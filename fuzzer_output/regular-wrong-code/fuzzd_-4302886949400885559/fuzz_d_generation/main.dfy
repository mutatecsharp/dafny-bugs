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
datatype D0 = DC0(cf0: bool)
datatype D1 = DC1(cf1: map<bool, bool>)
datatype D2 = DC2(cf2: int)
datatype D3 = DC3(cf3: string)
datatype D4 = DC4(cf4: map<bool, int>)
datatype D5 = DC6 | DC5(cf5: map<bool, seq<char>>) | DC7(cf6: D5)
class GlobalState {
	const f0 : bool
	const f1 : int
	const f2 : bool
	var f3 : int
	const f4 : string
	var f5 : int
	const f6 : int
	var f7 : bool
	var f8 : string
	const f9 : int
	const f10 : map<bool, int>
	const f11 : bool
	var f12 : bool
	const f13 : bool
	const f14 : string
	constructor (f0 : bool, f1 : int, f2 : bool, f3 : int, f4 : string, f5 : int, f6 : int, f7 : bool, f8 : string, f9 : int, f10 : map<bool, int>, f11 : bool, f12 : bool, f13 : bool, f14 : string) {
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
	}
	
}

class C0 {
	constructor () {
	}
	
	function fm4(p0: bool, globalState: GlobalState): bool {
		true
	}
}

function fm0(p0: int, p1: bool, p2: map<bool, bool>, globalState: GlobalState): char {
	'y'
}
function fm1(p0: bool, p1: map<bool, int>, globalState: GlobalState): int {
	|map[[|[!false, false]|, |multiset{!!!DC0(false).cf0}|, --0x339] := multiset{|"wirhrrwqa"|, --0x131, 0x34e, |[!true]|} * multiset{|multiset{!true}|}]|
}
function fm2(p0: int, p1: bool, globalState: GlobalState): bool {
	{[0x1a7], [|"kxsyhfc"|], [|map[false := false]|]} >= {[-0xb1], [|DC4(map[true := |[true]|]).cf4|, |[false, true, true]|, |map[0x120 := !false]|]}
}
function fm3(p0: bool, p1: int, globalState: GlobalState): seq<char> {
	['y', 'b', 'e']
}
function fm5(p0: bool, p1: int, p2: bool, p3: int, globalState: GlobalState): D0 {
	DC0(false)
}
function fm6(p0: bool, p1: bool, globalState: GlobalState): map<bool, int> {
	map[true := 0x320] + (map[true := 0xba] + map[!true := |(set v0 : int | v0 in (seq(abs(331), i0  => (206))) :: (safeModuloInt(v0, -|[false]|)))|])
}
function fm7(p0: map<bool, bool>, globalState: GlobalState): map<bool, seq<char>> {
	DC5(map[true := ['k', 'd']]).cf5
}
function fm8(p0: bool, p1: int, p2: string, p3: int, globalState: GlobalState): multiset<int> {
	multiset{|map[true := 0x3a6]|}
}
function fm9(p0: bool, globalState: GlobalState): seq<bool> {
	[if (true) then true else !!false, true, false, !(|multiset{|"d"|, |map[!true := true]|, 0x112, 0x179}| != -0x3a8), false <== false]
}
method m0(p0: bool, p1: bool, p2: bool, globalState: GlobalState) returns (r0: map<bool, bool>) {
	var v0 := 896;
	globalState.f12 := v0 >= v0;
	var v1: set<bool> := {p2};
	var v2: map<bool, int> := map[p1 := v0];
	var v3: map<int, bool> := map[fm1(p1, v2, globalState) := p2];
	if (v1 >= (v1 * {if (-v0 in v3) then v3[-v0] else p2})) {
		var v4: seq<bool> := [p0];
		v4 := ([false] + v4) + v4;
		var v5: array<bool> := new bool[12] [p2, p2, p1, p2, p2, !!p0, !p2, p1, p1, p0, !p0, p1];
		var v6: C0 := new C0();
		var v7: map<array<bool>, C0> := map[v5 := v6];
		v7 := v7[v5 := v6];
		v4 := v4 + v4;
		var v8 := DC1(map[p2 := p0]);
		match (v8) {
			case DC1(cf1) =>
				var v9 := new C0();
				var v10: map<C0, bool> := map[v6 := p2];
				var v11: multiset<map<bool, int>> := multiset{v2 + fm6(if (v6 in v10) then v10[v6] else p1, p0, globalState)};
				globalState.f5 := if (v2 in v11) then v11[v2] else v0;
				var v12: array<D3> := new D3[27](i0 => DC3("mm"));
				var v13 := "igqyhvkw";
				var v14 := DC3(v13);
				v12[safeIndex(607, v12.Length)] := v14;
				v12[safeIndex(607, v12.Length)] := v14.(cf3 := v13 + v13);
				var v15 := 'j';
				var v16 := DC2(v0);
				var v17: multiset<C0> := multiset{v9};
				var v18: array<int> := new int[10] [|v17|, v0, |v2|, |v13|, |[v0]|, |v4|, v0, v0, |v13|, |(v13[safeIndex(v0, |v13|) := v15])[safeIndex(v0, |v13[safeIndex(v0, |v13|) := v15]|) := v15]|];
				var v19: multiset<int> := multiset{v16.cf2, |map[!p0 := v18]|, v0 * fm1(p0, v2, globalState)};
				globalState.f7, v15, v19, globalState.f7 := fm1(p0, v2, globalState) > v0, v15, v19, p2;
		}
		
		if (!(|v4| > v0)) {
			v5[safeIndex(353, v5.Length)] := p1;
			v5[safeIndex(353, v5.Length)] := true;
			globalState.f7 := v0 >= v0;
			globalState.f12 := p0;
			var v20: array<int> := new int[24](i1 => safeDivisionInt(i1, v0));
			var v21: seq<array<int>> := [v20, v20];
			v21 := v21;
			globalState.f3 := v0;
		} else {
			globalState.f12 := v0 < (v0 - -v0);
			var v22 := "aov";
			var v23: map<C0, string> := map[v6 := v22];
			var v24 := DC2(-98);
			var v25: map<int, string> := map[v24.cf2 := v22];
			var v26: multiset<C0> := multiset{v6};
			var v27: array<int> := new int[21] [fm1(p2, v2, globalState), |(v23 + v23)|, v0, -v0, -fm1(p1, v2, globalState), v0, v24.cf2, v0, v0, |v25|, v0, |(v2 + v2)|, v0, v0, 0x255, v0, 353, -safeDivisionInt(v0, -|v1|), v0 - |v26|, v0, fm1(p0, v2, globalState)];
			v27[safeIndex(307, v27.Length)] := v0;
			v27[safeIndex(307, v27.Length)] := v0 * v0;
			var v28 := 'x';
			v28 := v28;
			var v29 := new C0();
			v5[safeIndex(821, v5.Length)] := p2;
			v5[safeIndex(821, v5.Length)] := p0;
		}
		
	} else {
		var v30 := new C0();
		var v31 := "corwclx";
		var v32: map<int, multiset<int>> := map[v0 := fm8(true, v0, v31, v0, globalState)];
		var v33: multiset<int> := multiset{v0};
		v32 := v32[v0 := v33];
		var v34: map<int, map<int, bool>> := map[0x128 := v3];
		v34 := v34[safeDivisionInt(v0, v0) := v3];
		globalState.f12 := p2;
		var v35: array<char> := new char[2](i2 => 'y');
		var v36 := 'b';
		v35[safeIndex(563, v35.Length)] := v36;
		v35[safeIndex(563, v35.Length)] := v36;
	}
	
	var i3 := 0;
	while (p0)
		decreases 100 - i3
	{
		if (i3 >= 100) {
			break;
		}
		
		i3 := i3 + 1;
		if (p1) {
			globalState.f3 := v0;
			var v37: array<D3> := new D3[5](i4 => DC3("woh"));
			var v38 := DC3(fm3(p0, v0, globalState));
			v37[safeIndex(903, v37.Length)] := v38;
			var v39 := "siujnxa";
			v37[safeIndex(903, v37.Length)] := DC3(v39);
			globalState.f7 := p1;
			var v40: array<bool> := new bool[13] [p1, p1, p2, p1, if (p0) then p2 else p2, p2, p1, p0, p0, fm2(-v0, p0, globalState) !in v1, p2, p0, p0 && p1];
			v40 := v40;
			var v41: C0 := new C0();
			var v42: map<bool, bool> := map[p1 := p2];
			var v43: multiset<bool> := multiset{false, true, if (true in v42) then v42[true] else p0};
			globalState.f12, globalState.f3, globalState.f12, globalState.f12, v41 := p0, v0, false, v43 > v43, v41;
		} else {
			var v44 := new C0();
			var v45 := "sb";
			globalState.f12 := v0 > |v45|;
			var v46: multiset<int> := multiset{v0};
			globalState.f7 := fm8(p0, v0, "eqibxy", fm1(p1, map[false := v0], globalState), globalState) > (v46 - v46);
			var v47: array<bool> := new bool[14](i5 => p1);
			var v48: map<bool, bool> := map[p0 := p2];
			var v49: seq<map<bool, bool>> := [v48, v48];
			var v50: map<array<bool>, map<bool, bool>> := map[v47 := v49[safeIndex(v0, |v49|)]];
			var v51: seq<bool> := [p2];
			var v52: map<int, map<bool, bool>> := map[|v51| := v48];
			globalState.f5, r0 := -v0, if (v47 in v50) then v50[v47] else if (0x300 in v52) then v52[0x300] else v48;
			var v53: array<string> := new string[10](i6 => v45);
			v53 := v53;
		}
		
		globalState.f12 := !true;
		var v54: array<bool> := new bool[8](i7 => p2);
		v54 := new bool[23](i8 => if (p2) then if (|[p0]| in v3) then v3[|[p0]|] else p1 else p1);
		var v55: map<array<bool>, bool> := map[v54 := p0];
		globalState.f5 := v0 * |(v55 + v55)|;
	}
	var v56 := 'w';
	var v57 := "rwuaf";
	if (!!(if (p0) then v56 !in v57 else p2)) {
		var v58: array<C0> := new C0[26];
		v58 := v58;
		var v59: set<int> := {|multiset([v0])|};
		var v60: map<bool, set<int>> := map[p0 := v59];
		v60 := v60[p2 := v59];
		var v61 := new C0();
		globalState.f12 := p0;
		var v62: seq<int> := [v0, v0, 0xb3, -v0];
		var v63: map<seq<int>, bool> := map[v62 := !(!p0 ==> p2)];
		v63 := v63[v62 + v62 := p2];
	} else {
		v57 := v57;
		var v64: multiset<bool> := multiset{p1};
		var v65: set<int> := {v0};
		var v66: multiset<int> := multiset{|v57|};
		var v67: multiset<multiset<int>> := multiset{v66, v66};
		var v68: map<multiset<multiset<int>>, map<bool, int>> := map[v67 := v2];
		var v69: seq<int> := [v0];
		var v70: array<int> := new int[19] [v0, v0, v0, v0, v0, v0, v0, v0, |v64|, |v57|, v0, v0, -|v69|, v0, v0, v0, v0, v0, v0];
		var v71: seq<array<int>> := [v70, v70];
		var v72: seq<array<int>> := [v71[safeIndex(v0, |v71|)]];
		var v73 := DC0(false);
		var v74 := DC0(v73.cf0);
		var v75: C0 := new C0();
		var v76: map<C0, string> := map[v75 := v57];
		var v77: array<int> := new int[25] [|[p2]|, |v64|, v0, v0, -0x2c9, fm1(p0, v2, globalState), if (p2) then v0 else 0x90, |(v65 * v65)|, v0, v0, fm1(fm2(661, p0, globalState), if (v67 in v68) then v68[v67] else v2, globalState), v0, |v72|, v0, if (v73.cf0) then v0 else v0, v0, v0, |v69|, |(v69 + v69)|, v0, |v76|, v0, v0, |v1|, v0];
		v70, v56 := v77, v56;
		var v78: map<bool, bool> := map[p1 := p1];
		var v79: map<set<int>, char> := map[{v0, v0, v0} := fm0(if (p2 in v64) then v64[p2] else v0, p1, v78, globalState)];
		v79 := v79[{v0, v0} * {v0, v0} := v57[safeIndex(v0, |v57|)]];
		v64 := v64;
		var v80 := new C0();
	}
	
	var v81: map<bool, bool> := map[true := p0];
	var v82 := DC1(v81);
	globalState.f7, globalState.f12 := match v82 {
		case DC1(cf1) => p2
	}, p0;
	for i9 := |v1| to |(v57 + v57)| {
		globalState.f3 := v0;
		var v83: seq<bool> := [p2];
		var v84 := DC0(p1);
		var v85: map<bool, seq<bool>> := map[p1 := v83];
		var v86: array<seq<bool>> := new seq<bool>[25] [v83 + [p0, p2], v83, v83 + v83, v83, v83 + v83, v83, [p1], v83, v83 + [p0, fm2(i9, p0, globalState)], v83, v83, v83, v83, v83[safeIndex(i9, |v83|) := p0], v83, ((v83 + v83)[safeIndex(v0, |(v83 + v83)|) := !p0])[safeIndex(v0, |(v83 + v83)[safeIndex(v0, |(v83 + v83)|) := !p0]|) := p1], fm9(v84.cf0, globalState), v83, v83, [p0, p1, !!p0, p2], v83, [p1], (v83 + v83)[safeIndex(v0, |(v83 + v83)|) := p2], v83 + v83, if ((if (i9 in v3) then v3[i9] else p1) in v85) then v85[if (i9 in v3) then v3[i9] else p1] else [p0]];
		v86[safeIndex(295, v86.Length)] := v83;
		v86[safeIndex(295, v86.Length)], globalState.f12, globalState.f7 := v83 + v83, true, p1;
		globalState.f3 := i9;
		globalState.f5 := i9;
	}
	r0 := v81;
}
method Main() {
	var v0 := "olqswci";
	var v1 := true;
	var v2 := -891;
	var globalState := new GlobalState(false, 0x73, true, -0x37f, v0, 403, 751, false, v0, 0x3be, (map[v1 := v2])[v1 := v2], true, false, false, v0);
	if (v1) {
		var v3: set<bool> := {DC0(v1).cf0, v1, v1, v1, v1};
		globalState.f5 := if (true) then v2 - |v3| else v2;
		var v4: multiset<int> := multiset{-0x327};
		v1 := v4[v2 := abs(837)] > v4;
		var v5: array<string> := new string[19];
		v5[safeIndex(379, v5.Length)] := v0;
		var v6 := DC0(v1);
		var v7: seq<bool> := [v1, v1];
		globalState.f12, globalState.f3, v5[safeIndex(379, v5.Length)], v6, globalState.f5 := v7 < v7, v2, v0 + v0, v6, safeModuloInt(v2 + v2, safeDivisionInt(v2, v2));
		var v8: array<bool> := new bool[7](i0 => multiset{v1} >= multiset{v1});
		v8[safeIndex(625, v8.Length)] := v1;
		v8[safeIndex(625, v8.Length)] := v1;
		var v9: array<int> := new int[12];
		v9[safeIndex(699, v9.Length)] := v2;
		v9[safeIndex(699, v9.Length)], globalState.f3 := v2, 0x29e;
	} else {
		var v10 := m0(true, v1, v1, globalState);
		var v11: array<D0> := new D0[6](i1 => DC0(v1));
		var v12 := DC0(DC0(v1).cf0);
		v11[safeIndex(38, v11.Length)] := v12;
		v11[safeIndex(38, v11.Length)] := if (v1) then v12 else v12;
		var v13 := 'e';
		var v14 := DC1(v10);
		v13 := fm0(v2, v1, v14.cf1, globalState);
		if (true) {
			var v15: seq<bool> := [true];
			var v16: set<bool> := {v1};
			var v17: set<char> := {fm0(|v16|, v1, v10, globalState), v13};
			var v18: map<char, int> := map[v0[safeIndex(0x354, |v0|)] := v2];
			var v20: map<bool, int> := map[!v1 := v2];
			var v21: map<map<map<bool, bool>, int>, int> := map[map[map[v1 := v1] := v2] := v2];
			var v22: map<map<bool, bool>, int> := map[v10 := |v15|];
			var v23: seq<int> := [v2];
			var v24: array<int> := new int[26] [|(v15 + v15)|, v2, safeModuloInt(v2, v2), v2, 737, |(if (v1) then v17 else set v19 : char | v19 in v18 :: (v19))|, v2, v2, v2, v2, fm1(!fm2(v2, v1, globalState), v20, globalState), v2, v2, v2, v2 - v2, v2, fm1(v1, map[v1 := v2], globalState), if (v22 in v21) then v21[v22] else |(seq(abs(847), i2  => ('i')))|, -v2, -0x353, v2, v2, v2 * v2, v2, if (v1) then v23[safeIndex(fm1(v1, v20, globalState), |v23|)] else -v2, v2 + v2];
			v24[safeIndex(86, v24.Length)] := v2;
			v24[safeIndex(86, v24.Length)] := |(seq(abs(0x73), i3  => (v13)))|;
			var v25: array<bool> := new bool[23];
			var v26: map<int, bool> := map[v24[safeIndex(86, v24.Length)] := !true];
			v25[safeIndex(400, v25.Length)] := if (v2 in v26) then v26[v2] else v1;
			v25[safeIndex(400, v25.Length)] := v1;
			v0, v13, globalState.f12, globalState.f5, v1 := if (v25[safeIndex(400, v25.Length)]) then fm3(v1, v24[safeIndex(86, v24.Length)], globalState) else v0[safeIndex(v2, |v0|) := v13] + v0, v0[safeIndex(v2, |v0|)], v25[safeIndex(400, v25.Length)], v2, v1;
			var v27: map<bool, string> := map[v25[safeIndex(400, v25.Length)] <==> fm2(|v0|, v1, globalState) := v0];
			v27 := v27;
			var v28: seq<seq<int>> := [[v24[safeIndex(86, v24.Length)]], v23, [v24[safeIndex(86, v24.Length)]], v23[safeIndex(v24[safeIndex(86, v24.Length)], |v23|) := v2], v23];
			var v29: multiset<seq<int>> := multiset{v28[safeIndex(v24[safeIndex(86, v24.Length)], |v28|)]};
			var v30 := m0(v25[safeIndex(400, v25.Length)], v25[safeIndex(400, v25.Length)], !(v29 == v29), globalState);
		} else {
			var v31 := new C0();
			v14 := v14;
			var v32: array<int> := new int[13](i4 => i4 * -0x39d);
			var v33: seq<array<int>> := [v32, v32, v32];
			var v34: map<map<bool, bool>, array<int>> := map[v10 := if (v1) then v33[safeIndex(|v0[safeIndex(|map[v1 := v2]|, |v0|) := v13]|, |v33|)] else v32];
			v34 := (v34 + v34) + v34;
			var v35: multiset<bool> := multiset{v1};
			var v36: map<int, int> := map[v2 := safeModuloInt(v2, if (false in v35) then v35[false] else v2)];
			var v37: map<bool, int> := map[v1 := v2];
			v36 := v36[fm1(v1, v37, globalState) := v2];
			var v38: array<string> := new string[22](i5 => v0 + v0);
			v38[safeIndex(205, v38.Length)] := v0 + "hydejw";
			v38[safeIndex(205, v38.Length)] := seq(abs(0xff), i6  => ('a'));
		}
		
		globalState.f7 := v1;
	}
	
	var i7 := 0;
	while (true)
		decreases 100 - i7
	{
		if (i7 >= 100) {
			break;
		}
		
		i7 := i7 + 1;
		var v39: set<bool> := {v1, false, v1, v1};
		var v40: multiset<int> := multiset{|v0|};
		var v41: seq<int> := [v2];
		var v42: array<bool> := new bool[16] [v1, if (v1) then v1 else v1, false, v1, v1, v1, -|v39| > v2, v40 !! (multiset{v2, v2})[v2 := abs(v2)], !(v41 < v41), v1, v1, v1, v1, v1, v1 in map[v1 := v1], v1];
		var v43: seq<bool> := [v1];
		var v44: array<int> := new int[8];
		v42, v43, v44, v2 := v42, v43, v44, v2 - (v2 - v2);
		v42[safeIndex(232, v42.Length)] := v1;
		v42[safeIndex(232, v42.Length)] := v40 <= multiset{|v0|};
		globalState.f12 := v42[safeIndex(232, v42.Length)];
		var v45 := DC2(|map[v42[safeIndex(232, v42.Length)] := fm2(v2, v1, globalState)]|);
		var v46: set<int> := {v45.cf2, -v2, v2};
		var v47: map<int, int> := map[-v2 := v2];
		v46 := v46 - ({|v47|, v2} - v46);
	}
	var v48: array<bool> := new bool[11](i8 => !v1);
	var v49: map<array<bool>, bool> := map[v48 := v1];
	v49 := v49[v48 := v1];
	for i9 := v2 to v2 {
		var v50 := new C0();
		var v51 := DC0(v1);
		var v52: set<D0> := {v51.(cf0 := v1)};
		v52, globalState.f7 := v52, v1;
		var v53 := m0(false, v1, v1, globalState);
		var v54: multiset<int> := multiset{0x317};
		var v55: seq<multiset<int>> := [v54, multiset{v2, v2}, v54];
		var v56: seq<int> := [i9, 736];
		globalState.f5, globalState.f5, v1, v1, v54 := 0x35a - v2, i9, v1, v1, (v55[safeIndex(v2, |v55|)])[if (v1) then v2 else v56[safeIndex(v2, |v56|)] := abs(v2)];
	}
	var v57: map<bool, int> := map[v1 := v2];
	for i10 := v2 to -fm1(v1, v57, globalState) {
		if (v1) {
			var v58: seq<string> := [v0];
			var v59: C0 := new C0();
			var v60: set<C0> := {v59};
			var v61: seq<set<C0>> := [v60, v60];
			var v62 := DC3(v0);
			var v63 := DC0(v1);
			var v64 := 'b';
			var v65: array<string> := new string[20] ["apenmlek", fm3(v1, -i10, globalState), v0, v0, v0, "jxraijmpf", v58[safeIndex(|v61[safeIndex(v2, |v61|)]|, |v58|)], v62.cf3 + v0, "vqhbcmux", v0, v62.cf3[safeIndex(i10, |v62.cf3|) := 'o'], v0, v0, v0, v0, fm3(v59.fm4(v1, globalState), i10, globalState), v0, (fm3(!v63.cf0, -v2, globalState))[safeIndex(i10, |fm3(!v63.cf0, -v2, globalState)|) := v64], seq(abs(247), i11  => (v64)), v0 + v0];
			v65[safeIndex(113, v65.Length)] := "evqq";
			v65[safeIndex(113, v65.Length)] := v0;
			globalState.f7 := v1;
			var v66: map<int, C0> := map[v2 := v59];
			var v67: seq<C0> := [if (v1) then v59 else if (v2 in v66) then v66[v2] else v59, v59];
			v59 := v67[safeIndex(v2, |v67|)];
			var v68 := new C0();
			var v69: array<array<int>> := new array<int>[16];
			var v70: array<int> := new int[12];
			v69[safeIndex(722, v69.Length)] := v70;
			v69[safeIndex(722, v69.Length)] := v70;
		} else {
			var v71 := m0(v1, v1, v1, globalState);
			globalState.f7 := v1;
			globalState.f7 := v0 < v0;
			var v72: array<int> := new int[5];
			v72[safeIndex(28, v72.Length)] := 1;
			v72[safeIndex(28, v72.Length)] := i10 - safeDivisionInt(i10, v2);
			globalState.f5 := fm1(v0 == (seq(abs(0xaa), i12  => ('p'))), v57, globalState);
		}
		
		var v73: seq<int> := [i10, v2];
		var v74: seq<seq<int>> := [v73, v73];
		var v75 := m0(false, v74 == [v73, [i10, i10, v2]], v1, globalState);
		globalState.f12 := v1;
		var v76: array<map<bool, int>> := new map<bool, int>[4];
		var v77: array<int> := new int[17];
		v77[safeIndex(26, v77.Length)] := i10 * i10;
		var v78: multiset<bool> := multiset{v1};
		var v79: multiset<int> := multiset{if (v1 in v78) then v78[v1] else v2};
		v77[safeIndex(79, v77.Length)] := |v79|;
		var v80: C0 := new C0();
		var v81: map<C0, string> := map[v80 := "pqq"];
		globalState.f5, v76, v77[safeIndex(26, v77.Length)], v77[safeIndex(79, v77.Length)] := -|(if (v80 in v81) then v81[v80] else v0)|, v76, -|v0| + v2, |v0|;
	}
	if (true) {
		var v82 := new C0();
		v57 := v57[v1 := safeModuloInt(v2, v2)];
		var v83: multiset<array<bool>> := multiset{v48};
		var v84: seq<int> := [|"uknqeb"|, v2];
		var v85: map<seq<int>, multiset<array<bool>>> := map[v84 := v83[v48 := abs(v2)]];
		v83 := if (v84 in v85) then v85[v84] else v83 + v83;
		globalState.f7 := false || v1;
		var v86: map<int, int> := map[v2 := v2];
		var v87: multiset<bool> := multiset{v1};
		v86 := v86[467 := |(if (!v1) then v87 else v87)|];
	} else {
		var v88: multiset<bool> := multiset{v1};
		var v89: seq<int> := [v2];
		var v90: set<array<bool>> := {v48};
		var v91 := 'd';
		var v92: seq<set<char>> := [{v91}];
		var v93: map<seq<set<char>>, string> := map[v92 := seq(abs(-0x1db), i14  => ('l'))];
		var v94: seq<bool> := [v1];
		var v95: array<int> := new int[4] [v2, v2, |v94|, v2];
		var v96: map<array<int>, bool> := map[v95 := v1];
		var v97: array<int> := new int[26] [v2, v2, v2, v2, v2, -621, safeDivisionInt(|(seq(abs(-603), i13  => ('m')))|, |v88|), fm1(v1, map[false := 0x2e6], globalState), v2, -310, v2, |{v1}|, |v57|, v2, if (!v1) then |v89| else v89[safeIndex(-|"oe"|, |v89|)], fm1(v1, v57, globalState), v2, fm1(v1, v57, globalState), safeDivisionInt(v2, v2), v2, safeModuloInt(v2, |v90|), v2, |(if (v92 in v93) then v93[v92] else v0)|, v2, v2, |v96|];
		v95[safeIndex(598, v95.Length)] := v2;
		var v98 := DC2(v2);
		v95[safeIndex(598, v95.Length)] := v98.cf2;
		var v99: set<int> := {if (v1) then v95[safeIndex(598, v95.Length)] else v95[safeIndex(598, v95.Length)]};
		var v100 := DC0(v1);
		var v101: array<D0> := new D0[7] [DC0(fm2(v2, v1, globalState)), v100, v100, v100, DC0(v1), fm5(false, 0x225, v1, v2, globalState), DC0(v1)];
		v101[safeIndex(574, v101.Length)] := v100;
		var v102: C0 := new C0();
		var v103: seq<C0> := [v102];
		v99, globalState.f12, v101[safeIndex(574, v101.Length)], v1 := v99, v1 ==> ([v102] == v103), DC0(fm6(v1, v1, globalState) != v57), ([!v1] + v94) < v94;
		var v104 := m0(v95[safeIndex(598, v95.Length)] in v89, v1, v1, globalState);
		var v105: array<set<bool>> := new set<bool>[10];
		var v106: set<bool> := {v1};
		v105[safeIndex(53, v105.Length)] := v106;
		v88, v105[safeIndex(53, v105.Length)], globalState.f3 := multiset{true}, v106 - (if (v1) then v106 else v106), v2;
		var v107 := new C0();
	}
	
	var v108 := DC1(map[v1 := v1]);
	var v109: set<D1> := {v108, v108};
	var v110 := DC2(|v109|);
	var v111: array<int> := new int[5] [v110.cf2, fm1(v1, v57, globalState), v2, v2, v2];
	v111 := new int[3];
	var v112 := 'j';
	var v113 := DC3(v0);
	var v114: map<bool, seq<char>> := map[v112 in v0 := v113.cf3 + v0];
	var v115: map<bool, bool> := map[v1 := v1];
	v114 := fm7(v115, globalState);
	v1 := v1;
	var v116 := m0(v1, !v1, v1, globalState);
	v2 := match v113 {
		case DC3(cf3) => v2 + |map[v1 := [v1]]|
	};
	if (!v1) {
		globalState.f7 := v1;
		var v117: seq<array<int>> := [v111];
		v111 := if (v1 && v1) then v111 else v117[safeIndex(v2, |v117|)];
		var v118: map<bool, array<int>> := map[v2 != (if (v1 in v57) then v57[v1] else v2) := v111];
		v118 := v118[!false := v111];
		globalState.f3 := v2;
		globalState.f8 := if (if (v1 in v115) then v115[v1] else v1) then v0[safeIndex(v2, |v0|) := v112] + (seq(abs(0x110), i15  => (v112)))[safeIndex(v2, |(seq(abs(0x110), i15  => (v112)))|) := v112] else v0;
	} else {
		var v119: C0 := new C0();
		var v120: map<C0, bool> := map[v119 := !v1];
		v1, v119, globalState.f7 := fm2(|map[v1 := v119]|, false, globalState), v119, if (v1) then v1 else if (v119 in v120) then v120[v119] else v1;
		var v121 := DC0(v1);
		v2 := |((v116 + v115) + map[v1 := v121.cf0])|;
		var v122: multiset<int> := multiset{|[v119]|};
		v112 := fm0(|v122|, v1, v116 + v116[v1 := v1], globalState);
		globalState.f7 := v1;
		globalState.f7 := true;
	}
	
	var i16 := 0;
	while (v1)
		decreases 100 - i16
	{
		if (i16 >= 100) {
			break;
		}
		
		i16 := i16 + 1;
		v1 := v1;
		var v123: C0 := new C0();
		v123 := v123;
		var v124 := new C0();
		var v125 := m0(v1, v1, v1 || v1, globalState);
	}
	for i17 := v2 to v2 - v2 {
		match (v113) {
			case DC3(cf3) =>
				globalState.f5 := i17;
				v57 := v57[!(v1 ==> v1) := i17];
				v48[safeIndex(769, v48.Length)] := fm2(v2, v1, globalState);
				v48[safeIndex(769, v48.Length)] := v1;
				globalState.f7 := fm2(v2 + i17, !v1, globalState);
		}
		
		var v126: set<int> := {-i17};
		globalState.f12 := !({v2} > v126);
		var v127: C0 := new C0();
		var v128: map<string, bool> := map[seq(abs(0x1b1), i18  => ('j')) := v1];
		var v129: seq<bool> := [v1];
		v0, globalState.f5, v127, v1 := (if (if (v0 in v128) then v128[v0] else false) then v0 else v0) + v0, if (!true in v57) then v57[!true] else |v129|, v127, v0 <= (v0 + v0);
		var v130: array<char> := new char[27](i20 => v112);
		var v131: map<array<char>, bool> := map[v130 := v1];
		v48[safeIndex(596, v48.Length)] := |(seq(abs(-0x127), i19  => (i17)))| > |v131|;
		v48[safeIndex(596, v48.Length)] := !(if (v1) then v2 <= |"dvpiwucvt"| else false);
	}
	v1 := v1 <== true;
	globalState.f12 := !((v0 + v0) <= v0);
	print v0, "\n";
	print v1, "\n";
	print v2, "\n";
	print globalState.f0, "\n";
	print globalState.f1, "\n";
	print globalState.f2, "\n";
	print globalState.f3, "\n";
	print globalState.f4, "\n";
	print globalState.f5, "\n";
	print globalState.f6, "\n";
	print globalState.f7, "\n";
	print globalState.f8, "\n";
	print globalState.f9, "\n";
	print globalState.f10 == map[true := -891], "\n";
	print globalState.f11, "\n";
	print globalState.f12, "\n";
	print globalState.f13, "\n";
	print globalState.f14, "\n";
	print i7, "\n";
	print v48[0], "\n";
	print v48[1], "\n";
	print v48[2], "\n";
	print v48[3], "\n";
	print v48[4], "\n";
	print v48[5], "\n";
	print v48[6], "\n";
	print v48[7], "\n";
	print v48[8], "\n";
	print v48[9], "\n";
	print v48[10], "\n";
	print |v49|, "\n";
	print v57 == map[true := 0], "\n";
	print v108.cf1 == map[true := true], "\n";
	print v109 == {DC1(map[true := true])}, "\n";
	print v110.cf2, "\n";
	print v112, "\n";
	print v113.cf3, "\n";
	print v114 == map[true := ['k', 'd']], "\n";
	print v115 == map[true := true], "\n";
	print v116 == map[true := true], "\n";
	print i16, "\n";
}