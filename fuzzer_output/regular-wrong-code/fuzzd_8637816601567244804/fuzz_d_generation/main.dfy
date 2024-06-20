function abs(x: int): int {
	if (x < 0) then -1 * x else x
}
function safeIndex(x: int, length: int): int 
	requires length > 0
{
	if (x < 0) then 0 else if (x >= length) then x % length else x
}
function safeDivisionInt(x1: int, x2: int): int {
	if (x2 == 0) then x1 else x1 / x2
}
function safeModuloInt(x1: int, x2: int): int {
	if (x2 == 0) then x1 else x1 % x2
}
datatype D0 = DC1(cf1: bool, cf2: map<bool, int>, cf3: bool, cf4: int, cf5: int) | DC2(cf6: int, cf7: bool, cf8: bool, cf9: int) | DC0(cf0: bool) | DC3(cf10: D0)
datatype D1 = DC5(cf12: bool, cf13: C0, cf14: bool) | DC4(cf11: string) | DC6(cf15: D1)
datatype D2 = DC8(cf17: int, cf18: int, cf19: C0) | DC7(cf16: map<bool, C0>)
class GlobalState {
	var f0 : bool
	const f1 : int
	const f2 : string
	const f3 : array<map<bool, int>>
	const f4 : char
	var f5 : int
	var f6 : array<set<int>>
	var f7 : bool
	const f8 : array<int>
	const f9 : array<array<char>>
	var f10 : array<int>
	const f11 : map<bool, int>
	var f12 : bool
	const f13 : bool
	var f14 : set<char>
	const f15 : bool
	constructor (f0 : bool, f1 : int, f2 : string, f3 : array<map<bool, int>>, f4 : char, f5 : int, f6 : array<set<int>>, f7 : bool, f8 : array<int>, f9 : array<array<char>>, f10 : array<int>, f11 : map<bool, int>, f12 : bool, f13 : bool, f14 : set<char>, f15 : bool) {
		this.f0 := f0;
		this.f1 := f1;
		this.f2 := f2;
		this.f3 := f3;
		this.f4 := f4;
		this.f5 := f5;
		this.f6 := f6;
		this.f7 := f7;
		this.f8 := f8;
		this.f9 := f9;
		this.f10 := f10;
		this.f11 := f11;
		this.f12 := f12;
		this.f13 := f13;
		this.f14 := f14;
		this.f15 := f15;
	}
	
}

class C0 {
	const f16 : int
	const f17 : string
	constructor (f16 : int, f17 : string) {
		this.f16 := f16;
		this.f17 := f17;
	}
	
	function fm1(p0: int, p1: map<D0, string>, p2: seq<bool>, globalState: GlobalState): int {
		|(if (!(multiset{false, true} >= multiset{DC1(true, map[true := f16], true, f16, f16).cf1})) then f17 else seq(abs(0x228), i0  => ('p')))|
	}
}

function fm0(p0: bool, p1: char, p2: char, p3: int, globalState: GlobalState): bool {
	true
}
function fm2(p0: int, p1: multiset<int>, globalState: GlobalState): string {
	seq(abs(0x2cf), i0  => ('q'))
}
function fm3(p0: int, p1: bool, p2: bool, globalState: GlobalState): int {
	0x292 - safeModuloInt(0x314, |[0xa1]|)
}
function fm4(p0: int, globalState: GlobalState): seq<bool> {
	[!!!false, false in map[!false := 971]]
}
function fm5(p0: int, globalState: GlobalState): char {
	match DC1(true, map[false := 0x1b3], false, -|(seq(abs(0xc0), i0  => ('t')))|, 848) {
		case DC1(cf1, cf2, cf3, cf4, cf5) => 'p'
		case DC2(cf6, cf7, cf8, cf9) => 'e'
		case DC0(cf0) => 'u'
		case DC3(cf10) => 'i'
	}
}
method m0(p0: bool, globalState: GlobalState) {
	var v0 := "i";
	v0 := v0;
	var v2 := 0x31b;
	var v3: map<int, int> := map[v2 := v2];
	var v4: map<map<int, int>, bool> := map[v3 := true];
	var v5: set<int> := {v2, |v4|};
	var v6: multiset<int> := multiset{v2, v2, fm3(v2, p0, p0, globalState), |v0|};
	v0 := fm2(|(map v1 : int | v1 in v5 :: (v1 + v2) := (p0))|, v6, globalState) + v0;
	var v7: array<bool> := new bool[11] [p0, p0, !p0, p0, p0, p0, p0, p0 <==> p0, p0, p0, p0];
	v7[safeIndex(469, v7.Length)] := p0;
	v7[safeIndex(469, v7.Length)] := !((v2 < v2) || ([p0] == fm4(v2, globalState)));
	var v8: array<int> := new int[24];
	var v9: map<int, bool> := map[v2 := v7[safeIndex(469, v7.Length)]];
	var v10: seq<int> := [|v9|];
	var v11: seq<bool> := [v7[safeIndex(469, v7.Length)]];
	v8[safeIndex(533, v8.Length)] := v10[safeIndex(|v11|, |v10|)];
	var v12 := DC2(|v0|, v7[safeIndex(469, v7.Length)], p0, |v9|);
	v8[safeIndex(533, v8.Length)] := v12.cf6;
	var i0 := 0;
	while (p0)
		decreases 100 - i0
	{
		if (i0 >= 100) {
			break;
		}
		
		i0 := i0 + 1;
		v8[safeIndex(533, v8.Length)] := -((v8[safeIndex(533, v8.Length)] + |v0|) + v2);
		var v13 := 'b';
		v7[safeIndex(469, v7.Length)] := !fm0(p0, fm5(v8[safeIndex(533, v8.Length)], globalState), v13, v2, globalState);
		globalState.f5 := v2;
		var v14: array<map<map<D1, bool>, int>> := new map<map<D1, bool>, int>[21];
		var v15 := DC4(seq(abs(0x34d), i1  => ('o')));
		var v16 := DC6(v15);
		var v17 := DC6(v16);
		var v18 := DC6(v16);
		var v19 := DC6(v16);
		var v20: map<D1, bool> := map[v19 := v7[safeIndex(469, v7.Length)]];
		var v21: map<map<D1, bool>, int> := map[v20 := fm3(0xac, false, p0, globalState)];
		v14[safeIndex(216, v14.Length)] := v21 + v21;
		v14[safeIndex(216, v14.Length)] := v21;
	}
	var v22 := 'b';
	var i2 := 0;
	while (v0 < (v0 + v0)[safeIndex(v8[safeIndex(533, v8.Length)], |(v0 + v0)|) := v22])
		decreases 100 - i2
	{
		if (i2 >= 100) {
			break;
		}
		
		i2 := i2 + 1;
		var v24: array<map<int, bool>> := new map<int, bool>[16] [v9, map[v8[safeIndex(533, v8.Length)] := v7[safeIndex(469, v7.Length)]], v9, v9, v9, v9, map v23 : int | v23 in v3 :: (v23 - v8[safeIndex(533, v8.Length)]) := (true), v9, v9, v9, v9, v9, v9, v9, v9, v9];
		var v25: map<array<map<int, bool>>, seq<bool>> := map[v24 := ([p0, p0])[safeIndex(|v10[safeIndex(|v10|, |v10|) := v8[safeIndex(533, v8.Length)]]|, |[p0, p0]|) := v7[safeIndex(469, v7.Length)]]];
		v25 := v25[v24 := v11];
		var v26: map<int, string> := map[v8[safeIndex(533, v8.Length)] := v0];
		var v27: map<bool, int> := map[p0 := |v26|];
		v27 := v27;
		globalState.f0 := !p0;
		var v28: C0 := new C0(v2, "bxfg");
		var v29 := DC5(v7[safeIndex(469, v7.Length)], v28, p0);
		globalState.f7 := v29.cf14;
	}
}
method Main() {
	var v0 := "sraobfhjd";
	var v1 := 178;
	var v2 := 'q';
	var v3: array<map<bool, int>> := new map<bool, int>[8](i0 => map[DC2(v1, true, true, |[false]|).cf7 := v1]);
	var v4: array<set<int>> := new set<int>[19](i1 => {v1});
	var v5: array<int> := new int[23](i2 => safeModuloInt(i2, 855));
	var v6: array<array<char>> := new array<char>[17];
	var v7 := false;
	var v8: set<char> := {v2, v2, 'u', v2};
	var globalState := new GlobalState(false, 0xed, v0[safeIndex(v1, |v0|) := v2], v3, 'w', -0x2e, v4, false, v5, v6, v5, map[v7 := -v1], false, true, v8, true);
	v2 := v2;
	var v9: multiset<bool> := multiset{v7};
	for i3 := v1 to if (v7 in v9) then v9[v7] else v1 {
		var v10 := DC2(0x3e6, v7, v7, |v0|);
		if (if (v7) then i3 >= v1 else (v10.(cf9 := |map[i3 := |v0|]|)).cf8) {
			var v11: array<bool> := new bool[3](i4 => !!!true);
			v11[safeIndex(829, v11.Length)] := |v9| > i3;
			v11[safeIndex(829, v11.Length)] := v7 ==> v7;
			var v12: seq<array<int>> := [v5];
			var v13: map<bool, bool> := map[v11[safeIndex(829, v11.Length)] := v7];
			var v14: array<array<int>> := new array<int>[5] [v5, v5, v5, v12[safeIndex(|v13|, |v12|)], v5];
			v14[safeIndex(694, v14.Length)] := v5;
			v14[safeIndex(694, v14.Length)] := v5;
			var v15: map<array<int>, bool> := map[v5 := v7 <== true];
			var v16: map<bool, array<int>> := map[v7 := v5];
			v15 := v15[if (v11[safeIndex(829, v11.Length)] in v16) then v16[v11[safeIndex(829, v11.Length)]] else v5 := i3 > i3];
			globalState.f5 := if (v11[safeIndex(829, v11.Length)]) then v1 else 823;
			globalState.f5 := v1 - i3;
		} else {
			m0(v7, globalState);
			globalState.f12 := fm0(v7 || v7, v2, v2, v1, globalState);
			v2 := v2;
			v1 := |v0|;
			globalState.f0 := true;
		}
		
		var v17: map<bool, bool> := map[true := v7];
		var v18: map<map<bool, bool>, string> := map[v17 := v0];
		globalState.f5 := |((if (v17 in v18) then v18[v17] else v0) + (if (v7) then v0 else v0))|;
		var v19 := new C0(i3, v0);
		globalState.f0 := if (v7) then v7 else v7;
	}
	m0(v7, globalState);
	for i5 := -(v1 + v1) to v1 {
		globalState.f10, v2, globalState.f12, globalState.f12 := v5, v2, v7 ==> v7, v7;
		var v20 := new C0(i5, v0 + v0);
		var v21: seq<int> := [v20.f16];
		v21 := if (v7) then seq(abs(346), i6  => (v20.f16)) else seq(abs(0x321), i7  => (i5));
		var v22 := DC0(v7);
		var v23: map<D0, string> := map[v22 := v20.f17];
		var v24 := DC4(v20.f17);
		v23 := v23[DC0(!v7) := v24.cf11 + v0];
	}
	var v25: multiset<int> := multiset{v1};
	globalState.f7 := fm2(v1, v25, globalState) <= (seq(abs(-49), i8  => (v2)));
	var v26: set<bool> := {v7, v7, v7};
	var v27: map<set<bool>, bool> := map[v26 := fm0(v7, v2, v2, -v1, globalState)];
	for i9 := |v27| to 103 {
		var v28: array<D1> := new D1[11];
		var v29: C0 := new C0(i9, "hrcnf");
		var v30 := DC5(v7, v29, v7);
		v28[safeIndex(761, v28.Length)] := v30.(cf12 := v7, cf14 := v7);
		v28[safeIndex(761, v28.Length)] := v30;
		var v31: set<C0> := {v29};
		globalState.f5, globalState.f5 := -v29.f16, |(v31 + v31)|;
		var v32 := DC4(v29.f17);
		var v33 := DC6(v32);
		v33 := DC6(v32);
		m0(v7, globalState);
	}
	var v34: map<bool, bool> := map[false := v7];
	v34 := v34[false := v7];
	var v35: C0 := new C0(v1, v0);
	globalState.f5 := |(map[v1 := v35] + map[v1 := v35])|;
	var v36: map<int, string> := map[v35.f16 := (seq(abs(816), i10  => (v2))) + "fal"];
	v36 := v36[-0x11a := v35.f17];
	var v37 := new C0(safeModuloInt(v35.f16, v35.f16), (seq(abs(41), i11  => (v2))) + v35.f17);
	m0(v37.f16 != v35.f16, globalState);
	var v38: map<char, int> := map[v2 := |v0|];
	v38 := (map v39 : char | v39 in [v2] :: (v39) := (|{v37.f16, v35.f16, v37.f16, |v34|, v1}|)) + v38;
	v7 := v7;
	globalState.f5 := v37.f16;
	m0(v7, globalState);
	var v40: map<bool, C0> := map[v7 := v35];
	globalState.f5 := (v37.f16 * |DC7(v40).cf16|) + 0xff;
	print v0, "\n";
	print v1, "\n";
	print v2, "\n";
	print v3[0] == map[true := 178], "\n";
	print v3[1] == map[true := 178], "\n";
	print v3[2] == map[true := 178], "\n";
	print v3[3] == map[true := 178], "\n";
	print v3[4] == map[true := 178], "\n";
	print v3[5] == map[true := 178], "\n";
	print v3[6] == map[true := 178], "\n";
	print v3[7] == map[true := 178], "\n";
	print v4[0] == {178}, "\n";
	print v4[1] == {178}, "\n";
	print v4[2] == {178}, "\n";
	print v4[3] == {178}, "\n";
	print v4[4] == {178}, "\n";
	print v4[5] == {178}, "\n";
	print v4[6] == {178}, "\n";
	print v4[7] == {178}, "\n";
	print v4[8] == {178}, "\n";
	print v4[9] == {178}, "\n";
	print v4[10] == {178}, "\n";
	print v4[11] == {178}, "\n";
	print v4[12] == {178}, "\n";
	print v4[13] == {178}, "\n";
	print v4[14] == {178}, "\n";
	print v4[15] == {178}, "\n";
	print v4[16] == {178}, "\n";
	print v4[17] == {178}, "\n";
	print v4[18] == {178}, "\n";
	print v5[0], "\n";
	print v5[1], "\n";
	print v5[2], "\n";
	print v5[3], "\n";
	print v5[4], "\n";
	print v5[5], "\n";
	print v5[6], "\n";
	print v5[7], "\n";
	print v5[8], "\n";
	print v5[9], "\n";
	print v5[10], "\n";
	print v5[11], "\n";
	print v5[12], "\n";
	print v5[13], "\n";
	print v5[14], "\n";
	print v5[15], "\n";
	print v5[16], "\n";
	print v5[17], "\n";
	print v5[18], "\n";
	print v5[19], "\n";
	print v5[20], "\n";
	print v5[21], "\n";
	print v5[22], "\n";
	print v7, "\n";
	print v8 == {'q', 'u'}, "\n";
	print globalState.f0, "\n";
	print globalState.f1, "\n";
	print globalState.f2, "\n";
	print globalState.f3[0] == map[true := 178], "\n";
	print globalState.f3[1] == map[true := 178], "\n";
	print globalState.f3[2] == map[true := 178], "\n";
	print globalState.f3[3] == map[true := 178], "\n";
	print globalState.f3[4] == map[true := 178], "\n";
	print globalState.f3[5] == map[true := 178], "\n";
	print globalState.f3[6] == map[true := 178], "\n";
	print globalState.f3[7] == map[true := 178], "\n";
	print globalState.f4, "\n";
	print globalState.f5, "\n";
	print globalState.f6[0] == {178}, "\n";
	print globalState.f6[1] == {178}, "\n";
	print globalState.f6[2] == {178}, "\n";
	print globalState.f6[3] == {178}, "\n";
	print globalState.f6[4] == {178}, "\n";
	print globalState.f6[5] == {178}, "\n";
	print globalState.f6[6] == {178}, "\n";
	print globalState.f6[7] == {178}, "\n";
	print globalState.f6[8] == {178}, "\n";
	print globalState.f6[9] == {178}, "\n";
	print globalState.f6[10] == {178}, "\n";
	print globalState.f6[11] == {178}, "\n";
	print globalState.f6[12] == {178}, "\n";
	print globalState.f6[13] == {178}, "\n";
	print globalState.f6[14] == {178}, "\n";
	print globalState.f6[15] == {178}, "\n";
	print globalState.f6[16] == {178}, "\n";
	print globalState.f6[17] == {178}, "\n";
	print globalState.f6[18] == {178}, "\n";
	print globalState.f7, "\n";
	print globalState.f8[0], "\n";
	print globalState.f8[1], "\n";
	print globalState.f8[2], "\n";
	print globalState.f8[3], "\n";
	print globalState.f8[4], "\n";
	print globalState.f8[5], "\n";
	print globalState.f8[6], "\n";
	print globalState.f8[7], "\n";
	print globalState.f8[8], "\n";
	print globalState.f8[9], "\n";
	print globalState.f8[10], "\n";
	print globalState.f8[11], "\n";
	print globalState.f8[12], "\n";
	print globalState.f8[13], "\n";
	print globalState.f8[14], "\n";
	print globalState.f8[15], "\n";
	print globalState.f8[16], "\n";
	print globalState.f8[17], "\n";
	print globalState.f8[18], "\n";
	print globalState.f8[19], "\n";
	print globalState.f8[20], "\n";
	print globalState.f8[21], "\n";
	print globalState.f8[22], "\n";
	print globalState.f10[0], "\n";
	print globalState.f10[1], "\n";
	print globalState.f10[2], "\n";
	print globalState.f10[3], "\n";
	print globalState.f10[4], "\n";
	print globalState.f10[5], "\n";
	print globalState.f10[6], "\n";
	print globalState.f10[7], "\n";
	print globalState.f10[8], "\n";
	print globalState.f10[9], "\n";
	print globalState.f10[10], "\n";
	print globalState.f10[11], "\n";
	print globalState.f10[12], "\n";
	print globalState.f10[13], "\n";
	print globalState.f10[14], "\n";
	print globalState.f10[15], "\n";
	print globalState.f10[16], "\n";
	print globalState.f10[17], "\n";
	print globalState.f10[18], "\n";
	print globalState.f10[19], "\n";
	print globalState.f10[20], "\n";
	print globalState.f10[21], "\n";
	print globalState.f10[22], "\n";
	print globalState.f11 == map[false := -178], "\n";
	print globalState.f12, "\n";
	print globalState.f13, "\n";
	print globalState.f14 == {'q', 'u'}, "\n";
	print globalState.f15, "\n";
	print v9 == multiset{false}, "\n";
	print v25 == multiset{178}, "\n";
	print v26 == {false}, "\n";
	print v27 == map[{false} := true], "\n";
	print v34 == map[false := false], "\n";
	print v35.f16, "\n";
	print v35.f17, "\n";
	print v36 == map[178 := "qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqfal", -282 := "sraobfhjd"], "\n";
	print v37.f16, "\n";
	print v37.f17, "\n";
	print v38 == map['q' := 9], "\n";
	print |v40|, "\n";
}