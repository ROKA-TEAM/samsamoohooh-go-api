package kakao

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"samsamoohooh-go-api/internal/domain"
	"samsamoohooh-go-api/internal/infra/config"
)

const (
	authURL     = "https://kauth.kakao.com/oauth/authorize"
	tokenURL    = "https://kauth.kakao.com/oauth/token" //nolint:gosec
	userInfoAPI = "https://kapi.kakao.com/v2/user/me"
)

var _ domain.OauthAuthorizationGrantService = (*OauthKakaoService)(nil)

type OauthKakaoService struct {
	config       *config.Config
	oauthConfig  *oauth2.Config
	userService  domain.UserService
	tokenService domain.TokenService
}

func NewOauthKakaoService(
	config *config.Config,
	userService domain.UserService,
	tokenService domain.TokenService,
) *OauthKakaoService {
	return &OauthKakaoService{
		config: config,
		oauthConfig: &oauth2.Config{
			ClientID:     config.Oauth.Kakao.ClientID,
			ClientSecret: config.Oauth.Kakao.ClientSecret,
			RedirectURL:  config.Oauth.Kakao.CallbackURL,
			Endpoint: oauth2.Endpoint{
				AuthURL:  authURL,
				TokenURL: tokenURL,
			},
		},
		userService:  userService,
		tokenService: tokenService,
	}
}

func (s OauthKakaoService) GetLoginURL(state string) string {
	return s.oauthConfig.AuthCodeURL(state)
}

func (s OauthKakaoService) Exchange(ctx context.Context, code string) (*domain.OauthPayload, error) {
	token, err := s.oauthConfig.Exchange(ctx, code)
	if err != nil {
		return nil, errors.Wrap(domain.ErrInternal, err.Error())
	}

	client := s.oauthConfig.Client(ctx, token)
	resp, err := client.Get(userInfoAPI)
	if err != nil {
		return nil, errors.Wrap(domain.ErrInternal, err.Error())
	}

	var respBody exchangeRespBody
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return nil, errors.Wrap(domain.ErrInternal, err.Error())
	}

	err = resp.Body.Close()
	if err != nil {
		return nil, errors.Wrap(domain.ErrInternal, err.Error())
	}

	return respBody.toDomain(), nil
}

func (s OauthKakaoService) AuthenticateOrRegister(ctx context.Context, code string) (string, string, error) {
	payload, err := s.Exchange(ctx, code)
	if err != nil {
		return "", "", errors.Wrap(domain.ErrInternal, err.Error())
	}

	user, err := s.userService.GetBySub(ctx, payload.Sub)

	if errors.Is(err, domain.ErrNotFound) {
		createdUser, err := s.userService.Create(ctx, &domain.User{
			Name:      payload.Name,
			Role:      domain.UserRoleGuest,
			Social:    domain.UserSocialKaKao,
			SocialSub: payload.Sub,
		})
		if err != nil {
			return "", "", err
		}

		user = createdUser
	} else if err != nil {
		return "", "", err
	}
	// 전에 등록한 사용자이다.

	// 토큰을 발급한다.
	accessToken, err := s.tokenService.GenerateAccessTokenString(user.ID, domain.TokenRoleType(user.Role))
	if err != nil {
		return "", "", err
	}

	refreshToken, err := s.tokenService.GenerateRefreshTokenString(user.ID, domain.TokenRoleType(user.Role))
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil

}
