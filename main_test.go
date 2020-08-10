package structenv

import (
	"os"
	"testing"
)

// func Test_SetFieldFromEnv_Success(t *testing.T) {
// 	const testEnv = "TEST_ENV"

// 	_ = os.Setenv(testEnv, "testenv")
// 	defer os.Unsetenv(testEnv)

// 	s := testStruct{
// 		Test: testEnv,
// 	}

// 	if err := setFieldFromEnv(reflect.ValueOf(&s).Elem().Field(0)); err != nil {
// 		t.Fatal(err)
// 	}

// 	if s.Test != "testenv" {
// 		t.Errorf("кек")
// 	}
// }

func Test_SetFromEnvs_Success(t *testing.T) {
	type goodStruct struct {
		User     string
		Password string `env:""`
	}

	gs := goodStruct{
		User:     "nobody",
		Password: "MY_PASSWORD",
	}

	os.Setenv("MY_PASSWORD", "password")
	defer os.Unsetenv("MY_PASSWORD")

	if err := SetFromEnvs(&gs); err != nil {
		t.Error(err)
	}

	if gs.Password != "password" {
		t.Errorf("expected: password; got: %s", gs.Password)
	}
}

func Test_SetFromEnvs_Success_NestedStructs(t *testing.T) {
	type Creds struct {
		User     string
		Password string `env:""`
	}

	type goodStruct struct {
		Creds Creds
		Old   int
	}

	gs := goodStruct{
		Creds: Creds{
			User:     "nobody",
			Password: "MY_PASSWORD",
		},
		Old: 100,
	}

	os.Setenv("MY_PASSWORD", "password")
	defer os.Unsetenv("MY_PASSWORD")

	if err := SetFromEnvs(&gs); err != nil {
		t.Error(err)
	}

	if gs.Creds.Password != "password" {
		t.Errorf("expected: password; got: %s", gs.Creds.Password)
	}

}

func Test_SetFromEnvs_Fail_TypeError(t *testing.T) {
	type badStruct struct {
		User     string
		Password int `env:"MY_PASSWORD"`
	}

	gs := badStruct{
		User: "nobody",
	}

	os.Setenv("MY_PASSWORD", "password")
	defer os.Unsetenv("MY_PASSWORD")

	if err := SetFromEnvs(&gs); err != nil {
		if _, ok := err.(TypeError); !ok {
			t.Error(err)
		}
	} else {
		t.Errorf("expected TypeError raising")
	}
}

func Test_SetFromEnvs_Fail_EnvError(t *testing.T) {
	type goodStruct struct {
		User     string
		Password string `env:"MY_PASSWORD"`
	}

	gs := goodStruct{
		User: "nobody",
	}

	if err := SetFromEnvs(&gs); err != nil {
		if _, ok := err.(EnvError); !ok {
			t.Error(err)
		}
	} else {
		t.Errorf("expected EnvError raising")
	}
}

func Test_SetFromEnvs_Fail_InNestedStructs(t *testing.T) {
	type Creds struct {
		User     string
		Password string `env:""`
	}

	type goodStruct struct {
		Creds Creds
		Old   int
	}

	gs := goodStruct{
		Creds: Creds{
			User:     "nobody",
			Password: "MY_PASSWORD",
		},
		Old: 100,
	}

	if err := SetFromEnvs(&gs); err != nil {
		if _, ok := err.(EnvError); !ok {
			t.Error(err)
		}
	} else {
		t.Errorf("expected EnvError raising")
	}
}
