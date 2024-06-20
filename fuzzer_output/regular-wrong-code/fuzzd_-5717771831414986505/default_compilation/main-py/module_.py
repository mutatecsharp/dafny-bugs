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
    def fm5(globalState):
        if (_dafny.SeqWithoutIsStrInference([True])) != (_dafny.SeqWithoutIsStrInference([False])):
            return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_0_i1_ in range(default__.abs(811))]) for d_1_i0_ in range(default__.abs(988))])
        elif True:
            return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tuqlhd")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wcvbklxx"))])

    @staticmethod
    def fm8(p0, p1, p2, globalState):
        return _dafny.SeqWithoutIsStrInference([default__.safeDivisionInt(411, 749) for d_2_i0_ in range(default__.abs(297))])

    @staticmethod
    def fm9(globalState):
        return ((_dafny.Map({True: True})) | (_dafny.Map({True: True}))) | ((_dafny.Map({False: True})) | (_dafny.Map({not(False): not(True)})))

    @staticmethod
    def fm11(p0, globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_3_i0_ in range(default__.abs(307))])

    @staticmethod
    def fm12(globalState):
        return _dafny.MultiSet([default__.safeDivisionInt(len(_dafny.Map({True: -509})), 214)])

    @staticmethod
    def fm13(p0, p1, p2, p3, globalState):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: _dafny.Seq
            for compr_0_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_4_i0_ in range(default__.abs(375))]): not(True)})).keys.Elements:
                d_5_v0_: _dafny.Seq = compr_0_
                if (d_5_v0_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_4_i0_ in range(default__.abs(375))]): not(True)})):
                    coll0_[d_5_v0_] = False
            return _dafny.Map(coll0_)
        return (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gyflnn")): True})) | (iife0_()
        )

    @staticmethod
    def fm14(p0, p1, p2, p3, globalState):
        return ((_dafny.MultiSet([True])) - (_dafny.MultiSet([False, not(True)]))).intersection(_dafny.MultiSet([True, True, True]))

    @staticmethod
    def fm15(p0, globalState):
        return ((_dafny.MultiSet([193, len(_dafny.Set({_dafny.CodePoint('l')})), 546])) - (_dafny.MultiSet([545, 965]))) - (_dafny.MultiSet([808]))

    @staticmethod
    def fm16(p0, p1, p2, globalState):
        return _dafny.CodePoint('k')

    @staticmethod
    def fm17(globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qrabdg"))

    @staticmethod
    def fm18(globalState):
        return (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([not(True), not(True)])) for d_6_i0_ in range(default__.abs(189))])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.Set({False})), -533, len(_dafny.Map({False: True}))]))

    @staticmethod
    def fm19(p0, globalState):
        return (((D7_DC19(_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True]))).cardinality]))}))).cf28 if True else _dafny.Map({False: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u')]))}))) | (_dafny.Map({not(not(True)): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n")))}))

    @staticmethod
    def fm20(p0, globalState):
        return (_dafny.Map({_dafny.SeqWithoutIsStrInference([False, not(False)]): False})) | ((_dafny.Map({_dafny.SeqWithoutIsStrInference([True]): not(True)})) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([True]): True})))

    @staticmethod
    def fm21(p0, globalState):
        source0_ = D5_DC15(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_7_i0_ in range(default__.abs(445))]))
        if source0_.is_DC16:
            d_8___mcc_h0_ = source0_.cf19
            d_9___mcc_h1_ = source0_.cf20
            d_10___mcc_h2_ = source0_.cf21
            d_11___mcc_h3_ = source0_.cf22
            d_12___mcc_h4_ = source0_.cf23
            d_13_cf23_ = d_12___mcc_h4_
            d_14_cf22_ = d_11___mcc_h3_
            d_15_cf21_ = d_10___mcc_h2_
            d_16_cf20_ = d_9___mcc_h1_
            d_17_cf19_ = d_8___mcc_h0_
            return D3_DC10(True, d_16_cf20_)
        elif True:
            d_18___mcc_h5_ = source0_.cf18
            d_19_cf18_ = d_18___mcc_h5_
            return D3_DC10(True, 713)

    @staticmethod
    def fm22(globalState):
        if (False if True else False):
            return D1_DC3(D0_DC1(), len(_dafny.Map({True: _dafny.SeqWithoutIsStrInference([False])})))
        elif True:
            def iife1_():
                coll1_ = _dafny.Map()
                compr_1_: int
                for compr_1_ in (_dafny.SeqWithoutIsStrInference([470, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_20_i0_ in range(default__.abs(334))])), 178])).Elements:
                    d_21_v0_: int = compr_1_
                    if (d_21_v0_) in (_dafny.SeqWithoutIsStrInference([470, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_20_i0_ in range(default__.abs(334))])), 178])):
                        coll1_[(d_21_v0_) + (493)] = 650
                return _dafny.Map(coll1_)
            return D1_DC3(D0_DC1(), len(iife1_()
))

    @staticmethod
    def fm23(p0, globalState):
        return _dafny.SeqWithoutIsStrInference([False])

    @staticmethod
    def fm24(p0, p1, p2, globalState):
        return (_dafny.Set({670, (0) - (len(_dafny.Set({(_dafny.MultiSet([900, 919, 804, 901])).cardinality, -743})))})) - ((_dafny.Set({-722, 152, len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True, True])), -562])), 865, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_22_i0_ in range(default__.abs(648))])), 132])), len(_dafny.Map({_dafny.Map({800: len(_dafny.SeqWithoutIsStrInference([331]))}): -891}))})) - (_dafny.Set({(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([False]))])).cardinality, -476, 959, len(_dafny.SeqWithoutIsStrInference([False, True]))})))

    @staticmethod
    def fm25(p0, p1, p2, globalState):
        return D0_DC0(len(_dafny.SeqWithoutIsStrInference([not(True), False, False])))

    @staticmethod
    def fm26(p0, globalState):
        def iife2_():
            def iife4_():
                coll4_ = _dafny.Map()
                compr_4_: int
                for compr_4_ in _dafny.IntegerRange(-93, 105):
                    d_23_v1_: int = compr_4_
                    if ((-93) <= (d_23_v1_)) and ((d_23_v1_) < (105)):
                        coll4_[(d_23_v1_) - (511)] = 96
                return _dafny.Map(coll4_)
            coll2_ = _dafny.Map()
            def iife3_():
                coll3_ = _dafny.Map()
                compr_3_: int
                for compr_3_ in _dafny.IntegerRange(-93, 105):
                    d_23_v1_: int = compr_3_
                    if ((-93) <= (d_23_v1_)) and ((d_23_v1_) < (105)):
                        coll3_[(d_23_v1_) - (511)] = 96
                return _dafny.Map(coll3_)
            compr_2_: _dafny.Map
            for compr_2_ in (_dafny.SeqWithoutIsStrInference([iife3_()
            ])).Elements:
                d_24_v0_: _dafny.Map = compr_2_
                if (d_24_v0_) in (_dafny.SeqWithoutIsStrInference([iife4_()
                ])):
                    coll2_[d_24_v0_] = False
            return _dafny.Map(coll2_)
        return (_dafny.Map({_dafny.Map({len(_dafny.SeqWithoutIsStrInference([212, 741])): 393}): True})) | (iife2_()
        )

    @staticmethod
    def fm27(p0, p1, p2, p3, globalState):
        return _dafny.Map({790: len(((D2_DC4(_dafny.Set({len(_dafny.Map({-544: not(True)})), (_dafny.MultiSet([(_dafny.MultiSet([len(_dafny.Map({(0) - (-406): not(True)}))])).cardinality, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sqqfutbxm")))])).cardinality}))).cf4) - (_dafny.Set({854})))})

    @staticmethod
    def fm28(p0, p1, p2, p3, globalState):
        if False:
            def iife5_():
                coll5_ = _dafny.Map()
                compr_5_: int
                for compr_5_ in _dafny.IntegerRange(-999, 240):
                    d_25_v0_: int = compr_5_
                    if ((-999) <= (d_25_v0_)) and ((d_25_v0_) < (240)):
                        coll5_[(d_25_v0_) - (len(_dafny.SeqWithoutIsStrInference([-219])))] = -167
                return _dafny.Map(coll5_)
            return (_dafny.MultiSet([_dafny.Map({-750: len(_dafny.Set({-201, 521}))}), iife5_()
            ])) - (_dafny.MultiSet([_dafny.Map({245: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a'), _dafny.CodePoint('p')]))}), _dafny.Map({-699: len(_dafny.Set({True}))}), _dafny.Map({999: 629}), _dafny.Map({len(_dafny.Set({len(_dafny.Map({True: True}))})): len(_dafny.Map({len(_dafny.Map({not(False): 52})): True}))}), _dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gsppdqr")))): 272})]))
        elif True:
            return _dafny.MultiSet([_dafny.Map({len(_dafny.Set({553, len(_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_26_i0_ in range(default__.abs(994))]))}))})): (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([D18_DC49(_dafny.SeqWithoutIsStrInference([True]), (_dafny.MultiSet([_dafny.CodePoint('c')])).cardinality)]))).cardinality}), _dafny.Map({len(_dafny.Set({-594})): 913})])
        elif True:
            def iife6_():
                coll6_ = _dafny.Set()
                compr_6_: int
                for compr_6_ in _dafny.IntegerRange(414, 809):
                    d_27_v1_: int = compr_6_
                    if ((414) <= (d_27_v1_)) and ((d_27_v1_) < (809)):
                        coll6_ = coll6_.union(_dafny.Set([(d_27_v1_) + (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_28_i2_ in range(default__.abs(232))])))]))
                return _dafny.Set(coll6_)
            return _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.Map({78: 765}), _dafny.Map({665: len(_dafny.Set({_dafny.Set({len(_dafny.SeqWithoutIsStrInference([_dafny.Map({len(iife6_()
): 906}) for d_29_i1_ in range(default__.abs(-553))]))})}))}), _dafny.Map({458: 240})]))

    @staticmethod
    def fm29(p0, p1, p2, globalState):
        return D7_DC21((len(_dafny.Map({(0) - (len(_dafny.Set({(_dafny.MultiSet([_dafny.CodePoint('u')])).cardinality}))): len(_dafny.SeqWithoutIsStrInference([False]))}))) - (len(_dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qyyg")))): 29}))), len(_dafny.SeqWithoutIsStrInference([False])))

    @staticmethod
    def fm30(globalState):
        return D11_DC32(_dafny.Map({True: False}))

    @staticmethod
    def fm31(globalState):
        def iife7_():
            coll7_ = _dafny.Map()
            compr_7_: int
            for compr_7_ in _dafny.IntegerRange(137, -918):
                d_30_v0_: int = compr_7_
                if ((137) <= (d_30_v0_)) and ((d_30_v0_) < (-918)):
                    coll7_[(d_30_v0_) * (-89)] = False
            return _dafny.Map(coll7_)
        def iife8_():
            coll8_ = _dafny.Map()
            compr_8_: int
            for compr_8_ in (_dafny.Map({432: 305})).keys.Elements:
                d_31_v1_: int = compr_8_
                if (d_31_v1_) in (_dafny.Map({432: 305})):
                    coll8_[default__.safeModuloInt(d_31_v1_, 45)] = True
            return _dafny.Map(coll8_)
        def iife9_():
            coll9_ = _dafny.Map()
            compr_9_: int
            for compr_9_ in _dafny.IntegerRange(561, 572):
                d_32_v2_: int = compr_9_
                if ((561) <= (d_32_v2_)) and ((d_32_v2_) < (572)):
                    coll9_[default__.safeDivisionInt(d_32_v2_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gswko"))))] = False
            return _dafny.Map(coll9_)
        return (_dafny.Set({_dafny.Map({700: False}), _dafny.Map({-206: True}), iife7_()
        , iife8_()
        , _dafny.Map({836: not(False)})})).intersection(_dafny.Set({(_dafny.Map({len(_dafny.Set({False})): not(True)})), _dafny.Map({(0) - (len(_dafny.Set({-814}))): False}), _dafny.Map({len(_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tlws")))})): False}), _dafny.Map({(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([292]))])).cardinality: True}), iife9_()
        }))

    @staticmethod
    def fm32(p0, p1, p2, globalState):
        return (_dafny.Set({False, False})) - (_dafny.Set({True}))

    @staticmethod
    def fm33(p0, p1, globalState):
        return _dafny.MultiSet([(_dafny.CodePoint('h') if not(True) else _dafny.CodePoint('t'))])

    @staticmethod
    def fm34(globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([_dafny.CodePoint('y'), _dafny.CodePoint('q')]) for d_33_i0_ in range(default__.abs(561))])

    @staticmethod
    def fm35(p0, p1, globalState):
        def lambda0_(source1_):
            if source1_.is_DC5:
                d_34___mcc_h0_ = source1_.cf5
                d_35_cf5_ = d_34___mcc_h0_
                return (_dafny.Set({d_35_cf5_, d_35_cf5_, d_35_cf5_, d_35_cf5_})).ispropersubset(_dafny.Set({d_35_cf5_}))
            elif source1_.is_DC6:
                return (False) == (not(not(False)))
            elif source1_.is_DC7:
                d_36___mcc_h1_ = source1_.cf6
                d_37___mcc_h2_ = source1_.cf7
                d_38___mcc_h3_ = source1_.cf8
                d_39_cf8_ = d_38___mcc_h3_
                d_40_cf7_ = d_37___mcc_h2_
                d_41_cf6_ = d_36___mcc_h1_
                return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fucxg"))) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gwvrtg")))
            elif source1_.is_DC4:
                d_42___mcc_h4_ = source1_.cf4
                d_43_cf4_ = d_42___mcc_h4_
                return not(False)
            elif True:
                d_44___mcc_h5_ = source1_.cf9
                d_45_cf9_ = d_44___mcc_h5_
                return True

        return not(lambda0_(D2_DC8(D2_DC8(D2_DC4(_dafny.Set({len(_dafny.Map({False: len(_dafny.Set({_dafny.CodePoint('s'), _dafny.CodePoint('t')}))})), 870, len(_dafny.Map({-713: False}))}))))))

    @staticmethod
    def fm36(p0, p1, p2, globalState):
        return D8_DC23(_dafny.Map({_dafny.SeqWithoutIsStrInference([not(False)]): 379}))

    @staticmethod
    def fm37(p0, globalState):
        def iife10_():
            coll10_ = _dafny.Map()
            compr_10_: str
            for compr_10_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_46_i0_ in range(default__.abs(193))])).Elements:
                d_47_v0_: str = compr_10_
                if (d_47_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_46_i0_ in range(default__.abs(193))])):
                    coll10_[d_47_v0_] = not(False)
            return _dafny.Map(coll10_)
        return ((_dafny.Map({_dafny.CodePoint('s'): True})) | (iife10_()
        )) | (_dafny.Map({_dafny.CodePoint('g'): False}))

    @staticmethod
    def Main(noArgsParameter__):
        d_48_v0_: _dafny.Seq
        d_48_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rwdkj"))
        d_49_v1_: int
        d_49_v1_ = 143
        d_50_v2_: str
        d_50_v2_ = _dafny.CodePoint('w')
        d_51_v3_: _dafny.Array
        def lambda1_(d_52_i0_):
            return _dafny.CodePoint('n')

        init0_ = lambda1_
        nw0_ = _dafny.Array(None, 8)
        for i0_0_ in range(nw0_.length(0)):
            nw0_[i0_0_] = init0_(i0_0_)
        d_51_v3_ = nw0_
        d_53_globalState_: GlobalState
        nw1_ = GlobalState()
        nw1_.ctor__(287, (d_48_v0_) + ((d_48_v0_).set(default__.safeIndex((0) - (d_49_v1_), len(d_48_v0_)), d_50_v2_)), d_51_v3_)
        d_53_globalState_ = nw1_
        d_54_v4_: bool
        d_54_v4_ = True
        d_55_v5_: _dafny.Map
        d_55_v5_ = _dafny.Map({d_50_v2_: d_54_v4_})
        d_56_v6_: _dafny.Array
        nw2_ = _dafny.Array(False, 16)
        d_56_v6_ = nw2_
        d_57_v7_: _dafny.Set
        d_57_v7_ = _dafny.Set({_dafny.CodePoint('o'), d_50_v2_})
        d_58_v8_: _dafny.Seq
        d_58_v8_ = _dafny.SeqWithoutIsStrInference([d_48_v0_, (d_48_v0_).set(default__.safeIndex(len(d_57_v7_), len(d_48_v0_)), d_50_v2_)])
        d_59_v9_: _dafny.MultiSet
        d_59_v9_ = _dafny.MultiSet([False])
        d_60_v10_: C5
        nw3_ = C5()
        nw3_.ctor__(d_49_v1_, ((d_55_v5_)[d_50_v2_] if (d_50_v2_) in (d_55_v5_) else d_54_v4_), d_56_v6_, (d_54_v4_ if d_54_v4_ else d_54_v4_), (d_54_v4_ if default__.fm35(d_58_v8_, d_54_v4_, d_53_globalState_) else False), d_59_v9_)
        d_60_v10_ = nw3_
        d_54_v4_ = (d_60_v10_).f10
        d_61_i1_: int
        d_61_i1_ = 0
        with _dafny.label("0"):
            while (d_60_v10_).f10:
                with _dafny.c_label("0"):
                    if (d_61_i1_) >= (100):
                        raise _dafny.Break("0")
                    d_61_i1_ = (d_61_i1_) + (1)
                    index0_ = default__.safeIndex(609, (d_56_v6_).length(0))
                    (d_56_v6_)[index0_] = (d_60_v10_).f10
                    index1_ = default__.safeIndex(609, (d_56_v6_).length(0))
                    (d_56_v6_)[index1_] = d_54_v4_
                    if (d_60_v10_).f10:
                        d_62_v11_: _dafny.Map
                        d_62_v11_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cfhadr")): (d_60_v10_).f10})
                        index2_ = default__.safeIndex(609, (d_56_v6_).length(0))
                        rhs0_ = ((d_60_v10_).fm4(d_49_v1_, 364, d_62_v11_, (d_60_v10_).f9, d_53_globalState_)) - (589)
                        rhs1_ = (not(not(not((d_56_v6_)[default__.safeIndex(609, (d_56_v6_).length(0))]))) if (d_60_v10_).f10 else d_54_v4_)
                        lhs0_ = d_53_globalState_
                        lhs1_ = d_56_v6_
                        lhs2_ = default__.safeIndex(609, (d_56_v6_).length(0))
                        lhs0_.f0 = rhs0_
                        lhs1_[lhs2_] = rhs1_
                        d_63_v12_: _dafny.Map
                        d_63_v12_ = _dafny.Map({(d_60_v10_).f9: (d_56_v6_)[default__.safeIndex(609, (d_56_v6_).length(0))]})
                        (d_60_v10_).m1((d_60_v10_).f9, (d_60_v10_).f10, ((d_63_v12_)[(d_60_v10_).f9] if ((d_60_v10_).f9) in (d_63_v12_) else (d_60_v10_).f10), d_53_globalState_)
                        d_64_v13_: _dafny.Seq
                        d_64_v13_ = _dafny.SeqWithoutIsStrInference([d_49_v1_, (0) - ((d_60_v10_).f9)])
                        d_65_v14_: _dafny.Seq
                        d_65_v14_ = _dafny.SeqWithoutIsStrInference([(d_64_v13_)[default__.safeIndex(91, len(d_64_v13_))], len(d_48_v0_)])
                        (d_53_globalState_).f0 = (0) - (default__.safeModuloInt(len(d_64_v13_), (d_60_v10_).f9))
                        (d_53_globalState_).f0 = ((d_49_v1_) - ((d_60_v10_).f9) if True else d_49_v1_)
                        d_54_v4_ = (d_50_v2_) in (d_48_v0_)
                    elif True:
                        (d_53_globalState_).f0 = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qr")))
                        d_66_v15_: _dafny.Set
                        d_66_v15_ = _dafny.Set({d_49_v1_})
                        d_67_v16_: _dafny.Map
                        d_67_v16_ = _dafny.Map({(d_60_v10_).f9: (d_60_v10_).f9})
                        d_68_v17_: _dafny.Map
                        d_68_v17_ = _dafny.Map({False: default__.fm16((d_60_v10_).f10, d_67_v16_, _dafny.SeqWithoutIsStrInference([d_50_v2_ for d_69_i2_ in range(default__.abs(970))]), d_53_globalState_)})
                        d_70_v18_: D19
                        d_70_v18_ = D19_DC51(d_68_v17_)
                        d_71_v19_: _dafny.Map
                        d_71_v19_ = _dafny.Map({d_48_v0_: (d_60_v10_).f10})
                        d_72_v21_: _dafny.Seq
                        d_72_v21_ = _dafny.SeqWithoutIsStrInference([(d_60_v10_).f9])
                        d_73_v22_: _dafny.Seq
                        d_73_v22_ = _dafny.SeqWithoutIsStrInference([d_54_v4_])
                        d_74_v23_: _dafny.Array
                        nw4_ = _dafny.Array(None, 27)
                        nw4_[int(0)] = (d_60_v10_).f9
                        nw4_[int(1)] = (d_60_v10_).f9
                        nw4_[int(2)] = (d_60_v10_).f9
                        nw4_[int(3)] = (d_60_v10_).f9
                        nw4_[int(4)] = len(d_66_v15_)
                        nw4_[int(5)] = (d_60_v10_).f9
                        nw4_[int(6)] = (d_60_v10_).fm4(825, (0) - (len((d_70_v18_).cf81)), d_71_v19_, (d_60_v10_).f9, d_53_globalState_)
                        nw4_[int(7)] = d_49_v1_
                        nw4_[int(8)] = len(d_48_v0_)
                        def iife11_():
                            coll11_ = _dafny.Map()
                            compr_11_: int
                            for compr_11_ in _dafny.IntegerRange(473, 327):
                                d_75_v20_: int = compr_11_
                                if ((473) <= (d_75_v20_)) and ((d_75_v20_) < (327)):
                                    coll11_[default__.safeDivisionInt(d_75_v20_, (d_60_v10_).f9)] = d_50_v2_
                            return _dafny.Map(coll11_)
                        nw4_[int(9)] = len(iife11_()
                        )
                        nw4_[int(10)] = (d_60_v10_).f9
                        nw4_[int(11)] = (d_60_v10_).f9
                        nw4_[int(12)] = d_49_v1_
                        nw4_[int(13)] = len(_dafny.SeqWithoutIsStrInference([d_50_v2_ for d_76_i3_ in range(default__.abs(-561))]))
                        nw4_[int(14)] = (d_60_v10_).f9
                        nw4_[int(15)] = len(d_72_v21_)
                        nw4_[int(16)] = (d_60_v10_).f9
                        nw4_[int(17)] = (d_60_v10_).f9
                        nw4_[int(18)] = (d_60_v10_).f9
                        nw4_[int(19)] = 307
                        nw4_[int(20)] = (d_60_v10_).f9
                        nw4_[int(21)] = (d_60_v10_).f9
                        nw4_[int(22)] = (d_60_v10_).f9
                        nw4_[int(23)] = len(d_48_v0_)
                        nw4_[int(24)] = (d_60_v10_).fm4((d_60_v10_).f9, len(d_73_v22_), d_71_v19_, (d_60_v10_).f9, d_53_globalState_)
                        nw4_[int(25)] = (d_60_v10_).f9
                        nw4_[int(26)] = len(d_73_v22_)
                        d_74_v23_ = nw4_
                        d_77_v24_: _dafny.Seq
                        d_77_v24_ = _dafny.SeqWithoutIsStrInference([d_74_v23_])
                        d_78_v25_: _dafny.Set
                        d_78_v25_ = _dafny.Set({(d_60_v10_).f10})
                        d_79_v26_: _dafny.MultiSet
                        d_79_v26_ = _dafny.MultiSet([len(_dafny.Map({(d_56_v6_)[default__.safeIndex(609, (d_56_v6_).length(0))]: (d_60_v10_).f10})), (0) - ((d_60_v10_).f9), (d_60_v10_).f9])
                        d_80_v28_: _dafny.Set
                        d_80_v28_ = _dafny.Set({d_48_v0_})
                        d_81_v29_: _dafny.MultiSet
                        d_81_v29_ = _dafny.MultiSet([_dafny.CodePoint('d'), d_50_v2_, d_50_v2_])
                        d_82_v30_: _dafny.Array
                        nw5_ = _dafny.Array(None, 21)
                        nw5_[int(0)] = len((d_77_v24_) + (d_77_v24_))
                        nw5_[int(1)] = default__.safeModuloInt(len(d_78_v25_), (d_60_v10_).f9)
                        nw5_[int(2)] = (d_60_v10_).f9
                        nw5_[int(3)] = 733
                        nw5_[int(4)] = d_49_v1_
                        nw5_[int(5)] = (d_79_v26_).cardinality
                        nw5_[int(6)] = (d_60_v10_).f9
                        nw5_[int(7)] = d_49_v1_
                        nw5_[int(8)] = d_49_v1_
                        nw5_[int(9)] = ((d_60_v10_).f9) + (len(d_66_v15_))
                        def iife12_():
                            coll12_ = _dafny.Map()
                            compr_12_: _dafny.Seq
                            for compr_12_ in (d_80_v28_).Elements:
                                d_83_v27_: _dafny.Seq = compr_12_
                                if (d_83_v27_) in (d_80_v28_):
                                    coll12_[d_83_v27_] = (d_60_v10_).f10
                            return _dafny.Map(coll12_)
                        nw5_[int(10)] = (d_60_v10_).fm4((d_60_v10_).f9, 520, (iife12_()
                        ).set(d_48_v0_, d_54_v4_), (d_81_v29_).cardinality, d_53_globalState_)
                        nw5_[int(11)] = 588
                        nw5_[int(12)] = d_49_v1_
                        nw5_[int(13)] = (d_60_v10_).f9
                        nw5_[int(14)] = ((d_60_v10_).f9 if (d_56_v6_)[default__.safeIndex(609, (d_56_v6_).length(0))] else (d_60_v10_).fm4(len(d_68_v17_), (0) - ((d_60_v10_).f9), d_71_v19_, d_49_v1_, d_53_globalState_))
                        nw5_[int(15)] = ((d_60_v10_).f9) + ((d_60_v10_).f9)
                        nw5_[int(16)] = (d_60_v10_).f9
                        nw5_[int(17)] = (d_60_v10_).f9
                        nw5_[int(18)] = ((d_60_v10_).f9) + (len(d_72_v21_))
                        nw5_[int(19)] = d_49_v1_
                        nw5_[int(20)] = len(d_48_v0_)
                        d_82_v30_ = nw5_
                        index3_ = default__.safeIndex(668, (d_82_v30_).length(0))
                        (d_82_v30_)[index3_] = (d_60_v10_).f9
                        d_84_v31_: D19
                        d_84_v31_ = D19_DC52(len(_dafny.Map({not((d_60_v10_).f10): d_49_v1_})))
                        index4_ = default__.safeIndex(668, (d_82_v30_).length(0))
                        def iife13_():
                            coll13_ = _dafny.Map()
                            compr_13_: int
                            for compr_13_ in _dafny.IntegerRange(779, 783):
                                d_85_v32_: int = compr_13_
                                if ((779) <= (d_85_v32_)) and ((d_85_v32_) < (783)):
                                    coll13_[(d_85_v32_) * (742)] = d_54_v4_
                            return _dafny.Map(coll13_)
                        (d_82_v30_)[index4_] = default__.safeDivisionInt((d_84_v31_).cf82, len((_dafny.Map({(d_60_v10_).f9: True})) | (iife13_()
                        )))
                        index5_ = default__.safeIndex(668, (d_82_v30_).length(0))
                        (d_82_v30_)[index5_] = d_49_v1_
                        d_54_v4_ = (d_60_v10_).f10
                        d_86_v33_: _dafny.Array
                        nw6_ = _dafny.Array(False, 3)
                        d_86_v33_ = nw6_
                        d_87_v34_: C4
                        nw7_ = C4()
                        nw7_.ctor__(d_86_v33_, d_54_v4_, False, d_59_v9_)
                        d_87_v34_ = nw7_
                        d_88_v35_: _dafny.Map
                        d_88_v35_ = _dafny.Map({d_54_v4_: d_87_v34_})
                        (d_53_globalState_).f0 = (d_60_v10_).fm4(len(d_88_v35_), (d_60_v10_).f9, d_71_v19_, len(d_48_v0_), d_53_globalState_)
                    d_89_v36_: _dafny.Map
                    d_89_v36_ = _dafny.Map({d_59_v9_: d_48_v0_})
                    d_90_v37_: D16
                    d_90_v37_ = D16_DC44(d_54_v4_, (d_56_v6_)[default__.safeIndex(609, (d_56_v6_).length(0))], len(d_89_v36_))
                    d_91_v38_: _dafny.Map
                    d_91_v38_ = _dafny.Map({d_90_v37_: d_49_v1_})
                    d_92_v39_: _dafny.Map
                    d_92_v39_ = _dafny.Map({d_49_v1_: (d_60_v10_).f10})
                    d_93_v40_: _dafny.Map
                    d_93_v40_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rwq")): d_54_v4_})
                    d_94_v41_: _dafny.Seq
                    d_94_v41_ = _dafny.SeqWithoutIsStrInference([((d_92_v39_)[(d_60_v10_).f9] if ((d_60_v10_).f9) in (d_92_v39_) else (d_60_v10_).f10), (d_60_v10_).f10, ((d_92_v39_)[(d_60_v10_).fm4((d_60_v10_).f9, (d_60_v10_).f9, d_93_v40_, (d_60_v10_).f9, d_53_globalState_)] if ((d_60_v10_).fm4((d_60_v10_).f9, (d_60_v10_).f9, d_93_v40_, (d_60_v10_).f9, d_53_globalState_)) in (d_92_v39_) else not(True)), (d_60_v10_).f10])
                    (d_60_v10_).m1(((d_91_v38_)[d_90_v37_] if (d_90_v37_) in (d_91_v38_) else (d_60_v10_).f9), (d_56_v6_)[default__.safeIndex(609, (d_56_v6_).length(0))], (d_94_v41_)[default__.safeIndex((d_60_v10_).f9, len(d_94_v41_))], d_53_globalState_)
                    d_95_v42_: _dafny.Seq
                    d_95_v42_ = _dafny.SeqWithoutIsStrInference([(d_60_v10_).f9])
                    d_96_v43_: T1
                    nw8_ = C1()
                    nw8_.ctor__((d_60_v10_).f10, d_48_v0_, d_54_v4_, d_59_v9_)
                    d_96_v43_ = nw8_
                    d_97_v44_: _dafny.Map
                    d_97_v44_ = _dafny.Map({d_96_v43_: (d_60_v10_).f10})
                    d_98_v45_: _dafny.Seq
                    d_98_v45_ = _dafny.SeqWithoutIsStrInference([(d_94_v41_).set(default__.safeIndex(-771, len(d_94_v41_)), not((d_96_v43_).f6))])
                    index6_ = default__.safeIndex(609, (d_56_v6_).length(0))
                    index7_ = default__.safeIndex(609, (d_56_v6_).length(0))
                    rhs2_ = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_99_i4_ in range(default__.abs(-143))])) + (d_48_v0_)
                    rhs3_ = ((d_95_v42_) + (_dafny.SeqWithoutIsStrInference([(d_60_v10_).f9, len(d_97_v44_), d_49_v1_, (0) - ((d_60_v10_).f9)]))).set(default__.safeIndex(default__.safeModuloInt(len(d_94_v41_), (d_60_v10_).f9), len((d_95_v42_) + (_dafny.SeqWithoutIsStrInference([(d_60_v10_).f9, len(d_97_v44_), d_49_v1_, (0) - ((d_60_v10_).f9)])))), (0) - (-259))
                    rhs4_ = (d_54_v4_) or ((d_60_v10_).f10)
                    rhs5_ = ((d_94_v41_) + (d_94_v41_)) not in (d_98_v45_)
                    rhs6_ = ((d_94_v41_) + (_dafny.SeqWithoutIsStrInference([(d_96_v43_).f6, (d_60_v10_).f10]))) + ((d_94_v41_) + (d_94_v41_))
                    lhs3_ = d_56_v6_
                    lhs4_ = default__.safeIndex(609, (d_56_v6_).length(0))
                    lhs5_ = d_56_v6_
                    lhs6_ = default__.safeIndex(609, (d_56_v6_).length(0))
                    d_48_v0_ = rhs2_
                    d_95_v42_ = rhs3_
                    lhs3_[lhs4_] = rhs4_
                    lhs5_[lhs6_] = rhs5_
                    d_94_v41_ = rhs6_
                    pass
            pass
        d_49_v1_ = d_49_v1_
        d_100_v46_: _dafny.Map
        d_100_v46_ = _dafny.Map({True: d_48_v0_})
        d_100_v46_ = d_100_v46_
        d_101_v47_: _dafny.Map
        d_101_v47_ = _dafny.Map({_dafny.Map({-434: (d_60_v10_).f9}): d_54_v4_})
        d_102_v48_: C0
        nw9_ = C0()
        nw9_.ctor__(d_101_v47_)
        d_102_v48_ = nw9_
        d_103_v49_: _dafny.Array
        def lambda2_(d_104_v1_):
            def lambda3_(d_105_i5_):
                return (d_105_i5_) + (d_104_v1_)

            return lambda3_

        init1_ = lambda2_(d_49_v1_)
        nw10_ = _dafny.Array(None, 12)
        for i0_1_ in range(nw10_.length(0)):
            nw10_[i0_1_] = init1_(i0_1_)
        d_103_v49_ = nw10_
        index8_ = default__.safeIndex(431, (d_103_v49_).length(0))
        (d_103_v49_)[index8_] = (d_60_v10_).f9
        d_106_v51_: _dafny.Array
        def lambda4_(d_107_v10_, d_108_v4_):
            def lambda5_(d_109_i6_):
                def iife14_():
                    coll14_ = _dafny.Map()
                    compr_14_: D8
                    for compr_14_ in (_dafny.MultiSet([D8_DC24((d_107_v10_).f9, d_108_v4_, _dafny.SeqWithoutIsStrInference([(d_107_v10_).f10]), 959, (d_107_v10_).f9)])).Elements:
                        d_110_v50_: D8 = compr_14_
                        if (d_110_v50_) in (_dafny.MultiSet([D8_DC24((d_107_v10_).f9, d_108_v4_, _dafny.SeqWithoutIsStrInference([(d_107_v10_).f10]), 959, (d_107_v10_).f9)])):
                            coll14_[d_110_v50_] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v'), _dafny.CodePoint('i')])
                    return _dafny.Map(coll14_)
                return iife14_()
                

            return lambda5_

        init2_ = lambda4_(d_60_v10_, d_54_v4_)
        nw11_ = _dafny.Array(None, 20)
        for i0_2_ in range(nw11_.length(0)):
            nw11_[i0_2_] = init2_(i0_2_)
        d_106_v51_ = nw11_
        d_111_v52_: _dafny.Seq
        d_111_v52_ = _dafny.SeqWithoutIsStrInference([(d_60_v10_).f10, (d_60_v10_).f10, (d_60_v10_).fm1((d_60_v10_).f10, -253, (d_60_v10_).f9, d_53_globalState_)])
        d_112_v53_: D8
        d_112_v53_ = D8_DC24((_dafny.MultiSet(d_111_v52_)).cardinality, (d_60_v10_).f10, d_111_v52_, (d_60_v10_).f9, d_49_v1_)
        d_113_v54_: _dafny.Map
        d_113_v54_ = _dafny.Map({d_49_v1_: (d_60_v10_).f9})
        d_114_v55_: _dafny.Map
        d_114_v55_ = _dafny.Map({d_112_v53_: _dafny.SeqWithoutIsStrInference([d_50_v2_, d_50_v2_, default__.fm16(d_54_v4_, d_113_v54_, d_48_v0_, d_53_globalState_), d_50_v2_])})
        index9_ = default__.safeIndex(25, (d_106_v51_).length(0))
        (d_106_v51_)[index9_] = d_114_v55_
        d_115_v57_: _dafny.Seq
        d_115_v57_ = _dafny.SeqWithoutIsStrInference([d_112_v53_, d_112_v53_])
        d_116_v58_: _dafny.Seq
        def iife15_():
            coll15_ = _dafny.Map()
            compr_15_: D8
            for compr_15_ in (d_115_v57_).Elements:
                d_117_v56_: D8 = compr_15_
                if (d_117_v56_) in (d_115_v57_):
                    coll15_[d_117_v56_] = d_48_v0_
            return _dafny.Map(coll15_)
        d_116_v58_ = _dafny.SeqWithoutIsStrInference([iife15_()
        ])
        index10_ = default__.safeIndex(431, (d_103_v49_).length(0))
        index11_ = default__.safeIndex(25, (d_106_v51_).length(0))
        rhs7_ = d_102_v48_
        rhs8_ = default__.safeModuloInt((0) - ((d_60_v10_).f9), (d_60_v10_).f9)
        rhs9_ = ((d_116_v58_)[default__.safeIndex(d_49_v1_, len(d_116_v58_))]) | (d_114_v55_)
        rhs10_ = (d_60_v10_).f9
        rhs11_ = (d_54_v4_ if (d_60_v10_).f10 else d_54_v4_)
        lhs7_ = d_103_v49_
        lhs8_ = default__.safeIndex(431, (d_103_v49_).length(0))
        lhs9_ = d_106_v51_
        lhs10_ = default__.safeIndex(25, (d_106_v51_).length(0))
        lhs11_ = d_53_globalState_
        d_102_v48_ = rhs7_
        lhs7_[lhs8_] = rhs8_
        lhs9_[lhs10_] = rhs9_
        lhs11_.f0 = rhs10_
        d_54_v4_ = rhs11_
        d_118_v60_: _dafny.Set
        d_118_v60_ = _dafny.Set({(d_103_v49_)[default__.safeIndex(431, (d_103_v49_).length(0))]})
        def iife16_():
            coll16_ = _dafny.Map()
            compr_16_: int
            for compr_16_ in (d_113_v54_).keys.Elements:
                d_119_v59_: int = compr_16_
                if (d_119_v59_) in (d_113_v54_):
                    coll16_[(d_119_v59_) * (-411)] = (d_60_v10_).f9
            return _dafny.Map(coll16_)
        source2_ = default__.fm36((d_113_v54_) | (iife16_()
        ), d_54_v4_, d_118_v60_, d_53_globalState_)
        if source2_.is_DC24:
            d_120___mcc_h0_ = source2_.cf34
            d_121___mcc_h1_ = source2_.cf35
            d_122___mcc_h2_ = source2_.cf36
            d_123___mcc_h3_ = source2_.cf37
            d_124___mcc_h4_ = source2_.cf38
            d_125_cf38_ = d_124___mcc_h4_
            d_126_cf37_ = d_123___mcc_h3_
            d_127_cf36_ = d_122___mcc_h2_
            d_128_cf35_ = d_121___mcc_h1_
            d_129_cf34_ = d_120___mcc_h0_
            index12_ = default__.safeIndex(431, (d_103_v49_).length(0))
            (d_103_v49_)[index12_] = ((0) - ((d_125_cf38_) * (d_126_cf37_))) - ((d_103_v49_)[default__.safeIndex(431, (d_103_v49_).length(0))])
            if (d_60_v10_).f10:
                index13_ = default__.safeIndex(431, (d_103_v49_).length(0))
                (d_103_v49_)[index13_] = 59
                d_130_v61_: _dafny.Map
                d_130_v61_ = _dafny.Map({(d_60_v10_).f9: d_128_cf35_})
                d_54_v4_ = ((d_130_v61_)[d_126_cf37_] if (d_126_cf37_) in (d_130_v61_) else (d_60_v10_).f10)
                d_125_cf38_ = (d_60_v10_).f9
                index14_ = default__.safeIndex(621, (d_56_v6_).length(0))
                (d_56_v6_)[index14_] = d_54_v4_
                d_131_v62_: _dafny.Map
                d_131_v62_ = _dafny.Map({d_48_v0_: not((d_60_v10_).f10)})
                d_132_v63_: _dafny.Map
                d_132_v63_ = _dafny.Map({d_128_cf35_: d_54_v4_})
                index15_ = default__.safeIndex(621, (d_56_v6_).length(0))
                (d_56_v6_)[index15_] = (d_60_v10_).fm1(not ((d_60_v10_).f10) or (d_54_v4_), default__.safeDivisionInt((d_60_v10_).fm4((d_60_v10_).f9, (d_60_v10_).f9, d_131_v62_, (d_60_v10_).f9, d_53_globalState_), len((d_132_v63_).set((d_60_v10_).fm2((d_60_v10_).f10, d_53_globalState_), (d_60_v10_).f10))), (d_103_v49_)[default__.safeIndex(431, (d_103_v49_).length(0))], d_53_globalState_)
                d_48_v0_ = _dafny.SeqWithoutIsStrInference([d_50_v2_ for d_133_i7_ in range(default__.abs(187))])
            elif True:
                (d_60_v10_).m2(d_50_v2_, d_53_globalState_)
                d_54_v4_ = not(d_128_cf35_)
                d_48_v0_ = d_48_v0_
                d_134_v64_: _dafny.Array
                nw12_ = _dafny.Array(_dafny.Array(None, 0), 15)
                d_134_v64_ = nw12_
                d_135_v65_: C2
                nw13_ = C2()
                nw13_.ctor__((d_59_v9_))
                d_135_v65_ = nw13_
                d_136_v66_: D9
                d_136_v66_ = D9_DC26(d_135_v65_)
                d_137_v67_: _dafny.MultiSet
                d_137_v67_ = _dafny.MultiSet([(d_60_v10_).f9])
                d_138_v68_: _dafny.Set
                d_138_v68_ = _dafny.Set({(d_60_v10_).f10, (d_60_v10_).f10})
                rhs12_ = (d_127_cf36_) <= (default__.fm23(((d_137_v67_)[(0) - (len(d_138_v68_))] if ((0) - (len(d_138_v68_))) in (d_137_v67_) else d_129_cf34_), d_53_globalState_))
                rhs13_ = d_134_v64_
                rhs14_ = d_50_v2_
                rhs15_ = d_136_v66_
                d_128_cf35_ = rhs12_
                d_134_v64_ = rhs13_
                d_50_v2_ = rhs14_
                d_136_v66_ = rhs15_
                index16_ = default__.safeIndex(438, (d_56_v6_).length(0))
                (d_56_v6_)[index16_] = d_54_v4_
                index17_ = default__.safeIndex(438, (d_56_v6_).length(0))
                (d_56_v6_)[index17_] = (d_60_v10_).f10
            d_139_v69_: _dafny.Map
            d_139_v69_ = _dafny.Map({d_54_v4_: _dafny.Set({d_129_cf34_})})
            d_140_v70_: _dafny.Set
            d_140_v70_ = _dafny.Set({((d_139_v69_)[d_54_v4_] if (d_54_v4_) in (d_139_v69_) else d_118_v60_), default__.fm24(len(d_48_v0_), True, d_49_v1_, d_53_globalState_)})
            def iife17_():
                coll17_ = _dafny.Set()
                compr_17_: _dafny.Set
                for compr_17_ in (d_140_v70_).Elements:
                    d_141_v71_: _dafny.Set = compr_17_
                    if (d_141_v71_) in (d_140_v70_):
                        coll17_ = coll17_.union(_dafny.Set([d_141_v71_]))
                return _dafny.Set(coll17_)
            d_140_v70_ = iife17_()
            
            d_142_v72_: _dafny.Array
            nw14_ = _dafny.Array(_dafny.Array(None, 0), 24)
            d_142_v72_ = nw14_
            index18_ = default__.safeIndex(99, (d_142_v72_).length(0))
            (d_142_v72_)[index18_] = d_51_v3_
            index19_ = default__.safeIndex(99, (d_142_v72_).length(0))
            (d_142_v72_)[index19_] = d_51_v3_
        elif source2_.is_DC25:
            d_143___mcc_h5_ = source2_.cf39
            d_144___mcc_h6_ = source2_.cf40
            d_145_cf40_ = d_144___mcc_h6_
            d_146_cf39_ = d_143___mcc_h5_
            d_147_v73_: _dafny.Map
            d_147_v73_ = _dafny.Map({(d_60_v10_).f10: d_103_v49_})
            d_147_v73_ = d_147_v73_
            rhs16_ = (d_60_v10_).f9
            rhs17_ = d_145_cf40_
            rhs18_ = not ((d_60_v10_).f10) or ((d_145_cf40_) in ((_dafny.Map({(d_60_v10_).f10: d_103_v49_})).set(False, d_103_v49_)))
            rhs19_ = d_48_v0_
            rhs20_ = True
            lhs12_ = d_53_globalState_
            lhs12_.f0 = rhs16_
            d_54_v4_ = rhs17_
            d_146_cf39_ = rhs18_
            d_48_v0_ = rhs19_
            d_146_cf39_ = rhs20_
            d_50_v2_ = _dafny.CodePoint('d')
            (d_53_globalState_).f0 = d_49_v1_
        elif True:
            d_148___mcc_h7_ = source2_.cf33
            d_149_cf33_ = d_148___mcc_h7_
            if (d_60_v10_).f10:
                (d_53_globalState_).f0 = (d_60_v10_).f9
                index20_ = default__.safeIndex(75, (d_56_v6_).length(0))
                (d_56_v6_)[index20_] = not((d_60_v10_).f10)
                index21_ = default__.safeIndex(75, (d_56_v6_).length(0))
                (d_56_v6_)[index21_] = not((d_60_v10_).f10)
                d_150_v74_: _dafny.Map
                d_150_v74_ = _dafny.Map({d_54_v4_: (d_56_v6_)[default__.safeIndex(75, (d_56_v6_).length(0))]})
                d_151_v75_: _dafny.Map
                d_151_v75_ = _dafny.Map({_dafny.MultiSet([d_54_v4_, (d_60_v10_).f10]): d_150_v74_})
                d_152_v77_: _dafny.Map
                d_152_v77_ = _dafny.Map({_dafny.MultiSet([d_54_v4_, not(d_54_v4_), (d_60_v10_).fm1((d_60_v10_).f10, len(d_48_v0_), len(d_111_v52_), d_53_globalState_), (d_60_v10_).f10]): len(d_150_v74_)})
                d_153_v79_: _dafny.Set
                d_153_v79_ = _dafny.Set({d_59_v9_})
                def iife18_():
                    coll18_ = _dafny.Map()
                    compr_18_: _dafny.MultiSet
                    for compr_18_ in (d_152_v77_).keys.Elements:
                        d_154_v76_: _dafny.MultiSet = compr_18_
                        if (d_154_v76_) in (d_152_v77_):
                            coll18_[d_154_v76_] = d_150_v74_
                    return _dafny.Map(coll18_)
                def iife19_():
                    coll19_ = _dafny.Map()
                    compr_19_: _dafny.MultiSet
                    for compr_19_ in (d_153_v79_).Elements:
                        d_155_v78_: _dafny.MultiSet = compr_19_
                        if (d_155_v78_) in (d_153_v79_):
                            coll19_[d_155_v78_] = d_150_v74_
                    return _dafny.Map(coll19_)
                d_151_v75_ = (d_151_v75_) | ((iife18_()
                ) | (iife19_()
                ))
                (d_53_globalState_).f0 = (d_60_v10_).f9
                d_156_v80_: _dafny.Array
                def lambda6_(d_157_v10_):
                    def lambda7_(d_158_i8_):
                        return _dafny.Map({(d_157_v10_).f10: -120})

                    return lambda7_

                init3_ = lambda6_(d_60_v10_)
                nw15_ = _dafny.Array(None, 19)
                for i0_3_ in range(nw15_.length(0)):
                    nw15_[i0_3_] = init3_(i0_3_)
                d_156_v80_ = nw15_
                d_159_v81_: bool
                d_160_v82_: int
                d_161_v83_: bool
                out0_: bool
                out1_: int
                out2_: bool
                out0_, out1_, out2_ = (d_60_v10_).m5(-675, d_59_v9_, d_156_v80_, d_54_v4_, d_53_globalState_)
                d_159_v81_ = out0_
                d_160_v82_ = out1_
                d_161_v83_ = out2_
            elif True:
                d_162_v84_: _dafny.Seq
                d_162_v84_ = _dafny.SeqWithoutIsStrInference([d_49_v1_])
                (d_60_v10_).m4((d_60_v10_).fm2((d_60_v10_).f10, d_53_globalState_), d_162_v84_, d_53_globalState_)
                d_163_v85_: _dafny.Array
                nw16_ = _dafny.Array(_dafny.Array(None, 0), 12)
                d_163_v85_ = nw16_
                index22_ = default__.safeIndex(677, (d_163_v85_).length(0))
                (d_163_v85_)[index22_] = d_56_v6_
                index23_ = default__.safeIndex(677, (d_163_v85_).length(0))
                (d_163_v85_)[index23_] = d_56_v6_
                d_164_v86_: _dafny.Seq
                d_164_v86_ = _dafny.SeqWithoutIsStrInference([d_102_v48_, d_102_v48_])
                d_165_v87_: D2
                d_165_v87_ = D2_DC7(d_50_v2_, d_113_v54_, (d_60_v10_).f9)
                d_166_v88_: _dafny.Map
                d_166_v88_ = _dafny.Map({d_48_v0_: d_54_v4_})
                d_167_v89_: _dafny.Array
                nw17_ = _dafny.Array(None, 28)
                nw17_[int(0)] = d_102_v48_
                nw17_[int(1)] = d_102_v48_
                nw17_[int(2)] = d_102_v48_
                nw17_[int(3)] = d_102_v48_
                nw17_[int(4)] = d_102_v48_
                nw17_[int(5)] = d_102_v48_
                nw17_[int(6)] = d_102_v48_
                nw17_[int(7)] = d_102_v48_
                nw17_[int(8)] = (d_164_v86_)[default__.safeIndex((d_60_v10_).fm4(d_49_v1_, (d_165_v87_).cf8, d_166_v88_, d_49_v1_, d_53_globalState_), len(d_164_v86_))]
                nw17_[int(9)] = d_102_v48_
                nw17_[int(10)] = d_102_v48_
                nw17_[int(11)] = d_102_v48_
                nw17_[int(12)] = d_102_v48_
                nw17_[int(13)] = d_102_v48_
                nw17_[int(14)] = d_102_v48_
                nw17_[int(15)] = d_102_v48_
                nw17_[int(16)] = d_102_v48_
                nw17_[int(17)] = d_102_v48_
                nw17_[int(18)] = d_102_v48_
                nw17_[int(19)] = d_102_v48_
                nw17_[int(20)] = d_102_v48_
                nw17_[int(21)] = d_102_v48_
                nw17_[int(22)] = d_102_v48_
                nw17_[int(23)] = d_102_v48_
                nw17_[int(24)] = d_102_v48_
                nw17_[int(25)] = d_102_v48_
                nw17_[int(26)] = d_102_v48_
                nw17_[int(27)] = d_102_v48_
                d_167_v89_ = nw17_
                index24_ = default__.safeIndex(708, (d_167_v89_).length(0))
                (d_167_v89_)[index24_] = d_102_v48_
                index25_ = default__.safeIndex(708, (d_167_v89_).length(0))
                rhs21_ = not((381) < ((d_103_v49_)[default__.safeIndex(431, (d_103_v49_).length(0))]))
                rhs22_ = (d_60_v10_).f9
                rhs23_ = d_102_v48_
                lhs13_ = d_53_globalState_
                lhs14_ = d_167_v89_
                lhs15_ = default__.safeIndex(708, (d_167_v89_).length(0))
                d_54_v4_ = rhs21_
                lhs13_.f0 = rhs22_
                lhs14_[lhs15_] = rhs23_
                index26_ = default__.safeIndex(431, (d_103_v49_).length(0))
                (d_103_v49_)[index26_] = (d_103_v49_)[default__.safeIndex(431, (d_103_v49_).length(0))]
                d_168_v90_: T0
                nw18_ = C2()
                nw18_.ctor__(d_59_v9_)
                d_168_v90_ = nw18_
                rhs24_ = (d_60_v10_).f10
                rhs25_ = d_168_v90_
                d_54_v4_ = rhs24_
                d_168_v90_ = rhs25_
            d_169_v91_: _dafny.Seq
            d_169_v91_ = _dafny.SeqWithoutIsStrInference([d_49_v1_])
            d_170_v92_: D15
            d_170_v92_ = D15_DC42(_dafny.MultiSet(d_169_v91_))
            d_171_v93_: _dafny.Seq
            d_171_v93_ = _dafny.SeqWithoutIsStrInference([D15_DC42(_dafny.MultiSet(d_169_v91_)), d_170_v92_])
            d_172_v94_: D21
            d_172_v94_ = D21_DC54(d_171_v93_)
            d_173_v95_: _dafny.MultiSet
            d_173_v95_ = _dafny.MultiSet([(d_103_v49_)[default__.safeIndex(431, (d_103_v49_).length(0))], (d_103_v49_)[default__.safeIndex(431, (d_103_v49_).length(0))], (d_103_v49_)[default__.safeIndex(431, (d_103_v49_).length(0))], (d_60_v10_).f9])
            if (d_49_v1_) > (len(((d_172_v94_).cf84) + (_dafny.SeqWithoutIsStrInference([D15_DC42(d_173_v95_), D15_DC42(d_173_v95_), d_170_v92_])))):
                d_174_v96_: C2
                nw19_ = C2()
                nw19_.ctor__(d_59_v9_)
                d_174_v96_ = nw19_
                d_175_v97_: _dafny.Map
                d_175_v97_ = _dafny.Map({d_174_v96_: (d_103_v49_)[default__.safeIndex(431, (d_103_v49_).length(0))]})
                d_176_v98_: _dafny.Seq
                d_176_v98_ = _dafny.SeqWithoutIsStrInference([d_56_v6_])
                d_177_v99_: _dafny.Map
                d_177_v99_ = _dafny.Map({d_176_v98_: (d_60_v10_).f10})
                (d_60_v10_).m1(len(d_175_v97_), ((d_177_v99_)[_dafny.SeqWithoutIsStrInference([d_56_v6_, d_56_v6_])] if (_dafny.SeqWithoutIsStrInference([d_56_v6_, d_56_v6_])) in (d_177_v99_) else (d_60_v10_).f10), (d_60_v10_).f10, d_53_globalState_)
                d_178_v100_: _dafny.Array
                nw20_ = _dafny.Array(None, 11)
                d_178_v100_ = nw20_
                index27_ = default__.safeIndex(111, (d_178_v100_).length(0))
                (d_178_v100_)[index27_] = d_60_v10_
                index28_ = default__.safeIndex(111, (d_178_v100_).length(0))
                (d_178_v100_)[index28_] = d_60_v10_
                index29_ = default__.safeIndex(431, (d_103_v49_).length(0))
                (d_103_v49_)[index29_] = (569) + (len(d_111_v52_))
                d_179_v101_: _dafny.Map
                d_179_v101_ = _dafny.Map({(d_60_v10_).f10: (d_60_v10_).f10})
                d_180_v102_: D13
                d_180_v102_ = D13_DC39((_dafny.Map({d_50_v2_: d_54_v4_})) | (d_55_v5_), d_56_v6_, d_111_v52_, len(d_179_v101_), (d_169_v91_) + (d_169_v91_))
                d_180_v102_ = D13_DC39(default__.fm37(d_50_v2_, d_53_globalState_), d_56_v6_, d_111_v52_, ((d_60_v10_).f9) - (134), (d_169_v91_ if True else default__.fm18(d_53_globalState_)))
                d_54_v4_ = (not ((d_60_v10_).f10) or (d_54_v4_)) in (_dafny.MultiSet(d_111_v52_))
            elif True:
                d_50_v2_ = d_50_v2_
                index30_ = default__.safeIndex(431, (d_103_v49_).length(0))
                (d_103_v49_)[index30_] = d_49_v1_
                index31_ = default__.safeIndex(8, (d_56_v6_).length(0))
                (d_56_v6_)[index31_] = (d_60_v10_).f10
                index32_ = default__.safeIndex(8, (d_56_v6_).length(0))
                (d_56_v6_)[index32_] = (d_48_v0_) <= ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "btussn"))))
                d_181_v103_: C6
                nw21_ = C6()
                nw21_.ctor__(d_118_v60_, d_50_v2_)
                d_181_v103_ = nw21_
                d_182_v104_: _dafny.Map
                d_182_v104_ = _dafny.Map({d_181_v103_: (d_56_v6_)[default__.safeIndex(8, (d_56_v6_).length(0))]})
                d_182_v104_ = (d_182_v104_).set((d_181_v103_ if d_54_v4_ else d_181_v103_), (d_60_v10_).f10)
                index33_ = default__.safeIndex(431, (d_103_v49_).length(0))
                index34_ = default__.safeIndex(8, (d_56_v6_).length(0))
                rhs26_ = (d_60_v10_).f9
                rhs27_ = (d_60_v10_).f10
                lhs16_ = d_103_v49_
                lhs17_ = default__.safeIndex(431, (d_103_v49_).length(0))
                lhs18_ = d_56_v6_
                lhs19_ = default__.safeIndex(8, (d_56_v6_).length(0))
                lhs16_[lhs17_] = rhs26_
                lhs18_[lhs19_] = rhs27_
            d_183_v105_: C3
            nw22_ = C3()
            nw22_.ctor__((d_111_v52_) <= (d_111_v52_), d_59_v9_)
            d_183_v105_ = nw22_
            d_54_v4_ = (d_60_v10_).f10
        d_184_v106_: _dafny.Array
        def lambda8_(d_185_v52_):
            def lambda9_(d_186_i9_):
                return D6_DC17(d_185_v52_)

            return lambda9_

        init4_ = lambda8_(d_111_v52_)
        nw23_ = _dafny.Array(None, 3)
        for i0_4_ in range(nw23_.length(0)):
            nw23_[i0_4_] = init4_(i0_4_)
        d_184_v106_ = nw23_
        index35_ = default__.safeIndex(622, (d_184_v106_).length(0))
        (d_184_v106_)[index35_] = D6_DC17(d_111_v52_)
        d_187_v107_: D6
        d_187_v107_ = D6_DC17(_dafny.SeqWithoutIsStrInference([False]))
        pat_let_tv0_ = d_111_v52_
        index36_ = default__.safeIndex(622, (d_184_v106_).length(0))
        def iife20_(_pat_let0_0):
            def iife21_(d_188_dt__update__tmp_h0_):
                def iife22_(_pat_let1_0):
                    def iife23_(d_189_dt__update_hcf24_h0_):
                        return D6_DC17(d_189_dt__update_hcf24_h0_)
                    return iife23_(_pat_let1_0)
                return iife22_(pat_let_tv0_)
            return iife21_(_pat_let0_0)
        (d_184_v106_)[index36_] = iife20_(d_187_v107_)
        d_54_v4_ = d_54_v4_
        d_190_v108_: C3
        nw24_ = C3()
        nw24_.ctor__((d_60_v10_).f10, d_59_v9_)
        d_190_v108_ = nw24_
        d_191_v109_: _dafny.Set
        d_191_v109_ = _dafny.Set({d_190_v108_, d_190_v108_, d_190_v108_})
        d_191_v109_ = d_191_v109_
        d_192_v110_: _dafny.Set
        d_192_v110_ = _dafny.Set({(d_60_v10_).f10, (d_60_v10_).fm1(d_54_v4_, d_49_v1_, len(d_118_v60_), d_53_globalState_), d_54_v4_})
        d_193_v111_: _dafny.Map
        d_193_v111_ = _dafny.Map({d_48_v0_: False})
        index37_ = default__.safeIndex(431, (d_103_v49_).length(0))
        rhs28_ = ((d_60_v10_).fm4(len(d_48_v0_), d_49_v1_, d_193_v111_, len(d_48_v0_), d_53_globalState_)) + (-538)
        rhs29_ = (d_60_v10_).f10
        rhs30_ = _dafny.Set({not((d_60_v10_).f10)})
        rhs31_ = (d_111_v52_)[default__.safeIndex((d_60_v10_).f9, len(d_111_v52_))]
        lhs20_ = d_103_v49_
        lhs21_ = default__.safeIndex(431, (d_103_v49_).length(0))
        lhs20_[lhs21_] = rhs28_
        d_54_v4_ = rhs29_
        d_192_v110_ = rhs30_
        d_54_v4_ = rhs31_
        d_194_v112_: D17
        d_194_v112_ = D17_DC47(d_54_v4_, (d_60_v10_).fm4((d_103_v49_)[default__.safeIndex(431, (d_103_v49_).length(0))], len(_dafny.Map({(d_60_v10_).f10: (d_60_v10_).fm4((d_103_v49_)[default__.safeIndex(431, (d_103_v49_).length(0))], len(d_48_v0_), d_193_v111_, (d_60_v10_).f9, d_53_globalState_)})), d_193_v111_, (d_103_v49_)[default__.safeIndex(431, (d_103_v49_).length(0))], d_53_globalState_), (d_60_v10_).f9)
        if (d_194_v112_).cf74:
            (d_190_v108_).m1(d_49_v1_, (d_60_v10_).f10, (d_60_v10_).f10, d_53_globalState_)
            d_195_v113_: _dafny.Map
            d_195_v113_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wfd")): d_192_v110_})
            d_195_v113_ = (d_195_v113_).set((d_48_v0_) + (d_48_v0_), d_192_v110_)
            (d_60_v10_).m1((d_103_v49_)[default__.safeIndex(431, (d_103_v49_).length(0))], (d_60_v10_).f10, True, d_53_globalState_)
            if (d_111_v52_)[default__.safeIndex(776, len(d_111_v52_))]:
                d_54_v4_ = ((d_54_v4_) or (d_54_v4_)) or ((d_49_v1_) in ((d_113_v54_).set(8, (d_103_v49_)[default__.safeIndex(431, (d_103_v49_).length(0))])))
                d_54_v4_ = ((d_60_v10_).f10 if (d_60_v10_).f10 else (d_54_v4_) or ((d_60_v10_).f10))
                d_196_v114_: _dafny.MultiSet
                d_196_v114_ = _dafny.MultiSet([d_192_v110_, _dafny.Set({d_54_v4_}), d_192_v110_, d_192_v110_])
                d_197_v115_: C3
                nw25_ = C3()
                nw25_.ctor__((d_60_v10_).fm1((d_60_v10_).f10, d_49_v1_, (d_196_v114_).cardinality, d_53_globalState_), (d_59_v9_) | (d_59_v9_))
                d_197_v115_ = nw25_
                rhs32_ = (d_60_v10_).f9
                rhs33_ = default__.safeDivisionInt(d_49_v1_, (0) - ((d_103_v49_)[default__.safeIndex(431, (d_103_v49_).length(0))]))
                rhs34_ = True
                lhs22_ = d_53_globalState_
                lhs23_ = d_53_globalState_
                lhs22_.f0 = rhs32_
                lhs23_.f0 = rhs33_
                d_54_v4_ = rhs34_
                d_198_v116_: _dafny.Array
                nw26_ = _dafny.Array(None, 16)
                d_198_v116_ = nw26_
                d_199_v117_: T2
                nw27_ = C5()
                nw27_.ctor__((d_60_v10_).f9, (d_60_v10_).f10, d_56_v6_, (d_60_v10_).f10, d_54_v4_, d_59_v9_)
                d_199_v117_ = nw27_
                index38_ = default__.safeIndex(22, (d_198_v116_).length(0))
                (d_198_v116_)[index38_] = d_199_v117_
                d_200_v118_: _dafny.Array
                nw28_ = _dafny.Array(_dafny.Map({}), 27)
                d_200_v118_ = nw28_
                index39_ = default__.safeIndex(431, (d_103_v49_).length(0))
                index40_ = default__.safeIndex(22, (d_198_v116_).length(0))
                rhs35_ = (d_60_v10_).f9
                rhs36_ = d_199_v117_
                rhs37_ = d_200_v118_
                rhs38_ = d_48_v0_
                lhs24_ = d_103_v49_
                lhs25_ = default__.safeIndex(431, (d_103_v49_).length(0))
                lhs26_ = d_198_v116_
                lhs27_ = default__.safeIndex(22, (d_198_v116_).length(0))
                lhs24_[lhs25_] = rhs35_
                lhs26_[lhs27_] = rhs36_
                d_200_v118_ = rhs37_
                d_48_v0_ = rhs38_
            elif True:
                d_201_v119_: _dafny.Array
                nw29_ = _dafny.Array(_dafny.Seq({}), 20)
                d_201_v119_ = nw29_
                d_202_v120_: _dafny.Seq
                d_202_v120_ = _dafny.SeqWithoutIsStrInference([d_103_v49_])
                index41_ = default__.safeIndex(154, (d_201_v119_).length(0))
                (d_201_v119_)[index41_] = (d_202_v120_) + (_dafny.SeqWithoutIsStrInference([d_103_v49_]))
                index42_ = default__.safeIndex(154, (d_201_v119_).length(0))
                (d_201_v119_)[index42_] = _dafny.SeqWithoutIsStrInference([d_103_v49_, d_103_v49_])
                d_203_v121_: _dafny.MultiSet
                d_203_v121_ = _dafny.MultiSet([d_49_v1_])
                d_203_v121_ = d_203_v121_
                d_204_v122_: C6
                nw30_ = C6()
                nw30_.ctor__(_dafny.Set({(d_60_v10_).f9}), d_50_v2_)
                d_204_v122_ = nw30_
                d_54_v4_ = (d_60_v10_).f10
                (d_53_globalState_).f0 = (d_60_v10_).f9
            d_205_v123_: _dafny.Map
            d_205_v123_ = _dafny.Map({(d_54_v4_) or (not((d_60_v10_).f10)): (False) == ((d_60_v10_).f10)})
            d_206_v124_: D12
            d_206_v124_ = D12_DC35(d_54_v4_)
            d_205_v123_ = (d_205_v123_).set((_dafny.SeqWithoutIsStrInference([not((d_206_v124_).cf55)])) != (d_111_v52_), d_54_v4_)
        elif True:
            d_207_v125_: D3
            d_207_v125_ = D3_DC9(d_103_v49_)
            d_208_v126_: D5
            d_208_v126_ = D5_DC16(d_207_v125_, 941, (d_60_v10_).f9, -528, d_102_v48_)
            d_209_v127_: _dafny.Map
            d_209_v127_ = _dafny.Map({d_208_v126_: d_103_v49_})
            d_103_v49_ = ((d_209_v127_)[d_208_v126_] if (d_208_v126_) in (d_209_v127_) else d_103_v49_)
            d_210_v128_: C3
            nw31_ = C3()
            nw31_.ctor__((d_60_v10_).f10, (_dafny.MultiSet([d_54_v4_, not((d_60_v10_).f10), (d_60_v10_).f10, (d_190_v108_).fm3(d_53_globalState_)])) | (d_59_v9_))
            d_210_v128_ = nw31_
            index43_ = default__.safeIndex(989, (d_56_v6_).length(0))
            (d_56_v6_)[index43_] = d_54_v4_
            index44_ = default__.safeIndex(989, (d_56_v6_).length(0))
            (d_56_v6_)[index44_] = d_54_v4_
            d_211_v129_: T0
            nw32_ = C2()
            nw32_.ctor__(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_54_v4_])))
            d_211_v129_ = nw32_
            d_212_v130_: _dafny.MultiSet
            d_212_v130_ = _dafny.MultiSet([d_211_v129_, d_211_v129_])
            d_213_v131_: D22
            d_213_v131_ = D22_DC57(d_211_v129_)
            d_214_v132_: _dafny.Seq
            d_214_v132_ = _dafny.SeqWithoutIsStrInference([d_211_v129_])
            d_215_v133_: _dafny.Array
            def lambda10_(d_216_v6_):
                def lambda11_(d_217_i10_):
                    return (d_216_v6_)[default__.safeIndex(989, (d_216_v6_).length(0))]

                return lambda11_

            init5_ = lambda10_(d_56_v6_)
            nw33_ = _dafny.Array(None, 10)
            for i0_5_ in range(nw33_.length(0)):
                nw33_[i0_5_] = init5_(i0_5_)
            d_215_v133_ = nw33_
            d_218_v134_: C4
            nw34_ = C4()
            nw34_.ctor__(d_215_v133_, (d_60_v10_).f10, d_54_v4_, d_59_v9_)
            d_218_v134_ = nw34_
            d_219_v135_: _dafny.Map
            d_219_v135_ = _dafny.Map({len(d_48_v0_): d_218_v134_})
            d_220_v136_: _dafny.Map
            d_220_v136_ = _dafny.Map({((d_219_v135_)[d_49_v1_] if (d_49_v1_) in (d_219_v135_) else d_218_v134_): d_212_v130_})
            d_221_v137_: _dafny.Array
            nw35_ = _dafny.Array(None, 23)
            nw35_[int(0)] = (d_212_v130_) | (_dafny.MultiSet([d_211_v129_]))
            nw35_[int(1)] = (d_212_v130_ if d_54_v4_ else _dafny.MultiSet([d_211_v129_, d_211_v129_, (d_213_v131_).cf88]))
            nw35_[int(2)] = d_212_v130_
            nw35_[int(3)] = (d_212_v130_).intersection(d_212_v130_)
            nw35_[int(4)] = d_212_v130_
            nw35_[int(5)] = d_212_v130_
            nw35_[int(6)] = (_dafny.MultiSet(d_214_v132_)).set(d_211_v129_, default__.abs((d_103_v49_)[default__.safeIndex(431, (d_103_v49_).length(0))]))
            nw35_[int(7)] = d_212_v130_
            nw35_[int(8)] = d_212_v130_
            nw35_[int(9)] = d_212_v130_
            nw35_[int(10)] = d_212_v130_
            nw35_[int(11)] = (d_212_v130_) | (_dafny.MultiSet([d_211_v129_, d_211_v129_]))
            nw35_[int(12)] = _dafny.MultiSet([d_211_v129_])
            nw35_[int(13)] = (d_212_v130_) - (d_212_v130_)
            nw35_[int(14)] = (_dafny.MultiSet(d_214_v132_)) | (d_212_v130_)
            nw35_[int(15)] = ((d_220_v136_)[d_218_v134_] if (d_218_v134_) in (d_220_v136_) else d_212_v130_)
            nw35_[int(16)] = (d_212_v130_) - (d_212_v130_)
            nw35_[int(17)] = (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_211_v129_]))).intersection((d_212_v130_).set(d_211_v129_, default__.abs((d_103_v49_)[default__.safeIndex(431, (d_103_v49_).length(0))])))
            nw35_[int(18)] = _dafny.MultiSet([d_211_v129_])
            nw35_[int(19)] = (d_212_v130_).intersection(d_212_v130_)
            nw35_[int(20)] = d_212_v130_
            nw35_[int(21)] = d_212_v130_
            nw35_[int(22)] = d_212_v130_
            d_221_v137_ = nw35_
            index45_ = default__.safeIndex(842, (d_221_v137_).length(0))
            (d_221_v137_)[index45_] = _dafny.MultiSet(d_214_v132_)
            index46_ = default__.safeIndex(842, (d_221_v137_).length(0))
            (d_221_v137_)[index46_] = d_212_v130_
            d_222_v138_: _dafny.Map
            d_222_v138_ = _dafny.Map({(default__.fm11((d_218_v134_).fm1(True, 357, d_49_v1_, d_53_globalState_), d_53_globalState_)) != (d_48_v0_): (d_218_v134_).fm6(d_53_globalState_)})
            d_222_v138_ = (d_222_v138_).set((d_60_v10_).f10, 36)
        (d_53_globalState_).f0 = d_49_v1_
        index47_ = default__.safeIndex(431, (d_103_v49_).length(0))
        (d_103_v49_)[index47_] = (d_103_v49_)[default__.safeIndex(431, (d_103_v49_).length(0))]
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_56_v6_).length(0)):
            d_223_i11_: int = guard_loop_0_
            if (True) and (((0) <= (d_223_i11_)) and ((d_223_i11_) < ((d_56_v6_).length(0)))):
                (d_56_v6_)[(d_223_i11_)] = (d_60_v10_).f10
        d_224_i12_: int
        d_224_i12_ = 0
        with _dafny.label("1"):
            while d_54_v4_:
                with _dafny.c_label("1"):
                    if (d_224_i12_) >= (100):
                        raise _dafny.Break("1")
                    d_224_i12_ = (d_224_i12_) + (1)
                    d_225_v139_: _dafny.Seq
                    d_225_v139_ = _dafny.SeqWithoutIsStrInference([len(d_48_v0_)])
                    d_226_v140_: _dafny.MultiSet
                    d_226_v140_ = _dafny.MultiSet([len((d_225_v139_) + (_dafny.SeqWithoutIsStrInference([(d_60_v10_).f9]))), (d_60_v10_).f9, (d_190_v108_).fm4(d_49_v1_, d_49_v1_, d_193_v111_, (d_60_v10_).f9, d_53_globalState_), (d_103_v49_)[default__.safeIndex(431, (d_103_v49_).length(0))]])
                    d_227_v141_: _dafny.Seq
                    d_227_v141_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(d_60_v10_).f9 for d_228_i13_ in range(default__.abs(279))])])
                    d_226_v140_ = _dafny.MultiSet((d_227_v141_)[default__.safeIndex((len(d_48_v0_)) + ((d_103_v49_)[default__.safeIndex(431, (d_103_v49_).length(0))]), len(d_227_v141_))])
                    d_49_v1_ = (d_60_v10_).f9
                    d_229_v142_: _dafny.Set
                    d_229_v142_ = _dafny.Set({d_56_v6_, d_56_v6_, d_56_v6_})
                    d_54_v4_ = ((d_229_v142_).intersection(_dafny.Set({d_56_v6_, d_56_v6_}))).isdisjoint(d_229_v142_)
                    d_230_v143_: _dafny.Array
                    nw36_ = _dafny.Array(_dafny.Array(None, 0), 1)
                    d_230_v143_ = nw36_
                    d_230_v143_ = d_230_v143_
                    pass
            pass
        _dafny.print((d_48_v0_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_49_v1_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_50_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_51_v3_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_51_v3_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_51_v3_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_51_v3_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_51_v3_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_51_v3_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_51_v3_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_51_v3_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_53_globalState_.f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_53_globalState_).f1).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_53_globalState_.f2)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_53_globalState_.f2)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_53_globalState_.f2)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_53_globalState_.f2)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_53_globalState_.f2)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_53_globalState_.f2)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_53_globalState_.f2)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_53_globalState_.f2)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_54_v4_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_55_v5_) == (_dafny.Map({_dafny.CodePoint('w'): True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_56_v6_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_56_v6_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_56_v6_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_56_v6_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_56_v6_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_56_v6_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_56_v6_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_56_v6_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_56_v6_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_56_v6_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_56_v6_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_56_v6_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_56_v6_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_56_v6_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_56_v6_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_56_v6_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_57_v7_) == (_dafny.Set({_dafny.CodePoint('o'), _dafny.CodePoint('w')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_58_v8_) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rwdkj")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rwwkj"))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_59_v9_) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_60_v10_).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_60_v10_).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_61_i1_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_100_v46_) == (_dafny.Map({True: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeerwdkj"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_101_v47_) == (_dafny.Map({_dafny.Map({-434: 143}): True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_102_v48_).f11) == (_dafny.Map({_dafny.Map({-434: 143}): True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_103_v49_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_103_v49_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_103_v49_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_103_v49_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_103_v49_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_103_v49_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_103_v49_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_103_v49_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_103_v49_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_103_v49_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_103_v49_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_103_v49_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_106_v51_)[0]) == (_dafny.Map({D8_DC24(143, True, _dafny.SeqWithoutIsStrInference([True]), 959, 143): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v'), _dafny.CodePoint('i')])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_106_v51_)[1]) == (_dafny.Map({D8_DC24(143, True, _dafny.SeqWithoutIsStrInference([True]), 959, 143): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v'), _dafny.CodePoint('i')])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_106_v51_)[2]) == (_dafny.Map({D8_DC24(143, True, _dafny.SeqWithoutIsStrInference([True]), 959, 143): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v'), _dafny.CodePoint('i')])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_106_v51_)[3]) == (_dafny.Map({D8_DC24(143, True, _dafny.SeqWithoutIsStrInference([True]), 959, 143): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v'), _dafny.CodePoint('i')])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_106_v51_)[4]) == (_dafny.Map({D8_DC24(143, True, _dafny.SeqWithoutIsStrInference([True]), 959, 143): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v'), _dafny.CodePoint('i')])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_106_v51_)[5]) == (_dafny.Map({D8_DC24(3, True, _dafny.SeqWithoutIsStrInference([True, True, True]), 143, 143): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w'), _dafny.CodePoint('w'), _dafny.CodePoint('k'), _dafny.CodePoint('w')])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_106_v51_)[6]) == (_dafny.Map({D8_DC24(143, True, _dafny.SeqWithoutIsStrInference([True]), 959, 143): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v'), _dafny.CodePoint('i')])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_106_v51_)[7]) == (_dafny.Map({D8_DC24(143, True, _dafny.SeqWithoutIsStrInference([True]), 959, 143): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v'), _dafny.CodePoint('i')])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_106_v51_)[8]) == (_dafny.Map({D8_DC24(143, True, _dafny.SeqWithoutIsStrInference([True]), 959, 143): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v'), _dafny.CodePoint('i')])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_106_v51_)[9]) == (_dafny.Map({D8_DC24(143, True, _dafny.SeqWithoutIsStrInference([True]), 959, 143): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v'), _dafny.CodePoint('i')])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_106_v51_)[10]) == (_dafny.Map({D8_DC24(143, True, _dafny.SeqWithoutIsStrInference([True]), 959, 143): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v'), _dafny.CodePoint('i')])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_106_v51_)[11]) == (_dafny.Map({D8_DC24(143, True, _dafny.SeqWithoutIsStrInference([True]), 959, 143): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v'), _dafny.CodePoint('i')])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_106_v51_)[12]) == (_dafny.Map({D8_DC24(143, True, _dafny.SeqWithoutIsStrInference([True]), 959, 143): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v'), _dafny.CodePoint('i')])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_106_v51_)[13]) == (_dafny.Map({D8_DC24(143, True, _dafny.SeqWithoutIsStrInference([True]), 959, 143): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v'), _dafny.CodePoint('i')])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_106_v51_)[14]) == (_dafny.Map({D8_DC24(143, True, _dafny.SeqWithoutIsStrInference([True]), 959, 143): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v'), _dafny.CodePoint('i')])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_106_v51_)[15]) == (_dafny.Map({D8_DC24(143, True, _dafny.SeqWithoutIsStrInference([True]), 959, 143): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v'), _dafny.CodePoint('i')])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_106_v51_)[16]) == (_dafny.Map({D8_DC24(143, True, _dafny.SeqWithoutIsStrInference([True]), 959, 143): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v'), _dafny.CodePoint('i')])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_106_v51_)[17]) == (_dafny.Map({D8_DC24(143, True, _dafny.SeqWithoutIsStrInference([True]), 959, 143): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v'), _dafny.CodePoint('i')])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_106_v51_)[18]) == (_dafny.Map({D8_DC24(143, True, _dafny.SeqWithoutIsStrInference([True]), 959, 143): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v'), _dafny.CodePoint('i')])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_106_v51_)[19]) == (_dafny.Map({D8_DC24(143, True, _dafny.SeqWithoutIsStrInference([True]), 959, 143): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v'), _dafny.CodePoint('i')])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_111_v52_) == (_dafny.SeqWithoutIsStrInference([True, True, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_112_v53_).cf34))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_112_v53_).cf35))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_112_v53_).cf36) == (_dafny.SeqWithoutIsStrInference([True, True, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_112_v53_).cf37))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_112_v53_).cf38))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_113_v54_) == (_dafny.Map({143: 143}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_114_v55_) == (_dafny.Map({D8_DC24(3, True, _dafny.SeqWithoutIsStrInference([True, True, True]), 143, 143): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w'), _dafny.CodePoint('w'), _dafny.CodePoint('k'), _dafny.CodePoint('w')])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_115_v57_) == (_dafny.SeqWithoutIsStrInference([D8_DC24(3, True, _dafny.SeqWithoutIsStrInference([True, True, True]), 143, 143), D8_DC24(3, True, _dafny.SeqWithoutIsStrInference([True, True, True]), 143, 143)]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_116_v58_) == (_dafny.SeqWithoutIsStrInference([_dafny.Map({D8_DC24(3, True, _dafny.SeqWithoutIsStrInference([True, True, True]), 143, 143): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeerwdkj"))})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_118_v60_) == (_dafny.Set({0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_184_v106_)[0]).cf24) == (_dafny.SeqWithoutIsStrInference([True, True, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_184_v106_)[1]).cf24) == (_dafny.SeqWithoutIsStrInference([True, True, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_184_v106_)[2]).cf24) == (_dafny.SeqWithoutIsStrInference([True, True, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_187_v107_).cf24) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_191_v109_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_192_v110_) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_193_v111_) == (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeerwdkj")): False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_194_v112_).cf74))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_194_v112_).cf75))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_194_v112_).cf76))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_224_i12_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D0_DC1)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)

class D0_DC1(D0, NamedTuple('DC1', [])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1)
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC0(D0, NamedTuple('DC0', [('cf0', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC0({_dafny.string_of(self.cf0)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC0) and self.cf0 == __o.cf0
    def __hash__(self) -> int:
        return super().__hash__()


class D1:
    @classmethod
    def default(cls, ):
        return lambda: D1_DC3(D0.default()(), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D1_DC2)

class D1_DC3(D1, NamedTuple('DC3', [('cf2', Any), ('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf2 == __o.cf2 and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC2(D1, NamedTuple('DC2', [('cf1', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC2({_dafny.string_of(self.cf1)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC2) and self.cf1 == __o.cf1
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC5(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D2_DC5)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D2_DC4)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D2_DC8)

class D2_DC5(D2, NamedTuple('DC5', [('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC5({_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC5) and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC6(D2, NamedTuple('DC6', [])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6)
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC7(D2, NamedTuple('DC7', [('cf6', Any), ('cf7', Any), ('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf6)}, {_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf6 == __o.cf6 and self.cf7 == __o.cf7 and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC4(D2, NamedTuple('DC4', [('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC4({_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC4) and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC8(D2, NamedTuple('DC8', [('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC8({_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC8) and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC10(False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D3_DC10)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D3_DC11)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D3_DC9)

class D3_DC10(D3, NamedTuple('DC10', [('cf11', Any), ('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC10({_dafny.string_of(self.cf11)}, {_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC10) and self.cf11 == __o.cf11 and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC11(D3, NamedTuple('DC11', [('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC11({_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC11) and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC9(D3, NamedTuple('DC9', [('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC13(False, _dafny.Map({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D4_DC13)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D4_DC12)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D4_DC14)

class D4_DC13(D4, NamedTuple('DC13', [('cf15', Any), ('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC13({_dafny.string_of(self.cf15)}, {_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC13) and self.cf15 == __o.cf15 and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC12(D4, NamedTuple('DC12', [('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC12({_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC12) and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC14(D4, NamedTuple('DC14', [('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC14({_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC14) and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC16(D3.default()(), int(0), int(0), int(0), None)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D5_DC16)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D5_DC15)

class D5_DC16(D5, NamedTuple('DC16', [('cf19', Any), ('cf20', Any), ('cf21', Any), ('cf22', Any), ('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC16({_dafny.string_of(self.cf19)}, {_dafny.string_of(self.cf20)}, {_dafny.string_of(self.cf21)}, {_dafny.string_of(self.cf22)}, {_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC16) and self.cf19 == __o.cf19 and self.cf20 == __o.cf20 and self.cf21 == __o.cf21 and self.cf22 == __o.cf22 and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC15(D5, NamedTuple('DC15', [('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC15({self.cf18.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC15) and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC18(None, _dafny.Seq({}), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D6_DC18)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D6_DC17)

class D6_DC18(D6, NamedTuple('DC18', [('cf25', Any), ('cf26', Any), ('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC18({_dafny.string_of(self.cf25)}, {_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC18) and self.cf25 == __o.cf25 and self.cf26 == __o.cf26 and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC17(D6, NamedTuple('DC17', [('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC17({_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC17) and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC20(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D7_DC20)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D7_DC21)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D7_DC19)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D7_DC22)

class D7_DC20(D7, NamedTuple('DC20', [('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC20({_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC20) and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC21(D7, NamedTuple('DC21', [('cf30', Any), ('cf31', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC21({_dafny.string_of(self.cf30)}, {_dafny.string_of(self.cf31)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC21) and self.cf30 == __o.cf30 and self.cf31 == __o.cf31
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC19(D7, NamedTuple('DC19', [('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC19({_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC19) and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC22(D7, NamedTuple('DC22', [('cf32', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC22({_dafny.string_of(self.cf32)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC22) and self.cf32 == __o.cf32
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC24(int(0), False, _dafny.Seq({}), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D8_DC24)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D8_DC25)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D8_DC23)

class D8_DC24(D8, NamedTuple('DC24', [('cf34', Any), ('cf35', Any), ('cf36', Any), ('cf37', Any), ('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC24({_dafny.string_of(self.cf34)}, {_dafny.string_of(self.cf35)}, {_dafny.string_of(self.cf36)}, {_dafny.string_of(self.cf37)}, {_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC24) and self.cf34 == __o.cf34 and self.cf35 == __o.cf35 and self.cf36 == __o.cf36 and self.cf37 == __o.cf37 and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC25(D8, NamedTuple('DC25', [('cf39', Any), ('cf40', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC25({_dafny.string_of(self.cf39)}, {_dafny.string_of(self.cf40)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC25) and self.cf39 == __o.cf39 and self.cf40 == __o.cf40
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC23(D8, NamedTuple('DC23', [('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC23({_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC23) and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC27(_dafny.Map({}), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D9_DC27)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D9_DC26)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D9_DC28)

class D9_DC27(D9, NamedTuple('DC27', [('cf42', Any), ('cf43', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC27({_dafny.string_of(self.cf42)}, {_dafny.string_of(self.cf43)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC27) and self.cf42 == __o.cf42 and self.cf43 == __o.cf43
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC26(D9, NamedTuple('DC26', [('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC26({_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC26) and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC28(D9, NamedTuple('DC28', [('cf44', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC28({_dafny.string_of(self.cf44)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC28) and self.cf44 == __o.cf44
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: D10_DC30(False, False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D10_DC30)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D10_DC31)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D10_DC29)

class D10_DC30(D10, NamedTuple('DC30', [('cf46', Any), ('cf47', Any), ('cf48', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC30({_dafny.string_of(self.cf46)}, {_dafny.string_of(self.cf47)}, {_dafny.string_of(self.cf48)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC30) and self.cf46 == __o.cf46 and self.cf47 == __o.cf47 and self.cf48 == __o.cf48
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC31(D10, NamedTuple('DC31', [('cf49', Any), ('cf50', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC31({_dafny.string_of(self.cf49)}, {_dafny.string_of(self.cf50)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC31) and self.cf49 == __o.cf49 and self.cf50 == __o.cf50
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC29(D10, NamedTuple('DC29', [('cf45', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC29({_dafny.string_of(self.cf45)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC29) and self.cf45 == __o.cf45
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC33(False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D11_DC33)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D11_DC32)

class D11_DC33(D11, NamedTuple('DC33', [('cf52', Any), ('cf53', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC33({_dafny.string_of(self.cf52)}, {_dafny.string_of(self.cf53)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC33) and self.cf52 == __o.cf52 and self.cf53 == __o.cf53
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC32(D11, NamedTuple('DC32', [('cf51', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC32({_dafny.string_of(self.cf51)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC32) and self.cf51 == __o.cf51
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC35(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D12_DC35)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D12_DC36)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D12_DC37)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D12_DC34)

class D12_DC35(D12, NamedTuple('DC35', [('cf55', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC35({_dafny.string_of(self.cf55)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC35) and self.cf55 == __o.cf55
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC36(D12, NamedTuple('DC36', [('cf56', Any), ('cf57', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC36({_dafny.string_of(self.cf56)}, {_dafny.string_of(self.cf57)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC36) and self.cf56 == __o.cf56 and self.cf57 == __o.cf57
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC37(D12, NamedTuple('DC37', [('cf58', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC37({_dafny.string_of(self.cf58)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC37) and self.cf58 == __o.cf58
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC34(D12, NamedTuple('DC34', [('cf54', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC34({_dafny.string_of(self.cf54)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC34) and self.cf54 == __o.cf54
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: D13_DC39(_dafny.Map({}), _dafny.Array(None, 0), _dafny.Seq({}), int(0), _dafny.Seq({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D13_DC39)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D13_DC38)

class D13_DC39(D13, NamedTuple('DC39', [('cf60', Any), ('cf61', Any), ('cf62', Any), ('cf63', Any), ('cf64', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC39({_dafny.string_of(self.cf60)}, {_dafny.string_of(self.cf61)}, {_dafny.string_of(self.cf62)}, {_dafny.string_of(self.cf63)}, {_dafny.string_of(self.cf64)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC39) and self.cf60 == __o.cf60 and self.cf61 == __o.cf61 and self.cf62 == __o.cf62 and self.cf63 == __o.cf63 and self.cf64 == __o.cf64
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC38(D13, NamedTuple('DC38', [('cf59', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC38({_dafny.string_of(self.cf59)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC38) and self.cf59 == __o.cf59
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D14_DC40)

class D14_DC40(D14, NamedTuple('DC40', [('cf65', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC40({_dafny.string_of(self.cf65)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC40) and self.cf65 == __o.cf65
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: D15_DC42(_dafny.MultiSet({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC42(self) -> bool:
        return isinstance(self, D15_DC42)
    @property
    def is_DC41(self) -> bool:
        return isinstance(self, D15_DC41)

class D15_DC42(D15, NamedTuple('DC42', [('cf67', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC42({_dafny.string_of(self.cf67)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC42) and self.cf67 == __o.cf67
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC41(D15, NamedTuple('DC41', [('cf66', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC41({_dafny.string_of(self.cf66)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC41) and self.cf66 == __o.cf66
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: D16_DC44(False, False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC44(self) -> bool:
        return isinstance(self, D16_DC44)
    @property
    def is_DC43(self) -> bool:
        return isinstance(self, D16_DC43)
    @property
    def is_DC45(self) -> bool:
        return isinstance(self, D16_DC45)

class D16_DC44(D16, NamedTuple('DC44', [('cf69', Any), ('cf70', Any), ('cf71', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC44({_dafny.string_of(self.cf69)}, {_dafny.string_of(self.cf70)}, {_dafny.string_of(self.cf71)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC44) and self.cf69 == __o.cf69 and self.cf70 == __o.cf70 and self.cf71 == __o.cf71
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC43(D16, NamedTuple('DC43', [('cf68', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC43({_dafny.string_of(self.cf68)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC43) and self.cf68 == __o.cf68
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC45(D16, NamedTuple('DC45', [('cf72', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC45({_dafny.string_of(self.cf72)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC45) and self.cf72 == __o.cf72
    def __hash__(self) -> int:
        return super().__hash__()


class D17:
    @classmethod
    def default(cls, ):
        return lambda: D17_DC47(False, int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC47(self) -> bool:
        return isinstance(self, D17_DC47)
    @property
    def is_DC46(self) -> bool:
        return isinstance(self, D17_DC46)

class D17_DC47(D17, NamedTuple('DC47', [('cf74', Any), ('cf75', Any), ('cf76', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC47({_dafny.string_of(self.cf74)}, {_dafny.string_of(self.cf75)}, {_dafny.string_of(self.cf76)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC47) and self.cf74 == __o.cf74 and self.cf75 == __o.cf75 and self.cf76 == __o.cf76
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC46(D17, NamedTuple('DC46', [('cf73', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC46({_dafny.string_of(self.cf73)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC46) and self.cf73 == __o.cf73
    def __hash__(self) -> int:
        return super().__hash__()


class D18:
    @classmethod
    def default(cls, ):
        return lambda: D18_DC49(_dafny.Seq({}), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC49(self) -> bool:
        return isinstance(self, D18_DC49)
    @property
    def is_DC48(self) -> bool:
        return isinstance(self, D18_DC48)
    @property
    def is_DC50(self) -> bool:
        return isinstance(self, D18_DC50)

class D18_DC49(D18, NamedTuple('DC49', [('cf78', Any), ('cf79', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC49({_dafny.string_of(self.cf78)}, {_dafny.string_of(self.cf79)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC49) and self.cf78 == __o.cf78 and self.cf79 == __o.cf79
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC48(D18, NamedTuple('DC48', [('cf77', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC48({_dafny.string_of(self.cf77)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC48) and self.cf77 == __o.cf77
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC50(D18, NamedTuple('DC50', [('cf80', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC50({_dafny.string_of(self.cf80)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC50) and self.cf80 == __o.cf80
    def __hash__(self) -> int:
        return super().__hash__()


class D19:
    @classmethod
    def default(cls, ):
        return lambda: D19_DC52(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC52(self) -> bool:
        return isinstance(self, D19_DC52)
    @property
    def is_DC51(self) -> bool:
        return isinstance(self, D19_DC51)

class D19_DC52(D19, NamedTuple('DC52', [('cf82', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC52({_dafny.string_of(self.cf82)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC52) and self.cf82 == __o.cf82
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC51(D19, NamedTuple('DC51', [('cf81', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC51({_dafny.string_of(self.cf81)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC51) and self.cf81 == __o.cf81
    def __hash__(self) -> int:
        return super().__hash__()


class D20:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC53(self) -> bool:
        return isinstance(self, D20_DC53)

class D20_DC53(D20, NamedTuple('DC53', [('cf83', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC53({_dafny.string_of(self.cf83)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC53) and self.cf83 == __o.cf83
    def __hash__(self) -> int:
        return super().__hash__()


class D21:
    @classmethod
    def default(cls, ):
        return lambda: D21_DC55(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC55(self) -> bool:
        return isinstance(self, D21_DC55)
    @property
    def is_DC56(self) -> bool:
        return isinstance(self, D21_DC56)
    @property
    def is_DC54(self) -> bool:
        return isinstance(self, D21_DC54)

class D21_DC55(D21, NamedTuple('DC55', [('cf85', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC55({_dafny.string_of(self.cf85)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC55) and self.cf85 == __o.cf85
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC56(D21, NamedTuple('DC56', [('cf86', Any), ('cf87', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC56({_dafny.string_of(self.cf86)}, {_dafny.string_of(self.cf87)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC56) and self.cf86 == __o.cf86 and self.cf87 == __o.cf87
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC54(D21, NamedTuple('DC54', [('cf84', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC54({_dafny.string_of(self.cf84)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC54) and self.cf84 == __o.cf84
    def __hash__(self) -> int:
        return super().__hash__()


class D22:
    @classmethod
    def default(cls, ):
        return lambda: D22_DC58()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC58(self) -> bool:
        return isinstance(self, D22_DC58)
    @property
    def is_DC57(self) -> bool:
        return isinstance(self, D22_DC57)

class D22_DC58(D22, NamedTuple('DC58', [])):
    def __dafnystr__(self) -> str:
        return f'D22.DC58'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC58)
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC57(D22, NamedTuple('DC57', [('cf88', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC57({_dafny.string_of(self.cf88)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC57) and self.cf88 == __o.cf88
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
    @property
    def f7(self):
        return self._f7
    @f7.setter
    def f7(self, value):
        self._f7 = value
    @property
    def f8(self):
        return self._f8
    @f8.setter
    def f8(self, value):
        self._f8 = value
    def m2(self, p0, globalState):
        pass

    def m3(self, globalState):
        pass


class GlobalState:
    def  __init__(self):
        self.f0: int = int(0)
        self.f2: _dafny.Array = _dafny.Array(None, 0)
        self._f1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2):
        (self).f0 = f0
        (self)._f1 = f1
        (self).f2 = f2

    @property
    def f1(self):
        return self._f1

class C0:
    def  __init__(self):
        self._f11: _dafny.Map = _dafny.Map({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f11):
        (self)._f11 = f11

    def fm10(self, p0, globalState):
        def iife24_():
            coll20_ = _dafny.Set()
            compr_20_: _dafny.Seq
            for compr_20_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bfuu")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qgnlm"))])).Elements:
                d_231_v0_: _dafny.Seq = compr_20_
                if (d_231_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bfuu")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qgnlm"))])):
                    coll20_ = coll20_.union(_dafny.Set([d_231_v0_]))
            return _dafny.Set(coll20_)
        return ((_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "abyn"))})) - (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tnbomydo"))}))).isdisjoint((iife24_()
        ) | (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hr"))})))

    @property
    def f11(self):
        return self._f11

class C1(T1, T0):
    def  __init__(self):
        self._f6: bool = False
        self._f5: _dafny.MultiSet = _dafny.MultiSet({})
        self._f12: bool = False
        self._f13: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    @property
    def f6(self):
        return self._f6
    @property
    def f5(self):
        return self._f5
    def ctor__(self, f12, f13, f6, f5):
        (self)._f12 = f12
        (self)._f13 = f13
        (self)._f6 = f6
        (self)._f5 = f5

    def fm3(self, globalState):
        return (_dafny.MultiSet([len(_dafny.Map({False: len((self).f13)}))])).ispropersubset(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([not((self).f6), (self).f12])), len(_dafny.SeqWithoutIsStrInference([579 for d_232_i0_ in range(default__.abs(323))])), 821, (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([(self).f12])), 562])).cardinality, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sejx"))).set(default__.safeIndex(len(_dafny.Map({(self).f12: True})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sejx")))), _dafny.CodePoint('i')))]))

    def fm4(self, p0, p1, p2, p3, globalState):
        return (0) - ((((self).f5 if ((0) - (-463)) >= (len(_dafny.Set({(self).f12, (self).f6}))) else ((self).f5).set((self).f12, default__.abs(-762)))).cardinality)

    def fm1(self, p0, p1, p2, globalState):
        return (self).f6

    def fm2(self, p0, globalState):
        return ((_dafny.Set({_dafny.SeqWithoutIsStrInference([(self).f6])})) | (_dafny.Set({_dafny.SeqWithoutIsStrInference([(self).f6]), _dafny.SeqWithoutIsStrInference([(self).f6, (self).f6, (self).f6])}))).isdisjoint(_dafny.Set({_dafny.SeqWithoutIsStrInference([(self).f6]), _dafny.SeqWithoutIsStrInference([(self).f12, True]), _dafny.SeqWithoutIsStrInference([not((self).f12)]), _dafny.SeqWithoutIsStrInference([(self).f12])}))

    def m1(self, p0, p1, p2, globalState):
        d_233_v0_: _dafny.Array
        nw37_ = _dafny.Array(None, 6)
        nw37_[int(0)] = p2
        nw37_[int(1)] = p2
        nw37_[int(2)] = (self).f6
        nw37_[int(3)] = p2
        nw37_[int(4)] = (self).f12
        nw37_[int(5)] = (self).f6
        d_233_v0_ = nw37_
        d_234_v1_: _dafny.MultiSet
        d_234_v1_ = _dafny.MultiSet([d_233_v0_])
        d_235_i0_: int
        d_235_i0_ = 0
        with _dafny.label("2"):
            while ((d_234_v1_) - (d_234_v1_)).issubset(d_234_v1_):
                with _dafny.c_label("2"):
                    if (d_235_i0_) >= (100):
                        raise _dafny.Break("2")
                    d_235_i0_ = (d_235_i0_) + (1)
                    d_236_v2_: _dafny.Set
                    d_236_v2_ = _dafny.Set({p0})
                    d_237_v3_: D2
                    d_237_v3_ = D2_DC4(d_236_v2_)
                    d_238_v4_: _dafny.Seq
                    d_238_v4_ = _dafny.SeqWithoutIsStrInference([d_237_v3_])
                    d_239_v5_: _dafny.Seq
                    d_239_v5_ = _dafny.SeqWithoutIsStrInference([d_238_v4_])
                    d_240_v6_: _dafny.Map
                    d_240_v6_ = _dafny.Map({p0: d_238_v4_})
                    d_239_v5_ = (_dafny.SeqWithoutIsStrInference([d_238_v4_])).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference([d_238_v4_]))), ((d_240_v6_)[p0] if (p0) in (d_240_v6_) else d_238_v4_))
                    d_241_v7_: _dafny.Array
                    nw38_ = _dafny.Array(D2.default()(), 17)
                    d_241_v7_ = nw38_
                    d_242_v8_: str
                    d_242_v8_ = _dafny.CodePoint('a')
                    d_243_v9_: _dafny.Map
                    d_243_v9_ = _dafny.Map({p0: -852})
                    d_244_v10_: D2
                    d_244_v10_ = D2_DC7(d_242_v8_, d_243_v9_, (0) - ((0) - (p0)))
                    index48_ = default__.safeIndex(949, (d_241_v7_).length(0))
                    (d_241_v7_)[index48_] = d_244_v10_
                    index49_ = default__.safeIndex(949, (d_241_v7_).length(0))
                    (d_241_v7_)[index49_] = d_244_v10_
                    d_245_v11_: bool
                    d_245_v11_ = True
                    d_245_v11_ = (self).f12
                    (globalState).f0 = p0
                    pass
            pass
        d_246_v12_: str
        d_246_v12_ = _dafny.CodePoint('n')
        d_247_v13_: _dafny.Map
        d_247_v13_ = _dafny.Map({((self).f5).cardinality: p0})
        d_248_v14_: D2
        d_248_v14_ = D2_DC7(d_246_v12_, d_247_v13_, p0)
        hi0_ = p0
        for d_249_i1_ in range(default__.safeModuloInt(p0, (d_248_v14_).cf8), hi0_):
            d_250_v15_: _dafny.Seq
            d_250_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gmpwl"))
            d_250_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ahtrhyr"))
            d_251_v16_: _dafny.MultiSet
            d_251_v16_ = _dafny.MultiSet([p0, p0])
            (globalState).f0 = (d_249_i1_ if p1 else (d_251_v16_).cardinality)
            d_252_v17_: _dafny.Seq
            d_252_v17_ = _dafny.SeqWithoutIsStrInference([p2])
            d_253_v18_: _dafny.Set
            d_253_v18_ = _dafny.Set({p1})
            d_252_v17_ = _dafny.SeqWithoutIsStrInference([((self).f12) and ((self).fm1(p2, p0, d_249_i1_, globalState)), (self).f6, (p2) not in (d_253_v18_)])
            d_254_v19_: bool
            d_254_v19_ = True
            d_255_v20_: _dafny.Seq
            d_255_v20_ = _dafny.SeqWithoutIsStrInference([d_247_v13_, (d_248_v14_).cf7])
            d_254_v19_ = (d_247_v13_) not in (d_255_v20_)
        (globalState).f0 = p0
        d_256_v21_: _dafny.Array
        def lambda12_(d_257_p0_):
            def lambda13_(d_258_i3_):
                return (d_258_i3_) - ((0) - (d_257_p0_))

            return lambda13_

        init6_ = lambda12_(p0)
        nw39_ = _dafny.Array(None, 13)
        for i0_6_ in range(nw39_.length(0)):
            nw39_[i0_6_] = init6_(i0_6_)
        d_256_v21_ = nw39_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_256_v21_).length(0)):
            d_259_i2_: int = guard_loop_1_
            if (True) and (((0) <= (d_259_i2_)) and ((d_259_i2_) < ((d_256_v21_).length(0)))):
                (d_256_v21_)[(d_259_i2_)] = default__.safeModuloInt(d_259_i2_, p0)
        d_260_v22_: D5
        d_260_v22_ = D5_DC15((self).f13)
        d_261_v23_: _dafny.MultiSet
        d_261_v23_ = _dafny.MultiSet([d_246_v12_, _dafny.CodePoint('q')])
        d_262_v24_: _dafny.Map
        d_262_v24_ = _dafny.Map({(d_260_v22_).cf18: d_261_v23_})
        d_262_v24_ = (d_262_v24_).set(((self).f13) + ((self).f13), d_261_v23_)
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_256_v21_).length(0)):
            d_263_i4_: int = guard_loop_2_
            if (True) and (((0) <= (d_263_i4_)) and ((d_263_i4_) < ((d_256_v21_).length(0)))):
                (d_256_v21_)[(d_263_i4_)] = (d_263_i4_) - (p0)

    def m7(self, p0, p1, p2, globalState):
        d_264_v0_: _dafny.Map
        d_264_v0_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([(self).f12])): (0) - (p2)})
        d_265_v1_: _dafny.Seq
        d_265_v1_ = _dafny.SeqWithoutIsStrInference([p2])
        d_266_v2_: _dafny.Set
        d_266_v2_ = _dafny.Set({(self).f6, not((self).f12)})
        d_267_v3_: _dafny.Map
        d_267_v3_ = _dafny.Map({(default__.fm14(d_264_v0_, (self).f12, (d_265_v1_)[default__.safeIndex(len((self).f13), len(d_265_v1_))], (self).f12, globalState)).cardinality: (d_266_v2_).issubset(d_266_v2_)})
        d_268_v4_: _dafny.Map
        d_268_v4_ = _dafny.Map({not((self).fm2((self).f12, globalState)): (self).f12})
        d_267_v3_ = (d_267_v3_).set((0) - ((len(d_268_v4_)) * (p2)), (self).f6)
        d_269_v5_: _dafny.Map
        d_269_v5_ = _dafny.Map({(self).f13: ((d_267_v3_)[p2] if (p2) in (d_267_v3_) else (self).f6)})
        d_270_i0_: int
        d_270_i0_ = 0
        with _dafny.label("3"):
            while ((len((self).f13)) - ((self).fm4(19, len((self).f13), d_269_v5_, p0, globalState))) > (p0):
                with _dafny.c_label("3"):
                    if (d_270_i0_) >= (100):
                        raise _dafny.Break("3")
                    d_270_i0_ = (d_270_i0_) + (1)
                    d_271_v6_: _dafny.Map
                    d_271_v6_ = _dafny.Map({d_264_v0_: not((self).f12)})
                    d_272_v7_: C0
                    nw40_ = C0()
                    nw40_.ctor__(d_271_v6_)
                    d_272_v7_ = nw40_
                    d_273_v8_: D2
                    d_273_v8_ = D2_DC7(_dafny.CodePoint('m'), d_264_v0_, p2)
                    d_274_v9_: _dafny.MultiSet
                    d_274_v9_ = _dafny.MultiSet([d_273_v8_])
                    d_275_v10_: str
                    d_275_v10_ = _dafny.CodePoint('i')
                    pat_let_tv1_ = d_266_v2_
                    pat_let_tv2_ = p0
                    pat_let_tv3_ = p1
                    pat_let_tv4_ = p2
                    pat_let_tv5_ = d_264_v0_
                    pat_let_tv6_ = d_264_v0_
                    pat_let_tv7_ = d_266_v2_
                    pat_let_tv8_ = p0
                    pat_let_tv9_ = p1
                    pat_let_tv10_ = p2
                    pat_let_tv11_ = d_264_v0_
                    pat_let_tv12_ = d_264_v0_
                    def iife27_(_pat_let4_0):
                        def iife28_(d_276_dt__update__tmp_h3_):
                            def iife29_(_pat_let5_0):
                                def iife30_(d_277_dt__update_hcf8_h2_):
                                    def iife31_(_pat_let6_0):
                                        def iife32_(d_278_dt__update_hcf7_h3_):
                                            return D2_DC7((d_276_dt__update__tmp_h3_).cf6, d_278_dt__update_hcf7_h3_, d_277_dt__update_hcf8_h2_)
                                        return iife32_(_pat_let6_0)
                                    return iife31_(_dafny.Map({(0) - (pat_let_tv2_): (pat_let_tv3_).cardinality}))
                                return iife30_(_pat_let5_0)
                            return iife29_(len(pat_let_tv1_))
                        return iife28_(_pat_let4_0)
                    def iife26_(_pat_let3_0):
                        def iife33_(d_279_dt__update__tmp_h4_):
                            def iife34_(_pat_let7_0):
                                def iife35_(d_280_dt__update_hcf8_h3_):
                                    def iife36_(_pat_let8_0):
                                        def iife37_(d_281_dt__update_hcf7_h4_):
                                            return D2_DC7((d_279_dt__update__tmp_h4_).cf6, d_281_dt__update_hcf7_h4_, d_280_dt__update_hcf8_h3_)
                                        return iife37_(_pat_let8_0)
                                    return iife36_(pat_let_tv5_)
                                return iife35_(_pat_let7_0)
                            return iife34_(pat_let_tv4_)
                        return iife33_(_pat_let3_0)
                    def iife25_(_pat_let2_0):
                        def iife38_(d_282_dt__update__tmp_h5_):
                            def iife39_(_pat_let9_0):
                                def iife40_(d_283_dt__update_hcf7_h5_):
                                    return D2_DC7((d_282_dt__update__tmp_h5_).cf6, d_283_dt__update_hcf7_h5_, (d_282_dt__update__tmp_h5_).cf8)
                                return iife40_(_pat_let9_0)
                            return iife39_(pat_let_tv6_)
                        return iife38_(_pat_let2_0)
                    def iife43_(_pat_let12_0):
                        def iife44_(d_284_dt__update__tmp_h0_):
                            def iife45_(_pat_let13_0):
                                def iife46_(d_285_dt__update_hcf8_h0_):
                                    def iife47_(_pat_let14_0):
                                        def iife48_(d_286_dt__update_hcf7_h0_):
                                            return D2_DC7((d_284_dt__update__tmp_h0_).cf6, d_286_dt__update_hcf7_h0_, d_285_dt__update_hcf8_h0_)
                                        return iife48_(_pat_let14_0)
                                    return iife47_(_dafny.Map({(0) - (pat_let_tv8_): (pat_let_tv9_).cardinality}))
                                return iife46_(_pat_let13_0)
                            return iife45_(len(pat_let_tv7_))
                        return iife44_(_pat_let12_0)
                    def iife42_(_pat_let11_0):
                        def iife49_(d_287_dt__update__tmp_h1_):
                            def iife50_(_pat_let15_0):
                                def iife51_(d_288_dt__update_hcf8_h1_):
                                    def iife52_(_pat_let16_0):
                                        def iife53_(d_289_dt__update_hcf7_h1_):
                                            return D2_DC7((d_287_dt__update__tmp_h1_).cf6, d_289_dt__update_hcf7_h1_, d_288_dt__update_hcf8_h1_)
                                        return iife53_(_pat_let16_0)
                                    return iife52_(pat_let_tv11_)
                                return iife51_(_pat_let15_0)
                            return iife50_(pat_let_tv10_)
                        return iife49_(_pat_let11_0)
                    def iife41_(_pat_let10_0):
                        def iife54_(d_290_dt__update__tmp_h2_):
                            def iife55_(_pat_let17_0):
                                def iife56_(d_291_dt__update_hcf7_h2_):
                                    return D2_DC7((d_290_dt__update__tmp_h2_).cf6, d_291_dt__update_hcf7_h2_, (d_290_dt__update__tmp_h2_).cf8)
                                return iife56_(_pat_let17_0)
                            return iife55_(pat_let_tv12_)
                        return iife54_(_pat_let10_0)
                    (globalState).f0 = ((d_274_v9_)[iife25_(iife26_(iife27_(D2_DC7(d_275_v10_, _dafny.Map({p2: p2}), p2))))] if (iife41_(iife42_(iife43_(D2_DC7(d_275_v10_, _dafny.Map({p2: p2}), p2))))) in (d_274_v9_) else (0) - (p0))
                    d_292_v11_: bool
                    d_292_v11_ = False
                    d_292_v11_ = (p2) < (p0)
                    d_293_v12_: D3
                    d_293_v12_ = D3_DC10((self).f6, p0)
                    (globalState).f0 = (d_293_v12_).cf12
                    pass
            pass
        (globalState).f0 = (0) - (p0)
        d_294_v13_: _dafny.Seq
        d_294_v13_ = _dafny.SeqWithoutIsStrInference([p1, default__.fm15((self).f6, globalState), (_dafny.MultiSet(d_265_v1_)) | (p1)])
        d_295_v14_: _dafny.Array
        def lambda14_(d_296_i1_):
            return _dafny.SeqWithoutIsStrInference([(self).f13 for d_297_i2_ in range(default__.abs(-706))])

        init7_ = lambda14_
        nw41_ = _dafny.Array(None, 20)
        for i0_7_ in range(nw41_.length(0)):
            nw41_[i0_7_] = init7_(i0_7_)
        d_295_v14_ = nw41_
        d_298_v15_: _dafny.Seq
        d_298_v15_ = _dafny.SeqWithoutIsStrInference([(self).f13])
        index50_ = default__.safeIndex(639, (d_295_v14_).length(0))
        (d_295_v14_)[index50_] = d_298_v15_
        d_299_v16_: _dafny.Array
        nw42_ = _dafny.Array(_dafny.CodePoint('D'), 18)
        d_299_v16_ = nw42_
        d_300_v17_: _dafny.Seq
        d_300_v17_ = _dafny.SeqWithoutIsStrInference([(self).f12, (self).f12])
        index51_ = default__.safeIndex(639, (d_295_v14_).length(0))
        rhs39_ = d_294_v13_
        rhs40_ = p0
        rhs41_ = d_299_v16_
        rhs42_ = _dafny.SeqWithoutIsStrInference([(self).f13 for d_301_i3_ in range(default__.abs(362))])
        rhs43_ = ((p2 if (self).fm2((self).f6, globalState) else p0)) - ((len(d_300_v17_) if (self).f12 else p0))
        lhs28_ = globalState
        lhs29_ = globalState
        lhs30_ = d_295_v14_
        lhs31_ = default__.safeIndex(639, (d_295_v14_).length(0))
        lhs32_ = globalState
        d_294_v13_ = rhs39_
        lhs28_.f0 = rhs40_
        lhs29_.f2 = rhs41_
        lhs30_[lhs31_] = rhs42_
        lhs32_.f0 = rhs43_
        d_302_v18_: bool
        d_302_v18_ = False
        d_303_v19_: _dafny.Seq
        d_303_v19_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gvyhp"))
        rhs44_ = ((0) - ((len(d_265_v1_) if (self).f12 else (p1).cardinality))) < (60)
        rhs45_ = d_303_v19_
        d_302_v18_ = rhs44_
        d_303_v19_ = rhs45_
        (globalState).f0 = (0) - (p2)

    def m8(self, p0, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: bool = False
        r3: bool = False
        d_304_v0_: int
        d_304_v0_ = 375
        d_305_v1_: _dafny.Map
        d_305_v1_ = _dafny.Map({d_304_v0_: 110})
        d_306_v2_: _dafny.Map
        d_306_v2_ = _dafny.Map({d_305_v1_: (self).f6})
        d_307_v3_: C0
        nw43_ = C0()
        nw43_.ctor__(d_306_v2_)
        d_307_v3_ = nw43_
        d_308_v4_: _dafny.Seq
        d_308_v4_ = _dafny.SeqWithoutIsStrInference([not((self).f6), (self).f6])
        d_309_v5_: _dafny.Map
        d_309_v5_ = _dafny.Map({d_307_v3_: len(d_308_v4_)})
        d_309_v5_ = (d_309_v5_).set(d_307_v3_, default__.safeModuloInt(d_304_v0_, d_304_v0_))
        d_310_v6_: _dafny.Seq
        d_310_v6_ = _dafny.SeqWithoutIsStrInference([d_304_v0_])
        d_311_v7_: _dafny.Map
        d_311_v7_ = _dafny.Map({p0: not((self).f12)})
        d_312_v8_: _dafny.Map
        d_312_v8_ = _dafny.Map({(d_310_v6_)[default__.safeIndex(d_304_v0_, len(d_310_v6_))]: (d_311_v7_) | (d_311_v7_)})
        d_312_v8_ = (d_312_v8_).set((d_304_v0_) - ((0) - (d_304_v0_)), d_311_v7_)
        r2 = not (p0) or ((self).fm3(globalState))
        d_313_i0_: int
        d_313_i0_ = 0
        with _dafny.label("4"):
            while (self).f6:
                with _dafny.c_label("4"):
                    if (d_313_i0_) >= (100):
                        raise _dafny.Break("4")
                    d_313_i0_ = (d_313_i0_) + (1)
                    d_304_v0_ = d_304_v0_
                    d_314_v9_: _dafny.Array
                    def lambda15_(d_315_i1_):
                        return (self).f12

                    init8_ = lambda15_
                    nw44_ = _dafny.Array(None, 25)
                    for i0_8_ in range(nw44_.length(0)):
                        nw44_[i0_8_] = init8_(i0_8_)
                    d_314_v9_ = nw44_
                    d_314_v9_ = d_314_v9_
                    d_307_v3_ = d_307_v3_
                    d_316_v10_: _dafny.Set
                    d_316_v10_ = _dafny.Set({d_304_v0_})
                    d_317_v11_: D6
                    d_317_v11_ = D6_DC18(d_307_v3_, d_308_v4_, len(d_316_v10_))
                    d_311_v7_ = _dafny.Map({(self).f12: (_dafny.MultiSet(d_308_v4_)) == (_dafny.MultiSet((d_317_v11_).cf26))})
                    pass
            pass
        d_318_v12_: _dafny.Array
        nw45_ = _dafny.Array(None, 7)
        nw45_[int(0)] = d_304_v0_
        nw45_[int(1)] = d_304_v0_
        nw45_[int(2)] = d_304_v0_
        nw45_[int(3)] = 686
        nw45_[int(4)] = ((d_305_v1_)[d_304_v0_] if (d_304_v0_) in (d_305_v1_) else d_304_v0_)
        nw45_[int(5)] = d_304_v0_
        nw45_[int(6)] = d_304_v0_
        d_318_v12_ = nw45_
        index52_ = default__.safeIndex(643, (d_318_v12_).length(0))
        (d_318_v12_)[index52_] = len((d_310_v6_).set(default__.safeIndex(d_304_v0_, len(d_310_v6_)), d_304_v0_))
        d_319_v13_: _dafny.Map
        d_319_v13_ = _dafny.Map({(d_308_v4_)[default__.safeIndex(d_304_v0_, len(d_308_v4_))]: len(d_310_v6_)})
        d_320_v14_: _dafny.Set
        d_320_v14_ = _dafny.Set({not((self).f6), (self).f6, (self).f12, (d_308_v4_)[default__.safeIndex(len((d_319_v13_).set((self).f6, d_304_v0_)), len(d_308_v4_))]})
        index53_ = default__.safeIndex(643, (d_318_v12_).length(0))
        (d_318_v12_)[index53_] = ((d_304_v0_) * (len(d_320_v14_)) if p0 else d_304_v0_)
        d_321_i2_: int
        d_321_i2_ = 0
        with _dafny.label("5"):
            while p0:
                with _dafny.c_label("5"):
                    if (d_321_i2_) >= (100):
                        raise _dafny.Break("5")
                    d_321_i2_ = (d_321_i2_) + (1)
                    d_322_v15_: C0
                    nw46_ = C0()
                    nw46_.ctor__(((d_307_v3_).f11) | (((d_307_v3_).f11).set(d_305_v1_, (self).f12)))
                    d_322_v15_ = nw46_
                    r2 = (self).f12
                    d_323_v16_: C0
                    nw47_ = C0()
                    nw47_.ctor__((d_307_v3_).f11)
                    d_323_v16_ = nw47_
                    d_324_v17_: _dafny.Array
                    nw48_ = _dafny.Array(False, 11)
                    d_324_v17_ = nw48_
                    index54_ = default__.safeIndex(35, (d_324_v17_).length(0))
                    (d_324_v17_)[index54_] = ((d_311_v7_)[p0] if (p0) in (d_311_v7_) else (self).f12)
                    index55_ = default__.safeIndex(35, (d_324_v17_).length(0))
                    (d_324_v17_)[index55_] = (self).f12
                    pass
            pass
        r0 = not ((self).f12) or ((self).f12)
        r1 = (d_304_v0_) * (default__.safeDivisionInt((0) - (d_304_v0_), 585))
        r2 = True
        r3 = (self).fm2(p0, globalState)
        return r0, r1, r2, r3

    @property
    def f12(self):
        return self._f12
    @property
    def f13(self):
        return self._f13

class C2(T0):
    def  __init__(self):
        self._f5: _dafny.MultiSet = _dafny.MultiSet({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    @property
    def f5(self):
        return self._f5
    def ctor__(self, f5):
        (self)._f5 = f5

    def fm1(self, p0, p1, p2, globalState):
        return not (True) or (True)

    def fm2(self, p0, globalState):
        return (len(_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uvl")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_325_i0_ in range(default__.abs(81))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hxubpnhri"))}))) > (211)

    def m9(self, p0, p1, globalState):
        r0: _dafny.Map = _dafny.Map({})
        d_326_v0_: str
        d_326_v0_ = _dafny.CodePoint('g')
        d_327_v1_: _dafny.Seq
        d_327_v1_ = _dafny.SeqWithoutIsStrInference([p0])
        d_328_v2_: _dafny.Map
        d_328_v2_ = _dafny.Map({p1: len(d_327_v1_)})
        d_329_v3_: _dafny.Map
        d_329_v3_ = _dafny.Map({p0: d_328_v2_})
        d_326_v0_ = default__.fm16(p0, ((d_329_v3_)[p0] if (p0) in (d_329_v3_) else d_328_v2_), default__.fm17(globalState), globalState)
        d_330_v4_: _dafny.Seq
        d_330_v4_ = _dafny.SeqWithoutIsStrInference([p1])
        d_330_v4_ = default__.fm18(globalState)
        d_331_v5_: _dafny.Map
        d_331_v5_ = _dafny.Map({p0: p1})
        d_332_v6_: _dafny.Map
        d_332_v6_ = _dafny.Map({not(p0): p0})
        d_333_v7_: _dafny.Seq
        d_333_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ibtnyjf"))
        d_334_v8_: D7
        d_334_v8_ = D7_DC19(_dafny.Map({p0: p1}))
        d_335_v9_: _dafny.Array
        nw49_ = _dafny.Array(None, 23)
        nw49_[int(0)] = (d_331_v5_ if p0 else _dafny.Map({True: p1}))
        nw49_[int(1)] = _dafny.Map({p0: p1})
        nw49_[int(2)] = d_331_v5_
        nw49_[int(3)] = d_331_v5_
        nw49_[int(4)] = _dafny.Map({p0: p1})
        nw49_[int(5)] = (d_331_v5_) | (d_331_v5_)
        nw49_[int(6)] = (d_331_v5_) | (d_331_v5_)
        nw49_[int(7)] = _dafny.Map({True: p1})
        nw49_[int(8)] = _dafny.Map({((d_332_v6_)[p0] if (p0) in (d_332_v6_) else p0): p1})
        nw49_[int(9)] = default__.fm19((D3_DC10(False, len(d_333_v7_))).cf12, globalState)
        nw49_[int(10)] = (d_331_v5_) | (d_331_v5_)
        nw49_[int(11)] = ((d_331_v5_).set(True, len(d_333_v7_)) if p0 else d_331_v5_)
        nw49_[int(12)] = d_331_v5_
        nw49_[int(13)] = d_331_v5_
        nw49_[int(14)] = (d_331_v5_ if p0 else d_331_v5_)
        nw49_[int(15)] = ((d_334_v8_).cf28) | (d_331_v5_)
        nw49_[int(16)] = d_331_v5_
        nw49_[int(17)] = d_331_v5_
        nw49_[int(18)] = ((d_331_v5_).set((self).fm2(False, globalState), p1) if not(True) else d_331_v5_)
        nw49_[int(19)] = d_331_v5_
        nw49_[int(20)] = default__.fm19((0) - (p1), globalState)
        nw49_[int(21)] = (d_331_v5_) | (_dafny.Map({p0: p1}))
        nw49_[int(22)] = (_dafny.Map({False: p1})) | (d_331_v5_)
        d_335_v9_ = nw49_
        d_336_v10_: _dafny.Array
        nw50_ = _dafny.Array(_dafny.Array(None, 0), 18)
        d_336_v10_ = nw50_
        d_337_v11_: bool
        d_337_v11_ = True
        d_338_v12_: _dafny.Set
        d_338_v12_ = _dafny.Set({p1, len(d_332_v6_)})
        d_339_v13_: _dafny.Map
        d_339_v13_ = _dafny.Map({p1: d_338_v12_})
        d_340_v14_: D5
        d_340_v14_ = D5_DC15(d_333_v7_)
        d_341_v15_: _dafny.MultiSet
        d_341_v15_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_326_v0_ for d_342_i0_ in range(default__.abs(674))]), d_333_v7_])
        rhs46_ = d_335_v9_
        rhs47_ = d_336_v10_
        rhs48_ = ((d_332_v6_)[(d_338_v12_).isdisjoint(((d_339_v13_)[p1] if (p1) in (d_339_v13_) else d_338_v12_))] if ((d_338_v12_).isdisjoint(((d_339_v13_)[p1] if (p1) in (d_339_v13_) else d_338_v12_))) in (d_332_v6_) else (d_333_v7_) <= ((d_340_v14_).cf18))
        rhs49_ = ((d_341_v15_).ispropersubset(d_341_v15_)) and (False)
        d_335_v9_ = rhs46_
        d_336_v10_ = rhs47_
        d_337_v11_ = rhs48_
        d_337_v11_ = rhs49_
        d_343_v16_: _dafny.Seq
        d_343_v16_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_326_v0_ for d_344_i1_ in range(default__.abs(165))])])
        d_345_v17_: _dafny.Array
        nw51_ = _dafny.Array(int(0), 24)
        d_345_v17_ = nw51_
        index56_ = default__.safeIndex(736, (d_345_v17_).length(0))
        (d_345_v17_)[index56_] = p1
        d_346_v19_: D8
        d_346_v19_ = D8_DC23(_dafny.Map({_dafny.SeqWithoutIsStrInference([d_337_v11_]): len(d_330_v4_)}))
        index57_ = default__.safeIndex(736, (d_345_v17_).length(0))
        def iife57_():
            coll21_ = _dafny.Map()
            compr_21_: int
            for compr_21_ in _dafny.IntegerRange(366, 185):
                d_347_v18_: int = compr_21_
                if ((366) <= (d_347_v18_)) and ((d_347_v18_) < (185)):
                    coll21_[(d_347_v18_) + ((_dafny.MultiSet([d_326_v0_])).cardinality)] = p1
            return _dafny.Map(coll21_)
        rhs50_ = (d_343_v16_) + ((d_343_v16_) + (d_343_v16_))
        rhs51_ = (p1) - (p1)
        rhs52_ = default__.fm16(True, (iife57_()
        ).set(len((d_346_v19_).cf33), p1), (d_333_v7_) + (d_333_v7_), globalState)
        rhs53_ = p1
        rhs54_ = p1
        lhs33_ = globalState
        lhs34_ = d_345_v17_
        lhs35_ = default__.safeIndex(736, (d_345_v17_).length(0))
        lhs36_ = globalState
        d_343_v16_ = rhs50_
        lhs33_.f0 = rhs51_
        d_326_v0_ = rhs52_
        lhs34_[lhs35_] = rhs53_
        lhs36_.f0 = rhs54_
        d_348_v20_: _dafny.Array
        def lambda16_(d_349_v7_):
            def lambda17_(d_350_i2_):
                return d_349_v7_

            return lambda17_

        init9_ = lambda16_(d_333_v7_)
        nw52_ = _dafny.Array(None, 8)
        for i0_9_ in range(nw52_.length(0)):
            nw52_[i0_9_] = init9_(i0_9_)
        d_348_v20_ = nw52_
        index58_ = default__.safeIndex(150, (d_348_v20_).length(0))
        (d_348_v20_)[index58_] = _dafny.SeqWithoutIsStrInference([d_326_v0_ for d_351_i3_ in range(default__.abs(333))])
        d_352_v21_: D1
        d_352_v21_ = D1_DC2(d_337_v11_)
        d_353_v22_: _dafny.Map
        d_353_v22_ = _dafny.Map({d_326_v0_: d_345_v17_})
        index59_ = default__.safeIndex(736, (d_345_v17_).length(0))
        index60_ = default__.safeIndex(150, (d_348_v20_).length(0))
        rhs55_ = p0
        rhs56_ = (d_345_v17_)[default__.safeIndex(736, (d_345_v17_).length(0))]
        rhs57_ = (d_333_v7_).set(default__.safeIndex(len(d_353_v22_), len(d_333_v7_)), default__.fm16(not(p0), _dafny.Map({(d_345_v17_)[default__.safeIndex(736, (d_345_v17_).length(0))]: (d_345_v17_)[default__.safeIndex(736, (d_345_v17_).length(0))]}), (d_343_v16_)[default__.safeIndex(-325, len(d_343_v16_))], globalState))
        rhs58_ = d_352_v21_
        lhs37_ = d_345_v17_
        lhs38_ = default__.safeIndex(736, (d_345_v17_).length(0))
        lhs39_ = d_348_v20_
        lhs40_ = default__.safeIndex(150, (d_348_v20_).length(0))
        d_337_v11_ = rhs55_
        lhs37_[lhs38_] = rhs56_
        lhs39_[lhs40_] = rhs57_
        d_352_v21_ = rhs58_
        (globalState).f0 = p1
        d_354_v23_: _dafny.Array
        nw53_ = _dafny.Array(False, 19)
        d_354_v23_ = nw53_
        d_355_v24_: _dafny.Map
        d_355_v24_ = _dafny.Map({d_354_v23_: p0})
        d_356_v25_: _dafny.Seq
        d_356_v25_ = _dafny.SeqWithoutIsStrInference([d_355_v24_, (d_355_v24_) | (d_355_v24_), d_355_v24_])
        r0 = (d_356_v25_)[default__.safeIndex(len(d_338_v12_), len(d_356_v25_))]
        return r0


class C3(T1, T0):
    def  __init__(self):
        self._f6: bool = False
        self._f5: _dafny.MultiSet = _dafny.MultiSet({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f6(self):
        return self._f6
    @property
    def f5(self):
        return self._f5
    def ctor__(self, f6, f5):
        (self)._f6 = f6
        (self)._f5 = f5

    def fm3(self, globalState):
        return (self).f6

    def fm4(self, p0, p1, p2, p3, globalState):
        return (len((_dafny.SeqWithoutIsStrInference([(self).f6])) + (_dafny.SeqWithoutIsStrInference([(self).f6])))) * ((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rassq")))) - (len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cax"))): (0) - (len(_dafny.SeqWithoutIsStrInference([not((self).f6)])))}))))

    def fm1(self, p0, p1, p2, globalState):
        source3_ = D0_DC0(-247)
        if source3_.is_DC1:
            return (self).f6
        elif True:
            d_357___mcc_h0_ = source3_.cf0
            d_358_cf0_ = d_357___mcc_h0_
            return not((self).f6)

    def fm2(self, p0, globalState):
        return (D1_DC2((self).f6)).cf1

    def m1(self, p0, p1, p2, globalState):
        d_359_v0_: D0
        d_359_v0_ = D0_DC1()
        d_359_v0_ = d_359_v0_
        d_360_v1_: _dafny.Seq
        d_360_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vekcmyumj"))
        d_361_v2_: D0
        d_361_v2_ = D0_DC0(len(d_360_v1_))
        pat_let_tv13_ = p0
        def iife58_(_pat_let18_0):
            def iife59_(d_362_dt__update__tmp_h0_):
                def iife60_(_pat_let19_0):
                    def iife61_(d_363_dt__update_hcf0_h0_):
                        return D0_DC0(d_363_dt__update_hcf0_h0_)
                    return iife61_(_pat_let19_0)
                return iife60_(pat_let_tv13_)
            return iife59_(_pat_let18_0)
        source4_ = iife58_(d_361_v2_)
        if source4_.is_DC1:
            if p1:
                d_364_v3_: _dafny.Map
                d_364_v3_ = _dafny.Map({p0: p2})
                d_365_v4_: str
                d_365_v4_ = _dafny.CodePoint('m')
                d_366_v5_: _dafny.Map
                d_366_v5_ = _dafny.Map({(-887) <= (len(d_364_v3_)): d_365_v4_})
                d_366_v5_ = (d_366_v5_).set((p0) != (p0), d_365_v4_)
                d_367_v6_: D1
                d_367_v6_ = D1_DC2((self).f6)
                d_368_v7_: _dafny.Map
                d_368_v7_ = _dafny.Map({(self).f6: p0})
                d_369_v8_: _dafny.Map
                d_369_v8_ = _dafny.Map({d_367_v6_: d_368_v7_})
                d_370_v9_: _dafny.Map
                d_370_v9_ = _dafny.Map({p0: (0) - (p0)})
                (globalState).f0 = (0) - ((p0 if (_dafny.Map({len(d_369_v8_): p0})) == (d_370_v9_) else p0))
                d_371_v10_: _dafny.Map
                d_371_v10_ = _dafny.Map({d_370_v9_: p1})
                d_372_v11_: C0
                nw54_ = C0()
                nw54_.ctor__(d_371_v10_)
                d_372_v11_ = nw54_
                d_373_v12_: _dafny.Map
                d_373_v12_ = _dafny.Map({d_372_v11_: p0})
                d_373_v12_ = (d_373_v12_).set(d_372_v11_, p0)
                d_374_v13_: C0
                nw55_ = C0()
                nw55_.ctor__(((d_372_v11_).f11) | (d_371_v10_))
                d_374_v13_ = nw55_
                (globalState).f0 = default__.safeModuloInt(p0, p0)
            elif True:
                d_375_v14_: bool
                d_375_v14_ = True
                d_375_v14_ = not(not((self).f6))
                d_376_v15_: _dafny.Array
                nw56_ = _dafny.Array(_dafny.Map({}), 22)
                d_376_v15_ = nw56_
                d_377_v16_: _dafny.Map
                d_377_v16_ = _dafny.Map({p0: p0})
                index61_ = default__.safeIndex(137, (d_376_v15_).length(0))
                (d_376_v15_)[index61_] = (d_377_v16_) | (d_377_v16_)
                index62_ = default__.safeIndex(137, (d_376_v15_).length(0))
                (d_376_v15_)[index62_] = d_377_v16_
                d_378_v17_: _dafny.Array
                nw57_ = _dafny.Array(False, 2)
                d_378_v17_ = nw57_
                index63_ = default__.safeIndex(682, (d_378_v17_).length(0))
                (d_378_v17_)[index63_] = d_375_v14_
                d_379_v18_: _dafny.Array
                nw58_ = _dafny.Array(int(0), 15)
                d_379_v18_ = nw58_
                d_380_v19_: _dafny.Map
                d_380_v19_ = _dafny.Map({d_379_v18_: p2})
                d_381_v20_: _dafny.Seq
                d_381_v20_ = _dafny.SeqWithoutIsStrInference([(self).f6])
                index64_ = default__.safeIndex(682, (d_378_v17_).length(0))
                rhs59_ = (self).f6
                rhs60_ = not((d_380_v19_) != ((_dafny.Map({d_379_v18_: p1})) | (d_380_v19_)))
                rhs61_ = not((d_381_v20_)[default__.safeIndex(p0, len(d_381_v20_))])
                lhs41_ = d_378_v17_
                lhs42_ = default__.safeIndex(682, (d_378_v17_).length(0))
                lhs41_[lhs42_] = rhs59_
                d_375_v14_ = rhs60_
                d_375_v14_ = rhs61_
                d_382_v21_: _dafny.Map
                d_382_v21_ = _dafny.Map({p1: ((d_378_v17_)[default__.safeIndex(682, (d_378_v17_).length(0))]) or ((self).f6)})
                d_382_v21_ = (d_382_v21_).set(False, (d_378_v17_)[default__.safeIndex(682, (d_378_v17_).length(0))])
                (globalState).f0 = ((p0) - (p0)) * ((p0) + (len(d_381_v20_)))
            (globalState).f0 = p0
            (globalState).f0 = (0) - (-971)
            d_383_v22_: bool
            d_383_v22_ = True
            d_384_v23_: _dafny.Set
            d_384_v23_ = _dafny.Set({p0})
            d_385_v24_: _dafny.Map
            d_385_v24_ = _dafny.Map({p0: d_359_v0_})
            d_386_v26_: _dafny.Map
            d_386_v26_ = _dafny.Map({False: p0})
            def iife62_():
                coll22_ = _dafny.Set()
                compr_22_: int
                for compr_22_ in (d_385_v24_).keys.Elements:
                    d_387_v25_: int = compr_22_
                    if (d_387_v25_) in (d_385_v24_):
                        coll22_ = coll22_.union(_dafny.Set([default__.safeDivisionInt(d_387_v25_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i"))))]))
                return _dafny.Set(coll22_)
            rhs62_ = not((d_384_v23_).isdisjoint(iife62_()
            ))
            rhs63_ = (default__.safeModuloInt(((d_386_v26_)[False] if (False) in (d_386_v26_) else p0), -562)) + (p0)
            lhs43_ = globalState
            d_383_v22_ = rhs62_
            lhs43_.f0 = rhs63_
        elif True:
            d_388___mcc_h0_ = source4_.cf0
            d_389_cf0_ = d_388___mcc_h0_
            d_390_v27_: _dafny.MultiSet
            d_390_v27_ = _dafny.MultiSet([d_389_cf0_])
            d_390_v27_ = d_390_v27_
            d_391_v28_: D1
            d_391_v28_ = D1_DC3(d_359_v0_, default__.safeDivisionInt(655, d_389_cf0_))
            d_392_v29_: bool
            d_392_v29_ = True
            d_393_v30_: _dafny.Array
            def lambda18_(d_394_v29_):
                def lambda19_(d_395_i0_):
                    return d_394_v29_

                return lambda19_

            init10_ = lambda18_(d_392_v29_)
            nw59_ = _dafny.Array(None, 7)
            for i0_10_ in range(nw59_.length(0)):
                nw59_[i0_10_] = init10_(i0_10_)
            d_393_v30_ = nw59_
            d_396_v31_: _dafny.Map
            d_396_v31_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ypkiarto")): (self).f6})
            d_397_v32_: _dafny.Seq
            d_397_v32_ = _dafny.SeqWithoutIsStrInference([p0, d_389_cf0_])
            rhs64_ = d_391_v28_
            rhs65_ = (self).fm4(d_389_cf0_, p0, (d_396_v31_) | (d_396_v31_), (d_397_v32_)[default__.safeIndex((0) - (d_389_cf0_), len(d_397_v32_))], globalState)
            rhs66_ = len(d_360_v1_)
            rhs67_ = p2
            rhs68_ = d_393_v30_
            lhs44_ = globalState
            lhs45_ = globalState
            d_391_v28_ = rhs64_
            lhs44_.f0 = rhs65_
            lhs45_.f0 = rhs66_
            d_392_v29_ = rhs67_
            d_393_v30_ = rhs68_
            d_398_v33_: _dafny.Map
            d_398_v33_ = _dafny.Map({len((d_397_v32_) + (d_397_v32_)): d_389_cf0_})
            d_399_v34_: _dafny.Set
            d_399_v34_ = _dafny.Set({(0) - (p0)})
            d_400_v35_: _dafny.Seq
            d_400_v35_ = _dafny.SeqWithoutIsStrInference([d_399_v34_])
            d_401_v40_: D2
            d_401_v40_ = D2_DC4(_dafny.Set({len(d_397_v32_), -705}))
            d_402_v42_: _dafny.Array
            nw60_ = _dafny.Array(None, 22)
            nw60_[int(0)] = d_399_v34_
            nw60_[int(1)] = (d_399_v34_) - (d_399_v34_)
            nw60_[int(2)] = d_399_v34_
            nw60_[int(3)] = (d_400_v35_)[default__.safeIndex(d_389_cf0_, len(d_400_v35_))]
            nw60_[int(4)] = d_399_v34_
            nw60_[int(5)] = d_399_v34_
            nw60_[int(6)] = _dafny.Set({p0})
            def iife63_():
                coll23_ = _dafny.Set()
                compr_23_: int
                for compr_23_ in _dafny.IntegerRange(753, 464):
                    d_403_v36_: int = compr_23_
                    if ((753) <= (d_403_v36_)) and ((d_403_v36_) < (464)):
                        coll23_ = coll23_.union(_dafny.Set([(d_403_v36_) + (p0)]))
                return _dafny.Set(coll23_)
            nw60_[int(7)] = iife63_()
            
            nw60_[int(8)] = d_399_v34_
            nw60_[int(9)] = d_399_v34_
            def iife64_():
                coll24_ = _dafny.Set()
                compr_24_: int
                for compr_24_ in (d_390_v27_).Elements:
                    d_404_v37_: int = compr_24_
                    if (d_404_v37_) in (d_390_v27_):
                        def iife65_():
                            coll25_ = _dafny.Map()
                            compr_25_: int
                            for compr_25_ in (_dafny.Set({113})).Elements:
                                d_405_v38_: int = compr_25_
                                if (d_405_v38_) in (_dafny.Set({113})):
                                    coll25_[(d_405_v38_) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y"))))] = not(False)
                            return _dafny.Map(coll25_)
                        coll24_ = coll24_.union(_dafny.Set([default__.safeDivisionInt(d_404_v37_, len(_dafny.Map({len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Map({False: True}))), (_dafny.MultiSet([False])).cardinality])): _dafny.SeqWithoutIsStrInference([711, len(iife65_()
)])})): len(_dafny.SeqWithoutIsStrInference([_dafny.Map({len(_dafny.SeqWithoutIsStrInference([True])): False})]))})))]))
                return _dafny.Set(coll24_)
            nw60_[int(10)] = iife64_()
            
            nw60_[int(11)] = (d_399_v34_).intersection(d_399_v34_)
            def iife66_():
                coll26_ = _dafny.Set()
                compr_26_: int
                for compr_26_ in _dafny.IntegerRange(636, 285):
                    d_406_v39_: int = compr_26_
                    if ((636) <= (d_406_v39_)) and ((d_406_v39_) < (285)):
                        coll26_ = coll26_.union(_dafny.Set([default__.safeDivisionInt(d_406_v39_, d_389_cf0_)]))
                return _dafny.Set(coll26_)
            nw60_[int(12)] = iife66_()
            
            nw60_[int(13)] = (d_399_v34_).intersection(_dafny.Set({p0}))
            nw60_[int(14)] = d_399_v34_
            nw60_[int(15)] = _dafny.Set({p0})
            nw60_[int(16)] = d_399_v34_
            nw60_[int(17)] = (d_401_v40_).cf4
            nw60_[int(18)] = d_399_v34_
            def iife67_():
                coll27_ = _dafny.Set()
                compr_27_: int
                for compr_27_ in _dafny.IntegerRange(-506, 339):
                    d_407_v41_: int = compr_27_
                    if ((-506) <= (d_407_v41_)) and ((d_407_v41_) < (339)):
                        coll27_ = coll27_.union(_dafny.Set([(d_407_v41_) * (25)]))
                return _dafny.Set(coll27_)
            nw60_[int(19)] = iife67_()
            
            nw60_[int(20)] = d_399_v34_
            nw60_[int(21)] = d_399_v34_
            d_402_v42_ = nw60_
            rhs69_ = d_398_v33_
            rhs70_ = (105) * ((((self).f5) | ((self).f5)).cardinality)
            rhs71_ = d_402_v42_
            d_398_v33_ = rhs69_
            d_389_cf0_ = rhs70_
            d_402_v42_ = rhs71_
            index65_ = default__.safeIndex(160, (d_393_v30_).length(0))
            (d_393_v30_)[index65_] = not((self).fm2(p2, globalState))
            d_408_v43_: _dafny.Array
            nw61_ = _dafny.Array(_dafny.Array(None, 0), 25)
            d_408_v43_ = nw61_
            d_409_v44_: _dafny.Array
            nw62_ = _dafny.Array(int(0), 3)
            d_409_v44_ = nw62_
            d_410_v45_: D3
            d_410_v45_ = D3_DC9(d_409_v44_)
            d_411_v46_: _dafny.Array
            nw63_ = _dafny.Array(None, 11)
            nw63_[int(0)] = d_409_v44_
            nw63_[int(1)] = d_409_v44_
            nw63_[int(2)] = d_409_v44_
            nw63_[int(3)] = (d_410_v45_).cf10
            nw63_[int(4)] = d_409_v44_
            nw63_[int(5)] = d_409_v44_
            nw63_[int(6)] = d_409_v44_
            nw63_[int(7)] = d_409_v44_
            nw63_[int(8)] = d_409_v44_
            nw63_[int(9)] = d_409_v44_
            nw63_[int(10)] = d_409_v44_
            d_411_v46_ = nw63_
            index66_ = default__.safeIndex(823, (d_408_v43_).length(0))
            (d_408_v43_)[index66_] = d_411_v46_
            index67_ = default__.safeIndex(26, (d_409_v44_).length(0))
            (d_409_v44_)[index67_] = p0
            d_412_v47_: _dafny.Array
            nw64_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 1)
            d_412_v47_ = nw64_
            index68_ = default__.safeIndex(717, (d_412_v47_).length(0))
            (d_412_v47_)[index68_] = d_360_v1_
            index69_ = default__.safeIndex(160, (d_393_v30_).length(0))
            index70_ = default__.safeIndex(823, (d_408_v43_).length(0))
            index71_ = default__.safeIndex(26, (d_409_v44_).length(0))
            index72_ = default__.safeIndex(717, (d_412_v47_).length(0))
            rhs72_ = d_392_v29_
            rhs73_ = d_411_v46_
            rhs74_ = (d_389_cf0_) * (p0)
            rhs75_ = d_360_v1_
            rhs76_ = (0) - (p0)
            lhs46_ = d_393_v30_
            lhs47_ = default__.safeIndex(160, (d_393_v30_).length(0))
            lhs48_ = d_408_v43_
            lhs49_ = default__.safeIndex(823, (d_408_v43_).length(0))
            lhs50_ = d_409_v44_
            lhs51_ = default__.safeIndex(26, (d_409_v44_).length(0))
            lhs52_ = d_412_v47_
            lhs53_ = default__.safeIndex(717, (d_412_v47_).length(0))
            lhs54_ = globalState
            lhs46_[lhs47_] = rhs72_
            lhs48_[lhs49_] = rhs73_
            lhs50_[lhs51_] = rhs74_
            lhs52_[lhs53_] = rhs75_
            lhs54_.f0 = rhs76_
        d_413_v48_: _dafny.Array
        nw65_ = _dafny.Array(int(0), 10)
        d_413_v48_ = nw65_
        d_414_v49_: _dafny.Map
        d_414_v49_ = _dafny.Map({((d_360_v1_).set(default__.safeIndex(p0, len(d_360_v1_)), _dafny.CodePoint('r'))) + (d_360_v1_): d_413_v48_})
        d_414_v49_ = (d_414_v49_).set((d_360_v1_) + (d_360_v1_), d_413_v48_)
        d_415_i1_: int
        d_415_i1_ = 0
        with _dafny.label("6"):
            while True:
                with _dafny.c_label("6"):
                    if (d_415_i1_) >= (100):
                        raise _dafny.Break("6")
                    d_415_i1_ = (d_415_i1_) + (1)
                    d_416_v50_: bool
                    d_416_v50_ = False
                    d_416_v50_ = not(((self).f6) or (d_416_v50_))
                    d_417_v51_: _dafny.Set
                    d_417_v51_ = _dafny.Set({(D2_DC5((self).f6)).cf5})
                    d_418_v52_: str
                    d_418_v52_ = _dafny.CodePoint('m')
                    (globalState).f0 = ((len(d_417_v51_)) - (len(_dafny.Set({d_418_v52_, d_418_v52_, d_418_v52_}))) if p1 else default__.safeDivisionInt(p0, p0))
                    d_419_v53_: _dafny.Map
                    d_419_v53_ = _dafny.Map({d_416_v50_: (self).fm3(globalState)})
                    d_420_v54_: _dafny.Set
                    d_420_v54_ = _dafny.Set({d_418_v52_})
                    d_421_v55_: _dafny.Map
                    d_421_v55_ = _dafny.Map({_dafny.Map({d_416_v50_: (self).f6}): d_420_v54_})
                    d_416_v50_ = (d_419_v53_) not in (d_421_v55_)
                    d_422_v56_: _dafny.Map
                    d_422_v56_ = _dafny.Map({p0: p0})
                    d_423_v57_: _dafny.Map
                    d_423_v57_ = _dafny.Map({d_422_v56_: not(p2)})
                    d_424_v58_: C0
                    nw66_ = C0()
                    nw66_.ctor__(d_423_v57_)
                    d_424_v58_ = nw66_
                    pass
            pass
        d_425_v59_: _dafny.Array
        nw67_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 29)
        d_425_v59_ = nw67_
        d_426_v60_: _dafny.Seq
        d_426_v60_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f"))])
        d_427_v61_: str
        d_427_v61_ = _dafny.CodePoint('b')
        d_428_v62_: _dafny.Map
        d_428_v62_ = _dafny.Map({p0: (0) - (p0)})
        d_429_v63_: D2
        d_429_v63_ = D2_DC7(d_427_v61_, d_428_v62_, p0)
        index73_ = default__.safeIndex(599, (d_425_v59_).length(0))
        (d_425_v59_)[index73_] = (d_426_v60_)[default__.safeIndex((self).fm4((d_429_v63_).cf8, p0, _dafny.Map({d_360_v1_: not((self).f6)}), p0, globalState), len(d_426_v60_))]
        index74_ = default__.safeIndex(599, (d_425_v59_).length(0))
        (d_425_v59_)[index74_] = (default__.fm11((self).f6, globalState)) + (d_360_v1_)
        d_430_v64_: _dafny.Map
        d_430_v64_ = _dafny.Map({d_360_v1_: p2})
        d_431_v65_: _dafny.Map
        d_431_v65_ = _dafny.Map({p1: (self).f6})
        hi1_ = (self).fm4(p0, p0, d_430_v64_, len(d_431_v65_), globalState)
        for d_432_i2_ in range(((default__.fm12(globalState)).intersection(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([p0])))).cardinality, hi1_):
            d_433_v66_: bool
            d_433_v66_ = False
            d_434_v67_: _dafny.Array
            nw68_ = _dafny.Array(_dafny.CodePoint('D'), 22)
            d_434_v67_ = nw68_
            d_435_v68_: _dafny.Seq
            d_435_v68_ = _dafny.SeqWithoutIsStrInference([d_434_v67_])
            d_436_v69_: _dafny.Array
            nw69_ = _dafny.Array(None, 12)
            nw69_[int(0)] = d_434_v67_
            nw69_[int(1)] = d_434_v67_
            nw69_[int(2)] = d_434_v67_
            nw69_[int(3)] = d_434_v67_
            nw69_[int(4)] = d_434_v67_
            nw69_[int(5)] = (d_435_v68_)[default__.safeIndex(d_432_i2_, len(d_435_v68_))]
            nw69_[int(6)] = d_434_v67_
            nw69_[int(7)] = d_434_v67_
            nw69_[int(8)] = d_434_v67_
            nw69_[int(9)] = d_434_v67_
            nw69_[int(10)] = d_434_v67_
            nw69_[int(11)] = d_434_v67_
            d_436_v69_ = nw69_
            d_437_v70_: _dafny.Seq
            d_437_v70_ = _dafny.SeqWithoutIsStrInference([(self).f6])
            d_438_v71_: _dafny.Seq
            d_438_v71_ = _dafny.SeqWithoutIsStrInference([(d_432_i2_) - (len(_dafny.Map({p0: len((d_425_v59_)[default__.safeIndex(599, (d_425_v59_).length(0))])}))), p0, (self).fm4(459, d_432_i2_, default__.fm13(len(d_437_v70_), d_429_v63_, p2, len(_dafny.Set({p1, (self).f6, (self).f6, p1})), globalState), 49, globalState)])
            d_439_v72_: _dafny.Map
            d_439_v72_ = _dafny.Map({d_427_v61_: d_436_v69_})
            d_440_v73_: D4
            d_440_v73_ = D4_DC12(((d_439_v72_)[d_427_v61_] if (d_427_v61_) in (d_439_v72_) else d_436_v69_))
            d_441_v74_: _dafny.Set
            d_441_v74_ = _dafny.Set({p0, (0) - (p0)})
            d_442_v75_: _dafny.Seq
            d_442_v75_ = _dafny.SeqWithoutIsStrInference([d_438_v71_, d_438_v71_])
            rhs77_ = not(p1)
            rhs78_ = (d_440_v73_).cf14
            rhs79_ = d_432_i2_
            rhs80_ = (default__.safeModuloInt(len(d_441_v74_), d_432_i2_)) * (default__.safeDivisionInt(d_432_i2_, d_432_i2_))
            rhs81_ = ((d_442_v75_)[default__.safeIndex(d_432_i2_, len(d_442_v75_))]) + (d_438_v71_)
            lhs55_ = globalState
            lhs56_ = globalState
            d_433_v66_ = rhs77_
            d_436_v69_ = rhs78_
            lhs55_.f0 = rhs79_
            lhs56_.f0 = rhs80_
            d_438_v71_ = rhs81_
            d_427_v61_ = d_427_v61_
            d_431_v65_ = (d_431_v65_).set((self).f6, d_433_v66_)
            d_443_v76_: _dafny.Map
            d_443_v76_ = _dafny.Map({p0: (self).f6})
            d_444_v77_: T1
            nw70_ = C1()
            nw70_.ctor__((self).f6, d_360_v1_, True, _dafny.MultiSet([False, d_433_v66_]))
            d_444_v77_ = nw70_
            d_445_v78_: _dafny.Array
            def lambda20_(d_446_v77_):
                def lambda21_(d_447_i3_):
                    return not((d_446_v77_).f6)

                return lambda21_

            init11_ = lambda20_(d_444_v77_)
            nw71_ = _dafny.Array(None, 18)
            for i0_11_ in range(nw71_.length(0)):
                nw71_[i0_11_] = init11_(i0_11_)
            d_445_v78_ = nw71_
            index75_ = default__.safeIndex(71, (d_445_v78_).length(0))
            (d_445_v78_)[index75_] = not((self).f6)
            d_448_v80_: _dafny.MultiSet
            d_448_v80_ = _dafny.MultiSet([d_361_v2_])
            index76_ = default__.safeIndex(71, (d_445_v78_).length(0))
            def iife68_():
                coll28_ = _dafny.Map()
                compr_28_: int
                for compr_28_ in _dafny.IntegerRange(938, -71):
                    d_449_v79_: int = compr_28_
                    if ((938) <= (d_449_v79_)) and ((d_449_v79_) < (-71)):
                        coll28_[(d_449_v79_) * (d_432_i2_)] = (d_444_v77_).f6
                return _dafny.Map(coll28_)
            rhs82_ = (d_443_v76_ if not((p0) <= (d_432_i2_)) else (iife68_()
             if (d_444_v77_).f6 else d_443_v76_))
            rhs83_ = d_444_v77_
            rhs84_ = d_427_v61_
            rhs85_ = ((d_448_v80_).set(d_361_v2_, default__.abs(d_432_i2_))).isdisjoint(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_361_v2_ for d_450_i4_ in range(default__.abs(522))])))
            lhs57_ = d_445_v78_
            lhs58_ = default__.safeIndex(71, (d_445_v78_).length(0))
            d_443_v76_ = rhs82_
            d_444_v77_ = rhs83_
            d_427_v61_ = rhs84_
            lhs57_[lhs58_] = rhs85_

    def m6(self, p0, p1, p2, p3, globalState):
        r0: D1 = D1.default()()
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r2: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r3: bool = False
        d_451_v0_: _dafny.Array
        nw72_ = _dafny.Array(int(0), 1)
        d_451_v0_ = nw72_
        nw73_ = _dafny.Array(int(0), 2)
        d_451_v0_ = nw73_
        r3 = p2
        if p2:
            if (self).f6:
                d_452_v1_: _dafny.Array
                nw74_ = _dafny.Array(False, 14)
                d_452_v1_ = nw74_
                index77_ = default__.safeIndex(765, (d_452_v1_).length(0))
                (d_452_v1_)[index77_] = (p0) == (p0)
                d_453_v2_: _dafny.Seq
                d_453_v2_ = _dafny.SeqWithoutIsStrInference([(self).fm1((self).fm3(globalState), p0, p0, globalState)])
                index78_ = default__.safeIndex(765, (d_452_v1_).length(0))
                (d_452_v1_)[index78_] = (p2) and ((p0) != (len(d_453_v2_)))
                (globalState).f0 = (0) - (p0)
                d_454_v3_: _dafny.Array
                nw75_ = _dafny.Array(_dafny.Map({}), 28)
                d_454_v3_ = nw75_
                d_455_v4_: C1
                nw76_ = C1()
                nw76_.ctor__(False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xy")), p2, _dafny.MultiSet(d_453_v2_))
                d_455_v4_ = nw76_
                d_456_v5_: _dafny.Array
                nw77_ = _dafny.Array(None, 29)
                nw77_[int(0)] = d_455_v4_
                nw77_[int(1)] = d_455_v4_
                nw77_[int(2)] = d_455_v4_
                nw77_[int(3)] = d_455_v4_
                nw77_[int(4)] = d_455_v4_
                nw77_[int(5)] = d_455_v4_
                nw77_[int(6)] = d_455_v4_
                nw77_[int(7)] = d_455_v4_
                nw77_[int(8)] = d_455_v4_
                nw77_[int(9)] = d_455_v4_
                nw77_[int(10)] = d_455_v4_
                nw77_[int(11)] = d_455_v4_
                nw77_[int(12)] = d_455_v4_
                nw77_[int(13)] = d_455_v4_
                nw77_[int(14)] = d_455_v4_
                nw77_[int(15)] = d_455_v4_
                nw77_[int(16)] = d_455_v4_
                nw77_[int(17)] = d_455_v4_
                nw77_[int(18)] = d_455_v4_
                nw77_[int(19)] = d_455_v4_
                nw77_[int(20)] = d_455_v4_
                nw77_[int(21)] = d_455_v4_
                nw77_[int(22)] = d_455_v4_
                nw77_[int(23)] = d_455_v4_
                nw77_[int(24)] = d_455_v4_
                nw77_[int(25)] = d_455_v4_
                nw77_[int(26)] = d_455_v4_
                nw77_[int(27)] = d_455_v4_
                nw77_[int(28)] = d_455_v4_
                d_456_v5_ = nw77_
                d_457_v6_: _dafny.Map
                d_457_v6_ = _dafny.Map({d_452_v1_: d_456_v5_})
                index79_ = default__.safeIndex(134, (d_454_v3_).length(0))
                (d_454_v3_)[index79_] = d_457_v6_
                index80_ = default__.safeIndex(134, (d_454_v3_).length(0))
                (d_454_v3_)[index80_] = d_457_v6_
                d_458_v7_: str
                d_458_v7_ = _dafny.CodePoint('r')
                d_459_v8_: bool
                d_460_v9_: int
                d_461_v10_: bool
                d_462_v11_: bool
                out3_: bool
                out4_: int
                out5_: bool
                out6_: bool
                out3_, out4_, out5_, out6_ = (d_455_v4_).m8((((d_455_v4_).f13).set(default__.safeIndex(p0, len((d_455_v4_).f13)), d_458_v7_)) < ((d_455_v4_).f13), globalState)
                d_459_v8_ = out3_
                d_460_v9_ = out4_
                d_461_v10_ = out5_
                d_462_v11_ = out6_
                d_463_v12_: _dafny.Map
                d_463_v12_ = _dafny.Map({d_460_v9_: (False if (d_455_v4_).f12 else (d_452_v1_)[default__.safeIndex(765, (d_452_v1_).length(0))])})
                d_463_v12_ = d_463_v12_
            elif True:
                (globalState).f0 = p0
                d_464_v13_: _dafny.Map
                d_464_v13_ = _dafny.Map({p2: p2})
                d_465_v14_: _dafny.Seq
                d_465_v14_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({p2: (self).f6}), d_464_v13_, d_464_v13_, d_464_v13_])
                d_466_v15_: _dafny.Map
                d_466_v15_ = _dafny.Map({False: p0})
                d_467_v16_: _dafny.Seq
                d_467_v16_ = _dafny.SeqWithoutIsStrInference([p2, (self).f6])
                d_468_v17_: _dafny.Map
                d_468_v17_ = _dafny.Map({(d_464_v13_) == ((d_465_v14_)[default__.safeIndex(p0, len(d_465_v14_))]): ((d_466_v15_)[(self).f6] if ((self).f6) in (d_466_v15_) else len(d_467_v16_))})
                d_468_v17_ = d_466_v15_
                r3 = p2
                (globalState).f0 = p0
                d_469_v18_: _dafny.Array
                nw78_ = _dafny.Array(None, 15)
                d_469_v18_ = nw78_
                d_470_v19_: T0
                nw79_ = C2()
                nw79_.ctor__((self).f5)
                d_470_v19_ = nw79_
                index81_ = default__.safeIndex(949, (d_469_v18_).length(0))
                (d_469_v18_)[index81_] = d_470_v19_
                index82_ = default__.safeIndex(949, (d_469_v18_).length(0))
                (d_469_v18_)[index82_] = (d_470_v19_ if False else d_470_v19_)
            if not(p2):
                d_471_v20_: _dafny.Seq
                d_471_v20_ = _dafny.SeqWithoutIsStrInference([(self).f6])
                d_472_v21_: C2
                nw80_ = C2()
                nw80_.ctor__(_dafny.MultiSet(d_471_v20_))
                d_472_v21_ = nw80_
                d_473_v22_: _dafny.Map
                out7_: _dafny.Map
                out7_ = (d_472_v21_).m9((p2) and ((self).f6), p0, globalState)
                d_473_v22_ = out7_
                d_474_v23_: _dafny.Seq
                d_474_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "edhl"))
                d_475_v24_: str
                d_475_v24_ = _dafny.CodePoint('q')
                d_476_v25_: _dafny.Map
                d_476_v25_ = _dafny.Map({d_474_v23_: (self).f6})
                (globalState).f0 = (self).fm4(default__.safeDivisionInt(len(d_474_v23_), 180), (p0) - (p0), (_dafny.Map({(d_474_v23_).set(default__.safeIndex(p0, len(d_474_v23_)), d_475_v24_): False}) if (self).fm3(globalState) else d_476_v25_), (((self).f5)[p2] if (p2) in ((self).f5) else p0), globalState)
                d_477_v26_: D1
                d_477_v26_ = D1_DC2((self).f6)
                d_477_v26_ = p3
                d_478_v27_: C2
                nw81_ = C2()
                nw81_.ctor__(_dafny.MultiSet([p2, (self).fm3(globalState), (self).fm2((self).f6, globalState)]))
                d_478_v27_ = nw81_
            elif True:
                d_479_v28_: _dafny.Seq
                d_479_v28_ = _dafny.SeqWithoutIsStrInference([(self).f6])
                d_480_v29_: _dafny.Map
                d_480_v29_ = _dafny.Map({d_479_v28_: p2})
                d_481_v30_: _dafny.Seq
                d_481_v30_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mflyhpk"))
                d_482_v31_: _dafny.Map
                d_482_v31_ = _dafny.Map({not((self).fm1(p2, p0, 46, globalState)): (self).f6})
                d_483_v32_: C1
                nw82_ = C1()
                nw82_.ctor__((self).f6, d_481_v30_, ((self).f6) not in (d_482_v31_), (self).f5)
                d_483_v32_ = nw82_
                rhs86_ = (default__.fm20(p0, globalState)).set(d_479_v28_, p2)
                rhs87_ = d_483_v32_
                d_480_v29_ = rhs86_
                d_483_v32_ = rhs87_
                d_484_v33_: _dafny.Array
                def lambda22_(d_485_i0_):
                    return (self).f6

                init12_ = lambda22_
                nw83_ = _dafny.Array(None, 22)
                for i0_12_ in range(nw83_.length(0)):
                    nw83_[i0_12_] = init12_(i0_12_)
                d_484_v33_ = nw83_
                d_486_v34_: _dafny.Set
                d_486_v34_ = _dafny.Set({d_484_v33_})
                d_486_v34_ = d_486_v34_
                d_487_v35_: _dafny.Map
                d_487_v35_ = _dafny.Map({(d_483_v32_).f13: (p0) - ((0) - (len((d_483_v32_).f13)))})
                d_488_v37_: _dafny.Map
                d_488_v37_ = _dafny.Map({p0: p0})
                def iife69_():
                    coll29_ = _dafny.Map()
                    compr_29_: int
                    for compr_29_ in (d_488_v37_).keys.Elements:
                        d_489_v36_: int = compr_29_
                        if (d_489_v36_) in (d_488_v37_):
                            coll29_[default__.safeModuloInt(d_489_v36_, p0)] = (d_483_v32_).f12
                    return _dafny.Map(coll29_)
                d_487_v35_ = (d_487_v35_).set(((d_483_v32_).f13) + (default__.fm11((d_483_v32_).f12, globalState)), (0) - ((len(_dafny.Set({iife69_()
                })) if p2 else (0) - (p0))))
                d_490_v38_: _dafny.Array
                nw84_ = _dafny.Array(_dafny.Array(None, 0), 7)
                d_490_v38_ = nw84_
                d_491_v39_: D3
                d_491_v39_ = D3_DC10(not(False), -542)
                d_492_v40_: D0
                d_492_v40_ = D0_DC0(p0)
                pat_let_tv14_ = p0
                pat_let_tv15_ = p0
                d_493_v41_: _dafny.Array
                nw85_ = _dafny.Array(None, 21)
                nw85_[int(0)] = d_491_v39_
                nw85_[int(1)] = d_491_v39_
                nw85_[int(2)] = d_491_v39_
                nw85_[int(3)] = D3_DC10(p2, 816)
                nw85_[int(4)] = d_491_v39_
                nw85_[int(5)] = d_491_v39_
                nw85_[int(6)] = d_491_v39_
                def iife70_(_pat_let20_0):
                    def iife71_(d_494_dt__update__tmp_h0_):
                        def iife72_(_pat_let21_0):
                            def iife73_(d_495_dt__update_hcf12_h0_):
                                return D3_DC10((d_494_dt__update__tmp_h0_).cf11, d_495_dt__update_hcf12_h0_)
                            return iife73_(_pat_let21_0)
                        return iife72_(pat_let_tv14_)
                    return iife71_(_pat_let20_0)
                nw85_[int(7)] = iife70_(default__.fm21((d_483_v32_).f12, globalState))
                nw85_[int(8)] = d_491_v39_
                nw85_[int(9)] = D3_DC10((d_483_v32_).f12, p0)
                nw85_[int(10)] = d_491_v39_
                nw85_[int(11)] = d_491_v39_
                nw85_[int(12)] = d_491_v39_
                nw85_[int(13)] = d_491_v39_
                nw85_[int(14)] = d_491_v39_
                nw85_[int(15)] = d_491_v39_
                nw85_[int(16)] = D3_DC10(p2, p0)
                nw85_[int(17)] = d_491_v39_
                nw85_[int(18)] = d_491_v39_
                def iife74_(_pat_let22_0):
                    def iife75_(d_496_dt__update__tmp_h1_):
                        def iife76_(_pat_let23_0):
                            def iife77_(d_497_dt__update_hcf12_h1_):
                                return D3_DC10((d_496_dt__update__tmp_h1_).cf11, d_497_dt__update_hcf12_h1_)
                            return iife77_(_pat_let23_0)
                        return iife76_(pat_let_tv15_)
                    return iife75_(_pat_let22_0)
                nw85_[int(19)] = iife74_(d_491_v39_)
                nw85_[int(20)] = D3_DC10(p2, (d_492_v40_).cf0)
                d_493_v41_ = nw85_
                index83_ = default__.safeIndex(606, (d_490_v38_).length(0))
                (d_490_v38_)[index83_] = d_493_v41_
                index84_ = default__.safeIndex(606, (d_490_v38_).length(0))
                (d_490_v38_)[index84_] = (d_493_v41_ if not(False) else d_493_v41_)
                d_484_v33_ = d_484_v33_
            d_498_v42_: _dafny.Seq
            d_498_v42_ = _dafny.SeqWithoutIsStrInference([(self).f6])
            d_499_v43_: _dafny.Set
            d_499_v43_ = _dafny.Set({len(_dafny.Set({p2, (self).f6})), (_dafny.MultiSet(d_498_v42_)).cardinality, p0, p0, p0})
            d_500_v44_: _dafny.Map
            d_500_v44_ = _dafny.Map({(self).f6: (self).f6})
            d_501_v45_: _dafny.Set
            d_501_v45_ = _dafny.Set({d_451_v0_})
            d_502_v46_: str
            d_502_v46_ = _dafny.CodePoint('n')
            d_503_v47_: _dafny.Seq
            d_503_v47_ = _dafny.SeqWithoutIsStrInference([d_502_v46_, d_502_v46_])
            d_504_v48_: _dafny.Map
            d_504_v48_ = _dafny.Map({len(d_503_v47_): (self).f6})
            d_505_v49_: _dafny.Seq
            d_505_v49_ = _dafny.SeqWithoutIsStrInference([(0) - (p0), p0])
            d_506_v50_: _dafny.Array
            nw86_ = _dafny.Array(None, 22)
            nw86_[int(0)] = not(True)
            nw86_[int(1)] = p2
            nw86_[int(2)] = not ((self).f6) or ((self).f6)
            nw86_[int(3)] = ((self).f6) == (p2)
            nw86_[int(4)] = (self).f6
            nw86_[int(5)] = (self).f6
            nw86_[int(6)] = p2
            nw86_[int(7)] = p2
            nw86_[int(8)] = (d_499_v43_).ispropersubset(_dafny.Set({-874}))
            nw86_[int(9)] = (self).f6
            nw86_[int(10)] = p2
            nw86_[int(11)] = p2
            nw86_[int(12)] = p2
            nw86_[int(13)] = p2
            nw86_[int(14)] = (default__.fm21(not((self).f6), globalState)).cf11
            nw86_[int(15)] = True
            nw86_[int(16)] = ((d_500_v44_)[not(False)] if (not(False)) in (d_500_v44_) else (self).f6)
            nw86_[int(17)] = (_dafny.Set({d_451_v0_, d_451_v0_})).issubset(d_501_v45_)
            nw86_[int(18)] = (self).fm3(globalState)
            nw86_[int(19)] = (len(d_504_v48_)) <= ((d_505_v49_)[default__.safeIndex(p0, len(d_505_v49_))])
            nw86_[int(20)] = p2
            nw86_[int(21)] = p2
            d_506_v50_ = nw86_
            index85_ = default__.safeIndex(422, (d_506_v50_).length(0))
            (d_506_v50_)[index85_] = p2
            index86_ = default__.safeIndex(422, (d_506_v50_).length(0))
            (d_506_v50_)[index86_] = False
            d_507_v51_: _dafny.Array
            def lambda23_(d_508_v47_, d_509_v50_, d_510_p0_):
                def lambda24_(d_511_i1_):
                    return (_dafny.Map({(self).f6: len(d_508_v47_)})) | (_dafny.Map({(d_509_v50_)[default__.safeIndex(422, (d_509_v50_).length(0))]: d_510_p0_}))

                return lambda24_

            init13_ = lambda23_(d_503_v47_, d_506_v50_, p0)
            nw87_ = _dafny.Array(None, 28)
            for i0_13_ in range(nw87_.length(0)):
                nw87_[i0_13_] = init13_(i0_13_)
            d_507_v51_ = nw87_
            d_512_v52_: _dafny.Map
            d_512_v52_ = _dafny.Map({not((d_506_v50_)[default__.safeIndex(422, (d_506_v50_).length(0))]): p0})
            index87_ = default__.safeIndex(742, (d_507_v51_).length(0))
            (d_507_v51_)[index87_] = d_512_v52_
            index88_ = default__.safeIndex(742, (d_507_v51_).length(0))
            def iife78_():
                coll30_ = _dafny.Set()
                compr_30_: int
                for compr_30_ in _dafny.IntegerRange(-695, 274):
                    d_513_v53_: int = compr_30_
                    if ((-695) <= (d_513_v53_)) and ((d_513_v53_) < (274)):
                        coll30_ = coll30_.union(_dafny.Set([(d_513_v53_) - (p0)]))
                return _dafny.Set(coll30_)
            rhs88_ = (d_499_v43_ if (self).f6 else (d_499_v43_) - (iife78_()
            ))
            rhs89_ = (self).fm3(globalState)
            rhs90_ = ((d_506_v50_)[default__.safeIndex(422, (d_506_v50_).length(0))] if (p0) > (p0) else p2)
            rhs91_ = ((d_512_v52_) | (default__.fm19(p0, globalState))).set(((d_506_v50_)[default__.safeIndex(422, (d_506_v50_).length(0))]) and ((self).f6), (438) * ((((self).f5).set((d_506_v50_)[default__.safeIndex(422, (d_506_v50_).length(0))], default__.abs(p0))).cardinality))
            lhs59_ = d_507_v51_
            lhs60_ = default__.safeIndex(742, (d_507_v51_).length(0))
            d_499_v43_ = rhs88_
            r3 = rhs89_
            r3 = rhs90_
            lhs59_[lhs60_] = rhs91_
            index89_ = default__.safeIndex(422, (d_506_v50_).length(0))
            (d_506_v50_)[index89_] = (self).fm1(False, p0, default__.safeDivisionInt(p0, p0), globalState)
        elif True:
            d_514_v54_: _dafny.Array
            nw88_ = _dafny.Array(None, 8)
            nw88_[int(0)] = not(not(p2))
            nw88_[int(1)] = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_515_i2_ in range(default__.abs(824))])) < (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yaxq")))
            nw88_[int(2)] = (self).f6
            nw88_[int(3)] = (p0) != (p0)
            nw88_[int(4)] = p2
            nw88_[int(5)] = (p0) != (p0)
            nw88_[int(6)] = p2
            nw88_[int(7)] = (self).f6
            d_514_v54_ = nw88_
            index90_ = default__.safeIndex(725, (d_514_v54_).length(0))
            (d_514_v54_)[index90_] = (self).f6
            d_516_v55_: _dafny.Seq
            d_516_v55_ = _dafny.SeqWithoutIsStrInference([p0, p0, -775, p0])
            d_517_v56_: _dafny.Seq
            d_517_v56_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jrojrlvjv"))
            d_518_v57_: _dafny.Map
            d_518_v57_ = _dafny.Map({d_517_v56_: (self).f6})
            d_519_v58_: _dafny.Set
            d_519_v58_ = _dafny.Set({p2, (self).f6})
            index91_ = default__.safeIndex(725, (d_514_v54_).length(0))
            rhs92_ = (self).fm4(default__.safeModuloInt((d_516_v55_)[default__.safeIndex(p0, len(d_516_v55_))], (0) - (p0)), 512, (d_518_v57_) | (d_518_v57_), 108, globalState)
            rhs93_ = p2
            rhs94_ = not((d_519_v58_).isdisjoint(d_519_v58_))
            lhs61_ = globalState
            lhs62_ = d_514_v54_
            lhs63_ = default__.safeIndex(725, (d_514_v54_).length(0))
            lhs61_.f0 = rhs92_
            lhs62_[lhs63_] = rhs93_
            r3 = rhs94_
            (globalState).f0 = p0
            rhs95_ = p0
            rhs96_ = ((len(d_519_v58_)) + (p0)) * (p0)
            lhs64_ = globalState
            lhs65_ = globalState
            lhs64_.f0 = rhs95_
            lhs65_.f0 = rhs96_
            d_520_v59_: str
            d_520_v59_ = _dafny.CodePoint('w')
            d_520_v59_ = d_520_v59_
            d_521_v60_: _dafny.Map
            d_521_v60_ = _dafny.Map({_dafny.Set({p0, p0, p0, 966}): _dafny.SeqWithoutIsStrInference([(d_514_v54_)[default__.safeIndex(725, (d_514_v54_).length(0))], True])})
            d_522_v61_: _dafny.Map
            d_522_v61_ = _dafny.Map({d_521_v60_: ((d_514_v54_)[default__.safeIndex(725, (d_514_v54_).length(0))]) and (p2)})
            d_522_v61_ = (d_522_v61_).set(d_521_v60_, not(not (not((self).f6)) or (False)))
        if (self).f6:
            d_523_v62_: _dafny.MultiSet
            d_523_v62_ = _dafny.MultiSet([(self).f6])
            d_523_v62_ = ((self).f5) - ((self).f5)
            d_524_v63_: str
            d_524_v63_ = _dafny.CodePoint('t')
            d_525_v64_: D2
            d_525_v64_ = D2_DC7(d_524_v63_, _dafny.Map({p0: p0}), p0)
            d_526_v65_: D2
            d_526_v65_ = D2_DC8(d_525_v64_)
            d_527_v66_: _dafny.Map
            d_527_v66_ = _dafny.Map({d_526_v65_: p0})
            d_528_v67_: _dafny.Map
            d_528_v67_ = _dafny.Map({(d_527_v66_).set(d_526_v65_, p0): p0})
            (globalState).f0 = ((d_528_v67_)[d_527_v66_] if (d_527_v66_) in (d_528_v67_) else p0)
            r2 = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ysdg"))) + (default__.fm11(p2, globalState))
            r3 = (self).f6
            r3 = (self).f6
        elif True:
            d_529_v68_: D3
            d_529_v68_ = D3_DC10(p2, p0)
            d_530_v69_: _dafny.Seq
            d_530_v69_ = _dafny.SeqWithoutIsStrInference([(d_529_v68_).cf12])
            d_530_v69_ = d_530_v69_
            d_531_v70_: C2
            nw89_ = C2()
            nw89_.ctor__((self).f5)
            d_531_v70_ = nw89_
            r3 = (self).f6
            d_532_v71_: _dafny.Map
            d_532_v71_ = _dafny.Map({622: ((self).f5).cardinality})
            r3 = ((_dafny.MultiSet([(d_531_v70_).fm1((self).f6, len(d_532_v71_), p0, globalState), p2])).issubset((self).f5)) == (not((self).f6))
            index92_ = default__.safeIndex(885, (d_451_v0_).length(0))
            (d_451_v0_)[index92_] = p0
            d_533_v72_: _dafny.Map
            d_533_v72_ = _dafny.Map({(self).f6: p2})
            d_534_v73_: _dafny.Seq
            d_534_v73_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qgee"))
            d_535_v74_: _dafny.Map
            d_535_v74_ = _dafny.Map({157: d_534_v73_})
            index93_ = default__.safeIndex(885, (d_451_v0_).length(0))
            rhs97_ = (default__.safeModuloInt(p0, len(d_533_v72_))) + (p0)
            rhs98_ = len(((d_535_v74_)[(p0 if (self).f6 else p0)] if ((p0 if (self).f6 else p0)) in (d_535_v74_) else (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_536_i3_ in range(default__.abs(666))])) + (d_534_v73_)))
            rhs99_ = (self).fm2(p2, globalState)
            lhs66_ = globalState
            lhs67_ = d_451_v0_
            lhs68_ = default__.safeIndex(885, (d_451_v0_).length(0))
            lhs66_.f0 = rhs97_
            lhs67_[lhs68_] = rhs98_
            r3 = rhs99_
        r3 = (self).f6
        index94_ = default__.safeIndex(85, (d_451_v0_).length(0))
        (d_451_v0_)[index94_] = p0
        d_537_v75_: _dafny.Seq
        d_537_v75_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "frah"))
        d_538_v76_: _dafny.Map
        d_538_v76_ = _dafny.Map({d_537_v75_: not(p2)})
        d_539_v77_: str
        d_539_v77_ = _dafny.CodePoint('t')
        index95_ = default__.safeIndex(85, (d_451_v0_).length(0))
        rhs100_ = (self).fm4(p0, (0) - (p0), d_538_v76_, p0, globalState)
        rhs101_ = (((self).f5)[(self).f6] if ((self).f6) in ((self).f5) else len((d_537_v75_).set(default__.safeIndex(len(d_537_v75_), len(d_537_v75_)), d_539_v77_)))
        lhs69_ = globalState
        lhs70_ = d_451_v0_
        lhs71_ = default__.safeIndex(85, (d_451_v0_).length(0))
        lhs69_.f0 = rhs100_
        lhs70_[lhs71_] = rhs101_
        r0 = default__.fm22(globalState)
        r1 = (d_537_v75_) + (d_537_v75_)
        r2 = (d_537_v75_) + (d_537_v75_)
        r3 = (self).f6
        return r0, r1, r2, r3


class C4(T2, T1, T0):
    def  __init__(self):
        self._f7: _dafny.Array = _dafny.Array(None, 0)
        self._f8: bool = False
        self._f6: bool = False
        self._f5: _dafny.MultiSet = _dafny.MultiSet({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    @property
    def f7(self):
        return self._f7
    @f7.setter
    def f7(self, value):
        self._f7 = value
    @property
    def f8(self):
        return self._f8
    @f8.setter
    def f8(self, value):
        self._f8 = value
    @property
    def f6(self):
        return self._f6
    @property
    def f5(self):
        return self._f5
    def ctor__(self, f7, f8, f6, f5):
        (self).f7 = f7
        (self).f8 = f8
        (self)._f6 = f6
        (self)._f5 = f5

    def fm3(self, globalState):
        return ((_dafny.SeqWithoutIsStrInference([False])) + (_dafny.SeqWithoutIsStrInference([self.f8, (self).f6, self.f8, self.f8, self.f8]))) <= (_dafny.SeqWithoutIsStrInference([self.f8, (self).f6, (self).f6, (self).f6, self.f8]))

    def fm4(self, p0, p1, p2, p3, globalState):
        return -256

    def fm1(self, p0, p1, p2, globalState):
        return self.f8

    def fm2(self, p0, globalState):
        return (self).f6

    def fm6(self, globalState):
        return (32) - ((0) - ((((self).f5)[self.f8] if (self.f8) in ((self).f5) else -44)))

    def fm7(self, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lsi"))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jxsfieg"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "isrelfj"))))

    def m2(self, p0, globalState):
        d_540_v0_: str
        d_540_v0_ = _dafny.CodePoint('o')
        rhs102_ = False
        rhs103_ = (self).fm6(globalState)
        rhs104_ = d_540_v0_
        lhs72_ = self
        lhs73_ = globalState
        lhs72_.f8 = rhs102_
        lhs73_.f0 = rhs103_
        d_540_v0_ = rhs104_
        d_541_v1_: _dafny.Map
        d_541_v1_ = _dafny.Map({(self).f6: self.f8})
        d_541_v1_ = (d_541_v1_).set(((self).f6) and ((self).f6), (self).f6)
        d_542_v2_: int
        d_542_v2_ = 725
        d_543_v3_: D0
        d_543_v3_ = D0_DC1()
        d_544_v4_: _dafny.Seq
        d_544_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "elhkbkopj"))
        d_545_v5_: _dafny.Seq
        d_545_v5_ = _dafny.SeqWithoutIsStrInference([d_544_v4_, d_544_v4_])
        d_546_v6_: _dafny.Array
        nw90_ = _dafny.Array(None, 18)
        nw90_[int(0)] = d_542_v2_
        nw90_[int(1)] = len(d_544_v4_)
        nw90_[int(2)] = d_542_v2_
        nw90_[int(3)] = d_542_v2_
        nw90_[int(4)] = 365
        nw90_[int(5)] = d_542_v2_
        nw90_[int(6)] = (self).fm6(globalState)
        nw90_[int(7)] = d_542_v2_
        nw90_[int(8)] = len(d_545_v5_)
        nw90_[int(9)] = d_542_v2_
        nw90_[int(10)] = d_542_v2_
        nw90_[int(11)] = d_542_v2_
        nw90_[int(12)] = len(d_544_v4_)
        nw90_[int(13)] = len(_dafny.SeqWithoutIsStrInference([d_542_v2_, d_542_v2_]))
        nw90_[int(14)] = d_542_v2_
        nw90_[int(15)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bjayaggc")))
        nw90_[int(16)] = d_542_v2_
        nw90_[int(17)] = len(default__.fm8(d_542_v2_, False, d_544_v4_, globalState))
        d_546_v6_ = nw90_
        d_547_v7_: _dafny.Map
        d_547_v7_ = _dafny.Map({d_543_v3_: d_546_v6_})
        d_548_v8_: _dafny.Map
        d_548_v8_ = _dafny.Map({d_542_v2_: len(d_547_v7_)})
        d_549_v9_: _dafny.Set
        d_549_v9_ = _dafny.Set({(self).fm4(d_542_v2_, d_542_v2_, _dafny.Map({d_544_v4_: self.f8}), d_542_v2_, globalState), d_542_v2_})
        d_550_v10_: _dafny.Array
        nw91_ = _dafny.Array(None, 27)
        nw91_[int(0)] = d_542_v2_
        nw91_[int(1)] = (d_542_v2_) * ((0) - (d_542_v2_))
        nw91_[int(2)] = len(d_548_v8_)
        nw91_[int(3)] = d_542_v2_
        nw91_[int(4)] = 737
        nw91_[int(5)] = d_542_v2_
        nw91_[int(6)] = 115
        nw91_[int(7)] = default__.safeModuloInt(len(d_544_v4_), d_542_v2_)
        nw91_[int(8)] = d_542_v2_
        nw91_[int(9)] = d_542_v2_
        nw91_[int(10)] = len(_dafny.Map({len(d_549_v9_): d_542_v2_}))
        nw91_[int(11)] = -640
        nw91_[int(12)] = (0) - ((d_542_v2_) + (d_542_v2_))
        nw91_[int(13)] = d_542_v2_
        nw91_[int(14)] = d_542_v2_
        nw91_[int(15)] = default__.safeDivisionInt(d_542_v2_, d_542_v2_)
        nw91_[int(16)] = d_542_v2_
        nw91_[int(17)] = ((self).f5).cardinality
        nw91_[int(18)] = d_542_v2_
        nw91_[int(19)] = len(_dafny.Map({self.f8: d_542_v2_}))
        nw91_[int(20)] = d_542_v2_
        nw91_[int(21)] = len(default__.fm9(globalState))
        nw91_[int(22)] = d_542_v2_
        nw91_[int(23)] = d_542_v2_
        nw91_[int(24)] = (0) - (d_542_v2_)
        nw91_[int(25)] = 518
        nw91_[int(26)] = d_542_v2_
        d_550_v10_ = nw91_
        d_550_v10_ = d_546_v6_
        d_551_i0_: int
        d_551_i0_ = 0
        with _dafny.label("7"):
            while (d_542_v2_) != ((len(d_544_v4_)) - ((0) - (d_542_v2_))):
                with _dafny.c_label("7"):
                    if (d_551_i0_) >= (100):
                        raise _dafny.Break("7")
                    d_551_i0_ = (d_551_i0_) + (1)
                    if (self).f6:
                        d_552_v12_: _dafny.Map
                        d_552_v12_ = _dafny.Map({d_543_v3_: p0})
                        d_553_v13_: _dafny.MultiSet
                        d_553_v13_ = _dafny.MultiSet([_dafny.Map({d_542_v2_: len(d_552_v12_)}), d_548_v8_])
                        d_554_v15_: _dafny.Seq
                        d_554_v15_ = _dafny.SeqWithoutIsStrInference([d_548_v8_, d_548_v8_])
                        d_555_v16_: C0
                        nw92_ = C0()
                        def iife79_():
                            coll31_ = _dafny.Map()
                            compr_31_: _dafny.Map
                            for compr_31_ in (d_553_v13_).Elements:
                                d_556_v11_: _dafny.Map = compr_31_
                                if (d_556_v11_) in (d_553_v13_):
                                    coll31_[d_556_v11_] = False
                            return _dafny.Map(coll31_)
                        def iife80_():
                            coll32_ = _dafny.Map()
                            compr_32_: _dafny.Map
                            for compr_32_ in (d_554_v15_).Elements:
                                d_557_v14_: _dafny.Map = compr_32_
                                if (d_557_v14_) in (d_554_v15_):
                                    coll32_[d_557_v14_] = self.f8
                            return _dafny.Map(coll32_)
                        nw92_.ctor__((iife79_()
                         if not(not(self.f8)) else iife80_()
                        ))
                        d_555_v16_ = nw92_
                        d_558_v17_: T1
                        nw93_ = C3()
                        nw93_.ctor__((self).f6, _dafny.MultiSet([(self).f6]))
                        d_558_v17_ = nw93_
                        d_559_v18_: _dafny.Set
                        d_559_v18_ = _dafny.Set({d_558_v17_, d_558_v17_})
                        arr0_ = self.f7
                        index96_ = default__.safeIndex(590, (self.f7).length(0))
                        arr0_[index96_] = self.f8
                        d_560_v19_: _dafny.Map
                        d_560_v19_ = _dafny.Map({913: not(((d_558_v17_).f6 if False else (self).f6))})
                        d_561_v20_: _dafny.Seq
                        d_561_v20_ = _dafny.SeqWithoutIsStrInference([d_542_v2_])
                        arr1_ = self.f7
                        index97_ = default__.safeIndex(590, (self.f7).length(0))
                        rhs105_ = _dafny.Set({d_558_v17_, d_558_v17_, d_558_v17_})
                        rhs106_ = ((d_560_v19_)[len((d_544_v4_) + ((d_544_v4_).set(default__.safeIndex(d_542_v2_, len(d_544_v4_)), d_540_v0_)))] if (len((d_544_v4_) + ((d_544_v4_).set(default__.safeIndex(d_542_v2_, len(d_544_v4_)), d_540_v0_)))) in (d_560_v19_) else (d_542_v2_) <= (len(d_561_v20_)))
                        rhs107_ = (d_558_v17_).f6
                        rhs108_ = d_543_v3_
                        lhs74_ = self
                        lhs75_ = self.f7
                        lhs76_ = default__.safeIndex(590, (self.f7).length(0))
                        d_559_v18_ = rhs105_
                        lhs74_.f8 = rhs106_
                        lhs75_[lhs76_] = rhs107_
                        d_543_v3_ = rhs108_
                        arr2_ = self.f7
                        index98_ = default__.safeIndex(590, (self.f7).length(0))
                        arr2_[index98_] = (d_558_v17_).f6
                        d_542_v2_ = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ulmddp")))
                        d_562_v21_: _dafny.Map
                        d_562_v21_ = _dafny.Map({d_544_v4_: self.f8})
                        (globalState).f0 = (self).fm4(d_542_v2_, d_542_v2_, d_562_v21_, d_542_v2_, globalState)
                    elif True:
                        d_563_v22_: _dafny.MultiSet
                        d_563_v22_ = _dafny.MultiSet([len(d_544_v4_)])
                        (self).f8 = (d_563_v22_).issubset(d_563_v22_)
                        arr3_ = self.f7
                        index99_ = default__.safeIndex(393, (self.f7).length(0))
                        arr3_[index99_] = (self).f6
                        d_564_v23_: _dafny.Seq
                        d_564_v23_ = _dafny.SeqWithoutIsStrInference([len(d_544_v4_)])
                        arr4_ = self.f7
                        index100_ = default__.safeIndex(393, (self.f7).length(0))
                        rhs109_ = (d_542_v2_) - (d_542_v2_)
                        rhs110_ = _dafny.CodePoint('g')
                        rhs111_ = (default__.safeModuloInt(len(d_564_v23_), ((self).f5).cardinality)) == (d_542_v2_)
                        rhs112_ = self.f7
                        rhs113_ = (0) - (len(default__.fm19((0) - (d_542_v2_), globalState)))
                        lhs77_ = globalState
                        lhs78_ = self.f7
                        lhs79_ = default__.safeIndex(393, (self.f7).length(0))
                        lhs80_ = self
                        lhs77_.f0 = rhs109_
                        d_540_v0_ = rhs110_
                        lhs78_[lhs79_] = rhs111_
                        lhs80_.f7 = rhs112_
                        d_542_v2_ = rhs113_
                        d_565_v24_: _dafny.Map
                        d_565_v24_ = _dafny.Map({d_542_v2_: p0})
                        d_540_v0_ = ((d_565_v24_)[(0) - (d_542_v2_)] if ((0) - (d_542_v2_)) in (d_565_v24_) else p0)
                        d_566_v25_: _dafny.Map
                        d_566_v25_ = _dafny.Map({default__.fm23(d_542_v2_, globalState): (self).fm3(globalState)})
                        d_566_v25_ = d_566_v25_
                        rhs114_ = (_dafny.MultiSet((d_564_v23_) + (d_564_v23_))) | (d_563_v22_)
                        rhs115_ = (0) - ((d_542_v2_) * (d_542_v2_))
                        lhs81_ = globalState
                        d_563_v22_ = rhs114_
                        lhs81_.f0 = rhs115_
                    (globalState).f0 = (((_dafny.MultiSet([(self).f6]) if (self).f6 else (self).f5)).cardinality) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ckkg"))))
                    d_567_v26_: _dafny.Map
                    d_567_v26_ = _dafny.Map({self.f8: d_542_v2_})
                    d_568_v27_: D7
                    d_568_v27_ = D7_DC19(d_567_v26_)
                    d_568_v27_ = D7_DC19(d_567_v26_)
                    d_569_v28_: C2
                    nw94_ = C2()
                    nw94_.ctor__((self).f5)
                    d_569_v28_ = nw94_
                    d_570_v29_: D9
                    d_570_v29_ = D9_DC26(d_569_v28_)
                    d_569_v28_ = (d_570_v29_).cf41
                    pass
            pass
        d_571_v30_: D3
        d_571_v30_ = D3_DC10(False, d_542_v2_)
        d_572_v31_: C1
        nw95_ = C1()
        nw95_.ctor__(self.f8, default__.fm17(globalState), (self).fm1((d_571_v30_).cf11, d_542_v2_, len(d_544_v4_), globalState), ((self).f5) | (_dafny.MultiSet([not(self.f8)])))
        d_572_v31_ = nw95_
        d_550_v10_ = d_550_v10_

    def m3(self, globalState):
        r0: _dafny.Set = _dafny.Set({})
        r1: int = int(0)
        r2: int = int(0)
        (self).f8 = (self).f6
        if True:
            arr5_ = self.f7
            index101_ = default__.safeIndex(189, (self.f7).length(0))
            arr5_[index101_] = True
            d_573_v0_: _dafny.Array
            def lambda25_(d_574_i0_):
                return (d_574_i0_) + (-680)

            init14_ = lambda25_
            nw96_ = _dafny.Array(None, 4)
            for i0_14_ in range(nw96_.length(0)):
                nw96_[i0_14_] = init14_(i0_14_)
            d_573_v0_ = nw96_
            d_575_v1_: int
            d_575_v1_ = 345
            index102_ = default__.safeIndex(443, (d_573_v0_).length(0))
            (d_573_v0_)[index102_] = (d_575_v1_) * (d_575_v1_)
            d_576_v2_: _dafny.Seq
            d_576_v2_ = _dafny.SeqWithoutIsStrInference([(self).fm2((self).f6, globalState), (self).f6, (self).f6, self.f8, self.f8])
            arr6_ = self.f7
            index103_ = default__.safeIndex(189, (self.f7).length(0))
            index104_ = default__.safeIndex(443, (d_573_v0_).length(0))
            rhs116_ = self.f7
            rhs117_ = (self).fm3(globalState)
            rhs118_ = (len(d_576_v2_)) == ((self).fm6(globalState))
            rhs119_ = d_575_v1_
            lhs82_ = self
            lhs83_ = self
            lhs84_ = self.f7
            lhs85_ = default__.safeIndex(189, (self.f7).length(0))
            lhs86_ = d_573_v0_
            lhs87_ = default__.safeIndex(443, (d_573_v0_).length(0))
            lhs82_.f7 = rhs116_
            lhs83_.f8 = rhs117_
            lhs84_[lhs85_] = rhs118_
            lhs86_[lhs87_] = rhs119_
            d_576_v2_ = d_576_v2_
            d_577_v3_: _dafny.Array
            def lambda26_(d_578_i1_):
                return False

            init15_ = lambda26_
            nw97_ = _dafny.Array(None, 18)
            for i0_15_ in range(nw97_.length(0)):
                nw97_[i0_15_] = init15_(i0_15_)
            d_577_v3_ = nw97_
            d_579_v4_: _dafny.MultiSet
            d_579_v4_ = _dafny.MultiSet([d_577_v3_])
            d_580_v5_: _dafny.Map
            d_580_v5_ = _dafny.Map({232: self.f8})
            d_581_v6_: _dafny.Map
            d_581_v6_ = _dafny.Map({d_579_v4_: d_580_v5_})
            d_581_v6_ = (d_581_v6_).set(d_579_v4_, _dafny.Map({d_575_v1_: (self).f6}))
            d_582_v7_: _dafny.Array
            def lambda27_(d_583_i2_):
                return _dafny.CodePoint('c')

            init16_ = lambda27_
            nw98_ = _dafny.Array(None, 28)
            for i0_16_ in range(nw98_.length(0)):
                nw98_[i0_16_] = init16_(i0_16_)
            d_582_v7_ = nw98_
            (globalState).f2 = d_582_v7_
            d_584_v8_: _dafny.Seq
            d_584_v8_ = _dafny.SeqWithoutIsStrInference([(d_573_v0_)[default__.safeIndex(443, (d_573_v0_).length(0))], len(d_576_v2_), d_575_v1_, 778])
            arr7_ = self.f7
            index105_ = default__.safeIndex(189, (self.f7).length(0))
            rhs120_ = d_573_v0_
            rhs121_ = (d_573_v0_)[default__.safeIndex(443, (d_573_v0_).length(0))]
            rhs122_ = ((d_584_v8_) + (d_584_v8_)) + (d_584_v8_)
            rhs123_ = (self).f6
            lhs88_ = globalState
            lhs89_ = self.f7
            lhs90_ = default__.safeIndex(189, (self.f7).length(0))
            d_573_v0_ = rhs120_
            lhs88_.f0 = rhs121_
            d_584_v8_ = rhs122_
            lhs89_[lhs90_] = rhs123_
        elif True:
            d_585_v9_: C2
            nw99_ = C2()
            nw99_.ctor__((self).f5)
            d_585_v9_ = nw99_
            (self).f8 = False
            d_586_v10_: str
            d_586_v10_ = _dafny.CodePoint('u')
            d_586_v10_ = d_586_v10_
            d_587_v11_: int
            d_587_v11_ = 442
            (globalState).f0 = d_587_v11_
            if not(False):
                d_588_v12_: _dafny.Seq
                d_588_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lh"))
                d_589_v13_: _dafny.Seq
                d_589_v13_ = _dafny.SeqWithoutIsStrInference([d_588_v12_])
                d_588_v12_ = ((d_589_v13_)[default__.safeIndex(d_587_v11_, len(d_589_v13_))]).set(default__.safeIndex((self).fm6(globalState), len((d_589_v13_)[default__.safeIndex(d_587_v11_, len(d_589_v13_))])), d_586_v10_)
                d_588_v12_ = d_588_v12_
                d_590_v14_: _dafny.Array
                nw100_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 29)
                d_590_v14_ = nw100_
                index106_ = default__.safeIndex(284, (d_590_v14_).length(0))
                (d_590_v14_)[index106_] = (d_588_v12_) + (d_588_v12_)
                index107_ = default__.safeIndex(284, (d_590_v14_).length(0))
                (d_590_v14_)[index107_] = (((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "px"))) + (d_588_v12_)).set(default__.safeIndex(d_587_v11_, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "px"))) + (d_588_v12_))), d_586_v10_)) + (d_588_v12_)
                d_586_v10_ = d_586_v10_
                d_591_v15_: _dafny.Seq
                d_591_v15_ = _dafny.SeqWithoutIsStrInference([(self).f6])
                d_592_v16_: _dafny.MultiSet
                d_592_v16_ = _dafny.MultiSet([d_591_v15_])
                d_593_v17_: D10
                d_593_v17_ = D10_DC29(d_592_v16_)
                d_594_v18_: _dafny.Array
                nw101_ = _dafny.Array(None, 17)
                nw101_[int(0)] = (d_592_v16_) - (d_592_v16_)
                nw101_[int(1)] = _dafny.MultiSet([d_591_v15_, _dafny.SeqWithoutIsStrInference([(self).f6])])
                nw101_[int(2)] = d_592_v16_
                nw101_[int(3)] = _dafny.MultiSet([d_591_v15_])
                nw101_[int(4)] = d_592_v16_
                nw101_[int(5)] = d_592_v16_
                nw101_[int(6)] = d_592_v16_
                nw101_[int(7)] = (_dafny.MultiSet([d_591_v15_])) | (d_592_v16_)
                nw101_[int(8)] = (d_592_v16_).set(d_591_v15_, default__.abs(d_587_v11_))
                nw101_[int(9)] = (d_592_v16_) | ((d_593_v17_).cf45)
                nw101_[int(10)] = d_592_v16_
                nw101_[int(11)] = (d_592_v16_).intersection(_dafny.MultiSet([d_591_v15_, d_591_v15_, d_591_v15_, d_591_v15_]))
                nw101_[int(12)] = (d_592_v16_) - (d_592_v16_)
                nw101_[int(13)] = d_592_v16_
                nw101_[int(14)] = d_592_v16_
                nw101_[int(15)] = _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([self.f8, (self).f6, True]) for d_595_i3_ in range(default__.abs(203))]))
                nw101_[int(16)] = d_592_v16_
                d_594_v18_ = nw101_
                d_596_v19_: _dafny.Seq
                d_596_v19_ = _dafny.SeqWithoutIsStrInference([d_591_v15_])
                index108_ = default__.safeIndex(423, (d_594_v18_).length(0))
                (d_594_v18_)[index108_] = _dafny.MultiSet(d_596_v19_)
                index109_ = default__.safeIndex(423, (d_594_v18_).length(0))
                (d_594_v18_)[index109_] = _dafny.MultiSet([d_591_v15_])
            elif True:
                arr8_ = self.f7
                index110_ = default__.safeIndex(240, (self.f7).length(0))
                arr8_[index110_] = self.f8
                d_597_v21_: D2
                def iife81_():
                    coll33_ = _dafny.Set()
                    compr_33_: int
                    for compr_33_ in _dafny.IntegerRange(851, -711):
                        d_598_v20_: int = compr_33_
                        if ((851) <= (d_598_v20_)) and ((d_598_v20_) < (-711)):
                            coll33_ = coll33_.union(_dafny.Set([(d_598_v20_) * (d_587_v11_)]))
                    return _dafny.Set(coll33_)
                d_597_v21_ = D2_DC4(iife81_()
)
                d_599_v22_: _dafny.Seq
                d_599_v22_ = _dafny.SeqWithoutIsStrInference([d_587_v11_, d_587_v11_])
                d_600_v23_: _dafny.Set
                d_600_v23_ = _dafny.Set({len(d_599_v22_), d_587_v11_})
                arr9_ = self.f7
                index111_ = default__.safeIndex(240, (self.f7).length(0))
                arr9_[index111_] = (d_600_v23_).issubset((d_597_v21_).cf4)
                d_601_v24_: _dafny.Seq
                d_601_v24_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kypgb"))])
                d_602_v25_: _dafny.Seq
                d_602_v25_ = _dafny.SeqWithoutIsStrInference([True])
                d_603_v26_: _dafny.Map
                d_603_v26_ = _dafny.Map({(d_602_v25_).set(default__.safeIndex(d_587_v11_, len(d_602_v25_)), self.f8): (d_601_v24_) + (d_601_v24_)})
                d_601_v24_ = ((d_603_v26_)[d_602_v25_] if (d_602_v25_) in (d_603_v26_) else _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ykalvw")) for d_604_i4_ in range(default__.abs(701))]))
                d_601_v24_ = d_601_v24_
                d_605_v27_: _dafny.Map
                d_605_v27_ = _dafny.Map({_dafny.MultiSet([(self.f7)[default__.safeIndex(240, (self.f7).length(0))]]): self.f8})
                d_605_v27_ = (d_605_v27_).set((self).f5, ((0) - (d_587_v11_)) < (d_587_v11_))
                d_606_v28_: _dafny.Map
                out8_: _dafny.Map
                out8_ = (d_585_v9_).m9(not((self).f6), d_587_v11_, globalState)
                d_606_v28_ = out8_
        d_607_v29_: _dafny.Seq
        d_607_v29_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_608_i5_ in range(default__.abs(250))])])
        arr10_ = self.f7
        index112_ = default__.safeIndex(183, (self.f7).length(0))
        arr10_[index112_] = (d_607_v29_) == (d_607_v29_)
        arr11_ = self.f7
        index113_ = default__.safeIndex(183, (self.f7).length(0))
        arr11_[index113_] = not((self).fm3(globalState))
        d_609_v30_: int
        d_609_v30_ = 967
        hi2_ = d_609_v30_
        for d_610_i6_ in range(d_609_v30_, hi2_):
            d_611_v31_: str
            d_611_v31_ = _dafny.CodePoint('g')
            d_612_v32_: _dafny.Map
            d_612_v32_ = _dafny.Map({d_609_v30_: ((self).f5).cardinality})
            arr12_ = self.f7
            index114_ = default__.safeIndex(183, (self.f7).length(0))
            arr12_[index114_] = ((d_610_i6_) + ((D2_DC7(d_611_v31_, d_612_v32_, -433)).cf8)) >= (d_609_v30_)
            d_613_v33_: _dafny.Array
            def lambda28_(d_614_i6_):
                def lambda29_(d_615_i7_):
                    return default__.safeModuloInt(d_615_i7_, d_614_i6_)

                return lambda29_

            init17_ = lambda28_(d_610_i6_)
            nw102_ = _dafny.Array(None, 29)
            for i0_17_ in range(nw102_.length(0)):
                nw102_[i0_17_] = init17_(i0_17_)
            d_613_v33_ = nw102_
            index115_ = default__.safeIndex(198, (d_613_v33_).length(0))
            (d_613_v33_)[index115_] = 728
            arr13_ = self.f7
            index116_ = default__.safeIndex(183, (self.f7).length(0))
            index117_ = default__.safeIndex(198, (d_613_v33_).length(0))
            rhs124_ = self.f8
            rhs125_ = (d_610_i6_) * (d_610_i6_)
            lhs91_ = self.f7
            lhs92_ = default__.safeIndex(183, (self.f7).length(0))
            lhs93_ = d_613_v33_
            lhs94_ = default__.safeIndex(198, (d_613_v33_).length(0))
            lhs91_[lhs92_] = rhs124_
            lhs93_[lhs94_] = rhs125_
            arr14_ = self.f7
            index118_ = default__.safeIndex(183, (self.f7).length(0))
            rhs126_ = (self.f7)[default__.safeIndex(183, (self.f7).length(0))]
            rhs127_ = (self.f7)[default__.safeIndex(183, (self.f7).length(0))]
            rhs128_ = (d_607_v29_ if self.f8 else d_607_v29_)
            lhs95_ = self
            lhs96_ = self.f7
            lhs97_ = default__.safeIndex(183, (self.f7).length(0))
            lhs95_.f8 = rhs126_
            lhs96_[lhs97_] = rhs127_
            d_607_v29_ = rhs128_
            d_616_v34_: _dafny.Seq
            d_616_v34_ = _dafny.SeqWithoutIsStrInference([d_609_v30_])
            d_617_v35_: _dafny.Seq
            d_617_v35_ = _dafny.SeqWithoutIsStrInference([d_616_v34_])
            d_618_v36_: _dafny.Seq
            d_618_v36_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hj"))
            d_619_v37_: _dafny.Map
            d_619_v37_ = _dafny.Map({(d_617_v35_)[default__.safeIndex(len(d_618_v36_), len(d_617_v35_))]: False})
            d_619_v37_ = d_619_v37_
        if False:
            (globalState).f0 = d_609_v30_
            d_620_v38_: _dafny.Array
            nw103_ = _dafny.Array(False, 28)
            d_620_v38_ = nw103_
            d_621_v39_: _dafny.Map
            d_621_v39_ = _dafny.Map({d_620_v38_: (self).f6})
            d_622_v40_: D9
            d_622_v40_ = D9_DC28(D9_DC27(d_621_v39_, (self.f7)[default__.safeIndex(183, (self.f7).length(0))]))
            d_622_v40_ = d_622_v40_
            (self).f8 = self.f8
            (globalState).f0 = (560 if (self.f7)[default__.safeIndex(183, (self.f7).length(0))] else 789)
            d_623_v41_: _dafny.MultiSet
            d_623_v41_ = _dafny.MultiSet([d_609_v30_, (self).fm6(globalState), (self).fm6(globalState)])
            d_624_v42_: _dafny.Set
            d_624_v42_ = _dafny.Set({d_609_v30_})
            d_625_v43_: _dafny.Seq
            d_625_v43_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gtark"))
            d_626_v44_: _dafny.Map
            d_626_v44_ = _dafny.Map({d_625_v43_: (self.f7)[default__.safeIndex(183, (self.f7).length(0))]})
            d_627_v45_: _dafny.Seq
            d_627_v45_ = _dafny.SeqWithoutIsStrInference([d_609_v30_, (self).fm4(d_609_v30_, d_609_v30_, d_626_v44_, d_609_v30_, globalState)])
            d_628_v46_: _dafny.Map
            d_628_v46_ = _dafny.Map({(0) - (d_609_v30_): self.f8})
            d_629_v47_: _dafny.Map
            d_629_v47_ = _dafny.Map({len(d_628_v46_): d_609_v30_})
            d_630_v48_: _dafny.Array
            nw104_ = _dafny.Array(None, 28)
            nw104_[int(0)] = d_609_v30_
            nw104_[int(1)] = d_609_v30_
            nw104_[int(2)] = d_609_v30_
            nw104_[int(3)] = d_609_v30_
            nw104_[int(4)] = d_609_v30_
            nw104_[int(5)] = (d_623_v41_).cardinality
            nw104_[int(6)] = d_609_v30_
            nw104_[int(7)] = 433
            nw104_[int(8)] = d_609_v30_
            nw104_[int(9)] = (default__.fm15((self).f6, globalState)).cardinality
            nw104_[int(10)] = d_609_v30_
            nw104_[int(11)] = len(d_624_v42_)
            nw104_[int(12)] = d_609_v30_
            nw104_[int(13)] = len(d_627_v45_)
            nw104_[int(14)] = d_609_v30_
            nw104_[int(15)] = d_609_v30_
            nw104_[int(16)] = d_609_v30_
            nw104_[int(17)] = d_609_v30_
            nw104_[int(18)] = d_609_v30_
            nw104_[int(19)] = d_609_v30_
            nw104_[int(20)] = d_609_v30_
            nw104_[int(21)] = (((self).f5).set(False, default__.abs(len(d_628_v46_)))).cardinality
            nw104_[int(22)] = d_609_v30_
            nw104_[int(23)] = len(d_629_v47_)
            nw104_[int(24)] = d_609_v30_
            nw104_[int(25)] = d_609_v30_
            nw104_[int(26)] = d_609_v30_
            nw104_[int(27)] = d_609_v30_
            d_630_v48_ = nw104_
            d_631_v49_: _dafny.Seq
            d_631_v49_ = _dafny.SeqWithoutIsStrInference([d_630_v48_])
            d_631_v49_ = _dafny.SeqWithoutIsStrInference([d_630_v48_, d_630_v48_, d_630_v48_])
        elif True:
            d_632_v50_: _dafny.Map
            d_632_v50_ = _dafny.Map({(d_609_v30_) > (d_609_v30_): (self.f7)[default__.safeIndex(183, (self.f7).length(0))]})
            d_632_v50_ = d_632_v50_
            d_633_v51_: _dafny.Seq
            d_633_v51_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ccqgh"))
            d_634_v52_: _dafny.Set
            d_634_v52_ = _dafny.Set({d_633_v51_})
            arr15_ = self.f7
            index119_ = default__.safeIndex(183, (self.f7).length(0))
            def iife82_():
                coll34_ = _dafny.Set()
                compr_34_: _dafny.Seq
                for compr_34_ in (d_607_v29_).Elements:
                    d_635_v53_: _dafny.Seq = compr_34_
                    if (d_635_v53_) in (d_607_v29_):
                        coll34_ = coll34_.union(_dafny.Set([d_635_v53_]))
                return _dafny.Set(coll34_)
            arr15_[index119_] = (iife82_()
            ).issubset(d_634_v52_)
            (globalState).f0 = d_609_v30_
            d_636_v54_: _dafny.Map
            d_636_v54_ = _dafny.Map({not(not(not((self).fm3(globalState)))): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pa"))})
            d_636_v54_ = (d_636_v54_).set((self.f7)[default__.safeIndex(183, (self.f7).length(0))], _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qc")))
            d_609_v30_ = default__.safeModuloInt(d_609_v30_, (d_609_v30_) * (167))
        d_637_v55_: _dafny.Seq
        d_637_v55_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "djxkfjewt"))
        d_638_v56_: _dafny.Seq
        d_638_v56_ = _dafny.SeqWithoutIsStrInference([d_609_v30_, len(d_637_v55_)])
        d_639_v57_: _dafny.Map
        d_639_v57_ = _dafny.Map({d_638_v56_: d_609_v30_})
        d_639_v57_ = (d_639_v57_).set(_dafny.SeqWithoutIsStrInference([d_609_v30_, d_609_v30_, len(d_637_v55_)]), (d_609_v30_) * ((0) - (len(_dafny.SeqWithoutIsStrInference([(self.f7)[default__.safeIndex(183, (self.f7).length(0))]])))))
        d_640_v58_: _dafny.Set
        d_640_v58_ = _dafny.Set({d_609_v30_, d_609_v30_, d_609_v30_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_641_i8_ in range(default__.abs(537))]))})
        d_642_v59_: _dafny.Map
        d_642_v59_ = _dafny.Map({not(self.f8): self.f8})
        d_643_v60_: _dafny.Map
        d_643_v60_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kdtsapr")): (self.f7)[default__.safeIndex(183, (self.f7).length(0))]})
        r0 = (d_640_v58_).intersection((_dafny.Set({d_609_v30_})).intersection(default__.fm24((self).fm4(d_609_v30_, len(d_642_v59_), d_643_v60_, d_609_v30_, globalState), (self).f6, 111, globalState)))
        r1 = (d_609_v30_) - (d_609_v30_)
        d_644_v61_: str
        d_644_v61_ = _dafny.CodePoint('r')
        d_645_v62_: _dafny.Map
        d_645_v62_ = _dafny.Map({d_644_v61_: self.f8})
        r2 = default__.safeModuloInt(d_609_v30_, len((d_645_v62_).set(d_644_v61_, (self).fm2(self.f8, globalState))))
        return r0, r1, r2

    def m1(self, p0, p1, p2, globalState):
        (self).f8 = True
        d_646_v0_: str
        d_646_v0_ = _dafny.CodePoint('t')
        d_647_v1_: D7
        d_647_v1_ = D7_DC21(p0, p0)
        d_648_v2_: _dafny.Map
        d_648_v2_ = _dafny.Map({(d_647_v1_).cf30: p0})
        d_649_v3_: D2
        d_649_v3_ = D2_DC7(d_646_v0_, d_648_v2_, p0)
        d_650_i0_: int
        d_650_i0_ = 0
        with _dafny.label("8"):
            pat_let_tv16_ = p0
            pat_let_tv17_ = p0
            pat_let_tv18_ = p0
            def lambda31_(source5_):
                if source5_.is_DC5:
                    d_656___mcc_h0_ = source5_.cf5
                    d_657_cf5_ = d_656___mcc_h0_
                    return d_657_cf5_
                elif source5_.is_DC6:
                    return (self).f6
                elif source5_.is_DC7:
                    d_658___mcc_h1_ = source5_.cf6
                    d_659___mcc_h2_ = source5_.cf7
                    d_660___mcc_h3_ = source5_.cf8
                    d_661_cf8_ = d_660___mcc_h3_
                    d_662_cf7_ = d_659___mcc_h2_
                    d_663_cf6_ = d_658___mcc_h1_
                    return (_dafny.Map({pat_let_tv16_: (self).f6})) == (_dafny.Map({d_661_cf8_: (self).f6}))
                elif source5_.is_DC4:
                    d_664___mcc_h4_ = source5_.cf4
                    d_665_cf4_ = d_664___mcc_h4_
                    return (pat_let_tv17_) <= (pat_let_tv18_)
                elif True:
                    d_666___mcc_h5_ = source5_.cf9
                    d_667_cf9_ = d_666___mcc_h5_
                    return self.f8

            while lambda31_(D2_DC8(d_649_v3_)):
                with _dafny.c_label("8"):
                    if (d_650_i0_) >= (100):
                        raise _dafny.Break("8")
                    d_650_i0_ = (d_650_i0_) + (1)
                    (self).f7 = self.f7
                    d_651_v4_: _dafny.Array
                    def lambda30_(d_652_i1_):
                        return D10_DC29(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([self.f8, (self).f6]), _dafny.SeqWithoutIsStrInference([(self).f6])]))

                    init18_ = lambda30_
                    nw105_ = _dafny.Array(None, 27)
                    for i0_18_ in range(nw105_.length(0)):
                        nw105_[i0_18_] = init18_(i0_18_)
                    d_651_v4_ = nw105_
                    d_651_v4_ = d_651_v4_
                    d_653_v5_: _dafny.Seq
                    d_653_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mvhgcimfk"))
                    d_654_v6_: D5
                    d_654_v6_ = D5_DC15(d_653_v5_)
                    d_655_v7_: D3
                    d_655_v7_ = D3_DC10((self).f6, len(_dafny.Map({d_654_v6_: True})))
                    (self).f8 = (d_655_v7_).cf11
                    (globalState).f0 = p0
                    pass
            pass
        d_668_v8_: _dafny.MultiSet
        d_668_v8_ = _dafny.MultiSet([p0])
        d_669_v9_: T0
        nw106_ = C2()
        nw106_.ctor__((self).f5)
        d_669_v9_ = nw106_
        rhs129_ = d_668_v8_
        rhs130_ = p0
        rhs131_ = (p2) == (p2)
        rhs132_ = p1
        rhs133_ = ((d_669_v9_ if (self).f6 else d_669_v9_) if p1 else d_669_v9_)
        lhs98_ = globalState
        lhs99_ = self
        lhs100_ = self
        d_668_v8_ = rhs129_
        lhs98_.f0 = rhs130_
        lhs99_.f8 = rhs131_
        lhs100_.f8 = rhs132_
        d_669_v9_ = rhs133_
        d_670_v10_: _dafny.Seq
        d_670_v10_ = _dafny.SeqWithoutIsStrInference([p0, p0, p0, p0, p0])
        d_671_v11_: _dafny.Seq
        d_671_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lpaexls"))
        d_672_v12_: _dafny.Map
        d_672_v12_ = _dafny.Map({(d_670_v10_) + (d_670_v10_): (d_646_v0_) in (d_671_v11_)})
        if ((d_672_v12_)[d_670_v10_] if (d_670_v10_) in (d_672_v12_) else self.f8):
            (globalState).f0 = p0
            d_673_v13_: D2
            d_673_v13_ = D2_DC8(d_649_v3_)
            source6_ = d_673_v13_
            if source6_.is_DC5:
                d_674___mcc_h6_ = source6_.cf5
                d_675_cf5_ = d_674___mcc_h6_
                d_646_v0_ = d_646_v0_
                d_676_v14_: _dafny.Array
                nw107_ = _dafny.Array(int(0), 17)
                d_676_v14_ = nw107_
                d_676_v14_ = d_676_v14_
                d_675_cf5_ = p2
                d_677_v15_: D7
                d_677_v15_ = D7_DC20(d_675_cf5_)
                (self).f8 = (d_677_v15_).cf29
            elif source6_.is_DC6:
                (self).f8 = p1
                d_678_v17_: D0
                d_678_v17_ = D0_DC1()
                d_679_v18_: D1
                d_679_v18_ = D1_DC3(d_678_v17_, p0)
                def iife83_():
                    coll35_ = _dafny.Map()
                    compr_35_: int
                    for compr_35_ in _dafny.IntegerRange(687, 872):
                        d_680_v16_: int = compr_35_
                        if ((687) <= (d_680_v16_)) and ((d_680_v16_) < (872)):
                            coll35_[default__.safeDivisionInt(d_680_v16_, p0)] = (self).f6
                    return _dafny.Map(coll35_)
                rhs134_ = not(((p0) + (p0)) not in (iife83_()
                ))
                rhs135_ = (d_679_v18_).cf3
                rhs136_ = self.f8
                lhs101_ = self
                lhs102_ = globalState
                lhs103_ = self
                lhs101_.f8 = rhs134_
                lhs102_.f0 = rhs135_
                lhs103_.f8 = rhs136_
                d_681_v19_: _dafny.Map
                d_681_v19_ = _dafny.Map({False: self.f8})
                d_681_v19_ = (d_681_v19_).set((self).f6, (self).f6)
                (globalState).f0 = (self).fm6(globalState)
            elif source6_.is_DC7:
                d_682___mcc_h7_ = source6_.cf6
                d_683___mcc_h8_ = source6_.cf7
                d_684___mcc_h9_ = source6_.cf8
                d_685_cf8_ = d_684___mcc_h9_
                d_686_cf7_ = d_683___mcc_h8_
                d_687_cf6_ = d_682___mcc_h7_
                (globalState).f0 = d_685_cf8_
                d_688_v20_: _dafny.Seq
                d_688_v20_ = _dafny.SeqWithoutIsStrInference([True, p1])
                d_689_v21_: _dafny.Map
                d_689_v21_ = _dafny.Map({D1_DC2(not((self).f6)): d_688_v20_})
                d_690_v22_: D1
                d_690_v22_ = D1_DC2((self).f6)
                d_689_v21_ = (d_689_v21_).set(d_690_v22_, default__.fm23(p0, globalState))
                arr16_ = self.f7
                index120_ = default__.safeIndex(682, (self.f7).length(0))
                arr16_[index120_] = p1
                d_691_v23_: _dafny.MultiSet
                d_691_v23_ = _dafny.MultiSet([d_670_v10_])
                d_692_v24_: D0
                d_692_v24_ = D0_DC1()
                d_693_v25_: D1
                d_693_v25_ = D1_DC3(d_692_v24_, p0)
                d_694_v26_: _dafny.Map
                d_694_v26_ = _dafny.Map({p2: self.f8})
                arr17_ = self.f7
                index121_ = default__.safeIndex(682, (self.f7).length(0))
                rhs137_ = (d_669_v9_).fm1((d_691_v23_) == (d_691_v23_), (p0) - (d_685_cf8_), (d_693_v25_).cf3, globalState)
                rhs138_ = (d_670_v10_) + (d_670_v10_)
                rhs139_ = (self).fm2((d_694_v26_) == (d_694_v26_), globalState)
                rhs140_ = d_670_v10_
                lhs104_ = self.f7
                lhs105_ = default__.safeIndex(682, (self.f7).length(0))
                lhs106_ = self
                lhs104_[lhs105_] = rhs137_
                d_670_v10_ = rhs138_
                lhs106_.f8 = rhs139_
                d_670_v10_ = rhs140_
                (globalState).f0 = (0) - (((self).fm6(globalState)) + (d_685_cf8_))
            elif source6_.is_DC4:
                d_695___mcc_h10_ = source6_.cf4
                d_696_cf4_ = d_695___mcc_h10_
                (self).f8 = (d_646_v0_) in ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vurhyya"))) + (_dafny.SeqWithoutIsStrInference([d_646_v0_ for d_697_i2_ in range(default__.abs(844))])))
                d_698_v27_: _dafny.Array
                nw108_ = _dafny.Array(None, 12)
                d_698_v27_ = nw108_
                (self).f8 = ((d_671_v11_)[default__.safeIndex((_dafny.MultiSet([d_698_v27_])).cardinality, len(d_671_v11_))]) in (d_671_v11_)
                (self).f8 = not(p2)
                d_699_v28_: _dafny.Map
                d_699_v28_ = _dafny.Map({(p1) and (self.f8): (self).fm6(globalState)})
                d_700_v29_: _dafny.Seq
                d_700_v29_ = _dafny.SeqWithoutIsStrInference([(self).f6])
                d_699_v28_ = (d_699_v28_).set((self).f6, (len((d_700_v29_).set(default__.safeIndex(p0, len(d_700_v29_)), (self).f6))) + (p0))
            elif True:
                d_701___mcc_h11_ = source6_.cf9
                d_702_cf9_ = d_701___mcc_h11_
                d_703_v30_: _dafny.Seq
                d_703_v30_ = _dafny.SeqWithoutIsStrInference([self.f8])
                d_703_v30_ = ((d_703_v30_) + (d_703_v30_)) + (_dafny.SeqWithoutIsStrInference([False]))
                (self).f8 = ((self).fm7(globalState)) <= (d_671_v11_)
                d_704_v31_: _dafny.Array
                nw109_ = _dafny.Array(int(0), 11)
                d_704_v31_ = nw109_
                d_705_v32_: _dafny.MultiSet
                d_705_v32_ = _dafny.MultiSet([d_704_v31_])
                (self).f8 = (d_705_v32_) == (d_705_v32_)
                (self).f7 = self.f7
            d_706_v33_: _dafny.Map
            d_706_v33_ = _dafny.Map({(self).f6: self.f8})
            d_706_v33_ = d_706_v33_
            if p2:
                def lambda32_(d_707_i3_):
                    return True

                init19_ = lambda32_
                nw110_ = _dafny.Array(None, 26)
                for i0_19_ in range(nw110_.length(0)):
                    nw110_[i0_19_] = init19_(i0_19_)
                (self).f7 = nw110_
                d_708_v34_: D0
                d_708_v34_ = D0_DC0(p0)
                d_709_v35_: _dafny.Seq
                d_709_v35_ = _dafny.SeqWithoutIsStrInference([D0_DC0(p0), d_708_v34_, d_708_v34_, (d_708_v34_ if p2 else d_708_v34_), default__.fm25((self).f6, p0, p2, globalState)])
                (globalState).f0 = len(d_709_v35_)
                d_710_v36_: _dafny.Map
                d_710_v36_ = _dafny.Map({self.f7: p2})
                d_711_v37_: D9
                d_711_v37_ = D9_DC27(d_710_v36_, (self.f8 if p2 else p1))
                d_711_v37_ = d_711_v37_
                (self).f8 = (self).f6
                d_712_v38_: _dafny.Array
                nw111_ = _dafny.Array(_dafny.Map({}), 9)
                d_712_v38_ = nw111_
                d_713_v40_: _dafny.Seq
                d_713_v40_ = _dafny.SeqWithoutIsStrInference([(self).f5])
                d_714_v41_: _dafny.Map
                d_714_v41_ = _dafny.Map({(self).f5: (self).f6})
                index122_ = default__.safeIndex(206, (d_712_v38_).length(0))
                def iife84_():
                    coll36_ = _dafny.Map()
                    compr_36_: _dafny.MultiSet
                    for compr_36_ in (d_713_v40_).Elements:
                        d_715_v39_: _dafny.MultiSet = compr_36_
                        if (d_715_v39_) in (d_713_v40_):
                            coll36_[d_715_v39_] = self.f8
                    return _dafny.Map(coll36_)
                (d_712_v38_)[index122_] = (iife84_()
                ) | (d_714_v41_)
                index123_ = default__.safeIndex(206, (d_712_v38_).length(0))
                (d_712_v38_)[index123_] = (d_714_v41_) | (d_714_v41_)
            elif True:
                d_646_v0_ = d_646_v0_
                d_716_v42_: C2
                nw112_ = C2()
                nw112_.ctor__((d_669_v9_).f5)
                d_716_v42_ = nw112_
                d_717_v43_: D2
                d_717_v43_ = D2_DC7(d_646_v0_, d_648_v2_, p0)
                (globalState).f0 = (d_717_v43_).cf8
                (self).f7 = self.f7
                (globalState).f0 = (0) - (p0)
            d_718_v44_: _dafny.Map
            d_718_v44_ = _dafny.Map({p0: (self).f6})
            d_719_v45_: C1
            nw113_ = C1()
            nw113_.ctor__(p2, (d_671_v11_) + (d_671_v11_), ((d_718_v44_)[p0] if (p0) in (d_718_v44_) else False), (self).f5)
            d_719_v45_ = nw113_
        elif True:
            d_720_v46_: _dafny.Array
            nw114_ = _dafny.Array(_dafny.Map({}), 28)
            d_720_v46_ = nw114_
            d_721_v47_: _dafny.Map
            d_721_v47_ = _dafny.Map({p2: (self).f6})
            index124_ = default__.safeIndex(378, (d_720_v46_).length(0))
            (d_720_v46_)[index124_] = (d_721_v47_) | (d_721_v47_)
            d_722_v48_: D11
            d_722_v48_ = D11_DC32(d_721_v47_)
            index125_ = default__.safeIndex(378, (d_720_v46_).length(0))
            (d_720_v46_)[index125_] = (d_721_v47_) | (((d_722_v48_).cf51) | (d_721_v47_))
            d_723_v49_: D1
            d_723_v49_ = D1_DC2(p1)
            source7_ = d_723_v49_
            if source7_.is_DC3:
                d_724___mcc_h12_ = source7_.cf2
                d_725___mcc_h13_ = source7_.cf3
                d_726_cf3_ = d_725___mcc_h13_
                d_727_cf2_ = d_724___mcc_h12_
                d_671_v11_ = d_671_v11_
                (globalState).f0 = (d_726_cf3_) - (p0)
                d_728_v50_: _dafny.Seq
                d_728_v50_ = _dafny.SeqWithoutIsStrInference([p2])
                d_729_v51_: C1
                nw115_ = C1()
                nw115_.ctor__(p2, d_671_v11_, (default__.fm23(d_726_cf3_, globalState)) < (d_728_v50_), (d_669_v9_).f5)
                d_729_v51_ = nw115_
                d_730_v52_: D3
                d_730_v52_ = D3_DC11(58)
                d_731_v53_: _dafny.Map
                d_731_v53_ = _dafny.Map({p2: d_730_v52_})
                d_726_cf3_ = (d_726_cf3_) - (default__.safeDivisionInt((0) - (p0), len(d_731_v53_)))
            elif True:
                d_732___mcc_h14_ = source7_.cf1
                d_733_cf1_ = d_732___mcc_h14_
                d_734_v54_: _dafny.Set
                d_734_v54_ = _dafny.Set({p2})
                d_734_v54_ = d_734_v54_
                d_735_v55_: _dafny.Array
                nw116_ = _dafny.Array(D2.default()(), 28)
                d_735_v55_ = nw116_
                d_735_v55_ = d_735_v55_
                (globalState).f0 = (default__.safeDivisionInt(p0, p0)) + (p0)
                d_736_v56_: _dafny.Array
                nw117_ = _dafny.Array(_dafny.Map({}), 9)
                d_736_v56_ = nw117_
                d_737_v57_: _dafny.Map
                d_737_v57_ = _dafny.Map({d_646_v0_: False})
                index126_ = default__.safeIndex(290, (d_736_v56_).length(0))
                (d_736_v56_)[index126_] = d_737_v57_
                d_738_v58_: D12
                d_738_v58_ = D12_DC34(d_737_v57_)
                index127_ = default__.safeIndex(290, (d_736_v56_).length(0))
                (d_736_v56_)[index127_] = ((d_738_v58_).cf54) | (d_737_v57_)
            d_739_v59_: _dafny.Map
            d_739_v59_ = _dafny.Map({self.f7: _dafny.MultiSet([p2])})
            d_740_v60_: _dafny.Array
            def lambda33_(d_741_i4_):
                return default__.safeModuloInt(d_741_i4_, -705)

            init20_ = lambda33_
            nw118_ = _dafny.Array(None, 7)
            for i0_20_ in range(nw118_.length(0)):
                nw118_[i0_20_] = init20_(i0_20_)
            d_740_v60_ = nw118_
            index128_ = default__.safeIndex(48, (d_740_v60_).length(0))
            (d_740_v60_)[index128_] = p0
            d_742_v61_: D13
            d_742_v61_ = D13_DC38(d_739_v59_)
            d_743_v62_: _dafny.Set
            d_743_v62_ = _dafny.Set({(self).f6, p2})
            d_744_v63_: _dafny.Seq
            d_744_v63_ = _dafny.SeqWithoutIsStrInference([self.f8])
            d_745_v64_: _dafny.Set
            d_745_v64_ = _dafny.Set({len(d_744_v63_)})
            index129_ = default__.safeIndex(48, (d_740_v60_).length(0))
            rhs141_ = (d_742_v61_).cf59
            rhs142_ = len((d_743_v62_).intersection((d_743_v62_).intersection(_dafny.Set({(d_744_v63_)[default__.safeIndex(p0, len(d_744_v63_))]}))))
            rhs143_ = not (p1) or (p2)
            rhs144_ = default__.safeDivisionInt(p0, (p0) + (len(d_745_v64_)))
            lhs107_ = d_740_v60_
            lhs108_ = default__.safeIndex(48, (d_740_v60_).length(0))
            lhs109_ = self
            lhs110_ = globalState
            d_739_v59_ = rhs141_
            lhs107_[lhs108_] = rhs142_
            lhs109_.f8 = rhs143_
            lhs110_.f0 = rhs144_
            (globalState).f0 = ((d_740_v60_)[default__.safeIndex(48, (d_740_v60_).length(0))]) + (p0)
            d_746_v65_: C2
            nw119_ = C2()
            nw119_.ctor__((d_669_v9_).f5)
            d_746_v65_ = nw119_
        if self.f8:
            (self).f8 = p2
            d_747_v66_: D5
            d_747_v66_ = D5_DC15(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uipsyskp")))
            pat_let_tv19_ = d_671_v11_
            d_748_v67_: _dafny.Map
            def iife85_(_pat_let24_0):
                def iife86_(d_749_dt__update__tmp_h0_):
                    def iife87_(_pat_let25_0):
                        def iife88_(d_750_dt__update_hcf18_h0_):
                            return D5_DC15(d_750_dt__update_hcf18_h0_)
                        return iife88_(_pat_let25_0)
                    return iife87_(pat_let_tv19_)
                return iife86_(_pat_let24_0)
            d_748_v67_ = _dafny.Map({d_670_v10_: iife85_(d_747_v66_)})
            d_748_v67_ = d_748_v67_
            d_751_v68_: _dafny.Seq
            d_751_v68_ = _dafny.SeqWithoutIsStrInference([not(p2)])
            d_648_v2_ = (d_648_v2_).set(p0, ((d_668_v8_)[p0] if (p0) in (d_668_v8_) else len(d_751_v68_)))
            d_752_v69_: D12
            d_752_v69_ = D12_DC36(((((self).f5)[self.f8] if (self.f8) in ((self).f5) else p0) if self.f8 else p0), p0)
            d_752_v69_ = d_752_v69_
            (self).f8 = True
        elif True:
            d_753_v70_: _dafny.Map
            d_753_v70_ = _dafny.Map({p1: self.f7})
            d_753_v70_ = (d_753_v70_) | (d_753_v70_)
            d_646_v0_ = d_646_v0_
            d_754_v71_: D3
            d_754_v71_ = D3_DC11(p0)
            d_755_v72_: _dafny.Map
            d_755_v72_ = _dafny.Map({(0) - ((d_754_v71_).cf13): d_671_v11_})
            d_756_v73_: _dafny.Seq
            d_756_v73_ = _dafny.SeqWithoutIsStrInference([p2])
            d_755_v72_ = (d_755_v72_).set(default__.safeModuloInt(len(d_756_v73_), p0), d_671_v11_)
            d_757_v74_: _dafny.Array
            def lambda34_(d_758_p0_):
                def lambda35_(d_759_i5_):
                    return (d_759_i5_) * (d_758_p0_)

                return lambda35_

            init21_ = lambda34_(p0)
            nw120_ = _dafny.Array(None, 8)
            for i0_21_ in range(nw120_.length(0)):
                nw120_[i0_21_] = init21_(i0_21_)
            d_757_v74_ = nw120_
            index130_ = default__.safeIndex(558, (d_757_v74_).length(0))
            (d_757_v74_)[index130_] = p0
            d_760_v75_: _dafny.Seq
            d_760_v75_ = _dafny.SeqWithoutIsStrInference([d_756_v73_, d_756_v73_])
            index131_ = default__.safeIndex(558, (d_757_v74_).length(0))
            rhs145_ = (d_670_v10_)[default__.safeIndex(len((d_760_v75_) + (d_760_v75_)), len(d_670_v10_))]
            rhs146_ = (default__.safeDivisionInt(p0, p0)) - (len(d_756_v73_))
            rhs147_ = p1
            rhs148_ = d_671_v11_
            rhs149_ = ((d_671_v11_).set(default__.safeIndex(p0, len(d_671_v11_)), d_646_v0_)) + ((d_671_v11_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uayxfpbbt"))))
            lhs111_ = d_757_v74_
            lhs112_ = default__.safeIndex(558, (d_757_v74_).length(0))
            lhs113_ = globalState
            lhs114_ = self
            lhs111_[lhs112_] = rhs145_
            lhs113_.f0 = rhs146_
            lhs114_.f8 = rhs147_
            d_671_v11_ = rhs148_
            d_671_v11_ = rhs149_
            d_761_v76_: C3
            nw121_ = C3()
            nw121_.ctor__(self.f8, (d_669_v9_).f5)
            d_761_v76_ = nw121_
        d_762_v77_: D12
        d_762_v77_ = D12_DC36(p0, p0)
        source8_ = d_762_v77_
        if source8_.is_DC35:
            d_763___mcc_h15_ = source8_.cf55
            d_764_cf55_ = d_763___mcc_h15_
            d_646_v0_ = d_646_v0_
            d_765_v78_: _dafny.Array
            nw122_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 15)
            d_765_v78_ = nw122_
            index132_ = default__.safeIndex(443, (d_765_v78_).length(0))
            (d_765_v78_)[index132_] = d_671_v11_
            d_766_v79_: _dafny.Array
            nw123_ = _dafny.Array(int(0), 12)
            d_766_v79_ = nw123_
            d_767_v80_: _dafny.Seq
            d_767_v80_ = _dafny.SeqWithoutIsStrInference([d_766_v79_])
            index133_ = default__.safeIndex(443, (d_765_v78_).length(0))
            (d_765_v78_)[index133_] = (d_671_v11_).set(default__.safeIndex(default__.safeDivisionInt(len(d_767_v80_), p0), len(d_671_v11_)), d_646_v0_)
            d_768_v81_: _dafny.Map
            d_768_v81_ = _dafny.Map({547: p2})
            d_769_v82_: _dafny.Map
            d_769_v82_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference([len(d_768_v81_), len(d_670_v10_)])) + (d_670_v10_): p0})
            index134_ = default__.safeIndex(386, (d_766_v79_).length(0))
            (d_766_v79_)[index134_] = (p0) - (p0)
            arr18_ = self.f7
            index135_ = default__.safeIndex(748, (self.f7).length(0))
            arr18_[index135_] = not (d_764_cf55_) or (not(p1))
            d_770_v83_: _dafny.Seq
            d_770_v83_ = _dafny.SeqWithoutIsStrInference([p1])
            d_771_v84_: _dafny.Set
            d_771_v84_ = _dafny.Set({p0, p0})
            d_772_v86_: _dafny.Seq
            d_772_v86_ = _dafny.SeqWithoutIsStrInference([d_648_v2_, d_648_v2_, d_648_v2_, d_648_v2_, d_648_v2_])
            d_773_v87_: C0
            nw124_ = C0()
            def iife89_():
                coll37_ = _dafny.Map()
                compr_37_: _dafny.Map
                for compr_37_ in (d_772_v86_).Elements:
                    d_774_v85_: _dafny.Map = compr_37_
                    if (d_774_v85_) in (d_772_v86_):
                        coll37_[d_774_v85_] = self.f8
                return _dafny.Map(coll37_)
            nw124_.ctor__(iife89_()
            )
            d_773_v87_ = nw124_
            d_775_v88_: _dafny.Map
            d_775_v88_ = _dafny.Map({(0) - (((d_669_v9_).f5).cardinality): d_773_v87_})
            index136_ = default__.safeIndex(386, (d_766_v79_).length(0))
            arr19_ = self.f7
            index137_ = default__.safeIndex(748, (self.f7).length(0))
            rhs150_ = _dafny.Map({d_670_v10_: len(d_770_v83_)})
            rhs151_ = (d_669_v9_).fm1(p2, p0, p0, globalState)
            rhs152_ = (p0) * (((d_670_v10_)[default__.safeIndex(len(d_771_v84_), len(d_670_v10_))]) - (len(d_775_v88_)))
            rhs153_ = not((self).f6)
            rhs154_ = (p2) and ((d_671_v11_) <= ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p"))).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p")))), default__.fm16(p2, d_648_v2_, (d_765_v78_)[default__.safeIndex(443, (d_765_v78_).length(0))], globalState))))
            lhs115_ = self
            lhs116_ = d_766_v79_
            lhs117_ = default__.safeIndex(386, (d_766_v79_).length(0))
            lhs118_ = self.f7
            lhs119_ = default__.safeIndex(748, (self.f7).length(0))
            d_769_v82_ = rhs150_
            lhs115_.f8 = rhs151_
            lhs116_[lhs117_] = rhs152_
            lhs118_[lhs119_] = rhs153_
            d_764_cf55_ = rhs154_
            if p1:
                index138_ = default__.safeIndex(443, (d_765_v78_).length(0))
                (d_765_v78_)[index138_] = d_671_v11_
                (self).f8 = self.f8
                d_776_v89_: C1
                nw125_ = C1()
                nw125_.ctor__(not (not(p2)) or (self.f8), d_671_v11_, p1, (d_669_v9_).f5)
                d_776_v89_ = nw125_
                d_777_v90_: _dafny.MultiSet
                d_777_v90_ = _dafny.MultiSet([default__.fm18(globalState)])
                d_778_v91_: _dafny.Array
                nw126_ = _dafny.Array(_dafny.MultiSet({}), 14)
                d_778_v91_ = nw126_
                d_779_v92_: _dafny.Map
                d_779_v92_ = _dafny.Map({(d_778_v91_ if d_764_cf55_ else d_778_v91_): p0})
                arr20_ = self.f7
                index139_ = default__.safeIndex(748, (self.f7).length(0))
                rhs155_ = (default__.safeModuloInt(p0, (0) - (p0))) * ((d_766_v79_)[default__.safeIndex(386, (d_766_v79_).length(0))])
                rhs156_ = (((d_777_v90_)[_dafny.SeqWithoutIsStrInference([p0 for d_780_i6_ in range(default__.abs(810))])] if (_dafny.SeqWithoutIsStrInference([p0 for d_781_i6_ in range(default__.abs(810))])) in (d_777_v90_) else p0)) not in (d_771_v84_)
                rhs157_ = ((d_779_v92_)[d_778_v91_] if (d_778_v91_) in (d_779_v92_) else ((d_766_v79_)[default__.safeIndex(386, (d_766_v79_).length(0))]) - ((d_766_v79_)[default__.safeIndex(386, (d_766_v79_).length(0))]))
                lhs120_ = globalState
                lhs121_ = self.f7
                lhs122_ = default__.safeIndex(748, (self.f7).length(0))
                lhs123_ = globalState
                lhs120_.f0 = rhs155_
                lhs121_[lhs122_] = rhs156_
                lhs123_.f0 = rhs157_
                d_782_v93_: _dafny.Map
                d_782_v93_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yc")): (d_776_v89_).f12})
                d_783_v94_: _dafny.Seq
                d_783_v94_ = _dafny.SeqWithoutIsStrInference([(d_776_v89_).f13, (d_765_v78_)[default__.safeIndex(443, (d_765_v78_).length(0))], (d_765_v78_)[default__.safeIndex(443, (d_765_v78_).length(0))]])
                d_784_v95_: _dafny.Map
                d_784_v95_ = _dafny.Map({(d_783_v94_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([(d_766_v79_)[default__.safeIndex(386, (d_766_v79_).length(0))] for d_785_i7_ in range(default__.abs(280))])), len(d_783_v94_))]: d_646_v0_})
                (globalState).f0 = (d_776_v89_).fm4(default__.safeDivisionInt((d_766_v79_)[default__.safeIndex(386, (d_766_v79_).length(0))], (d_766_v79_)[default__.safeIndex(386, (d_766_v79_).length(0))]), (d_766_v79_)[default__.safeIndex(386, (d_766_v79_).length(0))], d_782_v93_, len(d_784_v95_), globalState)
            elif True:
                d_786_v96_: C0
                nw127_ = C0()
                nw127_.ctor__(_dafny.Map({d_648_v2_: (self.f7)[default__.safeIndex(748, (self.f7).length(0))]}))
                d_786_v96_ = nw127_
                d_787_v97_: _dafny.Map
                d_787_v97_ = _dafny.Map({d_764_cf55_: True})
                arr21_ = self.f7
                index140_ = default__.safeIndex(748, (self.f7).length(0))
                rhs158_ = p1
                rhs159_ = (len((d_787_v97_) | (d_787_v97_))) * ((0) - (p0))
                rhs160_ = d_787_v97_
                lhs124_ = self.f7
                lhs125_ = default__.safeIndex(748, (self.f7).length(0))
                lhs126_ = globalState
                lhs124_[lhs125_] = rhs158_
                lhs126_.f0 = rhs159_
                d_787_v97_ = rhs160_
                d_788_v98_: _dafny.Map
                d_788_v98_ = _dafny.Map({d_646_v0_: d_764_cf55_})
                d_789_v99_: _dafny.MultiSet
                d_789_v99_ = _dafny.MultiSet([d_788_v98_])
                d_790_v100_: C1
                nw128_ = C1()
                nw128_.ctor__((d_789_v99_).issubset(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_788_v98_ for d_791_i8_ in range(default__.abs(564))]))), (d_765_v78_)[default__.safeIndex(443, (d_765_v78_).length(0))], (d_671_v11_) <= ((d_765_v78_)[default__.safeIndex(443, (d_765_v78_).length(0))]), (d_669_v9_).f5)
                d_790_v100_ = nw128_
                d_764_cf55_ = p2
                d_766_v79_ = d_766_v79_
        elif source8_.is_DC36:
            d_792___mcc_h16_ = source8_.cf56
            d_793___mcc_h17_ = source8_.cf57
            d_794_cf57_ = d_793___mcc_h17_
            d_795_cf56_ = d_792___mcc_h16_
            (self).f8 = not ((d_794_cf57_) > (d_794_cf57_)) or ((self.f8 if p1 else p1))
            d_670_v10_ = d_670_v10_
            d_794_cf57_ = ((d_670_v10_)[default__.safeIndex(d_795_cf56_, len(d_670_v10_))]) - ((d_668_v8_).cardinality)
            (globalState).f0 = d_795_cf56_
        elif source8_.is_DC37:
            d_796___mcc_h18_ = source8_.cf58
            d_797_cf58_ = d_796___mcc_h18_
            d_798_v101_: C1
            nw129_ = C1()
            nw129_.ctor__(p2, d_671_v11_, (self).f6, _dafny.MultiSet([p1]))
            d_798_v101_ = nw129_
            d_799_v102_: _dafny.Set
            d_799_v102_ = _dafny.Set({d_798_v101_})
            if (d_799_v102_).ispropersubset(_dafny.Set({d_798_v101_, d_798_v101_})):
                (self).f8 = self.f8
                d_800_v103_: C0
                nw130_ = C0()
                nw130_.ctor__(default__.fm26(p0, globalState))
                d_800_v103_ = nw130_
                (globalState).f0 = ((D6_DC18(d_800_v103_, default__.fm23(len(_dafny.Map({d_800_v103_: 179})), globalState), len(d_671_v11_))).cf27) - (p0)
                d_801_v104_: C1
                nw131_ = C1()
                nw131_.ctor__(((d_669_v9_).f5).isdisjoint((self).f5), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j")), self.f8, (self).f5)
                d_801_v104_ = nw131_
                d_648_v2_ = (d_648_v2_).set(p0, default__.safeModuloInt(((d_669_v9_).f5).cardinality, p0))
                d_802_v105_: C3
                nw132_ = C3()
                nw132_.ctor__(p1, _dafny.MultiSet([(d_798_v101_).f12]))
                d_802_v105_ = nw132_
            elif True:
                (self).f8 = (p0) < ((0) - (p0))
                (self).f8 = (p0) == (default__.safeModuloInt(p0, len((_dafny.Map({509: (d_798_v101_).f12})))))
                d_803_v106_: _dafny.Map
                d_803_v106_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_804_i9_ in range(default__.abs(-876))]): (d_798_v101_).f12})
                d_805_v107_: _dafny.Map
                d_805_v107_ = _dafny.Map({(d_798_v101_).fm4(p0, -167, d_803_v106_, (d_798_v101_).fm4((0) - (p0), p0, _dafny.Map({d_671_v11_: p1}), p0, globalState), globalState): d_797_cf58_})
                d_805_v107_ = (d_805_v107_).set((0) - (p0), default__.fm16(self.f8, d_648_v2_, _dafny.SeqWithoutIsStrInference([d_646_v0_ for d_806_i10_ in range(default__.abs(-403))]), globalState))
                d_807_v108_: _dafny.Seq
                d_807_v108_ = _dafny.SeqWithoutIsStrInference([(d_798_v101_).f12, p1])
                d_808_v109_: _dafny.Map
                d_808_v109_ = _dafny.Map({999: not(self.f8)})
                d_809_v110_: _dafny.Map
                d_809_v110_ = _dafny.Map({not(p1): (d_798_v101_).fm3(globalState)})
                d_810_v111_: D8
                d_810_v111_ = D8_DC24(p0, (d_798_v101_).f12, d_807_v108_, (0) - (len(default__.fm8(p0, ((d_808_v109_)[len(d_809_v110_)] if (len(d_809_v110_)) in (d_808_v109_) else p2), d_671_v11_, globalState))), p0)
                (globalState).f0 = len((d_810_v111_).cf36)
                d_811_v113_: _dafny.Set
                d_811_v113_ = _dafny.Set({((d_798_v101_).f13).set(default__.safeIndex(p0, len((d_798_v101_).f13)), d_646_v0_)})
                d_812_v114_: D15
                d_812_v114_ = D15_DC41(d_803_v106_)
                def iife90_():
                    coll38_ = _dafny.Map()
                    compr_38_: _dafny.Seq
                    for compr_38_ in (d_811_v113_).Elements:
                        d_813_v112_: _dafny.Seq = compr_38_
                        if (d_813_v112_) in (d_811_v113_):
                            coll38_[d_813_v112_] = True
                    return _dafny.Map(coll38_)
                (globalState).f0 = (self).fm4((self).fm4(p0, p0, iife90_()
                , (0) - (p0), globalState), 184, (d_812_v114_).cf66, (p0) - (len(d_805_v107_)), globalState)
            rhs161_ = (self).fm3(globalState)
            rhs162_ = p2
            lhs127_ = self
            lhs128_ = self
            lhs127_.f8 = rhs161_
            lhs128_.f8 = rhs162_
            d_646_v0_ = d_797_cf58_
            d_668_v8_ = d_668_v8_
        elif True:
            d_814___mcc_h19_ = source8_.cf54
            d_815_cf54_ = d_814___mcc_h19_
            d_670_v10_ = d_670_v10_
            d_816_v115_: C2
            nw133_ = C2()
            nw133_.ctor__((self).f5)
            d_816_v115_ = nw133_
            d_817_v116_: D9
            d_817_v116_ = D9_DC26(d_816_v115_)
            source9_ = d_817_v116_
            if source9_.is_DC27:
                d_818___mcc_h20_ = source9_.cf42
                d_819___mcc_h21_ = source9_.cf43
                d_820_cf43_ = d_819___mcc_h21_
                d_821_cf42_ = d_818___mcc_h20_
                d_671_v11_ = d_671_v11_
                d_822_v117_: _dafny.Array
                nw134_ = _dafny.Array(None, 14)
                nw134_[int(0)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sspyxy"))
                nw134_[int(1)] = d_671_v11_
                nw134_[int(2)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "enh"))
                nw134_[int(3)] = _dafny.SeqWithoutIsStrInference([d_646_v0_ for d_823_i11_ in range(default__.abs(128))])
                nw134_[int(4)] = d_671_v11_
                nw134_[int(5)] = d_671_v11_
                nw134_[int(6)] = d_671_v11_
                nw134_[int(7)] = d_671_v11_
                nw134_[int(8)] = d_671_v11_
                nw134_[int(9)] = d_671_v11_
                nw134_[int(10)] = d_671_v11_
                nw134_[int(11)] = d_671_v11_
                nw134_[int(12)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uu"))
                nw134_[int(13)] = _dafny.SeqWithoutIsStrInference([d_646_v0_ for d_824_i12_ in range(default__.abs(651))])
                d_822_v117_ = nw134_
                d_825_v118_: _dafny.Map
                d_825_v118_ = _dafny.Map({d_822_v117_: (self).f5})
                d_825_v118_ = (d_825_v118_).set(d_822_v117_, _dafny.MultiSet([(self).f6]))
                d_826_v119_: _dafny.Seq
                d_826_v119_ = _dafny.SeqWithoutIsStrInference([False])
                d_826_v119_ = d_826_v119_
                (globalState).f0 = (p0) - ((0) - ((((d_669_v9_).f5)[p1] if (p1) in ((d_669_v9_).f5) else p0)))
            elif source9_.is_DC26:
                d_827___mcc_h22_ = source9_.cf41
                d_828_cf41_ = d_827___mcc_h22_
                d_829_v120_: C1
                nw135_ = C1()
                nw135_.ctor__((self).f6, d_671_v11_, p1, (d_669_v9_).f5)
                d_829_v120_ = nw135_
                d_646_v0_ = (d_646_v0_ if (d_829_v120_).f12 else d_646_v0_)
                arr22_ = self.f7
                index141_ = default__.safeIndex(408, (self.f7).length(0))
                arr22_[index141_] = (d_829_v120_).f12
                d_830_v121_: _dafny.Map
                d_830_v121_ = _dafny.Map({(d_829_v120_).f13: p2})
                d_831_v122_: _dafny.Map
                d_831_v122_ = _dafny.Map({634: d_646_v0_})
                d_832_v123_: _dafny.Array
                nw136_ = _dafny.Array(None, 13)
                nw136_[int(0)] = p0
                nw136_[int(1)] = p0
                nw136_[int(2)] = p0
                nw136_[int(3)] = (self).fm6(globalState)
                nw136_[int(4)] = (d_829_v120_).fm4(109, p0, d_830_v121_, p0, globalState)
                nw136_[int(5)] = p0
                nw136_[int(6)] = p0
                nw136_[int(7)] = (len(d_831_v122_)) - ((0) - (p0))
                nw136_[int(8)] = p0
                nw136_[int(9)] = p0
                nw136_[int(10)] = p0
                nw136_[int(11)] = p0
                nw136_[int(12)] = ((d_648_v2_)[len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nqbu")))] if (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nqbu")))) in (d_648_v2_) else len(d_671_v11_))
                d_832_v123_ = nw136_
                index142_ = default__.safeIndex(136, (d_832_v123_).length(0))
                (d_832_v123_)[index142_] = (self).fm4(p0, p0, d_830_v121_, (d_670_v10_)[default__.safeIndex(p0, len(d_670_v10_))], globalState)
                d_833_v124_: _dafny.Set
                d_833_v124_ = _dafny.Set({9})
                arr23_ = self.f7
                index143_ = default__.safeIndex(408, (self.f7).length(0))
                index144_ = default__.safeIndex(136, (d_832_v123_).length(0))
                rhs163_ = p0
                rhs164_ = p2
                rhs165_ = default__.safeModuloInt(len((d_833_v124_) - (d_833_v124_)), ((d_669_v9_).f5).cardinality)
                rhs166_ = (self).f6
                lhs129_ = globalState
                lhs130_ = self.f7
                lhs131_ = default__.safeIndex(408, (self.f7).length(0))
                lhs132_ = d_832_v123_
                lhs133_ = default__.safeIndex(136, (d_832_v123_).length(0))
                lhs134_ = self
                lhs129_.f0 = rhs163_
                lhs130_[lhs131_] = rhs164_
                lhs132_[lhs133_] = rhs165_
                lhs134_.f8 = rhs166_
                d_834_v125_: _dafny.Set
                d_834_v125_ = _dafny.Set({(self).f6})
                d_835_v126_: _dafny.Set
                d_835_v126_ = _dafny.Set({_dafny.Set({p2, True}), d_834_v125_})
                d_836_v127_: _dafny.Map
                d_836_v127_ = _dafny.Map({len(d_835_v126_): (d_829_v120_).f12})
                d_836_v127_ = (d_836_v127_).set(p0, (_dafny.MultiSet([(self).fm2((self.f7)[default__.safeIndex(408, (self.f7).length(0))], globalState), p1])).isdisjoint((d_669_v9_).f5))
            elif True:
                d_837___mcc_h23_ = source9_.cf44
                d_838_cf44_ = d_837___mcc_h23_
                arr24_ = self.f7
                index145_ = default__.safeIndex(7, (self.f7).length(0))
                arr24_[index145_] = True
                d_839_v128_: C3
                nw137_ = C3()
                nw137_.ctor__((self).f6, (self).f5)
                d_839_v128_ = nw137_
                d_840_v129_: _dafny.Map
                d_840_v129_ = _dafny.Map({d_839_v128_: self.f7})
                d_841_v130_: _dafny.Seq
                d_841_v130_ = _dafny.SeqWithoutIsStrInference([self.f7, ((d_840_v129_)[d_839_v128_] if (d_839_v128_) in (d_840_v129_) else self.f7)])
                d_842_v131_: _dafny.Array
                def lambda36_(d_843_v10_):
                    def lambda37_(d_844_i13_):
                        return d_843_v10_

                    return lambda37_

                init22_ = lambda36_(d_670_v10_)
                nw138_ = _dafny.Array(None, 5)
                for i0_22_ in range(nw138_.length(0)):
                    nw138_[i0_22_] = init22_(i0_22_)
                d_842_v131_ = nw138_
                d_845_v132_: _dafny.Array
                nw139_ = _dafny.Array(int(0), 10)
                d_845_v132_ = nw139_
                d_846_v133_: _dafny.Map
                d_846_v133_ = _dafny.Map({d_842_v131_: len(_dafny.Map({p1: d_845_v132_}))})
                arr25_ = self.f7
                index146_ = default__.safeIndex(7, (self.f7).length(0))
                rhs167_ = (self).f6
                rhs168_ = len((d_841_v130_) + (d_841_v130_))
                rhs169_ = d_646_v0_
                rhs170_ = ((d_846_v133_)[d_842_v131_] if (d_842_v131_) in (d_846_v133_) else len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ibbdsuctt"))))
                rhs171_ = 635
                lhs135_ = self.f7
                lhs136_ = default__.safeIndex(7, (self.f7).length(0))
                lhs137_ = globalState
                lhs138_ = globalState
                lhs139_ = globalState
                lhs135_[lhs136_] = rhs167_
                lhs137_.f0 = rhs168_
                d_646_v0_ = rhs169_
                lhs138_.f0 = rhs170_
                lhs139_.f0 = rhs171_
                d_847_v134_: _dafny.Array
                nw140_ = _dafny.Array(D13.default()(), 21)
                d_847_v134_ = nw140_
                d_848_v135_: _dafny.Seq
                d_848_v135_ = _dafny.SeqWithoutIsStrInference([d_847_v134_, d_847_v134_, d_847_v134_])
                d_847_v134_ = (d_848_v135_)[default__.safeIndex(799, len(d_848_v135_))]
                d_849_v136_: _dafny.Map
                d_849_v136_ = _dafny.Map({p0: d_646_v0_})
                d_850_v137_: _dafny.Map
                d_850_v137_ = _dafny.Map({d_849_v136_: p0})
                d_851_v138_: C1
                nw141_ = C1()
                nw141_.ctor__(False, d_671_v11_, (_dafny.SeqWithoutIsStrInference([p0, p0, p0, p0])) < (d_670_v10_), (d_669_v9_).f5)
                d_851_v138_ = nw141_
                d_852_v139_: D12
                d_852_v139_ = D12_DC35(((self.f7)[default__.safeIndex(7, (self.f7).length(0))]) == ((d_851_v138_).f12))
                d_853_v141_: _dafny.Map
                d_853_v141_ = _dafny.Map({d_849_v136_: p1})
                d_854_v143_: _dafny.Map
                def iife91_():
                    coll39_ = _dafny.Map()
                    compr_39_: int
                    for compr_39_ in _dafny.IntegerRange(586, -57):
                        d_855_v142_: int = compr_39_
                        if ((586) <= (d_855_v142_)) and ((d_855_v142_) < (-57)):
                            coll39_[default__.safeModuloInt(d_855_v142_, p0)] = 854
                    return _dafny.Map(coll39_)
                d_854_v143_ = _dafny.Map({iife91_()
                : True})
                d_856_v144_: C0
                nw142_ = C0()
                nw142_.ctor__(d_854_v143_)
                d_856_v144_ = nw142_
                d_857_v145_: _dafny.Set
                d_857_v145_ = _dafny.Set({d_856_v144_, d_856_v144_, d_856_v144_, d_856_v144_})
                def iife92_():
                    coll40_ = _dafny.Map()
                    compr_40_: _dafny.Map
                    for compr_40_ in (d_853_v141_).keys.Elements:
                        d_858_v140_: _dafny.Map = compr_40_
                        if (d_858_v140_) in (d_853_v141_):
                            coll40_[d_858_v140_] = p0
                    return _dafny.Map(coll40_)
                def iife93_(_pat_let26_0):
                    def iife94_(d_859_dt__update__tmp_h1_):
                        def iife95_(_pat_let27_0):
                            def iife96_(d_860_dt__update_hcf55_h0_):
                                return D12_DC35(d_860_dt__update_hcf55_h0_)
                            return iife96_(_pat_let27_0)
                        return iife95_(self.f8)
                    return iife94_(_pat_let26_0)
                rhs172_ = (d_850_v137_) | ((iife92_()
                ) | (d_850_v137_))
                rhs173_ = default__.safeDivisionInt(p0, (p0) + ((0) - (p0)))
                rhs174_ = d_851_v138_
                rhs175_ = len((d_857_v145_ if ((d_851_v138_).f12) or (self.f8) else d_857_v145_))
                rhs176_ = iife93_(d_852_v139_)
                lhs140_ = globalState
                lhs141_ = globalState
                d_850_v137_ = rhs172_
                lhs140_.f0 = rhs173_
                d_851_v138_ = rhs174_
                lhs141_.f0 = rhs175_
                d_852_v139_ = rhs176_
                (globalState).f0 = p0
            d_861_v146_: _dafny.Map
            d_861_v146_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u"))) + (d_671_v11_): p0})
            d_862_v147_: _dafny.Seq
            d_862_v147_ = _dafny.SeqWithoutIsStrInference([d_861_v146_])
            d_861_v146_ = (((d_862_v147_)[default__.safeIndex(len(d_670_v10_), len(d_862_v147_))]) | (_dafny.Map({d_671_v11_: 777}))) | (d_861_v146_)
            def iife97_():
                coll41_ = _dafny.Map()
                compr_41_: int
                for compr_41_ in _dafny.IntegerRange(849, 589):
                    d_863_v148_: int = compr_41_
                    if ((849) <= (d_863_v148_)) and ((d_863_v148_) < (589)):
                        coll41_[default__.safeDivisionInt(d_863_v148_, p0)] = p0
                return _dafny.Map(coll41_)
            if (p2 if (d_648_v2_) != (iife97_()
            ) else ((self).f6) and ((self).f6)):
                d_864_v149_: _dafny.Set
                d_864_v149_ = _dafny.Set({d_671_v11_, d_671_v11_, d_671_v11_, (self).fm7(globalState)})
                d_864_v149_ = d_864_v149_
                d_865_v150_: _dafny.Set
                d_865_v150_ = _dafny.Set({p2, self.f8})
                d_866_v151_: _dafny.Map
                d_866_v151_ = _dafny.Map({(_dafny.Set({p1, p2, p1, self.f8})).issubset(d_865_v150_): p0})
                d_866_v151_ = (d_866_v151_).set((p2 if False else self.f8), p0)
                nw143_ = _dafny.Array(False, 27)
                (self).f7 = nw143_
                d_867_v152_: _dafny.Map
                d_867_v152_ = _dafny.Map({self.f7: self.f8})
                d_868_v153_: D9
                d_868_v153_ = D9_DC27(d_867_v152_, p1)
                d_869_v154_: _dafny.MultiSet
                d_869_v154_ = _dafny.MultiSet([D9_DC27(d_867_v152_, self.f8), d_868_v153_])
                d_870_v155_: D16
                d_870_v155_ = D16_DC43(d_869_v154_)
                d_871_v156_: _dafny.Seq
                d_871_v156_ = _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([d_868_v153_, d_868_v153_])) - (d_869_v154_), (d_870_v155_).cf68, (_dafny.MultiSet([d_868_v153_, d_868_v153_, d_868_v153_])) - (d_869_v154_)])
                (globalState).f0 = ((d_871_v156_)[default__.safeIndex(p0, len(d_871_v156_))]).cardinality
                d_872_v157_: _dafny.Map
                d_872_v157_ = _dafny.Map({d_648_v2_: self.f8})
                d_873_v158_: C0
                nw144_ = C0()
                nw144_.ctor__(d_872_v157_)
                d_873_v158_ = nw144_
                d_873_v158_ = d_873_v158_
            elif True:
                (self).f8 = (self).f6
                arr26_ = self.f7
                index147_ = default__.safeIndex(699, (self.f7).length(0))
                arr26_[index147_] = (p0) in (_dafny.Set({p0}))
                arr27_ = self.f7
                index148_ = default__.safeIndex(699, (self.f7).length(0))
                rhs177_ = (self).fm6(globalState)
                rhs178_ = (122) + (p0)
                rhs179_ = (d_671_v11_).set(default__.safeIndex((self).fm6(globalState), len(d_671_v11_)), d_646_v0_)
                rhs180_ = self.f8
                rhs181_ = self.f8
                lhs142_ = globalState
                lhs143_ = globalState
                lhs144_ = self
                lhs145_ = self.f7
                lhs146_ = default__.safeIndex(699, (self.f7).length(0))
                lhs142_.f0 = rhs177_
                lhs143_.f0 = rhs178_
                d_671_v11_ = rhs179_
                lhs144_.f8 = rhs180_
                lhs145_[lhs146_] = rhs181_
                d_874_v159_: _dafny.Map
                out9_: _dafny.Map
                out9_ = (d_816_v115_).m9(False, p0, globalState)
                d_874_v159_ = out9_
                (globalState).f0 = (self).fm6(globalState)
                d_875_v160_: _dafny.Array
                def lambda38_(d_876_p1_, d_877_v11_, d_878_p0_):
                    def lambda39_(d_879_i14_):
                        return (d_879_i14_) * (len(_dafny.Map({d_876_p1_: len(_dafny.Map({d_877_v11_: _dafny.Set({d_878_p0_, (0) - (d_878_p0_)})}))})))

                    return lambda39_

                init23_ = lambda38_(p1, d_671_v11_, p0)
                nw145_ = _dafny.Array(None, 10)
                for i0_23_ in range(nw145_.length(0)):
                    nw145_[i0_23_] = init23_(i0_23_)
                d_875_v160_ = nw145_
                d_875_v160_ = d_875_v160_


class C5(T2, T1, T0):
    def  __init__(self):
        self._f7: _dafny.Array = _dafny.Array(None, 0)
        self._f8: bool = False
        self._f6: bool = False
        self._f5: _dafny.MultiSet = _dafny.MultiSet({})
        self._f9: int = int(0)
        self._f10: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    @property
    def f7(self):
        return self._f7
    @f7.setter
    def f7(self, value):
        self._f7 = value
    @property
    def f8(self):
        return self._f8
    @f8.setter
    def f8(self, value):
        self._f8 = value
    @property
    def f6(self):
        return self._f6
    @property
    def f5(self):
        return self._f5
    def ctor__(self, f9, f10, f7, f8, f6, f5):
        (self)._f9 = f9
        (self)._f10 = f10
        (self).f7 = f7
        (self).f8 = f8
        (self)._f6 = f6
        (self)._f5 = f5

    def fm3(self, globalState):
        return ((self).f9) >= ((self).f9)

    def fm4(self, p0, p1, p2, p3, globalState):
        return len(((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rnpbenu"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ygnckedh")))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "klfmjbsc"))))

    def fm1(self, p0, p1, p2, globalState):
        return self.f8

    def fm2(self, p0, globalState):
        return (default__.safeDivisionInt((self).f9, (self).f9)) < ((self).f9)

    def m2(self, p0, globalState):
        d_880_v0_: str
        d_880_v0_ = _dafny.CodePoint('r')
        d_880_v0_ = d_880_v0_
        d_881_v1_: _dafny.Array
        def lambda40_(d_882_p0_):
            def lambda41_(d_883_i0_):
                return _dafny.SeqWithoutIsStrInference([d_882_p0_ for d_884_i1_ in range(default__.abs(-489))])

            return lambda41_

        init24_ = lambda40_(p0)
        nw146_ = _dafny.Array(None, 29)
        for i0_24_ in range(nw146_.length(0)):
            nw146_[i0_24_] = init24_(i0_24_)
        d_881_v1_ = nw146_
        d_881_v1_ = d_881_v1_
        d_885_v2_: _dafny.Map
        d_885_v2_ = _dafny.Map({False: (self).f10})
        d_886_v3_: _dafny.Seq
        d_886_v3_ = _dafny.SeqWithoutIsStrInference([not(self.f8), (self).fm1((self).f6, (self).f9, len(d_885_v2_), globalState), (self).f6, True])
        d_887_v4_: _dafny.Seq
        d_887_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nhj"))
        (self).f8 = ((len(d_886_v3_)) < ((self).f9) if self.f8 else (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "khisohk")) for d_888_i2_ in range(default__.abs(608))])) != (_dafny.SeqWithoutIsStrInference([d_887_v4_])))
        d_889_v5_: _dafny.Map
        d_889_v5_ = _dafny.Map({(self.f8) or (not((self).f10)): self.f7})
        d_889_v5_ = d_889_v5_
        if not(self.f8):
            d_890_v6_: _dafny.Seq
            d_890_v6_ = _dafny.SeqWithoutIsStrInference([self.f7, self.f7])
            d_891_v7_: _dafny.Map
            d_891_v7_ = _dafny.Map({(d_890_v6_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([self.f8])), len(d_890_v6_))]: d_880_v0_})
            d_880_v0_ = ((d_891_v7_)[self.f7] if (self.f7) in (d_891_v7_) else p0)
            d_892_v8_: _dafny.Array
            nw147_ = _dafny.Array(_dafny.Set({}), 23)
            d_892_v8_ = nw147_
            d_893_v9_: _dafny.Set
            d_893_v9_ = _dafny.Set({len(d_886_v3_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rkohfvwpq")))})
            index149_ = default__.safeIndex(58, (d_892_v8_).length(0))
            (d_892_v8_)[index149_] = d_893_v9_
            d_894_v10_: _dafny.Map
            d_894_v10_ = _dafny.Map({(self).f9: _dafny.Set({(self).f9, (self).f9, (self).f9})})
            d_895_v11_: _dafny.Map
            d_895_v11_ = _dafny.Map({False: ((d_894_v10_)[(self).f9] if ((self).f9) in (d_894_v10_) else _dafny.Set({(self).f9}))})
            index150_ = default__.safeIndex(58, (d_892_v8_).length(0))
            (d_892_v8_)[index150_] = (d_893_v9_).intersection(((d_895_v11_)[(self).f6] if ((self).f6) in (d_895_v11_) else d_893_v9_))
            if self.f8:
                (globalState).f0 = (self).f9
                index151_ = default__.safeIndex(206, (d_881_v1_).length(0))
                (d_881_v1_)[index151_] = d_887_v4_
                d_896_v12_: _dafny.Map
                d_896_v12_ = _dafny.Map({(self).f9: -811})
                d_897_v13_: _dafny.Seq
                d_897_v13_ = _dafny.SeqWithoutIsStrInference([((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jxhfdgy"))) + (_dafny.SeqWithoutIsStrInference([p0 for d_898_i3_ in range(default__.abs(225))]))).set(default__.safeIndex(len(d_887_v4_), len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jxhfdgy"))) + (_dafny.SeqWithoutIsStrInference([p0 for d_899_i3_ in range(default__.abs(225))])))), p0)])
                d_900_v15_: _dafny.Map
                d_900_v15_ = _dafny.Map({d_887_v4_: self.f8})
                d_901_v16_: D0
                d_901_v16_ = D0_DC1()
                d_902_v17_: _dafny.Set
                d_902_v17_ = _dafny.Set({d_901_v16_})
                d_903_v18_: _dafny.Set
                d_903_v18_ = _dafny.Set({((d_885_v2_)[(self).f10] if ((self).f10) in (d_885_v2_) else ((d_885_v2_)[(self).f10] if ((self).f10) in (d_885_v2_) else self.f8)), (self).f10, self.f8, self.f8})
                index152_ = default__.safeIndex(206, (d_881_v1_).length(0))
                def iife98_():
                    coll42_ = _dafny.Map()
                    compr_42_: int
                    for compr_42_ in _dafny.IntegerRange(-603, 543):
                        d_904_v14_: int = compr_42_
                        if ((-603) <= (d_904_v14_)) and ((d_904_v14_) < (543)):
                            coll42_[default__.safeDivisionInt(d_904_v14_, (self).f9)] = (D0_DC0(22)).cf0
                    return _dafny.Map(coll42_)
                rhs182_ = d_887_v4_
                rhs183_ = iife98_()

                rhs184_ = (self.f8) == ((_dafny.Set({len(d_885_v2_), -810, (self).f9})).isdisjoint(_dafny.Set({(self).fm4(len(_dafny.SeqWithoutIsStrInference([(self).f9, (self).f9, -857, (self).f9, (self).f9])), (self).f9, d_900_v15_, len(d_902_v17_), globalState)})))
                rhs185_ = (d_903_v18_).ispropersubset(d_903_v18_)
                rhs186_ = (d_897_v13_) + (default__.fm5(globalState))
                lhs147_ = d_881_v1_
                lhs148_ = default__.safeIndex(206, (d_881_v1_).length(0))
                lhs149_ = self
                lhs150_ = self
                lhs147_[lhs148_] = rhs182_
                d_896_v12_ = rhs183_
                lhs149_.f8 = rhs184_
                lhs150_.f8 = rhs185_
                d_897_v13_ = rhs186_
                d_905_v19_: _dafny.Map
                d_905_v19_ = _dafny.Map({(self).f10: _dafny.Map({172: (self).f9})})
                d_906_v20_: _dafny.Map
                d_906_v20_ = _dafny.Map({len(d_896_v12_): ((d_905_v19_)[(self).f6] if ((self).f6) in (d_905_v19_) else d_896_v12_)})
                d_907_v21_: _dafny.Seq
                d_907_v21_ = _dafny.SeqWithoutIsStrInference([((d_906_v20_)[(self).f9] if ((self).f9) in (d_906_v20_) else d_896_v12_)])
                d_908_v22_: _dafny.MultiSet
                d_908_v22_ = _dafny.MultiSet([(self).f9])
                d_909_v23_: _dafny.Seq
                d_909_v23_ = _dafny.SeqWithoutIsStrInference([len(_dafny.Map({(self).f9: d_880_v0_}))])
                d_910_v24_: _dafny.Seq
                d_910_v24_ = _dafny.SeqWithoutIsStrInference([d_909_v23_])
                d_911_v25_: _dafny.Map
                d_911_v25_ = _dafny.Map({(self).f9: (d_909_v23_).set(default__.safeIndex((self).f9, len(d_909_v23_)), (self).f9)})
                d_912_v26_: _dafny.Array
                nw148_ = _dafny.Array(None, 24)
                nw148_[int(0)] = (0) - ((self).f9)
                nw148_[int(1)] = (self).f9
                nw148_[int(2)] = ((self).f9) * ((self).f9)
                nw148_[int(3)] = (self).f9
                nw148_[int(4)] = (self).f9
                nw148_[int(5)] = (self).f9
                nw148_[int(6)] = ((self).f9) - ((self).f9)
                nw148_[int(7)] = ((0) - ((self).f9)) * ((self).f9)
                nw148_[int(8)] = (0) - (len((d_907_v21_)[default__.safeIndex((d_908_v22_).cardinality, len(d_907_v21_))]))
                nw148_[int(9)] = (0) - ((self).f9)
                nw148_[int(10)] = (self).f9
                nw148_[int(11)] = -585
                nw148_[int(12)] = (self).f9
                nw148_[int(13)] = (self).f9
                nw148_[int(14)] = (self).f9
                nw148_[int(15)] = len((d_910_v24_)[default__.safeIndex(len((d_887_v4_).set(default__.safeIndex((self).f9, len(d_887_v4_)), p0)), len(d_910_v24_))])
                nw148_[int(16)] = ((self).f9) + (len(d_911_v25_))
                nw148_[int(17)] = default__.safeModuloInt(470, 97)
                nw148_[int(18)] = len(((d_887_v4_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vhffqpp")))).set(default__.safeIndex((self).f9, len((d_887_v4_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vhffqpp"))))), p0))
                nw148_[int(19)] = (self).f9
                nw148_[int(20)] = default__.safeModuloInt(98, (self).f9)
                nw148_[int(21)] = (self).f9
                nw148_[int(22)] = (self).f9
                nw148_[int(23)] = -674
                d_912_v26_ = nw148_
                index153_ = default__.safeIndex(882, (d_912_v26_).length(0))
                (d_912_v26_)[index153_] = (0) - ((self).f9)
                d_913_v27_: D0
                d_913_v27_ = D0_DC0(len(d_886_v3_))
                index154_ = default__.safeIndex(882, (d_912_v26_).length(0))
                (d_912_v26_)[index154_] = (d_913_v27_).cf0
                (globalState).f0 = (self).f9
                d_914_v28_: T2
                nw149_ = C4()
                nw149_.ctor__(self.f7, (self).f10, not(self.f8), (self).f5)
                d_914_v28_ = nw149_
                d_915_v29_: _dafny.Array
                nw150_ = _dafny.Array(None, 10)
                nw150_[int(0)] = d_914_v28_
                nw150_[int(1)] = d_914_v28_
                nw150_[int(2)] = d_914_v28_
                nw150_[int(3)] = d_914_v28_
                nw150_[int(4)] = d_914_v28_
                nw150_[int(5)] = d_914_v28_
                nw150_[int(6)] = d_914_v28_
                nw150_[int(7)] = d_914_v28_
                nw150_[int(8)] = d_914_v28_
                nw150_[int(9)] = d_914_v28_
                d_915_v29_ = nw150_
                index155_ = default__.safeIndex(808, (d_915_v29_).length(0))
                (d_915_v29_)[index155_] = d_914_v28_
                index156_ = default__.safeIndex(808, (d_915_v29_).length(0))
                nw151_ = C4()
                nw151_.ctor__(d_914_v28_.f7, ((d_914_v28_).f5).ispropersubset(_dafny.MultiSet((d_886_v3_).set(default__.safeIndex((d_912_v26_)[default__.safeIndex(882, (d_912_v26_).length(0))], len(d_886_v3_)), (d_914_v28_).f6))), d_914_v28_.f8, (_dafny.MultiSet([self.f8, (self).f6])) - ((self).f5))
                (d_915_v29_)[index156_] = nw151_
            elif True:
                d_916_v30_: _dafny.Array
                nw152_ = _dafny.Array(_dafny.Array(None, 0), 14)
                d_916_v30_ = nw152_
                d_916_v30_ = d_916_v30_
                (self).f8 = False
                (self).f8 = ((self).f5).isdisjoint((self).f5)
                arr28_ = self.f7
                index157_ = default__.safeIndex(916, (self.f7).length(0))
                arr28_[index157_] = (self).f6
                arr29_ = self.f7
                index158_ = default__.safeIndex(916, (self.f7).length(0))
                rhs187_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ooxmqc"))
                rhs188_ = (self).f6
                rhs189_ = self.f8
                rhs190_ = (self).f9
                rhs191_ = d_880_v0_
                lhs151_ = self
                lhs152_ = self.f7
                lhs153_ = default__.safeIndex(916, (self.f7).length(0))
                lhs154_ = globalState
                d_887_v4_ = rhs187_
                lhs151_.f8 = rhs188_
                lhs152_[lhs153_] = rhs189_
                lhs154_.f0 = rhs190_
                d_880_v0_ = rhs191_
                arr30_ = self.f7
                index159_ = default__.safeIndex(916, (self.f7).length(0))
                arr30_[index159_] = ((self).f5).ispropersubset(((self).f5).set(self.f8, default__.abs((0) - ((self).f9))))
            (globalState).f0 = (self).f9
            d_917_v31_: _dafny.Array
            def lambda42_(d_918_i4_):
                return _dafny.Map({(self).f9: (self).f10})

            init25_ = lambda42_
            nw153_ = _dafny.Array(None, 29)
            for i0_25_ in range(nw153_.length(0)):
                nw153_[i0_25_] = init25_(i0_25_)
            d_917_v31_ = nw153_
            rhs192_ = (self).f9
            rhs193_ = d_917_v31_
            lhs155_ = globalState
            lhs155_.f0 = rhs192_
            d_917_v31_ = rhs193_
        elif True:
            (self).f8 = False
            d_919_v32_: C3
            nw154_ = C3()
            nw154_.ctor__((((d_885_v2_)[(self).f6] if ((self).f6) in (d_885_v2_) else (self).f10)) == (self.f8), (self).f5)
            d_919_v32_ = nw154_
            (self).f8 = not (self.f8) or ((self).f10)
            d_920_v33_: _dafny.Array
            nw155_ = _dafny.Array(int(0), 18)
            d_920_v33_ = nw155_
            d_921_v34_: _dafny.Seq
            d_921_v34_ = _dafny.SeqWithoutIsStrInference([self.f7])
            d_922_v35_: _dafny.Set
            d_922_v35_ = _dafny.Set({(self).f9})
            d_923_v36_: _dafny.MultiSet
            d_923_v36_ = _dafny.MultiSet([self.f7, self.f7, (d_921_v34_)[default__.safeIndex(len(d_922_v35_), len(d_921_v34_))], self.f7, self.f7])
            d_924_v37_: _dafny.Map
            d_924_v37_ = _dafny.Map({d_920_v33_: (0) - (((self).f9) * (((d_923_v36_)[self.f7] if (self.f7) in (d_923_v36_) else (self).f9)))})
            d_925_v38_: _dafny.Map
            d_925_v38_ = _dafny.Map({(self).f9: 289})
            d_926_v39_: C1
            nw156_ = C1()
            nw156_.ctor__(self.f8, _dafny.SeqWithoutIsStrInference([p0 for d_927_i5_ in range(default__.abs(500))]), (d_919_v32_).fm3(globalState), default__.fm14(d_925_v38_, self.f8, (self).f9, (self).f10, globalState))
            d_926_v39_ = nw156_
            d_928_v40_: _dafny.MultiSet
            d_928_v40_ = _dafny.MultiSet([d_926_v39_])
            d_924_v37_ = (d_924_v37_).set(d_920_v33_, (d_928_v40_).cardinality)
            d_929_v41_: _dafny.MultiSet
            d_929_v41_ = _dafny.MultiSet([(self).f9])
            d_930_v42_: D2
            d_930_v42_ = D2_DC7(p0, d_925_v38_, (self).f9)
            d_931_v43_: _dafny.Map
            d_931_v43_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([p0 for d_932_i6_ in range(default__.abs(-980))]): self.f8})
            (globalState).f0 = (d_926_v39_).fm4(((d_929_v41_)[608] if (608) in (d_929_v41_) else (self).f9), (self).f9, (default__.fm13((self).f9, d_930_v42_, (self).f10, len(d_886_v3_), globalState)) | (d_931_v43_), ((self).f9) * ((0) - ((self).f9)), globalState)
        d_933_v44_: _dafny.Map
        d_933_v44_ = _dafny.Map({(0) - (len(_dafny.Map({588: (self).f6}))): (self).f9})
        d_934_v45_: D2
        d_934_v45_ = D2_DC7(d_880_v0_, d_933_v44_, (self).f9)
        d_935_v46_: D12
        d_935_v46_ = D12_DC36(321, (self).f9)
        d_936_v47_: D15
        def iife99_(_pat_let28_0):
            def iife100_(d_937_dt__update__tmp_h0_):
                def iife101_(_pat_let29_0):
                    def iife102_(d_938_dt__update_hcf57_h0_):
                        return D12_DC36((d_937_dt__update__tmp_h0_).cf56, d_938_dt__update_hcf57_h0_)
                    return iife102_(_pat_let29_0)
                return iife101_((self).f9)
            return iife100_(_pat_let28_0)
        d_936_v47_ = D15_DC41(default__.fm13((self).f9, d_934_v45_, (self).f10, (iife99_(d_935_v46_)).cf57, globalState))
        d_939_i7_: int
        d_939_i7_ = 0
        with _dafny.label("9"):
            def lambda43_(source10_):
                if source10_.is_DC42:
                    d_942___mcc_h0_ = source10_.cf67
                    d_943_cf67_ = d_942___mcc_h0_
                    return (len(_dafny.Set({(self).f10, self.f8}))) == (len(_dafny.Set({(self).f10})))
                elif True:
                    d_944___mcc_h1_ = source10_.cf66
                    d_945_cf66_ = d_944___mcc_h1_
                    return (self).f6

            while lambda43_(d_936_v47_):
                with _dafny.c_label("9"):
                    if (d_939_i7_) >= (100):
                        raise _dafny.Break("9")
                    d_939_i7_ = (d_939_i7_) + (1)
                    d_940_v48_: C1
                    nw157_ = C1()
                    nw157_.ctor__((self.f8 if self.f8 else self.f8), (d_887_v4_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ceqq"))), (self).f6, (self).f5)
                    d_940_v48_ = nw157_
                    d_941_v49_: C4
                    nw158_ = C4()
                    nw158_.ctor__(self.f7, self.f8, (self).f10, ((self).f5) | (_dafny.MultiSet([not((self).f6)])))
                    d_941_v49_ = nw158_
                    (self).f8 = (self).f10
                    arr31_ = self.f7
                    index160_ = default__.safeIndex(379, (self.f7).length(0))
                    arr31_[index160_] = not((d_941_v49_).fm3(globalState))
                    arr32_ = self.f7
                    index161_ = default__.safeIndex(379, (self.f7).length(0))
                    arr32_[index161_] = (((d_940_v48_).f12 if (self).f10 else (self).f6)) not in (d_886_v3_)
                    pass
            pass

    def m3(self, globalState):
        r0: _dafny.Set = _dafny.Set({})
        r1: int = int(0)
        r2: int = int(0)
        arr33_ = self.f7
        index162_ = default__.safeIndex(238, (self.f7).length(0))
        arr33_[index162_] = not((self).f10)
        d_946_v0_: _dafny.Seq
        d_946_v0_ = _dafny.SeqWithoutIsStrInference([(self).f6, (self).f10])
        d_947_v1_: D10
        d_947_v1_ = D10_DC29(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_946_v0_])))
        d_948_v2_: _dafny.MultiSet
        d_948_v2_ = _dafny.MultiSet([d_947_v1_, d_947_v1_])
        arr34_ = self.f7
        index163_ = default__.safeIndex(238, (self.f7).length(0))
        arr34_[index163_] = not(((d_948_v2_).intersection(d_948_v2_)).issubset(d_948_v2_))
        d_949_v3_: D3
        d_949_v3_ = D3_DC10(not(self.f8), (self).f9)
        arr35_ = self.f7
        index164_ = default__.safeIndex(238, (self.f7).length(0))
        def lambda44_(source11_):
            if source11_.is_DC10:
                d_950___mcc_h0_ = source11_.cf11
                d_951___mcc_h1_ = source11_.cf12
                d_952_cf12_ = d_951___mcc_h1_
                d_953_cf11_ = d_950___mcc_h0_
                return (self).f10
            elif source11_.is_DC11:
                d_954___mcc_h2_ = source11_.cf13
                d_955_cf13_ = d_954___mcc_h2_
                return (self).f10
            elif True:
                d_956___mcc_h3_ = source11_.cf10
                d_957_cf10_ = d_956___mcc_h3_
                return (self).f6

        arr35_[index164_] = not(not(lambda44_(d_949_v3_)))
        d_958_v4_: _dafny.Seq
        d_958_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yie"))
        d_959_v5_: _dafny.Map
        d_959_v5_ = _dafny.Map({(d_958_v4_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kynauahc"))): (self).f6})
        d_959_v5_ = (d_959_v5_).set(d_958_v4_, not(self.f8))
        d_960_v6_: _dafny.MultiSet
        d_960_v6_ = _dafny.MultiSet([(self).f9])
        d_961_v7_: _dafny.Array
        nw159_ = _dafny.Array(None, 17)
        nw159_[int(0)] = d_946_v0_
        nw159_[int(1)] = d_946_v0_
        nw159_[int(2)] = (d_946_v0_) + (d_946_v0_)
        nw159_[int(3)] = d_946_v0_
        nw159_[int(4)] = (_dafny.SeqWithoutIsStrInference([not(True), (self.f7)[default__.safeIndex(238, (self.f7).length(0))]])) + (d_946_v0_)
        nw159_[int(5)] = default__.fm23(((d_960_v6_)[(self).f9] if ((self).f9) in (d_960_v6_) else (self).f9), globalState)
        nw159_[int(6)] = (d_946_v0_) + (d_946_v0_)
        nw159_[int(7)] = (d_946_v0_).set(default__.safeIndex((self).f9, len(d_946_v0_)), (self.f7)[default__.safeIndex(238, (self.f7).length(0))])
        nw159_[int(8)] = d_946_v0_
        nw159_[int(9)] = _dafny.SeqWithoutIsStrInference([(self).f6])
        nw159_[int(10)] = (d_946_v0_) + (d_946_v0_)
        nw159_[int(11)] = d_946_v0_
        nw159_[int(12)] = (d_946_v0_) + (d_946_v0_)
        nw159_[int(13)] = default__.fm23((self).f9, globalState)
        nw159_[int(14)] = (d_946_v0_) + ((d_946_v0_).set(default__.safeIndex((self).f9, len(d_946_v0_)), False))
        nw159_[int(15)] = (_dafny.SeqWithoutIsStrInference([self.f8])) + ((_dafny.SeqWithoutIsStrInference([(self).f10, not((self.f7)[default__.safeIndex(238, (self.f7).length(0))])])).set(default__.safeIndex((self).f9, len(_dafny.SeqWithoutIsStrInference([(self).f10, not((self.f7)[default__.safeIndex(238, (self.f7).length(0))])]))), (self).f10))
        nw159_[int(16)] = (default__.fm23(712, globalState)) + (d_946_v0_)
        d_961_v7_ = nw159_
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_961_v7_).length(0)):
            d_962_i0_: int = guard_loop_3_
            if (True) and (((0) <= (d_962_i0_)) and ((d_962_i0_) < ((d_961_v7_).length(0)))):
                (d_961_v7_)[(d_962_i0_)] = d_946_v0_
        d_963_v8_: D0
        d_963_v8_ = D0_DC1()
        d_963_v8_ = d_963_v8_
        d_964_v9_: str
        d_964_v9_ = _dafny.CodePoint('b')
        d_965_v10_: D10
        d_965_v10_ = D10_DC31(_dafny.SeqWithoutIsStrInference([(self).f5]), d_964_v9_)
        def lambda45_(source12_):
            if source12_.is_DC30:
                d_966___mcc_h4_ = source12_.cf46
                d_967___mcc_h5_ = source12_.cf47
                d_968___mcc_h6_ = source12_.cf48
                d_969_cf48_ = d_968___mcc_h6_
                d_970_cf47_ = d_967___mcc_h5_
                d_971_cf46_ = d_966___mcc_h4_
                return (self.f7)[default__.safeIndex(238, (self.f7).length(0))]
            elif source12_.is_DC31:
                d_972___mcc_h7_ = source12_.cf49
                d_973___mcc_h8_ = source12_.cf50
                d_974_cf50_ = d_973___mcc_h8_
                d_975_cf49_ = d_972___mcc_h7_
                return (self).f10
            elif True:
                d_976___mcc_h9_ = source12_.cf45
                d_977_cf45_ = d_976___mcc_h9_
                return (self).f6

        (self).f8 = lambda45_(d_965_v10_)
        d_978_v11_: D12
        d_978_v11_ = D12_DC34(_dafny.Map({d_964_v9_: self.f8}))
        d_979_v12_: _dafny.Set
        d_979_v12_ = _dafny.Set({(0) - ((self).f9), 32, default__.safeModuloInt(39, (self).f9), len(_dafny.Map({(self).f10: d_978_v11_})), len(d_958_v4_)})
        r0 = d_979_v12_
        d_980_v13_: _dafny.Array
        def lambda46_(d_981_i1_):
            return default__.safeModuloInt(d_981_i1_, (self).f9)

        init26_ = lambda46_
        nw160_ = _dafny.Array(None, 4)
        for i0_26_ in range(nw160_.length(0)):
            nw160_[i0_26_] = init26_(i0_26_)
        d_980_v13_ = nw160_
        d_982_v14_: _dafny.MultiSet
        d_982_v14_ = _dafny.MultiSet([d_980_v13_, d_980_v13_, d_980_v13_])
        d_983_v15_: _dafny.Map
        d_983_v15_ = _dafny.Map({(self).f9: (d_982_v14_).set(d_980_v13_, default__.abs((self).f9))})
        r1 = (self).fm4(217, (self).f9, d_959_v5_, ((((d_983_v15_)[(self).f9] if ((self).f9) in (d_983_v15_) else d_982_v14_)) | (d_982_v14_)).cardinality, globalState)
        r2 = (self).f9
        return r0, r1, r2

    def m1(self, p0, p1, p2, globalState):
        d_984_i0_: int
        d_984_i0_ = 0
        with _dafny.label("10"):
            while (self).f10:
                with _dafny.c_label("10"):
                    if (d_984_i0_) >= (100):
                        raise _dafny.Break("10")
                    d_984_i0_ = (d_984_i0_) + (1)
                    d_985_v0_: _dafny.Array
                    def lambda47_(d_986_i1_):
                        return (d_986_i1_) * ((self).f9)

                    init27_ = lambda47_
                    nw161_ = _dafny.Array(None, 21)
                    for i0_27_ in range(nw161_.length(0)):
                        nw161_[i0_27_] = init27_(i0_27_)
                    d_985_v0_ = nw161_
                    d_987_v1_: D3
                    d_987_v1_ = D3_DC9(d_985_v0_)
                    d_988_v2_: _dafny.Seq
                    d_988_v2_ = _dafny.SeqWithoutIsStrInference([p0])
                    d_989_v3_: _dafny.Seq
                    d_989_v3_ = _dafny.SeqWithoutIsStrInference([not((self).f6)])
                    d_990_v4_: _dafny.Map
                    d_990_v4_ = _dafny.Map({487: 828})
                    d_991_v5_: _dafny.Map
                    d_991_v5_ = _dafny.Map({d_990_v4_: (self).f6})
                    d_992_v6_: C0
                    nw162_ = C0()
                    nw162_.ctor__(d_991_v5_)
                    d_992_v6_ = nw162_
                    d_993_v7_: _dafny.Set
                    d_993_v7_ = _dafny.Set({(self).f6, self.f8, (self).f6, (self).f6, p2})
                    rhs194_ = (p0) * ((D5_DC16(d_987_v1_, len(d_988_v2_), 40, len(d_989_v3_), d_992_v6_)).cf20)
                    rhs195_ = (((d_990_v4_)[p0] if (p0) in (d_990_v4_) else len(d_993_v7_))) in (_dafny.SeqWithoutIsStrInference([p0, -625]))
                    rhs196_ = (d_988_v2_)[default__.safeIndex((self).f9, len(d_988_v2_))]
                    rhs197_ = (p0) > ((((self).f5)[p1] if (p1) in ((self).f5) else (self).f9))
                    lhs156_ = globalState
                    lhs157_ = self
                    lhs158_ = globalState
                    lhs159_ = self
                    lhs156_.f0 = rhs194_
                    lhs157_.f8 = rhs195_
                    lhs158_.f0 = rhs196_
                    lhs159_.f8 = rhs197_
                    (globalState).f0 = (self).f9
                    index165_ = default__.safeIndex(417, (d_985_v0_).length(0))
                    (d_985_v0_)[index165_] = (p0) - (p0)
                    index166_ = default__.safeIndex(417, (d_985_v0_).length(0))
                    (d_985_v0_)[index166_] = (0) - (default__.safeModuloInt(p0, default__.safeModuloInt(p0, (0) - (p0))))
                    d_994_v8_: _dafny.Seq
                    d_994_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bobaat"))
                    d_994_v8_ = (d_994_v8_ if False else (d_994_v8_) + (d_994_v8_))
                    pass
            pass
        d_995_v9_: _dafny.Set
        d_995_v9_ = _dafny.Set({(self).f6})
        d_996_v10_: _dafny.Seq
        d_996_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rhu"))
        d_997_v11_: D5
        d_997_v11_ = D5_DC15(d_996_v10_)
        d_998_v12_: _dafny.Map
        d_998_v12_ = _dafny.Map({(d_997_v11_).cf18: (self).f10})
        (globalState).f0 = ((self).fm4(len(d_995_v9_), (self).f9, d_998_v12_, (0) - ((self).f9), globalState)) * (p0)
        pat_let_tv20_ = p0
        pat_let_tv21_ = p0
        pat_let_tv22_ = p1
        pat_let_tv23_ = p1
        pat_let_tv24_ = p0
        def lambda48_(source13_):
            if source13_.is_DC20:
                d_999___mcc_h0_ = source13_.cf29
                d_1000_cf29_ = d_999___mcc_h0_
                return default__.safeModuloInt(pat_let_tv20_, pat_let_tv21_)
            elif source13_.is_DC21:
                d_1001___mcc_h1_ = source13_.cf30
                d_1002___mcc_h2_ = source13_.cf31
                d_1003_cf31_ = d_1002___mcc_h2_
                d_1004_cf30_ = d_1001___mcc_h1_
                return len(_dafny.Map({False: pat_let_tv22_}))
            elif source13_.is_DC19:
                d_1005___mcc_h3_ = source13_.cf28
                d_1006_cf28_ = d_1005___mcc_h3_
                return (self).f9
            elif True:
                d_1007___mcc_h4_ = source13_.cf32
                d_1008_cf32_ = d_1007___mcc_h4_
                return default__.safeModuloInt((_dafny.MultiSet([pat_let_tv23_, False])).cardinality, pat_let_tv24_)

        (globalState).f0 = lambda48_(D7_DC22(D7_DC21(p0, 734)))
        arr36_ = self.f7
        index167_ = default__.safeIndex(617, (self.f7).length(0))
        arr36_[index167_] = p1
        arr37_ = self.f7
        index168_ = default__.safeIndex(617, (self.f7).length(0))
        arr37_[index168_] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ibfacfd"))) < (d_996_v10_)
        d_1009_v13_: D7
        d_1009_v13_ = D7_DC21(p0, 383)
        d_1010_v14_: _dafny.MultiSet
        d_1010_v14_ = _dafny.MultiSet([d_1009_v13_])
        d_1011_v15_: _dafny.Seq
        d_1011_v15_ = _dafny.SeqWithoutIsStrInference([(self).f9, (d_1010_v14_).cardinality])
        d_1012_v16_: _dafny.Map
        d_1012_v16_ = _dafny.Map({_dafny.CodePoint('j'): (False) or (not((self).f10))})
        arr38_ = self.f7
        index169_ = default__.safeIndex(617, (self.f7).length(0))
        rhs198_ = (d_1011_v15_) + (_dafny.SeqWithoutIsStrInference([p0 for d_1013_i2_ in range(default__.abs(-447))]))
        rhs199_ = not(((d_1012_v16_)[(d_996_v10_)[default__.safeIndex((self).fm4(p0, (0) - ((self).f9), d_998_v12_, 464, globalState), len(d_996_v10_))]] if ((d_996_v10_)[default__.safeIndex((self).fm4(p0, (0) - ((self).f9), d_998_v12_, 464, globalState), len(d_996_v10_))]) in (d_1012_v16_) else (p0) == (p0)))
        lhs160_ = self.f7
        lhs161_ = default__.safeIndex(617, (self.f7).length(0))
        d_1011_v15_ = rhs198_
        lhs160_[lhs161_] = rhs199_
        d_1014_v17_: _dafny.Map
        d_1014_v17_ = _dafny.Map({(self).f9: len(d_996_v10_)})
        d_1015_v18_: _dafny.Map
        d_1015_v18_ = _dafny.Map({p0: ((d_1014_v17_)[(self).f9] if ((self).f9) in (d_1014_v17_) else (self).f9)})
        d_1016_i3_: int
        d_1016_i3_ = 0
        with _dafny.label("11"):
            while ((len(d_1011_v15_)) + (len(d_996_v10_))) >= (((default__.fm14(d_1015_v18_, False, 692, (self.f7)[default__.safeIndex(617, (self.f7).length(0))], globalState)).set(p1, default__.abs((d_1011_v15_)[default__.safeIndex((0) - ((self).f9), len(d_1011_v15_))]))).cardinality):
                with _dafny.c_label("11"):
                    if (d_1016_i3_) >= (100):
                        raise _dafny.Break("11")
                    d_1016_i3_ = (d_1016_i3_) + (1)
                    d_995_v9_ = _dafny.Set({(self).f6, p2})
                    arr39_ = self.f7
                    index170_ = default__.safeIndex(617, (self.f7).length(0))
                    arr39_[index170_] = (self).f6
                    d_1017_v19_: T0
                    nw163_ = C2()
                    nw163_.ctor__(_dafny.MultiSet([(self.f7)[default__.safeIndex(617, (self.f7).length(0))], (self.f7)[default__.safeIndex(617, (self.f7).length(0))], (self.f7)[default__.safeIndex(617, (self.f7).length(0))], True]))
                    d_1017_v19_ = nw163_
                    d_1018_v20_: _dafny.MultiSet
                    d_1018_v20_ = _dafny.MultiSet([d_1017_v19_, d_1017_v19_, d_1017_v19_])
                    d_1019_v21_: _dafny.Map
                    d_1019_v21_ = _dafny.Map({p2: self.f8})
                    d_1020_v22_: _dafny.Seq
                    d_1020_v22_ = _dafny.SeqWithoutIsStrInference([d_1019_v21_])
                    (self).f8 = (self).fm1((self).f6, ((d_1018_v20_)[d_1017_v19_] if (d_1017_v19_) in (d_1018_v20_) else len(d_1020_v22_)), (self).f9, globalState)
                    arr40_ = self.f7
                    index171_ = default__.safeIndex(617, (self.f7).length(0))
                    arr40_[index171_] = True
                    pass
            pass

    def m4(self, p0, p1, globalState):
        d_1021_v0_: _dafny.Map
        d_1021_v0_ = _dafny.Map({(self).f6: (self).f9})
        d_1021_v0_ = (d_1021_v0_).set((self).fm2((self).f10, globalState), (self).f9)
        d_1022_v1_: _dafny.Seq
        d_1022_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nttgnqle"))
        d_1023_v2_: _dafny.Map
        d_1023_v2_ = _dafny.Map({d_1022_v1_: True})
        d_1024_v3_: D16
        d_1024_v3_ = D16_DC44((self).fm2((self).f6, globalState), (self).fm3(globalState), (self).fm4((self).f9, (self).f9, d_1023_v2_, (self).f9, globalState))
        def lambda49_(source15_):
            if source15_.is_DC44:
                d_1025___mcc_h4_ = source15_.cf69
                d_1026___mcc_h5_ = source15_.cf70
                d_1027___mcc_h6_ = source15_.cf71
                d_1028_cf71_ = d_1027___mcc_h6_
                d_1029_cf70_ = d_1026___mcc_h5_
                d_1030_cf69_ = d_1025___mcc_h4_
                return D6_DC17(_dafny.SeqWithoutIsStrInference([not(d_1029_cf70_), d_1029_cf70_, True]))
            elif source15_.is_DC43:
                d_1031___mcc_h7_ = source15_.cf68
                d_1032_cf68_ = d_1031___mcc_h7_
                return D6_DC17(_dafny.SeqWithoutIsStrInference([self.f8]))
            elif True:
                d_1033___mcc_h8_ = source15_.cf72
                d_1034_cf72_ = d_1033___mcc_h8_
                return D6_DC17(_dafny.SeqWithoutIsStrInference([self.f8]))

        source14_ = lambda49_(d_1024_v3_)
        if source14_.is_DC18:
            d_1035___mcc_h0_ = source14_.cf25
            d_1036___mcc_h1_ = source14_.cf26
            d_1037___mcc_h2_ = source14_.cf27
            d_1038_cf27_ = d_1037___mcc_h2_
            d_1039_cf26_ = d_1036___mcc_h1_
            d_1040_cf25_ = d_1035___mcc_h0_
            d_1041_v4_: _dafny.Array
            nw164_ = _dafny.Array(int(0), 13)
            d_1041_v4_ = nw164_
            d_1042_v5_: _dafny.Map
            d_1042_v5_ = _dafny.Map({p0: d_1041_v4_})
            d_1042_v5_ = (d_1042_v5_).set(False, d_1041_v4_)
            d_1043_v6_: str
            d_1043_v6_ = _dafny.CodePoint('l')
            d_1044_v7_: _dafny.Array
            nw165_ = _dafny.Array(None, 8)
            nw165_[int(0)] = d_1043_v6_
            nw165_[int(1)] = d_1043_v6_
            nw165_[int(2)] = d_1043_v6_
            nw165_[int(3)] = d_1043_v6_
            nw165_[int(4)] = d_1043_v6_
            nw165_[int(5)] = d_1043_v6_
            nw165_[int(6)] = d_1043_v6_
            nw165_[int(7)] = d_1043_v6_
            d_1044_v7_ = nw165_
            d_1045_v8_: _dafny.Array
            nw166_ = _dafny.Array(None, 13)
            nw166_[int(0)] = d_1044_v7_
            nw166_[int(1)] = d_1044_v7_
            nw166_[int(2)] = d_1044_v7_
            nw166_[int(3)] = d_1044_v7_
            nw166_[int(4)] = d_1044_v7_
            nw166_[int(5)] = d_1044_v7_
            nw166_[int(6)] = d_1044_v7_
            nw166_[int(7)] = d_1044_v7_
            nw166_[int(8)] = d_1044_v7_
            nw166_[int(9)] = d_1044_v7_
            nw166_[int(10)] = d_1044_v7_
            nw166_[int(11)] = d_1044_v7_
            nw166_[int(12)] = d_1044_v7_
            d_1045_v8_ = nw166_
            d_1046_v9_: D4
            d_1046_v9_ = D4_DC12(d_1045_v8_)
            d_1047_v10_: D4
            d_1047_v10_ = D4_DC14(d_1046_v9_)
            d_1048_v11_: D4
            d_1048_v11_ = D4_DC14(d_1047_v10_)
            pat_let_tv25_ = d_1047_v10_
            pat_let_tv26_ = d_1047_v10_
            d_1049_v12_: _dafny.Array
            nw167_ = _dafny.Array(None, 20)
            nw167_[int(0)] = d_1048_v11_
            nw167_[int(1)] = d_1048_v11_
            nw167_[int(2)] = d_1048_v11_
            nw167_[int(3)] = d_1048_v11_
            nw167_[int(4)] = d_1048_v11_
            nw167_[int(5)] = d_1048_v11_
            nw167_[int(6)] = d_1048_v11_
            nw167_[int(7)] = d_1048_v11_
            nw167_[int(8)] = d_1048_v11_
            nw167_[int(9)] = D4_DC14(d_1047_v10_)
            nw167_[int(10)] = d_1048_v11_
            nw167_[int(11)] = d_1048_v11_
            nw167_[int(12)] = d_1048_v11_
            nw167_[int(13)] = d_1048_v11_
            nw167_[int(14)] = D4_DC14((d_1048_v11_).cf17)
            def iife103_(_pat_let30_0):
                def iife104_(d_1050_dt__update__tmp_h0_):
                    def iife105_(_pat_let31_0):
                        def iife106_(d_1051_dt__update_hcf17_h0_):
                            return D4_DC14(d_1051_dt__update_hcf17_h0_)
                        return iife106_(_pat_let31_0)
                    return iife105_(pat_let_tv25_)
                return iife104_(_pat_let30_0)
            nw167_[int(15)] = iife103_(D4_DC14(d_1047_v10_))
            nw167_[int(16)] = d_1048_v11_
            nw167_[int(17)] = d_1048_v11_
            nw167_[int(18)] = d_1048_v11_
            def iife107_(_pat_let32_0):
                def iife108_(d_1052_dt__update__tmp_h1_):
                    def iife109_(_pat_let33_0):
                        def iife110_(d_1053_dt__update_hcf17_h1_):
                            return D4_DC14(d_1053_dt__update_hcf17_h1_)
                        return iife110_(_pat_let33_0)
                    return iife109_(pat_let_tv26_)
                return iife108_(_pat_let32_0)
            nw167_[int(19)] = iife107_(d_1048_v11_)
            d_1049_v12_ = nw167_
            index172_ = default__.safeIndex(539, (d_1049_v12_).length(0))
            (d_1049_v12_)[index172_] = d_1048_v11_
            index173_ = default__.safeIndex(539, (d_1049_v12_).length(0))
            (d_1049_v12_)[index173_] = d_1048_v11_
            d_1054_v13_: C3
            nw168_ = C3()
            nw168_.ctor__((984) > ((self).f9), (self).f5)
            d_1054_v13_ = nw168_
            d_1055_v14_: C3
            nw169_ = C3()
            nw169_.ctor__((self).f6, (self).f5)
            d_1055_v14_ = nw169_
        elif True:
            d_1056___mcc_h3_ = source14_.cf24
            d_1057_cf24_ = d_1056___mcc_h3_
            d_1058_v15_: _dafny.Map
            d_1058_v15_ = _dafny.Map({True: self.f8})
            d_1059_v16_: _dafny.Array
            nw170_ = _dafny.Array(None, 14)
            nw170_[int(0)] = len(d_1058_v15_)
            nw170_[int(1)] = ((self).f9) * ((self).f9)
            nw170_[int(2)] = (self).f9
            nw170_[int(3)] = (self).f9
            nw170_[int(4)] = (self).f9
            nw170_[int(5)] = (self).f9
            nw170_[int(6)] = len(p1)
            nw170_[int(7)] = (0) - (((self).f9) * ((0) - ((self).f9)))
            nw170_[int(8)] = (self).f9
            nw170_[int(9)] = (self).f9
            nw170_[int(10)] = (self).f9
            nw170_[int(11)] = (self).f9
            nw170_[int(12)] = 635
            nw170_[int(13)] = (0) - ((self).f9)
            d_1059_v16_ = nw170_
            index174_ = default__.safeIndex(232, (d_1059_v16_).length(0))
            (d_1059_v16_)[index174_] = 600
            index175_ = default__.safeIndex(232, (d_1059_v16_).length(0))
            (d_1059_v16_)[index175_] = (self).f9
            d_1060_v17_: _dafny.Set
            d_1060_v17_ = _dafny.Set({(self).f10})
            d_1061_v18_: D0
            d_1061_v18_ = D0_DC0(len(d_1060_v17_))
            d_1062_v19_: _dafny.Map
            d_1062_v19_ = _dafny.Map({d_1061_v18_: (d_1022_v1_)[default__.safeIndex((self).f9, len(d_1022_v1_))]})
            d_1063_v20_: str
            d_1063_v20_ = _dafny.CodePoint('k')
            d_1062_v19_ = (d_1062_v19_).set((D0_DC0((self).f9) if self.f8 else d_1061_v18_), d_1063_v20_)
            d_1064_v21_: _dafny.Map
            d_1064_v21_ = _dafny.Map({(self).f9: not(p0)})
            if not((True if ((d_1064_v21_)[(self).f9] if ((self).f9) in (d_1064_v21_) else self.f8) else self.f8)):
                index176_ = default__.safeIndex(232, (d_1059_v16_).length(0))
                (d_1059_v16_)[index176_] = (self).f9
                index177_ = default__.safeIndex(232, (d_1059_v16_).length(0))
                rhs200_ = (True) and ((self).f10)
                rhs201_ = ((d_1022_v1_) + (d_1022_v1_)) + ((d_1022_v1_) + (d_1022_v1_))
                rhs202_ = (self).f6
                rhs203_ = (self).f6
                rhs204_ = (d_1059_v16_)[default__.safeIndex(232, (d_1059_v16_).length(0))]
                lhs162_ = self
                lhs163_ = self
                lhs164_ = self
                lhs165_ = d_1059_v16_
                lhs166_ = default__.safeIndex(232, (d_1059_v16_).length(0))
                lhs162_.f8 = rhs200_
                d_1022_v1_ = rhs201_
                lhs163_.f8 = rhs202_
                lhs164_.f8 = rhs203_
                lhs165_[lhs166_] = rhs204_
                index178_ = default__.safeIndex(232, (d_1059_v16_).length(0))
                def iife111_():
                    coll43_ = _dafny.Set()
                    compr_43_: int
                    for compr_43_ in _dafny.IntegerRange(314, 504):
                        d_1065_v22_: int = compr_43_
                        if ((314) <= (d_1065_v22_)) and ((d_1065_v22_) < (504)):
                            coll43_ = coll43_.union(_dafny.Set([(d_1065_v22_) * (((d_1021_v0_)[False] if (False) in (d_1021_v0_) else (d_1059_v16_)[default__.safeIndex(232, (d_1059_v16_).length(0))]))]))
                    return _dafny.Set(coll43_)
                rhs205_ = d_1022_v1_
                rhs206_ = (self).fm4((d_1059_v16_)[default__.safeIndex(232, (d_1059_v16_).length(0))], (d_1059_v16_)[default__.safeIndex(232, (d_1059_v16_).length(0))], d_1023_v2_, (d_1059_v16_)[default__.safeIndex(232, (d_1059_v16_).length(0))], globalState)
                rhs207_ = (default__.safeDivisionInt(len((p1).set(default__.safeIndex((d_1059_v16_)[default__.safeIndex(232, (d_1059_v16_).length(0))], len(p1)), -872)), (self).fm4(len(d_1058_v15_), len(d_1022_v1_), d_1023_v2_, len(d_1057_cf24_), globalState)) if (d_1057_cf24_)[default__.safeIndex((d_1059_v16_)[default__.safeIndex(232, (d_1059_v16_).length(0))], len(d_1057_cf24_))] else (d_1059_v16_)[default__.safeIndex(232, (d_1059_v16_).length(0))])
                rhs208_ = (not (self.f8) or ((self).f6)) or (True)
                rhs209_ = len(_dafny.SeqWithoutIsStrInference([(_dafny.Set({(d_1059_v16_)[default__.safeIndex(232, (d_1059_v16_).length(0))]})) | (iife111_()
) for d_1066_i0_ in range(default__.abs(452))]))
                lhs167_ = globalState
                lhs168_ = globalState
                lhs169_ = self
                lhs170_ = d_1059_v16_
                lhs171_ = default__.safeIndex(232, (d_1059_v16_).length(0))
                d_1022_v1_ = rhs205_
                lhs167_.f0 = rhs206_
                lhs168_.f0 = rhs207_
                lhs169_.f8 = rhs208_
                lhs170_[lhs171_] = rhs209_
                index179_ = default__.safeIndex(232, (d_1059_v16_).length(0))
                (d_1059_v16_)[index179_] = ((self).f9 if p0 else (d_1059_v16_)[default__.safeIndex(232, (d_1059_v16_).length(0))])
                d_1067_v23_: C2
                nw171_ = C2()
                nw171_.ctor__((self).f5)
                d_1067_v23_ = nw171_
                d_1068_v24_: _dafny.Set
                d_1068_v24_ = _dafny.Set({(0) - ((self).f9)})
                rhs210_ = d_1067_v23_
                rhs211_ = default__.safeDivisionInt(((self).f9) * ((self).fm4((d_1059_v16_)[default__.safeIndex(232, (d_1059_v16_).length(0))], (d_1059_v16_)[default__.safeIndex(232, (d_1059_v16_).length(0))], d_1023_v2_, len(d_1068_v24_), globalState)), (d_1059_v16_)[default__.safeIndex(232, (d_1059_v16_).length(0))])
                lhs172_ = globalState
                d_1067_v23_ = rhs210_
                lhs172_.f0 = rhs211_
            elif True:
                d_1069_v25_: _dafny.Seq
                d_1069_v25_ = _dafny.SeqWithoutIsStrInference([(d_1058_v15_).set(self.f8, self.f8)])
                (self).f8 = ((_dafny.SeqWithoutIsStrInference([d_1058_v15_ for d_1070_i1_ in range(default__.abs(652))])) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({self.f8: (d_1057_cf24_)[default__.safeIndex((d_1059_v16_)[default__.safeIndex(232, (d_1059_v16_).length(0))], len(d_1057_cf24_))]}) for d_1071_i2_ in range(default__.abs(110))]))) < (d_1069_v25_)
                (self).f8 = (self).f6
                arr41_ = self.f7
                index180_ = default__.safeIndex(848, (self.f7).length(0))
                arr41_[index180_] = (self).f6
                arr42_ = self.f7
                index181_ = default__.safeIndex(848, (self.f7).length(0))
                arr42_[index181_] = not((len(d_1060_v17_)) != (len(d_1057_cf24_)))
                arr43_ = self.f7
                index182_ = default__.safeIndex(848, (self.f7).length(0))
                arr43_[index182_] = (self.f7)[default__.safeIndex(848, (self.f7).length(0))]
                d_1072_v26_: C0
                nw172_ = C0()
                nw172_.ctor__(_dafny.Map({_dafny.Map({len(((d_1057_cf24_).set(default__.safeIndex((d_1059_v16_)[default__.safeIndex(232, (d_1059_v16_).length(0))], len(d_1057_cf24_)), (self).f10)).set(default__.safeIndex((d_1059_v16_)[default__.safeIndex(232, (d_1059_v16_).length(0))], len((d_1057_cf24_).set(default__.safeIndex((d_1059_v16_)[default__.safeIndex(232, (d_1059_v16_).length(0))], len(d_1057_cf24_)), (self).f10))), p0)): (d_1059_v16_)[default__.safeIndex(232, (d_1059_v16_).length(0))]}): (self).fm2((self).f10, globalState)}))
                d_1072_v26_ = nw172_
            d_1073_v27_: _dafny.MultiSet
            d_1073_v27_ = _dafny.MultiSet([(0) - ((d_1059_v16_)[default__.safeIndex(232, (d_1059_v16_).length(0))])])
            d_1074_v28_: _dafny.Set
            d_1074_v28_ = _dafny.Set({default__.safeModuloInt(-541, (self).fm4((d_1059_v16_)[default__.safeIndex(232, (d_1059_v16_).length(0))], (d_1059_v16_)[default__.safeIndex(232, (d_1059_v16_).length(0))], d_1023_v2_, (0) - ((d_1059_v16_)[default__.safeIndex(232, (d_1059_v16_).length(0))]), globalState))})
            rhs212_ = (d_1073_v27_).issubset(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f9 for d_1075_i3_ in range(default__.abs(713))])))
            rhs213_ = False
            rhs214_ = len(d_1074_v28_)
            lhs173_ = self
            lhs174_ = self
            lhs175_ = globalState
            lhs173_.f8 = rhs212_
            lhs174_.f8 = rhs213_
            lhs175_.f0 = rhs214_
        if (self).f6:
            d_1076_v29_: str
            d_1076_v29_ = _dafny.CodePoint('e')
            d_1077_v30_: _dafny.Map
            d_1077_v30_ = _dafny.Map({(self).f9: (self).f9})
            d_1078_v31_: D2
            d_1078_v31_ = D2_DC7(d_1076_v29_, d_1077_v30_, (self).f9)
            d_1079_v32_: D2
            d_1079_v32_ = D2_DC8(d_1078_v31_)
            source16_ = d_1079_v32_
            if source16_.is_DC5:
                d_1080___mcc_h9_ = source16_.cf5
                d_1081_cf5_ = d_1080___mcc_h9_
                d_1082_v33_: _dafny.Array
                nw173_ = _dafny.Array(int(0), 25)
                d_1082_v33_ = nw173_
                index183_ = default__.safeIndex(690, (d_1082_v33_).length(0))
                (d_1082_v33_)[index183_] = (self).f9
                index184_ = default__.safeIndex(837, (d_1082_v33_).length(0))
                (d_1082_v33_)[index184_] = ((self).f9) + ((self).f9)
                index185_ = default__.safeIndex(690, (d_1082_v33_).length(0))
                index186_ = default__.safeIndex(837, (d_1082_v33_).length(0))
                rhs215_ = (self).f9
                rhs216_ = ((self).f9) + (721)
                rhs217_ = (len(d_1022_v1_)) - (-353)
                rhs218_ = not((p0 if (self).f10 else (self).f10))
                lhs176_ = d_1082_v33_
                lhs177_ = default__.safeIndex(690, (d_1082_v33_).length(0))
                lhs178_ = d_1082_v33_
                lhs179_ = default__.safeIndex(837, (d_1082_v33_).length(0))
                lhs180_ = globalState
                lhs181_ = self
                lhs176_[lhs177_] = rhs215_
                lhs178_[lhs179_] = rhs216_
                lhs180_.f0 = rhs217_
                lhs181_.f8 = rhs218_
                (globalState).f0 = (self).f9
                d_1083_v34_: _dafny.Map
                d_1083_v34_ = _dafny.Map({(self).f10: (self).f10})
                index187_ = default__.safeIndex(690, (d_1082_v33_).length(0))
                (d_1082_v33_)[index187_] = ((len(d_1083_v34_)) - (301)) * ((d_1082_v33_)[default__.safeIndex(690, (d_1082_v33_).length(0))])
                d_1084_v35_: _dafny.Map
                d_1084_v35_ = _dafny.Map({(self).f9: d_1076_v29_})
                index188_ = default__.safeIndex(690, (d_1082_v33_).length(0))
                (d_1082_v33_)[index188_] = (0) - ((len(d_1084_v35_)) + ((self).f9))
            elif source16_.is_DC6:
                d_1085_v36_: T0
                nw174_ = C2()
                nw174_.ctor__((self).f5)
                d_1085_v36_ = nw174_
                d_1086_v37_: _dafny.Map
                d_1086_v37_ = _dafny.Map({d_1085_v36_: not(True)})
                d_1086_v37_ = (d_1086_v37_) | (d_1086_v37_)
                d_1087_v40_: _dafny.Map
                d_1087_v40_ = _dafny.Map({d_1077_v30_: p0})
                d_1088_v41_: C0
                nw175_ = C0()
                def iife112_():
                    def iife114_():
                        coll46_ = _dafny.Map()
                        compr_46_: int
                        for compr_46_ in _dafny.IntegerRange(50, -133):
                            d_1089_v39_: int = compr_46_
                            if ((50) <= (d_1089_v39_)) and ((d_1089_v39_) < (-133)):
                                coll46_[default__.safeDivisionInt(d_1089_v39_, (self).f9)] = self.f8
                        return _dafny.Map(coll46_)
                    coll44_ = _dafny.Map()
                    def iife113_():
                        coll45_ = _dafny.Map()
                        compr_45_: int
                        for compr_45_ in _dafny.IntegerRange(50, -133):
                            d_1089_v39_: int = compr_45_
                            if ((50) <= (d_1089_v39_)) and ((d_1089_v39_) < (-133)):
                                coll45_[default__.safeDivisionInt(d_1089_v39_, (self).f9)] = self.f8
                        return _dafny.Map(coll45_)
                    compr_44_: int
                    for compr_44_ in (iife113_()
                    ).keys.Elements:
                        d_1090_v38_: int = compr_44_
                        if (d_1090_v38_) in (iife114_()
                        ):
                            coll44_[(d_1090_v38_) + ((self).f9)] = 175
                    return _dafny.Map(coll44_)
                nw175_.ctor__((_dafny.Map({iife112_()
                : self.f8})) | (d_1087_v40_))
                d_1088_v41_ = nw175_
                d_1091_v42_: _dafny.Seq
                d_1091_v42_ = _dafny.SeqWithoutIsStrInference([len(p1), (self).f9])
                d_1092_v43_: _dafny.Array
                nw176_ = _dafny.Array(int(0), 28)
                d_1092_v43_ = nw176_
                d_1093_v44_: _dafny.Seq
                d_1093_v44_ = _dafny.SeqWithoutIsStrInference([p1, d_1091_v42_, (d_1091_v42_ if p0 else (d_1091_v42_).set(default__.safeIndex(len(d_1077_v30_), len(d_1091_v42_)), len(_dafny.Map({p0: (self).fm4(len(d_1022_v1_), (self).f9, _dafny.Map({d_1022_v1_: (self).f6}), (self).f9, globalState)})))), p1, p1])
                d_1094_v45_: D3
                d_1094_v45_ = D3_DC11((self).f9)
                d_1095_v46_: _dafny.Set
                d_1095_v46_ = _dafny.Set({d_1094_v45_})
                d_1096_v47_: _dafny.MultiSet
                d_1096_v47_ = _dafny.MultiSet([(self).f9])
                rhs219_ = (d_1093_v44_)[default__.safeIndex((self).f9, len(d_1093_v44_))]
                rhs220_ = d_1092_v43_
                rhs221_ = len(_dafny.Map({(_dafny.Set({D3_DC11((self).f9), d_1094_v45_})).issubset(d_1095_v46_): p0}))
                rhs222_ = default__.safeModuloInt((0) - ((self).f9), (d_1096_v47_).cardinality)
                lhs182_ = globalState
                lhs183_ = globalState
                d_1091_v42_ = rhs219_
                d_1092_v43_ = rhs220_
                lhs182_.f0 = rhs221_
                lhs183_.f0 = rhs222_
                d_1097_v48_: _dafny.Seq
                d_1097_v48_ = _dafny.SeqWithoutIsStrInference([d_1024_v3_, d_1024_v3_])
                d_1098_v49_: C4
                nw177_ = C4()
                nw177_.ctor__(self.f7, self.f8, p0, (_dafny.MultiSet([p0])).set(self.f8, default__.abs(len(d_1097_v48_))))
                d_1098_v49_ = nw177_
                d_1099_v50_: _dafny.Set
                d_1099_v50_ = _dafny.Set({d_1098_v49_})
                d_1100_v51_: _dafny.Seq
                d_1100_v51_ = _dafny.SeqWithoutIsStrInference([(d_1099_v50_).issubset(_dafny.Set({d_1098_v49_, d_1098_v49_}))])
                d_1101_v52_: _dafny.Map
                d_1101_v52_ = _dafny.Map({d_1076_v29_: (self).f9})
                (self).f8 = (d_1100_v51_)[default__.safeIndex(((d_1101_v52_)[d_1076_v29_] if (d_1076_v29_) in (d_1101_v52_) else (self).f9), len(d_1100_v51_))]
            elif source16_.is_DC7:
                d_1102___mcc_h10_ = source16_.cf6
                d_1103___mcc_h11_ = source16_.cf7
                d_1104___mcc_h12_ = source16_.cf8
                d_1105_cf8_ = d_1104___mcc_h12_
                d_1106_cf7_ = d_1103___mcc_h11_
                d_1107_cf6_ = d_1102___mcc_h10_
                d_1108_v53_: _dafny.Array
                def lambda50_(d_1109_p0_):
                    def lambda51_(d_1110_i4_):
                        return _dafny.Map({d_1109_p0_: d_1109_p0_})

                    return lambda51_

                init28_ = lambda50_(p0)
                nw178_ = _dafny.Array(None, 18)
                for i0_28_ in range(nw178_.length(0)):
                    nw178_[i0_28_] = init28_(i0_28_)
                d_1108_v53_ = nw178_
                def lambda52_(d_1111_i5_):
                    return _dafny.Map({self.f8: (self).f6})

                init29_ = lambda52_
                nw179_ = _dafny.Array(None, 19)
                for i0_29_ in range(nw179_.length(0)):
                    nw179_[i0_29_] = init29_(i0_29_)
                d_1108_v53_ = nw179_
                (self).f8 = self.f8
                (self).f8 = (self).f10
                rhs223_ = ((self).f9) - (d_1105_cf8_)
                rhs224_ = (self).f9
                rhs225_ = (((d_1021_v0_)[(self).f6] if ((self).f6) in (d_1021_v0_) else (self).f9)) + ((self).f9)
                lhs184_ = globalState
                lhs185_ = globalState
                d_1105_cf8_ = rhs223_
                lhs184_.f0 = rhs224_
                lhs185_.f0 = rhs225_
            elif source16_.is_DC4:
                d_1112___mcc_h13_ = source16_.cf4
                d_1113_cf4_ = d_1112___mcc_h13_
                d_1076_v29_ = d_1076_v29_
                (globalState).f0 = (self).f9
                d_1114_v54_: _dafny.Map
                d_1114_v54_ = _dafny.Map({d_1077_v30_: self.f8})
                d_1115_v55_: C0
                nw180_ = C0()
                nw180_.ctor__(d_1114_v54_)
                d_1115_v55_ = nw180_
                d_1116_v56_: _dafny.Array
                def lambda53_(d_1117_p0_, d_1118_v1_):
                    def lambda54_(d_1119_i6_):
                        return (d_1118_v1_ if d_1117_p0_ else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "esgdfxat")))

                    return lambda54_

                init30_ = lambda53_(p0, d_1022_v1_)
                nw181_ = _dafny.Array(None, 5)
                for i0_30_ in range(nw181_.length(0)):
                    nw181_[i0_30_] = init30_(i0_30_)
                d_1116_v56_ = nw181_
                index189_ = default__.safeIndex(733, (d_1116_v56_).length(0))
                (d_1116_v56_)[index189_] = d_1022_v1_
                index190_ = default__.safeIndex(733, (d_1116_v56_).length(0))
                (d_1116_v56_)[index190_] = (d_1022_v1_) + (d_1022_v1_)
            elif True:
                d_1120___mcc_h14_ = source16_.cf9
                d_1121_cf9_ = d_1120___mcc_h14_
                (globalState).f0 = ((self).f9) + (len(p1))
                d_1122_v57_: D5
                d_1122_v57_ = D5_DC15(d_1022_v1_)
                d_1123_v58_: _dafny.Map
                d_1123_v58_ = _dafny.Map({d_1122_v57_: d_1022_v1_})
                d_1124_v59_: C1
                nw182_ = C1()
                nw182_.ctor__(p0, ((d_1123_v58_)[d_1122_v57_] if (d_1122_v57_) in (d_1123_v58_) else d_1022_v1_), (self).f6, (_dafny.MultiSet([(self).f10])) | ((self).f5))
                d_1124_v59_ = nw182_
                d_1125_v60_: _dafny.Seq
                d_1125_v60_ = _dafny.SeqWithoutIsStrInference([(self).f10, False])
                (self).f8 = (d_1125_v60_)[default__.safeIndex((0) - (((self).f9) - ((self).f9)), len(d_1125_v60_))]
                (globalState).f0 = (0) - ((self).f9)
            d_1126_v61_: _dafny.Set
            d_1126_v61_ = _dafny.Set({d_1024_v3_})
            arr44_ = self.f7
            index191_ = default__.safeIndex(954, (self.f7).length(0))
            arr44_[index191_] = p0
            arr45_ = self.f7
            index192_ = default__.safeIndex(954, (self.f7).length(0))
            rhs226_ = d_1126_v61_
            rhs227_ = (self).f9
            rhs228_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uqurwber"))) <= (default__.fm17(globalState))
            rhs229_ = not(self.f8)
            rhs230_ = (self).f9
            lhs186_ = globalState
            lhs187_ = self.f7
            lhs188_ = default__.safeIndex(954, (self.f7).length(0))
            lhs189_ = self
            lhs190_ = globalState
            d_1126_v61_ = rhs226_
            lhs186_.f0 = rhs227_
            lhs187_[lhs188_] = rhs228_
            lhs189_.f8 = rhs229_
            lhs190_.f0 = rhs230_
            if (self.f7)[default__.safeIndex(954, (self.f7).length(0))]:
                d_1022_v1_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ism"))) + (d_1022_v1_)
                d_1127_v62_: _dafny.Array
                def lambda55_(d_1128_i7_):
                    return D1_DC2((self.f7)[default__.safeIndex(954, (self.f7).length(0))])

                init31_ = lambda55_
                nw183_ = _dafny.Array(None, 27)
                for i0_31_ in range(nw183_.length(0)):
                    nw183_[i0_31_] = init31_(i0_31_)
                d_1127_v62_ = nw183_
                d_1129_v63_: D1
                d_1129_v63_ = D1_DC2((self).f10)
                index193_ = default__.safeIndex(32, (d_1127_v62_).length(0))
                (d_1127_v62_)[index193_] = d_1129_v63_
                d_1130_v64_: _dafny.Array
                def lambda56_(d_1131_i8_):
                    return (d_1131_i8_) * ((0) - ((self).f9))

                init32_ = lambda56_
                nw184_ = _dafny.Array(None, 15)
                for i0_32_ in range(nw184_.length(0)):
                    nw184_[i0_32_] = init32_(i0_32_)
                d_1130_v64_ = nw184_
                index194_ = default__.safeIndex(476, (d_1130_v64_).length(0))
                (d_1130_v64_)[index194_] = (self).f9
                arr46_ = self.f7
                index195_ = default__.safeIndex(954, (self.f7).length(0))
                index196_ = default__.safeIndex(32, (d_1127_v62_).length(0))
                arr47_ = self.f7
                index197_ = default__.safeIndex(954, (self.f7).length(0))
                index198_ = default__.safeIndex(476, (d_1130_v64_).length(0))
                rhs231_ = (self.f7)[default__.safeIndex(954, (self.f7).length(0))]
                rhs232_ = ((self).f9) + (((self).f9) - ((self).f9))
                rhs233_ = D1_DC2(False)
                rhs234_ = not ((d_1022_v1_) <= (d_1022_v1_)) or (p0)
                rhs235_ = (self).f9
                lhs191_ = self.f7
                lhs192_ = default__.safeIndex(954, (self.f7).length(0))
                lhs193_ = globalState
                lhs194_ = d_1127_v62_
                lhs195_ = default__.safeIndex(32, (d_1127_v62_).length(0))
                lhs196_ = self.f7
                lhs197_ = default__.safeIndex(954, (self.f7).length(0))
                lhs198_ = d_1130_v64_
                lhs199_ = default__.safeIndex(476, (d_1130_v64_).length(0))
                lhs191_[lhs192_] = rhs231_
                lhs193_.f0 = rhs232_
                lhs194_[lhs195_] = rhs233_
                lhs196_[lhs197_] = rhs234_
                lhs198_[lhs199_] = rhs235_
                index199_ = default__.safeIndex(476, (d_1130_v64_).length(0))
                (d_1130_v64_)[index199_] = default__.safeModuloInt(default__.safeDivisionInt((self).f9, (self).f9), (0) - ((self).f9))
                index200_ = default__.safeIndex(476, (d_1130_v64_).length(0))
                (d_1130_v64_)[index200_] = default__.safeDivisionInt(((self).f9 if (self).f10 else 569), ((d_1077_v30_)[(self).f9] if ((self).f9) in (d_1077_v30_) else len(_dafny.Set({d_1024_v3_}))))
                d_1132_v65_: _dafny.Array
                def lambda57_(d_1133_i9_):
                    return (self).f6

                init33_ = lambda57_
                nw185_ = _dafny.Array(None, 20)
                for i0_33_ in range(nw185_.length(0)):
                    nw185_[i0_33_] = init33_(i0_33_)
                d_1132_v65_ = nw185_
                d_1134_v66_: _dafny.Map
                d_1134_v66_ = _dafny.Map({(self).f5: d_1132_v65_})
                (globalState).f0 = (self).fm4((self).f9, len(d_1134_v66_), _dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_1135_i10_ in range(default__.abs(405))]): (self).f10}), len(_dafny.Set({p0, (self).f6, p0})), globalState)
            elif True:
                (globalState).f0 = ((self).f9) - (((self).f9) + ((self).f9))
                (self).f8 = p0
                d_1136_v67_: _dafny.Array
                nw186_ = _dafny.Array(_dafny.Map({}), 7)
                d_1136_v67_ = nw186_
                d_1137_v68_: bool
                d_1138_v69_: int
                d_1139_v70_: bool
                out10_: bool
                out11_: int
                out12_: bool
                out10_, out11_, out12_ = (self).m5((self).f9, _dafny.MultiSet([(self).f10]), d_1136_v67_, p0, globalState)
                d_1137_v68_ = out10_
                d_1138_v69_ = out11_
                d_1139_v70_ = out12_
                d_1140_v71_: _dafny.Array
                def lambda58_(d_1141_i11_):
                    return (self).f10

                init34_ = lambda58_
                nw187_ = _dafny.Array(None, 20)
                for i0_34_ in range(nw187_.length(0)):
                    nw187_[i0_34_] = init34_(i0_34_)
                d_1140_v71_ = nw187_
                d_1140_v71_ = d_1140_v71_
                d_1142_v72_: C3
                nw188_ = C3()
                nw188_.ctor__(d_1139_v70_, (self).f5)
                d_1142_v72_ = nw188_
            d_1143_v73_: _dafny.Array
            nw189_ = _dafny.Array(False, 11)
            d_1143_v73_ = nw189_
            d_1144_v74_: _dafny.Map
            d_1144_v74_ = _dafny.Map({d_1143_v73_: (self).f10})
            d_1145_v75_: D9
            d_1145_v75_ = D9_DC27(d_1144_v74_, self.f8)
            d_1146_v76_: _dafny.Seq
            d_1146_v76_ = _dafny.SeqWithoutIsStrInference([True])
            d_1147_v77_: D8
            d_1147_v77_ = D8_DC24(((d_1021_v0_)[(self).f6] if ((self).f6) in (d_1021_v0_) else len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bwkuqlte")))), (d_1145_v75_).cf43, d_1146_v76_, ((self).f9 if (self).f6 else (self).f9), (self).f9)
            source17_ = d_1147_v77_
            if source17_.is_DC24:
                d_1148___mcc_h15_ = source17_.cf34
                d_1149___mcc_h16_ = source17_.cf35
                d_1150___mcc_h17_ = source17_.cf36
                d_1151___mcc_h18_ = source17_.cf37
                d_1152___mcc_h19_ = source17_.cf38
                d_1153_cf38_ = d_1152___mcc_h19_
                d_1154_cf37_ = d_1151___mcc_h18_
                d_1155_cf36_ = d_1150___mcc_h17_
                d_1156_cf35_ = d_1149___mcc_h16_
                d_1157_cf34_ = d_1148___mcc_h15_
                d_1158_v79_: _dafny.Map
                d_1158_v79_ = _dafny.Map({d_1157_cf34_: d_1156_cf35_})
                def iife115_():
                    coll47_ = _dafny.Map()
                    compr_47_: int
                    for compr_47_ in (d_1158_v79_).keys.Elements:
                        d_1159_v78_: int = compr_47_
                        if (d_1159_v78_) in (d_1158_v79_):
                            coll47_[(d_1159_v78_) - ((D0_DC0(d_1157_cf34_)).cf0)] = d_1153_cf38_
                    return _dafny.Map(coll47_)
                d_1076_v29_ = default__.fm16(p0, iife115_()
                , d_1022_v1_, globalState)
                d_1146_v76_ = ((d_1146_v76_) + (d_1155_cf36_)) + ((d_1155_cf36_) + (d_1146_v76_))
                d_1160_v80_: D9
                d_1160_v80_ = D9_DC27(d_1144_v74_, self.f8)
                d_1161_v81_: D9
                d_1161_v81_ = D9_DC28(d_1160_v80_)
                d_1162_v82_: C4
                nw190_ = C4()
                nw190_.ctor__(d_1143_v73_, d_1156_cf35_, False, (self).f5)
                d_1162_v82_ = nw190_
                d_1163_v83_: _dafny.Map
                d_1163_v83_ = _dafny.Map({d_1161_v81_: d_1162_v82_})
                d_1164_v84_: _dafny.Map
                d_1164_v84_ = _dafny.Map({(self).f10: d_1163_v83_})
                d_1165_v85_: C2
                nw191_ = C2()
                nw191_.ctor__((self).f5)
                d_1165_v85_ = nw191_
                d_1166_v86_: _dafny.Set
                d_1166_v86_ = _dafny.Set({(((d_1164_v84_)[(self).f10] if ((self).f10) in (d_1164_v84_) else d_1163_v83_)).set(D9_DC28(D9_DC26(d_1165_v85_)), d_1162_v82_), d_1163_v83_})
                d_1167_v87_: _dafny.Seq
                d_1167_v87_ = _dafny.SeqWithoutIsStrInference([d_1166_v86_])
                d_1168_v88_: _dafny.MultiSet
                d_1168_v88_ = _dafny.MultiSet([d_1153_cf38_])
                (globalState).f0 = (0) - (len((d_1167_v87_)[default__.safeIndex(((d_1168_v88_).intersection(_dafny.MultiSet(p1))).cardinality, len(d_1167_v87_))]))
                (globalState).f0 = (self).f9
            elif source17_.is_DC25:
                d_1169___mcc_h20_ = source17_.cf39
                d_1170___mcc_h21_ = source17_.cf40
                d_1171_cf40_ = d_1170___mcc_h21_
                d_1172_cf39_ = d_1169___mcc_h20_
                d_1077_v30_ = (d_1077_v30_) | (d_1077_v30_)
                d_1022_v1_ = d_1022_v1_
                d_1173_v89_: _dafny.Seq
                d_1173_v89_ = _dafny.SeqWithoutIsStrInference([d_1146_v76_])
                d_1174_v90_: _dafny.Map
                d_1174_v90_ = _dafny.Map({p0: d_1173_v89_})
                d_1174_v90_ = (d_1174_v90_).set(((self).fm2(d_1172_cf39_, globalState)) == ((self).f10), _dafny.SeqWithoutIsStrInference([d_1146_v76_ for d_1175_i12_ in range(default__.abs(810))]))
                (globalState).f0 = (self).f9
            elif True:
                d_1176___mcc_h22_ = source17_.cf33
                d_1177_cf33_ = d_1176___mcc_h22_
                (self).f8 = (self.f7)[default__.safeIndex(954, (self.f7).length(0))]
                arr48_ = self.f7
                index201_ = default__.safeIndex(954, (self.f7).length(0))
                rhs236_ = not(self.f8)
                rhs237_ = d_1146_v76_
                lhs200_ = self.f7
                lhs201_ = default__.safeIndex(954, (self.f7).length(0))
                lhs200_[lhs201_] = rhs236_
                d_1146_v76_ = rhs237_
                d_1178_v91_: _dafny.Array
                nw192_ = _dafny.Array(_dafny.Map({}), 27)
                d_1178_v91_ = nw192_
                d_1179_v92_: bool
                d_1180_v93_: int
                d_1181_v94_: bool
                out13_: bool
                out14_: int
                out15_: bool
                out13_, out14_, out15_ = (self).m5(default__.safeModuloInt((self).f9, (self).f9), _dafny.MultiSet([True, not((self).f10), not(p0)]), d_1178_v91_, p0, globalState)
                d_1179_v92_ = out13_
                d_1180_v93_ = out14_
                d_1181_v94_ = out15_
                d_1182_v95_: _dafny.Seq
                d_1182_v95_ = _dafny.SeqWithoutIsStrInference([d_1146_v76_])
                d_1183_v96_: _dafny.Map
                d_1183_v96_ = _dafny.Map({(self).f9: _dafny.SeqWithoutIsStrInference([len(_dafny.Set({d_1180_v93_})) for d_1184_i13_ in range(default__.abs(669))])})
                d_1185_v97_: _dafny.Map
                d_1185_v97_ = _dafny.Map({(self).f6: (self.f7)[default__.safeIndex(954, (self.f7).length(0))]})
                d_1186_v98_: _dafny.Array
                nw193_ = _dafny.Array(None, 24)
                nw193_[int(0)] = (566) * ((self).f9)
                nw193_[int(1)] = len(d_1182_v95_)
                nw193_[int(2)] = d_1180_v93_
                nw193_[int(3)] = (self).f9
                nw193_[int(4)] = (self).f9
                nw193_[int(5)] = (d_1180_v93_) - ((self).f9)
                nw193_[int(6)] = (d_1180_v93_ if p0 else d_1180_v93_)
                nw193_[int(7)] = (0) - (d_1180_v93_)
                nw193_[int(8)] = d_1180_v93_
                nw193_[int(9)] = -602
                nw193_[int(10)] = default__.safeDivisionInt(d_1180_v93_, d_1180_v93_)
                nw193_[int(11)] = (self).f9
                nw193_[int(12)] = (0) - (((self).f9) * ((self).f9))
                nw193_[int(13)] = ((self).f9) - ((self).f9)
                nw193_[int(14)] = len(d_1183_v96_)
                nw193_[int(15)] = default__.safeModuloInt(d_1180_v93_, -348)
                nw193_[int(16)] = ((self).f9 if False else (self).f9)
                nw193_[int(17)] = ((self).f9) * (d_1180_v93_)
                nw193_[int(18)] = (0) - ((d_1180_v93_ if (self).f6 else -714))
                nw193_[int(19)] = len(d_1021_v0_)
                nw193_[int(20)] = d_1180_v93_
                nw193_[int(21)] = d_1180_v93_
                nw193_[int(22)] = (self).f9
                nw193_[int(23)] = (self).fm4((self).f9, (self).f9, d_1023_v2_, len(d_1185_v97_), globalState)
                d_1186_v98_ = nw193_
                d_1187_v99_: _dafny.Seq
                d_1187_v99_ = _dafny.SeqWithoutIsStrInference([d_1186_v98_, (d_1186_v98_ if (self).f6 else d_1186_v98_), d_1186_v98_])
                d_1186_v98_ = (d_1187_v99_)[default__.safeIndex((self).f9, len(d_1187_v99_))]
            (globalState).f0 = (self).f9
        elif True:
            (self).f7 = self.f7
            d_1188_v100_: _dafny.Map
            d_1188_v100_ = _dafny.Map({335: (self).f6})
            d_1189_v101_: _dafny.Map
            d_1189_v101_ = _dafny.Map({(self).f9: (self).f9})
            d_1190_v102_: _dafny.MultiSet
            d_1190_v102_ = _dafny.MultiSet([default__.fm27(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ulvsylkks"))), ((d_1188_v100_)[(self).f9] if ((self).f9) in (d_1188_v100_) else not(self.f8)), (self).f9, d_1188_v100_, globalState), d_1189_v101_])
            d_1191_v103_: _dafny.Array
            nw194_ = _dafny.Array(D7.default()(), 19)
            d_1191_v103_ = nw194_
            d_1192_v104_: D7
            d_1192_v104_ = D7_DC21((self).f9, (0) - ((self).f9))
            index202_ = default__.safeIndex(103, (d_1191_v103_).length(0))
            (d_1191_v103_)[index202_] = d_1192_v104_
            d_1193_v105_: _dafny.Array
            nw195_ = _dafny.Array(_dafny.Array(None, 0), 20)
            d_1193_v105_ = nw195_
            d_1194_v106_: _dafny.Array
            def lambda59_(d_1195_i14_):
                return (d_1195_i14_) * (153)

            init35_ = lambda59_
            nw196_ = _dafny.Array(None, 23)
            for i0_35_ in range(nw196_.length(0)):
                nw196_[i0_35_] = init35_(i0_35_)
            d_1194_v106_ = nw196_
            index203_ = default__.safeIndex(31, (d_1194_v106_).length(0))
            (d_1194_v106_)[index203_] = (self).f9
            d_1196_v107_: _dafny.Map
            d_1196_v107_ = _dafny.Map({(self).fm4((self).f9, len(d_1022_v1_), d_1023_v2_, (self).fm4((self).f9, 46, d_1023_v2_, (self).f9, globalState), globalState): d_1194_v106_})
            d_1197_v108_: _dafny.MultiSet
            d_1197_v108_ = _dafny.MultiSet([len(d_1196_v107_)])
            index204_ = default__.safeIndex(103, (d_1191_v103_).length(0))
            index205_ = default__.safeIndex(31, (d_1194_v106_).length(0))
            rhs238_ = (default__.fm28((self).f9, d_1189_v101_, (self).f10, (self).f9, globalState)).intersection(d_1190_v102_)
            rhs239_ = default__.fm29(d_1022_v1_, p1, d_1197_v108_, globalState)
            rhs240_ = d_1193_v105_
            rhs241_ = (self).f9
            rhs242_ = len(d_1021_v0_)
            lhs202_ = d_1191_v103_
            lhs203_ = default__.safeIndex(103, (d_1191_v103_).length(0))
            lhs204_ = d_1194_v106_
            lhs205_ = default__.safeIndex(31, (d_1194_v106_).length(0))
            lhs206_ = globalState
            d_1190_v102_ = rhs238_
            lhs202_[lhs203_] = rhs239_
            d_1193_v105_ = rhs240_
            lhs204_[lhs205_] = rhs241_
            lhs206_.f0 = rhs242_
            if not((self).f6):
                d_1198_v109_: T2
                nw197_ = C4()
                nw197_.ctor__(self.f7, (self).f10, True, (self).f5)
                d_1198_v109_ = nw197_
                d_1198_v109_ = d_1198_v109_
                index206_ = default__.safeIndex(31, (d_1194_v106_).length(0))
                (d_1194_v106_)[index206_] = -90
                index207_ = default__.safeIndex(31, (d_1194_v106_).length(0))
                (d_1194_v106_)[index207_] = default__.safeModuloInt((p1)[default__.safeIndex((self).f9, len(p1))], (self).f9)
                (self).f8 = self.f8
                (d_1198_v109_).f8 = (self).f6
            elif True:
                d_1199_v110_: _dafny.Array
                nw198_ = _dafny.Array(_dafny.Array(None, 0), 27)
                d_1199_v110_ = nw198_
                index208_ = default__.safeIndex(315, (d_1199_v110_).length(0))
                (d_1199_v110_)[index208_] = self.f7
                index209_ = default__.safeIndex(315, (d_1199_v110_).length(0))
                (d_1199_v110_)[index209_] = (self.f7 if ((self).f5).ispropersubset((self).f5) else self.f7)
                (globalState).f0 = (self).f9
                (self).f8 = p0
                (globalState).f0 = (d_1194_v106_)[default__.safeIndex(31, (d_1194_v106_).length(0))]
                d_1200_v111_: _dafny.Seq
                d_1200_v111_ = _dafny.SeqWithoutIsStrInference([p1, p1, p1, _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([False, self.f8])) for d_1201_i15_ in range(default__.abs(172))]), default__.fm8((self).f9, self.f8, d_1022_v1_, globalState)])
                d_1200_v111_ = (_dafny.SeqWithoutIsStrInference([p1 for d_1202_i16_ in range(default__.abs(-870))])) + (d_1200_v111_)
            d_1203_v112_: _dafny.Seq
            d_1203_v112_ = _dafny.SeqWithoutIsStrInference([p1, (p1) + (p1), (p1) + (p1)])
            d_1203_v112_ = d_1203_v112_
            rhs243_ = (0) - (((d_1194_v106_)[default__.safeIndex(31, (d_1194_v106_).length(0))]) * (((d_1197_v108_)[len(d_1188_v100_)] if (len(d_1188_v100_)) in (d_1197_v108_) else (self).f9)))
            rhs244_ = d_1022_v1_
            rhs245_ = d_1194_v106_
            lhs207_ = globalState
            lhs207_.f0 = rhs243_
            d_1022_v1_ = rhs244_
            d_1194_v106_ = rhs245_
        d_1204_v113_: _dafny.Map
        d_1204_v113_ = _dafny.Map({(self).f9: default__.safeModuloInt((self).f9, (self).f9)})
        d_1204_v113_ = (d_1204_v113_).set((self).f9, (self).f9)
        d_1205_v114_: D11
        d_1205_v114_ = D11_DC32(_dafny.Map({(self).f10: False}))
        d_1206_v115_: _dafny.Map
        d_1206_v115_ = _dafny.Map({(self).f6: False})
        pat_let_tv27_ = p0
        d_1207_v116_: _dafny.Array
        nw199_ = _dafny.Array(None, 29)
        nw199_[int(0)] = d_1205_v114_
        nw199_[int(1)] = D11_DC32(default__.fm9(globalState))
        nw199_[int(2)] = d_1205_v114_
        nw199_[int(3)] = d_1205_v114_
        nw199_[int(4)] = d_1205_v114_
        nw199_[int(5)] = d_1205_v114_
        nw199_[int(6)] = d_1205_v114_
        nw199_[int(7)] = D11_DC32(d_1206_v115_)
        def iife116_(_pat_let34_0):
            def iife117_(d_1208_dt__update__tmp_h2_):
                def iife118_(_pat_let35_0):
                    def iife119_(d_1209_dt__update_hcf51_h0_):
                        return D11_DC32(d_1209_dt__update_hcf51_h0_)
                    return iife119_(_pat_let35_0)
                return iife118_(_dafny.Map({pat_let_tv27_: True}))
            return iife117_(_pat_let34_0)
        nw199_[int(8)] = iife116_(d_1205_v114_)
        nw199_[int(9)] = d_1205_v114_
        nw199_[int(10)] = d_1205_v114_
        nw199_[int(11)] = default__.fm30(globalState)
        nw199_[int(12)] = d_1205_v114_
        nw199_[int(13)] = d_1205_v114_
        nw199_[int(14)] = d_1205_v114_
        nw199_[int(15)] = d_1205_v114_
        nw199_[int(16)] = d_1205_v114_
        nw199_[int(17)] = d_1205_v114_
        nw199_[int(18)] = d_1205_v114_
        nw199_[int(19)] = d_1205_v114_
        nw199_[int(20)] = d_1205_v114_
        nw199_[int(21)] = D11_DC32(d_1206_v115_)
        nw199_[int(22)] = d_1205_v114_
        nw199_[int(23)] = d_1205_v114_
        nw199_[int(24)] = d_1205_v114_
        nw199_[int(25)] = d_1205_v114_
        nw199_[int(26)] = d_1205_v114_
        nw199_[int(27)] = D11_DC32(_dafny.Map({(self).f6: self.f8}))
        nw199_[int(28)] = d_1205_v114_
        d_1207_v116_ = nw199_
        index210_ = default__.safeIndex(382, (d_1207_v116_).length(0))
        (d_1207_v116_)[index210_] = d_1205_v114_
        index211_ = default__.safeIndex(382, (d_1207_v116_).length(0))
        (d_1207_v116_)[index211_] = d_1205_v114_
        d_1210_v117_: _dafny.Array
        nw200_ = _dafny.Array(_dafny.CodePoint('D'), 5)
        d_1210_v117_ = nw200_
        d_1211_v118_: str
        d_1211_v118_ = _dafny.CodePoint('q')
        index212_ = default__.safeIndex(33, (d_1210_v117_).length(0))
        (d_1210_v117_)[index212_] = d_1211_v118_
        d_1212_v119_: _dafny.Array
        nw201_ = _dafny.Array(_dafny.Set({}), 11)
        d_1212_v119_ = nw201_
        d_1213_v120_: _dafny.Map
        d_1213_v120_ = _dafny.Map({d_1204_v113_: self.f8})
        d_1214_v121_: C0
        nw202_ = C0()
        nw202_.ctor__(d_1213_v120_)
        d_1214_v121_ = nw202_
        d_1215_v122_: _dafny.Set
        d_1215_v122_ = _dafny.Set({d_1214_v121_, d_1214_v121_, d_1214_v121_})
        index213_ = default__.safeIndex(510, (d_1212_v119_).length(0))
        (d_1212_v119_)[index213_] = d_1215_v122_
        d_1216_v123_: D2
        d_1216_v123_ = D2_DC8(D2_DC8(D2_DC7(_dafny.CodePoint('e'), d_1204_v113_, (self).f9)))
        d_1217_v124_: _dafny.Set
        d_1217_v124_ = _dafny.Set({(self).f9, (self).f9})
        d_1218_v125_: _dafny.Seq
        d_1218_v125_ = _dafny.SeqWithoutIsStrInference([(self).f5, (self).f5])
        d_1219_v126_: D10
        d_1219_v126_ = D10_DC31(d_1218_v125_, d_1211_v118_)
        pat_let_tv28_ = d_1211_v118_
        pat_let_tv29_ = d_1211_v118_
        pat_let_tv30_ = d_1211_v118_
        index214_ = default__.safeIndex(33, (d_1210_v117_).length(0))
        index215_ = default__.safeIndex(510, (d_1212_v119_).length(0))
        def lambda60_(source18_):
            if source18_.is_DC30:
                d_1220___mcc_h23_ = source18_.cf46
                d_1221___mcc_h24_ = source18_.cf47
                d_1222___mcc_h25_ = source18_.cf48
                d_1223_cf48_ = d_1222___mcc_h25_
                d_1224_cf47_ = d_1221___mcc_h24_
                d_1225_cf46_ = d_1220___mcc_h23_
                return pat_let_tv28_
            elif source18_.is_DC31:
                d_1226___mcc_h26_ = source18_.cf49
                d_1227___mcc_h27_ = source18_.cf50
                d_1228_cf50_ = d_1227___mcc_h27_
                d_1229_cf49_ = d_1226___mcc_h26_
                return pat_let_tv29_
            elif True:
                d_1230___mcc_h28_ = source18_.cf45
                d_1231_cf45_ = d_1230___mcc_h28_
                return pat_let_tv30_

        rhs246_ = lambda60_(d_1219_v126_)
        rhs247_ = d_1215_v122_
        rhs248_ = d_1216_v123_
        rhs249_ = d_1217_v124_
        rhs250_ = (d_1021_v0_) | ((_dafny.Map({(self).f10: -252}) if (self).f6 else _dafny.Map({p0: (self).f9})))
        lhs208_ = d_1210_v117_
        lhs209_ = default__.safeIndex(33, (d_1210_v117_).length(0))
        lhs210_ = d_1212_v119_
        lhs211_ = default__.safeIndex(510, (d_1212_v119_).length(0))
        lhs208_[lhs209_] = rhs246_
        lhs210_[lhs211_] = rhs247_
        d_1216_v123_ = rhs248_
        d_1217_v124_ = rhs249_
        d_1021_v0_ = rhs250_

    def m5(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: bool = False
        d_1232_v2_: _dafny.Map
        d_1232_v2_ = _dafny.Map({p0: (self).f9})
        d_1233_v3_: _dafny.Set
        d_1233_v3_ = _dafny.Set({d_1232_v2_, d_1232_v2_})
        d_1234_v4_: C0
        nw203_ = C0()
        def iife120_():
            def iife122_():
                coll50_ = _dafny.Map()
                compr_50_: _dafny.Map
                for compr_50_ in (d_1233_v3_).Elements:
                    d_1235_v1_: _dafny.Map = compr_50_
                    if (d_1235_v1_) in (d_1233_v3_):
                        coll50_[d_1235_v1_] = len(_dafny.Map({(self).f9: (self).f6}))
                return _dafny.Map(coll50_)
            coll48_ = _dafny.Map()
            def iife121_():
                coll49_ = _dafny.Map()
                compr_49_: _dafny.Map
                for compr_49_ in (d_1233_v3_).Elements:
                    d_1235_v1_: _dafny.Map = compr_49_
                    if (d_1235_v1_) in (d_1233_v3_):
                        coll49_[d_1235_v1_] = len(_dafny.Map({(self).f9: (self).f6}))
                return _dafny.Map(coll49_)
            compr_48_: _dafny.Map
            for compr_48_ in (iife121_()
            ).keys.Elements:
                d_1236_v0_: _dafny.Map = compr_48_
                if (d_1236_v0_) in (iife122_()
                ):
                    coll48_[d_1236_v0_] = (self).f6
            return _dafny.Map(coll48_)
        nw203_.ctor__(iife120_()
        )
        d_1234_v4_ = nw203_
        d_1237_v6_: D2
        d_1237_v6_ = D2_DC5((self).f6)
        d_1238_v7_: _dafny.Map
        d_1238_v7_ = _dafny.Map({d_1237_v6_: len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([p1 for d_1239_i1_ in range(default__.abs(814))]))).cardinality for d_1240_i0_ in range(default__.abs(502))]))})
        d_1241_v8_: _dafny.Map
        d_1241_v8_ = _dafny.Map({d_1238_v7_: (self).f10})
        d_1242_v9_: D4
        def iife123_():
            coll51_ = _dafny.Map()
            compr_51_: _dafny.Map
            for compr_51_ in (d_1241_v8_).keys.Elements:
                d_1243_v5_: _dafny.Map = compr_51_
                if (d_1243_v5_) in (d_1241_v8_):
                    coll51_[d_1243_v5_] = False
            return _dafny.Map(coll51_)
        d_1242_v9_ = D4_DC13(self.f8, iife123_()
)
        d_1244_v10_: T2
        nw204_ = C4()
        nw204_.ctor__(self.f7, (self).f6, (not(p3)) and ((d_1242_v9_).cf15), p1)
        d_1244_v10_ = nw204_
        d_1245_v11_: _dafny.Map
        d_1245_v11_ = _dafny.Map({p0: (self).f10})
        d_1246_v12_: _dafny.Set
        d_1246_v12_ = _dafny.Set({d_1245_v11_, d_1245_v11_, (d_1245_v11_) | (d_1245_v11_)})
        d_1247_v13_: _dafny.MultiSet
        d_1247_v13_ = _dafny.MultiSet([d_1245_v11_, d_1245_v11_])
        d_1248_v15_: _dafny.Map
        def iife124_():
            coll52_ = _dafny.Set()
            compr_52_: _dafny.Map
            for compr_52_ in (d_1247_v13_).Elements:
                d_1249_v14_: _dafny.Map = compr_52_
                if (d_1249_v14_) in (d_1247_v13_):
                    coll52_ = coll52_.union(_dafny.Set([d_1249_v14_]))
            return _dafny.Set(coll52_)
        d_1248_v15_ = _dafny.Map({((self).f9) >= (p0): iife124_()
        })
        d_1250_v16_: _dafny.MultiSet
        d_1250_v16_ = _dafny.MultiSet([len(_dafny.Map({(self).f10: (self).f10}))])
        d_1251_v17_: _dafny.Map
        d_1251_v17_ = _dafny.Map({(d_1250_v16_).cardinality: d_1232_v2_})
        rhs251_ = (d_1244_v10_ if False else d_1244_v10_)
        rhs252_ = ((d_1248_v15_)[(d_1251_v17_) == ((d_1251_v17_).set((self).f9, d_1232_v2_))] if ((d_1251_v17_) == ((d_1251_v17_).set((self).f9, d_1232_v2_))) in (d_1248_v15_) else default__.fm31(globalState))
        rhs253_ = p0
        lhs212_ = globalState
        d_1244_v10_ = rhs251_
        d_1246_v12_ = rhs252_
        lhs212_.f0 = rhs253_
        d_1252_v18_: _dafny.Map
        d_1252_v18_ = _dafny.Map({d_1244_v10_.f7: p3})
        d_1253_v19_: D9
        d_1253_v19_ = D9_DC27(d_1252_v18_, (d_1244_v10_).f6)
        if not(((D9_DC27(d_1252_v18_, (self).f10) if (self).f10 else d_1253_v19_)).cf43):
            d_1232_v2_ = (d_1232_v2_).set(p0, ((self).f9) - ((self).f9))
            d_1254_v20_: _dafny.Map
            d_1254_v20_ = _dafny.Map({(d_1244_v10_).f6: (self).f6})
            d_1255_v21_: _dafny.Map
            d_1255_v21_ = _dafny.Map({True: len(d_1254_v20_)})
            d_1256_v22_: D7
            d_1256_v22_ = D7_DC19(d_1255_v21_)
            d_1257_v23_: _dafny.Set
            d_1257_v23_ = _dafny.Set({d_1256_v22_, d_1256_v22_, d_1256_v22_, d_1256_v22_})
            (d_1244_v10_).f8 = ((d_1257_v23_) | (d_1257_v23_)).isdisjoint((d_1257_v23_) - (_dafny.Set({d_1256_v22_})))
            (globalState).f0 = p0
            d_1258_v24_: _dafny.Set
            d_1258_v24_ = _dafny.Set({self.f8, d_1244_v10_.f8, (self).f10})
            r0 = (d_1258_v24_) == (default__.fm32((d_1244_v10_).f6, (self).f10, not(True), globalState))
            d_1259_v25_: _dafny.Seq
            d_1259_v25_ = _dafny.SeqWithoutIsStrInference([d_1244_v10_.f8])
            d_1260_v26_: _dafny.Array
            nw205_ = _dafny.Array(None, 6)
            nw205_[int(0)] = d_1259_v25_
            nw205_[int(1)] = d_1259_v25_
            nw205_[int(2)] = d_1259_v25_
            nw205_[int(3)] = d_1259_v25_
            nw205_[int(4)] = (d_1259_v25_).set(default__.safeIndex((self).f9, len(d_1259_v25_)), False)
            nw205_[int(5)] = d_1259_v25_
            d_1260_v26_ = nw205_
            d_1261_v27_: _dafny.Map
            d_1261_v27_ = _dafny.Map({d_1260_v26_: (self).f9})
            d_1261_v27_ = (d_1261_v27_).set(d_1260_v26_, 854)
        elif True:
            d_1262_v28_: str
            d_1262_v28_ = _dafny.CodePoint('x')
            d_1262_v28_ = _dafny.CodePoint('d')
            r1 = (self).f9
            d_1263_v29_: _dafny.Array
            def lambda61_(d_1264_v10_, d_1265_p0_):
                def lambda62_(d_1266_i2_):
                    return default__.safeDivisionInt(d_1266_i2_, len(_dafny.Map({d_1264_v10_.f8: d_1265_p0_})))

                return lambda62_

            init36_ = lambda61_(d_1244_v10_, p0)
            nw206_ = _dafny.Array(None, 8)
            for i0_36_ in range(nw206_.length(0)):
                nw206_[i0_36_] = init36_(i0_36_)
            d_1263_v29_ = nw206_
            d_1267_v30_: _dafny.Map
            d_1267_v30_ = _dafny.Map({d_1263_v29_: (self).f10})
            arr49_ = self.f7
            index216_ = default__.safeIndex(132, (self.f7).length(0))
            arr49_[index216_] = ((d_1267_v30_)[d_1263_v29_] if (d_1263_v29_) in (d_1267_v30_) else (d_1244_v10_).f6)
            d_1268_v31_: _dafny.Seq
            d_1268_v31_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w"))
            arr50_ = self.f7
            index217_ = default__.safeIndex(132, (self.f7).length(0))
            arr50_[index217_] = (d_1262_v28_) not in ((d_1268_v31_) + (d_1268_v31_))
            (globalState).f0 = (self).f9
            (d_1244_v10_).m2((d_1268_v31_)[default__.safeIndex((self).f9, len(d_1268_v31_))], globalState)
        if True:
            d_1269_v32_: str
            d_1269_v32_ = _dafny.CodePoint('e')
            (d_1244_v10_).m2(d_1269_v32_, globalState)
            if (p1) == (_dafny.MultiSet([(d_1244_v10_).fm3(globalState)])):
                r1 = (self).f9
                r0 = (d_1244_v10_).f6
                r2 = True
                d_1270_v33_: _dafny.Seq
                d_1270_v33_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nfwac"))
                d_1271_v34_: _dafny.Array
                nw207_ = _dafny.Array(None, 12)
                nw207_[int(0)] = (d_1270_v33_) + (d_1270_v33_)
                nw207_[int(1)] = d_1270_v33_
                nw207_[int(2)] = d_1270_v33_
                nw207_[int(3)] = d_1270_v33_
                nw207_[int(4)] = d_1270_v33_
                nw207_[int(5)] = d_1270_v33_
                nw207_[int(6)] = d_1270_v33_
                nw207_[int(7)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "grcliofu"))) + (d_1270_v33_)
                nw207_[int(8)] = (d_1270_v33_).set(default__.safeIndex(p0, len(d_1270_v33_)), _dafny.CodePoint('g'))
                nw207_[int(9)] = d_1270_v33_
                nw207_[int(10)] = d_1270_v33_
                nw207_[int(11)] = _dafny.SeqWithoutIsStrInference([d_1269_v32_ for d_1272_i3_ in range(default__.abs(115))])
                d_1271_v34_ = nw207_
                index218_ = default__.safeIndex(151, (d_1271_v34_).length(0))
                (d_1271_v34_)[index218_] = d_1270_v33_
                index219_ = default__.safeIndex(151, (d_1271_v34_).length(0))
                (d_1271_v34_)[index219_] = ((d_1270_v33_) + (d_1270_v33_)) + ((d_1270_v33_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qo"))))
                d_1273_v35_: C4
                nw208_ = C4()
                nw208_.ctor__(d_1244_v10_.f7, self.f8, d_1244_v10_.f8, default__.fm14(d_1232_v2_, self.f8, len(default__.fm17(globalState)), True, globalState))
                d_1273_v35_ = nw208_
            elif True:
                d_1274_v36_: _dafny.Seq
                d_1274_v36_ = _dafny.SeqWithoutIsStrInference([(self).f10, (self).f6, p3, (d_1244_v10_).f6])
                (d_1244_v10_).m1(p0, self.f8, (d_1244_v10_.f8) not in (d_1274_v36_), globalState)
                (self).f8 = (default__.fm14(_dafny.Map({622: p0}), d_1244_v10_.f8, p0, (d_1244_v10_).f6, globalState)).issubset((self).f5)
                (globalState).f0 = (p0) - ((self).fm4(len(_dafny.SeqWithoutIsStrInference([self.f8, (d_1244_v10_).f6])), (self).f9, _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wrmdqy")): (self).f6}), p0, globalState))
                (d_1244_v10_).m2(d_1269_v32_, globalState)
                arr51_ = d_1244_v10_.f7
                index220_ = default__.safeIndex(317, (d_1244_v10_.f7).length(0))
                arr51_[index220_] = True
                d_1275_v37_: _dafny.Array
                nw209_ = _dafny.Array(int(0), 6)
                d_1275_v37_ = nw209_
                index221_ = default__.safeIndex(468, (d_1275_v37_).length(0))
                (d_1275_v37_)[index221_] = (self).f9
                index222_ = default__.safeIndex(634, (d_1275_v37_).length(0))
                (d_1275_v37_)[index222_] = (0) - ((self).f9)
                d_1276_v38_: _dafny.Seq
                d_1276_v38_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))
                d_1277_v39_: _dafny.Map
                d_1277_v39_ = _dafny.Map({d_1276_v38_: (d_1244_v10_).f6})
                d_1278_v40_: D3
                d_1278_v40_ = D3_DC9(d_1275_v37_)
                d_1279_v41_: D5
                d_1279_v41_ = D5_DC16(d_1278_v40_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kd"))), p0, (self).f9, d_1234_v4_)
                arr52_ = d_1244_v10_.f7
                index223_ = default__.safeIndex(317, (d_1244_v10_.f7).length(0))
                index224_ = default__.safeIndex(468, (d_1275_v37_).length(0))
                index225_ = default__.safeIndex(634, (d_1275_v37_).length(0))
                rhs254_ = (p1).ispropersubset((d_1244_v10_).f5)
                rhs255_ = ((d_1244_v10_).fm4(p0, (self).f9, d_1277_v39_, (self).f9, globalState)) - (((d_1279_v41_).cf22) - ((self).f9))
                rhs256_ = default__.safeDivisionInt((self).fm4((self).f9, (0) - (p0), d_1277_v39_, (self).f9, globalState), (self).f9)
                lhs213_ = d_1244_v10_.f7
                lhs214_ = default__.safeIndex(317, (d_1244_v10_.f7).length(0))
                lhs215_ = d_1275_v37_
                lhs216_ = default__.safeIndex(468, (d_1275_v37_).length(0))
                lhs217_ = d_1275_v37_
                lhs218_ = default__.safeIndex(634, (d_1275_v37_).length(0))
                lhs213_[lhs214_] = rhs254_
                lhs215_[lhs216_] = rhs255_
                lhs217_[lhs218_] = rhs256_
            d_1280_v42_: _dafny.Array
            def lambda63_(d_1281_i4_):
                return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ljs"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ow")))

            init37_ = lambda63_
            nw210_ = _dafny.Array(None, 28)
            for i0_37_ in range(nw210_.length(0)):
                nw210_[i0_37_] = init37_(i0_37_)
            d_1280_v42_ = nw210_
            d_1282_v43_: _dafny.Seq
            d_1282_v43_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qwwlgc"))
            index226_ = default__.safeIndex(563, (d_1280_v42_).length(0))
            (d_1280_v42_)[index226_] = d_1282_v43_
            index227_ = default__.safeIndex(563, (d_1280_v42_).length(0))
            (d_1280_v42_)[index227_] = (d_1282_v43_ if (self).f6 else d_1282_v43_)
            d_1283_v44_: _dafny.Array
            nw211_ = _dafny.Array(_dafny.Seq({}), 9)
            d_1283_v44_ = nw211_
            d_1284_v45_: _dafny.Seq
            d_1284_v45_ = _dafny.SeqWithoutIsStrInference([len(d_1282_v43_), 776, p0, -193, 933])
            index228_ = default__.safeIndex(592, (d_1283_v44_).length(0))
            (d_1283_v44_)[index228_] = (d_1284_v45_) + (d_1284_v45_)
            d_1285_v46_: _dafny.Seq
            d_1285_v46_ = _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference([len(d_1245_v11_), p0])) + (d_1284_v45_)])
            d_1286_v47_: _dafny.Array
            nw212_ = _dafny.Array(None, 8)
            nw212_[int(0)] = 929
            nw212_[int(1)] = p0
            nw212_[int(2)] = p0
            nw212_[int(3)] = p0
            nw212_[int(4)] = p0
            nw212_[int(5)] = (self).f9
            nw212_[int(6)] = p0
            nw212_[int(7)] = len(d_1282_v43_)
            d_1286_v47_ = nw212_
            d_1287_v48_: _dafny.Seq
            d_1287_v48_ = _dafny.SeqWithoutIsStrInference([d_1286_v47_, d_1286_v47_])
            index229_ = default__.safeIndex(592, (d_1283_v44_).length(0))
            (d_1283_v44_)[index229_] = ((d_1285_v46_)[default__.safeIndex(p0, len(d_1285_v46_))]).set(default__.safeIndex((len(d_1287_v48_)) + ((0) - ((self).f9)), len((d_1285_v46_)[default__.safeIndex(p0, len(d_1285_v46_))])), default__.safeModuloInt(p0, -440))
            d_1288_v49_: _dafny.MultiSet
            d_1288_v49_ = _dafny.MultiSet([d_1269_v32_])
            d_1289_v50_: _dafny.Seq
            d_1289_v50_ = _dafny.SeqWithoutIsStrInference([d_1288_v49_, d_1288_v49_])
            d_1290_v51_: D17
            d_1290_v51_ = D17_DC46(d_1288_v49_)
            d_1291_v52_: _dafny.Array
            nw213_ = _dafny.Array(None, 23)
            nw213_[int(0)] = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_1269_v32_])])
            nw213_[int(1)] = d_1289_v50_
            nw213_[int(2)] = d_1289_v50_
            nw213_[int(3)] = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet((d_1280_v42_)[default__.safeIndex(563, (d_1280_v42_).length(0))]) for d_1292_i5_ in range(default__.abs(395))])
            nw213_[int(4)] = _dafny.SeqWithoutIsStrInference([d_1288_v49_, d_1288_v49_])
            nw213_[int(5)] = (d_1289_v50_).set(default__.safeIndex((_dafny.MultiSet([p3, d_1244_v10_.f8])).cardinality, len(d_1289_v50_)), d_1288_v49_)
            nw213_[int(6)] = d_1289_v50_
            nw213_[int(7)] = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_1269_v32_, d_1269_v32_]), d_1288_v49_, d_1288_v49_, d_1288_v49_, d_1288_v49_])
            nw213_[int(8)] = (d_1289_v50_) + (d_1289_v50_)
            nw213_[int(9)] = d_1289_v50_
            nw213_[int(10)] = (d_1289_v50_) + (_dafny.SeqWithoutIsStrInference([default__.fm33(False, (self).f9, globalState), d_1288_v49_]))
            nw213_[int(11)] = d_1289_v50_
            nw213_[int(12)] = (default__.fm34(globalState)) + (_dafny.SeqWithoutIsStrInference([d_1288_v49_]))
            nw213_[int(13)] = (_dafny.SeqWithoutIsStrInference([d_1288_v49_, d_1288_v49_])) + ((d_1289_v50_).set(default__.safeIndex((self).f9, len(d_1289_v50_)), _dafny.MultiSet([d_1269_v32_])))
            nw213_[int(14)] = (_dafny.SeqWithoutIsStrInference([d_1288_v49_]) if (self).f10 else d_1289_v50_)
            nw213_[int(15)] = d_1289_v50_
            nw213_[int(16)] = d_1289_v50_
            nw213_[int(17)] = (default__.fm34(globalState)) + (d_1289_v50_)
            nw213_[int(18)] = (_dafny.SeqWithoutIsStrInference([d_1288_v49_, d_1288_v49_, _dafny.MultiSet([d_1269_v32_])])) + (d_1289_v50_)
            nw213_[int(19)] = (_dafny.SeqWithoutIsStrInference([d_1288_v49_, (d_1290_v51_).cf73, _dafny.MultiSet([_dafny.CodePoint('g'), d_1269_v32_, d_1269_v32_]), d_1288_v49_, _dafny.MultiSet([d_1269_v32_])])) + (_dafny.SeqWithoutIsStrInference([d_1288_v49_ for d_1293_i6_ in range(default__.abs(130))]))
            nw213_[int(20)] = default__.fm34(globalState)
            nw213_[int(21)] = d_1289_v50_
            nw213_[int(22)] = (d_1289_v50_ if (d_1244_v10_).f6 else d_1289_v50_)
            d_1291_v52_ = nw213_
            index230_ = default__.safeIndex(413, (d_1291_v52_).length(0))
            (d_1291_v52_)[index230_] = (_dafny.SeqWithoutIsStrInference([(d_1288_v49_).set(d_1269_v32_, default__.abs((d_1284_v45_)[default__.safeIndex(-144, len(d_1284_v45_))])), d_1288_v49_, d_1288_v49_, d_1288_v49_])).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference([(d_1288_v49_).set(d_1269_v32_, default__.abs((d_1284_v45_)[default__.safeIndex(-144, len(d_1284_v45_))])), d_1288_v49_, d_1288_v49_, d_1288_v49_]))), d_1288_v49_)
            d_1294_v53_: _dafny.Seq
            d_1294_v53_ = _dafny.SeqWithoutIsStrInference([d_1289_v50_])
            index231_ = default__.safeIndex(413, (d_1291_v52_).length(0))
            rhs257_ = (self).f9
            rhs258_ = (d_1294_v53_)[default__.safeIndex((self).f9, len(d_1294_v53_))]
            rhs259_ = len(d_1284_v45_)
            lhs219_ = globalState
            lhs220_ = d_1291_v52_
            lhs221_ = default__.safeIndex(413, (d_1291_v52_).length(0))
            lhs222_ = globalState
            lhs219_.f0 = rhs257_
            lhs220_[lhs221_] = rhs258_
            lhs222_.f0 = rhs259_
        elif True:
            d_1295_v54_: _dafny.Array
            nw214_ = _dafny.Array(int(0), 2)
            d_1295_v54_ = nw214_
            index232_ = default__.safeIndex(640, (d_1295_v54_).length(0))
            (d_1295_v54_)[index232_] = p0
            index233_ = default__.safeIndex(640, (d_1295_v54_).length(0))
            rhs260_ = self.f8
            rhs261_ = not((self).f10)
            rhs262_ = p0
            lhs223_ = self
            lhs224_ = d_1295_v54_
            lhs225_ = default__.safeIndex(640, (d_1295_v54_).length(0))
            r0 = rhs260_
            lhs223_.f8 = rhs261_
            lhs224_[lhs225_] = rhs262_
            source19_ = D0_DC1()
            if source19_.is_DC1:
                (self).f8 = (d_1244_v10_).f6
                d_1296_v55_: _dafny.Seq
                d_1296_v55_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rq"))
                d_1297_v56_: _dafny.Map
                d_1297_v56_ = _dafny.Map({self.f7: d_1296_v55_})
                d_1298_v57_: C1
                nw215_ = C1()
                nw215_.ctor__(d_1244_v10_.f8, ((d_1297_v56_)[d_1244_v10_.f7] if (d_1244_v10_.f7) in (d_1297_v56_) else d_1296_v55_), self.f8, ((d_1244_v10_).f5).intersection(p1))
                d_1298_v57_ = nw215_
                d_1299_v58_: _dafny.Seq
                d_1299_v58_ = _dafny.SeqWithoutIsStrInference([not(p3)])
                index234_ = default__.safeIndex(640, (d_1295_v54_).length(0))
                (d_1295_v54_)[index234_] = default__.safeModuloInt(len(d_1299_v58_), (self).f9)
                arr53_ = d_1244_v10_.f7
                index235_ = default__.safeIndex(91, (d_1244_v10_.f7).length(0))
                arr53_[index235_] = not(not(d_1244_v10_.f8))
                d_1300_v59_: str
                d_1300_v59_ = _dafny.CodePoint('j')
                d_1301_v60_: _dafny.Seq
                d_1301_v60_ = _dafny.SeqWithoutIsStrInference([((d_1298_v57_).f13).set(default__.safeIndex((self).f9, len((d_1298_v57_).f13)), d_1300_v59_)])
                d_1302_v61_: _dafny.Seq
                d_1302_v61_ = _dafny.SeqWithoutIsStrInference([d_1245_v11_, d_1245_v11_, d_1245_v11_, d_1245_v11_])
                d_1303_v62_: _dafny.Seq
                d_1303_v62_ = _dafny.SeqWithoutIsStrInference([(self).f9])
                arr54_ = d_1244_v10_.f7
                index236_ = default__.safeIndex(91, (d_1244_v10_.f7).length(0))
                rhs263_ = len((((d_1301_v60_)[default__.safeIndex((_dafny.MultiSet((d_1302_v61_).set(default__.safeIndex((self).f9, len(d_1302_v61_)), d_1245_v11_))).cardinality, len(d_1301_v60_))]).set(default__.safeIndex((d_1295_v54_)[default__.safeIndex(640, (d_1295_v54_).length(0))], len((d_1301_v60_)[default__.safeIndex((_dafny.MultiSet((d_1302_v61_).set(default__.safeIndex((self).f9, len(d_1302_v61_)), d_1245_v11_))).cardinality, len(d_1301_v60_))])), d_1300_v59_)) + ((d_1298_v57_).f13))
                rhs264_ = p3
                rhs265_ = not((((self).f9) + ((self).f9)) != ((0) - ((0) - (len(d_1303_v62_)))))
                lhs226_ = globalState
                lhs227_ = d_1244_v10_.f7
                lhs228_ = default__.safeIndex(91, (d_1244_v10_.f7).length(0))
                lhs226_.f0 = rhs263_
                r2 = rhs264_
                lhs227_[lhs228_] = rhs265_
            elif True:
                d_1304___mcc_h0_ = source19_.cf0
                d_1305_cf0_ = d_1304___mcc_h0_
                (self).f8 = ((self).f10 if (d_1244_v10_).f6 else ((self).f6) == ((self).f10))
                d_1306_v63_: _dafny.Seq
                d_1306_v63_ = _dafny.SeqWithoutIsStrInference([d_1295_v54_])
                index237_ = default__.safeIndex(640, (d_1295_v54_).length(0))
                index238_ = default__.safeIndex(640, (d_1295_v54_).length(0))
                rhs266_ = self.f8
                rhs267_ = (self).f9
                rhs268_ = ((self).f9) - ((0) - ((d_1295_v54_)[default__.safeIndex(640, (d_1295_v54_).length(0))]))
                rhs269_ = d_1305_cf0_
                rhs270_ = (d_1306_v63_)[default__.safeIndex(default__.safeDivisionInt((d_1295_v54_)[default__.safeIndex(640, (d_1295_v54_).length(0))], (d_1295_v54_)[default__.safeIndex(640, (d_1295_v54_).length(0))]), len(d_1306_v63_))]
                lhs229_ = self
                lhs230_ = d_1295_v54_
                lhs231_ = default__.safeIndex(640, (d_1295_v54_).length(0))
                lhs232_ = d_1295_v54_
                lhs233_ = default__.safeIndex(640, (d_1295_v54_).length(0))
                lhs229_.f8 = rhs266_
                lhs230_[lhs231_] = rhs267_
                r1 = rhs268_
                lhs232_[lhs233_] = rhs269_
                d_1295_v54_ = rhs270_
                d_1307_v64_: D0
                d_1307_v64_ = D0_DC1()
                d_1308_v65_: _dafny.Seq
                d_1308_v65_ = _dafny.SeqWithoutIsStrInference([(d_1244_v10_).f6, (p0) >= (364), ((0) - (p0)) > (d_1305_cf0_)])
                rhs271_ = d_1307_v64_
                rhs272_ = (d_1308_v65_) + (d_1308_v65_)
                rhs273_ = (self).f9
                lhs234_ = globalState
                d_1307_v64_ = rhs271_
                d_1308_v65_ = rhs272_
                lhs234_.f0 = rhs273_
                r1 = default__.safeDivisionInt(d_1305_cf0_, p0)
            d_1309_v66_: _dafny.Seq
            d_1309_v66_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p1]), _dafny.SeqWithoutIsStrInference([(d_1244_v10_).f5])])
            d_1310_v67_: str
            d_1310_v67_ = _dafny.CodePoint('u')
            d_1311_v68_: D10
            d_1311_v68_ = D10_DC31((d_1309_v66_)[default__.safeIndex((self).f9, len(d_1309_v66_))], d_1310_v67_)
            d_1311_v68_ = d_1311_v68_
            d_1312_v69_: _dafny.Seq
            d_1312_v69_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mu"))
            d_1313_v70_: _dafny.Map
            d_1313_v70_ = _dafny.Map({(self).f10: len(d_1312_v69_)})
            d_1314_v71_: _dafny.Set
            d_1314_v71_ = _dafny.Set({(d_1244_v10_).f6})
            d_1315_v72_: T1
            nw216_ = C3()
            nw216_.ctor__((d_1244_v10_).fm1((self).fm3(globalState), len(d_1313_v70_), len(d_1314_v71_), globalState), _dafny.MultiSet([(self).f6]))
            d_1315_v72_ = nw216_
            d_1315_v72_ = d_1315_v72_
            d_1316_v73_: _dafny.Map
            d_1316_v73_ = _dafny.Map({(self).f10: (d_1234_v4_).f11})
            d_1317_v74_: C0
            nw217_ = C0()
            nw217_.ctor__(((d_1316_v73_)[(d_1244_v10_).f6] if ((d_1244_v10_).f6) in (d_1316_v73_) else _dafny.Map({d_1232_v2_: (d_1244_v10_).f6})))
            d_1317_v74_ = nw217_
        d_1318_v75_: _dafny.Seq
        d_1318_v75_ = _dafny.SeqWithoutIsStrInference([d_1244_v10_.f8])
        d_1319_v76_: D6
        d_1319_v76_ = D6_DC17(d_1318_v75_)
        d_1320_v77_: D11
        d_1320_v77_ = D11_DC33(d_1244_v10_.f8, False)
        pat_let_tv31_ = d_1244_v10_
        pat_let_tv32_ = p3
        d_1321_v78_: _dafny.Map
        def iife125_(_pat_let36_0):
            def iife126_(d_1322_dt__update__tmp_h0_):
                def iife127_(_pat_let37_0):
                    def iife128_(d_1323_dt__update_hcf24_h0_):
                        return D6_DC17(d_1323_dt__update_hcf24_h0_)
                    return iife128_(_pat_let37_0)
                return iife127_(_dafny.SeqWithoutIsStrInference([pat_let_tv31_.f8, pat_let_tv32_]))
            return iife126_(_pat_let36_0)
        d_1321_v78_ = _dafny.Map({iife125_(d_1319_v76_): d_1320_v77_})
        d_1321_v78_ = (d_1321_v78_) | (_dafny.Map({D6_DC17(d_1318_v75_): d_1320_v77_}))
        d_1324_i7_: int
        d_1324_i7_ = 0
        with _dafny.label("12"):
            while (self).f6:
                with _dafny.c_label("12"):
                    if (d_1324_i7_) >= (100):
                        raise _dafny.Break("12")
                    d_1324_i7_ = (d_1324_i7_) + (1)
                    d_1325_v79_: _dafny.Array
                    def lambda64_(d_1326_i8_):
                        return (d_1326_i8_) - ((self).f9)

                    init38_ = lambda64_
                    nw218_ = _dafny.Array(None, 15)
                    for i0_38_ in range(nw218_.length(0)):
                        nw218_[i0_38_] = init38_(i0_38_)
                    d_1325_v79_ = nw218_
                    index239_ = default__.safeIndex(424, (d_1325_v79_).length(0))
                    (d_1325_v79_)[index239_] = p0
                    index240_ = default__.safeIndex(424, (d_1325_v79_).length(0))
                    (d_1325_v79_)[index240_] = default__.safeModuloInt(p0, p0)
                    d_1327_v80_: _dafny.Seq
                    d_1327_v80_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet(d_1318_v75_)])
                    d_1328_v81_: D10
                    d_1328_v81_ = D10_DC31(d_1327_v80_, _dafny.CodePoint('q'))
                    d_1329_v82_: _dafny.Seq
                    d_1329_v82_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hm"))
                    if not(((d_1328_v81_).cf50) not in (d_1329_v82_)):
                        d_1330_v83_: _dafny.Map
                        d_1330_v83_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uexbyodl")): (self).f6})
                        r1 = (0) - ((self).fm4((self).f9, (d_1325_v79_)[default__.safeIndex(424, (d_1325_v79_).length(0))], d_1330_v83_, (0) - (len(d_1329_v82_)), globalState))
                        (d_1244_v10_).f8 = p3
                        (globalState).f0 = p0
                        arr55_ = d_1244_v10_.f7
                        index241_ = default__.safeIndex(665, (d_1244_v10_.f7).length(0))
                        arr55_[index241_] = p3
                        arr56_ = d_1244_v10_.f7
                        index242_ = default__.safeIndex(665, (d_1244_v10_.f7).length(0))
                        arr56_[index242_] = True
                        (d_1244_v10_).m2(default__.fm16((d_1244_v10_).f6, d_1232_v2_, d_1329_v82_, globalState), globalState)
                    elif True:
                        (globalState).f0 = p0
                        r2 = (self).f10
                        d_1331_v84_: _dafny.MultiSet
                        d_1331_v84_ = _dafny.MultiSet([d_1325_v79_])
                        d_1332_v85_: D18
                        d_1332_v85_ = D18_DC48(d_1331_v84_)
                        d_1331_v84_ = (d_1331_v84_).intersection((d_1331_v84_).intersection((d_1332_v85_).cf77))
                        (d_1244_v10_).f8 = not((self).f10)
                        d_1333_v86_: _dafny.Array
                        nw219_ = _dafny.Array(D4.default()(), 24)
                        d_1333_v86_ = nw219_
                        d_1334_v87_: str
                        d_1334_v87_ = _dafny.CodePoint('c')
                        d_1335_v88_: _dafny.Array
                        nw220_ = _dafny.Array(None, 22)
                        nw220_[int(0)] = d_1334_v87_
                        nw220_[int(1)] = d_1334_v87_
                        nw220_[int(2)] = d_1334_v87_
                        nw220_[int(3)] = d_1334_v87_
                        nw220_[int(4)] = d_1334_v87_
                        nw220_[int(5)] = _dafny.CodePoint('p')
                        nw220_[int(6)] = d_1334_v87_
                        nw220_[int(7)] = _dafny.CodePoint('a')
                        nw220_[int(8)] = d_1334_v87_
                        nw220_[int(9)] = d_1334_v87_
                        nw220_[int(10)] = d_1334_v87_
                        nw220_[int(11)] = (d_1328_v81_).cf50
                        nw220_[int(12)] = d_1334_v87_
                        nw220_[int(13)] = d_1334_v87_
                        nw220_[int(14)] = d_1334_v87_
                        nw220_[int(15)] = d_1334_v87_
                        nw220_[int(16)] = d_1334_v87_
                        nw220_[int(17)] = d_1334_v87_
                        nw220_[int(18)] = d_1334_v87_
                        nw220_[int(19)] = d_1334_v87_
                        nw220_[int(20)] = d_1334_v87_
                        nw220_[int(21)] = d_1334_v87_
                        d_1335_v88_ = nw220_
                        d_1336_v89_: _dafny.Array
                        nw221_ = _dafny.Array(None, 25)
                        nw221_[int(0)] = d_1335_v88_
                        nw221_[int(1)] = d_1335_v88_
                        nw221_[int(2)] = d_1335_v88_
                        nw221_[int(3)] = d_1335_v88_
                        nw221_[int(4)] = d_1335_v88_
                        nw221_[int(5)] = d_1335_v88_
                        nw221_[int(6)] = d_1335_v88_
                        nw221_[int(7)] = d_1335_v88_
                        nw221_[int(8)] = d_1335_v88_
                        nw221_[int(9)] = d_1335_v88_
                        nw221_[int(10)] = d_1335_v88_
                        nw221_[int(11)] = d_1335_v88_
                        nw221_[int(12)] = d_1335_v88_
                        nw221_[int(13)] = d_1335_v88_
                        nw221_[int(14)] = d_1335_v88_
                        nw221_[int(15)] = d_1335_v88_
                        nw221_[int(16)] = d_1335_v88_
                        nw221_[int(17)] = d_1335_v88_
                        nw221_[int(18)] = d_1335_v88_
                        nw221_[int(19)] = d_1335_v88_
                        nw221_[int(20)] = d_1335_v88_
                        nw221_[int(21)] = d_1335_v88_
                        nw221_[int(22)] = d_1335_v88_
                        nw221_[int(23)] = d_1335_v88_
                        nw221_[int(24)] = d_1335_v88_
                        d_1336_v89_ = nw221_
                        d_1337_v90_: D4
                        d_1337_v90_ = D4_DC12(d_1336_v89_)
                        d_1338_v91_: D4
                        d_1338_v91_ = D4_DC14(d_1337_v90_)
                        index243_ = default__.safeIndex(858, (d_1333_v86_).length(0))
                        (d_1333_v86_)[index243_] = d_1338_v91_
                        index244_ = default__.safeIndex(858, (d_1333_v86_).length(0))
                        rhs274_ = d_1338_v91_
                        rhs275_ = d_1232_v2_
                        rhs276_ = (len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lsl"))).set(default__.safeIndex((d_1325_v79_)[default__.safeIndex(424, (d_1325_v79_).length(0))], len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lsl")))), d_1334_v87_))) * ((d_1325_v79_)[default__.safeIndex(424, (d_1325_v79_).length(0))])
                        lhs235_ = d_1333_v86_
                        lhs236_ = default__.safeIndex(858, (d_1333_v86_).length(0))
                        lhs235_[lhs236_] = rhs274_
                        d_1232_v2_ = rhs275_
                        r1 = rhs276_
                    d_1339_v92_: C4
                    nw222_ = C4()
                    nw222_.ctor__(self.f7, ((self).f6) in (d_1318_v75_), (self).f10, (self).f5)
                    d_1339_v92_ = nw222_
                    d_1340_v93_: _dafny.Array
                    nw223_ = _dafny.Array(_dafny.Seq({}), 21)
                    d_1340_v93_ = nw223_
                    index245_ = default__.safeIndex(317, (d_1340_v93_).length(0))
                    (d_1340_v93_)[index245_] = _dafny.SeqWithoutIsStrInference([412])
                    d_1341_v94_: _dafny.Seq
                    d_1341_v94_ = _dafny.SeqWithoutIsStrInference([(self).f9])
                    d_1342_v95_: _dafny.Map
                    d_1342_v95_ = _dafny.Map({d_1329_v82_: p3})
                    index246_ = default__.safeIndex(317, (d_1340_v93_).length(0))
                    rhs277_ = ((d_1341_v94_) + (_dafny.SeqWithoutIsStrInference([(self).f9 for d_1343_i9_ in range(default__.abs(708))]))) + (d_1341_v94_)
                    rhs278_ = (752) > (default__.safeDivisionInt(194, (d_1244_v10_).fm4((d_1325_v79_)[default__.safeIndex(424, (d_1325_v79_).length(0))], (self).f9, d_1342_v95_, p0, globalState)))
                    rhs279_ = d_1329_v82_
                    rhs280_ = d_1244_v10_.f7
                    lhs237_ = d_1340_v93_
                    lhs238_ = default__.safeIndex(317, (d_1340_v93_).length(0))
                    lhs239_ = d_1244_v10_
                    lhs237_[lhs238_] = rhs277_
                    r0 = rhs278_
                    d_1329_v82_ = rhs279_
                    lhs239_.f7 = rhs280_
                    pass
            pass
        r0 = (d_1244_v10_).f6
        r1 = (self).f9
        r2 = (d_1244_v10_).f6
        return r0, r1, r2

    @property
    def f9(self):
        return self._f9
    @property
    def f10(self):
        return self._f10

class C6:
    def  __init__(self):
        self.f3: _dafny.Set = _dafny.Set({})
        self.f4: str = _dafny.CodePoint('D')
        pass

    def __dafnystr__(self) -> str:
        return "_module.C6"
    def ctor__(self, f3, f4):
        (self).f3 = f3
        (self).f4 = f4

    def fm0(self, p0, p1, p2, p3, globalState):
        if (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "adv"))) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iy"))):
            if False:
                return len(_dafny.Map({False: (0) - (len(_dafny.SeqWithoutIsStrInference([self.f4 for d_1344_i0_ in range(default__.abs(731))])))}))
            elif True:
                return 737
        elif True:
            return (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sayxmsb")))) + ((_dafny.MultiSet([True])).cardinality)

    def m0(self, p0, p1, p2, globalState):
        r0: int = int(0)
        r1: bool = False
        r2: int = int(0)
        d_1345_v0_: bool
        d_1345_v0_ = False
        d_1346_i0_: int
        d_1346_i0_ = 0
        with _dafny.label("13"):
            while d_1345_v0_:
                with _dafny.c_label("13"):
                    if (d_1346_i0_) >= (100):
                        raise _dafny.Break("13")
                    d_1346_i0_ = (d_1346_i0_) + (1)
                    d_1347_v1_: _dafny.Seq
                    d_1347_v1_ = _dafny.SeqWithoutIsStrInference([(p0) - (p2)])
                    r2 = (d_1347_v1_)[default__.safeIndex(p0, len(d_1347_v1_))]
                    d_1348_v2_: _dafny.Array
                    nw224_ = _dafny.Array(int(0), 18)
                    d_1348_v2_ = nw224_
                    index247_ = default__.safeIndex(440, (d_1348_v2_).length(0))
                    (d_1348_v2_)[index247_] = p0
                    d_1349_v3_: _dafny.Array
                    def lambda65_(d_1350_v0_):
                        def lambda66_(d_1351_i1_):
                            return d_1350_v0_

                        return lambda66_

                    init39_ = lambda65_(d_1345_v0_)
                    nw225_ = _dafny.Array(None, 5)
                    for i0_39_ in range(nw225_.length(0)):
                        nw225_[i0_39_] = init39_(i0_39_)
                    d_1349_v3_ = nw225_
                    index248_ = default__.safeIndex(684, (d_1348_v2_).length(0))
                    (d_1348_v2_)[index248_] = p2
                    d_1352_v4_: _dafny.Set
                    d_1352_v4_ = _dafny.Set({d_1349_v3_, d_1349_v3_})
                    index249_ = default__.safeIndex(440, (d_1348_v2_).length(0))
                    index250_ = default__.safeIndex(684, (d_1348_v2_).length(0))
                    rhs281_ = d_1348_v2_
                    rhs282_ = default__.safeDivisionInt(len((p1) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gutgmh")))), p0)
                    rhs283_ = d_1349_v3_
                    rhs284_ = default__.safeModuloInt(p0, len((d_1352_v4_).intersection(d_1352_v4_)))
                    lhs240_ = d_1348_v2_
                    lhs241_ = default__.safeIndex(440, (d_1348_v2_).length(0))
                    lhs242_ = d_1348_v2_
                    lhs243_ = default__.safeIndex(684, (d_1348_v2_).length(0))
                    d_1348_v2_ = rhs281_
                    lhs240_[lhs241_] = rhs282_
                    d_1349_v3_ = rhs283_
                    lhs242_[lhs243_] = rhs284_
                    d_1353_v5_: _dafny.Seq
                    d_1353_v5_ = _dafny.SeqWithoutIsStrInference([p1])
                    d_1354_v6_: C4
                    nw226_ = C4()
                    nw226_.ctor__(d_1349_v3_, default__.fm35(d_1353_v5_, d_1345_v0_, globalState), d_1345_v0_, _dafny.MultiSet([d_1345_v0_]))
                    d_1354_v6_ = nw226_
                    d_1355_v7_: _dafny.Set
                    d_1355_v7_ = _dafny.Set({d_1345_v0_, (d_1354_v6_).fm3(globalState)})
                    d_1356_v8_: _dafny.Map
                    d_1356_v8_ = _dafny.Map({d_1355_v7_: not(not(True))})
                    d_1356_v8_ = (d_1356_v8_).set(d_1355_v7_, d_1345_v0_)
                    pass
            pass
        if default__.fm35(_dafny.SeqWithoutIsStrInference([p1]), not(d_1345_v0_), globalState):
            d_1357_v9_: _dafny.Map
            d_1357_v9_ = _dafny.Map({p0: False})
            d_1358_v10_: _dafny.Map
            d_1358_v10_ = _dafny.Map({d_1357_v9_: d_1345_v0_})
            d_1359_v11_: _dafny.Seq
            d_1359_v11_ = _dafny.SeqWithoutIsStrInference([d_1358_v10_])
            d_1345_v0_ = (len((d_1359_v11_)[default__.safeIndex(len(p1), len(d_1359_v11_))])) in (self.f3)
            d_1360_v12_: _dafny.Array
            nw227_ = _dafny.Array(False, 15)
            d_1360_v12_ = nw227_
            d_1361_v13_: _dafny.Map
            d_1361_v13_ = _dafny.Map({d_1360_v12_: d_1345_v0_})
            d_1361_v13_ = d_1361_v13_
            d_1362_v14_: _dafny.Seq
            d_1362_v14_ = _dafny.SeqWithoutIsStrInference([True])
            d_1363_v15_: D18
            d_1363_v15_ = D18_DC49((d_1362_v14_).set(default__.safeIndex(p0, len(d_1362_v14_)), d_1345_v0_), 460)
            d_1364_v16_: _dafny.Map
            d_1364_v16_ = _dafny.Map({d_1363_v15_: False})
            if not(((d_1364_v16_)[d_1363_v15_] if (d_1363_v15_) in (d_1364_v16_) else d_1345_v0_)):
                (globalState).f0 = p0
                d_1365_v17_: _dafny.Array
                nw228_ = _dafny.Array(_dafny.Array(None, 0), 11)
                d_1365_v17_ = nw228_
                d_1366_v18_: _dafny.Array
                nw229_ = _dafny.Array(_dafny.Array(None, 0), 2)
                d_1366_v18_ = nw229_
                index251_ = default__.safeIndex(327, (d_1365_v17_).length(0))
                (d_1365_v17_)[index251_] = d_1366_v18_
                d_1367_v19_: _dafny.Map
                d_1367_v19_ = _dafny.Map({p0: p0})
                d_1368_v20_: _dafny.Map
                d_1368_v20_ = _dafny.Map({d_1367_v19_: d_1345_v0_})
                d_1369_v21_: C0
                nw230_ = C0()
                nw230_.ctor__(d_1368_v20_)
                d_1369_v21_ = nw230_
                d_1370_v22_: D6
                d_1370_v22_ = D6_DC18(d_1369_v21_, d_1362_v14_, p0)
                d_1371_v23_: _dafny.Seq
                d_1371_v23_ = _dafny.SeqWithoutIsStrInference([849, 655, (d_1370_v22_).cf27, p2])
                index252_ = default__.safeIndex(327, (d_1365_v17_).length(0))
                rhs285_ = len(_dafny.Set({d_1360_v12_, d_1360_v12_}))
                rhs286_ = d_1366_v18_
                rhs287_ = (len((d_1371_v23_) + (_dafny.SeqWithoutIsStrInference([p2])))) * (len(d_1362_v14_))
                lhs244_ = d_1365_v17_
                lhs245_ = default__.safeIndex(327, (d_1365_v17_).length(0))
                r2 = rhs285_
                lhs244_[lhs245_] = rhs286_
                r2 = rhs287_
                r1 = d_1345_v0_
                r1 = d_1345_v0_
                d_1372_v25_: D2
                d_1372_v25_ = D2_DC7(self.f4, d_1367_v19_, (0) - (p2))
                d_1373_v26_: _dafny.Map
                d_1373_v26_ = _dafny.Map({p2: d_1363_v15_})
                d_1374_v27_: _dafny.Seq
                def iife129_():
                    coll53_ = _dafny.Map()
                    compr_53_: int
                    for compr_53_ in ((d_1372_v25_).cf7).keys.Elements:
                        d_1375_v24_: int = compr_53_
                        if (d_1375_v24_) in ((d_1372_v25_).cf7):
                            coll53_[(d_1375_v24_) * (p0)] = D18_DC49(d_1362_v14_, len(d_1362_v14_))
                    return _dafny.Map(coll53_)
                d_1374_v27_ = _dafny.SeqWithoutIsStrInference([iife129_()
                , d_1373_v26_])
                pat_let_tv33_ = p0
                d_1376_v28_: _dafny.Map
                def iife130_(_pat_let38_0):
                    def iife131_(d_1377_dt__update__tmp_h0_):
                        def iife132_(_pat_let39_0):
                            def iife133_(d_1378_dt__update_hcf79_h0_):
                                return D18_DC49((d_1377_dt__update__tmp_h0_).cf78, d_1378_dt__update_hcf79_h0_)
                            return iife133_(_pat_let39_0)
                        return iife132_(pat_let_tv33_)
                    return iife131_(_pat_let38_0)
                d_1376_v28_ = _dafny.Map({self.f3: _dafny.Map({p0: iife130_(d_1363_v15_)})})
                (self).f4 = (p1)[default__.safeIndex((0) - (len((d_1374_v27_) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({829: d_1363_v15_}), ((d_1376_v28_)[self.f3] if (self.f3) in (d_1376_v28_) else d_1373_v26_)])))), len(p1))]
            elif True:
                d_1379_v29_: _dafny.Seq
                d_1379_v29_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sn"))
                d_1380_v30_: _dafny.Seq
                d_1380_v30_ = _dafny.SeqWithoutIsStrInference([p1, p1, p1, d_1379_v29_])
                d_1379_v29_ = ((d_1380_v30_)[default__.safeIndex((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(d_1379_v29_), p0, p2, p0]))).cardinality, len(d_1380_v30_))]) + (p1)
                (globalState).f0 = p0
                d_1381_v32_: _dafny.Map
                d_1381_v32_ = _dafny.Map({p2: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dtk")))})
                d_1382_v33_: _dafny.Map
                d_1382_v33_ = _dafny.Map({d_1381_v32_: p2})
                d_1383_v34_: C0
                nw231_ = C0()
                def iife134_():
                    coll54_ = _dafny.Map()
                    compr_54_: _dafny.Map
                    for compr_54_ in (d_1382_v33_).keys.Elements:
                        d_1384_v31_: _dafny.Map = compr_54_
                        if (d_1384_v31_) in (d_1382_v33_):
                            coll54_[d_1384_v31_] = not(d_1345_v0_)
                    return _dafny.Map(coll54_)
                nw231_.ctor__(iife134_()
                )
                d_1383_v34_ = nw231_
                index253_ = default__.safeIndex(485, (d_1360_v12_).length(0))
                (d_1360_v12_)[index253_] = d_1345_v0_
                index254_ = default__.safeIndex(485, (d_1360_v12_).length(0))
                (d_1360_v12_)[index254_] = False
                rhs288_ = (d_1360_v12_)[default__.safeIndex(485, (d_1360_v12_).length(0))]
                rhs289_ = d_1345_v0_
                rhs290_ = (d_1360_v12_)[default__.safeIndex(485, (d_1360_v12_).length(0))]
                rhs291_ = self.f4
                lhs246_ = self
                d_1345_v0_ = rhs288_
                d_1345_v0_ = rhs289_
                r1 = rhs290_
                lhs246_.f4 = rhs291_
            d_1385_v35_: _dafny.Map
            d_1385_v35_ = _dafny.Map({p2: p2})
            d_1386_v36_: _dafny.Map
            d_1386_v36_ = _dafny.Map({d_1385_v35_: d_1345_v0_})
            d_1387_v37_: C0
            nw232_ = C0()
            nw232_.ctor__(d_1386_v36_)
            d_1387_v37_ = nw232_
            d_1388_v38_: _dafny.Set
            d_1388_v38_ = _dafny.Set({d_1387_v37_})
            d_1389_v39_: _dafny.Map
            d_1389_v39_ = _dafny.Map({(p2) != (len(self.f3)): len(d_1388_v38_)})
            d_1389_v39_ = (d_1389_v39_).set((d_1387_v37_).fm10(len(p1), globalState), p0)
            d_1390_v40_: _dafny.Array
            nw233_ = _dafny.Array(None, 17)
            d_1390_v40_ = nw233_
            d_1391_v41_: D1
            d_1391_v41_ = D1_DC2(d_1345_v0_)
            d_1392_v42_: _dafny.MultiSet
            d_1392_v42_ = _dafny.MultiSet([d_1345_v0_, d_1345_v0_, False])
            d_1393_v43_: T1
            nw234_ = C3()
            nw234_.ctor__((d_1391_v41_).cf1, d_1392_v42_)
            d_1393_v43_ = nw234_
            index255_ = default__.safeIndex(758, (d_1390_v40_).length(0))
            (d_1390_v40_)[index255_] = d_1393_v43_
            index256_ = default__.safeIndex(758, (d_1390_v40_).length(0))
            (d_1390_v40_)[index256_] = d_1393_v43_
        elif True:
            d_1394_v44_: _dafny.Array
            def lambda67_(d_1395_v0_):
                def lambda68_(d_1396_i2_):
                    return d_1395_v0_

                return lambda68_

            init40_ = lambda67_(d_1345_v0_)
            nw235_ = _dafny.Array(None, 26)
            for i0_40_ in range(nw235_.length(0)):
                nw235_[i0_40_] = init40_(i0_40_)
            d_1394_v44_ = nw235_
            d_1397_v45_: _dafny.Seq
            d_1397_v45_ = _dafny.SeqWithoutIsStrInference([p1, p1, p1])
            d_1398_v46_: _dafny.Map
            d_1398_v46_ = _dafny.Map({False: default__.fm35(d_1397_v45_, d_1345_v0_, globalState)})
            d_1399_v47_: _dafny.MultiSet
            d_1399_v47_ = _dafny.MultiSet([True])
            d_1400_v48_: C5
            nw236_ = C5()
            nw236_.ctor__((0) - (p2), d_1345_v0_, d_1394_v44_, d_1345_v0_, not(((d_1398_v46_)[d_1345_v0_] if (d_1345_v0_) in (d_1398_v46_) else d_1345_v0_)), (d_1399_v47_).intersection(d_1399_v47_))
            d_1400_v48_ = nw236_
            rhs292_ = (D8_DC25(d_1345_v0_, (d_1400_v48_).f10)).cf39
            rhs293_ = p0
            d_1345_v0_ = rhs292_
            r2 = rhs293_
            if (_dafny.Set({d_1394_v44_, d_1394_v44_, d_1394_v44_})).isdisjoint(_dafny.Set({d_1394_v44_, d_1394_v44_, d_1394_v44_})):
                d_1401_v49_: _dafny.Set
                d_1401_v49_ = _dafny.Set({(d_1400_v48_).f10, d_1345_v0_})
                d_1402_v50_: D2
                d_1402_v50_ = D2_DC7(self.f4, _dafny.Map({len(d_1401_v49_): p2}), (d_1400_v48_).f9)
                d_1403_v51_: _dafny.Map
                d_1403_v51_ = _dafny.Map({False: p0})
                d_1404_v52_: _dafny.Array
                nw237_ = _dafny.Array(None, 29)
                nw237_[int(0)] = 563
                nw237_[int(1)] = 288
                nw237_[int(2)] = 712
                nw237_[int(3)] = (0) - ((d_1400_v48_).f9)
                nw237_[int(4)] = default__.safeDivisionInt(p0, 523)
                nw237_[int(5)] = default__.safeModuloInt((d_1400_v48_).f9, (d_1400_v48_).f9)
                nw237_[int(6)] = p2
                nw237_[int(7)] = (d_1402_v50_).cf8
                nw237_[int(8)] = default__.safeModuloInt(len(default__.fm17(globalState)), p2)
                nw237_[int(9)] = (0) - ((d_1400_v48_).f9)
                nw237_[int(10)] = default__.safeModuloInt((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pp")))), p2)
                nw237_[int(11)] = (p0 if d_1345_v0_ else len(d_1403_v51_))
                nw237_[int(12)] = (d_1400_v48_).f9
                nw237_[int(13)] = (d_1400_v48_).f9
                nw237_[int(14)] = (d_1400_v48_).f9
                nw237_[int(15)] = p0
                nw237_[int(16)] = -644
                nw237_[int(17)] = p2
                nw237_[int(18)] = (d_1400_v48_).f9
                nw237_[int(19)] = (d_1400_v48_).f9
                nw237_[int(20)] = -83
                nw237_[int(21)] = ((d_1400_v48_).f9) + (len(_dafny.Set({(d_1400_v48_).f9})))
                nw237_[int(22)] = p0
                nw237_[int(23)] = p2
                nw237_[int(24)] = p0
                nw237_[int(25)] = (d_1400_v48_).f9
                nw237_[int(26)] = (d_1400_v48_).f9
                nw237_[int(27)] = (p0) + (p2)
                nw237_[int(28)] = default__.safeDivisionInt(77, p0)
                d_1404_v52_ = nw237_
                index257_ = default__.safeIndex(699, (d_1404_v52_).length(0))
                (d_1404_v52_)[index257_] = p2
                d_1405_v53_: _dafny.Seq
                d_1405_v53_ = _dafny.SeqWithoutIsStrInference([(d_1400_v48_).f10])
                d_1406_v54_: _dafny.Map
                d_1406_v54_ = _dafny.Map({(d_1400_v48_).f10: (d_1405_v53_) + (d_1405_v53_)})
                index258_ = default__.safeIndex(699, (d_1404_v52_).length(0))
                rhs294_ = p0
                rhs295_ = default__.safeDivisionInt((self).fm0(d_1397_v45_, (d_1400_v48_).f9, d_1345_v0_, not(not((d_1400_v48_).f10)), globalState), 845)
                rhs296_ = (d_1406_v54_) | (d_1406_v54_)
                rhs297_ = default__.fm35((_dafny.SeqWithoutIsStrInference([p1, p1])) + (d_1397_v45_), d_1345_v0_, globalState)
                lhs247_ = globalState
                lhs248_ = d_1404_v52_
                lhs249_ = default__.safeIndex(699, (d_1404_v52_).length(0))
                lhs247_.f0 = rhs294_
                lhs248_[lhs249_] = rhs295_
                d_1406_v54_ = rhs296_
                d_1345_v0_ = rhs297_
                d_1407_v55_: _dafny.Map
                d_1407_v55_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cbep")): d_1345_v0_})
                rhs298_ = (d_1400_v48_).fm4((d_1404_v52_)[default__.safeIndex(699, (d_1404_v52_).length(0))], p2, d_1407_v55_, len(p1), globalState)
                rhs299_ = p2
                r2 = rhs298_
                r2 = rhs299_
                d_1408_v56_: _dafny.Seq
                d_1408_v56_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qwdwvuxp"))
                d_1409_v57_: T2
                nw238_ = C4()
                nw238_.ctor__(d_1394_v44_, d_1345_v0_, d_1345_v0_, d_1399_v47_)
                d_1409_v57_ = nw238_
                d_1410_v58_: _dafny.Map
                d_1410_v58_ = _dafny.Map({d_1409_v57_: D7_DC21(p2, (d_1404_v52_)[default__.safeIndex(699, (d_1404_v52_).length(0))])})
                d_1411_v59_: _dafny.Map
                d_1411_v59_ = _dafny.Map({d_1410_v58_: (d_1404_v52_)[default__.safeIndex(699, (d_1404_v52_).length(0))]})
                d_1412_v60_: _dafny.Seq
                d_1412_v60_ = _dafny.SeqWithoutIsStrInference([(d_1400_v48_).f9])
                rhs300_ = _dafny.SeqWithoutIsStrInference([self.f4 for d_1413_i3_ in range(default__.abs(-717))])
                rhs301_ = (d_1405_v53_)[default__.safeIndex((_dafny.MultiSet(d_1412_v60_)).cardinality, len(d_1405_v53_))]
                rhs302_ = (d_1411_v59_) | ((d_1411_v59_) | (d_1411_v59_))
                d_1408_v56_ = rhs300_
                d_1345_v0_ = rhs301_
                d_1411_v59_ = rhs302_
                arr57_ = d_1409_v57_.f7
                index259_ = default__.safeIndex(697, (d_1409_v57_.f7).length(0))
                arr57_[index259_] = ((d_1400_v48_).f10 if not(d_1409_v57_.f8) else (d_1409_v57_).f6)
                arr58_ = d_1409_v57_.f7
                index260_ = default__.safeIndex(697, (d_1409_v57_.f7).length(0))
                def iife135_():
                    coll55_ = _dafny.Set()
                    compr_55_: int
                    for compr_55_ in (d_1412_v60_).Elements:
                        d_1414_v61_: int = compr_55_
                        if (d_1414_v61_) in (d_1412_v60_):
                            coll55_ = coll55_.union(_dafny.Set([(d_1414_v61_) * (133)]))
                    return _dafny.Set(coll55_)
                rhs303_ = not((d_1400_v48_).f10)
                rhs304_ = (d_1404_v52_)[default__.safeIndex(699, (d_1404_v52_).length(0))]
                rhs305_ = (self.f3).issubset((_dafny.Set({p0, (d_1400_v48_).f9}) if False else iife135_()
                ))
                lhs250_ = d_1409_v57_.f7
                lhs251_ = default__.safeIndex(697, (d_1409_v57_.f7).length(0))
                lhs252_ = d_1409_v57_
                lhs250_[lhs251_] = rhs303_
                r2 = rhs304_
                lhs252_.f8 = rhs305_
                d_1415_v62_: C4
                nw239_ = C4()
                nw239_.ctor__(d_1394_v44_, (d_1409_v57_.f7)[default__.safeIndex(697, (d_1409_v57_.f7).length(0))], (d_1409_v57_.f7)[default__.safeIndex(697, (d_1409_v57_.f7).length(0))], d_1399_v47_)
                d_1415_v62_ = nw239_
                d_1416_v63_: _dafny.Seq
                d_1416_v63_ = _dafny.SeqWithoutIsStrInference([d_1415_v62_, d_1415_v62_])
                d_1417_v64_: _dafny.Map
                d_1417_v64_ = _dafny.Map({not((d_1409_v57_).f6): (d_1416_v63_)[default__.safeIndex(p2, len(d_1416_v63_))]})
                arr59_ = d_1409_v57_.f7
                index261_ = default__.safeIndex(697, (d_1409_v57_.f7).length(0))
                rhs306_ = ((d_1401_v49_) | (d_1401_v49_)).issubset(_dafny.Set({d_1345_v0_, (d_1409_v57_).f6, d_1409_v57_.f8, not((d_1400_v48_).f10), d_1345_v0_}))
                rhs307_ = d_1417_v64_
                rhs308_ = (d_1412_v60_)[default__.safeIndex(len(_dafny.Set({d_1404_v52_})), len(d_1412_v60_))]
                lhs253_ = d_1409_v57_.f7
                lhs254_ = default__.safeIndex(697, (d_1409_v57_.f7).length(0))
                lhs253_[lhs254_] = rhs306_
                d_1417_v64_ = rhs307_
                r0 = rhs308_
            elif True:
                d_1398_v46_ = (d_1398_v46_) | (d_1398_v46_)
                index262_ = default__.safeIndex(620, (d_1394_v44_).length(0))
                (d_1394_v44_)[index262_] = (d_1400_v48_).f10
                d_1418_v65_: _dafny.Map
                d_1418_v65_ = _dafny.Map({p2: (d_1400_v48_).f9})
                d_1419_v66_: D7
                d_1419_v66_ = D7_DC21((d_1400_v48_).f9, (p0) * (len(d_1418_v65_)))
                index263_ = default__.safeIndex(620, (d_1394_v44_).length(0))
                rhs309_ = (p2) >= ((d_1400_v48_).f9)
                rhs310_ = (d_1400_v48_).f9
                rhs311_ = d_1419_v66_
                rhs312_ = (0) - ((0) - (p2))
                rhs313_ = len((((p1) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jv")))) + (p1)).set(default__.safeIndex(228, len(((p1) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jv")))) + (p1))), self.f4))
                lhs255_ = d_1394_v44_
                lhs256_ = default__.safeIndex(620, (d_1394_v44_).length(0))
                lhs257_ = globalState
                lhs258_ = globalState
                lhs255_[lhs256_] = rhs309_
                r0 = rhs310_
                d_1419_v66_ = rhs311_
                lhs257_.f0 = rhs312_
                lhs258_.f0 = rhs313_
                d_1420_v67_: _dafny.Array
                def lambda69_(d_1421_v48_):
                    def lambda70_(d_1422_i4_):
                        return (d_1421_v48_).f10

                    return lambda70_

                init41_ = lambda69_(d_1400_v48_)
                nw240_ = _dafny.Array(None, 26)
                for i0_41_ in range(nw240_.length(0)):
                    nw240_[i0_41_] = init41_(i0_41_)
                d_1420_v67_ = nw240_
                d_1423_v68_: _dafny.Array
                nw241_ = _dafny.Array(_dafny.CodePoint('D'), 17)
                d_1423_v68_ = nw241_
                d_1424_v69_: _dafny.Map
                d_1424_v69_ = _dafny.Map({d_1423_v68_: d_1420_v67_})
                d_1425_v70_: _dafny.Seq
                d_1425_v70_ = _dafny.SeqWithoutIsStrInference([d_1420_v67_, ((d_1424_v69_)[d_1423_v68_] if (d_1423_v68_) in (d_1424_v69_) else d_1420_v67_), d_1420_v67_])
                d_1426_v71_: _dafny.MultiSet
                d_1426_v71_ = _dafny.MultiSet([391])
                rhs314_ = default__.safeDivisionInt((0) - (p2), (d_1400_v48_).f9)
                rhs315_ = not((default__.fm12(globalState)).isdisjoint((d_1426_v71_).set(336, default__.abs(p0))))
                rhs316_ = d_1425_v70_
                rhs317_ = ((d_1400_v48_).f10) or (not((d_1400_v48_).f10))
                rhs318_ = (d_1400_v48_).f9
                r2 = rhs314_
                d_1345_v0_ = rhs315_
                d_1425_v70_ = rhs316_
                d_1345_v0_ = rhs317_
                r2 = rhs318_
                d_1427_v72_: D1
                d_1427_v72_ = D1_DC2((d_1400_v48_).f10)
                r1 = (d_1427_v72_).cf1
                r0 = default__.safeModuloInt((p0) + (p2), len((p1) + (p1)))
            d_1428_v73_: _dafny.Seq
            d_1428_v73_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fieugd"))
            d_1428_v73_ = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_1429_i5_ in range(default__.abs(225))])) + (default__.fm11(default__.fm35(d_1397_v45_, d_1345_v0_, globalState), globalState))
            (globalState).f0 = (d_1400_v48_).f9
        d_1430_v74_: _dafny.Array
        nw242_ = _dafny.Array(False, 29)
        d_1430_v74_ = nw242_
        guard_loop_4_: int
        for guard_loop_4_ in _dafny.IntegerRange(0, (d_1430_v74_).length(0)):
            d_1431_i6_: int = guard_loop_4_
            if (True) and (((0) <= (d_1431_i6_)) and ((d_1431_i6_) < ((d_1430_v74_).length(0)))):
                (d_1430_v74_)[(d_1431_i6_)] = (D12_DC35(d_1345_v0_)).cf55
        d_1432_i7_: int
        d_1432_i7_ = 0
        with _dafny.label("14"):
            while (p2) >= (default__.safeModuloInt(p2, p0)):
                with _dafny.c_label("14"):
                    if (d_1432_i7_) >= (100):
                        raise _dafny.Break("14")
                    d_1432_i7_ = (d_1432_i7_) + (1)
                    d_1433_v75_: _dafny.Seq
                    d_1433_v75_ = _dafny.SeqWithoutIsStrInference([d_1345_v0_, d_1345_v0_, d_1345_v0_])
                    d_1434_v76_: _dafny.Set
                    d_1434_v76_ = _dafny.Set({d_1345_v0_, False})
                    d_1435_v77_: _dafny.Map
                    d_1435_v77_ = _dafny.Map({((d_1433_v75_)[default__.safeIndex(p0, len(d_1433_v75_))]) in (d_1434_v76_): d_1345_v0_})
                    r1 = not(((d_1435_v77_)[d_1345_v0_] if (d_1345_v0_) in (d_1435_v77_) else d_1345_v0_))
                    d_1436_v78_: _dafny.Map
                    d_1436_v78_ = _dafny.Map({p0: p2})
                    d_1437_v79_: _dafny.Map
                    d_1437_v79_ = _dafny.Map({d_1345_v0_: p2})
                    d_1345_v0_ = (((d_1436_v78_)[p2] if (p2) in (d_1436_v78_) else p0)) >= ((p0) - (((d_1437_v79_)[d_1345_v0_] if (d_1345_v0_) in (d_1437_v79_) else p2)))
                    d_1438_v80_: D2
                    d_1438_v80_ = D2_DC5(d_1345_v0_)
                    d_1439_v81_: D2
                    d_1439_v81_ = D2_DC8(d_1438_v80_)
                    d_1440_v82_: _dafny.Array
                    nw243_ = _dafny.Array(None, 5)
                    nw243_[int(0)] = d_1433_v75_
                    nw243_[int(1)] = _dafny.SeqWithoutIsStrInference([d_1345_v0_])
                    nw243_[int(2)] = d_1433_v75_
                    nw243_[int(3)] = d_1433_v75_
                    nw243_[int(4)] = d_1433_v75_
                    d_1440_v82_ = nw243_
                    index264_ = default__.safeIndex(126, (d_1440_v82_).length(0))
                    (d_1440_v82_)[index264_] = d_1433_v75_
                    d_1441_v83_: _dafny.Seq
                    d_1441_v83_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pgaotoy"))
                    index265_ = default__.safeIndex(126, (d_1440_v82_).length(0))
                    rhs319_ = D2_DC8(D2_DC5(d_1345_v0_))
                    rhs320_ = d_1433_v75_
                    rhs321_ = d_1441_v83_
                    lhs259_ = d_1440_v82_
                    lhs260_ = default__.safeIndex(126, (d_1440_v82_).length(0))
                    d_1439_v81_ = rhs319_
                    lhs259_[lhs260_] = rhs320_
                    d_1441_v83_ = rhs321_
                    d_1442_v84_: _dafny.MultiSet
                    d_1442_v84_ = _dafny.MultiSet([d_1345_v0_, d_1345_v0_, d_1345_v0_])
                    d_1443_v85_: C5
                    nw244_ = C5()
                    nw244_.ctor__(p2, d_1345_v0_, d_1430_v74_, d_1345_v0_, d_1345_v0_, d_1442_v84_)
                    d_1443_v85_ = nw244_
                    d_1444_v86_: _dafny.Map
                    d_1444_v86_ = _dafny.Map({d_1443_v85_: (d_1443_v85_).f9})
                    (globalState).f0 = (p0) + (((d_1444_v86_)[d_1443_v85_] if (d_1443_v85_) in (d_1444_v86_) else len(default__.fm23(p2, globalState))))
                    pass
            pass
        d_1445_v87_: C2
        nw245_ = C2()
        nw245_.ctor__(_dafny.MultiSet([d_1345_v0_]))
        d_1445_v87_ = nw245_
        d_1446_v88_: _dafny.Seq
        d_1446_v88_ = _dafny.SeqWithoutIsStrInference([d_1345_v0_])
        d_1447_v89_: _dafny.MultiSet
        d_1447_v89_ = _dafny.MultiSet([d_1345_v0_, d_1345_v0_, d_1345_v0_, False])
        if (d_1445_v87_).fm1(d_1345_v0_, (p2 if d_1345_v0_ else len(d_1446_v88_)), ((d_1447_v89_)[d_1345_v0_] if (d_1345_v0_) in (d_1447_v89_) else 979), globalState):
            d_1448_v90_: _dafny.Map
            d_1448_v90_ = _dafny.Map({d_1345_v0_: d_1345_v0_})
            d_1449_v91_: _dafny.Seq
            d_1449_v91_ = _dafny.SeqWithoutIsStrInference([(p2) * (len(d_1448_v90_)), (0) - (p0), p2])
            d_1449_v91_ = (d_1449_v91_) + (d_1449_v91_)
            d_1450_v92_: _dafny.Map
            d_1450_v92_ = _dafny.Map({not(False): p2})
            d_1451_v93_: _dafny.MultiSet
            d_1451_v93_ = _dafny.MultiSet([p2])
            d_1450_v92_ = (d_1450_v92_).set(d_1345_v0_, ((d_1451_v93_)[p0] if (p0) in (d_1451_v93_) else 369))
            d_1452_v94_: _dafny.Seq
            d_1452_v94_ = _dafny.SeqWithoutIsStrInference([p1])
            d_1453_v95_: _dafny.Map
            d_1453_v95_ = _dafny.Map({d_1430_v74_: p0})
            d_1454_v96_: _dafny.Array
            nw246_ = _dafny.Array(None, 21)
            nw246_[int(0)] = len(default__.fm18(globalState))
            nw246_[int(1)] = (d_1449_v91_)[default__.safeIndex(p0, len(d_1449_v91_))]
            nw246_[int(2)] = p0
            nw246_[int(3)] = p2
            nw246_[int(4)] = p0
            nw246_[int(5)] = p0
            nw246_[int(6)] = (p2 if False else p0)
            nw246_[int(7)] = (self).fm0(d_1452_v94_, len(self.f3), d_1345_v0_, d_1345_v0_, globalState)
            nw246_[int(8)] = p0
            nw246_[int(9)] = p2
            nw246_[int(10)] = p2
            nw246_[int(11)] = (p0 if d_1345_v0_ else p2)
            nw246_[int(12)] = p2
            nw246_[int(13)] = p0
            nw246_[int(14)] = p0
            nw246_[int(15)] = len(d_1453_v95_)
            nw246_[int(16)] = default__.safeModuloInt(p2, p2)
            nw246_[int(17)] = p2
            nw246_[int(18)] = p2
            nw246_[int(19)] = ((d_1447_v89_)[d_1345_v0_] if (d_1345_v0_) in (d_1447_v89_) else p2)
            nw246_[int(20)] = (self).fm0(_dafny.SeqWithoutIsStrInference([p1]), -966, d_1345_v0_, d_1345_v0_, globalState)
            d_1454_v96_ = nw246_
            d_1454_v96_ = d_1454_v96_
            d_1455_v97_: _dafny.Map
            d_1455_v97_ = _dafny.Map({p2: p0})
            d_1456_v98_: _dafny.Set
            d_1456_v98_ = _dafny.Set({(_dafny.Map({len(d_1448_v90_): -866})) | ((d_1455_v97_).set(p2, p0))})
            d_1457_v99_: _dafny.Seq
            d_1457_v99_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wxp"))
            d_1458_v100_: _dafny.Map
            d_1458_v100_ = _dafny.Map({d_1345_v0_: d_1455_v97_})
            rhs322_ = (d_1456_v98_) | ((d_1456_v98_ if not(default__.fm35(_dafny.SeqWithoutIsStrInference([p1 for d_1459_i8_ in range(default__.abs(132))]), not(d_1345_v0_), globalState)) else _dafny.Set({d_1455_v97_, ((d_1458_v100_)[d_1345_v0_] if (d_1345_v0_) in (d_1458_v100_) else d_1455_v97_)})))
            rhs323_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yohtecotu"))
            d_1456_v98_ = rhs322_
            d_1457_v99_ = rhs323_
            d_1460_v101_: C4
            nw247_ = C4()
            nw247_.ctor__(d_1430_v74_, False, d_1345_v0_, d_1447_v89_)
            d_1460_v101_ = nw247_
            d_1461_v102_: _dafny.Map
            d_1461_v102_ = _dafny.Map({d_1460_v101_: d_1451_v93_})
            d_1462_v103_: T2
            nw248_ = C4()
            nw248_.ctor__(d_1430_v74_, d_1345_v0_, (875) < (len(d_1461_v102_)), d_1447_v89_)
            d_1462_v103_ = nw248_
            rhs324_ = d_1462_v103_.f8
            rhs325_ = p0
            rhs326_ = d_1462_v103_
            rhs327_ = p0
            d_1345_v0_ = rhs324_
            r0 = rhs325_
            d_1462_v103_ = rhs326_
            r2 = rhs327_
        elif True:
            d_1463_v104_: _dafny.Array
            def lambda71_(d_1464_p1_):
                def lambda72_(d_1465_i9_):
                    return d_1464_p1_

                return lambda72_

            init42_ = lambda71_(p1)
            nw249_ = _dafny.Array(None, 19)
            for i0_42_ in range(nw249_.length(0)):
                nw249_[i0_42_] = init42_(i0_42_)
            d_1463_v104_ = nw249_
            index266_ = default__.safeIndex(507, (d_1463_v104_).length(0))
            (d_1463_v104_)[index266_] = p1
            d_1466_v105_: D5
            d_1466_v105_ = D5_DC15(p1)
            index267_ = default__.safeIndex(507, (d_1463_v104_).length(0))
            (d_1463_v104_)[index267_] = (((d_1466_v105_).cf18) + (_dafny.SeqWithoutIsStrInference([self.f4 for d_1467_i10_ in range(default__.abs(72))]))) + (_dafny.SeqWithoutIsStrInference([self.f4 for d_1468_i11_ in range(default__.abs(373))]))
            d_1345_v0_ = d_1345_v0_
            d_1345_v0_ = d_1345_v0_
            d_1469_v106_: _dafny.Array
            def lambda73_(d_1470_v0_):
                def lambda74_(d_1471_i12_):
                    return (d_1471_i12_) * (len(_dafny.Set({d_1470_v0_, d_1470_v0_})))

                return lambda74_

            init43_ = lambda73_(d_1345_v0_)
            nw250_ = _dafny.Array(None, 25)
            for i0_43_ in range(nw250_.length(0)):
                nw250_[i0_43_] = init43_(i0_43_)
            d_1469_v106_ = nw250_
            index268_ = default__.safeIndex(555, (d_1469_v106_).length(0))
            (d_1469_v106_)[index268_] = p2
            d_1472_v107_: _dafny.Seq
            d_1472_v107_ = _dafny.SeqWithoutIsStrInference([d_1446_v88_, _dafny.SeqWithoutIsStrInference([True]), d_1446_v88_, d_1446_v88_, d_1446_v88_])
            d_1473_v108_: C5
            nw251_ = C5()
            nw251_.ctor__(p0, d_1345_v0_, d_1430_v74_, False, d_1345_v0_, (d_1447_v89_).set(d_1345_v0_, default__.abs(p2)))
            d_1473_v108_ = nw251_
            d_1474_v109_: _dafny.Map
            d_1474_v109_ = _dafny.Map({d_1473_v108_: (d_1473_v108_).f10})
            d_1475_v110_: _dafny.MultiSet
            d_1475_v110_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([((d_1474_v109_)[d_1473_v108_] if (d_1473_v108_) in (d_1474_v109_) else not(not(not((d_1473_v108_).f10)))), d_1345_v0_]), d_1446_v88_])
            index269_ = default__.safeIndex(555, (d_1469_v106_).length(0))
            (d_1469_v106_)[index269_] = ((_dafny.MultiSet(d_1472_v107_)) - (d_1475_v110_)).cardinality
            r1 = not(not (False) or (not(d_1345_v0_)))
        d_1476_v111_: D17
        d_1476_v111_ = D17_DC47(d_1345_v0_, p0, p2)
        pat_let_tv34_ = p0
        def iife136_(_pat_let40_0):
            def iife137_(d_1477_dt__update__tmp_h1_):
                def iife138_(_pat_let41_0):
                    def iife139_(d_1478_dt__update_hcf75_h0_):
                        return D17_DC47((d_1477_dt__update__tmp_h1_).cf74, d_1478_dt__update_hcf75_h0_, (d_1477_dt__update__tmp_h1_).cf76)
                    return iife139_(_pat_let41_0)
                return iife138_(pat_let_tv34_)
            return iife137_(_pat_let40_0)
        r0 = (iife136_(d_1476_v111_)).cf75
        d_1479_v112_: _dafny.MultiSet
        d_1479_v112_ = _dafny.MultiSet([p0])
        d_1480_v113_: _dafny.Map
        d_1480_v113_ = _dafny.Map({p2: d_1345_v0_})
        d_1481_v114_: _dafny.Seq
        d_1481_v114_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uksqkb")), p1])
        r1 = not (((d_1480_v113_)[(self).fm0(d_1481_v114_, len(_dafny.SeqWithoutIsStrInference([p0, p0, p0])), d_1345_v0_, d_1345_v0_, globalState)] if ((self).fm0(d_1481_v114_, len(_dafny.SeqWithoutIsStrInference([p0, p0, p0])), d_1345_v0_, d_1345_v0_, globalState)) in (d_1480_v113_) else d_1345_v0_)) or ((_dafny.MultiSet([p2])).isdisjoint(d_1479_v112_))
        r2 = (0) - (p2)
        return r0, r1, r2

