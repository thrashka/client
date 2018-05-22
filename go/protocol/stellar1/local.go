// Auto-generated by avdl-compiler v1.3.22 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/stellar1/local.avdl

package stellar1

import (
	keybase1 "github.com/keybase/client/go/protocol/keybase1"
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type WalletAccountLocal struct {
	AccountID          AccountID `codec:"accountID" json:"accountID"`
	IsDefault          bool      `codec:"isDefault" json:"isDefault"`
	Name               string    `codec:"name" json:"name"`
	BalanceDescription string    `codec:"balanceDescription" json:"balanceDescription"`
}

func (o WalletAccountLocal) DeepCopy() WalletAccountLocal {
	return WalletAccountLocal{
		AccountID:          o.AccountID.DeepCopy(),
		IsDefault:          o.IsDefault,
		Name:               o.Name,
		BalanceDescription: o.BalanceDescription,
	}
}

type AccountAssetLocal struct {
	Name                   string `codec:"name" json:"name"`
	BalanceTotal           string `codec:"balanceTotal" json:"balanceTotal"`
	BalanceAvailableToSend string `codec:"balanceAvailableToSend" json:"balanceAvailableToSend"`
	AssetCode              string `codec:"assetCode" json:"assetCode"`
	Issuer                 string `codec:"issuer" json:"issuer"`
	Worth                  string `codec:"worth" json:"worth"`
	WorthCurrency          string `codec:"worthCurrency" json:"worthCurrency"`
}

func (o AccountAssetLocal) DeepCopy() AccountAssetLocal {
	return AccountAssetLocal{
		Name:                   o.Name,
		BalanceTotal:           o.BalanceTotal,
		BalanceAvailableToSend: o.BalanceAvailableToSend,
		AssetCode:              o.AssetCode,
		Issuer:                 o.Issuer,
		Worth:                  o.Worth,
		WorthCurrency:          o.WorthCurrency,
	}
}

type SendResultCLILocal struct {
	KbTxID KeybaseTransactionID     `codec:"KbTxID" json:"KbTxID"`
	TxID   TransactionID            `codec:"TxID" json:"TxID"`
	Relay  *SendRelayResultCLILocal `codec:"relay,omitempty" json:"relay,omitempty"`
}

func (o SendResultCLILocal) DeepCopy() SendResultCLILocal {
	return SendResultCLILocal{
		KbTxID: o.KbTxID.DeepCopy(),
		TxID:   o.TxID.DeepCopy(),
		Relay: (func(x *SendRelayResultCLILocal) *SendRelayResultCLILocal {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Relay),
	}
}

type SendRelayResultCLILocal struct {
	TeamID keybase1.TeamID `codec:"teamID" json:"teamID"`
}

func (o SendRelayResultCLILocal) DeepCopy() SendRelayResultCLILocal {
	return SendRelayResultCLILocal{
		TeamID: o.TeamID.DeepCopy(),
	}
}

type PaymentCLIOptionLocal struct {
	Payment *PaymentCLILocal `codec:"payment,omitempty" json:"payment,omitempty"`
	Err     string           `codec:"err" json:"err"`
}

func (o PaymentCLIOptionLocal) DeepCopy() PaymentCLIOptionLocal {
	return PaymentCLIOptionLocal{
		Payment: (func(x *PaymentCLILocal) *PaymentCLILocal {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Payment),
		Err: o.Err,
	}
}

type PaymentCLILocal struct {
	TxID            TransactionID `codec:"txID" json:"txID"`
	Time            TimeMs        `codec:"time" json:"time"`
	Status          string        `codec:"status" json:"status"`
	StatusDetail    string        `codec:"statusDetail" json:"statusDetail"`
	Amount          string        `codec:"amount" json:"amount"`
	Asset           Asset         `codec:"asset" json:"asset"`
	DisplayAmount   *string       `codec:"displayAmount,omitempty" json:"displayAmount,omitempty"`
	DisplayCurrency *string       `codec:"displayCurrency,omitempty" json:"displayCurrency,omitempty"`
	FromStellar     AccountID     `codec:"fromStellar" json:"fromStellar"`
	ToStellar       *AccountID    `codec:"toStellar,omitempty" json:"toStellar,omitempty"`
	FromUsername    *string       `codec:"fromUsername,omitempty" json:"fromUsername,omitempty"`
	ToUsername      *string       `codec:"toUsername,omitempty" json:"toUsername,omitempty"`
	Note            string        `codec:"note" json:"note"`
	NoteErr         string        `codec:"noteErr" json:"noteErr"`
}

func (o PaymentCLILocal) DeepCopy() PaymentCLILocal {
	return PaymentCLILocal{
		TxID:         o.TxID.DeepCopy(),
		Time:         o.Time.DeepCopy(),
		Status:       o.Status,
		StatusDetail: o.StatusDetail,
		Amount:       o.Amount,
		Asset:        o.Asset.DeepCopy(),
		DisplayAmount: (func(x *string) *string {
			if x == nil {
				return nil
			}
			tmp := (*x)
			return &tmp
		})(o.DisplayAmount),
		DisplayCurrency: (func(x *string) *string {
			if x == nil {
				return nil
			}
			tmp := (*x)
			return &tmp
		})(o.DisplayCurrency),
		FromStellar: o.FromStellar.DeepCopy(),
		ToStellar: (func(x *AccountID) *AccountID {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.ToStellar),
		FromUsername: (func(x *string) *string {
			if x == nil {
				return nil
			}
			tmp := (*x)
			return &tmp
		})(o.FromUsername),
		ToUsername: (func(x *string) *string {
			if x == nil {
				return nil
			}
			tmp := (*x)
			return &tmp
		})(o.ToUsername),
		Note:    o.Note,
		NoteErr: o.NoteErr,
	}
}

type OwnAccountCLILocal struct {
	AccountID    AccountID            `codec:"accountID" json:"accountID"`
	IsPrimary    bool                 `codec:"isPrimary" json:"isPrimary"`
	Name         string               `codec:"name" json:"name"`
	Balance      []Balance            `codec:"balance" json:"balance"`
	ExchangeRate *OutsideExchangeRate `codec:"exchangeRate,omitempty" json:"exchangeRate,omitempty"`
}

func (o OwnAccountCLILocal) DeepCopy() OwnAccountCLILocal {
	return OwnAccountCLILocal{
		AccountID: o.AccountID.DeepCopy(),
		IsPrimary: o.IsPrimary,
		Name:      o.Name,
		Balance: (func(x []Balance) []Balance {
			if x == nil {
				return nil
			}
			var ret []Balance
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Balance),
		ExchangeRate: (func(x *OutsideExchangeRate) *OutsideExchangeRate {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.ExchangeRate),
	}
}

type GetWalletAccountsLocalArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
}

type GetAccountAssetsLocalArg struct {
	SessionID int       `codec:"sessionID" json:"sessionID"`
	AccountID AccountID `codec:"accountID" json:"accountID"`
}

type ChangeWalletAccountNameLocalArg struct {
	SessionID int       `codec:"sessionID" json:"sessionID"`
	AccountID AccountID `codec:"accountID" json:"accountID"`
	NewName   string    `codec:"newName" json:"newName"`
}

type BalancesLocalArg struct {
	AccountID AccountID `codec:"accountID" json:"accountID"`
}

type SendCLILocalArg struct {
	Recipient       string `codec:"recipient" json:"recipient"`
	Amount          string `codec:"amount" json:"amount"`
	Asset           Asset  `codec:"asset" json:"asset"`
	Note            string `codec:"note" json:"note"`
	DisplayAmount   string `codec:"displayAmount" json:"displayAmount"`
	DisplayCurrency string `codec:"displayCurrency" json:"displayCurrency"`
}

type ClaimCLILocalArg struct {
	TxID string     `codec:"txID" json:"txID"`
	Into *AccountID `codec:"into,omitempty" json:"into,omitempty"`
}

type RecentPaymentsCLILocalArg struct {
	AccountID *AccountID `codec:"accountID,omitempty" json:"accountID,omitempty"`
}

type PaymentDetailCLILocalArg struct {
	TxID string `codec:"txID" json:"txID"`
}

type WalletInitLocalArg struct {
}

type WalletDumpLocalArg struct {
}

type WalletGetAccountsCLILocalArg struct {
}

type OwnAccountLocalArg struct {
	AccountID AccountID `codec:"accountID" json:"accountID"`
}

type ImportSecretKeyLocalArg struct {
	SecretKey   SecretKey `codec:"secretKey" json:"secretKey"`
	MakePrimary bool      `codec:"makePrimary" json:"makePrimary"`
}

type ExportSecretKeyLocalArg struct {
	AccountID AccountID `codec:"accountID" json:"accountID"`
}

type SetDisplayCurrencyArg struct {
	AccountID AccountID `codec:"accountID" json:"accountID"`
	Currency  string    `codec:"currency" json:"currency"`
}

type ExchangeRateLocalArg struct {
	Currency OutsideCurrencyCode `codec:"currency" json:"currency"`
}

type GetAvailableLocalCurrenciesArg struct {
}

type FormatLocalCurrencyStringArg struct {
	Amount string              `codec:"amount" json:"amount"`
	Code   OutsideCurrencyCode `codec:"code" json:"code"`
}

type LocalInterface interface {
	GetWalletAccountsLocal(context.Context, int) ([]WalletAccountLocal, error)
	GetAccountAssetsLocal(context.Context, GetAccountAssetsLocalArg) ([]AccountAssetLocal, error)
	ChangeWalletAccountNameLocal(context.Context, ChangeWalletAccountNameLocalArg) error
	BalancesLocal(context.Context, AccountID) ([]Balance, error)
	SendCLILocal(context.Context, SendCLILocalArg) (SendResultCLILocal, error)
	ClaimCLILocal(context.Context, ClaimCLILocalArg) (RelayClaimResult, error)
	RecentPaymentsCLILocal(context.Context, *AccountID) ([]PaymentCLIOptionLocal, error)
	PaymentDetailCLILocal(context.Context, string) (PaymentCLILocal, error)
	WalletInitLocal(context.Context) error
	WalletDumpLocal(context.Context) (Bundle, error)
	WalletGetAccountsCLILocal(context.Context) ([]OwnAccountCLILocal, error)
	OwnAccountLocal(context.Context, AccountID) (bool, error)
	ImportSecretKeyLocal(context.Context, ImportSecretKeyLocalArg) error
	ExportSecretKeyLocal(context.Context, AccountID) (SecretKey, error)
	SetDisplayCurrency(context.Context, SetDisplayCurrencyArg) error
	ExchangeRateLocal(context.Context, OutsideCurrencyCode) (OutsideExchangeRate, error)
	GetAvailableLocalCurrencies(context.Context) (map[OutsideCurrencyCode]OutsideCurrencyDefinition, error)
	FormatLocalCurrencyString(context.Context, FormatLocalCurrencyStringArg) (string, error)
}

func LocalProtocol(i LocalInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "stellar.1.local",
		Methods: map[string]rpc.ServeHandlerDescription{
			"GetWalletAccountsLocal": {
				MakeArg: func() interface{} {
					ret := make([]GetWalletAccountsLocalArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]GetWalletAccountsLocalArg)
					if !ok {
						err = rpc.NewTypeError((*[]GetWalletAccountsLocalArg)(nil), args)
						return
					}
					ret, err = i.GetWalletAccountsLocal(ctx, (*typedArgs)[0].SessionID)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"GetAccountAssetsLocal": {
				MakeArg: func() interface{} {
					ret := make([]GetAccountAssetsLocalArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]GetAccountAssetsLocalArg)
					if !ok {
						err = rpc.NewTypeError((*[]GetAccountAssetsLocalArg)(nil), args)
						return
					}
					ret, err = i.GetAccountAssetsLocal(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"changeWalletAccountNameLocal": {
				MakeArg: func() interface{} {
					ret := make([]ChangeWalletAccountNameLocalArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ChangeWalletAccountNameLocalArg)
					if !ok {
						err = rpc.NewTypeError((*[]ChangeWalletAccountNameLocalArg)(nil), args)
						return
					}
					err = i.ChangeWalletAccountNameLocal(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"balancesLocal": {
				MakeArg: func() interface{} {
					ret := make([]BalancesLocalArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]BalancesLocalArg)
					if !ok {
						err = rpc.NewTypeError((*[]BalancesLocalArg)(nil), args)
						return
					}
					ret, err = i.BalancesLocal(ctx, (*typedArgs)[0].AccountID)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"sendCLILocal": {
				MakeArg: func() interface{} {
					ret := make([]SendCLILocalArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]SendCLILocalArg)
					if !ok {
						err = rpc.NewTypeError((*[]SendCLILocalArg)(nil), args)
						return
					}
					ret, err = i.SendCLILocal(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"claimCLILocal": {
				MakeArg: func() interface{} {
					ret := make([]ClaimCLILocalArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ClaimCLILocalArg)
					if !ok {
						err = rpc.NewTypeError((*[]ClaimCLILocalArg)(nil), args)
						return
					}
					ret, err = i.ClaimCLILocal(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"recentPaymentsCLILocal": {
				MakeArg: func() interface{} {
					ret := make([]RecentPaymentsCLILocalArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]RecentPaymentsCLILocalArg)
					if !ok {
						err = rpc.NewTypeError((*[]RecentPaymentsCLILocalArg)(nil), args)
						return
					}
					ret, err = i.RecentPaymentsCLILocal(ctx, (*typedArgs)[0].AccountID)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"paymentDetailCLILocal": {
				MakeArg: func() interface{} {
					ret := make([]PaymentDetailCLILocalArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]PaymentDetailCLILocalArg)
					if !ok {
						err = rpc.NewTypeError((*[]PaymentDetailCLILocalArg)(nil), args)
						return
					}
					ret, err = i.PaymentDetailCLILocal(ctx, (*typedArgs)[0].TxID)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"walletInitLocal": {
				MakeArg: func() interface{} {
					ret := make([]WalletInitLocalArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					err = i.WalletInitLocal(ctx)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"walletDumpLocal": {
				MakeArg: func() interface{} {
					ret := make([]WalletDumpLocalArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					ret, err = i.WalletDumpLocal(ctx)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"walletGetAccountsCLILocal": {
				MakeArg: func() interface{} {
					ret := make([]WalletGetAccountsCLILocalArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					ret, err = i.WalletGetAccountsCLILocal(ctx)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"ownAccountLocal": {
				MakeArg: func() interface{} {
					ret := make([]OwnAccountLocalArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]OwnAccountLocalArg)
					if !ok {
						err = rpc.NewTypeError((*[]OwnAccountLocalArg)(nil), args)
						return
					}
					ret, err = i.OwnAccountLocal(ctx, (*typedArgs)[0].AccountID)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"importSecretKeyLocal": {
				MakeArg: func() interface{} {
					ret := make([]ImportSecretKeyLocalArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ImportSecretKeyLocalArg)
					if !ok {
						err = rpc.NewTypeError((*[]ImportSecretKeyLocalArg)(nil), args)
						return
					}
					err = i.ImportSecretKeyLocal(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"exportSecretKeyLocal": {
				MakeArg: func() interface{} {
					ret := make([]ExportSecretKeyLocalArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ExportSecretKeyLocalArg)
					if !ok {
						err = rpc.NewTypeError((*[]ExportSecretKeyLocalArg)(nil), args)
						return
					}
					ret, err = i.ExportSecretKeyLocal(ctx, (*typedArgs)[0].AccountID)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"setDisplayCurrency": {
				MakeArg: func() interface{} {
					ret := make([]SetDisplayCurrencyArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]SetDisplayCurrencyArg)
					if !ok {
						err = rpc.NewTypeError((*[]SetDisplayCurrencyArg)(nil), args)
						return
					}
					err = i.SetDisplayCurrency(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"exchangeRateLocal": {
				MakeArg: func() interface{} {
					ret := make([]ExchangeRateLocalArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ExchangeRateLocalArg)
					if !ok {
						err = rpc.NewTypeError((*[]ExchangeRateLocalArg)(nil), args)
						return
					}
					ret, err = i.ExchangeRateLocal(ctx, (*typedArgs)[0].Currency)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"getAvailableLocalCurrencies": {
				MakeArg: func() interface{} {
					ret := make([]GetAvailableLocalCurrenciesArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					ret, err = i.GetAvailableLocalCurrencies(ctx)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"formatLocalCurrencyString": {
				MakeArg: func() interface{} {
					ret := make([]FormatLocalCurrencyStringArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]FormatLocalCurrencyStringArg)
					if !ok {
						err = rpc.NewTypeError((*[]FormatLocalCurrencyStringArg)(nil), args)
						return
					}
					ret, err = i.FormatLocalCurrencyString(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type LocalClient struct {
	Cli rpc.GenericClient
}

func (c LocalClient) GetWalletAccountsLocal(ctx context.Context, sessionID int) (res []WalletAccountLocal, err error) {
	__arg := GetWalletAccountsLocalArg{SessionID: sessionID}
	err = c.Cli.Call(ctx, "stellar.1.local.GetWalletAccountsLocal", []interface{}{__arg}, &res)
	return
}

func (c LocalClient) GetAccountAssetsLocal(ctx context.Context, __arg GetAccountAssetsLocalArg) (res []AccountAssetLocal, err error) {
	err = c.Cli.Call(ctx, "stellar.1.local.GetAccountAssetsLocal", []interface{}{__arg}, &res)
	return
}

func (c LocalClient) ChangeWalletAccountNameLocal(ctx context.Context, __arg ChangeWalletAccountNameLocalArg) (err error) {
	err = c.Cli.Call(ctx, "stellar.1.local.changeWalletAccountNameLocal", []interface{}{__arg}, nil)
	return
}

func (c LocalClient) BalancesLocal(ctx context.Context, accountID AccountID) (res []Balance, err error) {
	__arg := BalancesLocalArg{AccountID: accountID}
	err = c.Cli.Call(ctx, "stellar.1.local.balancesLocal", []interface{}{__arg}, &res)
	return
}

func (c LocalClient) SendCLILocal(ctx context.Context, __arg SendCLILocalArg) (res SendResultCLILocal, err error) {
	err = c.Cli.Call(ctx, "stellar.1.local.sendCLILocal", []interface{}{__arg}, &res)
	return
}

func (c LocalClient) ClaimCLILocal(ctx context.Context, __arg ClaimCLILocalArg) (res RelayClaimResult, err error) {
	err = c.Cli.Call(ctx, "stellar.1.local.claimCLILocal", []interface{}{__arg}, &res)
	return
}

func (c LocalClient) RecentPaymentsCLILocal(ctx context.Context, accountID *AccountID) (res []PaymentCLIOptionLocal, err error) {
	__arg := RecentPaymentsCLILocalArg{AccountID: accountID}
	err = c.Cli.Call(ctx, "stellar.1.local.recentPaymentsCLILocal", []interface{}{__arg}, &res)
	return
}

func (c LocalClient) PaymentDetailCLILocal(ctx context.Context, txID string) (res PaymentCLILocal, err error) {
	__arg := PaymentDetailCLILocalArg{TxID: txID}
	err = c.Cli.Call(ctx, "stellar.1.local.paymentDetailCLILocal", []interface{}{__arg}, &res)
	return
}

func (c LocalClient) WalletInitLocal(ctx context.Context) (err error) {
	err = c.Cli.Call(ctx, "stellar.1.local.walletInitLocal", []interface{}{WalletInitLocalArg{}}, nil)
	return
}

func (c LocalClient) WalletDumpLocal(ctx context.Context) (res Bundle, err error) {
	err = c.Cli.Call(ctx, "stellar.1.local.walletDumpLocal", []interface{}{WalletDumpLocalArg{}}, &res)
	return
}

func (c LocalClient) WalletGetAccountsCLILocal(ctx context.Context) (res []OwnAccountCLILocal, err error) {
	err = c.Cli.Call(ctx, "stellar.1.local.walletGetAccountsCLILocal", []interface{}{WalletGetAccountsCLILocalArg{}}, &res)
	return
}

func (c LocalClient) OwnAccountLocal(ctx context.Context, accountID AccountID) (res bool, err error) {
	__arg := OwnAccountLocalArg{AccountID: accountID}
	err = c.Cli.Call(ctx, "stellar.1.local.ownAccountLocal", []interface{}{__arg}, &res)
	return
}

func (c LocalClient) ImportSecretKeyLocal(ctx context.Context, __arg ImportSecretKeyLocalArg) (err error) {
	err = c.Cli.Call(ctx, "stellar.1.local.importSecretKeyLocal", []interface{}{__arg}, nil)
	return
}

func (c LocalClient) ExportSecretKeyLocal(ctx context.Context, accountID AccountID) (res SecretKey, err error) {
	__arg := ExportSecretKeyLocalArg{AccountID: accountID}
	err = c.Cli.Call(ctx, "stellar.1.local.exportSecretKeyLocal", []interface{}{__arg}, &res)
	return
}

func (c LocalClient) SetDisplayCurrency(ctx context.Context, __arg SetDisplayCurrencyArg) (err error) {
	err = c.Cli.Call(ctx, "stellar.1.local.setDisplayCurrency", []interface{}{__arg}, nil)
	return
}

func (c LocalClient) ExchangeRateLocal(ctx context.Context, currency OutsideCurrencyCode) (res OutsideExchangeRate, err error) {
	__arg := ExchangeRateLocalArg{Currency: currency}
	err = c.Cli.Call(ctx, "stellar.1.local.exchangeRateLocal", []interface{}{__arg}, &res)
	return
}

func (c LocalClient) GetAvailableLocalCurrencies(ctx context.Context) (res map[OutsideCurrencyCode]OutsideCurrencyDefinition, err error) {
	err = c.Cli.Call(ctx, "stellar.1.local.getAvailableLocalCurrencies", []interface{}{GetAvailableLocalCurrenciesArg{}}, &res)
	return
}

func (c LocalClient) FormatLocalCurrencyString(ctx context.Context, __arg FormatLocalCurrencyStringArg) (res string, err error) {
	err = c.Cli.Call(ctx, "stellar.1.local.formatLocalCurrencyString", []interface{}{__arg}, &res)
	return
}
