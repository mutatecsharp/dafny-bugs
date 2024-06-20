datatype D0 = DC0(cf0: int)
datatype D1 = DC2(cf2: bool) | DC1(cf1: array<bool>) | DC3(cf3: D1)
datatype D2 = DC5(cf5: bool, cf6: int, cf7: array<bool>, cf8: multiset<int>) | DC6(cf9: int, cf10: bool, cf11: bool, cf12: int) | DC4(cf4: array<D1>) | DC7(cf13: D2)
datatype D3 = DC8(cf14: multiset<C0>)
datatype D4 = DC10(cf16: array<int>) | DC9(cf15: C0) | DC11(cf17: D4)
datatype D5 = DC13(cf19: int, cf20: int, cf21: bool) | DC12(cf18: string)
datatype D6 = DC15(cf23: int, cf24: int, cf25: seq<int>, cf26: int, cf27: set<int>) | DC16(cf28: int, cf29: bool, cf30: int, cf31: int, cf32: bool) | DC17(cf33: char, cf34: bool, cf35: bool, cf36: char) | DC14(cf22: C1)
datatype D7 = DC19(cf38: set<array<char>>, cf39: bool, cf40: bool) | DC20(cf41: set<array<D4>>, cf42: D5, cf43: int) | DC18(cf37: set<map<bool, D4>>) | DC21(cf44: D7)
datatype D8 = DC23(cf46: int, cf47: array<int>, cf48: int) | DC22(cf45: C2)
datatype D9 = DC25(cf50: int) | DC26(cf51: bool, cf52: set<bool>, cf53: array<bool>, cf54: array<seq<int>>) | DC27(cf55: int) | DC24(cf49: array<string>)
datatype D10 = DC29 | DC28(cf56: map<map<int, bool>, int>) | DC30(cf57: D10)
datatype D11 = DC32(cf59: int) | DC31(cf58: multiset<bool>) | DC33(cf60: D11)
datatype D12 = DC34(cf61: map<bool, char>)
class GlobalState {
	var f0 : bool
	const f1 : seq<bool>
	var f2 : int
	var f3 : bool
	const f4 : bool
	const f5 : char
	const f6 : int
	const f7 : array<bool>
	var f8 : multiset<int>
	var f9 : int
	var f10 : bool
	var f11 : array<bool>
	const f12 : map<bool, int>
	const f13 : string
	const f14 : bool
	const f15 : map<bool, bool>
	const f16 : bool
	constructor (f0 : bool, f1 : seq<bool>, f2 : int, f3 : bool, f4 : bool, f5 : char, f6 : int, f7 : array<bool>, f8 : multiset<int>, f9 : int, f10 : bool, f11 : array<bool>, f12 : map<bool, int>, f13 : string, f14 : bool, f15 : map<bool, bool>, f16 : bool) {
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
		this.f16 := f16;
	}
	
}

function fm0(globalState: GlobalState): int {
	|map[-120 := false]| / (-0x160 + |(seq(873, i0  => ('n')))|)
}
function fm1(p0: bool, p1: string, globalState: GlobalState): bool {
	'k' in (if (false) then "ewqpfrw" else "pfb")
}
function fm5(p0: int, p1: int, globalState: GlobalState): D1 {
	DC2(!(|map[false := false]| != |[seq(175, i0  => ('c')), "ol"]|))
}
function fm6(p0: bool, globalState: GlobalState): string {
	"k" + "pgy"
}
function fm7(p0: char, p1: bool, globalState: GlobalState): char {
	'k'
}
function fm8(p0: int, p1: int, globalState: GlobalState): set<int> {
	{0x230, -|"mng"| / 987, if (true) then |map[false := false]| else -0x114, 0xcb}
}
function fm9(p0: int, p1: bool, p2: bool, p3: int, globalState: GlobalState): multiset<bool> {
	multiset{true in {!false}, true <== true, true}
}
function fm10(globalState: GlobalState): map<bool, seq<bool>> {
	if (true) then map[false := [false, true]] + map[false := [true, true, true, false, false]] else map[!true := [false]]
}
function fm11(p0: int, p1: bool, p2: int, globalState: GlobalState): seq<int> {
	(if (false) then seq(0x27c, i0  => (|"fqfptov"|)) else [-151]) + ([|map[0x37a := false]|, 0x1f9] + [0x300])
}
function fm12(p0: int, p1: seq<set<int>>, p2: int, p3: int, globalState: GlobalState): D5 {
	match DC13(0x36f, |map[true := true]|, !false) {
		case DC13(cf19, cf20, cf21) => DC12("msdl")
		case DC12(cf18) => if (true) then DC12(cf18) else DC12(cf18)
	}
}
function fm13(globalState: GlobalState): map<map<int, bool>, int> {
	match DC12(seq(832, i0  => ('i'))) {
		case DC13(cf19, cf20, cf21) => map[map[cf20 := cf21] := cf19] + (map v0 : map<int, bool> | v0 in map[map[|{|{cf21}|}| := cf21] := cf21] :: (v0) := (cf20))
		case DC12(cf18) => (map v1 : map<int, bool> | v1 in map[map v2 : int | (0x359 <= v2) && (v2 < -0x15) :: (v2 * 0x122) := (false) := true] :: (v1) := (542)) + map[map[|map['x' := 0x147]| := !true] := 610]
	}
}
function fm14(p0: bool, p1: char, p2: bool, globalState: GlobalState): set<bool> {
	{true, false, |[0x142, |"t"|, |map[true := |(seq(0x22e, i0  => ('f')))|]|]| > |{0x4c, 0x70}|}
}
function fm15(p0: bool, p1: map<bool, int>, p2: bool, p3: bool, globalState: GlobalState): map<seq<int>, multiset<bool>> {
	map[[|map[false := true]|, |multiset([|[508, |map[false := 'o']|]|])|] := multiset([false, true, true])] + map[seq(644, i0  => (-|[false]|)) := multiset([false])]
}
function fm16(p0: int, p1: bool, p2: char, p3: int, globalState: GlobalState): map<bool, bool> {
	map[true := !!!false] + (map[true := false] + map[true := false])
}
function fm17(p0: int, p1: seq<D2>, globalState: GlobalState): D5 {
	DC13(|"qkgiwheqr"|, 0x12f * 832, true ==> true)
}
function fm18(p0: int, globalState: GlobalState): map<bool, char> {
	map[{map[true := !true], map[true := true]} != {map[true := false], map[false := !false]} := 'u']
}
method m0(p0: bool, p1: map<int, bool>, globalState: GlobalState) returns (r0: int, r1: int) {
	var v0: set<bool> := {p0};
	var v1: array<bool> := new bool[21];
	var v2: array<seq<int>> := new seq<int>[16](i0 => [468]);
	var v3: map<bool, bool> := map[p0 := p0];
	var v4 := "knce";
	var v5: map<bool, string> := map[p0 := v4];
	var v6: array<bool> := new bool[22] [false, DC26(p0, v0, v1, v2).cf51, p0, p0, if (fm1(true, seq(0xb0, i1  => ('d')), globalState) in v3) then v3[fm1(true, seq(0xb0, i1  => ('d')), globalState)] else !p0, p0, p0, p0, p0, p0, p0, fm1(p0, if (p0 in v5) then v5[p0] else v4, globalState), p0, p0, p0, false, p0, p0, p0, p0, p0, p0];
	var v7: seq<array<bool>> := [v6, v6];
	var v8 := 0x0;
	var v9: map<bool, seq<array<bool>>> := map[!(if (p0) then p0 else p0) := v7[v8 := v6]];
	v9 := v9[!p0 := v7];
	var v10: seq<int> := [v8];
	var v11 := DC13(v8, -v8, p0);
	globalState.f2 := |(match DC6(|v10|, v11.cf21, p0, v8) {
		case DC5(cf5, cf6, cf7, cf8) => {map[cf5 := 'a']}
		case DC6(cf9, cf10, cf11, cf12) => (set v13 : map<bool, char> | v13 in (map v12 : map<bool, char> | v12 in map[map[cf10 := 'e'] := cf12] :: (v12) := ('t')) :: (v13)) * {map[true := 'k'], map[cf10 := 'h'], map[v11.cf21 := 'f']}
		case DC4(cf4) => {map[p0 := 'a'], map[false := 'b']}
		case DC7(cf13) => {map[p0 := 'x']} - (set v14 : map<bool, char> | v14 in (seq(0x3cf, i2  => (map[false := 'h']))) :: (v14))
	})|;
	var v15 := DC0(v8);
	var v16: C3 := new C3(v4, v15, -0x8c);
	var v17: array<C3> := new C3[12] [v16, v16, v16, v16, v16, v16, v16, v16, v16, v16, v16, v16];
	v17[999] := v16;
	globalState.f2, v17[999] := v8 - (|v16.f19| + -v8), v16;
	for i3 := v8 to v8 {
		if (p0) {
			var v18 := new C2();
			var v19: multiset<bool> := multiset{p0};
			var v20: map<multiset<bool>, D5> := map[v19 := v11];
			v20 := v20[v19 := DC13(i3, |v10|, p0)];
			globalState.f9 := i3;
			globalState.f0 := false;
			v6[701] := p0;
			v6[701] := "ph" != (seq(0x25a, i4  => ('r')));
		} else {
			var v21: T0 := new C0(v15, 0x271);
			var v22: seq<T0> := [v21, v21, v21];
			var v23: seq<seq<T0>> := [v22[v21.f18 := v21]];
			var v24: array<seq<seq<T0>>> := new seq<seq<T0>>[14] [v23, [v22, v22], v23 + v23, v23 + v23, v23, if (p0) then v23 else v23, v23, v23, v23 + v23, v23, v23, v23 + [v22], v23, v23];
			v24[327] := v23;
			v24[327] := v23[i3 := v22 + v22];
			var v25: map<seq<int>, int> := map[v10 + fm11(i3, p0, v21.f18, globalState) := v8];
			var v26: map<bool, int> := map[p0 := v8];
			var v27: seq<map<bool, int>> := [v26];
			var v28: set<int> := {v8};
			var v29: map<int, map<bool, int>> := map[v8 := v27[|v28|]];
			var v30: multiset<int> := multiset{503};
			v25 := v25[v10 := -|(if (|v16.f19| in v29) then v29[|v16.f19|] else v26)| - |v30|];
			var v31: array<int> := new int[19](i5 => i5 + |[p0]|);
			var v32 := DC10(v31);
			v32 := v32;
			v31[778] := v8;
			v31[778] := v8;
			var v33: seq<bool> := [p0];
			var v34 := 'i';
			globalState.f2, v33, globalState.f9, globalState.f3 := v31[778], [v34 in v16.f19, p0], v31[778] / v21.f18, p0;
		}
		
		var v35 := new C2();
		globalState.f2 := v8;
		var v36 := 'n';
		var v37: C1 := new C1(v36);
		var v38: seq<C1> := [v37, v37, v37];
		v38, globalState.f0, globalState.f10, globalState.f9 := v38, i3 < (v10[v8] * i3), v16.f19 < (v16.f19 + v16.f19), i3;
	}
	var v39: multiset<bool> := multiset{p0, p0, p0};
	if (multiset{p0, p0} > v39) {
		var v40: array<int> := new int[25](i6 => i6 + v8);
		v40 := v40;
		var v41: map<int, int> := map[v8 := -v8];
		v41 := v41[v8 := v8];
		globalState.f9 := (-709 * v8) - (v8 * v8);
		globalState.f9 := v8;
		var v42 := DC16(v8, p0, v8, v8, p0);
		globalState.f9 := v42.cf28 + -0x306;
	} else {
		var v43: array<int> := new int[2](i7 => i7 - v8);
		v43 := new int[11];
		var v44: multiset<int> := multiset{v8, fm0(globalState)};
		var v45: set<int> := {v8};
		var v46 := new C0(v15.(cf0 := |v44|).(cf0 := v8).(cf0 := v8), |v45|);
		if (!!!p0 && !p0) {
			var v47: C0 := new C0(v15, 178);
			v43[642] := if (p0) then v8 else v8;
			var v48: map<bool, map<bool, bool>> := map[p0 := v3];
			v47, globalState.f0, v43[642] := v46, true, -(|v48| + v8);
			var v49 := new C0(v15, v43[642]);
			globalState.f3 := if (p0 in v3) then v3[p0] else false;
			globalState.f2 := v43[642];
			var v50: seq<bool> := ["vckfvwohs" < v16.f19];
			v50 := v50;
		} else {
			globalState.f2 := -((v8 * v8) - (v8 / v8));
			v1[665] := p0;
			v1[665] := p0;
			var v51 := DC32(v8);
			var v52 := 'q';
			var v53: C1 := new C1(v52);
			v8, globalState.f0, v51, v53 := v8 / v8, v1[665], v51, v53;
			var v54: map<C1, C0> := map[v53 := v46];
			var v55: seq<C0> := [v46];
			var v56: array<C0> := new C0[10] [if (p0) then v46 else v46, if (v53 in v54) then v54[v53] else v46, v46, v46, v46, v46, v55[|v10|], v46, v46, v46];
			v56[611] := v46;
			var v57 := DC6(v8, true, p0, v8);
			var v58: seq<D2> := [v57];
			var v59: multiset<D5> := multiset{v11, v11, fm17(|v4|, v58, globalState)};
			var v60 := DC3(DC2(v1[665]));
			var v61 := DC15(v8, v8, fm11(-v8, false, 0xda, globalState), v8, {-|"c"|});
			v56[611], v59, globalState.f9, globalState.f9, v60 := v46, v59 + (v59 - v59), |fm18(v8, globalState)|, -v61.cf24, v60.(cf3 := DC2(p0));
			var v62: map<char, bool> := map[v52 := !!p0];
			v62 := v62[v53.f20 := p0];
		}
		
		globalState.f10 := p0;
		var v63: C2 := new C2();
		var v64: multiset<C2> := multiset{v63, v63, v63, v63, v63};
		v64 := v64;
	}
	
	var v65 := 'j';
	var v66: C0 := new C0(v15, v8);
	v6[684] := p0;
	var v67: array<set<int>> := new set<int>[14];
	var v69: seq<set<int>> := [set v68 : int | (0x310 <= v68) && (v68 < 0xf3) :: (v68 / v8)];
	v67[925] := v69[0x4d];
	var v70: map<bool, int> := map[p0 := v8];
	var v71: multiset<C0> := multiset{v66};
	var v72: C2 := new C2();
	var v73: map<bool, C2> := map[p0 := v72];
	var v74: set<int> := {v8, |v73|, fm0(globalState)};
	v4, v65, v66, v6[684], v67[925] := v4, v16.f19[|v70| % |v71|], if (fm1(p0, v4, globalState)) then v66 else v66, p0, (v74 * v74) * (v74 * v74);
	r0 := if ('g' in v4) then v8 else v8;
	var v75: map<bool, char> := map[v6[684] := v65];
	var v76 := DC34(v75);
	var v77: map<int, map<bool, char>> := map[v8 := (v76.(cf61 := v75)).cf61];
	r1 := |v77[v8 % v8 := map[p0 := v65]]|;
}
trait T0 {
	const f17 : D0
	const f18 : int
	function fm2(globalState: GlobalState): D0 
}

class C0 extends T0 {
	constructor (f17 : D0, f18 : int) {
		this.f17 := f17;
		this.f18 := f18;
	}
	
	function fm2(globalState: GlobalState): D0 {
		match f17 {
			case DC0(cf0) => f17
		}
	}
}

class C1 {
	const f20 : char
	constructor (f20 : char) {
		this.f20 := f20;
	}
	
	method m3(p0: bool, globalState: GlobalState) returns (r0: int, r1: multiset<multiset<bool>>) {
		var i0 := 0;
		while (p0)
			decreases 100 - i0
		{
			if (i0 >= 100) {
				break;
			}
			
			i0 := i0 + 1;
			var v0: array<bool> := new bool[27];
			var v1 := 0x3de;
			var v2: map<array<bool>, int> := map[if (p0) then v0 else v0 := v1];
			var v3: seq<bool> := [(fm5(-v1, 121, globalState)).cf2];
			globalState.f2 := if (v0 in v2) then v2[v0] else -|v3|;
			if (p0) {
				var v4: map<int, int> := map[v1 := v1 - v1];
				v4 := v4[v1 := 0x130 - -v1];
				var v5, v6, v7, v8 := m4(globalState);
				globalState.f10 := v1 !in (set v9 : int | (959 <= v9) && (v9 < 0x207) :: (v9 + v6));
				v1 := v6;
				var v10 := DC1(v0);
				var v11 := DC3(v10);
				var v12 := DC1(v0);
				var v13: array<D1> := new D1[15] [v12, v12, v12, v12, v12, v12, v12, DC1(v0), v12, v12, v12, v12, v12, v12, v12];
				var v14: seq<array<D1>> := [v13, v13];
				var v15: array<array<D1>> := new array<D1>[19] [v13, v13, v13, v13, if (p0) then v13 else v13, v13, v13, v13, v13, if (v7) then v13 else v13, v13, v13, v13, v13, if (p0) then v13 else v13, v13, v13, v13, v14[v1]];
				var v16: map<int, bool> := map[v6 := v5];
				var v17 := "k";
				var v18: set<char> := {f20};
				var v19: array<int> := new int[12] [v1, v6, v6, v1, 840, --v1, |v4|, |v17|, |v18|, v6, v6, fm0(globalState)];
				var v20: map<bool, array<int>> := map[if (v6 in v16) then v16[v6] else v5 := v19];
				v11, v7, v15, v5, v0 := v11, (v6 != v1) in v20, v15, v6 == (-v1 / v1), v0;
			} else {
				var v21 := "dcfqu";
				v21, globalState.f2 := v21, if (p0) then v1 - v1 else fm0(globalState);
				var v22: map<array<bool>, bool> := map[v0 := p0];
				globalState.f0 := (v22 + v22) != v22;
				var v23 := DC0(-fm0(globalState));
				var v24 := new C0(v23, v1);
				var v25 := new C0(v23, v1);
				globalState.f2 := --v1;
			}
			
			if ((-0x55 % --525) < v1) {
				globalState.f10 := v3[v1 / -v1];
				v0[733] := p0;
				var v26 := "t";
				var v27: map<int, bool> := map[|(("nogbhr")[v1 := f20] + (fm6(fm1(p0, v26, globalState), globalState))[v1 := fm7(f20, p0, globalState)])| := {v1} !! fm8(|multiset(v26)|, v1, globalState)];
				v0[336] := p0;
				var v28: map<bool, bool> := map[p0 := p0];
				var v29: map<map<bool, bool>, int> := map[v28 := -v1];
				var v30: multiset<int> := multiset{v1};
				var v31: set<int> := {v1};
				var v32: seq<int> := [|v31|, v1, |v3|, v1, v1];
				var v34: map<int, map<int, bool>> := map[if (v1 in v30) then v30[v1] else |v32| := map v33 : int | v33 in v27 :: (v33 * 0xd6) := (p0)];
				var v35: set<bool> := {if (p0 in v28) then v28[p0] else p0, p0, p0};
				var v36 := DC2(p0);
				v0[733], v27, v0[336], globalState.f0, globalState.f9 := v28 !in v29, if (p0) then v27 else if (v1 in v34) then v34[v1] else v27, p0, p0 !in (v35 + {v36.cf2, p0, p0}), fm0(globalState);
				globalState.f0 := v36.cf2;
				var v37: map<char, bool> := map[f20 := p0];
				v37 := v37[f20 := p0];
				var v38: set<map<int, bool>> := {v27 + v27};
				v38 := v38;
			} else {
				var v39 := DC0(v1);
				var v40: C0 := new C0(v39, v1);
				v40 := v40;
				globalState.f0 := p0;
				globalState.f9 := v1;
				globalState.f2 := v1;
				var v41 := DC1(v0);
				var v42: map<char, D1> := map[f20 := v41];
				var v43 := "im";
				var v44: array<string> := new string[2] [v43, v43];
				v44[75] := v43;
				globalState.f3, globalState.f9, v42, v44[75] := p0, v1 / 256, map[f20 := v41], fm6(p0, globalState);
			}
			
			var v45 := DC0(v1);
			var v46 := new C0(v45, v1);
		}
		globalState.f3 := p0;
		var v47 := 825;
		r0 := -v47 / 127;
		var i1 := 0;
		while (!p0)
			decreases 100 - i1
		{
			if (i1 >= 100) {
				break;
			}
			
			i1 := i1 + 1;
			var v48: map<bool, bool> := map[p0 := p0];
			v48 := v48 + v48;
			globalState.f10 := !p0 && (--0x185 > |[p0, p0, p0]|);
			var v49 := "tht";
			v49 := v49 + v49;
			var v50 := DC0(v47);
			var v51: map<int, bool> := map[v50.cf0 := fm1(p0, v49, globalState)];
			var v52, v53 := m0(p0, v51, globalState);
		}
		var v54: array<bool> := new bool[26];
		forall i2 | 0 <= i2 < v54.Length {
			v54[i2] := !p0;
		}
		globalState.f2 := v47;
		var v55 := "iakx";
		r0 := if (p0) then |v55| else v47;
		var v56: multiset<bool> := multiset{p0};
		var v57: set<int> := {v47};
		var v58: seq<bool> := [p0, p0, p0, fm1(p0, v55, globalState), p0];
		var v59: multiset<multiset<bool>> := multiset{v56, v56, if (p0) then fm9(|v57|, p0, p0, --|v58|, globalState) else multiset(v58), v56[p0 := |v57|]};
		r1 := v59;
	}
	method m4(globalState: GlobalState) returns (r0: bool, r1: int, r2: bool, r3: bool) {
		var v0 := false;
		if (v0) {
			var v1 := 0x311;
			var v2: seq<bool> := [v0];
			var v3: multiset<bool> := multiset{v0, v1 == v1, v2[v1]};
			var v4: seq<int> := [v1, -v1];
			var v5 := "hfyg";
			v3, globalState.f3, globalState.f2 := multiset((if (v0) then [v0, v0] else v2)[0x3b3 := v0]) + v3, v0, (v1 / |v4|) - -(v1 + -|v5|);
			var v6: array<D1> := new D1[16];
			var v7 := DC4(v6);
			v6 := v7.cf4;
			var v8: array<int> := new int[27];
			v8[644] := v1 - v1;
			v8[644] := v1;
			globalState.f0 := false;
			var v9: array<seq<D2>> := new seq<D2>[23];
			var v10: seq<array<seq<D2>>> := [v9, v9];
			var v11: seq<array<seq<D2>>> := [v9, v10[-315]];
			var v12: map<bool, array<seq<D2>>> := map[v0 := v9];
			var v13: array<array<seq<D2>>> := new array<seq<D2>>[28] [v9, v9, v9, v9, v9, v9, v9, v9, v9, v9, v9, v11[v1], v9, v9, if (true) then v9 else v9, v9, v9, v9, v9, v9, if (true) then v11[v1] else v9, v9, v9, v9, v9, v9, if (v0 in v12) then v12[v0] else v9, v9];
			v13[554] := v9;
			v13[554] := v9;
		} else {
			var v14 := -271;
			var v15: seq<int> := [v14, v14, v14];
			globalState.f0 := (v14 % v14) >= v15[v14];
			var v16: multiset<bool> := multiset{v0, v0};
			var v17: map<bool, multiset<bool>> := map[v0 := v16];
			var v18 := DC0(|v17|);
			match (v18) {
				case DC0(cf0) =>
					var v19 := new C0(v18, 0x5a);
					var v20: map<bool, int> := map[!v0 := -251];
					var v21: seq<map<bool, int>> := [v20];
					v21 := seq(787, i0  => (v20));
					var v22 := 'h';
					var v23 := "nntcghypc";
					var v24: array<int> := new int[8];
					v24[426] := cf0;
					var v25: array<bool> := new bool[3];
					v25[148] := v16 <= v16;
					var v26: seq<array<bool>> := [v25];
					var v27: map<seq<array<bool>>, char> := map[v26 := v22];
					v22, v23, v24[426], v25[148] := if (v26 in v27) then v27[v26] else v22, (seq(0x255, i1  => (v23[cf0]))) + (seq(0x291, i2  => (v22))), cf0, !v0;
					v20 := v20[fm1(v0, v23, globalState) := v24[426]];
			}
			
			var v28: C0 := new C0(v18, v14);
			var v29: multiset<C0> := multiset{v28};
			var v30 := DC8(v29);
			var v31: map<int, int> := map[v14 := v14];
			var v32: array<bool> := new bool[25];
			var v33: multiset<int> := multiset{|fm8(v14, v14, globalState)|, v14};
			var v34 := DC5(v0, |v31|, v32, v33);
			var v35: seq<bool> := [v0, !v0, v0];
			var v36: array<seq<bool>> := new seq<bool>[16] [[v34.cf5, v0], [v0, v0], v35, v35, v35, v35, v35, v35, v35, v35, v35, v35, v35, v35, v35, v35];
			var v37: map<multiset<C0>, array<seq<bool>>> := map[v30.cf14 + v29 := v36];
			v37 := v37[v29 := v36];
			var v38 := "eancxn";
			v38 := v38;
			globalState.f0 := (v14 * v14) < -v14;
		}
		
		var i3 := 0;
		while (true)
			decreases 100 - i3
		{
			if (i3 >= 100) {
				break;
			}
			
			i3 := i3 + 1;
			var v39: seq<bool> := [v0];
			var v40 := "htdkpplcq";
			var v41: array<int> := new int[22](i4 => i4 + 0xdd);
			var v42: map<bool, array<int>> := map[true := v41];
			if (v39[-|v40| % |v42|]) {
				var v43: set<bool> := {v0};
				var v44 := 0xc0;
				var v45: map<bool, int> := map[v0 := v44];
				var v46 := DC0(|v45|);
				var v47: map<set<bool>, bool> := map[v43 := v39[v46.cf0]];
				v47 := v47;
				globalState.f3 := v0 || v0;
				r3 := v0;
				globalState.f10 := v0;
				var v48: seq<int> := [v44, v44];
				var v49: multiset<int> := multiset{|{f20, v40[v44]}|};
				var v50: array<bool> := new bool[1];
				var v51 := DC5(v0, v44, v50, multiset{v44});
				var v52: array<bool> := new bool[5] [multiset(v48) >= v49, v39[-v44], v0, v0, v51.cf5];
				v52[784] := v0;
				v52[784] := v0;
			} else {
				v40 := (seq(0x2fc, i5  => (f20))) + v40;
				var v53: map<seq<bool>, int> := map[v39 := 0x22b];
				var v54 := DC0(|v53|);
				var v55: C0 := new C0(v54, |v40|);
				var v56: seq<C0> := [v55, v55, v55];
				var v57 := 30;
				var v58 := DC9(v55);
				var v59: array<C0> := new C0[24] [v55, v55, v55, v55, v56[v57], v55, v55, v55, v55, v55, v55, v55, v55, v55, v55, v55, v55, v55, v55, v55, v55, v55, v55, v58.cf15];
				v59[899] := v55;
				v59[899] := v55;
				globalState.f2 := v57;
				var v60: array<bool> := new bool[17];
				var v61: map<bool, array<bool>> := map[v0 := v60];
				var v62: array<array<bool>> := new array<bool>[17] [v60, v60, v60, v60, v60, v60, v60, if (v0 in v61) then v61[v0] else v60, v60, v60, v60, v60, v60, v60, v60, v60, v60];
				v62[191] := v60;
				var v63: multiset<int> := multiset{v57, v57, v57};
				v57, v62[191] := |v63|, v60;
				globalState.f2 := v57;
			}
			
			v0 := v0;
			var v64: array<bool> := new bool[28] [v0, v0, v0, v0, fm1(true, v40, globalState), v0, fm1(v0, v40, globalState), v0, v0, v0, !v0, v0, v0, v0, v0, v0, true, v0, v0, v0, v0, v0, v0, false, v0, v0, v0, fm1(v0, v40[0x274 := f20], globalState)];
			var v65 := DC1(v64);
			var v66: seq<array<bool>> := [v64];
			var v67: array<array<bool>> := new array<bool>[25] [v64, v64, v64, v64, v64, v64, v64, if (v0) then v64 else v64, v64, v64, v64, v64, v64, v64, v64, v64, v64, v64, v64, v64, v64, v65.cf1, v64, if (v0) then v64 else v66[|v39|], v64];
			v67 := new array<bool>[5];
			var v68 := 0xe;
			r2 := (v68 / |{v40, v40}|) > |v39|;
		}
		var v69: array<bool> := new bool[3];
		var v70 := DC1(v69);
		var v71: map<D1, int> := map[v70 := 421];
		var v72 := 662;
		var v73: map<int, int> := map[if (v70 in v71) then v71[v70] else -v72 := v72];
		var v74: seq<char> := [f20];
		v73 := v73[|v74| - v72 := -351];
		v69[401] := v0;
		var v75: seq<int> := [v72];
		v69[401] := (v75 < v75) <== false;
		var v76: seq<bool> := [v69[401]];
		var v77 := DC2(v76[v72] <== v69[401]);
		v77 := v77.(cf2 := false);
		var v78: map<int, bool> := map[-v72 := v69[401]];
		var v79, v80 := m0(if (v76[v75[v72]]) then v69[401] else v69[401], v78[v72 := v0], globalState);
		r0 := !v69[401];
		var v81 := DC12(v74);
		r1 := |(v81.cf18 + (v74 + v74))|;
		r2 := fm1(v0, v74, globalState);
		r3 := !v0 <== (v69[401] <==> v69[401]);
	}
}

class C2 {
	constructor () {
	}
	
	method m2(globalState: GlobalState) returns (r0: int) {
		var v0 := 0x319;
		globalState.f3 := v0 == v0;
		var v1 := DC0(0x54);
		var v2 := 'f';
		var v3: C1 := new C1(v2);
		var v4: multiset<C1> := multiset{v3};
		var v5: multiset<int> := multiset{906};
		var v6 := new C0(v1, |{v4}| + (if (v0 in v5) then v5[v0] else 302));
		var v7 := true;
		var v8 := "rfvwqmbea";
		var v9 := DC13(v0, v0, fm1(!v7, v8, globalState));
		globalState.f3 := match v9 {
			case DC13(cf19, cf20, cf21) => false
			case DC12(cf18) => false
		};
		var v10: array<seq<bool>> := new seq<bool>[10];
		forall i0 | 0 <= i0 < v10.Length {
			v10[i0] := ([v7] + [v7]) + ([v7, v7] + [v7]);
		}
		var v11 := DC12(v8);
		if (match v11 {
			case DC13(cf19, cf20, cf21) => false
			case DC12(cf18) => !v7
		}) {
			v7 := v7;
			var v12: set<bool> := {!v7, v7, v7};
			var v13: array<bool> := new bool[24](i1 => false);
			var v14: map<int, array<bool>> := map[if (v7) then |v12| else v0 := v13];
			v14 := v14[0x179 := v13];
			if ((v0 + v0) > v0) {
				globalState.f2 := v0;
				globalState.f10 := !((if (v7) then v2 else v3.f20) in v8);
				globalState.f3 := (v0 / -v0) < v0;
				r0 := v0;
				var v15 := DC14(v3);
				var v16: array<C1> := new C1[9] [if (v7) then v3 else v15.cf22, v3, v3, v3, v3, v3, v3, v3, v3];
				v16 := v16;
			} else {
				var v17: array<string> := new string[16];
				v17[553] := if (v7) then v8[v0 := v2] else v8;
				var v18: array<T0> := new T0[23];
				var v19: multiset<array<T0>> := multiset{v18};
				globalState.f0, v17[553], globalState.f3, r0 := if (v7) then !v7 else v7, (seq(143, i2  => (v3.f20)))[-v0 := 'h'], (v19 + v19) >= v19, v0;
				var v20: map<int, bool> := map[v0 / |v5| := v7];
				var v21: set<char> := {'e'};
				v20 := v20[-fm0(globalState) + v0 := v21 > v21];
				v13[626] := true;
				var v22: multiset<bool> := multiset{v7, v7};
				v13[626] := if (!v7 ==> v7) then v22 > v22 else v7;
				v7 := !(v17[553] == v17[553]);
				var v23: map<int, char> := map[0x25e := v2];
				var v24: seq<map<int, char>> := [v23];
				globalState.f0 := v0 <= |v24[0x2e8]|;
			}
			
			var v25: array<set<array<bool>>> := new set<array<bool>>[6];
			var v26: set<array<bool>> := {v13};
			v25[453] := v26;
			v25[453] := v26 - v26;
			globalState.f0 := v7;
		} else {
			globalState.f2 := v0;
			var v27: map<int, C1> := map[|v8| := v3];
			var v28: map<map<int, C1>, bool> := map[v27 := !v7];
			v28 := v28[v27 := false];
			if (v7) {
				var v29 := new C1(if (v7) then v3.f20 else 'm');
				globalState.f9 := (v0 % |fm10(globalState)|) - (v0 + |v8|);
				var v30: array<bool> := new bool[5](i3 => true);
				var v31 := DC5(v7, 324, v30, v5);
				var v32: multiset<D2> := multiset{v31};
				var v33: map<bool, string> := map[v7 := "saqvvnypo"];
				var v34: seq<multiset<D2>> := [v32, v32[v31 := -|v33|]];
				v34 := v34;
				globalState.f0 := !v7 ==> v7;
				globalState.f2 := fm0(globalState);
			} else {
				var v35 := new C1(v3.f20);
				var v36: multiset<string> := multiset{v8};
				var v37: seq<D5> := [v9, DC13(v0, |v36|, v7)];
				var v38: seq<bool> := [v7, v37 <= v37[v0 := v9]];
				globalState.f10 := v38[|[v7, false]| % v0];
				globalState.f2 := v0;
				var v39: map<D5, char> := map[v11 := 'd'];
				var v40: array<char> := new char[27] [v2, fm7(v2, fm1(v7, seq(0x127, i4  => (v3.f20)), globalState), globalState), v35.f20, v3.f20, v3.f20, v3.f20, v3.f20, v35.f20, v3.f20, 'b', v2, v3.f20, v2, v2, v3.f20, v2, v3.f20, if (v11 in v39) then v39[v11] else 'f', v3.f20, v35.f20, v35.f20, v2, v2, v3.f20, 'j', v3.f20, v2];
				v40[199] := v35.f20;
				var v41: array<set<bool>> := new set<bool>[5];
				var v42: map<array<set<bool>>, char> := map[v41 := v3.f20];
				var v43: seq<array<set<bool>>> := [v41];
				v40[199] := if (v43[-v0] in v42) then v42[v43[-v0]] else v3.f20;
				var v44: array<bool> := new bool[10](i5 => v7);
				var v45: multiset<seq<int>> := multiset{[v0]};
				var v46: seq<multiset<seq<int>>> := [v45];
				v44[330] := v46[|[v3.f20]|] < v45;
				var v47 := DC17(v40[199], v7, !v7, v3.f20);
				v44[330] := v47.cf34;
			}
			
			globalState.f2 := 0x59;
			var v48 := new C1('d');
		}
		
		var v49: set<bool> := {v7};
		var v50 := DC16(v0, v7, |v8|, |v49|, v7);
		if ((if (v7) then v7 else v50.cf32) <== v7) {
			var v51: seq<int> := [|[v7, v7]| - v0, v0];
			v51 := v51;
			v7 := v7;
			var v52 := new C1('k');
			globalState.f3 := v7;
			v2 := v52.f20;
		} else {
			var v53: C0 := new C0(v1, v0);
			var v54: map<int, bool> := map[v0 := v7];
			var v55: seq<bool> := [v7, v7];
			var v56: array<bool> := new bool[25] [v54 == v54, v7, v0 <= v0, !v7, v7, v7, v55[v0], v7, if (!false) then v7 else v7, fm1(v7, fm6(!v7, globalState), globalState), v7, v7, |v49| >= v0, v5 != v5, v7, v7, v7, v7, v7, true, v7, v50.cf29, v7, v7, v7];
			v8, v53, v8, globalState.f11 := v8 + v8, v53, v8, v56;
			globalState.f9 := v0;
			var v57: map<bool, bool> := map[v7 := v7];
			var v58: seq<D5> := [v9.(cf19 := |v55|), v9];
			v57 := v57[fm1(v7, seq(269, i6  => (v3.f20)), globalState) := v58 < (seq(0x30e, i7  => (v9)))];
			var v59 := DC2(v7);
			var v60 := DC3(v59);
			v60 := v60;
			if (v7) {
				globalState.f10 := if (v7) then if (v7) then v7 else fm1(v7, "xtbk", globalState) else v7;
				var v61: array<int> := new int[21];
				v61 := v61;
				globalState.f3 := v7;
				var v62: map<multiset<int>, int> := map[multiset{|v57|} := v0];
				v62 := v62;
				globalState.f2 := v0 - v0;
			} else {
				globalState.f10 := v7;
				var v63: array<int> := new int[21];
				var v64 := DC10(v63);
				var v65 := DC11(v64);
				var v66: seq<int> := [v0, v0];
				var v67: map<D4, seq<int>> := map[v65 := v66];
				var v68: map<seq<int>, int> := map[if (v65 in v67) then v67[v65] else v66 := v0 % v0];
				v68 := v68[if (v7) then v66 else v66 := 329];
				v60 := if (v7 || v7) then v60 else v60;
				v0 := v0;
				var v69: map<array<bool>, bool> := map[v56 := v7];
				var v71: set<int> := {v0, v0};
				globalState.f0 := if (v56 in v69) then v69[v56] else (set v70 : int | (-694 <= v70) && (v70 < -854) :: (v70 - v0)) > v71;
			}
			
		}
		
		r0 := -(v0 - (v0 * v0));
	}
}

class C3 extends T0 {
	const f19 : string
	constructor (f19 : string, f17 : D0, f18 : int) {
		this.f19 := f19;
		this.f17 := f17;
		this.f18 := f18;
	}
	
	function fm2(globalState: GlobalState): D0 {
		f17
	}
	function fm3(p0: bool, p1: int, globalState: GlobalState): int {
		f18
	}
	function fm4(p0: int, p1: string, p2: seq<int>, globalState: GlobalState): int {
		f18 % (f18 + f18)
	}
	method m1(globalState: GlobalState) returns (r0: array<int>, r1: bool) {
		var v0: seq<int> := [f18];
		for i0 := -v0[-233] to f18 * |f19| {
			match (f17) {
				case DC0(cf0) =>
					globalState.f2 := i0;
					var v1 := true;
					var v2: map<bool, int> := map[!(if (v1) then v1 else v1) := cf0];
					globalState.f2 := -(if ((i0 <= |f19|) in v2) then v2[i0 <= |f19|] else 0x2d0);
					globalState.f9 := |f19|;
					var v3: array<array<bool>> := new array<bool>[18];
					var v4: map<bool, bool> := map[v1 := v1];
					var v5: array<bool> := new bool[6] [v1, if (!v1 in v4) then v4[!v1] else v1, v1, fm1(true, "ti", globalState), v1, true];
					v3[652] := v5;
					v5[99] := true;
					var v6 := DC1(v5);
					var v7: seq<bool> := [v1];
					v3[652], globalState.f2, globalState.f0, v5[99] := v6.cf1, i0, if (v7[i0]) then v1 else v1, v1;
			}
			
			var v8: array<bool> := new bool[19];
			var v9 := DC1(v8);
			match (v9) {
				case DC2(cf2) =>
					var v10 := 'x';
					var v11 := new C1(v10);
					globalState.f9 := f18;
					var v12 := new C0(f17, -f18);
					var v13: map<int, int> := map[i0 := 549];
					var v14: map<bool, int> := map[cf2 := f18];
					v8[734] := {i0, f18} != fm8(if (cf2 in v14) then v14[cf2] else i0, i0, globalState);
					v10, globalState.f9, v13, v8[734] := v10, i0, v13 + v13, i0 < f18;
				case DC1(cf1) =>
					var v15 := false;
					v8[660] := v15;
					var v16: T0 := new C0(DC0(f18), i0);
					var v17: set<T0> := {v16, v16};
					v8[660] := (v17 - v17) > v17;
					var v18: array<int> := new int[27];
					v18[162] := 874;
					v8[660], v18[162] := v8[660], f18;
					v16 := v16;
					var v19: set<int> := {f18, i0};
					var v20: map<set<int>, array<bool>> := map[v19 := cf1];
					v20 := v20[v19 := cf1];
				case DC3(cf3) =>
					var v21 := true;
					var v22 := 't';
					var v23: C1 := new C1(v22);
					var v24: seq<C1> := [if (v21) then v23 else v23, v23, v23, if (v21) then v23 else v23, v23];
					v24 := v24 + v24;
					var v25: C0 := new C0(DC0(|f19|), f18);
					var v26 := DC9(v25);
					var v27 := DC11(v26);
					var v28 := DC11(v27);
					var v29: map<bool, D4> := map[v21 := v28];
					var v30: set<map<bool, D4>> := {v29};
					var v31 := DC18(v30);
					v30 := v31.cf37;
					var v32, v33 := v23.m3(f18 == f18, globalState);
					var v34: C2 := new C2();
					var v35: array<C2> := new C2[13] [v34, v34, v34, v34, DC22(v34).cf45, v34, if (v21) then v34 else v34, v34, v34, v34, v34, v34, v34];
					v35[400] := v34;
					v35[400] := v34;
			}
			
			var v36: array<set<D2>> := new set<D2>[12];
			var v37 := false;
			var v38: map<int, bool> := map[|f19| := v37];
			var v40: set<int> := {-f18, i0};
			var v41: seq<map<int, bool>> := [v38, v38, map v39 : int | v39 in v40 :: (v39 / i0) := (v37)];
			var v42: map<bool, int> := map[v37 := fm4(|v41|, f19, v0, globalState)];
			var v43: array<C2> := new C2[6];
			var v44: multiset<array<C2>> := multiset{v43, v43, v43, v43, v43};
			var v45: multiset<int> := multiset{|v42|, if (v43 in v44) then v44[v43] else f18};
			var v46 := DC5(v37, -f18, v8, v45);
			var v47: set<D2> := {DC5(v37, i0, v8, v45), v46};
			v36[807] := v47;
			v36[807] := v47;
			var v48: T0 := new C0(f17, f18);
			var v49: map<T0, bool> := map[v48 := v37];
			globalState.f3 := (v49 + v49) != (v49 + map[v48 := v37]);
		}
		for i1 := f18 to -0x6d {
			var v50 := false;
			globalState.f2 := (if (v50) then fm3(v50, f18, globalState) else f18) + f18;
			var v51 := 'x';
			v51 := v51;
			globalState.f9 := i1 + -v0[|f19|];
			v50 := v50;
		}
		var v52: array<int> := new int[25](i2 => i2 + -0x6b);
		v52[497] := f18;
		v52[497] := f18 - -fm0(globalState);
		var v53: array<string> := new string[14];
		var v54: array<bool> := new bool[16](i3 => "alyvhe" != f19);
		var v55 := true;
		v54[961] := v55;
		var v56: multiset<int> := multiset{|map[|f19| := v55]|};
		v53, v54[961] := v53, v56 >= v56;
		var v57 := "baxh";
		var v58: multiset<array<int>> := multiset{v52};
		v57, globalState.f9, v58, globalState.f0 := f19, f18, v58, if (v54[961]) then v54[961] else v54[961];
		var i4 := 0;
		while (v55)
			decreases 100 - i4
		{
			if (i4 >= 100) {
				break;
			}
			
			i4 := i4 + 1;
			if (v55) {
				var v59: map<int, bool> := map[f18 := v54[961]];
				var v60, v61 := m0(v55, v59, globalState);
				var v63: set<map<int, bool>> := {v59};
				var v64 := DC15(|(map v62 : int | (-0x142 <= v62) && (v62 < 0x268) :: (v62 + v52[497]) := (-v61))|, v61, [f18], f18, {v60, |v63|, -v52[497], v60, v61});
				var v65 := 'c';
				var v66: map<bool, int> := map[v54[961] := f18];
				v64, globalState.f10, v65, v60, globalState.f3 := v64, (v66 + v66) != (v66 + v66), v65, v61 * (if (v54[961]) then v60 else v60), fm1(fm1(v54[961], v57, globalState), ("g")[v52[497] := v65], globalState);
				var v67: array<map<C1, int>> := new map<C1, int>[1];
				var v68: C1 := new C1(v65);
				var v69: map<C1, int> := map[v68 := v52[497]];
				v67[264] := v69;
				var v70 := DC1(v54);
				var v71: seq<D1> := [v70];
				globalState.f9, v61, v54[961], v67[264], globalState.f2 := v52[497], fm0(globalState) * (v52[497] * |v71|), v54[961], map[v68 := 670] + (map[v68 := 0x8f] + v69), v52[497];
				globalState.f3 := v61 <= v52[497];
				globalState.f9 := |f19|;
			} else {
				globalState.f2 := -0x84 % f18;
				globalState.f2 := f18;
				v57 := f19;
				v0 := v0;
				var v72: map<int, int> := map[f18 := |f19| / v52[497]];
				v72 := v72[v52[497] := 0x37];
			}
			
			globalState.f9 := f18 - (if (v55) then fm3(v55, f18, globalState) else 0x352);
			var v73: seq<string> := [seq(547, i6  => ('f')), "bj", v57, v57, v57];
			var v74: seq<string> := [seq(0x234, i5  => ('h')), f19, f19, v73[f18]];
			var v75: seq<bool> := [v54[961]];
			v52[497], v52[497] := f18, |v73[fm3(v75[v52[497]], |v75|, globalState)]| / v52[497];
			globalState.f9 := v52[497] * -v52[497];
		}
		r0 := v52;
		r1 := v54[961];
	}
}

method Main() {
	var v0 := true;
	var v1: array<bool> := new bool[21](i0 => v0);
	var v2 := 213;
	var v3: map<bool, int> := map[v0 := -v2];
	var v4 := "l";
	var v5: map<bool, bool> := map[v0 := v0];
	var globalState := new GlobalState(true, [v0], 172, false, true, 'f', 669, v1, multiset{v2, v2}, 0x101, false, if (v0) then v1 else v1, v3, v4, false, v5 + v5, true);
	if (v0) {
		v3 := v3[v0 := v2];
		var v6: array<seq<bool>> := new seq<bool>[12](i1 => [false]);
		var v7: seq<bool> := [v0];
		v6[676] := v7;
		v6[676] := v7[if (v0) then -v2 else v2 := v7[-v2]];
		var v8: map<int, int> := map[v2 := v2];
		var v9: map<map<int, int>, bool> := map[v8 := v0];
		v2 := v2 - |v9|;
		var v10: set<int> := {v2};
		var v11: seq<int> := [0x139];
		var v12: seq<int> := [|v11|];
		var v13: array<int> := new int[28] [v2, 0x4d, |v6[676]| * v2, v2, fm0(globalState) * -v2, v2, fm0(globalState), v2, v2, 0x163 / v2, |v10| + v2, v2, v2, 0x150, v2, fm0(globalState) - v2, v2, v2, v2, v2 / v2, v2, if (v0) then v2 else v2, |v4|, v2, v2 / |v11|, v2, v2, fm0(globalState)];
		var v14: map<int, array<int>> := map[-0x291 := v13];
		v1, globalState.f9, v13 := v1, (v2 * v2) - (v2 / v2), if ((v2 % v2) in v14) then v14[v2 % v2] else v13;
		var v15: array<array<bool>> := new array<bool>[12];
		v15 := v15;
	} else {
		var v16 := 'r';
		var v17: map<char, int> := map[v16 := v2];
		globalState.f2 := |(map[v17 := v0])[v17 := v0]|;
		var v18: seq<bool> := [v0];
		v18 := [!v0];
		var v20: multiset<int> := multiset{v2, v2, v2, v2, v2};
		var v21: multiset<bool> := multiset{v0, v0};
		var v22, v23 := m0(v4 <= v4, (map v19 : int | v19 in v20 :: (v19 % v2) := (v0))[|v21| := v18[v2]], globalState);
		var v24: seq<string> := [v4, seq(0x30f, i2  => (v16)), v4, v4];
		if (fm1(false, v24[v22], globalState)) {
			v16 := v16;
			globalState.f3 := fm0(globalState) > v2;
			var v25 := DC0(v23);
			var v26: map<int, bool> := map[v25.cf0 := !v0];
			var v27, v28 := m0(v0, v26, globalState);
			globalState.f9 := v28;
			var v29: map<seq<bool>, bool> := map[v18 + [false] := v0];
			var v30: seq<int> := [v27];
			var v31: multiset<seq<int>> := multiset{v30};
			globalState.f3 := !(if (v18 in v29) then v29[v18] else v31 >= multiset{v30, v30});
		} else {
			v16 := v16;
			v18 := v18 + v18;
			var v32 := new C1(v16);
			var v33: map<bool, seq<bool>> := map[v0 := v18];
			var v34: set<seq<bool>> := {v18, v18, v18, v18, if (v18[v2] in v33) then v33[v18[v2]] else v18};
			globalState.f10 := v34 < (v34 + v34);
			v3 := v3[if (v0) then v0 else v0 := v23];
		}
		
		globalState.f10, globalState.f0, globalState.f0 := !true ==> true, multiset(v18 + [v0, v0, v0]) !! v21, v2 <= v22;
	}
	
	var v35: multiset<int> := multiset{v2, v2, v2};
	var v36 := DC16(v2, v0, |"b"|, v2, v0);
	globalState.f9, globalState.f9 := if (-(v36.cf31 / v2) in v35) then v35[-(v36.cf31 / v2)] else -v2, v2;
	var v37: seq<int> := [-496, v2, v2];
	v37 := (v37 + [v2]) + v37;
	var v38 := DC2(false);
	var v39: multiset<bool> := multiset{v0};
	var v40: set<bool> := {v0, v0};
	var v41: array<D1> := new D1[20] [v38, v38, v38.(cf2 := v0), v38, v38, v38, DC2(v0), DC2(v0), v38, v38, v38, v38, v38.(cf2 := v0), v38.(cf2 := fm1(v0, "jol", globalState)), v38, v38, v38, fm5(v2, if (false in v39) then v39[false] else |v40|, globalState), v38, v38];
	var v42 := DC4(v41);
	v41 := if (v4 < "e") then v41 else v42.cf4;
	var v43 := 'y';
	var v44: C1 := new C1(v43);
	var v45 := DC14(v44);
	match (v45.(cf22 := if (v0) then v44 else v44)) {
		case DC15(cf23, cf24, cf25, cf26, cf27) =>
			var v46: multiset<char> := multiset{v44.f20, v43, v43};
			globalState.f9, v2 := cf23, if (760 >= |map[fm0(globalState) := cf23]|) then 0x1be else if (v0) then |v46| else cf23;
			var v47 := DC0(|v4|);
			var v48: T0 := new C0(v47, cf23);
			var v49: array<map<bool, C0>> := new map<bool, C0>[16];
			var v50: C0 := new C0(DC0(-cf26), v48.f18);
			var v51: map<bool, C0> := map[v0 := v50];
			v49[787] := v51 + v51;
			var v52: multiset<array<bool>> := multiset{v1};
			var v53: map<int, int> := map[cf26 := cf26];
			var v54: map<bool, map<int, int>> := map[false := v53];
			v48, cf25, v44, v49[787] := v48, fm11(|v4| % |multiset(v37)|, v0, if (v1 in v52) then v52[v1] else |(if (v0 in v54) then v54[v0] else map[v48.f18 := v2])|, globalState), v44, v51;
			var v55: array<int> := new int[11];
			v55 := v55;
			globalState.f0 := v0;
		case DC16(cf28, cf29, cf30, cf31, cf32) =>
			var v56, v57 := v44.m3(v0, globalState);
			var v58: seq<bool> := [cf29];
			var v59: map<int, seq<int>> := map[cf28 := fm11(cf30, v0, cf28, globalState)];
			var v60: seq<seq<int>> := [v37[-15 := v56]];
			var v61: map<string, bool> := map[v4 := cf32];
			var v62: array<seq<int>> := new seq<int>[13] [v37, v37, v37, v37, v37 + v37, v37, [fm0(globalState), 0x127], [cf31, 173, |multiset(v58[cf28 := v0])|, 3], v37[v2 := cf28], if (cf28 in v59) then v59[cf28] else v37, v37 + (seq(0x97, i3  => (|v40|))), v37, (v60[cf31])[|v61| := cf30]];
			v62[373] := v37;
			v62[373] := v37;
			var v63: array<int> := new int[3];
			var v64 := DC23(cf28, v63, cf31);
			var v65: set<int> := {cf28};
			var v66: map<int, set<int>> := map[v2 := v65];
			var v67: seq<set<int>> := [v65, if (|v5| in v66) then v66[|v5|] else v65];
			match (fm12(v64.cf46, v67, cf30, cf31, globalState)) {
				case DC13(cf19, cf20, cf21) =>
					globalState.f2 := cf20;
					v63[881] := |map['p' := v3]|;
					var v68: array<char> := new char[5];
					v68[898] := v44.f20;
					var v69: map<int, int> := map[cf31 := |v58|];
					var v70: map<map<int, int>, string> := map[v69 := v4];
					var v71: map<int, map<int, int>> := map[cf20 := v69];
					var v72 := DC12(v4);
					var v73: array<string> := new string[28] [v4, "sfvx" + "tunb", "laijoy", "l", v4, v4, if ((if (|v62[373]| in v71) then v71[|v62[373]|] else v69) in v70) then v70[if (|v62[373]| in v71) then v71[|v62[373]|] else v69] else v4, v4[v2 := v44.f20], v4 + v4, v72.cf18[v56 := v43], v4, if (cf29) then v4 else v4, if (fm1(cf29, v4, globalState)) then (fm6(cf32, globalState))[v56 := v44.f20] else v4, "vrxbpbgci", v4, v4, v4, "uci", v4, v4, DC12(v4).cf18, v4, v4, v4, "smf", v4, "hrnoni", "abfgyma"];
					var v74: set<seq<bool>> := {v58, [cf29, v0, true]};
					var v75 := DC24(v73);
					v63[881], globalState.f2, globalState.f3, v68[898], v73 := if ((v74 < v74) in v3) then v3[v74 < v74] else cf30, (0x32b - |v40|) % cf28, v0, v44.f20, v75.cf49;
					var v76: C2 := new C2();
					var v77 := DC22(v76);
					var v78 := DC0(v56);
					var v79: C0 := new C0(v78, cf30);
					v58, globalState.f10, v77, v79 := v58[v2 % v63[881] := true], cf32, v77, v79;
					var v80 := DC9(v79);
					v80 := v80;
				case DC12(cf18) =>
					v62[373] := v37;
					var v81: map<int, bool> := map[-0x299 := cf29];
					var v82: map<map<int, bool>, int> := map[v81 := v2];
					var v84: multiset<map<int, bool>> := multiset{v81};
					var v86 := DC28(map[v81 := -cf30]);
					var v88: array<map<map<int, bool>, int>> := new map<map<int, bool>, int>[27] [fm13(globalState), v82 + v82, v82, v82, v82, v82 + v82, (map v83 : map<int, bool> | v83 in v84 :: (v83) := (cf28))[v81 := |fm8(cf28, -0x34c, globalState)|], v82 + v82, v82, (map v85 : map<int, bool> | v85 in v84 :: (v85) := (cf30))[v81 := cf31], map[v81 := |"uavoowfky"|], v82, v82 + v82, v86.cf56[map[v56 := v0] := v2], v82, v82, v82, v82, v82 + v82, map[map v87 : int | v87 in v65 :: (v87 - cf30) := (cf32) := cf30], v82, v82, v82, v82, map[map[cf31 := cf32] := cf28], v82 + v82, v82 + v82];
					v88[509] := v82;
					v88[509], cf31 := (map[v81 := v2] + (map v89 : map<int, bool> | v89 in v82 :: (v89) := (|map[|map[v4 := -0x10d]| := |[false]|]|))) + v82, cf28;
					var v90 := DC6(-696, cf32, cf32, v56);
					v90 := v90;
					cf28 := v2 * |map[v1 := v63]|;
			}
			
			var v91: C2 := new C2();
			v91 := v91;
		case DC17(cf33, cf34, cf35, cf36) =>
			globalState.f2 := v2 * (-556 + v2);
			cf34 := false;
			var v92: seq<seq<int>> := [v37, v37, [|v4|], v37];
			globalState.f10 := !(v92[-0x3b0] < v37);
			globalState.f0 := true;
		case DC14(cf22) =>
			globalState.f0 := true;
			globalState.f9 := v2;
			var v93 := DC27(v2);
			if ((v2 + (v93.(cf55 := v2)).cf55) !in (map v94 : int | v94 in [v2] :: (v94 * v2) := (v2))) {
				var v95: seq<string> := [v4, v4, v4, v4, "w"];
				var v96: map<string, bool> := map[v95[v2] := !v0];
				v96 := v96;
				var v97, v98 := cf22.m3(v0, globalState);
				var v99 := DC0(-v97);
				var v100: T0 := new C0(v99, |v4|);
				v100 := v100;
				globalState.f0 := v0;
				var v101: seq<bool> := [v0];
				globalState.f3 := v101 <= [v0];
			} else {
				var v102: C2 := new C2();
				v102 := new C2();
				var v103 := new C1(cf22.f20);
				var v104, v105 := cf22.m3(!v0, globalState);
				globalState.f10 := v0;
				globalState.f9 := v2;
			}
			
			var v106: array<int> := new int[5](i4 => i4 + |[v0]|);
			v106[565] := 0xaa % v2;
			v106[565] := |v40|;
	}
	
	var v107 := DC25(v2);
	var v108: array<D9> := new D9[6] [v107, v107.(cf50 := v2), v107, if (v0) then v107 else v107, v107, v107];
	v108[481] := v107;
	v108[481] := v107.(cf50 := |(v4 + v4)|);
	if ((v2 / v2) != |(if (v0) then [v2] else seq(0x1c5, i5  => (-v2)))|) {
		var v109: map<char, bool> := map[v44.f20 := v0];
		v109, globalState.f0 := v109, v0;
		var v110: seq<string> := [v4 + v4, "k", v4];
		v37, globalState.f0, v110 := v37, if (if (v0) then v0 else false) then v0 else v0 && v0, v110;
		v4 := v4;
		var v111: map<int, char> := map[v2 := v43];
		var v112: seq<bool> := [v0, true, !v0, v0];
		v40 := fm14(v2 > v2, if (|v112| in v111) then v111[|v112|] else fm7(v44.f20, v0, globalState), v0, globalState);
		if (fm14(v0, 'c', !v0, globalState) != v40) {
			var v113 := DC4(v41);
			var v114 := DC7(v113);
			v114 := v114;
			var v115: array<map<char, bool>> := new map<char, bool>[6](i6 => v109);
			globalState.f9, globalState.f11, v115 := |v4|, v1, v115;
			v0, v43 := v0 <==> (v2 > 0x2d1), fm7(v44.f20, v0, globalState);
			var v116: map<bool, string> := map[fm1(fm1(v0, v4, globalState), fm6(v0, globalState), globalState) := v4 + "eph"];
			v116 := v116[v0 := "aaf"];
			globalState.f10 := !(if (v0 ==> v0) then v0 <==> v0 else v0);
		} else {
			globalState.f10 := v0;
			globalState.f2 := |fm8(v2, 873, globalState)| % v2;
			var v117: C2 := new C2();
			var v118: multiset<C2> := multiset{v117};
			globalState.f0 := (v118 - v118) >= multiset{v117};
			var v119: map<seq<int>, bool> := map[fm11(v2, v0, v2, globalState) := -v2 >= v2];
			v119 := v119;
			var v120: array<int> := new int[14](i7 => i7 - |{'n'}|);
			var v121 := DC15(|v4|, v2, v37, v2, {v2});
			var v122 := DC0(v121.cf23);
			var v123: C0 := new C0(v122, |v4|);
			var v124: map<C0, array<int>> := map[v123 := v120];
			var v125: array<array<int>> := new array<int>[25] [v120, if (v123 in v124) then v124[v123] else v120, v120, v120, v120, v120, v120, v120, v120, v120, v120, v120, v120, v120, v120, v120, v120, v120, v120, v120, v120, v120, v120, v120, v120];
			v125[587] := v120;
			v125[587] := v120;
		}
		
	} else {
		globalState.f3 := v0;
		var v126: map<seq<int>, multiset<bool>> := map[[v2, 681] := v39 - v39];
		var v127: set<int> := {v2};
		v126 := fm15(|v127| == v2, v3[true := -0xb2], v0, false, globalState);
		var v128: map<int, string> := map[v2 := DC12(v4).cf18];
		v4 := (if (v2 in v128) then v128[v2] else v4[v2 := v44.f20])[|"kr"| := v43];
		globalState.f3 := v0;
		var v129: array<array<bool>> := new array<bool>[27];
		v129[543] := v1;
		v129[543] := v1;
	}
	
	var v130: set<C1> := {v44, v44, v44};
	globalState.f3 := v44 in v130;
	var v131: set<int> := {|v4|};
	if ({v2} !! v131) {
		var v132, v133, v134, v135 := v44.m4(globalState);
		var v136 := DC0(v133);
		var v137: C0 := new C0(v136, v133);
		var v138: multiset<C0> := multiset{v137, v137, v137, v137};
		var v139 := DC8(v138);
		match (v139.(cf14 := v138)) {
			case DC8(cf14) =>
				var v140 := new C2();
				var v141: array<array<int>> := new array<int>[17];
				var v142: array<int> := new int[6] [v2, fm0(globalState), v2, v2, v2, |v4|];
				v141[718] := v142;
				v141[718] := v142;
				globalState.f2 := v133;
				v43 := v44.f20;
		}
		
		v1, globalState.f10 := v1, v0;
		globalState.f2 := v133;
		var v143: seq<bool> := [v134, v0];
		var v144: map<seq<bool>, bool> := map[v143 := v38.cf2];
		var v145: map<int, int> := map[|v144| := if (v132 in v3) then v3[v132] else v133];
		v133 := if (v133 in v145) then v145[v133] else v133;
	} else {
		var v146: array<int> := new int[22] [v2, v2, v2, v2, v2, v2, v2, v2, v2, v2, v2, v2, v2, v2, v2, |"ify"|, v2, v2, v2, v2, v2, v2];
		var v147 := DC10(v146);
		var v148 := DC11(v147);
		var v149: map<bool, D4> := map[v0 := v148.(cf17 := DC10(v146))];
		var v150: set<map<bool, D4>> := {v149, v149, map[v0 := v148.(cf17 := v147)], v149};
		match (DC18(v150 + {v149, v149, v149, v149})) {
			case DC19(cf38, cf39, cf40) =>
				var v151 := DC7(DC4(v41));
				var v152 := DC4(v41);
				v151 := DC7(v152);
				globalState.f9 := if (false) then v2 + v2 else |("mject" + v4)|;
				var v153 := DC0(v2);
				var v154 := new C0(v153, v2);
				var v155: array<string> := new string[17];
				var v156 := DC24(v155);
				var v157 := DC24(v156.cf49);
				var v158: seq<bool> := [cf39];
				globalState.f2, cf39, globalState.f9, v156 := 0x175, fm1(cf39, v4, globalState), fm0(globalState), if (!DC5(v0, |v158|, v1, v35).cf5) then if (cf39) then v156 else v156 else v157;
			case DC20(cf41, cf42, cf43) =>
				globalState.f0 := v0;
				v1[870] := v0;
				v1[870], cf43 := true, (cf43 + cf43) + -v2;
				var v159 := DC0(|(seq(0x44, i8  => (cf43)))|);
				var v160: map<int, bool> := map[cf43 := true];
				var v161: C0 := new C0(v159, |v160|);
				var v162 := DC9(v161);
				var v163: map<bool, C0> := map[v0 := v161];
				var v164: array<C0> := new C0[10] [v161, v161, v161, v162.cf15, v161, v161, v161, v161, v161, if (v1[870] in v163) then v163[v1[870]] else v161];
				v164[182] := v161;
				v164[182] := v161;
				var v165, v166 := m0(v1[870], v160, globalState);
			case DC18(cf37) =>
				v4, globalState.f9 := v4 + ([v44.f20, v44.f20] + v4), v2 * (v2 * v2);
				v146[356] := v2;
				var v167: array<D10> := new D10[25];
				var v168 := DC29();
				v167[109] := v168;
				v146[805] := fm0(globalState) / |v131|;
				var v169: C0 := new C0(DC0(|v37|), v2);
				var v170 := DC6(v2, true, v0, -|map[if (v0 in v5) then v5[v0] else true := v169]|);
				v146[356], globalState.f3, v167[109], globalState.f9, v146[805] := v2 + v2, v0, DC29(), fm0(globalState), v170.cf12;
				var v171: map<int, int> := map[v2 := v146[356]];
				v0 := if (!(|"hiwhxnyyd"| in v171)) then v0 else v0;
				var v172: map<int, bool> := map[v2 / |v4| := v0];
				v172 := v172[v146[356] - v146[356] := if (v0) then v0 else v0];
			case DC21(cf44) =>
				v1[987] := v2 != v2;
				v1[987] := v0;
				var v173 := DC0(-v2);
				var v174: multiset<char> := multiset{v44.f20, v43};
				var v175: T0 := new C0(v173, |v174|);
				var v176: map<int, T0> := map[v2 := v175];
				var v177: map<int, int> := map[|{true}| := |v176|];
				v146, v177 := v146, v177 + v177;
				v43 := v44.f20;
				var v178: array<bool> := new bool[22](i9 => true);
				var v179 := DC5(v0, v175.f18, v178, v35);
				var v180: set<D2> := {v179};
				var v181: map<set<D2>, int> := map[v180 - v180 := -v37[v175.f18]];
				v181 := v181[v180 := |(fm16(v175.f18, v1[987], 'b', v2, globalState) + v5)|];
		}
		
		var v182: seq<string> := [v4, v4];
		var v183: array<string> := new string[28] [v4, fm6(v0, globalState), seq(-0x2e7, i10  => (v44.f20)), v4, v4, v4, v4, fm6(v0, globalState), v4 + "l", v4 + v4, "veh", seq(0x27e, i11  => (v43)), v4, v4 + v4, v4, v4, v4[v2 := v44.f20], seq(489, i12  => (v44.f20)), v4, v182[-0x168], v4, v4, "psmshitr", v4[v2 := 'f'], seq(478, i13  => ('t')), v4 + (seq(-0x390, i14  => ('s'))), v4, v4];
		v183[306] := v4;
		v183[306] := seq(762, i15  => (v44.f20));
		v1[379] := v37 <= v37;
		v1[379] := v108[481].cf50 == v2;
		var v184 := DC31(v39);
		globalState.f2 := (v2 * |v184.cf58|) / v2;
		var v185 := new C1(v43);
	}
	
	var v186: array<set<int>> := new set<int>[27](i16 => v131);
	v186[599] := v131 - v131;
	v186[599] := v131;
	globalState.f0 := (v2 - v2) >= v2;
	v3 := v3 + (v3 + v3);
	var v187 := DC15(818, v2, v37, v2, {v2});
	for i17 := |v187.cf27| to v2 {
		globalState.f0, globalState.f0, v2 := v0, v0, i17;
		var v188: array<int> := new int[18](i18 => i18 * i17);
		v188[44] := v2;
		var v189 := DC10(v188);
		var v190: array<D4> := new D4[13] [v189, v189, v189, v189, v189, v189, v189, v189, v189.(cf16 := v188), v189, v189, DC10(v188), v189];
		var v191: map<int, int> := map[i17 := -v187.cf26];
		var v192: set<array<D4>> := {v190, v190};
		v188[44] := (DC20({v190, v190}, DC13(fm0(globalState), fm0(globalState), v0), |v191|).(cf43 := v2, cf41 := v192)).cf43;
		var v193 := DC0(i17);
		var v194: T0 := new C0(v193, v2 - -v2);
		var v195: map<bool, T0> := map[fm1(v0, "pdygc", globalState) := v194];
		v194 := if (v0 in v195) then v195[v0] else v194;
		var v196 := DC31(v39);
		var v197 := DC33(v196);
		var v198 := DC33(v197);
		v198 := v198;
	}
	forall i19 | 0 <= i19 < v1.Length {
		v1[i19] := v2 != (|map[v2 := false]| + |v4|);
	}
	var v199: set<string> := {v4, v4, "r"};
	v199 := v199;
	for i20 := v2 % v2 to v2 {
		globalState.f9 := v2;
		if (!v0) {
			var v200 := new C3(v4, DC0(0x279), v2);
			var v201 := DC0(i20);
			var v202: C0 := new C0(v201, v200.fm3(v0, i20, globalState));
			var v203: multiset<C0> := multiset{v202, v202, v202, v202};
			var v204: array<int> := new int[11] [v2, i20, i20, |v203|, v200.fm3(v0, v2, globalState), v2, i20, 834, v2, v2, -510];
			var v205 := DC10(v204);
			v205 := DC10(v204);
			var v206: multiset<array<bool>> := multiset{v1, v1, v1, v1, v1};
			v206 := v206;
			var v207, v208 := v200.m1(globalState);
			var v209: map<array<bool>, set<bool>> := map[v1 := v40];
			v207, v2, v200 := v207, |(v209 + v209)[v1 := v40 + v40]|, v200;
		} else {
			var v210: C2 := new C2();
			var v211: multiset<C2> := multiset{v210};
			var v212: set<multiset<C2>> := {v211, v211};
			globalState.f0 := v212 > (v212 - v212);
			var v213: map<C2, bool> := map[v210 := false];
			var v215: map<int, bool> := map[-v2 := !v0];
			var v216: set<map<int, bool>> := {v215, v215, v215, map[v2 := v0], v215};
			var v217: map<bool, D10> := map[if (v210 in v213) then v213[v210] else v0 := DC28(map v214 : map<int, bool> | v214 in v216 :: (v214) := (v2))];
			var v218: map<bool, map<bool, D10>> := map[v0 := v217];
			v218 := v218;
			var v219 := DC0(i20);
			var v220: T0 := new C0(v219, i20);
			v220 := new C0(v219, v220.f18);
			globalState.f2 := if (fm6(v0, globalState) != fm6(v38.cf2, globalState)) then v2 / v2 else v220.f18;
			v215 := v215[if (v0) then 0x39d else |[v0, v0, v0]| := v0];
		}
		
		globalState.f3 := fm11(i20, v0, i20, globalState) < [i20, i20];
		var v221: array<char> := new char[14];
		var v222: set<array<char>> := {v221, v221};
		var v223 := DC19(v222, v0, v0);
		var v224 := DC21(v223);
		v224 := v224;
	}
}