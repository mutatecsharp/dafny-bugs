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
datatype D0 = DC0 | DC1(cf0: int, cf1: seq<int>) | DC2(cf2: bool, cf3: array<string>, cf4: int)
datatype D1 = DC3(cf5: string)
class GlobalState {
	const f0 : string
	const f1 : bool
	var f2 : int
	const f3 : string
	constructor (f0 : string, f1 : bool, f2 : int, f3 : string) {
		this.f0 := f0;
		this.f1 := f1;
		this.f2 := f2;
		this.f3 := f3;
	}
	
}

class C0 {
	constructor () {
	}
	
	function fm2(p0: int, p1: D0, globalState: GlobalState): string {
		("ieh" + "ewuihk") + "h"
	}
}

function fm0(p0: bool, p1: int, globalState: GlobalState): int {
	safeModuloInt(0x341, |(map v0 : int | (0x235 <= v0) && (v0 < 0x282) :: (safeDivisionInt(v0, |"doerfd"|)) := (|(map v1 : seq<char> | v1 in multiset{['f']} :: (v1) := (0x18))|))|)
}
function fm1(globalState: GlobalState): char {
	'w'
}
function fm3(p0: seq<int>, p1: bool, p2: bool, p3: bool, globalState: GlobalState): multiset<int> {
	multiset{-329, --|multiset{true}|, 696, |{true, true, !true}|} - (multiset{|(seq(abs(982), i0  => ("omngle")))|} * multiset{|map[--597 := !false]|})
}
function fm4(p0: bool, globalState: GlobalState): bool {
	(if (true) then -0x5c else |map[!true := |multiset([|DC3(seq(abs(605), i0  => ('c'))).cf5|])|]|) <= safeDivisionInt(|"esnepq"|, 0x19e)
}
function fm5(p0: int, p1: bool, p2: bool, globalState: GlobalState): seq<bool> {
	[false, |"aahphab"| == ---0x3d2, !true, -0x2eb >= -0x3e0, true && !!true]
}
function fm6(p0: bool, p1: bool, globalState: GlobalState): string {
	match DC3("jkgbfxm") {
		case DC3(cf5) => cf5
	}
}
function fm7(p0: bool, p1: int, globalState: GlobalState): seq<int> {
	([|multiset{true, true, !false, true}|, |[true, true]|] + [-0x14e]) + ([|"cudaub"|] + [0x28a])
}
function fm8(globalState: GlobalState): multiset<bool> {
	(multiset{false} + multiset{true, false, true}) - multiset{true, true}
}
method m0(p0: int, p1: int, globalState: GlobalState) returns (r0: array<bool>, r1: int) {
	globalState.f2 := p0;
	var v0 := false;
	var v1: seq<bool> := [v0, v0, v0];
	v1 := v1;
	var v2: set<int> := {p1};
	var i0 := 0;
	while (!!((v2 <= v2) || true))
		decreases 100 - i0
	{
		if (i0 >= 100) {
			break;
		}
		
		i0 := i0 + 1;
		var v3 := 'r';
		var v4: map<bool, char> := map[v0 := v3];
		var v5: array<char> := new char[15] [v3, fm1(globalState), v3, v3, v3, v3, v3, v3, if (v0) then v3 else v3, if (v0 in v4) then v4[v0] else 'b', v3, v3, v3, v3, v3];
		v5[safeIndex(34, v5.Length)] := v3;
		v5[safeIndex(34, v5.Length)] := v3;
		r1 := fm0(v0, p0, globalState);
		var v6 := new C0();
		var v7: multiset<bool> := multiset{true, v0, v0};
		v0 := |(v7 + v7)| != p0;
	}
	var v8 := "bhpgj";
	for i1 := |v8| to safeDivisionInt(0x17c, 0x3c) {
		var v9: array<int> := new int[1] [i1];
		var v10: seq<array<int>> := [v9, v9, v9];
		globalState.f2 := |v10|;
		v8 := v8 + v8;
		var v11 := 'n';
		v8 := v8 + (fm6(!false, v0, globalState))[safeIndex(p1, |fm6(!false, v0, globalState)|) := v11];
		if (v0) {
			var v12: map<int, string> := map[830 * i1 := v8];
			v12 := v12[p1 := v8];
			globalState.f2 := p1;
			var v13: C0 := new C0();
			var v14: array<C0> := new C0[24] [v13, v13, v13, v13, v13, v13, v13, v13, v13, v13, if (v0) then v13 else v13, v13, v13, v13, v13, v13, if (v0) then v13 else v13, v13, v13, v13, v13, v13, if (v0) then v13 else v13, v13];
			v14 := new C0[27];
			var v15: multiset<bool> := multiset{v0, v0};
			globalState.f2 := if (!(v11 !in v8) in v15) then v15[!(v11 !in v8)] else i1;
			var v16: seq<int> := [p1, safeDivisionInt(p0, p0)];
			v16 := fm7(v0, if (v0) then 0x239 else i1, globalState);
		} else {
			var v17: array<string> := new string[16](i2 => v8);
			var v18 := DC2(v0, v17, p0);
			var v19: map<int, bool> := map[p0 := v18.cf2];
			var v20: multiset<int> := multiset{i1, i1, p1};
			v0 := if (safeModuloInt(i1, |v8[safeIndex(if (|v8| in v20) then v20[|v8|] else p0, |v8|) := v11]|) in v19) then v19[safeModuloInt(i1, |v8[safeIndex(if (|v8| in v20) then v20[|v8|] else p0, |v8|) := v11]|)] else v0;
			var v21: seq<int> := [fm0(v0, p0, globalState), p1];
			var v22: set<char> := {v11, v11};
			var v23: seq<seq<int>> := [v21];
			var v24 := DC1(p0, v21[safeIndex(|v22|, |v21|) := |v23|]);
			globalState.f2 := v24.cf0;
			v0 := v0;
			var v25 := DC0();
			globalState.f2, v25 := i1, v25;
			var v26: multiset<bool> := multiset{v0};
			v0 := (fm8(globalState) * v26) !! (v26 + multiset{v0});
		}
		
	}
	var v27: array<string> := new string[9];
	var v28: seq<int> := [p0, p0];
	var v29 := DC2(v0, v27, DC1(p1, v28).cf0);
	var i3 := 0;
	while (!v29.cf2)
		decreases 100 - i3
	{
		if (i3 >= 100) {
			break;
		}
		
		i3 := i3 + 1;
		var v30: array<C0> := new C0[1];
		var v31: C0 := new C0();
		v30[safeIndex(252, v30.Length)] := v31;
		v30[safeIndex(252, v30.Length)] := v31;
		var v32 := DC1(p1, v28);
		globalState.f2 := safeModuloInt(|[v0, v0]|, p1 - v32.cf0);
		v28 := v28;
		var v33: map<seq<int>, string> := map[v28 := v8 + "odihkxdux"];
		r1 := |v33|;
	}
	var v35: array<set<char>> := new set<char>[29](i5 => {'b'} - (set v34 : char | v34 in multiset(v8) :: (v34)));
	forall i4 | 0 <= i4 < v35.Length {
		v35[i4] := {'b'} - {'f'};
	}
	var v36: array<bool> := new bool[19];
	r0 := v36;
	r1 := p1;
}
method Main() {
	var v0 := "rhvwcv";
	var globalState := new GlobalState(v0 + v0, true, 0x19b, v0);
	var v1 := 387;
	globalState.f2 := v1;
	var v2 := false;
	v2 := v2;
	if (v2) {
		var v3, v4 := m0(v1, v1 - v1, globalState);
		var v5: multiset<seq<int>> := multiset{[-0x8, v4]};
		var v6: seq<multiset<seq<int>>> := [v5, v5];
		v2 := v5 > v6[safeIndex(-v4, |v6|)];
		var v7: array<map<bool, int>> := new map<bool, int>[16](i0 => map[v2 := |v0|]);
		v3[safeIndex(98, v3.Length)] := v2;
		var v8: map<int, int> := map[v1 := safeModuloInt(|v0|, v4)];
		var v9: set<int> := {v1, v4, v1};
		globalState.f2, v7, v3[safeIndex(98, v3.Length)] := |v8|, v7, v9 !! (v9 * v9);
		var v10, v11 := m0(safeDivisionInt(v1, v4), v1, globalState);
		v0 := "nwstnawx";
	} else {
		var v12: array<int> := new int[4];
		var v13: map<int, array<int>> := map[v1 := v12];
		var v14: array<array<int>> := new array<int>[1] [if (v1 in v13) then v13[v1] else v12];
		v14[safeIndex(281, v14.Length)] := v12;
		v14[safeIndex(281, v14.Length)] := v12;
		globalState.f2 := v1;
		var v15, v16 := m0(v1, --0x358, globalState);
		v14[safeIndex(281, v14.Length)] := v12;
		var v17: map<array<bool>, bool> := map[v15 := v16 >= v16];
		v17 := v17 + map[v15 := false];
	}
	
	for i1 := |("suibphoto" + v0)| to v1 {
		var v18: array<bool> := new bool[12](i2 => v1 != i1);
		v18 := v18;
		v1 := i1;
		var v19: seq<bool> := [false, v2];
		var v20, v21 := m0(0x297 * fm0(v2, |v19|, globalState), i1, globalState);
		v2, globalState.f2 := v2, -0x151;
	}
	var v22: set<bool> := {v2, v2, true};
	var v23: array<string> := new string[3];
	var v24 := DC2(v22 != v22, v23, v1);
	match (v24) {
		case DC0() =>
			var v25: map<int, bool> := map[v1 := v2];
			var v26: seq<bool> := [v2];
			var v27: map<seq<bool>, bool> := map[v26 := true];
			var v28: array<int> := new int[12] [v1, fm0(v2, v1, globalState), -0x20e, -0x390, v1, v1, --0x37e, v1, v1 - |v25|, |v27|, if (true) then v1 else v1, v1];
			var v29: array<char> := new char[5](i3 => 'd');
			var v30: map<array<char>, int> := map[v29 := v1];
			var v31: map<int, int> := map[v1 := |v30|];
			var v32: map<bool, map<int, int>> := map[true := v31];
			v28[safeIndex(377, v28.Length)] := |(if (v2 in v32) then v32[v2] else map v33 : int | (0x3bb <= v33) && (v33 < 0x148) :: (v33 * v1) := (v1))|;
			v28[safeIndex(377, v28.Length)] := fm0(v2, 0x15c, globalState);
			var v34 := 'o';
			var v35: array<bool> := new bool[4];
			v35[safeIndex(709, v35.Length)] := v2;
			var v36: map<D0, char> := map[if (v2) then v24 else DC2(false, v23, v1) := v34];
			v2, v28[safeIndex(377, v28.Length)], v34, v35[safeIndex(709, v35.Length)] := v2, safeModuloInt(v28[safeIndex(377, v28.Length)], v28[safeIndex(377, v28.Length)]), if ((if (false) then DC2(v2, v23, v28[safeIndex(377, v28.Length)]) else v24) in v36) then v36[if (false) then DC2(v2, v23, v28[safeIndex(377, v28.Length)]) else v24] else fm1(globalState), (if (!v2) then v2 else v2) ==> v2;
			var v37 := new C0();
			var v38: seq<int> := [-296, v1];
			v0, globalState.f2 := v0, v38[safeIndex(safeDivisionInt(v28[safeIndex(377, v28.Length)], v1), |v38|)];
		case DC1(cf0, cf1) =>
			var v39: array<int> := new int[18];
			v39[safeIndex(335, v39.Length)] := v1;
			v39[safeIndex(335, v39.Length)] := cf0 - cf0;
			v2 := v2;
			v39 := v39;
			v2 := v2;
		case DC2(cf2, cf3, cf4) =>
			v0 := v0;
			cf2 := v24.cf2;
			globalState.f2 := v1 + (v1 - v1);
			var v41: C0 := new C0();
			var v42: map<map<int, bool>, C0> := map[map v40 : int | (0x39f <= v40) && (v40 < -180) :: (safeDivisionInt(v40, v1)) := (v2) := v41];
			var v43, v44 := m0(|v42|, -(if (cf2) then v1 else 0x1fa), globalState);
	}
	
	var v45: C0 := new C0();
	var v46: multiset<C0> := multiset{v45, v45};
	var v47: seq<int> := [v1, v1, v1, |v46|];
	var v48: multiset<int> := multiset{-0xfe, v47[safeIndex(|v0|, |v47|)], v1};
	for i4 := v1 to if (-0x30b in v48) then v48[-0x30b] else v1 {
		var v49: seq<bool> := [v2];
		var v50: map<seq<bool>, bool> := map[v49 := v2];
		var v51 := 's';
		v0 := (((seq(abs(218), i5  => ('s'))) + v0)[safeIndex(|v50| + i4, |((seq(abs(218), i5  => ('s'))) + v0)|) := v51])[safeIndex(-v1, |((seq(abs(218), i5  => ('s'))) + v0)[safeIndex(|v50| + i4, |((seq(abs(218), i5  => ('s'))) + v0)|) := v51]|) := 'l'];
		var v52: map<int, int> := map[v1 + i4 := 807];
		v52 := v52[|{i4, i4, v1}| + v1 := safeDivisionInt(i4, v1)];
		var v53: array<int> := new int[20];
		var v54: set<array<int>> := {v53, v53};
		v2 := v54 >= v54;
		v2 := v49[safeIndex(v1, |v49|)];
	}
	var v55: array<array<int>> := new array<int>[2];
	v55 := v55;
	var v56: seq<bool> := [v2, v2];
	if ((v1 * |v56|) in v47) {
		v1 := if (false) then v1 else v1 + |(seq(abs(561), i6  => ('l')))|;
		var v57: map<char, bool> := map[fm1(globalState) := v2];
		var v58 := 'a';
		v2 := !(if (v58 in v57) then v57[v58] else !(v0 == v0));
		globalState.f2 := v1;
		v24 := v24;
		var v59: map<int, int> := map[0x3e1 := v1];
		v2 := 0x306 <= safeDivisionInt(--|v0|, |v59|);
	} else {
		var v60 := new C0();
		v2 := |v47| != v1;
		var v61, v62 := m0(v1, v1, globalState);
		if (v2) {
			var v63: array<int> := new int[25] [fm0(!v2, v1, globalState), v1, v1, v1, v62 + v1, v62 - v1, v1, v1, 0xa + v1, |"fatyy"|, v62, v1, v62, v1, v1, v1, v62 * --fm0(v2, |v22|, globalState), |v56| * 0x7, -v62, -v62, safeModuloInt(v62, |"iasqrm"|), -v62, v1, v47[safeIndex(v1, |v47|)], v62];
			v63 := v63;
			v2 := !(v48 <= (fm3([v1], v2, v2, v2, globalState) - v48));
			v63[safeIndex(57, v63.Length)] := v1;
			v63[safeIndex(57, v63.Length)] := |{if (v2) then v2 else true, v48 >= v48[v62 := abs(v1)], v0 <= v0}|;
			var v64: map<bool, bool> := map[v2 := fm4(v2, globalState)];
			var v65: map<int, bool> := map[fm0(v2, v62, globalState) := v2];
			v2 := if ((|v65| >= v63[safeIndex(57, v63.Length)]) in v64) then v64[|v65| >= v63[safeIndex(57, v63.Length)]] else !v2;
			var v66 := new C0();
		} else {
			var v67: array<seq<C0>> := new seq<C0>[20];
			var v68: array<int> := new int[8](i7 => i7 - v1);
			v68[safeIndex(580, v68.Length)] := v1;
			var v69: map<int, int> := map[|v0| := |v22|];
			var v70: map<array<bool>, map<int, int>> := map[v61 := v69];
			var v71: map<array<int>, array<seq<C0>>> := map[v68 := v67];
			v2, v2, v67, v68[safeIndex(580, v68.Length)] := v70 != v70, v2, if (v68 in v71) then v71[v68] else v67, -v62;
			var v72: set<seq<bool>> := {v56[safeIndex(v1, |v56|) := v2] + v56, v56 + v56};
			v72 := {v56, fm5(v68[safeIndex(580, v68.Length)], v2, v2, globalState) + [v2], v56};
			var v73 := 'a';
			var v74: map<multiset<int>, char> := map[v48 := v73];
			v74 := v74[v48 := if (false) then v73 else v73];
			var v75: array<D0> := new D0[24];
			v75[safeIndex(483, v75.Length)] := v24;
			v75[safeIndex(483, v75.Length)] := v24.(cf3 := v23);
			var v76: array<char> := new char[24](i8 => 'd');
			v76[safeIndex(942, v76.Length)] := v73;
			v76[safeIndex(942, v76.Length)], v48, v62 := v73, v48[v68[safeIndex(580, v68.Length)] := abs(v68[safeIndex(580, v68.Length)])], v68[safeIndex(580, v68.Length)];
		}
		
		var v77: multiset<bool> := multiset{!v2};
		v77 := v77;
	}
	
	v45 := v45;
	var v78 := 'i';
	v2, v2, v78, v1 := v2, v2, v78, |(multiset([v1, v1]) * (if (v2) then v48 else v48))|;
	var v79 := DC0();
	var v80: array<int> := new int[1] [|map[v1 := v79]|];
	forall i9 | 0 <= i9 < v80.Length {
		v80[i9] := safeModuloInt(i9, v1);
	}
	var v81 := new C0();
	v2 := v2;
	var v82: map<bool, int> := map[v2 := v1];
	if (if (v2) then v82 != v82 else !v2) {
		var v83, v84 := m0(v1, v1, globalState);
		v2 := !((seq(abs(0x271), i10  => (v78))) != v0);
		var v85, v86 := m0(safeDivisionInt(|v48|, v1), v1, globalState);
		v2 := v78 !in (v0 + v0);
		var v87 := new C0();
	} else {
		v1 := v1;
		globalState.f2, v1, v82 := -v1, v1, v82;
		var v88: array<array<char>> := new array<char>[12];
		var v89: array<char> := new char[7](i11 => v78);
		v88[safeIndex(65, v88.Length)] := v89;
		v88[safeIndex(65, v88.Length)] := v89;
		var v90, v91 := m0(|v0|, v1 + v1, globalState);
		v0 := v0;
	}
	
	var v92: array<bool> := new bool[9];
	v92 := v92;
	var v93: seq<array<int>> := [v80];
	var v94: map<seq<int>, seq<array<int>>> := map[v47 := v93 + v93];
	v94 := v94[v47 := v93];
	print v0, "\n";
	print globalState.f0, "\n";
	print globalState.f1, "\n";
	print globalState.f2, "\n";
	print globalState.f3, "\n";
	print v1, "\n";
	print v2, "\n";
	print v22 == {false, true}, "\n";
	print v24.cf2, "\n";
	print v24.cf4, "\n";
	print |v46|, "\n";
	print v47 == [386, 386, 386, 2], "\n";
	print v48 == multiset{-254, 386, 386}, "\n";
	print v56 == [false, false], "\n";
	print v78, "\n";
	print v80[0], "\n";
	print v82 == map[false := 2], "\n";
	print |v93|, "\n";
	print |v94|, "\n";
}