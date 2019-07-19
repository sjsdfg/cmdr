/*
 * Copyright © 2019 Hedzr Yeh.
 */

package cmdr

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

// NewOptions returns an `Options` structure pointer
func NewOptions() *Options {
	return &Options{
		entries:   make(map[string]interface{}),
		hierarchy: make(map[string]interface{}),
		rw:        new(sync.RWMutex),
	}
}

// NewOptionsWith returns an `Options` structure pointer
func NewOptionsWith(entries map[string]interface{}) *Options {
	return &Options{
		entries:   entries,
		hierarchy: make(map[string]interface{}),
		rw:        new(sync.RWMutex),
	}
}

//
//
//

// GetBool returns the bool value of an `Option` key. Such as:
// ```golang
// cmdr.Get("app.logger.enable") => true,...
// ```
//
func GetBool(key string) bool {
	return rxxtOptions.GetBool(key)
}

// GetBoolP returns the bool value of an `Option` key. Such as:
// ```golang
// cmdr.GetP("app.logger", "enable") => true,...
// ```
func GetBoolP(prefix, key string) bool {
	return rxxtOptions.GetBool(fmt.Sprintf("%s.%s", prefix, key))
}

// GetBoolR returns the bool value of an `Option` key with [WrapWithRxxtPrefix]. Such as:
// ```golang
// cmdr.GetBoolR("logger.enable") => true,...
// ```
//
func GetBoolR(key string) bool {
	return rxxtOptions.GetBool(wrapWithRxxtPrefix(key))
}

// GetBoolRP returns the bool value of an `Option` key with [WrapWithRxxtPrefix]. Such as:
// ```golang
// cmdr.GetBoolRP("logger", "enable") => true,...
// ```
func GetBoolRP(prefix, key string) bool {
	return rxxtOptions.GetBool(wrapWithRxxtPrefix(fmt.Sprintf("%s.%s", prefix, key)))
}

// GetInt returns the int value of an `Option` key.
func GetInt(key string) int {
	return int(rxxtOptions.GetInt(key))
}

// GetIntP returns the int value of an `Option` key.
func GetIntP(prefix, key string) int {
	return int(rxxtOptions.GetInt(fmt.Sprintf("%s.%s", prefix, key)))
}

// GetIntR returns the int value of an `Option` key with [WrapWithRxxtPrefix].
func GetIntR(key string) int {
	return int(rxxtOptions.GetInt(wrapWithRxxtPrefix(key)))
}

// GetIntRP returns the int value of an `Option` key with [WrapWithRxxtPrefix].
func GetIntRP(prefix, key string) int {
	return int(rxxtOptions.GetInt(wrapWithRxxtPrefix(fmt.Sprintf("%s.%s", prefix, key))))
}

// GetInt64 returns the int64 value of an `Option` key.
func GetInt64(key string) int64 {
	return rxxtOptions.GetInt(key)
}

// GetInt64P returns the int64 value of an `Option` key.
func GetInt64P(prefix, key string) int64 {
	return rxxtOptions.GetInt(fmt.Sprintf("%s.%s", prefix, key))
}

// GetInt64R returns the int64 value of an `Option` key with [WrapWithRxxtPrefix].
func GetInt64R(key string) int64 {
	return rxxtOptions.GetInt(wrapWithRxxtPrefix(key))
}

// GetInt64RP returns the int64 value of an `Option` key with [WrapWithRxxtPrefix].
func GetInt64RP(prefix, key string) int64 {
	return rxxtOptions.GetInt(wrapWithRxxtPrefix(fmt.Sprintf("%s.%s", prefix, key)))
}

// GetUint returns the uint value of an `Option` key.
func GetUint(key string) uint {
	return uint(rxxtOptions.GetUint(key))
}

// GetUintP returns the uint value of an `Option` key.
func GetUintP(prefix, key string) uint {
	return uint(rxxtOptions.GetUint(fmt.Sprintf("%s.%s", prefix, key)))
}

// GetUintR returns the uint value of an `Option` key with [WrapWithRxxtPrefix].
func GetUintR(key string) uint {
	return uint(rxxtOptions.GetUint(wrapWithRxxtPrefix(key)))
}

// GetUintRP returns the uint value of an `Option` key with [WrapWithRxxtPrefix].
func GetUintRP(prefix, key string) uint {
	return uint(rxxtOptions.GetUint(wrapWithRxxtPrefix(fmt.Sprintf("%s.%s", prefix, key))))
}

// GetUint64 returns the uint64 value of an `Option` key.
func GetUint64(key string) uint64 {
	return rxxtOptions.GetUint(key)
}

// GetUint64P returns the uint64 value of an `Option` key.
func GetUint64P(prefix, key string) uint64 {
	return rxxtOptions.GetUint(fmt.Sprintf("%s.%s", prefix, key))
}

// GetUint64R returns the uint64 value of an `Option` key with [WrapWithRxxtPrefix].
func GetUint64R(key string) uint64 {
	return rxxtOptions.GetUint(wrapWithRxxtPrefix(key))
}

// GetUint64RP returns the uint64 value of an `Option` key with [WrapWithRxxtPrefix].
func GetUint64RP(prefix, key string) uint64 {
	return rxxtOptions.GetUint(wrapWithRxxtPrefix(fmt.Sprintf("%s.%s", prefix, key)))
}

// GetFloat32 returns the float32 value of an `Option` key.
func GetFloat32(key string) float32 {
	return float32(rxxtOptions.GetFloat32(key))
}

// GetFloat32P returns the float32 value of an `Option` key.
func GetFloat32P(prefix, key string) float32 {
	return float32(rxxtOptions.GetFloat32(fmt.Sprintf("%s.%s", prefix, key)))
}

// GetFloat32R returns the float32 value of an `Option` key with [WrapWithRxxtPrefix].
func GetFloat32R(key string) float32 {
	return float32(rxxtOptions.GetFloat32(wrapWithRxxtPrefix(key)))
}

// GetFloat32RP returns the float32 value of an `Option` key with [WrapWithRxxtPrefix].
func GetFloat32RP(prefix, key string) float32 {
	return float32(rxxtOptions.GetFloat32(wrapWithRxxtPrefix(fmt.Sprintf("%s.%s", prefix, key))))
}

// GetFloat64 returns the float64 value of an `Option` key.
func GetFloat64(key string) float64 {
	return rxxtOptions.GetFloat64(key)
}

// GetFloat64P returns the float64 value of an `Option` key.
func GetFloat64P(prefix, key string) float64 {
	return rxxtOptions.GetFloat64(fmt.Sprintf("%s.%s", prefix, key))
}

// GetFloat64R returns the float64 value of an `Option` key with [WrapWithRxxtPrefix].
func GetFloat64R(key string) float64 {
	return rxxtOptions.GetFloat64(wrapWithRxxtPrefix(key))
}

// GetFloat64RP returns the float64 value of an `Option` key with [WrapWithRxxtPrefix].
func GetFloat64RP(prefix, key string) float64 {
	return rxxtOptions.GetFloat64(wrapWithRxxtPrefix(fmt.Sprintf("%s.%s", prefix, key)))
}

// GetString returns the string value of an `Option` key.
func GetString(key string) string {
	return rxxtOptions.GetString(key)
}

// GetStringP returns the string value of an `Option` key.
func GetStringP(prefix, key string) string {
	return rxxtOptions.GetString(fmt.Sprintf("%s.%s", prefix, key))
}

// GetStringR returns the string value of an `Option` key with [WrapWithRxxtPrefix].
func GetStringR(key string) string {
	return rxxtOptions.GetString(wrapWithRxxtPrefix(key))
}

// GetStringRP returns the string value of an `Option` key with [WrapWithRxxtPrefix].
func GetStringRP(prefix, key string) string {
	return rxxtOptions.GetString(wrapWithRxxtPrefix(fmt.Sprintf("%s.%s", prefix, key)))
}

// GetStringSlice returns the string slice value of an `Option` key.
func GetStringSlice(key string) []string {
	return rxxtOptions.GetStringSlice(key)
}

// GetStringSliceP returns the string slice value of an `Option` key.
func GetStringSliceP(prefix, key string) []string {
	return rxxtOptions.GetStringSlice(fmt.Sprintf("%s.%s", prefix, key))
}

// GetStringSliceR returns the string slice value of an `Option` key with [WrapWithRxxtPrefix].
func GetStringSliceR(key string) []string {
	return rxxtOptions.GetStringSlice(wrapWithRxxtPrefix(key))
}

// GetStringSliceRP returns the string slice value of an `Option` key with [WrapWithRxxtPrefix].
func GetStringSliceRP(prefix, key string) []string {
	return rxxtOptions.GetStringSlice(wrapWithRxxtPrefix(fmt.Sprintf("%s.%s", prefix, key)))
}

// Get returns the generic value of an `Option` key with [WrapWithRxxtPrefix]. Such as:
// ```golang
// cmdr.Get("app.logger.level") => 'DEBUG',...
// ```
//
func Get(key string) interface{} {
	return rxxtOptions.Get(key)
}

// GetR returns the generic value of an `Option` key with [WrapWithRxxtPrefix]. Such as:
// ```golang
// cmdr.GetR("logger.level") => 'DEBUG',...
// ```
//
func GetR(key string) interface{} {
	return rxxtOptions.Get(wrapWithRxxtPrefix(key))
}

// GetMap an `Option` by key string, it returns a hierarchy map or nil
func GetMap(key string) map[string]interface{} {
	return rxxtOptions.GetMap(key)
}

// GetMapR an `Option` by key string with [WrapWithRxxtPrefix], it returns a hierarchy map or nil
func GetMapR(key string) map[string]interface{} {
	return rxxtOptions.GetMap(wrapWithRxxtPrefix(key))
}

// GetSectionFrom returns error while cannot yaml Marshal and Unmarshal
// `cmdr.GetSectionFrom(sectionKeyPath, &holder)` could load all sub-tree nodes from sectionKeyPath and transform them into holder structure, such as:
// ```go
//  type ServerConfig struct {
//    Port int
//    HttpMode int
//    EnableTls bool
//  }
//  var serverConfig = new(ServerConfig)
//  cmdr.GetSectionFrom("server", &serverConfig)
//  assert serverConfig.Port == 7100
// ```
func GetSectionFrom(sectionKeyPath string, holder interface{}) (err error) {
	var b []byte
	fObj := GetMapR(sectionKeyPath)
	if fObj != nil {
		b, err = yaml.Marshal(fObj)
		if err == nil {
			err = yaml.Unmarshal(b, holder)
			// if err == nil {
			// 	logrus.Debugf("configuration section got: %v", configHolder)
			// }
		}
	}
	return
}

// Get an `Option` by key string, eg:
// ```golang
// cmdr.Get("app.logger.level") => 'DEBUG',...
// ```
//
func (s *Options) Get(key string) interface{} {
	defer s.rw.RUnlock()
	s.rw.RLock()
	return s.entries[key]
}

// GetMap an `Option` by key string, it returns a hierarchy map or nil
func (s *Options) GetMap(key string) map[string]interface{} {
	defer s.rw.RUnlock()
	s.rw.RLock()

	return s.getMapNoLock(key)
}

func (s *Options) getMapNoLock(key string) (m map[string]interface{}) {
	a := strings.Split(key, ".")
	if len(a) > 0 {
		m = s.getMap(s.hierarchy, a[0], a[1:]...)
	}
	return
}

func (s *Options) getMap(vp map[string]interface{}, key string, remains ...string) map[string]interface{} {
	if len(remains) > 0 {
		if v, ok := vp[key]; ok {
			if vm, ok := v.(map[string]interface{}); ok {
				return s.getMap(vm, remains[0], remains[1:]...)
			}
		}
		return nil
	}

	if v, ok := vp[key]; ok {
		if vm, ok := v.(map[string]interface{}); ok {
			return vm
		}
		return vp
	}
	return nil
}

// GetBool returns the bool value of an `Option` key.
func (s *Options) GetBool(key string) (ret bool) {
	switch strings.ToLower(s.GetString(key)) {
	case "1", "y", "t", "yes", "true", "ok", "on":
		ret = true
	}
	return
}

// GetInt returns the int64 value of an `Option` key.
func (s *Options) GetInt(key string) (ir int64) {
	if ir64, err := strconv.ParseInt(s.GetString(key), 10, 64); err == nil {
		ir = ir64
	}
	return
}

// GetUint returns the uint64 value of an `Option` key.
func (s *Options) GetUint(key string) (ir uint64) {
	if ir64, err := strconv.ParseUint(s.GetString(key), 10, 64); err == nil {
		ir = ir64
	}
	return
}

// GetFloat32 returns the float32 value of an `Option` key.
func (s *Options) GetFloat32(key string) (ir float32) {
	if ir64, err := strconv.ParseFloat(s.GetString(key), 10); err == nil {
		ir = float32(ir64)
	}
	return
}

// GetFloat64 returns the float64 value of an `Option` key.
func (s *Options) GetFloat64(key string) (ir float64) {
	if ir64, err := strconv.ParseFloat(s.GetString(key), 10); err == nil {
		ir = ir64
	}
	return
}

// GetStringSlice returns the string slice value of an `Option` key.
func (s *Options) GetStringSlice(key string) (ir []string) {
	// envkey := s.envKey(key)
	// if s, ok := os.LookupEnv(envkey); ok {
	// 	ir = strings.Split(s, ",")
	// }

	defer s.rw.RUnlock()
	s.rw.RLock()

	if v, ok := s.entries[key]; ok {
		vvv := reflect.ValueOf(v)
		switch vvv.Kind() {
		case reflect.String:
			ir = strings.Split(v.(string), ",")
		case reflect.Slice:
			if r, ok := v.([]string); ok {
				ir = r
			} else if ri, ok := v.([]int); ok {
				for _, rii := range ri {
					ir = append(ir, strconv.Itoa(rii))
				}
			} else if ri, ok := v.([]byte); ok {
				ir = strings.Split(string(ri), ",")
			} else {
				for i := 0; i < vvv.Len(); i++ {
					ir = append(ir, fmt.Sprintf("%v", vvv.Index(i).Interface()))
				}
			}
		default:
			ir = strings.Split(fmt.Sprintf("%v", v), ",")
		}
	}
	return
}

func stringSliceToIntSlice(in []string) (out []int) {
	for _, ii := range in {
		if i, err := strconv.Atoi(ii); err == nil {
			out = append(out, i)
		}
	}
	return
}

// GetIntSlice returns the int slice value of an `Option` key.
func GetIntSlice(key string) []int {
	return rxxtOptions.GetIntSlice(key)
}

// GetIntSliceP returns the int slice value of an `Option` key.
func GetIntSliceP(prefix, key string) []int {
	return rxxtOptions.GetIntSlice(fmt.Sprintf("%s.%s", prefix, key))
}

// GetIntSliceR returns the int slice value of an `Option` key with [WrapWithRxxtPrefix].
func GetIntSliceR(key string) []int {
	return rxxtOptions.GetIntSlice(wrapWithRxxtPrefix(key))
}

// GetIntSliceRP returns the int slice value of an `Option` key with [WrapWithRxxtPrefix].
func GetIntSliceRP(prefix, key string) []int {
	return rxxtOptions.GetIntSlice(wrapWithRxxtPrefix(fmt.Sprintf("%s.%s", prefix, key)))
}

// GetIntSlice returns the string slice value of an `Option` key.
func (s *Options) GetIntSlice(key string) (ir []int) {
	// envkey := s.envKey(key)
	// if s, ok := os.LookupEnv(envkey); ok {
	// 	ir = stringSliceToIntSlice(strings.Split(s, ","))
	// }

	defer s.rw.RUnlock()
	s.rw.RLock()

	if v, ok := s.entries[key]; ok {
		vvv := reflect.ValueOf(v)
		switch vvv.Kind() {
		case reflect.String:
			ir = stringSliceToIntSlice(strings.Split(v.(string), ","))
		case reflect.Slice:
			if r, ok := v.([]string); ok {
				ir = stringSliceToIntSlice(r)
			} else if ri, ok := v.([]int); ok {
				ir = ri
			} else if ri, ok := v.([]byte); ok {
				xx := strings.Split(string(ri), ",")
				ir = stringSliceToIntSlice(xx)
			} else {
				var xx []string
				for i := 0; i < vvv.Len(); i++ {
					xx = append(xx, fmt.Sprintf("%v", vvv.Index(i).Interface()))
				}
				ir = stringSliceToIntSlice(xx)
			}
		default:
			ir = stringSliceToIntSlice(strings.Split(fmt.Sprintf("%v", v), ","))
		}
	}
	return
}

// GetDuration returns the int slice value of an `Option` key.
func GetDuration(key string) time.Duration {
	return rxxtOptions.GetDuration(key)
}

// GetDurationP returns the int slice value of an `Option` key.
func GetDurationP(prefix, key string) time.Duration {
	return rxxtOptions.GetDuration(fmt.Sprintf("%s.%s", prefix, key))
}

// GetDurationR returns the int slice value of an `Option` key.
func GetDurationR(key string) time.Duration {
	return rxxtOptions.GetDuration(wrapWithRxxtPrefix(key))
}

// GetDurationRP returns the int slice value of an `Option` key.
func GetDurationRP(prefix, key string) time.Duration {
	return rxxtOptions.GetDuration(wrapWithRxxtPrefix(fmt.Sprintf("%s.%s", prefix, key)))
}

// GetDuration returns the time duration value of an `Option` key.
func (s *Options) GetDuration(key string) (ir time.Duration) {
	str := s.GetString(key)
	ir, _ = time.ParseDuration(str)
	return
}

// GetString returns the string value of an `Option` key.
func (s *Options) GetString(key string) (ret string) {
	// envkey := s.envKey(key)
	// if s, ok := os.LookupEnv(envkey); ok {
	// 	ret = s
	// }

	defer s.rw.RUnlock()
	s.rw.RLock()

	if v, ok := s.entries[key]; ok {
		switch reflect.ValueOf(v).Kind() {
		case reflect.String:
			ret = v.(string)
		default:
			if v != nil {
				ret = fmt.Sprintf("%v", v)
			}
		}
	}
	return
}

func (s *Options) buildAutomaticEnv(rootCmd *RootCommand) (err error) {
	// p := strings.Join(EnvPrefix,"_")
	p := strings.Join(RxxtPrefix, ".")
	for key := range s.entries {
		ek := s.envKey(key)
		if v, ok := os.LookupEnv(ek); ok {
			if strings.HasPrefix(key, p) {
				s.Set(key[len(p)+1:], v)
			} else {
				s.Set(key, v)
			}
		}
	}
	return
}

func (s *Options) envKey(key string) (envkey string) {
	key = strings.ReplaceAll(key, ".", "_")
	key = strings.ReplaceAll(key, "-", "_")
	envkey = strings.Join(append(EnvPrefix, strings.ToUpper(key)), "_")
	return
}

// WrapWithRxxtPrefix wrap an key with [RxxtPrefix], for [GetXxx(key)] and [GetXxxP(prefix,key)]
func WrapWithRxxtPrefix(key string) string {
	return wrapWithRxxtPrefix(key)
}

func wrapWithRxxtPrefix(key string) string {
	if len(RxxtPrefix) == 0 {
		return key
	}
	p := strings.Join(RxxtPrefix, ".")
	if len(key) == 0 {
		return p
	}
	return p + "." + key
}

// Set set the value of an `Option` key (with prefix auto-wrap). The key MUST not have an `app` prefix. eg:
// ```golang
// cmdr.Set("logger.level", "DEBUG")
// cmdr.Set("ms.tags.port", 8500)
// ...
// cmdr.Set("debug", true)
// cmdr.GetBool("app.debug") => true
// ```
//
func Set(key string, val interface{}) {
	rxxtOptions.Set(key, val)
}

// SetNx but without prefix auto-wrapped.
// `rxxtPrefix` is a string slice to define the prefix string array, default is ["app"].
// So, cmdr.Set("debug", true) will put an real entry with (`app.debug`, true).
func SetNx(key string, val interface{}) {
	rxxtOptions.SetNx(key, val)
}

// Set set the value of an `Option` key. The key MUST not have an `app` prefix. eg:
// ```golang
// cmdr.Set("debug", true)
// cmdr.GetBool("app.debug") => true
// ```
func (s *Options) Set(key string, val interface{}) {
	k := wrapWithRxxtPrefix(key)
	s.SetNx(k, val)
}

// SetNx but without prefix auto-wrapped.
// `rxxtPrefix` is a string slice to define the prefix string array, default is ["app"].
// So, cmdr.Set("debug", true) will put an real entry with (`app.debug`, true).
func (s *Options) SetNx(key string, val interface{}) {
	defer s.rw.Unlock()
	s.rw.Lock()

	if val == nil {
		if s.getMapNoLock(key) != nil {
			// don't set a branch node to nil if it have children.
			return
		}
	}

	s.entries[key] = val
	a := strings.Split(key, ".")
	mergeMap(s.hierarchy, a[0], et(a, 1, val))
}

func mergeMap(m map[string]interface{}, key string, val interface{}) map[string]interface{} {
	if z, ok := m[key]; ok {
		if zm, ok := z.(map[string]interface{}); ok {
			if vm, ok := val.(map[string]interface{}); ok {
				for k, v := range vm {
					zm = mergeMap(zm, k, v)
				}
				m[key] = zm
			} else {
				m[key] = val
			}
		} else {
			m[key] = val
		}
	} else {
		m[key] = val
	}
	return m
}

func et(keys []string, ix int, val interface{}) interface{} {
	if ix <= len(keys)-1 {
		p := make(map[string]interface{})
		p[keys[ix]] = et(keys, ix+1, val)
		return p
	}
	return val
}

// ResetOptions to reset the exists `Options`, so that you could follow a `LoadConfigFile()` with it.
func ResetOptions() {
	rxxtOptions.Reset()
}

// Reset the exists `Options`, so that you could follow a `LoadConfigFile()` with it.
func (s *Options) Reset() {
	defer s.rw.Unlock()
	s.rw.Lock()

	s.entries = nil
	time.Sleep(100 * time.Millisecond)
	s.entries = make(map[string]interface{})
}

func mx(pre, k string) string {
	if len(pre) == 0 {
		return k
	}
	return pre + "." + k
}

func mxIx(pre string, k interface{}) string {
	if len(pre) == 0 {
		return fmt.Sprintf("%v", k)
	}
	return fmt.Sprintf("%v.%v", pre, k)
}

func (s *Options) loopMapMap(kdot string, m map[string]map[string]interface{}) (err error) {
	for k, v := range m {
		if err = s.loopMap(mx(kdot, k), v); err != nil {
			return
		}
	}
	return
}

func (s *Options) loopMap(kdot string, m map[string]interface{}) (err error) {
	for k, v := range m {
		if vm, ok := v.(map[interface{}]interface{}); ok {
			if err = s.loopIxMap(mx(kdot, k), vm); err != nil {
				return
			}
		} else if vm, ok := v.(map[string]interface{}); ok {
			if err = s.loopMap(mx(kdot, k), vm); err != nil {
				return
			}
		} else {
			s.SetNx(mx(kdot, k), v)
		}
	}
	return
}

func (s *Options) loopIxMap(kdot string, m map[interface{}]interface{}) (err error) {
	for k, v := range m {
		if vm, ok := v.(map[interface{}]interface{}); ok {
			if err = s.loopIxMap(mxIx(kdot, k), vm); err != nil {
				return
			}
			// } else if vm, ok := v.(map[string]interface{}); ok {
			// 	if err = s.loopMap(mxIx(kdot, k), vm); err != nil {
			// 		return
			// 	}
		} else {
			s.SetNx(mxIx(kdot, k), v)
		}
	}
	return
}

// DumpAsString for debugging.
func DumpAsString() (str string) {
	return rxxtOptions.DumpAsString()
}

// SaveAsYaml to Save all config entries as a yaml file
func SaveAsYaml(filename string) (err error) {
	obj := rxxtOptions.GetHierarchyList()

	b, err := yaml.Marshal(obj)
	if err != nil {
		return
	}

	err = ioutil.WriteFile(filename, b, 0644)
	return
}

// SaveAsJSON to Save all config entries as a json file
func SaveAsJSON(filename string) (err error) {
	obj := rxxtOptions.GetHierarchyList()

	b, err := json.Marshal(obj)
	if err != nil {
		return
	}

	err = ioutil.WriteFile(filename, b, 0644)
	return
}

// SaveAsToml to Save all config entries as a toml file
func SaveAsToml(filename string) (err error) {
	obj := rxxtOptions.GetHierarchyList()
	err = SaveObjAsToml(obj, filename)
	return
}

// SaveObjAsToml to Save an object as a toml file
func SaveObjAsToml(obj interface{}, filename string) (err error) {
	f, err := os.Create(filename)
	if err != nil {
		return
	}

	e := toml.NewEncoder(bufio.NewWriter(f))
	if err = e.Encode(obj); err != nil {
		return
	}

	// err = ioutil.WriteFile(filename, b, 0644)
	return
}

// DumpAsString for debugging.
func (s *Options) DumpAsString() (str string) {
	k3 := make([]string, 0)
	for k := range s.entries {
		k3 = append(k3, k)
	}
	sort.Strings(k3)

	for _, k := range k3 {
		str = str + fmt.Sprintf("%-48v => %v\n", k, s.entries[k])
	}
	str += "---------------------------------\n"

	b, err := yaml.Marshal(s.hierarchy)
	if err == nil {
		str += string(b)
	}
	return
}

// GetHierarchyList returns the hierarchy data for dumping
func (s *Options) GetHierarchyList() map[string]interface{} {
	defer s.rw.RUnlock()
	s.rw.RLock()
	return s.hierarchy
}
