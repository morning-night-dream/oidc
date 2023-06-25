package rp

type RP struct {
	ClientID    string   // RPを識別するためのID
	RedirectURI string   // ログイン後にリダイレクトさせるURL
	Scopes      []string //
	AuthURL     string   // OPの認証エンドポイント
	TokenURL    string   // OPのトークンエンドポイント
	UserInfoURL string   // OPのユーザー情報エンドポイント
}
