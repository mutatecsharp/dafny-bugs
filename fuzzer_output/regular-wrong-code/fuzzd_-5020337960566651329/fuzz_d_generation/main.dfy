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
datatype D0 = DC1 | DC0(cf0: bool)
datatype D1 = DC3(cf2: char, cf3: int) | DC4(cf4: map<int, map<C0, bool>>, cf5: int) | DC2(cf1: C0)
datatype D2 = DC5(cf6: seq<int>)
datatype D3 = DC7(cf8: bool, cf9: bool) | DC6(cf7: map<array<bool>, array<C0>>)
datatype D4 = DC8(cf10: array<int>)
class GlobalState {
	constructor () {
	}
	
}

class C0 {
	const f0 : seq<int>
	constructor (f0 : seq<int>) {
		this.f0 := f0;
	}
	
	function fm5(p0: bool, p1: bool, p2: bool, p3: int, globalState: GlobalState): int {
		-344
	}
}

function fm0(globalState: GlobalState): set<bool> {
	{true, true} + {true}
}
function fm1(p0: bool, p1: bool, p2: int, globalState: GlobalState): bool {
	true
}
function fm2(p0: int, globalState: GlobalState): D0 {
	DC1()
}
function fm3(globalState: GlobalState): string {
	match DC0(true) {
		case DC1() => "ttioiues"
		case DC0(cf0) => "p"
	}
}
function fm4(p0: int, p1: int, p2: bool, p3: bool, globalState: GlobalState): int {
	-413
}
function fm6(p0: int, p1: bool, p2: bool, globalState: GlobalState): D2 {
	DC5(seq(abs(0x359), i0  => (|"meafobh"|)))
}
function fm7(p0: bool, globalState: GlobalState): D0 {
	match DC1() {
		case DC1() => DC0(true)
		case DC0(cf0) => DC0(cf0)
	}
}
method m0(p0: bool, p1: bool, globalState: GlobalState) returns (r0: seq<int>, r1: int, r2: bool) {
	var v0: array<bool> := new bool[7];
	forall i0 | 0 <= i0 < v0.Length {
		v0[i0] := if (|"icxytex"| != 0x19a) then p0 else p0;
	}
	var v1 := 0xe;
	var v2: seq<int> := [v1];
	var v3: C0 := new C0(v2);
	var v4: map<C0, bool> := map[v3 := p1];
	var v5: map<int, map<C0, bool>> := map[v1 := v4];
	var v6: seq<map<int, map<C0, bool>>> := [v5];
	var v7 := DC4(v6[safeIndex(v1, |v6|)], v1);
	match (fm2(v7.cf5, globalState)) {
		case DC1() =>
			r2 := false;
			var v8: seq<bool> := [p0, p0];
			var v9: map<bool, int> := map[v8[safeIndex(v1, |v8|)] := v1];
			var v10: array<int> := new int[12];
			var v11: multiset<int> := multiset{v1, fm4(v1, v1, p0, p0, globalState), |v9|, |[v10, v10, v10]|};
			var v12 := DC0(multiset{v1} == v11);
			var v13 := "avymisae";
			var v14: map<bool, string> := map[p1 := v13];
			var v15: map<bool, string> := map[!p1 := if (p0 in v14) then v14[p0] else v13];
			var v16: array<string> := new string[16] [v13, v13 + v13, v13, v13, v13 + v13, if (p1 in v15) then v15[p1] else v13[safeIndex(v1, |v13|) := 'w'], v13, seq(abs(0x1ea), i1  => ('m')), v13, if (p0) then v13 else v13, v13, v13 + v13, "rwsgvn", fm3(globalState), v13 + v13, v13];
			var v17 := 'v';
			var v18: array<char> := new char[4] [v17, v17, v17, v17];
			var v19: seq<C0> := [v3];
			var v20: array<C0> := new C0[10] [v19[safeIndex(v1, |v19|)], v3, v3, v3, v3, v3, v3, v3, v3, v3];
			var v21: map<int, array<C0>> := map[|v8| := v20];
			var v22: map<array<bool>, array<C0>> := map[v0 := if (v1 in v21) then v21[v1] else v20];
			var v23: map<array<char>, map<array<bool>, array<C0>>> := map[v18 := v22];
			var v24: seq<map<array<bool>, array<C0>>> := [v22];
			var v25: array<map<array<bool>, array<C0>>> := new map<array<bool>, array<C0>>[23] [if (v18 in v23) then v23[v18] else map[v0 := v20], v22, v22 + v22, v22, v22, v22, v22, v22, v22 + map[v0 := v20], DC6(map[v0 := v20]).cf7, v22, v22, v22, v22, v22, v22 + map[v0 := v20], v22, map[v0 := v20], v24[safeIndex(v1, |v24|)], v22 + v22, v22 + v22, v22, v22[v0 := v20]];
			v25[safeIndex(548, v25.Length)] := v22;
			r2, v12, v16, v25[safeIndex(548, v25.Length)] := p0, v12, v16, v22 + v22;
			var v26: array<array<int>> := new array<int>[25];
			v26[safeIndex(132, v26.Length)] := v10;
			var v27: map<bool, multiset<int>> := map[p1 := v11];
			var v28: multiset<bool> := multiset{p1, p0};
			var v29: seq<array<int>> := [v10];
			var v30: map<bool, array<int>> := map[(if (p0 in v27) then v27[p0] else multiset{v1, |v28|, 0x39, v1}) > v11 := v29[safeIndex(v1, |v29|)]];
			var v31 := DC8(v10);
			v26[safeIndex(132, v26.Length)] := if (p1 in v30) then v30[p1] else v31.cf10;
			r2 := !!!p0;
		case DC0(cf0) =>
			var v32 := DC5(seq(abs(-0x32), i2  => (v1)));
			v2 := v3.f0 + v32.cf6;
			var v33 := 'y';
			v33 := v33;
			r1 := v1 - -0x26;
			var v34: set<D1> := {v7};
			var v35: map<C0, set<D1>> := map[v3 := v34 * v34];
			var v36: seq<set<D1>> := [v34, v34, v34, v34, v34];
			v35 := v35[v3 := v36[safeIndex(v1, |v36|)]];
	}
	
	v0[safeIndex(314, v0.Length)] := p0;
	var v37 := "mexshnevu";
	var v38: multiset<bool> := multiset{fm1(p1, p0, |[v2]|, globalState), p0};
	var v39: seq<bool> := [p0, p0];
	var v40: multiset<C0> := multiset{v3};
	r2, r2, v0[safeIndex(314, v0.Length)], v37 := p0, (v38 <= v38) !in v39, v40 > v40, v37;
	var v41 := new C0(v2 + v3.f0);
	var v42: array<map<bool, D0>> := new map<bool, D0>[9](i3 => if (false) then map[v0[safeIndex(314, v0.Length)] := DC1()] else map[v0[safeIndex(314, v0.Length)] := DC1()]);
	var v43 := DC1();
	var v44: map<bool, D0> := map[p0 := v43];
	v42[safeIndex(646, v42.Length)] := v44;
	v42[safeIndex(646, v42.Length)] := match fm7(v0[safeIndex(314, v0.Length)], globalState) {
		case DC1() => v44
		case DC0(cf0) => v44
	};
	var v45: array<array<map<int, int>>> := new array<map<int, int>>[7];
	var v46: array<map<int, int>> := new map<int, int>[26](i4 => map[|{v1, |multiset{v1}|}| := v1]);
	v45[safeIndex(152, v45.Length)] := v46;
	v45[safeIndex(152, v45.Length)] := v46;
	r0 := [0x8b, v1, v1] + v2;
	r1 := v1;
	r2 := p0;
}
method Main() {
	var globalState := new GlobalState();
	var v0: array<bool> := new bool[7];
	var v1: set<array<bool>> := {v0};
	v1 := v1;
	var v2 := 98;
	var v3 := true;
	var v4: map<int, int> := map[|{v3, v3, v3}| := v2];
	var v6 := DC0(false);
	v2, v4, v3 := v2, (map v5 : int | (-0x2d5 <= v5) && (v5 < -518) :: (safeModuloInt(v5, v2)) := (--v2)) + v4, v3 <==> v6.cf0;
	forall i0 | 0 <= i0 < v0.Length {
		v0[i0] := v3;
	}
	var v7 := 'y';
	var v8: map<char, bool> := map[v7 := v3];
	for i1 := |v8| to v2 {
		var v9 := DC1();
		v9 := v9;
		var v10: set<bool> := {v3, v3};
		var v11, v12, v13 := m0(!(fm0(globalState) >= v10), true <== true, globalState);
		v3, v7, v3 := fm1(v13, v13, v12, globalState), v7, v13;
		v10 := v10;
	}
	var v14 := "braqfma";
	var v15 := DC1();
	var v16: array<D0> := new D0[12] [v15, v15, v15, v15, v15, fm2(v2, globalState), v15, DC1(), v15, v15, v15, v15];
	var v17: map<array<D0>, bool> := map[v16 := v3];
	v14, v3, v3, v17, v6 := (v14 + v14) + "dvtinys", v3, v3, v17 + v17, v6;
	var v18: multiset<string> := multiset{v14, v14, v14};
	v3, v2 := true, v2 + (|v18| + v2);
	var i2 := 0;
	while (v3 <==> !v3)
		decreases 100 - i2
	{
		if (i2 >= 100) {
			break;
		}
		
		i2 := i2 + 1;
		v2 := 175 + v2;
		var v19, v20, v21 := m0(false, v3, globalState);
		var v22: multiset<bool> := multiset{v21};
		v21 := v2 != safeModuloInt(v20, |v22|);
		v21 := v21;
	}
	v14 := fm3(globalState) + "ttjivi";
	if (!v3) {
		v3 := v3;
		var v23: map<int, bool> := map[v2 := v3];
		var v24: map<bool, int> := map[v3 := v2];
		var v25: seq<map<bool, int>> := [map[v3 := |v23|], v24];
		var v26: map<bool, int> := map[false := fm4(v2, |v25|, v3, v3, globalState) - -707];
		v24 := v24[v3 := v2 - v2];
		var v27: seq<bool> := [v3];
		v3 := v27[safeIndex(v2, |v27|)];
		v2 := 0x280;
		v2 := v2;
	} else {
		if (v3) {
			var v28: array<int> := new int[19];
			v28[safeIndex(611, v28.Length)] := v2;
			var v29: multiset<char> := multiset{v7, 'k', v7};
			var v30: map<bool, int> := map[v3 := -0x22];
			var v31: multiset<bool> := multiset{v3, v3, v3};
			var v32: multiset<int> := multiset{if (v7 in v29) then v29[v7] else fm4(v2, v2, false, v3, globalState), |multiset{v3}|, v2, |v30[fm1(v3, v3, v2, globalState) := v2]| + |v31|, v2};
			v28[safeIndex(611, v28.Length)] := if (v2 in v32) then v32[v2] else safeDivisionInt(850, -v2);
			var v33: seq<bool> := [false];
			var v34 := new C0([|multiset(v33)|]);
			v14 := v14;
			v7 := v14[safeIndex(v2, |v14|)];
			var v35: array<C0> := new C0[9];
			var v36: map<array<C0>, bool> := map[v35 := v3];
			var v37, v38, v39 := m0(v3, !(-|v36| < v34.fm5(v3, v3, v3, v28[safeIndex(611, v28.Length)], globalState)), globalState);
		} else {
			var v40: set<bool> := {!v3};
			var v41: map<D0, set<bool>> := map[v6 := v40];
			v3 := if (v3) then v3 else v41 == v41;
			v2 := v2 + -v2;
			v2 := v2;
			var v42: map<bool, bool> := map[v3 := v3];
			v42 := v42 + v42;
			var v43: seq<bool> := [true, v3, false, v3, v3];
			var v44: multiset<D0> := multiset{DC1()};
			var v45: map<bool, int> := map[v43[safeIndex(v2, |v43|)] := if (DC1() in v44) then v44[DC1()] else v2];
			v2 := safeModuloInt(|v45|, v2) + v2;
		}
		
		var v46: map<bool, bool> := map[v3 := v3];
		v46 := v46[v3 := v6.cf0];
		var v47: multiset<bool> := multiset{v3, v3, v3};
		var v48: C0 := new C0([|v14|, 0x2c2, |v47|]);
		var v49: map<bool, C0> := map[false := v48];
		var v50 := DC2(if (v3 in v49) then v49[v3] else v48);
		var v51: array<C0> := new C0[16] [v48, v48, v48, v48, if (true) then v48 else v48, v48, v48, v48, v48, v48, v48, v48, v50.cf1, v48, v48, v48];
		v51[safeIndex(539, v51.Length)] := v48;
		v51[safeIndex(539, v51.Length)] := v48;
		v3 := v3;
		v0[safeIndex(866, v0.Length)] := v3 <==> v3;
		v0[safeIndex(866, v0.Length)] := v3;
	}
	
	if (v6.cf0) {
		var v52: seq<bool> := [v3];
		var v53: multiset<bool> := multiset{v3, v3};
		v2 := if (multiset(v52) > v53) then v2 else -|v52| * v2;
		var v54: set<bool> := {true};
		v54 := if (v3) then {false, v3} - {!v3} else {v3, v3, v3};
		var v55: array<int> := new int[25](i3 => i3 - v2);
		v55[safeIndex(637, v55.Length)] := v2;
		v55[safeIndex(597, v55.Length)] := -0x10a;
		var v56: seq<int> := [v2];
		var v57: seq<int> := [|v56|, |v56|];
		var v58: C0 := new C0(v57);
		var v59: map<C0, bool> := map[v58 := true];
		var v60: map<int, map<C0, bool>> := map[|v14| := v59];
		v2, v3, v55[safeIndex(637, v55.Length)], v2, v55[safeIndex(597, v55.Length)] := -v2, v6.cf0, v2, 0xf6, -safeDivisionInt(|v14|, DC4(v60, v2).cf5) - v58.fm5(false, v3, v3, fm4(v2, v2, v3, v3, globalState), globalState);
		v3 := multiset(v52) >= v53;
		v3 := true;
	} else {
		v4 := v4[v2 := -v2];
		v2 := 0x3a7;
		v2 := -v2;
		v3 := !v3;
		var v61: map<multiset<bool>, int> := map[multiset{v3} := v2];
		var v62: multiset<bool> := multiset{v3};
		v3 := (if (v62 in v61) then v61[v62] else v2) <= v2;
	}
	
	var i4 := 0;
	while (v3)
		decreases 100 - i4
	{
		if (i4 >= 100) {
			break;
		}
		
		i4 := i4 + 1;
		v2 := v2;
		var v63: map<int, bool> := map[--v2 := v3];
		var v64 := new C0((fm6(v2, v3, if (0x342 in v63) then v63[0x342] else v3, globalState)).cf6);
		v14 := v14;
		v2 := if (v3) then v2 else v2;
	}
	v3 := v3;
	v3 := fm4(v2, -v2, false, true, globalState) <= v2;
	var v65: set<int> := {|v14|};
	var v66: seq<int> := [v2];
	var v67: C0 := new C0(v66);
	var v68: map<C0, bool> := map[v67 := v3];
	var v69: map<int, map<C0, bool>> := map[0xe8 := v68];
	var v70 := DC4(v69, |multiset{!v3, v3}|);
	var v71: map<set<int>, D1> := map[v65 := v70];
	var v72: seq<int> := [|v71|, --0x38f, |v14|];
	var v73 := new C0(v66);
	var v74, v75, v76 := m0(v3, v6.cf0, globalState);
	v75 := if (v3) then -0x226 else safeModuloInt(-v2, -v75);
	print v0[0], "\n";
	print v0[1], "\n";
	print v0[2], "\n";
	print v0[3], "\n";
	print v0[4], "\n";
	print v0[5], "\n";
	print v0[6], "\n";
	print |v1|, "\n";
	print v2, "\n";
	print v3, "\n";
	print v4 == map[59 := 98, 60 := 98, 61 := 98, 62 := 98, 63 := 98, 64 := 98, 65 := 98, 66 := 98, 67 := 98, 68 := 98, 69 := 98, 70 := 98, 71 := 98, 72 := 98, 73 := 98, 74 := 98, 75 := 98, 76 := 98, 77 := 98, 78 := 98, 79 := 98, 80 := 98, 81 := 98, 82 := 98, 83 := 98, 84 := 98, 85 := 98, 86 := 98, 87 := 98, 88 := 98, 89 := 98, 90 := 98, 91 := 98, 92 := 98, 93 := 98, 94 := 98, 95 := 98, 96 := 98, 97 := 98, 0 := 98, 1 := 98, 2 := 98, 3 := 98, 4 := 98, 5 := 98, 6 := 98, 7 := 98, 8 := 98, 9 := 98, 10 := 98, 11 := 98, 12 := 98, 13 := 98, 14 := 98, 15 := 98, 16 := 98, 17 := 98, 18 := 98, 19 := 98, 20 := 98, 21 := 98, 22 := 98, 23 := 98, 24 := 98, 25 := 98, 26 := 98, 27 := 98, 28 := 98, 29 := 98, 30 := 98, 31 := 98, 32 := 98, 33 := 98, 34 := 98, 35 := 98, 36 := 98, 37 := 98, 38 := 98, 39 := 98, 40 := 98, 41 := 98, 42 := 98, 43 := 98, 44 := 98, 45 := 98, 46 := 98, 47 := 98, 48 := 98, 49 := 98, 50 := 98, 51 := 98, 52 := 98, 53 := 98, 54 := 98, 55 := 98, 56 := 98, 57 := 98, 58 := 98, 199 := -199], "\n";
	print v6.cf0, "\n";
	print v7, "\n";
	print v8 == map['y' := false], "\n";
	print v14, "\n";
	print |v17|, "\n";
	print v18 == multiset{"braqfmabraqfmadvtinys", "braqfmabraqfmadvtinys", "braqfmabraqfmadvtinys"}, "\n";
	print i2, "\n";
	print i4, "\n";
	print v65 == {7}, "\n";
	print v66 == [-935], "\n";
	print v67.f0 == [-935], "\n";
	print |v68|, "\n";
	print |v69|, "\n";
	print |v70.cf4|, "\n";
	print v70.cf5, "\n";
	print |v71|, "\n";
	print v72 == [1, 911, 7], "\n";
	print v73.f0 == [-935], "\n";
	print v74 == [139, 14, 14, 14], "\n";
	print v75, "\n";
	print v76, "\n";
}