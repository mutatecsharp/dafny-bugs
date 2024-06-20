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
datatype D1 = DC2(cf2: int, cf3: int, cf4: int, cf5: int) | DC1(cf1: map<int, int>) | DC3(cf6: D1)
datatype D2 = DC5(cf8: bool, cf9: bool) | DC6(cf10: bool, cf11: bool, cf12: char) | DC4(cf7: set<C0>) | DC7(cf13: D2)
datatype D3 = DC9(cf15: bool) | DC10(cf16: D2, cf17: char, cf18: int) | DC8(cf14: set<char>)
datatype D4 = DC12(cf20: bool, cf21: bool, cf22: bool, cf23: seq<string>) | DC11(cf19: array<D3>)
class GlobalState {
	const f0 : bool
	const f1 : bool
	const f2 : bool
	const f3 : bool
	const f4 : map<int, seq<bool>>
	const f5 : int
	var f6 : int
	const f7 : bool
	const f8 : string
	const f9 : string
	var f10 : int
	const f11 : array<bool>
	const f12 : bool
	const f13 : int
	constructor (f0 : bool, f1 : bool, f2 : bool, f3 : bool, f4 : map<int, seq<bool>>, f5 : int, f6 : int, f7 : bool, f8 : string, f9 : string, f10 : int, f11 : array<bool>, f12 : bool, f13 : int) {
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
	}
	
}

class C0 {
	const f14 : int
	constructor (f14 : int) {
		this.f14 := f14;
	}
	
	function fm4(p0: bool, globalState: GlobalState): bool {
		DC0(true).cf0
	}
	function fm5(globalState: GlobalState): bool {
		multiset{!true, false, !false, true} !! (multiset{true, false} + multiset{true})
	}
}

function fm0(globalState: GlobalState): int {
	DC2(|map[false := 0x382]|, |"mcftfnqg"|, -0x187, |(map v0 : int | (194 <= v0) && (v0 < 0x13d) :: (v0 - -949) := ({false}))|).cf3
}
function fm1(p0: int, p1: bool, p2: bool, p3: seq<map<int, bool>>, globalState: GlobalState): bool {
	{false} !! ({true} - {true, false, false})
}
function fm2(p0: int, p1: bool, p2: bool, globalState: GlobalState): map<int, bool> {
	map[|"p"| := false] + map[|{false, true, !!false, !false}| := false]
}
function fm3(p0: int, p1: int, p2: bool, globalState: GlobalState): char {
	if (|(seq(abs(0x327), i0  => ('a')))| == |{!!!false, !!!!true, !true, false}|) then 'd' else 'g'
}
function fm6(p0: bool, p1: bool, p2: bool, p3: bool, globalState: GlobalState): map<set<int>, int> {
	map[{0x2f7} := |{true, false}|] + map[set v0 : int | v0 in map[|(seq(abs(-0x2ad), i0  => (-0x3e2)))| := true] :: (v0 - |(seq(abs(0x4f), i1  => (['m', 'w', 'k', 'b', 'g'])))|) := 930]
}
function fm7(p0: bool, p1: char, globalState: GlobalState): string {
	("j" + "iptpuc") + (seq(abs(838), i0  => ('s')))
}
function fm8(p0: bool, globalState: GlobalState): seq<bool> {
	match DC2(0xf3, |map[|map[false := "adlori"]| := DC9(true)]|, |"omdgnphda"|, |"gtoykysed"|) {
		case DC2(cf2, cf3, cf4, cf5) => [!false, true] + [!true]
		case DC1(cf1) => [false, false] + [!true, true]
		case DC3(cf6) => [true, false, false]
	}
}
method m0(p0: bool, p1: int, globalState: GlobalState) {
	globalState.f10 := p1 * (p1 * p1);
	var v0: C0 := new C0(p1);
	var v1: map<C0, int> := map[v0 := v0.f14];
	var v2: seq<bool> := [p0];
	var v3 := "iaqmfjp";
	var v4: multiset<bool> := multiset{v2[safeIndex(v0.f14, |v2|)], p0, p0};
	var v5: array<bool> := new bool[12] [v0 in v1, v2[safeIndex(|v3|, |v2|)], p0, !(p0 || p0), p0, p0, p0, p0, -|v3| != |"rpjgtvwk"|, p0, v4 > v4, v0.fm4(p0, globalState)];
	v5[safeIndex(212, v5.Length)] := true;
	v5[safeIndex(212, v5.Length)] := p0;
	var v6 := new C0(safeDivisionInt(p1, p1));
	var v7: map<bool, int> := map[p0 := |v2|];
	globalState.f10 := (if (p0 in v7) then v7[p0] else p1) - v0.f14;
	var v8: seq<C0> := [v6, v0];
	v8 := v8;
	if (!v5[safeIndex(212, v5.Length)]) {
		globalState.f6 := v6.f14;
		globalState.f10 := p1;
		v7 := v7[p0 := 0x320];
		var v9: array<int> := new int[5];
		v9[safeIndex(881, v9.Length)] := p1;
		var v10: map<bool, array<bool>> := map[p0 := v5];
		v3, globalState.f10, v9[safeIndex(881, v9.Length)], v5[safeIndex(212, v5.Length)] := v3, v6.f14, |v10|, !p0;
		var v11: map<bool, seq<map<bool, bool>>> := map[v5[safeIndex(212, v5.Length)] := seq(abs(0x337), i0  => (map[v5[safeIndex(212, v5.Length)] := v5[safeIndex(212, v5.Length)]]))];
		var v12: map<bool, bool> := map[v5[safeIndex(212, v5.Length)] := p0];
		var v13: seq<map<bool, bool>> := [v12, v12];
		v11 := v11[v5[safeIndex(212, v5.Length)] := v13];
	} else {
		var v14: map<int, int> := map[0x2c9 := |map[v5[safeIndex(212, v5.Length)] := v5[safeIndex(212, v5.Length)]]|];
		var v15 := DC1(v14);
		match (v15.(cf1 := v14)) {
			case DC2(cf2, cf3, cf4, cf5) =>
				var v16: multiset<string> := multiset{v3};
				var v17: map<int, string> := map[|v16| := v3];
				v3 := if (v0.f14 in v17) then v17[v0.f14] else seq(abs(0xa6), i1  => ('h'));
				var v18 := new C0(v0.f14);
				v5[safeIndex(212, v5.Length)] := cf4 > cf4;
				cf3 := v6.f14;
			case DC1(cf1) =>
				v5[safeIndex(212, v5.Length)] := (-v6.f14 + -|v3|) <= (if (p0 in v7) then v7[p0] else v6.f14);
				var v19: array<int> := new int[6];
				v19[safeIndex(236, v19.Length)] := v0.f14;
				globalState.f10, v19[safeIndex(236, v19.Length)], globalState.f10 := safeModuloInt(-|v3|, v6.f14), 0x4a - p1, safeModuloInt(v0.f14, p1);
				var v20 := 'c';
				v20 := v20;
				v19 := v19;
			case DC3(cf6) =>
				globalState.f10 := 0x238;
				v5[safeIndex(212, v5.Length)] := v6.fm5(globalState);
				var v21 := new C0(v0.f14);
				var v22: array<map<bool, seq<bool>>> := new map<bool, seq<bool>>[16];
				var v23: map<bool, seq<bool>> := map[!p0 := v2[safeIndex(v21.f14, |v2|) := v5[safeIndex(212, v5.Length)]]];
				v22[safeIndex(241, v22.Length)] := v23;
				var v24 := 'd';
				var v25: map<map<bool, int>, C0> := map[map[p0 := |[p0]|] := v21];
				var v26: map<char, C0> := map[v24 := if (v7 in v25) then v25[v7] else v21];
				globalState.f10, v22[safeIndex(241, v22.Length)], v26 := v6.f14, v23, v26 + map[v24 := v0];
		}
		
		var v27: map<int, bool> := map[p1 := v5[safeIndex(212, v5.Length)]];
		globalState.f10 := |v27| - -p1;
		if (!(v3 == v3)) {
			v5[safeIndex(212, v5.Length)] := false;
			var v28: seq<int> := [p1, v0.f14];
			var v29 := DC5(p0, p0);
			var v30: multiset<D2> := multiset{v29, v29};
			var v31: multiset<int> := multiset{0x238, |v30|, v6.f14};
			var v32: map<int, seq<int>> := map[p1 := v28[safeIndex(v0.f14, |v28|) := |v31|]];
			var v33: set<int> := {v0.f14, 831};
			v5[safeIndex(212, v5.Length)] := |v32| in (if (p0) then v33 else v33);
			var v34 := 'l';
			var v35: map<map<bool, bool>, char> := map[map[p0 := v5[safeIndex(212, v5.Length)]] := v34];
			globalState.f10 := v28[safeIndex(|v35| - v6.f14, |v28|)];
			var v36 := DC1(v14);
			var v37 := DC3(v36);
			v37 := if (p0) then DC3(v36) else v37;
			globalState.f10 := p1;
		} else {
			globalState.f6 := v0.f14;
			var v38 := 'g';
			v5[safeIndex(212, v5.Length)] := -|(v3 + v3)[safeIndex(v0.f14, |(v3 + v3)|) := v38]| != v0.f14;
			var v39 := DC6(!v6.fm5(globalState), v6.f14 < |v3|, fm3(v0.f14, v6.f14, p0, globalState));
			v39 := v39.(cf11 := p0 !in v2);
			var v40 := DC2(v0.f14, v0.f14, v6.f14, fm0(globalState));
			var v41 := DC3(v40);
			var v42: array<seq<int>> := new seq<int>[1];
			var v43: seq<int> := [|map[v7 := v6.f14]|, v6.f14];
			v42[safeIndex(163, v42.Length)] := v43 + v43;
			v41, v5[safeIndex(212, v5.Length)], v42[safeIndex(163, v42.Length)] := v41, p0 <==> v5[safeIndex(212, v5.Length)], v43 + v43;
			v5[safeIndex(212, v5.Length)] := false;
		}
		
		var v44 := DC0(false);
		v44 := v44;
		var v45 := DC2(-506, |v14|, v6.f14, v6.f14);
		var v46: seq<int> := [-v6.f14, v45.cf4];
		v5[safeIndex(212, v5.Length)] := (v6.f14 - -v46[safeIndex(v0.f14, |v46|)]) > |v46|;
	}
	
}
method Main() {
	var v0 := 0x1c9;
	var v1 := false;
	var v2: map<int, bool> := map[v0 := v1];
	var v3: seq<bool> := [if (|"koq"| in v2) then v2[|"koq"|] else v1];
	var v4: map<int, seq<bool>> := map[0x39c := v3];
	var v5 := "rlnyx";
	var v6: array<bool> := new bool[14];
	var globalState := new GlobalState(true, true, true, true, v4 + v4[v0 := v3], -0x8, 0x100, false, v5, v5, 133, v6, true, 0xbf);
	if (v1 <== (--v0 > fm0(globalState))) {
		var v7: array<map<bool, bool>> := new map<bool, bool>[21];
		var v8: map<array<map<bool, bool>>, bool> := map[v7 := v1];
		v8 := v8[v7 := v1];
		v0 := v0;
		globalState.f6 := -930;
		var v9: array<int> := new int[24];
		v9 := v9;
		var v10: seq<int> := [v0];
		var v11: multiset<int> := multiset{v0};
		var v12: set<multiset<int>> := {v11, v11};
		var v13: map<int, int> := map[|v12| := v0];
		var v14: set<int> := {0x140, |v5|, |v10|, v0, |v13|};
		v1 := {v0, -0x2f3, v0, v0} > (v14 + {v0, |"ryggvk"|});
	} else {
		globalState.f10 := -v0;
		var v15: map<bool, int> := map[v1 := v0];
		var v16: seq<int> := [|v2|, |v15|];
		v1 := fm1(v16[safeIndex(0x8b, |v16|)], false, v1, [v2, fm2(v0, v1, v1, globalState)] + [v2], globalState);
		var v17: map<int, int> := map[v0 := fm0(globalState)];
		v17 := v17[|(v5[safeIndex(v0, |v5|) := fm3(v0, v0, v1, globalState)] + v5)| := v0];
		globalState.f6 := safeDivisionInt(v0, v0);
		v1 := v1;
	}
	
	var v18: map<int, int> := map[v0 := v0];
	var v20: seq<map<int, bool>> := [map v19 : int | (0x398 <= v19) && (v19 < -0x322) :: (v19 + |map[v1 := v1]|) := (v1)];
	var v21: map<int, seq<map<int, bool>>> := map[|v18| := v20];
	v6[safeIndex(506, v6.Length)] := fm1(v0, true, false, if (v0 in v21) then v21[v0] else v20, globalState);
	v6[safeIndex(506, v6.Length)] := v1;
	var v22 := 'i';
	var v23: map<string, bool> := map[("c")[safeIndex(v0, |"c"|) := v22] := v6[safeIndex(506, v6.Length)]];
	v1 := (if (v5 in v23) then v23[v5] else true) <== v6[safeIndex(506, v6.Length)];
	var v24: array<int> := new int[4](i0 => safeDivisionInt(i0, v0));
	v24[safeIndex(76, v24.Length)] := v0;
	v24[safeIndex(597, v24.Length)] := |multiset{v6[safeIndex(506, v6.Length)], v1}|;
	var v25: seq<int> := [v0, v0, v0];
	v24[safeIndex(76, v24.Length)], globalState.f6, v22, v24[safeIndex(597, v24.Length)], v25 := safeModuloInt(v0, -fm0(globalState)), (v0 + |v18|) * v0, v5[safeIndex(v0, |v5|)], -v0, v25;
	for i1 := v24[safeIndex(76, v24.Length)] to v24[safeIndex(76, v24.Length)] {
		var v26: set<bool> := {v1};
		var v27: map<bool, int> := map[v1 := |v26|];
		v27 := v27[true := 0x1e2];
		m0(v6[safeIndex(506, v6.Length)], --|(v25 + v25)|, globalState);
		v3 := v3 + v3;
		var v28: map<seq<bool>, array<int>> := map[v3 := v24];
		var v29: seq<array<int>> := [v24, v24, if ([v1] in v28) then v28[[v1]] else v24];
		v24 := v29[safeIndex(v0, |v29|)];
	}
	forall i2 | 0 <= i2 < v6.Length {
		v6[i2] := v1;
	}
	var v30: map<int, char> := map[v24[safeIndex(76, v24.Length)] := v22];
	v30 := v30[v0 := v22];
	if (v1) {
		globalState.f6 := v24[safeIndex(76, v24.Length)] * v24[safeIndex(76, v24.Length)];
		if (!false || fm1(|v5|, v6[safeIndex(506, v6.Length)], v6[safeIndex(506, v6.Length)], v20, globalState)) {
			v1 := v6[safeIndex(506, v6.Length)];
			v6 := v6;
			v5 := v5;
			var v31 := new C0(v24[safeIndex(76, v24.Length)]);
			var v32: map<seq<bool>, bool> := map[v3 + v3 := v6[safeIndex(506, v6.Length)] <== v6[safeIndex(506, v6.Length)]];
			v32 := v32[v3[safeIndex(v0, |v3|) := v6[safeIndex(506, v6.Length)]] := !(v3[safeIndex(v0, |v3|)] <==> v6[safeIndex(506, v6.Length)])];
		} else {
			var v33: seq<array<bool>> := [v6];
			m0(v6[safeIndex(506, v6.Length)], |v33|, globalState);
			v24[safeIndex(76, v24.Length)] := fm0(globalState);
			v1 := v1;
			var v34: seq<string> := [v5, v5];
			var v35: map<int, seq<int>> := map[v24[safeIndex(76, v24.Length)] := [v0, v24[safeIndex(76, v24.Length)]]];
			v1, globalState.f10, v1, globalState.f6, v6[safeIndex(506, v6.Length)] := v6[safeIndex(506, v6.Length)], 0x25a, !!v6[safeIndex(506, v6.Length)], safeDivisionInt(v0, |v34[safeIndex(0x316, |v34|) := "c"]|), (v24[safeIndex(76, v24.Length)] * |v3|) in v35;
			var v36 := new C0(0x1de);
		}
		
		v6[safeIndex(506, v6.Length)] := v1;
		var v37: map<string, array<bool>> := map[v5 := v6];
		v6 := if (((seq(abs(76), i3  => (v22))) + v5)[safeIndex(v0, |((seq(abs(76), i3  => (v22))) + v5)|) := 'c'] in v37) then v37[((seq(abs(76), i3  => (v22))) + v5)[safeIndex(v0, |((seq(abs(76), i3  => (v22))) + v5)|) := 'c']] else v6;
		var v38: set<seq<int>> := {v25};
		m0({[-0x90]} > v38, v24[safeIndex(76, v24.Length)], globalState);
	} else {
		v22 := 'y';
		v5 := v5;
		if (|v5| == v0) {
			var v39: array<multiset<seq<C0>>> := new multiset<seq<C0>>[25];
			var v40: C0 := new C0(v24[safeIndex(76, v24.Length)]);
			var v41: seq<C0> := [v40];
			var v42: multiset<seq<C0>> := multiset{v41, v41};
			var v43: map<array<int>, multiset<seq<C0>>> := map[v24 := v42];
			var v44: seq<seq<C0>> := [v41];
			var v45: map<bool, bool> := map[v6[safeIndex(506, v6.Length)] := v1];
			var v46: map<int, map<bool, bool>> := map[-v0 := v45];
			v39[safeIndex(793, v39.Length)] := (if (v24 in v43) then v43[v24] else multiset(v44))[v41 := abs(|v46|)];
			v39[safeIndex(793, v39.Length)] := v42;
			m0(v6[safeIndex(506, v6.Length)], v24[safeIndex(76, v24.Length)], globalState);
			v25 := v25 + (v25 + v25);
			m0(v1, v0, globalState);
			v18 := DC1(v18).cf1;
		} else {
			v6[safeIndex(506, v6.Length)] := false;
			var v47: array<map<int, int>> := new map<int, int>[14];
			var v48: map<array<map<int, int>>, int> := map[v47 := |v25|];
			v48 := v48[v47 := |v3| - v24[safeIndex(76, v24.Length)]];
			v24[safeIndex(76, v24.Length)] := 0x36e;
			v4 := v4;
			var v49: set<int> := {v24[safeIndex(76, v24.Length)], v24[safeIndex(76, v24.Length)], v24[safeIndex(76, v24.Length)]};
			var v50: map<bool, bool> := map[v6[safeIndex(506, v6.Length)] := v6[safeIndex(506, v6.Length)]];
			var v51: map<set<int>, int> := map[v49 := |v50[v6[safeIndex(506, v6.Length)] := v6[safeIndex(506, v6.Length)]]|];
			var v52: C0 := new C0(v24[safeIndex(76, v24.Length)]);
			var v53: seq<C0> := [v52, v52];
			var v54: seq<map<set<int>, int>> := [v51, v51];
			var v55: array<map<set<int>, int>> := new map<set<int>, int>[19] [map[{v0} := v0], v51[v49 := |v53|], v51 + v51, v51, v54[safeIndex(v0, |v54|)], v51 + v51, v51, v51, fm6(v6[safeIndex(506, v6.Length)], v1, v6[safeIndex(506, v6.Length)], v1, globalState), v51, v51 + v51, map[v49 := 0x2df], v51, v51, v51 + v51, v51, v51[v49 := -v24[safeIndex(76, v24.Length)]], v51, v51];
			v55[safeIndex(728, v55.Length)] := v51;
			var v56: seq<string> := [v5];
			v24[safeIndex(76, v24.Length)], v1, globalState.f6, v55[safeIndex(728, v55.Length)] := |v56|, v6[safeIndex(506, v6.Length)], fm0(globalState), v51 + v51;
		}
		
		var v57: map<bool, int> := map[v1 := v0];
		v57 := v57;
		v6[safeIndex(506, v6.Length)] := v6[safeIndex(506, v6.Length)];
	}
	
	for i4 := v24[safeIndex(76, v24.Length)] to 0x3e6 {
		var v58: array<C0> := new C0[6];
		var v59: C0 := new C0(|"ook"|);
		v58[safeIndex(728, v58.Length)] := v59;
		v58[safeIndex(728, v58.Length)] := v59;
		var v60: multiset<bool> := multiset{v1};
		var v61 := DC2(v59.f14, |v60|, 0xfa, i4);
		var v62 := DC3(v61);
		var v63 := DC3(v61);
		v63 := v63.(cf6 := v61);
		var v64: map<C0, bool> := map[v58[safeIndex(728, v58.Length)] := v6[safeIndex(506, v6.Length)]];
		if (if (v58[safeIndex(728, v58.Length)] in v64) then v64[v58[safeIndex(728, v58.Length)]] else v1 && v6[safeIndex(506, v6.Length)]) {
			var v65: set<bool> := {v6[safeIndex(506, v6.Length)]};
			var v66: multiset<array<bool>> := multiset{v6};
			v24 := new int[14] [i4, v59.f14, i4, safeDivisionInt(|v60|, i4), -0x1ea, safeModuloInt(v59.f14, 461), v0, v59.f14, v24[safeIndex(76, v24.Length)], |v65|, v59.f14, v59.f14, -v59.f14, |v66| * v0];
			m0(!v1, |(seq(abs(556), i5  => (v22)))|, globalState);
			m0(!v1 && v1, v59.f14, globalState);
			v59 := v59;
			var v67: multiset<char> := multiset{v22, v22};
			var v68 := new C0(v24[safeIndex(76, v24.Length)] * (if (v22 in v67) then v67[v22] else i4));
		} else {
			var v69 := DC1(v18 + v18);
			var v70: set<C0> := {v58[safeIndex(728, v58.Length)]};
			var v71 := DC4(v70);
			v6[safeIndex(506, v6.Length)], v1, v6[safeIndex(506, v6.Length)], v69, v6[safeIndex(506, v6.Length)] := !fm1(v24[safeIndex(76, v24.Length)], v6[safeIndex(506, v6.Length)], v6[safeIndex(506, v6.Length)], [v2], globalState), (if (v6[safeIndex(506, v6.Length)]) then v70 else v71.cf7) !! v70, v1, v69, v6[safeIndex(506, v6.Length)];
			v24[safeIndex(76, v24.Length)] := -safeModuloInt(v0, 0x173);
			v24 := new int[7](i6 => safeDivisionInt(i6, -|[v6[safeIndex(506, v6.Length)], v1, v6[safeIndex(506, v6.Length)], true]|));
			var v72 := new C0(-(v0 - v59.f14));
			m0(v1, safeDivisionInt(v59.f14, v59.f14), globalState);
		}
		
		v1 := v6[safeIndex(506, v6.Length)];
	}
	var v73: array<string> := new string[25];
	v73, v6[safeIndex(506, v6.Length)] := v73, v1;
	v24[safeIndex(76, v24.Length)] := --v24[safeIndex(76, v24.Length)];
	var v74: C0 := new C0(v0);
	var v75: seq<C0> := [v74];
	var v76: multiset<C0> := multiset{v74};
	var v77: set<int> := {|v75|, v0, -v74.f14, |v3|, |v76|};
	v1 := (v77 + v77) !! (v77 + v77);
	if (v1) {
		var v78: map<bool, int> := map[v1 := v74.f14];
		var v79: seq<map<bool, int>> := [v78];
		var v80: array<seq<seq<bool>>> := new seq<seq<bool>>[10](i7 => seq(abs(0x2bc), i8  => (v3)));
		var v81: seq<seq<bool>> := [[v1, v1]];
		v80[safeIndex(593, v80.Length)] := v81;
		var v82: map<char, char> := map[v22 := v22];
		var v83: set<map<char, char>> := {v82};
		var v84 := DC6(!v6[safeIndex(506, v6.Length)], v1, v22);
		var v86: set<char> := {v22};
		var v87 := DC8(v86);
		v79, v80[safeIndex(593, v80.Length)], v83, v84 := v79[safeIndex(v74.f14, |v79|) := v78], v81, v83 * (set v88 : map<char, char> | v88 in multiset{v82, map v85 : char | v85 in v87.cf14 :: (v85) := (v22)} :: (v88)), v84;
		m0(DC5(v1, v6[safeIndex(506, v6.Length)]).cf9, safeDivisionInt(v74.f14, v74.f14), globalState);
		v24[safeIndex(76, v24.Length)] := v74.f14;
		v1, v2, v24[safeIndex(76, v24.Length)], globalState.f10, globalState.f6 := |v5| != |v18|, map v89 : int | (-0x2d9 <= v89) && (v89 < 0x2a) :: (safeModuloInt(v89, v74.f14)) := (v1), 0x204, |v5| * v74.f14, -v0;
		v2 := v2[v74.f14 := v1];
	} else {
		if (v1) {
			var v90: seq<set<int>> := [v77];
			v6[safeIndex(506, v6.Length)], v1, v77, v22 := v74.fm4(v1, globalState), v1, v90[safeIndex(v74.f14 + v74.f14, |v90|)], v22;
			var v91: seq<set<char>> := [{v22}];
			var v92: map<C0, set<char>> := map[v74 := v91[safeIndex(v74.f14, |v91|)]];
			var v93 := DC8(if (v74 in v92) then v92[v74] else {v22});
			var v94: multiset<D3> := multiset{v93};
			var v95: multiset<bool> := multiset{v1, v1 || v1, v6[safeIndex(506, v6.Length)], v1, v94 == v94};
			v95 := (v95 * v95) - v95;
			m0(v1, v74.f14, globalState);
			var v96: array<C0> := new C0[28];
			v96 := v96;
			var v97: multiset<int> := multiset{235};
			globalState.f10 := v25[safeIndex(if (v74.f14 in v97) then v97[v74.f14] else v74.f14, |v25|)];
		} else {
			v24 := new int[18];
			var v98: multiset<char> := multiset{v22};
			v24[safeIndex(76, v24.Length)] := |(v98 + multiset(v5))| - v74.f14;
			var v99: set<C0> := {v75[safeIndex(v0, |v75|)]};
			var v100 := DC4(v99 * v99);
			v100 := v100;
			var v101 := DC5(!v6[safeIndex(506, v6.Length)], v1);
			var v102: map<set<array<int>>, D2> := map[{v24, v24, v24} := if (v6[safeIndex(506, v6.Length)]) then v101 else v101];
			v102 := v102[{v24, v24} := v101];
			m0(v6[safeIndex(506, v6.Length)], v74.f14, globalState);
		}
		
		v24 := v24;
		v24[safeIndex(76, v24.Length)] := -v74.f14;
		var v103: multiset<bool> := multiset{v1};
		var v104: map<map<int, bool>, bool> := map[v2 := !v3[safeIndex(if (v6[safeIndex(506, v6.Length)] in v103) then v103[v6[safeIndex(506, v6.Length)]] else v24[safeIndex(76, v24.Length)], |v3|)]];
		v6[safeIndex(506, v6.Length)], globalState.f6, v104 := v103 > v103, safeModuloInt(v0, v74.f14), v104 + v104;
		if (v1) {
			v6[safeIndex(506, v6.Length)] := v0 != v74.f14;
			v24[safeIndex(76, v24.Length)] := v24[safeIndex(76, v24.Length)];
			v1 := v6[safeIndex(506, v6.Length)];
			var v105: seq<seq<int>> := [v25];
			v105 := v105;
			v1 := v6[safeIndex(506, v6.Length)];
		} else {
			v1 := [0x263, v0] < v25;
			v1 := v1 <== v6[safeIndex(506, v6.Length)];
			var v106: array<multiset<C0>> := new multiset<C0>[9];
			v106[safeIndex(243, v106.Length)] := v76;
			v106[safeIndex(243, v106.Length)] := (multiset{v74} + multiset{v74}) * v76[v74 := abs(v24[safeIndex(76, v24.Length)])];
			globalState.f10 := v74.f14 - v74.f14;
			v6[safeIndex(506, v6.Length)] := false;
		}
		
	}
	
	v6[safeIndex(506, v6.Length)] := v6[safeIndex(506, v6.Length)];
	var v107 := DC5(v1, v6[safeIndex(506, v6.Length)]);
	var v108 := DC7(v107);
	match (if (v1) then v108 else v108) {
		case DC5(cf8, cf9) =>
			var v109: map<char, C0> := map[v22 := v74];
			v109 := (map['l' := v74])[v22 := v74];
			var v110: map<int, D3> := map[v0 := DC10(DC5(cf9, cf8), v22, |v5|)];
			var v111 := DC5(cf9, cf8);
			var v112 := DC10(v111, v22, v74.f14);
			v110 := v110[v74.f14 := v112];
			cf9 := !(v5 != "fglj");
			v1 := v74.fm5(globalState);
		case DC6(cf10, cf11, cf12) =>
			v24 := new int[29](i9 => safeDivisionInt(i9, |v25|));
			var v113 := new C0(v74.f14);
			var v114: map<bool, bool> := map[v1 := cf11];
			v6[safeIndex(506, v6.Length)] := if (v6[safeIndex(506, v6.Length)] in v114) then v114[v6[safeIndex(506, v6.Length)]] else false;
			m0(v1, 0x284, globalState);
		case DC4(cf7) =>
			if (v1) {
				var v115 := DC9(v6[safeIndex(506, v6.Length)]);
				var v116: map<int, D3> := map[-(v74.f14 + v0) := v115];
				v116 := v116[v74.f14 := v115.(cf15 := !v1)];
				var v117 := new C0(v74.f14);
				v1 := v1;
				var v118: map<int, array<int>> := map[v0 := v24];
				v118 := v118[v117.f14 := v24] + (v118 + v118);
				v74 := v74;
			} else {
				v18 := v18[|[v1, v6[safeIndex(506, v6.Length)]]| := v74.f14];
				var v119: array<array<string>> := new array<string>[2] [v73, v73];
				v119[safeIndex(43, v119.Length)] := if (v6[safeIndex(506, v6.Length)]) then v73 else v73;
				var v120 := DC4(cf7);
				var v121: set<D2> := {v120, v120, v120, v120, DC4(cf7)};
				v119[safeIndex(43, v119.Length)], v24[safeIndex(76, v24.Length)], v1, globalState.f6, globalState.f6 := v73, v74.f14, v74.f14 == v0, |(v121 + v121)|, safeDivisionInt(v74.f14, |v5|);
				v0 := fm0(globalState);
				v22 := v22;
				var v122 := new C0(v74.f14);
			}
			
			var v123 := new C0(-safeModuloInt(v24[safeIndex(76, v24.Length)], 0x326));
			if ((multiset{v74} * v76) < v76) {
				v24[safeIndex(76, v24.Length)] := |v5| - v74.f14;
				globalState.f6 := v74.f14;
				globalState.f6 := v74.f14;
				var v124: array<D3> := new D3[28];
				var v125: multiset<array<D3>> := multiset{v124, v124, v124, v124, v124};
				v6, globalState.f6, globalState.f6, globalState.f6 := v6, if (DC11(v124).cf19 in v125) then v125[DC11(v124).cf19] else v24[safeIndex(76, v24.Length)] + v74.f14, v123.f14, v74.f14;
				var v126: multiset<array<bool>> := multiset{v6};
				globalState.f6 := safeDivisionInt(v74.f14 + |v126|, --(v24[safeIndex(76, v24.Length)] * v0));
			} else {
				m0(v1, -(v123.f14 + 0xfe), globalState);
				v5 := v5;
				v1 := |v5| <= v24[safeIndex(76, v24.Length)];
				var v127: multiset<int> := multiset{v74.f14};
				globalState.f6, globalState.f6 := (if (v24[safeIndex(76, v24.Length)] in v127) then v127[v24[safeIndex(76, v24.Length)]] else v123.f14) + (v74.f14 * |fm7(v6[safeIndex(506, v6.Length)], v22, globalState)|), -v123.f14;
				v24 := new int[23];
			}
			
			var v128: map<C0, seq<int>> := map[v74 := v25];
			var v129: seq<map<C0, seq<int>>> := [v128 + v128];
			v128 := v129[safeIndex(v74.f14, |v129|)];
		case DC7(cf13) =>
			v74 := new C0(safeDivisionInt(v74.f14, v0));
			v6[safeIndex(506, v6.Length)], globalState.f6 := v24[safeIndex(76, v24.Length)] >= v74.f14, v24[safeIndex(76, v24.Length)];
			v1 := [v1] == fm8(v6[safeIndex(506, v6.Length)], globalState);
			if (v3[safeIndex(v74.f14, |v3|)]) {
				var v130 := DC8({v22, v22, v22, v22, 'v'});
				var v131: set<char> := {v22, v22};
				v130 := v130.(cf14 := v131);
				v130 := if (v6[safeIndex(506, v6.Length)]) then v130 else v130;
				v73[safeIndex(437, v73.Length)] := v5;
				v73[safeIndex(437, v73.Length)] := v5 + v5;
				var v132: array<C0> := new C0[15];
				v132 := v132;
				globalState.f6 := v74.f14;
			} else {
				globalState.f6 := -v24[safeIndex(76, v24.Length)];
				var v133: seq<array<int>> := [v24, v24, v24, v24];
				var v134: multiset<array<int>> := multiset{v133[safeIndex(v24[safeIndex(76, v24.Length)], |v133|)]};
				v6[safeIndex(506, v6.Length)] := v134 > multiset{v24};
				v24[safeIndex(76, v24.Length)] := |(v5 + fm7(v6[safeIndex(506, v6.Length)], fm3(v24[safeIndex(76, v24.Length)], -0x1e, v6[safeIndex(506, v6.Length)], globalState), globalState))|;
				globalState.f10 := -v0;
				var v135: map<set<int>, int> := map[v77 := fm0(globalState)];
				v135 := v135[v77 := v74.f14];
			}
			
	}
	
	m0(false, v74.f14, globalState);
	print v0, "\n";
	print v1, "\n";
	print v2 == map[457 := false], "\n";
	print v3 == [false], "\n";
	print v4 == map[924 := [false]], "\n";
	print v5, "\n";
	print v6[0], "\n";
	print v6[1], "\n";
	print v6[2], "\n";
	print v6[3], "\n";
	print v6[4], "\n";
	print v6[5], "\n";
	print v6[6], "\n";
	print v6[7], "\n";
	print v6[8], "\n";
	print v6[9], "\n";
	print v6[10], "\n";
	print v6[11], "\n";
	print v6[12], "\n";
	print v6[13], "\n";
	print globalState.f0, "\n";
	print globalState.f1, "\n";
	print globalState.f2, "\n";
	print globalState.f3, "\n";
	print globalState.f4 == map[924 := [false], 457 := [false]], "\n";
	print globalState.f5, "\n";
	print globalState.f6, "\n";
	print globalState.f7, "\n";
	print globalState.f8, "\n";
	print globalState.f9, "\n";
	print globalState.f10, "\n";
	print globalState.f11[0], "\n";
	print globalState.f11[1], "\n";
	print globalState.f11[2], "\n";
	print globalState.f11[3], "\n";
	print globalState.f11[4], "\n";
	print globalState.f11[5], "\n";
	print globalState.f11[6], "\n";
	print globalState.f11[7], "\n";
	print globalState.f11[8], "\n";
	print globalState.f11[9], "\n";
	print globalState.f11[10], "\n";
	print globalState.f11[11], "\n";
	print globalState.f11[12], "\n";
	print globalState.f11[13], "\n";
	print globalState.f12, "\n";
	print globalState.f13, "\n";
	print v18 == map[457 := 457], "\n";
	print v20 == [map[]], "\n";
	print v21 == map[1 := [map[]]], "\n";
	print v22, "\n";
	print v23 == map["i" := true], "\n";
	print v24[4], "\n";
	print v25 == [457, 457, 457], "\n";
	print v30 == map[1 := 'n', 457 := 'n'], "\n";
	print v74.f14, "\n";
	print |v75|, "\n";
	print |v76|, "\n";
	print v77 == {1, 457, -457}, "\n";
	print v107.cf8, "\n";
	print v107.cf9, "\n";
	print v108.cf13.cf8, "\n";
	print v108.cf13.cf9, "\n";
}