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
    def fm2(p0, p1, p2, globalState):
        return (_dafny.MultiSet([False])) == (_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([not(False)])) + (_dafny.SeqWithoutIsStrInference([False]))))

    @staticmethod
    def fm6(globalState):
        return (len(_dafny.Set({False, True}))) < (743)

    @staticmethod
    def fm9(p0, p1, p2, globalState):
        return _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gyeq"))) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gvla"))), (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([192, 920]))) == (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({606}))]))]))])

    @staticmethod
    def fm10(p0, p1, p2, p3, globalState):
        return _dafny.Set({(-374) + (980), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_0_i0_ in range(default__.abs(252))])), (len(_dafny.Set({True}))) * (99)})

    @staticmethod
    def fm11(p0, globalState):
        def lambda0_(source0_):
            if source0_.is_DC23:
                d_1___mcc_h0_ = source0_.cf35
                d_2___mcc_h1_ = source0_.cf36
                d_3___mcc_h2_ = source0_.cf37
                d_4_cf37_ = d_3___mcc_h2_
                d_5_cf36_ = d_2___mcc_h1_
                d_6_cf35_ = d_1___mcc_h0_
                return (_dafny.Map({d_5_cf36_: d_5_cf36_})) | (_dafny.Map({True: d_6_cf35_}))
            elif source0_.is_DC24:
                return (_dafny.Map({not(True): False})) | (_dafny.Map({not(False): False}))
            elif True:
                d_7___mcc_h3_ = source0_.cf34
                d_8_cf34_ = d_7___mcc_h3_
                return (D16_DC39(_dafny.Map({False: False}))).cf57

        return len(lambda0_(D8_DC24()))

    @staticmethod
    def fm12(p0, p1, p2, globalState):
        return D0_DC2((D8_DC23(True, True, len(_dafny.Set({800})))).cf36)

    @staticmethod
    def fm15(p0, p1, p2, globalState):
        return ((_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ftpspqpu")))})) | (_dafny.Map({True: 740}))) | ((_dafny.Map({True: -788})) | (_dafny.Map({False: 673})))

    @staticmethod
    def fm16(globalState):
        source1_ = D8_DC24()
        if source1_.is_DC23:
            d_9___mcc_h0_ = source1_.cf35
            d_10___mcc_h1_ = source1_.cf36
            d_11___mcc_h2_ = source1_.cf37
            d_12_cf37_ = d_11___mcc_h2_
            d_13_cf36_ = d_10___mcc_h1_
            d_14_cf35_ = d_9___mcc_h0_
            return _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([263 for d_15_i0_ in range(default__.abs(676))]))
        elif source1_.is_DC24:
            return _dafny.MultiSet([1])
        elif True:
            d_16___mcc_h3_ = source1_.cf34
            d_17_cf34_ = d_16___mcc_h3_
            return _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([905]))

    @staticmethod
    def fm17(p0, globalState):
        source2_ = D7_DC18(_dafny.Map({946: True}))
        if source2_.is_DC19:
            d_18___mcc_h0_ = source2_.cf27
            d_19___mcc_h1_ = source2_.cf28
            d_20___mcc_h2_ = source2_.cf29
            d_21___mcc_h3_ = source2_.cf30
            d_22_cf30_ = d_21___mcc_h3_
            d_23_cf29_ = d_20___mcc_h2_
            d_24_cf28_ = d_19___mcc_h1_
            d_25_cf27_ = d_18___mcc_h0_
            return D0_DC1(d_23_cf29_, _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "whpnil"))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xdqneud")))
        elif source2_.is_DC20:
            d_26___mcc_h4_ = source2_.cf31
            d_27___mcc_h5_ = source2_.cf32
            d_28_cf32_ = d_27___mcc_h5_
            d_29_cf31_ = d_26___mcc_h4_
            if d_28_cf32_:
                return D0_DC1(d_28_cf32_, _dafny.SeqWithoutIsStrInference([d_29_cf31_]), (D7_DC20(d_29_cf31_, d_28_cf32_)).cf31)
            elif True:
                return D0_DC1(True, _dafny.SeqWithoutIsStrInference([d_29_cf31_, d_29_cf31_, d_29_cf31_, d_29_cf31_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_30_i0_ in range(default__.abs(-869))])]), d_29_cf31_)
        elif source2_.is_DC18:
            d_31___mcc_h6_ = source2_.cf26
            d_32_cf26_ = d_31___mcc_h6_
            return D0_DC1(True, _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_33_i1_ in range(default__.abs(816))])]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "narkygd")))
        elif True:
            d_34___mcc_h7_ = source2_.cf33
            d_35_cf33_ = d_34___mcc_h7_
            return D0_DC1(not(False), _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u")) for d_36_i2_ in range(default__.abs(558))]), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_37_i3_ in range(default__.abs(863))]))

    @staticmethod
    def fm18(p0, p1, globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([D0_DC0(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hwq")))) for d_38_i0_ in range(default__.abs(256))]), _dafny.SeqWithoutIsStrInference([D0_DC0(694)]), _dafny.SeqWithoutIsStrInference([D0_DC0(len(_dafny.Map({True: True}))) for d_39_i1_ in range(default__.abs(559))])]) if not(False) else _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([D0_DC0(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dnmlywcha")))) for d_40_i2_ in range(default__.abs(865))])]))) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([D0_DC0(len(_dafny.Set({False, False}))), D0_DC0(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bgl"))))]) for d_41_i3_ in range(default__.abs(-414))]))

    @staticmethod
    def fm19(p0, p1, globalState):
        return not (False) or (not(not(True)))

    @staticmethod
    def fm20(p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_42_i0_ in range(default__.abs(583))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h")))

    @staticmethod
    def fm22(p0, p1, p2, globalState):
        return ((_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([D14_DC35(_dafny.MultiSet([_dafny.CodePoint('m')])), D14_DC35(_dafny.MultiSet([_dafny.CodePoint('p')]))])) + (_dafny.SeqWithoutIsStrInference([D14_DC35(_dafny.MultiSet([_dafny.CodePoint('a')]))])))) - (_dafny.MultiSet([D14_DC35(_dafny.MultiSet([_dafny.CodePoint('f')])), D14_DC35(_dafny.MultiSet([_dafny.CodePoint('k')])), D14_DC35(_dafny.MultiSet([_dafny.CodePoint('p')])), D14_DC35(_dafny.MultiSet([_dafny.CodePoint('e'), _dafny.CodePoint('t'), _dafny.CodePoint('j')])), D14_DC35(_dafny.MultiSet([_dafny.CodePoint('b'), _dafny.CodePoint('w')]))]))).cardinality

    @staticmethod
    def fm23(p0, p1, p2, p3, globalState):
        source3_ = D9_DC27(not(False), len(_dafny.SeqWithoutIsStrInference([-433, 600])), False)
        if source3_.is_DC26:
            d_43___mcc_h0_ = source3_.cf39
            d_44___mcc_h1_ = source3_.cf40
            d_45___mcc_h2_ = source3_.cf41
            d_46_cf41_ = d_45___mcc_h2_
            d_47_cf40_ = d_44___mcc_h1_
            d_48_cf39_ = d_43___mcc_h0_
            return d_47_cf40_
        elif source3_.is_DC27:
            d_49___mcc_h3_ = source3_.cf42
            d_50___mcc_h4_ = source3_.cf43
            d_51___mcc_h5_ = source3_.cf44
            d_52_cf44_ = d_51___mcc_h5_
            d_53_cf43_ = d_50___mcc_h4_
            d_54_cf42_ = d_49___mcc_h3_
            return _dafny.CodePoint('a')
        elif True:
            d_55___mcc_h6_ = source3_.cf38
            d_56_cf38_ = d_55___mcc_h6_
            return _dafny.CodePoint('e')

    @staticmethod
    def fm24(globalState):
        return ((_dafny.Map({_dafny.CodePoint('g'): _dafny.Map({True: False})}) if True else _dafny.Map({(D9_DC26(True, _dafny.CodePoint('a'), True)).cf40: _dafny.Map({not(False): False})}))) | (_dafny.Map({_dafny.CodePoint('f'): _dafny.Map({True: True})}))

    @staticmethod
    def fm25(p0, p1, globalState):
        return ((_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([-739, 362, len(_dafny.SeqWithoutIsStrInference([False]))]))])) | (_dafny.MultiSet([600, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([801 for d_57_i0_ in range(default__.abs(612))]))).cardinality, 708, (0) - (len(_dafny.SeqWithoutIsStrInference([not(True), not(True)]))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hbr")))]))).intersection(((D12_DC31(_dafny.MultiSet([108]))).cf48) | (_dafny.MultiSet([417])))

    @staticmethod
    def fm26(p0, p1, globalState):
        source4_ = D4_DC12(642)
        if source4_.is_DC11:
            return D1_DC5(not(False), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hknqq")), _dafny.SeqWithoutIsStrInference([not(False), False]))
        elif source4_.is_DC12:
            d_58___mcc_h0_ = source4_.cf18
            d_59_cf18_ = d_58___mcc_h0_
            return D1_DC5(True, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tdv")), _dafny.SeqWithoutIsStrInference([True]))
        elif True:
            d_60___mcc_h1_ = source4_.cf17
            d_61_cf17_ = d_60___mcc_h1_
            return D1_DC5(True, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "spg")), _dafny.SeqWithoutIsStrInference([False]))

    @staticmethod
    def fm27(p0, p1, p2, globalState):
        return _dafny.SeqWithoutIsStrInference([not(not (True) or (False))])

    @staticmethod
    def fm28(globalState):
        if (497) > (len(_dafny.Map({_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_62_i0_ in range(default__.abs(984))]))]): len(_dafny.Map({False: True}))}))):
            if not(False):
                return D0_DC3(D0_DC2(True))
            elif True:
                return D0_DC3(D0_DC3(D0_DC1(not(False), _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_63_i1_ in range(default__.abs(442))]), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_64_i2_ in range(default__.abs(958))])]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fwyar")))))
        elif True:
            return D0_DC3(D0_DC0(296))

    @staticmethod
    def fm29(p0, p1, p2, p3, globalState):
        return D2_DC7((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tfptdctot")))) == (855), _dafny.Map({False: D0_DC3(D0_DC0(272))}), (987) * (len(_dafny.SeqWithoutIsStrInference([False, True, True, not(False), True]))), not(not (True) or (False)))

    @staticmethod
    def fm30(globalState):
        return (_dafny.SeqWithoutIsStrInference([93])) + (_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Set({753, 220}))) for d_65_i0_ in range(default__.abs(25))]))

    @staticmethod
    def fm31(globalState):
        return _dafny.Map({len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_66_i0_ in range(default__.abs(-89))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mijsmoem")))): len(_dafny.Map({False: (D2_DC7(not(True), _dafny.Map({False: D0_DC3(D0_DC0(len(_dafny.Map({True: 430}))))}), 710, True)).cf13}))})

    @staticmethod
    def fm32(p0, p1, p2, p3, globalState):
        return ((_dafny.Map({990: not(True)})) | (_dafny.Map({len(_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([False]))})): False}))) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([False])): False}))

    @staticmethod
    def fm33(p0, p1, p2, globalState):
        return (_dafny.Set({not(False), True, False, True})) | ((_dafny.Set({False, not(False), not(True), True})) - ((_dafny.Set({False}))))

    @staticmethod
    def fm34(globalState):
        if (_dafny.Set({14, (0) - ((_dafny.MultiSet([False, False, not(False)])).cardinality)})).issubset(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "amok")))})):
            def iife0_():
                coll0_ = _dafny.Map()
                compr_0_: int
                for compr_0_ in _dafny.IntegerRange(299, 729):
                    d_67_v0_: int = compr_0_
                    if ((299) <= (d_67_v0_)) and ((d_67_v0_) < (729)):
                        coll0_[(d_67_v0_) * (len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({(_dafny.MultiSet([True])).cardinality: False})), len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mwcb")))}))])))] = _dafny.CodePoint('u')
                return _dafny.Map(coll0_)
            def iife1_():
                coll1_ = _dafny.Map()
                compr_1_: int
                for compr_1_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: False})) for d_70_i2_ in range(default__.abs(933))]))])).Elements:
                    d_71_v1_: int = compr_1_
                    if (d_71_v1_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: False})) for d_70_i2_ in range(default__.abs(933))]))])):
                        def iife2_():
                            coll2_ = _dafny.Set()
                            compr_2_: int
                            for compr_2_ in _dafny.IntegerRange(228, 681):
                                d_72_v2_: int = compr_2_
                                if ((228) <= (d_72_v2_)) and ((d_72_v2_) < (681)):
                                    coll2_ = coll2_.union(_dafny.Set([default__.safeModuloInt(d_72_v2_, 70)]))
                            return _dafny.Set(coll2_)
                        coll1_[default__.safeDivisionInt(d_71_v1_, 672)] = len(_dafny.SeqWithoutIsStrInference([228, len(iife2_()
                        )]))
                return _dafny.Map(coll1_)
            return (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([-406, len(_dafny.SeqWithoutIsStrInference([False]))]), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([not(True), True]))])])) - (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([len(iife0_()
            ), len(_dafny.Map({not(True): len(_dafny.Map({False: False}))}))]), _dafny.SeqWithoutIsStrInference([962 for d_68_i0_ in range(default__.abs(777))]), _dafny.SeqWithoutIsStrInference([159 for d_69_i1_ in range(default__.abs(589))]), _dafny.SeqWithoutIsStrInference([len(iife1_()
            )])]))
        elif True:
            return (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "juf"))), len(_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_73_i3_ in range(default__.abs(851))]))})), len(_dafny.Map({False: True})), (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([919, len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uise"))): 23}))]))).cardinality]), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([not(not(True))]))]), _dafny.SeqWithoutIsStrInference([739, 490, len(_dafny.SeqWithoutIsStrInference([738]))])])) | (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([not(False), False])), len(_dafny.SeqWithoutIsStrInference([False, False])), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_74_i4_ in range(default__.abs(487))]))])]))

    @staticmethod
    def fm35(p0, p1, p2, globalState):
        return D4_DC12(len((_dafny.SeqWithoutIsStrInference([-151, len(_dafny.Set({788}))])) + (_dafny.SeqWithoutIsStrInference([-655]))))

    @staticmethod
    def fm36(globalState):
        return _dafny.Map({_dafny.CodePoint('g'): 238})

    @staticmethod
    def fm37(p0, p1, p2, p3, globalState):
        return (_dafny.Map({True: True})) | ((_dafny.Map({False: True})) | (_dafny.Map({False: True})))

    @staticmethod
    def fm38(globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([-886 for d_75_i0_ in range(default__.abs(697))])])

    @staticmethod
    def fm39(p0, globalState):
        return ((_dafny.MultiSet([not(True), not(False)])) | (_dafny.MultiSet([False, True, False]))).intersection((_dafny.MultiSet([False])) - (_dafny.MultiSet([True, False])))

    @staticmethod
    def fm40(p0, p1, globalState):
        return D7_DC20((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_76_i0_ in range(default__.abs(-531))])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_77_i1_ in range(default__.abs(874))])), not((len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([False])): 979}))) < ((0) - ((0) - (len(_dafny.SeqWithoutIsStrInference([True, False])))))))

    @staticmethod
    def fm41(p0, p1, globalState):
        return D9_DC26((not(True)) not in (_dafny.Set({True, True})), _dafny.CodePoint('s'), not(not(not(True))))

    @staticmethod
    def fm42(p0, p1, p2, p3, globalState):
        return (_dafny.MultiSet([D0_DC1(False, _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jtdyx")) for d_78_i0_ in range(default__.abs(-783))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pj"))), D0_DC1(not(False), _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "woy")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dwuebhwk"))]), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_79_i1_ in range(default__.abs(321))])), D0_DC1(False, _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ue"))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wo")))])).intersection(_dafny.MultiSet([D0_DC1(not(not(True)), _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wg")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jwmnfd"))]), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_80_i2_ in range(default__.abs(475))])), D0_DC1(False, _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jmo"))]), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_81_i3_ in range(default__.abs(395))])), D0_DC1(False, _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "boc")) for d_82_i4_ in range(default__.abs(851))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cihhuxg")))]))

    @staticmethod
    def fm43(p0, globalState):
        return ((_dafny.Set({_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True, False])), _dafny.MultiSet([False, False]), _dafny.MultiSet([not(not(True))])})).intersection(_dafny.Set({_dafny.MultiSet([not(True)])}))).intersection((_dafny.Set({_dafny.MultiSet([True, True, True])})) | (_dafny.Set({_dafny.MultiSet([True, False, True, False]), _dafny.MultiSet([True])})))

    @staticmethod
    def fm44(p0, p1, p2, globalState):
        if (False if False else False):
            return _dafny.MultiSet([_dafny.Map({not(not(True)): not(not(True))}), _dafny.Map({True: False})])
        elif True:
            return (_dafny.MultiSet([_dafny.Map({True: False})])) | (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.Map({False: False}), _dafny.Map({False: not(False)}), _dafny.Map({False: False})])))

    @staticmethod
    def Main(noArgsParameter__):
        d_83_v0_: _dafny.Array
        nw0_ = _dafny.Array(False, 14)
        d_83_v0_ = nw0_
        d_84_v1_: _dafny.Array
        nw1_ = _dafny.Array(_dafny.Seq({}), 13)
        d_84_v1_ = nw1_
        d_85_v2_: _dafny.Seq
        d_85_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fno"))
        d_86_v3_: int
        d_86_v3_ = 636
        d_87_v4_: str
        d_87_v4_ = _dafny.CodePoint('a')
        d_88_v5_: _dafny.Set
        d_88_v5_ = _dafny.Set({(d_85_v2_)[default__.safeIndex(d_86_v3_, len(d_85_v2_))], d_87_v4_})
        d_89_v6_: _dafny.Map
        d_89_v6_ = _dafny.Map({d_88_v5_: d_86_v3_})
        d_90_v7_: _dafny.Seq
        d_90_v7_ = _dafny.SeqWithoutIsStrInference([d_89_v6_, d_89_v6_])
        d_91_v8_: bool
        d_91_v8_ = False
        d_92_v9_: _dafny.MultiSet
        d_92_v9_ = _dafny.MultiSet([d_91_v8_])
        d_93_v10_: _dafny.Map
        d_93_v10_ = _dafny.Map({d_91_v8_: d_91_v8_})
        d_94_v11_: _dafny.MultiSet
        d_94_v11_ = _dafny.MultiSet([_dafny.Map({d_91_v8_: d_91_v8_}), d_93_v10_, d_93_v10_])
        d_95_v12_: _dafny.Map
        d_95_v12_ = _dafny.Map({d_86_v3_: d_92_v9_})
        d_96_v13_: _dafny.Set
        d_96_v13_ = _dafny.Set({d_91_v8_})
        d_97_globalState_: GlobalState
        nw2_ = GlobalState()
        nw2_.ctor__(_dafny.CodePoint('b'), d_83_v0_, d_84_v1_, True, (d_90_v7_)[default__.safeIndex(((d_92_v9_)[True] if (True) in (d_92_v9_) else (d_94_v11_).cardinality), len(d_90_v7_))], True, -209, (d_95_v12_) | (d_95_v12_), (d_85_v2_) + (d_85_v2_), True, True, d_96_v13_, d_85_v2_, False)
        d_97_globalState_ = nw2_
        d_83_v0_ = d_83_v0_
        d_98_v14_: _dafny.Array
        nw3_ = _dafny.Array(int(0), 3)
        d_98_v14_ = nw3_
        index0_ = default__.safeIndex(484, (d_98_v14_).length(0))
        (d_98_v14_)[index0_] = d_86_v3_
        index1_ = default__.safeIndex(484, (d_98_v14_).length(0))
        (d_98_v14_)[index1_] = d_86_v3_
        index2_ = default__.safeIndex(454, (d_83_v0_).length(0))
        (d_83_v0_)[index2_] = ((d_98_v14_)[default__.safeIndex(484, (d_98_v14_).length(0))]) != ((0) - (len(_dafny.Map({d_85_v2_: d_85_v2_}))))
        d_99_v16_: _dafny.Seq
        d_99_v16_ = _dafny.SeqWithoutIsStrInference([(d_98_v14_)[default__.safeIndex(484, (d_98_v14_).length(0))], (d_98_v14_)[default__.safeIndex(484, (d_98_v14_).length(0))]])
        d_100_v17_: _dafny.Seq
        def iife3_():
            coll3_ = _dafny.Map()
            compr_3_: int
            for compr_3_ in (d_99_v16_).Elements:
                d_101_v15_: int = compr_3_
                if (d_101_v15_) in (d_99_v16_):
                    coll3_[(d_101_v15_) * ((d_98_v14_)[default__.safeIndex(484, (d_98_v14_).length(0))])] = d_91_v8_
            return _dafny.Map(coll3_)
        d_100_v17_ = _dafny.SeqWithoutIsStrInference([len(iife3_()
        )])
        d_102_v18_: _dafny.MultiSet
        d_102_v18_ = _dafny.MultiSet([d_99_v16_])
        d_103_v19_: _dafny.Map
        d_103_v19_ = _dafny.Map({964: d_86_v3_})
        d_104_v20_: _dafny.MultiSet
        d_104_v20_ = _dafny.MultiSet([(d_98_v14_)[default__.safeIndex(484, (d_98_v14_).length(0))], (d_98_v14_)[default__.safeIndex(484, (d_98_v14_).length(0))], (d_92_v9_).cardinality, ((d_103_v19_)[d_86_v3_] if (d_86_v3_) in (d_103_v19_) else (d_98_v14_)[default__.safeIndex(484, (d_98_v14_).length(0))])])
        index3_ = default__.safeIndex(484, (d_98_v14_).length(0))
        index4_ = default__.safeIndex(454, (d_83_v0_).length(0))
        index5_ = default__.safeIndex(484, (d_98_v14_).length(0))
        rhs0_ = default__.safeDivisionInt((0) - ((d_98_v14_)[default__.safeIndex(484, (d_98_v14_).length(0))]), d_86_v3_)
        rhs1_ = ((d_104_v20_) | (d_104_v20_)).issubset(_dafny.MultiSet([(d_98_v14_)[default__.safeIndex(484, (d_98_v14_).length(0))], ((d_102_v18_)[d_100_v17_] if (d_100_v17_) in (d_102_v18_) else d_86_v3_), d_86_v3_, -296]))
        rhs2_ = ((d_85_v2_)[default__.safeIndex((0) - (len(d_93_v10_)), len(d_85_v2_))]) in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qahncvdqm")))
        rhs3_ = ((145) + (482)) < (default__.safeModuloInt(d_86_v3_, (d_98_v14_)[default__.safeIndex(484, (d_98_v14_).length(0))]))
        rhs4_ = (d_98_v14_)[default__.safeIndex(484, (d_98_v14_).length(0))]
        lhs0_ = d_98_v14_
        lhs1_ = default__.safeIndex(484, (d_98_v14_).length(0))
        lhs2_ = d_97_globalState_
        lhs3_ = d_83_v0_
        lhs4_ = default__.safeIndex(454, (d_83_v0_).length(0))
        lhs5_ = d_98_v14_
        lhs6_ = default__.safeIndex(484, (d_98_v14_).length(0))
        lhs0_[lhs1_] = rhs0_
        d_91_v8_ = rhs1_
        lhs2_.f9 = rhs2_
        lhs3_[lhs4_] = rhs3_
        lhs5_[lhs6_] = rhs4_
        d_105_v21_: _dafny.Map
        d_105_v21_ = _dafny.Map({(_dafny.Map({d_91_v8_: d_91_v8_})).set((d_83_v0_)[default__.safeIndex(454, (d_83_v0_).length(0))], (d_83_v0_)[default__.safeIndex(454, (d_83_v0_).length(0))]): d_98_v14_})
        d_98_v14_ = ((d_105_v21_)[d_93_v10_] if (d_93_v10_) in (d_105_v21_) else d_98_v14_)
        d_86_v3_ = (0) - (-954)
        d_106_v22_: C7
        nw4_ = C7()
        nw4_.ctor__()
        d_106_v22_ = nw4_
        d_107_v23_: C5
        nw5_ = C5()
        nw5_.ctor__()
        d_107_v23_ = nw5_
        d_91_v8_ = (d_91_v8_ if (d_107_v23_) not in (_dafny.SeqWithoutIsStrInference([d_107_v23_])) else (d_83_v0_)[default__.safeIndex(454, (d_83_v0_).length(0))])
        hi0_ = len(default__.fm20((d_98_v14_)[default__.safeIndex(484, (d_98_v14_).length(0))], d_91_v8_, d_91_v8_, d_86_v3_, d_97_globalState_))
        for d_108_i0_ in range(((d_98_v14_)[default__.safeIndex(484, (d_98_v14_).length(0))]) * (d_86_v3_), hi0_):
            d_109_v24_: _dafny.Array
            nw6_ = _dafny.Array(None, 5)
            nw6_[int(0)] = d_87_v4_
            nw6_[int(1)] = d_87_v4_
            nw6_[int(2)] = d_87_v4_
            nw6_[int(3)] = d_87_v4_
            nw6_[int(4)] = d_87_v4_
            d_109_v24_ = nw6_
            d_110_v25_: _dafny.Map
            d_110_v25_ = _dafny.Map({d_109_v24_: d_86_v3_})
            d_110_v25_ = _dafny.Map({d_109_v24_: d_108_i0_})
            d_93_v10_ = (d_93_v10_).set(d_91_v8_, (d_83_v0_)[default__.safeIndex(454, (d_83_v0_).length(0))])
            d_111_v26_: C1
            nw7_ = C1()
            nw7_.ctor__()
            d_111_v26_ = nw7_
            d_85_v2_ = _dafny.SeqWithoutIsStrInference([d_87_v4_ for d_112_i1_ in range(default__.abs(607))])
        d_113_v27_: _dafny.Map
        d_113_v27_ = _dafny.Map({_dafny.CodePoint('p'): len(_dafny.SeqWithoutIsStrInference([d_87_v4_ for d_114_i2_ in range(default__.abs(-90))]))})
        d_115_v28_: D9
        d_115_v28_ = D9_DC25(d_113_v27_)
        d_116_v29_: _dafny.Map
        d_116_v29_ = _dafny.Map({d_86_v3_: d_115_v28_})
        rhs5_ = len(d_116_v29_)
        rhs6_ = d_85_v2_
        d_86_v3_ = rhs5_
        d_85_v2_ = rhs6_
        d_117_v30_: _dafny.Map
        d_117_v30_ = _dafny.Map({(d_98_v14_)[default__.safeIndex(484, (d_98_v14_).length(0))]: d_85_v2_})
        hi1_ = len(((d_117_v30_)[288] if (288) in (d_117_v30_) else d_85_v2_))
        for d_118_i3_ in range((d_98_v14_)[default__.safeIndex(484, (d_98_v14_).length(0))], hi1_):
            d_119_v31_: _dafny.MultiSet
            d_119_v31_ = _dafny.MultiSet([d_87_v4_, d_87_v4_])
            d_120_v32_: D14
            d_120_v32_ = D14_DC35(d_119_v31_)
            d_121_v33_: _dafny.Set
            d_121_v33_ = _dafny.Set({len(d_85_v2_)})
            d_122_v34_: _dafny.Seq
            d_122_v34_ = _dafny.SeqWithoutIsStrInference([default__.fm19(len(d_121_v33_), (d_83_v0_)[default__.safeIndex(454, (d_83_v0_).length(0))], d_97_globalState_)])
            if not ((d_119_v31_).issubset((d_120_v32_).cf53)) or ((d_122_v34_) <= (d_122_v34_)):
                (d_97_globalState_).f5 = (d_83_v0_)[default__.safeIndex(454, (d_83_v0_).length(0))]
                d_86_v3_ = default__.safeDivisionInt(default__.safeDivisionInt(85, (d_98_v14_)[default__.safeIndex(484, (d_98_v14_).length(0))]), default__.safeDivisionInt(len(d_85_v2_), (d_98_v14_)[default__.safeIndex(484, (d_98_v14_).length(0))]))
                d_123_v35_: _dafny.Set
                d_123_v35_ = _dafny.Set({d_85_v2_})
                d_124_v36_: T0
                nw8_ = C2()
                nw8_.ctor__()
                d_124_v36_ = nw8_
                d_125_v37_: D4
                d_125_v37_ = D4_DC12(default__.safeDivisionInt(-645, (d_92_v9_).cardinality))
                d_126_v38_: _dafny.Map
                d_126_v38_ = _dafny.Map({44: d_123_v35_})
                d_127_v39_: _dafny.Map
                d_127_v39_ = _dafny.Map({(d_83_v0_)[default__.safeIndex(454, (d_83_v0_).length(0))]: d_123_v35_})
                rhs7_ = (((d_126_v38_)[(d_98_v14_)[default__.safeIndex(484, (d_98_v14_).length(0))]] if ((d_98_v14_)[default__.safeIndex(484, (d_98_v14_).length(0))]) in (d_126_v38_) else d_123_v35_)).intersection((d_123_v35_) | (((d_127_v39_)[d_91_v8_] if (d_91_v8_) in (d_127_v39_) else d_123_v35_)))
                rhs8_ = d_124_v36_
                rhs9_ = d_125_v37_
                rhs10_ = d_91_v8_
                lhs7_ = d_97_globalState_
                d_123_v35_ = rhs7_
                d_124_v36_ = rhs8_
                d_125_v37_ = rhs9_
                lhs7_.f5 = rhs10_
                d_128_v40_: C1
                nw9_ = C1()
                nw9_.ctor__()
                d_128_v40_ = nw9_
                d_129_v41_: _dafny.Map
                d_129_v41_ = _dafny.Map({not((d_83_v0_)[default__.safeIndex(454, (d_83_v0_).length(0))]): (d_98_v14_)[default__.safeIndex(484, (d_98_v14_).length(0))]})
                d_130_v42_: _dafny.Map
                d_130_v42_ = _dafny.Map({(d_100_v17_)[default__.safeIndex((d_98_v14_)[default__.safeIndex(484, (d_98_v14_).length(0))], len(d_100_v17_))]: d_128_v40_})
                index6_ = default__.safeIndex(454, (d_83_v0_).length(0))
                rhs11_ = ((d_130_v42_)[d_118_i3_] if (d_118_i3_) in (d_130_v42_) else d_128_v40_)
                rhs12_ = (_dafny.MultiSet([d_86_v3_])).ispropersubset(_dafny.MultiSet([d_118_i3_, (d_99_v16_)[default__.safeIndex(d_118_i3_, len(d_99_v16_))]]))
                rhs13_ = (d_122_v34_) + (d_122_v34_)
                rhs14_ = d_103_v19_
                rhs15_ = (_dafny.Map({((d_93_v10_)[(d_83_v0_)[default__.safeIndex(454, (d_83_v0_).length(0))]] if ((d_83_v0_)[default__.safeIndex(454, (d_83_v0_).length(0))]) in (d_93_v10_) else d_91_v8_): len(d_85_v2_)})) | ((d_129_v41_) | (d_129_v41_))
                lhs8_ = d_83_v0_
                lhs9_ = default__.safeIndex(454, (d_83_v0_).length(0))
                d_128_v40_ = rhs11_
                lhs8_[lhs9_] = rhs12_
                d_122_v34_ = rhs13_
                d_103_v19_ = rhs14_
                d_129_v41_ = rhs15_
                d_131_v43_: C6
                nw10_ = C6()
                nw10_.ctor__()
                d_131_v43_ = nw10_
            elif True:
                index7_ = default__.safeIndex(484, (d_98_v14_).length(0))
                (d_98_v14_)[index7_] = (((d_98_v14_)[default__.safeIndex(484, (d_98_v14_).length(0))]) * (d_86_v3_)) + ((0) - (default__.safeDivisionInt(d_118_i3_, (d_98_v14_)[default__.safeIndex(484, (d_98_v14_).length(0))])))
                d_132_v44_: _dafny.Map
                d_132_v44_ = _dafny.Map({d_98_v14_: (d_83_v0_)[default__.safeIndex(454, (d_83_v0_).length(0))]})
                d_133_v45_: _dafny.Array
                nw11_ = _dafny.Array(_dafny.MultiSet({}), 7)
                d_133_v45_ = nw11_
                d_134_v46_: D1
                d_134_v46_ = D1_DC5((d_83_v0_)[default__.safeIndex(454, (d_83_v0_).length(0))], d_85_v2_, d_122_v34_)
                d_135_v47_: _dafny.MultiSet
                d_135_v47_ = _dafny.MultiSet([D1_DC5((d_83_v0_)[default__.safeIndex(454, (d_83_v0_).length(0))], _dafny.SeqWithoutIsStrInference([d_87_v4_ for d_136_i4_ in range(default__.abs(700))]), d_122_v34_), d_134_v46_])
                index8_ = default__.safeIndex(403, (d_133_v45_).length(0))
                (d_133_v45_)[index8_] = d_135_v47_
                index9_ = default__.safeIndex(403, (d_133_v45_).length(0))
                rhs16_ = d_132_v44_
                rhs17_ = _dafny.MultiSet([d_134_v46_, d_134_v46_, default__.fm26(True, d_91_v8_, d_97_globalState_)])
                rhs18_ = d_115_v28_
                lhs10_ = d_133_v45_
                lhs11_ = default__.safeIndex(403, (d_133_v45_).length(0))
                d_132_v44_ = rhs16_
                lhs10_[lhs11_] = rhs17_
                d_115_v28_ = rhs18_
                (d_97_globalState_).f5 = not(d_91_v8_)
                d_137_v48_: _dafny.Map
                d_137_v48_ = _dafny.Map({(d_98_v14_)[default__.safeIndex(484, (d_98_v14_).length(0))]: d_91_v8_})
                d_138_v49_: D7
                d_138_v49_ = D7_DC18(d_137_v48_)
                pat_let_tv0_ = d_91_v8_
                def iife4_(_pat_let0_0):
                    def iife5_(d_139_dt__update__tmp_h0_):
                        def iife6_(_pat_let1_0):
                            def iife7_(d_140_dt__update_hcf26_h0_):
                                return D7_DC18(d_140_dt__update_hcf26_h0_)
                            return iife7_(_pat_let1_0)
                        return iife6_(_dafny.Map({506: pat_let_tv0_}))
                    return iife5_(_pat_let0_0)
                d_138_v49_ = iife4_(d_138_v49_)
                (d_97_globalState_).f13 = default__.fm6(d_97_globalState_)
            d_141_v50_: T1
            nw12_ = C3()
            nw12_.ctor__(not((d_83_v0_)[default__.safeIndex(454, (d_83_v0_).length(0))]), d_118_i3_)
            d_141_v50_ = nw12_
            d_142_v51_: _dafny.Seq
            d_142_v51_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([default__.fm11(d_91_v8_, d_97_globalState_), len(d_122_v34_)]), d_104_v20_])
            (d_107_v23_).m3(d_86_v3_, d_141_v50_, (d_141_v50_).fm0(d_86_v3_, d_87_v4_, d_97_globalState_), (d_142_v51_)[default__.safeIndex((d_98_v14_)[default__.safeIndex(484, (d_98_v14_).length(0))], len(d_142_v51_))], d_97_globalState_)
            index10_ = default__.safeIndex(454, (d_83_v0_).length(0))
            (d_83_v0_)[index10_] = (d_83_v0_)[default__.safeIndex(454, (d_83_v0_).length(0))]
            index11_ = default__.safeIndex(484, (d_98_v14_).length(0))
            (d_98_v14_)[index11_] = (0) - ((0) - (d_86_v3_))
        d_143_v52_: _dafny.Map
        d_143_v52_ = _dafny.Map({d_86_v3_: (d_83_v0_)[default__.safeIndex(454, (d_83_v0_).length(0))]})
        d_144_v53_: D7
        d_144_v53_ = D7_DC18(d_143_v52_)
        d_145_v54_: _dafny.Seq
        d_145_v54_ = _dafny.SeqWithoutIsStrInference([(d_83_v0_)[default__.safeIndex(454, (d_83_v0_).length(0))]])
        d_146_v55_: _dafny.Seq
        d_146_v55_ = _dafny.SeqWithoutIsStrInference([d_85_v2_, d_85_v2_])
        d_147_v56_: D0
        d_147_v56_ = D0_DC1(d_91_v8_, d_146_v55_, d_85_v2_)
        d_148_v57_: D0
        d_148_v57_ = D0_DC3(D0_DC3(d_147_v56_))
        pat_let_tv1_ = d_144_v53_
        pat_let_tv2_ = d_144_v53_
        pat_let_tv3_ = d_143_v52_
        pat_let_tv4_ = d_143_v52_
        pat_let_tv5_ = d_144_v53_
        def lambda1_(source5_):
            if source5_.is_DC1:
                d_149___mcc_h0_ = source5_.cf1
                d_150___mcc_h1_ = source5_.cf2
                d_151___mcc_h2_ = source5_.cf3
                d_152_cf3_ = d_151___mcc_h2_
                d_153_cf2_ = d_150___mcc_h1_
                d_154_cf1_ = d_149___mcc_h0_
                return pat_let_tv1_
            elif source5_.is_DC2:
                d_155___mcc_h3_ = source5_.cf4
                d_156_cf4_ = d_155___mcc_h3_
                return pat_let_tv2_
            elif source5_.is_DC0:
                d_157___mcc_h4_ = source5_.cf0
                d_158_cf0_ = d_157___mcc_h4_
                d_159_dt__update__tmp_h1_ = D7_DC18(pat_let_tv3_)
                d_160_dt__update_hcf26_h1_ = pat_let_tv4_
                return D7_DC18(d_160_dt__update_hcf26_h1_)
            elif True:
                d_161___mcc_h5_ = source5_.cf5
                d_162_cf5_ = d_161___mcc_h5_
                return D7_DC18((pat_let_tv5_).cf26)

        rhs19_ = d_93_v10_
        rhs20_ = (d_98_v14_)[default__.safeIndex(484, (d_98_v14_).length(0))]
        rhs21_ = (d_145_v54_)[default__.safeIndex(d_86_v3_, len(d_145_v54_))]
        rhs22_ = d_85_v2_
        rhs23_ = lambda1_(d_148_v57_)
        lhs12_ = d_97_globalState_
        d_93_v10_ = rhs19_
        d_86_v3_ = rhs20_
        lhs12_.f13 = rhs21_
        d_85_v2_ = rhs22_
        d_144_v53_ = rhs23_
        d_163_v58_: _dafny.Map
        d_163_v58_ = _dafny.Map({d_91_v8_: d_86_v3_})
        d_164_v59_: D9
        d_164_v59_ = D9_DC26((d_83_v0_)[default__.safeIndex(454, (d_83_v0_).length(0))], d_87_v4_, d_91_v8_)
        d_165_v60_: _dafny.Array
        nw13_ = _dafny.Array(None, 29)
        nw13_[int(0)] = d_87_v4_
        nw13_[int(1)] = _dafny.CodePoint('y')
        nw13_[int(2)] = d_87_v4_
        nw13_[int(3)] = d_87_v4_
        nw13_[int(4)] = default__.fm23((d_83_v0_)[default__.safeIndex(454, (d_83_v0_).length(0))], _dafny.Map({d_87_v4_: d_93_v10_}), False, ((d_163_v58_)[d_91_v8_] if (d_91_v8_) in (d_163_v58_) else (d_98_v14_)[default__.safeIndex(484, (d_98_v14_).length(0))]), d_97_globalState_)
        nw13_[int(5)] = d_87_v4_
        nw13_[int(6)] = d_87_v4_
        nw13_[int(7)] = d_87_v4_
        nw13_[int(8)] = d_87_v4_
        nw13_[int(9)] = d_87_v4_
        nw13_[int(10)] = d_87_v4_
        nw13_[int(11)] = d_87_v4_
        nw13_[int(12)] = d_87_v4_
        nw13_[int(13)] = d_87_v4_
        nw13_[int(14)] = (d_164_v59_).cf40
        nw13_[int(15)] = d_87_v4_
        nw13_[int(16)] = d_87_v4_
        nw13_[int(17)] = d_87_v4_
        nw13_[int(18)] = d_87_v4_
        nw13_[int(19)] = d_87_v4_
        nw13_[int(20)] = d_87_v4_
        nw13_[int(21)] = d_87_v4_
        nw13_[int(22)] = d_87_v4_
        nw13_[int(23)] = d_87_v4_
        nw13_[int(24)] = d_87_v4_
        nw13_[int(25)] = d_87_v4_
        nw13_[int(26)] = d_87_v4_
        nw13_[int(27)] = d_87_v4_
        nw13_[int(28)] = d_87_v4_
        d_165_v60_ = nw13_
        index12_ = default__.safeIndex(873, (d_165_v60_).length(0))
        (d_165_v60_)[index12_] = d_87_v4_
        index13_ = default__.safeIndex(873, (d_165_v60_).length(0))
        (d_165_v60_)[index13_] = d_87_v4_
        d_166_v61_: T1
        nw14_ = C5()
        nw14_.ctor__()
        d_166_v61_ = nw14_
        d_167_v62_: _dafny.Map
        d_167_v62_ = _dafny.Map({False: d_166_v61_})
        d_168_v63_: _dafny.Set
        d_168_v63_ = _dafny.Set({d_167_v62_, d_167_v62_})
        d_169_i5_: int
        d_169_i5_ = 0
        with _dafny.label("0"):
            while ((d_168_v63_ if (d_83_v0_)[default__.safeIndex(454, (d_83_v0_).length(0))] else d_168_v63_)).isdisjoint(d_168_v63_):
                with _dafny.c_label("0"):
                    if (d_169_i5_) >= (100):
                        raise _dafny.Break("0")
                    d_169_i5_ = (d_169_i5_) + (1)
                    d_86_v3_ = ((0) - ((0) - (((d_98_v14_)[default__.safeIndex(484, (d_98_v14_).length(0))]) * (len(_dafny.SeqWithoutIsStrInference([357, d_86_v3_, d_86_v3_, (d_98_v14_)[default__.safeIndex(484, (d_98_v14_).length(0))]])))))) + ((D0_DC0((d_98_v14_)[default__.safeIndex(484, (d_98_v14_).length(0))])).cf0)
                    (d_97_globalState_).f13 = (d_91_v8_ if (d_83_v0_)[default__.safeIndex(454, (d_83_v0_).length(0))] else d_91_v8_)
                    d_170_v64_: _dafny.Map
                    d_170_v64_ = _dafny.Map({d_86_v3_: (d_98_v14_ if (d_83_v0_)[default__.safeIndex(454, (d_83_v0_).length(0))] else d_98_v14_)})
                    d_171_v65_: D11
                    d_171_v65_ = D11_DC30(False)
                    d_172_v66_: _dafny.Map
                    d_172_v66_ = _dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference([(d_165_v60_)[default__.safeIndex(873, (d_165_v60_).length(0))] for d_173_i6_ in range(default__.abs(465))]))): d_171_v65_})
                    d_174_v67_: D0
                    d_174_v67_ = D0_DC2((d_83_v0_)[default__.safeIndex(454, (d_83_v0_).length(0))])
                    index14_ = default__.safeIndex(484, (d_98_v14_).length(0))
                    index15_ = default__.safeIndex(484, (d_98_v14_).length(0))
                    rhs24_ = default__.safeDivisionInt((len(d_172_v66_) if d_91_v8_ else d_86_v3_), (d_98_v14_)[default__.safeIndex(484, (d_98_v14_).length(0))])
                    rhs25_ = _dafny.Map({(d_98_v14_)[default__.safeIndex(484, (d_98_v14_).length(0))]: d_98_v14_})
                    rhs26_ = d_166_v61_
                    rhs27_ = (d_107_v23_).fm5(d_174_v67_, d_97_globalState_)
                    rhs28_ = not(False)
                    lhs13_ = d_98_v14_
                    lhs14_ = default__.safeIndex(484, (d_98_v14_).length(0))
                    lhs15_ = d_98_v14_
                    lhs16_ = default__.safeIndex(484, (d_98_v14_).length(0))
                    lhs17_ = d_97_globalState_
                    lhs13_[lhs14_] = rhs24_
                    d_170_v64_ = rhs25_
                    d_166_v61_ = rhs26_
                    lhs15_[lhs16_] = rhs27_
                    lhs17_.f13 = rhs28_
                    d_175_v68_: C1
                    nw15_ = C1()
                    nw15_.ctor__()
                    d_175_v68_ = nw15_
                    d_176_v69_: _dafny.Set
                    d_176_v69_ = _dafny.Set({d_175_v68_, d_175_v68_})
                    (d_97_globalState_).f3 = ((d_176_v69_).intersection((_dafny.Set({d_175_v68_})))) == ((d_176_v69_ if default__.fm6(d_97_globalState_) else d_176_v69_))
                    pass
            pass
        if d_91_v8_:
            (d_97_globalState_).f13 = not(default__.fm2((_dafny.CodePoint('d') if d_91_v8_ else d_87_v4_), d_148_v57_, (d_98_v14_)[default__.safeIndex(484, (d_98_v14_).length(0))], d_97_globalState_))
            d_177_v70_: _dafny.Set
            d_177_v70_ = _dafny.Set({(d_98_v14_)[default__.safeIndex(484, (d_98_v14_).length(0))]})
            d_178_v71_: _dafny.Map
            d_178_v71_ = _dafny.Map({not(((d_83_v0_)[default__.safeIndex(454, (d_83_v0_).length(0))]) == (True)): (d_177_v70_) | (d_177_v70_)})
            d_179_v72_: D4
            d_179_v72_ = D4_DC10(d_177_v70_)
            d_178_v71_ = (d_178_v71_).set((d_83_v0_)[default__.safeIndex(454, (d_83_v0_).length(0))], (d_179_v72_).cf17)
            (d_97_globalState_).f13 = d_91_v8_
            d_180_v73_: _dafny.Map
            d_180_v73_ = _dafny.Map({d_98_v14_: d_85_v2_})
            d_181_v74_: _dafny.Array
            nw16_ = _dafny.Array(_dafny.Array(None, 0), 10)
            d_181_v74_ = nw16_
            d_182_v75_: _dafny.Array
            nw17_ = _dafny.Array(None, 26)
            d_182_v75_ = nw17_
            index16_ = default__.safeIndex(958, (d_181_v74_).length(0))
            (d_181_v74_)[index16_] = d_182_v75_
            d_183_v76_: _dafny.Seq
            d_183_v76_ = _dafny.SeqWithoutIsStrInference([d_83_v0_, d_83_v0_, d_83_v0_, d_83_v0_, d_83_v0_])
            index17_ = default__.safeIndex(958, (d_181_v74_).length(0))
            rhs29_ = ((d_180_v73_) | (d_180_v73_)) | (((d_180_v73_).set(d_98_v14_, d_85_v2_) if False else d_180_v73_))
            rhs30_ = False
            rhs31_ = d_182_v75_
            rhs32_ = ((d_91_v8_) in (d_145_v54_)) in ((d_145_v54_) + (d_145_v54_))
            rhs33_ = ((d_183_v76_) + (d_183_v76_)) + (d_183_v76_)
            lhs18_ = d_97_globalState_
            lhs19_ = d_181_v74_
            lhs20_ = default__.safeIndex(958, (d_181_v74_).length(0))
            lhs21_ = d_97_globalState_
            d_180_v73_ = rhs29_
            lhs18_.f9 = rhs30_
            lhs19_[lhs20_] = rhs31_
            lhs21_.f13 = rhs32_
            d_183_v76_ = rhs33_
            (d_107_v23_).m3((d_86_v3_) + (d_86_v3_), d_166_v61_, (d_98_v14_)[default__.safeIndex(484, (d_98_v14_).length(0))], (default__.fm16(d_97_globalState_)) - (d_104_v20_), d_97_globalState_)
        elif True:
            d_145_v54_ = _dafny.SeqWithoutIsStrInference([(d_83_v0_)[default__.safeIndex(454, (d_83_v0_).length(0))]])
            d_184_v77_: _dafny.Map
            d_184_v77_ = _dafny.Map({(d_98_v14_)[default__.safeIndex(484, (d_98_v14_).length(0))]: d_166_v61_})
            d_185_v78_: C2
            nw18_ = C2()
            nw18_.ctor__()
            d_185_v78_ = nw18_
            d_186_v79_: _dafny.Map
            d_186_v79_ = _dafny.Map({(d_98_v14_)[default__.safeIndex(484, (d_98_v14_).length(0))]: d_185_v78_})
            (d_107_v23_).m3(default__.safeModuloInt((d_98_v14_)[default__.safeIndex(484, (d_98_v14_).length(0))], d_86_v3_), ((d_184_v77_)[d_86_v3_] if (d_86_v3_) in (d_184_v77_) else d_166_v61_), (_dafny.MultiSet([d_186_v79_])).cardinality, _dafny.MultiSet([len(_dafny.Map({(d_98_v14_)[default__.safeIndex(484, (d_98_v14_).length(0))]: len(d_143_v52_)})), (d_98_v14_)[default__.safeIndex(484, (d_98_v14_).length(0))], len(d_85_v2_), (d_98_v14_)[default__.safeIndex(484, (d_98_v14_).length(0))]]), d_97_globalState_)
            d_187_v81_: _dafny.Seq
            def iife8_():
                coll4_ = _dafny.Map()
                compr_4_: int
                for compr_4_ in _dafny.IntegerRange(-331, -333):
                    d_188_v80_: int = compr_4_
                    if ((-331) <= (d_188_v80_)) and ((d_188_v80_) < (-333)):
                        coll4_[default__.safeModuloInt(d_188_v80_, (d_98_v14_)[default__.safeIndex(484, (d_98_v14_).length(0))])] = (d_83_v0_)[default__.safeIndex(454, (d_83_v0_).length(0))]
                return _dafny.Map(coll4_)
            d_187_v81_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_86_v3_: True}), d_143_v52_, iife8_()
            , d_143_v52_, _dafny.Map({(d_99_v16_)[default__.safeIndex((d_98_v14_)[default__.safeIndex(484, (d_98_v14_).length(0))], len(d_99_v16_))]: (d_83_v0_)[default__.safeIndex(454, (d_83_v0_).length(0))]})])
            d_86_v3_ = (0) - ((len(d_187_v81_) if not(d_91_v8_) else len(d_96_v13_)))
            d_189_v82_: _dafny.Map
            d_189_v82_ = _dafny.Map({d_86_v3_: _dafny.Map({d_91_v8_: d_91_v8_})})
            d_190_v83_: _dafny.Seq
            d_190_v83_ = _dafny.SeqWithoutIsStrInference([((d_189_v82_)[d_86_v3_] if (d_86_v3_) in (d_189_v82_) else d_93_v10_), d_93_v10_])
            d_191_v84_: _dafny.Seq
            d_191_v84_ = _dafny.SeqWithoutIsStrInference([d_190_v83_, d_190_v83_])
            d_192_v85_: _dafny.Array
            nw19_ = _dafny.Array(None, 16)
            nw19_[int(0)] = (d_94_v11_).set(_dafny.Map({(d_83_v0_)[default__.safeIndex(454, (d_83_v0_).length(0))]: (d_83_v0_)[default__.safeIndex(454, (d_83_v0_).length(0))]}), default__.abs(len((d_143_v52_).set((d_98_v14_)[default__.safeIndex(484, (d_98_v14_).length(0))], d_91_v8_))))
            nw19_[int(1)] = d_94_v11_
            nw19_[int(2)] = _dafny.MultiSet([d_93_v10_])
            nw19_[int(3)] = d_94_v11_
            nw19_[int(4)] = default__.fm44((d_98_v14_)[default__.safeIndex(484, (d_98_v14_).length(0))], d_86_v3_, d_91_v8_, d_97_globalState_)
            nw19_[int(5)] = (d_94_v11_) - (d_94_v11_)
            nw19_[int(6)] = d_94_v11_
            nw19_[int(7)] = d_94_v11_
            nw19_[int(8)] = d_94_v11_
            nw19_[int(9)] = _dafny.MultiSet((d_191_v84_)[default__.safeIndex((0) - ((0) - ((d_98_v14_)[default__.safeIndex(484, (d_98_v14_).length(0))])), len(d_191_v84_))])
            nw19_[int(10)] = _dafny.MultiSet([_dafny.Map({(d_83_v0_)[default__.safeIndex(454, (d_83_v0_).length(0))]: False}), d_93_v10_, d_93_v10_, d_93_v10_])
            nw19_[int(11)] = (d_94_v11_ if (d_83_v0_)[default__.safeIndex(454, (d_83_v0_).length(0))] else d_94_v11_)
            nw19_[int(12)] = d_94_v11_
            nw19_[int(13)] = d_94_v11_
            nw19_[int(14)] = (d_94_v11_).set(d_93_v10_, default__.abs(d_86_v3_))
            nw19_[int(15)] = d_94_v11_
            d_192_v85_ = nw19_
            index18_ = default__.safeIndex(542, (d_192_v85_).length(0))
            (d_192_v85_)[index18_] = (d_94_v11_) | (d_94_v11_)
            index19_ = default__.safeIndex(542, (d_192_v85_).length(0))
            index20_ = default__.safeIndex(454, (d_83_v0_).length(0))
            rhs34_ = default__.safeModuloInt((d_86_v3_ if (d_83_v0_)[default__.safeIndex(454, (d_83_v0_).length(0))] else d_86_v3_), d_86_v3_)
            rhs35_ = _dafny.MultiSet([_dafny.Map({(d_83_v0_)[default__.safeIndex(454, (d_83_v0_).length(0))]: d_91_v8_})])
            rhs36_ = not((d_83_v0_)[default__.safeIndex(454, (d_83_v0_).length(0))])
            rhs37_ = d_91_v8_
            lhs22_ = d_192_v85_
            lhs23_ = default__.safeIndex(542, (d_192_v85_).length(0))
            lhs24_ = d_83_v0_
            lhs25_ = default__.safeIndex(454, (d_83_v0_).length(0))
            lhs26_ = d_97_globalState_
            d_86_v3_ = rhs34_
            lhs22_[lhs23_] = rhs35_
            lhs24_[lhs25_] = rhs36_
            lhs26_.f9 = rhs37_
            d_86_v3_ = d_86_v3_
        d_193_v86_: T0
        nw20_ = C2()
        nw20_.ctor__()
        d_193_v86_ = nw20_
        d_193_v86_ = d_193_v86_
        d_194_v87_: bool
        d_195_v88_: int
        out0_: bool
        out1_: int
        out0_, out1_ = (d_106_v22_).m0(d_87_v4_, d_97_globalState_)
        d_194_v87_ = out0_
        d_195_v88_ = out1_
        _dafny.print(_dafny.string_of((d_83_v0_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_85_v2_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_86_v3_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_87_v4_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_88_v5_) == (_dafny.Set({_dafny.CodePoint('f'), _dafny.CodePoint('a')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_89_v6_) == (_dafny.Map({_dafny.Set({_dafny.CodePoint('f'), _dafny.CodePoint('a')}): 636}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_90_v7_) == (_dafny.SeqWithoutIsStrInference([_dafny.Map({_dafny.Set({_dafny.CodePoint('f'), _dafny.CodePoint('a')}): 636}), _dafny.Map({_dafny.Set({_dafny.CodePoint('f'), _dafny.CodePoint('a')}): 636})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_91_v8_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_92_v9_) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_93_v10_) == (_dafny.Map({False: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_94_v11_) == (_dafny.MultiSet([_dafny.Map({False: False}), _dafny.Map({False: False}), _dafny.Map({False: False})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_95_v12_) == (_dafny.Map({636: _dafny.MultiSet([False])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_96_v13_) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_97_globalState_).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_97_globalState_).f1)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_97_globalState_.f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_97_globalState_.f4) == (_dafny.Map({_dafny.Set({_dafny.CodePoint('f'), _dafny.CodePoint('a')}): 636}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_97_globalState_.f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_97_globalState_).f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_97_globalState_).f7) == (_dafny.Map({636: _dafny.MultiSet([False])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_97_globalState_).f8).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_97_globalState_.f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_97_globalState_).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_97_globalState_).f11) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_97_globalState_).f12).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_97_globalState_.f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_98_v14_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_99_v16_) == (_dafny.SeqWithoutIsStrInference([636, 636]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_100_v17_) == (_dafny.SeqWithoutIsStrInference([1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_102_v18_) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([636, 636])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_103_v19_) == (_dafny.Map({964: 636}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_104_v20_) == (_dafny.MultiSet([636, 636, 636, 1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_105_v21_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_113_v27_) == (_dafny.Map({_dafny.CodePoint('p'): 90}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_115_v28_).cf38) == (_dafny.Map({_dafny.CodePoint('p'): 90}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_116_v29_) == (_dafny.Map({954: D9_DC25(_dafny.Map({_dafny.CodePoint('p'): 90}))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_117_v30_) == (_dafny.Map({636: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fno"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_143_v52_) == (_dafny.Map({1: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_144_v53_).cf26) == (_dafny.Map({1: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_145_v54_) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_146_v55_) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fno")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fno"))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_147_v56_).cf1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_147_v56_).cf2) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fno")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fno"))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_147_v56_).cf3).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_148_v57_).cf5).cf5).cf1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((((d_148_v57_).cf5).cf5).cf2) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fno")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fno"))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((((d_148_v57_).cf5).cf5).cf3).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_163_v58_) == (_dafny.Map({False: 636}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_v59_).cf39))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_v59_).cf40))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_v59_).cf41))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_165_v60_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_165_v60_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_165_v60_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_165_v60_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_165_v60_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_165_v60_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_165_v60_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_165_v60_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_165_v60_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_165_v60_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_165_v60_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_165_v60_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_165_v60_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_165_v60_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_165_v60_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_165_v60_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_165_v60_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_165_v60_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_165_v60_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_165_v60_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_165_v60_)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_165_v60_)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_165_v60_)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_165_v60_)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_165_v60_)[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_165_v60_)[25]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_165_v60_)[26]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_165_v60_)[27]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_165_v60_)[28]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_167_v62_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_168_v63_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_169_i5_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_194_v87_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_195_v88_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(False, _dafny.Seq({}), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
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

class D0_DC1(D0, NamedTuple('DC1', [('cf1', Any), ('cf2', Any), ('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf1)}, {_dafny.string_of(self.cf2)}, {self.cf3.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf1 == __o.cf1 and self.cf2 == __o.cf2 and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC2(D0, NamedTuple('DC2', [('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2({_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2) and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC0(D0, NamedTuple('DC0', [('cf0', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC0({_dafny.string_of(self.cf0)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC0) and self.cf0 == __o.cf0
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC3(D0, NamedTuple('DC3', [('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC3({_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC3) and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()


class D1:
    @classmethod
    def default(cls, ):
        return lambda: D1_DC5(False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.Seq({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D1_DC5)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)

class D1_DC5(D1, NamedTuple('DC5', [('cf7', Any), ('cf8', Any), ('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf7)}, {self.cf8.VerbatimString(True)}, {_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf7 == __o.cf7 and self.cf8 == __o.cf8 and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC4(D1, NamedTuple('DC4', [('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC7(False, _dafny.Map({}), int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D2_DC8)

class D2_DC7(D2, NamedTuple('DC7', [('cf11', Any), ('cf12', Any), ('cf13', Any), ('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf11)}, {_dafny.string_of(self.cf12)}, {_dafny.string_of(self.cf13)}, {_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf11 == __o.cf11 and self.cf12 == __o.cf12 and self.cf13 == __o.cf13 and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC6(D2, NamedTuple('DC6', [('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC8(D2, NamedTuple('DC8', [('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC8({_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC8) and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Set({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D3_DC9)

class D3_DC9(D3, NamedTuple('DC9', [('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC11()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D4_DC11)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D4_DC12)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D4_DC10)

class D4_DC11(D4, NamedTuple('DC11', [])):
    def __dafnystr__(self) -> str:
        return f'D4.DC11'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC11)
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC12(D4, NamedTuple('DC12', [('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC12({_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC12) and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC10(D4, NamedTuple('DC10', [('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC10({_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC10) and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC14(False, False, _dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D5_DC14)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D5_DC13)

class D5_DC14(D5, NamedTuple('DC14', [('cf20', Any), ('cf21', Any), ('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC14({_dafny.string_of(self.cf20)}, {_dafny.string_of(self.cf21)}, {_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC14) and self.cf20 == __o.cf20 and self.cf21 == __o.cf21 and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC13(D5, NamedTuple('DC13', [('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC13({_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC13) and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC16(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D6_DC16)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D6_DC15)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D6_DC17)

class D6_DC16(D6, NamedTuple('DC16', [('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC16({_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC16) and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC15(D6, NamedTuple('DC15', [('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC15({_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC15) and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC17(D6, NamedTuple('DC17', [('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC17({_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC17) and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC19(D1.default()(), int(0), False, _dafny.Array(None, 0))
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
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D7_DC21)

class D7_DC19(D7, NamedTuple('DC19', [('cf27', Any), ('cf28', Any), ('cf29', Any), ('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC19({_dafny.string_of(self.cf27)}, {_dafny.string_of(self.cf28)}, {_dafny.string_of(self.cf29)}, {_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC19) and self.cf27 == __o.cf27 and self.cf28 == __o.cf28 and self.cf29 == __o.cf29 and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC20(D7, NamedTuple('DC20', [('cf31', Any), ('cf32', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC20({self.cf31.VerbatimString(True)}, {_dafny.string_of(self.cf32)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC20) and self.cf31 == __o.cf31 and self.cf32 == __o.cf32
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC18(D7, NamedTuple('DC18', [('cf26', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC18({_dafny.string_of(self.cf26)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC18) and self.cf26 == __o.cf26
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC21(D7, NamedTuple('DC21', [('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC21({_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC21) and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC23(False, False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D8_DC23)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D8_DC24)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D8_DC22)

class D8_DC23(D8, NamedTuple('DC23', [('cf35', Any), ('cf36', Any), ('cf37', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC23({_dafny.string_of(self.cf35)}, {_dafny.string_of(self.cf36)}, {_dafny.string_of(self.cf37)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC23) and self.cf35 == __o.cf35 and self.cf36 == __o.cf36 and self.cf37 == __o.cf37
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC24(D8, NamedTuple('DC24', [])):
    def __dafnystr__(self) -> str:
        return f'D8.DC24'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC24)
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC22(D8, NamedTuple('DC22', [('cf34', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC22({_dafny.string_of(self.cf34)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC22) and self.cf34 == __o.cf34
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC26(False, _dafny.CodePoint('D'), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D9_DC26)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D9_DC27)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D9_DC25)

class D9_DC26(D9, NamedTuple('DC26', [('cf39', Any), ('cf40', Any), ('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC26({_dafny.string_of(self.cf39)}, {_dafny.string_of(self.cf40)}, {_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC26) and self.cf39 == __o.cf39 and self.cf40 == __o.cf40 and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC27(D9, NamedTuple('DC27', [('cf42', Any), ('cf43', Any), ('cf44', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC27({_dafny.string_of(self.cf42)}, {_dafny.string_of(self.cf43)}, {_dafny.string_of(self.cf44)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC27) and self.cf42 == __o.cf42 and self.cf43 == __o.cf43 and self.cf44 == __o.cf44
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC25(D9, NamedTuple('DC25', [('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC25({_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC25) and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D10_DC28)

class D10_DC28(D10, NamedTuple('DC28', [('cf45', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC28({_dafny.string_of(self.cf45)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC28) and self.cf45 == __o.cf45
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC30(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D11_DC30)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D11_DC29)

class D11_DC30(D11, NamedTuple('DC30', [('cf47', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC30({_dafny.string_of(self.cf47)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC30) and self.cf47 == __o.cf47
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC29(D11, NamedTuple('DC29', [('cf46', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC29({_dafny.string_of(self.cf46)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC29) and self.cf46 == __o.cf46
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC32(False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D12_DC32)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D12_DC31)

class D12_DC32(D12, NamedTuple('DC32', [('cf49', Any), ('cf50', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC32({_dafny.string_of(self.cf49)}, {_dafny.string_of(self.cf50)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC32) and self.cf49 == __o.cf49 and self.cf50 == __o.cf50
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC31(D12, NamedTuple('DC31', [('cf48', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC31({_dafny.string_of(self.cf48)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC31) and self.cf48 == __o.cf48
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: D13_DC34(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D13_DC34)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D13_DC33)

class D13_DC34(D13, NamedTuple('DC34', [('cf52', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC34({_dafny.string_of(self.cf52)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC34) and self.cf52 == __o.cf52
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC33(D13, NamedTuple('DC33', [('cf51', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC33({_dafny.string_of(self.cf51)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC33) and self.cf51 == __o.cf51
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: D14_DC36(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D14_DC36)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D14_DC35)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D14_DC37)

class D14_DC36(D14, NamedTuple('DC36', [('cf54', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC36({_dafny.string_of(self.cf54)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC36) and self.cf54 == __o.cf54
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC35(D14, NamedTuple('DC35', [('cf53', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC35({_dafny.string_of(self.cf53)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC35) and self.cf53 == __o.cf53
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC37(D14, NamedTuple('DC37', [('cf55', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC37({_dafny.string_of(self.cf55)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC37) and self.cf55 == __o.cf55
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Set({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D15_DC38)

class D15_DC38(D15, NamedTuple('DC38', [('cf56', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC38({_dafny.string_of(self.cf56)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC38) and self.cf56 == __o.cf56
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: D16_DC40(False, False, int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D16_DC40)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D16_DC39)

class D16_DC40(D16, NamedTuple('DC40', [('cf58', Any), ('cf59', Any), ('cf60', Any), ('cf61', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC40({_dafny.string_of(self.cf58)}, {_dafny.string_of(self.cf59)}, {_dafny.string_of(self.cf60)}, {_dafny.string_of(self.cf61)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC40) and self.cf58 == __o.cf58 and self.cf59 == __o.cf59 and self.cf60 == __o.cf60 and self.cf61 == __o.cf61
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC39(D16, NamedTuple('DC39', [('cf57', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC39({_dafny.string_of(self.cf57)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC39) and self.cf57 == __o.cf57
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass

class T1:
    pass
    def m0(self, p0, globalState):
        pass


class GlobalState:
    def  __init__(self):
        self.f2: _dafny.Array = _dafny.Array(None, 0)
        self.f3: bool = False
        self.f4: _dafny.Map = _dafny.Map({})
        self.f5: bool = False
        self.f9: bool = False
        self.f13: bool = False
        self._f0: str = _dafny.CodePoint('D')
        self._f1: _dafny.Array = _dafny.Array(None, 0)
        self._f6: int = int(0)
        self._f7: _dafny.Map = _dafny.Map({})
        self._f8: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f10: bool = False
        self._f11: _dafny.Set = _dafny.Set({})
        self._f12: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13):
        (self)._f0 = f0
        (self)._f1 = f1
        (self).f2 = f2
        (self).f3 = f3
        (self).f4 = f4
        (self).f5 = f5
        (self)._f6 = f6
        (self)._f7 = f7
        (self)._f8 = f8
        (self).f9 = f9
        (self)._f10 = f10
        (self)._f11 = f11
        (self)._f12 = f12
        (self).f13 = f13

    @property
    def f0(self):
        return self._f0
    @property
    def f1(self):
        return self._f1
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
    def f10(self):
        return self._f10
    @property
    def f11(self):
        return self._f11
    @property
    def f12(self):
        return self._f12

class C0:
    def  __init__(self):
        self._f18: _dafny.MultiSet = _dafny.MultiSet({})
        self._f19: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f18, f19):
        (self)._f18 = f18
        (self)._f19 = f19

    def fm21(self, p0, p1, p2, p3, globalState):
        if (self).f19:
            return _dafny.SeqWithoutIsStrInference([not((self).f19)])
        elif True:
            return _dafny.SeqWithoutIsStrInference([False, (self).f19])

    @property
    def f18(self):
        return self._f18
    @property
    def f19(self):
        return self._f19

class C1:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    def ctor__(self):
        pass
        pass

    def m8(self, globalState):
        r0: int = int(0)
        d_196_v0_: bool
        d_196_v0_ = False
        d_197_v1_: int
        d_197_v1_ = -189
        d_198_v2_: D0
        d_198_v2_ = D0_DC2(d_196_v0_)
        d_199_v3_: _dafny.Seq
        d_199_v3_ = _dafny.SeqWithoutIsStrInference([d_198_v2_, D0_DC2(d_196_v0_), d_198_v2_, d_198_v2_, d_198_v2_])
        (globalState).f13 = (False if d_196_v0_ else ((0) - (d_197_v1_)) <= (len(d_199_v3_)))
        d_200_v4_: _dafny.Map
        d_200_v4_ = _dafny.Map({d_196_v0_: d_196_v0_})
        d_200_v4_ = (d_200_v4_).set((d_197_v1_) >= (-972), d_196_v0_)
        if not(d_196_v0_):
            d_201_v5_: _dafny.Array
            def lambda2_(d_202_i0_):
                return default__.safeDivisionInt(d_202_i0_, -167)

            init0_ = lambda2_
            nw21_ = _dafny.Array(None, 4)
            for i0_0_ in range(nw21_.length(0)):
                nw21_[i0_0_] = init0_(i0_0_)
            d_201_v5_ = nw21_
            d_203_v6_: _dafny.Map
            d_203_v6_ = _dafny.Map({d_197_v1_: d_201_v5_})
            d_204_v7_: _dafny.Map
            d_204_v7_ = _dafny.Map({default__.fm19(d_197_v1_, d_196_v0_, globalState): ((d_203_v6_)[d_197_v1_] if (d_197_v1_) in (d_203_v6_) else d_201_v5_)})
            d_205_v8_: D1
            d_205_v8_ = D1_DC4(d_204_v7_)
            d_204_v7_ = (d_205_v8_).cf6
            d_206_v9_: _dafny.Seq
            d_206_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a"))
            d_206_v9_ = d_206_v9_
            d_207_v10_: _dafny.Array
            def lambda3_(d_208_v0_, d_209_v4_):
                def lambda4_(d_210_i1_):
                    return (((d_209_v4_)[d_208_v0_] if (d_208_v0_) in (d_209_v4_) else True)) == (d_208_v0_)

                return lambda4_

            init1_ = lambda3_(d_196_v0_, d_200_v4_)
            nw22_ = _dafny.Array(None, 6)
            for i0_1_ in range(nw22_.length(0)):
                nw22_[i0_1_] = init1_(i0_1_)
            d_207_v10_ = nw22_
            d_211_v11_: _dafny.Seq
            d_211_v11_ = _dafny.SeqWithoutIsStrInference([d_196_v0_])
            d_212_v12_: _dafny.Seq
            d_212_v12_ = _dafny.SeqWithoutIsStrInference([d_196_v0_, d_196_v0_, (d_211_v11_)[default__.safeIndex(d_197_v1_, len(d_211_v11_))]])
            index21_ = default__.safeIndex(624, (d_207_v10_).length(0))
            (d_207_v10_)[index21_] = (d_212_v12_)[default__.safeIndex(d_197_v1_, len(d_212_v12_))]
            index22_ = default__.safeIndex(819, (d_201_v5_).length(0))
            (d_201_v5_)[index22_] = ((0) - (len(default__.fm20(len(d_206_v9_), d_196_v0_, d_196_v0_, d_197_v1_, globalState)))) - (d_197_v1_)
            index23_ = default__.safeIndex(624, (d_207_v10_).length(0))
            index24_ = default__.safeIndex(819, (d_201_v5_).length(0))
            rhs38_ = not(False)
            rhs39_ = ((d_197_v1_) + (d_197_v1_)) - ((d_197_v1_) + (d_197_v1_))
            rhs40_ = (((0) - (-89)) * (d_197_v1_)) != (default__.safeDivisionInt(d_197_v1_, d_197_v1_))
            lhs27_ = d_207_v10_
            lhs28_ = default__.safeIndex(624, (d_207_v10_).length(0))
            lhs29_ = d_201_v5_
            lhs30_ = default__.safeIndex(819, (d_201_v5_).length(0))
            lhs31_ = globalState
            lhs27_[lhs28_] = rhs38_
            lhs29_[lhs30_] = rhs39_
            lhs31_.f3 = rhs40_
            (globalState).f3 = True
            d_213_v13_: _dafny.MultiSet
            d_213_v13_ = _dafny.MultiSet([not(d_196_v0_), d_196_v0_, (d_207_v10_)[default__.safeIndex(624, (d_207_v10_).length(0))]])
            r0 = default__.safeDivisionInt(((d_213_v13_)[d_196_v0_] if (d_196_v0_) in (d_213_v13_) else (d_201_v5_)[default__.safeIndex(819, (d_201_v5_).length(0))]), (d_201_v5_)[default__.safeIndex(819, (d_201_v5_).length(0))])
        elif True:
            d_214_v14_: _dafny.Array
            nw23_ = _dafny.Array(int(0), 12)
            d_214_v14_ = nw23_
            d_215_v15_: _dafny.Seq
            d_215_v15_ = _dafny.SeqWithoutIsStrInference([d_197_v1_])
            index25_ = default__.safeIndex(755, (d_214_v14_).length(0))
            (d_214_v14_)[index25_] = len((d_215_v15_).set(default__.safeIndex(d_197_v1_, len(d_215_v15_)), d_197_v1_))
            index26_ = default__.safeIndex(755, (d_214_v14_).length(0))
            (d_214_v14_)[index26_] = (0) - (d_197_v1_)
            d_216_v16_: _dafny.Array
            nw24_ = _dafny.Array(_dafny.Array(None, 0), 20)
            d_216_v16_ = nw24_
            d_217_v17_: _dafny.Array
            nw25_ = _dafny.Array(_dafny.CodePoint('D'), 18)
            d_217_v17_ = nw25_
            d_218_v18_: _dafny.Seq
            d_218_v18_ = _dafny.SeqWithoutIsStrInference([d_217_v17_, d_217_v17_, d_217_v17_, d_217_v17_, d_217_v17_])
            d_219_v19_: _dafny.Array
            nw26_ = _dafny.Array(None, 28)
            nw26_[int(0)] = d_217_v17_
            nw26_[int(1)] = d_217_v17_
            nw26_[int(2)] = d_217_v17_
            nw26_[int(3)] = d_217_v17_
            nw26_[int(4)] = d_217_v17_
            nw26_[int(5)] = d_217_v17_
            nw26_[int(6)] = d_217_v17_
            nw26_[int(7)] = (d_218_v18_)[default__.safeIndex(-11, len(d_218_v18_))]
            nw26_[int(8)] = d_217_v17_
            nw26_[int(9)] = d_217_v17_
            nw26_[int(10)] = d_217_v17_
            nw26_[int(11)] = (d_217_v17_ if d_196_v0_ else d_217_v17_)
            nw26_[int(12)] = d_217_v17_
            nw26_[int(13)] = d_217_v17_
            nw26_[int(14)] = d_217_v17_
            nw26_[int(15)] = d_217_v17_
            nw26_[int(16)] = d_217_v17_
            nw26_[int(17)] = d_217_v17_
            nw26_[int(18)] = d_217_v17_
            nw26_[int(19)] = d_217_v17_
            nw26_[int(20)] = d_217_v17_
            nw26_[int(21)] = (d_217_v17_ if d_196_v0_ else d_217_v17_)
            nw26_[int(22)] = (d_217_v17_ if d_196_v0_ else d_217_v17_)
            nw26_[int(23)] = d_217_v17_
            nw26_[int(24)] = d_217_v17_
            nw26_[int(25)] = d_217_v17_
            nw26_[int(26)] = d_217_v17_
            nw26_[int(27)] = d_217_v17_
            d_219_v19_ = nw26_
            index27_ = default__.safeIndex(513, (d_219_v19_).length(0))
            (d_219_v19_)[index27_] = d_217_v17_
            index28_ = default__.safeIndex(513, (d_219_v19_).length(0))
            rhs41_ = d_216_v16_
            rhs42_ = d_217_v17_
            rhs43_ = d_196_v0_
            lhs32_ = d_219_v19_
            lhs33_ = default__.safeIndex(513, (d_219_v19_).length(0))
            d_216_v16_ = rhs41_
            lhs32_[lhs33_] = rhs42_
            d_196_v0_ = rhs43_
            d_220_v20_: _dafny.Seq
            d_220_v20_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ddwvkjgl"))
            if (d_220_v20_) != (d_220_v20_):
                d_221_v21_: _dafny.MultiSet
                d_221_v21_ = _dafny.MultiSet([d_197_v1_, d_197_v1_])
                d_222_v22_: _dafny.Map
                d_222_v22_ = _dafny.Map({(d_214_v14_)[default__.safeIndex(755, (d_214_v14_).length(0))]: (d_221_v21_).intersection(d_221_v21_)})
                d_222_v22_ = (d_222_v22_).set(((d_214_v14_)[default__.safeIndex(755, (d_214_v14_).length(0))]) + (d_197_v1_), (_dafny.MultiSet(d_215_v15_)).intersection((d_221_v21_).set((d_214_v14_)[default__.safeIndex(755, (d_214_v14_).length(0))], default__.abs(d_197_v1_))))
                d_223_v23_: _dafny.Seq
                d_223_v23_ = _dafny.SeqWithoutIsStrInference([d_220_v20_, d_220_v20_])
                d_224_v24_: _dafny.Seq
                d_224_v24_ = _dafny.SeqWithoutIsStrInference([(d_223_v23_)[default__.safeIndex(d_197_v1_, len(d_223_v23_))], d_220_v20_])
                d_225_v25_: _dafny.Seq
                d_225_v25_ = _dafny.SeqWithoutIsStrInference([d_196_v0_])
                d_226_v26_: D1
                d_226_v26_ = D1_DC5(d_196_v0_, d_220_v20_, d_225_v25_)
                d_227_v27_: _dafny.Map
                d_227_v27_ = _dafny.Map({(d_214_v14_)[default__.safeIndex(755, (d_214_v14_).length(0))]: (d_214_v14_)[default__.safeIndex(755, (d_214_v14_).length(0))]})
                d_228_v28_: _dafny.Map
                d_228_v28_ = _dafny.Map({d_200_v4_: d_196_v0_})
                d_229_v29_: _dafny.Set
                d_229_v29_ = _dafny.Set({d_196_v0_})
                d_230_v30_: _dafny.MultiSet
                d_230_v30_ = _dafny.MultiSet([d_229_v29_])
                d_231_v31_: _dafny.Array
                nw27_ = _dafny.Array(None, 27)
                nw27_[int(0)] = d_196_v0_
                nw27_[int(1)] = (False) and (d_196_v0_)
                nw27_[int(2)] = True
                nw27_[int(3)] = d_196_v0_
                nw27_[int(4)] = d_196_v0_
                nw27_[int(5)] = not(d_196_v0_)
                nw27_[int(6)] = d_196_v0_
                nw27_[int(7)] = (d_220_v20_) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rexhpd")))
                nw27_[int(8)] = d_196_v0_
                nw27_[int(9)] = d_196_v0_
                nw27_[int(10)] = not(d_196_v0_)
                nw27_[int(11)] = (d_197_v1_) > (len((d_223_v23_)[default__.safeIndex(((d_221_v21_)[(d_214_v14_)[default__.safeIndex(755, (d_214_v14_).length(0))]] if ((d_214_v14_)[default__.safeIndex(755, (d_214_v14_).length(0))]) in (d_221_v21_) else (d_214_v14_)[default__.safeIndex(755, (d_214_v14_).length(0))]), len(d_223_v23_))]))
                nw27_[int(12)] = not(d_196_v0_)
                nw27_[int(13)] = True
                nw27_[int(14)] = (d_226_v26_).cf7
                nw27_[int(15)] = d_196_v0_
                nw27_[int(16)] = ((d_227_v27_).set((d_214_v14_)[default__.safeIndex(755, (d_214_v14_).length(0))], (d_214_v14_)[default__.safeIndex(755, (d_214_v14_).length(0))])) == (d_227_v27_)
                nw27_[int(17)] = d_196_v0_
                nw27_[int(18)] = (True if ((d_228_v28_)[d_200_v4_] if (d_200_v4_) in (d_228_v28_) else d_196_v0_) else not(not(d_196_v0_)))
                nw27_[int(19)] = True
                nw27_[int(20)] = (d_221_v21_).ispropersubset(d_221_v21_)
                nw27_[int(21)] = d_196_v0_
                nw27_[int(22)] = d_196_v0_
                nw27_[int(23)] = (((d_230_v30_)[d_229_v29_] if (d_229_v29_) in (d_230_v30_) else (d_214_v14_)[default__.safeIndex(755, (d_214_v14_).length(0))])) == (d_197_v1_)
                nw27_[int(24)] = d_196_v0_
                nw27_[int(25)] = (d_225_v25_) == (d_225_v25_)
                nw27_[int(26)] = d_196_v0_
                d_231_v31_ = nw27_
                index29_ = default__.safeIndex(182, (d_231_v31_).length(0))
                (d_231_v31_)[index29_] = (d_220_v20_) <= (d_220_v20_)
                index30_ = default__.safeIndex(182, (d_231_v31_).length(0))
                (d_231_v31_)[index30_] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h"))) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kfkrjpla")))
                d_232_v32_: D0
                d_232_v32_ = D0_DC1(d_196_v0_, d_223_v23_, d_220_v20_)
                d_233_v33_: _dafny.MultiSet
                d_233_v33_ = _dafny.MultiSet([d_232_v32_, d_232_v32_, d_232_v32_, d_232_v32_])
                d_234_v34_: C0
                nw28_ = C0()
                nw28_.ctor__((D2_DC6(d_233_v33_)).cf10, not(not(False)))
                d_234_v34_ = nw28_
                d_231_v31_ = d_231_v31_
                d_235_v35_: _dafny.Array
                def lambda5_(d_236_v20_):
                    def lambda6_(d_237_i2_):
                        return (d_236_v20_) + (d_236_v20_)

                    return lambda6_

                init2_ = lambda5_(d_220_v20_)
                nw29_ = _dafny.Array(None, 11)
                for i0_2_ in range(nw29_.length(0)):
                    nw29_[i0_2_] = init2_(i0_2_)
                d_235_v35_ = nw29_
                index31_ = default__.safeIndex(664, (d_235_v35_).length(0))
                (d_235_v35_)[index31_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mcihcac"))
                d_238_v36_: _dafny.Map
                d_238_v36_ = _dafny.Map({(d_234_v34_).f19: ((d_214_v14_)[default__.safeIndex(755, (d_214_v14_).length(0))] if (d_234_v34_).f19 else len(_dafny.Set({584, (d_214_v14_)[default__.safeIndex(755, (d_214_v14_).length(0))]})))})
                index32_ = default__.safeIndex(664, (d_235_v35_).length(0))
                rhs44_ = (d_220_v20_) + ((d_220_v20_) + (d_220_v20_))
                rhs45_ = (d_214_v14_)[default__.safeIndex(755, (d_214_v14_).length(0))]
                rhs46_ = d_238_v36_
                lhs34_ = d_235_v35_
                lhs35_ = default__.safeIndex(664, (d_235_v35_).length(0))
                lhs34_[lhs35_] = rhs44_
                r0 = rhs45_
                d_238_v36_ = rhs46_
            elif True:
                (globalState).f3 = (d_198_v2_).cf4
                d_239_v37_: str
                d_239_v37_ = _dafny.CodePoint('s')
                d_240_v38_: _dafny.Array
                nw30_ = _dafny.Array(False, 19)
                d_240_v38_ = nw30_
                d_241_v39_: _dafny.MultiSet
                d_241_v39_ = _dafny.MultiSet([d_197_v1_, d_197_v1_, (d_214_v14_)[default__.safeIndex(755, (d_214_v14_).length(0))], (d_214_v14_)[default__.safeIndex(755, (d_214_v14_).length(0))]])
                d_242_v40_: _dafny.Seq
                d_242_v40_ = _dafny.SeqWithoutIsStrInference([d_220_v20_])
                d_243_v41_: D0
                d_243_v41_ = D0_DC2(d_196_v0_)
                d_244_v42_: D0
                d_244_v42_ = D0_DC3(d_243_v41_)
                d_245_v43_: _dafny.Map
                d_245_v43_ = _dafny.Map({(D0_DC1(d_196_v0_, d_242_v40_, d_220_v20_)).cf1: d_244_v42_})
                d_246_v44_: D2
                d_246_v44_ = D2_DC7(d_196_v0_, d_245_v43_, d_197_v1_, d_196_v0_)
                d_247_v45_: _dafny.Map
                d_247_v45_ = _dafny.Map({d_220_v20_: d_240_v38_})
                d_248_v46_: _dafny.Set
                d_248_v46_ = _dafny.Set({(0) - ((d_214_v14_)[default__.safeIndex(755, (d_214_v14_).length(0))])})
                index33_ = default__.safeIndex(755, (d_214_v14_).length(0))
                rhs47_ = default__.fm22(d_246_v44_, d_244_v42_, d_196_v0_, globalState)
                rhs48_ = default__.fm22(d_246_v44_, d_244_v42_, d_196_v0_, globalState)
                rhs49_ = default__.fm23(d_196_v0_, default__.fm24(globalState), d_196_v0_, (d_214_v14_)[default__.safeIndex(755, (d_214_v14_).length(0))], globalState)
                rhs50_ = ((d_247_v45_)[(d_220_v20_) + (d_220_v20_)] if ((d_220_v20_) + (d_220_v20_)) in (d_247_v45_) else d_240_v38_)
                rhs51_ = (d_241_v39_).set((len(d_220_v20_)) * (d_197_v1_), default__.abs(len((d_248_v46_) - (d_248_v46_))))
                lhs36_ = d_214_v14_
                lhs37_ = default__.safeIndex(755, (d_214_v14_).length(0))
                lhs36_[lhs37_] = rhs47_
                r0 = rhs48_
                d_239_v37_ = rhs49_
                d_240_v38_ = rhs50_
                d_241_v39_ = rhs51_
                d_249_v47_: D0
                d_249_v47_ = D0_DC1(d_196_v0_, d_242_v40_, default__.fm20(d_197_v1_, not(d_196_v0_), d_196_v0_, len(d_220_v20_), globalState))
                d_250_v48_: C0
                nw31_ = C0()
                nw31_.ctor__(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_249_v47_, d_249_v47_])), d_196_v0_)
                d_250_v48_ = nw31_
                index34_ = default__.safeIndex(755, (d_214_v14_).length(0))
                (d_214_v14_)[index34_] = (0) - ((d_197_v1_) + (default__.safeDivisionInt((d_214_v14_)[default__.safeIndex(755, (d_214_v14_).length(0))], (0) - (d_197_v1_))))
                d_251_v49_: C0
                nw32_ = C0()
                nw32_.ctor__(((d_250_v48_).f18).set(d_249_v47_, default__.abs(len(d_248_v46_))), not((d_250_v48_).f19))
                d_251_v49_ = nw32_
            (globalState).f5 = d_196_v0_
            d_252_v50_: D0
            d_252_v50_ = D0_DC0((d_214_v14_)[default__.safeIndex(755, (d_214_v14_).length(0))])
            d_253_v51_: D0
            d_253_v51_ = D0_DC3(d_252_v50_)
            d_254_v52_: _dafny.Map
            d_254_v52_ = _dafny.Map({d_196_v0_: d_253_v51_})
            d_255_v53_: D2
            d_255_v53_ = D2_DC7(d_196_v0_, d_254_v52_, d_197_v1_, False)
            index35_ = default__.safeIndex(755, (d_214_v14_).length(0))
            (d_214_v14_)[index35_] = default__.fm22(d_255_v53_, D0_DC3(d_252_v50_), not(d_196_v0_), globalState)
        d_256_v54_: _dafny.Map
        d_256_v54_ = _dafny.Map({d_197_v1_: d_197_v1_})
        d_257_v55_: _dafny.Seq
        d_257_v55_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tlaqlophn"))
        d_256_v54_ = (d_256_v54_).set(len(d_257_v55_), d_197_v1_)
        d_258_v56_: _dafny.Map
        d_258_v56_ = _dafny.Map({d_196_v0_: d_197_v1_})
        d_259_v57_: str
        d_259_v57_ = _dafny.CodePoint('g')
        d_260_v58_: _dafny.MultiSet
        d_260_v58_ = _dafny.MultiSet([d_259_v57_])
        r0 = (d_197_v1_) + (((d_258_v56_)[d_196_v0_] if (d_196_v0_) in (d_258_v56_) else ((d_260_v58_)[d_259_v57_] if (d_259_v57_) in (d_260_v58_) else d_197_v1_)))
        d_261_i3_: int
        d_261_i3_ = 0
        with _dafny.label("1"):
            while d_196_v0_:
                with _dafny.c_label("1"):
                    if (d_261_i3_) >= (100):
                        raise _dafny.Break("1")
                    d_261_i3_ = (d_261_i3_) + (1)
                    if (default__.fm19(d_197_v1_, d_196_v0_, globalState)) and (d_196_v0_):
                        d_262_v59_: _dafny.Array
                        def lambda7_(d_263_v0_):
                            def lambda8_(d_264_i4_):
                                return (d_264_i4_) - (len(_dafny.SeqWithoutIsStrInference([d_263_v0_])))

                            return lambda8_

                        init3_ = lambda7_(d_196_v0_)
                        nw33_ = _dafny.Array(None, 20)
                        for i0_3_ in range(nw33_.length(0)):
                            nw33_[i0_3_] = init3_(i0_3_)
                        d_262_v59_ = nw33_
                        index36_ = default__.safeIndex(961, (d_262_v59_).length(0))
                        (d_262_v59_)[index36_] = (0) - (d_197_v1_)
                        d_265_v60_: _dafny.Map
                        d_265_v60_ = _dafny.Map({d_197_v1_: d_196_v0_})
                        d_266_v61_: D0
                        d_266_v61_ = D0_DC3(D0_DC0(len(d_265_v60_)))
                        d_267_v62_: D0
                        d_267_v62_ = D0_DC3(d_266_v61_)
                        d_268_v63_: _dafny.Map
                        d_268_v63_ = _dafny.Map({d_196_v0_: d_267_v62_})
                        d_269_v64_: D2
                        d_269_v64_ = D2_DC7(d_196_v0_, d_268_v63_, d_197_v1_, d_196_v0_)
                        d_270_v65_: _dafny.Seq
                        d_270_v65_ = _dafny.SeqWithoutIsStrInference([d_196_v0_])
                        index37_ = default__.safeIndex(961, (d_262_v59_).length(0))
                        rhs52_ = default__.safeModuloInt(d_197_v1_, len(default__.fm20(len(d_258_v56_), d_196_v0_, d_196_v0_, default__.fm22(d_269_v64_, d_267_v62_, d_196_v0_, globalState), globalState)))
                        rhs53_ = not (d_196_v0_) or (d_196_v0_)
                        rhs54_ = len(d_270_v65_)
                        lhs38_ = d_262_v59_
                        lhs39_ = default__.safeIndex(961, (d_262_v59_).length(0))
                        lhs40_ = globalState
                        lhs38_[lhs39_] = rhs52_
                        lhs40_.f5 = rhs53_
                        d_197_v1_ = rhs54_
                        d_271_v66_: _dafny.MultiSet
                        d_271_v66_ = _dafny.MultiSet([(d_262_v59_)[default__.safeIndex(961, (d_262_v59_).length(0))]])
                        d_272_v67_: _dafny.Seq
                        d_272_v67_ = _dafny.SeqWithoutIsStrInference([(d_262_v59_)[default__.safeIndex(961, (d_262_v59_).length(0))], len(_dafny.SeqWithoutIsStrInference([(d_262_v59_)[default__.safeIndex(961, (d_262_v59_).length(0))]]))])
                        rhs55_ = ((default__.fm25(d_196_v0_, d_272_v67_, globalState)).intersection(d_271_v66_)) | ((_dafny.MultiSet([-802, d_197_v1_, d_197_v1_])) | (d_271_v66_))
                        rhs56_ = d_196_v0_
                        lhs41_ = globalState
                        d_271_v66_ = rhs55_
                        lhs41_.f13 = rhs56_
                        d_273_v68_: D1
                        d_273_v68_ = D1_DC5(d_196_v0_, d_257_v55_, (d_270_v65_).set(default__.safeIndex((0) - (d_197_v1_), len(d_270_v65_)), (d_270_v65_)[default__.safeIndex(d_197_v1_, len(d_270_v65_))]))
                        pat_let_tv6_ = d_257_v55_
                        d_274_v69_: _dafny.Array
                        nw34_ = _dafny.Array(None, 5)
                        nw34_[int(0)] = default__.fm26(d_196_v0_, d_196_v0_, globalState)
                        nw34_[int(1)] = d_273_v68_
                        nw34_[int(2)] = d_273_v68_
                        nw34_[int(3)] = d_273_v68_
                        def iife9_(_pat_let2_0):
                            def iife10_(d_275_dt__update__tmp_h0_):
                                def iife11_(_pat_let3_0):
                                    def iife12_(d_276_dt__update_hcf8_h0_):
                                        return D1_DC5((d_275_dt__update__tmp_h0_).cf7, d_276_dt__update_hcf8_h0_, (d_275_dt__update__tmp_h0_).cf9)
                                    return iife12_(_pat_let3_0)
                                return iife11_(pat_let_tv6_)
                            return iife10_(_pat_let2_0)
                        nw34_[int(4)] = iife9_(default__.fm26(d_196_v0_, d_196_v0_, globalState))
                        d_274_v69_ = nw34_
                        index38_ = default__.safeIndex(711, (d_274_v69_).length(0))
                        (d_274_v69_)[index38_] = d_273_v68_
                        index39_ = default__.safeIndex(711, (d_274_v69_).length(0))
                        (d_274_v69_)[index39_] = d_273_v68_
                        d_277_v70_: _dafny.Set
                        d_277_v70_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([(d_271_v66_).cardinality]), _dafny.SeqWithoutIsStrInference([d_197_v1_ for d_278_i5_ in range(default__.abs(-342))])})
                        index40_ = default__.safeIndex(961, (d_262_v59_).length(0))
                        index41_ = default__.safeIndex(961, (d_262_v59_).length(0))
                        rhs57_ = default__.safeDivisionInt((d_262_v59_)[default__.safeIndex(961, (d_262_v59_).length(0))], len((default__.fm20((0) - (len(d_272_v67_)), not(d_196_v0_), True, (d_262_v59_)[default__.safeIndex(961, (d_262_v59_).length(0))], globalState)) + (d_257_v55_)))
                        rhs58_ = (_dafny.Set({d_272_v67_})).ispropersubset(d_277_v70_)
                        rhs59_ = ((0) - (((d_262_v59_)[default__.safeIndex(961, (d_262_v59_).length(0))] if d_196_v0_ else 273))) + ((d_197_v1_) - (d_197_v1_))
                        rhs60_ = d_200_v4_
                        lhs42_ = d_262_v59_
                        lhs43_ = default__.safeIndex(961, (d_262_v59_).length(0))
                        lhs44_ = d_262_v59_
                        lhs45_ = default__.safeIndex(961, (d_262_v59_).length(0))
                        lhs42_[lhs43_] = rhs57_
                        d_196_v0_ = rhs58_
                        lhs44_[lhs45_] = rhs59_
                        d_200_v4_ = rhs60_
                        r0 = (d_272_v67_)[default__.safeIndex((d_262_v59_)[default__.safeIndex(961, (d_262_v59_).length(0))], len(d_272_v67_))]
                    elif True:
                        d_279_v71_: _dafny.MultiSet
                        d_279_v71_ = _dafny.MultiSet([len(_dafny.Map({d_196_v0_: d_259_v57_}))])
                        d_280_v72_: _dafny.Seq
                        d_280_v72_ = _dafny.SeqWithoutIsStrInference([len(d_257_v55_), d_197_v1_])
                        d_281_v73_: _dafny.MultiSet
                        d_281_v73_ = _dafny.MultiSet([d_280_v72_])
                        d_282_v74_: _dafny.Seq
                        d_282_v74_ = _dafny.SeqWithoutIsStrInference([d_256_v54_, d_256_v54_, d_256_v54_, d_256_v54_])
                        d_283_v75_: _dafny.Array
                        nw35_ = _dafny.Array(None, 14)
                        nw35_[int(0)] = ((d_279_v71_).set(d_197_v1_, default__.abs(d_197_v1_))).issubset(d_279_v71_)
                        nw35_[int(1)] = d_196_v0_
                        nw35_[int(2)] = False
                        nw35_[int(3)] = d_196_v0_
                        nw35_[int(4)] = (d_196_v0_) or (d_196_v0_)
                        nw35_[int(5)] = ((d_281_v73_).cardinality) != (d_197_v1_)
                        nw35_[int(6)] = d_196_v0_
                        nw35_[int(7)] = d_196_v0_
                        nw35_[int(8)] = d_196_v0_
                        nw35_[int(9)] = not (True) or (False)
                        nw35_[int(10)] = (False) or (d_196_v0_)
                        nw35_[int(11)] = (d_282_v74_) != (_dafny.SeqWithoutIsStrInference([_dafny.Map({d_197_v1_: d_197_v1_})]))
                        nw35_[int(12)] = d_196_v0_
                        nw35_[int(13)] = d_196_v0_
                        d_283_v75_ = nw35_
                        index42_ = default__.safeIndex(103, (d_283_v75_).length(0))
                        (d_283_v75_)[index42_] = d_196_v0_
                        index43_ = default__.safeIndex(103, (d_283_v75_).length(0))
                        (d_283_v75_)[index43_] = ((d_257_v55_) <= (d_257_v55_)) or (d_196_v0_)
                        d_284_v76_: D0
                        d_284_v76_ = D0_DC1((d_283_v75_)[default__.safeIndex(103, (d_283_v75_).length(0))], _dafny.SeqWithoutIsStrInference([d_257_v55_ for d_285_i6_ in range(default__.abs(608))]), d_257_v55_)
                        d_286_v77_: C0
                        nw36_ = C0()
                        nw36_.ctor__(_dafny.MultiSet([d_284_v76_, d_284_v76_, d_284_v76_]), d_196_v0_)
                        d_286_v77_ = nw36_
                        d_287_v78_: _dafny.Array
                        nw37_ = _dafny.Array(None, 17)
                        nw37_[int(0)] = d_286_v77_
                        nw37_[int(1)] = d_286_v77_
                        nw37_[int(2)] = d_286_v77_
                        nw37_[int(3)] = d_286_v77_
                        nw37_[int(4)] = d_286_v77_
                        nw37_[int(5)] = d_286_v77_
                        nw37_[int(6)] = d_286_v77_
                        nw37_[int(7)] = d_286_v77_
                        nw37_[int(8)] = d_286_v77_
                        nw37_[int(9)] = d_286_v77_
                        nw37_[int(10)] = d_286_v77_
                        nw37_[int(11)] = d_286_v77_
                        nw37_[int(12)] = d_286_v77_
                        nw37_[int(13)] = d_286_v77_
                        nw37_[int(14)] = d_286_v77_
                        nw37_[int(15)] = d_286_v77_
                        nw37_[int(16)] = d_286_v77_
                        d_287_v78_ = nw37_
                        d_288_v79_: _dafny.Map
                        d_288_v79_ = _dafny.Map({d_196_v0_: (d_287_v78_ if (d_283_v75_)[default__.safeIndex(103, (d_283_v75_).length(0))] else d_287_v78_)})
                        d_288_v79_ = (d_288_v79_).set(False, d_287_v78_)
                        d_257_v55_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "phrjej"))) + (d_257_v55_)
                        d_289_v80_: _dafny.Map
                        d_289_v80_ = _dafny.Map({(d_286_v77_).f19: d_258_v56_})
                        d_290_v81_: _dafny.Seq
                        d_290_v81_ = _dafny.SeqWithoutIsStrInference([True])
                        d_258_v56_ = ((d_258_v56_ if d_196_v0_ else _dafny.Map({(d_286_v77_).f19: d_197_v1_}))) | (((d_289_v80_)[(d_290_v81_)[default__.safeIndex(d_197_v1_, len(d_290_v81_))]] if ((d_290_v81_)[default__.safeIndex(d_197_v1_, len(d_290_v81_))]) in (d_289_v80_) else d_258_v56_))
                        index44_ = default__.safeIndex(103, (d_283_v75_).length(0))
                        (d_283_v75_)[index44_] = ((d_200_v4_)[(d_283_v75_)[default__.safeIndex(103, (d_283_v75_).length(0))]] if ((d_283_v75_)[default__.safeIndex(103, (d_283_v75_).length(0))]) in (d_200_v4_) else not((d_286_v77_).f19))
                    d_291_v83_: _dafny.MultiSet
                    d_291_v83_ = _dafny.MultiSet([d_197_v1_])
                    d_292_v84_: _dafny.Seq
                    def iife13_():
                        coll5_ = _dafny.Map()
                        compr_5_: int
                        for compr_5_ in (d_291_v83_).Elements:
                            d_293_v82_: int = compr_5_
                            if (d_293_v82_) in (d_291_v83_):
                                coll5_[default__.safeModuloInt(d_293_v82_, ((d_258_v56_)[d_196_v0_] if (d_196_v0_) in (d_258_v56_) else len(d_257_v55_)))] = len(_dafny.Map({len(_dafny.Set({d_197_v1_})): d_196_v0_}))
                        return _dafny.Map(coll5_)
                    d_292_v84_ = _dafny.SeqWithoutIsStrInference([d_197_v1_, len(iife13_()
                    ), d_197_v1_])
                    r0 = (d_292_v84_)[default__.safeIndex((d_292_v84_)[default__.safeIndex(d_197_v1_, len(d_292_v84_))], len(d_292_v84_))]
                    d_294_v86_: _dafny.Map
                    d_294_v86_ = _dafny.Map({d_197_v1_: d_257_v55_})
                    d_295_v88_: _dafny.Array
                    nw38_ = _dafny.Array(None, 11)
                    def iife14_():
                        coll6_ = _dafny.Map()
                        compr_6_: int
                        for compr_6_ in _dafny.IntegerRange(955, 953):
                            d_296_v85_: int = compr_6_
                            if ((955) <= (d_296_v85_)) and ((d_296_v85_) < (953)):
                                coll6_[default__.safeDivisionInt(d_296_v85_, -853)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_297_i7_ in range(default__.abs(346))])
                        return _dafny.Map(coll6_)
                    nw38_[int(0)] = (iife14_()
                    ) | (d_294_v86_)
                    nw38_[int(1)] = d_294_v86_
                    nw38_[int(2)] = (_dafny.Map({d_197_v1_: d_257_v55_})) | (d_294_v86_)
                    nw38_[int(3)] = ((d_294_v86_).set(418, default__.fm20(len(d_256_v54_), d_196_v0_, d_196_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gj"))), globalState))) | (d_294_v86_)
                    nw38_[int(4)] = (d_294_v86_) | (_dafny.Map({d_197_v1_: d_257_v55_}))
                    nw38_[int(5)] = _dafny.Map({d_197_v1_: d_257_v55_})
                    nw38_[int(6)] = d_294_v86_
                    nw38_[int(7)] = d_294_v86_
                    nw38_[int(8)] = (d_294_v86_) | (d_294_v86_)
                    nw38_[int(9)] = d_294_v86_
                    def iife15_():
                        coll7_ = _dafny.Map()
                        compr_7_: int
                        for compr_7_ in _dafny.IntegerRange(81, 133):
                            d_298_v87_: int = compr_7_
                            if ((81) <= (d_298_v87_)) and ((d_298_v87_) < (133)):
                                coll7_[(d_298_v87_) - ((d_260_v58_).cardinality)] = d_257_v55_
                        return _dafny.Map(coll7_)
                    nw38_[int(10)] = (iife15_()
                    ).set(d_197_v1_, d_257_v55_)
                    d_295_v88_ = nw38_
                    index45_ = default__.safeIndex(30, (d_295_v88_).length(0))
                    (d_295_v88_)[index45_] = (d_294_v86_) | (d_294_v86_)
                    d_299_v89_: _dafny.Seq
                    d_299_v89_ = _dafny.SeqWithoutIsStrInference([d_294_v86_])
                    index46_ = default__.safeIndex(30, (d_295_v88_).length(0))
                    (d_295_v88_)[index46_] = ((d_294_v86_) | ((d_299_v89_)[default__.safeIndex(d_197_v1_, len(d_299_v89_))])) | (d_294_v86_)
                    d_197_v1_ = d_197_v1_
                    pass
            pass
        d_300_v90_: _dafny.Seq
        d_300_v90_ = _dafny.SeqWithoutIsStrInference([d_197_v1_, d_197_v1_, d_197_v1_])
        d_301_v91_: D1
        d_301_v91_ = D1_DC5(d_196_v0_, d_257_v55_, default__.fm27(len(d_300_v90_), d_197_v1_, d_197_v1_, globalState))
        pat_let_tv7_ = d_197_v1_
        pat_let_tv8_ = d_197_v1_
        def lambda9_(source6_):
            if source6_.is_DC5:
                d_302___mcc_h0_ = source6_.cf7
                d_303___mcc_h1_ = source6_.cf8
                d_304___mcc_h2_ = source6_.cf9
                d_305_cf9_ = d_304___mcc_h2_
                d_306_cf8_ = d_303___mcc_h1_
                d_307_cf7_ = d_302___mcc_h0_
                return pat_let_tv7_
            elif True:
                d_308___mcc_h3_ = source6_.cf6
                d_309_cf6_ = d_308___mcc_h3_
                return pat_let_tv8_

        r0 = lambda9_(d_301_v91_)
        return r0


class C2(T0):
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    def ctor__(self):
        pass
        pass

    def fm0(self, p0, p1, globalState):
        return 226

    def fm1(self, p0, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fbtqkyvh"))

    def fm14(self, p0, p1, p2, globalState):
        return (_dafny.SeqWithoutIsStrInference([(D0_DC0(len(_dafny.Map({False: True})))).cf0 for d_310_i0_ in range(default__.abs(944))])) in (_dafny.Map({_dafny.SeqWithoutIsStrInference([-772 for d_311_i1_ in range(default__.abs(354))]): _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True, False]))]))}))

    def m6(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: bool = False
        (globalState).f13 = not (p2) or (not (p2) or (p2))
        r1 = default__.safeModuloInt(len(default__.fm15(p2, p2, p2, globalState)), ((self).fm0(len(p0), _dafny.CodePoint('p'), globalState) if p2 else p3))
        d_312_v0_: _dafny.Set
        d_312_v0_ = _dafny.Set({p3})
        hi2_ = len((d_312_v0_).intersection(d_312_v0_))
        for d_313_i0_ in range(p3, hi2_):
            d_314_v1_: _dafny.Array
            nw39_ = _dafny.Array(int(0), 16)
            d_314_v1_ = nw39_
            d_314_v1_ = d_314_v1_
            r1 = default__.safeDivisionInt(default__.safeDivisionInt(p3, d_313_i0_), p3)
            r1 = d_313_i0_
            d_315_v2_: _dafny.Array
            nw40_ = _dafny.Array(_dafny.Array(None, 0), 17)
            d_315_v2_ = nw40_
            d_316_v3_: _dafny.Array
            nw41_ = _dafny.Array(_dafny.Map({}), 25)
            d_316_v3_ = nw41_
            index47_ = default__.safeIndex(298, (d_315_v2_).length(0))
            (d_315_v2_)[index47_] = (d_316_v3_ if p2 else d_316_v3_)
            index48_ = default__.safeIndex(298, (d_315_v2_).length(0))
            (d_315_v2_)[index48_] = d_316_v3_
        d_317_v4_: _dafny.Map
        d_317_v4_ = _dafny.Map({p2: p2})
        d_318_v5_: _dafny.Seq
        d_318_v5_ = _dafny.SeqWithoutIsStrInference([(self).fm14(_dafny.Set({p3, len(d_317_v4_)}), not(p2), p1, globalState), False, True])
        d_319_v6_: str
        d_319_v6_ = _dafny.CodePoint('a')
        if (d_318_v5_)[default__.safeIndex(len(((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mkjhatikf"))) + (p0)).set(default__.safeIndex(631, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mkjhatikf"))) + (p0))), d_319_v6_)), len(d_318_v5_))]:
            d_320_v7_: _dafny.Map
            d_320_v7_ = _dafny.Map({p2: p3})
            d_321_v8_: _dafny.Array
            nw42_ = _dafny.Array(int(0), 3)
            d_321_v8_ = nw42_
            d_322_v9_: _dafny.Map
            d_322_v9_ = _dafny.Map({p3: -712})
            d_323_v10_: _dafny.Map
            d_323_v10_ = _dafny.Map({d_321_v8_: len(d_322_v9_)})
            d_324_v11_: _dafny.Array
            nw43_ = _dafny.Array(None, 12)
            nw43_[int(0)] = -737
            nw43_[int(1)] = ((d_320_v7_)[p2] if (p2) in (d_320_v7_) else p3)
            nw43_[int(2)] = p3
            nw43_[int(3)] = p3
            nw43_[int(4)] = p3
            nw43_[int(5)] = p3
            nw43_[int(6)] = p3
            nw43_[int(7)] = p3
            nw43_[int(8)] = p3
            nw43_[int(9)] = p3
            nw43_[int(10)] = len(d_323_v10_)
            nw43_[int(11)] = p3
            d_324_v11_ = nw43_
            d_325_v12_: _dafny.Map
            d_325_v12_ = _dafny.Map({p2: d_321_v8_})
            d_326_v13_: _dafny.Map
            d_326_v13_ = _dafny.Map({p3: d_319_v6_})
            d_327_v15_: _dafny.MultiSet
            def iife16_():
                coll8_ = _dafny.Set()
                compr_8_: int
                for compr_8_ in (d_326_v13_).keys.Elements:
                    d_328_v14_: int = compr_8_
                    if (d_328_v14_) in (d_326_v13_):
                        coll8_ = coll8_.union(_dafny.Set([(d_328_v14_) * (290)]))
                return _dafny.Set(coll8_)
            def iife17_():
                coll9_ = _dafny.Set()
                compr_9_: int
                for compr_9_ in (d_326_v13_).keys.Elements:
                    d_330_v14_: int = compr_9_
                    if (d_330_v14_) in (d_326_v13_):
                        coll9_ = coll9_.union(_dafny.Set([(d_330_v14_) * (290)]))
                return _dafny.Set(coll9_)
            d_327_v15_ = _dafny.MultiSet([((d_325_v12_)[(self).fm14(iife16_()
            , p2, _dafny.SeqWithoutIsStrInference([d_319_v6_ for d_329_i1_ in range(default__.abs(900))]), globalState)] if ((self).fm14(iife17_()
            , p2, _dafny.SeqWithoutIsStrInference([d_319_v6_ for d_331_i1_ in range(default__.abs(900))]), globalState)) in (d_325_v12_) else d_324_v11_), d_324_v11_, d_321_v8_, d_324_v11_, d_324_v11_])
            d_332_v16_: _dafny.MultiSet
            d_332_v16_ = _dafny.MultiSet([p3])
            rhs61_ = (d_327_v15_) | (d_327_v15_)
            rhs62_ = (d_332_v16_) - (default__.fm16(globalState))
            d_327_v15_ = rhs61_
            d_332_v16_ = rhs62_
            d_333_v17_: D0
            d_333_v17_ = D0_DC3(D0_DC0((self).fm0(144, _dafny.CodePoint('w'), globalState)))
            d_334_v18_: _dafny.Map
            d_334_v18_ = _dafny.Map({default__.safeDivisionInt(p3, 698): d_333_v17_})
            d_335_v20_: _dafny.Map
            d_335_v20_ = _dafny.Map({p3: p2})
            def iife18_():
                coll10_ = _dafny.Map()
                compr_10_: int
                for compr_10_ in (d_335_v20_).keys.Elements:
                    d_336_v19_: int = compr_10_
                    if (d_336_v19_) in (d_335_v20_):
                        coll10_[default__.safeDivisionInt(d_336_v19_, p3)] = d_333_v17_
                return _dafny.Map(coll10_)
            d_334_v18_ = (d_334_v18_) | (iife18_()
            )
            (globalState).f9 = p2
            d_337_v21_: _dafny.Array
            nw44_ = _dafny.Array(D0.default()(), 17)
            d_337_v21_ = nw44_
            index49_ = default__.safeIndex(373, (d_337_v21_).length(0))
            (d_337_v21_)[index49_] = default__.fm17(p2, globalState)
            d_338_v22_: _dafny.Map
            d_338_v22_ = _dafny.Map({p2: d_312_v0_})
            d_339_v23_: _dafny.Seq
            d_339_v23_ = _dafny.SeqWithoutIsStrInference([p3])
            d_340_v25_: _dafny.Seq
            def iife19_():
                coll11_ = _dafny.Set()
                compr_11_: int
                for compr_11_ in (d_339_v23_).Elements:
                    d_341_v24_: int = compr_11_
                    if (d_341_v24_) in (d_339_v23_):
                        coll11_ = coll11_.union(_dafny.Set([(d_341_v24_) + (len(_dafny.SeqWithoutIsStrInference([False])))]))
                return _dafny.Set(coll11_)
            d_340_v25_ = _dafny.SeqWithoutIsStrInference([(self).fm1(_dafny.Map({len(((d_338_v22_)[p2] if (p2) in (d_338_v22_) else iife19_()
            )): p3}), globalState), _dafny.SeqWithoutIsStrInference([d_319_v6_ for d_342_i2_ in range(default__.abs(-726))]), p1])
            d_343_v26_: _dafny.Map
            d_343_v26_ = _dafny.Map({(self).fm0(p3, d_319_v6_, globalState): d_340_v25_})
            d_344_v27_: D0
            d_344_v27_ = D0_DC1(p2, ((d_343_v26_)[p3] if (p3) in (d_343_v26_) else _dafny.SeqWithoutIsStrInference([p1])), (p0) + (p0))
            index50_ = default__.safeIndex(373, (d_337_v21_).length(0))
            (d_337_v21_)[index50_] = d_344_v27_
            d_345_v28_: _dafny.Array
            def lambda10_(d_346_p2_):
                def lambda11_(d_347_i3_):
                    return d_346_p2_

                return lambda11_

            init4_ = lambda10_(p2)
            nw45_ = _dafny.Array(None, 26)
            for i0_4_ in range(nw45_.length(0)):
                nw45_[i0_4_] = init4_(i0_4_)
            d_345_v28_ = nw45_
            index51_ = default__.safeIndex(650, (d_345_v28_).length(0))
            (d_345_v28_)[index51_] = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_348_i4_ in range(default__.abs(65))])) == (p1)
            index52_ = default__.safeIndex(650, (d_345_v28_).length(0))
            (d_345_v28_)[index52_] = not ((p2) or (p2)) or ((p3) < (len(d_318_v5_)))
        elif True:
            rhs63_ = p3
            rhs64_ = False
            lhs46_ = globalState
            r1 = rhs63_
            lhs46_.f3 = rhs64_
            (globalState).f3 = p2
            r1 = p3
            d_349_v29_: D0
            d_349_v29_ = D0_DC0(p3)
            d_350_v30_: _dafny.Seq
            d_350_v30_ = _dafny.SeqWithoutIsStrInference([d_349_v29_])
            d_351_v31_: _dafny.Seq
            d_351_v31_ = _dafny.SeqWithoutIsStrInference([d_350_v30_])
            d_352_v32_: _dafny.MultiSet
            d_352_v32_ = _dafny.MultiSet([p3, p3])
            d_353_v33_: _dafny.Seq
            d_353_v33_ = _dafny.SeqWithoutIsStrInference([d_352_v32_])
            d_351_v31_ = default__.fm18(d_319_v6_, ((d_353_v33_)[default__.safeIndex(p3, len(d_353_v33_))]).cardinality, globalState)
            r1 = p3
        d_354_v34_: _dafny.MultiSet
        d_354_v34_ = _dafny.MultiSet([d_312_v0_, d_312_v0_])
        hi3_ = ((0) - (len(_dafny.Map({p3: ((d_354_v34_)[d_312_v0_] if (d_312_v0_) in (d_354_v34_) else p3)})))) * ((0) - (p3))
        for d_355_i5_ in range(p3, hi3_):
            d_319_v6_ = d_319_v6_
            d_319_v6_ = _dafny.CodePoint('c')
            d_356_v35_: _dafny.Seq
            d_356_v35_ = _dafny.SeqWithoutIsStrInference([d_355_i5_])
            d_357_v36_: _dafny.Map
            d_357_v36_ = _dafny.Map({d_355_i5_: len(d_356_v35_)})
            d_358_v37_: _dafny.Map
            d_358_v37_ = _dafny.Map({d_355_i5_: len(d_357_v36_)})
            d_359_v38_: _dafny.Seq
            d_359_v38_ = _dafny.SeqWithoutIsStrInference([p0, p1, _dafny.SeqWithoutIsStrInference([d_319_v6_ for d_360_i6_ in range(default__.abs(29))]), (self).fm1(d_358_v37_, globalState)])
            d_361_v39_: D0
            d_361_v39_ = D0_DC1(p2, d_359_v38_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "unnkkr")))
            d_362_v40_: _dafny.MultiSet
            d_362_v40_ = _dafny.MultiSet([d_361_v39_])
            d_363_v41_: C0
            nw46_ = C0()
            nw46_.ctor__((d_362_v40_).set(d_361_v39_, default__.abs(d_355_i5_)), p2)
            d_363_v41_ = nw46_
            d_364_v42_: _dafny.Seq
            d_364_v42_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qumwtaj"))
            d_364_v42_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "juj"))
        hi4_ = default__.safeDivisionInt((0) - (p3), p3)
        for d_365_i7_ in range(p3, hi4_):
            d_366_v43_: _dafny.Seq
            d_366_v43_ = _dafny.SeqWithoutIsStrInference([p3, d_365_i7_])
            d_367_v44_: _dafny.Array
            nw47_ = _dafny.Array(None, 4)
            nw47_[int(0)] = p2
            nw47_[int(1)] = not(p2)
            nw47_[int(2)] = p2
            nw47_[int(3)] = p2
            d_367_v44_ = nw47_
            d_368_v45_: _dafny.Map
            d_368_v45_ = _dafny.Map({(d_366_v43_)[default__.safeIndex((0) - (p3), len(d_366_v43_))]: d_367_v44_})
            d_369_v46_: _dafny.Map
            d_369_v46_ = _dafny.Map({len(d_366_v43_): d_318_v5_})
            d_368_v45_ = (d_368_v45_).set(len(((d_369_v46_)[d_365_i7_] if (d_365_i7_) in (d_369_v46_) else d_318_v5_)), d_367_v44_)
            d_319_v6_ = d_319_v6_
            r1 = (self).fm0(d_365_i7_, d_319_v6_, globalState)
            d_370_v47_: D1
            d_370_v47_ = D1_DC5(p2, p0, (_dafny.SeqWithoutIsStrInference([p2, p2])) + (d_318_v5_))
            d_370_v47_ = D1_DC5(p2, (p0) + (p0), (d_318_v5_) + (d_318_v5_))
        d_371_v48_: _dafny.Map
        d_371_v48_ = _dafny.Map({p3: not (p2) or (True)})
        d_372_v49_: _dafny.Seq
        d_372_v49_ = _dafny.SeqWithoutIsStrInference([p3])
        r0 = not(((d_371_v48_)[(d_372_v49_)[default__.safeIndex(p3, len(d_372_v49_))]] if ((d_372_v49_)[default__.safeIndex(p3, len(d_372_v49_))]) in (d_371_v48_) else not(p2)))
        r1 = p3
        r2 = p2
        return r0, r1, r2

    def m7(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: _dafny.Seq = _dafny.Seq({})
        r2: bool = False
        d_373_v0_: str
        d_373_v0_ = _dafny.CodePoint('b')
        d_374_v1_: _dafny.MultiSet
        d_374_v1_ = _dafny.MultiSet([d_373_v0_, _dafny.CodePoint('b'), d_373_v0_, d_373_v0_])
        d_375_v2_: _dafny.Map
        d_375_v2_ = _dafny.Map({p3: (d_374_v1_).cardinality})
        d_376_v3_: _dafny.MultiSet
        d_376_v3_ = _dafny.MultiSet([d_375_v2_])
        d_377_v4_: _dafny.MultiSet
        d_377_v4_ = _dafny.MultiSet([not(p3), p3])
        d_378_v5_: _dafny.Seq
        d_378_v5_ = _dafny.SeqWithoutIsStrInference([True])
        d_379_v6_: _dafny.Array
        nw48_ = _dafny.Array(int(0), 12)
        d_379_v6_ = nw48_
        d_380_v7_: _dafny.Map
        d_380_v7_ = _dafny.Map({len(d_378_v5_): d_379_v6_})
        d_381_v8_: _dafny.Map
        d_381_v8_ = _dafny.Map({len(p0): p2})
        d_382_v9_: _dafny.Seq
        d_382_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uwerfkjaf"))
        d_383_v10_: _dafny.Map
        d_383_v10_ = _dafny.Map({(0) - ((0) - (len(d_382_v9_))): p3})
        d_384_v11_: _dafny.Array
        nw49_ = _dafny.Array(None, 25)
        nw49_[int(0)] = (0) - (p1)
        nw49_[int(1)] = (p2) - (p2)
        nw49_[int(2)] = (435 if p3 else p1)
        nw49_[int(3)] = p2
        nw49_[int(4)] = (p1) + (766)
        nw49_[int(5)] = (d_376_v3_).cardinality
        nw49_[int(6)] = p2
        nw49_[int(7)] = default__.safeModuloInt((d_377_v4_).cardinality, p2)
        nw49_[int(8)] = (len(d_378_v5_)) - (len(d_380_v7_))
        nw49_[int(9)] = (0) - (p1)
        nw49_[int(10)] = ((d_381_v8_)[(0) - (p1)] if ((0) - (p1)) in (d_381_v8_) else 970)
        nw49_[int(11)] = p1
        nw49_[int(12)] = (p2) * (p1)
        nw49_[int(13)] = len((d_382_v9_) + (d_382_v9_))
        nw49_[int(14)] = -260
        nw49_[int(15)] = p2
        nw49_[int(16)] = p2
        nw49_[int(17)] = ((d_375_v2_)[p3] if (p3) in (d_375_v2_) else p1)
        nw49_[int(18)] = p2
        nw49_[int(19)] = -272
        nw49_[int(20)] = p2
        nw49_[int(21)] = 156
        nw49_[int(22)] = p1
        nw49_[int(23)] = len(d_383_v10_)
        nw49_[int(24)] = p2
        d_384_v11_ = nw49_
        def lambda12_(d_385_p3_, d_386_p1_):
            def lambda13_(d_387_i0_):
                return (d_387_i0_) * (len(_dafny.Map({d_385_p3_: d_386_p1_})))

            return lambda13_

        init5_ = lambda12_(p3, p1)
        nw50_ = _dafny.Array(None, 8)
        for i0_5_ in range(nw50_.length(0)):
            nw50_[i0_5_] = init5_(i0_5_)
        d_384_v11_ = nw50_
        d_388_i1_: int
        d_388_i1_ = 0
        with _dafny.label("2"):
            while p3:
                with _dafny.c_label("2"):
                    if (d_388_i1_) >= (100):
                        raise _dafny.Break("2")
                    d_388_i1_ = (d_388_i1_) + (1)
                    d_389_v12_: _dafny.Set
                    d_389_v12_ = _dafny.Set({p1})
                    d_390_v13_: _dafny.Map
                    d_390_v13_ = _dafny.Map({d_389_v12_: d_382_v9_})
                    d_391_v14_: _dafny.Set
                    d_391_v14_ = _dafny.Set({(p1 if False else len(((d_390_v13_)[d_389_v12_] if (d_389_v12_) in (d_390_v13_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j"))))), p1})
                    d_389_v12_ = d_389_v12_
                    d_392_v15_: _dafny.Map
                    d_392_v15_ = _dafny.Map({p3: default__.fm28(globalState)})
                    d_393_v16_: D2
                    d_393_v16_ = D2_DC7(p3, (d_392_v15_) | (d_392_v15_), p1, p3)
                    source7_ = d_393_v16_
                    if source7_.is_DC7:
                        d_394___mcc_h0_ = source7_.cf11
                        d_395___mcc_h1_ = source7_.cf12
                        d_396___mcc_h2_ = source7_.cf13
                        d_397___mcc_h3_ = source7_.cf14
                        d_398_cf14_ = d_397___mcc_h3_
                        d_399_cf13_ = d_396___mcc_h2_
                        d_400_cf12_ = d_395___mcc_h1_
                        d_401_cf11_ = d_394___mcc_h0_
                        d_399_cf13_ = 655
                        d_402_v17_: _dafny.Seq
                        d_402_v17_ = _dafny.SeqWithoutIsStrInference([d_382_v9_, d_382_v9_])
                        d_403_v18_: D0
                        d_403_v18_ = D0_DC1(p3, d_402_v17_, _dafny.SeqWithoutIsStrInference([d_373_v0_ for d_404_i2_ in range(default__.abs(574))]))
                        d_405_v19_: _dafny.MultiSet
                        d_405_v19_ = _dafny.MultiSet([d_403_v18_, d_403_v18_])
                        d_406_v20_: C0
                        nw51_ = C0()
                        nw51_.ctor__(d_405_v19_, d_401_cf11_)
                        d_406_v20_ = nw51_
                        d_407_v21_: _dafny.Map
                        d_407_v21_ = _dafny.Map({d_401_cf11_: (d_406_v20_).f19})
                        (globalState).f5 = ((d_407_v21_)[p3] if (p3) in (d_407_v21_) else d_401_cf11_)
                        d_408_v22_: _dafny.Map
                        d_408_v22_ = _dafny.Map({p3: p0})
                        d_409_v23_: D0
                        d_409_v23_ = D0_DC2(True)
                        d_410_v24_: _dafny.Array
                        nw52_ = _dafny.Array(None, 26)
                        nw52_[int(0)] = (len(d_382_v9_)) != (d_399_cf13_)
                        nw52_[int(1)] = True
                        nw52_[int(2)] = ((d_406_v20_).f19) == (False)
                        nw52_[int(3)] = d_401_cf11_
                        nw52_[int(4)] = (d_383_v10_) == (d_383_v10_)
                        nw52_[int(5)] = not((d_406_v20_).f19)
                        nw52_[int(6)] = d_398_cf14_
                        nw52_[int(7)] = d_398_cf14_
                        nw52_[int(8)] = (d_399_cf13_) < (len(p0))
                        nw52_[int(9)] = not(d_398_cf14_)
                        nw52_[int(10)] = p3
                        nw52_[int(11)] = d_401_cf11_
                        nw52_[int(12)] = (p0).isdisjoint(((d_408_v22_)[d_398_cf14_] if (d_398_cf14_) in (d_408_v22_) else p0))
                        nw52_[int(13)] = p3
                        nw52_[int(14)] = (d_406_v20_).f19
                        nw52_[int(15)] = (d_401_cf11_) == (d_401_cf11_)
                        nw52_[int(16)] = (d_406_v20_).f19
                        nw52_[int(17)] = (not((d_406_v20_).f19)) == (d_398_cf14_)
                        nw52_[int(18)] = d_398_cf14_
                        nw52_[int(19)] = p3
                        nw52_[int(20)] = default__.fm19(d_399_cf13_, d_401_cf11_, globalState)
                        nw52_[int(21)] = (d_409_v23_).cf4
                        nw52_[int(22)] = not(p3)
                        nw52_[int(23)] = False
                        nw52_[int(24)] = not((d_389_v12_).issubset(d_391_v14_))
                        nw52_[int(25)] = (d_406_v20_).f19
                        d_410_v24_ = nw52_
                        index53_ = default__.safeIndex(271, (d_410_v24_).length(0))
                        (d_410_v24_)[index53_] = (d_406_v20_).f19
                        d_411_v25_: _dafny.Seq
                        d_411_v25_ = _dafny.SeqWithoutIsStrInference([p2, len(d_382_v9_), p1, p2, d_399_cf13_])
                        d_412_v26_: _dafny.Map
                        d_412_v26_ = _dafny.Map({990: d_382_v9_})
                        index54_ = default__.safeIndex(271, (d_410_v24_).length(0))
                        (d_410_v24_)[index54_] = (self).fm14(d_389_v12_, (len(d_382_v9_)) in (d_411_v25_), ((d_412_v26_)[p1] if (p1) in (d_412_v26_) else d_382_v9_), globalState)
                    elif source7_.is_DC6:
                        d_413___mcc_h4_ = source7_.cf10
                        d_414_cf10_ = d_413___mcc_h4_
                        d_415_v27_: _dafny.Seq
                        d_415_v27_ = _dafny.SeqWithoutIsStrInference([-738])
                        index55_ = default__.safeIndex(326, (d_379_v6_).length(0))
                        (d_379_v6_)[index55_] = (p2) * ((d_415_v27_)[default__.safeIndex(p2, len(d_415_v27_))])
                        d_416_v28_: D0
                        d_416_v28_ = D0_DC0(p1)
                        index56_ = default__.safeIndex(326, (d_379_v6_).length(0))
                        rhs65_ = d_415_v27_
                        rhs66_ = (p2) != (p1)
                        rhs67_ = default__.fm22(d_393_v16_, D0_DC3(d_416_v28_), p3, globalState)
                        lhs47_ = d_379_v6_
                        lhs48_ = default__.safeIndex(326, (d_379_v6_).length(0))
                        d_415_v27_ = rhs65_
                        r2 = rhs66_
                        lhs47_[lhs48_] = rhs67_
                        d_417_v29_: bool
                        d_418_v30_: int
                        d_419_v31_: bool
                        out2_: bool
                        out3_: int
                        out4_: bool
                        out2_, out3_, out4_ = (self).m6(d_382_v9_, d_382_v9_, p3, p2, globalState)
                        d_417_v29_ = out2_
                        d_418_v30_ = out3_
                        d_419_v31_ = out4_
                        d_420_v32_: D0
                        d_420_v32_ = D0_DC3(D0_DC0((d_379_v6_)[default__.safeIndex(326, (d_379_v6_).length(0))]))
                        pat_let_tv9_ = d_416_v28_
                        index57_ = default__.safeIndex(326, (d_379_v6_).length(0))
                        def iife20_(_pat_let4_0):
                            def iife21_(d_421_dt__update__tmp_h0_):
                                def iife22_(_pat_let5_0):
                                    def iife23_(d_422_dt__update_hcf5_h0_):
                                        return D0_DC3(d_422_dt__update_hcf5_h0_)
                                    return iife23_(_pat_let5_0)
                                return iife22_(pat_let_tv9_)
                            return iife21_(_pat_let4_0)
                        (d_379_v6_)[index57_] = (default__.fm22(d_393_v16_, iife20_(d_420_v32_), d_417_v29_, globalState)) - (len(d_378_v5_))
                        d_423_v33_: _dafny.Map
                        d_423_v33_ = _dafny.Map({d_419_v31_: d_419_v31_})
                        d_424_v34_: _dafny.Map
                        d_424_v34_ = _dafny.Map({d_373_v0_: d_423_v33_})
                        d_373_v0_ = default__.fm23(d_419_v31_, d_424_v34_, d_417_v29_, -87, globalState)
                    elif True:
                        d_425___mcc_h5_ = source7_.cf15
                        d_426_cf15_ = d_425___mcc_h5_
                        d_427_v35_: C1
                        nw53_ = C1()
                        nw53_.ctor__()
                        d_427_v35_ = nw53_
                        d_428_v36_: _dafny.Map
                        d_428_v36_ = _dafny.Map({p2: d_373_v0_})
                        d_428_v36_ = (d_428_v36_).set((p1) - (p1), d_373_v0_)
                        (globalState).f9 = p3
                        d_429_v37_: _dafny.Array
                        nw54_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 25)
                        d_429_v37_ = nw54_
                        index58_ = default__.safeIndex(158, (d_429_v37_).length(0))
                        (d_429_v37_)[index58_] = (d_382_v9_) + (d_382_v9_)
                        d_430_v38_: _dafny.Map
                        d_430_v38_ = _dafny.Map({p3: d_382_v9_})
                        d_431_v39_: _dafny.Seq
                        d_431_v39_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_373_v0_, d_373_v0_]), d_382_v9_, ((d_430_v38_)[p3] if (p3) in (d_430_v38_) else d_382_v9_), d_382_v9_, _dafny.SeqWithoutIsStrInference([d_373_v0_])])
                        index59_ = default__.safeIndex(158, (d_429_v37_).length(0))
                        (d_429_v37_)[index59_] = ((d_431_v39_)[default__.safeIndex((0) - ((self).fm0(p2, d_373_v0_, globalState)), len(d_431_v39_))]) + (d_382_v9_)
                    d_432_v40_: _dafny.Seq
                    d_432_v40_ = _dafny.SeqWithoutIsStrInference([d_382_v9_])
                    d_433_v41_: D0
                    d_433_v41_ = D0_DC1(p3, d_432_v40_, d_382_v9_)
                    d_434_v42_: _dafny.MultiSet
                    d_434_v42_ = _dafny.MultiSet([d_433_v41_, d_433_v41_])
                    d_435_v43_: C0
                    nw55_ = C0()
                    nw55_.ctor__(d_434_v42_, p3)
                    d_435_v43_ = nw55_
                    d_436_v44_: _dafny.Seq
                    d_436_v44_ = _dafny.SeqWithoutIsStrInference([p1, p2])
                    d_437_v45_: _dafny.Map
                    d_437_v45_ = _dafny.Map({d_436_v44_: p1})
                    d_437_v45_ = (d_437_v45_).set(d_436_v44_, p2)
                    pass
            pass
        d_438_v46_: int
        d_438_v46_ = -694
        d_438_v46_ = 278
        d_439_v47_: _dafny.Set
        d_439_v47_ = _dafny.Set({p1})
        source8_ = D0_DC0(len(d_439_v47_))
        if source8_.is_DC1:
            d_440___mcc_h6_ = source8_.cf1
            d_441___mcc_h7_ = source8_.cf2
            d_442___mcc_h8_ = source8_.cf3
            d_443_cf3_ = d_442___mcc_h8_
            d_444_cf2_ = d_441___mcc_h7_
            d_445_cf1_ = d_440___mcc_h6_
            d_446_v48_: _dafny.Array
            nw56_ = _dafny.Array(_dafny.Map({}), 19)
            d_446_v48_ = nw56_
            d_447_v49_: _dafny.Seq
            d_447_v49_ = _dafny.SeqWithoutIsStrInference([p2, p1])
            d_448_v50_: _dafny.Map
            d_448_v50_ = _dafny.Map({D1_DC5(d_445_cf1_, d_443_cf3_, d_378_v5_): d_447_v49_})
            index60_ = default__.safeIndex(99, (d_446_v48_).length(0))
            (d_446_v48_)[index60_] = d_448_v50_
            d_449_v51_: D0
            d_449_v51_ = D0_DC1(p3, d_444_cf2_, d_443_cf3_)
            d_450_v52_: _dafny.Map
            d_450_v52_ = _dafny.Map({True: D0_DC3(d_449_v51_)})
            d_451_v53_: D2
            d_451_v53_ = D2_DC7(p3, d_450_v52_, len(_dafny.SeqWithoutIsStrInference([d_445_cf1_])), d_445_cf1_)
            index61_ = default__.safeIndex(285, (d_384_v11_).length(0))
            (d_384_v11_)[index61_] = ((d_451_v53_).cf13) * (d_438_v46_)
            index62_ = default__.safeIndex(99, (d_446_v48_).length(0))
            index63_ = default__.safeIndex(285, (d_384_v11_).length(0))
            rhs68_ = d_448_v50_
            rhs69_ = d_443_cf3_
            rhs70_ = d_438_v46_
            rhs71_ = p3
            rhs72_ = (0) - (default__.safeModuloInt((len(d_439_v47_)) - (p1), p2))
            lhs49_ = d_446_v48_
            lhs50_ = default__.safeIndex(99, (d_446_v48_).length(0))
            lhs51_ = d_384_v11_
            lhs52_ = default__.safeIndex(285, (d_384_v11_).length(0))
            lhs53_ = globalState
            lhs49_[lhs50_] = rhs68_
            d_443_cf3_ = rhs69_
            lhs51_[lhs52_] = rhs70_
            lhs53_.f13 = rhs71_
            d_438_v46_ = rhs72_
            index64_ = default__.safeIndex(285, (d_384_v11_).length(0))
            (d_384_v11_)[index64_] = default__.safeDivisionInt(d_438_v46_, default__.safeModuloInt(p1, p1))
            d_452_v54_: _dafny.Map
            d_452_v54_ = _dafny.Map({p3: _dafny.Map({d_445_cf1_: d_379_v6_})})
            d_453_v55_: _dafny.Map
            d_453_v55_ = _dafny.Map({p3: d_379_v6_})
            d_454_v56_: D1
            d_454_v56_ = D1_DC4(((d_452_v54_)[d_445_cf1_] if (d_445_cf1_) in (d_452_v54_) else d_453_v55_))
            d_455_v57_: _dafny.Map
            d_455_v57_ = _dafny.Map({d_454_v56_: (d_451_v53_).cf13})
            d_455_v57_ = (d_455_v57_).set(d_454_v56_, len(d_444_cf2_))
            index65_ = default__.safeIndex(285, (d_384_v11_).length(0))
            rhs73_ = (d_375_v2_) | (d_375_v2_)
            rhs74_ = -205
            rhs75_ = (d_443_cf3_) + (d_382_v9_)
            rhs76_ = (d_378_v5_ if d_445_cf1_ else (d_378_v5_) + (d_378_v5_))
            lhs54_ = d_384_v11_
            lhs55_ = default__.safeIndex(285, (d_384_v11_).length(0))
            d_375_v2_ = rhs73_
            lhs54_[lhs55_] = rhs74_
            d_443_cf3_ = rhs75_
            r0 = rhs76_
        elif source8_.is_DC2:
            d_456___mcc_h9_ = source8_.cf4
            d_457_cf4_ = d_456___mcc_h9_
            d_458_v58_: _dafny.Array
            nw57_ = _dafny.Array(None, 2)
            nw57_[int(0)] = d_457_cf4_
            nw57_[int(1)] = default__.fm19(len(_dafny.SeqWithoutIsStrInference([d_373_v0_ for d_459_i3_ in range(default__.abs(-373))])), d_457_cf4_, globalState)
            d_458_v58_ = nw57_
            d_460_v59_: _dafny.Seq
            d_460_v59_ = _dafny.SeqWithoutIsStrInference([d_382_v9_, _dafny.SeqWithoutIsStrInference([d_373_v0_ for d_461_i4_ in range(default__.abs(68))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))])
            d_462_v60_: _dafny.Seq
            d_462_v60_ = _dafny.SeqWithoutIsStrInference([len(d_460_v59_)])
            d_463_v61_: _dafny.Map
            d_463_v61_ = _dafny.Map({d_458_v58_: (d_462_v60_) + (d_462_v60_)})
            d_438_v46_ = (0) - (len(d_463_v61_))
            d_464_v62_: _dafny.Array
            nw58_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 8)
            d_464_v62_ = nw58_
            index66_ = default__.safeIndex(753, (d_464_v62_).length(0))
            (d_464_v62_)[index66_] = (d_382_v9_) + (d_382_v9_)
            index67_ = default__.safeIndex(510, (d_379_v6_).length(0))
            (d_379_v6_)[index67_] = default__.safeDivisionInt(len(_dafny.Map({d_457_cf4_: p3})), len(p0))
            index68_ = default__.safeIndex(753, (d_464_v62_).length(0))
            index69_ = default__.safeIndex(510, (d_379_v6_).length(0))
            rhs77_ = d_384_v11_
            rhs78_ = d_382_v9_
            rhs79_ = (0) - ((0) - ((default__.safeDivisionInt(len(d_462_v60_), len(d_375_v2_))) - (p2)))
            rhs80_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dwq"))
            lhs56_ = d_464_v62_
            lhs57_ = default__.safeIndex(753, (d_464_v62_).length(0))
            lhs58_ = d_379_v6_
            lhs59_ = default__.safeIndex(510, (d_379_v6_).length(0))
            d_379_v6_ = rhs77_
            lhs56_[lhs57_] = rhs78_
            lhs58_[lhs59_] = rhs79_
            d_382_v9_ = rhs80_
            d_465_v63_: D0
            d_465_v63_ = D0_DC0((p1) + (p2))
            source9_ = d_465_v63_
            if source9_.is_DC1:
                d_466___mcc_h12_ = source9_.cf1
                d_467___mcc_h13_ = source9_.cf2
                d_468___mcc_h14_ = source9_.cf3
                d_469_cf3_ = d_468___mcc_h14_
                d_470_cf2_ = d_467___mcc_h13_
                d_471_cf1_ = d_466___mcc_h12_
                (globalState).f9 = d_471_cf1_
                d_472_v64_: D0
                d_472_v64_ = D0_DC1(d_471_cf1_, d_460_v59_, (d_464_v62_)[default__.safeIndex(753, (d_464_v62_).length(0))])
                d_473_v65_: D0
                d_473_v65_ = D0_DC3(d_472_v64_)
                d_474_v66_: D0
                d_474_v66_ = D0_DC3(d_472_v64_)
                d_475_v67_: D2
                d_475_v67_ = D2_DC7(d_457_cf4_, _dafny.Map({d_471_cf1_: d_474_v66_}), 779, d_471_cf1_)
                d_438_v46_ = default__.fm22(d_475_v67_, d_474_v66_, (d_378_v5_)[default__.safeIndex((_dafny.MultiSet([d_438_v46_])).cardinality, len(d_378_v5_))], globalState)
                d_476_v68_: _dafny.MultiSet
                d_476_v68_ = _dafny.MultiSet([d_384_v11_, d_384_v11_])
                d_477_v69_: _dafny.Seq
                d_477_v69_ = _dafny.SeqWithoutIsStrInference([(d_476_v68_).set(d_384_v11_, default__.abs(d_438_v46_))])
                index70_ = default__.safeIndex(510, (d_379_v6_).length(0))
                (d_379_v6_)[index70_] = (0) - (((d_477_v69_)[default__.safeIndex(706, len(d_477_v69_))]).cardinality)
                d_458_v58_ = d_458_v58_
            elif source9_.is_DC2:
                d_478___mcc_h15_ = source9_.cf4
                d_479_cf4_ = d_478___mcc_h15_
                d_480_v70_: D0
                d_480_v70_ = D0_DC0(p1)
                d_481_v71_: D0
                d_481_v71_ = D0_DC3(d_480_v70_)
                d_482_v72_: _dafny.Map
                d_482_v72_ = _dafny.Map({p3: d_481_v71_})
                d_483_v73_: D2
                d_483_v73_ = D2_DC7(p3, d_482_v72_, (0) - (len(_dafny.SeqWithoutIsStrInference([d_378_v5_ for d_484_i5_ in range(default__.abs(804))]))), p3)
                d_485_v74_: _dafny.Set
                d_485_v74_ = _dafny.Set({d_373_v0_})
                pat_let_tv10_ = p3
                pat_let_tv11_ = d_482_v72_
                pat_let_tv12_ = d_378_v5_
                pat_let_tv13_ = d_479_cf4_
                pat_let_tv14_ = p3
                d_486_v77_: _dafny.Array
                nw59_ = _dafny.Array(None, 26)
                nw59_[int(0)] = d_483_v73_
                nw59_[int(1)] = d_483_v73_
                nw59_[int(2)] = d_483_v73_
                def iife24_(_pat_let6_0):
                    def iife25_(d_487_dt__update__tmp_h1_):
                        def iife26_(_pat_let7_0):
                            def iife27_(d_488_dt__update_hcf11_h0_):
                                def iife28_(_pat_let8_0):
                                    def iife29_(d_489_dt__update_hcf12_h0_):
                                        return D2_DC7(d_488_dt__update_hcf11_h0_, d_489_dt__update_hcf12_h0_, (d_487_dt__update__tmp_h1_).cf13, (d_487_dt__update__tmp_h1_).cf14)
                                    return iife29_(_pat_let8_0)
                                return iife28_(pat_let_tv11_)
                            return iife27_(_pat_let7_0)
                        return iife26_(pat_let_tv10_)
                    return iife25_(_pat_let6_0)
                nw59_[int(3)] = iife24_(d_483_v73_)
                nw59_[int(4)] = d_483_v73_
                nw59_[int(5)] = d_483_v73_
                def iife30_(_pat_let9_0):
                    def iife31_(d_490_dt__update__tmp_h2_):
                        def iife32_(_pat_let10_0):
                            def iife33_(d_491_dt__update_hcf13_h0_):
                                def iife34_(_pat_let11_0):
                                    def iife35_(d_492_dt__update_hcf14_h0_):
                                        def iife36_(_pat_let12_0):
                                            def iife37_(d_493_dt__update_hcf11_h1_):
                                                return D2_DC7(d_493_dt__update_hcf11_h1_, (d_490_dt__update__tmp_h2_).cf12, d_491_dt__update_hcf13_h0_, d_492_dt__update_hcf14_h0_)
                                            return iife37_(_pat_let12_0)
                                        return iife36_(pat_let_tv14_)
                                    return iife35_(_pat_let11_0)
                                return iife34_(pat_let_tv13_)
                            return iife33_(_pat_let10_0)
                        return iife32_(len(pat_let_tv12_))
                    return iife31_(_pat_let9_0)
                nw59_[int(6)] = iife30_(d_483_v73_)
                nw59_[int(7)] = d_483_v73_
                nw59_[int(8)] = d_483_v73_
                nw59_[int(9)] = d_483_v73_
                nw59_[int(10)] = d_483_v73_
                nw59_[int(11)] = d_483_v73_
                nw59_[int(12)] = d_483_v73_
                nw59_[int(13)] = d_483_v73_
                nw59_[int(14)] = d_483_v73_
                nw59_[int(15)] = d_483_v73_
                nw59_[int(16)] = D2_DC7(p3, (d_482_v72_).set(p3, d_481_v71_), len(d_485_v74_), d_479_cf4_)
                nw59_[int(17)] = default__.fm29(len(d_375_v2_), (d_379_v6_)[default__.safeIndex(510, (d_379_v6_).length(0))], (d_464_v62_)[default__.safeIndex(753, (d_464_v62_).length(0))], p1, globalState)
                nw59_[int(18)] = D2_DC7((d_378_v5_)[default__.safeIndex(d_438_v46_, len(d_378_v5_))], d_482_v72_, (d_379_v6_)[default__.safeIndex(510, (d_379_v6_).length(0))], p3)
                nw59_[int(19)] = d_483_v73_
                nw59_[int(20)] = d_483_v73_
                nw59_[int(21)] = d_483_v73_
                nw59_[int(22)] = d_483_v73_
                def iife38_():
                    def iife40_():
                        coll14_ = _dafny.Map()
                        compr_14_: int
                        for compr_14_ in _dafny.IntegerRange(288, -460):
                            d_494_v75_: int = compr_14_
                            if ((288) <= (d_494_v75_)) and ((d_494_v75_) < (-460)):
                                coll14_[(d_494_v75_) - (p1)] = d_479_cf4_
                        return _dafny.Map(coll14_)
                    coll12_ = _dafny.Set()
                    def iife39_():
                        coll13_ = _dafny.Map()
                        compr_13_: int
                        for compr_13_ in _dafny.IntegerRange(288, -460):
                            d_494_v75_: int = compr_13_
                            if ((288) <= (d_494_v75_)) and ((d_494_v75_) < (-460)):
                                coll13_[(d_494_v75_) - (p1)] = d_479_cf4_
                        return _dafny.Map(coll13_)
                    compr_12_: int
                    for compr_12_ in (iife39_()
                    ).keys.Elements:
                        d_495_v76_: int = compr_12_
                        if (d_495_v76_) in (iife40_()
                        ):
                            coll12_ = coll12_.union(_dafny.Set([(d_495_v76_) - (-977)]))
                    return _dafny.Set(coll12_)
                nw59_[int(23)] = (d_483_v73_ if d_479_cf4_ else D2_DC7((self).fm14(iife38_()
, False, d_382_v9_, globalState), d_482_v72_, len(_dafny.Map({default__.fm30(globalState): d_479_cf4_})), p3))
                nw59_[int(24)] = d_483_v73_
                nw59_[int(25)] = d_483_v73_
                d_486_v77_ = nw59_
                index71_ = default__.safeIndex(262, (d_486_v77_).length(0))
                (d_486_v77_)[index71_] = D2_DC7(p3, d_482_v72_, (d_379_v6_)[default__.safeIndex(510, (d_379_v6_).length(0))], p3)
                index72_ = default__.safeIndex(262, (d_486_v77_).length(0))
                (d_486_v77_)[index72_] = default__.fm29((0) - (len((d_378_v5_) + (d_378_v5_))), (p2) + ((d_379_v6_)[default__.safeIndex(510, (d_379_v6_).length(0))]), d_382_v9_, default__.safeDivisionInt((d_379_v6_)[default__.safeIndex(510, (d_379_v6_).length(0))], (d_483_v73_).cf13), globalState)
                d_496_v78_: _dafny.Map
                d_496_v78_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gortewya")): not(d_479_cf4_)})
                d_497_v79_: D0
                d_497_v79_ = D0_DC1(d_479_cf4_, d_460_v59_, (d_460_v59_)[default__.safeIndex(d_438_v46_, len(d_460_v59_))])
                d_498_v80_: _dafny.Set
                d_498_v80_ = _dafny.Set({p3})
                d_496_v78_ = (d_496_v78_).set((d_497_v79_).cf3, ((d_498_v80_)).issubset(p0))
                d_384_v11_ = d_384_v11_
                d_499_v81_: _dafny.MultiSet
                d_499_v81_ = _dafny.MultiSet([D0_DC1(d_479_cf4_, _dafny.SeqWithoutIsStrInference([d_382_v9_ for d_500_i6_ in range(default__.abs(-952))]), _dafny.SeqWithoutIsStrInference([d_373_v0_ for d_501_i7_ in range(default__.abs(494))]))])
                d_502_v82_: _dafny.Seq
                d_502_v82_ = _dafny.SeqWithoutIsStrInference([d_499_v81_, d_499_v81_])
                d_503_v83_: _dafny.MultiSet
                d_503_v83_ = _dafny.MultiSet([(d_379_v6_)[default__.safeIndex(510, (d_379_v6_).length(0))]])
                d_504_v84_: _dafny.Seq
                d_504_v84_ = _dafny.SeqWithoutIsStrInference([(d_502_v82_)[default__.safeIndex((d_503_v83_).cardinality, len(d_502_v82_))], d_499_v81_])
                d_505_v85_: _dafny.Map
                d_505_v85_ = _dafny.Map({p1: d_502_v82_})
                d_505_v85_ = (d_505_v85_).set(((d_379_v6_)[default__.safeIndex(510, (d_379_v6_).length(0))]) + (d_438_v46_), d_502_v82_)
            elif source9_.is_DC0:
                d_506___mcc_h16_ = source9_.cf0
                d_507_cf0_ = d_506___mcc_h16_
                nw60_ = _dafny.Array(int(0), 3)
                d_384_v11_ = nw60_
                rhs81_ = d_373_v0_
                rhs82_ = (_dafny.MultiSet([d_373_v0_])) - ((d_374_v1_) | (d_374_v1_))
                rhs83_ = len(((d_464_v62_)[default__.safeIndex(753, (d_464_v62_).length(0))]) + ((d_464_v62_)[default__.safeIndex(753, (d_464_v62_).length(0))]))
                rhs84_ = (0) - (d_438_v46_)
                rhs85_ = (d_379_v6_)[default__.safeIndex(510, (d_379_v6_).length(0))]
                d_373_v0_ = rhs81_
                d_374_v1_ = rhs82_
                d_438_v46_ = rhs83_
                d_438_v46_ = rhs84_
                d_438_v46_ = rhs85_
                d_508_v86_: _dafny.Map
                d_508_v86_ = _dafny.Map({len(d_383_v10_): d_460_v59_})
                d_509_v87_: D0
                d_509_v87_ = D0_DC1(p3, _dafny.SeqWithoutIsStrInference([d_382_v9_]), (self).fm1(_dafny.Map({d_438_v46_: len(d_381_v8_)}), globalState))
                d_510_v88_: _dafny.MultiSet
                d_510_v88_ = _dafny.MultiSet([D0_DC1(p3, _dafny.SeqWithoutIsStrInference([d_382_v9_ for d_511_i8_ in range(default__.abs(638))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cv"))), D0_DC1(d_457_cf4_, d_460_v59_, (d_464_v62_)[default__.safeIndex(753, (d_464_v62_).length(0))]), default__.fm17(p3, globalState), D0_DC1(p3, ((d_508_v86_)[d_438_v46_] if (d_438_v46_) in (d_508_v86_) else _dafny.SeqWithoutIsStrInference([(d_464_v62_)[default__.safeIndex(753, (d_464_v62_).length(0))] for d_512_i9_ in range(default__.abs(105))])), d_382_v9_), d_509_v87_])
                d_513_v89_: C0
                nw61_ = C0()
                nw61_.ctor__(d_510_v88_, d_457_cf4_)
                d_513_v89_ = nw61_
                d_514_v90_: _dafny.Seq
                d_514_v90_ = _dafny.SeqWithoutIsStrInference([default__.fm17((d_513_v89_).f19, globalState)])
                d_515_v91_: C0
                nw62_ = C0()
                nw62_.ctor__((_dafny.MultiSet(d_514_v90_)).set(default__.fm17((d_513_v89_).f19, globalState), default__.abs(d_438_v46_)), d_457_cf4_)
                d_515_v91_ = nw62_
            elif True:
                d_516___mcc_h17_ = source9_.cf5
                d_517_cf5_ = d_516___mcc_h17_
                d_518_v92_: _dafny.Map
                d_518_v92_ = _dafny.Map({p3: default__.fm19(len(d_462_v60_), p3, globalState)})
                d_519_v93_: D0
                d_519_v93_ = D0_DC1(d_457_cf4_, d_460_v59_, d_382_v9_)
                d_520_v94_: _dafny.Map
                d_520_v94_ = _dafny.Map({d_457_cf4_: D0_DC3(d_519_v93_)})
                d_521_v95_: D2
                d_521_v95_ = D2_DC7(((d_383_v10_)[(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(d_379_v6_)[default__.safeIndex(510, (d_379_v6_).length(0))] for d_522_i10_ in range(default__.abs(890))]))).cardinality] if ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(d_379_v6_)[default__.safeIndex(510, (d_379_v6_).length(0))] for d_523_i10_ in range(default__.abs(890))]))).cardinality) in (d_383_v10_) else p3), d_520_v94_, d_438_v46_, not(d_457_cf4_))
                d_524_v96_: _dafny.Seq
                d_524_v96_ = _dafny.SeqWithoutIsStrInference([d_521_v95_])
                d_525_v97_: _dafny.Map
                d_525_v97_ = _dafny.Map({p2: d_458_v58_})
                d_526_v98_: _dafny.Map
                d_526_v98_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference([len(d_518_v92_)])).set(default__.safeIndex(len(d_462_v60_), len(_dafny.SeqWithoutIsStrInference([len(d_518_v92_)]))), len(d_524_v96_)): len(d_525_v97_)})
                d_527_v99_: _dafny.Array
                nw63_ = _dafny.Array(_dafny.Array(None, 0), 1)
                d_527_v99_ = nw63_
                d_528_v100_: _dafny.Array
                nw64_ = _dafny.Array(_dafny.Array(None, 0), 11)
                d_528_v100_ = nw64_
                index73_ = default__.safeIndex(699, (d_527_v99_).length(0))
                (d_527_v99_)[index73_] = d_528_v100_
                index74_ = default__.safeIndex(699, (d_527_v99_).length(0))
                rhs86_ = d_526_v98_
                rhs87_ = d_528_v100_
                lhs60_ = d_527_v99_
                lhs61_ = default__.safeIndex(699, (d_527_v99_).length(0))
                d_526_v98_ = rhs86_
                lhs60_[lhs61_] = rhs87_
                d_529_v101_: _dafny.Map
                d_529_v101_ = _dafny.Map({(d_464_v62_)[default__.safeIndex(753, (d_464_v62_).length(0))]: d_462_v60_})
                d_530_v102_: _dafny.Array
                nw65_ = _dafny.Array(None, 10)
                nw65_[int(0)] = (d_462_v60_).set(default__.safeIndex(len((d_464_v62_)[default__.safeIndex(753, (d_464_v62_).length(0))]), len(d_462_v60_)), p1)
                nw65_[int(1)] = d_462_v60_
                nw65_[int(2)] = (d_462_v60_) + (d_462_v60_)
                nw65_[int(3)] = (d_462_v60_ if False else default__.fm30(globalState))
                nw65_[int(4)] = ((d_529_v101_)[default__.fm20(p1, p3, d_457_cf4_, d_438_v46_, globalState)] if (default__.fm20(p1, p3, d_457_cf4_, d_438_v46_, globalState)) in (d_529_v101_) else d_462_v60_)
                nw65_[int(5)] = _dafny.SeqWithoutIsStrInference([len(d_462_v60_) for d_531_i11_ in range(default__.abs(-514))])
                nw65_[int(6)] = d_462_v60_
                nw65_[int(7)] = d_462_v60_
                nw65_[int(8)] = d_462_v60_
                nw65_[int(9)] = (_dafny.SeqWithoutIsStrInference([(d_379_v6_)[default__.safeIndex(510, (d_379_v6_).length(0))]])) + (_dafny.SeqWithoutIsStrInference([-161 for d_532_i12_ in range(default__.abs(224))]))
                d_530_v102_ = nw65_
                (globalState).f2 = d_530_v102_
                d_533_v103_: D1
                d_533_v103_ = D1_DC5(not(p3), (d_464_v62_)[default__.safeIndex(753, (d_464_v62_).length(0))], d_378_v5_)
                rhs88_ = (((0) - (len(d_382_v9_))) + (p2)) * (p1)
                rhs89_ = (d_533_v103_).cf7
                rhs90_ = d_457_cf4_
                lhs62_ = globalState
                lhs63_ = globalState
                d_438_v46_ = rhs88_
                lhs62_.f13 = rhs89_
                lhs63_.f5 = rhs90_
                d_534_v104_: _dafny.Map
                d_534_v104_ = _dafny.Map({p3: d_384_v11_})
                d_535_v105_: D1
                d_535_v105_ = D1_DC4(d_534_v104_)
                d_536_v106_: _dafny.Map
                d_536_v106_ = _dafny.Map({d_462_v60_: d_535_v105_})
                d_537_v107_: _dafny.Map
                d_537_v107_ = _dafny.Map({d_536_v106_: p3})
                d_537_v107_ = (d_537_v107_).set(d_536_v106_, p3)
            d_538_v108_: _dafny.Array
            nw66_ = _dafny.Array(None, 8)
            d_538_v108_ = nw66_
            d_539_v109_: D0
            d_539_v109_ = D0_DC1(d_457_cf4_, d_460_v59_, (d_464_v62_)[default__.safeIndex(753, (d_464_v62_).length(0))])
            d_540_v110_: D0
            d_540_v110_ = D0_DC1(False, (d_539_v109_).cf2, (d_464_v62_)[default__.safeIndex(753, (d_464_v62_).length(0))])
            d_541_v111_: _dafny.MultiSet
            d_541_v111_ = _dafny.MultiSet([d_540_v110_])
            d_542_v112_: C0
            nw67_ = C0()
            nw67_.ctor__(d_541_v111_, (d_539_v109_).cf1)
            d_542_v112_ = nw67_
            d_543_v113_: _dafny.Map
            d_543_v113_ = _dafny.Map({d_538_v108_: d_542_v112_})
            d_544_v114_: _dafny.Map
            d_544_v114_ = _dafny.Map({d_543_v113_: p1})
            index75_ = default__.safeIndex(810, (d_458_v58_).length(0))
            (d_458_v58_)[index75_] = ((0) - ((0) - ((self).fm0(((d_544_v114_)[d_543_v113_] if (d_543_v113_) in (d_544_v114_) else -858), d_373_v0_, globalState)))) != (p1)
            index76_ = default__.safeIndex(810, (d_458_v58_).length(0))
            (d_458_v58_)[index76_] = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ooskv"))) + (default__.fm20(d_438_v46_, d_457_cf4_, (d_542_v112_).f19, d_438_v46_, globalState))) < (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xfq")))
        elif source8_.is_DC0:
            d_545___mcc_h10_ = source8_.cf0
            d_546_cf0_ = d_545___mcc_h10_
            if p3:
                d_547_v115_: _dafny.Seq
                d_547_v115_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_373_v0_ for d_548_i13_ in range(default__.abs(857))]), (d_382_v9_).set(default__.safeIndex(d_546_cf0_, len(d_382_v9_)), d_373_v0_)])
                d_549_v116_: D0
                d_549_v116_ = D0_DC1(p3, d_547_v115_, d_382_v9_)
                d_550_v117_: D0
                d_550_v117_ = D0_DC3(d_549_v116_)
                d_551_v118_: D0
                d_551_v118_ = D0_DC3(d_549_v116_)
                d_552_v119_: D2
                d_552_v119_ = D2_DC7(False, _dafny.Map({default__.fm19(len(d_439_v47_), p3, globalState): d_551_v118_}), p2, p3)
                d_382_v9_ = (self).fm1((_dafny.Map({d_438_v46_: (d_552_v119_).cf13})) | (default__.fm31(globalState)), globalState)
                d_553_v120_: D0
                d_553_v120_ = D0_DC1(p3, (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uert")) for d_554_i14_ in range(default__.abs(110))])).set(default__.safeIndex(d_438_v46_, len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uert")) for d_555_i14_ in range(default__.abs(110))]))), d_382_v9_), d_382_v9_)
                d_556_v121_: _dafny.MultiSet
                d_556_v121_ = _dafny.MultiSet([d_553_v120_])
                d_557_v122_: _dafny.Array
                nw68_ = _dafny.Array(False, 14)
                d_557_v122_ = nw68_
                d_558_v123_: _dafny.Map
                d_558_v123_ = _dafny.Map({d_557_v122_: (d_377_v4_).cardinality})
                d_559_v124_: C0
                nw69_ = C0()
                nw69_.ctor__((d_556_v121_).set(D0_DC1(p3, d_547_v115_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xbt"))), default__.abs(len(d_558_v123_))), p3)
                d_559_v124_ = nw69_
                d_560_v125_: bool
                d_561_v126_: int
                d_562_v127_: bool
                out5_: bool
                out6_: int
                out7_: bool
                out5_, out6_, out7_ = (self).m6(d_382_v9_, d_382_v9_, p3, 163, globalState)
                d_560_v125_ = out5_
                d_561_v126_ = out6_
                d_562_v127_ = out7_
                d_563_v128_: bool
                d_564_v129_: int
                d_565_v130_: bool
                out8_: bool
                out9_: int
                out10_: bool
                out8_, out9_, out10_ = (self).m6(d_382_v9_, d_382_v9_, d_560_v125_, p2, globalState)
                d_563_v128_ = out8_
                d_564_v129_ = out9_
                d_565_v130_ = out10_
                d_561_v126_ = (default__.safeModuloInt(d_546_cf0_, (0) - (d_546_cf0_))) + (d_546_cf0_)
            elif True:
                d_438_v46_ = p1
                d_566_v131_: _dafny.Array
                nw70_ = _dafny.Array(_dafny.MultiSet({}), 4)
                d_566_v131_ = nw70_
                rhs91_ = d_566_v131_
                rhs92_ = p3
                lhs64_ = globalState
                d_566_v131_ = rhs91_
                lhs64_.f13 = rhs92_
                (globalState).f3 = p3
                d_567_v132_: D1
                d_567_v132_ = D1_DC5(p3, d_382_v9_, d_378_v5_)
                d_568_v133_: _dafny.Map
                d_568_v133_ = _dafny.Map({d_567_v132_: d_382_v9_})
                d_568_v133_ = d_568_v133_
                d_569_v134_: _dafny.Set
                d_569_v134_ = p0
                d_570_v135_: _dafny.Seq
                d_570_v135_ = _dafny.SeqWithoutIsStrInference([d_382_v9_])
                d_571_v136_: D0
                d_571_v136_ = D0_DC1(p3, d_570_v135_, d_382_v9_)
                d_572_v137_: _dafny.MultiSet
                d_572_v137_ = _dafny.MultiSet([d_571_v136_])
                d_573_v138_: C0
                nw71_ = C0()
                nw71_.ctor__(d_572_v137_, ((d_383_v10_)[(d_377_v4_).cardinality] if ((d_377_v4_).cardinality) in (d_383_v10_) else p3))
                d_573_v138_ = nw71_
                d_574_v139_: _dafny.Seq
                d_574_v139_ = _dafny.SeqWithoutIsStrInference([d_573_v138_])
                rhs93_ = (p0) | (p0)
                rhs94_ = d_379_v6_
                rhs95_ = (d_574_v139_) + (d_574_v139_)
                d_569_v134_ = rhs93_
                d_384_v11_ = rhs94_
                d_574_v139_ = rhs95_
            d_575_v140_: _dafny.Map
            d_575_v140_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "astkcoj")): p3})
            d_576_v142_: _dafny.Map
            d_576_v142_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v")): len(d_378_v5_)})
            def iife41_():
                coll15_ = _dafny.Map()
                compr_15_: _dafny.Seq
                for compr_15_ in ((d_576_v142_).set(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rjqeb")), len(d_382_v9_))).keys.Elements:
                    d_577_v141_: _dafny.Seq = compr_15_
                    if (d_577_v141_) in ((d_576_v142_).set(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rjqeb")), len(d_382_v9_))):
                        coll15_[d_577_v141_] = p3
                return _dafny.Map(coll15_)
            d_575_v140_ = iife41_()
            
            d_578_v143_: _dafny.Map
            d_578_v143_ = _dafny.Map({p3: p3})
            d_579_v144_: _dafny.Set
            d_579_v144_ = _dafny.Set({((d_578_v143_)[(self).fm14(_dafny.Set({p1}), p3, d_382_v9_, globalState)] if ((self).fm14(_dafny.Set({p1}), p3, d_382_v9_, globalState)) in (d_578_v143_) else p3), p3})
            d_579_v144_ = p0
            d_580_v145_: _dafny.Seq
            d_580_v145_ = _dafny.SeqWithoutIsStrInference([d_382_v9_, d_382_v9_])
            d_581_v146_: D0
            d_581_v146_ = D0_DC0(len((D0_DC1(not(p3), (d_580_v145_).set(default__.safeIndex(p1, len(d_580_v145_)), d_382_v9_), d_382_v9_)).cf3))
            d_582_v147_: D0
            d_582_v147_ = D0_DC3(d_581_v146_)
            d_582_v147_ = d_582_v147_
        elif True:
            d_583___mcc_h11_ = source8_.cf5
            d_584_cf5_ = d_583___mcc_h11_
            r2 = (d_382_v9_) < ((d_382_v9_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ienvw"))))
            d_383_v10_ = (d_383_v10_).set(default__.safeModuloInt(p2, -748), (True) and (p3))
            d_585_v148_: _dafny.Array
            nw72_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 2)
            d_585_v148_ = nw72_
            d_586_v149_: _dafny.Map
            d_586_v149_ = _dafny.Map({p3: d_382_v9_})
            index77_ = default__.safeIndex(837, (d_585_v148_).length(0))
            (d_585_v148_)[index77_] = (d_382_v9_) + (((d_586_v149_)[True] if (True) in (d_586_v149_) else d_382_v9_))
            index78_ = default__.safeIndex(837, (d_585_v148_).length(0))
            (d_585_v148_)[index78_] = default__.fm20(p2, p3, (d_382_v9_) < (_dafny.SeqWithoutIsStrInference([d_373_v0_ for d_587_i15_ in range(default__.abs(933))])), p1, globalState)
            (globalState).f3 = p3
        d_588_v150_: _dafny.Array
        def lambda14_(d_589_p3_, d_590_v9_):
            def lambda15_(d_591_i16_):
                return (D0_DC1(d_589_p3_, _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "usbo"))]), d_590_v9_)).cf1

            return lambda15_

        init6_ = lambda14_(p3, d_382_v9_)
        nw73_ = _dafny.Array(None, 29)
        for i0_6_ in range(nw73_.length(0)):
            nw73_[i0_6_] = init6_(i0_6_)
        d_588_v150_ = nw73_
        index79_ = default__.safeIndex(568, (d_588_v150_).length(0))
        (d_588_v150_)[index79_] = p3
        d_592_v151_: _dafny.Map
        d_592_v151_ = _dafny.Map({p3: p3})
        d_593_v152_: _dafny.Map
        d_593_v152_ = _dafny.Map({d_592_v151_: p3})
        index80_ = default__.safeIndex(436, (d_588_v150_).length(0))
        (d_588_v150_)[index80_] = (d_592_v151_) in (d_593_v152_)
        d_594_v153_: _dafny.MultiSet
        d_594_v153_ = _dafny.MultiSet([-21, (0) - ((0) - (d_438_v46_))])
        d_595_v154_: _dafny.Seq
        d_595_v154_ = _dafny.SeqWithoutIsStrInference([d_594_v153_])
        index81_ = default__.safeIndex(568, (d_588_v150_).length(0))
        index82_ = default__.safeIndex(436, (d_588_v150_).length(0))
        rhs96_ = p3
        rhs97_ = (self).fm14(d_439_v47_, p3, d_382_v9_, globalState)
        rhs98_ = (d_595_v154_) != ((_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([len(d_381_v8_), p2, p1]), d_594_v153_]) if p3 else d_595_v154_))
        rhs99_ = d_438_v46_
        lhs65_ = d_588_v150_
        lhs66_ = default__.safeIndex(568, (d_588_v150_).length(0))
        lhs67_ = d_588_v150_
        lhs68_ = default__.safeIndex(436, (d_588_v150_).length(0))
        lhs69_ = globalState
        lhs65_[lhs66_] = rhs96_
        lhs67_[lhs68_] = rhs97_
        lhs69_.f5 = rhs98_
        d_438_v46_ = rhs99_
        d_596_v155_: D0
        d_596_v155_ = D0_DC0(p2)
        d_597_v156_: D0
        d_597_v156_ = D0_DC3(D0_DC3(d_596_v155_))
        d_598_v157_: _dafny.Map
        d_598_v157_ = _dafny.Map({p3: d_597_v156_})
        d_599_v158_: _dafny.Set
        d_599_v158_ = _dafny.Set({_dafny.Map({(0) - (p1): default__.fm22(D2_DC7((d_588_v150_)[default__.safeIndex(568, (d_588_v150_).length(0))], d_598_v157_, (_dafny.MultiSet([p3, False])).cardinality, p3), default__.fm28(globalState), (d_588_v150_)[default__.safeIndex(568, (d_588_v150_).length(0))], globalState)}), d_381_v8_})
        d_600_v159_: _dafny.Seq
        d_600_v159_ = _dafny.SeqWithoutIsStrInference([d_382_v9_, d_382_v9_])
        d_601_v160_: D0
        d_601_v160_ = D0_DC1((d_588_v150_)[default__.safeIndex(568, (d_588_v150_).length(0))], d_600_v159_, d_382_v9_)
        d_602_v161_: D2
        d_602_v161_ = D2_DC6(_dafny.MultiSet([d_601_v160_, D0_DC1((d_588_v150_)[default__.safeIndex(568, (d_588_v150_).length(0))], d_600_v159_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bfwqt"))), d_601_v160_]))
        pat_let_tv15_ = d_599_v158_
        pat_let_tv16_ = d_599_v158_
        pat_let_tv17_ = d_599_v158_
        def lambda16_(source10_):
            if source10_.is_DC7:
                d_603___mcc_h18_ = source10_.cf11
                d_604___mcc_h19_ = source10_.cf12
                d_605___mcc_h20_ = source10_.cf13
                d_606___mcc_h21_ = source10_.cf14
                d_607_cf14_ = d_606___mcc_h21_
                d_608_cf13_ = d_605___mcc_h20_
                d_609_cf12_ = d_604___mcc_h19_
                d_610_cf11_ = d_603___mcc_h18_
                return pat_let_tv15_
            elif source10_.is_DC6:
                d_611___mcc_h22_ = source10_.cf10
                d_612_cf10_ = d_611___mcc_h22_
                return pat_let_tv16_
            elif True:
                d_613___mcc_h23_ = source10_.cf15
                d_614_cf15_ = d_613___mcc_h23_
                return pat_let_tv17_

        d_599_v158_ = lambda16_(d_602_v161_)
        r0 = (d_378_v5_) + ((d_378_v5_) + (d_378_v5_))
        r1 = d_378_v5_
        r2 = (d_588_v150_)[default__.safeIndex(568, (d_588_v150_).length(0))]
        return r0, r1, r2


class C3(T1, T0):
    def  __init__(self):
        self._f16: bool = False
        self._f17: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    def ctor__(self, f16, f17):
        (self)._f16 = f16
        (self)._f17 = f17

    def fm0(self, p0, p1, globalState):
        return (0) - ((self).f17)

    def fm1(self, p0, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vcy"))

    def fm13(self, p0, p1, p2, p3, globalState):
        def iife42_():
            coll16_ = _dafny.Set()
            compr_16_: _dafny.Seq
            for compr_16_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(self).f16]) for d_615_i0_ in range(default__.abs(216))])).Elements:
                d_616_v0_: _dafny.Seq = compr_16_
                if (d_616_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(self).f16]) for d_615_i0_ in range(default__.abs(216))])):
                    coll16_ = coll16_.union(_dafny.Set([d_616_v0_]))
            return _dafny.Set(coll16_)
        return ((_dafny.Set({_dafny.SeqWithoutIsStrInference([(self).f16, (self).f16]), _dafny.SeqWithoutIsStrInference([(self).f16]), _dafny.SeqWithoutIsStrInference([(self).f16])})).intersection(_dafny.Set({_dafny.SeqWithoutIsStrInference([(self).f16]), _dafny.SeqWithoutIsStrInference([(self).f16, (self).f16, (self).f16, (self).f16, (self).f16])}))).isdisjoint((iife42_()
        ).intersection(_dafny.Set({_dafny.SeqWithoutIsStrInference([(self).f16])})))

    def m0(self, p0, globalState):
        r0: bool = False
        r1: int = int(0)
        d_617_v0_: C2
        nw74_ = C2()
        nw74_.ctor__()
        d_617_v0_ = nw74_
        d_618_v1_: _dafny.Array
        nw75_ = _dafny.Array(False, 26)
        d_618_v1_ = nw75_
        d_618_v1_ = d_618_v1_
        r1 = (self).f17
        (globalState).f3 = ((self).f16) or ((self).f16)
        d_619_v2_: _dafny.Array
        nw76_ = _dafny.Array(int(0), 5)
        d_619_v2_ = nw76_
        d_620_v3_: _dafny.Map
        d_620_v3_ = _dafny.Map({d_619_v2_: (self).f16})
        d_621_v4_: _dafny.Seq
        d_621_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eehtvwmi"))
        d_622_v5_: _dafny.Seq
        d_622_v5_ = _dafny.SeqWithoutIsStrInference([(self).f16])
        d_623_v6_: D1
        d_623_v6_ = D1_DC5((self).f16, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dr")), d_622_v5_)
        d_624_v7_: _dafny.Map
        d_624_v7_ = _dafny.Map({d_620_v3_: len((d_621_v4_ if (self).f16 else (d_623_v6_).cf8))})
        d_624_v7_ = (d_624_v7_).set(d_620_v3_, (self).fm0((self).f17, p0, globalState))
        d_625_v8_: C1
        nw77_ = C1()
        nw77_.ctor__()
        d_625_v8_ = nw77_
        r0 = (self).f16
        d_626_v9_: _dafny.Seq
        d_626_v9_ = _dafny.SeqWithoutIsStrInference([d_621_v4_, _dafny.SeqWithoutIsStrInference([p0 for d_627_i0_ in range(default__.abs(828))])])
        d_628_v10_: D0
        d_628_v10_ = D0_DC1(True, d_626_v9_, d_621_v4_)
        d_629_v11_: _dafny.MultiSet
        def iife43_(_pat_let13_0):
            def iife44_(d_630_dt__update__tmp_h0_):
                def iife45_(_pat_let14_0):
                    def iife46_(d_631_dt__update_hcf1_h0_):
                        return D0_DC1(d_631_dt__update_hcf1_h0_, (d_630_dt__update__tmp_h0_).cf2, (d_630_dt__update__tmp_h0_).cf3)
                    return iife46_(_pat_let14_0)
                return iife45_((self).f16)
            return iife44_(_pat_let13_0)
        d_629_v11_ = _dafny.MultiSet([(iife43_(d_628_v10_)).cf3])
        r1 = ((self).f17) + ((d_629_v11_).cardinality)
        return r0, r1

    @property
    def f16(self):
        return self._f16
    @property
    def f17(self):
        return self._f17

class C4:
    def  __init__(self):
        self._f14: int = int(0)
        self._f15: _dafny.Map = _dafny.Map({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    def ctor__(self, f14, f15):
        (self)._f14 = f14
        (self)._f15 = f15

    def fm7(self, p0, p1, p2, p3, globalState):
        def iife47_():
            coll17_ = _dafny.Map()
            compr_17_: int
            for compr_17_ in (_dafny.Set({(self).f14})).Elements:
                d_632_v0_: int = compr_17_
                if (d_632_v0_) in (_dafny.Set({(self).f14})):
                    coll17_[default__.safeModuloInt(d_632_v0_, -958)] = False
            return _dafny.Map(coll17_)
        return (_dafny.SeqWithoutIsStrInference([(D0_DC0((0) - ((self).f14))).cf0])) + ((_dafny.SeqWithoutIsStrInference([len(iife47_()
        ), 832, (self).f14])) + (_dafny.SeqWithoutIsStrInference([(self).f14])))

    def fm8(self, p0, p1, p2, globalState):
        def iife48_():
            coll18_ = _dafny.Map()
            compr_18_: int
            for compr_18_ in (_dafny.Set({(self).f14})).Elements:
                d_633_v0_: int = compr_18_
                if (d_633_v0_) in (_dafny.Set({(self).f14})):
                    coll18_[default__.safeDivisionInt(d_633_v0_, (self).f14)] = True
            return _dafny.Map(coll18_)
        return not ((_dafny.Set({(self).f14})).isdisjoint(_dafny.Set({(self).f14, len(iife48_()
        ), (0) - ((self).f14), (_dafny.MultiSet([(self).f14])).cardinality}))) or ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "axmcm"))) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "elfoqgqbn"))))

    def m5(self, p0, globalState):
        d_634_v0_: bool
        d_634_v0_ = False
        d_635_i0_: int
        d_635_i0_ = 0
        with _dafny.label("3"):
            while d_634_v0_:
                with _dafny.c_label("3"):
                    if (d_635_i0_) >= (100):
                        raise _dafny.Break("3")
                    d_635_i0_ = (d_635_i0_) + (1)
                    d_636_v1_: D0
                    d_636_v1_ = D0_DC0(default__.safeDivisionInt((self).f14, (self).f14))
                    d_637_v2_: _dafny.Map
                    d_637_v2_ = _dafny.Map({(self).f14: not(d_634_v0_)})
                    d_638_v3_: _dafny.Seq
                    d_638_v3_ = _dafny.SeqWithoutIsStrInference([((d_637_v2_)[(self).f14] if ((self).f14) in (d_637_v2_) else d_634_v0_), d_634_v0_])
                    pat_let_tv18_ = d_638_v3_
                    def iife49_(_pat_let15_0):
                        def iife50_(d_639_dt__update__tmp_h0_):
                            def iife51_(_pat_let16_0):
                                def iife52_(d_640_dt__update_hcf0_h0_):
                                    return D0_DC0(d_640_dt__update_hcf0_h0_)
                                return iife52_(_pat_let16_0)
                            return iife51_(len(pat_let_tv18_))
                        return iife50_(_pat_let15_0)
                    d_636_v1_ = iife49_(d_636_v1_)
                    d_641_v4_: int
                    d_641_v4_ = 665
                    d_641_v4_ = (self).f14
                    d_642_v5_: _dafny.Map
                    d_642_v5_ = _dafny.Map({d_634_v0_: d_634_v0_})
                    if (d_634_v0_ if not((_dafny.Set({(self).f14})).isdisjoint(_dafny.Set({(0) - (d_641_v4_), d_641_v4_, len(d_642_v5_), (self).f14}))) else d_634_v0_):
                        d_643_v6_: _dafny.Map
                        d_643_v6_ = _dafny.Map({d_634_v0_: d_641_v4_})
                        d_644_v7_: _dafny.MultiSet
                        d_644_v7_ = _dafny.MultiSet([(self).f14, d_641_v4_])
                        d_645_v8_: _dafny.Map
                        d_645_v8_ = _dafny.Map({_dafny.Set({d_634_v0_, d_634_v0_, d_634_v0_}): (self).f14})
                        d_646_v9_: _dafny.Map
                        d_646_v9_ = _dafny.Map({d_634_v0_: d_645_v8_})
                        d_647_v10_: _dafny.Array
                        nw78_ = _dafny.Array(None, 26)
                        nw78_[int(0)] = ((d_643_v6_)[d_634_v0_] if (d_634_v0_) in (d_643_v6_) else (self).f14)
                        nw78_[int(1)] = (self).f14
                        nw78_[int(2)] = ((d_644_v7_).cardinality) - ((self).f14)
                        nw78_[int(3)] = (self).f14
                        nw78_[int(4)] = (d_636_v1_).cf0
                        nw78_[int(5)] = (self).f14
                        nw78_[int(6)] = 321
                        nw78_[int(7)] = (self).f14
                        nw78_[int(8)] = len(((d_646_v9_)[False] if (False) in (d_646_v9_) else d_645_v8_))
                        nw78_[int(9)] = (self).f14
                        nw78_[int(10)] = default__.safeModuloInt(d_641_v4_, 841)
                        nw78_[int(11)] = (self).f14
                        nw78_[int(12)] = (self).f14
                        nw78_[int(13)] = len((default__.fm9(-178, d_638_v3_, p0, globalState)) + (d_638_v3_))
                        nw78_[int(14)] = d_641_v4_
                        nw78_[int(15)] = (self).f14
                        nw78_[int(16)] = len(default__.fm10(d_634_v0_, 569, (self).fm8(d_641_v4_, d_634_v0_, d_634_v0_, globalState), not(d_634_v0_), globalState))
                        nw78_[int(17)] = ((d_644_v7_).cardinality) - ((0) - (default__.fm11(d_634_v0_, globalState)))
                        nw78_[int(18)] = d_641_v4_
                        nw78_[int(19)] = default__.safeDivisionInt((0) - (d_641_v4_), d_641_v4_)
                        nw78_[int(20)] = d_641_v4_
                        nw78_[int(21)] = (self).f14
                        nw78_[int(22)] = (0) - ((d_641_v4_) - (d_641_v4_))
                        nw78_[int(23)] = ((self).f14) + ((self).f14)
                        nw78_[int(24)] = (0) - (d_641_v4_)
                        nw78_[int(25)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sdnbh")))
                        d_647_v10_ = nw78_
                        index83_ = default__.safeIndex(723, (d_647_v10_).length(0))
                        (d_647_v10_)[index83_] = d_641_v4_
                        d_648_v11_: _dafny.MultiSet
                        d_648_v11_ = _dafny.MultiSet([default__.fm12(d_634_v0_, d_634_v0_, (self).f14, globalState)])
                        index84_ = default__.safeIndex(723, (d_647_v10_).length(0))
                        def iife53_():
                            coll19_ = _dafny.Set()
                            compr_19_: D0
                            for compr_19_ in (d_648_v11_).Elements:
                                d_649_v12_: D0 = compr_19_
                                if (d_649_v12_) in (d_648_v11_):
                                    coll19_ = coll19_.union(_dafny.Set([d_649_v12_]))
                            return _dafny.Set(coll19_)
                        (d_647_v10_)[index84_] = len(iife53_()
                        )
                        (globalState).f13 = d_634_v0_
                        (globalState).f5 = d_634_v0_
                        d_650_v13_: C2
                        nw79_ = C2()
                        nw79_.ctor__()
                        d_650_v13_ = nw79_
                        d_651_v14_: _dafny.Set
                        d_651_v14_ = _dafny.Set({d_634_v0_, (d_643_v6_) == (default__.fm15(d_634_v0_, d_634_v0_, d_634_v0_, globalState))})
                        d_652_v15_: D1
                        d_652_v15_ = D1_DC5(d_634_v0_, p0, _dafny.SeqWithoutIsStrInference([True]))
                        d_653_v16_: _dafny.Array
                        def lambda17_(d_654_v0_):
                            def lambda18_(d_655_i1_):
                                return d_654_v0_

                            return lambda18_

                        init7_ = lambda17_(d_634_v0_)
                        nw80_ = _dafny.Array(None, 24)
                        for i0_7_ in range(nw80_.length(0)):
                            nw80_[i0_7_] = init7_(i0_7_)
                        d_653_v16_ = nw80_
                        index85_ = default__.safeIndex(347, (d_653_v16_).length(0))
                        (d_653_v16_)[index85_] = (d_638_v3_)[default__.safeIndex(d_641_v4_, len(d_638_v3_))]
                        d_656_v17_: _dafny.Set
                        d_656_v17_ = _dafny.Set({(self).f14, d_641_v4_, (self).f14, (0) - ((self).f14)})
                        index86_ = default__.safeIndex(347, (d_653_v16_).length(0))
                        rhs100_ = d_651_v14_
                        rhs101_ = d_652_v15_
                        rhs102_ = (d_650_v13_).fm14(d_656_v17_, d_634_v0_, p0, globalState)
                        lhs70_ = d_653_v16_
                        lhs71_ = default__.safeIndex(347, (d_653_v16_).length(0))
                        d_651_v14_ = rhs100_
                        d_652_v15_ = rhs101_
                        lhs70_[lhs71_] = rhs102_
                    elif True:
                        d_657_v18_: C2
                        nw81_ = C2()
                        nw81_.ctor__()
                        d_657_v18_ = nw81_
                        d_657_v18_ = d_657_v18_
                        d_658_v19_: _dafny.Array
                        nw82_ = _dafny.Array(False, 17)
                        d_658_v19_ = nw82_
                        index87_ = default__.safeIndex(996, (d_658_v19_).length(0))
                        (d_658_v19_)[index87_] = d_634_v0_
                        d_659_v20_: _dafny.Seq
                        d_659_v20_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rtdk"))
                        d_660_v21_: _dafny.Seq
                        d_660_v21_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([d_658_v19_, d_658_v19_])), (self).f14, (self).f14, d_641_v4_, len(d_638_v3_)])
                        index88_ = default__.safeIndex(996, (d_658_v19_).length(0))
                        rhs103_ = (d_660_v21_) < (d_660_v21_)
                        rhs104_ = p0
                        lhs72_ = d_658_v19_
                        lhs73_ = default__.safeIndex(996, (d_658_v19_).length(0))
                        lhs72_[lhs73_] = rhs103_
                        d_659_v20_ = rhs104_
                        d_661_v23_: _dafny.MultiSet
                        def iife54_():
                            coll20_ = _dafny.Set()
                            compr_20_: int
                            for compr_20_ in _dafny.IntegerRange(720, -698):
                                d_662_v22_: int = compr_20_
                                if ((720) <= (d_662_v22_)) and ((d_662_v22_) < (-698)):
                                    coll20_ = coll20_.union(_dafny.Set([default__.safeDivisionInt(d_662_v22_, (self).f14)]))
                            return _dafny.Set(coll20_)
                        d_661_v23_ = _dafny.MultiSet([len(iife54_()
                        )])
                        d_661_v23_ = d_661_v23_
                        d_663_v24_: _dafny.Seq
                        d_663_v24_ = _dafny.SeqWithoutIsStrInference([d_659_v20_])
                        (globalState).f13 = not((d_663_v24_) == (d_663_v24_))
                        d_664_v25_: _dafny.Map
                        d_664_v25_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([(d_658_v19_)[default__.safeIndex(996, (d_658_v19_).length(0))], (d_658_v19_)[default__.safeIndex(996, (d_658_v19_).length(0))], (d_658_v19_)[default__.safeIndex(996, (d_658_v19_).length(0))]]): d_658_v19_})
                        d_658_v19_ = ((d_664_v25_)[d_638_v3_] if (d_638_v3_) in (d_664_v25_) else d_658_v19_)
                    d_665_v26_: _dafny.Seq
                    d_665_v26_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y"))
                    d_666_v27_: _dafny.Set
                    d_666_v27_ = _dafny.Set({(0) - (d_641_v4_), (self).f14})
                    d_667_v28_: _dafny.Map
                    d_667_v28_ = _dafny.Map({d_634_v0_: d_641_v4_})
                    d_668_v29_: _dafny.Map
                    d_668_v29_ = _dafny.Map({d_666_v27_: len(d_667_v28_)})
                    def iife55_():
                        coll21_ = _dafny.Map()
                        compr_21_: _dafny.Set
                        for compr_21_ in (d_668_v29_).keys.Elements:
                            d_669_v30_: _dafny.Set = compr_21_
                            if (d_669_v30_) in (d_668_v29_):
                                coll21_[d_669_v30_] = 615
                        return _dafny.Map(coll21_)
                    rhs105_ = ((d_665_v26_) + (d_665_v26_)) + (d_665_v26_)
                    rhs106_ = (d_668_v29_) | ((d_668_v29_) | (iife55_()
                    ))
                    d_665_v26_ = rhs105_
                    d_668_v29_ = rhs106_
                    pass
            pass
        d_670_v31_: _dafny.Array
        nw83_ = _dafny.Array(_dafny.Map({}), 21)
        d_670_v31_ = nw83_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_670_v31_).length(0)):
            d_671_i2_: int = guard_loop_0_
            if (True) and (((0) <= (d_671_i2_)) and ((d_671_i2_) < ((d_670_v31_).length(0)))):
                (d_670_v31_)[(d_671_i2_)] = (_dafny.Map({213: 746})) | ((_dafny.Map({469: (self).f14})) | (_dafny.Map({len(p0): (self).f14})))
        hi5_ = (self).f14
        for d_672_i3_ in range(len(default__.fm32(True, d_634_v0_, p0, d_634_v0_, globalState)), hi5_):
            d_673_v32_: _dafny.Seq
            d_673_v32_ = _dafny.SeqWithoutIsStrInference([d_634_v0_, (p0) <= (p0)])
            if (d_673_v32_)[default__.safeIndex(760, len(d_673_v32_))]:
                d_674_v33_: int
                d_674_v33_ = 412
                d_674_v33_ = 491
                d_675_v34_: _dafny.Array
                nw84_ = _dafny.Array(int(0), 1)
                d_675_v34_ = nw84_
                index89_ = default__.safeIndex(31, (d_675_v34_).length(0))
                (d_675_v34_)[index89_] = d_672_i3_
                index90_ = default__.safeIndex(31, (d_675_v34_).length(0))
                (d_675_v34_)[index90_] = d_672_i3_
                d_676_v35_: _dafny.Set
                d_676_v35_ = _dafny.Set({default__.fm11(False, globalState)})
                d_677_v36_: D4
                d_677_v36_ = D4_DC10(d_676_v35_)
                d_678_v37_: _dafny.Map
                d_678_v37_ = _dafny.Map({not((d_676_v35_).isdisjoint((d_677_v36_).cf17)): d_672_i3_})
                d_678_v37_ = (d_678_v37_).set(d_634_v0_, d_674_v33_)
                d_679_v38_: C2
                nw85_ = C2()
                nw85_.ctor__()
                d_679_v38_ = nw85_
                d_674_v33_ = d_672_i3_
            elif True:
                d_680_v39_: int
                d_680_v39_ = 528
                d_680_v39_ = (0) - ((((self).f14) - (d_680_v39_)) * (d_672_i3_))
                (globalState).f5 = ((self).f14) >= (d_680_v39_)
                d_681_v40_: str
                d_681_v40_ = _dafny.CodePoint('j')
                d_682_v41_: _dafny.MultiSet
                d_682_v41_ = _dafny.MultiSet([d_681_v40_])
                d_680_v39_ = ((d_682_v41_) | (_dafny.MultiSet(p0))).cardinality
                d_683_v42_: _dafny.Array
                def lambda19_(d_684_v0_):
                    def lambda20_(d_685_i4_):
                        return d_684_v0_

                    return lambda20_

                init8_ = lambda19_(d_634_v0_)
                nw86_ = _dafny.Array(None, 27)
                for i0_8_ in range(nw86_.length(0)):
                    nw86_[i0_8_] = init8_(i0_8_)
                d_683_v42_ = nw86_
                d_686_v43_: D5
                d_686_v43_ = D5_DC13(d_683_v42_)
                d_683_v42_ = (d_686_v43_).cf19
                d_687_v44_: D0
                d_687_v44_ = D0_DC2(d_634_v0_)
                d_688_v45_: _dafny.Seq
                d_688_v45_ = _dafny.SeqWithoutIsStrInference([p0, p0])
                d_689_v46_: D0
                d_689_v46_ = D0_DC1(default__.fm19(d_672_i3_, (d_687_v44_).cf4, globalState), d_688_v45_, p0)
                d_690_v47_: _dafny.MultiSet
                d_690_v47_ = _dafny.MultiSet([d_689_v46_])
                d_691_v48_: _dafny.Seq
                d_691_v48_ = _dafny.SeqWithoutIsStrInference([d_690_v47_])
                d_692_v49_: C0
                nw87_ = C0()
                nw87_.ctor__((d_691_v48_)[default__.safeIndex(d_680_v39_, len(d_691_v48_))], False)
                d_692_v49_ = nw87_
            d_693_v50_: _dafny.Array
            nw88_ = _dafny.Array(_dafny.Array(None, 0), 5)
            d_693_v50_ = nw88_
            d_694_v51_: _dafny.Array
            nw89_ = _dafny.Array(int(0), 16)
            d_694_v51_ = nw89_
            index91_ = default__.safeIndex(135, (d_693_v50_).length(0))
            (d_693_v50_)[index91_] = d_694_v51_
            d_695_v53_: _dafny.Map
            def iife56_():
                coll22_ = _dafny.Set()
                compr_22_: int
                for compr_22_ in (_dafny.Map({(self).f14: p0})).keys.Elements:
                    d_696_v52_: int = compr_22_
                    if (d_696_v52_) in (_dafny.Map({(self).f14: p0})):
                        coll22_ = coll22_.union(_dafny.Set([default__.safeDivisionInt(d_696_v52_, 275)]))
                return _dafny.Set(coll22_)
            d_695_v53_ = _dafny.Map({len(iife56_()
            ): 332})
            d_697_v54_: _dafny.MultiSet
            d_697_v54_ = _dafny.MultiSet([d_672_i3_, (self).f14, len(d_695_v53_), d_672_i3_])
            index92_ = default__.safeIndex(135, (d_693_v50_).length(0))
            (d_693_v50_)[index92_] = (d_694_v51_ if (d_697_v54_).issubset(_dafny.MultiSet([d_672_i3_])) else d_694_v51_)
            d_695_v53_ = (d_695_v53_).set(((self).f14) - ((self).f14), default__.safeModuloInt(d_672_i3_, (self).f14))
            d_698_v55_: int
            d_698_v55_ = 977
            d_699_v56_: str
            d_699_v56_ = _dafny.CodePoint('x')
            d_698_v55_ = len((p0).set(default__.safeIndex((902) + (d_672_i3_), len(p0)), d_699_v56_))
        d_700_v57_: str
        d_700_v57_ = _dafny.CodePoint('a')
        d_701_v58_: _dafny.MultiSet
        d_701_v58_ = _dafny.MultiSet([d_700_v57_])
        d_702_v59_: _dafny.Seq
        d_702_v59_ = _dafny.SeqWithoutIsStrInference([d_634_v0_])
        d_703_v60_: int
        d_703_v60_ = 976
        d_704_v61_: _dafny.Seq
        d_704_v61_ = _dafny.SeqWithoutIsStrInference([(self).f14, (self).f14])
        d_705_v62_: _dafny.Map
        d_705_v62_ = _dafny.Map({d_634_v0_: d_703_v60_})
        rhs107_ = not((-758) == (default__.fm11(not(d_634_v0_), globalState)))
        rhs108_ = ((d_701_v58_).intersection(_dafny.MultiSet([d_700_v57_]))) - (d_701_v58_)
        rhs109_ = default__.fm27((self).f14, d_703_v60_, default__.safeModuloInt(len(d_704_v61_), len(d_704_v61_)), globalState)
        rhs110_ = d_634_v0_
        rhs111_ = (len(d_705_v62_)) * ((self).f14)
        lhs74_ = globalState
        lhs75_ = globalState
        lhs74_.f13 = rhs107_
        d_701_v58_ = rhs108_
        d_702_v59_ = rhs109_
        lhs75_.f13 = rhs110_
        d_703_v60_ = rhs111_
        d_706_v63_: D0
        d_706_v63_ = D0_DC1(d_634_v0_, _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_700_v57_ for d_707_i6_ in range(default__.abs(104))]) for d_708_i5_ in range(default__.abs(210))]), p0)
        d_709_v64_: _dafny.MultiSet
        d_709_v64_ = _dafny.MultiSet([d_706_v63_, d_706_v63_, D0_DC1(d_634_v0_, _dafny.SeqWithoutIsStrInference([p0 for d_710_i7_ in range(default__.abs(67))]), p0)])
        d_711_v65_: D2
        d_711_v65_ = D2_DC6(d_709_v64_)
        pat_let_tv19_ = d_709_v64_
        pat_let_tv20_ = d_709_v64_
        def iife57_(_pat_let17_0):
            def iife58_(d_712_dt__update__tmp_h1_):
                def iife59_(_pat_let18_0):
                    def iife60_(d_713_dt__update_hcf10_h0_):
                        return D2_DC6(d_713_dt__update_hcf10_h0_)
                    return iife60_(_pat_let18_0)
                return iife59_((pat_let_tv19_) | (pat_let_tv20_))
            return iife58_(_pat_let17_0)
        source11_ = iife57_(d_711_v65_)
        if source11_.is_DC7:
            d_714___mcc_h0_ = source11_.cf11
            d_715___mcc_h1_ = source11_.cf12
            d_716___mcc_h2_ = source11_.cf13
            d_717___mcc_h3_ = source11_.cf14
            d_718_cf14_ = d_717___mcc_h3_
            d_719_cf13_ = d_716___mcc_h2_
            d_720_cf12_ = d_715___mcc_h1_
            d_721_cf11_ = d_714___mcc_h0_
            d_722_v66_: _dafny.Seq
            d_722_v66_ = _dafny.SeqWithoutIsStrInference([d_706_v63_])
            d_723_v67_: C0
            nw90_ = C0()
            nw90_.ctor__((d_709_v64_) | ((_dafny.MultiSet(d_722_v66_)).set(d_706_v63_, default__.abs(default__.fm11(d_721_cf11_, globalState)))), d_721_cf11_)
            d_723_v67_ = nw90_
            d_703_v60_ = (self).f14
            d_724_v68_: _dafny.Map
            d_724_v68_ = _dafny.Map({d_719_cf13_: default__.safeDivisionInt((0) - (d_703_v60_), -627)})
            d_724_v68_ = (d_724_v68_).set(d_703_v60_, d_703_v60_)
            d_725_v69_: D2
            d_725_v69_ = D2_DC7(d_634_v0_, d_720_cf12_, d_703_v60_, d_634_v0_)
            d_726_v70_: D0
            d_726_v70_ = D0_DC2(not((d_723_v67_).f19))
            d_727_v71_: D0
            d_727_v71_ = D0_DC3(D0_DC3(d_726_v70_))
            d_728_v72_: _dafny.Seq
            d_728_v72_ = _dafny.SeqWithoutIsStrInference([(self).fm7(p0, d_718_cf14_, d_721_cf11_, _dafny.Map({d_718_cf14_: d_721_cf11_}), globalState)])
            d_729_v73_: _dafny.Array
            nw91_ = _dafny.Array(None, 20)
            nw91_[int(0)] = len(default__.fm10(d_718_cf14_, d_719_cf13_, d_634_v0_, (d_723_v67_).f19, globalState))
            nw91_[int(1)] = d_703_v60_
            nw91_[int(2)] = d_703_v60_
            nw91_[int(3)] = d_703_v60_
            nw91_[int(4)] = (0) - ((d_719_cf13_ if d_718_cf14_ else d_703_v60_))
            nw91_[int(5)] = d_719_cf13_
            nw91_[int(6)] = (d_719_cf13_) * (len(d_704_v61_))
            nw91_[int(7)] = default__.fm22(d_725_v69_, d_727_v71_, False, globalState)
            nw91_[int(8)] = len((p0) + (_dafny.SeqWithoutIsStrInference([d_700_v57_ for d_730_i8_ in range(default__.abs(405))])))
            nw91_[int(9)] = default__.safeModuloInt((self).f14, d_719_cf13_)
            nw91_[int(10)] = len(d_728_v72_)
            nw91_[int(11)] = (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sxb")))) - (len(_dafny.Map({len(d_702_v59_): (d_723_v67_).f19})))
            nw91_[int(12)] = d_703_v60_
            nw91_[int(13)] = (d_703_v60_) - ((self).f14)
            nw91_[int(14)] = d_703_v60_
            nw91_[int(15)] = 884
            nw91_[int(16)] = len(d_724_v68_)
            nw91_[int(17)] = 945
            nw91_[int(18)] = default__.safeDivisionInt(d_719_cf13_, d_703_v60_)
            nw91_[int(19)] = (self).f14
            d_729_v73_ = nw91_
            index93_ = default__.safeIndex(368, (d_729_v73_).length(0))
            (d_729_v73_)[index93_] = (self).f14
            d_731_v74_: D1
            d_731_v74_ = D1_DC4(_dafny.Map({d_718_cf14_: d_729_v73_}))
            d_732_v75_: _dafny.Map
            d_732_v75_ = _dafny.Map({d_718_cf14_: d_729_v73_})
            d_733_v76_: D6
            d_733_v76_ = D6_DC15(_dafny.MultiSet([d_721_cf11_, (d_723_v67_).f19]))
            d_734_v77_: D4
            d_734_v77_ = D4_DC12(((d_733_v76_).cf23).cardinality)
            index94_ = default__.safeIndex(368, (d_729_v73_).length(0))
            rhs112_ = default__.safeDivisionInt((d_703_v60_) - (42), (d_703_v60_) * ((self).f14))
            rhs113_ = True
            rhs114_ = D1_DC4(d_732_v75_)
            rhs115_ = (0) - ((d_734_v77_).cf18)
            lhs76_ = d_729_v73_
            lhs77_ = default__.safeIndex(368, (d_729_v73_).length(0))
            lhs78_ = globalState
            lhs76_[lhs77_] = rhs112_
            lhs78_.f13 = rhs113_
            d_731_v74_ = rhs114_
            d_703_v60_ = rhs115_
        elif source11_.is_DC6:
            d_735___mcc_h4_ = source11_.cf10
            d_736_cf10_ = d_735___mcc_h4_
            d_737_v78_: _dafny.Array
            nw92_ = _dafny.Array(False, 5)
            d_737_v78_ = nw92_
            index95_ = default__.safeIndex(34, (d_737_v78_).length(0))
            (d_737_v78_)[index95_] = False
            d_738_v79_: D1
            d_738_v79_ = D1_DC5(d_634_v0_, p0, d_702_v59_)
            d_739_v80_: _dafny.Seq
            d_739_v80_ = _dafny.SeqWithoutIsStrInference([d_738_v79_])
            d_740_v81_: _dafny.MultiSet
            d_740_v81_ = _dafny.MultiSet([(d_738_v79_) in (d_739_v80_), d_634_v0_])
            d_741_v82_: _dafny.Array
            def lambda21_(d_742_p0_):
                def lambda22_(d_743_i9_):
                    return default__.safeModuloInt(d_743_i9_, len(d_742_p0_))

                return lambda22_

            init9_ = lambda21_(p0)
            nw93_ = _dafny.Array(None, 23)
            for i0_9_ in range(nw93_.length(0)):
                nw93_[i0_9_] = init9_(i0_9_)
            d_741_v82_ = nw93_
            index96_ = default__.safeIndex(217, (d_741_v82_).length(0))
            (d_741_v82_)[index96_] = (self).f14
            index97_ = default__.safeIndex(34, (d_737_v78_).length(0))
            index98_ = default__.safeIndex(217, (d_741_v82_).length(0))
            rhs116_ = d_634_v0_
            rhs117_ = (d_740_v81_) - (d_740_v81_)
            rhs118_ = (_dafny.SeqWithoutIsStrInference([(self).f14 for d_744_i10_ in range(default__.abs(661))])) <= (d_704_v61_)
            rhs119_ = (self).f14
            rhs120_ = not(d_634_v0_)
            lhs79_ = d_737_v78_
            lhs80_ = default__.safeIndex(34, (d_737_v78_).length(0))
            lhs81_ = globalState
            lhs82_ = d_741_v82_
            lhs83_ = default__.safeIndex(217, (d_741_v82_).length(0))
            lhs84_ = globalState
            lhs79_[lhs80_] = rhs116_
            d_740_v81_ = rhs117_
            lhs81_.f13 = rhs118_
            lhs82_[lhs83_] = rhs119_
            lhs84_.f13 = rhs120_
            (globalState).f5 = (d_737_v78_)[default__.safeIndex(34, (d_737_v78_).length(0))]
            d_745_v83_: _dafny.Seq
            d_745_v83_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "umgbbbrln"))
            d_745_v83_ = (p0) + (d_745_v83_)
            d_746_v84_: D6
            d_746_v84_ = D6_DC16((self).f14)
            (globalState).f5 = ((d_746_v84_).cf24) >= (d_703_v60_)
        elif True:
            d_747___mcc_h5_ = source11_.cf15
            d_748_cf15_ = d_747___mcc_h5_
            d_700_v57_ = (d_700_v57_ if not((d_703_v60_) >= (default__.fm11(d_634_v0_, globalState))) else d_700_v57_)
            (globalState).f13 = not(d_634_v0_)
            d_749_v85_: _dafny.Array
            nw94_ = _dafny.Array(_dafny.Map({}), 19)
            d_749_v85_ = nw94_
            d_750_v86_: _dafny.Array
            def lambda23_(d_751_i11_):
                return True

            init10_ = lambda23_
            nw95_ = _dafny.Array(None, 16)
            for i0_10_ in range(nw95_.length(0)):
                nw95_[i0_10_] = init10_(i0_10_)
            d_750_v86_ = nw95_
            index99_ = default__.safeIndex(451, (d_749_v85_).length(0))
            (d_749_v85_)[index99_] = _dafny.Map({d_750_v86_: d_634_v0_})
            d_752_v87_: _dafny.Array
            def lambda24_(d_753_v0_):
                def lambda25_(d_754_i12_):
                    return (_dafny.Set({d_753_v0_})).intersection(_dafny.Set({d_753_v0_}))

                return lambda25_

            init11_ = lambda24_(d_634_v0_)
            nw96_ = _dafny.Array(None, 12)
            for i0_11_ in range(nw96_.length(0)):
                nw96_[i0_11_] = init11_(i0_11_)
            d_752_v87_ = nw96_
            d_755_v88_: _dafny.Set
            d_755_v88_ = _dafny.Set({d_634_v0_, d_634_v0_, d_634_v0_, d_634_v0_, d_634_v0_})
            index100_ = default__.safeIndex(615, (d_752_v87_).length(0))
            (d_752_v87_)[index100_] = d_755_v88_
            d_756_v89_: _dafny.MultiSet
            d_756_v89_ = _dafny.MultiSet([d_634_v0_])
            d_757_v90_: _dafny.Map
            d_757_v90_ = _dafny.Map({d_750_v86_: (d_702_v59_)[default__.safeIndex((self).f14, len(d_702_v59_))]})
            d_758_v91_: _dafny.MultiSet
            d_758_v91_ = _dafny.MultiSet([d_703_v60_])
            d_759_v92_: _dafny.MultiSet
            d_759_v92_ = _dafny.MultiSet([(d_758_v91_).cardinality])
            index101_ = default__.safeIndex(451, (d_749_v85_).length(0))
            index102_ = default__.safeIndex(615, (d_752_v87_).length(0))
            rhs121_ = (((d_704_v61_) + (d_704_v61_)).set(default__.safeIndex(138, len((d_704_v61_) + (d_704_v61_))), (self).f14) if d_634_v0_ else d_704_v61_)
            rhs122_ = (d_757_v90_) | (d_757_v90_)
            rhs123_ = ((default__.fm33(d_634_v0_, d_634_v0_, (d_702_v59_)[default__.safeIndex(len(_dafny.Map({(self).f14: p0})), len(d_702_v59_))], globalState)) - (d_755_v88_)) | ((d_755_v88_) | (_dafny.Set({False})))
            rhs124_ = _dafny.MultiSet(d_702_v59_)
            rhs125_ = (d_750_v86_ if ((d_758_v91_).set(d_703_v60_, default__.abs(d_703_v60_))).issubset(d_759_v92_) else d_750_v86_)
            lhs85_ = d_749_v85_
            lhs86_ = default__.safeIndex(451, (d_749_v85_).length(0))
            lhs87_ = d_752_v87_
            lhs88_ = default__.safeIndex(615, (d_752_v87_).length(0))
            d_704_v61_ = rhs121_
            lhs85_[lhs86_] = rhs122_
            lhs87_[lhs88_] = rhs123_
            d_756_v89_ = rhs124_
            d_750_v86_ = rhs125_
            index103_ = default__.safeIndex(250, (d_750_v86_).length(0))
            (d_750_v86_)[index103_] = d_634_v0_
            index104_ = default__.safeIndex(250, (d_750_v86_).length(0))
            (d_750_v86_)[index104_] = d_634_v0_
        d_760_i13_: int
        d_760_i13_ = 0
        with _dafny.label("4"):
            while not(not (d_634_v0_) or (not (d_634_v0_) or (d_634_v0_))):
                with _dafny.c_label("4"):
                    if (d_760_i13_) >= (100):
                        raise _dafny.Break("4")
                    d_760_i13_ = (d_760_i13_) + (1)
                    d_761_v93_: _dafny.Seq
                    d_761_v93_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hglx"))
                    d_761_v93_ = (d_761_v93_).set(default__.safeIndex(((0) - ((self).f14)) + ((self).f14), len(d_761_v93_)), d_700_v57_)
                    d_703_v60_ = len(p0)
                    d_762_v94_: D1
                    d_762_v94_ = D1_DC5((self).fm8((self).f14, d_634_v0_, d_634_v0_, globalState), p0, d_702_v59_)
                    d_763_v95_: C0
                    nw97_ = C0()
                    nw97_.ctor__(d_709_v64_, (d_762_v94_).cf7)
                    d_763_v95_ = nw97_
                    d_702_v59_ = ((d_702_v59_) + (_dafny.SeqWithoutIsStrInference([d_634_v0_, d_634_v0_]))) + (d_702_v59_)
                    pass
            pass

    @property
    def f14(self):
        return self._f14
    @property
    def f15(self):
        return self._f15

class C5(T1, T0):
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    def ctor__(self):
        pass
        pass

    def fm0(self, p0, p1, globalState):
        def iife61_():
            coll23_ = _dafny.Set()
            compr_23_: int
            for compr_23_ in _dafny.IntegerRange(947, 745):
                d_764_v0_: int = compr_23_
                if ((947) <= (d_764_v0_)) and ((d_764_v0_) < (745)):
                    coll23_ = coll23_.union(_dafny.Set([default__.safeModuloInt(d_764_v0_, len(_dafny.SeqWithoutIsStrInference([False])))]))
            return _dafny.Set(coll23_)
        return ((len(iife61_()
        )) - (len(_dafny.Map({False: False})))) + (274)

    def fm1(self, p0, globalState):
        return _dafny.SeqWithoutIsStrInference([(_dafny.CodePoint('e') if False else _dafny.CodePoint('c')) for d_765_i0_ in range(default__.abs(714))])

    def fm4(self, p0, p1, p2, p3, globalState):
        def iife62_():
            def iife63_():
                coll25_ = _dafny.Map()
                compr_25_: int
                for compr_25_ in (_dafny.Set({(_dafny.MultiSet([636])).cardinality})).Elements:
                    d_767_v1_: int = compr_25_
                    if (d_767_v1_) in (_dafny.Set({(_dafny.MultiSet([636])).cardinality})):
                        coll25_[default__.safeDivisionInt(d_767_v1_, len(_dafny.SeqWithoutIsStrInference([True])))] = -483
                return _dafny.Map(coll25_)
            coll24_ = _dafny.Map()
            compr_24_: int
            for compr_24_ in _dafny.IntegerRange(-249, 694):
                d_766_v0_: int = compr_24_
                if ((-249) <= (d_766_v0_)) and ((d_766_v0_) < (694)):
                    coll24_[(d_766_v0_) * (len(iife63_()
                    ))] = 365
            return _dafny.Map(coll24_)
        return ((_dafny.SeqWithoutIsStrInference([len(_dafny.Map({825: len(_dafny.Set({False}))})), 764])) + (_dafny.SeqWithoutIsStrInference([-875, 578]))) + ((_dafny.SeqWithoutIsStrInference([856])) + (_dafny.SeqWithoutIsStrInference([-773, len(iife62_()
        )])))

    def fm5(self, p0, globalState):
        return 732

    def m0(self, p0, globalState):
        r0: bool = False
        r1: int = int(0)
        d_768_v0_: bool
        d_768_v0_ = True
        d_769_v1_: _dafny.Seq
        d_769_v1_ = _dafny.SeqWithoutIsStrInference([default__.fm6(globalState), d_768_v0_, d_768_v0_])
        d_770_v2_: int
        d_770_v2_ = -350
        d_771_v3_: _dafny.Seq
        d_771_v3_ = _dafny.SeqWithoutIsStrInference([d_769_v1_, (d_769_v1_).set(default__.safeIndex(d_770_v2_, len(d_769_v1_)), d_768_v0_)])
        d_772_v4_: _dafny.Seq
        d_772_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dnhvpnw"))
        d_773_v5_: _dafny.Seq
        d_773_v5_ = _dafny.SeqWithoutIsStrInference([d_772_v4_, d_772_v4_])
        d_774_i0_: int
        d_774_i0_ = 0
        with _dafny.label("5"):
            while (d_769_v1_) == (((d_771_v3_)[default__.safeIndex(len(d_773_v5_), len(d_771_v3_))]) + (_dafny.SeqWithoutIsStrInference([d_768_v0_]))):
                with _dafny.c_label("5"):
                    if (d_774_i0_) >= (100):
                        raise _dafny.Break("5")
                    d_774_i0_ = (d_774_i0_) + (1)
                    d_775_v6_: C2
                    nw98_ = C2()
                    nw98_.ctor__()
                    d_775_v6_ = nw98_
                    d_770_v2_ = default__.safeDivisionInt((0) - (d_770_v2_), d_770_v2_)
                    if d_768_v0_:
                        d_770_v2_ = d_770_v2_
                        r1 = d_770_v2_
                        d_776_v8_: _dafny.MultiSet
                        d_776_v8_ = _dafny.MultiSet([d_770_v2_])
                        d_777_v9_: _dafny.Map
                        d_777_v9_ = _dafny.Map({d_770_v2_: d_770_v2_})
                        d_778_v10_: _dafny.Seq
                        def iife64_():
                            coll26_ = _dafny.Map()
                            compr_26_: int
                            for compr_26_ in (_dafny.MultiSet([-454, d_770_v2_, 576, d_770_v2_, d_770_v2_])).Elements:
                                d_779_v7_: int = compr_26_
                                if (d_779_v7_) in (_dafny.MultiSet([-454, d_770_v2_, 576, d_770_v2_, d_770_v2_])):
                                    coll26_[default__.safeDivisionInt(d_779_v7_, d_770_v2_)] = d_770_v2_
                            return _dafny.Map(coll26_)
                        def iife65_():
                            coll27_ = _dafny.Map()
                            compr_27_: int
                            for compr_27_ in (_dafny.MultiSet([-454, d_770_v2_, 576, d_770_v2_, d_770_v2_])).Elements:
                                d_781_v7_: int = compr_27_
                                if (d_781_v7_) in (_dafny.MultiSet([-454, d_770_v2_, 576, d_770_v2_, d_770_v2_])):
                                    coll27_[default__.safeDivisionInt(d_781_v7_, d_770_v2_)] = d_770_v2_
                            return _dafny.Map(coll27_)
                        d_778_v10_ = _dafny.SeqWithoutIsStrInference([d_770_v2_, (d_770_v2_) - (len((_dafny.SeqWithoutIsStrInference([iife64_()
 for d_780_i1_ in range(default__.abs(866))])).set(default__.safeIndex((0) - ((d_776_v8_).cardinality), len(_dafny.SeqWithoutIsStrInference([iife65_()
 for d_782_i1_ in range(default__.abs(866))]))), d_777_v9_))), len(d_772_v4_)])
                        rhs126_ = d_768_v0_
                        rhs127_ = (d_770_v2_) * (d_770_v2_)
                        rhs128_ = d_778_v10_
                        r0 = rhs126_
                        d_770_v2_ = rhs127_
                        d_778_v10_ = rhs128_
                        d_783_v11_: str
                        d_783_v11_ = _dafny.CodePoint('b')
                        d_783_v11_ = d_783_v11_
                        d_784_v12_: _dafny.Array
                        def lambda26_(d_785_v2_):
                            def lambda27_(d_786_i2_):
                                return default__.safeModuloInt(d_786_i2_, d_785_v2_)

                            return lambda27_

                        init12_ = lambda26_(d_770_v2_)
                        nw99_ = _dafny.Array(None, 20)
                        for i0_12_ in range(nw99_.length(0)):
                            nw99_[i0_12_] = init12_(i0_12_)
                        d_784_v12_ = nw99_
                        d_787_v13_: _dafny.Map
                        d_787_v13_ = _dafny.Map({d_784_v12_: _dafny.SeqWithoutIsStrInference([p0 for d_788_i3_ in range(default__.abs(805))])})
                        d_789_v14_: D0
                        d_789_v14_ = D0_DC1(False, d_773_v5_, ((d_787_v13_)[d_784_v12_] if (d_784_v12_) in (d_787_v13_) else d_772_v4_))
                        d_790_v15_: _dafny.MultiSet
                        d_790_v15_ = _dafny.MultiSet([d_789_v14_, d_789_v14_])
                        d_791_v16_: D2
                        d_791_v16_ = D2_DC6(d_790_v15_)
                        d_792_v17_: C0
                        nw100_ = C0()
                        nw100_.ctor__((d_791_v16_).cf10, d_768_v0_)
                        d_792_v17_ = nw100_
                    elif True:
                        d_768_v0_ = (d_768_v0_) or (True)
                        (globalState).f13 = d_768_v0_
                        (globalState).f13 = d_768_v0_
                        d_769_v1_ = d_769_v1_
                        (globalState).f9 = not(d_768_v0_)
                    d_793_v18_: _dafny.Map
                    d_793_v18_ = _dafny.Map({-708: d_768_v0_})
                    d_794_v20_: _dafny.Set
                    d_794_v20_ = _dafny.Set({d_768_v0_})
                    d_795_v21_: _dafny.Map
                    d_795_v21_ = _dafny.Map({len(d_794_v20_): d_770_v2_})
                    d_796_v22_: D7
                    def iife66_():
                        coll28_ = _dafny.Map()
                        compr_28_: int
                        for compr_28_ in (d_795_v21_).keys.Elements:
                            d_797_v19_: int = compr_28_
                            if (d_797_v19_) in (d_795_v21_):
                                coll28_[(d_797_v19_) * (523)] = True
                        return _dafny.Map(coll28_)
                    d_796_v22_ = D7_DC18(iife66_()
)
                    d_793_v18_ = (d_796_v22_).cf26
                    pass
            pass
        if d_768_v0_:
            d_798_v23_: _dafny.Array
            def lambda28_(d_799_p0_):
                def lambda29_(d_800_i4_):
                    return d_799_p0_

                return lambda29_

            init13_ = lambda28_(p0)
            nw101_ = _dafny.Array(None, 12)
            for i0_13_ in range(nw101_.length(0)):
                nw101_[i0_13_] = init13_(i0_13_)
            d_798_v23_ = nw101_
            index105_ = default__.safeIndex(92, (d_798_v23_).length(0))
            (d_798_v23_)[index105_] = p0
            index106_ = default__.safeIndex(92, (d_798_v23_).length(0))
            (d_798_v23_)[index106_] = p0
            d_801_v24_: D0
            d_801_v24_ = D0_DC3(D0_DC1(d_768_v0_, d_773_v5_, d_772_v4_))
            d_802_v25_: _dafny.Map
            d_802_v25_ = _dafny.Map({d_768_v0_: d_801_v24_})
            d_803_v26_: D2
            d_803_v26_ = D2_DC7(d_768_v0_, d_802_v25_, d_770_v2_, d_768_v0_)
            d_804_v27_: D0
            d_804_v27_ = D0_DC1(True, d_773_v5_, d_772_v4_)
            d_770_v2_ = ((d_770_v2_) - (default__.fm22(d_803_v26_, D0_DC3(d_804_v27_), d_768_v0_, globalState))) + (default__.fm22(D2_DC7(d_768_v0_, d_802_v25_, 168, d_768_v0_), d_801_v24_, d_768_v0_, globalState))
            d_805_v28_: _dafny.Array
            nw102_ = _dafny.Array(None, 27)
            nw102_[int(0)] = d_768_v0_
            nw102_[int(1)] = d_768_v0_
            nw102_[int(2)] = d_768_v0_
            nw102_[int(3)] = d_768_v0_
            nw102_[int(4)] = not(d_768_v0_)
            nw102_[int(5)] = not(d_768_v0_)
            nw102_[int(6)] = d_768_v0_
            nw102_[int(7)] = d_768_v0_
            nw102_[int(8)] = d_768_v0_
            nw102_[int(9)] = d_768_v0_
            nw102_[int(10)] = d_768_v0_
            nw102_[int(11)] = d_768_v0_
            nw102_[int(12)] = d_768_v0_
            nw102_[int(13)] = not(d_768_v0_)
            nw102_[int(14)] = not(d_768_v0_)
            nw102_[int(15)] = d_768_v0_
            nw102_[int(16)] = d_768_v0_
            nw102_[int(17)] = d_768_v0_
            nw102_[int(18)] = d_768_v0_
            nw102_[int(19)] = d_768_v0_
            nw102_[int(20)] = d_768_v0_
            nw102_[int(21)] = False
            nw102_[int(22)] = True
            nw102_[int(23)] = d_768_v0_
            nw102_[int(24)] = d_768_v0_
            nw102_[int(25)] = d_768_v0_
            nw102_[int(26)] = d_768_v0_
            d_805_v28_ = nw102_
            d_806_v29_: _dafny.Array
            nw103_ = _dafny.Array(None, 16)
            nw103_[int(0)] = d_805_v28_
            nw103_[int(1)] = d_805_v28_
            nw103_[int(2)] = d_805_v28_
            nw103_[int(3)] = d_805_v28_
            nw103_[int(4)] = d_805_v28_
            nw103_[int(5)] = d_805_v28_
            nw103_[int(6)] = d_805_v28_
            nw103_[int(7)] = d_805_v28_
            nw103_[int(8)] = d_805_v28_
            nw103_[int(9)] = d_805_v28_
            nw103_[int(10)] = d_805_v28_
            nw103_[int(11)] = d_805_v28_
            nw103_[int(12)] = d_805_v28_
            nw103_[int(13)] = d_805_v28_
            nw103_[int(14)] = d_805_v28_
            nw103_[int(15)] = d_805_v28_
            d_806_v29_ = nw103_
            d_806_v29_ = d_806_v29_
            d_807_v30_: T1
            nw104_ = C3()
            nw104_.ctor__(d_768_v0_, d_770_v2_)
            d_807_v30_ = nw104_
            d_808_v31_: _dafny.Set
            d_808_v31_ = _dafny.Set({d_768_v0_})
            d_809_v32_: _dafny.MultiSet
            d_809_v32_ = _dafny.MultiSet([d_770_v2_])
            (self).m3((d_770_v2_) + (d_770_v2_), d_807_v30_, len((d_808_v31_) | (_dafny.Set({d_768_v0_, d_768_v0_, d_768_v0_}))), (d_809_v32_) | (d_809_v32_), globalState)
            d_810_v33_: _dafny.Set
            d_810_v33_ = d_808_v31_
            d_811_v34_: _dafny.Set
            d_811_v34_ = _dafny.Set({d_810_v33_})
            d_812_v35_: _dafny.Map
            d_812_v35_ = _dafny.Map({d_811_v34_: d_807_v30_})
            (globalState).f3 = not((d_812_v35_) != (d_812_v35_))
        elif True:
            d_813_v36_: _dafny.MultiSet
            d_813_v36_ = _dafny.MultiSet([d_768_v0_])
            d_814_v37_: _dafny.Seq
            d_814_v37_ = _dafny.SeqWithoutIsStrInference([((d_813_v36_)[d_768_v0_] if (d_768_v0_) in (d_813_v36_) else d_770_v2_), d_770_v2_, d_770_v2_, d_770_v2_])
            d_770_v2_ = default__.safeModuloInt(d_770_v2_, (((d_813_v36_)[d_768_v0_] if (d_768_v0_) in (d_813_v36_) else len(d_814_v37_))) + (default__.fm11(default__.fm6(globalState), globalState)))
            r0 = d_768_v0_
            d_815_v38_: str
            d_815_v38_ = _dafny.CodePoint('e')
            d_816_v39_: _dafny.Array
            def lambda30_(d_817_v2_):
                def lambda31_(d_818_i5_):
                    return (d_818_i5_) - (d_817_v2_)

                return lambda31_

            init14_ = lambda30_(d_770_v2_)
            nw105_ = _dafny.Array(None, 26)
            for i0_14_ in range(nw105_.length(0)):
                nw105_[i0_14_] = init14_(i0_14_)
            d_816_v39_ = nw105_
            index107_ = default__.safeIndex(448, (d_816_v39_).length(0))
            (d_816_v39_)[index107_] = len((d_772_v4_) + (d_772_v4_))
            index108_ = default__.safeIndex(837, (d_816_v39_).length(0))
            (d_816_v39_)[index108_] = (d_814_v37_)[default__.safeIndex(d_770_v2_, len(d_814_v37_))]
            d_819_v40_: D0
            d_819_v40_ = D0_DC2(d_768_v0_)
            d_820_v41_: _dafny.Seq
            d_820_v41_ = _dafny.SeqWithoutIsStrInference([d_819_v40_])
            d_821_v42_: _dafny.Map
            d_821_v42_ = _dafny.Map({d_770_v2_: not(False)})
            index109_ = default__.safeIndex(448, (d_816_v39_).length(0))
            index110_ = default__.safeIndex(837, (d_816_v39_).length(0))
            rhs129_ = d_815_v38_
            rhs130_ = len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ebekj"))) + ((d_773_v5_)[default__.safeIndex(d_770_v2_, len(d_773_v5_))]))
            rhs131_ = len(d_769_v1_)
            rhs132_ = (-52) - (len(d_821_v42_))
            rhs133_ = d_820_v41_
            lhs89_ = d_816_v39_
            lhs90_ = default__.safeIndex(448, (d_816_v39_).length(0))
            lhs91_ = d_816_v39_
            lhs92_ = default__.safeIndex(837, (d_816_v39_).length(0))
            d_815_v38_ = rhs129_
            lhs89_[lhs90_] = rhs130_
            lhs91_[lhs92_] = rhs131_
            r1 = rhs132_
            d_820_v41_ = rhs133_
            r1 = d_770_v2_
            d_822_v43_: C2
            nw106_ = C2()
            nw106_.ctor__()
            d_822_v43_ = nw106_
        if (d_768_v0_) and (d_768_v0_):
            d_823_v44_: _dafny.Array
            nw107_ = _dafny.Array(None, 19)
            nw107_[int(0)] = p0
            nw107_[int(1)] = p0
            nw107_[int(2)] = p0
            nw107_[int(3)] = p0
            nw107_[int(4)] = _dafny.CodePoint('c')
            nw107_[int(5)] = p0
            nw107_[int(6)] = _dafny.CodePoint('b')
            nw107_[int(7)] = p0
            nw107_[int(8)] = p0
            nw107_[int(9)] = p0
            nw107_[int(10)] = p0
            nw107_[int(11)] = p0
            nw107_[int(12)] = _dafny.CodePoint('e')
            nw107_[int(13)] = p0
            nw107_[int(14)] = p0
            nw107_[int(15)] = p0
            nw107_[int(16)] = _dafny.CodePoint('q')
            nw107_[int(17)] = p0
            nw107_[int(18)] = p0
            d_823_v44_ = nw107_
            index111_ = default__.safeIndex(138, (d_823_v44_).length(0))
            (d_823_v44_)[index111_] = p0
            index112_ = default__.safeIndex(138, (d_823_v44_).length(0))
            (d_823_v44_)[index112_] = p0
            d_824_v45_: _dafny.Array
            def lambda32_(d_825_v2_):
                def lambda33_(d_826_i6_):
                    return default__.safeDivisionInt(d_826_i6_, d_825_v2_)

                return lambda33_

            init15_ = lambda32_(d_770_v2_)
            nw108_ = _dafny.Array(None, 28)
            for i0_15_ in range(nw108_.length(0)):
                nw108_[i0_15_] = init15_(i0_15_)
            d_824_v45_ = nw108_
            d_827_v46_: _dafny.Set
            d_827_v46_ = _dafny.Set({d_768_v0_, default__.fm6(globalState), d_768_v0_, not(d_768_v0_), d_768_v0_})
            index113_ = default__.safeIndex(312, (d_824_v45_).length(0))
            (d_824_v45_)[index113_] = len(d_827_v46_)
            d_828_v47_: _dafny.Set
            d_828_v47_ = _dafny.Set({d_770_v2_})
            d_829_v48_: _dafny.Map
            d_829_v48_ = _dafny.Map({d_770_v2_: len(d_828_v47_)})
            index114_ = default__.safeIndex(312, (d_824_v45_).length(0))
            (d_824_v45_)[index114_] = (len(d_829_v48_)) * (d_770_v2_)
            d_830_v49_: _dafny.Array
            nw109_ = _dafny.Array(_dafny.Set({}), 19)
            d_830_v49_ = nw109_
            d_831_v50_: _dafny.Set
            d_831_v50_ = _dafny.Set({d_824_v45_})
            index115_ = default__.safeIndex(637, (d_830_v49_).length(0))
            (d_830_v49_)[index115_] = d_831_v50_
            index116_ = default__.safeIndex(637, (d_830_v49_).length(0))
            (d_830_v49_)[index116_] = d_831_v50_
            d_770_v2_ = ((d_829_v48_)[(0) - ((d_824_v45_)[default__.safeIndex(312, (d_824_v45_).length(0))])] if ((0) - ((d_824_v45_)[default__.safeIndex(312, (d_824_v45_).length(0))])) in (d_829_v48_) else (d_770_v2_) - (d_770_v2_))
            r1 = (d_824_v45_)[default__.safeIndex(312, (d_824_v45_).length(0))]
        elif True:
            (globalState).f9 = (d_768_v0_ if d_768_v0_ else d_768_v0_)
            d_832_v51_: _dafny.Set
            d_832_v51_ = _dafny.Set({not(False), d_768_v0_, d_768_v0_})
            d_833_v52_: _dafny.Map
            d_833_v52_ = _dafny.Map({len(d_832_v51_): d_770_v2_})
            d_833_v52_ = (d_833_v52_).set((151) * (147), (0) - (default__.safeDivisionInt(d_770_v2_, d_770_v2_)))
            d_834_v53_: _dafny.Seq
            d_834_v53_ = _dafny.SeqWithoutIsStrInference([len(d_772_v4_)])
            d_835_v54_: _dafny.MultiSet
            d_835_v54_ = _dafny.MultiSet([d_834_v53_, _dafny.SeqWithoutIsStrInference([d_770_v2_, d_770_v2_, d_770_v2_, d_770_v2_])])
            d_836_v55_: D1
            d_836_v55_ = D1_DC5(d_768_v0_, d_772_v4_, d_769_v1_)
            d_837_v56_: _dafny.Array
            nw110_ = _dafny.Array(None, 24)
            nw110_[int(0)] = d_772_v4_
            nw110_[int(1)] = d_772_v4_
            nw110_[int(2)] = _dafny.SeqWithoutIsStrInference([p0 for d_838_i7_ in range(default__.abs(-339))])
            nw110_[int(3)] = d_772_v4_
            nw110_[int(4)] = _dafny.SeqWithoutIsStrInference([p0 for d_839_i8_ in range(default__.abs(932))])
            nw110_[int(5)] = _dafny.SeqWithoutIsStrInference([p0 for d_840_i9_ in range(default__.abs(-895))])
            nw110_[int(6)] = d_772_v4_
            nw110_[int(7)] = d_772_v4_
            nw110_[int(8)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rl"))
            nw110_[int(9)] = d_772_v4_
            nw110_[int(10)] = _dafny.SeqWithoutIsStrInference([p0 for d_841_i10_ in range(default__.abs(-2))])
            nw110_[int(11)] = d_772_v4_
            nw110_[int(12)] = default__.fm20(len(d_834_v53_), d_768_v0_, d_768_v0_, d_770_v2_, globalState)
            nw110_[int(13)] = d_772_v4_
            nw110_[int(14)] = d_772_v4_
            nw110_[int(15)] = d_772_v4_
            nw110_[int(16)] = d_772_v4_
            nw110_[int(17)] = d_772_v4_
            nw110_[int(18)] = d_772_v4_
            nw110_[int(19)] = d_772_v4_
            nw110_[int(20)] = d_772_v4_
            nw110_[int(21)] = d_772_v4_
            nw110_[int(22)] = d_772_v4_
            nw110_[int(23)] = d_772_v4_
            d_837_v56_ = nw110_
            d_842_v57_: D7
            d_842_v57_ = D7_DC19(d_836_v55_, d_770_v2_, d_768_v0_, d_837_v56_)
            d_843_v58_: D0
            d_843_v58_ = D0_DC2(False)
            d_844_v59_: _dafny.Seq
            d_844_v59_ = _dafny.SeqWithoutIsStrInference([d_834_v53_, d_834_v53_, ((d_834_v53_).set(default__.safeIndex(d_770_v2_, len(d_834_v53_)), 927)).set(default__.safeIndex((d_842_v57_).cf28, len((d_834_v53_).set(default__.safeIndex(d_770_v2_, len(d_834_v53_)), 927))), (self).fm5(d_843_v58_, globalState)), _dafny.SeqWithoutIsStrInference([d_770_v2_]), d_834_v53_])
            if ((d_835_v54_).intersection(default__.fm34(globalState))).isdisjoint(_dafny.MultiSet(d_844_v59_)):
                d_845_v60_: _dafny.Map
                d_845_v60_ = _dafny.Map({d_768_v0_: d_770_v2_})
                d_845_v60_ = (d_845_v60_).set(True, d_770_v2_)
                d_846_v61_: D0
                d_846_v61_ = D0_DC1(d_768_v0_, d_773_v5_, d_772_v4_)
                d_847_v62_: _dafny.MultiSet
                d_847_v62_ = _dafny.MultiSet([d_846_v61_])
                d_848_v63_: C0
                nw111_ = C0()
                nw111_.ctor__(d_847_v62_, d_768_v0_)
                d_848_v63_ = nw111_
                d_849_v64_: _dafny.Array
                nw112_ = _dafny.Array(False, 27)
                d_849_v64_ = nw112_
                d_850_v65_: bool
                d_851_v66_: int
                d_852_v67_: int
                out11_: bool
                out12_: int
                out13_: int
                out11_, out12_, out13_ = (self).m4(d_834_v53_, d_849_v64_, globalState)
                d_850_v65_ = out11_
                d_851_v66_ = out12_
                d_852_v67_ = out13_
                index117_ = default__.safeIndex(365, (d_849_v64_).length(0))
                (d_849_v64_)[index117_] = (d_768_v0_ if d_850_v65_ else d_768_v0_)
                index118_ = default__.safeIndex(365, (d_849_v64_).length(0))
                (d_849_v64_)[index118_] = d_850_v65_
                d_853_v68_: _dafny.Array
                nw113_ = _dafny.Array(_dafny.Array(None, 0), 21)
                d_853_v68_ = nw113_
                index119_ = default__.safeIndex(773, (d_853_v68_).length(0))
                (d_853_v68_)[index119_] = d_849_v64_
                index120_ = default__.safeIndex(773, (d_853_v68_).length(0))
                (d_853_v68_)[index120_] = d_849_v64_
            elif True:
                d_854_v69_: _dafny.MultiSet
                d_854_v69_ = _dafny.MultiSet([d_770_v2_])
                (globalState).f5 = not((_dafny.MultiSet(d_834_v53_)).issubset(d_854_v69_))
                (globalState).f13 = ((0) - (d_770_v2_)) <= ((0) - (-799))
                d_855_v70_: _dafny.Array
                nw114_ = _dafny.Array(None, 12)
                nw114_[int(0)] = (len((d_772_v4_).set(default__.safeIndex(d_770_v2_, len(d_772_v4_)), p0)) if d_768_v0_ else d_770_v2_)
                nw114_[int(1)] = len((d_769_v1_) + (_dafny.SeqWithoutIsStrInference([d_768_v0_])))
                nw114_[int(2)] = (d_770_v2_) + (d_770_v2_)
                nw114_[int(3)] = d_770_v2_
                nw114_[int(4)] = d_770_v2_
                nw114_[int(5)] = (0) - (d_770_v2_)
                nw114_[int(6)] = (0) - (d_770_v2_)
                nw114_[int(7)] = default__.safeModuloInt(d_770_v2_, d_770_v2_)
                nw114_[int(8)] = d_770_v2_
                nw114_[int(9)] = d_770_v2_
                nw114_[int(10)] = (d_770_v2_) + (d_770_v2_)
                nw114_[int(11)] = len(d_772_v4_)
                d_855_v70_ = nw114_
                index121_ = default__.safeIndex(459, (d_855_v70_).length(0))
                (d_855_v70_)[index121_] = (d_770_v2_) + (d_770_v2_)
                d_856_v71_: _dafny.Array
                nw115_ = _dafny.Array(False, 11)
                d_856_v71_ = nw115_
                index122_ = default__.safeIndex(83, (d_856_v71_).length(0))
                (d_856_v71_)[index122_] = not(not(not(d_768_v0_)))
                index123_ = default__.safeIndex(459, (d_855_v70_).length(0))
                index124_ = default__.safeIndex(83, (d_856_v71_).length(0))
                rhs134_ = d_770_v2_
                rhs135_ = d_768_v0_
                rhs136_ = d_768_v0_
                lhs93_ = d_855_v70_
                lhs94_ = default__.safeIndex(459, (d_855_v70_).length(0))
                lhs95_ = d_856_v71_
                lhs96_ = default__.safeIndex(83, (d_856_v71_).length(0))
                lhs97_ = globalState
                lhs93_[lhs94_] = rhs134_
                lhs95_[lhs96_] = rhs135_
                lhs97_.f9 = rhs136_
                d_857_v72_: _dafny.Array
                def lambda34_(d_858_v53_):
                    def lambda35_(d_859_i11_):
                        return d_858_v53_

                    return lambda35_

                init16_ = lambda34_(d_834_v53_)
                nw116_ = _dafny.Array(None, 11)
                for i0_16_ in range(nw116_.length(0)):
                    nw116_[i0_16_] = init16_(i0_16_)
                d_857_v72_ = nw116_
                index125_ = default__.safeIndex(705, (d_857_v72_).length(0))
                (d_857_v72_)[index125_] = (d_834_v53_) + (d_834_v53_)
                d_860_v73_: str
                d_860_v73_ = _dafny.CodePoint('o')
                index126_ = default__.safeIndex(705, (d_857_v72_).length(0))
                rhs137_ = ((d_834_v53_) + ((d_834_v53_ if True else _dafny.SeqWithoutIsStrInference([len(_dafny.Map({(d_856_v71_)[default__.safeIndex(83, (d_856_v71_).length(0))]: d_770_v2_})) for d_861_i12_ in range(default__.abs(735))])))).set(default__.safeIndex((0) - (d_770_v2_), len((d_834_v53_) + ((d_834_v53_ if True else _dafny.SeqWithoutIsStrInference([len(_dafny.Map({(d_856_v71_)[default__.safeIndex(83, (d_856_v71_).length(0))]: d_770_v2_})) for d_862_i12_ in range(default__.abs(735))]))))), (d_855_v70_)[default__.safeIndex(459, (d_855_v70_).length(0))])
                rhs138_ = (d_855_v70_)[default__.safeIndex(459, (d_855_v70_).length(0))]
                rhs139_ = (((d_772_v4_).set(default__.safeIndex(d_770_v2_, len(d_772_v4_)), p0)) + (d_772_v4_)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wje")))
                rhs140_ = (d_772_v4_)[default__.safeIndex(len(d_772_v4_), len(d_772_v4_))]
                lhs98_ = d_857_v72_
                lhs99_ = default__.safeIndex(705, (d_857_v72_).length(0))
                lhs98_[lhs99_] = rhs137_
                d_770_v2_ = rhs138_
                d_772_v4_ = rhs139_
                d_860_v73_ = rhs140_
                d_860_v73_ = p0
            d_863_v74_: D0
            d_863_v74_ = D0_DC1(d_768_v0_, d_773_v5_, d_772_v4_)
            d_864_v75_: _dafny.MultiSet
            d_864_v75_ = _dafny.MultiSet([d_863_v74_, d_863_v74_, d_863_v74_, d_863_v74_, d_863_v74_])
            d_865_v76_: C0
            nw117_ = C0()
            nw117_.ctor__(d_864_v75_, not(d_768_v0_))
            d_865_v76_ = nw117_
            d_866_v77_: str
            d_866_v77_ = _dafny.CodePoint('s')
            d_866_v77_ = d_866_v77_
        d_867_v78_: _dafny.Set
        d_867_v78_ = _dafny.Set({d_768_v0_, d_768_v0_, d_768_v0_})
        d_867_v78_ = (d_867_v78_).intersection(d_867_v78_)
        d_868_v79_: D0
        d_868_v79_ = D0_DC2(d_768_v0_)
        d_869_v80_: C4
        nw118_ = C4()
        nw118_.ctor__(d_770_v2_, (_dafny.Map({_dafny.CodePoint('r'): (self).fm5(d_868_v79_, globalState)})) | (_dafny.Map({p0: d_770_v2_})))
        d_869_v80_ = nw118_
        d_870_i13_: int
        d_870_i13_ = 0
        with _dafny.label("6"):
            while d_768_v0_:
                with _dafny.c_label("6"):
                    if (d_870_i13_) >= (100):
                        raise _dafny.Break("6")
                    d_870_i13_ = (d_870_i13_) + (1)
                    d_871_v81_: _dafny.MultiSet
                    d_871_v81_ = _dafny.MultiSet([(d_869_v80_).f14, ((d_869_v80_).f14) * (len(d_769_v1_))])
                    d_871_v81_ = d_871_v81_
                    d_770_v2_ = (0) - (default__.fm11(d_768_v0_, globalState))
                    d_770_v2_ = (d_869_v80_).f14
                    (globalState).f5 = d_768_v0_
                    pass
            pass
        r0 = (d_768_v0_) or (False)
        r1 = (0) - ((d_770_v2_) * (len(((d_869_v80_).f15) | ((d_869_v80_).f15))))
        return r0, r1

    def m3(self, p0, p1, p2, p3, globalState):
        d_872_v0_: _dafny.Array
        nw119_ = _dafny.Array(_dafny.CodePoint('D'), 6)
        d_872_v0_ = nw119_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_872_v0_).length(0)):
            d_873_i0_: int = guard_loop_1_
            if (True) and (((0) <= (d_873_i0_)) and ((d_873_i0_) < ((d_872_v0_).length(0)))):
                (d_872_v0_)[(d_873_i0_)] = (_dafny.CodePoint('y') if True else _dafny.CodePoint('a'))
        d_874_v1_: _dafny.Array
        nw120_ = _dafny.Array(False, 28)
        d_874_v1_ = nw120_
        d_875_v2_: bool
        d_875_v2_ = True
        index127_ = default__.safeIndex(89, (d_874_v1_).length(0))
        (d_874_v1_)[index127_] = not(not(d_875_v2_))
        d_876_v3_: int
        d_876_v3_ = -355
        d_877_v4_: _dafny.Seq
        d_877_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hlo"))
        d_878_v5_: _dafny.Map
        d_878_v5_ = _dafny.Map({504: True})
        d_879_v6_: D7
        d_879_v6_ = D7_DC18(d_878_v5_)
        d_880_v7_: D7
        d_880_v7_ = D7_DC21(d_879_v6_)
        d_881_v8_: _dafny.Map
        d_881_v8_ = _dafny.Map({d_876_v3_: d_880_v7_})
        d_882_v9_: _dafny.Set
        d_882_v9_ = _dafny.Set({len(_dafny.Map({d_875_v2_: d_881_v8_}))})
        d_883_v10_: _dafny.Set
        d_883_v10_ = _dafny.Set({d_875_v2_})
        index128_ = default__.safeIndex(89, (d_874_v1_).length(0))
        rhs141_ = not((len(_dafny.SeqWithoutIsStrInference([p0 for d_884_i1_ in range(default__.abs(7))]))) not in (d_882_v9_))
        rhs142_ = ((0) - ((0) - (p0))) * (len((d_883_v10_ if d_875_v2_ else d_883_v10_)))
        rhs143_ = d_875_v2_
        rhs144_ = d_877_v4_
        lhs100_ = d_874_v1_
        lhs101_ = default__.safeIndex(89, (d_874_v1_).length(0))
        lhs100_[lhs101_] = rhs141_
        d_876_v3_ = rhs142_
        d_875_v2_ = rhs143_
        d_877_v4_ = rhs144_
        d_876_v3_ = ((p3)[p2] if (p2) in (p3) else default__.safeDivisionInt(p0, d_876_v3_))
        (globalState).f3 = d_875_v2_
        (globalState).f9 = not ((d_874_v1_)[default__.safeIndex(89, (d_874_v1_).length(0))]) or (d_875_v2_)
        d_885_v11_: D0
        d_885_v11_ = D0_DC2(d_875_v2_)
        d_886_v12_: D0
        d_886_v12_ = D0_DC3(d_885_v11_)
        source12_ = d_886_v12_
        if source12_.is_DC1:
            d_887___mcc_h0_ = source12_.cf1
            d_888___mcc_h1_ = source12_.cf2
            d_889___mcc_h2_ = source12_.cf3
            d_890_cf3_ = d_889___mcc_h2_
            d_891_cf2_ = d_888___mcc_h1_
            d_892_cf1_ = d_887___mcc_h0_
            d_876_v3_ = default__.safeDivisionInt(d_876_v3_, (0) - (p2))
            d_876_v3_ = ((0) - (p0)) * ((p2) * (p0))
            d_893_v13_: D4
            d_893_v13_ = D4_DC11()
            source13_ = d_893_v13_
            if source13_.is_DC11:
                d_894_v14_: _dafny.Array
                nw121_ = _dafny.Array(int(0), 11)
                d_894_v14_ = nw121_
                index129_ = default__.safeIndex(381, (d_894_v14_).length(0))
                (d_894_v14_)[index129_] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lklrsh")))
                index130_ = default__.safeIndex(381, (d_894_v14_).length(0))
                (d_894_v14_)[index130_] = p0
                d_895_v15_: str
                d_895_v15_ = _dafny.CodePoint('n')
                d_896_v16_: _dafny.MultiSet
                d_896_v16_ = _dafny.MultiSet([d_895_v15_])
                d_876_v3_ = (d_896_v16_).cardinality
                d_897_v17_: _dafny.Map
                d_897_v17_ = _dafny.Map({d_892_cf1_: d_893_v13_})
                d_898_v18_: C3
                nw122_ = C3()
                nw122_.ctor__((d_874_v1_)[default__.safeIndex(89, (d_874_v1_).length(0))], (p2 if (d_874_v1_)[default__.safeIndex(89, (d_874_v1_).length(0))] else len(d_897_v17_)))
                d_898_v18_ = nw122_
                d_882_v9_ = default__.fm10((d_874_v1_)[default__.safeIndex(89, (d_874_v1_).length(0))], p2, d_892_cf1_, (p2) >= (p2), globalState)
            elif source13_.is_DC12:
                d_899___mcc_h6_ = source13_.cf18
                d_900_cf18_ = d_899___mcc_h6_
                d_882_v9_ = d_882_v9_
                d_901_v19_: _dafny.Array
                nw123_ = _dafny.Array(int(0), 13)
                d_901_v19_ = nw123_
                index131_ = default__.safeIndex(206, (d_901_v19_).length(0))
                (d_901_v19_)[index131_] = (d_900_cf18_) - (-548)
                index132_ = default__.safeIndex(206, (d_901_v19_).length(0))
                (d_901_v19_)[index132_] = (0) - (p0)
                (globalState).f13 = (d_874_v1_)[default__.safeIndex(89, (d_874_v1_).length(0))]
                d_902_v20_: _dafny.Seq
                d_902_v20_ = _dafny.SeqWithoutIsStrInference([d_892_cf1_])
                d_903_v21_: _dafny.MultiSet
                d_903_v21_ = _dafny.MultiSet([d_872_v0_, d_872_v0_, d_872_v0_])
                d_904_v22_: _dafny.Map
                d_904_v22_ = _dafny.Map({d_902_v20_: (d_903_v21_) | (d_903_v21_)})
                d_904_v22_ = (d_904_v22_).set(d_902_v20_, (d_903_v21_ if (d_874_v1_)[default__.safeIndex(89, (d_874_v1_).length(0))] else d_903_v21_))
            elif True:
                d_905___mcc_h7_ = source13_.cf17
                d_906_cf17_ = d_905___mcc_h7_
                d_907_v23_: str
                d_907_v23_ = _dafny.CodePoint('e')
                d_908_v25_: _dafny.Seq
                d_908_v25_ = _dafny.SeqWithoutIsStrInference([d_906_cf17_])
                def iife67_():
                    coll29_ = _dafny.Map()
                    compr_29_: _dafny.Set
                    for compr_29_ in (_dafny.MultiSet(d_908_v25_)).Elements:
                        d_909_v24_: _dafny.Set = compr_29_
                        if (d_909_v24_) in (_dafny.MultiSet(d_908_v25_)):
                            coll29_[d_909_v24_] = d_876_v3_
                    return _dafny.Map(coll29_)
                d_907_v23_ = (d_890_cf3_)[default__.safeIndex((len(iife67_()
                ) if (d_874_v1_)[default__.safeIndex(89, (d_874_v1_).length(0))] else p0), len(d_890_cf3_))]
                d_910_v26_: _dafny.Map
                d_910_v26_ = _dafny.Map({not(d_892_cf1_): -362})
                d_911_v27_: _dafny.Map
                d_911_v27_ = _dafny.Map({p0: p2})
                d_912_v28_: _dafny.Array
                nw124_ = _dafny.Array(None, 11)
                nw124_[int(0)] = p0
                nw124_[int(1)] = len(d_910_v26_)
                nw124_[int(2)] = (0) - (p2)
                nw124_[int(3)] = 578
                nw124_[int(4)] = p0
                nw124_[int(5)] = p0
                nw124_[int(6)] = p2
                nw124_[int(7)] = d_876_v3_
                nw124_[int(8)] = len((_dafny.Map({d_876_v3_: d_876_v3_})) | (d_911_v27_))
                nw124_[int(9)] = (0) - (d_876_v3_)
                nw124_[int(10)] = default__.safeModuloInt(p0, p2)
                d_912_v28_ = nw124_
                index133_ = default__.safeIndex(744, (d_912_v28_).length(0))
                (d_912_v28_)[index133_] = p0
                d_913_v29_: D0
                d_913_v29_ = D0_DC0(p2)
                index134_ = default__.safeIndex(112, (d_912_v28_).length(0))
                (d_912_v28_)[index134_] = (-951) * (d_876_v3_)
                d_914_v30_: _dafny.MultiSet
                d_914_v30_ = _dafny.MultiSet([d_874_v1_])
                d_915_v31_: _dafny.Seq
                d_915_v31_ = _dafny.SeqWithoutIsStrInference([not(d_875_v2_)])
                d_916_v32_: _dafny.Map
                d_916_v32_ = _dafny.Map({d_915_v31_: d_876_v3_})
                index135_ = default__.safeIndex(744, (d_912_v28_).length(0))
                index136_ = default__.safeIndex(112, (d_912_v28_).length(0))
                rhs145_ = default__.fm19((0) - (d_876_v3_), d_875_v2_, globalState)
                rhs146_ = p2
                rhs147_ = D0_DC0((0) - (p0))
                rhs148_ = ((d_916_v32_)[(d_915_v31_) + (d_915_v31_)] if ((d_915_v31_) + (d_915_v31_)) in (d_916_v32_) else len(d_911_v27_))
                rhs149_ = d_914_v30_
                lhs102_ = globalState
                lhs103_ = d_912_v28_
                lhs104_ = default__.safeIndex(744, (d_912_v28_).length(0))
                lhs105_ = d_912_v28_
                lhs106_ = default__.safeIndex(112, (d_912_v28_).length(0))
                lhs102_.f3 = rhs145_
                lhs103_[lhs104_] = rhs146_
                d_913_v29_ = rhs147_
                lhs105_[lhs106_] = rhs148_
                d_914_v30_ = rhs149_
                d_917_v33_: D0
                d_917_v33_ = D0_DC2(d_892_cf1_)
                (globalState).f9 = ((self).fm5(d_917_v33_, globalState)) > (len(d_906_cf17_))
                index137_ = default__.safeIndex(744, (d_912_v28_).length(0))
                (d_912_v28_)[index137_] = default__.fm11((d_874_v1_)[default__.safeIndex(89, (d_874_v1_).length(0))], globalState)
            index138_ = default__.safeIndex(89, (d_874_v1_).length(0))
            (d_874_v1_)[index138_] = (d_874_v1_)[default__.safeIndex(89, (d_874_v1_).length(0))]
        elif source12_.is_DC2:
            d_918___mcc_h3_ = source12_.cf4
            d_919_cf4_ = d_918___mcc_h3_
            d_920_v35_: _dafny.Seq
            d_920_v35_ = _dafny.SeqWithoutIsStrInference([d_876_v3_, p2])
            d_921_v36_: _dafny.Map
            d_921_v36_ = _dafny.Map({len(d_920_v35_): p0})
            d_922_v37_: C4
            nw125_ = C4()
            def iife68_():
                coll30_ = _dafny.Map()
                compr_30_: str
                for compr_30_ in ((p1).fm1(d_921_v36_, globalState)).Elements:
                    d_923_v34_: str = compr_30_
                    if (d_923_v34_) in ((p1).fm1(d_921_v36_, globalState)):
                        coll30_[d_923_v34_] = (0) - (p2)
                return _dafny.Map(coll30_)
            nw125_.ctor__(13, iife68_()
            )
            d_922_v37_ = nw125_
            d_919_cf4_ = not((d_874_v1_)[default__.safeIndex(89, (d_874_v1_).length(0))])
            index139_ = default__.safeIndex(89, (d_874_v1_).length(0))
            (d_874_v1_)[index139_] = d_875_v2_
            d_876_v3_ = (d_922_v37_).f14
        elif source12_.is_DC0:
            d_924___mcc_h4_ = source12_.cf0
            d_925_cf0_ = d_924___mcc_h4_
            d_926_v38_: str
            d_926_v38_ = _dafny.CodePoint('j')
            index140_ = default__.safeIndex(869, (d_872_v0_).length(0))
            (d_872_v0_)[index140_] = d_926_v38_
            index141_ = default__.safeIndex(869, (d_872_v0_).length(0))
            (d_872_v0_)[index141_] = d_926_v38_
            (globalState).f3 = (d_874_v1_)[default__.safeIndex(89, (d_874_v1_).length(0))]
            def iife69_():
                coll31_ = _dafny.Set()
                compr_31_: int
                for compr_31_ in _dafny.IntegerRange(21, -521):
                    d_927_v39_: int = compr_31_
                    if ((21) <= (d_927_v39_)) and ((d_927_v39_) < (-521)):
                        coll31_ = coll31_.union(_dafny.Set([default__.safeDivisionInt(d_927_v39_, d_876_v3_)]))
                return _dafny.Set(coll31_)
            source14_ = D4_DC10((iife69_()
).intersection(d_882_v9_))
            if source14_.is_DC11:
                d_928_v40_: _dafny.Seq
                d_928_v40_ = _dafny.SeqWithoutIsStrInference([(0) - (p0), p0, len(d_883_v10_)])
                d_929_v41_: _dafny.Map
                d_929_v41_ = _dafny.Map({d_875_v2_: d_928_v40_})
                d_930_v42_: D6
                d_930_v42_ = D6_DC16(len(d_928_v40_))
                d_931_v43_: D0
                d_931_v43_ = D0_DC0(p0)
                d_932_v44_: _dafny.Array
                nw126_ = _dafny.Array(None, 21)
                nw126_[int(0)] = d_928_v40_
                nw126_[int(1)] = ((d_929_v41_)[d_875_v2_] if (d_875_v2_) in (d_929_v41_) else d_928_v40_)
                nw126_[int(2)] = (_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vcqodo")))) for d_933_i2_ in range(default__.abs(919))])) + (d_928_v40_)
                nw126_[int(3)] = d_928_v40_
                nw126_[int(4)] = (d_928_v40_).set(default__.safeIndex(p0, len(d_928_v40_)), (_dafny.MultiSet([d_875_v2_, (d_874_v1_)[default__.safeIndex(89, (d_874_v1_).length(0))], d_875_v2_])).cardinality)
                nw126_[int(5)] = d_928_v40_
                nw126_[int(6)] = _dafny.SeqWithoutIsStrInference([8 for d_934_i3_ in range(default__.abs(-667))])
                nw126_[int(7)] = _dafny.SeqWithoutIsStrInference([d_925_cf0_ for d_935_i4_ in range(default__.abs(268))])
                nw126_[int(8)] = d_928_v40_
                nw126_[int(9)] = d_928_v40_
                nw126_[int(10)] = d_928_v40_
                nw126_[int(11)] = d_928_v40_
                nw126_[int(12)] = d_928_v40_
                nw126_[int(13)] = _dafny.SeqWithoutIsStrInference([(d_930_v42_).cf24, p2, d_925_cf0_])
                nw126_[int(14)] = d_928_v40_
                nw126_[int(15)] = (d_928_v40_ if d_875_v2_ else d_928_v40_)
                nw126_[int(16)] = (_dafny.SeqWithoutIsStrInference([d_876_v3_ for d_936_i5_ in range(default__.abs(895))])) + (d_928_v40_)
                nw126_[int(17)] = (_dafny.SeqWithoutIsStrInference([p0, d_925_cf0_, (d_931_v43_).cf0])) + (d_928_v40_)
                nw126_[int(18)] = (d_928_v40_) + (d_928_v40_)
                nw126_[int(19)] = d_928_v40_
                nw126_[int(20)] = (default__.fm30(globalState)) + ((d_928_v40_).set(default__.safeIndex(p0, len(d_928_v40_)), len(default__.fm20(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ybsadm"))), d_875_v2_, d_875_v2_, d_876_v3_, globalState))))
                d_932_v44_ = nw126_
                index142_ = default__.safeIndex(947, (d_932_v44_).length(0))
                (d_932_v44_)[index142_] = d_928_v40_
                index143_ = default__.safeIndex(947, (d_932_v44_).length(0))
                (d_932_v44_)[index143_] = _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([547, p0])).cardinality, d_925_cf0_, d_876_v3_])
                d_937_v46_: _dafny.MultiSet
                d_937_v46_ = _dafny.MultiSet([d_926_v38_])
                d_938_v47_: _dafny.Seq
                d_938_v47_ = _dafny.SeqWithoutIsStrInference([not(not(((d_878_v5_)[p2] if (p2) in (d_878_v5_) else (d_874_v1_)[default__.safeIndex(89, (d_874_v1_).length(0))]))), d_875_v2_])
                def iife70_():
                    coll32_ = _dafny.Map()
                    compr_32_: str
                    for compr_32_ in ((d_937_v46_).set((d_872_v0_)[default__.safeIndex(869, (d_872_v0_).length(0))], default__.abs(p2))).Elements:
                        d_939_v45_: str = compr_32_
                        if (d_939_v45_) in ((d_937_v46_).set((d_872_v0_)[default__.safeIndex(869, (d_872_v0_).length(0))], default__.abs(p2))):
                            coll32_[d_939_v45_] = _dafny.Map({(d_874_v1_)[default__.safeIndex(89, (d_874_v1_).length(0))]: (d_874_v1_)[default__.safeIndex(89, (d_874_v1_).length(0))]})
                    return _dafny.Map(coll32_)
                rhs150_ = (default__.fm23(not(not(d_875_v2_)), iife70_()
                , d_875_v2_, p2, globalState)) not in (d_877_v4_)
                rhs151_ = (len(d_938_v47_)) + (d_876_v3_)
                rhs152_ = 831
                lhs107_ = globalState
                lhs107_.f9 = rhs150_
                d_876_v3_ = rhs151_
                d_876_v3_ = rhs152_
                (globalState).f13 = (d_875_v2_) and (not((d_874_v1_)[default__.safeIndex(89, (d_874_v1_).length(0))]))
                d_940_v48_: D7
                d_940_v48_ = D7_DC18(d_878_v5_)
                d_941_v49_: _dafny.Map
                d_941_v49_ = _dafny.Map({d_940_v48_: p1})
                d_876_v3_ = len(_dafny.Set({(d_940_v48_) not in (d_941_v49_), False}))
            elif source14_.is_DC12:
                d_942___mcc_h8_ = source14_.cf18
                d_943_cf18_ = d_942___mcc_h8_
                (globalState).f3 = ((d_877_v4_) + (_dafny.SeqWithoutIsStrInference([(d_872_v0_)[default__.safeIndex(869, (d_872_v0_).length(0))] for d_944_i6_ in range(default__.abs(770))]))) <= (d_877_v4_)
                d_880_v7_ = D7_DC21(d_879_v6_)
                index144_ = default__.safeIndex(89, (d_874_v1_).length(0))
                (d_874_v1_)[index144_] = (d_874_v1_)[default__.safeIndex(89, (d_874_v1_).length(0))]
                d_945_v50_: D4
                d_945_v50_ = D4_DC11()
                d_946_v51_: _dafny.Array
                nw127_ = _dafny.Array(int(0), 1)
                d_946_v51_ = nw127_
                index145_ = default__.safeIndex(573, (d_946_v51_).length(0))
                (d_946_v51_)[index145_] = d_925_cf0_
                index146_ = default__.safeIndex(573, (d_946_v51_).length(0))
                rhs153_ = not((d_874_v1_)[default__.safeIndex(89, (d_874_v1_).length(0))])
                rhs154_ = d_945_v50_
                rhs155_ = len((_dafny.SeqWithoutIsStrInference([d_926_v38_ for d_947_i7_ in range(default__.abs(957))])) + (((d_877_v4_).set(default__.safeIndex(d_943_cf18_, len(d_877_v4_)), _dafny.CodePoint('u'))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qwrmofy")))))
                lhs108_ = globalState
                lhs109_ = d_946_v51_
                lhs110_ = default__.safeIndex(573, (d_946_v51_).length(0))
                lhs108_.f9 = rhs153_
                d_945_v50_ = rhs154_
                lhs109_[lhs110_] = rhs155_
            elif True:
                d_948___mcc_h9_ = source14_.cf17
                d_949_cf17_ = d_948___mcc_h9_
                d_950_v52_: _dafny.Array
                nw128_ = _dafny.Array(_dafny.Map({}), 12)
                d_950_v52_ = nw128_
                d_950_v52_ = d_950_v52_
                d_951_v53_: _dafny.Map
                d_951_v53_ = _dafny.Map({d_876_v3_: d_877_v4_})
                d_952_v54_: _dafny.Map
                d_952_v54_ = _dafny.Map({d_874_v1_: 198})
                d_953_v55_: _dafny.Map
                d_953_v55_ = _dafny.Map({len(d_952_v54_): p2})
                d_951_v53_ = (d_951_v53_).set((0) - (p0), (p1).fm1(d_953_v55_, globalState))
                d_954_v56_: _dafny.Array
                def lambda36_(d_955_v2_, d_956_v12_, d_957_v4_, d_958_v38_):
                    def lambda37_(d_959_i8_):
                        return D2_DC8(D2_DC7(d_955_v2_, _dafny.Map({d_955_v2_: d_956_v12_}), len((d_957_v4_).set(default__.safeIndex(len(d_957_v4_), len(d_957_v4_)), d_958_v38_)), d_955_v2_))

                    return lambda37_

                init17_ = lambda36_(d_875_v2_, d_886_v12_, d_877_v4_, d_926_v38_)
                nw129_ = _dafny.Array(None, 16)
                for i0_17_ in range(nw129_.length(0)):
                    nw129_[i0_17_] = init17_(i0_17_)
                d_954_v56_ = nw129_
                d_960_v57_: _dafny.Map
                d_960_v57_ = _dafny.Map({(d_875_v2_) == ((d_874_v1_)[default__.safeIndex(89, (d_874_v1_).length(0))]): d_954_v56_})
                d_960_v57_ = (d_960_v57_).set((d_875_v2_ if d_875_v2_ else d_875_v2_), d_954_v56_)
                d_961_v58_: _dafny.Array
                nw130_ = _dafny.Array(int(0), 18)
                d_961_v58_ = nw130_
                d_962_v59_: _dafny.Map
                d_962_v59_ = _dafny.Map({(d_874_v1_)[default__.safeIndex(89, (d_874_v1_).length(0))]: d_961_v58_})
                d_961_v58_ = ((d_962_v59_)[(d_874_v1_)[default__.safeIndex(89, (d_874_v1_).length(0))]] if ((d_874_v1_)[default__.safeIndex(89, (d_874_v1_).length(0))]) in (d_962_v59_) else d_961_v58_)
            d_925_cf0_ = (0) - (p0)
        elif True:
            d_963___mcc_h5_ = source12_.cf5
            d_964_cf5_ = d_963___mcc_h5_
            d_965_v60_: str
            d_965_v60_ = _dafny.CodePoint('t')
            d_966_v61_: _dafny.Map
            d_966_v61_ = _dafny.Map({d_965_v60_: d_877_v4_})
            index147_ = default__.safeIndex(89, (d_874_v1_).length(0))
            (d_874_v1_)[index147_] = (_dafny.CodePoint('s')) in (d_966_v61_)
            (globalState).f9 = d_875_v2_
            d_967_v62_: _dafny.Seq
            d_967_v62_ = _dafny.SeqWithoutIsStrInference([d_877_v4_])
            d_968_v63_: _dafny.Map
            d_968_v63_ = _dafny.Map({d_876_v3_: d_877_v4_})
            d_969_v64_: D0
            d_969_v64_ = D0_DC1((d_874_v1_)[default__.safeIndex(89, (d_874_v1_).length(0))], d_967_v62_, ((d_968_v63_)[p0] if (p0) in (d_968_v63_) else d_877_v4_))
            d_970_v65_: _dafny.MultiSet
            d_970_v65_ = _dafny.MultiSet([d_969_v64_])
            d_971_v66_: D2
            d_971_v66_ = D2_DC6(d_970_v65_)
            pat_let_tv21_ = d_970_v65_
            d_972_v67_: D2
            def iife71_(_pat_let19_0):
                def iife72_(d_973_dt__update__tmp_h0_):
                    def iife73_(_pat_let20_0):
                        def iife74_(d_974_dt__update_hcf10_h0_):
                            return D2_DC6(d_974_dt__update_hcf10_h0_)
                        return iife74_(_pat_let20_0)
                    return iife73_(pat_let_tv21_)
                return iife72_(_pat_let19_0)
            d_972_v67_ = D2_DC6((iife71_(d_971_v66_)).cf10)
            source15_ = d_971_v66_
            if source15_.is_DC7:
                d_975___mcc_h10_ = source15_.cf11
                d_976___mcc_h11_ = source15_.cf12
                d_977___mcc_h12_ = source15_.cf13
                d_978___mcc_h13_ = source15_.cf14
                d_979_cf14_ = d_978___mcc_h13_
                d_980_cf13_ = d_977___mcc_h12_
                d_981_cf12_ = d_976___mcc_h11_
                d_982_cf11_ = d_975___mcc_h10_
                d_983_v68_: _dafny.Array
                def lambda38_(d_984_cf13_):
                    def lambda39_(d_985_i9_):
                        return (d_985_i9_) - (d_984_cf13_)

                    return lambda39_

                init18_ = lambda38_(d_980_cf13_)
                nw131_ = _dafny.Array(None, 29)
                for i0_18_ in range(nw131_.length(0)):
                    nw131_[i0_18_] = init18_(i0_18_)
                d_983_v68_ = nw131_
                index148_ = default__.safeIndex(110, (d_983_v68_).length(0))
                (d_983_v68_)[index148_] = d_980_cf13_
                index149_ = default__.safeIndex(678, (d_983_v68_).length(0))
                (d_983_v68_)[index149_] = p2
                d_986_v69_: _dafny.Map
                d_986_v69_ = _dafny.Map({d_874_v1_: d_979_cf14_})
                d_987_v70_: _dafny.Map
                d_987_v70_ = _dafny.Map({d_982_cf11_: d_980_cf13_})
                index150_ = default__.safeIndex(110, (d_983_v68_).length(0))
                index151_ = default__.safeIndex(678, (d_983_v68_).length(0))
                rhs156_ = (0) - (default__.safeDivisionInt(len((_dafny.Map({d_874_v1_: d_875_v2_})) | (d_986_v69_)), default__.safeDivisionInt(-701, ((d_987_v70_)[d_875_v2_] if (d_875_v2_) in (d_987_v70_) else 214))))
                rhs157_ = (d_874_v1_)[default__.safeIndex(89, (d_874_v1_).length(0))]
                rhs158_ = len(d_877_v4_)
                rhs159_ = d_876_v3_
                lhs111_ = d_983_v68_
                lhs112_ = default__.safeIndex(110, (d_983_v68_).length(0))
                lhs113_ = d_983_v68_
                lhs114_ = default__.safeIndex(678, (d_983_v68_).length(0))
                d_980_cf13_ = rhs156_
                d_875_v2_ = rhs157_
                lhs111_[lhs112_] = rhs158_
                lhs113_[lhs114_] = rhs159_
                d_874_v1_ = d_874_v1_
                d_988_v71_: _dafny.Seq
                d_988_v71_ = _dafny.SeqWithoutIsStrInference([d_982_cf11_])
                d_989_v72_: _dafny.Seq
                d_989_v72_ = _dafny.SeqWithoutIsStrInference([d_988_v71_])
                d_990_v73_: D2
                d_990_v73_ = D2_DC7((d_874_v1_)[default__.safeIndex(89, (d_874_v1_).length(0))], d_981_cf12_, default__.fm11(d_979_cf14_, globalState), d_982_cf11_)
                pat_let_tv22_ = d_981_cf12_
                def iife75_(_pat_let21_0):
                    def iife76_(d_991_dt__update__tmp_h1_):
                        def iife77_(_pat_let22_0):
                            def iife78_(d_992_dt__update_hcf12_h0_):
                                return D2_DC7((d_991_dt__update__tmp_h1_).cf11, d_992_dt__update_hcf12_h0_, (d_991_dt__update__tmp_h1_).cf13, (d_991_dt__update__tmp_h1_).cf14)
                            return iife78_(_pat_let22_0)
                        return iife77_(pat_let_tv22_)
                    return iife76_(_pat_let21_0)
                d_980_cf13_ = (143) + (len((d_989_v72_)[default__.safeIndex((iife75_(d_990_v73_)).cf13, len(d_989_v72_))]))
                d_993_v74_: C1
                nw132_ = C1()
                nw132_.ctor__()
                d_993_v74_ = nw132_
            elif source15_.is_DC6:
                d_994___mcc_h14_ = source15_.cf10
                d_995_cf10_ = d_994___mcc_h14_
                d_996_v75_: C1
                nw133_ = C1()
                nw133_.ctor__()
                d_996_v75_ = nw133_
                (globalState).f13 = True
                d_997_v76_: D2
                d_997_v76_ = D2_DC7(d_875_v2_, _dafny.Map({d_875_v2_: D0_DC3(d_885_v11_)}), p0, d_875_v2_)
                d_998_v77_: D2
                d_998_v77_ = D2_DC8(d_997_v76_)
                d_999_v78_: _dafny.Map
                d_999_v78_ = _dafny.Map({d_998_v77_: p2})
                d_1000_v79_: _dafny.Map
                d_1000_v79_ = _dafny.Map({d_875_v2_: d_876_v3_})
                d_1001_v80_: _dafny.Map
                d_1001_v80_ = _dafny.Map({(d_874_v1_)[default__.safeIndex(89, (d_874_v1_).length(0))]: (d_874_v1_)[default__.safeIndex(89, (d_874_v1_).length(0))]})
                d_999_v78_ = (d_999_v78_).set(D2_DC8(d_997_v76_), ((d_1000_v79_)[((d_1001_v80_)[(d_874_v1_)[default__.safeIndex(89, (d_874_v1_).length(0))]] if ((d_874_v1_)[default__.safeIndex(89, (d_874_v1_).length(0))]) in (d_1001_v80_) else d_875_v2_)] if (((d_1001_v80_)[(d_874_v1_)[default__.safeIndex(89, (d_874_v1_).length(0))]] if ((d_874_v1_)[default__.safeIndex(89, (d_874_v1_).length(0))]) in (d_1001_v80_) else d_875_v2_)) in (d_1000_v79_) else p2))
                (globalState).f9 = d_875_v2_
            elif True:
                d_1002___mcc_h15_ = source15_.cf15
                d_1003_cf15_ = d_1002___mcc_h15_
                d_1004_v81_: _dafny.Map
                d_1004_v81_ = _dafny.Map({d_875_v2_: d_886_v12_})
                d_1005_v82_: D2
                d_1005_v82_ = D2_DC7(False, d_1004_v81_, (0) - (p0), d_875_v2_)
                d_1006_v83_: _dafny.Map
                d_1006_v83_ = _dafny.Map({d_965_v60_: d_875_v2_})
                d_876_v3_ = (0) - (((default__.fm22(d_1005_v82_, d_886_v12_, ((d_1006_v83_)[d_965_v60_] if (d_965_v60_) in (d_1006_v83_) else d_875_v2_), globalState)) * (p2)) + (p2))
                d_1007_v84_: _dafny.Seq
                d_1007_v84_ = _dafny.SeqWithoutIsStrInference([d_875_v2_, d_875_v2_, d_875_v2_])
                d_1008_v85_: _dafny.Map
                d_1008_v85_ = _dafny.Map({p0: _dafny.CodePoint('k')})
                d_1009_v86_: _dafny.Map
                d_1009_v86_ = _dafny.Map({(d_874_v1_)[default__.safeIndex(89, (d_874_v1_).length(0))]: (d_874_v1_)[default__.safeIndex(89, (d_874_v1_).length(0))]})
                d_1010_v87_: _dafny.Map
                d_1010_v87_ = _dafny.Map({d_875_v2_: not(((d_1009_v86_)[(d_874_v1_)[default__.safeIndex(89, (d_874_v1_).length(0))]] if ((d_874_v1_)[default__.safeIndex(89, (d_874_v1_).length(0))]) in (d_1009_v86_) else default__.fm6(globalState)))})
                d_1011_v88_: _dafny.Seq
                d_1011_v88_ = _dafny.SeqWithoutIsStrInference([p0, -878, d_876_v3_, p0, len(d_1010_v87_)])
                d_1012_v89_: _dafny.Array
                nw134_ = _dafny.Array(None, 13)
                nw134_[int(0)] = (0) - (len(d_1007_v84_))
                nw134_[int(1)] = (d_876_v3_) - (-518)
                nw134_[int(2)] = d_876_v3_
                nw134_[int(3)] = d_876_v3_
                nw134_[int(4)] = len(d_1008_v85_)
                nw134_[int(5)] = d_876_v3_
                nw134_[int(6)] = (d_1011_v88_)[default__.safeIndex(p2, len(d_1011_v88_))]
                nw134_[int(7)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uvsvmpmou")))
                nw134_[int(8)] = p0
                nw134_[int(9)] = (0) - (-64)
                nw134_[int(10)] = p0
                nw134_[int(11)] = (0) - ((0) - (default__.safeModuloInt(p2, d_876_v3_)))
                nw134_[int(12)] = default__.safeModuloInt(p0, p2)
                d_1012_v89_ = nw134_
                index152_ = default__.safeIndex(89, (d_874_v1_).length(0))
                rhs160_ = True
                rhs161_ = d_1012_v89_
                rhs162_ = d_875_v2_
                rhs163_ = (d_965_v60_) not in ((_dafny.SeqWithoutIsStrInference([d_965_v60_ for d_1013_i10_ in range(default__.abs(893))])) + (d_877_v4_))
                rhs164_ = p2
                lhs115_ = globalState
                lhs116_ = globalState
                lhs117_ = d_874_v1_
                lhs118_ = default__.safeIndex(89, (d_874_v1_).length(0))
                lhs115_.f3 = rhs160_
                d_1012_v89_ = rhs161_
                lhs116_.f5 = rhs162_
                lhs117_[lhs118_] = rhs163_
                d_876_v3_ = rhs164_
                (globalState).f3 = (d_874_v1_)[default__.safeIndex(89, (d_874_v1_).length(0))]
                d_1014_v90_: _dafny.Array
                nw135_ = _dafny.Array(_dafny.Array(None, 0), 10)
                d_1014_v90_ = nw135_
                index153_ = default__.safeIndex(624, (d_1014_v90_).length(0))
                (d_1014_v90_)[index153_] = d_1012_v89_
                index154_ = default__.safeIndex(624, (d_1014_v90_).length(0))
                (d_1014_v90_)[index154_] = d_1012_v89_
            d_1015_v91_: _dafny.Array
            nw136_ = _dafny.Array(int(0), 4)
            d_1015_v91_ = nw136_
            index155_ = default__.safeIndex(43, (d_1015_v91_).length(0))
            (d_1015_v91_)[index155_] = default__.fm11((d_874_v1_)[default__.safeIndex(89, (d_874_v1_).length(0))], globalState)
            index156_ = default__.safeIndex(43, (d_1015_v91_).length(0))
            (d_1015_v91_)[index156_] = p0

    def m4(self, p0, p1, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: int = int(0)
        d_1016_v0_: _dafny.Seq
        d_1016_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "scsqu"))
        d_1016_v0_ = d_1016_v0_
        d_1017_v1_: bool
        d_1017_v1_ = False
        r2 = (self).fm5(D0_DC2(d_1017_v1_), globalState)
        d_1018_v2_: int
        d_1018_v2_ = 931
        hi6_ = d_1018_v2_
        for d_1019_i0_ in range(len(d_1016_v0_), hi6_):
            r2 = (d_1018_v2_) * (d_1019_i0_)
            d_1020_v3_: _dafny.Map
            d_1020_v3_ = _dafny.Map({d_1018_v2_: default__.fm6(globalState)})
            d_1021_v4_: _dafny.Array
            nw137_ = _dafny.Array(_dafny.Map({}), 18)
            d_1021_v4_ = nw137_
            d_1022_v5_: _dafny.Array
            nw138_ = _dafny.Array(None, 14)
            nw138_[int(0)] = d_1017_v1_
            nw138_[int(1)] = not(d_1017_v1_)
            nw138_[int(2)] = d_1017_v1_
            nw138_[int(3)] = (d_1017_v1_) and (d_1017_v1_)
            nw138_[int(4)] = (d_1018_v2_) <= (d_1018_v2_)
            nw138_[int(5)] = False
            nw138_[int(6)] = (d_1018_v2_) < (807)
            nw138_[int(7)] = d_1017_v1_
            nw138_[int(8)] = d_1017_v1_
            nw138_[int(9)] = d_1017_v1_
            nw138_[int(10)] = (((d_1020_v3_)[350] if (350) in (d_1020_v3_) else d_1017_v1_)) == (d_1017_v1_)
            nw138_[int(11)] = d_1017_v1_
            nw138_[int(12)] = False
            nw138_[int(13)] = False
            d_1022_v5_ = nw138_
            def iife79_():
                coll33_ = _dafny.Map()
                compr_33_: int
                for compr_33_ in _dafny.IntegerRange(791, -974):
                    d_1023_v6_: int = compr_33_
                    if ((791) <= (d_1023_v6_)) and ((d_1023_v6_) < (-974)):
                        coll33_[(d_1023_v6_) - (d_1018_v2_)] = d_1017_v1_
                return _dafny.Map(coll33_)
            rhs165_ = (iife79_()
            ).set(d_1019_i0_, d_1017_v1_)
            rhs166_ = d_1019_i0_
            rhs167_ = d_1021_v4_
            rhs168_ = (d_1022_v5_ if not(not(not(default__.fm6(globalState)))) else p1)
            d_1020_v3_ = rhs165_
            r1 = rhs166_
            d_1021_v4_ = rhs167_
            d_1022_v5_ = rhs168_
            r2 = d_1019_i0_
            d_1024_v7_: _dafny.Array
            nw139_ = _dafny.Array(int(0), 26)
            d_1024_v7_ = nw139_
            index157_ = default__.safeIndex(474, (d_1024_v7_).length(0))
            (d_1024_v7_)[index157_] = d_1018_v2_
            d_1025_v8_: _dafny.Seq
            d_1025_v8_ = _dafny.SeqWithoutIsStrInference([d_1017_v1_])
            d_1026_v9_: D1
            d_1026_v9_ = D1_DC5(d_1017_v1_, d_1016_v0_, d_1025_v8_)
            d_1027_v10_: _dafny.Seq
            d_1027_v10_ = _dafny.SeqWithoutIsStrInference([d_1017_v1_, not((d_1026_v9_).cf7)])
            index158_ = default__.safeIndex(474, (d_1024_v7_).length(0))
            (d_1024_v7_)[index158_] = len(((((_dafny.SeqWithoutIsStrInference([d_1017_v1_, d_1017_v1_, default__.fm6(globalState)])) + (d_1025_v8_)).set(default__.safeIndex((0) - (d_1018_v2_), len((_dafny.SeqWithoutIsStrInference([d_1017_v1_, d_1017_v1_, default__.fm6(globalState)])) + (d_1025_v8_))), d_1017_v1_)).set(default__.safeIndex(d_1018_v2_, len(((_dafny.SeqWithoutIsStrInference([d_1017_v1_, d_1017_v1_, default__.fm6(globalState)])) + (d_1025_v8_)).set(default__.safeIndex((0) - (d_1018_v2_), len((_dafny.SeqWithoutIsStrInference([d_1017_v1_, d_1017_v1_, default__.fm6(globalState)])) + (d_1025_v8_))), d_1017_v1_))), d_1017_v1_)) + (d_1027_v10_))
        d_1028_i1_: int
        d_1028_i1_ = 0
        with _dafny.label("7"):
            while not(default__.fm6(globalState)):
                with _dafny.c_label("7"):
                    if (d_1028_i1_) >= (100):
                        raise _dafny.Break("7")
                    d_1028_i1_ = (d_1028_i1_) + (1)
                    d_1029_v11_: _dafny.Seq
                    d_1029_v11_ = _dafny.SeqWithoutIsStrInference([d_1017_v1_])
                    d_1030_v12_: str
                    d_1030_v12_ = _dafny.CodePoint('c')
                    source16_ = default__.fm35(((d_1029_v11_) + (d_1029_v11_)).set(default__.safeIndex(d_1018_v2_, len((d_1029_v11_) + (d_1029_v11_))), d_1017_v1_), not (d_1017_v1_) or (d_1017_v1_), (_dafny.SeqWithoutIsStrInference([(d_1016_v0_)[default__.safeIndex(d_1018_v2_, len(d_1016_v0_))] for d_1031_i2_ in range(default__.abs(703))])).set(default__.safeIndex(d_1018_v2_, len(_dafny.SeqWithoutIsStrInference([(d_1016_v0_)[default__.safeIndex(d_1018_v2_, len(d_1016_v0_))] for d_1032_i2_ in range(default__.abs(703))]))), d_1030_v12_), globalState)
                    if source16_.is_DC11:
                        d_1016_v0_ = d_1016_v0_
                        d_1033_v13_: _dafny.MultiSet
                        d_1033_v13_ = _dafny.MultiSet([d_1029_v11_, (d_1029_v11_) + (d_1029_v11_), d_1029_v11_])
                        rhs169_ = (p0)[default__.safeIndex(d_1018_v2_, len(p0))]
                        rhs170_ = (_dafny.MultiSet([d_1029_v11_, d_1029_v11_, d_1029_v11_])) - (d_1033_v13_)
                        rhs171_ = (0) - (-497)
                        r2 = rhs169_
                        d_1033_v13_ = rhs170_
                        d_1018_v2_ = rhs171_
                        d_1034_v14_: _dafny.MultiSet
                        d_1034_v14_ = _dafny.MultiSet([d_1016_v0_])
                        rhs172_ = default__.safeDivisionInt((d_1018_v2_ if default__.fm6(globalState) else (0) - (d_1018_v2_)), d_1018_v2_)
                        rhs173_ = _dafny.MultiSet([d_1016_v0_])
                        rhs174_ = d_1018_v2_
                        rhs175_ = default__.fm6(globalState)
                        rhs176_ = not(not(d_1017_v1_))
                        lhs119_ = globalState
                        lhs120_ = globalState
                        r2 = rhs172_
                        d_1034_v14_ = rhs173_
                        r1 = rhs174_
                        lhs119_.f3 = rhs175_
                        lhs120_.f3 = rhs176_
                        d_1035_v15_: _dafny.Map
                        d_1035_v15_ = _dafny.Map({d_1017_v1_: d_1017_v1_})
                        d_1036_v16_: _dafny.Map
                        d_1036_v16_ = _dafny.Map({d_1030_v12_: d_1035_v15_})
                        d_1037_v17_: _dafny.Array
                        nw140_ = _dafny.Array(None, 19)
                        nw140_[int(0)] = d_1030_v12_
                        nw140_[int(1)] = d_1030_v12_
                        nw140_[int(2)] = d_1030_v12_
                        nw140_[int(3)] = d_1030_v12_
                        nw140_[int(4)] = d_1030_v12_
                        nw140_[int(5)] = d_1030_v12_
                        nw140_[int(6)] = d_1030_v12_
                        nw140_[int(7)] = d_1030_v12_
                        nw140_[int(8)] = d_1030_v12_
                        nw140_[int(9)] = d_1030_v12_
                        nw140_[int(10)] = d_1030_v12_
                        nw140_[int(11)] = _dafny.CodePoint('a')
                        nw140_[int(12)] = d_1030_v12_
                        nw140_[int(13)] = d_1030_v12_
                        nw140_[int(14)] = d_1030_v12_
                        nw140_[int(15)] = default__.fm23(d_1017_v1_, d_1036_v16_, d_1017_v1_, d_1018_v2_, globalState)
                        nw140_[int(16)] = d_1030_v12_
                        nw140_[int(17)] = d_1030_v12_
                        nw140_[int(18)] = _dafny.CodePoint('o')
                        d_1037_v17_ = nw140_
                        d_1038_v20_: _dafny.Set
                        d_1038_v20_ = _dafny.Set({d_1018_v2_})
                        d_1039_v21_: _dafny.Map
                        def iife80_():
                            coll34_ = _dafny.Map()
                            compr_34_: int
                            for compr_34_ in (d_1038_v20_).Elements:
                                d_1040_v19_: int = compr_34_
                                if (d_1040_v19_) in (d_1038_v20_):
                                    coll34_[default__.safeDivisionInt(d_1040_v19_, d_1018_v2_)] = d_1017_v1_
                            return _dafny.Map(coll34_)
                        d_1039_v21_ = _dafny.Map({d_1030_v12_: iife80_()
                        })
                        index159_ = default__.safeIndex(240, (d_1037_v17_).length(0))
                        def iife81_():
                            coll35_ = _dafny.Map()
                            compr_35_: str
                            for compr_35_ in (d_1039_v21_).keys.Elements:
                                d_1041_v18_: str = compr_35_
                                if (d_1041_v18_) in (d_1039_v21_):
                                    coll35_[d_1041_v18_] = (_dafny.Map({d_1017_v1_: d_1017_v1_})).set(d_1017_v1_, d_1017_v1_)
                            return _dafny.Map(coll35_)
                        (d_1037_v17_)[index159_] = default__.fm23(d_1017_v1_, iife81_()
                        , d_1017_v1_, d_1018_v2_, globalState)
                        index160_ = default__.safeIndex(240, (d_1037_v17_).length(0))
                        (d_1037_v17_)[index160_] = d_1030_v12_
                    elif source16_.is_DC12:
                        d_1042___mcc_h0_ = source16_.cf18
                        d_1043_cf18_ = d_1042___mcc_h0_
                        (globalState).f3 = d_1017_v1_
                        r2 = default__.fm11(d_1017_v1_, globalState)
                        (globalState).f13 = True
                        d_1017_v1_ = d_1017_v1_
                    elif True:
                        d_1044___mcc_h1_ = source16_.cf17
                        d_1045_cf17_ = d_1044___mcc_h1_
                        d_1046_v22_: _dafny.Map
                        d_1046_v22_ = _dafny.Map({(d_1045_cf17_).intersection(d_1045_cf17_): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bgddjiqf"))})
                        d_1046_v22_ = (d_1046_v22_).set(d_1045_cf17_, (d_1016_v0_) + (d_1016_v0_))
                        d_1047_v23_: _dafny.Set
                        d_1047_v23_ = _dafny.Set({d_1017_v1_})
                        d_1048_v24_: _dafny.Map
                        d_1048_v24_ = _dafny.Map({len(_dafny.Map({d_1017_v1_: d_1029_v11_})): (d_1018_v2_) * (len(d_1047_v23_))})
                        d_1048_v24_ = d_1048_v24_
                        d_1017_v1_ = False
                        r2 = d_1018_v2_
                    d_1049_v25_: _dafny.Array
                    def lambda40_(d_1050_v2_):
                        def lambda41_(d_1051_i3_):
                            return (d_1051_i3_) + (d_1050_v2_)

                        return lambda41_

                    init19_ = lambda40_(d_1018_v2_)
                    nw141_ = _dafny.Array(None, 3)
                    for i0_19_ in range(nw141_.length(0)):
                        nw141_[i0_19_] = init19_(i0_19_)
                    d_1049_v25_ = nw141_
                    d_1052_v26_: _dafny.Map
                    d_1052_v26_ = _dafny.Map({d_1018_v2_: len(_dafny.Map({(self).fm0(d_1018_v2_, d_1030_v12_, globalState): False}))})
                    index161_ = default__.safeIndex(903, (d_1049_v25_).length(0))
                    (d_1049_v25_)[index161_] = len(d_1052_v26_)
                    index162_ = default__.safeIndex(903, (d_1049_v25_).length(0))
                    (d_1049_v25_)[index162_] = len(d_1016_v0_)
                    d_1053_v27_: _dafny.MultiSet
                    d_1053_v27_ = _dafny.MultiSet([468])
                    d_1053_v27_ = _dafny.MultiSet([(d_1049_v25_)[default__.safeIndex(903, (d_1049_v25_).length(0))], (d_1049_v25_)[default__.safeIndex(903, (d_1049_v25_).length(0))]])
                    d_1030_v12_ = _dafny.CodePoint('a')
                    pass
            pass
        r2 = d_1018_v2_
        d_1054_v28_: _dafny.Set
        d_1054_v28_ = _dafny.Set({(d_1018_v2_) == (d_1018_v2_)})
        d_1054_v28_ = default__.fm33(d_1017_v1_, d_1017_v1_, d_1017_v1_, globalState)
        r0 = False
        d_1055_v29_: D0
        d_1055_v29_ = D0_DC0(d_1018_v2_)
        d_1056_v30_: D0
        d_1056_v30_ = D0_DC3(d_1055_v29_)
        d_1057_v31_: _dafny.Map
        d_1057_v31_ = _dafny.Map({d_1017_v1_: d_1056_v30_})
        d_1058_v32_: D2
        d_1058_v32_ = D2_DC7(d_1017_v1_, d_1057_v31_, d_1018_v2_, False)
        r1 = default__.fm22(d_1058_v32_, D0_DC3(d_1055_v29_), (len(p0)) == (229), globalState)
        r2 = d_1018_v2_
        return r0, r1, r2


class C6(T1, T0):
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C6"
    def ctor__(self):
        pass
        pass

    def fm0(self, p0, p1, globalState):
        return (default__.safeModuloInt((0) - ((0) - (-869)), (D0_DC0(211)).cf0)) * (-444)

    def fm1(self, p0, globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tiwttktph"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cd")))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aprlnhebx")))

    def fm3(self, p0, p1, p2, globalState):
        return False

    def m0(self, p0, globalState):
        r0: bool = False
        r1: int = int(0)
        d_1059_v0_: bool
        d_1059_v0_ = True
        d_1060_v1_: int
        d_1060_v1_ = 619
        d_1061_v2_: T1
        nw142_ = C3()
        nw142_.ctor__(d_1059_v0_, d_1060_v1_)
        d_1061_v2_ = nw142_
        d_1062_v3_: _dafny.Set
        d_1062_v3_ = _dafny.Set({d_1061_v2_, d_1061_v2_})
        d_1063_v4_: _dafny.Set
        d_1063_v4_ = _dafny.Set({len(_dafny.SeqWithoutIsStrInference([p0 for d_1064_i0_ in range(default__.abs(381))])), len(d_1062_v3_)})
        if ((_dafny.Set({d_1060_v1_})).isdisjoint(d_1063_v4_) if d_1059_v0_ else d_1059_v0_):
            d_1065_v5_: _dafny.Seq
            d_1065_v5_ = _dafny.SeqWithoutIsStrInference([d_1059_v0_, d_1059_v0_, d_1059_v0_])
            d_1066_v6_: _dafny.Map
            d_1067_v7_: _dafny.MultiSet
            d_1068_v8_: bool
            d_1069_v9_: _dafny.Array
            out14_: _dafny.Map
            out15_: _dafny.MultiSet
            out16_: bool
            out17_: _dafny.Array
            out14_, out15_, out16_, out17_ = (self).m2(d_1065_v5_, d_1059_v0_, globalState)
            d_1066_v6_ = out14_
            d_1067_v7_ = out15_
            d_1068_v8_ = out16_
            d_1069_v9_ = out17_
            d_1070_v10_: _dafny.Array
            nw143_ = _dafny.Array(False, 4)
            d_1070_v10_ = nw143_
            index163_ = default__.safeIndex(425, (d_1070_v10_).length(0))
            (d_1070_v10_)[index163_] = d_1059_v0_
            index164_ = default__.safeIndex(425, (d_1070_v10_).length(0))
            rhs177_ = (d_1060_v1_) > (d_1060_v1_)
            rhs178_ = (d_1060_v1_) * (d_1060_v1_)
            rhs179_ = default__.fm19(d_1060_v1_, d_1068_v8_, globalState)
            lhs121_ = globalState
            lhs122_ = d_1070_v10_
            lhs123_ = default__.safeIndex(425, (d_1070_v10_).length(0))
            lhs121_.f5 = rhs177_
            d_1060_v1_ = rhs178_
            lhs122_[lhs123_] = rhs179_
            d_1071_v11_: _dafny.Seq
            d_1071_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mcaxr"))
            d_1072_v12_: _dafny.Seq
            d_1072_v12_ = _dafny.SeqWithoutIsStrInference([d_1071_v11_])
            d_1073_v13_: D0
            d_1073_v13_ = D0_DC1(d_1068_v8_, d_1072_v12_, d_1071_v11_)
            d_1074_v14_: D2
            d_1074_v14_ = D2_DC6(_dafny.MultiSet([d_1073_v13_]))
            source17_ = d_1074_v14_
            if source17_.is_DC7:
                d_1075___mcc_h0_ = source17_.cf11
                d_1076___mcc_h1_ = source17_.cf12
                d_1077___mcc_h2_ = source17_.cf13
                d_1078___mcc_h3_ = source17_.cf14
                d_1079_cf14_ = d_1078___mcc_h3_
                d_1080_cf13_ = d_1077___mcc_h2_
                d_1081_cf12_ = d_1076___mcc_h1_
                d_1082_cf11_ = d_1075___mcc_h0_
                d_1083_v15_: _dafny.Array
                def lambda42_(d_1084_p0_, d_1085_v11_, d_1086_cf14_):
                    def lambda43_(d_1087_i1_):
                        return (_dafny.SeqWithoutIsStrInference([d_1084_p0_ for d_1088_i2_ in range(default__.abs(229))])) + ((D7_DC20(d_1085_v11_, d_1086_cf14_)).cf31)

                    return lambda43_

                init20_ = lambda42_(p0, d_1071_v11_, d_1079_cf14_)
                nw144_ = _dafny.Array(None, 24)
                for i0_20_ in range(nw144_.length(0)):
                    nw144_[i0_20_] = init20_(i0_20_)
                d_1083_v15_ = nw144_
                index165_ = default__.safeIndex(990, (d_1083_v15_).length(0))
                (d_1083_v15_)[index165_] = d_1071_v11_
                d_1089_v16_: _dafny.Map
                d_1089_v16_ = _dafny.Map({d_1082_cf11_: d_1080_cf13_})
                index166_ = default__.safeIndex(990, (d_1083_v15_).length(0))
                (d_1083_v15_)[index166_] = default__.fm20(d_1080_cf13_, d_1059_v0_, (((d_1089_v16_)[d_1082_cf11_] if (d_1082_cf11_) in (d_1089_v16_) else len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "prnfmcqd"))))) > (d_1060_v1_), d_1060_v1_, globalState)
                r1 = d_1080_cf13_
                d_1090_v17_: _dafny.Seq
                d_1090_v17_ = _dafny.SeqWithoutIsStrInference([d_1065_v5_])
                d_1091_v18_: _dafny.Seq
                d_1091_v18_ = _dafny.SeqWithoutIsStrInference([d_1060_v1_, len(d_1090_v17_), d_1080_cf13_, d_1060_v1_])
                d_1092_v19_: _dafny.Map
                d_1092_v19_ = _dafny.Map({d_1060_v1_: d_1071_v11_})
                d_1080_cf13_ = ((self).fm0(d_1060_v1_, p0, globalState)) * ((d_1091_v18_)[default__.safeIndex(len(((d_1092_v19_)[d_1080_cf13_] if (d_1080_cf13_) in (d_1092_v19_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uqbs")))), len(d_1091_v18_))])
                d_1079_cf14_ = not((d_1060_v1_) != (939))
            elif source17_.is_DC6:
                d_1093___mcc_h4_ = source17_.cf10
                d_1094_cf10_ = d_1093___mcc_h4_
                d_1095_v20_: C4
                nw145_ = C4()
                nw145_.ctor__(default__.safeModuloInt(459, d_1060_v1_), default__.fm36(globalState))
                d_1095_v20_ = nw145_
                d_1096_v21_: _dafny.Map
                d_1096_v21_ = _dafny.Map({d_1070_v10_: d_1068_v8_})
                d_1096_v21_ = ((_dafny.Map({d_1070_v10_: (d_1070_v10_)[default__.safeIndex(425, (d_1070_v10_).length(0))]})) | (d_1096_v21_)) | (d_1096_v21_)
                d_1060_v1_ = 999
                d_1097_v22_: _dafny.Map
                d_1097_v22_ = _dafny.Map({d_1068_v8_: (d_1059_v0_) not in (d_1065_v5_)})
                d_1097_v22_ = (d_1097_v22_).set(d_1068_v8_, not(d_1059_v0_))
            elif True:
                d_1098___mcc_h5_ = source17_.cf15
                d_1099_cf15_ = d_1098___mcc_h5_
                r1 = default__.safeModuloInt(d_1060_v1_, d_1060_v1_)
                (globalState).f9 = (d_1070_v10_)[default__.safeIndex(425, (d_1070_v10_).length(0))]
                (globalState).f13 = ((d_1063_v4_).issubset(d_1063_v4_)) == (not(not (d_1068_v8_) or (d_1068_v8_)))
                d_1068_v8_ = ((850 if d_1068_v8_ else d_1060_v1_)) <= (d_1060_v1_)
            d_1100_v23_: C2
            nw146_ = C2()
            nw146_.ctor__()
            d_1100_v23_ = nw146_
            d_1101_v24_: _dafny.MultiSet
            d_1101_v24_ = _dafny.MultiSet([d_1100_v23_])
            (globalState).f9 = not((d_1101_v24_).isdisjoint((_dafny.MultiSet([d_1100_v23_])).set(d_1100_v23_, default__.abs(len(d_1065_v5_)))))
            if (d_1070_v10_)[default__.safeIndex(425, (d_1070_v10_).length(0))]:
                d_1060_v1_ = d_1060_v1_
                d_1102_v25_: _dafny.MultiSet
                d_1102_v25_ = _dafny.MultiSet([D0_DC1(True, d_1072_v12_, d_1071_v11_)])
                d_1103_v26_: C0
                nw147_ = C0()
                nw147_.ctor__((d_1102_v25_).intersection(d_1102_v25_), d_1068_v8_)
                d_1103_v26_ = nw147_
                d_1104_v27_: _dafny.Map
                d_1105_v28_: _dafny.MultiSet
                d_1106_v29_: bool
                d_1107_v30_: _dafny.Array
                out18_: _dafny.Map
                out19_: _dafny.MultiSet
                out20_: bool
                out21_: _dafny.Array
                out18_, out19_, out20_, out21_ = (self).m2(d_1065_v5_, not(not(d_1059_v0_)), globalState)
                d_1104_v27_ = out18_
                d_1105_v28_ = out19_
                d_1106_v29_ = out20_
                d_1107_v30_ = out21_
                d_1108_v31_: D5
                d_1108_v31_ = D5_DC13(d_1070_v10_)
                d_1109_v32_: _dafny.MultiSet
                d_1109_v32_ = _dafny.MultiSet([d_1108_v31_])
                d_1059_v0_ = (d_1109_v32_).ispropersubset(d_1109_v32_)
                index167_ = default__.safeIndex(425, (d_1070_v10_).length(0))
                (d_1070_v10_)[index167_] = (d_1063_v4_).isdisjoint(d_1063_v4_)
            elif True:
                d_1110_v33_: bool
                d_1111_v34_: int
                d_1112_v35_: bool
                out22_: bool
                out23_: int
                out24_: bool
                out22_, out23_, out24_ = (d_1100_v23_).m6(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_1113_i3_ in range(default__.abs(-864))]), d_1071_v11_, d_1068_v8_, d_1060_v1_, globalState)
                d_1110_v33_ = out22_
                d_1111_v34_ = out23_
                d_1112_v35_ = out24_
                d_1114_v36_: C2
                nw148_ = C2()
                nw148_.ctor__()
                d_1114_v36_ = nw148_
                d_1115_v37_: _dafny.Map
                d_1115_v37_ = _dafny.Map({d_1111_v34_: (d_1070_v10_)[default__.safeIndex(425, (d_1070_v10_).length(0))]})
                d_1116_v38_: _dafny.Seq
                d_1116_v38_ = _dafny.SeqWithoutIsStrInference([-119, d_1111_v34_])
                d_1117_v39_: _dafny.Map
                d_1117_v39_ = _dafny.Map({d_1112_v35_: (d_1116_v38_)[default__.safeIndex(d_1111_v34_, len(d_1116_v38_))]})
                d_1118_v40_: _dafny.Array
                nw149_ = _dafny.Array(None, 25)
                nw149_[int(0)] = d_1060_v1_
                nw149_[int(1)] = d_1060_v1_
                nw149_[int(2)] = d_1111_v34_
                nw149_[int(3)] = d_1060_v1_
                nw149_[int(4)] = d_1060_v1_
                nw149_[int(5)] = (0) - (len(_dafny.SeqWithoutIsStrInference([d_1059_v0_, d_1110_v33_])))
                nw149_[int(6)] = len(_dafny.Set({_dafny.SeqWithoutIsStrInference([p0 for d_1119_i4_ in range(default__.abs(580))]), d_1071_v11_, _dafny.SeqWithoutIsStrInference([p0 for d_1120_i5_ in range(default__.abs(566))])}))
                nw149_[int(7)] = d_1111_v34_
                nw149_[int(8)] = d_1060_v1_
                nw149_[int(9)] = len(_dafny.Map({d_1068_v8_: p0}))
                nw149_[int(10)] = d_1111_v34_
                nw149_[int(11)] = d_1060_v1_
                nw149_[int(12)] = d_1060_v1_
                nw149_[int(13)] = len(d_1071_v11_)
                nw149_[int(14)] = 36
                nw149_[int(15)] = (d_1114_v36_).fm0(d_1060_v1_, p0, globalState)
                nw149_[int(16)] = (d_1116_v38_)[default__.safeIndex(len(d_1117_v39_), len(d_1116_v38_))]
                nw149_[int(17)] = d_1060_v1_
                nw149_[int(18)] = d_1111_v34_
                nw149_[int(19)] = d_1060_v1_
                nw149_[int(20)] = d_1060_v1_
                nw149_[int(21)] = d_1060_v1_
                nw149_[int(22)] = d_1111_v34_
                nw149_[int(23)] = d_1060_v1_
                nw149_[int(24)] = d_1060_v1_
                d_1118_v40_ = nw149_
                d_1121_v41_: _dafny.Map
                d_1121_v41_ = _dafny.Map({d_1118_v40_: (0) - (d_1060_v1_)})
                d_1122_v42_: D8
                d_1122_v42_ = D8_DC22(d_1121_v41_)
                d_1123_v43_: _dafny.Set
                d_1123_v43_ = _dafny.Set({d_1059_v0_, d_1112_v35_, d_1059_v0_})
                d_1115_v37_ = (d_1115_v37_).set((d_1116_v38_)[default__.safeIndex(len((d_1122_v42_).cf34), len(d_1116_v38_))], (d_1123_v43_).ispropersubset(d_1123_v43_))
                d_1071_v11_ = ((d_1061_v2_).fm1(_dafny.Map({553: d_1111_v34_}), globalState)).set(default__.safeIndex(233, len((d_1061_v2_).fm1(_dafny.Map({553: d_1111_v34_}), globalState))), (p0 if not(d_1068_v8_) else p0))
                d_1110_v33_ = (d_1070_v10_)[default__.safeIndex(425, (d_1070_v10_).length(0))]
        elif True:
            source18_ = D7_DC18(_dafny.Map({d_1060_v1_: True}))
            if source18_.is_DC19:
                d_1124___mcc_h6_ = source18_.cf27
                d_1125___mcc_h7_ = source18_.cf28
                d_1126___mcc_h8_ = source18_.cf29
                d_1127___mcc_h9_ = source18_.cf30
                d_1128_cf30_ = d_1127___mcc_h9_
                d_1129_cf29_ = d_1126___mcc_h8_
                d_1130_cf28_ = d_1125___mcc_h7_
                d_1131_cf27_ = d_1124___mcc_h6_
                d_1132_v44_: _dafny.Seq
                d_1132_v44_ = _dafny.SeqWithoutIsStrInference([True])
                d_1133_v45_: _dafny.Map
                d_1134_v46_: _dafny.MultiSet
                d_1135_v47_: bool
                d_1136_v48_: _dafny.Array
                out25_: _dafny.Map
                out26_: _dafny.MultiSet
                out27_: bool
                out28_: _dafny.Array
                out25_, out26_, out27_, out28_ = (self).m2((d_1132_v44_) + (d_1132_v44_), (d_1130_cf28_) < (d_1060_v1_), globalState)
                d_1133_v45_ = out25_
                d_1134_v46_ = out26_
                d_1135_v47_ = out27_
                d_1136_v48_ = out28_
                d_1060_v1_ = d_1130_cf28_
                d_1137_v49_: _dafny.Map
                d_1137_v49_ = _dafny.Map({True: (d_1135_v47_ if d_1129_cf29_ else d_1135_v47_)})
                d_1138_v50_: _dafny.Seq
                d_1138_v50_ = _dafny.SeqWithoutIsStrInference([d_1137_v49_])
                d_1139_v51_: _dafny.Seq
                d_1139_v51_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tvnpxdv"))
                d_1140_v52_: _dafny.Seq
                d_1140_v52_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "anlbx")), d_1139_v51_])
                d_1141_v53_: D0
                d_1141_v53_ = D0_DC1(d_1135_v47_, (d_1140_v52_).set(default__.safeIndex(516, len(d_1140_v52_)), d_1139_v51_), d_1139_v51_)
                d_1142_v54_: _dafny.MultiSet
                d_1142_v54_ = _dafny.MultiSet([d_1141_v53_, d_1141_v53_])
                d_1143_v55_: C0
                nw150_ = C0()
                nw150_.ctor__(d_1142_v54_, d_1059_v0_)
                d_1143_v55_ = nw150_
                d_1144_v56_: _dafny.Map
                d_1144_v56_ = _dafny.Map({d_1143_v55_: d_1129_cf29_})
                d_1145_v57_: _dafny.Map
                d_1145_v57_ = _dafny.Map({not(d_1129_cf29_): len((d_1144_v56_).set(d_1143_v55_, (d_1143_v55_).f19))})
                d_1137_v49_ = (d_1137_v49_) | ((d_1138_v50_)[default__.safeIndex(((d_1145_v57_)[d_1129_cf29_] if (d_1129_cf29_) in (d_1145_v57_) else d_1130_cf28_), len(d_1138_v50_))])
                d_1146_v58_: T0
                nw151_ = C2()
                nw151_.ctor__()
                d_1146_v58_ = nw151_
                d_1146_v58_ = d_1146_v58_
            elif source18_.is_DC20:
                d_1147___mcc_h10_ = source18_.cf31
                d_1148___mcc_h11_ = source18_.cf32
                d_1149_cf32_ = d_1148___mcc_h11_
                d_1150_cf31_ = d_1147___mcc_h10_
                d_1151_v59_: _dafny.Map
                d_1151_v59_ = _dafny.Map({d_1060_v1_: 207})
                d_1152_v60_: _dafny.Seq
                d_1152_v60_ = _dafny.SeqWithoutIsStrInference([d_1150_cf31_])
                d_1153_v61_: D0
                d_1153_v61_ = D0_DC1(d_1059_v0_, d_1152_v60_, d_1150_cf31_)
                d_1154_v62_: _dafny.Map
                d_1154_v62_ = _dafny.Map({default__.fm19(len(_dafny.Set({d_1149_cf32_, d_1059_v0_, d_1149_cf32_, True})), d_1059_v0_, globalState): D0_DC3(d_1153_v61_)})
                d_1155_v63_: D2
                d_1155_v63_ = D2_DC7(d_1059_v0_, d_1154_v62_, len(d_1150_cf31_), True)
                d_1156_v64_: _dafny.Map
                d_1156_v64_ = _dafny.Map({d_1060_v1_: _dafny.CodePoint('t')})
                d_1157_v65_: _dafny.MultiSet
                d_1157_v65_ = _dafny.MultiSet([_dafny.Map({d_1059_v0_: d_1149_cf32_})])
                d_1158_v67_: _dafny.Seq
                def iife82_():
                    coll36_ = _dafny.Map()
                    compr_36_: int
                    for compr_36_ in _dafny.IntegerRange(713, 736):
                        d_1159_v66_: int = compr_36_
                        if ((713) <= (d_1159_v66_)) and ((d_1159_v66_) < (736)):
                            coll36_[default__.safeModuloInt(d_1159_v66_, d_1060_v1_)] = d_1149_cf32_
                    return _dafny.Map(coll36_)
                d_1158_v67_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({(d_1157_v65_).cardinality: d_1149_cf32_}), iife82_()
                ])
                d_1160_v68_: _dafny.Set
                d_1160_v68_ = _dafny.Set({d_1059_v0_})
                d_1161_v69_: _dafny.Array
                nw152_ = _dafny.Array(None, 20)
                nw152_[int(0)] = d_1060_v1_
                nw152_[int(1)] = (d_1060_v1_) - (d_1060_v1_)
                nw152_[int(2)] = (len(_dafny.Set({d_1149_cf32_}))) - (len((d_1151_v59_).set(d_1060_v1_, d_1060_v1_)))
                nw152_[int(3)] = default__.safeDivisionInt(len(d_1150_cf31_), d_1060_v1_)
                nw152_[int(4)] = d_1060_v1_
                nw152_[int(5)] = d_1060_v1_
                nw152_[int(6)] = default__.safeModuloInt(d_1060_v1_, d_1060_v1_)
                nw152_[int(7)] = d_1060_v1_
                nw152_[int(8)] = (0) - (d_1060_v1_)
                nw152_[int(9)] = default__.fm22(d_1155_v63_, D0_DC3(d_1153_v61_), d_1149_cf32_, globalState)
                nw152_[int(10)] = (d_1060_v1_) - (d_1060_v1_)
                nw152_[int(11)] = d_1060_v1_
                nw152_[int(12)] = d_1060_v1_
                nw152_[int(13)] = d_1060_v1_
                nw152_[int(14)] = 183
                nw152_[int(15)] = (0) - (default__.safeModuloInt(d_1060_v1_, len(default__.fm37(not(d_1149_cf32_), d_1060_v1_, d_1060_v1_, ((d_1156_v64_)[d_1060_v1_] if (d_1060_v1_) in (d_1156_v64_) else p0), globalState))))
                nw152_[int(16)] = d_1060_v1_
                nw152_[int(17)] = (len(d_1158_v67_)) - (len(d_1150_cf31_))
                nw152_[int(18)] = len(d_1160_v68_)
                nw152_[int(19)] = (342) + (138)
                d_1161_v69_ = nw152_
                index168_ = default__.safeIndex(424, (d_1161_v69_).length(0))
                (d_1161_v69_)[index168_] = len(d_1151_v59_)
                index169_ = default__.safeIndex(424, (d_1161_v69_).length(0))
                (d_1161_v69_)[index169_] = 674
                d_1162_v70_: _dafny.Array
                nw153_ = _dafny.Array(False, 18)
                d_1162_v70_ = nw153_
                d_1162_v70_ = d_1162_v70_
                d_1060_v1_ = (0) - (len(_dafny.SeqWithoutIsStrInference([d_1149_cf32_, default__.fm6(globalState), d_1149_cf32_, not(d_1059_v0_), d_1059_v0_])))
                r1 = d_1060_v1_
            elif source18_.is_DC18:
                d_1163___mcc_h12_ = source18_.cf26
                d_1164_cf26_ = d_1163___mcc_h12_
                d_1165_v71_: _dafny.Array
                def lambda44_(d_1166_v4_):
                    def lambda45_(d_1167_i6_):
                        return (d_1166_v4_ if True else d_1166_v4_)

                    return lambda45_

                init21_ = lambda44_(d_1063_v4_)
                nw154_ = _dafny.Array(None, 14)
                for i0_21_ in range(nw154_.length(0)):
                    nw154_[i0_21_] = init21_(i0_21_)
                d_1165_v71_ = nw154_
                index170_ = default__.safeIndex(231, (d_1165_v71_).length(0))
                (d_1165_v71_)[index170_] = (d_1063_v4_ if d_1059_v0_ else d_1063_v4_)
                index171_ = default__.safeIndex(231, (d_1165_v71_).length(0))
                (d_1165_v71_)[index171_] = d_1063_v4_
                d_1168_v72_: _dafny.Array
                def lambda46_(d_1169_v0_):
                    def lambda47_(d_1170_i7_):
                        return default__.safeDivisionInt(d_1170_i7_, len(_dafny.Map({d_1169_v0_: d_1169_v0_})))

                    return lambda47_

                init22_ = lambda46_(d_1059_v0_)
                nw155_ = _dafny.Array(None, 1)
                for i0_22_ in range(nw155_.length(0)):
                    nw155_[i0_22_] = init22_(i0_22_)
                d_1168_v72_ = nw155_
                d_1171_v73_: _dafny.Seq
                d_1171_v73_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "taqaliqve"))
                index172_ = default__.safeIndex(882, (d_1168_v72_).length(0))
                (d_1168_v72_)[index172_] = (0) - (default__.safeModuloInt(len(d_1171_v73_), d_1060_v1_))
                index173_ = default__.safeIndex(882, (d_1168_v72_).length(0))
                (d_1168_v72_)[index173_] = d_1060_v1_
                d_1172_v74_: _dafny.Array
                nw156_ = _dafny.Array(_dafny.Map({}), 20)
                d_1172_v74_ = nw156_
                d_1173_v75_: _dafny.Map
                d_1173_v75_ = _dafny.Map({d_1168_v72_: len(d_1171_v73_)})
                d_1174_v76_: D8
                d_1174_v76_ = D8_DC22((d_1173_v75_).set(d_1168_v72_, (d_1168_v72_)[default__.safeIndex(882, (d_1168_v72_).length(0))]))
                d_1175_v77_: _dafny.Array
                nw157_ = _dafny.Array(None, 3)
                nw157_[int(0)] = d_1059_v0_
                nw157_[int(1)] = d_1059_v0_
                nw157_[int(2)] = d_1059_v0_
                d_1175_v77_ = nw157_
                index174_ = default__.safeIndex(695, (d_1175_v77_).length(0))
                (d_1175_v77_)[index174_] = d_1059_v0_
                d_1176_v78_: _dafny.Set
                d_1176_v78_ = _dafny.Set({d_1059_v0_, not(d_1059_v0_), d_1059_v0_, d_1059_v0_, d_1059_v0_})
                pat_let_tv23_ = d_1173_v75_
                index175_ = default__.safeIndex(882, (d_1168_v72_).length(0))
                index176_ = default__.safeIndex(695, (d_1175_v77_).length(0))
                def iife83_(_pat_let23_0):
                    def iife84_(d_1177_dt__update__tmp_h0_):
                        def iife85_(_pat_let24_0):
                            def iife86_(d_1178_dt__update_hcf34_h0_):
                                return D8_DC22(d_1178_dt__update_hcf34_h0_)
                            return iife86_(_pat_let24_0)
                        return iife85_(pat_let_tv23_)
                    return iife84_(_pat_let23_0)
                rhs180_ = d_1172_v74_
                rhs181_ = not(d_1059_v0_)
                rhs182_ = ((self).fm0(d_1060_v1_, p0, globalState)) - (d_1060_v1_)
                rhs183_ = iife83_(d_1174_v76_)
                rhs184_ = (default__.fm33(d_1059_v0_, d_1059_v0_, d_1059_v0_, globalState)).isdisjoint((d_1176_v78_).intersection(d_1176_v78_))
                lhs124_ = globalState
                lhs125_ = d_1168_v72_
                lhs126_ = default__.safeIndex(882, (d_1168_v72_).length(0))
                lhs127_ = d_1175_v77_
                lhs128_ = default__.safeIndex(695, (d_1175_v77_).length(0))
                d_1172_v74_ = rhs180_
                lhs124_.f3 = rhs181_
                lhs125_[lhs126_] = rhs182_
                d_1174_v76_ = rhs183_
                lhs127_[lhs128_] = rhs184_
                d_1179_v79_: D5
                d_1179_v79_ = D5_DC13(d_1175_v77_)
                d_1175_v77_ = (d_1179_v79_).cf19
            elif True:
                d_1180___mcc_h13_ = source18_.cf33
                d_1181_cf33_ = d_1180___mcc_h13_
                d_1182_v80_: _dafny.Seq
                d_1182_v80_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pdxrtkj"))
                r1 = default__.safeDivisionInt(len(d_1182_v80_), default__.safeModuloInt(d_1060_v1_, d_1060_v1_))
                d_1060_v1_ = (0) - (d_1060_v1_)
                d_1183_v81_: _dafny.Array
                nw158_ = _dafny.Array(_dafny.Array(None, 0), 8)
                d_1183_v81_ = nw158_
                d_1184_v82_: _dafny.Array
                def lambda48_(d_1185_v0_):
                    def lambda49_(d_1186_i8_):
                        return d_1185_v0_

                    return lambda49_

                init23_ = lambda48_(d_1059_v0_)
                nw159_ = _dafny.Array(None, 13)
                for i0_23_ in range(nw159_.length(0)):
                    nw159_[i0_23_] = init23_(i0_23_)
                d_1184_v82_ = nw159_
                index177_ = default__.safeIndex(858, (d_1183_v81_).length(0))
                (d_1183_v81_)[index177_] = d_1184_v82_
                index178_ = default__.safeIndex(858, (d_1183_v81_).length(0))
                (d_1183_v81_)[index178_] = d_1184_v82_
                d_1187_v83_: D4
                d_1187_v83_ = D4_DC12(d_1060_v1_)
                d_1188_v84_: _dafny.Map
                d_1188_v84_ = _dafny.Map({d_1187_v83_: ((0) - ((d_1061_v2_).fm0(d_1060_v1_, p0, globalState)) if d_1059_v0_ else d_1060_v1_)})
                r1 = ((d_1188_v84_)[default__.fm35(_dafny.SeqWithoutIsStrInference([True]), d_1059_v0_, _dafny.SeqWithoutIsStrInference([p0 for d_1189_i9_ in range(default__.abs(82))]), globalState)] if (default__.fm35(_dafny.SeqWithoutIsStrInference([True]), d_1059_v0_, _dafny.SeqWithoutIsStrInference([p0 for d_1190_i9_ in range(default__.abs(82))]), globalState)) in (d_1188_v84_) else default__.safeModuloInt(279, d_1060_v1_))
            d_1191_v85_: _dafny.Array
            nw160_ = _dafny.Array(int(0), 8)
            d_1191_v85_ = nw160_
            d_1191_v85_ = d_1191_v85_
            d_1192_v86_: _dafny.Map
            d_1192_v86_ = _dafny.Map({d_1060_v1_: (False) == (d_1059_v0_)})
            d_1193_v87_: _dafny.Seq
            d_1193_v87_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lgyduwq"))
            rhs185_ = d_1192_v86_
            rhs186_ = len(d_1193_v87_)
            rhs187_ = d_1191_v85_
            d_1192_v86_ = rhs185_
            r1 = rhs186_
            d_1191_v85_ = rhs187_
            d_1194_v88_: _dafny.Map
            d_1194_v88_ = _dafny.Map({(d_1060_v1_) <= (d_1060_v1_): d_1059_v0_})
            d_1194_v88_ = (d_1194_v88_).set(d_1059_v0_, False)
            if d_1059_v0_:
                d_1195_v89_: _dafny.Seq
                d_1195_v89_ = _dafny.SeqWithoutIsStrInference([d_1059_v0_, d_1059_v0_, d_1059_v0_, d_1059_v0_, True])
                r1 = len(_dafny.SeqWithoutIsStrInference([(0) - (len(d_1195_v89_))]))
                d_1196_v90_: D1
                d_1196_v90_ = D1_DC5(d_1059_v0_, d_1193_v87_, d_1195_v89_)
                d_1197_v91_: _dafny.Set
                d_1197_v91_ = _dafny.Set({_dafny.Map({760: d_1196_v90_})})
                d_1198_v92_: _dafny.Map
                d_1198_v92_ = _dafny.Map({d_1060_v1_: d_1196_v90_})
                d_1197_v91_ = (d_1197_v91_) | (_dafny.Set({(d_1198_v92_).set(d_1060_v1_, d_1196_v90_)}))
                d_1199_v93_: C3
                nw161_ = C3()
                nw161_.ctor__((_dafny.Set({len(d_1193_v87_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kenqbixrf")))})).ispropersubset(d_1063_v4_), d_1060_v1_)
                d_1199_v93_ = nw161_
                d_1195_v89_ = ((d_1195_v89_) + (d_1195_v89_)) + (d_1195_v89_)
                d_1200_v94_: _dafny.Seq
                d_1200_v94_ = _dafny.SeqWithoutIsStrInference([d_1060_v1_])
                d_1201_v95_: _dafny.Array
                nw162_ = _dafny.Array(D0.default()(), 29)
                d_1201_v95_ = nw162_
                d_1202_v96_: _dafny.Set
                d_1202_v96_ = _dafny.Set({(D5_DC14((d_1199_v93_).fm13((d_1199_v93_).f17, d_1200_v94_, (d_1199_v93_).f17, d_1059_v0_, globalState), d_1059_v0_, d_1201_v95_)).cf20, default__.fm19((d_1199_v93_).f17, (d_1199_v93_).f16, globalState)})
                d_1203_v97_: _dafny.Seq
                d_1203_v97_ = _dafny.SeqWithoutIsStrInference([d_1202_v96_, d_1202_v96_, d_1202_v96_, d_1202_v96_])
                d_1202_v96_ = ((d_1202_v96_) | (d_1202_v96_) if ((d_1199_v93_).f17) != (-723) else ((d_1203_v97_)[default__.safeIndex(d_1060_v1_, len(d_1203_v97_))]) - (_dafny.Set({(d_1199_v93_).f16})))
            elif True:
                (globalState).f13 = d_1059_v0_
                r1 = (0) - ((0) - (d_1060_v1_))
                d_1204_v98_: _dafny.Seq
                d_1204_v98_ = _dafny.SeqWithoutIsStrInference([d_1059_v0_, d_1059_v0_])
                r0 = ((d_1204_v98_) + (d_1204_v98_)) == ((d_1204_v98_) + (d_1204_v98_))
                r1 = d_1060_v1_
                rhs188_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jmukv"))
                rhs189_ = ((d_1193_v87_).set(default__.safeIndex(d_1060_v1_, len(d_1193_v87_)), _dafny.CodePoint('c'))).set(default__.safeIndex(((0) - (d_1060_v1_)) * (d_1060_v1_), len((d_1193_v87_).set(default__.safeIndex(d_1060_v1_, len(d_1193_v87_)), _dafny.CodePoint('c')))), p0)
                d_1193_v87_ = rhs188_
                d_1193_v87_ = rhs189_
        d_1205_v99_: D0
        d_1205_v99_ = D0_DC0(d_1060_v1_)
        d_1206_v100_: D0
        d_1206_v100_ = D0_DC3(d_1205_v99_)
        d_1207_v101_: _dafny.Map
        d_1207_v101_ = _dafny.Map({d_1059_v0_: d_1206_v100_})
        d_1208_v102_: D7
        d_1208_v102_ = D7_DC20(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "khrd")), d_1059_v0_)
        d_1209_v103_: D2
        d_1209_v103_ = D2_DC7(d_1059_v0_, d_1207_v101_, d_1060_v1_, (d_1208_v102_).cf32)
        d_1210_v104_: _dafny.Array
        nw163_ = _dafny.Array(None, 19)
        nw163_[int(0)] = d_1059_v0_
        nw163_[int(1)] = not(d_1059_v0_)
        nw163_[int(2)] = d_1059_v0_
        nw163_[int(3)] = d_1059_v0_
        nw163_[int(4)] = False
        nw163_[int(5)] = d_1059_v0_
        nw163_[int(6)] = d_1059_v0_
        nw163_[int(7)] = d_1059_v0_
        nw163_[int(8)] = not(d_1059_v0_)
        nw163_[int(9)] = False
        nw163_[int(10)] = d_1059_v0_
        nw163_[int(11)] = d_1059_v0_
        nw163_[int(12)] = (self).fm3(True, d_1059_v0_, 248, globalState)
        nw163_[int(13)] = d_1059_v0_
        nw163_[int(14)] = d_1059_v0_
        nw163_[int(15)] = (d_1209_v103_).cf11
        nw163_[int(16)] = d_1059_v0_
        nw163_[int(17)] = d_1059_v0_
        nw163_[int(18)] = d_1059_v0_
        d_1210_v104_ = nw163_
        d_1211_v105_: _dafny.Array
        nw164_ = _dafny.Array(None, 9)
        nw164_[int(0)] = d_1210_v104_
        nw164_[int(1)] = d_1210_v104_
        nw164_[int(2)] = d_1210_v104_
        nw164_[int(3)] = d_1210_v104_
        nw164_[int(4)] = d_1210_v104_
        nw164_[int(5)] = d_1210_v104_
        nw164_[int(6)] = d_1210_v104_
        nw164_[int(7)] = d_1210_v104_
        nw164_[int(8)] = d_1210_v104_
        d_1211_v105_ = nw164_
        d_1212_v106_: _dafny.Seq
        d_1212_v106_ = _dafny.SeqWithoutIsStrInference([d_1210_v104_, d_1210_v104_])
        index179_ = default__.safeIndex(498, (d_1211_v105_).length(0))
        (d_1211_v105_)[index179_] = (d_1212_v106_)[default__.safeIndex(-808, len(d_1212_v106_))]
        index180_ = default__.safeIndex(498, (d_1211_v105_).length(0))
        (d_1211_v105_)[index180_] = d_1210_v104_
        index181_ = default__.safeIndex(879, (d_1210_v104_).length(0))
        (d_1210_v104_)[index181_] = d_1059_v0_
        index182_ = default__.safeIndex(879, (d_1210_v104_).length(0))
        (d_1210_v104_)[index182_] = not (d_1059_v0_) or ((not(d_1059_v0_) if d_1059_v0_ else d_1059_v0_))
        d_1213_v107_: C3
        nw165_ = C3()
        nw165_.ctor__((d_1210_v104_)[default__.safeIndex(879, (d_1210_v104_).length(0))], d_1060_v1_)
        d_1213_v107_ = nw165_
        d_1214_i10_: int
        d_1214_i10_ = 0
        with _dafny.label("8"):
            while (d_1213_v107_).f16:
                with _dafny.c_label("8"):
                    if (d_1214_i10_) >= (100):
                        raise _dafny.Break("8")
                    d_1214_i10_ = (d_1214_i10_) + (1)
                    d_1215_v108_: _dafny.Seq
                    d_1215_v108_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pjglis"))
                    d_1215_v108_ = (d_1215_v108_ if d_1059_v0_ else d_1215_v108_)
                    d_1216_v109_: _dafny.Seq
                    d_1216_v109_ = _dafny.SeqWithoutIsStrInference([d_1215_v108_])
                    d_1217_v110_: _dafny.Seq
                    d_1217_v110_ = _dafny.SeqWithoutIsStrInference([D0_DC1((d_1210_v104_)[default__.safeIndex(879, (d_1210_v104_).length(0))], d_1216_v109_, d_1215_v108_)])
                    d_1218_v111_: D0
                    d_1218_v111_ = D0_DC1((d_1210_v104_)[default__.safeIndex(879, (d_1210_v104_).length(0))], d_1216_v109_, d_1215_v108_)
                    d_1219_v112_: _dafny.MultiSet
                    d_1219_v112_ = _dafny.MultiSet([d_1218_v111_, d_1218_v111_, d_1218_v111_])
                    pat_let_tv24_ = d_1219_v112_
                    def iife87_(_pat_let25_0):
                        def iife88_(d_1220_dt__update__tmp_h1_):
                            def iife89_(_pat_let26_0):
                                def iife90_(d_1221_dt__update_hcf15_h0_):
                                    return D2_DC8(d_1221_dt__update_hcf15_h0_)
                                return iife90_(_pat_let26_0)
                            return iife89_(D2_DC6(pat_let_tv24_))
                        return iife88_(_pat_let25_0)
                    source19_ = iife87_(D2_DC8(D2_DC6(_dafny.MultiSet(d_1217_v110_))))
                    if source19_.is_DC7:
                        d_1222___mcc_h14_ = source19_.cf11
                        d_1223___mcc_h15_ = source19_.cf12
                        d_1224___mcc_h16_ = source19_.cf13
                        d_1225___mcc_h17_ = source19_.cf14
                        d_1226_cf14_ = d_1225___mcc_h17_
                        d_1227_cf13_ = d_1224___mcc_h16_
                        d_1228_cf12_ = d_1223___mcc_h15_
                        d_1229_cf11_ = d_1222___mcc_h14_
                        index183_ = default__.safeIndex(879, (d_1210_v104_).length(0))
                        (d_1210_v104_)[index183_] = True
                        index184_ = default__.safeIndex(879, (d_1210_v104_).length(0))
                        (d_1210_v104_)[index184_] = d_1226_cf14_
                        d_1230_v113_: _dafny.Map
                        d_1230_v113_ = _dafny.Map({p0: (d_1213_v107_).f17})
                        d_1231_v114_: D9
                        d_1231_v114_ = D9_DC25(d_1230_v113_)
                        d_1232_v115_: C4
                        nw166_ = C4()
                        nw166_.ctor__(d_1060_v1_, (d_1231_v114_).cf38)
                        d_1232_v115_ = nw166_
                        d_1233_v116_: _dafny.Seq
                        d_1233_v116_ = _dafny.SeqWithoutIsStrInference([d_1226_cf14_, True])
                        index185_ = default__.safeIndex(879, (d_1210_v104_).length(0))
                        (d_1210_v104_)[index185_] = (d_1233_v116_)[default__.safeIndex((d_1213_v107_).f17, len(d_1233_v116_))]
                    elif source19_.is_DC6:
                        d_1234___mcc_h18_ = source19_.cf10
                        d_1235_cf10_ = d_1234___mcc_h18_
                        d_1236_v117_: D9
                        d_1236_v117_ = D9_DC27((d_1210_v104_)[default__.safeIndex(879, (d_1210_v104_).length(0))], (d_1213_v107_).f17, (d_1210_v104_)[default__.safeIndex(879, (d_1210_v104_).length(0))])
                        d_1237_v118_: _dafny.MultiSet
                        d_1237_v118_ = _dafny.MultiSet([(d_1213_v107_).f16])
                        rhs190_ = ((d_1237_v118_ if (d_1213_v107_).f16 else d_1237_v118_)).cardinality
                        rhs191_ = d_1236_v117_
                        d_1060_v1_ = rhs190_
                        d_1236_v117_ = rhs191_
                        d_1238_v119_: _dafny.MultiSet
                        d_1238_v119_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([len(default__.fm30(globalState))]))])
                        index186_ = default__.safeIndex(879, (d_1210_v104_).length(0))
                        (d_1210_v104_)[index186_] = (d_1238_v119_).issubset(_dafny.MultiSet([(d_1213_v107_).f17]))
                        d_1215_v108_ = (d_1215_v108_) + ((d_1215_v108_) + (d_1215_v108_))
                        d_1239_v120_: _dafny.Seq
                        d_1239_v120_ = _dafny.SeqWithoutIsStrInference([d_1213_v107_])
                        index187_ = default__.safeIndex(879, (d_1210_v104_).length(0))
                        rhs192_ = (d_1239_v120_) < (d_1239_v120_)
                        rhs193_ = (d_1213_v107_).f17
                        lhs129_ = d_1210_v104_
                        lhs130_ = default__.safeIndex(879, (d_1210_v104_).length(0))
                        lhs129_[lhs130_] = rhs192_
                        r1 = rhs193_
                    elif True:
                        d_1240___mcc_h19_ = source19_.cf15
                        d_1241_cf15_ = d_1240___mcc_h19_
                        d_1242_v121_: D5
                        d_1242_v121_ = D5_DC13((d_1211_v105_)[default__.safeIndex(498, (d_1211_v105_).length(0))])
                        d_1243_v122_: D5
                        d_1243_v122_ = D5_DC13((d_1242_v121_).cf19)
                        index188_ = default__.safeIndex(498, (d_1211_v105_).length(0))
                        (d_1211_v105_)[index188_] = (d_1242_v121_).cf19
                        (globalState).f5 = ((d_1213_v107_).f17) <= ((d_1213_v107_).f17)
                        d_1244_v123_: _dafny.Map
                        d_1244_v123_ = _dafny.Map({(d_1213_v107_).f17: (d_1213_v107_).f17})
                        r1 = (default__.safeDivisionInt((0) - (d_1060_v1_), len((self).fm1(d_1244_v123_, globalState)))) * ((0) - (len(d_1215_v108_)))
                        d_1245_v124_: _dafny.Array
                        d_1245_v124_ = d_1211_v105_
                        d_1211_v105_ = (d_1245_v124_)
                    d_1246_v125_: _dafny.Set
                    d_1246_v125_ = _dafny.Set({(d_1210_v104_)[default__.safeIndex(879, (d_1210_v104_).length(0))], default__.fm19((d_1213_v107_).f17, d_1059_v0_, globalState)})
                    d_1246_v125_ = d_1246_v125_
                    d_1247_v126_: _dafny.Array
                    nw167_ = _dafny.Array(int(0), 2)
                    d_1247_v126_ = nw167_
                    d_1248_v127_: _dafny.Map
                    d_1248_v127_ = _dafny.Map({d_1247_v126_: (d_1213_v107_).f17})
                    d_1249_v128_: D8
                    d_1249_v128_ = D8_DC22((d_1248_v127_).set(d_1247_v126_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_1250_i11_ in range(default__.abs(129))]))))
                    d_1249_v128_ = d_1249_v128_
                    pass
            pass
        d_1251_v129_: _dafny.Array
        nw168_ = _dafny.Array(None, 19)
        d_1251_v129_ = nw168_
        d_1252_v130_: T0
        nw169_ = C2()
        nw169_.ctor__()
        d_1252_v130_ = nw169_
        index189_ = default__.safeIndex(294, (d_1251_v129_).length(0))
        (d_1251_v129_)[index189_] = d_1252_v130_
        d_1253_v131_: _dafny.Array
        nw170_ = _dafny.Array(_dafny.Set({}), 13)
        d_1253_v131_ = nw170_
        index190_ = default__.safeIndex(360, (d_1253_v131_).length(0))
        (d_1253_v131_)[index190_] = d_1063_v4_
        index191_ = default__.safeIndex(294, (d_1251_v129_).length(0))
        index192_ = default__.safeIndex(360, (d_1253_v131_).length(0))
        rhs194_ = d_1252_v130_
        rhs195_ = d_1063_v4_
        lhs131_ = d_1251_v129_
        lhs132_ = default__.safeIndex(294, (d_1251_v129_).length(0))
        lhs133_ = d_1253_v131_
        lhs134_ = default__.safeIndex(360, (d_1253_v131_).length(0))
        lhs131_[lhs132_] = rhs194_
        lhs133_[lhs134_] = rhs195_
        d_1254_v132_: _dafny.Seq
        d_1254_v132_ = _dafny.SeqWithoutIsStrInference([not(default__.fm6(globalState)), (d_1213_v107_).f16, ((d_1210_v104_)[default__.safeIndex(879, (d_1210_v104_).length(0))] if (d_1213_v107_).f16 else d_1059_v0_)])
        r0 = (d_1254_v132_)[default__.safeIndex((d_1213_v107_).f17, len(d_1254_v132_))]
        d_1255_v133_: _dafny.Map
        d_1255_v133_ = _dafny.Map({True: d_1254_v132_})
        r1 = len((_dafny.Map({(d_1210_v104_)[default__.safeIndex(879, (d_1210_v104_).length(0))]: d_1254_v132_}) if (d_1210_v104_)[default__.safeIndex(879, (d_1210_v104_).length(0))] else d_1255_v133_))
        return r0, r1

    def m2(self, p0, p1, globalState):
        r0: _dafny.Map = _dafny.Map({})
        r1: _dafny.MultiSet = _dafny.MultiSet({})
        r2: bool = False
        r3: _dafny.Array = _dafny.Array(None, 0)
        d_1256_v0_: _dafny.Seq
        d_1256_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ypgvmngw"))
        d_1257_v1_: _dafny.MultiSet
        d_1257_v1_ = _dafny.MultiSet([len(d_1256_v0_)])
        d_1258_v2_: _dafny.Map
        d_1258_v2_ = _dafny.Map({p1: d_1257_v1_})
        d_1259_v3_: _dafny.Map
        d_1259_v3_ = _dafny.Map({d_1258_v2_: d_1257_v1_})
        d_1260_v4_: int
        d_1260_v4_ = -81
        d_1261_v5_: _dafny.Seq
        d_1261_v5_ = _dafny.SeqWithoutIsStrInference([d_1260_v4_])
        d_1259_v3_ = (d_1259_v3_).set(d_1258_v2_, _dafny.MultiSet((d_1261_v5_ if p1 else d_1261_v5_)))
        d_1260_v4_ = default__.fm11((p1) and (True), globalState)
        d_1262_v6_: _dafny.Array
        nw171_ = _dafny.Array(_dafny.Map({}), 4)
        d_1262_v6_ = nw171_
        d_1263_v7_: _dafny.Map
        d_1263_v7_ = _dafny.Map({p1: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_1264_i0_ in range(default__.abs(732))])})
        index193_ = default__.safeIndex(73, (d_1262_v6_).length(0))
        (d_1262_v6_)[index193_] = (d_1263_v7_) | (d_1263_v7_)
        d_1265_v8_: str
        d_1265_v8_ = _dafny.CodePoint('l')
        d_1266_v9_: _dafny.Map
        d_1266_v9_ = _dafny.Map({d_1261_v5_: d_1260_v4_})
        d_1267_v10_: _dafny.Map
        d_1267_v10_ = _dafny.Map({d_1260_v4_: (d_1257_v1_).cardinality})
        d_1268_v11_: _dafny.Map
        d_1268_v11_ = _dafny.Map({len(d_1267_v10_): d_1260_v4_})
        index194_ = default__.safeIndex(73, (d_1262_v6_).length(0))
        rhs196_ = (0) - (d_1260_v4_)
        rhs197_ = (d_1263_v7_) | (_dafny.Map({p1: d_1256_v0_}))
        rhs198_ = d_1265_v8_
        rhs199_ = d_1266_v9_
        rhs200_ = (self).fm1(d_1268_v11_, globalState)
        lhs135_ = d_1262_v6_
        lhs136_ = default__.safeIndex(73, (d_1262_v6_).length(0))
        d_1260_v4_ = rhs196_
        lhs135_[lhs136_] = rhs197_
        d_1265_v8_ = rhs198_
        d_1266_v9_ = rhs199_
        d_1256_v0_ = rhs200_
        d_1265_v8_ = d_1265_v8_
        d_1269_v12_: _dafny.Set
        d_1269_v12_ = _dafny.Set({d_1260_v4_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cdtw")))})
        d_1269_v12_ = d_1269_v12_
        if p1:
            d_1256_v0_ = (d_1256_v0_) + (d_1256_v0_)
            d_1270_v13_: _dafny.Map
            d_1270_v13_ = _dafny.Map({d_1260_v4_: p1})
            d_1270_v13_ = ((d_1270_v13_) | (d_1270_v13_)) | (d_1270_v13_)
            d_1256_v0_ = d_1256_v0_
            d_1271_v14_: _dafny.Seq
            d_1271_v14_ = _dafny.SeqWithoutIsStrInference([d_1261_v5_, d_1261_v5_])
            d_1272_v15_: _dafny.Seq
            d_1272_v15_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1261_v5_]), (d_1271_v14_) + (d_1271_v14_), default__.fm38(globalState)])
            d_1273_v16_: D0
            d_1273_v16_ = D0_DC0(len(p0))
            d_1274_v17_: D0
            d_1274_v17_ = D0_DC3(d_1273_v16_)
            d_1275_v18_: _dafny.Map
            d_1275_v18_ = _dafny.Map({p1: d_1274_v17_})
            d_1271_v14_ = (d_1272_v15_)[default__.safeIndex(default__.fm22(D2_DC7(p1, d_1275_v18_, d_1260_v4_, p1), d_1274_v17_, p1, globalState), len(d_1272_v15_))]
            d_1276_v19_: _dafny.Map
            d_1276_v19_ = _dafny.Map({False: p1})
            d_1277_v20_: _dafny.Array
            nw172_ = _dafny.Array(None, 17)
            nw172_[int(0)] = p1
            nw172_[int(1)] = ((d_1276_v19_)[p1] if (p1) in (d_1276_v19_) else p1)
            nw172_[int(2)] = p1
            nw172_[int(3)] = p1
            nw172_[int(4)] = True
            nw172_[int(5)] = True
            nw172_[int(6)] = not(p1)
            nw172_[int(7)] = True
            nw172_[int(8)] = (d_1269_v12_).ispropersubset(d_1269_v12_)
            nw172_[int(9)] = p1
            nw172_[int(10)] = p1
            nw172_[int(11)] = not((p1) == (p1))
            nw172_[int(12)] = p1
            nw172_[int(13)] = not (default__.fm19(d_1260_v4_, p1, globalState)) or (p1)
            nw172_[int(14)] = p1
            nw172_[int(15)] = ((0) - (d_1260_v4_)) < (d_1260_v4_)
            nw172_[int(16)] = not(p1)
            d_1277_v20_ = nw172_
            index195_ = default__.safeIndex(288, (d_1277_v20_).length(0))
            (d_1277_v20_)[index195_] = False
            index196_ = default__.safeIndex(288, (d_1277_v20_).length(0))
            (d_1277_v20_)[index196_] = p1
        elif True:
            d_1278_v21_: _dafny.Map
            d_1278_v21_ = _dafny.Map({d_1265_v8_: d_1260_v4_})
            d_1279_v22_: C4
            nw173_ = C4()
            nw173_.ctor__((d_1260_v4_) * (d_1260_v4_), d_1278_v21_)
            d_1279_v22_ = nw173_
            d_1280_v23_: _dafny.Map
            d_1280_v23_ = _dafny.Map({(0) - ((0) - ((d_1279_v22_).f14)): p0})
            (globalState).f5 = (p0) != (((d_1280_v23_)[(d_1279_v22_).f14] if ((d_1279_v22_).f14) in (d_1280_v23_) else _dafny.SeqWithoutIsStrInference([p1])))
            (globalState).f13 = p1
            d_1281_v24_: D4
            d_1281_v24_ = D4_DC11()
            source20_ = d_1281_v24_
            if source20_.is_DC11:
                d_1282_v25_: D9
                d_1282_v25_ = D9_DC26(p1, d_1265_v8_, (p0)[default__.safeIndex((d_1279_v22_).f14, len(p0))])
                d_1283_v26_: _dafny.MultiSet
                d_1283_v26_ = _dafny.MultiSet([p1, (d_1282_v25_).cf41])
                d_1283_v26_ = (d_1283_v26_) | ((d_1283_v26_).intersection(default__.fm39(p1, globalState)))
                d_1284_v27_: _dafny.Set
                d_1284_v27_ = _dafny.Set({(d_1256_v0_) + (d_1256_v0_)})
                d_1284_v27_ = d_1284_v27_
                d_1285_v28_: _dafny.Array
                nw174_ = _dafny.Array(int(0), 10)
                d_1285_v28_ = nw174_
                d_1286_v29_: _dafny.Seq
                d_1286_v29_ = _dafny.SeqWithoutIsStrInference([d_1285_v28_, d_1285_v28_])
                d_1287_v30_: _dafny.Set
                d_1287_v30_ = _dafny.Set({p1, p1})
                d_1288_v31_: _dafny.Map
                d_1288_v31_ = _dafny.Map({p1: d_1287_v30_})
                d_1289_v32_: _dafny.Map
                d_1289_v32_ = _dafny.Map({p1: p1})
                d_1290_v33_: D8
                d_1290_v33_ = D8_DC23(p1, p1, len(d_1289_v32_))
                d_1291_v34_: _dafny.Map
                d_1291_v34_ = _dafny.Map({p1: len(d_1287_v30_)})
                d_1292_v35_: _dafny.Array
                nw175_ = _dafny.Array(None, 19)
                nw175_[int(0)] = p1
                nw175_[int(1)] = not(((0) - ((0) - ((d_1279_v22_).f14))) <= (len(d_1286_v29_)))
                nw175_[int(2)] = p1
                nw175_[int(3)] = not(p1)
                nw175_[int(4)] = ((d_1261_v5_).set(default__.safeIndex((d_1279_v22_).f14, len(d_1261_v5_)), len(d_1288_v31_))) < (_dafny.SeqWithoutIsStrInference([d_1260_v4_, (d_1279_v22_).f14]))
                nw175_[int(5)] = (d_1290_v33_).cf36
                nw175_[int(6)] = (p1) in (d_1291_v34_)
                nw175_[int(7)] = (d_1260_v4_) <= (d_1260_v4_)
                nw175_[int(8)] = p1
                nw175_[int(9)] = p1
                nw175_[int(10)] = not (p1) or (True)
                nw175_[int(11)] = p1
                nw175_[int(12)] = p1
                nw175_[int(13)] = not(p1)
                nw175_[int(14)] = not(p1)
                nw175_[int(15)] = not(p1)
                nw175_[int(16)] = (p1) or (p1)
                nw175_[int(17)] = False
                nw175_[int(18)] = p1
                d_1292_v35_ = nw175_
                index197_ = default__.safeIndex(434, (d_1292_v35_).length(0))
                (d_1292_v35_)[index197_] = p1
                index198_ = default__.safeIndex(434, (d_1292_v35_).length(0))
                (d_1292_v35_)[index198_] = default__.fm19((d_1279_v22_).f14, p1, globalState)
                d_1293_v36_: _dafny.Map
                d_1293_v36_ = _dafny.Map({(_dafny.MultiSet(d_1261_v5_)).cardinality: d_1286_v29_})
                rhs201_ = (d_1286_v29_) == (((d_1293_v36_)[(d_1279_v22_).f14] if ((d_1279_v22_).f14) in (d_1293_v36_) else _dafny.SeqWithoutIsStrInference([d_1285_v28_])))
                rhs202_ = (d_1256_v0_) + (d_1256_v0_)
                lhs137_ = globalState
                lhs137_.f3 = rhs201_
                d_1256_v0_ = rhs202_
            elif source20_.is_DC12:
                d_1294___mcc_h0_ = source20_.cf18
                d_1295_cf18_ = d_1294___mcc_h0_
                d_1296_v37_: _dafny.Array
                nw176_ = _dafny.Array(int(0), 6)
                d_1296_v37_ = nw176_
                d_1297_v38_: _dafny.Seq
                d_1297_v38_ = _dafny.SeqWithoutIsStrInference([d_1296_v37_, d_1296_v37_, d_1296_v37_])
                d_1260_v4_ = len(d_1297_v38_)
                index199_ = default__.safeIndex(864, (d_1296_v37_).length(0))
                (d_1296_v37_)[index199_] = d_1295_cf18_
                index200_ = default__.safeIndex(864, (d_1296_v37_).length(0))
                (d_1296_v37_)[index200_] = default__.fm22(default__.fm29((d_1261_v5_)[default__.safeIndex(-759, len(d_1261_v5_))], d_1295_cf18_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kkiqvf")), d_1295_cf18_, globalState), D0_DC3(D0_DC0(d_1295_cf18_)), p1, globalState)
                index201_ = default__.safeIndex(864, (d_1296_v37_).length(0))
                (d_1296_v37_)[index201_] = default__.fm11(p1, globalState)
                index202_ = default__.safeIndex(864, (d_1296_v37_).length(0))
                (d_1296_v37_)[index202_] = d_1295_cf18_
            elif True:
                d_1298___mcc_h1_ = source20_.cf17
                d_1299_cf17_ = d_1298___mcc_h1_
                (globalState).f3 = p1
                d_1300_v39_: _dafny.Array
                def lambda50_(d_1301_cf17_):
                    def lambda51_(d_1302_i1_):
                        return (d_1302_i1_) + (len(_dafny.SeqWithoutIsStrInference([len(d_1301_cf17_) for d_1303_i2_ in range(default__.abs(317))])))

                    return lambda51_

                init24_ = lambda50_(d_1299_cf17_)
                nw177_ = _dafny.Array(None, 9)
                for i0_24_ in range(nw177_.length(0)):
                    nw177_[i0_24_] = init24_(i0_24_)
                d_1300_v39_ = nw177_
                index203_ = default__.safeIndex(678, (d_1300_v39_).length(0))
                (d_1300_v39_)[index203_] = (0) - (((d_1279_v22_).f14 if p1 else (d_1279_v22_).f14))
                index204_ = default__.safeIndex(684, (d_1300_v39_).length(0))
                (d_1300_v39_)[index204_] = (d_1279_v22_).f14
                d_1304_v40_: _dafny.Array
                nw178_ = _dafny.Array(None, 3)
                nw178_[int(0)] = not(p1)
                nw178_[int(1)] = (d_1265_v8_) in (d_1256_v0_)
                nw178_[int(2)] = p1
                d_1304_v40_ = nw178_
                index205_ = default__.safeIndex(574, (d_1304_v40_).length(0))
                (d_1304_v40_)[index205_] = p1
                index206_ = default__.safeIndex(678, (d_1300_v39_).length(0))
                index207_ = default__.safeIndex(684, (d_1300_v39_).length(0))
                index208_ = default__.safeIndex(574, (d_1304_v40_).length(0))
                rhs203_ = (d_1279_v22_).f14
                rhs204_ = (len((((d_1263_v7_)[p1] if (p1) in (d_1263_v7_) else (self).fm1(d_1268_v11_, globalState))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wucl"))))) + ((_dafny.MultiSet([False])).cardinality)
                rhs205_ = p1
                rhs206_ = (d_1260_v4_) == (d_1260_v4_)
                lhs138_ = d_1300_v39_
                lhs139_ = default__.safeIndex(678, (d_1300_v39_).length(0))
                lhs140_ = d_1300_v39_
                lhs141_ = default__.safeIndex(684, (d_1300_v39_).length(0))
                lhs142_ = d_1304_v40_
                lhs143_ = default__.safeIndex(574, (d_1304_v40_).length(0))
                lhs138_[lhs139_] = rhs203_
                lhs140_[lhs141_] = rhs204_
                lhs142_[lhs143_] = rhs205_
                r2 = rhs206_
                d_1305_v41_: _dafny.Map
                d_1305_v41_ = _dafny.Map({(d_1304_v40_)[default__.safeIndex(574, (d_1304_v40_).length(0))]: p1})
                d_1306_v42_: _dafny.Map
                d_1306_v42_ = _dafny.Map({d_1265_v8_: d_1305_v41_})
                d_1307_v43_: C4
                nw179_ = C4()
                nw179_.ctor__((d_1279_v22_).f14, _dafny.Map({default__.fm23((d_1304_v40_)[default__.safeIndex(574, (d_1304_v40_).length(0))], d_1306_v42_, (d_1304_v40_)[default__.safeIndex(574, (d_1304_v40_).length(0))], (d_1279_v22_).f14, globalState): 848}))
                d_1307_v43_ = nw179_
                (globalState).f13 = (d_1304_v40_)[default__.safeIndex(574, (d_1304_v40_).length(0))]
            (d_1279_v22_).m5((d_1256_v0_ if True else d_1256_v0_), globalState)
        d_1308_v44_: _dafny.Seq
        d_1308_v44_ = _dafny.SeqWithoutIsStrInference([d_1256_v0_])
        d_1309_v45_: D0
        d_1309_v45_ = D0_DC1(p1, d_1308_v44_, d_1256_v0_)
        d_1310_v46_: _dafny.Map
        d_1310_v46_ = _dafny.Map({(d_1309_v45_).cf3: d_1268_v11_})
        r0 = ((d_1310_v46_)[_dafny.SeqWithoutIsStrInference([d_1265_v8_ for d_1311_i3_ in range(default__.abs(67))])] if (_dafny.SeqWithoutIsStrInference([d_1265_v8_ for d_1312_i3_ in range(default__.abs(67))])) in (d_1310_v46_) else (default__.fm31(globalState)) | (d_1268_v11_))
        r1 = default__.fm25(p1, default__.fm30(globalState), globalState)
        r2 = p1
        d_1313_v47_: _dafny.Array
        nw180_ = _dafny.Array(_dafny.MultiSet({}), 16)
        d_1313_v47_ = nw180_
        r3 = d_1313_v47_
        return r0, r1, r2, r3


class C7(T1, T0):
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C7"
    def ctor__(self):
        pass
        pass

    def fm0(self, p0, p1, globalState):
        return (0) - ((0) - ((D0_DC0((0) - ((0) - (len(_dafny.SeqWithoutIsStrInference([True, True, False])))))).cf0))

    def fm1(self, p0, globalState):
        return (D0_DC1(False, _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cfv")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lo"))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lohm")))).cf3

    def m0(self, p0, globalState):
        r0: bool = False
        r1: int = int(0)
        d_1314_v0_: int
        d_1314_v0_ = 57
        d_1315_v1_: _dafny.Map
        d_1315_v1_ = _dafny.Map({d_1314_v0_: d_1314_v0_})
        d_1315_v1_ = (d_1315_v1_).set(d_1314_v0_, 31)
        d_1316_v2_: bool
        d_1316_v2_ = False
        d_1317_v3_: D0
        d_1317_v3_ = D0_DC0(664)
        d_1318_v4_: D0
        d_1318_v4_ = D0_DC3(d_1317_v3_)
        pat_let_tv25_ = d_1317_v3_
        def iife91_(_pat_let27_0):
            def iife92_(d_1319_dt__update__tmp_h0_):
                def iife93_(_pat_let28_0):
                    def iife94_(d_1320_dt__update_hcf5_h0_):
                        return D0_DC3(d_1320_dt__update_hcf5_h0_)
                    return iife94_(_pat_let28_0)
                return iife93_(pat_let_tv25_)
            return iife92_(_pat_let27_0)
        (globalState).f13 = default__.fm2(p0, (iife91_(d_1318_v4_) if d_1316_v2_ else d_1318_v4_), default__.safeDivisionInt(d_1314_v0_, (0) - (d_1314_v0_)), globalState)
        d_1314_v0_ = (self).fm0((self).fm0(d_1314_v0_, p0, globalState), p0, globalState)
        d_1321_v5_: _dafny.Seq
        d_1321_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iqufbbgna"))
        d_1322_v6_: _dafny.Seq
        d_1322_v6_ = _dafny.SeqWithoutIsStrInference([d_1321_v5_, d_1321_v5_])
        d_1323_v7_: D0
        d_1323_v7_ = D0_DC1(d_1316_v2_, d_1322_v6_, (d_1321_v5_ if d_1316_v2_ else d_1321_v5_))
        source21_ = d_1323_v7_
        if source21_.is_DC1:
            d_1324___mcc_h0_ = source21_.cf1
            d_1325___mcc_h1_ = source21_.cf2
            d_1326___mcc_h2_ = source21_.cf3
            d_1327_cf3_ = d_1326___mcc_h2_
            d_1328_cf2_ = d_1325___mcc_h1_
            d_1329_cf1_ = d_1324___mcc_h0_
            d_1330_v8_: C5
            nw181_ = C5()
            nw181_.ctor__()
            d_1330_v8_ = nw181_
            d_1331_v9_: T1
            nw182_ = C5()
            nw182_.ctor__()
            d_1331_v9_ = nw182_
            d_1332_v10_: _dafny.Array
            def lambda52_(d_1333_v2_):
                def lambda53_(d_1334_i0_):
                    return (D0_DC2(d_1333_v2_)).cf4

                return lambda53_

            init25_ = lambda52_(d_1316_v2_)
            nw183_ = _dafny.Array(None, 13)
            for i0_25_ in range(nw183_.length(0)):
                nw183_[i0_25_] = init25_(i0_25_)
            d_1332_v10_ = nw183_
            index209_ = default__.safeIndex(335, (d_1332_v10_).length(0))
            (d_1332_v10_)[index209_] = d_1316_v2_
            d_1335_v11_: _dafny.Seq
            d_1335_v11_ = _dafny.SeqWithoutIsStrInference([d_1329_cf1_, d_1329_cf1_, d_1316_v2_, d_1316_v2_])
            index210_ = default__.safeIndex(335, (d_1332_v10_).length(0))
            rhs207_ = d_1329_cf1_
            rhs208_ = (d_1331_v9_ if d_1329_cf1_ else (d_1331_v9_ if d_1329_cf1_ else d_1331_v9_))
            rhs209_ = d_1316_v2_
            rhs210_ = (d_1316_v2_) and ((len(d_1335_v11_)) < (len(_dafny.SeqWithoutIsStrInference([d_1314_v0_, 463, d_1314_v0_]))))
            lhs144_ = globalState
            lhs145_ = d_1332_v10_
            lhs146_ = default__.safeIndex(335, (d_1332_v10_).length(0))
            d_1316_v2_ = rhs207_
            d_1331_v9_ = rhs208_
            lhs144_.f3 = rhs209_
            lhs145_[lhs146_] = rhs210_
            if d_1316_v2_:
                d_1335_v11_ = (default__.fm27((d_1314_v0_) - (d_1314_v0_), (d_1314_v0_) * (d_1314_v0_), (d_1314_v0_) + (d_1314_v0_), globalState)).set(default__.safeIndex(d_1314_v0_, len(default__.fm27((d_1314_v0_) - (d_1314_v0_), (d_1314_v0_) * (d_1314_v0_), (d_1314_v0_) + (d_1314_v0_), globalState))), (d_1332_v10_)[default__.safeIndex(335, (d_1332_v10_).length(0))])
                d_1336_v12_: str
                d_1336_v12_ = _dafny.CodePoint('o')
                d_1336_v12_ = p0
                (globalState).f3 = not(d_1316_v2_)
                d_1337_v13_: C5
                nw184_ = C5()
                nw184_.ctor__()
                d_1337_v13_ = nw184_
                d_1338_v14_: _dafny.Set
                d_1338_v14_ = _dafny.Set({d_1314_v0_, d_1314_v0_})
                d_1339_v15_: _dafny.Seq
                d_1339_v15_ = _dafny.SeqWithoutIsStrInference([len(d_1338_v14_), 83])
                d_1339_v15_ = (d_1339_v15_) + (d_1339_v15_)
            elif True:
                d_1314_v0_ = d_1314_v0_
                r1 = d_1314_v0_
                d_1340_v16_: _dafny.Map
                d_1340_v16_ = _dafny.Map({(d_1332_v10_)[default__.safeIndex(335, (d_1332_v10_).length(0))]: d_1318_v4_})
                d_1341_v17_: _dafny.Map
                d_1341_v17_ = _dafny.Map({d_1314_v0_: d_1316_v2_})
                d_1342_v18_: T0
                nw185_ = C2()
                nw185_.ctor__()
                d_1342_v18_ = nw185_
                d_1343_v19_: _dafny.Map
                d_1343_v19_ = _dafny.Map({len(d_1341_v17_): d_1342_v18_})
                d_1344_v20_: D2
                d_1344_v20_ = D2_DC7(True, d_1340_v16_, len((d_1343_v19_).set(d_1314_v0_, d_1342_v18_)), d_1329_cf1_)
                d_1345_v21_: D2
                d_1345_v21_ = D2_DC7((d_1344_v20_).cf14, d_1340_v16_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vwxlge"))), d_1316_v2_)
                d_1314_v0_ = default__.fm22(d_1345_v21_, d_1318_v4_, True, globalState)
                d_1346_v22_: _dafny.Set
                d_1346_v22_ = _dafny.Set({not(d_1329_cf1_), d_1316_v2_, d_1316_v2_})
                d_1347_v23_: _dafny.Set
                d_1347_v23_ = _dafny.Set({d_1346_v22_})
                d_1348_v24_: D6
                d_1348_v24_ = D6_DC16(len(_dafny.SeqWithoutIsStrInference([p0 for d_1349_i1_ in range(default__.abs(-888))])))
                index211_ = default__.safeIndex(335, (d_1332_v10_).length(0))
                rhs211_ = (d_1347_v23_).isdisjoint(_dafny.Set({default__.fm33(d_1329_cf1_, d_1329_cf1_, (d_1332_v10_)[default__.safeIndex(335, (d_1332_v10_).length(0))], globalState), d_1346_v22_, d_1346_v22_, d_1346_v22_, d_1346_v22_}))
                rhs212_ = (d_1314_v0_) * (d_1314_v0_)
                rhs213_ = d_1329_cf1_
                rhs214_ = (d_1348_v24_).cf24
                rhs215_ = (d_1332_v10_)[default__.safeIndex(335, (d_1332_v10_).length(0))]
                lhs147_ = globalState
                lhs148_ = d_1332_v10_
                lhs149_ = default__.safeIndex(335, (d_1332_v10_).length(0))
                lhs150_ = globalState
                lhs147_.f5 = rhs211_
                r1 = rhs212_
                lhs148_[lhs149_] = rhs213_
                r1 = rhs214_
                lhs150_.f3 = rhs215_
                r1 = (d_1314_v0_) + (d_1314_v0_)
            d_1350_v25_: _dafny.Array
            nw186_ = _dafny.Array(int(0), 27)
            d_1350_v25_ = nw186_
            r1 = len(_dafny.Set({d_1350_v25_, d_1350_v25_, d_1350_v25_, d_1350_v25_}))
        elif source21_.is_DC2:
            d_1351___mcc_h3_ = source21_.cf4
            d_1352_cf4_ = d_1351___mcc_h3_
            d_1353_v26_: _dafny.Array
            nw187_ = _dafny.Array(int(0), 21)
            d_1353_v26_ = nw187_
            d_1354_v27_: _dafny.Map
            d_1354_v27_ = _dafny.Map({d_1353_v26_: 636})
            d_1355_v28_: D8
            d_1355_v28_ = D8_DC22(d_1354_v27_)
            d_1356_v29_: _dafny.Seq
            d_1356_v29_ = _dafny.SeqWithoutIsStrInference([d_1314_v0_])
            pat_let_tv26_ = d_1353_v26_
            pat_let_tv27_ = d_1314_v0_
            def iife95_(_pat_let29_0):
                def iife96_(d_1357_dt__update__tmp_h1_):
                    def iife97_(_pat_let30_0):
                        def iife98_(d_1358_dt__update_hcf34_h0_):
                            return D8_DC22(d_1358_dt__update_hcf34_h0_)
                        return iife98_(_pat_let30_0)
                    return iife97_(_dafny.Map({pat_let_tv26_: pat_let_tv27_}))
                return iife96_(_pat_let29_0)
            rhs216_ = iife95_(d_1355_v28_)
            rhs217_ = (0) - ((d_1314_v0_ if d_1316_v2_ else d_1314_v0_))
            rhs218_ = (len((d_1356_v29_ if d_1352_cf4_ else _dafny.SeqWithoutIsStrInference([d_1314_v0_ for d_1359_i2_ in range(default__.abs(41))])))) * ((0) - (d_1314_v0_))
            d_1355_v28_ = rhs216_
            r1 = rhs217_
            r1 = rhs218_
            d_1360_v30_: _dafny.Map
            d_1360_v30_ = _dafny.Map({d_1316_v2_: len(_dafny.Set({d_1352_cf4_, d_1316_v2_, default__.fm6(globalState)}))})
            d_1315_v1_ = (d_1315_v1_).set(len(d_1360_v30_), (0) - (default__.safeModuloInt(-311, d_1314_v0_)))
            d_1361_v32_: C4
            nw188_ = C4()
            def iife99_():
                coll37_ = _dafny.Map()
                compr_37_: str
                for compr_37_ in (_dafny.MultiSet([p0, p0])).Elements:
                    d_1362_v31_: str = compr_37_
                    if (d_1362_v31_) in (_dafny.MultiSet([p0, p0])):
                        coll37_[d_1362_v31_] = d_1314_v0_
                return _dafny.Map(coll37_)
            nw188_.ctor__(d_1314_v0_, iife99_()
            )
            d_1361_v32_ = nw188_
            d_1363_v33_: C1
            nw189_ = C1()
            nw189_.ctor__()
            d_1363_v33_ = nw189_
            nw190_ = C1()
            nw190_.ctor__()
            d_1363_v33_ = nw190_
        elif source21_.is_DC0:
            d_1364___mcc_h4_ = source21_.cf0
            d_1365_cf0_ = d_1364___mcc_h4_
            d_1366_v34_: _dafny.Map
            d_1366_v34_ = _dafny.Map({not((d_1316_v2_ if d_1316_v2_ else d_1316_v2_)): d_1316_v2_})
            d_1366_v34_ = (d_1366_v34_).set(d_1316_v2_, d_1316_v2_)
            (globalState).f9 = not(d_1316_v2_)
            d_1365_cf0_ = d_1314_v0_
            d_1367_v35_: _dafny.Map
            d_1367_v35_ = _dafny.Map({d_1321_v5_: (not(d_1316_v2_) if d_1316_v2_ else d_1316_v2_)})
            d_1367_v35_ = (d_1367_v35_).set(d_1321_v5_, not((default__.fm11(d_1316_v2_, globalState)) == (d_1314_v0_)))
        elif True:
            d_1368___mcc_h5_ = source21_.cf5
            d_1369_cf5_ = d_1368___mcc_h5_
            d_1370_v36_: _dafny.MultiSet
            d_1370_v36_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([p0])])
            d_1371_v37_: _dafny.Map
            d_1371_v37_ = _dafny.Map({(d_1370_v36_ if d_1316_v2_ else d_1370_v36_): d_1314_v0_})
            d_1372_v38_: _dafny.Set
            d_1372_v38_ = _dafny.Set({not(d_1316_v2_)})
            d_1373_v39_: _dafny.Seq
            d_1373_v39_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({d_1316_v2_}), d_1372_v38_, _dafny.Set({False, d_1316_v2_})])
            d_1371_v37_ = (d_1371_v37_).set(d_1370_v36_, len((d_1373_v39_)[default__.safeIndex(d_1314_v0_, len(d_1373_v39_))]))
            d_1374_v40_: D4
            d_1374_v40_ = D4_DC12(d_1314_v0_)
            source22_ = d_1374_v40_
            if source22_.is_DC11:
                r0 = (d_1316_v2_) or (d_1316_v2_)
                d_1375_v41_: _dafny.Map
                d_1375_v41_ = _dafny.Map({d_1316_v2_: d_1316_v2_})
                d_1376_v42_: _dafny.Set
                d_1376_v42_ = _dafny.Set({d_1314_v0_, d_1314_v0_, d_1314_v0_, (len(d_1321_v5_)) + (len(d_1375_v41_))})
                d_1376_v42_ = d_1376_v42_
                d_1377_v43_: _dafny.Array
                def lambda54_(d_1378_v42_, d_1379_v0_):
                    def lambda55_(d_1380_i3_):
                        return (_dafny.MultiSet([_dafny.Set({d_1379_v0_})])).ispropersubset(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_1378_v42_, d_1378_v42_])))

                    return lambda55_

                init26_ = lambda54_(d_1376_v42_, d_1314_v0_)
                nw191_ = _dafny.Array(None, 29)
                for i0_26_ in range(nw191_.length(0)):
                    nw191_[i0_26_] = init26_(i0_26_)
                d_1377_v43_ = nw191_
                d_1381_v44_: D8
                d_1381_v44_ = D8_DC23(d_1316_v2_, True, d_1314_v0_)
                index212_ = default__.safeIndex(569, (d_1377_v43_).length(0))
                (d_1377_v43_)[index212_] = (d_1381_v44_).cf35
                index213_ = default__.safeIndex(569, (d_1377_v43_).length(0))
                (d_1377_v43_)[index213_] = d_1316_v2_
                rhs219_ = (d_1377_v43_)[default__.safeIndex(569, (d_1377_v43_).length(0))]
                rhs220_ = d_1321_v5_
                lhs151_ = globalState
                lhs151_.f9 = rhs219_
                d_1321_v5_ = rhs220_
            elif source22_.is_DC12:
                d_1382___mcc_h6_ = source22_.cf18
                d_1383_cf18_ = d_1382___mcc_h6_
                d_1384_v45_: _dafny.Set
                d_1384_v45_ = _dafny.Set({d_1383_cf18_, (d_1370_v36_).cardinality})
                d_1385_v46_: _dafny.MultiSet
                d_1385_v46_ = _dafny.MultiSet([d_1384_v45_])
                rhs221_ = (d_1383_cf18_ if d_1316_v2_ else len(_dafny.SeqWithoutIsStrInference([p0 for d_1386_i4_ in range(default__.abs(600))])))
                rhs222_ = (d_1385_v46_).isdisjoint(d_1385_v46_)
                lhs152_ = globalState
                d_1383_cf18_ = rhs221_
                lhs152_.f13 = rhs222_
                d_1387_v47_: C1
                nw192_ = C1()
                nw192_.ctor__()
                d_1387_v47_ = nw192_
                d_1388_v48_: _dafny.Map
                d_1388_v48_ = _dafny.Map({_dafny.Map({d_1314_v0_: d_1383_cf18_}): d_1384_v45_})
                d_1389_v49_: _dafny.Array
                nw193_ = _dafny.Array(None, 19)
                nw193_[int(0)] = d_1316_v2_
                nw193_[int(1)] = not(d_1316_v2_)
                nw193_[int(2)] = d_1316_v2_
                nw193_[int(3)] = (d_1314_v0_) <= (d_1383_cf18_)
                nw193_[int(4)] = d_1316_v2_
                nw193_[int(5)] = d_1316_v2_
                nw193_[int(6)] = d_1316_v2_
                nw193_[int(7)] = d_1316_v2_
                nw193_[int(8)] = d_1316_v2_
                nw193_[int(9)] = d_1316_v2_
                nw193_[int(10)] = (d_1315_v1_) in (d_1388_v48_)
                nw193_[int(11)] = d_1316_v2_
                nw193_[int(12)] = d_1316_v2_
                nw193_[int(13)] = (d_1314_v0_) < (d_1314_v0_)
                nw193_[int(14)] = d_1316_v2_
                nw193_[int(15)] = d_1316_v2_
                nw193_[int(16)] = (d_1316_v2_) and (d_1316_v2_)
                nw193_[int(17)] = d_1316_v2_
                nw193_[int(18)] = d_1316_v2_
                d_1389_v49_ = nw193_
                d_1390_v50_: _dafny.MultiSet
                d_1390_v50_ = _dafny.MultiSet([d_1316_v2_])
                d_1391_v51_: _dafny.Map
                d_1391_v51_ = _dafny.Map({d_1383_cf18_: d_1387_v47_})
                d_1392_v52_: _dafny.Map
                d_1392_v52_ = _dafny.Map({d_1316_v2_: 815})
                rhs223_ = (not(d_1316_v2_)) or ((((d_1390_v50_).set(d_1316_v2_, default__.abs(-436))).cardinality) == ((0) - (d_1314_v0_)))
                rhs224_ = ((d_1391_v51_)[d_1383_cf18_] if (d_1383_cf18_) in (d_1391_v51_) else d_1387_v47_)
                rhs225_ = d_1389_v49_
                rhs226_ = (d_1383_cf18_) + (len(d_1321_v5_))
                rhs227_ = (self).fm0(((d_1392_v52_)[d_1316_v2_] if (d_1316_v2_) in (d_1392_v52_) else d_1383_cf18_), (p0 if not(d_1316_v2_) else p0), globalState)
                lhs153_ = globalState
                lhs153_.f13 = rhs223_
                d_1387_v47_ = rhs224_
                d_1389_v49_ = rhs225_
                r1 = rhs226_
                r1 = rhs227_
                d_1316_v2_ = d_1316_v2_
                d_1321_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uf"))
            elif True:
                d_1393___mcc_h7_ = source22_.cf17
                d_1394_cf17_ = d_1393___mcc_h7_
                d_1395_v53_: _dafny.Array
                nw194_ = _dafny.Array(False, 7)
                d_1395_v53_ = nw194_
                d_1396_v54_: _dafny.MultiSet
                d_1396_v54_ = _dafny.MultiSet([d_1316_v2_])
                d_1397_v55_: D6
                d_1397_v55_ = D6_DC15(d_1396_v54_)
                d_1398_v56_: _dafny.Map
                d_1398_v56_ = _dafny.Map({d_1316_v2_: (not(not(d_1316_v2_)) if d_1316_v2_ else True)})
                rhs228_ = d_1395_v53_
                rhs229_ = d_1397_v55_
                rhs230_ = d_1314_v0_
                rhs231_ = ((d_1398_v56_)[d_1316_v2_] if (d_1316_v2_) in (d_1398_v56_) else d_1316_v2_)
                rhs232_ = (0) - (default__.fm11((d_1314_v0_) == (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lecq")))), globalState))
                lhs154_ = globalState
                d_1395_v53_ = rhs228_
                d_1397_v55_ = rhs229_
                r1 = rhs230_
                lhs154_.f9 = rhs231_
                d_1314_v0_ = rhs232_
                index214_ = default__.safeIndex(502, (d_1395_v53_).length(0))
                (d_1395_v53_)[index214_] = d_1316_v2_
                index215_ = default__.safeIndex(37, (d_1395_v53_).length(0))
                (d_1395_v53_)[index215_] = not (default__.fm6(globalState)) or (d_1316_v2_)
                d_1399_v57_: _dafny.Seq
                d_1399_v57_ = _dafny.SeqWithoutIsStrInference([d_1316_v2_, d_1316_v2_])
                d_1400_v58_: _dafny.MultiSet
                d_1400_v58_ = _dafny.MultiSet([p0])
                d_1401_v59_: _dafny.MultiSet
                d_1401_v59_ = _dafny.MultiSet([(0) - ((len(d_1399_v57_)) * ((d_1400_v58_).cardinality)), (d_1314_v0_) - (d_1314_v0_), d_1314_v0_])
                index216_ = default__.safeIndex(502, (d_1395_v53_).length(0))
                index217_ = default__.safeIndex(37, (d_1395_v53_).length(0))
                rhs233_ = d_1316_v2_
                rhs234_ = (d_1314_v0_) < (d_1314_v0_)
                rhs235_ = (d_1401_v59_) - (d_1401_v59_)
                lhs155_ = d_1395_v53_
                lhs156_ = default__.safeIndex(502, (d_1395_v53_).length(0))
                lhs157_ = d_1395_v53_
                lhs158_ = default__.safeIndex(37, (d_1395_v53_).length(0))
                lhs155_[lhs156_] = rhs233_
                lhs157_[lhs158_] = rhs234_
                d_1401_v59_ = rhs235_
                d_1402_v60_: _dafny.Map
                d_1402_v60_ = _dafny.Map({True: d_1318_v4_})
                d_1403_v61_: D2
                d_1403_v61_ = D2_DC7((d_1395_v53_)[default__.safeIndex(502, (d_1395_v53_).length(0))], d_1402_v60_, 229, d_1316_v2_)
                d_1404_v62_: _dafny.Array
                def lambda56_(d_1405_v0_):
                    def lambda57_(d_1406_i5_):
                        return (d_1406_i5_) + (d_1405_v0_)

                    return lambda57_

                init27_ = lambda56_(d_1314_v0_)
                nw195_ = _dafny.Array(None, 4)
                for i0_27_ in range(nw195_.length(0)):
                    nw195_[i0_27_] = init27_(i0_27_)
                d_1404_v62_ = nw195_
                d_1407_v63_: _dafny.Map
                d_1407_v63_ = _dafny.Map({d_1404_v62_: d_1314_v0_})
                pat_let_tv28_ = d_1317_v3_
                def iife100_(_pat_let31_0):
                    def iife101_(d_1408_dt__update__tmp_h2_):
                        def iife102_(_pat_let32_0):
                            def iife103_(d_1409_dt__update_hcf5_h1_):
                                return D0_DC3(d_1409_dt__update_hcf5_h1_)
                            return iife103_(_pat_let32_0)
                        return iife102_(pat_let_tv28_)
                    return iife101_(_pat_let31_0)
                rhs236_ = ((d_1401_v59_).intersection(d_1401_v59_)).set(default__.fm22(d_1403_v61_, d_1318_v4_, (d_1395_v53_)[default__.safeIndex(502, (d_1395_v53_).length(0))], globalState), default__.abs((0) - (d_1314_v0_)))
                rhs237_ = default__.fm22(d_1403_v61_, iife100_(d_1318_v4_), (d_1399_v57_)[default__.safeIndex(d_1314_v0_, len(d_1399_v57_))], globalState)
                rhs238_ = not((d_1404_v62_) in (d_1407_v63_))
                rhs239_ = default__.fm19(d_1314_v0_, not (d_1316_v2_) or (d_1316_v2_), globalState)
                lhs159_ = globalState
                lhs160_ = globalState
                d_1401_v59_ = rhs236_
                d_1314_v0_ = rhs237_
                lhs159_.f5 = rhs238_
                lhs160_.f13 = rhs239_
                d_1410_v64_: D8
                d_1410_v64_ = D8_DC24()
                d_1410_v64_ = d_1410_v64_
            d_1411_v65_: C6
            nw196_ = C6()
            nw196_.ctor__()
            d_1411_v65_ = nw196_
            d_1412_v66_: _dafny.Seq
            d_1412_v66_ = _dafny.SeqWithoutIsStrInference([d_1411_v65_])
            d_1413_v67_: D11
            d_1413_v67_ = D11_DC29(_dafny.MultiSet(d_1412_v66_))
            d_1414_v68_: _dafny.Array
            nw197_ = _dafny.Array(None, 6)
            nw197_[int(0)] = default__.safeModuloInt(d_1314_v0_, 940)
            nw197_[int(1)] = d_1314_v0_
            nw197_[int(2)] = (self).fm0(d_1314_v0_, p0, globalState)
            nw197_[int(3)] = default__.safeDivisionInt(587, d_1314_v0_)
            nw197_[int(4)] = d_1314_v0_
            nw197_[int(5)] = (d_1314_v0_) + (((d_1413_v67_).cf46).cardinality)
            d_1414_v68_ = nw197_
            d_1414_v68_ = d_1414_v68_
            d_1415_v69_: _dafny.Seq
            d_1415_v69_ = _dafny.SeqWithoutIsStrInference([d_1314_v0_])
            d_1416_v70_: _dafny.Map
            d_1416_v70_ = _dafny.Map({len(d_1415_v69_): d_1321_v5_})
            d_1416_v70_ = (d_1416_v70_) | (d_1416_v70_)
        d_1417_v71_: _dafny.Map
        d_1417_v71_ = _dafny.Map({d_1316_v2_: d_1316_v2_})
        d_1418_v72_: _dafny.Map
        d_1418_v72_ = _dafny.Map({d_1316_v2_: d_1417_v71_})
        d_1419_v73_: _dafny.Seq
        d_1419_v73_ = _dafny.SeqWithoutIsStrInference([d_1314_v0_, d_1314_v0_])
        d_1420_v74_: _dafny.Seq
        d_1420_v74_ = _dafny.SeqWithoutIsStrInference([d_1316_v2_])
        d_1421_v75_: _dafny.Array
        def lambda58_(d_1422_v5_):
            def lambda59_(d_1423_i6_):
                return d_1422_v5_

            return lambda59_

        init28_ = lambda58_(d_1321_v5_)
        nw198_ = _dafny.Array(None, 25)
        for i0_28_ in range(nw198_.length(0)):
            nw198_[i0_28_] = init28_(i0_28_)
        d_1421_v75_ = nw198_
        d_1424_v76_: D7
        d_1424_v76_ = D7_DC19(D1_DC5(d_1316_v2_, d_1321_v5_, d_1420_v74_), len(d_1321_v5_), d_1316_v2_, d_1421_v75_)
        d_1425_v77_: _dafny.Seq
        d_1425_v77_ = _dafny.SeqWithoutIsStrInference([d_1419_v73_])
        d_1426_v78_: _dafny.Seq
        d_1426_v78_ = _dafny.SeqWithoutIsStrInference([(d_1425_v77_)[default__.safeIndex(d_1314_v0_, len(d_1425_v77_))], d_1419_v73_])
        d_1427_v79_: _dafny.Array
        nw199_ = _dafny.Array(None, 15)
        nw199_[int(0)] = d_1314_v0_
        nw199_[int(1)] = (0) - (d_1314_v0_)
        nw199_[int(2)] = d_1314_v0_
        nw199_[int(3)] = d_1314_v0_
        nw199_[int(4)] = len((((d_1418_v72_)[d_1316_v2_] if (d_1316_v2_) in (d_1418_v72_) else _dafny.Map({d_1316_v2_: d_1316_v2_}))) | (d_1417_v71_))
        nw199_[int(5)] = d_1314_v0_
        nw199_[int(6)] = d_1314_v0_
        nw199_[int(7)] = (0) - (d_1314_v0_)
        nw199_[int(8)] = d_1314_v0_
        nw199_[int(9)] = len(d_1419_v73_)
        nw199_[int(10)] = ((d_1424_v76_).cf28 if not(d_1316_v2_) else len(d_1425_v77_))
        nw199_[int(11)] = d_1314_v0_
        nw199_[int(12)] = d_1314_v0_
        nw199_[int(13)] = (d_1314_v0_) * (d_1314_v0_)
        nw199_[int(14)] = d_1314_v0_
        d_1427_v79_ = nw199_
        index218_ = default__.safeIndex(320, (d_1427_v79_).length(0))
        (d_1427_v79_)[index218_] = (460) + (d_1314_v0_)
        index219_ = default__.safeIndex(320, (d_1427_v79_).length(0))
        (d_1427_v79_)[index219_] = (0) - ((0) - (d_1314_v0_))
        d_1428_i7_: int
        d_1428_i7_ = 0
        with _dafny.label("9"):
            while d_1316_v2_:
                with _dafny.c_label("9"):
                    if (d_1428_i7_) >= (100):
                        raise _dafny.Break("9")
                    d_1428_i7_ = (d_1428_i7_) + (1)
                    d_1429_v80_: _dafny.Map
                    d_1429_v80_ = _dafny.Map({p0: len(d_1321_v5_)})
                    d_1430_v81_: C4
                    nw200_ = C4()
                    nw200_.ctor__((d_1427_v79_)[default__.safeIndex(320, (d_1427_v79_).length(0))], d_1429_v80_)
                    d_1430_v81_ = nw200_
                    d_1431_v82_: _dafny.Map
                    d_1431_v82_ = _dafny.Map({p0: d_1430_v81_})
                    d_1431_v82_ = ((d_1431_v82_) | (d_1431_v82_)) | ((d_1431_v82_) | (d_1431_v82_))
                    (globalState).f13 = d_1316_v2_
                    d_1432_v83_: _dafny.Map
                    d_1432_v83_ = _dafny.Map({not(d_1316_v2_): d_1318_v4_})
                    d_1433_v86_: _dafny.MultiSet
                    d_1433_v86_ = _dafny.MultiSet([p0, p0, p0])
                    d_1434_v87_: D9
                    d_1434_v87_ = D9_DC27(not(d_1316_v2_), len(d_1321_v5_), d_1316_v2_)
                    d_1435_v88_: D2
                    def iife104_():
                        def iife106_():
                            coll40_ = _dafny.Map()
                            compr_40_: str
                            for compr_40_ in (d_1433_v86_).Elements:
                                d_1436_v85_: str = compr_40_
                                if (d_1436_v85_) in (d_1433_v86_):
                                    coll40_[d_1436_v85_] = (d_1430_v81_).f14
                            return _dafny.Map(coll40_)
                        coll38_ = _dafny.Map()
                        def iife105_():
                            coll39_ = _dafny.Map()
                            compr_39_: str
                            for compr_39_ in (d_1433_v86_).Elements:
                                d_1436_v85_: str = compr_39_
                                if (d_1436_v85_) in (d_1433_v86_):
                                    coll39_[d_1436_v85_] = (d_1430_v81_).f14
                            return _dafny.Map(coll39_)
                        compr_38_: str
                        for compr_38_ in (iife105_()
                        ).keys.Elements:
                            d_1437_v84_: str = compr_38_
                            if (d_1437_v84_) in (iife106_()
                            ):
                                coll38_[d_1437_v84_] = _dafny.Set({d_1316_v2_})
                        return _dafny.Map(coll38_)
                    d_1435_v88_ = D2_DC7(d_1316_v2_, d_1432_v83_, len(iife104_()
), not((d_1434_v87_).cf42))
                    r1 = default__.fm22(d_1435_v88_, d_1318_v4_, d_1316_v2_, globalState)
                    d_1438_v89_: _dafny.Array
                    nw201_ = _dafny.Array(_dafny.Array(None, 0), 25)
                    d_1438_v89_ = nw201_
                    d_1439_v90_: C1
                    nw202_ = C1()
                    nw202_.ctor__()
                    d_1439_v90_ = nw202_
                    d_1440_v91_: _dafny.Array
                    nw203_ = _dafny.Array(None, 8)
                    nw203_[int(0)] = d_1439_v90_
                    nw203_[int(1)] = d_1439_v90_
                    nw203_[int(2)] = d_1439_v90_
                    nw203_[int(3)] = d_1439_v90_
                    nw203_[int(4)] = d_1439_v90_
                    nw203_[int(5)] = d_1439_v90_
                    nw203_[int(6)] = d_1439_v90_
                    nw203_[int(7)] = d_1439_v90_
                    d_1440_v91_ = nw203_
                    d_1441_v92_: _dafny.Map
                    d_1441_v92_ = _dafny.Map({d_1440_v91_: (d_1430_v81_).f14})
                    rhs240_ = d_1438_v89_
                    rhs241_ = ((d_1441_v92_)[d_1440_v91_] if (d_1440_v91_) in (d_1441_v92_) else (d_1314_v0_) * (448))
                    d_1438_v89_ = rhs240_
                    d_1314_v0_ = rhs241_
                    pass
            pass
        d_1442_v93_: _dafny.MultiSet
        d_1442_v93_ = _dafny.MultiSet([d_1316_v2_])
        d_1443_v94_: _dafny.MultiSet
        d_1443_v94_ = _dafny.MultiSet([(d_1427_v79_)[default__.safeIndex(320, (d_1427_v79_).length(0))], (d_1442_v93_).cardinality])
        r0 = (d_1443_v94_).issubset((_dafny.MultiSet([d_1314_v0_, (d_1427_v79_)[default__.safeIndex(320, (d_1427_v79_).length(0))]])).set(d_1314_v0_, default__.abs((d_1427_v79_)[default__.safeIndex(320, (d_1427_v79_).length(0))])))
        r1 = (d_1427_v79_)[default__.safeIndex(320, (d_1427_v79_).length(0))]
        return r0, r1

    def m1(self, p0, p1, globalState):
        r0: _dafny.Set = _dafny.Set({})
        r1: bool = False
        r2: int = int(0)
        r3: bool = False
        d_1444_v0_: _dafny.MultiSet
        d_1444_v0_ = _dafny.MultiSet([p0, p0, True])
        d_1445_v1_: D6
        d_1445_v1_ = D6_DC16(((d_1444_v0_).cardinality if p0 else 489))
        source23_ = d_1445_v1_
        if source23_.is_DC16:
            d_1446___mcc_h0_ = source23_.cf24
            d_1447_cf24_ = d_1446___mcc_h0_
            d_1448_v2_: D0
            d_1448_v2_ = D0_DC0(d_1447_cf24_)
            d_1449_v3_: D0
            d_1449_v3_ = D0_DC3(d_1448_v2_)
            d_1450_v4_: _dafny.Map
            d_1450_v4_ = _dafny.Map({p0: d_1449_v3_})
            d_1451_v5_: D2
            d_1451_v5_ = D2_DC7(p0, d_1450_v4_, d_1447_cf24_, p0)
            d_1452_v6_: _dafny.Seq
            d_1452_v6_ = _dafny.SeqWithoutIsStrInference([True])
            d_1453_v7_: str
            d_1453_v7_ = _dafny.CodePoint('f')
            d_1454_v8_: _dafny.Array
            nw204_ = _dafny.Array(None, 22)
            nw204_[int(0)] = p0
            nw204_[int(1)] = p0
            nw204_[int(2)] = default__.fm2(d_1453_v7_, d_1449_v3_, d_1447_cf24_, globalState)
            nw204_[int(3)] = p0
            nw204_[int(4)] = p0
            nw204_[int(5)] = p0
            nw204_[int(6)] = p0
            nw204_[int(7)] = p0
            nw204_[int(8)] = p0
            nw204_[int(9)] = not(p0)
            nw204_[int(10)] = p0
            nw204_[int(11)] = p0
            nw204_[int(12)] = True
            nw204_[int(13)] = not(False)
            nw204_[int(14)] = p0
            nw204_[int(15)] = p0
            nw204_[int(16)] = p0
            nw204_[int(17)] = p0
            nw204_[int(18)] = not(p0)
            nw204_[int(19)] = p0
            nw204_[int(20)] = p0
            nw204_[int(21)] = p0
            d_1454_v8_ = nw204_
            d_1455_v9_: _dafny.Seq
            d_1455_v9_ = _dafny.SeqWithoutIsStrInference([d_1454_v8_, d_1454_v8_])
            d_1456_v10_: _dafny.Map
            d_1456_v10_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_1453_v7_ for d_1457_i0_ in range(default__.abs(-271))]): p0})
            d_1458_v12_: _dafny.Seq
            d_1458_v12_ = _dafny.SeqWithoutIsStrInference([601, d_1447_cf24_])
            d_1459_v13_: _dafny.Map
            d_1459_v13_ = _dafny.Map({d_1447_cf24_: d_1447_cf24_})
            d_1460_v14_: _dafny.Map
            d_1460_v14_ = _dafny.Map({p0: p0})
            d_1461_v15_: _dafny.Set
            d_1461_v15_ = _dafny.Set({True})
            d_1462_v16_: _dafny.Array
            nw205_ = _dafny.Array(None, 24)
            nw205_[int(0)] = default__.fm22(d_1451_v5_, d_1449_v3_, p0, globalState)
            nw205_[int(1)] = (d_1447_cf24_) + ((0) - (d_1447_cf24_))
            nw205_[int(2)] = default__.safeModuloInt(d_1447_cf24_, (0) - (d_1447_cf24_))
            nw205_[int(3)] = d_1447_cf24_
            nw205_[int(4)] = default__.fm22(D2_DC7(p0, d_1450_v4_, d_1447_cf24_, p0), d_1449_v3_, p0, globalState)
            nw205_[int(5)] = d_1447_cf24_
            nw205_[int(6)] = d_1447_cf24_
            nw205_[int(7)] = default__.fm11(p0, globalState)
            nw205_[int(8)] = d_1447_cf24_
            nw205_[int(9)] = len(d_1452_v6_)
            nw205_[int(10)] = d_1447_cf24_
            nw205_[int(11)] = len(d_1455_v9_)
            def iife107_():
                coll41_ = _dafny.Set()
                compr_41_: _dafny.Seq
                for compr_41_ in (d_1456_v10_).keys.Elements:
                    d_1463_v11_: _dafny.Seq = compr_41_
                    if (d_1463_v11_) in (d_1456_v10_):
                        coll41_ = coll41_.union(_dafny.Set([d_1463_v11_]))
                return _dafny.Set(coll41_)
            nw205_[int(12)] = len((_dafny.SeqWithoutIsStrInference([len(d_1452_v6_), len(iife107_()
            )])) + (d_1458_v12_))
            nw205_[int(13)] = d_1447_cf24_
            nw205_[int(14)] = (len(d_1459_v13_)) * (len(d_1460_v14_))
            nw205_[int(15)] = len(_dafny.Map({p0: len(p1)}))
            nw205_[int(16)] = d_1447_cf24_
            nw205_[int(17)] = len(d_1461_v15_)
            nw205_[int(18)] = d_1447_cf24_
            nw205_[int(19)] = d_1447_cf24_
            nw205_[int(20)] = (d_1447_cf24_) - (d_1447_cf24_)
            nw205_[int(21)] = default__.fm22(d_1451_v5_, d_1449_v3_, p0, globalState)
            nw205_[int(22)] = ((d_1444_v0_).cardinality) * (d_1447_cf24_)
            nw205_[int(23)] = d_1447_cf24_
            d_1462_v16_ = nw205_
            rhs242_ = (default__.safeDivisionInt(d_1447_cf24_, d_1447_cf24_)) * (d_1447_cf24_)
            rhs243_ = d_1462_v16_
            d_1447_cf24_ = rhs242_
            d_1462_v16_ = rhs243_
            d_1464_v17_: C5
            nw206_ = C5()
            nw206_.ctor__()
            d_1464_v17_ = nw206_
            d_1465_v18_: _dafny.Map
            d_1465_v18_ = _dafny.Map({d_1464_v17_: d_1462_v16_})
            d_1465_v18_ = (d_1465_v18_).set(d_1464_v17_, d_1462_v16_)
            d_1466_v19_: _dafny.Seq
            d_1466_v19_ = _dafny.SeqWithoutIsStrInference([p1, p1, p1, p1])
            d_1447_cf24_ = default__.safeModuloInt(54, len((d_1466_v19_)[default__.safeIndex(((d_1444_v0_)[p0] if (p0) in (d_1444_v0_) else d_1447_cf24_), len(d_1466_v19_))]))
            d_1467_v20_: D8
            d_1467_v20_ = D8_DC24()
            d_1467_v20_ = d_1467_v20_
        elif source23_.is_DC15:
            d_1468___mcc_h1_ = source23_.cf23
            d_1469_cf23_ = d_1468___mcc_h1_
            d_1470_v21_: int
            d_1470_v21_ = -864
            (globalState).f3 = (d_1470_v21_) >= (d_1470_v21_)
            r2 = d_1470_v21_
            if (True if p0 else (d_1470_v21_) == (d_1470_v21_)):
                d_1470_v21_ = d_1470_v21_
                (globalState).f3 = (d_1470_v21_) >= (d_1470_v21_)
                d_1471_v22_: _dafny.MultiSet
                d_1471_v22_ = _dafny.MultiSet([d_1470_v21_, d_1470_v21_, d_1470_v21_])
                d_1472_v23_: D12
                d_1472_v23_ = D12_DC31(d_1471_v22_)
                d_1473_v24_: _dafny.Array
                nw207_ = _dafny.Array(int(0), 19)
                d_1473_v24_ = nw207_
                d_1474_v25_: _dafny.MultiSet
                d_1474_v25_ = _dafny.MultiSet([d_1473_v24_, d_1473_v24_])
                d_1475_v26_: _dafny.Map
                d_1475_v26_ = _dafny.Map({d_1473_v24_: d_1473_v24_})
                rhs244_ = p0
                rhs245_ = (d_1472_v23_).cf48
                rhs246_ = p0
                rhs247_ = (0) - (((d_1474_v25_)[(((d_1475_v26_)[d_1473_v24_] if (d_1473_v24_) in (d_1475_v26_) else d_1473_v24_) if p0 else d_1473_v24_)] if ((((d_1475_v26_)[d_1473_v24_] if (d_1473_v24_) in (d_1475_v26_) else d_1473_v24_) if p0 else d_1473_v24_)) in (d_1474_v25_) else d_1470_v21_))
                rhs248_ = default__.safeDivisionInt(d_1470_v21_, d_1470_v21_)
                lhs161_ = globalState
                lhs162_ = globalState
                lhs161_.f3 = rhs244_
                d_1471_v22_ = rhs245_
                lhs162_.f3 = rhs246_
                d_1470_v21_ = rhs247_
                d_1470_v21_ = rhs248_
                r2 = d_1470_v21_
                d_1476_v27_: _dafny.Array
                nw208_ = _dafny.Array(False, 7)
                d_1476_v27_ = nw208_
                d_1476_v27_ = d_1476_v27_
            elif True:
                (globalState).f9 = (not(p0)) or ((p0) or (default__.fm19(len(p1), p0, globalState)))
                d_1477_v28_: _dafny.Map
                d_1477_v28_ = _dafny.Map({(d_1470_v21_) * ((0) - (d_1470_v21_)): d_1470_v21_})
                d_1477_v28_ = d_1477_v28_
                (globalState).f13 = (d_1470_v21_) > (303)
                r2 = d_1470_v21_
                d_1478_v29_: _dafny.Array
                nw209_ = _dafny.Array(False, 26)
                d_1478_v29_ = nw209_
                d_1479_v30_: _dafny.Array
                nw210_ = _dafny.Array(None, 17)
                nw210_[int(0)] = d_1478_v29_
                nw210_[int(1)] = d_1478_v29_
                nw210_[int(2)] = d_1478_v29_
                nw210_[int(3)] = d_1478_v29_
                nw210_[int(4)] = d_1478_v29_
                nw210_[int(5)] = d_1478_v29_
                nw210_[int(6)] = d_1478_v29_
                nw210_[int(7)] = d_1478_v29_
                nw210_[int(8)] = d_1478_v29_
                nw210_[int(9)] = d_1478_v29_
                nw210_[int(10)] = d_1478_v29_
                nw210_[int(11)] = d_1478_v29_
                nw210_[int(12)] = d_1478_v29_
                nw210_[int(13)] = d_1478_v29_
                nw210_[int(14)] = d_1478_v29_
                nw210_[int(15)] = d_1478_v29_
                nw210_[int(16)] = d_1478_v29_
                d_1479_v30_ = nw210_
                index220_ = default__.safeIndex(669, (d_1479_v30_).length(0))
                (d_1479_v30_)[index220_] = d_1478_v29_
                index221_ = default__.safeIndex(669, (d_1479_v30_).length(0))
                (d_1479_v30_)[index221_] = d_1478_v29_
            d_1480_v31_: _dafny.Seq
            d_1480_v31_ = _dafny.SeqWithoutIsStrInference([p0])
            d_1481_v32_: _dafny.Seq
            d_1481_v32_ = _dafny.SeqWithoutIsStrInference([p1])
            d_1482_v33_: D0
            d_1482_v33_ = D0_DC1(not(p0), d_1481_v32_, p1)
            d_1483_v34_: D0
            d_1483_v34_ = D0_DC3(d_1482_v33_)
            d_1484_v35_: _dafny.Map
            d_1484_v35_ = _dafny.Map({(d_1480_v31_)[default__.safeIndex(d_1470_v21_, len(d_1480_v31_))]: d_1483_v34_})
            d_1485_v36_: D2
            d_1485_v36_ = D2_DC7(False, d_1484_v35_, d_1470_v21_, p0)
            d_1486_v37_: str
            d_1486_v37_ = _dafny.CodePoint('u')
            d_1487_v38_: _dafny.MultiSet
            d_1487_v38_ = _dafny.MultiSet([d_1486_v37_])
            d_1488_v39_: _dafny.MultiSet
            d_1488_v39_ = _dafny.MultiSet([_dafny.MultiSet([not(p0)])])
            d_1489_v40_: _dafny.Seq
            d_1489_v40_ = _dafny.SeqWithoutIsStrInference([d_1470_v21_])
            d_1490_v41_: _dafny.Map
            d_1490_v41_ = _dafny.Map({p0: False})
            d_1491_v42_: _dafny.Seq
            d_1491_v42_ = _dafny.SeqWithoutIsStrInference([d_1490_v41_])
            d_1492_v43_: _dafny.Array
            nw211_ = _dafny.Array(None, 27)
            nw211_[int(0)] = d_1470_v21_
            nw211_[int(1)] = (d_1470_v21_) * (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_1493_i1_ in range(default__.abs(819))])))
            nw211_[int(2)] = (0) - (default__.fm22(d_1485_v36_, d_1483_v34_, p0, globalState))
            nw211_[int(3)] = d_1470_v21_
            nw211_[int(4)] = d_1470_v21_
            nw211_[int(5)] = ((d_1487_v38_)[d_1486_v37_] if (d_1486_v37_) in (d_1487_v38_) else d_1470_v21_)
            nw211_[int(6)] = d_1470_v21_
            nw211_[int(7)] = d_1470_v21_
            nw211_[int(8)] = d_1470_v21_
            nw211_[int(9)] = ((d_1488_v39_)[d_1469_cf23_] if (d_1469_cf23_) in (d_1488_v39_) else d_1470_v21_)
            nw211_[int(10)] = d_1470_v21_
            nw211_[int(11)] = d_1470_v21_
            nw211_[int(12)] = d_1470_v21_
            nw211_[int(13)] = (d_1489_v40_)[default__.safeIndex(d_1470_v21_, len(d_1489_v40_))]
            nw211_[int(14)] = (d_1470_v21_ if p0 else d_1470_v21_)
            nw211_[int(15)] = 542
            nw211_[int(16)] = len((p1) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "da"))))
            nw211_[int(17)] = (0) - ((d_1470_v21_) + (471))
            nw211_[int(18)] = d_1470_v21_
            nw211_[int(19)] = (d_1470_v21_ if p0 else d_1470_v21_)
            nw211_[int(20)] = d_1470_v21_
            nw211_[int(21)] = default__.safeDivisionInt(len(d_1491_v42_), (0) - (d_1470_v21_))
            nw211_[int(22)] = (0) - ((d_1470_v21_) + (d_1470_v21_))
            nw211_[int(23)] = d_1470_v21_
            nw211_[int(24)] = (d_1470_v21_) + (d_1470_v21_)
            nw211_[int(25)] = d_1470_v21_
            nw211_[int(26)] = default__.safeDivisionInt(d_1470_v21_, (0) - (d_1470_v21_))
            d_1492_v43_ = nw211_
            d_1494_v44_: T0
            nw212_ = C2()
            nw212_.ctor__()
            d_1494_v44_ = nw212_
            d_1495_v45_: _dafny.Array
            nw213_ = _dafny.Array(None, 29)
            d_1495_v45_ = nw213_
            d_1496_v46_: D0
            d_1496_v46_ = D0_DC1(p0, d_1481_v32_, p1)
            d_1497_v47_: _dafny.Seq
            d_1497_v47_ = _dafny.SeqWithoutIsStrInference([d_1496_v46_])
            d_1498_v48_: C0
            nw214_ = C0()
            nw214_.ctor__((_dafny.MultiSet(d_1497_v47_)).set(d_1496_v46_, default__.abs(d_1470_v21_)), p0)
            d_1498_v48_ = nw214_
            index222_ = default__.safeIndex(648, (d_1495_v45_).length(0))
            (d_1495_v45_)[index222_] = d_1498_v48_
            d_1499_v49_: _dafny.Set
            d_1499_v49_ = _dafny.Set({(d_1498_v48_).f19})
            index223_ = default__.safeIndex(648, (d_1495_v45_).length(0))
            rhs249_ = default__.fm19((d_1470_v21_ if p0 else len(d_1499_v49_)), (d_1498_v48_).f19, globalState)
            rhs250_ = (D13_DC33(d_1492_v43_)).cf51
            rhs251_ = d_1494_v44_
            rhs252_ = d_1498_v48_
            rhs253_ = p0
            lhs163_ = globalState
            lhs164_ = d_1495_v45_
            lhs165_ = default__.safeIndex(648, (d_1495_v45_).length(0))
            lhs166_ = globalState
            lhs163_.f3 = rhs249_
            d_1492_v43_ = rhs250_
            d_1494_v44_ = rhs251_
            lhs164_[lhs165_] = rhs252_
            lhs166_.f5 = rhs253_
        elif True:
            d_1500___mcc_h2_ = source23_.cf25
            d_1501_cf25_ = d_1500___mcc_h2_
            d_1502_v50_: int
            d_1502_v50_ = 377
            d_1503_v51_: _dafny.Set
            d_1503_v51_ = _dafny.Set({p0, p0})
            d_1504_v52_: _dafny.Map
            d_1504_v52_ = _dafny.Map({p0: _dafny.SeqWithoutIsStrInference([p0])})
            d_1505_v53_: _dafny.Map
            d_1505_v53_ = _dafny.Map({p0: d_1502_v50_})
            d_1506_v54_: _dafny.Set
            d_1506_v54_ = _dafny.Set({_dafny.Map({p0: d_1502_v50_}), d_1505_v53_, d_1505_v53_, d_1505_v53_})
            d_1507_v55_: _dafny.Map
            d_1507_v55_ = _dafny.Map({d_1502_v50_: p0})
            d_1508_v56_: _dafny.Seq
            d_1508_v56_ = _dafny.SeqWithoutIsStrInference([d_1502_v50_, d_1502_v50_])
            d_1509_v57_: _dafny.Array
            nw215_ = _dafny.Array(None, 15)
            nw215_[int(0)] = (d_1502_v50_) - (d_1502_v50_)
            nw215_[int(1)] = d_1502_v50_
            nw215_[int(2)] = (d_1502_v50_) * (d_1502_v50_)
            nw215_[int(3)] = d_1502_v50_
            nw215_[int(4)] = len((d_1503_v51_) - (default__.fm33(p0, p0, not(p0), globalState)))
            nw215_[int(5)] = len(d_1504_v52_)
            nw215_[int(6)] = (0) - (len(p1))
            nw215_[int(7)] = default__.fm11(p0, globalState)
            nw215_[int(8)] = len((p1) + (p1))
            nw215_[int(9)] = default__.safeDivisionInt(d_1502_v50_, d_1502_v50_)
            nw215_[int(10)] = (124) * (d_1502_v50_)
            nw215_[int(11)] = len((d_1506_v54_) | (d_1506_v54_))
            nw215_[int(12)] = len((d_1507_v55_) | (d_1507_v55_))
            nw215_[int(13)] = len(_dafny.SeqWithoutIsStrInference([d_1508_v56_, d_1508_v56_, d_1508_v56_, d_1508_v56_]))
            nw215_[int(14)] = d_1502_v50_
            d_1509_v57_ = nw215_
            d_1509_v57_ = d_1509_v57_
            r3 = p0
            d_1502_v50_ = (0) - (d_1502_v50_)
            d_1510_v58_: _dafny.Seq
            d_1510_v58_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vcvnthiyb"))
            d_1510_v58_ = p1
        if p0:
            d_1511_v59_: _dafny.Array
            nw216_ = _dafny.Array(False, 3)
            d_1511_v59_ = nw216_
            d_1512_v60_: _dafny.Set
            d_1512_v60_ = _dafny.Set({d_1511_v59_, d_1511_v59_, d_1511_v59_, d_1511_v59_, d_1511_v59_})
            d_1512_v60_ = d_1512_v60_
            index224_ = default__.safeIndex(683, (d_1511_v59_).length(0))
            (d_1511_v59_)[index224_] = False
            d_1513_v61_: int
            d_1513_v61_ = 330
            d_1514_v62_: _dafny.Map
            d_1514_v62_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_1513_v61_ for d_1515_i2_ in range(default__.abs(-713))])): d_1513_v61_})
            d_1516_v63_: _dafny.Array
            nw217_ = _dafny.Array(None, 24)
            nw217_[int(0)] = (0) - (d_1513_v61_)
            nw217_[int(1)] = d_1513_v61_
            nw217_[int(2)] = d_1513_v61_
            nw217_[int(3)] = d_1513_v61_
            nw217_[int(4)] = (0) - (d_1513_v61_)
            nw217_[int(5)] = d_1513_v61_
            nw217_[int(6)] = d_1513_v61_
            nw217_[int(7)] = d_1513_v61_
            nw217_[int(8)] = 440
            nw217_[int(9)] = ((d_1514_v62_)[d_1513_v61_] if (d_1513_v61_) in (d_1514_v62_) else d_1513_v61_)
            nw217_[int(10)] = d_1513_v61_
            nw217_[int(11)] = d_1513_v61_
            nw217_[int(12)] = d_1513_v61_
            nw217_[int(13)] = d_1513_v61_
            nw217_[int(14)] = d_1513_v61_
            nw217_[int(15)] = 688
            nw217_[int(16)] = d_1513_v61_
            nw217_[int(17)] = d_1513_v61_
            nw217_[int(18)] = d_1513_v61_
            nw217_[int(19)] = default__.fm11(p0, globalState)
            nw217_[int(20)] = ((d_1444_v0_)[p0] if (p0) in (d_1444_v0_) else d_1513_v61_)
            nw217_[int(21)] = d_1513_v61_
            nw217_[int(22)] = d_1513_v61_
            nw217_[int(23)] = len(d_1512_v60_)
            d_1516_v63_ = nw217_
            d_1517_v64_: _dafny.Seq
            d_1517_v64_ = _dafny.SeqWithoutIsStrInference([d_1516_v63_])
            d_1518_v65_: _dafny.Map
            d_1518_v65_ = _dafny.Map({p0: p0})
            d_1519_v66_: _dafny.Seq
            d_1519_v66_ = _dafny.SeqWithoutIsStrInference([p0, p0, p0])
            d_1520_v67_: _dafny.Array
            nw218_ = _dafny.Array(None, 14)
            nw218_[int(0)] = (d_1513_v61_) + (233)
            nw218_[int(1)] = (d_1513_v61_) * (len(d_1517_v64_))
            nw218_[int(2)] = d_1513_v61_
            nw218_[int(3)] = d_1513_v61_
            nw218_[int(4)] = d_1513_v61_
            nw218_[int(5)] = d_1513_v61_
            nw218_[int(6)] = d_1513_v61_
            nw218_[int(7)] = d_1513_v61_
            nw218_[int(8)] = d_1513_v61_
            nw218_[int(9)] = d_1513_v61_
            nw218_[int(10)] = len(d_1518_v65_)
            nw218_[int(11)] = (d_1513_v61_) + (d_1513_v61_)
            nw218_[int(12)] = (d_1513_v61_) * (len(d_1519_v66_))
            nw218_[int(13)] = d_1513_v61_
            d_1520_v67_ = nw218_
            index225_ = default__.safeIndex(756, (d_1520_v67_).length(0))
            (d_1520_v67_)[index225_] = d_1513_v61_
            d_1521_v68_: _dafny.Seq
            d_1521_v68_ = _dafny.SeqWithoutIsStrInference([d_1513_v61_, d_1513_v61_, d_1513_v61_, d_1513_v61_])
            d_1522_v69_: C5
            nw219_ = C5()
            nw219_.ctor__()
            d_1522_v69_ = nw219_
            d_1523_v70_: _dafny.Map
            d_1523_v70_ = _dafny.Map({p0: d_1522_v69_})
            d_1524_v71_: _dafny.Seq
            d_1524_v71_ = _dafny.SeqWithoutIsStrInference([(d_1523_v70_).set(False, d_1522_v69_), d_1523_v70_, d_1523_v70_, d_1523_v70_])
            d_1525_v72_: _dafny.Seq
            d_1525_v72_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(D4_DC12(725)).cf18 for d_1526_i3_ in range(default__.abs(679))]), (d_1521_v68_).set(default__.safeIndex(d_1513_v61_, len(d_1521_v68_)), 173), d_1521_v68_, (d_1521_v68_ if (d_1519_v66_)[default__.safeIndex(len(d_1524_v71_), len(d_1519_v66_))] else d_1521_v68_), (_dafny.SeqWithoutIsStrInference([len(d_1514_v62_) for d_1527_i4_ in range(default__.abs(178))])) + (default__.fm30(globalState))])
            d_1528_v73_: T0
            nw220_ = C2()
            nw220_.ctor__()
            d_1528_v73_ = nw220_
            d_1529_v74_: _dafny.Map
            d_1529_v74_ = _dafny.Map({d_1528_v73_: d_1521_v68_})
            d_1530_v75_: _dafny.Map
            d_1530_v75_ = _dafny.Map({d_1520_v67_: len(d_1529_v74_)})
            d_1531_v76_: _dafny.MultiSet
            d_1531_v76_ = _dafny.MultiSet([d_1513_v61_])
            index226_ = default__.safeIndex(683, (d_1511_v59_).length(0))
            index227_ = default__.safeIndex(756, (d_1520_v67_).length(0))
            rhs254_ = d_1513_v61_
            rhs255_ = p0
            rhs256_ = ((d_1530_v75_)[d_1520_v67_] if (d_1520_v67_) in (d_1530_v75_) else ((d_1531_v76_).cardinality) * (623))
            rhs257_ = len(d_1518_v65_)
            rhs258_ = _dafny.SeqWithoutIsStrInference([(d_1521_v68_) + (d_1521_v68_) for d_1532_i5_ in range(default__.abs(868))])
            lhs167_ = d_1511_v59_
            lhs168_ = default__.safeIndex(683, (d_1511_v59_).length(0))
            lhs169_ = d_1520_v67_
            lhs170_ = default__.safeIndex(756, (d_1520_v67_).length(0))
            r2 = rhs254_
            lhs167_[lhs168_] = rhs255_
            r2 = rhs256_
            lhs169_[lhs170_] = rhs257_
            d_1525_v72_ = rhs258_
            index228_ = default__.safeIndex(683, (d_1511_v59_).length(0))
            (d_1511_v59_)[index228_] = not(p0)
            if not(False):
                index229_ = default__.safeIndex(683, (d_1511_v59_).length(0))
                (d_1511_v59_)[index229_] = (d_1519_v66_)[default__.safeIndex(d_1513_v61_, len(d_1519_v66_))]
                d_1513_v61_ = 653
                r1 = p0
                d_1533_v77_: _dafny.Array
                nw221_ = _dafny.Array(_dafny.Seq({}), 17)
                d_1533_v77_ = nw221_
                d_1534_v78_: _dafny.Seq
                d_1534_v78_ = _dafny.SeqWithoutIsStrInference([d_1533_v77_, d_1533_v77_, d_1533_v77_])
                d_1535_v79_: _dafny.Array
                nw222_ = _dafny.Array(None, 27)
                nw222_[int(0)] = d_1533_v77_
                nw222_[int(1)] = d_1533_v77_
                nw222_[int(2)] = d_1533_v77_
                nw222_[int(3)] = d_1533_v77_
                nw222_[int(4)] = d_1533_v77_
                nw222_[int(5)] = d_1533_v77_
                nw222_[int(6)] = d_1533_v77_
                nw222_[int(7)] = d_1533_v77_
                nw222_[int(8)] = d_1533_v77_
                nw222_[int(9)] = d_1533_v77_
                nw222_[int(10)] = d_1533_v77_
                nw222_[int(11)] = d_1533_v77_
                nw222_[int(12)] = d_1533_v77_
                nw222_[int(13)] = d_1533_v77_
                nw222_[int(14)] = d_1533_v77_
                nw222_[int(15)] = d_1533_v77_
                nw222_[int(16)] = d_1533_v77_
                nw222_[int(17)] = d_1533_v77_
                nw222_[int(18)] = d_1533_v77_
                nw222_[int(19)] = (d_1534_v78_)[default__.safeIndex((d_1520_v67_)[default__.safeIndex(756, (d_1520_v67_).length(0))], len(d_1534_v78_))]
                nw222_[int(20)] = d_1533_v77_
                nw222_[int(21)] = d_1533_v77_
                nw222_[int(22)] = d_1533_v77_
                nw222_[int(23)] = d_1533_v77_
                nw222_[int(24)] = d_1533_v77_
                nw222_[int(25)] = d_1533_v77_
                nw222_[int(26)] = d_1533_v77_
                d_1535_v79_ = nw222_
                index230_ = default__.safeIndex(493, (d_1535_v79_).length(0))
                (d_1535_v79_)[index230_] = d_1533_v77_
                index231_ = default__.safeIndex(493, (d_1535_v79_).length(0))
                index232_ = default__.safeIndex(683, (d_1511_v59_).length(0))
                rhs259_ = d_1533_v77_
                rhs260_ = (((d_1444_v0_).set(p0, default__.abs(d_1513_v61_))).issubset(d_1444_v0_) if not(p0) else p0)
                lhs171_ = d_1535_v79_
                lhs172_ = default__.safeIndex(493, (d_1535_v79_).length(0))
                lhs173_ = d_1511_v59_
                lhs174_ = default__.safeIndex(683, (d_1511_v59_).length(0))
                lhs171_[lhs172_] = rhs259_
                lhs173_[lhs174_] = rhs260_
                r2 = d_1513_v61_
            elif True:
                d_1536_v80_: _dafny.Set
                d_1536_v80_ = _dafny.Set({p0})
                r0 = ((d_1536_v80_).intersection(default__.fm33(p0, (d_1511_v59_)[default__.safeIndex(683, (d_1511_v59_).length(0))], (d_1511_v59_)[default__.safeIndex(683, (d_1511_v59_).length(0))], globalState))) | (d_1536_v80_)
                d_1537_v81_: D0
                d_1537_v81_ = D0_DC0(len(d_1519_v66_))
                d_1538_v82_: _dafny.Array
                nw223_ = _dafny.Array(None, 4)
                nw223_[int(0)] = d_1537_v81_
                nw223_[int(1)] = d_1537_v81_
                nw223_[int(2)] = D0_DC0(len(p1))
                nw223_[int(3)] = d_1537_v81_
                d_1538_v82_ = nw223_
                d_1539_v83_: D7
                d_1539_v83_ = D7_DC20(p1, default__.fm2(_dafny.CodePoint('u'), default__.fm28(globalState), (d_1520_v67_)[default__.safeIndex(756, (d_1520_v67_).length(0))], globalState))
                d_1540_v84_: _dafny.Seq
                d_1540_v84_ = _dafny.SeqWithoutIsStrInference([p1])
                d_1541_v85_: _dafny.Map
                d_1541_v85_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bjx")): (_dafny.SeqWithoutIsStrInference([p1 for d_1542_i6_ in range(default__.abs(416))])) + (d_1540_v84_)})
                d_1543_v86_: _dafny.Set
                d_1543_v86_ = _dafny.Set({d_1513_v61_, d_1513_v61_})
                d_1544_v87_: _dafny.Map
                d_1544_v87_ = _dafny.Map({p1: d_1543_v86_})
                d_1545_v88_: D4
                d_1545_v88_ = D4_DC10(((d_1544_v87_)[p1] if (p1) in (d_1544_v87_) else d_1543_v86_))
                index233_ = default__.safeIndex(756, (d_1520_v67_).length(0))
                rhs261_ = len(((d_1541_v85_)[p1] if (p1) in (d_1541_v85_) else d_1540_v84_))
                rhs262_ = d_1538_v82_
                rhs263_ = (d_1513_v61_) * ((0) - (d_1513_v61_))
                rhs264_ = default__.fm40(d_1513_v61_, d_1545_v88_, globalState)
                rhs265_ = (d_1513_v61_) + (d_1513_v61_)
                lhs175_ = d_1520_v67_
                lhs176_ = default__.safeIndex(756, (d_1520_v67_).length(0))
                lhs175_[lhs176_] = rhs261_
                d_1538_v82_ = rhs262_
                d_1513_v61_ = rhs263_
                d_1539_v83_ = rhs264_
                r2 = rhs265_
                d_1546_v89_: C3
                nw224_ = C3()
                nw224_.ctor__((d_1513_v61_) <= (d_1513_v61_), (d_1520_v67_)[default__.safeIndex(756, (d_1520_v67_).length(0))])
                d_1546_v89_ = nw224_
                d_1547_v90_: _dafny.Set
                d_1547_v90_ = _dafny.Set({d_1539_v83_})
                d_1548_v91_: _dafny.Seq
                d_1548_v91_ = _dafny.SeqWithoutIsStrInference([D7_DC20(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_1549_i7_ in range(default__.abs(965))]), p0)])
                index234_ = default__.safeIndex(683, (d_1511_v59_).length(0))
                def iife108_():
                    coll42_ = _dafny.Set()
                    compr_42_: D7
                    for compr_42_ in (d_1548_v91_).Elements:
                        d_1550_v92_: D7 = compr_42_
                        if (d_1550_v92_) in (d_1548_v91_):
                            coll42_ = coll42_.union(_dafny.Set([d_1550_v92_]))
                    return _dafny.Set(coll42_)
                (d_1511_v59_)[index234_] = ((d_1547_v90_).isdisjoint(iife108_()
                ) if (d_1511_v59_)[default__.safeIndex(683, (d_1511_v59_).length(0))] else p0)
                d_1551_v93_: _dafny.Array
                nw225_ = _dafny.Array(_dafny.Array(None, 0), 11)
                d_1551_v93_ = nw225_
                d_1551_v93_ = d_1551_v93_
            d_1552_v94_: D13
            d_1552_v94_ = D13_DC33(d_1516_v63_)
            d_1553_v95_: _dafny.Array
            nw226_ = _dafny.Array(None, 11)
            nw226_[int(0)] = d_1520_v67_
            nw226_[int(1)] = d_1520_v67_
            nw226_[int(2)] = d_1520_v67_
            nw226_[int(3)] = (d_1520_v67_ if p0 else d_1516_v63_)
            nw226_[int(4)] = (d_1517_v64_)[default__.safeIndex((0) - (-168), len(d_1517_v64_))]
            nw226_[int(5)] = (d_1552_v94_).cf51
            nw226_[int(6)] = d_1516_v63_
            nw226_[int(7)] = d_1516_v63_
            nw226_[int(8)] = d_1516_v63_
            nw226_[int(9)] = d_1516_v63_
            nw226_[int(10)] = d_1516_v63_
            d_1553_v95_ = nw226_
            index235_ = default__.safeIndex(800, (d_1553_v95_).length(0))
            (d_1553_v95_)[index235_] = d_1520_v67_
            d_1554_v96_: D5
            d_1554_v96_ = D5_DC13(d_1511_v59_)
            index236_ = default__.safeIndex(800, (d_1553_v95_).length(0))
            index237_ = default__.safeIndex(683, (d_1511_v59_).length(0))
            rhs266_ = not(p0)
            rhs267_ = d_1520_v67_
            rhs268_ = p0
            rhs269_ = (p0 if p0 else p0)
            rhs270_ = (d_1554_v96_).cf19
            lhs177_ = globalState
            lhs178_ = d_1553_v95_
            lhs179_ = default__.safeIndex(800, (d_1553_v95_).length(0))
            lhs180_ = globalState
            lhs181_ = d_1511_v59_
            lhs182_ = default__.safeIndex(683, (d_1511_v59_).length(0))
            lhs177_.f13 = rhs266_
            lhs178_[lhs179_] = rhs267_
            lhs180_.f9 = rhs268_
            lhs181_[lhs182_] = rhs269_
            d_1511_v59_ = rhs270_
        elif True:
            d_1555_v97_: _dafny.Set
            d_1555_v97_ = _dafny.Set({not(False), p0})
            r0 = (d_1555_v97_).intersection(_dafny.Set({p0, p0, p0}))
            d_1556_v98_: _dafny.Array
            nw227_ = _dafny.Array(False, 18)
            d_1556_v98_ = nw227_
            index238_ = default__.safeIndex(668, (d_1556_v98_).length(0))
            (d_1556_v98_)[index238_] = True
            index239_ = default__.safeIndex(668, (d_1556_v98_).length(0))
            (d_1556_v98_)[index239_] = p0
            d_1557_v99_: int
            d_1557_v99_ = 564
            r3 = (836) <= (default__.safeDivisionInt(d_1557_v99_, d_1557_v99_))
            (globalState).f9 = not(p0)
            d_1558_v100_: _dafny.Seq
            d_1558_v100_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bspwfx"))
            d_1558_v100_ = (p1) + ((p1) + (d_1558_v100_))
        (globalState).f9 = default__.fm19(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gb"))), not(not(p0)), globalState)
        if p0:
            d_1559_v101_: str
            d_1559_v101_ = _dafny.CodePoint('s')
            d_1559_v101_ = d_1559_v101_
            d_1560_v102_: _dafny.Seq
            d_1560_v102_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xmcve"))
            d_1560_v102_ = _dafny.SeqWithoutIsStrInference([d_1559_v101_ for d_1561_i8_ in range(default__.abs(-191))])
            d_1562_v103_: _dafny.Seq
            d_1562_v103_ = _dafny.SeqWithoutIsStrInference([p0])
            d_1563_v104_: _dafny.Set
            d_1563_v104_ = _dafny.Set({d_1562_v103_, d_1562_v103_})
            r2 = len((d_1563_v104_) | (d_1563_v104_))
            d_1564_v105_: int
            d_1564_v105_ = 252
            r2 = (0) - (d_1564_v105_)
            d_1565_v106_: D9
            d_1565_v106_ = D9_DC26(p0, d_1559_v101_, p0)
            d_1566_v107_: _dafny.Map
            d_1566_v107_ = _dafny.Map({p0: p0})
            d_1567_v108_: _dafny.Map
            d_1567_v108_ = _dafny.Map({d_1559_v101_: d_1566_v107_})
            pat_let_tv29_ = p0
            pat_let_tv30_ = p0
            pat_let_tv31_ = d_1567_v108_
            pat_let_tv32_ = p0
            pat_let_tv33_ = d_1564_v105_
            pat_let_tv34_ = globalState
            d_1568_v109_: _dafny.Array
            nw228_ = _dafny.Array(None, 9)
            nw228_[int(0)] = _dafny.CodePoint('b')
            nw228_[int(1)] = d_1559_v101_
            nw228_[int(2)] = _dafny.CodePoint('u')
            nw228_[int(3)] = d_1559_v101_
            nw228_[int(4)] = d_1559_v101_
            def iife109_(_pat_let33_0):
                def iife110_(d_1569_dt__update__tmp_h0_):
                    def iife111_(_pat_let34_0):
                        def iife112_(d_1570_dt__update_hcf41_h0_):
                            def iife113_(_pat_let35_0):
                                def iife114_(d_1571_dt__update_hcf40_h0_):
                                    return D9_DC26((d_1569_dt__update__tmp_h0_).cf39, d_1571_dt__update_hcf40_h0_, d_1570_dt__update_hcf41_h0_)
                                return iife114_(_pat_let35_0)
                            return iife113_(default__.fm23(pat_let_tv30_, pat_let_tv31_, pat_let_tv32_, pat_let_tv33_, pat_let_tv34_))
                        return iife112_(_pat_let34_0)
                    return iife111_(pat_let_tv29_)
                return iife110_(_pat_let33_0)
            nw228_[int(5)] = (iife109_(d_1565_v106_)).cf40
            nw228_[int(6)] = d_1559_v101_
            nw228_[int(7)] = _dafny.CodePoint('s')
            nw228_[int(8)] = d_1559_v101_
            d_1568_v109_ = nw228_
            index240_ = default__.safeIndex(105, (d_1568_v109_).length(0))
            (d_1568_v109_)[index240_] = _dafny.CodePoint('w')
            d_1572_v110_: _dafny.Set
            d_1572_v110_ = _dafny.Set({True})
            d_1573_v111_: D7
            d_1573_v111_ = D7_DC20(d_1560_v102_, p0)
            index241_ = default__.safeIndex(105, (d_1568_v109_).length(0))
            rhs271_ = (default__.fm41(p0, True, globalState)).cf39
            rhs272_ = default__.fm23((d_1572_v110_).ispropersubset(_dafny.Set({p0, p0})), d_1567_v108_, (d_1573_v111_).cf32, d_1564_v105_, globalState)
            lhs183_ = globalState
            lhs184_ = d_1568_v109_
            lhs185_ = default__.safeIndex(105, (d_1568_v109_).length(0))
            lhs183_.f5 = rhs271_
            lhs184_[lhs185_] = rhs272_
        elif True:
            d_1574_v112_: int
            d_1574_v112_ = 579
            r2 = d_1574_v112_
            if p0:
                d_1575_v113_: _dafny.Array
                nw229_ = _dafny.Array(int(0), 11)
                d_1575_v113_ = nw229_
                d_1576_v114_: _dafny.Map
                d_1576_v114_ = _dafny.Map({d_1575_v113_: p0})
                d_1576_v114_ = d_1576_v114_
                d_1577_v115_: _dafny.Array
                nw230_ = _dafny.Array(False, 27)
                d_1577_v115_ = nw230_
                d_1578_v116_: _dafny.Seq
                d_1578_v116_ = _dafny.SeqWithoutIsStrInference([p0])
                index242_ = default__.safeIndex(694, (d_1577_v115_).length(0))
                (d_1577_v115_)[index242_] = (p0) not in (d_1578_v116_)
                index243_ = default__.safeIndex(694, (d_1577_v115_).length(0))
                rhs273_ = p0
                rhs274_ = d_1574_v112_
                rhs275_ = default__.safeDivisionInt(len(p1), d_1574_v112_)
                lhs186_ = d_1577_v115_
                lhs187_ = default__.safeIndex(694, (d_1577_v115_).length(0))
                lhs186_[lhs187_] = rhs273_
                r2 = rhs274_
                r2 = rhs275_
                d_1579_v117_: _dafny.Array
                nw231_ = _dafny.Array(None, 18)
                d_1579_v117_ = nw231_
                d_1580_v118_: C2
                nw232_ = C2()
                nw232_.ctor__()
                d_1580_v118_ = nw232_
                index244_ = default__.safeIndex(697, (d_1579_v117_).length(0))
                (d_1579_v117_)[index244_] = d_1580_v118_
                index245_ = default__.safeIndex(697, (d_1579_v117_).length(0))
                (d_1579_v117_)[index245_] = d_1580_v118_
                d_1581_v119_: _dafny.Map
                d_1581_v119_ = _dafny.Map({not((d_1577_v115_)[default__.safeIndex(694, (d_1577_v115_).length(0))]): d_1575_v113_})
                d_1582_v120_: _dafny.Set
                d_1582_v120_ = _dafny.Set({(d_1574_v112_) - (d_1574_v112_), 38, len(d_1581_v119_)})
                d_1582_v120_ = d_1582_v120_
                (globalState).f9 = (d_1577_v115_)[default__.safeIndex(694, (d_1577_v115_).length(0))]
            elif True:
                d_1583_v121_: str
                d_1583_v121_ = _dafny.CodePoint('x')
                d_1584_v122_: _dafny.Map
                d_1584_v122_ = _dafny.Map({d_1583_v121_: d_1574_v112_})
                d_1585_v123_: C4
                nw233_ = C4()
                nw233_.ctor__(d_1574_v112_, (d_1584_v122_) | (d_1584_v122_))
                d_1585_v123_ = nw233_
                d_1586_v124_: _dafny.Map
                d_1586_v124_ = _dafny.Map({len(p1): d_1574_v112_})
                d_1587_v125_: _dafny.MultiSet
                d_1587_v125_ = _dafny.MultiSet([_dafny.CodePoint('o'), d_1583_v121_])
                d_1588_v126_: _dafny.Array
                nw234_ = _dafny.Array(None, 17)
                nw234_[int(0)] = (d_1585_v123_).f14
                nw234_[int(1)] = len(((self).fm1(d_1586_v124_, globalState)).set(default__.safeIndex(d_1574_v112_, len((self).fm1(d_1586_v124_, globalState))), d_1583_v121_))
                nw234_[int(2)] = (d_1585_v123_).f14
                nw234_[int(3)] = (d_1585_v123_).f14
                nw234_[int(4)] = ((d_1587_v125_).set(d_1583_v121_, default__.abs((d_1585_v123_).f14))).cardinality
                nw234_[int(5)] = (d_1585_v123_).f14
                nw234_[int(6)] = d_1574_v112_
                nw234_[int(7)] = d_1574_v112_
                nw234_[int(8)] = (d_1585_v123_).f14
                nw234_[int(9)] = (d_1585_v123_).f14
                nw234_[int(10)] = -354
                nw234_[int(11)] = (d_1585_v123_).f14
                nw234_[int(12)] = len(_dafny.SeqWithoutIsStrInference([(d_1585_v123_).f14 for d_1589_i9_ in range(default__.abs(148))]))
                nw234_[int(13)] = d_1574_v112_
                nw234_[int(14)] = len(p1)
                nw234_[int(15)] = d_1574_v112_
                nw234_[int(16)] = (d_1585_v123_).f14
                d_1588_v126_ = nw234_
                d_1590_v127_: _dafny.Seq
                d_1590_v127_ = _dafny.SeqWithoutIsStrInference([d_1588_v126_])
                d_1591_v128_: _dafny.Set
                d_1591_v128_ = _dafny.Set({not(p0)})
                d_1592_v129_: _dafny.Map
                d_1592_v129_ = _dafny.Map({_dafny.MultiSet([p0, True]): (d_1591_v128_).isdisjoint(d_1591_v128_)})
                rhs276_ = default__.safeModuloInt((d_1585_v123_).f14, (d_1585_v123_).f14)
                rhs277_ = ((d_1592_v129_)[d_1444_v0_] if (d_1444_v0_) in (d_1592_v129_) else p0)
                rhs278_ = d_1590_v127_
                lhs188_ = globalState
                r2 = rhs276_
                lhs188_.f9 = rhs277_
                d_1590_v127_ = rhs278_
                d_1593_v130_: _dafny.Map
                d_1593_v130_ = _dafny.Map({p0: p0})
                d_1593_v130_ = d_1593_v130_
                d_1594_v131_: _dafny.Seq
                d_1594_v131_ = _dafny.SeqWithoutIsStrInference([p0, p0])
                d_1595_v132_: C0
                nw235_ = C0()
                nw235_.ctor__(default__.fm42(p0, len(d_1594_v131_), (d_1585_v123_).f15, _dafny.SeqWithoutIsStrInference([p0]), globalState), (not(p0)) and (True))
                d_1595_v132_ = nw235_
                d_1595_v132_ = d_1595_v132_
                r1 = p0
            d_1596_v133_: _dafny.Array
            nw236_ = _dafny.Array(_dafny.MultiSet({}), 22)
            d_1596_v133_ = nw236_
            d_1597_v134_: _dafny.Map
            d_1597_v134_ = _dafny.Map({p0: p0})
            d_1598_v135_: _dafny.Seq
            d_1598_v135_ = _dafny.SeqWithoutIsStrInference([d_1597_v134_])
            d_1599_v136_: _dafny.MultiSet
            d_1599_v136_ = _dafny.MultiSet([d_1597_v134_, (d_1598_v135_)[default__.safeIndex(d_1574_v112_, len(d_1598_v135_))]])
            index246_ = default__.safeIndex(321, (d_1596_v133_).length(0))
            (d_1596_v133_)[index246_] = d_1599_v136_
            index247_ = default__.safeIndex(321, (d_1596_v133_).length(0))
            (d_1596_v133_)[index247_] = (d_1599_v136_) | (_dafny.MultiSet([d_1597_v134_, d_1597_v134_]))
            d_1600_v137_: C1
            nw237_ = C1()
            nw237_.ctor__()
            d_1600_v137_ = nw237_
            d_1601_v138_: _dafny.Array
            nw238_ = _dafny.Array(_dafny.Array(None, 0), 9)
            d_1601_v138_ = nw238_
            d_1601_v138_ = d_1601_v138_
        d_1602_v139_: _dafny.Array
        nw239_ = _dafny.Array(False, 10)
        d_1602_v139_ = nw239_
        index248_ = default__.safeIndex(303, (d_1602_v139_).length(0))
        (d_1602_v139_)[index248_] = not(p0)
        d_1603_v140_: _dafny.Map
        d_1603_v140_ = _dafny.Map({p0: p0})
        d_1604_v141_: int
        d_1604_v141_ = 18
        d_1605_v142_: _dafny.Seq
        d_1605_v142_ = _dafny.SeqWithoutIsStrInference([p0, p0, p0, False, True])
        d_1606_v143_: D0
        d_1606_v143_ = D0_DC2(not((d_1605_v142_)[default__.safeIndex(d_1604_v141_, len(d_1605_v142_))]))
        d_1607_v144_: D0
        d_1607_v144_ = D0_DC3(d_1606_v143_)
        index249_ = default__.safeIndex(303, (d_1602_v139_).length(0))
        (d_1602_v139_)[index249_] = ((d_1603_v140_)[(p0 if p0 else p0)] if ((p0 if p0 else p0)) in (d_1603_v140_) else default__.fm2((p1)[default__.safeIndex(d_1604_v141_, len(p1))], d_1607_v144_, d_1604_v141_, globalState))
        if (210) >= (default__.safeDivisionInt(d_1604_v141_, len((d_1605_v142_).set(default__.safeIndex(d_1604_v141_, len(d_1605_v142_)), p0)))):
            d_1608_v145_: _dafny.Seq
            d_1608_v145_ = _dafny.SeqWithoutIsStrInference([d_1602_v139_, d_1602_v139_])
            d_1609_v146_: _dafny.Map
            d_1609_v146_ = _dafny.Map({(d_1602_v139_)[default__.safeIndex(303, (d_1602_v139_).length(0))]: d_1602_v139_})
            d_1610_v147_: _dafny.Array
            nw240_ = _dafny.Array(None, 26)
            nw240_[int(0)] = d_1602_v139_
            nw240_[int(1)] = d_1602_v139_
            nw240_[int(2)] = d_1602_v139_
            nw240_[int(3)] = d_1602_v139_
            nw240_[int(4)] = d_1602_v139_
            nw240_[int(5)] = d_1602_v139_
            nw240_[int(6)] = d_1602_v139_
            nw240_[int(7)] = d_1602_v139_
            nw240_[int(8)] = d_1602_v139_
            nw240_[int(9)] = d_1602_v139_
            nw240_[int(10)] = d_1602_v139_
            nw240_[int(11)] = d_1602_v139_
            nw240_[int(12)] = d_1602_v139_
            nw240_[int(13)] = d_1602_v139_
            nw240_[int(14)] = (d_1608_v145_)[default__.safeIndex(d_1604_v141_, len(d_1608_v145_))]
            nw240_[int(15)] = d_1602_v139_
            nw240_[int(16)] = d_1602_v139_
            nw240_[int(17)] = d_1602_v139_
            nw240_[int(18)] = d_1602_v139_
            nw240_[int(19)] = d_1602_v139_
            nw240_[int(20)] = d_1602_v139_
            nw240_[int(21)] = d_1602_v139_
            nw240_[int(22)] = d_1602_v139_
            nw240_[int(23)] = d_1602_v139_
            nw240_[int(24)] = ((d_1609_v146_)[(d_1602_v139_)[default__.safeIndex(303, (d_1602_v139_).length(0))]] if ((d_1602_v139_)[default__.safeIndex(303, (d_1602_v139_).length(0))]) in (d_1609_v146_) else d_1602_v139_)
            nw240_[int(25)] = d_1602_v139_
            d_1610_v147_ = nw240_
            index250_ = default__.safeIndex(909, (d_1610_v147_).length(0))
            (d_1610_v147_)[index250_] = d_1602_v139_
            index251_ = default__.safeIndex(909, (d_1610_v147_).length(0))
            (d_1610_v147_)[index251_] = (D5_DC13(d_1602_v139_)).cf19
            d_1611_v148_: _dafny.Seq
            d_1611_v148_ = _dafny.SeqWithoutIsStrInference([d_1604_v141_])
            d_1612_v149_: _dafny.Seq
            d_1612_v149_ = _dafny.SeqWithoutIsStrInference([d_1608_v145_])
            d_1613_v150_: _dafny.Array
            nw241_ = _dafny.Array(None, 24)
            nw241_[int(0)] = (d_1608_v145_).set(default__.safeIndex(d_1604_v141_, len(d_1608_v145_)), (d_1610_v147_)[default__.safeIndex(909, (d_1610_v147_).length(0))])
            nw241_[int(1)] = d_1608_v145_
            nw241_[int(2)] = d_1608_v145_
            nw241_[int(3)] = _dafny.SeqWithoutIsStrInference([(d_1610_v147_)[default__.safeIndex(909, (d_1610_v147_).length(0))]])
            nw241_[int(4)] = (d_1608_v145_) + (d_1608_v145_)
            nw241_[int(5)] = d_1608_v145_
            nw241_[int(6)] = d_1608_v145_
            nw241_[int(7)] = d_1608_v145_
            nw241_[int(8)] = (d_1608_v145_).set(default__.safeIndex((0) - (len(d_1611_v148_)), len(d_1608_v145_)), d_1602_v139_)
            nw241_[int(9)] = (d_1608_v145_) + (d_1608_v145_)
            nw241_[int(10)] = (d_1608_v145_) + (d_1608_v145_)
            nw241_[int(11)] = d_1608_v145_
            nw241_[int(12)] = (_dafny.SeqWithoutIsStrInference([(d_1610_v147_)[default__.safeIndex(909, (d_1610_v147_).length(0))]])) + (_dafny.SeqWithoutIsStrInference([(d_1610_v147_)[default__.safeIndex(909, (d_1610_v147_).length(0))], d_1602_v139_]))
            nw241_[int(13)] = d_1608_v145_
            nw241_[int(14)] = _dafny.SeqWithoutIsStrInference([(d_1610_v147_)[default__.safeIndex(909, (d_1610_v147_).length(0))], (d_1610_v147_)[default__.safeIndex(909, (d_1610_v147_).length(0))], d_1602_v139_])
            nw241_[int(15)] = (d_1608_v145_) + (d_1608_v145_)
            nw241_[int(16)] = ((d_1612_v149_)[default__.safeIndex(d_1604_v141_, len(d_1612_v149_))]) + (d_1608_v145_)
            nw241_[int(17)] = _dafny.SeqWithoutIsStrInference([(d_1610_v147_)[default__.safeIndex(909, (d_1610_v147_).length(0))]])
            nw241_[int(18)] = d_1608_v145_
            nw241_[int(19)] = _dafny.SeqWithoutIsStrInference([d_1602_v139_, d_1602_v139_])
            nw241_[int(20)] = (_dafny.SeqWithoutIsStrInference([d_1602_v139_, d_1602_v139_])).set(default__.safeIndex(d_1604_v141_, len(_dafny.SeqWithoutIsStrInference([d_1602_v139_, d_1602_v139_]))), d_1602_v139_)
            nw241_[int(21)] = d_1608_v145_
            nw241_[int(22)] = d_1608_v145_
            nw241_[int(23)] = _dafny.SeqWithoutIsStrInference([d_1602_v139_, d_1602_v139_])
            d_1613_v150_ = nw241_
            index252_ = default__.safeIndex(134, (d_1613_v150_).length(0))
            (d_1613_v150_)[index252_] = d_1608_v145_
            d_1614_v151_: _dafny.Map
            d_1614_v151_ = _dafny.Map({(d_1602_v139_)[default__.safeIndex(303, (d_1602_v139_).length(0))]: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_1615_i11_ in range(default__.abs(791))])})
            d_1616_v152_: _dafny.Seq
            d_1616_v152_ = _dafny.SeqWithoutIsStrInference([p1, p1])
            d_1617_v153_: str
            d_1617_v153_ = _dafny.CodePoint('e')
            d_1618_v154_: D0
            d_1618_v154_ = D0_DC1(False, (d_1616_v152_).set(default__.safeIndex(d_1604_v141_, len(d_1616_v152_)), (p1).set(default__.safeIndex(d_1604_v141_, len(p1)), _dafny.CodePoint('f'))), (p1).set(default__.safeIndex(d_1604_v141_, len(p1)), d_1617_v153_))
            pat_let_tv35_ = d_1602_v139_
            pat_let_tv36_ = d_1602_v139_
            pat_let_tv37_ = d_1616_v152_
            d_1619_v155_: _dafny.Array
            nw242_ = _dafny.Array(None, 26)
            nw242_[int(0)] = _dafny.SeqWithoutIsStrInference([(p1)[default__.safeIndex((0) - (d_1604_v141_), len(p1))] for d_1620_i10_ in range(default__.abs(249))])
            nw242_[int(1)] = ((d_1614_v151_)[p0] if (p0) in (d_1614_v151_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sfabqetxt")))
            nw242_[int(2)] = (p1) + (p1)
            nw242_[int(3)] = p1
            nw242_[int(4)] = p1
            nw242_[int(5)] = p1
            nw242_[int(6)] = p1
            nw242_[int(7)] = p1
            nw242_[int(8)] = p1
            nw242_[int(9)] = (p1) + (p1)
            nw242_[int(10)] = p1
            nw242_[int(11)] = p1
            nw242_[int(12)] = p1
            nw242_[int(13)] = p1
            nw242_[int(14)] = p1
            nw242_[int(15)] = (p1) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uatwgx")))
            def iife115_(_pat_let36_0):
                def iife116_(d_1621_dt__update__tmp_h1_):
                    def iife117_(_pat_let37_0):
                        def iife118_(d_1622_dt__update_hcf1_h0_):
                            def iife119_(_pat_let38_0):
                                def iife120_(d_1623_dt__update_hcf2_h0_):
                                    return D0_DC1(d_1622_dt__update_hcf1_h0_, d_1623_dt__update_hcf2_h0_, (d_1621_dt__update__tmp_h1_).cf3)
                                return iife120_(_pat_let38_0)
                            return iife119_(pat_let_tv37_)
                        return iife118_(_pat_let37_0)
                    return iife117_((pat_let_tv36_)[default__.safeIndex(303, (pat_let_tv35_).length(0))])
                return iife116_(_pat_let36_0)
            nw242_[int(16)] = (p1) + ((iife115_(d_1618_v154_)).cf3)
            nw242_[int(17)] = _dafny.SeqWithoutIsStrInference([d_1617_v153_ for d_1624_i12_ in range(default__.abs(494))])
            nw242_[int(18)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wocuatbi"))
            nw242_[int(19)] = p1
            nw242_[int(20)] = (p1) + (p1)
            nw242_[int(21)] = (D1_DC5(p0, p1, (_dafny.SeqWithoutIsStrInference([(d_1602_v139_)[default__.safeIndex(303, (d_1602_v139_).length(0))], False])).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_1617_v153_ for d_1625_i13_ in range(default__.abs(984))])), len(_dafny.SeqWithoutIsStrInference([(d_1602_v139_)[default__.safeIndex(303, (d_1602_v139_).length(0))], False]))), (d_1605_v142_)[default__.safeIndex(d_1604_v141_, len(d_1605_v142_))]))).cf8
            nw242_[int(22)] = p1
            nw242_[int(23)] = (p1) + (p1)
            nw242_[int(24)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jp"))
            nw242_[int(25)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "scmvf"))
            d_1619_v155_ = nw242_
            index253_ = default__.safeIndex(438, (d_1619_v155_).length(0))
            (d_1619_v155_)[index253_] = p1
            d_1626_v156_: _dafny.MultiSet
            d_1626_v156_ = _dafny.MultiSet([d_1604_v141_])
            index254_ = default__.safeIndex(134, (d_1613_v150_).length(0))
            index255_ = default__.safeIndex(438, (d_1619_v155_).length(0))
            rhs279_ = (_dafny.MultiSet([d_1604_v141_, d_1604_v141_, d_1604_v141_, len(p1)])).isdisjoint(d_1626_v156_)
            rhs280_ = (d_1602_v139_)[default__.safeIndex(303, (d_1602_v139_).length(0))]
            rhs281_ = (d_1608_v145_) + (d_1608_v145_)
            rhs282_ = p1
            lhs189_ = globalState
            lhs190_ = globalState
            lhs191_ = d_1613_v150_
            lhs192_ = default__.safeIndex(134, (d_1613_v150_).length(0))
            lhs193_ = d_1619_v155_
            lhs194_ = default__.safeIndex(438, (d_1619_v155_).length(0))
            lhs189_.f9 = rhs279_
            lhs190_.f13 = rhs280_
            lhs191_[lhs192_] = rhs281_
            lhs193_[lhs194_] = rhs282_
            d_1627_v157_: _dafny.Set
            d_1627_v157_ = _dafny.Set({(d_1602_v139_)[default__.safeIndex(303, (d_1602_v139_).length(0))]})
            d_1628_v158_: _dafny.Seq
            d_1628_v158_ = _dafny.SeqWithoutIsStrInference([d_1627_v157_])
            d_1629_v159_: _dafny.MultiSet
            d_1629_v159_ = _dafny.MultiSet([d_1627_v157_])
            d_1605_v142_ = (d_1605_v142_ if (_dafny.MultiSet(d_1628_v158_)).ispropersubset(d_1629_v159_) else d_1605_v142_)
            d_1630_v160_: _dafny.Array
            nw243_ = _dafny.Array(int(0), 22)
            d_1630_v160_ = nw243_
            d_1630_v160_ = d_1630_v160_
            index256_ = default__.safeIndex(438, (d_1619_v155_).length(0))
            (d_1619_v155_)[index256_] = p1
        elif True:
            r2 = (0) - ((d_1604_v141_) - (d_1604_v141_))
            if (((d_1602_v139_)[default__.safeIndex(303, (d_1602_v139_).length(0))] if (d_1602_v139_)[default__.safeIndex(303, (d_1602_v139_).length(0))] else (d_1602_v139_)[default__.safeIndex(303, (d_1602_v139_).length(0))]) if (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_1631_i14_ in range(default__.abs(540))]))) <= (d_1604_v141_) else (d_1602_v139_)[default__.safeIndex(303, (d_1602_v139_).length(0))]):
                d_1632_v161_: _dafny.Seq
                d_1632_v161_ = _dafny.SeqWithoutIsStrInference([p1])
                d_1633_v162_: D0
                d_1633_v162_ = D0_DC1((d_1602_v139_)[default__.safeIndex(303, (d_1602_v139_).length(0))], d_1632_v161_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_1634_i15_ in range(default__.abs(592))]))
                d_1635_v163_: _dafny.MultiSet
                d_1635_v163_ = _dafny.MultiSet([d_1633_v162_])
                d_1636_v164_: C0
                nw244_ = C0()
                nw244_.ctor__(d_1635_v163_, p0)
                d_1636_v164_ = nw244_
                d_1637_v165_: _dafny.Seq
                d_1637_v165_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hrsxsmked"))
                d_1638_v166_: _dafny.Map
                d_1638_v166_ = _dafny.Map({d_1604_v141_: False})
                d_1639_v167_: _dafny.Map
                d_1639_v167_ = _dafny.Map({d_1604_v141_: len(d_1605_v142_)})
                d_1640_v168_: _dafny.Set
                d_1640_v168_ = _dafny.Set({default__.safeDivisionInt(d_1604_v141_, len(d_1638_v166_)), len(d_1639_v167_), d_1604_v141_})
                d_1641_v169_: D1
                d_1641_v169_ = D1_DC5((d_1602_v139_)[default__.safeIndex(303, (d_1602_v139_).length(0))], d_1637_v165_, d_1605_v142_)
                rhs283_ = (d_1637_v165_) + ((d_1641_v169_).cf8)
                rhs284_ = (d_1602_v139_)[default__.safeIndex(303, (d_1602_v139_).length(0))]
                rhs285_ = d_1640_v168_
                lhs195_ = globalState
                d_1637_v165_ = rhs283_
                lhs195_.f9 = rhs284_
                d_1640_v168_ = rhs285_
                d_1642_v170_: _dafny.Seq
                d_1642_v170_ = _dafny.SeqWithoutIsStrInference([d_1604_v141_, (d_1604_v141_ if (d_1636_v164_).f19 else d_1604_v141_), d_1604_v141_])
                d_1642_v170_ = _dafny.SeqWithoutIsStrInference([len(d_1603_v140_), d_1604_v141_, d_1604_v141_, len(d_1603_v140_)])
                d_1604_v141_ = d_1604_v141_
                d_1643_v171_: T1
                nw245_ = C5()
                nw245_.ctor__()
                d_1643_v171_ = nw245_
                d_1644_v172_: _dafny.Set
                d_1644_v172_ = _dafny.Set({d_1643_v171_})
                d_1645_v173_: D13
                d_1645_v173_ = D13_DC34(default__.safeModuloInt(d_1604_v141_, len(d_1644_v172_)))
                d_1645_v173_ = d_1645_v173_
            elif True:
                d_1646_v174_: C1
                nw246_ = C1()
                nw246_.ctor__()
                d_1646_v174_ = nw246_
                d_1647_v175_: C6
                nw247_ = C6()
                nw247_.ctor__()
                d_1647_v175_ = nw247_
                d_1648_v176_: _dafny.Seq
                d_1648_v176_ = _dafny.SeqWithoutIsStrInference([d_1604_v141_])
                d_1649_v177_: _dafny.Map
                d_1649_v177_ = _dafny.Map({p0: d_1648_v176_})
                d_1650_v178_: _dafny.Seq
                d_1650_v178_ = _dafny.SeqWithoutIsStrInference([d_1649_v177_])
                d_1651_v179_: D11
                d_1651_v179_ = D11_DC29((_dafny.MultiSet([d_1647_v175_])).set(d_1647_v175_, default__.abs(len((d_1650_v178_)[default__.safeIndex(d_1604_v141_, len(d_1650_v178_))]))))
                d_1651_v179_ = d_1651_v179_
                (globalState).f9 = p0
                d_1652_v180_: _dafny.Map
                d_1652_v180_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_1653_i16_ in range(default__.abs(836))]): d_1604_v141_})
                d_1654_v181_: T1
                nw248_ = C3()
                nw248_.ctor__((d_1602_v139_)[default__.safeIndex(303, (d_1602_v139_).length(0))], (d_1444_v0_).cardinality)
                d_1654_v181_ = nw248_
                d_1655_v182_: _dafny.MultiSet
                d_1655_v182_ = _dafny.MultiSet([d_1654_v181_])
                d_1652_v180_ = (d_1652_v180_).set(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_1656_i17_ in range(default__.abs(768))]), (d_1655_v182_).cardinality)
                d_1657_v183_: _dafny.Array
                nw249_ = _dafny.Array(_dafny.CodePoint('D'), 26)
                d_1657_v183_ = nw249_
                d_1658_v184_: str
                d_1658_v184_ = _dafny.CodePoint('b')
                index257_ = default__.safeIndex(563, (d_1657_v183_).length(0))
                (d_1657_v183_)[index257_] = d_1658_v184_
                index258_ = default__.safeIndex(563, (d_1657_v183_).length(0))
                (d_1657_v183_)[index258_] = (p1)[default__.safeIndex((len(d_1605_v142_)) - (d_1604_v141_), len(p1))]
            d_1659_v185_: _dafny.Array
            nw250_ = _dafny.Array(_dafny.CodePoint('D'), 14)
            d_1659_v185_ = nw250_
            index259_ = default__.safeIndex(587, (d_1659_v185_).length(0))
            (d_1659_v185_)[index259_] = _dafny.CodePoint('c')
            d_1660_v186_: _dafny.Map
            d_1660_v186_ = _dafny.Map({(p1)[default__.safeIndex(d_1604_v141_, len(p1))]: d_1603_v140_})
            d_1661_v187_: _dafny.Seq
            d_1661_v187_ = _dafny.SeqWithoutIsStrInference([d_1604_v141_, d_1604_v141_])
            index260_ = default__.safeIndex(303, (d_1602_v139_).length(0))
            index261_ = default__.safeIndex(587, (d_1659_v185_).length(0))
            rhs286_ = (d_1602_v139_)[default__.safeIndex(303, (d_1602_v139_).length(0))]
            rhs287_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wkpov"))) == (p1)
            rhs288_ = default__.fm23((p1) <= (p1), d_1660_v186_, (d_1602_v139_)[default__.safeIndex(303, (d_1602_v139_).length(0))], (d_1661_v187_)[default__.safeIndex(d_1604_v141_, len(d_1661_v187_))], globalState)
            rhs289_ = 986
            rhs290_ = d_1604_v141_
            lhs196_ = globalState
            lhs197_ = d_1602_v139_
            lhs198_ = default__.safeIndex(303, (d_1602_v139_).length(0))
            lhs199_ = d_1659_v185_
            lhs200_ = default__.safeIndex(587, (d_1659_v185_).length(0))
            lhs196_.f3 = rhs286_
            lhs197_[lhs198_] = rhs287_
            lhs199_[lhs200_] = rhs288_
            r2 = rhs289_
            r2 = rhs290_
            (globalState).f5 = True
            d_1662_v188_: _dafny.Set
            d_1662_v188_ = _dafny.Set({d_1444_v0_, d_1444_v0_})
            d_1663_v189_: _dafny.MultiSet
            d_1663_v189_ = _dafny.MultiSet([759])
            d_1664_v190_: _dafny.Map
            d_1664_v190_ = _dafny.Map({(d_1662_v188_).isdisjoint(default__.fm43((d_1602_v139_)[default__.safeIndex(303, (d_1602_v139_).length(0))], globalState)): d_1663_v189_})
            d_1665_v191_: _dafny.Map
            d_1665_v191_ = _dafny.Map({d_1602_v139_: _dafny.SeqWithoutIsStrInference([False, p0])})
            d_1664_v190_ = (d_1664_v190_).set(default__.fm2(default__.fm23((d_1602_v139_)[default__.safeIndex(303, (d_1602_v139_).length(0))], d_1660_v186_, (d_1602_v139_)[default__.safeIndex(303, (d_1602_v139_).length(0))], len(d_1665_v191_), globalState), D0_DC3(d_1606_v143_), d_1604_v141_, globalState), d_1663_v189_)
        d_1666_v192_: _dafny.Set
        d_1666_v192_ = _dafny.Set({default__.fm6(globalState), (d_1602_v139_)[default__.safeIndex(303, (d_1602_v139_).length(0))]})
        r0 = d_1666_v192_
        r1 = not((d_1602_v139_)[default__.safeIndex(303, (d_1602_v139_).length(0))])
        r2 = d_1604_v141_
        r3 = p0
        return r0, r1, r2, r3

