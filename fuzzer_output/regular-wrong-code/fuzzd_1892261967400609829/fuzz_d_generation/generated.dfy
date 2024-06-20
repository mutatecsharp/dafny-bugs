datatype D0 = DC1(cf1: bool) | DC2(cf2: int, cf3: int) | DC0(cf0: array<int>)
datatype D1 = DC3(cf4: char)
datatype D2 = DC5(cf6: bool, cf7: bool, cf8: bool, cf9: int, cf10: bool) | DC6 | DC4(cf5: map<int, int>) | DC7(cf11: D2)
datatype D3 = DC9(cf13: bool, cf14: map<int, int>) | DC8(cf12: set<map<int, int>>)
datatype D4 = DC11(cf16: map<int, D3>, cf17: array<bool>) | DC10(cf15: C0)
datatype D5 = DC13 | DC14(cf19: map<map<bool, bool>, seq<D0>>, cf20: int) | DC12(cf18: map<bool, int>) | DC15(cf21: D5)
datatype D6 = DC17 | DC16(cf22: set<int>)
datatype D7 = DC18(cf23: map<bool, bool>)
datatype D8 = DC20(cf25: bool, cf26: int) | DC19(cf24: map<int, char>) | DC21(cf27: D8)
class GlobalState {
	var f0 : map<set<int>, int>
	const f1 : set<int>
	const f2 : int
	var f3 : bool
	const f4 : int
	const f5 : char
	const f6 : map<int, seq<int>>
	const f7 : int
	const f8 : int
	var f9 : int
	const f10 : set<int>
	const f11 : int
	const f12 : bool
	const f13 : bool
	var f14 : bool
	const f15 : int
	const f16 : bool
	const f17 : array<char>
	const f18 : int
	const f19 : map<bool, bool>
	const f20 : array<string>
	var f21 : string
	var f22 : int
	var f23 : set<seq<int>>
	var f24 : bool
	var f25 : int
	constructor (f0 : map<set<int>, int>, f1 : set<int>, f2 : int, f3 : bool, f4 : int, f5 : char, f6 : map<int, seq<int>>, f7 : int, f8 : int, f9 : int, f10 : set<int>, f11 : int, f12 : bool, f13 : bool, f14 : bool, f15 : int, f16 : bool, f17 : array<char>, f18 : int, f19 : map<bool, bool>, f20 : array<string>, f21 : string, f22 : int, f23 : set<seq<int>>, f24 : bool, f25 : int) {
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
		this.f17 := f17;
		this.f18 := f18;
		this.f19 := f19;
		this.f20 := f20;
		this.f21 := f21;
		this.f22 := f22;
		this.f23 := f23;
		this.f24 := f24;
		this.f25 := f25;
	}
	
}

function fm0(p0: int, p1: bool, globalState: GlobalState): bool {
	if (multiset{map v0 : int | (0x126 <= v0) && (v0 < 982) :: (v0 + |map[true := |multiset(seq(0x225, i0  => (|map[false := !true]|)))|]|) := (|{false}|)} >= multiset{map[|multiset{true, false, false}| := |(seq(252, i1  => (|multiset{false, false}|)))|]}) then true else false
}
function fm1(p0: int, p1: bool, p2: string, p3: bool, globalState: GlobalState): int {
	|(map[false := false] + map[true := true])| % (|multiset{-146}| * -0x259)
}
function fm2(p0: int, p1: int, p2: bool, p3: int, globalState: GlobalState): string {
	"t"
}
function fm3(p0: int, p1: bool, globalState: GlobalState): set<bool> {
	{true, true, multiset{0x2b5, 0x2b1} < multiset{|multiset{false, false}|, |[true, true]|}}
}
function fm4(p0: D1, p1: seq<bool>, globalState: GlobalState): char {
	'p'
}
function fm6(p0: int, globalState: GlobalState): seq<bool> {
	[!(!true <==> false), !!true, [{DC20(false, 0x3c2).cf26, -0x320}] <= [{-845}], false, 0x30e == 0x389]
}
function fm7(p0: bool, globalState: GlobalState): map<bool, int> {
	map[false := 0x20a] + (map[!DC9(true, map[|"lihiauqtn"| := |{|[654]|, -0x31f}|]).cf13 := |map[0x277 := false]|] + map[false := |[0x2b2]|])
}
function fm8(globalState: GlobalState): seq<map<bool, int>> {
	([map[DC9(false, map[0x46 := 0x118]).cf13 := 299]] + (seq(-53, i0  => (map[false := -|multiset{true}|])))) + [map[!true := |{true}|]]
}
function fm9(p0: bool, p1: int, globalState: GlobalState): set<string> {
	set v0 : string | v0 in map[seq(-0x2a2, i0  => ('r')) := !true] :: (v0)
}
function fm10(p0: char, p1: int, globalState: GlobalState): map<int, D3> {
	map[-0x7a := DC9(false, map[|map[false := |{'w', 'k', 'e'}|]| := --32])]
}
function fm11(p0: bool, globalState: GlobalState): seq<int> {
	seq(0x14d, i0  => (|"hqgfpgnr"|))
}
function fm12(p0: seq<bool>, p1: bool, globalState: GlobalState): map<int, int> {
	(map[0x44 := 47] + map[|map[true := true]| := |[0x4f]|]) + (map[--498 := DC2(|map[[true] := true]|, |map[true := false]|).cf3] + (map v0 : int | (0x136 <= v0) && (v0 < -0x3b3) :: (v0 * |multiset{|multiset{true, false}|}|) := (-|"cjm"|)))
}
function fm13(globalState: GlobalState): D2 {
	DC4(map[|map[true := true]| := -|multiset{-|multiset{DC9(false, map[|multiset([|"iqbk"|])| := |[false, false, true, true, true]|])}|}|] + map[-0x7 := |"dnpmcdw"|])
}
function fm14(p0: seq<int>, p1: int, p2: int, p3: multiset<int>, globalState: GlobalState): map<int, char> {
	map[|map[false := 'y']| - 96 := 'o']
}
method m0(p0: bool, globalState: GlobalState) returns (r0: seq<bool>, r1: int, r2: array<int>, r3: char) {
	var i0 := 0;
	while (p0)
		decreases 100 - i0
	{
		if (i0 >= 100) {
			break;
		}
		
		i0 := i0 + 1;
		var v0 := new C0();
		globalState.f3 := !p0;
		var v1: seq<bool> := [p0];
		globalState.f3 := if (p0) then p0 else multiset(v1) < multiset{p0};
		var v2: set<string> := {"jbq" + "jjrtckkdd"};
		var v3 := 0x15a;
		v2 := fm9(p0, v3, globalState);
	}
	var v4 := 346;
	var v5 := DC2(v4, v4);
	var v6: seq<int> := [|"dfufllfmk"|, v4];
	globalState.f22 := v5.cf2 - |v6|;
	var v7 := DC6();
	v7 := v7;
	match (DC20(p0, v4)) {
		case DC20(cf25, cf26) =>
			var v8: array<bool> := new bool[18](i1 => p0);
			var v9: map<int, int> := map[v4 := v4];
			var v10: set<map<int, int>> := {v9, v9[|[cf26]| := cf26], v9, v9, v9};
			v8[554] := v10 == v10;
			v8[554] := p0;
			var v11 := 'l';
			var v12: multiset<int> := multiset{v4, v4, cf26, cf26};
			var v13: map<char, int> := map[v11 := cf26 * |v12|];
			v13 := v13[v11 := if (p0) then v4 else v4];
			var v14: map<bool, int> := map[v4 != 0x248 := cf26 % v4];
			v14 := v14 + map[v8[554] := cf26];
			var v15: map<bool, bool> := map[p0 := v8[554]];
			v15 := v15[v8[554] := p0];
		case DC19(cf24) =>
			globalState.f14 := if (false) then p0 else if (p0) then p0 else p0;
			var v16 := 'a';
			var v17: map<char, char> := map[v16 := 'k'];
			var v18: set<bool> := {p0};
			if (v4 <= (|v17| % |v18|)) {
				var v19 := new C0();
				var v20 := "ntnbowsn";
				globalState.f21 := v20;
				var v21 := new C0();
				v16 := v16;
				globalState.f22 := v4;
			} else {
				globalState.f14 := p0;
				var v22: multiset<bool> := multiset{p0};
				v22 := v22;
				globalState.f25 := v4 * -892;
				var v23: array<C0> := new C0[21];
				var v24: C0 := new C0();
				var v25: map<C0, array<C0>> := map[v24 := v23];
				var v26: seq<array<C0>> := [v23, if (v24 in v25) then v25[v24] else v23];
				globalState.f3 := |(v26 + [v23])| <= v4;
				var v27: map<bool, bool> := map[p0 := p0];
				var v28 := DC1(p0);
				var v29: seq<D0> := [v28];
				var v30: map<map<bool, bool>, seq<D0>> := map[v27 := v29];
				var v31: map<int, int> := map[DC14(v30, v4).cf20 := 309];
				var v32 := DC9(p0, v31);
				var v33: map<int, D3> := map[|(seq(357, i2  => (v16)))| := v32];
				var v34: array<bool> := new bool[24];
				var v35 := DC11(v33, v34);
				v35 := v35;
			}
			
			var v36 := "llmnukb";
			var v37: map<char, string> := map[v16 := v36];
			match (DC1(fm0(|(if (v16 in v37) then v37[v16] else v36)|, p0, globalState))) {
				case DC1(cf1) =>
					globalState.f24 := true;
					globalState.f9 := (if (cf1) then v4 else |"iwdi"|) * v4;
					var v38: seq<bool> := [cf1];
					var v39: seq<seq<bool>> := [v38 + [p0, p0]];
					r0 := v39[|"bkf"|];
					var v40: map<int, int> := map[|map[v4 := v38]| := 177];
					var v41 := DC9(fm0(v4, cf1, globalState), v40 + v40);
					v41 := v41;
				case DC2(cf2, cf3) =>
					var v42 := new C0();
					var v43: set<char> := {v16, v16, v16};
					var v44: array<set<char>> := new set<char>[7] [v43, v43, {v16, v16, v16, v16}, v43, if (!p0) then v43 else v43, v43, v43];
					var v45: array<int> := new int[16](i3 => i3 * |v36|);
					v45[389] := 0x227;
					var v46: array<bool> := new bool[5](i4 => p0);
					var v47 := DC11(fm10(v16, |v6|, globalState), v46);
					var v48: map<D4, bool> := map[v47 := p0];
					var v49: array<map<D4, bool>> := new map<D4, bool>[9] [v48, v48, map[v47 := true] + map[v47 := p0], map[v47 := p0], v48 + (map[v47 := p0])[v47 := p0], v48, v48, v48 + v48, map[v47 := p0]];
					v49[150] := v48;
					v44, v45[389], globalState.f24, v49[150], globalState.f3 := v44, cf3, !!(if (p0) then p0 else p0), v48 + (v48 + v48), p0;
					var v50: set<array<bool>> := {v46, v46};
					v50, r3, globalState.f9 := v50 * (if (p0) then v50 else {v46}), v16, v45[389];
					var v51: map<bool, int> := map[p0 := cf2];
					var v52: map<map<bool, int>, bool> := map[v51 := if (p0) then !p0 else !p0];
					v52 := v52[v51 := p0];
				case DC0(cf0) =>
					var v53 := new C0();
					globalState.f21 := "ukla";
					var v54: seq<seq<int>> := [fm11(p0, globalState), fm11(p0, globalState), v6];
					r1 := |fm2(v53.fm5(globalState), |v6|, p0, v4, globalState)| * |v54|;
					globalState.f9 := v4;
			}
			
			var v55: C0 := new C0();
			var v56: array<C0> := new C0[1] [v55];
			v56[359] := v55;
			v56[359] := v55;
		case DC21(cf27) =>
			var v57 := new C0();
			r1 := 0x314;
			var v58: seq<bool> := [!p0];
			var v59: map<seq<bool>, bool> := map[v58 := p0];
			v59 := v59;
			var v60: array<char> := new char[8];
			var v61 := 'x';
			v60[970] := if (p0) then v61 else v61;
			var v62: map<bool, bool> := map[p0 := p0];
			var v63 := DC1(p0);
			var v64: seq<D0> := [v63, v63];
			var v65: map<map<bool, bool>, seq<D0>> := map[v62 := v64];
			var v66 := DC14(v65[v62 := v64], v4);
			r1, globalState.f24, v60[970], globalState.f25 := v66.cf20, !!p0, 'c', -v4 + v4;
	}
	
	var v67: C0 := new C0();
	var v68: map<C0, int> := map[v67 := v67.fm5(globalState)];
	var v69: map<int, int> := map[if (v67 in v68) then v68[v67] else 0x56 := v4];
	v4, v69 := v4, v69 + (v69 + v69);
	for i5 := v4 to v4 {
		var v70: array<bool> := new bool[2] [p0, fm0(v4, true, globalState)];
		v70[919] := false;
		var v71: multiset<bool> := multiset{p0};
		v70[919] := fm0(i5, 0x1a1 == (if (p0 in v71) then v71[p0] else i5), globalState);
		globalState.f14 := v70[919];
		var v72: seq<bool> := [p0, !p0, p0, p0];
		var v73 := "fjmbsow";
		globalState.f9 := (v5.cf2 - |v72|) * |((seq(997, i6  => ('v'))) + v73)|;
		var v74 := DC4(fm12(v72, v70[919], globalState));
		v74 := fm13(globalState);
	}
	var v75: seq<bool> := [p0, !p0, false, false];
	var v76: seq<seq<bool>> := [v75];
	var v77: multiset<int> := multiset{v4, v4};
	r0 := v76[if (-v4 in v77) then v77[-v4] else |fm14(seq(0x31e, i7  => (v4)), v4, v4, v77, globalState)|];
	r1 := v4;
	var v78: array<int> := new int[8](i8 => i8 % v4);
	var v79: seq<array<int>> := [v78, v78, v78, v78];
	r2 := v79[v4];
	var v80 := 'f';
	r3 := v80;
}
class C0 {
	constructor () {
	}
	
	function fm5(globalState: GlobalState): int {
		|"rpbp"| - (0x2c0 * 0x203)
	}
}

method Main() {
	var v0 := 0x2b3;
	var v1: set<int> := {v0};
	var v2: map<set<int>, int> := map[v1 := v0];
	var v3: seq<int> := [v0, v0];
	var v4: map<int, seq<int>> := map[|{v0}| := v3];
	var v5: array<int> := new int[24];
	var v6 := DC0(v5);
	var v7: array<array<int>> := new array<int>[4] [v5, v5, v5, v6.cf0];
	var v8: multiset<int> := multiset{|v3|};
	var v9 := "px";
	var v10: map<array<array<int>>, set<int>> := map[v7 := {|v8|, v0, |v9|, v0}];
	var v11: array<char> := new char[23](i0 => DC3('j').cf4);
	var v12 := true;
	var v13: map<bool, bool> := map[v12 := v12];
	var v14: array<string> := new string[2] ["eqobbdls", "rgyeuu"];
	var v15: set<seq<int>> := {v3};
	var globalState := new GlobalState(v2 + v2, v1, 0x1bd, false, 988, 'l', v4, 496, -0xe9, 0x3aa, if (v7 in v10) then v10[v7] else v1, 987, false, true, false, 99, false, v11, 0xa1, v13, v14, "txxtamr", 0x241, v15, false, 0xb5);
	var v16 := 'j';
	var v17: map<map<bool, int>, char> := map[map[v12 := v0] := v16];
	v17 := v17;
	var i1 := 0;
	while (v12)
		decreases 100 - i1
	{
		if (i1 >= 100) {
			break;
		}
		
		i1 := i1 + 1;
		var v18: seq<bool> := [v12];
		globalState.f14 := v18[v0];
		var v19 := DC1(v12);
		var v20: map<D0, int> := map[v19 := |v8|];
		var v21: map<int, bool> := map[v0 := v12];
		var v22: set<bool> := {v12};
		var v23: map<bool, int> := map[if (|v22| in v21) then v21[|v22|] else v12 := v0];
		v20 := v20[v19 := |v23|];
		v5[874] := |v9|;
		var v24: map<bool, char> := map[v18[|v18|] := v16];
		var v25: map<int, array<int>> := map[|v24| := v5];
		v5[874] := v0 + |(v25 + v25)|;
		var v26: map<D1, bool> := map[DC3(v16) := v12];
		var v27 := DC3(v16);
		v26 := v26[v27 := v12];
	}
	var v28, v29, v30, v31 := m0(fm0(|"gb"|, v12, globalState), globalState);
	globalState.f22 := -v29;
	if (v12) {
		globalState.f14 := v12;
		var v32 := DC1(v12);
		match (v32) {
			case DC1(cf1) =>
				var v33: array<bool> := new bool[8] [v29 != v29, true, v12, v12, v1 !! v1, cf1, 0x1ec == --v0, cf1];
				v33 := v33;
				globalState.f14 := !v12 || cf1;
				globalState.f22 := v29;
				v5[176] := v29;
				v5[176] := 0x372;
			case DC2(cf2, cf3) =>
				var v34: multiset<bool> := multiset{v12, !v12};
				var v36: map<int, int> := map[cf2 := v0];
				var v37: map<bool, int> := map[v12 := fm1(716, v28[cf2], v9, v12, globalState)];
				var v40: array<map<int, int>> := new map<int, int>[19] [map[|v34| := cf3] + (map v35 : int | (0x2cd <= v35) && (v35 < 364) :: (v35 * |map[map[v12 := cf2] := true]|) := (|multiset{cf2, -v0}|)), v36, v36, v36, map[cf2 := v29] + v36, v36, v36, v36[|v37| := cf3], v36, if (v12) then v36 else DC4(v36).cf5, v36, map v38 : int | (-473 <= v38) && (v38 < 31) :: (v38 + -v29) := (-cf2), map[v29 := cf2] + v36, v36 + v36, v36, v36, v36, v36, (map v39 : int | v39 in v36[|v28| := cf3] :: (v39 % 0x3de) := (|v28|))[cf2 := --0x3e2]];
				var v41: map<int, D0> := map[v29 := v32];
				v30[329] := fm1(|v41|, v12, "vjvcgp", v12, globalState);
				v29, globalState.f22, v40, globalState.f14, v30[329] := cf2, cf2, v40, !v12, v0;
				var v42: set<map<int, int>> := {v36 + v36[0xc := v30[329]]};
				var v43: map<map<int, int>, bool> := map[v36 := v12];
				v42 := (DC8({v36}).(cf12 := v42).(cf12 := v42)).cf12 * (v42 * (set v44 : map<int, int> | v44 in v43 :: (v44)));
				globalState.f22 := fm1(|multiset{true}| % --v29, v12, fm2(v29, v0, v28[|(seq(0x3a5, i2  => (v29)))|], -cf3, globalState), v12, globalState);
				var v45: set<bool> := {true, v12};
				v45 := (v45 - fm3(v30[329], v12, globalState)) + ({v12, v12} - v45);
			case DC0(cf0) =>
				globalState.f25 := |v1|;
				globalState.f9 := v0;
				var v46, v47, v48, v49 := m0(!v12, globalState);
				var v50: multiset<bool> := multiset{v12, v12, v12, v12, v12};
				var v51 := DC3(v16);
				v47, v49 := |v50| - --(v29 / |v3|), fm4(v51, [true], globalState);
		}
		
		var v52: map<seq<int>, int> := map[[v0] := v0];
		var v54: seq<seq<int>> := [v3, v3[v0 := v0]];
		v52 := v52 + ((map v53 : seq<int> | v53 in v54 :: (v53) := (v0)) + v52);
		var v55: multiset<bool> := multiset{v12, v12, true, v12};
		var v56 := DC2(0x26f, fm1(v29, v12, v9, v12, globalState));
		match (DC5(v1 > v1, (if (false in v55) then v55[false] else v56.cf2) < v29, true, v29, fm1(v29, false, fm2(v29, v29, v12, v0, globalState), v12, globalState) < v0)) {
			case DC5(cf6, cf7, cf8, cf9, cf10) =>
				var v57: set<set<int>> := {v1};
				cf7 := v57 >= (v57 * v57);
				v29, v5, globalState.f21, globalState.f21 := --(cf9 * v29), if (v12) then v5 else if (cf8) then v5 else v5, v9, fm2(cf9, v0, cf7, |v28|, globalState) + v9;
				v5[193] := v0;
				v5[193] := v29;
				var v58, v59, v60, v61 := m0(!v12, globalState);
			case DC6() =>
				var v62 := new C0();
				var v63: multiset<array<char>> := multiset{v11};
				var v64: map<multiset<array<char>>, C0> := map[v63 := v62];
				var v65: seq<C0> := [v62];
				var v66: map<bool, C0> := map[v12 := v62];
				var v67: array<C0> := new C0[26] [v62, v62, v62, v62, v62, v62, v62, v62, if (fm0(v0, v12, globalState)) then v62 else v62, if (v63 in v64) then v64[v63] else v65[v0], v62, v62, v62, v62, v62, v62, v62, v62, v62, v62, if (fm0(v0, v12, globalState) in v66) then v66[fm0(v0, v12, globalState)] else v62, if (v12 in v66) then v66[v12] else v62, v62, v62, v62, v62];
				v67[521] := v62;
				v67[521] := if (v12) then v62 else v62;
				var v68 := new C0();
				var v69, v70, v71, v72 := m0(v12, globalState);
			case DC4(cf5) =>
				var v73: C0 := new C0();
				var v74: map<C0, C0> := map[v73 := v73];
				var v75 := DC10(v73);
				v73, globalState.f22 := if (v75.cf15 in v74) then v74[v75.cf15] else v73, if ((if (v12 in v13) then v13[v12] else v12) in v55) then v55[if (v12 in v13) then v13[v12] else v12] else v0 / v29;
				var v76: array<set<bool>> := new set<bool>[21];
				globalState.f25, globalState.f14, v31, globalState.f14, v76 := v29, v3[v0] >= -0x127, v31, !v12, v76;
				globalState.f21 := v9;
				var v77, v78, v79, v80 := m0(0x34d > v29, globalState);
			case DC7(cf11) =>
				var v81: array<bool> := new bool[11];
				v81[634] := v32.cf1;
				var v82: seq<multiset<bool>> := [multiset(fm6(v0, globalState))];
				v81[634] := (if (v12) then v55 else multiset{v12, v12}) >= (v82[v29] * v55[v12 := -v0]);
				globalState.f24 := if (v12) then v8 > multiset{v29} else v12;
				var v83: map<int, int> := map[|v3| := v29];
				var v84: set<bool> := {v12};
				var v85: set<map<int, int>> := {v83, map[|v9| := v29], map[|v84| := v29], v83, v83};
				var v86 := DC8(v85);
				v86 := v86;
				globalState.f14 := v81[634];
		}
		
		var v87, v88, v89, v90 := m0(v12, globalState);
	} else {
		match (DC3(v31)) {
			case DC3(cf4) =>
				var v91: set<bool> := {v12};
				v91 := v91;
				var v92: array<bool> := new bool[27];
				v92[263] := v12;
				var v93: array<seq<array<int>>> := new seq<array<int>>[2];
				var v94: seq<array<int>> := [v5, v30];
				v93[48] := v94;
				var v95: map<bool, int> := map[!v12 := |v28|];
				var v96: map<map<int, bool>, array<int>> := map[map[|v95| := v12] := v30];
				var v97: map<int, bool> := map[v29 := v12];
				v30, globalState.f14, v92[263], globalState.f22, v93[48] := if (v97 in v96) then v96[v97] else if (v12) then v5 else v30, v12, !v12, v3[v29], v94 + v94;
				var v98 := new C0();
				var v99, v100, v101, v102 := m0(fm0(v0, v92[263], globalState), globalState);
		}
		
		v9 := ("aml" + v9) + (v9 + v9);
		globalState.f14 := if (v29 != v0) then v12 else if (v12) then v12 else v12;
		globalState.f24 := v12;
		var v103: map<int, string> := map[v29 := v9];
		v103 := v103[|map[v0 := !false]| := v9];
	}
	
	forall i3 | 0 <= i3 < v5.Length {
		v5[i3] := i3 + 0x6e;
	}
	var v104: map<int, char> := map[-v29 := v31];
	globalState.f24 := (if (v0 in v104) then v104[v0] else 'g') in v9[-v0 := v31];
	var i4 := 0;
	while (v12)
		decreases 100 - i4
	{
		if (i4 >= 100) {
			break;
		}
		
		i4 := i4 + 1;
		var v105: seq<array<int>> := [v5, v5, v5];
		v30[236] := |[v28[v29]]| - |v105|;
		var v106: seq<map<bool, bool>> := [v13];
		var v107: map<int, bool> := map[|v106| := v12];
		var v108: map<map<int, bool>, int> := map[v107 := v0];
		var v111: seq<map<int, bool>> := [v107, v107, map v110 : int | v110 in multiset{v29, v0} :: (v110 + v0) := (v12)];
		globalState.f9, v30[236] := |((v108 + (map v109 : map<int, bool> | v109 in v111 :: (v109) := (v29))) + v108)|, v29;
		var v112: set<bool> := {v12, v12};
		var v113: map<map<int, char>, bool> := map[map[v29 := v16] := true];
		globalState.f9 := |map[-fm1(|v112|, v12, v9, v12, globalState) % |v113| := v30[236]]|;
		v30[236] := v29 % |v107|;
		globalState.f22 := -v30[236];
	}
	globalState.f25 := v29;
	for i5 := |v9| to 63 {
		var v114: map<int, int> := map[i5 := v29];
		var v115: set<map<int, int>> := {v114};
		var v116 := DC8(v115);
		match (v116.(cf12 := {v114, v114} + v115)) {
			case DC9(cf13, cf14) =>
				var v117: map<seq<bool>, int> := map[v28 + v28 := v29];
				v117 := v117[v28 := -i5];
				var v118: array<bool> := new bool[16] [cf13, v12, v12, cf13, cf13, cf13, v12, cf13, cf13, v12, v12, v12, v12, v12, cf13, cf13];
				var v119: map<int, array<bool>> := map[i5 := v118];
				globalState.f3 := (if (!cf13) then v119 else v119) != v119;
				v30[920] := fm1(v29, v12, v9, cf13, globalState) % -v29;
				v30[920] := v0;
				v28 := fm6(|v9|, globalState);
			case DC8(cf12) =>
				var v120: map<bool, int> := map[v12 := 875];
				v120 := v120[true := if (v12) then |[v0]| else v29];
				globalState.f3 := v12;
				globalState.f9 := v29;
				var v121 := new C0();
		}
		
		if (false) {
			var v122, v123, v124, v125 := m0(!v12, globalState);
			var v126 := DC4(v114);
			var v127 := DC9(v12, v126.cf5);
			var v128: seq<map<int, int>> := [v114[v0 := |v9|]];
			var v129: map<int, D3> := map[v0 := DC9(v127.cf13, v128[fm1(v0, v12, seq(0x19a, i6  => (v125)), v12, globalState)])];
			var v130: map<map<int, D3>, char> := map[v129 + v129 := v31];
			v130 := v130[v129 := v16];
			var v131: array<multiset<seq<int>>> := new multiset<seq<int>>[24];
			var v132: seq<seq<int>> := [v3, v3, [v29, -i5]];
			v131[194] := multiset(v132);
			v131[194] := multiset{[i5, 683, |v3|]};
			globalState.f22 := fm1(v29, v12, seq(0x327, i7  => (v125)), fm0(i5, v12, globalState), globalState) + 0xe4;
			var v133 := new C0();
		} else {
			var v134: array<array<array<bool>>> := new array<array<bool>>[29];
			var v135: array<bool> := new bool[19](i8 => true);
			var v136: array<array<bool>> := new array<bool>[12] [v135, v135, v135, v135, v135, v135, v135, v135, v135, v135, v135, v135];
			v134[944] := v136;
			var v137: map<array<int>, array<int>> := map[v5 := v5];
			v12, v134[944], globalState.f25 := if (!true) then v12 && fm0(|v137|, v12, globalState) else v12, v136, v0;
			var v138 := new C0();
			v5[854] := 0x3c4;
			var v139: map<C0, C0> := map[v138 := v138];
			globalState.f24, v138, globalState.f22, globalState.f25, v5[854] := fm0(v29, v12, globalState), if (v138 in v139) then v139[v138] else v138, v29, v3[v0], -v0;
			var v140 := DC10(v138);
			var v141: map<D4, int> := map[v140 := v5[854]];
			v141 := v141[v140 := v5[854]];
			var v142 := new C0();
		}
		
		var v143: array<map<bool, int>> := new map<bool, int>[19];
		var v144: map<bool, int> := map[!v12 := v29];
		v143[937] := v144;
		var v145 := DC12(v144);
		v143[937] := (v145.(cf18 := v144).(cf18 := map[v12 := i5]).(cf18 := v144)).cf18;
		var v146: multiset<bool> := multiset{v12, v12};
		var v147: seq<multiset<bool>> := [multiset{true}, v146, v146, v146];
		globalState.f14 := if (v8 <= v8) then v12 else v147[v29] !! v146;
	}
	var v148: C0 := new C0();
	var v149: seq<C0> := [v148, v148];
	var v150: map<int, string> := map[v0 := v9];
	v148 := v149[fm1(v29, v12, if (v0 in v150) then v150[v0] else seq(-0xd1, i9  => (v16)), v12, globalState)];
	v5[998] := -v29;
	v5[998], v16, globalState.f3, globalState.f24 := v0, v31, v0 <= v0, v12;
	var v151: map<int, bool> := map[v5[998] := v12];
	var v152: multiset<bool> := multiset{if (v29 in v151) then v151[v29] else !v12};
	if (v152 > multiset(v28 + v28)) {
		globalState.f24 := true;
		globalState.f22 := v5[998];
		var v153: map<int, int> := map[v5[998] := v5[998]];
		var v154: map<int, D3> := map[v0 := DC9(false, v153)];
		var v155: array<bool> := new bool[15];
		var v156 := DC11(v154, v155);
		var v157: map<seq<bool>, D4> := map[[v12] := v156];
		var v158: map<int, int> := map[v0 := |v157|];
		match (DC4(v158 + v153)) {
			case DC5(cf6, cf7, cf8, cf9, cf10) =>
				var v159, v160, v161, v162 := m0(fm0(v5[998], cf6, globalState), globalState);
				var v163 := DC16(v1);
				var v164 := DC18(v13);
				v1, globalState.f24, v149, globalState.f9 := (v163.(cf22 := v1)).cf22, v12, v149, |v164.cf23|;
				globalState.f14 := v12;
				var v165 := new C0();
			case DC6() =>
				var v166, v167, v168, v169 := m0(v12, globalState);
				v155[756] := v12;
				v155[756] := v29 <= |v1|;
				v28 := v166;
				globalState.f9 := 0x66;
			case DC4(cf5) =>
				globalState.f21 := v9;
				globalState.f24 := (v29 / v5[998]) > |v28|;
				var v170: array<map<int, char>> := new map<int, char>[1];
				var v171: map<C0, map<int, char>> := map[v148 := v104];
				var v172 := DC19(if (v148 in v171) then v171[v148] else v104);
				v170[224] := v172.cf24 + v104;
				var v173: seq<map<int, char>> := [v104[v5[998] := v31], v104, v104];
				v170[224] := v104 + v173[v5[998]];
				var v174: map<seq<bool>, seq<int>> := map[v28 := v3 + v3];
				v174 := v174[v28 + v28 := v3 + v3];
			case DC7(cf11) =>
				var v175: seq<map<bool, int>> := [fm7(v12, globalState)];
				globalState.f14 := v175 <= (fm8(globalState) + v175);
				var v176, v177, v178, v179 := m0(v12, globalState);
				v31 := fm4(DC3(v179), v28, globalState);
				globalState.f25, globalState.f24 := -0x79 * v5[998], v12;
		}
		
		globalState.f14 := !v12;
		if (v152 >= multiset(v28)) {
			v31 := v31;
			v31 := if (false) then v16 else v31;
			v5 := v30;
			var v180 := new C0();
			var v181: set<bool> := {true, v29 in v3, v12, v0 != |v28|, v12 <==> v12};
			v5[998], v5, v181, globalState.f25, globalState.f24 := if (v29 in v158) then v158[v29] else v5[998], v5, (v181 + v181) + v181, v0, v12;
		} else {
			var v182 := new C0();
			globalState.f25 := -v29;
			globalState.f25 := |([v12, v12] + v28)|;
			globalState.f25 := v29;
			var v183 := new C0();
		}
		
	} else {
		var v185: map<int, int> := map[v29 := |(map v184 : int | v184 in v3 :: (v184 + v0) := (|v8|))|];
		globalState.f25, globalState.f3, v1 := -((|v185| + v5[998]) - (v29 / v29)), v12, v1;
		var v186, v187, v188, v189 := m0(v12, globalState);
		var v190, v191, v192, v193 := m0(v12 <== v12, globalState);
		var v194: array<C0> := new C0[10];
		v194[53] := v148;
		var v195: array<bool> := new bool[14] [v12, v12, if (v12) then v12 else v12, !v12 && v12, true, !v12, v12, v12, v12, v12, v12, if (v12) then v12 else v12, v29 != v29, v12];
		v195[198] := v12;
		var v196: set<string> := {v9, v9, "bxdpjq", fm2(v5[998], v29, false, v5[998], globalState), v9};
		v194[53], v11, v3, globalState.f3, v195[198] := v148, v11, v3, ("b" != (seq(0x30d, i10  => (v193)))) in (v190 + v190), v196 > v196;
		if (false) {
			globalState.f9 := -v187;
			globalState.f25 := v148.fm5(globalState);
			var v197, v198, v199, v200 := m0(v12, globalState);
			var v201 := new C0();
			var v202: map<bool, int> := map[v195[198] := v0];
			v202 := v202;
		} else {
			globalState.f25 := |v190|;
			globalState.f24 := (v191 + -v29) >= |v3|;
			var v203 := new C0();
			globalState.f14 := v195[198];
			globalState.f24 := v195[198];
		}
		
	}
	
	var v204: set<bool> := {v28[|v9|], v12, v12};
	if ((v204 - v204) !! v204) {
		var v205: array<bool> := new bool[17];
		v205[644] := v12;
		var v206: map<int, int> := map[|v151| := -0xd4];
		var v207 := DC4(v206);
		v205[644] := fm0(|(fm2(v5[998], v5[998], v12, v29, globalState) + v9)|, v5[998] !in v207.cf5, globalState);
		globalState.f22 := -v5[998];
		var v208, v209, v210, v211 := m0(v205[644], globalState);
		v30 := v210;
		globalState.f24 := v205[644];
	} else {
		if (v28[v29] <== (|v28| != v5[998])) {
			var v212: array<array<array<int>>> := new array<array<int>>[19] [v7, v7, v7, v7, if (v12) then v7 else v7, v7, v7, v7, v7, v7, v7, v7, v7, if (v12) then v7 else v7, v7, v7, v7, v7, v7];
			v212[637] := v7;
			v212[637] := v7;
			globalState.f24 := v12 && (v12 ==> v12);
			var v213 := DC17();
			globalState.f21, globalState.f21, v213 := v9, "smrlm", v213;
			v30 := v30;
			globalState.f24 := fm0(v29, fm0(v0, v12, globalState), globalState);
		} else {
			var v214, v215, v216, v217 := m0(v12, globalState);
			globalState.f14 := v12;
			var v218: map<array<int>, bool> := map[v216 := false];
			v218 := v218 + v218;
			var v219, v220, v221, v222 := m0(v12, globalState);
			globalState.f21 := v9;
		}
		
		var v223, v224, v225, v226 := m0(v204 !! v204, globalState);
		var v227, v228, v229, v230 := m0(v12, globalState);
		globalState.f24 := v12;
		globalState.f14 := v12;
	}
	
	globalState.f9 := v5[998];
	v0 := v148.fm5(globalState);
}