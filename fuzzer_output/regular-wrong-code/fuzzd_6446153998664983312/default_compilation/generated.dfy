datatype D0 = DC1 | DC2(cf1: bool, cf2: int) | DC0(cf0: bool)
datatype D1 = DC4(cf4: bool, cf5: map<bool, int>, cf6: bool) | DC3(cf3: map<bool, int>) | DC5(cf7: D1)
datatype D2 = DC7 | DC8(cf9: bool, cf10: bool) | DC6(cf8: set<bool>)
datatype D3 = DC9(cf11: multiset<array<bool>>)
datatype D4 = DC11 | DC12(cf13: bool, cf14: T0, cf15: bool, cf16: bool) | DC13(cf17: int) | DC10(cf12: array<int>)
datatype D5 = DC15(cf19: int, cf20: bool) | DC14(cf18: seq<int>)
class GlobalState {
	const f0 : multiset<seq<bool>>
	const f1 : array<bool>
	const f2 : int
	const f3 : bool
	const f4 : int
	var f5 : bool
	const f6 : map<bool, int>
	const f7 : map<int, bool>
	var f8 : bool
	const f9 : bool
	constructor (f0 : multiset<seq<bool>>, f1 : array<bool>, f2 : int, f3 : bool, f4 : int, f5 : bool, f6 : map<bool, int>, f7 : map<int, bool>, f8 : bool, f9 : bool) {
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

function fm0(p0: int, p1: map<seq<bool>, bool>, globalState: GlobalState): int {
	|({!false} + {false})|
}
function fm1(p0: int, p1: bool, p2: int, p3: bool, globalState: GlobalState): multiset<bool> {
	multiset{!false} * (if (false) then multiset([false]) else multiset([!true]))
}
function fm2(p0: bool, p1: seq<bool>, globalState: GlobalState): bool {
	!true
}
function fm3(globalState: GlobalState): set<bool> {
	{true, false, false} - ({true} * {false, !false})
}
function fm4(p0: char, p1: string, p2: bool, globalState: GlobalState): seq<int> {
	(if (true) then DC14([DC13(|[|{0x282}|, |map[true := false]|, |(seq(0x150, i0  => ('v')))|]|).cf17]) else DC14([0x2a3, 827, |map[718 := -0x195]|])).cf18
}
function fm5(p0: char, globalState: GlobalState): string {
	match DC8(true, false) {
		case DC7() => "ammm" + (seq(832, i0  => ('o')))
		case DC8(cf9, cf10) => "widm"
		case DC6(cf8) => if (true) then "ult" else "krthaq"
	}
}
function fm8(globalState: GlobalState): multiset<int> {
	match if (!false) then DC14([268]) else DC14([-0x10b]) {
		case DC15(cf19, cf20) => if (cf20) then multiset([-|map[cf20 := cf20]|]) else multiset{|map['n' := "isbarghd"]|}
		case DC14(cf18) => multiset{|[!false]|}
	}
}
function fm9(p0: int, p1: int, p2: int, globalState: GlobalState): set<char> {
	(set v0 : char | v0 in (seq(0x3b7, i0  => ('y'))) :: (v0)) * {'f', 'm', 'i', 'f', 'x'}
}
function fm10(p0: int, p1: bool, p2: int, p3: int, globalState: GlobalState): seq<bool> {
	[false, -0x182 > 0x1e6, !({'s', 'q'} >= (set v0 : char | v0 in ['v'] :: (v0))), !(if (false) then false else !!true)]
}
function fm11(globalState: GlobalState): map<seq<bool>, bool> {
	if ({true, !false} >= {true, !false}) then map[[false] := !false] else map[[false] := false] + map[[false, true] := true]
}
function fm12(p0: map<int, int>, p1: set<int>, p2: int, p3: bool, globalState: GlobalState): map<int, string> {
	map[-0x1d0 := "creunkjsj"] + (map[|multiset{true}| := "ft"] + map[|"lvvvveiw"| := seq(0x110, i0  => ('g'))])
}
method m0(p0: bool, p1: seq<bool>, globalState: GlobalState) returns (r0: string) {
	var v0: array<bool> := new bool[26];
	forall i0 | 0 <= i0 < v0.Length {
		v0[i0] := false;
	}
	var v1 := 0x22b;
	var v2 := DC0(p0);
	var v3: C0 := new C0(v1, p0, v2);
	var v4 := 'p';
	var v5: set<char> := {v4};
	var v6: set<set<char>> := {v5, v5, v5, v5, fm9(v3.f12, v1, |p1|, globalState)};
	for i1 := if (p0) then v1 else v1 to |map[v3 := |v6|]| {
		var v7 := "lotbhuwwl";
		var v8: set<int> := {|v7|};
		var v9: seq<seq<bool>> := [p1];
		var v10: map<int, int> := map[v3.f12 := i1];
		var v11: multiset<array<bool>> := multiset{v0};
		var v12 := DC9(v11);
		var v13: array<seq<bool>> := new seq<bool>[22] [p1, p1[i1 := !!p0], p1, p1, p1, p1, v9[v1], p1 + ([!true])[v3.f12 := p0], p1 + p1, fm10(880, p0, v3.f12, i1, globalState), p1, p1, p1, if (p0) then p1 else [p0], fm10(v3.f12, p0, if (|v12.cf11| in v10) then v10[|v12.cf11|] else i1, i1, globalState), p1, if (p0) then p1 else p1[v1 := p0], [p0, p0, p0, p0], if (p0) then p1 else p1, [p0, p0, p0, p0, !p0], [fm2(p0, p1, globalState), p0], p1];
		v13[907] := p1;
		v8, globalState.f5, v13[907], v3, globalState.f5 := v8, p0, p1, v3, p0;
		var v14 := DC7();
		var v15: set<D2> := {v14};
		if ((v3.f12 * |v15|) > i1) {
			var v16: array<array<int>> := new array<int>[21];
			v16 := v16;
			var v17 := new C0(|v7| * i1, if (p0) then p0 else p0, v2);
			globalState.f8, v1, v2, globalState.f8 := !(p0 ==> true), i1 - (v17.f12 / -i1), v2.(cf0 := if (!p0) then p0 else p0), !(p0 ==> p0);
			var v18: seq<int> := [v3.f12];
			var v19: map<seq<int>, C0> := map[v18 := v17];
			var v20: map<map<seq<int>, C0>, map<int, int>> := map[map[v18 := v17] + v19 := v10 + v10];
			v20 := v20[v19[v18 := v17] := map[0x1e5 := v3.f12]];
			var v21: seq<C0> := [v17, v3, v17, v3];
			v21 := ([v17, v3] + v21) + v21;
		} else {
			var v22: array<int> := new int[4] [-460, v3.f12, v3.f12 % v3.f12, v3.f12];
			v22 := v22;
			var v23: multiset<int> := multiset{fm0(-v1, map[p1 := p0], globalState)};
			var v24 := DC2(p0, |v23|);
			var v25: map<seq<bool>, bool> := map[v13[907][9 := p0] := p0];
			var v26: seq<map<seq<bool>, bool>> := [map[p1 := p0], v25, v25];
			v1, v1, v4, v24 := -v3.f12, if (v3.f12 in v23) then v23[v3.f12] else fm0(v3.f12, v26[i1], globalState), v4, v24;
			globalState.f5 := true;
			var v27: seq<int> := [fm0(|v7|, fm11(globalState), globalState), |v13[907]|];
			var v28: map<char, int> := map[v4 := v3.f12];
			globalState.f5 := v27[0x23b] in (fm12(v10, v8, |v28|, p0, globalState))[v3.f12 := v7];
			var v29: T0 := new C0(v1, p0, v2);
			var v30: map<int, T0> := map[|p1| := v29];
			var v31: set<map<int, T0>> := {v30};
			v31 := v31;
		}
		
		var v32: set<bool> := {p0};
		var v33: map<int, set<bool>> := map[459 := v32];
		var v34: multiset<int> := multiset{|(v33 + v33)|, v3.f12 + v1, v1, 0x361, i1 + v3.f12};
		v34 := v34;
		v3 := v3;
	}
	var v35: map<bool, multiset<bool>> := map[p0 := multiset{p0}];
	var v36: map<bool, bool> := map[p0 := p0];
	globalState.f8 := v3.fm7(v2, v35, if (if (p0 in v36) then v36[p0] else p0) then v1 else 0x261, globalState);
	var v37: T0 := new C0(--0x236, p0, v2);
	var v38: map<int, T0> := map[v3.f12 := v37];
	var v39: multiset<T0> := multiset{v37, if (v3.f12 in v38) then v38[v3.f12] else v37};
	var v40: map<multiset<T0>, bool> := map[v39 := !(v1 != 0x39a)];
	globalState.f8 := if (v39 in v40) then v40[v39] else v37.f10;
	var i2 := 0;
	while (DC8(v37.f10, p0).cf10)
		decreases 100 - i2
	{
		if (i2 >= 100) {
			break;
		}
		
		i2 := i2 + 1;
		v1, v1, globalState.f8 := 107, v1, false;
		if (v3.f12 == v3.f12) {
			v1 := -v1;
			var v41: seq<int> := [v3.f12, v1];
			var v42 := DC8(v37.f10, v37.f10);
			var v43: multiset<bool> := multiset{v37.f10};
			var v44: seq<multiset<bool>> := [v43];
			var v45: array<int> := new int[18](i3 => i3 * v1);
			var v46: map<bool, array<int>> := map[p1[v3.f12] := v45];
			v41, globalState.f8, v1, v42, v37 := v41 + (if (v37.f10) then v41 else [v3.f12]), p0 <== (v43 !in v44), 854 * (v3.f12 * |v46|), v42, if (v3.f12 in v38) then v38[v3.f12] else v37;
			v45[287] := v3.f12;
			var v47: multiset<array<bool>> := multiset{v0};
			var v48: set<seq<int>> := {v41};
			var v49: seq<C0> := [v3];
			var v50: map<bool, seq<C0>> := map[v48 >= v48 := v49 + v49];
			v45[287], v1, v1, v1, globalState.f8 := |(v47 + v47)| * v3.f12, v3.f12, v3.f12, |v50|, v37.f10 || v37.f10;
			v37 := v37;
			var v51: array<D1> := new D1[1](i4 => DC4(v37.f10, map[p0 := v3.f12], !v37.f10));
			var v52: map<bool, int> := map[v37.f10 := v3.f12];
			var v53 := DC4(v37.f10, v52, true);
			v51[111] := v53;
			v51[111] := v53;
		} else {
			var v54 := DC1();
			v54 := v54;
			var v55 := DC8(v37.f10, v37.f10);
			var v56: multiset<bool> := multiset{v55.cf9, v37.f10, v37.f10};
			globalState.f8 := if (v37.f10) then v37.f10 else v3.fm7(DC0(false), map[v37.f10 := v56], v3.f12, globalState);
			v37 := v37;
			v37 := v37;
			var v57: array<C0> := new C0[9];
			var v58: map<int, C0> := map[v3.f12 := v3];
			v57[387] := if (v37.f10) then v3 else if (v3.f12 in v58) then v58[v3.f12] else v3;
			v57[387] := v3;
		}
		
		var v59 := DC2(v37.f10, v3.f12);
		var v60: map<seq<bool>, bool> := map[p1 := v37.f10];
		var v61 := DC12(v37.f10, v37, true, v37.f10);
		var v62: map<C0, D4> := map[v3 := v61];
		var v63: seq<int> := [-0x1cf];
		var v64: array<int> := new int[19] [-181, v3.f12, 0x292, v59.cf2, v1, v3.f12, fm0(v3.f12, v60, globalState), v3.f12, |v62|, -0x19d, v1, v1, v1, v3.f12, 528, v1, v3.f12, v1, |v63|];
		var v65 := DC10(v64);
		var v66: map<int, array<int>> := map[0x11 := v65.cf12];
		v66 := v66[if (p0) then v1 else -648 := v64];
		globalState.f8 := v37.f10 <==> (fm2(v37.f10, p1, globalState) && v37.f10);
	}
	v37 := new C0(v3.f12, p0, v37.f11);
	var v67 := "qf";
	r0 := v67 + v67;
}
trait T0 {
	const f10 : bool
	const f11 : D0
	function fm6(p0: bool, p1: int, p2: bool, globalState: GlobalState): bool 
	function fm7(p0: D0, p1: map<bool, multiset<bool>>, p2: int, globalState: GlobalState): bool 
}

class C0 extends T0 {
	const f12 : int
	constructor (f12 : int, f10 : bool, f11 : D0) {
		this.f12 := f12;
		this.f10 := f10;
		this.f11 := f11;
	}
	
	function fm6(p0: bool, p1: int, p2: bool, globalState: GlobalState): bool {
		f12 > -f12
	}
	function fm7(p0: D0, p1: map<bool, multiset<bool>>, p2: int, globalState: GlobalState): bool {
		false ==> false
	}
}

method Main() {
	var v0 := false;
	var v1: seq<bool> := [v0, v0];
	var v2 := 147;
	var v3: multiset<seq<bool>> := multiset{[true], v1, v1[-758 := false], v1[v2 := !v0]};
	var v4 := DC0(v0);
	var v5: array<bool> := new bool[11] [v0, v4.cf0, v0, v0, v0, v0, v0, v0, v0, v0, v0];
	var v6: map<bool, int> := map[v0 := v2];
	var v7 := DC4(v0, v6, v0);
	var v8: map<int, bool> := map[|[multiset{v0}]| := v0];
	var globalState := new GlobalState(v3, v5, -299, true, -0x17a, true, v7.cf5, v8, false, true);
	var v9: multiset<bool> := multiset{!v0, v0, v0};
	if (!((v9 * v9) <= v9)) {
		var v10: map<seq<bool>, bool> := map[v1 := v0];
		var v11 := "dsjek";
		v2 := fm0(v2, v10, globalState) + (|v11| + v2);
		var v12: array<int> := new int[28](i0 => i0 % v2);
		v12 := v12;
		var v13 := m0(v0, v1, globalState);
		v12[136] := v2;
		var v14: set<bool> := {v0, v0, false, v0, v0};
		var v15: set<int> := {|v14|};
		v12[136] := |v15| + v2;
		var v16: map<bool, bool> := map[v0 := v0];
		var v17: seq<int> := [if (!(if (true in v16) then v16[true] else v0)) then 446 else v12[136]];
		var v18: map<bool, seq<int>> := map[v0 := v17];
		v17 := (v17 + (if (v0 in v18) then v18[v0] else v17)[v2 := v12[136]]) + [v12[136], 0x248];
	} else {
		v9 := v9 + fm1(v2, v0, v2, v0, globalState);
		globalState.f8 := DC4(v0, v6, v0).cf6;
		var v19 := DC3(v6);
		var v20 := DC5(v19);
		v5[158] := false;
		v2, v8, v20, v2, v5[158] := v2, v8, v20, v2, v0;
		v5[158] := fm2(v5[158], v1 + v1, globalState);
		globalState.f8, v5[158] := v0 ==> v0, v5[158];
	}
	
	globalState.f8 := v0;
	v2 := v2;
	for i1 := v2 to v2 {
		var v21: set<int> := {-v2, i1, i1};
		var v22 := "wgy";
		var v23 := DC1();
		var v24: map<string, D0> := map[v22 := v23];
		v2 := if (v9 != v9) then |v21| else i1 * |v24|;
		var v25: map<int, int> := map[i1 := v2];
		v9 := (multiset{v0, v0, v0, v0})[(if (v2 in v25) then v25[v2] else |v22|) <= i1 := i1];
		globalState.f8 := v0;
		if (!(v0 <==> !v0) ==> v0) {
			var v26 := 'f';
			var v27: array<char> := new char[3] [v26, if (!v0) then 'y' else v26, v26];
			v27 := v27;
			globalState.f8 := v0;
			v1 := (v1 + v1) + v1;
			var v28 := m0(v0, v1 + [!v0], globalState);
			v5[470] := true;
			var v29: array<int> := new int[2](i2 => i2 % v2);
			v29[391] := -v2;
			var v30 := DC6(fm3(globalState));
			var v31: set<bool> := {v0};
			var v32: seq<int> := [v2, -223];
			v5[470], v2, v2, v29[391] := (v30.cf8 - v31) !! (if (v0) then {v0, v0} else v31), v2 / v2, v32[i1 * i1], -(i1 + 0x145);
		} else {
			v22 := v22;
			v2 := i1;
			var v33: map<seq<bool>, bool> := map[[false] := v0];
			globalState.f5 := fm2(fm0(i1, v33, globalState) != v2, v1, globalState);
			v4 := v4;
			var v34 := 'i';
			v34 := v34;
		}
		
	}
	var v36: map<int, int> := map[fm0(v2, map v35 : seq<bool> | v35 in map[v1 := v0] :: (v35) := (v0), globalState) := v2];
	var v37: multiset<int> := multiset{0x2cd};
	var v38 := "toc";
	var v39 := 'y';
	var v40: array<int> := new int[11](i3 => i3 / v2);
	var v41: seq<array<int>> := [v40];
	v0, v2, globalState.f8, v2 := v36 == (map[0xf9 := if (v2 in v37) then v37[v2] else |v38|] + v36), |fm4(v39, fm5(v39, globalState) + v38, v0, globalState)|, v0, -|(([v40] + v41) + (v41 + [v40, v40, v40, v40]))|;
	forall i4 | 0 <= i4 < v5.Length {
		v5[i4] := v2 > (v2 * v2);
	}
	v38 := v38;
	v2 := v2;
	var v43: map<int, map<int, bool>> := map[|map[v2 := v2]| := map v42 : int | (0x2d8 <= v42) && (v42 < 0x396) :: (v42 / v2) := (v0)];
	v5[797] := 927 !in v43;
	var v45: set<seq<bool>> := {v1};
	v5[797], v39, v2, v2 := v0, if (true) then if (v0) then v39 else v39 else v39, v2 % (v2 % fm0(v2, map v44 : seq<bool> | v44 in v45 :: (v44) := (v0), globalState)), 0x259 + v2;
	var v46: map<bool, bool> := map[v5[797] := v0 <==> v5[797]];
	var v48: seq<int> := [v2];
	v46 := v46[|v37| in (map v47 : int | v47 in v48 :: (v47 - v2) := (v37)) := v5[797]];
	var v49 := new C0(v2, v0, v4);
	v2 := v49.f12 * (v49.f12 - v48[0x82]);
	for i5 := |v8| + v49.f12 to v2 {
		globalState.f8 := v0;
		var v50 := new C0(-(v49.f12 * |v6|), if (v49.f12 in v8) then v8[v49.f12] else true, v4);
		v5[797] := !((fm8(globalState) * multiset(v48)) < v37);
		var v51: seq<C0> := [v50, v50, v50];
		if (|v51| == v2) {
			var v52 := m0(v5[797], v1, globalState);
			var v53 := m0(v49.fm6(v0, v2, true, globalState), v1, globalState);
			globalState.f8 := v0;
			globalState.f5, v2 := v0, -v49.f12 / (v49.f12 - -0x3d);
			v5[797] := v0;
		} else {
			v2 := v50.f12;
			v2 := -((v2 - 0x1ca) / -v49.f12);
			v39 := v39;
			var v54 := DC6({v5[797]});
			var v55: map<D2, int> := map[v54 := v50.f12];
			v2 := if (v54 in v55) then v55[v54] else v49.f12 + v49.f12;
			var v56 := m0(v5[797], v1, globalState);
		}
		
	}
	v5[797] := (DC8(v0, v5[797]).(cf10 := v5[797])).cf9;
	var v57 := new C0(v49.f12, v0, DC0(v5[797]));
	var v58 := m0(v5[797], [v5[797]], globalState);
}