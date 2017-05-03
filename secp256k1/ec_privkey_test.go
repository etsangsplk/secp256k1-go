package secp256k1

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"testing"
)

type PrivkeyTweakAddTestCase struct {
	PrivateKey string `yaml:"privkey"`
	Tweak      string `yaml:"tweak"`
	Tweaked    string `yaml:"tweaked"`
}

func (t *PrivkeyTweakAddTestCase) GetPrivateKey() []byte {
	public, err := hex.DecodeString(t.PrivateKey)
	if err != nil {
		panic("Invalid private key")
	}
	return public
}
func (t *PrivkeyTweakAddTestCase) GetTweak() []byte {
	tweak, err := hex.DecodeString(t.Tweak)
	if err != nil {
		panic(err)
	}
	return tweak
}
func (t *PrivkeyTweakAddTestCase) GetTweaked() []byte {
	tweaked, err := hex.DecodeString(t.Tweaked)
	if err != nil {
		panic(err)
	}
	return tweaked
}

type PrivkeyTweakAddFixtures []PrivkeyTweakAddTestCase

func GetPrivkeyTweakAddFixtures() PrivkeyTweakAddFixtures {
	source := readFile(PrivkeyTweakAddTestVectors)
	testCase := PrivkeyTweakAddFixtures{}
	err := yaml.Unmarshal(source, &testCase)
	if err != nil {
		panic(err)
	}
	return testCase
}

func TestPrivkeyTweakAddFixtures(t *testing.T) {
	ctx, err := ContextCreate(ContextSign | ContextVerify)
	if err != nil {
		panic(err)
	}

	fixtures := GetPrivkeyTweakAddFixtures()

	for i := 0; i < 1; i++ {
		fixture := fixtures[i]
		priv := fixture.GetPrivateKey()
		tweak := fixture.GetTweak()

		r, err := EcPrivkeyTweakAdd(ctx, priv, tweak)
		spOK(t, r, err)

		assert.Equal(t, fixture.GetTweaked(), priv)
	}
}

type PrivkeyTweakMulTestCase struct {
	PrivateKey string `yaml:"privkey"`
	Tweak      string `yaml:"tweak"`
	Tweaked    string `yaml:"tweaked"`
}

func (t *PrivkeyTweakMulTestCase) GetPrivateKey() []byte {
	public, err := hex.DecodeString(t.PrivateKey)
	if err != nil {
		panic("Invalid private key")
	}
	return public
}
func (t *PrivkeyTweakMulTestCase) GetTweak() []byte {
	tweak, err := hex.DecodeString(t.Tweak)
	if err != nil {
		panic(err)
	}
	return tweak
}
func (t *PrivkeyTweakMulTestCase) GetTweaked() []byte {
	tweaked, err := hex.DecodeString(t.Tweaked)
	if err != nil {
		panic(err)
	}
	return tweaked
}

type PrivkeyTweakMulFixtures []PrivkeyTweakMulTestCase

func GetPrivkeyTweakMulFixtures() PrivkeyTweakMulFixtures {
	source := readFile(PrivkeyTweakMulTestVectors)
	testCase := PrivkeyTweakMulFixtures{}
	err := yaml.Unmarshal(source, &testCase)
	if err != nil {
		panic(err)
	}
	return testCase
}

func TestPrivkeyTweakMulFixtures(t *testing.T) {
	ctx, err := ContextCreate(ContextSign | ContextVerify)
	if err != nil {
		panic(err)
	}

	fixtures := GetPrivkeyTweakMulFixtures()

	for i := 0; i < 1; i++ {
		fixture := fixtures[i]
		priv := fixture.GetPrivateKey()
		tweak := fixture.GetTweak()

		r, err := EcPrivkeyTweakMul(ctx, priv, tweak)
		spOK(t, r, err)

		assert.Equal(t, fixture.GetTweaked(), priv)
	}
}

func TestPrivkeyVerifyFixtures(t *testing.T) {
	ctx, err := ContextCreate(ContextSign | ContextVerify)
	if err != nil {
		panic(err)
	}

	fixtures := GetPrivkeyTweakMulFixtures()

	for i := 0; i < 1; i++ {
		fixture := fixtures[i]
		priv := fixture.GetPrivateKey()
		result := EcSeckeyVerify(ctx, priv)
		assert.Equal(t, 1, result)
	}
}

func TestPrivkeyTweakAddChecksTweakSize(t *testing.T) {
	ctx, err := ContextCreate(ContextSign | ContextVerify)
	if err != nil {
		panic(err)
	}

	priv, _ := hex.DecodeString("e9a06e539d6bf5cf1ca5c41b59121fa3df07a338322405a312c67b6349a707e9")
	badTweak, _ := hex.DecodeString("AAAA")

	r, err := EcPrivkeyTweakAdd(ctx, priv, badTweak)
	assert.Error(t, err)
	assert.Equal(t, 0, r)
	assert.Equal(t, ErrorTweakSize, err.Error())
}

func TestPrivkeyTweakMulChecksTweakSize(t *testing.T) {
	ctx, err := ContextCreate(ContextSign | ContextVerify)
	if err != nil {
		panic(err)
	}

	priv, _ := hex.DecodeString("e9a06e539d6bf5cf1ca5c41b59121fa3df07a338322405a312c67b6349a707e9")
	badTweak, _ := hex.DecodeString("AAAA")

	r, err := EcPrivkeyTweakMul(ctx, priv, badTweak)
	assert.Error(t, err)
	assert.Equal(t, 0, r)
	assert.Equal(t, ErrorTweakSize, err.Error())
}


func TestPrivkeyTweakAddChecksPrivkeySize(t *testing.T) {
	ctx, err := ContextCreate(ContextSign | ContextVerify)
	if err != nil {
		panic(err)
	}

	tweak, _ := hex.DecodeString("e9a06e539d6bf5cf1ca5c41b59121fa3df07a338322405a312c67b6349a707e9")
	priv, _ := hex.DecodeString("AAAA")

	r, err := EcPrivkeyTweakAdd(ctx, priv, tweak)
	assert.Error(t, err)
	assert.Equal(t, 0, r)
	assert.Equal(t, ErrorPrivateKeySize, err.Error())
}

func TestPrivkeyTweakMulChecksPrivkeySize(t *testing.T) {
	ctx, err := ContextCreate(ContextSign | ContextVerify)
	if err != nil {
		panic(err)
	}

	priv, _ := hex.DecodeString("AAAA")
	tweak, _ := hex.DecodeString("e9a06e539d6bf5cf1ca5c41b59121fa3df07a338322405a312c67b6349a707e9")

	r, err := EcPrivkeyTweakMul(ctx, priv, tweak)
	assert.Error(t, err)
	assert.Equal(t, 0, r)
	assert.Equal(t, ErrorPrivateKeySize, err.Error())
}