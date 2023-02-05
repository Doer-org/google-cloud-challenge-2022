package config

import (
	"testing"
)

func TestDsn_Dsn(t *testing.T) {
	// FIXME:
	t.Skip()
	tests := []struct {
		name    string
		wantDsn string
		wantErr error
	}{
		{
			name:    "必要な環境変数がすべて設定されている場合、DSNを返す",
			wantDsn: "user=hoge password=hoge host=hoge port=hoge dbname=hoge sslmode=disable",
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("POSTGRES_HOST", "hoge")
			t.Setenv("POSTGRES_PORT", "hoge")
			t.Setenv("POSTGRES_USER", "hoge")
			t.Setenv("POSTGRES_PASSWORD", "hoge")
			t.Setenv("POSTGRES_DB", "hoge")

			gotDSN, gotErr := DSN()
			if tt.wantDsn != gotDSN {
				t.Errorf("TestDsn_Dsn Error : wantDsb %s, but gotDsb %s", tt.wantDsn, gotDSN)
			}

			if tt.wantErr != gotErr {
				t.Errorf("TestDsn_Dsn Error : wantErr %s, but gotErr %s", tt.wantErr, gotErr)
			}
		})
	}
}
