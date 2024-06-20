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
    def fm0(p0, globalState):
        return default__.safeModuloInt(280, (-786) * ((0) - (-585)))

    @staticmethod
    def fm7(p0, p1, globalState):
        return _dafny.Set({True})

    @staticmethod
    def fm10(p0, p1, p2, globalState):
        if ((0) - (-465)) <= (701):
            return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_0_i0_ in range(default__.abs(377))])
        elif True:
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xyrvtep"))

    @staticmethod
    def fm11(p0, globalState):
        return _dafny.CodePoint('e')

    @staticmethod
    def fm12(p0, p1, globalState):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: int
            for compr_0_ in _dafny.IntegerRange(828, 917):
                d_1_v0_: int = compr_0_
                if ((828) <= (d_1_v0_)) and ((d_1_v0_) < (917)):
                    coll0_[(d_1_v0_) - (890)] = not(False)
            return _dafny.Map(coll0_)
        def iife1_():
            coll1_ = _dafny.Map()
            compr_1_: int
            for compr_1_ in (_dafny.Map({len(_dafny.Map({False: True})): True})).keys.Elements:
                d_2_v1_: int = compr_1_
                if (d_2_v1_) in (_dafny.Map({len(_dafny.Map({False: True})): True})):
                    coll1_[(d_2_v1_) * (len(_dafny.SeqWithoutIsStrInference([True])))] = False
            return _dafny.Map(coll1_)
        return not((_dafny.MultiSet([_dafny.Map({len(_dafny.Map({True: _dafny.CodePoint('l')})): not(True)})])).isdisjoint((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.Map({len(_dafny.Set({True})): True})])) if False else _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vkne"))): False}), iife0_()
        , iife1_()
        , _dafny.Map({len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iwq")))})): True}), _dafny.Map({562: False})])))))

    @staticmethod
    def fm14(p0, p1, p2, globalState):
        def iife2_():
            coll2_ = _dafny.Set()
            compr_2_: int
            for compr_2_ in _dafny.IntegerRange(518, 111):
                d_4_v0_: int = compr_2_
                if ((518) <= (d_4_v0_)) and ((d_4_v0_) < (111)):
                    coll2_ = coll2_.union(_dafny.Set([(d_4_v0_) + (828)]))
            return _dafny.Set(coll2_)
        return _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gcrwhf")) for d_3_i0_ in range(default__.abs(606))])), (83) - (len(_dafny.Map({518: _dafny.CodePoint('k')}))), ((0) - (len(_dafny.Map({641: False})))) - ((0) - (len(iife2_()
        )))])

    @staticmethod
    def fm15(p0, p1, p2, globalState):
        return D3_DC10(_dafny.MultiSet([False, True]))

    @staticmethod
    def fm16(globalState):
        return (_dafny.SeqWithoutIsStrInference([True])) + (_dafny.SeqWithoutIsStrInference([not(False)]))

    @staticmethod
    def fm17(p0, p1, p2, globalState):
        return ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([757, 86])).cardinality for d_5_i0_ in range(default__.abs(490))]))) - (_dafny.MultiSet([488]))).intersection(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_6_i2_ in range(default__.abs(-421))])) for d_7_i1_ in range(default__.abs(880))])))

    @staticmethod
    def fm18(p0, p1, globalState):
        def iife3_():
            coll3_ = _dafny.Map()
            compr_3_: int
            for compr_3_ in (_dafny.Set({len(_dafny.Map({True: False}))})).Elements:
                d_8_v0_: int = compr_3_
                if (d_8_v0_) in (_dafny.Set({len(_dafny.Map({True: False}))})):
                    coll3_[(d_8_v0_) * ((_dafny.MultiSet([len(_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wnrwxf")))}))])).cardinality)] = True
            return _dafny.Map(coll3_)
        return (_dafny.Map({914: not(True)})) | (iife3_()
        )

    @staticmethod
    def fm19(p0, p1, globalState):
        return _dafny.Map({True: True})

    @staticmethod
    def fm20(p0, p1, p2, globalState):
        return ((_dafny.Map({_dafny.MultiSet([len(_dafny.Map({(_dafny.MultiSet([False])).cardinality: (0) - (len(_dafny.SeqWithoutIsStrInference([not(True), False, True, False])))})), 671]): _dafny.SeqWithoutIsStrInference([-936 for d_9_i0_ in range(default__.abs(996))])})) | (_dafny.Map({_dafny.MultiSet([len(_dafny.Map({False: 125}))]): _dafny.SeqWithoutIsStrInference([443, 719])}))) | (_dafny.Map({_dafny.MultiSet([931]): _dafny.SeqWithoutIsStrInference([856])}))

    @staticmethod
    def fm22(p0, globalState):
        return (_dafny.SeqWithoutIsStrInference([True, False, not(True)])) + (_dafny.SeqWithoutIsStrInference([not(True)]))

    @staticmethod
    def fm23(p0, p1, p2, globalState):
        source0_ = D12_DC30(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oausrgoo")), 442, _dafny.CodePoint('l'))
        if source0_.is_DC29:
            d_10___mcc_h0_ = source0_.cf45
            d_11___mcc_h1_ = source0_.cf46
            d_12___mcc_h2_ = source0_.cf47
            d_13_cf47_ = d_12___mcc_h2_
            d_14_cf46_ = d_11___mcc_h1_
            d_15_cf45_ = d_10___mcc_h0_
            return (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_16_i0_ in range(default__.abs(494))])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_17_i1_ in range(default__.abs(539))]))
        elif source0_.is_DC30:
            d_18___mcc_h3_ = source0_.cf48
            d_19___mcc_h4_ = source0_.cf49
            d_20___mcc_h5_ = source0_.cf50
            d_21_cf50_ = d_20___mcc_h5_
            d_22_cf49_ = d_19___mcc_h4_
            d_23_cf48_ = d_18___mcc_h3_
            return (d_23_cf48_) + ((d_23_cf48_).set(default__.safeIndex(-209, len(d_23_cf48_)), _dafny.CodePoint('o')))
        elif source0_.is_DC31:
            d_24___mcc_h6_ = source0_.cf51
            d_25___mcc_h7_ = source0_.cf52
            d_26___mcc_h8_ = source0_.cf53
            d_27___mcc_h9_ = source0_.cf54
            d_28_cf54_ = d_27___mcc_h9_
            d_29_cf53_ = d_26___mcc_h8_
            d_30_cf52_ = d_25___mcc_h7_
            d_31_cf51_ = d_24___mcc_h6_
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xdvd"))
        elif True:
            d_32___mcc_h10_ = source0_.cf44
            d_33_cf44_ = d_32___mcc_h10_
            return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_34_i2_ in range(default__.abs(995))])

    @staticmethod
    def fm24(p0, p1, globalState):
        return (_dafny.MultiSet([False, not(not(False)), False, True])) - ((_dafny.MultiSet([False])).intersection(_dafny.MultiSet([True])))

    @staticmethod
    def fm25(p0, p1, p2, p3, globalState):
        return D1_DC5(len((_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ptgi"))})) | (_dafny.Set({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_35_i0_ in range(default__.abs(474))])}))), False, not((_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([False, True])])) <= (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([False, not(True), True, True, not(False)]), _dafny.MultiSet([not(True), True])]))), _dafny.Map({_dafny.MultiSet([77]): _dafny.SeqWithoutIsStrInference([len(_dafny.Set({True, True}))])}))

    @staticmethod
    def fm26(p0, p1, p2, globalState):
        return ((_dafny.Map({False: _dafny.SeqWithoutIsStrInference([len(_dafny.Map({479: False})), (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([True]))])).cardinality, 538, 754, (0) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([False, True]))])))])})) | (_dafny.Map({False: _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h')]))])) for d_36_i0_ in range(default__.abs(949))])}))) | (_dafny.Map({True: _dafny.SeqWithoutIsStrInference([230, -506])}))

    @staticmethod
    def fm27(p0, globalState):
        return _dafny.Set({(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_37_i0_ in range(default__.abs(105))])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a')])), (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_38_i1_ in range(default__.abs(486))])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_39_i2_ in range(default__.abs(660))])), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t'), _dafny.CodePoint('p'), _dafny.CodePoint('n')])})

    @staticmethod
    def fm28(p0, globalState):
        return (_dafny.Set({True})) - (_dafny.Set({not(False), True}))

    @staticmethod
    def fm29(globalState):
        return D1_DC6(538, (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([not(False), not(True), False, False, False]), _dafny.SeqWithoutIsStrInference([True])])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False])])), len(_dafny.Map({-211: True})), False)

    @staticmethod
    def fm30(p0, globalState):
        return (_dafny.Map({not(True): True})) | (_dafny.Map({True: True}))

    @staticmethod
    def fm32(p0, p1, p2, globalState):
        def iife4_():
            def iife6_():
                coll6_ = _dafny.Map()
                compr_6_: str
                for compr_6_ in (_dafny.Set({_dafny.CodePoint('g')})).Elements:
                    d_44_v1_: str = compr_6_
                    if (d_44_v1_) in (_dafny.Set({_dafny.CodePoint('g')})):
                        coll6_[d_44_v1_] = len(_dafny.SeqWithoutIsStrInference([not(True)]))
                return _dafny.Map(coll6_)
            coll4_ = _dafny.Map()
            def iife5_():
                coll5_ = _dafny.Map()
                compr_5_: str
                for compr_5_ in (_dafny.Set({_dafny.CodePoint('g')})).Elements:
                    d_44_v1_: str = compr_5_
                    if (d_44_v1_) in (_dafny.Set({_dafny.CodePoint('g')})):
                        coll5_[d_44_v1_] = len(_dafny.SeqWithoutIsStrInference([not(True)]))
                return _dafny.Map(coll5_)
            compr_4_: str
            for compr_4_ in (iife5_()
            ).keys.Elements:
                d_45_v0_: str = compr_4_
                if (d_45_v0_) in (iife6_()
                ):
                    coll4_[d_45_v0_] = not(not(True))
            return _dafny.Map(coll4_)
        return _dafny.MultiSet(((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference([507]))): (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xwenmsl")))])).cardinality}))]))]) for d_40_i0_ in range(default__.abs(974))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.Set({False}))]) for d_41_i1_ in range(default__.abs(-309))])) if True else _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([-134 for d_42_i2_ in range(default__.abs(904))]), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qqmyop"))) for d_43_i3_ in range(default__.abs(-210))]), _dafny.SeqWithoutIsStrInference([len(_dafny.Map({len((D15_DC37(iife4_()
)).cf64): True}))])])))

    @staticmethod
    def fm33(p0, p1, p2, p3, globalState):
        def iife7_():
            coll7_ = _dafny.Map()
            compr_7_: int
            for compr_7_ in (_dafny.Map({9: 301})).keys.Elements:
                d_46_v0_: int = compr_7_
                if (d_46_v0_) in (_dafny.Map({9: 301})):
                    coll7_[(d_46_v0_) - (len(_dafny.Map({(D12_DC28(_dafny.MultiSet([-568, 314]))).cf44: 341})))] = 503
            return _dafny.Map(coll7_)
        return (_dafny.Map({805: -464})) | (iife7_()
        )

    @staticmethod
    def fm34(p0, p1, p2, globalState):
        return _dafny.SeqWithoutIsStrInference([False, (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([507])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ghvpdnr"))), len(_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Map({D1_DC4(_dafny.Map({_dafny.MultiSet([True]): True})): (_dafny.MultiSet([True])).cardinality}))), 179]))]), _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: len(_dafny.Map({True: False}))})), len(_dafny.SeqWithoutIsStrInference([94, len(_dafny.Map({449: len(_dafny.SeqWithoutIsStrInference([985, 165]))})), len(_dafny.Map({not(True): True})), 425]))]))])) for d_47_i0_ in range(default__.abs(504))])), _dafny.MultiSet([len(_dafny.Map({188: -128})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ahqc")))])])) == (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([942])]))])

    @staticmethod
    def fm35(p0, p1, p2, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cxtyubjr"))

    @staticmethod
    def fm36(p0, p1, p2, globalState):
        def iife8_():
            coll8_ = _dafny.Map()
            compr_8_: str
            for compr_8_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d'), _dafny.CodePoint('x'), _dafny.CodePoint('f'), _dafny.CodePoint('q')])).Elements:
                d_49_v0_: str = compr_8_
                if (d_49_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d'), _dafny.CodePoint('x'), _dafny.CodePoint('f'), _dafny.CodePoint('q')])):
                    coll8_[d_49_v0_] = 386
            return _dafny.Map(coll8_)
        return ((_dafny.SeqWithoutIsStrInference([_dafny.Map({_dafny.CodePoint('e'): len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_48_i0_ in range(default__.abs(758))]))})])) + (_dafny.SeqWithoutIsStrInference([iife8_()
        ]))) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({_dafny.CodePoint('p'): 538}), _dafny.Map({(D4_DC13(_dafny.CodePoint('e'))).cf24: (_dafny.MultiSet([762, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "quneatwcc")))])).cardinality})]))

    @staticmethod
    def fm37(p0, p1, p2, globalState):
        source1_ = D15_DC38()
        if source1_.is_DC38:
            return D7_DC18(_dafny.SeqWithoutIsStrInference([_dafny.Map({_dafny.CodePoint('k'): -356})]))
        elif True:
            d_50___mcc_h0_ = source1_.cf64
            d_51_cf64_ = d_50___mcc_h0_
            return D7_DC18(_dafny.SeqWithoutIsStrInference([_dafny.Map({_dafny.CodePoint('e'): 217})]))

    @staticmethod
    def fm38(p0, p1, globalState):
        return ((_dafny.Set({True})).intersection(_dafny.Set({not(False)}))) - (_dafny.Set({False, False, False, False}))

    @staticmethod
    def fm39(globalState):
        def iife9_():
            coll9_ = _dafny.Set()
            compr_9_: int
            for compr_9_ in _dafny.IntegerRange(562, 427):
                d_52_v0_: int = compr_9_
                if ((562) <= (d_52_v0_)) and ((d_52_v0_) < (427)):
                    coll9_ = coll9_.union(_dafny.Set([default__.safeDivisionInt(d_52_v0_, len(_dafny.SeqWithoutIsStrInference([357, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True]))).cardinality])))]))
            return _dafny.Set(coll9_)
        return ((_dafny.Set({_dafny.SeqWithoutIsStrInference([426, 976, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rkh")))])})) - (_dafny.Set({_dafny.SeqWithoutIsStrInference([len(iife9_()
) for d_53_i0_ in range(default__.abs(-561))])}))).intersection((_dafny.Set({_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ddtk")))])))])})).intersection(_dafny.Set({_dafny.SeqWithoutIsStrInference([80]), _dafny.SeqWithoutIsStrInference([798]), _dafny.SeqWithoutIsStrInference([-939])})))

    @staticmethod
    def fm40(p0, p1, p2, globalState):
        return D3_DC11(len(_dafny.Map({560: True})))

    @staticmethod
    def fm41(p0, p1, p2, p3, globalState):
        def iife10_():
            coll10_ = _dafny.Map()
            compr_10_: int
            for compr_10_ in _dafny.IntegerRange(273, 623):
                d_55_v0_: int = compr_10_
                if ((273) <= (d_55_v0_)) and ((d_55_v0_) < (623)):
                    coll10_[(d_55_v0_) + (518)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dbeopaxvs")))
            return _dafny.Map(coll10_)
        def iife11_():
            coll11_ = _dafny.Map()
            compr_11_: int
            for compr_11_ in _dafny.IntegerRange(478, 125):
                d_56_v1_: int = compr_11_
                if ((478) <= (d_56_v1_)) and ((d_56_v1_) < (125)):
                    coll11_[default__.safeDivisionInt(d_56_v1_, len(_dafny.Map({True: -611})))] = _dafny.CodePoint('g')
            return _dafny.Map(coll11_)
        source2_ = D12_DC31(len(_dafny.SeqWithoutIsStrInference([False, (D12_DC29(_dafny.SeqWithoutIsStrInference([887, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_54_i0_ in range(default__.abs(355))]))]), len(iife10_()
), False)).cf47])), False, False, len(iife11_()
))
        if source2_.is_DC29:
            d_57___mcc_h0_ = source2_.cf45
            d_58___mcc_h1_ = source2_.cf46
            d_59___mcc_h2_ = source2_.cf47
            d_60_cf47_ = d_59___mcc_h2_
            d_61_cf46_ = d_58___mcc_h1_
            d_62_cf45_ = d_57___mcc_h0_
            return D4_DC13(_dafny.CodePoint('i'))
        elif source2_.is_DC30:
            d_63___mcc_h3_ = source2_.cf48
            d_64___mcc_h4_ = source2_.cf49
            d_65___mcc_h5_ = source2_.cf50
            d_66_cf50_ = d_65___mcc_h5_
            d_67_cf49_ = d_64___mcc_h4_
            d_68_cf48_ = d_63___mcc_h3_
            return D4_DC13(d_66_cf50_)
        elif source2_.is_DC31:
            d_69___mcc_h6_ = source2_.cf51
            d_70___mcc_h7_ = source2_.cf52
            d_71___mcc_h8_ = source2_.cf53
            d_72___mcc_h9_ = source2_.cf54
            d_73_cf54_ = d_72___mcc_h9_
            d_74_cf53_ = d_71___mcc_h8_
            d_75_cf52_ = d_70___mcc_h7_
            d_76_cf51_ = d_69___mcc_h6_
            return D4_DC13(_dafny.CodePoint('t'))
        elif True:
            d_77___mcc_h10_ = source2_.cf44
            d_78_cf44_ = d_77___mcc_h10_
            return D4_DC13(_dafny.CodePoint('c'))

    @staticmethod
    def fm42(p0, p1, p2, globalState):
        def iife12_():
            def iife14_():
                coll14_ = _dafny.Map()
                compr_14_: int
                for compr_14_ in _dafny.IntegerRange(322, 477):
                    d_79_v1_: int = compr_14_
                    if ((322) <= (d_79_v1_)) and ((d_79_v1_) < (477)):
                        coll14_[(d_79_v1_) - (26)] = False
                return _dafny.Map(coll14_)
            coll12_ = _dafny.Map()
            def iife13_():
                coll13_ = _dafny.Map()
                compr_13_: int
                for compr_13_ in _dafny.IntegerRange(322, 477):
                    d_79_v1_: int = compr_13_
                    if ((322) <= (d_79_v1_)) and ((d_79_v1_) < (477)):
                        coll13_[(d_79_v1_) - (26)] = False
                return _dafny.Map(coll13_)
            compr_12_: int
            for compr_12_ in (_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iatc"))), len(_dafny.Map({len(iife13_()
            ): True})), (_dafny.MultiSet([True])).cardinality})).Elements:
                d_80_v0_: int = compr_12_
                if (d_80_v0_) in (_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iatc"))), len(_dafny.Map({len(iife14_()
                ): True})), (_dafny.MultiSet([True])).cardinality})):
                    coll12_[(d_80_v0_) + (905)] = False
            return _dafny.Map(coll12_)
        return _dafny.Set({130, (483) * (460), (954) - (len(iife12_()
        )), 997, ((0) - (-292)) * ((0) - (len(_dafny.Map({_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([253])), -790]): 334}))))})

    @staticmethod
    def fm43(p0, p1, globalState):
        if (908) != (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qpimlci")))):
            return _dafny.Map({668: _dafny.SeqWithoutIsStrInference([(0) - (-988)])})
        elif True:
            return _dafny.Map({(_dafny.MultiSet([_dafny.CodePoint('x')])).cardinality: _dafny.SeqWithoutIsStrInference([251 for d_81_i0_ in range(default__.abs(956))])})

    @staticmethod
    def fm44(p0, p1, p2, globalState):
        return ((_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nlcjilh")): False})) | (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cec")): False}))) | (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q")): True}))

    @staticmethod
    def fm45(p0, globalState):
        return _dafny.MultiSet([(_dafny.Map({True: 480})) | (_dafny.Map({True: 554}))])

    @staticmethod
    def fm46(p0, p1, p2, globalState):
        if (len(_dafny.Map({True: True}))) <= (-344):
            return (_dafny.SeqWithoutIsStrInference([_dafny.Map({False: 102}) for d_82_i0_ in range(default__.abs(-546))])) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({True: len(_dafny.Set({True}))})]))
        elif True:
            return _dafny.SeqWithoutIsStrInference([_dafny.Map({not(True): -804}), _dafny.Map({False: 625})])

    @staticmethod
    def m8(globalState):
        d_83_v0_: bool
        d_83_v0_ = False
        d_84_v1_: _dafny.Seq
        d_84_v1_ = _dafny.SeqWithoutIsStrInference([d_83_v0_, True, d_83_v0_, d_83_v0_, d_83_v0_])
        d_85_v2_: int
        d_85_v2_ = 667
        d_86_v3_: _dafny.Seq
        d_86_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rcmkx"))
        (globalState).f13 = ((0) - (len(d_84_v1_))) < ((0) - (default__.safeDivisionInt(d_85_v2_, len(d_86_v3_))))
        d_87_v4_: _dafny.Array
        nw0_ = _dafny.Array(_dafny.CodePoint('D'), 29)
        d_87_v4_ = nw0_
        d_88_v5_: _dafny.Array
        d_88_v5_ = d_87_v4_
        source3_ = d_88_v5_
        d_89___mcc_h0_ = source3_
        d_90_cf42_ = d_89___mcc_h0_
        d_91_v6_: _dafny.Array
        nw1_ = _dafny.Array(False, 23)
        d_91_v6_ = nw1_
        d_92_v7_: D4
        d_92_v7_ = D4_DC14(d_83_v0_)
        d_93_v8_: D4
        d_93_v8_ = D4_DC15(d_92_v7_)
        d_94_v9_: _dafny.MultiSet
        d_94_v9_ = _dafny.MultiSet([d_85_v2_])
        d_95_v10_: _dafny.Array
        nw2_ = _dafny.Array(None, 10)
        nw2_[int(0)] = -911
        nw2_[int(1)] = (0) - (d_85_v2_)
        nw2_[int(2)] = (d_94_v9_).cardinality
        nw2_[int(3)] = 994
        nw2_[int(4)] = d_85_v2_
        nw2_[int(5)] = d_85_v2_
        nw2_[int(6)] = d_85_v2_
        nw2_[int(7)] = d_85_v2_
        nw2_[int(8)] = d_85_v2_
        nw2_[int(9)] = d_85_v2_
        d_95_v10_ = nw2_
        d_96_v11_: _dafny.Array
        d_96_v11_ = d_95_v10_
        d_97_v12_: _dafny.Array
        nw3_ = _dafny.Array(None, 14)
        nw3_[int(0)] = d_95_v10_
        nw3_[int(1)] = d_95_v10_
        nw3_[int(2)] = d_95_v10_
        nw3_[int(3)] = (d_96_v11_)
        nw3_[int(4)] = d_95_v10_
        nw3_[int(5)] = d_95_v10_
        nw3_[int(6)] = d_95_v10_
        nw3_[int(7)] = d_95_v10_
        nw3_[int(8)] = d_95_v10_
        nw3_[int(9)] = d_95_v10_
        nw3_[int(10)] = d_95_v10_
        nw3_[int(11)] = d_95_v10_
        nw3_[int(12)] = d_95_v10_
        nw3_[int(13)] = d_95_v10_
        d_97_v12_ = nw3_
        d_98_v13_: T1
        nw4_ = C5()
        nw4_.ctor__(d_97_v12_, d_85_v2_)
        d_98_v13_ = nw4_
        d_99_v14_: _dafny.Seq
        d_99_v14_ = _dafny.SeqWithoutIsStrInference([d_90_cf42_])
        d_100_v15_: _dafny.Map
        d_100_v15_ = _dafny.Map({d_85_v2_: d_85_v2_})
        d_101_v16_: str
        d_101_v16_ = _dafny.CodePoint('i')
        d_102_v17_: _dafny.Array
        nw5_ = _dafny.Array(None, 28)
        nw5_[int(0)] = 670
        nw5_[int(1)] = default__.fm0(len(d_86_v3_), globalState)
        nw5_[int(2)] = (203 if d_83_v0_ else d_85_v2_)
        nw5_[int(3)] = d_85_v2_
        nw5_[int(4)] = -330
        nw5_[int(5)] = d_85_v2_
        nw5_[int(6)] = d_85_v2_
        nw5_[int(7)] = (D7_DC20(_dafny.SeqWithoutIsStrInference([d_85_v2_ for d_103_i0_ in range(default__.abs(133))]), d_91_v6_, d_93_v8_, d_83_v0_, d_85_v2_)).cf34
        nw5_[int(8)] = (0) - (d_85_v2_)
        nw5_[int(9)] = len(d_84_v1_)
        nw5_[int(10)] = d_85_v2_
        nw5_[int(11)] = len(_dafny.Map({d_98_v13_: d_91_v6_}))
        nw5_[int(12)] = (0) - (d_85_v2_)
        nw5_[int(13)] = d_85_v2_
        nw5_[int(14)] = d_85_v2_
        nw5_[int(15)] = len((_dafny.SeqWithoutIsStrInference([d_83_v0_, default__.fm12(d_83_v0_, d_83_v0_, globalState), d_83_v0_, not(False)])).set(default__.safeIndex(d_85_v2_, len(_dafny.SeqWithoutIsStrInference([d_83_v0_, default__.fm12(d_83_v0_, d_83_v0_, globalState), d_83_v0_, not(False)]))), d_83_v0_))
        nw5_[int(16)] = d_85_v2_
        nw5_[int(17)] = d_85_v2_
        nw5_[int(18)] = d_85_v2_
        nw5_[int(19)] = len((d_99_v14_) + (d_99_v14_))
        nw5_[int(20)] = d_85_v2_
        nw5_[int(21)] = 115
        nw5_[int(22)] = (0) - (len(default__.fm44(d_85_v2_, 361, ((d_86_v3_).set(default__.safeIndex(default__.fm0(len(d_100_v15_), globalState), len(d_86_v3_)), d_101_v16_)).set(default__.safeIndex((d_98_v13_).fm2(globalState), len((d_86_v3_).set(default__.safeIndex(default__.fm0(len(d_100_v15_), globalState), len(d_86_v3_)), d_101_v16_))), d_101_v16_), globalState)))
        nw5_[int(23)] = -60
        nw5_[int(24)] = d_85_v2_
        nw5_[int(25)] = default__.safeModuloInt(d_85_v2_, len(d_86_v3_))
        nw5_[int(26)] = d_85_v2_
        nw5_[int(27)] = d_85_v2_
        d_102_v17_ = nw5_
        d_104_v18_: _dafny.Set
        d_104_v18_ = _dafny.Set({d_85_v2_})
        index0_ = default__.safeIndex(785, (d_95_v10_).length(0))
        (d_95_v10_)[index0_] = default__.safeModuloInt(len(d_104_v18_), d_85_v2_)
        index1_ = default__.safeIndex(892, (d_102_v17_).length(0))
        (d_102_v17_)[index1_] = d_85_v2_
        d_105_v19_: _dafny.Array
        nw6_ = _dafny.Array(_dafny.Array(None, 0), 27)
        d_105_v19_ = nw6_
        d_106_v20_: _dafny.Array
        nw7_ = _dafny.Array(_dafny.Array(None, 0), 4)
        d_106_v20_ = nw7_
        index2_ = default__.safeIndex(954, (d_105_v19_).length(0))
        (d_105_v19_)[index2_] = d_106_v20_
        d_107_v21_: C5
        nw8_ = C5()
        nw8_.ctor__(d_97_v12_, (0) - (d_85_v2_))
        d_107_v21_ = nw8_
        d_108_v22_: _dafny.Set
        d_108_v22_ = _dafny.Set({d_107_v21_, d_107_v21_})
        d_109_v23_: _dafny.Map
        d_109_v23_ = _dafny.Map({d_83_v0_: d_83_v0_})
        d_110_v24_: _dafny.MultiSet
        d_110_v24_ = _dafny.MultiSet([d_83_v0_, d_83_v0_])
        d_111_v25_: _dafny.Map
        d_111_v25_ = _dafny.Map({d_101_v16_: (d_110_v24_).isdisjoint(default__.fm24(d_85_v2_, d_83_v0_, globalState))})
        d_112_v26_: _dafny.Map
        d_112_v26_ = _dafny.Map({143: d_106_v20_})
        index3_ = default__.safeIndex(785, (d_95_v10_).length(0))
        index4_ = default__.safeIndex(892, (d_102_v17_).length(0))
        index5_ = default__.safeIndex(954, (d_105_v19_).length(0))
        rhs0_ = not((d_83_v0_ if (d_108_v22_).issubset(d_108_v22_) else ((d_109_v23_)[d_83_v0_] if (d_83_v0_) in (d_109_v23_) else d_83_v0_)))
        rhs1_ = d_85_v2_
        rhs2_ = ((d_111_v25_)[_dafny.CodePoint('r')] if (_dafny.CodePoint('r')) in (d_111_v25_) else not((d_83_v0_) or (d_83_v0_)))
        rhs3_ = ((d_107_v21_.f30 if d_83_v0_ else d_85_v2_)) - (len((d_86_v3_) + (d_86_v3_)))
        rhs4_ = (((d_112_v26_)[d_107_v21_.f30] if (d_107_v21_.f30) in (d_112_v26_) else d_106_v20_) if d_83_v0_ else d_106_v20_)
        lhs0_ = d_95_v10_
        lhs1_ = default__.safeIndex(785, (d_95_v10_).length(0))
        lhs2_ = globalState
        lhs3_ = d_102_v17_
        lhs4_ = default__.safeIndex(892, (d_102_v17_).length(0))
        lhs5_ = d_105_v19_
        lhs6_ = default__.safeIndex(954, (d_105_v19_).length(0))
        d_83_v0_ = rhs0_
        lhs0_[lhs1_] = rhs1_
        lhs2_.f13 = rhs2_
        lhs3_[lhs4_] = rhs3_
        lhs5_[lhs6_] = rhs4_
        (globalState).f24 = d_104_v18_
        if ((d_84_v1_)[default__.safeIndex((d_102_v17_)[default__.safeIndex(892, (d_102_v17_).length(0))], len(d_84_v1_))] if d_83_v0_ else not(d_83_v0_)):
            d_113_v27_: C0
            nw9_ = C0()
            nw9_.ctor__(not(False))
            d_113_v27_ = nw9_
            rhs5_ = d_113_v27_
            rhs6_ = d_91_v6_
            rhs7_ = d_85_v2_
            lhs7_ = d_107_v21_
            d_113_v27_ = rhs5_
            d_91_v6_ = rhs6_
            lhs7_.f30 = rhs7_
            d_114_v28_: _dafny.Array
            nw10_ = _dafny.Array(None, 27)
            d_114_v28_ = nw10_
            index6_ = default__.safeIndex(980, (d_114_v28_).length(0))
            (d_114_v28_)[index6_] = d_98_v13_
            index7_ = default__.safeIndex(980, (d_114_v28_).length(0))
            (d_114_v28_)[index7_] = d_98_v13_
            d_115_v29_: _dafny.Map
            d_115_v29_ = _dafny.Map({211: d_91_v6_})
            d_116_v30_: _dafny.Set
            d_116_v30_ = _dafny.Set({d_83_v0_})
            d_117_v31_: _dafny.Seq
            d_117_v31_ = _dafny.SeqWithoutIsStrInference([d_104_v18_, d_104_v18_, default__.fm42((d_95_v10_)[default__.safeIndex(785, (d_95_v10_).length(0))], d_116_v30_, (d_95_v10_)[default__.safeIndex(785, (d_95_v10_).length(0))], globalState)])
            d_118_v32_: _dafny.Map
            d_118_v32_ = _dafny.Map({(0) - (d_107_v21_.f30): (d_113_v27_).f32})
            index8_ = default__.safeIndex(785, (d_95_v10_).length(0))
            index9_ = default__.safeIndex(892, (d_102_v17_).length(0))
            rhs8_ = d_115_v29_
            rhs9_ = ((d_104_v18_) - (d_104_v18_)).issubset((d_104_v18_) | ((d_117_v31_)[default__.safeIndex((d_95_v10_)[default__.safeIndex(785, (d_95_v10_).length(0))], len(d_117_v31_))]))
            rhs10_ = default__.safeModuloInt(len(d_109_v23_), d_85_v2_)
            rhs11_ = ((0) - (d_107_v21_.f30) if not((d_84_v1_)[default__.safeIndex((d_94_v9_).cardinality, len(d_84_v1_))]) else (len(_dafny.SeqWithoutIsStrInference([((d_118_v32_)[(d_102_v17_)[default__.safeIndex(892, (d_102_v17_).length(0))]] if ((d_102_v17_)[default__.safeIndex(892, (d_102_v17_).length(0))]) in (d_118_v32_) else d_83_v0_), False])) if d_83_v0_ else d_107_v21_.f30))
            rhs12_ = (0) - ((d_95_v10_)[default__.safeIndex(785, (d_95_v10_).length(0))])
            lhs8_ = globalState
            lhs9_ = globalState
            lhs10_ = d_95_v10_
            lhs11_ = default__.safeIndex(785, (d_95_v10_).length(0))
            lhs12_ = d_102_v17_
            lhs13_ = default__.safeIndex(892, (d_102_v17_).length(0))
            d_115_v29_ = rhs8_
            lhs8_.f22 = rhs9_
            lhs9_.f10 = rhs10_
            lhs10_[lhs11_] = rhs11_
            lhs12_[lhs13_] = rhs12_
            d_119_v33_: _dafny.Map
            d_119_v33_ = _dafny.Map({(d_113_v27_).f32: d_107_v21_.f30})
            d_120_v34_: _dafny.MultiSet
            d_120_v34_ = _dafny.MultiSet([d_119_v33_])
            d_121_v35_: _dafny.Seq
            d_121_v35_ = _dafny.SeqWithoutIsStrInference([d_116_v30_, d_116_v30_])
            d_122_v36_: _dafny.Array
            nw11_ = _dafny.Array(None, 25)
            nw11_[int(0)] = (d_120_v34_).intersection(_dafny.MultiSet([d_119_v33_]))
            nw11_[int(1)] = default__.fm45((0) - ((d_95_v10_)[default__.safeIndex(785, (d_95_v10_).length(0))]), globalState)
            nw11_[int(2)] = _dafny.MultiSet([d_119_v33_, d_119_v33_, d_119_v33_, d_119_v33_, d_119_v33_])
            nw11_[int(3)] = d_120_v34_
            nw11_[int(4)] = d_120_v34_
            nw11_[int(5)] = d_120_v34_
            nw11_[int(6)] = d_120_v34_
            nw11_[int(7)] = d_120_v34_
            nw11_[int(8)] = (d_120_v34_).set(d_119_v33_, default__.abs((d_98_v13_).fm2(globalState)))
            nw11_[int(9)] = d_120_v34_
            nw11_[int(10)] = _dafny.MultiSet(default__.fm46((d_102_v17_)[default__.safeIndex(892, (d_102_v17_).length(0))], d_85_v2_, d_107_v21_.f30, globalState))
            nw11_[int(11)] = (d_120_v34_).intersection(d_120_v34_)
            nw11_[int(12)] = d_120_v34_
            nw11_[int(13)] = (d_120_v34_).intersection(d_120_v34_)
            nw11_[int(14)] = default__.fm45((d_107_v21_).fm3((d_102_v17_)[default__.safeIndex(892, (d_102_v17_).length(0))], (d_113_v27_).f32, -344, D0_DC2((d_102_v17_)[default__.safeIndex(892, (d_102_v17_).length(0))], True), globalState), globalState)
            nw11_[int(15)] = d_120_v34_
            nw11_[int(16)] = d_120_v34_
            nw11_[int(17)] = d_120_v34_
            nw11_[int(18)] = d_120_v34_
            nw11_[int(19)] = _dafny.MultiSet([_dafny.Map({(d_113_v27_).f32: len(d_121_v35_)})])
            nw11_[int(20)] = (d_120_v34_).intersection(d_120_v34_)
            nw11_[int(21)] = d_120_v34_
            nw11_[int(22)] = d_120_v34_
            nw11_[int(23)] = d_120_v34_
            nw11_[int(24)] = ((d_120_v34_).set(d_119_v33_, default__.abs(d_107_v21_.f30))) | (d_120_v34_)
            d_122_v36_ = nw11_
            index10_ = default__.safeIndex(865, (d_122_v36_).length(0))
            (d_122_v36_)[index10_] = d_120_v34_
            index11_ = default__.safeIndex(865, (d_122_v36_).length(0))
            (d_122_v36_)[index11_] = d_120_v34_
            (globalState).f19 = d_101_v16_
        elif True:
            (globalState).f7 = d_83_v0_
            (globalState).f13 = d_83_v0_
            d_123_v37_: C1
            nw12_ = C1()
            nw12_.ctor__()
            d_123_v37_ = nw12_
            d_123_v37_ = d_123_v37_
            d_107_v21_ = d_107_v21_
            (d_107_v21_).f30 = 667
        d_84_v1_ = d_84_v1_
        d_124_v38_: D14
        d_124_v38_ = D14_DC34(True)
        d_125_v39_: _dafny.Array
        nw13_ = _dafny.Array(None, 9)
        nw13_[int(0)] = d_83_v0_
        nw13_[int(1)] = d_83_v0_
        nw13_[int(2)] = (d_84_v1_) == (d_84_v1_)
        nw13_[int(3)] = d_83_v0_
        nw13_[int(4)] = (d_83_v0_) == (d_83_v0_)
        nw13_[int(5)] = d_83_v0_
        nw13_[int(6)] = d_83_v0_
        nw13_[int(7)] = d_83_v0_
        nw13_[int(8)] = False
        d_125_v39_ = nw13_
        index12_ = default__.safeIndex(546, (d_125_v39_).length(0))
        (d_125_v39_)[index12_] = d_83_v0_
        d_126_v40_: str
        d_126_v40_ = _dafny.CodePoint('v')
        index13_ = default__.safeIndex(546, (d_125_v39_).length(0))
        rhs13_ = D14_DC34(d_83_v0_)
        rhs14_ = d_85_v2_
        rhs15_ = d_83_v0_
        rhs16_ = d_85_v2_
        rhs17_ = default__.fm10(_dafny.Map({d_85_v2_: d_85_v2_}), d_126_v40_, d_83_v0_, globalState)
        lhs14_ = globalState
        lhs15_ = d_125_v39_
        lhs16_ = default__.safeIndex(546, (d_125_v39_).length(0))
        d_124_v38_ = rhs13_
        lhs14_.f10 = rhs14_
        lhs15_[lhs16_] = rhs15_
        d_85_v2_ = rhs16_
        d_86_v3_ = rhs17_
        (globalState).f10 = len(d_84_v1_)
        (globalState).f7 = (d_125_v39_)[default__.safeIndex(546, (d_125_v39_).length(0))]
        d_127_v41_: _dafny.Map
        d_127_v41_ = _dafny.Map({(d_125_v39_)[default__.safeIndex(546, (d_125_v39_).length(0))]: (d_125_v39_)[default__.safeIndex(546, (d_125_v39_).length(0))]})
        d_128_v42_: _dafny.Map
        d_128_v42_ = _dafny.Map({(0) - (d_85_v2_): d_85_v2_})
        d_129_v43_: _dafny.Seq
        d_129_v43_ = _dafny.SeqWithoutIsStrInference([len(d_127_v41_), 81, len(d_128_v42_)])
        hi0_ = len(d_129_v43_)
        for d_130_i1_ in range(d_85_v2_, hi0_):
            d_131_v44_: D4
            d_131_v44_ = D4_DC14(d_83_v0_)
            d_132_v45_: _dafny.Seq
            d_132_v45_ = _dafny.SeqWithoutIsStrInference([d_131_v44_, d_131_v44_, d_131_v44_])
            d_132_v45_ = d_132_v45_
            d_133_v46_: _dafny.MultiSet
            d_133_v46_ = _dafny.MultiSet([d_83_v0_, d_83_v0_])
            (globalState).f10 = (((d_133_v46_)[d_83_v0_] if (d_83_v0_) in (d_133_v46_) else d_85_v2_)) + (d_85_v2_)
            d_134_v47_: _dafny.Set
            d_134_v47_ = _dafny.Set({d_130_i1_, d_130_i1_, d_85_v2_})
            rhs18_ = d_87_v4_
            rhs19_ = d_125_v39_
            rhs20_ = (d_127_v41_) == (d_127_v41_)
            rhs21_ = ((d_125_v39_)[default__.safeIndex(546, (d_125_v39_).length(0))] if ((0) - (d_130_i1_)) not in (d_134_v47_) else (d_125_v39_)[default__.safeIndex(546, (d_125_v39_).length(0))])
            lhs17_ = globalState
            lhs18_ = globalState
            d_87_v4_ = rhs18_
            d_125_v39_ = rhs19_
            lhs17_.f13 = rhs20_
            lhs18_.f7 = rhs21_
            d_135_v48_: D7
            d_135_v48_ = D7_DC19()
            source4_ = d_135_v48_
            if source4_.is_DC19:
                d_125_v39_ = d_125_v39_
                (globalState).f13 = True
                d_136_v49_: _dafny.Array
                nw14_ = _dafny.Array(None, 7)
                nw14_[int(0)] = d_86_v3_
                nw14_[int(1)] = d_86_v3_
                nw14_[int(2)] = d_86_v3_
                nw14_[int(3)] = (_dafny.SeqWithoutIsStrInference([d_126_v40_])) + (d_86_v3_)
                nw14_[int(4)] = d_86_v3_
                nw14_[int(5)] = d_86_v3_
                nw14_[int(6)] = d_86_v3_
                d_136_v49_ = nw14_
                index14_ = default__.safeIndex(468, (d_136_v49_).length(0))
                (d_136_v49_)[index14_] = d_86_v3_
                d_137_v50_: C0
                nw15_ = C0()
                nw15_.ctor__(False)
                d_137_v50_ = nw15_
                index15_ = default__.safeIndex(468, (d_136_v49_).length(0))
                rhs22_ = d_86_v3_
                rhs23_ = d_137_v50_
                lhs19_ = d_136_v49_
                lhs20_ = default__.safeIndex(468, (d_136_v49_).length(0))
                lhs19_[lhs20_] = rhs22_
                d_137_v50_ = rhs23_
                d_138_v51_: _dafny.Map
                d_138_v51_ = _dafny.Map({d_126_v40_: d_85_v2_})
                d_139_v52_: _dafny.Seq
                d_139_v52_ = _dafny.SeqWithoutIsStrInference([d_138_v51_, d_138_v51_, d_138_v51_, d_138_v51_, _dafny.Map({d_126_v40_: d_130_i1_})])
                d_140_v53_: D7
                d_140_v53_ = D7_DC18(d_139_v52_)
                d_140_v53_ = d_140_v53_
            elif source4_.is_DC20:
                d_141___mcc_h1_ = source4_.cf30
                d_142___mcc_h2_ = source4_.cf31
                d_143___mcc_h3_ = source4_.cf32
                d_144___mcc_h4_ = source4_.cf33
                d_145___mcc_h5_ = source4_.cf34
                d_146_cf34_ = d_145___mcc_h5_
                d_147_cf33_ = d_144___mcc_h4_
                d_148_cf32_ = d_143___mcc_h3_
                d_149_cf31_ = d_142___mcc_h2_
                d_150_cf30_ = d_141___mcc_h1_
                (globalState).f22 = d_83_v0_
                d_151_v54_: _dafny.Map
                d_151_v54_ = _dafny.Map({d_85_v2_: not (True) or (d_147_cf33_)})
                d_151_v54_ = (d_151_v54_).set(default__.safeModuloInt(d_146_cf34_, d_146_cf34_), d_83_v0_)
                d_129_v43_ = _dafny.SeqWithoutIsStrInference([default__.fm0(d_85_v2_, globalState), d_146_cf34_, d_130_i1_])
                d_152_v55_: _dafny.Array
                def lambda0_(d_153_v3_):
                    def lambda1_(d_154_i2_):
                        return d_153_v3_

                    return lambda1_

                init0_ = lambda0_(d_86_v3_)
                nw16_ = _dafny.Array(None, 19)
                for i0_0_ in range(nw16_.length(0)):
                    nw16_[i0_0_] = init0_(i0_0_)
                d_152_v55_ = nw16_
                index16_ = default__.safeIndex(429, (d_152_v55_).length(0))
                (d_152_v55_)[index16_] = d_86_v3_
                index17_ = default__.safeIndex(429, (d_152_v55_).length(0))
                (d_152_v55_)[index17_] = (((d_86_v3_ if d_147_cf33_ else d_86_v3_)).set(default__.safeIndex(d_130_i1_, len((d_86_v3_ if d_147_cf33_ else d_86_v3_))), d_126_v40_)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uyvafgoga")))
            elif True:
                d_155___mcc_h6_ = source4_.cf29
                d_156_cf29_ = d_155___mcc_h6_
                index18_ = default__.safeIndex(546, (d_125_v39_).length(0))
                rhs24_ = (d_125_v39_)[default__.safeIndex(546, (d_125_v39_).length(0))]
                rhs25_ = (d_125_v39_)[default__.safeIndex(546, (d_125_v39_).length(0))]
                rhs26_ = (0) - (d_85_v2_)
                lhs21_ = globalState
                lhs22_ = d_125_v39_
                lhs23_ = default__.safeIndex(546, (d_125_v39_).length(0))
                lhs24_ = globalState
                lhs21_.f22 = rhs24_
                lhs22_[lhs23_] = rhs25_
                lhs24_.f10 = rhs26_
                d_157_v56_: _dafny.Seq
                d_157_v56_ = _dafny.SeqWithoutIsStrInference([d_84_v1_])
                d_158_v57_: _dafny.Array
                nw17_ = _dafny.Array(None, 2)
                nw17_[int(0)] = ((d_157_v56_)[default__.safeIndex(d_130_i1_, len(d_157_v56_))]) + (d_84_v1_)
                nw17_[int(1)] = _dafny.SeqWithoutIsStrInference([(d_125_v39_)[default__.safeIndex(546, (d_125_v39_).length(0))], d_83_v0_, (d_125_v39_)[default__.safeIndex(546, (d_125_v39_).length(0))], (d_125_v39_)[default__.safeIndex(546, (d_125_v39_).length(0))], d_83_v0_])
                d_158_v57_ = nw17_
                index19_ = default__.safeIndex(701, (d_158_v57_).length(0))
                (d_158_v57_)[index19_] = d_84_v1_
                index20_ = default__.safeIndex(701, (d_158_v57_).length(0))
                (d_158_v57_)[index20_] = (_dafny.SeqWithoutIsStrInference([(d_125_v39_)[default__.safeIndex(546, (d_125_v39_).length(0))], (d_125_v39_)[default__.safeIndex(546, (d_125_v39_).length(0))]]) if d_83_v0_ else d_84_v1_)
                (globalState).f10 = (0) - (d_85_v2_)
                d_85_v2_ = 889

    @staticmethod
    def Main(noArgsParameter__):
        d_159_v0_: _dafny.Array
        nw18_ = _dafny.Array(False, 13)
        d_159_v0_ = nw18_
        d_160_v1_: _dafny.MultiSet
        d_160_v1_ = _dafny.MultiSet([d_159_v0_, d_159_v0_])
        d_161_v2_: bool
        d_161_v2_ = False
        d_162_v3_: _dafny.Seq
        d_162_v3_ = _dafny.SeqWithoutIsStrInference([d_161_v2_])
        d_163_v4_: _dafny.Seq
        d_163_v4_ = _dafny.SeqWithoutIsStrInference([((d_160_v1_)[d_159_v0_] if (d_159_v0_) in (d_160_v1_) else len(d_162_v3_))])
        d_164_v5_: _dafny.Map
        d_164_v5_ = _dafny.Map({d_163_v4_: d_159_v0_})
        d_165_v6_: int
        d_165_v6_ = -551
        d_166_v7_: _dafny.Array
        nw19_ = _dafny.Array(None, 2)
        nw19_[int(0)] = (d_162_v3_).set(default__.safeIndex(d_165_v6_, len(d_162_v3_)), d_161_v2_)
        nw19_[int(1)] = d_162_v3_
        d_166_v7_ = nw19_
        d_167_v8_: _dafny.Seq
        d_167_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wrohtasav"))
        d_168_v9_: _dafny.MultiSet
        d_168_v9_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([d_165_v6_ for d_169_i0_ in range(default__.abs(177))]))])
        d_170_v10_: _dafny.Map
        d_170_v10_ = _dafny.Map({(d_167_v8_)[default__.safeIndex(d_165_v6_, len(d_167_v8_))]: d_168_v9_})
        d_171_v11_: _dafny.Array
        nw20_ = _dafny.Array(None, 4)
        nw20_[int(0)] = ((d_170_v10_)[_dafny.CodePoint('n')] if (_dafny.CodePoint('n')) in (d_170_v10_) else d_168_v9_)
        nw20_[int(1)] = d_168_v9_
        nw20_[int(2)] = d_168_v9_
        nw20_[int(3)] = _dafny.MultiSet([d_165_v6_, 985])
        d_171_v11_ = nw20_
        d_172_v12_: _dafny.Array
        nw21_ = _dafny.Array(_dafny.Array(None, 0), 12)
        d_172_v12_ = nw21_
        d_173_v13_: str
        d_173_v13_ = _dafny.CodePoint('s')
        d_174_v14_: _dafny.Map
        d_174_v14_ = _dafny.Map({d_173_v13_: d_161_v2_})
        d_175_v15_: _dafny.Map
        d_175_v15_ = _dafny.Map({d_174_v14_: d_161_v2_})
        d_176_v16_: _dafny.Map
        d_176_v16_ = _dafny.Map({d_159_v0_: d_161_v2_})
        d_177_v17_: _dafny.Seq
        d_177_v17_ = _dafny.SeqWithoutIsStrInference([d_167_v8_])
        d_178_v18_: _dafny.Array
        nw22_ = _dafny.Array(None, 18)
        nw22_[int(0)] = d_167_v8_
        nw22_[int(1)] = d_167_v8_
        nw22_[int(2)] = d_167_v8_
        nw22_[int(3)] = d_167_v8_
        nw22_[int(4)] = d_167_v8_
        nw22_[int(5)] = (d_167_v8_).set(default__.safeIndex(d_165_v6_, len(d_167_v8_)), d_173_v13_)
        nw22_[int(6)] = d_167_v8_
        nw22_[int(7)] = (d_177_v17_)[default__.safeIndex(d_165_v6_, len(d_177_v17_))]
        nw22_[int(8)] = d_167_v8_
        nw22_[int(9)] = d_167_v8_
        nw22_[int(10)] = (D0_DC0(d_167_v8_)).cf0
        nw22_[int(11)] = d_167_v8_
        nw22_[int(12)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vfacyd"))).set(default__.safeIndex(d_165_v6_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vfacyd")))), (d_167_v8_)[default__.safeIndex(len(d_162_v3_), len(d_167_v8_))])
        nw22_[int(13)] = d_167_v8_
        nw22_[int(14)] = d_167_v8_
        nw22_[int(15)] = d_167_v8_
        nw22_[int(16)] = d_167_v8_
        nw22_[int(17)] = (d_167_v8_).set(default__.safeIndex(d_165_v6_, len(d_167_v8_)), d_173_v13_)
        d_178_v18_ = nw22_
        d_179_v19_: _dafny.Set
        d_179_v19_ = _dafny.Set({len(d_162_v3_)})
        d_180_globalState_: GlobalState
        nw23_ = GlobalState()
        nw23_.ctor__((d_164_v5_) | (d_164_v5_), d_166_v7_, True, _dafny.CodePoint('h'), False, d_171_v11_, d_172_v12_, False, d_168_v9_, 974, -767, d_175_v15_, d_176_v16_, True, True, True, False, 811, True, _dafny.CodePoint('d'), True, d_178_v18_, True, True, d_179_v19_, (d_162_v3_) + (d_162_v3_), True, 80, _dafny.CodePoint('x'))
        d_180_globalState_ = nw23_
        d_181_v20_: _dafny.Map
        d_181_v20_ = _dafny.Map({len(d_163_v4_): False})
        d_181_v20_ = d_181_v20_
        d_182_v21_: D0
        d_182_v21_ = D0_DC1(default__.fm0((_dafny.MultiSet(d_163_v4_)).cardinality, d_180_globalState_), d_165_v6_, d_165_v6_, d_159_v0_)
        pat_let_tv0_ = d_182_v21_
        def iife15_(_pat_let0_0):
            def iife16_(d_183_dt__update__tmp_h0_):
                def iife17_(_pat_let1_0):
                    def iife18_(d_184_dt__update_hcf7_h0_):
                        return D0_DC3(d_184_dt__update_hcf7_h0_)
                    return iife18_(_pat_let1_0)
                return iife17_(pat_let_tv0_)
            return iife16_(_pat_let0_0)
        source5_ = iife15_(D0_DC3(d_182_v21_))
        if source5_.is_DC1:
            d_185___mcc_h0_ = source5_.cf1
            d_186___mcc_h1_ = source5_.cf2
            d_187___mcc_h2_ = source5_.cf3
            d_188___mcc_h3_ = source5_.cf4
            d_189_cf4_ = d_188___mcc_h3_
            d_190_cf3_ = d_187___mcc_h2_
            d_191_cf2_ = d_186___mcc_h1_
            d_192_cf1_ = d_185___mcc_h0_
            (d_180_globalState_).f7 = (d_190_cf3_) >= (d_190_cf3_)
            d_193_v22_: D0
            d_193_v22_ = D0_DC2(d_191_cf2_, d_161_v2_)
            d_193_v22_ = d_193_v22_
            d_194_v23_: C1
            nw24_ = C1()
            nw24_.ctor__()
            d_194_v23_ = nw24_
            d_195_v24_: D0
            d_196_v25_: bool
            out0_: D0
            out1_: bool
            out0_, out1_ = (d_194_v23_).m1((d_163_v4_).set(default__.safeIndex(d_192_cf1_, len(d_163_v4_)), d_192_cf1_), (d_192_cf1_ if d_161_v2_ else d_192_cf1_), d_161_v2_, d_180_globalState_)
            d_195_v24_ = out0_
            d_196_v25_ = out1_
        elif source5_.is_DC2:
            d_197___mcc_h4_ = source5_.cf5
            d_198___mcc_h5_ = source5_.cf6
            d_199_cf6_ = d_198___mcc_h5_
            d_200_cf5_ = d_197___mcc_h4_
            d_201_v26_: _dafny.Map
            d_201_v26_ = _dafny.Map({d_161_v2_: d_199_cf6_})
            d_181_v20_ = (d_181_v20_).set(default__.safeModuloInt(len(d_181_v20_), d_200_cf5_), not((d_200_cf5_) >= (len(((_dafny.SeqWithoutIsStrInference([not(((d_201_v26_)[d_161_v2_] if (d_161_v2_) in (d_201_v26_) else d_161_v2_)), d_199_cf6_])).set(default__.safeIndex(d_200_cf5_, len(_dafny.SeqWithoutIsStrInference([not(((d_201_v26_)[d_161_v2_] if (d_161_v2_) in (d_201_v26_) else d_161_v2_)), d_199_cf6_]))), d_199_cf6_)).set(default__.safeIndex(d_200_cf5_, len((_dafny.SeqWithoutIsStrInference([not(((d_201_v26_)[d_161_v2_] if (d_161_v2_) in (d_201_v26_) else d_161_v2_)), d_199_cf6_])).set(default__.safeIndex(d_200_cf5_, len(_dafny.SeqWithoutIsStrInference([not(((d_201_v26_)[d_161_v2_] if (d_161_v2_) in (d_201_v26_) else d_161_v2_)), d_199_cf6_]))), d_199_cf6_))), d_161_v2_)))))
            default__.m8(d_180_globalState_)
            d_202_v27_: _dafny.Array
            def lambda2_(d_203_i1_):
                return (d_203_i1_) * (-936)

            init1_ = lambda2_
            nw25_ = _dafny.Array(None, 12)
            for i0_1_ in range(nw25_.length(0)):
                nw25_[i0_1_] = init1_(i0_1_)
            d_202_v27_ = nw25_
            index21_ = default__.safeIndex(117, (d_202_v27_).length(0))
            (d_202_v27_)[index21_] = d_165_v6_
            index22_ = default__.safeIndex(117, (d_202_v27_).length(0))
            (d_202_v27_)[index22_] = d_200_cf5_
            index23_ = default__.safeIndex(952, (d_159_v0_).length(0))
            (d_159_v0_)[index23_] = (d_162_v3_)[default__.safeIndex((d_202_v27_)[default__.safeIndex(117, (d_202_v27_).length(0))], len(d_162_v3_))]
            index24_ = default__.safeIndex(952, (d_159_v0_).length(0))
            (d_159_v0_)[index24_] = d_199_cf6_
        elif source5_.is_DC0:
            d_204___mcc_h6_ = source5_.cf0
            d_205_cf0_ = d_204___mcc_h6_
            default__.m8(d_180_globalState_)
            d_206_v28_: _dafny.Array
            nw26_ = _dafny.Array(None, 21)
            nw26_[int(0)] = d_173_v13_
            nw26_[int(1)] = d_173_v13_
            nw26_[int(2)] = _dafny.CodePoint('j')
            nw26_[int(3)] = d_173_v13_
            nw26_[int(4)] = d_173_v13_
            nw26_[int(5)] = (d_205_cf0_)[default__.safeIndex((0) - (d_165_v6_), len(d_205_cf0_))]
            nw26_[int(6)] = _dafny.CodePoint('d')
            nw26_[int(7)] = d_173_v13_
            nw26_[int(8)] = d_173_v13_
            nw26_[int(9)] = d_173_v13_
            nw26_[int(10)] = _dafny.CodePoint('p')
            nw26_[int(11)] = (d_167_v8_)[default__.safeIndex((d_168_v9_).cardinality, len(d_167_v8_))]
            nw26_[int(12)] = d_173_v13_
            nw26_[int(13)] = d_173_v13_
            nw26_[int(14)] = d_173_v13_
            nw26_[int(15)] = d_173_v13_
            nw26_[int(16)] = d_173_v13_
            nw26_[int(17)] = d_173_v13_
            nw26_[int(18)] = d_173_v13_
            nw26_[int(19)] = d_173_v13_
            nw26_[int(20)] = d_173_v13_
            d_206_v28_ = nw26_
            index25_ = default__.safeIndex(472, (d_206_v28_).length(0))
            (d_206_v28_)[index25_] = _dafny.CodePoint('l')
            d_207_v29_: _dafny.Map
            d_207_v29_ = _dafny.Map({d_159_v0_: d_173_v13_})
            index26_ = default__.safeIndex(472, (d_206_v28_).length(0))
            (d_206_v28_)[index26_] = (d_173_v13_ if d_161_v2_ else ((d_207_v29_)[d_159_v0_] if (d_159_v0_) in (d_207_v29_) else d_173_v13_))
            (d_180_globalState_).f10 = (d_165_v6_) * ((0) - ((d_165_v6_) + (d_165_v6_)))
            d_208_v30_: D0
            d_208_v30_ = D0_DC3((d_182_v21_ if d_161_v2_ else d_182_v21_))
            source6_ = d_208_v30_
            if source6_.is_DC1:
                d_209___mcc_h8_ = source6_.cf1
                d_210___mcc_h9_ = source6_.cf2
                d_211___mcc_h10_ = source6_.cf3
                d_212___mcc_h11_ = source6_.cf4
                d_213_cf4_ = d_212___mcc_h11_
                d_214_cf3_ = d_211___mcc_h10_
                d_215_cf2_ = d_210___mcc_h9_
                d_216_cf1_ = d_209___mcc_h8_
                index27_ = default__.safeIndex(377, (d_213_cf4_).length(0))
                (d_213_cf4_)[index27_] = default__.fm12(d_161_v2_, d_161_v2_, d_180_globalState_)
                index28_ = default__.safeIndex(377, (d_213_cf4_).length(0))
                (d_213_cf4_)[index28_] = d_161_v2_
                d_214_cf3_ = d_214_cf3_
                default__.m8(d_180_globalState_)
                d_205_cf0_ = d_205_cf0_
            elif source6_.is_DC2:
                d_217___mcc_h12_ = source6_.cf5
                d_218___mcc_h13_ = source6_.cf6
                d_219_cf6_ = d_218___mcc_h13_
                d_220_cf5_ = d_217___mcc_h12_
                d_221_v31_: _dafny.Array
                nw27_ = _dafny.Array(None, 16)
                d_221_v31_ = nw27_
                d_222_v32_: _dafny.Array
                nw28_ = _dafny.Array(_dafny.Array(None, 0), 25)
                d_222_v32_ = nw28_
                d_223_v33_: T1
                nw29_ = C4()
                nw29_.ctor__(d_222_v32_)
                d_223_v33_ = nw29_
                d_224_v34_: _dafny.Map
                d_224_v34_ = _dafny.Map({d_221_v31_: d_223_v33_})
                d_224_v34_ = (d_224_v34_).set(d_221_v31_, d_223_v33_)
                d_225_v35_: _dafny.Array
                nw30_ = _dafny.Array(None, 4)
                nw30_[int(0)] = d_220_cf5_
                nw30_[int(1)] = d_165_v6_
                nw30_[int(2)] = (0) - (d_220_cf5_)
                nw30_[int(3)] = d_165_v6_
                d_225_v35_ = nw30_
                d_225_v35_ = (d_225_v35_ if d_219_cf6_ else d_225_v35_)
                d_226_v36_: _dafny.Array
                d_226_v36_ = d_206_v28_
                rhs27_ = d_225_v35_
                rhs28_ = (d_165_v6_) != (d_220_cf5_)
                rhs29_ = d_206_v28_
                lhs25_ = d_180_globalState_
                d_225_v35_ = rhs27_
                lhs25_.f7 = rhs28_
                d_226_v36_ = rhs29_
                index29_ = default__.safeIndex(289, (d_225_v35_).length(0))
                (d_225_v35_)[index29_] = d_220_cf5_
                index30_ = default__.safeIndex(289, (d_225_v35_).length(0))
                (d_225_v35_)[index30_] = d_220_cf5_
            elif source6_.is_DC0:
                d_227___mcc_h14_ = source6_.cf0
                d_228_cf0_ = d_227___mcc_h14_
                d_229_v37_: _dafny.Seq
                d_229_v37_ = _dafny.SeqWithoutIsStrInference([d_181_v20_])
                d_229_v37_ = ((d_229_v37_) + (_dafny.SeqWithoutIsStrInference([d_181_v20_])) if d_161_v2_ else _dafny.SeqWithoutIsStrInference([d_181_v20_]))
                d_230_v38_: _dafny.Array
                nw31_ = _dafny.Array(int(0), 29)
                d_230_v38_ = nw31_
                d_231_v39_: _dafny.Seq
                d_231_v39_ = _dafny.SeqWithoutIsStrInference([d_230_v38_, d_230_v38_, d_230_v38_, d_230_v38_, d_230_v38_])
                index31_ = default__.safeIndex(701, (d_230_v38_).length(0))
                (d_230_v38_)[index31_] = len(d_231_v39_)
                index32_ = default__.safeIndex(701, (d_230_v38_).length(0))
                (d_230_v38_)[index32_] = (d_163_v4_)[default__.safeIndex((640) - (len(d_179_v19_)), len(d_163_v4_))]
                index33_ = default__.safeIndex(701, (d_230_v38_).length(0))
                (d_230_v38_)[index33_] = d_165_v6_
                index34_ = default__.safeIndex(491, (d_159_v0_).length(0))
                (d_159_v0_)[index34_] = default__.fm12(False, False, d_180_globalState_)
                index35_ = default__.safeIndex(491, (d_159_v0_).length(0))
                (d_159_v0_)[index35_] = not((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gr"))) not in (d_177_v17_))
            elif True:
                d_232___mcc_h15_ = source6_.cf7
                d_233_cf7_ = d_232___mcc_h15_
                d_165_v6_ = d_165_v6_
                (d_180_globalState_).f10 = (d_163_v4_)[default__.safeIndex((0) - (d_165_v6_), len(d_163_v4_))]
                default__.m8(d_180_globalState_)
                d_234_v40_: _dafny.Map
                d_234_v40_ = _dafny.Map({d_161_v2_: d_206_v28_})
                d_234_v40_ = (d_234_v40_).set(d_161_v2_, d_206_v28_)
        elif True:
            d_235___mcc_h7_ = source5_.cf7
            d_236_cf7_ = d_235___mcc_h7_
            d_166_v7_ = d_166_v7_
            if (d_168_v9_).issubset(_dafny.MultiSet([d_165_v6_])):
                d_237_v41_: C0
                nw32_ = C0()
                nw32_.ctor__(not(not (default__.fm12(d_161_v2_, d_161_v2_, d_180_globalState_)) or (d_161_v2_)))
                d_237_v41_ = nw32_
                (d_180_globalState_).f10 = default__.fm0(default__.safeModuloInt((0) - (d_165_v6_), d_165_v6_), d_180_globalState_)
                (d_180_globalState_).f7 = (d_237_v41_).f32
                d_238_v42_: _dafny.Array
                def lambda3_(d_239_v6_):
                    def lambda4_(d_240_i2_):
                        return (d_240_i2_) * (d_239_v6_)

                    return lambda4_

                init2_ = lambda3_(d_165_v6_)
                nw33_ = _dafny.Array(None, 16)
                for i0_2_ in range(nw33_.length(0)):
                    nw33_[i0_2_] = init2_(i0_2_)
                d_238_v42_ = nw33_
                index36_ = default__.safeIndex(251, (d_238_v42_).length(0))
                (d_238_v42_)[index36_] = d_165_v6_
                d_241_v44_: _dafny.Seq
                def iife19_():
                    coll15_ = _dafny.Map()
                    compr_15_: int
                    for compr_15_ in _dafny.IntegerRange(599, -29):
                        d_242_v43_: int = compr_15_
                        if ((599) <= (d_242_v43_)) and ((d_242_v43_) < (-29)):
                            coll15_[default__.safeDivisionInt(d_242_v43_, len(d_167_v8_))] = d_161_v2_
                    return _dafny.Map(coll15_)
                d_241_v44_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({default__.fm0(d_165_v6_, d_180_globalState_): True}), iife19_()
                ])
                index37_ = default__.safeIndex(251, (d_238_v42_).length(0))
                (d_238_v42_)[index37_] = (default__.safeDivisionInt(d_165_v6_, d_165_v6_)) + ((d_165_v6_) + ((0) - (len(_dafny.SeqWithoutIsStrInference([len(d_241_v44_)])))))
                (d_180_globalState_).f22 = (d_237_v41_).f32
            elif True:
                d_243_v45_: _dafny.Array
                nw34_ = _dafny.Array(int(0), 3)
                d_243_v45_ = nw34_
                d_243_v45_ = d_243_v45_
                d_244_v46_: C1
                nw35_ = C1()
                nw35_.ctor__()
                d_244_v46_ = nw35_
                d_245_v47_: _dafny.MultiSet
                d_245_v47_ = _dafny.MultiSet([d_161_v2_])
                d_246_v48_: _dafny.Map
                d_246_v48_ = _dafny.Map({d_245_v47_: d_244_v46_})
                d_247_v49_: _dafny.Map
                d_247_v49_ = _dafny.Map({len(d_167_v8_): d_244_v46_})
                d_248_v50_: _dafny.Array
                nw36_ = _dafny.Array(None, 26)
                nw36_[int(0)] = d_244_v46_
                nw36_[int(1)] = d_244_v46_
                nw36_[int(2)] = d_244_v46_
                nw36_[int(3)] = d_244_v46_
                nw36_[int(4)] = d_244_v46_
                nw36_[int(5)] = d_244_v46_
                nw36_[int(6)] = d_244_v46_
                nw36_[int(7)] = d_244_v46_
                nw36_[int(8)] = d_244_v46_
                nw36_[int(9)] = d_244_v46_
                nw36_[int(10)] = d_244_v46_
                nw36_[int(11)] = ((d_246_v48_)[d_245_v47_] if (d_245_v47_) in (d_246_v48_) else d_244_v46_)
                nw36_[int(12)] = d_244_v46_
                nw36_[int(13)] = d_244_v46_
                nw36_[int(14)] = d_244_v46_
                nw36_[int(15)] = d_244_v46_
                nw36_[int(16)] = d_244_v46_
                nw36_[int(17)] = d_244_v46_
                nw36_[int(18)] = d_244_v46_
                nw36_[int(19)] = d_244_v46_
                nw36_[int(20)] = d_244_v46_
                nw36_[int(21)] = d_244_v46_
                nw36_[int(22)] = d_244_v46_
                nw36_[int(23)] = d_244_v46_
                nw36_[int(24)] = ((d_247_v49_)[len(_dafny.SeqWithoutIsStrInference([d_161_v2_, d_161_v2_, d_161_v2_]))] if (len(_dafny.SeqWithoutIsStrInference([d_161_v2_, d_161_v2_, d_161_v2_]))) in (d_247_v49_) else d_244_v46_)
                nw36_[int(25)] = d_244_v46_
                d_248_v50_ = nw36_
                index38_ = default__.safeIndex(796, (d_248_v50_).length(0))
                (d_248_v50_)[index38_] = d_244_v46_
                index39_ = default__.safeIndex(796, (d_248_v50_).length(0))
                (d_248_v50_)[index39_] = d_244_v46_
                index40_ = default__.safeIndex(796, (d_248_v50_).length(0))
                (d_248_v50_)[index40_] = (d_244_v46_ if d_161_v2_ else d_244_v46_)
                default__.m8(d_180_globalState_)
                index41_ = default__.safeIndex(587, (d_159_v0_).length(0))
                (d_159_v0_)[index41_] = d_161_v2_
                index42_ = default__.safeIndex(587, (d_159_v0_).length(0))
                (d_159_v0_)[index42_] = not(False)
            d_249_v51_: _dafny.Array
            nw37_ = _dafny.Array(_dafny.Array(None, 0), 6)
            d_249_v51_ = nw37_
            d_250_v52_: _dafny.Map
            d_250_v52_ = _dafny.Map({d_165_v6_: d_249_v51_})
            d_251_v53_: T1
            nw38_ = C4()
            nw38_.ctor__(((d_250_v52_)[d_165_v6_] if (d_165_v6_) in (d_250_v52_) else d_249_v51_))
            d_251_v53_ = nw38_
            rhs30_ = not (not((d_177_v17_) != (d_177_v17_))) or (d_161_v2_)
            rhs31_ = d_251_v53_
            lhs26_ = d_180_globalState_
            lhs26_.f7 = rhs30_
            d_251_v53_ = rhs31_
            rhs32_ = (d_161_v2_) == (d_161_v2_)
            rhs33_ = d_173_v13_
            lhs27_ = d_180_globalState_
            lhs28_ = d_180_globalState_
            lhs27_.f13 = rhs32_
            lhs28_.f19 = rhs33_
        d_252_v54_: _dafny.Set
        d_252_v54_ = _dafny.Set({d_161_v2_})
        source7_ = D8_DC22(len(d_252_v54_), d_161_v2_, d_173_v13_)
        if source7_.is_DC22:
            d_253___mcc_h16_ = source7_.cf36
            d_254___mcc_h17_ = source7_.cf37
            d_255___mcc_h18_ = source7_.cf38
            d_256_cf38_ = d_255___mcc_h18_
            d_257_cf37_ = d_254___mcc_h17_
            d_258_cf36_ = d_253___mcc_h16_
            d_259_v55_: _dafny.Array
            nw39_ = _dafny.Array(int(0), 28)
            d_259_v55_ = nw39_
            d_260_v56_: _dafny.Seq
            d_260_v56_ = _dafny.SeqWithoutIsStrInference([d_259_v55_])
            (d_180_globalState_).f22 = (d_165_v6_) > (len(d_260_v56_))
            d_257_cf37_ = ((d_163_v4_ if not(d_257_cf37_) else d_163_v4_)) <= ((_dafny.SeqWithoutIsStrInference([d_165_v6_ for d_261_i3_ in range(default__.abs(251))])) + (d_163_v4_))
            d_262_v57_: C1
            nw40_ = C1()
            nw40_.ctor__()
            d_262_v57_ = nw40_
            (d_180_globalState_).f10 = default__.fm0(d_165_v6_, d_180_globalState_)
        elif source7_.is_DC23:
            d_263___mcc_h19_ = source7_.cf39
            d_264_cf39_ = d_263___mcc_h19_
            d_265_v58_: _dafny.Seq
            d_265_v58_ = _dafny.SeqWithoutIsStrInference([d_159_v0_])
            d_266_v59_: _dafny.Array
            nw41_ = _dafny.Array(int(0), 6)
            d_266_v59_ = nw41_
            index43_ = default__.safeIndex(599, (d_266_v59_).length(0))
            (d_266_v59_)[index43_] = d_165_v6_
            d_267_v60_: _dafny.Array
            nw42_ = _dafny.Array(_dafny.Array(None, 0), 25)
            d_267_v60_ = nw42_
            d_268_v61_: T0
            nw43_ = C4()
            nw43_.ctor__(d_267_v60_)
            d_268_v61_ = nw43_
            d_269_v62_: _dafny.Seq
            d_269_v62_ = _dafny.SeqWithoutIsStrInference([d_268_v61_])
            d_270_v63_: _dafny.Map
            d_270_v63_ = _dafny.Map({default__.safeDivisionInt(d_165_v6_, (0) - (d_165_v6_)): (d_269_v62_)[default__.safeIndex(d_165_v6_, len(d_269_v62_))]})
            d_271_v64_: _dafny.Map
            d_271_v64_ = _dafny.Map({d_161_v2_: d_161_v2_})
            index44_ = default__.safeIndex(599, (d_266_v59_).length(0))
            rhs34_ = d_265_v58_
            rhs35_ = (d_165_v6_) * ((len(d_271_v64_)) * (d_165_v6_))
            rhs36_ = (d_270_v63_) | ((d_270_v63_) | (d_270_v63_))
            lhs29_ = d_266_v59_
            lhs30_ = default__.safeIndex(599, (d_266_v59_).length(0))
            d_265_v58_ = rhs34_
            lhs29_[lhs30_] = rhs35_
            d_270_v63_ = rhs36_
            d_272_v65_: _dafny.Map
            d_272_v65_ = _dafny.Map({d_266_v59_: d_161_v2_})
            d_273_v66_: _dafny.Map
            d_273_v66_ = _dafny.Map({d_161_v2_: d_173_v13_})
            d_274_v67_: _dafny.Array
            nw44_ = _dafny.Array(None, 16)
            nw44_[int(0)] = d_173_v13_
            nw44_[int(1)] = d_173_v13_
            nw44_[int(2)] = d_173_v13_
            nw44_[int(3)] = d_173_v13_
            nw44_[int(4)] = d_173_v13_
            nw44_[int(5)] = _dafny.CodePoint('l')
            nw44_[int(6)] = d_173_v13_
            nw44_[int(7)] = (d_173_v13_ if d_264_cf39_ else d_173_v13_)
            nw44_[int(8)] = d_173_v13_
            nw44_[int(9)] = _dafny.CodePoint('n')
            nw44_[int(10)] = d_173_v13_
            nw44_[int(11)] = default__.fm11(len((d_272_v65_).set(d_266_v59_, d_264_cf39_)), d_180_globalState_)
            nw44_[int(12)] = ((d_273_v66_)[d_264_cf39_] if (d_264_cf39_) in (d_273_v66_) else d_173_v13_)
            nw44_[int(13)] = d_173_v13_
            nw44_[int(14)] = d_173_v13_
            nw44_[int(15)] = d_173_v13_
            d_274_v67_ = nw44_
            index45_ = default__.safeIndex(990, (d_274_v67_).length(0))
            (d_274_v67_)[index45_] = d_173_v13_
            d_275_v68_: D8
            d_275_v68_ = D8_DC22(245, d_161_v2_, d_173_v13_)
            index46_ = default__.safeIndex(990, (d_274_v67_).length(0))
            (d_274_v67_)[index46_] = (d_173_v13_ if d_264_cf39_ else (d_275_v68_).cf38)
            d_276_v69_: _dafny.Map
            d_276_v69_ = _dafny.Map({(d_266_v59_)[default__.safeIndex(599, (d_266_v59_).length(0))]: d_163_v4_})
            d_276_v69_ = (d_276_v69_).set(d_165_v6_, d_163_v4_)
            d_277_v70_: C3
            nw45_ = C3()
            nw45_.ctor__()
            d_277_v70_ = nw45_
            d_278_v71_: _dafny.Seq
            d_278_v71_ = _dafny.SeqWithoutIsStrInference([d_277_v70_])
            d_165_v6_ = (0) - ((d_165_v6_) * ((len(d_278_v71_)) - (d_165_v6_)))
        elif source7_.is_DC21:
            d_279___mcc_h20_ = source7_.cf35
            d_280_cf35_ = d_279___mcc_h20_
            index47_ = default__.safeIndex(633, (d_159_v0_).length(0))
            (d_159_v0_)[index47_] = d_161_v2_
            index48_ = default__.safeIndex(633, (d_159_v0_).length(0))
            (d_159_v0_)[index48_] = not(d_161_v2_)
            d_252_v54_ = d_252_v54_
            (d_180_globalState_).f10 = default__.safeModuloInt(d_165_v6_, (d_165_v6_) - (d_165_v6_))
            d_281_v72_: _dafny.Map
            d_281_v72_ = _dafny.Map({d_165_v6_: d_165_v6_})
            d_282_v73_: _dafny.MultiSet
            d_282_v73_ = _dafny.MultiSet([d_173_v13_])
            (d_180_globalState_).f10 = ((d_281_v72_)[(0) - ((0) - ((d_282_v73_).cardinality))] if ((0) - ((0) - ((d_282_v73_).cardinality))) in (d_281_v72_) else d_165_v6_)
        elif True:
            d_283___mcc_h21_ = source7_.cf40
            d_284_cf40_ = d_283___mcc_h21_
            d_285_v74_: _dafny.Array
            nw46_ = _dafny.Array(None, 23)
            nw46_[int(0)] = d_173_v13_
            nw46_[int(1)] = d_173_v13_
            nw46_[int(2)] = default__.fm11(d_165_v6_, d_180_globalState_)
            nw46_[int(3)] = d_173_v13_
            nw46_[int(4)] = d_173_v13_
            nw46_[int(5)] = d_173_v13_
            nw46_[int(6)] = d_173_v13_
            nw46_[int(7)] = d_173_v13_
            nw46_[int(8)] = d_173_v13_
            nw46_[int(9)] = d_173_v13_
            nw46_[int(10)] = d_173_v13_
            nw46_[int(11)] = d_173_v13_
            nw46_[int(12)] = d_173_v13_
            nw46_[int(13)] = d_173_v13_
            nw46_[int(14)] = d_173_v13_
            nw46_[int(15)] = d_173_v13_
            nw46_[int(16)] = d_173_v13_
            nw46_[int(17)] = d_173_v13_
            nw46_[int(18)] = d_173_v13_
            nw46_[int(19)] = d_173_v13_
            nw46_[int(20)] = d_173_v13_
            nw46_[int(21)] = d_173_v13_
            nw46_[int(22)] = d_173_v13_
            d_285_v74_ = nw46_
            index49_ = default__.safeIndex(398, (d_285_v74_).length(0))
            (d_285_v74_)[index49_] = d_173_v13_
            index50_ = default__.safeIndex(398, (d_285_v74_).length(0))
            (d_285_v74_)[index50_] = _dafny.CodePoint('n')
            d_286_v75_: D4
            d_286_v75_ = D4_DC14(d_161_v2_)
            d_287_v76_: _dafny.Seq
            d_287_v76_ = _dafny.SeqWithoutIsStrInference([d_163_v4_, d_163_v4_])
            d_288_v77_: D0
            d_288_v77_ = D0_DC2(len(d_287_v76_), d_161_v2_)
            source8_ = (D0_DC2(334, not(d_161_v2_)) if (d_286_v75_).cf25 else d_288_v77_)
            if source8_.is_DC1:
                d_289___mcc_h22_ = source8_.cf1
                d_290___mcc_h23_ = source8_.cf2
                d_291___mcc_h24_ = source8_.cf3
                d_292___mcc_h25_ = source8_.cf4
                d_293_cf4_ = d_292___mcc_h25_
                d_294_cf3_ = d_291___mcc_h24_
                d_295_cf2_ = d_290___mcc_h23_
                d_296_cf1_ = d_289___mcc_h22_
                (d_180_globalState_).f13 = not((default__.fm12(False, d_161_v2_, d_180_globalState_)) == (not(d_161_v2_)))
                index51_ = default__.safeIndex(398, (d_285_v74_).length(0))
                (d_285_v74_)[index51_] = (d_285_v74_)[default__.safeIndex(398, (d_285_v74_).length(0))]
                default__.m8(d_180_globalState_)
                default__.m8(d_180_globalState_)
            elif source8_.is_DC2:
                d_297___mcc_h26_ = source8_.cf5
                d_298___mcc_h27_ = source8_.cf6
                d_299_cf6_ = d_298___mcc_h27_
                d_300_cf5_ = d_297___mcc_h26_
                default__.m8(d_180_globalState_)
                (d_180_globalState_).f7 = d_161_v2_
                d_301_v78_: C1
                nw47_ = C1()
                nw47_.ctor__()
                d_301_v78_ = nw47_
                d_301_v78_ = d_301_v78_
                d_302_v79_: _dafny.Array
                nw48_ = _dafny.Array(D3.default()(), 15)
                d_302_v79_ = nw48_
                d_303_v80_: _dafny.MultiSet
                d_303_v80_ = _dafny.MultiSet([d_299_cf6_])
                index52_ = default__.safeIndex(644, (d_302_v79_).length(0))
                (d_302_v79_)[index52_] = D3_DC10(d_303_v80_)
                d_304_v81_: D3
                d_304_v81_ = D3_DC10(d_303_v80_)
                index53_ = default__.safeIndex(644, (d_302_v79_).length(0))
                (d_302_v79_)[index53_] = d_304_v81_
            elif source8_.is_DC0:
                d_305___mcc_h28_ = source8_.cf0
                d_306_cf0_ = d_305___mcc_h28_
                d_307_v82_: C0
                nw49_ = C0()
                nw49_.ctor__(False)
                d_307_v82_ = nw49_
                d_308_v83_: _dafny.Map
                d_308_v83_ = _dafny.Map({d_307_v82_: d_165_v6_})
                d_308_v83_ = (d_308_v83_).set(d_307_v82_, d_165_v6_)
                d_309_v84_: _dafny.Map
                d_309_v84_ = _dafny.Map({d_161_v2_: d_161_v2_})
                index54_ = default__.safeIndex(367, (d_159_v0_).length(0))
                (d_159_v0_)[index54_] = ((d_309_v84_)[default__.fm12((d_307_v82_).f32, ((d_181_v20_)[d_165_v6_] if (d_165_v6_) in (d_181_v20_) else d_161_v2_), d_180_globalState_)] if (default__.fm12((d_307_v82_).f32, ((d_181_v20_)[d_165_v6_] if (d_165_v6_) in (d_181_v20_) else d_161_v2_), d_180_globalState_)) in (d_309_v84_) else not(d_161_v2_))
                index55_ = default__.safeIndex(367, (d_159_v0_).length(0))
                (d_159_v0_)[index55_] = d_161_v2_
                d_168_v9_ = d_168_v9_
                (d_180_globalState_).f10 = default__.safeDivisionInt(d_165_v6_, (d_165_v6_) * ((0) - (len(_dafny.Set({d_165_v6_})))))
            elif True:
                d_310___mcc_h29_ = source8_.cf7
                d_311_cf7_ = d_310___mcc_h29_
                d_312_v85_: _dafny.Array
                nw50_ = _dafny.Array(_dafny.Array(None, 0), 3)
                d_312_v85_ = nw50_
                index56_ = default__.safeIndex(4, (d_312_v85_).length(0))
                (d_312_v85_)[index56_] = d_159_v0_
                index57_ = default__.safeIndex(398, (d_285_v74_).length(0))
                index58_ = default__.safeIndex(4, (d_312_v85_).length(0))
                rhs37_ = default__.fm11(d_165_v6_, d_180_globalState_)
                rhs38_ = d_159_v0_
                rhs39_ = not(d_161_v2_)
                lhs31_ = d_285_v74_
                lhs32_ = default__.safeIndex(398, (d_285_v74_).length(0))
                lhs33_ = d_312_v85_
                lhs34_ = default__.safeIndex(4, (d_312_v85_).length(0))
                lhs31_[lhs32_] = rhs37_
                lhs33_[lhs34_] = rhs38_
                d_161_v2_ = rhs39_
                d_165_v6_ = d_165_v6_
                (d_180_globalState_).f10 = d_165_v6_
                index59_ = default__.safeIndex(240, (d_171_v11_).length(0))
                (d_171_v11_)[index59_] = d_168_v9_
                index60_ = default__.safeIndex(240, (d_171_v11_).length(0))
                rhs40_ = d_165_v6_
                rhs41_ = (0) - (d_165_v6_)
                rhs42_ = _dafny.MultiSet([d_165_v6_, d_165_v6_])
                lhs35_ = d_180_globalState_
                lhs36_ = d_171_v11_
                lhs37_ = default__.safeIndex(240, (d_171_v11_).length(0))
                lhs35_.f10 = rhs40_
                d_165_v6_ = rhs41_
                lhs36_[lhs37_] = rhs42_
            d_313_v86_: _dafny.Map
            d_313_v86_ = _dafny.Map({d_165_v6_: default__.safeDivisionInt(len(d_252_v54_), len(d_167_v8_))})
            d_314_v87_: C1
            nw51_ = C1()
            nw51_.ctor__()
            d_314_v87_ = nw51_
            d_315_v88_: _dafny.Seq
            d_315_v88_ = _dafny.SeqWithoutIsStrInference([d_314_v87_])
            d_313_v86_ = (d_313_v86_).set(len(_dafny.Map({(d_315_v88_).set(default__.safeIndex(((d_168_v9_)[d_165_v6_] if (d_165_v6_) in (d_168_v9_) else d_165_v6_), len(d_315_v88_)), d_314_v87_): d_161_v2_})), d_165_v6_)
            d_316_v89_: _dafny.Array
            nw52_ = _dafny.Array(_dafny.Array(None, 0), 11)
            d_316_v89_ = nw52_
            d_317_v90_: T1
            nw53_ = C5()
            nw53_.ctor__(d_316_v89_, -624)
            d_317_v90_ = nw53_
            d_317_v90_ = (d_317_v90_ if d_161_v2_ else d_317_v90_)
        (d_180_globalState_).f13 = d_161_v2_
        d_318_i4_: int
        d_318_i4_ = 0
        with _dafny.label("0"):
            while d_161_v2_:
                with _dafny.c_label("0"):
                    if (d_318_i4_) >= (100):
                        raise _dafny.Break("0")
                    d_318_i4_ = (d_318_i4_) + (1)
                    d_319_v91_: _dafny.Map
                    d_319_v91_ = _dafny.Map({d_162_v3_: d_165_v6_})
                    d_320_v92_: _dafny.Map
                    d_320_v92_ = _dafny.Map({d_165_v6_: (0) - (len(d_319_v91_))})
                    d_321_v93_: C0
                    nw54_ = C0()
                    nw54_.ctor__((default__.fm10(d_320_v92_, d_173_v13_, True, d_180_globalState_)) <= (d_167_v8_))
                    d_321_v93_ = nw54_
                    d_322_v94_: _dafny.Array
                    def lambda5_(d_323_v6_):
                        def lambda6_(d_324_i5_):
                            return (d_324_i5_) - (d_323_v6_)

                        return lambda6_

                    init3_ = lambda5_(d_165_v6_)
                    nw55_ = _dafny.Array(None, 1)
                    for i0_3_ in range(nw55_.length(0)):
                        nw55_[i0_3_] = init3_(i0_3_)
                    d_322_v94_ = nw55_
                    d_325_v95_: _dafny.Map
                    d_325_v95_ = _dafny.Map({d_322_v94_: d_165_v6_})
                    (d_180_globalState_).f22 = default__.fm12(d_161_v2_, (d_325_v95_) != (d_325_v95_), d_180_globalState_)
                    d_326_v96_: C3
                    nw56_ = C3()
                    nw56_.ctor__()
                    d_326_v96_ = nw56_
                    d_326_v96_ = d_326_v96_
                    d_327_v97_: _dafny.Array
                    nw57_ = _dafny.Array(_dafny.Array(None, 0), 22)
                    d_327_v97_ = nw57_
                    d_328_v98_: C4
                    nw58_ = C4()
                    nw58_.ctor__(d_327_v97_)
                    d_328_v98_ = nw58_
                    index61_ = default__.safeIndex(899, (d_322_v94_).length(0))
                    (d_322_v94_)[index61_] = d_165_v6_
                    index62_ = default__.safeIndex(899, (d_322_v94_).length(0))
                    rhs43_ = (d_173_v13_) not in (d_167_v8_)
                    rhs44_ = d_328_v98_
                    rhs45_ = d_167_v8_
                    rhs46_ = d_165_v6_
                    lhs38_ = d_180_globalState_
                    lhs39_ = d_322_v94_
                    lhs40_ = default__.safeIndex(899, (d_322_v94_).length(0))
                    lhs38_.f7 = rhs43_
                    d_328_v98_ = rhs44_
                    d_167_v8_ = rhs45_
                    lhs39_[lhs40_] = rhs46_
                    pass
            pass
        d_329_v100_: _dafny.Map
        def iife20_():
            coll16_ = _dafny.Map()
            compr_16_: str
            for compr_16_ in (d_167_v8_).Elements:
                d_330_v99_: str = compr_16_
                if (d_330_v99_) in (d_167_v8_):
                    coll16_[d_330_v99_] = d_165_v6_
            return _dafny.Map(coll16_)
        d_329_v100_ = _dafny.Map({d_165_v6_: iife20_()
        })
        (d_180_globalState_).f10 = ((len(d_329_v100_)) * (d_165_v6_) if default__.fm12(not(d_161_v2_), d_161_v2_, d_180_globalState_) else default__.safeDivisionInt(len(d_167_v8_), d_165_v6_))
        d_331_v101_: D12
        d_331_v101_ = D12_DC28(_dafny.MultiSet(d_163_v4_))
        (d_180_globalState_).f13 = ((d_168_v9_).intersection((d_331_v101_).cf44)).ispropersubset(d_168_v9_)
        d_332_v102_: _dafny.Array
        nw59_ = _dafny.Array(None, 3)
        nw59_[int(0)] = len((_dafny.SeqWithoutIsStrInference([d_161_v2_])) + (d_162_v3_))
        nw59_[int(1)] = d_165_v6_
        nw59_[int(2)] = len(d_167_v8_)
        d_332_v102_ = nw59_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_332_v102_).length(0)):
            d_333_i6_: int = guard_loop_0_
            if (True) and (((0) <= (d_333_i6_)) and ((d_333_i6_) < ((d_332_v102_).length(0)))):
                (d_332_v102_)[(d_333_i6_)] = (d_333_i6_) - ((d_165_v6_) - (d_165_v6_))
        if not(((d_162_v3_) + ((D8_DC21(d_162_v3_)).cf35)) == (d_162_v3_)):
            d_334_v103_: _dafny.Map
            d_334_v103_ = _dafny.Map({d_173_v13_: (d_165_v6_) + (d_165_v6_)})
            d_334_v103_ = (d_334_v103_).set(d_173_v13_, 526)
            d_335_v104_: D3
            d_335_v104_ = D3_DC10(_dafny.MultiSet(d_162_v3_))
            pat_let_tv1_ = d_161_v2_
            pat_let_tv2_ = d_161_v2_
            pat_let_tv3_ = d_180_globalState_
            pat_let_tv4_ = d_162_v3_
            def iife21_(_pat_let2_0):
                def iife22_(d_336_dt__update__tmp_h2_):
                    def iife23_(_pat_let3_0):
                        def iife24_(d_337_dt__update_hcf21_h0_):
                            return D3_DC10(d_337_dt__update_hcf21_h0_)
                        return iife24_(_pat_let3_0)
                    return iife23_((_dafny.MultiSet([default__.fm12(pat_let_tv1_, pat_let_tv2_, pat_let_tv3_)])) - (_dafny.MultiSet(pat_let_tv4_)))
                return iife22_(_pat_let2_0)
            source9_ = iife21_(d_335_v104_)
            if source9_.is_DC11:
                d_338___mcc_h30_ = source9_.cf22
                d_339_cf22_ = d_338___mcc_h30_
                d_340_v105_: C3
                nw60_ = C3()
                nw60_.ctor__()
                d_340_v105_ = nw60_
                d_341_v106_: _dafny.MultiSet
                d_341_v106_ = _dafny.MultiSet([d_340_v105_])
                (d_180_globalState_).f22 = not((((d_341_v106_)[d_340_v105_] if (d_340_v105_) in (d_341_v106_) else (d_340_v105_).fm8(d_339_cf22_, d_165_v6_, d_339_cf22_, d_180_globalState_))) > (((d_168_v9_).cardinality) * (d_339_cf22_)))
                d_342_v107_: D4
                d_342_v107_ = D4_DC13(d_173_v13_)
                d_343_v108_: D4
                d_343_v108_ = D4_DC15(d_342_v107_)
                d_344_v109_: D7
                d_344_v109_ = D7_DC20(_dafny.SeqWithoutIsStrInference([d_339_cf22_]), d_159_v0_, d_343_v108_, d_161_v2_, (d_165_v6_ if d_161_v2_ else d_339_cf22_))
                d_344_v109_ = d_344_v109_
                d_345_v110_: _dafny.Array
                nw61_ = _dafny.Array(_dafny.Array(None, 0), 16)
                d_345_v110_ = nw61_
                d_346_v111_: C4
                nw62_ = C4()
                nw62_.ctor__(d_345_v110_)
                d_346_v111_ = nw62_
                d_347_v112_: T1
                nw63_ = C4()
                nw63_.ctor__((d_346_v111_).f31)
                d_347_v112_ = nw63_
                d_348_v113_: _dafny.Array
                d_348_v113_ = d_332_v102_
                d_349_v114_: _dafny.Array
                nw64_ = _dafny.Array(None, 12)
                nw64_[int(0)] = d_332_v102_
                nw64_[int(1)] = (d_348_v113_)
                nw64_[int(2)] = d_332_v102_
                nw64_[int(3)] = d_332_v102_
                nw64_[int(4)] = d_332_v102_
                nw64_[int(5)] = d_332_v102_
                nw64_[int(6)] = d_332_v102_
                nw64_[int(7)] = d_332_v102_
                nw64_[int(8)] = d_332_v102_
                nw64_[int(9)] = d_332_v102_
                nw64_[int(10)] = d_332_v102_
                nw64_[int(11)] = d_332_v102_
                d_349_v114_ = nw64_
                nw65_ = C5()
                nw65_.ctor__(d_349_v114_, ((0) - (d_165_v6_)) + (d_165_v6_))
                d_347_v112_ = nw65_
            elif source9_.is_DC12:
                d_350___mcc_h31_ = source9_.cf23
                d_351_cf23_ = d_350___mcc_h31_
                (d_180_globalState_).f10 = d_165_v6_
                d_352_v115_: _dafny.Map
                d_352_v115_ = _dafny.Map({d_173_v13_: d_252_v54_})
                d_353_v116_: _dafny.Array
                nw66_ = _dafny.Array(None, 13)
                nw66_[int(0)] = d_252_v54_
                nw66_[int(1)] = (_dafny.Set({d_351_cf23_})) | (d_252_v54_)
                nw66_[int(2)] = d_252_v54_
                nw66_[int(3)] = d_252_v54_
                nw66_[int(4)] = d_252_v54_
                nw66_[int(5)] = _dafny.Set({d_161_v2_, d_351_cf23_})
                nw66_[int(6)] = d_252_v54_
                nw66_[int(7)] = (((d_352_v115_)[d_173_v13_] if (d_173_v13_) in (d_352_v115_) else d_252_v54_) if d_351_cf23_ else _dafny.Set({d_161_v2_, d_351_cf23_}))
                nw66_[int(8)] = _dafny.Set({d_351_cf23_, (d_162_v3_)[default__.safeIndex(-580, len(d_162_v3_))], d_351_cf23_, default__.fm12(not(d_351_cf23_), False, d_180_globalState_)})
                nw66_[int(9)] = d_252_v54_
                nw66_[int(10)] = d_252_v54_
                nw66_[int(11)] = d_252_v54_
                nw66_[int(12)] = d_252_v54_
                d_353_v116_ = nw66_
                d_354_v117_: _dafny.Map
                d_354_v117_ = _dafny.Map({d_161_v2_: not(True)})
                d_355_v118_: _dafny.Map
                d_355_v118_ = (_dafny.Map({d_161_v2_: d_351_cf23_})).set(False, d_351_cf23_)
                d_356_v119_: D12
                d_356_v119_ = D12_DC29(_dafny.SeqWithoutIsStrInference([351 for d_357_i7_ in range(default__.abs(993))]), d_165_v6_, d_161_v2_)
                d_358_v120_: _dafny.Array
                nw67_ = _dafny.Array(None, 26)
                nw67_[int(0)] = (d_354_v117_ if d_161_v2_ else d_355_v118_)
                nw67_[int(1)] = d_355_v118_
                nw67_[int(2)] = d_355_v118_
                nw67_[int(3)] = d_354_v117_
                nw67_[int(4)] = d_354_v117_
                nw67_[int(5)] = d_355_v118_
                nw67_[int(6)] = d_355_v118_
                nw67_[int(7)] = d_355_v118_
                nw67_[int(8)] = d_354_v117_
                nw67_[int(9)] = d_355_v118_
                nw67_[int(10)] = d_355_v118_
                nw67_[int(11)] = d_355_v118_
                nw67_[int(12)] = d_355_v118_
                nw67_[int(13)] = _dafny.Map({d_161_v2_: d_351_cf23_})
                nw67_[int(14)] = d_355_v118_
                nw67_[int(15)] = d_355_v118_
                nw67_[int(16)] = (d_354_v117_).set(d_161_v2_, not(d_351_cf23_))
                nw67_[int(17)] = d_354_v117_
                nw67_[int(18)] = d_355_v118_
                nw67_[int(19)] = d_355_v118_
                nw67_[int(20)] = d_355_v118_
                nw67_[int(21)] = d_355_v118_
                nw67_[int(22)] = d_355_v118_
                nw67_[int(23)] = d_355_v118_
                nw67_[int(24)] = d_355_v118_
                nw67_[int(25)] = _dafny.Map({d_161_v2_: (d_356_v119_).cf47})
                d_358_v120_ = nw67_
                index63_ = default__.safeIndex(180, (d_358_v120_).length(0))
                (d_358_v120_)[index63_] = (d_354_v117_).set(d_351_cf23_, d_161_v2_)
                index64_ = default__.safeIndex(860, (d_332_v102_).length(0))
                (d_332_v102_)[index64_] = default__.fm0(504, d_180_globalState_)
                d_359_v121_: _dafny.Set
                d_359_v121_ = _dafny.Set({(d_167_v8_).set(default__.safeIndex((0) - (d_165_v6_), len(d_167_v8_)), (d_167_v8_)[default__.safeIndex(d_165_v6_, len(d_167_v8_))])})
                index65_ = default__.safeIndex(180, (d_358_v120_).length(0))
                index66_ = default__.safeIndex(860, (d_332_v102_).length(0))
                rhs47_ = d_353_v116_
                rhs48_ = (d_165_v6_) + (len(d_167_v8_))
                rhs49_ = d_355_v118_
                rhs50_ = default__.safeDivisionInt(d_165_v6_, len(d_359_v121_))
                rhs51_ = 951
                lhs41_ = d_358_v120_
                lhs42_ = default__.safeIndex(180, (d_358_v120_).length(0))
                lhs43_ = d_332_v102_
                lhs44_ = default__.safeIndex(860, (d_332_v102_).length(0))
                d_353_v116_ = rhs47_
                d_165_v6_ = rhs48_
                lhs41_[lhs42_] = rhs49_
                lhs43_[lhs44_] = rhs50_
                d_165_v6_ = rhs51_
                (d_180_globalState_).f10 = (d_332_v102_)[default__.safeIndex(860, (d_332_v102_).length(0))]
                d_360_v122_: _dafny.Map
                d_360_v122_ = _dafny.Map({d_167_v8_: (d_181_v20_) == (d_181_v20_)})
                d_361_v123_: _dafny.Array
                nw68_ = _dafny.Array(_dafny.Array(None, 0), 18)
                d_361_v123_ = nw68_
                index67_ = default__.safeIndex(108, (d_361_v123_).length(0))
                (d_361_v123_)[index67_] = d_178_v18_
                d_362_v124_: _dafny.Array
                nw69_ = _dafny.Array(_dafny.Array(None, 0), 16)
                d_362_v124_ = nw69_
                d_363_v125_: C4
                nw70_ = C4()
                nw70_.ctor__(d_362_v124_)
                d_363_v125_ = nw70_
                d_364_v126_: _dafny.MultiSet
                d_364_v126_ = _dafny.MultiSet([d_363_v125_])
                d_365_v127_: _dafny.Seq
                d_365_v127_ = _dafny.SeqWithoutIsStrInference([d_364_v126_])
                index68_ = default__.safeIndex(156, (d_159_v0_).length(0))
                (d_159_v0_)[index68_] = (_dafny.MultiSet([d_363_v125_])) in (d_365_v127_)
                d_366_v128_: _dafny.Map
                d_366_v128_ = _dafny.Map({387: d_167_v8_})
                d_367_v129_: _dafny.Map
                d_367_v129_ = _dafny.Map({True: len(((d_366_v128_)[len(d_167_v8_)] if (len(d_167_v8_)) in (d_366_v128_) else d_167_v8_))})
                d_368_v130_: _dafny.Array
                nw71_ = _dafny.Array(_dafny.Array(None, 0), 27)
                d_368_v130_ = nw71_
                d_369_v131_: C5
                nw72_ = C5()
                nw72_.ctor__(d_368_v130_, -280)
                d_369_v131_ = nw72_
                d_370_v132_: _dafny.Map
                d_370_v132_ = _dafny.Map({d_168_v9_: d_369_v131_})
                index69_ = default__.safeIndex(860, (d_332_v102_).length(0))
                index70_ = default__.safeIndex(108, (d_361_v123_).length(0))
                index71_ = default__.safeIndex(156, (d_159_v0_).length(0))
                rhs52_ = (d_165_v6_) * (((d_367_v129_)[d_161_v2_] if (d_161_v2_) in (d_367_v129_) else d_165_v6_))
                rhs53_ = (_dafny.Map({d_167_v8_: d_351_cf23_})) | (d_360_v122_)
                rhs54_ = d_178_v18_
                rhs55_ = (default__.fm0(len(d_370_v132_), d_180_globalState_)) >= (-786)
                rhs56_ = d_161_v2_
                lhs45_ = d_332_v102_
                lhs46_ = default__.safeIndex(860, (d_332_v102_).length(0))
                lhs47_ = d_361_v123_
                lhs48_ = default__.safeIndex(108, (d_361_v123_).length(0))
                lhs49_ = d_180_globalState_
                lhs50_ = d_159_v0_
                lhs51_ = default__.safeIndex(156, (d_159_v0_).length(0))
                lhs45_[lhs46_] = rhs52_
                d_360_v122_ = rhs53_
                lhs47_[lhs48_] = rhs54_
                lhs49_.f7 = rhs55_
                lhs50_[lhs51_] = rhs56_
            elif True:
                d_371___mcc_h32_ = source9_.cf21
                d_372_cf21_ = d_371___mcc_h32_
                rhs57_ = default__.fm12(d_161_v2_, d_161_v2_, d_180_globalState_)
                rhs58_ = True
                rhs59_ = _dafny.SeqWithoutIsStrInference([d_165_v6_, d_165_v6_, d_165_v6_, d_165_v6_, d_165_v6_])
                lhs52_ = d_180_globalState_
                lhs53_ = d_180_globalState_
                lhs52_.f7 = rhs57_
                lhs53_.f7 = rhs58_
                d_163_v4_ = rhs59_
                d_167_v8_ = (((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qmpwft"))) + (d_167_v8_)) + (d_167_v8_)).set(default__.safeIndex(d_165_v6_, len(((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qmpwft"))) + (d_167_v8_)) + (d_167_v8_))), d_173_v13_)
                d_373_v133_: _dafny.Array
                nw73_ = _dafny.Array(None, 24)
                d_373_v133_ = nw73_
                d_374_v134_: C3
                nw74_ = C3()
                nw74_.ctor__()
                d_374_v134_ = nw74_
                index72_ = default__.safeIndex(127, (d_373_v133_).length(0))
                (d_373_v133_)[index72_] = d_374_v134_
                index73_ = default__.safeIndex(127, (d_373_v133_).length(0))
                (d_373_v133_)[index73_] = d_374_v134_
                index74_ = default__.safeIndex(23, (d_159_v0_).length(0))
                (d_159_v0_)[index74_] = default__.fm12(not(d_161_v2_), d_161_v2_, d_180_globalState_)
                index75_ = default__.safeIndex(23, (d_159_v0_).length(0))
                rhs60_ = d_165_v6_
                rhs61_ = default__.fm12(d_161_v2_, d_161_v2_, d_180_globalState_)
                rhs62_ = ((d_374_v134_).fm8(d_165_v6_, 184, d_165_v6_, d_180_globalState_) if True else (d_374_v134_).fm8(len(d_179_v19_), d_165_v6_, len(d_162_v3_), d_180_globalState_))
                lhs54_ = d_159_v0_
                lhs55_ = default__.safeIndex(23, (d_159_v0_).length(0))
                d_165_v6_ = rhs60_
                lhs54_[lhs55_] = rhs61_
                d_165_v6_ = rhs62_
            d_375_v135_: _dafny.Array
            nw75_ = _dafny.Array(None, 5)
            nw75_[int(0)] = d_332_v102_
            nw75_[int(1)] = d_332_v102_
            nw75_[int(2)] = d_332_v102_
            nw75_[int(3)] = d_332_v102_
            nw75_[int(4)] = d_332_v102_
            d_375_v135_ = nw75_
            d_376_v136_: _dafny.Map
            d_376_v136_ = _dafny.Map({_dafny.MultiSet(d_163_v4_): d_165_v6_})
            d_377_v137_: _dafny.Map
            d_377_v137_ = _dafny.Map({d_376_v136_: _dafny.Set({d_161_v2_})})
            d_378_v139_: _dafny.Map
            def iife25_():
                coll17_ = _dafny.Map()
                compr_17_: int
                for compr_17_ in (d_179_v19_).Elements:
                    d_379_v138_: int = compr_17_
                    if (d_379_v138_) in (d_179_v19_):
                        coll17_[(d_379_v138_) + (-583)] = d_165_v6_
                return _dafny.Map(coll17_)
            d_378_v139_ = _dafny.Map({((d_377_v137_)[d_376_v136_] if (d_376_v136_) in (d_377_v137_) else d_252_v54_): (0) - ((_dafny.MultiSet([iife25_()
            ])).cardinality)})
            d_380_v140_: T1
            nw76_ = C5()
            nw76_.ctor__(d_375_v135_, (0) - (((d_378_v139_)[d_252_v54_] if (d_252_v54_) in (d_378_v139_) else d_165_v6_)))
            d_380_v140_ = nw76_
            d_381_v141_: _dafny.Array
            nw77_ = _dafny.Array(None, 2)
            nw77_[int(0)] = d_380_v140_
            nw77_[int(1)] = d_380_v140_
            d_381_v141_ = nw77_
            index76_ = default__.safeIndex(288, (d_381_v141_).length(0))
            (d_381_v141_)[index76_] = d_380_v140_
            index77_ = default__.safeIndex(288, (d_381_v141_).length(0))
            rhs63_ = default__.safeModuloInt(d_165_v6_, (0) - ((0) - (d_165_v6_)))
            rhs64_ = d_380_v140_
            rhs65_ = (0) - (len(_dafny.SeqWithoutIsStrInference([((d_181_v20_)[len(d_179_v19_)] if (len(d_179_v19_)) in (d_181_v20_) else d_161_v2_)])))
            lhs56_ = d_180_globalState_
            lhs57_ = d_381_v141_
            lhs58_ = default__.safeIndex(288, (d_381_v141_).length(0))
            lhs59_ = d_180_globalState_
            lhs56_.f10 = rhs63_
            lhs57_[lhs58_] = rhs64_
            lhs59_.f10 = rhs65_
            d_382_v142_: D12
            d_382_v142_ = D12_DC31(d_165_v6_, not(False), d_161_v2_, (d_168_v9_).cardinality)
            d_383_v143_: D8
            d_383_v143_ = D8_DC23((d_382_v142_).cf52)
            d_383_v143_ = d_383_v143_
            d_384_v144_: C0
            nw78_ = C0()
            nw78_.ctor__(d_161_v2_)
            d_384_v144_ = nw78_
        elif True:
            default__.m8(d_180_globalState_)
            (d_180_globalState_).f10 = (0) - (default__.safeModuloInt(default__.fm0(d_165_v6_, d_180_globalState_), d_165_v6_))
            (d_180_globalState_).f7 = not((d_167_v8_) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "glmboqi"))))
            d_385_v145_: C2
            nw79_ = C2()
            nw79_.ctor__((D0_DC0(d_167_v8_)).cf0)
            d_385_v145_ = nw79_
            (d_385_v145_).f33 = d_385_v145_.f33
        d_332_v102_ = d_332_v102_
        if d_161_v2_:
            d_165_v6_ = d_165_v6_
            d_386_v146_: _dafny.Map
            d_386_v146_ = _dafny.Map({True: d_165_v6_})
            d_386_v146_ = (d_386_v146_).set((d_161_v2_ if d_161_v2_ else d_161_v2_), d_165_v6_)
            nw80_ = _dafny.Array(None, 2)
            nw80_[int(0)] = not(not (d_161_v2_) or (d_161_v2_))
            nw80_[int(1)] = d_161_v2_
            d_159_v0_ = nw80_
            d_387_v147_: _dafny.Array
            nw81_ = _dafny.Array(D8.default()(), 18)
            d_387_v147_ = nw81_
            d_388_v148_: D8
            d_388_v148_ = D8_DC22(d_165_v6_, d_161_v2_, _dafny.CodePoint('v'))
            index78_ = default__.safeIndex(886, (d_387_v147_).length(0))
            (d_387_v147_)[index78_] = d_388_v148_
            index79_ = default__.safeIndex(886, (d_387_v147_).length(0))
            (d_387_v147_)[index79_] = d_388_v148_
            d_389_v149_: _dafny.Array
            nw82_ = _dafny.Array(_dafny.Array(None, 0), 24)
            d_389_v149_ = nw82_
            d_390_v150_: _dafny.Array
            d_390_v150_ = d_389_v149_
            d_391_v151_: C4
            nw83_ = C4()
            nw83_.ctor__((d_390_v150_))
            d_391_v151_ = nw83_
        elif True:
            d_167_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "faliulvks"))
            d_392_v152_: _dafny.Array
            nw84_ = _dafny.Array(None, 22)
            nw84_[int(0)] = d_332_v102_
            nw84_[int(1)] = d_332_v102_
            nw84_[int(2)] = d_332_v102_
            nw84_[int(3)] = d_332_v102_
            nw84_[int(4)] = d_332_v102_
            nw84_[int(5)] = d_332_v102_
            nw84_[int(6)] = d_332_v102_
            nw84_[int(7)] = d_332_v102_
            nw84_[int(8)] = d_332_v102_
            nw84_[int(9)] = d_332_v102_
            nw84_[int(10)] = d_332_v102_
            nw84_[int(11)] = d_332_v102_
            nw84_[int(12)] = d_332_v102_
            nw84_[int(13)] = d_332_v102_
            nw84_[int(14)] = d_332_v102_
            nw84_[int(15)] = d_332_v102_
            nw84_[int(16)] = d_332_v102_
            nw84_[int(17)] = d_332_v102_
            nw84_[int(18)] = d_332_v102_
            nw84_[int(19)] = d_332_v102_
            nw84_[int(20)] = d_332_v102_
            nw84_[int(21)] = d_332_v102_
            d_392_v152_ = nw84_
            d_393_v153_: C5
            nw85_ = C5()
            nw85_.ctor__(d_392_v152_, d_165_v6_)
            d_393_v153_ = nw85_
            d_394_v154_: _dafny.MultiSet
            d_394_v154_ = _dafny.MultiSet([d_332_v102_, d_332_v102_, d_332_v102_])
            rhs66_ = d_393_v153_
            rhs67_ = (d_161_v2_) in (d_252_v54_)
            rhs68_ = d_332_v102_
            rhs69_ = _dafny.MultiSet([d_332_v102_, d_332_v102_, d_332_v102_])
            rhs70_ = d_332_v102_
            lhs60_ = d_180_globalState_
            d_393_v153_ = rhs66_
            lhs60_.f13 = rhs67_
            d_332_v102_ = rhs68_
            d_394_v154_ = rhs69_
            d_332_v102_ = rhs70_
            d_395_v155_: _dafny.Array
            nw86_ = _dafny.Array(_dafny.Array(None, 0), 28)
            d_395_v155_ = nw86_
            d_396_v156_: T1
            nw87_ = C4()
            nw87_.ctor__(d_395_v155_)
            d_396_v156_ = nw87_
            d_397_v157_: _dafny.Set
            d_397_v157_ = _dafny.Set({(D14_DC33(d_396_v156_)).cf56, d_396_v156_, d_396_v156_})
            d_398_v158_: _dafny.Map
            d_398_v158_ = _dafny.Map({d_397_v157_: (d_177_v17_)[default__.safeIndex(d_393_v153_.f30, len(d_177_v17_))]})
            d_398_v158_ = (d_398_v158_).set(d_397_v157_, (d_167_v8_) + (d_167_v8_))
            d_399_v159_: _dafny.Array
            def lambda7_(d_400_v2_, d_401_v3_, d_402_v6_):
                def lambda8_(d_403_i8_):
                    return (_dafny.Map({d_400_v2_: (d_401_v3_)[default__.safeIndex(d_402_v6_, len(d_401_v3_))]})) | (_dafny.Map({d_400_v2_: d_400_v2_}))

                return lambda8_

            init4_ = lambda7_(d_161_v2_, d_162_v3_, d_165_v6_)
            nw88_ = _dafny.Array(None, 24)
            for i0_4_ in range(nw88_.length(0)):
                nw88_[i0_4_] = init4_(i0_4_)
            d_399_v159_ = nw88_
            d_404_v160_: _dafny.Map
            d_404_v160_ = _dafny.Map({d_161_v2_: d_161_v2_})
            index80_ = default__.safeIndex(459, (d_399_v159_).length(0))
            (d_399_v159_)[index80_] = (d_404_v160_).set(d_161_v2_, (d_393_v153_).fm4(d_167_v8_, d_165_v6_, d_180_globalState_))
            index81_ = default__.safeIndex(459, (d_399_v159_).length(0))
            (d_399_v159_)[index81_] = ((d_404_v160_) | (d_404_v160_)) | (_dafny.Map({False: d_161_v2_}))
            (d_393_v153_).f30 = d_393_v153_.f30
        d_405_v161_: D4
        d_405_v161_ = D4_DC15(D4_DC14(d_161_v2_))
        (d_180_globalState_).f10 = (D7_DC20((d_163_v4_).set(default__.safeIndex(d_165_v6_, len(d_163_v4_)), d_165_v6_), d_159_v0_, d_405_v161_, d_161_v2_, d_165_v6_)).cf34
        if d_161_v2_:
            d_406_v162_: D7
            d_406_v162_ = D7_DC19()
            rhs71_ = d_161_v2_
            rhs72_ = default__.fm12(d_161_v2_, d_161_v2_, d_180_globalState_)
            rhs73_ = D7_DC19()
            lhs61_ = d_180_globalState_
            lhs62_ = d_180_globalState_
            lhs61_.f13 = rhs71_
            lhs62_.f13 = rhs72_
            d_406_v162_ = rhs73_
            (d_180_globalState_).f13 = d_161_v2_
            if (d_162_v3_)[default__.safeIndex(537, len(d_162_v3_))]:
                (d_180_globalState_).f10 = ((0) - (d_165_v6_)) * (d_165_v6_)
                d_407_v163_: _dafny.Map
                d_407_v163_ = _dafny.Map({d_163_v4_: d_161_v2_})
                d_407_v163_ = d_407_v163_
                d_167_v8_ = d_167_v8_
                d_163_v4_ = (d_163_v4_) + (_dafny.SeqWithoutIsStrInference([d_165_v6_]))
                d_408_v164_: D4
                d_408_v164_ = D4_DC13(d_173_v13_)
                pat_let_tv5_ = d_173_v13_
                d_409_v165_: _dafny.Map
                def iife26_(_pat_let4_0):
                    def iife27_(d_410_dt__update__tmp_h5_):
                        def iife28_(_pat_let5_0):
                            def iife29_(d_411_dt__update_hcf24_h0_):
                                return D4_DC13(d_411_dt__update_hcf24_h0_)
                            return iife29_(_pat_let5_0)
                        return iife28_(pat_let_tv5_)
                    return iife27_(_pat_let4_0)
                d_409_v165_ = _dafny.Map({iife26_(d_408_v164_): (d_165_v6_ if d_161_v2_ else d_165_v6_)})
                (d_180_globalState_).f10 = ((d_409_v165_)[d_408_v164_] if (d_408_v164_) in (d_409_v165_) else d_165_v6_)
            elif True:
                d_412_v166_: _dafny.Map
                d_412_v166_ = _dafny.Map({not(d_161_v2_): d_161_v2_})
                d_413_v167_: _dafny.Map
                d_413_v167_ = _dafny.Map({(d_412_v166_).set(False, d_161_v2_): d_165_v6_})
                rhs74_ = (802) * (d_165_v6_)
                rhs75_ = (d_159_v0_ if d_161_v2_ else d_159_v0_)
                rhs76_ = d_413_v167_
                rhs77_ = not((d_181_v20_) != (_dafny.Map({d_165_v6_: d_161_v2_})))
                rhs78_ = d_161_v2_
                lhs63_ = d_180_globalState_
                lhs64_ = d_180_globalState_
                lhs65_ = d_180_globalState_
                lhs63_.f10 = rhs74_
                d_159_v0_ = rhs75_
                d_413_v167_ = rhs76_
                lhs64_.f22 = rhs77_
                lhs65_.f7 = rhs78_
                index82_ = default__.safeIndex(739, (d_159_v0_).length(0))
                (d_159_v0_)[index82_] = ((d_412_v166_)[not(d_161_v2_)] if (not(d_161_v2_)) in (d_412_v166_) else d_161_v2_)
                index83_ = default__.safeIndex(739, (d_159_v0_).length(0))
                (d_159_v0_)[index83_] = d_161_v2_
                d_414_v168_: _dafny.Array
                def lambda9_(d_415_v0_):
                    def lambda10_(d_416_i9_):
                        return (d_415_v0_)[default__.safeIndex(739, (d_415_v0_).length(0))]

                    return lambda10_

                init5_ = lambda9_(d_159_v0_)
                nw89_ = _dafny.Array(None, 25)
                for i0_5_ in range(nw89_.length(0)):
                    nw89_[i0_5_] = init5_(i0_5_)
                d_414_v168_ = nw89_
                d_417_v169_: _dafny.Map
                d_417_v169_ = _dafny.Map({d_414_v168_: d_332_v102_})
                d_417_v169_ = (d_417_v169_).set(d_414_v168_, d_332_v102_)
                d_181_v20_ = (d_181_v20_).set(len(d_167_v8_), (d_159_v0_)[default__.safeIndex(739, (d_159_v0_).length(0))])
                (d_180_globalState_).f10 = d_165_v6_
            d_418_v170_: C2
            nw90_ = C2()
            nw90_.ctor__(d_167_v8_)
            d_418_v170_ = nw90_
            index84_ = default__.safeIndex(511, (d_332_v102_).length(0))
            (d_332_v102_)[index84_] = len(d_167_v8_)
            index85_ = default__.safeIndex(511, (d_332_v102_).length(0))
            (d_332_v102_)[index85_] = len((d_162_v3_).set(default__.safeIndex(len((d_179_v19_).intersection(d_179_v19_)), len(d_162_v3_)), not(d_161_v2_)))
        elif True:
            d_419_v171_: _dafny.Map
            d_419_v171_ = _dafny.Map({-338: d_332_v102_})
            d_420_v172_: _dafny.Array
            d_420_v172_ = d_332_v102_
            d_421_v173_: _dafny.Array
            nw91_ = _dafny.Array(None, 23)
            nw91_[int(0)] = d_332_v102_
            nw91_[int(1)] = d_332_v102_
            nw91_[int(2)] = d_332_v102_
            nw91_[int(3)] = d_332_v102_
            nw91_[int(4)] = ((d_419_v171_)[d_165_v6_] if (d_165_v6_) in (d_419_v171_) else d_332_v102_)
            nw91_[int(5)] = d_332_v102_
            nw91_[int(6)] = d_332_v102_
            nw91_[int(7)] = d_332_v102_
            nw91_[int(8)] = d_332_v102_
            nw91_[int(9)] = d_332_v102_
            nw91_[int(10)] = d_332_v102_
            nw91_[int(11)] = d_332_v102_
            nw91_[int(12)] = (d_420_v172_)
            nw91_[int(13)] = d_332_v102_
            nw91_[int(14)] = d_332_v102_
            nw91_[int(15)] = d_332_v102_
            nw91_[int(16)] = d_332_v102_
            nw91_[int(17)] = d_332_v102_
            nw91_[int(18)] = d_332_v102_
            nw91_[int(19)] = d_332_v102_
            nw91_[int(20)] = d_332_v102_
            nw91_[int(21)] = d_332_v102_
            nw91_[int(22)] = d_332_v102_
            d_421_v173_ = nw91_
            d_422_v174_: C5
            nw92_ = C5()
            nw92_.ctor__(d_421_v173_, d_165_v6_)
            d_422_v174_ = nw92_
            (d_180_globalState_).f13 = d_161_v2_
            d_423_v175_: _dafny.Set
            d_423_v175_ = _dafny.Set({(d_168_v9_) | (_dafny.MultiSet([len(d_163_v4_)])), d_168_v9_, (d_168_v9_) | (d_168_v9_)})
            rhs79_ = (d_423_v175_).intersection(d_423_v175_)
            rhs80_ = d_161_v2_
            rhs81_ = len(default__.fm43((d_173_v13_ if d_161_v2_ else d_173_v13_), _dafny.Map({d_161_v2_: d_173_v13_}), d_180_globalState_))
            lhs66_ = d_180_globalState_
            lhs67_ = d_180_globalState_
            d_423_v175_ = rhs79_
            lhs66_.f7 = rhs80_
            lhs67_.f10 = rhs81_
            (d_180_globalState_).f10 = len(d_167_v8_)
            d_424_v176_: _dafny.Map
            d_424_v176_ = _dafny.Map({d_173_v13_: d_159_v0_})
            d_425_v177_: _dafny.Map
            d_425_v177_ = _dafny.Map({d_422_v174_: d_161_v2_})
            d_426_v178_: _dafny.Seq
            d_426_v178_ = _dafny.SeqWithoutIsStrInference([d_159_v0_, d_159_v0_])
            d_427_v179_: _dafny.Map
            d_427_v179_ = _dafny.Map({len(d_425_v177_): _dafny.Map({_dafny.CodePoint('o'): (d_426_v178_)[default__.safeIndex(381, len(d_426_v178_))]})})
            d_428_v180_: _dafny.Array
            nw93_ = _dafny.Array(None, 16)
            nw93_[int(0)] = d_424_v176_
            nw93_[int(1)] = (d_424_v176_) | (d_424_v176_)
            nw93_[int(2)] = d_424_v176_
            nw93_[int(3)] = (_dafny.Map({default__.fm11(d_165_v6_, d_180_globalState_): d_159_v0_})) | (d_424_v176_)
            nw93_[int(4)] = d_424_v176_
            nw93_[int(5)] = ((d_427_v179_)[d_422_v174_.f30] if (d_422_v174_.f30) in (d_427_v179_) else d_424_v176_)
            nw93_[int(6)] = (d_424_v176_) | (d_424_v176_)
            nw93_[int(7)] = (d_424_v176_) | (d_424_v176_)
            nw93_[int(8)] = d_424_v176_
            nw93_[int(9)] = d_424_v176_
            nw93_[int(10)] = _dafny.Map({d_173_v13_: d_159_v0_})
            nw93_[int(11)] = (d_424_v176_).set(_dafny.CodePoint('d'), d_159_v0_)
            nw93_[int(12)] = _dafny.Map({d_173_v13_: d_159_v0_})
            nw93_[int(13)] = d_424_v176_
            nw93_[int(14)] = (d_424_v176_).set(d_173_v13_, d_159_v0_)
            nw93_[int(15)] = (d_424_v176_) | (d_424_v176_)
            d_428_v180_ = nw93_
            index86_ = default__.safeIndex(522, (d_428_v180_).length(0))
            (d_428_v180_)[index86_] = d_424_v176_
            index87_ = default__.safeIndex(522, (d_428_v180_).length(0))
            (d_428_v180_)[index87_] = ((d_424_v176_) | ((d_424_v176_).set(d_173_v13_, d_159_v0_))) | (d_424_v176_)
        d_167_v8_ = (d_167_v8_) + (d_167_v8_)
        default__.m8(d_180_globalState_)
        default__.m8(d_180_globalState_)
        _dafny.print(_dafny.string_of((d_159_v0_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_159_v0_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_160_v1_).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_161_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_162_v3_) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_163_v4_) == (_dafny.SeqWithoutIsStrInference([-551, -551, -551, -551, -551]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_164_v5_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_165_v6_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_166_v7_)[0]) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_166_v7_)[1]) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_167_v8_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v9_) == (_dafny.MultiSet([177]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_170_v10_) == (_dafny.Map({_dafny.CodePoint('w'): _dafny.MultiSet([177])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_171_v11_)[0]) == (_dafny.MultiSet([177]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_171_v11_)[1]) == (_dafny.MultiSet([177]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_171_v11_)[2]) == (_dafny.MultiSet([177]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_171_v11_)[3]) == (_dafny.MultiSet([-551, 985]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_173_v13_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_174_v14_) == (_dafny.Map({_dafny.CodePoint('s'): False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_175_v15_) == (_dafny.Map({_dafny.Map({_dafny.CodePoint('s'): False}): False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_176_v16_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_177_v17_) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wrohtasav"))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_178_v18_)[0]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_178_v18_)[1]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_178_v18_)[2]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_178_v18_)[3]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_178_v18_)[4]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_178_v18_)[5]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_178_v18_)[6]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_178_v18_)[7]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_178_v18_)[8]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_178_v18_)[9]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_178_v18_)[10]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_178_v18_)[11]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_178_v18_)[12]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_178_v18_)[13]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_178_v18_)[14]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_178_v18_)[15]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_178_v18_)[16]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_178_v18_)[17]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_179_v19_) == (_dafny.Set({1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_180_globalState_).f0)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_180_globalState_).f1)[0]) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_180_globalState_).f1)[1]) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_180_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_180_globalState_).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_180_globalState_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_180_globalState_).f5)[0]) == (_dafny.MultiSet([177]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_180_globalState_).f5)[1]) == (_dafny.MultiSet([177]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_180_globalState_).f5)[2]) == (_dafny.MultiSet([177]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_180_globalState_).f5)[3]) == (_dafny.MultiSet([-551, 985]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_180_globalState_.f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_globalState_).f8) == (_dafny.MultiSet([177]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_180_globalState_).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_180_globalState_.f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_180_globalState_.f11) == (_dafny.Map({_dafny.Map({_dafny.CodePoint('s'): False}): False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_180_globalState_).f12)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_180_globalState_.f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_180_globalState_).f14))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_180_globalState_).f15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_180_globalState_).f16))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_180_globalState_).f17))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_180_globalState_).f18))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_180_globalState_.f19))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_180_globalState_).f20))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_180_globalState_).f21)[0]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_180_globalState_).f21)[1]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_180_globalState_).f21)[2]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_180_globalState_).f21)[3]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_180_globalState_).f21)[4]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_180_globalState_).f21)[5]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_180_globalState_).f21)[6]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_180_globalState_).f21)[7]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_180_globalState_).f21)[8]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_180_globalState_).f21)[9]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_180_globalState_).f21)[10]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_180_globalState_).f21)[11]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_180_globalState_).f21)[12]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_180_globalState_).f21)[13]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_180_globalState_).f21)[14]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_180_globalState_).f21)[15]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_180_globalState_).f21)[16]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_180_globalState_).f21)[17]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_180_globalState_.f22))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_180_globalState_).f23))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_180_globalState_.f24) == (_dafny.Set({667}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_globalState_).f25) == (_dafny.SeqWithoutIsStrInference([False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_180_globalState_).f26))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_180_globalState_).f27))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_180_globalState_).f28))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_v20_) == (_dafny.Map({1: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_182_v21_).cf1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_182_v21_).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_182_v21_).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_182_v21_).cf4)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_182_v21_).cf4)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_252_v54_) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_318_i4_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_329_v100_) == (_dafny.Map({-551: _dafny.Map({_dafny.CodePoint('w'): -551, _dafny.CodePoint('r'): -551, _dafny.CodePoint('o'): -551, _dafny.CodePoint('h'): -551, _dafny.CodePoint('t'): -551, _dafny.CodePoint('a'): -551, _dafny.CodePoint('s'): -551, _dafny.CodePoint('v'): -551})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_331_v101_).cf44) == (_dafny.MultiSet([2]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_332_v102_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_332_v102_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_332_v102_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_405_v161_).cf26).cf25))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(int(0), int(0), int(0), _dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D0_DC1)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D0_DC2)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D0_DC3)

class D0_DC1(D0, NamedTuple('DC1', [('cf1', Any), ('cf2', Any), ('cf3', Any), ('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf1)}, {_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)}, {_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf1 == __o.cf1 and self.cf2 == __o.cf2 and self.cf3 == __o.cf3 and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC2(D0, NamedTuple('DC2', [('cf5', Any), ('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2({_dafny.string_of(self.cf5)}, {_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2) and self.cf5 == __o.cf5 and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC0(D0, NamedTuple('DC0', [('cf0', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC0({self.cf0.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC0) and self.cf0 == __o.cf0
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC3(D0, NamedTuple('DC3', [('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC3({_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC3) and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()


class D1:
    @classmethod
    def default(cls, ):
        return lambda: D1_DC5(int(0), False, False, _dafny.Map({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D1_DC5)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D1_DC6)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)

class D1_DC5(D1, NamedTuple('DC5', [('cf9', Any), ('cf10', Any), ('cf11', Any), ('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf9)}, {_dafny.string_of(self.cf10)}, {_dafny.string_of(self.cf11)}, {_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf9 == __o.cf9 and self.cf10 == __o.cf10 and self.cf11 == __o.cf11 and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC6(D1, NamedTuple('DC6', [('cf13', Any), ('cf14', Any), ('cf15', Any), ('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC6({_dafny.string_of(self.cf13)}, {_dafny.string_of(self.cf14)}, {_dafny.string_of(self.cf15)}, {_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC6) and self.cf13 == __o.cf13 and self.cf14 == __o.cf14 and self.cf15 == __o.cf15 and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC4(D1, NamedTuple('DC4', [('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC8(False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D2_DC8)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D2_DC9)

class D2_DC8(D2, NamedTuple('DC8', [('cf18', Any), ('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC8({_dafny.string_of(self.cf18)}, {_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC8) and self.cf18 == __o.cf18 and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC7(D2, NamedTuple('DC7', [('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC9(D2, NamedTuple('DC9', [('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC9({_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC9) and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC11(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D3_DC11)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D3_DC12)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D3_DC10)

class D3_DC11(D3, NamedTuple('DC11', [('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC11({_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC11) and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC12(D3, NamedTuple('DC12', [('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC12({_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC12) and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC10(D3, NamedTuple('DC10', [('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC10({_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC10) and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC14(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D4_DC14)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D4_DC13)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D4_DC15)

class D4_DC14(D4, NamedTuple('DC14', [('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC14({_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC14) and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC13(D4, NamedTuple('DC13', [('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC13({_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC13) and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC15(D4, NamedTuple('DC15', [('cf26', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC15({_dafny.string_of(self.cf26)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC15) and self.cf26 == __o.cf26
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D5_DC16)

class D5_DC16(D5, NamedTuple('DC16', [('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC16({_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC16) and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D6_DC17)

class D6_DC17(D6, NamedTuple('DC17', [('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC17({_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC17) and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC19()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D7_DC19)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D7_DC20)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D7_DC18)

class D7_DC19(D7, NamedTuple('DC19', [])):
    def __dafnystr__(self) -> str:
        return f'D7.DC19'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC19)
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC20(D7, NamedTuple('DC20', [('cf30', Any), ('cf31', Any), ('cf32', Any), ('cf33', Any), ('cf34', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC20({_dafny.string_of(self.cf30)}, {_dafny.string_of(self.cf31)}, {_dafny.string_of(self.cf32)}, {_dafny.string_of(self.cf33)}, {_dafny.string_of(self.cf34)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC20) and self.cf30 == __o.cf30 and self.cf31 == __o.cf31 and self.cf32 == __o.cf32 and self.cf33 == __o.cf33 and self.cf34 == __o.cf34
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC18(D7, NamedTuple('DC18', [('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC18({_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC18) and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC22(int(0), False, _dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D8_DC22)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D8_DC23)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D8_DC21)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D8_DC24)

class D8_DC22(D8, NamedTuple('DC22', [('cf36', Any), ('cf37', Any), ('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC22({_dafny.string_of(self.cf36)}, {_dafny.string_of(self.cf37)}, {_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC22) and self.cf36 == __o.cf36 and self.cf37 == __o.cf37 and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC23(D8, NamedTuple('DC23', [('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC23({_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC23) and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC21(D8, NamedTuple('DC21', [('cf35', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC21({_dafny.string_of(self.cf35)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC21) and self.cf35 == __o.cf35
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC24(D8, NamedTuple('DC24', [('cf40', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC24({_dafny.string_of(self.cf40)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC24) and self.cf40 == __o.cf40
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D9_DC25)

class D9_DC25(D9, NamedTuple('DC25', [('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC25({_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC25) and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D10_DC26)

class D10_DC26(D10, NamedTuple('DC26', [('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC26({_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC26) and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D11_DC27)

class D11_DC27(D11, NamedTuple('DC27', [('cf43', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC27({_dafny.string_of(self.cf43)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC27) and self.cf43 == __o.cf43
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC29(_dafny.Seq({}), int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D12_DC29)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D12_DC30)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D12_DC31)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D12_DC28)

class D12_DC29(D12, NamedTuple('DC29', [('cf45', Any), ('cf46', Any), ('cf47', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC29({_dafny.string_of(self.cf45)}, {_dafny.string_of(self.cf46)}, {_dafny.string_of(self.cf47)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC29) and self.cf45 == __o.cf45 and self.cf46 == __o.cf46 and self.cf47 == __o.cf47
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC30(D12, NamedTuple('DC30', [('cf48', Any), ('cf49', Any), ('cf50', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC30({self.cf48.VerbatimString(True)}, {_dafny.string_of(self.cf49)}, {_dafny.string_of(self.cf50)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC30) and self.cf48 == __o.cf48 and self.cf49 == __o.cf49 and self.cf50 == __o.cf50
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC31(D12, NamedTuple('DC31', [('cf51', Any), ('cf52', Any), ('cf53', Any), ('cf54', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC31({_dafny.string_of(self.cf51)}, {_dafny.string_of(self.cf52)}, {_dafny.string_of(self.cf53)}, {_dafny.string_of(self.cf54)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC31) and self.cf51 == __o.cf51 and self.cf52 == __o.cf52 and self.cf53 == __o.cf53 and self.cf54 == __o.cf54
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC28(D12, NamedTuple('DC28', [('cf44', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC28({_dafny.string_of(self.cf44)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC28) and self.cf44 == __o.cf44
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D13_DC32)

class D13_DC32(D13, NamedTuple('DC32', [('cf55', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC32({_dafny.string_of(self.cf55)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC32) and self.cf55 == __o.cf55
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: D14_DC34(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D14_DC34)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D14_DC35)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D14_DC36)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D14_DC33)

class D14_DC34(D14, NamedTuple('DC34', [('cf57', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC34({_dafny.string_of(self.cf57)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC34) and self.cf57 == __o.cf57
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC35(D14, NamedTuple('DC35', [('cf58', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC35({_dafny.string_of(self.cf58)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC35) and self.cf58 == __o.cf58
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC36(D14, NamedTuple('DC36', [('cf59', Any), ('cf60', Any), ('cf61', Any), ('cf62', Any), ('cf63', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC36({_dafny.string_of(self.cf59)}, {_dafny.string_of(self.cf60)}, {_dafny.string_of(self.cf61)}, {_dafny.string_of(self.cf62)}, {_dafny.string_of(self.cf63)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC36) and self.cf59 == __o.cf59 and self.cf60 == __o.cf60 and self.cf61 == __o.cf61 and self.cf62 == __o.cf62 and self.cf63 == __o.cf63
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC33(D14, NamedTuple('DC33', [('cf56', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC33({_dafny.string_of(self.cf56)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC33) and self.cf56 == __o.cf56
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: D15_DC38()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D15_DC38)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D15_DC37)

class D15_DC38(D15, NamedTuple('DC38', [])):
    def __dafnystr__(self) -> str:
        return f'D15.DC38'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC38)
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC37(D15, NamedTuple('DC37', [('cf64', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC37({_dafny.string_of(self.cf64)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC37) and self.cf64 == __o.cf64
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    def m0(self, p0, p1, p2, p3, globalState):
        pass

    def m1(self, p0, p1, p2, globalState):
        pass


class T1:
    pass

class GlobalState:
    def  __init__(self):
        self.f6: _dafny.Array = _dafny.Array(None, 0)
        self.f7: bool = False
        self.f10: int = int(0)
        self.f11: _dafny.Map = _dafny.Map({})
        self.f13: bool = False
        self.f19: str = _dafny.CodePoint('D')
        self.f22: bool = False
        self.f24: _dafny.Set = _dafny.Set({})
        self._f0: _dafny.Map = _dafny.Map({})
        self._f1: _dafny.Array = _dafny.Array(None, 0)
        self._f2: bool = False
        self._f3: str = _dafny.CodePoint('D')
        self._f4: bool = False
        self._f5: _dafny.Array = _dafny.Array(None, 0)
        self._f8: _dafny.MultiSet = _dafny.MultiSet({})
        self._f9: int = int(0)
        self._f12: _dafny.Map = _dafny.Map({})
        self._f14: bool = False
        self._f15: bool = False
        self._f16: bool = False
        self._f17: int = int(0)
        self._f18: bool = False
        self._f20: bool = False
        self._f21: _dafny.Array = _dafny.Array(None, 0)
        self._f23: bool = False
        self._f25: _dafny.Seq = _dafny.Seq({})
        self._f26: bool = False
        self._f27: int = int(0)
        self._f28: str = _dafny.CodePoint('D')
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22, f23, f24, f25, f26, f27, f28):
        (self)._f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2
        (self)._f3 = f3
        (self)._f4 = f4
        (self)._f5 = f5
        (self).f6 = f6
        (self).f7 = f7
        (self)._f8 = f8
        (self)._f9 = f9
        (self).f10 = f10
        (self).f11 = f11
        (self)._f12 = f12
        (self).f13 = f13
        (self)._f14 = f14
        (self)._f15 = f15
        (self)._f16 = f16
        (self)._f17 = f17
        (self)._f18 = f18
        (self).f19 = f19
        (self)._f20 = f20
        (self)._f21 = f21
        (self).f22 = f22
        (self)._f23 = f23
        (self).f24 = f24
        (self)._f25 = f25
        (self)._f26 = f26
        (self)._f27 = f27
        (self)._f28 = f28

    @property
    def f0(self):
        return self._f0
    @property
    def f1(self):
        return self._f1
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
    def f5(self):
        return self._f5
    @property
    def f8(self):
        return self._f8
    @property
    def f9(self):
        return self._f9
    @property
    def f12(self):
        return self._f12
    @property
    def f14(self):
        return self._f14
    @property
    def f15(self):
        return self._f15
    @property
    def f16(self):
        return self._f16
    @property
    def f17(self):
        return self._f17
    @property
    def f18(self):
        return self._f18
    @property
    def f20(self):
        return self._f20
    @property
    def f21(self):
        return self._f21
    @property
    def f23(self):
        return self._f23
    @property
    def f25(self):
        return self._f25
    @property
    def f26(self):
        return self._f26
    @property
    def f27(self):
        return self._f27
    @property
    def f28(self):
        return self._f28

class C0:
    def  __init__(self):
        self._f32: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f32):
        (self)._f32 = f32

    def fm13(self, p0, globalState):
        def lambda11_(source10_):
            if source10_.is_DC5:
                d_429___mcc_h0_ = source10_.cf9
                d_430___mcc_h1_ = source10_.cf10
                d_431___mcc_h2_ = source10_.cf11
                d_432___mcc_h3_ = source10_.cf12
                d_433_cf12_ = d_432___mcc_h3_
                d_434_cf11_ = d_431___mcc_h2_
                d_435_cf10_ = d_430___mcc_h1_
                d_436_cf9_ = d_429___mcc_h0_
                return 732
            elif source10_.is_DC6:
                d_437___mcc_h4_ = source10_.cf13
                d_438___mcc_h5_ = source10_.cf14
                d_439___mcc_h6_ = source10_.cf15
                d_440___mcc_h7_ = source10_.cf16
                d_441_cf16_ = d_440___mcc_h7_
                d_442_cf15_ = d_439___mcc_h6_
                d_443_cf14_ = d_438___mcc_h5_
                d_444_cf13_ = d_437___mcc_h4_
                return (0) - (default__.safeModuloInt(d_442_cf15_, len(_dafny.Map({d_441_cf16_: d_444_cf13_}))))
            elif True:
                d_445___mcc_h8_ = source10_.cf8
                d_446_cf8_ = d_445___mcc_h8_
                return default__.safeDivisionInt(632, len(_dafny.Set({339})))

        return (0) - (lambda11_(D1_DC6((D0_DC2(891, (self).f32)).cf5, _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(self).f32]), _dafny.SeqWithoutIsStrInference([not((self).f32), (self).f32])]), 124, (self).f32)))

    @property
    def f32(self):
        return self._f32

class C1(T0):
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    def ctor__(self):
        pass
        pass

    def fm1(self, p0, p1, p2, p3, globalState):
        return _dafny.Map({_dafny.MultiSet([not(False)]): (_dafny.Map({-535: not(True)})) not in (_dafny.SeqWithoutIsStrInference([_dafny.Map({len(_dafny.Map({False: (_dafny.MultiSet([True, False, False, True, True])).cardinality})): False}) for d_447_i0_ in range(default__.abs(946))]))})

    def fm2(self, globalState):
        return (0) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_448_i0_ in range(default__.abs(778))])), default__.safeDivisionInt(-566, -96), (964) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v")))), (68) * (-443)])))

    def fm21(self, p0, globalState):
        return (_dafny.Set({(_dafny.MultiSet([-402])).cardinality, -703, len(_dafny.SeqWithoutIsStrInference([False])), (_dafny.MultiSet([True, True])).cardinality})).isdisjoint(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g"))), len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "thqpyy"))): len(_dafny.SeqWithoutIsStrInference([False, True]))}))}))

    def m0(self, p0, p1, p2, p3, globalState):
        if p1:
            d_449_v0_: _dafny.Array
            def lambda12_(d_450_i0_):
                return (d_450_i0_) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ecpxkbg"))))

            init6_ = lambda12_
            nw94_ = _dafny.Array(None, 13)
            for i0_6_ in range(nw94_.length(0)):
                nw94_[i0_6_] = init6_(i0_6_)
            d_449_v0_ = nw94_
            index88_ = default__.safeIndex(510, (d_449_v0_).length(0))
            (d_449_v0_)[index88_] = p0
            d_451_v1_: _dafny.MultiSet
            d_451_v1_ = _dafny.MultiSet([-622, p2, p0, p2])
            index89_ = default__.safeIndex(665, (d_449_v0_).length(0))
            (d_449_v0_)[index89_] = ((d_451_v1_)[p0] if (p0) in (d_451_v1_) else p0)
            d_452_v2_: _dafny.Array
            nw95_ = _dafny.Array(None, 2)
            nw95_[int(0)] = (p1) == (p1)
            nw95_[int(1)] = p1
            d_452_v2_ = nw95_
            index90_ = default__.safeIndex(858, (d_452_v2_).length(0))
            (d_452_v2_)[index90_] = p3
            d_453_v3_: _dafny.Seq
            d_453_v3_ = _dafny.SeqWithoutIsStrInference([p0, p0])
            d_454_v4_: _dafny.Set
            d_454_v4_ = _dafny.Set({len(d_453_v3_), p2})
            d_455_v5_: _dafny.Seq
            d_455_v5_ = _dafny.SeqWithoutIsStrInference([p0, p2, len(_dafny.Map({p3: False})), (0) - (len(d_454_v4_))])
            d_456_v6_: _dafny.Set
            d_456_v6_ = _dafny.Set({p1, not(False), p1, p3})
            index91_ = default__.safeIndex(510, (d_449_v0_).length(0))
            index92_ = default__.safeIndex(665, (d_449_v0_).length(0))
            index93_ = default__.safeIndex(858, (d_452_v2_).length(0))
            rhs82_ = p0
            rhs83_ = default__.fm0(p0, globalState)
            rhs84_ = (p0 if p1 else (d_453_v3_)[default__.safeIndex(len(d_456_v6_), len(d_453_v3_))])
            rhs85_ = p1
            lhs68_ = globalState
            lhs69_ = d_449_v0_
            lhs70_ = default__.safeIndex(510, (d_449_v0_).length(0))
            lhs71_ = d_449_v0_
            lhs72_ = default__.safeIndex(665, (d_449_v0_).length(0))
            lhs73_ = d_452_v2_
            lhs74_ = default__.safeIndex(858, (d_452_v2_).length(0))
            lhs68_.f10 = rhs82_
            lhs69_[lhs70_] = rhs83_
            lhs71_[lhs72_] = rhs84_
            lhs73_[lhs74_] = rhs85_
            d_457_v7_: _dafny.Map
            d_457_v7_ = _dafny.Map({(d_452_v2_)[default__.safeIndex(858, (d_452_v2_).length(0))]: p3})
            d_458_v8_: _dafny.Seq
            d_458_v8_ = _dafny.SeqWithoutIsStrInference([((d_457_v7_)[p3] if (p3) in (d_457_v7_) else True), not (p1) or (p3), p1])
            index94_ = default__.safeIndex(510, (d_449_v0_).length(0))
            rhs86_ = (default__.safeDivisionInt(p0, default__.fm0(p0, globalState))) + (p0)
            rhs87_ = (d_458_v8_) + ((d_458_v8_) + (d_458_v8_))
            rhs88_ = (d_449_v0_ if False else d_449_v0_)
            lhs75_ = d_449_v0_
            lhs76_ = default__.safeIndex(510, (d_449_v0_).length(0))
            lhs75_[lhs76_] = rhs86_
            d_458_v8_ = rhs87_
            d_449_v0_ = rhs88_
            if (self).fm21(p0, globalState):
                d_459_v9_: str
                d_459_v9_ = _dafny.CodePoint('w')
                (globalState).f19 = d_459_v9_
                (globalState).f13 = True
                d_460_v10_: C0
                nw96_ = C0()
                nw96_.ctor__(p3)
                d_460_v10_ = nw96_
                (globalState).f22 = p3
                index95_ = default__.safeIndex(858, (d_452_v2_).length(0))
                (d_452_v2_)[index95_] = (self).fm21((0) - (p2), globalState)
            elif True:
                d_458_v8_ = (_dafny.SeqWithoutIsStrInference([False, (self).fm21((0) - ((d_449_v0_)[default__.safeIndex(510, (d_449_v0_).length(0))]), globalState), p1])) + ((default__.fm22((d_452_v2_)[default__.safeIndex(858, (d_452_v2_).length(0))], globalState)) + (d_458_v8_))
                d_461_v11_: _dafny.Seq
                d_461_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xfsnymi"))
                d_462_v12_: D0
                d_462_v12_ = D0_DC0(d_461_v11_)
                (globalState).f10 = len((d_462_v12_).cf0)
                (globalState).f13 = p1
                nw97_ = _dafny.Array(int(0), 9)
                d_449_v0_ = nw97_
                (globalState).f7 = (p1) in ((d_458_v8_) + (d_458_v8_))
            index96_ = default__.safeIndex(510, (d_449_v0_).length(0))
            (d_449_v0_)[index96_] = ((0) - (p0)) - (p2)
            (globalState).f13 = p1
        elif True:
            d_463_v13_: _dafny.MultiSet
            d_463_v13_ = _dafny.MultiSet([p1, p1, p3, p1, p1])
            d_464_v14_: _dafny.Seq
            d_464_v14_ = _dafny.SeqWithoutIsStrInference([False, (self).fm21(p0, globalState), p1, p1])
            d_465_v15_: _dafny.Set
            d_465_v15_ = _dafny.Set({p1})
            d_466_v16_: _dafny.Map
            d_466_v16_ = _dafny.Map({p0: p0})
            d_467_v17_: _dafny.Seq
            d_467_v17_ = _dafny.SeqWithoutIsStrInference([d_464_v14_])
            d_468_v18_: D1
            d_468_v18_ = D1_DC6(p2, d_467_v17_, p2, p3)
            d_469_v19_: _dafny.Seq
            d_469_v19_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sj"))
            d_470_v20_: _dafny.Array
            nw98_ = _dafny.Array(None, 24)
            nw98_[int(0)] = d_463_v13_
            nw98_[int(1)] = d_463_v13_
            nw98_[int(2)] = d_463_v13_
            nw98_[int(3)] = (d_463_v13_ if (d_464_v14_)[default__.safeIndex(len(d_465_v15_), len(d_464_v14_))] else _dafny.MultiSet([(d_464_v14_)[default__.safeIndex(len(default__.fm23(D2_DC7(d_466_v16_), p0, (d_468_v18_).cf15, globalState)), len(d_464_v14_))], p3, False]))
            nw98_[int(4)] = d_463_v13_
            nw98_[int(5)] = (d_463_v13_ if p3 else d_463_v13_)
            nw98_[int(6)] = _dafny.MultiSet(d_464_v14_)
            nw98_[int(7)] = (d_463_v13_) | (d_463_v13_)
            nw98_[int(8)] = default__.fm24(len(d_469_v19_), p3, globalState)
            nw98_[int(9)] = (d_463_v13_) | (d_463_v13_)
            nw98_[int(10)] = d_463_v13_
            nw98_[int(11)] = d_463_v13_
            nw98_[int(12)] = (_dafny.MultiSet(d_464_v14_)) - (d_463_v13_)
            nw98_[int(13)] = d_463_v13_
            nw98_[int(14)] = d_463_v13_
            nw98_[int(15)] = d_463_v13_
            nw98_[int(16)] = d_463_v13_
            nw98_[int(17)] = d_463_v13_
            nw98_[int(18)] = _dafny.MultiSet([False, p3])
            nw98_[int(19)] = d_463_v13_
            nw98_[int(20)] = d_463_v13_
            nw98_[int(21)] = (d_463_v13_) - (d_463_v13_)
            nw98_[int(22)] = d_463_v13_
            nw98_[int(23)] = d_463_v13_
            d_470_v20_ = nw98_
            d_470_v20_ = d_470_v20_
            d_471_v21_: _dafny.Seq
            d_471_v21_ = _dafny.SeqWithoutIsStrInference([228])
            d_471_v21_ = (d_471_v21_) + (d_471_v21_)
            d_472_v22_: str
            d_472_v22_ = _dafny.CodePoint('n')
            (globalState).f19 = d_472_v22_
            if (p0) == (len(d_471_v21_)):
                d_473_v23_: _dafny.Array
                nw99_ = _dafny.Array(False, 5)
                d_473_v23_ = nw99_
                index97_ = default__.safeIndex(793, (d_473_v23_).length(0))
                (d_473_v23_)[index97_] = False
                index98_ = default__.safeIndex(793, (d_473_v23_).length(0))
                rhs89_ = p3
                rhs90_ = d_472_v22_
                rhs91_ = ((p1 if p1 else p1)) == (p1)
                lhs77_ = d_473_v23_
                lhs78_ = default__.safeIndex(793, (d_473_v23_).length(0))
                lhs79_ = globalState
                lhs77_[lhs78_] = rhs89_
                d_472_v22_ = rhs90_
                lhs79_.f22 = rhs91_
                d_474_v24_: D2
                d_474_v24_ = D2_DC7(d_466_v16_)
                d_475_v25_: _dafny.Map
                d_475_v25_ = _dafny.Map({d_473_v23_: d_469_v19_})
                d_476_v26_: _dafny.Array
                nw100_ = _dafny.Array(None, 16)
                nw100_[int(0)] = default__.fm23(d_474_v24_, (0) - (p2), p2, globalState)
                nw100_[int(1)] = _dafny.SeqWithoutIsStrInference([d_472_v22_ for d_477_i1_ in range(default__.abs(400))])
                nw100_[int(2)] = (d_469_v19_) + (d_469_v19_)
                nw100_[int(3)] = (d_469_v19_) + (_dafny.SeqWithoutIsStrInference([d_472_v22_ for d_478_i2_ in range(default__.abs(945))]))
                nw100_[int(4)] = (d_469_v19_) + (d_469_v19_)
                nw100_[int(5)] = ((d_475_v25_)[d_473_v23_] if (d_473_v23_) in (d_475_v25_) else d_469_v19_)
                nw100_[int(6)] = d_469_v19_
                nw100_[int(7)] = d_469_v19_
                nw100_[int(8)] = _dafny.SeqWithoutIsStrInference([(d_469_v19_)[default__.safeIndex(len(d_469_v19_), len(d_469_v19_))] for d_479_i3_ in range(default__.abs(-190))])
                nw100_[int(9)] = d_469_v19_
                nw100_[int(10)] = d_469_v19_
                nw100_[int(11)] = default__.fm23(d_474_v24_, -323, p2, globalState)
                nw100_[int(12)] = (d_469_v19_).set(default__.safeIndex(len(d_471_v21_), len(d_469_v19_)), d_472_v22_)
                nw100_[int(13)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ubj"))
                nw100_[int(14)] = d_469_v19_
                nw100_[int(15)] = d_469_v19_
                d_476_v26_ = nw100_
                index99_ = default__.safeIndex(923, (d_476_v26_).length(0))
                (d_476_v26_)[index99_] = d_469_v19_
                index100_ = default__.safeIndex(793, (d_473_v23_).length(0))
                index101_ = default__.safeIndex(923, (d_476_v26_).length(0))
                rhs92_ = (self).fm21(532, globalState)
                rhs93_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fye"))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uerjb"))) + (d_469_v19_))
                rhs94_ = False
                lhs80_ = d_473_v23_
                lhs81_ = default__.safeIndex(793, (d_473_v23_).length(0))
                lhs82_ = d_476_v26_
                lhs83_ = default__.safeIndex(923, (d_476_v26_).length(0))
                lhs84_ = globalState
                lhs80_[lhs81_] = rhs92_
                lhs82_[lhs83_] = rhs93_
                lhs84_.f22 = rhs94_
                (globalState).f10 = p0
                d_480_v27_: _dafny.Array
                nw101_ = _dafny.Array(_dafny.Map({}), 17)
                d_480_v27_ = nw101_
                d_481_v28_: _dafny.Map
                d_481_v28_ = _dafny.Map({p2: (d_473_v23_)[default__.safeIndex(793, (d_473_v23_).length(0))]})
                index102_ = default__.safeIndex(353, (d_480_v27_).length(0))
                (d_480_v27_)[index102_] = d_481_v28_
                d_482_v29_: _dafny.Array
                nw102_ = _dafny.Array(None, 24)
                d_482_v29_ = nw102_
                d_483_v30_: C0
                nw103_ = C0()
                nw103_.ctor__((d_473_v23_)[default__.safeIndex(793, (d_473_v23_).length(0))])
                d_483_v30_ = nw103_
                index103_ = default__.safeIndex(303, (d_482_v29_).length(0))
                (d_482_v29_)[index103_] = d_483_v30_
                d_484_v31_: _dafny.Array
                def lambda13_(d_485_p0_):
                    def lambda14_(d_486_i4_):
                        return default__.safeModuloInt(d_486_i4_, d_485_p0_)

                    return lambda14_

                init7_ = lambda13_(p0)
                nw104_ = _dafny.Array(None, 8)
                for i0_7_ in range(nw104_.length(0)):
                    nw104_[i0_7_] = init7_(i0_7_)
                d_484_v31_ = nw104_
                index104_ = default__.safeIndex(648, (d_484_v31_).length(0))
                (d_484_v31_)[index104_] = default__.safeDivisionInt(p2, p0)
                d_487_v32_: D3
                d_487_v32_ = D3_DC10(_dafny.MultiSet(d_464_v14_))
                d_488_v33_: _dafny.Map
                d_488_v33_ = _dafny.Map({(d_476_v26_)[default__.safeIndex(923, (d_476_v26_).length(0))]: d_484_v31_})
                d_489_v34_: _dafny.Array
                d_489_v34_ = d_484_v31_
                d_490_v35_: _dafny.Seq
                d_490_v35_ = _dafny.SeqWithoutIsStrInference([d_484_v31_])
                d_491_v37_: _dafny.Array
                nw105_ = _dafny.Array(None, 14)
                nw105_[int(0)] = d_484_v31_
                nw105_[int(1)] = d_484_v31_
                nw105_[int(2)] = d_484_v31_
                nw105_[int(3)] = d_484_v31_
                nw105_[int(4)] = d_484_v31_
                nw105_[int(5)] = (((d_488_v33_)[d_469_v19_] if (d_469_v19_) in (d_488_v33_) else d_484_v31_) if (self).fm21(p2, globalState) else d_484_v31_)
                nw105_[int(6)] = d_484_v31_
                nw105_[int(7)] = d_484_v31_
                nw105_[int(8)] = d_484_v31_
                nw105_[int(9)] = (d_489_v34_)
                def iife30_():
                    coll18_ = _dafny.Map()
                    compr_18_: int
                    for compr_18_ in (_dafny.SeqWithoutIsStrInference([-571])).Elements:
                        d_492_v36_: int = compr_18_
                        if (d_492_v36_) in (_dafny.SeqWithoutIsStrInference([-571])):
                            coll18_[(d_492_v36_) * ((d_471_v21_)[default__.safeIndex(p0, len(d_471_v21_))])] = p0
                    return _dafny.Map(coll18_)
                nw105_[int(10)] = (d_490_v35_)[default__.safeIndex((0) - (len(iife30_()
                )), len(d_490_v35_))]
                nw105_[int(11)] = d_484_v31_
                nw105_[int(12)] = d_484_v31_
                nw105_[int(13)] = d_484_v31_
                d_491_v37_ = nw105_
                index105_ = default__.safeIndex(353, (d_480_v27_).length(0))
                index106_ = default__.safeIndex(303, (d_482_v29_).length(0))
                index107_ = default__.safeIndex(648, (d_484_v31_).length(0))
                def iife31_():
                    coll19_ = _dafny.Map()
                    compr_19_: int
                    for compr_19_ in _dafny.IntegerRange(207, 900):
                        d_493_v38_: int = compr_19_
                        if ((207) <= (d_493_v38_)) and ((d_493_v38_) < (900)):
                            coll19_[(d_493_v38_) + (p0)] = (p0) == ((0) - ((0) - (len(d_471_v21_))))
                    return _dafny.Map(coll19_)
                rhs95_ = iife31_()

                rhs96_ = d_483_v30_
                rhs97_ = (p2) + (p2)
                rhs98_ = d_487_v32_
                rhs99_ = (d_491_v37_ if (p3) == (not((d_473_v23_)[default__.safeIndex(793, (d_473_v23_).length(0))])) else d_491_v37_)
                lhs85_ = d_480_v27_
                lhs86_ = default__.safeIndex(353, (d_480_v27_).length(0))
                lhs87_ = d_482_v29_
                lhs88_ = default__.safeIndex(303, (d_482_v29_).length(0))
                lhs89_ = d_484_v31_
                lhs90_ = default__.safeIndex(648, (d_484_v31_).length(0))
                lhs85_[lhs86_] = rhs95_
                lhs87_[lhs88_] = rhs96_
                lhs89_[lhs90_] = rhs97_
                d_487_v32_ = rhs98_
                d_491_v37_ = rhs99_
                d_494_v39_: _dafny.Map
                d_494_v39_ = _dafny.Map({(d_473_v23_)[default__.safeIndex(793, (d_473_v23_).length(0))]: (d_484_v31_)[default__.safeIndex(648, (d_484_v31_).length(0))]})
                d_495_v40_: _dafny.Map
                d_495_v40_ = _dafny.Map({(d_473_v23_)[default__.safeIndex(793, (d_473_v23_).length(0))]: d_494_v39_})
                index108_ = default__.safeIndex(923, (d_476_v26_).length(0))
                (d_476_v26_)[index108_] = ((d_469_v19_) + (((d_476_v26_)[default__.safeIndex(923, (d_476_v26_).length(0))]) + ((d_476_v26_)[default__.safeIndex(923, (d_476_v26_).length(0))]))).set(default__.safeIndex(len((d_495_v40_) | (d_495_v40_)), len((d_469_v19_) + (((d_476_v26_)[default__.safeIndex(923, (d_476_v26_).length(0))]) + ((d_476_v26_)[default__.safeIndex(923, (d_476_v26_).length(0))])))), d_472_v22_)
            elif True:
                d_496_v41_: _dafny.Array
                def lambda15_(d_497_v13_, d_498_p1_):
                    def lambda16_(d_499_i5_):
                        return D1_DC4(_dafny.Map({d_497_v13_: d_498_p1_}))

                    return lambda16_

                init8_ = lambda15_(d_463_v13_, p1)
                nw106_ = _dafny.Array(None, 16)
                for i0_8_ in range(nw106_.length(0)):
                    nw106_[i0_8_] = init8_(i0_8_)
                d_496_v41_ = nw106_
                d_500_v43_: _dafny.Map
                d_500_v43_ = _dafny.Map({d_463_v13_: p0})
                d_501_v44_: D1
                def iife32_():
                    coll20_ = _dafny.Map()
                    compr_20_: _dafny.MultiSet
                    for compr_20_ in (d_500_v43_).keys.Elements:
                        d_502_v42_: _dafny.MultiSet = compr_20_
                        if (d_502_v42_) in (d_500_v43_):
                            coll20_[d_502_v42_] = p3
                    return _dafny.Map(coll20_)
                d_501_v44_ = D1_DC4(iife32_()
)
                index109_ = default__.safeIndex(81, (d_496_v41_).length(0))
                (d_496_v41_)[index109_] = d_501_v44_
                d_503_v45_: _dafny.Map
                d_503_v45_ = _dafny.Map({d_463_v13_: p3})
                pat_let_tv6_ = p1
                pat_let_tv7_ = p1
                index110_ = default__.safeIndex(81, (d_496_v41_).length(0))
                def iife33_(_pat_let6_0):
                    def iife34_(d_504_dt__update__tmp_h0_):
                        def iife35_(_pat_let7_0):
                            def iife36_(d_505_dt__update_hcf8_h0_):
                                return D1_DC4(d_505_dt__update_hcf8_h0_)
                            return iife36_(_pat_let7_0)
                        return iife35_(_dafny.Map({_dafny.MultiSet([pat_let_tv6_]): pat_let_tv7_}))
                    return iife34_(_pat_let6_0)
                (d_496_v41_)[index110_] = iife33_(D1_DC4(d_503_v45_))
                d_506_v46_: _dafny.Array
                def lambda17_(d_507_i6_):
                    return (d_507_i6_) * (-650)

                init9_ = lambda17_
                nw107_ = _dafny.Array(None, 23)
                for i0_9_ in range(nw107_.length(0)):
                    nw107_[i0_9_] = init9_(i0_9_)
                d_506_v46_ = nw107_
                index111_ = default__.safeIndex(438, (d_506_v46_).length(0))
                (d_506_v46_)[index111_] = p2
                index112_ = default__.safeIndex(438, (d_506_v46_).length(0))
                rhs100_ = p2
                rhs101_ = d_472_v22_
                rhs102_ = (0) - (p2)
                lhs91_ = d_506_v46_
                lhs92_ = default__.safeIndex(438, (d_506_v46_).length(0))
                lhs93_ = globalState
                lhs94_ = globalState
                lhs91_[lhs92_] = rhs100_
                lhs93_.f19 = rhs101_
                lhs94_.f10 = rhs102_
                d_468_v18_ = (d_468_v18_ if False else d_468_v18_)
                d_508_v47_: _dafny.Array
                def lambda18_(d_509_p1_, d_510_v19_, d_511_p3_):
                    def lambda19_(d_512_i7_):
                        return (_dafny.Map({_dafny.Map({20: d_509_p1_}): d_509_p1_})) | (_dafny.Map({_dafny.Map({len(d_510_v19_): False}): d_511_p3_}))

                    return lambda19_

                init10_ = lambda18_(p1, d_469_v19_, p3)
                nw108_ = _dafny.Array(None, 27)
                for i0_10_ in range(nw108_.length(0)):
                    nw108_[i0_10_] = init10_(i0_10_)
                d_508_v47_ = nw108_
                d_513_v48_: _dafny.Map
                d_513_v48_ = _dafny.Map({p0: p1})
                d_514_v49_: _dafny.Map
                d_514_v49_ = _dafny.Map({d_513_v48_: (self).fm21((d_506_v46_)[default__.safeIndex(438, (d_506_v46_).length(0))], globalState)})
                index113_ = default__.safeIndex(622, (d_508_v47_).length(0))
                (d_508_v47_)[index113_] = d_514_v49_
                index114_ = default__.safeIndex(438, (d_506_v46_).length(0))
                index115_ = default__.safeIndex(622, (d_508_v47_).length(0))
                rhs103_ = p0
                rhs104_ = d_514_v49_
                rhs105_ = ((d_471_v21_) + (d_471_v21_)) + (_dafny.SeqWithoutIsStrInference([208]))
                rhs106_ = p1
                lhs95_ = d_506_v46_
                lhs96_ = default__.safeIndex(438, (d_506_v46_).length(0))
                lhs97_ = d_508_v47_
                lhs98_ = default__.safeIndex(622, (d_508_v47_).length(0))
                lhs99_ = globalState
                lhs95_[lhs96_] = rhs103_
                lhs97_[lhs98_] = rhs104_
                d_471_v21_ = rhs105_
                lhs99_.f13 = rhs106_
                d_515_v50_: D0
                d_515_v50_ = D0_DC2(default__.fm0(p2, globalState), p1)
                d_516_v51_: _dafny.Array
                nw109_ = _dafny.Array(None, 15)
                nw109_[int(0)] = p3
                nw109_[int(1)] = p3
                nw109_[int(2)] = p3
                nw109_[int(3)] = not(p1)
                nw109_[int(4)] = not(not((p3) not in (d_465_v15_)))
                nw109_[int(5)] = (867) <= (len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({p3: d_472_v22_})) for d_517_i8_ in range(default__.abs(-93))])))
                nw109_[int(6)] = (p0) >= (p0)
                nw109_[int(7)] = p1
                nw109_[int(8)] = p1
                nw109_[int(9)] = p3
                nw109_[int(10)] = (not(p3)) or (p1)
                nw109_[int(11)] = p1
                nw109_[int(12)] = p1
                nw109_[int(13)] = p1
                nw109_[int(14)] = (p0) < ((d_515_v50_).cf5)
                d_516_v51_ = nw109_
                index116_ = default__.safeIndex(863, (d_516_v51_).length(0))
                (d_516_v51_)[index116_] = p1
                d_518_v52_: D0
                d_518_v52_ = D0_DC1(p0, len(_dafny.Map({p3: p0})), p2, d_516_v51_)
                d_519_v53_: D0
                d_519_v53_ = D0_DC3(d_518_v52_)
                index117_ = default__.safeIndex(863, (d_516_v51_).length(0))
                index118_ = default__.safeIndex(438, (d_506_v46_).length(0))
                index119_ = default__.safeIndex(438, (d_506_v46_).length(0))
                rhs107_ = p1
                rhs108_ = (d_519_v53_ if p3 else d_519_v53_)
                rhs109_ = ((p2) * ((d_506_v46_)[default__.safeIndex(438, (d_506_v46_).length(0))])) * ((d_506_v46_)[default__.safeIndex(438, (d_506_v46_).length(0))])
                rhs110_ = p0
                lhs100_ = d_516_v51_
                lhs101_ = default__.safeIndex(863, (d_516_v51_).length(0))
                lhs102_ = d_506_v46_
                lhs103_ = default__.safeIndex(438, (d_506_v46_).length(0))
                lhs104_ = d_506_v46_
                lhs105_ = default__.safeIndex(438, (d_506_v46_).length(0))
                lhs100_[lhs101_] = rhs107_
                d_519_v53_ = rhs108_
                lhs102_[lhs103_] = rhs109_
                lhs104_[lhs105_] = rhs110_
            d_520_v54_: _dafny.Array
            def lambda20_(d_521_p3_):
                def lambda21_(d_522_i9_):
                    return d_521_p3_

                return lambda21_

            init11_ = lambda20_(p3)
            nw110_ = _dafny.Array(None, 19)
            for i0_11_ in range(nw110_.length(0)):
                nw110_[i0_11_] = init11_(i0_11_)
            d_520_v54_ = nw110_
            d_523_v55_: D0
            d_523_v55_ = D0_DC1(p2, 933, -728, d_520_v54_)
            d_524_v56_: _dafny.Seq
            d_524_v56_ = _dafny.SeqWithoutIsStrInference([d_523_v55_, d_523_v55_])
            d_525_v57_: _dafny.Map
            d_525_v57_ = _dafny.Map({p1: d_524_v56_})
            d_526_v58_: _dafny.Set
            d_526_v58_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([d_523_v55_, d_523_v55_, d_523_v55_]), (((d_525_v57_)[p3] if (p3) in (d_525_v57_) else d_524_v56_)) + (d_524_v56_)})
            d_527_v59_: _dafny.Array
            nw111_ = _dafny.Array(_dafny.Array(None, 0), 1)
            d_527_v59_ = nw111_
            d_528_v60_: _dafny.Array
            nw112_ = _dafny.Array(D0.default()(), 12)
            d_528_v60_ = nw112_
            index120_ = default__.safeIndex(921, (d_527_v59_).length(0))
            (d_527_v59_)[index120_] = d_528_v60_
            d_529_v61_: _dafny.Array
            def lambda22_(d_530_v22_):
                def lambda23_(d_531_i10_):
                    return d_530_v22_

                return lambda23_

            init12_ = lambda22_(d_472_v22_)
            nw113_ = _dafny.Array(None, 6)
            for i0_12_ in range(nw113_.length(0)):
                nw113_[i0_12_] = init12_(i0_12_)
            d_529_v61_ = nw113_
            index121_ = default__.safeIndex(837, (d_529_v61_).length(0))
            (d_529_v61_)[index121_] = _dafny.CodePoint('e')
            index122_ = default__.safeIndex(921, (d_527_v59_).length(0))
            index123_ = default__.safeIndex(837, (d_529_v61_).length(0))
            rhs111_ = d_526_v58_
            rhs112_ = d_528_v60_
            rhs113_ = (p3) not in ((d_464_v14_) + ((d_464_v14_).set(default__.safeIndex(len(d_469_v19_), len(d_464_v14_)), p1)))
            rhs114_ = p2
            rhs115_ = d_472_v22_
            lhs106_ = d_527_v59_
            lhs107_ = default__.safeIndex(921, (d_527_v59_).length(0))
            lhs108_ = globalState
            lhs109_ = globalState
            lhs110_ = d_529_v61_
            lhs111_ = default__.safeIndex(837, (d_529_v61_).length(0))
            d_526_v58_ = rhs111_
            lhs106_[lhs107_] = rhs112_
            lhs108_.f13 = rhs113_
            lhs109_.f10 = rhs114_
            lhs110_[lhs111_] = rhs115_
        hi1_ = default__.safeModuloInt(61, 131)
        for d_532_i11_ in range(461, hi1_):
            d_533_v62_: _dafny.Map
            d_533_v62_ = _dafny.Map({p1: p0})
            d_533_v62_ = d_533_v62_
            d_534_v63_: _dafny.Seq
            d_534_v63_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hrq"))
            d_535_v64_: _dafny.Seq
            d_535_v64_ = _dafny.SeqWithoutIsStrInference([len(d_534_v63_)])
            d_536_v65_: _dafny.Set
            d_536_v65_ = _dafny.Set({len(d_535_v64_)})
            d_537_v66_: C0
            nw114_ = C0()
            nw114_.ctor__(not(p1))
            d_537_v66_ = nw114_
            d_538_v67_: _dafny.Seq
            d_538_v67_ = _dafny.SeqWithoutIsStrInference([d_537_v66_])
            d_539_v68_: str
            d_539_v68_ = _dafny.CodePoint('v')
            d_540_v69_: _dafny.Set
            d_540_v69_ = _dafny.Set({_dafny.CodePoint('f'), d_539_v68_, d_539_v68_, d_539_v68_})
            d_541_v70_: _dafny.MultiSet
            d_541_v70_ = _dafny.MultiSet([(d_538_v67_).set(default__.safeIndex(len(d_540_v69_), len(d_538_v67_)), d_537_v66_), d_538_v67_])
            d_542_v71_: _dafny.MultiSet
            d_542_v71_ = _dafny.MultiSet([p2])
            d_543_v72_: _dafny.Array
            nw115_ = _dafny.Array(None, 24)
            nw115_[int(0)] = p0
            nw115_[int(1)] = p0
            nw115_[int(2)] = d_532_i11_
            nw115_[int(3)] = p0
            nw115_[int(4)] = (p0) * (len(d_536_v65_))
            nw115_[int(5)] = p2
            nw115_[int(6)] = d_532_i11_
            nw115_[int(7)] = len(d_534_v63_)
            nw115_[int(8)] = d_532_i11_
            nw115_[int(9)] = len(d_534_v63_)
            nw115_[int(10)] = len(default__.fm22(p3, globalState))
            nw115_[int(11)] = default__.safeModuloInt(len(d_534_v63_), d_532_i11_)
            nw115_[int(12)] = d_532_i11_
            nw115_[int(13)] = (self).fm2(globalState)
            nw115_[int(14)] = ((d_541_v70_)[d_538_v67_] if (d_538_v67_) in (d_541_v70_) else default__.fm0((d_542_v71_).cardinality, globalState))
            nw115_[int(15)] = p2
            nw115_[int(16)] = p0
            nw115_[int(17)] = p2
            nw115_[int(18)] = d_532_i11_
            nw115_[int(19)] = p0
            nw115_[int(20)] = default__.safeModuloInt((d_537_v66_).fm13(p2, globalState), len(_dafny.SeqWithoutIsStrInference([d_539_v68_, d_539_v68_, d_539_v68_])))
            nw115_[int(21)] = len(_dafny.Map({(d_537_v66_).f32: d_539_v68_}))
            nw115_[int(22)] = default__.safeDivisionInt((d_537_v66_).fm13(p0, globalState), default__.fm0(d_532_i11_, globalState))
            nw115_[int(23)] = p0
            d_543_v72_ = nw115_
            d_544_v73_: _dafny.Set
            d_544_v73_ = _dafny.Set({(d_537_v66_).f32})
            index124_ = default__.safeIndex(573, (d_543_v72_).length(0))
            (d_543_v72_)[index124_] = default__.safeModuloInt(len(d_544_v73_), p2)
            index125_ = default__.safeIndex(573, (d_543_v72_).length(0))
            (d_543_v72_)[index125_] = p2
            d_545_v74_: _dafny.Array
            def lambda24_(d_546_v66_, d_547_i11_, d_548_v72_):
                def lambda25_(d_549_i12_):
                    return (len(_dafny.Map({(d_546_v66_).f32: D2_DC7(_dafny.Map({d_547_i11_: (d_548_v72_)[default__.safeIndex(573, (d_548_v72_).length(0))]}))}))) > (d_547_i11_)

                return lambda25_

            init13_ = lambda24_(d_537_v66_, d_532_i11_, d_543_v72_)
            nw116_ = _dafny.Array(None, 4)
            for i0_13_ in range(nw116_.length(0)):
                nw116_[i0_13_] = init13_(i0_13_)
            d_545_v74_ = nw116_
            index126_ = default__.safeIndex(971, (d_545_v74_).length(0))
            (d_545_v74_)[index126_] = p1
            index127_ = default__.safeIndex(573, (d_543_v72_).length(0))
            index128_ = default__.safeIndex(573, (d_543_v72_).length(0))
            index129_ = default__.safeIndex(971, (d_545_v74_).length(0))
            rhs116_ = (0) - (default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([(d_535_v64_)[default__.safeIndex(d_532_i11_, len(d_535_v64_))], (0) - (d_532_i11_), len(d_534_v63_)])), p2))
            rhs117_ = d_532_i11_
            rhs118_ = (d_532_i11_) >= (d_532_i11_)
            lhs112_ = d_543_v72_
            lhs113_ = default__.safeIndex(573, (d_543_v72_).length(0))
            lhs114_ = d_543_v72_
            lhs115_ = default__.safeIndex(573, (d_543_v72_).length(0))
            lhs116_ = d_545_v74_
            lhs117_ = default__.safeIndex(971, (d_545_v74_).length(0))
            lhs112_[lhs113_] = rhs116_
            lhs114_[lhs115_] = rhs117_
            lhs116_[lhs117_] = rhs118_
            d_550_v75_: C0
            nw117_ = C0()
            nw117_.ctor__((d_545_v74_)[default__.safeIndex(971, (d_545_v74_).length(0))])
            d_550_v75_ = nw117_
        d_551_v76_: _dafny.Array
        nw118_ = _dafny.Array(_dafny.Array(None, 0), 17)
        d_551_v76_ = nw118_
        d_552_v77_: _dafny.Seq
        d_552_v77_ = _dafny.SeqWithoutIsStrInference([p2])
        d_553_v78_: _dafny.Seq
        d_553_v78_ = _dafny.SeqWithoutIsStrInference([(self).fm21(p0, globalState)])
        d_554_v79_: _dafny.MultiSet
        d_554_v79_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_555_i13_ in range(default__.abs(28))]))])
        d_556_v80_: _dafny.Seq
        d_556_v80_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cxrrfwi"))
        d_557_v81_: _dafny.Map
        d_557_v81_ = _dafny.Map({(self).fm2(globalState): d_556_v80_})
        d_558_v82_: _dafny.Array
        nw119_ = _dafny.Array(None, 10)
        nw119_[int(0)] = p2
        nw119_[int(1)] = p2
        nw119_[int(2)] = len(d_552_v77_)
        nw119_[int(3)] = 380
        nw119_[int(4)] = len(_dafny.Map({p0: len(d_553_v78_)}))
        nw119_[int(5)] = default__.fm0(p0, globalState)
        nw119_[int(6)] = p2
        nw119_[int(7)] = 411
        nw119_[int(8)] = ((d_554_v79_)[p0] if (p0) in (d_554_v79_) else len(((d_557_v81_)[p0] if (p0) in (d_557_v81_) else d_556_v80_)))
        nw119_[int(9)] = p0
        d_558_v82_ = nw119_
        d_559_v83_: _dafny.Array
        nw120_ = _dafny.Array(None, 3)
        nw120_[int(0)] = d_558_v82_
        nw120_[int(1)] = d_558_v82_
        nw120_[int(2)] = d_558_v82_
        d_559_v83_ = nw120_
        index130_ = default__.safeIndex(692, (d_551_v76_).length(0))
        (d_551_v76_)[index130_] = d_559_v83_
        d_560_v84_: _dafny.Map
        d_560_v84_ = _dafny.Map({p3: d_559_v83_})
        index131_ = default__.safeIndex(692, (d_551_v76_).length(0))
        (d_551_v76_)[index131_] = ((d_560_v84_)[p1] if (p1) in (d_560_v84_) else d_559_v83_)
        d_561_v85_: _dafny.Seq
        d_561_v85_ = _dafny.SeqWithoutIsStrInference([d_553_v78_, d_553_v78_])
        d_562_v86_: D1
        d_562_v86_ = D1_DC6(p2, d_561_v85_, p0, p3)
        d_563_v87_: _dafny.Map
        d_563_v87_ = _dafny.Map({True: p3})
        d_564_v88_: _dafny.Map
        d_564_v88_ = _dafny.Map({((d_563_v87_)[p3] if (p3) in (d_563_v87_) else p1): p1})
        d_565_v89_: D0
        d_565_v89_ = D0_DC2(p0, p1)
        d_566_v90_: _dafny.Array
        nw121_ = _dafny.Array(None, 19)
        nw121_[int(0)] = p3
        nw121_[int(1)] = p1
        nw121_[int(2)] = p1
        nw121_[int(3)] = p1
        nw121_[int(4)] = (d_562_v86_).cf16
        nw121_[int(5)] = p3
        nw121_[int(6)] = p3
        nw121_[int(7)] = ((d_563_v87_)[((d_564_v88_)[p3] if (p3) in (d_564_v88_) else p1)] if (((d_564_v88_)[p3] if (p3) in (d_564_v88_) else p1)) in (d_563_v87_) else p3)
        nw121_[int(8)] = p1
        nw121_[int(9)] = p3
        nw121_[int(10)] = p1
        nw121_[int(11)] = (d_565_v89_).cf6
        nw121_[int(12)] = (d_553_v78_)[default__.safeIndex(p0, len(d_553_v78_))]
        nw121_[int(13)] = p1
        nw121_[int(14)] = p1
        nw121_[int(15)] = p3
        nw121_[int(16)] = p3
        nw121_[int(17)] = not(p1)
        nw121_[int(18)] = False
        d_566_v90_ = nw121_
        d_567_v91_: _dafny.Set
        d_567_v91_ = _dafny.Set({d_566_v90_})
        d_568_v92_: _dafny.Set
        d_568_v92_ = _dafny.Set({(d_552_v77_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cqty"))), len(d_552_v77_))]})
        d_569_v93_: _dafny.Seq
        d_569_v93_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({d_566_v90_, d_566_v90_, d_566_v90_}), d_567_v91_])
        rhs119_ = len(((d_568_v92_) | (_dafny.Set({(0) - (p2)}))).intersection((d_568_v92_ if p1 else d_568_v92_)))
        rhs120_ = (d_569_v93_)[default__.safeIndex((p0) * (p2), len(d_569_v93_))]
        rhs121_ = p0
        lhs118_ = globalState
        lhs119_ = globalState
        lhs118_.f10 = rhs119_
        d_567_v91_ = rhs120_
        lhs119_.f10 = rhs121_
        (globalState).f13 = p3
        d_570_v94_: str
        d_570_v94_ = _dafny.CodePoint('y')
        d_571_v95_: D4
        d_571_v95_ = D4_DC13(d_570_v94_)
        pat_let_tv8_ = d_570_v94_
        pat_let_tv9_ = d_570_v94_
        pat_let_tv10_ = d_570_v94_
        def lambda26_(source11_):
            if source11_.is_DC14:
                d_572___mcc_h0_ = source11_.cf25
                d_573_cf25_ = d_572___mcc_h0_
                return pat_let_tv8_
            elif source11_.is_DC13:
                d_574___mcc_h1_ = source11_.cf24
                d_575_cf24_ = d_574___mcc_h1_
                return pat_let_tv9_
            elif True:
                d_576___mcc_h2_ = source11_.cf26
                d_577_cf26_ = d_576___mcc_h2_
                return pat_let_tv10_

        rhs122_ = p0
        rhs123_ = p2
        rhs124_ = (d_570_v94_ if p1 else d_570_v94_)
        rhs125_ = lambda26_(d_571_v95_)
        lhs120_ = globalState
        lhs121_ = globalState
        lhs122_ = globalState
        lhs123_ = globalState
        lhs120_.f10 = rhs122_
        lhs121_.f10 = rhs123_
        lhs122_.f19 = rhs124_
        lhs123_.f19 = rhs125_

    def m1(self, p0, p1, p2, globalState):
        r0: D0 = D0.default()()
        r1: bool = False
        d_578_v0_: _dafny.Seq
        d_578_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "btqf"))
        d_578_v0_ = d_578_v0_
        if p2:
            (globalState).f13 = not(p2)
            d_579_v1_: _dafny.Map
            d_579_v1_ = _dafny.Map({408: not(p2)})
            d_580_v2_: _dafny.Array
            nw122_ = _dafny.Array(None, 17)
            nw122_[int(0)] = (0) - (p1)
            nw122_[int(1)] = (p1) + (p1)
            nw122_[int(2)] = len((d_578_v0_) + (d_578_v0_))
            nw122_[int(3)] = (0) - ((0) - ((0) - ((p1) * (len(d_578_v0_)))))
            nw122_[int(4)] = len(p0)
            nw122_[int(5)] = (0) - ((p1) * (p1))
            nw122_[int(6)] = p1
            nw122_[int(7)] = p1
            nw122_[int(8)] = p1
            nw122_[int(9)] = p1
            nw122_[int(10)] = p1
            nw122_[int(11)] = p1
            nw122_[int(12)] = len(d_579_v1_)
            nw122_[int(13)] = p1
            nw122_[int(14)] = p1
            nw122_[int(15)] = (0) - (p1)
            nw122_[int(16)] = len((d_578_v0_) + (d_578_v0_))
            d_580_v2_ = nw122_
            d_580_v2_ = (d_580_v2_ if not(p2) else d_580_v2_)
            d_581_v3_: _dafny.Array
            nw123_ = _dafny.Array(_dafny.Array(None, 0), 12)
            d_581_v3_ = nw123_
            d_582_v4_: _dafny.Seq
            d_582_v4_ = _dafny.SeqWithoutIsStrInference([p2, p2, p2])
            rhs126_ = (p1 if not(p2) else (p1 if (d_582_v4_)[default__.safeIndex(p1, len(d_582_v4_))] else -576))
            rhs127_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ocl"))
            rhs128_ = d_581_v3_
            rhs129_ = default__.safeModuloInt(548, p1)
            rhs130_ = (self).fm21((0) - (p1), globalState)
            lhs124_ = globalState
            lhs125_ = globalState
            lhs126_ = globalState
            lhs124_.f10 = rhs126_
            d_578_v0_ = rhs127_
            d_581_v3_ = rhs128_
            lhs125_.f10 = rhs129_
            lhs126_.f13 = rhs130_
            if p2:
                d_583_v5_: _dafny.Seq
                d_583_v5_ = _dafny.SeqWithoutIsStrInference([p1, p1])
                d_583_v5_ = ((d_583_v5_) + (_dafny.SeqWithoutIsStrInference([p1, (0) - (p1), p1]))).set(default__.safeIndex(p1, len((d_583_v5_) + (_dafny.SeqWithoutIsStrInference([p1, (0) - (p1), p1])))), (p1 if p2 else p1))
                (globalState).f7 = p2
                d_584_v6_: _dafny.MultiSet
                d_584_v6_ = _dafny.MultiSet([p1, len(d_578_v0_)])
                d_585_v7_: _dafny.Map
                d_585_v7_ = _dafny.Map({d_584_v6_: (p0).set(default__.safeIndex(p1, len(p0)), (_dafny.MultiSet([True])).cardinality)})
                d_586_v8_: D1
                d_586_v8_ = D1_DC5(p1, not(p2), p2, d_585_v7_)
                d_587_v9_: _dafny.Array
                nw124_ = _dafny.Array(None, 10)
                nw124_[int(0)] = default__.fm25(p2, not(p2), p1, p2, globalState)
                nw124_[int(1)] = d_586_v8_
                nw124_[int(2)] = d_586_v8_
                nw124_[int(3)] = d_586_v8_
                nw124_[int(4)] = d_586_v8_
                nw124_[int(5)] = d_586_v8_
                nw124_[int(6)] = d_586_v8_
                nw124_[int(7)] = d_586_v8_
                nw124_[int(8)] = D1_DC5(p1, p2, True, d_585_v7_)
                nw124_[int(9)] = d_586_v8_
                d_587_v9_ = nw124_
                index132_ = default__.safeIndex(758, (d_587_v9_).length(0))
                (d_587_v9_)[index132_] = d_586_v8_
                index133_ = default__.safeIndex(758, (d_587_v9_).length(0))
                (d_587_v9_)[index133_] = d_586_v8_
                d_588_v10_: str
                out2_: str
                out2_ = (self).m6((len(d_583_v5_)) - (p1), (len(d_578_v0_)) - (p1), (len(d_582_v4_)) * (p1), default__.safeDivisionInt(p1, p1), globalState)
                d_588_v10_ = out2_
                d_589_v11_: _dafny.Set
                d_589_v11_ = _dafny.Set({True, p2, p2})
                (globalState).f13 = not (not((p2) in (d_589_v11_))) or (p2)
            elif True:
                d_590_v12_: _dafny.Map
                d_590_v12_ = _dafny.Map({p2: ((p0).set(default__.safeIndex(p1, len(p0)), p1)) + (p0)})
                d_591_v13_: _dafny.MultiSet
                d_591_v13_ = _dafny.MultiSet([p2])
                d_592_v14_: _dafny.Map
                d_592_v14_ = _dafny.Map({d_591_v13_: True})
                d_593_v15_: D1
                d_593_v15_ = D1_DC4(d_592_v14_)
                d_594_v16_: _dafny.Seq
                d_594_v16_ = _dafny.SeqWithoutIsStrInference([d_593_v15_, d_593_v15_])
                d_595_v17_: str
                d_595_v17_ = _dafny.CodePoint('k')
                d_596_v18_: _dafny.Map
                d_596_v18_ = _dafny.Map({d_580_v2_: d_595_v17_})
                d_597_v19_: _dafny.Set
                d_597_v19_ = _dafny.Set({p2, p2})
                d_598_v20_: _dafny.MultiSet
                d_598_v20_ = _dafny.MultiSet([d_595_v17_])
                d_599_v21_: _dafny.Map
                d_599_v21_ = _dafny.Map({d_598_v20_: p2})
                d_600_v22_: _dafny.MultiSet
                d_600_v22_ = _dafny.MultiSet([len(d_594_v16_), (len(d_596_v18_)) * (p1), p1, default__.safeDivisionInt(len(d_597_v19_), len(d_599_v21_)), p1])
                d_601_v23_: D0
                d_601_v23_ = D0_DC2(p1, p2)
                d_602_v24_: _dafny.Map
                d_602_v24_ = _dafny.Map({(d_601_v23_).cf6: p2})
                rhs131_ = len((d_582_v4_ if p2 else (d_582_v4_) + (default__.fm22(p2, globalState))))
                rhs132_ = default__.fm26(p1, ((d_602_v24_)[p2] if (p2) in (d_602_v24_) else p2), (d_578_v0_)[default__.safeIndex((d_600_v22_).cardinality, len(d_578_v0_))], globalState)
                rhs133_ = d_600_v22_
                rhs134_ = True
                rhs135_ = p1
                lhs127_ = globalState
                lhs128_ = globalState
                lhs129_ = globalState
                lhs127_.f10 = rhs131_
                d_590_v12_ = rhs132_
                d_600_v22_ = rhs133_
                lhs128_.f13 = rhs134_
                lhs129_.f10 = rhs135_
                d_603_v25_: _dafny.Set
                d_603_v25_ = _dafny.Set({(_dafny.SeqWithoutIsStrInference([d_595_v17_])) + ((d_578_v0_).set(default__.safeIndex((0) - (p1), len(d_578_v0_)), d_595_v17_))})
                d_604_v26_: _dafny.Map
                d_604_v26_ = _dafny.Map({d_597_v19_: p2})
                rhs136_ = default__.fm27(d_604_v26_, globalState)
                rhs137_ = p2
                lhs130_ = globalState
                d_603_v25_ = rhs136_
                lhs130_.f22 = rhs137_
                d_605_v27_: _dafny.Array
                nw125_ = _dafny.Array(None, 27)
                nw125_[int(0)] = p2
                nw125_[int(1)] = p2
                nw125_[int(2)] = p2
                nw125_[int(3)] = p2
                nw125_[int(4)] = p2
                nw125_[int(5)] = p2
                nw125_[int(6)] = p2
                nw125_[int(7)] = p2
                nw125_[int(8)] = p2
                nw125_[int(9)] = False
                nw125_[int(10)] = p2
                nw125_[int(11)] = (self).fm21(p1, globalState)
                nw125_[int(12)] = p2
                nw125_[int(13)] = p2
                nw125_[int(14)] = False
                nw125_[int(15)] = p2
                nw125_[int(16)] = p2
                nw125_[int(17)] = p2
                nw125_[int(18)] = p2
                nw125_[int(19)] = not(p2)
                nw125_[int(20)] = p2
                nw125_[int(21)] = not(False)
                nw125_[int(22)] = p2
                nw125_[int(23)] = p2
                nw125_[int(24)] = p2
                nw125_[int(25)] = (d_582_v4_)[default__.safeIndex(p1, len(d_582_v4_))]
                nw125_[int(26)] = p2
                d_605_v27_ = nw125_
                d_606_v28_: _dafny.Seq
                d_606_v28_ = _dafny.SeqWithoutIsStrInference([d_605_v27_, d_605_v27_])
                d_607_v29_: _dafny.Map
                d_607_v29_ = _dafny.Map({(d_606_v28_)[default__.safeIndex(p1, len(d_606_v28_))]: len(d_578_v0_)})
                d_608_v30_: _dafny.Map
                d_608_v30_ = _dafny.Map({d_607_v29_: (d_595_v17_) in (d_578_v0_)})
                d_608_v30_ = (d_608_v30_).set(d_607_v29_, p2)
                (globalState).f7 = p2
                index134_ = default__.safeIndex(242, (d_581_v3_).length(0))
                (d_581_v3_)[index134_] = d_605_v27_
                d_609_v31_: D4
                d_609_v31_ = D4_DC14(p2)
                index135_ = default__.safeIndex(242, (d_581_v3_).length(0))
                rhs138_ = d_605_v27_
                rhs139_ = p2
                rhs140_ = d_609_v31_
                lhs131_ = d_581_v3_
                lhs132_ = default__.safeIndex(242, (d_581_v3_).length(0))
                lhs133_ = globalState
                lhs131_[lhs132_] = rhs138_
                lhs133_.f22 = rhs139_
                d_609_v31_ = rhs140_
            (globalState).f10 = len(((d_582_v4_) + (default__.fm22(p2, globalState))) + (d_582_v4_))
        elif True:
            d_610_v32_: _dafny.Set
            d_610_v32_ = _dafny.Set({d_578_v0_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_611_i0_ in range(default__.abs(303))])})
            (globalState).f7 = (self).fm21(len(d_610_v32_), globalState)
            d_578_v0_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_612_i1_ in range(default__.abs(-309))])
            d_613_v33_: _dafny.Set
            d_613_v33_ = _dafny.Set({p1, p1, p1, p1})
            d_614_v34_: _dafny.Array
            nw126_ = _dafny.Array(None, 19)
            nw126_[int(0)] = False
            nw126_[int(1)] = False
            nw126_[int(2)] = (self).fm21(-719, globalState)
            nw126_[int(3)] = (d_613_v33_).issubset(d_613_v33_)
            nw126_[int(4)] = p2
            nw126_[int(5)] = p2
            nw126_[int(6)] = False
            nw126_[int(7)] = (len(default__.fm28(p2, globalState))) > (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iqstee"))))
            nw126_[int(8)] = not(True)
            nw126_[int(9)] = p2
            nw126_[int(10)] = p2
            nw126_[int(11)] = (p2) == (p2)
            nw126_[int(12)] = True
            nw126_[int(13)] = p2
            nw126_[int(14)] = p2
            nw126_[int(15)] = not(False)
            nw126_[int(16)] = p2
            nw126_[int(17)] = p2
            nw126_[int(18)] = p2
            d_614_v34_ = nw126_
            index136_ = default__.safeIndex(839, (d_614_v34_).length(0))
            (d_614_v34_)[index136_] = (p1) <= (p1)
            index137_ = default__.safeIndex(839, (d_614_v34_).length(0))
            (d_614_v34_)[index137_] = p2
            (globalState).f13 = (d_614_v34_)[default__.safeIndex(839, (d_614_v34_).length(0))]
            d_615_v35_: _dafny.Set
            d_615_v35_ = _dafny.Set({(d_614_v34_)[default__.safeIndex(839, (d_614_v34_).length(0))]})
            d_616_v36_: _dafny.Seq
            d_616_v36_ = _dafny.SeqWithoutIsStrInference([(d_614_v34_)[default__.safeIndex(839, (d_614_v34_).length(0))]])
            d_617_v37_: _dafny.MultiSet
            d_617_v37_ = _dafny.MultiSet([p1])
            d_618_v38_: _dafny.Set
            d_618_v38_ = _dafny.Set({not(p2), (d_615_v35_).ispropersubset(_dafny.Set({p2})), ((d_616_v36_)[default__.safeIndex(p1, len(d_616_v36_))] if not(p2) else not(p2)), (d_614_v34_)[default__.safeIndex(839, (d_614_v34_).length(0))], (d_617_v37_).ispropersubset(d_617_v37_)})
            d_615_v35_ = d_618_v38_
        d_619_v39_: _dafny.Array
        nw127_ = _dafny.Array(False, 13)
        d_619_v39_ = nw127_
        d_620_v40_: _dafny.Seq
        d_620_v40_ = _dafny.SeqWithoutIsStrInference([p2])
        d_621_v41_: D1
        d_621_v41_ = D1_DC6(p1, _dafny.SeqWithoutIsStrInference([d_620_v40_]), p1, p2)
        rhs141_ = d_619_v39_
        rhs142_ = p1
        rhs143_ = default__.fm29(globalState)
        lhs134_ = globalState
        d_619_v39_ = rhs141_
        lhs134_.f10 = rhs142_
        d_621_v41_ = rhs143_
        d_622_v42_: _dafny.Array
        nw128_ = _dafny.Array(int(0), 9)
        d_622_v42_ = nw128_
        d_622_v42_ = d_622_v42_
        d_623_v43_: _dafny.Array
        nw129_ = _dafny.Array(_dafny.Array(None, 0), 27)
        d_623_v43_ = nw129_
        index138_ = default__.safeIndex(960, (d_623_v43_).length(0))
        (d_623_v43_)[index138_] = d_619_v39_
        index139_ = default__.safeIndex(960, (d_623_v43_).length(0))
        nw130_ = _dafny.Array(False, 14)
        (d_623_v43_)[index139_] = nw130_
        (globalState).f10 = p1
        d_624_v44_: D0
        d_624_v44_ = D0_DC1(914, p1, p1, (d_623_v43_)[default__.safeIndex(960, (d_623_v43_).length(0))])
        r0 = (d_624_v44_ if not((p2) == (p2)) else d_624_v44_)
        r1 = (d_620_v40_)[default__.safeIndex(p1, len(d_620_v40_))]
        return r0, r1

    def m6(self, p0, p1, p2, p3, globalState):
        r0: str = _dafny.CodePoint('D')
        d_625_v0_: _dafny.Seq
        d_625_v0_ = _dafny.SeqWithoutIsStrInference([p2])
        d_625_v0_ = d_625_v0_
        d_626_v1_: _dafny.Array
        nw131_ = _dafny.Array(int(0), 16)
        d_626_v1_ = nw131_
        index140_ = default__.safeIndex(935, (d_626_v1_).length(0))
        (d_626_v1_)[index140_] = p2
        index141_ = default__.safeIndex(935, (d_626_v1_).length(0))
        (d_626_v1_)[index141_] = p1
        d_627_v2_: bool
        d_627_v2_ = False
        index142_ = default__.safeIndex(935, (d_626_v1_).length(0))
        (d_626_v1_)[index142_] = (len(_dafny.SeqWithoutIsStrInference([(0) - (p1) for d_628_i0_ in range(default__.abs(835))])) if not(d_627_v2_) else p2)
        d_629_v3_: _dafny.Seq
        d_629_v3_ = _dafny.SeqWithoutIsStrInference([d_627_v2_])
        (globalState).f10 = ((d_626_v1_)[default__.safeIndex(935, (d_626_v1_).length(0))]) * ((0) - (len((d_629_v3_).set(default__.safeIndex(p3, len(d_629_v3_)), d_627_v2_))))
        if (d_627_v2_) == (d_627_v2_):
            d_630_v4_: _dafny.Seq
            d_630_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ffxpkyq"))
            index143_ = default__.safeIndex(935, (d_626_v1_).length(0))
            rhs144_ = default__.safeModuloInt((0) - ((len(_dafny.SeqWithoutIsStrInference([d_627_v2_]))) * (p0)), (d_626_v1_)[default__.safeIndex(935, (d_626_v1_).length(0))])
            rhs145_ = ((0) - (default__.safeModuloInt(len(d_625_v0_), p3))) + (default__.safeDivisionInt(len(d_630_v4_), (d_626_v1_)[default__.safeIndex(935, (d_626_v1_).length(0))]))
            rhs146_ = p2
            lhs135_ = globalState
            lhs136_ = globalState
            lhs137_ = d_626_v1_
            lhs138_ = default__.safeIndex(935, (d_626_v1_).length(0))
            lhs135_.f10 = rhs144_
            lhs136_.f10 = rhs145_
            lhs137_[lhs138_] = rhs146_
            d_631_v5_: C0
            nw132_ = C0()
            nw132_.ctor__((d_630_v4_) != (d_630_v4_))
            d_631_v5_ = nw132_
            d_632_v6_: D3
            d_632_v6_ = D3_DC12((d_631_v5_).f32)
            d_633_v7_: _dafny.MultiSet
            d_633_v7_ = _dafny.MultiSet([p0])
            d_634_v8_: D1
            d_634_v8_ = D1_DC5(p1, (D3_DC12((d_631_v5_).f32)).cf23, not((d_632_v6_).cf23), _dafny.Map({d_633_v7_: d_625_v0_}))
            d_634_v8_ = d_634_v8_
            d_635_v9_: _dafny.Map
            d_635_v9_ = _dafny.Map({-298: d_631_v5_})
            (globalState).f10 = default__.safeModuloInt(default__.safeModuloInt(len(d_635_v9_), (0) - (p3)), (p3) + (p2))
            d_636_v10_: str
            d_636_v10_ = _dafny.CodePoint('n')
            (globalState).f19 = d_636_v10_
        elif True:
            index144_ = default__.safeIndex(935, (d_626_v1_).length(0))
            rhs147_ = (self).fm21((0) - (p3), globalState)
            rhs148_ = p1
            rhs149_ = not((p3) > (p0))
            lhs139_ = globalState
            lhs140_ = d_626_v1_
            lhs141_ = default__.safeIndex(935, (d_626_v1_).length(0))
            lhs142_ = globalState
            lhs139_.f22 = rhs147_
            lhs140_[lhs141_] = rhs148_
            lhs142_.f7 = rhs149_
            d_637_v11_: C0
            nw133_ = C0()
            nw133_.ctor__((self).fm21(p2, globalState))
            d_637_v11_ = nw133_
            d_638_v12_: _dafny.MultiSet
            d_638_v12_ = _dafny.MultiSet([(d_637_v11_).f32])
            d_639_v13_: _dafny.Map
            d_639_v13_ = _dafny.Map({not(not((d_637_v11_).f32)): not(not((d_637_v11_).f32))})
            (globalState).f10 = (((d_638_v12_)[False] if (False) in (d_638_v12_) else len(d_639_v13_)) if (d_637_v11_).f32 else (d_626_v1_)[default__.safeIndex(935, (d_626_v1_).length(0))])
            d_640_v14_: C0
            nw134_ = C0()
            nw134_.ctor__(d_627_v2_)
            d_640_v14_ = nw134_
            d_641_v15_: _dafny.Seq
            d_641_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lroq"))
            d_642_v16_: _dafny.Map
            d_642_v16_ = _dafny.Map({(d_626_v1_)[default__.safeIndex(935, (d_626_v1_).length(0))]: p3})
            d_643_v17_: D2
            d_643_v17_ = D2_DC7(d_642_v16_)
            d_644_v18_: str
            d_644_v18_ = _dafny.CodePoint('x')
            d_645_v19_: D3
            d_645_v19_ = D3_DC11(p3)
            d_646_v20_: _dafny.Seq
            d_646_v20_ = _dafny.SeqWithoutIsStrInference([d_641_v15_])
            d_647_v21_: _dafny.Array
            nw135_ = _dafny.Array(None, 28)
            nw135_[int(0)] = d_641_v15_
            nw135_[int(1)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_648_i1_ in range(default__.abs(663))])
            nw135_[int(2)] = (default__.fm23(d_643_v17_, p2, (0) - ((d_626_v1_)[default__.safeIndex(935, (d_626_v1_).length(0))]), globalState)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dtnxh")))
            nw135_[int(3)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_649_i2_ in range(default__.abs(620))])
            nw135_[int(4)] = d_641_v15_
            nw135_[int(5)] = d_641_v15_
            nw135_[int(6)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vhcdvwyj"))
            nw135_[int(7)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "chvp"))
            nw135_[int(8)] = (d_641_v15_) + (d_641_v15_)
            nw135_[int(9)] = d_641_v15_
            nw135_[int(10)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_650_i3_ in range(default__.abs(44))])
            nw135_[int(11)] = d_641_v15_
            nw135_[int(12)] = (d_641_v15_).set(default__.safeIndex(781, len(d_641_v15_)), d_644_v18_)
            nw135_[int(13)] = (d_641_v15_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xjxlm")))
            nw135_[int(14)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kxppprism"))) + (_dafny.SeqWithoutIsStrInference([d_644_v18_ for d_651_i4_ in range(default__.abs(677))]))
            nw135_[int(15)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qghd"))
            nw135_[int(16)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ajvtl"))) + (d_641_v15_)
            nw135_[int(17)] = _dafny.SeqWithoutIsStrInference([d_644_v18_ for d_652_i5_ in range(default__.abs(137))])
            nw135_[int(18)] = d_641_v15_
            nw135_[int(19)] = default__.fm23(D2_DC7(d_642_v16_), 422, (d_626_v1_)[default__.safeIndex(935, (d_626_v1_).length(0))], globalState)
            nw135_[int(20)] = d_641_v15_
            nw135_[int(21)] = (_dafny.SeqWithoutIsStrInference([d_644_v18_ for d_653_i6_ in range(default__.abs(785))])) + (d_641_v15_)
            nw135_[int(22)] = d_641_v15_
            nw135_[int(23)] = _dafny.SeqWithoutIsStrInference([d_644_v18_ for d_654_i7_ in range(default__.abs(-926))])
            nw135_[int(24)] = d_641_v15_
            nw135_[int(25)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ytkqvx"))
            nw135_[int(26)] = ((d_641_v15_).set(default__.safeIndex((d_645_v19_).cf22, len(d_641_v15_)), d_644_v18_)) + (d_641_v15_)
            nw135_[int(27)] = (d_646_v20_)[default__.safeIndex((d_626_v1_)[default__.safeIndex(935, (d_626_v1_).length(0))], len(d_646_v20_))]
            d_647_v21_ = nw135_
            index145_ = default__.safeIndex(835, (d_647_v21_).length(0))
            (d_647_v21_)[index145_] = d_641_v15_
            d_655_v22_: _dafny.Array
            nw136_ = _dafny.Array(D4.default()(), 28)
            d_655_v22_ = nw136_
            d_656_v23_: D4
            d_656_v23_ = D4_DC13(d_644_v18_)
            d_657_v24_: D4
            d_657_v24_ = D4_DC15(d_656_v23_)
            index146_ = default__.safeIndex(31, (d_655_v22_).length(0))
            (d_655_v22_)[index146_] = d_657_v24_
            index147_ = default__.safeIndex(835, (d_647_v21_).length(0))
            index148_ = default__.safeIndex(935, (d_626_v1_).length(0))
            index149_ = default__.safeIndex(31, (d_655_v22_).length(0))
            rhs150_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "akub"))
            rhs151_ = len(d_629_v3_)
            rhs152_ = d_657_v24_
            rhs153_ = (d_637_v11_).f32
            lhs143_ = d_647_v21_
            lhs144_ = default__.safeIndex(835, (d_647_v21_).length(0))
            lhs145_ = d_626_v1_
            lhs146_ = default__.safeIndex(935, (d_626_v1_).length(0))
            lhs147_ = d_655_v22_
            lhs148_ = default__.safeIndex(31, (d_655_v22_).length(0))
            lhs149_ = globalState
            lhs143_[lhs144_] = rhs150_
            lhs145_[lhs146_] = rhs151_
            lhs147_[lhs148_] = rhs152_
            lhs149_.f7 = rhs153_
        d_658_v25_: _dafny.Array
        nw137_ = _dafny.Array(False, 14)
        d_658_v25_ = nw137_
        d_659_v26_: _dafny.Seq
        d_659_v26_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cfgaepsi"))
        rhs154_ = d_658_v25_
        rhs155_ = d_659_v26_
        rhs156_ = d_625_v0_
        rhs157_ = p3
        lhs150_ = globalState
        d_658_v25_ = rhs154_
        d_659_v26_ = rhs155_
        d_625_v0_ = rhs156_
        lhs150_.f10 = rhs157_
        d_660_v27_: str
        d_660_v27_ = _dafny.CodePoint('m')
        r0 = d_660_v27_
        return r0

    def m7(self, p0, p1, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: bool = False
        d_661_v0_: _dafny.Set
        d_661_v0_ = _dafny.Set({False})
        d_662_v1_: bool
        d_662_v1_ = True
        d_663_v2_: _dafny.Array
        nw138_ = _dafny.Array(None, 6)
        nw138_[int(0)] = d_662_v1_
        nw138_[int(1)] = d_662_v1_
        nw138_[int(2)] = d_662_v1_
        nw138_[int(3)] = d_662_v1_
        nw138_[int(4)] = d_662_v1_
        nw138_[int(5)] = d_662_v1_
        d_663_v2_ = nw138_
        pat_let_tv11_ = d_663_v2_
        pat_let_tv12_ = d_662_v1_
        pat_let_tv13_ = d_663_v2_
        def iife37_(_pat_let8_0):
            def iife38_(d_664_dt__update__tmp_h0_):
                def iife39_(_pat_let9_0):
                    def iife40_(d_665_dt__update_hcf4_h0_):
                        return D0_DC1((d_664_dt__update__tmp_h0_).cf1, (d_664_dt__update__tmp_h0_).cf2, (d_664_dt__update__tmp_h0_).cf3, d_665_dt__update_hcf4_h0_)
                    return iife40_(_pat_let9_0)
                return iife39_((pat_let_tv11_ if pat_let_tv12_ else pat_let_tv13_))
            return iife38_(_pat_let8_0)
        source12_ = iife37_(D0_DC1(len(d_661_v0_), -556, p0, d_663_v2_))
        if source12_.is_DC1:
            d_666___mcc_h0_ = source12_.cf1
            d_667___mcc_h1_ = source12_.cf2
            d_668___mcc_h2_ = source12_.cf3
            d_669___mcc_h3_ = source12_.cf4
            d_670_cf4_ = d_669___mcc_h3_
            d_671_cf3_ = d_668___mcc_h2_
            d_672_cf2_ = d_667___mcc_h1_
            d_673_cf1_ = d_666___mcc_h0_
            d_674_v3_: _dafny.Array
            def lambda27_(d_675_v0_):
                def lambda28_(d_676_i0_):
                    return (d_675_v0_) | (d_675_v0_)

                return lambda28_

            init14_ = lambda27_(d_661_v0_)
            nw139_ = _dafny.Array(None, 25)
            for i0_14_ in range(nw139_.length(0)):
                nw139_[i0_14_] = init14_(i0_14_)
            d_674_v3_ = nw139_
            d_674_v3_ = d_674_v3_
            d_677_v4_: _dafny.Seq
            d_677_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gteyicgu"))
            d_678_v5_: D0
            d_678_v5_ = D0_DC0(d_677_v4_)
            d_679_v6_: _dafny.Set
            d_679_v6_ = _dafny.Set({d_678_v5_})
            d_672_cf2_ = len(d_679_v6_)
            d_680_v7_: _dafny.MultiSet
            d_680_v7_ = _dafny.MultiSet([(d_677_v4_) + (d_677_v4_)])
            d_680_v7_ = _dafny.MultiSet([d_677_v4_, d_677_v4_])
            d_681_v8_: _dafny.Array
            nw140_ = _dafny.Array(None, 4)
            nw140_[int(0)] = d_673_cf1_
            nw140_[int(1)] = (0) - ((len(_dafny.SeqWithoutIsStrInference([p0, d_671_cf3_, len(d_677_v4_), 21]))) + (d_673_cf1_))
            nw140_[int(2)] = d_672_cf2_
            nw140_[int(3)] = (500) + (d_673_cf1_)
            d_681_v8_ = nw140_
            index150_ = default__.safeIndex(706, (d_681_v8_).length(0))
            (d_681_v8_)[index150_] = (0) - (p1)
            index151_ = default__.safeIndex(706, (d_681_v8_).length(0))
            (d_681_v8_)[index151_] = d_673_cf1_
        elif source12_.is_DC2:
            d_682___mcc_h4_ = source12_.cf5
            d_683___mcc_h5_ = source12_.cf6
            d_684_cf6_ = d_683___mcc_h5_
            d_685_cf5_ = d_682___mcc_h4_
            d_686_v9_: _dafny.Seq
            d_686_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rpgfswiru"))
            (globalState).f13 = not ((d_686_v9_) <= (d_686_v9_)) or ((d_685_cf5_) == (len(default__.fm22(d_684_cf6_, globalState))))
            d_687_v10_: _dafny.Array
            def lambda29_(d_688_cf5_):
                def lambda30_(d_689_i1_):
                    return default__.safeModuloInt(d_689_i1_, d_688_cf5_)

                return lambda30_

            init15_ = lambda29_(d_685_cf5_)
            nw141_ = _dafny.Array(None, 5)
            for i0_15_ in range(nw141_.length(0)):
                nw141_[i0_15_] = init15_(i0_15_)
            d_687_v10_ = nw141_
            d_690_v11_: _dafny.Seq
            d_690_v11_ = _dafny.SeqWithoutIsStrInference([d_685_cf5_])
            index152_ = default__.safeIndex(735, (d_687_v10_).length(0))
            (d_687_v10_)[index152_] = (d_690_v11_)[default__.safeIndex((self).fm2(globalState), len(d_690_v11_))]
            index153_ = default__.safeIndex(735, (d_687_v10_).length(0))
            (d_687_v10_)[index153_] = p1
            r0 = d_687_v10_
            d_685_cf5_ = d_685_cf5_
        elif source12_.is_DC0:
            d_691___mcc_h6_ = source12_.cf0
            d_692_cf0_ = d_691___mcc_h6_
            index154_ = default__.safeIndex(164, (d_663_v2_).length(0))
            (d_663_v2_)[index154_] = not(d_662_v1_)
            d_693_v12_: D0
            d_693_v12_ = D0_DC2(527, d_662_v1_)
            pat_let_tv14_ = d_662_v1_
            index155_ = default__.safeIndex(164, (d_663_v2_).length(0))
            def iife41_(_pat_let10_0):
                def iife42_(d_694_dt__update__tmp_h1_):
                    def iife43_(_pat_let11_0):
                        def iife44_(d_695_dt__update_hcf6_h0_):
                            return D0_DC2((d_694_dt__update__tmp_h1_).cf5, d_695_dt__update_hcf6_h0_)
                        return iife44_(_pat_let11_0)
                    return iife43_(pat_let_tv14_)
                return iife42_(_pat_let10_0)
            (d_663_v2_)[index155_] = ((d_662_v1_) or ((iife41_(d_693_v12_)).cf6)) or (False)
            d_696_v13_: _dafny.Array
            nw142_ = _dafny.Array(int(0), 16)
            d_696_v13_ = nw142_
            r0 = d_696_v13_
            d_697_v14_: D3
            d_697_v14_ = D3_DC11(p1)
            (globalState).f10 = (d_697_v14_).cf22
            d_698_v15_: _dafny.Map
            d_698_v15_ = _dafny.Map({d_697_v14_: not(d_662_v1_)})
            d_699_v16_: _dafny.Seq
            d_699_v16_ = _dafny.SeqWithoutIsStrInference([d_698_v15_])
            d_699_v16_ = d_699_v16_
        elif True:
            d_700___mcc_h7_ = source12_.cf7
            d_701_cf7_ = d_700___mcc_h7_
            d_702_v17_: str
            d_702_v17_ = _dafny.CodePoint('a')
            d_703_v18_: _dafny.Seq
            d_703_v18_ = _dafny.SeqWithoutIsStrInference([d_702_v17_, d_702_v17_, _dafny.CodePoint('v'), d_702_v17_, _dafny.CodePoint('g')])
            d_703_v18_ = (d_703_v18_) + (_dafny.SeqWithoutIsStrInference([d_702_v17_, d_702_v17_, d_702_v17_]))
            d_704_v19_: _dafny.Seq
            d_704_v19_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_702_v17_ for d_705_i2_ in range(default__.abs(671))])])
            d_703_v18_ = (d_704_v19_)[default__.safeIndex(len((d_661_v0_ if d_662_v1_ else d_661_v0_)), len(d_704_v19_))]
            d_706_v20_: D0
            d_706_v20_ = D0_DC0(d_703_v18_)
            pat_let_tv15_ = d_703_v18_
            pat_let_tv16_ = d_703_v18_
            def iife45_(_pat_let12_0):
                def iife46_(d_707_dt__update__tmp_h2_):
                    def iife47_(_pat_let13_0):
                        def iife48_(d_708_dt__update_hcf0_h0_):
                            return D0_DC0(d_708_dt__update_hcf0_h0_)
                        return iife48_(_pat_let13_0)
                    return iife47_((pat_let_tv15_) + (pat_let_tv16_))
                return iife46_(_pat_let12_0)
            source13_ = iife45_(d_706_v20_)
            if source13_.is_DC1:
                d_709___mcc_h8_ = source13_.cf1
                d_710___mcc_h9_ = source13_.cf2
                d_711___mcc_h10_ = source13_.cf3
                d_712___mcc_h11_ = source13_.cf4
                d_713_cf4_ = d_712___mcc_h11_
                d_714_cf3_ = d_711___mcc_h10_
                d_715_cf2_ = d_710___mcc_h9_
                d_716_cf1_ = d_709___mcc_h8_
                d_717_v21_: _dafny.Array
                nw143_ = _dafny.Array(_dafny.Array(None, 0), 26)
                d_717_v21_ = nw143_
                d_718_v22_: _dafny.Set
                d_718_v22_ = _dafny.Set({d_714_cf3_, d_715_cf2_})
                rhs158_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l")) if d_662_v1_ else d_703_v18_)
                rhs159_ = d_662_v1_
                rhs160_ = d_717_v21_
                rhs161_ = ((len(d_718_v22_)) > (d_714_cf3_)) == (d_662_v1_)
                lhs151_ = globalState
                d_703_v18_ = rhs158_
                d_662_v1_ = rhs159_
                d_717_v21_ = rhs160_
                lhs151_.f7 = rhs161_
                d_719_v23_: D3
                d_719_v23_ = D3_DC11(p1)
                d_719_v23_ = d_719_v23_
                d_720_v24_: D0
                d_720_v24_ = D0_DC2(p1, d_662_v1_)
                d_721_v25_: D0
                d_721_v25_ = D0_DC3(D0_DC3(d_720_v24_))
                d_722_v26_: _dafny.Array
                nw144_ = _dafny.Array(None, 2)
                nw144_[int(0)] = d_721_v25_
                nw144_[int(1)] = (d_721_v25_ if d_662_v1_ else d_721_v25_)
                d_722_v26_ = nw144_
                index156_ = default__.safeIndex(431, (d_722_v26_).length(0))
                (d_722_v26_)[index156_] = (d_721_v25_ if d_662_v1_ else d_721_v25_)
                d_723_v27_: _dafny.Map
                d_723_v27_ = _dafny.Map({_dafny.Map({d_715_cf2_: d_662_v1_}): d_662_v1_})
                d_724_v28_: D0
                d_724_v28_ = D0_DC2(p0, d_662_v1_)
                index157_ = default__.safeIndex(431, (d_722_v26_).length(0))
                rhs162_ = (d_704_v19_)[default__.safeIndex((0) - (d_714_cf3_), len(d_704_v19_))]
                rhs163_ = d_721_v25_
                rhs164_ = (d_662_v1_ if d_662_v1_ else (d_724_v28_).cf6)
                rhs165_ = d_723_v27_
                lhs152_ = d_722_v26_
                lhs153_ = default__.safeIndex(431, (d_722_v26_).length(0))
                lhs154_ = globalState
                d_703_v18_ = rhs162_
                lhs152_[lhs153_] = rhs163_
                lhs154_.f13 = rhs164_
                d_723_v27_ = rhs165_
                d_725_v29_: _dafny.Array
                nw145_ = _dafny.Array(int(0), 29)
                d_725_v29_ = nw145_
                d_726_v30_: _dafny.Array
                d_726_v30_ = d_725_v29_
                d_727_v31_: _dafny.Seq
                d_727_v31_ = _dafny.SeqWithoutIsStrInference([(d_726_v30_), d_725_v29_, d_725_v29_])
                rhs166_ = (d_727_v31_)[default__.safeIndex(p1, len(d_727_v31_))]
                rhs167_ = d_663_v2_
                rhs168_ = d_662_v1_
                rhs169_ = d_716_cf1_
                r0 = rhs166_
                d_663_v2_ = rhs167_
                r1 = rhs168_
                d_715_cf2_ = rhs169_
            elif source13_.is_DC2:
                d_728___mcc_h12_ = source13_.cf5
                d_729___mcc_h13_ = source13_.cf6
                d_730_cf6_ = d_729___mcc_h13_
                d_731_cf5_ = d_728___mcc_h12_
                (globalState).f13 = d_730_cf6_
                index158_ = default__.safeIndex(413, (d_663_v2_).length(0))
                (d_663_v2_)[index158_] = not((d_703_v18_) == (d_703_v18_))
                d_732_v32_: _dafny.Map
                d_732_v32_ = _dafny.Map({d_662_v1_: not(d_662_v1_)})
                index159_ = default__.safeIndex(413, (d_663_v2_).length(0))
                rhs170_ = default__.safeModuloInt(p1, p1)
                rhs171_ = d_730_cf6_
                rhs172_ = (default__.safeDivisionInt(p1, p0) if (default__.fm30(True, globalState)) == (d_732_v32_) else (0) - (d_731_cf5_))
                lhs155_ = globalState
                lhs156_ = d_663_v2_
                lhs157_ = default__.safeIndex(413, (d_663_v2_).length(0))
                lhs158_ = globalState
                lhs155_.f10 = rhs170_
                lhs156_[lhs157_] = rhs171_
                lhs158_.f10 = rhs172_
                d_733_v33_: _dafny.Seq
                d_733_v33_ = _dafny.SeqWithoutIsStrInference([(d_663_v2_)[default__.safeIndex(413, (d_663_v2_).length(0))]])
                d_734_v34_: _dafny.Seq
                d_734_v34_ = _dafny.SeqWithoutIsStrInference([(d_733_v33_)[default__.safeIndex(p0, len(d_733_v33_))], (d_663_v2_)[default__.safeIndex(413, (d_663_v2_).length(0))]])
                d_734_v34_ = d_733_v33_
                d_735_v35_: _dafny.Array
                nw146_ = _dafny.Array(False, 19)
                d_735_v35_ = nw146_
                d_736_v36_: _dafny.Seq
                d_736_v36_ = _dafny.SeqWithoutIsStrInference([d_735_v35_])
                r1 = ((d_736_v36_) + (d_736_v36_)) <= (d_736_v36_)
            elif source13_.is_DC0:
                d_737___mcc_h14_ = source13_.cf0
                d_738_cf0_ = d_737___mcc_h14_
                d_739_v37_: _dafny.MultiSet
                d_739_v37_ = _dafny.MultiSet([d_662_v1_])
                d_740_v38_: _dafny.Map
                d_740_v38_ = _dafny.Map({p1: _dafny.Map({p0: d_662_v1_})})
                d_741_v39_: _dafny.Seq
                d_741_v39_ = _dafny.SeqWithoutIsStrInference([default__.fm0(p1, globalState), (self).fm2(globalState), (d_739_v37_).cardinality, p0, len(d_740_v38_)])
                d_742_v40_: str
                out3_: str
                out3_ = (self).m6(default__.fm0(p0, globalState), (p0) * (p1), (d_741_v39_)[default__.safeIndex(p0, len(d_741_v39_))], p0, globalState)
                d_742_v40_ = out3_
                (globalState).f22 = d_662_v1_
                d_743_v41_: _dafny.Array
                def lambda31_(d_744_p1_):
                    def lambda32_(d_745_i3_):
                        return (d_745_i3_) * (d_744_p1_)

                    return lambda32_

                init16_ = lambda31_(p1)
                nw147_ = _dafny.Array(None, 24)
                for i0_16_ in range(nw147_.length(0)):
                    nw147_[i0_16_] = init16_(i0_16_)
                d_743_v41_ = nw147_
                d_746_v42_: _dafny.Map
                d_746_v42_ = _dafny.Map({d_662_v1_: d_743_v41_})
                d_746_v42_ = (d_746_v42_).set(d_662_v1_, d_743_v41_)
                (globalState).f13 = ((_dafny.SeqWithoutIsStrInference([d_702_v17_ for d_747_i4_ in range(default__.abs(214))])) + (d_703_v18_)) < ((d_706_v20_).cf0)
            elif True:
                d_748___mcc_h15_ = source13_.cf7
                d_749_cf7_ = d_748___mcc_h15_
                d_750_v43_: _dafny.Array
                nw148_ = _dafny.Array(_dafny.CodePoint('D'), 8)
                d_750_v43_ = nw148_
                d_751_v44_: _dafny.Seq
                d_751_v44_ = _dafny.SeqWithoutIsStrInference([d_750_v43_])
                rhs173_ = _dafny.CodePoint('p')
                rhs174_ = d_662_v1_
                rhs175_ = (d_751_v44_)[default__.safeIndex(89, len(d_751_v44_))]
                lhs159_ = globalState
                d_702_v17_ = rhs173_
                lhs159_.f22 = rhs174_
                d_750_v43_ = rhs175_
                d_663_v2_ = d_663_v2_
                (globalState).f10 = default__.safeDivisionInt(p0, (0) - (p1))
                (globalState).f22 = d_662_v1_
            d_752_v45_: _dafny.Array
            def lambda33_(d_753_v1_):
                def lambda34_(d_754_i5_):
                    return (D4_DC13(_dafny.CodePoint('i')) if (D2_DC8(d_753_v1_, d_753_v1_)).cf19 else D4_DC13(_dafny.CodePoint('o')))

                return lambda34_

            init17_ = lambda33_(d_662_v1_)
            nw149_ = _dafny.Array(None, 4)
            for i0_17_ in range(nw149_.length(0)):
                nw149_[i0_17_] = init17_(i0_17_)
            d_752_v45_ = nw149_
            d_755_v46_: D4
            d_755_v46_ = D4_DC13(d_702_v17_)
            pat_let_tv17_ = d_702_v17_
            index160_ = default__.safeIndex(70, (d_752_v45_).length(0))
            def iife49_(_pat_let14_0):
                def iife50_(d_756_dt__update__tmp_h3_):
                    def iife51_(_pat_let15_0):
                        def iife52_(d_757_dt__update_hcf24_h0_):
                            return D4_DC13(d_757_dt__update_hcf24_h0_)
                        return iife52_(_pat_let15_0)
                    return iife51_(pat_let_tv17_)
                return iife50_(_pat_let14_0)
            (d_752_v45_)[index160_] = iife49_(d_755_v46_)
            pat_let_tv18_ = d_702_v17_
            index161_ = default__.safeIndex(70, (d_752_v45_).length(0))
            def iife53_(_pat_let16_0):
                def iife54_(d_758_dt__update__tmp_h4_):
                    def iife55_(_pat_let17_0):
                        def iife56_(d_759_dt__update_hcf24_h1_):
                            return D4_DC13(d_759_dt__update_hcf24_h1_)
                        return iife56_(_pat_let17_0)
                    return iife55_(pat_let_tv18_)
                return iife54_(_pat_let16_0)
            (d_752_v45_)[index161_] = iife53_(d_755_v46_)
        d_760_v47_: _dafny.Map
        d_760_v47_ = _dafny.Map({not((self).fm21(p1, globalState)): d_663_v2_})
        d_760_v47_ = (d_760_v47_).set((d_662_v1_ if (self).fm21(p0, globalState) else d_662_v1_), d_663_v2_)
        d_761_v48_: _dafny.Seq
        d_761_v48_ = _dafny.SeqWithoutIsStrInference([d_662_v1_])
        hi2_ = p1
        for d_762_i6_ in range(len(d_761_v48_), hi2_):
            index162_ = default__.safeIndex(129, (d_663_v2_).length(0))
            (d_663_v2_)[index162_] = True
            d_763_v49_: _dafny.Map
            d_763_v49_ = _dafny.Map({d_662_v1_: p0})
            d_764_v50_: _dafny.Seq
            d_764_v50_ = _dafny.SeqWithoutIsStrInference([d_762_i6_])
            d_765_v51_: _dafny.Seq
            d_765_v51_ = _dafny.SeqWithoutIsStrInference([((d_763_v49_)[not(d_662_v1_)] if (not(d_662_v1_)) in (d_763_v49_) else p1), d_762_i6_, (d_764_v50_)[default__.safeIndex(p0, len(d_764_v50_))], d_762_i6_, d_762_i6_])
            index163_ = default__.safeIndex(129, (d_663_v2_).length(0))
            (d_663_v2_)[index163_] = ((d_765_v51_)[default__.safeIndex(d_762_i6_, len(d_765_v51_))]) > (d_762_i6_)
            (globalState).f22 = not((d_761_v48_)[default__.safeIndex(p1, len(d_761_v48_))])
            d_766_v52_: D3
            d_766_v52_ = D3_DC11(p1)
            source14_ = d_766_v52_
            if source14_.is_DC11:
                d_767___mcc_h16_ = source14_.cf22
                d_768_cf22_ = d_767___mcc_h16_
                d_769_v53_: _dafny.Seq
                d_769_v53_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lkekff"))
                d_769_v53_ = ((d_769_v53_) + (d_769_v53_)) + (d_769_v53_)
                d_770_v54_: _dafny.Array
                nw150_ = _dafny.Array(_dafny.CodePoint('D'), 17)
                d_770_v54_ = nw150_
                index164_ = default__.safeIndex(428, (d_770_v54_).length(0))
                (d_770_v54_)[index164_] = _dafny.CodePoint('e')
                d_771_v55_: _dafny.Set
                d_771_v55_ = _dafny.Set({d_768_cf22_})
                index165_ = default__.safeIndex(428, (d_770_v54_).length(0))
                (d_770_v54_)[index165_] = (d_769_v53_)[default__.safeIndex((0) - (len((d_771_v55_).intersection(d_771_v55_))), len(d_769_v53_))]
                d_772_v56_: _dafny.Map
                d_772_v56_ = _dafny.Map({_dafny.MultiSet([(d_663_v2_)[default__.safeIndex(129, (d_663_v2_).length(0))], d_662_v1_]): (d_663_v2_)[default__.safeIndex(129, (d_663_v2_).length(0))]})
                d_773_v57_: D1
                d_773_v57_ = D1_DC4(d_772_v56_)
                d_773_v57_ = d_773_v57_
                d_774_v58_: _dafny.MultiSet
                d_774_v58_ = _dafny.MultiSet([(d_663_v2_)[default__.safeIndex(129, (d_663_v2_).length(0))], (d_663_v2_)[default__.safeIndex(129, (d_663_v2_).length(0))], d_662_v1_])
                d_775_v59_: str
                out4_: str
                out4_ = (self).m6(d_762_i6_, d_768_cf22_, d_762_i6_, default__.safeModuloInt(d_768_cf22_, ((d_774_v58_)[d_662_v1_] if (d_662_v1_) in (d_774_v58_) else (0) - (-473))), globalState)
                d_775_v59_ = out4_
            elif source14_.is_DC12:
                d_776___mcc_h17_ = source14_.cf23
                d_777_cf23_ = d_776___mcc_h17_
                d_778_v60_: _dafny.Seq
                d_778_v60_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bxjrgwaol"))
                d_779_v61_: _dafny.Array
                nw151_ = _dafny.Array(None, 6)
                nw151_[int(0)] = d_777_cf23_
                nw151_[int(1)] = d_662_v1_
                nw151_[int(2)] = d_662_v1_
                nw151_[int(3)] = d_777_cf23_
                nw151_[int(4)] = (d_663_v2_)[default__.safeIndex(129, (d_663_v2_).length(0))]
                nw151_[int(5)] = d_777_cf23_
                d_779_v61_ = nw151_
                d_780_v62_: _dafny.MultiSet
                d_780_v62_ = _dafny.MultiSet([False])
                d_781_v63_: _dafny.Map
                d_781_v63_ = _dafny.Map({d_766_v52_: (d_663_v2_)[default__.safeIndex(129, (d_663_v2_).length(0))]})
                d_782_v64_: _dafny.Seq
                d_782_v64_ = _dafny.SeqWithoutIsStrInference([d_781_v63_])
                d_783_v65_: _dafny.Map
                d_783_v65_ = _dafny.Map({d_777_cf23_: (d_663_v2_)[default__.safeIndex(129, (d_663_v2_).length(0))]})
                d_784_v66_: D4
                d_784_v66_ = D4_DC14(d_662_v1_)
                d_785_v67_: D4
                d_785_v67_ = D4_DC15(d_784_v66_)
                d_786_v68_: D4
                d_786_v68_ = D4_DC15(d_785_v67_)
                d_787_v69_: D7
                d_787_v69_ = D7_DC20(_dafny.SeqWithoutIsStrInference([590]), d_779_v61_, d_786_v68_, (d_663_v2_)[default__.safeIndex(129, (d_663_v2_).length(0))], p1)
                d_788_v70_: _dafny.Map
                d_788_v70_ = _dafny.Map({414: (d_787_v69_).cf34})
                d_789_v71_: _dafny.Array
                nw152_ = _dafny.Array(None, 25)
                nw152_[int(0)] = default__.safeModuloInt(len((D7_DC18(_dafny.SeqWithoutIsStrInference([_dafny.Map({_dafny.CodePoint('m'): len(d_761_v48_)}) for d_790_i7_ in range(default__.abs(933))]))).cf29), len(d_778_v60_))
                nw152_[int(1)] = p1
                nw152_[int(2)] = 698
                nw152_[int(3)] = p0
                nw152_[int(4)] = p0
                nw152_[int(5)] = p1
                nw152_[int(6)] = 249
                nw152_[int(7)] = (0) - (default__.safeDivisionInt((0) - (p0), len(d_761_v48_)))
                nw152_[int(8)] = 504
                nw152_[int(9)] = p0
                nw152_[int(10)] = p0
                nw152_[int(11)] = ((_dafny.MultiSet([d_779_v61_])).cardinality) - (-328)
                nw152_[int(12)] = (d_780_v62_).cardinality
                nw152_[int(13)] = default__.safeDivisionInt(d_762_i6_, (0) - (d_762_i6_))
                nw152_[int(14)] = len((d_782_v64_)[default__.safeIndex((d_780_v62_).cardinality, len(d_782_v64_))])
                nw152_[int(15)] = len((d_778_v60_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gicvf"))))
                nw152_[int(16)] = p0
                nw152_[int(17)] = 264
                nw152_[int(18)] = p0
                nw152_[int(19)] = -87
                nw152_[int(20)] = (d_762_i6_ if d_662_v1_ else d_762_i6_)
                nw152_[int(21)] = (d_762_i6_) + (p0)
                nw152_[int(22)] = 487
                nw152_[int(23)] = default__.safeModuloInt(len(d_783_v65_), len(d_788_v70_))
                nw152_[int(24)] = d_762_i6_
                d_789_v71_ = nw152_
                index166_ = default__.safeIndex(364, (d_789_v71_).length(0))
                (d_789_v71_)[index166_] = len(d_778_v60_)
                index167_ = default__.safeIndex(364, (d_789_v71_).length(0))
                rhs176_ = (d_761_v48_) + (d_761_v48_)
                rhs177_ = (_dafny.SeqWithoutIsStrInference([d_762_i6_ for d_791_i8_ in range(default__.abs(912))])) + (d_764_v50_)
                rhs178_ = p0
                lhs160_ = d_789_v71_
                lhs161_ = default__.safeIndex(364, (d_789_v71_).length(0))
                d_761_v48_ = rhs176_
                d_765_v51_ = rhs177_
                lhs160_[lhs161_] = rhs178_
                d_783_v65_ = (d_783_v65_).set((d_663_v2_)[default__.safeIndex(129, (d_663_v2_).length(0))], not(d_777_cf23_))
                (globalState).f7 = (d_761_v48_)[default__.safeIndex(d_762_i6_, len(d_761_v48_))]
                d_783_v65_ = (d_783_v65_ if d_777_cf23_ else (d_783_v65_) | (d_783_v65_))
            elif True:
                d_792___mcc_h18_ = source14_.cf21
                d_793_cf21_ = d_792___mcc_h18_
                (globalState).f7 = (d_663_v2_)[default__.safeIndex(129, (d_663_v2_).length(0))]
                (globalState).f13 = (d_663_v2_)[default__.safeIndex(129, (d_663_v2_).length(0))]
                d_794_v72_: _dafny.Map
                d_794_v72_ = _dafny.Map({d_762_i6_: d_761_v48_})
                d_795_v73_: C0
                nw153_ = C0()
                nw153_.ctor__((d_663_v2_)[default__.safeIndex(129, (d_663_v2_).length(0))])
                d_795_v73_ = nw153_
                d_796_v74_: _dafny.Seq
                d_796_v74_ = _dafny.SeqWithoutIsStrInference([d_795_v73_])
                d_797_v75_: _dafny.Seq
                d_797_v75_ = _dafny.SeqWithoutIsStrInference([d_761_v48_, ((d_794_v72_)[len(d_796_v74_)] if (len(d_796_v74_)) in (d_794_v72_) else d_761_v48_)])
                d_798_v76_: D1
                d_798_v76_ = D1_DC6(d_762_i6_, d_797_v75_, d_762_i6_, (d_795_v73_).f32)
                d_799_v77_: _dafny.Set
                d_799_v77_ = _dafny.Set({d_798_v76_})
                d_800_v78_: _dafny.MultiSet
                d_800_v78_ = _dafny.MultiSet([d_799_v77_])
                d_801_v79_: D3
                d_801_v79_ = D3_DC12((d_795_v73_).f32)
                d_802_v80_: _dafny.Map
                d_802_v80_ = _dafny.Map({d_801_v79_: d_800_v78_})
                (globalState).f22 = ((_dafny.MultiSet([d_799_v77_])) - (((d_802_v80_)[d_801_v79_] if (d_801_v79_) in (d_802_v80_) else d_800_v78_))).issubset(d_800_v78_)
                d_803_v81_: _dafny.Seq
                d_803_v81_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "epv"))
                d_804_v82_: C0
                nw154_ = C0()
                nw154_.ctor__((len(d_803_v81_)) > (p0))
                d_804_v82_ = nw154_
            d_805_v83_: C0
            nw155_ = C0()
            nw155_.ctor__(not (d_662_v1_) or ((d_663_v2_)[default__.safeIndex(129, (d_663_v2_).length(0))]))
            d_805_v83_ = nw155_
        d_806_v84_: D4
        d_806_v84_ = D4_DC15(D4_DC14(d_662_v1_))
        d_807_v85_: _dafny.Map
        d_807_v85_ = _dafny.Map({(d_661_v0_).ispropersubset(d_661_v0_): d_806_v84_})
        d_807_v85_ = (_dafny.Map({False: d_806_v84_})) | (d_807_v85_)
        d_808_v86_: _dafny.Seq
        d_808_v86_ = _dafny.SeqWithoutIsStrInference([d_663_v2_, d_663_v2_, d_663_v2_, (d_663_v2_ if True else d_663_v2_), d_663_v2_])
        d_663_v2_ = (d_808_v86_)[default__.safeIndex(p0, len(d_808_v86_))]
        d_809_v88_: _dafny.Seq
        d_809_v88_ = _dafny.SeqWithoutIsStrInference([p1])
        d_810_v89_: _dafny.Map
        d_810_v89_ = _dafny.Map({p1: d_662_v1_})
        def iife57_():
            coll21_ = _dafny.Map()
            compr_21_: int
            for compr_21_ in (d_809_v88_).Elements:
                d_811_v87_: int = compr_21_
                if (d_811_v87_) in (d_809_v88_):
                    coll21_[(d_811_v87_) - (p0)] = d_662_v1_
            return _dafny.Map(coll21_)
        (globalState).f13 = (d_662_v1_ if (iife57_()
        ) == (d_810_v89_) else d_662_v1_)
        def lambda35_(d_812_p1_):
            def lambda36_(d_813_i9_):
                return default__.safeModuloInt(d_813_i9_, d_812_p1_)

            return lambda36_

        init18_ = lambda35_(p1)
        nw156_ = _dafny.Array(None, 13)
        for i0_18_ in range(nw156_.length(0)):
            nw156_[i0_18_] = init18_(i0_18_)
        r0 = nw156_
        r1 = d_662_v1_
        return r0, r1


class C2(T1, T0):
    def  __init__(self):
        self.f33: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    def ctor__(self, f33):
        (self).f33 = f33

    def fm1(self, p0, p1, p2, p3, globalState):
        def iife58_():
            coll22_ = _dafny.Map()
            compr_22_: _dafny.MultiSet
            for compr_22_ in (_dafny.MultiSet([_dafny.MultiSet([True, not(True)])])).Elements:
                d_814_v0_: _dafny.MultiSet = compr_22_
                if (d_814_v0_) in (_dafny.MultiSet([_dafny.MultiSet([True, not(True)])])):
                    coll22_[d_814_v0_] = False
            return _dafny.Map(coll22_)
        return ((iife58_()
        ) | (_dafny.Map({_dafny.MultiSet([False]): True}))) | (_dafny.Map({_dafny.MultiSet([not(True)]): not(False)}))

    def fm2(self, globalState):
        return ((999) - (len(_dafny.Set({(0) - (-668), -602})))) + ((len(_dafny.Set({-570, len(self.f33)})) if True else len(_dafny.Set({not(False), True, False}))))

    def fm31(self, p0, globalState):
        return not ((True if not(True) else False)) or ((_dafny.CodePoint('n')) not in (self.f33))

    def m0(self, p0, p1, p2, p3, globalState):
        d_815_v0_: _dafny.Array
        nw157_ = _dafny.Array(int(0), 23)
        d_815_v0_ = nw157_
        d_816_v1_: _dafny.MultiSet
        d_816_v1_ = _dafny.MultiSet([p1, p1])
        d_817_v2_: _dafny.Array
        nw158_ = _dafny.Array(None, 8)
        nw158_[int(0)] = False
        nw158_[int(1)] = p3
        nw158_[int(2)] = True
        nw158_[int(3)] = p1
        nw158_[int(4)] = p3
        nw158_[int(5)] = p3
        nw158_[int(6)] = p3
        nw158_[int(7)] = p1
        d_817_v2_ = nw158_
        d_818_v3_: _dafny.MultiSet
        d_818_v3_ = _dafny.MultiSet([d_817_v2_])
        d_819_v4_: _dafny.MultiSet
        d_819_v4_ = _dafny.MultiSet([(d_816_v1_).cardinality, len(self.f33), (d_818_v3_).cardinality])
        d_820_v5_: C0
        nw159_ = C0()
        nw159_.ctor__(p1)
        d_820_v5_ = nw159_
        d_821_v6_: _dafny.Set
        d_821_v6_ = _dafny.Set({d_820_v5_, d_820_v5_})
        d_822_v7_: _dafny.Map
        d_822_v7_ = _dafny.Map({d_819_v4_: len(d_821_v6_)})
        d_823_v8_: _dafny.Set
        d_823_v8_ = _dafny.Set({p1})
        index168_ = default__.safeIndex(146, (d_815_v0_).length(0))
        (d_815_v0_)[index168_] = default__.safeDivisionInt(855, ((d_822_v7_)[d_819_v4_] if (d_819_v4_) in (d_822_v7_) else len(d_823_v8_)))
        d_824_v9_: D3
        d_824_v9_ = D3_DC10(d_816_v1_)
        d_825_v10_: _dafny.Set
        d_825_v10_ = _dafny.Set({d_824_v9_, D3_DC10(d_816_v1_), d_824_v9_, d_824_v9_})
        d_826_v11_: _dafny.Map
        d_826_v11_ = _dafny.Map({(self).fm31(d_819_v4_, globalState): d_825_v10_})
        index169_ = default__.safeIndex(146, (d_815_v0_).length(0))
        (d_815_v0_)[index169_] = len((d_826_v11_).set((self).fm31(d_819_v4_, globalState), (d_825_v10_).intersection(d_825_v10_)))
        d_827_v12_: _dafny.Array
        nw160_ = _dafny.Array(_dafny.Array(None, 0), 2)
        d_827_v12_ = nw160_
        d_827_v12_ = d_827_v12_
        d_828_v13_: _dafny.Array
        nw161_ = _dafny.Array(_dafny.MultiSet({}), 12)
        d_828_v13_ = nw161_
        d_829_v14_: str
        d_829_v14_ = _dafny.CodePoint('h')
        d_830_v15_: _dafny.Seq
        d_830_v15_ = _dafny.SeqWithoutIsStrInference([default__.fm0(p2, globalState)])
        d_831_v16_: _dafny.MultiSet
        d_831_v16_ = _dafny.MultiSet([(((d_830_v15_).set(default__.safeIndex(468, len(d_830_v15_)), p2)).set(default__.safeIndex(p0, len((d_830_v15_).set(default__.safeIndex(468, len(d_830_v15_)), p2))), len(_dafny.Map({(d_820_v5_).f32: (0) - (p2)})))).set(default__.safeIndex((d_815_v0_)[default__.safeIndex(146, (d_815_v0_).length(0))], len(((d_830_v15_).set(default__.safeIndex(468, len(d_830_v15_)), p2)).set(default__.safeIndex(p0, len((d_830_v15_).set(default__.safeIndex(468, len(d_830_v15_)), p2))), len(_dafny.Map({(d_820_v5_).f32: (0) - (p2)}))))), p0), _dafny.SeqWithoutIsStrInference([-686 for d_832_i0_ in range(default__.abs(261))])])
        index170_ = default__.safeIndex(710, (d_828_v13_).length(0))
        (d_828_v13_)[index170_] = (default__.fm32(d_829_v14_, (d_820_v5_).f32, d_829_v14_, globalState)) - (d_831_v16_)
        d_833_v17_: _dafny.Seq
        d_833_v17_ = _dafny.SeqWithoutIsStrInference([p1])
        d_834_v18_: _dafny.Map
        d_834_v18_ = _dafny.Map({p2: len(d_833_v17_)})
        d_835_v19_: _dafny.Seq
        d_835_v19_ = _dafny.SeqWithoutIsStrInference([d_834_v18_])
        d_836_v20_: _dafny.Map
        d_836_v20_ = _dafny.Map({d_835_v19_: _dafny.MultiSet([d_830_v15_])})
        index171_ = default__.safeIndex(710, (d_828_v13_).length(0))
        (d_828_v13_)[index171_] = ((d_836_v20_)[d_835_v19_] if (d_835_v19_) in (d_836_v20_) else d_831_v16_)
        if p1:
            (self).f33 = self.f33
            if not(p1):
                d_837_v21_: _dafny.Map
                d_837_v21_ = _dafny.Map({not((d_820_v5_).f32): (d_820_v5_).f32})
                d_837_v21_ = (d_837_v21_) | (d_837_v21_)
                rhs179_ = d_815_v0_
                rhs180_ = default__.fm0(p0, globalState)
                lhs162_ = globalState
                d_815_v0_ = rhs179_
                lhs162_.f10 = rhs180_
                d_838_v22_: C0
                nw162_ = C0()
                nw162_.ctor__((p1) or (p3))
                d_838_v22_ = nw162_
                (globalState).f10 = (0) - ((default__.safeModuloInt((d_815_v0_)[default__.safeIndex(146, (d_815_v0_).length(0))], p2)) * (((d_815_v0_)[default__.safeIndex(146, (d_815_v0_).length(0))]) * (p2)))
                (self).f33 = self.f33
            elif True:
                (globalState).f10 = p0
                (globalState).f10 = default__.safeDivisionInt((d_815_v0_)[default__.safeIndex(146, (d_815_v0_).length(0))], len(d_833_v17_))
                (globalState).f10 = default__.fm0((self).fm2(globalState), globalState)
                (globalState).f7 = (True if (self).fm31(d_819_v4_, globalState) else (p0) != (len(self.f33)))
                index172_ = default__.safeIndex(146, (d_815_v0_).length(0))
                (d_815_v0_)[index172_] = (self).fm2(globalState)
            d_839_v23_: D2
            d_839_v23_ = D2_DC9(D2_DC9(D2_DC7(d_834_v18_)))
            d_840_v24_: _dafny.Map
            d_840_v24_ = _dafny.Map({D2_DC9(d_839_v23_): (d_820_v5_).f32})
            d_841_v25_: D2
            d_841_v25_ = D2_DC9(d_839_v23_)
            d_840_v24_ = (d_840_v24_).set(d_841_v25_, (self).fm31(_dafny.MultiSet(d_830_v15_), globalState))
            d_842_v26_: _dafny.Set
            d_842_v26_ = _dafny.Set({len(d_823_v8_)})
            d_843_v27_: _dafny.Map
            d_843_v27_ = _dafny.Map({p1: default__.safeDivisionInt(-76, len(d_842_v26_))})
            d_843_v27_ = (d_843_v27_).set(p1, (d_815_v0_)[default__.safeIndex(146, (d_815_v0_).length(0))])
            d_844_v28_: C1
            nw163_ = C1()
            nw163_.ctor__()
            d_844_v28_ = nw163_
        elif True:
            d_845_v29_: _dafny.Map
            d_845_v29_ = _dafny.Map({(p3) or (False): p2})
            (globalState).f10 = ((d_845_v29_)[(d_820_v5_).f32] if ((d_820_v5_).f32) in (d_845_v29_) else (self).fm2(globalState))
            index173_ = default__.safeIndex(168, (d_817_v2_).length(0))
            (d_817_v2_)[index173_] = True
            index174_ = default__.safeIndex(168, (d_817_v2_).length(0))
            def iife59_():
                coll23_ = _dafny.Map()
                compr_23_: int
                for compr_23_ in _dafny.IntegerRange(416, 822):
                    d_846_v30_: int = compr_23_
                    if ((416) <= (d_846_v30_)) and ((d_846_v30_) < (822)):
                        coll23_[default__.safeDivisionInt(d_846_v30_, p0)] = p3
                return _dafny.Map(coll23_)
            rhs181_ = ((d_815_v0_)[default__.safeIndex(146, (d_815_v0_).length(0))]) not in (d_834_v18_)
            rhs182_ = (len(iife59_()
            ) if p3 else default__.safeModuloInt(p0, len(d_830_v15_)))
            lhs163_ = d_817_v2_
            lhs164_ = default__.safeIndex(168, (d_817_v2_).length(0))
            lhs165_ = globalState
            lhs163_[lhs164_] = rhs181_
            lhs165_.f10 = rhs182_
            (globalState).f13 = p3
            d_847_v31_: D8
            d_847_v31_ = D8_DC22((d_815_v0_)[default__.safeIndex(146, (d_815_v0_).length(0))], (d_817_v2_)[default__.safeIndex(168, (d_817_v2_).length(0))], d_829_v14_)
            d_848_v32_: _dafny.Map
            d_848_v32_ = _dafny.Map({not((d_847_v31_).cf37): (d_820_v5_).f32})
            source15_ = d_848_v32_
            d_849___mcc_h0_ = source15_
            d_850_cf27_ = d_849___mcc_h0_
            index175_ = default__.safeIndex(146, (d_815_v0_).length(0))
            rhs183_ = False
            rhs184_ = ((d_830_v15_) + (d_830_v15_)) <= (d_830_v15_)
            rhs185_ = p2
            lhs166_ = globalState
            lhs167_ = globalState
            lhs168_ = d_815_v0_
            lhs169_ = default__.safeIndex(146, (d_815_v0_).length(0))
            lhs166_.f13 = rhs183_
            lhs167_.f13 = rhs184_
            lhs168_[lhs169_] = rhs185_
            (globalState).f7 = True
            d_851_v33_: _dafny.Array
            nw164_ = _dafny.Array(None, 9)
            nw164_[int(0)] = (675) >= ((d_815_v0_)[default__.safeIndex(146, (d_815_v0_).length(0))])
            nw164_[int(1)] = not((d_819_v4_).isdisjoint(d_819_v4_))
            nw164_[int(2)] = (self).fm31(d_819_v4_, globalState)
            nw164_[int(3)] = (d_820_v5_).f32
            nw164_[int(4)] = (d_820_v5_).f32
            nw164_[int(5)] = (d_820_v5_).f32
            nw164_[int(6)] = False
            nw164_[int(7)] = ((d_815_v0_)[default__.safeIndex(146, (d_815_v0_).length(0))]) > (p0)
            nw164_[int(8)] = (d_820_v5_).f32
            d_851_v33_ = nw164_
            rhs186_ = d_851_v33_
            rhs187_ = p2
            rhs188_ = (self.f33) + (self.f33)
            rhs189_ = (0) - ((default__.fm0(p2, globalState)) + (default__.fm0((0) - (p2), globalState)))
            lhs170_ = globalState
            lhs171_ = self
            lhs172_ = globalState
            d_851_v33_ = rhs186_
            lhs170_.f10 = rhs187_
            lhs171_.f33 = rhs188_
            lhs172_.f10 = rhs189_
            (globalState).f22 = not(not ((d_820_v5_).f32) or ((self).fm31(d_819_v4_, globalState)))
            d_852_v34_: _dafny.Array
            nw165_ = _dafny.Array(False, 10)
            d_852_v34_ = nw165_
            d_853_v35_: D0
            d_853_v35_ = D0_DC1(p0, p2, (0) - (p2), d_852_v34_)
            d_854_v36_: _dafny.Map
            d_854_v36_ = _dafny.Map({len(d_833_v17_): (d_817_v2_)[default__.safeIndex(168, (d_817_v2_).length(0))]})
            d_853_v35_ = D0_DC1(len(d_854_v36_), (len(self.f33)) + (len(default__.fm33((d_820_v5_).f32, (d_820_v5_).f32, p3, (d_820_v5_).f32, globalState))), ((d_834_v18_)[(d_815_v0_)[default__.safeIndex(146, (d_815_v0_).length(0))]] if ((d_815_v0_)[default__.safeIndex(146, (d_815_v0_).length(0))]) in (d_834_v18_) else (d_815_v0_)[default__.safeIndex(146, (d_815_v0_).length(0))]), d_852_v34_)
        if p1:
            if False:
                (self).f33 = ((self.f33) + ((self.f33).set(default__.safeIndex(p0, len(self.f33)), _dafny.CodePoint('u')))) + (self.f33)
                d_833_v17_ = (d_833_v17_) + ((d_833_v17_) + (d_833_v17_))
                (globalState).f13 = ((d_820_v5_).f32 if True else (d_820_v5_).f32)
                (globalState).f13 = p3
                (globalState).f10 = (0) - (p0)
            elif True:
                d_855_v37_: _dafny.Seq
                d_855_v37_ = _dafny.SeqWithoutIsStrInference([d_815_v0_])
                d_855_v37_ = d_855_v37_
                d_856_v38_: C1
                nw166_ = C1()
                nw166_.ctor__()
                d_856_v38_ = nw166_
                (globalState).f10 = default__.fm0(p0, globalState)
                (globalState).f22 = (d_856_v38_).fm21((d_815_v0_)[default__.safeIndex(146, (d_815_v0_).length(0))], globalState)
                d_857_v39_: _dafny.Array
                nw167_ = _dafny.Array(_dafny.MultiSet({}), 7)
                d_857_v39_ = nw167_
                index176_ = default__.safeIndex(967, (d_857_v39_).length(0))
                (d_857_v39_)[index176_] = d_816_v1_
                index177_ = default__.safeIndex(967, (d_857_v39_).length(0))
                (d_857_v39_)[index177_] = _dafny.MultiSet((d_833_v17_) + ((d_833_v17_) + (d_833_v17_)))
            d_858_v40_: C0
            nw168_ = C0()
            nw168_.ctor__(not ((d_820_v5_).f32) or (not(p3)))
            d_858_v40_ = nw168_
            d_859_v41_: _dafny.Array
            nw169_ = _dafny.Array(_dafny.Array(None, 0), 27)
            d_859_v41_ = nw169_
            d_859_v41_ = (d_859_v41_ if (p0) < (len(default__.fm34(p1, (self.f33)[default__.safeIndex(p2, len(self.f33))], (d_815_v0_)[default__.safeIndex(146, (d_815_v0_).length(0))], globalState))) else d_859_v41_)
            d_860_v42_: _dafny.Seq
            d_860_v42_ = _dafny.SeqWithoutIsStrInference([((d_833_v17_).set(default__.safeIndex(p0, len(d_833_v17_)), p1)) + (d_833_v17_), _dafny.SeqWithoutIsStrInference([(d_820_v5_).f32, False, (d_858_v40_).f32, (self).fm31(_dafny.MultiSet([p2]), globalState)])])
            d_861_v43_: _dafny.Map
            d_861_v43_ = _dafny.Map({(d_858_v40_).f32: d_817_v2_})
            d_862_v44_: _dafny.Set
            d_862_v44_ = _dafny.Set({((d_861_v43_)[p3] if (p3) in (d_861_v43_) else d_817_v2_)})
            d_863_v45_: D4
            d_863_v45_ = D4_DC13(d_829_v14_)
            d_864_v46_: D4
            d_864_v46_ = D4_DC15(d_863_v45_)
            d_865_v47_: D7
            d_865_v47_ = D7_DC20(d_830_v15_, d_817_v2_, d_864_v46_, (d_820_v5_).f32, (d_815_v0_)[default__.safeIndex(146, (d_815_v0_).length(0))])
            index178_ = default__.safeIndex(146, (d_815_v0_).length(0))
            rhs190_ = _dafny.SeqWithoutIsStrInference([d_833_v17_, d_833_v17_, _dafny.SeqWithoutIsStrInference([not(p3)]), d_833_v17_])
            rhs191_ = default__.safeModuloInt((p0) * (p0), p0)
            rhs192_ = (d_862_v44_).isdisjoint(_dafny.Set({d_817_v2_, d_817_v2_, (d_865_v47_).cf31, d_817_v2_}))
            lhs173_ = d_815_v0_
            lhs174_ = default__.safeIndex(146, (d_815_v0_).length(0))
            lhs175_ = globalState
            d_860_v42_ = rhs190_
            lhs173_[lhs174_] = rhs191_
            lhs175_.f22 = rhs192_
            (globalState).f7 = (d_820_v5_).f32
        elif True:
            d_866_v48_: _dafny.Map
            d_866_v48_ = _dafny.Map({(d_820_v5_).f32: (d_820_v5_).f32})
            index179_ = default__.safeIndex(146, (d_815_v0_).length(0))
            (d_815_v0_)[index179_] = (0) - ((p0) + (len((d_866_v48_).set(p1, (d_820_v5_).f32))))
            d_867_v49_: _dafny.Map
            d_867_v49_ = _dafny.Map({self.f33: p1})
            d_866_v48_ = (d_866_v48_).set((d_820_v5_).f32, ((d_867_v49_)[self.f33] if (self.f33) in (d_867_v49_) else p1))
            d_868_v50_: C1
            nw170_ = C1()
            nw170_.ctor__()
            d_868_v50_ = nw170_
            d_869_v51_: _dafny.Map
            d_869_v51_ = _dafny.Map({_dafny.CodePoint('t'): len((_dafny.SeqWithoutIsStrInference([d_830_v15_, _dafny.SeqWithoutIsStrInference([len(self.f33), p0])])).set(default__.safeIndex(len(self.f33), len(_dafny.SeqWithoutIsStrInference([d_830_v15_, _dafny.SeqWithoutIsStrInference([len(self.f33), p0])]))), d_830_v15_))})
            d_870_v52_: _dafny.Map
            d_870_v52_ = _dafny.Map({d_869_v51_: default__.fm35((d_815_v0_)[default__.safeIndex(146, (d_815_v0_).length(0))], p3, not((d_820_v5_).f32), globalState)})
            def iife60_():
                coll24_ = _dafny.Map()
                compr_24_: _dafny.Map
                for compr_24_ in (default__.fm36(d_829_v14_, D3_DC10(d_816_v1_), p1, globalState)).Elements:
                    d_872_v53_: _dafny.Map = compr_24_
                    if (d_872_v53_) in (default__.fm36(d_829_v14_, D3_DC10(d_816_v1_), p1, globalState)):
                        coll24_[d_872_v53_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wqaidjcb"))
                return _dafny.Map(coll24_)
            d_870_v52_ = ((d_870_v52_).set(d_869_v51_, _dafny.SeqWithoutIsStrInference([d_829_v14_ for d_871_i1_ in range(default__.abs(-70))])) if (self.f33) < (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nos"))) else (d_870_v52_) | (iife60_()
            ))
            (globalState).f7 = p1
        index180_ = default__.safeIndex(146, (d_815_v0_).length(0))
        (d_815_v0_)[index180_] = 921

    def m1(self, p0, p1, p2, globalState):
        r0: D0 = D0.default()()
        r1: bool = False
        d_873_v0_: _dafny.Array
        def lambda37_(d_874_p2_, d_875_p1_):
            def lambda38_(d_876_i0_):
                return D7_DC18(_dafny.SeqWithoutIsStrInference([_dafny.Map({_dafny.CodePoint('g'): len(_dafny.Map({d_874_p2_: d_874_p2_}))}), _dafny.Map({_dafny.CodePoint('f'): d_875_p1_})]))

            return lambda38_

        init19_ = lambda37_(p2, p1)
        nw171_ = _dafny.Array(None, 3)
        for i0_19_ in range(nw171_.length(0)):
            nw171_[i0_19_] = init19_(i0_19_)
        d_873_v0_ = nw171_
        d_877_v1_: str
        d_877_v1_ = _dafny.CodePoint('k')
        d_878_v2_: _dafny.Seq
        d_878_v2_ = _dafny.SeqWithoutIsStrInference([p2, p2, p2, p2, p2])
        d_879_v3_: _dafny.Map
        d_879_v3_ = _dafny.Map({d_877_v1_: len(d_878_v2_)})
        index181_ = default__.safeIndex(314, (d_873_v0_).length(0))
        (d_873_v0_)[index181_] = D7_DC18((_dafny.SeqWithoutIsStrInference([d_879_v3_])).set(default__.safeIndex(p1, len(_dafny.SeqWithoutIsStrInference([d_879_v3_]))), d_879_v3_))
        index182_ = default__.safeIndex(314, (d_873_v0_).length(0))
        (d_873_v0_)[index182_] = default__.fm37(p1, p1, p1, globalState)
        (self).f33 = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xcb"))
        d_880_v4_: _dafny.Array
        def lambda39_(d_881_p1_):
            def lambda40_(d_882_i1_):
                return default__.safeModuloInt(d_882_i1_, d_881_p1_)

            return lambda40_

        init20_ = lambda39_(p1)
        nw172_ = _dafny.Array(None, 16)
        for i0_20_ in range(nw172_.length(0)):
            nw172_[i0_20_] = init20_(i0_20_)
        d_880_v4_ = nw172_
        index183_ = default__.safeIndex(887, (d_880_v4_).length(0))
        (d_880_v4_)[index183_] = p1
        d_883_v5_: _dafny.Set
        d_883_v5_ = _dafny.Set({p1, p1, len(_dafny.SeqWithoutIsStrInference([p2]))})
        d_884_v6_: _dafny.MultiSet
        d_884_v6_ = _dafny.MultiSet([default__.fm0(len(d_883_v5_), globalState), p1])
        d_885_v7_: D8
        d_885_v7_ = D8_DC22(p1, p2, d_877_v1_)
        d_886_v8_: _dafny.Map
        d_886_v8_ = _dafny.Map({p1: p1})
        d_887_v10_: D1
        d_887_v10_ = D1_DC6(p1, _dafny.SeqWithoutIsStrInference([d_878_v2_, _dafny.SeqWithoutIsStrInference([p2])]), p1, p2)
        d_888_v11_: _dafny.Array
        nw173_ = _dafny.Array(None, 3)
        nw173_[int(0)] = (((d_884_v6_)[(d_885_v7_).cf36] if ((d_885_v7_).cf36) in (d_884_v6_) else p1)) >= (len(self.f33))
        def iife61_():
            coll25_ = _dafny.Set()
            compr_25_: int
            for compr_25_ in (d_886_v8_).keys.Elements:
                d_889_v9_: int = compr_25_
                if (d_889_v9_) in (d_886_v8_):
                    coll25_ = coll25_.union(_dafny.Set([(d_889_v9_) * (-326)]))
            return _dafny.Set(coll25_)
        nw173_[int(1)] = (iife61_()
        ).ispropersubset(d_883_v5_)
        nw173_[int(2)] = (d_887_v10_).cf16
        d_888_v11_ = nw173_
        index184_ = default__.safeIndex(951, (d_888_v11_).length(0))
        (d_888_v11_)[index184_] = not(p2)
        index185_ = default__.safeIndex(405, (d_880_v4_).length(0))
        (d_880_v4_)[index185_] = p1
        index186_ = default__.safeIndex(887, (d_880_v4_).length(0))
        index187_ = default__.safeIndex(951, (d_888_v11_).length(0))
        index188_ = default__.safeIndex(405, (d_880_v4_).length(0))
        rhs193_ = len((d_878_v2_) + ((d_878_v2_).set(default__.safeIndex((self).fm2(globalState), len(d_878_v2_)), (self).fm31(d_884_v6_, globalState))))
        rhs194_ = not (not((_dafny.Set({p2, p2})).isdisjoint(default__.fm38(len(p0), True, globalState)))) or (p2)
        rhs195_ = (p1) * (default__.safeModuloInt(-561, p1))
        rhs196_ = (d_884_v6_).cardinality
        rhs197_ = p1
        lhs176_ = d_880_v4_
        lhs177_ = default__.safeIndex(887, (d_880_v4_).length(0))
        lhs178_ = d_888_v11_
        lhs179_ = default__.safeIndex(951, (d_888_v11_).length(0))
        lhs180_ = globalState
        lhs181_ = d_880_v4_
        lhs182_ = default__.safeIndex(405, (d_880_v4_).length(0))
        lhs183_ = globalState
        lhs176_[lhs177_] = rhs193_
        lhs178_[lhs179_] = rhs194_
        lhs180_.f10 = rhs195_
        lhs181_[lhs182_] = rhs196_
        lhs183_.f10 = rhs197_
        _ingredientsd_0 = []
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_880_v4_).length(0)):
            d_890_i2_: int = guard_loop_1_
            if (True) and (((0) <= (d_890_i2_)) and ((d_890_i2_) < ((d_880_v4_).length(0)))):
                _ingredientsd_0.append((d_880_v4_, int(d_890_i2_), default__.safeModuloInt(d_890_i2_, (d_880_v4_)[default__.safeIndex(887, (d_880_v4_).length(0))])))
        for _tupd_0 in _ingredientsd_0:
            _tupd_0[0][_tupd_0[1]] = _tupd_0[2]
        d_891_v12_: D0
        d_891_v12_ = D0_DC0(self.f33)
        d_892_v13_: _dafny.Map
        d_892_v13_ = _dafny.Map({(d_880_v4_)[default__.safeIndex(887, (d_880_v4_).length(0))]: self.f33})
        d_893_v14_: _dafny.Array
        nw174_ = _dafny.Array(None, 20)
        nw174_[int(0)] = (d_891_v12_).cf0
        nw174_[int(1)] = (D0_DC0(self.f33)).cf0
        nw174_[int(2)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gbvposwab"))
        nw174_[int(3)] = (self.f33).set(default__.safeIndex(p1, len(self.f33)), d_877_v1_)
        nw174_[int(4)] = self.f33
        nw174_[int(5)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iwtorenfq"))
        nw174_[int(6)] = self.f33
        nw174_[int(7)] = ((d_892_v13_)[p1] if (p1) in (d_892_v13_) else self.f33)
        nw174_[int(8)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "scc"))
        nw174_[int(9)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qltrd"))) + (self.f33)
        nw174_[int(10)] = _dafny.SeqWithoutIsStrInference([d_877_v1_ for d_894_i3_ in range(default__.abs(30))])
        nw174_[int(11)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ssiokybl"))
        nw174_[int(12)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "atla"))
        nw174_[int(13)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j"))
        nw174_[int(14)] = (self.f33).set(default__.safeIndex(len(d_886_v8_), len(self.f33)), d_877_v1_)
        nw174_[int(15)] = (self.f33) + (default__.fm35((d_880_v4_)[default__.safeIndex(887, (d_880_v4_).length(0))], not(p2), (d_888_v11_)[default__.safeIndex(951, (d_888_v11_).length(0))], globalState))
        nw174_[int(16)] = self.f33
        nw174_[int(17)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uokfrfjh"))
        nw174_[int(18)] = self.f33
        nw174_[int(19)] = default__.fm35(308, p2, False, globalState)
        d_893_v14_ = nw174_
        d_895_v15_: _dafny.Map
        d_895_v15_ = _dafny.Map({p2: (d_888_v11_)[default__.safeIndex(951, (d_888_v11_).length(0))]})
        index189_ = default__.safeIndex(597, (d_893_v14_).length(0))
        (d_893_v14_)[index189_] = (self.f33).set(default__.safeIndex(len(self.f33), len(self.f33)), (self.f33)[default__.safeIndex(len(d_895_v15_), len(self.f33))])
        index190_ = default__.safeIndex(597, (d_893_v14_).length(0))
        (d_893_v14_)[index190_] = self.f33
        d_896_v16_: _dafny.Set
        d_896_v16_ = _dafny.Set({p2, (_dafny.MultiSet([len(d_878_v2_)])).ispropersubset(d_884_v6_)})
        d_897_v17_: _dafny.Seq
        d_897_v17_ = _dafny.SeqWithoutIsStrInference([d_896_v16_])
        rhs198_ = ((default__.fm38(p1, p2, globalState)).intersection(d_896_v16_)) - ((d_897_v17_)[default__.safeIndex(p1, len(d_897_v17_))])
        rhs199_ = d_877_v1_
        d_896_v16_ = rhs198_
        d_877_v1_ = rhs199_
        r0 = D0_DC1(p1, default__.safeModuloInt(p1, p1), p1, d_888_v11_)
        r1 = p2
        return r0, r1


class C3(T0):
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    def ctor__(self):
        pass
        pass

    def fm1(self, p0, p1, p2, p3, globalState):
        return (_dafny.Map({_dafny.MultiSet([not(False), True, False]): not(False)})) | (_dafny.Map({_dafny.MultiSet([True, not(True)]): not(True)}))

    def fm2(self, globalState):
        def iife62_():
            coll26_ = _dafny.Set()
            compr_26_: int
            for compr_26_ in _dafny.IntegerRange(24, 182):
                d_898_v0_: int = compr_26_
                if ((24) <= (d_898_v0_)) and ((d_898_v0_) < (182)):
                    coll26_ = coll26_.union(_dafny.Set([default__.safeModuloInt(d_898_v0_, 21)]))
            return _dafny.Set(coll26_)
        source16_ = D1_DC6(len(_dafny.Map({False: len(iife62_()
)})), _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False, False]), _dafny.SeqWithoutIsStrInference([False])]), 423, not(True))
        if source16_.is_DC5:
            d_899___mcc_h0_ = source16_.cf9
            d_900___mcc_h1_ = source16_.cf10
            d_901___mcc_h2_ = source16_.cf11
            d_902___mcc_h3_ = source16_.cf12
            d_903_cf12_ = d_902___mcc_h3_
            d_904_cf11_ = d_901___mcc_h2_
            d_905_cf10_ = d_900___mcc_h1_
            d_906_cf9_ = d_899___mcc_h0_
            return default__.safeDivisionInt(d_906_cf9_, d_906_cf9_)
        elif source16_.is_DC6:
            d_907___mcc_h4_ = source16_.cf13
            d_908___mcc_h5_ = source16_.cf14
            d_909___mcc_h6_ = source16_.cf15
            d_910___mcc_h7_ = source16_.cf16
            d_911_cf16_ = d_910___mcc_h7_
            d_912_cf15_ = d_909___mcc_h6_
            d_913_cf14_ = d_908___mcc_h5_
            d_914_cf13_ = d_907___mcc_h4_
            return d_914_cf13_
        elif True:
            d_915___mcc_h8_ = source16_.cf8
            d_916_cf8_ = d_915___mcc_h8_
            return default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vpngnpve"))), len(_dafny.Map({872: -382})))

    def fm8(self, p0, p1, p2, globalState):
        return default__.safeModuloInt((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yxgl")))) * (-516), (183) * (len((D0_DC0(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ulgojwmoa")))).cf0)))

    def fm9(self, p0, p1, p2, globalState):
        def iife63_():
            def iife65_():
                coll29_ = _dafny.Map()
                compr_29_: int
                for compr_29_ in _dafny.IntegerRange(521, 997):
                    d_917_v1_: int = compr_29_
                    if ((521) <= (d_917_v1_)) and ((d_917_v1_) < (997)):
                        coll29_[default__.safeDivisionInt(d_917_v1_, 433)] = len(_dafny.SeqWithoutIsStrInference([False]))
                return _dafny.Map(coll29_)
            coll27_ = _dafny.Map()
            def iife64_():
                coll28_ = _dafny.Map()
                compr_28_: int
                for compr_28_ in _dafny.IntegerRange(521, 997):
                    d_917_v1_: int = compr_28_
                    if ((521) <= (d_917_v1_)) and ((d_917_v1_) < (997)):
                        coll28_[default__.safeDivisionInt(d_917_v1_, 433)] = len(_dafny.SeqWithoutIsStrInference([False]))
                return _dafny.Map(coll28_)
            compr_27_: _dafny.MultiSet
            for compr_27_ in (_dafny.Map({_dafny.MultiSet([True, not(False), True, False, False]): len(_dafny.Map({-510: len(iife64_()
            )}))})).keys.Elements:
                d_918_v0_: _dafny.MultiSet = compr_27_
                if (d_918_v0_) in (_dafny.Map({_dafny.MultiSet([True, not(False), True, False, False]): len(_dafny.Map({-510: len(iife65_()
                )}))})):
                    coll27_[d_918_v0_] = False
            return _dafny.Map(coll27_)
        return D1_DC4(((D1_DC4(iife63_()
)).cf8) | (_dafny.Map({_dafny.MultiSet([False, False, False, False]): False})))

    def m0(self, p0, p1, p2, p3, globalState):
        d_919_v0_: _dafny.Map
        d_919_v0_ = _dafny.Map({p0: 604})
        d_920_v1_: _dafny.Map
        d_920_v1_ = _dafny.Map({p0: len(d_919_v0_)})
        d_921_v2_: _dafny.Seq
        d_921_v2_ = _dafny.SeqWithoutIsStrInference([d_919_v0_])
        (globalState).f22 = ((d_921_v2_).set(default__.safeIndex(p0, len(d_921_v2_)), (D2_DC7(d_920_v1_)).cf17)) <= ((d_921_v2_) + (d_921_v2_))
        d_922_v3_: _dafny.Array
        nw175_ = _dafny.Array(int(0), 6)
        d_922_v3_ = nw175_
        index191_ = default__.safeIndex(663, (d_922_v3_).length(0))
        (d_922_v3_)[index191_] = p0
        index192_ = default__.safeIndex(663, (d_922_v3_).length(0))
        (d_922_v3_)[index192_] = default__.safeModuloInt(p2, p0)
        d_923_v4_: _dafny.Seq
        d_923_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iiynvtexy"))
        d_924_v5_: D0
        d_924_v5_ = D0_DC0(d_923_v4_)
        d_925_v6_: str
        d_925_v6_ = _dafny.CodePoint('j')
        d_926_v7_: _dafny.Map
        d_926_v7_ = _dafny.Map({p1: d_923_v4_})
        pat_let_tv19_ = d_923_v4_
        pat_let_tv20_ = d_923_v4_
        pat_let_tv21_ = d_924_v5_
        pat_let_tv22_ = d_923_v4_
        d_927_v8_: _dafny.Array
        nw176_ = _dafny.Array(None, 16)
        nw176_[int(0)] = d_924_v5_
        nw176_[int(1)] = d_924_v5_
        nw176_[int(2)] = D0_DC0((d_923_v4_).set(default__.safeIndex((d_922_v3_)[default__.safeIndex(663, (d_922_v3_).length(0))], len(d_923_v4_)), d_925_v6_))
        nw176_[int(3)] = d_924_v5_
        nw176_[int(4)] = d_924_v5_
        def iife67_(_pat_let19_0):
            def iife68_(d_928_dt__update__tmp_h0_):
                def iife69_(_pat_let20_0):
                    def iife70_(d_929_dt__update_hcf0_h0_):
                        return D0_DC0(d_929_dt__update_hcf0_h0_)
                    return iife70_(_pat_let20_0)
                return iife69_(pat_let_tv19_)
            return iife68_(_pat_let19_0)
        def iife66_(_pat_let18_0):
            def iife71_(d_930_dt__update__tmp_h1_):
                def iife72_(_pat_let21_0):
                    def iife73_(d_931_dt__update_hcf0_h1_):
                        return D0_DC0(d_931_dt__update_hcf0_h1_)
                    return iife73_(_pat_let21_0)
                return iife72_(pat_let_tv20_)
            return iife71_(_pat_let18_0)
        nw176_[int(5)] = iife66_(iife67_(d_924_v5_))
        nw176_[int(6)] = D0_DC0(default__.fm10(_dafny.Map({(d_922_v3_)[default__.safeIndex(663, (d_922_v3_).length(0))]: p0}), default__.fm11((d_922_v3_)[default__.safeIndex(663, (d_922_v3_).length(0))], globalState), p3, globalState))
        nw176_[int(7)] = d_924_v5_
        nw176_[int(8)] = d_924_v5_
        nw176_[int(9)] = d_924_v5_
        nw176_[int(10)] = d_924_v5_
        def iife74_(_pat_let22_0):
            def iife75_(d_932_dt__update__tmp_h2_):
                def iife76_(_pat_let23_0):
                    def iife77_(d_933_dt__update_hcf0_h2_):
                        return D0_DC0(d_933_dt__update_hcf0_h2_)
                    return iife77_(_pat_let23_0)
                return iife76_((pat_let_tv21_).cf0)
            return iife75_(_pat_let22_0)
        nw176_[int(11)] = iife74_(d_924_v5_)
        nw176_[int(12)] = d_924_v5_
        def iife78_(_pat_let24_0):
            def iife79_(d_934_dt__update__tmp_h3_):
                def iife80_(_pat_let25_0):
                    def iife81_(d_935_dt__update_hcf0_h3_):
                        return D0_DC0(d_935_dt__update_hcf0_h3_)
                    return iife81_(_pat_let25_0)
                return iife80_(pat_let_tv22_)
            return iife79_(_pat_let24_0)
        nw176_[int(13)] = iife78_(D0_DC0(((d_926_v7_)[p3] if (p3) in (d_926_v7_) else d_923_v4_)))
        nw176_[int(14)] = d_924_v5_
        nw176_[int(15)] = d_924_v5_
        d_927_v8_ = nw176_
        index193_ = default__.safeIndex(494, (d_927_v8_).length(0))
        (d_927_v8_)[index193_] = d_924_v5_
        d_936_v9_: _dafny.Seq
        d_936_v9_ = _dafny.SeqWithoutIsStrInference([True, p1, not(p1)])
        pat_let_tv23_ = d_923_v4_
        pat_let_tv24_ = p2
        pat_let_tv25_ = d_923_v4_
        pat_let_tv26_ = d_925_v6_
        index194_ = default__.safeIndex(494, (d_927_v8_).length(0))
        def iife82_(_pat_let26_0):
            def iife83_(d_937_dt__update__tmp_h4_):
                def iife84_(_pat_let27_0):
                    def iife85_(d_938_dt__update_hcf0_h4_):
                        return D0_DC0(d_938_dt__update_hcf0_h4_)
                    return iife85_(_pat_let27_0)
                return iife84_(((pat_let_tv23_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q")))).set(default__.safeIndex(pat_let_tv24_, len((pat_let_tv25_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q"))))), pat_let_tv26_))
            return iife83_(_pat_let26_0)
        (d_927_v8_)[index194_] = iife82_((d_924_v5_ if not((d_936_v9_)[default__.safeIndex(p0, len(d_936_v9_))]) else D0_DC0((d_924_v5_).cf0)))
        d_939_v10_: _dafny.Map
        d_939_v10_ = _dafny.Map({len(d_936_v9_): p3})
        d_940_v11_: _dafny.Map
        d_940_v11_ = _dafny.Map({len((d_939_v10_).set((d_922_v3_)[default__.safeIndex(663, (d_922_v3_).length(0))], p1)): d_922_v3_})
        d_940_v11_ = (d_940_v11_).set(p0, d_922_v3_)
        if (p3) == (not(p1)):
            index195_ = default__.safeIndex(663, (d_922_v3_).length(0))
            (d_922_v3_)[index195_] = len(default__.fm10((d_919_v0_) | (d_920_v1_), _dafny.CodePoint('s'), p3, globalState))
            d_941_v12_: _dafny.MultiSet
            d_941_v12_ = _dafny.MultiSet([p2])
            d_942_v13_: _dafny.Map
            d_942_v13_ = _dafny.Map({d_941_v12_: (_dafny.SeqWithoutIsStrInference([p0 for d_943_i0_ in range(default__.abs(510))])).set(default__.safeIndex(p2, len(_dafny.SeqWithoutIsStrInference([p0 for d_944_i0_ in range(default__.abs(510))]))), p0)})
            d_945_v14_: D1
            d_945_v14_ = D1_DC5(len(d_936_v9_), p1, p1, d_942_v13_)
            pat_let_tv27_ = p3
            pat_let_tv28_ = p2
            pat_let_tv29_ = d_942_v13_
            index196_ = default__.safeIndex(663, (d_922_v3_).length(0))
            def iife86_(_pat_let28_0):
                def iife87_(d_946_dt__update__tmp_h5_):
                    def iife88_(_pat_let29_0):
                        def iife89_(d_947_dt__update_hcf11_h0_):
                            def iife90_(_pat_let30_0):
                                def iife91_(d_948_dt__update_hcf9_h0_):
                                    def iife92_(_pat_let31_0):
                                        def iife93_(d_949_dt__update_hcf12_h0_):
                                            return D1_DC5(d_948_dt__update_hcf9_h0_, (d_946_dt__update__tmp_h5_).cf10, d_947_dt__update_hcf11_h0_, d_949_dt__update_hcf12_h0_)
                                        return iife93_(_pat_let31_0)
                                    return iife92_(pat_let_tv29_)
                                return iife91_(_pat_let30_0)
                            return iife90_(pat_let_tv28_)
                        return iife89_(_pat_let29_0)
                    return iife88_(pat_let_tv27_)
                return iife87_(_pat_let28_0)
            (d_922_v3_)[index196_] = (0) - ((len(d_939_v10_)) - ((iife86_(d_945_v14_)).cf9))
            index197_ = default__.safeIndex(663, (d_922_v3_).length(0))
            (d_922_v3_)[index197_] = p2
            (globalState).f22 = default__.fm12(p1, p1, globalState)
            d_950_v15_: _dafny.Map
            d_950_v15_ = _dafny.Map({D0_DC0(d_923_v4_): p1})
            d_951_v16_: _dafny.Array
            nw177_ = _dafny.Array(None, 21)
            nw177_[int(0)] = p1
            nw177_[int(1)] = p1
            nw177_[int(2)] = ((d_950_v15_)[(d_927_v8_)[default__.safeIndex(494, (d_927_v8_).length(0))]] if ((d_927_v8_)[default__.safeIndex(494, (d_927_v8_).length(0))]) in (d_950_v15_) else p1)
            nw177_[int(3)] = p1
            nw177_[int(4)] = False
            nw177_[int(5)] = p3
            nw177_[int(6)] = True
            nw177_[int(7)] = p3
            nw177_[int(8)] = False
            nw177_[int(9)] = p3
            nw177_[int(10)] = p1
            nw177_[int(11)] = p3
            nw177_[int(12)] = p3
            nw177_[int(13)] = p1
            nw177_[int(14)] = p1
            nw177_[int(15)] = p1
            nw177_[int(16)] = True
            nw177_[int(17)] = p1
            nw177_[int(18)] = p3
            nw177_[int(19)] = p1
            nw177_[int(20)] = False
            d_951_v16_ = nw177_
            d_952_v17_: _dafny.Map
            d_952_v17_ = _dafny.Map({p2: d_951_v16_})
            d_953_v18_: _dafny.Seq
            d_953_v18_ = _dafny.SeqWithoutIsStrInference([d_951_v16_, d_951_v16_, d_951_v16_])
            d_952_v17_ = (d_952_v17_).set(p0, (d_953_v18_)[default__.safeIndex((self).fm2(globalState), len(d_953_v18_))])
        elif True:
            d_954_v19_: C0
            nw178_ = C0()
            nw178_.ctor__(p3)
            d_954_v19_ = nw178_
            d_923_v4_ = _dafny.SeqWithoutIsStrInference([d_925_v6_ for d_955_i1_ in range(default__.abs(133))])
            (globalState).f10 = 718
            if default__.fm12(not(p3), (d_954_v19_).f32, globalState):
                d_956_v20_: _dafny.Set
                d_956_v20_ = _dafny.Set({p3, p3, (d_954_v19_).f32, True, p3})
                d_957_v21_: _dafny.Seq
                d_957_v21_ = _dafny.SeqWithoutIsStrInference([p0])
                index198_ = default__.safeIndex(663, (d_922_v3_).length(0))
                (d_922_v3_)[index198_] = ((len(d_926_v7_)) * (len(d_956_v20_))) + (len((d_957_v21_) + (default__.fm14(False, -248, p2, globalState))))
                d_958_v22_: bool
                d_959_v23_: int
                d_960_v24_: T0
                out5_: bool
                out6_: int
                out7_: T0
                out5_, out6_, out7_ = (self).m4(d_923_v4_, len(d_957_v21_), globalState)
                d_958_v22_ = out5_
                d_959_v23_ = out6_
                d_960_v24_ = out7_
                (globalState).f22 = (d_936_v9_)[default__.safeIndex(p0, len(d_936_v9_))]
                d_961_v25_: _dafny.Map
                d_961_v25_ = _dafny.Map({d_922_v3_: _dafny.CodePoint('u')})
                d_961_v25_ = (d_961_v25_).set(d_922_v3_, d_925_v6_)
                d_962_v26_: C0
                nw179_ = C0()
                nw179_.ctor__(p1)
                d_962_v26_ = nw179_
            elif True:
                d_963_v27_: C0
                nw180_ = C0()
                nw180_.ctor__((_dafny.Map({not(p3): d_936_v9_})) == (_dafny.Map({not((d_936_v9_)[default__.safeIndex(-175, len(d_936_v9_))]): d_936_v9_})))
                d_963_v27_ = nw180_
                (globalState).f22 = not((d_963_v27_).f32)
                (globalState).f13 = p1
                d_939_v10_ = (d_939_v10_).set(p0, not((d_954_v19_).f32))
                (globalState).f10 = default__.safeDivisionInt((len(d_939_v10_)) + (len(d_923_v4_)), ((0) - (p2)) + ((d_922_v3_)[default__.safeIndex(663, (d_922_v3_).length(0))]))
            (globalState).f19 = (d_925_v6_ if (d_954_v19_).f32 else (d_925_v6_ if (d_954_v19_).f32 else d_925_v6_))
        d_964_v28_: _dafny.Array
        nw181_ = _dafny.Array(None, 14)
        nw181_[int(0)] = d_939_v10_
        nw181_[int(1)] = (d_939_v10_) | ((d_939_v10_).set((d_922_v3_)[default__.safeIndex(663, (d_922_v3_).length(0))], True))
        nw181_[int(2)] = _dafny.Map({(0) - (p0): p3})
        nw181_[int(3)] = (d_939_v10_) | (d_939_v10_)
        nw181_[int(4)] = d_939_v10_
        nw181_[int(5)] = _dafny.Map({(d_922_v3_)[default__.safeIndex(663, (d_922_v3_).length(0))]: p1})
        nw181_[int(6)] = d_939_v10_
        nw181_[int(7)] = (d_939_v10_ if p1 else _dafny.Map({(d_922_v3_)[default__.safeIndex(663, (d_922_v3_).length(0))]: p1}))
        nw181_[int(8)] = d_939_v10_
        nw181_[int(9)] = d_939_v10_
        nw181_[int(10)] = (_dafny.Map({(d_922_v3_)[default__.safeIndex(663, (d_922_v3_).length(0))]: not(p1)})) | (d_939_v10_)
        nw181_[int(11)] = (d_939_v10_) | ((d_939_v10_).set(p2, p1))
        nw181_[int(12)] = (d_939_v10_) | ((d_939_v10_).set((d_922_v3_)[default__.safeIndex(663, (d_922_v3_).length(0))], not(p3)))
        nw181_[int(13)] = d_939_v10_
        d_964_v28_ = nw181_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_964_v28_).length(0)):
            d_965_i2_: int = guard_loop_2_
            if (True) and (((0) <= (d_965_i2_)) and ((d_965_i2_) < ((d_964_v28_).length(0)))):
                (d_964_v28_)[(d_965_i2_)] = d_939_v10_

    def m1(self, p0, p1, p2, globalState):
        r0: D0 = D0.default()()
        r1: bool = False
        d_966_v0_: _dafny.Array
        nw182_ = _dafny.Array(_dafny.Array(None, 0), 11)
        d_966_v0_ = nw182_
        d_967_v1_: _dafny.Array
        nw183_ = _dafny.Array(False, 22)
        d_967_v1_ = nw183_
        index199_ = default__.safeIndex(300, (d_966_v0_).length(0))
        (d_966_v0_)[index199_] = d_967_v1_
        index200_ = default__.safeIndex(300, (d_966_v0_).length(0))
        (d_966_v0_)[index200_] = (d_967_v1_ if default__.fm12(p2, p2, globalState) else d_967_v1_)
        d_968_v2_: _dafny.Seq
        d_968_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nnaijwfws"))
        (globalState).f10 = len(d_968_v2_)
        index201_ = default__.safeIndex(582, (d_967_v1_).length(0))
        (d_967_v1_)[index201_] = p2
        index202_ = default__.safeIndex(582, (d_967_v1_).length(0))
        (d_967_v1_)[index202_] = (D1_DC6(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ydkxlx"))), _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p2, p2, p2, p2, p2])]), p1, not(p2))).cf16
        d_969_v3_: C0
        nw184_ = C0()
        nw184_.ctor__(True)
        d_969_v3_ = nw184_
        d_970_i0_: int
        d_970_i0_ = 0
        with _dafny.label("1"):
            while p2:
                with _dafny.c_label("1"):
                    if (d_970_i0_) >= (100):
                        raise _dafny.Break("1")
                    d_970_i0_ = (d_970_i0_) + (1)
                    (globalState).f13 = (d_969_v3_).f32
                    d_971_v4_: C0
                    nw185_ = C0()
                    nw185_.ctor__(p2)
                    d_971_v4_ = nw185_
                    d_972_v5_: str
                    d_972_v5_ = _dafny.CodePoint('y')
                    d_973_v6_: _dafny.MultiSet
                    d_973_v6_ = _dafny.MultiSet([(d_971_v4_).f32, not((d_969_v3_).f32)])
                    d_974_v7_: C0
                    nw186_ = C0()
                    nw186_.ctor__(not(((default__.fm15(d_972_v5_, p1, default__.fm12((d_967_v1_)[default__.safeIndex(582, (d_967_v1_).length(0))], p2, globalState), globalState)).cf21) != (d_973_v6_)))
                    d_974_v7_ = nw186_
                    d_975_v8_: _dafny.Map
                    d_975_v8_ = _dafny.Map({p1: len(_dafny.Map({d_972_v5_: D0_DC2(p1, False)}))})
                    d_976_v9_: _dafny.Seq
                    d_976_v9_ = _dafny.SeqWithoutIsStrInference([p1, ((0) - (p1)) * (p1), p1, (0) - (((d_975_v8_)[p1] if (p1) in (d_975_v8_) else p1)), -191])
                    d_977_v10_: _dafny.Seq
                    d_977_v10_ = _dafny.SeqWithoutIsStrInference([(d_967_v1_)[default__.safeIndex(582, (d_967_v1_).length(0))]])
                    d_978_v11_: _dafny.Set
                    d_978_v11_ = _dafny.Set({(d_967_v1_)[default__.safeIndex(582, (d_967_v1_).length(0))], (d_967_v1_)[default__.safeIndex(582, (d_967_v1_).length(0))], p2})
                    rhs200_ = (p2 if p2 else (d_977_v10_) == (default__.fm16(globalState)))
                    rhs201_ = (d_976_v9_) + ((d_976_v9_) + (d_976_v9_))
                    rhs202_ = p1
                    rhs203_ = p2
                    rhs204_ = len((d_978_v11_).intersection((_dafny.Set({True})) - (d_978_v11_)))
                    lhs184_ = globalState
                    lhs185_ = globalState
                    lhs186_ = globalState
                    lhs187_ = globalState
                    lhs184_.f7 = rhs200_
                    d_976_v9_ = rhs201_
                    lhs185_.f10 = rhs202_
                    lhs186_.f22 = rhs203_
                    lhs187_.f10 = rhs204_
                    pass
            pass
        d_979_v12_: _dafny.Array
        nw187_ = _dafny.Array(_dafny.CodePoint('D'), 9)
        d_979_v12_ = nw187_
        d_980_v13_: str
        d_980_v13_ = _dafny.CodePoint('h')
        index203_ = default__.safeIndex(80, (d_979_v12_).length(0))
        (d_979_v12_)[index203_] = d_980_v13_
        index204_ = default__.safeIndex(80, (d_979_v12_).length(0))
        (d_979_v12_)[index204_] = d_980_v13_
        d_981_v14_: D0
        d_981_v14_ = D0_DC1(p1, p1, (0) - (len(d_968_v2_)), (d_966_v0_)[default__.safeIndex(300, (d_966_v0_).length(0))])
        d_982_v15_: _dafny.Seq
        d_982_v15_ = _dafny.SeqWithoutIsStrInference([d_969_v3_])
        d_983_v16_: _dafny.Set
        d_983_v16_ = _dafny.Set({d_982_v15_})
        pat_let_tv30_ = d_983_v16_
        pat_let_tv31_ = p2
        pat_let_tv32_ = p1
        pat_let_tv33_ = d_967_v1_
        def iife94_(_pat_let32_0):
            def iife95_(d_984_dt__update__tmp_h0_):
                def iife96_(_pat_let33_0):
                    def iife97_(d_985_dt__update_hcf3_h0_):
                        def iife98_(_pat_let34_0):
                            def iife99_(d_986_dt__update_hcf1_h0_):
                                def iife100_(_pat_let35_0):
                                    def iife101_(d_987_dt__update_hcf4_h0_):
                                        return D0_DC1(d_986_dt__update_hcf1_h0_, (d_984_dt__update__tmp_h0_).cf2, d_985_dt__update_hcf3_h0_, d_987_dt__update_hcf4_h0_)
                                    return iife101_(_pat_let35_0)
                                return iife100_(pat_let_tv33_)
                            return iife99_(_pat_let34_0)
                        return iife98_(((0) - (len(pat_let_tv30_)) if pat_let_tv31_ else pat_let_tv32_))
                    return iife97_(_pat_let33_0)
                return iife96_(-614)
            return iife95_(_pat_let32_0)
        r0 = iife94_(d_981_v14_)
        d_988_v17_: _dafny.MultiSet
        d_988_v17_ = _dafny.MultiSet([p1, p1, p1])
        d_989_v18_: _dafny.MultiSet
        d_989_v18_ = _dafny.MultiSet([d_988_v17_])
        r1 = ((d_989_v18_).set(d_988_v17_, default__.abs(p1))).isdisjoint(d_989_v18_)
        return r0, r1

    def m4(self, p0, p1, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: T0 = None
        d_990_v0_: bool
        d_990_v0_ = False
        d_991_i0_: int
        d_991_i0_ = 0
        with _dafny.label("2"):
            while d_990_v0_:
                with _dafny.c_label("2"):
                    if (d_991_i0_) >= (100):
                        raise _dafny.Break("2")
                    d_991_i0_ = (d_991_i0_) + (1)
                    d_992_v1_: _dafny.Array
                    nw188_ = _dafny.Array(int(0), 19)
                    d_992_v1_ = nw188_
                    d_993_v2_: _dafny.MultiSet
                    d_993_v2_ = _dafny.MultiSet([d_992_v1_, d_992_v1_])
                    d_994_v3_: _dafny.Seq
                    d_994_v3_ = _dafny.SeqWithoutIsStrInference([d_990_v0_, (d_993_v2_) != (d_993_v2_), d_990_v0_, not((d_990_v0_ if default__.fm12(d_990_v0_, d_990_v0_, globalState) else d_990_v0_))])
                    d_994_v3_ = d_994_v3_
                    d_995_v4_: _dafny.Seq
                    d_995_v4_ = _dafny.SeqWithoutIsStrInference([d_994_v3_, d_994_v3_, d_994_v3_, default__.fm16(globalState)])
                    d_996_v5_: D1
                    d_996_v5_ = D1_DC6(p1, d_995_v4_, p1, d_990_v0_)
                    index205_ = default__.safeIndex(201, (d_992_v1_).length(0))
                    (d_992_v1_)[index205_] = len(p0)
                    d_997_v6_: _dafny.Seq
                    d_997_v6_ = _dafny.SeqWithoutIsStrInference([p1])
                    d_998_v7_: _dafny.Set
                    d_998_v7_ = _dafny.Set({p1, p1})
                    d_999_v8_: str
                    d_999_v8_ = _dafny.CodePoint('a')
                    d_1000_v9_: _dafny.Map
                    d_1000_v9_ = _dafny.Map({default__.fm17((0) - (p1), d_990_v0_, d_999_v8_, globalState): d_997_v6_})
                    d_1001_v10_: D1
                    d_1001_v10_ = D1_DC5(p1, not(d_990_v0_), d_990_v0_, d_1000_v9_)
                    index206_ = default__.safeIndex(201, (d_992_v1_).length(0))
                    rhs205_ = d_990_v0_
                    rhs206_ = ((d_997_v6_)[default__.safeIndex(p1, len(d_997_v6_))]) not in (d_998_v7_)
                    rhs207_ = default__.fm12(d_990_v0_, d_990_v0_, globalState)
                    rhs208_ = d_996_v5_
                    rhs209_ = (d_1001_v10_).cf9
                    lhs188_ = globalState
                    lhs189_ = d_992_v1_
                    lhs190_ = default__.safeIndex(201, (d_992_v1_).length(0))
                    lhs188_.f7 = rhs205_
                    r0 = rhs206_
                    d_990_v0_ = rhs207_
                    d_996_v5_ = rhs208_
                    lhs189_[lhs190_] = rhs209_
                    d_1002_v11_: _dafny.Array
                    nw189_ = _dafny.Array(False, 24)
                    d_1002_v11_ = nw189_
                    index207_ = default__.safeIndex(349, (d_1002_v11_).length(0))
                    (d_1002_v11_)[index207_] = True
                    index208_ = default__.safeIndex(349, (d_1002_v11_).length(0))
                    (d_1002_v11_)[index208_] = d_990_v0_
                    d_1003_v12_: D3
                    d_1003_v12_ = D3_DC11((d_992_v1_)[default__.safeIndex(201, (d_992_v1_).length(0))])
                    d_1004_v13_: _dafny.Set
                    d_1004_v13_ = _dafny.Set({d_992_v1_, d_992_v1_})
                    pat_let_tv34_ = d_992_v1_
                    pat_let_tv35_ = d_992_v1_
                    index209_ = default__.safeIndex(201, (d_992_v1_).length(0))
                    def iife102_(_pat_let36_0):
                        def iife103_(d_1005_dt__update__tmp_h0_):
                            def iife104_(_pat_let37_0):
                                def iife105_(d_1006_dt__update_hcf22_h0_):
                                    return D3_DC11(d_1006_dt__update_hcf22_h0_)
                                return iife105_(_pat_let37_0)
                            return iife104_((0) - ((pat_let_tv35_)[default__.safeIndex(201, (pat_let_tv34_).length(0))]))
                        return iife103_(_pat_let36_0)
                    (d_992_v1_)[index209_] = (0) - (((iife102_(d_1003_v12_)).cf22) + ((p1) - (len(d_1004_v13_))))
                    pass
            pass
        hi3_ = p1
        for d_1007_i1_ in range(p1, hi3_):
            d_1008_v14_: _dafny.Map
            d_1008_v14_ = _dafny.Map({p1: d_990_v0_})
            d_1009_v15_: _dafny.MultiSet
            d_1009_v15_ = _dafny.MultiSet([(0) - (len(d_1008_v14_))])
            d_1010_v16_: _dafny.Seq
            d_1010_v16_ = _dafny.SeqWithoutIsStrInference([d_990_v0_, d_990_v0_])
            d_1011_v17_: _dafny.Set
            d_1011_v17_ = _dafny.Set({d_990_v0_})
            d_1012_v18_: _dafny.Array
            nw190_ = _dafny.Array(None, 28)
            nw190_[int(0)] = (d_1009_v15_).isdisjoint(_dafny.MultiSet([default__.fm0(p1, globalState), d_1007_i1_]))
            nw190_[int(1)] = d_990_v0_
            nw190_[int(2)] = (default__.fm0(p1, globalState)) < (d_1007_i1_)
            nw190_[int(3)] = d_990_v0_
            nw190_[int(4)] = d_990_v0_
            nw190_[int(5)] = d_990_v0_
            nw190_[int(6)] = d_990_v0_
            nw190_[int(7)] = d_990_v0_
            nw190_[int(8)] = d_990_v0_
            nw190_[int(9)] = (d_1007_i1_) <= (p1)
            nw190_[int(10)] = d_990_v0_
            nw190_[int(11)] = (p1) >= (d_1007_i1_)
            nw190_[int(12)] = not(d_990_v0_)
            nw190_[int(13)] = (d_990_v0_) in (d_1010_v16_)
            nw190_[int(14)] = d_990_v0_
            nw190_[int(15)] = default__.fm12(not(d_990_v0_), d_990_v0_, globalState)
            nw190_[int(16)] = d_990_v0_
            nw190_[int(17)] = d_990_v0_
            nw190_[int(18)] = not (d_990_v0_) or (d_990_v0_)
            nw190_[int(19)] = d_990_v0_
            nw190_[int(20)] = d_990_v0_
            nw190_[int(21)] = d_990_v0_
            nw190_[int(22)] = d_990_v0_
            nw190_[int(23)] = (d_990_v0_) == (d_990_v0_)
            nw190_[int(24)] = not(True)
            nw190_[int(25)] = (not(d_990_v0_)) or (d_990_v0_)
            nw190_[int(26)] = d_990_v0_
            nw190_[int(27)] = (d_1011_v17_).issubset(d_1011_v17_)
            d_1012_v18_ = nw190_
            index210_ = default__.safeIndex(212, (d_1012_v18_).length(0))
            (d_1012_v18_)[index210_] = False
            d_1013_v19_: _dafny.Set
            d_1013_v19_ = _dafny.Set({p1})
            index211_ = default__.safeIndex(212, (d_1012_v18_).length(0))
            rhs210_ = default__.fm12(d_990_v0_, (d_1007_i1_) < (len(d_1013_v19_)), globalState)
            rhs211_ = False
            rhs212_ = p1
            lhs191_ = d_1012_v18_
            lhs192_ = default__.safeIndex(212, (d_1012_v18_).length(0))
            lhs193_ = globalState
            r0 = rhs210_
            lhs191_[lhs192_] = rhs211_
            lhs193_.f10 = rhs212_
            d_1014_v20_: _dafny.Array
            def lambda41_(d_1015_v14_):
                def lambda42_(d_1016_i2_):
                    return d_1015_v14_

                return lambda42_

            init21_ = lambda41_(d_1008_v14_)
            nw191_ = _dafny.Array(None, 15)
            for i0_21_ in range(nw191_.length(0)):
                nw191_[i0_21_] = init21_(i0_21_)
            d_1014_v20_ = nw191_
            index212_ = default__.safeIndex(420, (d_1014_v20_).length(0))
            (d_1014_v20_)[index212_] = default__.fm18(p1, d_990_v0_, globalState)
            index213_ = default__.safeIndex(420, (d_1014_v20_).length(0))
            (d_1014_v20_)[index213_] = d_1008_v14_
            d_1017_v21_: _dafny.Array
            def lambda43_(d_1018_p0_):
                def lambda44_(d_1019_i3_):
                    return d_1018_p0_

                return lambda44_

            init22_ = lambda43_(p0)
            nw192_ = _dafny.Array(None, 16)
            for i0_22_ in range(nw192_.length(0)):
                nw192_[i0_22_] = init22_(i0_22_)
            d_1017_v21_ = nw192_
            d_1020_v22_: _dafny.Map
            d_1020_v22_ = _dafny.Map({-100: p0})
            index214_ = default__.safeIndex(897, (d_1017_v21_).length(0))
            (d_1017_v21_)[index214_] = (((d_1020_v22_)[d_1007_i1_] if (d_1007_i1_) in (d_1020_v22_) else p0)).set(default__.safeIndex(p1, len(((d_1020_v22_)[d_1007_i1_] if (d_1007_i1_) in (d_1020_v22_) else p0))), default__.fm11(p1, globalState))
            d_1021_v23_: str
            d_1021_v23_ = _dafny.CodePoint('t')
            index215_ = default__.safeIndex(897, (d_1017_v21_).length(0))
            (d_1017_v21_)[index215_] = (_dafny.SeqWithoutIsStrInference([d_1021_v23_])) + (p0)
            d_1022_v24_: C0
            nw193_ = C0()
            nw193_.ctor__(default__.fm12((d_1012_v18_)[default__.safeIndex(212, (d_1012_v18_).length(0))], (d_1012_v18_)[default__.safeIndex(212, (d_1012_v18_).length(0))], globalState))
            d_1022_v24_ = nw193_
        hi4_ = p1
        for d_1023_i4_ in range(len(p0), hi4_):
            d_1024_v25_: _dafny.Map
            d_1024_v25_ = _dafny.Map({d_990_v0_: d_990_v0_})
            d_1024_v25_ = (d_1024_v25_).set(d_990_v0_, False)
            d_1025_v26_: str
            d_1025_v26_ = _dafny.CodePoint('q')
            d_1026_v27_: _dafny.MultiSet
            d_1026_v27_ = _dafny.MultiSet([d_1025_v26_, d_1025_v26_, d_1025_v26_])
            r1 = ((d_1026_v27_).cardinality if d_990_v0_ else 331)
            d_1027_v28_: _dafny.Array
            nw194_ = _dafny.Array(int(0), 8)
            d_1027_v28_ = nw194_
            d_1028_v29_: _dafny.Map
            d_1028_v29_ = _dafny.Map({d_1027_v28_: True})
            d_1029_v30_: _dafny.MultiSet
            d_1029_v30_ = _dafny.MultiSet([d_990_v0_])
            d_1028_v29_ = (d_1028_v29_).set(d_1027_v28_, (d_1029_v30_).isdisjoint(d_1029_v30_))
            d_1030_v31_: _dafny.Seq
            d_1030_v31_ = _dafny.SeqWithoutIsStrInference([d_990_v0_])
            if (d_1030_v31_)[default__.safeIndex(d_1023_i4_, len(d_1030_v31_))]:
                (globalState).f10 = (p1) + (default__.fm0(d_1023_i4_, globalState))
                d_1031_v32_: _dafny.Array
                nw195_ = _dafny.Array(_dafny.Map({}), 12)
                d_1031_v32_ = nw195_
                d_1032_v33_: _dafny.Map
                d_1032_v33_ = _dafny.Map({d_1031_v32_: (p1) * (d_1023_i4_)})
                (globalState).f10 = (0) - (((d_1032_v33_)[d_1031_v32_] if (d_1031_v32_) in (d_1032_v33_) else (d_1023_i4_) + (d_1023_i4_)))
                index216_ = default__.safeIndex(697, (d_1027_v28_).length(0))
                (d_1027_v28_)[index216_] = p1
                index217_ = default__.safeIndex(697, (d_1027_v28_).length(0))
                (d_1027_v28_)[index217_] = p1
                (globalState).f22 = not(default__.fm12(not(d_990_v0_), d_990_v0_, globalState))
                d_1033_v34_: _dafny.Array
                def lambda45_(d_1034_v25_):
                    def lambda46_(d_1035_i5_):
                        return d_1034_v25_

                    return lambda46_

                init23_ = lambda45_(d_1024_v25_)
                nw196_ = _dafny.Array(None, 27)
                for i0_23_ in range(nw196_.length(0)):
                    nw196_[i0_23_] = init23_(i0_23_)
                d_1033_v34_ = nw196_
                d_1036_v35_: _dafny.Seq
                d_1036_v35_ = _dafny.SeqWithoutIsStrInference([d_1024_v25_, d_1024_v25_, d_1024_v25_, d_1024_v25_, d_1024_v25_])
                index218_ = default__.safeIndex(507, (d_1033_v34_).length(0))
                (d_1033_v34_)[index218_] = (d_1036_v35_)[default__.safeIndex(p1, len(d_1036_v35_))]
                index219_ = default__.safeIndex(507, (d_1033_v34_).length(0))
                (d_1033_v34_)[index219_] = (d_1024_v25_ if default__.fm12(d_990_v0_, d_990_v0_, globalState) else d_1024_v25_)
            elif True:
                d_1037_v36_: _dafny.Array
                def lambda47_(d_1038_v0_):
                    def lambda48_(d_1039_i6_):
                        return d_1038_v0_

                    return lambda48_

                init24_ = lambda47_(d_990_v0_)
                nw197_ = _dafny.Array(None, 7)
                for i0_24_ in range(nw197_.length(0)):
                    nw197_[i0_24_] = init24_(i0_24_)
                d_1037_v36_ = nw197_
                index220_ = default__.safeIndex(514, (d_1037_v36_).length(0))
                (d_1037_v36_)[index220_] = d_990_v0_
                index221_ = default__.safeIndex(514, (d_1037_v36_).length(0))
                (d_1037_v36_)[index221_] = True
                d_1025_v26_ = d_1025_v26_
                d_1040_v37_: _dafny.Map
                d_1040_v37_ = _dafny.Map({d_1025_v26_: d_990_v0_})
                d_1041_v38_: _dafny.Map
                d_1041_v38_ = _dafny.Map({d_1023_i4_: True})
                index222_ = default__.safeIndex(514, (d_1037_v36_).length(0))
                (d_1037_v36_)[index222_] = (default__.fm12((d_1037_v36_)[default__.safeIndex(514, (d_1037_v36_).length(0))], ((d_1040_v37_)[d_1025_v26_] if (d_1025_v26_) in (d_1040_v37_) else ((d_1041_v38_)[p1] if (p1) in (d_1041_v38_) else True)), globalState)) or (d_990_v0_)
                d_990_v0_ = True
                d_1042_v39_: _dafny.Array
                nw198_ = _dafny.Array(_dafny.CodePoint('D'), 15)
                d_1042_v39_ = nw198_
                index223_ = default__.safeIndex(740, (d_1042_v39_).length(0))
                (d_1042_v39_)[index223_] = d_1025_v26_
                d_1043_v40_: D4
                d_1043_v40_ = D4_DC13((p0)[default__.safeIndex(d_1023_i4_, len(p0))])
                index224_ = default__.safeIndex(740, (d_1042_v39_).length(0))
                (d_1042_v39_)[index224_] = (d_1043_v40_).cf24
        (globalState).f13 = not(d_990_v0_)
        d_1044_v41_: _dafny.Seq
        d_1044_v41_ = _dafny.SeqWithoutIsStrInference([False])
        if ((d_1044_v41_)[default__.safeIndex(p1, len(d_1044_v41_))]) == (default__.fm12(d_990_v0_, d_990_v0_, globalState)):
            d_1045_v42_: D0
            d_1045_v42_ = D0_DC2(p1, d_990_v0_)
            d_1046_v43_: _dafny.MultiSet
            d_1046_v43_ = _dafny.MultiSet([d_1045_v42_])
            (globalState).f22 = (d_1046_v43_).ispropersubset(d_1046_v43_)
            (globalState).f10 = (p1 if d_990_v0_ else p1)
            d_1047_v44_: D3
            d_1047_v44_ = D3_DC11((p1) + (p1))
            source17_ = d_1047_v44_
            if source17_.is_DC11:
                d_1048___mcc_h0_ = source17_.cf22
                d_1049_cf22_ = d_1048___mcc_h0_
                d_1050_v45_: _dafny.Map
                d_1050_v45_ = _dafny.Map({default__.fm11(d_1049_cf22_, globalState): d_1044_v41_})
                d_1050_v45_ = d_1050_v45_
                d_1051_v46_: _dafny.Map
                d_1051_v46_ = _dafny.Map({p1: False})
                d_1052_v47_: _dafny.Map
                d_1052_v47_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference([p1 for d_1053_i7_ in range(default__.abs(733))])) + (default__.fm14(d_990_v0_, d_1049_cf22_, default__.fm0(d_1049_cf22_, globalState), globalState)): d_1051_v46_})
                d_1054_v48_: _dafny.Seq
                d_1054_v48_ = _dafny.SeqWithoutIsStrInference([p1])
                d_1052_v47_ = (d_1052_v47_).set(d_1054_v48_, d_1051_v46_)
                d_1055_v49_: C0
                nw199_ = C0()
                nw199_.ctor__(d_990_v0_)
                d_1055_v49_ = nw199_
                d_1056_v50_: _dafny.Map
                d_1056_v50_ = _dafny.Map({(default__.fm12(d_990_v0_, d_990_v0_, globalState) if d_990_v0_ else d_990_v0_): (len(_dafny.SeqWithoutIsStrInference([d_1055_v49_, d_1055_v49_])) if d_990_v0_ else (self).fm8(p1, d_1049_cf22_, p1, globalState))})
                d_1056_v50_ = (d_1056_v50_).set(not(d_990_v0_), len((p0) + (p0)))
                d_1057_v51_: _dafny.Map
                d_1057_v51_ = _dafny.Map({d_990_v0_: (d_1055_v49_).f32})
                d_1058_v52_: _dafny.Map
                d_1058_v52_ = d_1057_v51_
                d_1059_v53_: _dafny.MultiSet
                d_1059_v53_ = _dafny.MultiSet([d_1049_cf22_, d_1049_cf22_])
                d_1060_v54_: _dafny.Array
                nw200_ = _dafny.Array(None, 14)
                nw200_[int(0)] = _dafny.Map({False: d_990_v0_})
                nw200_[int(1)] = _dafny.Map({d_990_v0_: (d_1055_v49_).f32})
                nw200_[int(2)] = d_1057_v51_
                nw200_[int(3)] = d_1057_v51_
                nw200_[int(4)] = d_1057_v51_
                nw200_[int(5)] = d_1057_v51_
                nw200_[int(6)] = d_1057_v51_
                nw200_[int(7)] = (_dafny.Map({(d_1055_v49_).f32: (d_1055_v49_).f32})) | (d_1057_v51_)
                nw200_[int(8)] = d_1057_v51_
                nw200_[int(9)] = d_1057_v51_
                nw200_[int(10)] = (d_1058_v52_)
                nw200_[int(11)] = (_dafny.Map({(d_1055_v49_).f32: (d_1055_v49_).f32})).set((d_1055_v49_).f32, d_990_v0_)
                nw200_[int(12)] = default__.fm19(d_1059_v53_, p1, globalState)
                nw200_[int(13)] = (d_1057_v51_) | (d_1057_v51_)
                d_1060_v54_ = nw200_
                d_1060_v54_ = d_1060_v54_
            elif source17_.is_DC12:
                d_1061___mcc_h1_ = source17_.cf23
                d_1062_cf23_ = d_1061___mcc_h1_
                d_1063_v55_: _dafny.Array
                nw201_ = _dafny.Array(int(0), 8)
                d_1063_v55_ = nw201_
                index225_ = default__.safeIndex(538, (d_1063_v55_).length(0))
                (d_1063_v55_)[index225_] = p1
                d_1064_v56_: _dafny.MultiSet
                d_1064_v56_ = _dafny.MultiSet([p0, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_1065_i8_ in range(default__.abs(702))])])
                d_1066_v57_: C0
                nw202_ = C0()
                nw202_.ctor__(False)
                d_1066_v57_ = nw202_
                index226_ = default__.safeIndex(538, (d_1063_v55_).length(0))
                rhs213_ = (d_1044_v41_)[default__.safeIndex(((d_1064_v56_) | (d_1064_v56_)).cardinality, len(d_1044_v41_))]
                rhs214_ = (len(_dafny.SeqWithoutIsStrInference([d_1066_v57_]))) + (p1)
                rhs215_ = (0) - (p1)
                rhs216_ = p1
                rhs217_ = p1
                lhs194_ = globalState
                lhs195_ = globalState
                lhs196_ = globalState
                lhs197_ = d_1063_v55_
                lhs198_ = default__.safeIndex(538, (d_1063_v55_).length(0))
                lhs194_.f13 = rhs213_
                lhs195_.f10 = rhs214_
                r1 = rhs215_
                lhs196_.f10 = rhs216_
                lhs197_[lhs198_] = rhs217_
                (globalState).f10 = p1
                d_1067_v58_: _dafny.Map
                d_1067_v58_ = _dafny.Map({False: (d_1063_v55_)[default__.safeIndex(538, (d_1063_v55_).length(0))]})
                d_1068_v59_: _dafny.Map
                d_1068_v59_ = _dafny.Map({p1: ((d_1067_v58_)[True] if (True) in (d_1067_v58_) else 891)})
                d_1068_v59_ = (d_1068_v59_).set(len(p0), (d_1063_v55_)[default__.safeIndex(538, (d_1063_v55_).length(0))])
                (globalState).f7 = d_990_v0_
            elif True:
                d_1069___mcc_h2_ = source17_.cf21
                d_1070_cf21_ = d_1069___mcc_h2_
                d_1071_v60_: _dafny.Set
                d_1071_v60_ = _dafny.Set({len(p0)})
                d_1072_v61_: _dafny.Seq
                d_1072_v61_ = _dafny.SeqWithoutIsStrInference([(0) - ((self).fm8(p1, len(p0), 790, globalState)), 100, len(d_1071_v60_)])
                d_1073_v62_: _dafny.Map
                d_1073_v62_ = _dafny.Map({_dafny.MultiSet(d_1072_v61_): d_1072_v61_})
                d_1074_v63_: D1
                d_1074_v63_ = D1_DC5(p1, d_990_v0_, d_990_v0_, default__.fm20(p1, d_990_v0_, (0) - (p1), globalState))
                d_1075_v64_: _dafny.Seq
                d_1075_v64_ = _dafny.SeqWithoutIsStrInference([D1_DC5(len(p0), d_990_v0_, d_990_v0_, d_1073_v62_), d_1074_v63_, D1_DC5(p1, not(d_990_v0_), not(d_990_v0_), d_1073_v62_), d_1074_v63_])
                d_1075_v64_ = d_1075_v64_
                d_1076_v65_: _dafny.Array
                nw203_ = _dafny.Array(_dafny.Set({}), 25)
                d_1076_v65_ = nw203_
                d_1077_v66_: _dafny.Map
                d_1077_v66_ = _dafny.Map({p1: (0) - (p1)})
                index227_ = default__.safeIndex(709, (d_1076_v65_).length(0))
                def iife106_():
                    coll30_ = _dafny.Set()
                    compr_30_: int
                    for compr_30_ in (d_1077_v66_).keys.Elements:
                        d_1078_v67_: int = compr_30_
                        if (d_1078_v67_) in (d_1077_v66_):
                            coll30_ = coll30_.union(_dafny.Set([(d_1078_v67_) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mbi"))))]))
                    return _dafny.Set(coll30_)
                (d_1076_v65_)[index227_] = (d_1071_v60_) | (iife106_()
                )
                index228_ = default__.safeIndex(709, (d_1076_v65_).length(0))
                (d_1076_v65_)[index228_] = _dafny.Set({p1})
                d_1079_v68_: str
                d_1079_v68_ = _dafny.CodePoint('d')
                d_1080_v69_: D4
                d_1080_v69_ = D4_DC13(d_1079_v68_)
                d_1081_v70_: _dafny.Seq
                d_1081_v70_ = _dafny.SeqWithoutIsStrInference([d_1080_v69_])
                r1 = len(((d_1081_v70_) + (d_1081_v70_)) + (_dafny.SeqWithoutIsStrInference([D4_DC13(d_1079_v68_)])))
                d_1082_v71_: _dafny.Set
                d_1082_v71_ = _dafny.Set({d_1079_v68_, d_1079_v68_})
                r0 = (d_1082_v71_).issubset(d_1082_v71_)
            (globalState).f22 = (d_990_v0_) or (d_990_v0_)
            (globalState).f10 = p1
        elif True:
            (globalState).f7 = d_990_v0_
            d_1083_v72_: _dafny.Array
            nw204_ = _dafny.Array(None, 8)
            nw204_[int(0)] = d_990_v0_
            nw204_[int(1)] = default__.fm12(d_990_v0_, d_990_v0_, globalState)
            nw204_[int(2)] = d_990_v0_
            nw204_[int(3)] = (d_1044_v41_)[default__.safeIndex(p1, len(d_1044_v41_))]
            nw204_[int(4)] = (d_990_v0_ if (d_1044_v41_)[default__.safeIndex(default__.fm0(p1, globalState), len(d_1044_v41_))] else d_990_v0_)
            nw204_[int(5)] = d_990_v0_
            nw204_[int(6)] = (not(d_990_v0_)) and (d_990_v0_)
            nw204_[int(7)] = d_990_v0_
            d_1083_v72_ = nw204_
            d_1084_v73_: _dafny.Map
            d_1084_v73_ = _dafny.Map({d_990_v0_: d_990_v0_})
            index229_ = default__.safeIndex(557, (d_1083_v72_).length(0))
            (d_1083_v72_)[index229_] = (d_1084_v73_) == (d_1084_v73_)
            index230_ = default__.safeIndex(557, (d_1083_v72_).length(0))
            (d_1083_v72_)[index230_] = d_990_v0_
            d_1085_v74_: _dafny.Seq
            d_1085_v74_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rlvkgkw"))
            rhs218_ = d_1085_v74_
            rhs219_ = (p0) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "opjwmmjub")))
            d_1085_v74_ = rhs218_
            d_1085_v74_ = rhs219_
            d_1086_v75_: C0
            nw205_ = C0()
            nw205_.ctor__((d_1083_v72_)[default__.safeIndex(557, (d_1083_v72_).length(0))])
            d_1086_v75_ = nw205_
            d_1087_v76_: T0
            nw206_ = C1()
            nw206_.ctor__()
            d_1087_v76_ = nw206_
            d_1088_v77_: str
            d_1088_v77_ = _dafny.CodePoint('h')
            d_1089_v78_: _dafny.Map
            d_1089_v78_ = _dafny.Map({d_1087_v76_: len((default__.fm10(_dafny.Map({p1: p1}), d_1088_v77_, (d_1086_v75_).f32, globalState)) + (p0))})
            d_1089_v78_ = (d_1089_v78_).set(d_1087_v76_, p1)
        d_1090_v79_: _dafny.MultiSet
        d_1090_v79_ = _dafny.MultiSet([p1])
        d_1091_v80_: _dafny.Map
        d_1091_v80_ = _dafny.Map({len(d_1044_v41_): d_990_v0_})
        d_1092_v81_: _dafny.Map
        d_1092_v81_ = _dafny.Map({(d_1090_v79_).set((0) - (len(d_1091_v80_)), default__.abs(-908)): _dafny.SeqWithoutIsStrInference([p1 for d_1093_i9_ in range(default__.abs(687))])})
        d_1094_v82_: D1
        d_1094_v82_ = D1_DC5(len(d_1044_v41_), (p1) <= (p1), d_990_v0_, d_1092_v81_)
        source18_ = d_1094_v82_
        if source18_.is_DC5:
            d_1095___mcc_h3_ = source18_.cf9
            d_1096___mcc_h4_ = source18_.cf10
            d_1097___mcc_h5_ = source18_.cf11
            d_1098___mcc_h6_ = source18_.cf12
            d_1099_cf12_ = d_1098___mcc_h6_
            d_1100_cf11_ = d_1097___mcc_h5_
            d_1101_cf10_ = d_1096___mcc_h4_
            d_1102_cf9_ = d_1095___mcc_h3_
            if d_1101_cf10_:
                d_1103_v83_: _dafny.MultiSet
                d_1103_v83_ = _dafny.MultiSet([d_1100_cf11_, d_1101_cf10_])
                d_1104_v84_: _dafny.Seq
                d_1104_v84_ = _dafny.SeqWithoutIsStrInference([d_1103_v83_, d_1103_v83_])
                d_1105_v85_: _dafny.Array
                nw207_ = _dafny.Array(int(0), 21)
                d_1105_v85_ = nw207_
                rhs220_ = p1
                rhs221_ = ((d_1104_v84_) + (d_1104_v84_)).set(default__.safeIndex((_dafny.MultiSet([d_1105_v85_, d_1105_v85_, d_1105_v85_])).cardinality, len((d_1104_v84_) + (d_1104_v84_))), _dafny.MultiSet([d_990_v0_, d_990_v0_]))
                d_1102_cf9_ = rhs220_
                d_1104_v84_ = rhs221_
                d_1106_v86_: _dafny.Map
                d_1106_v86_ = _dafny.Map({d_990_v0_: d_1100_cf11_})
                d_1107_v87_: str
                d_1107_v87_ = _dafny.CodePoint('p')
                d_1108_v88_: _dafny.Seq
                d_1108_v88_ = _dafny.SeqWithoutIsStrInference([len(d_1106_v86_), d_1102_cf9_, len(_dafny.Set({d_1107_v87_})), d_1102_cf9_])
                d_1109_v89_: _dafny.Seq
                d_1109_v89_ = _dafny.SeqWithoutIsStrInference([d_1108_v88_, d_1108_v88_])
                d_1110_v90_: _dafny.Map
                d_1110_v90_ = _dafny.Map({((d_1109_v89_)[default__.safeIndex(d_1102_cf9_, len(d_1109_v89_))] if d_1100_cf11_ else d_1108_v88_): d_1100_cf11_})
                d_1110_v90_ = d_1110_v90_
                d_1111_v91_: _dafny.MultiSet
                d_1111_v91_ = _dafny.MultiSet([d_1107_v87_, d_1107_v87_])
                d_1112_v92_: _dafny.Set
                d_1112_v92_ = _dafny.Set({d_1102_cf9_, (d_1111_v91_).cardinality})
                d_1113_v93_: D0
                d_1113_v93_ = D0_DC2(p1, False)
                d_1114_v94_: _dafny.Map
                d_1114_v94_ = _dafny.Map({not(d_1101_cf10_): p1})
                d_1115_v95_: _dafny.Array
                nw208_ = _dafny.Array(None, 22)
                nw208_[int(0)] = d_1101_cf10_
                nw208_[int(1)] = d_1100_cf11_
                nw208_[int(2)] = (_dafny.MultiSet([default__.fm11(len(d_1044_v41_), globalState)])).ispropersubset(_dafny.MultiSet([d_1107_v87_]))
                nw208_[int(3)] = True
                nw208_[int(4)] = d_1101_cf10_
                nw208_[int(5)] = (len(d_1112_v92_)) < (p1)
                nw208_[int(6)] = d_1101_cf10_
                nw208_[int(7)] = ((d_1106_v86_)[False] if (False) in (d_1106_v86_) else d_1101_cf10_)
                nw208_[int(8)] = d_1100_cf11_
                nw208_[int(9)] = (d_1113_v93_).cf6
                nw208_[int(10)] = d_990_v0_
                nw208_[int(11)] = default__.fm12(d_990_v0_, d_990_v0_, globalState)
                nw208_[int(12)] = default__.fm12(d_990_v0_, False, globalState)
                nw208_[int(13)] = default__.fm12(d_1100_cf11_, not(d_1100_cf11_), globalState)
                nw208_[int(14)] = d_1100_cf11_
                nw208_[int(15)] = d_990_v0_
                nw208_[int(16)] = (d_1114_v94_) == (_dafny.Map({d_990_v0_: d_1102_cf9_}))
                nw208_[int(17)] = d_990_v0_
                nw208_[int(18)] = not(d_990_v0_)
                nw208_[int(19)] = d_990_v0_
                nw208_[int(20)] = (d_1090_v79_) != (d_1090_v79_)
                nw208_[int(21)] = d_1100_cf11_
                d_1115_v95_ = nw208_
                index231_ = default__.safeIndex(525, (d_1115_v95_).length(0))
                (d_1115_v95_)[index231_] = d_1101_cf10_
                index232_ = default__.safeIndex(525, (d_1115_v95_).length(0))
                (d_1115_v95_)[index232_] = not((d_1103_v83_).isdisjoint(d_1103_v83_))
                d_1116_v96_: _dafny.Set
                d_1116_v96_ = _dafny.Set({d_990_v0_, d_990_v0_})
                d_1117_v97_: _dafny.Map
                d_1117_v97_ = _dafny.Map({d_1115_v95_: d_1116_v96_})
                d_1118_v98_: D8
                d_1118_v98_ = D8_DC21(d_1044_v41_)
                d_1119_v99_: _dafny.Map
                d_1119_v99_ = _dafny.Map({d_1117_v97_: ((d_1118_v98_).cf35) + (d_1044_v41_)})
                d_1119_v99_ = (d_1119_v99_).set(d_1117_v97_, d_1044_v41_)
                (globalState).f24 = (d_1112_v92_).intersection(d_1112_v92_)
            elif True:
                d_1120_v100_: T1
                nw209_ = C2()
                nw209_.ctor__((p0) + (p0))
                d_1120_v100_ = nw209_
                d_1121_v101_: _dafny.Array
                nw210_ = _dafny.Array(_dafny.CodePoint('D'), 3)
                d_1121_v101_ = nw210_
                d_1122_v102_: str
                d_1122_v102_ = _dafny.CodePoint('i')
                index233_ = default__.safeIndex(930, (d_1121_v101_).length(0))
                (d_1121_v101_)[index233_] = d_1122_v102_
                d_1123_v103_: _dafny.Array
                nw211_ = _dafny.Array(_dafny.Map({}), 28)
                d_1123_v103_ = nw211_
                index234_ = default__.safeIndex(930, (d_1121_v101_).length(0))
                rhs222_ = default__.safeModuloInt(len(p0), p1)
                rhs223_ = d_1120_v100_
                rhs224_ = d_1122_v102_
                rhs225_ = d_1123_v103_
                lhs199_ = globalState
                lhs200_ = d_1121_v101_
                lhs201_ = default__.safeIndex(930, (d_1121_v101_).length(0))
                lhs199_.f10 = rhs222_
                d_1120_v100_ = rhs223_
                lhs200_[lhs201_] = rhs224_
                d_1123_v103_ = rhs225_
                d_1124_v104_: _dafny.Seq
                d_1124_v104_ = _dafny.SeqWithoutIsStrInference([d_1102_cf9_, p1])
                (globalState).f10 = ((0) - ((d_1124_v104_)[default__.safeIndex(d_1102_cf9_, len(d_1124_v104_))])) * (d_1102_cf9_)
                r1 = (0) - (p1)
                (globalState).f10 = d_1102_cf9_
                d_1125_v105_: _dafny.Map
                d_1125_v105_ = _dafny.Map({d_1100_cf11_: d_1122_v102_})
                d_1126_v106_: _dafny.Array
                nw212_ = _dafny.Array(None, 3)
                nw212_[int(0)] = (d_1125_v105_) | (d_1125_v105_)
                nw212_[int(1)] = _dafny.Map({d_1100_cf11_: (d_1121_v101_)[default__.safeIndex(930, (d_1121_v101_).length(0))]})
                nw212_[int(2)] = d_1125_v105_
                d_1126_v106_ = nw212_
                index235_ = default__.safeIndex(669, (d_1126_v106_).length(0))
                (d_1126_v106_)[index235_] = d_1125_v105_
                d_1127_v107_: _dafny.Map
                d_1127_v107_ = _dafny.Map({True: d_1101_cf10_})
                index236_ = default__.safeIndex(669, (d_1126_v106_).length(0))
                (d_1126_v106_)[index236_] = (d_1125_v105_) | (_dafny.Map({((d_1127_v107_)[d_990_v0_] if (d_990_v0_) in (d_1127_v107_) else d_1101_cf10_): d_1122_v102_}))
            d_1128_v108_: _dafny.Array
            nw213_ = _dafny.Array(int(0), 6)
            d_1128_v108_ = nw213_
            rhs226_ = d_1128_v108_
            rhs227_ = (d_1102_cf9_) >= (p1)
            rhs228_ = p1
            rhs229_ = d_1101_cf10_
            rhs230_ = d_1102_cf9_
            lhs202_ = globalState
            lhs203_ = globalState
            lhs204_ = globalState
            lhs205_ = globalState
            d_1128_v108_ = rhs226_
            lhs202_.f13 = rhs227_
            lhs203_.f10 = rhs228_
            lhs204_.f13 = rhs229_
            lhs205_.f10 = rhs230_
            d_1091_v80_ = (d_1091_v80_).set(235, (d_1100_cf11_) == (d_1101_cf10_))
            d_1129_v109_: _dafny.Array
            d_1129_v109_ = d_1128_v108_
            source19_ = d_1128_v108_
            d_1130___mcc_h12_ = source19_
            d_1131_cf28_ = d_1130___mcc_h12_
            (globalState).f10 = p1
            d_1132_v110_: _dafny.Map
            d_1132_v110_ = _dafny.Map({len(_dafny.Set({d_1128_v108_})): (0) - (d_1102_cf9_)})
            d_1133_v112_: _dafny.MultiSet
            d_1133_v112_ = _dafny.MultiSet([d_990_v0_, d_1100_cf11_])
            d_1134_v113_: _dafny.Seq
            d_1134_v113_ = _dafny.SeqWithoutIsStrInference([d_1133_v112_, _dafny.MultiSet([True]), _dafny.MultiSet(d_1044_v41_)])
            def iife107_():
                coll31_ = _dafny.Map()
                compr_31_: _dafny.MultiSet
                for compr_31_ in (d_1134_v113_).Elements:
                    d_1135_v111_: _dafny.MultiSet = compr_31_
                    if (d_1135_v111_) in (d_1134_v113_):
                        coll31_[d_1135_v111_] = d_1102_cf9_
                return _dafny.Map(coll31_)
            d_1132_v110_ = (d_1132_v110_).set(default__.safeDivisionInt((self).fm8(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "isvwjc"))), -69, d_1102_cf9_, globalState), p1), len(iife107_()
            ))
            d_1136_v114_: _dafny.Array
            def lambda49_(d_1137_cf11_):
                def lambda50_(d_1138_i10_):
                    return d_1137_cf11_

                return lambda50_

            init25_ = lambda49_(d_1100_cf11_)
            nw214_ = _dafny.Array(None, 5)
            for i0_25_ in range(nw214_.length(0)):
                nw214_[i0_25_] = init25_(i0_25_)
            d_1136_v114_ = nw214_
            d_1139_v115_: _dafny.Map
            d_1139_v115_ = _dafny.Map({d_1102_cf9_: d_1136_v114_})
            d_1140_v116_: _dafny.Seq
            d_1140_v116_ = _dafny.SeqWithoutIsStrInference([d_1102_cf9_, d_1102_cf9_])
            d_1139_v115_ = (d_1139_v115_).set(len(d_1140_v116_), d_1136_v114_)
            (globalState).f10 = d_1102_cf9_
        elif source18_.is_DC6:
            d_1141___mcc_h7_ = source18_.cf13
            d_1142___mcc_h8_ = source18_.cf14
            d_1143___mcc_h9_ = source18_.cf15
            d_1144___mcc_h10_ = source18_.cf16
            d_1145_cf16_ = d_1144___mcc_h10_
            d_1146_cf15_ = d_1143___mcc_h9_
            d_1147_cf14_ = d_1142___mcc_h8_
            d_1148_cf13_ = d_1141___mcc_h7_
            d_1149_v117_: _dafny.Seq
            d_1149_v117_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "esqyiagdp"))
            d_1150_v118_: str
            d_1150_v118_ = _dafny.CodePoint('w')
            d_1151_v119_: _dafny.Map
            d_1151_v119_ = _dafny.Map({d_990_v0_: (d_1149_v117_).set(default__.safeIndex(d_1148_cf13_, len(d_1149_v117_)), d_1150_v118_)})
            d_1149_v117_ = (d_1149_v117_).set(default__.safeIndex(len((d_1149_v117_ if not(d_990_v0_) else ((d_1151_v119_)[d_990_v0_] if (d_990_v0_) in (d_1151_v119_) else p0))), len(d_1149_v117_)), d_1150_v118_)
            (globalState).f13 = (d_1145_cf16_ if d_990_v0_ else (default__.fm0(d_1148_cf13_, globalState)) < (p1))
            d_1152_v120_: _dafny.Map
            d_1152_v120_ = _dafny.Map({not((d_1044_v41_)[default__.safeIndex(271, len(d_1044_v41_))]): True})
            d_1153_v121_: _dafny.Set
            d_1153_v121_ = _dafny.Set({d_1145_cf16_, d_990_v0_})
            (globalState).f13 = ((d_1152_v120_)[(d_1153_v121_).issubset(d_1153_v121_)] if ((d_1153_v121_).issubset(d_1153_v121_)) in (d_1152_v120_) else d_990_v0_)
            d_1154_v122_: _dafny.Array
            nw215_ = _dafny.Array(False, 8)
            d_1154_v122_ = nw215_
            d_1155_v123_: _dafny.Map
            d_1155_v123_ = _dafny.Map({default__.safeDivisionInt(824, d_1148_cf13_): d_1154_v122_})
            d_1155_v123_ = (d_1155_v123_).set(len(d_1152_v120_), d_1154_v122_)
        elif True:
            d_1156___mcc_h11_ = source18_.cf8
            d_1157_cf8_ = d_1156___mcc_h11_
            d_1158_v124_: _dafny.MultiSet
            d_1158_v124_ = _dafny.MultiSet([d_990_v0_])
            d_1159_v125_: _dafny.Seq
            d_1159_v125_ = _dafny.SeqWithoutIsStrInference([p1])
            d_1160_v126_: _dafny.Map
            d_1160_v126_ = _dafny.Map({not((d_1158_v124_).ispropersubset(d_1158_v124_)): d_1159_v125_})
            d_1160_v126_ = d_1160_v126_
            d_1161_v127_: _dafny.Array
            def lambda51_(d_1162_v0_):
                def lambda52_(d_1163_i11_):
                    return d_1162_v0_

                return lambda52_

            init26_ = lambda51_(d_990_v0_)
            nw216_ = _dafny.Array(None, 15)
            for i0_26_ in range(nw216_.length(0)):
                nw216_[i0_26_] = init26_(i0_26_)
            d_1161_v127_ = nw216_
            index237_ = default__.safeIndex(345, (d_1161_v127_).length(0))
            (d_1161_v127_)[index237_] = (p1) >= ((d_1159_v125_)[default__.safeIndex(p1, len(d_1159_v125_))])
            index238_ = default__.safeIndex(345, (d_1161_v127_).length(0))
            (d_1161_v127_)[index238_] = (_dafny.SeqWithoutIsStrInference([d_990_v0_, d_990_v0_])) != ((_dafny.SeqWithoutIsStrInference([d_990_v0_]) if d_990_v0_ else d_1044_v41_))
            (globalState).f13 = (d_1161_v127_)[default__.safeIndex(345, (d_1161_v127_).length(0))]
            d_1164_v128_: _dafny.Array
            nw217_ = _dafny.Array(_dafny.Array(None, 0), 5)
            d_1164_v128_ = nw217_
            d_1165_v129_: _dafny.Set
            d_1165_v129_ = _dafny.Set({True})
            d_1166_v130_: _dafny.Array
            nw218_ = _dafny.Array(None, 9)
            nw218_[int(0)] = d_1090_v79_
            nw218_[int(1)] = d_1090_v79_
            nw218_[int(2)] = d_1090_v79_
            nw218_[int(3)] = _dafny.MultiSet([p1, p1])
            nw218_[int(4)] = _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([961 for d_1167_i12_ in range(default__.abs(30))]))
            nw218_[int(5)] = d_1090_v79_
            nw218_[int(6)] = d_1090_v79_
            nw218_[int(7)] = d_1090_v79_
            nw218_[int(8)] = _dafny.MultiSet([len(d_1165_v129_), p1])
            d_1166_v130_ = nw218_
            index239_ = default__.safeIndex(19, (d_1164_v128_).length(0))
            (d_1164_v128_)[index239_] = d_1166_v130_
            index240_ = default__.safeIndex(19, (d_1164_v128_).length(0))
            nw219_ = _dafny.Array(_dafny.MultiSet({}), 11)
            (d_1164_v128_)[index240_] = nw219_
        d_1168_v131_: _dafny.Array
        nw220_ = _dafny.Array(int(0), 4)
        d_1168_v131_ = nw220_
        d_1169_v132_: _dafny.MultiSet
        d_1169_v132_ = _dafny.MultiSet([d_1168_v131_, d_1168_v131_])
        r0 = (d_1169_v132_).ispropersubset((d_1169_v132_) | (d_1169_v132_))
        r1 = (p1) - (367)
        d_1170_v133_: T0
        nw221_ = C1()
        nw221_.ctor__()
        d_1170_v133_ = nw221_
        r2 = d_1170_v133_
        return r0, r1, r2

    def m5(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: int = int(0)
        (globalState).f22 = False
        (globalState).f22 = p3
        d_1171_v0_: _dafny.Set
        d_1171_v0_ = _dafny.Set({p0, (0) - (p0), p0})
        hi5_ = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_1172_i1_ in range(default__.abs(212))]))
        for d_1173_i0_ in range((len(d_1171_v0_)) - (p0), hi5_):
            d_1174_v1_: _dafny.Seq
            d_1174_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rfyol"))
            d_1174_v1_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "arqqntbg"))) + (d_1174_v1_)) + (d_1174_v1_)
            d_1175_v2_: _dafny.Array
            nw222_ = _dafny.Array(_dafny.Set({}), 4)
            d_1175_v2_ = nw222_
            d_1176_v3_: _dafny.Map
            d_1176_v3_ = _dafny.Map({default__.fm12(p3, not(p2), globalState): d_1175_v2_})
            d_1176_v3_ = (d_1176_v3_).set(p3, d_1175_v2_)
            rhs231_ = p2
            rhs232_ = default__.safeModuloInt(len((d_1174_v1_) + (d_1174_v1_)), d_1173_i0_)
            lhs206_ = globalState
            lhs207_ = globalState
            lhs206_.f13 = rhs231_
            lhs207_.f10 = rhs232_
            d_1177_v4_: _dafny.Array
            def lambda53_(d_1178_p3_):
                def lambda54_(d_1179_i2_):
                    return d_1178_p3_

                return lambda54_

            init27_ = lambda53_(p3)
            nw223_ = _dafny.Array(None, 23)
            for i0_27_ in range(nw223_.length(0)):
                nw223_[i0_27_] = init27_(i0_27_)
            d_1177_v4_ = nw223_
            index241_ = default__.safeIndex(69, (d_1177_v4_).length(0))
            (d_1177_v4_)[index241_] = (d_1171_v0_).issubset(d_1171_v0_)
            index242_ = default__.safeIndex(69, (d_1177_v4_).length(0))
            (d_1177_v4_)[index242_] = p3
        d_1180_v5_: str
        d_1180_v5_ = _dafny.CodePoint('s')
        d_1181_v7_: _dafny.MultiSet
        d_1181_v7_ = _dafny.MultiSet([p3])
        d_1182_v8_: _dafny.Seq
        d_1182_v8_ = _dafny.SeqWithoutIsStrInference([d_1181_v7_])
        d_1183_v9_: D1
        def iife108_():
            coll32_ = _dafny.Map()
            compr_32_: _dafny.MultiSet
            for compr_32_ in (d_1182_v8_).Elements:
                d_1184_v6_: _dafny.MultiSet = compr_32_
                if (d_1184_v6_) in (d_1182_v8_):
                    coll32_[d_1184_v6_] = p3
            return _dafny.Map(coll32_)
        d_1183_v9_ = D1_DC4(iife108_()
)
        d_1185_v10_: _dafny.Map
        d_1185_v10_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bk"))).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bk")))), d_1180_v5_): d_1183_v9_})
        d_1186_v11_: _dafny.Map
        d_1186_v11_ = _dafny.Map({p2: d_1185_v10_})
        d_1187_v12_: _dafny.Seq
        d_1187_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rwbrc"))
        d_1186_v11_ = (d_1186_v11_).set(not(((d_1187_v12_).set(default__.safeIndex(p0, len(d_1187_v12_)), d_1180_v5_)) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wouynh")))), (d_1185_v10_).set(d_1187_v12_, d_1183_v9_))
        d_1188_i3_: int
        d_1188_i3_ = 0
        with _dafny.label("3"):
            while p1:
                with _dafny.c_label("3"):
                    if (d_1188_i3_) >= (100):
                        raise _dafny.Break("3")
                    d_1188_i3_ = (d_1188_i3_) + (1)
                    d_1189_v13_: _dafny.Map
                    d_1189_v13_ = _dafny.Map({p0: len(d_1187_v12_)})
                    d_1190_v14_: _dafny.Seq
                    d_1190_v14_ = _dafny.SeqWithoutIsStrInference([p1, p1])
                    d_1189_v13_ = default__.fm33((p3) and (True), (p2) == (p2), (d_1190_v14_)[default__.safeIndex(p0, len(d_1190_v14_))], p1, globalState)
                    d_1191_v15_: C0
                    nw224_ = C0()
                    nw224_.ctor__(p2)
                    d_1191_v15_ = nw224_
                    d_1192_v16_: C0
                    nw225_ = C0()
                    nw225_.ctor__((p0) != (p0))
                    d_1192_v16_ = nw225_
                    d_1193_v17_: _dafny.MultiSet
                    d_1193_v17_ = _dafny.MultiSet([p0, len(d_1187_v12_)])
                    d_1194_v18_: _dafny.Map
                    d_1194_v18_ = _dafny.Map({(d_1190_v14_) == (d_1190_v14_): d_1193_v17_})
                    d_1195_v19_: C2
                    nw226_ = C2()
                    nw226_.ctor__(_dafny.SeqWithoutIsStrInference([d_1180_v5_ for d_1196_i4_ in range(default__.abs(45))]))
                    d_1195_v19_ = nw226_
                    d_1197_v20_: _dafny.Map
                    d_1197_v20_ = _dafny.Map({d_1195_v19_: p0})
                    d_1198_v21_: _dafny.Map
                    d_1198_v21_ = _dafny.Map({p3: len(d_1195_v19_.f33)})
                    rhs233_ = len(d_1194_v18_)
                    rhs234_ = ((d_1197_v20_)[d_1195_v19_] if (d_1195_v19_) in (d_1197_v20_) else p0)
                    rhs235_ = (0) - (len(((_dafny.Map({p1: (0) - (p0)})) | (d_1198_v21_)).set((d_1192_v16_).f32, p0)))
                    lhs208_ = globalState
                    lhs209_ = globalState
                    lhs210_ = globalState
                    lhs208_.f10 = rhs233_
                    lhs209_.f10 = rhs234_
                    lhs210_.f10 = rhs235_
                    pass
            pass
        d_1199_v22_: _dafny.Array
        nw227_ = _dafny.Array(int(0), 23)
        d_1199_v22_ = nw227_
        index243_ = default__.safeIndex(537, (d_1199_v22_).length(0))
        (d_1199_v22_)[index243_] = p0
        index244_ = default__.safeIndex(537, (d_1199_v22_).length(0))
        (d_1199_v22_)[index244_] = p0
        r0 = p0
        r1 = (d_1199_v22_)[default__.safeIndex(537, (d_1199_v22_).length(0))]
        return r0, r1


class C4(T0, T1):
    def  __init__(self):
        self._f31: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    def ctor__(self, f31):
        (self)._f31 = f31

    def fm1(self, p0, p1, p2, p3, globalState):
        if (False if False else not(True)):
            return (_dafny.Map({_dafny.MultiSet([not(True)]): True})) | (_dafny.Map({_dafny.MultiSet([True, True]): True}))
        elif True:
            return (_dafny.Map({_dafny.MultiSet([False]): not(True)})) | (_dafny.Map({_dafny.MultiSet([True]): False}))

    def fm2(self, globalState):
        return (0) - (((973) + ((0) - (len(_dafny.SeqWithoutIsStrInference([631])))) if (False if False else False) else len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t')]) if True else _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_1200_i0_ in range(default__.abs(357))])))))

    def fm5(self, p0, p1, p2, globalState):
        if not((_dafny.MultiSet([-735])) != (_dafny.MultiSet([669]))):
            return True
        elif not(True):
            return False
        elif True:
            return False

    def fm6(self, globalState):
        return -665

    def m0(self, p0, p1, p2, p3, globalState):
        rhs236_ = p3
        rhs237_ = p1
        rhs238_ = p2
        lhs211_ = globalState
        lhs212_ = globalState
        lhs213_ = globalState
        lhs211_.f13 = rhs236_
        lhs212_.f22 = rhs237_
        lhs213_.f10 = rhs238_
        d_1201_i0_: int
        d_1201_i0_ = 0
        with _dafny.label("4"):
            while (p2) <= ((363 if p1 else p2)):
                with _dafny.c_label("4"):
                    if (d_1201_i0_) >= (100):
                        raise _dafny.Break("4")
                    d_1201_i0_ = (d_1201_i0_) + (1)
                    d_1202_v0_: _dafny.Seq
                    d_1202_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gjwrovu"))
                    d_1203_v1_: str
                    d_1203_v1_ = _dafny.CodePoint('o')
                    d_1204_v2_: _dafny.Map
                    d_1204_v2_ = _dafny.Map({len((d_1202_v0_).set(default__.safeIndex((0) - (len(_dafny.Map({p2: 993}))), len(d_1202_v0_)), d_1203_v1_)): _dafny.MultiSet([p2])})
                    d_1204_v2_ = (d_1204_v2_).set(p2, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([979, p0, p0]))).set(p0, default__.abs(p2)))
                    d_1205_v3_: _dafny.Set
                    d_1205_v3_ = _dafny.Set({p3, p3})
                    d_1206_v4_: T0
                    nw228_ = C3()
                    nw228_.ctor__()
                    d_1206_v4_ = nw228_
                    d_1207_v5_: _dafny.Seq
                    d_1207_v5_ = _dafny.SeqWithoutIsStrInference([d_1206_v4_])
                    d_1208_v6_: _dafny.MultiSet
                    d_1208_v6_ = _dafny.MultiSet([d_1205_v3_, d_1205_v3_])
                    if ((_dafny.MultiSet([default__.fm7(p1, p3, globalState), d_1205_v3_, _dafny.Set({p1})])).set(_dafny.Set({p3}), default__.abs(len(d_1207_v5_)))).issubset(d_1208_v6_):
                        d_1209_v7_: D2
                        d_1209_v7_ = D2_DC8(p3, p1)
                        d_1210_v8_: D2
                        d_1210_v8_ = D2_DC9(d_1209_v7_)
                        d_1211_v9_: _dafny.Map
                        d_1211_v9_ = _dafny.Map({d_1203_v1_: default__.safeDivisionInt(p0, len((_dafny.Map({p0: d_1210_v8_})).set(p2, d_1210_v8_)))})
                        d_1211_v9_ = (d_1211_v9_).set(d_1203_v1_, p2)
                        d_1212_v10_: _dafny.Array
                        nw229_ = _dafny.Array(int(0), 11)
                        d_1212_v10_ = nw229_
                        index245_ = default__.safeIndex(954, (d_1212_v10_).length(0))
                        (d_1212_v10_)[index245_] = p0
                        index246_ = default__.safeIndex(954, (d_1212_v10_).length(0))
                        (d_1212_v10_)[index246_] = default__.safeDivisionInt((0) - (p0), p0)
                        d_1212_v10_ = d_1212_v10_
                        d_1213_v11_: _dafny.Seq
                        d_1213_v11_ = _dafny.SeqWithoutIsStrInference([742, (self).fm2(globalState), (d_1212_v10_)[default__.safeIndex(954, (d_1212_v10_).length(0))]])
                        d_1214_v12_: _dafny.Map
                        d_1214_v12_ = _dafny.Map({p1: len(d_1213_v11_)})
                        d_1215_v13_: _dafny.Set
                        d_1215_v13_ = _dafny.Set({_dafny.Map({p3: p2})})
                        (globalState).f7 = ((_dafny.Set({d_1214_v12_})) - (d_1215_v13_)) != (_dafny.Set({d_1214_v12_, d_1214_v12_, d_1214_v12_, _dafny.Map({p1: 208}), d_1214_v12_}))
                        d_1216_v14_: _dafny.Array
                        nw230_ = _dafny.Array(_dafny.Array(None, 0), 7)
                        d_1216_v14_ = nw230_
                        d_1217_v15_: _dafny.Array
                        nw231_ = _dafny.Array(False, 22)
                        d_1217_v15_ = nw231_
                        index247_ = default__.safeIndex(476, (d_1216_v14_).length(0))
                        (d_1216_v14_)[index247_] = d_1217_v15_
                        index248_ = default__.safeIndex(476, (d_1216_v14_).length(0))
                        (d_1216_v14_)[index248_] = d_1217_v15_
                    elif True:
                        d_1218_v16_: _dafny.Array
                        def lambda55_(d_1219_p1_):
                            def lambda56_(d_1220_i1_):
                                return d_1219_p1_

                            return lambda56_

                        init28_ = lambda55_(p1)
                        nw232_ = _dafny.Array(None, 15)
                        for i0_28_ in range(nw232_.length(0)):
                            nw232_[i0_28_] = init28_(i0_28_)
                        d_1218_v16_ = nw232_
                        index249_ = default__.safeIndex(900, (d_1218_v16_).length(0))
                        (d_1218_v16_)[index249_] = p3
                        index250_ = default__.safeIndex(900, (d_1218_v16_).length(0))
                        (d_1218_v16_)[index250_] = True
                        d_1221_v17_: _dafny.Array
                        def lambda57_(d_1222_v0_):
                            def lambda58_(d_1223_i2_):
                                return (d_1223_i2_) + (len(d_1222_v0_))

                            return lambda58_

                        init29_ = lambda57_(d_1202_v0_)
                        nw233_ = _dafny.Array(None, 14)
                        for i0_29_ in range(nw233_.length(0)):
                            nw233_[i0_29_] = init29_(i0_29_)
                        d_1221_v17_ = nw233_
                        index251_ = default__.safeIndex(86, (d_1221_v17_).length(0))
                        (d_1221_v17_)[index251_] = p2
                        index252_ = default__.safeIndex(86, (d_1221_v17_).length(0))
                        rhs239_ = p0
                        rhs240_ = (d_1203_v1_) not in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xnwqqkdl")))
                        rhs241_ = p2
                        lhs214_ = globalState
                        lhs215_ = globalState
                        lhs216_ = d_1221_v17_
                        lhs217_ = default__.safeIndex(86, (d_1221_v17_).length(0))
                        lhs214_.f10 = rhs239_
                        lhs215_.f13 = rhs240_
                        lhs216_[lhs217_] = rhs241_
                        d_1224_v18_: C0
                        nw234_ = C0()
                        nw234_.ctor__(False)
                        d_1224_v18_ = nw234_
                        d_1225_v19_: _dafny.Map
                        d_1225_v19_ = _dafny.Map({False: d_1203_v1_})
                        d_1226_v20_: _dafny.Seq
                        d_1226_v20_ = _dafny.SeqWithoutIsStrInference([p1])
                        d_1227_v21_: _dafny.Seq
                        d_1227_v21_ = _dafny.SeqWithoutIsStrInference([(d_1226_v20_)[default__.safeIndex(p2, len(d_1226_v20_))]])
                        d_1225_v19_ = (d_1225_v19_).set((d_1227_v21_) != (d_1227_v21_), d_1203_v1_)
                        (globalState).f10 = (0) - (p0)
                    d_1228_v23_: _dafny.Map
                    d_1228_v23_ = _dafny.Map({d_1203_v1_: p2})
                    d_1229_v24_: _dafny.MultiSet
                    d_1229_v24_ = _dafny.MultiSet([p1, p1])
                    d_1230_v25_: _dafny.Map
                    d_1230_v25_ = _dafny.Map({d_1229_v24_: p3})
                    d_1231_v26_: D1
                    d_1231_v26_ = D1_DC4(d_1230_v25_)
                    pat_let_tv36_ = d_1230_v25_
                    d_1232_v27_: _dafny.Map
                    def iife109_():
                        coll33_ = _dafny.Map()
                        compr_33_: str
                        for compr_33_ in (d_1228_v23_).keys.Elements:
                            d_1233_v22_: str = compr_33_
                            if (d_1233_v22_) in (d_1228_v23_):
                                coll33_[d_1233_v22_] = p2
                        return _dafny.Map(coll33_)
                    def iife110_(_pat_let38_0):
                        def iife111_(d_1234_dt__update__tmp_h0_):
                            def iife112_(_pat_let39_0):
                                def iife113_(d_1235_dt__update_hcf8_h0_):
                                    return D1_DC4(d_1235_dt__update_hcf8_h0_)
                                return iife113_(_pat_let39_0)
                            return iife112_(pat_let_tv36_)
                        return iife111_(_pat_let38_0)
                    d_1232_v27_ = _dafny.Map({len(iife109_()
                    ): iife110_(d_1231_v26_)})
                    d_1236_v28_: _dafny.MultiSet
                    d_1236_v28_ = _dafny.MultiSet([d_1232_v27_, _dafny.Map({p0: d_1231_v26_}), d_1232_v27_])
                    d_1237_v29_: _dafny.Map
                    d_1237_v29_ = _dafny.Map({False: (d_1236_v28_) | (d_1236_v28_)})
                    d_1238_v30_: _dafny.Set
                    d_1238_v30_ = _dafny.Set({d_1202_v0_})
                    d_1239_v31_: _dafny.Seq
                    d_1239_v31_ = _dafny.SeqWithoutIsStrInference([d_1232_v27_])
                    d_1237_v29_ = (d_1237_v29_).set(not((d_1238_v30_).ispropersubset(d_1238_v30_)), _dafny.MultiSet((d_1239_v31_) + (d_1239_v31_)))
                    if p3:
                        (globalState).f10 = p2
                        d_1240_v32_: _dafny.Array
                        def lambda59_(d_1241_i3_):
                            return True

                        init30_ = lambda59_
                        nw235_ = _dafny.Array(None, 8)
                        for i0_30_ in range(nw235_.length(0)):
                            nw235_[i0_30_] = init30_(i0_30_)
                        d_1240_v32_ = nw235_
                        index253_ = default__.safeIndex(49, (d_1240_v32_).length(0))
                        (d_1240_v32_)[index253_] = p3
                        index254_ = default__.safeIndex(49, (d_1240_v32_).length(0))
                        (d_1240_v32_)[index254_] = (p3) and (p3)
                        d_1242_v33_: C0
                        nw236_ = C0()
                        nw236_.ctor__(p1)
                        d_1242_v33_ = nw236_
                        (globalState).f10 = 383
                        d_1243_v34_: _dafny.Seq
                        d_1243_v34_ = _dafny.SeqWithoutIsStrInference([((0) - ((0) - ((self).fm2(globalState)))) * (p2), len((d_1202_v0_ if not((d_1240_v32_)[default__.safeIndex(49, (d_1240_v32_).length(0))]) else d_1202_v0_)), default__.safeModuloInt((0) - (p2), p2)])
                        d_1243_v34_ = d_1243_v34_
                    elif True:
                        d_1244_v35_: C3
                        nw237_ = C3()
                        nw237_.ctor__()
                        d_1244_v35_ = nw237_
                        d_1245_v36_: _dafny.Seq
                        d_1245_v36_ = _dafny.SeqWithoutIsStrInference([d_1202_v0_])
                        d_1246_v37_: _dafny.Map
                        d_1246_v37_ = _dafny.Map({p3: p0})
                        d_1247_v38_: _dafny.Array
                        nw238_ = _dafny.Array(None, 16)
                        nw238_[int(0)] = default__.fm10(_dafny.Map({len(d_1202_v0_): len(d_1205_v3_)}), d_1203_v1_, p1, globalState)
                        nw238_[int(1)] = (d_1202_v0_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gdhniefl")))
                        nw238_[int(2)] = d_1202_v0_
                        nw238_[int(3)] = d_1202_v0_
                        nw238_[int(4)] = d_1202_v0_
                        nw238_[int(5)] = d_1202_v0_
                        nw238_[int(6)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "arq"))
                        nw238_[int(7)] = d_1202_v0_
                        nw238_[int(8)] = (_dafny.SeqWithoutIsStrInference([d_1203_v1_ for d_1248_i4_ in range(default__.abs(777))])) + (d_1202_v0_)
                        nw238_[int(9)] = d_1202_v0_
                        nw238_[int(10)] = ((d_1245_v36_)[default__.safeIndex(len(d_1246_v37_), len(d_1245_v36_))]) + (d_1202_v0_)
                        nw238_[int(11)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pl"))
                        nw238_[int(12)] = d_1202_v0_
                        nw238_[int(13)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "abm"))
                        nw238_[int(14)] = (_dafny.SeqWithoutIsStrInference([d_1203_v1_ for d_1249_i5_ in range(default__.abs(381))])) + (d_1202_v0_)
                        nw238_[int(15)] = d_1202_v0_
                        d_1247_v38_ = nw238_
                        index255_ = default__.safeIndex(692, (d_1247_v38_).length(0))
                        (d_1247_v38_)[index255_] = (d_1202_v0_) + (d_1202_v0_)
                        index256_ = default__.safeIndex(692, (d_1247_v38_).length(0))
                        (d_1247_v38_)[index256_] = (_dafny.SeqWithoutIsStrInference([d_1203_v1_ for d_1250_i6_ in range(default__.abs(914))])) + ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_1251_i7_ in range(default__.abs(577))])) + (d_1202_v0_))
                        (globalState).f10 = p0
                        d_1252_v39_: _dafny.Map
                        d_1252_v39_ = _dafny.Map({p0: p3})
                        (globalState).f10 = len((d_1252_v39_).set(591, p1))
                        d_1253_v40_: C1
                        nw239_ = C1()
                        nw239_.ctor__()
                        d_1253_v40_ = nw239_
                    pass
            pass
        hi6_ = -238
        for d_1254_i8_ in range(p0, hi6_):
            d_1255_v41_: _dafny.MultiSet
            d_1255_v41_ = _dafny.MultiSet([p1, p3])
            d_1256_v42_: D3
            d_1256_v42_ = D3_DC10(d_1255_v41_)
            d_1256_v42_ = (d_1256_v42_ if p3 else d_1256_v42_)
            d_1257_v43_: T0
            nw240_ = C1()
            nw240_.ctor__()
            d_1257_v43_ = nw240_
            d_1257_v43_ = d_1257_v43_
            d_1258_v44_: _dafny.Seq
            d_1258_v44_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v"))
            d_1259_v45_: _dafny.Map
            d_1259_v45_ = _dafny.Map({d_1258_v44_: len(d_1258_v44_)})
            d_1260_v46_: _dafny.Map
            d_1260_v46_ = _dafny.Map({p3: p1})
            d_1261_v47_: _dafny.Map
            d_1261_v47_ = _dafny.Map({d_1254_i8_: p2})
            d_1262_v48_: _dafny.MultiSet
            d_1262_v48_ = _dafny.MultiSet([D2_DC7(d_1261_v47_)])
            d_1263_v49_: _dafny.Seq
            d_1263_v49_ = _dafny.SeqWithoutIsStrInference([d_1262_v48_])
            d_1264_v50_: _dafny.Seq
            d_1264_v50_ = _dafny.SeqWithoutIsStrInference([len(d_1260_v46_), ((d_1263_v49_)[default__.safeIndex(p2, len(d_1263_v49_))]).cardinality])
            d_1265_v51_: _dafny.Array
            nw241_ = _dafny.Array(None, 15)
            nw241_[int(0)] = p0
            nw241_[int(1)] = ((d_1259_v45_)[d_1258_v44_] if (d_1258_v44_) in (d_1259_v45_) else p2)
            nw241_[int(2)] = default__.safeDivisionInt(682, -113)
            nw241_[int(3)] = d_1254_i8_
            nw241_[int(4)] = default__.fm0(d_1254_i8_, globalState)
            nw241_[int(5)] = ((d_1255_v41_).cardinality) * (p2)
            nw241_[int(6)] = 584
            nw241_[int(7)] = p0
            nw241_[int(8)] = p0
            nw241_[int(9)] = (d_1254_i8_) - ((0) - (-7))
            nw241_[int(10)] = p2
            nw241_[int(11)] = p2
            nw241_[int(12)] = (d_1257_v43_).fm2(globalState)
            nw241_[int(13)] = (d_1264_v50_)[default__.safeIndex(p0, len(d_1264_v50_))]
            nw241_[int(14)] = p0
            d_1265_v51_ = nw241_
            index257_ = default__.safeIndex(233, (d_1265_v51_).length(0))
            (d_1265_v51_)[index257_] = (p2 if p1 else p2)
            index258_ = default__.safeIndex(233, (d_1265_v51_).length(0))
            rhs242_ = (d_1254_i8_) + (p0)
            rhs243_ = p2
            lhs218_ = globalState
            lhs219_ = d_1265_v51_
            lhs220_ = default__.safeIndex(233, (d_1265_v51_).length(0))
            lhs218_.f10 = rhs242_
            lhs219_[lhs220_] = rhs243_
            if not(not(p1)):
                (globalState).f10 = default__.safeDivisionInt((d_1265_v51_)[default__.safeIndex(233, (d_1265_v51_).length(0))], d_1254_i8_)
                (globalState).f7 = not(p1)
                d_1266_v52_: str
                d_1266_v52_ = _dafny.CodePoint('l')
                d_1267_v53_: D4
                d_1267_v53_ = D4_DC13(d_1266_v52_)
                d_1268_v54_: D4
                d_1268_v54_ = D4_DC15(d_1267_v53_)
                d_1269_v55_: D4
                d_1269_v55_ = D4_DC15(d_1267_v53_)
                d_1270_v56_: _dafny.Map
                d_1270_v56_ = _dafny.Map({d_1254_i8_: d_1269_v55_})
                d_1270_v56_ = d_1270_v56_
                d_1271_v57_: _dafny.Array
                nw242_ = _dafny.Array(False, 26)
                d_1271_v57_ = nw242_
                d_1272_v58_: _dafny.Set
                d_1272_v58_ = _dafny.Set({d_1271_v57_, d_1271_v57_, d_1271_v57_})
                (globalState).f22 = (d_1271_v57_) not in (d_1272_v58_)
                d_1258_v44_ = d_1258_v44_
            elif True:
                d_1273_v59_: _dafny.Set
                d_1273_v59_ = _dafny.Set({d_1254_i8_})
                d_1274_v60_: D8
                d_1274_v60_ = D8_DC23((d_1273_v59_).ispropersubset(d_1273_v59_))
                d_1274_v60_ = d_1274_v60_
                (globalState).f13 = p1
                d_1275_v61_: C2
                nw243_ = C2()
                nw243_.ctor__(d_1258_v44_)
                d_1275_v61_ = nw243_
                (globalState).f13 = p1
                index259_ = default__.safeIndex(233, (d_1265_v51_).length(0))
                (d_1265_v51_)[index259_] = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_1276_i9_ in range(default__.abs(728))]))
        d_1277_v62_: _dafny.Array
        nw244_ = _dafny.Array(_dafny.Seq({}), 26)
        d_1277_v62_ = nw244_
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_1277_v62_).length(0)):
            d_1278_i10_: int = guard_loop_3_
            if (True) and (((0) <= (d_1278_i10_)) and ((d_1278_i10_) < ((d_1277_v62_).length(0)))):
                (d_1277_v62_)[(d_1278_i10_)] = (_dafny.SeqWithoutIsStrInference([D4_DC13(_dafny.CodePoint('n')) for d_1279_i11_ in range(default__.abs(704))])) + (_dafny.SeqWithoutIsStrInference([D4_DC13(_dafny.CodePoint('q')) for d_1280_i12_ in range(default__.abs(83))]))
        if (not (p3) or (p1) if not((p0) != (p2)) else p3):
            (globalState).f10 = p2
            d_1281_v63_: C0
            nw245_ = C0()
            nw245_.ctor__(p1)
            d_1281_v63_ = nw245_
            d_1282_v64_: _dafny.Array
            def lambda60_(d_1283_i13_):
                return _dafny.SeqWithoutIsStrInference([False])

            init31_ = lambda60_
            nw246_ = _dafny.Array(None, 25)
            for i0_31_ in range(nw246_.length(0)):
                nw246_[i0_31_] = init31_(i0_31_)
            d_1282_v64_ = nw246_
            d_1284_v65_: _dafny.Seq
            d_1284_v65_ = _dafny.SeqWithoutIsStrInference([d_1282_v64_])
            d_1285_v66_: str
            d_1285_v66_ = _dafny.CodePoint('a')
            d_1286_v67_: _dafny.Map
            d_1286_v67_ = _dafny.Map({(d_1284_v65_)[default__.safeIndex(p2, len(d_1284_v65_))]: d_1285_v66_})
            d_1286_v67_ = (d_1286_v67_).set(d_1282_v64_, d_1285_v66_)
            (globalState).f7 = p3
            d_1287_v68_: _dafny.MultiSet
            d_1287_v68_ = _dafny.MultiSet([p3, True, p3, p1, p3])
            d_1288_v69_: _dafny.Seq
            d_1288_v69_ = _dafny.SeqWithoutIsStrInference([(d_1287_v68_ if (d_1281_v63_).f32 else d_1287_v68_)])
            d_1288_v69_ = d_1288_v69_
        elif True:
            d_1289_v70_: _dafny.Map
            d_1289_v70_ = _dafny.Map({p2: p1})
            d_1290_v71_: _dafny.MultiSet
            d_1290_v71_ = _dafny.MultiSet([p3, p1, p1, ((d_1289_v70_)[p2] if (p2) in (d_1289_v70_) else p3), p3])
            d_1291_v72_: _dafny.Set
            d_1291_v72_ = _dafny.Set({d_1290_v71_, d_1290_v71_, _dafny.MultiSet([not(p3), p1])})
            d_1291_v72_ = d_1291_v72_
            pat_let_tv37_ = d_1290_v71_
            pat_let_tv38_ = p1
            pat_let_tv39_ = p2
            def iife114_(_pat_let40_0):
                def iife115_(d_1292_dt__update__tmp_h1_):
                    def iife116_(_pat_let41_0):
                        def iife117_(d_1293_dt__update_hcf21_h0_):
                            return D3_DC10(d_1293_dt__update_hcf21_h0_)
                        return iife117_(_pat_let41_0)
                    return iife116_((pat_let_tv37_).set(pat_let_tv38_, default__.abs((0) - (pat_let_tv39_))))
                return iife115_(_pat_let40_0)
            source20_ = iife114_(D3_DC10(default__.fm24(p2, p1, globalState)))
            if source20_.is_DC11:
                d_1294___mcc_h0_ = source20_.cf22
                d_1295_cf22_ = d_1294___mcc_h0_
                d_1296_v73_: _dafny.Array
                def lambda61_(d_1297_p1_):
                    def lambda62_(d_1298_i14_):
                        return d_1297_p1_

                    return lambda62_

                init32_ = lambda61_(p1)
                nw247_ = _dafny.Array(None, 10)
                for i0_32_ in range(nw247_.length(0)):
                    nw247_[i0_32_] = init32_(i0_32_)
                d_1296_v73_ = nw247_
                d_1299_v74_: _dafny.Set
                d_1299_v74_ = _dafny.Set({d_1296_v73_, d_1296_v73_})
                d_1300_v75_: _dafny.Seq
                d_1300_v75_ = _dafny.SeqWithoutIsStrInference([d_1299_v74_, d_1299_v74_, d_1299_v74_])
                (globalState).f10 = len((d_1300_v75_)[default__.safeIndex(p2, len(d_1300_v75_))])
                d_1301_v76_: C0
                nw248_ = C0()
                nw248_.ctor__(p3)
                d_1301_v76_ = nw248_
                d_1301_v76_ = d_1301_v76_
                (globalState).f19 = _dafny.CodePoint('f')
                d_1302_v77_: C1
                nw249_ = C1()
                nw249_.ctor__()
                d_1302_v77_ = nw249_
                d_1303_v78_: _dafny.Set
                d_1303_v78_ = _dafny.Set({d_1295_cf22_})
                d_1304_v79_: _dafny.Seq
                d_1304_v79_ = _dafny.SeqWithoutIsStrInference([(d_1303_v78_) == (d_1303_v78_)])
                d_1305_v80_: _dafny.Seq
                d_1305_v80_ = _dafny.SeqWithoutIsStrInference([d_1295_cf22_, d_1295_cf22_])
                d_1306_v81_: _dafny.Seq
                d_1306_v81_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pwm"))
                rhs244_ = d_1302_v77_
                rhs245_ = (d_1305_v80_)[default__.safeIndex(p0, len(d_1305_v80_))]
                rhs246_ = ((self).fm2(globalState)) != (len(d_1306_v81_))
                rhs247_ = (((d_1304_v79_).set(default__.safeIndex(d_1295_cf22_, len(d_1304_v79_)), (d_1301_v76_).f32)) + (d_1304_v79_) if False else d_1304_v79_)
                lhs221_ = globalState
                lhs222_ = globalState
                d_1302_v77_ = rhs244_
                lhs221_.f10 = rhs245_
                lhs222_.f22 = rhs246_
                d_1304_v79_ = rhs247_
            elif source20_.is_DC12:
                d_1307___mcc_h1_ = source20_.cf23
                d_1308_cf23_ = d_1307___mcc_h1_
                d_1309_v82_: str
                d_1309_v82_ = _dafny.CodePoint('j')
                d_1310_v83_: D8
                d_1310_v83_ = D8_DC22(652, p1, d_1309_v82_)
                d_1308_cf23_ = ((0) - ((d_1310_v83_).cf36)) != (p2)
                d_1311_v84_: _dafny.Set
                d_1311_v84_ = _dafny.Set({708})
                d_1312_v85_: _dafny.Seq
                d_1312_v85_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tmjx"))
                d_1313_v86_: _dafny.MultiSet
                d_1313_v86_ = _dafny.MultiSet([p2, len(d_1311_v84_), default__.safeModuloInt(len(default__.fm39(globalState)), p2), (0) - (len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ywd"))) + (d_1312_v85_)))])
                d_1314_v87_: _dafny.Seq
                d_1314_v87_ = _dafny.SeqWithoutIsStrInference([d_1312_v85_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vwntchfj"))])
                d_1315_v88_: _dafny.Seq
                d_1315_v88_ = _dafny.SeqWithoutIsStrInference([(self).fm5(p3, d_1314_v87_, False, globalState), d_1308_cf23_])
                d_1316_v89_: _dafny.Seq
                d_1316_v89_ = _dafny.SeqWithoutIsStrInference([len(d_1315_v88_), 184])
                d_1313_v86_ = (d_1313_v86_) | (_dafny.MultiSet(d_1316_v89_))
                (globalState).f10 = default__.safeDivisionInt((p0) - (p0), 774)
                d_1317_v90_: _dafny.Array
                nw250_ = _dafny.Array(None, 11)
                nw250_[int(0)] = _dafny.SeqWithoutIsStrInference([d_1309_v82_ for d_1318_i15_ in range(default__.abs(433))])
                nw250_[int(1)] = ((d_1312_v85_).set(default__.safeIndex(p2, len(d_1312_v85_)), d_1309_v82_)).set(default__.safeIndex(len((_dafny.SeqWithoutIsStrInference([d_1309_v82_ for d_1319_i16_ in range(default__.abs(465))])).set(default__.safeIndex(len(d_1312_v85_), len(_dafny.SeqWithoutIsStrInference([d_1309_v82_ for d_1320_i16_ in range(default__.abs(465))]))), d_1309_v82_)), len((d_1312_v85_).set(default__.safeIndex(p2, len(d_1312_v85_)), d_1309_v82_))), _dafny.CodePoint('h'))
                nw250_[int(2)] = d_1312_v85_
                nw250_[int(3)] = _dafny.SeqWithoutIsStrInference([d_1309_v82_ for d_1321_i17_ in range(default__.abs(14))])
                nw250_[int(4)] = d_1312_v85_
                nw250_[int(5)] = d_1312_v85_
                nw250_[int(6)] = default__.fm10(default__.fm33(p3, p3, True, True, globalState), d_1309_v82_, p3, globalState)
                nw250_[int(7)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pyctsgo"))
                nw250_[int(8)] = (d_1312_v85_).set(default__.safeIndex(len(d_1312_v85_), len(d_1312_v85_)), d_1309_v82_)
                nw250_[int(9)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qdgk"))
                nw250_[int(10)] = d_1312_v85_
                d_1317_v90_ = nw250_
                d_1322_v91_: _dafny.Map
                d_1322_v91_ = _dafny.Map({d_1317_v90_: d_1309_v82_})
                d_1322_v91_ = (d_1322_v91_).set(d_1317_v90_, _dafny.CodePoint('u'))
            elif True:
                d_1323___mcc_h2_ = source20_.cf21
                d_1324_cf21_ = d_1323___mcc_h2_
                d_1325_v92_: _dafny.Array
                nw251_ = _dafny.Array(_dafny.Map({}), 10)
                d_1325_v92_ = nw251_
                d_1326_v93_: _dafny.Map
                d_1326_v93_ = _dafny.Map({p1: p1})
                d_1327_v94_: _dafny.Map
                d_1327_v94_ = (d_1326_v93_).set(p1, p1)
                index260_ = default__.safeIndex(494, (d_1325_v92_).length(0))
                (d_1325_v92_)[index260_] = d_1327_v94_
                index261_ = default__.safeIndex(494, (d_1325_v92_).length(0))
                (d_1325_v92_)[index261_] = d_1326_v93_
                d_1328_v95_: _dafny.Seq
                d_1328_v95_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tdhv"))
                d_1329_v96_: C2
                nw252_ = C2()
                nw252_.ctor__(d_1328_v95_)
                d_1329_v96_ = nw252_
                d_1330_v97_: D8
                d_1330_v97_ = D8_DC21(_dafny.SeqWithoutIsStrInference([False, p1, not(p3)]))
                d_1330_v97_ = d_1330_v97_
                d_1331_v98_: C1
                nw253_ = C1()
                nw253_.ctor__()
                d_1331_v98_ = nw253_
            if p3:
                d_1332_v99_: _dafny.Seq
                d_1332_v99_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qvt"))
                d_1333_v100_: C2
                nw254_ = C2()
                nw254_.ctor__((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rlamqeya"))) + (d_1332_v99_))
                d_1333_v100_ = nw254_
                d_1334_v101_: _dafny.Array
                nw255_ = _dafny.Array(_dafny.Array(None, 0), 11)
                d_1334_v101_ = nw255_
                d_1335_v102_: _dafny.Seq
                d_1335_v102_ = _dafny.SeqWithoutIsStrInference([p1, p3])
                d_1336_v103_: _dafny.Seq
                d_1336_v103_ = _dafny.SeqWithoutIsStrInference([d_1333_v100_.f33])
                d_1337_v104_: _dafny.Seq
                d_1337_v104_ = _dafny.SeqWithoutIsStrInference([p0, len(d_1336_v103_)])
                d_1338_v105_: _dafny.MultiSet
                d_1338_v105_ = _dafny.MultiSet([p2])
                d_1339_v106_: _dafny.Map
                d_1339_v106_ = _dafny.Map({(d_1335_v102_)[default__.safeIndex(p2, len(d_1335_v102_))]: (d_1337_v104_)[default__.safeIndex((d_1338_v105_).cardinality, len(d_1337_v104_))]})
                d_1340_v107_: _dafny.Map
                d_1340_v107_ = _dafny.Map({(d_1339_v106_) == (d_1339_v106_): d_1334_v101_})
                d_1334_v101_ = ((d_1340_v107_)[(d_1335_v102_)[default__.safeIndex(p2, len(d_1335_v102_))]] if ((d_1335_v102_)[default__.safeIndex(p2, len(d_1335_v102_))]) in (d_1340_v107_) else d_1334_v101_)
                d_1341_v108_: _dafny.Set
                d_1341_v108_ = _dafny.Set({p1})
                d_1342_v109_: D8
                d_1342_v109_ = D8_DC23(p1)
                d_1343_v110_: D8
                d_1343_v110_ = D8_DC24(d_1342_v109_)
                d_1344_v111_: str
                d_1344_v111_ = _dafny.CodePoint('r')
                rhs248_ = (d_1341_v108_) - (d_1341_v108_)
                rhs249_ = d_1344_v111_
                rhs250_ = p3
                rhs251_ = d_1343_v110_
                rhs252_ = not(not(p1))
                lhs223_ = globalState
                lhs224_ = globalState
                lhs225_ = globalState
                d_1341_v108_ = rhs248_
                lhs223_.f19 = rhs249_
                lhs224_.f13 = rhs250_
                d_1343_v110_ = rhs251_
                lhs225_.f7 = rhs252_
                (globalState).f19 = d_1344_v111_
                d_1345_v112_: _dafny.Map
                d_1345_v112_ = _dafny.Map({d_1333_v100_.f33: p2})
                (globalState).f10 = (0) - ((len(d_1345_v112_)) - (p0))
            elif True:
                d_1346_v113_: _dafny.Array
                nw256_ = _dafny.Array(int(0), 21)
                d_1346_v113_ = nw256_
                d_1347_v114_: _dafny.Map
                d_1347_v114_ = _dafny.Map({p1: p1})
                index262_ = default__.safeIndex(332, (d_1346_v113_).length(0))
                (d_1346_v113_)[index262_] = (len(d_1347_v114_)) * (p2)
                d_1348_v115_: _dafny.Map
                d_1348_v115_ = d_1347_v114_
                d_1349_v116_: _dafny.Seq
                d_1349_v116_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "esrblhn"))
                index263_ = default__.safeIndex(626, (d_1346_v113_).length(0))
                (d_1346_v113_)[index263_] = default__.safeModuloInt(p0, len(d_1349_v116_))
                index264_ = default__.safeIndex(739, (d_1346_v113_).length(0))
                (d_1346_v113_)[index264_] = p2
                index265_ = default__.safeIndex(332, (d_1346_v113_).length(0))
                index266_ = default__.safeIndex(626, (d_1346_v113_).length(0))
                index267_ = default__.safeIndex(739, (d_1346_v113_).length(0))
                rhs253_ = ((p0) * (p2)) * (373)
                rhs254_ = (self).fm5(p3, _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ochwdrjm")), d_1349_v116_]), p1, globalState)
                rhs255_ = d_1348_v115_
                rhs256_ = (0) - (((self).fm6(globalState)) * (len((d_1289_v70_) | (d_1289_v70_))))
                rhs257_ = len(d_1349_v116_)
                lhs226_ = d_1346_v113_
                lhs227_ = default__.safeIndex(332, (d_1346_v113_).length(0))
                lhs228_ = globalState
                lhs229_ = d_1346_v113_
                lhs230_ = default__.safeIndex(626, (d_1346_v113_).length(0))
                lhs231_ = d_1346_v113_
                lhs232_ = default__.safeIndex(739, (d_1346_v113_).length(0))
                lhs226_[lhs227_] = rhs253_
                lhs228_.f22 = rhs254_
                d_1348_v115_ = rhs255_
                lhs229_[lhs230_] = rhs256_
                lhs231_[lhs232_] = rhs257_
                d_1350_v117_: _dafny.Array
                def lambda63_(d_1351_p1_):
                    def lambda64_(d_1352_i18_):
                        return d_1351_p1_

                    return lambda64_

                init33_ = lambda63_(p1)
                nw257_ = _dafny.Array(None, 8)
                for i0_33_ in range(nw257_.length(0)):
                    nw257_[i0_33_] = init33_(i0_33_)
                d_1350_v117_ = nw257_
                d_1353_v118_: _dafny.Array
                nw258_ = _dafny.Array(None, 6)
                nw258_[int(0)] = d_1350_v117_
                nw258_[int(1)] = d_1350_v117_
                nw258_[int(2)] = d_1350_v117_
                nw258_[int(3)] = d_1350_v117_
                nw258_[int(4)] = d_1350_v117_
                nw258_[int(5)] = (d_1350_v117_ if p1 else d_1350_v117_)
                d_1353_v118_ = nw258_
                index268_ = default__.safeIndex(189, (d_1353_v118_).length(0))
                (d_1353_v118_)[index268_] = d_1350_v117_
                d_1354_v119_: C2
                nw259_ = C2()
                nw259_.ctor__(d_1349_v116_)
                d_1354_v119_ = nw259_
                d_1355_v120_: _dafny.Set
                d_1355_v120_ = _dafny.Set({p1})
                d_1356_v121_: _dafny.Map
                d_1356_v121_ = _dafny.Map({p3: d_1354_v119_})
                index269_ = default__.safeIndex(189, (d_1353_v118_).length(0))
                rhs258_ = (d_1355_v120_) != ((d_1355_v120_ if p3 else d_1355_v120_))
                rhs259_ = d_1350_v117_
                rhs260_ = ((d_1356_v121_)[p3] if (p3) in (d_1356_v121_) else d_1354_v119_)
                rhs261_ = p3
                lhs233_ = globalState
                lhs234_ = d_1353_v118_
                lhs235_ = default__.safeIndex(189, (d_1353_v118_).length(0))
                lhs236_ = globalState
                lhs233_.f13 = rhs258_
                lhs234_[lhs235_] = rhs259_
                d_1354_v119_ = rhs260_
                lhs236_.f22 = rhs261_
                (globalState).f10 = default__.safeModuloInt((p2) * (p0), p2)
                arr0_ = (d_1353_v118_)[default__.safeIndex(189, (d_1353_v118_).length(0))]
                index270_ = default__.safeIndex(716, ((d_1353_v118_)[default__.safeIndex(189, (d_1353_v118_).length(0))]).length(0))
                arr0_[index270_] = p3
                arr1_ = (d_1353_v118_)[default__.safeIndex(189, (d_1353_v118_).length(0))]
                index271_ = default__.safeIndex(716, ((d_1353_v118_)[default__.safeIndex(189, (d_1353_v118_).length(0))]).length(0))
                arr1_[index271_] = p3
                (globalState).f10 = p0
            rhs262_ = p1
            rhs263_ = p1
            lhs237_ = globalState
            lhs238_ = globalState
            lhs237_.f13 = rhs262_
            lhs238_.f13 = rhs263_
            d_1357_v123_: _dafny.Seq
            d_1357_v123_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "idmp"))
            d_1358_v124_: _dafny.MultiSet
            d_1358_v124_ = _dafny.MultiSet([d_1357_v123_])
            def iife118_():
                coll34_ = _dafny.Map()
                compr_34_: _dafny.Seq
                for compr_34_ in (d_1358_v124_).Elements:
                    d_1359_v122_: _dafny.Seq = compr_34_
                    if (d_1359_v122_) in (d_1358_v124_):
                        coll34_[d_1359_v122_] = p1
                return _dafny.Map(coll34_)
            (globalState).f10 = len(iife118_()
            )
        d_1360_v125_: _dafny.Array
        nw260_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 6)
        d_1360_v125_ = nw260_
        guard_loop_4_: int
        for guard_loop_4_ in _dafny.IntegerRange(0, (d_1360_v125_).length(0)):
            d_1361_i19_: int = guard_loop_4_
            if (True) and (((0) <= (d_1361_i19_)) and ((d_1361_i19_) < ((d_1360_v125_).length(0)))):
                (d_1360_v125_)[(d_1361_i19_)] = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "roywi"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fllqnoebg")))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_1362_i20_ in range(default__.abs(919))]))

    def m1(self, p0, p1, p2, globalState):
        r0: D0 = D0.default()()
        r1: bool = False
        r1 = p2
        d_1363_i0_: int
        d_1363_i0_ = 0
        with _dafny.label("5"):
            while p2:
                with _dafny.c_label("5"):
                    if (d_1363_i0_) >= (100):
                        raise _dafny.Break("5")
                    d_1363_i0_ = (d_1363_i0_) + (1)
                    d_1364_v0_: _dafny.Seq
                    d_1364_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oekpnmql"))
                    d_1365_v1_: _dafny.Map
                    d_1365_v1_ = _dafny.Map({p2: 776})
                    d_1366_v2_: _dafny.Map
                    d_1366_v2_ = _dafny.Map({p1: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i"))})
                    d_1367_v3_: _dafny.Set
                    d_1367_v3_ = _dafny.Set({p2, not(p2), p2})
                    d_1368_v4_: str
                    d_1368_v4_ = _dafny.CodePoint('k')
                    rhs264_ = ((d_1366_v2_)[-300] if (-300) in (d_1366_v2_) else ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "phjwav"))) + (d_1364_v0_)).set(default__.safeIndex(len(d_1367_v3_), len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "phjwav"))) + (d_1364_v0_))), d_1368_v4_))
                    rhs265_ = (d_1365_v1_ if p2 else d_1365_v1_)
                    d_1364_v0_ = rhs264_
                    d_1365_v1_ = rhs265_
                    d_1369_v5_: _dafny.Set
                    d_1369_v5_ = _dafny.Set({d_1364_v0_, d_1364_v0_, d_1364_v0_})
                    d_1370_v6_: _dafny.Map
                    d_1370_v6_ = _dafny.Map({(p1) + ((0) - ((self).fm2(globalState))): d_1369_v5_})
                    d_1371_v7_: _dafny.Map
                    d_1371_v7_ = _dafny.Map({990: (self).fm2(globalState)})
                    d_1370_v6_ = (d_1370_v6_).set(len(default__.fm23(D2_DC7(d_1371_v7_), 728, p1, globalState)), (d_1369_v5_) - (d_1369_v5_))
                    (globalState).f7 = p2
                    d_1372_v8_: _dafny.Array
                    out8_: _dafny.Array
                    out8_ = (self).m2(globalState)
                    d_1372_v8_ = out8_
                    pass
            pass
        (globalState).f10 = p1
        d_1373_v9_: _dafny.MultiSet
        d_1373_v9_ = _dafny.MultiSet([p2])
        d_1374_v10_: D3
        d_1374_v10_ = D3_DC10(d_1373_v9_)
        if ((d_1374_v10_).cf21).isdisjoint((d_1373_v9_) | (d_1373_v9_)):
            d_1375_v11_: _dafny.Array
            nw261_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 16)
            d_1375_v11_ = nw261_
            index272_ = default__.safeIndex(36, (d_1375_v11_).length(0))
            (d_1375_v11_)[index272_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xbbmhsxm"))
            d_1376_v12_: _dafny.Seq
            d_1376_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oobix"))
            index273_ = default__.safeIndex(36, (d_1375_v11_).length(0))
            (d_1375_v11_)[index273_] = (d_1376_v12_) + (d_1376_v12_)
            d_1377_v14_: D0
            d_1377_v14_ = D0_DC2(p1, not(p2))
            d_1378_v15_: _dafny.Array
            nw262_ = _dafny.Array(None, 19)
            nw262_[int(0)] = True
            nw262_[int(1)] = p2
            nw262_[int(2)] = p2
            nw262_[int(3)] = p2
            nw262_[int(4)] = False
            nw262_[int(5)] = p2
            nw262_[int(6)] = p2
            nw262_[int(7)] = p2
            nw262_[int(8)] = False
            nw262_[int(9)] = p2
            nw262_[int(10)] = p2
            nw262_[int(11)] = p2
            nw262_[int(12)] = p2
            nw262_[int(13)] = p2
            nw262_[int(14)] = (d_1377_v14_).cf6
            nw262_[int(15)] = p2
            nw262_[int(16)] = p2
            nw262_[int(17)] = p2
            nw262_[int(18)] = p2
            d_1378_v15_ = nw262_
            d_1379_v16_: D4
            d_1379_v16_ = D4_DC14(p2)
            d_1380_v17_: D4
            d_1380_v17_ = D4_DC15(d_1379_v16_)
            d_1381_v18_: D7
            d_1381_v18_ = D7_DC20(p0, d_1378_v15_, d_1380_v17_, p2, p1)
            d_1382_v19_: _dafny.Seq
            d_1382_v19_ = _dafny.SeqWithoutIsStrInference([p2])
            d_1383_v20_: _dafny.Map
            d_1383_v20_ = _dafny.Map({p1: -839})
            d_1384_v22_: _dafny.Map
            d_1384_v22_ = _dafny.Map({p1: p2})
            d_1385_v23_: _dafny.Map
            def iife119_():
                coll35_ = _dafny.Map()
                compr_35_: int
                for compr_35_ in (d_1384_v22_).keys.Elements:
                    d_1387_v21_: int = compr_35_
                    if (d_1387_v21_) in (d_1384_v22_):
                        coll35_[default__.safeDivisionInt(d_1387_v21_, -486)] = p2
                return _dafny.Map(coll35_)
            d_1385_v23_ = _dafny.Map({(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([len(d_1376_v12_) for d_1386_i1_ in range(default__.abs(91))])), len(iife119_()
            )])).cardinality: (self).fm5(p2, _dafny.SeqWithoutIsStrInference([d_1376_v12_ for d_1388_i2_ in range(default__.abs(884))]), p2, globalState)})
            d_1389_v24_: _dafny.Array
            nw263_ = _dafny.Array(None, 22)
            nw263_[int(0)] = (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "axtltt"))))
            nw263_[int(1)] = p1
            nw263_[int(2)] = p1
            nw263_[int(3)] = p1
            nw263_[int(4)] = p1
            nw263_[int(5)] = p1
            nw263_[int(6)] = p1
            nw263_[int(7)] = p1
            nw263_[int(8)] = p1
            nw263_[int(9)] = p1
            def iife120_():
                coll36_ = _dafny.Map()
                compr_36_: int
                for compr_36_ in ((d_1381_v18_).cf30).Elements:
                    d_1390_v13_: int = compr_36_
                    if (d_1390_v13_) in ((d_1381_v18_).cf30):
                        coll36_[(d_1390_v13_) - ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([p1]))).cardinality)] = p1
                return _dafny.Map(coll36_)
            nw263_[int(10)] = len(iife120_()
            )
            nw263_[int(11)] = p1
            nw263_[int(12)] = (0) - (p1)
            nw263_[int(13)] = p1
            nw263_[int(14)] = len(d_1382_v19_)
            nw263_[int(15)] = len(d_1383_v20_)
            nw263_[int(16)] = p1
            nw263_[int(17)] = p1
            nw263_[int(18)] = 218
            nw263_[int(19)] = len(d_1385_v23_)
            nw263_[int(20)] = len(_dafny.SeqWithoutIsStrInference([p2, not(False)]))
            nw263_[int(21)] = p1
            d_1389_v24_ = nw263_
            d_1391_v25_: _dafny.MultiSet
            d_1391_v25_ = _dafny.MultiSet([d_1389_v24_])
            d_1392_v26_: _dafny.MultiSet
            d_1392_v26_ = _dafny.MultiSet([default__.safeModuloInt(p1, (d_1391_v25_).cardinality), p1, p1, ((p0)[default__.safeIndex(p1, len(p0))]) - (p1)])
            d_1392_v26_ = (d_1392_v26_ if p2 else d_1392_v26_)
            d_1393_v27_: _dafny.Set
            d_1393_v27_ = _dafny.Set({p1})
            (globalState).f24 = d_1393_v27_
            (globalState).f7 = (p2) == (p2)
            d_1382_v19_ = _dafny.SeqWithoutIsStrInference([p2, p2])
        elif True:
            d_1394_v28_: _dafny.Array
            nw264_ = _dafny.Array(_dafny.MultiSet({}), 5)
            d_1394_v28_ = nw264_
            d_1395_v29_: _dafny.Array
            nw265_ = _dafny.Array(False, 6)
            d_1395_v29_ = nw265_
            index274_ = default__.safeIndex(312, (d_1394_v28_).length(0))
            (d_1394_v28_)[index274_] = _dafny.MultiSet([d_1395_v29_])
            d_1396_v30_: _dafny.MultiSet
            d_1396_v30_ = _dafny.MultiSet([d_1395_v29_, d_1395_v29_])
            index275_ = default__.safeIndex(312, (d_1394_v28_).length(0))
            (d_1394_v28_)[index275_] = d_1396_v30_
            (globalState).f19 = _dafny.CodePoint('b')
            d_1395_v29_ = d_1395_v29_
            d_1397_v31_: str
            d_1397_v31_ = _dafny.CodePoint('d')
            d_1398_v32_: _dafny.Seq
            d_1398_v32_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kxbccg"))
            d_1399_v33_: _dafny.Map
            d_1399_v33_ = _dafny.Map({d_1397_v31_: d_1398_v32_})
            (globalState).f10 = default__.safeModuloInt(len(d_1399_v33_), p1)
            d_1400_v34_: C1
            nw266_ = C1()
            nw266_.ctor__()
            d_1400_v34_ = nw266_
        if p2:
            d_1401_v35_: _dafny.Seq
            d_1401_v35_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g"))
            d_1402_v36_: C2
            nw267_ = C2()
            nw267_.ctor__((d_1401_v35_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "enwdan"))))
            d_1402_v36_ = nw267_
            (globalState).f10 = p1
            d_1403_v37_: D0
            d_1403_v37_ = D0_DC2(default__.fm0(p1, globalState), False)
            pat_let_tv40_ = p1
            def iife121_(_pat_let42_0):
                def iife122_(d_1404_dt__update__tmp_h0_):
                    def iife123_(_pat_let43_0):
                        def iife124_(d_1405_dt__update_hcf5_h0_):
                            return D0_DC2(d_1405_dt__update_hcf5_h0_, (d_1404_dt__update__tmp_h0_).cf6)
                        return iife124_(_pat_let43_0)
                    return iife123_(pat_let_tv40_)
                return iife122_(_pat_let42_0)
            if (iife121_(d_1403_v37_)).cf6:
                (globalState).f13 = p2
                d_1406_v40_: _dafny.Map
                d_1406_v40_ = _dafny.Map({p2: p2})
                d_1407_v41_: _dafny.Seq
                d_1407_v41_ = _dafny.SeqWithoutIsStrInference([((d_1406_v40_)[p2] if (p2) in (d_1406_v40_) else p2)])
                d_1408_v42_: _dafny.Set
                d_1408_v42_ = _dafny.Set({d_1407_v41_})
                def iife125_():
                    def iife127_():
                        coll39_ = _dafny.Map()
                        compr_39_: _dafny.Seq
                        for compr_39_ in (d_1408_v42_).Elements:
                            d_1409_v39_: _dafny.Seq = compr_39_
                            if (d_1409_v39_) in (d_1408_v42_):
                                coll39_[d_1409_v39_] = p2
                        return _dafny.Map(coll39_)
                    coll37_ = _dafny.Map()
                    def iife126_():
                        coll38_ = _dafny.Map()
                        compr_38_: _dafny.Seq
                        for compr_38_ in (d_1408_v42_).Elements:
                            d_1409_v39_: _dafny.Seq = compr_38_
                            if (d_1409_v39_) in (d_1408_v42_):
                                coll38_[d_1409_v39_] = p2
                        return _dafny.Map(coll38_)
                    compr_37_: _dafny.Seq
                    for compr_37_ in (iife126_()
                    ).keys.Elements:
                        d_1410_v38_: _dafny.Seq = compr_37_
                        if (d_1410_v38_) in (iife127_()
                        ):
                            coll37_[d_1410_v38_] = (d_1373_v9_).cardinality
                    return _dafny.Map(coll37_)
                rhs266_ = ((self).fm2(globalState) if not(p2) else len(iife125_()
                ))
                rhs267_ = d_1373_v9_
                lhs239_ = globalState
                lhs239_.f10 = rhs266_
                d_1373_v9_ = rhs267_
                d_1411_v43_: str
                d_1411_v43_ = _dafny.CodePoint('c')
                d_1412_v44_: _dafny.Map
                d_1412_v44_ = _dafny.Map({p1: p2})
                d_1413_v45_: _dafny.Map
                d_1413_v45_ = _dafny.Map({d_1411_v43_: len(d_1412_v44_)})
                d_1414_v46_: _dafny.Map
                d_1414_v46_ = _dafny.Map({p1: (p0).set(default__.safeIndex(len(d_1413_v45_), len(p0)), p1)})
                (d_1402_v36_).m0(len(((d_1414_v46_)[p1] if (p1) in (d_1414_v46_) else p0)), p2, p1, not(p2), globalState)
                (globalState).f13 = ((0) - (p1)) < (p1)
                d_1415_v47_: _dafny.Map
                d_1415_v47_ = _dafny.Map({False: p1})
                d_1415_v47_ = (d_1415_v47_).set((p1) <= (p1), p1)
            elif True:
                (globalState).f13 = p2
                (globalState).f10 = (p0)[default__.safeIndex(p1, len(p0))]
                d_1416_v48_: _dafny.Array
                nw268_ = _dafny.Array(False, 6)
                d_1416_v48_ = nw268_
                (globalState).f10 = (_dafny.MultiSet([d_1416_v48_, d_1416_v48_])).cardinality
                d_1417_v49_: str
                d_1417_v49_ = _dafny.CodePoint('h')
                (globalState).f19 = d_1417_v49_
                d_1418_v50_: _dafny.Map
                d_1418_v50_ = _dafny.Map({p2: d_1402_v36_.f33})
                d_1419_v51_: _dafny.Seq
                d_1419_v51_ = _dafny.SeqWithoutIsStrInference([d_1402_v36_.f33, (((d_1418_v50_)[p2] if (p2) in (d_1418_v50_) else d_1402_v36_.f33)).set(default__.safeIndex(137, len(((d_1418_v50_)[p2] if (p2) in (d_1418_v50_) else d_1402_v36_.f33))), d_1417_v49_), (d_1401_v35_) + (d_1401_v35_)])
                d_1419_v51_ = ((d_1419_v51_) + (d_1419_v51_)) + (d_1419_v51_)
            (d_1402_v36_).m0(len(d_1401_v35_), p2, (0) - (default__.safeModuloInt(p1, p1)), True, globalState)
            (globalState).f10 = (0) - (default__.safeDivisionInt((p1) - (p1), p1))
        elif True:
            (globalState).f13 = p2
            (globalState).f10 = ((d_1373_v9_).cardinality) + ((146) + (p1))
            d_1420_v52_: _dafny.Map
            d_1420_v52_ = _dafny.Map({default__.safeDivisionInt((self).fm6(globalState), p1): (len(_dafny.Map({0: p0}))) * (p1)})
            d_1420_v52_ = (d_1420_v52_).set(p1, 841)
            (globalState).f10 = 622
            d_1421_v53_: _dafny.Map
            d_1421_v53_ = _dafny.Map({p2: len(_dafny.SeqWithoutIsStrInference([p2]))})
            d_1422_v54_: _dafny.MultiSet
            d_1422_v54_ = _dafny.MultiSet([844])
            d_1423_v55_: _dafny.Seq
            d_1423_v55_ = _dafny.SeqWithoutIsStrInference([d_1422_v54_])
            d_1424_v56_: _dafny.Seq
            d_1424_v56_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gyebnjvkd"))
            d_1425_v57_: _dafny.Map
            d_1425_v57_ = _dafny.Map({p1: p2})
            d_1426_v58_: _dafny.Array
            nw269_ = _dafny.Array(None, 26)
            nw269_[int(0)] = len(_dafny.Map({_dafny.SeqWithoutIsStrInference([p1]): p2}))
            nw269_[int(1)] = p1
            nw269_[int(2)] = p1
            nw269_[int(3)] = p1
            nw269_[int(4)] = p1
            nw269_[int(5)] = p1
            nw269_[int(6)] = len(d_1421_v53_)
            nw269_[int(7)] = p1
            nw269_[int(8)] = len(d_1423_v55_)
            nw269_[int(9)] = p1
            nw269_[int(10)] = p1
            nw269_[int(11)] = p1
            nw269_[int(12)] = p1
            nw269_[int(13)] = p1
            nw269_[int(14)] = ((d_1422_v54_)[p1] if (p1) in (d_1422_v54_) else p1)
            nw269_[int(15)] = ((d_1422_v54_).set(p1, default__.abs(p1))).cardinality
            nw269_[int(16)] = len(d_1424_v56_)
            nw269_[int(17)] = p1
            nw269_[int(18)] = p1
            nw269_[int(19)] = len(d_1424_v56_)
            nw269_[int(20)] = p1
            nw269_[int(21)] = len(d_1425_v57_)
            nw269_[int(22)] = p1
            nw269_[int(23)] = p1
            nw269_[int(24)] = p1
            nw269_[int(25)] = p1
            d_1426_v58_ = nw269_
            d_1427_v59_: _dafny.Map
            d_1427_v59_ = _dafny.Map({d_1426_v58_: p2})
            (globalState).f10 = (447) + (len(d_1427_v59_))
        d_1428_v60_: _dafny.Map
        d_1428_v60_ = _dafny.Map({default__.safeModuloInt((0) - (p1), -889): (0) - ((p1 if p2 else p1))})
        d_1428_v60_ = (d_1428_v60_).set(p1, 565)
        d_1429_v61_: _dafny.Seq
        d_1429_v61_ = _dafny.SeqWithoutIsStrInference([p2])
        d_1430_v62_: _dafny.Map
        d_1430_v62_ = _dafny.Map({p2: p2})
        d_1431_v63_: _dafny.MultiSet
        d_1431_v63_ = _dafny.MultiSet([p1])
        d_1432_v64_: _dafny.Map
        d_1432_v64_ = _dafny.Map({d_1431_v63_: p0})
        d_1433_v65_: D1
        d_1433_v65_ = D1_DC5(p1, p2, ((d_1430_v62_)[False] if (False) in (d_1430_v62_) else not(p2)), d_1432_v64_)
        d_1434_v66_: _dafny.Array
        nw270_ = _dafny.Array(False, 17)
        d_1434_v66_ = nw270_
        d_1435_v67_: D0
        d_1435_v67_ = D0_DC1(len((d_1429_v61_).set(default__.safeIndex(p1, len(d_1429_v61_)), p2)), p1, ((0) - (p1) if (d_1433_v65_).cf10 else len(d_1429_v61_)), d_1434_v66_)
        r0 = d_1435_v67_
        r1 = (d_1431_v63_).issubset((d_1431_v63_ if True else d_1431_v63_))
        return r0, r1

    def m2(self, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        d_1436_v0_: _dafny.Array
        nw271_ = _dafny.Array(False, 8)
        d_1436_v0_ = nw271_
        guard_loop_5_: int
        for guard_loop_5_ in _dafny.IntegerRange(0, (d_1436_v0_).length(0)):
            d_1437_i0_: int = guard_loop_5_
            if (True) and (((0) <= (d_1437_i0_)) and ((d_1437_i0_) < ((d_1436_v0_).length(0)))):
                (d_1436_v0_)[(d_1437_i0_)] = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gkgylxbum"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o")))) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "txvt")))
        d_1438_v1_: int
        d_1438_v1_ = 109
        d_1439_v2_: bool
        d_1439_v2_ = False
        d_1440_v3_: D0
        d_1440_v3_ = D0_DC2(d_1438_v1_, d_1439_v2_)
        d_1441_v4_: _dafny.Seq
        d_1441_v4_ = _dafny.SeqWithoutIsStrInference([d_1440_v3_, d_1440_v3_, d_1440_v3_])
        d_1442_v6_: _dafny.MultiSet
        def iife128_():
            coll40_ = _dafny.Map()
            compr_40_: int
            for compr_40_ in _dafny.IntegerRange(-15, -440):
                d_1443_v5_: int = compr_40_
                if ((-15) <= (d_1443_v5_)) and ((d_1443_v5_) < (-440)):
                    coll40_[(d_1443_v5_) - (len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lem"))): d_1438_v1_})))] = 518
            return _dafny.Map(coll40_)
        d_1442_v6_ = _dafny.MultiSet([iife128_()
        ])
        d_1444_v8_: D0
        def iife129_():
            coll41_ = _dafny.Set()
            compr_41_: _dafny.Map
            for compr_41_ in (d_1442_v6_).Elements:
                d_1445_v7_: _dafny.Map = compr_41_
                if (d_1445_v7_) in (d_1442_v6_):
                    coll41_ = coll41_.union(_dafny.Set([d_1445_v7_]))
            return _dafny.Set(coll41_)
        d_1444_v8_ = D0_DC3((d_1441_v4_)[default__.safeIndex(len(iife129_()
), len(d_1441_v4_))])
        d_1446_v9_: _dafny.Map
        d_1446_v9_ = _dafny.Map({d_1438_v1_: d_1444_v8_})
        d_1447_v10_: _dafny.Seq
        d_1447_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w"))
        d_1448_v11_: D3
        d_1448_v11_ = D3_DC11(d_1438_v1_)
        d_1449_v12_: _dafny.MultiSet
        d_1449_v12_ = _dafny.MultiSet([d_1438_v1_])
        d_1450_v13_: _dafny.Map
        d_1450_v13_ = _dafny.Map({d_1449_v12_: _dafny.SeqWithoutIsStrInference([-553 for d_1451_i1_ in range(default__.abs(808))])})
        d_1452_v14_: D1
        d_1452_v14_ = D1_DC5(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "csvrecknv"))), d_1439_v2_, d_1439_v2_, d_1450_v13_)
        d_1453_v15_: D1
        d_1453_v15_ = D1_DC5(d_1438_v1_, False, d_1439_v2_, (d_1452_v14_).cf12)
        d_1454_v16_: _dafny.Array
        def lambda65_(d_1455_v10_):
            def lambda66_(d_1456_i2_):
                return default__.safeDivisionInt(d_1456_i2_, len(d_1455_v10_))

            return lambda66_

        init34_ = lambda65_(d_1447_v10_)
        nw272_ = _dafny.Array(None, 3)
        for i0_34_ in range(nw272_.length(0)):
            nw272_[i0_34_] = init34_(i0_34_)
        d_1454_v16_ = nw272_
        d_1457_v17_: _dafny.Map
        d_1457_v17_ = _dafny.Map({d_1453_v15_: d_1454_v16_})
        d_1458_v18_: _dafny.Array
        nw273_ = _dafny.Array(None, 12)
        nw273_[int(0)] = -908
        nw273_[int(1)] = d_1438_v1_
        nw273_[int(2)] = len(d_1446_v9_)
        nw273_[int(3)] = (len(d_1447_v10_)) + ((0) - (d_1438_v1_))
        nw273_[int(4)] = d_1438_v1_
        nw273_[int(5)] = (d_1448_v11_).cf22
        nw273_[int(6)] = (0) - (d_1438_v1_)
        nw273_[int(7)] = d_1438_v1_
        nw273_[int(8)] = d_1438_v1_
        nw273_[int(9)] = (0) - (len((d_1457_v17_) | (_dafny.Map({default__.fm25(d_1439_v2_, d_1439_v2_, d_1438_v1_, d_1439_v2_, globalState): d_1454_v16_}))))
        nw273_[int(10)] = d_1438_v1_
        nw273_[int(11)] = d_1438_v1_
        d_1458_v18_ = nw273_
        index276_ = default__.safeIndex(476, (d_1458_v18_).length(0))
        (d_1458_v18_)[index276_] = d_1438_v1_
        index277_ = default__.safeIndex(476, (d_1458_v18_).length(0))
        (d_1458_v18_)[index277_] = d_1438_v1_
        d_1459_v19_: _dafny.Map
        d_1459_v19_ = _dafny.Map({((d_1458_v18_)[default__.safeIndex(476, (d_1458_v18_).length(0))]) <= (d_1438_v1_): _dafny.Set({d_1454_v16_, d_1454_v16_, d_1458_v18_})})
        d_1460_v20_: _dafny.Set
        d_1460_v20_ = _dafny.Set({d_1458_v18_})
        d_1459_v19_ = (d_1459_v19_).set(d_1439_v2_, d_1460_v20_)
        d_1461_v21_: D2
        d_1461_v21_ = D2_DC8(d_1439_v2_, d_1439_v2_)
        pat_let_tv41_ = d_1438_v1_
        pat_let_tv42_ = d_1438_v1_
        pat_let_tv43_ = d_1439_v2_
        def lambda67_(source21_):
            if source21_.is_DC8:
                d_1462___mcc_h0_ = source21_.cf18
                d_1463___mcc_h1_ = source21_.cf19
                d_1464_cf19_ = d_1463___mcc_h1_
                d_1465_cf18_ = d_1462___mcc_h0_
                return pat_let_tv41_
            elif source21_.is_DC7:
                d_1466___mcc_h2_ = source21_.cf17
                d_1467_cf17_ = d_1466___mcc_h2_
                return pat_let_tv42_
            elif True:
                d_1468___mcc_h3_ = source21_.cf20
                d_1469_cf20_ = d_1468___mcc_h3_
                return 866

        def iife130_(_pat_let44_0):
            def iife131_(d_1470_dt__update__tmp_h0_):
                def iife132_(_pat_let45_0):
                    def iife133_(d_1471_dt__update_hcf18_h0_):
                        return D2_DC8(d_1471_dt__update_hcf18_h0_, (d_1470_dt__update__tmp_h0_).cf19)
                    return iife133_(_pat_let45_0)
                return iife132_(pat_let_tv43_)
            return iife131_(_pat_let44_0)
        d_1438_v1_ = lambda67_(iife130_(d_1461_v21_))
        index278_ = default__.safeIndex(476, (d_1458_v18_).length(0))
        (d_1458_v18_)[index278_] = (0) - ((0) - ((d_1458_v18_)[default__.safeIndex(476, (d_1458_v18_).length(0))]))
        d_1472_v22_: D1
        d_1473_v23_: int
        d_1474_v24_: _dafny.Seq
        out9_: D1
        out10_: int
        out11_: _dafny.Seq
        out9_, out10_, out11_ = (self).m3((d_1439_v2_ if d_1439_v2_ else d_1439_v2_), d_1436_v0_, globalState)
        d_1472_v22_ = out9_
        d_1473_v23_ = out10_
        d_1474_v24_ = out11_
        r0 = d_1454_v16_
        return r0

    def m3(self, p0, p1, globalState):
        r0: D1 = D1.default()()
        r1: int = int(0)
        r2: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        d_1475_v0_: _dafny.Seq
        d_1475_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ri"))
        d_1476_v1_: int
        d_1476_v1_ = 212
        d_1477_v2_: _dafny.Map
        d_1477_v2_ = _dafny.Map({d_1475_v0_: _dafny.Set({(0) - (d_1476_v1_), d_1476_v1_})})
        d_1478_v3_: _dafny.Set
        d_1478_v3_ = _dafny.Set({d_1476_v1_})
        d_1477_v2_ = (d_1477_v2_).set(d_1475_v0_, d_1478_v3_)
        (globalState).f13 = p0
        d_1479_v4_: _dafny.Seq
        d_1479_v4_ = _dafny.SeqWithoutIsStrInference([p0])
        d_1480_v5_: _dafny.Seq
        d_1480_v5_ = _dafny.SeqWithoutIsStrInference([(d_1479_v4_)[default__.safeIndex(d_1476_v1_, len(d_1479_v4_))]])
        d_1481_v6_: _dafny.Map
        d_1481_v6_ = _dafny.Map({len(d_1479_v4_): d_1476_v1_})
        d_1481_v6_ = (d_1481_v6_).set(d_1476_v1_, d_1476_v1_)
        d_1482_v7_: C0
        nw274_ = C0()
        nw274_.ctor__(p0)
        d_1482_v7_ = nw274_
        d_1482_v7_ = d_1482_v7_
        (globalState).f13 = (default__.fm11(d_1476_v1_, globalState)) not in ((d_1475_v0_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_1483_i0_ in range(default__.abs(668))])))
        d_1484_v8_: C3
        nw275_ = C3()
        nw275_.ctor__()
        d_1484_v8_ = nw275_
        r0 = default__.fm29(globalState)
        d_1485_v9_: _dafny.Map
        d_1485_v9_ = _dafny.Map({p0: p0})
        d_1486_v10_: _dafny.Set
        d_1486_v10_ = _dafny.Set({default__.fm30((d_1482_v7_).f32, globalState), d_1485_v9_})
        r1 = default__.safeModuloInt((0) - (default__.safeModuloInt(d_1476_v1_, len(d_1486_v10_))), (self).fm6(globalState))
        r2 = default__.fm35(d_1476_v1_, p0, ((d_1482_v7_).f32 if (d_1482_v7_).f32 else p0), globalState)
        return r0, r1, r2

    @property
    def f31(self):
        return self._f31

class C5(T1, T0):
    def  __init__(self):
        self.f30: int = int(0)
        self._f29: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    def ctor__(self, f29, f30):
        (self)._f29 = f29
        (self).f30 = f30

    def fm1(self, p0, p1, p2, p3, globalState):
        def iife134_():
            coll42_ = _dafny.Map()
            compr_42_: _dafny.MultiSet
            for compr_42_ in (_dafny.Map({_dafny.MultiSet([False, True, False]): self.f30})).keys.Elements:
                d_1487_v0_: _dafny.MultiSet = compr_42_
                if (d_1487_v0_) in (_dafny.Map({_dafny.MultiSet([False, True, False]): self.f30})):
                    coll42_[d_1487_v0_] = True
            return _dafny.Map(coll42_)
        return (D1_DC4(iife134_()
)).cf8

    def fm2(self, globalState):
        return self.f30

    def fm3(self, p0, p1, p2, p3, globalState):
        return (self.f30) * (len(_dafny.Map({self.f30: not(False)})))

    def fm4(self, p0, p1, globalState):
        return False

    def m0(self, p0, p1, p2, p3, globalState):
        if p1:
            d_1488_v0_: _dafny.Array
            nw276_ = _dafny.Array(_dafny.MultiSet({}), 11)
            d_1488_v0_ = nw276_
            d_1489_v1_: _dafny.Array
            def lambda68_(d_1490_p3_):
                def lambda69_(d_1491_i0_):
                    return d_1490_p3_

                return lambda69_

            init35_ = lambda68_(p3)
            nw277_ = _dafny.Array(None, 1)
            for i0_35_ in range(nw277_.length(0)):
                nw277_[i0_35_] = init35_(i0_35_)
            d_1489_v1_ = nw277_
            d_1492_v2_: _dafny.MultiSet
            d_1492_v2_ = _dafny.MultiSet([d_1489_v1_])
            index279_ = default__.safeIndex(161, (d_1488_v0_).length(0))
            (d_1488_v0_)[index279_] = d_1492_v2_
            index280_ = default__.safeIndex(161, (d_1488_v0_).length(0))
            rhs268_ = (d_1492_v2_) - (d_1492_v2_)
            rhs269_ = (p1) == (p1)
            lhs240_ = d_1488_v0_
            lhs241_ = default__.safeIndex(161, (d_1488_v0_).length(0))
            lhs242_ = globalState
            lhs240_[lhs241_] = rhs268_
            lhs242_.f22 = rhs269_
            index281_ = default__.safeIndex(417, (d_1489_v1_).length(0))
            (d_1489_v1_)[index281_] = p3
            d_1493_v3_: _dafny.Seq
            d_1493_v3_ = _dafny.SeqWithoutIsStrInference([True])
            d_1494_v4_: _dafny.Set
            d_1494_v4_ = _dafny.Set({p1, p3, (d_1493_v3_)[default__.safeIndex(self.f30, len(d_1493_v3_))]})
            index282_ = default__.safeIndex(417, (d_1489_v1_).length(0))
            (d_1489_v1_)[index282_] = ((_dafny.Set({False})) - (_dafny.Set({not(p1), p1, p1, True, p1}))).ispropersubset(d_1494_v4_)
            d_1495_v5_: _dafny.Array
            def lambda70_(d_1496_p2_, d_1497_p0_, d_1498_p1_):
                def lambda71_(d_1499_i1_):
                    return (_dafny.SeqWithoutIsStrInference([d_1496_p2_ for d_1500_i2_ in range(default__.abs(723))])) + (_dafny.SeqWithoutIsStrInference([(D0_DC2(d_1497_p0_, d_1498_p1_)).cf5 for d_1501_i3_ in range(default__.abs(604))]))

                return lambda71_

            init36_ = lambda70_(p2, p0, p1)
            nw278_ = _dafny.Array(None, 7)
            for i0_36_ in range(nw278_.length(0)):
                nw278_[i0_36_] = init36_(i0_36_)
            d_1495_v5_ = nw278_
            d_1502_v6_: _dafny.MultiSet
            d_1502_v6_ = _dafny.MultiSet([p0, 259])
            rhs270_ = d_1495_v5_
            rhs271_ = d_1502_v6_
            rhs272_ = ((d_1489_v1_)[default__.safeIndex(417, (d_1489_v1_).length(0))]) and (not((d_1489_v1_)[default__.safeIndex(417, (d_1489_v1_).length(0))]))
            lhs243_ = globalState
            d_1495_v5_ = rhs270_
            d_1502_v6_ = rhs271_
            lhs243_.f7 = rhs272_
            d_1503_v7_: _dafny.Seq
            d_1503_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "spjomng"))
            d_1504_v8_: D0
            d_1504_v8_ = D0_DC0((d_1503_v7_) + (d_1503_v7_))
            d_1504_v8_ = d_1504_v8_
            d_1505_v9_: _dafny.Map
            d_1505_v9_ = _dafny.Map({(d_1494_v4_).ispropersubset(d_1494_v4_): p3})
            rhs273_ = ((d_1505_v9_)[not(p3)] if (not(p3)) in (d_1505_v9_) else p1)
            rhs274_ = self.f30
            rhs275_ = (0) - ((0) - ((p2) - (p0)))
            rhs276_ = d_1503_v7_
            rhs277_ = (self).fm2(globalState)
            lhs244_ = globalState
            lhs245_ = globalState
            lhs246_ = self
            lhs247_ = self
            lhs244_.f7 = rhs273_
            lhs245_.f10 = rhs274_
            lhs246_.f30 = rhs275_
            d_1503_v7_ = rhs276_
            lhs247_.f30 = rhs277_
        elif True:
            (globalState).f7 = p1
            d_1506_v10_: _dafny.Array
            nw279_ = _dafny.Array(_dafny.Array(None, 0), 8)
            d_1506_v10_ = nw279_
            d_1507_v11_: T1
            nw280_ = C4()
            nw280_.ctor__(d_1506_v10_)
            d_1507_v11_ = nw280_
            d_1508_v12_: _dafny.Map
            d_1508_v12_ = _dafny.Map({d_1507_v11_: p0})
            d_1509_v13_: _dafny.Array
            nw281_ = _dafny.Array(None, 3)
            nw281_[int(0)] = _dafny.Map({d_1507_v11_: p0})
            nw281_[int(1)] = d_1508_v12_
            nw281_[int(2)] = d_1508_v12_
            d_1509_v13_ = nw281_
            index283_ = default__.safeIndex(587, (d_1509_v13_).length(0))
            (d_1509_v13_)[index283_] = d_1508_v12_
            index284_ = default__.safeIndex(587, (d_1509_v13_).length(0))
            (d_1509_v13_)[index284_] = (d_1508_v12_) | (d_1508_v12_)
            d_1510_v14_: _dafny.Array
            def lambda72_(d_1511_i4_):
                return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qj"))

            init37_ = lambda72_
            nw282_ = _dafny.Array(None, 23)
            for i0_37_ in range(nw282_.length(0)):
                nw282_[i0_37_] = init37_(i0_37_)
            d_1510_v14_ = nw282_
            d_1512_v15_: _dafny.Seq
            d_1512_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vkarjbt"))
            index285_ = default__.safeIndex(929, (d_1510_v14_).length(0))
            (d_1510_v14_)[index285_] = (d_1512_v15_) + (d_1512_v15_)
            index286_ = default__.safeIndex(929, (d_1510_v14_).length(0))
            (d_1510_v14_)[index286_] = default__.fm35(603, p1, (self.f30) != (162), globalState)
            d_1513_v17_: _dafny.MultiSet
            def iife135_():
                coll43_ = _dafny.Map()
                compr_43_: int
                for compr_43_ in _dafny.IntegerRange(469, 878):
                    d_1514_v16_: int = compr_43_
                    if ((469) <= (d_1514_v16_)) and ((d_1514_v16_) < (878)):
                        coll43_[(d_1514_v16_) + (self.f30)] = p1
                return _dafny.Map(coll43_)
            d_1513_v17_ = _dafny.MultiSet([iife135_()
            ])
            d_1515_v18_: _dafny.Seq
            d_1515_v18_ = _dafny.SeqWithoutIsStrInference([d_1513_v17_])
            d_1516_v19_: _dafny.Map
            d_1516_v19_ = _dafny.Map({p2: self.f30})
            d_1517_v20_: _dafny.Array
            def lambda73_(d_1518_p1_):
                def lambda74_(d_1519_i5_):
                    return d_1518_p1_

                return lambda74_

            init38_ = lambda73_(p1)
            nw283_ = _dafny.Array(None, 7)
            for i0_38_ in range(nw283_.length(0)):
                nw283_[i0_38_] = init38_(i0_38_)
            d_1517_v20_ = nw283_
            d_1520_v21_: _dafny.Map
            d_1520_v21_ = _dafny.Map({d_1516_v19_: d_1517_v20_})
            d_1521_v22_: _dafny.Map
            d_1521_v22_ = _dafny.Map({p1: p0})
            d_1522_v23_: _dafny.Map
            d_1522_v23_ = _dafny.Map({d_1521_v22_: p3})
            d_1523_v24_: _dafny.Map
            d_1523_v24_ = _dafny.Map({p0: p3})
            d_1524_v25_: D2
            d_1524_v25_ = D2_DC7(_dafny.Map({((d_1516_v19_)[self.f30] if (self.f30) in (d_1516_v19_) else p2): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h")))}))
            d_1525_v26_: _dafny.Map
            d_1525_v26_ = _dafny.Map({d_1524_v25_: p0})
            d_1526_v27_: _dafny.Seq
            d_1526_v27_ = _dafny.SeqWithoutIsStrInference([d_1525_v26_, d_1525_v26_, _dafny.Map({D2_DC7(d_1516_v19_): p0}), d_1525_v26_, d_1525_v26_])
            d_1527_v28_: _dafny.Map
            d_1527_v28_ = _dafny.Map({((d_1515_v18_)[default__.safeIndex(len(_dafny.Map({len(d_1520_v21_): p0})), len(d_1515_v18_))] if p1 else (_dafny.MultiSet([_dafny.Map({((d_1521_v22_)[p1] if (p1) in (d_1521_v22_) else len(d_1522_v23_)): p1}), d_1523_v24_])).set(d_1523_v24_, default__.abs(p2))): d_1526_v27_})
            d_1527_v28_ = (d_1527_v28_).set(d_1513_v17_, _dafny.SeqWithoutIsStrInference([d_1525_v26_ for d_1528_i6_ in range(default__.abs(11))]))
            d_1529_v29_: str
            d_1529_v29_ = _dafny.CodePoint('c')
            index287_ = default__.safeIndex(929, (d_1510_v14_).length(0))
            (d_1510_v14_)[index287_] = (default__.fm23(D2_DC7(d_1516_v19_), (0) - (((d_1521_v22_)[p1] if (p1) in (d_1521_v22_) else p2)), (0) - (p0), globalState)).set(default__.safeIndex(p0, len(default__.fm23(D2_DC7(d_1516_v19_), (0) - (((d_1521_v22_)[p1] if (p1) in (d_1521_v22_) else p2)), (0) - (p0), globalState))), d_1529_v29_)
        (globalState).f10 = self.f30
        d_1530_v30_: _dafny.Array
        nw284_ = _dafny.Array(int(0), 13)
        d_1530_v30_ = nw284_
        d_1531_v31_: str
        d_1531_v31_ = _dafny.CodePoint('a')
        d_1532_v32_: _dafny.MultiSet
        d_1532_v32_ = _dafny.MultiSet([d_1531_v31_, d_1531_v31_, d_1531_v31_])
        index288_ = default__.safeIndex(535, (d_1530_v30_).length(0))
        (d_1530_v30_)[index288_] = (d_1532_v32_).cardinality
        index289_ = default__.safeIndex(16, (d_1530_v30_).length(0))
        (d_1530_v30_)[index289_] = default__.safeModuloInt(p2, self.f30)
        d_1533_v33_: _dafny.Map
        d_1533_v33_ = _dafny.Map({not(p3): 554})
        d_1534_v34_: _dafny.MultiSet
        d_1534_v34_ = _dafny.MultiSet([len(d_1533_v33_)])
        d_1535_v36_: D3
        d_1535_v36_ = D3_DC11(p0)
        d_1536_v37_: _dafny.Seq
        def iife136_():
            coll44_ = _dafny.Map()
            compr_44_: int
            for compr_44_ in (default__.fm14(p1, p2, (d_1535_v36_).cf22, globalState)).Elements:
                d_1537_v35_: int = compr_44_
                if (d_1537_v35_) in (default__.fm14(p1, p2, (d_1535_v36_).cf22, globalState)):
                    coll44_[default__.safeModuloInt(d_1537_v35_, len(_dafny.Map({p1: d_1534_v34_})))] = d_1531_v31_
            return _dafny.Map(coll44_)
        d_1536_v37_ = _dafny.SeqWithoutIsStrInference([(d_1534_v34_).cardinality, len(iife136_()
        ), (self.f30) + (default__.fm0(p0, globalState))])
        index290_ = default__.safeIndex(535, (d_1530_v30_).length(0))
        index291_ = default__.safeIndex(16, (d_1530_v30_).length(0))
        rhs278_ = (0) - ((p0) - (p2))
        rhs279_ = p2
        rhs280_ = (d_1536_v37_)[default__.safeIndex((p0) * (p0), len(d_1536_v37_))]
        rhs281_ = (len((_dafny.SeqWithoutIsStrInference([(0) - (p0), -948])) + (d_1536_v37_))) * ((p2) + (self.f30))
        rhs282_ = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iqxo")))
        lhs248_ = d_1530_v30_
        lhs249_ = default__.safeIndex(535, (d_1530_v30_).length(0))
        lhs250_ = self
        lhs251_ = d_1530_v30_
        lhs252_ = default__.safeIndex(16, (d_1530_v30_).length(0))
        lhs253_ = self
        lhs254_ = globalState
        lhs248_[lhs249_] = rhs278_
        lhs250_.f30 = rhs279_
        lhs251_[lhs252_] = rhs280_
        lhs253_.f30 = rhs281_
        lhs254_.f10 = rhs282_
        d_1538_v38_: D0
        d_1538_v38_ = D0_DC2((d_1530_v30_)[default__.safeIndex(535, (d_1530_v30_).length(0))], not(p3))
        d_1538_v38_ = d_1538_v38_
        d_1539_v39_: C3
        nw285_ = C3()
        nw285_.ctor__()
        d_1539_v39_ = nw285_
        d_1540_v40_: _dafny.MultiSet
        d_1540_v40_ = _dafny.MultiSet([d_1539_v39_])
        d_1541_v41_: _dafny.Map
        d_1541_v41_ = _dafny.Map({(d_1540_v40_).cardinality: (d_1539_v39_).fm8(p2, -44, p0, globalState)})
        d_1541_v41_ = (d_1541_v41_).set(744, (d_1530_v30_)[default__.safeIndex(535, (d_1530_v30_).length(0))])
        d_1542_i7_: int
        d_1542_i7_ = 0
        with _dafny.label("6"):
            while p3:
                with _dafny.c_label("6"):
                    if (d_1542_i7_) >= (100):
                        raise _dafny.Break("6")
                    d_1542_i7_ = (d_1542_i7_) + (1)
                    if p3:
                        d_1543_v42_: _dafny.Array
                        def lambda75_(d_1544_p3_, d_1545_p1_):
                            def lambda76_(d_1546_i8_):
                                return (d_1545_p1_ if d_1544_p3_ else d_1544_p3_)

                            return lambda76_

                        init39_ = lambda75_(p3, p1)
                        nw286_ = _dafny.Array(None, 11)
                        for i0_39_ in range(nw286_.length(0)):
                            nw286_[i0_39_] = init39_(i0_39_)
                        d_1543_v42_ = nw286_
                        index292_ = default__.safeIndex(491, (d_1543_v42_).length(0))
                        (d_1543_v42_)[index292_] = p1
                        d_1547_v43_: _dafny.Seq
                        d_1547_v43_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fxrsbmfcb"))
                        d_1548_v44_: D4
                        d_1548_v44_ = D4_DC14(p1)
                        d_1549_v45_: D4
                        d_1549_v45_ = D4_DC15(d_1548_v44_)
                        d_1550_v46_: D7
                        d_1550_v46_ = D7_DC20((_dafny.SeqWithoutIsStrInference([self.f30])).set(default__.safeIndex(len(d_1541_v41_), len(_dafny.SeqWithoutIsStrInference([self.f30]))), len(d_1547_v43_)), d_1543_v42_, d_1549_v45_, p3, self.f30)
                        d_1551_v47_: D0
                        d_1551_v47_ = D0_DC0(d_1547_v43_)
                        index293_ = default__.safeIndex(491, (d_1543_v42_).length(0))
                        rhs283_ = (d_1550_v46_).cf33
                        rhs284_ = p3
                        rhs285_ = ((d_1547_v43_) + ((d_1551_v47_).cf0)) == ((d_1547_v43_) + (_dafny.SeqWithoutIsStrInference([d_1531_v31_ for d_1552_i9_ in range(default__.abs(936))])))
                        rhs286_ = p2
                        lhs255_ = d_1543_v42_
                        lhs256_ = default__.safeIndex(491, (d_1543_v42_).length(0))
                        lhs257_ = globalState
                        lhs258_ = globalState
                        lhs259_ = globalState
                        lhs255_[lhs256_] = rhs283_
                        lhs257_.f7 = rhs284_
                        lhs258_.f13 = rhs285_
                        lhs259_.f10 = rhs286_
                        d_1553_v49_: _dafny.Seq
                        d_1553_v49_ = _dafny.SeqWithoutIsStrInference([(d_1543_v42_)[default__.safeIndex(491, (d_1543_v42_).length(0))]])
                        d_1554_v50_: _dafny.Seq
                        d_1554_v50_ = _dafny.SeqWithoutIsStrInference([d_1553_v49_])
                        pat_let_tv44_ = d_1536_v37_
                        pat_let_tv45_ = d_1554_v50_
                        pat_let_tv46_ = d_1554_v50_
                        pat_let_tv47_ = d_1543_v42_
                        pat_let_tv48_ = d_1543_v42_
                        pat_let_tv49_ = d_1530_v30_
                        pat_let_tv50_ = d_1530_v30_
                        def iife137_(_pat_let46_0):
                            def iife138_(d_1555_dt__update__tmp_h0_):
                                def iife139_(_pat_let47_0):
                                    def iife140_(d_1556_dt__update_hcf30_h0_):
                                        def iife142_():
                                            coll45_ = _dafny.Map()
                                            compr_45_: _dafny.Seq
                                            for compr_45_ in (_dafny.MultiSet(pat_let_tv45_)).Elements:
                                                d_1557_v48_: _dafny.Seq = compr_45_
                                                if (d_1557_v48_) in (_dafny.MultiSet(pat_let_tv46_)):
                                                    coll45_[d_1557_v48_] = (pat_let_tv48_)[default__.safeIndex(491, (pat_let_tv47_).length(0))]
                                            return _dafny.Map(coll45_)
                                        def iife141_(_pat_let48_0):
                                            def iife143_(d_1558_dt__update_hcf34_h0_):
                                                return D7_DC20(d_1556_dt__update_hcf30_h0_, (d_1555_dt__update__tmp_h0_).cf31, (d_1555_dt__update__tmp_h0_).cf32, (d_1555_dt__update__tmp_h0_).cf33, d_1558_dt__update_hcf34_h0_)
                                            return iife143_(_pat_let48_0)
                                        return iife141_(default__.safeDivisionInt(len(iife142_()
                                        ), (pat_let_tv50_)[default__.safeIndex(535, (pat_let_tv49_).length(0))]))
                                    return iife140_(_pat_let47_0)
                                return iife139_(pat_let_tv44_)
                            return iife138_(_pat_let46_0)
                        d_1550_v46_ = iife137_(d_1550_v46_)
                        d_1559_v51_: _dafny.Array
                        def lambda77_(d_1560_v49_, d_1561_v42_):
                            def lambda78_(d_1562_i10_):
                                return ((D8_DC21(d_1560_v49_)).cf35) + (_dafny.SeqWithoutIsStrInference([(d_1561_v42_)[default__.safeIndex(491, (d_1561_v42_).length(0))]]))

                            return lambda78_

                        init40_ = lambda77_(d_1553_v49_, d_1543_v42_)
                        nw287_ = _dafny.Array(None, 25)
                        for i0_40_ in range(nw287_.length(0)):
                            nw287_[i0_40_] = init40_(i0_40_)
                        d_1559_v51_ = nw287_
                        d_1559_v51_ = d_1559_v51_
                        d_1563_v52_: _dafny.Array
                        nw288_ = _dafny.Array(D1.default()(), 14)
                        d_1563_v52_ = nw288_
                        d_1564_v53_: _dafny.Map
                        d_1564_v53_ = _dafny.Map({(d_1543_v42_)[default__.safeIndex(491, (d_1543_v42_).length(0))]: d_1563_v52_})
                        d_1565_v54_: _dafny.Seq
                        d_1565_v54_ = _dafny.SeqWithoutIsStrInference([((d_1564_v53_)[False] if (False) in (d_1564_v53_) else d_1563_v52_), d_1563_v52_, d_1563_v52_, d_1563_v52_])
                        d_1566_v55_: _dafny.Array
                        nw289_ = _dafny.Array(None, 6)
                        nw289_[int(0)] = d_1563_v52_
                        nw289_[int(1)] = (d_1565_v54_)[default__.safeIndex((d_1536_v37_)[default__.safeIndex(self.f30, len(d_1536_v37_))], len(d_1565_v54_))]
                        nw289_[int(2)] = d_1563_v52_
                        nw289_[int(3)] = d_1563_v52_
                        nw289_[int(4)] = d_1563_v52_
                        nw289_[int(5)] = d_1563_v52_
                        d_1566_v55_ = nw289_
                        d_1567_v56_: T1
                        nw290_ = C4()
                        nw290_.ctor__(d_1566_v55_)
                        d_1567_v56_ = nw290_
                        d_1567_v56_ = d_1567_v56_
                        d_1568_v57_: _dafny.Array
                        nw291_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 11)
                        d_1568_v57_ = nw291_
                        index294_ = default__.safeIndex(248, (d_1568_v57_).length(0))
                        (d_1568_v57_)[index294_] = _dafny.SeqWithoutIsStrInference([(d_1547_v43_)[default__.safeIndex(len(_dafny.Map({d_1531_v31_: p2})), len(d_1547_v43_))] for d_1569_i11_ in range(default__.abs(427))])
                        index295_ = default__.safeIndex(248, (d_1568_v57_).length(0))
                        rhs287_ = default__.fm34(p1, d_1531_v31_, ((d_1539_v39_).fm2(globalState) if (d_1543_v42_)[default__.safeIndex(491, (d_1543_v42_).length(0))] else len(d_1533_v33_)), globalState)
                        rhs288_ = (_dafny.SeqWithoutIsStrInference([d_1531_v31_ for d_1570_i12_ in range(default__.abs(-404))])) + (d_1547_v43_)
                        lhs260_ = d_1568_v57_
                        lhs261_ = default__.safeIndex(248, (d_1568_v57_).length(0))
                        d_1553_v49_ = rhs287_
                        lhs260_[lhs261_] = rhs288_
                    elif True:
                        d_1571_v58_: _dafny.Array
                        def lambda79_(d_1572_v36_):
                            def lambda80_(d_1573_i13_):
                                return d_1572_v36_

                            return lambda80_

                        init41_ = lambda79_(d_1535_v36_)
                        nw292_ = _dafny.Array(None, 1)
                        for i0_41_ in range(nw292_.length(0)):
                            nw292_[i0_41_] = init41_(i0_41_)
                        d_1571_v58_ = nw292_
                        index296_ = default__.safeIndex(34, (d_1571_v58_).length(0))
                        (d_1571_v58_)[index296_] = d_1535_v36_
                        pat_let_tv51_ = d_1536_v37_
                        pat_let_tv52_ = d_1536_v37_
                        index297_ = default__.safeIndex(34, (d_1571_v58_).length(0))
                        def iife144_(_pat_let49_0):
                            def iife145_(d_1574_dt__update__tmp_h1_):
                                def iife146_(_pat_let50_0):
                                    def iife147_(d_1575_dt__update_hcf22_h0_):
                                        return D3_DC11(d_1575_dt__update_hcf22_h0_)
                                    return iife147_(_pat_let50_0)
                                return iife146_((pat_let_tv51_)[default__.safeIndex(323, len(pat_let_tv52_))])
                            return iife145_(_pat_let49_0)
                        (d_1571_v58_)[index297_] = iife144_(default__.fm40((d_1534_v34_).cardinality, False, p3, globalState))
                        index298_ = default__.safeIndex(535, (d_1530_v30_).length(0))
                        index299_ = default__.safeIndex(535, (d_1530_v30_).length(0))
                        rhs289_ = (d_1530_v30_)[default__.safeIndex(535, (d_1530_v30_).length(0))]
                        rhs290_ = p0
                        rhs291_ = self.f30
                        lhs262_ = d_1530_v30_
                        lhs263_ = default__.safeIndex(535, (d_1530_v30_).length(0))
                        lhs264_ = d_1530_v30_
                        lhs265_ = default__.safeIndex(535, (d_1530_v30_).length(0))
                        lhs266_ = globalState
                        lhs262_[lhs263_] = rhs289_
                        lhs264_[lhs265_] = rhs290_
                        lhs266_.f10 = rhs291_
                        (globalState).f7 = p3
                        d_1576_v59_: C1
                        nw293_ = C1()
                        nw293_.ctor__()
                        d_1576_v59_ = nw293_
                        d_1577_v60_: _dafny.Map
                        d_1577_v60_ = _dafny.Map({d_1576_v59_: p1})
                        d_1578_v61_: _dafny.Map
                        d_1578_v61_ = _dafny.Map({d_1531_v31_: p1})
                        d_1577_v60_ = (d_1577_v60_).set(d_1576_v59_, ((d_1578_v61_)[d_1531_v31_] if (d_1531_v31_) in (d_1578_v61_) else True))
                        d_1579_v62_: _dafny.Map
                        d_1579_v62_ = _dafny.Map({478: d_1533_v33_})
                        index300_ = default__.safeIndex(535, (d_1530_v30_).length(0))
                        (d_1530_v30_)[index300_] = (((d_1530_v30_)[default__.safeIndex(535, (d_1530_v30_).length(0))] if p1 else len(d_1579_v62_))) * (p2)
                    d_1580_v63_: _dafny.Seq
                    d_1580_v63_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v"))
                    rhs292_ = (_dafny.SeqWithoutIsStrInference([d_1531_v31_ for d_1581_i14_ in range(default__.abs(46))])).set(default__.safeIndex(self.f30, len(_dafny.SeqWithoutIsStrInference([d_1531_v31_ for d_1582_i14_ in range(default__.abs(46))]))), d_1531_v31_)
                    rhs293_ = ((d_1534_v34_) - (_dafny.MultiSet(d_1536_v37_))).cardinality
                    lhs267_ = globalState
                    d_1580_v63_ = rhs292_
                    lhs267_.f10 = rhs293_
                    if p3:
                        d_1583_v64_: _dafny.Array
                        nw294_ = _dafny.Array(D1.default()(), 1)
                        d_1583_v64_ = nw294_
                        d_1584_v65_: _dafny.Array
                        nw295_ = _dafny.Array(None, 10)
                        nw295_[int(0)] = d_1583_v64_
                        nw295_[int(1)] = d_1583_v64_
                        nw295_[int(2)] = d_1583_v64_
                        nw295_[int(3)] = d_1583_v64_
                        nw295_[int(4)] = d_1583_v64_
                        nw295_[int(5)] = d_1583_v64_
                        nw295_[int(6)] = d_1583_v64_
                        nw295_[int(7)] = d_1583_v64_
                        nw295_[int(8)] = d_1583_v64_
                        nw295_[int(9)] = d_1583_v64_
                        d_1584_v65_ = nw295_
                        d_1585_v66_: C4
                        nw296_ = C4()
                        nw296_.ctor__(d_1584_v65_)
                        d_1585_v66_ = nw296_
                        d_1586_v67_: _dafny.MultiSet
                        d_1586_v67_ = _dafny.MultiSet([d_1530_v30_, d_1530_v30_])
                        (globalState).f10 = ((d_1586_v67_)[d_1530_v30_] if (d_1530_v30_) in (d_1586_v67_) else self.f30)
                        (globalState).f10 = -41
                        (globalState).f13 = not(((p0) - (p0)) < (self.f30))
                        d_1587_v68_: _dafny.Seq
                        d_1587_v68_ = _dafny.SeqWithoutIsStrInference([p3])
                        d_1533_v33_ = (d_1533_v33_).set(not (p3) or (p3), (len(d_1587_v68_)) - (self.f30))
                    elif True:
                        d_1588_v69_: _dafny.Map
                        d_1588_v69_ = _dafny.Map({d_1530_v30_: (self).fm4(d_1580_v63_, self.f30, globalState)})
                        d_1589_v70_: _dafny.Seq
                        d_1589_v70_ = _dafny.SeqWithoutIsStrInference([d_1588_v69_, d_1588_v69_, d_1588_v69_, d_1588_v69_, _dafny.Map({d_1530_v30_: p3})])
                        d_1590_v71_: _dafny.Seq
                        d_1590_v71_ = _dafny.SeqWithoutIsStrInference([d_1530_v30_, d_1530_v30_])
                        d_1588_v69_ = ((d_1589_v70_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([881])), len(d_1589_v70_))]) | (_dafny.Map({(d_1590_v71_)[default__.safeIndex(self.f30, len(d_1590_v71_))]: True}))
                        d_1591_v72_: _dafny.Seq
                        d_1591_v72_ = _dafny.SeqWithoutIsStrInference([D4_DC14(p1)])
                        d_1592_v73_: _dafny.Map
                        d_1592_v73_ = _dafny.Map({d_1591_v72_: self.f30})
                        rhs294_ = d_1592_v73_
                        rhs295_ = d_1580_v63_
                        d_1592_v73_ = rhs294_
                        d_1580_v63_ = rhs295_
                        d_1593_v74_: _dafny.Array
                        nw297_ = _dafny.Array(D7.default()(), 22)
                        d_1593_v74_ = nw297_
                        d_1594_v75_: _dafny.Map
                        d_1594_v75_ = _dafny.Map({p1: True})
                        d_1595_v76_: _dafny.Seq
                        d_1595_v76_ = _dafny.SeqWithoutIsStrInference([d_1594_v75_])
                        d_1596_v77_: _dafny.Array
                        nw298_ = _dafny.Array(None, 26)
                        nw298_[int(0)] = p3
                        nw298_[int(1)] = True
                        nw298_[int(2)] = not(p1)
                        nw298_[int(3)] = p1
                        nw298_[int(4)] = False
                        nw298_[int(5)] = p3
                        nw298_[int(6)] = p1
                        nw298_[int(7)] = not(True)
                        nw298_[int(8)] = p3
                        nw298_[int(9)] = p3
                        nw298_[int(10)] = p3
                        nw298_[int(11)] = p3
                        nw298_[int(12)] = (self).fm4(d_1580_v63_, len(d_1595_v76_), globalState)
                        nw298_[int(13)] = p1
                        nw298_[int(14)] = p1
                        nw298_[int(15)] = p1
                        nw298_[int(16)] = p3
                        nw298_[int(17)] = not(False)
                        nw298_[int(18)] = False
                        nw298_[int(19)] = p3
                        nw298_[int(20)] = p1
                        nw298_[int(21)] = (self).fm4(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "whc")), len(_dafny.Map({(d_1530_v30_)[default__.safeIndex(535, (d_1530_v30_).length(0))]: 138})), globalState)
                        nw298_[int(22)] = p1
                        nw298_[int(23)] = p3
                        nw298_[int(24)] = p3
                        nw298_[int(25)] = p3
                        d_1596_v77_ = nw298_
                        d_1597_v78_: D4
                        d_1597_v78_ = D4_DC14(p3)
                        d_1598_v79_: D4
                        d_1598_v79_ = D4_DC15(d_1597_v78_)
                        d_1599_v80_: D7
                        d_1599_v80_ = D7_DC20(d_1536_v37_, d_1596_v77_, d_1598_v79_, p3, p0)
                        index301_ = default__.safeIndex(516, (d_1593_v74_).length(0))
                        (d_1593_v74_)[index301_] = d_1599_v80_
                        index302_ = default__.safeIndex(516, (d_1593_v74_).length(0))
                        (d_1593_v74_)[index302_] = D7_DC20((_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "flxdys"))) for d_1600_i15_ in range(default__.abs(312))])) + (_dafny.SeqWithoutIsStrInference([p0 for d_1601_i16_ in range(default__.abs(-914))])), d_1596_v77_, d_1598_v79_, p3, p0)
                        d_1602_v81_: _dafny.Map
                        d_1602_v81_ = _dafny.Map({(d_1580_v63_) + (default__.fm10(d_1541_v41_, d_1531_v31_, p3, globalState)): _dafny.SeqWithoutIsStrInference([(d_1530_v30_)[default__.safeIndex(535, (d_1530_v30_).length(0))]])})
                        d_1602_v81_ = (d_1602_v81_).set(d_1580_v63_, d_1536_v37_)
                        d_1603_v82_: _dafny.Map
                        d_1603_v82_ = _dafny.Map({680: d_1596_v77_})
                        d_1596_v77_ = ((d_1603_v82_)[(d_1530_v30_)[default__.safeIndex(535, (d_1530_v30_).length(0))]] if ((d_1530_v30_)[default__.safeIndex(535, (d_1530_v30_).length(0))]) in (d_1603_v82_) else d_1596_v77_)
                    d_1604_v83_: _dafny.Array
                    nw299_ = _dafny.Array(False, 29)
                    d_1604_v83_ = nw299_
                    index303_ = default__.safeIndex(222, (d_1604_v83_).length(0))
                    (d_1604_v83_)[index303_] = ((d_1530_v30_)[default__.safeIndex(535, (d_1530_v30_).length(0))]) >= ((d_1530_v30_)[default__.safeIndex(535, (d_1530_v30_).length(0))])
                    index304_ = default__.safeIndex(222, (d_1604_v83_).length(0))
                    (d_1604_v83_)[index304_] = (self).fm4(d_1580_v63_, len(d_1580_v63_), globalState)
                    pass
            pass

    def m1(self, p0, p1, p2, globalState):
        r0: D0 = D0.default()()
        r1: bool = False
        d_1605_v0_: _dafny.Map
        d_1605_v0_ = _dafny.Map({self.f30: p1})
        hi7_ = default__.safeDivisionInt(p1, len(d_1605_v0_))
        for d_1606_i0_ in range(p1, hi7_):
            d_1607_v1_: _dafny.Seq
            d_1607_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ctawxugn"))
            d_1608_v2_: _dafny.Seq
            d_1608_v2_ = _dafny.SeqWithoutIsStrInference([p2, p2])
            d_1609_v3_: _dafny.Map
            d_1609_v3_ = _dafny.Map({(self).fm4(d_1607_v1_, (_dafny.MultiSet(d_1608_v2_)).cardinality, globalState): -828})
            d_1610_v4_: _dafny.MultiSet
            d_1610_v4_ = _dafny.MultiSet([311])
            d_1611_v5_: D2
            d_1611_v5_ = D2_DC7(d_1605_v0_)
            d_1612_v6_: C2
            nw300_ = C2()
            nw300_.ctor__(d_1607_v1_)
            d_1612_v6_ = nw300_
            d_1613_v7_: _dafny.MultiSet
            d_1613_v7_ = _dafny.MultiSet([d_1612_v6_])
            d_1614_v8_: _dafny.Map
            d_1614_v8_ = _dafny.Map({len(p0): d_1613_v7_})
            d_1615_v9_: _dafny.Array
            nw301_ = _dafny.Array(None, 27)
            nw301_[int(0)] = not(True)
            nw301_[int(1)] = p2
            nw301_[int(2)] = (d_1606_i0_) == (p1)
            nw301_[int(3)] = not(p2)
            nw301_[int(4)] = p2
            nw301_[int(5)] = p2
            nw301_[int(6)] = (self).fm4(d_1607_v1_, self.f30, globalState)
            nw301_[int(7)] = p2
            nw301_[int(8)] = p2
            nw301_[int(9)] = p2
            nw301_[int(10)] = (len(d_1605_v0_)) < (232)
            nw301_[int(11)] = (p2 if not(p2) else p2)
            nw301_[int(12)] = not (p2) or (p2)
            nw301_[int(13)] = p2
            nw301_[int(14)] = (len(_dafny.Map({True: p2}))) >= (len(d_1609_v3_))
            nw301_[int(15)] = (801) >= ((_dafny.MultiSet([len(d_1608_v2_), 657, p1, 56])).cardinality)
            nw301_[int(16)] = (d_1610_v4_).ispropersubset(_dafny.MultiSet([d_1606_i0_, (0) - (d_1606_i0_), len(d_1607_v1_), len(default__.fm23(d_1611_v5_, len(_dafny.SeqWithoutIsStrInference([len(d_1607_v1_)])), self.f30, globalState)), len(d_1608_v2_)]))
            nw301_[int(17)] = p2
            nw301_[int(18)] = (self.f30) > (p1)
            nw301_[int(19)] = p2
            nw301_[int(20)] = p2
            nw301_[int(21)] = p2
            nw301_[int(22)] = p2
            nw301_[int(23)] = p2
            nw301_[int(24)] = (d_1606_i0_) in (d_1614_v8_)
            nw301_[int(25)] = (((d_1609_v3_)[not(not(not(p2)))] if (not(not(not(p2)))) in (d_1609_v3_) else d_1606_i0_)) >= (self.f30)
            nw301_[int(26)] = p2
            d_1615_v9_ = nw301_
            rhs296_ = (d_1606_i0_) - (376)
            rhs297_ = d_1615_v9_
            rhs298_ = (d_1608_v2_)[default__.safeIndex(d_1606_i0_, len(d_1608_v2_))]
            lhs268_ = globalState
            lhs269_ = globalState
            lhs268_.f10 = rhs296_
            d_1615_v9_ = rhs297_
            lhs269_.f7 = rhs298_
            d_1616_v10_: str
            d_1616_v10_ = _dafny.CodePoint('b')
            rhs299_ = d_1606_i0_
            rhs300_ = d_1616_v10_
            rhs301_ = p1
            lhs270_ = globalState
            lhs271_ = globalState
            lhs272_ = self
            lhs270_.f10 = rhs299_
            lhs271_.f19 = rhs300_
            lhs272_.f30 = rhs301_
            d_1617_v11_: C3
            nw302_ = C3()
            nw302_.ctor__()
            d_1617_v11_ = nw302_
            d_1618_v12_: D0
            d_1618_v12_ = D0_DC2(p1, p2)
            d_1619_v13_: D0
            d_1619_v13_ = D0_DC3(d_1618_v12_)
            d_1620_v14_: D0
            d_1620_v14_ = D0_DC3(d_1618_v12_)
            pat_let_tv53_ = d_1619_v13_
            pat_let_tv54_ = d_1618_v12_
            d_1621_v15_: _dafny.Array
            nw303_ = _dafny.Array(None, 20)
            nw303_[int(0)] = d_1620_v14_
            nw303_[int(1)] = d_1620_v14_
            nw303_[int(2)] = d_1620_v14_
            nw303_[int(3)] = d_1620_v14_
            nw303_[int(4)] = D0_DC3(d_1619_v13_)
            def iife148_(_pat_let51_0):
                def iife149_(d_1622_dt__update__tmp_h0_):
                    def iife150_(_pat_let52_0):
                        def iife151_(d_1623_dt__update_hcf7_h0_):
                            return D0_DC3(d_1623_dt__update_hcf7_h0_)
                        return iife151_(_pat_let52_0)
                    return iife150_(pat_let_tv53_)
                return iife149_(_pat_let51_0)
            nw303_[int(5)] = iife148_(d_1620_v14_)
            def iife152_(_pat_let53_0):
                def iife153_(d_1624_dt__update__tmp_h1_):
                    def iife154_(_pat_let54_0):
                        def iife155_(d_1625_dt__update_hcf7_h1_):
                            return D0_DC3(d_1625_dt__update_hcf7_h1_)
                        return iife155_(_pat_let54_0)
                    return iife154_(pat_let_tv54_)
                return iife153_(_pat_let53_0)
            nw303_[int(6)] = iife152_(d_1620_v14_)
            nw303_[int(7)] = d_1620_v14_
            nw303_[int(8)] = d_1620_v14_
            nw303_[int(9)] = d_1620_v14_
            nw303_[int(10)] = D0_DC3(d_1619_v13_)
            nw303_[int(11)] = d_1620_v14_
            nw303_[int(12)] = D0_DC3(d_1618_v12_)
            nw303_[int(13)] = D0_DC3(d_1618_v12_)
            nw303_[int(14)] = d_1620_v14_
            nw303_[int(15)] = d_1620_v14_
            nw303_[int(16)] = d_1620_v14_
            nw303_[int(17)] = d_1620_v14_
            nw303_[int(18)] = d_1620_v14_
            nw303_[int(19)] = D0_DC3(d_1619_v13_)
            d_1621_v15_ = nw303_
            index305_ = default__.safeIndex(205, (d_1621_v15_).length(0))
            (d_1621_v15_)[index305_] = d_1620_v14_
            index306_ = default__.safeIndex(205, (d_1621_v15_).length(0))
            (d_1621_v15_)[index306_] = D0_DC3(d_1619_v13_)
        d_1626_v17_: _dafny.MultiSet
        d_1626_v17_ = _dafny.MultiSet([self.f30])
        d_1627_v18_: D2
        def iife156_():
            coll46_ = _dafny.Map()
            compr_46_: int
            for compr_46_ in (d_1626_v17_).Elements:
                d_1628_v16_: int = compr_46_
                if (d_1628_v16_) in (d_1626_v17_):
                    coll46_[(d_1628_v16_) + (645)] = p1
            return _dafny.Map(coll46_)
        d_1627_v18_ = D2_DC7(iife156_()
)
        d_1629_v19_: _dafny.Map
        d_1629_v19_ = _dafny.Map({p2: p2})
        d_1630_v20_: D3
        d_1630_v20_ = D3_DC12((self).fm4(default__.fm23(d_1627_v18_, self.f30, len(d_1629_v19_), globalState), p1, globalState))
        d_1630_v20_ = d_1630_v20_
        d_1631_v21_: _dafny.Array
        def lambda81_(d_1632_p1_):
            def lambda82_(d_1633_i1_):
                return (d_1633_i1_) + (d_1632_p1_)

            return lambda82_

        init42_ = lambda81_(p1)
        nw304_ = _dafny.Array(None, 20)
        for i0_42_ in range(nw304_.length(0)):
            nw304_[i0_42_] = init42_(i0_42_)
        d_1631_v21_ = nw304_
        if (_dafny.MultiSet([d_1631_v21_, d_1631_v21_])).isdisjoint(_dafny.MultiSet([d_1631_v21_])):
            d_1634_v22_: _dafny.MultiSet
            d_1634_v22_ = _dafny.MultiSet([p2, True, p2, default__.fm12(p2, p2, globalState), p2])
            d_1635_v23_: D1
            d_1635_v23_ = D1_DC4(_dafny.Map({d_1634_v22_: p2}))
            source22_ = d_1635_v23_
            if source22_.is_DC5:
                d_1636___mcc_h0_ = source22_.cf9
                d_1637___mcc_h1_ = source22_.cf10
                d_1638___mcc_h2_ = source22_.cf11
                d_1639___mcc_h3_ = source22_.cf12
                d_1640_cf12_ = d_1639___mcc_h3_
                d_1641_cf11_ = d_1638___mcc_h2_
                d_1642_cf10_ = d_1637___mcc_h1_
                d_1643_cf9_ = d_1636___mcc_h0_
                d_1626_v17_ = d_1626_v17_
                d_1644_v24_: D8
                d_1644_v24_ = D8_DC22(98, d_1641_cf11_, _dafny.CodePoint('b'))
                d_1645_v25_: str
                d_1645_v25_ = _dafny.CodePoint('o')
                d_1646_v26_: _dafny.Seq
                d_1646_v26_ = _dafny.SeqWithoutIsStrInference([d_1644_v24_, d_1644_v24_, d_1644_v24_, D8_DC22(d_1643_cf9_, p2, d_1645_v25_)])
                d_1646_v26_ = d_1646_v26_
                d_1647_v27_: _dafny.Array
                nw305_ = _dafny.Array(_dafny.Array(None, 0), 29)
                d_1647_v27_ = nw305_
                d_1648_v28_: T1
                nw306_ = C4()
                nw306_.ctor__(d_1647_v27_)
                d_1648_v28_ = nw306_
                d_1649_v29_: _dafny.Seq
                d_1649_v29_ = _dafny.SeqWithoutIsStrInference([d_1641_cf11_])
                d_1650_v30_: _dafny.Seq
                d_1650_v30_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ra"))
                d_1651_v32_: D8
                d_1651_v32_ = D8_DC21(d_1649_v29_)
                d_1652_v33_: _dafny.Seq
                d_1652_v33_ = _dafny.SeqWithoutIsStrInference([d_1651_v32_])
                d_1653_v34_: _dafny.Map
                d_1653_v34_ = _dafny.Map({d_1651_v32_: len(d_1650_v30_)})
                d_1654_v35_: _dafny.Array
                nw307_ = _dafny.Array(None, 5)
                def iife157_():
                    coll47_ = _dafny.Map()
                    compr_47_: D8
                    for compr_47_ in (d_1652_v33_).Elements:
                        d_1655_v31_: D8 = compr_47_
                        if (d_1655_v31_) in (d_1652_v33_):
                            coll47_[d_1655_v31_] = len(_dafny.Set({(d_1626_v17_).cardinality}))
                    return _dafny.Map(coll47_)
                nw307_[int(0)] = (_dafny.Map({D8_DC21(d_1649_v29_): len(d_1650_v30_)})) | (iife157_()
                )
                nw307_[int(1)] = (d_1653_v34_).set(D8_DC21((d_1649_v29_).set(default__.safeIndex(d_1643_cf9_, len(d_1649_v29_)), p2)), len(d_1650_v30_))
                nw307_[int(2)] = d_1653_v34_
                nw307_[int(3)] = d_1653_v34_
                nw307_[int(4)] = (_dafny.Map({d_1651_v32_: self.f30})) | (d_1653_v34_)
                d_1654_v35_ = nw307_
                index307_ = default__.safeIndex(948, (d_1654_v35_).length(0))
                (d_1654_v35_)[index307_] = _dafny.Map({d_1651_v32_: self.f30})
                index308_ = default__.safeIndex(948, (d_1654_v35_).length(0))
                rhs302_ = d_1642_cf10_
                rhs303_ = d_1648_v28_
                rhs304_ = d_1653_v34_
                lhs273_ = globalState
                lhs274_ = d_1654_v35_
                lhs275_ = default__.safeIndex(948, (d_1654_v35_).length(0))
                lhs273_.f22 = rhs302_
                d_1648_v28_ = rhs303_
                lhs274_[lhs275_] = rhs304_
                d_1643_cf9_ = (default__.safeDivisionInt(p1, self.f30)) * (((0) - (self.f30) if d_1642_cf10_ else len(d_1650_v30_)))
            elif source22_.is_DC6:
                d_1656___mcc_h4_ = source22_.cf13
                d_1657___mcc_h5_ = source22_.cf14
                d_1658___mcc_h6_ = source22_.cf15
                d_1659___mcc_h7_ = source22_.cf16
                d_1660_cf16_ = d_1659___mcc_h7_
                d_1661_cf15_ = d_1658___mcc_h6_
                d_1662_cf14_ = d_1657___mcc_h5_
                d_1663_cf13_ = d_1656___mcc_h4_
                d_1664_v36_: _dafny.Seq
                d_1664_v36_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rmnqe"))
                d_1665_v37_: _dafny.Map
                d_1665_v37_ = _dafny.Map({d_1660_cf16_: d_1664_v36_})
                d_1666_v38_: _dafny.Set
                d_1666_v38_ = _dafny.Set({True})
                index309_ = default__.safeIndex(647, (d_1631_v21_).length(0))
                (d_1631_v21_)[index309_] = (0) - (((_dafny.MultiSet([len(d_1665_v37_), d_1663_cf13_, 287])).cardinality if d_1660_cf16_ else len(d_1666_v38_)))
                d_1667_v39_: _dafny.Array
                nw308_ = _dafny.Array(None, 1)
                nw308_[int(0)] = _dafny.CodePoint('h')
                d_1667_v39_ = nw308_
                d_1668_v40_: str
                d_1668_v40_ = _dafny.CodePoint('t')
                index310_ = default__.safeIndex(112, (d_1667_v39_).length(0))
                (d_1667_v39_)[index310_] = d_1668_v40_
                d_1669_v41_: D8
                d_1669_v41_ = D8_DC23(d_1660_cf16_)
                d_1670_v42_: _dafny.Array
                nw309_ = _dafny.Array(None, 17)
                nw309_[int(0)] = not(p2)
                nw309_[int(1)] = p2
                nw309_[int(2)] = p2
                nw309_[int(3)] = p2
                nw309_[int(4)] = p2
                nw309_[int(5)] = p2
                nw309_[int(6)] = p2
                nw309_[int(7)] = default__.fm12(d_1660_cf16_, p2, globalState)
                nw309_[int(8)] = p2
                nw309_[int(9)] = d_1660_cf16_
                nw309_[int(10)] = d_1660_cf16_
                nw309_[int(11)] = p2
                nw309_[int(12)] = p2
                nw309_[int(13)] = p2
                nw309_[int(14)] = d_1660_cf16_
                nw309_[int(15)] = d_1660_cf16_
                nw309_[int(16)] = False
                d_1670_v42_ = nw309_
                d_1671_v43_: D4
                d_1671_v43_ = D4_DC13(d_1668_v40_)
                d_1672_v44_: D4
                d_1672_v44_ = D4_DC15(d_1671_v43_)
                d_1673_v45_: _dafny.Map
                d_1673_v45_ = _dafny.Map({d_1626_v17_: _dafny.SeqWithoutIsStrInference([d_1663_cf13_, self.f30])})
                d_1674_v46_: D1
                d_1674_v46_ = D1_DC5(d_1663_cf13_, d_1660_cf16_, d_1660_cf16_, d_1673_v45_)
                d_1675_v47_: _dafny.Map
                d_1675_v47_ = _dafny.Map({self.f30: d_1674_v46_})
                d_1676_v49_: D7
                def iife158_():
                    coll48_ = _dafny.Set()
                    compr_48_: int
                    for compr_48_ in (d_1675_v47_).keys.Elements:
                        d_1677_v48_: int = compr_48_
                        if (d_1677_v48_) in (d_1675_v47_):
                            coll48_ = coll48_.union(_dafny.Set([default__.safeModuloInt(d_1677_v48_, 330)]))
                    return _dafny.Set(coll48_)
                d_1676_v49_ = D7_DC20(p0, d_1670_v42_, d_1672_v44_, not(d_1660_cf16_), (p0)[default__.safeIndex(len(iife158_()
), len(p0))])
                index311_ = default__.safeIndex(647, (d_1631_v21_).length(0))
                index312_ = default__.safeIndex(112, (d_1667_v39_).length(0))
                rhs305_ = (0) - ((p0)[default__.safeIndex(d_1663_cf13_, len(p0))])
                rhs306_ = (d_1669_v41_).cf39
                rhs307_ = (d_1676_v49_).cf34
                rhs308_ = p1
                rhs309_ = (d_1668_v40_ if (default__.fm0(d_1663_cf13_, globalState)) <= (d_1663_cf13_) else d_1668_v40_)
                lhs276_ = d_1631_v21_
                lhs277_ = default__.safeIndex(647, (d_1631_v21_).length(0))
                lhs278_ = globalState
                lhs279_ = globalState
                lhs280_ = d_1667_v39_
                lhs281_ = default__.safeIndex(112, (d_1667_v39_).length(0))
                lhs276_[lhs277_] = rhs305_
                lhs278_.f22 = rhs306_
                lhs279_.f10 = rhs307_
                d_1663_cf13_ = rhs308_
                lhs280_[lhs281_] = rhs309_
                d_1678_v50_: _dafny.MultiSet
                d_1678_v50_ = _dafny.MultiSet([d_1664_v36_, d_1664_v36_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ybyitirvg"))])
                d_1678_v50_ = d_1678_v50_
                d_1679_v51_: C0
                nw310_ = C0()
                nw310_.ctor__(p2)
                d_1679_v51_ = nw310_
                index313_ = default__.safeIndex(631, (d_1670_v42_).length(0))
                (d_1670_v42_)[index313_] = p2
                index314_ = default__.safeIndex(941, (d_1670_v42_).length(0))
                (d_1670_v42_)[index314_] = (d_1661_cf15_) < ((0) - (len(_dafny.SeqWithoutIsStrInference([401 for d_1680_i2_ in range(default__.abs(919))]))))
                d_1681_v52_: _dafny.Seq
                d_1681_v52_ = _dafny.SeqWithoutIsStrInference([d_1666_v38_, d_1666_v38_])
                index315_ = default__.safeIndex(631, (d_1670_v42_).length(0))
                index316_ = default__.safeIndex(941, (d_1670_v42_).length(0))
                rhs310_ = (d_1679_v51_).fm13(d_1663_cf13_, globalState)
                rhs311_ = not(d_1660_cf16_)
                rhs312_ = not(not(not(p2)))
                rhs313_ = (d_1681_v52_) <= (d_1681_v52_)
                lhs282_ = globalState
                lhs283_ = d_1670_v42_
                lhs284_ = default__.safeIndex(631, (d_1670_v42_).length(0))
                lhs285_ = d_1670_v42_
                lhs286_ = default__.safeIndex(941, (d_1670_v42_).length(0))
                lhs287_ = globalState
                lhs282_.f10 = rhs310_
                lhs283_[lhs284_] = rhs311_
                lhs285_[lhs286_] = rhs312_
                lhs287_.f13 = rhs313_
            elif True:
                d_1682___mcc_h8_ = source22_.cf8
                d_1683_cf8_ = d_1682___mcc_h8_
                d_1684_v53_: _dafny.Seq
                d_1684_v53_ = _dafny.SeqWithoutIsStrInference([p2, p2, p2])
                d_1685_v54_: _dafny.Seq
                d_1685_v54_ = _dafny.SeqWithoutIsStrInference([d_1684_v53_])
                d_1686_v55_: D1
                d_1686_v55_ = D1_DC6(p1, d_1685_v54_, p1, p2)
                d_1687_v56_: _dafny.Map
                d_1687_v56_ = _dafny.Map({d_1686_v55_: d_1634_v22_})
                d_1687_v56_ = (d_1687_v56_).set(d_1686_v55_, default__.fm24(199, p2, globalState))
                d_1688_v57_: _dafny.Map
                d_1688_v57_ = _dafny.Map({default__.safeModuloInt(728, p1): not (True) or (p2)})
                d_1689_v58_: _dafny.Seq
                d_1689_v58_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nvdnrsmts"))
                d_1688_v57_ = (d_1688_v57_).set(p1, not((self).fm4(d_1689_v58_, self.f30, globalState)))
                d_1690_v59_: D8
                d_1690_v59_ = D8_DC21(d_1684_v53_)
                d_1691_v60_: _dafny.Map
                d_1691_v60_ = _dafny.Map({d_1690_v59_: d_1631_v21_})
                d_1692_v61_: _dafny.Map
                d_1692_v61_ = d_1691_v60_
                d_1693_v62_: str
                d_1693_v62_ = _dafny.CodePoint('u')
                rhs314_ = (d_1692_v61_)
                rhs315_ = d_1693_v62_
                rhs316_ = not (p2) or ((p1) != (self.f30))
                lhs288_ = globalState
                lhs289_ = globalState
                d_1691_v60_ = rhs314_
                lhs288_.f19 = rhs315_
                lhs289_.f7 = rhs316_
                d_1694_v63_: _dafny.Map
                d_1694_v63_ = _dafny.Map({self.f30: (d_1689_v58_) + (d_1689_v58_)})
                d_1694_v63_ = (d_1694_v63_).set((self.f30) - (p1), d_1689_v58_)
            (globalState).f10 = (p1) - (627)
            d_1695_v64_: _dafny.Seq
            d_1695_v64_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vvk"))
            d_1696_v65_: _dafny.Array
            nw311_ = _dafny.Array(_dafny.CodePoint('D'), 27)
            d_1696_v65_ = nw311_
            d_1697_v66_: _dafny.Map
            d_1697_v66_ = _dafny.Map({d_1695_v64_: d_1696_v65_})
            d_1697_v66_ = (_dafny.Map({d_1695_v64_: d_1696_v65_})) | ((_dafny.Map({d_1695_v64_: d_1696_v65_})) | (d_1697_v66_))
            d_1696_v65_ = d_1696_v65_
            d_1698_v67_: C0
            nw312_ = C0()
            nw312_.ctor__(not((p1) <= (self.f30)))
            d_1698_v67_ = nw312_
        elif True:
            d_1699_v68_: C3
            nw313_ = C3()
            nw313_.ctor__()
            d_1699_v68_ = nw313_
            d_1700_v69_: _dafny.Map
            d_1700_v69_ = _dafny.Map({p1: d_1699_v68_})
            rhs317_ = p2
            rhs318_ = d_1700_v69_
            lhs290_ = globalState
            lhs290_.f13 = rhs317_
            d_1700_v69_ = rhs318_
            (globalState).f7 = p2
            if p2:
                d_1701_v70_: str
                d_1701_v70_ = _dafny.CodePoint('y')
                (globalState).f19 = d_1701_v70_
                d_1702_v71_: _dafny.MultiSet
                d_1702_v71_ = _dafny.MultiSet([p2, not(p2)])
                d_1703_v72_: _dafny.Map
                d_1703_v72_ = _dafny.Map({p2: d_1631_v21_})
                d_1704_v73_: _dafny.Map
                d_1704_v73_ = _dafny.Map({d_1702_v71_: ((d_1703_v72_)[not(True)] if (not(True)) in (d_1703_v72_) else d_1631_v21_)})
                d_1704_v73_ = (d_1704_v73_).set(d_1702_v71_, d_1631_v21_)
                (globalState).f10 = ((0) - (default__.safeDivisionInt((_dafny.MultiSet([p2])).cardinality, p1))) + (self.f30)
                d_1705_v74_: _dafny.Seq
                d_1705_v74_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lcy"))
                d_1706_v75_: _dafny.Map
                d_1706_v75_ = _dafny.Map({p2: len(d_1705_v74_)})
                (self).f30 = ((0) - ((p1) - (len(d_1706_v75_)))) * (self.f30)
                d_1707_v76_: _dafny.Array
                def lambda83_(d_1708_i3_):
                    return True

                init43_ = lambda83_
                nw314_ = _dafny.Array(None, 7)
                for i0_43_ in range(nw314_.length(0)):
                    nw314_[i0_43_] = init43_(i0_43_)
                d_1707_v76_ = nw314_
                index317_ = default__.safeIndex(79, (d_1707_v76_).length(0))
                (d_1707_v76_)[index317_] = p2
                d_1709_v77_: _dafny.Seq
                d_1709_v77_ = _dafny.SeqWithoutIsStrInference([p2, False])
                d_1710_v78_: _dafny.Set
                d_1710_v78_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([p1]), p0, p0, p0})
                index318_ = default__.safeIndex(79, (d_1707_v76_).length(0))
                rhs319_ = p2
                rhs320_ = default__.fm22((_dafny.Set({p0, p0, _dafny.SeqWithoutIsStrInference([len(p0)])})).issubset(d_1710_v78_), globalState)
                rhs321_ = p2
                lhs291_ = d_1707_v76_
                lhs292_ = default__.safeIndex(79, (d_1707_v76_).length(0))
                lhs293_ = globalState
                lhs291_[lhs292_] = rhs319_
                d_1709_v77_ = rhs320_
                lhs293_.f13 = rhs321_
            elif True:
                d_1711_v79_: _dafny.Seq
                d_1711_v79_ = _dafny.SeqWithoutIsStrInference([default__.fm12(True, p2, globalState)])
                d_1712_v80_: _dafny.Seq
                d_1712_v80_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "shrji"))
                d_1713_v81_: C0
                nw315_ = C0()
                nw315_.ctor__((d_1711_v79_)[default__.safeIndex(len(d_1712_v80_), len(d_1711_v79_))])
                d_1713_v81_ = nw315_
                d_1714_v82_: _dafny.Seq
                d_1714_v82_ = _dafny.SeqWithoutIsStrInference([d_1713_v81_])
                d_1715_v83_: _dafny.Array
                nw316_ = _dafny.Array(None, 23)
                nw316_[int(0)] = d_1713_v81_
                nw316_[int(1)] = d_1713_v81_
                nw316_[int(2)] = d_1713_v81_
                nw316_[int(3)] = d_1713_v81_
                nw316_[int(4)] = d_1713_v81_
                nw316_[int(5)] = d_1713_v81_
                nw316_[int(6)] = (d_1714_v82_)[default__.safeIndex(self.f30, len(d_1714_v82_))]
                nw316_[int(7)] = d_1713_v81_
                nw316_[int(8)] = d_1713_v81_
                nw316_[int(9)] = d_1713_v81_
                nw316_[int(10)] = d_1713_v81_
                nw316_[int(11)] = d_1713_v81_
                nw316_[int(12)] = (d_1713_v81_ if p2 else d_1713_v81_)
                nw316_[int(13)] = d_1713_v81_
                nw316_[int(14)] = d_1713_v81_
                nw316_[int(15)] = d_1713_v81_
                nw316_[int(16)] = (d_1713_v81_ if (d_1713_v81_).f32 else d_1713_v81_)
                nw316_[int(17)] = d_1713_v81_
                nw316_[int(18)] = d_1713_v81_
                nw316_[int(19)] = d_1713_v81_
                nw316_[int(20)] = d_1713_v81_
                nw316_[int(21)] = d_1713_v81_
                nw316_[int(22)] = d_1713_v81_
                d_1715_v83_ = nw316_
                index319_ = default__.safeIndex(455, (d_1715_v83_).length(0))
                (d_1715_v83_)[index319_] = d_1713_v81_
                index320_ = default__.safeIndex(455, (d_1715_v83_).length(0))
                (d_1715_v83_)[index320_] = d_1713_v81_
                (globalState).f10 = len(d_1712_v80_)
                d_1716_v84_: _dafny.Map
                d_1716_v84_ = _dafny.Map({len(p0): _dafny.SeqWithoutIsStrInference([self.f30 for d_1717_i4_ in range(default__.abs(345))])})
                d_1718_v85_: _dafny.Array
                def lambda84_(d_1719_p2_):
                    def lambda85_(d_1720_i5_):
                        return not(d_1719_p2_)

                    return lambda85_

                init44_ = lambda84_(p2)
                nw317_ = _dafny.Array(None, 22)
                for i0_44_ in range(nw317_.length(0)):
                    nw317_[i0_44_] = init44_(i0_44_)
                d_1718_v85_ = nw317_
                d_1721_v86_: str
                d_1721_v86_ = _dafny.CodePoint('x')
                d_1722_v87_: D4
                d_1722_v87_ = D4_DC13(d_1721_v86_)
                d_1723_v88_: D4
                d_1723_v88_ = D4_DC15(d_1722_v87_)
                d_1724_v89_: D7
                d_1724_v89_ = D7_DC20(((d_1716_v84_)[897] if (897) in (d_1716_v84_) else p0), d_1718_v85_, d_1723_v88_, p2, (self).fm2(globalState))
                rhs322_ = self.f30
                rhs323_ = (d_1724_v89_ if (p1) <= (self.f30) else d_1724_v89_)
                lhs294_ = self
                lhs294_.f30 = rhs322_
                d_1724_v89_ = rhs323_
                (globalState).f10 = (0) - (self.f30)
                d_1725_v90_: _dafny.Map
                d_1725_v90_ = _dafny.Map({p1: p2})
                d_1726_v91_: _dafny.Map
                d_1726_v91_ = _dafny.Map({_dafny.CodePoint('l'): (d_1713_v81_).f32})
                d_1727_v92_: _dafny.Array
                nw318_ = _dafny.Array(D4.default()(), 11)
                d_1727_v92_ = nw318_
                d_1728_v93_: _dafny.Seq
                d_1728_v93_ = _dafny.SeqWithoutIsStrInference([d_1712_v80_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e"))])
                index321_ = default__.safeIndex(114, (d_1727_v92_).length(0))
                (d_1727_v92_)[index321_] = default__.fm41(len((d_1728_v93_).set(default__.safeIndex(self.f30, len(d_1728_v93_)), d_1712_v80_)), p2, default__.fm24(93, p2, globalState), d_1626_v17_, globalState)
                d_1729_v95_: D4
                d_1729_v95_ = D4_DC13(default__.fm11(p1, globalState))
                index322_ = default__.safeIndex(114, (d_1727_v92_).length(0))
                def iife159_():
                    coll49_ = _dafny.Map()
                    compr_49_: int
                    for compr_49_ in (d_1605_v0_).keys.Elements:
                        d_1730_v94_: int = compr_49_
                        if (d_1730_v94_) in (d_1605_v0_):
                            coll49_[default__.safeDivisionInt(d_1730_v94_, (d_1626_v17_).cardinality)] = (d_1713_v81_).f32
                    return _dafny.Map(coll49_)
                rhs324_ = (iife159_()
                ) | (d_1725_v90_)
                rhs325_ = d_1726_v91_
                rhs326_ = d_1729_v95_
                lhs295_ = d_1727_v92_
                lhs296_ = default__.safeIndex(114, (d_1727_v92_).length(0))
                d_1725_v90_ = rhs324_
                d_1726_v91_ = rhs325_
                lhs295_[lhs296_] = rhs326_
            if not (p2) or (False):
                d_1731_v96_: str
                d_1731_v96_ = _dafny.CodePoint('i')
                d_1732_v97_: _dafny.MultiSet
                d_1732_v97_ = _dafny.MultiSet([d_1731_v96_])
                rhs327_ = (_dafny.CodePoint('q') if p2 else d_1731_v96_)
                rhs328_ = p1
                rhs329_ = ((p0)[default__.safeIndex(((d_1732_v97_).set(d_1731_v96_, default__.abs(p1))).cardinality, len(p0))]) - (self.f30)
                rhs330_ = self.f30
                lhs297_ = globalState
                lhs298_ = globalState
                lhs299_ = self
                lhs300_ = self
                lhs297_.f19 = rhs327_
                lhs298_.f10 = rhs328_
                lhs299_.f30 = rhs329_
                lhs300_.f30 = rhs330_
                d_1733_v98_: _dafny.Set
                d_1733_v98_ = _dafny.Set({p2, p2, p2, p2})
                (globalState).f13 = not((d_1733_v98_) == (d_1733_v98_))
                (self).f30 = (default__.safeDivisionInt(self.f30, self.f30)) - (self.f30)
                (globalState).f13 = not (not(True)) or (True)
                (globalState).f10 = len((p0) + (p0))
            elif True:
                index323_ = default__.safeIndex(567, (d_1631_v21_).length(0))
                (d_1631_v21_)[index323_] = (0) - (self.f30)
                index324_ = default__.safeIndex(567, (d_1631_v21_).length(0))
                (d_1631_v21_)[index324_] = (p1) * (((d_1699_v68_).fm2(globalState)) * (p1))
                (globalState).f10 = (p1 if (d_1629_v19_) != (d_1629_v19_) else self.f30)
                d_1734_v99_: _dafny.Seq
                d_1734_v99_ = _dafny.SeqWithoutIsStrInference([p2])
                d_1735_v100_: _dafny.MultiSet
                d_1735_v100_ = _dafny.MultiSet([d_1734_v99_, d_1734_v99_])
                d_1736_v101_: _dafny.MultiSet
                d_1736_v101_ = _dafny.MultiSet([p2])
                (globalState).f10 = ((d_1735_v100_).cardinality if not(p2) else (0) - (((d_1631_v21_)[default__.safeIndex(567, (d_1631_v21_).length(0))]) + ((_dafny.MultiSet([(d_1631_v21_)[default__.safeIndex(567, (d_1631_v21_).length(0))], (d_1736_v101_).cardinality, self.f30])).cardinality)))
                d_1737_v102_: _dafny.Map
                d_1737_v102_ = _dafny.Map({p2: (d_1631_v21_)[default__.safeIndex(567, (d_1631_v21_).length(0))]})
                d_1738_v103_: _dafny.Seq
                d_1738_v103_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rbconal"))
                d_1737_v102_ = (d_1737_v102_).set(p2, len(d_1738_v103_))
                (globalState).f13 = p2
            d_1739_v104_: _dafny.Seq
            d_1739_v104_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "llhcw"))
            d_1740_v105_: D0
            d_1740_v105_ = D0_DC2((((d_1626_v17_)[714] if (714) in (d_1626_v17_) else self.f30) if not(p2) else (0) - (len(d_1739_v104_))), not(p2))
            d_1740_v105_ = D0_DC2((0) - (default__.safeDivisionInt(self.f30, (0) - (p1))), p2)
        rhs331_ = (self.f30 if (self.f30) < (len(_dafny.SeqWithoutIsStrInference([p2]))) else self.f30)
        rhs332_ = p2
        rhs333_ = ((p0)[default__.safeIndex(len(p0), len(p0))] if p2 else len(_dafny.Set({p2})))
        lhs301_ = globalState
        lhs302_ = self
        lhs301_.f10 = rhs331_
        r1 = rhs332_
        lhs302_.f30 = rhs333_
        if p2:
            d_1741_v106_: _dafny.Array
            def lambda86_(d_1742_p2_):
                def lambda87_(d_1743_i6_):
                    return d_1742_p2_

                return lambda87_

            init45_ = lambda86_(p2)
            nw319_ = _dafny.Array(None, 10)
            for i0_45_ in range(nw319_.length(0)):
                nw319_[i0_45_] = init45_(i0_45_)
            d_1741_v106_ = nw319_
            index325_ = default__.safeIndex(119, (d_1741_v106_).length(0))
            (d_1741_v106_)[index325_] = p2
            d_1744_v107_: _dafny.Array
            def lambda88_(d_1745_i7_):
                return _dafny.CodePoint('x')

            init46_ = lambda88_
            nw320_ = _dafny.Array(None, 1)
            for i0_46_ in range(nw320_.length(0)):
                nw320_[i0_46_] = init46_(i0_46_)
            d_1744_v107_ = nw320_
            d_1746_v108_: _dafny.Array
            d_1746_v108_ = d_1744_v107_
            d_1747_v109_: _dafny.Seq
            d_1747_v109_ = _dafny.SeqWithoutIsStrInference([d_1744_v107_])
            d_1748_v110_: _dafny.Array
            nw321_ = _dafny.Array(None, 17)
            nw321_[int(0)] = d_1744_v107_
            nw321_[int(1)] = d_1744_v107_
            nw321_[int(2)] = d_1744_v107_
            nw321_[int(3)] = d_1744_v107_
            nw321_[int(4)] = d_1744_v107_
            nw321_[int(5)] = d_1744_v107_
            nw321_[int(6)] = d_1744_v107_
            nw321_[int(7)] = d_1744_v107_
            nw321_[int(8)] = d_1744_v107_
            nw321_[int(9)] = d_1744_v107_
            nw321_[int(10)] = d_1744_v107_
            nw321_[int(11)] = d_1744_v107_
            nw321_[int(12)] = d_1744_v107_
            nw321_[int(13)] = (d_1746_v108_)
            nw321_[int(14)] = d_1744_v107_
            nw321_[int(15)] = d_1744_v107_
            nw321_[int(16)] = (d_1747_v109_)[default__.safeIndex(self.f30, len(d_1747_v109_))]
            d_1748_v110_ = nw321_
            d_1749_v111_: _dafny.Seq
            d_1749_v111_ = _dafny.SeqWithoutIsStrInference([d_1748_v110_])
            index326_ = default__.safeIndex(119, (d_1741_v106_).length(0))
            rhs334_ = p2
            rhs335_ = p2
            rhs336_ = (d_1749_v111_)[default__.safeIndex(self.f30, len(d_1749_v111_))]
            lhs303_ = d_1741_v106_
            lhs304_ = default__.safeIndex(119, (d_1741_v106_).length(0))
            lhs305_ = globalState
            lhs306_ = globalState
            lhs303_[lhs304_] = rhs334_
            lhs305_.f22 = rhs335_
            lhs306_.f6 = rhs336_
            d_1750_v112_: _dafny.MultiSet
            d_1750_v112_ = _dafny.MultiSet([_dafny.Set({self.f30})])
            if (d_1750_v112_).ispropersubset(d_1750_v112_):
                d_1751_v113_: _dafny.Seq
                d_1751_v113_ = _dafny.SeqWithoutIsStrInference([(d_1626_v17_).issubset(_dafny.MultiSet([len(_dafny.Map({False: p1})), p1, p1, self.f30])), p2])
                d_1751_v113_ = default__.fm22(not((d_1741_v106_)[default__.safeIndex(119, (d_1741_v106_).length(0))]), globalState)
                d_1752_v114_: _dafny.Seq
                d_1752_v114_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bvnrgdjr"))
                d_1753_v115_: _dafny.Seq
                d_1753_v115_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([self.f30 for d_1754_i9_ in range(default__.abs(412))])])
                d_1755_v116_: str
                d_1755_v116_ = _dafny.CodePoint('m')
                d_1756_v117_: D0
                d_1756_v117_ = D0_DC0(d_1752_v114_)
                d_1757_v118_: _dafny.Array
                nw322_ = _dafny.Array(None, 11)
                nw322_[int(0)] = ((d_1752_v114_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_1758_i8_ in range(default__.abs(177))]))).set(default__.safeIndex(len(d_1753_v115_), len((d_1752_v114_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_1759_i8_ in range(default__.abs(177))])))), d_1755_v116_)
                nw322_[int(1)] = d_1752_v114_
                nw322_[int(2)] = d_1752_v114_
                nw322_[int(3)] = d_1752_v114_
                nw322_[int(4)] = d_1752_v114_
                nw322_[int(5)] = _dafny.SeqWithoutIsStrInference([d_1755_v116_ for d_1760_i10_ in range(default__.abs(-829))])
                nw322_[int(6)] = d_1752_v114_
                nw322_[int(7)] = (_dafny.SeqWithoutIsStrInference([d_1755_v116_ for d_1761_i11_ in range(default__.abs(528))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aiqhwkuhp")))
                nw322_[int(8)] = (d_1756_v117_).cf0
                nw322_[int(9)] = d_1752_v114_
                nw322_[int(10)] = (d_1752_v114_) + (d_1752_v114_)
                d_1757_v118_ = nw322_
                index327_ = default__.safeIndex(600, (d_1757_v118_).length(0))
                (d_1757_v118_)[index327_] = d_1752_v114_
                d_1762_v119_: _dafny.Map
                d_1762_v119_ = _dafny.Map({(d_1741_v106_)[default__.safeIndex(119, (d_1741_v106_).length(0))]: _dafny.SeqWithoutIsStrInference([d_1755_v116_ for d_1763_i12_ in range(default__.abs(-220))])})
                d_1764_v120_: _dafny.Map
                d_1764_v120_ = _dafny.Map({d_1741_v106_: p1})
                index328_ = default__.safeIndex(600, (d_1757_v118_).length(0))
                (d_1757_v118_)[index328_] = ((d_1762_v119_)[(p2) and ((d_1741_v106_)[default__.safeIndex(119, (d_1741_v106_).length(0))])] if ((p2) and ((d_1741_v106_)[default__.safeIndex(119, (d_1741_v106_).length(0))])) in (d_1762_v119_) else default__.fm23(d_1627_v18_, p1, len((d_1764_v120_).set(d_1741_v106_, p1)), globalState))
                index329_ = default__.safeIndex(600, (d_1757_v118_).length(0))
                (d_1757_v118_)[index329_] = (((d_1757_v118_)[default__.safeIndex(600, (d_1757_v118_).length(0))]) + (d_1752_v114_)).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d"))), len(((d_1757_v118_)[default__.safeIndex(600, (d_1757_v118_).length(0))]) + (d_1752_v114_))), d_1755_v116_)
                (globalState).f22 = (d_1741_v106_)[default__.safeIndex(119, (d_1741_v106_).length(0))]
                d_1765_v121_: _dafny.Seq
                d_1765_v121_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_1755_v116_: 467})])
                d_1766_v122_: D7
                d_1766_v122_ = D7_DC18(d_1765_v121_)
                pat_let_tv55_ = d_1765_v121_
                d_1767_v123_: _dafny.Seq
                def iife160_(_pat_let55_0):
                    def iife161_(d_1768_dt__update__tmp_h2_):
                        def iife162_(_pat_let56_0):
                            def iife163_(d_1769_dt__update_hcf29_h0_):
                                return D7_DC18(d_1769_dt__update_hcf29_h0_)
                            return iife163_(_pat_let56_0)
                        return iife162_(pat_let_tv55_)
                    return iife161_(_pat_let55_0)
                d_1767_v123_ = _dafny.SeqWithoutIsStrInference([d_1766_v122_, iife160_(d_1766_v122_)])
                (globalState).f13 = (self).fm4((d_1757_v118_)[default__.safeIndex(600, (d_1757_v118_).length(0))], len(d_1767_v123_), globalState)
            elif True:
                (globalState).f10 = default__.safeDivisionInt(default__.fm0(-529, globalState), self.f30)
                d_1770_v124_: _dafny.MultiSet
                d_1770_v124_ = _dafny.MultiSet([p2])
                d_1771_v125_: _dafny.Seq
                d_1771_v125_ = _dafny.SeqWithoutIsStrInference([(d_1741_v106_)[default__.safeIndex(119, (d_1741_v106_).length(0))]])
                d_1772_v126_: _dafny.Set
                d_1772_v126_ = _dafny.Set({d_1741_v106_, d_1741_v106_})
                rhs337_ = _dafny.MultiSet([p2, (_dafny.Set({len(d_1771_v125_)})) != (_dafny.Set({p1})), not((_dafny.Set({d_1741_v106_})) == (d_1772_v126_)), (_dafny.MultiSet([self.f30])).ispropersubset(_dafny.MultiSet([p1]))])
                rhs338_ = p2
                rhs339_ = p1
                lhs307_ = globalState
                lhs308_ = self
                d_1770_v124_ = rhs337_
                lhs307_.f7 = rhs338_
                lhs308_.f30 = rhs339_
                d_1773_v127_: _dafny.Array
                nw323_ = _dafny.Array(D1.default()(), 19)
                d_1773_v127_ = nw323_
                d_1774_v128_: _dafny.Map
                d_1774_v128_ = _dafny.Map({_dafny.MultiSet(d_1771_v125_): p2})
                d_1775_v129_: D1
                d_1775_v129_ = D1_DC4(d_1774_v128_)
                d_1776_v130_: _dafny.Map
                d_1776_v130_ = _dafny.Map({d_1775_v129_: d_1773_v127_})
                d_1777_v131_: _dafny.Map
                d_1777_v131_ = _dafny.Map({p2: d_1773_v127_})
                d_1778_v132_: _dafny.Array
                nw324_ = _dafny.Array(None, 27)
                nw324_[int(0)] = d_1773_v127_
                nw324_[int(1)] = d_1773_v127_
                nw324_[int(2)] = d_1773_v127_
                nw324_[int(3)] = d_1773_v127_
                nw324_[int(4)] = d_1773_v127_
                nw324_[int(5)] = d_1773_v127_
                nw324_[int(6)] = d_1773_v127_
                nw324_[int(7)] = d_1773_v127_
                nw324_[int(8)] = d_1773_v127_
                nw324_[int(9)] = d_1773_v127_
                nw324_[int(10)] = d_1773_v127_
                nw324_[int(11)] = d_1773_v127_
                nw324_[int(12)] = d_1773_v127_
                nw324_[int(13)] = d_1773_v127_
                nw324_[int(14)] = d_1773_v127_
                nw324_[int(15)] = d_1773_v127_
                nw324_[int(16)] = d_1773_v127_
                nw324_[int(17)] = d_1773_v127_
                nw324_[int(18)] = d_1773_v127_
                nw324_[int(19)] = d_1773_v127_
                nw324_[int(20)] = d_1773_v127_
                nw324_[int(21)] = ((d_1776_v130_)[d_1775_v129_] if (d_1775_v129_) in (d_1776_v130_) else d_1773_v127_)
                nw324_[int(22)] = d_1773_v127_
                nw324_[int(23)] = d_1773_v127_
                nw324_[int(24)] = d_1773_v127_
                nw324_[int(25)] = d_1773_v127_
                nw324_[int(26)] = ((d_1777_v131_)[p2] if (p2) in (d_1777_v131_) else d_1773_v127_)
                d_1778_v132_ = nw324_
                d_1779_v133_: C4
                nw325_ = C4()
                nw325_.ctor__(d_1778_v132_)
                d_1779_v133_ = nw325_
                d_1780_v134_: _dafny.Map
                d_1780_v134_ = _dafny.Map({p2: (d_1779_v133_ if p2 else d_1779_v133_)})
                d_1780_v134_ = (d_1780_v134_).set((d_1741_v106_)[default__.safeIndex(119, (d_1741_v106_).length(0))], d_1779_v133_)
                d_1781_v135_: _dafny.Set
                d_1781_v135_ = _dafny.Set({((0) - (self.f30)) + (self.f30), self.f30})
                (globalState).f24 = d_1781_v135_
                d_1782_v136_: C4
                nw326_ = C4()
                nw326_.ctor__((d_1779_v133_).f31)
                d_1782_v136_ = nw326_
            (globalState).f13 = ((0) - (self.f30)) >= (default__.safeDivisionInt(self.f30, p1))
            index330_ = default__.safeIndex(119, (d_1741_v106_).length(0))
            (d_1741_v106_)[index330_] = p2
            (globalState).f13 = p2
        elif True:
            d_1783_v137_: _dafny.Array
            nw327_ = _dafny.Array(_dafny.Array(None, 0), 19)
            d_1783_v137_ = nw327_
            index331_ = default__.safeIndex(77, ((self).f29).length(0))
            ((self).f29)[index331_] = d_1631_v21_
            index332_ = default__.safeIndex(77, ((self).f29).length(0))
            rhs340_ = d_1783_v137_
            rhs341_ = self.f30
            rhs342_ = d_1631_v21_
            rhs343_ = self.f30
            lhs309_ = globalState
            lhs310_ = (self).f29
            lhs311_ = default__.safeIndex(77, ((self).f29).length(0))
            lhs312_ = globalState
            d_1783_v137_ = rhs340_
            lhs309_.f10 = rhs341_
            lhs310_[lhs311_] = rhs342_
            lhs312_.f10 = rhs343_
            if p2:
                d_1784_v138_: _dafny.MultiSet
                d_1784_v138_ = _dafny.MultiSet([p2])
                (globalState).f10 = (default__.safeDivisionInt(p1, (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ekgpmyva")))))) + (((d_1784_v138_) | (_dafny.MultiSet([p2]))).cardinality)
                (globalState).f10 = default__.safeModuloInt(p1, default__.safeDivisionInt(self.f30, self.f30))
                (globalState).f13 = p2
                (globalState).f7 = p2
                (globalState).f22 = False
            elif True:
                (globalState).f13 = p2
                index333_ = default__.safeIndex(999, (d_1631_v21_).length(0))
                (d_1631_v21_)[index333_] = p1
                index334_ = default__.safeIndex(999, (d_1631_v21_).length(0))
                (d_1631_v21_)[index334_] = self.f30
                index335_ = default__.safeIndex(77, ((self).f29).length(0))
                ((self).f29)[index335_] = ((self).f29)[default__.safeIndex(77, ((self).f29).length(0))]
                d_1785_v139_: _dafny.Seq
                d_1785_v139_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lhnsen"))
                d_1785_v139_ = d_1785_v139_
                d_1786_v140_: _dafny.Array
                nw328_ = _dafny.Array(None, 10)
                nw328_[int(0)] = (p2) == (p2)
                nw328_[int(1)] = p2
                nw328_[int(2)] = p2
                nw328_[int(3)] = p2
                nw328_[int(4)] = p2
                nw328_[int(5)] = (p2 if p2 else p2)
                nw328_[int(6)] = p2
                nw328_[int(7)] = p2
                nw328_[int(8)] = True
                nw328_[int(9)] = p2
                d_1786_v140_ = nw328_
                index336_ = default__.safeIndex(573, (d_1786_v140_).length(0))
                (d_1786_v140_)[index336_] = not(p2)
                index337_ = default__.safeIndex(573, (d_1786_v140_).length(0))
                (d_1786_v140_)[index337_] = p2
            d_1787_v141_: _dafny.Map
            d_1787_v141_ = _dafny.Map({(0) - (self.f30): p2})
            d_1788_v142_: _dafny.Map
            d_1788_v142_ = _dafny.Map({d_1626_v17_: p0})
            d_1789_v143_: D1
            d_1789_v143_ = D1_DC5(len(d_1787_v141_), True, p2, d_1788_v142_)
            d_1790_v144_: _dafny.Seq
            d_1790_v144_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ipg"))
            d_1791_v145_: _dafny.Array
            nw329_ = _dafny.Array(None, 27)
            nw329_[int(0)] = p2
            nw329_[int(1)] = p2
            nw329_[int(2)] = p2
            nw329_[int(3)] = p2
            nw329_[int(4)] = p2
            nw329_[int(5)] = p2
            nw329_[int(6)] = not(p2)
            nw329_[int(7)] = p2
            nw329_[int(8)] = (d_1789_v143_).cf11
            nw329_[int(9)] = p2
            nw329_[int(10)] = p2
            nw329_[int(11)] = p2
            nw329_[int(12)] = p2
            nw329_[int(13)] = p2
            nw329_[int(14)] = p2
            nw329_[int(15)] = not(True)
            nw329_[int(16)] = p2
            nw329_[int(17)] = True
            nw329_[int(18)] = p2
            nw329_[int(19)] = p2
            nw329_[int(20)] = p2
            nw329_[int(21)] = (self).fm4(d_1790_v144_, p1, globalState)
            nw329_[int(22)] = p2
            nw329_[int(23)] = True
            nw329_[int(24)] = p2
            nw329_[int(25)] = p2
            nw329_[int(26)] = p2
            d_1791_v145_ = nw329_
            d_1792_v146_: D4
            d_1792_v146_ = D4_DC14(p2)
            d_1793_v147_: D4
            d_1793_v147_ = D4_DC15(d_1792_v146_)
            d_1794_v148_: D7
            d_1794_v148_ = D7_DC20(p0, d_1791_v145_, D4_DC15(d_1792_v146_), (self.f30) not in (p0), p1)
            d_1795_v149_: _dafny.Seq
            d_1795_v149_ = _dafny.SeqWithoutIsStrInference([p0, _dafny.SeqWithoutIsStrInference([self.f30, (0) - (len(_dafny.SeqWithoutIsStrInference([(0) - (self.f30) for d_1796_i13_ in range(default__.abs(937))]))), self.f30]), p0, p0, p0])
            d_1797_v150_: D4
            d_1797_v150_ = D4_DC15(d_1792_v146_)
            d_1794_v148_ = (D7_DC20((p0).set(default__.safeIndex(len((d_1795_v149_)[default__.safeIndex(self.f30, len(d_1795_v149_))]), len(p0)), p1), d_1791_v145_, d_1797_v150_, p2, len(_dafny.Map({self.f30: (0) - (p1)}))) if p2 else d_1794_v148_)
            d_1798_v151_: C1
            nw330_ = C1()
            nw330_.ctor__()
            d_1798_v151_ = nw330_
            if ((0) - (len(_dafny.Map({self.f30: (self).fm4(d_1790_v144_, self.f30, globalState)})))) >= (len((p0).set(default__.safeIndex(self.f30, len(p0)), len(p0)))):
                d_1799_v152_: D0
                d_1799_v152_ = D0_DC2((0) - (p1), p2)
                d_1800_v153_: _dafny.Seq
                d_1800_v153_ = _dafny.SeqWithoutIsStrInference([d_1799_v152_])
                d_1801_v154_: _dafny.Seq
                d_1801_v154_ = _dafny.SeqWithoutIsStrInference([d_1800_v153_, d_1800_v153_, (d_1800_v153_).set(default__.safeIndex(self.f30, len(d_1800_v153_)), d_1799_v152_), _dafny.SeqWithoutIsStrInference([d_1799_v152_ for d_1802_i14_ in range(default__.abs(367))]), d_1800_v153_])
                d_1803_v155_: _dafny.Map
                d_1803_v155_ = _dafny.Map({p2: (d_1801_v154_).set(default__.safeIndex(len(d_1605_v0_), len(d_1801_v154_)), d_1800_v153_)})
                d_1801_v154_ = (d_1801_v154_ if p2 else ((d_1803_v155_)[False] if (False) in (d_1803_v155_) else d_1801_v154_))
                d_1804_v156_: C3
                nw331_ = C3()
                nw331_.ctor__()
                d_1804_v156_ = nw331_
                (globalState).f10 = p1
                d_1805_v157_: _dafny.Map
                d_1805_v157_ = _dafny.Map({(d_1798_v151_).fm21(p1, globalState): (len(d_1790_v144_)) + (self.f30)})
                (globalState).f10 = len(d_1805_v157_)
                d_1806_v158_: _dafny.Array
                nw332_ = _dafny.Array(D1.default()(), 22)
                d_1806_v158_ = nw332_
                d_1807_v159_: _dafny.Array
                d_1807_v159_ = d_1806_v158_
                d_1808_v160_: _dafny.Map
                d_1808_v160_ = _dafny.Map({p2: d_1806_v158_})
                d_1809_v161_: _dafny.Array
                nw333_ = _dafny.Array(None, 12)
                nw333_[int(0)] = d_1806_v158_
                nw333_[int(1)] = d_1806_v158_
                nw333_[int(2)] = d_1806_v158_
                nw333_[int(3)] = d_1806_v158_
                nw333_[int(4)] = d_1806_v158_
                nw333_[int(5)] = d_1806_v158_
                nw333_[int(6)] = d_1806_v158_
                nw333_[int(7)] = d_1806_v158_
                nw333_[int(8)] = (d_1807_v159_)
                nw333_[int(9)] = ((d_1808_v160_)[p2] if (p2) in (d_1808_v160_) else d_1806_v158_)
                nw333_[int(10)] = d_1806_v158_
                nw333_[int(11)] = d_1806_v158_
                d_1809_v161_ = nw333_
                d_1810_v162_: T1
                nw334_ = C4()
                nw334_.ctor__(d_1809_v161_)
                d_1810_v162_ = nw334_
                d_1810_v162_ = d_1810_v162_
            elif True:
                (globalState).f7 = p2
                d_1811_v163_: C1
                nw335_ = C1()
                nw335_.ctor__()
                d_1811_v163_ = nw335_
                (globalState).f22 = p2
                d_1812_v164_: _dafny.Seq
                d_1812_v164_ = _dafny.SeqWithoutIsStrInference([p2])
                d_1813_v165_: str
                d_1813_v165_ = _dafny.CodePoint('y')
                d_1626_v17_ = (default__.fm17(self.f30, (d_1812_v164_)[default__.safeIndex(p1, len(d_1812_v164_))], d_1813_v165_, globalState)).intersection(d_1626_v17_)
                d_1814_v166_: C0
                nw336_ = C0()
                nw336_.ctor__((p2 if True else p2))
                d_1814_v166_ = nw336_
        d_1815_i15_: int
        d_1815_i15_ = 0
        with _dafny.label("7"):
            while p2:
                with _dafny.c_label("7"):
                    if (d_1815_i15_) >= (100):
                        raise _dafny.Break("7")
                    d_1815_i15_ = (d_1815_i15_) + (1)
                    (globalState).f22 = ((d_1629_v19_)[p2] if (p2) in (d_1629_v19_) else False)
                    (globalState).f22 = p2
                    if p2:
                        (globalState).f10 = ((p0)[default__.safeIndex(p1, len(p0))]) - (self.f30)
                        (globalState).f10 = default__.safeModuloInt(p1, (0) - (self.f30))
                        rhs344_ = d_1631_v21_
                        rhs345_ = default__.fm0((p1) + (p1), globalState)
                        lhs313_ = globalState
                        d_1631_v21_ = rhs344_
                        lhs313_.f10 = rhs345_
                        d_1816_v167_: C3
                        nw337_ = C3()
                        nw337_.ctor__()
                        d_1816_v167_ = nw337_
                        d_1817_v168_: _dafny.Seq
                        d_1817_v168_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kdqjxd"))
                        d_1818_v169_: C2
                        nw338_ = C2()
                        nw338_.ctor__(d_1817_v168_)
                        d_1818_v169_ = nw338_
                    elif True:
                        d_1631_v21_ = d_1631_v21_
                        d_1819_v170_: _dafny.Array
                        nw339_ = _dafny.Array(None, 10)
                        nw339_[int(0)] = d_1631_v21_
                        nw339_[int(1)] = (d_1631_v21_ if p2 else d_1631_v21_)
                        nw339_[int(2)] = d_1631_v21_
                        nw339_[int(3)] = d_1631_v21_
                        nw339_[int(4)] = d_1631_v21_
                        nw339_[int(5)] = (d_1631_v21_ if p2 else d_1631_v21_)
                        nw339_[int(6)] = d_1631_v21_
                        nw339_[int(7)] = d_1631_v21_
                        nw339_[int(8)] = d_1631_v21_
                        nw339_[int(9)] = d_1631_v21_
                        d_1819_v170_ = nw339_
                        d_1819_v170_ = d_1819_v170_
                        d_1820_v171_: _dafny.Array
                        nw340_ = _dafny.Array(_dafny.MultiSet({}), 1)
                        d_1820_v171_ = nw340_
                        d_1820_v171_ = d_1820_v171_
                        d_1821_v172_: str
                        d_1821_v172_ = _dafny.CodePoint('o')
                        d_1822_v173_: D8
                        d_1822_v173_ = D8_DC22(p1, not(p2), d_1821_v172_)
                        d_1823_v174_: _dafny.Set
                        d_1823_v174_ = _dafny.Set({d_1822_v173_})
                        d_1824_v175_: _dafny.Seq
                        d_1824_v175_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lfkqt"))
                        d_1823_v174_ = ((d_1823_v174_ if (self).fm4(d_1824_v175_, p1, globalState) else d_1823_v174_)) | (d_1823_v174_)
                        d_1605_v0_ = (_dafny.Map({(0) - (self.f30): p1})) | (d_1605_v0_)
                    d_1825_v176_: C1
                    nw341_ = C1()
                    nw341_.ctor__()
                    d_1825_v176_ = nw341_
                    d_1826_v177_: _dafny.Map
                    d_1826_v177_ = _dafny.Map({p2: d_1825_v176_})
                    if (p2) in ((d_1826_v177_) | (d_1826_v177_)):
                        d_1827_v178_: _dafny.Map
                        d_1827_v178_ = _dafny.Map({p2: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q"))})
                        d_1828_v179_: _dafny.Seq
                        d_1828_v179_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "djlw"))
                        d_1827_v178_ = (d_1827_v178_).set(p2, (d_1828_v179_) + (d_1828_v179_))
                        r1 = not(p2)
                        d_1829_v180_: _dafny.Array
                        nw342_ = _dafny.Array(None, 7)
                        nw342_[int(0)] = d_1828_v179_
                        nw342_[int(1)] = (d_1828_v179_) + (d_1828_v179_)
                        nw342_[int(2)] = d_1828_v179_
                        nw342_[int(3)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vxtcqx"))
                        nw342_[int(4)] = (d_1828_v179_) + (d_1828_v179_)
                        nw342_[int(5)] = (d_1828_v179_) + (d_1828_v179_)
                        nw342_[int(6)] = default__.fm23(D2_DC7(d_1605_v0_), self.f30, (0) - (self.f30), globalState)
                        d_1829_v180_ = nw342_
                        index338_ = default__.safeIndex(719, (d_1829_v180_).length(0))
                        (d_1829_v180_)[index338_] = d_1828_v179_
                        index339_ = default__.safeIndex(719, (d_1829_v180_).length(0))
                        (d_1829_v180_)[index339_] = d_1828_v179_
                        index340_ = default__.safeIndex(475, (d_1631_v21_).length(0))
                        (d_1631_v21_)[index340_] = default__.safeModuloInt(self.f30, -430)
                        d_1830_v181_: _dafny.Map
                        d_1830_v181_ = _dafny.Map({d_1825_v176_: d_1631_v21_})
                        d_1831_v182_: _dafny.Array
                        nw343_ = _dafny.Array(_dafny.MultiSet({}), 7)
                        d_1831_v182_ = nw343_
                        index341_ = default__.safeIndex(260, (d_1831_v182_).length(0))
                        (d_1831_v182_)[index341_] = _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_1828_v179_, d_1828_v179_, (d_1829_v180_)[default__.safeIndex(719, (d_1829_v180_).length(0))]]))
                        d_1832_v183_: C0
                        nw344_ = C0()
                        nw344_.ctor__(True)
                        d_1832_v183_ = nw344_
                        d_1833_v184_: _dafny.Seq
                        d_1833_v184_ = _dafny.SeqWithoutIsStrInference([(d_1832_v183_).f32, (d_1832_v183_).f32])
                        d_1834_v185_: D8
                        d_1834_v185_ = D8_DC21(d_1833_v184_)
                        d_1835_v186_: D8
                        d_1835_v186_ = D8_DC24(d_1834_v185_)
                        d_1836_v187_: _dafny.Map
                        d_1836_v187_ = _dafny.Map({d_1832_v183_: d_1835_v186_})
                        d_1837_v188_: _dafny.Map
                        d_1837_v188_ = _dafny.Map({p2: p1})
                        d_1838_v189_: _dafny.Map
                        d_1838_v189_ = _dafny.Map({(d_1836_v187_) | (d_1836_v187_): _dafny.MultiSet([default__.fm23(d_1627_v18_, (0) - (len(d_1837_v188_)), p1, globalState), (d_1829_v180_)[default__.safeIndex(719, (d_1829_v180_).length(0))], (d_1829_v180_)[default__.safeIndex(719, (d_1829_v180_).length(0))]])})
                        index342_ = default__.safeIndex(475, (d_1631_v21_).length(0))
                        index343_ = default__.safeIndex(260, (d_1831_v182_).length(0))
                        rhs346_ = p1
                        rhs347_ = (_dafny.Set({p1})).isdisjoint(default__.fm42((0) - (self.f30), _dafny.Set({p2, not(p2)}), p1, globalState))
                        rhs348_ = len((((d_1827_v178_)[p2] if (p2) in (d_1827_v178_) else default__.fm23(d_1627_v18_, p1, self.f30, globalState)) if p2 else (d_1829_v180_)[default__.safeIndex(719, (d_1829_v180_).length(0))]))
                        rhs349_ = d_1830_v181_
                        rhs350_ = ((d_1838_v189_)[(_dafny.Map({d_1832_v183_: d_1835_v186_})) | (d_1836_v187_)] if ((_dafny.Map({d_1832_v183_: d_1835_v186_})) | (d_1836_v187_)) in (d_1838_v189_) else (_dafny.MultiSet([(d_1829_v180_)[default__.safeIndex(719, (d_1829_v180_).length(0))], (d_1829_v180_)[default__.safeIndex(719, (d_1829_v180_).length(0))]])).intersection(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(d_1829_v180_)[default__.safeIndex(719, (d_1829_v180_).length(0))] for d_1839_i16_ in range(default__.abs(883))]))))
                        lhs314_ = globalState
                        lhs315_ = globalState
                        lhs316_ = d_1631_v21_
                        lhs317_ = default__.safeIndex(475, (d_1631_v21_).length(0))
                        lhs318_ = d_1831_v182_
                        lhs319_ = default__.safeIndex(260, (d_1831_v182_).length(0))
                        lhs314_.f10 = rhs346_
                        lhs315_.f7 = rhs347_
                        lhs316_[lhs317_] = rhs348_
                        d_1830_v181_ = rhs349_
                        lhs318_[lhs319_] = rhs350_
                        (globalState).f22 = p2
                    elif True:
                        (globalState).f13 = p2
                        d_1840_v190_: _dafny.Map
                        d_1840_v190_ = _dafny.Map({self.f30: p2})
                        d_1841_v191_: D2
                        d_1841_v191_ = D2_DC8(p2, not(((d_1840_v190_)[p1] if (p1) in (d_1840_v190_) else True)))
                        d_1841_v191_ = d_1841_v191_
                        d_1842_v192_: _dafny.Seq
                        d_1842_v192_ = _dafny.SeqWithoutIsStrInference([default__.fm12(p2, p2, globalState), p2])
                        d_1843_v193_: _dafny.Seq
                        d_1843_v193_ = _dafny.SeqWithoutIsStrInference([d_1605_v0_, _dafny.Map({(_dafny.MultiSet(d_1842_v192_)).cardinality: self.f30})])
                        d_1844_v194_: _dafny.Map
                        d_1844_v194_ = _dafny.Map({p1: d_1843_v193_})
                        d_1845_v195_: _dafny.Map
                        d_1845_v195_ = _dafny.Map({p2: self.f30})
                        d_1844_v194_ = (d_1844_v194_).set(((d_1845_v195_)[not(False)] if (not(False)) in (d_1845_v195_) else len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nqoaxevgx")))]))), d_1843_v193_)
                        (d_1825_v176_).m0(436, (p1) in (p0), self.f30, p2, globalState)
                        (globalState).f10 = -389
                    pass
            pass
        d_1846_v196_: _dafny.Seq
        d_1846_v196_ = _dafny.SeqWithoutIsStrInference([not(not(p2)), p2, p2])
        d_1847_v197_: _dafny.Seq
        d_1847_v197_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "htgli"))
        d_1848_v198_: D0
        d_1848_v198_ = D0_DC2(len(p0), p2)
        d_1849_v199_: _dafny.Array
        nw345_ = _dafny.Array(False, 5)
        d_1849_v199_ = nw345_
        d_1850_v200_: D0
        d_1850_v200_ = D0_DC1(len((d_1846_v196_) + (d_1846_v196_)), p1, (self).fm3(len(d_1847_v197_), p2, (self).fm2(globalState), d_1848_v198_, globalState), d_1849_v199_)
        r0 = d_1850_v200_
        r1 = p2
        return r0, r1

    @property
    def f29(self):
        return self._f29
