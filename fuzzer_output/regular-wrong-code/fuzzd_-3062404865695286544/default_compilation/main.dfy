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
datatype D0 = DC1(cf1: int, cf2: set<bool>, cf3: int, cf4: array<seq<bool>>, cf5: char) | DC2(cf6: map<int, int>, cf7: map<bool, bool>, cf8: array<bool>) | DC3(cf9: int) | DC0(cf0: array<bool>)
datatype D1 = DC5(cf11: int, cf12: map<int, int>, cf13: int) | DC6(cf14: char, cf15: char) | DC7(cf16: char) | DC4(cf10: bool)
datatype D2 = DC9(cf18: char) | DC10(cf19: bool, cf20: bool) | DC11(cf21: int, cf22: int) | DC8(cf17: C1) | DC12(cf23: D2)
datatype D3 = DC14(cf25: C1, cf26: char, cf27: C1, cf28: C0, cf29: bool) | DC13(cf24: array<char>) | DC15(cf30: D3)
datatype D4 = DC17(cf32: int, cf33: bool, cf34: int, cf35: int, cf36: C1) | DC16(cf31: seq<string>)
datatype D5 = DC19(cf38: bool, cf39: bool, cf40: bool) | DC18(cf37: array<int>) | DC20(cf41: D5)
datatype D6 = DC22(cf43: bool) | DC23(cf44: int, cf45: seq<int>, cf46: int) | DC21(cf42: set<multiset<bool>>) | DC24(cf47: D6)
class GlobalState {
	var f0 : int
	var f1 : int
	const f2 : bool
	const f3 : int
	const f4 : int
	var f5 : bool
	constructor (f0 : int, f1 : int, f2 : bool, f3 : int, f4 : int, f5 : bool) {
		this.f0 := f0;
		this.f1 := f1;
		this.f2 := f2;
		this.f3 := f3;
		this.f4 := f4;
		this.f5 := f5;
	}
	
}

class C0 {
	var f8 : map<int, int>
	const f9 : bool
	constructor (f8 : map<int, int>, f9 : bool) {
		this.f8 := f8;
		this.f9 := f9;
	}
	
	function fm3(p0: char, p1: char, p2: int, globalState: GlobalState): int {
		--0x98 - (0x353 * |map[f9 := 0x3ac]|)
	}
	function fm4(globalState: GlobalState): bool {
		(multiset{[true], [f9], [f9, f9]} + multiset{[false, f9], [f9]}) !! (multiset{[f9], [f9]} - multiset{[f9], [f9], [f9, false, f9]})
	}
}

class C1 {
	const f6 : char
	const f7 : bool
	constructor (f6 : char, f7 : bool) {
		this.f6 := f6;
		this.f7 := f7;
	}
	
	function fm0(p0: string, p1: char, globalState: GlobalState): string {
		(seq(abs(-0xbc), i0  => (f6))) + ("qkbbvoi" + "dcitg")
	}
	method m0(p0: int, p1: int, globalState: GlobalState) returns (r0: int, r1: int, r2: array<int>) {
		if (f7) {
			var v0: array<map<int, seq<bool>>> := new map<int, seq<bool>>[21](i0 => if (f7) then map[p1 := [f7]] else map[p1 := [f7, f7]]);
			var v1: seq<bool> := [f7];
			var v2: map<int, seq<bool>> := map[p0 := v1];
			v0[safeIndex(572, v0.Length)] := v2;
			var v3: array<int> := new int[5];
			v3[safeIndex(365, v3.Length)] := p1;
			globalState.f1, v0[safeIndex(572, v0.Length)], r1, v3[safeIndex(365, v3.Length)] := safeModuloInt(p1, |v1|), if (fm1(globalState)) then map[p1 := v1] else v2 + v2, 0x30, 0x2aa;
			globalState.f5 := if (false) then f7 else f7;
			var v4 := "cfkjko";
			var v5: map<string, int> := map[v4 := v3[safeIndex(365, v3.Length)]];
			v5 := v5[v4 := v3[safeIndex(365, v3.Length)]];
			var v6: map<bool, int> := map[f7 := p1];
			var v7: map<string, map<bool, bool>> := map[v4 := map[f7 := f7]];
			var v8: map<bool, bool> := map[!f7 := f7];
			globalState.f1 := if (f7 !in v6) then fm2(fm1(globalState), globalState) else if (f7) then |(if (v4 in v7) then v7[v4] else v8)| else v3[safeIndex(365, v3.Length)];
			var v9: map<char, bool> := map[f6 := f7 || f7];
			v9 := v9[f6 := false];
		} else {
			var v10: map<int, bool> := map[fm2(f7, globalState) := f7];
			var v11: multiset<bool> := multiset{f7};
			var v12 := new C0(fm5(|v10|, f7, f7, globalState), multiset{f7, f7} == v11);
			var v13: array<bool> := new bool[28];
			var v14 := DC0(v13);
			v13 := v14.cf0;
			var v15: map<bool, int> := map[v12.f9 := p0];
			v15 := v15[v12.f9 := fm2(f7, globalState)];
			var v16 := 'q';
			v16 := if (!v12.f9) then f6 else f6;
			var v17: array<set<C0>> := new set<C0>[18];
			var v18: set<array<bool>> := {v13, v13};
			var v19: map<set<array<bool>>, bool> := map[v18 := false];
			v17 := if (if (v12.f9) then if (v18 in v19) then v19[v18] else v12.fm4(globalState) else f7) then v17 else v17;
		}
		
		var v20 := "swfjwfa";
		var v21 := DC4(false);
		var v22: set<int> := {p0, |v20|, p1};
		var v23: multiset<int> := multiset{|v22|};
		var v24: array<bool> := new bool[15] [f7, !f7, p1 > p0, !(v20 == v20), !v21.cf10, f7, !fm1(globalState), f7, !f7, v23 == v23, fm1(globalState), f7, f7, f7, f7];
		v24[safeIndex(448, v24.Length)] := p0 < fm2(f7, globalState);
		v24[safeIndex(448, v24.Length)] := f7;
		globalState.f5 := f7;
		var v25: map<int, int> := map[-0x2a1 := p1];
		var v26: set<string> := {(v20 + v20)[safeIndex(|v25|, |(v20 + v20)|) := 'h']};
		r1 := |v26|;
		var v27: set<bool> := {f7, f7, f7};
		var v28: array<seq<bool>> := new seq<bool>[14];
		var v29 := DC1(p1, v27, p0, v28, f6);
		var v30: array<D0> := new D0[19] [v29, v29, v29, v29, v29, v29, v29, v29, v29, v29, v29, v29, v29, v29, v29, v29, v29, v29, v29];
		var v31: map<array<D0>, bool> := map[v30 := f7];
		v31 := v31[v30 := v24[safeIndex(448, v24.Length)]];
		globalState.f5 := v24[safeIndex(448, v24.Length)];
		var v32: map<bool, int> := map[false := 192];
		r0 := safeModuloInt(if (v24[safeIndex(448, v24.Length)]) then |{true}| else p1, |v32|);
		var v33: C0 := new C0(v25, !f7);
		var v34: array<C0> := new C0[14] [v33, v33, v33, v33, v33, v33, v33, v33, v33, v33, v33, v33, v33, v33];
		var v35: map<bool, bool> := map[true := false];
		r1 := |map[if (f7) then v34 else v34 := if (if (f7 in v35) then v35[f7] else f7) then f7 else true]|;
		var v36: array<int> := new int[4];
		r2 := v36;
	}
}

function fm1(globalState: GlobalState): bool {
	(|multiset{!true, true}| + |"y"|) < ----0x101
}
function fm2(p0: bool, globalState: GlobalState): int {
	(|(map v0 : int | (0x62 <= v0) && (v0 < 163) :: (safeDivisionInt(v0, 0x83)) := (false))| - |[map[-0x2d2 := true], map[0x189 := false], map[|(seq(abs(-0x187), i0  => ('h')))| := false], map v1 : int | v1 in map[0x7a := 0x109] :: (v1 - 798) := (true), map[|(seq(abs(-223), i1  => ('y')))| := true]]|) - 0x75
}
function fm5(p0: int, p1: bool, p2: bool, globalState: GlobalState): map<int, int> {
	(map v0 : int | v0 in [|multiset{|map[|[false]| := true]|}|] :: (v0 * 0x322) := (-596)) + map[567 := |[[true, true], [false, false], [true]]|]
}
function fm6(globalState: GlobalState): char {
	if (multiset{|(seq(abs(-0x301), i0  => (985)))|, -713} >= multiset{|(seq(abs(837), i1  => ('s')))|, |multiset([|map[!false := 47]|, -|"okfva"|])|, -0x237, 0x318, -0x39}) then 'l' else 'k'
}
function fm7(p0: int, p1: int, p2: bool, globalState: GlobalState): D1 {
	DC4(!false)
}
function fm8(p0: int, p1: bool, p2: bool, p3: map<seq<bool>, int>, globalState: GlobalState): map<D2, int> {
	map v0 : D2 | v0 in map[DC9('k') := false] :: (v0) := (|("ev" + "sbujxi")|)
}
function fm9(p0: bool, p1: bool, globalState: GlobalState): multiset<map<bool, bool>> {
	if (true) then multiset{map[true := true]} + multiset{map[true := true], map[false := DC10(false, !false).cf20]} else multiset{map[true := false]} + multiset{map[true := true], map[false := false]}
}
function fm10(p0: int, p1: set<bool>, globalState: GlobalState): multiset<bool> {
	multiset{!(-0x2a5 > 0x304), DC21({multiset{false, !true}}).cf42 > {multiset{false, !false}}}
}
method Main() {
	var globalState := new GlobalState(470, -0x2ec, true, 0xdc, 0x227, false);
	var v0 := 'e';
	var v1 := new C1(v0, false);
	globalState.f1 := |multiset(seq(abs(-0x113), i0  => ([v1.f7])))|;
	var v2 := DC6('o', fm6(globalState));
	var v3: seq<D1> := [DC6(v0, v1.f6), if (false) then v2 else v2];
	v3 := [v2, v2, v2, v2, v2];
	var v4 := 0x16;
	var v5: map<int, int> := map[v4 := v4];
	var v6: map<bool, bool> := map[!v1.f7 := true];
	var v7: array<bool> := new bool[16](i1 => v1.f7);
	var v8 := DC2(v5, v6, v7);
	match (v8.(cf8 := v7)) {
		case DC1(cf1, cf2, cf3, cf4, cf5) =>
			v7 := new bool[2] [v1.f7, (fm7(0x2f6, v4, false, globalState)).cf10];
			var v9: map<bool, C1> := map[false := if (v1.f7) then v1 else v1];
			v9 := v9[v1.f7 := v1];
			v5 := v5[v4 := v4];
			var v10: multiset<int> := multiset{cf1};
			globalState.f5 := (v10 >= v10) <==> fm1(globalState);
		case DC2(cf6, cf7, cf8) =>
			var v11: seq<int> := [v4, v4, v4];
			var v12: map<seq<int>, int> := map[v11 := v4];
			var v13, v14, v15 := v1.m0(if (v11 in v12) then v12[v11] else 105, v4, globalState);
			var v16 := "olipx";
			v16 := (seq(abs(0x12d), i2  => (v1.f6))) + (seq(abs(434), i3  => (v1.f6)));
			globalState.f5 := v1.f7;
			var v17 := DC3(v14);
			var v18 := DC5(v17.cf9, cf6, |v16|);
			v15[safeIndex(445, v15.Length)] := v18.cf11;
			var v19: set<bool> := {v1.f7};
			var v20: array<seq<bool>> := new seq<bool>[23](i4 => [v1.f7]);
			var v21 := DC1(v4, v19, v14, v20, v0);
			var v22: map<char, bool> := map[v21.cf5 := !v1.f7];
			var v23: seq<map<char, bool>> := [v22[v1.f6 := v1.f7]];
			v15[safeIndex(445, v15.Length)] := |v23|;
		case DC3(cf9) =>
			var v24: array<int> := new int[14];
			var v25: seq<array<int>> := [v24];
			v25 := v25 + v25[safeIndex(v4, |v25|) := v24];
			globalState.f5 := !v1.f7;
			var v26: set<C1> := {DC8(v1).cf17};
			if (v26 !! (v26 + v26)) {
				var v27: array<map<D2, int>> := new map<D2, int>[16](i5 => map[DC9(v1.f6) := v4]);
				var v28: map<D2, int> := map[DC9(v1.f6) := v4];
				v27[safeIndex(150, v27.Length)] := if (!v1.f7) then v28 else v28;
				v27[safeIndex(150, v27.Length)] := v28 + fm8(v4, v1.f7, v1.f7, map[[v1.f7] := cf9], globalState);
				v1, globalState.f5 := v1, v1.f7;
				v7[safeIndex(698, v7.Length)] := v1.f7;
				var v29: set<bool> := {!v1.f7, v1.f7 || v1.f7, v1.f7};
				var v30: map<map<int, int>, bool> := map[v5 := !v1.f7];
				v7[safeIndex(698, v7.Length)], v29 := |v30| >= v4, v29;
				var v31: C0 := new C0(v5, v7[safeIndex(698, v7.Length)]);
				v31 := v31;
				var v32: array<string> := new string[18];
				var v33 := "bhphawf";
				v32[safeIndex(990, v32.Length)] := v33;
				v32[safeIndex(990, v32.Length)] := v33 + v33;
			} else {
				v7[safeIndex(981, v7.Length)] := v1.f7;
				v7[safeIndex(981, v7.Length)] := if (v1.f7) then cf9 == -0xbe else fm1(globalState);
				v1 := v1;
				var v34 := new C1(v1.f6, v1.f7);
				var v35: seq<int> := [v4, cf9];
				globalState.f5 := !((v35 <= v35) <== v1.f7);
				var v36 := DC8(v1);
				var v37: array<C1> := new C1[14] [v36.cf17, v1, v1, v34, v34, v34, v1, v34, v34, v1, if (v1.f7) then v1 else v1, v1, v34, v34];
				v37[safeIndex(895, v37.Length)] := v1;
				v37[safeIndex(895, v37.Length)] := v1;
			}
			
			var v38 := "qsq";
			v4 := v4 - |v38|;
		case DC0(cf0) =>
			var v39: array<int> := new int[24](i6 => i6 + v4);
			v39[safeIndex(624, v39.Length)] := v4;
			v39[safeIndex(624, v39.Length)], globalState.f1 := v4, 0x118;
			var v40 := DC0(v7);
			var v41: map<D0, int> := map[v40 := v39[safeIndex(624, v39.Length)]];
			var v42 := "htil";
			var v43 := DC5(|v42|, v5, v39[safeIndex(624, v39.Length)]);
			v39[safeIndex(624, v39.Length)] := if (v1.f7) then safeModuloInt(if (v40 in v41) then v41[v40] else v39[safeIndex(624, v39.Length)], 0x26f) else -v43.cf13;
			v7 := v7;
			v6 := v6[v1.f7 := v1.f7];
	}
	
	for i7 := v4 to v4 {
		globalState.f1 := i7 - fm2(false, globalState);
		var v44: array<int> := new int[22];
		var v45 := "a";
		v44[safeIndex(413, v44.Length)] := v4 + -|v45|;
		v44[safeIndex(413, v44.Length)] := v4;
		var v46: map<string, int> := map["ekadg" := v4];
		v46 := map["tpsq" := v44[safeIndex(413, v44.Length)]] + v46;
		v7[safeIndex(132, v7.Length)] := !false;
		v7[safeIndex(132, v7.Length)], globalState.f5, v0 := v1.f7, v1.f7, v1.f6;
	}
	forall i8 | 0 <= i8 < v7.Length {
		v7[i8] := v1.f7;
	}
	var v47 := "fyyyndfmr";
	v47 := if (v1.f7) then v47 else "vintro";
	for i9 := v4 to 320 {
		var v48, v49, v50 := v1.m0(v4, v4, globalState);
		var v51 := DC8(v1);
		var v52 := DC12(v51);
		match (v52) {
			case DC9(cf18) =>
				v50[safeIndex(687, v50.Length)] := -|map[v1.f7 := i9]|;
				v50[safeIndex(687, v50.Length)] := i9;
				var v53: array<array<bool>> := new array<bool>[4];
				v53[safeIndex(613, v53.Length)] := v7;
				v53[safeIndex(613, v53.Length)] := v7;
				var v54 := DC7(v1.f6);
				var v55: map<D1, int> := map[v54 := -safeDivisionInt(-|"pb"|, -v49)];
				v55 := v55;
				v5 := v5[v4 + i9 := 0x141];
			case DC10(cf19, cf20) =>
				var v56: map<bool, int> := map[true := v48];
				v50[safeIndex(536, v50.Length)] := |v56[v1.f7 := v48]|;
				v50[safeIndex(536, v50.Length)] := safeDivisionInt(|v47|, v49);
				v7[safeIndex(6, v7.Length)] := cf19;
				v7[safeIndex(6, v7.Length)] := true ==> (v0 !in v47);
				var v57: map<bool, set<bool>> := map[fm2(v1.f7, globalState) != v50[safeIndex(536, v50.Length)] := {v7[safeIndex(6, v7.Length)]}];
				var v58 := DC4(cf19);
				var v59: map<D1, string> := map[v58 := "g"];
				var v60: seq<int> := [v49];
				var v61: set<bool> := {cf20, true, fm1(globalState), cf19, cf19};
				cf19, globalState.f1, v57, v1 := (if (DC4(cf19) in v59) then v59[DC4(cf19)] else v47) != v47, -v49, v57[v50[safeIndex(536, v50.Length)] in v60 := {cf19, cf20, cf19} * v61], v1;
				var v62: map<int, multiset<C1>> := map[v49 := multiset{v1, v1}];
				var v63 := DC8(v1);
				var v64: multiset<C1> := multiset{v1, v1, v1, v63.cf17, v1};
				var v65, v66, v67 := v1.m0(|v60|, |(if (v49 in v62) then v62[v49] else v64)|, globalState);
			case DC11(cf21, cf22) =>
				cf21, v50 := v4, v50;
				globalState.f5 := true;
				globalState.f5 := v1.f7;
				var v68: map<int, bool> := map[v49 := true];
				v68 := v68[-12 := false];
			case DC8(cf17) =>
				v7 := v7;
				globalState.f1 := v48;
				var v69: map<int, bool> := map[|v47| := !v1.f7];
				var v70, v71, v72 := v1.m0(v49, safeModuloInt(|v69|, 440), globalState);
				globalState.f5 := v1.f7;
			case DC12(cf23) =>
				globalState.f5 := !fm1(globalState);
				var v73: C0 := new C0(v5, v1.f7);
				var v74: multiset<C0> := multiset{v73, v73, v73, v73};
				var v75: seq<int> := [|v74|];
				var v76: seq<seq<int>> := [v75];
				var v77: set<int> := {v48, |v76|};
				globalState.f5, v77 := v1.f7, v77;
				v0 := v0;
				var v78: multiset<array<bool>> := multiset{v7, v7};
				var v79: map<int, multiset<array<bool>>> := map[v49 := (multiset{v7})[v7 := abs(v49)]];
				var v80: array<multiset<array<bool>>> := new multiset<array<bool>>[14] [v78, v78, v78, v78, v78[v7 := abs(v48)], v78, v78 - v78, v78 + v78, v78, v78, v78, if (v48 in v79) then v79[v48] else v78, v78 - v78, multiset{v7}];
				v80[safeIndex(52, v80.Length)] := ((multiset{v7})[v7 := abs(i9)])[v7 := abs(|v47|)];
				var v81: seq<array<bool>> := [v7];
				var v82: seq<multiset<array<bool>>> := [v78, v78, v78, v78, multiset{v7, v7, v81[safeIndex(v4, |v81|)], v7, v7}];
				var v83: map<bool, int> := map[v1.f7 := v4];
				var v84: map<set<int>, map<bool, int>> := map[v77 := v83];
				v80[safeIndex(52, v80.Length)] := v82[safeIndex(|(if (v77 in v84) then v84[v77] else v83)|, |v82|)] - v78;
		}
		
		v7[safeIndex(859, v7.Length)] := v1.f7;
		v7[safeIndex(859, v7.Length)] := true;
		var v85 := new C1('b', false);
	}
	var i10 := 0;
	while (v1.f7)
		decreases 100 - i10
	{
		if (i10 >= 100) {
			break;
		}
		
		i10 := i10 + 1;
		if (v1.f7) {
			globalState.f1 := -safeDivisionInt(v4, v4);
			var v86 := DC9(v1.f6);
			v86 := v86;
			var v87, v88, v89 := v1.m0(v4, v4, globalState);
			var v90: map<array<bool>, int> := map[v7 := v4];
			v90 := v90[v7 := v4];
			v47 := v1.fm0("arwyvlusb", v1.f6, globalState);
		} else {
			v0 := v47[safeIndex(v4, |v47|)];
			v7[safeIndex(41, v7.Length)] := v1.f7;
			v47, globalState.f5, v7[safeIndex(41, v7.Length)] := (v47 + v47) + "hgpygk", v1.f7, v1.f7;
			globalState.f5 := !(v4 < v4);
			var v91: array<char> := new char[2] [v0, v0];
			var v92: seq<array<char>> := [v91];
			var v93 := DC13(v91);
			var v94: array<array<char>> := new array<char>[26] [v91, v91, v91, v91, v91, v91, v91, v91, v91, if (v1.f7) then v91 else v91, v91, v92[safeIndex(v4, |v92|)], v91, v91, v91, v91, v91, v91, v91, v93.cf24, v91, v91, v91, v91, v91, v91];
			v94 := v94;
			var v95: multiset<int> := multiset{v4};
			var v96: array<seq<array<D2>>> := new seq<array<D2>>[22];
			var v97: array<D2> := new D2[3](i11 => DC9(v1.f6));
			var v98: seq<array<D2>> := [v97, v97];
			v96[safeIndex(652, v96.Length)] := v98;
			var v99: set<bool> := {v7[safeIndex(41, v7.Length)]};
			v7[safeIndex(41, v7.Length)], v95, globalState.f1, v96[safeIndex(652, v96.Length)], globalState.f5 := v7[safeIndex(41, v7.Length)], v95, safeModuloInt(v4, v4), [v97], (if (v1.f7) then v99 else v99) > {v1.f7};
		}
		
		if (v47 == v47) {
			var v100: map<int, bool> := map[v4 := v1.f7];
			v100 := v100[safeDivisionInt(v4, |multiset{v47}|) := v1.f7];
			v7[safeIndex(823, v7.Length)] := v1.f7;
			var v101: array<D3> := new D3[7];
			var v102: map<array<D3>, bool> := map[v101 := v1.f7];
			var v103: multiset<int> := multiset{v4, v4, v4};
			var v104: set<bool> := {!v1.f7, v1.f7};
			v7[safeIndex(823, v7.Length)], globalState.f1, globalState.f0 := v1.f7, --0x2df, |v102[v101 := v103 > multiset{v4, v4, |v104|, v4, v4}]|;
			var v105: array<array<int>> := new array<int>[7];
			var v106: map<array<array<int>>, bool> := map[v105 := v1.f7];
			v106 := v106[v105 := v1.f7];
			var v107: map<C1, int> := map[v1 := 0x326];
			var v108: seq<bool> := [v1.f7, fm1(globalState), v1.f7, !v1.f7, v7[safeIndex(823, v7.Length)]];
			var v109 := new C0(v5[if (v1 in v107) then v107[v1] else |v108| := v4], v7[safeIndex(823, v7.Length)]);
			v4 := v4;
		} else {
			v7[safeIndex(315, v7.Length)] := fm1(globalState);
			v7[safeIndex(315, v7.Length)] := v1.f7;
			var v110: map<int, string> := map[v4 := v47];
			var v111: map<int, map<int, string>> := map[-0x2d := v110];
			v111 := v111[v4 := v110];
			var v112 := DC4(v1.f7);
			v112 := v112.(cf10 := v7[safeIndex(315, v7.Length)]);
			globalState.f0 := -v4;
			var v113 := new C1('j', v7[safeIndex(315, v7.Length)]);
		}
		
		var v114: seq<int> := [fm2(v1.f7, globalState), 555];
		var v116: set<int> := {v4};
		globalState.f5 := (set v115 : int | v115 in v114 :: (v115 - -0x29c)) > v116;
		var v117: array<string> := new string[2] [v47, if (!v1.f7) then "tdaboijva" else v47];
		v117[safeIndex(996, v117.Length)] := v47;
		v117[safeIndex(996, v117.Length)] := v1.fm0(v1.fm0("gqains", v0, globalState), v1.f6, globalState);
	}
	var v118: multiset<map<bool, bool>> := multiset{v6};
	var v119: array<int> := new int[11] [|(seq(abs(-948), i12  => (v1.f6)))|, v4, v4, safeModuloInt(0x26d, 0x299), |v6|, fm2(if (false in v6) then v6[false] else v1.f7, globalState), v4, v4, v4, v4, |(fm9(v1.f7, v1.f7, globalState) - v118)|];
	v119[safeIndex(229, v119.Length)] := v4;
	var v120: multiset<int> := multiset{v4, v4};
	var v121: seq<int> := [|v47|];
	v119[safeIndex(229, v119.Length)], v4, v4, globalState.f5 := v4, if (-0x2b6 in v120) then v120[-0x2b6] else v121[safeIndex(fm2(v1.f7, globalState), |v121|)], v4, !(!v1.f7 <== v1.f7);
	v119[safeIndex(229, v119.Length)] := |(seq(abs(0x368), i13  => (v0)))|;
	var v122: seq<bool> := [v1.f7 <== v1.f7];
	globalState.f5 := v122[safeIndex(v4, |v122|)];
	var v123: map<bool, int> := map[v1.f7 := v119[safeIndex(229, v119.Length)]];
	var v124 := DC3(v4);
	var v125: map<bool, D0> := map[v122[safeIndex(|v120[if (!!v1.f7 in v123) then v123[!!v1.f7] else v4 := abs(v119[safeIndex(229, v119.Length)])]|, |v122|)] := v124];
	v125 := v125[v1.f7 := v124];
	if (|v47| <= v4) {
		var v126: set<int> := {0x84, v4};
		v126 := v126 * ((set v127 : int | (0x3a0 <= v127) && (v127 < 795) :: (v127 * |(seq(abs(0x15b), i14  => (v119[safeIndex(229, v119.Length)])))|)) - v126);
		globalState.f5 := !v1.f7;
		var v128: seq<map<int, int>> := [v5];
		var v130: map<int, bool> := map[v119[safeIndex(229, v119.Length)] := v1.f7];
		var v131: seq<map<int, bool>> := [v130];
		var v132 := new C0((v128[safeIndex(|v121|, |v128|)])[v4 := |(map v129 : map<int, bool> | v129 in v131 :: (v129) := (if (true in v6) then v6[true] else v1.f7))|], v1.f7);
		var v133 := DC5(safeDivisionInt(v119[safeIndex(229, v119.Length)], v4), v132.f8 + v5, if (v4 in v132.f8) then v132.f8[v4] else v4);
		match (v133) {
			case DC5(cf11, cf12, cf13) =>
				var v134, v135, v136 := v1.m0(v119[safeIndex(229, v119.Length)], cf11, globalState);
				var v137 := new C0(map[|v121| := |v47|] + cf12, true);
				var v138: multiset<char> := multiset{v1.f6};
				v138 := v138;
				v0 := v47[safeIndex(cf11, |v47|)];
			case DC6(cf14, cf15) =>
				globalState.f5, globalState.f1 := v132.f9, safeDivisionInt(0x254, v4);
				var v139: array<C0> := new C0[12] [v132, v132, v132, v132, v132, v132, if (true) then v132 else v132, v132, v132, v132, v132, v132];
				v139[safeIndex(857, v139.Length)] := v132;
				var v140: array<char> := new char[5];
				var v141: array<array<char>> := new array<char>[17] [v140, v140, v140, v140, v140, v140, v140, v140, v140, if (v1.f7) then v140 else v140, v140, v140, v140, v140, v140, v140, v140];
				v139[safeIndex(857, v139.Length)], globalState.f1, v141, globalState.f1, globalState.f5 := v132, v119[safeIndex(229, v119.Length)], v141, v119[safeIndex(229, v119.Length)], !v1.f7;
				cf15 := v1.f6;
				globalState.f5 := v132.f9 || v1.f7;
			case DC7(cf16) =>
				var v142, v143, v144 := v1.m0(v119[safeIndex(229, v119.Length)], v4, globalState);
				var v145: map<multiset<bool>, int> := map[multiset{v1.f7, v1.f7, v1.f7, v1.f7} := v4];
				v145 := v145[fm10(105, {v1.f7, false, v1.f7, v132.f9, !!!v1.f7}, globalState) := if (!v1.f7) then v4 else v119[safeIndex(229, v119.Length)]];
				var v146 := DC10(v132.f9, v1.f7);
				v146 := DC10(v132.f9, v1.f7);
				var v147: multiset<bool> := multiset{v1.f7, v132.f9};
				v7[safeIndex(807, v7.Length)] := -|v147| >= v4;
				globalState.f5, v7[safeIndex(807, v7.Length)], globalState.f0 := v119[safeIndex(229, v119.Length)] != v119[safeIndex(229, v119.Length)], v132.f9, -safeModuloInt(v142, v142 - |v123|);
			case DC4(cf10) =>
				v7[safeIndex(137, v7.Length)] := cf10;
				v7[safeIndex(137, v7.Length)] := cf10;
				var v148 := DC10(true, v132.fm4(globalState));
				v7[safeIndex(137, v7.Length)] := v148.cf20;
				v0 := v1.f6;
				globalState.f5 := cf10 <==> v132.f9;
		}
		
		var v149: array<string> := new string[21];
		v149[safeIndex(774, v149.Length)] := v47 + v47;
		var v150 := DC10(v1.f7, v1.f7);
		var v151 := DC14(v1, v1.f6, v1, v132, v132.f9);
		var v152: set<D3> := {DC14(v1, v1.f6, v1, v132, v132.f9), v151};
		globalState.f5, v119[safeIndex(229, v119.Length)], v149[safeIndex(774, v149.Length)] := !(!!v150.cf19 && (v152 !! v152)), v119[safeIndex(229, v119.Length)], v47;
	} else {
		var v153: seq<string> := [v47, v47 + v47, v47];
		var v154 := DC16(v153);
		v153 := (v153 + v154.cf31[safeIndex(-v119[safeIndex(229, v119.Length)], |v154.cf31|) := v47]) + v153;
		var v155 := DC18(v119);
		v119 := v155.cf37;
		var v156, v157, v158 := v1.m0(v119[safeIndex(229, v119.Length)], if (v1.f7 in v123) then v123[v1.f7] else v4, globalState);
		globalState.f5 := v1.f7;
		v7[safeIndex(704, v7.Length)] := v1.f7;
		var v159: C0 := new C0(map[v119[safeIndex(229, v119.Length)] := v157], v1.f7);
		var v160: seq<C0> := [v159, v159, v159];
		globalState.f1, v0, v7[safeIndex(704, v7.Length)], v156, v119[safeIndex(229, v119.Length)] := v157, fm6(globalState), -|v160| < v4, v157 - v157, 0x3af;
	}
	
	var v161: set<int> := {110, v119[safeIndex(229, v119.Length)]};
	var v162: multiset<string> := multiset{v47};
	var v163, v164, v165 := v1.m0(|v161| - |v162|, if (v1.f7) then v119[safeIndex(229, v119.Length)] else v4, globalState);
	v164 := 0x2ed;
	print globalState.f0, "\n";
	print globalState.f1, "\n";
	print globalState.f2, "\n";
	print globalState.f3, "\n";
	print globalState.f4, "\n";
	print globalState.f5, "\n";
	print v0, "\n";
	print v1.f6, "\n";
	print v1.f7, "\n";
	print v2.cf14, "\n";
	print v2.cf15, "\n";
	print v3 == [DC6('o', 'k'), DC6('o', 'k'), DC6('o', 'k'), DC6('o', 'k'), DC6('o', 'k')], "\n";
	print v4, "\n";
	print v5 == map[22 := 22], "\n";
	print v6 == map[true := true], "\n";
	print v7[0], "\n";
	print v7[1], "\n";
	print v7[2], "\n";
	print v7[3], "\n";
	print v7[4], "\n";
	print v7[5], "\n";
	print v7[6], "\n";
	print v7[7], "\n";
	print v7[8], "\n";
	print v7[9], "\n";
	print v7[10], "\n";
	print v7[11], "\n";
	print v7[12], "\n";
	print v7[13], "\n";
	print v7[14], "\n";
	print v7[15], "\n";
	print v8.cf6 == map[22 := 22], "\n";
	print v8.cf7 == map[true := true], "\n";
	print v8.cf8[0], "\n";
	print v8.cf8[1], "\n";
	print v8.cf8[2], "\n";
	print v8.cf8[3], "\n";
	print v8.cf8[4], "\n";
	print v8.cf8[5], "\n";
	print v8.cf8[6], "\n";
	print v8.cf8[7], "\n";
	print v8.cf8[8], "\n";
	print v8.cf8[9], "\n";
	print v8.cf8[10], "\n";
	print v8.cf8[11], "\n";
	print v8.cf8[12], "\n";
	print v8.cf8[13], "\n";
	print v8.cf8[14], "\n";
	print v8.cf8[15], "\n";
	print v47, "\n";
	print i10, "\n";
	print v118 == multiset{map[true := true]}, "\n";
	print v119[0], "\n";
	print v119[1], "\n";
	print v119[2], "\n";
	print v119[3], "\n";
	print v119[4], "\n";
	print v119[5], "\n";
	print v119[6], "\n";
	print v119[7], "\n";
	print v119[8], "\n";
	print v119[9], "\n";
	print v119[10], "\n";
	print v120 == multiset{22, 22}, "\n";
	print v121 == [6], "\n";
	print v122 == [true], "\n";
	print v123 == map[false := 872], "\n";
	print v124.cf9, "\n";
	print v125 == map[true := DC3(22), false := DC3(22)], "\n";
	print v161 == {110, 872}, "\n";
	print v162 == multiset{"vintro"}, "\n";
	print v163, "\n";
	print v164, "\n";
}