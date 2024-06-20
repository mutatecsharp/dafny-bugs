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
datatype D0 = DC0(cf0: char)
datatype D1 = DC2(cf2: set<seq<int>>) | DC3(cf3: multiset<bool>, cf4: bool, cf5: bool) | DC1(cf1: array<int>)
datatype D2 = DC4(cf6: int)
datatype D3 = DC6(cf8: int, cf9: int, cf10: bool) | DC5(cf7: map<array<bool>, int>) | DC7(cf11: D3)
datatype D4 = DC8(cf12: seq<C0>)
datatype D5 = DC9(cf13: string)
datatype D6 = DC11(cf15: int) | DC10(cf14: C0)
datatype D7 = DC13(cf17: bool) | DC12(cf16: multiset<seq<int>>) | DC14(cf18: D7)
datatype D8 = DC16(cf20: bool, cf21: multiset<bool>, cf22: seq<D7>, cf23: map<int, int>) | DC17(cf24: seq<char>, cf25: int, cf26: int, cf27: int) | DC15(cf19: seq<int>) | DC18(cf28: D8)
datatype D9 = DC20(cf30: array<C0>, cf31: bool, cf32: int) | DC19(cf29: map<int, seq<bool>>)
datatype D10 = DC22(cf34: bool, cf35: int, cf36: bool, cf37: int) | DC21(cf33: seq<bool>) | DC23(cf38: D10)
trait T0 {
	var f0 : bool
	function fm4(globalState: GlobalState): D2 
	function fm5(p0: int, p1: bool, globalState: GlobalState): int 
}

class GlobalState {
	constructor () {
	}
	
}

class C0 extends T0 {
	constructor (f0 : bool) {
		this.f0 := f0;
	}
	
	function fm4(globalState: GlobalState): D2 {
		match DC2({[|map[840 := f0]|]}) {
			case DC2(cf2) => DC4(-0x12b)
			case DC3(cf3, cf4, cf5) => DC4(|{cf4}|)
			case DC1(cf1) => DC4(-|(map v0 : int | (617 <= v0) && (v0 < 0x24a) :: (v0 * |"vo"|) := (129))|)
		}
	}
	function fm5(p0: int, p1: bool, globalState: GlobalState): int {
		-(39 * DC4(0x1c4).cf6)
	}
	function fm6(p0: int, p1: multiset<map<bool, bool>>, p2: int, globalState: GlobalState): bool {
		{seq(abs(0xde), i0  => (-0x18a)), seq(abs(-0x310), i1  => (|"hpmca"|)), [|"hs"|, |"gjhdka"|, -0x2fd], seq(abs(0x1ff), i2  => (240)), [0x32f, 633, DC4(|{-0x15b}|).cf6, |map[false := true]|, 390]} <= {[0x247]}
	}
}

function fm0(p0: string, p1: int, p2: int, globalState: GlobalState): char {
	if (|[!!false]| < |(seq(abs(-917), i0  => ('m')))|) then DC0('h').cf0 else 'c'
}
function fm1(p0: string, p1: int, p2: set<bool>, globalState: GlobalState): bool {
	true
}
function fm2(p0: bool, globalState: GlobalState): int {
	(|map['x' := true]| - |[|[|multiset([490])|]|]|) * |(map[true := "wtycevgkj"] + map[!!true := DC9("lqmpa").cf13])|
}
function fm3(globalState: GlobalState): seq<int> {
	([|map[229 := |(seq(abs(0x346), i0  => ('d')))|]|, -|(seq(abs(0x126), i1  => ('b')))|] + (seq(abs(720), i2  => (|"auo"|)))) + ([|(seq(abs(-51), i3  => (0x13e)))|, |[true, !!!true, true]|] + (seq(abs(0x1f1), i4  => (0xf5))))
}
function fm7(p0: int, p1: int, globalState: GlobalState): string {
	"p" + "rjwgaoud"
}
function fm8(p0: bool, p1: bool, p2: int, globalState: GlobalState): D5 {
	match DC14(DC14(DC14(DC14(DC12(multiset{[0x3dd, |(set v2 : int | v2 in (map v0 : int | (0x2e <= v0) && (v0 < 0x1e3) :: (safeModuloInt(v0, |(map v1 : int | v1 in multiset{-0x138, |multiset{0xdc}|} :: (safeDivisionInt(v1, -0x28b)) := (0x22b))|)) := (0x383)) :: (v2 * |map[|map[!false := |{899, 0x2bf, |"llm"|}|]| := -0x2c2]|))|]}))))) {
		case DC13(cf17) => DC9(seq(abs(37), i0  => ('l')))
		case DC12(cf16) => DC9("y")
		case DC14(cf18) => DC9("hvccdvgrp")
	}
}
function fm9(p0: int, p1: int, p2: int, p3: char, globalState: GlobalState): set<map<int, bool>> {
	{map[916 := true] + map[|map[false := -0x263]| := false], map[0x3df := !true]}
}
function fm10(p0: bool, p1: int, p2: string, p3: seq<seq<int>>, globalState: GlobalState): D0 {
	match if (!true) then DC11(-0x2a1) else DC11(|{|"gqh"|}|) {
		case DC11(cf15) => DC0('g')
		case DC10(cf14) => DC0('r')
	}
}
function fm11(globalState: GlobalState): map<bool, set<char>> {
	(map[false := {'i'}] + map[false := {'g'}]) + map[false := set v0 : char | v0 in map['b' := -0x202] :: (v0)]
}
function fm12(p0: bool, p1: char, p2: bool, p3: int, globalState: GlobalState): multiset<seq<int>> {
	(if (false) then multiset{seq(abs(0x86), i0  => (0x3af))} else multiset{[0x7, -|map[0x208 := !true]|, 0x1f0, ---0x2f]}) + multiset(if (!true) then [[975]] else [seq(abs(907), i1  => (---0x312)), [|multiset{0x32b, |(seq(abs(0x3cd), i2  => ('x')))|, 948}|, |(map v0 : int | v0 in (map v1 : int | (481 <= v1) && (v1 < 0x251) :: (safeModuloInt(v1, 0x34b)) := (|[false]|)) :: (safeDivisionInt(v0, -620)) := ("vyeoofmh"))|, |"lmjx"|]])
}
function fm13(p0: bool, p1: bool, p2: bool, globalState: GlobalState): set<bool> {
	{true} * {false, false}
}
function fm14(p0: int, globalState: GlobalState): set<seq<int>> {
	{(seq(abs(-0xad), i0  => (-|{!false, true}|))) + [|[|{|"vetmckar"|}|]|], [|(map v0 : multiset<bool> | v0 in [multiset{true, true}, multiset{true}] :: (v0) := (|[-|map[true := false]|]|))|, |{!!!true}|, 0x29e], [0x332, |DC21([!true]).cf33|, 699]}
}
function fm15(globalState: GlobalState): map<int, bool> {
	(map[-782 := false] + map[|multiset{|multiset{-698}|}| := !true]) + map[-0x17d := !false]
}
function fm16(globalState: GlobalState): set<int> {
	{|[[false, true]]|}
}
function fm17(p0: int, p1: int, p2: bool, globalState: GlobalState): D1 {
	if ({[false]} !! {[false, true]}) then if (true) then DC3(multiset{!false, false, false}, false, true) else DC3(multiset{true}, false, false) else DC3(multiset{!!true, true}, !DC13(true).cf17, true)
}
method m0(p0: map<array<bool>, int>, p1: bool, globalState: GlobalState) returns (r0: int, r1: int, r2: int, r3: set<set<bool>>) {
	var v0 := 0xe8;
	var v1: seq<int> := [v0];
	var v2 := DC15(v1);
	v1 := v2.cf19;
	var v3 := DC13(p1);
	var v4 := DC14(v3);
	var v5: map<D7, bool> := map[v4 := p1];
	v5 := v5[v4 := false];
	var v6: map<bool, int> := map[p1 := v0];
	var v7: set<map<bool, int>> := {v6, v6};
	var i0 := 0;
	while ((set v8 : map<bool, int> | v8 in v7 :: (v8)) > v7)
		decreases 100 - i0
	{
		if (i0 >= 100) {
			break;
		}
		
		i0 := i0 + 1;
		var v9: array<int> := new int[8](i1 => i1 * v0);
		v9 := DC1(v9).cf1;
		var v10: array<multiset<seq<int>>> := new multiset<seq<int>>[18](i2 => multiset{v1} - multiset{v1, ([v0, |v1|])[safeIndex(v0, |[v0, |v1|]|) := -|[p1]|]});
		var v11: seq<seq<int>> := [v1];
		v10[safeIndex(232, v10.Length)] := multiset([[v0]] + v11);
		var v12: C0 := new C0(p1);
		var v13: map<C0, seq<int>> := map[v12 := v1];
		var v14: multiset<seq<int>> := multiset{[v0, v0], if (v12 in v13) then v13[v12] else v1};
		v10[safeIndex(232, v10.Length)] := DC12(v14).cf16;
		var v15 := false;
		var v16 := 'r';
		var v17 := DC0(v16);
		var v18: seq<char> := [v17.cf0];
		v15, v12 := fm7(|v18|, v0, globalState) != (if (p1) then "wrtno" else v18), v12;
		v1 := [v0] + v1;
	}
	var v19 := true;
	v19 := p1;
	var i3 := 0;
	while (p1)
		decreases 100 - i3
	{
		if (i3 >= 100) {
			break;
		}
		
		i3 := i3 + 1;
		v19 := p1;
		r0 := v0;
		var v20: seq<bool> := [false];
		var v21: seq<D7> := [v4, DC14(v3)];
		var v22: map<int, int> := map[v0 := v0];
		var v23 := DC16(v19, multiset(v20), v21, v22);
		var v24 := DC18(v23);
		var v25 := DC18(v23);
		var v26 := DC18(v23);
		var v27 := DC18(v26);
		var v28 := DC18(v26);
		match (v28) {
			case DC16(cf20, cf21, cf22, cf23) =>
				var v29: array<map<int, bool>> := new map<int, bool>[15];
				v29[safeIndex(447, v29.Length)] := fm15(globalState);
				var v30: map<int, bool> := map[v0 := cf20];
				var v31: seq<map<int, bool>> := [v30, v30];
				var v32 := "nixsur";
				var v33: map<string, int> := map[v32 := v0];
				var v34: seq<map<string, int>> := [v33, v33];
				var v35: set<bool> := {v19, v19, p1};
				v29[safeIndex(447, v29.Length)] := v31[safeIndex(|v34[safeIndex(|{v0, v0}|, |v34|)]|, |v31|)] + (v30[-v0 := p1])[v0 := fm1("hmgxqq", v0, v35, globalState)];
				var v36: T0 := new C0(p1);
				v36 := v36;
				var v37: set<int> := {-|map[v30[v0 := cf20] := v19]|};
				var v38: array<set<int>> := new set<int>[4] [v37, v37, {v0, v0}, v37];
				v38[safeIndex(317, v38.Length)] := v37;
				var v39: array<int> := new int[22](i4 => safeDivisionInt(i4, v0));
				v39[safeIndex(833, v39.Length)] := safeModuloInt(v0, |v22|);
				v39[safeIndex(982, v39.Length)] := |v1|;
				var v40: C0 := new C0(p1);
				var v41: seq<C0> := [v40];
				var v42 := DC8(([v41[safeIndex(v0, |v41|)]])[safeIndex(|v41|, |[v41[safeIndex(v0, |v41|)]]|) := v40] + v41);
				v38[safeIndex(317, v38.Length)], v32, v39[safeIndex(833, v39.Length)], v39[safeIndex(982, v39.Length)], v42 := (v37 - fm16(globalState)) * v37, v32, v1[safeIndex(v0, |v1|)], v0 * |v32|, v42;
				var v43: multiset<seq<int>> := multiset{[v0, v39[safeIndex(833, v39.Length)]], v1, v1, v1};
				var v44 := DC12(v43);
				cf20, r2, v44 := if (v36.f0) then cf20 else v36.f0, v39[safeIndex(833, v39.Length)], v44.(cf16 := v43);
			case DC17(cf24, cf25, cf26, cf27) =>
				var v45 := 'f';
				var v46 := DC6(v0, cf26, p1);
				var v47: set<bool> := {v46.cf10, p1, p1};
				var v48: map<int, bool> := map[v1[safeIndex(cf25, |v1|)] := fm1(("efgeywy")[safeIndex(cf25, |"efgeywy"|) := v45], v0, v47, globalState)];
				v48 := v48[cf27 := v19];
				v19 := cf25 <= v1[safeIndex(-DC11(cf27).cf15, |v1|)];
				var v49: C0 := new C0(v19);
				v49 := v49;
				var v50: array<int> := new int[4];
				v50 := v50;
			case DC15(cf19) =>
				var v51 := DC4(v0);
				var v52: map<bool, D2> := map[v19 := v51];
				var v53: map<bool, bool> := map[p1 := v19];
				var v54: map<map<bool, D2>, bool> := map[v52 := if (v19 in v53) then v53[v19] else p1];
				v54 := v54[v52 := p1];
				v19 := if (v19 in v53) then v53[v19] else !p1 ==> p1;
				v19 := p1;
				var v55: set<bool> := {v19, v19};
				v19 := (v55 * v55) > v55;
			case DC18(cf28) =>
				var v56: multiset<bool> := multiset{v20[safeIndex(v0, |v20|)], p1};
				var v57: set<bool> := {p1, p1};
				v19 := !fm1(fm7(if (v19 in v56) then v56[v19] else v0, v0, globalState), v0, v57, globalState);
				v1 := (if (p1) then v1 else seq(abs(0x38b), i5  => (v0)))[safeIndex(v0, |(if (p1) then v1 else seq(abs(0x38b), i5  => (v0)))|) := v0];
				var v58: C0 := new C0(p1);
				var v59: map<C0, bool> := map[v58 := p1];
				var v60: map<map<C0, bool>, bool> := map[v59 := p1];
				var v61: C0 := new C0(if (v59 in v60) then v60[v59] else p1);
				var v62 := "owkgj";
				var v63: map<int, string> := map[v0 := v62];
				var v64 := 'd';
				var v65: map<bool, char> := map[v19 := v64];
				var v66: map<int, C0> := map[|(if (v19) then if (v0 in v63) then v63[v0] else v62 else v62[safeIndex(v0, |v62|) := if (p1 in v65) then v65[p1] else v64])| := v61];
				var v67 := DC10(v61);
				v61 := if (|(v20[safeIndex(302, |v20|) := v19] + [v19, p1, true, v19])| in v66) then v66[|(v20[safeIndex(302, |v20|) := v19] + [v19, p1, true, v19])|] else v67.cf14;
				v19 := p1;
		}
		
		var v68 := "emnf";
		if (fm1(v68, v0, fm13(v19, v19, p1, globalState), globalState)) {
			v19 := v19;
			var v69: set<bool> := {v19, v19};
			var v70 := new C0(p1 || fm1(v68, |"ertqblss"|, v69, globalState));
			var v71: array<int> := new int[16];
			v71[safeIndex(624, v71.Length)] := v0;
			v19, v68, v71[safeIndex(624, v71.Length)] := v1 <= v1, v68, v0;
			var v72: array<bool> := new bool[22](i6 => false);
			v72[safeIndex(852, v72.Length)] := p1;
			v72[safeIndex(852, v72.Length)] := p1;
			var v73 := new C0((fm17(v0, 0x2f3, p1, globalState)).cf5);
		} else {
			var v74: array<int> := new int[23];
			v74[safeIndex(860, v74.Length)] := v0;
			var v75 := 'd';
			var v76: seq<string> := [v68];
			var v77 := DC4(v0);
			var v78: map<int, seq<string>> := map[v0 := v76];
			var v79: seq<seq<string>> := [if (-272 in v78) then v78[-272] else v76[safeIndex(v0, |v76|) := v68]];
			v74[safeIndex(860, v74.Length)], v75, v75, v76, v0 := v77.cf6, v75, v68[safeIndex(v0, |v68|)], v79[safeIndex(fm2(p1, globalState), |v79|)] + v76, safeDivisionInt(v0, v1[safeIndex(v0, |v1|)]);
			var v80: array<T0> := new T0[5];
			v80 := new T0[7];
			var v82 := DC19(map v81 : int | (-0x167 <= v81) && (v81 < 0xcb) :: (safeModuloInt(v81, v0)) := (v20));
			v74[safeIndex(860, v74.Length)] := |v82.cf29|;
			v19 := v19;
			var v83 := new C0(v19);
		}
		
	}
	var i7 := 0;
	while (!!v19)
		decreases 100 - i7
	{
		if (i7 >= 100) {
			break;
		}
		
		i7 := i7 + 1;
		var v84: seq<bool> := [!p1, v19];
		var v85: multiset<bool> := multiset{v84[safeIndex(v0, |v84|)]};
		var v86 := DC3(v85, v19, v19);
		var v87: set<bool> := {p1, v19, p1, v86.cf4};
		v19 := fm1(seq(abs(0xfd), i8  => ('l')), v0, v87, globalState);
		if (0x1c0 == v0) {
			var v88 := "nlwntcgy";
			var v89: C0 := new C0(p1);
			var v90 := 'm';
			v88 := v88[safeIndex(v0 * |{v89}|, |v88|) := if (p1) then v90 else 'j'];
			var v91: array<D8> := new D8[19](i9 => v2);
			var v92: set<array<D8>> := {v91};
			v19 := (if (v19) then 0x209 else v0) != |(v92 * v92)|;
			v1 := v1;
			var v93: array<bool> := new bool[26];
			v93[safeIndex(75, v93.Length)] := v19;
			v93[safeIndex(75, v93.Length)] := false;
			v88 := fm7(v0, v0, globalState);
		} else {
			var v94: map<bool, bool> := map[false := v19];
			v94 := v94[v19 := p1 ==> p1];
			r1 := v0;
			var v95: array<C0> := new C0[20];
			var v96: C0 := new C0(!p1);
			v95[safeIndex(849, v95.Length)] := v96;
			v95[safeIndex(849, v95.Length)] := v96;
			var v97 := new C0(p1);
			var v98: map<int, bool> := map[v0 := !p1];
			var v99: set<seq<int>> := {v1, v1};
			var v100 := DC2(v99);
			var v101 := 'w';
			var v102: seq<map<bool, bool>> := [v94];
			var v103: multiset<char> := multiset{v101, if (v97.fm6(v0, multiset(v102), v0, globalState)) then v101 else v101, v101, v101, v101};
			v19, v98, v100, v95, v103 := v0 in v98, (map v104 : int | v104 in fm3(globalState) :: (v104 - -v0) := (false)) + v98[0x22c := p1], v100, v95, v103 * v103;
		}
		
		var v105 := "hr";
		v0 := -safeModuloInt(0x14 + 0x28f, safeModuloInt(v0, |v105|));
		v19 := p1 && (if (v19) then v19 else false);
	}
	r0 := v0;
	r1 := v0;
	r2 := v0;
	var v106: set<bool> := {true};
	var v107: set<set<bool>> := {v106, v106, v106};
	r3 := v107;
}
method Main() {
	var globalState := new GlobalState();
	var v0: array<bool> := new bool[24];
	var v1 := 365;
	var v2 := true;
	var v3, v4, v5, v6 := m0(map[v0 := v1], v2, globalState);
	var v7 := "gbehkj";
	v4 := safeDivisionInt(safeModuloInt(-v4, |v7|), 0x32d);
	var v8 := 'k';
	var v9: map<bool, int> := map[v2 := v1];
	var v10 := DC0('u');
	var v11: array<char> := new char[20] [v8, v8, v8, v8, fm0(v7, -v5, v5, globalState), fm0(v7, |[|v9|, v5, v1]|, v1, globalState), v8, v8, v8, v8, 'c', v8, v10.cf0, fm0(seq(abs(0x1b9), i0  => (v8)), v5, v3, globalState), v8, v8, v7[safeIndex(v3, |v7|)], v8, fm0(seq(abs(916), i1  => ('j')), v3, v3, globalState), 'x'];
	v11 := v11;
	var v12: map<int, int> := map[v4 := v4];
	if ((if (v4 in v12) then v12[v4] else v4) != v5) {
		if (true) {
			var v13: array<int> := new int[28];
			v13 := v13;
			var v14: seq<int> := [v4, v4];
			v9 := v9[v2 := |v14|];
			var v15 := DC1(v13);
			v13 := if (!v2) then v13 else v15.cf1;
			v0[safeIndex(925, v0.Length)] := v2;
			v0[safeIndex(925, v0.Length)] := v2;
			v0[safeIndex(925, v0.Length)] := v0[safeIndex(925, v0.Length)];
		} else {
			var v16: map<array<bool>, int> := map[v0 := v4];
			var v17, v18, v19, v20 := m0(v16 + v16, true, globalState);
			v2 := v2;
			var v21: seq<bool> := [v2, v2];
			var v22: set<int> := {|v21|};
			var v23, v24, v25, v26 := m0(v16, v22 == v22, globalState);
			v23 := safeDivisionInt(v25, v25);
			var v27 := DC3(multiset([v2, v2]), v2, v2);
			v2 := (v27.(cf5 := v2)).cf4 && v2;
		}
		
		var v28: multiset<int> := multiset{v5, |(seq(abs(448), i2  => (multiset{0x1f})))|, v4, -0xf8};
		v2 := v28 !! (v28 + v28);
		var v29: multiset<string> := multiset{seq(abs(259), i3  => (v8))};
		var v30: seq<bool> := [v2];
		var v31 := DC4(|v30|);
		var v32: set<bool> := {v2, v2};
		var v33: set<bool> := {!v2, v2, fm1(v7, v31.cf6, v32, globalState), !!v2};
		v2, v2, v7, v8, v3 := v2, !(multiset{v7} <= (v29 * v29)), seq(abs(232), i4  => (v8)), v8, |v32|;
		var v34: map<array<bool>, int> := map[v0 := v4];
		var v35 := DC5(v34);
		var v36, v37, v38, v39 := m0(v35.cf7, !true, globalState);
		var v40: array<seq<string>> := new seq<string>[28];
		var v41: seq<string> := [v7, v7, v7];
		v40[safeIndex(468, v40.Length)] := v41 + v41;
		v40[safeIndex(468, v40.Length)] := [v7, v7] + [v7];
	} else {
		v1 := v1;
		v2 := false;
		v12 := v12[v3 := 0x109 + fm2(v2, globalState)];
		v2, v8, v4 := v2, v8, v3;
		v0[safeIndex(603, v0.Length)] := v2;
		v0[safeIndex(603, v0.Length)] := (v4 < v5) <==> v2;
	}
	
	var v42: set<bool> := {v2};
	if (if (true) then fm1(("pkkborlrx")[safeIndex(v5, |"pkkborlrx"|) := v8], fm2(v2, globalState), v42, globalState) else v42 == v42) {
		var v43: multiset<bool> := multiset{v2};
		var v44: map<array<bool>, int> := map[v0 := |v43|];
		var v45 := DC5(v44);
		var v46, v47, v48, v49 := m0(if (v2) then v44 else v45.cf7, v2, globalState);
		var v50, v51, v52, v53 := m0(v44, v2, globalState);
		var v54, v55, v56, v57 := m0(v44 + v44, v2, globalState);
		var v58, v59, v60, v61 := m0(v44 + v44, v2, globalState);
		v2, v2, v2, v11, v2 := !(v2 || v2), v2, v2, if (v2) then v11 else v11, v2;
	} else {
		var v62: map<array<bool>, int> := map[v0 := v1];
		var v63, v64, v65, v66 := m0(v62, v2, globalState);
		v65 := v3 * |v9|;
		v2 := !((if (v2) then v5 else if (v2 in v9) then v9[v2] else |[v2, v2]|) !in fm3(globalState));
		var v67: seq<string> := [v7, v7, v7];
		var v68 := new C0(fm1(v67[safeIndex(|[v3]|, |v67|)], v1, v42, globalState));
		var v69: array<seq<int>> := new seq<int>[9];
		var v70: seq<int> := [fm2(v2, globalState)];
		v69[safeIndex(525, v69.Length)] := v70;
		v69[safeIndex(525, v69.Length)] := v70;
	}
	
	var v71: C0 := new C0(v2);
	var v72: seq<bool> := [true];
	for i5 := 0x31c to -(if (false) then |[v71]| else -|v72|) {
		var v73: map<array<bool>, int> := map[v0 := v4];
		var v74, v75, v76, v77 := m0(v73 + v73, fm1(seq(abs(-0x31b), i6  => (v8)), v3, v42, globalState), globalState);
		var v78, v79, v80, v81 := m0(if (v2) then v73 else v73, v42 !! {v2, v2, false, v2, v2}, globalState);
		v2 := !false;
		var v82: seq<C0> := [v71, v71];
		var v83: array<seq<C0>> := new seq<C0>[3] [v82[safeIndex(v79, |v82|) := v71], v82, v82];
		v83[safeIndex(180, v83.Length)] := v82;
		var v84: map<bool, bool> := map[v2 := v2];
		var v85 := DC3(multiset{v2, if (v2 in v84) then v84[v2] else v2}, true, v2);
		var v86: set<int> := {38};
		var v87: array<int> := new int[7] [0x33a, v3, if (v85.cf4 in v9) then v9[v85.cf4] else v76, |v86|, v3, v76, v80];
		var v88: map<array<int>, seq<C0>> := map[v87 := v82];
		var v89: seq<array<int>> := [v87];
		v83[safeIndex(180, v83.Length)] := if (v89[safeIndex(|v7|, |v89|)] in v88) then v88[v89[safeIndex(|v7|, |v89|)]] else [v71] + v82;
	}
	var v90 := new C0(|v72| < |([v72[safeIndex(-989, |v72|)]])[safeIndex(|[v4, v1]|, |[v72[safeIndex(-989, |v72|)]]|) := v2]|);
	for i7 := fm2(v2, globalState) to v5 {
		var v91: seq<int> := [|v7|];
		var v92: map<array<bool>, int> := map[v0 := |v91|];
		var v93, v94, v95, v96 := m0(v92 + v92, v2, globalState);
		v2 := v2;
		var v97: array<seq<C0>> := new seq<C0>[15];
		var v98: seq<C0> := [v71, v90];
		v97[safeIndex(935, v97.Length)] := v98 + v98;
		var v99 := DC8([v71]);
		v97[safeIndex(935, v97.Length)] := v98 + v99.cf12;
		v71 := v71;
	}
	if (match if (v2) then DC0(v8) else DC0(v8) {
		case DC0(cf0) => v5 >= |map[0x2d := cf0]|
	}) {
		var v100: T0 := new C0(v2);
		var v101: map<bool, T0> := map[false := v100];
		var v102: array<T0> := new T0[18] [v100, v100, v100, v100, v100, v100, v100, v100, v100, v100, v100, v100, v100, if (true) then v100 else v100, v100, if (!v100.f0 in v101) then v101[!v100.f0] else v100, v100, v100];
		v102, v2 := v102, if (if (v100.f0) then !v100.f0 else v100.f0) then v100.f0 else v100.f0;
		v4 := v4;
		var v103: seq<string> := [v7, v7, v7, "hucmr"];
		var v104: seq<int> := [|v72|, 0x1df];
		var v105: map<seq<string>, char> := map[v103 + v103[safeIndex(v104[safeIndex(v4, |v104|)], |v103|) := seq(abs(0x5), i8  => (v8))] := v8];
		v105 := v105[[fm7(v3, v4, globalState), v7, v7] + [v7] := 'e'];
		var v106: array<int> := new int[7](i9 => i9 * 0x1b7);
		v106[safeIndex(858, v106.Length)] := v4;
		var v107 := DC4(v3);
		var v108: set<D2> := {if (v100.f0) then DC4(v1) else v107, v107};
		v106[safeIndex(858, v106.Length)] := |v108|;
		var v109 := DC9(v7);
		v7 := (if (false) then v109 else fm8(v2, !v2, v1, globalState)).cf13;
	} else {
		var v110 := new C0(v4 > v1);
		var v111 := DC3(multiset{v2, v2}, v2, v2);
		var v112 := DC10(v90);
		var v113: seq<C0> := [v71, v90, if (v2) then v71 else v71, v90, if (v111.cf4) then v110 else v112.cf14];
		var v114: array<set<map<int, bool>>> := new set<map<int, bool>>[17];
		var v115: map<int, bool> := map[v1 := v2];
		var v116: set<map<int, bool>> := {v115};
		v114[safeIndex(537, v114.Length)] := v116;
		var v117: T0 := new C0(v2);
		var v118: map<string, T0> := map[v7 := v117];
		v4, v113, v114[safeIndex(537, v114.Length)] := 0x2d1, (v113 + (v113 + v113))[safeIndex(|v118|, |(v113 + (v113 + v113))|) := v110], fm9(|(v7 + "icsjbibk")|, |v72[safeIndex(v4, |v72|) := v117.f0]|, |map[v117.f0 := -0x349]| + v117.fm5(v3, true, globalState), v8, globalState);
		var v119: multiset<bool> := multiset{v2};
		var v120: multiset<multiset<bool>> := multiset{v119};
		v4 := |v120|;
		v1 := fm2(v117.f0 && !false, globalState);
		if (v2) {
			var v121: seq<set<int>> := [{-218, v4, -0x208, DC11(v1).cf15, v1}];
			v117.f0 := v121 == v121;
			var v122: seq<int> := [v3, -v3];
			v122 := v122;
			v8 := 'o';
			var v123 := new C0(v2);
			v2 := v2;
		} else {
			var v124 := new C0(v117.f0);
			var v125 := new C0(v2);
			var v126: array<string> := new string[17];
			v126[safeIndex(692, v126.Length)] := seq(abs(0x15b), i10  => (v8));
			v126[safeIndex(692, v126.Length)] := v7;
			v117.f0 := v117.f0;
			var v127 := new C0(v2);
		}
		
	}
	
	var v128 := new C0(v2);
	if (v1 < (v4 - v5)) {
		var v129: map<bool, bool> := map[v2 := v2];
		var v130: multiset<map<bool, bool>> := multiset{v129};
		v2 := v71.fm6(-v5, v130 + v130, --v4, globalState);
		v8, v7, v4 := v10.cf0, seq(abs(-426), i11  => (v8)), 406;
		var v131: map<char, int> := map[fm0(v7, v3, |v9|, globalState) := |v7|];
		var v132: set<map<char, int>> := {v131};
		v132 := v132;
		v2 := v2;
		var v133: map<array<bool>, int> := map[v0 := 940];
		var v134: multiset<bool> := multiset{v2, v2};
		var v135, v136, v137, v138 := m0(v133, (DC3(v134, !v2, v2).(cf3 := v134)).cf5, globalState);
	} else {
		var v139: array<map<bool, bool>> := new map<bool, bool>[16];
		v139 := new map<bool, bool>[2](i12 => map[v2 := v2] + map[v2 := v2]);
		var v140: seq<set<bool>> := [v42];
		var v141: map<int, seq<set<bool>>> := map[v5 := v140];
		v141 := v141[safeModuloInt(v4, v4) := v140 + v140];
		var v142: map<int, string> := map[v3 := v7];
		var v143: seq<string> := [v7, v7, "ffchstvvi" + v7];
		v72, v142, v4, v7, v143 := v72, v142, v128.fm5(v5, false, globalState), v7, (seq(abs(-0x372), i13  => (v7))) + v143;
		v90 := v90;
		var v144: array<int> := new int[4](i14 => safeDivisionInt(i14, |v12|));
		v144[safeIndex(620, v144.Length)] := if (false) then -v4 else v5;
		var v145: multiset<string> := multiset{v7, DC9("stnwqxv").cf13, seq(abs(0x3d), i15  => (v8))};
		v2, v144[safeIndex(620, v144.Length)], v5 := (v145 - v145) >= v145, v5 * v3, v90.fm5(v4, !v2, globalState);
	}
	
	var i16 := 0;
	while (v72[safeIndex(v1, |v72|)])
		decreases 100 - i16
	{
		if (i16 >= 100) {
			break;
		}
		
		i16 := i16 + 1;
		v128 := v90;
		var v146: map<int, bool> := map[v4 * v128.fm5(|v42|, v2, globalState) := v2];
		v146 := v146;
		var v147: multiset<int> := multiset{v4, 0x7a};
		var v148: seq<int> := [0xb0, v5, -0xaa, fm2(fm1(v7, 943, {v2}, globalState), globalState)];
		v2, v8, v4 := v147 >= multiset(v148), v8, if (!false || v2) then safeDivisionInt(v5, v5) else |v7|;
		v5 := |(fm8(v2, v2, |v72|, globalState).(cf13 := v7)).cf13|;
	}
	v5 := -v1;
	var v149: multiset<map<bool, bool>> := multiset{map[v2 := v2]};
	var v150: map<bool, bool> := map[v2 := v2];
	if (v128.fm6(|v7|, v149 + v149[v150 := abs(v4)], 342, globalState)) {
		v0, v3, v2 := v0, -(v5 + v5), v2;
		v7 := seq(abs(0x20a), i17  => (v8));
		if (v2) {
			v0 := v0;
			v3 := safeDivisionInt(-0x34d, v4);
			v7 := v7;
			var v151 := new C0(!v2);
			var v152: array<D0> := new D0[16](i18 => v10);
			v152[safeIndex(273, v152.Length)] := v10;
			var v153: seq<int> := [v5];
			var v154: seq<seq<int>> := [v153];
			v152[safeIndex(273, v152.Length)] := fm10(v2, -v4, "tdbylyj", v154, globalState).(cf0 := v8);
		} else {
			v0[safeIndex(457, v0.Length)] := v2;
			v0[safeIndex(457, v0.Length)] := v2;
			var v155 := new C0(v2 <== v0[safeIndex(457, v0.Length)]);
			var v156: set<char> := {v8};
			var v157: map<bool, set<char>> := map[if (v2) then v72[safeIndex(v1, |v72|)] else v2 := {v8, v8, v8} + v156];
			var v158: set<int> := {v5};
			var v159: multiset<bool> := multiset{v0[safeIndex(457, v0.Length)]};
			var v160 := DC3(v159, v2, v2);
			v4, v0[safeIndex(457, v0.Length)], v5, v157 := safeDivisionInt(v1, -0x3c7 - v1), if (v2) then v2 else v0[safeIndex(457, v0.Length)], safeModuloInt(v4, |v158| * |map[v7 := v160.cf4]|), v157 + (fm11(globalState) + v157);
			var v161: array<bool> := new bool[15];
			var v162: map<array<bool>, int> := map[v161 := v155.fm5(v5, v0[safeIndex(457, v0.Length)], globalState)];
			var v163: map<C0, map<array<bool>, int>> := map[v128 := v162];
			var v164, v165, v166, v167 := m0(if (v90 in v163) then v163[v90] else v162, v2, globalState);
			var v168: array<seq<bool>> := new seq<bool>[3](i19 => v72);
			var v169: multiset<C0> := multiset{v128};
			v0[safeIndex(457, v0.Length)], v168, v3 := (multiset{v128, v128} - v169) <= v169, if (!v0[safeIndex(457, v0.Length)]) then v168 else v168, v4;
		}
		
		if (v128.fm6(|v42|, v149, v4, globalState)) {
			v5 := 826;
			var v170: multiset<int> := multiset{v4};
			var v171: map<bool, array<bool>> := map[v2 := v0];
			var v172: array<int> := new int[18] [v5 - v1, |v12|, |(multiset{v4})[v3 := abs(v4)]|, v5, safeDivisionInt(|v7|, |v42|), v4, -(v3 * v5), if (v2) then -|v170| else v5, |(map[v2 := v0] + v171)|, v3, v4, v4, v4, -0x36c, v5, |{v2, v2, v2}|, -v90.fm5(0x3b3, v2, globalState), v4];
			v172[safeIndex(493, v172.Length)] := v5;
			v172[safeIndex(493, v172.Length)] := -0x49;
			var v173, v174, v175, v176 := m0(map[v0 := v5], v2, globalState);
			var v177: multiset<bool> := multiset{v2};
			v3 := v174 * |v177|;
			v1 := v128.fm5(v172[safeIndex(493, v172.Length)] * v175, false, globalState);
		} else {
			v5 := v5 + v4;
			var v178 := DC6(-|v72|, v4, v2);
			v2 := v178.cf10;
			var v179: map<array<bool>, int> := map[v0 := v3];
			var v180, v181, v182, v183 := m0(v179 + v179, if (v2 in v150) then v150[v2] else v2, globalState);
			var v185 := DC9(v7);
			var v186: map<int, D5> := map[417 := v185];
			var v187: map<int, char> := map[|v72| := v8];
			var v188: seq<map<int, char>> := [map v184 : int | v184 in v186 :: (v184 * v181) := (v8), v187 + v187, v187 + v187, v187];
			var v189: array<array<int>> := new array<int>[27];
			var v190: array<int> := new int[17](i20 => safeModuloInt(i20, v1));
			v189[safeIndex(693, v189.Length)] := v190;
			var v191: map<set<bool>, bool> := map[v42 := v2];
			v188, v71, v189[safeIndex(693, v189.Length)], v182, v2 := if (false) then v188 else v188, v90, v190, v181 + -|v7|, v191 == v191;
			v181, v180, v8, v0 := -v3 + 0xe5, v178.cf8, v8, v0;
		}
		
		var v192: map<array<bool>, int> := map[v0 := |v72|];
		var v193, v194, v195, v196 := m0(v192, v2, globalState);
	} else {
		v2 := v2;
		var v197: array<int> := new int[4];
		v197[safeIndex(942, v197.Length)] := |("odeh" + v7)|;
		v197[safeIndex(942, v197.Length)] := v5;
		var v198 := DC12(fm12(true, v8, v2, v1, globalState));
		var v199: map<array<bool>, int> := map[v0 := |v198.cf16|];
		var v200, v201, v202, v203 := m0(v199, v2, globalState);
		v2 := !v2;
		var v204: array<array<bool>> := new array<bool>[2];
		v1, v204, v2, v2 := v1, v204, v2 ==> v2, v2;
	}
	
	v0[safeIndex(901, v0.Length)] := v2;
	v0[safeIndex(901, v0.Length)] := fm1(v7, |v7|, v42, globalState);
	var v205: seq<int> := [|v7|, |v72|];
	var v206: multiset<seq<int>> := multiset{v205};
	var v207 := DC12(v206);
	var v208 := DC12(v207.cf16);
	var v209 := DC14(v208);
	match (v209.(cf18 := v208)) {
		case DC13(cf17) =>
			v4 := -v3;
			cf17 := cf17;
			var v210: map<bool, C0> := map[cf17 := v71];
			var v211: array<C0> := new C0[18] [v71, v90, v90, v90, v128, v71, v71, v128, v71, v71, v128, if (v2 in v210) then v210[v2] else v128, v128, v128, v90, v90, v90, v71];
			v211[safeIndex(917, v211.Length)] := v90;
			v211[safeIndex(917, v211.Length)] := v128;
			v1 := safeDivisionInt(v3, v5) * (|v72| - 283);
		case DC12(cf16) =>
			v4, v3, v2, v8, v0[safeIndex(901, v0.Length)] := v5, v4, ("xya" + "kbshdbhf") != (seq(abs(-774), i21  => (v8))), v8, v2;
			var v212: set<int> := {v3};
			v0[safeIndex(901, v0.Length)] := v71.fm6(safeDivisionInt(v5, |v7|), v149, -(v5 - |v212|), globalState);
			var v213: multiset<bool> := multiset{v2};
			var v214 := DC3(v213, v2, v0[safeIndex(901, v0.Length)]);
			match (if (true) then v214 else v214) {
				case DC2(cf2) =>
					v5 := --v4;
					var v215, v216, v217, v218 := m0(map[v0 := v3], v2, globalState);
					var v219: map<array<bool>, int> := map[v0 := |v205|];
					var v220, v221, v222, v223 := m0(v219 + v219, v2, globalState);
					var v224 := DC6(safeDivisionInt(|v213|, v5), 225 - -v4, v0[safeIndex(901, v0.Length)]);
					v224 := v224.(cf9 := |(v7 + v7)|);
				case DC3(cf3, cf4, cf5) =>
					v72 := v72 + v72;
					var v225: map<array<bool>, int> := map[v0 := 0x2b3];
					var v226: seq<C0> := [v71];
					var v227: map<seq<C0>, bool> := map[v226 := v71.fm6(v5, v149, v1, globalState)];
					var v228: map<map<seq<C0>, bool>, bool> := map[v227 := false];
					var v229, v230, v231, v232 := m0(v225, if (v227 in v228) then v228[v227] else v0[safeIndex(901, v0.Length)], globalState);
					v7 := fm7(v1, v229, globalState) + v7;
					v4 := -v4;
				case DC1(cf1) =>
					var v233: map<array<bool>, int> := map[v0 := |v7|];
					var v234: map<C0, int> := map[v90 := v4];
					var v235, v236, v237, v238 := m0(v233 + v233[v0 := v4], v1 < |v234|, globalState);
					var v239, v240, v241, v242 := m0(v233, v0[safeIndex(901, v0.Length)], globalState);
					v42 := fm13(v7 <= v7, !(v2 in v213), false, globalState);
					var v243 := DC11(v241);
					var v244: set<D6> := {v243, DC11(v240), v243};
					var v245: T0 := new C0(v2);
					var v246: seq<T0> := [v245, v245];
					var v247: array<set<D7>> := new set<D7>[2];
					v244, v246, v247 := v244, [v245, v245] + ([v245] + v246), v247;
			}
			
			var v248: map<bool, array<bool>> := map[v2 := v0];
			v248 := v248['m' in v7 := v0];
		case DC14(cf18) =>
			if (v0[safeIndex(901, v0.Length)]) {
				v4 := v5;
				var v249: set<seq<int>> := {v205, v205, v205, v205};
				v0[safeIndex(901, v0.Length)] := !(fm14(|v9|, globalState) >= v249);
				v8 := v8;
				var v250: map<array<bool>, int> := map[v0 := v5];
				var v251, v252, v253, v254 := m0(v250, v0[safeIndex(901, v0.Length)], globalState);
				var v255: set<int> := {if (!!v2 in v9) then v9[!!v2] else -|v7|};
				var v256: map<seq<bool>, int> := map[[!v0[safeIndex(901, v0.Length)], v0[safeIndex(901, v0.Length)], v2] := |v255|];
				var v257: seq<seq<bool>> := [v72];
				v256 := v256[v257[safeIndex(v251, |v257|)] + v72 := v5];
			} else {
				v0[safeIndex(901, v0.Length)] := v2;
				var v258: multiset<string> := multiset{v7, v7};
				var v259: map<seq<int>, string> := map[seq(abs(0x29d), i22  => (v1)) := seq(abs(23), i23  => (v8))];
				v0 := new bool[19] [if (v0[safeIndex(901, v0.Length)]) then v2 else v0[safeIndex(901, v0.Length)], v2, v72 <= [v2], v0[safeIndex(901, v0.Length)], v0[safeIndex(901, v0.Length)], v2, v72[safeIndex(|v258|, |v72|)], true <==> v0[safeIndex(901, v0.Length)], v2, v2, v0[safeIndex(901, v0.Length)], v0[safeIndex(901, v0.Length)], v2, v2, v0[safeIndex(901, v0.Length)], fm1(if (v205 in v259) then v259[v205] else v7, v3, v42, globalState), DC13(v0[safeIndex(901, v0.Length)]).cf17, v0[safeIndex(901, v0.Length)], !v0[safeIndex(901, v0.Length)]];
				v2 := !!(v7 < v7);
				v8 := v8;
				var v260: set<int> := {324, v1};
				var v261 := new C0(v260 !! v260);
			}
			
			v0 := new bool[15](i24 => false);
			var v262: array<map<bool, bool>> := new map<bool, bool>[20];
			v262[safeIndex(131, v262.Length)] := v150 + v150;
			v262[safeIndex(131, v262.Length)] := v150 + v150;
			var v263: map<seq<int>, int> := map[v205 + v205 := v5];
			v263 := v263;
	}
	
	print v0[0], "\n";
	print v0[1], "\n";
	print v0[2], "\n";
	print v0[3], "\n";
	print v0[4], "\n";
	print v0[5], "\n";
	print v0[6], "\n";
	print v0[7], "\n";
	print v0[8], "\n";
	print v0[9], "\n";
	print v0[10], "\n";
	print v0[11], "\n";
	print v0[12], "\n";
	print v0[13], "\n";
	print v0[14], "\n";
	print v1, "\n";
	print v2, "\n";
	print v3, "\n";
	print v4, "\n";
	print v5, "\n";
	print v6 == {{true}}, "\n";
	print v7 == ['k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k', 'k'], "\n";
	print v8, "\n";
	print v9 == map[true := 365], "\n";
	print v10.cf0, "\n";
	print v11[0], "\n";
	print v11[1], "\n";
	print v11[2], "\n";
	print v11[3], "\n";
	print v11[4], "\n";
	print v11[5], "\n";
	print v11[6], "\n";
	print v11[7], "\n";
	print v11[8], "\n";
	print v11[9], "\n";
	print v11[10], "\n";
	print v11[11], "\n";
	print v11[12], "\n";
	print v11[13], "\n";
	print v11[14], "\n";
	print v11[15], "\n";
	print v11[16], "\n";
	print v11[17], "\n";
	print v11[18], "\n";
	print v11[19], "\n";
	print v12 == map[0 := 265], "\n";
	print v42 == {false}, "\n";
	print v72 == [true], "\n";
	print i16, "\n";
	print v149 == multiset{map[false := false]}, "\n";
	print v150 == map[false := false], "\n";
	print v205 == [426, 1], "\n";
	print v206 == multiset{[426, 1]}, "\n";
	print v207.cf16 == multiset{[426, 1]}, "\n";
	print v208.cf16 == multiset{[426, 1]}, "\n";
	print v209.cf18.cf16 == multiset{[426, 1]}, "\n";
}