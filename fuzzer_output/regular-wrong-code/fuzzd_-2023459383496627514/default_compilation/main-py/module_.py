import sys
from typing import Callable, Any, TypeVar, NamedTuple
from math import floor
from itertools import count

import module_
import _dafny
import System_

# Module: module_

class default__:
    def  __init__(self):
        pass

    @staticmethod
    def abs(x):
        if (x) < (0):
            return (-1) * (x)
        elif True:
            return x

    @staticmethod
    def safeIndex(x, length):
        if (x) < (0):
            return 0
        elif (x) >= (length):
            return _dafny.euclidian_modulus(x, length)
        elif True:
            return x

    @staticmethod
    def safeDivisionInt(x1, x2):
        if (x2) == (0):
            return x1
        elif True:
            return _dafny.euclidian_division(x1, x2)

    @staticmethod
    def safeModuloInt(x1, x2):
        if (x2) == (0):
            return x1
        elif True:
            return _dafny.euclidian_modulus(x1, x2)

    @staticmethod
    def fm0(globalState):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: int
            for compr_0_ in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([68 for d_0_i0_ in range(default__.abs(163))]))).Elements:
                d_1_v0_: int = compr_0_
                if (d_1_v0_) in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([68 for d_0_i0_ in range(default__.abs(163))]))):
                    coll0_[(d_1_v0_) + (-976)] = not(False)
            return _dafny.Map(coll0_)
        return _dafny.Map({(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ibq"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qjdrauoj"))): default__.safeDivisionInt(len(iife0_()
        ), 120)})

    @staticmethod
    def fm1(globalState):
        return _dafny.CodePoint('h')

    @staticmethod
    def fm2(p0, p1, p2, p3, globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k'), _dafny.CodePoint('g'), _dafny.CodePoint('i')])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n'), _dafny.CodePoint('o')]))) + ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a')])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x'), _dafny.CodePoint('j'), _dafny.CodePoint('k'), _dafny.CodePoint('d'), _dafny.CodePoint('d')])))

    @staticmethod
    def fm3(p0, globalState):
        def iife1_():
            coll1_ = _dafny.Map()
            compr_1_: str
            for compr_1_ in (_dafny.Map({_dafny.CodePoint('f'): _dafny.Map({False: True})})).keys.Elements:
                d_2_v0_: str = compr_1_
                if (d_2_v0_) in (_dafny.Map({_dafny.CodePoint('f'): _dafny.Map({False: True})})):
                    coll1_[d_2_v0_] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p')])
            return _dafny.Map(coll1_)
        return ((_dafny.Map({_dafny.CodePoint('m'): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g')])})) | (_dafny.Map({_dafny.CodePoint('b'): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p'), _dafny.CodePoint('a')])}))) | (iife1_()
        )

    @staticmethod
    def fm4(p0, globalState):
        source0_ = D3_DC8(_dafny.CodePoint('h'), not(False), (D0_DC1(False, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ndasii"))))).cf1, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aqwpe")))
        if source0_.is_DC8:
            d_3___mcc_h0_ = source0_.cf17
            d_4___mcc_h1_ = source0_.cf18
            d_5___mcc_h2_ = source0_.cf19
            d_6___mcc_h3_ = source0_.cf20
            d_7_cf20_ = d_6___mcc_h3_
            d_8_cf19_ = d_5___mcc_h2_
            d_9_cf18_ = d_4___mcc_h1_
            d_10_cf17_ = d_3___mcc_h0_
            return not(d_9_cf18_)
        elif source0_.is_DC7:
            d_11___mcc_h4_ = source0_.cf16
            d_12_cf16_ = d_11___mcc_h4_
            return (False) == (True)
        elif True:
            d_13___mcc_h5_ = source0_.cf21
            d_14_cf21_ = d_13___mcc_h5_
            return (True) or (True)

    @staticmethod
    def fm8(p0, globalState):
        return _dafny.SeqWithoutIsStrInference([247, default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([-345, 863])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xdtj")))), -684])

    @staticmethod
    def fm12(globalState):
        def iife2_():
            coll2_ = _dafny.Map()
            compr_2_: int
            for compr_2_ in _dafny.IntegerRange(998, 270):
                d_15_v0_: int = compr_2_
                if ((998) <= (d_15_v0_)) and ((d_15_v0_) < (270)):
                    coll2_[(d_15_v0_) - (322)] = len(_dafny.Map({True: len(_dafny.Set({True}))}))
            return _dafny.Map(coll2_)
        return _dafny.SeqWithoutIsStrInference([((0) - ((0) - (len(_dafny.Map({D0_DC1(False, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b")))): False}))))) >= ((0) - (len((D6_DC17((0) - (len(iife2_()
)), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_16_i0_ in range(default__.abs(653))]), False)).cf39))), not (True) or (not(False))])

    @staticmethod
    def fm14(p0, globalState):
        return (_dafny.Set({not(not(False)), not(not(True))})) | (_dafny.Set({False}))

    @staticmethod
    def fm15(p0, p1, globalState):
        return ((_dafny.SeqWithoutIsStrInference([True, True]) if True else _dafny.SeqWithoutIsStrInference([False, True]))) + (_dafny.SeqWithoutIsStrInference([False]))

    @staticmethod
    def fm16(p0, globalState):
        return D0_DC1(True, 9)

    @staticmethod
    def fm17(p0, p1, globalState):
        return _dafny.Map({len((_dafny.Map({False: D3_DC9(D3_DC9(D3_DC8(_dafny.CodePoint('k'), False, True, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bvdjclmpx")))))})) | (_dafny.Map({True: D3_DC9(D3_DC9(D3_DC9(D3_DC7(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_17_i0_ in range(default__.abs(92))])))))}))): (_dafny.MultiSet([False])) | (_dafny.MultiSet([not(True)]))})

    @staticmethod
    def fm18(p0, p1, globalState):
        return ((_dafny.Map({False: not(False)}) if not(not(True)) else _dafny.Map({not(not(True)): False}))) | ((_dafny.Map({True: False})) | (_dafny.Map({False: True})))

    @staticmethod
    def fm19(p0, p1, globalState):
        if True:
            return default__.safeDivisionInt(len(_dafny.Map({717: len(_dafny.Set({False}))})), 59)
        elif True:
            return (_dafny.MultiSet([False, False, False, False, True])).cardinality

    @staticmethod
    def fm20(p0, p1, globalState):
        return ((D1_DC4(not(False), True, 747, _dafny.SeqWithoutIsStrInference([883, 416]), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u')]))) if True else D1_DC4(False, True, -294, (D1_DC4(False, False, len(_dafny.Map({len(_dafny.Map({not(True): False})): False})), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mpil"))), 275, 715, len(_dafny.SeqWithoutIsStrInference([True])), (0) - (-940)]), 17)).cf8, len(_dafny.Map({len(_dafny.Map({306: not(True)})): (0) - (len(_dafny.Map({False: 230})))}))))).cf8

    @staticmethod
    def fm21(globalState):
        source1_ = D3_DC8(_dafny.CodePoint('t'), True, False, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_18_i0_ in range(default__.abs(14))]))
        if source1_.is_DC8:
            d_19___mcc_h0_ = source1_.cf17
            d_20___mcc_h1_ = source1_.cf18
            d_21___mcc_h2_ = source1_.cf19
            d_22___mcc_h3_ = source1_.cf20
            d_23_cf20_ = d_22___mcc_h3_
            d_24_cf19_ = d_21___mcc_h2_
            d_25_cf18_ = d_20___mcc_h1_
            d_26_cf17_ = d_19___mcc_h0_
            return D0_DC0((d_23_cf20_)[default__.safeIndex(669, len(d_23_cf20_))])
        elif source1_.is_DC7:
            d_27___mcc_h4_ = source1_.cf16
            d_28_cf16_ = d_27___mcc_h4_
            return D0_DC0(_dafny.CodePoint('d'))
        elif True:
            d_29___mcc_h5_ = source1_.cf21
            d_30_cf21_ = d_29___mcc_h5_
            return D0_DC0(_dafny.CodePoint('e'))

    @staticmethod
    def fm22(globalState):
        return _dafny.Set({(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "icl"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x"))), (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_31_i0_ in range(default__.abs(522))])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_32_i1_ in range(default__.abs(202))]))})

    @staticmethod
    def fm23(p0, p1, p2, p3, globalState):
        return D3_DC8(_dafny.CodePoint('w'), False, ((0) - (len(_dafny.Map({-362: True})))) >= ((0) - ((_dafny.MultiSet([not(not(False)), True, False])).cardinality)), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ykgprxn"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bkcd"))))

    @staticmethod
    def fm24(p0, p1, p2, p3, globalState):
        return _dafny.MultiSet(((_dafny.SeqWithoutIsStrInference([True, False])) + (_dafny.SeqWithoutIsStrInference([not(False), True]))) + ((_dafny.SeqWithoutIsStrInference([not(False)])) + (_dafny.SeqWithoutIsStrInference([not(False)]))))

    @staticmethod
    def fm25(p0, globalState):
        return (_dafny.Map({501: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lifhwcw"))})) | (_dafny.Map({360: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "altifx"))}))

    @staticmethod
    def fm26(globalState):
        return ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_33_i1_ in range(default__.abs(-977))]) for d_34_i0_ in range(default__.abs(161))]))).intersection(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ujto")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pfcmr")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jnkits")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_35_i2_ in range(default__.abs(70))])]))) - ((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ynytpecsu"))]) if True else _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x"))])))

    @staticmethod
    def fm27(p0, p1, p2, p3, globalState):
        def iife3_():
            coll3_ = _dafny.Map()
            compr_3_: int
            for compr_3_ in _dafny.IntegerRange(12, 21):
                d_37_v0_: int = compr_3_
                if ((12) <= (d_37_v0_)) and ((d_37_v0_) < (21)):
                    coll3_[(d_37_v0_) * (len(_dafny.SeqWithoutIsStrInference([279])))] = len(_dafny.Map({False: _dafny.CodePoint('c')}))
            return _dafny.Map(coll3_)
        return D1_DC5(((0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k')])))) - (4), False, (len(_dafny.Map({False: (0) - (-758)})) if False else (0) - (len(_dafny.Map({True: (_dafny.MultiSet([715])).cardinality})))), (_dafny.Set({len(_dafny.SeqWithoutIsStrInference([764 for d_36_i0_ in range(default__.abs(41))])), 938, 739, (_dafny.MultiSet([not(True)])).cardinality, 0})).issubset(_dafny.Set({352, len(_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wfi"))), 763])))])), 980, len(iife3_()
)})), -549)

    @staticmethod
    def fm28(p0, globalState):
        return ((_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_38_i0_ in range(default__.abs(465))]))})) | (_dafny.Map({True: 895}))) | ((_dafny.Map({not(False): 855}) if True else _dafny.Map({False: 677})))

    @staticmethod
    def fm29(p0, p1, p2, globalState):
        return (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jcgiuvj"))) for d_39_i0_ in range(default__.abs(242))]))])) | (_dafny.MultiSet([-791, len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s"))])), -678})), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_40_i1_ in range(default__.abs(-166))]))]))

    @staticmethod
    def fm32(p0, p1, p2, p3, globalState):
        return _dafny.SeqWithoutIsStrInference([((_dafny.MultiSet([657])).cardinality if not(not(True)) else 769)])

    @staticmethod
    def fm33(p0, p1, p2, p3, globalState):
        def iife4_():
            coll4_ = _dafny.Map()
            compr_4_: int
            for compr_4_ in _dafny.IntegerRange(473, 374):
                d_41_v0_: int = compr_4_
                if ((473) <= (d_41_v0_)) and ((d_41_v0_) < (374)):
                    coll4_[(d_41_v0_) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jwdxy"))))] = False
            return _dafny.Map(coll4_)
        source2_ = D8_DC19(iife4_()
)
        if source2_.is_DC20:
            d_42___mcc_h0_ = source2_.cf43
            d_43___mcc_h1_ = source2_.cf44
            d_44___mcc_h2_ = source2_.cf45
            d_45___mcc_h3_ = source2_.cf46
            d_46___mcc_h4_ = source2_.cf47
            d_47_cf47_ = d_46___mcc_h4_
            d_48_cf46_ = d_45___mcc_h3_
            d_49_cf45_ = d_44___mcc_h2_
            d_50_cf44_ = d_43___mcc_h1_
            d_51_cf43_ = d_42___mcc_h0_
            return _dafny.SeqWithoutIsStrInference([d_47_cf47_])
        elif True:
            d_52___mcc_h5_ = source2_.cf42
            d_53_cf42_ = d_52___mcc_h5_
            return _dafny.SeqWithoutIsStrInference([False, False])

    @staticmethod
    def fm34(p0, p1, p2, p3, globalState):
        def iife5_():
            coll5_ = _dafny.Map()
            compr_5_: int
            for compr_5_ in _dafny.IntegerRange(229, -103):
                d_54_v0_: int = compr_5_
                if ((229) <= (d_54_v0_)) and ((d_54_v0_) < (-103)):
                    coll5_[default__.safeModuloInt(d_54_v0_, -349)] = True
            return _dafny.Map(coll5_)
        return ((_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sqdajpdkh"))): True})) | (_dafny.Map({311: not(False)}))) | ((iife5_()
        ) | (_dafny.Map({len(_dafny.Set({_dafny.CodePoint('j')})): False})))

    @staticmethod
    def fm35(globalState):
        if False:
            if True:
                return D6_DC17(121, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h")), (D9_DC22(19, True, _dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_55_i0_ in range(default__.abs(173))])): False}))).cf50)
            elif True:
                return D6_DC17(119, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ndg")), False)
        elif True:
            return D6_DC17(695, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "stgojmq")), True)

    @staticmethod
    def m0(p0, p1, globalState):
        r0: bool = False
        r1: _dafny.MultiSet = _dafny.MultiSet({})
        r2: _dafny.Set = _dafny.Set({})
        r3: int = int(0)
        d_56_v0_: bool
        d_56_v0_ = False
        d_57_i0_: int
        d_57_i0_ = 0
        with _dafny.label("0"):
            while d_56_v0_:
                with _dafny.c_label("0"):
                    if (d_57_i0_) >= (100):
                        raise _dafny.Break("0")
                    d_57_i0_ = (d_57_i0_) + (1)
                    d_58_v1_: _dafny.Array
                    nw0_ = _dafny.Array(int(0), 19)
                    d_58_v1_ = nw0_
                    d_59_v2_: int
                    d_59_v2_ = 849
                    d_60_v3_: _dafny.Set
                    d_60_v3_ = _dafny.Set({len((p1).set(default__.safeIndex(d_59_v2_, len(p1)), not(d_56_v0_))), d_59_v2_, d_59_v2_})
                    index0_ = default__.safeIndex(859, (d_58_v1_).length(0))
                    (d_58_v1_)[index0_] = (d_59_v2_) * (len(d_60_v3_))
                    d_61_v4_: _dafny.Seq
                    d_61_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bcuys"))
                    index1_ = default__.safeIndex(859, (d_58_v1_).length(0))
                    (d_58_v1_)[index1_] = len(((d_61_v4_) + (d_61_v4_)) + ((d_61_v4_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_62_i1_ in range(default__.abs(829))]))))
                    r3 = d_59_v2_
                    r0 = d_56_v0_
                    r0 = d_56_v0_
                    pass
            pass
        d_63_v5_: _dafny.Seq
        d_63_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "btljib"))
        d_64_v6_: int
        d_64_v6_ = -243
        d_65_v7_: _dafny.Map
        d_65_v7_ = _dafny.Map({d_64_v6_: not(d_56_v0_)})
        d_66_v8_: D8
        d_66_v8_ = D8_DC19(d_65_v7_)
        pat_let_tv0_ = d_63_v5_
        pat_let_tv1_ = d_65_v7_
        def lambda0_(source3_):
            if source3_.is_DC20:
                d_67___mcc_h0_ = source3_.cf43
                d_68___mcc_h1_ = source3_.cf44
                d_69___mcc_h2_ = source3_.cf45
                d_70___mcc_h3_ = source3_.cf46
                d_71___mcc_h4_ = source3_.cf47
                d_72_cf47_ = d_71___mcc_h4_
                d_73_cf46_ = d_70___mcc_h3_
                d_74_cf45_ = d_69___mcc_h2_
                d_75_cf44_ = d_68___mcc_h1_
                d_76_cf43_ = d_67___mcc_h0_
                return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iuyjsjqt"))
            elif True:
                d_77___mcc_h5_ = source3_.cf42
                d_78_cf42_ = d_77___mcc_h5_
                return pat_let_tv0_

        def iife6_(_pat_let0_0):
            def iife7_(d_79_dt__update__tmp_h0_):
                def iife8_(_pat_let1_0):
                    def iife9_(d_80_dt__update_hcf42_h0_):
                        return D8_DC19(d_80_dt__update_hcf42_h0_)
                    return iife9_(_pat_let1_0)
                return iife8_(pat_let_tv1_)
            return iife7_(_pat_let0_0)
        d_63_v5_ = lambda0_(iife6_(d_66_v8_))
        d_81_v9_: C0
        nw1_ = C0()
        nw1_.ctor__()
        d_81_v9_ = nw1_
        d_82_v10_: _dafny.MultiSet
        d_82_v10_ = _dafny.MultiSet([d_56_v0_, d_56_v0_])
        r1 = ((_dafny.MultiSet([not(d_56_v0_)])).intersection((_dafny.MultiSet(p1)).set(d_56_v0_, default__.abs(d_64_v6_)))) | ((d_82_v10_) | (_dafny.MultiSet([not(d_56_v0_)])))
        r3 = (d_64_v6_) + ((0) - (d_64_v6_))
        r3 = (default__.fm35(globalState)).cf38
        r0 = (d_64_v6_) < (d_64_v6_)
        r1 = d_82_v10_
        d_83_v11_: D3
        d_83_v11_ = D3_DC7(d_63_v5_)
        pat_let_tv2_ = p0
        pat_let_tv3_ = d_64_v6_
        pat_let_tv4_ = p0
        pat_let_tv5_ = p1
        pat_let_tv6_ = d_64_v6_
        pat_let_tv7_ = p1
        pat_let_tv8_ = d_56_v0_
        pat_let_tv9_ = d_56_v0_
        pat_let_tv10_ = d_56_v0_
        pat_let_tv11_ = d_56_v0_
        pat_let_tv12_ = d_56_v0_
        def lambda1_(source4_):
            if source4_.is_DC8:
                d_84___mcc_h6_ = source4_.cf17
                d_85___mcc_h7_ = source4_.cf18
                d_86___mcc_h8_ = source4_.cf19
                d_87___mcc_h9_ = source4_.cf20
                d_88_cf20_ = d_87___mcc_h9_
                d_89_cf19_ = d_86___mcc_h8_
                d_90_cf18_ = d_85___mcc_h7_
                d_91_cf17_ = d_84___mcc_h6_
                return (_dafny.Set({not(d_90_cf18_), d_90_cf18_, (pat_let_tv2_)[default__.safeIndex(pat_let_tv3_, len(pat_let_tv4_))]})) | (_dafny.Set({(pat_let_tv5_)[default__.safeIndex(pat_let_tv6_, len(pat_let_tv7_))], d_90_cf18_}))
            elif source4_.is_DC7:
                d_92___mcc_h10_ = source4_.cf16
                d_93_cf16_ = d_92___mcc_h10_
                return (_dafny.Set({pat_let_tv8_})) - (_dafny.Set({pat_let_tv9_}))
            elif True:
                d_94___mcc_h11_ = source4_.cf21
                d_95_cf21_ = d_94___mcc_h11_
                return (_dafny.Set({pat_let_tv10_})).intersection(_dafny.Set({pat_let_tv11_, pat_let_tv12_}))

        r2 = lambda1_(d_83_v11_)
        d_96_v12_: str
        d_96_v12_ = _dafny.CodePoint('w')
        d_97_v13_: _dafny.Map
        d_97_v13_ = _dafny.Map({d_64_v6_: d_64_v6_})
        d_98_v14_: _dafny.Seq
        d_98_v14_ = _dafny.SeqWithoutIsStrInference([len(d_97_v13_), d_64_v6_])
        d_99_v15_: _dafny.MultiSet
        d_99_v15_ = _dafny.MultiSet([d_64_v6_])
        r3 = (default__.fm19(d_96_v12_, d_56_v0_, globalState) if (d_64_v6_) <= (d_64_v6_) else len((d_98_v14_) + (_dafny.SeqWithoutIsStrInference([(d_99_v15_).cardinality]))))
        return r0, r1, r2, r3

    @staticmethod
    def Main(noArgsParameter__):
        d_100_v0_: bool
        d_100_v0_ = False
        d_101_v1_: _dafny.Array
        nw2_ = _dafny.Array(None, 26)
        nw2_[int(0)] = d_100_v0_
        nw2_[int(1)] = d_100_v0_
        nw2_[int(2)] = d_100_v0_
        nw2_[int(3)] = d_100_v0_
        nw2_[int(4)] = d_100_v0_
        nw2_[int(5)] = d_100_v0_
        nw2_[int(6)] = d_100_v0_
        nw2_[int(7)] = d_100_v0_
        nw2_[int(8)] = d_100_v0_
        nw2_[int(9)] = d_100_v0_
        nw2_[int(10)] = d_100_v0_
        nw2_[int(11)] = False
        nw2_[int(12)] = True
        nw2_[int(13)] = not(d_100_v0_)
        nw2_[int(14)] = not(d_100_v0_)
        nw2_[int(15)] = d_100_v0_
        nw2_[int(16)] = False
        nw2_[int(17)] = d_100_v0_
        nw2_[int(18)] = False
        nw2_[int(19)] = d_100_v0_
        nw2_[int(20)] = d_100_v0_
        nw2_[int(21)] = d_100_v0_
        nw2_[int(22)] = d_100_v0_
        nw2_[int(23)] = d_100_v0_
        nw2_[int(24)] = d_100_v0_
        nw2_[int(25)] = d_100_v0_
        d_101_v1_ = nw2_
        d_102_v2_: int
        d_102_v2_ = 599
        d_103_v3_: _dafny.MultiSet
        d_103_v3_ = _dafny.MultiSet([d_102_v2_])
        d_104_v4_: _dafny.Seq
        d_104_v4_ = _dafny.SeqWithoutIsStrInference([d_102_v2_])
        d_105_globalState_: GlobalState
        nw3_ = GlobalState()
        nw3_.ctor__(False, 38, 586, False, d_101_v1_, d_101_v1_, d_103_v3_, d_104_v4_, True, 249)
        d_105_globalState_ = nw3_
        d_106_v5_: _dafny.Map
        d_106_v5_ = _dafny.Map({True: default__.safeModuloInt(-550, len(d_104_v4_))})
        d_106_v5_ = d_106_v5_
        d_107_v6_: _dafny.Seq
        d_107_v6_ = _dafny.SeqWithoutIsStrInference([d_100_v0_, d_100_v0_])
        d_108_v7_: bool
        d_109_v8_: _dafny.MultiSet
        d_110_v9_: _dafny.Set
        d_111_v10_: int
        out0_: bool
        out1_: _dafny.MultiSet
        out2_: _dafny.Set
        out3_: int
        out0_, out1_, out2_, out3_ = default__.m0(d_107_v6_, d_107_v6_, d_105_globalState_)
        d_108_v7_ = out0_
        d_109_v8_ = out1_
        d_110_v9_ = out2_
        d_111_v10_ = out3_
        d_111_v10_ = d_102_v2_
        d_112_v11_: _dafny.Set
        d_112_v11_ = _dafny.Set({d_102_v2_, d_111_v10_, d_111_v10_, d_111_v10_, d_102_v2_})
        d_112_v11_ = d_112_v11_
        d_113_v12_: _dafny.Array
        nw4_ = _dafny.Array(int(0), 22)
        d_113_v12_ = nw4_
        index2_ = default__.safeIndex(378, (d_113_v12_).length(0))
        (d_113_v12_)[index2_] = (0) - ((d_102_v2_) * ((0) - (len(d_104_v4_))))
        index3_ = default__.safeIndex(204, (d_101_v1_).length(0))
        (d_101_v1_)[index3_] = not (d_108_v7_) or (d_108_v7_)
        d_114_v13_: _dafny.Seq
        d_114_v13_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fe"))
        d_115_v14_: _dafny.Map
        d_115_v14_ = _dafny.Map({d_114_v13_: d_102_v2_})
        index4_ = default__.safeIndex(378, (d_113_v12_).length(0))
        index5_ = default__.safeIndex(204, (d_101_v1_).length(0))
        rhs0_ = (d_112_v11_) != (d_112_v11_)
        rhs1_ = len(((d_114_v13_).set(default__.safeIndex(d_102_v2_, len(d_114_v13_)), _dafny.CodePoint('b'))) + (d_114_v13_))
        rhs2_ = d_111_v10_
        rhs3_ = d_100_v0_
        rhs4_ = (d_115_v14_ if d_108_v7_ else default__.fm0(d_105_globalState_))
        lhs0_ = d_113_v12_
        lhs1_ = default__.safeIndex(378, (d_113_v12_).length(0))
        lhs2_ = d_101_v1_
        lhs3_ = default__.safeIndex(204, (d_101_v1_).length(0))
        d_100_v0_ = rhs0_
        d_102_v2_ = rhs1_
        lhs0_[lhs1_] = rhs2_
        lhs2_[lhs3_] = rhs3_
        d_115_v14_ = rhs4_
        if (d_112_v11_).issubset((d_112_v11_).intersection(d_112_v11_)):
            index6_ = default__.safeIndex(204, (d_101_v1_).length(0))
            (d_101_v1_)[index6_] = (((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(d_113_v12_)[default__.safeIndex(378, (d_113_v12_).length(0))] for d_116_i0_ in range(default__.abs(496))]))).set(d_102_v2_, default__.abs(d_111_v10_))).set((d_113_v12_)[default__.safeIndex(378, (d_113_v12_).length(0))], default__.abs((0) - (d_102_v2_)))).ispropersubset(d_103_v3_)
            index7_ = default__.safeIndex(378, (d_113_v12_).length(0))
            index8_ = default__.safeIndex(378, (d_113_v12_).length(0))
            rhs5_ = (default__.safeDivisionInt((d_113_v12_)[default__.safeIndex(378, (d_113_v12_).length(0))], len(d_114_v13_))) * ((0) - ((d_113_v12_)[default__.safeIndex(378, (d_113_v12_).length(0))]))
            rhs6_ = default__.safeModuloInt((d_113_v12_)[default__.safeIndex(378, (d_113_v12_).length(0))], (d_113_v12_)[default__.safeIndex(378, (d_113_v12_).length(0))])
            rhs7_ = d_114_v13_
            lhs4_ = d_113_v12_
            lhs5_ = default__.safeIndex(378, (d_113_v12_).length(0))
            lhs6_ = d_113_v12_
            lhs7_ = default__.safeIndex(378, (d_113_v12_).length(0))
            lhs4_[lhs5_] = rhs5_
            lhs6_[lhs7_] = rhs6_
            d_114_v13_ = rhs7_
            d_117_v15_: _dafny.Map
            d_117_v15_ = _dafny.Map({d_108_v7_: (d_101_v1_)[default__.safeIndex(204, (d_101_v1_).length(0))]})
            d_118_v16_: _dafny.Map
            d_118_v16_ = _dafny.Map({(d_104_v4_)[default__.safeIndex(len(d_117_v15_), len(d_104_v4_))]: d_102_v2_})
            index9_ = default__.safeIndex(378, (d_113_v12_).length(0))
            rhs8_ = d_103_v3_
            rhs9_ = (0) - (default__.safeDivisionInt((d_102_v2_) * (len(d_118_v16_)), 373))
            rhs10_ = d_114_v13_
            rhs11_ = 192
            lhs8_ = d_113_v12_
            lhs9_ = default__.safeIndex(378, (d_113_v12_).length(0))
            d_103_v3_ = rhs8_
            lhs8_[lhs9_] = rhs9_
            d_114_v13_ = rhs10_
            d_102_v2_ = rhs11_
            d_114_v13_ = d_114_v13_
            d_119_v18_: D0
            d_119_v18_ = D0_DC0(_dafny.CodePoint('v'))
            d_120_v19_: str
            d_120_v19_ = _dafny.CodePoint('x')
            d_121_v20_: _dafny.Map
            d_121_v20_ = _dafny.Map({d_120_v19_: d_114_v13_})
            d_122_v22_: _dafny.Set
            d_122_v22_ = _dafny.Set({d_120_v19_})
            d_123_v24_: _dafny.Map
            d_123_v24_ = _dafny.Map({d_120_v19_: d_117_v15_})
            d_124_v26_: _dafny.Array
            nw5_ = _dafny.Array(None, 29)
            def iife10_():
                coll6_ = _dafny.Map()
                compr_6_: str
                for compr_6_ in (d_114_v13_).Elements:
                    d_125_v17_: str = compr_6_
                    if (d_125_v17_) in (d_114_v13_):
                        coll6_[d_125_v17_] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p')])
                return _dafny.Map(coll6_)
            nw5_[int(0)] = iife10_()
            
            nw5_[int(1)] = (_dafny.Map({default__.fm1(d_105_globalState_): d_114_v13_})).set((d_119_v18_).cf0, default__.fm2(d_100_v0_, (d_113_v12_)[default__.safeIndex(378, (d_113_v12_).length(0))], _dafny.CodePoint('q'), (d_113_v12_)[default__.safeIndex(378, (d_113_v12_).length(0))], d_105_globalState_))
            nw5_[int(2)] = d_121_v20_
            nw5_[int(3)] = d_121_v20_
            nw5_[int(4)] = default__.fm3(d_100_v0_, d_105_globalState_)
            nw5_[int(5)] = d_121_v20_
            nw5_[int(6)] = d_121_v20_
            nw5_[int(7)] = d_121_v20_
            nw5_[int(8)] = _dafny.Map({d_120_v19_: _dafny.SeqWithoutIsStrInference([default__.fm1(d_105_globalState_)])})
            nw5_[int(9)] = (d_121_v20_) | (d_121_v20_)
            def iife11_():
                coll7_ = _dafny.Map()
                compr_7_: str
                for compr_7_ in (d_122_v22_).Elements:
                    d_126_v21_: str = compr_7_
                    if (d_126_v21_) in (d_122_v22_):
                        coll7_[d_126_v21_] = _dafny.SeqWithoutIsStrInference([d_120_v19_])
                return _dafny.Map(coll7_)
            nw5_[int(10)] = iife11_()
            
            nw5_[int(11)] = default__.fm3(d_100_v0_, d_105_globalState_)
            nw5_[int(12)] = d_121_v20_
            nw5_[int(13)] = d_121_v20_
            nw5_[int(14)] = (d_121_v20_).set(d_120_v19_, d_114_v13_)
            def iife12_():
                coll8_ = _dafny.Map()
                compr_8_: str
                for compr_8_ in (d_123_v24_).keys.Elements:
                    d_127_v23_: str = compr_8_
                    if (d_127_v23_) in (d_123_v24_):
                        coll8_[d_127_v23_] = d_114_v13_
                return _dafny.Map(coll8_)
            nw5_[int(15)] = iife12_()
            
            nw5_[int(16)] = d_121_v20_
            nw5_[int(17)] = d_121_v20_
            nw5_[int(18)] = d_121_v20_
            nw5_[int(19)] = d_121_v20_
            nw5_[int(20)] = d_121_v20_
            nw5_[int(21)] = (d_121_v20_) | (d_121_v20_)
            def iife13_():
                coll9_ = _dafny.Map()
                compr_9_: str
                for compr_9_ in (d_122_v22_).Elements:
                    d_128_v25_: str = compr_9_
                    if (d_128_v25_) in (d_122_v22_):
                        coll9_[d_128_v25_] = d_114_v13_
                return _dafny.Map(coll9_)
            nw5_[int(22)] = (d_121_v20_ if d_100_v0_ else iife13_()
            )
            nw5_[int(23)] = (d_121_v20_).set(d_120_v19_, d_114_v13_)
            nw5_[int(24)] = _dafny.Map({d_120_v19_: d_114_v13_})
            nw5_[int(25)] = _dafny.Map({(d_119_v18_).cf0: _dafny.SeqWithoutIsStrInference([d_120_v19_ for d_129_i1_ in range(default__.abs(655))])})
            nw5_[int(26)] = default__.fm3((d_101_v1_)[default__.safeIndex(204, (d_101_v1_).length(0))], d_105_globalState_)
            nw5_[int(27)] = d_121_v20_
            nw5_[int(28)] = _dafny.Map({d_120_v19_: (d_114_v13_).set(default__.safeIndex(d_102_v2_, len(d_114_v13_)), d_120_v19_)})
            d_124_v26_ = nw5_
            index10_ = default__.safeIndex(4, (d_124_v26_).length(0))
            (d_124_v26_)[index10_] = d_121_v20_
            d_130_v27_: _dafny.Seq
            d_130_v27_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_120_v19_: (d_114_v13_).set(default__.safeIndex((d_113_v12_)[default__.safeIndex(378, (d_113_v12_).length(0))], len(d_114_v13_)), d_120_v19_)}), d_121_v20_, d_121_v20_, d_121_v20_, (d_121_v20_) | (d_121_v20_)])
            index11_ = default__.safeIndex(4, (d_124_v26_).length(0))
            (d_124_v26_)[index11_] = (d_130_v27_)[default__.safeIndex(d_102_v2_, len(d_130_v27_))]
        elif True:
            d_131_v28_: _dafny.Map
            d_131_v28_ = _dafny.Map({d_101_v1_: (d_101_v1_)[default__.safeIndex(204, (d_101_v1_).length(0))]})
            d_131_v28_ = (d_131_v28_).set(d_101_v1_, not(not (d_100_v0_) or (d_108_v7_)))
            d_109_v8_ = d_109_v8_
            index12_ = default__.safeIndex(204, (d_101_v1_).length(0))
            (d_101_v1_)[index12_] = (d_109_v8_).ispropersubset(d_109_v8_)
            index13_ = default__.safeIndex(204, (d_101_v1_).length(0))
            (d_101_v1_)[index13_] = not(default__.fm4(d_112_v11_, d_105_globalState_))
            d_132_v29_: D6
            d_132_v29_ = D6_DC17(d_102_v2_, d_114_v13_, d_100_v0_)
            d_133_v30_: C1
            nw6_ = C1()
            nw6_.ctor__((-225) <= ((d_113_v12_)[default__.safeIndex(378, (d_113_v12_).length(0))]), (d_132_v29_).cf40, d_109_v8_)
            d_133_v30_ = nw6_
        d_134_v31_: bool
        d_135_v32_: _dafny.MultiSet
        d_136_v33_: _dafny.Set
        d_137_v34_: int
        out4_: bool
        out5_: _dafny.MultiSet
        out6_: _dafny.Set
        out7_: int
        out4_, out5_, out6_, out7_ = default__.m0(d_107_v6_, (d_107_v6_ if d_108_v7_ else d_107_v6_), d_105_globalState_)
        d_134_v31_ = out4_
        d_135_v32_ = out5_
        d_136_v33_ = out6_
        d_137_v34_ = out7_
        d_138_v35_: _dafny.Array
        nw7_ = _dafny.Array(_dafny.Array(None, 0), 6)
        d_138_v35_ = nw7_
        index14_ = default__.safeIndex(699, (d_138_v35_).length(0))
        (d_138_v35_)[index14_] = d_113_v12_
        index15_ = default__.safeIndex(699, (d_138_v35_).length(0))
        (d_138_v35_)[index15_] = d_113_v12_
        d_139_v36_: D9
        d_139_v36_ = D9_DC23(d_137_v34_)
        d_140_v37_: _dafny.Map
        d_140_v37_ = _dafny.Map({(d_113_v12_)[default__.safeIndex(378, (d_113_v12_).length(0))]: d_102_v2_})
        hi0_ = ((d_140_v37_)[d_111_v10_] if (d_111_v10_) in (d_140_v37_) else d_111_v10_)
        for d_141_i2_ in range((d_139_v36_).cf52, hi0_):
            d_142_v38_: C5
            nw8_ = C5()
            nw8_.ctor__(d_100_v0_, d_109_v8_)
            d_142_v38_ = nw8_
            d_143_v39_: _dafny.Map
            d_143_v39_ = _dafny.Map({d_142_v38_: d_108_v7_})
            d_144_v40_: T0
            nw9_ = C2()
            nw9_.ctor__(not((d_142_v38_).f11), d_141_i2_, d_109_v8_)
            d_144_v40_ = nw9_
            d_145_v41_: _dafny.Array
            nw10_ = _dafny.Array(None, 8)
            nw10_[int(0)] = d_144_v40_
            nw10_[int(1)] = d_144_v40_
            nw10_[int(2)] = d_144_v40_
            nw10_[int(3)] = d_144_v40_
            nw10_[int(4)] = d_144_v40_
            nw10_[int(5)] = d_144_v40_
            nw10_[int(6)] = d_144_v40_
            nw10_[int(7)] = d_144_v40_
            d_145_v41_ = nw10_
            d_146_v42_: _dafny.Map
            d_146_v42_ = _dafny.Map({d_145_v41_: d_108_v7_})
            index16_ = default__.safeIndex(204, (d_101_v1_).length(0))
            (d_101_v1_)[index16_] = not (((d_143_v39_)[(d_142_v38_)] if ((d_142_v38_)) in (d_143_v39_) else ((d_146_v42_)[d_145_v41_] if (d_145_v41_) in (d_146_v42_) else True))) or ((d_142_v38_).f11)
            d_147_v43_: bool
            d_148_v44_: _dafny.MultiSet
            d_149_v45_: _dafny.Set
            d_150_v46_: int
            out8_: bool
            out9_: _dafny.MultiSet
            out10_: _dafny.Set
            out11_: int
            out8_, out9_, out10_, out11_ = default__.m0(_dafny.SeqWithoutIsStrInference([(d_101_v1_)[default__.safeIndex(204, (d_101_v1_).length(0))]]), d_107_v6_, d_105_globalState_)
            d_147_v43_ = out8_
            d_148_v44_ = out9_
            d_149_v45_ = out10_
            d_150_v46_ = out11_
            d_151_v47_: _dafny.Seq
            d_151_v47_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet(d_107_v6_)])
            d_152_v48_: str
            d_152_v48_ = _dafny.CodePoint('p')
            d_153_v49_: _dafny.Map
            d_153_v49_ = _dafny.Map({default__.fm2((d_142_v38_).fm5(d_149_v45_, len(d_151_v47_), d_150_v46_, d_105_globalState_), d_141_i2_, d_152_v48_, d_150_v46_, d_105_globalState_): d_147_v43_})
            d_154_v50_: _dafny.Map
            d_154_v50_ = _dafny.Map({d_147_v43_: d_108_v7_})
            d_153_v49_ = (d_153_v49_).set(default__.fm2(((d_154_v50_)[(d_142_v38_).f11] if ((d_142_v38_).f11) in (d_154_v50_) else d_147_v43_), d_102_v2_, d_152_v48_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ivq"))), d_105_globalState_), d_134_v31_)
            d_155_v51_: C0
            nw11_ = C0()
            nw11_.ctor__()
            d_155_v51_ = nw11_
            d_155_v51_ = d_155_v51_
        d_156_v52_: str
        d_156_v52_ = _dafny.CodePoint('f')
        d_157_v53_: _dafny.Seq
        d_157_v53_ = _dafny.SeqWithoutIsStrInference([d_114_v13_, d_114_v13_, default__.fm2(d_100_v0_, d_111_v10_, d_156_v52_, len(d_140_v37_), d_105_globalState_), d_114_v13_])
        d_158_v54_: _dafny.Array
        nw12_ = _dafny.Array(None, 24)
        nw12_[int(0)] = (d_157_v53_)[default__.safeIndex(d_111_v10_, len(d_157_v53_))]
        nw12_[int(1)] = d_114_v13_
        nw12_[int(2)] = d_114_v13_
        nw12_[int(3)] = d_114_v13_
        nw12_[int(4)] = ((d_114_v13_) + (((d_114_v13_).set(default__.safeIndex((d_113_v12_)[default__.safeIndex(378, (d_113_v12_).length(0))], len(d_114_v13_)), d_156_v52_)).set(default__.safeIndex(d_111_v10_, len((d_114_v13_).set(default__.safeIndex((d_113_v12_)[default__.safeIndex(378, (d_113_v12_).length(0))], len(d_114_v13_)), d_156_v52_))), (d_114_v13_)[default__.safeIndex(d_102_v2_, len(d_114_v13_))]))).set(default__.safeIndex((d_113_v12_)[default__.safeIndex(378, (d_113_v12_).length(0))], len((d_114_v13_) + (((d_114_v13_).set(default__.safeIndex((d_113_v12_)[default__.safeIndex(378, (d_113_v12_).length(0))], len(d_114_v13_)), d_156_v52_)).set(default__.safeIndex(d_111_v10_, len((d_114_v13_).set(default__.safeIndex((d_113_v12_)[default__.safeIndex(378, (d_113_v12_).length(0))], len(d_114_v13_)), d_156_v52_))), (d_114_v13_)[default__.safeIndex(d_102_v2_, len(d_114_v13_))])))), default__.fm1(d_105_globalState_))
        nw12_[int(5)] = d_114_v13_
        nw12_[int(6)] = d_114_v13_
        nw12_[int(7)] = _dafny.SeqWithoutIsStrInference([d_156_v52_ for d_159_i3_ in range(default__.abs(629))])
        nw12_[int(8)] = default__.fm2(d_108_v7_, d_111_v10_, d_156_v52_, d_137_v34_, d_105_globalState_)
        nw12_[int(9)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "muykribt"))
        nw12_[int(10)] = (d_114_v13_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "atwbgc")))
        nw12_[int(11)] = d_114_v13_
        nw12_[int(12)] = d_114_v13_
        nw12_[int(13)] = d_114_v13_
        nw12_[int(14)] = d_114_v13_
        nw12_[int(15)] = d_114_v13_
        nw12_[int(16)] = (d_114_v13_) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lx"))).set(default__.safeIndex(d_111_v10_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lx")))), d_156_v52_))
        nw12_[int(17)] = (d_114_v13_).set(default__.safeIndex(d_111_v10_, len(d_114_v13_)), d_156_v52_)
        nw12_[int(18)] = d_114_v13_
        nw12_[int(19)] = (d_114_v13_) + (d_114_v13_)
        nw12_[int(20)] = (d_114_v13_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yyix")))
        nw12_[int(21)] = (d_114_v13_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hkxekp")))
        nw12_[int(22)] = (d_114_v13_) + (default__.fm2((d_101_v1_)[default__.safeIndex(204, (d_101_v1_).length(0))], 228, d_156_v52_, d_102_v2_, d_105_globalState_))
        nw12_[int(23)] = (d_114_v13_ if d_100_v0_ else _dafny.SeqWithoutIsStrInference([d_156_v52_ for d_160_i4_ in range(default__.abs(-159))]))
        d_158_v54_ = nw12_
        index17_ = default__.safeIndex(830, (d_158_v54_).length(0))
        (d_158_v54_)[index17_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hythm"))
        index18_ = default__.safeIndex(830, (d_158_v54_).length(0))
        index19_ = default__.safeIndex(378, (d_113_v12_).length(0))
        rhs12_ = (d_114_v13_) + (default__.fm2(d_100_v0_, (D0_DC1(d_100_v0_, (d_113_v12_)[default__.safeIndex(378, (d_113_v12_).length(0))])).cf2, d_156_v52_, 432, d_105_globalState_))
        rhs13_ = (d_113_v12_)[default__.safeIndex(378, (d_113_v12_).length(0))]
        lhs10_ = d_158_v54_
        lhs11_ = default__.safeIndex(830, (d_158_v54_).length(0))
        lhs12_ = d_113_v12_
        lhs13_ = default__.safeIndex(378, (d_113_v12_).length(0))
        lhs10_[lhs11_] = rhs12_
        lhs12_[lhs13_] = rhs13_
        hi1_ = (0) - ((d_102_v2_) * (d_111_v10_))
        for d_161_i5_ in range((d_113_v12_)[default__.safeIndex(378, (d_113_v12_).length(0))], hi1_):
            d_162_v55_: _dafny.Array
            nw13_ = _dafny.Array(None, 12)
            d_162_v55_ = nw13_
            d_163_v56_: _dafny.Map
            d_163_v56_ = _dafny.Map({366: d_162_v55_})
            d_163_v56_ = d_163_v56_
            d_164_v57_: _dafny.Array
            nw14_ = _dafny.Array(None, 28)
            d_164_v57_ = nw14_
            nw15_ = _dafny.Array(None, 26)
            d_164_v57_ = nw15_
            index20_ = default__.safeIndex(204, (d_101_v1_).length(0))
            (d_101_v1_)[index20_] = ((d_107_v6_) + ((d_107_v6_).set(default__.safeIndex(314, len(d_107_v6_)), d_134_v31_))) < (d_107_v6_)
            d_165_v58_: _dafny.Map
            d_165_v58_ = _dafny.Map({d_161_i5_: (d_101_v1_)[default__.safeIndex(204, (d_101_v1_).length(0))]})
            d_102_v2_ = (d_111_v10_) - (len(d_165_v58_))
        d_106_v5_ = (d_106_v5_).set((d_108_v7_) not in (d_107_v6_), (0) - ((d_113_v12_)[default__.safeIndex(378, (d_113_v12_).length(0))]))
        d_112_v11_ = d_112_v11_
        d_166_v59_: T1
        nw16_ = C3()
        nw16_.ctor__(d_134_v31_)
        d_166_v59_ = nw16_
        d_167_v60_: _dafny.Seq
        d_167_v60_ = _dafny.SeqWithoutIsStrInference([d_166_v59_, d_166_v59_, d_166_v59_, d_166_v59_, d_166_v59_])
        d_167_v60_ = d_167_v60_
        d_168_v61_: _dafny.MultiSet
        d_168_v61_ = d_135_v32_
        pat_let_tv13_ = d_113_v12_
        pat_let_tv14_ = d_113_v12_
        def lambda2_(source5_):
            d_169___mcc_h0_ = source5_
            d_170_cf41_ = d_169___mcc_h0_
            return (381) != ((pat_let_tv14_)[default__.safeIndex(378, (pat_let_tv13_).length(0))])

        d_108_v7_ = lambda2_(d_168_v61_)
        d_171_v62_: bool
        d_172_v63_: _dafny.MultiSet
        d_173_v64_: _dafny.Set
        d_174_v65_: int
        out12_: bool
        out13_: _dafny.MultiSet
        out14_: _dafny.Set
        out15_: int
        out12_, out13_, out14_, out15_ = default__.m0(d_107_v6_, d_107_v6_, d_105_globalState_)
        d_171_v62_ = out12_
        d_172_v63_ = out13_
        d_173_v64_ = out14_
        d_174_v65_ = out15_
        _dafny.print(_dafny.string_of(d_100_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_101_v1_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_101_v1_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_101_v1_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_101_v1_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_101_v1_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_101_v1_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_101_v1_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_101_v1_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_101_v1_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_101_v1_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_101_v1_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_101_v1_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_101_v1_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_101_v1_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_101_v1_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_101_v1_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_101_v1_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_101_v1_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_101_v1_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_101_v1_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_101_v1_)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_101_v1_)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_101_v1_)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_101_v1_)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_101_v1_)[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_101_v1_)[25]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_102_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_103_v3_) == (_dafny.MultiSet([599]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_104_v4_) == (_dafny.SeqWithoutIsStrInference([599]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_105_globalState_).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_105_globalState_.f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_105_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_105_globalState_).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_105_globalState_).f4)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_105_globalState_).f4)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_105_globalState_).f4)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_105_globalState_).f4)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_105_globalState_).f4)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_105_globalState_).f4)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_105_globalState_).f4)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_105_globalState_).f4)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_105_globalState_).f4)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_105_globalState_).f4)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_105_globalState_).f4)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_105_globalState_).f4)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_105_globalState_).f4)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_105_globalState_).f4)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_105_globalState_).f4)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_105_globalState_).f4)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_105_globalState_).f4)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_105_globalState_).f4)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_105_globalState_).f4)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_105_globalState_).f4)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_105_globalState_).f4)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_105_globalState_).f4)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_105_globalState_).f4)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_105_globalState_).f4)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_105_globalState_).f4)[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_105_globalState_).f4)[25]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_105_globalState_.f5)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_105_globalState_.f5)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_105_globalState_.f5)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_105_globalState_.f5)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_105_globalState_.f5)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_105_globalState_.f5)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_105_globalState_.f5)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_105_globalState_.f5)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_105_globalState_.f5)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_105_globalState_.f5)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_105_globalState_.f5)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_105_globalState_.f5)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_105_globalState_.f5)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_105_globalState_.f5)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_105_globalState_.f5)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_105_globalState_.f5)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_105_globalState_.f5)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_105_globalState_.f5)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_105_globalState_.f5)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_105_globalState_.f5)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_105_globalState_.f5)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_105_globalState_.f5)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_105_globalState_.f5)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_105_globalState_.f5)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_105_globalState_.f5)[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_105_globalState_.f5)[25]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_105_globalState_).f6) == (_dafny.MultiSet([599]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_105_globalState_).f7) == (_dafny.SeqWithoutIsStrInference([599]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_105_globalState_).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_105_globalState_).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_106_v5_) == (_dafny.Map({True: 0, False: 0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_107_v6_) == (_dafny.SeqWithoutIsStrInference([False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_108_v7_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_109_v8_) == (_dafny.MultiSet([False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_110_v9_) == (_dafny.Set({}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_111_v10_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_112_v11_) == (_dafny.Set({599}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_113_v12_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_114_v13_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_115_v14_) == (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ibqqjdrauoj")): 0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_134_v31_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_135_v32_) == (_dafny.MultiSet([False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_136_v33_) == (_dafny.Set({}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_137_v34_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_138_v35_)[3])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_139_v36_).cf52))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_140_v37_) == (_dafny.Map({0: 192}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_156_v52_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_157_v53_) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fe")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fe")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k'), _dafny.CodePoint('g'), _dafny.CodePoint('i'), _dafny.CodePoint('n'), _dafny.CodePoint('o'), _dafny.CodePoint('a'), _dafny.CodePoint('x'), _dafny.CodePoint('j'), _dafny.CodePoint('k'), _dafny.CodePoint('d'), _dafny.CodePoint('d')]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fe"))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_158_v54_)[0]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_158_v54_)[1]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_158_v54_)[2]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_158_v54_)[3]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_158_v54_)[4]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_158_v54_)[5]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_158_v54_)[6]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_158_v54_)[7]) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_158_v54_)[8]) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k'), _dafny.CodePoint('g'), _dafny.CodePoint('i'), _dafny.CodePoint('n'), _dafny.CodePoint('o'), _dafny.CodePoint('a'), _dafny.CodePoint('x'), _dafny.CodePoint('j'), _dafny.CodePoint('k'), _dafny.CodePoint('d'), _dafny.CodePoint('d')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_158_v54_)[9]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_158_v54_)[10]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_158_v54_)[11]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_158_v54_)[12]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_158_v54_)[13]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_158_v54_)[14]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_158_v54_)[15]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_158_v54_)[16]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_158_v54_)[17]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_158_v54_)[18]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_158_v54_)[19]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_158_v54_)[20]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_158_v54_)[21]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_158_v54_)[22]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_158_v54_)[23]) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_167_v60_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_168_v61_)) == (_dafny.MultiSet([False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_171_v62_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_172_v63_) == (_dafny.MultiSet([False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_173_v64_) == (_dafny.Set({}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_174_v65_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D0_DC1)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D0_DC2)

class D0_DC1(D0, NamedTuple('DC1', [('cf1', Any), ('cf2', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf1)}, {_dafny.string_of(self.cf2)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf1 == __o.cf1 and self.cf2 == __o.cf2
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC0(D0, NamedTuple('DC0', [('cf0', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC0({_dafny.string_of(self.cf0)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC0) and self.cf0 == __o.cf0
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC2(D0, NamedTuple('DC2', [('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2({_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2) and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()


class D1:
    @classmethod
    def default(cls, ):
        return lambda: D1_DC4(False, False, int(0), _dafny.Seq({}), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D1_DC5)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)

class D1_DC4(D1, NamedTuple('DC4', [('cf5', Any), ('cf6', Any), ('cf7', Any), ('cf8', Any), ('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf5)}, {_dafny.string_of(self.cf6)}, {_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)}, {_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf5 == __o.cf5 and self.cf6 == __o.cf6 and self.cf7 == __o.cf7 and self.cf8 == __o.cf8 and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC5(D1, NamedTuple('DC5', [('cf10', Any), ('cf11', Any), ('cf12', Any), ('cf13', Any), ('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf10)}, {_dafny.string_of(self.cf11)}, {_dafny.string_of(self.cf12)}, {_dafny.string_of(self.cf13)}, {_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf10 == __o.cf10 and self.cf11 == __o.cf11 and self.cf12 == __o.cf12 and self.cf13 == __o.cf13 and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC3(D1, NamedTuple('DC3', [('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)

class D2_DC6(D2, NamedTuple('DC6', [('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC8(_dafny.CodePoint('D'), False, False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D3_DC8)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D3_DC7)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D3_DC9)

class D3_DC8(D3, NamedTuple('DC8', [('cf17', Any), ('cf18', Any), ('cf19', Any), ('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC8({_dafny.string_of(self.cf17)}, {_dafny.string_of(self.cf18)}, {_dafny.string_of(self.cf19)}, {self.cf20.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC8) and self.cf17 == __o.cf17 and self.cf18 == __o.cf18 and self.cf19 == __o.cf19 and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC7(D3, NamedTuple('DC7', [('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC7({self.cf16.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC7) and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC9(D3, NamedTuple('DC9', [('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC11(D0.default()(), _dafny.Array(None, 0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D4_DC11)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D4_DC10)

class D4_DC11(D4, NamedTuple('DC11', [('cf23', Any), ('cf24', Any), ('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC11({_dafny.string_of(self.cf23)}, {_dafny.string_of(self.cf24)}, {_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC11) and self.cf23 == __o.cf23 and self.cf24 == __o.cf24 and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC10(D4, NamedTuple('DC10', [('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC10({_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC10) and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC13(int(0), _dafny.Array(None, 0), _dafny.Array(None, 0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D5_DC13)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D5_DC14)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D5_DC12)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D5_DC15)

class D5_DC13(D5, NamedTuple('DC13', [('cf27', Any), ('cf28', Any), ('cf29', Any), ('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC13({_dafny.string_of(self.cf27)}, {_dafny.string_of(self.cf28)}, {_dafny.string_of(self.cf29)}, {_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC13) and self.cf27 == __o.cf27 and self.cf28 == __o.cf28 and self.cf29 == __o.cf29 and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC14(D5, NamedTuple('DC14', [('cf31', Any), ('cf32', Any), ('cf33', Any), ('cf34', Any), ('cf35', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC14({_dafny.string_of(self.cf31)}, {_dafny.string_of(self.cf32)}, {_dafny.string_of(self.cf33)}, {_dafny.string_of(self.cf34)}, {_dafny.string_of(self.cf35)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC14) and self.cf31 == __o.cf31 and self.cf32 == __o.cf32 and self.cf33 == __o.cf33 and self.cf34 == __o.cf34 and self.cf35 == __o.cf35
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC12(D5, NamedTuple('DC12', [('cf26', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC12({_dafny.string_of(self.cf26)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC12) and self.cf26 == __o.cf26
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC15(D5, NamedTuple('DC15', [('cf36', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC15({_dafny.string_of(self.cf36)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC15) and self.cf36 == __o.cf36
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC17(int(0), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D6_DC17)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D6_DC16)

class D6_DC17(D6, NamedTuple('DC17', [('cf38', Any), ('cf39', Any), ('cf40', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC17({_dafny.string_of(self.cf38)}, {self.cf39.VerbatimString(True)}, {_dafny.string_of(self.cf40)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC17) and self.cf38 == __o.cf38 and self.cf39 == __o.cf39 and self.cf40 == __o.cf40
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC16(D6, NamedTuple('DC16', [('cf37', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC16({_dafny.string_of(self.cf37)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC16) and self.cf37 == __o.cf37
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D7_DC18)

class D7_DC18(D7, NamedTuple('DC18', [('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC18({_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC18) and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC20(_dafny.CodePoint('D'), False, False, _dafny.Map({}), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D8_DC20)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D8_DC19)

class D8_DC20(D8, NamedTuple('DC20', [('cf43', Any), ('cf44', Any), ('cf45', Any), ('cf46', Any), ('cf47', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC20({_dafny.string_of(self.cf43)}, {_dafny.string_of(self.cf44)}, {_dafny.string_of(self.cf45)}, {_dafny.string_of(self.cf46)}, {_dafny.string_of(self.cf47)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC20) and self.cf43 == __o.cf43 and self.cf44 == __o.cf44 and self.cf45 == __o.cf45 and self.cf46 == __o.cf46 and self.cf47 == __o.cf47
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC19(D8, NamedTuple('DC19', [('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC19({_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC19) and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC22(int(0), False, _dafny.Map({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D9_DC22)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D9_DC23)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D9_DC24)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D9_DC21)

class D9_DC22(D9, NamedTuple('DC22', [('cf49', Any), ('cf50', Any), ('cf51', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC22({_dafny.string_of(self.cf49)}, {_dafny.string_of(self.cf50)}, {_dafny.string_of(self.cf51)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC22) and self.cf49 == __o.cf49 and self.cf50 == __o.cf50 and self.cf51 == __o.cf51
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC23(D9, NamedTuple('DC23', [('cf52', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC23({_dafny.string_of(self.cf52)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC23) and self.cf52 == __o.cf52
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC24(D9, NamedTuple('DC24', [('cf53', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC24({_dafny.string_of(self.cf53)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC24) and self.cf53 == __o.cf53
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC21(D9, NamedTuple('DC21', [('cf48', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC21({_dafny.string_of(self.cf48)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC21) and self.cf48 == __o.cf48
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: None
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D10_DC25)

class D10_DC25(D10, NamedTuple('DC25', [('cf54', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC25({_dafny.string_of(self.cf54)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC25) and self.cf54 == __o.cf54
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass

class T1:
    pass
    def m1(self, p0, p1, p2, globalState):
        pass


class T2:
    pass

class GlobalState:
    def  __init__(self):
        self.f1: int = int(0)
        self.f5: _dafny.Array = _dafny.Array(None, 0)
        self._f0: bool = False
        self._f2: int = int(0)
        self._f3: bool = False
        self._f4: _dafny.Array = _dafny.Array(None, 0)
        self._f6: _dafny.MultiSet = _dafny.MultiSet({})
        self._f7: _dafny.Seq = _dafny.Seq({})
        self._f8: bool = False
        self._f9: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9):
        (self)._f0 = f0
        (self).f1 = f1
        (self)._f2 = f2
        (self)._f3 = f3
        (self)._f4 = f4
        (self).f5 = f5
        (self)._f6 = f6
        (self)._f7 = f7
        (self)._f8 = f8
        (self)._f9 = f9

    @property
    def f0(self):
        return self._f0
    @property
    def f2(self):
        return self._f2
    @property
    def f3(self):
        return self._f3
    @property
    def f4(self):
        return self._f4
    @property
    def f6(self):
        return self._f6
    @property
    def f7(self):
        return self._f7
    @property
    def f8(self):
        return self._f8
    @property
    def f9(self):
        return self._f9

class C0(T2):
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self):
        pass
        pass

    def fm5(self, p0, p1, p2, globalState):
        return True

    def fm6(self, p0, p1, globalState):
        return default__.safeDivisionInt(len((_dafny.SeqWithoutIsStrInference([len(_dafny.Set({False, True})), -766])) + (_dafny.SeqWithoutIsStrInference([131]))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "igq"))))

    def fm13(self, p0, globalState):
        return (_dafny.CodePoint('c')) not in (((D3_DC8(_dafny.CodePoint('h'), True, False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xnohburd")))).cf20) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wami"))))


class C1(T0):
    def  __init__(self):
        self._f10: _dafny.MultiSet = _dafny.MultiSet({})
        self._f14: bool = False
        self._f15: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    @property
    def f10(self):
        return self._f10
    def ctor__(self, f14, f15, f10):
        (self)._f14 = f14
        (self)._f15 = f15
        (self)._f10 = f10

    def m5(self, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: int = int(0)
        r2: _dafny.Map = _dafny.Map({})
        d_175_v0_: bool
        d_175_v0_ = False
        d_175_v0_ = d_175_v0_
        d_176_v1_: C0
        nw17_ = C0()
        nw17_.ctor__()
        d_176_v1_ = nw17_
        d_177_v2_: D4
        d_177_v2_ = D4_DC10(d_176_v1_)
        d_178_v3_: int
        d_178_v3_ = 979
        d_179_v4_: str
        d_179_v4_ = _dafny.CodePoint('a')
        rhs14_ = (d_177_v2_).cf22
        rhs15_ = (default__.fm2(d_175_v0_, d_178_v3_, d_179_v4_, d_178_v3_, globalState)) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w")))
        d_176_v1_ = rhs14_
        d_175_v0_ = rhs15_
        d_180_v5_: _dafny.Seq
        d_180_v5_ = _dafny.SeqWithoutIsStrInference([d_178_v3_])
        hi2_ = (d_180_v5_)[default__.safeIndex(d_178_v3_, len(d_180_v5_))]
        for d_181_i0_ in range(d_178_v3_, hi2_):
            d_182_v6_: _dafny.Array
            nw18_ = _dafny.Array(int(0), 18)
            d_182_v6_ = nw18_
            d_183_v7_: _dafny.Map
            d_183_v7_ = _dafny.Map({True: d_182_v6_})
            d_184_v8_: _dafny.Array
            nw19_ = _dafny.Array(None, 4)
            nw19_[int(0)] = d_182_v6_
            nw19_[int(1)] = d_182_v6_
            nw19_[int(2)] = ((d_183_v7_)[d_175_v0_] if (d_175_v0_) in (d_183_v7_) else d_182_v6_)
            nw19_[int(3)] = d_182_v6_
            d_184_v8_ = nw19_
            index21_ = default__.safeIndex(734, (d_184_v8_).length(0))
            (d_184_v8_)[index21_] = d_182_v6_
            index22_ = default__.safeIndex(734, (d_184_v8_).length(0))
            (d_184_v8_)[index22_] = (d_182_v6_ if False else d_182_v6_)
            d_185_v9_: _dafny.Array
            def lambda3_(d_186_v4_):
                def lambda4_(d_187_i1_):
                    return (_dafny.SeqWithoutIsStrInference([d_186_v4_])) + (_dafny.SeqWithoutIsStrInference([d_186_v4_ for d_188_i2_ in range(default__.abs(214))]))

                return lambda4_

            init0_ = lambda3_(d_179_v4_)
            nw20_ = _dafny.Array(None, 10)
            for i0_0_ in range(nw20_.length(0)):
                nw20_[i0_0_] = init0_(i0_0_)
            d_185_v9_ = nw20_
            d_189_v10_: _dafny.Seq
            d_189_v10_ = _dafny.SeqWithoutIsStrInference([d_179_v4_])
            d_190_v11_: _dafny.MultiSet
            d_190_v11_ = _dafny.MultiSet([d_181_i0_, d_178_v3_, len(d_189_v10_), d_181_i0_, len(d_180_v5_)])
            d_191_v12_: _dafny.Seq
            d_191_v12_ = _dafny.SeqWithoutIsStrInference([(d_189_v10_)[default__.safeIndex((d_190_v11_).cardinality, len(d_189_v10_))], d_179_v4_, d_179_v4_, d_179_v4_, _dafny.CodePoint('s')])
            index23_ = default__.safeIndex(239, (d_185_v9_).length(0))
            (d_185_v9_)[index23_] = (d_191_v12_) + (d_191_v12_)
            index24_ = default__.safeIndex(239, (d_185_v9_).length(0))
            (d_185_v9_)[index24_] = d_191_v12_
            index25_ = default__.safeIndex(239, (d_185_v9_).length(0))
            (d_185_v9_)[index25_] = _dafny.SeqWithoutIsStrInference([d_179_v4_ for d_192_i3_ in range(default__.abs(-892))])
            d_193_v13_: C0
            nw21_ = C0()
            nw21_.ctor__()
            d_193_v13_ = nw21_
        d_194_v14_: _dafny.Set
        d_194_v14_ = _dafny.Set({d_178_v3_})
        d_195_v16_: _dafny.Seq
        d_195_v16_ = _dafny.SeqWithoutIsStrInference([d_194_v14_, _dafny.Set({len(d_180_v5_), d_178_v3_})])
        d_196_v17_: _dafny.Array
        nw22_ = _dafny.Array(None, 17)
        nw22_[int(0)] = d_194_v14_
        nw22_[int(1)] = (_dafny.Set({d_178_v3_, (d_176_v1_).fm6(d_175_v0_, d_178_v3_, globalState)})) | (d_194_v14_)
        nw22_[int(2)] = d_194_v14_
        nw22_[int(3)] = d_194_v14_
        nw22_[int(4)] = d_194_v14_
        nw22_[int(5)] = d_194_v14_
        nw22_[int(6)] = (d_194_v14_) | (d_194_v14_)
        def iife14_():
            coll10_ = _dafny.Set()
            compr_10_: int
            for compr_10_ in _dafny.IntegerRange(35, 104):
                d_197_v15_: int = compr_10_
                if ((35) <= (d_197_v15_)) and ((d_197_v15_) < (104)):
                    coll10_ = coll10_.union(_dafny.Set([default__.safeModuloInt(d_197_v15_, (((self).f10)[(self).f14] if ((self).f14) in ((self).f10) else len(_dafny.Set({(self).f14}))))]))
            return _dafny.Set(coll10_)
        nw22_[int(7)] = iife14_()
        
        nw22_[int(8)] = d_194_v14_
        nw22_[int(9)] = (d_194_v14_) | (d_194_v14_)
        nw22_[int(10)] = d_194_v14_
        nw22_[int(11)] = d_194_v14_
        nw22_[int(12)] = (d_195_v16_)[default__.safeIndex((0) - (d_178_v3_), len(d_195_v16_))]
        nw22_[int(13)] = d_194_v14_
        nw22_[int(14)] = d_194_v14_
        nw22_[int(15)] = d_194_v14_
        nw22_[int(16)] = d_194_v14_
        d_196_v17_ = nw22_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_196_v17_).length(0)):
            d_198_i4_: int = guard_loop_0_
            if (True) and (((0) <= (d_198_i4_)) and ((d_198_i4_) < ((d_196_v17_).length(0)))):
                (d_196_v17_)[(d_198_i4_)] = d_194_v14_
        d_199_i5_: int
        d_199_i5_ = 0
        with _dafny.label("1"):
            while default__.fm4((d_194_v14_ if (self).f14 else d_194_v14_), globalState):
                with _dafny.c_label("1"):
                    if (d_199_i5_) >= (100):
                        raise _dafny.Break("1")
                    d_199_i5_ = (d_199_i5_) + (1)
                    r1 = 618
                    if (self).f15:
                        d_200_v18_: T2
                        nw23_ = C0()
                        nw23_.ctor__()
                        d_200_v18_ = nw23_
                        d_200_v18_ = d_200_v18_
                        d_201_v19_: _dafny.Seq
                        d_201_v19_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "efljrb"))
                        d_202_v20_: _dafny.Map
                        d_202_v20_ = _dafny.Map({d_201_v19_: True})
                        d_203_v22_: _dafny.Seq
                        d_203_v22_ = _dafny.SeqWithoutIsStrInference([d_201_v19_])
                        d_204_v23_: _dafny.Map
                        d_204_v23_ = _dafny.Map({(self).f14: (self).f15})
                        d_205_v24_: _dafny.Seq
                        d_205_v24_ = _dafny.SeqWithoutIsStrInference([d_204_v23_, d_204_v23_, _dafny.Map({d_175_v0_: (self).f15}), d_204_v23_, d_204_v23_])
                        d_206_v25_: _dafny.Seq
                        d_206_v25_ = _dafny.SeqWithoutIsStrInference([(d_205_v24_)[default__.safeIndex(d_178_v3_, len(d_205_v24_))]])
                        d_207_v26_: _dafny.Map
                        d_207_v26_ = _dafny.Map({((d_205_v24_)[default__.safeIndex(d_178_v3_, len(d_205_v24_))]).set(d_175_v0_, (self).f14): 433})
                        d_208_v28_: _dafny.Array
                        nw24_ = _dafny.Array(None, 27)
                        def iife15_():
                            coll11_ = _dafny.Map()
                            compr_11_: _dafny.Seq
                            for compr_11_ in (d_203_v22_).Elements:
                                d_209_v21_: _dafny.Seq = compr_11_
                                if (d_209_v21_) in (d_203_v22_):
                                    coll11_[d_209_v21_] = d_175_v0_
                            return _dafny.Map(coll11_)
                        nw24_[int(0)] = len((d_202_v20_) | (iife15_()
                        ))
                        nw24_[int(1)] = d_178_v3_
                        nw24_[int(2)] = d_178_v3_
                        nw24_[int(3)] = (d_178_v3_) * (d_178_v3_)
                        nw24_[int(4)] = ((self).f10).cardinality
                        nw24_[int(5)] = (d_200_v18_).fm6((self).f14, len(d_201_v19_), globalState)
                        nw24_[int(6)] = d_178_v3_
                        nw24_[int(7)] = d_178_v3_
                        nw24_[int(8)] = d_178_v3_
                        nw24_[int(9)] = d_178_v3_
                        nw24_[int(10)] = d_178_v3_
                        nw24_[int(11)] = len(d_201_v19_)
                        nw24_[int(12)] = 350
                        nw24_[int(13)] = len(_dafny.SeqWithoutIsStrInference([d_179_v4_ for d_210_i6_ in range(default__.abs(439))]))
                        nw24_[int(14)] = d_178_v3_
                        nw24_[int(15)] = d_178_v3_
                        nw24_[int(16)] = (d_176_v1_).fm6(not((self).f15), d_178_v3_, globalState)
                        nw24_[int(17)] = len(d_194_v14_)
                        nw24_[int(18)] = d_178_v3_
                        nw24_[int(19)] = default__.safeModuloInt(len((d_180_v5_).set(default__.safeIndex(d_178_v3_, len(d_180_v5_)), d_178_v3_)), d_178_v3_)
                        nw24_[int(20)] = (d_176_v1_).fm6((self).f15, d_178_v3_, globalState)
                        nw24_[int(21)] = d_178_v3_
                        nw24_[int(22)] = d_178_v3_
                        nw24_[int(23)] = default__.safeModuloInt(d_178_v3_, d_178_v3_)
                        nw24_[int(24)] = (d_178_v3_) + (d_178_v3_)
                        def iife16_():
                            coll12_ = _dafny.Set()
                            compr_12_: int
                            for compr_12_ in _dafny.IntegerRange(218, 629):
                                d_211_v27_: int = compr_12_
                                if ((218) <= (d_211_v27_)) and ((d_211_v27_) < (629)):
                                    coll12_ = coll12_.union(_dafny.Set([(d_211_v27_) * (d_178_v3_)]))
                            return _dafny.Set(coll12_)
                        nw24_[int(25)] = ((d_207_v26_)[d_204_v23_] if (d_204_v23_) in (d_207_v26_) else (0) - (len(iife16_()
                        )))
                        nw24_[int(26)] = d_178_v3_
                        d_208_v28_ = nw24_
                        index26_ = default__.safeIndex(329, (d_208_v28_).length(0))
                        (d_208_v28_)[index26_] = 883
                        index27_ = default__.safeIndex(329, (d_208_v28_).length(0))
                        (d_208_v28_)[index27_] = d_178_v3_
                        d_175_v0_ = False
                        d_212_v29_: _dafny.Array
                        def lambda5_(d_213_i7_):
                            return (self).f15

                        init1_ = lambda5_
                        nw25_ = _dafny.Array(None, 9)
                        for i0_1_ in range(nw25_.length(0)):
                            nw25_[i0_1_] = init1_(i0_1_)
                        d_212_v29_ = nw25_
                        r0 = d_212_v29_
                        rhs16_ = (((d_208_v28_)[default__.safeIndex(329, (d_208_v28_).length(0))] if (self).f15 else len(d_201_v19_))) + (d_178_v3_)
                        rhs17_ = (self).f14
                        lhs14_ = globalState
                        lhs14_.f1 = rhs16_
                        d_175_v0_ = rhs17_
                    elif True:
                        d_214_v30_: _dafny.Array
                        nw26_ = _dafny.Array(int(0), 3)
                        d_214_v30_ = nw26_
                        index28_ = default__.safeIndex(249, (d_214_v30_).length(0))
                        (d_214_v30_)[index28_] = default__.safeDivisionInt(d_178_v3_, ((self).f10).cardinality)
                        index29_ = default__.safeIndex(249, (d_214_v30_).length(0))
                        (d_214_v30_)[index29_] = (((self).f10)[(d_178_v3_) == (d_178_v3_)] if ((d_178_v3_) == (d_178_v3_)) in ((self).f10) else d_178_v3_)
                        d_215_v31_: _dafny.Seq
                        d_215_v31_ = _dafny.SeqWithoutIsStrInference([(self).f14, (self).f14])
                        rhs18_ = (d_215_v31_) + (d_215_v31_)
                        rhs19_ = d_175_v0_
                        rhs20_ = 490
                        lhs15_ = globalState
                        d_215_v31_ = rhs18_
                        d_175_v0_ = rhs19_
                        lhs15_.f1 = rhs20_
                        d_216_v32_: D1
                        d_216_v32_ = D1_DC5(d_178_v3_, (self).f14, 166, not (d_175_v0_) or (d_175_v0_), d_178_v3_)
                        d_217_v33_: _dafny.Seq
                        d_217_v33_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ko"))
                        d_216_v32_ = D1_DC5((d_214_v30_)[default__.safeIndex(249, (d_214_v30_).length(0))], (self).f14, default__.safeModuloInt(791, (d_214_v30_)[default__.safeIndex(249, (d_214_v30_).length(0))]), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "scqxnyhq"))) <= (d_217_v33_), len((d_215_v31_) + (d_215_v31_)))
                        d_218_v34_: _dafny.MultiSet
                        d_218_v34_ = _dafny.MultiSet([330])
                        d_178_v3_ = ((d_218_v34_)[(d_214_v30_)[default__.safeIndex(249, (d_214_v30_).length(0))]] if ((d_214_v30_)[default__.safeIndex(249, (d_214_v30_).length(0))]) in (d_218_v34_) else (d_214_v30_)[default__.safeIndex(249, (d_214_v30_).length(0))])
                        d_219_v35_: _dafny.Array
                        def lambda6_(d_220_i8_):
                            return (self).f14

                        init2_ = lambda6_
                        nw27_ = _dafny.Array(None, 28)
                        for i0_2_ in range(nw27_.length(0)):
                            nw27_[i0_2_] = init2_(i0_2_)
                        d_219_v35_ = nw27_
                        index30_ = default__.safeIndex(411, (d_219_v35_).length(0))
                        (d_219_v35_)[index30_] = d_175_v0_
                        index31_ = default__.safeIndex(411, (d_219_v35_).length(0))
                        rhs21_ = d_175_v0_
                        rhs22_ = not((self).f14)
                        rhs23_ = d_217_v33_
                        rhs24_ = (d_214_v30_)[default__.safeIndex(249, (d_214_v30_).length(0))]
                        rhs25_ = (d_214_v30_)[default__.safeIndex(249, (d_214_v30_).length(0))]
                        lhs16_ = d_219_v35_
                        lhs17_ = default__.safeIndex(411, (d_219_v35_).length(0))
                        lhs18_ = globalState
                        lhs19_ = globalState
                        lhs16_[lhs17_] = rhs21_
                        d_175_v0_ = rhs22_
                        d_217_v33_ = rhs23_
                        lhs18_.f1 = rhs24_
                        lhs19_.f1 = rhs25_
                    d_221_v36_: _dafny.MultiSet
                    d_221_v36_ = _dafny.MultiSet(d_180_v5_)
                    source6_ = d_221_v36_
                    d_222___mcc_h0_ = source6_
                    d_223_cf15_ = d_222___mcc_h0_
                    d_224_v37_: C0
                    nw28_ = C0()
                    nw28_.ctor__()
                    d_224_v37_ = nw28_
                    d_225_v38_: C0
                    nw29_ = C0()
                    nw29_.ctor__()
                    d_225_v38_ = nw29_
                    d_226_v39_: _dafny.Array
                    nw30_ = _dafny.Array(None, 20)
                    nw30_[int(0)] = (self).f15
                    nw30_[int(1)] = (self).f15
                    nw30_[int(2)] = d_175_v0_
                    nw30_[int(3)] = d_175_v0_
                    nw30_[int(4)] = (self).f15
                    nw30_[int(5)] = (self).f14
                    nw30_[int(6)] = default__.fm4(d_194_v14_, globalState)
                    nw30_[int(7)] = False
                    nw30_[int(8)] = d_175_v0_
                    nw30_[int(9)] = d_175_v0_
                    nw30_[int(10)] = (self).f15
                    nw30_[int(11)] = (self).f15
                    nw30_[int(12)] = True
                    nw30_[int(13)] = not(True)
                    nw30_[int(14)] = (self).f14
                    nw30_[int(15)] = (self).f15
                    nw30_[int(16)] = (self).f15
                    nw30_[int(17)] = (self).f15
                    nw30_[int(18)] = d_175_v0_
                    nw30_[int(19)] = (self).f15
                    d_226_v39_ = nw30_
                    d_227_v40_: _dafny.MultiSet
                    d_227_v40_ = _dafny.MultiSet([d_226_v39_])
                    d_175_v0_ = ((d_227_v40_).intersection(d_227_v40_)).isdisjoint((d_227_v40_) - (d_227_v40_))
                    (globalState).f1 = (d_178_v3_) + (d_178_v3_)
                    rhs26_ = not(d_175_v0_)
                    rhs27_ = (self).f14
                    rhs28_ = (-259) != ((d_176_v1_).fm6((self).f14, d_178_v3_, globalState))
                    rhs29_ = 2
                    d_175_v0_ = rhs26_
                    d_175_v0_ = rhs27_
                    d_175_v0_ = rhs28_
                    r1 = rhs29_
                    pass
            pass
        hi3_ = d_178_v3_
        for d_228_i9_ in range((d_178_v3_) * (d_178_v3_), hi3_):
            d_229_v41_: _dafny.Map
            d_229_v41_ = _dafny.Map({d_178_v3_: d_228_i9_})
            d_229_v41_ = (d_229_v41_).set((0) - ((d_228_i9_) - (341)), d_228_i9_)
            d_230_v42_: _dafny.Seq
            d_230_v42_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vkio"))
            d_231_v43_: _dafny.Seq
            d_231_v43_ = _dafny.SeqWithoutIsStrInference([(self).f14])
            d_232_v44_: _dafny.Array
            def lambda7_(d_233_i10_):
                return (self).f15

            init3_ = lambda7_
            nw31_ = _dafny.Array(None, 23)
            for i0_3_ in range(nw31_.length(0)):
                nw31_[i0_3_] = init3_(i0_3_)
            d_232_v44_ = nw31_
            d_234_v45_: _dafny.Map
            d_234_v45_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([default__.fm1(globalState), d_179_v4_, d_179_v4_])): d_232_v44_})
            d_235_v48_: _dafny.Array
            nw32_ = _dafny.Array(None, 25)
            nw32_[int(0)] = (self).f14
            nw32_[int(1)] = (self).f14
            nw32_[int(2)] = (self).f14
            nw32_[int(3)] = (self).f15
            nw32_[int(4)] = ((self).f10).issubset((self).f10)
            nw32_[int(5)] = d_175_v0_
            nw32_[int(6)] = ((self).f14) or ((self).f14)
            nw32_[int(7)] = (self).f15
            nw32_[int(8)] = d_175_v0_
            nw32_[int(9)] = (d_178_v3_) >= (d_228_i9_)
            nw32_[int(10)] = d_175_v0_
            nw32_[int(11)] = (self).f15
            nw32_[int(12)] = (d_230_v42_) != (d_230_v42_)
            nw32_[int(13)] = not((self).f15)
            nw32_[int(14)] = (d_176_v1_).fm5(default__.fm14((self).f14, globalState), d_228_i9_, d_228_i9_, globalState)
            nw32_[int(15)] = not((self).f14)
            nw32_[int(16)] = (d_231_v43_) == ((d_231_v43_).set(default__.safeIndex(len(d_231_v43_), len(d_231_v43_)), (self).f15))
            nw32_[int(17)] = (self).f15
            nw32_[int(18)] = True
            nw32_[int(19)] = ((d_176_v1_).fm13((self).f15, globalState)) == (not((self).f15))
            nw32_[int(20)] = (self).f14
            nw32_[int(21)] = (self).f15
            nw32_[int(22)] = (d_228_i9_) > (-262)
            nw32_[int(23)] = (d_228_i9_) <= (len(d_234_v45_))
            def iife17_():
                coll13_ = _dafny.Set()
                compr_13_: int
                for compr_13_ in _dafny.IntegerRange(708, 526):
                    d_236_v46_: int = compr_13_
                    if ((708) <= (d_236_v46_)) and ((d_236_v46_) < (526)):
                        coll13_ = coll13_.union(_dafny.Set([default__.safeDivisionInt(d_236_v46_, d_178_v3_)]))
                return _dafny.Set(coll13_)
            def iife18_():
                coll14_ = _dafny.Set()
                compr_14_: int
                for compr_14_ in _dafny.IntegerRange(184, 772):
                    d_237_v47_: int = compr_14_
                    if ((184) <= (d_237_v47_)) and ((d_237_v47_) < (772)):
                        coll14_ = coll14_.union(_dafny.Set([(d_237_v47_) * (d_178_v3_)]))
                return _dafny.Set(coll14_)
            nw32_[int(24)] = (iife17_()
            ).isdisjoint(iife18_()
            )
            d_235_v48_ = nw32_
            index32_ = default__.safeIndex(941, (d_235_v48_).length(0))
            (d_235_v48_)[index32_] = d_175_v0_
            d_238_v49_: _dafny.Map
            d_238_v49_ = _dafny.Map({d_175_v0_: d_230_v42_})
            d_239_v50_: _dafny.MultiSet
            d_239_v50_ = _dafny.MultiSet([len(d_238_v49_)])
            d_240_v51_: _dafny.Map
            d_240_v51_ = _dafny.Map({(self).f14: (d_239_v50_).cardinality})
            index33_ = default__.safeIndex(941, (d_235_v48_).length(0))
            (d_235_v48_)[index33_] = ((_dafny.Map({False: 842})) | (d_240_v51_)) != (_dafny.Map({(self).f15: d_228_i9_}))
            if (self).f14:
                d_230_v42_ = d_230_v42_
                def iife19_(_pat_let2_0):
                    def iife20_(d_241_dt__update__tmp_h0_):
                        def iife21_(_pat_let3_0):
                            def iife22_(d_243_dt__update_hcf16_h0_):
                                return D3_DC7(d_243_dt__update_hcf16_h0_)
                            return iife22_(_pat_let3_0)
                        return iife21_(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_242_i11_ in range(default__.abs(228))]))
                    return iife20_(_pat_let2_0)
                d_230_v42_ = (iife19_(D3_DC7(d_230_v42_))).cf16
                index34_ = default__.safeIndex(941, (d_235_v48_).length(0))
                (d_235_v48_)[index34_] = (self).f15
                d_244_v52_: _dafny.Array
                nw33_ = _dafny.Array(int(0), 4)
                d_244_v52_ = nw33_
                index35_ = default__.safeIndex(154, (d_244_v52_).length(0))
                (d_244_v52_)[index35_] = d_228_i9_
                index36_ = default__.safeIndex(154, (d_244_v52_).length(0))
                rhs30_ = d_178_v3_
                rhs31_ = d_178_v3_
                lhs20_ = d_244_v52_
                lhs21_ = default__.safeIndex(154, (d_244_v52_).length(0))
                r1 = rhs30_
                lhs20_[lhs21_] = rhs31_
                d_245_v53_: _dafny.Set
                d_245_v53_ = _dafny.Set({d_175_v0_})
                index37_ = default__.safeIndex(154, (d_244_v52_).length(0))
                (d_244_v52_)[index37_] = (0) - ((d_180_v5_)[default__.safeIndex(len(d_245_v53_), len(d_180_v5_))])
            elif True:
                d_246_v54_: _dafny.MultiSet
                d_246_v54_ = _dafny.MultiSet([(d_230_v42_ if (self).f14 else d_230_v42_)])
                d_246_v54_ = (d_246_v54_).intersection(((d_246_v54_).set(d_230_v42_, default__.abs(len(d_231_v43_)))) | (d_246_v54_))
                d_247_v55_: _dafny.Array
                nw34_ = _dafny.Array(int(0), 15)
                d_247_v55_ = nw34_
                index38_ = default__.safeIndex(300, (d_247_v55_).length(0))
                (d_247_v55_)[index38_] = d_228_i9_
                d_248_v56_: _dafny.Map
                d_248_v56_ = _dafny.Map({d_178_v3_: d_175_v0_})
                index39_ = default__.safeIndex(300, (d_247_v55_).length(0))
                rhs32_ = d_178_v3_
                rhs33_ = (d_176_v1_).fm6(((d_248_v56_)[d_178_v3_] if (d_178_v3_) in (d_248_v56_) else (self).f15), d_178_v3_, globalState)
                lhs22_ = globalState
                lhs23_ = d_247_v55_
                lhs24_ = default__.safeIndex(300, (d_247_v55_).length(0))
                lhs22_.f1 = rhs32_
                lhs23_[lhs24_] = rhs33_
                d_178_v3_ = len(d_230_v42_)
                d_249_v57_: _dafny.MultiSet
                d_249_v57_ = _dafny.MultiSet([d_179_v4_, d_179_v4_, d_179_v4_])
                d_250_v58_: D0
                d_250_v58_ = D0_DC1((self).f15, (d_249_v57_).cardinality)
                index40_ = default__.safeIndex(941, (d_235_v48_).length(0))
                (d_235_v48_)[index40_] = (d_250_v58_).cf1
                d_251_v59_: C0
                nw35_ = C0()
                nw35_.ctor__()
                d_251_v59_ = nw35_
            d_231_v43_ = (d_231_v43_) + (d_231_v43_)
        d_252_v60_: _dafny.Array
        def lambda8_(d_253_v0_):
            def lambda9_(d_254_i12_):
                return (_dafny.Set({True, d_253_v0_})).ispropersubset(_dafny.Set({d_253_v0_}))

            return lambda9_

        init4_ = lambda8_(d_175_v0_)
        nw36_ = _dafny.Array(None, 2)
        for i0_4_ in range(nw36_.length(0)):
            nw36_[i0_4_] = init4_(i0_4_)
        d_252_v60_ = nw36_
        r0 = d_252_v60_
        r1 = 34
        d_255_v61_: _dafny.Map
        d_255_v61_ = _dafny.Map({(self).f15: (self).f15})
        r2 = ((d_255_v61_) | (d_255_v61_)) | (d_255_v61_)
        return r0, r1, r2

    def m6(self, p0, p1, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: _dafny.Map = _dafny.Map({})
        r2: C0 = None
        d_256_v1_: _dafny.Seq
        def iife23_():
            coll15_ = _dafny.Set()
            compr_15_: int
            for compr_15_ in _dafny.IntegerRange(-679, 733):
                d_257_v0_: int = compr_15_
                if ((-679) <= (d_257_v0_)) and ((d_257_v0_) < (733)):
                    coll15_ = coll15_.union(_dafny.Set([(d_257_v0_) * (p1)]))
            return _dafny.Set(coll15_)
        d_256_v1_ = _dafny.SeqWithoutIsStrInference([default__.fm4(iife23_()
        , globalState)])
        if (d_256_v1_)[default__.safeIndex(p1, len(d_256_v1_))]:
            d_258_v2_: _dafny.Array
            def lambda10_(d_259_p1_):
                def lambda11_(d_260_i0_):
                    return (_dafny.Map({d_259_p1_: False})).set(d_259_p1_, (self).f14)

                return lambda11_

            init5_ = lambda10_(p1)
            nw37_ = _dafny.Array(None, 25)
            for i0_5_ in range(nw37_.length(0)):
                nw37_[i0_5_] = init5_(i0_5_)
            d_258_v2_ = nw37_
            d_258_v2_ = (d_258_v2_ if not((self).f15) else d_258_v2_)
            d_261_v3_: _dafny.Array
            nw38_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 11)
            d_261_v3_ = nw38_
            d_262_v4_: _dafny.Seq
            d_262_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xcoyciqnu"))
            index41_ = default__.safeIndex(231, (d_261_v3_).length(0))
            (d_261_v3_)[index41_] = d_262_v4_
            index42_ = default__.safeIndex(231, (d_261_v3_).length(0))
            (d_261_v3_)[index42_] = d_262_v4_
            d_263_v5_: _dafny.Map
            d_263_v5_ = _dafny.Map({p1: p1})
            d_263_v5_ = (d_263_v5_) | (d_263_v5_)
            d_264_v6_: _dafny.Seq
            d_264_v6_ = _dafny.SeqWithoutIsStrInference([(d_261_v3_)[default__.safeIndex(231, (d_261_v3_).length(0))], (d_261_v3_)[default__.safeIndex(231, (d_261_v3_).length(0))]])
            if ((d_264_v6_) + (d_264_v6_)) <= (d_264_v6_):
                d_265_v7_: _dafny.MultiSet
                d_265_v7_ = _dafny.MultiSet([(self).f14, (self).f15, (self).f14, (self).f15, (self).f15])
                d_266_v8_: bool
                d_266_v8_ = False
                d_267_v9_: _dafny.Set
                d_267_v9_ = _dafny.Set({((d_261_v3_)[default__.safeIndex(231, (d_261_v3_).length(0))]) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "twrc")))})
                d_268_v10_: D5
                d_268_v10_ = D5_DC12(d_267_v9_)
                rhs34_ = _dafny.MultiSet(default__.fm15(_dafny.CodePoint('w'), 767, globalState))
                rhs35_ = not ((self).f14) or ((self).f14)
                rhs36_ = (d_268_v10_).cf26
                d_265_v7_ = rhs34_
                d_266_v8_ = rhs35_
                d_267_v9_ = rhs36_
                d_269_v11_: str
                d_269_v11_ = _dafny.CodePoint('d')
                d_270_v12_: _dafny.Array
                nw39_ = _dafny.Array(int(0), 12)
                d_270_v12_ = nw39_
                d_271_v13_: _dafny.Array
                nw40_ = _dafny.Array(None, 11)
                nw40_[int(0)] = (0) - (p1)
                nw40_[int(1)] = (p1) - (p1)
                nw40_[int(2)] = p1
                nw40_[int(3)] = p1
                nw40_[int(4)] = p1
                nw40_[int(5)] = len((default__.fm2(d_266_v8_, (0) - (p1), d_269_v11_, p1, globalState)) + (d_262_v4_))
                nw40_[int(6)] = p1
                nw40_[int(7)] = default__.safeDivisionInt(p1, p1)
                nw40_[int(8)] = p1
                nw40_[int(9)] = len(_dafny.Map({d_266_v8_: d_270_v12_}))
                nw40_[int(10)] = p1
                d_271_v13_ = nw40_
                index43_ = default__.safeIndex(909, (d_271_v13_).length(0))
                (d_271_v13_)[index43_] = p1
                d_272_v14_: D0
                d_272_v14_ = D0_DC1((D1_DC5(p1, (self).f15, p1, d_266_v8_, p1)).cf11, p1)
                d_273_v15_: _dafny.Array
                nw41_ = _dafny.Array(False, 22)
                d_273_v15_ = nw41_
                d_274_v16_: D4
                d_274_v16_ = D4_DC11(d_272_v14_, d_273_v15_, p1)
                index44_ = default__.safeIndex(909, (d_271_v13_).length(0))
                (d_271_v13_)[index44_] = (d_274_v16_).cf25
                d_275_v17_: C0
                nw42_ = C0()
                nw42_.ctor__()
                d_275_v17_ = nw42_
                d_269_v11_ = d_269_v11_
                d_276_v18_: _dafny.Seq
                d_276_v18_ = _dafny.SeqWithoutIsStrInference([(0) - ((d_271_v13_)[default__.safeIndex(909, (d_271_v13_).length(0))])])
                (globalState).f1 = (((self).f10).cardinality) - (default__.safeModuloInt((d_271_v13_)[default__.safeIndex(909, (d_271_v13_).length(0))], (d_276_v18_)[default__.safeIndex(p1, len(d_276_v18_))]))
            elif True:
                d_277_v19_: D1
                d_277_v19_ = D1_DC5(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o"))), (self).f14, p1, (self).f15, len(_dafny.Map({p1: (self).f14})))
                d_278_v20_: _dafny.Map
                d_278_v20_ = _dafny.Map({d_277_v19_: (self).f15})
                pat_let_tv15_ = p1
                pat_let_tv16_ = p1
                def iife24_(_pat_let4_0):
                    def iife25_(d_279_dt__update__tmp_h0_):
                        def iife26_(_pat_let5_0):
                            def iife27_(d_280_dt__update_hcf10_h0_):
                                def iife28_(_pat_let6_0):
                                    def iife29_(d_281_dt__update_hcf14_h0_):
                                        def iife30_(_pat_let7_0):
                                            def iife31_(d_282_dt__update_hcf11_h0_):
                                                return D1_DC5(d_280_dt__update_hcf10_h0_, d_282_dt__update_hcf11_h0_, (d_279_dt__update__tmp_h0_).cf12, (d_279_dt__update__tmp_h0_).cf13, d_281_dt__update_hcf14_h0_)
                                            return iife31_(_pat_let7_0)
                                        return iife30_((self).f14)
                                    return iife29_(_pat_let6_0)
                                return iife28_(pat_let_tv16_)
                            return iife27_(_pat_let5_0)
                        return iife26_(pat_let_tv15_)
                    return iife25_(_pat_let4_0)
                d_278_v20_ = (d_278_v20_).set(iife24_(d_277_v19_), (self).f14)
                index45_ = default__.safeIndex(231, (d_261_v3_).length(0))
                (d_261_v3_)[index45_] = (d_261_v3_)[default__.safeIndex(231, (d_261_v3_).length(0))]
                d_283_v21_: _dafny.Seq
                d_283_v21_ = _dafny.SeqWithoutIsStrInference([d_256_v1_])
                d_284_v22_: str
                d_284_v22_ = _dafny.CodePoint('m')
                d_285_v23_: _dafny.Map
                d_285_v23_ = _dafny.Map({d_284_v22_: (self).f15})
                d_286_v24_: bool
                d_287_v25_: _dafny.MultiSet
                d_288_v26_: _dafny.Set
                d_289_v27_: int
                out16_: bool
                out17_: _dafny.MultiSet
                out18_: _dafny.Set
                out19_: int
                out16_, out17_, out18_, out19_ = default__.m0(((d_256_v1_).set(default__.safeIndex(p1, len(d_256_v1_)), False)) + (d_256_v1_), (d_283_v21_)[default__.safeIndex(len(d_285_v23_), len(d_283_v21_))], globalState)
                d_286_v24_ = out16_
                d_287_v25_ = out17_
                d_288_v26_ = out18_
                d_289_v27_ = out19_
                d_262_v4_ = (d_262_v4_).set(default__.safeIndex((0) - (default__.safeModuloInt(len(d_262_v4_), ((self).f10).cardinality)), len(d_262_v4_)), d_284_v22_)
                d_286_v24_ = d_286_v24_
            d_290_v28_: C0
            nw43_ = C0()
            nw43_.ctor__()
            d_290_v28_ = nw43_
        elif True:
            d_291_v29_: C0
            nw44_ = C0()
            nw44_.ctor__()
            d_291_v29_ = nw44_
            d_292_v30_: _dafny.Map
            d_292_v30_ = _dafny.Map({p1: (self).f15})
            d_293_v31_: _dafny.Set
            d_293_v31_ = _dafny.Set({p1})
            d_292_v30_ = (d_292_v30_).set(default__.safeModuloInt(p1, p1), ((self).f14 if not((self).f14) else default__.fm4(d_293_v31_, globalState)))
            d_294_v32_: bool
            d_294_v32_ = False
            d_294_v32_ = (self).f15
            d_295_v33_: str
            d_295_v33_ = _dafny.CodePoint('o')
            d_296_v34_: _dafny.Set
            d_296_v34_ = _dafny.Set({d_295_v33_, d_295_v33_, d_295_v33_, d_295_v33_, d_295_v33_})
            d_297_v35_: _dafny.Map
            d_297_v35_ = _dafny.Map({426: p1})
            d_298_v36_: _dafny.Array
            nw45_ = _dafny.Array(None, 20)
            nw45_[int(0)] = p1
            nw45_[int(1)] = p1
            nw45_[int(2)] = p1
            nw45_[int(3)] = (d_291_v29_).fm6(d_294_v32_, p1, globalState)
            nw45_[int(4)] = p1
            nw45_[int(5)] = p1
            nw45_[int(6)] = p1
            nw45_[int(7)] = p1
            nw45_[int(8)] = ((d_297_v35_)[p1] if (p1) in (d_297_v35_) else p1)
            nw45_[int(9)] = 746
            nw45_[int(10)] = p1
            nw45_[int(11)] = (0) - ((d_291_v29_).fm6((self).f15, p1, globalState))
            nw45_[int(12)] = p1
            nw45_[int(13)] = p1
            nw45_[int(14)] = (d_291_v29_).fm6((self).f14, (0) - (p1), globalState)
            nw45_[int(15)] = len(d_256_v1_)
            nw45_[int(16)] = p1
            nw45_[int(17)] = p1
            nw45_[int(18)] = p1
            nw45_[int(19)] = p1
            d_298_v36_ = nw45_
            d_299_v37_: D0
            d_299_v37_ = D0_DC1(not((self).f14), p1)
            d_300_v38_: _dafny.Map
            d_300_v38_ = _dafny.Map({(self).f15: d_299_v37_})
            d_301_v39_: _dafny.Map
            d_301_v39_ = _dafny.Map({d_298_v36_: (self).f15})
            d_302_v40_: _dafny.Seq
            d_302_v40_ = _dafny.SeqWithoutIsStrInference([d_300_v38_, d_300_v38_])
            pat_let_tv17_ = d_294_v32_
            d_303_v41_: _dafny.Array
            nw46_ = _dafny.Array(None, 15)
            nw46_[int(0)] = d_300_v38_
            nw46_[int(1)] = d_300_v38_
            nw46_[int(2)] = d_300_v38_
            nw46_[int(3)] = _dafny.Map({((d_301_v39_)[d_298_v36_] if (d_298_v36_) in (d_301_v39_) else (self).f15): d_299_v37_})
            def iife32_(_pat_let8_0):
                def iife33_(d_304_dt__update__tmp_h1_):
                    def iife34_(_pat_let9_0):
                        def iife35_(d_305_dt__update_hcf1_h0_):
                            return D0_DC1(d_305_dt__update_hcf1_h0_, (d_304_dt__update__tmp_h1_).cf2)
                        return iife35_(_pat_let9_0)
                    return iife34_(pat_let_tv17_)
                return iife33_(_pat_let8_0)
            nw46_[int(4)] = _dafny.Map({d_294_v32_: iife32_(d_299_v37_)})
            nw46_[int(5)] = (d_300_v38_).set(d_294_v32_, d_299_v37_)
            nw46_[int(6)] = d_300_v38_
            nw46_[int(7)] = d_300_v38_
            nw46_[int(8)] = ((_dafny.Map({d_294_v32_: d_299_v37_})).set(d_294_v32_, d_299_v37_)).set((self).f15, d_299_v37_)
            nw46_[int(9)] = _dafny.Map({d_294_v32_: d_299_v37_})
            nw46_[int(10)] = (d_302_v40_)[default__.safeIndex(len(d_297_v35_), len(d_302_v40_))]
            nw46_[int(11)] = d_300_v38_
            nw46_[int(12)] = d_300_v38_
            nw46_[int(13)] = _dafny.Map({d_294_v32_: D0_DC1((self).f14, p1)})
            nw46_[int(14)] = d_300_v38_
            d_303_v41_ = nw46_
            d_306_v42_: D5
            d_306_v42_ = D5_DC13((0) - ((p1) * (len(d_296_v34_))), d_298_v36_, d_303_v41_, p1)
            source7_ = d_306_v42_
            if source7_.is_DC13:
                d_307___mcc_h0_ = source7_.cf27
                d_308___mcc_h1_ = source7_.cf28
                d_309___mcc_h2_ = source7_.cf29
                d_310___mcc_h3_ = source7_.cf30
                d_311_cf30_ = d_310___mcc_h3_
                d_312_cf29_ = d_309___mcc_h2_
                d_313_cf28_ = d_308___mcc_h1_
                d_314_cf27_ = d_307___mcc_h0_
                d_315_v43_: _dafny.Array
                nw47_ = _dafny.Array(None, 17)
                nw47_[int(0)] = (self).f15
                nw47_[int(1)] = not(d_294_v32_)
                nw47_[int(2)] = False
                nw47_[int(3)] = d_294_v32_
                nw47_[int(4)] = d_294_v32_
                nw47_[int(5)] = d_294_v32_
                nw47_[int(6)] = (self).f14
                nw47_[int(7)] = ((self).f15) and (not(d_294_v32_))
                nw47_[int(8)] = (self).f15
                nw47_[int(9)] = (self).f15
                nw47_[int(10)] = (self).f14
                nw47_[int(11)] = d_294_v32_
                nw47_[int(12)] = (self).f14
                nw47_[int(13)] = d_294_v32_
                nw47_[int(14)] = (self).f14
                nw47_[int(15)] = (self).f15
                nw47_[int(16)] = d_294_v32_
                d_315_v43_ = nw47_
                index46_ = default__.safeIndex(55, (d_315_v43_).length(0))
                (d_315_v43_)[index46_] = (self).f14
                d_316_v44_: _dafny.Seq
                d_316_v44_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))
                index47_ = default__.safeIndex(55, (d_315_v43_).length(0))
                (d_315_v43_)[index47_] = (d_316_v44_) <= (d_316_v44_)
                d_317_v45_: _dafny.Array
                nw48_ = _dafny.Array(_dafny.CodePoint('D'), 19)
                d_317_v45_ = nw48_
                index48_ = default__.safeIndex(674, (d_317_v45_).length(0))
                (d_317_v45_)[index48_] = _dafny.CodePoint('c')
                index49_ = default__.safeIndex(674, (d_317_v45_).length(0))
                (d_317_v45_)[index49_] = d_295_v33_
                d_318_v46_: _dafny.Map
                d_318_v46_ = _dafny.Map({d_256_v1_: True})
                d_292_v30_ = (d_292_v30_).set(d_314_cf27_, ((d_318_v46_)[d_256_v1_] if (d_256_v1_) in (d_318_v46_) else d_294_v32_))
                d_319_v47_: _dafny.Array
                def lambda12_(d_320_i1_):
                    return _dafny.SeqWithoutIsStrInference([(self).f15])

                init6_ = lambda12_
                nw49_ = _dafny.Array(None, 2)
                for i0_6_ in range(nw49_.length(0)):
                    nw49_[i0_6_] = init6_(i0_6_)
                d_319_v47_ = nw49_
                index50_ = default__.safeIndex(170, (d_319_v47_).length(0))
                (d_319_v47_)[index50_] = (d_256_v1_) + (d_256_v1_)
                index51_ = default__.safeIndex(170, (d_319_v47_).length(0))
                (d_319_v47_)[index51_] = d_256_v1_
            elif source7_.is_DC14:
                d_321___mcc_h4_ = source7_.cf31
                d_322___mcc_h5_ = source7_.cf32
                d_323___mcc_h6_ = source7_.cf33
                d_324___mcc_h7_ = source7_.cf34
                d_325___mcc_h8_ = source7_.cf35
                d_326_cf35_ = d_325___mcc_h8_
                d_327_cf34_ = d_324___mcc_h7_
                d_328_cf33_ = d_323___mcc_h6_
                d_329_cf32_ = d_322___mcc_h5_
                d_330_cf31_ = d_321___mcc_h4_
                index52_ = default__.safeIndex(790, (d_298_v36_).length(0))
                (d_298_v36_)[index52_] = (d_329_cf32_ if (self).f14 else len(d_293_v31_))
                index53_ = default__.safeIndex(790, (d_298_v36_).length(0))
                (d_298_v36_)[index53_] = (p1) * (-398)
                def iife36_():
                    coll16_ = _dafny.Set()
                    compr_16_: int
                    for compr_16_ in (d_293_v31_).Elements:
                        d_331_v48_: int = compr_16_
                        if (d_331_v48_) in (d_293_v31_):
                            coll16_ = coll16_.union(_dafny.Set([(d_331_v48_) * (721)]))
                    return _dafny.Set(coll16_)
                d_293_v31_ = (d_293_v31_) | (iife36_()
                )
                d_294_v32_ = True
                d_332_v49_: _dafny.Array
                nw50_ = _dafny.Array(D0.default()(), 15)
                d_332_v49_ = nw50_
                index54_ = default__.safeIndex(464, (d_332_v49_).length(0))
                (d_332_v49_)[index54_] = d_299_v37_
                d_333_v50_: D0
                d_333_v50_ = D0_DC0(d_295_v33_)
                index55_ = default__.safeIndex(464, (d_332_v49_).length(0))
                (d_332_v49_)[index55_] = default__.fm16(d_333_v50_, globalState)
            elif source7_.is_DC12:
                d_334___mcc_h9_ = source7_.cf26
                d_335_cf26_ = d_334___mcc_h9_
                d_336_v51_: _dafny.Seq
                d_336_v51_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uhkhhuwaq"))
                d_336_v51_ = d_336_v51_
                (globalState).f1 = ((0) - (p1)) * (p1)
                index56_ = default__.safeIndex(635, (d_298_v36_).length(0))
                (d_298_v36_)[index56_] = len(d_293_v31_)
                index57_ = default__.safeIndex(635, (d_298_v36_).length(0))
                rhs37_ = (self).f14
                rhs38_ = ((_dafny.MultiSet([False])) - (((self).f10) - (_dafny.MultiSet([default__.fm4(d_293_v31_, globalState)])))).cardinality
                lhs25_ = d_298_v36_
                lhs26_ = default__.safeIndex(635, (d_298_v36_).length(0))
                d_294_v32_ = rhs37_
                lhs25_[lhs26_] = rhs38_
                index58_ = default__.safeIndex(635, (d_298_v36_).length(0))
                (d_298_v36_)[index58_] = p1
            elif True:
                d_337___mcc_h10_ = source7_.cf36
                d_338_cf36_ = d_337___mcc_h10_
                d_339_v52_: _dafny.Array
                def lambda13_(d_340_i2_):
                    return (self).f14

                init7_ = lambda13_
                nw51_ = _dafny.Array(None, 22)
                for i0_7_ in range(nw51_.length(0)):
                    nw51_[i0_7_] = init7_(i0_7_)
                d_339_v52_ = nw51_
                (globalState).f5 = ((d_339_v52_ if (self).f15 else d_339_v52_) if d_294_v32_ else d_339_v52_)
                (globalState).f1 = p1
                d_341_v53_: _dafny.Seq
                d_341_v53_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jfclgutmi"))
                d_341_v53_ = d_341_v53_
                d_294_v32_ = ((p1) != (p1)) and ((d_256_v1_)[default__.safeIndex(p1, len(d_256_v1_))])
            d_342_v54_: _dafny.Seq
            d_342_v54_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iexqde"))
            d_343_v55_: _dafny.MultiSet
            d_343_v55_ = _dafny.MultiSet([p1, p1, len(d_342_v54_)])
            d_344_v56_: _dafny.Set
            d_344_v56_ = _dafny.Set({(self).f15})
            d_345_v57_: _dafny.Seq
            d_345_v57_ = _dafny.SeqWithoutIsStrInference([default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([p1, p1, ((d_297_v35_)[p1] if (p1) in (d_297_v35_) else len(d_344_v56_)), p1])), p1), len(d_344_v56_), p1, len(_dafny.Map({default__.fm1(globalState): d_342_v54_})), p1])
            d_343_v55_ = _dafny.MultiSet(d_345_v57_)
        (globalState).f1 = p1
        d_346_v58_: _dafny.Map
        d_346_v58_ = _dafny.Map({(self).f15: d_256_v1_})
        d_347_v59_: _dafny.Set
        d_347_v59_ = _dafny.Set({p1, p1})
        d_348_v60_: str
        d_348_v60_ = _dafny.CodePoint('i')
        d_349_v61_: _dafny.Seq
        d_349_v61_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))
        d_350_v62_: _dafny.MultiSet
        d_350_v62_ = _dafny.MultiSet([p1, p1])
        rhs39_ = default__.safeModuloInt(p1, p1)
        rhs40_ = ((default__.fm15(d_348_v60_, p1, globalState)) + (d_256_v1_) if default__.fm4(d_347_v59_, globalState) else d_256_v1_)
        rhs41_ = (d_346_v58_).set((self).f14, (((d_256_v1_).set(default__.safeIndex(len(d_349_v61_), len(d_256_v1_)), (self).f14)).set(default__.safeIndex((d_350_v62_).cardinality, len((d_256_v1_).set(default__.safeIndex(len(d_349_v61_), len(d_256_v1_)), (self).f14))), (self).f15)) + (d_256_v1_))
        lhs27_ = globalState
        lhs27_.f1 = rhs39_
        d_256_v1_ = rhs40_
        d_346_v58_ = rhs41_
        d_351_v63_: _dafny.Set
        d_351_v63_ = _dafny.Set({(self).f15})
        hi4_ = 427
        for d_352_i3_ in range(len((d_351_v63_) - (d_351_v63_)), hi4_):
            d_353_v64_: C0
            nw52_ = C0()
            nw52_.ctor__()
            d_353_v64_ = nw52_
            d_354_v65_: D4
            d_354_v65_ = D4_DC10(d_353_v64_)
            d_355_v67_: D6
            def iife37_():
                coll17_ = _dafny.Set()
                compr_17_: int
                for compr_17_ in _dafny.IntegerRange(490, 676):
                    d_356_v66_: int = compr_17_
                    if ((490) <= (d_356_v66_)) and ((d_356_v66_) < (676)):
                        coll17_ = coll17_.union(_dafny.Set([(d_356_v66_) - (len(d_349_v61_))]))
                return _dafny.Set(coll17_)
            d_355_v67_ = D6_DC16(iife37_()
)
            d_354_v65_ = (d_354_v65_ if default__.fm4((d_355_v67_).cf37, globalState) else d_354_v65_)
            d_357_v68_: C0
            nw53_ = C0()
            nw53_.ctor__()
            d_357_v68_ = nw53_
            if (self).f14:
                d_256_v1_ = d_256_v1_
                d_358_v69_: _dafny.Array
                nw54_ = _dafny.Array(_dafny.Map({}), 15)
                d_358_v69_ = nw54_
                d_359_v70_: _dafny.Map
                d_359_v70_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ey"))])): (self).f10})
                index59_ = default__.safeIndex(90, (d_358_v69_).length(0))
                (d_358_v69_)[index59_] = d_359_v70_
                d_360_v71_: _dafny.Seq
                d_360_v71_ = _dafny.SeqWithoutIsStrInference([p1, d_352_i3_])
                d_361_v72_: D1
                d_361_v72_ = D1_DC4((self).f15, False, d_352_i3_, d_360_v71_, p1)
                d_362_v74_: _dafny.Array
                nw55_ = _dafny.Array(None, 14)
                def iife38_():
                    coll18_ = _dafny.Set()
                    compr_18_: int
                    for compr_18_ in _dafny.IntegerRange(771, -659):
                        d_363_v73_: int = compr_18_
                        if ((771) <= (d_363_v73_)) and ((d_363_v73_) < (-659)):
                            coll18_ = coll18_.union(_dafny.Set([(d_363_v73_) * (d_352_i3_)]))
                    return _dafny.Set(coll18_)
                nw55_[int(0)] = default__.fm4(iife38_()
                , globalState)
                nw55_[int(1)] = not((self).f14)
                nw55_[int(2)] = not((self).f14)
                nw55_[int(3)] = (self).f15
                nw55_[int(4)] = (self).f14
                nw55_[int(5)] = (self).f15
                nw55_[int(6)] = (self).f14
                nw55_[int(7)] = (self).f15
                nw55_[int(8)] = (self).f15
                nw55_[int(9)] = (self).f14
                nw55_[int(10)] = (self).f14
                nw55_[int(11)] = (self).f15
                nw55_[int(12)] = (self).f14
                nw55_[int(13)] = (self).f14
                d_362_v74_ = nw55_
                d_364_v75_: _dafny.Map
                d_364_v75_ = _dafny.Map({not((self).f14): d_362_v74_})
                d_365_v76_: _dafny.Seq
                d_365_v76_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({(d_361_v72_).cf6: d_362_v74_}), d_364_v75_, _dafny.Map({not((self).f15): d_362_v74_}), d_364_v75_])
                index60_ = default__.safeIndex(90, (d_358_v69_).length(0))
                (d_358_v69_)[index60_] = default__.fm17(_dafny.MultiSet([d_348_v60_]), len((d_365_v76_)[default__.safeIndex(d_352_i3_, len(d_365_v76_))]), globalState)
                d_366_v77_: _dafny.Map
                d_366_v77_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jtyk"))) == (d_349_v61_): (self).f14})
                d_366_v77_ = (d_366_v77_).set((self).f15, (p1) != (d_352_i3_))
                d_367_v78_: bool
                d_367_v78_ = False
                d_367_v78_ = not(((d_366_v77_)[(self).f15] if ((self).f15) in (d_366_v77_) else (d_347_v59_).issubset(d_347_v59_)))
                rhs42_ = p1
                rhs43_ = d_256_v1_
                lhs28_ = globalState
                lhs28_.f1 = rhs42_
                d_256_v1_ = rhs43_
            elif True:
                (globalState).f1 = (((self).f10) | ((_dafny.MultiSet(d_256_v1_)) | ((self).f10))).cardinality
                d_368_v79_: C0
                nw56_ = C0()
                nw56_.ctor__()
                d_368_v79_ = nw56_
                d_369_v80_: C0
                nw57_ = C0()
                nw57_.ctor__()
                d_369_v80_ = nw57_
                (globalState).f1 = ((d_350_v62_)[d_352_i3_] if (d_352_i3_) in (d_350_v62_) else d_352_i3_)
                d_370_v82_: _dafny.Array
                def lambda14_(d_371_i3_):
                    def lambda15_(d_372_i4_):
                        def iife39_():
                            coll19_ = _dafny.Map()
                            compr_19_: _dafny.Map
                            for compr_19_ in (_dafny.Map({_dafny.Map({(self).f14: (self).f15}): d_371_i3_})).keys.Elements:
                                d_373_v81_: _dafny.Map = compr_19_
                                if (d_373_v81_) in (_dafny.Map({_dafny.Map({(self).f14: (self).f15}): d_371_i3_})):
                                    coll19_[d_373_v81_] = (self).f14
                            return _dafny.Map(coll19_)
                        return (d_372_i4_) * (len(iife39_()
                        ))

                    return lambda15_

                init8_ = lambda14_(d_352_i3_)
                nw58_ = _dafny.Array(None, 10)
                for i0_8_ in range(nw58_.length(0)):
                    nw58_[i0_8_] = init8_(i0_8_)
                d_370_v82_ = nw58_
                d_374_v83_: _dafny.Map
                d_374_v83_ = _dafny.Map({d_352_i3_: ((self).f10).cardinality})
                d_375_v84_: _dafny.Map
                d_375_v84_ = _dafny.Map({p1: len(d_374_v83_)})
                index61_ = default__.safeIndex(672, (d_370_v82_).length(0))
                (d_370_v82_)[index61_] = len((d_374_v83_).set(p1, ((d_374_v83_)[d_352_i3_] if (d_352_i3_) in (d_374_v83_) else len(d_374_v83_))))
                index62_ = default__.safeIndex(672, (d_370_v82_).length(0))
                (d_370_v82_)[index62_] = ((d_353_v64_).fm6((self).f15, p1, globalState)) * (len(d_349_v61_))
            d_376_v85_: bool
            d_376_v85_ = True
            d_376_v85_ = not((True if d_376_v85_ else d_376_v85_))
        d_377_v86_: D0
        d_377_v86_ = D0_DC0(d_348_v60_)
        def lambda16_(source8_):
            if source8_.is_DC1:
                d_378___mcc_h11_ = source8_.cf1
                d_379___mcc_h12_ = source8_.cf2
                d_380_cf2_ = d_379___mcc_h12_
                d_381_cf1_ = d_378___mcc_h11_
                return False
            elif source8_.is_DC0:
                d_382___mcc_h13_ = source8_.cf0
                d_383_cf0_ = d_382___mcc_h13_
                return (self).f15
            elif True:
                d_384___mcc_h14_ = source8_.cf3
                d_385_cf3_ = d_384___mcc_h14_
                return True

        if lambda16_(d_377_v86_):
            if (((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wny"))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q"))).set(default__.safeIndex(p1, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q")))), d_348_v60_))).set(default__.safeIndex((0) - (p1), len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wny"))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q"))).set(default__.safeIndex(p1, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q")))), d_348_v60_)))), d_348_v60_)) == (d_349_v61_):
                d_386_v87_: C0
                nw59_ = C0()
                nw59_.ctor__()
                d_386_v87_ = nw59_
                d_387_v88_: _dafny.Map
                d_387_v88_ = _dafny.Map({p1: d_386_v87_})
                d_388_v89_: _dafny.Array
                nw60_ = _dafny.Array(None, 28)
                nw60_[int(0)] = d_386_v87_
                nw60_[int(1)] = (d_386_v87_ if (self).f15 else d_386_v87_)
                nw60_[int(2)] = d_386_v87_
                nw60_[int(3)] = d_386_v87_
                nw60_[int(4)] = d_386_v87_
                nw60_[int(5)] = ((d_387_v88_)[p1] if (p1) in (d_387_v88_) else d_386_v87_)
                nw60_[int(6)] = d_386_v87_
                nw60_[int(7)] = d_386_v87_
                nw60_[int(8)] = d_386_v87_
                nw60_[int(9)] = d_386_v87_
                nw60_[int(10)] = d_386_v87_
                nw60_[int(11)] = d_386_v87_
                nw60_[int(12)] = d_386_v87_
                nw60_[int(13)] = d_386_v87_
                nw60_[int(14)] = d_386_v87_
                nw60_[int(15)] = d_386_v87_
                nw60_[int(16)] = d_386_v87_
                nw60_[int(17)] = d_386_v87_
                nw60_[int(18)] = d_386_v87_
                nw60_[int(19)] = d_386_v87_
                nw60_[int(20)] = d_386_v87_
                nw60_[int(21)] = d_386_v87_
                nw60_[int(22)] = d_386_v87_
                nw60_[int(23)] = (d_386_v87_ if (self).f15 else d_386_v87_)
                nw60_[int(24)] = d_386_v87_
                nw60_[int(25)] = d_386_v87_
                nw60_[int(26)] = d_386_v87_
                nw60_[int(27)] = d_386_v87_
                d_388_v89_ = nw60_
                d_388_v89_ = d_388_v89_
                d_389_v90_: _dafny.Seq
                d_389_v90_ = _dafny.SeqWithoutIsStrInference([p1, p1])
                d_390_v91_: _dafny.Seq
                d_390_v91_ = _dafny.SeqWithoutIsStrInference([d_351_v63_, d_351_v63_, d_351_v63_, d_351_v63_, d_351_v63_])
                d_391_v94_: _dafny.Array
                nw61_ = _dafny.Array(None, 20)
                nw61_[int(0)] = (p1) * (p1)
                nw61_[int(1)] = (0) - (p1)
                nw61_[int(2)] = p1
                nw61_[int(3)] = p1
                nw61_[int(4)] = (d_389_v90_)[default__.safeIndex((0) - (p1), len(d_389_v90_))]
                nw61_[int(5)] = p1
                nw61_[int(6)] = p1
                nw61_[int(7)] = (0) - ((default__.fm16(d_377_v86_, globalState)).cf2)
                nw61_[int(8)] = len((d_390_v91_)[default__.safeIndex(p1, len(d_390_v91_))])
                nw61_[int(9)] = p1
                nw61_[int(10)] = (0) - (p1)
                def iife40_():
                    coll20_ = _dafny.Map()
                    compr_20_: int
                    for compr_20_ in _dafny.IntegerRange(861, -298):
                        d_392_v92_: int = compr_20_
                        if ((861) <= (d_392_v92_)) and ((d_392_v92_) < (-298)):
                            coll20_[(d_392_v92_) + (p1)] = (D0_DC1((self).f15, p1)).cf2
                    return _dafny.Map(coll20_)
                nw61_[int(11)] = len(iife40_()
                )
                nw61_[int(12)] = p1
                nw61_[int(13)] = p1
                nw61_[int(14)] = p1
                nw61_[int(15)] = (0) - (((self).f10).cardinality)
                nw61_[int(16)] = len(d_349_v61_)
                def iife41_():
                    coll21_ = _dafny.Map()
                    compr_21_: int
                    for compr_21_ in _dafny.IntegerRange(830, 26):
                        d_393_v93_: int = compr_21_
                        if ((830) <= (d_393_v93_)) and ((d_393_v93_) < (26)):
                            coll21_[(d_393_v93_) * (len(d_351_v63_))] = (self).f14
                    return _dafny.Map(coll21_)
                nw61_[int(17)] = default__.safeDivisionInt(p1, len(default__.fm18((self).f15, iife41_()
                , globalState)))
                nw61_[int(18)] = (p1) - (p1)
                nw61_[int(19)] = (d_386_v87_).fm6((self).f15, p1, globalState)
                d_391_v94_ = nw61_
                index63_ = default__.safeIndex(684, (d_391_v94_).length(0))
                (d_391_v94_)[index63_] = p1
                index64_ = default__.safeIndex(684, (d_391_v94_).length(0))
                rhs44_ = d_348_v60_
                rhs45_ = (len((d_389_v90_) + (d_389_v90_))) + (-840)
                rhs46_ = (0) - ((0) - (((0) - (-26) if (self).f15 else p1)))
                rhs47_ = (p1 if True else p1)
                rhs48_ = p1
                lhs29_ = d_391_v94_
                lhs30_ = default__.safeIndex(684, (d_391_v94_).length(0))
                lhs31_ = globalState
                lhs32_ = globalState
                lhs33_ = globalState
                d_348_v60_ = rhs44_
                lhs29_[lhs30_] = rhs45_
                lhs31_.f1 = rhs46_
                lhs32_.f1 = rhs47_
                lhs33_.f1 = rhs48_
                d_394_v95_: bool
                d_394_v95_ = False
                d_395_v96_: _dafny.Array
                nw62_ = _dafny.Array(False, 21)
                d_395_v96_ = nw62_
                d_396_v97_: _dafny.Array
                nw63_ = _dafny.Array(None, 13)
                nw63_[int(0)] = d_395_v96_
                nw63_[int(1)] = d_395_v96_
                nw63_[int(2)] = d_395_v96_
                nw63_[int(3)] = d_395_v96_
                nw63_[int(4)] = d_395_v96_
                nw63_[int(5)] = d_395_v96_
                nw63_[int(6)] = d_395_v96_
                nw63_[int(7)] = d_395_v96_
                nw63_[int(8)] = d_395_v96_
                nw63_[int(9)] = d_395_v96_
                nw63_[int(10)] = d_395_v96_
                nw63_[int(11)] = d_395_v96_
                nw63_[int(12)] = d_395_v96_
                d_396_v97_ = nw63_
                d_397_v98_: _dafny.Array
                def lambda17_(d_398_v60_):
                    def lambda18_(d_399_i5_):
                        return d_398_v60_

                    return lambda18_

                init9_ = lambda17_(d_348_v60_)
                nw64_ = _dafny.Array(None, 6)
                for i0_9_ in range(nw64_.length(0)):
                    nw64_[i0_9_] = init9_(i0_9_)
                d_397_v98_ = nw64_
                d_400_v99_: _dafny.Seq
                d_400_v99_ = _dafny.SeqWithoutIsStrInference([d_397_v98_, d_397_v98_, d_397_v98_])
                rhs49_ = (_dafny.Set({(self).f15, (self).f14})).ispropersubset(d_351_v63_)
                rhs50_ = p1
                rhs51_ = d_396_v97_
                rhs52_ = (d_386_v87_).fm5(d_351_v63_, len(d_400_v99_), (p1) + (p1), globalState)
                lhs34_ = globalState
                d_394_v95_ = rhs49_
                lhs34_.f1 = rhs50_
                d_396_v97_ = rhs51_
                d_394_v95_ = rhs52_
                d_401_v100_: _dafny.Array
                def lambda19_(d_402_v61_):
                    def lambda20_(d_403_i6_):
                        return d_402_v61_

                    return lambda20_

                init10_ = lambda19_(d_349_v61_)
                nw65_ = _dafny.Array(None, 11)
                for i0_10_ in range(nw65_.length(0)):
                    nw65_[i0_10_] = init10_(i0_10_)
                d_401_v100_ = nw65_
                index65_ = default__.safeIndex(333, (d_396_v97_).length(0))
                (d_396_v97_)[index65_] = d_395_v96_
                index66_ = default__.safeIndex(333, (d_396_v97_).length(0))
                rhs53_ = d_395_v96_
                rhs54_ = d_401_v100_
                rhs55_ = d_395_v96_
                lhs35_ = globalState
                lhs36_ = d_396_v97_
                lhs37_ = default__.safeIndex(333, (d_396_v97_).length(0))
                lhs35_.f5 = rhs53_
                d_401_v100_ = rhs54_
                lhs36_[lhs37_] = rhs55_
            elif True:
                d_404_v101_: _dafny.Array
                def lambda21_(d_405_i7_):
                    return (self).f15

                init11_ = lambda21_
                nw66_ = _dafny.Array(None, 20)
                for i0_11_ in range(nw66_.length(0)):
                    nw66_[i0_11_] = init11_(i0_11_)
                d_404_v101_ = nw66_
                d_406_v102_: _dafny.MultiSet
                d_406_v102_ = _dafny.MultiSet([d_404_v101_, d_404_v101_])
                (globalState).f1 = default__.safeDivisionInt(default__.safeModuloInt(p1, ((d_406_v102_).set(d_404_v101_, default__.abs(p1))).cardinality), p1)
                d_407_v103_: bool
                d_407_v103_ = False
                d_407_v103_ = (self).f14
                d_408_v104_: _dafny.Map
                d_408_v104_ = _dafny.Map({(self).f14: len(_dafny.Map({(self).f14: default__.fm19(d_348_v60_, (self).f14, globalState)}))})
                d_409_v105_: _dafny.Map
                d_409_v105_ = _dafny.Map({p1: p1})
                d_410_v106_: _dafny.Seq
                d_410_v106_ = _dafny.SeqWithoutIsStrInference([p1])
                d_411_v107_: _dafny.Array
                nw67_ = _dafny.Array(None, 29)
                nw67_[int(0)] = p1
                nw67_[int(1)] = p1
                nw67_[int(2)] = -308
                nw67_[int(3)] = p1
                nw67_[int(4)] = p1
                nw67_[int(5)] = p1
                nw67_[int(6)] = p1
                nw67_[int(7)] = len(d_349_v61_)
                nw67_[int(8)] = p1
                nw67_[int(9)] = p1
                nw67_[int(10)] = len(d_347_v59_)
                nw67_[int(11)] = p1
                nw67_[int(12)] = p1
                nw67_[int(13)] = len(_dafny.SeqWithoutIsStrInference([(self).f14]))
                nw67_[int(14)] = p1
                nw67_[int(15)] = ((self).f10).cardinality
                nw67_[int(16)] = p1
                nw67_[int(17)] = len(d_408_v104_)
                nw67_[int(18)] = p1
                nw67_[int(19)] = p1
                nw67_[int(20)] = len(_dafny.SeqWithoutIsStrInference([d_348_v60_ for d_412_i8_ in range(default__.abs(308))]))
                nw67_[int(21)] = p1
                nw67_[int(22)] = len(d_349_v61_)
                nw67_[int(23)] = len(d_349_v61_)
                nw67_[int(24)] = len(d_256_v1_)
                nw67_[int(25)] = p1
                nw67_[int(26)] = ((d_409_v105_)[p1] if (p1) in (d_409_v105_) else (0) - (default__.fm19(d_348_v60_, (self).f14, globalState)))
                nw67_[int(27)] = p1
                nw67_[int(28)] = len(d_410_v106_)
                d_411_v107_ = nw67_
                d_413_v108_: _dafny.MultiSet
                d_413_v108_ = _dafny.MultiSet([d_411_v107_])
                d_414_v109_: _dafny.Map
                d_414_v109_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([(0) - (p1)])): d_413_v108_})
                d_415_v110_: _dafny.Map
                d_415_v110_ = _dafny.Map({(self).f14: d_411_v107_})
                d_416_v111_: _dafny.Array
                nw68_ = _dafny.Array(None, 9)
                nw68_[int(0)] = d_413_v108_
                nw68_[int(1)] = d_413_v108_
                nw68_[int(2)] = _dafny.MultiSet([d_411_v107_])
                nw68_[int(3)] = d_413_v108_
                nw68_[int(4)] = ((d_414_v109_)[p1] if (p1) in (d_414_v109_) else d_413_v108_)
                nw68_[int(5)] = d_413_v108_
                nw68_[int(6)] = d_413_v108_
                nw68_[int(7)] = _dafny.MultiSet([((d_415_v110_)[d_407_v103_] if (d_407_v103_) in (d_415_v110_) else d_411_v107_)])
                nw68_[int(8)] = d_413_v108_
                d_416_v111_ = nw68_
                index67_ = default__.safeIndex(183, (d_416_v111_).length(0))
                (d_416_v111_)[index67_] = (d_413_v108_) | (d_413_v108_)
                index68_ = default__.safeIndex(183, (d_416_v111_).length(0))
                (d_416_v111_)[index68_] = d_413_v108_
                d_349_v61_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))
                (globalState).f1 = default__.safeDivisionInt(-725, len(d_408_v104_))
            d_417_v112_: bool
            d_417_v112_ = True
            d_418_v113_: D6
            d_418_v113_ = D6_DC17(p1, d_349_v61_, (self).f14)
            d_417_v112_ = not(((d_418_v113_ if (d_256_v1_)[default__.safeIndex(p1, len(d_256_v1_))] else d_418_v113_)).cf40)
            d_419_v114_: _dafny.Map
            d_419_v114_ = _dafny.Map({(d_256_v1_)[default__.safeIndex(p1, len(d_256_v1_))]: p1})
            d_419_v114_ = _dafny.Map({not(d_417_v112_): (p1) - (p1)})
            d_420_v116_: D5
            d_420_v116_ = D5_DC12(d_351_v63_)
            def iife42_():
                coll22_ = _dafny.Map()
                compr_22_: int
                for compr_22_ in ((_dafny.Map({p1: d_420_v116_})).set(p1, d_420_v116_)).keys.Elements:
                    d_421_v115_: int = compr_22_
                    if (d_421_v115_) in ((_dafny.Map({p1: d_420_v116_})).set(p1, d_420_v116_)):
                        coll22_[default__.safeModuloInt(d_421_v115_, p1)] = (_dafny.Map({len(d_256_v1_): (self).f14}) if d_417_v112_ else _dafny.Map({p1: (self).f14}))
                return _dafny.Map(coll22_)
            rhs56_ = default__.fm19(d_348_v60_, (self).f15, globalState)
            rhs57_ = p1
            rhs58_ = (0) - (len(iife42_()
            ))
            rhs59_ = ((d_349_v61_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_422_i9_ in range(default__.abs(156))])) if (self).f14 else d_349_v61_)
            lhs38_ = globalState
            lhs39_ = globalState
            lhs40_ = globalState
            lhs38_.f1 = rhs56_
            lhs39_.f1 = rhs57_
            lhs40_.f1 = rhs58_
            d_349_v61_ = rhs59_
            if (self).f15:
                d_417_v112_ = d_417_v112_
                d_417_v112_ = (self).f15
                d_423_v117_: bool
                d_424_v118_: _dafny.MultiSet
                d_425_v119_: _dafny.Set
                d_426_v120_: int
                out20_: bool
                out21_: _dafny.MultiSet
                out22_: _dafny.Set
                out23_: int
                out20_, out21_, out22_, out23_ = default__.m0(((d_256_v1_).set(default__.safeIndex(9, len(d_256_v1_)), (self).f15)) + (d_256_v1_), d_256_v1_, globalState)
                d_423_v117_ = out20_
                d_424_v118_ = out21_
                d_425_v119_ = out22_
                d_426_v120_ = out23_
                d_417_v112_ = d_423_v117_
                d_427_v121_: _dafny.Map
                d_427_v121_ = _dafny.Map({(_dafny.MultiSet([381])).cardinality: (self).f14})
                d_428_v122_: _dafny.MultiSet
                d_428_v122_ = _dafny.MultiSet([D0_DC1((self).f15, len(d_427_v121_))])
                d_429_v123_: D0
                d_429_v123_ = D0_DC1((self).f14, p1)
                (globalState).f1 = default__.safeDivisionInt(d_426_v120_, ((d_428_v122_)[d_429_v123_] if (d_429_v123_) in (d_428_v122_) else d_426_v120_))
            elif True:
                d_430_v124_: C0
                nw69_ = C0()
                nw69_.ctor__()
                d_430_v124_ = nw69_
                d_431_v125_: _dafny.Array
                nw70_ = _dafny.Array(_dafny.Array(None, 0), 21)
                d_431_v125_ = nw70_
                d_432_v126_: _dafny.Array
                nw71_ = _dafny.Array(int(0), 23)
                d_432_v126_ = nw71_
                index69_ = default__.safeIndex(274, (d_431_v125_).length(0))
                (d_431_v125_)[index69_] = d_432_v126_
                d_433_v128_: _dafny.Map
                d_433_v128_ = _dafny.Map({p1: p1})
                d_434_v129_: _dafny.Map
                def iife43_():
                    coll23_ = _dafny.Map()
                    compr_23_: int
                    for compr_23_ in (d_433_v128_).keys.Elements:
                        d_435_v127_: int = compr_23_
                        if (d_435_v127_) in (d_433_v128_):
                            coll23_[default__.safeModuloInt(d_435_v127_, p1)] = p1
                    return _dafny.Map(coll23_)
                d_434_v129_ = _dafny.Map({(iife43_()
                ) | (d_433_v128_): d_432_v126_})
                index70_ = default__.safeIndex(274, (d_431_v125_).length(0))
                (d_431_v125_)[index70_] = ((d_434_v129_)[d_433_v128_] if (d_433_v128_) in (d_434_v129_) else d_432_v126_)
                d_436_v130_: _dafny.Seq
                d_436_v130_ = _dafny.SeqWithoutIsStrInference([d_349_v61_])
                d_437_v131_: _dafny.Map
                d_437_v131_ = _dafny.Map({(self).f14: (d_436_v130_)[default__.safeIndex(p1, len(d_436_v130_))]})
                d_349_v61_ = ((d_437_v131_)[(self).f15] if ((self).f15) in (d_437_v131_) else d_349_v61_)
                d_417_v112_ = (self).f15
                d_417_v112_ = (self).f14
        elif True:
            d_438_v132_: _dafny.Array
            def lambda22_(d_439_v61_):
                def lambda23_(d_440_i10_):
                    return default__.safeDivisionInt(d_440_i10_, len(d_439_v61_))

                return lambda23_

            init12_ = lambda22_(d_349_v61_)
            nw72_ = _dafny.Array(None, 27)
            for i0_12_ in range(nw72_.length(0)):
                nw72_[i0_12_] = init12_(i0_12_)
            d_438_v132_ = nw72_
            d_438_v132_ = d_438_v132_
            d_441_v133_: bool
            d_441_v133_ = True
            d_442_v134_: _dafny.Map
            d_442_v134_ = _dafny.Map({((self).f10).cardinality: False})
            d_443_v135_: _dafny.Map
            d_443_v135_ = _dafny.Map({d_442_v134_: (0) - (len(_dafny.Map({p1: p1})))})
            d_444_v136_: _dafny.Array
            def lambda24_(d_445_v1_, d_446_p1_):
                def lambda25_(d_447_i11_):
                    return (d_445_v1_)[default__.safeIndex(d_446_p1_, len(d_445_v1_))]

                return lambda25_

            init13_ = lambda24_(d_256_v1_, p1)
            nw73_ = _dafny.Array(None, 21)
            for i0_13_ in range(nw73_.length(0)):
                nw73_[i0_13_] = init13_(i0_13_)
            d_444_v136_ = nw73_
            d_448_v137_: _dafny.MultiSet
            d_448_v137_ = _dafny.MultiSet([d_444_v136_])
            d_441_v133_ = (p1) >= (len((d_443_v135_).set(_dafny.Map({p1: d_441_v133_}), ((d_448_v137_)[d_444_v136_] if (d_444_v136_) in (d_448_v137_) else (0) - (p1)))))
            d_441_v133_ = not ((self).f14) or ((self).f15)
            d_441_v133_ = (self).f14
            if (self).f15:
                d_441_v133_ = (self).f15
                (globalState).f1 = ((549) - (p1)) - (default__.safeDivisionInt((0) - (p1), p1))
                d_449_v138_: _dafny.Map
                d_449_v138_ = _dafny.Map({(self).f15: (self).f15})
                d_450_v139_: _dafny.Map
                d_450_v139_ = _dafny.Map({len(d_449_v138_): (d_350_v62_) - (_dafny.MultiSet([p1]))})
                d_451_v140_: _dafny.MultiSet
                d_451_v140_ = (self).f10
                d_452_v141_: C0
                nw74_ = C0()
                nw74_.ctor__()
                d_452_v141_ = nw74_
                rhs60_ = (p1) * (len(d_442_v134_))
                rhs61_ = _dafny.Map({p1: d_350_v62_})
                rhs62_ = (((self).f10)) != ((self).f10)
                rhs63_ = d_452_v141_
                rhs64_ = ((((self).f10).cardinality) * (p1)) * ((0) - (p1))
                lhs41_ = globalState
                lhs42_ = globalState
                lhs41_.f1 = rhs60_
                d_450_v139_ = rhs61_
                d_441_v133_ = rhs62_
                r2 = rhs63_
                lhs42_.f1 = rhs64_
                d_453_v142_: _dafny.Array
                nw75_ = _dafny.Array(None, 26)
                d_453_v142_ = nw75_
                d_453_v142_ = d_453_v142_
                d_454_v143_: D1
                d_454_v143_ = D1_DC4((self).f14, False, p1, default__.fm20(p1, 713, globalState), -757)
                index71_ = default__.safeIndex(678, (d_444_v136_).length(0))
                (d_444_v136_)[index71_] = (d_454_v143_).cf6
                d_455_v144_: _dafny.Map
                d_455_v144_ = _dafny.Map({d_441_v133_: p1})
                index72_ = default__.safeIndex(678, (d_444_v136_).length(0))
                (d_444_v136_)[index72_] = not (((d_442_v134_)[len(d_455_v144_)] if (len(d_455_v144_)) in (d_442_v134_) else (self).f15)) or ((not(not(not(True))) if not(not((d_256_v1_)[default__.safeIndex(p1, len(d_256_v1_))])) else (self).f15))
            elif True:
                d_456_v145_: _dafny.Array
                nw76_ = _dafny.Array(None, 3)
                nw76_[int(0)] = _dafny.MultiSet([d_444_v136_, d_444_v136_, d_444_v136_])
                nw76_[int(1)] = _dafny.MultiSet([d_444_v136_, d_444_v136_])
                nw76_[int(2)] = d_448_v137_
                d_456_v145_ = nw76_
                index73_ = default__.safeIndex(508, (d_456_v145_).length(0))
                (d_456_v145_)[index73_] = (d_448_v137_) - (d_448_v137_)
                index74_ = default__.safeIndex(508, (d_456_v145_).length(0))
                (d_456_v145_)[index74_] = d_448_v137_
                d_457_v146_: D1
                d_457_v146_ = D1_DC4(d_441_v133_, (self).f14, p1, _dafny.SeqWithoutIsStrInference([-93 for d_458_i12_ in range(default__.abs(81))]), p1)
                index75_ = default__.safeIndex(994, (d_438_v132_).length(0))
                (d_438_v132_)[index75_] = (d_457_v146_).cf9
                index76_ = default__.safeIndex(994, (d_438_v132_).length(0))
                (d_438_v132_)[index76_] = default__.safeDivisionInt(p1, p1)
                d_441_v133_ = not((self).f14)
                d_459_v147_: C0
                nw77_ = C0()
                nw77_.ctor__()
                d_459_v147_ = nw77_
                d_460_v148_: C0
                nw78_ = C0()
                nw78_.ctor__()
                d_460_v148_ = nw78_
        if (self).f15:
            if (d_256_v1_)[default__.safeIndex(p1, len(d_256_v1_))]:
                d_461_v149_: bool
                d_461_v149_ = False
                d_461_v149_ = ((p1) * (p1)) >= (p1)
                d_462_v150_: C0
                nw79_ = C0()
                nw79_.ctor__()
                d_462_v150_ = nw79_
                (globalState).f1 = p1
                d_463_v151_: C0
                nw80_ = C0()
                nw80_.ctor__()
                d_463_v151_ = nw80_
                d_349_v61_ = (d_349_v61_) + (d_349_v61_)
            elif True:
                d_464_v152_: _dafny.Array
                def lambda26_(d_465_i13_):
                    return default__.safeModuloInt(d_465_i13_, 86)

                init14_ = lambda26_
                nw81_ = _dafny.Array(None, 25)
                for i0_14_ in range(nw81_.length(0)):
                    nw81_[i0_14_] = init14_(i0_14_)
                d_464_v152_ = nw81_
                index77_ = default__.safeIndex(484, (d_464_v152_).length(0))
                (d_464_v152_)[index77_] = default__.safeDivisionInt(default__.fm19(d_348_v60_, (self).f15, globalState), (0) - (p1))
                index78_ = default__.safeIndex(484, (d_464_v152_).length(0))
                (d_464_v152_)[index78_] = default__.fm19(d_348_v60_, (self).f15, globalState)
                d_466_v153_: _dafny.Array
                nw82_ = _dafny.Array(_dafny.Seq({}), 19)
                d_466_v153_ = nw82_
                index79_ = default__.safeIndex(209, (d_466_v153_).length(0))
                (d_466_v153_)[index79_] = _dafny.SeqWithoutIsStrInference([not((self).f14)])
                d_467_v154_: bool
                d_467_v154_ = True
                index80_ = default__.safeIndex(209, (d_466_v153_).length(0))
                rhs65_ = d_256_v1_
                rhs66_ = default__.fm4(d_347_v59_, globalState)
                rhs67_ = not((default__.fm19(d_348_v60_, (self).f14, globalState)) > (p1))
                lhs43_ = d_466_v153_
                lhs44_ = default__.safeIndex(209, (d_466_v153_).length(0))
                lhs43_[lhs44_] = rhs65_
                d_467_v154_ = rhs66_
                d_467_v154_ = rhs67_
                d_468_v155_: C0
                nw83_ = C0()
                nw83_.ctor__()
                d_468_v155_ = nw83_
                d_469_v156_: C0
                nw84_ = C0()
                nw84_.ctor__()
                d_469_v156_ = nw84_
                d_349_v61_ = default__.fm2(d_467_v154_, (d_464_v152_)[default__.safeIndex(484, (d_464_v152_).length(0))], d_348_v60_, (d_464_v152_)[default__.safeIndex(484, (d_464_v152_).length(0))], globalState)
            d_470_v157_: _dafny.Array
            nw85_ = _dafny.Array(False, 29)
            d_470_v157_ = nw85_
            index81_ = default__.safeIndex(210, (d_470_v157_).length(0))
            (d_470_v157_)[index81_] = (self).f15
            index82_ = default__.safeIndex(210, (d_470_v157_).length(0))
            (d_470_v157_)[index82_] = False
            d_471_v158_: _dafny.Array
            def lambda27_(d_472_v60_):
                def lambda28_(d_473_i14_):
                    return d_472_v60_

                return lambda28_

            init15_ = lambda27_(d_348_v60_)
            nw86_ = _dafny.Array(None, 26)
            for i0_15_ in range(nw86_.length(0)):
                nw86_[i0_15_] = init15_(i0_15_)
            d_471_v158_ = nw86_
            index83_ = default__.safeIndex(821, (d_471_v158_).length(0))
            (d_471_v158_)[index83_] = d_348_v60_
            index84_ = default__.safeIndex(821, (d_471_v158_).length(0))
            (d_471_v158_)[index84_] = default__.fm1(globalState)
            d_474_v159_: _dafny.Array
            nw87_ = _dafny.Array(int(0), 8)
            d_474_v159_ = nw87_
            d_475_v160_: _dafny.Map
            d_475_v160_ = _dafny.Map({d_474_v159_: p1})
            d_475_v160_ = ((d_475_v160_) | (d_475_v160_)) | (d_475_v160_)
            d_476_v161_: _dafny.Map
            d_476_v161_ = _dafny.Map({p1: (self).f15})
            d_477_v163_: D8
            d_477_v163_ = D8_DC19(d_476_v161_)
            d_478_v164_: _dafny.Set
            def iife44_():
                coll24_ = _dafny.Map()
                compr_24_: int
                for compr_24_ in _dafny.IntegerRange(976, 896):
                    d_479_v162_: int = compr_24_
                    if ((976) <= (d_479_v162_)) and ((d_479_v162_) < (896)):
                        coll24_[(d_479_v162_) - (p1)] = True
                return _dafny.Map(coll24_)
            d_478_v164_ = _dafny.Set({d_476_v161_, d_476_v161_, (iife44_()
            ) | ((d_477_v163_).cf42), d_476_v161_})
            d_480_v165_: _dafny.Seq
            d_480_v165_ = _dafny.SeqWithoutIsStrInference([d_478_v164_])
            d_478_v164_ = (d_478_v164_) | ((d_480_v165_)[default__.safeIndex(p1, len(d_480_v165_))])
        elif True:
            def iife45_():
                coll25_ = _dafny.Set()
                compr_25_: int
                for compr_25_ in _dafny.IntegerRange(52, -560):
                    d_481_v166_: int = compr_25_
                    if ((52) <= (d_481_v166_)) and ((d_481_v166_) < (-560)):
                        coll25_ = coll25_.union(_dafny.Set([default__.safeModuloInt(d_481_v166_, len(_dafny.SeqWithoutIsStrInference([p1, p1, len(_dafny.Map({(self).f14: len(_dafny.SeqWithoutIsStrInference([_dafny.Map({(self).f15: _dafny.MultiSet(d_256_v1_)}) for d_482_i15_ in range(default__.abs(385))]))}))])))]))
                return _dafny.Set(coll25_)
            if default__.fm4((iife45_()
            ) - (d_347_v59_), globalState):
                (globalState).f1 = p1
                d_483_v167_: _dafny.Array
                nw88_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 9)
                d_483_v167_ = nw88_
                index85_ = default__.safeIndex(416, (d_483_v167_).length(0))
                (d_483_v167_)[index85_] = ((d_349_v61_) + (d_349_v61_)).set(default__.safeIndex(p1, len((d_349_v61_) + (d_349_v61_))), d_348_v60_)
                index86_ = default__.safeIndex(416, (d_483_v167_).length(0))
                (d_483_v167_)[index86_] = (_dafny.SeqWithoutIsStrInference([d_348_v60_ for d_484_i16_ in range(default__.abs(285))])) + ((d_349_v61_).set(default__.safeIndex(p1, len(d_349_v61_)), d_348_v60_))
                d_485_v168_: bool
                d_485_v168_ = True
                d_485_v168_ = (self).f14
                (globalState).f1 = (0) - ((0) - (p1))
                d_486_v169_: _dafny.Map
                d_486_v169_ = _dafny.Map({((d_256_v1_) + (d_256_v1_)).set(default__.safeIndex(len(d_349_v61_), len((d_256_v1_) + (d_256_v1_))), (self).f14): (p1) * (p1)})
                d_487_v170_: D3
                d_487_v170_ = D3_DC7((d_483_v167_)[default__.safeIndex(416, (d_483_v167_).length(0))])
                d_488_v171_: _dafny.Map
                d_488_v171_ = _dafny.Map({d_487_v170_: False})
                d_489_v172_: _dafny.Map
                d_489_v172_ = _dafny.Map({(self).f15: not((self).f14)})
                d_490_v173_: _dafny.Map
                d_490_v173_ = _dafny.Map({(D0_DC1((self).f14, p1)).cf1: d_489_v172_})
                rhs68_ = ((d_486_v169_)[d_256_v1_] if (d_256_v1_) in (d_486_v169_) else p1)
                rhs69_ = (((d_488_v171_)[d_487_v170_] if (d_487_v170_) in (d_488_v171_) else ((d_489_v172_)[(self).f15] if ((self).f15) in (d_489_v172_) else (self).f14))) in ((((d_490_v173_)[d_485_v168_] if (d_485_v168_) in (d_490_v173_) else d_489_v172_)) | (d_489_v172_))
                rhs70_ = default__.safeDivisionInt(p1, default__.fm19(d_348_v60_, (self).f15, globalState))
                lhs45_ = globalState
                lhs46_ = globalState
                lhs45_.f1 = rhs68_
                d_485_v168_ = rhs69_
                lhs46_.f1 = rhs70_
            elif True:
                d_491_v174_: _dafny.MultiSet
                d_491_v174_ = (self).f10
                d_491_v174_ = _dafny.MultiSet([(d_256_v1_)[default__.safeIndex(p1, len(d_256_v1_))]])
                d_492_v175_: _dafny.Array
                nw89_ = _dafny.Array(None, 3)
                nw89_[int(0)] = d_377_v86_
                nw89_[int(1)] = default__.fm21(globalState)
                nw89_[int(2)] = d_377_v86_
                d_492_v175_ = nw89_
                d_493_v176_: _dafny.Map
                d_493_v176_ = _dafny.Map({d_492_v175_: (d_256_v1_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ale"))), len(d_256_v1_))]})
                d_494_v177_: _dafny.Seq
                d_494_v177_ = _dafny.SeqWithoutIsStrInference([d_347_v59_, d_347_v59_, d_347_v59_])
                d_495_v178_: _dafny.Map
                d_495_v178_ = _dafny.Map({len((d_494_v177_)[default__.safeIndex(p1, len(d_494_v177_))]): d_256_v1_})
                d_496_v179_: _dafny.Map
                d_496_v179_ = _dafny.Map({d_495_v178_: d_492_v175_})
                d_493_v176_ = (d_493_v176_).set(((d_496_v179_)[d_495_v178_] if (d_495_v178_) in (d_496_v179_) else d_492_v175_), (d_349_v61_) == (d_349_v61_))
                d_497_v180_: bool
                d_497_v180_ = False
                d_498_v181_: _dafny.Array
                nw90_ = _dafny.Array(int(0), 25)
                d_498_v181_ = nw90_
                index87_ = default__.safeIndex(95, (d_498_v181_).length(0))
                (d_498_v181_)[index87_] = p1
                index88_ = default__.safeIndex(95, (d_498_v181_).length(0))
                rhs71_ = False
                rhs72_ = p1
                lhs47_ = d_498_v181_
                lhs48_ = default__.safeIndex(95, (d_498_v181_).length(0))
                d_497_v180_ = rhs71_
                lhs47_[lhs48_] = rhs72_
                index89_ = default__.safeIndex(95, (d_498_v181_).length(0))
                (d_498_v181_)[index89_] = default__.safeDivisionInt(((d_498_v181_)[default__.safeIndex(95, (d_498_v181_).length(0))]) * ((d_498_v181_)[default__.safeIndex(95, (d_498_v181_).length(0))]), default__.fm19(default__.fm1(globalState), (self).f15, globalState))
                d_499_v182_: D1
                d_499_v182_ = D1_DC5(-460, (self).f15, (d_498_v181_)[default__.safeIndex(95, (d_498_v181_).length(0))], (self).f15, p1)
                d_500_v184_: _dafny.Seq
                d_500_v184_ = _dafny.SeqWithoutIsStrInference([(d_498_v181_)[default__.safeIndex(95, (d_498_v181_).length(0))]])
                d_501_v185_: _dafny.Seq
                d_501_v185_ = _dafny.SeqWithoutIsStrInference([d_499_v182_, d_499_v182_, d_499_v182_, D1_DC5(len(d_351_v63_), False, p1, True, len(d_500_v184_))])
                d_502_v188_: _dafny.Set
                d_502_v188_ = _dafny.Set({d_499_v182_})
                d_503_v189_: _dafny.Map
                d_503_v189_ = _dafny.Map({d_497_v180_: default__.safeModuloInt(p1, p1)})
                d_504_v190_: _dafny.Array
                nw91_ = _dafny.Array(False, 21)
                d_504_v190_ = nw91_
                d_505_v191_: D1
                d_505_v191_ = D1_DC3(d_504_v190_)
                d_506_v192_: _dafny.Set
                d_506_v192_ = _dafny.Set({d_505_v191_, d_505_v191_})
                d_507_v193_: _dafny.Map
                d_507_v193_ = _dafny.Map({d_506_v192_: p1})
                d_508_v194_: _dafny.Seq
                d_508_v194_ = _dafny.SeqWithoutIsStrInference([d_507_v193_, d_507_v193_, (_dafny.Map({d_506_v192_: len(d_256_v1_)})).set(_dafny.Set({d_505_v191_}), default__.fm19(d_348_v60_, (self).f15, globalState))])
                index90_ = default__.safeIndex(95, (d_498_v181_).length(0))
                index91_ = default__.safeIndex(95, (d_498_v181_).length(0))
                def iife46_():
                    def iife50_():
                        def iife52_():
                            coll32_ = _dafny.Set()
                            compr_32_: D1
                            for compr_32_ in (d_501_v185_).Elements:
                                d_514_v186_: D1 = compr_32_
                                if (d_514_v186_) in (d_501_v185_):
                                    coll32_ = coll32_.union(_dafny.Set([d_514_v186_]))
                            return _dafny.Set(coll32_)
                        coll30_ = _dafny.Map()
                        def iife51_():
                            coll31_ = _dafny.Set()
                            compr_31_: D1
                            for compr_31_ in (d_501_v185_).Elements:
                                d_513_v186_: D1 = compr_31_
                                if (d_513_v186_) in (d_501_v185_):
                                    coll31_ = coll31_.union(_dafny.Set([d_513_v186_]))
                            return _dafny.Set(coll31_)
                        compr_30_: D1
                        for compr_30_ in (iife51_()
                        ).Elements:
                            d_510_v183_: D1 = compr_30_
                            if (d_510_v183_) in (iife52_()
                            ):
                                coll30_[d_510_v183_] = (self).f15
                        return _dafny.Map(coll30_)
                    coll26_ = _dafny.Set()
                    def iife47_():
                        def iife49_():
                            coll29_ = _dafny.Set()
                            compr_29_: D1
                            for compr_29_ in (d_501_v185_).Elements:
                                d_511_v186_: D1 = compr_29_
                                if (d_511_v186_) in (d_501_v185_):
                                    coll29_ = coll29_.union(_dafny.Set([d_511_v186_]))
                            return _dafny.Set(coll29_)
                        coll27_ = _dafny.Map()
                        def iife48_():
                            coll28_ = _dafny.Set()
                            compr_28_: D1
                            for compr_28_ in (d_501_v185_).Elements:
                                d_509_v186_: D1 = compr_28_
                                if (d_509_v186_) in (d_501_v185_):
                                    coll28_ = coll28_.union(_dafny.Set([d_509_v186_]))
                            return _dafny.Set(coll28_)
                        compr_27_: D1
                        for compr_27_ in (iife48_()
                        ).Elements:
                            d_510_v183_: D1 = compr_27_
                            if (d_510_v183_) in (iife49_()
                            ):
                                coll27_[d_510_v183_] = (self).f15
                        return _dafny.Map(coll27_)
                    compr_26_: D1
                    for compr_26_ in (iife47_()
                    ).keys.Elements:
                        d_512_v187_: D1 = compr_26_
                        if (d_512_v187_) in (iife50_()
                        ):
                            coll26_ = coll26_.union(_dafny.Set([d_512_v187_]))
                    return _dafny.Set(coll26_)
                rhs73_ = ((_dafny.Set({d_499_v182_, D1_DC5(default__.fm19(d_348_v60_, (self).f15, globalState), False, (d_498_v181_)[default__.safeIndex(95, (d_498_v181_).length(0))], (self).f14, len(d_347_v59_))})) - (iife46_()
                )).isdisjoint(d_502_v188_)
                rhs74_ = ((d_503_v189_)[(self).f15] if ((self).f15) in (d_503_v189_) else p1)
                rhs75_ = not(True)
                rhs76_ = len((d_508_v194_)[default__.safeIndex((d_498_v181_)[default__.safeIndex(95, (d_498_v181_).length(0))], len(d_508_v194_))])
                lhs49_ = d_498_v181_
                lhs50_ = default__.safeIndex(95, (d_498_v181_).length(0))
                lhs51_ = d_498_v181_
                lhs52_ = default__.safeIndex(95, (d_498_v181_).length(0))
                d_497_v180_ = rhs73_
                lhs49_[lhs50_] = rhs74_
                d_497_v180_ = rhs75_
                lhs51_[lhs52_] = rhs76_
            d_515_v195_: _dafny.Set
            d_515_v195_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nthxycdvu"))})
            d_516_v196_: _dafny.Map
            d_516_v196_ = _dafny.Map({p1: d_515_v195_})
            d_517_v197_: _dafny.Array
            nw92_ = _dafny.Array(None, 9)
            nw92_[int(0)] = d_515_v195_
            nw92_[int(1)] = d_515_v195_
            nw92_[int(2)] = d_515_v195_
            nw92_[int(3)] = d_515_v195_
            nw92_[int(4)] = d_515_v195_
            nw92_[int(5)] = ((d_516_v196_)[p1] if (p1) in (d_516_v196_) else d_515_v195_)
            nw92_[int(6)] = _dafny.Set({d_349_v61_, d_349_v61_})
            nw92_[int(7)] = ((d_516_v196_)[p1] if (p1) in (d_516_v196_) else d_515_v195_)
            nw92_[int(8)] = (d_515_v195_) - (d_515_v195_)
            d_517_v197_ = nw92_
            index92_ = default__.safeIndex(973, (d_517_v197_).length(0))
            (d_517_v197_)[index92_] = (_dafny.Set({d_349_v61_})) - (d_515_v195_)
            index93_ = default__.safeIndex(973, (d_517_v197_).length(0))
            (d_517_v197_)[index93_] = default__.fm22(globalState)
            d_518_v198_: bool
            d_518_v198_ = False
            d_518_v198_ = not(((p1) == (p1)) or ((d_518_v198_) in (d_256_v1_)))
            d_519_v199_: T2
            nw93_ = C0()
            nw93_.ctor__()
            d_519_v199_ = nw93_
            d_519_v199_ = d_519_v199_
            d_520_v200_: _dafny.Array
            def lambda29_(d_521_v1_, d_522_v61_):
                def lambda30_(d_523_i17_):
                    return (d_521_v1_)[default__.safeIndex(len(d_522_v61_), len(d_521_v1_))]

                return lambda30_

            init16_ = lambda29_(d_256_v1_, d_349_v61_)
            nw94_ = _dafny.Array(None, 28)
            for i0_16_ in range(nw94_.length(0)):
                nw94_[i0_16_] = init16_(i0_16_)
            d_520_v200_ = nw94_
            index94_ = default__.safeIndex(561, (d_520_v200_).length(0))
            (d_520_v200_)[index94_] = d_518_v198_
            d_524_v201_: _dafny.Map
            d_524_v201_ = _dafny.Map({default__.safeDivisionInt(default__.fm19(d_348_v60_, True, globalState), 824): not((self).f14)})
            index95_ = default__.safeIndex(561, (d_520_v200_).length(0))
            (d_520_v200_)[index95_] = not(((d_524_v201_)[(0) - (default__.fm19(d_348_v60_, d_518_v198_, globalState))] if ((0) - (default__.fm19(d_348_v60_, d_518_v198_, globalState))) in (d_524_v201_) else d_518_v198_))
        d_525_v202_: _dafny.Array
        def lambda31_(d_526_p1_):
            def lambda32_(d_527_i18_):
                return ((0) - (d_526_p1_)) > (d_526_p1_)

            return lambda32_

        init17_ = lambda31_(p1)
        nw95_ = _dafny.Array(None, 3)
        for i0_17_ in range(nw95_.length(0)):
            nw95_[i0_17_] = init17_(i0_17_)
        d_525_v202_ = nw95_
        r0 = d_525_v202_
        d_528_v204_: _dafny.Map
        d_528_v204_ = _dafny.Map({p1: (self).f14})
        def iife53_():
            coll33_ = _dafny.Map()
            compr_33_: int
            for compr_33_ in _dafny.IntegerRange(171, 252):
                d_529_v203_: int = compr_33_
                if ((171) <= (d_529_v203_)) and ((d_529_v203_) < (252)):
                    coll33_[(d_529_v203_) + (429)] = (self).f14
            return _dafny.Map(coll33_)
        r1 = (iife53_()
        ) | (d_528_v204_)
        d_530_v205_: C0
        nw96_ = C0()
        nw96_.ctor__()
        d_530_v205_ = nw96_
        r2 = d_530_v205_
        return r0, r1, r2

    @property
    def f14(self):
        return self._f14
    @property
    def f15(self):
        return self._f15

class C2(T0):
    def  __init__(self):
        self._f10: _dafny.MultiSet = _dafny.MultiSet({})
        self.f12: bool = False
        self._f13: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    @property
    def f10(self):
        return self._f10
    def ctor__(self, f12, f13, f10):
        (self).f12 = f12
        (self)._f13 = f13
        (self)._f10 = f10

    def fm10(self, p0, p1, globalState):
        return _dafny.CodePoint('n')

    def fm11(self, p0, globalState):
        if not(self.f12):
            return (self).f13
        elif True:
            return (self).f13

    def m4(self, p0, p1, p2, globalState):
        r0: _dafny.Map = _dafny.Map({})
        d_531_v0_: _dafny.Seq
        d_531_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fsvmc"))
        (globalState).f1 = len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_532_i0_ in range(default__.abs(-614))])) + (d_531_v0_))
        d_533_v1_: _dafny.Array
        nw97_ = _dafny.Array(False, 14)
        d_533_v1_ = nw97_
        d_534_v2_: _dafny.Map
        d_534_v2_ = _dafny.Map({d_533_v1_: default__.fm12(globalState)})
        d_535_v3_: _dafny.Seq
        d_535_v3_ = _dafny.SeqWithoutIsStrInference([not(self.f12)])
        d_534_v2_ = (d_534_v2_).set(d_533_v1_, d_535_v3_)
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_533_v1_).length(0)):
            d_536_i1_: int = guard_loop_1_
            if (True) and (((0) <= (d_536_i1_)) and ((d_536_i1_) < ((d_533_v1_).length(0)))):
                (d_533_v1_)[(d_536_i1_)] = self.f12
        d_537_i2_: int
        d_537_i2_ = 0
        with _dafny.label("2"):
            while self.f12:
                with _dafny.c_label("2"):
                    if (d_537_i2_) >= (100):
                        raise _dafny.Break("2")
                    d_537_i2_ = (d_537_i2_) + (1)
                    d_538_v4_: _dafny.Set
                    d_538_v4_ = _dafny.Set({len(d_535_v3_), p1})
                    d_539_v5_: _dafny.Set
                    d_539_v5_ = _dafny.Set({len(d_538_v4_), (self).f13})
                    if (_dafny.Set({p0, p1})).isdisjoint(d_539_v5_):
                        index96_ = default__.safeIndex(766, (d_533_v1_).length(0))
                        (d_533_v1_)[index96_] = self.f12
                        d_540_v6_: _dafny.Set
                        d_540_v6_ = _dafny.Set({default__.fm4(d_539_v5_, globalState)})
                        index97_ = default__.safeIndex(766, (d_533_v1_).length(0))
                        (d_533_v1_)[index97_] = (_dafny.Set({default__.fm4(_dafny.Set({p1, p1, p1}), globalState), self.f12, self.f12, self.f12})).isdisjoint(d_540_v6_)
                        d_541_v7_: D0
                        d_541_v7_ = D0_DC1((d_533_v1_)[default__.safeIndex(766, (d_533_v1_).length(0))], p0)
                        d_542_v8_: D0
                        d_542_v8_ = D0_DC2(d_541_v7_)
                        d_543_v9_: _dafny.MultiSet
                        d_543_v9_ = _dafny.MultiSet([d_542_v8_])
                        d_544_v10_: _dafny.Map
                        d_544_v10_ = _dafny.Map({d_543_v9_: p2})
                        d_544_v10_ = (d_544_v10_).set(d_543_v9_, p2)
                        d_545_v11_: _dafny.MultiSet
                        d_545_v11_ = _dafny.MultiSet([p1, p1])
                        d_546_v12_: _dafny.Map
                        d_546_v12_ = _dafny.Map({(_dafny.Set({p1, p1, (self).fm11(_dafny.MultiSet(p2), globalState), (self).f13})).intersection(d_539_v5_): ((self).f10).set(self.f12, default__.abs((d_545_v11_).cardinality))})
                        d_546_v12_ = (d_546_v12_).set((_dafny.Set({p1, p0})) - (d_538_v4_), ((self).f10).set((d_533_v1_)[default__.safeIndex(766, (d_533_v1_).length(0))], default__.abs(p0)))
                        d_547_v13_: str
                        d_547_v13_ = _dafny.CodePoint('t')
                        d_548_v14_: _dafny.Seq
                        d_548_v14_ = _dafny.SeqWithoutIsStrInference([(d_545_v11_).set((self).f13, default__.abs((self).f13)), _dafny.MultiSet(p2), d_545_v11_])
                        d_549_v15_: _dafny.Array
                        nw98_ = _dafny.Array(None, 17)
                        nw98_[int(0)] = d_531_v0_
                        nw98_[int(1)] = d_531_v0_
                        nw98_[int(2)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_550_i3_ in range(default__.abs(450))])
                        nw98_[int(3)] = default__.fm2((d_533_v1_)[default__.safeIndex(766, (d_533_v1_).length(0))], p0, d_547_v13_, ((d_548_v14_)[default__.safeIndex(p1, len(d_548_v14_))]).cardinality, globalState)
                        nw98_[int(4)] = d_531_v0_
                        nw98_[int(5)] = (d_531_v0_).set(default__.safeIndex(len(d_531_v0_), len(d_531_v0_)), d_547_v13_)
                        nw98_[int(6)] = d_531_v0_
                        nw98_[int(7)] = d_531_v0_
                        nw98_[int(8)] = (d_531_v0_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qmb")))
                        nw98_[int(9)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jqnofwwqn"))) + (d_531_v0_)
                        nw98_[int(10)] = d_531_v0_
                        nw98_[int(11)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "majd"))
                        nw98_[int(12)] = d_531_v0_
                        nw98_[int(13)] = d_531_v0_
                        nw98_[int(14)] = d_531_v0_
                        nw98_[int(15)] = (d_531_v0_) + (_dafny.SeqWithoutIsStrInference([d_547_v13_ for d_551_i4_ in range(default__.abs(-611))]))
                        nw98_[int(16)] = _dafny.SeqWithoutIsStrInference([d_547_v13_ for d_552_i5_ in range(default__.abs(-701))])
                        d_549_v15_ = nw98_
                        index98_ = default__.safeIndex(153, (d_549_v15_).length(0))
                        (d_549_v15_)[index98_] = d_531_v0_
                        index99_ = default__.safeIndex(153, (d_549_v15_).length(0))
                        (d_549_v15_)[index99_] = d_531_v0_
                        d_553_v16_: _dafny.Array
                        nw99_ = _dafny.Array(int(0), 12)
                        d_553_v16_ = nw99_
                        index100_ = default__.safeIndex(972, (d_553_v16_).length(0))
                        (d_553_v16_)[index100_] = ((p2)[default__.safeIndex((self).f13, len(p2))] if (d_533_v1_)[default__.safeIndex(766, (d_533_v1_).length(0))] else p1)
                        index101_ = default__.safeIndex(766, (d_533_v1_).length(0))
                        index102_ = default__.safeIndex(972, (d_553_v16_).length(0))
                        index103_ = default__.safeIndex(153, (d_549_v15_).length(0))
                        rhs77_ = (self).f13
                        rhs78_ = self.f12
                        rhs79_ = (p0) * (p0)
                        rhs80_ = (d_549_v15_)[default__.safeIndex(153, (d_549_v15_).length(0))]
                        lhs53_ = globalState
                        lhs54_ = d_533_v1_
                        lhs55_ = default__.safeIndex(766, (d_533_v1_).length(0))
                        lhs56_ = d_553_v16_
                        lhs57_ = default__.safeIndex(972, (d_553_v16_).length(0))
                        lhs58_ = d_549_v15_
                        lhs59_ = default__.safeIndex(153, (d_549_v15_).length(0))
                        lhs53_.f1 = rhs77_
                        lhs54_[lhs55_] = rhs78_
                        lhs56_[lhs57_] = rhs79_
                        lhs58_[lhs59_] = rhs80_
                    elif True:
                        d_554_v17_: _dafny.Map
                        d_554_v17_ = _dafny.Map({386: p0})
                        d_554_v17_ = d_554_v17_
                        d_538_v4_ = d_538_v4_
                        (self).f12 = self.f12
                        d_555_v18_: D1
                        d_555_v18_ = D1_DC3(d_533_v1_)
                        d_556_v19_: _dafny.Array
                        nw100_ = _dafny.Array(None, 19)
                        nw100_[int(0)] = d_533_v1_
                        nw100_[int(1)] = d_533_v1_
                        nw100_[int(2)] = d_533_v1_
                        nw100_[int(3)] = d_533_v1_
                        nw100_[int(4)] = d_533_v1_
                        nw100_[int(5)] = d_533_v1_
                        nw100_[int(6)] = (d_555_v18_).cf4
                        nw100_[int(7)] = d_533_v1_
                        nw100_[int(8)] = d_533_v1_
                        nw100_[int(9)] = d_533_v1_
                        nw100_[int(10)] = d_533_v1_
                        nw100_[int(11)] = d_533_v1_
                        nw100_[int(12)] = d_533_v1_
                        nw100_[int(13)] = d_533_v1_
                        nw100_[int(14)] = d_533_v1_
                        nw100_[int(15)] = d_533_v1_
                        nw100_[int(16)] = d_533_v1_
                        nw100_[int(17)] = d_533_v1_
                        nw100_[int(18)] = d_533_v1_
                        d_556_v19_ = nw100_
                        d_557_v20_: _dafny.Map
                        d_557_v20_ = _dafny.Map({not(True): d_556_v19_})
                        d_557_v20_ = (d_557_v20_).set(self.f12, d_556_v19_)
                        d_558_v21_: _dafny.Map
                        d_558_v21_ = _dafny.Map({d_555_v18_: self.f12})
                        d_559_v22_: _dafny.MultiSet
                        d_559_v22_ = _dafny.MultiSet([(self).f13, (self).f13])
                        d_560_v23_: str
                        d_560_v23_ = _dafny.CodePoint('c')
                        d_561_v24_: D1
                        d_561_v24_ = D1_DC4(self.f12, self.f12, len(d_558_v21_), _dafny.SeqWithoutIsStrInference([len(d_531_v0_), len(default__.fm2(self.f12, (d_559_v22_).cardinality, d_560_v23_, p0, globalState))]), (self).fm11(_dafny.MultiSet([len(d_531_v0_)]), globalState))
                        d_562_v25_: _dafny.Seq
                        d_562_v25_ = _dafny.SeqWithoutIsStrInference([d_561_v24_, d_561_v24_, d_561_v24_, d_561_v24_, d_561_v24_])
                        d_563_v26_: _dafny.Map
                        d_563_v26_ = _dafny.Map({self.f12: d_562_v25_})
                        d_564_v27_: D0
                        d_564_v27_ = D0_DC1(not((len(d_563_v26_)) != ((self).f13)), 156)
                        index104_ = default__.safeIndex(218, (d_533_v1_).length(0))
                        (d_533_v1_)[index104_] = self.f12
                        index105_ = default__.safeIndex(963, (d_556_v19_).length(0))
                        (d_556_v19_)[index105_] = d_533_v1_
                        index106_ = default__.safeIndex(218, (d_533_v1_).length(0))
                        index107_ = default__.safeIndex(963, (d_556_v19_).length(0))
                        rhs81_ = d_564_v27_
                        rhs82_ = (not((d_559_v22_).issubset(d_559_v22_)) if self.f12 else self.f12)
                        rhs83_ = default__.safeDivisionInt(((self).f13) * ((self).fm11(d_559_v22_, globalState)), len((d_554_v17_ if self.f12 else d_554_v17_)))
                        rhs84_ = d_533_v1_
                        lhs60_ = d_533_v1_
                        lhs61_ = default__.safeIndex(218, (d_533_v1_).length(0))
                        lhs62_ = globalState
                        lhs63_ = d_556_v19_
                        lhs64_ = default__.safeIndex(963, (d_556_v19_).length(0))
                        d_564_v27_ = rhs81_
                        lhs60_[lhs61_] = rhs82_
                        lhs62_.f1 = rhs83_
                        lhs63_[lhs64_] = rhs84_
                    d_565_v28_: D1
                    d_565_v28_ = D1_DC3(d_533_v1_)
                    d_566_v29_: _dafny.Array
                    nw101_ = _dafny.Array(None, 3)
                    nw101_[int(0)] = d_565_v28_
                    nw101_[int(1)] = d_565_v28_
                    nw101_[int(2)] = d_565_v28_
                    d_566_v29_ = nw101_
                    index108_ = default__.safeIndex(245, (d_566_v29_).length(0))
                    (d_566_v29_)[index108_] = d_565_v28_
                    index109_ = default__.safeIndex(245, (d_566_v29_).length(0))
                    (d_566_v29_)[index109_] = d_565_v28_
                    if self.f12:
                        d_567_v30_: str
                        d_567_v30_ = _dafny.CodePoint('n')
                        d_568_v31_: _dafny.Set
                        d_568_v31_ = _dafny.Set({False, self.f12})
                        d_569_v32_: _dafny.Set
                        d_569_v32_ = _dafny.Set({default__.fm2(self.f12, len(d_568_v31_), d_567_v30_, p1, globalState), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h"))})
                        d_570_v33_: D0
                        d_570_v33_ = D0_DC1((p0) >= (-711), len(d_569_v32_))
                        rhs85_ = (self).f13
                        rhs86_ = _dafny.CodePoint('c')
                        rhs87_ = d_570_v33_
                        lhs65_ = globalState
                        lhs65_.f1 = rhs85_
                        d_567_v30_ = rhs86_
                        d_570_v33_ = rhs87_
                        rhs88_ = (self).f13
                        rhs89_ = p1
                        lhs66_ = globalState
                        lhs67_ = globalState
                        lhs66_.f1 = rhs88_
                        lhs67_.f1 = rhs89_
                        d_571_v34_: _dafny.Map
                        d_571_v34_ = _dafny.Map({-505: self.f12})
                        d_572_v35_: _dafny.Map
                        d_572_v35_ = _dafny.Map({self.f12: (self).f13})
                        d_571_v34_ = (d_571_v34_).set(len((d_572_v35_).set(self.f12, p0)), self.f12)
                        d_573_v36_: _dafny.Array
                        nw102_ = _dafny.Array(_dafny.MultiSet({}), 26)
                        d_573_v36_ = nw102_
                        def lambda33_(d_574_i6_):
                            return _dafny.MultiSet([(self).f13, (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([self.f12]))])).cardinality])

                        init18_ = lambda33_
                        nw103_ = _dafny.Array(None, 13)
                        for i0_18_ in range(nw103_.length(0)):
                            nw103_[i0_18_] = init18_(i0_18_)
                        d_573_v36_ = nw103_
                        (self).f12 = self.f12
                    elif True:
                        d_575_v37_: _dafny.Array
                        def lambda34_(d_576_v3_):
                            def lambda35_(d_577_i7_):
                                return d_576_v3_

                            return lambda35_

                        init19_ = lambda34_(d_535_v3_)
                        nw104_ = _dafny.Array(None, 27)
                        for i0_19_ in range(nw104_.length(0)):
                            nw104_[i0_19_] = init19_(i0_19_)
                        d_575_v37_ = nw104_
                        index110_ = default__.safeIndex(561, (d_575_v37_).length(0))
                        (d_575_v37_)[index110_] = d_535_v3_
                        index111_ = default__.safeIndex(561, (d_575_v37_).length(0))
                        (d_575_v37_)[index111_] = d_535_v3_
                        (globalState).f1 = p1
                        index112_ = default__.safeIndex(245, (d_566_v29_).length(0))
                        (d_566_v29_)[index112_] = (d_566_v29_)[default__.safeIndex(245, (d_566_v29_).length(0))]
                        (globalState).f1 = 66
                        d_578_v38_: _dafny.MultiSet
                        d_578_v38_ = _dafny.MultiSet([self.f12, not(self.f12), self.f12, self.f12])
                        d_579_v39_: _dafny.MultiSet
                        d_579_v39_ = _dafny.MultiSet([(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "micx")))), p0])
                        d_580_v40_: _dafny.Map
                        d_580_v40_ = _dafny.Map({(self).fm11(d_579_v39_, globalState): self.f12})
                        d_581_v41_: D0
                        d_581_v41_ = D0_DC1(self.f12, p0)
                        rhs90_ = ((d_578_v38_) | (_dafny.MultiSet(default__.fm12(globalState)))).intersection(d_578_v38_)
                        rhs91_ = (d_533_v1_ if (False if not(((d_580_v40_)[(d_581_v41_).cf2] if ((d_581_v41_).cf2) in (d_580_v40_) else self.f12)) else self.f12) else d_533_v1_)
                        lhs68_ = globalState
                        d_578_v38_ = rhs90_
                        lhs68_.f5 = rhs91_
                    if (d_531_v0_) < (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hmmeoafg"))):
                        d_582_v42_: str
                        d_582_v42_ = _dafny.CodePoint('g')
                        d_583_v43_: D0
                        d_583_v43_ = D0_DC0(d_582_v42_)
                        d_584_v44_: D0
                        d_584_v44_ = D0_DC1(not(self.f12), (0) - ((p2)[default__.safeIndex(p1, len(p2))]))
                        index113_ = default__.safeIndex(164, (d_533_v1_).length(0))
                        (d_533_v1_)[index113_] = True
                        d_585_v45_: _dafny.MultiSet
                        d_585_v45_ = _dafny.MultiSet([p0, 756, (self).f13, (self).f13, (p2)[default__.safeIndex(p1, len(p2))]])
                        d_586_v46_: _dafny.MultiSet
                        d_586_v46_ = d_585_v45_
                        index114_ = default__.safeIndex(164, (d_533_v1_).length(0))
                        def iife55_(_pat_let11_0):
                            def iife56_(d_587_dt__update__tmp_h0_):
                                def iife57_(_pat_let12_0):
                                    def iife58_(d_588_dt__update_hcf1_h0_):
                                        return D0_DC1(d_588_dt__update_hcf1_h0_, (d_587_dt__update__tmp_h0_).cf2)
                                    return iife58_(_pat_let12_0)
                                return iife57_(self.f12)
                            return iife56_(_pat_let11_0)
                        def iife54_(_pat_let10_0):
                            def iife59_(d_589_dt__update__tmp_h1_):
                                def iife60_(_pat_let13_0):
                                    def iife61_(d_590_dt__update_hcf1_h1_):
                                        return D0_DC1(d_590_dt__update_hcf1_h1_, (d_589_dt__update__tmp_h1_).cf2)
                                    return iife61_(_pat_let13_0)
                                return iife60_(self.f12)
                            return iife59_(_pat_let10_0)
                        rhs92_ = (self).fm11(((d_586_v46_)) | (d_585_v45_), globalState)
                        rhs93_ = d_583_v43_
                        rhs94_ = iife54_(iife55_(D0_DC1(self.f12, p1)))
                        rhs95_ = self.f12
                        lhs69_ = globalState
                        lhs70_ = d_533_v1_
                        lhs71_ = default__.safeIndex(164, (d_533_v1_).length(0))
                        lhs69_.f1 = rhs92_
                        d_583_v43_ = rhs93_
                        d_584_v44_ = rhs94_
                        lhs70_[lhs71_] = rhs95_
                        d_591_v47_: _dafny.Seq
                        d_591_v47_ = _dafny.SeqWithoutIsStrInference([d_531_v0_, d_531_v0_, d_531_v0_, d_531_v0_])
                        (globalState).f1 = default__.safeDivisionInt((0) - (len((d_591_v47_)[default__.safeIndex(p1, len(d_591_v47_))])), 910)
                        d_531_v0_ = (((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bvmvcibbo"))) + (d_531_v0_)) + (((d_531_v0_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hlqwmddvn"))), len(d_531_v0_)), d_582_v42_)) + (d_531_v0_))).set(default__.safeIndex((self).f13, len(((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bvmvcibbo"))) + (d_531_v0_)) + (((d_531_v0_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hlqwmddvn"))), len(d_531_v0_)), d_582_v42_)) + (d_531_v0_)))), d_582_v42_)
                        (globalState).f1 = p0
                        index115_ = default__.safeIndex(164, (d_533_v1_).length(0))
                        (d_533_v1_)[index115_] = self.f12
                    elif True:
                        (globalState).f1 = p1
                        (self).f12 = default__.fm4(d_539_v5_, globalState)
                        d_592_v48_: _dafny.Map
                        d_592_v48_ = _dafny.Map({self.f12: default__.fm4(d_538_v4_, globalState)})
                        d_592_v48_ = (d_592_v48_).set(self.f12, self.f12)
                        (self).f12 = self.f12
                        index116_ = default__.safeIndex(817, (d_533_v1_).length(0))
                        (d_533_v1_)[index116_] = self.f12
                        d_593_v49_: str
                        d_593_v49_ = _dafny.CodePoint('q')
                        d_594_v50_: _dafny.Array
                        nw105_ = _dafny.Array(None, 22)
                        nw105_[int(0)] = d_593_v49_
                        nw105_[int(1)] = d_593_v49_
                        nw105_[int(2)] = d_593_v49_
                        nw105_[int(3)] = d_593_v49_
                        nw105_[int(4)] = d_593_v49_
                        nw105_[int(5)] = d_593_v49_
                        nw105_[int(6)] = d_593_v49_
                        nw105_[int(7)] = d_593_v49_
                        nw105_[int(8)] = d_593_v49_
                        nw105_[int(9)] = d_593_v49_
                        nw105_[int(10)] = d_593_v49_
                        nw105_[int(11)] = d_593_v49_
                        nw105_[int(12)] = d_593_v49_
                        nw105_[int(13)] = d_593_v49_
                        nw105_[int(14)] = d_593_v49_
                        nw105_[int(15)] = d_593_v49_
                        nw105_[int(16)] = d_593_v49_
                        nw105_[int(17)] = d_593_v49_
                        nw105_[int(18)] = d_593_v49_
                        nw105_[int(19)] = d_593_v49_
                        nw105_[int(20)] = d_593_v49_
                        nw105_[int(21)] = d_593_v49_
                        d_594_v50_ = nw105_
                        d_595_v51_: _dafny.Map
                        d_595_v51_ = _dafny.Map({False: len(p2)})
                        d_596_v52_: _dafny.Map
                        d_596_v52_ = _dafny.Map({d_594_v50_: ((d_595_v51_)[self.f12] if (self.f12) in (d_595_v51_) else p0)})
                        index117_ = default__.safeIndex(817, (d_533_v1_).length(0))
                        rhs96_ = (True) and (not(self.f12))
                        rhs97_ = self.f12
                        rhs98_ = d_596_v52_
                        rhs99_ = (((_dafny.SeqWithoutIsStrInference([d_593_v49_ for d_597_i8_ in range(default__.abs(508))])) + (d_531_v0_)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "anvdwavn")))).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([self.f12])), len(((_dafny.SeqWithoutIsStrInference([d_593_v49_ for d_598_i8_ in range(default__.abs(508))])) + (d_531_v0_)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "anvdwavn"))))), d_593_v49_)
                        rhs100_ = self.f12
                        lhs72_ = self
                        lhs73_ = d_533_v1_
                        lhs74_ = default__.safeIndex(817, (d_533_v1_).length(0))
                        lhs75_ = self
                        lhs72_.f12 = rhs96_
                        lhs73_[lhs74_] = rhs97_
                        d_596_v52_ = rhs98_
                        d_531_v0_ = rhs99_
                        lhs75_.f12 = rhs100_
                    pass
            pass
        d_599_v53_: T2
        nw106_ = C0()
        nw106_.ctor__()
        d_599_v53_ = nw106_
        d_600_v54_: _dafny.Seq
        d_600_v54_ = _dafny.SeqWithoutIsStrInference([d_599_v53_, d_599_v53_, d_599_v53_, d_599_v53_, d_599_v53_])
        d_601_v55_: _dafny.Map
        d_601_v55_ = _dafny.Map({self.f12: self.f12})
        d_602_v56_: _dafny.MultiSet
        d_602_v56_ = _dafny.MultiSet([(0) - (len(d_601_v55_)), p1, len(_dafny.Map({p0: self.f12})), (self).f13])
        d_603_v57_: _dafny.MultiSet
        d_603_v57_ = _dafny.MultiSet([623, (d_602_v56_).cardinality, (self).f13])
        d_604_v58_: _dafny.Map
        d_604_v58_ = _dafny.Map({False: d_599_v53_})
        d_605_v59_: _dafny.Map
        d_605_v59_ = _dafny.Map({802: d_599_v53_})
        d_606_v60_: _dafny.Map
        d_606_v60_ = _dafny.Map({self.f12: d_533_v1_})
        d_607_v61_: _dafny.Array
        nw107_ = _dafny.Array(None, 20)
        nw107_[int(0)] = d_599_v53_
        nw107_[int(1)] = d_599_v53_
        nw107_[int(2)] = d_599_v53_
        nw107_[int(3)] = d_599_v53_
        nw107_[int(4)] = d_599_v53_
        nw107_[int(5)] = (d_600_v54_)[default__.safeIndex((self).fm11(d_602_v56_, globalState), len(d_600_v54_))]
        nw107_[int(6)] = d_599_v53_
        nw107_[int(7)] = d_599_v53_
        nw107_[int(8)] = d_599_v53_
        nw107_[int(9)] = ((d_604_v58_)[self.f12] if (self.f12) in (d_604_v58_) else d_599_v53_)
        nw107_[int(10)] = d_599_v53_
        nw107_[int(11)] = d_599_v53_
        nw107_[int(12)] = d_599_v53_
        nw107_[int(13)] = d_599_v53_
        nw107_[int(14)] = d_599_v53_
        nw107_[int(15)] = d_599_v53_
        nw107_[int(16)] = d_599_v53_
        nw107_[int(17)] = d_599_v53_
        nw107_[int(18)] = d_599_v53_
        nw107_[int(19)] = ((d_605_v59_)[len(d_606_v60_)] if (len(d_606_v60_)) in (d_605_v59_) else d_599_v53_)
        d_607_v61_ = nw107_
        d_608_v62_: _dafny.Seq
        d_608_v62_ = _dafny.SeqWithoutIsStrInference([d_607_v61_, d_607_v61_, d_607_v61_, d_607_v61_])
        d_608_v62_ = (d_608_v62_) + ((d_608_v62_) + (d_608_v62_))
        d_609_v63_: C0
        nw108_ = C0()
        nw108_.ctor__()
        d_609_v63_ = nw108_
        d_610_v64_: T0
        nw109_ = C1()
        nw109_.ctor__(self.f12, self.f12, (self).f10)
        d_610_v64_ = nw109_
        d_611_v65_: _dafny.Set
        d_611_v65_ = _dafny.Set({d_610_v64_})
        d_612_v66_: _dafny.Map
        d_612_v66_ = _dafny.Map({d_611_v65_: False})
        r0 = (d_612_v66_) | (d_612_v66_)
        return r0

    @property
    def f13(self):
        return self._f13

class C3(T1):
    def  __init__(self):
        self._f16: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    def ctor__(self, f16):
        (self)._f16 = f16

    def fm30(self, globalState):
        def iife62_():
            coll34_ = _dafny.Set()
            compr_34_: int
            for compr_34_ in _dafny.IntegerRange(-12, -592):
                d_614_v0_: int = compr_34_
                if ((-12) <= (d_614_v0_)) and ((d_614_v0_) < (-592)):
                    coll34_ = coll34_.union(_dafny.Set([(d_614_v0_) * (len(_dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ypcu")))): not((self).f16)})))]))
            return _dafny.Set(coll34_)
        return (_dafny.Map({_dafny.Set({len(_dafny.Map({(self).f16: (self).f16})), len(_dafny.Set({(0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_613_i0_ in range(default__.abs(-822))]))), len(_dafny.Map({-739: (self).f16})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nfqbkcer")))})), 956, (0) - (len(_dafny.Map({(self).f16: 949})))}): D9_DC23(48)})) | (_dafny.Map({iife62_()
        : D9_DC23(len(_dafny.Map({304: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uqqfvcsyh")))})))}))

    def fm31(self, p0, p1, p2, globalState):
        return 806

    def m1(self, p0, p1, p2, globalState):
        d_615_v0_: _dafny.Set
        d_615_v0_ = _dafny.Set({p1, -38})
        d_616_v1_: _dafny.Seq
        d_616_v1_ = _dafny.SeqWithoutIsStrInference([p1])
        (self).m7(default__.fm4(d_615_v0_, globalState), d_616_v1_, globalState)
        d_617_v2_: _dafny.Array
        nw110_ = _dafny.Array(False, 10)
        d_617_v2_ = nw110_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_617_v2_).length(0)):
            d_618_i0_: int = guard_loop_2_
            if (True) and (((0) <= (d_618_i0_)) and ((d_618_i0_) < ((d_617_v2_).length(0)))):
                (d_617_v2_)[(d_618_i0_)] = (self).f16
        d_619_v3_: _dafny.Set
        d_619_v3_ = _dafny.Set({False})
        d_620_v4_: _dafny.Seq
        d_620_v4_ = _dafny.SeqWithoutIsStrInference([not((self).f16)])
        d_621_v5_: C1
        nw111_ = C1()
        nw111_.ctor__((self).f16, (self).f16, _dafny.MultiSet(d_620_v4_))
        d_621_v5_ = nw111_
        d_622_v6_: _dafny.Seq
        d_622_v6_ = _dafny.SeqWithoutIsStrInference([d_621_v5_, d_621_v5_])
        d_623_v7_: _dafny.MultiSet
        d_623_v7_ = _dafny.MultiSet([p1, len(d_619_v3_), p1, len(d_622_v6_), p1])
        d_624_v8_: _dafny.Array
        nw112_ = _dafny.Array(_dafny.Array(None, 0), 14)
        d_624_v8_ = nw112_
        d_625_v9_: _dafny.Map
        d_625_v9_ = _dafny.Map({(d_623_v7_) - (d_623_v7_): d_624_v8_})
        d_625_v9_ = (d_625_v9_).set(d_623_v7_, d_624_v8_)
        if (793) <= ((0) - (-839)):
            (globalState).f1 = p1
            d_626_v10_: _dafny.Seq
            d_626_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cxakbjhpc"))
            d_627_v11_: D1
            d_627_v11_ = D1_DC5(748, (d_621_v5_).f15, len(d_626_v10_), (self).f16, p1)
            d_628_v12_: _dafny.Map
            d_628_v12_ = _dafny.Map({p1: p1})
            d_629_v13_: _dafny.Map
            d_629_v13_ = _dafny.Map({(self).f16: (d_621_v5_).f15})
            d_630_v14_: _dafny.Array
            nw113_ = _dafny.Array(None, 11)
            nw113_[int(0)] = (d_626_v10_) + (d_626_v10_)
            nw113_[int(1)] = d_626_v10_
            nw113_[int(2)] = d_626_v10_
            nw113_[int(3)] = d_626_v10_
            nw113_[int(4)] = d_626_v10_
            nw113_[int(5)] = (((d_626_v10_).set(default__.safeIndex(len(default__.fm32((d_627_v11_).cf14, p0, d_628_v12_, (d_621_v5_).f15, globalState)), len(d_626_v10_)), p0)).set(default__.safeIndex(len(d_629_v13_), len((d_626_v10_).set(default__.safeIndex(len(default__.fm32((d_627_v11_).cf14, p0, d_628_v12_, (d_621_v5_).f15, globalState)), len(d_626_v10_)), p0))), p0) if (d_621_v5_).f14 else d_626_v10_)
            nw113_[int(6)] = (d_626_v10_) + (d_626_v10_)
            nw113_[int(7)] = (_dafny.SeqWithoutIsStrInference([p0 for d_631_i1_ in range(default__.abs(555))]) if (d_621_v5_).f14 else d_626_v10_)
            nw113_[int(8)] = (d_626_v10_) + (d_626_v10_)
            nw113_[int(9)] = d_626_v10_
            nw113_[int(10)] = d_626_v10_
            d_630_v14_ = nw113_
            d_630_v14_ = d_630_v14_
            d_632_v15_: _dafny.Array
            def lambda36_(d_633_p1_):
                def lambda37_(d_634_i2_):
                    return default__.safeModuloInt(d_634_i2_, d_633_p1_)

                return lambda37_

            init20_ = lambda36_(p1)
            nw114_ = _dafny.Array(None, 28)
            for i0_20_ in range(nw114_.length(0)):
                nw114_[i0_20_] = init20_(i0_20_)
            d_632_v15_ = nw114_
            index118_ = default__.safeIndex(713, (d_632_v15_).length(0))
            (d_632_v15_)[index118_] = p1
            d_635_v16_: _dafny.Seq
            d_635_v16_ = _dafny.SeqWithoutIsStrInference([d_626_v10_])
            d_636_v17_: _dafny.MultiSet
            d_636_v17_ = _dafny.MultiSet([(d_621_v5_).f15, (d_621_v5_).f14, not(False)])
            index119_ = default__.safeIndex(713, (d_632_v15_).length(0))
            rhs101_ = (0) - (default__.safeDivisionInt((71 if (self).f16 else len(d_635_v16_)), (p1) * (len(d_629_v13_))))
            rhs102_ = p1
            rhs103_ = len(default__.fm33(((d_636_v17_)[(d_621_v5_).f15] if ((d_621_v5_).f15) in (d_636_v17_) else p1), (p1) > (p1), (d_621_v5_).f15, p1, globalState))
            rhs104_ = ((d_623_v7_)[p1] if (p1) in (d_623_v7_) else (p1) * (len(d_620_v4_)))
            lhs76_ = globalState
            lhs77_ = globalState
            lhs78_ = d_632_v15_
            lhs79_ = default__.safeIndex(713, (d_632_v15_).length(0))
            lhs80_ = globalState
            lhs76_.f1 = rhs101_
            lhs77_.f1 = rhs102_
            lhs78_[lhs79_] = rhs103_
            lhs80_.f1 = rhs104_
            d_615_v0_ = d_615_v0_
            (globalState).f1 = (d_632_v15_)[default__.safeIndex(713, (d_632_v15_).length(0))]
        elif True:
            (globalState).f1 = ((d_623_v7_)[len((_dafny.SeqWithoutIsStrInference([(self).f16]) if (self).f16 else d_620_v4_))] if (len((_dafny.SeqWithoutIsStrInference([(self).f16]) if (self).f16 else d_620_v4_))) in (d_623_v7_) else p1)
            d_637_v18_: bool
            d_637_v18_ = False
            rhs105_ = p1
            rhs106_ = (d_621_v5_).f15
            lhs81_ = globalState
            lhs81_.f1 = rhs105_
            d_637_v18_ = rhs106_
            d_638_v19_: _dafny.Array
            nw115_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 22)
            d_638_v19_ = nw115_
            def lambda38_(d_639_i3_):
                return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dbl"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x")))

            init21_ = lambda38_
            nw116_ = _dafny.Array(None, 5)
            for i0_21_ in range(nw116_.length(0)):
                nw116_[i0_21_] = init21_(i0_21_)
            d_638_v19_ = nw116_
            d_640_v20_: _dafny.MultiSet
            d_640_v20_ = _dafny.MultiSet([(d_621_v5_).f14])
            d_640_v20_ = (d_640_v20_) - (d_640_v20_)
            d_641_v21_: _dafny.Seq
            d_641_v21_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xwcdnmpw"))
            d_642_v22_: _dafny.Map
            d_642_v22_ = _dafny.Map({len(_dafny.Set({(self).f16})): d_641_v21_})
            rhs107_ = (p1) == (p1)
            rhs108_ = _dafny.Map({(p1) + (p1): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oumom"))})
            d_637_v18_ = rhs107_
            d_642_v22_ = rhs108_
        hi5_ = default__.safeDivisionInt(p1, p1)
        for d_643_i4_ in range(default__.safeModuloInt(p1, p1), hi5_):
            d_644_v23_: _dafny.Array
            def lambda39_(d_645_p1_):
                def lambda40_(d_646_i5_):
                    return (d_646_i5_) - (d_645_p1_)

                return lambda40_

            init22_ = lambda39_(p1)
            nw117_ = _dafny.Array(None, 27)
            for i0_22_ in range(nw117_.length(0)):
                nw117_[i0_22_] = init22_(i0_22_)
            d_644_v23_ = nw117_
            index120_ = default__.safeIndex(656, (d_644_v23_).length(0))
            (d_644_v23_)[index120_] = d_643_i4_
            index121_ = default__.safeIndex(656, (d_644_v23_).length(0))
            (d_644_v23_)[index121_] = p1
            d_647_v24_: str
            d_647_v24_ = _dafny.CodePoint('h')
            d_648_v25_: _dafny.Seq
            d_648_v25_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qs"))
            d_649_v26_: _dafny.Map
            d_649_v26_ = _dafny.Map({not((len(d_648_v25_)) >= (d_643_i4_)): (d_621_v5_).f15})
            d_650_v27_: _dafny.Map
            d_650_v27_ = _dafny.Map({p1: (self).f16})
            d_651_v28_: D9
            d_651_v28_ = D9_DC22(p1, (self).f16, d_650_v27_)
            d_652_v29_: _dafny.Map
            d_652_v29_ = _dafny.Map({(d_621_v5_).f15: d_650_v27_})
            d_653_v30_: _dafny.Seq
            d_653_v30_ = _dafny.SeqWithoutIsStrInference([d_650_v27_, default__.fm34((d_651_v28_).cf50, d_620_v4_, (d_616_v1_)[default__.safeIndex(len(d_615_v0_), len(d_616_v1_))], (d_621_v5_).f15, globalState), d_650_v27_, d_650_v27_, ((d_652_v29_)[(d_621_v5_).f14] if ((d_621_v5_).f14) in (d_652_v29_) else ((d_652_v29_)[(d_621_v5_).f15] if ((d_621_v5_).f15) in (d_652_v29_) else d_650_v27_))])
            d_654_v31_: _dafny.Map
            d_654_v31_ = _dafny.Map({(d_621_v5_).f14: p1})
            d_655_v32_: D8
            d_655_v32_ = D8_DC20(d_647_v24_, (d_621_v5_).f15, default__.fm4(_dafny.Set({(d_644_v23_)[default__.safeIndex(656, (d_644_v23_).length(0))], d_643_i4_, p1}), globalState), d_654_v31_, (d_621_v5_).f14)
            d_656_v33_: _dafny.Map
            d_656_v33_ = _dafny.Map({d_655_v32_: p1})
            d_657_v34_: _dafny.Map
            d_657_v34_ = _dafny.Map({p1: _dafny.Map({D8_DC20(d_647_v24_, (d_621_v5_).f15, (d_621_v5_).f15, _dafny.Map({(d_621_v5_).f15: p1}), (self).f16): d_643_i4_})})
            index122_ = default__.safeIndex(656, (d_644_v23_).length(0))
            rhs109_ = (0) - ((self).fm31(d_647_v24_, (self).f16, (_dafny.Map({(d_644_v23_)[default__.safeIndex(656, (d_644_v23_).length(0))]: d_656_v33_}) if (d_621_v5_).f15 else d_657_v34_), globalState))
            rhs110_ = default__.fm1(globalState)
            rhs111_ = (d_649_v26_) | (d_649_v26_)
            rhs112_ = ((d_653_v30_) + (d_653_v30_)) + (d_653_v30_)
            lhs82_ = d_644_v23_
            lhs83_ = default__.safeIndex(656, (d_644_v23_).length(0))
            lhs82_[lhs83_] = rhs109_
            d_647_v24_ = rhs110_
            d_649_v26_ = rhs111_
            d_653_v30_ = rhs112_
            d_648_v25_ = d_648_v25_
            d_658_v35_: bool
            d_658_v35_ = True
            d_658_v35_ = (self).f16
        d_659_v36_: _dafny.Array
        nw118_ = _dafny.Array(int(0), 14)
        d_659_v36_ = nw118_
        d_660_v37_: _dafny.Seq
        d_660_v37_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nqgf"))
        index123_ = default__.safeIndex(450, (d_659_v36_).length(0))
        (d_659_v36_)[index123_] = len(d_660_v37_)
        index124_ = default__.safeIndex(450, (d_659_v36_).length(0))
        (d_659_v36_)[index124_] = p1

    def m7(self, p0, p1, globalState):
        d_661_v0_: int
        d_661_v0_ = 53
        d_662_v1_: _dafny.Map
        d_662_v1_ = _dafny.Map({(d_661_v0_) + (d_661_v0_): d_661_v0_})
        d_662_v1_ = (d_662_v1_).set(d_661_v0_, d_661_v0_)
        hi6_ = d_661_v0_
        for d_663_i0_ in range(default__.safeDivisionInt(d_661_v0_, d_661_v0_), hi6_):
            d_661_v0_ = default__.safeModuloInt(d_663_i0_, 689)
            d_664_v2_: _dafny.Set
            d_664_v2_ = _dafny.Set({(d_663_i0_) - (d_661_v0_), 597, d_663_i0_, d_663_i0_})
            (globalState).f1 = len(d_664_v2_)
            if (self).f16:
                (globalState).f1 = default__.safeModuloInt(d_663_i0_, d_661_v0_)
                (globalState).f1 = d_661_v0_
                (globalState).f1 = default__.safeModuloInt((0) - (d_663_i0_), (0) - (len((d_664_v2_ if p0 else _dafny.Set({d_663_i0_})))))
                d_665_v3_: _dafny.MultiSet
                d_665_v3_ = _dafny.MultiSet([p0])
                d_666_v4_: C1
                nw119_ = C1()
                nw119_.ctor__(False, p0, d_665_v3_)
                d_666_v4_ = nw119_
                d_667_v5_: _dafny.Seq
                d_667_v5_ = _dafny.SeqWithoutIsStrInference([d_666_v4_, d_666_v4_, d_666_v4_])
                d_661_v0_ = default__.safeDivisionInt((0) - (len(d_667_v5_)), (0) - (d_663_i0_))
                d_668_v6_: _dafny.Array
                def lambda41_(d_669_i1_):
                    return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wruaxewg"))

                init23_ = lambda41_
                nw120_ = _dafny.Array(None, 23)
                for i0_23_ in range(nw120_.length(0)):
                    nw120_[i0_23_] = init23_(i0_23_)
                d_668_v6_ = nw120_
                d_670_v7_: _dafny.Array
                nw121_ = _dafny.Array(None, 12)
                nw121_[int(0)] = d_661_v0_
                nw121_[int(1)] = 174
                nw121_[int(2)] = d_661_v0_
                nw121_[int(3)] = d_661_v0_
                nw121_[int(4)] = d_663_i0_
                nw121_[int(5)] = d_663_i0_
                nw121_[int(6)] = d_661_v0_
                nw121_[int(7)] = d_663_i0_
                nw121_[int(8)] = 687
                nw121_[int(9)] = d_663_i0_
                nw121_[int(10)] = d_663_i0_
                nw121_[int(11)] = d_661_v0_
                d_670_v7_ = nw121_
                d_671_v8_: _dafny.Array
                nw122_ = _dafny.Array(_dafny.Map({}), 20)
                d_671_v8_ = nw122_
                d_672_v9_: D5
                d_672_v9_ = D5_DC13(d_663_i0_, d_670_v7_, d_671_v8_, d_663_i0_)
                d_673_v10_: _dafny.Map
                d_673_v10_ = _dafny.Map({d_668_v6_: d_672_v9_})
                d_673_v10_ = (d_673_v10_).set(d_668_v6_, d_672_v9_)
            elif True:
                d_674_v11_: D9
                d_674_v11_ = D9_DC24(d_663_i0_)
                d_675_v12_: _dafny.Map
                d_675_v12_ = _dafny.Map({d_674_v11_: (self).f16})
                (globalState).f1 = (p1)[default__.safeIndex(len(d_675_v12_), len(p1))]
                d_676_v13_: C0
                nw123_ = C0()
                nw123_.ctor__()
                d_676_v13_ = nw123_
                d_676_v13_ = d_676_v13_
                d_677_v14_: _dafny.MultiSet
                d_677_v14_ = _dafny.MultiSet([p0])
                d_677_v14_ = d_677_v14_
                d_678_v15_: C2
                nw124_ = C2()
                nw124_.ctor__(p0, d_661_v0_, d_677_v14_)
                d_678_v15_ = nw124_
                (globalState).f1 = d_661_v0_
            d_679_v16_: C0
            nw125_ = C0()
            nw125_.ctor__()
            d_679_v16_ = nw125_
        hi7_ = len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "edunij"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rmlmct"))))
        for d_680_i2_ in range(d_661_v0_, hi7_):
            d_681_v17_: _dafny.Set
            d_681_v17_ = _dafny.Set({not(True), (self).f16})
            d_682_v18_: D0
            d_682_v18_ = D0_DC1(p0, d_661_v0_)
            d_683_v19_: _dafny.Array
            nw126_ = _dafny.Array(None, 4)
            nw126_[int(0)] = d_680_i2_
            nw126_[int(1)] = d_680_i2_
            nw126_[int(2)] = (0) - (len(d_681_v17_))
            nw126_[int(3)] = (d_682_v18_).cf2
            d_683_v19_ = nw126_
            index125_ = default__.safeIndex(2, (d_683_v19_).length(0))
            (d_683_v19_)[index125_] = d_661_v0_
            d_684_v20_: _dafny.Seq
            d_684_v20_ = _dafny.SeqWithoutIsStrInference([p0, (self).f16])
            index126_ = default__.safeIndex(2, (d_683_v19_).length(0))
            (d_683_v19_)[index126_] = ((0) - (len((d_684_v20_ if (self).f16 else (d_684_v20_).set(default__.safeIndex(d_680_i2_, len(d_684_v20_)), (self).f16))))) + (d_661_v0_)
            d_685_v21_: _dafny.Array
            def lambda42_(d_686_v0_, d_687_i2_):
                def lambda43_(d_688_i3_):
                    return _dafny.Map({d_686_v0_: d_687_i2_})

                return lambda43_

            init24_ = lambda42_(d_661_v0_, d_680_i2_)
            nw127_ = _dafny.Array(None, 11)
            for i0_24_ in range(nw127_.length(0)):
                nw127_[i0_24_] = init24_(i0_24_)
            d_685_v21_ = nw127_
            d_685_v21_ = d_685_v21_
            d_689_v22_: _dafny.MultiSet
            d_689_v22_ = _dafny.MultiSet([d_661_v0_])
            d_661_v0_ = ((d_689_v22_)[(p1)[default__.safeIndex(d_661_v0_, len(p1))]] if ((p1)[default__.safeIndex(d_661_v0_, len(p1))]) in (d_689_v22_) else default__.safeModuloInt(d_680_i2_, 409))
            d_690_v23_: str
            d_690_v23_ = _dafny.CodePoint('m')
            d_691_v24_: _dafny.Seq
            d_691_v24_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({p0: p0}), _dafny.Map({(self).f16: (self).f16})])
            d_692_v25_: _dafny.Map
            d_692_v25_ = _dafny.Map({p0: d_661_v0_})
            d_693_v26_: D8
            d_693_v26_ = D8_DC20(d_690_v23_, (self).f16, not((self).f16), d_692_v25_, p0)
            d_694_v27_: _dafny.Map
            d_694_v27_ = _dafny.Map({d_693_v26_: d_661_v0_})
            d_695_v28_: _dafny.Map
            d_695_v28_ = _dafny.Map({len(d_691_v24_): d_694_v27_})
            index127_ = default__.safeIndex(2, (d_683_v19_).length(0))
            (d_683_v19_)[index127_] = (self).fm31(d_690_v23_, (len(p1)) >= ((d_683_v19_)[default__.safeIndex(2, (d_683_v19_).length(0))]), d_695_v28_, globalState)
        d_696_v29_: _dafny.MultiSet
        d_696_v29_ = _dafny.MultiSet([d_661_v0_, d_661_v0_])
        d_697_v30_: _dafny.Array
        nw128_ = _dafny.Array(None, 4)
        nw128_[int(0)] = d_696_v29_
        nw128_[int(1)] = d_696_v29_
        nw128_[int(2)] = (d_696_v29_) - (_dafny.MultiSet(p1))
        nw128_[int(3)] = _dafny.MultiSet(p1)
        d_697_v30_ = nw128_
        d_697_v30_ = d_697_v30_
        d_698_v31_: _dafny.Map
        d_698_v31_ = _dafny.Map({d_662_v1_: d_661_v0_})
        d_699_v32_: bool
        d_699_v32_ = True
        d_700_v33_: _dafny.Seq
        d_700_v33_ = _dafny.SeqWithoutIsStrInference([p0])
        d_701_v34_: _dafny.Seq
        d_701_v34_ = _dafny.SeqWithoutIsStrInference([d_700_v33_])
        d_702_v35_: _dafny.Set
        d_702_v35_ = _dafny.Set({p0})
        rhs113_ = (d_698_v31_).set(((d_662_v1_).set(len((d_701_v34_)[default__.safeIndex(d_661_v0_, len(d_701_v34_))]), d_661_v0_)) | (d_662_v1_), len((d_702_v35_).intersection(d_702_v35_)))
        rhs114_ = (self).f16
        d_698_v31_ = rhs113_
        d_699_v32_ = rhs114_
        d_703_v36_: str
        d_703_v36_ = _dafny.CodePoint('l')
        d_704_v37_: _dafny.Array
        nw129_ = _dafny.Array(False, 5)
        d_704_v37_ = nw129_
        index128_ = default__.safeIndex(257, (d_704_v37_).length(0))
        (d_704_v37_)[index128_] = d_699_v32_
        index129_ = default__.safeIndex(257, (d_704_v37_).length(0))
        rhs115_ = default__.fm1(globalState)
        rhs116_ = p0
        lhs84_ = d_704_v37_
        lhs85_ = default__.safeIndex(257, (d_704_v37_).length(0))
        d_703_v36_ = rhs115_
        lhs84_[lhs85_] = rhs116_

    def m8(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        d_705_v0_: bool
        d_705_v0_ = False
        d_705_v0_ = (True) and ((p2 if True else p2))
        d_706_v1_: _dafny.Seq
        d_706_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "le"))
        d_707_v2_: _dafny.Map
        d_707_v2_ = _dafny.Map({p2: len(d_706_v1_)})
        d_708_v3_: _dafny.Set
        d_708_v3_ = _dafny.Set({len(d_707_v2_)})
        d_709_v4_: _dafny.MultiSet
        d_709_v4_ = _dafny.MultiSet([p3])
        d_710_v5_: _dafny.MultiSet
        d_710_v5_ = _dafny.MultiSet([p0, (self).f16])
        d_711_v6_: C1
        nw130_ = C1()
        nw130_.ctor__((p2) and (default__.fm4(d_708_v3_, globalState)), (d_709_v4_).ispropersubset(d_709_v4_), (d_710_v5_).set((self).f16, default__.abs((0) - (p3))))
        d_711_v6_ = nw130_
        if default__.fm4(d_708_v3_, globalState):
            d_712_v7_: _dafny.Array
            nw131_ = _dafny.Array(int(0), 12)
            d_712_v7_ = nw131_
            d_713_v8_: _dafny.Map
            d_713_v8_ = _dafny.Map({p3: False})
            d_714_v9_: _dafny.Map
            d_714_v9_ = _dafny.Map({(self).f16: ((d_713_v8_)[593] if (593) in (d_713_v8_) else True)})
            index130_ = default__.safeIndex(634, (d_712_v7_).length(0))
            (d_712_v7_)[index130_] = len((d_714_v9_) | ((d_714_v9_).set(not(p2), p2)))
            index131_ = default__.safeIndex(634, (d_712_v7_).length(0))
            (d_712_v7_)[index131_] = default__.safeDivisionInt(len((d_706_v1_) + (d_706_v1_)), p1)
            d_715_v10_: _dafny.Seq
            d_715_v10_ = _dafny.SeqWithoutIsStrInference([(d_711_v6_).f14, ((d_712_v7_)[default__.safeIndex(634, (d_712_v7_).length(0))]) >= (p1), (d_711_v6_).f14, not((self).f16), (d_711_v6_).f14])
            d_715_v10_ = (d_715_v10_ if False else (d_715_v10_ if (d_711_v6_).f14 else d_715_v10_))
            d_716_v11_: D8
            d_716_v11_ = D8_DC19(d_713_v8_)
            d_717_v12_: _dafny.Seq
            d_717_v12_ = _dafny.SeqWithoutIsStrInference([d_716_v11_, d_716_v11_])
            d_718_v13_: _dafny.Seq
            d_718_v13_ = _dafny.SeqWithoutIsStrInference([d_717_v12_, _dafny.SeqWithoutIsStrInference([d_716_v11_ for d_719_i0_ in range(default__.abs(-57))]), d_717_v12_, d_717_v12_])
            d_720_v15_: _dafny.Set
            d_720_v15_ = _dafny.Set({d_717_v12_})
            def iife63_():
                coll35_ = _dafny.Set()
                compr_35_: _dafny.Seq
                for compr_35_ in (d_718_v13_).Elements:
                    d_721_v14_: _dafny.Seq = compr_35_
                    if (d_721_v14_) in (d_718_v13_):
                        coll35_ = coll35_.union(_dafny.Set([d_721_v14_]))
                return _dafny.Set(coll35_)
            if ((d_720_v15_) | (d_720_v15_)).ispropersubset(iife63_()
            ):
                d_722_v16_: C1
                nw132_ = C1()
                nw132_.ctor__(p2, (d_711_v6_).f14, _dafny.MultiSet([p2, (d_711_v6_).f14, p2]))
                d_722_v16_ = nw132_
                (globalState).f1 = (p3) * ((p3) * ((d_712_v7_)[default__.safeIndex(634, (d_712_v7_).length(0))]))
                index132_ = default__.safeIndex(634, (d_712_v7_).length(0))
                (d_712_v7_)[index132_] = (0) - ((0) - (p3))
                r0 = (p1) + (p3)
                d_705_v0_ = (d_722_v16_).f15
            elif True:
                d_710_v5_ = d_710_v5_
                d_705_v0_ = (d_711_v6_).f14
                d_723_v17_: _dafny.Array
                nw133_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 22)
                d_723_v17_ = nw133_
                index133_ = default__.safeIndex(574, (d_723_v17_).length(0))
                (d_723_v17_)[index133_] = d_706_v1_
                index134_ = default__.safeIndex(574, (d_723_v17_).length(0))
                (d_723_v17_)[index134_] = ((d_706_v1_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v")))) + (d_706_v1_)
                d_705_v0_ = p0
                nw134_ = C1()
                nw134_.ctor__((d_711_v6_).f14, (d_715_v10_)[default__.safeIndex(p1, len(d_715_v10_))], d_710_v5_)
                d_711_v6_ = nw134_
            d_724_v18_: D0
            d_724_v18_ = D0_DC1((self).f16, (d_712_v7_)[default__.safeIndex(634, (d_712_v7_).length(0))])
            d_725_v20_: _dafny.Map
            d_725_v20_ = _dafny.Map({395: (d_712_v7_)[default__.safeIndex(634, (d_712_v7_).length(0))]})
            d_726_v21_: _dafny.MultiSet
            def iife64_():
                coll36_ = _dafny.Map()
                compr_36_: int
                for compr_36_ in (d_725_v20_).keys.Elements:
                    d_727_v19_: int = compr_36_
                    if (d_727_v19_) in (d_725_v20_):
                        coll36_[default__.safeModuloInt(d_727_v19_, p1)] = len(_dafny.SeqWithoutIsStrInference([p1]))
                return _dafny.Map(coll36_)
            d_726_v21_ = _dafny.MultiSet([iife64_()
            ])
            index135_ = default__.safeIndex(634, (d_712_v7_).length(0))
            (d_712_v7_)[index135_] = ((d_724_v18_).cf2) + (((d_726_v21_)[d_725_v20_] if (d_725_v20_) in (d_726_v21_) else p3))
            d_728_v22_: C0
            nw135_ = C0()
            nw135_.ctor__()
            d_728_v22_ = nw135_
        elif True:
            d_729_v23_: _dafny.Array
            def lambda44_(d_730_p0_, d_731_v3_):
                def lambda45_(d_732_i1_):
                    return _dafny.Map({d_730_p0_: len(d_731_v3_)})

                return lambda45_

            init25_ = lambda44_(p0, d_708_v3_)
            nw136_ = _dafny.Array(None, 21)
            for i0_25_ in range(nw136_.length(0)):
                nw136_[i0_25_] = init25_(i0_25_)
            d_729_v23_ = nw136_
            index136_ = default__.safeIndex(716, (d_729_v23_).length(0))
            (d_729_v23_)[index136_] = d_707_v2_
            index137_ = default__.safeIndex(716, (d_729_v23_).length(0))
            (d_729_v23_)[index137_] = d_707_v2_
            r0 = (p3) + (714)
            d_733_v24_: T2
            nw137_ = C0()
            nw137_.ctor__()
            d_733_v24_ = nw137_
            d_734_v25_: _dafny.Map
            d_734_v25_ = _dafny.Map({d_733_v24_: True})
            d_734_v25_ = (d_734_v25_).set(d_733_v24_, p2)
            d_735_v26_: _dafny.Map
            d_735_v26_ = _dafny.Map({(p1) > (952): (d_711_v6_).f14})
            d_735_v26_ = (d_735_v26_).set(not((d_706_v1_) == (d_706_v1_)), d_705_v0_)
            if (p0) == ((d_711_v6_).f15):
                d_705_v0_ = (d_711_v6_).f15
                (globalState).f1 = p1
                d_736_v27_: _dafny.Map
                d_736_v27_ = _dafny.Map({p3: (d_711_v6_).f14})
                d_737_v28_: _dafny.Seq
                d_737_v28_ = _dafny.SeqWithoutIsStrInference([p1, p3])
                d_738_v29_: D1
                d_738_v29_ = D1_DC4(p2, (d_711_v6_).f15, p1, d_737_v28_, p3)
                d_739_v30_: _dafny.Seq
                d_739_v30_ = _dafny.SeqWithoutIsStrInference([p0, (d_711_v6_).f14])
                d_740_v31_: _dafny.Array
                nw138_ = _dafny.Array(None, 26)
                nw138_[int(0)] = (len(d_736_v27_)) - (p1)
                nw138_[int(1)] = p3
                nw138_[int(2)] = len((d_735_v26_) | (d_735_v26_))
                nw138_[int(3)] = p1
                nw138_[int(4)] = (d_710_v5_).cardinality
                nw138_[int(5)] = p1
                nw138_[int(6)] = (d_738_v29_).cf9
                nw138_[int(7)] = (d_733_v24_).fm6(p2, p1, globalState)
                nw138_[int(8)] = p1
                nw138_[int(9)] = ((d_709_v4_)[p3] if (p3) in (d_709_v4_) else p3)
                nw138_[int(10)] = p1
                nw138_[int(11)] = ((0) - (p1)) * ((d_710_v5_).cardinality)
                nw138_[int(12)] = p3
                nw138_[int(13)] = p1
                nw138_[int(14)] = len(d_739_v30_)
                nw138_[int(15)] = p3
                nw138_[int(16)] = len(d_735_v26_)
                nw138_[int(17)] = p3
                nw138_[int(18)] = p1
                nw138_[int(19)] = p3
                nw138_[int(20)] = p3
                nw138_[int(21)] = p1
                nw138_[int(22)] = (p1) + (len(d_706_v1_))
                nw138_[int(23)] = p1
                nw138_[int(24)] = p3
                nw138_[int(25)] = default__.safeDivisionInt(p1, p1)
                d_740_v31_ = nw138_
                index138_ = default__.safeIndex(787, (d_740_v31_).length(0))
                (d_740_v31_)[index138_] = p1
                index139_ = default__.safeIndex(787, (d_740_v31_).length(0))
                (d_740_v31_)[index139_] = default__.safeDivisionInt(p1, p3)
                d_705_v0_ = (self).f16
                d_705_v0_ = ((0) - ((d_740_v31_)[default__.safeIndex(787, (d_740_v31_).length(0))])) < (len(d_708_v3_))
            elif True:
                d_711_v6_ = d_711_v6_
                (globalState).f1 = p3
                d_741_v32_: _dafny.Seq
                d_741_v32_ = _dafny.SeqWithoutIsStrInference([len(d_706_v1_)])
                d_742_v33_: _dafny.Seq
                d_742_v33_ = _dafny.SeqWithoutIsStrInference([p2])
                d_743_v34_: C2
                nw139_ = C2()
                nw139_.ctor__(d_705_v0_, default__.safeModuloInt(-699, len(d_741_v32_)), _dafny.MultiSet((d_742_v33_) + (d_742_v33_)))
                d_743_v34_ = nw139_
                d_744_v35_: _dafny.Array
                def lambda46_(d_745_p3_):
                    def lambda47_(d_746_i2_):
                        return (d_746_i2_) + (d_745_p3_)

                    return lambda47_

                init26_ = lambda46_(p3)
                nw140_ = _dafny.Array(None, 2)
                for i0_26_ in range(nw140_.length(0)):
                    nw140_[i0_26_] = init26_(i0_26_)
                d_744_v35_ = nw140_
                index140_ = default__.safeIndex(295, (d_744_v35_).length(0))
                (d_744_v35_)[index140_] = (p1) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "et"))))
                d_747_v36_: str
                d_747_v36_ = _dafny.CodePoint('f')
                d_748_v37_: D0
                d_748_v37_ = D0_DC0(d_747_v36_)
                d_749_v38_: D0
                d_749_v38_ = D0_DC1((d_711_v6_).f15, p3)
                index141_ = default__.safeIndex(295, (d_744_v35_).length(0))
                (d_744_v35_)[index141_] = len(((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))) + (d_706_v1_)).set(default__.safeIndex(default__.safeModuloInt(p3, p3), len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))) + (d_706_v1_))), (d_743_v34_).fm10(d_748_v37_, D0_DC2(d_749_v38_), globalState)))
                r0 = (d_743_v34_).f13
        d_750_v39_: _dafny.Seq
        d_750_v39_ = _dafny.SeqWithoutIsStrInference([p2])
        d_751_v40_: C0
        nw141_ = C0()
        nw141_.ctor__()
        d_751_v40_ = nw141_
        d_752_v41_: _dafny.Seq
        d_752_v41_ = _dafny.SeqWithoutIsStrInference([d_751_v40_])
        d_753_v42_: _dafny.Seq
        d_753_v42_ = _dafny.SeqWithoutIsStrInference([d_752_v41_])
        d_754_v43_: _dafny.Map
        d_754_v43_ = _dafny.Map({d_750_v39_: (d_752_v41_) not in ((d_753_v42_).set(default__.safeIndex(p1, len(d_753_v42_)), d_752_v41_))})
        d_754_v43_ = (d_754_v43_).set((default__.fm33(p3, p0, (d_711_v6_).f15, p3, globalState)) + (((_dafny.SeqWithoutIsStrInference([p2])).set(default__.safeIndex(p3, len(_dafny.SeqWithoutIsStrInference([p2]))), (d_751_v40_).fm13(p2, globalState))).set(default__.safeIndex(p1, len((_dafny.SeqWithoutIsStrInference([p2])).set(default__.safeIndex(p3, len(_dafny.SeqWithoutIsStrInference([p2]))), (d_751_v40_).fm13(p2, globalState)))), p0)), d_705_v0_)
        d_755_i3_: int
        d_755_i3_ = 0
        with _dafny.label("3"):
            while (d_711_v6_).f14:
                with _dafny.c_label("3"):
                    if (d_755_i3_) >= (100):
                        raise _dafny.Break("3")
                    d_755_i3_ = (d_755_i3_) + (1)
                    d_756_v44_: _dafny.Seq
                    d_756_v44_ = _dafny.SeqWithoutIsStrInference([d_706_v1_])
                    (globalState).f1 = len(d_756_v44_)
                    d_757_v45_: C2
                    nw142_ = C2()
                    nw142_.ctor__((d_711_v6_).f15, 677, _dafny.MultiSet(d_750_v39_))
                    d_757_v45_ = nw142_
                    d_758_v46_: _dafny.Array
                    nw143_ = _dafny.Array(_dafny.CodePoint('D'), 11)
                    d_758_v46_ = nw143_
                    d_759_v47_: _dafny.Set
                    d_759_v47_ = _dafny.Set({(d_751_v40_).fm13(d_757_v45_.f12, globalState)})
                    d_760_v48_: _dafny.Seq
                    d_760_v48_ = _dafny.SeqWithoutIsStrInference([(d_757_v45_).f13, len(d_759_v47_), p3])
                    d_761_v49_: _dafny.Array
                    nw144_ = _dafny.Array(None, 8)
                    nw144_[int(0)] = p3
                    nw144_[int(1)] = p1
                    nw144_[int(2)] = (0) - (len(d_750_v39_))
                    nw144_[int(3)] = p1
                    nw144_[int(4)] = len(d_760_v48_)
                    nw144_[int(5)] = (d_757_v45_).f13
                    nw144_[int(6)] = (d_757_v45_).f13
                    nw144_[int(7)] = len(d_708_v3_)
                    d_761_v49_ = nw144_
                    d_762_v50_: _dafny.Map
                    d_762_v50_ = _dafny.Map({d_758_v46_: d_761_v49_})
                    d_762_v50_ = d_762_v50_
                    d_705_v0_ = (d_751_v40_).fm5((d_759_v47_).intersection(d_759_v47_), p3, len(d_707_v2_), globalState)
                    pass
            pass
        d_763_v51_: C1
        nw145_ = C1()
        nw145_.ctor__(p2, (self).f16, (d_710_v5_).set(p0, default__.abs(p3)))
        d_763_v51_ = nw145_
        r0 = p1
        return r0

    @property
    def f16(self):
        return self._f16

class C4(T1):
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    def ctor__(self):
        pass
        pass

    def fm9(self, p0, p1, globalState):
        return 921

    def m1(self, p0, p1, p2, globalState):
        d_764_v0_: bool
        d_764_v0_ = False
        d_765_i0_: int
        d_765_i0_ = 0
        with _dafny.label("4"):
            while d_764_v0_:
                with _dafny.c_label("4"):
                    if (d_765_i0_) >= (100):
                        raise _dafny.Break("4")
                    d_765_i0_ = (d_765_i0_) + (1)
                    d_766_v1_: _dafny.Set
                    d_766_v1_ = _dafny.Set({(0) - (p1), (p1) - (p1)})
                    (globalState).f1 = len(d_766_v1_)
                    d_767_v2_: _dafny.Array
                    def lambda48_(d_768_v0_):
                        def lambda49_(d_769_i1_):
                            return d_768_v0_

                        return lambda49_

                    init27_ = lambda48_(d_764_v0_)
                    nw146_ = _dafny.Array(None, 21)
                    for i0_27_ in range(nw146_.length(0)):
                        nw146_[i0_27_] = init27_(i0_27_)
                    d_767_v2_ = nw146_
                    index142_ = default__.safeIndex(910, (d_767_v2_).length(0))
                    (d_767_v2_)[index142_] = True
                    d_770_v3_: _dafny.MultiSet
                    d_770_v3_ = _dafny.MultiSet([d_764_v0_, d_764_v0_])
                    d_771_v4_: _dafny.Map
                    d_771_v4_ = _dafny.Map({p1: (d_770_v3_).issubset(d_770_v3_)})
                    d_772_v5_: _dafny.Seq
                    d_772_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "skejjki"))
                    d_773_v6_: _dafny.Map
                    d_773_v6_ = _dafny.Map({d_764_v0_: d_772_v5_})
                    d_774_v7_: _dafny.MultiSet
                    d_774_v7_ = _dafny.MultiSet([p1, len(((d_773_v6_)[not(not((p2).cf1))] if (not(not((p2).cf1))) in (d_773_v6_) else d_772_v5_))])
                    d_775_v8_: _dafny.Array
                    nw147_ = _dafny.Array(int(0), 16)
                    d_775_v8_ = nw147_
                    d_776_v9_: _dafny.Seq
                    d_776_v9_ = _dafny.SeqWithoutIsStrInference([d_775_v8_])
                    index143_ = default__.safeIndex(910, (d_767_v2_).length(0))
                    rhs117_ = p1
                    rhs118_ = ((d_774_v7_)[(p1) * (p1)] if ((p1) * (p1)) in (d_774_v7_) else len(d_776_v9_))
                    rhs119_ = d_764_v0_
                    rhs120_ = d_771_v4_
                    lhs86_ = globalState
                    lhs87_ = globalState
                    lhs88_ = d_767_v2_
                    lhs89_ = default__.safeIndex(910, (d_767_v2_).length(0))
                    lhs86_.f1 = rhs117_
                    lhs87_.f1 = rhs118_
                    lhs88_[lhs89_] = rhs119_
                    d_771_v4_ = rhs120_
                    d_777_v10_: C1
                    nw148_ = C1()
                    nw148_.ctor__((d_767_v2_)[default__.safeIndex(910, (d_767_v2_).length(0))], (d_767_v2_)[default__.safeIndex(910, (d_767_v2_).length(0))], d_770_v3_)
                    d_777_v10_ = nw148_
                    d_778_v11_: _dafny.Map
                    d_778_v11_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eo")): not((d_777_v10_).f14)})
                    d_766_v1_ = (d_766_v1_ if ((d_778_v11_)[_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ei"))] if (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ei"))) in (d_778_v11_) else (d_767_v2_)[default__.safeIndex(910, (d_767_v2_).length(0))]) else d_766_v1_)
                    pass
            pass
        d_779_v12_: _dafny.Seq
        d_779_v12_ = _dafny.SeqWithoutIsStrInference([not(d_764_v0_)])
        d_780_v13_: _dafny.Map
        d_780_v13_ = _dafny.Map({d_764_v0_: d_779_v12_})
        d_781_v14_: _dafny.MultiSet
        d_781_v14_ = _dafny.MultiSet([p1])
        d_782_v15_: _dafny.Seq
        d_782_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ycsjd"))
        d_783_v16_: _dafny.Array
        nw149_ = _dafny.Array(None, 14)
        nw149_[int(0)] = p1
        nw149_[int(1)] = (0) - (len(d_780_v13_))
        nw149_[int(2)] = (d_781_v14_).cardinality
        nw149_[int(3)] = p1
        nw149_[int(4)] = p1
        nw149_[int(5)] = p1
        nw149_[int(6)] = (0) - (p1)
        nw149_[int(7)] = p1
        nw149_[int(8)] = p1
        nw149_[int(9)] = p1
        nw149_[int(10)] = (self).fm9(_dafny.SeqWithoutIsStrInference([p0 for d_784_i2_ in range(default__.abs(256))]), p1, globalState)
        nw149_[int(11)] = len(d_782_v15_)
        nw149_[int(12)] = p1
        nw149_[int(13)] = (p1) * (p1)
        d_783_v16_ = nw149_
        d_785_v17_: _dafny.MultiSet
        d_785_v17_ = _dafny.MultiSet([d_764_v0_])
        d_786_v18_: C1
        nw150_ = C1()
        nw150_.ctor__(d_764_v0_, d_764_v0_, d_785_v17_)
        d_786_v18_ = nw150_
        d_787_v19_: _dafny.Map
        d_787_v19_ = _dafny.Map({d_764_v0_: d_786_v18_})
        d_788_v20_: _dafny.Map
        d_788_v20_ = _dafny.Map({not(False): d_787_v19_})
        index144_ = default__.safeIndex(733, (d_783_v16_).length(0))
        (d_783_v16_)[index144_] = default__.safeModuloInt(len(d_788_v20_), p1)
        index145_ = default__.safeIndex(733, (d_783_v16_).length(0))
        (d_783_v16_)[index145_] = (p1) * (p1)
        d_789_v21_: _dafny.Array
        nw151_ = _dafny.Array(_dafny.Seq({}), 19)
        d_789_v21_ = nw151_
        index146_ = default__.safeIndex(824, (d_789_v21_).length(0))
        (d_789_v21_)[index146_] = d_779_v12_
        index147_ = default__.safeIndex(824, (d_789_v21_).length(0))
        (d_789_v21_)[index147_] = d_779_v12_
        hi8_ = p1
        for d_790_i3_ in range((self).fm9(d_782_v15_, (self).fm9(d_782_v15_, (0) - ((d_783_v16_)[default__.safeIndex(733, (d_783_v16_).length(0))]), globalState), globalState), hi8_):
            d_764_v0_ = not(not((False) or ((d_786_v18_).f15)))
            d_791_v22_: _dafny.Seq
            d_791_v22_ = _dafny.SeqWithoutIsStrInference([d_790_i3_, (d_783_v16_)[default__.safeIndex(733, (d_783_v16_).length(0))], (d_783_v16_)[default__.safeIndex(733, (d_783_v16_).length(0))]])
            d_792_v23_: _dafny.Seq
            d_792_v23_ = _dafny.SeqWithoutIsStrInference([d_791_v22_, d_791_v22_, d_791_v22_])
            index148_ = default__.safeIndex(733, (d_783_v16_).length(0))
            rhs121_ = (d_792_v23_)[default__.safeIndex((0) - (p1), len(d_792_v23_))]
            rhs122_ = default__.safeDivisionInt((len((d_789_v21_)[default__.safeIndex(824, (d_789_v21_).length(0))])) * (p1), (d_783_v16_)[default__.safeIndex(733, (d_783_v16_).length(0))])
            lhs90_ = d_783_v16_
            lhs91_ = default__.safeIndex(733, (d_783_v16_).length(0))
            d_791_v22_ = rhs121_
            lhs90_[lhs91_] = rhs122_
            d_793_v24_: _dafny.Array
            def lambda50_(d_794_i3_, d_795_v16_):
                def lambda51_(d_796_i4_):
                    return (d_794_i3_) == ((d_795_v16_)[default__.safeIndex(733, (d_795_v16_).length(0))])

                return lambda51_

            init28_ = lambda50_(d_790_i3_, d_783_v16_)
            nw152_ = _dafny.Array(None, 4)
            for i0_28_ in range(nw152_.length(0)):
                nw152_[i0_28_] = init28_(i0_28_)
            d_793_v24_ = nw152_
            index149_ = default__.safeIndex(82, (d_793_v24_).length(0))
            (d_793_v24_)[index149_] = (d_786_v18_).f14
            d_797_v25_: _dafny.Set
            d_797_v25_ = _dafny.Set({(0) - ((d_783_v16_)[default__.safeIndex(733, (d_783_v16_).length(0))])})
            index150_ = default__.safeIndex(82, (d_793_v24_).length(0))
            (d_793_v24_)[index150_] = default__.fm4(d_797_v25_, globalState)
            d_798_v26_: _dafny.Map
            d_798_v26_ = _dafny.Map({p0: d_782_v15_})
            d_799_v27_: D6
            d_799_v27_ = D6_DC17((d_783_v16_)[default__.safeIndex(733, (d_783_v16_).length(0))], _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hetkxtwq")), (d_786_v18_).f15)
            d_800_v28_: _dafny.Map
            d_800_v28_ = _dafny.Map({p1: (d_786_v18_).f15})
            d_801_v29_: D6
            d_801_v29_ = D6_DC17((d_783_v16_)[default__.safeIndex(733, (d_783_v16_).length(0))], (d_799_v27_).cf39, ((d_800_v28_)[(d_783_v16_)[default__.safeIndex(733, (d_783_v16_).length(0))]] if ((d_783_v16_)[default__.safeIndex(733, (d_783_v16_).length(0))]) in (d_800_v28_) else d_764_v0_))
            (globalState).f1 = (0) - (default__.safeDivisionInt(d_790_i3_, len((((d_798_v26_)[p0] if (p0) in (d_798_v26_) else d_782_v15_)) + ((d_801_v29_).cf39))))
        _ingredientsd_0 = []
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_783_v16_).length(0)):
            d_802_i5_: int = guard_loop_3_
            if (True) and (((0) <= (d_802_i5_)) and ((d_802_i5_) < ((d_783_v16_).length(0)))):
                _ingredientsd_0.append((d_783_v16_, int(d_802_i5_), default__.safeModuloInt(d_802_i5_, len((_dafny.Map({(d_783_v16_)[default__.safeIndex(733, (d_783_v16_).length(0))]: (d_786_v18_).f15})) | ((D8_DC19(_dafny.Map({p1: (d_786_v18_).f15}))).cf42)))))
        for _tupd_0 in _ingredientsd_0:
            _tupd_0[0][_tupd_0[1]] = _tupd_0[2]
        pat_let_tv18_ = p1
        pat_let_tv19_ = p1
        pat_let_tv20_ = d_781_v14_
        pat_let_tv21_ = d_781_v14_
        def lambda52_(source9_):
            if source9_.is_DC1:
                d_803___mcc_h0_ = source9_.cf1
                d_804___mcc_h1_ = source9_.cf2
                d_805_cf2_ = d_804___mcc_h1_
                d_806_cf1_ = d_803___mcc_h0_
                return _dafny.MultiSet((_dafny.SeqWithoutIsStrInference([pat_let_tv18_])) + (_dafny.SeqWithoutIsStrInference([pat_let_tv19_, len(_dafny.SeqWithoutIsStrInference([-747 for d_807_i6_ in range(default__.abs(176))]))])))
            elif source9_.is_DC0:
                d_808___mcc_h2_ = source9_.cf0
                d_809_cf0_ = d_808___mcc_h2_
                return pat_let_tv20_
            elif True:
                d_810___mcc_h3_ = source9_.cf3
                d_811_cf3_ = d_810___mcc_h3_
                return pat_let_tv21_

        rhs123_ = default__.safeDivisionInt(default__.fm19(p0, d_764_v0_, globalState), default__.safeDivisionInt(p1, p1))
        rhs124_ = (d_783_v16_ if (d_786_v18_).f14 else d_783_v16_)
        rhs125_ = not ((d_786_v18_).f15) or ((d_786_v18_).f14)
        rhs126_ = lambda52_(default__.fm21(globalState))
        lhs92_ = globalState
        lhs92_.f1 = rhs123_
        d_783_v16_ = rhs124_
        d_764_v0_ = rhs125_
        d_781_v14_ = rhs126_

    def m2(self, globalState):
        r0: _dafny.Set = _dafny.Set({})
        r1: int = int(0)
        r2: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        d_812_v0_: bool
        d_812_v0_ = False
        d_813_i0_: int
        d_813_i0_ = 0
        with _dafny.label("5"):
            while d_812_v0_:
                with _dafny.c_label("5"):
                    if (d_813_i0_) >= (100):
                        raise _dafny.Break("5")
                    d_813_i0_ = (d_813_i0_) + (1)
                    d_814_v1_: _dafny.Array
                    nw153_ = _dafny.Array(False, 6)
                    d_814_v1_ = nw153_
                    index151_ = default__.safeIndex(14, (d_814_v1_).length(0))
                    (d_814_v1_)[index151_] = d_812_v0_
                    index152_ = default__.safeIndex(14, (d_814_v1_).length(0))
                    (d_814_v1_)[index152_] = d_812_v0_
                    d_815_v2_: int
                    d_815_v2_ = 809
                    d_816_v3_: _dafny.Array
                    nw154_ = _dafny.Array(None, 4)
                    nw154_[int(0)] = d_815_v2_
                    nw154_[int(1)] = (0) - (d_815_v2_)
                    nw154_[int(2)] = 315
                    nw154_[int(3)] = d_815_v2_
                    d_816_v3_ = nw154_
                    d_817_v4_: _dafny.Array
                    nw155_ = _dafny.Array(None, 2)
                    nw155_[int(0)] = d_816_v3_
                    nw155_[int(1)] = (d_816_v3_ if d_812_v0_ else d_816_v3_)
                    d_817_v4_ = nw155_
                    d_817_v4_ = d_817_v4_
                    d_818_v5_: _dafny.Map
                    d_818_v5_ = _dafny.Map({(d_814_v1_)[default__.safeIndex(14, (d_814_v1_).length(0))]: d_815_v2_})
                    d_819_v6_: _dafny.Set
                    d_819_v6_ = _dafny.Set({d_815_v2_, d_815_v2_, len(((d_818_v5_).set((d_814_v1_)[default__.safeIndex(14, (d_814_v1_).length(0))], 642)).set((d_814_v1_)[default__.safeIndex(14, (d_814_v1_).length(0))], d_815_v2_))})
                    if not(default__.fm4(d_819_v6_, globalState)):
                        d_820_v7_: bool
                        d_821_v8_: _dafny.Array
                        d_822_v9_: int
                        d_823_v10_: int
                        out24_: bool
                        out25_: _dafny.Array
                        out26_: int
                        out27_: int
                        out24_, out25_, out26_, out27_ = (self).m3(globalState)
                        d_820_v7_ = out24_
                        d_821_v8_ = out25_
                        d_822_v9_ = out26_
                        d_823_v10_ = out27_
                        nw156_ = _dafny.Array(int(0), 23)
                        d_816_v3_ = nw156_
                        d_824_v11_: str
                        d_824_v11_ = _dafny.CodePoint('j')
                        d_825_v12_: D3
                        d_825_v12_ = D3_DC8(d_824_v11_, d_812_v0_, default__.fm4(d_819_v6_, globalState), _dafny.SeqWithoutIsStrInference([d_824_v11_ for d_826_i1_ in range(default__.abs(662))]))
                        d_812_v0_ = (d_825_v12_).cf19
                        d_827_v13_: _dafny.Map
                        d_827_v13_ = _dafny.Map({(d_814_v1_)[default__.safeIndex(14, (d_814_v1_).length(0))]: not(d_812_v0_)})
                        d_820_v7_ = not(not (((d_827_v13_)[d_820_v7_] if (d_820_v7_) in (d_827_v13_) else d_820_v7_)) or (d_812_v0_))
                        d_816_v3_ = d_816_v3_
                    elif True:
                        d_828_v14_: _dafny.MultiSet
                        d_828_v14_ = _dafny.MultiSet([(d_814_v1_)[default__.safeIndex(14, (d_814_v1_).length(0))], (d_814_v1_)[default__.safeIndex(14, (d_814_v1_).length(0))], (d_814_v1_)[default__.safeIndex(14, (d_814_v1_).length(0))], d_812_v0_, (d_814_v1_)[default__.safeIndex(14, (d_814_v1_).length(0))]])
                        index153_ = default__.safeIndex(14, (d_814_v1_).length(0))
                        (d_814_v1_)[index153_] = ((d_828_v14_).cardinality) < (d_815_v2_)
                        d_829_v15_: _dafny.Map
                        d_829_v15_ = _dafny.Map({(0) - (d_815_v2_): (d_815_v2_) + (d_815_v2_)})
                        d_830_v16_: _dafny.Seq
                        d_830_v16_ = _dafny.SeqWithoutIsStrInference([d_818_v5_, d_818_v5_, _dafny.Map({(d_814_v1_)[default__.safeIndex(14, (d_814_v1_).length(0))]: 127})])
                        d_831_v17_: _dafny.MultiSet
                        d_831_v17_ = _dafny.MultiSet([d_815_v2_])
                        d_829_v15_ = (d_829_v15_).set((len((d_830_v16_)[default__.safeIndex(d_815_v2_, len(d_830_v16_))])) + (((d_831_v17_)[d_815_v2_] if (d_815_v2_) in (d_831_v17_) else d_815_v2_)), (d_815_v2_) + (d_815_v2_))
                        d_832_v18_: str
                        d_832_v18_ = _dafny.CodePoint('k')
                        d_832_v18_ = _dafny.CodePoint('i')
                        d_833_v19_: _dafny.Seq
                        d_833_v19_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sby"))
                        d_834_v20_: _dafny.Map
                        d_834_v20_ = _dafny.Map({d_817_v4_: d_833_v19_})
                        d_835_v21_: _dafny.Seq
                        d_835_v21_ = _dafny.SeqWithoutIsStrInference([d_833_v19_])
                        d_836_v22_: _dafny.Seq
                        d_836_v22_ = _dafny.SeqWithoutIsStrInference([d_833_v19_, d_833_v19_, _dafny.SeqWithoutIsStrInference([d_832_v18_ for d_837_i5_ in range(default__.abs(-274))]), (d_835_v21_)[default__.safeIndex(d_815_v2_, len(d_835_v21_))], default__.fm2(d_812_v0_, len(d_833_v19_), d_832_v18_, d_815_v2_, globalState)])
                        d_838_v23_: _dafny.Array
                        nw157_ = _dafny.Array(None, 24)
                        nw157_[int(0)] = d_833_v19_
                        nw157_[int(1)] = (_dafny.SeqWithoutIsStrInference([d_832_v18_ for d_839_i2_ in range(default__.abs(757))])) + (d_833_v19_)
                        nw157_[int(2)] = d_833_v19_
                        nw157_[int(3)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "llcxt"))
                        nw157_[int(4)] = d_833_v19_
                        nw157_[int(5)] = ((d_834_v20_)[d_817_v4_] if (d_817_v4_) in (d_834_v20_) else d_833_v19_)
                        nw157_[int(6)] = d_833_v19_
                        nw157_[int(7)] = d_833_v19_
                        nw157_[int(8)] = d_833_v19_
                        nw157_[int(9)] = d_833_v19_
                        nw157_[int(10)] = _dafny.SeqWithoutIsStrInference([d_832_v18_ for d_840_i3_ in range(default__.abs(377))])
                        nw157_[int(11)] = (d_833_v19_).set(default__.safeIndex(981, len(d_833_v19_)), d_832_v18_)
                        nw157_[int(12)] = (d_833_v19_) + (default__.fm2((d_814_v1_)[default__.safeIndex(14, (d_814_v1_).length(0))], d_815_v2_, d_832_v18_, d_815_v2_, globalState))
                        nw157_[int(13)] = d_833_v19_
                        nw157_[int(14)] = (d_833_v19_) + (d_833_v19_)
                        nw157_[int(15)] = _dafny.SeqWithoutIsStrInference([d_832_v18_ for d_841_i4_ in range(default__.abs(914))])
                        nw157_[int(16)] = (d_833_v19_) + (d_833_v19_)
                        nw157_[int(17)] = d_833_v19_
                        nw157_[int(18)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aqevycure"))) + (d_833_v19_)
                        nw157_[int(19)] = default__.fm2(d_812_v0_, d_815_v2_, d_832_v18_, (0) - (len(d_833_v19_)), globalState)
                        nw157_[int(20)] = (d_833_v19_).set(default__.safeIndex(d_815_v2_, len(d_833_v19_)), d_832_v18_)
                        nw157_[int(21)] = (d_835_v21_)[default__.safeIndex(default__.fm19(d_832_v18_, True, globalState), len(d_835_v21_))]
                        nw157_[int(22)] = d_833_v19_
                        nw157_[int(23)] = d_833_v19_
                        d_838_v23_ = nw157_
                        index154_ = default__.safeIndex(24, (d_838_v23_).length(0))
                        (d_838_v23_)[index154_] = (d_833_v19_).set(default__.safeIndex(d_815_v2_, len(d_833_v19_)), d_832_v18_)
                        index155_ = default__.safeIndex(24, (d_838_v23_).length(0))
                        (d_838_v23_)[index155_] = d_833_v19_
                        d_842_v24_: D3
                        d_842_v24_ = D3_DC8(d_832_v18_, (d_814_v1_)[default__.safeIndex(14, (d_814_v1_).length(0))], (d_814_v1_)[default__.safeIndex(14, (d_814_v1_).length(0))], (d_838_v23_)[default__.safeIndex(24, (d_838_v23_).length(0))])
                        d_843_v25_: D3
                        d_843_v25_ = D3_DC9(d_842_v24_)
                        d_843_v25_ = d_843_v25_
                    d_844_v26_: _dafny.Seq
                    d_844_v26_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oxvklsjlo"))
                    d_845_v27_: str
                    d_845_v27_ = _dafny.CodePoint('u')
                    d_846_v28_: _dafny.Map
                    d_846_v28_ = _dafny.Map({d_812_v0_: d_845_v27_})
                    d_847_v29_: _dafny.Map
                    d_847_v29_ = _dafny.Map({(d_814_v1_)[default__.safeIndex(14, (d_814_v1_).length(0))]: (d_814_v1_)[default__.safeIndex(14, (d_814_v1_).length(0))]})
                    d_848_v30_: _dafny.Map
                    d_848_v30_ = _dafny.Map({-195: d_812_v0_})
                    d_849_v31_: _dafny.Map
                    d_849_v31_ = _dafny.Map({d_815_v2_: len(d_848_v30_)})
                    d_850_v32_: _dafny.Array
                    nw158_ = _dafny.Array(None, 19)
                    nw158_[int(0)] = d_844_v26_
                    nw158_[int(1)] = d_844_v26_
                    nw158_[int(2)] = (d_844_v26_) + (d_844_v26_)
                    nw158_[int(3)] = d_844_v26_
                    nw158_[int(4)] = default__.fm2((d_814_v1_)[default__.safeIndex(14, (d_814_v1_).length(0))], d_815_v2_, ((d_846_v28_)[(d_814_v1_)[default__.safeIndex(14, (d_814_v1_).length(0))]] if ((d_814_v1_)[default__.safeIndex(14, (d_814_v1_).length(0))]) in (d_846_v28_) else d_845_v27_), len(d_847_v29_), globalState)
                    nw158_[int(5)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "le"))
                    nw158_[int(6)] = d_844_v26_
                    nw158_[int(7)] = ((d_844_v26_).set(default__.safeIndex(d_815_v2_, len(d_844_v26_)), d_845_v27_)) + (d_844_v26_)
                    nw158_[int(8)] = d_844_v26_
                    nw158_[int(9)] = default__.fm2((d_814_v1_)[default__.safeIndex(14, (d_814_v1_).length(0))], d_815_v2_, _dafny.CodePoint('i'), d_815_v2_, globalState)
                    nw158_[int(10)] = _dafny.SeqWithoutIsStrInference([d_845_v27_ for d_851_i6_ in range(default__.abs(525))])
                    nw158_[int(11)] = d_844_v26_
                    nw158_[int(12)] = d_844_v26_
                    nw158_[int(13)] = (d_844_v26_) + (d_844_v26_)
                    nw158_[int(14)] = _dafny.SeqWithoutIsStrInference([d_845_v27_ for d_852_i7_ in range(default__.abs(579))])
                    nw158_[int(15)] = ((default__.fm23(d_845_v27_, len(d_849_v31_), d_812_v0_, (d_814_v1_)[default__.safeIndex(14, (d_814_v1_).length(0))], globalState)).cf20) + (d_844_v26_)
                    nw158_[int(16)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rkcfltsn"))
                    nw158_[int(17)] = _dafny.SeqWithoutIsStrInference([d_845_v27_ for d_853_i8_ in range(default__.abs(818))])
                    nw158_[int(18)] = (d_844_v26_) + (d_844_v26_)
                    d_850_v32_ = nw158_
                    index156_ = default__.safeIndex(49, (d_850_v32_).length(0))
                    (d_850_v32_)[index156_] = d_844_v26_
                    d_854_v33_: _dafny.Seq
                    d_854_v33_ = _dafny.SeqWithoutIsStrInference([d_844_v26_])
                    index157_ = default__.safeIndex(49, (d_850_v32_).length(0))
                    (d_850_v32_)[index157_] = ((d_854_v33_)[default__.safeIndex(d_815_v2_, len(d_854_v33_))]) + (d_844_v26_)
                    pass
            pass
        d_855_i9_: int
        d_855_i9_ = 0
        with _dafny.label("6"):
            while not(d_812_v0_):
                with _dafny.c_label("6"):
                    if (d_855_i9_) >= (100):
                        raise _dafny.Break("6")
                    d_855_i9_ = (d_855_i9_) + (1)
                    d_856_v34_: int
                    d_856_v34_ = 522
                    d_857_v35_: _dafny.MultiSet
                    d_857_v35_ = _dafny.MultiSet([d_812_v0_])
                    d_858_v36_: C1
                    nw159_ = C1()
                    nw159_.ctor__(d_812_v0_, d_812_v0_, d_857_v35_)
                    d_858_v36_ = nw159_
                    d_859_v37_: _dafny.Map
                    d_859_v37_ = _dafny.Map({d_812_v0_: d_858_v36_})
                    (globalState).f1 = default__.safeDivisionInt(d_856_v34_, (len(d_859_v37_) if d_812_v0_ else d_856_v34_))
                    d_860_v38_: _dafny.MultiSet
                    d_860_v38_ = _dafny.MultiSet([d_856_v34_])
                    d_861_v39_: _dafny.MultiSet
                    d_861_v39_ = d_860_v38_
                    d_862_v40_: _dafny.Map
                    d_862_v40_ = _dafny.Map({d_861_v39_: (d_856_v34_) != ((d_860_v38_).cardinality)})
                    d_863_v41_: _dafny.Set
                    d_863_v41_ = _dafny.Set({d_856_v34_, (d_857_v35_).cardinality})
                    d_864_v42_: _dafny.Map
                    d_864_v42_ = _dafny.Map({d_856_v34_: default__.fm4(d_863_v41_, globalState)})
                    d_865_v43_: _dafny.Map
                    d_865_v43_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_866_i10_ in range(default__.abs(929))]): d_856_v34_})
                    d_862_v40_ = (d_862_v40_).set(d_860_v38_, ((d_864_v42_)[len(d_865_v43_)] if (len(d_865_v43_)) in (d_864_v42_) else d_812_v0_))
                    d_812_v0_ = not(d_812_v0_)
                    d_867_v44_: C0
                    nw160_ = C0()
                    nw160_.ctor__()
                    d_867_v44_ = nw160_
                    pass
            pass
        d_868_v45_: _dafny.Seq
        d_868_v45_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "evwq"))
        d_869_v46_: _dafny.MultiSet
        d_869_v46_ = _dafny.MultiSet([d_868_v45_])
        d_870_v47_: int
        d_870_v47_ = 502
        d_871_v48_: _dafny.Seq
        d_871_v48_ = _dafny.SeqWithoutIsStrInference([(0) - (default__.safeModuloInt((d_869_v46_).cardinality, (0) - (len(d_868_v45_)))), 485, default__.fm19(default__.fm1(globalState), d_812_v0_, globalState), d_870_v47_])
        d_871_v48_ = (d_871_v48_) + (d_871_v48_)
        d_872_i11_: int
        d_872_i11_ = 0
        with _dafny.label("7"):
            while True:
                with _dafny.c_label("7"):
                    if (d_872_i11_) >= (100):
                        raise _dafny.Break("7")
                    d_872_i11_ = (d_872_i11_) + (1)
                    d_873_v49_: _dafny.Map
                    d_873_v49_ = _dafny.Map({d_870_v47_: d_812_v0_})
                    d_874_v50_: _dafny.Seq
                    d_874_v50_ = _dafny.SeqWithoutIsStrInference([d_812_v0_])
                    d_875_v51_: _dafny.MultiSet
                    d_875_v51_ = _dafny.MultiSet([d_812_v0_, d_812_v0_])
                    d_876_v52_: _dafny.Map
                    d_876_v52_ = _dafny.Map({d_812_v0_: d_875_v51_})
                    d_877_v53_: _dafny.Map
                    d_877_v53_ = _dafny.Map({d_812_v0_: default__.fm4(_dafny.Set({-482}), globalState)})
                    d_878_v54_: _dafny.MultiSet
                    d_878_v54_ = _dafny.MultiSet([d_870_v47_, len(_dafny.SeqWithoutIsStrInference([d_812_v0_]))])
                    d_879_v55_: _dafny.Set
                    d_879_v55_ = _dafny.Set({d_870_v47_, d_870_v47_, d_870_v47_, d_870_v47_, (d_878_v54_).cardinality})
                    d_880_v56_: _dafny.Map
                    d_880_v56_ = _dafny.Map({d_877_v53_: _dafny.MultiSet([len(d_879_v55_), d_870_v47_])})
                    d_881_v57_: _dafny.Map
                    d_881_v57_ = _dafny.Map({len(d_880_v56_): len(d_873_v49_)})
                    d_882_v58_: _dafny.Array
                    nw161_ = _dafny.Array(None, 19)
                    nw161_[int(0)] = ((d_873_v49_)[d_870_v47_] if (d_870_v47_) in (d_873_v49_) else d_812_v0_)
                    nw161_[int(1)] = d_812_v0_
                    nw161_[int(2)] = d_812_v0_
                    nw161_[int(3)] = not((d_812_v0_) in (d_874_v50_))
                    nw161_[int(4)] = d_812_v0_
                    nw161_[int(5)] = d_812_v0_
                    nw161_[int(6)] = (default__.fm24(len(d_874_v50_), d_812_v0_, d_873_v49_, d_870_v47_, globalState)).issubset(((d_876_v52_)[d_812_v0_] if (d_812_v0_) in (d_876_v52_) else d_875_v51_))
                    nw161_[int(7)] = False
                    nw161_[int(8)] = d_812_v0_
                    nw161_[int(9)] = d_812_v0_
                    nw161_[int(10)] = True
                    nw161_[int(11)] = d_812_v0_
                    nw161_[int(12)] = (_dafny.Map({863: d_870_v47_})) == (d_881_v57_)
                    nw161_[int(13)] = d_812_v0_
                    nw161_[int(14)] = not (True) or (d_812_v0_)
                    nw161_[int(15)] = d_812_v0_
                    nw161_[int(16)] = (206) < (d_870_v47_)
                    nw161_[int(17)] = d_812_v0_
                    nw161_[int(18)] = d_812_v0_
                    d_882_v58_ = nw161_
                    index158_ = default__.safeIndex(128, (d_882_v58_).length(0))
                    (d_882_v58_)[index158_] = not((d_870_v47_) > (len(d_877_v53_)))
                    index159_ = default__.safeIndex(128, (d_882_v58_).length(0))
                    (d_882_v58_)[index159_] = (d_812_v0_ if d_812_v0_ else d_812_v0_)
                    d_883_v59_: str
                    d_883_v59_ = _dafny.CodePoint('b')
                    d_884_v60_: D3
                    d_884_v60_ = D3_DC8(d_883_v59_, d_812_v0_, d_812_v0_, d_868_v45_)
                    d_885_v61_: _dafny.Array
                    nw162_ = _dafny.Array(None, 21)
                    nw162_[int(0)] = d_868_v45_
                    nw162_[int(1)] = d_868_v45_
                    nw162_[int(2)] = d_868_v45_
                    nw162_[int(3)] = d_868_v45_
                    nw162_[int(4)] = d_868_v45_
                    nw162_[int(5)] = d_868_v45_
                    nw162_[int(6)] = d_868_v45_
                    nw162_[int(7)] = d_868_v45_
                    nw162_[int(8)] = d_868_v45_
                    nw162_[int(9)] = default__.fm2(d_812_v0_, d_870_v47_, d_883_v59_, d_870_v47_, globalState)
                    nw162_[int(10)] = d_868_v45_
                    nw162_[int(11)] = (d_868_v45_).set(default__.safeIndex(d_870_v47_, len(d_868_v45_)), d_883_v59_)
                    nw162_[int(12)] = d_868_v45_
                    nw162_[int(13)] = d_868_v45_
                    nw162_[int(14)] = _dafny.SeqWithoutIsStrInference([d_883_v59_ for d_886_i12_ in range(default__.abs(308))])
                    nw162_[int(15)] = (d_884_v60_).cf20
                    nw162_[int(16)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fbkhxmvx"))
                    nw162_[int(17)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ooxi"))
                    nw162_[int(18)] = d_868_v45_
                    nw162_[int(19)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vvrnb"))
                    nw162_[int(20)] = d_868_v45_
                    d_885_v61_ = nw162_
                    d_887_v62_: _dafny.Map
                    d_887_v62_ = _dafny.Map({d_885_v61_: (d_874_v50_)[default__.safeIndex(d_870_v47_, len(d_874_v50_))]})
                    d_812_v0_ = ((d_887_v62_)[(d_885_v61_ if d_812_v0_ else d_885_v61_)] if ((d_885_v61_ if d_812_v0_ else d_885_v61_)) in (d_887_v62_) else d_812_v0_)
                    d_888_v63_: C2
                    nw163_ = C2()
                    nw163_.ctor__(d_812_v0_, len(d_877_v53_), d_875_v51_)
                    d_888_v63_ = nw163_
                    d_889_v64_: C0
                    nw164_ = C0()
                    nw164_.ctor__()
                    d_889_v64_ = nw164_
                    pass
            pass
        d_890_v65_: _dafny.Array
        nw165_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 17)
        d_890_v65_ = nw165_
        guard_loop_4_: int
        for guard_loop_4_ in _dafny.IntegerRange(0, (d_890_v65_).length(0)):
            d_891_i13_: int = guard_loop_4_
            if (True) and (((0) <= (d_891_i13_)) and ((d_891_i13_) < ((d_890_v65_).length(0)))):
                (d_890_v65_)[(d_891_i13_)] = d_868_v45_
        d_892_v66_: _dafny.MultiSet
        d_892_v66_ = _dafny.MultiSet([d_812_v0_, False, d_812_v0_, d_812_v0_, d_812_v0_])
        d_893_v67_: str
        d_893_v67_ = _dafny.CodePoint('b')
        d_894_v68_: _dafny.Set
        d_894_v68_ = _dafny.Set({d_812_v0_})
        d_895_v69_: _dafny.Array
        nw166_ = _dafny.Array(None, 28)
        nw166_[int(0)] = d_870_v47_
        nw166_[int(1)] = d_870_v47_
        nw166_[int(2)] = d_870_v47_
        nw166_[int(3)] = (d_892_v66_).cardinality
        nw166_[int(4)] = d_870_v47_
        nw166_[int(5)] = d_870_v47_
        nw166_[int(6)] = d_870_v47_
        nw166_[int(7)] = d_870_v47_
        nw166_[int(8)] = ((d_892_v66_)[d_812_v0_] if (d_812_v0_) in (d_892_v66_) else d_870_v47_)
        nw166_[int(9)] = d_870_v47_
        nw166_[int(10)] = default__.fm19(d_893_v67_, d_812_v0_, globalState)
        nw166_[int(11)] = d_870_v47_
        nw166_[int(12)] = d_870_v47_
        nw166_[int(13)] = d_870_v47_
        nw166_[int(14)] = (self).fm9(d_868_v45_, (0) - (d_870_v47_), globalState)
        nw166_[int(15)] = 240
        nw166_[int(16)] = d_870_v47_
        nw166_[int(17)] = (0) - ((0) - ((d_871_v48_)[default__.safeIndex(d_870_v47_, len(d_871_v48_))]))
        nw166_[int(18)] = d_870_v47_
        nw166_[int(19)] = d_870_v47_
        nw166_[int(20)] = d_870_v47_
        nw166_[int(21)] = d_870_v47_
        nw166_[int(22)] = d_870_v47_
        nw166_[int(23)] = d_870_v47_
        nw166_[int(24)] = d_870_v47_
        nw166_[int(25)] = len(d_894_v68_)
        nw166_[int(26)] = d_870_v47_
        nw166_[int(27)] = d_870_v47_
        d_895_v69_ = nw166_
        d_896_v70_: _dafny.Array
        nw167_ = _dafny.Array(_dafny.Map({}), 13)
        d_896_v70_ = nw167_
        source10_ = D5_DC13(d_870_v47_, d_895_v69_, d_896_v70_, (200) - (len(d_868_v45_)))
        if source10_.is_DC13:
            d_897___mcc_h0_ = source10_.cf27
            d_898___mcc_h1_ = source10_.cf28
            d_899___mcc_h2_ = source10_.cf29
            d_900___mcc_h3_ = source10_.cf30
            d_901_cf30_ = d_900___mcc_h3_
            d_902_cf29_ = d_899___mcc_h2_
            d_903_cf28_ = d_898___mcc_h1_
            d_904_cf27_ = d_897___mcc_h0_
            d_905_v71_: D9
            d_905_v71_ = D9_DC21(_dafny.SeqWithoutIsStrInference([d_812_v0_]))
            d_906_v72_: _dafny.Seq
            d_906_v72_ = _dafny.SeqWithoutIsStrInference([d_812_v0_])
            d_907_v73_: bool
            d_908_v74_: _dafny.MultiSet
            d_909_v75_: _dafny.Set
            d_910_v76_: int
            out28_: bool
            out29_: _dafny.MultiSet
            out30_: _dafny.Set
            out31_: int
            out28_, out29_, out30_, out31_ = default__.m0((d_905_v71_).cf48, d_906_v72_, globalState)
            d_907_v73_ = out28_
            d_908_v74_ = out29_
            d_909_v75_ = out30_
            d_910_v76_ = out31_
            d_911_v77_: _dafny.Map
            d_911_v77_ = _dafny.Map({d_904_cf27_: (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "shpmce"))).set(default__.safeIndex(-490, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "shpmce")))), _dafny.CodePoint('q'))})
            d_912_v78_: D9
            d_912_v78_ = D9_DC24(d_870_v47_)
            d_911_v77_ = default__.fm25(d_912_v78_, globalState)
            d_907_v73_ = not((((0) - (len(d_868_v45_))) + ((0) - (d_904_cf27_))) == (d_904_cf27_))
            d_901_cf30_ = d_910_v76_
        elif source10_.is_DC14:
            d_913___mcc_h4_ = source10_.cf31
            d_914___mcc_h5_ = source10_.cf32
            d_915___mcc_h6_ = source10_.cf33
            d_916___mcc_h7_ = source10_.cf34
            d_917___mcc_h8_ = source10_.cf35
            d_918_cf35_ = d_917___mcc_h8_
            d_919_cf34_ = d_916___mcc_h7_
            d_920_cf33_ = d_915___mcc_h6_
            d_921_cf32_ = d_914___mcc_h5_
            d_922_cf31_ = d_913___mcc_h4_
            d_871_v48_ = _dafny.SeqWithoutIsStrInference([(d_870_v47_) + (d_870_v47_), d_870_v47_])
            d_923_v79_: T0
            nw168_ = C1()
            nw168_.ctor__(d_812_v0_, d_812_v0_, _dafny.MultiSet([d_812_v0_]))
            d_923_v79_ = nw168_
            d_924_v80_: _dafny.Array
            nw169_ = _dafny.Array(None, 15)
            nw169_[int(0)] = d_923_v79_
            nw169_[int(1)] = d_923_v79_
            nw169_[int(2)] = d_923_v79_
            nw169_[int(3)] = d_923_v79_
            nw169_[int(4)] = d_923_v79_
            nw169_[int(5)] = d_923_v79_
            nw169_[int(6)] = (d_923_v79_ if not(False) else d_923_v79_)
            nw169_[int(7)] = d_923_v79_
            nw169_[int(8)] = d_923_v79_
            nw169_[int(9)] = d_923_v79_
            nw169_[int(10)] = d_923_v79_
            nw169_[int(11)] = d_923_v79_
            nw169_[int(12)] = d_923_v79_
            nw169_[int(13)] = d_923_v79_
            nw169_[int(14)] = d_923_v79_
            d_924_v80_ = nw169_
            index160_ = default__.safeIndex(798, (d_924_v80_).length(0))
            (d_924_v80_)[index160_] = d_923_v79_
            index161_ = default__.safeIndex(798, (d_924_v80_).length(0))
            nw170_ = C1()
            nw170_.ctor__(d_812_v0_, (d_868_v45_) <= (d_868_v45_), d_892_v66_)
            (d_924_v80_)[index161_] = nw170_
            d_925_v81_: _dafny.Set
            d_925_v81_ = _dafny.Set({d_870_v47_, d_921_cf32_})
            d_926_v82_: _dafny.Map
            d_926_v82_ = _dafny.Map({(0) - (len(d_925_v81_)): d_812_v0_})
            d_927_v83_: _dafny.MultiSet
            d_927_v83_ = _dafny.MultiSet([d_870_v47_, d_920_cf33_, d_920_cf33_, len(d_926_v82_)])
            d_928_v84_: _dafny.Seq
            d_928_v84_ = _dafny.SeqWithoutIsStrInference([(d_927_v83_) - (_dafny.MultiSet([d_870_v47_])), (_dafny.MultiSet([d_920_cf33_])) - (d_927_v83_), (d_927_v83_) | (d_927_v83_), _dafny.MultiSet(d_871_v48_)])
            d_929_v85_: _dafny.Seq
            d_929_v85_ = _dafny.SeqWithoutIsStrInference([d_928_v84_, d_928_v84_, _dafny.SeqWithoutIsStrInference([d_927_v83_ for d_930_i14_ in range(default__.abs(390))]), d_928_v84_, _dafny.SeqWithoutIsStrInference([d_927_v83_])])
            d_931_v86_: _dafny.Map
            d_931_v86_ = _dafny.Map({(d_927_v83_).cardinality: d_920_cf33_})
            d_928_v84_ = ((d_928_v84_) + (d_928_v84_)) + ((d_929_v85_)[default__.safeIndex(len(d_931_v86_), len(d_929_v85_))])
            index162_ = default__.safeIndex(798, (d_895_v69_).length(0))
            (d_895_v69_)[index162_] = default__.fm19(default__.fm1(globalState), d_812_v0_, globalState)
            index163_ = default__.safeIndex(798, (d_895_v69_).length(0))
            (d_895_v69_)[index163_] = default__.fm19(d_893_v67_, not(default__.fm4(d_925_v81_, globalState)), globalState)
        elif source10_.is_DC12:
            d_932___mcc_h9_ = source10_.cf26
            d_933_cf26_ = d_932___mcc_h9_
            d_934_v87_: _dafny.Map
            d_934_v87_ = _dafny.Map({d_870_v47_: d_870_v47_})
            rhs127_ = not(d_812_v0_)
            rhs128_ = d_934_v87_
            rhs129_ = d_893_v67_
            d_812_v0_ = rhs127_
            d_934_v87_ = rhs128_
            d_893_v67_ = rhs129_
            rhs130_ = (0) - (default__.safeModuloInt(d_870_v47_, d_870_v47_))
            rhs131_ = (d_870_v47_) - (len(d_868_v45_))
            rhs132_ = 561
            lhs93_ = globalState
            r1 = rhs130_
            lhs93_.f1 = rhs131_
            r1 = rhs132_
            index164_ = default__.safeIndex(469, (d_890_v65_).length(0))
            (d_890_v65_)[index164_] = (d_868_v45_).set(default__.safeIndex((0) - (d_870_v47_), len(d_868_v45_)), d_893_v67_)
            d_935_v88_: _dafny.Array
            nw171_ = _dafny.Array(False, 29)
            d_935_v88_ = nw171_
            d_936_v89_: _dafny.Seq
            d_936_v89_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_937_i15_ in range(default__.abs(440))])])
            index165_ = default__.safeIndex(469, (d_890_v65_).length(0))
            rhs133_ = d_935_v88_
            rhs134_ = (d_936_v89_)[default__.safeIndex(d_870_v47_, len(d_936_v89_))]
            rhs135_ = (d_868_v45_) + (_dafny.SeqWithoutIsStrInference([d_893_v67_]))
            rhs136_ = d_868_v45_
            rhs137_ = not(not(d_812_v0_))
            lhs94_ = globalState
            lhs95_ = d_890_v65_
            lhs96_ = default__.safeIndex(469, (d_890_v65_).length(0))
            lhs94_.f5 = rhs133_
            lhs95_[lhs96_] = rhs134_
            r2 = rhs135_
            d_868_v45_ = rhs136_
            d_812_v0_ = rhs137_
            d_938_v90_: _dafny.Array
            def lambda53_(d_939_v0_):
                def lambda54_(d_940_i16_):
                    return _dafny.SeqWithoutIsStrInference([d_939_v0_, d_939_v0_, d_939_v0_, not(True)])

                return lambda54_

            init29_ = lambda53_(d_812_v0_)
            nw172_ = _dafny.Array(None, 7)
            for i0_29_ in range(nw172_.length(0)):
                nw172_[i0_29_] = init29_(i0_29_)
            d_938_v90_ = nw172_
            d_941_v91_: _dafny.Seq
            d_941_v91_ = _dafny.SeqWithoutIsStrInference([not(True)])
            index166_ = default__.safeIndex(696, (d_938_v90_).length(0))
            (d_938_v90_)[index166_] = d_941_v91_
            d_942_v92_: _dafny.Array
            nw173_ = _dafny.Array(_dafny.Map({}), 20)
            d_942_v92_ = nw173_
            index167_ = default__.safeIndex(868, (d_942_v92_).length(0))
            (d_942_v92_)[index167_] = (_dafny.Map({d_871_v48_: d_812_v0_})).set(default__.fm20(d_870_v47_, d_870_v47_, globalState), d_812_v0_)
            d_943_v93_: _dafny.Set
            d_943_v93_ = _dafny.Set({d_870_v47_, d_870_v47_})
            d_944_v94_: _dafny.MultiSet
            d_944_v94_ = _dafny.MultiSet([len(_dafny.Set({d_812_v0_})), d_870_v47_, d_870_v47_, (self).fm9((d_890_v65_)[default__.safeIndex(469, (d_890_v65_).length(0))], 169, globalState)])
            index168_ = default__.safeIndex(696, (d_938_v90_).length(0))
            index169_ = default__.safeIndex(868, (d_942_v92_).length(0))
            rhs138_ = d_941_v91_
            rhs139_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([len(d_943_v93_)]): ((d_944_v94_).cardinality) <= (424)})
            rhs140_ = not(d_812_v0_)
            lhs97_ = d_938_v90_
            lhs98_ = default__.safeIndex(696, (d_938_v90_).length(0))
            lhs99_ = d_942_v92_
            lhs100_ = default__.safeIndex(868, (d_942_v92_).length(0))
            lhs97_[lhs98_] = rhs138_
            lhs99_[lhs100_] = rhs139_
            d_812_v0_ = rhs140_
        elif True:
            d_945___mcc_h10_ = source10_.cf36
            d_946_cf36_ = d_945___mcc_h10_
            d_947_v95_: C2
            nw174_ = C2()
            nw174_.ctor__(d_812_v0_, d_870_v47_, _dafny.MultiSet([False]))
            d_947_v95_ = nw174_
            d_947_v95_ = d_947_v95_
            index170_ = default__.safeIndex(149, (d_895_v69_).length(0))
            (d_895_v69_)[index170_] = default__.safeDivisionInt(d_870_v47_, d_870_v47_)
            index171_ = default__.safeIndex(149, (d_895_v69_).length(0))
            rhs141_ = d_870_v47_
            rhs142_ = (0) - ((d_947_v95_).f13)
            rhs143_ = (not(False)) == ((d_870_v47_) > (-387))
            rhs144_ = not(d_812_v0_)
            lhs101_ = d_895_v69_
            lhs102_ = default__.safeIndex(149, (d_895_v69_).length(0))
            lhs103_ = d_947_v95_
            lhs101_[lhs102_] = rhs141_
            r1 = rhs142_
            d_812_v0_ = rhs143_
            lhs103_.f12 = rhs144_
            d_948_v96_: _dafny.Seq
            d_948_v96_ = _dafny.SeqWithoutIsStrInference([d_812_v0_])
            d_949_v97_: _dafny.Set
            d_949_v97_ = _dafny.Set({len((d_948_v96_).set(default__.safeIndex((d_895_v69_)[default__.safeIndex(149, (d_895_v69_).length(0))], len(d_948_v96_)), d_947_v95_.f12)), (d_895_v69_)[default__.safeIndex(149, (d_895_v69_).length(0))]})
            d_950_v98_: _dafny.Array
            nw175_ = _dafny.Array(None, 11)
            nw175_[int(0)] = d_812_v0_
            nw175_[int(1)] = d_947_v95_.f12
            nw175_[int(2)] = default__.fm4(d_949_v97_, globalState)
            nw175_[int(3)] = not((d_947_v95_.f12) and (True))
            nw175_[int(4)] = d_947_v95_.f12
            nw175_[int(5)] = d_947_v95_.f12
            nw175_[int(6)] = (d_947_v95_.f12 if d_947_v95_.f12 else True)
            nw175_[int(7)] = not(d_812_v0_)
            nw175_[int(8)] = d_947_v95_.f12
            nw175_[int(9)] = not(((d_948_v96_)[default__.safeIndex((d_947_v95_).f13, len(d_948_v96_))] if d_947_v95_.f12 else d_947_v95_.f12))
            nw175_[int(10)] = ((d_895_v69_)[default__.safeIndex(149, (d_895_v69_).length(0))]) <= (575)
            d_950_v98_ = nw175_
            index172_ = default__.safeIndex(254, (d_950_v98_).length(0))
            (d_950_v98_)[index172_] = d_812_v0_
            index173_ = default__.safeIndex(254, (d_950_v98_).length(0))
            (d_950_v98_)[index173_] = ((d_894_v68_) | (d_894_v68_)).isdisjoint(d_894_v68_)
            (d_947_v95_).f12 = False
        d_951_v99_: _dafny.Seq
        d_951_v99_ = _dafny.SeqWithoutIsStrInference([d_812_v0_, d_812_v0_])
        d_952_v100_: _dafny.Map
        d_952_v100_ = _dafny.Map({d_870_v47_: d_870_v47_})
        d_953_v101_: _dafny.Set
        d_953_v101_ = _dafny.Set({d_870_v47_, len(d_871_v48_), d_870_v47_})
        d_954_v102_: _dafny.Set
        d_954_v102_ = _dafny.Set({d_870_v47_, default__.fm19(d_893_v67_, (d_951_v99_)[default__.safeIndex(d_870_v47_, len(d_951_v99_))], globalState), (d_870_v47_) - (d_870_v47_), ((d_952_v100_)[(self).fm9(_dafny.SeqWithoutIsStrInference([d_893_v67_ for d_955_i17_ in range(default__.abs(433))]), (D6_DC17(d_870_v47_, d_868_v45_, default__.fm4(d_953_v101_, globalState))).cf38, globalState)] if ((self).fm9(_dafny.SeqWithoutIsStrInference([d_893_v67_ for d_956_i17_ in range(default__.abs(433))]), (D6_DC17(d_870_v47_, d_868_v45_, default__.fm4(d_953_v101_, globalState))).cf38, globalState)) in (d_952_v100_) else 153)})
        r0 = d_954_v102_
        r1 = d_870_v47_
        r2 = (d_868_v45_) + (default__.fm2(d_812_v0_, d_870_v47_, d_893_v67_, len(_dafny.Map({d_870_v47_: d_812_v0_})), globalState))
        return r0, r1, r2

    def m3(self, globalState):
        r0: bool = False
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: int = int(0)
        r3: int = int(0)
        if False:
            d_957_v0_: _dafny.Seq
            d_957_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dtaw"))
            d_958_v1_: _dafny.Set
            d_958_v1_ = _dafny.Set({d_957_v0_})
            d_959_v2_: _dafny.MultiSet
            d_959_v2_ = _dafny.MultiSet([d_957_v0_])
            d_960_v3_: _dafny.Seq
            d_960_v3_ = _dafny.SeqWithoutIsStrInference([d_959_v2_])
            d_961_v4_: int
            d_961_v4_ = 794
            def iife65_():
                coll37_ = _dafny.Set()
                compr_37_: _dafny.Seq
                for compr_37_ in ((d_960_v3_)[default__.safeIndex(d_961_v4_, len(d_960_v3_))]).Elements:
                    d_962_v5_: _dafny.Seq = compr_37_
                    if (d_962_v5_) in ((d_960_v3_)[default__.safeIndex(d_961_v4_, len(d_960_v3_))]):
                        coll37_ = coll37_.union(_dafny.Set([d_962_v5_]))
                return _dafny.Set(coll37_)
            d_958_v1_ = iife65_()
            
            d_963_v6_: _dafny.Map
            d_963_v6_ = _dafny.Map({len(d_957_v0_): d_961_v4_})
            d_963_v6_ = (d_963_v6_).set((0) - ((d_961_v4_) + (d_961_v4_)), (0) - (d_961_v4_))
            d_964_v7_: bool
            d_964_v7_ = True
            d_965_v8_: _dafny.Seq
            d_965_v8_ = _dafny.SeqWithoutIsStrInference([True, d_964_v7_])
            d_966_v9_: str
            d_966_v9_ = _dafny.CodePoint('s')
            (globalState).f1 = len(((((d_957_v0_) + (d_957_v0_)).set(default__.safeIndex(len(d_965_v8_), len((d_957_v0_) + (d_957_v0_))), _dafny.CodePoint('t')) if ((self).fm9(d_957_v0_, d_961_v4_, globalState)) > (d_961_v4_) else d_957_v0_)).set(default__.safeIndex(d_961_v4_, len((((d_957_v0_) + (d_957_v0_)).set(default__.safeIndex(len(d_965_v8_), len((d_957_v0_) + (d_957_v0_))), _dafny.CodePoint('t')) if ((self).fm9(d_957_v0_, d_961_v4_, globalState)) > (d_961_v4_) else d_957_v0_))), d_966_v9_))
            d_965_v8_ = d_965_v8_
            d_967_v10_: _dafny.Map
            d_967_v10_ = _dafny.Map({d_964_v7_: d_966_v9_})
            d_967_v10_ = (d_967_v10_).set(d_964_v7_, d_966_v9_)
        elif True:
            d_968_v11_: int
            d_968_v11_ = -824
            (globalState).f1 = (d_968_v11_) * (-559)
            d_969_v12_: str
            d_969_v12_ = _dafny.CodePoint('u')
            d_970_v13_: bool
            d_970_v13_ = True
            d_971_v14_: _dafny.MultiSet
            d_971_v14_ = _dafny.MultiSet([d_970_v13_, not(True)])
            d_972_v15_: _dafny.Set
            d_972_v15_ = _dafny.Set({(d_971_v14_).cardinality, d_968_v11_})
            d_973_v16_: _dafny.Map
            d_973_v16_ = _dafny.Map({d_970_v13_: d_968_v11_})
            d_974_v17_: D8
            d_974_v17_ = D8_DC20(d_969_v12_, default__.fm4(d_972_v15_, globalState), d_970_v13_, d_973_v16_, d_970_v13_)
            d_975_v18_: _dafny.Map
            d_975_v18_ = _dafny.Map({(d_974_v17_).cf43: d_969_v12_})
            d_976_v19_: _dafny.Seq
            d_976_v19_ = _dafny.SeqWithoutIsStrInference([d_969_v12_, _dafny.CodePoint('g'), d_969_v12_])
            d_975_v18_ = (d_975_v18_).set(d_969_v12_, (d_976_v19_)[default__.safeIndex(d_968_v11_, len(d_976_v19_))])
            d_977_v20_: _dafny.Array
            def lambda55_(d_978_v12_):
                def lambda56_(d_979_i0_):
                    return d_978_v12_

                return lambda56_

            init30_ = lambda55_(d_969_v12_)
            nw176_ = _dafny.Array(None, 16)
            for i0_30_ in range(nw176_.length(0)):
                nw176_[i0_30_] = init30_(i0_30_)
            d_977_v20_ = nw176_
            d_977_v20_ = d_977_v20_
            r0 = not (d_970_v13_) or (d_970_v13_)
            d_980_v21_: _dafny.Map
            d_980_v21_ = _dafny.Map({d_968_v11_: False})
            rhs145_ = d_968_v11_
            rhs146_ = d_970_v13_
            rhs147_ = default__.fm4(d_972_v15_, globalState)
            rhs148_ = (len(d_980_v21_)) * ((d_968_v11_) + (d_968_v11_))
            rhs149_ = default__.safeDivisionInt(d_968_v11_, d_968_v11_)
            lhs104_ = globalState
            lhs104_.f1 = rhs145_
            d_970_v13_ = rhs146_
            d_970_v13_ = rhs147_
            r3 = rhs148_
            r3 = rhs149_
        d_981_v22_: bool
        d_981_v22_ = True
        d_982_v23_: _dafny.MultiSet
        d_982_v23_ = _dafny.MultiSet([d_981_v22_])
        d_983_v24_: _dafny.MultiSet
        d_983_v24_ = d_982_v23_
        source11_ = d_983_v24_
        d_984___mcc_h0_ = source11_
        d_985_cf41_ = d_984___mcc_h0_
        d_986_v25_: _dafny.Seq
        d_986_v25_ = _dafny.SeqWithoutIsStrInference([False, d_981_v22_])
        d_987_v26_: bool
        d_988_v27_: _dafny.MultiSet
        d_989_v28_: _dafny.Set
        d_990_v29_: int
        out32_: bool
        out33_: _dafny.MultiSet
        out34_: _dafny.Set
        out35_: int
        out32_, out33_, out34_, out35_ = default__.m0((d_986_v25_) + (d_986_v25_), d_986_v25_, globalState)
        d_987_v26_ = out32_
        d_988_v27_ = out33_
        d_989_v28_ = out34_
        d_990_v29_ = out35_
        d_991_v30_: C0
        nw177_ = C0()
        nw177_.ctor__()
        d_991_v30_ = nw177_
        d_990_v29_ = d_990_v29_
        d_992_v31_: str
        d_992_v31_ = _dafny.CodePoint('t')
        (globalState).f1 = (0) - (default__.fm19(d_992_v31_, d_987_v26_, globalState))
        d_993_v32_: int
        d_993_v32_ = -895
        d_994_v33_: _dafny.Map
        d_994_v33_ = _dafny.Map({d_993_v32_: d_981_v22_})
        d_995_v34_: _dafny.Array
        nw178_ = _dafny.Array(None, 5)
        nw178_[int(0)] = not(d_981_v22_)
        nw178_[int(1)] = ((d_994_v33_)[d_993_v32_] if (d_993_v32_) in (d_994_v33_) else d_981_v22_)
        nw178_[int(2)] = (len(_dafny.SeqWithoutIsStrInference([d_993_v32_]))) > (d_993_v32_)
        nw178_[int(3)] = d_981_v22_
        nw178_[int(4)] = d_981_v22_
        d_995_v34_ = nw178_
        d_996_v35_: _dafny.Seq
        d_996_v35_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l"))
        d_997_v36_: _dafny.MultiSet
        d_997_v36_ = _dafny.MultiSet([d_996_v35_])
        index174_ = default__.safeIndex(408, (d_995_v34_).length(0))
        (d_995_v34_)[index174_] = (d_997_v36_).isdisjoint(default__.fm26(globalState))
        index175_ = default__.safeIndex(408, (d_995_v34_).length(0))
        (d_995_v34_)[index175_] = d_981_v22_
        d_998_v37_: _dafny.Set
        d_998_v37_ = _dafny.Set({d_993_v32_, (0) - (d_993_v32_)})
        d_999_v38_: _dafny.Seq
        d_999_v38_ = _dafny.SeqWithoutIsStrInference([default__.fm4(d_998_v37_, globalState)])
        d_1000_v39_: D9
        d_1000_v39_ = D9_DC21(d_999_v38_)
        pat_let_tv22_ = d_999_v38_
        def iife66_(_pat_let14_0):
            def iife67_(d_1001_dt__update__tmp_h0_):
                def iife68_(_pat_let15_0):
                    def iife69_(d_1002_dt__update_hcf48_h0_):
                        return D9_DC21(d_1002_dt__update_hcf48_h0_)
                    return iife69_(_pat_let15_0)
                return iife68_(pat_let_tv22_)
            return iife67_(_pat_let14_0)
        source12_ = iife66_(d_1000_v39_)
        if source12_.is_DC22:
            d_1003___mcc_h1_ = source12_.cf49
            d_1004___mcc_h2_ = source12_.cf50
            d_1005___mcc_h3_ = source12_.cf51
            d_1006_cf51_ = d_1005___mcc_h3_
            d_1007_cf50_ = d_1004___mcc_h2_
            d_1008_cf49_ = d_1003___mcc_h1_
            d_1009_v40_: C2
            nw179_ = C2()
            nw179_.ctor__(False, (d_1008_cf49_) + (506), (d_982_v23_) - (d_982_v23_))
            d_1009_v40_ = nw179_
            d_1010_v41_: _dafny.Array
            nw180_ = _dafny.Array(_dafny.Array(None, 0), 12)
            d_1010_v41_ = nw180_
            d_1011_v42_: _dafny.Map
            d_1011_v42_ = _dafny.Map({d_1008_cf49_: ((d_982_v23_)[(d_995_v34_)[default__.safeIndex(408, (d_995_v34_).length(0))]] if ((d_995_v34_)[default__.safeIndex(408, (d_995_v34_).length(0))]) in (d_982_v23_) else d_993_v32_)})
            d_1012_v43_: str
            d_1012_v43_ = _dafny.CodePoint('v')
            d_1013_v44_: _dafny.Set
            d_1013_v44_ = _dafny.Set({d_995_v34_})
            d_1014_v45_: _dafny.MultiSet
            d_1014_v45_ = _dafny.MultiSet([(d_1009_v40_).f13])
            d_1015_v46_: _dafny.Map
            d_1015_v46_ = _dafny.Map({(d_999_v38_)[default__.safeIndex(d_993_v32_, len(d_999_v38_))]: (d_995_v34_)[default__.safeIndex(408, (d_995_v34_).length(0))]})
            d_1016_v47_: _dafny.Map
            d_1016_v47_ = _dafny.Map({d_996_v35_: default__.fm27(d_1012_v43_, -714, d_1009_v40_.f12, default__.fm24(d_1008_cf49_, (d_995_v34_)[default__.safeIndex(408, (d_995_v34_).length(0))], d_994_v33_, d_1008_cf49_, globalState), globalState)})
            d_1017_v48_: _dafny.Array
            nw181_ = _dafny.Array(None, 16)
            nw181_[int(0)] = len(_dafny.SeqWithoutIsStrInference([d_1007_cf50_, True, d_981_v22_]))
            nw181_[int(1)] = ((d_1011_v42_)[len(_dafny.SeqWithoutIsStrInference([d_1012_v43_]))] if (len(_dafny.SeqWithoutIsStrInference([d_1012_v43_]))) in (d_1011_v42_) else len(d_1013_v44_))
            nw181_[int(2)] = (0) - ((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vtmvqyk")))) - (d_993_v32_))
            nw181_[int(3)] = (len(d_996_v35_)) + (len(d_996_v35_))
            nw181_[int(4)] = d_1008_cf49_
            nw181_[int(5)] = 720
            nw181_[int(6)] = default__.safeDivisionInt(((d_1014_v45_)[(d_1009_v40_).f13] if ((d_1009_v40_).f13) in (d_1014_v45_) else len(d_1015_v46_)), d_1008_cf49_)
            nw181_[int(7)] = -4
            nw181_[int(8)] = default__.safeModuloInt(d_1008_cf49_, d_1008_cf49_)
            nw181_[int(9)] = (d_1009_v40_).f13
            nw181_[int(10)] = d_993_v32_
            nw181_[int(11)] = default__.safeModuloInt(-734, (0) - ((d_1009_v40_).fm11(d_1014_v45_, globalState)))
            nw181_[int(12)] = (d_1009_v40_).f13
            nw181_[int(13)] = len(d_1016_v47_)
            nw181_[int(14)] = (d_993_v32_) * (len(d_996_v35_))
            nw181_[int(15)] = len(d_999_v38_)
            d_1017_v48_ = nw181_
            index176_ = default__.safeIndex(306, (d_1017_v48_).length(0))
            (d_1017_v48_)[index176_] = (0) - ((d_1009_v40_).f13)
            d_1018_v49_: _dafny.Map
            d_1018_v49_ = _dafny.Map({d_981_v22_: default__.fm28((d_1009_v40_).fm11(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_1008_cf49_])), globalState), globalState)})
            d_1019_v50_: _dafny.Map
            d_1019_v50_ = _dafny.Map({False: d_993_v32_})
            d_1020_v51_: _dafny.Map
            d_1020_v51_ = _dafny.Map({((d_1018_v49_)[True] if (True) in (d_1018_v49_) else d_1019_v50_): d_1010_v41_})
            d_1021_v52_: _dafny.Seq
            d_1021_v52_ = _dafny.SeqWithoutIsStrInference([(d_1019_v50_ if d_1007_cf50_ else _dafny.Map({False: d_1008_cf49_})), d_1019_v50_])
            index177_ = default__.safeIndex(306, (d_1017_v48_).length(0))
            rhs150_ = ((d_1020_v51_)[d_1019_v50_] if (d_1019_v50_) in (d_1020_v51_) else d_1010_v41_)
            rhs151_ = len(d_1021_v52_)
            lhs105_ = d_1017_v48_
            lhs106_ = default__.safeIndex(306, (d_1017_v48_).length(0))
            d_1010_v41_ = rhs150_
            lhs105_[lhs106_] = rhs151_
            d_981_v22_ = d_981_v22_
            d_1022_v53_: C0
            nw182_ = C0()
            nw182_.ctor__()
            d_1022_v53_ = nw182_
            d_1023_v54_: _dafny.Map
            d_1023_v54_ = _dafny.Map({d_1022_v53_: False})
            d_1024_v55_: _dafny.Map
            d_1024_v55_ = _dafny.Map({(d_999_v38_) + ((d_999_v38_).set(default__.safeIndex(len(d_1023_v54_), len(d_999_v38_)), d_981_v22_)): (0) - ((d_1017_v48_)[default__.safeIndex(306, (d_1017_v48_).length(0))])})
            d_1024_v55_ = (d_1024_v55_).set(d_999_v38_, d_993_v32_)
        elif source12_.is_DC23:
            d_1025___mcc_h4_ = source12_.cf52
            d_1026_cf52_ = d_1025___mcc_h4_
            r3 = d_1026_cf52_
            d_1027_v56_: _dafny.Seq
            d_1027_v56_ = _dafny.SeqWithoutIsStrInference([((d_1000_v39_).cf48) + (_dafny.SeqWithoutIsStrInference([(d_995_v34_)[default__.safeIndex(408, (d_995_v34_).length(0))]]))])
            rhs152_ = d_981_v22_
            rhs153_ = d_1027_v56_
            r0 = rhs152_
            d_1027_v56_ = rhs153_
            d_1028_v57_: _dafny.MultiSet
            d_1028_v57_ = _dafny.MultiSet([d_993_v32_, d_993_v32_, d_993_v32_])
            d_1029_v58_: _dafny.Seq
            d_1029_v58_ = _dafny.SeqWithoutIsStrInference([d_1026_cf52_])
            if (_dafny.MultiSet(d_1029_v58_)).ispropersubset((_dafny.MultiSet(default__.fm20(d_1026_cf52_, d_1026_cf52_, globalState))).intersection(d_1028_v57_)):
                d_1030_v59_: str
                d_1030_v59_ = _dafny.CodePoint('n')
                d_1031_v60_: D3
                d_1031_v60_ = D3_DC8(d_1030_v59_, (d_995_v34_)[default__.safeIndex(408, (d_995_v34_).length(0))], d_981_v22_, d_996_v35_)
                pat_let_tv23_ = d_995_v34_
                pat_let_tv24_ = d_995_v34_
                pat_let_tv25_ = d_1030_v59_
                def iife70_(_pat_let16_0):
                    def iife71_(d_1032_dt__update__tmp_h1_):
                        def iife72_(_pat_let17_0):
                            def iife73_(d_1033_dt__update_hcf18_h0_):
                                def iife74_(_pat_let18_0):
                                    def iife75_(d_1034_dt__update_hcf17_h0_):
                                        return D3_DC8(d_1034_dt__update_hcf17_h0_, d_1033_dt__update_hcf18_h0_, (d_1032_dt__update__tmp_h1_).cf19, (d_1032_dt__update__tmp_h1_).cf20)
                                    return iife75_(_pat_let18_0)
                                return iife74_(pat_let_tv25_)
                            return iife73_(_pat_let17_0)
                        return iife72_((pat_let_tv24_)[default__.safeIndex(408, (pat_let_tv23_).length(0))])
                    return iife71_(_pat_let16_0)
                d_1031_v60_ = iife70_(d_1031_v60_)
                d_1035_v61_: C0
                nw183_ = C0()
                nw183_.ctor__()
                d_1035_v61_ = nw183_
                d_1036_v62_: _dafny.Map
                d_1036_v62_ = _dafny.Map({(d_995_v34_)[default__.safeIndex(408, (d_995_v34_).length(0))]: d_1026_cf52_})
                d_1037_v63_: _dafny.Map
                d_1037_v63_ = _dafny.Map({-903: d_995_v34_})
                d_1038_v64_: _dafny.Array
                nw184_ = _dafny.Array(None, 6)
                nw184_[int(0)] = (len(d_1036_v62_)) + (d_993_v32_)
                nw184_[int(1)] = len(d_1037_v63_)
                nw184_[int(2)] = d_993_v32_
                nw184_[int(3)] = ((d_1028_v57_)[d_993_v32_] if (d_993_v32_) in (d_1028_v57_) else d_1026_cf52_)
                nw184_[int(4)] = (len(d_998_v37_)) * (d_1026_cf52_)
                nw184_[int(5)] = (0) - (d_993_v32_)
                d_1038_v64_ = nw184_
                index178_ = default__.safeIndex(934, (d_1038_v64_).length(0))
                (d_1038_v64_)[index178_] = (d_993_v32_ if (d_995_v34_)[default__.safeIndex(408, (d_995_v34_).length(0))] else (0) - ((d_982_v23_).cardinality))
                index179_ = default__.safeIndex(934, (d_1038_v64_).length(0))
                rhs154_ = _dafny.SeqWithoutIsStrInference([d_1030_v59_ for d_1039_i1_ in range(default__.abs(956))])
                rhs155_ = d_993_v32_
                lhs107_ = d_1038_v64_
                lhs108_ = default__.safeIndex(934, (d_1038_v64_).length(0))
                d_996_v35_ = rhs154_
                lhs107_[lhs108_] = rhs155_
                d_1040_v65_: C2
                nw185_ = C2()
                nw185_.ctor__((d_993_v32_) >= ((d_1038_v64_)[default__.safeIndex(934, (d_1038_v64_).length(0))]), 373, (d_983_v24_))
                d_1040_v65_ = nw185_
                d_1041_v66_: _dafny.Map
                d_1041_v66_ = _dafny.Map({d_1040_v65_.f12: (d_995_v34_)[default__.safeIndex(408, (d_995_v34_).length(0))]})
                d_981_v22_ = (d_981_v22_) == (((d_1041_v66_)[True] if (True) in (d_1041_v66_) else d_1040_v65_.f12))
            elif True:
                (globalState).f1 = ((0) - (d_993_v32_)) * (d_993_v32_)
                d_1042_v67_: _dafny.Set
                d_1042_v67_ = _dafny.Set({d_996_v35_})
                r3 = (0) - ((d_993_v32_) + (((d_1029_v58_)[default__.safeIndex(len(d_1042_v67_), len(d_1029_v58_))]) * (d_1026_cf52_)))
                def lambda57_(d_1043_v34_):
                    def lambda58_(d_1044_i2_):
                        return (d_1043_v34_)[default__.safeIndex(408, (d_1043_v34_).length(0))]

                    return lambda58_

                init31_ = lambda57_(d_995_v34_)
                nw186_ = _dafny.Array(None, 1)
                for i0_31_ in range(nw186_.length(0)):
                    nw186_[i0_31_] = init31_(i0_31_)
                d_995_v34_ = nw186_
                d_1045_v68_: str
                d_1045_v68_ = _dafny.CodePoint('a')
                r3 = default__.fm19(d_1045_v68_, d_981_v22_, globalState)
                index180_ = default__.safeIndex(408, (d_995_v34_).length(0))
                index181_ = default__.safeIndex(408, (d_995_v34_).length(0))
                rhs156_ = (d_1026_cf52_) + (d_993_v32_)
                rhs157_ = (d_995_v34_)[default__.safeIndex(408, (d_995_v34_).length(0))]
                rhs158_ = d_1026_cf52_
                rhs159_ = (d_995_v34_)[default__.safeIndex(408, (d_995_v34_).length(0))]
                rhs160_ = (d_995_v34_)[default__.safeIndex(408, (d_995_v34_).length(0))]
                lhs109_ = d_995_v34_
                lhs110_ = default__.safeIndex(408, (d_995_v34_).length(0))
                lhs111_ = d_995_v34_
                lhs112_ = default__.safeIndex(408, (d_995_v34_).length(0))
                d_993_v32_ = rhs156_
                lhs109_[lhs110_] = rhs157_
                r2 = rhs158_
                lhs111_[lhs112_] = rhs159_
                d_981_v22_ = rhs160_
            d_1046_v69_: _dafny.Array
            def lambda59_(d_1047_v58_):
                def lambda60_(d_1048_i3_):
                    return d_1047_v58_

                return lambda60_

            init32_ = lambda59_(d_1029_v58_)
            nw187_ = _dafny.Array(None, 8)
            for i0_32_ in range(nw187_.length(0)):
                nw187_[i0_32_] = init32_(i0_32_)
            d_1046_v69_ = nw187_
            d_1049_v70_: D6
            d_1049_v70_ = D6_DC17(276, d_996_v35_, d_981_v22_)
            d_1050_v71_: _dafny.Map
            d_1050_v71_ = _dafny.Map({d_999_v38_: d_981_v22_})
            d_1051_v72_: _dafny.Map
            d_1051_v72_ = _dafny.Map({d_1026_cf52_: d_1050_v71_})
            rhs161_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "csulwnn"))
            rhs162_ = (d_1049_v70_).cf39
            rhs163_ = d_1046_v69_
            rhs164_ = (len(d_1029_v58_)) > (len(d_1029_v58_))
            rhs165_ = ((_dafny.SeqWithoutIsStrInference([d_981_v22_, d_981_v22_, True, (d_995_v34_)[default__.safeIndex(408, (d_995_v34_).length(0))], False]) if (d_995_v34_)[default__.safeIndex(408, (d_995_v34_).length(0))] else d_999_v38_)) not in ((((d_1051_v72_)[d_993_v32_] if (d_993_v32_) in (d_1051_v72_) else d_1050_v71_)) | (_dafny.Map({d_999_v38_: (d_995_v34_)[default__.safeIndex(408, (d_995_v34_).length(0))]})))
            d_996_v35_ = rhs161_
            d_996_v35_ = rhs162_
            d_1046_v69_ = rhs163_
            r0 = rhs164_
            r0 = rhs165_
        elif source12_.is_DC24:
            d_1052___mcc_h5_ = source12_.cf53
            d_1053_cf53_ = d_1052___mcc_h5_
            d_1054_v73_: C1
            nw188_ = C1()
            nw188_.ctor__(default__.fm4(_dafny.Set({d_993_v32_, len(d_999_v38_)}), globalState), d_981_v22_, (d_982_v23_).intersection(_dafny.MultiSet(d_999_v38_)))
            d_1054_v73_ = nw188_
            index182_ = default__.safeIndex(408, (d_995_v34_).length(0))
            (d_995_v34_)[index182_] = (d_995_v34_)[default__.safeIndex(408, (d_995_v34_).length(0))]
            r0 = (d_1054_v73_).f15
            d_1055_v74_: _dafny.Seq
            d_1055_v74_ = _dafny.SeqWithoutIsStrInference([(d_998_v37_).intersection(d_998_v37_), d_998_v37_])
            d_998_v37_ = (d_1055_v74_)[default__.safeIndex(len(d_999_v38_), len(d_1055_v74_))]
        elif True:
            d_1056___mcc_h6_ = source12_.cf48
            d_1057_cf48_ = d_1056___mcc_h6_
            r0 = d_981_v22_
            d_996_v35_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rd"))
            d_1058_v75_: str
            d_1058_v75_ = _dafny.CodePoint('o')
            d_981_v22_ = not((d_1058_v75_) not in ((((d_996_v35_).set(default__.safeIndex(d_993_v32_, len(d_996_v35_)), d_1058_v75_)) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_1059_i4_ in range(default__.abs(-326))]))).set(default__.safeIndex(d_993_v32_, len(((d_996_v35_).set(default__.safeIndex(d_993_v32_, len(d_996_v35_)), d_1058_v75_)) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_1060_i4_ in range(default__.abs(-326))])))), d_1058_v75_)))
            d_1061_v76_: _dafny.Map
            d_1061_v76_ = _dafny.Map({(d_995_v34_)[default__.safeIndex(408, (d_995_v34_).length(0))]: default__.fm4(d_998_v37_, globalState)})
            d_1062_v77_: _dafny.Map
            d_1062_v77_ = _dafny.Map({d_993_v32_: d_1061_v76_})
            d_1063_v78_: _dafny.Set
            d_1063_v78_ = _dafny.Set({(d_995_v34_)[default__.safeIndex(408, (d_995_v34_).length(0))]})
            d_1061_v76_ = ((d_1062_v77_)[default__.safeModuloInt((self).fm9(d_996_v35_, len(d_1063_v78_), globalState), 990)] if (default__.safeModuloInt((self).fm9(d_996_v35_, len(d_1063_v78_), globalState), 990)) in (d_1062_v77_) else (default__.fm18((d_995_v34_)[default__.safeIndex(408, (d_995_v34_).length(0))], d_994_v33_, globalState)).set(d_981_v22_, (d_995_v34_)[default__.safeIndex(408, (d_995_v34_).length(0))]))
        if d_981_v22_:
            d_993_v32_ = d_993_v32_
            d_1064_v79_: _dafny.MultiSet
            d_1064_v79_ = _dafny.MultiSet([(0) - (d_993_v32_), d_993_v32_, d_993_v32_])
            d_1064_v79_ = ((default__.fm29(314, d_981_v22_, False, globalState)).set(d_993_v32_, default__.abs(len(d_996_v35_)))) | ((d_1064_v79_ if d_981_v22_ else d_1064_v79_))
            d_1065_v80_: str
            d_1065_v80_ = _dafny.CodePoint('b')
            d_1066_v81_: D3
            d_1066_v81_ = D3_DC8(d_1065_v80_, True, d_981_v22_, default__.fm2(not(not(not(d_981_v22_))), d_993_v32_, d_1065_v80_, d_993_v32_, globalState))
            d_1067_v82_: _dafny.Seq
            d_1067_v82_ = _dafny.SeqWithoutIsStrInference([d_995_v34_])
            d_1068_v83_: _dafny.Set
            d_1068_v83_ = _dafny.Set({d_981_v22_, (d_995_v34_)[default__.safeIndex(408, (d_995_v34_).length(0))], (d_995_v34_)[default__.safeIndex(408, (d_995_v34_).length(0))], (d_995_v34_)[default__.safeIndex(408, (d_995_v34_).length(0))]})
            (globalState).f1 = (self).fm9((((d_1066_v81_).cf20).set(default__.safeIndex(d_993_v32_, len((d_1066_v81_).cf20)), default__.fm1(globalState))).set(default__.safeIndex(len(d_1067_v82_), len(((d_1066_v81_).cf20).set(default__.safeIndex(d_993_v32_, len((d_1066_v81_).cf20)), default__.fm1(globalState)))), d_1065_v80_), (len(d_1068_v83_)) + (d_993_v32_), globalState)
            d_1069_v84_: T1
            nw189_ = C3()
            nw189_.ctor__(d_981_v22_)
            d_1069_v84_ = nw189_
            d_1070_v85_: _dafny.Map
            d_1070_v85_ = _dafny.Map({d_1069_v84_: (default__.fm4(d_998_v37_, globalState)) == (False)})
            index183_ = default__.safeIndex(408, (d_995_v34_).length(0))
            (d_995_v34_)[index183_] = ((d_1070_v85_)[d_1069_v84_] if (d_1069_v84_) in (d_1070_v85_) else d_981_v22_)
            index184_ = default__.safeIndex(408, (d_995_v34_).length(0))
            (d_995_v34_)[index184_] = not(default__.fm4(d_998_v37_, globalState))
        elif True:
            d_1071_v86_: _dafny.Array
            nw190_ = _dafny.Array(None, 3)
            nw190_[int(0)] = d_996_v35_
            nw190_[int(1)] = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rxgwn"))).set(default__.safeIndex((d_982_v23_).cardinality, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rxgwn")))), (d_996_v35_)[default__.safeIndex(d_993_v32_, len(d_996_v35_))])) + (d_996_v35_)
            nw190_[int(2)] = d_996_v35_
            d_1071_v86_ = nw190_
            d_1072_v87_: _dafny.Array
            nw191_ = _dafny.Array(_dafny.Array(None, 0), 21)
            d_1072_v87_ = nw191_
            d_1073_v88_: _dafny.Array
            def lambda61_(d_1074_v37_):
                def lambda62_(d_1075_i5_):
                    return d_1074_v37_

                return lambda62_

            init33_ = lambda61_(d_998_v37_)
            nw192_ = _dafny.Array(None, 24)
            for i0_33_ in range(nw192_.length(0)):
                nw192_[i0_33_] = init33_(i0_33_)
            d_1073_v88_ = nw192_
            index185_ = default__.safeIndex(767, (d_1073_v88_).length(0))
            (d_1073_v88_)[index185_] = d_998_v37_
            index186_ = default__.safeIndex(408, (d_995_v34_).length(0))
            index187_ = default__.safeIndex(767, (d_1073_v88_).length(0))
            rhs166_ = (d_995_v34_)[default__.safeIndex(408, (d_995_v34_).length(0))]
            rhs167_ = d_1071_v86_
            rhs168_ = d_993_v32_
            rhs169_ = d_1072_v87_
            rhs170_ = d_998_v37_
            lhs113_ = d_995_v34_
            lhs114_ = default__.safeIndex(408, (d_995_v34_).length(0))
            lhs115_ = globalState
            lhs116_ = d_1073_v88_
            lhs117_ = default__.safeIndex(767, (d_1073_v88_).length(0))
            lhs113_[lhs114_] = rhs166_
            d_1071_v86_ = rhs167_
            lhs115_.f1 = rhs168_
            d_1072_v87_ = rhs169_
            lhs116_[lhs117_] = rhs170_
            d_1076_v89_: _dafny.Map
            d_1076_v89_ = _dafny.Map({d_995_v34_: d_993_v32_})
            d_1076_v89_ = (d_1076_v89_).set(d_995_v34_, (0) - ((d_993_v32_) - (d_993_v32_)))
            (globalState).f1 = d_993_v32_
            d_1077_v90_: _dafny.Set
            d_1077_v90_ = _dafny.Set({d_998_v37_})
            d_1078_v91_: _dafny.Array
            nw193_ = _dafny.Array(None, 10)
            nw193_[int(0)] = (0) - (d_993_v32_)
            nw193_[int(1)] = d_993_v32_
            nw193_[int(2)] = d_993_v32_
            nw193_[int(3)] = d_993_v32_
            nw193_[int(4)] = d_993_v32_
            nw193_[int(5)] = len(d_1076_v89_)
            nw193_[int(6)] = len((d_1077_v90_ if (d_995_v34_)[default__.safeIndex(408, (d_995_v34_).length(0))] else d_1077_v90_))
            nw193_[int(7)] = d_993_v32_
            nw193_[int(8)] = d_993_v32_
            nw193_[int(9)] = (847) + (d_993_v32_)
            d_1078_v91_ = nw193_
            d_1078_v91_ = d_1078_v91_
            index188_ = default__.safeIndex(112, (d_1071_v86_).length(0))
            (d_1071_v86_)[index188_] = d_996_v35_
            index189_ = default__.safeIndex(112, (d_1071_v86_).length(0))
            (d_1071_v86_)[index189_] = d_996_v35_
        d_1079_v93_: _dafny.Array
        def lambda63_(d_1080_v32_):
            def lambda64_(d_1081_i7_):
                def iife76_():
                    coll38_ = _dafny.Set()
                    compr_38_: int
                    for compr_38_ in _dafny.IntegerRange(555, 410):
                        d_1082_v92_: int = compr_38_
                        if ((555) <= (d_1082_v92_)) and ((d_1082_v92_) < (410)):
                            coll38_ = coll38_.union(_dafny.Set([(d_1082_v92_) * (d_1080_v32_)]))
                    return _dafny.Set(coll38_)
                return iife76_()
                

            return lambda64_

        init34_ = lambda63_(d_993_v32_)
        nw194_ = _dafny.Array(None, 8)
        for i0_34_ in range(nw194_.length(0)):
            nw194_[i0_34_] = init34_(i0_34_)
        d_1079_v93_ = nw194_
        guard_loop_5_: int
        for guard_loop_5_ in _dafny.IntegerRange(0, (d_1079_v93_).length(0)):
            d_1083_i6_: int = guard_loop_5_
            if (True) and (((0) <= (d_1083_i6_)) and ((d_1083_i6_) < ((d_1079_v93_).length(0)))):
                (d_1079_v93_)[(d_1083_i6_)] = (d_998_v37_).intersection(d_998_v37_)
        d_1084_v94_: D0
        d_1084_v94_ = D0_DC0(_dafny.CodePoint('e'))
        pat_let_tv26_ = d_995_v34_
        pat_let_tv27_ = d_995_v34_
        pat_let_tv28_ = d_996_v35_
        def lambda65_(source13_):
            if source13_.is_DC1:
                d_1085___mcc_h7_ = source13_.cf1
                d_1086___mcc_h8_ = source13_.cf2
                d_1087_cf2_ = d_1086___mcc_h8_
                d_1088_cf1_ = d_1085___mcc_h7_
                return not(d_1088_cf1_)
            elif source13_.is_DC0:
                d_1089___mcc_h9_ = source13_.cf0
                d_1090_cf0_ = d_1089___mcc_h9_
                return not(not((pat_let_tv27_)[default__.safeIndex(408, (pat_let_tv26_).length(0))]))
            elif True:
                d_1091___mcc_h10_ = source13_.cf3
                d_1092_cf3_ = d_1091___mcc_h10_
                return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xq"))) != (pat_let_tv28_)

        r0 = not(lambda65_(D0_DC2(d_1084_v94_)))
        nw195_ = _dafny.Array(None, 4)
        nw195_[int(0)] = d_995_v34_
        nw195_[int(1)] = d_995_v34_
        nw195_[int(2)] = d_995_v34_
        nw195_[int(3)] = d_995_v34_
        r1 = nw195_
        r2 = d_993_v32_
        d_1093_v95_: str
        d_1093_v95_ = _dafny.CodePoint('w')
        d_1094_v96_: _dafny.Map
        d_1094_v96_ = _dafny.Map({(d_995_v34_)[default__.safeIndex(408, (d_995_v34_).length(0))]: default__.fm19(d_1093_v95_, (d_995_v34_)[default__.safeIndex(408, (d_995_v34_).length(0))], globalState)})
        r3 = (d_993_v32_) - (default__.safeModuloInt(d_993_v32_, ((d_1094_v96_)[True] if (True) in (d_1094_v96_) else d_993_v32_)))
        return r0, r1, r2, r3


class C5(T0, T1, T2):
    def  __init__(self):
        self._f10: _dafny.MultiSet = _dafny.MultiSet({})
        self._f11: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    @property
    def f10(self):
        return self._f10
    def ctor__(self, f11, f10):
        (self)._f11 = f11
        (self)._f10 = f10

    def fm5(self, p0, p1, p2, globalState):
        return False

    def fm6(self, p0, p1, globalState):
        return len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_1095_i0_ in range(default__.abs(921))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pqoiu"))))

    def fm7(self, p0, globalState):
        return False

    def m1(self, p0, p1, p2, globalState):
        d_1096_v0_: _dafny.Seq
        d_1096_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wbyaw"))
        d_1097_v1_: D0
        d_1097_v1_ = D0_DC0(p0)
        d_1098_v2_: _dafny.Map
        d_1098_v2_ = _dafny.Map({p1: p1})
        d_1099_v3_: _dafny.MultiSet
        d_1099_v3_ = _dafny.MultiSet([721])
        d_1100_v4_: _dafny.Set
        d_1100_v4_ = _dafny.Set({len(d_1098_v2_), (D0_DC1((self).f11, (d_1099_v3_).cardinality)).cf2})
        d_1101_v5_: _dafny.Seq
        d_1101_v5_ = _dafny.SeqWithoutIsStrInference([(self).f11])
        d_1102_v6_: _dafny.Seq
        d_1102_v6_ = _dafny.SeqWithoutIsStrInference([p1, p1, p1])
        d_1103_v7_: _dafny.Set
        d_1103_v7_ = _dafny.Set({(self).f11})
        d_1104_v8_: _dafny.Array
        def lambda66_(d_1105_p2_):
            def lambda67_(d_1106_i0_):
                return (d_1106_i0_) - ((d_1105_p2_).cf2)

            return lambda67_

        init35_ = lambda66_(p2)
        nw196_ = _dafny.Array(None, 23)
        for i0_35_ in range(nw196_.length(0)):
            nw196_[i0_35_] = init35_(i0_35_)
        d_1104_v8_ = nw196_
        d_1107_v9_: _dafny.Map
        d_1107_v9_ = _dafny.Map({d_1104_v8_: p1})
        d_1108_v10_: _dafny.Array
        nw197_ = _dafny.Array(None, 19)
        nw197_[int(0)] = default__.safeModuloInt((self).fm6((self).f11, 657, globalState), 641)
        nw197_[int(1)] = p1
        nw197_[int(2)] = -842
        nw197_[int(3)] = p1
        nw197_[int(4)] = p1
        nw197_[int(5)] = p1
        nw197_[int(6)] = (self).fm6(default__.fm4(d_1100_v4_, globalState), p1, globalState)
        nw197_[int(7)] = default__.safeDivisionInt(len((d_1101_v5_).set(default__.safeIndex(p1, len(d_1101_v5_)), (self).f11)), (d_1102_v6_)[default__.safeIndex(p1, len(d_1102_v6_))])
        nw197_[int(8)] = p1
        nw197_[int(9)] = len(d_1103_v7_)
        nw197_[int(10)] = (p1) * ((0) - (p1))
        nw197_[int(11)] = (self).fm6(False, p1, globalState)
        nw197_[int(12)] = len(d_1107_v9_)
        nw197_[int(13)] = len((d_1096_v0_) + (d_1096_v0_))
        nw197_[int(14)] = (p1) * (407)
        nw197_[int(15)] = p1
        nw197_[int(16)] = p1
        nw197_[int(17)] = 954
        nw197_[int(18)] = p1
        d_1108_v10_ = nw197_
        d_1109_v11_: D0
        d_1109_v11_ = D0_DC1((self).f11, p1)
        d_1110_v12_: D0
        d_1110_v12_ = D0_DC2(d_1109_v11_)
        d_1111_v13_: _dafny.Map
        d_1111_v13_ = _dafny.Map({p1: p0})
        rhs171_ = (d_1096_v0_) + (d_1096_v0_)
        rhs172_ = d_1097_v1_
        rhs173_ = d_1104_v8_
        rhs174_ = d_1110_v12_
        rhs175_ = d_1111_v13_
        d_1096_v0_ = rhs171_
        d_1097_v1_ = rhs172_
        d_1108_v10_ = rhs173_
        d_1110_v12_ = rhs174_
        d_1111_v13_ = rhs175_
        if ((d_1102_v6_ if (p2).cf1 else d_1102_v6_)) <= (default__.fm8(default__.fm4(d_1100_v4_, globalState), globalState)):
            d_1112_v14_: bool
            d_1112_v14_ = False
            d_1112_v14_ = not((423) > (p1))
            d_1113_v15_: bool
            d_1114_v16_: _dafny.MultiSet
            d_1115_v17_: _dafny.Set
            d_1116_v18_: int
            out36_: bool
            out37_: _dafny.MultiSet
            out38_: _dafny.Set
            out39_: int
            out36_, out37_, out38_, out39_ = default__.m0(_dafny.SeqWithoutIsStrInference([True, not(not(True))]), d_1101_v5_, globalState)
            d_1113_v15_ = out36_
            d_1114_v16_ = out37_
            d_1115_v17_ = out38_
            d_1116_v18_ = out39_
            d_1116_v18_ = (0) - (default__.safeDivisionInt((len(d_1096_v0_)) - (d_1116_v18_), d_1116_v18_))
            d_1112_v14_ = (d_1116_v18_) != (p1)
            d_1117_v19_: _dafny.Map
            d_1117_v19_ = _dafny.Map({d_1116_v18_: (self).f11})
            (globalState).f1 = (0) - ((d_1102_v6_)[default__.safeIndex((d_1116_v18_) * (len(d_1117_v19_)), len(d_1102_v6_))])
        elif True:
            d_1118_v20_: C3
            nw198_ = C3()
            nw198_.ctor__((self).f11)
            d_1118_v20_ = nw198_
            d_1119_v21_: bool
            d_1119_v21_ = True
            rhs176_ = d_1119_v21_
            rhs177_ = (self).f11
            d_1119_v21_ = rhs176_
            d_1119_v21_ = rhs177_
            if d_1119_v21_:
                d_1120_v22_: _dafny.Map
                d_1120_v22_ = _dafny.Map({False: (self).f11})
                d_1121_v24_: _dafny.Map
                def iife77_():
                    coll39_ = _dafny.Map()
                    compr_39_: int
                    for compr_39_ in _dafny.IntegerRange(213, 198):
                        d_1122_v23_: int = compr_39_
                        if ((213) <= (d_1122_v23_)) and ((d_1122_v23_) < (198)):
                            coll39_[default__.safeDivisionInt(d_1122_v23_, p1)] = p1
                    return _dafny.Map(coll39_)
                d_1121_v24_ = _dafny.Map({len(iife77_()
                ): d_1119_v21_})
                d_1123_v25_: _dafny.Set
                d_1123_v25_ = _dafny.Set({(d_1120_v22_).set((d_1118_v20_).f16, False), default__.fm18(True, d_1121_v24_, globalState), (d_1120_v22_ if d_1119_v21_ else d_1120_v22_)})
                (globalState).f1 = len(d_1123_v25_)
                d_1124_v26_: _dafny.Array
                def lambda68_(d_1125_v7_):
                    def lambda69_(d_1126_i1_):
                        return (d_1125_v7_) | (_dafny.Set({False}))

                    return lambda69_

                init36_ = lambda68_(d_1103_v7_)
                nw199_ = _dafny.Array(None, 9)
                for i0_36_ in range(nw199_.length(0)):
                    nw199_[i0_36_] = init36_(i0_36_)
                d_1124_v26_ = nw199_
                index190_ = default__.safeIndex(18, (d_1124_v26_).length(0))
                (d_1124_v26_)[index190_] = _dafny.Set({(d_1118_v20_).f16, (d_1118_v20_).f16, (d_1118_v20_).f16})
                index191_ = default__.safeIndex(18, (d_1124_v26_).length(0))
                (d_1124_v26_)[index191_] = d_1103_v7_
                (globalState).f1 = (483) - (len(d_1096_v0_))
                d_1127_v27_: _dafny.Array
                nw200_ = _dafny.Array(False, 15)
                d_1127_v27_ = nw200_
                d_1128_v28_: _dafny.Seq
                d_1128_v28_ = _dafny.SeqWithoutIsStrInference([d_1127_v27_, d_1127_v27_])
                d_1129_v29_: _dafny.Seq
                d_1129_v29_ = _dafny.SeqWithoutIsStrInference([(d_1127_v27_ if (d_1118_v20_).f16 else d_1127_v27_), (d_1127_v27_ if (d_1118_v20_).f16 else d_1127_v27_), d_1127_v27_, (d_1128_v28_)[default__.safeIndex(p1, len(d_1128_v28_))]])
                rhs178_ = p1
                rhs179_ = (d_1128_v28_)[default__.safeIndex((0) - (p1), len(d_1128_v28_))]
                rhs180_ = d_1104_v8_
                rhs181_ = d_1096_v0_
                lhs118_ = globalState
                lhs119_ = globalState
                lhs118_.f1 = rhs178_
                lhs119_.f5 = rhs179_
                d_1104_v8_ = rhs180_
                d_1096_v0_ = rhs181_
                (globalState).f1 = len(default__.fm14((d_1118_v20_).f16, globalState))
            elif True:
                d_1130_v30_: _dafny.Array
                def lambda70_(d_1131_v20_):
                    def lambda71_(d_1132_i2_):
                        return (d_1131_v20_).f16

                    return lambda71_

                init37_ = lambda70_(d_1118_v20_)
                nw201_ = _dafny.Array(None, 20)
                for i0_37_ in range(nw201_.length(0)):
                    nw201_[i0_37_] = init37_(i0_37_)
                d_1130_v30_ = nw201_
                d_1133_v31_: D6
                d_1133_v31_ = D6_DC17(len(d_1101_v5_), d_1096_v0_, True)
                index192_ = default__.safeIndex(188, (d_1130_v30_).length(0))
                (d_1130_v30_)[index192_] = (d_1119_v21_) and ((d_1133_v31_).cf40)
                index193_ = default__.safeIndex(737, (d_1130_v30_).length(0))
                (d_1130_v30_)[index193_] = (d_1118_v20_).f16
                d_1134_v32_: _dafny.Map
                d_1134_v32_ = _dafny.Map({d_1119_v21_: p1})
                d_1135_v33_: _dafny.Map
                d_1135_v33_ = _dafny.Map({p1: (self).f11})
                d_1136_v34_: _dafny.Seq
                d_1136_v34_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xrghkd"))])
                d_1137_v35_: C2
                nw202_ = C2()
                nw202_.ctor__((self).f11, p1, (self).f10)
                d_1137_v35_ = nw202_
                d_1138_v36_: _dafny.Map
                d_1138_v36_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({(d_1118_v20_).f16: (d_1118_v20_).f16})), len(d_1135_v33_), len(d_1136_v34_)])).set(default__.safeIndex(p1, len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({(d_1118_v20_).f16: (d_1118_v20_).f16})), len(d_1135_v33_), len(d_1136_v34_)]))), p1): d_1137_v35_})
                index194_ = default__.safeIndex(188, (d_1130_v30_).length(0))
                index195_ = default__.safeIndex(737, (d_1130_v30_).length(0))
                rhs182_ = ((d_1134_v32_ if d_1119_v21_ else d_1134_v32_)) == ((d_1134_v32_) | (d_1134_v32_))
                rhs183_ = (d_1138_v36_) == (_dafny.Map({d_1102_v6_: d_1137_v35_}))
                lhs120_ = d_1130_v30_
                lhs121_ = default__.safeIndex(188, (d_1130_v30_).length(0))
                lhs122_ = d_1130_v30_
                lhs123_ = default__.safeIndex(737, (d_1130_v30_).length(0))
                lhs120_[lhs121_] = rhs182_
                lhs122_[lhs123_] = rhs183_
                index196_ = default__.safeIndex(188, (d_1130_v30_).length(0))
                rhs184_ = ((False) == ((d_1130_v30_)[default__.safeIndex(188, (d_1130_v30_).length(0))])) and ((d_1118_v20_).f16)
                rhs185_ = 730
                lhs124_ = d_1130_v30_
                lhs125_ = default__.safeIndex(188, (d_1130_v30_).length(0))
                lhs126_ = globalState
                lhs124_[lhs125_] = rhs184_
                lhs126_.f1 = rhs185_
                d_1139_v37_: _dafny.Array
                nw203_ = _dafny.Array(_dafny.CodePoint('D'), 23)
                d_1139_v37_ = nw203_
                index197_ = default__.safeIndex(834, (d_1139_v37_).length(0))
                (d_1139_v37_)[index197_] = _dafny.CodePoint('s')
                d_1140_v38_: D8
                d_1140_v38_ = D8_DC20(p0, (self).f11, d_1119_v21_, d_1134_v32_, d_1119_v21_)
                index198_ = default__.safeIndex(834, (d_1139_v37_).length(0))
                rhs186_ = (((d_1134_v32_)[(self).f11] if ((self).f11) in (d_1134_v32_) else p1)) - (default__.safeDivisionInt(p1, (d_1137_v35_).f13))
                rhs187_ = (p0 if (d_1118_v20_).f16 else (p0 if d_1137_v35_.f12 else (d_1140_v38_).cf43))
                rhs188_ = (d_1137_v35_).f13
                lhs127_ = globalState
                lhs128_ = d_1139_v37_
                lhs129_ = default__.safeIndex(834, (d_1139_v37_).length(0))
                lhs130_ = globalState
                lhs127_.f1 = rhs186_
                lhs128_[lhs129_] = rhs187_
                lhs130_.f1 = rhs188_
                (d_1137_v35_).f12 = (d_1130_v30_)[default__.safeIndex(188, (d_1130_v30_).length(0))]
                index199_ = default__.safeIndex(188, (d_1130_v30_).length(0))
                (d_1130_v30_)[index199_] = d_1119_v21_
            if (self).f11:
                d_1101_v5_ = d_1101_v5_
                (globalState).f1 = (p1) * (p1)
                d_1096_v0_ = ((_dafny.SeqWithoutIsStrInference([p0 for d_1141_i3_ in range(default__.abs(260))])) + (d_1096_v0_)) + (d_1096_v0_)
                d_1142_v39_: _dafny.Map
                d_1142_v39_ = _dafny.Map({d_1104_v8_: not (d_1119_v21_) or ((d_1118_v20_).f16)})
                d_1142_v39_ = _dafny.Map({d_1104_v8_: (d_1118_v20_).f16})
                d_1143_v40_: _dafny.Map
                d_1143_v40_ = _dafny.Map({d_1119_v21_: p1})
                index200_ = default__.safeIndex(172, (d_1104_v8_).length(0))
                (d_1104_v8_)[index200_] = len((d_1143_v40_).set((self).f11, p1))
                index201_ = default__.safeIndex(172, (d_1104_v8_).length(0))
                (d_1104_v8_)[index201_] = p1
            elif True:
                d_1144_v41_: _dafny.Array
                nw204_ = _dafny.Array(False, 9)
                d_1144_v41_ = nw204_
                (globalState).f5 = d_1144_v41_
                index202_ = default__.safeIndex(486, (d_1144_v41_).length(0))
                (d_1144_v41_)[index202_] = (d_1118_v20_).f16
                index203_ = default__.safeIndex(486, (d_1144_v41_).length(0))
                (d_1144_v41_)[index203_] = d_1119_v21_
                index204_ = default__.safeIndex(486, (d_1144_v41_).length(0))
                index205_ = default__.safeIndex(486, (d_1144_v41_).length(0))
                rhs189_ = (322) - (p1)
                rhs190_ = (d_1118_v20_).f16
                rhs191_ = (618) + (p1)
                rhs192_ = p1
                rhs193_ = (self).fm5((d_1103_v7_) | (d_1103_v7_), p1, (d_1102_v6_)[default__.safeIndex(p1, len(d_1102_v6_))], globalState)
                lhs131_ = globalState
                lhs132_ = d_1144_v41_
                lhs133_ = default__.safeIndex(486, (d_1144_v41_).length(0))
                lhs134_ = globalState
                lhs135_ = globalState
                lhs136_ = d_1144_v41_
                lhs137_ = default__.safeIndex(486, (d_1144_v41_).length(0))
                lhs131_.f1 = rhs189_
                lhs132_[lhs133_] = rhs190_
                lhs134_.f1 = rhs191_
                lhs135_.f1 = rhs192_
                lhs136_[lhs137_] = rhs193_
                (globalState).f1 = default__.safeModuloInt(p1, p1)
                d_1119_v21_ = (d_1118_v20_).f16
            d_1145_v42_: _dafny.Map
            d_1145_v42_ = _dafny.Map({p1: (d_1118_v20_).f16})
            d_1145_v42_ = (d_1145_v42_) | (d_1145_v42_)
        if (self).f11:
            d_1146_v43_: _dafny.Array
            nw205_ = _dafny.Array(None, 12)
            nw205_[int(0)] = (self).f11
            nw205_[int(1)] = (self).f11
            nw205_[int(2)] = (self).f11
            nw205_[int(3)] = (self).f11
            nw205_[int(4)] = (self).f11
            nw205_[int(5)] = (self).f11
            nw205_[int(6)] = (self).f11
            nw205_[int(7)] = not((self).f11)
            nw205_[int(8)] = False
            nw205_[int(9)] = (self).f11
            nw205_[int(10)] = (self).f11
            nw205_[int(11)] = True
            d_1146_v43_ = nw205_
            d_1147_v44_: _dafny.Array
            nw206_ = _dafny.Array(None, 5)
            nw206_[int(0)] = (d_1146_v43_ if (self).f11 else d_1146_v43_)
            nw206_[int(1)] = d_1146_v43_
            nw206_[int(2)] = d_1146_v43_
            nw206_[int(3)] = d_1146_v43_
            nw206_[int(4)] = d_1146_v43_
            d_1147_v44_ = nw206_
            index206_ = default__.safeIndex(309, (d_1147_v44_).length(0))
            (d_1147_v44_)[index206_] = d_1146_v43_
            index207_ = default__.safeIndex(309, (d_1147_v44_).length(0))
            (d_1147_v44_)[index207_] = d_1146_v43_
            index208_ = default__.safeIndex(852, (d_1104_v8_).length(0))
            (d_1104_v8_)[index208_] = p1
            d_1148_v45_: bool
            d_1148_v45_ = False
            index209_ = default__.safeIndex(852, (d_1104_v8_).length(0))
            rhs194_ = (d_1102_v6_) + ((_dafny.SeqWithoutIsStrInference([p1, p1])) + (d_1102_v6_))
            rhs195_ = p1
            rhs196_ = (self).fm6(d_1148_v45_, p1, globalState)
            rhs197_ = (not(False)) and ((self).f11)
            lhs138_ = globalState
            lhs139_ = d_1104_v8_
            lhs140_ = default__.safeIndex(852, (d_1104_v8_).length(0))
            d_1102_v6_ = rhs194_
            lhs138_.f1 = rhs195_
            lhs139_[lhs140_] = rhs196_
            d_1148_v45_ = rhs197_
            arr0_ = (d_1147_v44_)[default__.safeIndex(309, (d_1147_v44_).length(0))]
            index210_ = default__.safeIndex(877, ((d_1147_v44_)[default__.safeIndex(309, (d_1147_v44_).length(0))]).length(0))
            arr0_[index210_] = d_1148_v45_
            arr1_ = (d_1147_v44_)[default__.safeIndex(309, (d_1147_v44_).length(0))]
            index211_ = default__.safeIndex(877, ((d_1147_v44_)[default__.safeIndex(309, (d_1147_v44_).length(0))]).length(0))
            arr1_[index211_] = (default__.fm19(p0, d_1148_v45_, globalState)) >= (p1)
            arr2_ = (d_1147_v44_)[default__.safeIndex(309, (d_1147_v44_).length(0))]
            index212_ = default__.safeIndex(877, ((d_1147_v44_)[default__.safeIndex(309, (d_1147_v44_).length(0))]).length(0))
            arr2_[index212_] = (self).f11
            d_1149_v46_: str
            d_1149_v46_ = _dafny.CodePoint('x')
            d_1149_v46_ = p0
        elif True:
            d_1150_v47_: bool
            d_1150_v47_ = True
            d_1150_v47_ = (p1) == (p1)
            if d_1150_v47_:
                (globalState).f1 = ((p1) * (p1) if not ((self).f11) or (d_1150_v47_) else p1)
                d_1150_v47_ = not(False)
                d_1150_v47_ = not(d_1150_v47_)
                (globalState).f1 = p1
                d_1151_v48_: T2
                nw207_ = C0()
                nw207_.ctor__()
                d_1151_v48_ = nw207_
                d_1151_v48_ = d_1151_v48_
            elif True:
                d_1152_v49_: C4
                nw208_ = C4()
                nw208_.ctor__()
                d_1152_v49_ = nw208_
                d_1153_v50_: D3
                d_1153_v50_ = D3_DC8(_dafny.CodePoint('m'), (self).f11, default__.fm4(d_1100_v4_, globalState), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jhadxaocd"))) + ((d_1096_v0_).set(default__.safeIndex(p1, len(d_1096_v0_)), _dafny.CodePoint('k'))))
                d_1153_v50_ = d_1153_v50_
                (globalState).f1 = (0) - (default__.safeDivisionInt(p1, (0) - ((((self).f10) - (((self).f10).set(d_1150_v47_, default__.abs(p1)))).cardinality)))
                d_1154_v51_: C2
                nw209_ = C2()
                nw209_.ctor__((len(d_1102_v6_)) not in (d_1100_v4_), (p1) * (len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kwiask"))).set(default__.safeIndex(len((d_1096_v0_).set(default__.safeIndex(p1, len(d_1096_v0_)), p0)), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kwiask")))), _dafny.CodePoint('e')))), (self).f10)
                d_1154_v51_ = nw209_
                (globalState).f1 = (d_1154_v51_).f13
            d_1155_v52_: C2
            nw210_ = C2()
            nw210_.ctor__((d_1099_v3_).ispropersubset(d_1099_v3_), p1, (self).f10)
            d_1155_v52_ = nw210_
            d_1156_v53_: _dafny.Array
            nw211_ = _dafny.Array(None, 3)
            nw211_[int(0)] = ((d_1096_v0_) + (d_1096_v0_)).set(default__.safeIndex(p1, len((d_1096_v0_) + (d_1096_v0_))), _dafny.CodePoint('d'))
            nw211_[int(1)] = default__.fm2(d_1155_v52_.f12, len(d_1100_v4_), p0, (d_1155_v52_).f13, globalState)
            nw211_[int(2)] = (d_1096_v0_).set(default__.safeIndex((d_1155_v52_).f13, len(d_1096_v0_)), p0)
            d_1156_v53_ = nw211_
            index213_ = default__.safeIndex(418, (d_1156_v53_).length(0))
            (d_1156_v53_)[index213_] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xonhdkbwi"))) + (d_1096_v0_)
            index214_ = default__.safeIndex(418, (d_1156_v53_).length(0))
            (d_1156_v53_)[index214_] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_1157_i4_ in range(default__.abs(187))])
            d_1158_v54_: _dafny.Map
            d_1158_v54_ = _dafny.Map({not((self).f11): (d_1108_v10_ if (self).f11 else d_1108_v10_)})
            d_1108_v10_ = ((d_1158_v54_)[(d_1101_v5_) == (d_1101_v5_)] if ((d_1101_v5_) == (d_1101_v5_)) in (d_1158_v54_) else d_1104_v8_)
        d_1159_v55_: bool
        d_1159_v55_ = False
        d_1160_v56_: _dafny.Array
        nw212_ = _dafny.Array(False, 3)
        d_1160_v56_ = nw212_
        d_1161_v57_: _dafny.Map
        d_1161_v57_ = _dafny.Map({d_1160_v56_: (self).f11})
        d_1162_v58_: _dafny.Map
        d_1162_v58_ = _dafny.Map({(self).f11: p0})
        d_1163_v59_: _dafny.Map
        d_1163_v59_ = _dafny.Map({(len(d_1162_v58_)) + (p1): (d_1096_v0_).set(default__.safeIndex(p1, len(d_1096_v0_)), p0)})
        d_1164_v60_: _dafny.Map
        d_1164_v60_ = _dafny.Map({default__.fm4(d_1100_v4_, globalState): d_1096_v0_})
        pat_let_tv29_ = p1
        def iife78_(_pat_let19_0):
            def iife79_(d_1165_dt__update__tmp_h0_):
                def iife80_(_pat_let20_0):
                    def iife81_(d_1166_dt__update_hcf2_h0_):
                        return D0_DC1((d_1165_dt__update__tmp_h0_).cf1, d_1166_dt__update_hcf2_h0_)
                    return iife81_(_pat_let20_0)
                return iife80_(pat_let_tv29_)
            return iife79_(_pat_let19_0)
        rhs198_ = ((((_dafny.SeqWithoutIsStrInference([p1]) if d_1159_v55_ else d_1102_v6_)).set(default__.safeIndex(p1, len((_dafny.SeqWithoutIsStrInference([p1]) if d_1159_v55_ else d_1102_v6_))), p1)).set(default__.safeIndex(len(d_1161_v57_), len(((_dafny.SeqWithoutIsStrInference([p1]) if d_1159_v55_ else d_1102_v6_)).set(default__.safeIndex(p1, len((_dafny.SeqWithoutIsStrInference([p1]) if d_1159_v55_ else d_1102_v6_))), p1))), (d_1102_v6_)[default__.safeIndex(p1, len(d_1102_v6_))])) + ((d_1102_v6_ if (self).f11 else d_1102_v6_))
        rhs199_ = (d_1101_v5_) < (default__.fm33(p1, (self).f11, (self).f11, p1, globalState))
        rhs200_ = (iife78_(p2)).cf1
        rhs201_ = (((d_1163_v59_)[p1] if (p1) in (d_1163_v59_) else ((d_1164_v60_)[(self).f11] if ((self).f11) in (d_1164_v60_) else d_1096_v0_))).set(default__.safeIndex((0) - (len(d_1096_v0_)), len(((d_1163_v59_)[p1] if (p1) in (d_1163_v59_) else ((d_1164_v60_)[(self).f11] if ((self).f11) in (d_1164_v60_) else d_1096_v0_)))), p0)
        rhs202_ = not((self).fm7(p1, globalState))
        d_1102_v6_ = rhs198_
        d_1159_v55_ = rhs199_
        d_1159_v55_ = rhs200_
        d_1096_v0_ = rhs201_
        d_1159_v55_ = rhs202_
        d_1167_v61_: _dafny.Array
        def lambda72_(d_1168_v4_):
            def lambda73_(d_1169_i6_):
                return d_1168_v4_

            return lambda73_

        init38_ = lambda72_(d_1100_v4_)
        nw213_ = _dafny.Array(None, 3)
        for i0_38_ in range(nw213_.length(0)):
            nw213_[i0_38_] = init38_(i0_38_)
        d_1167_v61_ = nw213_
        guard_loop_6_: int
        for guard_loop_6_ in _dafny.IntegerRange(0, (d_1167_v61_).length(0)):
            d_1170_i5_: int = guard_loop_6_
            if (True) and (((0) <= (d_1170_i5_)) and ((d_1170_i5_) < ((d_1167_v61_).length(0)))):
                def iife82_():
                    coll40_ = _dafny.Set()
                    compr_40_: int
                    for compr_40_ in _dafny.IntegerRange(102, 717):
                        d_1171_v62_: int = compr_40_
                        if ((102) <= (d_1171_v62_)) and ((d_1171_v62_) < (717)):
                            coll40_ = coll40_.union(_dafny.Set([default__.safeDivisionInt(d_1171_v62_, len(d_1101_v5_))]))
                    return _dafny.Set(coll40_)
                (d_1167_v61_)[(d_1170_i5_)] = (_dafny.Set({p1})).intersection((d_1100_v4_).intersection(iife82_()
                ))
        d_1172_v63_: _dafny.Array
        nw214_ = _dafny.Array(_dafny.Set({}), 8)
        d_1172_v63_ = nw214_
        index215_ = default__.safeIndex(667, (d_1172_v63_).length(0))
        (d_1172_v63_)[index215_] = d_1103_v7_
        index216_ = default__.safeIndex(667, (d_1172_v63_).length(0))
        (d_1172_v63_)[index216_] = d_1103_v7_

    @property
    def f11(self):
        return self._f11
