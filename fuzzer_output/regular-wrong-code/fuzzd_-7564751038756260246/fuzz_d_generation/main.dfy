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
datatype D0 = DC0(cf0: multiset<bool>)
datatype D1 = DC2(cf2: set<string>, cf3: char) | DC1(cf1: bool)
datatype D2 = DC3(cf4: seq<bool>)
datatype D3 = DC5(cf6: bool, cf7: int) | DC4(cf5: string)
datatype D4 = DC7(cf9: bool, cf10: C0, cf11: bool, cf12: bool) | DC8 | DC6(cf8: seq<D3>)
datatype D5 = DC9(cf13: array<int>)
datatype D6 = DC11(cf15: int, cf16: int) | DC12(cf17: C0, cf18: bool) | DC13(cf19: int) | DC10(cf14: seq<int>)
datatype D7 = DC15(cf21: int, cf22: int) | DC16(cf23: int, cf24: int, cf25: bool, cf26: bool, cf27: int) | DC14(cf20: map<C0, bool>) | DC17(cf28: D7)
datatype D8 = DC18(cf29: set<int>)
class GlobalState {
	var f0 : int
	const f1 : int
	var f2 : string
	const f3 : array<char>
	const f4 : bool
	const f5 : bool
	const f6 : array<bool>
	const f7 : int
	var f8 : int
	const f9 : set<char>
	const f10 : int
	constructor (f0 : int, f1 : int, f2 : string, f3 : array<char>, f4 : bool, f5 : bool, f6 : array<bool>, f7 : int, f8 : int, f9 : set<char>, f10 : int) {
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
	}
	
}

class C0 {
	constructor () {
	}
	
}

function fm0(globalState: GlobalState): bool {
	DC16(0xa3, 125, true, false, 0x114).cf26
}
function fm1(p0: char, p1: bool, p2: bool, globalState: GlobalState): set<int> {
	({0x183} * {-0x325}) + DC18({0x270, DC13(510).cf19}).cf29
}
function fm2(p0: map<int, int>, p1: bool, p2: int, globalState: GlobalState): char {
	'f'
}
function fm3(globalState: GlobalState): seq<bool> {
	([true] + [true]) + [true, false]
}
function fm4(p0: bool, p1: int, globalState: GlobalState): int {
	549
}
function fm5(p0: int, p1: bool, globalState: GlobalState): map<int, int> {
	map[|[false, true, !true, !false, false]| := --safeModuloInt(|[true]|, 0xc0)]
}
function fm6(p0: char, p1: string, p2: bool, globalState: GlobalState): D0 {
	DC0(multiset([!false]) - multiset([!true]))
}
function fm7(globalState: GlobalState): map<int, bool> {
	map[315 := false] + (map v0 : int | (0x20d <= v0) && (v0 < 56) :: (v0 - |multiset{false}|) := (!!false))
}
function fm8(p0: int, p1: map<int, int>, p2: int, globalState: GlobalState): string {
	"drfsgt" + ("b" + "wtfucath")
}
function fm9(globalState: GlobalState): D3 {
	if (false) then DC4(seq(abs(0x360), i0  => ('n'))) else DC4(seq(abs(-0x193), i1  => ('x')))
}
function fm10(p0: bool, p1: int, p2: int, globalState: GlobalState): map<map<int, int>, bool> {
	match if (false) then DC15(--|map[0x3d7 := |map[false := true]|]|, |map[false := multiset{true}]|) else DC15(|[[!false]]|, |"ujipcla"|) {
		case DC15(cf21, cf22) => map v0 : map<int, int> | v0 in multiset([map[|"htgdju"| := cf21], map[cf21 := cf22], map[cf21 := cf21]]) :: (v0) := (true)
		case DC16(cf23, cf24, cf25, cf26, cf27) => map[map[cf23 := cf24] := cf25]
		case DC14(cf20) => map[map v1 : int | v1 in map[0x33 := true] :: (v1 * |{false, false}|) := (0x1a4) := false] + map[map v2 : int | v2 in multiset{|(map v3 : D2 | v3 in multiset{DC3([true, true])} :: (v3) := (-0x21e))|, |(map v4 : int | v4 in [|(seq(abs(0x2a0), i0  => ('a')))|] :: (v4 + 0x195) := (190))|, |(seq(abs(344), i1  => (83)))|, 0x15e} :: (safeDivisionInt(v2, |{!true}|)) := (-|map[false := |[true, false, true]|]|) := !true]
		case DC17(cf28) => map v5 : map<int, int> | v5 in [map[|[true]| := |{0x292, |(seq(abs(528), i2  => ('j')))|, -688, -524, -555}|], map[|multiset([|"fvgbw"|, |map[true := 0x18a]|])| := -341]] :: (v5) := (false)
	}
}
function fm11(p0: int, p1: int, p2: map<bool, int>, globalState: GlobalState): map<bool, int> {
	map[false := -|(set v0 : int | (0x1ec <= v0) && (v0 < 0x349) :: (v0 - -|[0x100]|))|] + (map[!!true := |map[|"orpdwcbhu"| := |(seq(abs(893), i0  => ('g')))|]|] + map[false := |(map v1 : int | v1 in [|multiset{false, false}|] :: (v1 + 0x1bd) := (|map[true := 908]|))|])
}
method m0(p0: bool, p1: int, p2: char, p3: bool, globalState: GlobalState) {
	var i0 := 0;
	while (p3)
		decreases 100 - i0
	{
		if (i0 >= 100) {
			break;
		}
		
		i0 := i0 + 1;
		var v0: map<bool, bool> := map[false := p0];
		globalState.f8 := |v0[p3 := p0]|;
		var v1 := new C0();
		var v2 := new C0();
		var v3 := DC5(p3, p1);
		var v4 := "bkerwni";
		var v5: set<bool> := {false, p0};
		var v6: map<C0, bool> := map[v2 := p3];
		var v7: seq<map<C0, bool>> := [v6];
		var v8: array<bool> := new bool[18];
		var v9: map<int, array<bool>> := map[|[p0, p3]| := v8];
		var v10: array<int> := new int[21] [p1, safeDivisionInt(p1, p1), v3.cf7, p1, -p1, p1, p1, |v4|, |v5|, p1, |v4|, p1, 0x20e, p1, p1, safeDivisionInt(p1, |v7[safeIndex(p1, |v7|) := v6]|), p1, |v4|, p1, |(v9 + v9)|, p1];
		v10[safeIndex(786, v10.Length)] := |v0[p3 := true]|;
		v10[safeIndex(786, v10.Length)] := p1;
	}
	var v11: seq<int> := [-p1];
	var v12 := "owdvbmnf";
	var v13: C0 := new C0();
	var v14: seq<C0> := [v13, v13];
	var v15: seq<seq<C0>> := [v14, v14, v14, [v13, v13, v13]];
	var v16: map<int, bool> := map[v11[safeIndex(|v12|, |v11|)] := v14 in v15];
	var i1 := 0;
	while (if (p1 in v16) then v16[p1] else p3)
		decreases 100 - i1
	{
		if (i1 >= 100) {
			break;
		}
		
		i1 := i1 + 1;
		if (!!!p0) {
			var v17: seq<bool> := [p0, p0];
			var v18: map<seq<bool>, C0> := map[v17 := v13];
			var v19: map<C0, bool> := map[if (v17 in v18) then v18[v17] else v13 := v11 < v11];
			v19 := map[v13 := p3];
			globalState.f8 := fm4(p3, p1, globalState) * p1;
			var v20: array<bool> := new bool[22](i2 => |multiset{p3}| < v11[safeIndex(p1, |v11|)]);
			var v21: multiset<bool> := multiset{p3};
			v20[safeIndex(304, v20.Length)] := v21 != v21;
			var v22: multiset<int> := multiset{p1};
			var v23: map<int, int> := map[-0x10e := p1];
			var v24: set<int> := {357};
			var v25: map<C0, set<int>> := map[v13 := v24];
			v20[safeIndex(304, v20.Length)] := -(v11[safeIndex(|v22|, |v11|)] - -(if (p1 in v23) then v23[p1] else |v25|)) > |v17|;
			globalState.f8 := 0xb6;
			var v26: set<C0> := {v13, v13};
			var v27: map<int, set<C0>> := map[p1 := v26];
			v20[safeIndex(304, v20.Length)] := v26 !! (if (p1 in v27) then v27[p1] else v26);
		} else {
			var v28 := false;
			globalState.f8, v28, v12, v28 := v11[safeIndex(0x88 + p1, |v11|)], !v28, v12, p2 in v12;
			var v29: map<bool, bool> := map[!p3 := p0];
			var v30: seq<map<bool, bool>> := [v29, map[p3 := v28]];
			v30 := (v30 + v30) + v30;
			var v31: map<C0, char> := map[v13 := p2];
			v31 := v31[v13 := p2];
			var v32 := new C0();
			var v33: multiset<char> := multiset{p2, 'q', p2, 'f', p2};
			var v34: set<int> := {p1};
			var v35: map<int, int> := map[p1 := p1];
			var v36: array<int> := new int[19] [p1, 293, p1, if (p2 in v33) then v33[p2] else p1, |(v11 + v11)|, p1, -fm4(p0, p1, globalState), p1, p1, |map[v32 := |v11|]|, -DC5(false, 0x296).cf7, p1, 185, |v34| + p1, safeModuloInt(p1, p1), |(seq(abs(0x35e), i3  => (0x146)))|, p1, safeModuloInt(-257, p1), |(v35 + map[p1 := p1])|];
			v36[safeIndex(249, v36.Length)] := 506;
			var v37: map<bool, string> := map[p0 := v12];
			v36[safeIndex(249, v36.Length)] := p1 + safeDivisionInt(0xf, |v37|);
		}
		
		var v38 := false;
		v38 := p0;
		var v39 := new C0();
		var v40: array<int> := new int[26];
		var v41: map<array<int>, string> := map[v40 := v12 + v12[safeIndex(p1, |v12|) := 'e']];
		var v42 := DC5(true, p1);
		v41, v38, globalState.f0, globalState.f0 := v41 + v41, p3, v42.cf7, fm4(v38, -p1, globalState);
	}
	match (fm9(globalState)) {
		case DC5(cf6, cf7) =>
			var v43 := new C0();
			var v44 := 'g';
			v44 := if (p0) then fm2(map[p1 := -p1], !!cf6, cf7, globalState) else v44;
			cf6 := p0;
			cf6 := if (!cf6) then p3 else p0;
		case DC4(cf5) =>
			var v45: map<seq<int>, string> := map[v11 := cf5];
			var v46 := DC10([p1]);
			v45 := v45[(v11[safeIndex(|v46.cf14|, |v11|) := v11[safeIndex(p1, |v11|)]])[safeIndex(p1, |v11[safeIndex(|v46.cf14|, |v11|) := v11[safeIndex(p1, |v11|)]]|) := p1] := "imvdwrgjx"];
			var v47 := new C0();
			globalState.f0 := p1;
			var v48: array<int> := new int[10];
			v48[safeIndex(289, v48.Length)] := |cf5| + |(map v49 : int | (0x190 <= v49) && (v49 < 0x25a) :: (safeDivisionInt(v49, p1)) := (p2))|;
			v48[safeIndex(289, v48.Length)] := p1;
	}
	
	var v50: multiset<int> := multiset{p1};
	var v51: seq<bool> := [if (!p3) then !p3 else p0, v50 < v50];
	v51 := v51;
	var v52 := new C0();
	for i4 := p1 to -(p1 + p1) {
		var v53: array<int> := new int[3](i5 => i5 * p1);
		v53[safeIndex(646, v53.Length)] := 0x2ef - p1;
		var v54: array<bool> := new bool[6];
		var v55: map<array<bool>, bool> := map[v54 := v51[safeIndex(p1, |v51|)]];
		globalState.f0, globalState.f2, v53[safeIndex(646, v53.Length)], v55, globalState.f8 := -((if (p0) then p1 else i4) + safeModuloInt(i4, i4)), if (p3) then v12 else ("br")[safeIndex(p1, |"br"|) := p2] + v12, safeModuloInt(i4, p1), v55, safeDivisionInt(i4, p1 - p1);
		if (false) {
			v54[safeIndex(643, v54.Length)] := p3;
			v54[safeIndex(643, v54.Length)] := if (!!p3) then !(p1 != i4) else false && p0;
			var v56: map<bool, bool> := map[p0 := v54[safeIndex(643, v54.Length)]];
			var v57: map<int, map<bool, bool>> := map[-v53[safeIndex(646, v53.Length)] := v56];
			v57 := v57[-i4 := v56];
			v54[safeIndex(643, v54.Length)] := p3;
			var v58: map<map<map<int, int>, bool>, int> := map[fm10(p0, i4, -914, globalState) := p1];
			var v60: map<int, int> := map[0x2e0 := p1];
			var v61: set<string> := {"dxrmfw", v12};
			var v62 := DC2(v61, p2);
			var v63: map<map<int, int>, D1> := map[v60 := v62];
			v58 := v58[map v59 : map<int, int> | v59 in v63 :: (v59) := (p0) := p1];
			v54[safeIndex(643, v54.Length)] := p3;
		} else {
			var v64 := true;
			v64 := v53[safeIndex(646, v53.Length)] != i4;
			v64 := v11 < (v11 + v11[safeIndex(v53[safeIndex(646, v53.Length)], |v11|) := p1]);
			var v65: seq<multiset<C0>> := [multiset(v14) * multiset{v52}, multiset(if (p0) then v14 else v14), multiset((if (p0) then [v52] else v14)[safeIndex(p1, |(if (p0) then [v52] else v14)|) := v13])];
			v13, v52, v65 := v13, v13, v65;
			var v66 := DC8();
			v66 := v66;
			var v67: array<seq<map<bool, int>>> := new seq<map<bool, int>>[11];
			var v68: map<bool, int> := map[p0 := |multiset{v51[safeIndex(p1, |v51|)]}|];
			var v69: seq<map<bool, int>> := [fm11(|v11|, v53[safeIndex(646, v53.Length)], v68, globalState)];
			v67[safeIndex(913, v67.Length)] := v69;
			var v70: set<int> := {p1};
			v64, v67[safeIndex(913, v67.Length)], v53[safeIndex(646, v53.Length)], globalState.f0 := if (v12 <= v12) then v64 else p0, v69, v53[safeIndex(646, v53.Length)], safeDivisionInt(v53[safeIndex(646, v53.Length)], |v70|);
		}
		
		v54[safeIndex(400, v54.Length)] := v51[safeIndex(i4, |v51|)];
		var v71: multiset<array<int>> := multiset{v53, v53, v53};
		v54[safeIndex(400, v54.Length)] := v71 !! v71;
		var v72: map<bool, bool> := map[p0 <==> p3 := v50 < v50];
		var v73: map<C0, bool> := map[v13 := v54[safeIndex(400, v54.Length)]];
		var v74 := DC14(v73);
		var v75: map<multiset<int>, int> := map[v50 := |v74.cf20[v13 := v54[safeIndex(400, v54.Length)]]|];
		v72 := v72[|v75| < -p1 := v54[safeIndex(400, v54.Length)]];
	}
}
method Main() {
	var v0 := "si";
	var v1: array<char> := new char[11];
	var v2 := false;
	var v3: array<bool> := new bool[3] [!v2, !v2, v2];
	var v4 := 'j';
	var v5: set<char> := {v4};
	var globalState := new GlobalState(-0xa4, 378, v0, v1, true, true, v3, 223, 0x247, v5, 782);
	v2 := !(!v2 && v2);
	var v6 := 520;
	var v7: map<array<char>, int> := map[if (v2) then v1 else v1 := if (fm0(globalState)) then v6 else v6];
	var v8: seq<bool> := [v2];
	var v9: set<int> := {v6, v6, |v8|};
	v7 := v7[v1 := |v9| - v6];
	forall i0 | 0 <= i0 < v3.Length {
		v3[i0] := DC0(multiset{false}).cf0 >= multiset{false, v2, false, v2};
	}
	var v10: multiset<bool> := multiset{!v2, v2};
	var v11 := DC0(v10);
	match (v11) {
		case DC0(cf0) =>
			if (!(!(v9 > v9) <== v2)) {
				m0(if (v2) then fm0(globalState) else v8[safeIndex(v6, |v8|)], |multiset(v8)|, v4, v2, globalState);
				var v12: set<bool> := {false};
				var v13: seq<int> := [|v8|];
				var v14: map<set<int>, int> := map[v9 - {|v12|, -0x134} := v13[safeIndex(v6, |v13|)] + v6];
				v2, globalState.f8, globalState.f8, v14 := fm0(globalState), v6, v6, v14[fm1('d', v2, v2, globalState) := 0x50];
				v2 := v2;
				globalState.f0 := v6;
				var v15: multiset<seq<bool>> := multiset{v8, [v2, v2]};
				var v16: map<multiset<seq<bool>>, string> := map[v15 := v0];
				var v17: seq<multiset<seq<bool>>> := [v15, v15, v15];
				v16 := v16[v17[safeIndex(v6, |v17|)] := v0];
			} else {
				v1[safeIndex(178, v1.Length)] := v4;
				var v18: map<int, int> := map[v6 := v6];
				v1[safeIndex(178, v1.Length)] := fm2(v18, v2, v6, globalState);
				var v19: set<bool> := {v2};
				var v20: map<char, string> := map[v4 := seq(abs(0x1ae), i1  => (v4))];
				var v21: map<set<bool>, bool> := map[v19 := v20 != v20];
				v21 := v21[v19 := v2];
				globalState.f0 := v6;
				globalState.f2 := v0;
				m0(v8[safeIndex(v6, |v8|)], v6, v4, v2, globalState);
			}
			
			globalState.f0 := v6;
			var v22 := DC1(true);
			if (v22.cf1) {
				var v23: array<array<bool>> := new array<bool>[18];
				var v24: map<array<bool>, bool> := map[v3 := v2];
				var v25: map<array<array<bool>>, int> := map[v23 := v6 * |v24|];
				v25 := v25[v23 := v6];
				v8 := fm3(globalState) + v8;
				globalState.f0 := v6 - v6;
				var v26 := new C0();
				var v27: array<array<seq<char>>> := new array<seq<char>>[11];
				var v28: map<bool, seq<char>> := map[v2 := v0];
				var v29: map<C0, seq<char>> := map[v26 := if (!v2 in v28) then v28[!v2] else seq(abs(0x17), i2  => (v4))];
				var v30: array<seq<char>> := new seq<char>[8] [[v4, 'g', v4, v4], if (v26 in v29) then v29[v26] else seq(abs(-93), i3  => (v4)), v0, v0, v0, seq(abs(123), i4  => ('t')), v0, v0];
				v27[safeIndex(264, v27.Length)] := v30;
				v1[safeIndex(325, v1.Length)] := v4;
				var v31: array<C0> := new C0[27];
				var v32: map<array<C0>, array<seq<char>>> := map[v31 := v30];
				v27[safeIndex(264, v27.Length)], v23, v1[safeIndex(325, v1.Length)], globalState.f0 := if (v31 in v32) then v32[v31] else v30, v23, v4, v6;
			} else {
				var v33: array<array<char>> := new array<char>[27];
				v33[safeIndex(214, v33.Length)] := v1;
				v33[safeIndex(214, v33.Length)], globalState.f8 := v1, -(v6 + v6);
				var v34: map<bool, int> := map[v2 := v6];
				v34 := v34[v2 := v6];
				var v35: map<bool, bool> := map[v2 := v2];
				var v36: seq<set<int>> := [v9];
				var v37: map<int, int> := map[v6 := |v0|];
				var v38: multiset<char> := multiset{v4};
				v9, v6, globalState.f0, v2, v35 := v36[safeIndex(v6 * v6, |v36|)], fm4(v2, v6, globalState), safeDivisionInt(v6, v6), (|v37| == |v38|) ==> (v2 && v2), (v35 + v35[v2 := true])[fm0(globalState) := v2];
				var v39: seq<map<int, int>> := [v37];
				var v40: map<int, bool> := map[v6 := v2];
				var v41: map<int, bool> := map[|{if (v6 in v40) then v40[v6] else v2}| := v2];
				var v43: array<map<int, int>> := new map<int, int>[10] [v37, v39[safeIndex(0x2b, |v39|)], fm5(0xb7, v2, globalState), v37 + v37, v37, v37, fm5(|"aywqiw"|, v2, globalState) + v37, map[0x179 := |(set v42 : int | v42 in v41 :: (v42 * 51))|], v37 + fm5(v6, v2, globalState), v37];
				v43[safeIndex(809, v43.Length)] := map[v6 := |v8|];
				var v44: set<string> := {"nagnmac"};
				var v45 := DC2(v44, v4);
				var v46: array<C0> := new C0[7];
				var v47: map<D0, array<C0>> := map[DC0(cf0) := if (v2) then v46 else v46];
				var v48: map<int, map<int, int>> := map[v6 := v37];
				v43[safeIndex(809, v43.Length)], v45, v47, v2 := v37[v6 := v6] + (if (v6 in v48) then v48[v6] else v37), v45, v47 + (v47 + map[v11 := v46]), v6 >= v6;
				v2 := false;
			}
			
			m0(v2, -v6, v4, -278 >= v6, globalState);
	}
	
	v2 := v2;
	v3[safeIndex(849, v3.Length)] := v2;
	v3[safeIndex(849, v3.Length)] := !v2;
	var v49: seq<int> := [fm4(v3[safeIndex(849, v3.Length)], v6, globalState)];
	var v50: map<seq<int>, int> := map[v49 := v6];
	var v51: map<bool, int> := map[v3[safeIndex(849, v3.Length)] := |v0[safeIndex(0x7a, |v0|) := v4]|];
	var v52: map<map<seq<int>, int>, map<bool, int>> := map[v50 := v51];
	v52 := v52[map[v49 := v6] + (map v53 : seq<int> | v53 in v50 :: (v53) := (|v9|)) := map[v3[safeIndex(849, v3.Length)] := v6]];
	if (v2 || (v6 >= v6)) {
		globalState.f8 := v6;
		var v54: map<array<bool>, bool> := map[v3 := v2];
		v54 := (v54 + v54)[v3 := true];
		var v55: seq<string> := ["pnjsvda"];
		var v56 := DC3((v8[safeIndex(|v55|, |v8|) := v2])[safeIndex(|v0|, |v8[safeIndex(|v55|, |v8|) := v2]|) := false]);
		v8 := v56.cf4;
		var v57: C0 := new C0();
		var v58: seq<C0> := [v57];
		var v59: set<C0> := {v57};
		var v60: map<int, C0> := map[v6 := v57];
		var v61 := DC4(seq(abs(0xf4), i5  => (v4)));
		var v62: seq<D0> := [v11, fm6('i', v61.cf5, v2, globalState)];
		m0({v58[safeIndex(v6, |v58|)], v57, v57, v57, v57} !! v59, |(v60 + map[v6 := v57])|, v4, v62 != [v11, v11], globalState);
		var v63: array<int> := new int[3](i6 => i6 * v6);
		v63 := v63;
	} else {
		var v64: map<int, int> := map[v6 := v6];
		v64 := map[|v0| := v6] + map[v6 := v6];
		var v65 := DC4(v0);
		var v66: seq<D3> := [v65, v65];
		var v67: map<bool, seq<D3>> := map[fm0(globalState) := v66];
		var v68: map<bool, seq<D3>> := map[!v3[safeIndex(849, v3.Length)] := if (v3[safeIndex(849, v3.Length)] in v67) then v67[v3[safeIndex(849, v3.Length)]] else [v65]];
		var v69 := DC6(v66);
		v67 := v67[v3[safeIndex(849, v3.Length)] := v69.cf8];
		var v70: map<int, bool> := map[v6 := v2];
		v3[safeIndex(849, v3.Length)] := v70 != fm7(globalState);
		v3[safeIndex(849, v3.Length)] := !v2;
		v9 := v9;
	}
	
	v2 := v2;
	var v71: array<map<D2, bool>> := new map<D2, bool>[4];
	forall i7 | 0 <= i7 < v71.Length {
		v71[i7] := (map[DC3(v8) := v3[safeIndex(849, v3.Length)]] + (map v72 : D2 | v72 in (map v73 : D2 | v73 in (seq(abs(0x35e), i8  => (DC3([false])))) :: (v73) := (map[!true := v6])) :: (v72) := (v3[safeIndex(849, v3.Length)]))) + (if (v3[safeIndex(849, v3.Length)]) then map[DC3(v8) := !v3[safeIndex(849, v3.Length)]] else map[DC3(v8) := v3[safeIndex(849, v3.Length)]]);
	}
	var v74 := new C0();
	var v75: map<int, int> := map[v6 := v6];
	v0 := fm8(v6, v75, |v49|, globalState) + (v0 + v0);
	var v76: map<C0, int> := map[v74 := v6];
	v76 := v76[v74 := v6];
	v2 := v3[safeIndex(849, v3.Length)];
	var v77: array<int> := new int[23];
	v77 := v77;
	var v78: multiset<int> := multiset{v6};
	for i9 := fm4(false, |v8[safeIndex(79, |v8|) := DC1(v3[safeIndex(849, v3.Length)]).cf1]|, globalState) to |v78| * v6 {
		var v79: seq<array<bool>> := [v3, v3, v3];
		v79 := [v3];
		var v80 := DC9(v77);
		var v81: seq<array<int>> := [v77, v80.cf13, v77, v77];
		v81 := v81;
		var v82 := DC5(v2, v49[safeIndex(i9, |v49|)]);
		v3[safeIndex(849, v3.Length)] := v6 > (v82.cf7 * -0x318);
		v3[safeIndex(849, v3.Length)], v3[safeIndex(849, v3.Length)], globalState.f8, v9, v74 := !v2, (v10 * v10) < v10, |(([v3[safeIndex(849, v3.Length)]] + v8) + [v2, v3[safeIndex(849, v3.Length)]])|, {v6} - v9, v74;
	}
	print v0, "\n";
	print v1[2], "\n";
	print v1[6], "\n";
	print v2, "\n";
	print v3[0], "\n";
	print v3[1], "\n";
	print v3[2], "\n";
	print v4, "\n";
	print v5 == {'j'}, "\n";
	print globalState.f0, "\n";
	print globalState.f1, "\n";
	print globalState.f2, "\n";
	print globalState.f3[2], "\n";
	print globalState.f3[6], "\n";
	print globalState.f4, "\n";
	print globalState.f5, "\n";
	print globalState.f6[0], "\n";
	print globalState.f6[1], "\n";
	print globalState.f6[2], "\n";
	print globalState.f7, "\n";
	print globalState.f8, "\n";
	print globalState.f9 == {'j'}, "\n";
	print globalState.f10, "\n";
	print v6, "\n";
	print |v7|, "\n";
	print v8 == [true, true, false, false, true], "\n";
	print v9 == {520, 1}, "\n";
	print v10 == multiset{false, true}, "\n";
	print v11.cf0 == multiset{false, true}, "\n";
	print v49 == [549], "\n";
	print v50 == map[[549] := 520], "\n";
	print v51 == map[false := 2], "\n";
	print v52 == map[map[[549] := 520] := map[false := 2], map[[549] := 2] := map[false := 520]], "\n";
	print v71[0] == map[DC3([true, true, false, false, true]) := false, DC3([false]) := false], "\n";
	print v71[1] == map[DC3([true, true, false, false, true]) := false, DC3([false]) := false], "\n";
	print v71[2] == map[DC3([true, true, false, false, true]) := false, DC3([false]) := false], "\n";
	print v71[3] == map[DC3([true, true, false, false, true]) := false, DC3([false]) := false], "\n";
	print v75 == map[520 := 520], "\n";
	print |v76|, "\n";
	print v78 == multiset{520}, "\n";
}