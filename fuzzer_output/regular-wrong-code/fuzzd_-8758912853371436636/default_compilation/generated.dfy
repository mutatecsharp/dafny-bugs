datatype D0 = DC0(cf0: int) | DC1(cf1: bool, cf2: int, cf3: int, cf4: bool, cf5: int) | DC2(cf6: D0)
datatype D1 = DC4(cf8: array<int>) | DC3(cf7: T0) | DC5(cf9: D1)
datatype D2 = DC7(cf11: seq<char>, cf12: int, cf13: int, cf14: bool) | DC8(cf15: bool, cf16: T0) | DC9(cf17: bool) | DC6(cf10: string) | DC10(cf18: D2)
datatype D3 = DC12 | DC11(cf19: set<char>)
datatype D4 = DC14(cf21: seq<bool>, cf22: char, cf23: int, cf24: int) | DC13(cf20: char) | DC15(cf25: D4)
class GlobalState {
	const f0 : int
	const f1 : int
	const f2 : int
	const f3 : set<int>
	const f4 : int
	const f5 : char
	const f6 : bool
	const f7 : int
	const f8 : bool
	constructor (f0 : int, f1 : int, f2 : int, f3 : set<int>, f4 : int, f5 : char, f6 : bool, f7 : int, f8 : bool) {
		this.f0 := f0;
		this.f1 := f1;
		this.f2 := f2;
		this.f3 := f3;
		this.f4 := f4;
		this.f5 := f5;
		this.f6 := f6;
		this.f7 := f7;
		this.f8 := f8;
	}
	
}

function fm0(globalState: GlobalState): string {
	((seq(0x269, i0  => ('c'))) + "ocmloxxh") + ("nwbfxj" + "ctmkds")
}
function fm2(p0: bool, p1: char, p2: int, p3: bool, globalState: GlobalState): bool {
	|(if (false) then map[0x1ac := true] else map[-0x2ba := !false])| == |((set v1 : int | v1 in (map v0 : int | (0x38b <= v0) && (v0 < 0x1a7) :: (v0 % |multiset{0x351, 0x3ce}|) := ('v')) :: (v1 % 0x2f)) + {0x3bc, 0x2bd})|
}
function fm3(p0: int, globalState: GlobalState): char {
	if (|(map v0 : char | v0 in (seq(0x1f6, i0  => ('q'))) :: (v0) := (|[|"vaujf"|]|))| == --|"uat"|) then 'd' else if (false) then 'u' else 'w'
}
function fm4(p0: bool, p1: bool, globalState: GlobalState): multiset<bool> {
	multiset{|map[|map[true := !!false]| := multiset{|map['i' := !false]|, |[-421, |map[532 := false]|, |"rnwxeunx"|]|, -0x34e, 0xf6, 0x3aa}]| >= -324}
}
function fm5(p0: bool, globalState: GlobalState): int {
	|((if (!true) then [true, false, true, true, false] else [true]) + [true])|
}
method m0(p0: bool, p1: bool, globalState: GlobalState) {
	var i0 := 0;
	while (p1)
		decreases 100 - i0
	{
		if (i0 >= 100) {
			break;
		}
		
		i0 := i0 + 1;
		var v0: array<bool> := new bool[12];
		v0 := v0;
		var v1 := 0x23;
		var v2: array<char> := new char[9](i1 => 'j');
		var v3: map<int, seq<array<char>>> := map[v1 := [v2]];
		var v4: map<int, bool> := map[v1 := p1];
		var v5: seq<array<char>> := [v2, v2];
		var v6 := new C0(if (-fm5(if (-v1 in v4) then v4[-v1] else p0, globalState) in v3) then v3[-fm5(if (-v1 in v4) then v4[-v1] else p0, globalState)] else v5);
		var v7 := true;
		var v8 := "k";
		v7 := !(v1 >= |v8|);
		v7 := p1;
	}
	var v9 := "peut";
	var v10: array<char> := new char[22];
	var v11: seq<array<char>> := [v10, v10, v10, v10, v10];
	var v12: C0 := new C0(v11);
	var v13: map<bool, C0> := map[(seq(-0x36d, i2  => ('w'))) != v9 := v12];
	v13 := v13[p0 := v12];
	var v14 := DC9(p1);
	v9 := match v14 {
		case DC7(cf11, cf12, cf13, cf14) => "ba"
		case DC8(cf15, cf16) => v9
		case DC9(cf17) => v9 + v9
		case DC6(cf10) => cf10 + v9
		case DC10(cf18) => "eohfqvsl"
	};
	v9 := v9;
	if (p0 <== (p0 <== p0)) {
		var v15 := true;
		var v16 := 0x15c;
		var v17: multiset<int> := multiset{v16};
		v15 := (|v9| * v16) in v17;
		v15 := p1;
		v15 := v15;
		var v18: T0 := new C0(v11);
		var v19 := DC3(v18);
		var v20 := DC5(v19);
		var v21: map<bool, seq<D1>> := map[v15 := [v20]];
		var v22: seq<D1> := [DC5(v19), v20, v20, DC5(v19), v20];
		var v23: seq<bool> := [p1, p1];
		var v24: seq<seq<D1>> := [if (!p0 in v21) then v21[!p0] else v22, [v20] + v22, if (v23[v16]) then v22 else v22];
		v16 := --|v24[fm5(v15, globalState)]|;
		v18 := v18;
	} else {
		var v25 := 534;
		var v26 := true;
		var v27: map<bool, bool> := map[!p0 := p0];
		v25, v25, v26 := v25, if (p0 || (if (v26 in v27) then v27[v26] else p0)) then v25 else v25, true;
		if (if (v26) then p1 <== p1 else fm5(v26, globalState) < v25) {
			var v28: array<C0> := new C0[22];
			v28[184] := v12;
			v28[184] := v12;
			var v29: seq<map<bool, C0>> := [v13, v13];
			v25 := |v29|;
			v25 := v25;
			v26 := p0;
			var v30: array<int> := new int[3](i3 => i3 / v25);
			v30 := v30;
		} else {
			v26 := p0;
			var v31 := new C0(v11 + v11);
			var v32 := 'u';
			var v33: multiset<char> := multiset{v32};
			v25 := if (v32 in v33) then v33[v32] else v31.fm1(p0, v26, true, p1, globalState);
			v12 := v12;
			var v34: T0 := new C0([v10]);
			v34 := v34;
		}
		
		var v35: set<bool> := {false};
		var v36 := DC7(v9, |v35|, v25, !true);
		var v37: map<D2, bool> := map[v36 := v26];
		v37 := map[v36 := v26];
		var v38 := 't';
		var v39 := DC1(v26, fm5(v26, globalState), v25, false, v25);
		var v40: map<char, int> := map[v38 := v39.cf3];
		v40 := v40['h' := v25];
		if (v26) {
			v25 := v25 * (|v9| - 0x1b4);
			v38 := v38;
			var v41 := new C0(v11);
			var v42: set<int> := {v25};
			var v43: seq<C0> := [v41, v41];
			var v44: array<int> := new int[15] [-(v25 % -v25), |v42| - v25, v25, v25, |v9|, v25, |v43|, v25, v36.cf13, -|v9|, v25, v36.cf13, v25, v25, v25];
			v44[780] := v25;
			var v45: seq<bool> := [v26, p0];
			var v46 := DC14(v45, 'l', v25, v25);
			var v47: set<char> := {v46.cf22, v38, v38};
			var v48 := DC11(v47);
			v26, v44[780] := p0 ==> true, |(v48.(cf19 := set v49 : char | v49 in v47 :: (v49))).cf19|;
			v45 := v45;
		} else {
			v25 := v25;
			var v50: T0 := new C0(v11);
			var v51 := DC8(p0, v50);
			v26 := (if (true) then v51 else v51).cf15;
			var v52: multiset<char> := multiset{'v', v38};
			var v53: map<int, string> := map[v39.cf2 := v9];
			var v54: seq<bool> := [multiset{v38, v38} >= v52, !false, v25 <= |(if (v25 in v53) then v53[v25] else v9)|];
			v54 := v54 + v54[|v9| := p1];
			v25 := v25 + v25;
			v25 := v25;
		}
		
	}
	
	var v55 := 0x73;
	var v56: set<int> := {v55, v55};
	var v57: map<set<int>, string> := map[v56 := v9];
	v57 := v57 + v57;
}
trait T0 {
	var f9 : seq<array<char>>
	function fm1(p0: bool, p1: bool, p2: bool, p3: bool, globalState: GlobalState): int 
}

class C0 extends T0 {
	constructor (f9 : seq<array<char>>) {
		this.f9 := f9;
	}
	
	function fm1(p0: bool, p1: bool, p2: bool, p3: bool, globalState: GlobalState): int {
		|((seq(134, i0  => (0x8d))) + (seq(-0x3c7, i1  => (0x51))))| - -|([true, !false, DC1(true, 0x1e6, |[true, true]|, true, |multiset([308])|).cf4] + [true])|
	}
}

method Main() {
	var v0 := 159;
	var v1 := "vqomatru";
	var v2: seq<int> := [v0, --v0, v0, |v1|];
	var globalState := new GlobalState(0x35e, 771, 785, set v3 : int | v3 in v2 :: (v3 + 0x1dd), -185, 'u', false, 0x181, true);
	var v4 := DC0(|(fm0(globalState) + v1)|);
	match (v4) {
		case DC0(cf0) =>
			v0 := cf0;
			var v5 := false;
			v5 := v5;
			var v6: set<string> := {v1, "erdcbq", v1 + v1};
			var v7: array<int> := new int[13];
			v7[895] := v0;
			v6, v7[895] := v6 - {v1}, v0;
			var v8: seq<bool> := [v5, v5 && v5];
			v5 := v8[cf0];
		case DC1(cf1, cf2, cf3, cf4, cf5) =>
			var v9: set<bool> := {false};
			var v10: set<set<bool>> := {v9, v9};
			cf4 := !((v10 + v10) >= v10);
			m0(cf1, cf1, globalState);
			var v11: array<bool> := new bool[23];
			v11[611] := cf1;
			v11[611] := if (cf1) then v0 > cf2 else cf4 ==> false;
			var v12: array<char> := new char[1] ['o'];
			var v13: seq<array<char>> := [v12];
			var v14 := new C0(v13[cf5 := v12]);
		case DC2(cf6) =>
			var v15 := true;
			var v16: map<bool, int> := map[v15 <==> v15 := v0];
			v16 := v16[v15 := v0];
			var v17 := 'd';
			v15 := fm2(!v15, v17, v0, if (!v15) then v15 else v15, globalState);
			var v18: map<int, bool> := map[-v0 := v15 || v15];
			v18 := v18[-0x22e := v15 || false];
			var v19: array<array<D0>> := new array<D0>[1];
			var v20: array<D0> := new D0[9];
			v19[937] := v20;
			v19[937] := v20;
	}
	
	var v21: array<set<C0>> := new set<C0>[14];
	var v22 := 'h';
	var v23: array<char> := new char[12] [v22, v22, v22, v22, v22, v22, v22, 'q', v22, v22, v22, v22];
	var v24: seq<array<char>> := [v23];
	var v25: C0 := new C0(v24);
	var v26: set<C0> := {v25, v25};
	v21[611] := v26;
	v21[611] := v26;
	match (v4) {
		case DC0(cf0) =>
			v1 := "rllcb" + ("t")[v0 := v22];
			var v27: T0 := new C0([v23, v23]);
			v27 := v27;
			v0 := v0;
			var v28: array<bool> := new bool[9](i0 => false <==> true);
			var v29 := true;
			v28[208] := v29;
			var v30: set<bool> := {v29};
			v28[208] := if (v29) then v29 ==> DC1(v29, |v26|, -v0, v29, 0x22c).cf4 else |v30| <= v0;
		case DC1(cf1, cf2, cf3, cf4, cf5) =>
			var v31: array<int> := new int[22];
			var v32: array<bool> := new bool[14];
			var v33: map<array<int>, array<bool>> := map[v31 := v32];
			v33 := v33[v31 := v32];
			v25 := v25;
			v1 := v1 + v1;
			if (cf4) {
				var v34 := DC1(false, cf5, v0, cf1, cf2);
				cf3 := v34.cf2;
				var v35: seq<array<bool>> := [v32, v32];
				var v36: seq<array<bool>> := [v35[cf3], v32, v32];
				var v37: map<bool, int> := map[cf1 := |(v35 + v35)|];
				var v38: seq<bool> := [cf1];
				var v39: map<string, seq<bool>> := map[v1 := v38];
				v37 := v37[map[v1 := v38] == v39 := v0];
				v31[327] := if (cf4) then cf2 else --v0;
				var v40: set<int> := {v25.fm1(fm2(cf1, 'w', |v2|, !cf4, globalState), cf1, cf1, cf1, globalState), cf3, cf2};
				v31[327] := v0 + |v40|;
				cf1 := v34.cf1;
				cf1 := !!fm2(cf4, if (cf1) then v22 else fm3(v34.cf3, globalState), cf3 - -v0, cf1, globalState);
			} else {
				m0(cf4 <==> false, cf1, globalState);
				v0 := v2[cf5] + -0x1db;
				v32[154] := cf1;
				v32[154] := v25 in v26;
				var v41: array<map<bool, int>> := new map<bool, int>[19];
				v41 := v41;
				var v42: array<array<int>> := new array<int>[15];
				v42 := v42;
			}
			
		case DC2(cf6) =>
			var v43 := false;
			if (!v43) {
				v0 := v0;
				var v44: seq<bool> := [v43, v43, v43, v43, v43];
				var v45: multiset<bool> := multiset{fm2(v43, v22, v0, !v43, globalState)};
				var v46: map<int, int> := map[v0 := v0];
				var v47 := DC1(v43, |v45|, |v46|, v43, v0);
				var v48: T0 := new C0(v24);
				var v49: map<T0, bool> := map[v48 := v43];
				var v50: multiset<int> := multiset{v0, |(seq(0x211, i1  => (v22)))|};
				var v51 := DC1(v47.cf4, v0, |v49[v48 := v43]|, v43, if (v0 in v50) then v50[v0] else v2[v0]);
				var v52: multiset<C0> := multiset{v25};
				var v53: array<int> := new int[12] [v0, |v44|, v0, 916, v47.cf5, v0, |v52| / -v48.fm1(v43, v43, v43, false, globalState), |v1|, v0, 0x32a, v0, v0];
				v53[266] := v0;
				v53[266] := v48.fm1(v0 < v0, v43, v43, if (fm2(fm2(v43, v22, v0, v43, globalState), v22, v0, v43, globalState)) then fm2(v43, v22, |v1|, v43, globalState) else false, globalState);
				v53[266] := v25.fm1(v43, fm2(v43, v22, -0x3c2, v43, globalState), v43, !(v50 >= v50), globalState);
				v0 := |v1|;
				v43 := !v43;
			} else {
				var v54: T0 := new C0(v24);
				var v55: seq<T0> := [v54, v54, v54, v54];
				var v56 := DC3(v55[v0]);
				var v57: array<T0> := new T0[21] [v54, DC3(v54).cf7, v54, v56.cf7, v54, v54, v54, v54, v54, v54, v54, v54, v54, v54, v54, v54, v54, v54, v54, v54, v54];
				v57[216] := v54;
				v57[216] := if (v43) then v54 else if (v43) then v54 else v54;
				var v58: map<bool, bool> := map[v43 := false];
				v58 := v58[v43 := v43];
				var v59: array<bool> := new bool[14];
				v59 := v59;
				var v60: array<map<int, char>> := new map<int, char>[24];
				v60 := new map<int, char>[17];
				var v61: seq<bool> := [v43, v43, v43];
				v61 := v61;
			}
			
			var v62: array<int> := new int[16];
			v62[992] := v0;
			v62[992] := v0;
			v43 := v43;
			v0 := v0;
	}
	
	var v63 := true;
	if (v63) {
		v63 := v63;
		var v64: map<int, bool> := map[v0 := false];
		v0 := (|v64| * v0) + v0;
		v1 := fm0(globalState);
		if (!v63) {
			v63 := v63;
			var v65: array<bool> := new bool[20];
			var v66: seq<bool> := [v63];
			var v67: set<int> := {|v66|, v0};
			var v68 := DC1(v63, v0, v0, v63, |v67|);
			v65[698] := v68.cf1;
			v65[698] := true;
			v65[698] := v63;
			var v69: map<bool, seq<bool>> := map[v65[698] := v66];
			v69 := v69;
			var v70: T0 := new C0(v24);
			var v71: map<T0, int> := map[v70 := v0];
			v71 := v71[v70 := v0];
		} else {
			var v72: map<bool, char> := map[v63 := 'm'];
			v72 := v72;
			var v73: multiset<bool> := multiset{v63};
			var v74: seq<bool> := [v63];
			var v75: array<multiset<bool>> := new multiset<bool>[26] [v73, fm4(false, v63, globalState), v73, multiset{false, v63, v63, v63, v63}, v73, v73, v73, multiset{v63}, v73, v73, fm4(v63, false, globalState), v73[v63 := v0], v73, v73, v73, v73, v73, v73, multiset{v63}, v73, multiset{v63}, v73, v73, fm4(fm2(v63, v1[|v64|], v0, v63, globalState), v74[0x60], globalState), v73, v73];
			var v76: map<array<multiset<bool>>, bool> := map[v75 := v63 <== v63];
			v63 := if (v75 in v76) then v76[v75] else v63;
			v0 := v0;
			v21[611] := v21[611] - v26;
			var v77: array<string> := new string[12](i2 => v1);
			v77[809] := v1;
			var v78: array<bool> := new bool[23];
			v78[965] := !v63;
			v77[809], v78[965] := v1, !v63 && v63;
		}
		
		var v79: seq<bool> := [v63];
		v63 := !!v63 <== fm2(v63, v22, |v2|, fm2(v63, v22, |v79|, v63, globalState), globalState);
	} else {
		v23[18] := v1[v0];
		v63, v0, v23[18], v0, v63 := false, v0, v22, if (v63) then v0 else v0, true;
		var v80: array<seq<int>> := new seq<int>[28](i3 => v2 + [v0, v0, v0]);
		v80 := v80;
		var v81: map<seq<int>, string> := map[v2 := v1];
		v81 := v81;
		v63 := v63;
		v23[18] := v23[18];
	}
	
	m0(v63, v63 <==> v63, globalState);
	v63 := v63;
	v63 := v63;
	var v82: array<seq<int>> := new seq<int>[1](i4 => v2);
	var v83: multiset<bool> := multiset{v63, v63};
	var v84: map<array<seq<int>>, multiset<bool>> := map[v82 := v83];
	var v85: seq<multiset<bool>> := [if (v82 in v84) then v84[v82] else v83, v83, v83];
	v85 := v85;
	var v86 := new C0(v24);
	m0(v63, v63, globalState);
	var v87: array<int> := new int[16](i5 => i5 % v2[v0]);
	v87[544] := v0;
	var v88: map<bool, int> := map[v63 := 211];
	var v89: set<int> := {|v2|};
	v87[544] := -((if (v63 in v88) then v88[v63] else v0) % |map[|v89| := v25.fm1(false, v63, !v63, !v63, globalState)]|);
	v87[544] := -(v0 % v87[544]);
	v87[544] := v0 - v87[544];
	var v90: array<string> := new string[2];
	v90[918] := v1;
	v90[918] := v1;
	var v91: map<int, int> := map[v87[544] := v87[544]];
	v91 := if (v63) then map v92 : int | v92 in v2 :: (v92 * v0) := (v87[544]) else v91;
	var v93: set<multiset<bool>> := {v83, v83[v63 := v87[544]], multiset{v63}, v83};
	if ({v83, v83} !! v93) {
		var v94: seq<string> := [v1, "rqfejscgi", v1, seq(-0x166, i6  => (v22))];
		v90[918], v87[544] := v94[v87[544]] + (v1 + v90[918]), v87[544] - v87[544];
		var v95: multiset<int> := multiset{v0};
		var v97: map<multiset<int>, set<int>> := map[v95 := v89];
		m0(multiset([v86.fm1(v63, v63, v63, v63, globalState), v0]) <= v95, (set v96 : int | (-0x3e <= v96) && (v96 < 728) :: (v96 - v0)) !! (if (multiset{|"wbacwiua"|} in v97) then v97[multiset{|"wbacwiua"|}] else v89), globalState);
		v87[544] := -v0;
		var v98 := DC7(v90[918], |v1|, v87[544], v63);
		var v99: multiset<string> := multiset{v98.cf11, v1, v90[918]};
		var v100: map<multiset<string>, bool> := map[v99 := v63];
		v100 := v100[multiset{v1} - v99 := false];
		var v101: map<int, bool> := map[v87[544] := v63];
		v101 := map[v86.fm1(v63, v63, v63, v63, globalState) := v63];
	} else {
		v63 := v63;
		var v102: T0 := new C0(v24 + v24);
		var v103: seq<bool> := [v63, v63, false, v63];
		var v104 := DC7(v90[918][v0 := v22], v0, -|{v63}|, v63);
		var v105: map<int, bool> := map[|v1| := v63];
		var v106 := DC4(v87);
		var v107: seq<array<int>> := [v87, v106.cf8];
		v0, v102, v63, v102 := v2[v102.fm1(v63, v103[v104.cf13], v63, !!v63, globalState)], v102, if (!fm2(!v63, fm3(|v105|, globalState), v87[544], v63, globalState)) then !(v87[544] > |v107|) else v63, v102;
		v22 := if (v63 <==> v63) then v22 else v22;
		v25 := v25;
		v91 := v91[v0 := v102.fm1(v63, v63, true, v63, globalState)];
	}
	
}