// Copyright © 2020 Hedzr Yeh.

package tool

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"gopkg.in/hedzr/errors.v2"
	"io"
	"io/ioutil"
	"math"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

// ParseComplex converts a string to complex number.
//
// Examples:
//
//    c1 := cmdr.ParseComplex("3-4i")
//    c2 := cmdr.ParseComplex("3.13+4.79i")
func ParseComplex(s string) (v complex128) {
	return a2complexShort(s)
}

// ParseComplexX converts a string to complex number.
// If the string is not valid complex format, return err not nil.
//
// Examples:
//
//    c1 := cmdr.ParseComplex("3-4i")
//    c2 := cmdr.ParseComplex("3.13+4.79i")
func ParseComplexX(s string) (v complex128, err error) {
	return a2complex(s)
}

func a2complexShort(s string) (v complex128) {
	v, _ = a2complex(s)
	return
}

func a2complex(s string) (v complex128, err error) {
	s = strings.TrimSpace(strings.TrimRightFunc(strings.TrimLeftFunc(s, func(r rune) bool {
		return r == '('
	}), func(r rune) bool {
		return r == ')'
	}))

	if i := strings.IndexAny(s, "+-"); i >= 0 {
		rr, ii := s[0:i], s[i:]
		if j := strings.Index(ii, "i"); j >= 0 {
			var ff, fi float64
			ff, err = strconv.ParseFloat(strings.TrimSpace(rr), 64)
			if err != nil {
				return
			}
			fi, err = strconv.ParseFloat(strings.TrimSpace(ii[0:j]), 64)
			if err != nil {
				return
			}

			v = complex(ff, fi)
			return
		}
		err = errors.New("for a complex number, the imaginary part should end with 'i', such as '3+4i'")
		return

		// err = errors.New("not valid complex number.")
	}

	var ff float64
	ff, err = strconv.ParseFloat(strings.TrimSpace(s), 64)
	if err != nil {
		return
	}
	v = complex(ff, 0)
	return
}

//
// external
//

// Launch executes a command setting both standard input, output and error.
func Launch(cmd string, args ...string) (err error) {
	c := exec.Command(cmd, args...)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	err = c.Run()

	if err != nil {
		if _, isExitError := err.(*exec.ExitError); isExitError {
			err = nil
		}
	}
	return
}

// // LaunchSudo executes a command under "sudo".
// func LaunchSudo(cmd string, args ...string) error {
// 	return Launch("sudo", append([]string{cmd}, args...)...)
// }

//
// editor
//

// func getEditor() (string, error) {
// 	if GetEditor != nil {
// 		return GetEditor()
// 	}
// 	return exec.LookPath(DefaultEditor)
// }

func randomFilename() (fn string) {
	buf := make([]byte, 16)
	fn = os.Getenv("HOME") + ".CMDR_EDIT_FILE"
	if _, err := rand.Read(buf); err == nil {
		fn = fmt.Sprintf("%v/.CMDR_%x", os.Getenv("HOME"), buf)
	}
	return
}

// const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// RandomStringPure generate a random string with length specified.
func RandomStringPure(length int) (result string) {
	buf := make([]byte, length)
	if _, err := rand.Read(buf); err == nil {
		result = string(buf)
	}
	return
	// source:=rand.NewSource(time.Now().UnixNano())
	// b := make([]byte, length)
	// for i := range b {
	// 	b[i] = charset[source.Int63()%int64(len(charset))]
	// }
	// return string(b)
}

// LaunchEditor launches the specified editor
func LaunchEditor(editor string) (content []byte, err error) {
	return launchEditorWith(editor, randomFilename())
}

// LaunchEditorWith launches the specified editor with a filename
func LaunchEditorWith(editor string, filename string) (content []byte, err error) {
	return launchEditorWith(editor, filename)
}

func launchEditorWith(editor, filename string) (content []byte, err error) {
	cmd := exec.Command(editor, filename)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()

	if err != nil {
		if _, isExitError := err.(*exec.ExitError); !isExitError {
			return
		}
	}

	content, err = ioutil.ReadFile(filename)
	if err != nil {
		return []byte{}, nil
	}
	return
}

// Soundex returns the english word's soundex value, such as: 'tags' => 't322'
func Soundex(s string) (snd4 string) {
	return soundex(s)
}

func soundex(s string) (snd4 string) {
	// if len(s) == 0 {
	// 	return
	// }

	var src, tgt []rune
	src = []rune(s)

	i := 0
	for ; i < len(src); i++ {
		if !(src[i] == '-' || src[i] == '~' || src[i] == '+') {
			// first char
			tgt = append(tgt, src[i])
			break
		}
	}

	for ; i < len(src); i++ {
		ch := src[i]
		switch ch {
		case 'a', 'e', 'i', 'o', 'u', 'y', 'h', 'w': // do nothing to remove it
		case 'b', 'f', 'p', 'v':
			tgt = append(tgt, '1')
		case 'c', 'g', 'j', 'k', 'q', 's', 'x', 'z':
			tgt = append(tgt, '2')
		case 'd', 't':
			tgt = append(tgt, '3')
		case 'l':
			tgt = append(tgt, '4')
		case 'm', 'n':
			tgt = append(tgt, '5')
		case 'r':
			tgt = append(tgt, '6')
		}
	}

	snd4 = string(tgt)
	return
}

// StringMetricFactor for JaroWinklerDistance algorithm
const StringMetricFactor = 100000000000

type (
	// StringDistance is an interface for string metric.
	// A string metric is a metric that measures distance between two strings.
	// In most case, it means that the edit distance about those two strings.
	// This is saying, it is how many times are needed while you were
	// modifying string to another one, note that inserting, deleting,
	// substing one character means once.
	StringDistance interface {
		Calc(s1, s2 string, opts ...DistanceOption) (distance int)
	}

	// DistanceOption is a functional options prototype
	DistanceOption func(StringDistance)
)

// JaroWinklerDistance returns an calculator for two strings distance metric, with Jaro-Winkler algorithm.
func JaroWinklerDistance(opts ...DistanceOption) StringDistance {
	x := &jaroWinklerDistance{threshold: 0.7, factor: StringMetricFactor}
	for _, c := range opts {
		c(x)
	}
	return x
}

// JWWithThreshold sets the threshold for Jaro-Winkler algorithm.
func JWWithThreshold(threshold float64) DistanceOption {
	return func(distance StringDistance) {
		if v, ok := distance.(*jaroWinklerDistance); ok {
			v.threshold = threshold
		}
	}
}

type jaroWinklerDistance struct {
	threshold float64
	factor    float64

	matches        int
	maxLength      int
	transpositions int // transpositions is a double number here
	prefix         int

	distance float64
}

func (s *jaroWinklerDistance) Calc(src1, src2 string, opts ...DistanceOption) (distance int) {
	s1, s2 := []rune(src1), []rune(src2)
	lenMax, lenMin := len(s1), len(s2)

	var sMax, sMin []rune
	if lenMax > lenMin {
		sMax, sMin = s1, s2
	} else {
		sMax, sMin = s2, s1
		lenMax, lenMin = lenMin, lenMax
	}
	s.maxLength = lenMax

	iMatchIndexes, matchFlags := s.match(sMax, sMin, lenMax, lenMin)
	s.findTranspositions(sMax, sMin, lenMax, lenMin, iMatchIndexes, matchFlags)

	// println("  matches, transpositions, prefix: ", s.matches, s.transpositions, s.prefix)

	if s.matches == 0 {
		s.distance = 0
		return 0
	}

	m := float64(s.matches)
	jaroDistance := m/float64(lenMax) + m/float64(lenMin)
	jaroDistance += (m - float64(s.transpositions)/2) / m
	jaroDistance /= 3

	var jw float64
	if jaroDistance < s.threshold {
		jw = jaroDistance
	} else {
		jw = jaroDistance + math.Min(0.1, 1/float64(s.maxLength))*float64(s.prefix)*(1-jaroDistance)
	}

	// println("  jaro, jw: ", jaroDistance, jw)

	s.distance = jw * s.factor
	distance = int(math.Round(s.distance))
	return
}

func (s *jaroWinklerDistance) match(sMax, sMin []rune, lenMax, lenMin int) (iMatchIndexes []int, matchFlags []bool) {
	iRange := Max(lenMax/2-1, 0)
	iMatchIndexes = make([]int, lenMin)
	for i := 0; i < lenMin; i++ {
		iMatchIndexes[i] = -1
	}

	s.prefix, s.matches = 0, 0
	for mi := 0; mi < len(sMin); mi++ {
		if sMax[mi] == sMin[mi] {
			s.prefix++
		} else {
			break
		}
	}
	s.matches = s.prefix

	matchFlags = make([]bool, lenMax)

	for mi := s.prefix; mi < lenMin; mi++ {
		c1 := sMin[mi]
		xi, xn := Max(mi-iRange, s.prefix), lenMax // min(mi+iRange-1, lenMax)
		for ; xi < xn; xi++ {
			if !matchFlags[xi] && c1 == sMax[xi] {
				iMatchIndexes[mi] = xi
				matchFlags[xi] = true
				s.matches++
				break
			}
		}
	}
	return
}

func (s *jaroWinklerDistance) findTranspositions(sMax, sMin []rune, lenMax, lenMin int, iMatchIndexes []int, matchFlags []bool) {
	ms1, ms2 := make([]rune, s.matches), make([]rune, s.matches)
	for i, si := 0, 0; i < lenMin; i++ {
		if iMatchIndexes[i] != -1 {
			ms1[si] = sMin[i]
			si++
		}
	}
	for i, si := 0, 0; i < lenMax; i++ {
		if matchFlags[i] {
			ms2[si] = sMax[i]
			si++
		}
	}
	// fmt.Printf("iMatchIndexes, s1, s2: %v, %v, %v\n", iMatchIndexes, string(sMax), string(sMin))
	// println("     ms1, ms2: ", string(ms1), string(ms2))

	s.transpositions = 0
	for mi := 0; mi < len(ms1); mi++ {
		if ms1[mi] != ms2[mi] {
			s.transpositions++
		}
	}
}

// Max return the larger one of int
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Min return the less one of int
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// StripPrefix strips the prefix 'p' from a string 's'
func StripPrefix(s, p string) string {
	return stripPrefix(s, p)
}

func stripPrefix(s, p string) string {
	if strings.HasPrefix(s, p) {
		return s[len(p):]
	}
	return s
}

// StripOrderPrefix strips the prefix string fragment for sorting order.
// see also: Command.Group, Flag.Group, ...
// An order prefix is a dotted string with multiple alphabet and digit. Such as:
// "zzzz.", "0001.", "700.", "A1." ...
func StripOrderPrefix(s string) string {
	if xre.MatchString(s) {
		s = s[strings.Index(s, ".")+1:]
	}
	return s
}

// HasOrderPrefix tests whether an order prefix is present or not.
// An order prefix is a dotted string with multiple alphabet and digit. Such as:
// "zzzz.", "0001.", "700.", "A1." ...
func HasOrderPrefix(s string) bool {
	return xre.MatchString(s)
}

var (
	xre = regexp.MustCompile(`^[0-9A-Za-z]+\.(.+)$`)
)

// IsDigitHeavy tests if the whole string is digit
func IsDigitHeavy(s string) bool {
	m, _ := regexp.MatchString("^\\d+$", s)
	// if err != nil {
	// 	return false
	// }
	return m
}

// PressEnterToContinue lets program pause and wait for user's ENTER key press in console/terminal
func PressEnterToContinue(in io.Reader, msg ...string) (input string) {
	if len(msg) > 0 && len(msg[0]) > 0 {
		fmt.Print(msg[0])
	} else {
		fmt.Print("Press 'Enter' to continue...")
	}
	b, _ := bufio.NewReader(in).ReadBytes('\n')
	return strings.TrimRight(string(b), "\n")
}

// PressAnyKeyToContinue lets program pause and wait for user's ANY key press in console/terminal
func PressAnyKeyToContinue(in io.Reader, msg ...string) (input string) {
	if len(msg) > 0 && len(msg[0]) > 0 {
		fmt.Print(msg[0])
	} else {
		fmt.Print("Press any key to continue...")
	}
	_, _ = fmt.Fscanf(in, "%s", &input)
	return
}

// SavedOsArgs is a copy of os.Args, just for testing
var SavedOsArgs []string

func init() {
	if SavedOsArgs == nil {
		// bug: can't copt slice to slice: _ = StandardCopier.Copy(&SavedOsArgs, &os.Args)
		for _, s := range os.Args {
			SavedOsArgs = append(SavedOsArgs, s)
		}
	}
}
