package main

import (
    "crypto/sha256"
    "fmt"
	"encoding/base64"
    "github.com/Nik-U/pbc"
)


	var str_params = `type a
		q 7952436261441942116724005628190220211674820351038725044719754524973612317658021198716510489828805006470399946340981009347486324095778164329267829132520191
		h 10882555391405823409412640212961128637981856544833210764164836988830046619205446557430811225581697238537472
		r 730750818665452757176057050065048642452048576511
		exp2 159
		exp1 110
		sign1 1
		sign0 -1
	`
	//生成元P 由CompressedBytes导出的数据
	var str_P = "BjH+pEP5ZYgTzmhn2zAlf5nHt6T+F+uJX6BrTePi2EoVypoZYWyjHS/J/eqW3zODZ4Q0mFWb5yMs9nHoyuhqdQE="
	//系统主密钥s 由Bytes导出
	var str_s = "VbZDgs8v+Tm5g7cK16eZVdDTGV4="
func main() {
    // The authority generates system parameters
    // params := pbc.GenerateA(160, 512)
	fmt.Println("\n\n一种改进的基于身份的签名方案\n\n")

	params,_ := pbc.NewParamsFromString(str_params)
	fmt.Println("\n\n系统参数params:\n",params)
    pairing := params.NewPairing()
	

	P_bytes,_ := base64.StdEncoding.DecodeString(str_P)
	P := pairing.NewG1().SetCompressedBytes(P_bytes)
	fmt.Println("\n\n生成元P:\n",str_P)

	s_bytes,_ := base64.StdEncoding.DecodeString(str_s)
	s := pairing.NewZr().SetBytes(s_bytes)
	fmt.Println("\n\n系统主密钥s:\n", str_s)

	Ppub := pairing.NewG1().MulZn(P,s)
	Ppub_bytes := Ppub.Bytes()
	str_Ppub := base64.StdEncoding.EncodeToString(Ppub_bytes)
	fmt.Println("\n\n系统公钥Ppub:\n",str_Ppub)
	ID := "ID"
	Qid := pairing.NewG1().SetFromStringHash(ID,sha256.New())
	str_Qid:= base64.StdEncoding.EncodeToString(Qid.Bytes())
	fmt.Println("\n\nstep 4 ---- ID公钥Qid:\n",str_Qid)
	Did := pairing.NewG1().MulZn(Qid,s)
	str_Did:= base64.StdEncoding.EncodeToString(Did.Bytes())
	fmt.Println("\n\nstep 5 ---- ID私钥Did:\n",str_Did)
	
	k := pairing.NewZr().Rand()
	str_k:= base64.StdEncoding.EncodeToString(k.Bytes())
	fmt.Println("\n\nstep 6 ---- 签名随机参数k:\n",str_k)
	j1 := pairing.NewGT().Pair(Qid,Ppub)
	R := pairing.NewGT().PowZn(j1,k)
	str_R := base64.StdEncoding.EncodeToString(R.Bytes())
	fmt.Println("\n\nstep 7 ---- 签名的一部分R:\n",str_R)
	
	message := "origin message"
	m := pairing.NewZr().SetFromStringHash(message, sha256.New())
	fmt.Println("\n\nstep 8 ---- message:\n",message)
	mpz1 := R.X()
	h1 := pairing.NewZr().MulBig(m, mpz1)
	str_h1 := base64.StdEncoding.EncodeToString(h1.Bytes())
	fmt.Println("\n\nstep 9 ---- h1:\n",str_h1)
	j2 := pairing.NewZr().Add(k,h1)
	V := pairing.NewG1().MulZn(Did, j2)
	str_V := base64.StdEncoding.EncodeToString(V.Bytes())
	fmt.Println("\n\n签名为sigma:(R,V)\n")
	fmt.Println("\nR:",str_R)
	fmt.Println("\nV:",str_V)


	//验证阶段
	c1 := pairing.NewGT().Pair(V,P)
	str_c1 := base64.StdEncoding.EncodeToString(c1.Bytes())
	fmt.Println("\n\ne(V,P)=\n",str_c1)
	c2 := pairing.NewGT().Pair(Qid,Ppub)
	c3 := pairing.NewGT().PowZn(c2,h1)
	c4 := pairing.NewGT().Mul(R, c3)
	str_c4 := base64.StdEncoding.EncodeToString(c4.Bytes())
	
	fmt.Println("\n\nc4:\n",str_c4)
	c5 := pairing.NewGT().Pow2Zn(c2,k,c2,h1)
	str_c5 := base64.StdEncoding.EncodeToString(c5.Bytes())
	fmt.Println("\n\nc5:\n",str_c5)

}
