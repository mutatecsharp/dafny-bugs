datatype D0 = DC0(cf0: bool)
datatype D1 = DC2 | DC1(cf1: char) | DC3(cf2: D1)
datatype D2 = DC5(cf4: bool, cf5: int, cf6: int, cf7: int, cf8: bool) | DC4(cf3: multiset<bool>)
datatype D3 = DC7(cf10: bool) | DC6(cf9: set<int>)
datatype D4 = DC9(cf12: int, cf13: seq<array<bool>>, cf14: int, cf15: int) | DC10(cf16: int, cf17: bool) | DC8(cf11: array<array<int>>)
datatype D5 = DC12(cf19: bool) | DC13(cf20: bool, cf21: char) | DC11(cf18: map<bool, int>) | DC14(cf22: D5)
datatype D6 = DC16(cf24: map<multiset<bool>, int>, cf25: seq<bool>) | DC15(cf23: array<bool>)
datatype D7 = DC18(cf27: bool, cf28: bool) | DC19(cf29: int, cf30: bool, cf31: int) | DC20(cf32: char, cf33: int) | DC17(cf26: string)
datatype D8 = DC22(cf35: int) | DC21(cf34: C3)
datatype D9 = DC24(cf37: int, cf38: int) | DC23(cf36: array<int>)
datatype D10 = DC26(cf40: int) | DC25(cf39: map<int, int>) | DC27(cf41: D10)
datatype D11 = DC29(cf43: map<int, int>) | DC30 | DC28(cf42: T1) | DC31(cf44: D11)
datatype D12 = DC33(cf46: string) | DC32(cf45: map<T0, array<T1>>)
datatype D13 = DC34(cf47: C1)
datatype D14 = DC36(cf49: bool, cf50: seq<bool>, cf51: seq<C2>, cf52: seq<seq<D8>>, cf53: char) | DC35(cf48: array<string>) | DC37(cf54: D14)
class GlobalState {
	const f0 : int
	const f1 : bool
	const f2 : map<string, array<bool>>
	var f3 : bool
	var f4 : bool
	const f5 : int
	const f6 : string
	constructor (f0 : int, f1 : bool, f2 : map<string, array<bool>>, f3 : bool, f4 : bool, f5 : int, f6 : string) {
		this.f0 := f0;
		this.f1 := f1;
		this.f2 := f2;
		this.f3 := f3;
		this.f4 := f4;
		this.f5 := f5;
		this.f6 := f6;
	}
	
}

function fm5(globalState: GlobalState): char {
	if (608 != 0x382) then 'v' else 'l'
}
function fm8(p0: bool, p1: seq<int>, p2: bool, globalState: GlobalState): string {
	seq(100, i0  => (DC1('d').cf1))
}
function fm9(p0: bool, globalState: GlobalState): map<bool, bool> {
	(map[true := false] + map[true := false]) + map[!!true := true]
}
function fm10(p0: string, p1: bool, p2: int, p3: int, globalState: GlobalState): seq<D0> {
	match DC30() {
		case DC29(cf43) => (seq(-0xde, i0  => (DC0(true)))) + [DC0(false), DC0(false), DC0(false)]
		case DC30() => [DC0(!false), DC0(!!!false), DC0(!true), DC0(true)]
		case DC28(cf42) => [DC0(true), DC0(true), DC0(!!false), DC0(true)] + (seq(0x251, i1  => (DC0(false))))
		case DC31(cf44) => [DC0(!false)] + [DC0(true)]
	}
}
function fm16(p0: bool, globalState: GlobalState): int {
	--|(({true, false} - {false}) * {!!false})|
}
function fm17(p0: char, p1: bool, globalState: GlobalState): map<bool, seq<int>> {
	map[true := [|(seq(-0x303, i0  => ('e')))|]]
}
function fm20(p0: bool, p1: bool, p2: bool, globalState: GlobalState): int {
	--0x1cd * (|multiset(seq(-0x157, i0  => (|map['e' := true]|)))| + 518)
}
function fm21(p0: int, p1: int, p2: int, p3: string, globalState: GlobalState): seq<int> {
	([|map["yvab" := true]|] + [|(map v0 : int | (927 <= v0) && (v0 < 999) :: (v0 + |"ykwfkuyl"|) := (298))|, 0x138, |"h"|, 792]) + [0x399]
}
function fm22(globalState: GlobalState): multiset<D1> {
	multiset{DC3(DC2())} + multiset{DC3(DC1('y')), DC3(DC2()), DC3(DC2())}
}
function fm23(p0: int, globalState: GlobalState): D1 {
	match DC0(false) {
		case DC0(cf0) => DC3(DC2())
	}
}
function fm24(p0: bool, p1: int, p2: seq<char>, globalState: GlobalState): bool {
	!(|[!true, !false]| != 0x23b)
}
function fm25(globalState: GlobalState): map<bool, bool> {
	map[false := !true] + (map[true := false] + map[false := false])
}
function fm26(p0: multiset<int>, p1: bool, globalState: GlobalState): D3 {
	DC6({644} + (set v0 : int | (-0x38b <= v0) && (v0 < -0xda) :: (v0 * -0x3d2)))
}
function fm27(globalState: GlobalState): set<multiset<bool>> {
	{multiset{true, true, false}} * {multiset{true}}
}
function fm28(p0: bool, globalState: GlobalState): multiset<bool> {
	DC4(multiset([!true])).cf3
}
function fm29(p0: bool, p1: multiset<int>, globalState: GlobalState): set<int> {
	match DC25(map[529 := |DC4(multiset{false}).cf3|]) {
		case DC26(cf40) => DC6({|(map v0 : int | (-0x392 <= v0) && (v0 < 258) :: (v0 - cf40) := (true))|}).cf9
		case DC25(cf39) => set v1 : int | (0x23e <= v1) && (v1 < 0x168) :: (v1 / |[|(set v2 : int | (0x304 <= v2) && (v2 < 0x215) :: (v2 + |[|multiset{true, false}|, |[-0x134]|, -0x392, |map[false := true]|, |multiset{50}|]|))|]|)
		case DC27(cf41) => {|"xlnmqckl"|}
	}
}
function fm30(p0: bool, globalState: GlobalState): seq<bool> {
	if (0x2d0 < |(map v0 : D3 | v0 in map[DC6({|"rsyxbepi"|, |[|map[|(set v1 : int | (0x355 <= v1) && (v1 < 631) :: (v1 + |{false, false}|))| := |{true}|]|, |{true, false, !false}|]|}) := "veyog"] :: (v0) := (-0x289))|) then [!true] + [false] else if (!!true) then [true, true, true, false, true] else [true, !true]
}
function fm31(p0: int, p1: multiset<bool>, globalState: GlobalState): set<map<int, int>> {
	{map[-|[!!false, false]| := |[617]|] + map[0x37 := |(seq(-0x3a8, i0  => ('y')))|], map[0x1a8 := DC22(|[-954, |"qejrjkfg"|]|).cf35] + map[|(seq(140, i1  => (360)))| := |"dhvkev"|], map[0x2c9 := |(seq(108, i2  => ('q')))|]}
}
function fm32(globalState: GlobalState): D5 {
	match DC13(false, 'y') {
		case DC12(cf19) => DC13(cf19, 'r')
		case DC13(cf20, cf21) => DC13(cf20, cf21)
		case DC11(cf18) => DC13(true, 'f')
		case DC14(cf22) => DC13(false, 'f')
	}
}
function fm33(globalState: GlobalState): D0 {
	DC0(--|"qqnloowd"| == -548)
}
function fm34(p0: bool, p1: bool, globalState: GlobalState): set<D6> {
	{DC16(map[multiset{false} := |map[false := 0x9b]|], [true]), DC16(map[multiset{false} := |[|[true]|]|], [false])} * (set v0 : D6 | v0 in multiset([DC16(map[multiset{true} := 0x2bb], [false, true])]) :: (v0))
}
function fm35(p0: int, globalState: GlobalState): D6 {
	DC16((map v0 : multiset<bool> | v0 in multiset{multiset{true}, multiset{true}} :: (v0) := (|[false]|)) + (map v1 : multiset<bool> | v1 in (map v2 : multiset<bool> | v2 in (seq(83, i0  => (multiset{false, true}))) :: (v2) := (0x307)) :: (v1) := (|[true, true, false, false, !false]|)), [false])
}
function fm36(p0: bool, p1: bool, globalState: GlobalState): set<multiset<D1>> {
	(set v0 : multiset<D1> | v0 in {multiset{DC3(DC2()), DC3(DC1('f'))}} :: (v0)) + ({multiset{DC3(DC3(DC2())), DC3(DC2()), DC3(DC2()), DC3(DC2())}, multiset{DC3(DC3(DC2()))}} * {multiset([DC3(DC2()), DC3(DC1('m')), DC3(DC2())])})
}
function fm37(globalState: GlobalState): map<int, string> {
	if (!(|"npm"| < -0x332)) then map v0 : int | v0 in {0x5d, |[false]|, 0x2b6, 980} :: (v0 + 302) := (seq(0xb8, i0  => ('f'))) else (map v1 : int | v1 in [0x2dc, 0x37e] :: (v1 - -0x272) := ("nf")) + map[|"r"| := seq(-0x228, i1  => ('r'))]
}
function fm38(p0: bool, p1: bool, p2: bool, p3: int, globalState: GlobalState): map<int, int> {
	if (|(map v0 : multiset<bool> | v0 in [multiset{false}, multiset([false, false])] :: (v0) := (-0xed))| < 0x2f) then map[--0x113 := 0x127] + (map v1 : int | v1 in map[|multiset{!false}| := false] :: (v1 * --0x2e8) := (-645)) else map[|(map v2 : int | (0x377 <= v2) && (v2 < -0x33a) :: (v2 % -0x33c) := (true))| := 28]
}
function fm39(p0: int, p1: D3, p2: bool, p3: bool, globalState: GlobalState): map<bool, int> {
	(map[true := |multiset([|(map v0 : int | v0 in map[0x1f4 := |map[[|multiset{!true}|] := true]|] :: (v0 % -296) := ({true}))|])|] + map[true := |{true, true}|]) + (map[false := |(map v1 : int | v1 in (seq(667, i0  => (0x195))) :: (v1 + |"qqnmhka"|) := (0x3c5))|] + map[!!!true := 0xe0])
}
method m0(p0: int, p1: bool, p2: bool, globalState: GlobalState) returns (r0: array<string>, r1: bool) {
	var v0: seq<int> := [-0x68];
	var v1: map<seq<int>, bool> := map[v0 := p2];
	var v2 := DC18(p2, if ((seq(0x2ec, i0  => (p0))) in v1) then v1[seq(0x2ec, i0  => (p0))] else p1);
	match (v2) {
		case DC18(cf27, cf28) =>
			var v3: seq<bool> := [!true];
			var v4: map<D7, seq<bool>> := map[DC17("mlycpc") := v3];
			var v5: map<map<D7, seq<bool>>, bool> := map[v4 := if (p2) then true else p1];
			v5 := v5[v4 := cf28];
			var v6: map<int, bool> := map[p0 := cf27];
			var v7: map<map<int, bool>, int> := map[v6 := p0];
			var v8 := "n";
			var v9 := new C5((seq(0x38f, i1  => ('k')))[|v7| := 'n'] + v8, --p0, p0, p0);
			var v11: C4 := new C4(p1, |(map v10 : int | v10 in v6 :: (v10 % v9.f10) := (v9.f10))|, |fm21(p0, -p0, p0, "etffgh", globalState)|);
			var v12: multiset<C4> := multiset{v11};
			var v13: seq<C4> := [v11];
			var v14: multiset<multiset<C4>> := multiset{v12, v12, multiset{v11, v11}, multiset(v13)};
			globalState.f3 := v14 >= multiset{v12, multiset{v11}};
			var v15 := -0x3da;
			v15 := |{v9.f10, 0x36d, p0, v15}|;
		case DC19(cf29, cf30, cf31) =>
			var v16 := new C3(p2, fm20(p2, p1, !p1, globalState) % cf31, -0x2c5);
			var v17 := "rtmdkqwc";
			var v18: set<int> := {v16.fm13(v17, globalState), p0};
			var v19: map<set<int>, bool> := map[v18 * {|[p1, !!!p1]|, p0} := if (v16.f12) then true else cf30];
			v19 := v19[v18 := cf30 ==> v16.f12];
			cf31 := cf31;
			cf30 := "ruevlxjng" <= (v17 + v17);
		case DC20(cf32, cf33) =>
			globalState.f4 := p1;
			var v20: array<int> := new int[4];
			var v21: array<map<int, string>> := new map<int, string>[12];
			var v22 := "xjka";
			var v23: map<int, string> := map[cf33 := v22];
			v21[509] := v23;
			var v24 := DC23(v20);
			var v25: map<bool, bool> := map[p2 := p2];
			var v26: seq<bool> := [p1, (if (false in v25) then v25[false] else p1) <== p1];
			var v27: multiset<int> := multiset{-877};
			var v28: set<int> := {p0};
			var v29: C5 := new C5(v22, cf33, |v27|, |v28|);
			var v30: map<C5, bool> := map[v29 := p2];
			var v31: map<int, map<int, string>> := map[v29.f10 % p0 := fm37(globalState) + fm37(globalState)];
			globalState.f4, v20, globalState.f4, v21[509] := p1, v24.cf36, v26[|v30| * v29.f10], if (-v29.f10 in v31) then v31[-v29.f10] else v23 + v23;
			var v32: array<string> := new string[17];
			var v33: map<array<string>, int> := map[v32 := -969];
			v33 := v33[v32 := cf33];
			cf33 := (cf33 - cf33) % p0;
		case DC17(cf26) =>
			var v34: seq<bool> := [!true];
			var v35: array<int> := new int[18] [|fm30(false, globalState)|, 0x7a, |[true]|, fm20(p1, p2, p1, globalState), p0, p0, p0, p0, p0, p0, 674, p0, p0, -|v34|, -p0, p0, p0, 0x375];
			var v36: seq<array<int>> := [v35, v35];
			var v37: map<bool, array<int>> := map[p1 := v36[p0]];
			var v38: map<int, bool> := map[-p0 := p1];
			v37 := v37[if (p0 in v38) then v38[p0] else p1 := v35];
			globalState.f4 := if (p2) then if (p2) then false else p1 else true;
			if (p0 > 0xf0) {
				var v39: C4 := new C4(p1, |(fm30(!!false, globalState))[p0 := p1]|, -p0);
				var v40: T0 := new C1(p0, |"mhnx"|);
				var v41: seq<T0> := [v40, v40];
				v39, v41 := v39, v41;
				var v42 := 0x0;
				v42 := (v40.f8 % v40.f7) % --v42;
				v0 := if (p1) then v0 else v0;
				v35[581] := -p0;
				var v43: map<bool, int> := map[p1 := p0];
				v35[581] := v42 * (if (fm24(v39.f11, v40.f7, cf26, globalState) in v43) then v43[fm24(v39.f11, v40.f7, cf26, globalState)] else v42);
				var v44: multiset<int> := multiset{|fm28(p2, globalState)|, v42, |cf26|, v42};
				var v45: array<bool> := new bool[9] [v39.f11, p2, v42 <= v40.f7, p2, p1 ==> p1, p2, true, v44 == v44, v39.f11];
				var v46 := DC25(fm38(v39.fm1({p2}, p0, -v42, 554, globalState), true, fm24(p2, 816, cf26, globalState), v40.f8, globalState));
				var v47: set<int> := {|v46.cf39|};
				v45[781] := v47 <= (set v48 : int | (0x6b <= v48) && (v48 < 0x282) :: (v48 * v40.f8));
				var v49: array<char> := new char[29];
				v49[950] := 'e';
				v45[931] := p1;
				var v50: array<map<bool, int>> := new map<bool, int>[18](i2 => map[p1 := v40.f8]);
				v50[514] := v43 + v43;
				var v51 := 'v';
				var v52 := DC6(v47);
				v45[781], v42, v49[950], v45[931], v50[514] := v34[v40.f8], v42, v51, p2, fm39(v40.f8, v52, p2, v39.f11, globalState);
			} else {
				v35[18] := p0;
				v35[18] := |v0| % p0;
				v35[18] := v35[18];
				v35[18] := p0;
				var v53: map<bool, bool> := map[p2 := p2];
				var v54 := 'c';
				var v55: array<bool> := new bool[23] [p2, p2, p2, !fm24(p2, |v53|, [v54], globalState), !p1, p2, p1, p2, p2, p2, p1, p2, p2, p2, p1, p1, true, p1, p1, p1, p2, p1, true];
				var v56: seq<array<bool>> := [v55];
				var v57: multiset<int> := multiset{v35[18], |cf26|};
				v35[18] := DC9(p0, v56, v35[18], if ((if (516 in v57) then v57[516] else v35[18]) in v57) then v57[if (516 in v57) then v57[516] else v35[18]] else 0x248).cf15 + -v35[18];
				var v58: array<array<int>> := new array<int>[16];
				var v59: array<int> := new int[18];
				v58[22] := v59;
				v58[22] := v59;
			}
			
			var v60: C3 := new C3(p1, p0, p0);
			var v61 := DC21(v60);
			var v62: map<D8, bool> := map[v61 := p2];
			v62 := v62[DC21(v60) := false];
	}
	
	var v63 := 's';
	var v64: array<bool> := new bool[9] [!!p2, p2, p1, p2, p1, true, p1, false, p1];
	var v65 := DC15(v64);
	var v66 := "nbthqcwcx";
	var v67: map<D6, string> := map[v65 := v66];
	var i3 := 0;
	while (v63 in (if (v65 in v67) then v67[v65] else v66))
		decreases 100 - i3
	{
		if (i3 >= 100) {
			break;
		}
		
		i3 := i3 + 1;
		var v68: T0 := new C5("e", p0, p0, p0);
		var v69: T1 := new C5(v66, v68.f8, |"ks"|, v68.f7);
		var v70 := DC28(v69);
		var v71: array<T1> := new T1[15] [v69, v69, v69, v70.cf42, v69, v69, v69, v69, v69, v69, v69, v69, v69, v69, v69];
		var v72: map<T0, array<T1>> := map[v68 := v71];
		var v73 := DC32(v72);
		var v74: seq<map<T0, array<T1>>> := [map[v68 := v71]];
		var v75: map<int, map<T0, array<T1>>> := map[v68.f8 := v72];
		var v76: array<map<T0, array<T1>>> := new map<T0, array<T1>>[21] [v72, v73.cf45 + v72, v74[0x2c5], map[v68 := v71], v72, map[v68 := v71], v72, v72, v72, v72[v68 := v71], v72, map[v68 := v71], v72, v72, v72 + v72, v72, v72, v72, v72[v68 := v71], (v72[v68 := v71])[v68 := v71], if (fm16(p2, globalState) in v75) then v75[fm16(p2, globalState)] else v72];
		v76[145] := v72;
		v76[145] := v72 + (v72 + v72);
		var v77: map<bool, T1> := map[true := v69];
		v77 := v77[fm24(p2, |v66[v68.f7 := v63]|, v66, globalState) := v69];
		v64 := new bool[23];
		var v78: map<T0, T1> := map[v68 := v69];
		v78 := map[v68 := v69];
	}
	v64[630] := v66 <= v66;
	v64[630] := p0 >= p0;
	var v79: array<array<int>> := new array<int>[3];
	var v80: array<int> := new int[19];
	v79[944] := v80;
	v79[944] := v80;
	if (v64[630]) {
		var v81 := DC19(p0, v64[630], p0);
		var v82: seq<D7> := [v81];
		var v83: multiset<int> := multiset{p0, p0, |v82|, p0, |v66|};
		v83 := v83;
		v79[944][980] := 0xed;
		var v84: multiset<bool> := multiset{p2, v64[630]};
		var v85: C1 := new C1(p0, p0);
		var v86: map<C1, bool> := map[v85 := p1];
		var v87 := DC34(v85);
		var v88 := DC23(v79[944]);
		var v89: map<int, array<int>> := map[p0 := v88.cf36];
		var v90: array<D9> := new D9[18];
		var v91: map<bool, array<D9>> := map[p1 := v90];
		v64[630], globalState.f3, v79[944], globalState.f4, v79[944][980] := v84 !! v84, if (v87.cf47 in v86) then v86[v87.cf47] else true, if (|(map[p2 := v90] + v91)| in v89) then v89[|(map[p2 := v90] + v91)|] else if (p1) then v79[944] else v80, false, p0 - -0x21d;
		globalState.f4 := true;
		var v92: map<seq<int>, int> := map[v0[-0x2df := p0] + [v79[944][980]] := p0];
		v92 := (v92 + v92) + v92;
		var v93: array<seq<bool>> := new seq<bool>[5];
		v93 := if (true) then v93 else v93;
	} else {
		var v94 := -0x6a;
		var v95: multiset<bool> := multiset{p2};
		globalState.f4, v94, globalState.f4, globalState.f3 := false, v94, -v94 == p0, (0x109 % (if (p2 in v95) then v95[p2] else -v94)) >= p0;
		globalState.f4 := p2;
		var v96: map<bool, multiset<int>> := map[|v66| != p0 := multiset(v0)];
		v96 := v96[p2 := multiset(v0)];
		v79[944][869] := p0;
		v79[944][869] := fm16(p1, globalState) - p0;
		var v97 := DC1('d');
		match (DC3(v97)) {
			case DC2() =>
				r0 := new string[7] [v66, "mfn", if (!p1) then v66 else v66, "ilvyw", v66, v66, v66];
				v94 := v94;
				var v98 := DC7(p1);
				v98 := v98;
				var v99: array<array<bool>> := new array<bool>[9];
				v99, v79[944][869] := v99, v79[944][869];
			case DC1(cf1) =>
				globalState.f3 := cf1 !in "lq";
				var v100: T1 := new C5(v66, v79[944][869], p0, 780);
				var v101: map<T1, int> := map[v100 := p0];
				var v102: map<map<T1, int>, bool> := map[v101 + v101 := p1];
				v102 := v102[v101 := p2];
				v64[630], v0 := v64[630], v0;
				var v103: array<seq<array<bool>>> := new seq<array<bool>>[10];
				var v104: seq<array<bool>> := [v64, v64, v64];
				v103[655] := v104;
				v103[655] := (v104 + v104) + v104;
			case DC3(cf2) =>
				var v105: C6 := new C6(|v0|, p0);
				var v106: seq<C6> := [v105, v105];
				var v107: map<int, int> := map[|v106| := v94];
				var v108: C4 := new C4(p2, v94, v94);
				var v109: map<C4, int> := map[v108 := v79[944][869]];
				var v110: C4 := new C4(p1, p0, if (v108 in v109) then v109[v108] else 0x145);
				v94 := ((if (|map[v64[630] := v110]| in v107) then v107[|map[v64[630] := v110]|] else |v0|) / v94) + (if (0x40 in v107) then v107[0x40] else -0x132);
				var v111 := new C2(v79[944][869], v79[944][869]);
				v64, globalState.f3, v79[944][869] := v64, (v66 + (seq(616, i4  => (v63)))) == v66, v94;
				var v112: array<C3> := new C3[19];
				var v113: C3 := new C3(p1, p0, p0);
				v112[994] := v113;
				v112[994] := v113;
		}
		
	}
	
	var v114: map<multiset<string>, bool> := map[multiset{v66} := !v64[630]];
	var v115: multiset<string> := multiset{v66};
	var v116: C0 := new C0(v80, p0, p0);
	var v117: map<int, C0> := map[p0 := v116];
	var v118: map<int, map<int, C0>> := map[p0 := v117];
	v114 := v114[v115 := p0 in v118];
	var v119: array<string> := new string[9] [v66, v66, v66, v66, v66, v66, v66, v66, v66];
	var v120 := DC35(v119);
	r0 := v120.cf48;
	var v121 := DC13(p2, v63);
	r1 := match v121 {
		case DC12(cf19) => v66 <= DC17("v").cf26
		case DC13(cf20, cf21) => cf20
		case DC11(cf18) => p2
		case DC14(cf22) => true
	};
}
trait T0 {
	const f7 : int
	const f8 : int
}

trait T1 extends T0 {
	function fm0(p0: bool, globalState: GlobalState): set<set<bool>> 
	function fm1(p0: set<bool>, p1: int, p2: int, p3: int, globalState: GlobalState): bool 
}

trait T2 extends T1 {
	function fm2(p0: bool, globalState: GlobalState): string 
	method m1(p0: seq<char>, globalState: GlobalState) 
	method m2(p0: int, p1: bool, p2: int, p3: bool, globalState: GlobalState) returns (r0: seq<bool>, r1: int, r2: bool) 
}

class C0 extends T1 {
	var f13 : array<int>
	constructor (f13 : array<int>, f7 : int, f8 : int) {
		this.f13 := f13;
		this.f7 := f7;
		this.f8 := f8;
	}
	
	function fm0(p0: bool, globalState: GlobalState): set<set<bool>> {
		{{false} + {false}, {false} - {true, !true}, {true, false, false}}
	}
	function fm1(p0: set<bool>, p1: int, p2: int, p3: int, globalState: GlobalState): bool {
		({"jg"} + {"q", "xbglft", seq(739, i0  => ('o')), seq(-0x333, i1  => ('b')), "gkesmfsv"}) >= ({"nup"} + (set v0 : string | v0 in (seq(620, i2  => (seq(0x21b, i3  => ('h'))))) :: (v0)))
	}
}

class C1 extends T0 {
	constructor (f7 : int, f8 : int) {
		this.f7 := f7;
		this.f8 := f8;
	}
	
	function fm18(globalState: GlobalState): string {
		"jof" + ("hno" + (seq(0x14, i0  => ('k'))))
	}
	function fm19(p0: bool, p1: bool, p2: int, p3: int, globalState: GlobalState): char {
		DC1('m').cf1
	}
	method m6(globalState: GlobalState) returns (r0: set<bool>, r1: int, r2: int) {
		var v0: array<bool> := new bool[6](i0 => f7 > f7);
		var v1 := false;
		v0[464] := v1;
		var v2 := DC0(v1);
		var v3: seq<bool> := [v1, v1, if (v1) then v1 else v1];
		var v4: set<int> := {0x135};
		var v5 := DC5(v1, f7, |v4|, f7, v1);
		globalState.f4, v0[464], v2, globalState.f4 := v1, v3[f8], match v5 {
			case DC5(cf4, cf5, cf6, cf7, cf8) => v2
			case DC4(cf3) => v2
		}, v2.cf0;
		var v6: map<int, int> := map[f7 := f7];
		var v7 := "tuyj";
		for i1 := if (fm20(v1, v0[464], true, globalState) in v6) then v6[fm20(v1, v0[464], true, globalState)] else |v7| to 0x1a6 {
			var v8: array<int> := new int[2];
			v8[310] := f7;
			var v9: array<seq<int>> := new seq<int>[27](i2 => [-f7, f7, if (f8 in v6) then v6[f8] else f8]);
			v8[310], v1, v9, r2 := f8, v0[464], v9, v5.cf7;
			var v10: C0 := new C0(v8, i1, i1);
			v10 := v10;
			var v11 := 'y';
			v11 := v11;
			var v12: map<char, bool> := map['w' := v1];
			var v13: map<map<char, bool>, int> := map[v12 := f8];
			if (v12 in v13) {
				globalState.f3 := v3[v5.cf5];
				globalState.f4, v8[310] := v1, v8[310];
				var v14: map<D0, int> := map[v2 := v8[310]];
				var v15: set<map<D0, int>> := {v14};
				var v16: seq<set<int>> := [v4];
				v8[310], v7, v15, v0[464], v4 := v8[310], v7, v15, v1, v16[f8 + 310];
				globalState.f3 := i1 <= f8;
				var v17: map<bool, int> := map[i1 >= f8 := v8[310]];
				v17 := v17;
			} else {
				v7 := v7 + v7;
				var v18 := DC1(v11);
				var v19: seq<D1> := [v18, v18, DC1('c')];
				var v20 := new C0(v8, |fm21(715, |v19|, f7, v7, globalState)|, f8);
				var v21: map<seq<D2>, bool> := map[[v5] := v1];
				var v22: seq<D2> := [v5];
				v21 := v21[v22 := false];
				v8[310] := v8[310] / (fm20(v0[464], DC0(v1).cf0, v0[464], globalState) / 0xe1);
				var v23 := DC1(v11);
				var v24 := DC3(v23);
				var v25: multiset<D1> := multiset{v24, DC3(v23), v24, v24};
				v25 := fm22(globalState);
			}
			
		}
		match (DC5(v3[f7], f7, |v4|, f7, v0[464])) {
			case DC5(cf4, cf5, cf6, cf7, cf8) =>
				var v26 := 'q';
				var v27: map<string, bool> := map[("xbuccmlho")[f8 := v26] := v0[464]];
				var v28: map<bool, bool> := map[v0[464] := v27 != map[seq(787, i3  => (v26)) := cf4]];
				v28 := v28[cf4 := cf8];
				var v29 := DC3(fm23(|"qf"|, globalState));
				var v30 := DC3(v29);
				var v31: seq<int> := [cf5];
				v30, cf4, v0[464], v7 := v30, v1, cf4, ((seq(-0x195, i4  => (v26))) + ((seq(0x25a, i5  => (v26))) + v7))[cf7 - |v31| := v26];
				cf4 := fm24(cf7 < v5.cf5, cf5, ['c'], globalState);
				cf5, v0[464], cf8 := f8, cf4 <== cf4, !v1;
			case DC4(cf3) =>
				var v32: array<int> := new int[6](i6 => i6 * f7);
				v32[161] := f8;
				var v33: set<bool> := {v0[464]};
				var v34: C0 := new C0(v32, |v7|, f7);
				var v35: map<bool, C0> := map[v1 := if (false) then v34 else v34];
				v0[464], r0, v32[161], r2 := !!(!(v0[464] || !v0[464]) && v1), if (v1) then v33 * v33 else v33, --|v35|, f7 + f7;
				if (-f7 != (f7 * -f8)) {
					v32[161] := 117;
					var v36: seq<multiset<bool>> := [cf3];
					var v37: array<multiset<bool>> := new multiset<bool>[11] [cf3[v3[f8] := v32[161]] + multiset([!v0[464], v1]), cf3, cf3, multiset(v3), cf3, cf3, cf3, cf3 + cf3, v36[-f8] + cf3, cf3, cf3];
					v37[908] := cf3;
					v37[908] := (cf3 * multiset{v1}) - cf3;
					v32[161] := f8 * (if (v1) then -v32[161] else v32[161]);
					var v38 := 'u';
					var v39: map<char, bool> := map[v38 := v0[464]];
					v39 := v39[v38 := v0[464]];
					v32[161] := v32[161];
				} else {
					r1 := if (v0[464]) then if (v1) then -f8 else f8 else v32[161];
					var v40: seq<D1> := [DC2()];
					v40 := v40;
					var v41 := DC4(cf3);
					v41 := v41;
					var v42 := new C0(v32, -810 / v32[161], v32[161] + |v7|);
					var v43: array<C0> := new C0[14];
					v43[650] := v34;
					var v44: map<int, bool> := map[f8 := v1];
					var v45: multiset<int> := multiset{fm20(v0[464], v1, if (f7 in v44) then v44[f7] else v0[464], globalState), |v33|, f8};
					v43[650], r1 := v42, |v45|;
				}
				
				v0[464] := v3[if (v1) then v32[161] else v32[161]];
				var v46: set<D0> := {v2};
				var v47: map<set<D0>, D2> := map[v46 := DC5(v34.fm1(v33, v32[161], v32[161], f7, globalState), fm20(v1, v0[464], v0[464], globalState), f8, -0x1a2, v1)];
				v32[161] := |v47|;
		}
		
		var v48 := 'x';
		var v49 := DC1(v48);
		var v50 := DC3(v49);
		match (v50) {
			case DC2() =>
				var v51: set<array<bool>> := {v0};
				if ((v51 >= v51) && v1) {
					globalState.f4 := v1;
					var v52 := DC1(if (true) then v48 else fm19(v1, v5.cf8, f8, f7, globalState));
					v52 := v52;
					var v53: array<int> := new int[9];
					var v54: map<int, array<bool>> := map[f7 := v0];
					var v55 := new C0(v53, -f7, |v54|);
					var v56: map<bool, bool> := map[v0[464] := 0x12b <= fm20(v1, v0[464], v0[464], globalState)];
					v56 := fm25(globalState);
					r2 := (if (v1) then f8 else f7) / f8;
				} else {
					var v57: map<int, bool> := map[f8 * fm20(v1, v1, v0[464], globalState) := v0[464]];
					v57 := v57[f7 := v0[464]];
					var v58: array<int> := new int[26];
					v58[299] := f8;
					v58[299] := f7;
					var v59: seq<int> := [v58[299]];
					var v60: multiset<int> := multiset{v58[299]};
					var v61: map<seq<int>, bool> := map[v59[fm20(fm24(v1, |"trxrrblv"|, [v48, v48], globalState), v0[464], v1, globalState) := f8] := (fm26(v60, v0[464], globalState)).cf9 > v4];
					v61 := v61;
					r2 := f8;
					v1 := v3[v58[299]];
				}
				
				r1 := v5.cf5;
				globalState.f4 := !v1;
				var v62: seq<int> := [f7, f8];
				r2 := -476 - v62[fm20(v0[464], v1, false, globalState)];
			case DC1(cf1) =>
				var v63: array<int> := new int[27](i7 => i7 - f7);
				var v64: map<bool, int> := map[v1 := -f7];
				v63[300] := |v64| * v5.cf7;
				v63[300] := -0x375;
				var v65: multiset<int> := multiset{v63[300], f7};
				globalState.f4 := v65 == v65;
				var v66 := DC7(v0[464]);
				var v67: map<bool, D3> := map[v1 := v66];
				v67 := v67[if (v1) then v0[464] else v0[464] := v66];
				if (v0[464]) {
					globalState.f4 := v0[464];
					cf1 := cf1;
					globalState.f4 := (0x3bd - 0x15e) == f8;
					var v68: array<array<int>> := new array<int>[17] [v63, v63, v63, v63, v63, v63, v63, v63, v63, v63, v63, v63, v63, v63, v63, v63, v63];
					var v69: seq<array<array<int>>> := [v68, v68];
					var v70: map<bool, bool> := map[v0[464] := v1];
					var v71: array<array<array<int>>> := new array<array<int>>[4] [v68, v69[f8], v68, v69[|v70|]];
					v71[830] := v68;
					var v72 := DC8(v68);
					v71[830] := v72.cf11;
					var v73: map<int, bool> := map[v63[300] := true];
					v63[300] := (|[v1]| - fm20(v1, v1, v0[464], globalState)) % (f8 / |v73|);
				} else {
					var v74: array<string> := new string[7] [v7, v7, fm18(globalState), v7, v7, fm18(globalState), v7];
					var v75: map<bool, array<string>> := map[v1 := v74];
					v75 := v75[v1 := v74];
					var v76: seq<multiset<bool>> := [multiset(v3)];
					var v78: map<bool, bool> := map[v1 := v5.cf8];
					var v79: multiset<bool> := multiset{if (v0[464] in v78) then v78[v0[464]] else v0[464]};
					var v80: set<multiset<bool>> := {v79};
					v0[464] := (if (fm24(v1, f8, v7, globalState)) then set v77 : multiset<bool> | v77 in v76 :: (v77) else fm27(globalState)) != v80;
					globalState.f3 := (v3 + v3) != v3;
					globalState.f4, r1 := v0[464], f8;
					v6 := v6[v63[300] := -f7] + v6;
				}
				
			case DC3(cf2) =>
				var v81: multiset<bool> := multiset{v0[464]};
				var v82: seq<multiset<bool>> := [v81, v81, multiset{!false, false}];
				var v83: multiset<int> := multiset{f7};
				var v84: array<int> := new int[24] [|(seq(970, i8  => (f7)))|, f8, f7, f7, f8, 0x1e7, f8, f8, |v7|, f7, f7, 0x390, |{f8, v5.cf5, fm20(v1, v1, v0[464], globalState)}|, f7, |v82[f8 := v81]|, f7, |v6|, f7, |v7[f8 := v48]|, |v83|, 0x261, f8, f8, f8];
				var v85: map<bool, char> := map[v1 := 'l'];
				var v86 := new C0(v84, if (f7 in v83) then v83[f7] else |v85|, f8);
				v0[464], v0[464], globalState.f3, v7 := fm24(v0[464], f8, v7, globalState), v1, fm20(v0[464], v0[464], true, globalState) == f7, seq(0x5a, i9  => (v48));
				var v87: array<seq<int>> := new seq<int>[25](i10 => [|{DC2(), DC2(), DC2(), DC2(), DC2()}|, f8, |[v1, v0[464], v1]|, f8, |v83|]);
				var v88: seq<array<bool>> := [v0];
				v87[487] := fm21(fm20(v0[464], v0[464], true, globalState), f7, |v88[f7 := v0]|, v7, globalState);
				var v89: map<bool, int> := map[v0[464] := f8];
				var v90: seq<int> := [-|v89|];
				v87[487] := [|v3|, f7] + v90;
				if (v1) {
					var v91 := DC4(fm28(v0[464], globalState) - multiset{v0[464]});
					v91 := v91;
					v86.f13[799] := -f8;
					v86.f13[799] := |{v1, v0[464]}|;
					var v92: multiset<map<int, int>> := multiset{v6, v6, v6};
					v0[464] := v92 > v92;
					v81 := v81;
					var v93: map<int, bool> := map[v86.f13[799] := v1];
					var v94: array<array<int>> := new array<int>[5];
					v94[675] := v86.f13;
					v86.f13[799], v93, r1, r1, v94[675] := if ((v0[464] && v1) in v89) then v89[v0[464] && v1] else f7 % v86.f13[799], v93, f7, f8 % f8, v86.f13;
				} else {
					r2 := (f8 % |map[v86.f13 := v1]|) + f8;
					var v95: map<bool, bool> := map[v0[464] := |multiset(v87[487][|DC11(v89).cf18| := f7])| <= |{f7}|];
					var v96: seq<map<bool, bool>> := [v95, v95, v95, v95 + v95];
					v95 := v96[f7];
					var v97: array<string> := new string[21];
					var v98: multiset<char> := multiset{v48};
					globalState.f4, v97, v0[464] := v98 > v98, v97, v1;
					r2 := f8;
					r1 := f8;
				}
				
		}
		
		var v99 := DC10(f7, v0[464]);
		r2 := |(match v99 {
			case DC9(cf12, cf13, cf14, cf15) => "cradsw"
			case DC10(cf16, cf17) => "bqqe"
			case DC8(cf11) => v7
		})|;
		globalState.f4 := (if (v0[464]) then v1 else v3[f7]) && v1;
		var v100: set<bool> := {v1};
		r0 := v100 * (v100 * v100);
		r1 := f8;
		var v101: multiset<array<bool>> := multiset{v0};
		r2 := f8 + (if (false) then |v101| else f7);
	}
}

class C2 extends T0, T2 {
	constructor (f7 : int, f8 : int) {
		this.f7 := f7;
		this.f8 := f8;
	}
	
	function fm2(p0: bool, globalState: GlobalState): string {
		"ikasy"
	}
	function fm0(p0: bool, globalState: GlobalState): set<set<bool>> {
		{{false, false, true, false} - {false}, {false, true} - {true, false, true, true}, {!false, true} * {DC0(false).cf0}}
	}
	function fm1(p0: set<bool>, p1: int, p2: int, p3: int, globalState: GlobalState): bool {
		(if (false) then DC0(false) else DC0(true)).cf0
	}
	function fm14(globalState: GlobalState): string {
		"qx" + "cv"
	}
	function fm15(globalState: GlobalState): bool {
		false
	}
	method m1(p0: seq<char>, globalState: GlobalState) {
		var v0: array<bool> := new bool[19](i0 => f8 < f8);
		v0 := v0;
		var v1 := false;
		var v2 := DC5(v1, f8, f7, f7, v1);
		for i1 := f7 to v2.cf5 {
			var v3 := "rsl";
			v3 := p0 + p0;
			var v4: array<array<bool>> := new array<bool>[1];
			v4 := new array<bool>[10];
			var v5: array<int> := new int[7](i2 => i2 + i1);
			var v6: map<int, int> := map[0x177 := f7];
			var v7 := new C0(v5, |(v3 + p0)|, |v6|);
			if (0x35c < |fm2(v1, globalState)|) {
				var v8 := -417;
				v8 := -0x25a - f8;
				v8 := v8;
				globalState.f4 := v1;
				globalState.f3 := if (v1) then v1 else i1 >= -0x3c5;
				v0[310] := v1;
				v0[310] := !v1;
			} else {
				var v9: multiset<int> := multiset{f8};
				var v10: seq<multiset<int>> := [v9, v9, v9];
				var v11: map<array<bool>, int> := map[v0 := f7];
				var v12: map<int, map<int, int>> := map[|v11| := v6];
				var v13: map<int, map<int, map<int, int>>> := map[|v10| := v12 + v12];
				v13 := v13[fm16(v1, globalState) := v12[i1 := v6]];
				var v14: seq<int> := [f8];
				var v15: seq<seq<int>> := [v14];
				var v16: multiset<bool> := multiset{v1, v1, !v1, v15[v2.cf5] == v14};
				var v17 := 'a';
				var v18: map<bool, int> := map[v1 := |v3[-i1 := v17]|];
				var v19: seq<bool> := [true, !!v1, v1];
				globalState.f3, v16, v18, v19 := false, (v16 + v16) + v16, v18 + v18, v19;
				var v20 := new C0(v7.f13, f8, f7);
				var v21: array<char> := new char[2];
				v21[361] := p0[fm16(false, globalState)];
				v21[361] := v17;
				var v22 := new C0(if (v1) then v7.f13 else v5, i1 % |(seq(0x11e, i3  => (v21[361])))|, fm16(v1, globalState));
			}
			
		}
		var v24: multiset<bool> := multiset{v1, v1};
		var v25: set<int> := {|v24|, f8};
		var v26: array<set<int>> := new set<int>[2] [set v23 : int | (695 <= v23) && (v23 < -840) :: (v23 * -f7), v25];
		var v27: map<int, array<set<int>>> := map[---f8 := v26];
		var v28: array<array<set<int>>> := new array<set<int>>[11] [if (v1) then v26 else v26, if (f8 in v27) then v27[f8] else v26, v26, v26, v26, v26, v26, v26, v26, v26, v26];
		v28[734] := v26;
		v28[734] := v26;
		var v29: multiset<int> := multiset{0x2cb, -f7};
		var v30: map<bool, bool> := map[v1 := v1];
		var v31: array<int> := new int[25] [|map[|v29| := p0[|v24| := 'y']]|, f8, |v25|, 285, f7, f8, f8, f8, |p0|, f7, f8, f8, 891, fm16(v1, globalState), -f7, f8, f7, |p0|, -|v30|, f8, f7, f7, f7, f7, |p0|];
		var v32 := new C0(v31, f7, f7);
		globalState.f4 := v1;
		var v33 := new C0(v32.f13, -f8, |(multiset{"trnq", p0, p0, p0, seq(0x35, i4  => ('o'))})[p0 := f7]|);
	}
	method m2(p0: int, p1: bool, p2: int, p3: bool, globalState: GlobalState) returns (r0: seq<bool>, r1: int, r2: bool) {
		var v0: array<int> := new int[2](i0 => i0 / f8);
		var v1 := new C0(v0, p0 % f8, 0x2fd);
		globalState.f4 := p1;
		var v2: seq<bool> := [true, p3];
		var v3 := DC4(multiset(v2));
		match (v3) {
			case DC5(cf4, cf5, cf6, cf7, cf8) =>
				if (!(f8 < cf5)) {
					globalState.f4 := cf4;
					r2 := if (p3) then cf8 else if (cf8) then cf8 else false;
					globalState.f3 := p1;
					var v4 := 'c';
					var v5: map<bool, char> := map[if (true) then p3 else false := v4];
					v5 := v5[p1 := 'e'];
					var v6 := "ksccjtf";
					v6 := v6;
				} else {
					cf6 := |fm14(globalState)| % cf7;
					var v7 := 'g';
					v7, r0 := v7, v2;
					var v8: array<string> := new string[10];
					var v9 := "ucud";
					v8[902] := v9[448 := v7];
					v8[902] := v9 + (seq(-60, i1  => (v7)));
					r0 := [p3] + (v2 + [p3]);
					var v10: array<bool> := new bool[13];
					v10[272] := cf8;
					var v11: seq<int> := [cf6];
					var v12: map<bool, seq<int>> := map[cf4 := v11];
					v10[272], r2, cf6, v11, v12 := p3, p1, if (p1) then cf6 else cf6, v11, fm17(v7, cf4, globalState);
				}
				
				if (p1) {
					var v13 := "v";
					v13 := (seq(-0x210, i2  => ('d'))) + v13;
					var v14: map<int, string> := map[f7 := "lb"];
					v14 := v14[cf6 := v13];
					var v15: T0 := new C1(f7, f8);
					v15 := if (p1) then v15 else v15;
					var v16: set<bool> := {cf8};
					var v17: map<int, int> := map[v15.f7 := 0x341];
					globalState.f3 := v1.fm1(v16, f8, -|v17|, v15.f8, globalState);
					var v18: map<bool, int> := map[p3 := v15.f8];
					v18 := v18[p1 := |(map v19 : int | (0x1b5 <= v19) && (v19 < -935) :: (v19 % cf5) := (p3))| - p0];
				} else {
					var v20: map<bool, bool> := map[!p3 := cf5 < p2];
					v20 := v20[cf8 := cf8];
					globalState.f4 := cf6 > (0x393 % cf7);
					var v21 := 'v';
					var v22: seq<char> := [v21];
					v1.f13[904] := |(v22 + ['j'])|;
					var v23: map<bool, int> := map[!cf8 := p0];
					var v24: multiset<map<bool, int>> := multiset{map[p3 := |v22|]};
					v1.f13[904] := if (cf8) then if (p3 in v23) then v23[p3] else |v24| else f7 - p2;
					v20 := v20[false := |v22| == cf7];
					var v25: seq<int> := [cf5, cf7];
					var v26: multiset<seq<int>> := multiset{v25};
					v26 := multiset{v25, v25};
				}
				
				cf5 := cf7;
				globalState.f4 := cf4;
			case DC4(cf3) =>
				r1 := f8;
				var v27 := 'y';
				var v28: set<char> := {v27, v27, 'd'};
				var v29: seq<set<char>> := [v28];
				globalState.f4 := p2 >= |v29|;
				var v30: array<bool> := new bool[18];
				var v31: seq<array<bool>> := [v30];
				var v32 := DC9(p0, v31, f7, p0);
				var v33: array<array<int>> := new array<int>[3];
				var v34 := DC8(v33);
				var v35: map<bool, D4> := map[p1 := v34];
				match (DC5(true, f8, 475, v32.cf14, p3).(cf8 := false, cf5 := p0).(cf6 := --|v35|)) {
					case DC5(cf4, cf5, cf6, cf7, cf8) =>
						var v36: map<bool, array<int>> := map[!cf8 := v1.f13];
						var v37: set<int> := {|v36|};
						var v38 := DC6(v37 + v37);
						v38 := v38;
						globalState.f4 := p3;
						var v39: set<bool> := {p1};
						v30[84] := true;
						var v40: seq<set<bool>> := [v39, v39, {v2[f7]}];
						v39, v30[84] := v40[|v2| + cf7], if (p3) then f8 == p0 else p1;
						cf6 := f8;
					case DC4(cf3) =>
						globalState.f4 := f7 > p2;
						r1 := --0x17c;
						var v41: array<set<int>> := new set<int>[5];
						var v44: set<int> := {p0};
						v41[757] := (set v43 : int | v43 in (set v42 : int | (-205 <= v42) && (v42 < 0x2e) :: (v42 / |"vudgnl"|)) :: (v43 / |(seq(0x1b0, i3  => ('o')))|)) - v44;
						var v46 := DC6({p2});
						r2, r0, r1, v32, v41[757] := (v28 * (set v45 : char | v45 in (seq(0x3da, i4  => (v27))) :: (v45))) >= v28, v2, f7, v32.(cf14 := p2 - |v31|), v46.cf9;
						v30 := v30;
				}
				
				var v47: multiset<int> := multiset{p0, p2};
				var v48 := new C1(--fm16(true, globalState), -|fm29(p1, v47, globalState)|);
		}
		
		var v49: multiset<bool> := multiset{p1};
		var v50: seq<int> := [fm20(p3, !false, p3, globalState), if (p3 in v49) then v49[p3] else -f7];
		for i5 := |(if (p3) then v50 else v50)| to p0 {
			globalState.f3 := p1;
			var v51 := "nvldbus";
			var v52: map<bool, int> := map[p3 := |v51|];
			var v53: multiset<int> := multiset{p0};
			v52 := v52[p3 := if (p2 in v53) then v53[p2] else p0];
			var v54: array<bool> := new bool[13](i6 => p3);
			var v55 := DC9(p2, [v54, v54, v54, v54, v54], 0x2b0, f8);
			v1.f13[503] := (if (p1 in v52) then v52[p1] else p2) / v55.cf14;
			v1.f13[503] := p0 * f7;
			var v56: map<int, array<int>> := map[|v51| := v0];
			var v57: map<map<int, array<int>>, multiset<int>> := map[v56 := v53];
			v57 := v57[v56 := multiset(if (p1) then v50 else [|v53|, f7, -0x2cd, 620, 742])];
		}
		r1 := fm16(fm15(globalState), globalState);
		forall i7 | 0 <= i7 < v0.Length {
			v0[i7] := i7 % f8;
		}
		var v58: set<bool> := {p3};
		r0 := v2[-f7 := f7 != |v58|];
		var v59: set<int> := {p0};
		var v60 := DC6(v59);
		r1 := match v60 {
			case DC7(cf10) => if (cf10) then p0 else |multiset{p2, -0x328, f8, |"fjemwxaw"|, f8}|
			case DC6(cf9) => 0x14c
		};
		var v61 := "rtyih";
		r2 := fm20(false, p3, p1, globalState) != |v61|;
	}
}

class C3 extends T2, T0 {
	const f12 : bool
	constructor (f12 : bool, f7 : int, f8 : int) {
		this.f12 := f12;
		this.f7 := f7;
		this.f8 := f8;
	}
	
	function fm2(p0: bool, globalState: GlobalState): string {
		"skp"
	}
	function fm0(p0: bool, globalState: GlobalState): set<set<bool>> {
		match DC0(f12) {
			case DC0(cf0) => if (f12) then {{false, cf0}} else {{cf0}, {f12}}
		}
	}
	function fm1(p0: set<bool>, p1: int, p2: int, p3: int, globalState: GlobalState): bool {
		(|{f12, f12, f12, f12, !f12}| % |(map v0 : int | (0xea <= v0) && (v0 < 0x156) :: (v0 - 0x38a) := (f12))|) >= (|multiset([f8, |"xutwsx"|, f8])| - f8)
	}
	function fm12(globalState: GlobalState): bool {
		f12
	}
	function fm13(p0: string, globalState: GlobalState): int {
		f8
	}
	method m1(p0: seq<char>, globalState: GlobalState) {
		var v0: seq<bool> := [f12, !f12, f12];
		for i0 := -|(v0 + v0)| to -f7 {
			var v1 := DC0(f12);
			if (!(i0 > |multiset{v1}|)) {
				var v2: array<int> := new int[22](i1 => i1 % i0);
				var v3: multiset<int> := multiset{f7};
				var v4: map<array<int>, multiset<int>> := map[v2 := v3];
				var v5: array<bool> := new bool[29](i2 => multiset(seq(0x2df, i3  => ('b'))) != multiset{'e'});
				v5[259] := f12;
				var v6: set<int> := {f7, f8};
				var v7: seq<int> := [|v6|];
				v4, globalState.f3, globalState.f4, v5[259] := v4 + v4, (v7 + [f7]) != (v7 + v7), f12, f12;
				var v8 := 'r';
				var v9 := DC1(DC1('j').cf1);
				v8 := v9.cf1;
				globalState.f3 := v5[259];
				globalState.f3 := v0[i0];
				v5[259] := f12;
			} else {
				globalState.f4 := f12 && (if (f12) then f12 else f12);
				var v10: multiset<seq<bool>> := multiset{v0};
				globalState.f3 := (v10 - v10) > v10;
				var v11: array<bool> := new bool[28](i4 => f12);
				v11 := v11;
				var v12: seq<seq<bool>> := [v0];
				v12 := v12;
				var v13 := 658;
				v13 := f7 + (f8 % f7);
			}
			
			var v14 := 'g';
			v14 := v14;
			globalState.f4 := f12;
			var v15: map<int, int> := map[f8 := i0 % |multiset{f12}|];
			v15 := v15;
		}
		var v16: multiset<int> := multiset{f8};
		var v17, v18 := m0(f7, fm1({f12, f12}, -0x232, f7, |map[|{0x37d}| := f8]|, globalState), v16 >= v16, globalState);
		var v19 := 'b';
		var v20: map<bool, char> := map[fm12(globalState) := v19];
		v20 := v20;
		var i5 := 0;
		while (f12)
			decreases 100 - i5
		{
			if (i5 >= 100) {
				break;
			}
			
			i5 := i5 + 1;
			var v21 := "xronwep";
			v21, v18 := p0, f12;
			v18 := f12;
			var v22 := 0x303;
			var v23: array<bool> := new bool[1];
			var v24 := DC1(v19);
			var v25: multiset<D1> := multiset{v24};
			var v26: seq<multiset<D1>> := [v25];
			v22, v23, v19, v21, v26 := -(0x283 % 663), v23, v19, p0 + v21, (seq(0x2fb, i6  => (multiset([v24, v24])))) + v26;
			var v27: map<bool, bool> := map[v18 := v18];
			var v28: map<map<bool, bool>, array<bool>> := map[v27 + v27 := v23];
			v28 := v28[v27 + v27 := v23];
		}
		var v29: multiset<bool> := multiset{false};
		var v30: seq<int> := [|v29|];
		for i7 := v30[f8] to f7 * f7 {
			var v31, v32 := m0(f7 + -|multiset{"vvjdbcbj", seq(-0x1b3, i8  => (v19))}|, v18, v18 <==> true, globalState);
			globalState.f3 := DC0(v32).cf0;
			v0 := v0 + (v0 + v0);
			v30, globalState.f3 := v30, f12;
		}
		var v33: seq<seq<bool>> := [if (v18) then v0 else v0];
		v0 := v33[f8 / f7];
	}
	method m2(p0: int, p1: bool, p2: int, p3: bool, globalState: GlobalState) returns (r0: seq<bool>, r1: int, r2: bool) {
		var v0 := "dwwkgew";
		r1 := 0x10b / |v0|;
		var v1 := 'm';
		var v2: array<bool> := new bool[20](i0 => -p2 > f8);
		v1, v2, r1, r1 := 'y', v2, -|v0|, 0xff;
		r2 := p3;
		var v3: seq<bool> := [f12, true, p1];
		var v4: multiset<bool> := multiset{!p1, v3[p2]};
		var v5 := DC4(v4);
		v4 := (v5.(cf3 := v4)).cf3;
		for i1 := p2 to -f7 {
			r1 := -((p0 % -574) % 0x1bc);
			var v6 := DC0(f12);
			match (v6) {
				case DC0(cf0) =>
					r1 := p0;
					r1 := f8;
					var v7: map<int, int> := map[f8 := fm13(v0, globalState)];
					r1 := --(|v7| % |([p1] + v3)|);
					var v8: T2 := new C2(p2, p2);
					v8 := v8;
			}
			
			var v9 := DC12(!fm12(globalState));
			var v10: multiset<int> := multiset{-0x2bf, p0, |v0|};
			var v11: seq<int> := [i1];
			r2, globalState.f3, v9, v10 := f7 < |fm30(p3, globalState)|, !(p3 || f12), v9, v10 - (v10 + multiset(v11));
			globalState.f4 := !(if (p1) then false else !p1);
		}
		var v12: array<int> := new int[24](i3 => i3 % |v3[f7 := p1]|);
		forall i2 | 0 <= i2 < v12.Length {
			v12[i2] := i2 % (if (f12) then p2 else -f8);
		}
		r0 := v3;
		r1 := p2 * f7;
		r2 := p3;
	}
}

class C4 extends T1 {
	const f11 : bool
	constructor (f11 : bool, f7 : int, f8 : int) {
		this.f11 := f11;
		this.f7 := f7;
		this.f8 := f8;
	}
	
	function fm0(p0: bool, globalState: GlobalState): set<set<bool>> {
		({{f11}} + {{f11, f11}, {f11, false, false, f11}}) + ({{f11}} + (set v1 : set<bool> | v1 in (set v0 : set<bool> | v0 in multiset{{f11}, {f11, f11, f11}} :: (v0)) :: (v1)))
	}
	function fm1(p0: set<bool>, p1: int, p2: int, p3: int, globalState: GlobalState): bool {
		|(seq(0x1e0, i0  => ('g')))| < f7
	}
	function fm11(globalState: GlobalState): int {
		if (f11) then f8 else |([f7, f7] + (seq(0x2be, i0  => (0x1c5))))|
	}
	method m5(p0: bool, p1: int, p2: int, p3: int, globalState: GlobalState) returns (r0: bool, r1: bool, r2: bool, r3: bool) {
		r1 := p0;
		var v0 := new C2(0x219, -183);
		var v1 := new C2(p2, f7 / f7);
		var i0 := 0;
		while (if (f11) then p0 else p0)
			decreases 100 - i0
		{
			if (i0 >= 100) {
				break;
			}
			
			i0 := i0 + 1;
			var v2 := 0xf1;
			v2 := p2;
			var v3 := new C1(f7 % 83, f7 % fm16(true, globalState));
			var v4: array<array<int>> := new array<int>[9];
			v4 := v4;
			var v5: map<bool, int> := map[p0 := p2];
			var v6: map<map<bool, int>, seq<bool>> := map[v5 := [p0, f11]];
			v2 := -p2 / (v2 / |v6|);
		}
		match (fm26(multiset{p2}, f11, globalState)) {
			case DC7(cf10) =>
				var v7 := 0xb8;
				var v8: seq<bool> := [f11, false];
				var v9: seq<seq<bool>> := [v8, v8];
				v7 := (-f8 % v7) * |v9[|"lhbhuxmc"|]|;
				var v10: multiset<int> := multiset{v7 * p3};
				var v11: seq<int> := [p1];
				v7 := -(if (p2 in v10) then v10[p2] else |(v11 + v11)|);
				v11 := v11;
				var v12 := "md";
				var v13: array<int> := new int[10] [f8, p1, f8, -318, p2, v7, |v12|, v11[v7], f8, p3];
				var v14: array<array<int>> := new array<int>[1] [v13];
				v14[304] := v13;
				v14[304] := v13;
			case DC6(cf9) =>
				var v15 := "u";
				var v16: set<string> := {v15};
				var v17: array<bool> := new bool[25](i1 => p0);
				var v18: map<string, array<bool>> := map[v15 := v17];
				var v19 := -109;
				var v20: seq<map<string, array<bool>>> := [map["ekpn" := v17], v18];
				var v21: set<bool> := {p0};
				var v22: multiset<bool> := multiset{f11};
				v16, v18, v19 := v16 - v16, v20[|v21|], |(v22 * v22)|;
				r3 := !f11;
				var v23: seq<bool> := [p0, p0];
				var v24: seq<int> := [0x50, f7, p1];
				globalState.f4 := (v23 + v23)[|v24| := f11] == v23;
				v23 := v23[607 := p0];
		}
		
		var v25 := 321;
		v25 := p2;
		r0 := p0;
		r1 := f11;
		r2 := p0 ==> f11;
		r3 := !f11;
	}
}

class C5 extends T0, T1 {
	const f9 : string
	const f10 : int
	constructor (f9 : string, f10 : int, f7 : int, f8 : int) {
		this.f9 := f9;
		this.f10 := f10;
		this.f7 := f7;
		this.f8 := f8;
	}
	
	function fm0(p0: bool, globalState: GlobalState): set<set<bool>> {
		({{false, false}} - {{true, true}, {false, true}}) + {{false}, {true, false, !true, false, true}}
	}
	function fm1(p0: set<bool>, p1: int, p2: int, p3: int, globalState: GlobalState): bool {
		false && ([{|[f7]|, f10, f8, |multiset([f10, f10])|, |f9|}, {f10}] != [{0x40}])
	}
	function fm6(p0: char, p1: string, p2: bool, globalState: GlobalState): int {
		f8
	}
	function fm7(p0: int, globalState: GlobalState): int {
		f10 * f7
	}
	method m4(p0: int, globalState: GlobalState) returns (r0: int) {
		var v0: array<int> := new int[11](i0 => i0 * f7);
		v0[993] := f7;
		var v1 := false;
		v0[993] := fm6('a', fm8(v1, [f8, p0, p0], v1, globalState), v1, globalState);
		var v2: seq<bool> := [v1, v1, v1];
		var v3: map<string, bool> := map[f9 := v1];
		var v4 := 'g';
		var v5: map<int, int> := map[p0 := |f9|];
		var v6 := DC0(v1);
		var v7: set<int> := {f8, |f9|};
		var v8: map<bool, bool> := map[v1 := false];
		var v9: multiset<map<bool, bool>> := multiset{fm9(v1, globalState)};
		var v10: set<bool> := {v1};
		var v11: multiset<bool> := multiset{v1, v1};
		var v12: seq<int> := [f7];
		var v13: map<multiset<int>, bool> := map[((multiset([|v12|]))[p0 := f7])[p0 := f8] := v1];
		var v14: array<bool> := new bool[27] [v1, v1, v2[f8], if ("twaaoc" in v3) then v3["twaaoc"] else v1, v1, v1, v1, v1, v1, -fm6(v4, f9, v1, globalState) < -|multiset{v1}|, !v1, v1, |v2| !in v5, v6.cf0, v1, v1, v7 <= v7, v8 !in v9, v1, v1, v1 || v1, !(v1 && v1), v10 !! v10, v11 !! multiset{v1, v1}, false, v1, if (multiset{v0[993], v0[993], p0, f7} in v13) then v13[multiset{v0[993], v0[993], p0, f7}] else !v1];
		v14[888] := if (v1) then v1 else !v1;
		var v15: multiset<int> := multiset{f8};
		v0[993], v14[888] := fm7(|v15[v0[993] := f7]| % p0, globalState), v2[f7];
		var v16: seq<array<int>> := [v0, v0];
		v0[993] := |(v16 + v16)[|map[v14[888] := v14[888]]| + f8 := v0]|;
		if (v1) {
			var v17: map<char, seq<D0>> := map[v4 := (fm10(f9, v14[888], -p0, f7, globalState))[f7 := DC0(v1)]];
			var v18: seq<D0> := [v6, v6];
			v17 := v17[v4 := v18];
			v1 := if (v14[888] in v8) then v8[v14[888]] else !v14[888];
			v14[888] := v1;
			var v19: map<bool, array<bool>> := map[v1 := v14];
			var v20: array<array<bool>> := new array<bool>[15] [v14, v14, v14, v14, if (v1 in v19) then v19[v1] else v14, v14, v14, v14, v14, v14, v14, v14, if (true in v19) then v19[true] else v14, v14, v14];
			v20[888] := v14;
			v20[888] := new bool[23];
			v14 := v20[888];
		} else {
			var v22: set<string> := {f9};
			r0, v14[888] := f7, (set v21 : string | v21 in [f9[v0[993] := 'f'], f9] :: (v21)) > v22;
			var v23 := new C3("xs" <= f9, f7, v0[993]);
			var v24 := new C1(f7 + f7, -|v12| - f7);
			if (v1) {
				v15 := multiset{-599, f8, fm7(v0[993], globalState)};
				var v25: C0 := new C0(v0, 0x3df, f10);
				var v26: seq<C0> := [v25];
				var v27: set<seq<C0>> := {v26};
				globalState.f3 := (-|v27| / p0) >= (fm16(v1, globalState) % v0[993]);
				v0[993] := v0[993] % 0x3c5;
				v0[993] := 0x25a;
				var v28 := new C4(v23.f12, f10, p0);
			} else {
				v12 := v12 + v12;
				var v29 := new C1(f8, -fm7(f8, globalState));
				var v30: map<bool, int> := map[v23.f12 := |v2|];
				globalState.f4 := v30[v23.f12 := f10] != v30;
				globalState.f4 := v14[888];
				v0[993] := f10;
			}
			
			var v31: map<multiset<int>, string> := map[v15 + v15 := "fvhdv"];
			v31 := v31[multiset(v12) := f9 + f9];
		}
		
		var v32: C4 := new C4(v14[888], v0[993], f10);
		var v33: seq<C4> := [v32, v32];
		v33 := v33;
		for i1 := p0 to f7 {
			var v34 := "tjdnlna";
			var v35: array<array<bool>> := new array<bool>[14];
			v35[690] := v14;
			var v36: array<char> := new char[2] ['f', v4];
			v36[652] := v4;
			var v37: map<seq<int>, bool> := map[[p0, |v15|] := v14[888]];
			v34, v0[993], globalState.f3, v35[690], v36[652] := "olimw", |(v37 + v37)|, f8 < (f8 % f8), v14, v4;
			var v38 := DC7(v32.f11);
			match (v38) {
				case DC7(cf10) =>
					var v39: set<map<int, int>> := {v5, v5, v5, v5, v5};
					var v41: array<set<map<int, int>>> := new set<map<int, int>>[15] [fm31(f8, multiset(fm30(cf10, globalState)), globalState), v39, set v40 : map<int, int> | v40 in (seq(0x182, i2  => (v5[|[i1]| := p0]))) :: (v40), v39, v39, v39, v39, v39 * v39, v39, fm31(v0[993], v11, globalState) - v39, {map[f8 := -f10]}, v39, v39, v39, fm31(f7, multiset(v2), globalState)];
					v41[196] := v39;
					var v42: seq<set<map<int, int>>> := [v39 * v39];
					v34, v1, v41[196], v0[993], v36[652] := f9, f8 < f10, v42[-f10], f7 + (|"cniy"| % |f9[-f8 := v36[652]]|), v4;
					v0 := v0;
					v14[888], v0, v12, cf10, v9 := true, v0, [v12[|v15|]], v32.f11 <==> v14[888], v9;
					v3 := v3[f9 := v14[888]];
				case DC6(cf9) =>
					globalState.f4 := v1;
					var v43: map<set<int>, string> := map[v7 := f9 + v34];
					v34 := if ((cf9 + cf9) in v43) then v43[cf9 + cf9] else v34;
					v0[993] := i1;
					globalState.f4 := v32.f11;
			}
			
			var v44 := DC1('u');
			var v45 := DC3(v44);
			var v46, v47, v48, v49 := v32.m5(v32.f11, |v12|, 0x2a5, |{v45}|, globalState);
			r0 := |(v8 + v8)| / v0[993];
		}
		r0 := f8;
	}
}

class C6 extends T2 {
	constructor (f7 : int, f8 : int) {
		this.f7 := f7;
		this.f8 := f8;
	}
	
	function fm2(p0: bool, globalState: GlobalState): string {
		("lysobld" + (seq(0x2dc, i0  => ('x')))) + "exbusrxu"
	}
	function fm0(p0: bool, globalState: GlobalState): set<set<bool>> {
		{{false, !!true, !false}} * ({{true, false}} - {{false}, {false, true, false, true, true}})
	}
	function fm1(p0: set<bool>, p1: int, p2: int, p3: int, globalState: GlobalState): bool {
		970 <= f7
	}
	function fm3(p0: bool, p1: bool, p2: bool, p3: bool, globalState: GlobalState): int {
		if (DC0(!false).cf0) then -0x179 else |"sv"|
	}
	function fm4(p0: bool, p1: bool, globalState: GlobalState): int {
		f7
	}
	method m1(p0: seq<char>, globalState: GlobalState) {
		var v0 := -0x1df;
		var v1 := "hndygtt";
		var v2 := true;
		var v3: multiset<bool> := multiset{v2};
		globalState.f3, v0, v1 := v2, if (v2) then f8 * (if (v2 in v3) then v3[v2] else f8) else f7 + f8, v1[f7 / v0 := fm5(globalState)];
		var v4: array<int> := new int[28];
		var v5: multiset<int> := multiset{|v1|};
		var v6: map<bool, multiset<int>> := map[v2 := v5];
		var v7: seq<int> := [-|(if (v2 in v6) then v6[v2] else v5)|, f7];
		var v8: T1 := new C0(v4, f8, f8 % v7[f7]);
		v8 := v8;
		var v9: seq<multiset<int>> := [v5[v0 := v0]];
		var i0 := 0;
		while (v9 < [v5, v5])
			decreases 100 - i0
		{
			if (i0 >= 100) {
				break;
			}
			
			i0 := i0 + 1;
			if (if (v2) then p0 == (seq(0x3aa, i1  => ('e'))) else if (v2) then v2 else v2) {
				var v10 := 'r';
				v1 := (v1 + v1)[v8.f7 / f7 := v10];
				var v11: array<bool> := new bool[26](i2 => p0 == v1);
				var v12 := DC13(true, v10);
				var v13 := DC15(v11);
				v11, v12 := v13.cf23, DC13(v2, v10);
				var v14: array<D2> := new D2[9];
				var v15: map<bool, array<D2>> := map[v2 := v14];
				v14 := if (v2 in v15) then v15[v2] else v14;
				var v16: set<char> := {v10, v10};
				v16 := v16 + v16;
				var v17: set<string> := {fm8(v2, v7, v2, globalState), p0 + p0};
				var v18: map<string, seq<bool>> := map[seq(-0x35f, i3  => (v10)) := [v2]];
				v17 := (v17 - {v1}) + (set v19 : string | v19 in v18 :: (v19));
			} else {
				v4[65] := if (!v2 in v3) then v3[!v2] else f7;
				v4[65] := fm20(v2, v2, v2, globalState) / v8.f8;
				var v20 := DC17(p0);
				var v21: seq<string> := [v1, fm8(v2, v7, v2, globalState)];
				var v22: set<string> := {v1, "j"};
				var v23: array<string> := new string[21] [v1, v1, v1, v1 + v1, seq(0x16a, i4  => ('b')), v1 + p0, v20.cf26, p0 + (seq(0xa, i5  => ('l'))), seq(0x325, i6  => ('u')), v1, p0, seq(0x207, i7  => ('o')), v20.cf26 + "klcneosa", "vbsipip", "wudxpkey", p0, v1 + v21[|v22|], v1, v1 + v1, p0[|v5| := fm5(globalState)], v1];
				v23[153] := v1;
				var v24: seq<bool> := [v2];
				var v25: map<int, bool> := map[v0 := DC12(v2).cf19];
				var v26: map<int, int> := map[|v24| := |v25|];
				var v27 := 'j';
				v23[153] := v1[|v26| := v27] + "vhe";
				var v28: array<bool> := new bool[12];
				v28[300] := v2;
				v28[300] := if (!(v2 ==> v2)) then v2 else v2;
				var v29: array<seq<seq<bool>>> := new seq<seq<bool>>[9](i8 => [v24]);
				var v30: seq<seq<bool>> := [v24, v24];
				v29[565] := v30 + v30;
				var v31 := DC13(v28[300], v27);
				var v32: map<D5, char> := map[v31 := v27];
				var v33 := DC20(v27, |(if (true) then v32 else v32)|);
				v4[65], v29[565], globalState.f3, v33, globalState.f4 := v8.f7, v30, v2, DC20(v27, --v8.f8), !v28[300];
				v0 := v8.f7;
			}
			
			var v34 := 'k';
			var v35 := DC13(v2, v34);
			match (v35) {
				case DC12(cf19) =>
					v4[170] := -v8.f7;
					v4[170] := f8;
					v0 := v8.f7;
					var v36: map<bool, int> := map[cf19 := v4[170]];
					v4[170] := |v36| / v4[170];
					var v37: array<bool> := new bool[19](i9 => v2);
					var v38: seq<array<bool>> := [v37];
					var v39 := DC9(f8, v38, -f8, -v8.f8);
					var v40: map<D4, int> := map[v39.(cf12 := v0, cf15 := v0) := v8.f7];
					var v41: map<int, bool> := map[f8 := cf19];
					v40 := v40[v39 := |v41|];
				case DC13(cf20, cf21) =>
					var v42: set<bool> := {!cf20};
					var v43: set<set<bool>> := {{!v2}, v42, v42 - v42};
					v43 := {v42, v42, v42, {cf20}} + v43;
					var v44: array<bool> := new bool[15] [cf20, cf20, v2, v2, cf20, cf20, true, true, cf20, v2, cf20, true, cf20, true, cf20];
					var v45: multiset<array<bool>> := multiset{v44};
					globalState.f3 := (if (cf20) then v45 else v45) >= v45;
					v4[172] := f8;
					v4[172] := -(-|v3| * v8.f8);
					var v46: array<array<bool>> := new array<bool>[17] [v44, v44, v44, v44, v44, v44, v44, v44, v44, v44, v44, v44, v44, v44, v44, v44, v44];
					v46[149] := v44;
					v46[149] := if ({cf20, v2, v2, v2, cf20} > v42) then v44 else v44;
				case DC11(cf18) =>
					var v47: C5 := new C5("gk", v8.f8, -fm3(v2, v2, v2, v2, globalState), v8.f7);
					v47 := v47;
					var v48: seq<seq<int>> := [v7];
					v48 := v48;
					var v49: T2 := new C3(v2, |p0| / v8.f7, -v8.f8);
					v49 := new C2(0x39a, -v8.f8);
					var v50, v51, v52 := m3(f7 > v49.f8, v0, if (v2) then v4 else v4, fm4(v2, v2, globalState), globalState);
				case DC14(cf22) =>
					var v53 := DC2();
					var v54: seq<D1> := [v53];
					v54 := v54;
					var v55: map<char, int> := map[v34 := v8.f8];
					v0 := |v55|;
					v0 := v0;
					v0 := |p0|;
			}
			
			var v56: map<int, bool> := map[v0 := v2];
			var v57: map<map<int, bool>, array<int>> := map[v56 := v4];
			var v58: map<array<int>, char> := map[v4 := v34];
			var v59: map<char, array<int>> := map[v34 := v4];
			var v60: array<map<array<int>, char>> := new map<array<int>, char>[14] [map[if (map[|v7| := v2] in v57) then v57[map[|v7| := v2]] else v4 := v34], v58, v58, v58 + v58, v58, v58, map[v4 := v34], v58[if (v34 in v59) then v59[v34] else v4 := v34], v58 + map[v4 := v34], v58, if (v2) then v58 else v58, map[v4 := v34] + v58, v58, map[v4 := v34]];
			v60[75] := map[v4 := v34];
			v60[75] := v58;
			var v61: seq<array<int>> := [v4, v4];
			var v62: array<array<int>> := new array<int>[17] [v4, v4, v4, v4, v4, v4, v4, v4, v4, v4, v4, v61[f8], v4, v4, v4, v4, v4];
			var v63 := DC8(v62);
			match (v63) {
				case DC9(cf12, cf13, cf14, cf15) =>
					var v64: seq<seq<char>> := [v1, [v34, v34], v1];
					cf14 := (v8.f7 * |v3|) % ((if (v2 in v3) then v3[v2] else v8.f8) + |v64|);
					var v65: map<bool, int> := map[v2 := 0x287];
					var v66: set<int> := {v8.f8, v8.f7 / 901, --(if (v2 in v65) then v65[v2] else |"hcam"|), 0x14a};
					v66 := (v66 * {v8.f7}) - v66;
					var v67: array<bool> := new bool[20];
					v67 := if (v2 <==> v2) then v67 else v67;
					v67[29] := v2;
					v67[29] := v2 ==> v2;
				case DC10(cf16, cf17) =>
					var v68: C5 := new C5(v1, 386, |v5|, v8.f7);
					var v69: map<int, C5> := map[cf16 := v68];
					var v70: map<int, map<int, C5>> := map[f8 := v69];
					v70 := v70[v8.f7 := (map[f7 := v68])[0x166 := v68]];
					var v71: map<bool, string> := map[!v2 := v68.f9];
					var v72: map<bool, bool> := map[true := v2];
					var v73 := DC20(v34, v68.f10);
					var v74: seq<D7> := [v73];
					var v75: C0 := new C0(v4, 741, |multiset{cf17}|);
					var v76: map<array<int>, C0> := map[v4 := v75];
					var v77: array<bool> := new bool[26] [v2, v2, !(v68.f9 <= (if (cf17 in v71) then v71[cf17] else "ejyl")), v2, false, f8 >= -0x32, DC12(v2).cf19, cf17 && cf17, v2 <==> v2, false in v72, false, !('q' !in v1), v74 < v74, v2, if (v8.f7 in v56) then v56[v8.f7] else if (-0x38f in v56) then v56[-0x38f] else v2, cf17, v2, v2, (fm32(globalState)).cf20, cf17, !cf17, 391 >= v8.f7, cf17, v4 in v76, v2, cf17];
					v77 := v77;
					globalState.f4 := v2;
					var v78: set<string> := {v1};
					v77[432] := "rhndgwb" in v78;
					v4[814] := -(if (v8.f8 in v5) then v5[v8.f8] else 0x2ce) % cf16;
					v77[432], v34, cf16, v4[814], globalState.f4 := !((v8.f7 / v8.f8) < v8.f8), fm5(globalState), v8.f8, v8.f8, cf16 == v8.f7;
				case DC8(cf11) =>
					var v79: set<int> := {v8.f8};
					var v80: map<char, set<int>> := map[v1[f7] := v79];
					v79 := if (v34 in v80) then v80[v34] else v79;
					v0 := v8.f7 - v8.f7;
					var v81: multiset<char> := multiset{v34};
					globalState.f4 := if (v2) then v2 else v34 !in v81;
					var v82: array<bool> := new bool[12];
					var v83: array<array<bool>> := new array<bool>[4] [v82, if (v2) then v82 else v82, v82, v82];
					v83[381] := v82;
					var v84: map<bool, bool> := map[!v2 := v2];
					v4, v0, v83[381], v0, globalState.f4 := v4, 0x97, v82, v0, if (v2 in v84) then v84[v2] else v2;
			}
			
		}
		var v85 := 's';
		var v86: map<seq<int>, char> := map[(seq(0x114, i10  => (|map[v2 := v2]|))) + v7 := v85];
		v86 := v86[seq(-0x3d5, i11  => (v8.f7)) := v85];
		var v87: map<int, int> := map[v0 := v8.f8];
		v87 := v87[f8 := -0x1cc * v8.f7];
		var v88: map<bool, map<int, int>> := map[v2 := v87];
		var v89: seq<bool> := [false, v2, v2];
		v88 := v88[v89[f7] := v87 + v87];
	}
	method m2(p0: int, p1: bool, p2: int, p3: bool, globalState: GlobalState) returns (r0: seq<bool>, r1: int, r2: bool) {
		if (p3) {
			globalState.f4 := true;
			var v0 := "bjrtfrpd";
			v0 := v0 + v0;
			globalState.f4 := f7 != p0;
			var v1 := 'q';
			var v2 := DC14(DC13(p3, v1));
			v2 := if (p3) then v2 else v2;
			var v3: array<int> := new int[16](i0 => i0 % f7);
			v3[840] := p2;
			v3[840] := f7;
		} else {
			var v4: multiset<bool> := multiset{p3};
			var v5: map<multiset<bool>, int> := map[v4 := p2];
			var v6: multiset<int> := multiset{-0xfd};
			var v7: seq<char> := ['b', 'h'];
			var v8: seq<bool> := [fm24(p3, -|v4|, v7, globalState)];
			var v9 := DC16(v5[v4 := |v6|], v8);
			match (v9) {
				case DC16(cf24, cf25) =>
					var v10 := 'o';
					var v11 := DC20(v10, f7);
					r1 := v11.cf33;
					var v12: T1 := new C4(p1, p0, f7);
					var v13: multiset<T1> := multiset{v12};
					var v14: map<bool, bool> := map[false := p1];
					var v15: map<int, int> := map[v12.f8 := p0];
					var v16: map<string, int> := map[v7 := if (p1 in v4) then v4[p1] else p2];
					var v17: seq<int> := [|cf25|, if (v7 in v16) then v16[v7] else v12.f7, v12.f7, v12.f7];
					var v18: set<bool> := {true, p3, p3};
					var v19: array<int> := new int[18] [0x376, f8, p2 / -p2, p0, f7, |"yiww"|, |v13|, f7, v12.f7 - 0x27c, f8, v12.f7, f7, v12.f8, p0, |v14| + p0, (if (!p1 in v4) then v4[!p1] else if (v12.f7 in v15) then v15[v12.f7] else f8) + v17[|v18|], if (p1) then f8 else -0x24d, |cf25| / 755];
					v19[248] := p0 / p2;
					v19[248] := f8;
					var v20, v21, v22 := m3(p3, |map[0x2b1 := |v7|]|, v19, f8, globalState);
					var v23, v24, v25 := m3(false, f8 - v12.f7, v19, -f7, globalState);
				case DC15(cf23) =>
					r1 := f7;
					var v26 := 'e';
					var v27 := DC20(if (p3) then v26 else v26, f8);
					var v28 := DC13(p1, v26);
					var v29: array<int> := new int[8](i1 => i1 * p0);
					var v30: map<D5, array<int>> := map[v28 := v29];
					var v31: seq<map<D5, array<int>>> := [v30];
					globalState.f3, globalState.f3, v27 := v31[f7] == (v30[v28 := v29])[v28 := v29], if (p3) then p1 else p1, DC20(v26, -p0);
					cf23[844] := p3;
					cf23[844] := !p3;
					r1 := |((seq(0x369, i2  => (v26))) + (seq(0x13, i3  => ('i'))))| * -0x18e;
			}
			
			var v32: map<bool, int> := map[p3 := 662];
			var v33 := DC11(v32);
			match (v33) {
				case DC12(cf19) =>
					var v34: seq<D6> := [v9];
					var v35 := DC6({-0x137, p0, f7, f8, f8});
					var v36: set<int> := {p2, f8};
					var v37: map<int, D3> := map[|v34| % f8 := v35.(cf9 := v36)];
					v37 := v37[-p0 * p0 := v35];
					v7 := v7;
					globalState.f4 := p1;
					r1 := -788;
				case DC13(cf20, cf21) =>
					var v38: array<bool> := new bool[8] [f7 < f7, if (p3) then p3 else cf20, p1, true, cf20, cf20, cf20, cf20];
					r1, v38, r1 := p0, v38, p2 - (p2 / fm3(cf20, p1, p3, true, globalState));
					var v39: seq<array<bool>> := [v38];
					var v40: array<seq<int>> := new seq<int>[5];
					var v41: seq<int> := [f8];
					v40[310] := v41;
					v39, v40[310], v7, cf21 := if (p1) then v39 else v39, (v41 + [f7]) + (v41 + v41), v7, cf21;
					var v42: C1 := new C1(f7, -p0);
					var v43: map<int, bool> := map[p2 := !p3];
					var v44: set<int> := {f7, p2, p0, -|v43|};
					var v45: map<C1, int> := map[v42 := |v44|];
					v45 := v45[v42 := p0];
					var v46: array<int> := new int[23](i4 => i4 % f8);
					v46[684] := p2 / f7;
					var v47: map<int, int> := map[f7 := |v7|];
					var v48: map<bool, map<int, int>> := map[p1 := v47];
					v46[684] := (p0 % |(if (p3 in v48) then v48[p3] else v47)|) + f8;
				case DC11(cf18) =>
					globalState.f3 := p3;
					var v49 := 'j';
					var v50: multiset<char> := multiset{v49};
					var v51: seq<int> := [|v50|, p0, p2];
					var v52 := new C5(v7, f7, v51[p0], f8);
					r1 := f7;
					r1 := p2;
				case DC14(cf22) =>
					v32 := v32[true := -f7];
					r1 := f8 + (if (p3) then f8 else f8);
					var v53: array<int> := new int[7](i5 => i5 % 0x288);
					v53[659] := |v7|;
					v53[659] := p0;
					v53[659] := 0x2f6;
			}
			
			var v54 := DC3(fm23(f8, globalState));
			match (v54) {
				case DC2() =>
					var v55: array<bool> := new bool[15](i6 => p1);
					v55[715] := p1;
					var v56: array<int> := new int[24];
					v56[870] := p2;
					var v57: C1 := new C1(f7, f8);
					var v58: seq<C1> := [v57];
					v55[715], globalState.f4, r1, v56[870] := p1 && (v58 != v58), p0 < |v8|, fm16(!p1, globalState), p2 + 0x2e2;
					var v59: map<int, int> := map[(if (v55[715] in v4) then v4[v55[715]] else f8) - f7 := p0];
					v59 := v59[594 / 0x93 := p2];
					v56[870] := |(v7 + (if (v55[715]) then v7 else v7))|;
					v56[870] := p2 * v56[870];
				case DC1(cf1) =>
					var v60: map<int, bool> := map[f8 := p1];
					var v61: map<int, int> := map[f8 := |map[p2 := p3]|];
					var v62: seq<map<int, int>> := [v61, v61];
					var v63: map<seq<char>, bool> := map[v7 := p3];
					var v64: array<int> := new int[27] [f8, |v60|, fm16(p1, globalState), fm20(p3, p1, p1, globalState), f7, -f7, f7, p0, |v7|, |v62|, |v8|, -f8, 0x1f4, p0, p2, 532, p2, p2, |v7|, p0, f7, f8, f8, |v63|, 0x112, p2, p0];
					var v65: seq<array<int>> := [v64, v64, v64, v64, v64];
					var v66: array<array<int>> := new array<int>[10] [v64, v64, v64, v64, v64, v64, v65[-p2], v64, v64, v64];
					v66[861] := v64;
					v66[861] := v64;
					var v67: set<string> := {v7};
					globalState.f4 := (if (v8[p0]) then v67 else v67) >= (v67 - {"hwlofrqbi"});
					var v68: map<bool, string> := map[p1 := v7];
					v64[155] := |v68|;
					var v69: seq<int> := [p0];
					var v70: map<char, string> := map[cf1 := v7];
					var v71: map<bool, array<array<int>>> := map[p3 := if (p3) then v66 else v66];
					v64[155], v69, globalState.f3, r1, v66 := -f8 + f7, v69, if (cf1 !in v7) then p1 else if (f8 in v60) then v60[f8] else !false, f7 / |v70|, if (true in v71) then v71[true] else v66;
					v64[155] := -(-f7 / f8) * 0x286;
				case DC3(cf2) =>
					var v72: set<int> := {p2};
					v72 := v72;
					var v73: T0 := new C5("svx", |(seq(0x214, i7  => ('f')))|, 0x87, f8);
					var v74: map<string, T0> := map[v7 := v73];
					var v75: seq<T0> := [if (fm2(p3, globalState) in v74) then v74[fm2(p3, globalState)] else v73, v73];
					var v76 := new C2(|v75|, fm20(p3, false, p1, globalState));
					v8 := v8;
					var v77: array<array<array<int>>> := new array<array<int>>[3];
					var v78: array<array<int>> := new array<int>[6];
					v77[418] := v78;
					var v79: array<seq<C1>> := new seq<C1>[16];
					var v80 := DC19(p0, p1, f7);
					var v81: map<int, bool> := map[v73.f8 - p2 := v80.cf31 <= p2];
					v77[418], v7, v79, globalState.f3, v81 := v78, v7, v79, p3, v81;
			}
			
			r1 := p2;
			var v82: array<array<int>> := new array<int>[15];
			var v83: array<bool> := new bool[19];
			var v84: map<bool, array<bool>> := map[p3 := v83];
			var v85 := 'w';
			var v86: array<int> := new int[14] [if (p0 in v6) then v6[p0] else f8, f7, p0, |v84|, f7, -p0, p2, fm4(true, true, globalState), p2, fm20(p1, p1, p3, globalState), f7, 0x3c5, |(multiset{v85})[v85 := |v7|]|, p0];
			v82[933] := v86;
			var v87 := DC4(v4);
			var v88: multiset<D2> := multiset{v87, v87, v87};
			var v89: map<int, bool> := map[f7 := p3];
			r1, v82[933], v88, r1 := if (p1) then -0xc5 else f8 / |v89|, v86, multiset{v87} + v88, f8 + f7;
		}
		
		if (p3) {
			var v90 := 'h';
			v90 := v90;
			r1 := if (!true) then f7 else if (p1) then p0 else f7;
			var v91: array<int> := new int[8](i8 => i8 * -0x13d);
			v91[552] := 0x275;
			var v92: multiset<int> := multiset{p2, p2};
			var v94: set<set<int>> := {fm29(p3, v92, globalState), (set v93 : int | (0xe0 <= v93) && (v93 < 0x1de) :: (v93 + 701)) * {-460, f8, -p2}};
			var v95 := "ctt";
			var v96: set<int> := {788};
			v91[552], v94, r1 := -(f7 % |v95|), (v94 - {v96}) - v94, p0;
			r1 := fm3(p3, p1, p3, false, globalState) * (-0xe2 * v91[552]);
			var v97: seq<bool> := [p1, DC13(fm24(p1, f7, v95, globalState), v90).cf20, !p3];
			var v98: T2 := new C2(f8, p2);
			var v99: map<int, T2> := map[|v96| := v98];
			var v100: array<bool> := new bool[4] [p1, v97[|v99|], p3, p3];
			v100[225] := p1;
			v100[225] := v97[731];
		} else {
			var v101: map<bool, bool> := map[false := p1];
			r1 := |v101[p3 := p1]| * (|"ymjehvl"| + f7);
			var v102 := "uaehms";
			var v103: seq<string> := [v102, seq(0x112, i9  => ('u')), v102];
			if ((v102 + v102) != (v103[f7] + v102)) {
				r1 := f8;
				var v104: seq<bool> := [p1];
				var v105 := 'i';
				var v106: seq<int> := [f8];
				var v107: array<int> := new int[10] [|v104[p2 := p1]|, DC20(v105, |v103|).cf33, f7, p2 * |map[!p3 := p1]|, if (p3) then p2 else f7, |v106| + -f8, p2, 545 - --0x2d3, p2, f7];
				var v108: map<char, bool> := map[v105 := p3];
				var v109: C5 := new C5(v102, f7, f7, -|v108|);
				var v110: map<int, C5> := map[f7 := v109];
				v107[17] := f8 - |v110[f8 := v109]|;
				v107[17], r1 := f7, f8;
				v105 := v105;
				var v111: map<char, int> := map['k' := f8];
				var v113: set<char> := {fm5(globalState)};
				var v114: seq<set<char>> := [{v105} + (set v112 : char | v112 in v111 :: (v112)), v113];
				v114 := v114 + (v114 + v114);
				v107[17] := f7 + f7;
			} else {
				var v115: seq<bool> := [p1];
				var v116 := DC5(p1, f8, |v115[|v115| := p1]|, f7, fm24(p3, 0x3e6, v102, globalState));
				var v117: set<bool> := {p3};
				var v118: array<bool> := new bool[17] [!p3, p1, !!(if (v116.cf4) then p3 else fm1(v117, f7, -0x2cf, --p0, globalState)), p3, p3, p3, v102 < (seq(0x117, i10  => ('b'))), p1, p1, true <== p3, p3, p3, p1, true <== p1, p3, p3, p3];
				v118[271] := true;
				v118[271] := f8 != (f8 + p2);
				var v119: set<int> := {--|v101|, p2};
				r1 := p0 * |v119|;
				var v120: map<int, int> := map[p0 := f7];
				var v121 := new C1(if (p0 in v120) then v120[p0] else f7, if (!p3) then f7 else fm3(p1, true, v118[271], p1, globalState));
				var v122 := DC13(p1, 'q');
				var v123: map<int, seq<D5>> := map[fm16(p3, globalState) % p0 := [v122, v122, fm32(globalState), fm32(globalState)]];
				var v124: multiset<bool> := multiset{p3};
				var v125: multiset<int> := multiset{if (p3 in v124) then v124[p3] else f7};
				var v126: seq<int> := [|v125| / 0x33b, -(-0x1c3 + p0), p0, f7 % p2];
				globalState.f4, v102, v123, v118[271], v126 := if (p1) then false else p1, v102, v123, v118[271], v126;
				r1 := |(v102 + v102)| + (if (fm20(p3, v118[271], p1, globalState) in v125) then v125[fm20(p3, v118[271], p1, globalState)] else |v115|);
			}
			
			var v127 := new C3(!p3, f7, p2);
			var v128: array<string> := new string[10];
			v128[414] := v102;
			v128[414] := v102 + v102;
			var v129: set<bool> := {v127.f12};
			var v130: seq<bool> := [p1, p3];
			var v131: array<int> := new int[18] [|v102|, -(f8 % p2), p0, p2, p2, -f8, f7, f8, |v129|, |v130|, -703, p2 % |v129|, |v127.fm2(true, globalState)|, p0, f7 + p2, f8, 0x7d - f8, p2];
			v131 := v131;
		}
		
		if (p0 >= p2) {
			var v132: seq<int> := [p0];
			var v133: seq<seq<int>> := [v132, seq(0x306, i11  => (0x1af))];
			var v134: map<int, int> := map[f7 := |v133|];
			v134 := v134[-0x10b := f8 + p0];
			var v135: multiset<bool> := multiset{p1, p3, p3, !p3, !p1};
			var v136: map<multiset<bool>, bool> := map[v135 := p1];
			v136 := v136[v135 := f8 < f7];
			var v137: map<int, bool> := map[f8 := p3];
			v137 := v137[-f7 := true];
			globalState.f4 := p3;
			globalState.f3 := p0 > f7;
		} else {
			var v138: seq<int> := [p2];
			var v139 := "fxgvl";
			var v140: map<int, int> := map[|v138| := |v139|];
			var v141: set<bool> := {p1, p1, p1, p3, p3};
			var v142: multiset<set<bool>> := multiset{v141};
			var v143: map<bool, bool> := map[!false := true];
			var v144: multiset<int> := multiset{p2};
			var v145: multiset<multiset<int>> := multiset{v144, v144};
			var v146: array<int> := new int[9] [f7, -f8, f7, 34, if (v144 in v145) then v145[v144] else p0, |map[p1 := !p3]|, p0, |v139|, p2];
			var v147: map<array<int>, bool> := map[v146 := p3];
			var v148: array<int> := new int[26] [p2, if (p0 in v140) then v140[p0] else p0, f8, f8, 0xc9, p2, if (v141 in v142) then v142[v141] else |v143|, p2, -0x1b2, v138[p0], fm3(p3, p1, p1, if (p3 in v143) then v143[p3] else p1, globalState), p2, f7, p2, |v138|, f7, f7, |v138|, -|v147|, p0, p0, p2, p2, fm20(p3, true, false, globalState), p2, p0];
			var v149 := new C0(v146, f7, f8);
			var v150: seq<bool> := [p1];
			r0 := v150 + v150;
			var v151: map<bool, int> := map[f7 < f7 := f7];
			v151 := v151[p3 := p2 - f8];
			r1 := p2;
			v149.f13[37] := p2;
			v149.f13[37] := p2;
		}
		
		if (p1) {
			var v152 := 'e';
			var v153 := "vlqsggw";
			var v154: seq<char> := [v153[p0], v152, fm5(globalState)];
			v152 := if (p1) then v152 else v154[f8];
			var v155: array<int> := new int[13](i12 => i12 / f8);
			var v156: map<array<int>, int> := map[v155 := p2];
			v156 := v156[v155 := if (!p3) then f8 else p0];
			r1 := p0;
			if (p3) {
				var v157 := new C2(-|v154|, f8);
				globalState.f3 := p1;
				v155 := v155;
				var v158: seq<bool> := [p1];
				var v159: seq<seq<bool>> := [v158];
				var v160: seq<int> := [f8, p0, |v159[p0]|];
				var v161: multiset<int> := multiset{p0, f8, -p0};
				v155[918] := -(|"ikupdc"| * (if (-p2 in v161) then v161[-p2] else |v160|));
				var v162: map<int, seq<int>> := map[|v158| := v160 + v160];
				var v163: map<bool, bool> := map[p1 := p3];
				var v164: multiset<seq<bool>> := multiset{v159[|v163|]};
				r1, v160, v155[918] := f7, if (p0 in v162) then v162[p0] else v160, (-|v164| + p2) - f7;
				globalState.f4 := !p1;
			} else {
				var v165: set<bool> := {p3, p1, p1};
				var v166: set<bool> := {p3, p1, fm1(v165, p2, -90, |v153|, globalState), !p1, p3};
				var v167 := DC2();
				var v168 := DC3(v167);
				var v169 := DC3(v168);
				var v170 := DC3(v169);
				var v171: multiset<D1> := multiset{DC3(v169), v170};
				var v172: array<bool> := new bool[15] [!p1, p3, p1, p3, fm1(v165, p2, if (v170 in v171) then v171[v170] else p0, f7, globalState), p1, p3, p1, p3, p3, p3, p3, p3, p1, p1];
				var v173 := DC15(v172);
				v173, v172, r1, r1 := v173, if (p1) then v172 else v172, fm4(p3, p3, globalState), f7;
				v172[781] := false;
				var v174 := DC0(p3);
				var v175: array<D0> := new D0[16] [DC0(p1), v174, v174.(cf0 := p1), DC0(p1), v174, v174, v174, fm33(globalState), fm33(globalState), v174, v174, v174, v174, v174, v174, v174];
				v172[781], v175 := (fm1(v166, f8, f8, f8, globalState) || p3) <== p1, v175;
				var v176: map<bool, int> := map[p3 := p2];
				var v177: seq<bool> := [p3];
				v172[781], v172[781] := (if (p3 in v176) then v176[p3] else p0) >= (p2 / f8), v177[p2];
				var v178: seq<string> := [v154[f8 := v152], v154, "bp", v153, v153];
				var v179: map<multiset<string>, array<bool>> := map[multiset{v153, v178[p2]} := v172];
				var v180: multiset<string> := multiset{v154[p2 := v152]};
				var v181: set<array<bool>> := {if (v180 in v179) then v179[v180] else v172};
				var v182 := DC11(v176);
				v181, v182, globalState.f4, r1 := v181 + v181, v182, v177[|(seq(0xcd, i13  => (v152)))[p2 := 'o']|], p2;
				v177 := v177;
			}
			
			r1 := p2;
		} else {
			globalState.f4 := p3;
			globalState.f3 := (f8 > -0x14) <==> p1;
			r1 := 0xc1 / p0;
			var v183: array<bool> := new bool[25];
			var v184 := DC15(v183);
			v183 := v184.cf23;
			var v185: seq<bool> := [p1, p3, true];
			var v186 := DC16((fm35(p2, globalState)).cf24, v185);
			var v187: map<multiset<bool>, int> := map[multiset{p1} := p0];
			var v188: set<D6> := {DC16(v187, v185)};
			var v189: seq<set<D6>> := [fm34(p1, p3, globalState), {v186, v186} * {v186, v186}, v188, v188, v188 + v188];
			v189 := seq(705, i14  => (v188));
		}
		
		var v190: array<bool> := new bool[24];
		v190[351] := p3;
		v190[351] := p3;
		r1 := p2;
		var v191: seq<bool> := [p3];
		r0 := v191;
		var v192 := DC11((map[v190[351] := p2])[p1 := f8]);
		r1 := -match v192 {
			case DC12(cf19) => -p0
			case DC13(cf20, cf21) => 862
			case DC11(cf18) => f8
			case DC14(cf22) => DC5(DC7(v190[351]).cf10, f7, f8, p0, DC0(p3).cf0).cf7 % -0xae
		};
		r2 := !p3;
	}
	method m3(p0: bool, p1: int, p2: array<int>, p3: int, globalState: GlobalState) returns (r0: int, r1: int, r2: bool) {
		var v0: C2 := new C2(f8, 0x2af);
		var v1: set<C2> := {v0, v0, v0, v0, v0};
		var v2: map<int, set<C2>> := map[f8 := v1 * v1];
		v1 := if (p1 in v2) then v2[p1] else v1;
		var v3: seq<int> := [f8];
		r0 := v3[p3];
		var v4: T1 := new C0(p2, f8, fm20(p0, p0, p0, globalState) + f8);
		v4 := v4;
		var v5 := DC10(-f7, f8 > p3);
		var v6: array<multiset<bool>> := new multiset<bool>[11](i0 => multiset{p0} - multiset{p0, p0, p0, p0, p0});
		v5, v6 := v5, v6;
		for i1 := 0x133 to -0x308 {
			var v7 := DC19(p1, p0, v4.f7);
			var v8: array<bool> := new bool[2](i2 => p0);
			var v9: multiset<array<bool>> := multiset{v8, v8, v8, v8};
			var v10: multiset<int> := multiset{|map[v4 := p0]|};
			var v11 := new C4(multiset{v7.cf29, |v9|} > v10, v4.f8, v4.f8);
			var v12 := 'y';
			v12 := v12;
			var v13: seq<bool> := [p0, p0];
			var v14: seq<seq<bool>> := [v13];
			var v15: C3 := new C3(v11.f11, v4.f7, f7);
			var v16: set<C3> := {v15};
			var v17: C3 := new C3(!(v13 == v14[p1]), 0x11, |v16| + p1);
			var v18 := "eu";
			var v19 := DC21(v15);
			v17, r0, v18 := v19.cf34, v4.f7, (v17.fm2(v17.f12, globalState))[v4.f8 := if (v15.f12) then v12 else v12];
			r2 := v15.f12;
		}
		var v20 := 'a';
		var v21: set<bool> := {false};
		var v22: map<bool, set<bool>> := map[!true := v21];
		var v23: map<int, set<bool>> := map[v4.f7 := if (true in v22) then v22[true] else v21];
		var v24: seq<bool> := [p0, p0];
		var v25: map<char, set<bool>> := map[v20 := if (|v24| in v23) then v23[|v24|] else {!p0}];
		v25 := v25[v20 := v21];
		r0 := 0x368;
		r1 := v4.f7;
		r2 := p0;
	}
}

method Main() {
	var v0: array<bool> := new bool[18](i1 => false);
	var v1: map<string, array<bool>> := map[seq(230, i0  => ('n')) := v0];
	var v2 := "w";
	var globalState := new GlobalState(0x7c, true, v1 + v1, false, true, 0x263, v2);
	var v3 := false;
	globalState.f3 := v3;
	var v4 := 0x2f4;
	var i2 := 0;
	while ((v4 % v4) <= v4)
		decreases 100 - i2
	{
		if (i2 >= 100) {
			break;
		}
		
		i2 := i2 + 1;
		var v5, v6 := m0(v4, v3, v3, globalState);
		var v7: map<bool, bool> := map[true := v6];
		var v8: seq<bool> := [v3];
		v7 := v7[v3 in v8 := v3];
		var v9: set<string> := {v2};
		var v10 := new C2(|v9| - v4, -v4);
		v7 := v7[v3 := v3];
	}
	var v11: seq<int> := [0x58];
	globalState.f4 := v11[-v4] == |v2|;
	var v12: map<int, bool> := map[823 := false];
	var v13: C4 := new C4(v3, |map[v3 := v3]|, v4);
	var v14: seq<C4> := [v13, v13];
	var v15: array<int> := new int[23] [218, |v2|, v4, |v12|, v4, 709, v4, |v14|, v4, v4, v4, |[v13.f11, v13.f11]|, v4, -v4, 0, v4, v4, v4, v4, v4, v4, v4, v4];
	var v16: seq<array<int>> := [v15];
	var v17 := new C0(v16[v4], 0x20d, v4);
	v0[913] := !!(v13.f11 <== v13.f11);
	var v18: array<array<int>> := new array<int>[1] [v15];
	v18[677] := v17.f13;
	var v19: seq<bool> := [v13.f11, fm8(true, v11, v3, globalState) == v2, if (!v13.f11) then v3 else v13.f11, v12 != v12, v13.f11 <==> v3];
	var v20 := 'b';
	var v21 := DC1(v20);
	var v22 := DC3(DC1(v20));
	var v23: multiset<D1> := multiset{DC3(v21), v22};
	var v24: set<multiset<D1>> := {v23, v23, multiset{v22}};
	v0[913], v18[677], v19, v4, v13 := !(if (v4 >= |v2|) then v3 else v3), v15, v19, fm20(v3, fm36(v3, v13.f11, globalState) > v24, v13.f11, globalState), v13;
	var v25: multiset<array<int>> := multiset{v17.f13};
	var v26: set<bool> := {v13.f11};
	var v27: map<bool, bool> := map[v13.f11 := v17.fm1(v26, v4, v4, v4, globalState)];
	var v28: map<int, multiset<array<int>>> := map[-v4 := v25[v17.f13 := v4]];
	v0, v4, v25 := v0, if (if (v13.f11 in v27) then v27[v13.f11] else v13.f11) then |(v26 - v26)| else v4, if (-v4 in v28) then v28[-v4] else v25;
	var v29: T2 := new C2(v4, v4);
	var v30: map<T2, map<int, bool>> := map[v29 := v12];
	var v31 := new C2(v4, 0x164 + |(if (v29 in v30) then v30[v29] else map[|map[v29.f8 := v17]| := v0[913]])|);
	var i3 := 0;
	while ((v2 + v2) < (seq(0x2ab, i4  => (v20))))
		decreases 100 - i3
	{
		if (i3 >= 100) {
			break;
		}
		
		i3 := i3 + 1;
		v0[913] := !v13.f11;
		var v32: array<seq<bool>> := new seq<bool>[17];
		var v33: map<multiset<bool>, int> := map[multiset{v0[913]} := v29.f7];
		var v34 := DC16(v33, v19);
		v32[35] := v34.cf25;
		v32[35] := v19;
		var v35 := DC20(v20, v4);
		var v36: map<D7, char> := map[v35 := v20];
		var v37 := new C4(v3, -|[v36, v36]|, v29.f7);
		v2 := if (v0[913]) then v31.fm14(globalState) + v2 else seq(220, i5  => (v20));
	}
	var v38: array<seq<set<bool>>> := new seq<set<bool>>[5](i6 => [{v0[913], v13.f11}]);
	var v39: seq<set<bool>> := [v26];
	v38[708] := v39;
	var v40 := DC10(v4, v13.f11);
	v0[913], v38[708] := v40.cf17, v39 + v39;
	var v41: C5 := new C5(v2, v4, v29.f8, v29.f7);
	var v42: map<C5, int> := map[v41 := v41.f10];
	var v43: multiset<array<bool>> := multiset{v0, v0};
	globalState.f4 := v13.fm1(v26, v29.f8, if (v41 in v42) then v42[v41] else v4, if (v0 in v43) then v43[v0] else |v11|, globalState);
	globalState.f3 := !(|fm9(v3, globalState)| > (-|v41.f9[|v26| := v20]| * v29.f7));
	for i7 := v29.f7 - |v11| to v4 {
		v4 := v4;
		var v44: seq<T2> := [v29, v29];
		var v45: C2 := new C2(v29.f7, |(v44 + v44)|);
		v45 := v45;
		v45 := v45;
		v45.m1(v2, globalState);
	}
	var v46: array<D8> := new D8[9];
	forall i8 | 0 <= i8 < v46.Length {
		v46[i8] := DC22(-v41.f10);
	}
	forall i9 | 0 <= i9 < v17.f13.Length {
		v17.f13[i9] := i9 / v29.f7;
	}
	globalState.f3 := !v3;
	v2 := v2;
}