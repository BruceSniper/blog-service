/**
 * @Author:      zhangyuyao
 * @Description: TODO
 * @File:        setting.go
 * @Version:     1.0.0
 * @Date:        2021/8/22 3:26 下午
 */

package setting

import "github.com/spf13/viper"

type Setting struct {
	vp *viper.Viper
}

func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &Setting{vp}, nil
}
