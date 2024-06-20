datatype D0 = DC1(cf1: array<bool>) | DC2(cf2: bool, cf3: map<int, bool>, cf4: map<bool, char>, cf5: bool) | DC0(cf0: bool) | DC3(cf6: D0)
datatype D1 = DC5(cf8: bool) | DC4(cf7: array<int>) | DC6(cf9: D1)
datatype D2 = DC8(cf11: bool, cf12: int, cf13: bool) | DC9(cf14: int) | DC7(cf10: int)
datatype D3 = DC10(cf15: string)
datatype D4 = DC11(cf16: set<bool>)
datatype D5 = DC13(cf18: bool, cf19: bool) | DC14(cf20: bool) | DC12(cf17: seq<multiset<bool>>)
datatype D6 = DC16(cf22: array<bool>, cf23: int) | DC17(cf24: bool) | DC18(cf25: int) | DC15(cf21: seq<int>) | DC19(cf26: D6)
datatype D7 = DC21 | DC20(cf27: map<bool, array<int>>)
datatype D8 = DC23(cf29: int) | DC24(cf30: set<bool>, cf31: seq<bool>) | DC22(cf28: multiset<string>)
datatype D9 = DC26 | DC27(cf33: C1, cf34: int, cf35: int, cf36: int, cf37: map<map<int, bool>, char>) | DC25(cf32: char) | DC28(cf38: D9)
datatype D10 = DC30(cf40: int) | DC29(cf39: T2) | DC31(cf41: D10)
datatype D11 = DC33(cf43: map<bool, bool>, cf44: int, cf45: bool, cf46: bool, cf47: int) | DC34(cf48: D0, cf49: int, cf50: map<C4, C1>) | DC32(cf42: map<bool, array<bool>>) | DC35(cf51: D11)
datatype D12 = DC37 | DC38(cf53: D8, cf54: int, cf55: int) | DC39(cf56: bool) | DC36(cf52: set<D3>)
class GlobalState {
	const f0 : multiset<int>
	const f1 : bool
	const f2 : bool
	const f3 : bool
	const f4 : int
	var f5 : int
	const f6 : bool
	const f7 : bool
	const f8 : int
	const f9 : array<char>
	constructor (f0 : multiset<int>, f1 : bool, f2 : bool, f3 : bool, f4 : int, f5 : int, f6 : bool, f7 : bool, f8 : int, f9 : array<char>) {
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
	}
	
}

function fm0(p0: int, p1: int, globalState: GlobalState): int {
	0x358 % -0x1bc
}
function fm1(p0: int, p1: bool, p2: int, p3: bool, globalState: GlobalState): bool {
	false && (-0x208 > |DC10(seq(-618, i0  => ('c'))).cf15|)
}
function fm3(p0: bool, p1: bool, p2: int, globalState: GlobalState): map<int, bool> {
	map v0 : int | (-0x9e <= v0) && (v0 < 0xe4) :: (v0 % -841) := (true ==> false)
}
function fm9(p0: bool, globalState: GlobalState): seq<bool> {
	(if (false) then [true] else [true]) + ([false] + [true, true])
}
function fm10(p0: int, globalState: GlobalState): map<int, bool> {
	(if (true) then map[0x3c8 := true] else map[-0x73 := true]) + map[-0x3d5 := !false]
}
function fm11(p0: int, p1: int, p2: bool, globalState: GlobalState): string {
	"t"
}
function fm12(p0: int, p1: int, p2: int, p3: int, globalState: GlobalState): map<int, int> {
	map[-DC7(0x284).cf10 := 76]
}
function fm13(p0: bool, p1: bool, globalState: GlobalState): set<int> {
	(set v0 : int | (0x1fb <= v0) && (v0 < 0x17) :: (v0 % |[true, false]|)) + ({|"gj"|, |map[true := |[false, false]|]|} * {-|map[true := 'q']|, |map[seq(383, i0  => ('h')) := true]|, |"dyck"|, |{"r"}|, |map[true := |[|multiset{seq(-0x42, i1  => ('g')), seq(-0x72, i2  => ('e'))}|]|]|})
}
function fm14(p0: int, p1: D2, globalState: GlobalState): set<bool> {
	{true, {|"sqbhaywk"|, 0xa0} > {|(seq(0x23c, i0  => ('x')))|, |(set v0 : int | (0x217 <= v0) && (v0 < 0x300) :: (v0 % 0x1c7))|}}
}
function fm15(p0: map<bool, bool>, p1: int, globalState: GlobalState): char {
	'e'
}
function fm16(p0: bool, p1: int, p2: D6, p3: int, globalState: GlobalState): multiset<seq<bool>> {
	multiset{[false, true], [true, true, false, false], [false, true], [false]} - multiset(if (!true) then [DC24({!true, true}, [!false]).cf31, [true], [true, true]] else [[true], [!false], [true], [false, true, true, false], [true, true]])
}
function fm17(p0: bool, p1: bool, globalState: GlobalState): D2 {
	match DC18(|[|"hqs"|]|) {
		case DC16(cf22, cf23) => DC8(false, |DC36({DC10("rtpervwnl")}).cf52|, false)
		case DC17(cf24) => DC8(cf24, 877, cf24)
		case DC18(cf25) => DC8(false, cf25, !!!true)
		case DC15(cf21) => DC8(false, |(map v0 : set<int> | v0 in {{-|(seq(0x287, i0  => ('b')))|}, {-146, 401, -0x125, |cf21|, |multiset{|multiset{-|multiset{0x1d7}|, 0x182}|, |map[false := false]|}|}} :: (v0) := (|(seq(0x2c4, i1  => ('f')))|))|, false)
		case DC19(cf26) => DC8(!true, -990, false)
	}
}
function fm18(p0: int, p1: int, p2: bool, globalState: GlobalState): map<bool, bool> {
	map[!!false || false := false]
}
function fm19(p0: int, globalState: GlobalState): multiset<bool> {
	multiset{false} + multiset{true}
}
function fm20(p0: int, p1: int, p2: int, p3: D5, globalState: GlobalState): set<char> {
	{'v', 'u', if (true) then DC25('k').cf32 else 'w'}
}
function fm21(p0: int, p1: int, p2: char, p3: int, globalState: GlobalState): multiset<int> {
	match DC23(663) {
		case DC23(cf29) => multiset{cf29, cf29, |[true, false]|} + multiset{0xf0}
		case DC24(cf30, cf31) => multiset{267}
		case DC22(cf28) => multiset{|map['i' := 0x3e0]|} - multiset([|"jgetcs"|])
	}
}
function fm22(p0: bool, globalState: GlobalState): D8 {
	DC24({false, true}, [true, true] + [false])
}
function fm23(p0: bool, p1: int, p2: int, globalState: GlobalState): seq<int> {
	(if (!false) then seq(0x1cb, i0  => (|(seq(0xdc, i1  => ('y')))|)) else [|map[map[-19 := false] := !false]|, 0x31]) + (seq(-0x3b9, i2  => (-0x121)))
}
function fm24(p0: bool, p1: int, globalState: GlobalState): seq<D5> {
	([DC14(true), DC14(true)] + [DC14(true), DC14(false)]) + [DC14(true), DC14(true)]
}
function fm25(p0: bool, p1: bool, p2: int, globalState: GlobalState): D5 {
	DC14(!false)
}
function fm26(p0: int, p1: char, p2: bool, p3: char, globalState: GlobalState): seq<D5> {
	seq(0x2e3, i0  => (DC13(true, false)))
}
function fm27(p0: map<int, bool>, p1: bool, p2: string, globalState: GlobalState): D9 {
	DC26()
}
method m0(p0: int, p1: int, p2: int, p3: bool, globalState: GlobalState) returns (r0: int, r1: map<bool, int>, r2: string) {
	if (p3) {
		var v0 := "gnxtuopx";
		var v1: multiset<bool> := multiset{p3};
		var v2: array<D2> := new D2[22](i2 => DC7(|multiset{v1}|));
		var v3: T2 := new C3(v2, 'v', 419);
		var v4: map<seq<int>, T2> := map[seq(668, i1  => (-0xc2)) := v3];
		var v5: map<bool, string> := map[p3 := seq(-89, i3  => ('y'))];
		var v6: array<string> := new string[13] [v0, fm11(if (p3 in v1) then v1[p3] else |v0|, p2, p3, globalState), seq(0x78, i0  => ('b')), v0, v0 + fm11(|v4|, p0, p3, globalState), "exf", "ivgkwls", if (p3 in v5) then v5[p3] else v0, v0, v0, v0, v0, v0];
		v6[44] := seq(0x1d9, i4  => ('u'));
		var v7 := 'c';
		globalState.f5, v6[44] := p1 * p0, v0[v3.f12 := v7];
		var v8: seq<int> := [v3.f12];
		var v9: map<int, int> := map[|v6[44]| := p0];
		var v10: array<seq<int>> := new seq<int>[13] [v8, v8 + v8, v8, v8, v8, v8 + [p1], v8, v8, fm23(p3, 16, -0x1d9, globalState), v8, [|v9|], v8 + v8, [p2]];
		v10[191] := v8;
		var v11: map<bool, seq<int>> := map[p3 := v8];
		v10[191] := if (true in v11) then v11[true] else v8;
		var v12 := DC0(p3);
		match (v12) {
			case DC1(cf1) =>
				var v13: map<int, bool> := map[v3.f12 := p3];
				v13 := v13[p2 := p3];
				v3.f12 := p1;
				var v14 := DC14(false);
				var v15: seq<D5> := [DC14(p3), v14];
				var v16: seq<seq<D5>> := [v15, v15];
				var v17: seq<array<D2>> := [v2, v2];
				var v18: T1 := new C3(v17[|multiset{-p2}|], 's', |multiset{p3, p3}|);
				var v19: multiset<T1> := multiset{v18, v18};
				var v20: map<multiset<T1>, char> := map[v19[v18 := v3.f12] := v7];
				var v21: multiset<int> := multiset{p2, v3.f12};
				var v22: map<seq<D5>, T2> := map[(v16[|v20|] + fm24(p3, v3.f12, globalState))[|v21| := DC14(p3)] := v3];
				v22 := v22[v15 + v15[v18.f12 := v14] := v3];
				var v23 := false;
				v23 := p3;
			case DC2(cf2, cf3, cf4, cf5) =>
				v8 := [p2, -p0, 0x3cd] + [|v0|];
				var v24: map<char, bool> := map[v7 := p3];
				var v25: array<int> := new int[18] [v3.f12 - p2, |v6[44]|, -(v3.f12 / p1), p0, v3.f12, p0 * |"die"|, -(fm0(493, |v24|, globalState) * 685), p1, |v6[44]|, -v3.f12, p2, p1, p1, p1, |v0|, |(seq(0x21c, i5  => (|v8|)))| - p0, v3.f12, v3.f12];
				v25[571] := v3.f12 * v3.f12;
				var v26: set<string> := {"scckxsky", v6[44]};
				v25[571] := if (cf5) then p2 else |v26|;
				v7 := v7;
				cf5 := !(p2 < p0);
			case DC0(cf0) =>
				var v28: set<int> := {p1};
				v3 := new C0(map v27 : int | v27 in fm23(cf0, v3.f12, p2, globalState) :: (v27 * v3.f12) := (cf0), |v28|);
				var v29: array<bool> := new bool[22];
				v29[609] := cf0;
				var v30 := DC14(false);
				var v31: seq<string> := [v6[44]];
				v29[609], v30, cf0 := ["y", v0] < v31, fm25(cf0, cf0, fm0(p2, -p2, globalState) % v3.f12, globalState), cf0;
				v29[609] := cf0;
				globalState.f5, globalState.f5, globalState.f5, v7 := v3.f12, if (v29[609]) then |v6[44]| else v10[191][p2], -v3.f12, v7;
			case DC3(cf6) =>
				var v32: array<C3> := new C3[14];
				var v33: C3 := new C3(v2, v7, p1);
				v32[0] := v33;
				v32[0] := new C3(v2, v7, -p1);
				var v34: array<int> := new int[25];
				var v35: multiset<array<int>> := multiset{v34, v34};
				v3.f12 := if (v34 in v35) then v35[v34] else 0x1b0;
				var v36: map<int, bool> := map[v3.f12 * v3.f12 := p3];
				v36 := v36[v3.f12 * |v1| := p3];
				var v37: set<int> := {v3.f12};
				var v38 := new C4(|(v37 + v37)|, if (fm1(-0x94, p3, p0, p3, globalState)) then v3.f12 else v3.f12);
		}
		
		var v39: seq<bool> := [p3];
		if ([p3] != (v39 + v39)) {
			v3.f12 := if (p0 in v9) then v9[p0] else p0;
			var v40: map<bool, int> := map[p3 := p1];
			var v41: array<bool> := new bool[15] [fm1(p0, false, v8[|v1|], p3, globalState), p3, p3, p3, p3, p3, p3 || !false, false, p3, !(fm23(p3, fm0(|v40|, p2, globalState), p1, globalState) <= v10[191]), p3, p3, p3, p3, v39[489]];
			v41[757] := p3;
			v41[757] := p3;
			v41[757] := v41[757];
			v41[757] := p3;
			v3.f12 := if (v41[757] in v40) then v40[v41[757]] else p2;
		} else {
			var v42: array<T0> := new T0[14];
			var v43: array<bool> := new bool[4](i6 => p3);
			var v44: map<array<T0>, array<bool>> := map[v42 := v43];
			v44 := v44;
			var v45 := DC13(p3, p3);
			var v46: seq<D5> := [v45, v45, v45, v45];
			var v47: map<bool, bool> := map[p3 := p3];
			var v48: set<seq<D5>> := {v46, [v45, v45] + v46, fm26(p0, fm15(v47, v3.f12, globalState), p3, v7, globalState)};
			v48 := v48;
			var v49 := false;
			var v50: array<map<int, array<bool>>> := new map<int, array<bool>>[12];
			var v51: map<int, array<bool>> := map[p1 := v43];
			v50[178] := v51;
			v49, v50[178] := fm1(p0, v49, p0 + -415, p3, globalState), v51;
			var v52 := DC29(v3);
			var v53: array<T2> := new T2[26] [v3, v3, v3, v3, v3, v3, v3, if (!true) then v3 else v3, v3, v3, v3, v3, v3, v3, v3, v3, v3, v3, v3, v3, v3, if (!p3) then v3 else v3, v3, v3, if (v49) then v52.cf39 else v3, v3];
			v53[84] := v3;
			var v54 := DC23(v3.f12);
			var v55 := DC26();
			var v56: C1 := new C1('g', v7, 0x2b8);
			var v57: set<C1> := {v56, v56};
			var v58: map<int, bool> := map[v3.f12 := p3];
			v3, v53[84], v7, v54, v55 := v3, if (v49) then v3 else v3, v7, DC23(p2 % |v57|), fm27(v58, p3, v0 + v6[44], globalState);
			v0 := seq(0x31, i7  => (v56.f15));
		}
		
		var v59: multiset<int> := multiset{p2};
		var v60: map<bool, int> := map[p3 := p1];
		r0 := (v3.f12 / -35) / (if (-p1 in v59) then v59[-p1] else |v60|);
	} else {
		var v61 := true;
		var v62: multiset<bool> := multiset{p3};
		var v63: multiset<int> := multiset{p1, |v62|};
		v61 := multiset{fm0(0x360, p1, globalState)} == (v63 + v63);
		var v64 := DC30(p0);
		globalState.f5 := v64.cf40 - (p0 * p1);
		var v65: array<bool> := new bool[24] [v61, v61, p3, v61, v61, true, false, !v61, !v61, p3, v61, v61, true, p3, p3, p3, p3, v61, p3, false, p3, v61, DC13(v61, p3).cf19, p3];
		var v66: map<bool, array<bool>> := map[v61 := v65];
		var v67 := DC32(v66);
		v66 := v67.cf42 + (v66 + v66);
		globalState.f5 := if (p3) then p2 else p2 * p1;
		var v68 := new C4(p1, p0);
	}
	
	var v69: set<bool> := {p3};
	var v70: seq<int> := [p2, p0];
	var v71: array<int> := new int[17] [0x286 - p1, p2, p1 + |v69|, if (false) then p1 else |[p0]|, p0, p2, p1, p1 + p1, p1, |v70|, p1, |[p3, p3, p3]|, 0x28c / p0, p0, p0, p2, p2 % p2];
	v71 := v71;
	var v72 := false;
	var v73 := 'y';
	var v74: C1 := new C1(v73, v73, p2);
	var v75 := "g";
	v72, v72, v72, v74, globalState.f5 := if (v72) then v72 else -|v75| > p1, if (p3) then v72 else v72, v72, v74, p2;
	var v76: array<bool> := new bool[21](i8 => -476 <= |v75[-p1 := v74.f15]|);
	v76[94] := p3;
	var v77: map<int, string> := map[|(seq(-0x289, i9  => (v74.f15)))| := v75];
	var v78 := DC4(v71);
	var v79: set<array<int>> := {v78.cf7, v71};
	v76[94], r0 := fm1(p0, v72, v70[p0], p3, globalState), fm0(|(if (-0x78 in v77) then v77[-0x78] else v75)| % p2, |(if (p3) then v79 else v79)|, globalState);
	var v80 := new C4(-263 + p0, p2);
	v72 := true;
	r0 := v80.f11 % v80.f10;
	var v81: map<bool, int> := map[p3 := 821];
	r1 := v81;
	var v82: map<int, int> := map[v80.f11 := p0];
	var v83: seq<string> := [seq(0x15b, i10  => ('d')), v75, (v75[-p1 := v73] + v75)[|v82| := v73]];
	r2 := v83[-0x37a];
}
trait T0 {
	var f12 : int
}

trait T1 extends T0 {
	const f13 : char
	function fm4(globalState: GlobalState): int 
	function fm5(p0: bool, p1: bool, p2: bool, p3: bool, globalState: GlobalState): seq<int> 
	method m2(globalState: GlobalState) returns (r0: bool, r1: array<int>, r2: map<char, bool>, r3: int) 
}

trait T2 extends T0 {
	function fm6(globalState: GlobalState): D0 
}

class C0 extends T2 {
	var f16 : map<int, bool>
	constructor (f16 : map<int, bool>, f12 : int) {
		this.f16 := f16;
		this.f12 := f12;
	}
	
	function fm6(globalState: GlobalState): D0 {
		DC0("qxlwa" != (seq(0x172, i0  => ('t'))))
	}
	function fm8(p0: bool, p1: bool, p2: int, p3: bool, globalState: GlobalState): string {
		"lnadqpvyl" + (if (false) then "yutm" else "slatllus")
	}
}

class C1 extends T1 {
	const f15 : char
	constructor (f15 : char, f13 : char, f12 : int) {
		this.f15 := f15;
		this.f13 := f13;
		this.f12 := f12;
	}
	
	function fm4(globalState: GlobalState): int {
		f12
	}
	function fm5(p0: bool, p1: bool, p2: bool, p3: bool, globalState: GlobalState): seq<int> {
		if (false) then [f12] else if (true) then seq(155, i0  => (f12)) else [0xea, f12]
	}
	method m2(globalState: GlobalState) returns (r0: bool, r1: array<int>, r2: map<char, bool>, r3: int) {
		var v0: array<int> := new int[12](i0 => i0 * f12);
		v0[216] := f12;
		v0[216] := f12;
		var v1 := false;
		var v2 := DC5(v1);
		v2 := DC5(v1);
		if (if (v1) then v1 else v1) {
			var v3: set<bool> := {v1, v1};
			var v4: map<int, bool> := map[|v3| := false];
			var v5: T2 := new C0(v4, v0[216]);
			var v6: seq<bool> := [v1, v1, v1];
			var v7: seq<array<int>> := [v0];
			r0, r0, r0, globalState.f5 := true, v1 ==> (v5 in [v5]), !v6[-v5.f12] <== v1, |((v7 + v7) + v7)[f12 := v0]|;
			var v8: seq<int> := [f12];
			var v9: multiset<int> := multiset{v5.f12};
			var v10: map<bool, multiset<int>> := map[v1 := v9];
			var v11: map<map<bool, multiset<int>>, bool> := map[v10 := v1];
			var v12: map<bool, int> := map[v1 := v0[216]];
			var v13 := DC7(|map[if (f12 in v9) then v9[f12] else v0[216] := if (v1 in v12) then v12[v1] else f12]|);
			v1, globalState.f5, r0, v8 := v1, v5.f12, if (v10 in v11) then v11[v10] else !(v1 || v1), v8[if (v1) then v13.cf10 else -345 := v5.f12];
			if (false) {
				v1, r0 := if (v1) then v1 else v1 !in v6, v1;
				var v14 := new C0(v4, DC9(-0x342).cf14);
				var v15 := 'n';
				v15 := f13;
				var v16 := "iiotodfiq";
				v16 := (v16 + v16) + v16;
				v6 := fm9(!v1, globalState) + [v1, true, v1];
			} else {
				var v17 := "q";
				var v18: seq<string> := [v17, seq(0x2ce, i1  => (f15)), v17 + v17];
				var v19: map<int, int> := map[-v5.f12 := 0xd2];
				v17 := v18[0x1b8 + (if (|v17| in v19) then v19[|v17|] else v0[216])];
				var v20: C0 := new C0(v4, v0[216]);
				v20 := v20;
				var v21: array<T2> := new T2[22];
				v21 := v21;
				var v22: multiset<seq<bool>> := multiset{v6};
				var v23: seq<seq<bool>> := [v6];
				v22 := multiset(v23);
				v20.f16 := v20.f16;
			}
			
			v1 := "h" == "lbyqqyufh";
			var v24: C0 := new C0(map[f12 := v1], v0[216]);
			var v25: multiset<C0> := multiset{v24, v24, v24};
			var v26: map<bool, char> := map[v1 := f15];
			var v27 := DC2(v1, fm10(|v25|, globalState), v26, true);
			var v28: map<D0, int> := map[v27 := f12];
			globalState.f5 := (if (true) then if (v27 in v28) then v28[v27] else v0[216] else v5.f12) / (v5.f12 - v0[216]);
		} else {
			if (if (v1) then v1 else false) {
				var v29: map<char, bool> := map['n' := v1];
				var v30 := "ynbd";
				v1 := (if ('w' in v29) then v29['w'] else v1) ==> fm1(|v30|, fm1(f12, !v1, v0[216], v1, globalState), v0[216], v1, globalState);
				var v31: map<int, bool> := map[-503 := v1];
				var v32 := new C0(v31 + v31, v0[216]);
				v30 := v30;
				r0 := v1;
				r0 := fm1(|(v30 + (seq(198, i2  => (f13)))[|multiset{f12}| := 'q'])|, f12 < -|"r"|, v0[216], v1 && v1, globalState);
			} else {
				var v33: seq<bool> := [v1, v1, v1];
				var v34 := "cqtlrd";
				var v35: seq<string> := [v34];
				var v36 := DC0(true);
				var v37 := DC3(v36);
				var v38: array<bool> := new bool[23] [v1, v1, false, v1, fm1(f12, v1, -989, v1, globalState), v1, v1, v1, v1, v1, v1, v1, false, v1, v1, v1, false, v1, v1, v1, v1, v1, v1];
				var v39: array<bool> := new bool[19] [v1, v1, v1, v33[f12], v1, fm1(f12, true, 0x2cc, v1, globalState), v1 <== v1, v1 ==> v1, v1, v1, fm11(v0[216], f12, v1, globalState) <= v34, v34 != v35[|v33[--v0[216] := false]|], v1, v1, v37 in map[v37 := v38], v1, v1, !v1, v1];
				r0, globalState.f5, globalState.f5, v39, r0 := v1, -754, v0[216] * v0[216], v39, !(true <== !v1);
				var v40: multiset<int> := multiset{f12, v0[216]};
				var v41: map<bool, bool> := map[v1 := v1];
				var v42: map<int, bool> := map[if (|v41| in v40) then v40[|v41|] else v0[216] := false];
				var v43 := DC10(fm11(fm0(f12, f12, globalState), v0[216], v1, globalState));
				var v44: map<bool, int> := map[v1 := |v43.cf15|];
				var v45 := new C0(v42, |v44|);
				var v46: array<map<bool, bool>> := new map<bool, bool>[7](i3 => v41);
				v46[873] := v41 + map[v1 := v1];
				v46[873] := v41;
				var v47: map<seq<bool>, int> := map[v33 := f12];
				v47 := v47[[fm1(v0[216], v1, f12, v1, globalState)] := f12];
				var v48: set<bool> := {true};
				var v49: map<set<bool>, int> := map[v48 := if (v0[216] in v40) then v40[v0[216]] else |v34|];
				var v50 := 's';
				var v51: array<map<bool, int>> := new map<bool, int>[2](i4 => v44);
				v51[647] := v44;
				var v53: seq<set<bool>> := [v48];
				var v54 := DC11(v48);
				var v55 := DC8(v1, v0[216], v1);
				v49, v50, v51[647] := ((map v52 : set<bool> | v52 in v53 :: (v52) := (v0[216]))[v54.cf16 := v55.cf12])[{v1, v1, v1} := fm0(f12, v0[216], globalState)], v50, v44;
			}
			
			var v56: seq<bool> := [v1];
			var v57 := DC7(0x2c2);
			var v58: multiset<int> := multiset{|([v56, [false]] + (seq(97, i5  => (v56))))|, f12, v0[216], v57.cf10, f12 - f12};
			v58 := v58;
			v1 := v0[216] <= f12;
			var v59: map<int, bool> := map[v0[216] := v1];
			var v60: multiset<bool> := multiset{v1, false, v1, v1, v1};
			var v63: array<map<int, bool>> := new map<int, bool>[13] [v59 + map[v0[216] := v1], map[v0[216] := v1], v59 + v59, v59, v59, fm10(|v60|, globalState), v59, v59 + v59, v59, v59, map v61 : int | (14 <= v61) && (v61 < 0x2b2) :: (v61 % f12) := (!v1), v59, map v62 : int | (0x17e <= v62) && (v62 < 85) :: (v62 + f12) := (v1)];
			v63[948] := v59[f12 := v1];
			v63[948] := v59;
			var v64: array<D2> := new D2[4](i6 => DC9(524));
			var v65: map<multiset<int>, array<D2>> := map[v58 := v64];
			v65 := v65[multiset{f12} := v64];
		}
		
		var v66: map<int, bool> := map[f12 := !!true];
		var v67: multiset<bool> := multiset{v1, v1};
		var v68: map<int, int> := map[20 := |v67|];
		var v71: seq<set<int>> := [{-|{v66}|}, set v69 : int | v69 in v68 :: (v69 % |map[DC2(!true, map v70 : int | v70 in [|map[0xf2 := false]|, 0xc0] :: (v70 + 0x353) := (!true), map[true := 'f'], true).cf2 := true]|)];
		var v72 := "lvadps";
		var v73: set<int> := {f12};
		if (v71[|v72|] == (v73 - v73)) {
			var v74: seq<int> := [v0[216]];
			r0 := fm1(v0[216], false, v0[216], v0[216] > |v74|, globalState);
			var v75: array<bool> := new bool[14];
			var v76 := DC0(v1);
			v75[89] := v76.cf0;
			var v77: array<D3> := new D3[28];
			var v78: set<array<D3>> := {v77, v77};
			v0[216], v75[89], r3, r3 := |v78|, -v0[216] <= (v0[216] % v0[216]), f12, v0[216];
			var v79 := new C0(map[550 := true], 688);
			var v80: array<seq<seq<bool>>> := new seq<seq<bool>>[29];
			var v81: seq<seq<bool>> := [[v75[89]], [v75[89], v1]];
			v80[563] := v81;
			v80[563] := v81;
			v75 := v75;
		} else {
			var v82 := new C0(map[|v72| := true], f12);
			var v83 := new C0(v66, f12);
			v83 := v83;
			var v84: array<string> := new string[11](i7 => v72 + v72);
			v84[857] := "r";
			var v85 := DC10(v72);
			v84[857] := (if (v1) then v72 else "hu") + v85.cf15;
			v84[857] := (v72 + "y") + "bsdwgeit";
		}
		
		var v86: seq<bool> := [v1];
		var v87 := DC7(|v86|);
		var v88: seq<int> := [v87.cf10];
		var v89 := DC4(v0);
		var v90: map<int, array<int>> := map[v88[f12] := v89.cf7];
		v90 := v90[v0[216] + v0[216] := v0];
		for i8 := v0[216] to f12 {
			var v91: set<seq<bool>> := {v86};
			var v92: map<seq<bool>, multiset<bool>> := map[[v1, v1] := multiset{v1, v1, false, fm1(|v73|, v1, f12, v1, globalState)}];
			if (if (true) then v91 > (set v93 : seq<bool> | v93 in v92 :: (v93)) else !!(if (v1) then v1 else v1)) {
				v72 := v72 + (seq(-898, i9  => (f15)));
				v72 := v72;
				r0 := v1;
				var v94 := 'p';
				v94 := f13;
				var v95: array<D0> := new D0[18];
				v95 := v95;
			} else {
				var v96: array<map<bool, D1>> := new map<bool, D1>[7];
				var v97: map<array<map<bool, D1>>, bool> := map[v96 := v1];
				v97 := v97[v96 := v1 <== fm1(f12, !v1, i8, v1, globalState)];
				var v98: T2 := new C0(map[|v88| := v1], f12);
				var v99: seq<T2> := [v98, v98, v98, v98, v98];
				var v100: map<seq<T2>, int> := map[v99 := v98.f12];
				v100 := v100[[v98, v98, v98, v98] := f12 * v0[216]];
				v0[216] := |v72[fm4(globalState) := f15]|;
				v68 := fm12(-(f12 - v0[216]), v98.f12, f12, fm0(|v86|, |v68|, globalState) + v0[216], globalState);
				var v101: set<map<int, int>> := {v68, v68, v68};
				var v102: array<bool> := new bool[13] [v1, v1, v1, v1, v1, v1, true, !v1, v1, v1 in v86, v1, v1, v101 >= v101];
				v102 := if (if (|v68| in v66) then v66[|v68|] else v1) then v102 else v102;
			}
			
			v88 := v88;
			v1 := !true;
			var v103 := DC0(v1);
			v1 := !v103.cf0;
		}
		r0 := !fm1(f12, v1, f12 * f12, false, globalState);
		r1 := v0;
		var v104: set<seq<int>> := {v88};
		var v105: multiset<int> := multiset{v0[216], f12};
		var v106: map<char, bool> := map['b' := !(if (|v104| in v66) then v66[|v104|] else fm1(f12, v1, |v105|, v1, globalState))];
		r2 := v106 + v106;
		r3 := |v72|;
	}
	method m5(globalState: GlobalState) returns (r0: int, r1: bool, r2: map<set<bool>, multiset<int>>) {
		var v0: set<int> := {f12, |"lqyamhoax"|};
		r1 := v0 >= v0;
		f12 := match DC9(f12) {
			case DC8(cf11, cf12, cf13) => cf12
			case DC9(cf14) => cf14
			case DC7(cf10) => cf10
		};
		var v1 := true;
		var i0 := 0;
		while (v1)
			decreases 100 - i0
		{
			if (i0 >= 100) {
				break;
			}
			
			i0 := i0 + 1;
			var v2: seq<bool> := [v1, v1];
			var v3 := "iy";
			var v4: map<int, bool> := map[if (v1) then f12 else f12 := fm0(|v2|, |v3|, globalState) <= f12];
			v4 := v4[f12 := !v1];
			var v5: array<array<bool>> := new array<bool>[10];
			var v6 := DC8(v1, f12, v1);
			var v7: array<bool> := new bool[5] [v6.cf11, v1, v1, v1, v1];
			v5[717] := v7;
			v1, v5[717] := !v1 <== (f12 in fm13(v1, v1, globalState)), v7;
			var v8: set<bool> := {v1};
			var v9 := DC9(f12);
			v8 := fm14(f12, v9, globalState);
			var v10: seq<int> := [f12];
			var v11: multiset<char> := multiset{'j', f15, f15};
			globalState.f5 := f12 % v10[if (f15 in v11) then v11[f15] else f12];
		}
		r1 := if (v1) then if (v1) then v1 else v1 else !v1;
		var v12: seq<int> := [f12, 164 / f12, f12];
		var v13: seq<bool> := [true, v1];
		var v14 := "iiunb";
		v12 := (fm5(v13 <= v13, v1, v1, true, globalState))[|v14| - f12 := -(f12 * f12)];
		var v15: array<array<set<int>>> := new array<set<int>>[26];
		var v16: array<set<int>> := new set<int>[25](i1 => v0);
		v15[52] := v16;
		var v17: map<bool, array<set<int>>> := map[v1 := v16];
		v15[52] := if (fm1(-|v14|, v1, f12, v1, globalState) in v17) then v17[fm1(-|v14|, v1, f12, v1, globalState)] else v16;
		r0 := -match DC0(fm1(f12, !v1, f12, v1, globalState)) {
			case DC1(cf1) => 0x250 / f12
			case DC2(cf2, cf3, cf4, cf5) => f12 + f12
			case DC0(cf0) => f12
			case DC3(cf6) => 0x67
		};
		var v18 := DC5(v1);
		var v19: map<bool, int> := map[v18.cf8 := 0x26f];
		r1 := if (!v1) then v1 else fm1(f12, !v1, -|v19|, v1, globalState);
		var v20: set<bool> := {v1, false, v1, v1};
		var v21: map<int, bool> := map[f12 := v1];
		var v22: map<bool, bool> := map[true := if (421 in v21) then v21[421] else !v1];
		var v23: multiset<int> := multiset{f12, |v22|};
		var v24: map<set<bool>, multiset<int>> := map[{v1} := v23];
		var v26: seq<set<bool>> := [v20];
		r2 := map[v20 := v23] + (v24 + (map v25 : set<bool> | v25 in v26 :: (v25) := (v23)));
	}
}

class C2 extends T0 {
	const f17 : bool
	constructor (f17 : bool, f12 : int) {
		this.f17 := f17;
		this.f12 := f12;
	}
	
}

class C3 extends T1, T2 {
	const f14 : array<D2>
	constructor (f14 : array<D2>, f13 : char, f12 : int) {
		this.f14 := f14;
		this.f13 := f13;
		this.f12 := f12;
	}
	
	function fm4(globalState: GlobalState): int {
		f12 / (632 % |multiset{true}|)
	}
	function fm5(p0: bool, p1: bool, p2: bool, p3: bool, globalState: GlobalState): seq<int> {
		[f12] + [f12]
	}
	function fm6(globalState: GlobalState): D0 {
		DC0(false)
	}
	function fm7(globalState: GlobalState): int {
		|((if (false) then "kdgs" else "xerenpimg") + ("sugwvjq" + (seq(0x237, i0  => (f13)))))|
	}
	method m2(globalState: GlobalState) returns (r0: bool, r1: array<int>, r2: map<char, bool>, r3: int) {
		var v0: array<int> := new int[4](i0 => i0 + f12);
		v0[341] := f12 % f12;
		var v1 := true;
		var v2 := DC8(v1, f12, v1);
		v0[341] := match v2 {
			case DC8(cf11, cf12, cf13) => cf12 + |[true]|
			case DC9(cf14) => f12
			case DC7(cf10) => f12
		};
		var i1 := 0;
		while (v1)
			decreases 100 - i1
		{
			if (i1 >= 100) {
				break;
			}
			
			i1 := i1 + 1;
			var v3 := "dii";
			globalState.f5 := fm0(|v3|, v0[341], globalState) + |v3[v0[341] := f13]|;
			var v4: map<int, bool> := map[0x35a := true];
			var v5 := new C0(v4, fm0(v0[341], v0[341], globalState) / v0[341]);
			var v6: seq<bool> := [v1];
			v6 := if (v1) then v6 else ([v1, v1, v1, v1, v1])[-f12 := v1];
			var v7: array<array<char>> := new array<char>[8];
			var v8: array<char> := new char[20];
			v7[302] := v8;
			var v9: seq<array<char>> := [v8];
			var v10: multiset<char> := multiset{f13, f13, f13, f13};
			v7[302] := v9[if (f13 in v10) then v10[f13] else f12];
		}
		var v11: map<int, array<int>> := map[v0[341] := v0];
		v11 := v11[-(0x1f1 - v0[341]) := v0];
		r3 := v0[341];
		v1 := v0[341] < v0[341];
		r0 := v1;
		var v12: multiset<bool> := multiset{v1, v1};
		var v13 := DC12([v12, v12]);
		r0 := (v13.cf17 + (seq(0x25d, i2  => (v12)))) <= (seq(0x39c, i3  => (v12)));
		r1 := v0;
		var v14: map<char, bool> := map[f13 := v1];
		r2 := v14 + v14;
		var v15: set<bool> := {v1, fm1(f12, v1, v0[341], v1, globalState)};
		var v16: set<set<bool>> := {{v1, v1, v1, false}, {v1, false, v1, v1}, v15};
		r3 := |v16|;
	}
	method m3(p0: multiset<bool>, globalState: GlobalState) returns (r0: int, r1: bool, r2: int) {
		var v0 := true;
		var i0 := 0;
		while (v0)
			decreases 100 - i0
		{
			if (i0 >= 100) {
				break;
			}
			
			i0 := i0 + 1;
			var v1: map<int, int> := map[0x386 := f12];
			var v2: array<array<bool>> := new array<bool>[7];
			var v3: seq<int> := [f12, 0x34a];
			var v4: seq<int> := [f12, |v3|];
			var v5: array<bool> := new bool[22] [v0, true, v0, v0, v0, fm1(f12, v0, f12, v0, globalState), v0, v0, v0, v0, v0, v0, false, v0, v0, fm1(|v3|, v0, 0x81, v0, globalState), fm1(f12, v0, f12, v0, globalState), v0, v0, v0, v0, v0];
			v2[919] := v5;
			v1, v2[919] := v1, if (v0) then v5 else v5;
			var v6 := "doktuoubt";
			if ((|map[|map[-|v6| := f12]| := v0]| % 0xb) == |(v6 + (seq(94, i1  => (v6[f12]))))|) {
				var v7: map<bool, int> := map[v0 := f12];
				globalState.f5 := (|v7| / f12) * f12;
				f12 := f12;
				var v8: array<int> := new int[27];
				v8[331] := |[v8]| / f12;
				v8[331] := --f12;
				var v9: seq<string> := [seq(-0x28f, i2  => (f13))];
				var v10 := DC8(v0, v8[331], v0);
				var v11: map<int, D2> := map[|v9[v8[331]]| := v10];
				var v12: seq<bool> := [v0];
				v8[331] := |((v11[f12 := v10.(cf11 := v0)])[|v6| := v10])[v8[331] := DC8(v12[f12], v8[331], !!v0)]|;
				var v13: set<char> := {f13, f13, v6[0x16b], 'x'};
				var v14: set<bool> := {v0};
				var v15: seq<set<bool>> := [v14, v14];
				var v16 := DC15(v4);
				var v17: map<set<char>, seq<int>> := map[v13 := if (v0) then [|v15[f12]|, -fm0(v8[331], f12, globalState)] else v16.cf21];
				v17 := map[v13 + v13 := v4];
			} else {
				var v18: set<seq<int>> := {v3};
				var v19: map<set<seq<int>>, string> := map[v18 := v6];
				v0 := !(("hvvjigs" < (if (v18 in v19) then v19[v18] else v6)) <== v0);
				var v20: array<int> := new int[11];
				v20[737] := f12;
				var v21 := DC0(!v0);
				var v22 := DC3(v21);
				v20[663] := 0x37e;
				v20[928] := f12;
				v20[66] := -(if (f12 in v1) then v1[f12] else |v6|);
				v20[737], v22, v20[663], v20[928], v20[66] := f12, v22.(cf6 := DC3(v21)), f12, f12, |(if (v0 || !v0) then p0 else p0)|;
				r1 := v0;
				var v23: set<bool> := {v0, v0};
				f12 := if (v0) then v20[737] else |(v4 + fm5(v0, v0, v0, fm1(if (f12 in v1) then v1[f12] else |v23|, v0, v20[737], true, globalState), globalState))|;
				r1 := !v0;
			}
			
			v0 := false;
			var v24 := DC14(v0);
			var v25: set<bool> := {v0};
			var v26: multiset<int> := multiset{f12, |v25|};
			match (v24.(cf20 := v26 >= v26)) {
				case DC13(cf18, cf19) =>
					var v27: array<int> := new int[1] [f12 * v4[f12]];
					v27[237] := f12;
					v27[237] := f12;
					v6 := "nbfyp";
					globalState.f5 := -f12;
					r1, r2 := cf18, -f12;
				case DC14(cf20) =>
					var v28: seq<bool> := ['n' in v6];
					v0 := v28[f12];
					var v29: multiset<string> := multiset{"i", "tf", "tbjwfs", v6};
					r1 := fm1(f12, !fm1(f12, cf20, f12, v0, globalState), |(v29 * multiset{seq(0x28d, i3  => (f13)), seq(-0x2e1, i4  => (f13))})|, cf20, globalState);
					var v30: map<bool, bool> := map[false := v0];
					var v31: seq<set<bool>> := [v25, v25, v25, v25, {true, cf20, v0}];
					var v32: map<seq<set<bool>>, char> := map[v31 := f13];
					var v33: array<char> := new char[10] [f13, fm15(v30, |v6|, globalState), if (v31 in v32) then v32[v31] else f13, f13, f13, f13, f13, f13, f13, f13];
					v33, globalState.f5 := v33, -f12 - f12;
					v0 := cf20;
				case DC12(cf17) =>
					v6 := v6[f12 := f13];
					var v34: array<int> := new int[1](i5 => i5 + -|(seq(0x1eb, i6  => (f13)))|);
					v34 := if (v0 <==> v0) then v34 else if (v0) then v34 else v34;
					r1 := !!v0;
					var v35 := new C1(f13, f13, f12);
			}
			
		}
		var v36: map<bool, bool> := map[v0 := v0];
		var v37 := "gpstgyprm";
		var v38: seq<bool> := [fm1(|v36|, v0, |v37|, v0, globalState)];
		var v39: set<bool> := {fm1(f12, false, |v38|, v0, globalState), v0, if (v0 in v36) then v36[v0] else v0, v0};
		match (DC11(v39)) {
			case DC11(cf16) =>
				var v40: array<bool> := new bool[8] [!fm1(0x30f, v0, f12, v0, globalState), v37 <= v37, true, v0, v0, f12 > f12, !(!v0 || v0), v0];
				v40[553] := v0;
				var v41: seq<int> := [|v37|];
				var v42: map<bool, char> := map[v0 := f13];
				var v44: set<char> := {f13};
				v40[553] := if (v0) then |v41| == f12 else (if (v0 in v42) then v42[v0] else f13) !in (map v43 : char | v43 in v44 :: (v43) := (v0));
				var v46: map<int, int> := map[f12 := f12];
				var v47: map<int, int> := map[f12 := |v46|];
				var v48: map<int, bool> := map[|v47| := v0];
				var v49 := new C0((map v45 : int | v45 in v41 :: (v45 - --485) := (!v40[553])) + v48, f12);
				var v50: array<int> := new int[22](i7 => i7 - f12);
				v50[919] := f12;
				v50[919] := f12 + |cf16|;
				v49.f16 := v49.f16[437 := f12 !in fm5(!v0, v0, v0, v0, globalState)];
		}
		
		var v51: array<bool> := new bool[27];
		v51[348] := v0;
		var v52: seq<multiset<bool>> := [p0, p0];
		var v53 := DC12(v52);
		var v54: seq<D5> := [v53];
		v51[348] := ([v53] + v54[f12 := v53]) <= v54;
		var v55: multiset<seq<bool>> := multiset{v38, [v51[348]], v38};
		var v56: set<int> := {f12, f12};
		var v57: seq<int> := [f12, f12, |v56|, f12];
		var v58 := DC15(v57);
		v0 := (v55 + fm16(v0, f12, v58, f12, globalState)) >= (v55 * v55);
		if ((if (true) then v37 else v37) != v37) {
			var v59 := DC8(v51[348], f12, v0);
			var v60: array<D2> := new D2[29] [v59, v59, v59, v59, v59, fm17(v0, fm1(f12, v0, f12, v0, globalState), globalState), v59, v59, v59, v59.(cf13 := v0), v59, v59, v59, DC8(false, |v38|, v0), v59, v59, v59, v59, v59, v59, v59, v59, v59, v59, v59, fm17(v0, v0, globalState), v59, v59, v59];
			var v61: seq<array<D2>> := [v60, v60];
			v61 := if (v51[348]) then v61 else v61;
			var v62, v63, v64 := m0(f12, 0x71 - -f12, f12, v51[348], globalState);
			v0 := v51[348];
			var v65: map<seq<bool>, bool> := map[[!v0] := v0];
			v65 := v65[v38 := v51[348]];
			v51[348] := v0;
		} else {
			v51[348] := v51[348];
			var v66: array<map<int, int>> := new map<int, int>[27](i8 => map[f12 := -f12]);
			var v67: map<int, bool> := map[f12 := false];
			var v68: seq<seq<bool>> := [v38, [v51[348], false, false, v0]];
			var v69: map<int, int> := map[|v67| := DC9(fm0(f12, |multiset(v68[0x198])|, globalState)).cf14];
			v66[218] := v69 + map[f12 := f12];
			var v71: multiset<char> := multiset{f13};
			v66[218], v37, v0, v37 := map v70 : int | v70 in v57 :: (v70 * 0x385) := (f12), v37, (multiset{f13, 'f', f13} - v71) >= multiset(fm11(f12, |multiset(v57)|, v51[348], globalState)), v37;
			var v72: array<array<bool>> := new array<bool>[9];
			var v73: map<bool, array<array<bool>>> := map[fm1(f12, false, f12, false, globalState) := v72];
			v73 := v73[v0 := if (v0 in v73) then v73[v0] else v72];
			var v74 := 'o';
			v74 := fm15(v36, v57[|v37|], globalState);
			v56 := v56;
		}
		
		var i9 := 0;
		while (v38[-f12 + f12])
			decreases 100 - i9
		{
			if (i9 >= 100) {
				break;
			}
			
			i9 := i9 + 1;
			globalState.f5 := 479;
			var v75: map<int, bool> := map[f12 := v51[348]];
			v75 := v75[f12 := v51[348] <== v51[348]];
			var v76 := 'u';
			v76 := v76;
			v0 := if (v51[348]) then v51[348] !in {v0, v0} else f12 == -f12;
		}
		r0 := 0x340 * f12;
		r1 := v51[348];
		var v77: multiset<string> := multiset{v37[f12 := f13], v37};
		r2 := (if ("ptyrk" in v77) then v77["ptyrk"] else f12) % (f12 * f12);
	}
	method m4(p0: int, p1: bool, p2: string, globalState: GlobalState) returns (r0: bool, r1: bool, r2: int) {
		for i0 := p0 to f12 - f12 {
			var v0: map<bool, int> := map[false := f12];
			var v1: multiset<map<bool, int>> := multiset{v0, v0};
			v1 := multiset([v0]);
			var v2: array<array<int>> := new array<int>[29];
			var v3 := DC8(p1, -238, p1);
			var v4: set<bool> := {v3.cf11};
			var v5: map<int, bool> := map[p0 := p1];
			var v6: T0 := new C2(true, -i0);
			var v7: seq<int> := [p0, p0, p0, p0, |map[p0 := v6]|];
			var v8: map<int, int> := map[DC8(p1, p0, p1).cf12 := f12];
			var v9: array<int> := new int[25] [f12, p0, |fm18(|p2|, 421, p1, globalState)|, f12, f12, p0, |v4|, |p2|, f12, f12, i0, 0xcb, p0, i0, p0, i0, -|p2|, 0x3c8, f12, i0, |v5[|v7| := p1]|, p0, p0, |v8|, i0];
			v2[47] := v9;
			var v10 := DC15(if (p1) then v7 else v7);
			r2, v2[47], r1, r1, v10 := v6.f12, v9, p1, p1, DC15(seq(0x2ac, i1  => (|p2|)));
			v2[47] := new int[8];
			var v11: seq<bool> := [p1, p1];
			var v12: array<bool> := new bool[16] [p1, fm1(p0, p1, -0xc3, p1, globalState), p1, p1, p1, p1, p1, p1, false, p1, v11[p0], p1, p1, p1, p1, p1];
			var v13: map<int, array<bool>> := map[i0 := v12];
			var v14: seq<array<bool>> := [v12];
			v13 := v13[f12 := v14[v7[f12]]];
		}
		if (p1) {
			var v15: set<int> := {f12};
			var v16: map<bool, int> := map[p1 := f12];
			var v17: map<int, bool> := map[f12 := p1];
			var v18: T2 := new C0(v17, f12);
			var v19: seq<T2> := [v18];
			var v20: multiset<bool> := multiset{false};
			var v21: array<int> := new int[29] [p0, f12, -f12, f12, |v15|, p0, p0, p0, f12, f12, p0, f12, if (p1 in v16) then v16[p1] else p0, p0, |v19[|v20| := v18]|, p0, p0, -p0, v18.f12, f12, p0, v18.f12, f12, -p0, f12, p0, p0, f12, v18.f12];
			var v22 := DC4(v21);
			var v23: map<map<bool, array<int>>, bool> := map[map[p1 := v22.cf7] := if (p1) then p1 else p1];
			var v24: map<bool, array<int>> := map[p1 := v21];
			var v25 := DC20(v24);
			var v26: seq<bool> := [!p1];
			v23 := v23[v25.cf27 := !fm1(|multiset(v26)|, false, p0, p1, globalState)];
			r2 := f12;
			r0 := p1;
			var v27 := 'b';
			v27 := f13;
			var v28: array<bool> := new bool[4];
			v28[175] := p1;
			v28[175] := p1;
		} else {
			var v29 := "cggpcap";
			v29 := v29;
			var v30: C0 := new C0((map[f12 := true])[0x19 := p1], f12);
			v30 := v30;
			if (p1) {
				var v31 := 'e';
				v31 := f13;
				var v32: seq<bool> := [!p1];
				var v33: map<bool, bool> := map[v32[0x203] := p1];
				var v34: set<bool> := {p1, false, p1};
				var v35: array<bool> := new bool[22] [p1, map[p1 := p1] == v33, fm1(f12, p1, p0, p1, globalState), p1, p1, p0 >= p0, p1 <== p1, p1, p1, p1, p1, p1, p1, p1, p1, if (p1) then p1 else p1, p1, v34 !! v34, true, p1, p1, if (p1) then p1 else p1];
				v35[581] := p1;
				var v36: seq<int> := [-803];
				v35[581] := |{v35}| != |v36|;
				var v37: array<string> := new string[5](i2 => v29 + p2);
				v37[252] := p2;
				r0, v37[252], v31, v35[581] := p1, p2 + "b", 'd', p1;
				var v40 := new C0(v30.f16 + (map v38 : int | v38 in (set v39 : int | (0x133 <= v39) && (v39 < 0x7b) :: (v39 / f12)) :: (v38 - p0) := (v35[581])), f12);
				v36 := [f12, f12 % 0x1dc, p0];
			} else {
				var v41: map<bool, bool> := map[true := p1];
				var v42: multiset<int> := multiset{|(seq(0x28d, i3  => (f13)))|, |v41|};
				var v43: seq<int> := [-|(multiset{v30.f16})[v30.f16 := -p0]|, |[p1, p1]|];
				var v44: set<seq<int>> := {v43};
				var v45: map<multiset<int>, int> := map[v42 := |v44|];
				v45 := map[v42 := f12];
				globalState.f5 := p0;
				v29 := p2[p0 := f13];
				var v46: array<int> := new int[16];
				v46[47] := p0;
				v46[47] := -|v29|;
				var v47: map<bool, char> := map[p1 := f13];
				var v48: array<map<bool, char>> := new map<bool, char>[29] [v47, map[p1 := 'v'], v47, map[p1 := f13], v47[!p1 := f13] + v47, v47 + v47, v47, v47, v47 + v47, v47, v47, v47 + v47, if (p1) then v47 else map[p1 := f13], v47 + v47, v47, map[p1 := f13], v47, (v47[p1 := f13])[p1 := 't'] + v47, v47, v47, v47, v47, v47, v47, v47, v47[false := f13], v47, v47 + v47, v47];
				v48[706] := (map[p1 := f13])[p1 := f13];
				v48[706] := v47;
			}
			
			var v49: array<D6> := new D6[6];
			var v50: map<bool, bool> := map[p1 := p1];
			var v51: array<bool> := new bool[28] [p1, p1, p1, false, false, p1, false, p1, p1, p1, p1, p1, true, p1, p1, p1, p1, p1, p1, p1, if (p1 in v50) then v50[p1] else p1, p1, p1, p1, p1, false, !false, p1];
			var v52 := DC16(v51, f12);
			var v53 := DC19(v52);
			v49[651] := v53;
			var v54: array<char> := new char[10](i4 => f13);
			v54[233] := 'l';
			var v55: map<int, int> := map[p0 := 0x360];
			r1, v49[651], v54[233], v55 := p1, v53, 'k', v55 + v55;
			var v56: seq<bool> := [p1];
			var v57: seq<seq<bool>> := [v56];
			var v58: array<seq<bool>> := new seq<bool>[9] [v56, v56, [p1], v56, v56, v56, v57[|v57[p0 := fm9(p1, globalState)]|], v56, v56];
			var v59: seq<array<seq<bool>>> := [v58];
			var v60: array<array<seq<bool>>> := new array<seq<bool>>[6] [v58, v58, v58, v58, v58, v59[-f12]];
			v60[871] := v58;
			v60[871] := v58;
		}
		
		var v61: multiset<int> := multiset{p0};
		var i5 := 0;
		while (p0 !in v61)
			decreases 100 - i5
		{
			if (i5 >= 100) {
				break;
			}
			
			i5 := i5 + 1;
			globalState.f5, r2 := -f12, f12;
			var v62 := DC14(p1);
			match (v62) {
				case DC13(cf18, cf19) =>
					var v63: map<int, bool> := map[f12 := !p1];
					cf18 := if (f12 in v63) then v63[f12] else !cf19;
					r1 := !!p1;
					globalState.f5 := if (cf19) then 0x1fb else f12;
					var v64: array<int> := new int[25];
					v64[197] := p0 % |p2|;
					v64[197] := -p0;
				case DC14(cf20) =>
					var v65: seq<bool> := [p1];
					var v66: array<bool> := new bool[21] [!cf20, p1, cf20, p1, !p1, cf20, p1, !(p0 < p0), p1, !cf20, p1, cf20, v65[p0], cf20, p1 ==> p1, if (p1) then false else cf20, cf20, !(if (cf20) then cf20 else p1), f12 >= f12, p1 <==> cf20, cf20];
					v66[866] := p0 == p0;
					var v67: seq<int> := [850];
					v66[866] := [f12] < (if (p1) then v67 else fm5(p1, cf20, cf20, !p1, globalState));
					globalState.f5 := p0;
					var v68 := DC8(cf20, p0, cf20);
					v66[866] := !fm1(0x2f5, fm1(f12, v66[866], f12, v66[866], globalState), |map[--788 := p1]|, v68.cf13, globalState);
					var v69: set<map<bool, bool>> := {map[cf20 := !fm1(-0x2e9, v66[866], p0, cf20, globalState)]};
					var v71: multiset<char> := multiset{'j'};
					var v72: array<int> := new int[17] [0xb0, f12, p0 % -f12, if (false) then -|v69| else f12, p0 % 440, p0, f12, p0, f12, |(map v70 : char | v70 in v71 :: (v70) := (cf20))|, p0, --239, f12 + |multiset(v65)|, f12 + p0, f12, 0x218, f12];
					v72 := v72;
				case DC12(cf17) =>
					var v73: C1 := new C1(f13, f13, f12);
					var v74: array<int> := new int[28](i6 => i6 / |{|p2|}|);
					v74[985] := 0x12b;
					var v75 := 'm';
					v73, v74[985], v75 := v73, p0, v73.f15;
					v75 := v73.f15;
					globalState.f5 := v73.fm4(globalState);
					var v76, v77, v78 := m0(v74[985], v74[985] * f12, v74[985] + v74[985], true, globalState);
			}
			
			var v79: map<int, string> := map[-f12 := p2];
			var v80: seq<bool> := [true];
			v79 := v79[|v80| := p2 + p2];
			var v81: array<int> := new int[3];
			v81, f12 := v81, -fm0(p0, f12, globalState);
		}
		for i7 := -0x30b * 0x165 to f12 {
			r0 := v61 >= multiset{p0};
			var v82 := new C0(map[i7 := p1], i7);
			var v83: map<bool, bool> := map[p1 := p1];
			var v84: seq<bool> := [p1, if (!p1 in v83) then v83[!p1] else p1, p1, p1, p1];
			var v85: seq<seq<bool>> := [v84];
			if (|v85| >= f12) {
				var v86: array<bool> := new bool[22];
				v86[844] := p1;
				v86[844] := p2 != p2;
				var v87: array<int> := new int[10];
				v87[923] := i7;
				var v88: set<bool> := {false, v86[844]};
				var v89: map<bool, set<bool>> := map[p1 := v88];
				v87[923] := |(if ((f12 == i7) in v89) then v89[f12 == i7] else v88)|;
				v86[844] := DC0(p1).cf0 <==> v86[844];
				r1 := v86[844];
				var v90: seq<int> := [796];
				var v91: multiset<seq<int>> := multiset{v90};
				v87[923], globalState.f5, v87[923] := i7, if ((v90 + v90) in v91) then v91[v90 + v90] else p0, i7;
			} else {
				globalState.f5 := -i7;
				var v92: multiset<bool> := multiset{p1, p1};
				v92 := multiset{p1};
				var v93: array<bool> := new bool[22];
				v93[5] := fm1(f12, p1, i7, false, globalState);
				v93[5] := p1;
				var v94 := 'j';
				v94 := f13;
				var v95, v96, v97 := m0(f12, -fm0(p0, p0, globalState), 0x196, v93[5], globalState);
			}
			
			var v98 := "b";
			v98 := v98 + ((seq(0x33f, i8  => (f13))) + "hxb");
		}
		var v99: array<bool> := new bool[16](i9 => p1);
		var v100 := DC1(v99);
		match (v100) {
			case DC1(cf1) =>
				v99[979] := !!p1;
				v99[979] := p1;
				var v101 := "uadjarwl";
				v101 := v101;
				var v102 := DC9(p0);
				var v103 := DC11(fm14(-0x296, v102, globalState));
				match (v103) {
					case DC11(cf16) =>
						var v104: map<bool, bool> := map[!p1 := true];
						v104 := v104[fm1(p0, p1, p0, p1, globalState) := v99[979]];
						var v105: seq<int> := [f12, p0, p0, f12 * p0];
						v105 := v105;
						v61 := v61;
						var v106: T1 := new C1('l', f13, 738);
						v106 := v106;
				}
				
				r2, v99[979], f12 := |v101| * p0, p1, f12 / 0x3d3;
			case DC2(cf2, cf3, cf4, cf5) =>
				r0 := v61 <= multiset([|[f12, p0, f12]|, -|p2|]);
				var v107: T0 := new C2(cf2, 0x171);
				v107, cf2 := v107, cf2;
				var v108 := new C1(f13, f13, f12);
				if (!true) {
					v99[100] := cf2;
					v99[100] := !true;
					globalState.f5 := |p2|;
					var v109: set<bool> := {cf2, fm1(p0, cf5, p0, v99[100], globalState)};
					var v110 := DC11(v109);
					v107.f12, v110 := p0, v110;
					var v111 := new C0(cf3 + map[|p2| := v99[100]], p0);
					var v112: array<bool> := new bool[28](i10 => cf5);
					var v113 := DC16(v112, v107.f12);
					var v115: seq<bool> := [false];
					var v116: map<bool, int> := map[v99[100] := p0];
					var v117: seq<int> := [v107.f12, 0x1d7, f12];
					v107.f12, v99[100], v99[100] := v113.cf23, --|(map v114 : int | (945 <= v114) && (v114 < -0x77) :: (v114 % v107.f12) := (v107.f12))| < v107.f12, if (v115[|v116|] ==> false) then cf2 else fm1(|v117|, p1, f12, fm1(0x262, v99[100], v107.f12, cf2, globalState), globalState);
				} else {
					var v118: array<int> := new int[11];
					var v119: set<map<int, bool>> := {cf3};
					v118, v107.f12 := v118, |v119| % 0x328;
					r2 := v107.f12;
					r1 := !!cf2;
					var v120 := "nbjki";
					v120 := "c" + p2;
					var v121: seq<int> := [v107.f12];
					var v122: map<int, char> := map[|v121| := v108.f15];
					var v123: multiset<map<int, char>> := multiset{map[-p0 := f13], v122, v122[f12 := 'h'], v122 + v122, v122 + v122};
					v123 := v123;
				}
				
			case DC0(cf0) =>
				var v124: multiset<bool> := multiset{cf0};
				var v125: set<bool> := {v124 >= v124};
				v125 := v125;
				v99[537] := cf0;
				v99[537] := cf0;
				var v126: array<bool> := new bool[19];
				var v127 := DC16(v126, p0);
				match (v127) {
					case DC16(cf22, cf23) =>
						var v128: array<int> := new int[15];
						v128[718] := cf23;
						var v129: seq<array<bool>> := [cf22];
						v128[718] := |(v129 + v129)| * p0;
						var v130: map<int, int> := map[v128[718] := p0 / p0];
						v128[718] := |v130|;
						var v131 := DC8(cf0, 0x120, cf0);
						f12 := fm0(v128[718] % v131.cf12, 37, globalState);
						v99[537] := [cf22, v126] <= v129;
					case DC17(cf24) =>
						r0 := f12 > f12;
						var v132: T0 := new C2(p1, f12);
						var v133: map<T0, int> := map[v132 := 0x397];
						var v135: multiset<char> := multiset{f13, f13};
						var v136: C2 := new C2(false, |(map v134 : char | v134 in v135 :: (v134) := (p0))|);
						v133 := v133[v132 := |map[-0x1fb := v136]| / p0];
						var v137: seq<bool> := [p1];
						var v138, v139, v140 := m3(if (cf0) then multiset(v137) else v124, globalState);
						var v141: seq<int> := [p0, f12];
						var v142: multiset<seq<int>> := multiset{v141, v141, v141};
						v99[537] := (v142 * v142) == v142;
					case DC18(cf25) =>
						cf25 := -cf25;
						var v143 := "r";
						v143 := "kj";
						v99[537] := cf0;
						var v144 := 'l';
						var v145: map<int, bool> := map[f12 := cf0];
						var v146: map<bool, int> := map[cf0 := cf25];
						var v147: C0 := new C0(v145, 0x5f * |v146|);
						var v148 := DC18(f12 % (if (v99[537] in v124) then v124[v99[537]] else cf25));
						f12, v126, v144, v147, v148 := if (0x384 in v61) then v61[0x384] else f12, v126, v144, v147, v148;
					case DC15(cf21) =>
						var v149: multiset<string> := multiset{p2, p2};
						cf0 := v149 !! (v149 - v149);
						var v150 := new C1(f13, f13, -857);
						var v151 := DC15(cf21 + cf21);
						v151 := DC15(seq(872, i11  => (|cf21|)));
						var v152: array<seq<bool>> := new seq<bool>[11](i12 => [v99[537]]);
						var v153: map<int, int> := map[-f12 := f12];
						var v154: set<int> := {-|v153|};
						r0, v99[537], globalState.f5, v152, cf0 := v154 >= v154, v99[537], -p0, v152, false;
					case DC19(cf26) =>
						globalState.f5 := p0;
						var v155: map<multiset<bool>, int> := map[v124 := f12];
						var v156 := DC8(cf0, -p0, p1);
						v155 := v155[fm19(v156.cf12, globalState) := p0];
						v99[537] := cf0;
						var v157: array<int> := new int[16];
						var v158: map<bool, array<int>> := map[v99[537] := v157];
						var v159: array<array<int>> := new array<int>[5] [v157, v157, v157, if (cf0 in v158) then v158[cf0] else v157, v157];
						v159[407] := v157;
						v159[407] := v157;
				}
				
				r1 := !p1;
			case DC3(cf6) =>
				var v160: array<int> := new int[19];
				var v161: set<array<int>> := {v160};
				var v162: map<int, int> := map[p0 := |v161|];
				var v163: map<int, map<int, int>> := map[p0 := v162];
				v163 := v163[|p2| := v162];
				if (p1) {
					var v164: set<bool> := {p1};
					var v165 := DC17(p1);
					var v166: seq<bool> := [v164 >= v164, if (p1) then p1 else p1, 188 < p0, true, v165.cf24];
					r1 := v166[f12];
					v160 := v160;
					var v167: array<map<bool, bool>> := new map<bool, bool>[8];
					var v168: map<bool, bool> := map[p1 := p1];
					v167[274] := v168;
					v167[274] := v168 + (fm18(f12, f12, p1, globalState) + v168);
					r2 := -p0 * f12;
					var v169: multiset<bool> := multiset{p1, true};
					var v170, v171, v172 := m3(v169 - multiset(([p1])[p0 := p1]), globalState);
				} else {
					v160[731] := |"va"|;
					v160[731] := |p2|;
					v160[731] := f12;
					var v173: seq<int> := [p0];
					var v174: map<map<int, int>, int> := map[map[if (f12 in v61) then v61[f12] else v173[|p2|] := |p2|] := v160[731]];
					v174 := if (p1) then v174 else v174;
					var v175: array<string> := new string[23];
					v175[232] := p2 + p2;
					v175[232] := p2 + "kemlgdi";
					var v176: set<bool> := {p1};
					var v177: seq<set<bool>> := [{p1, !true}, v176, v176];
					var v178: set<set<bool>> := {v177[v160[731]], v176, v176};
					var v179 := DC17(true);
					var v180: seq<array<bool>> := [v99, v99, v99];
					v178, v179, v180, globalState.f5, globalState.f5 := v178, v179, (v180 + v180)[v173[f12] := v99], --f12 % f12, -0x152;
				}
				
				var v181: array<string> := new string[3](i13 => p2);
				v181 := if (p1) then v181 else v181;
				var v182: seq<bool> := [p1, p1];
				var v183: set<bool> := {p1};
				var v184: seq<int> := [p0, p0];
				v182 := fm9([|v162[f12 := 0x69]|, p0, -|v183|] <= v184, globalState);
		}
		
		var v185: multiset<bool> := multiset{p1};
		var v186: map<bool, multiset<bool>> := map[!p1 := v185];
		var v187: set<bool> := {true, p1};
		r0 := !fm1(f12, (if (p1 in v186) then v186[p1] else v185) >= v185, |v187|, false, globalState);
		r0 := |v61| > p0;
		r1 := (p1 ==> p1) <== (p1 <==> p1);
		r2 := p0 / f12;
	}
}

class C4 {
	const f10 : int
	var f11 : int
	constructor (f10 : int, f11 : int) {
		this.f10 := f10;
		this.f11 := f11;
	}
	
	function fm2(p0: bool, globalState: GlobalState): int {
		f10
	}
	method m1(p0: string, p1: char, p2: bool, p3: bool, globalState: GlobalState) returns (r0: int, r1: array<bool>, r2: set<array<D0>>, r3: array<int>) {
		var v1: set<int> := {0x3e7, f10, |"stclb"|, -f10, f11};
		var v2 := DC2(p3, map v0 : int | v0 in v1 :: (v0 * f10) := (true), (map[p2 := p1])[p3 := p1], p2);
		match (v2) {
			case DC1(cf1) =>
				r1 := cf1;
				if (if (p3) then p3 else p2) {
					var v3: multiset<int> := multiset{f11};
					var v4 := DC8(p3, |v3|, p2);
					var v5: map<int, bool> := map[f10 := p3];
					var v6: map<bool, int> := map[p2 := f11];
					var v7: seq<int> := [f11];
					var v8: map<int, char> := map[f11 := p1];
					var v9: array<int> := new int[27] [f11, fm0(f10, f10, globalState), f10, |p0| / |p0|, 528, f10, -v4.cf12, f10, |(if (p2) then v5 else fm3(p2, p3, |v6|, globalState))|, f11, f10, 0xab, f10, -0x208, f10, v7[0x292], |(p0 + "tup")|, |v8|, 106 / f11, f10 - f11, f11, f11 / f11, f10, if (p3) then f11 else f11, f10, f11, f11];
					var v10 := DC9(0x30b);
					v9[128] := v10.cf14;
					v9[128] := f11;
					var v11 := true;
					v11 := false;
					var v12: set<multiset<int>> := {v3, v3, multiset{f10}};
					v12, v11, r0, v11 := v12, false, f10 % f11, v11 <== p2;
					var v13: set<bool> := {p2};
					var v14 := new C2(v13 <= v13, v9[128]);
					globalState.f5 := v9[128];
				} else {
					var v15 := false;
					var v16: multiset<array<bool>> := multiset{cf1, cf1};
					v15 := p2 <== (v16 !! multiset{cf1, cf1, cf1});
					var v17: array<int> := new int[3];
					v17[625] := f11;
					v17[625] := f11;
					v15 := (if (v15) then f11 else f10) < |p0|;
					globalState.f5 := v17[625] % f11;
					v15 := p2 || (p0 != p0);
				}
				
				var v18 := DC13(true, false);
				var v19: multiset<bool> := multiset{p3, v18.cf19};
				var v20: seq<int> := [681, f11];
				var v21: seq<multiset<bool>> := [v19, v19, v19[p2 := |v20|], v19, v19];
				match (DC12(v21)) {
					case DC13(cf18, cf19) =>
						var v22: C1 := new C1(p1, p1, f10);
						v22, cf18 := v22, cf19;
						globalState.f5 := 0x73;
						var v23: multiset<int> := multiset{f11, |p0|};
						var v24 := DC12([v19, v19]);
						f11 := |fm20(fm2(cf18, globalState), f11, |v23|, v24, globalState)|;
						globalState.f5 := |v19|;
					case DC14(cf20) =>
						var v25: array<string> := new string[2];
						v25[37] := p0;
						v25[37] := p0;
						var v26: C1 := new C1(p1, p1, |v25[37]|);
						v26 := if (false) then v26 else v26;
						var v27: array<int> := new int[7];
						r3 := v27;
						var v28: multiset<string> := multiset{p0, "ed"};
						var v29: map<array<bool>, bool> := map[cf1 := p0 !in v28];
						v29 := v29[cf1 := fm1(f10, !cf20, f10, p3, globalState)];
					case DC12(cf17) =>
						var v30 := new C1(p1, p1, f11);
						r0 := f11;
						var v31 := true;
						v31 := p2;
						var v32: array<string> := new string[25];
						v32[252] := p0;
						v32[252] := p0 + p0;
				}
				
				var v33: array<int> := new int[1];
				v33[558] := f10;
				v33[558] := f11;
			case DC2(cf2, cf3, cf4, cf5) =>
				var v34 := DC21();
				v34 := DC21();
				if (false) {
					cf5 := cf5;
					var v35: seq<int> := [f11, f11, |map[p2 := DC7(f10).cf10]|, 0x215];
					var v36 := new C2(p3, -(|v35| % -f11));
					var v37: map<char, bool> := map[p1 := v36.f17];
					var v38: multiset<bool> := multiset{if ('a' in v37) then v37['a'] else !p2, cf5};
					globalState.f5 := if ((f10 < f11) in v38) then v38[f10 < f11] else |cf3|;
					var v39: array<bool> := new bool[28] [cf5, p3, cf2, v36.f17, p3, cf5, v36.f17, cf5, p3, v36.f17, v36.f17, cf2, cf5, p2, true, p2, cf2, cf5, cf5, p3, true, p2, p3, false, p3, p2, p2, p2];
					var v40: T1 := new C1(p1, 'l', DC16(v39, f10).cf23);
					var v41: array<string> := new string[6](i0 => p0);
					var v42: seq<T1> := [v40, v40];
					v40, cf5, v41, v40.f12 := v42[-414], ((seq(644, i1  => (v40.f13))) + p0) !in ((seq(0x12e, i2  => (p0))) + [p0, p0]), v41, -f10;
					var v43: multiset<string> := multiset{fm11(v40.f12, v40.f12, p2, globalState)};
					var v44: array<D2> := new D2[23];
					var v45: set<char> := {v40.f13};
					var v46: map<bool, int> := map[cf2 := fm0(|v45|, f11, globalState)];
					var v47: C3 := new C3(v44, p1, if (p2 in v46) then v46[p2] else f11);
					var v48: map<multiset<string>, C3> := map[DC22(v43).cf28 := v47];
					v48 := v48[multiset{p0, p0, p0, p0, p0} := v47];
				} else {
					var v49: array<multiset<int>> := new multiset<int>[8];
					var v50: set<bool> := {p2, cf5};
					var v51: multiset<int> := multiset{|v50|};
					v49[693] := v51;
					v49[693] := v51;
					var v52: array<int> := new int[10];
					var v53: map<bool, int> := map[p3 := f11];
					v52[478] := |v53[p2 := -f11]|;
					var v54: map<int, int> := map[|p0| := f10];
					v52[478] := (f11 - |v54|) * f11;
					cf2 := cf2;
					cf2 := fm1(f11, p3, f10, cf5, globalState);
					var v55: set<array<int>> := {v52, v52, v52};
					var v56 := DC13(true, -f11 != |v55|);
					v56 := v56;
				}
				
				globalState.f5 := |fm10(-926 % f11, globalState)|;
				var v57: seq<bool> := [cf2];
				var v58: map<seq<bool>, set<bool>> := map[v57 := {p3}];
				var v59: set<bool> := {p2, p3};
				cf5 := ((if (v57 in v58) then v58[v57] else v59) - {p3, p3}) > v59;
			case DC0(cf0) =>
				var v60: multiset<set<int>> := multiset{v1};
				cf0 := v60 > v60;
				var v61: array<int> := new int[8] [f11, f11, -f11, |v1|, f10, f11, f10, |map[f11 := f11]|];
				var v62: array<bool> := new bool[16] [cf0, false, p3, p3, p2, p2, p2, p3, p2, fm1(f10, p3, f10, p2, globalState), p3, p2, cf0, cf0, p3, cf0];
				var v63: map<array<int>, array<bool>> := map[v61 := v62];
				v62[731] := p2;
				v63, v62[731] := v63[v61 := v62] + v63, p2;
				var v64: multiset<string> := multiset{"hs", "yn"};
				var v65 := DC22(v64);
				match (v65) {
					case DC23(cf29) =>
						var v66: map<bool, array<int>> := map[p2 := v61];
						var v67: T1 := new C1(p1, p1, cf29);
						var v68: map<D7, T1> := map[DC20(map[v62[731] := v61]).(cf27 := v66) := v67];
						var v69: map<map<D7, T1>, bool> := map[v68 := !p3];
						var v70: map<int, map<map<D7, T1>, bool>> := map[cf29 := v69];
						var v71 := DC7(v67.f12);
						var v72: map<bool, bool> := map[v62[731] := v62[731]];
						var v73: array<D2> := new D2[15] [v71, v71.(cf10 := cf29), v71, DC7(|v72|), v71, v71, DC7(f10), v71, v71, v71, v71, v71, DC7(0x21b), v71, v71];
						var v74: C3 := new C3(v73, p1, -596);
						var v75: seq<C3> := [v74];
						v70 := v70[cf29 / |v75| := v69];
						var v76 := DC20(v66);
						var v77: seq<D7> := [v76];
						v77 := ((v77 + v77) + v77)[v67.f12 := v76];
						var v78: array<D4> := new D4[29](i3 => DC11({p3}));
						var v79 := new C3(v74.f14, p1, |multiset{v78, v78}|);
						var v80 := 'u';
						var v81: map<bool, char> := map[p2 := v67.f13];
						v80 := if ((if (v62[731]) then p2 else p2) in v81) then v81[if (v62[731]) then p2 else p2] else p1;
					case DC24(cf30, cf31) =>
						var v82 := DC7(f11);
						var v83: array<D2> := new D2[13] [v82, v82, v82, v82, v82, v82, v82, v82, v82, v82, v82, v82, v82];
						var v84 := DC25(p1);
						var v85 := new C3(v83, v84.cf32, |{cf0}|);
						v62[731] := true;
						var v86: map<bool, bool> := map[true := !!p3];
						var v87: seq<int> := [if (!v62[731]) then |v86| else f11];
						globalState.f5 := v87[--f10];
						var v88 := DC18(|map[!cf0 := p2]|);
						v62[731], f11, v62[731], v62[731] := v62[731], v88.cf25, fm1(f10, cf0, -(f11 % 0xfa), p2, globalState), p2;
					case DC22(cf28) =>
						var v89: map<int, int> := map[f10 := f11];
						var v90: seq<int> := [0x35a];
						v89 := v89[f11 := v90[|p0|]];
						var v91: T1 := new C1(p0[f11], p1, f10);
						var v92: array<D2> := new D2[12](i4 => DC7(|{v91.f12, f10, v91.f12, |p0|}|));
						v91 := new C3(v92, v91.f13, v90[f10]);
						var v93 := 'n';
						v93 := 'a';
						v1 := v1;
				}
				
				v62 := v62;
			case DC3(cf6) =>
				var v94: array<bool> := new bool[21];
				var v95 := DC1(v94);
				var v97: seq<int> := [|[|(map v96 : char | v96 in p0 :: (v96) := (p2))|]|, 257];
				var v98: map<array<bool>, int> := map[v95.cf1 := |(v97 + [f11])|];
				v94[125] := p3;
				var v99: multiset<string> := multiset{p0, p0 + (seq(0x1bf, i5  => (p1)))};
				var v100: multiset<int> := multiset{f11};
				var v101: set<multiset<int>> := {multiset(v97), v100, v100};
				v98, globalState.f5, v94[125], v99 := v98, f11, v101 != (v101 * v101), v99;
				var v102 := 'u';
				v102 := v102;
				var v103: map<string, char> := map[p0 := p1];
				var v104: map<map<string, char>, int> := map[v103 := f11];
				v104 := v104[v103 := f10];
				v94[125] := p2;
		}
		
		var v105 := new C1(p1, p1, 554);
		var v106: multiset<int> := multiset{-0x386};
		var v107: map<int, multiset<int>> := map[0x92 := v106];
		if (v106[f11 := f10] <= (if (p3) then v106 else if (f11 in v107) then v107[f11] else v106)) {
			var v108: map<bool, int> := map[p2 := f10];
			var v109: map<bool, int> := map[p2 := |v108|];
			var v110: T0 := new C2(p3, if (p2 in v109) then v109[p2] else f11);
			var v111: array<int> := new int[9](i6 => i6 - v110.f12);
			var v112: map<int, array<int>> := map[v110.f12 := v111];
			var v113: seq<array<int>> := [v111];
			f11, globalState.f5, v110, v110.f12 := |((v106 * v106) - v106[f10 := f10])|, |(map[f10 := v111] + v112[f11 := v111])[f11 := v113[f11]]|, v110, f11;
			var v114: array<array<int>> := new array<int>[25] [v111, v111, v111, v111, v111, v111, v111, v111, v111, v111, v111, v111, v111, v111, v111, v111, v111, v111, v111, v111, v111, v111, v111, v111, v111];
			v114[501] := v111;
			var v115: seq<bool> := [p3, !true];
			var v116: array<bool> := new bool[17](i7 => p3);
			var v117: map<int, array<bool>> := map[v110.f12 := v116];
			v114[501] := new int[5] [|v115|, f11, |v117[f11 := v116]|, v110.f12, f10];
			v110.f12 := |(p0 + (p0 + p0))|;
			var v118: array<char> := new char[10](i8 => v105.f15);
			v118[310] := v105.f15;
			v118[310] := p1;
			v116[824] := !p2;
			var v119: map<bool, bool> := map[p3 := p2];
			v116[824] := if (p3) then if (fm1(v105.fm4(globalState), p3, f11, !false, globalState) in v119) then v119[fm1(v105.fm4(globalState), p3, f11, !false, globalState)] else p2 else p3;
		} else {
			var v120: array<int> := new int[1](i9 => i9 % f11);
			var v121 := DC4(v120);
			var v122: set<D1> := {DC4(v120), v121, v121};
			var v123: map<char, set<D1>> := map[v105.f15 := v122];
			var v124: map<D7, int> := map[DC20(map[p3 := v120]) := f11];
			var v125: seq<int> := [f10];
			var v126: map<bool, seq<int>> := map[p2 := v125];
			var v127: seq<bool> := [{v121} == (if (v105.f15 in v123) then v123[v105.f15] else v122), -(if (DC20(map[p3 := v120]) in v124) then v124[DC20(map[p3 := v120])] else |v126|) == f11, p3, f10 > f11];
			v127 := fm9(p2, globalState);
			var v128 := false;
			v128 := p2;
			var v129 := DC5(p2);
			var v130 := DC6(v129);
			match (v130) {
				case DC5(cf8) =>
					var v131: map<bool, string> := map[!v128 := p0];
					var v132: seq<string> := [if (p3 in v131) then v131[p3] else p0];
					var v133: map<int, seq<string>> := map[f11 := v132];
					var v134: map<int, seq<string>> := map[f11 := if (f11 in v133) then v133[f11] else v132];
					v132 := if ((f10 + |v127|) in v134) then v134[f10 + |v127|] else v132;
					cf8 := p2;
					var v136 := new C0(map v135 : int | (-0x8a <= v135) && (v135 < 0xe3) :: (v135 + f11) := (cf8), f11);
					var v137 := DC5(v128);
					var v138 := new C2(v137.cf8, f10);
				case DC4(cf7) =>
					cf7[549] := f11;
					cf7[549] := f11;
					var v140: map<int, char> := map[cf7[549] := v105.f15];
					var v141: map<int, bool> := map[f10 := true];
					var v142 := new C0((map v139 : int | v139 in v140 :: (v139 + f11) := (p2)) + v141, f10);
					var v143 := "puri";
					v143 := ("s" + (seq(651, i10  => (v105.f15)))) + v143;
					var v144 := DC10(v143);
					v143 := v144.cf15 + (seq(-971, i11  => (v105.f15)));
				case DC6(cf9) =>
					v128 := p3;
					var v145 := DC14(p2);
					var v146: array<bool> := new bool[11] [p3, v128, [f10] == v125, v128, false, v128, v128, p2, p3, !fm1(f10, !v128, f10, v128, globalState), v145.cf20];
					var v147: set<bool> := {p3, !p3};
					v146[69] := fm1(|v147|, p3, f10, p2, globalState);
					v146[69] := true;
					var v148: array<seq<int>> := new seq<int>[3](i12 => [|p0|] + [f11]);
					v148[38] := v125;
					v148[38] := v125[f11 := f10];
					v128 := v128;
			}
			
			v128 := !fm1(f11, v128, f10, p3 <==> p3, globalState);
			var v149: map<bool, bool> := map[p2 := v128];
			globalState.f5 := if (fm1(f11, v128, f10, v128, globalState)) then |v149| else f10;
		}
		
		var i13 := 0;
		while (!(!p3 ==> true))
			decreases 100 - i13
		{
			if (i13 >= 100) {
				break;
			}
			
			i13 := i13 + 1;
			var v150 := false;
			var v151: array<bool> := new bool[22](i14 => v150);
			var v152: set<array<bool>> := {v151};
			v150 := (v152 - v152) > v152;
			var v154: multiset<bool> := multiset{false, false};
			var v155: map<int, bool> := map[|v154| := true];
			var v156 := DC27(v105, f11, f11, f10, map v153 : map<int, bool> | v153 in multiset([v155]) :: (v153) := (v105.f15));
			v150 := fm1(f10, p3, |({f10, f11, v156.cf34, f10} + v1)|, !!(if (fm1(-f11, p3, f10, false, globalState)) then !p3 else v150), globalState);
			var v157: seq<int> := [f11];
			v157 := v157;
			if (v150) {
				var v158: array<seq<int>> := new seq<int>[5] [v157, v157, v157, v157, v157];
				v158[474] := [v157[|p0|], f10];
				var v159: array<D2> := new D2[12](i15 => DC7(f10));
				var v160: T1 := new C3(v159, p0[f10], f10);
				v158[474], globalState.f5, r0, v160, v150 := [-0x3cd, v160.f12, -0xb9] + v157, -939, fm2(false, globalState) * v160.f12, v160, fm1(599, v150, f11, v150, globalState);
				var v161: map<bool, int> := map[v150 := f10];
				var v162: seq<bool> := [true];
				v150, v150, v150 := v160.f12 <= |(v161 + v161)|, fm1(|v155|, false in v162, if (v150) then f10 else -|p0|, p3, globalState), v160.f12 == f10;
				v151[155] := !p2;
				v151[155] := p3;
				var v163: seq<map<bool, int>> := [map[true := f10]];
				var v164: T0 := new C2(p3, -fm0(|v163[f10]|, |v161|, globalState));
				v164 := v164;
				v151[155] := p3;
			} else {
				v150 := p3;
				var v165: array<D5> := new D5[5](i16 => DC13(!v150, !p2));
				var v166 := DC13(!p3, p3);
				v165[792] := v166;
				v165[792] := v166;
				var v167 := new C1(v105.f15, v105.f15, f11);
				var v168 := new C1(v167.f15, v167.f15, f10);
				var v169: array<int> := new int[5];
				var v170: set<bool> := {p2, true, v150};
				v169[924] := |v170|;
				v169[924] := f10;
			}
			
		}
		if (DC8(p3, f11, p2).cf11) {
			var v171: map<int, bool> := map[f10 := p2];
			var v172: C0 := new C0(v171, f10);
			var v173: map<char, C0> := map[v105.f15 := v172];
			v173 := v173;
			if (|"eubmv"| > f10) {
				var v174, v175, v176 := v105.m5(globalState);
				var v177: set<bool> := {v175};
				var v178 := DC0(true);
				var v179: array<set<bool>> := new set<bool>[28] [v177, v177, {p2}, v177 + v177, {true}, {v175}, v177, {p3, p3} - v177, v177, {v178.cf0, p3, !v175, fm1(0x332, p3, v174, p2, globalState)}, v177, v177, {p2, p2}, {p3}, v177, v177, v177, v177, {v175, v175, p3}, v177 - v177, v177, v177 + {p3}, v177, v177, v177 - v177, v177, v177 * {p3}, v177];
				var v180 := DC11(v177);
				v179[504] := v180.cf16 * v180.cf16;
				v175, v179[504] := false, v177 - (v177 - v177);
				var v181: array<string> := new string[8](i17 => "hhjkitey" + p0);
				v181[207] := p0 + p0;
				var v182: seq<bool> := [fm1(v174, false, f10, p3, globalState)];
				var v183 := DC24(v177, v182);
				var v184: multiset<D8> := multiset{v183.(cf30 := v179[504])};
				v181[207] := if (v184 >= v184) then p0 else p0;
				var v185: array<int> := new int[4](i18 => i18 + |v106|);
				var v186: array<D2> := new D2[2](i19 => DC7(|v1|));
				var v187: T1 := new C3(v186, v105.f15, f10);
				var v188: seq<T1> := [v187];
				var v189: map<array<int>, seq<T1>> := map[v185 := v188];
				v189 := v189[if (false) then v185 else v185 := v188 + v188];
				globalState.f5 := f10 / f10;
			} else {
				var v190 := false;
				v190 := p3;
				var v191: map<char, multiset<int>> := map[v105.f15 := multiset(v105.fm5(p2, !p2, false, !p3, globalState))];
				var v192: seq<int> := [f10];
				var v193 := DC0(p2);
				var v194: seq<bool> := [v190];
				var v196: array<bool> := new bool[20] [false, (if (v105.f15 in v191) then v191[v105.f15] else v106)[f10 := -600] >= fm21(f10, f10, 'h', f10, globalState), p2, p3, p2, p2, !p2, |v192| <= (if (f11 in v106) then v106[f11] else f11), v193.cf0, p3, fm1(|v192|, fm1(|{f10}|, p3, |"wma"|, v194[f10], globalState), --f10, v190, globalState), true, true, (seq(0x3d, i20  => ('k'))) < p0, p2, !(v1 >= (set v195 : int | (0x264 <= v195) && (v195 < -261) :: (v195 / f10))), v190, v190, f10 >= f11, false];
				v196[763] := p2;
				var v197: set<bool> := {false};
				var v198 := DC11(v197);
				var v199: seq<C0> := [v172];
				var v200: array<int> := new int[16];
				var v201: seq<array<int>> := [v200, v200, v200];
				globalState.f5, globalState.f5, v196[763], v198, r3 := |v1| % |v199|, f10 * f10, v106 > (v106 * v106), v198, v201[f11];
				var v202 := DC5(p3);
				v190 := v202.cf8;
				v172 := new C0(map v203 : int | (361 <= v203) && (v203 < 634) :: (v203 / f10) := (true), f11);
				var v204: map<seq<int>, seq<bool>> := map[v192 + v192 := v194];
				v194 := if (v192 in v204) then v204[v192] else [v190];
			}
			
			var v205: map<bool, bool> := map[!p3 := p2 && p3];
			v205 := v205[p3 := p3];
			var v206: seq<int> := [f11];
			globalState.f5 := v206[|v106|];
			var v207: array<bool> := new bool[21];
			v207[232] := if (p3) then p2 else p3;
			v207[232] := p3;
		} else {
			var v208: map<bool, bool> := map[p2 := p3];
			v208 := v208[p3 := fm1(f10, p2, f11, !p2, globalState)];
			var v209: array<set<bool>> := new set<bool>[22](i21 => {p3, p2, p2, !true, p2});
			v209 := v209;
			var v210 := DC7(f11);
			var v211: array<D2> := new D2[3] [v210, DC7(f11), v210];
			var v212: C3 := new C3(v211, v105.f15, f11);
			var v213: map<C3, int> := map[v212 := f11];
			var v214: multiset<map<C3, int>> := multiset{v213, v213};
			var v215: multiset<bool> := multiset{v214 !! v214};
			v215 := v215;
			var v216: array<int> := new int[10](i22 => i22 - f10);
			v216[546] := -f10;
			v216[546] := |p0|;
			var v217: array<bool> := new bool[25];
			var v218 := DC1(v217);
			v218 := v218;
		}
		
		v1 := v1;
		r0 := -0x38b;
		var v219: multiset<bool> := multiset{p2};
		var v220: array<bool> := new bool[5] [p0 < p0, f11 == |(seq(0x12e, i23  => (v105.f15)))|, !p3, !fm1(f10, p2, |v219|, p2, globalState), p3];
		r1 := v220;
		var v221: array<D0> := new D0[3](i24 => v2);
		var v222: set<array<D0>> := {v221};
		r2 := (v222 * v222) + v222;
		var v223: array<int> := new int[1];
		r3 := v223;
	}
}

method Main() {
	var v0 := 0xac;
	var v1: seq<int> := [v0];
	var v2: map<int, seq<int>> := map[v0 := v1];
	var v3: multiset<int> := multiset{v0};
	var v4 := 's';
	var v5 := "o";
	var v6: array<char> := new char[11] [v4, v4, v4, v4, v4, v4, v4, v4, v4, v5[v0], v4];
	var globalState := new GlobalState(multiset(if (|v1| in v2) then v2[|v1|] else [0x8a]) - v3, true, false, true, 841, 0x152, true, false, -0xc0, v6);
	var v7: map<int, int> := map[v0 := 447];
	var v8 := true;
	var v9, v10, v11 := m0(if (0x31 in v7) then v7[0x31] else v0, v0, v0, v8, globalState);
	v8 := v8;
	for i0 := v9 to v0 {
		var v12: map<int, string> := map[-0x10d - fm0(v9, v9, globalState) := "ytxvk"];
		v5 := if (-v9 in v12) then v12[-v9] else v11;
		v8 := v8;
		var v13: seq<string> := ["xrlcvxj", v5, "enoyjic"];
		var v14: seq<bool> := [v8];
		var v15 := DC0(v8);
		var v16 := DC0(v15.cf0);
		v13, v8, v8, v0, v8 := v13 + (v13 + v13), !v8, v8 <== (|v14| > i0), i0, v15.cf0;
		var v17: array<set<int>> := new set<int>[29];
		v17 := v17;
	}
	var v18: array<int> := new int[20];
	var v19: array<array<int>> := new array<int>[14] [v18, v18, DC4(v18).cf7, v18, v18, v18, v18, v18, v18, v18, v18, v18, v18, v18];
	v19[329] := v18;
	v19[329] := v18;
	var v20: set<int> := {-0x90};
	var v21, v22, v23 := m0(v0, v9, -(if (true) then v1[-|v20|] else |v5|), fm1(-v0, v8, -267, !true, globalState), globalState);
	var v24: array<D2> := new D2[5](i1 => DC7(v0));
	var v25 := new C3(v24, 'x', -(v21 / v21));
	var v26: array<C1> := new C1[23];
	var v27: C1 := new C1(v4, v4, v9);
	var v28: map<int, bool> := map[v0 := v8];
	var v29: map<map<int, bool>, char> := map[v28 := v27.f15];
	var v30 := DC27(v27, |v3|, |v5|, -208, v29);
	v26[747] := v30.cf33;
	v26[747] := new C1(v4, v27.f15, |v11| - v0);
	v18[190] := v0;
	v18[190] := -v25.fm4(globalState);
	var v31: array<C3> := new C3[24];
	v31[244] := v25;
	v31[244] := if (v8) then v25 else v25;
	var v32, v33, v34, v35 := v27.m2(globalState);
	var v36 := DC14(v8);
	if (v36.cf20) {
		var v37: array<set<T2>> := new set<T2>[12];
		var v38: T0 := new C2(v32, v25.fm7(globalState));
		var v39: map<array<set<T2>>, T0> := map[v37 := v38];
		v39 := v39[v37 := v38];
		var v40: array<bool> := new bool[5] [!v32, !v8, v32, |multiset{v35}| < v21, v8];
		v40[953] := v8;
		var v41 := DC8(v18[190] >= v9, if (!v32) then v1[v21] else |v3|, [v27.f15, 'n'] <= v23);
		var v42: map<D1, array<int>> := map[DC4(v19[329]) := v33];
		var v43: map<bool, set<int>> := map[v32 := {-0x370, v21}];
		var v44: array<set<int>> := new set<int>[12] [{v35} - v20, fm13(false, v8, globalState), v20, {-v21} * v20, {v35}, v20, {|v42|} - (if (v32 in v43) then v43[v32] else v20), {v38.f12}, v20 + v20, v20 - v20, {v18[190]}, v20];
		v44[177] := {v18[190], -v0};
		v40[953], v9, v41, v44[177] := v8, v25.fm7(globalState), v41.(cf11 := false, cf13 := true), v20;
		var v45: map<bool, bool> := map[v8 := true];
		var v46: multiset<bool> := multiset{v8};
		var v47: seq<multiset<bool>> := [v46];
		var v48 := DC12([multiset{v32, !v32}]);
		var v49: seq<D5> := [DC12(v47), DC12(v47), v48, v48, DC12(v47)];
		var v50 := DC15(v27.fm5(if (!false in v45) then v45[!false] else v40[953], v8, v32, fm1(v35, true, v1[|v49|], v32, globalState), globalState));
		v50 := DC15(v1).(cf21 := v1);
		var v51: array<D3> := new D3[24](i2 => if (v40[953]) then DC10(v11) else DC10(v23));
		v51 := v51;
		var v52: set<bool> := {true, v40[953], v32};
		if (fm1(v18[190], fm1(v9, if (v32 in v45) then v45[v32] else v32, v38.f12, fm1(0x7d, v40[953], |v52|, v32, globalState), globalState), v25.fm7(globalState), v32, globalState)) {
			v40[953] := (fm21(v0, v21, v4, -v35, globalState) >= multiset(v1)) <== v8;
			v8 := v32;
			var v53, v54, v55, v56 := v25.m2(globalState);
			var v57: map<T0, bool> := map[v38 := v8];
			var v58 := new C2(v8, |(v57 + v57)|);
			v40 := v40;
		} else {
			globalState.f5 := -v18[190];
			var v59: map<int, string> := map[v25.fm7(globalState) := v5];
			var v60: set<C1> := {v27};
			v59 := v59[|v60| * 627 := v11];
			v32 := v40[953] ==> v40[953];
			v18[190] := -(v0 / 241) % fm0(if (v18[190] in v7) then v7[v18[190]] else v0, 0x292, globalState);
			v23 := seq(-635, i3  => (v11[v9]));
		}
		
	} else {
		var v61 := new C3(v24, 'j', v0);
		var v62: seq<bool> := [v32];
		v62 := v62[v21 := if (true) then v32 else fm1(v0, v8, v0, v32, globalState)];
		var v63: map<array<int>, int> := map[v33 := v21];
		if (v63[v18 := v21] != v63) {
			var v64: map<bool, array<C1>> := map[v32 := v26];
			var v65: array<set<int>> := new set<int>[3] [v20, v20, {|fm19(v0, globalState)|, -v18[190], v21}];
			v65[588] := v20;
			var v66 := DC9(v0);
			v64, v65[588], v66, v23 := v64[v3 >= v3 := v26], {|v61.fm5(v32, v32, v32, v8, globalState)| + v18[190], v18[190]}, v66, v23;
			var v67 := new C2(v8, v9);
			var v68: array<T2> := new T2[11];
			var v69: T2 := new C3(v61.f14, 'r', 0x10c);
			v68[729] := v69;
			v8, v68[729] := v62[v69.f12] <==> v32, if (v8) then v69 else v69;
			var v70, v71, v72, v73 := v25.m2(globalState);
			v23 := v5;
		} else {
			var v74: C4 := new C4(v18[190], v9);
			var v75: seq<C4> := [v74, v74, v74];
			var v76: map<int, seq<C4>> := map[|v22| := v75];
			v21 := |(if (v18[190] in v76) then v76[v18[190]] else v75)| / v35;
			var v77: C0 := new C0(v28, v35);
			var v78: map<C0, string> := map[v77 := v5[v9 := v27.f15]];
			v35 := |(if (false) then v78 + v78 else v78 + v78)|;
			v8 := v32;
			v8 := false;
			v19[329] := new int[6](i4 => i4 % v74.f10);
		}
		
		var v79 := DC26();
		v79 := if (!v32) then v79 else v79;
		var v80, v81, v82 := v25.m3(multiset{false, !v32}, globalState);
	}
	
	var v83: multiset<bool> := multiset{v8};
	if (fm1(|map[v32 := v83]|, v8, 0x34c, v21 < |v83|, globalState)) {
		v35 := 214;
		v28 := v28[if (v8) then v18[190] else -v18[190] := false];
		var v84: map<bool, array<C3>> := map[v8 := v31];
		var v85: map<map<int, bool>, array<C3>> := map[v28 + fm3(v32, v8, v25.fm7(globalState), globalState) := if (v8 in v84) then v84[v8] else v31];
		var v87: map<bool, char> := map[fm1(0x2e, v32, 0x3bf, v8, globalState) := v27.f15];
		var v88 := DC2(v8, map v86 : int | (0x259 <= v86) && (v86 < 587) :: (v86 / |v5|) := (v32), v87, v8);
		var v89 := DC2(v8, v88.cf3, v87[v8 := v27.f15], v8);
		var v90: map<int, array<C3>> := map[v35 := v31];
		v85 := map[v88.cf3 := if (0x2be in v90) then v90[0x2be] else v31];
		v8 := v8;
		var v91: set<bool> := {v32};
		var v92: seq<bool> := [v32];
		var v93: seq<seq<bool>> := [v92, [v8, v32, v32], v92];
		var v94 := DC24(v91, v93[v21]);
		var v95: set<D8> := {fm22(v32, globalState), v94};
		var v96: map<set<D8>, int> := map[v95 * v95 := if (v32) then |v1| else v35];
		v96, v31[244], v8 := v96 + v96, v25, true;
	} else {
		v32 := v8;
		v4 := v4;
		v4 := if (!v32) then 'm' else v11[v21];
		globalState.f5 := v0;
		globalState.f5 := 0x242;
	}
	
	var v97, v98, v99, v100 := v27.m2(globalState);
	var v101 := new C3(v25.f14, v4, |"m"|);
	var i5 := 0;
	while (true)
		decreases 100 - i5
	{
		if (i5 >= 100) {
			break;
		}
		
		i5 := i5 + 1;
		globalState.f5 := v21;
		if (v32 ==> v97) {
			var v102, v103, v104, v105 := v27.m2(globalState);
			v35 := if (v97) then v25.fm4(globalState) else v9 - v9;
			var v106: array<bool> := new bool[4](i6 => v1 < v1);
			v106[299] := v97;
			v103, v106[299] := v19[329], !v97;
			v23, v21, v18[190], v97 := v5, v35, v18[190], true;
			var v107 := new C3(v25.f14, v27.f15, -(v105 - v35));
		} else {
			v11 := if (v8) then v23 + v5 else v5 + v23;
			v18[190] := v35;
			var v108: set<char> := {v4};
			v108, v97, globalState.f5, v1 := ({v27.f15, v4} + {v4}) + v108, (v100 - v0) != v1[v9], v9, v1;
			var v109 := DC10(v5 + v5);
			v109 := v109;
			var v110, v111, v112, v113 := v27.m2(globalState);
		}
		
		if (v8) {
			var v114 := DC15(v1[v35 := v21]);
			v114 := if (!v8) then v114 else v114;
			v1 := [|v20|] + v1;
			v1 := [v100, v18[190], v9, v0] + [--v0];
			v18[190] := 399;
			v0 := v21;
		} else {
			var v115 := new C4(v9 + v35, |v22| % v21);
			var v116, v117, v118 := v101.m4(v9, v8, v5 + (seq(-0x29d, i7  => (v27.f15))), globalState);
			var v119: array<bool> := new bool[16];
			v119[8] := |v20| != v115.f10;
			var v120: map<int, set<int>> := map[v0 := {v21}];
			v119[8] := (if (-|v23| in v120) then v120[-|v23|] else v20) >= v20;
			var v121, v122, v123 := v101.m3(v83, globalState);
			v119[8] := false;
		}
		
		v0 := v18[190];
	}
	v35, v32 := -v9, fm1(|v10|, v97, v100, v97, globalState);
}