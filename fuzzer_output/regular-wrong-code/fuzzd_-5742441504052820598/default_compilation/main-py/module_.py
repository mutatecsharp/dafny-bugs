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
    def fm0(p0, p1, globalState):
        return False

    @staticmethod
    def fm1(globalState):
        return (default__.safeDivisionInt(720, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))))) + (default__.safeDivisionInt(len(_dafny.Set({False})), len(_dafny.Set({False}))))

    @staticmethod
    def fm2(p0, p1, p2, globalState):
        return (_dafny.SeqWithoutIsStrInference([True])) + (_dafny.SeqWithoutIsStrInference([True, False]))

    @staticmethod
    def fm3(p0, p1, p2, globalState):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: str
            for compr_0_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g'), _dafny.CodePoint('w'), _dafny.CodePoint('i')])).Elements:
                d_0_v0_: str = compr_0_
                if (d_0_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g'), _dafny.CodePoint('w'), _dafny.CodePoint('i')])):
                    coll0_[d_0_v0_] = (D4_DC13(D1_DC4(61, _dafny.SeqWithoutIsStrInference([False, True]), len(_dafny.Map({True: False}))), True, _dafny.SeqWithoutIsStrInference([True]))).cf25
            return _dafny.Map(coll0_)
        def iife1_():
            coll1_ = _dafny.Map()
            compr_1_: str
            for compr_1_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j')])).Elements:
                d_1_v1_: str = compr_1_
                if (d_1_v1_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j')])):
                    coll1_[d_1_v1_] = True
            return _dafny.Map(coll1_)
        return ((iife0_()
         if not(False) else iife1_()
        )) | (_dafny.Map({_dafny.CodePoint('h'): not(not(False))}))

    @staticmethod
    def fm4(globalState):
        return _dafny.Set({_dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o')]))): not(not(True))})})

    @staticmethod
    def fm13(globalState):
        return ((_dafny.Map({True: _dafny.SeqWithoutIsStrInference([not(False), False])})) | (_dafny.Map({True: _dafny.SeqWithoutIsStrInference([False, False, (D5_DC17(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([not(True)])), True)).cf35, False])}))) | ((_dafny.Map({True: _dafny.SeqWithoutIsStrInference([False, False])})) | (_dafny.Map({True: _dafny.SeqWithoutIsStrInference([False])})))

    @staticmethod
    def fm14(p0, p1, p2, p3, globalState):
        source0_ = (D12_DC37(True, True, False) if True else D12_DC37(False, not(True), True))
        if source0_.is_DC37:
            d_2___mcc_h0_ = source0_.cf65
            d_3___mcc_h1_ = source0_.cf66
            d_4___mcc_h2_ = source0_.cf67
            d_5_cf67_ = d_4___mcc_h2_
            d_6_cf66_ = d_3___mcc_h1_
            d_7_cf65_ = d_2___mcc_h0_
            def iife2_():
                coll2_ = _dafny.Set()
                compr_2_: int
                for compr_2_ in (_dafny.SeqWithoutIsStrInference([-911, -848])).Elements:
                    d_8_v0_: int = compr_2_
                    if (d_8_v0_) in (_dafny.SeqWithoutIsStrInference([-911, -848])):
                        coll2_ = coll2_.union(_dafny.Set([(d_8_v0_) * (len(_dafny.SeqWithoutIsStrInference([False])))]))
                return _dafny.Set(coll2_)
            return _dafny.SeqWithoutIsStrInference([len(iife2_()
            )])
        elif source0_.is_DC38:
            d_9___mcc_h3_ = source0_.cf68
            d_10_cf68_ = d_9___mcc_h3_
            return (_dafny.SeqWithoutIsStrInference([-928])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: True})), d_10_cf68_, d_10_cf68_]))
        elif True:
            d_11___mcc_h4_ = source0_.cf64
            d_12_cf64_ = d_11___mcc_h4_
            return (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_13_i0_ in range(default__.abs(689))]))}))])) + (_dafny.SeqWithoutIsStrInference([718]))

    @staticmethod
    def fm16(p0, p1, p2, p3, globalState):
        def iife3_():
            coll3_ = _dafny.Set()
            compr_3_: int
            for compr_3_ in _dafny.IntegerRange(692, -359):
                d_14_v0_: int = compr_3_
                if ((692) <= (d_14_v0_)) and ((d_14_v0_) < (-359)):
                    coll3_ = coll3_.union(_dafny.Set([(d_14_v0_) + (676)]))
            return _dafny.Set(coll3_)
        return (iife3_()
        ) - ((_dafny.Set({160, len(_dafny.Set({False}))})) - (_dafny.Set({(_dafny.MultiSet([True, (D3_DC10(len(_dafny.Set({371, -573, 30, 327, -973})), 173, True, 916, False)).cf19])).cardinality, 889})))

    @staticmethod
    def fm17(p0, p1, globalState):
        source1_ = D25_DC66()
        if source1_.is_DC66:
            if not(not(not(not(False)))):
                return _dafny.Set({True, False, False, not(True)})
            elif True:
                return _dafny.Set({True})
        elif source1_.is_DC65:
            d_15___mcc_h0_ = source1_.cf112
            d_16_cf112_ = d_15___mcc_h0_
            return _dafny.Set({not(True)})
        elif True:
            d_17___mcc_h1_ = source1_.cf113
            d_18_cf113_ = d_17___mcc_h1_
            return (_dafny.Set({False})) | (_dafny.Set({True, True}))

    @staticmethod
    def fm18(p0, p1, p2, globalState):
        return (_dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_19_i0_ in range(default__.abs(352))]): len(_dafny.SeqWithoutIsStrInference([True, not(False)]))})) | (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "txtdxytq")): 825}))

    @staticmethod
    def fm19(p0, p1, globalState):
        return (_dafny.Map({True: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "audp"))})) | (_dafny.Map({not(True): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lpvkko"))}))

    @staticmethod
    def fm20(p0, p1, globalState):
        return D3_DC10(-898, default__.safeModuloInt(301, -715), (len(_dafny.SeqWithoutIsStrInference([513, 745, (0) - (len(_dafny.SeqWithoutIsStrInference([not(not(True))]))), (_dafny.MultiSet([True, False])).cardinality, -721]))) != (len(_dafny.Map({14: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_20_i0_ in range(default__.abs(766))]))}))), default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "daknbvfwa"))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "svirrblg")))), (509) > (924))

    @staticmethod
    def fm21(p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference([(0) - ((0) - (-843))])) + (_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([D17_DC46(True, len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([True]))})), (_dafny.MultiSet([len(_dafny.Set({False, False, True}))])).cardinality, True)])))]))

    @staticmethod
    def fm22(p0, p1, p2, p3, globalState):
        source2_ = D7_DC25(False, True, 506)
        if source2_.is_DC25:
            d_21___mcc_h0_ = source2_.cf45
            d_22___mcc_h1_ = source2_.cf46
            d_23___mcc_h2_ = source2_.cf47
            d_24_cf47_ = d_23___mcc_h2_
            d_25_cf46_ = d_22___mcc_h1_
            d_26_cf45_ = d_21___mcc_h0_
            return D2_DC5(_dafny.MultiSet([d_26_cf45_]))
        elif source2_.is_DC26:
            d_27___mcc_h3_ = source2_.cf48
            d_28___mcc_h4_ = source2_.cf49
            d_29_cf49_ = d_28___mcc_h4_
            d_30_cf48_ = d_27___mcc_h3_
            return D2_DC5(_dafny.MultiSet([False]))
        elif source2_.is_DC27:
            d_31___mcc_h5_ = source2_.cf50
            d_32___mcc_h6_ = source2_.cf51
            d_33___mcc_h7_ = source2_.cf52
            d_34_cf52_ = d_33___mcc_h7_
            d_35_cf51_ = d_32___mcc_h6_
            d_36_cf50_ = d_31___mcc_h5_
            return D2_DC5(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True, False, True, True, True])))
        elif True:
            d_37___mcc_h8_ = source2_.cf44
            d_38_cf44_ = d_37___mcc_h8_
            return D2_DC5(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True, False])))

    @staticmethod
    def fm23(p0, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iexcfynql"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tco")))

    @staticmethod
    def fm24(globalState):
        return ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(0) - (-845), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qhpo"))), -683, len(_dafny.SeqWithoutIsStrInference([not(True), not(True)]))]), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True, False]))])) for d_39_i0_ in range(default__.abs(828))]))])]))) - (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([377 for d_40_i1_ in range(default__.abs(-450))]), _dafny.SeqWithoutIsStrInference([699, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yycq")))])]))) | ((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Map({not(False): _dafny.CodePoint('m')}))) for d_41_i2_ in range(default__.abs(418))])]) if not(False) else _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_42_i3_ in range(default__.abs(381))])), 326, len(_dafny.SeqWithoutIsStrInference([False])), (0) - ((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gpcwyeyi"))])).cardinality), len(_dafny.SeqWithoutIsStrInference([True]))]), _dafny.SeqWithoutIsStrInference([266, 550])])))

    @staticmethod
    def fm25(globalState):
        return ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False]))) - (_dafny.MultiSet([False]))) - ((_dafny.MultiSet([False, False, True, False])) - (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True]))))

    @staticmethod
    def fm26(p0, p1, p2, p3, globalState):
        if (_dafny.Set({False})).isdisjoint(_dafny.Set({(D14_DC42(887, True, _dafny.Set({not(True)}), 841)).cf73, True})):
            return _dafny.Map({True: _dafny.Map({521: False})})
        elif True:
            return _dafny.Map({not(False): _dafny.Map({len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False]))).cardinality])): True})})

    @staticmethod
    def fm27(p0, globalState):
        if True:
            if False:
                return _dafny.CodePoint('h')
            elif True:
                return _dafny.CodePoint('x')
        elif True:
            return _dafny.CodePoint('x')

    @staticmethod
    def fm28(p0, globalState):
        return _dafny.Map({(not(not(not(True))) if False else True): (160) * (269)})

    @staticmethod
    def fm29(p0, globalState):
        return (_dafny.Map({_dafny.CodePoint('y'): _dafny.Map({True: _dafny.CodePoint('s')})})) | (_dafny.Map({_dafny.CodePoint('y'): _dafny.Map({False: _dafny.CodePoint('y')})}))

    @staticmethod
    def fm30(globalState):
        source3_ = D20_DC54(False, True, len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pdadbgu"))): len(_dafny.SeqWithoutIsStrInference([_dafny.Map({869: not(True)}) for d_43_i0_ in range(default__.abs(720))]))})), True)
        if source3_.is_DC54:
            d_44___mcc_h0_ = source3_.cf97
            d_45___mcc_h1_ = source3_.cf98
            d_46___mcc_h2_ = source3_.cf99
            d_47___mcc_h3_ = source3_.cf100
            d_48_cf100_ = d_47___mcc_h3_
            d_49_cf99_ = d_46___mcc_h2_
            d_50_cf98_ = d_45___mcc_h1_
            d_51_cf97_ = d_44___mcc_h0_
            return D12_DC37(d_51_cf97_, True, False)
        elif source3_.is_DC55:
            d_52___mcc_h4_ = source3_.cf101
            d_53___mcc_h5_ = source3_.cf102
            d_54___mcc_h6_ = source3_.cf103
            d_55_cf103_ = d_54___mcc_h6_
            d_56_cf102_ = d_53___mcc_h5_
            d_57_cf101_ = d_52___mcc_h4_
            return D12_DC37(True, not(not(True)), True)
        elif source3_.is_DC53:
            d_58___mcc_h7_ = source3_.cf96
            d_59_cf96_ = d_58___mcc_h7_
            return D12_DC37(True, False, False)
        elif True:
            d_60___mcc_h8_ = source3_.cf104
            d_61_cf104_ = d_60___mcc_h8_
            return D12_DC37(False, False, False)

    @staticmethod
    def fm31(p0, p1, p2, globalState):
        def iife4_():
            coll4_ = _dafny.Map()
            compr_4_: int
            for compr_4_ in (_dafny.SeqWithoutIsStrInference([99])).Elements:
                d_63_v0_: int = compr_4_
                if (d_63_v0_) in (_dafny.SeqWithoutIsStrInference([99])):
                    coll4_[default__.safeModuloInt(d_63_v0_, len(_dafny.Map({724: True})))] = True
            return _dafny.Map(coll4_)
        return (D28_DC70(_dafny.Map({True: _dafny.Set({(_dafny.MultiSet([not(True)])).cardinality, len(_dafny.Set({_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_62_i0_ in range(default__.abs(231))])): not(False)}), _dafny.Map({-273: True}), iife4_()
, _dafny.Map({len(_dafny.SeqWithoutIsStrInference([False])): True})}))})}))).cf116

    @staticmethod
    def fm32(p0, p1, globalState):
        return D7_DC26((-449) * ((0) - ((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ewkrdwrs")))))), len((_dafny.SeqWithoutIsStrInference([-355 for d_64_i0_ in range(default__.abs(584))])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True, not(True)]))]))))

    @staticmethod
    def fm33(p0, globalState):
        return (_dafny.MultiSet([-998, 980]))

    @staticmethod
    def fm34(p0, p1, p2, globalState):
        def iife5_():
            def iife7_():
                coll7_ = _dafny.Map()
                compr_7_: str
                for compr_7_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y')])).Elements:
                    d_65_v1_: str = compr_7_
                    if (d_65_v1_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y')])):
                        coll7_[d_65_v1_] = 758
                return _dafny.Map(coll7_)
            coll5_ = _dafny.Map()
            def iife6_():
                coll6_ = _dafny.Map()
                compr_6_: str
                for compr_6_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y')])).Elements:
                    d_65_v1_: str = compr_6_
                    if (d_65_v1_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y')])):
                        coll6_[d_65_v1_] = 758
                return _dafny.Map(coll6_)
            compr_5_: str
            for compr_5_ in (iife6_()
            ).keys.Elements:
                d_66_v0_: str = compr_5_
                if (d_66_v0_) in (iife7_()
                ):
                    coll5_[d_66_v0_] = _dafny.Map({len(_dafny.Map({True: -263})): True})
            return _dafny.Map(coll5_)
        def iife8_():
            coll8_ = _dafny.Map()
            compr_8_: str
            for compr_8_ in (_dafny.Set({_dafny.CodePoint('y')})).Elements:
                d_67_v2_: str = compr_8_
                if (d_67_v2_) in (_dafny.Set({_dafny.CodePoint('y')})):
                    coll8_[d_67_v2_] = _dafny.Map({-484: False})
            return _dafny.Map(coll8_)
        return (iife5_()
        ) | (iife8_()
        )

    @staticmethod
    def fm35(p0, p1, globalState):
        if True:
            return D4_DC13(D1_DC4(107, _dafny.SeqWithoutIsStrInference([False]), len(_dafny.Map({812: True}))), True, _dafny.SeqWithoutIsStrInference([False]))
        elif True:
            return D4_DC13(D1_DC4(26, _dafny.SeqWithoutIsStrInference([True]), 932), True, _dafny.SeqWithoutIsStrInference([False]))

    @staticmethod
    def fm36(globalState):
        return ((_dafny.SeqWithoutIsStrInference([D4_DC14(D4_DC14(D4_DC14(D4_DC13(D1_DC4(836, _dafny.SeqWithoutIsStrInference([True]), len(_dafny.Map({False: False}))), False, _dafny.SeqWithoutIsStrInference([True])))))]) if not(True) else _dafny.SeqWithoutIsStrInference([D4_DC14(D4_DC13(D1_DC4((0) - (len(_dafny.Map({False: False}))), _dafny.SeqWithoutIsStrInference([False, False, False]), len(_dafny.Map({True: 973}))), True, _dafny.SeqWithoutIsStrInference([True]))), D4_DC14(D4_DC13(D1_DC4(len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_68_i0_ in range(default__.abs(158))])): True})), _dafny.SeqWithoutIsStrInference([False]), (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(0) - ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yjundkn")))]))).cardinality) for d_69_i1_ in range(default__.abs(13))]))]))).cardinality), False, _dafny.SeqWithoutIsStrInference([True]))), D4_DC14(D4_DC13(D1_DC4(521, (D1_DC4((_dafny.MultiSet([False, True])).cardinality, _dafny.SeqWithoutIsStrInference([False]), 842)).cf7, 861), False, _dafny.SeqWithoutIsStrInference([True]))), D4_DC14(D4_DC14(D4_DC13(D1_DC4(len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_70_i2_ in range(default__.abs(420))])), 162])), _dafny.SeqWithoutIsStrInference([False]), 831), not(False), _dafny.SeqWithoutIsStrInference([True])))), D4_DC14(D4_DC14(D4_DC13(D1_DC4(854, _dafny.SeqWithoutIsStrInference([True]), len(_dafny.Set({True, True}))), True, _dafny.SeqWithoutIsStrInference([False]))))]))) + (_dafny.SeqWithoutIsStrInference([D4_DC14(D4_DC13(D1_DC4(219, _dafny.SeqWithoutIsStrInference([True]), len(_dafny.Map({_dafny.Map({True: False}): not(True)}))), not(False), _dafny.SeqWithoutIsStrInference([True, True]))), D4_DC14(D4_DC14(D4_DC13(D1_DC4(-33, _dafny.SeqWithoutIsStrInference([True, False, True]), -491), True, _dafny.SeqWithoutIsStrInference([True, True, True])))), D4_DC14(D4_DC13(D1_DC4(-856, _dafny.SeqWithoutIsStrInference([False]), 476), not(not(False)), _dafny.SeqWithoutIsStrInference([False]))), D4_DC14(D4_DC12(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rqsle"))))]))

    @staticmethod
    def fm37(p0, p1, p2, globalState):
        def iife9_():
            coll9_ = _dafny.Map()
            compr_9_: int
            for compr_9_ in _dafny.IntegerRange(448, -536):
                d_71_v0_: int = compr_9_
                if ((448) <= (d_71_v0_)) and ((d_71_v0_) < (-536)):
                    coll9_[(d_71_v0_) + (409)] = True
            return _dafny.Map(coll9_)
        def iife10_():
            coll10_ = _dafny.Map()
            compr_10_: int
            for compr_10_ in _dafny.IntegerRange(457, 315):
                d_72_v1_: int = compr_10_
                if ((457) <= (d_72_v1_)) and ((d_72_v1_) < (315)):
                    coll10_[(d_72_v1_) * (470)] = (D3_DC10(138, len(_dafny.Set({True})), False, len((D4_DC12(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rdwpis")))).cf23), not(not(True)))).cf21
            return _dafny.Map(coll10_)
        return (iife9_()
        ) | (iife10_()
        )

    @staticmethod
    def fm38(p0, p1, p2, globalState):
        return D12_DC38(default__.safeDivisionInt(234, (_dafny.MultiSet([False])).cardinality))

    @staticmethod
    def fm39(p0, globalState):
        return _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ov")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vywr")), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ci"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dxwjujvlh")))])

    @staticmethod
    def fm40(globalState):
        return ((_dafny.Map({False: True})) | (_dafny.Map({False: not(False)}))) | ((_dafny.Map({True: True})) | (_dafny.Map({not(False): False})))

    @staticmethod
    def fm41(globalState):
        return (D29_DC74(_dafny.SeqWithoutIsStrInference([D1_DC3(_dafny.Map({False: 405}))]))).cf121

    @staticmethod
    def fm42(p0, p1, p2, p3, globalState):
        return D5_DC16(True, not (True) or (False), (not(True) if False else False), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ulqhcujj"))), False)

    @staticmethod
    def fm43(globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q'), _dafny.CodePoint('f')]))])) + (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([_dafny.CodePoint('m')]) for d_73_i0_ in range(default__.abs(961))]))) + ((_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([_dafny.CodePoint('y'), _dafny.CodePoint('d'), _dafny.CodePoint('b'), _dafny.CodePoint('l'), _dafny.CodePoint('h')])])) + (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([_dafny.CodePoint('d')])])))

    @staticmethod
    def fm44(p0, p1, p2, p3, globalState):
        return ((D30_DC79(_dafny.Map({_dafny.SeqWithoutIsStrInference([True, True]): False}))).cf132) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([False, False, True, True]): True}))

    @staticmethod
    def fm45(p0, p1, globalState):
        if False:
            def iife11_():
                coll11_ = _dafny.Map()
                compr_11_: int
                for compr_11_ in _dafny.IntegerRange(-925, 757):
                    d_74_v0_: int = compr_11_
                    if ((-925) <= (d_74_v0_)) and ((d_74_v0_) < (757)):
                        coll11_[(d_74_v0_) * (353)] = -712
                return _dafny.Map(coll11_)
            return (_dafny.Map({601: -351})) | (iife11_()
            )
        elif True:
            def iife12_():
                coll12_ = _dafny.Map()
                compr_12_: int
                for compr_12_ in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([444 for d_75_i0_ in range(default__.abs(282))])])).cardinality]))).Elements:
                    d_76_v1_: int = compr_12_
                    if (d_76_v1_) in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([444 for d_75_i0_ in range(default__.abs(282))])])).cardinality]))):
                        coll12_[default__.safeModuloInt(d_76_v1_, 914)] = -295
                return _dafny.Map(coll12_)
            return (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.Map({False: _dafny.CodePoint('e')}), _dafny.Map({False: _dafny.CodePoint('p')})])): len(_dafny.SeqWithoutIsStrInference([True]))})) | (iife12_()
            )

    @staticmethod
    def fm46(p0, p1, p2, globalState):
        def iife13_():
            coll13_ = _dafny.Map()
            compr_13_: D3
            for compr_13_ in (_dafny.Set({D3_DC10(962, 743, False, 127, not(True))})).Elements:
                d_77_v0_: D3 = compr_13_
                if (d_77_v0_) in (_dafny.Set({D3_DC10(962, 743, False, 127, not(True))})):
                    coll13_[d_77_v0_] = True
            return _dafny.Map(coll13_)
        def iife14_():
            coll14_ = _dafny.Map()
            compr_14_: int
            for compr_14_ in (_dafny.SeqWithoutIsStrInference([468 for d_78_i0_ in range(default__.abs(876))])).Elements:
                d_79_v1_: int = compr_14_
                if (d_79_v1_) in (_dafny.SeqWithoutIsStrInference([468 for d_78_i0_ in range(default__.abs(876))])):
                    coll14_[default__.safeModuloInt(d_79_v1_, len(_dafny.SeqWithoutIsStrInference([True, False])))] = 67
            return _dafny.Map(coll14_)
        def iife15_():
            coll15_ = _dafny.Map()
            compr_15_: D3
            for compr_15_ in (_dafny.SeqWithoutIsStrInference([D3_DC10(len(_dafny.SeqWithoutIsStrInference([False, True])), (0) - (len(_dafny.Map({False: True}))), False, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b"))), False) for d_80_i1_ in range(default__.abs(727))])).Elements:
                d_81_v2_: D3 = compr_15_
                if (d_81_v2_) in (_dafny.SeqWithoutIsStrInference([D3_DC10(len(_dafny.SeqWithoutIsStrInference([False, True])), (0) - (len(_dafny.Map({False: True}))), False, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b"))), False) for d_80_i1_ in range(default__.abs(727))])):
                    coll15_[d_81_v2_] = False
            return _dafny.Map(coll15_)
        return (iife13_()
        ) | ((_dafny.Map({D3_DC10((0) - (len(iife14_()
)), 971, True, -176, not(False)): False})) | (iife15_()
        ))

    @staticmethod
    def fm47(p0, p1, p2, p3, globalState):
        return D2_DC6((256) * (len(_dafny.Map({len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q'), _dafny.CodePoint('s')])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rne")))})): 733}))), not(((D12_DC37(False, False, True)).cf66) or (True)))

    @staticmethod
    def fm48(globalState):
        return (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oshom"))), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_82_i0_ in range(default__.abs(847))]))])) | (_dafny.MultiSet([755]))

    @staticmethod
    def fm49(p0, p1, p2, globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([-982, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c")))])])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_83_i1_ in range(default__.abs(166))])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ourm")))]) for d_84_i0_ in range(default__.abs(207))]))) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([not(False), not(True), True, not(False)])) for d_85_i3_ in range(default__.abs(-651))]) for d_86_i2_ in range(default__.abs(795))]))

    @staticmethod
    def fm50(p0, p1, p2, globalState):
        return _dafny.Map({_dafny.CodePoint('f'): len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aqrhmq"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ioc"))))})

    @staticmethod
    def fm51(p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference([D0_DC0(_dafny.SeqWithoutIsStrInference([214, 168, -128, (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vx"))))])), D0_DC0(_dafny.SeqWithoutIsStrInference([-549, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eurutqjd")))])), (D0_DC0(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ioa")))])) if not(True) else D0_DC0(_dafny.SeqWithoutIsStrInference([921 for d_87_i0_ in range(default__.abs(291))]))), D0_DC0(_dafny.SeqWithoutIsStrInference([318])), D0_DC0(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([479 for d_88_i1_ in range(default__.abs(4))]))]))])

    @staticmethod
    def fm52(p0, globalState):
        def iife16_():
            def iife28_():
                def iife34_():
                    def iife37_():
                        coll37_ = _dafny.Map()
                        compr_37_: int
                        for compr_37_ in _dafny.IntegerRange(448, -932):
                            d_89_v0_: int = compr_37_
                            if ((448) <= (d_89_v0_)) and ((d_89_v0_) < (-932)):
                                coll37_[default__.safeModuloInt(d_89_v0_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_90_i0_ in range(default__.abs(265))])))] = -858
                        return _dafny.Map(coll37_)
                    def iife38_():
                        coll38_ = _dafny.Map()
                        compr_38_: int
                        for compr_38_ in _dafny.IntegerRange(442, -884):
                            d_91_v1_: int = compr_38_
                            if ((442) <= (d_91_v1_)) and ((d_91_v1_) < (-884)):
                                coll38_[(d_91_v1_) - (len(_dafny.SeqWithoutIsStrInference([False])))] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sdy")))
                        return _dafny.Map(coll38_)
                    coll34_ = _dafny.Set()
                    def iife35_():
                        coll35_ = _dafny.Map()
                        compr_35_: int
                        for compr_35_ in _dafny.IntegerRange(448, -932):
                            d_89_v0_: int = compr_35_
                            if ((448) <= (d_89_v0_)) and ((d_89_v0_) < (-932)):
                                coll35_[default__.safeModuloInt(d_89_v0_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_90_i0_ in range(default__.abs(265))])))] = -858
                        return _dafny.Map(coll35_)
                    def iife36_():
                        coll36_ = _dafny.Map()
                        compr_36_: int
                        for compr_36_ in _dafny.IntegerRange(442, -884):
                            d_91_v1_: int = compr_36_
                            if ((442) <= (d_91_v1_)) and ((d_91_v1_) < (-884)):
                                coll36_[(d_91_v1_) - (len(_dafny.SeqWithoutIsStrInference([False])))] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sdy")))
                        return _dafny.Map(coll36_)
                    compr_34_: _dafny.Map
                    for compr_34_ in (_dafny.MultiSet([_dafny.Map({680: -8}), iife35_()
                    , iife36_()
                    ])).Elements:
                        d_98_v2_: _dafny.Map = compr_34_
                        if (d_98_v2_) in (_dafny.MultiSet([_dafny.Map({680: -8}), iife37_()
                        , iife38_()
                        ])):
                            coll34_ = coll34_.union(_dafny.Set([d_98_v2_]))
                    return _dafny.Set(coll34_)
                coll28_ = _dafny.Set()
                def iife29_():
                    def iife32_():
                        coll32_ = _dafny.Map()
                        compr_32_: int
                        for compr_32_ in _dafny.IntegerRange(448, -932):
                            d_89_v0_: int = compr_32_
                            if ((448) <= (d_89_v0_)) and ((d_89_v0_) < (-932)):
                                coll32_[default__.safeModuloInt(d_89_v0_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_90_i0_ in range(default__.abs(265))])))] = -858
                        return _dafny.Map(coll32_)
                    def iife33_():
                        coll33_ = _dafny.Map()
                        compr_33_: int
                        for compr_33_ in _dafny.IntegerRange(442, -884):
                            d_91_v1_: int = compr_33_
                            if ((442) <= (d_91_v1_)) and ((d_91_v1_) < (-884)):
                                coll33_[(d_91_v1_) - (len(_dafny.SeqWithoutIsStrInference([False])))] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sdy")))
                        return _dafny.Map(coll33_)
                    coll29_ = _dafny.Set()
                    def iife30_():
                        coll30_ = _dafny.Map()
                        compr_30_: int
                        for compr_30_ in _dafny.IntegerRange(448, -932):
                            d_89_v0_: int = compr_30_
                            if ((448) <= (d_89_v0_)) and ((d_89_v0_) < (-932)):
                                coll30_[default__.safeModuloInt(d_89_v0_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_90_i0_ in range(default__.abs(265))])))] = -858
                        return _dafny.Map(coll30_)
                    def iife31_():
                        coll31_ = _dafny.Map()
                        compr_31_: int
                        for compr_31_ in _dafny.IntegerRange(442, -884):
                            d_91_v1_: int = compr_31_
                            if ((442) <= (d_91_v1_)) and ((d_91_v1_) < (-884)):
                                coll31_[(d_91_v1_) - (len(_dafny.SeqWithoutIsStrInference([False])))] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sdy")))
                        return _dafny.Map(coll31_)
                    compr_29_: _dafny.Map
                    for compr_29_ in (_dafny.MultiSet([_dafny.Map({680: -8}), iife30_()
                    , iife31_()
                    ])).Elements:
                        d_96_v2_: _dafny.Map = compr_29_
                        if (d_96_v2_) in (_dafny.MultiSet([_dafny.Map({680: -8}), iife32_()
                        , iife33_()
                        ])):
                            coll29_ = coll29_.union(_dafny.Set([d_96_v2_]))
                    return _dafny.Set(coll29_)
                compr_28_: _dafny.Map
                for compr_28_ in (iife29_()
                ).Elements:
                    d_97_v3_: _dafny.Map = compr_28_
                    if (d_97_v3_) in (iife34_()
                    ):
                        coll28_ = coll28_.union(_dafny.Set([d_97_v3_]))
                return _dafny.Set(coll28_)
            coll16_ = _dafny.Set()
            def iife17_():
                def iife23_():
                    def iife26_():
                        coll26_ = _dafny.Map()
                        compr_26_: int
                        for compr_26_ in _dafny.IntegerRange(448, -932):
                            d_89_v0_: int = compr_26_
                            if ((448) <= (d_89_v0_)) and ((d_89_v0_) < (-932)):
                                coll26_[default__.safeModuloInt(d_89_v0_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_90_i0_ in range(default__.abs(265))])))] = -858
                        return _dafny.Map(coll26_)
                    def iife27_():
                        coll27_ = _dafny.Map()
                        compr_27_: int
                        for compr_27_ in _dafny.IntegerRange(442, -884):
                            d_91_v1_: int = compr_27_
                            if ((442) <= (d_91_v1_)) and ((d_91_v1_) < (-884)):
                                coll27_[(d_91_v1_) - (len(_dafny.SeqWithoutIsStrInference([False])))] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sdy")))
                        return _dafny.Map(coll27_)
                    coll23_ = _dafny.Set()
                    def iife24_():
                        coll24_ = _dafny.Map()
                        compr_24_: int
                        for compr_24_ in _dafny.IntegerRange(448, -932):
                            d_89_v0_: int = compr_24_
                            if ((448) <= (d_89_v0_)) and ((d_89_v0_) < (-932)):
                                coll24_[default__.safeModuloInt(d_89_v0_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_90_i0_ in range(default__.abs(265))])))] = -858
                        return _dafny.Map(coll24_)
                    def iife25_():
                        coll25_ = _dafny.Map()
                        compr_25_: int
                        for compr_25_ in _dafny.IntegerRange(442, -884):
                            d_91_v1_: int = compr_25_
                            if ((442) <= (d_91_v1_)) and ((d_91_v1_) < (-884)):
                                coll25_[(d_91_v1_) - (len(_dafny.SeqWithoutIsStrInference([False])))] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sdy")))
                        return _dafny.Map(coll25_)
                    compr_23_: _dafny.Map
                    for compr_23_ in (_dafny.MultiSet([_dafny.Map({680: -8}), iife24_()
                    , iife25_()
                    ])).Elements:
                        d_94_v2_: _dafny.Map = compr_23_
                        if (d_94_v2_) in (_dafny.MultiSet([_dafny.Map({680: -8}), iife26_()
                        , iife27_()
                        ])):
                            coll23_ = coll23_.union(_dafny.Set([d_94_v2_]))
                    return _dafny.Set(coll23_)
                coll17_ = _dafny.Set()
                def iife18_():
                    def iife21_():
                        coll21_ = _dafny.Map()
                        compr_21_: int
                        for compr_21_ in _dafny.IntegerRange(448, -932):
                            d_89_v0_: int = compr_21_
                            if ((448) <= (d_89_v0_)) and ((d_89_v0_) < (-932)):
                                coll21_[default__.safeModuloInt(d_89_v0_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_90_i0_ in range(default__.abs(265))])))] = -858
                        return _dafny.Map(coll21_)
                    def iife22_():
                        coll22_ = _dafny.Map()
                        compr_22_: int
                        for compr_22_ in _dafny.IntegerRange(442, -884):
                            d_91_v1_: int = compr_22_
                            if ((442) <= (d_91_v1_)) and ((d_91_v1_) < (-884)):
                                coll22_[(d_91_v1_) - (len(_dafny.SeqWithoutIsStrInference([False])))] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sdy")))
                        return _dafny.Map(coll22_)
                    coll18_ = _dafny.Set()
                    def iife19_():
                        coll19_ = _dafny.Map()
                        compr_19_: int
                        for compr_19_ in _dafny.IntegerRange(448, -932):
                            d_89_v0_: int = compr_19_
                            if ((448) <= (d_89_v0_)) and ((d_89_v0_) < (-932)):
                                coll19_[default__.safeModuloInt(d_89_v0_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_90_i0_ in range(default__.abs(265))])))] = -858
                        return _dafny.Map(coll19_)
                    def iife20_():
                        coll20_ = _dafny.Map()
                        compr_20_: int
                        for compr_20_ in _dafny.IntegerRange(442, -884):
                            d_91_v1_: int = compr_20_
                            if ((442) <= (d_91_v1_)) and ((d_91_v1_) < (-884)):
                                coll20_[(d_91_v1_) - (len(_dafny.SeqWithoutIsStrInference([False])))] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sdy")))
                        return _dafny.Map(coll20_)
                    compr_18_: _dafny.Map
                    for compr_18_ in (_dafny.MultiSet([_dafny.Map({680: -8}), iife19_()
                    , iife20_()
                    ])).Elements:
                        d_92_v2_: _dafny.Map = compr_18_
                        if (d_92_v2_) in (_dafny.MultiSet([_dafny.Map({680: -8}), iife21_()
                        , iife22_()
                        ])):
                            coll18_ = coll18_.union(_dafny.Set([d_92_v2_]))
                    return _dafny.Set(coll18_)
                compr_17_: _dafny.Map
                for compr_17_ in (iife18_()
                ).Elements:
                    d_93_v3_: _dafny.Map = compr_17_
                    if (d_93_v3_) in (iife23_()
                    ):
                        coll17_ = coll17_.union(_dafny.Set([d_93_v3_]))
                return _dafny.Set(coll17_)
            compr_16_: _dafny.Map
            for compr_16_ in (iife17_()
            ).Elements:
                d_95_v4_: _dafny.Map = compr_16_
                if (d_95_v4_) in (iife28_()
                ):
                    coll16_ = coll16_.union(_dafny.Set([d_95_v4_]))
            return _dafny.Set(coll16_)
        def iife39_():
            def iife43_():
                coll43_ = _dafny.Map()
                compr_43_: int
                for compr_43_ in _dafny.IntegerRange(930, 503):
                    d_99_v5_: int = compr_43_
                    if ((930) <= (d_99_v5_)) and ((d_99_v5_) < (503)):
                        coll43_[(d_99_v5_) + (369)] = 50
                return _dafny.Map(coll43_)
            def iife44_():
                coll44_ = _dafny.Map()
                compr_44_: int
                for compr_44_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tu"))), 494])).Elements:
                    d_100_v6_: int = compr_44_
                    if (d_100_v6_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tu"))), 494])):
                        coll44_[default__.safeDivisionInt(d_100_v6_, -900)] = 595
                return _dafny.Map(coll44_)
            def iife45_():
                coll45_ = _dafny.Map()
                compr_45_: int
                for compr_45_ in _dafny.IntegerRange(210, 130):
                    d_101_v7_: int = compr_45_
                    if ((210) <= (d_101_v7_)) and ((d_101_v7_) < (130)):
                        coll45_[default__.safeDivisionInt(d_101_v7_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v'), _dafny.CodePoint('i')])))] = 327
                return _dafny.Map(coll45_)
            coll39_ = _dafny.Set()
            def iife40_():
                coll40_ = _dafny.Map()
                compr_40_: int
                for compr_40_ in _dafny.IntegerRange(930, 503):
                    d_99_v5_: int = compr_40_
                    if ((930) <= (d_99_v5_)) and ((d_99_v5_) < (503)):
                        coll40_[(d_99_v5_) + (369)] = 50
                return _dafny.Map(coll40_)
            def iife41_():
                coll41_ = _dafny.Map()
                compr_41_: int
                for compr_41_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tu"))), 494])).Elements:
                    d_100_v6_: int = compr_41_
                    if (d_100_v6_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tu"))), 494])):
                        coll41_[default__.safeDivisionInt(d_100_v6_, -900)] = 595
                return _dafny.Map(coll41_)
            def iife42_():
                coll42_ = _dafny.Map()
                compr_42_: int
                for compr_42_ in _dafny.IntegerRange(210, 130):
                    d_101_v7_: int = compr_42_
                    if ((210) <= (d_101_v7_)) and ((d_101_v7_) < (130)):
                        coll42_[default__.safeDivisionInt(d_101_v7_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v'), _dafny.CodePoint('i')])))] = 327
                return _dafny.Map(coll42_)
            compr_39_: _dafny.Map
            for compr_39_ in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([iife40_()
            , iife41_()
            , iife42_()
            ]))).Elements:
                d_102_v8_: _dafny.Map = compr_39_
                if (d_102_v8_) in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([iife43_()
                , iife44_()
                , iife45_()
                ]))):
                    coll39_ = coll39_.union(_dafny.Set([d_102_v8_]))
            return _dafny.Set(coll39_)
        def iife46_():
            coll46_ = _dafny.Set()
            compr_46_: _dafny.Map
            for compr_46_ in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sxrucrnd")))): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wrre")))})]))).Elements:
                d_103_v9_: _dafny.Map = compr_46_
                if (d_103_v9_) in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sxrucrnd")))): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wrre")))})]))):
                    coll46_ = coll46_.union(_dafny.Set([d_103_v9_]))
            return _dafny.Set(coll46_)
        return (iife16_()
        ) - ((iife39_()
        ).intersection((D31_DC83(iife46_()
)).cf139))

    @staticmethod
    def fm53(p0, p1, globalState):
        return (_dafny.Map({_dafny.SeqWithoutIsStrInference([575, len(_dafny.Map({804: 583})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b")))]): 774})) | ((_dafny.Map({_dafny.SeqWithoutIsStrInference([261]): (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sffbrlbb"))))})) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ahxykkxr")))]): -484})))

    @staticmethod
    def fm54(p0, p1, p2, globalState):
        source4_ = D7_DC25(True, True, 36)
        if source4_.is_DC25:
            d_104___mcc_h0_ = source4_.cf45
            d_105___mcc_h1_ = source4_.cf46
            d_106___mcc_h2_ = source4_.cf47
            d_107_cf47_ = d_106___mcc_h2_
            d_108_cf46_ = d_105___mcc_h1_
            d_109_cf45_ = d_104___mcc_h0_
            return (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ryinnfi")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jiaps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uoh"))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i"))]))
        elif source4_.is_DC26:
            d_110___mcc_h3_ = source4_.cf48
            d_111___mcc_h4_ = source4_.cf49
            d_112_cf49_ = d_111___mcc_h4_
            d_113_cf48_ = d_110___mcc_h3_
            return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wdhfq"))])
        elif source4_.is_DC27:
            d_114___mcc_h5_ = source4_.cf50
            d_115___mcc_h6_ = source4_.cf51
            d_116___mcc_h7_ = source4_.cf52
            d_117_cf52_ = d_116___mcc_h7_
            d_118_cf51_ = d_115___mcc_h6_
            d_119_cf50_ = d_114___mcc_h5_
            return (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xodt"))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lmgchov")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jvb")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cdbl"))]))
        elif True:
            d_120___mcc_h8_ = source4_.cf44
            d_121_cf44_ = d_120___mcc_h8_
            if not(True):
                return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xxwbaqyk"))])
            elif True:
                return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_122_i1_ in range(default__.abs(419))]) for d_123_i0_ in range(default__.abs(576))])

    @staticmethod
    def fm55(p0, p1, p2, globalState):
        return (_dafny.Map({-244: _dafny.CodePoint('p')})) | (_dafny.Map({693: _dafny.CodePoint('a')}))

    @staticmethod
    def fm56(p0, globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([False, False]), (_dafny.MultiSet([not(False)])) - (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True, not(not(True))]))), _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False])), (_dafny.MultiSet([False])) - (_dafny.MultiSet([True, not(not(False)), True]))])

    @staticmethod
    def fm57(p0, p1, p2, globalState):
        return _dafny.Set({(_dafny.CodePoint('u') if not(False) else _dafny.CodePoint('s'))})

    @staticmethod
    def fm58(p0, globalState):
        return D6_DC22()

    @staticmethod
    def m0(p0, p1, p2, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: _dafny.Array = _dafny.Array(None, 0)
        r3: bool = False
        d_124_v0_: int
        d_124_v0_ = -490
        if not(not(default__.fm0(d_124_v0_, (d_124_v0_) - (d_124_v0_), globalState))):
            d_125_v1_: bool
            d_125_v1_ = False
            r3 = d_125_v1_
            d_126_v2_: _dafny.Array
            nw0_ = _dafny.Array(False, 12)
            d_126_v2_ = nw0_
            d_127_v3_: _dafny.MultiSet
            d_127_v3_ = _dafny.MultiSet([d_126_v2_, d_126_v2_])
            d_128_v4_: C6
            nw1_ = C6()
            nw1_.ctor__(True, (d_127_v3_).intersection(d_127_v3_), d_124_v0_)
            d_128_v4_ = nw1_
            nw2_ = C6()
            nw2_.ctor__(d_125_v1_, d_127_v3_, (default__.fm1(globalState) if d_125_v1_ else len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q")))))
            d_128_v4_ = nw2_
            d_129_v5_: D7
            d_129_v5_ = D7_DC25(d_125_v1_, d_125_v1_, d_124_v0_)
            d_129_v5_ = d_129_v5_
            d_130_v6_: _dafny.Set
            d_130_v6_ = _dafny.Set({(d_128_v4_).f8, (d_128_v4_).f8, d_125_v1_, (d_128_v4_).f8})
            d_131_v7_: D14
            d_131_v7_ = D14_DC41(d_130_v6_)
            d_132_v8_: _dafny.Seq
            d_132_v8_ = _dafny.SeqWithoutIsStrInference([d_131_v7_, d_131_v7_, d_131_v7_])
            d_133_v9_: _dafny.Seq
            d_133_v9_ = _dafny.SeqWithoutIsStrInference([(d_128_v4_).f8])
            d_134_v10_: _dafny.Map
            d_134_v10_ = _dafny.Map({d_132_v8_: ((_dafny.SeqWithoutIsStrInference([d_125_v1_, d_125_v1_])) + (d_133_v9_)).set(default__.safeIndex(d_124_v0_, len((_dafny.SeqWithoutIsStrInference([d_125_v1_, d_125_v1_])) + (d_133_v9_))), (d_128_v4_).f8)})
            pat_let_tv0_ = d_130_v6_
            def iife47_(_pat_let0_0):
                def iife48_(d_135_dt__update__tmp_h0_):
                    def iife49_(_pat_let1_0):
                        def iife50_(d_136_dt__update_hcf71_h0_):
                            return D14_DC41(d_136_dt__update_hcf71_h0_)
                        return iife50_(_pat_let1_0)
                    return iife49_(pat_let_tv0_)
                return iife48_(_pat_let0_0)
            d_134_v10_ = (d_134_v10_).set(_dafny.SeqWithoutIsStrInference([iife47_(d_131_v7_), d_131_v7_, d_131_v7_]), (d_133_v9_) + (d_133_v9_))
            r1 = d_124_v0_
        elif True:
            d_137_v11_: bool
            d_137_v11_ = True
            d_138_v12_: _dafny.Map
            d_138_v12_ = _dafny.Map({d_137_v11_: d_124_v0_})
            d_139_v13_: C1
            nw3_ = C1()
            nw3_.ctor__(len(p1), ((d_138_v12_)[d_137_v11_] if (d_137_v11_) in (d_138_v12_) else d_124_v0_), d_124_v0_, d_137_v11_)
            d_139_v13_ = nw3_
            d_140_v14_: _dafny.Seq
            d_140_v14_ = _dafny.SeqWithoutIsStrInference([d_137_v11_, True, False])
            nw4_ = C1()
            nw4_.ctor__(d_139_v13_.f17, d_139_v13_.f17, d_124_v0_, (d_140_v14_)[default__.safeIndex(d_124_v0_, len(d_140_v14_))])
            d_139_v13_ = nw4_
            d_141_v15_: _dafny.Seq
            d_141_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mq"))
            d_142_v16_: _dafny.Seq
            d_142_v16_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_143_i1_ in range(default__.abs(958))])])
            d_144_v17_: _dafny.Seq
            d_144_v17_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_145_i0_ in range(default__.abs(602))]), d_141_v15_, (d_142_v16_)[default__.safeIndex(d_139_v13_.f17, len(d_142_v16_))]])
            r3 = not ((d_144_v17_) < (d_142_v16_)) or (d_137_v11_)
            r0 = default__.fm1(globalState)
            d_141_v15_ = d_141_v15_
            d_146_v18_: T2
            nw5_ = C1()
            nw5_.ctor__(d_139_v13_.f17, -873, d_124_v0_, d_137_v11_)
            d_146_v18_ = nw5_
            nw6_ = C1()
            nw6_.ctor__(d_139_v13_.f17, (d_146_v18_).f15, (d_146_v18_).f15, d_137_v11_)
            d_146_v18_ = nw6_
        d_147_v19_: _dafny.Map
        d_147_v19_ = _dafny.Map({d_124_v0_: d_124_v0_})
        d_148_v20_: bool
        d_148_v20_ = False
        d_149_v21_: C5
        nw7_ = C5()
        nw7_.ctor__(((d_147_v19_)[(_dafny.MultiSet([d_148_v20_])).cardinality] if ((_dafny.MultiSet([d_148_v20_])).cardinality) in (d_147_v19_) else d_124_v0_))
        d_149_v21_ = nw7_
        d_150_v22_: _dafny.Map
        d_150_v22_ = _dafny.Map({(d_124_v0_ if d_148_v20_ else d_124_v0_): d_147_v19_})
        rhs0_ = d_148_v20_
        rhs1_ = d_149_v21_
        rhs2_ = (d_149_v21_).fm8(globalState)
        rhs3_ = d_150_v22_
        r3 = rhs0_
        d_149_v21_ = rhs1_
        r3 = rhs2_
        d_150_v22_ = rhs3_
        d_151_v23_: _dafny.Seq
        d_151_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wamgpv"))
        d_152_v24_: C8
        nw8_ = C8()
        nw8_.ctor__()
        d_152_v24_ = nw8_
        d_153_v25_: _dafny.Set
        d_153_v25_ = _dafny.Set({d_152_v24_, d_152_v24_})
        d_154_v26_: _dafny.Seq
        d_154_v26_ = _dafny.SeqWithoutIsStrInference([d_148_v20_, False, d_148_v20_, d_148_v20_, True])
        d_155_v27_: _dafny.Seq
        d_155_v27_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_148_v20_, d_148_v20_, not(d_148_v20_)]), d_154_v26_])
        d_156_v28_: _dafny.Set
        d_156_v28_ = _dafny.Set({d_124_v0_, d_124_v0_})
        d_157_v29_: _dafny.Seq
        d_157_v29_ = _dafny.SeqWithoutIsStrInference([d_147_v19_, d_147_v19_])
        d_158_v30_: _dafny.MultiSet
        d_158_v30_ = _dafny.MultiSet([d_148_v20_, d_148_v20_, d_148_v20_, d_148_v20_])
        d_159_v31_: _dafny.Array
        nw9_ = _dafny.Array(None, 23)
        nw9_[int(0)] = default__.safeModuloInt(default__.fm1(globalState), d_124_v0_)
        nw9_[int(1)] = len(d_151_v23_)
        nw9_[int(2)] = len((_dafny.Set({d_152_v24_})) | (d_153_v25_))
        nw9_[int(3)] = d_124_v0_
        nw9_[int(4)] = d_124_v0_
        nw9_[int(5)] = (0) - (-640)
        nw9_[int(6)] = len(((d_155_v27_)[default__.safeIndex(268, len(d_155_v27_))]).set(default__.safeIndex(d_124_v0_, len((d_155_v27_)[default__.safeIndex(268, len(d_155_v27_))])), d_148_v20_))
        nw9_[int(7)] = d_124_v0_
        nw9_[int(8)] = len(d_151_v23_)
        nw9_[int(9)] = (d_124_v0_ if d_148_v20_ else d_124_v0_)
        nw9_[int(10)] = (d_152_v24_).fm6(d_124_v0_, d_156_v28_, d_148_v20_, len(d_151_v23_), globalState)
        nw9_[int(11)] = d_124_v0_
        nw9_[int(12)] = d_124_v0_
        nw9_[int(13)] = d_124_v0_
        nw9_[int(14)] = len(d_147_v19_)
        nw9_[int(15)] = d_124_v0_
        nw9_[int(16)] = ((0) - (d_124_v0_)) * (d_124_v0_)
        nw9_[int(17)] = d_124_v0_
        nw9_[int(18)] = d_124_v0_
        nw9_[int(19)] = len((d_147_v19_) | ((d_157_v29_)[default__.safeIndex(d_124_v0_, len(d_157_v29_))]))
        nw9_[int(20)] = -763
        nw9_[int(21)] = (d_158_v30_).cardinality
        nw9_[int(22)] = d_124_v0_
        d_159_v31_ = nw9_
        r2 = d_159_v31_
        d_160_v32_: int
        d_161_v33_: bool
        out0_: int
        out1_: bool
        out0_, out1_ = (d_152_v24_).m5((p1) + (p1), globalState)
        d_160_v32_ = out0_
        d_161_v33_ = out1_
        d_151_v23_ = (d_151_v23_).set(default__.safeIndex(d_124_v0_, len(d_151_v23_)), _dafny.CodePoint('v'))
        r0 = (0) - (default__.safeModuloInt(d_124_v0_, default__.safeDivisionInt(d_160_v32_, ((d_147_v19_)[d_124_v0_] if (d_124_v0_) in (d_147_v19_) else 550))))
        d_162_v34_: D3
        d_162_v34_ = D3_DC10(d_160_v32_, d_124_v0_, False, d_160_v32_, d_148_v20_)
        pat_let_tv1_ = d_124_v0_
        pat_let_tv2_ = d_124_v0_
        pat_let_tv3_ = d_124_v0_
        pat_let_tv4_ = d_160_v32_
        pat_let_tv5_ = d_158_v30_
        pat_let_tv6_ = d_158_v30_
        def lambda0_(source5_):
            if source5_.is_DC8:
                d_163___mcc_h0_ = source5_.cf13
                d_164___mcc_h1_ = source5_.cf14
                d_165_cf14_ = d_164___mcc_h1_
                d_166_cf13_ = d_163___mcc_h0_
                return pat_let_tv1_
            elif source5_.is_DC9:
                d_167___mcc_h2_ = source5_.cf15
                d_168___mcc_h3_ = source5_.cf16
                d_169_cf16_ = d_168___mcc_h3_
                d_170_cf15_ = d_167___mcc_h2_
                return pat_let_tv2_
            elif source5_.is_DC10:
                d_171___mcc_h4_ = source5_.cf17
                d_172___mcc_h5_ = source5_.cf18
                d_173___mcc_h6_ = source5_.cf19
                d_174___mcc_h7_ = source5_.cf20
                d_175___mcc_h8_ = source5_.cf21
                d_176_cf21_ = d_175___mcc_h8_
                d_177_cf20_ = d_174___mcc_h7_
                d_178_cf19_ = d_173___mcc_h6_
                d_179_cf18_ = d_172___mcc_h5_
                d_180_cf17_ = d_171___mcc_h4_
                return default__.safeDivisionInt(d_179_cf18_, d_177_cf20_)
            elif source5_.is_DC7:
                d_181___mcc_h9_ = source5_.cf12
                d_182_cf12_ = d_181___mcc_h9_
                return default__.safeDivisionInt(213, pat_let_tv3_)
            elif True:
                d_183___mcc_h10_ = source5_.cf22
                d_184_cf22_ = d_183___mcc_h10_
                return (pat_let_tv4_) + (len(_dafny.Set({pat_let_tv5_, pat_let_tv6_})))

        r0 = lambda0_(d_162_v34_)
        r1 = d_124_v0_
        r2 = d_159_v31_
        r3 = d_148_v20_
        return r0, r1, r2, r3

    @staticmethod
    def Main(noArgsParameter__):
        d_185_globalState_: GlobalState
        nw10_ = GlobalState()
        nw10_.ctor__(False, -879, 334)
        d_185_globalState_ = nw10_
        d_186_v0_: bool
        d_186_v0_ = False
        d_187_v1_: _dafny.Map
        d_187_v1_ = _dafny.Map({d_186_v0_: d_186_v0_})
        d_188_v2_: int
        d_188_v2_ = 229
        d_189_i0_: int
        d_189_i0_ = 0
        with _dafny.label("0"):
            while ((d_187_v1_)[default__.fm0(d_188_v2_, 136, d_185_globalState_)] if (default__.fm0(d_188_v2_, 136, d_185_globalState_)) in (d_187_v1_) else d_186_v0_):
                with _dafny.c_label("0"):
                    if (d_189_i0_) >= (100):
                        raise _dafny.Break("0")
                    d_189_i0_ = (d_189_i0_) + (1)
                    d_190_v3_: _dafny.Seq
                    d_190_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mkfno"))
                    d_191_v4_: str
                    d_191_v4_ = _dafny.CodePoint('b')
                    d_190_v3_ = (((d_190_v3_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_192_i1_ in range(default__.abs(603))]))) + (d_190_v3_)).set(default__.safeIndex(default__.fm1(d_185_globalState_), len(((d_190_v3_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_193_i1_ in range(default__.abs(603))]))) + (d_190_v3_))), d_191_v4_)
                    d_194_v5_: _dafny.Map
                    d_194_v5_ = _dafny.Map({-343: d_186_v0_})
                    d_195_v6_: _dafny.Seq
                    d_195_v6_ = _dafny.SeqWithoutIsStrInference([(0) - (d_188_v2_)])
                    d_196_v7_: D0
                    d_196_v7_ = D0_DC0(d_195_v6_)
                    d_197_v8_: _dafny.Array
                    nw11_ = _dafny.Array(int(0), 22)
                    d_197_v8_ = nw11_
                    d_198_v9_: _dafny.Map
                    d_198_v9_ = _dafny.Map({len(d_190_v3_): d_197_v8_})
                    d_199_v10_: int
                    d_200_v11_: int
                    d_201_v12_: _dafny.Array
                    d_202_v13_: bool
                    out2_: int
                    out3_: int
                    out4_: _dafny.Array
                    out5_: bool
                    out2_, out3_, out4_, out5_ = default__.m0(d_194_v5_, ((d_196_v7_).cf0) + (d_195_v6_), d_198_v9_, d_185_globalState_)
                    d_199_v10_ = out2_
                    d_200_v11_ = out3_
                    d_201_v12_ = out4_
                    d_202_v13_ = out5_
                    d_200_v11_ = d_199_v10_
                    index0_ = default__.safeIndex(304, (d_201_v12_).length(0))
                    (d_201_v12_)[index0_] = -77
                    index1_ = default__.safeIndex(304, (d_201_v12_).length(0))
                    rhs4_ = _dafny.SeqWithoutIsStrInference([d_188_v2_ for d_203_i2_ in range(default__.abs(191))])
                    rhs5_ = d_200_v11_
                    lhs0_ = d_201_v12_
                    lhs1_ = default__.safeIndex(304, (d_201_v12_).length(0))
                    d_195_v6_ = rhs4_
                    lhs0_[lhs1_] = rhs5_
                    pass
            pass
        d_204_v14_: _dafny.Set
        d_204_v14_ = _dafny.Set({d_186_v0_, d_186_v0_})
        d_205_v15_: _dafny.Seq
        d_205_v15_ = _dafny.SeqWithoutIsStrInference([d_188_v2_, d_188_v2_, d_188_v2_, len(d_204_v14_)])
        d_206_v16_: D0
        d_206_v16_ = D0_DC0(d_205_v15_)
        d_206_v16_ = d_206_v16_
        d_207_v17_: _dafny.Array
        nw12_ = _dafny.Array(int(0), 4)
        d_207_v17_ = nw12_
        d_208_v18_: _dafny.Map
        d_208_v18_ = _dafny.Map({d_188_v2_: d_207_v17_})
        d_209_v19_: int
        d_210_v20_: int
        d_211_v21_: _dafny.Array
        d_212_v22_: bool
        out6_: int
        out7_: int
        out8_: _dafny.Array
        out9_: bool
        out6_, out7_, out8_, out9_ = default__.m0(_dafny.Map({d_188_v2_: d_186_v0_}), (d_205_v15_) + (_dafny.SeqWithoutIsStrInference([d_188_v2_])), d_208_v18_, d_185_globalState_)
        d_209_v19_ = out6_
        d_210_v20_ = out7_
        d_211_v21_ = out8_
        d_212_v22_ = out9_
        d_213_v23_: _dafny.Seq
        d_213_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sc"))
        d_214_v24_: _dafny.MultiSet
        d_214_v24_ = _dafny.MultiSet([955])
        d_215_v25_: _dafny.Array
        nw13_ = _dafny.Array(D0.default()(), 27)
        d_215_v25_ = nw13_
        d_216_v26_: _dafny.Set
        d_216_v26_ = _dafny.Set({d_215_v25_})
        d_217_v27_: _dafny.Seq
        d_217_v27_ = _dafny.SeqWithoutIsStrInference([d_212_v22_])
        d_218_v28_: _dafny.MultiSet
        d_218_v28_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_186_v0_, d_212_v22_]), d_217_v27_, d_217_v27_])
        d_219_v29_: str
        d_219_v29_ = _dafny.CodePoint('w')
        d_220_v30_: _dafny.Seq
        d_220_v30_ = _dafny.SeqWithoutIsStrInference([default__.fm2(d_219_v29_, d_210_v20_, d_210_v20_, d_185_globalState_)])
        rhs6_ = (d_209_v19_) == ((d_214_v24_).cardinality)
        rhs7_ = (d_210_v20_) + (264)
        rhs8_ = (d_216_v26_).ispropersubset(d_216_v26_)
        rhs9_ = d_213_v23_
        rhs10_ = ((_dafny.MultiSet(d_220_v30_)).set(d_217_v27_, default__.abs(d_209_v19_))).issubset(d_218_v28_)
        d_186_v0_ = rhs6_
        d_209_v19_ = rhs7_
        d_212_v22_ = rhs8_
        d_213_v23_ = rhs9_
        d_186_v0_ = rhs10_
        d_221_v31_: D0
        d_221_v31_ = D0_DC1(d_209_v19_, d_209_v19_, d_186_v0_)
        d_222_v32_: D0
        d_222_v32_ = D0_DC2(d_221_v31_)
        d_223_v33_: D0
        d_223_v33_ = D0_DC2(d_221_v31_)
        d_223_v33_ = D0_DC2(d_221_v31_)
        d_224_v34_: C1
        nw14_ = C1()
        nw14_.ctor__((d_210_v20_ if (d_217_v27_)[default__.safeIndex(d_209_v19_, len(d_217_v27_))] else -733), d_209_v19_, d_209_v19_, not(d_186_v0_))
        d_224_v34_ = nw14_
        d_225_v35_: _dafny.Map
        d_225_v35_ = _dafny.Map({d_188_v2_: d_186_v0_})
        d_226_v36_: D2
        d_226_v36_ = D2_DC6(d_209_v19_, d_212_v22_)
        d_227_v37_: _dafny.Map
        d_227_v37_ = _dafny.Map({default__.safeModuloInt((0) - (d_209_v19_), len(d_225_v35_)): (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nyrdn"))).set(default__.safeIndex(len(d_205_v15_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nyrdn")))), (d_224_v34_).fm10(-249, d_186_v0_, d_212_v22_, d_226_v36_, d_185_globalState_))})
        d_227_v37_ = (d_227_v37_).set((len(d_187_v1_)) + (799), d_213_v23_)
        (d_224_v34_).f17 = default__.fm1(d_185_globalState_)
        d_228_v38_: _dafny.Set
        d_228_v38_ = _dafny.Set({d_210_v20_, -99})
        d_229_i3_: int
        d_229_i3_ = 0
        with _dafny.label("1"):
            while (d_228_v38_).ispropersubset(d_228_v38_):
                with _dafny.c_label("1"):
                    if (d_229_i3_) >= (100):
                        raise _dafny.Break("1")
                    d_229_i3_ = (d_229_i3_) + (1)
                    (d_224_v34_).f17 = (((d_214_v24_) - (d_214_v24_)).set(d_224_v34_.f17, default__.abs(d_224_v34_.f17))).cardinality
                    d_230_v39_: C7
                    nw15_ = C7()
                    nw15_.ctor__(d_213_v23_, d_209_v19_)
                    d_230_v39_ = nw15_
                    nw16_ = C7()
                    nw16_.ctor__(_dafny.SeqWithoutIsStrInference([d_219_v29_ for d_231_i4_ in range(default__.abs(-810))]), d_224_v34_.f17)
                    d_230_v39_ = nw16_
                    d_232_v40_: _dafny.Array
                    def lambda1_(d_233_v15_):
                        def lambda2_(d_234_i5_):
                            return d_233_v15_

                        return lambda2_

                    init0_ = lambda1_(d_205_v15_)
                    nw17_ = _dafny.Array(None, 10)
                    for i0_0_ in range(nw17_.length(0)):
                        nw17_[i0_0_] = init0_(i0_0_)
                    d_232_v40_ = nw17_
                    d_235_v41_: _dafny.Map
                    d_235_v41_ = _dafny.Map({d_232_v40_: not(d_186_v0_)})
                    if ((d_235_v41_)[d_232_v40_] if (d_232_v40_) in (d_235_v41_) else not(d_212_v22_)):
                        d_236_v42_: _dafny.Array
                        nw18_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 23)
                        d_236_v42_ = nw18_
                        index2_ = default__.safeIndex(174, (d_236_v42_).length(0))
                        (d_236_v42_)[index2_] = d_230_v39_.f7
                        index3_ = default__.safeIndex(174, (d_236_v42_).length(0))
                        (d_236_v42_)[index3_] = d_230_v39_.f7
                        d_186_v0_ = ((d_230_v39_.f7) + (d_230_v39_.f7)) < ((d_236_v42_)[default__.safeIndex(174, (d_236_v42_).length(0))])
                        d_212_v22_ = d_212_v22_
                        d_237_v43_: C8
                        nw19_ = C8()
                        nw19_.ctor__()
                        d_237_v43_ = nw19_
                        d_238_v44_: _dafny.Set
                        d_238_v44_ = _dafny.Set({d_237_v43_, d_237_v43_})
                        d_239_v45_: _dafny.Map
                        d_239_v45_ = _dafny.Map({d_188_v2_: d_238_v44_})
                        index4_ = default__.safeIndex(968, (d_207_v17_).length(0))
                        (d_207_v17_)[index4_] = len(((d_239_v45_)[d_188_v2_] if (d_188_v2_) in (d_239_v45_) else d_238_v44_))
                        index5_ = default__.safeIndex(968, (d_207_v17_).length(0))
                        (d_207_v17_)[index5_] = default__.safeDivisionInt((d_214_v24_).cardinality, default__.safeDivisionInt(d_224_v34_.f17, d_224_v34_.f17))
                        d_240_v46_: C0
                        nw20_ = C0()
                        nw20_.ctor__(d_224_v34_.f17)
                        d_240_v46_ = nw20_
                        d_241_v47_: D20
                        d_241_v47_ = D20_DC55(d_230_v39_.f7, d_240_v46_, d_224_v34_.f17)
                        (d_230_v39_).f7 = (d_241_v47_).cf101
                    elif True:
                        d_242_v48_: _dafny.MultiSet
                        d_242_v48_ = _dafny.MultiSet([d_219_v29_])
                        d_243_v49_: _dafny.Map
                        d_243_v49_ = _dafny.Map({d_209_v19_: d_242_v48_})
                        rhs11_ = d_243_v49_
                        rhs12_ = (d_224_v34_).fm5(default__.fm23(d_186_v0_, d_185_globalState_), (d_212_v22_) or (d_186_v0_), d_185_globalState_)
                        d_243_v49_ = rhs11_
                        d_212_v22_ = rhs12_
                        d_244_v50_: D6
                        d_244_v50_ = D6_DC22()
                        d_245_v51_: _dafny.Set
                        d_245_v51_ = _dafny.Set({d_244_v50_, default__.fm58(((d_187_v1_)[d_186_v0_] if (d_186_v0_) in (d_187_v1_) else d_212_v22_), d_185_globalState_)})
                        d_246_v52_: _dafny.Seq
                        d_246_v52_ = _dafny.SeqWithoutIsStrInference([d_244_v50_])
                        d_247_v54_: _dafny.Array
                        nw21_ = _dafny.Array(None, 13)
                        nw21_[int(0)] = d_245_v51_
                        nw21_[int(1)] = d_245_v51_
                        nw21_[int(2)] = d_245_v51_
                        nw21_[int(3)] = (d_245_v51_) - (d_245_v51_)
                        nw21_[int(4)] = d_245_v51_
                        nw21_[int(5)] = d_245_v51_
                        nw21_[int(6)] = (_dafny.Set({d_244_v50_}) if d_186_v0_ else d_245_v51_)
                        nw21_[int(7)] = d_245_v51_
                        nw21_[int(8)] = (d_245_v51_ if True else d_245_v51_)
                        nw21_[int(9)] = d_245_v51_
                        nw21_[int(10)] = d_245_v51_
                        def iife51_():
                            coll47_ = _dafny.Set()
                            compr_47_: D6
                            for compr_47_ in (d_246_v52_).Elements:
                                d_248_v53_: D6 = compr_47_
                                if (d_248_v53_) in (d_246_v52_):
                                    coll47_ = coll47_.union(_dafny.Set([d_248_v53_]))
                            return _dafny.Set(coll47_)
                        nw21_[int(11)] = iife51_()
                        
                        nw21_[int(12)] = d_245_v51_
                        d_247_v54_ = nw21_
                        d_247_v54_ = d_247_v54_
                        d_249_v55_: _dafny.Array
                        nw22_ = _dafny.Array(False, 19)
                        d_249_v55_ = nw22_
                        d_249_v55_ = d_249_v55_
                        d_186_v0_ = d_186_v0_
                        d_186_v0_ = (d_230_v39_).fm5(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vtsqthjmq")), (d_228_v38_).isdisjoint(d_228_v38_), d_185_globalState_)
                    d_250_v56_: _dafny.MultiSet
                    d_250_v56_ = _dafny.MultiSet([d_213_v23_, d_230_v39_.f7])
                    d_250_v56_ = d_250_v56_
                    pass
            pass
        d_251_v57_: C0
        nw23_ = C0()
        nw23_.ctor__(d_188_v2_)
        d_251_v57_ = nw23_
        d_252_v58_: _dafny.Array
        def lambda3_(d_253_v22_):
            def lambda4_(d_254_i6_):
                return d_253_v22_

            return lambda4_

        init1_ = lambda3_(d_212_v22_)
        nw24_ = _dafny.Array(None, 22)
        for i0_1_ in range(nw24_.length(0)):
            nw24_[i0_1_] = init1_(i0_1_)
        d_252_v58_ = nw24_
        index6_ = default__.safeIndex(297, (d_252_v58_).length(0))
        (d_252_v58_)[index6_] = not(d_186_v0_)
        d_255_v59_: _dafny.Set
        d_255_v59_ = _dafny.Set({d_252_v58_})
        index7_ = default__.safeIndex(297, (d_252_v58_).length(0))
        (d_252_v58_)[index7_] = not ((d_224_v34_.f17) == (len(d_255_v59_))) or (d_186_v0_)
        d_256_v60_: _dafny.Map
        d_256_v60_ = _dafny.Map({d_224_v34_.f17: d_219_v29_})
        d_257_v61_: C3
        nw25_ = C3()
        nw25_.ctor__(len(_dafny.SeqWithoutIsStrInference([d_188_v2_])), d_213_v23_, default__.fm1(d_185_globalState_))
        d_257_v61_ = nw25_
        d_258_v62_: _dafny.Map
        d_258_v62_ = _dafny.Map({d_257_v61_: d_212_v22_})
        d_259_v63_: D3
        d_259_v63_ = D3_DC8(d_219_v29_, d_252_v58_)
        d_260_v64_: _dafny.Map
        d_260_v64_ = _dafny.Map({d_205_v15_: d_209_v19_})
        d_261_v66_: _dafny.Seq
        d_261_v66_ = _dafny.SeqWithoutIsStrInference([d_256_v60_])
        d_262_v67_: _dafny.Array
        nw26_ = _dafny.Array(None, 29)
        nw26_[int(0)] = d_256_v60_
        nw26_[int(1)] = (d_256_v60_) | (_dafny.Map({len(_dafny.Map({d_258_v62_: d_259_v63_})): d_219_v29_}))
        nw26_[int(2)] = _dafny.Map({d_188_v2_: d_219_v29_})
        nw26_[int(3)] = (_dafny.Map({d_251_v57_.f18: d_219_v29_})) | (d_256_v60_)
        nw26_[int(4)] = d_256_v60_
        nw26_[int(5)] = (d_256_v60_) | (d_256_v60_)
        nw26_[int(6)] = d_256_v60_
        nw26_[int(7)] = d_256_v60_
        nw26_[int(8)] = d_256_v60_
        nw26_[int(9)] = (_dafny.Map({len(d_204_v14_): d_219_v29_})).set(len(d_260_v64_), d_219_v29_)
        nw26_[int(10)] = d_256_v60_
        nw26_[int(11)] = d_256_v60_
        nw26_[int(12)] = d_256_v60_
        nw26_[int(13)] = d_256_v60_
        nw26_[int(14)] = d_256_v60_
        nw26_[int(15)] = (d_256_v60_).set(d_209_v19_, d_219_v29_)
        nw26_[int(16)] = d_256_v60_
        nw26_[int(17)] = d_256_v60_
        def iife52_():
            coll48_ = _dafny.Map()
            compr_48_: int
            for compr_48_ in _dafny.IntegerRange(320, 930):
                d_263_v65_: int = compr_48_
                if ((320) <= (d_263_v65_)) and ((d_263_v65_) < (930)):
                    coll48_[default__.safeDivisionInt(d_263_v65_, d_224_v34_.f17)] = d_219_v29_
            return _dafny.Map(coll48_)
        nw26_[int(18)] = (iife52_()
        ) | (d_256_v60_)
        nw26_[int(19)] = default__.fm55(-817, d_224_v34_.f17, (d_257_v61_).f11, d_185_globalState_)
        nw26_[int(20)] = d_256_v60_
        nw26_[int(21)] = (d_256_v60_).set(d_224_v34_.f17, _dafny.CodePoint('b'))
        nw26_[int(22)] = (d_261_v66_)[default__.safeIndex(len(default__.fm37(d_186_v0_, not(d_186_v0_), d_251_v57_.f18, d_185_globalState_)), len(d_261_v66_))]
        nw26_[int(23)] = default__.fm55((0) - (d_209_v19_), 575, d_209_v19_, d_185_globalState_)
        nw26_[int(24)] = d_256_v60_
        nw26_[int(25)] = (d_256_v60_) | (d_256_v60_)
        nw26_[int(26)] = d_256_v60_
        nw26_[int(27)] = (d_256_v60_).set(len(d_217_v27_), _dafny.CodePoint('e'))
        nw26_[int(28)] = d_256_v60_
        d_262_v67_ = nw26_
        d_262_v67_ = d_262_v67_
        d_264_v68_: D24
        d_264_v68_ = D24_DC63(_dafny.MultiSet(d_213_v23_))
        pat_let_tv7_ = d_252_v58_
        pat_let_tv8_ = d_252_v58_
        pat_let_tv9_ = d_186_v0_
        def lambda5_(source6_):
            if source6_.is_DC64:
                return not((pat_let_tv8_)[default__.safeIndex(297, (pat_let_tv7_).length(0))])
            elif True:
                d_265___mcc_h0_ = source6_.cf111
                d_266_cf111_ = d_265___mcc_h0_
                return pat_let_tv9_

        rhs13_ = lambda5_(d_264_v68_)
        rhs14_ = ((default__.fm23((d_252_v58_)[default__.safeIndex(297, (d_252_v58_).length(0))], d_185_globalState_)).set(default__.safeIndex((d_257_v61_).f11, len(default__.fm23((d_252_v58_)[default__.safeIndex(297, (d_252_v58_).length(0))], d_185_globalState_))), d_219_v29_)) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rtmcb")))
        d_186_v0_ = rhs13_
        d_212_v22_ = rhs14_
        d_267_v69_: C8
        nw27_ = C8()
        nw27_.ctor__()
        d_267_v69_ = nw27_
        d_268_v70_: T1
        nw28_ = C4()
        nw28_.ctor__(default__.safeDivisionInt((_dafny.MultiSet([len((d_257_v61_).f12), len((d_257_v61_).f12)])).cardinality, d_209_v19_))
        d_268_v70_ = nw28_
        d_268_v70_ = d_268_v70_
        d_269_v71_: C5
        nw29_ = C5()
        nw29_.ctor__(d_210_v20_)
        d_269_v71_ = nw29_
        hi0_ = d_224_v34_.f17
        for d_270_i7_ in range(len((_dafny.Map({D6_DC22(): d_269_v71_})).set(default__.fm58(d_186_v0_, d_185_globalState_), d_269_v71_)), hi0_):
            index8_ = default__.safeIndex(554, (d_211_v21_).length(0))
            (d_211_v21_)[index8_] = (0) - ((d_257_v61_).f11)
            d_271_v72_: _dafny.Map
            d_271_v72_ = _dafny.Map({d_269_v71_: (d_252_v58_)[default__.safeIndex(297, (d_252_v58_).length(0))]})
            index9_ = default__.safeIndex(554, (d_211_v21_).length(0))
            (d_211_v21_)[index9_] = len((d_271_v72_ if ((d_257_v61_).f12) <= ((d_257_v61_).f12) else d_271_v72_))
            d_272_v73_: C9
            nw30_ = C9()
            nw30_.ctor__(d_251_v57_.f18, True, (d_211_v21_)[default__.safeIndex(554, (d_211_v21_).length(0))])
            d_272_v73_ = nw30_
            d_273_v74_: _dafny.Map
            d_273_v74_ = _dafny.Map({d_272_v73_: d_213_v23_})
            d_274_v75_: _dafny.Array
            nw31_ = _dafny.Array(None, 18)
            nw31_[int(0)] = ((d_257_v61_).f12) + (d_213_v23_)
            nw31_[int(1)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uyi")) if (d_252_v58_)[default__.safeIndex(297, (d_252_v58_).length(0))] else _dafny.SeqWithoutIsStrInference([d_219_v29_ for d_275_i8_ in range(default__.abs(230))]))
            nw31_[int(2)] = _dafny.SeqWithoutIsStrInference([d_219_v29_ for d_276_i9_ in range(default__.abs(433))])
            nw31_[int(3)] = (d_257_v61_).f12
            nw31_[int(4)] = (d_257_v61_).f12
            nw31_[int(5)] = ((d_257_v61_).f12) + ((d_257_v61_).f12)
            nw31_[int(6)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yndayr"))
            nw31_[int(7)] = ((d_273_v74_)[d_272_v73_] if (d_272_v73_) in (d_273_v74_) else d_213_v23_)
            nw31_[int(8)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ntmnqxyon"))
            nw31_[int(9)] = ((d_257_v61_).f12) + ((d_257_v61_).f12)
            nw31_[int(10)] = (d_213_v23_) + (d_213_v23_)
            nw31_[int(11)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lggkiehv"))
            nw31_[int(12)] = (d_257_v61_).f12
            nw31_[int(13)] = (d_257_v61_).f12
            nw31_[int(14)] = (d_257_v61_).f12
            nw31_[int(15)] = d_213_v23_
            nw31_[int(16)] = d_213_v23_
            nw31_[int(17)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b"))
            d_274_v75_ = nw31_
            index10_ = default__.safeIndex(636, (d_274_v75_).length(0))
            (d_274_v75_)[index10_] = ((d_257_v61_).f12).set(default__.safeIndex((d_257_v61_).f11, len((d_257_v61_).f12)), d_219_v29_)
            index11_ = default__.safeIndex(636, (d_274_v75_).length(0))
            index12_ = default__.safeIndex(554, (d_211_v21_).length(0))
            rhs15_ = (d_257_v61_).f12
            rhs16_ = (d_267_v69_).fm6(d_272_v73_.f5, d_228_v38_, (d_212_v22_) and (d_272_v73_.f6), (d_224_v34_.f17) + (d_224_v34_.f17), d_185_globalState_)
            rhs17_ = d_214_v24_
            lhs2_ = d_274_v75_
            lhs3_ = default__.safeIndex(636, (d_274_v75_).length(0))
            lhs4_ = d_211_v21_
            lhs5_ = default__.safeIndex(554, (d_211_v21_).length(0))
            lhs2_[lhs3_] = rhs15_
            lhs4_[lhs5_] = rhs16_
            d_214_v24_ = rhs17_
            d_277_v76_: _dafny.Array
            nw32_ = _dafny.Array(_dafny.Seq({}), 14)
            d_277_v76_ = nw32_
            d_278_v77_: _dafny.Seq
            d_278_v77_ = _dafny.SeqWithoutIsStrInference([d_187_v1_, _dafny.Map({False: True})])
            index13_ = default__.safeIndex(664, (d_277_v76_).length(0))
            (d_277_v76_)[index13_] = (d_278_v77_) + (d_278_v77_)
            index14_ = default__.safeIndex(664, (d_277_v76_).length(0))
            (d_277_v76_)[index14_] = (_dafny.SeqWithoutIsStrInference([_dafny.Map({True: d_186_v0_}) for d_279_i10_ in range(default__.abs(-122))])) + ((d_278_v77_).set(default__.safeIndex(-478, len(d_278_v77_)), _dafny.Map({(d_252_v58_)[default__.safeIndex(297, (d_252_v58_).length(0))]: d_212_v22_})))
            d_188_v2_ = (d_211_v21_)[default__.safeIndex(554, (d_211_v21_).length(0))]
        _dafny.print(_dafny.string_of((d_185_globalState_).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_185_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_185_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_186_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_187_v1_) == (_dafny.Map({False: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_188_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_189_i0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_v14_) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_205_v15_) == (_dafny.SeqWithoutIsStrInference([229, 229, 229, 1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_206_v16_).cf0) == (_dafny.SeqWithoutIsStrInference([229, 229, 229, 1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_208_v18_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_209_v19_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_210_v20_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_211_v21_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_211_v21_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_211_v21_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_211_v21_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_211_v21_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_211_v21_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_211_v21_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_211_v21_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_211_v21_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_211_v21_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_211_v21_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_211_v21_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_211_v21_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_211_v21_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_211_v21_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_211_v21_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_211_v21_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_211_v21_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_211_v21_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_211_v21_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_211_v21_)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_211_v21_)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_211_v21_)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_212_v22_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_213_v23_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_214_v24_) == (_dafny.MultiSet([955]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_216_v26_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_217_v27_) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_218_v28_) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([False, False]), _dafny.SeqWithoutIsStrInference([False]), _dafny.SeqWithoutIsStrInference([False])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_219_v29_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_220_v30_) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True, True, False])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_221_v31_).cf1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_221_v31_).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_221_v31_).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_222_v32_).cf4).cf1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_222_v32_).cf4).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_222_v32_).cf4).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_223_v33_).cf4).cf1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_223_v33_).cf4).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_223_v33_).cf4).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_224_v34_.f17))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_225_v35_) == (_dafny.Map({229: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_226_v36_).cf10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_226_v36_).cf11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_227_v37_) == (_dafny.Map({0: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nyrdh")), 800: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sc"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_228_v38_) == (_dafny.Set({-490, -99}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_229_i3_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_251_v57_.f18))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_252_v58_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_252_v58_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_252_v58_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_252_v58_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_252_v58_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_252_v58_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_252_v58_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_252_v58_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_252_v58_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_252_v58_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_252_v58_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_252_v58_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_252_v58_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_252_v58_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_252_v58_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_252_v58_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_252_v58_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_252_v58_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_252_v58_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_252_v58_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_252_v58_)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_252_v58_)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_255_v59_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_256_v60_) == (_dafny.Map({721: _dafny.CodePoint('w')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_257_v61_).f11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_257_v61_).f12).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_258_v62_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_259_v63_).cf13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_259_v63_).cf14)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_259_v63_).cf14)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_259_v63_).cf14)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_259_v63_).cf14)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_259_v63_).cf14)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_259_v63_).cf14)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_259_v63_).cf14)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_259_v63_).cf14)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_259_v63_).cf14)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_259_v63_).cf14)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_259_v63_).cf14)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_259_v63_).cf14)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_259_v63_).cf14)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_259_v63_).cf14)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_259_v63_).cf14)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_259_v63_).cf14)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_259_v63_).cf14)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_259_v63_).cf14)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_259_v63_).cf14)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_259_v63_).cf14)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_259_v63_).cf14)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_259_v63_).cf14)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_260_v64_) == (_dafny.Map({_dafny.SeqWithoutIsStrInference([229, 229, 229, 1]): -226}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_261_v66_) == (_dafny.SeqWithoutIsStrInference([_dafny.Map({721: _dafny.CodePoint('w')})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v67_)[0]) == (_dafny.Map({721: _dafny.CodePoint('w')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v67_)[1]) == (_dafny.Map({721: _dafny.CodePoint('w'), 1: _dafny.CodePoint('w')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v67_)[2]) == (_dafny.Map({229: _dafny.CodePoint('w')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v67_)[3]) == (_dafny.Map({229: _dafny.CodePoint('w'), 721: _dafny.CodePoint('w')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v67_)[4]) == (_dafny.Map({721: _dafny.CodePoint('w')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v67_)[5]) == (_dafny.Map({721: _dafny.CodePoint('w')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v67_)[6]) == (_dafny.Map({721: _dafny.CodePoint('w')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v67_)[7]) == (_dafny.Map({721: _dafny.CodePoint('w')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v67_)[8]) == (_dafny.Map({721: _dafny.CodePoint('w')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v67_)[9]) == (_dafny.Map({1: _dafny.CodePoint('w')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v67_)[10]) == (_dafny.Map({721: _dafny.CodePoint('w')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v67_)[11]) == (_dafny.Map({721: _dafny.CodePoint('w')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v67_)[12]) == (_dafny.Map({721: _dafny.CodePoint('w')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v67_)[13]) == (_dafny.Map({721: _dafny.CodePoint('w')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v67_)[14]) == (_dafny.Map({721: _dafny.CodePoint('w')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v67_)[15]) == (_dafny.Map({721: _dafny.CodePoint('w'), -226: _dafny.CodePoint('w')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v67_)[16]) == (_dafny.Map({721: _dafny.CodePoint('w')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v67_)[17]) == (_dafny.Map({721: _dafny.CodePoint('w')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v67_)[18]) == (_dafny.Map({0: _dafny.CodePoint('w'), 1: _dafny.CodePoint('w'), 721: _dafny.CodePoint('w')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v67_)[19]) == (_dafny.Map({-244: _dafny.CodePoint('p'), 693: _dafny.CodePoint('a')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v67_)[20]) == (_dafny.Map({721: _dafny.CodePoint('w')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v67_)[21]) == (_dafny.Map({721: _dafny.CodePoint('b')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v67_)[22]) == (_dafny.Map({721: _dafny.CodePoint('w')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v67_)[23]) == (_dafny.Map({-244: _dafny.CodePoint('p'), 693: _dafny.CodePoint('a')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v67_)[24]) == (_dafny.Map({721: _dafny.CodePoint('w')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v67_)[25]) == (_dafny.Map({721: _dafny.CodePoint('w')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v67_)[26]) == (_dafny.Map({721: _dafny.CodePoint('w')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v67_)[27]) == (_dafny.Map({721: _dafny.CodePoint('w'), 1: _dafny.CodePoint('e')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v67_)[28]) == (_dafny.Map({721: _dafny.CodePoint('w')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_264_v68_).cf111) == (_dafny.MultiSet([_dafny.CodePoint('s'), _dafny.CodePoint('c')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_268_v70_.f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(int(0), int(0), False)
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

class D0_DC1(D0, NamedTuple('DC1', [('cf1', Any), ('cf2', Any), ('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf1)}, {_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf1 == __o.cf1 and self.cf2 == __o.cf2 and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC0(D0, NamedTuple('DC0', [('cf0', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC0({_dafny.string_of(self.cf0)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC0) and self.cf0 == __o.cf0
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC2(D0, NamedTuple('DC2', [('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2({_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2) and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()


class D1:
    @classmethod
    def default(cls, ):
        return lambda: D1_DC4(int(0), _dafny.Seq({}), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)

class D1_DC4(D1, NamedTuple('DC4', [('cf6', Any), ('cf7', Any), ('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf6)}, {_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf6 == __o.cf6 and self.cf7 == __o.cf7 and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC3(D1, NamedTuple('DC3', [('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC6(int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D2_DC5)

class D2_DC6(D2, NamedTuple('DC6', [('cf10', Any), ('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf10)}, {_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf10 == __o.cf10 and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC5(D2, NamedTuple('DC5', [('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC5({_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC5) and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC8(_dafny.CodePoint('D'), _dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D3_DC8)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D3_DC9)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D3_DC10)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D3_DC7)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D3_DC11)

class D3_DC8(D3, NamedTuple('DC8', [('cf13', Any), ('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC8({_dafny.string_of(self.cf13)}, {_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC8) and self.cf13 == __o.cf13 and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC9(D3, NamedTuple('DC9', [('cf15', Any), ('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf15)}, {_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf15 == __o.cf15 and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC10(D3, NamedTuple('DC10', [('cf17', Any), ('cf18', Any), ('cf19', Any), ('cf20', Any), ('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC10({_dafny.string_of(self.cf17)}, {_dafny.string_of(self.cf18)}, {_dafny.string_of(self.cf19)}, {_dafny.string_of(self.cf20)}, {_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC10) and self.cf17 == __o.cf17 and self.cf18 == __o.cf18 and self.cf19 == __o.cf19 and self.cf20 == __o.cf20 and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC7(D3, NamedTuple('DC7', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC7({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC7) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC11(D3, NamedTuple('DC11', [('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC11({_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC11) and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC13(D1.default()(), False, _dafny.Seq({}))
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

class D4_DC13(D4, NamedTuple('DC13', [('cf24', Any), ('cf25', Any), ('cf26', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC13({_dafny.string_of(self.cf24)}, {_dafny.string_of(self.cf25)}, {_dafny.string_of(self.cf26)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC13) and self.cf24 == __o.cf24 and self.cf25 == __o.cf25 and self.cf26 == __o.cf26
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC12(D4, NamedTuple('DC12', [('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC12({self.cf23.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC12) and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC14(D4, NamedTuple('DC14', [('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC14({_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC14) and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC16(False, False, False, int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D5_DC16)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D5_DC17)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D5_DC18)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D5_DC15)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D5_DC19)

class D5_DC16(D5, NamedTuple('DC16', [('cf29', Any), ('cf30', Any), ('cf31', Any), ('cf32', Any), ('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC16({_dafny.string_of(self.cf29)}, {_dafny.string_of(self.cf30)}, {_dafny.string_of(self.cf31)}, {_dafny.string_of(self.cf32)}, {_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC16) and self.cf29 == __o.cf29 and self.cf30 == __o.cf30 and self.cf31 == __o.cf31 and self.cf32 == __o.cf32 and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC17(D5, NamedTuple('DC17', [('cf34', Any), ('cf35', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC17({_dafny.string_of(self.cf34)}, {_dafny.string_of(self.cf35)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC17) and self.cf34 == __o.cf34 and self.cf35 == __o.cf35
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC18(D5, NamedTuple('DC18', [('cf36', Any), ('cf37', Any), ('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC18({_dafny.string_of(self.cf36)}, {_dafny.string_of(self.cf37)}, {_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC18) and self.cf36 == __o.cf36 and self.cf37 == __o.cf37 and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC15(D5, NamedTuple('DC15', [('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC15({_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC15) and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC19(D5, NamedTuple('DC19', [('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC19({_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC19) and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC21(int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D6_DC21)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D6_DC22)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D6_DC20)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D6_DC23)

class D6_DC21(D6, NamedTuple('DC21', [('cf41', Any), ('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC21({_dafny.string_of(self.cf41)}, {_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC21) and self.cf41 == __o.cf41 and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC22(D6, NamedTuple('DC22', [])):
    def __dafnystr__(self) -> str:
        return f'D6.DC22'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC22)
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC20(D6, NamedTuple('DC20', [('cf40', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC20({_dafny.string_of(self.cf40)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC20) and self.cf40 == __o.cf40
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC23(D6, NamedTuple('DC23', [('cf43', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC23({_dafny.string_of(self.cf43)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC23) and self.cf43 == __o.cf43
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC25(False, False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D7_DC25)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D7_DC26)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D7_DC27)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D7_DC24)

class D7_DC25(D7, NamedTuple('DC25', [('cf45', Any), ('cf46', Any), ('cf47', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC25({_dafny.string_of(self.cf45)}, {_dafny.string_of(self.cf46)}, {_dafny.string_of(self.cf47)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC25) and self.cf45 == __o.cf45 and self.cf46 == __o.cf46 and self.cf47 == __o.cf47
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC26(D7, NamedTuple('DC26', [('cf48', Any), ('cf49', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC26({_dafny.string_of(self.cf48)}, {_dafny.string_of(self.cf49)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC26) and self.cf48 == __o.cf48 and self.cf49 == __o.cf49
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC27(D7, NamedTuple('DC27', [('cf50', Any), ('cf51', Any), ('cf52', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC27({_dafny.string_of(self.cf50)}, {_dafny.string_of(self.cf51)}, {_dafny.string_of(self.cf52)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC27) and self.cf50 == __o.cf50 and self.cf51 == __o.cf51 and self.cf52 == __o.cf52
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC24(D7, NamedTuple('DC24', [('cf44', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC24({_dafny.string_of(self.cf44)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC24) and self.cf44 == __o.cf44
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC29()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D8_DC29)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D8_DC30)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D8_DC28)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D8_DC31)

class D8_DC29(D8, NamedTuple('DC29', [])):
    def __dafnystr__(self) -> str:
        return f'D8.DC29'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC29)
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC30(D8, NamedTuple('DC30', [('cf54', Any), ('cf55', Any), ('cf56', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC30({_dafny.string_of(self.cf54)}, {_dafny.string_of(self.cf55)}, {_dafny.string_of(self.cf56)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC30) and self.cf54 == __o.cf54 and self.cf55 == __o.cf55 and self.cf56 == __o.cf56
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC28(D8, NamedTuple('DC28', [('cf53', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC28({_dafny.string_of(self.cf53)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC28) and self.cf53 == __o.cf53
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC31(D8, NamedTuple('DC31', [('cf57', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC31({_dafny.string_of(self.cf57)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC31) and self.cf57 == __o.cf57
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D9_DC32)

class D9_DC32(D9, NamedTuple('DC32', [('cf58', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC32({_dafny.string_of(self.cf58)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC32) and self.cf58 == __o.cf58
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D10_DC33)

class D10_DC33(D10, NamedTuple('DC33', [('cf59', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC33({_dafny.string_of(self.cf59)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC33) and self.cf59 == __o.cf59
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC35(int(0), False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D11_DC35)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D11_DC34)

class D11_DC35(D11, NamedTuple('DC35', [('cf61', Any), ('cf62', Any), ('cf63', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC35({_dafny.string_of(self.cf61)}, {_dafny.string_of(self.cf62)}, {_dafny.string_of(self.cf63)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC35) and self.cf61 == __o.cf61 and self.cf62 == __o.cf62 and self.cf63 == __o.cf63
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC34(D11, NamedTuple('DC34', [('cf60', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC34({_dafny.string_of(self.cf60)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC34) and self.cf60 == __o.cf60
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC37(False, False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D12_DC37)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D12_DC38)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D12_DC36)

class D12_DC37(D12, NamedTuple('DC37', [('cf65', Any), ('cf66', Any), ('cf67', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC37({_dafny.string_of(self.cf65)}, {_dafny.string_of(self.cf66)}, {_dafny.string_of(self.cf67)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC37) and self.cf65 == __o.cf65 and self.cf66 == __o.cf66 and self.cf67 == __o.cf67
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC38(D12, NamedTuple('DC38', [('cf68', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC38({_dafny.string_of(self.cf68)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC38) and self.cf68 == __o.cf68
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC36(D12, NamedTuple('DC36', [('cf64', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC36({_dafny.string_of(self.cf64)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC36) and self.cf64 == __o.cf64
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: D13_DC40(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D13_DC40)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D13_DC39)

class D13_DC40(D13, NamedTuple('DC40', [('cf70', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC40({_dafny.string_of(self.cf70)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC40) and self.cf70 == __o.cf70
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC39(D13, NamedTuple('DC39', [('cf69', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC39({_dafny.string_of(self.cf69)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC39) and self.cf69 == __o.cf69
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: D14_DC42(int(0), False, _dafny.Set({}), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC42(self) -> bool:
        return isinstance(self, D14_DC42)
    @property
    def is_DC41(self) -> bool:
        return isinstance(self, D14_DC41)

class D14_DC42(D14, NamedTuple('DC42', [('cf72', Any), ('cf73', Any), ('cf74', Any), ('cf75', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC42({_dafny.string_of(self.cf72)}, {_dafny.string_of(self.cf73)}, {_dafny.string_of(self.cf74)}, {_dafny.string_of(self.cf75)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC42) and self.cf72 == __o.cf72 and self.cf73 == __o.cf73 and self.cf74 == __o.cf74 and self.cf75 == __o.cf75
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC41(D14, NamedTuple('DC41', [('cf71', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC41({_dafny.string_of(self.cf71)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC41) and self.cf71 == __o.cf71
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC43(self) -> bool:
        return isinstance(self, D15_DC43)

class D15_DC43(D15, NamedTuple('DC43', [('cf76', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC43({_dafny.string_of(self.cf76)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC43) and self.cf76 == __o.cf76
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC44(self) -> bool:
        return isinstance(self, D16_DC44)

class D16_DC44(D16, NamedTuple('DC44', [('cf77', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC44({_dafny.string_of(self.cf77)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC44) and self.cf77 == __o.cf77
    def __hash__(self) -> int:
        return super().__hash__()


class D17:
    @classmethod
    def default(cls, ):
        return lambda: D17_DC46(False, int(0), int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC46(self) -> bool:
        return isinstance(self, D17_DC46)
    @property
    def is_DC45(self) -> bool:
        return isinstance(self, D17_DC45)

class D17_DC46(D17, NamedTuple('DC46', [('cf79', Any), ('cf80', Any), ('cf81', Any), ('cf82', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC46({_dafny.string_of(self.cf79)}, {_dafny.string_of(self.cf80)}, {_dafny.string_of(self.cf81)}, {_dafny.string_of(self.cf82)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC46) and self.cf79 == __o.cf79 and self.cf80 == __o.cf80 and self.cf81 == __o.cf81 and self.cf82 == __o.cf82
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC45(D17, NamedTuple('DC45', [('cf78', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC45({_dafny.string_of(self.cf78)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC45) and self.cf78 == __o.cf78
    def __hash__(self) -> int:
        return super().__hash__()


class D18:
    @classmethod
    def default(cls, ):
        return lambda: D18_DC48(False, False, D1.default()(), None)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC48(self) -> bool:
        return isinstance(self, D18_DC48)
    @property
    def is_DC47(self) -> bool:
        return isinstance(self, D18_DC47)
    @property
    def is_DC49(self) -> bool:
        return isinstance(self, D18_DC49)

class D18_DC48(D18, NamedTuple('DC48', [('cf84', Any), ('cf85', Any), ('cf86', Any), ('cf87', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC48({_dafny.string_of(self.cf84)}, {_dafny.string_of(self.cf85)}, {_dafny.string_of(self.cf86)}, {_dafny.string_of(self.cf87)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC48) and self.cf84 == __o.cf84 and self.cf85 == __o.cf85 and self.cf86 == __o.cf86 and self.cf87 == __o.cf87
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC47(D18, NamedTuple('DC47', [('cf83', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC47({_dafny.string_of(self.cf83)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC47) and self.cf83 == __o.cf83
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC49(D18, NamedTuple('DC49', [('cf88', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC49({_dafny.string_of(self.cf88)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC49) and self.cf88 == __o.cf88
    def __hash__(self) -> int:
        return super().__hash__()


class D19:
    @classmethod
    def default(cls, ):
        return lambda: D19_DC51(int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC51(self) -> bool:
        return isinstance(self, D19_DC51)
    @property
    def is_DC52(self) -> bool:
        return isinstance(self, D19_DC52)
    @property
    def is_DC50(self) -> bool:
        return isinstance(self, D19_DC50)

class D19_DC51(D19, NamedTuple('DC51', [('cf90', Any), ('cf91', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC51({_dafny.string_of(self.cf90)}, {_dafny.string_of(self.cf91)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC51) and self.cf90 == __o.cf90 and self.cf91 == __o.cf91
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC52(D19, NamedTuple('DC52', [('cf92', Any), ('cf93', Any), ('cf94', Any), ('cf95', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC52({_dafny.string_of(self.cf92)}, {_dafny.string_of(self.cf93)}, {_dafny.string_of(self.cf94)}, {_dafny.string_of(self.cf95)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC52) and self.cf92 == __o.cf92 and self.cf93 == __o.cf93 and self.cf94 == __o.cf94 and self.cf95 == __o.cf95
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC50(D19, NamedTuple('DC50', [('cf89', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC50({_dafny.string_of(self.cf89)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC50) and self.cf89 == __o.cf89
    def __hash__(self) -> int:
        return super().__hash__()


class D20:
    @classmethod
    def default(cls, ):
        return lambda: D20_DC54(False, False, int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC54(self) -> bool:
        return isinstance(self, D20_DC54)
    @property
    def is_DC55(self) -> bool:
        return isinstance(self, D20_DC55)
    @property
    def is_DC53(self) -> bool:
        return isinstance(self, D20_DC53)
    @property
    def is_DC56(self) -> bool:
        return isinstance(self, D20_DC56)

class D20_DC54(D20, NamedTuple('DC54', [('cf97', Any), ('cf98', Any), ('cf99', Any), ('cf100', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC54({_dafny.string_of(self.cf97)}, {_dafny.string_of(self.cf98)}, {_dafny.string_of(self.cf99)}, {_dafny.string_of(self.cf100)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC54) and self.cf97 == __o.cf97 and self.cf98 == __o.cf98 and self.cf99 == __o.cf99 and self.cf100 == __o.cf100
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC55(D20, NamedTuple('DC55', [('cf101', Any), ('cf102', Any), ('cf103', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC55({self.cf101.VerbatimString(True)}, {_dafny.string_of(self.cf102)}, {_dafny.string_of(self.cf103)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC55) and self.cf101 == __o.cf101 and self.cf102 == __o.cf102 and self.cf103 == __o.cf103
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC53(D20, NamedTuple('DC53', [('cf96', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC53({_dafny.string_of(self.cf96)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC53) and self.cf96 == __o.cf96
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC56(D20, NamedTuple('DC56', [('cf104', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC56({_dafny.string_of(self.cf104)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC56) and self.cf104 == __o.cf104
    def __hash__(self) -> int:
        return super().__hash__()


class D21:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC57(self) -> bool:
        return isinstance(self, D21_DC57)

class D21_DC57(D21, NamedTuple('DC57', [('cf105', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC57({_dafny.string_of(self.cf105)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC57) and self.cf105 == __o.cf105
    def __hash__(self) -> int:
        return super().__hash__()


class D22:
    @classmethod
    def default(cls, ):
        return lambda: D22_DC59(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC59(self) -> bool:
        return isinstance(self, D22_DC59)
    @property
    def is_DC58(self) -> bool:
        return isinstance(self, D22_DC58)

class D22_DC59(D22, NamedTuple('DC59', [('cf107', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC59({_dafny.string_of(self.cf107)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC59) and self.cf107 == __o.cf107
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC58(D22, NamedTuple('DC58', [('cf106', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC58({_dafny.string_of(self.cf106)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC58) and self.cf106 == __o.cf106
    def __hash__(self) -> int:
        return super().__hash__()


class D23:
    @classmethod
    def default(cls, ):
        return lambda: D23_DC61(_dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC61(self) -> bool:
        return isinstance(self, D23_DC61)
    @property
    def is_DC60(self) -> bool:
        return isinstance(self, D23_DC60)
    @property
    def is_DC62(self) -> bool:
        return isinstance(self, D23_DC62)

class D23_DC61(D23, NamedTuple('DC61', [('cf109', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC61({_dafny.string_of(self.cf109)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC61) and self.cf109 == __o.cf109
    def __hash__(self) -> int:
        return super().__hash__()

class D23_DC60(D23, NamedTuple('DC60', [('cf108', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC60({_dafny.string_of(self.cf108)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC60) and self.cf108 == __o.cf108
    def __hash__(self) -> int:
        return super().__hash__()

class D23_DC62(D23, NamedTuple('DC62', [('cf110', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC62({_dafny.string_of(self.cf110)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC62) and self.cf110 == __o.cf110
    def __hash__(self) -> int:
        return super().__hash__()


class D24:
    @classmethod
    def default(cls, ):
        return lambda: D24_DC64()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC64(self) -> bool:
        return isinstance(self, D24_DC64)
    @property
    def is_DC63(self) -> bool:
        return isinstance(self, D24_DC63)

class D24_DC64(D24, NamedTuple('DC64', [])):
    def __dafnystr__(self) -> str:
        return f'D24.DC64'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC64)
    def __hash__(self) -> int:
        return super().__hash__()

class D24_DC63(D24, NamedTuple('DC63', [('cf111', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC63({_dafny.string_of(self.cf111)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC63) and self.cf111 == __o.cf111
    def __hash__(self) -> int:
        return super().__hash__()


class D25:
    @classmethod
    def default(cls, ):
        return lambda: D25_DC66()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC66(self) -> bool:
        return isinstance(self, D25_DC66)
    @property
    def is_DC65(self) -> bool:
        return isinstance(self, D25_DC65)
    @property
    def is_DC67(self) -> bool:
        return isinstance(self, D25_DC67)

class D25_DC66(D25, NamedTuple('DC66', [])):
    def __dafnystr__(self) -> str:
        return f'D25.DC66'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC66)
    def __hash__(self) -> int:
        return super().__hash__()

class D25_DC65(D25, NamedTuple('DC65', [('cf112', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC65({_dafny.string_of(self.cf112)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC65) and self.cf112 == __o.cf112
    def __hash__(self) -> int:
        return super().__hash__()

class D25_DC67(D25, NamedTuple('DC67', [('cf113', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC67({_dafny.string_of(self.cf113)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC67) and self.cf113 == __o.cf113
    def __hash__(self) -> int:
        return super().__hash__()


class D26:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC68(self) -> bool:
        return isinstance(self, D26_DC68)

class D26_DC68(D26, NamedTuple('DC68', [('cf114', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC68({_dafny.string_of(self.cf114)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC68) and self.cf114 == __o.cf114
    def __hash__(self) -> int:
        return super().__hash__()


class D27:
    @classmethod
    def default(cls, ):
        return lambda: None
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC69(self) -> bool:
        return isinstance(self, D27_DC69)

class D27_DC69(D27, NamedTuple('DC69', [('cf115', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC69({_dafny.string_of(self.cf115)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC69) and self.cf115 == __o.cf115
    def __hash__(self) -> int:
        return super().__hash__()


class D28:
    @classmethod
    def default(cls, ):
        return lambda: D28_DC71(False, int(0), _dafny.Seq({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC71(self) -> bool:
        return isinstance(self, D28_DC71)
    @property
    def is_DC72(self) -> bool:
        return isinstance(self, D28_DC72)
    @property
    def is_DC70(self) -> bool:
        return isinstance(self, D28_DC70)
    @property
    def is_DC73(self) -> bool:
        return isinstance(self, D28_DC73)

class D28_DC71(D28, NamedTuple('DC71', [('cf117', Any), ('cf118', Any), ('cf119', Any)])):
    def __dafnystr__(self) -> str:
        return f'D28.DC71({_dafny.string_of(self.cf117)}, {_dafny.string_of(self.cf118)}, {_dafny.string_of(self.cf119)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D28_DC71) and self.cf117 == __o.cf117 and self.cf118 == __o.cf118 and self.cf119 == __o.cf119
    def __hash__(self) -> int:
        return super().__hash__()

class D28_DC72(D28, NamedTuple('DC72', [])):
    def __dafnystr__(self) -> str:
        return f'D28.DC72'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D28_DC72)
    def __hash__(self) -> int:
        return super().__hash__()

class D28_DC70(D28, NamedTuple('DC70', [('cf116', Any)])):
    def __dafnystr__(self) -> str:
        return f'D28.DC70({_dafny.string_of(self.cf116)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D28_DC70) and self.cf116 == __o.cf116
    def __hash__(self) -> int:
        return super().__hash__()

class D28_DC73(D28, NamedTuple('DC73', [('cf120', Any)])):
    def __dafnystr__(self) -> str:
        return f'D28.DC73({_dafny.string_of(self.cf120)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D28_DC73) and self.cf120 == __o.cf120
    def __hash__(self) -> int:
        return super().__hash__()


class D29:
    @classmethod
    def default(cls, ):
        return lambda: D29_DC75(_dafny.Map({}), False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC75(self) -> bool:
        return isinstance(self, D29_DC75)
    @property
    def is_DC76(self) -> bool:
        return isinstance(self, D29_DC76)
    @property
    def is_DC77(self) -> bool:
        return isinstance(self, D29_DC77)
    @property
    def is_DC74(self) -> bool:
        return isinstance(self, D29_DC74)
    @property
    def is_DC78(self) -> bool:
        return isinstance(self, D29_DC78)

class D29_DC75(D29, NamedTuple('DC75', [('cf122', Any), ('cf123', Any), ('cf124', Any)])):
    def __dafnystr__(self) -> str:
        return f'D29.DC75({_dafny.string_of(self.cf122)}, {_dafny.string_of(self.cf123)}, {_dafny.string_of(self.cf124)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D29_DC75) and self.cf122 == __o.cf122 and self.cf123 == __o.cf123 and self.cf124 == __o.cf124
    def __hash__(self) -> int:
        return super().__hash__()

class D29_DC76(D29, NamedTuple('DC76', [('cf125', Any), ('cf126', Any), ('cf127', Any), ('cf128', Any)])):
    def __dafnystr__(self) -> str:
        return f'D29.DC76({_dafny.string_of(self.cf125)}, {_dafny.string_of(self.cf126)}, {_dafny.string_of(self.cf127)}, {_dafny.string_of(self.cf128)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D29_DC76) and self.cf125 == __o.cf125 and self.cf126 == __o.cf126 and self.cf127 == __o.cf127 and self.cf128 == __o.cf128
    def __hash__(self) -> int:
        return super().__hash__()

class D29_DC77(D29, NamedTuple('DC77', [('cf129', Any), ('cf130', Any)])):
    def __dafnystr__(self) -> str:
        return f'D29.DC77({_dafny.string_of(self.cf129)}, {_dafny.string_of(self.cf130)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D29_DC77) and self.cf129 == __o.cf129 and self.cf130 == __o.cf130
    def __hash__(self) -> int:
        return super().__hash__()

class D29_DC74(D29, NamedTuple('DC74', [('cf121', Any)])):
    def __dafnystr__(self) -> str:
        return f'D29.DC74({_dafny.string_of(self.cf121)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D29_DC74) and self.cf121 == __o.cf121
    def __hash__(self) -> int:
        return super().__hash__()

class D29_DC78(D29, NamedTuple('DC78', [('cf131', Any)])):
    def __dafnystr__(self) -> str:
        return f'D29.DC78({_dafny.string_of(self.cf131)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D29_DC78) and self.cf131 == __o.cf131
    def __hash__(self) -> int:
        return super().__hash__()


class D30:
    @classmethod
    def default(cls, ):
        return lambda: D30_DC80(int(0), _dafny.CodePoint('D'), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC80(self) -> bool:
        return isinstance(self, D30_DC80)
    @property
    def is_DC81(self) -> bool:
        return isinstance(self, D30_DC81)
    @property
    def is_DC82(self) -> bool:
        return isinstance(self, D30_DC82)
    @property
    def is_DC79(self) -> bool:
        return isinstance(self, D30_DC79)

class D30_DC80(D30, NamedTuple('DC80', [('cf133', Any), ('cf134', Any), ('cf135', Any)])):
    def __dafnystr__(self) -> str:
        return f'D30.DC80({_dafny.string_of(self.cf133)}, {_dafny.string_of(self.cf134)}, {_dafny.string_of(self.cf135)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D30_DC80) and self.cf133 == __o.cf133 and self.cf134 == __o.cf134 and self.cf135 == __o.cf135
    def __hash__(self) -> int:
        return super().__hash__()

class D30_DC81(D30, NamedTuple('DC81', [('cf136', Any), ('cf137', Any)])):
    def __dafnystr__(self) -> str:
        return f'D30.DC81({self.cf136.VerbatimString(True)}, {_dafny.string_of(self.cf137)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D30_DC81) and self.cf136 == __o.cf136 and self.cf137 == __o.cf137
    def __hash__(self) -> int:
        return super().__hash__()

class D30_DC82(D30, NamedTuple('DC82', [('cf138', Any)])):
    def __dafnystr__(self) -> str:
        return f'D30.DC82({_dafny.string_of(self.cf138)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D30_DC82) and self.cf138 == __o.cf138
    def __hash__(self) -> int:
        return super().__hash__()

class D30_DC79(D30, NamedTuple('DC79', [('cf132', Any)])):
    def __dafnystr__(self) -> str:
        return f'D30.DC79({_dafny.string_of(self.cf132)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D30_DC79) and self.cf132 == __o.cf132
    def __hash__(self) -> int:
        return super().__hash__()


class D31:
    @classmethod
    def default(cls, ):
        return lambda: D31_DC84(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC84(self) -> bool:
        return isinstance(self, D31_DC84)
    @property
    def is_DC85(self) -> bool:
        return isinstance(self, D31_DC85)
    @property
    def is_DC83(self) -> bool:
        return isinstance(self, D31_DC83)

class D31_DC84(D31, NamedTuple('DC84', [('cf140', Any)])):
    def __dafnystr__(self) -> str:
        return f'D31.DC84({_dafny.string_of(self.cf140)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D31_DC84) and self.cf140 == __o.cf140
    def __hash__(self) -> int:
        return super().__hash__()

class D31_DC85(D31, NamedTuple('DC85', [('cf141', Any), ('cf142', Any), ('cf143', Any), ('cf144', Any), ('cf145', Any)])):
    def __dafnystr__(self) -> str:
        return f'D31.DC85({_dafny.string_of(self.cf141)}, {_dafny.string_of(self.cf142)}, {_dafny.string_of(self.cf143)}, {_dafny.string_of(self.cf144)}, {_dafny.string_of(self.cf145)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D31_DC85) and self.cf141 == __o.cf141 and self.cf142 == __o.cf142 and self.cf143 == __o.cf143 and self.cf144 == __o.cf144 and self.cf145 == __o.cf145
    def __hash__(self) -> int:
        return super().__hash__()

class D31_DC83(D31, NamedTuple('DC83', [('cf139', Any)])):
    def __dafnystr__(self) -> str:
        return f'D31.DC83({_dafny.string_of(self.cf139)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D31_DC83) and self.cf139 == __o.cf139
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    @property
    def f4(self):
        return self._f4
    @f4.setter
    def f4(self, value):
        self._f4 = value
    def m3(self, p0, p1, p2, p3, globalState):
        pass

    def m4(self, p0, p1, globalState):
        pass


class T1:
    pass
    def m7(self, p0, p1, p2, p3, globalState):
        pass

    def m8(self, p0, globalState):
        pass


class T2:
    pass
    def m11(self, globalState):
        pass

    def m12(self, p0, p1, p2, p3, globalState):
        pass


class GlobalState:
    def  __init__(self):
        self._f0: bool = False
        self._f1: int = int(0)
        self._f2: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2):
        (self)._f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2

    @property
    def f0(self):
        return self._f0
    @property
    def f1(self):
        return self._f1
    @property
    def f2(self):
        return self._f2

class C0:
    def  __init__(self):
        self.f18: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f18):
        (self).f18 = f18

    def fm15(self, p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference([self.f18, (0) - (self.f18)])) + ((_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([False])).cardinality for d_280_i0_ in range(default__.abs(926))])) + (_dafny.SeqWithoutIsStrInference([(0) - (-413) for d_281_i1_ in range(default__.abs(-543))])))


class C1(T1, T0, T2):
    def  __init__(self):
        self._f4: int = int(0)
        self._f15: int = int(0)
        self._f16: bool = False
        self.f17: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    @property
    def f4(self):
        return self._f4
    @f4.setter
    def f4(self, value):
        self._f4 = value
    @property
    def f15(self):
        return self._f15
    @property
    def f16(self):
        return self._f16
    def ctor__(self, f17, f4, f15, f16):
        (self).f17 = f17
        (self).f4 = f4
        (self)._f15 = f15
        (self)._f16 = f16

    def fm10(self, p0, p1, p2, p3, globalState):
        if (self).f16:
            return _dafny.CodePoint('h')
        elif True:
            return _dafny.CodePoint('j')

    def fm5(self, p0, p1, globalState):
        def iife53_():
            coll49_ = _dafny.Set()
            compr_49_: int
            for compr_49_ in _dafny.IntegerRange(755, 736):
                d_282_v0_: int = compr_49_
                if ((755) <= (d_282_v0_)) and ((d_282_v0_) < (736)):
                    coll49_ = coll49_.union(_dafny.Set([default__.safeModuloInt(d_282_v0_, (0) - (self.f4))]))
            return _dafny.Set(coll49_)
        return (_dafny.Set({self.f17, self.f17})).ispropersubset((_dafny.Set({self.f4})).intersection(iife53_()
        ))

    def m7(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: _dafny.Array = _dafny.Array(None, 0)
        r3: int = int(0)
        d_283_i0_: int
        d_283_i0_ = 0
        with _dafny.label("2"):
            while (self).f16:
                with _dafny.c_label("2"):
                    if (d_283_i0_) >= (100):
                        raise _dafny.Break("2")
                    d_283_i0_ = (d_283_i0_) + (1)
                    d_284_v0_: _dafny.Array
                    nw33_ = _dafny.Array(int(0), 1)
                    d_284_v0_ = nw33_
                    index15_ = default__.safeIndex(253, (d_284_v0_).length(0))
                    (d_284_v0_)[index15_] = ((self).f15) + ((self).f15)
                    index16_ = default__.safeIndex(253, (d_284_v0_).length(0))
                    (d_284_v0_)[index16_] = default__.fm1(globalState)
                    d_285_v1_: C0
                    nw34_ = C0()
                    nw34_.ctor__((len(default__.fm16((self).f16, _dafny.SeqWithoutIsStrInference([p1, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "duuor"))]), not(p2), 466, globalState))) + (default__.fm1(globalState)))
                    d_285_v1_ = nw34_
                    r2 = d_284_v0_
                    d_286_v2_: bool
                    d_286_v2_ = True
                    d_286_v2_ = (self).fm5(p1, d_286_v2_, globalState)
                    pass
            pass
        d_287_v3_: bool
        d_287_v3_ = True
        d_287_v3_ = (self).f16
        d_288_v4_: _dafny.MultiSet
        d_288_v4_ = _dafny.MultiSet([d_287_v3_])
        d_289_v5_: D0
        d_289_v5_ = D0_DC1(p0, self.f4, p2)
        hi1_ = ((d_288_v4_)[True] if (True) in (d_288_v4_) else (d_289_v5_).cf1)
        for d_290_i1_ in range(p0, hi1_):
            if (True if not(not (d_287_v3_) or (d_287_v3_)) else (d_287_v3_) and (p3)):
                d_291_v6_: _dafny.Seq
                d_291_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kosv"))
                d_291_v6_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qgeknb"))) + (d_291_v6_)
                d_292_v7_: C0
                nw35_ = C0()
                nw35_.ctor__(-745)
                d_292_v7_ = nw35_
                d_293_v8_: C0
                nw36_ = C0()
                nw36_.ctor__(default__.fm1(globalState))
                d_293_v8_ = nw36_
                d_294_v9_: _dafny.Set
                d_294_v9_ = _dafny.Set({(self).f15})
                d_295_v10_: _dafny.Array
                nw37_ = _dafny.Array(False, 2)
                d_295_v10_ = nw37_
                index17_ = default__.safeIndex(21, (d_295_v10_).length(0))
                (d_295_v10_)[index17_] = p2
                d_296_v11_: _dafny.MultiSet
                d_296_v11_ = _dafny.MultiSet([p0])
                d_297_v12_: _dafny.Seq
                d_297_v12_ = _dafny.SeqWithoutIsStrInference([p1, d_291_v6_])
                index18_ = default__.safeIndex(21, (d_295_v10_).length(0))
                rhs18_ = (d_296_v11_).isdisjoint(d_296_v11_)
                rhs19_ = default__.fm16((855) < (len(p1)), d_297_v12_, p3, d_292_v7_.f18, globalState)
                rhs20_ = ((p0) == (d_292_v7_.f18)) == (p3)
                lhs6_ = d_295_v10_
                lhs7_ = default__.safeIndex(21, (d_295_v10_).length(0))
                d_287_v3_ = rhs18_
                d_294_v9_ = rhs19_
                lhs6_[lhs7_] = rhs20_
                d_298_v13_: _dafny.Seq
                d_298_v13_ = _dafny.SeqWithoutIsStrInference([d_295_v10_])
                d_298_v13_ = _dafny.SeqWithoutIsStrInference([d_295_v10_, d_295_v10_])
            elif True:
                d_299_v14_: _dafny.MultiSet
                d_299_v14_ = _dafny.MultiSet([self.f17])
                d_300_v15_: _dafny.Seq
                d_300_v15_ = _dafny.SeqWithoutIsStrInference([self.f4])
                d_301_v16_: _dafny.Array
                nw38_ = _dafny.Array(None, 18)
                nw38_[int(0)] = default__.safeDivisionInt((self).f15, self.f4)
                nw38_[int(1)] = d_290_i1_
                nw38_[int(2)] = (d_299_v14_).cardinality
                nw38_[int(3)] = p0
                nw38_[int(4)] = (0) - (len(p1))
                nw38_[int(5)] = self.f4
                nw38_[int(6)] = d_290_i1_
                nw38_[int(7)] = self.f17
                nw38_[int(8)] = (self).f15
                nw38_[int(9)] = self.f17
                nw38_[int(10)] = 64
                nw38_[int(11)] = (0) - (d_290_i1_)
                nw38_[int(12)] = self.f17
                nw38_[int(13)] = d_290_i1_
                nw38_[int(14)] = default__.safeDivisionInt((d_300_v15_)[default__.safeIndex(self.f17, len(d_300_v15_))], (self).f15)
                nw38_[int(15)] = (p0) * (534)
                nw38_[int(16)] = d_290_i1_
                nw38_[int(17)] = (0) - (len((p1) + (p1)))
                d_301_v16_ = nw38_
                index19_ = default__.safeIndex(235, (d_301_v16_).length(0))
                (d_301_v16_)[index19_] = p0
                d_302_v17_: _dafny.Map
                d_302_v17_ = _dafny.Map({len(_dafny.Map({self.f4: _dafny.Set({((d_299_v14_)[(self).f15] if ((self).f15) in (d_299_v14_) else d_290_i1_), -617, self.f17})})): (d_290_i1_) * ((self).f15)})
                index20_ = default__.safeIndex(235, (d_301_v16_).length(0))
                rhs21_ = default__.safeDivisionInt(((self).f15) - ((d_299_v14_).cardinality), len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([p1 for d_303_i3_ in range(default__.abs(418))])) for d_304_i2_ in range(default__.abs(780))])))
                rhs22_ = len(d_302_v17_)
                rhs23_ = p0
                lhs8_ = self
                lhs9_ = d_301_v16_
                lhs10_ = default__.safeIndex(235, (d_301_v16_).length(0))
                lhs8_.f4 = rhs21_
                r3 = rhs22_
                lhs9_[lhs10_] = rhs23_
                d_305_v18_: _dafny.Map
                d_305_v18_ = _dafny.Map({8: p3})
                def iife54_():
                    coll50_ = _dafny.Map()
                    compr_50_: int
                    for compr_50_ in _dafny.IntegerRange(-652, -6):
                        d_306_v19_: int = compr_50_
                        if ((-652) <= (d_306_v19_)) and ((d_306_v19_) < (-6)):
                            coll50_[default__.safeModuloInt(d_306_v19_, len(d_300_v15_))] = (p1) != (p1)
                    return _dafny.Map(coll50_)
                d_305_v18_ = iife54_()
                
                r3 = p0
                d_307_v20_: _dafny.Set
                d_307_v20_ = _dafny.Set({972})
                d_308_v21_: D3
                d_308_v21_ = D3_DC7(d_307_v20_)
                d_287_v3_ = ((d_307_v20_).intersection((d_308_v21_).cf12)).isdisjoint(d_307_v20_)
                index21_ = default__.safeIndex(235, (d_301_v16_).length(0))
                (d_301_v16_)[index21_] = d_290_i1_
            rhs24_ = (self).f15
            rhs25_ = d_287_v3_
            rhs26_ = self.f4
            r3 = rhs24_
            d_287_v3_ = rhs25_
            r0 = rhs26_
            r0 = p0
            d_309_v22_: _dafny.Array
            def lambda6_(d_310_i4_):
                return (self).f16

            init2_ = lambda6_
            nw39_ = _dafny.Array(None, 10)
            for i0_2_ in range(nw39_.length(0)):
                nw39_[i0_2_] = init2_(i0_2_)
            d_309_v22_ = nw39_
            index22_ = default__.safeIndex(468, (d_309_v22_).length(0))
            (d_309_v22_)[index22_] = not ((self).f16) or ((self).f16)
            d_311_v23_: _dafny.Seq
            d_311_v23_ = _dafny.SeqWithoutIsStrInference([self.f17])
            index23_ = default__.safeIndex(468, (d_309_v22_).length(0))
            (d_309_v22_)[index23_] = ((p0) * (len(p1))) < (((0) - (-218)) - (len(d_311_v23_)))
        d_312_v24_: str
        d_312_v24_ = _dafny.CodePoint('g')
        d_313_v25_: _dafny.Seq
        d_313_v25_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i'), d_312_v24_])])
        d_314_v26_: _dafny.Set
        d_314_v26_ = _dafny.Set({(d_312_v24_) not in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dhrfx"))), (self).f16, (d_313_v25_) <= ((d_313_v25_).set(default__.safeIndex(-674, len(d_313_v25_)), p1))})
        d_314_v26_ = (d_314_v26_) | (default__.fm17(p3, default__.fm1(globalState), globalState))
        if p2:
            d_315_v27_: D0
            d_315_v27_ = D0_DC1(p0, (0) - ((self).f15), d_287_v3_)
            d_316_v28_: D0
            d_316_v28_ = D0_DC2(D0_DC2(d_315_v27_))
            d_317_v29_: _dafny.Map
            d_317_v29_ = _dafny.Map({d_287_v3_: (self).f16})
            d_318_v30_: _dafny.MultiSet
            d_318_v30_ = _dafny.MultiSet([d_312_v24_])
            d_319_v31_: _dafny.Seq
            d_319_v31_ = _dafny.SeqWithoutIsStrInference([default__.fm0((0) - ((self).f15), self.f4, globalState)])
            d_320_v32_: _dafny.Map
            d_320_v32_ = _dafny.Map({len(d_319_v31_): d_287_v3_})
            d_321_v33_: _dafny.Array
            nw40_ = _dafny.Array(None, 28)
            nw40_[int(0)] = len(_dafny.Set({d_316_v28_, d_316_v28_}))
            nw40_[int(1)] = -980
            nw40_[int(2)] = self.f17
            nw40_[int(3)] = (self).f15
            nw40_[int(4)] = len(p1)
            nw40_[int(5)] = len(p1)
            nw40_[int(6)] = p0
            nw40_[int(7)] = len(d_314_v26_)
            nw40_[int(8)] = self.f4
            nw40_[int(9)] = (self).f15
            nw40_[int(10)] = p0
            nw40_[int(11)] = ((d_288_v4_).set(p2, default__.abs(len(d_317_v29_)))).cardinality
            nw40_[int(12)] = self.f17
            nw40_[int(13)] = (d_318_v30_).cardinality
            nw40_[int(14)] = self.f17
            nw40_[int(15)] = self.f4
            nw40_[int(16)] = (self).f15
            nw40_[int(17)] = self.f17
            nw40_[int(18)] = self.f4
            nw40_[int(19)] = p0
            nw40_[int(20)] = len((d_319_v31_).set(default__.safeIndex((0) - (len(p1)), len(d_319_v31_)), not(p2)))
            nw40_[int(21)] = default__.fm1(globalState)
            nw40_[int(22)] = self.f17
            nw40_[int(23)] = len(d_320_v32_)
            nw40_[int(24)] = (self).f15
            nw40_[int(25)] = self.f4
            nw40_[int(26)] = len(p1)
            nw40_[int(27)] = self.f17
            d_321_v33_ = nw40_
            d_322_v34_: _dafny.MultiSet
            d_322_v34_ = _dafny.MultiSet([d_321_v33_])
            d_287_v3_ = (d_322_v34_).issubset(d_322_v34_)
            (self).f4 = self.f4
            d_323_v35_: _dafny.Seq
            d_323_v35_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wy"))
            rhs27_ = p0
            rhs28_ = d_323_v35_
            lhs11_ = self
            lhs11_.f17 = rhs27_
            d_323_v35_ = rhs28_
            d_324_v36_: _dafny.MultiSet
            d_324_v36_ = _dafny.MultiSet([(self).f15])
            (self).f17 = ((d_324_v36_)[default__.safeModuloInt((self).f15, self.f17)] if (default__.safeModuloInt((self).f15, self.f17)) in (d_324_v36_) else (self).f15)
            d_325_v37_: _dafny.Array
            def lambda7_(d_326_v31_):
                def lambda8_(d_327_i5_):
                    return D1_DC4(self.f17, d_326_v31_, self.f17)

                return lambda8_

            init3_ = lambda7_(d_319_v31_)
            nw41_ = _dafny.Array(None, 8)
            for i0_3_ in range(nw41_.length(0)):
                nw41_[i0_3_] = init3_(i0_3_)
            d_325_v37_ = nw41_
            d_328_v38_: D1
            d_328_v38_ = D1_DC4(len(d_323_v35_), d_319_v31_, self.f4)
            index24_ = default__.safeIndex(360, (d_325_v37_).length(0))
            (d_325_v37_)[index24_] = d_328_v38_
            index25_ = default__.safeIndex(360, (d_325_v37_).length(0))
            (d_325_v37_)[index25_] = d_328_v38_
        elif True:
            d_329_v39_: _dafny.Array
            def lambda9_(d_330_i6_):
                return default__.safeDivisionInt(d_330_i6_, self.f4)

            init4_ = lambda9_
            nw42_ = _dafny.Array(None, 10)
            for i0_4_ in range(nw42_.length(0)):
                nw42_[i0_4_] = init4_(i0_4_)
            d_329_v39_ = nw42_
            d_331_v40_: _dafny.Seq
            d_331_v40_ = _dafny.SeqWithoutIsStrInference([d_329_v39_, d_329_v39_, d_329_v39_])
            d_332_v41_: _dafny.Seq
            d_332_v41_ = _dafny.SeqWithoutIsStrInference([(d_331_v40_)[default__.safeIndex(p0, len(d_331_v40_))]])
            d_332_v41_ = d_332_v41_
            d_333_v42_: _dafny.Array
            nw43_ = _dafny.Array(False, 19)
            d_333_v42_ = nw43_
            index26_ = default__.safeIndex(996, (d_333_v42_).length(0))
            (d_333_v42_)[index26_] = p2
            d_334_v43_: _dafny.MultiSet
            d_334_v43_ = _dafny.MultiSet([d_333_v42_])
            index27_ = default__.safeIndex(996, (d_333_v42_).length(0))
            rhs29_ = (_dafny.MultiSet([d_333_v42_])).ispropersubset(d_334_v43_)
            rhs30_ = d_314_v26_
            lhs12_ = d_333_v42_
            lhs13_ = default__.safeIndex(996, (d_333_v42_).length(0))
            lhs12_[lhs13_] = rhs29_
            d_314_v26_ = rhs30_
            d_335_v44_: _dafny.Array
            def lambda10_(d_336_i7_):
                return _dafny.CodePoint('m')

            init5_ = lambda10_
            nw44_ = _dafny.Array(None, 5)
            for i0_5_ in range(nw44_.length(0)):
                nw44_[i0_5_] = init5_(i0_5_)
            d_335_v44_ = nw44_
            index28_ = default__.safeIndex(906, (d_335_v44_).length(0))
            (d_335_v44_)[index28_] = d_312_v24_
            index29_ = default__.safeIndex(906, (d_335_v44_).length(0))
            (d_335_v44_)[index29_] = (p1)[default__.safeIndex((self).f15, len(p1))]
            d_337_v45_: _dafny.MultiSet
            d_337_v45_ = _dafny.MultiSet([d_329_v39_])
            d_337_v45_ = d_337_v45_
            d_329_v39_ = d_329_v39_
        d_338_v46_: _dafny.Seq
        d_338_v46_ = _dafny.SeqWithoutIsStrInference([d_287_v3_, (self).f16, d_287_v3_])
        (self).f4 = (len((d_338_v46_) + (d_338_v46_))) * ((self).f15)
        r0 = len(p1)
        d_339_v47_: _dafny.Array
        nw45_ = _dafny.Array(_dafny.Array(None, 0), 23)
        d_339_v47_ = nw45_
        r1 = d_339_v47_
        def lambda11_(d_340_p0_):
            def lambda12_(d_341_i8_):
                return (d_341_i8_) - (d_340_p0_)

            return lambda12_

        init6_ = lambda11_(p0)
        nw46_ = _dafny.Array(None, 18)
        for i0_6_ in range(nw46_.length(0)):
            nw46_[i0_6_] = init6_(i0_6_)
        r2 = nw46_
        r3 = (self).f15
        return r0, r1, r2, r3

    def m8(self, p0, globalState):
        d_342_v0_: _dafny.Seq
        d_342_v0_ = _dafny.SeqWithoutIsStrInference([False])
        d_343_v1_: D1
        d_343_v1_ = D1_DC4((0) - ((len(_dafny.SeqWithoutIsStrInference([(self).f15 for d_344_i0_ in range(default__.abs(-33))]))) * ((self).f15)), d_342_v0_, 873)
        d_343_v1_ = d_343_v1_
        hi2_ = self.f4
        for d_345_i1_ in range(len(default__.fm18(self.f17, (self).f15, self.f17, globalState)), hi2_):
            (self).f17 = self.f17
            d_346_v2_: _dafny.Seq
            d_347_v3_: _dafny.Map
            out10_: _dafny.Seq
            out11_: _dafny.Map
            out10_, out11_ = (self).m14(globalState)
            d_346_v2_ = out10_
            d_347_v3_ = out11_
            d_348_v4_: _dafny.Map
            d_348_v4_ = _dafny.Map({(self).f16: d_346_v2_})
            d_349_v5_: bool
            d_349_v5_ = False
            rhs31_ = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "onsmmxab")))
            rhs32_ = default__.fm19(_dafny.SeqWithoutIsStrInference([(self).f15 for d_350_i2_ in range(default__.abs(188))]), (self).f15, globalState)
            rhs33_ = (self).f16
            lhs14_ = self
            lhs14_.f4 = rhs31_
            d_348_v4_ = rhs32_
            d_349_v5_ = rhs33_
            (self).f17 = d_345_i1_
        d_351_v6_: _dafny.Map
        d_351_v6_ = _dafny.Map({self.f4: (self).f16})
        d_352_v7_: _dafny.Seq
        d_352_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mvm"))
        d_353_v8_: _dafny.Array
        nw47_ = _dafny.Array(None, 28)
        nw47_[int(0)] = (self).f15
        nw47_[int(1)] = (self).f15
        nw47_[int(2)] = (self).f15
        nw47_[int(3)] = self.f4
        nw47_[int(4)] = (self).f15
        nw47_[int(5)] = (0) - (len((d_352_v7_ if ((d_351_v6_)[self.f17] if (self.f17) in (d_351_v6_) else False) else d_352_v7_)))
        nw47_[int(6)] = self.f4
        nw47_[int(7)] = (369) - (self.f17)
        nw47_[int(8)] = self.f4
        nw47_[int(9)] = self.f4
        nw47_[int(10)] = self.f17
        nw47_[int(11)] = default__.fm1(globalState)
        nw47_[int(12)] = 647
        nw47_[int(13)] = (self).f15
        nw47_[int(14)] = (self).f15
        nw47_[int(15)] = self.f17
        nw47_[int(16)] = 806
        nw47_[int(17)] = self.f4
        nw47_[int(18)] = (self).f15
        nw47_[int(19)] = self.f4
        nw47_[int(20)] = self.f17
        nw47_[int(21)] = self.f4
        nw47_[int(22)] = (self.f17) * (self.f17)
        nw47_[int(23)] = len(_dafny.Map({(self).f16: self.f17}))
        nw47_[int(24)] = (self).f15
        nw47_[int(25)] = self.f17
        nw47_[int(26)] = (self).f15
        nw47_[int(27)] = (self.f4) + (len(d_352_v7_))
        d_353_v8_ = nw47_
        index30_ = default__.safeIndex(119, (d_353_v8_).length(0))
        (d_353_v8_)[index30_] = len(d_351_v6_)
        index31_ = default__.safeIndex(119, (d_353_v8_).length(0))
        (d_353_v8_)[index31_] = (D3_DC10(503, (self).f15, (self).f16, (self).f15, (d_342_v0_)[default__.safeIndex(self.f17, len(d_342_v0_))])).cf18
        d_354_v9_: _dafny.Array
        nw48_ = _dafny.Array(False, 7)
        d_354_v9_ = nw48_
        d_354_v9_ = d_354_v9_
        d_355_v10_: _dafny.MultiSet
        d_355_v10_ = _dafny.MultiSet([self.f17])
        d_356_i3_: int
        d_356_i3_ = 0
        with _dafny.label("3"):
            while not((d_355_v10_).isdisjoint((d_355_v10_) - (d_355_v10_))):
                with _dafny.c_label("3"):
                    if (d_356_i3_) >= (100):
                        raise _dafny.Break("3")
                    d_356_i3_ = (d_356_i3_) + (1)
                    index32_ = default__.safeIndex(119, (d_353_v8_).length(0))
                    (d_353_v8_)[index32_] = (d_353_v8_)[default__.safeIndex(119, (d_353_v8_).length(0))]
                    d_357_v11_: bool
                    d_357_v11_ = True
                    d_357_v11_ = (default__.safeDivisionInt((self).f15, (self).f15)) > (self.f4)
                    d_358_v12_: _dafny.Array
                    nw49_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 10)
                    d_358_v12_ = nw49_
                    rhs34_ = (0) - (default__.safeDivisionInt(self.f17, default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_359_i4_ in range(default__.abs(909))])), 224)))
                    rhs35_ = d_358_v12_
                    lhs15_ = self
                    lhs15_.f4 = rhs34_
                    d_358_v12_ = rhs35_
                    d_360_v13_: C0
                    nw50_ = C0()
                    nw50_.ctor__(default__.safeModuloInt((d_353_v8_)[default__.safeIndex(119, (d_353_v8_).length(0))], (0) - (self.f17)))
                    d_360_v13_ = nw50_
                    pass
            pass
        (self).f4 = (((self).f15) * ((0) - (self.f4))) - (self.f17)

    def m3(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        d_361_v0_: str
        d_361_v0_ = _dafny.CodePoint('i')
        d_362_v1_: _dafny.Array
        def lambda13_(d_363_p0_):
            def lambda14_(d_364_i0_):
                return d_363_p0_

            return lambda14_

        init7_ = lambda13_(p0)
        nw51_ = _dafny.Array(None, 26)
        for i0_7_ in range(nw51_.length(0)):
            nw51_[i0_7_] = init7_(i0_7_)
        d_362_v1_ = nw51_
        d_365_v2_: D3
        d_365_v2_ = D3_DC8(d_361_v0_, d_362_v1_)
        source7_ = d_365_v2_
        if source7_.is_DC8:
            d_366___mcc_h0_ = source7_.cf13
            d_367___mcc_h1_ = source7_.cf14
            d_368_cf14_ = d_367___mcc_h1_
            d_369_cf13_ = d_366___mcc_h0_
            d_370_v3_: _dafny.Seq
            d_370_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kfitckc"))
            d_371_v4_: D2
            d_371_v4_ = D2_DC6(len(d_370_v3_), p3)
            d_372_v5_: _dafny.Seq
            d_372_v5_ = _dafny.SeqWithoutIsStrInference([d_371_v4_, d_371_v4_])
            d_373_v6_: _dafny.Map
            d_373_v6_ = _dafny.Map({p3: (self).f15})
            pat_let_tv10_ = globalState
            pat_let_tv11_ = globalState
            def iife55_(_pat_let2_0):
                def iife56_(d_374_dt__update__tmp_h0_):
                    def iife57_(_pat_let3_0):
                        def iife58_(d_375_dt__update_hcf10_h0_):
                            return D2_DC6(d_375_dt__update_hcf10_h0_, (d_374_dt__update__tmp_h0_).cf11)
                        return iife58_(_pat_let3_0)
                    return iife57_(default__.fm1(pat_let_tv10_))
                return iife56_(_pat_let2_0)
            def iife59_(_pat_let4_0):
                def iife60_(d_376_dt__update__tmp_h1_):
                    def iife61_(_pat_let5_0):
                        def iife62_(d_377_dt__update_hcf10_h1_):
                            return D2_DC6(d_377_dt__update_hcf10_h1_, (d_376_dt__update__tmp_h1_).cf11)
                        return iife62_(_pat_let5_0)
                    return iife61_(default__.fm1(pat_let_tv11_))
                return iife60_(_pat_let4_0)
            (self).f4 = len(((d_372_v5_).set(default__.safeIndex(p1, len(d_372_v5_)), iife55_(d_371_v4_))).set(default__.safeIndex(((d_373_v6_)[p2] if (p2) in (d_373_v6_) else default__.fm1(globalState)), len((d_372_v5_).set(default__.safeIndex(p1, len(d_372_v5_)), iife59_(d_371_v4_)))), d_371_v4_))
            r0 = (self).f16
            (self).f4 = ((self.f4) * (self.f4)) - (self.f17)
            d_370_v3_ = d_370_v3_
        elif source7_.is_DC9:
            d_378___mcc_h2_ = source7_.cf15
            d_379___mcc_h3_ = source7_.cf16
            d_380_cf16_ = d_379___mcc_h3_
            d_381_cf15_ = d_378___mcc_h2_
            d_361_v0_ = d_361_v0_
            d_382_v7_: _dafny.Seq
            d_382_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nbfax"))
            d_383_v8_: _dafny.Map
            d_383_v8_ = _dafny.Map({d_382_v7_: (d_361_v0_) in (d_382_v7_)})
            d_384_v9_: _dafny.Array
            nw52_ = _dafny.Array(int(0), 26)
            d_384_v9_ = nw52_
            rhs36_ = d_381_cf15_
            rhs37_ = p1
            rhs38_ = d_383_v8_
            rhs39_ = d_384_v9_
            r0 = rhs36_
            d_380_cf16_ = rhs37_
            d_383_v8_ = rhs38_
            d_384_v9_ = rhs39_
            r0 = ((self).f16) and (p2)
            if p2:
                index33_ = default__.safeIndex(819, (d_384_v9_).length(0))
                (d_384_v9_)[index33_] = (self).f15
                d_385_v10_: _dafny.MultiSet
                d_385_v10_ = _dafny.MultiSet([d_381_cf15_, (self).f16, p3, (self).f16, not(not(True))])
                d_386_v11_: _dafny.Map
                d_386_v11_ = _dafny.Map({d_385_v10_: True})
                d_387_v12_: _dafny.Set
                d_387_v12_ = _dafny.Set({(self).f16})
                index34_ = default__.safeIndex(819, (d_384_v9_).length(0))
                (d_384_v9_)[index34_] = default__.safeDivisionInt(len(d_386_v11_), (len(d_387_v12_)) - (d_380_cf16_))
                d_388_v13_: D4
                d_388_v13_ = D4_DC12(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_389_i1_ in range(default__.abs(757))]))
                d_390_v14_: _dafny.Seq
                d_390_v14_ = _dafny.SeqWithoutIsStrInference([d_382_v7_, d_382_v7_])
                rhs40_ = d_362_v1_
                rhs41_ = (d_388_v13_).cf23
                rhs42_ = ((len(d_390_v14_)) * ((self).f15)) - (default__.safeDivisionInt(p1, self.f4))
                lhs16_ = self
                d_362_v1_ = rhs40_
                d_382_v7_ = rhs41_
                lhs16_.f17 = rhs42_
                rhs43_ = d_381_cf15_
                rhs44_ = not (d_381_cf15_) or ((self).f16)
                r0 = rhs43_
                r0 = rhs44_
                d_391_v15_: _dafny.Map
                d_391_v15_ = _dafny.Map({p3: p0})
                d_392_v16_: _dafny.Map
                d_392_v16_ = _dafny.Map({not(not(((d_391_v15_)[p3] if (p3) in (d_391_v15_) else p3))): d_361_v0_})
                rhs45_ = (d_392_v16_) | (d_392_v16_)
                rhs46_ = default__.safeModuloInt(d_380_cf16_, (p1 if d_381_cf15_ else p1))
                d_392_v16_ = rhs45_
                d_380_cf16_ = rhs46_
                d_393_v17_: _dafny.Seq
                d_393_v17_ = _dafny.SeqWithoutIsStrInference([p2])
                d_394_v18_: D1
                d_394_v18_ = D1_DC4(d_380_cf16_, d_393_v17_, self.f17)
                d_395_v19_: D4
                d_395_v19_ = D4_DC13(d_394_v18_, p2, d_393_v17_)
                d_396_v20_: _dafny.Map
                d_396_v20_ = _dafny.Map({self.f4: d_395_v19_})
                d_396_v20_ = (d_396_v20_).set(self.f4, D4_DC13(D1_DC4((d_394_v18_).cf6, d_393_v17_, len(d_382_v7_)), p3, _dafny.SeqWithoutIsStrInference([p3, d_381_cf15_, p0])))
            elif True:
                (self).f4 = self.f4
                d_397_v21_: _dafny.Seq
                d_397_v21_ = _dafny.SeqWithoutIsStrInference([(self).f15, (0) - ((self).f15), self.f17, self.f17])
                d_397_v21_ = _dafny.SeqWithoutIsStrInference([((self).f15 if p3 else 229) for d_398_i2_ in range(default__.abs(763))])
                d_399_v22_: C0
                nw53_ = C0()
                nw53_.ctor__(self.f17)
                d_399_v22_ = nw53_
                d_400_v23_: _dafny.Array
                def lambda15_(d_401_v21_):
                    def lambda16_(d_402_i3_):
                        return _dafny.MultiSet(d_401_v21_)

                    return lambda16_

                init8_ = lambda15_(d_397_v21_)
                nw54_ = _dafny.Array(None, 28)
                for i0_8_ in range(nw54_.length(0)):
                    nw54_[i0_8_] = init8_(i0_8_)
                d_400_v23_ = nw54_
                d_403_v24_: _dafny.Map
                d_403_v24_ = _dafny.Map({d_400_v23_: (self).f15})
                d_404_v25_: _dafny.Map
                d_404_v25_ = _dafny.Map({d_381_cf15_: (self).f16})
                d_403_v24_ = (d_403_v24_).set((D5_DC15(d_400_v23_)).cf28, (len((d_404_v25_).set(p3, p2))) + (len(d_397_v21_)))
                d_405_v26_: _dafny.Array
                nw55_ = _dafny.Array(None, 6)
                nw55_[int(0)] = d_362_v1_
                nw55_[int(1)] = d_362_v1_
                nw55_[int(2)] = d_362_v1_
                nw55_[int(3)] = d_362_v1_
                nw55_[int(4)] = d_362_v1_
                nw55_[int(5)] = d_362_v1_
                d_405_v26_ = nw55_
                d_406_v27_: _dafny.Map
                d_406_v27_ = _dafny.Map({497: d_405_v26_})
                d_406_v27_ = (d_406_v27_).set(len(d_404_v25_), d_405_v26_)
        elif source7_.is_DC10:
            d_407___mcc_h4_ = source7_.cf17
            d_408___mcc_h5_ = source7_.cf18
            d_409___mcc_h6_ = source7_.cf19
            d_410___mcc_h7_ = source7_.cf20
            d_411___mcc_h8_ = source7_.cf21
            d_412_cf21_ = d_411___mcc_h8_
            d_413_cf20_ = d_410___mcc_h7_
            d_414_cf19_ = d_409___mcc_h6_
            d_415_cf18_ = d_408___mcc_h5_
            d_416_cf17_ = d_407___mcc_h4_
            d_417_v28_: _dafny.Seq
            d_417_v28_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g"))
            if ((self).f16) == (not((d_417_v28_) < (d_417_v28_))):
                d_418_v29_: _dafny.Seq
                d_418_v29_ = _dafny.SeqWithoutIsStrInference([(0) - ((self).f15)])
                d_419_v30_: _dafny.MultiSet
                d_419_v30_ = _dafny.MultiSet([d_418_v29_])
                d_420_v31_: _dafny.Seq
                d_420_v31_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([958]) for d_421_i4_ in range(default__.abs(461))])), d_419_v30_])
                r0 = (p2 if p3 else (d_419_v30_).issubset((d_420_v31_)[default__.safeIndex(self.f17, len(d_420_v31_))]))
                d_422_v32_: _dafny.Set
                d_422_v32_ = _dafny.Set({default__.fm1(globalState), (self).f15})
                d_423_v33_: D3
                d_423_v33_ = D3_DC10(len(d_417_v28_), d_416_cf17_, d_412_cf21_, d_416_cf17_, not(d_412_cf21_))
                d_424_v34_: _dafny.Seq
                d_424_v34_ = _dafny.SeqWithoutIsStrInference([False])
                d_425_v35_: _dafny.MultiSet
                d_425_v35_ = _dafny.MultiSet([self.f17])
                pat_let_tv12_ = d_413_cf20_
                pat_let_tv13_ = d_424_v34_
                pat_let_tv14_ = p1
                pat_let_tv15_ = d_424_v34_
                pat_let_tv16_ = d_413_cf20_
                pat_let_tv17_ = d_414_cf19_
                pat_let_tv18_ = p0
                pat_let_tv19_ = d_424_v34_
                pat_let_tv20_ = p3
                pat_let_tv21_ = p2
                d_426_v36_: _dafny.Array
                nw56_ = _dafny.Array(None, 18)
                nw56_[int(0)] = default__.fm20((self).fm5(_dafny.SeqWithoutIsStrInference([d_361_v0_ for d_427_i5_ in range(default__.abs(870))]), p0, globalState), d_422_v32_, globalState)
                nw56_[int(1)] = d_423_v33_
                nw56_[int(2)] = d_423_v33_
                nw56_[int(3)] = d_423_v33_
                nw56_[int(4)] = d_423_v33_
                nw56_[int(5)] = d_423_v33_
                def iife64_(_pat_let7_0):
                    def iife65_(d_428_dt__update__tmp_h2_):
                        def iife66_(_pat_let8_0):
                            def iife67_(d_429_dt__update_hcf20_h0_):
                                def iife68_(_pat_let9_0):
                                    def iife69_(d_430_dt__update_hcf19_h0_):
                                        return D3_DC10((d_428_dt__update__tmp_h2_).cf17, (d_428_dt__update__tmp_h2_).cf18, d_430_dt__update_hcf19_h0_, d_429_dt__update_hcf20_h0_, (d_428_dt__update__tmp_h2_).cf21)
                                    return iife69_(_pat_let9_0)
                                return iife68_((pat_let_tv13_)[default__.safeIndex(pat_let_tv14_, len(pat_let_tv15_))])
                            return iife67_(_pat_let8_0)
                        return iife66_(pat_let_tv12_)
                    return iife65_(_pat_let7_0)
                def iife63_(_pat_let6_0):
                    def iife70_(d_431_dt__update__tmp_h3_):
                        def iife71_(_pat_let10_0):
                            def iife72_(d_432_dt__update_hcf18_h0_):
                                def iife73_(_pat_let11_0):
                                    def iife74_(d_433_dt__update_hcf20_h1_):
                                        return D3_DC10((d_431_dt__update__tmp_h3_).cf17, d_432_dt__update_hcf18_h0_, (d_431_dt__update__tmp_h3_).cf19, d_433_dt__update_hcf20_h1_, (d_431_dt__update__tmp_h3_).cf21)
                                    return iife74_(_pat_let11_0)
                                return iife73_(pat_let_tv16_)
                            return iife72_(_pat_let10_0)
                        return iife71_((self).f15)
                    return iife70_(_pat_let6_0)
                nw56_[int(6)] = iife63_(iife64_(d_423_v33_))
                def iife75_(_pat_let12_0):
                    def iife76_(d_434_dt__update__tmp_h4_):
                        def iife77_(_pat_let13_0):
                            def iife78_(d_435_dt__update_hcf19_h1_):
                                def iife79_(_pat_let14_0):
                                    def iife80_(d_436_dt__update_hcf21_h0_):
                                        return D3_DC10((d_434_dt__update__tmp_h4_).cf17, (d_434_dt__update__tmp_h4_).cf18, d_435_dt__update_hcf19_h1_, (d_434_dt__update__tmp_h4_).cf20, d_436_dt__update_hcf21_h0_)
                                    return iife80_(_pat_let14_0)
                                return iife79_(pat_let_tv18_)
                            return iife78_(_pat_let13_0)
                        return iife77_(pat_let_tv17_)
                    return iife76_(_pat_let12_0)
                nw56_[int(7)] = iife75_(d_423_v33_)
                nw56_[int(8)] = d_423_v33_
                def iife81_(_pat_let15_0):
                    def iife82_(d_437_dt__update__tmp_h5_):
                        def iife83_(_pat_let16_0):
                            def iife84_(d_438_dt__update_hcf18_h1_):
                                def iife85_(_pat_let17_0):
                                    def iife86_(d_439_dt__update_hcf17_h0_):
                                        def iife87_(_pat_let18_0):
                                            def iife88_(d_440_dt__update_hcf21_h1_):
                                                def iife89_(_pat_let19_0):
                                                    def iife90_(d_441_dt__update_hcf19_h2_):
                                                        return D3_DC10(d_439_dt__update_hcf17_h0_, d_438_dt__update_hcf18_h1_, d_441_dt__update_hcf19_h2_, (d_437_dt__update__tmp_h5_).cf20, d_440_dt__update_hcf21_h1_)
                                                    return iife90_(_pat_let19_0)
                                                return iife89_(pat_let_tv21_)
                                            return iife88_(_pat_let18_0)
                                        return iife87_(pat_let_tv20_)
                                    return iife86_(_pat_let17_0)
                                return iife85_(len(pat_let_tv19_))
                            return iife84_(_pat_let16_0)
                        return iife83_(519)
                    return iife82_(_pat_let15_0)
                nw56_[int(9)] = iife81_(D3_DC10((d_425_v35_).cardinality, (d_425_v35_).cardinality, p2, d_413_cf20_, d_412_cf21_))
                nw56_[int(10)] = d_423_v33_
                nw56_[int(11)] = d_423_v33_
                nw56_[int(12)] = D3_DC10(-833, p1, (self).f16, self.f4, default__.fm0(d_416_cf17_, len(d_417_v28_), globalState))
                nw56_[int(13)] = d_423_v33_
                nw56_[int(14)] = d_423_v33_
                nw56_[int(15)] = d_423_v33_
                nw56_[int(16)] = d_423_v33_
                nw56_[int(17)] = d_423_v33_
                d_426_v36_ = nw56_
                d_442_v37_: _dafny.Map
                d_442_v37_ = _dafny.Map({False: d_412_cf21_})
                d_443_v38_: _dafny.Map
                d_443_v38_ = _dafny.Map({p2: len(_dafny.Map({d_442_v37_: (d_424_v34_)[default__.safeIndex(self.f4, len(d_424_v34_))]}))})
                index35_ = default__.safeIndex(346, (d_426_v36_).length(0))
                (d_426_v36_)[index35_] = D3_DC10(self.f17, (0) - (d_416_cf17_), p2, len(d_443_v38_), p2)
                index36_ = default__.safeIndex(346, (d_426_v36_).length(0))
                (d_426_v36_)[index36_] = d_423_v33_
                d_416_cf17_ = default__.safeDivisionInt(d_413_cf20_, d_416_cf17_)
                d_424_v34_ = (_dafny.SeqWithoutIsStrInference([(self).f16, p2, p3])) + ((d_424_v34_) + (d_424_v34_))
                d_444_v39_: _dafny.Map
                d_444_v39_ = _dafny.Map({d_416_cf17_: d_362_v1_})
                d_444_v39_ = d_444_v39_
            elif True:
                d_445_v40_: _dafny.Array
                nw57_ = _dafny.Array(_dafny.Array(None, 0), 8)
                d_445_v40_ = nw57_
                d_446_v42_: _dafny.Map
                d_446_v42_ = _dafny.Map({d_413_cf20_: (self).f16})
                d_447_v43_: _dafny.Map
                d_447_v43_ = _dafny.Map({p1: d_361_v0_})
                d_448_v45_: _dafny.Array
                nw58_ = _dafny.Array(None, 6)
                def iife91_():
                    coll51_ = _dafny.Map()
                    compr_51_: int
                    for compr_51_ in (d_446_v42_).keys.Elements:
                        d_449_v41_: int = compr_51_
                        if (d_449_v41_) in (d_446_v42_):
                            coll51_[(d_449_v41_) - (p1)] = _dafny.CodePoint('d')
                    return _dafny.Map(coll51_)
                nw58_[int(0)] = iife91_()
                
                nw58_[int(1)] = (d_447_v43_).set(p1, d_361_v0_)
                def iife92_():
                    coll52_ = _dafny.Map()
                    compr_52_: int
                    for compr_52_ in _dafny.IntegerRange(513, 508):
                        d_450_v44_: int = compr_52_
                        if ((513) <= (d_450_v44_)) and ((d_450_v44_) < (508)):
                            coll52_[default__.safeModuloInt(d_450_v44_, p1)] = _dafny.CodePoint('w')
                    return _dafny.Map(coll52_)
                nw58_[int(2)] = iife92_()
                
                nw58_[int(3)] = d_447_v43_
                nw58_[int(4)] = d_447_v43_
                nw58_[int(5)] = d_447_v43_
                d_448_v45_ = nw58_
                index37_ = default__.safeIndex(987, (d_445_v40_).length(0))
                (d_445_v40_)[index37_] = d_448_v45_
                index38_ = default__.safeIndex(987, (d_445_v40_).length(0))
                (d_445_v40_)[index38_] = d_448_v45_
                d_451_v46_: _dafny.Set
                d_451_v46_ = _dafny.Set({d_416_cf17_, self.f4, p1})
                d_451_v46_ = d_451_v46_
                d_361_v0_ = d_361_v0_
                d_452_v47_: _dafny.Seq
                d_452_v47_ = _dafny.SeqWithoutIsStrInference([not(d_414_cf19_), d_412_cf21_, d_412_cf21_, not(p0)])
                index39_ = default__.safeIndex(693, (d_362_v1_).length(0))
                (d_362_v1_)[index39_] = (d_452_v47_)[default__.safeIndex((self).f15, len(d_452_v47_))]
                index40_ = default__.safeIndex(693, (d_362_v1_).length(0))
                (d_362_v1_)[index40_] = p3
                d_453_v48_: C0
                nw59_ = C0()
                nw59_.ctor__(self.f17)
                d_453_v48_ = nw59_
            d_454_v49_: _dafny.Map
            d_454_v49_ = _dafny.Map({(self).f16: (self).f15})
            d_455_v50_: _dafny.Seq
            d_455_v50_ = _dafny.SeqWithoutIsStrInference([(self).f16, d_412_cf21_])
            d_454_v49_ = (d_454_v49_).set((d_455_v50_)[default__.safeIndex(d_413_cf20_, len(d_455_v50_))], d_415_cf18_)
            d_456_v51_: _dafny.Map
            d_456_v51_ = _dafny.Map({self.f17: d_362_v1_})
            d_413_cf20_ = default__.safeModuloInt((454) - ((self).f15), default__.safeDivisionInt(len(d_456_v51_), self.f4))
            d_457_v52_: _dafny.Seq
            d_457_v52_ = _dafny.SeqWithoutIsStrInference([618, (self).f15])
            d_458_v53_: _dafny.Set
            d_458_v53_ = _dafny.Set({(self).f16})
            d_459_v54_: _dafny.Map
            d_459_v54_ = _dafny.Map({p0: (self).f16})
            d_460_v55_: D4
            d_460_v55_ = D4_DC13(D1_DC4(d_413_cf20_, d_455_v50_, len(d_458_v53_)), ((d_459_v54_)[p2] if (p2) in (d_459_v54_) else False), d_455_v50_)
            d_461_v56_: _dafny.MultiSet
            d_461_v56_ = _dafny.MultiSet([False])
            def iife93_():
                coll53_ = _dafny.Set()
                compr_53_: _dafny.MultiSet
                for compr_53_ in (_dafny.Set({d_461_v56_})).Elements:
                    d_462_v57_: _dafny.MultiSet = compr_53_
                    if (d_462_v57_) in (_dafny.Set({d_461_v56_})):
                        coll53_ = coll53_.union(_dafny.Set([d_462_v57_]))
                return _dafny.Set(coll53_)
            d_457_v52_ = (d_457_v52_) + ((d_457_v52_) + (default__.fm21(len((d_455_v50_).set(default__.safeIndex(self.f17, len(d_455_v50_)), (self).fm5(d_417_v28_, (d_460_v55_).cf25, globalState))), d_414_cf19_, p0, iife93_()
            , globalState)))
        elif source7_.is_DC7:
            d_463___mcc_h9_ = source7_.cf12
            d_464_cf12_ = d_463___mcc_h9_
            d_465_v58_: _dafny.Seq
            d_465_v58_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nkuxq"))
            d_465_v58_ = ((_dafny.SeqWithoutIsStrInference([d_361_v0_ for d_466_i6_ in range(default__.abs(781))])).set(default__.safeIndex(271, len(_dafny.SeqWithoutIsStrInference([d_361_v0_ for d_467_i6_ in range(default__.abs(781))]))), d_361_v0_)) + (d_465_v58_)
            r0 = (p2 if (self).f16 else (self).f16)
            if p3:
                d_468_v59_: _dafny.Array
                nw60_ = _dafny.Array(int(0), 19)
                d_468_v59_ = nw60_
                d_469_v60_: _dafny.Map
                d_469_v60_ = _dafny.Map({not(p2): d_468_v59_})
                (self).f17 = len(((d_469_v60_).set(p3, d_468_v59_)).set(not(p3), d_468_v59_))
                d_470_v61_: _dafny.Seq
                d_470_v61_ = _dafny.SeqWithoutIsStrInference([d_464_cf12_])
                d_471_v62_: _dafny.Map
                d_471_v62_ = _dafny.Map({default__.safeDivisionInt(self.f17, p1): (_dafny.Set({p1})) == ((d_470_v61_)[default__.safeIndex(self.f4, len(d_470_v61_))])})
                def iife94_():
                    coll54_ = _dafny.Map()
                    compr_54_: int
                    for compr_54_ in _dafny.IntegerRange(-445, -923):
                        d_472_v63_: int = compr_54_
                        if ((-445) <= (d_472_v63_)) and ((d_472_v63_) < (-923)):
                            coll54_[default__.safeDivisionInt(d_472_v63_, p1)] = (p3) and ((self).f16)
                    return _dafny.Map(coll54_)
                d_471_v62_ = iife94_()
                
                d_473_v64_: _dafny.Set
                d_473_v64_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([self.f4])})
                d_474_v65_: C0
                nw61_ = C0()
                nw61_.ctor__((p1) - (len(d_473_v64_)))
                d_474_v65_ = nw61_
                (self).f17 = default__.fm1(globalState)
                (self).f4 = (0) - ((0) - ((self).f15))
            elif True:
                d_475_v66_: _dafny.MultiSet
                d_475_v66_ = _dafny.MultiSet([(self).f15, 589])
                (self).f4 = (((d_475_v66_)[p1] if (p1) in (d_475_v66_) else self.f4)) - (default__.safeModuloInt(self.f4, 679))
                d_465_v58_ = d_465_v58_
                d_476_v67_: _dafny.Array
                nw62_ = _dafny.Array(_dafny.Set({}), 18)
                d_476_v67_ = nw62_
                index41_ = default__.safeIndex(0, (d_476_v67_).length(0))
                (d_476_v67_)[index41_] = d_464_cf12_
                d_477_v68_: _dafny.Array
                nw63_ = _dafny.Array(D5.default()(), 7)
                d_477_v68_ = nw63_
                d_478_v69_: D5
                d_478_v69_ = D5_DC16((self).f16, not(p0), p2, p1, not(True))
                index42_ = default__.safeIndex(113, (d_477_v68_).length(0))
                (d_477_v68_)[index42_] = d_478_v69_
                index43_ = default__.safeIndex(0, (d_476_v67_).length(0))
                index44_ = default__.safeIndex(113, (d_477_v68_).length(0))
                rhs47_ = True
                rhs48_ = p1
                rhs49_ = d_464_cf12_
                rhs50_ = d_362_v1_
                rhs51_ = d_478_v69_
                lhs17_ = self
                lhs18_ = d_476_v67_
                lhs19_ = default__.safeIndex(0, (d_476_v67_).length(0))
                lhs20_ = d_477_v68_
                lhs21_ = default__.safeIndex(113, (d_477_v68_).length(0))
                r0 = rhs47_
                lhs17_.f17 = rhs48_
                lhs18_[lhs19_] = rhs49_
                d_362_v1_ = rhs50_
                lhs20_[lhs21_] = rhs51_
                d_479_v70_: _dafny.Array
                nw64_ = _dafny.Array(None, 19)
                nw64_[int(0)] = d_362_v1_
                nw64_[int(1)] = d_362_v1_
                nw64_[int(2)] = d_362_v1_
                nw64_[int(3)] = d_362_v1_
                nw64_[int(4)] = d_362_v1_
                nw64_[int(5)] = d_362_v1_
                nw64_[int(6)] = d_362_v1_
                nw64_[int(7)] = d_362_v1_
                nw64_[int(8)] = d_362_v1_
                nw64_[int(9)] = d_362_v1_
                nw64_[int(10)] = d_362_v1_
                nw64_[int(11)] = d_362_v1_
                nw64_[int(12)] = d_362_v1_
                nw64_[int(13)] = d_362_v1_
                nw64_[int(14)] = d_362_v1_
                nw64_[int(15)] = d_362_v1_
                nw64_[int(16)] = d_362_v1_
                nw64_[int(17)] = d_362_v1_
                nw64_[int(18)] = d_362_v1_
                d_479_v70_ = nw64_
                d_480_v71_: _dafny.Map
                d_480_v71_ = _dafny.Map({d_479_v70_: d_362_v1_})
                d_481_v72_: _dafny.Seq
                d_481_v72_ = _dafny.SeqWithoutIsStrInference([d_362_v1_])
                d_482_v73_: _dafny.Array
                nw65_ = _dafny.Array(None, 26)
                nw65_[int(0)] = d_362_v1_
                nw65_[int(1)] = d_362_v1_
                nw65_[int(2)] = d_362_v1_
                nw65_[int(3)] = d_362_v1_
                nw65_[int(4)] = d_362_v1_
                nw65_[int(5)] = (d_362_v1_ if (self).f16 else d_362_v1_)
                nw65_[int(6)] = d_362_v1_
                nw65_[int(7)] = d_362_v1_
                nw65_[int(8)] = (d_362_v1_ if p0 else d_362_v1_)
                nw65_[int(9)] = ((d_480_v71_)[d_479_v70_] if (d_479_v70_) in (d_480_v71_) else d_362_v1_)
                nw65_[int(10)] = d_362_v1_
                nw65_[int(11)] = d_362_v1_
                nw65_[int(12)] = d_362_v1_
                nw65_[int(13)] = d_362_v1_
                nw65_[int(14)] = d_362_v1_
                nw65_[int(15)] = d_362_v1_
                nw65_[int(16)] = (d_481_v72_)[default__.safeIndex(self.f17, len(d_481_v72_))]
                nw65_[int(17)] = d_362_v1_
                nw65_[int(18)] = (d_362_v1_ if p0 else d_362_v1_)
                nw65_[int(19)] = d_362_v1_
                nw65_[int(20)] = (D3_DC8(d_361_v0_, d_362_v1_)).cf14
                nw65_[int(21)] = d_362_v1_
                nw65_[int(22)] = d_362_v1_
                nw65_[int(23)] = d_362_v1_
                nw65_[int(24)] = d_362_v1_
                nw65_[int(25)] = d_362_v1_
                d_482_v73_ = nw65_
                index45_ = default__.safeIndex(166, (d_482_v73_).length(0))
                (d_482_v73_)[index45_] = d_362_v1_
                index46_ = default__.safeIndex(166, (d_482_v73_).length(0))
                (d_482_v73_)[index46_] = d_362_v1_
                (self).f17 = self.f17
            d_483_v74_: C0
            nw66_ = C0()
            nw66_.ctor__((self.f4) * (428))
            d_483_v74_ = nw66_
        elif True:
            d_484___mcc_h10_ = source7_.cf22
            d_485_cf22_ = d_484___mcc_h10_
            d_486_v75_: _dafny.Map
            d_486_v75_ = _dafny.Map({p0: (self).f16})
            d_486_v75_ = (d_486_v75_).set(p2, not(p0))
            r0 = p0
            d_487_v76_: _dafny.Map
            d_487_v76_ = _dafny.Map({p1: p0})
            d_488_v77_: _dafny.MultiSet
            d_488_v77_ = _dafny.MultiSet([p1])
            d_489_v78_: _dafny.MultiSet
            d_489_v78_ = _dafny.MultiSet([len(d_487_v76_), (d_488_v77_).cardinality])
            d_490_v79_: _dafny.Seq
            d_490_v79_ = _dafny.SeqWithoutIsStrInference([p2])
            d_491_v80_: _dafny.Seq
            d_491_v80_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rmirrqiph"))
            d_492_v81_: _dafny.Seq
            d_492_v81_ = _dafny.SeqWithoutIsStrInference([(self).f15])
            r0 = (((d_489_v78_).set(len(d_490_v79_), default__.abs(len(d_491_v80_)))) | (_dafny.MultiSet([p1, self.f17]))).isdisjoint((d_488_v77_).set((self).f15, default__.abs((d_492_v81_)[default__.safeIndex(len(d_491_v80_), len(d_492_v81_))])))
            r0 = (self).f16
        d_493_v82_: _dafny.Map
        d_493_v82_ = _dafny.Map({p0: not(False)})
        d_494_v83_: _dafny.Map
        d_494_v83_ = _dafny.Map({(self).f15: not(True)})
        d_495_v84_: _dafny.Array
        nw67_ = _dafny.Array(None, 23)
        nw67_[int(0)] = (self).f15
        nw67_[int(1)] = self.f17
        nw67_[int(2)] = -770
        nw67_[int(3)] = 777
        nw67_[int(4)] = self.f4
        nw67_[int(5)] = (self).f15
        nw67_[int(6)] = p1
        nw67_[int(7)] = p1
        nw67_[int(8)] = self.f4
        nw67_[int(9)] = self.f4
        nw67_[int(10)] = len(d_493_v82_)
        nw67_[int(11)] = p1
        nw67_[int(12)] = self.f17
        nw67_[int(13)] = self.f4
        nw67_[int(14)] = len(d_494_v83_)
        nw67_[int(15)] = p1
        nw67_[int(16)] = self.f4
        nw67_[int(17)] = p1
        nw67_[int(18)] = self.f17
        nw67_[int(19)] = self.f4
        nw67_[int(20)] = self.f4
        nw67_[int(21)] = 285
        nw67_[int(22)] = self.f17
        d_495_v84_ = nw67_
        d_496_v85_: _dafny.Array
        nw68_ = _dafny.Array(None, 26)
        nw68_[int(0)] = d_495_v84_
        nw68_[int(1)] = (d_495_v84_ if default__.fm0(self.f17, 445, globalState) else d_495_v84_)
        nw68_[int(2)] = d_495_v84_
        nw68_[int(3)] = d_495_v84_
        nw68_[int(4)] = d_495_v84_
        nw68_[int(5)] = d_495_v84_
        nw68_[int(6)] = d_495_v84_
        nw68_[int(7)] = d_495_v84_
        nw68_[int(8)] = d_495_v84_
        nw68_[int(9)] = d_495_v84_
        nw68_[int(10)] = d_495_v84_
        nw68_[int(11)] = d_495_v84_
        nw68_[int(12)] = d_495_v84_
        nw68_[int(13)] = d_495_v84_
        nw68_[int(14)] = d_495_v84_
        nw68_[int(15)] = d_495_v84_
        nw68_[int(16)] = d_495_v84_
        nw68_[int(17)] = d_495_v84_
        nw68_[int(18)] = d_495_v84_
        nw68_[int(19)] = d_495_v84_
        nw68_[int(20)] = d_495_v84_
        nw68_[int(21)] = d_495_v84_
        nw68_[int(22)] = d_495_v84_
        nw68_[int(23)] = d_495_v84_
        nw68_[int(24)] = d_495_v84_
        nw68_[int(25)] = d_495_v84_
        d_496_v85_ = nw68_
        index47_ = default__.safeIndex(640, (d_496_v85_).length(0))
        (d_496_v85_)[index47_] = d_495_v84_
        index48_ = default__.safeIndex(640, (d_496_v85_).length(0))
        def lambda17_(d_497_i7_):
            return (d_497_i7_) - ((self).f15)

        init9_ = lambda17_
        nw69_ = _dafny.Array(None, 14)
        for i0_9_ in range(nw69_.length(0)):
            nw69_[i0_9_] = init9_(i0_9_)
        (d_496_v85_)[index48_] = nw69_
        d_498_v86_: _dafny.Seq
        d_498_v86_ = _dafny.SeqWithoutIsStrInference([p0])
        r0 = (not (p2) or ((self).f16)) and ((self.f17) <= (len(_dafny.Map({d_498_v86_: _dafny.CodePoint('j')}))))
        d_499_v87_: _dafny.Seq
        d_499_v87_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kaewedw"))
        rhs52_ = ((d_499_v87_) + (d_499_v87_)) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "salckobh")))
        rhs53_ = default__.fm1(globalState)
        lhs22_ = self
        r0 = rhs52_
        lhs22_.f4 = rhs53_
        (self).f4 = -878
        hi3_ = p1
        for d_500_i8_ in range(self.f17, hi3_):
            d_501_v88_: _dafny.Seq
            d_501_v88_ = _dafny.SeqWithoutIsStrInference([self.f4])
            index49_ = default__.safeIndex(563, (d_495_v84_).length(0))
            (d_495_v84_)[index49_] = ((_dafny.MultiSet(d_501_v88_)).cardinality) * ((self).f15)
            d_502_v89_: _dafny.Set
            d_502_v89_ = _dafny.Set({d_361_v0_, _dafny.CodePoint('b'), d_361_v0_})
            index50_ = default__.safeIndex(563, (d_495_v84_).length(0))
            rhs54_ = (default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g"))), p1) if (p3 if (self).f16 else p3) else d_500_i8_)
            rhs55_ = (((0) - (self.f17)) - (p1)) < ((len(d_502_v89_) if (self).f16 else self.f17))
            rhs56_ = (d_493_v82_).set(p0, (self).f16)
            lhs23_ = d_495_v84_
            lhs24_ = default__.safeIndex(563, (d_495_v84_).length(0))
            lhs23_[lhs24_] = rhs54_
            r0 = rhs55_
            d_493_v82_ = rhs56_
            d_503_v90_: _dafny.Map
            d_503_v90_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "htlnukccc")): (d_495_v84_)[default__.safeIndex(563, (d_495_v84_).length(0))]})
            d_504_v91_: _dafny.Map
            d_504_v91_ = _dafny.Map({d_501_v88_: ((0) - (((_dafny.MultiSet([self.f4])).set(len(d_503_v90_), default__.abs(121))).cardinality)) * (-417)})
            d_505_v92_: _dafny.Map
            d_505_v92_ = _dafny.Map({True: d_361_v0_})
            (self).f17 = ((d_504_v91_)[d_501_v88_] if (d_501_v88_) in (d_504_v91_) else len((d_505_v92_) | (d_505_v92_)))
            r0 = p0
            (self).f4 = d_500_i8_
        r0 = p3
        return r0

    def m4(self, p0, p1, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r2: _dafny.Array = _dafny.Array(None, 0)
        d_506_v0_: _dafny.Map
        d_506_v0_ = _dafny.Map({self.f17: not(p1)})
        d_507_v1_: D6
        d_507_v1_ = D6_DC20(d_506_v0_)
        d_508_v2_: _dafny.Seq
        d_508_v2_ = _dafny.SeqWithoutIsStrInference([(self).f15])
        d_509_v3_: _dafny.Set
        d_509_v3_ = _dafny.Set({p1, (self).f16})
        d_510_v4_: _dafny.Map
        d_510_v4_ = _dafny.Map({len(d_509_v3_): p0})
        d_511_v5_: _dafny.Array
        nw70_ = _dafny.Array(int(0), 7)
        d_511_v5_ = nw70_
        d_512_v6_: _dafny.Map
        d_512_v6_ = _dafny.Map({len(d_510_v4_): d_511_v5_})
        d_513_v7_: int
        d_514_v8_: int
        d_515_v9_: _dafny.Array
        d_516_v10_: bool
        out12_: int
        out13_: int
        out14_: _dafny.Array
        out15_: bool
        out12_, out13_, out14_, out15_ = default__.m0((d_507_v1_).cf40, d_508_v2_, d_512_v6_, globalState)
        d_513_v7_ = out12_
        d_514_v8_ = out13_
        d_515_v9_ = out14_
        d_516_v10_ = out15_
        index51_ = default__.safeIndex(357, (d_511_v5_).length(0))
        (d_511_v5_)[index51_] = (self).f15
        index52_ = default__.safeIndex(357, (d_511_v5_).length(0))
        (d_511_v5_)[index52_] = (self).f15
        d_517_v11_: _dafny.Array
        nw71_ = _dafny.Array(_dafny.CodePoint('D'), 9)
        d_517_v11_ = nw71_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_517_v11_).length(0)):
            d_518_i0_: int = guard_loop_0_
            if (True) and (((0) <= (d_518_i0_)) and ((d_518_i0_) < ((d_517_v11_).length(0)))):
                (d_517_v11_)[(d_518_i0_)] = (_dafny.CodePoint('n') if p1 else _dafny.CodePoint('e'))
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_511_v5_).length(0)):
            d_519_i1_: int = guard_loop_1_
            if (True) and (((0) <= (d_519_i1_)) and ((d_519_i1_) < ((d_511_v5_).length(0)))):
                (d_511_v5_)[(d_519_i1_)] = (d_519_i1_) + ((self).f15)
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_511_v5_).length(0)):
            d_520_i2_: int = guard_loop_2_
            if (True) and (((0) <= (d_520_i2_)) and ((d_520_i2_) < ((d_511_v5_).length(0)))):
                (d_511_v5_)[(d_520_i2_)] = default__.safeModuloInt(d_520_i2_, (self).f15)
        d_521_i3_: int
        d_521_i3_ = 0
        with _dafny.label("4"):
            while not(p1):
                with _dafny.c_label("4"):
                    if (d_521_i3_) >= (100):
                        raise _dafny.Break("4")
                    d_521_i3_ = (d_521_i3_) + (1)
                    d_510_v4_ = (d_510_v4_).set((0) - (self.f4), self.f17)
                    d_522_v12_: _dafny.Array
                    nw72_ = _dafny.Array(_dafny.Map({}), 4)
                    d_522_v12_ = nw72_
                    d_523_v13_: _dafny.MultiSet
                    d_523_v13_ = _dafny.MultiSet([p1])
                    d_522_v12_ = (d_522_v12_ if (d_523_v13_).ispropersubset(d_523_v13_) else d_522_v12_)
                    d_524_v14_: _dafny.MultiSet
                    d_524_v14_ = _dafny.MultiSet([(self).f15, self.f4])
                    d_514_v8_ = (0) - (((d_524_v14_).cardinality) - (len((_dafny.Set({(self).f16})) - (d_509_v3_))))
                    d_514_v8_ = (default__.safeModuloInt(self.f17, p0)) - (self.f17)
                    pass
            pass
        d_525_v15_: _dafny.Seq
        d_525_v15_ = _dafny.SeqWithoutIsStrInference([d_508_v2_])
        d_526_v16_: D0
        d_526_v16_ = D0_DC0((d_525_v15_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hx"))), len(d_525_v15_))])
        d_527_v17_: _dafny.Seq
        d_527_v17_ = _dafny.SeqWithoutIsStrInference([D0_DC0(d_508_v2_), d_526_v16_])
        r0 = ((d_527_v17_) + (d_527_v17_)) + (d_527_v17_)
        d_528_v18_: _dafny.Seq
        d_528_v18_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aaab"))
        d_529_v19_: _dafny.Seq
        d_529_v19_ = _dafny.SeqWithoutIsStrInference([(d_528_v18_) + (d_528_v18_)])
        r1 = (d_529_v19_)[default__.safeIndex((0) - (default__.fm1(globalState)), len(d_529_v19_))]
        d_530_v20_: _dafny.Map
        d_530_v20_ = _dafny.Map({d_516_v10_: d_515_v9_})
        d_531_v21_: str
        d_531_v21_ = _dafny.CodePoint('p')
        r2 = ((d_530_v20_)[(d_516_v10_) in (default__.fm2(d_531_v21_, (self).f15, d_513_v7_, globalState))] if ((d_516_v10_) in (default__.fm2(d_531_v21_, (self).f15, d_513_v7_, globalState))) in (d_530_v20_) else d_511_v5_)
        return r0, r1, r2

    def m11(self, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: _dafny.Set = _dafny.Set({})
        d_532_v0_: bool
        d_532_v0_ = True
        d_532_v0_ = not(d_532_v0_)
        if not(default__.fm0(self.f4, self.f4, globalState)):
            d_533_v1_: _dafny.Map
            d_533_v1_ = _dafny.Map({d_532_v0_: not((self).f16)})
            d_534_v2_: C0
            nw73_ = C0()
            nw73_.ctor__(len(d_533_v1_))
            d_534_v2_ = nw73_
            d_535_v3_: _dafny.MultiSet
            d_535_v3_ = _dafny.MultiSet([self.f4])
            d_535_v3_ = d_535_v3_
            d_536_v4_: str
            d_536_v4_ = _dafny.CodePoint('s')
            d_537_v5_: D2
            d_537_v5_ = D2_DC6(self.f17, (self).f16)
            d_536_v4_ = (self).fm10(self.f4, (self).f16, (self).f16, d_537_v5_, globalState)
            d_538_v6_: _dafny.Seq
            d_538_v6_ = _dafny.SeqWithoutIsStrInference([(self).f16, not(d_532_v0_)])
            d_539_v7_: _dafny.Seq
            d_539_v7_ = _dafny.SeqWithoutIsStrInference([d_538_v6_, d_538_v6_])
            d_540_v8_: _dafny.Seq
            d_540_v8_ = _dafny.SeqWithoutIsStrInference([self.f17, (self).f15, d_534_v2_.f18, (self).f15, self.f17])
            d_541_v9_: _dafny.Seq
            d_541_v9_ = _dafny.SeqWithoutIsStrInference([d_540_v8_])
            d_542_v11_: _dafny.Set
            d_542_v11_ = _dafny.Set({d_540_v8_})
            def iife95_():
                coll55_ = _dafny.Set()
                compr_55_: _dafny.Seq
                for compr_55_ in (d_541_v9_).Elements:
                    d_543_v10_: _dafny.Seq = compr_55_
                    if (d_543_v10_) in (d_541_v9_):
                        coll55_ = coll55_.union(_dafny.Set([d_543_v10_]))
                return _dafny.Set(coll55_)
            (d_534_v2_).f18 = (len(d_539_v7_)) - (len((iife95_()
            ) - (d_542_v11_)))
            d_544_v12_: _dafny.Array
            def lambda18_(d_545_v0_):
                def lambda19_(d_546_i0_):
                    return d_545_v0_

                return lambda19_

            init10_ = lambda18_(d_532_v0_)
            nw74_ = _dafny.Array(None, 7)
            for i0_10_ in range(nw74_.length(0)):
                nw74_[i0_10_] = init10_(i0_10_)
            d_544_v12_ = nw74_
            d_547_v13_: _dafny.MultiSet
            d_547_v13_ = _dafny.MultiSet([d_544_v12_])
            if (d_547_v13_).isdisjoint((d_547_v13_).set(d_544_v12_, default__.abs((0) - (self.f4)))):
                d_548_v14_: _dafny.Array
                nw75_ = _dafny.Array(D6.default()(), 14)
                d_548_v14_ = nw75_
                d_548_v14_ = d_548_v14_
                d_549_v15_: C0
                nw76_ = C0()
                nw76_.ctor__(d_534_v2_.f18)
                d_549_v15_ = nw76_
                d_550_v16_: D3
                d_550_v16_ = D3_DC8(d_536_v4_, d_544_v12_)
                d_551_v17_: _dafny.Array
                nw77_ = _dafny.Array(None, 16)
                nw77_[int(0)] = d_544_v12_
                nw77_[int(1)] = d_544_v12_
                nw77_[int(2)] = d_544_v12_
                nw77_[int(3)] = d_544_v12_
                nw77_[int(4)] = d_544_v12_
                nw77_[int(5)] = (d_544_v12_ if (self).f16 else d_544_v12_)
                nw77_[int(6)] = d_544_v12_
                nw77_[int(7)] = d_544_v12_
                nw77_[int(8)] = d_544_v12_
                nw77_[int(9)] = d_544_v12_
                nw77_[int(10)] = d_544_v12_
                nw77_[int(11)] = d_544_v12_
                nw77_[int(12)] = d_544_v12_
                nw77_[int(13)] = (d_550_v16_).cf14
                nw77_[int(14)] = d_544_v12_
                nw77_[int(15)] = d_544_v12_
                d_551_v17_ = nw77_
                d_551_v17_ = d_551_v17_
                d_533_v1_ = (d_533_v1_).set(True, ((self).f16 if True else (self).f16))
                d_540_v8_ = ((d_540_v8_) + (d_540_v8_)) + (d_540_v8_)
            elif True:
                d_552_v18_: _dafny.Map
                d_552_v18_ = _dafny.Map({len(d_538_v6_): d_532_v0_})
                d_552_v18_ = (d_552_v18_).set((0) - (self.f4), (self).f16)
                d_553_v19_: _dafny.Map
                d_553_v19_ = _dafny.Map({_dafny.CodePoint('s'): d_544_v12_})
                d_554_v20_: _dafny.Array
                nw78_ = _dafny.Array(None, 17)
                nw78_[int(0)] = 875
                nw78_[int(1)] = d_534_v2_.f18
                nw78_[int(2)] = (0) - (d_534_v2_.f18)
                nw78_[int(3)] = len((d_542_v11_) - (d_542_v11_))
                nw78_[int(4)] = self.f4
                nw78_[int(5)] = d_534_v2_.f18
                nw78_[int(6)] = d_534_v2_.f18
                nw78_[int(7)] = d_534_v2_.f18
                nw78_[int(8)] = default__.fm1(globalState)
                nw78_[int(9)] = d_534_v2_.f18
                nw78_[int(10)] = d_534_v2_.f18
                nw78_[int(11)] = len(d_553_v19_)
                nw78_[int(12)] = 531
                nw78_[int(13)] = self.f4
                nw78_[int(14)] = default__.safeDivisionInt(self.f4, self.f17)
                nw78_[int(15)] = d_534_v2_.f18
                nw78_[int(16)] = self.f17
                d_554_v20_ = nw78_
                index53_ = default__.safeIndex(450, (d_554_v20_).length(0))
                (d_554_v20_)[index53_] = (0) - (d_534_v2_.f18)
                index54_ = default__.safeIndex(450, (d_554_v20_).length(0))
                (d_554_v20_)[index54_] = self.f4
                (d_534_v2_).f18 = (d_554_v20_)[default__.safeIndex(450, (d_554_v20_).length(0))]
                d_555_v21_: _dafny.Map
                d_555_v21_ = _dafny.Map({(d_534_v2_.f18) < ((self).f15): d_536_v4_})
                d_536_v4_ = ((d_555_v21_)[d_532_v0_] if (d_532_v0_) in (d_555_v21_) else d_536_v4_)
                d_544_v12_ = d_544_v12_
        elif True:
            d_556_v22_: _dafny.Set
            d_556_v22_ = _dafny.Set({538})
            (self).f4 = default__.safeModuloInt((0) - (len((d_556_v22_) - (d_556_v22_))), self.f4)
            d_532_v0_ = (self).f16
            d_557_v23_: _dafny.Array
            def lambda20_(d_558_i1_):
                return (self).f16

            init11_ = lambda20_
            nw79_ = _dafny.Array(None, 16)
            for i0_11_ in range(nw79_.length(0)):
                nw79_[i0_11_] = init11_(i0_11_)
            d_557_v23_ = nw79_
            d_559_v24_: _dafny.MultiSet
            d_559_v24_ = _dafny.MultiSet([d_557_v23_])
            d_559_v24_ = d_559_v24_
            d_560_v25_: _dafny.Seq
            d_560_v25_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hmde"))
            d_561_v26_: C0
            nw80_ = C0()
            nw80_.ctor__(len(d_560_v25_))
            d_561_v26_ = nw80_
            d_562_v27_: D5
            d_562_v27_ = D5_DC18(d_561_v26_, (self).f16, d_532_v0_)
            d_563_v28_: _dafny.Set
            d_563_v28_ = _dafny.Set({(d_562_v27_).cf38})
            d_564_v29_: _dafny.Seq
            d_564_v29_ = _dafny.SeqWithoutIsStrInference([d_561_v26_, d_561_v26_, d_561_v26_])
            d_565_v30_: D2
            d_565_v30_ = D2_DC6(self.f4, False)
            d_566_v31_: _dafny.Map
            d_566_v31_ = _dafny.Map({len(d_560_v25_): len(d_560_v25_)})
            d_567_v32_: _dafny.Map
            d_567_v32_ = _dafny.Map({d_532_v0_: len(d_566_v31_)})
            d_568_v33_: D4
            d_568_v33_ = D4_DC12(d_560_v25_)
            d_569_v34_: _dafny.Map
            d_569_v34_ = _dafny.Map({d_568_v33_: d_532_v0_})
            d_570_v35_: _dafny.Array
            nw81_ = _dafny.Array(None, 5)
            nw81_[int(0)] = self.f17
            nw81_[int(1)] = d_561_v26_.f18
            nw81_[int(2)] = self.f17
            nw81_[int(3)] = (self).f15
            nw81_[int(4)] = len(d_569_v34_)
            d_570_v35_ = nw81_
            d_571_v36_: _dafny.Seq
            d_571_v36_ = _dafny.SeqWithoutIsStrInference([d_570_v35_, d_570_v35_, d_570_v35_, d_570_v35_, d_570_v35_])
            d_572_v37_: _dafny.Map
            d_572_v37_ = _dafny.Map({d_557_v23_: d_561_v26_.f18})
            d_573_v38_: _dafny.Array
            nw82_ = _dafny.Array(None, 23)
            nw82_[int(0)] = default__.safeDivisionInt(d_561_v26_.f18, (self).f15)
            nw82_[int(1)] = (self).f15
            nw82_[int(2)] = (self).f15
            nw82_[int(3)] = d_561_v26_.f18
            nw82_[int(4)] = len((d_563_v28_).intersection(d_563_v28_))
            nw82_[int(5)] = self.f4
            nw82_[int(6)] = self.f4
            nw82_[int(7)] = default__.fm1(globalState)
            nw82_[int(8)] = default__.safeModuloInt(len(d_564_v29_), (d_565_v30_).cf10)
            nw82_[int(9)] = default__.safeModuloInt(self.f4, self.f4)
            nw82_[int(10)] = (self).f15
            nw82_[int(11)] = ((d_567_v32_)[not(d_532_v0_)] if (not(d_532_v0_)) in (d_567_v32_) else len(d_571_v36_))
            nw82_[int(12)] = d_561_v26_.f18
            nw82_[int(13)] = d_561_v26_.f18
            nw82_[int(14)] = d_561_v26_.f18
            nw82_[int(15)] = default__.safeDivisionInt((self).f15, len(d_560_v25_))
            nw82_[int(16)] = d_561_v26_.f18
            nw82_[int(17)] = (self).f15
            nw82_[int(18)] = ((d_572_v37_)[d_557_v23_] if (d_557_v23_) in (d_572_v37_) else d_561_v26_.f18)
            nw82_[int(19)] = self.f4
            nw82_[int(20)] = self.f17
            nw82_[int(21)] = default__.safeDivisionInt(self.f17, d_561_v26_.f18)
            nw82_[int(22)] = 270
            d_573_v38_ = nw82_
            index55_ = default__.safeIndex(887, (d_573_v38_).length(0))
            (d_573_v38_)[index55_] = ((self).f15) + (self.f4)
            d_574_v39_: D6
            d_574_v39_ = D6_DC21((self).f15, (self).f16)
            index56_ = default__.safeIndex(887, (d_573_v38_).length(0))
            (d_573_v38_)[index56_] = (d_574_v39_).cf41
        d_575_v40_: _dafny.Map
        d_575_v40_ = _dafny.Map({self.f4: self.f4})
        d_576_v41_: _dafny.Seq
        d_576_v41_ = _dafny.SeqWithoutIsStrInference([d_532_v0_, d_532_v0_])
        d_577_v42_: _dafny.Seq
        d_577_v42_ = _dafny.SeqWithoutIsStrInference([self.f4, (self).f15, self.f4, self.f4, len(d_576_v41_)])
        d_578_v43_: _dafny.Map
        d_578_v43_ = _dafny.Map({len(d_577_v42_): False})
        d_579_v44_: _dafny.Seq
        d_579_v44_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xug"))
        d_580_v45_: _dafny.MultiSet
        d_580_v45_ = _dafny.MultiSet([self.f4, (self).f15])
        d_581_v46_: _dafny.MultiSet
        d_581_v46_ = _dafny.MultiSet([(d_580_v45_).cardinality])
        d_582_v47_: _dafny.Array
        nw83_ = _dafny.Array(None, 22)
        nw83_[int(0)] = len(_dafny.Map({d_532_v0_: d_575_v40_}))
        nw83_[int(1)] = len(_dafny.Map({(self).f16: self.f4}))
        nw83_[int(2)] = self.f4
        nw83_[int(3)] = 153
        nw83_[int(4)] = len(d_577_v42_)
        nw83_[int(5)] = self.f17
        nw83_[int(6)] = len(d_578_v43_)
        nw83_[int(7)] = len(_dafny.SeqWithoutIsStrInference([(self).f15 for d_583_i3_ in range(default__.abs(370))]))
        nw83_[int(8)] = self.f17
        nw83_[int(9)] = self.f4
        nw83_[int(10)] = self.f4
        nw83_[int(11)] = self.f17
        nw83_[int(12)] = self.f17
        nw83_[int(13)] = self.f4
        nw83_[int(14)] = self.f17
        nw83_[int(15)] = self.f17
        nw83_[int(16)] = (0) - (len(d_579_v44_))
        nw83_[int(17)] = self.f4
        nw83_[int(18)] = (self).f15
        nw83_[int(19)] = len(d_579_v44_)
        nw83_[int(20)] = self.f4
        nw83_[int(21)] = (d_581_v46_).cardinality
        d_582_v47_ = nw83_
        d_584_v48_: _dafny.Seq
        d_584_v48_ = _dafny.SeqWithoutIsStrInference([d_582_v47_])
        hi4_ = len(d_584_v48_)
        for d_585_i2_ in range(self.f4, hi4_):
            index57_ = default__.safeIndex(367, (d_582_v47_).length(0))
            (d_582_v47_)[index57_] = (0) - ((self.f17) - (self.f17))
            index58_ = default__.safeIndex(367, (d_582_v47_).length(0))
            (d_582_v47_)[index58_] = (0) - ((self.f4) + ((self).f15))
            (self).f17 = (self.f17) + (default__.fm1(globalState))
            d_532_v0_ = False
            d_586_v49_: str
            d_586_v49_ = _dafny.CodePoint('x')
            d_586_v49_ = d_586_v49_
        d_587_v50_: _dafny.Map
        d_587_v50_ = _dafny.Map({self.f4: d_575_v40_})
        d_588_v51_: C0
        nw84_ = C0()
        nw84_.ctor__(len(d_587_v50_))
        d_588_v51_ = nw84_
        d_589_v52_: _dafny.Array
        nw85_ = _dafny.Array(False, 8)
        d_589_v52_ = nw85_
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_589_v52_).length(0)):
            d_590_i4_: int = guard_loop_3_
            if (True) and (((0) <= (d_590_i4_)) and ((d_590_i4_) < ((d_589_v52_).length(0)))):
                def iife96_():
                    def iife98_():
                        coll58_ = _dafny.Map()
                        compr_58_: _dafny.Seq
                        for compr_58_ in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_579_v44_]))).Elements:
                            d_592_v54_: _dafny.Seq = compr_58_
                            if (d_592_v54_) in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_579_v44_]))):
                                coll58_[d_592_v54_] = (D7_DC24(_dafny.Map({False: (self).f16}))).cf44
                        return _dafny.Map(coll58_)
                    coll56_ = _dafny.Map()
                    def iife97_():
                        coll57_ = _dafny.Map()
                        compr_57_: _dafny.Seq
                        for compr_57_ in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_579_v44_]))).Elements:
                            d_592_v54_: _dafny.Seq = compr_57_
                            if (d_592_v54_) in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_579_v44_]))):
                                coll57_[d_592_v54_] = (D7_DC24(_dafny.Map({False: (self).f16}))).cf44
                        return _dafny.Map(coll57_)
                    compr_56_: _dafny.Seq
                    for compr_56_ in (iife97_()
                    ).keys.Elements:
                        d_593_v53_: _dafny.Seq = compr_56_
                        if (d_593_v53_) in (iife98_()
                        ):
                            coll56_[d_593_v53_] = _dafny.CodePoint('o')
                    return _dafny.Map(coll56_)
                (d_589_v52_)[(d_590_i4_)] = ((d_579_v44_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_591_i5_ in range(default__.abs(445))]))) in (iife96_()
                )
        d_594_v55_: _dafny.Set
        d_594_v55_ = _dafny.Set({(self).f15, self.f17})
        hi5_ = ((self).f15) + (len(d_579_v44_))
        for d_595_i6_ in range((0) - (len(d_594_v55_)), hi5_):
            index59_ = default__.safeIndex(495, (d_589_v52_).length(0))
            (d_589_v52_)[index59_] = (self).f16
            d_596_v56_: D2
            d_596_v56_ = D2_DC5((_dafny.MultiSet([d_532_v0_, d_532_v0_])).set((self).f16, default__.abs(len(d_577_v42_))))
            d_597_v57_: _dafny.Set
            d_597_v57_ = _dafny.Set({not((self).f16), d_532_v0_})
            d_598_v58_: _dafny.Map
            d_598_v58_ = _dafny.Map({(self).f15: d_597_v57_})
            d_599_v59_: D1
            d_599_v59_ = D1_DC4(d_588_v51_.f18, d_576_v41_, self.f17)
            index60_ = default__.safeIndex(495, (d_589_v52_).length(0))
            rhs57_ = ((d_532_v0_) == (d_532_v0_)) and (False)
            rhs58_ = d_532_v0_
            rhs59_ = default__.fm22((d_597_v57_).ispropersubset(((d_598_v58_)[d_588_v51_.f18] if (d_588_v51_.f18) in (d_598_v58_) else d_597_v57_)), (0) - (self.f4), self.f4, (D7_DC25(d_532_v0_, not((self).f16), self.f17)).cf46, globalState)
            rhs60_ = (self).f16
            rhs61_ = (D4_DC13(d_599_v59_, (self).f16, d_576_v41_)).cf26
            lhs25_ = d_589_v52_
            lhs26_ = default__.safeIndex(495, (d_589_v52_).length(0))
            d_532_v0_ = rhs57_
            lhs25_[lhs26_] = rhs58_
            d_596_v56_ = rhs59_
            d_532_v0_ = rhs60_
            d_576_v41_ = rhs61_
            index61_ = default__.safeIndex(131, (d_582_v47_).length(0))
            (d_582_v47_)[index61_] = (self).f15
            index62_ = default__.safeIndex(131, (d_582_v47_).length(0))
            (d_582_v47_)[index62_] = d_588_v51_.f18
            (self).f17 = (d_582_v47_)[default__.safeIndex(131, (d_582_v47_).length(0))]
            index63_ = default__.safeIndex(495, (d_589_v52_).length(0))
            (d_589_v52_)[index63_] = False
        d_600_v60_: _dafny.Array
        nw86_ = _dafny.Array(_dafny.CodePoint('D'), 13)
        d_600_v60_ = nw86_
        r0 = (d_600_v60_ if (self.f17) <= ((0) - (d_588_v51_.f18)) else d_600_v60_)
        d_601_v61_: _dafny.Set
        d_601_v61_ = _dafny.Set({d_589_v52_, d_589_v52_})
        r1 = ((d_601_v61_) - (d_601_v61_) if d_532_v0_ else d_601_v61_)
        return r0, r1

    def m12(self, p0, p1, p2, p3, globalState):
        r0: _dafny.MultiSet = _dafny.MultiSet({})
        r1: int = int(0)
        d_602_v0_: bool
        d_602_v0_ = True
        d_602_v0_ = p0
        d_603_v1_: _dafny.Array
        nw87_ = _dafny.Array(None, 5)
        nw87_[int(0)] = self.f4
        nw87_[int(1)] = default__.fm1(globalState)
        nw87_[int(2)] = default__.safeModuloInt(p2, self.f4)
        nw87_[int(3)] = self.f4
        nw87_[int(4)] = p2
        d_603_v1_ = nw87_
        index64_ = default__.safeIndex(926, (d_603_v1_).length(0))
        def iife99_():
            coll59_ = _dafny.Map()
            compr_59_: int
            for compr_59_ in (_dafny.MultiSet([p2])).Elements:
                d_604_v2_: int = compr_59_
                if (d_604_v2_) in (_dafny.MultiSet([p2])):
                    coll59_[default__.safeModuloInt(d_604_v2_, len(_dafny.Map({p1: (self).f15})))] = True
            return _dafny.Map(coll59_)
        (d_603_v1_)[index64_] = len(iife99_()
        )
        d_605_v3_: _dafny.MultiSet
        d_605_v3_ = _dafny.MultiSet([p0])
        index65_ = default__.safeIndex(926, (d_603_v1_).length(0))
        (d_603_v1_)[index65_] = (((d_605_v3_)[p1] if (p1) in (d_605_v3_) else default__.fm1(globalState)) if p1 else (self).f15)
        if not(d_602_v0_):
            d_606_v4_: C0
            nw88_ = C0()
            nw88_.ctor__(default__.fm1(globalState))
            d_606_v4_ = nw88_
            if (d_602_v0_) in (d_605_v3_):
                d_607_v5_: _dafny.Seq
                d_607_v5_ = _dafny.SeqWithoutIsStrInference([-234])
                d_608_v6_: _dafny.Seq
                d_608_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sy"))
                d_609_v7_: _dafny.MultiSet
                d_609_v7_ = _dafny.MultiSet([d_603_v1_])
                d_610_v8_: _dafny.Seq
                d_610_v8_ = _dafny.SeqWithoutIsStrInference([d_606_v4_])
                d_611_v9_: _dafny.Map
                d_611_v9_ = _dafny.Map({(d_603_v1_)[default__.safeIndex(926, (d_603_v1_).length(0))]: (d_610_v8_)[default__.safeIndex((self).f15, len(d_610_v8_))]})
                d_612_v10_: _dafny.Map
                d_612_v10_ = _dafny.Map({d_611_v9_: False})
                d_613_v11_: _dafny.Map
                d_613_v11_ = _dafny.Map({p1: 259})
                d_614_v12_: _dafny.MultiSet
                d_614_v12_ = _dafny.MultiSet([((d_613_v11_)[p1] if (p1) in (d_613_v11_) else (self).f15)])
                d_615_v13_: _dafny.Array
                nw89_ = _dafny.Array(None, 26)
                nw89_[int(0)] = p1
                nw89_[int(1)] = d_602_v0_
                nw89_[int(2)] = not(not(default__.fm0((d_607_v5_)[default__.safeIndex(d_606_v4_.f18, len(d_607_v5_))], self.f17, globalState)))
                nw89_[int(3)] = p0
                nw89_[int(4)] = not((self).fm5(d_608_v6_, p1, globalState))
                nw89_[int(5)] = p0
                nw89_[int(6)] = not(d_602_v0_)
                nw89_[int(7)] = p1
                nw89_[int(8)] = d_602_v0_
                nw89_[int(9)] = p1
                nw89_[int(10)] = (self).fm5(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_616_i0_ in range(default__.abs(406))]), (self).f16, globalState)
                nw89_[int(11)] = (-7) == (self.f17)
                nw89_[int(12)] = p0
                nw89_[int(13)] = d_602_v0_
                nw89_[int(14)] = (self).f16
                nw89_[int(15)] = not(p1)
                nw89_[int(16)] = True
                nw89_[int(17)] = (_dafny.MultiSet([d_603_v1_, d_603_v1_])).issubset(d_609_v7_)
                nw89_[int(18)] = p1
                nw89_[int(19)] = p0
                nw89_[int(20)] = ((d_612_v10_)[d_611_v9_] if (d_611_v9_) in (d_612_v10_) else d_602_v0_)
                nw89_[int(21)] = p1
                nw89_[int(22)] = (d_614_v12_).isdisjoint(d_614_v12_)
                nw89_[int(23)] = ((self).f15) != (d_606_v4_.f18)
                nw89_[int(24)] = not(p1)
                nw89_[int(25)] = p0
                d_615_v13_ = nw89_
                index66_ = default__.safeIndex(512, (d_615_v13_).length(0))
                (d_615_v13_)[index66_] = (self).f16
                d_617_v14_: _dafny.Map
                d_617_v14_ = _dafny.Map({(d_605_v3_).cardinality: (self).f16})
                index67_ = default__.safeIndex(512, (d_615_v13_).length(0))
                (d_615_v13_)[index67_] = ((self).f16) or (((self).f16 if d_602_v0_ else ((d_617_v14_)[d_606_v4_.f18] if (d_606_v4_.f18) in (d_617_v14_) else p1)))
                d_618_v15_: _dafny.Seq
                d_619_v16_: _dafny.Map
                out16_: _dafny.Seq
                out17_: _dafny.Map
                out16_, out17_ = (self).m14(globalState)
                d_618_v15_ = out16_
                d_619_v16_ = out17_
                d_620_v17_: _dafny.Seq
                d_620_v17_ = _dafny.SeqWithoutIsStrInference([d_608_v6_, d_618_v15_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "olg")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tudfubgu"))])
                d_621_v18_: _dafny.Seq
                d_621_v18_ = _dafny.SeqWithoutIsStrInference([(d_620_v17_)[default__.safeIndex(35, len(d_620_v17_))], d_618_v15_])
                d_622_v19_: _dafny.Array
                nw90_ = _dafny.Array(None, 13)
                nw90_[int(0)] = d_620_v17_
                nw90_[int(1)] = d_620_v17_
                nw90_[int(2)] = _dafny.SeqWithoutIsStrInference([d_618_v15_, d_608_v6_, d_618_v15_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qymygl"))])
                nw90_[int(3)] = d_621_v18_
                nw90_[int(4)] = (d_621_v18_) + (d_621_v18_)
                nw90_[int(5)] = d_621_v18_
                nw90_[int(6)] = d_621_v18_
                nw90_[int(7)] = d_621_v18_
                nw90_[int(8)] = (d_620_v17_) + (d_620_v17_)
                nw90_[int(9)] = d_621_v18_
                nw90_[int(10)] = d_620_v17_
                nw90_[int(11)] = d_621_v18_
                nw90_[int(12)] = d_621_v18_
                d_622_v19_ = nw90_
                index68_ = default__.safeIndex(921, (d_622_v19_).length(0))
                (d_622_v19_)[index68_] = (d_620_v17_) + (d_621_v18_)
                index69_ = default__.safeIndex(921, (d_622_v19_).length(0))
                (d_622_v19_)[index69_] = d_621_v18_
                d_623_v20_: D7
                d_623_v20_ = D7_DC24(_dafny.Map({d_602_v0_: p1}))
                d_624_v21_: _dafny.Map
                d_624_v21_ = _dafny.Map({True: (d_615_v13_)[default__.safeIndex(512, (d_615_v13_).length(0))]})
                d_625_v22_: _dafny.Array
                nw91_ = _dafny.Array(None, 10)
                nw91_[int(0)] = d_623_v20_
                nw91_[int(1)] = d_623_v20_
                nw91_[int(2)] = D7_DC24(d_624_v21_)
                nw91_[int(3)] = d_623_v20_
                nw91_[int(4)] = d_623_v20_
                nw91_[int(5)] = d_623_v20_
                nw91_[int(6)] = d_623_v20_
                nw91_[int(7)] = d_623_v20_
                nw91_[int(8)] = d_623_v20_
                nw91_[int(9)] = d_623_v20_
                d_625_v22_ = nw91_
                d_626_v23_: _dafny.Map
                d_626_v23_ = _dafny.Map({d_625_v22_: d_624_v21_})
                (d_606_v4_).f18 = len((d_626_v23_ if not((self).f16) else _dafny.Map({d_625_v22_: d_624_v21_})))
                d_627_v24_: _dafny.Set
                d_627_v24_ = _dafny.Set({not(p0), (self).f16, (d_615_v13_)[default__.safeIndex(512, (d_615_v13_).length(0))]})
                index70_ = default__.safeIndex(512, (d_615_v13_).length(0))
                (d_615_v13_)[index70_] = (((self).f16) == (p0)) in (d_627_v24_)
            elif True:
                d_628_v25_: _dafny.Array
                nw92_ = _dafny.Array(None, 17)
                nw92_[int(0)] = (self).f16
                nw92_[int(1)] = not(False)
                nw92_[int(2)] = p1
                nw92_[int(3)] = p1
                nw92_[int(4)] = p1
                nw92_[int(5)] = p0
                nw92_[int(6)] = False
                nw92_[int(7)] = p1
                nw92_[int(8)] = p0
                nw92_[int(9)] = (self).f16
                nw92_[int(10)] = p1
                nw92_[int(11)] = p1
                nw92_[int(12)] = p1
                nw92_[int(13)] = p0
                nw92_[int(14)] = (self).f16
                nw92_[int(15)] = d_602_v0_
                nw92_[int(16)] = p0
                d_628_v25_ = nw92_
                d_629_v26_: _dafny.Map
                d_629_v26_ = _dafny.Map({d_628_v25_: True})
                d_630_v27_: _dafny.Map
                d_630_v27_ = _dafny.Map({(0) - ((d_603_v1_)[default__.safeIndex(926, (d_603_v1_).length(0))]): d_602_v0_})
                d_631_v28_: _dafny.Map
                d_631_v28_ = _dafny.Map({((d_630_v27_)[d_606_v4_.f18] if (d_606_v4_.f18) in (d_630_v27_) else (self).f16): p2})
                d_629_v26_ = (d_629_v26_).set(d_628_v25_, default__.fm0(((d_631_v28_)[p0] if (p0) in (d_631_v28_) else 881), d_606_v4_.f18, globalState))
                d_632_v29_: D6
                d_632_v29_ = D6_DC21((d_603_v1_)[default__.safeIndex(926, (d_603_v1_).length(0))], not(p0))
                d_633_v30_: _dafny.Set
                d_633_v30_ = _dafny.Set({((d_632_v29_).cf41) > (p2)})
                d_634_v31_: str
                d_634_v31_ = _dafny.CodePoint('l')
                d_635_v32_: _dafny.Seq
                d_635_v32_ = _dafny.SeqWithoutIsStrInference([d_634_v31_])
                rhs62_ = d_633_v30_
                rhs63_ = (((_dafny.MultiSet(d_635_v32_)).cardinality) * (76)) <= ((self.f4) - (self.f17))
                d_633_v30_ = rhs62_
                d_602_v0_ = rhs63_
                d_634_v31_ = (d_634_v31_ if p1 else d_634_v31_)
                (self).f4 = (d_603_v1_)[default__.safeIndex(926, (d_603_v1_).length(0))]
                d_636_v33_: _dafny.Array
                nw93_ = _dafny.Array(D7.default()(), 5)
                d_636_v33_ = nw93_
                d_636_v33_ = d_636_v33_
            d_637_v34_: D0
            d_637_v34_ = D0_DC1(self.f4, default__.fm1(globalState), p1)
            if (False) == ((d_637_v34_).cf3):
                d_638_v35_: D5
                d_638_v35_ = D5_DC18(d_606_v4_, True, (default__.fm1(globalState)) > (self.f17))
                d_639_v36_: _dafny.Set
                d_639_v36_ = _dafny.Set({p1, (self).f16, (self).f16})
                rhs64_ = True
                rhs65_ = d_638_v35_
                rhs66_ = p1
                rhs67_ = (d_639_v36_) - (d_639_v36_)
                d_602_v0_ = rhs64_
                d_638_v35_ = rhs65_
                d_602_v0_ = rhs66_
                d_639_v36_ = rhs67_
                d_640_v37_: _dafny.Map
                d_640_v37_ = _dafny.Map({p2: True})
                d_641_v39_: _dafny.Seq
                def iife100_():
                    coll60_ = _dafny.Map()
                    compr_60_: int
                    for compr_60_ in _dafny.IntegerRange(927, 656):
                        d_642_v38_: int = compr_60_
                        if ((927) <= (d_642_v38_)) and ((d_642_v38_) < (656)):
                            coll60_[default__.safeModuloInt(d_642_v38_, 630)] = True
                    return _dafny.Map(coll60_)
                d_641_v39_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({(self).f15: p0}), d_640_v37_, iife100_()
                , d_640_v37_, d_640_v37_])
                d_643_v40_: _dafny.MultiSet
                d_643_v40_ = _dafny.MultiSet([d_640_v37_])
                d_644_v41_: bool
                d_645_v42_: _dafny.Seq
                d_646_v43_: _dafny.Array
                out18_: bool
                out19_: _dafny.Seq
                out20_: _dafny.Array
                out18_, out19_, out20_ = (self).m13(115, (_dafny.MultiSet(d_641_v39_)) - (d_643_v40_), d_602_v0_, globalState)
                d_644_v41_ = out18_
                d_645_v42_ = out19_
                d_646_v43_ = out20_
                d_647_v44_: _dafny.Array
                nw94_ = _dafny.Array(None, 22)
                nw94_[int(0)] = p0
                nw94_[int(1)] = d_602_v0_
                nw94_[int(2)] = p0
                nw94_[int(3)] = p0
                nw94_[int(4)] = False
                nw94_[int(5)] = False
                nw94_[int(6)] = (self).f16
                nw94_[int(7)] = False
                nw94_[int(8)] = (self).f16
                nw94_[int(9)] = d_644_v41_
                nw94_[int(10)] = d_644_v41_
                nw94_[int(11)] = d_644_v41_
                nw94_[int(12)] = p1
                nw94_[int(13)] = True
                nw94_[int(14)] = p1
                nw94_[int(15)] = p1
                nw94_[int(16)] = (self).f16
                nw94_[int(17)] = d_602_v0_
                nw94_[int(18)] = d_602_v0_
                nw94_[int(19)] = default__.fm0(d_606_v4_.f18, len(_dafny.Map({-454: False})), globalState)
                nw94_[int(20)] = d_602_v0_
                nw94_[int(21)] = (self).f16
                d_647_v44_ = nw94_
                d_648_v45_: _dafny.Seq
                d_648_v45_ = _dafny.SeqWithoutIsStrInference([d_647_v44_])
                d_649_v46_: _dafny.Seq
                d_649_v46_ = _dafny.SeqWithoutIsStrInference([(self).f15, len(_dafny.SeqWithoutIsStrInference([402, (d_603_v1_)[default__.safeIndex(926, (d_603_v1_).length(0))], len(d_648_v45_), d_606_v4_.f18, (d_603_v1_)[default__.safeIndex(926, (d_603_v1_).length(0))]]))])
                d_650_v47_: _dafny.MultiSet
                d_650_v47_ = _dafny.MultiSet([d_606_v4_.f18])
                d_651_v48_: _dafny.Seq
                d_651_v48_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(d_650_v47_).cardinality, d_606_v4_.f18, p2, d_606_v4_.f18]), d_649_v46_, d_649_v46_])
                rhs68_ = ((0) - (((d_603_v1_)[default__.safeIndex(926, (d_603_v1_).length(0))]) - (len(d_649_v46_)))) * (len((((d_651_v48_)[default__.safeIndex((d_603_v1_)[default__.safeIndex(926, (d_603_v1_).length(0))], len(d_651_v48_))]).set(default__.safeIndex((d_603_v1_)[default__.safeIndex(926, (d_603_v1_).length(0))], len((d_651_v48_)[default__.safeIndex((d_603_v1_)[default__.safeIndex(926, (d_603_v1_).length(0))], len(d_651_v48_))])), (d_605_v3_).cardinality)).set(default__.safeIndex(self.f17, len(((d_651_v48_)[default__.safeIndex((d_603_v1_)[default__.safeIndex(926, (d_603_v1_).length(0))], len(d_651_v48_))]).set(default__.safeIndex((d_603_v1_)[default__.safeIndex(926, (d_603_v1_).length(0))], len((d_651_v48_)[default__.safeIndex((d_603_v1_)[default__.safeIndex(926, (d_603_v1_).length(0))], len(d_651_v48_))])), (d_605_v3_).cardinality))), d_606_v4_.f18)))
                rhs69_ = p0
                rhs70_ = p0
                lhs27_ = self
                lhs27_.f4 = rhs68_
                d_602_v0_ = rhs69_
                d_644_v41_ = rhs70_
                d_652_v49_: _dafny.Set
                d_652_v49_ = _dafny.Set({(_dafny.MultiSet([(d_603_v1_)[default__.safeIndex(926, (d_603_v1_).length(0))], (0) - (p2)])).cardinality})
                def iife101_():
                    coll61_ = _dafny.Set()
                    compr_61_: int
                    for compr_61_ in _dafny.IntegerRange(99, 198):
                        d_653_v50_: int = compr_61_
                        if ((99) <= (d_653_v50_)) and ((d_653_v50_) < (198)):
                            coll61_ = coll61_.union(_dafny.Set([default__.safeModuloInt(d_653_v50_, ((d_650_v47_)[d_606_v4_.f18] if (d_606_v4_.f18) in (d_650_v47_) else len(_dafny.Map({d_602_v0_: self.f4}))))]))
                    return _dafny.Set(coll61_)
                (d_606_v4_).f18 = (self.f4 if (iife101_()
                ).ispropersubset(d_652_v49_) else default__.safeDivisionInt(d_606_v4_.f18, self.f4))
                d_654_v51_: str
                d_654_v51_ = _dafny.CodePoint('g')
                d_655_v52_: _dafny.Seq
                d_655_v52_ = _dafny.SeqWithoutIsStrInference([d_654_v51_])
                rhs71_ = (default__.fm1(globalState)) != (p2)
                rhs72_ = ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_656_i1_ in range(default__.abs(33))])).set(default__.safeIndex(p2, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_657_i1_ in range(default__.abs(33))]))), d_654_v51_) if d_644_v41_ else (d_655_v52_) + (d_655_v52_))
                rhs73_ = d_654_v51_
                d_602_v0_ = rhs71_
                d_655_v52_ = rhs72_
                d_654_v51_ = rhs73_
            elif True:
                d_658_v53_: _dafny.Seq
                d_658_v53_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vdp"))
                d_659_v54_: _dafny.Map
                d_659_v54_ = _dafny.Map({(d_658_v53_) != (d_658_v53_): d_602_v0_})
                d_659_v54_ = (d_659_v54_).set(p1, p0)
                d_660_v55_: _dafny.MultiSet
                d_660_v55_ = _dafny.MultiSet([d_606_v4_.f18])
                d_661_v56_: _dafny.Seq
                d_661_v56_ = _dafny.SeqWithoutIsStrInference([d_660_v55_])
                d_660_v55_ = ((d_661_v56_)[default__.safeIndex(d_606_v4_.f18, len(d_661_v56_))] if p0 else d_660_v55_)
                index71_ = default__.safeIndex(926, (d_603_v1_).length(0))
                (d_603_v1_)[index71_] = self.f17
                d_662_v57_: _dafny.Map
                d_662_v57_ = _dafny.Map({p0: (self).f15})
                d_662_v57_ = (d_662_v57_).set(not (not(d_602_v0_)) or (not(p0)), self.f4)
                d_663_v58_: _dafny.Array
                nw95_ = _dafny.Array(D1.default()(), 28)
                d_663_v58_ = nw95_
                d_664_v59_: D1
                d_664_v59_ = D1_DC4(d_606_v4_.f18, _dafny.SeqWithoutIsStrInference([d_602_v0_]), self.f17)
                index72_ = default__.safeIndex(48, (d_663_v58_).length(0))
                (d_663_v58_)[index72_] = d_664_v59_
                d_665_v60_: str
                d_665_v60_ = _dafny.CodePoint('e')
                d_666_v61_: _dafny.Array
                def lambda21_(d_667_v60_):
                    def lambda22_(d_668_i2_):
                        return d_667_v60_

                    return lambda22_

                init12_ = lambda21_(d_665_v60_)
                nw96_ = _dafny.Array(None, 13)
                for i0_12_ in range(nw96_.length(0)):
                    nw96_[i0_12_] = init12_(i0_12_)
                d_666_v61_ = nw96_
                d_669_v62_: _dafny.Seq
                d_669_v62_ = _dafny.SeqWithoutIsStrInference([d_666_v61_])
                index73_ = default__.safeIndex(48, (d_663_v58_).length(0))
                index74_ = default__.safeIndex(926, (d_603_v1_).length(0))
                rhs74_ = d_664_v59_
                rhs75_ = d_606_v4_.f18
                rhs76_ = d_665_v60_
                rhs77_ = (d_669_v62_)[default__.safeIndex((d_606_v4_.f18) + (d_606_v4_.f18), len(d_669_v62_))]
                rhs78_ = ((d_603_v1_)[default__.safeIndex(926, (d_603_v1_).length(0))]) * (default__.safeDivisionInt(p2, (self).f15))
                lhs28_ = d_663_v58_
                lhs29_ = default__.safeIndex(48, (d_663_v58_).length(0))
                lhs30_ = d_603_v1_
                lhs31_ = default__.safeIndex(926, (d_603_v1_).length(0))
                lhs28_[lhs29_] = rhs74_
                r1 = rhs75_
                d_665_v60_ = rhs76_
                d_666_v61_ = rhs77_
                lhs30_[lhs31_] = rhs78_
            d_602_v0_ = (not (not((self).f16)) or (p0)) == ((self).f16)
            d_670_v63_: C0
            nw97_ = C0()
            nw97_.ctor__(p2)
            d_670_v63_ = nw97_
        elif True:
            d_671_v64_: _dafny.Set
            d_671_v64_ = _dafny.Set({(self).f16})
            d_672_v65_: _dafny.Seq
            d_672_v65_ = _dafny.SeqWithoutIsStrInference([d_671_v64_, d_671_v64_])
            d_673_v66_: _dafny.Seq
            d_673_v66_ = _dafny.SeqWithoutIsStrInference([(d_672_v65_)[default__.safeIndex(self.f17, len(d_672_v65_))], d_671_v64_, d_671_v64_])
            d_674_v67_: _dafny.Array
            nw98_ = _dafny.Array(None, 23)
            nw98_[int(0)] = (_dafny.Set({default__.fm0((self).f15, (self).f15, globalState), d_602_v0_, True, True})) | (d_671_v64_)
            nw98_[int(1)] = d_671_v64_
            nw98_[int(2)] = d_671_v64_
            nw98_[int(3)] = (d_671_v64_) | (d_671_v64_)
            nw98_[int(4)] = d_671_v64_
            nw98_[int(5)] = _dafny.Set({not(p1), (self).f16})
            nw98_[int(6)] = _dafny.Set({d_602_v0_})
            nw98_[int(7)] = (_dafny.Set({d_602_v0_})).intersection(d_671_v64_)
            nw98_[int(8)] = (d_671_v64_) | (default__.fm17(d_602_v0_, p2, globalState))
            nw98_[int(9)] = d_671_v64_
            nw98_[int(10)] = (d_671_v64_) - (d_671_v64_)
            nw98_[int(11)] = d_671_v64_
            nw98_[int(12)] = d_671_v64_
            nw98_[int(13)] = (d_673_v66_)[default__.safeIndex((self).f15, len(d_673_v66_))]
            nw98_[int(14)] = _dafny.Set({(self).f16})
            nw98_[int(15)] = d_671_v64_
            nw98_[int(16)] = d_671_v64_
            nw98_[int(17)] = d_671_v64_
            nw98_[int(18)] = _dafny.Set({d_602_v0_})
            nw98_[int(19)] = d_671_v64_
            nw98_[int(20)] = d_671_v64_
            nw98_[int(21)] = d_671_v64_
            nw98_[int(22)] = (d_671_v64_ if p0 else _dafny.Set({p0}))
            d_674_v67_ = nw98_
            d_674_v67_ = d_674_v67_
            d_675_v68_: str
            d_675_v68_ = _dafny.CodePoint('t')
            d_676_v69_: _dafny.MultiSet
            d_676_v69_ = _dafny.MultiSet([d_675_v68_, _dafny.CodePoint('i'), d_675_v68_])
            d_677_v70_: _dafny.Map
            d_677_v70_ = _dafny.Map({_dafny.MultiSet([not(not(p0)), (self).f16]): len(_dafny.Map({d_676_v69_: self.f4}))})
            d_677_v70_ = ((d_677_v70_) | (d_677_v70_)) | (d_677_v70_)
            d_678_v71_: _dafny.Map
            d_678_v71_ = _dafny.Map({self.f4: 378})
            d_679_v72_: _dafny.Seq
            d_679_v72_ = _dafny.SeqWithoutIsStrInference([(self).f15, len(d_678_v71_)])
            d_680_v73_: _dafny.Map
            d_680_v73_ = _dafny.Map({default__.fm1(globalState): d_603_v1_})
            d_681_v74_: int
            d_682_v75_: int
            d_683_v76_: _dafny.Array
            d_684_v77_: bool
            out21_: int
            out22_: int
            out23_: _dafny.Array
            out24_: bool
            out21_, out22_, out23_, out24_ = default__.m0(_dafny.Map({(d_603_v1_)[default__.safeIndex(926, (d_603_v1_).length(0))]: p0}), d_679_v72_, d_680_v73_, globalState)
            d_681_v74_ = out21_
            d_682_v75_ = out22_
            d_683_v76_ = out23_
            d_684_v77_ = out24_
            d_685_v78_: _dafny.Array
            nw99_ = _dafny.Array(_dafny.Array(None, 0), 28)
            d_685_v78_ = nw99_
            d_686_v79_: _dafny.Array
            nw100_ = _dafny.Array(False, 8)
            d_686_v79_ = nw100_
            index75_ = default__.safeIndex(924, (d_685_v78_).length(0))
            (d_685_v78_)[index75_] = d_686_v79_
            d_687_v80_: _dafny.Seq
            d_687_v80_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nxtwfsnhl"))
            d_688_v81_: _dafny.Set
            d_688_v81_ = _dafny.Set({d_681_v74_})
            d_689_v82_: D1
            d_689_v82_ = D1_DC4(self.f17, _dafny.SeqWithoutIsStrInference([d_602_v0_]), (self).f15)
            d_690_v83_: D4
            d_690_v83_ = D4_DC13(d_689_v82_, (self).f16, _dafny.SeqWithoutIsStrInference([d_684_v77_, d_602_v0_]))
            d_691_v84_: _dafny.Map
            d_691_v84_ = _dafny.Map({(d_687_v80_)[default__.safeIndex(len((d_690_v83_).cf26), len(d_687_v80_))]: p1})
            index76_ = default__.safeIndex(924, (d_685_v78_).length(0))
            nw101_ = _dafny.Array(None, 22)
            nw101_[int(0)] = (self.f17) >= (d_682_v75_)
            nw101_[int(1)] = p1
            nw101_[int(2)] = p0
            nw101_[int(3)] = d_684_v77_
            nw101_[int(4)] = (d_605_v3_).issubset(d_605_v3_)
            nw101_[int(5)] = d_684_v77_
            nw101_[int(6)] = not((self).f16)
            nw101_[int(7)] = d_602_v0_
            nw101_[int(8)] = (self).f16
            nw101_[int(9)] = not(p1)
            nw101_[int(10)] = (d_602_v0_) or (p1)
            nw101_[int(11)] = not((d_687_v80_) <= (d_687_v80_))
            nw101_[int(12)] = (d_687_v80_) <= (d_687_v80_)
            nw101_[int(13)] = not(False)
            nw101_[int(14)] = d_602_v0_
            nw101_[int(15)] = (p0) and (d_602_v0_)
            nw101_[int(16)] = d_602_v0_
            nw101_[int(17)] = (d_688_v81_).issubset(_dafny.Set({self.f17, p2}))
            nw101_[int(18)] = p1
            nw101_[int(19)] = not((p0 if False else d_602_v0_))
            nw101_[int(20)] = ((d_691_v84_)[d_675_v68_] if (d_675_v68_) in (d_691_v84_) else d_602_v0_)
            nw101_[int(21)] = not(p1)
            (d_685_v78_)[index76_] = nw101_
            d_692_v85_: _dafny.Map
            d_692_v85_ = _dafny.Map({p2: d_675_v68_})
            d_693_v86_: _dafny.Map
            d_693_v86_ = _dafny.Map({(d_603_v1_)[default__.safeIndex(926, (d_603_v1_).length(0))]: _dafny.Set({p0, p0})})
            d_692_v85_ = (d_692_v85_).set(len(d_693_v86_), (d_687_v80_)[default__.safeIndex(d_682_v75_, len(d_687_v80_))])
        if (self).f16:
            d_694_v87_: C0
            nw102_ = C0()
            nw102_.ctor__(self.f4)
            d_694_v87_ = nw102_
            d_695_v88_: _dafny.Map
            d_695_v88_ = _dafny.Map({self.f4: (d_694_v87_ if False else d_694_v87_)})
            d_695_v88_ = (d_695_v88_).set(d_694_v87_.f18, d_694_v87_)
            (d_694_v87_).f18 = (d_603_v1_)[default__.safeIndex(926, (d_603_v1_).length(0))]
            d_696_v89_: _dafny.Set
            d_696_v89_ = _dafny.Set({p1, p1})
            d_697_v90_: _dafny.Map
            d_697_v90_ = _dafny.Map({self.f17: len((d_696_v89_) - (d_696_v89_))})
            d_697_v90_ = (d_697_v90_).set(((self).f15) - (self.f17), ((0) - ((d_603_v1_)[default__.safeIndex(926, (d_603_v1_).length(0))]) if d_602_v0_ else default__.fm1(globalState)))
            d_602_v0_ = d_602_v0_
            index77_ = default__.safeIndex(926, (d_603_v1_).length(0))
            (d_603_v1_)[index77_] = default__.safeModuloInt((self).f15, (0) - (p2))
        elif True:
            if not((self).f16):
                d_698_v91_: C0
                nw103_ = C0()
                nw103_.ctor__((d_603_v1_)[default__.safeIndex(926, (d_603_v1_).length(0))])
                d_698_v91_ = nw103_
                r1 = d_698_v91_.f18
                d_699_v92_: _dafny.Array
                nw104_ = _dafny.Array(_dafny.Array(None, 0), 18)
                d_699_v92_ = nw104_
                index78_ = default__.safeIndex(543, (d_699_v92_).length(0))
                (d_699_v92_)[index78_] = p3
                d_700_v93_: D8
                d_700_v93_ = D8_DC30(762, p3, p1)
                index79_ = default__.safeIndex(543, (d_699_v92_).length(0))
                (d_699_v92_)[index79_] = (d_700_v93_).cf55
                d_701_v94_: _dafny.Set
                d_701_v94_ = _dafny.Set({self.f17, len(default__.fm23(not(True), globalState)), p2})
                d_602_v0_ = (d_701_v94_).issubset(d_701_v94_)
                d_602_v0_ = not(default__.fm0((self.f17) - (self.f17), (0) - (self.f4), globalState))
            elif True:
                d_602_v0_ = not (d_602_v0_) or (p1)
                d_602_v0_ = not(True)
                d_602_v0_ = p0
                d_702_v95_: _dafny.Seq
                d_702_v95_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "epthots"))
                d_703_v96_: _dafny.MultiSet
                d_703_v96_ = _dafny.MultiSet([len(d_702_v95_)])
                d_704_v97_: D2
                d_704_v97_ = D2_DC6((0) - ((d_703_v96_).cardinality), d_602_v0_)
                d_705_v98_: _dafny.Map
                def iife102_(_pat_let20_0):
                    def iife103_(d_706_dt__update__tmp_h0_):
                        def iife104_(_pat_let21_0):
                            def iife105_(d_707_dt__update_hcf10_h0_):
                                return D2_DC6(d_707_dt__update_hcf10_h0_, (d_706_dt__update__tmp_h0_).cf11)
                            return iife105_(_pat_let21_0)
                        return iife104_(-370)
                    return iife103_(_pat_let20_0)
                d_705_v98_ = _dafny.Map({iife102_(d_704_v97_): p0})
                d_705_v98_ = d_705_v98_
                d_708_v100_: _dafny.Seq
                d_708_v100_ = _dafny.SeqWithoutIsStrInference([d_702_v95_])
                d_709_v102_: C0
                nw105_ = C0()
                def iife106_():
                    def iife108_():
                        coll64_ = _dafny.Map()
                        compr_64_: _dafny.Seq
                        for compr_64_ in (_dafny.SeqWithoutIsStrInference([(d_708_v100_)[default__.safeIndex(len(d_702_v95_), len(d_708_v100_))]])).Elements:
                            d_710_v99_: _dafny.Seq = compr_64_
                            if (d_710_v99_) in (_dafny.SeqWithoutIsStrInference([(d_708_v100_)[default__.safeIndex(len(d_702_v95_), len(d_708_v100_))]])):
                                coll64_[d_710_v99_] = False
                        return _dafny.Map(coll64_)
                    coll62_ = _dafny.Set()
                    def iife107_():
                        coll63_ = _dafny.Map()
                        compr_63_: _dafny.Seq
                        for compr_63_ in (_dafny.SeqWithoutIsStrInference([(d_708_v100_)[default__.safeIndex(len(d_702_v95_), len(d_708_v100_))]])).Elements:
                            d_710_v99_: _dafny.Seq = compr_63_
                            if (d_710_v99_) in (_dafny.SeqWithoutIsStrInference([(d_708_v100_)[default__.safeIndex(len(d_702_v95_), len(d_708_v100_))]])):
                                coll63_[d_710_v99_] = False
                        return _dafny.Map(coll63_)
                    compr_62_: _dafny.Seq
                    for compr_62_ in (iife107_()
                    ).keys.Elements:
                        d_711_v101_: _dafny.Seq = compr_62_
                        if (d_711_v101_) in (iife108_()
                        ):
                            coll62_ = coll62_.union(_dafny.Set([d_711_v101_]))
                    return _dafny.Set(coll62_)
                nw105_.ctor__(len(iife106_()
                ))
                d_709_v102_ = nw105_
            if d_602_v0_:
                (self).f17 = default__.fm1(globalState)
                d_712_v103_: _dafny.Seq
                d_712_v103_ = _dafny.SeqWithoutIsStrInference([self.f4])
                d_713_v104_: _dafny.Map
                d_713_v104_ = _dafny.Map({len(d_712_v103_): d_712_v103_})
                d_714_v106_: D3
                def iife109_():
                    coll65_ = _dafny.Set()
                    compr_65_: int
                    for compr_65_ in (d_713_v104_).keys.Elements:
                        d_715_v105_: int = compr_65_
                        if (d_715_v105_) in (d_713_v104_):
                            coll65_ = coll65_.union(_dafny.Set([default__.safeModuloInt(d_715_v105_, len(_dafny.Map({False: _dafny.SeqWithoutIsStrInference([536])})))]))
                    return _dafny.Set(coll65_)
                d_714_v106_ = D3_DC7(iife109_()
)
                d_716_v107_: _dafny.MultiSet
                d_716_v107_ = _dafny.MultiSet([d_714_v106_])
                (self).f17 = (((d_716_v107_)[d_714_v106_] if (d_714_v106_) in (d_716_v107_) else (0) - ((d_603_v1_)[default__.safeIndex(926, (d_603_v1_).length(0))]))) * ((714 if not(True) else len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gxd")))))
                d_717_v108_: _dafny.Array
                nw106_ = _dafny.Array(False, 14)
                d_717_v108_ = nw106_
                index80_ = default__.safeIndex(107, (d_717_v108_).length(0))
                (d_717_v108_)[index80_] = True
                index81_ = default__.safeIndex(107, (d_717_v108_).length(0))
                (d_717_v108_)[index81_] = not(p0)
                d_718_v109_: _dafny.Map
                d_718_v109_ = _dafny.Map({False: (d_603_v1_)[default__.safeIndex(926, (d_603_v1_).length(0))]})
                d_718_v109_ = (d_718_v109_).set(d_602_v0_, self.f4)
                index82_ = default__.safeIndex(107, (d_717_v108_).length(0))
                (d_717_v108_)[index82_] = (self).f16
            elif True:
                d_719_v110_: _dafny.Map
                d_719_v110_ = _dafny.Map({(self).f16: (self).f16})
                d_719_v110_ = (d_719_v110_).set(p1, p0)
                d_720_v112_: _dafny.Array
                def lambda23_(d_721_v0_):
                    def lambda24_(d_722_i3_):
                        def iife110_():
                            coll66_ = _dafny.Set()
                            compr_66_: int
                            for compr_66_ in _dafny.IntegerRange(588, 85):
                                d_723_v111_: int = compr_66_
                                if ((588) <= (d_723_v111_)) and ((d_723_v111_) < (85)):
                                    coll66_ = coll66_.union(_dafny.Set([(d_723_v111_) * ((self).f15)]))
                            return _dafny.Set(coll66_)
                        return (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_721_v0_])): self.f17})) | (_dafny.Map({self.f4: (0) - (len(iife110_()
                        ))}))

                    return lambda24_

                init13_ = lambda23_(d_602_v0_)
                nw107_ = _dafny.Array(None, 23)
                for i0_13_ in range(nw107_.length(0)):
                    nw107_[i0_13_] = init13_(i0_13_)
                d_720_v112_ = nw107_
                d_724_v113_: _dafny.Map
                d_724_v113_ = _dafny.Map({d_602_v0_: 918})
                d_725_v114_: _dafny.Seq
                d_725_v114_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "crsfw"))
                d_726_v115_: _dafny.Map
                d_726_v115_ = _dafny.Map({len(d_724_v113_): len(d_725_v114_)})
                index83_ = default__.safeIndex(453, (d_720_v112_).length(0))
                (d_720_v112_)[index83_] = (d_726_v115_) | (d_726_v115_)
                index84_ = default__.safeIndex(453, (d_720_v112_).length(0))
                (d_720_v112_)[index84_] = d_726_v115_
                d_602_v0_ = p0
                d_727_v116_: _dafny.MultiSet
                d_727_v116_ = _dafny.MultiSet([self.f17])
                d_602_v0_ = ((d_727_v116_)).ispropersubset(d_727_v116_)
                d_728_v117_: _dafny.Seq
                d_728_v117_ = _dafny.SeqWithoutIsStrInference([self.f4])
                d_729_v118_: _dafny.Seq
                d_729_v118_ = _dafny.SeqWithoutIsStrInference([(d_603_v1_)[default__.safeIndex(926, (d_603_v1_).length(0))], p2, (d_603_v1_)[default__.safeIndex(926, (d_603_v1_).length(0))], (d_728_v117_)[default__.safeIndex(self.f4, len(d_728_v117_))], self.f4])
                r1 = ((d_724_v113_)[p1] if (p1) in (d_724_v113_) else default__.safeModuloInt(len(d_728_v117_), (self).f15))
            d_730_v119_: C0
            nw108_ = C0()
            nw108_.ctor__((d_603_v1_)[default__.safeIndex(926, (d_603_v1_).length(0))])
            d_730_v119_ = nw108_
            d_731_v120_: _dafny.Seq
            d_731_v120_ = _dafny.SeqWithoutIsStrInference([p2])
            d_732_v121_: _dafny.Seq
            d_732_v121_ = _dafny.SeqWithoutIsStrInference([d_731_v120_, d_731_v120_, d_731_v120_, d_731_v120_])
            d_733_v122_: _dafny.Map
            d_733_v122_ = _dafny.Map({True: p0})
            d_734_v123_: _dafny.Map
            d_734_v123_ = _dafny.Map({p1: (d_732_v121_)[default__.safeIndex(len(d_733_v122_), len(d_732_v121_))]})
            d_734_v123_ = (d_734_v123_).set(p1, d_731_v120_)
            d_735_v124_: _dafny.Seq
            d_735_v124_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "up"))
            d_736_v125_: str
            d_736_v125_ = _dafny.CodePoint('c')
            (self).f17 = len((((d_735_v124_) + (d_735_v124_)).set(default__.safeIndex((d_603_v1_)[default__.safeIndex(926, (d_603_v1_).length(0))], len((d_735_v124_) + (d_735_v124_))), d_736_v125_) if p1 else d_735_v124_))
        d_737_v126_: _dafny.Array
        nw109_ = _dafny.Array(_dafny.MultiSet({}), 17)
        d_737_v126_ = nw109_
        guard_loop_4_: int
        for guard_loop_4_ in _dafny.IntegerRange(0, (d_737_v126_).length(0)):
            d_738_i4_: int = guard_loop_4_
            if (True) and (((0) <= (d_738_i4_)) and ((d_738_i4_) < ((d_737_v126_).length(0)))):
                (d_737_v126_)[(d_738_i4_)] = ((d_605_v3_) - (d_605_v3_)).set(p1, default__.abs(len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wtjosb"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tplfnqn"))))))
        d_739_v127_: _dafny.Seq
        d_739_v127_ = _dafny.SeqWithoutIsStrInference([(self).f15])
        d_739_v127_ = (_dafny.SeqWithoutIsStrInference([(D1_DC4(self.f4, _dafny.SeqWithoutIsStrInference([(self).f16]), p2)).cf6 for d_740_i5_ in range(default__.abs(388))])) + ((d_739_v127_) + (_dafny.SeqWithoutIsStrInference([self.f17, (d_603_v1_)[default__.safeIndex(926, (d_603_v1_).length(0))]])))
        d_741_v128_: _dafny.MultiSet
        d_741_v128_ = _dafny.MultiSet([d_739_v127_, (d_739_v127_).set(default__.safeIndex(len(d_739_v127_), len(d_739_v127_)), self.f4), d_739_v127_, d_739_v127_, d_739_v127_])
        r0 = (d_741_v128_).intersection(default__.fm24(globalState))
        r1 = self.f17
        return r0, r1

    def m13(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: _dafny.Seq = _dafny.Seq({})
        r2: _dafny.Array = _dafny.Array(None, 0)
        d_742_v0_: _dafny.Array
        def lambda25_(d_743_p2_):
            def lambda26_(d_744_i0_):
                return (_dafny.SeqWithoutIsStrInference([(self).f16])) + (_dafny.SeqWithoutIsStrInference([d_743_p2_]))

            return lambda26_

        init14_ = lambda25_(p2)
        nw110_ = _dafny.Array(None, 18)
        for i0_14_ in range(nw110_.length(0)):
            nw110_[i0_14_] = init14_(i0_14_)
        d_742_v0_ = nw110_
        d_745_v1_: _dafny.Seq
        d_745_v1_ = _dafny.SeqWithoutIsStrInference([p2])
        index85_ = default__.safeIndex(331, (d_742_v0_).length(0))
        (d_742_v0_)[index85_] = d_745_v1_
        index86_ = default__.safeIndex(331, (d_742_v0_).length(0))
        (d_742_v0_)[index86_] = ((_dafny.SeqWithoutIsStrInference([not((self).f16)]) if p2 else d_745_v1_)) + (_dafny.SeqWithoutIsStrInference([(self).f16]))
        d_746_v2_: _dafny.Array
        nw111_ = _dafny.Array(int(0), 24)
        d_746_v2_ = nw111_
        index87_ = default__.safeIndex(274, (d_746_v2_).length(0))
        (d_746_v2_)[index87_] = 327
        d_747_v3_: _dafny.Seq
        d_747_v3_ = _dafny.SeqWithoutIsStrInference([(0) - ((self).f15), p0, (self.f4) + (690), p0, 36])
        index88_ = default__.safeIndex(274, (d_746_v2_).length(0))
        (d_746_v2_)[index88_] = (d_747_v3_)[default__.safeIndex(self.f4, len(d_747_v3_))]
        d_748_v4_: _dafny.Array
        d_748_v4_ = d_746_v2_
        d_746_v2_ = (d_748_v4_)
        d_749_v6_: _dafny.Array
        def lambda27_(d_750_i2_):
            def iife111_():
                coll67_ = _dafny.Map()
                compr_67_: str
                for compr_67_ in (_dafny.Set({_dafny.CodePoint('n'), _dafny.CodePoint('d')})).Elements:
                    d_751_v5_: str = compr_67_
                    if (d_751_v5_) in (_dafny.Set({_dafny.CodePoint('n'), _dafny.CodePoint('d')})):
                        coll67_[d_751_v5_] = (self).f16
                return _dafny.Map(coll67_)
            return _dafny.Map({len(iife111_()
            ): D6_DC21((self).f15, (self).f16)})

        init15_ = lambda27_
        nw112_ = _dafny.Array(None, 17)
        for i0_15_ in range(nw112_.length(0)):
            nw112_[i0_15_] = init15_(i0_15_)
        d_749_v6_ = nw112_
        def iife112_():
            coll68_ = _dafny.Map()
            compr_68_: int
            for compr_68_ in (_dafny.MultiSet([self.f17])).Elements:
                d_753_v7_: int = compr_68_
                if (d_753_v7_) in (_dafny.MultiSet([self.f17])):
                    coll68_[default__.safeModuloInt(d_753_v7_, self.f4)] = D6_DC21(self.f17, (self).f16)
            return _dafny.Map(coll68_)
        _ingredientsd_0 = []
        guard_loop_5_: int
        for guard_loop_5_ in _dafny.IntegerRange(0, (d_749_v6_).length(0)):
            d_752_i1_: int = guard_loop_5_
            if (True) and (((0) <= (d_752_i1_)) and ((d_752_i1_) < ((d_749_v6_).length(0)))):
                _ingredientsd_0.append((d_749_v6_, int(d_752_i1_), ((_dafny.Map({len(_dafny.Map({_dafny.SeqWithoutIsStrInference([(self).f15]): p2})): D6_DC21((d_746_v2_)[default__.safeIndex(274, (d_746_v2_).length(0))], p2)})) | (_dafny.Map({self.f4: D6_DC21(self.f4, True)}))) | ((iife112_()
                ) | (_dafny.Map({p0: D6_DC21(p0, p2)})))))
        for _tupd_0 in _ingredientsd_0:
            _tupd_0[0][_tupd_0[1]] = _tupd_0[2]
        d_754_v8_: _dafny.Seq
        d_754_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qrgi"))
        d_755_v9_: D6
        d_755_v9_ = D6_DC21(len(d_754_v8_), (self).f16)
        d_756_v10_: _dafny.Array
        nw113_ = _dafny.Array(_dafny.CodePoint('D'), 13)
        d_756_v10_ = nw113_
        d_757_v11_: _dafny.Map
        d_757_v11_ = _dafny.Map({(d_755_v9_).cf42: d_756_v10_})
        d_757_v11_ = (d_757_v11_).set(p2, d_756_v10_)
        d_758_v12_: _dafny.Array
        def lambda28_(d_759_v8_):
            def lambda29_(d_760_i4_):
                return d_759_v8_

            return lambda29_

        init16_ = lambda28_(d_754_v8_)
        nw114_ = _dafny.Array(None, 9)
        for i0_16_ in range(nw114_.length(0)):
            nw114_[i0_16_] = init16_(i0_16_)
        d_758_v12_ = nw114_
        guard_loop_6_: int
        for guard_loop_6_ in _dafny.IntegerRange(0, (d_758_v12_).length(0)):
            d_761_i3_: int = guard_loop_6_
            if (True) and (((0) <= (d_761_i3_)) and ((d_761_i3_) < ((d_758_v12_).length(0)))):
                (d_758_v12_)[(d_761_i3_)] = ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_762_i5_ in range(default__.abs(886))])) + (d_754_v8_)) + (d_754_v8_)
        r0 = ((d_746_v2_)[default__.safeIndex(274, (d_746_v2_).length(0))]) <= ((d_746_v2_)[default__.safeIndex(274, (d_746_v2_).length(0))])
        d_763_v13_: _dafny.Set
        d_763_v13_ = _dafny.Set({(self).f16, p2})
        d_764_v14_: _dafny.Seq
        d_764_v14_ = _dafny.SeqWithoutIsStrInference([d_763_v13_])
        r1 = (d_764_v14_) + (d_764_v14_)
        d_765_v15_: _dafny.Map
        d_765_v15_ = _dafny.Map({False: d_746_v2_})
        d_766_v16_: _dafny.Array
        nw115_ = _dafny.Array(None, 29)
        nw115_[int(0)] = d_746_v2_
        nw115_[int(1)] = d_746_v2_
        nw115_[int(2)] = d_746_v2_
        nw115_[int(3)] = d_746_v2_
        nw115_[int(4)] = d_746_v2_
        nw115_[int(5)] = d_746_v2_
        nw115_[int(6)] = d_746_v2_
        nw115_[int(7)] = d_746_v2_
        nw115_[int(8)] = d_746_v2_
        nw115_[int(9)] = d_746_v2_
        nw115_[int(10)] = d_746_v2_
        nw115_[int(11)] = d_746_v2_
        nw115_[int(12)] = d_746_v2_
        nw115_[int(13)] = d_746_v2_
        nw115_[int(14)] = d_746_v2_
        nw115_[int(15)] = d_746_v2_
        nw115_[int(16)] = d_746_v2_
        nw115_[int(17)] = d_746_v2_
        nw115_[int(18)] = d_746_v2_
        nw115_[int(19)] = d_746_v2_
        nw115_[int(20)] = ((d_765_v15_)[p2] if (p2) in (d_765_v15_) else d_746_v2_)
        nw115_[int(21)] = d_746_v2_
        nw115_[int(22)] = d_746_v2_
        nw115_[int(23)] = d_746_v2_
        nw115_[int(24)] = d_746_v2_
        nw115_[int(25)] = d_746_v2_
        nw115_[int(26)] = d_746_v2_
        nw115_[int(27)] = d_746_v2_
        nw115_[int(28)] = d_746_v2_
        d_766_v16_ = nw115_
        d_767_v17_: D11
        d_767_v17_ = D11_DC34(d_766_v16_)
        r2 = (d_767_v17_).cf60
        return r0, r1, r2

    def m14(self, globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r1: _dafny.Map = _dafny.Map({})
        d_768_v0_: _dafny.Map
        d_768_v0_ = _dafny.Map({not ((self).f16) or ((self).f16): (self).f16})
        d_768_v0_ = (d_768_v0_).set((self).f16, (self).f16)
        d_769_v1_: str
        d_769_v1_ = _dafny.CodePoint('e')
        d_770_v2_: _dafny.MultiSet
        d_770_v2_ = _dafny.MultiSet([d_769_v1_, d_769_v1_])
        d_771_v3_: _dafny.Seq
        d_771_v3_ = _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([d_769_v1_])) | (d_770_v2_)])
        d_771_v3_ = (d_771_v3_ if default__.fm0(self.f4, (self).f15, globalState) else (d_771_v3_).set(default__.safeIndex(230, len(d_771_v3_)), d_770_v2_))
        if not ((self).f16) or ((self).f16):
            d_772_v4_: _dafny.Map
            d_772_v4_ = _dafny.Map({self.f4: (self).f16})
            d_773_v5_: _dafny.Seq
            d_773_v5_ = _dafny.SeqWithoutIsStrInference([-367])
            d_774_v6_: _dafny.Array
            def lambda30_(d_775_i0_):
                return default__.safeModuloInt(d_775_i0_, self.f4)

            init17_ = lambda30_
            nw116_ = _dafny.Array(None, 22)
            for i0_17_ in range(nw116_.length(0)):
                nw116_[i0_17_] = init17_(i0_17_)
            d_774_v6_ = nw116_
            d_776_v7_: _dafny.Map
            d_776_v7_ = _dafny.Map({(self).f15: d_774_v6_})
            d_777_v8_: int
            d_778_v9_: int
            d_779_v10_: _dafny.Array
            d_780_v11_: bool
            out25_: int
            out26_: int
            out27_: _dafny.Array
            out28_: bool
            out25_, out26_, out27_, out28_ = default__.m0(d_772_v4_, d_773_v5_, d_776_v7_, globalState)
            d_777_v8_ = out25_
            d_778_v9_ = out26_
            d_779_v10_ = out27_
            d_780_v11_ = out28_
            d_781_v12_: _dafny.Array
            nw117_ = _dafny.Array(_dafny.Array(None, 0), 11)
            d_781_v12_ = nw117_
            d_782_v13_: _dafny.Array
            def lambda31_(d_783_v11_):
                def lambda32_(d_784_i1_):
                    return d_783_v11_

                return lambda32_

            init18_ = lambda31_(d_780_v11_)
            nw118_ = _dafny.Array(None, 18)
            for i0_18_ in range(nw118_.length(0)):
                nw118_[i0_18_] = init18_(i0_18_)
            d_782_v13_ = nw118_
            index89_ = default__.safeIndex(825, (d_781_v12_).length(0))
            (d_781_v12_)[index89_] = d_782_v13_
            index90_ = default__.safeIndex(825, (d_781_v12_).length(0))
            (d_781_v12_)[index90_] = d_782_v13_
            d_785_v14_: _dafny.Seq
            d_785_v14_ = _dafny.SeqWithoutIsStrInference([d_773_v5_])
            index91_ = default__.safeIndex(954, (d_774_v6_).length(0))
            (d_774_v6_)[index91_] = len(((d_785_v14_)[default__.safeIndex(self.f4, len(d_785_v14_))] if d_780_v11_ else d_773_v5_))
            index92_ = default__.safeIndex(954, (d_774_v6_).length(0))
            (d_774_v6_)[index92_] = (0) - ((default__.fm1(globalState)) * ((self.f17) + ((self).f15)))
            d_777_v8_ = self.f4
            d_786_v15_: _dafny.MultiSet
            d_786_v15_ = _dafny.MultiSet([(self).f16, d_780_v11_, d_780_v11_])
            index93_ = default__.safeIndex(954, (d_774_v6_).length(0))
            (d_774_v6_)[index93_] = ((d_773_v5_)[default__.safeIndex(self.f4, len(d_773_v5_))]) * (((d_786_v15_)[True] if (True) in (d_786_v15_) else (0) - ((d_774_v6_)[default__.safeIndex(954, (d_774_v6_).length(0))])))
        elif True:
            d_787_v16_: _dafny.Seq
            d_787_v16_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "huaylrld"))
            (self).f17 = len(((d_787_v16_) + (d_787_v16_)) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dv"))) + (d_787_v16_)))
            d_788_v17_: _dafny.Map
            d_788_v17_ = _dafny.Map({len(d_787_v16_): default__.fm0(self.f4, -153, globalState)})
            d_789_v18_: _dafny.Seq
            d_789_v18_ = _dafny.SeqWithoutIsStrInference([941])
            d_790_v19_: _dafny.Array
            nw119_ = _dafny.Array(None, 10)
            nw119_[int(0)] = self.f4
            nw119_[int(1)] = self.f17
            nw119_[int(2)] = len(d_788_v17_)
            nw119_[int(3)] = self.f17
            nw119_[int(4)] = self.f17
            nw119_[int(5)] = self.f17
            nw119_[int(6)] = len((d_787_v16_).set(default__.safeIndex(self.f4, len(d_787_v16_)), d_769_v1_))
            nw119_[int(7)] = (d_789_v18_)[default__.safeIndex((self).f15, len(d_789_v18_))]
            nw119_[int(8)] = (self).f15
            nw119_[int(9)] = self.f17
            d_790_v19_ = nw119_
            d_791_v20_: _dafny.Set
            d_791_v20_ = _dafny.Set({d_790_v19_})
            (self).f17 = (0) - (default__.safeDivisionInt((len(d_791_v20_)) * ((self).f15), 863))
            d_792_v21_: C0
            nw120_ = C0()
            nw120_.ctor__(default__.fm1(globalState))
            d_792_v21_ = nw120_
            (d_792_v21_).f18 = d_792_v21_.f18
            d_793_v22_: _dafny.Array
            nw121_ = _dafny.Array(_dafny.MultiSet({}), 16)
            d_793_v22_ = nw121_
            d_794_v23_: _dafny.Seq
            d_794_v23_ = _dafny.SeqWithoutIsStrInference([d_793_v22_, d_793_v22_, d_793_v22_, d_793_v22_])
            d_795_v24_: _dafny.Array
            nw122_ = _dafny.Array(None, 10)
            nw122_[int(0)] = (d_794_v23_)[default__.safeIndex((0) - ((0) - (self.f4)), len(d_794_v23_))]
            nw122_[int(1)] = d_793_v22_
            nw122_[int(2)] = d_793_v22_
            nw122_[int(3)] = d_793_v22_
            nw122_[int(4)] = (d_794_v23_)[default__.safeIndex((default__.fm25(globalState)).cardinality, len(d_794_v23_))]
            nw122_[int(5)] = d_793_v22_
            nw122_[int(6)] = d_793_v22_
            nw122_[int(7)] = d_793_v22_
            nw122_[int(8)] = d_793_v22_
            nw122_[int(9)] = d_793_v22_
            d_795_v24_ = nw122_
            d_796_v25_: _dafny.Array
            nw123_ = _dafny.Array(False, 27)
            d_796_v25_ = nw123_
            d_797_v26_: _dafny.Map
            d_797_v26_ = _dafny.Map({(self).f16: (d_788_v17_) | (d_788_v17_)})
            d_798_v27_: _dafny.MultiSet
            d_798_v27_ = _dafny.MultiSet([not((self).fm5(d_787_v16_, (self).f16, globalState)), (self).f16])
            d_799_v28_: D5
            d_799_v28_ = D5_DC17(d_798_v27_, (self).f16)
            d_800_v29_: _dafny.Map
            d_800_v29_ = _dafny.Map({self.f4: d_799_v28_})
            d_801_v30_: _dafny.Set
            d_801_v30_ = _dafny.Set({(self).f16})
            d_802_v31_: _dafny.Map
            d_802_v31_ = _dafny.Map({d_789_v18_: d_792_v21_.f18})
            def iife113_(_pat_let22_0):
                def iife114_(d_803_dt__update__tmp_h0_):
                    def iife115_(_pat_let23_0):
                        def iife116_(d_804_dt__update_hcf8_h0_):
                            return D1_DC4((d_803_dt__update__tmp_h0_).cf6, (d_803_dt__update__tmp_h0_).cf7, d_804_dt__update_hcf8_h0_)
                        return iife116_(_pat_let23_0)
                    return iife115_(-59)
                return iife114_(_pat_let22_0)
            rhs79_ = d_795_v24_
            rhs80_ = d_792_v21_.f18
            rhs81_ = d_796_v25_
            rhs82_ = default__.safeModuloInt((iife113_(D1_DC4(len(d_800_v29_), _dafny.SeqWithoutIsStrInference([(self).f16]), len(d_801_v30_)))).cf8, (self.f4) * (self.f4))
            rhs83_ = (default__.fm26(len(_dafny.Map({not(not((self).f16)): (self).f15})), False, (d_802_v31_).set(d_789_v18_, d_792_v21_.f18), (self).f16, globalState)) | ((d_797_v26_) | (d_797_v26_))
            lhs32_ = d_792_v21_
            lhs33_ = self
            d_795_v24_ = rhs79_
            lhs32_.f18 = rhs80_
            d_796_v25_ = rhs81_
            lhs33_.f17 = rhs82_
            d_797_v26_ = rhs83_
        d_805_v32_: bool
        d_805_v32_ = True
        d_805_v32_ = False
        d_806_v33_: C0
        nw124_ = C0()
        nw124_.ctor__(default__.fm1(globalState))
        d_806_v33_ = nw124_
        rhs84_ = (d_806_v33_ if not((self).f16) else d_806_v33_)
        rhs85_ = not((self).f16)
        d_806_v33_ = rhs84_
        d_805_v32_ = rhs85_
        hi6_ = d_806_v33_.f18
        for d_807_i2_ in range(140, hi6_):
            d_808_v34_: C0
            nw125_ = C0()
            nw125_.ctor__(len(_dafny.Set({(self).f16})))
            d_808_v34_ = nw125_
            d_809_v35_: C0
            nw126_ = C0()
            nw126_.ctor__(d_806_v33_.f18)
            d_809_v35_ = nw126_
            d_810_v36_: _dafny.Array
            nw127_ = _dafny.Array(False, 6)
            d_810_v36_ = nw127_
            d_811_v37_: _dafny.Map
            d_811_v37_ = _dafny.Map({self.f4: d_805_v32_})
            index94_ = default__.safeIndex(892, (d_810_v36_).length(0))
            (d_810_v36_)[index94_] = (d_806_v33_.f18) in (d_811_v37_)
            index95_ = default__.safeIndex(892, (d_810_v36_).length(0))
            (d_810_v36_)[index95_] = (self).f16
            d_812_v38_: _dafny.Array
            nw128_ = _dafny.Array(_dafny.MultiSet({}), 14)
            d_812_v38_ = nw128_
            nw129_ = _dafny.Array(_dafny.MultiSet({}), 14)
            d_812_v38_ = nw129_
        r0 = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eqo")) if d_805_v32_ else _dafny.SeqWithoutIsStrInference([d_769_v1_ for d_813_i3_ in range(default__.abs(741))]))
        d_814_v39_: _dafny.Map
        d_814_v39_ = _dafny.Map({d_805_v32_: _dafny.Map({(self).f15: _dafny.Set({self.f17})})})
        d_815_v40_: _dafny.Set
        d_815_v40_ = _dafny.Set({self.f17})
        d_816_v41_: _dafny.Map
        d_816_v41_ = _dafny.Map({d_806_v33_.f18: d_815_v40_})
        r1 = ((d_814_v39_)[(True) or (d_805_v32_)] if ((True) or (d_805_v32_)) in (d_814_v39_) else d_816_v41_)
        return r0, r1


class C2:
    def  __init__(self):
        self.f13: bool = False
        self._f14: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    def ctor__(self, f13, f14):
        (self).f13 = f13
        (self)._f14 = f14

    def m10(self, p0, p1, p2, globalState):
        r0: bool = False
        d_817_i0_: int
        d_817_i0_ = 0
        with _dafny.label("5"):
            while self.f13:
                with _dafny.c_label("5"):
                    if (d_817_i0_) >= (100):
                        raise _dafny.Break("5")
                    d_817_i0_ = (d_817_i0_) + (1)
                    d_818_v0_: _dafny.Map
                    d_818_v0_ = _dafny.Map({(self).f14: self.f13})
                    d_819_v1_: _dafny.Array
                    nw130_ = _dafny.Array(int(0), 3)
                    d_819_v1_ = nw130_
                    d_820_v2_: _dafny.Map
                    d_820_v2_ = _dafny.Map({len(p1): d_819_v1_})
                    d_821_v3_: int
                    d_822_v4_: int
                    d_823_v5_: _dafny.Array
                    d_824_v6_: bool
                    out29_: int
                    out30_: int
                    out31_: _dafny.Array
                    out32_: bool
                    out29_, out30_, out31_, out32_ = default__.m0(d_818_v0_, _dafny.SeqWithoutIsStrInference([(self).f14 for d_825_i1_ in range(default__.abs(-965))]), d_820_v2_, globalState)
                    d_821_v3_ = out29_
                    d_822_v4_ = out30_
                    d_823_v5_ = out31_
                    d_824_v6_ = out32_
                    d_826_v7_: _dafny.Map
                    d_826_v7_ = _dafny.Map({d_824_v6_: d_822_v4_})
                    d_827_v8_: D1
                    d_827_v8_ = D1_DC3((d_826_v7_) | (d_826_v7_))
                    d_827_v8_ = d_827_v8_
                    d_828_v9_: _dafny.Seq
                    d_828_v9_ = _dafny.SeqWithoutIsStrInference([not(d_824_v6_), d_824_v6_])
                    d_828_v9_ = (d_828_v9_) + (d_828_v9_)
                    d_823_v5_ = (d_819_v1_ if d_824_v6_ else d_819_v1_)
                    pass
            pass
        d_829_v10_: C0
        nw131_ = C0()
        nw131_.ctor__((self).f14)
        d_829_v10_ = nw131_
        d_830_v11_: _dafny.Seq
        d_830_v11_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tsjfgtnuq")))])
        d_831_v12_: _dafny.Map
        d_831_v12_ = _dafny.Map({(self).f14: len(d_830_v11_)})
        hi7_ = d_829_v10_.f18
        for d_832_i2_ in range((0) - ((len(d_831_v12_)) - (d_829_v10_.f18)), hi7_):
            d_833_v13_: _dafny.Array
            nw132_ = _dafny.Array(_dafny.Set({}), 24)
            d_833_v13_ = nw132_
            d_833_v13_ = d_833_v13_
            d_834_v14_: str
            d_834_v14_ = _dafny.CodePoint('o')
            d_835_v15_: D0
            d_835_v15_ = D0_DC2(D0_DC1(d_829_v10_.f18, 560, self.f13))
            d_836_v16_: D0
            d_836_v16_ = D0_DC0(d_830_v11_)
            d_837_v17_: _dafny.MultiSet
            d_837_v17_ = _dafny.MultiSet([d_835_v15_, d_835_v15_, D0_DC2(d_836_v16_), d_835_v15_, d_835_v15_])
            d_838_v18_: _dafny.Map
            d_838_v18_ = _dafny.Map({(((d_837_v17_).set(d_835_v15_, default__.abs(d_829_v10_.f18))).set(d_835_v15_, default__.abs(d_829_v10_.f18)) if default__.fm0(len(_dafny.SeqWithoutIsStrInference([d_834_v14_])), (self).f14, globalState) else d_837_v17_): (_dafny.SeqWithoutIsStrInference([len(p1) for d_839_i3_ in range(default__.abs(429))])) <= (d_830_v11_)})
            d_838_v18_ = (d_838_v18_).set(d_837_v17_, (d_834_v14_) not in (p1))
            d_840_v19_: _dafny.Map
            d_840_v19_ = _dafny.Map({(self).f14: default__.fm0(default__.fm1(globalState), d_829_v10_.f18, globalState)})
            d_841_v20_: _dafny.Map
            d_841_v20_ = _dafny.Map({d_829_v10_.f18: (default__.fm0(len(d_840_v19_), d_832_i2_, globalState)) == (self.f13)})
            d_840_v19_ = d_840_v19_
            d_830_v11_ = d_830_v11_
        d_842_v21_: _dafny.Map
        d_842_v21_ = _dafny.Map({False: (self).f14})
        d_843_v22_: _dafny.Map
        d_843_v22_ = _dafny.Map({not (self.f13) or (self.f13): (len(d_842_v21_)) + (default__.fm1(globalState))})
        d_842_v21_ = (d_842_v21_).set(self.f13, 339)
        d_844_i4_: int
        d_844_i4_ = 0
        with _dafny.label("6"):
            while self.f13:
                with _dafny.c_label("6"):
                    if (d_844_i4_) >= (100):
                        raise _dafny.Break("6")
                    d_844_i4_ = (d_844_i4_) + (1)
                    (self).f13 = default__.fm0((self).f14, d_829_v10_.f18, globalState)
                    d_845_v23_: _dafny.Seq
                    d_845_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oxbr"))
                    d_845_v23_ = p1
                    d_846_v24_: _dafny.Seq
                    d_846_v24_ = _dafny.SeqWithoutIsStrInference([self.f13, self.f13, self.f13])
                    d_847_v25_: _dafny.Map
                    d_847_v25_ = _dafny.Map({d_829_v10_.f18: d_846_v24_})
                    d_848_v26_: D12
                    d_848_v26_ = D12_DC36(d_847_v25_)
                    d_849_v27_: str
                    d_849_v27_ = _dafny.CodePoint('x')
                    d_850_v28_: _dafny.MultiSet
                    d_850_v28_ = _dafny.MultiSet([_dafny.CodePoint('w'), d_849_v27_, d_849_v27_])
                    d_847_v25_ = ((d_848_v26_).cf64).set(d_829_v10_.f18, (d_846_v24_).set(default__.safeIndex((d_850_v28_).cardinality, len(d_846_v24_)), self.f13))
                    d_851_v29_: _dafny.Array
                    nw133_ = _dafny.Array(False, 15)
                    d_851_v29_ = nw133_
                    index96_ = default__.safeIndex(147, (d_851_v29_).length(0))
                    (d_851_v29_)[index96_] = True
                    index97_ = default__.safeIndex(147, (d_851_v29_).length(0))
                    (d_851_v29_)[index97_] = self.f13
                    pass
            pass
        d_852_v30_: _dafny.Array
        def lambda33_(d_853_p1_):
            def lambda34_(d_854_i5_):
                return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "djlrxsrs"))) + (d_853_p1_)

            return lambda34_

        init19_ = lambda33_(p1)
        nw134_ = _dafny.Array(None, 1)
        for i0_19_ in range(nw134_.length(0)):
            nw134_[i0_19_] = init19_(i0_19_)
        d_852_v30_ = nw134_
        index98_ = default__.safeIndex(108, (d_852_v30_).length(0))
        (d_852_v30_)[index98_] = p1
        index99_ = default__.safeIndex(108, (d_852_v30_).length(0))
        (d_852_v30_)[index99_] = (p1) + (p1)
        d_855_v31_: str
        d_855_v31_ = _dafny.CodePoint('k')
        d_856_v32_: _dafny.Array
        nw135_ = _dafny.Array(None, 4)
        nw135_[int(0)] = default__.fm0(d_829_v10_.f18, d_829_v10_.f18, globalState)
        nw135_[int(1)] = self.f13
        nw135_[int(2)] = True
        nw135_[int(3)] = self.f13
        d_856_v32_ = nw135_
        d_857_v33_: D3
        d_857_v33_ = D3_DC8(d_855_v31_, d_856_v32_)
        d_858_v34_: _dafny.Array
        nw136_ = _dafny.Array(None, 17)
        nw136_[int(0)] = d_855_v31_
        nw136_[int(1)] = d_855_v31_
        nw136_[int(2)] = d_855_v31_
        nw136_[int(3)] = default__.fm27(self.f13, globalState)
        nw136_[int(4)] = _dafny.CodePoint('f')
        nw136_[int(5)] = d_855_v31_
        nw136_[int(6)] = d_855_v31_
        nw136_[int(7)] = _dafny.CodePoint('s')
        nw136_[int(8)] = d_855_v31_
        nw136_[int(9)] = d_855_v31_
        nw136_[int(10)] = d_855_v31_
        nw136_[int(11)] = d_855_v31_
        nw136_[int(12)] = (d_857_v33_).cf13
        nw136_[int(13)] = d_855_v31_
        nw136_[int(14)] = d_855_v31_
        nw136_[int(15)] = d_855_v31_
        nw136_[int(16)] = d_855_v31_
        d_858_v34_ = nw136_
        d_859_v35_: _dafny.Set
        d_859_v35_ = _dafny.Set({d_858_v34_, d_858_v34_, d_858_v34_})
        r0 = (d_859_v35_).isdisjoint(d_859_v35_)
        return r0

    @property
    def f14(self):
        return self._f14

class C3(T0):
    def  __init__(self):
        self._f4: int = int(0)
        self._f11: int = int(0)
        self._f12: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f4(self):
        return self._f4
    @f4.setter
    def f4(self, value):
        self._f4 = value
    def ctor__(self, f11, f12, f4):
        (self)._f11 = f11
        (self)._f12 = f12
        (self).f4 = f4

    def fm5(self, p0, p1, globalState):
        source8_ = D2_DC6(self.f4, not(False))
        if source8_.is_DC6:
            d_860___mcc_h0_ = source8_.cf10
            d_861___mcc_h1_ = source8_.cf11
            d_862_cf11_ = d_861___mcc_h1_
            d_863_cf10_ = d_860___mcc_h0_
            return d_862_cf11_
        elif True:
            d_864___mcc_h2_ = source8_.cf9
            d_865_cf9_ = d_864___mcc_h2_
            return False

    def fm11(self, p0, p1, globalState):
        return (0) - ((self).f11)

    def fm12(self, globalState):
        return _dafny.SeqWithoutIsStrInference([False, not((not(True) if not(False) else False))])

    def m3(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        if default__.fm0((p1) * (p1), (self).f11, globalState):
            d_866_v0_: _dafny.Seq
            d_866_v0_ = _dafny.SeqWithoutIsStrInference([p0])
            d_867_v1_: _dafny.Seq
            d_867_v1_ = _dafny.SeqWithoutIsStrInference([(d_866_v0_)[default__.safeIndex(self.f4, len(d_866_v0_))], p0])
            d_868_v2_: _dafny.Map
            d_868_v2_ = _dafny.Map({True: d_867_v1_})
            d_869_v3_: D1
            d_869_v3_ = D1_DC4(p1, d_866_v0_, (self).f11)
            d_870_v4_: _dafny.Array
            nw137_ = _dafny.Array(None, 27)
            nw137_[int(0)] = (d_868_v2_) | (d_868_v2_)
            nw137_[int(1)] = d_868_v2_
            nw137_[int(2)] = (d_868_v2_) | (d_868_v2_)
            nw137_[int(3)] = _dafny.Map({p0: d_866_v0_})
            nw137_[int(4)] = d_868_v2_
            nw137_[int(5)] = (d_868_v2_ if p0 else default__.fm13(globalState))
            nw137_[int(6)] = default__.fm13(globalState)
            nw137_[int(7)] = (d_868_v2_) | (_dafny.Map({p0: d_866_v0_}))
            nw137_[int(8)] = (d_868_v2_) | (d_868_v2_)
            nw137_[int(9)] = d_868_v2_
            nw137_[int(10)] = _dafny.Map({p3: _dafny.SeqWithoutIsStrInference([p3, p2])})
            nw137_[int(11)] = (d_868_v2_) | (d_868_v2_)
            nw137_[int(12)] = d_868_v2_
            nw137_[int(13)] = d_868_v2_
            nw137_[int(14)] = d_868_v2_
            nw137_[int(15)] = (d_868_v2_) | (default__.fm13(globalState))
            nw137_[int(16)] = d_868_v2_
            nw137_[int(17)] = default__.fm13(globalState)
            nw137_[int(18)] = (d_868_v2_) | (_dafny.Map({p2: (d_869_v3_).cf7}))
            nw137_[int(19)] = d_868_v2_
            nw137_[int(20)] = (d_868_v2_) | (d_868_v2_)
            nw137_[int(21)] = d_868_v2_
            nw137_[int(22)] = d_868_v2_
            nw137_[int(23)] = d_868_v2_
            nw137_[int(24)] = d_868_v2_
            nw137_[int(25)] = d_868_v2_
            nw137_[int(26)] = (d_868_v2_) | (d_868_v2_)
            d_870_v4_ = nw137_
            index100_ = default__.safeIndex(472, (d_870_v4_).length(0))
            (d_870_v4_)[index100_] = d_868_v2_
            index101_ = default__.safeIndex(472, (d_870_v4_).length(0))
            (d_870_v4_)[index101_] = (((d_868_v2_).set(p0, d_866_v0_)) | (d_868_v2_)) | (default__.fm13(globalState))
            r0 = not (p2) or (default__.fm0(len((self).f12), (0) - (-892), globalState))
            d_871_v5_: _dafny.Map
            d_871_v5_ = _dafny.Map({330: p1})
            d_872_v6_: str
            d_872_v6_ = _dafny.CodePoint('l')
            d_873_v7_: _dafny.Map
            d_873_v7_ = _dafny.Map({default__.fm14(p1, len(d_871_v5_), p2, d_872_v6_, globalState): p3})
            d_874_v8_: _dafny.Array
            def lambda35_(d_875_i0_):
                return False

            init20_ = lambda35_
            nw138_ = _dafny.Array(None, 20)
            for i0_20_ in range(nw138_.length(0)):
                nw138_[i0_20_] = init20_(i0_20_)
            d_874_v8_ = nw138_
            d_876_v9_: _dafny.MultiSet
            d_876_v9_ = _dafny.MultiSet([d_874_v8_])
            r0 = not((((d_873_v7_)[_dafny.SeqWithoutIsStrInference([(self).f11])] if (_dafny.SeqWithoutIsStrInference([(self).f11])) in (d_873_v7_) else p2)) or ((_dafny.MultiSet([d_874_v8_])).isdisjoint(d_876_v9_)))
            d_877_v10_: _dafny.Map
            d_877_v10_ = _dafny.Map({p0: d_874_v8_})
            (self).f4 = len(((d_877_v10_) | (d_877_v10_)) | (d_877_v10_))
            d_878_v11_: _dafny.Seq
            d_878_v11_ = _dafny.SeqWithoutIsStrInference([-643, self.f4, (self).f11])
            r0 = (len(d_878_v11_)) >= ((self).f11)
        elif True:
            if p2:
                d_879_v12_: str
                d_879_v12_ = _dafny.CodePoint('y')
                d_880_v13_: _dafny.Map
                d_880_v13_ = _dafny.Map({d_879_v12_: p2})
                d_881_v14_: C1
                nw139_ = C1()
                nw139_.ctor__((len(d_880_v13_) if p3 else (self).f11), (self).f11, (self).f11, not(p3))
                d_881_v14_ = nw139_
                d_882_v15_: D4
                d_882_v15_ = D4_DC12((self).f12)
                d_883_v16_: _dafny.Set
                d_883_v16_ = _dafny.Set({p2})
                d_884_v17_: _dafny.Array
                nw140_ = _dafny.Array(None, 28)
                nw140_[int(0)] = (self).f12
                nw140_[int(1)] = (self).f12
                nw140_[int(2)] = (d_882_v15_).cf23
                nw140_[int(3)] = (self).f12
                nw140_[int(4)] = (self).f12
                nw140_[int(5)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a"))
                nw140_[int(6)] = ((self).f12).set(default__.safeIndex(len(d_883_v16_), len((self).f12)), d_879_v12_)
                nw140_[int(7)] = (self).f12
                nw140_[int(8)] = (self).f12
                nw140_[int(9)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tnnmncn"))
                nw140_[int(10)] = (self).f12
                nw140_[int(11)] = _dafny.SeqWithoutIsStrInference([d_879_v12_ for d_885_i1_ in range(default__.abs(671))])
                nw140_[int(12)] = (self).f12
                nw140_[int(13)] = (self).f12
                nw140_[int(14)] = _dafny.SeqWithoutIsStrInference([d_879_v12_ for d_886_i2_ in range(default__.abs(467))])
                nw140_[int(15)] = (self).f12
                nw140_[int(16)] = default__.fm23(p2, globalState)
                nw140_[int(17)] = (self).f12
                nw140_[int(18)] = (self).f12
                nw140_[int(19)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e"))
                nw140_[int(20)] = (self).f12
                nw140_[int(21)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gkj"))
                nw140_[int(22)] = (self).f12
                nw140_[int(23)] = (self).f12
                nw140_[int(24)] = (self).f12
                nw140_[int(25)] = (self).f12
                nw140_[int(26)] = (self).f12
                nw140_[int(27)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "beajp"))
                d_884_v17_ = nw140_
                d_887_v18_: D8
                d_887_v18_ = D8_DC30(p1, d_884_v17_, p3)
                d_888_v19_: _dafny.Seq
                d_888_v19_ = _dafny.SeqWithoutIsStrInference([d_887_v18_])
                d_889_v20_: _dafny.MultiSet
                d_889_v20_ = _dafny.MultiSet([len(d_888_v19_)])
                d_890_v21_: T2
                nw141_ = C1()
                nw141_.ctor__(d_881_v14_.f17, (d_889_v20_).cardinality, ((self).f11) * (p1), p3)
                d_890_v21_ = nw141_
                d_890_v21_ = d_890_v21_
                r0 = (d_890_v21_).f16
                r0 = p3
                d_891_v22_: C0
                nw142_ = C0()
                nw142_.ctor__(default__.safeModuloInt(self.f4, len(_dafny.SeqWithoutIsStrInference([d_879_v12_ for d_892_i3_ in range(default__.abs(13))]))))
                d_891_v22_ = nw142_
            elif True:
                rhs86_ = p0
                rhs87_ = p2
                r0 = rhs86_
                r0 = rhs87_
                (self).f4 = p1
                d_893_v23_: C1
                nw143_ = C1()
                nw143_.ctor__((449 if p0 else p1), p1, (self).f11, p3)
                d_893_v23_ = nw143_
                d_894_v24_: _dafny.Array
                d_895_v25_: _dafny.Set
                out33_: _dafny.Array
                out34_: _dafny.Set
                out33_, out34_ = (d_893_v23_).m11(globalState)
                d_894_v24_ = out33_
                d_895_v25_ = out34_
                d_896_v26_: _dafny.Array
                def lambda36_(d_897_p3_):
                    def lambda37_(d_898_i4_):
                        return not(d_897_p3_)

                    return lambda37_

                init21_ = lambda36_(p3)
                nw144_ = _dafny.Array(None, 1)
                for i0_21_ in range(nw144_.length(0)):
                    nw144_[i0_21_] = init21_(i0_21_)
                d_896_v26_ = nw144_
                index102_ = default__.safeIndex(460, (d_896_v26_).length(0))
                (d_896_v26_)[index102_] = p0
                d_899_v27_: _dafny.Seq
                d_899_v27_ = _dafny.SeqWithoutIsStrInference([p3])
                index103_ = default__.safeIndex(460, (d_896_v26_).length(0))
                (d_896_v26_)[index103_] = ((d_899_v27_)[default__.safeIndex((self).f11, len(d_899_v27_))]) == (p2)
            d_900_v28_: D1
            d_900_v28_ = D1_DC4(len((self).f12), _dafny.SeqWithoutIsStrInference([p0, p0]), (self).f11)
            d_901_v29_: _dafny.Array
            nw145_ = _dafny.Array(None, 1)
            nw145_[int(0)] = p1
            d_901_v29_ = nw145_
            index104_ = default__.safeIndex(42, (d_901_v29_).length(0))
            (d_901_v29_)[index104_] = (self).f11
            d_902_v30_: _dafny.Array
            def lambda38_(d_903_i5_):
                return (self).f12

            init22_ = lambda38_
            nw146_ = _dafny.Array(None, 18)
            for i0_22_ in range(nw146_.length(0)):
                nw146_[i0_22_] = init22_(i0_22_)
            d_902_v30_ = nw146_
            d_904_v31_: D8
            d_904_v31_ = D8_DC30(p1, d_902_v30_, p0)
            d_905_v32_: _dafny.Seq
            d_905_v32_ = _dafny.SeqWithoutIsStrInference([(d_904_v31_).cf56])
            d_906_v33_: str
            d_906_v33_ = _dafny.CodePoint('e')
            d_907_v34_: T1
            nw147_ = C1()
            nw147_.ctor__(self.f4, (self).f11, p1, p3)
            d_907_v34_ = nw147_
            d_908_v35_: _dafny.Map
            d_908_v35_ = _dafny.Map({d_907_v34_: d_905_v32_})
            d_909_v36_: _dafny.Map
            d_909_v36_ = _dafny.Map({False: (self).f11})
            d_910_v37_: _dafny.Seq
            d_910_v37_ = _dafny.SeqWithoutIsStrInference([d_905_v32_, _dafny.SeqWithoutIsStrInference([p3]), _dafny.SeqWithoutIsStrInference([(d_907_v34_).fm5(((self).f12).set(default__.safeIndex(404, len((self).f12)), d_906_v33_), p2, globalState), p2]), d_905_v32_])
            d_911_v38_: _dafny.Array
            nw148_ = _dafny.Array(None, 26)
            nw148_[int(0)] = d_905_v32_
            nw148_[int(1)] = default__.fm2(d_906_v33_, 204, p1, globalState)
            nw148_[int(2)] = (_dafny.SeqWithoutIsStrInference([p2, p0])) + (d_905_v32_)
            nw148_[int(3)] = ((d_908_v35_)[d_907_v34_] if (d_907_v34_) in (d_908_v35_) else d_905_v32_)
            nw148_[int(4)] = d_905_v32_
            nw148_[int(5)] = d_905_v32_
            nw148_[int(6)] = d_905_v32_
            nw148_[int(7)] = d_905_v32_
            nw148_[int(8)] = d_905_v32_
            nw148_[int(9)] = (d_905_v32_) + (d_905_v32_)
            nw148_[int(10)] = d_905_v32_
            nw148_[int(11)] = d_905_v32_
            nw148_[int(12)] = d_905_v32_
            nw148_[int(13)] = d_905_v32_
            nw148_[int(14)] = default__.fm2(default__.fm27(p0, globalState), d_907_v34_.f4, len(d_909_v36_), globalState)
            nw148_[int(15)] = d_905_v32_
            nw148_[int(16)] = d_905_v32_
            nw148_[int(17)] = d_905_v32_
            nw148_[int(18)] = d_905_v32_
            nw148_[int(19)] = d_905_v32_
            nw148_[int(20)] = (d_905_v32_) + ((d_905_v32_).set(default__.safeIndex(len((self).f12), len(d_905_v32_)), p3))
            nw148_[int(21)] = _dafny.SeqWithoutIsStrInference([p2])
            nw148_[int(22)] = (d_905_v32_) + (d_905_v32_)
            nw148_[int(23)] = (d_910_v37_)[default__.safeIndex(self.f4, len(d_910_v37_))]
            nw148_[int(24)] = d_905_v32_
            nw148_[int(25)] = d_905_v32_
            d_911_v38_ = nw148_
            index105_ = default__.safeIndex(897, (d_911_v38_).length(0))
            (d_911_v38_)[index105_] = d_905_v32_
            pat_let_tv22_ = d_910_v37_
            pat_let_tv23_ = d_910_v37_
            index106_ = default__.safeIndex(42, (d_901_v29_).length(0))
            index107_ = default__.safeIndex(897, (d_911_v38_).length(0))
            def iife117_(_pat_let24_0):
                def iife118_(d_912_dt__update__tmp_h0_):
                    def iife119_(_pat_let25_0):
                        def iife120_(d_913_dt__update_hcf7_h0_):
                            return D1_DC4((d_912_dt__update__tmp_h0_).cf6, d_913_dt__update_hcf7_h0_, (d_912_dt__update__tmp_h0_).cf8)
                        return iife120_(_pat_let25_0)
                    return iife119_((pat_let_tv22_)[default__.safeIndex((self).f11, len(pat_let_tv23_))])
                return iife118_(_pat_let24_0)
            rhs88_ = len(_dafny.Map({(0) - (len((self).f12)): d_907_v34_.f4}))
            rhs89_ = iife117_(d_900_v28_)
            rhs90_ = default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ciegxmx"))), self.f4)
            rhs91_ = _dafny.SeqWithoutIsStrInference([(d_905_v32_)[default__.safeIndex(d_907_v34_.f4, len(d_905_v32_))]])
            lhs34_ = self
            lhs35_ = d_901_v29_
            lhs36_ = default__.safeIndex(42, (d_901_v29_).length(0))
            lhs37_ = d_911_v38_
            lhs38_ = default__.safeIndex(897, (d_911_v38_).length(0))
            lhs34_.f4 = rhs88_
            d_900_v28_ = rhs89_
            lhs35_[lhs36_] = rhs90_
            lhs37_[lhs38_] = rhs91_
            d_914_v39_: _dafny.Map
            d_914_v39_ = _dafny.Map({False: p3})
            d_915_v40_: C2
            nw149_ = C2()
            nw149_.ctor__(p3, len(d_914_v39_))
            d_915_v40_ = nw149_
            d_916_v41_: _dafny.Map
            d_916_v41_ = _dafny.Map({d_915_v40_: p2})
            d_917_v42_: _dafny.Seq
            d_917_v42_ = _dafny.SeqWithoutIsStrInference([d_916_v41_])
            d_918_v43_: _dafny.Map
            d_918_v43_ = _dafny.Map({547: d_917_v42_})
            d_919_v44_: D13
            d_919_v44_ = D13_DC39(d_917_v42_)
            d_920_v45_: _dafny.Seq
            d_920_v45_ = _dafny.SeqWithoutIsStrInference([d_907_v34_.f4])
            d_921_v46_: _dafny.Map
            d_921_v46_ = _dafny.Map({((d_918_v43_)[len(((self).f12).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([p0])), len((self).f12)), d_906_v33_))] if (len(((self).f12).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([p0])), len((self).f12)), d_906_v33_))) in (d_918_v43_) else (d_919_v44_).cf69): (d_920_v45_)[default__.safeIndex((d_915_v40_).f14, len(d_920_v45_))]})
            d_922_v47_: _dafny.MultiSet
            d_922_v47_ = _dafny.MultiSet([not(p3)])
            d_921_v46_ = (d_921_v46_).set(((d_918_v43_)[d_907_v34_.f4] if (d_907_v34_.f4) in (d_918_v43_) else d_917_v42_), (0) - (((d_915_v40_).f14) * (((d_922_v47_)[d_915_v40_.f13] if (d_915_v40_.f13) in (d_922_v47_) else p1))))
            d_923_v48_: _dafny.Array
            nw150_ = _dafny.Array(False, 4)
            d_923_v48_ = nw150_
            d_924_v49_: _dafny.Map
            d_924_v49_ = _dafny.Map({d_923_v48_: d_919_v44_})
            d_925_v50_: _dafny.Set
            d_925_v50_ = _dafny.Set({(self).f11})
            d_926_v51_: _dafny.Map
            d_926_v51_ = _dafny.Map({d_915_v40_.f13: d_925_v50_})
            rhs92_ = (d_907_v34_).fm5((self).f12, p2, globalState)
            rhs93_ = (0) - ((self).fm11(d_926_v51_, p3, globalState))
            rhs94_ = d_924_v49_
            rhs95_ = len(_dafny.Map({d_907_v34_.f4: (0) - ((d_920_v45_)[default__.safeIndex(self.f4, len(d_920_v45_))])}))
            lhs39_ = d_915_v40_
            lhs40_ = d_907_v34_
            lhs41_ = self
            lhs39_.f13 = rhs92_
            lhs40_.f4 = rhs93_
            d_924_v49_ = rhs94_
            lhs41_.f4 = rhs95_
            d_927_v52_: D5
            d_927_v52_ = D5_DC17(d_922_v47_, p2)
            d_928_v53_: _dafny.Map
            d_928_v53_ = _dafny.Map({(d_901_v29_)[default__.safeIndex(42, (d_901_v29_).length(0))]: d_927_v52_})
            d_929_v54_: _dafny.Set
            d_929_v54_ = _dafny.Set({d_928_v53_})
            d_930_v56_: _dafny.Map
            def iife121_():
                coll69_ = _dafny.Map()
                compr_69_: int
                for compr_69_ in _dafny.IntegerRange(726, -809):
                    d_931_v55_: int = compr_69_
                    if ((726) <= (d_931_v55_)) and ((d_931_v55_) < (-809)):
                        coll69_[default__.safeDivisionInt(d_931_v55_, (d_915_v40_).f14)] = d_927_v52_
                return _dafny.Map(coll69_)
            d_930_v56_ = _dafny.Map({iife121_()
            : d_907_v34_.f4})
            def iife122_():
                coll70_ = _dafny.Set()
                compr_70_: _dafny.Map
                for compr_70_ in (d_930_v56_).keys.Elements:
                    d_932_v57_: _dafny.Map = compr_70_
                    if (d_932_v57_) in (d_930_v56_):
                        coll70_ = coll70_.union(_dafny.Set([d_932_v57_]))
                return _dafny.Set(coll70_)
            (d_915_v40_).f13 = (d_929_v54_).ispropersubset(iife122_()
            )
        if p0:
            d_933_v58_: _dafny.Seq
            d_933_v58_ = _dafny.SeqWithoutIsStrInference([p2, p0])
            r0 = (False if p0 else not((d_933_v58_)[default__.safeIndex(self.f4, len(d_933_v58_))]))
            d_934_v59_: str
            d_934_v59_ = _dafny.CodePoint('f')
            d_935_v60_: _dafny.Map
            d_935_v60_ = _dafny.Map({p3: p3})
            d_936_v61_: _dafny.Array
            nw151_ = _dafny.Array(None, 20)
            nw151_[int(0)] = not(p0)
            nw151_[int(1)] = not(p0)
            nw151_[int(2)] = p2
            nw151_[int(3)] = p2
            nw151_[int(4)] = default__.fm0(self.f4, self.f4, globalState)
            nw151_[int(5)] = p0
            nw151_[int(6)] = p0
            nw151_[int(7)] = p0
            nw151_[int(8)] = p0
            nw151_[int(9)] = p3
            nw151_[int(10)] = p3
            nw151_[int(11)] = ((d_935_v60_)[True] if (True) in (d_935_v60_) else p2)
            nw151_[int(12)] = p2
            nw151_[int(13)] = not(p3)
            nw151_[int(14)] = p2
            nw151_[int(15)] = p0
            nw151_[int(16)] = p3
            nw151_[int(17)] = p0
            nw151_[int(18)] = p3
            nw151_[int(19)] = p2
            d_936_v61_ = nw151_
            d_937_v62_: _dafny.Map
            d_937_v62_ = _dafny.Map({(d_934_v59_ if p2 else d_934_v59_): d_936_v61_})
            d_938_v63_: D7
            d_938_v63_ = D7_DC25(p2, True, p1)
            d_939_v64_: _dafny.MultiSet
            d_939_v64_ = _dafny.MultiSet([default__.fm1(globalState), (-906) + (self.f4), (self).f11])
            rhs96_ = ((self).fm5(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ayeurpjei")), p0, globalState) if True else (d_938_v63_).cf45)
            rhs97_ = (0) - (((d_939_v64_)[self.f4] if (self.f4) in (d_939_v64_) else (0) - ((self).f11)))
            rhs98_ = d_937_v62_
            lhs42_ = self
            r0 = rhs96_
            lhs42_.f4 = rhs97_
            d_937_v62_ = rhs98_
            d_934_v59_ = d_934_v59_
            d_940_v65_: _dafny.MultiSet
            d_940_v65_ = _dafny.MultiSet([p0])
            d_941_v66_: _dafny.Map
            d_941_v66_ = _dafny.Map({self.f4: p3})
            d_942_v67_: _dafny.Seq
            d_942_v67_ = _dafny.SeqWithoutIsStrInference([((d_940_v65_)[p3] if (p3) in (d_940_v65_) else (d_940_v65_).cardinality), len(d_941_v66_), self.f4, -895])
            d_943_v68_: D0
            d_943_v68_ = D0_DC0(d_942_v67_)
            d_944_v69_: C2
            nw152_ = C2()
            nw152_.ctor__(((d_943_v68_).cf0) != (d_942_v67_), p1)
            d_944_v69_ = nw152_
            (self).f4 = self.f4
        elif True:
            r0 = p3
            d_945_v70_: str
            d_945_v70_ = _dafny.CodePoint('u')
            d_945_v70_ = _dafny.CodePoint('y')
            d_946_v71_: _dafny.Seq
            d_946_v71_ = _dafny.SeqWithoutIsStrInference([(self).f11])
            d_947_v72_: D0
            d_947_v72_ = D0_DC0(d_946_v71_)
            d_948_v73_: _dafny.Map
            d_948_v73_ = _dafny.Map({d_947_v72_: p1})
            pat_let_tv24_ = d_946_v71_
            def iife123_(_pat_let26_0):
                def iife124_(d_949_dt__update__tmp_h1_):
                    def iife125_(_pat_let27_0):
                        def iife126_(d_950_dt__update_hcf0_h0_):
                            return D0_DC0(d_950_dt__update_hcf0_h0_)
                        return iife126_(_pat_let27_0)
                    return iife125_(pat_let_tv24_)
                return iife124_(_pat_let26_0)
            d_948_v73_ = (d_948_v73_).set(iife123_(d_947_v72_), self.f4)
            d_951_v74_: _dafny.Array
            nw153_ = _dafny.Array(_dafny.Set({}), 27)
            d_951_v74_ = nw153_
            d_951_v74_ = d_951_v74_
            d_952_v75_: _dafny.Array
            nw154_ = _dafny.Array(int(0), 11)
            d_952_v75_ = nw154_
            index108_ = default__.safeIndex(910, (d_952_v75_).length(0))
            (d_952_v75_)[index108_] = p1
            index109_ = default__.safeIndex(910, (d_952_v75_).length(0))
            (d_952_v75_)[index109_] = len((self).f12)
        hi8_ = p1
        for d_953_i6_ in range(p1, hi8_):
            rhs99_ = len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_954_i7_ in range(default__.abs(-26))])) + ((self).f12))
            rhs100_ = p3
            lhs43_ = self
            lhs43_.f4 = rhs99_
            r0 = rhs100_
            if not(True):
                d_955_v76_: _dafny.Seq
                d_955_v76_ = _dafny.SeqWithoutIsStrInference([(self).f11, (self).f11, (self).f11])
                d_956_v77_: _dafny.MultiSet
                d_956_v77_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qtj")))])
                d_957_v78_: _dafny.Seq
                d_957_v78_ = _dafny.SeqWithoutIsStrInference([len(_dafny.Set({len(d_955_v76_), (d_956_v77_).cardinality, self.f4, (self).f11, (self).f11})), (d_953_i6_) - (default__.fm1(globalState))])
                d_957_v78_ = d_955_v76_
                d_958_v79_: _dafny.Array
                nw155_ = _dafny.Array(_dafny.Map({}), 27)
                d_958_v79_ = nw155_
                d_959_v80_: _dafny.Array
                nw156_ = _dafny.Array(False, 25)
                d_959_v80_ = nw156_
                d_960_v81_: _dafny.Map
                d_960_v81_ = _dafny.Map({p2: d_959_v80_})
                index110_ = default__.safeIndex(24, (d_958_v79_).length(0))
                (d_958_v79_)[index110_] = (d_960_v81_) | (d_960_v81_)
                index111_ = default__.safeIndex(24, (d_958_v79_).length(0))
                (d_958_v79_)[index111_] = d_960_v81_
                d_961_v82_: _dafny.Map
                d_961_v82_ = _dafny.Map({p3: d_953_i6_})
                d_962_v83_: _dafny.Seq
                d_962_v83_ = _dafny.SeqWithoutIsStrInference([d_961_v82_, d_961_v82_, (d_961_v82_ if p3 else _dafny.Map({True: (self).f11}))])
                d_962_v83_ = ((_dafny.SeqWithoutIsStrInference([d_961_v82_, d_961_v82_])) + (d_962_v83_) if not((d_956_v77_).isdisjoint(d_956_v77_)) else d_962_v83_)
                d_963_v84_: _dafny.Set
                d_963_v84_ = _dafny.Set({p3, p3})
                d_964_v85_: _dafny.MultiSet
                d_964_v85_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dpauedy")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qusq"))])
                d_965_v86_: _dafny.Array
                nw157_ = _dafny.Array(None, 12)
                nw157_[int(0)] = d_963_v84_
                nw157_[int(1)] = d_963_v84_
                nw157_[int(2)] = d_963_v84_
                nw157_[int(3)] = d_963_v84_
                nw157_[int(4)] = d_963_v84_
                nw157_[int(5)] = d_963_v84_
                nw157_[int(6)] = d_963_v84_
                nw157_[int(7)] = (d_963_v84_).intersection(d_963_v84_)
                nw157_[int(8)] = default__.fm17(p0, self.f4, globalState)
                nw157_[int(9)] = (_dafny.Set({default__.fm0((d_964_v85_).cardinality, d_953_i6_, globalState)}) if False else d_963_v84_)
                nw157_[int(10)] = d_963_v84_
                nw157_[int(11)] = (d_963_v84_).intersection(d_963_v84_)
                d_965_v86_ = nw157_
                index112_ = default__.safeIndex(361, (d_965_v86_).length(0))
                (d_965_v86_)[index112_] = (d_963_v84_).intersection(d_963_v84_)
                d_966_v87_: _dafny.Array
                def lambda39_(d_967_i6_):
                    def lambda40_(d_968_i8_):
                        return default__.safeDivisionInt(d_968_i8_, d_967_i6_)

                    return lambda40_

                init23_ = lambda39_(d_953_i6_)
                nw158_ = _dafny.Array(None, 10)
                for i0_23_ in range(nw158_.length(0)):
                    nw158_[i0_23_] = init23_(i0_23_)
                d_966_v87_ = nw158_
                d_969_v88_: str
                d_969_v88_ = _dafny.CodePoint('v')
                d_970_v89_: D3
                d_970_v89_ = D3_DC8(d_969_v88_, d_959_v80_)
                d_971_v90_: _dafny.Array
                nw159_ = _dafny.Array(None, 15)
                nw159_[int(0)] = _dafny.CodePoint('l')
                nw159_[int(1)] = d_969_v88_
                nw159_[int(2)] = d_969_v88_
                nw159_[int(3)] = d_969_v88_
                nw159_[int(4)] = d_969_v88_
                nw159_[int(5)] = d_969_v88_
                nw159_[int(6)] = (d_970_v89_).cf13
                nw159_[int(7)] = d_969_v88_
                nw159_[int(8)] = ((self).f12)[default__.safeIndex((d_955_v76_)[default__.safeIndex(d_953_i6_, len(d_955_v76_))], len((self).f12))]
                nw159_[int(9)] = _dafny.CodePoint('p')
                nw159_[int(10)] = d_969_v88_
                nw159_[int(11)] = _dafny.CodePoint('f')
                nw159_[int(12)] = d_969_v88_
                nw159_[int(13)] = d_969_v88_
                nw159_[int(14)] = d_969_v88_
                d_971_v90_ = nw159_
                d_972_v91_: _dafny.Map
                d_972_v91_ = _dafny.Map({d_966_v87_: d_971_v90_})
                index113_ = default__.safeIndex(361, (d_965_v86_).length(0))
                (d_965_v86_)[index113_] = ((D14_DC42(len(d_972_v91_), p0, d_963_v84_, p1)).cf74) - (d_963_v84_)
                d_973_v92_: _dafny.Array
                nw160_ = _dafny.Array(_dafny.Array(None, 0), 13)
                d_973_v92_ = nw160_
                index114_ = default__.safeIndex(690, (d_973_v92_).length(0))
                (d_973_v92_)[index114_] = d_959_v80_
                d_974_v93_: _dafny.MultiSet
                d_974_v93_ = _dafny.MultiSet([p3])
                index115_ = default__.safeIndex(690, (d_973_v92_).length(0))
                rhs101_ = (self.f4) * ((len(d_961_v82_)) + (self.f4))
                rhs102_ = (d_956_v77_) | (d_956_v77_)
                rhs103_ = d_959_v80_
                rhs104_ = not((d_953_i6_) not in ((d_956_v77_) - (d_956_v77_)))
                rhs105_ = d_974_v93_
                lhs44_ = self
                lhs45_ = d_973_v92_
                lhs46_ = default__.safeIndex(690, (d_973_v92_).length(0))
                lhs44_.f4 = rhs101_
                d_956_v77_ = rhs102_
                lhs45_[lhs46_] = rhs103_
                r0 = rhs104_
                d_974_v93_ = rhs105_
            elif True:
                r0 = not((self).fm5((self).f12, p3, globalState))
                d_975_v94_: _dafny.Array
                nw161_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 3)
                d_975_v94_ = nw161_
                d_976_v95_: _dafny.Map
                d_976_v95_ = _dafny.Map({d_975_v94_: ((self).f12) != ((self).f12)})
                d_976_v95_ = (d_976_v95_).set(d_975_v94_, p2)
                d_977_v96_: _dafny.Array
                nw162_ = _dafny.Array(int(0), 20)
                d_977_v96_ = nw162_
                d_978_v97_: D11
                d_978_v97_ = D11_DC35(p1, p3, p3)
                d_979_v98_: _dafny.Set
                d_979_v98_ = _dafny.Set({p3, p0, p0, False, p0})
                index116_ = default__.safeIndex(359, (d_977_v96_).length(0))
                (d_977_v96_)[index116_] = len((_dafny.Set({p3}) if (d_978_v97_).cf63 else d_979_v98_))
                index117_ = default__.safeIndex(359, (d_977_v96_).length(0))
                (d_977_v96_)[index117_] = p1
                d_980_v99_: str
                d_980_v99_ = _dafny.CodePoint('a')
                d_981_v100_: _dafny.Seq
                d_981_v100_ = _dafny.SeqWithoutIsStrInference([p2, p2, p0])
                d_982_v101_: _dafny.Seq
                d_982_v101_ = _dafny.SeqWithoutIsStrInference([not((d_981_v100_)[default__.safeIndex(self.f4, len(d_981_v100_))])])
                d_983_v102_: _dafny.Set
                d_983_v102_ = _dafny.Set({d_953_i6_})
                d_984_v103_: _dafny.Map
                d_984_v103_ = _dafny.Map({d_983_v102_: p3})
                d_985_v104_: _dafny.Seq
                d_985_v104_ = _dafny.SeqWithoutIsStrInference([(d_977_v96_)[default__.safeIndex(359, (d_977_v96_).length(0))], (d_977_v96_)[default__.safeIndex(359, (d_977_v96_).length(0))]])
                d_986_v105_: _dafny.Array
                nw163_ = _dafny.Array(None, 22)
                nw163_[int(0)] = p3
                nw163_[int(1)] = (d_980_v99_) not in ((self).f12)
                nw163_[int(2)] = p3
                nw163_[int(3)] = not((p2) and (p0))
                nw163_[int(4)] = p3
                nw163_[int(5)] = p2
                nw163_[int(6)] = p3
                nw163_[int(7)] = p3
                nw163_[int(8)] = not(p2)
                nw163_[int(9)] = p2
                nw163_[int(10)] = p3
                nw163_[int(11)] = p0
                nw163_[int(12)] = not (p0) or (p0)
                nw163_[int(13)] = p3
                nw163_[int(14)] = p0
                nw163_[int(15)] = not((d_981_v100_)[default__.safeIndex(len(d_984_v103_), len(d_981_v100_))])
                nw163_[int(16)] = (d_985_v104_) < (_dafny.SeqWithoutIsStrInference([p1, len((self).f12)]))
                nw163_[int(17)] = p2
                nw163_[int(18)] = True
                nw163_[int(19)] = p2
                nw163_[int(20)] = p2
                nw163_[int(21)] = p3
                d_986_v105_ = nw163_
                d_986_v105_ = d_986_v105_
                r0 = (d_981_v100_)[default__.safeIndex((self).f11, len(d_981_v100_))]
            d_987_v106_: _dafny.Seq
            d_987_v106_ = _dafny.SeqWithoutIsStrInference([p3])
            r0 = (d_987_v106_)[default__.safeIndex((0) - ((self).f11), len(d_987_v106_))]
            d_988_v107_: _dafny.Array
            def lambda41_(d_989_p0_):
                def lambda42_(d_990_i9_):
                    return d_989_p0_

                return lambda42_

            init24_ = lambda41_(p0)
            nw164_ = _dafny.Array(None, 1)
            for i0_24_ in range(nw164_.length(0)):
                nw164_[i0_24_] = init24_(i0_24_)
            d_988_v107_ = nw164_
            nw165_ = _dafny.Array(None, 10)
            nw165_[int(0)] = p2
            nw165_[int(1)] = not(p2)
            nw165_[int(2)] = p3
            nw165_[int(3)] = False
            nw165_[int(4)] = p0
            nw165_[int(5)] = p2
            nw165_[int(6)] = not(False)
            nw165_[int(7)] = p3
            nw165_[int(8)] = not(not(True))
            nw165_[int(9)] = p2
            d_988_v107_ = nw165_
        d_991_v108_: str
        d_991_v108_ = _dafny.CodePoint('i')
        d_992_v109_: _dafny.Array
        nw166_ = _dafny.Array(None, 15)
        nw166_[int(0)] = d_991_v108_
        nw166_[int(1)] = d_991_v108_
        nw166_[int(2)] = ((self).f12)[default__.safeIndex(self.f4, len((self).f12))]
        nw166_[int(3)] = default__.fm27(p3, globalState)
        nw166_[int(4)] = d_991_v108_
        nw166_[int(5)] = (_dafny.CodePoint('e') if default__.fm0(self.f4, 444, globalState) else d_991_v108_)
        nw166_[int(6)] = d_991_v108_
        nw166_[int(7)] = d_991_v108_
        nw166_[int(8)] = d_991_v108_
        nw166_[int(9)] = d_991_v108_
        nw166_[int(10)] = d_991_v108_
        nw166_[int(11)] = d_991_v108_
        nw166_[int(12)] = d_991_v108_
        nw166_[int(13)] = d_991_v108_
        nw166_[int(14)] = d_991_v108_
        d_992_v109_ = nw166_
        index118_ = default__.safeIndex(942, (d_992_v109_).length(0))
        (d_992_v109_)[index118_] = d_991_v108_
        index119_ = default__.safeIndex(942, (d_992_v109_).length(0))
        (d_992_v109_)[index119_] = d_991_v108_
        d_993_v110_: _dafny.Set
        d_993_v110_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([(d_992_v109_)[default__.safeIndex(942, (d_992_v109_).length(0))] for d_994_i10_ in range(default__.abs(601))])})
        if (self).fm5((self).f12, (d_993_v110_).issubset(_dafny.Set({(self).f12, (self).f12})), globalState):
            (self).f4 = self.f4
            index120_ = default__.safeIndex(942, (d_992_v109_).length(0))
            (d_992_v109_)[index120_] = (d_992_v109_)[default__.safeIndex(942, (d_992_v109_).length(0))]
            d_992_v109_ = d_992_v109_
            d_995_v111_: _dafny.Seq
            d_995_v111_ = _dafny.SeqWithoutIsStrInference([p2, p3, p0])
            d_996_v112_: C2
            nw167_ = C2()
            nw167_.ctor__((d_995_v111_)[default__.safeIndex(self.f4, len(d_995_v111_))], (self).f11)
            d_996_v112_ = nw167_
            d_997_v113_: _dafny.Set
            d_997_v113_ = _dafny.Set({(d_992_v109_)[default__.safeIndex(942, (d_992_v109_).length(0))], (d_992_v109_)[default__.safeIndex(942, (d_992_v109_).length(0))], (d_992_v109_)[default__.safeIndex(942, (d_992_v109_).length(0))]})
            d_998_v114_: _dafny.Map
            d_998_v114_ = _dafny.Map({(self).f11: d_997_v113_})
            d_998_v114_ = (d_998_v114_).set(p1, d_997_v113_)
        elif True:
            d_999_v115_: _dafny.MultiSet
            d_999_v115_ = _dafny.MultiSet([(self).f11])
            d_1000_v116_: _dafny.Map
            d_1000_v116_ = _dafny.Map({p2: True})
            d_1001_v117_: _dafny.Seq
            d_1001_v117_ = _dafny.SeqWithoutIsStrInference([(self).f11, (0) - ((self).f11)])
            d_1002_v118_: _dafny.MultiSet
            d_1002_v118_ = _dafny.MultiSet([d_991_v108_, (d_992_v109_)[default__.safeIndex(942, (d_992_v109_).length(0))], ((self).f12)[default__.safeIndex((self).f11, len((self).f12))], (d_992_v109_)[default__.safeIndex(942, (d_992_v109_).length(0))]])
            d_1003_v119_: _dafny.Map
            d_1003_v119_ = _dafny.Map({(self).f11: self.f4})
            d_1004_v120_: _dafny.Seq
            d_1004_v120_ = _dafny.SeqWithoutIsStrInference([(d_1003_v119_).set(62, self.f4)])
            d_1005_v121_: _dafny.Array
            nw168_ = _dafny.Array(None, 23)
            nw168_[int(0)] = ((d_999_v115_)[len(d_1000_v116_)] if (len(d_1000_v116_)) in (d_999_v115_) else p1)
            nw168_[int(1)] = p1
            nw168_[int(2)] = default__.fm1(globalState)
            nw168_[int(3)] = (self).f11
            nw168_[int(4)] = self.f4
            nw168_[int(5)] = (d_1001_v117_)[default__.safeIndex((d_1002_v118_).cardinality, len(d_1001_v117_))]
            nw168_[int(6)] = (self.f4) - (p1)
            nw168_[int(7)] = default__.fm1(globalState)
            nw168_[int(8)] = 925
            nw168_[int(9)] = p1
            nw168_[int(10)] = (self).f11
            nw168_[int(11)] = p1
            nw168_[int(12)] = (self).f11
            nw168_[int(13)] = len((self).f12)
            nw168_[int(14)] = (self.f4 if p0 else len(d_1004_v120_))
            nw168_[int(15)] = self.f4
            nw168_[int(16)] = (self).f11
            nw168_[int(17)] = ((self).f11) * (len(d_1003_v119_))
            nw168_[int(18)] = self.f4
            nw168_[int(19)] = self.f4
            nw168_[int(20)] = len(_dafny.SeqWithoutIsStrInference([p0]))
            nw168_[int(21)] = (self).f11
            nw168_[int(22)] = default__.fm1(globalState)
            d_1005_v121_ = nw168_
            index121_ = default__.safeIndex(889, (d_1005_v121_).length(0))
            (d_1005_v121_)[index121_] = default__.fm1(globalState)
            index122_ = default__.safeIndex(889, (d_1005_v121_).length(0))
            rhs106_ = self.f4
            rhs107_ = not(False)
            lhs47_ = d_1005_v121_
            lhs48_ = default__.safeIndex(889, (d_1005_v121_).length(0))
            lhs47_[lhs48_] = rhs106_
            r0 = rhs107_
            index123_ = default__.safeIndex(889, (d_1005_v121_).length(0))
            (d_1005_v121_)[index123_] = p1
            d_1006_v122_: _dafny.Map
            d_1006_v122_ = _dafny.Map({default__.fm1(globalState): not(not(p2))})
            if ((d_1006_v122_)[(self).f11] if ((self).f11) in (d_1006_v122_) else not((False) and (False))):
                r0 = (len(_dafny.Set({p1}))) < (322)
                d_1007_v123_: D12
                d_1007_v123_ = D12_DC37(((self).f12) != ((self).f12), (True) == (p2), p2)
                pat_let_tv25_ = p3
                def iife127_(_pat_let28_0):
                    def iife128_(d_1008_dt__update__tmp_h2_):
                        def iife129_(_pat_let29_0):
                            def iife130_(d_1009_dt__update_hcf67_h0_):
                                return D12_DC37((d_1008_dt__update__tmp_h2_).cf65, (d_1008_dt__update__tmp_h2_).cf66, d_1009_dt__update_hcf67_h0_)
                            return iife130_(_pat_let29_0)
                        return iife129_(pat_let_tv25_)
                    return iife128_(_pat_let28_0)
                d_1007_v123_ = iife127_(d_1007_v123_)
                d_1010_v124_: _dafny.Set
                d_1010_v124_ = _dafny.Set({(self).f11, (self).f11})
                d_1011_v125_: _dafny.Map
                d_1011_v125_ = _dafny.Map({p3: d_1010_v124_})
                d_1012_v126_: D5
                d_1012_v126_ = D5_DC16(p0, False, p0, self.f4, p3)
                rhs108_ = default__.safeModuloInt((self).fm11(d_1011_v125_, p2, globalState), self.f4)
                rhs109_ = (d_1012_v126_).cf31
                lhs49_ = self
                lhs49_.f4 = rhs108_
                r0 = rhs109_
                r0 = p0
                r0 = p0
            elif True:
                d_1013_v127_: _dafny.Array
                def lambda43_(d_1014_v108_, d_1015_p1_, d_1016_p2_):
                    def lambda44_(d_1017_i11_):
                        return (_dafny.Map({_dafny.CodePoint('j'): (self).f11})) | (_dafny.Map({d_1014_v108_: (0) - (len(_dafny.Map({d_1015_p1_: D14_DC41(_dafny.Set({d_1016_p2_}))})))}))

                    return lambda44_

                init25_ = lambda43_(d_991_v108_, p1, p2)
                nw169_ = _dafny.Array(None, 21)
                for i0_25_ in range(nw169_.length(0)):
                    nw169_[i0_25_] = init25_(i0_25_)
                d_1013_v127_ = nw169_
                d_1018_v128_: _dafny.Map
                d_1018_v128_ = _dafny.Map({d_991_v108_: 820})
                index124_ = default__.safeIndex(657, (d_1013_v127_).length(0))
                def iife131_():
                    coll71_ = _dafny.Map()
                    compr_71_: str
                    for compr_71_ in (_dafny.Set({d_991_v108_, (d_992_v109_)[default__.safeIndex(942, (d_992_v109_).length(0))], (d_992_v109_)[default__.safeIndex(942, (d_992_v109_).length(0))], (d_992_v109_)[default__.safeIndex(942, (d_992_v109_).length(0))]})).Elements:
                        d_1019_v129_: str = compr_71_
                        if (d_1019_v129_) in (_dafny.Set({d_991_v108_, (d_992_v109_)[default__.safeIndex(942, (d_992_v109_).length(0))], (d_992_v109_)[default__.safeIndex(942, (d_992_v109_).length(0))], (d_992_v109_)[default__.safeIndex(942, (d_992_v109_).length(0))]})):
                            coll71_[d_1019_v129_] = 298
                    return _dafny.Map(coll71_)
                (d_1013_v127_)[index124_] = (d_1018_v128_) | (iife131_()
                )
                d_1020_v130_: _dafny.Seq
                d_1020_v130_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xdr"))
                d_1021_v131_: _dafny.Map
                d_1021_v131_ = _dafny.Map({p2: default__.safeDivisionInt((self).f11, (d_1005_v121_)[default__.safeIndex(889, (d_1005_v121_).length(0))])})
                index125_ = default__.safeIndex(657, (d_1013_v127_).length(0))
                index126_ = default__.safeIndex(889, (d_1005_v121_).length(0))
                rhs110_ = d_1018_v128_
                rhs111_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "amrqusko"))) + ((self).f12)
                rhs112_ = (0) - (self.f4)
                rhs113_ = (d_1021_v131_) | (_dafny.Map({p3: p1}))
                lhs50_ = d_1013_v127_
                lhs51_ = default__.safeIndex(657, (d_1013_v127_).length(0))
                lhs52_ = d_1005_v121_
                lhs53_ = default__.safeIndex(889, (d_1005_v121_).length(0))
                lhs50_[lhs51_] = rhs110_
                d_1020_v130_ = rhs111_
                lhs52_[lhs53_] = rhs112_
                d_1021_v131_ = rhs113_
                d_1020_v130_ = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_1022_i12_ in range(default__.abs(419))])) + ((self).f12)
                d_1023_v132_: C2
                nw170_ = C2()
                nw170_.ctor__(not(p2), default__.safeModuloInt((self).f11, 374))
                d_1023_v132_ = nw170_
                d_1024_v133_: _dafny.Seq
                d_1024_v133_ = _dafny.SeqWithoutIsStrInference([d_1020_v130_])
                d_1020_v130_ = ((self).f12) + ((d_1024_v133_)[default__.safeIndex((self).f11, len(d_1024_v133_))])
                index127_ = default__.safeIndex(889, (d_1005_v121_).length(0))
                (d_1005_v121_)[index127_] = (self).f11
            d_1025_v134_: _dafny.MultiSet
            d_1025_v134_ = _dafny.MultiSet([p3])
            (self).f4 = ((d_1025_v134_)[p2] if (p2) in (d_1025_v134_) else (self).f11)
            index128_ = default__.safeIndex(942, (d_992_v109_).length(0))
            rhs114_ = (d_1005_v121_)[default__.safeIndex(889, (d_1005_v121_).length(0))]
            rhs115_ = ((_dafny.MultiSet(d_1001_v117_)).intersection(d_999_v115_)).intersection(d_999_v115_)
            rhs116_ = d_991_v108_
            lhs54_ = self
            lhs55_ = d_992_v109_
            lhs56_ = default__.safeIndex(942, (d_992_v109_).length(0))
            lhs54_.f4 = rhs114_
            d_999_v115_ = rhs115_
            lhs55_[lhs56_] = rhs116_
        index129_ = default__.safeIndex(942, (d_992_v109_).length(0))
        rhs117_ = p2
        rhs118_ = (d_992_v109_)[default__.safeIndex(942, (d_992_v109_).length(0))]
        lhs57_ = d_992_v109_
        lhs58_ = default__.safeIndex(942, (d_992_v109_).length(0))
        r0 = rhs117_
        lhs57_[lhs58_] = rhs118_
        d_1026_v135_: _dafny.Set
        d_1026_v135_ = _dafny.Set({(self).f11})
        d_1027_v136_: _dafny.Set
        d_1027_v136_ = _dafny.Set({d_1026_v135_})
        d_1028_v137_: _dafny.Map
        d_1028_v137_ = _dafny.Map({d_1026_v135_: d_991_v108_})
        def iife132_():
            coll72_ = _dafny.Set()
            compr_72_: _dafny.Set
            for compr_72_ in (d_1028_v137_).keys.Elements:
                d_1029_v138_: _dafny.Set = compr_72_
                if (d_1029_v138_) in (d_1028_v137_):
                    coll72_ = coll72_.union(_dafny.Set([d_1029_v138_]))
            return _dafny.Set(coll72_)
        r0 = (d_1027_v136_).issubset(iife132_()
        )
        return r0

    def m4(self, p0, p1, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r2: _dafny.Array = _dafny.Array(None, 0)
        d_1030_v0_: _dafny.Map
        d_1030_v0_ = _dafny.Map({len(default__.fm28(451, globalState)): p1})
        hi9_ = -770
        for d_1031_i0_ in range((self.f4 if ((d_1030_v0_)[p0] if (p0) in (d_1030_v0_) else p1) else p0), hi9_):
            d_1032_v1_: _dafny.Set
            d_1032_v1_ = _dafny.Set({p1})
            d_1033_v2_: D14
            d_1033_v2_ = D14_DC41(d_1032_v1_)
            source9_ = d_1033_v2_
            if source9_.is_DC42:
                d_1034___mcc_h0_ = source9_.cf72
                d_1035___mcc_h1_ = source9_.cf73
                d_1036___mcc_h2_ = source9_.cf74
                d_1037___mcc_h3_ = source9_.cf75
                d_1038_cf75_ = d_1037___mcc_h3_
                d_1039_cf74_ = d_1036___mcc_h2_
                d_1040_cf73_ = d_1035___mcc_h1_
                d_1041_cf72_ = d_1034___mcc_h0_
                (self).f4 = (0) - (p0)
                d_1042_v3_: _dafny.MultiSet
                d_1042_v3_ = _dafny.MultiSet([len((self).f12)])
                d_1043_v4_: _dafny.Seq
                d_1043_v4_ = _dafny.SeqWithoutIsStrInference([d_1038_cf75_, (d_1042_v3_).cardinality])
                d_1044_v5_: _dafny.Seq
                d_1044_v5_ = _dafny.SeqWithoutIsStrInference([p1, d_1040_cf73_])
                d_1045_v6_: _dafny.Seq
                d_1045_v6_ = _dafny.SeqWithoutIsStrInference([d_1044_v5_, d_1044_v5_])
                d_1046_v7_: _dafny.Array
                nw171_ = _dafny.Array(None, 23)
                nw171_[int(0)] = (d_1038_cf75_ if d_1040_cf73_ else 504)
                nw171_[int(1)] = (d_1041_cf72_) * ((self).f11)
                nw171_[int(2)] = self.f4
                nw171_[int(3)] = self.f4
                nw171_[int(4)] = d_1038_cf75_
                nw171_[int(5)] = self.f4
                nw171_[int(6)] = default__.fm1(globalState)
                nw171_[int(7)] = len(_dafny.Set({p1}))
                nw171_[int(8)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ixctxti")))
                nw171_[int(9)] = (self).f11
                nw171_[int(10)] = (0) - ((d_1031_i0_) * (p0))
                nw171_[int(11)] = d_1041_cf72_
                nw171_[int(12)] = (d_1043_v4_)[default__.safeIndex(d_1041_cf72_, len(d_1043_v4_))]
                nw171_[int(13)] = len(default__.fm29(p1, globalState))
                nw171_[int(14)] = len(_dafny.SeqWithoutIsStrInference([d_1038_cf75_]))
                nw171_[int(15)] = 281
                nw171_[int(16)] = 3
                nw171_[int(17)] = self.f4
                nw171_[int(18)] = d_1041_cf72_
                nw171_[int(19)] = d_1041_cf72_
                nw171_[int(20)] = ((self).f11) - ((D7_DC27((self).f11, p0, len(d_1045_v6_))).cf51)
                nw171_[int(21)] = 211
                nw171_[int(22)] = -695
                d_1046_v7_ = nw171_
                d_1047_v8_: _dafny.Seq
                d_1047_v8_ = _dafny.SeqWithoutIsStrInference([d_1030_v0_])
                index130_ = default__.safeIndex(113, (d_1046_v7_).length(0))
                (d_1046_v7_)[index130_] = len(_dafny.Set({(0) - (len((d_1047_v8_)[default__.safeIndex(p0, len(d_1047_v8_))])), d_1041_cf72_, 406}))
                d_1048_v9_: _dafny.Array
                nw172_ = _dafny.Array(False, 13)
                d_1048_v9_ = nw172_
                index131_ = default__.safeIndex(958, (d_1048_v9_).length(0))
                (d_1048_v9_)[index131_] = d_1040_cf73_
                d_1049_v10_: _dafny.Array
                nw173_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 27)
                d_1049_v10_ = nw173_
                index132_ = default__.safeIndex(887, (d_1049_v10_).length(0))
                (d_1049_v10_)[index132_] = (self).f12
                d_1050_v11_: _dafny.MultiSet
                d_1050_v11_ = _dafny.MultiSet([p1, True, d_1040_cf73_])
                index133_ = default__.safeIndex(113, (d_1046_v7_).length(0))
                index134_ = default__.safeIndex(958, (d_1048_v9_).length(0))
                index135_ = default__.safeIndex(887, (d_1049_v10_).length(0))
                rhs119_ = (((0) - ((self).f11) if True else (0) - ((self).f11))) * (((d_1050_v11_)[p1] if (p1) in (d_1050_v11_) else d_1038_cf75_))
                rhs120_ = ((d_1038_cf75_ if p1 else (self).f11)) == (d_1031_i0_)
                rhs121_ = (self).f12
                lhs59_ = d_1046_v7_
                lhs60_ = default__.safeIndex(113, (d_1046_v7_).length(0))
                lhs61_ = d_1048_v9_
                lhs62_ = default__.safeIndex(958, (d_1048_v9_).length(0))
                lhs63_ = d_1049_v10_
                lhs64_ = default__.safeIndex(887, (d_1049_v10_).length(0))
                lhs59_[lhs60_] = rhs119_
                lhs61_[lhs62_] = rhs120_
                lhs63_[lhs64_] = rhs121_
                d_1041_cf72_ = d_1041_cf72_
                d_1051_v12_: D12
                d_1051_v12_ = D12_DC37(p1, p1, d_1040_cf73_)
                d_1052_v13_: _dafny.Seq
                d_1052_v13_ = _dafny.SeqWithoutIsStrInference([d_1051_v12_, default__.fm30(globalState)])
                d_1052_v13_ = d_1052_v13_
            elif True:
                d_1053___mcc_h4_ = source9_.cf71
                d_1054_cf71_ = d_1053___mcc_h4_
                d_1055_v14_: _dafny.Seq
                d_1055_v14_ = _dafny.SeqWithoutIsStrInference([p1])
                d_1056_v15_: _dafny.Set
                d_1056_v15_ = _dafny.Set({p0, default__.fm1(globalState)})
                (self).f4 = default__.safeModuloInt(len(d_1055_v14_), (self).fm11((default__.fm31(False, p1, p0, globalState)).set(p1, d_1056_v15_), p1, globalState))
                d_1057_v16_: str
                d_1057_v16_ = _dafny.CodePoint('x')
                d_1058_v17_: _dafny.Map
                d_1058_v17_ = _dafny.Map({self.f4: d_1057_v16_})
                d_1059_v18_: _dafny.Seq
                d_1059_v18_ = _dafny.SeqWithoutIsStrInference([(self).f12, _dafny.SeqWithoutIsStrInference([d_1057_v16_ for d_1060_i1_ in range(default__.abs(149))]), (self).f12])
                d_1061_v19_: _dafny.Array
                nw174_ = _dafny.Array(None, 23)
                nw174_[int(0)] = d_1057_v16_
                nw174_[int(1)] = d_1057_v16_
                nw174_[int(2)] = d_1057_v16_
                nw174_[int(3)] = d_1057_v16_
                nw174_[int(4)] = d_1057_v16_
                nw174_[int(5)] = d_1057_v16_
                nw174_[int(6)] = (d_1057_v16_ if p1 else d_1057_v16_)
                nw174_[int(7)] = d_1057_v16_
                nw174_[int(8)] = d_1057_v16_
                nw174_[int(9)] = d_1057_v16_
                nw174_[int(10)] = d_1057_v16_
                nw174_[int(11)] = d_1057_v16_
                nw174_[int(12)] = d_1057_v16_
                nw174_[int(13)] = d_1057_v16_
                nw174_[int(14)] = ((d_1058_v17_)[(self).f11] if ((self).f11) in (d_1058_v17_) else _dafny.CodePoint('s'))
                nw174_[int(15)] = d_1057_v16_
                nw174_[int(16)] = d_1057_v16_
                nw174_[int(17)] = d_1057_v16_
                nw174_[int(18)] = default__.fm27(p1, globalState)
                nw174_[int(19)] = d_1057_v16_
                nw174_[int(20)] = ((self).f12)[default__.safeIndex(len(default__.fm16(p1, d_1059_v18_, p1, p0, globalState)), len((self).f12))]
                nw174_[int(21)] = d_1057_v16_
                nw174_[int(22)] = _dafny.CodePoint('x')
                d_1061_v19_ = nw174_
                d_1062_v20_: _dafny.MultiSet
                d_1062_v20_ = _dafny.MultiSet([p0, p0])
                rhs122_ = d_1061_v19_
                rhs123_ = d_1062_v20_
                d_1061_v19_ = rhs122_
                d_1062_v20_ = rhs123_
                d_1063_v21_: D7
                d_1063_v21_ = D7_DC26(p0, len(d_1059_v18_))
                d_1064_v22_: D13
                d_1064_v22_ = D13_DC40(p1)
                d_1065_v23_: _dafny.MultiSet
                d_1065_v23_ = _dafny.MultiSet([_dafny.MultiSet([p1, p1, p1])])
                d_1066_v24_: _dafny.MultiSet
                d_1066_v24_ = _dafny.MultiSet([default__.fm0(len((self).f12), self.f4, globalState), p1])
                d_1067_v25_: _dafny.Map
                d_1067_v25_ = _dafny.Map({(0) - (((d_1065_v23_)[d_1066_v24_] if (d_1066_v24_) in (d_1065_v23_) else len(_dafny.SeqWithoutIsStrInference([p1, p1])))): p0})
                pat_let_tv26_ = d_1062_v20_
                pat_let_tv27_ = globalState
                pat_let_tv28_ = globalState
                pat_let_tv29_ = d_1062_v20_
                d_1068_v26_: _dafny.Array
                nw175_ = _dafny.Array(None, 27)
                def iife133_(_pat_let30_0):
                    def iife134_(d_1069_dt__update__tmp_h0_):
                        def iife135_(_pat_let31_0):
                            def iife136_(d_1070_dt__update_hcf49_h0_):
                                return D7_DC26((d_1069_dt__update__tmp_h0_).cf48, d_1070_dt__update_hcf49_h0_)
                            return iife136_(_pat_let31_0)
                        return iife135_(((pat_let_tv26_)[default__.fm1(pat_let_tv27_)] if (default__.fm1(pat_let_tv28_)) in (pat_let_tv29_) else 106))
                    return iife134_(_pat_let30_0)
                nw175_[int(0)] = iife133_(d_1063_v21_)
                nw175_[int(1)] = (d_1063_v21_ if not(p1) else d_1063_v21_)
                nw175_[int(2)] = d_1063_v21_
                nw175_[int(3)] = D7_DC26(d_1031_i0_, len(d_1030_v0_))
                nw175_[int(4)] = d_1063_v21_
                nw175_[int(5)] = d_1063_v21_
                nw175_[int(6)] = d_1063_v21_
                nw175_[int(7)] = d_1063_v21_
                nw175_[int(8)] = d_1063_v21_
                nw175_[int(9)] = d_1063_v21_
                nw175_[int(10)] = D7_DC26(p0, 560)
                nw175_[int(11)] = d_1063_v21_
                nw175_[int(12)] = d_1063_v21_
                nw175_[int(13)] = d_1063_v21_
                nw175_[int(14)] = D7_DC26(self.f4, default__.fm1(globalState))
                nw175_[int(15)] = d_1063_v21_
                nw175_[int(16)] = d_1063_v21_
                nw175_[int(17)] = default__.fm32(d_1057_v16_, d_1064_v22_, globalState)
                nw175_[int(18)] = d_1063_v21_
                nw175_[int(19)] = d_1063_v21_
                nw175_[int(20)] = d_1063_v21_
                nw175_[int(21)] = d_1063_v21_
                nw175_[int(22)] = D7_DC26(self.f4, ((d_1067_v25_)[len(d_1030_v0_)] if (len(d_1030_v0_)) in (d_1067_v25_) else p0))
                nw175_[int(23)] = default__.fm32(d_1057_v16_, d_1064_v22_, globalState)
                nw175_[int(24)] = d_1063_v21_
                nw175_[int(25)] = d_1063_v21_
                nw175_[int(26)] = d_1063_v21_
                d_1068_v26_ = nw175_
                index136_ = default__.safeIndex(338, (d_1068_v26_).length(0))
                (d_1068_v26_)[index136_] = d_1063_v21_
                index137_ = default__.safeIndex(338, (d_1068_v26_).length(0))
                (d_1068_v26_)[index137_] = default__.fm32(d_1057_v16_, (d_1064_v22_ if p1 else d_1064_v22_), globalState)
                d_1071_v27_: bool
                d_1071_v27_ = True
                d_1072_v28_: D14
                d_1072_v28_ = D14_DC42((d_1066_v24_).cardinality, p1, d_1054_cf71_, d_1031_i0_)
                d_1073_v29_: _dafny.Seq
                d_1073_v29_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({p1}), _dafny.Set({True}), (d_1072_v28_).cf74, _dafny.Set({d_1071_v27_, p1})])
                d_1074_v30_: _dafny.Map
                d_1074_v30_ = _dafny.Map({(self).f11: (self).f12})
                rhs124_ = d_1071_v27_
                rhs125_ = len(d_1055_v14_)
                rhs126_ = (d_1073_v29_) <= (_dafny.SeqWithoutIsStrInference([d_1054_cf71_ for d_1075_i2_ in range(default__.abs(954))]))
                rhs127_ = (d_1074_v30_) == (_dafny.Map({p0: (self).f12}))
                lhs65_ = self
                d_1071_v27_ = rhs124_
                lhs65_.f4 = rhs125_
                d_1071_v27_ = rhs126_
                d_1071_v27_ = rhs127_
            d_1076_v31_: _dafny.MultiSet
            d_1076_v31_ = _dafny.MultiSet([p1, p1, p1])
            d_1077_v32_: _dafny.Seq
            d_1077_v32_ = _dafny.SeqWithoutIsStrInference([False])
            d_1078_v33_: _dafny.Set
            d_1078_v33_ = _dafny.Set({d_1077_v32_})
            (self).f4 = ((d_1076_v31_).cardinality) - (len(d_1078_v33_))
            d_1079_v34_: bool
            d_1079_v34_ = True
            d_1079_v34_ = not(p1)
            d_1080_v35_: _dafny.Array
            nw176_ = _dafny.Array(_dafny.Array(None, 0), 28)
            d_1080_v35_ = nw176_
            d_1080_v35_ = d_1080_v35_
        d_1081_v36_: _dafny.Array
        nw177_ = _dafny.Array(None, 10)
        nw177_[int(0)] = p1
        nw177_[int(1)] = p1
        nw177_[int(2)] = p1
        nw177_[int(3)] = p1
        nw177_[int(4)] = p1
        nw177_[int(5)] = p1
        nw177_[int(6)] = p1
        nw177_[int(7)] = p1
        nw177_[int(8)] = p1
        nw177_[int(9)] = p1
        d_1081_v36_ = nw177_
        d_1082_v37_: D3
        d_1082_v37_ = D3_DC8(_dafny.CodePoint('h'), d_1081_v36_)
        d_1083_v38_: _dafny.Array
        nw178_ = _dafny.Array(None, 7)
        nw178_[int(0)] = (self).f12
        nw178_[int(1)] = ((self).f12) + ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_1084_i4_ in range(default__.abs(-800))])).set(default__.safeIndex(self.f4, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_1085_i4_ in range(default__.abs(-800))]))), (d_1082_v37_).cf13))
        nw178_[int(2)] = (self).f12
        nw178_[int(3)] = (self).f12
        nw178_[int(4)] = (self).f12
        nw178_[int(5)] = (self).f12
        nw178_[int(6)] = (self).f12
        d_1083_v38_ = nw178_
        guard_loop_7_: int
        for guard_loop_7_ in _dafny.IntegerRange(0, (d_1083_v38_).length(0)):
            d_1086_i3_: int = guard_loop_7_
            if (True) and (((0) <= (d_1086_i3_)) and ((d_1086_i3_) < ((d_1083_v38_).length(0)))):
                (d_1083_v38_)[(d_1086_i3_)] = ((self).f12) + (((self).f12).set(default__.safeIndex(p0, len((self).f12)), _dafny.CodePoint('q')))
        d_1087_v39_: D3
        d_1087_v39_ = D3_DC9(p1, self.f4)
        d_1088_v40_: _dafny.Seq
        d_1088_v40_ = _dafny.SeqWithoutIsStrInference([p0])
        d_1089_v41_: _dafny.Array
        nw179_ = _dafny.Array(None, 4)
        nw179_[int(0)] = ((0) - ((d_1087_v39_).cf16) if p1 else self.f4)
        nw179_[int(1)] = -241
        nw179_[int(2)] = (0) - (default__.safeModuloInt(self.f4, p0))
        nw179_[int(3)] = (d_1088_v40_)[default__.safeIndex(self.f4, len(d_1088_v40_))]
        d_1089_v41_ = nw179_
        index138_ = default__.safeIndex(628, (d_1089_v41_).length(0))
        (d_1089_v41_)[index138_] = self.f4
        index139_ = default__.safeIndex(628, (d_1089_v41_).length(0))
        (d_1089_v41_)[index139_] = default__.fm1(globalState)
        d_1090_v42_: _dafny.Array
        d_1090_v42_ = d_1089_v41_
        d_1091_v43_: _dafny.Map
        d_1091_v43_ = _dafny.Map({len((self).f12): (d_1090_v42_)})
        hi10_ = len(d_1091_v43_)
        for d_1092_i5_ in range(default__.safeDivisionInt((self).f11, (d_1089_v41_)[default__.safeIndex(628, (d_1089_v41_).length(0))]), hi10_):
            d_1093_v44_: _dafny.Seq
            d_1093_v44_ = _dafny.SeqWithoutIsStrInference([p1, not(p1)])
            d_1093_v44_ = (d_1093_v44_) + ((_dafny.SeqWithoutIsStrInference([p1, p1, p1, p1])).set(default__.safeIndex(len((self).f12), len(_dafny.SeqWithoutIsStrInference([p1, p1, p1, p1]))), p1))
            d_1094_v45_: bool
            d_1094_v45_ = False
            d_1095_v46_: D6
            d_1095_v46_ = D6_DC21((self).f11, d_1094_v45_)
            d_1094_v45_ = (d_1095_v46_).cf42
            (self).f4 = (d_1089_v41_)[default__.safeIndex(628, (d_1089_v41_).length(0))]
            d_1096_v47_: str
            d_1096_v47_ = _dafny.CodePoint('n')
            d_1097_v48_: _dafny.Seq
            d_1097_v48_ = _dafny.SeqWithoutIsStrInference([((self).f12) + ((self).f12), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_1098_i6_ in range(default__.abs(73))]), (((self).f12).set(default__.safeIndex(self.f4, len((self).f12)), d_1096_v47_)) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_1099_i7_ in range(default__.abs(169))])), (self).f12])
            d_1097_v48_ = d_1097_v48_
        d_1100_v49_: C0
        nw180_ = C0()
        nw180_.ctor__(self.f4)
        d_1100_v49_ = nw180_
        d_1101_i8_: int
        d_1101_i8_ = 0
        with _dafny.label("7"):
            while p1:
                with _dafny.c_label("7"):
                    if (d_1101_i8_) >= (100):
                        raise _dafny.Break("7")
                    d_1101_i8_ = (d_1101_i8_) + (1)
                    d_1102_v50_: bool
                    d_1102_v50_ = False
                    d_1102_v50_ = d_1102_v50_
                    if p1:
                        d_1103_v51_: _dafny.Seq
                        d_1103_v51_ = _dafny.SeqWithoutIsStrInference([p1])
                        d_1104_v52_: _dafny.Map
                        d_1104_v52_ = _dafny.Map({p1: d_1102_v50_})
                        d_1105_v53_: _dafny.Map
                        d_1105_v53_ = _dafny.Map({(d_1103_v51_)[default__.safeIndex((d_1089_v41_)[default__.safeIndex(628, (d_1089_v41_).length(0))], len(d_1103_v51_))]: d_1104_v52_})
                        d_1106_v54_: D7
                        d_1106_v54_ = D7_DC24(((d_1105_v53_)[p1] if (p1) in (d_1105_v53_) else d_1104_v52_))
                        d_1106_v54_ = D7_DC24(d_1104_v52_)
                        index140_ = default__.safeIndex(628, (d_1089_v41_).length(0))
                        (d_1089_v41_)[index140_] = len(_dafny.SeqWithoutIsStrInference([(self).f12 for d_1107_i9_ in range(default__.abs(451))]))
                        d_1108_v55_: _dafny.MultiSet
                        d_1108_v55_ = _dafny.MultiSet([d_1100_v49_.f18])
                        d_1109_v56_: str
                        d_1109_v56_ = _dafny.CodePoint('i')
                        d_1110_v57_: _dafny.Map
                        d_1110_v57_ = _dafny.Map({(d_1108_v55_).cardinality: d_1109_v56_})
                        d_1111_v58_: _dafny.Map
                        d_1111_v58_ = d_1110_v57_
                        (d_1100_v49_).f18 = (len((d_1111_v58_))) - (d_1100_v49_.f18)
                        d_1102_v50_ = (d_1102_v50_) or (d_1102_v50_)
                        d_1112_v59_: C0
                        nw181_ = C0()
                        nw181_.ctor__(d_1100_v49_.f18)
                        d_1112_v59_ = nw181_
                    elif True:
                        d_1113_v60_: _dafny.MultiSet
                        d_1113_v60_ = _dafny.MultiSet([d_1102_v50_, d_1102_v50_, not(default__.fm0(len(_dafny.Set({p1})), -945, globalState))])
                        d_1114_v61_: C0
                        nw182_ = C0()
                        nw182_.ctor__(((d_1113_v60_)[p1] if (p1) in (d_1113_v60_) else p0))
                        d_1114_v61_ = nw182_
                        d_1115_v62_: _dafny.Set
                        d_1115_v62_ = _dafny.Set({default__.fm1(globalState), 961})
                        index141_ = default__.safeIndex(947, (d_1081_v36_).length(0))
                        (d_1081_v36_)[index141_] = (d_1115_v62_) != (d_1115_v62_)
                        index142_ = default__.safeIndex(947, (d_1081_v36_).length(0))
                        (d_1081_v36_)[index142_] = p1
                        d_1116_v63_: _dafny.Map
                        d_1116_v63_ = _dafny.Map({(d_1081_v36_)[default__.safeIndex(947, (d_1081_v36_).length(0))]: d_1115_v62_})
                        d_1117_v64_: D4
                        d_1117_v64_ = D4_DC12((self).f12)
                        d_1118_v65_: _dafny.Set
                        d_1118_v65_ = _dafny.Set({((self).fm11(d_1116_v63_, p1, globalState)) <= ((self).f11), not(((d_1081_v36_)[default__.safeIndex(947, (d_1081_v36_).length(0))] if d_1102_v50_ else True)), (self).fm5((d_1117_v64_).cf23, not((d_1081_v36_)[default__.safeIndex(947, (d_1081_v36_).length(0))]), globalState), p1})
                        d_1118_v65_ = _dafny.Set({(d_1081_v36_)[default__.safeIndex(947, (d_1081_v36_).length(0))], not(p1), (d_1081_v36_)[default__.safeIndex(947, (d_1081_v36_).length(0))], (d_1081_v36_)[default__.safeIndex(947, (d_1081_v36_).length(0))]})
                        d_1119_v66_: C2
                        nw183_ = C2()
                        nw183_.ctor__(p1, (self).f11)
                        d_1119_v66_ = nw183_
                        d_1120_v67_: _dafny.MultiSet
                        d_1120_v67_ = _dafny.MultiSet([d_1100_v49_.f18, (d_1119_v66_).f14, (d_1119_v66_).f14])
                        d_1121_v68_: _dafny.Map
                        d_1121_v68_ = _dafny.Map({(d_1081_v36_)[default__.safeIndex(947, (d_1081_v36_).length(0))]: len((self).f12)})
                        d_1122_v69_: str
                        d_1122_v69_ = _dafny.CodePoint('p')
                        d_1123_v70_: _dafny.Seq
                        d_1123_v70_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1100_v49_.f18]), _dafny.SeqWithoutIsStrInference([(d_1089_v41_)[default__.safeIndex(628, (d_1089_v41_).length(0))] for d_1124_i10_ in range(default__.abs(624))])])
                        (self).f4 = ((0) - (((0) - ((d_1120_v67_).cardinality) if p1 else d_1100_v49_.f18))) + ((len(((self).f12).set(default__.safeIndex(len(d_1121_v68_), len((self).f12)), d_1122_v69_)) if d_1102_v50_ else len(d_1123_v70_)))
                    d_1125_v71_: _dafny.Seq
                    d_1125_v71_ = _dafny.SeqWithoutIsStrInference([False])
                    d_1126_v72_: _dafny.Map
                    d_1126_v72_ = _dafny.Map({p1: len(d_1125_v71_)})
                    d_1127_v73_: _dafny.Map
                    d_1127_v73_ = _dafny.Map({(-847) + ((d_1089_v41_)[default__.safeIndex(628, (d_1089_v41_).length(0))]): len(d_1126_v72_)})
                    d_1127_v73_ = (d_1127_v73_).set(p0, (p0) + (253))
                    (self).f4 = (self).f11
                    pass
            pass
        d_1128_v74_: D0
        d_1128_v74_ = D0_DC0(d_1088_v40_)
        d_1129_v75_: _dafny.Seq
        d_1129_v75_ = _dafny.SeqWithoutIsStrInference([d_1128_v74_])
        r0 = (_dafny.SeqWithoutIsStrInference([d_1128_v74_])) + (d_1129_v75_)
        r1 = (self).f12
        r2 = d_1089_v41_
        return r0, r1, r2

    @property
    def f11(self):
        return self._f11
    @property
    def f12(self):
        return self._f12

class C4(T0, T1):
    def  __init__(self):
        self._f4: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    @property
    def f4(self):
        return self._f4
    @f4.setter
    def f4(self, value):
        self._f4 = value
    def ctor__(self, f4):
        (self).f4 = f4

    def fm5(self, p0, p1, globalState):
        if not (not(False)) or (True):
            def iife137_():
                coll73_ = _dafny.Set()
                compr_73_: _dafny.Seq
                for compr_73_ in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m'), _dafny.CodePoint('e')])])).Elements:
                    d_1130_v0_: _dafny.Seq = compr_73_
                    if (d_1130_v0_) in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m'), _dafny.CodePoint('e')])])):
                        coll73_ = coll73_.union(_dafny.Set([d_1130_v0_]))
                return _dafny.Set(coll73_)
            return (iife137_()
            ).isdisjoint(_dafny.Set({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_1131_i0_ in range(default__.abs(-323))]), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v'), _dafny.CodePoint('o')]), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e'), _dafny.CodePoint('t')])}))
        elif True:
            return True

    def fm10(self, p0, p1, p2, p3, globalState):
        return _dafny.CodePoint('b')

    def m3(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        hi11_ = self.f4
        for d_1132_i0_ in range(default__.safeDivisionInt(p1, p1), hi11_):
            (self).f4 = default__.fm1(globalState)
            d_1133_v0_: _dafny.Map
            d_1133_v0_ = _dafny.Map({d_1132_i0_: p3})
            d_1134_v1_: _dafny.Map
            d_1134_v1_ = _dafny.Map({p2: ((d_1133_v0_)[len(_dafny.SeqWithoutIsStrInference([p2]))] if (len(_dafny.SeqWithoutIsStrInference([p2]))) in (d_1133_v0_) else p0)})
            d_1135_v2_: _dafny.Seq
            d_1135_v2_ = _dafny.SeqWithoutIsStrInference([d_1134_v1_, d_1134_v1_, (d_1134_v1_) | (d_1134_v1_), d_1134_v1_])
            d_1136_v3_: D1
            d_1136_v3_ = D1_DC3(_dafny.Map({p2: d_1132_i0_}))
            d_1137_v4_: _dafny.Set
            d_1137_v4_ = _dafny.Set({d_1136_v3_, d_1136_v3_})
            rhs128_ = (d_1135_v2_) + (_dafny.SeqWithoutIsStrInference([d_1134_v1_, _dafny.Map({p3: not(p3)}), _dafny.Map({p2: p3}), d_1134_v1_, d_1134_v1_]))
            rhs129_ = (d_1137_v4_) | (d_1137_v4_)
            d_1135_v2_ = rhs128_
            d_1137_v4_ = rhs129_
            if not(p0):
                d_1138_v5_: _dafny.Seq
                d_1138_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "poldhr"))
                (self).f4 = len((d_1138_v5_) + (d_1138_v5_))
                d_1139_v6_: _dafny.Array
                def lambda45_(d_1140_v5_):
                    def lambda46_(d_1141_i1_):
                        return default__.safeModuloInt(d_1141_i1_, len((d_1140_v5_).set(default__.safeIndex(self.f4, len(d_1140_v5_)), _dafny.CodePoint('t'))))

                    return lambda46_

                init26_ = lambda45_(d_1138_v5_)
                nw184_ = _dafny.Array(None, 23)
                for i0_26_ in range(nw184_.length(0)):
                    nw184_[i0_26_] = init26_(i0_26_)
                d_1139_v6_ = nw184_
                d_1142_v7_: _dafny.Map
                d_1142_v7_ = _dafny.Map({d_1139_v6_: p2})
                d_1142_v7_ = ((d_1142_v7_) | (d_1142_v7_)) | ((d_1142_v7_) | (d_1142_v7_))
                d_1143_v8_: _dafny.Array
                nw185_ = _dafny.Array(False, 3)
                d_1143_v8_ = nw185_
                d_1144_v9_: _dafny.Map
                d_1144_v9_ = _dafny.Map({d_1132_i0_: d_1143_v8_})
                d_1145_v10_: _dafny.MultiSet
                d_1145_v10_ = _dafny.MultiSet([(0) - (d_1132_i0_), self.f4, len(d_1144_v9_)])
                d_1146_v11_: T0
                nw186_ = C1()
                nw186_.ctor__(((d_1145_v10_)[len(d_1138_v5_)] if (len(d_1138_v5_)) in (d_1145_v10_) else self.f4), (d_1145_v10_).cardinality, p1, p2)
                d_1146_v11_ = nw186_
                d_1147_v12_: str
                d_1147_v12_ = _dafny.CodePoint('k')
                nw187_ = C3()
                nw187_.ctor__(d_1132_i0_, _dafny.SeqWithoutIsStrInference([d_1147_v12_, d_1147_v12_, d_1147_v12_]), d_1132_i0_)
                d_1146_v11_ = nw187_
                d_1148_v13_: _dafny.Map
                d_1148_v13_ = _dafny.Map({d_1138_v5_: d_1146_v11_.f4})
                index143_ = default__.safeIndex(594, (d_1139_v6_).length(0))
                (d_1139_v6_)[index143_] = default__.fm1(globalState)
                index144_ = default__.safeIndex(946, (d_1143_v8_).length(0))
                (d_1143_v8_)[index144_] = p2
                index145_ = default__.safeIndex(594, (d_1139_v6_).length(0))
                index146_ = default__.safeIndex(946, (d_1143_v8_).length(0))
                rhs130_ = d_1148_v13_
                rhs131_ = self.f4
                rhs132_ = p2
                rhs133_ = d_1147_v12_
                rhs134_ = (p2) == (p2)
                lhs66_ = d_1139_v6_
                lhs67_ = default__.safeIndex(594, (d_1139_v6_).length(0))
                lhs68_ = d_1143_v8_
                lhs69_ = default__.safeIndex(946, (d_1143_v8_).length(0))
                d_1148_v13_ = rhs130_
                lhs66_[lhs67_] = rhs131_
                lhs68_[lhs69_] = rhs132_
                d_1147_v12_ = rhs133_
                r0 = rhs134_
                d_1149_v14_: _dafny.Map
                d_1149_v14_ = _dafny.Map({d_1143_v8_: 378})
                d_1150_v15_: _dafny.Map
                d_1150_v15_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_1147_v12_ for d_1151_i2_ in range(default__.abs(-895))]): p2})
                d_1149_v14_ = (d_1149_v14_).set(d_1143_v8_, len((d_1150_v15_) | (d_1150_v15_)))
            elif True:
                (self).f4 = d_1132_i0_
                d_1152_v16_: _dafny.Seq
                d_1152_v16_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wljqfik"))
                d_1153_v17_: T1
                nw188_ = C1()
                nw188_.ctor__(len(d_1152_v16_), self.f4, self.f4, p3)
                d_1153_v17_ = nw188_
                d_1154_v18_: _dafny.Map
                d_1154_v18_ = _dafny.Map({d_1153_v17_: _dafny.Map({True: p3})})
                r0 = (d_1154_v18_) == (d_1154_v18_)
                d_1134_v1_ = ((d_1134_v1_) | (d_1134_v1_)).set(not (False) or (p0), p2)
                (self).f4 = d_1153_v17_.f4
                d_1155_v19_: _dafny.Seq
                d_1155_v19_ = _dafny.SeqWithoutIsStrInference([p3, p3])
                d_1156_v20_: _dafny.Map
                d_1156_v20_ = _dafny.Map({(d_1155_v19_).set(default__.safeIndex(866, len(d_1155_v19_)), p2): d_1152_v16_})
                d_1157_v21_: _dafny.MultiSet
                d_1157_v21_ = _dafny.MultiSet([d_1153_v17_.f4])
                d_1158_v22_: _dafny.Set
                d_1158_v22_ = _dafny.Set({not(not(p2))})
                d_1159_v23_: _dafny.Seq
                d_1159_v23_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([len(d_1158_v22_)])])
                d_1160_v24_: str
                d_1160_v24_ = _dafny.CodePoint('s')
                d_1161_v25_: _dafny.Array
                nw189_ = _dafny.Array(None, 22)
                nw189_[int(0)] = p2
                nw189_[int(1)] = not(p2)
                nw189_[int(2)] = not((d_1153_v17_).fm5(((d_1156_v20_)[d_1155_v19_] if (d_1155_v19_) in (d_1156_v20_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))), not(p3), globalState))
                nw189_[int(3)] = p0
                nw189_[int(4)] = not(p2)
                nw189_[int(5)] = p0
                nw189_[int(6)] = p0
                nw189_[int(7)] = p0
                nw189_[int(8)] = (d_1153_v17_).fm5(d_1152_v16_, p3, globalState)
                nw189_[int(9)] = p0
                nw189_[int(10)] = p3
                nw189_[int(11)] = (773) >= (p1)
                nw189_[int(12)] = not(not(p3))
                nw189_[int(13)] = p3
                nw189_[int(14)] = p0
                nw189_[int(15)] = (d_1153_v17_.f4) < (d_1153_v17_.f4)
                nw189_[int(16)] = False
                nw189_[int(17)] = p0
                nw189_[int(18)] = p0
                nw189_[int(19)] = p3
                nw189_[int(20)] = p0
                nw189_[int(21)] = (d_1157_v21_) != ((d_1159_v23_)[default__.safeIndex((0) - ((_dafny.MultiSet([d_1160_v24_])).cardinality), len(d_1159_v23_))])
                d_1161_v25_ = nw189_
                d_1162_v26_: D13
                d_1162_v26_ = D13_DC40(p0)
                index147_ = default__.safeIndex(589, (d_1161_v25_).length(0))
                (d_1161_v25_)[index147_] = ((d_1134_v1_)[p2] if (p2) in (d_1134_v1_) else (d_1162_v26_).cf70)
                index148_ = default__.safeIndex(589, (d_1161_v25_).length(0))
                (d_1161_v25_)[index148_] = (d_1155_v19_)[default__.safeIndex(len(d_1152_v16_), len(d_1155_v19_))]
            r0 = p3
        d_1163_v27_: _dafny.Set
        d_1163_v27_ = _dafny.Set({p1})
        r0 = (d_1163_v27_).ispropersubset(d_1163_v27_)
        d_1164_v28_: _dafny.MultiSet
        d_1164_v28_ = _dafny.MultiSet([self.f4])
        pat_let_tv30_ = p3
        pat_let_tv31_ = p0
        def lambda47_(source10_):
            d_1165___mcc_h0_ = source10_
            d_1166_cf58_ = d_1165___mcc_h0_
            return (pat_let_tv30_) == (pat_let_tv31_)

        r0 = not(lambda47_(d_1164_v28_))
        d_1167_v29_: _dafny.MultiSet
        d_1167_v29_ = _dafny.MultiSet([p1])
        pat_let_tv32_ = p3
        def lambda48_(source11_):
            d_1168___mcc_h1_ = source11_
            d_1169_cf58_ = d_1168___mcc_h1_
            return pat_let_tv32_

        if lambda48_(d_1167_v29_):
            d_1170_v30_: _dafny.Map
            d_1170_v30_ = _dafny.Map({not(False): 641})
            if (D11_DC35(len(d_1170_v30_), p0, p3)).cf62:
                r0 = p0
                (self).f4 = default__.safeModuloInt(self.f4, (p1 if p0 else p1))
                (self).f4 = p1
                d_1171_v31_: D3
                d_1171_v31_ = D3_DC9(p2, p1)
                d_1172_v32_: D2
                d_1172_v32_ = D2_DC6(self.f4, p0)
                pat_let_tv33_ = p0
                pat_let_tv34_ = d_1172_v32_
                pat_let_tv35_ = globalState
                pat_let_tv36_ = p0
                d_1173_v33_: C2
                nw190_ = C2()
                def iife138_(_pat_let32_0):
                    def iife139_(d_1174_dt__update__tmp_h0_):
                        def iife140_(_pat_let33_0):
                            def iife141_(d_1175_dt__update_hcf16_h0_):
                                return D3_DC9((d_1174_dt__update__tmp_h0_).cf15, d_1175_dt__update_hcf16_h0_)
                            return iife141_(_pat_let33_0)
                        return iife140_(len(_dafny.Map({(self).fm10(self.f4, False, pat_let_tv33_, pat_let_tv34_, pat_let_tv35_): pat_let_tv36_})))
                    return iife139_(_pat_let32_0)
                nw190_.ctor__((iife138_(d_1171_v31_)).cf15, self.f4)
                d_1173_v33_ = nw190_
                d_1176_v34_: _dafny.Seq
                d_1176_v34_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ffsj"))
                d_1177_v35_: str
                d_1177_v35_ = _dafny.CodePoint('e')
                d_1178_v36_: _dafny.Array
                nw191_ = _dafny.Array(False, 2)
                d_1178_v36_ = nw191_
                d_1179_v37_: D3
                d_1179_v37_ = D3_DC8(d_1177_v35_, d_1178_v36_)
                d_1176_v34_ = (((d_1176_v34_).set(default__.safeIndex((d_1173_v33_).f14, len(d_1176_v34_)), d_1177_v35_)) + ((d_1176_v34_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "coyuqojaw"))))).set(default__.safeIndex(len(d_1170_v30_), len(((d_1176_v34_).set(default__.safeIndex((d_1173_v33_).f14, len(d_1176_v34_)), d_1177_v35_)) + ((d_1176_v34_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "coyuqojaw")))))), (d_1179_v37_).cf13)
            elif True:
                d_1180_v38_: _dafny.Array
                nw192_ = _dafny.Array(int(0), 22)
                d_1180_v38_ = nw192_
                d_1181_v41_: _dafny.Array
                def lambda49_(d_1182_i3_):
                    def iife142_():
                        def iife144_():
                            coll76_ = _dafny.Set()
                            compr_76_: str
                            for compr_76_ in (_dafny.Set({_dafny.CodePoint('e')})).Elements:
                                d_1185_v40_: str = compr_76_
                                if (d_1185_v40_) in (_dafny.Set({_dafny.CodePoint('e')})):
                                    coll76_ = coll76_.union(_dafny.Set([d_1185_v40_]))
                            return _dafny.Set(coll76_)
                        coll74_ = _dafny.Map()
                        def iife143_():
                            coll75_ = _dafny.Set()
                            compr_75_: str
                            for compr_75_ in (_dafny.Set({_dafny.CodePoint('e')})).Elements:
                                d_1183_v40_: str = compr_75_
                                if (d_1183_v40_) in (_dafny.Set({_dafny.CodePoint('e')})):
                                    coll75_ = coll75_.union(_dafny.Set([d_1183_v40_]))
                            return _dafny.Set(coll75_)
                        compr_74_: str
                        for compr_74_ in (iife143_()
                        ).Elements:
                            d_1184_v39_: str = compr_74_
                            if (d_1184_v39_) in (iife144_()
                            ):
                                coll74_[d_1184_v39_] = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_1186_i4_ in range(default__.abs(672))]))
                        return _dafny.Map(coll74_)
                    return iife142_()
                    

                init27_ = lambda49_
                nw193_ = _dafny.Array(None, 8)
                for i0_27_ in range(nw193_.length(0)):
                    nw193_[i0_27_] = init27_(i0_27_)
                d_1181_v41_ = nw193_
                d_1187_v42_: _dafny.Seq
                d_1187_v42_ = _dafny.SeqWithoutIsStrInference([p3, p0, p3, p2, p3])
                d_1188_v43_: _dafny.Seq
                d_1188_v43_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vrtblnli"))
                d_1189_v44_: str
                d_1189_v44_ = _dafny.CodePoint('e')
                d_1190_v45_: _dafny.Array
                nw194_ = _dafny.Array(None, 28)
                nw194_[int(0)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yv"))
                nw194_[int(1)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_1191_i5_ in range(default__.abs(76))])
                nw194_[int(2)] = d_1188_v43_
                nw194_[int(3)] = d_1188_v43_
                nw194_[int(4)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qjvnief"))
                nw194_[int(5)] = d_1188_v43_
                nw194_[int(6)] = d_1188_v43_
                nw194_[int(7)] = d_1188_v43_
                nw194_[int(8)] = d_1188_v43_
                nw194_[int(9)] = d_1188_v43_
                nw194_[int(10)] = (d_1188_v43_).set(default__.safeIndex(self.f4, len(d_1188_v43_)), d_1189_v44_)
                nw194_[int(11)] = d_1188_v43_
                nw194_[int(12)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wtmo"))
                nw194_[int(13)] = d_1188_v43_
                nw194_[int(14)] = d_1188_v43_
                nw194_[int(15)] = d_1188_v43_
                nw194_[int(16)] = (D4_DC12(d_1188_v43_)).cf23
                nw194_[int(17)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aiywbb"))
                nw194_[int(18)] = d_1188_v43_
                nw194_[int(19)] = d_1188_v43_
                nw194_[int(20)] = d_1188_v43_
                nw194_[int(21)] = d_1188_v43_
                nw194_[int(22)] = d_1188_v43_
                nw194_[int(23)] = d_1188_v43_
                nw194_[int(24)] = d_1188_v43_
                nw194_[int(25)] = d_1188_v43_
                nw194_[int(26)] = d_1188_v43_
                nw194_[int(27)] = d_1188_v43_
                d_1190_v45_ = nw194_
                d_1192_v46_: D8
                d_1192_v46_ = D8_DC30(p1, d_1190_v45_, p0)
                d_1193_v47_: D8
                d_1193_v47_ = D8_DC31(d_1192_v46_)
                d_1194_v48_: _dafny.Set
                d_1194_v48_ = _dafny.Set({d_1193_v47_, d_1193_v47_, d_1193_v47_})
                d_1195_v49_: _dafny.Array
                nw195_ = _dafny.Array(None, 16)
                nw195_[int(0)] = p3
                nw195_[int(1)] = (d_1187_v42_)[default__.safeIndex(p1, len(d_1187_v42_))]
                nw195_[int(2)] = default__.fm0(p1, p1, globalState)
                nw195_[int(3)] = p0
                nw195_[int(4)] = p3
                nw195_[int(5)] = p0
                nw195_[int(6)] = p3
                nw195_[int(7)] = p2
                nw195_[int(8)] = p2
                nw195_[int(9)] = p3
                nw195_[int(10)] = p0
                nw195_[int(11)] = (_dafny.Set({d_1193_v47_})).issubset(d_1194_v48_)
                nw195_[int(12)] = p3
                nw195_[int(13)] = True
                nw195_[int(14)] = p3
                nw195_[int(15)] = (self.f4) > (p1)
                d_1195_v49_ = nw195_
                index149_ = default__.safeIndex(776, (d_1195_v49_).length(0))
                (d_1195_v49_)[index149_] = p3
                index150_ = default__.safeIndex(776, (d_1195_v49_).length(0))
                rhs135_ = (p2) and (True)
                rhs136_ = d_1180_v38_
                rhs137_ = d_1181_v41_
                rhs138_ = p2
                rhs139_ = (default__.safeModuloInt(self.f4, default__.fm1(globalState))) < ((self.f4) * (p1))
                lhs70_ = d_1195_v49_
                lhs71_ = default__.safeIndex(776, (d_1195_v49_).length(0))
                r0 = rhs135_
                d_1180_v38_ = rhs136_
                d_1181_v41_ = rhs137_
                lhs70_[lhs71_] = rhs138_
                r0 = rhs139_
                d_1196_v50_: _dafny.Map
                d_1196_v50_ = _dafny.Map({default__.fm33(p3, globalState): (self).fm5(d_1188_v43_, p0, globalState)})
                d_1197_v51_: _dafny.Seq
                d_1197_v51_ = _dafny.SeqWithoutIsStrInference([p1])
                d_1196_v50_ = (d_1196_v50_).set(_dafny.MultiSet((d_1197_v51_) + (d_1197_v51_)), p2)
                r0 = default__.fm0(111, len(_dafny.Set({self.f4})), globalState)
                d_1198_v52_: _dafny.Map
                d_1198_v52_ = _dafny.Map({p3: d_1180_v38_})
                d_1198_v52_ = (d_1198_v52_).set((p1) > (p1), d_1180_v38_)
                d_1199_v53_: _dafny.Map
                d_1199_v53_ = _dafny.Map({(d_1195_v49_)[default__.safeIndex(776, (d_1195_v49_).length(0))]: True})
                d_1200_v54_: D7
                d_1200_v54_ = D7_DC24(d_1199_v53_)
                d_1201_v55_: _dafny.MultiSet
                d_1201_v55_ = _dafny.MultiSet([d_1200_v54_])
                d_1202_v56_: _dafny.Map
                d_1202_v56_ = _dafny.Map({(self.f4) > (self.f4): d_1201_v55_})
                d_1202_v56_ = (_dafny.Map({False: d_1201_v55_})) | (d_1202_v56_)
            d_1203_v57_: _dafny.MultiSet
            d_1203_v57_ = _dafny.MultiSet([(p0) or (not(True)), p0, not(False)])
            d_1204_v58_: _dafny.Seq
            d_1204_v58_ = _dafny.SeqWithoutIsStrInference([p0])
            rhs140_ = not(p0)
            rhs141_ = p3
            rhs142_ = self.f4
            rhs143_ = (_dafny.MultiSet([p2])) - ((d_1203_v57_).intersection(_dafny.MultiSet(d_1204_v58_)))
            lhs72_ = self
            r0 = rhs140_
            r0 = rhs141_
            lhs72_.f4 = rhs142_
            d_1203_v57_ = rhs143_
            d_1205_v59_: _dafny.Array
            nw196_ = _dafny.Array(D8.default()(), 14)
            d_1205_v59_ = nw196_
            d_1206_v60_: _dafny.Seq
            d_1206_v60_ = _dafny.SeqWithoutIsStrInference([d_1205_v59_])
            d_1207_v61_: _dafny.Seq
            d_1207_v61_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ymylgyne"))
            d_1208_v62_: _dafny.Seq
            d_1208_v62_ = _dafny.SeqWithoutIsStrInference([d_1207_v61_])
            d_1209_v63_: _dafny.Set
            d_1209_v63_ = _dafny.Set({p0})
            d_1210_v64_: _dafny.Map
            d_1210_v64_ = _dafny.Map({d_1206_v60_: ((d_1208_v62_)[default__.safeIndex(len(d_1209_v63_), len(d_1208_v62_))]) + (d_1207_v61_)})
            d_1210_v64_ = (d_1210_v64_).set(d_1206_v60_, d_1207_v61_)
            r0 = (self).fm5((d_1207_v61_) + (d_1207_v61_), p3, globalState)
            if p2:
                d_1211_v65_: _dafny.Array
                nw197_ = _dafny.Array(int(0), 16)
                d_1211_v65_ = nw197_
                d_1212_v66_: _dafny.Array
                d_1212_v66_ = d_1211_v65_
                d_1212_v66_ = d_1212_v66_
                d_1207_v61_ = d_1207_v61_
                d_1213_v67_: _dafny.Map
                d_1213_v67_ = _dafny.Map({self.f4: False})
                d_1214_v69_: C1
                nw198_ = C1()
                nw198_.ctor__(self.f4, len(d_1213_v67_), p1, p0)
                d_1214_v69_ = nw198_
                d_1215_v70_: _dafny.Map
                def iife145_():
                    coll77_ = _dafny.Map()
                    compr_77_: str
                    for compr_77_ in (d_1207_v61_).Elements:
                        d_1216_v68_: str = compr_77_
                        if (d_1216_v68_) in (d_1207_v61_):
                            coll77_[d_1216_v68_] = d_1213_v67_
                    return _dafny.Map(coll77_)
                d_1215_v70_ = _dafny.Map({(_dafny.Map({_dafny.CodePoint('o'): d_1213_v67_})) | (iife145_()
                ): d_1214_v69_})
                d_1217_v71_: _dafny.Seq
                d_1217_v71_ = _dafny.SeqWithoutIsStrInference([p1])
                d_1218_v72_: D0
                d_1218_v72_ = D0_DC0(d_1217_v71_)
                d_1219_v73_: _dafny.Map
                d_1219_v73_ = _dafny.Map({(d_1218_v72_).cf0: _dafny.Map({d_1214_v69_.f17: p1})})
                d_1215_v70_ = (d_1215_v70_).set(default__.fm34(d_1219_v73_, p1, not(p2), globalState), d_1214_v69_)
                d_1220_v74_: _dafny.Array
                def lambda50_(d_1221_i6_):
                    return _dafny.CodePoint('f')

                init28_ = lambda50_
                nw199_ = _dafny.Array(None, 20)
                for i0_28_ in range(nw199_.length(0)):
                    nw199_[i0_28_] = init28_(i0_28_)
                d_1220_v74_ = nw199_
                d_1222_v75_: str
                d_1222_v75_ = _dafny.CodePoint('v')
                index151_ = default__.safeIndex(353, (d_1220_v74_).length(0))
                (d_1220_v74_)[index151_] = d_1222_v75_
                d_1223_v76_: _dafny.MultiSet
                d_1223_v76_ = _dafny.MultiSet([d_1217_v71_])
                d_1224_v77_: D3
                d_1224_v77_ = D3_DC10(p1, len(d_1217_v71_), p3, (d_1223_v76_).cardinality, p0)
                d_1225_v78_: D2
                d_1225_v78_ = D2_DC6(747, p2)
                index152_ = default__.safeIndex(353, (d_1220_v74_).length(0))
                (d_1220_v74_)[index152_] = (d_1222_v75_ if p0 else (d_1214_v69_).fm10(p1, (d_1224_v77_).cf19, p0, d_1225_v78_, globalState))
                d_1226_v79_: _dafny.Array
                def lambda51_(d_1227_v78_, d_1228_v58_, d_1229_v69_, d_1230_v57_, d_1231_p2_, d_1232_p0_):
                    def lambda52_(d_1233_i7_):
                        return D4_DC13(D1_DC4(len(_dafny.Map({True: d_1227_v78_})), d_1228_v58_, d_1229_v69_.f17), (D5_DC17(d_1230_v57_, d_1231_p2_)).cf35, _dafny.SeqWithoutIsStrInference([d_1232_p0_]))

                    return lambda52_

                init29_ = lambda51_(d_1225_v78_, d_1204_v58_, d_1214_v69_, d_1203_v57_, p2, p0)
                nw200_ = _dafny.Array(None, 16)
                for i0_29_ in range(nw200_.length(0)):
                    nw200_[i0_29_] = init29_(i0_29_)
                d_1226_v79_ = nw200_
                d_1234_v80_: _dafny.Map
                d_1234_v80_ = _dafny.Map({p2: p0})
                d_1235_v81_: D1
                d_1235_v81_ = D1_DC4(len(d_1170_v30_), d_1204_v58_, len(d_1234_v80_))
                d_1236_v82_: D4
                d_1236_v82_ = D4_DC13(d_1235_v81_, not((d_1204_v58_)[default__.safeIndex(self.f4, len(d_1204_v58_))]), default__.fm2(d_1222_v75_, d_1214_v69_.f17, self.f4, globalState))
                index153_ = default__.safeIndex(364, (d_1226_v79_).length(0))
                (d_1226_v79_)[index153_] = d_1236_v82_
                index154_ = default__.safeIndex(364, (d_1226_v79_).length(0))
                (d_1226_v79_)[index154_] = default__.fm35(default__.fm1(globalState), _dafny.SeqWithoutIsStrInference([self.f4, 518, self.f4, p1, self.f4]), globalState)
            elif True:
                (self).f4 = p1
                d_1237_v83_: str
                d_1237_v83_ = _dafny.CodePoint('x')
                d_1238_v84_: _dafny.Array
                nw201_ = _dafny.Array(None, 1)
                nw201_[int(0)] = (d_1237_v83_) not in (d_1207_v61_)
                d_1238_v84_ = nw201_
                d_1238_v84_ = (d_1238_v84_ if p3 else d_1238_v84_)
                r0 = True
                (self).f4 = p1
                (self).f4 = default__.safeModuloInt((p1) - (p1), p1)
        elif True:
            d_1239_v85_: _dafny.Array
            nw202_ = _dafny.Array(False, 4)
            d_1239_v85_ = nw202_
            index155_ = default__.safeIndex(0, (d_1239_v85_).length(0))
            (d_1239_v85_)[index155_] = p0
            index156_ = default__.safeIndex(0, (d_1239_v85_).length(0))
            (d_1239_v85_)[index156_] = p3
            d_1240_v86_: _dafny.Map
            d_1240_v86_ = _dafny.Map({p1: p1})
            d_1241_v87_: _dafny.Map
            d_1241_v87_ = _dafny.Map({(d_1239_v85_)[default__.safeIndex(0, (d_1239_v85_).length(0))]: len(d_1240_v86_)})
            d_1242_v88_: _dafny.MultiSet
            d_1242_v88_ = _dafny.MultiSet([d_1241_v87_, d_1241_v87_])
            d_1243_v89_: C2
            nw203_ = C2()
            nw203_.ctor__(p0, (0) - (((d_1242_v88_).intersection(d_1242_v88_)).cardinality))
            d_1243_v89_ = nw203_
            (self).f4 = default__.safeModuloInt(p1, self.f4)
            (self).f4 = (((d_1243_v89_).f14) * (p1)) * (p1)
            (self).f4 = (d_1243_v89_).f14
        d_1244_v90_: _dafny.Seq
        d_1244_v90_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fwfno"))
        d_1245_v91_: _dafny.MultiSet
        d_1245_v91_ = _dafny.MultiSet([d_1244_v90_])
        if (d_1245_v91_).isdisjoint(d_1245_v91_):
            d_1246_v92_: _dafny.Array
            nw204_ = _dafny.Array(D3.default()(), 28)
            d_1246_v92_ = nw204_
            d_1247_v93_: str
            d_1247_v93_ = _dafny.CodePoint('u')
            d_1248_v94_: _dafny.Array
            nw205_ = _dafny.Array(False, 3)
            d_1248_v94_ = nw205_
            d_1249_v95_: D3
            d_1249_v95_ = D3_DC8(d_1247_v93_, d_1248_v94_)
            index157_ = default__.safeIndex(252, (d_1246_v92_).length(0))
            (d_1246_v92_)[index157_] = D3_DC11(d_1249_v95_)
            d_1250_v96_: D3
            d_1250_v96_ = D3_DC11(d_1249_v95_)
            index158_ = default__.safeIndex(252, (d_1246_v92_).length(0))
            (d_1246_v92_)[index158_] = d_1250_v96_
            d_1251_v97_: _dafny.Seq
            d_1251_v97_ = _dafny.SeqWithoutIsStrInference([False])
            d_1252_v98_: _dafny.Map
            d_1252_v98_ = _dafny.Map({p0: p1})
            d_1253_v99_: _dafny.Array
            nw206_ = _dafny.Array(None, 5)
            nw206_[int(0)] = 803
            nw206_[int(1)] = p1
            nw206_[int(2)] = (_dafny.MultiSet([len(d_1251_v97_), p1])).cardinality
            nw206_[int(3)] = len(d_1252_v98_)
            nw206_[int(4)] = p1
            d_1253_v99_ = nw206_
            index159_ = default__.safeIndex(908, (d_1253_v99_).length(0))
            (d_1253_v99_)[index159_] = self.f4
            index160_ = default__.safeIndex(908, (d_1253_v99_).length(0))
            (d_1253_v99_)[index160_] = default__.safeDivisionInt((self.f4) + (p1), default__.safeDivisionInt(self.f4, p1))
            r0 = (not (not(p3)) or (not(p3))) == (not(p3))
            index161_ = default__.safeIndex(908, (d_1253_v99_).length(0))
            (d_1253_v99_)[index161_] = default__.safeDivisionInt(self.f4, self.f4)
            r0 = p3
        elif True:
            r0 = not((default__.fm0(self.f4, -381, globalState)) or (p3))
            d_1254_v100_: _dafny.Seq
            d_1254_v100_ = _dafny.SeqWithoutIsStrInference([d_1244_v90_, d_1244_v90_, d_1244_v90_])
            if (d_1254_v100_) == (d_1254_v100_):
                d_1255_v101_: str
                d_1255_v101_ = _dafny.CodePoint('j')
                d_1256_v102_: _dafny.Map
                d_1256_v102_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([p0, p0, not(not(p3))])): d_1255_v101_})
                d_1255_v101_ = ((d_1256_v102_)[p1] if (p1) in (d_1256_v102_) else d_1255_v101_)
                d_1257_v103_: _dafny.Array
                def lambda53_(d_1258_p0_, d_1259_p3_, d_1260_p2_, d_1261_p1_):
                    def lambda54_(d_1262_i8_):
                        return ((D14_DC42(len(_dafny.Map({True: _dafny.Map({len(_dafny.Map({d_1258_p0_: d_1259_p3_})): not(False)})})), d_1260_p2_, _dafny.Set({False}), d_1261_p1_)).cf75) < (self.f4)

                    return lambda54_

                init30_ = lambda53_(p0, p3, p2, p1)
                nw207_ = _dafny.Array(None, 10)
                for i0_30_ in range(nw207_.length(0)):
                    nw207_[i0_30_] = init30_(i0_30_)
                d_1257_v103_ = nw207_
                index162_ = default__.safeIndex(923, (d_1257_v103_).length(0))
                (d_1257_v103_)[index162_] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cvdgjyj"))) not in ((d_1254_v100_).set(default__.safeIndex(len(d_1244_v90_), len(d_1254_v100_)), d_1244_v90_))
                d_1263_v104_: _dafny.Map
                d_1263_v104_ = _dafny.Map({p1: p1})
                index163_ = default__.safeIndex(923, (d_1257_v103_).length(0))
                (d_1257_v103_)[index163_] = ((d_1263_v104_) | (d_1263_v104_)) != ((d_1263_v104_) | (_dafny.Map({p1: self.f4})))
                index164_ = default__.safeIndex(923, (d_1257_v103_).length(0))
                (d_1257_v103_)[index164_] = True
                d_1244_v90_ = d_1244_v90_
                index165_ = default__.safeIndex(923, (d_1257_v103_).length(0))
                (d_1257_v103_)[index165_] = (d_1257_v103_)[default__.safeIndex(923, (d_1257_v103_).length(0))]
            elif True:
                d_1264_v105_: _dafny.Array
                def lambda55_(d_1265_i9_):
                    return (d_1265_i9_) * (self.f4)

                init31_ = lambda55_
                nw208_ = _dafny.Array(None, 13)
                for i0_31_ in range(nw208_.length(0)):
                    nw208_[i0_31_] = init31_(i0_31_)
                d_1264_v105_ = nw208_
                d_1266_v106_: _dafny.Map
                d_1266_v106_ = _dafny.Map({d_1264_v105_: self.f4})
                d_1266_v106_ = ((d_1266_v106_) | (d_1266_v106_) if p0 else (_dafny.Map({d_1264_v105_: p1})) | (_dafny.Map({d_1264_v105_: self.f4})))
                (self).f4 = (0) - (self.f4)
                (self).f4 = self.f4
                d_1267_v107_: _dafny.MultiSet
                d_1267_v107_ = _dafny.MultiSet([d_1163_v27_, d_1163_v27_, d_1163_v27_])
                d_1268_v108_: _dafny.Map
                d_1268_v108_ = _dafny.Map({p1: default__.safeModuloInt(((d_1267_v107_)[d_1163_v27_] if (d_1163_v27_) in (d_1267_v107_) else p1), self.f4)})
                d_1268_v108_ = (d_1268_v108_).set(p1, p1)
                (self).f4 = p1
            d_1269_v109_: _dafny.Seq
            d_1269_v109_ = _dafny.SeqWithoutIsStrInference([p1])
            d_1270_v110_: _dafny.Seq
            d_1270_v110_ = _dafny.SeqWithoutIsStrInference([default__.fm0(len(_dafny.Map({p3: p3})), self.f4, globalState)])
            d_1271_v111_: str
            d_1271_v111_ = _dafny.CodePoint('p')
            d_1272_v112_: _dafny.Array
            nw209_ = _dafny.Array(None, 29)
            nw209_[int(0)] = len(d_1269_v109_)
            nw209_[int(1)] = self.f4
            nw209_[int(2)] = len(_dafny.SeqWithoutIsStrInference([True]))
            nw209_[int(3)] = 35
            nw209_[int(4)] = len(d_1270_v110_)
            nw209_[int(5)] = self.f4
            nw209_[int(6)] = (self.f4) * (self.f4)
            nw209_[int(7)] = len(d_1270_v110_)
            nw209_[int(8)] = p1
            nw209_[int(9)] = len(_dafny.SeqWithoutIsStrInference([d_1271_v111_]))
            nw209_[int(10)] = self.f4
            nw209_[int(11)] = (d_1167_v29_).cardinality
            nw209_[int(12)] = p1
            nw209_[int(13)] = p1
            nw209_[int(14)] = self.f4
            nw209_[int(15)] = (918) * ((0) - (self.f4))
            nw209_[int(16)] = self.f4
            nw209_[int(17)] = p1
            nw209_[int(18)] = self.f4
            nw209_[int(19)] = self.f4
            nw209_[int(20)] = p1
            nw209_[int(21)] = self.f4
            nw209_[int(22)] = p1
            nw209_[int(23)] = p1
            nw209_[int(24)] = self.f4
            nw209_[int(25)] = p1
            nw209_[int(26)] = p1
            nw209_[int(27)] = p1
            nw209_[int(28)] = self.f4
            d_1272_v112_ = nw209_
            index166_ = default__.safeIndex(222, (d_1272_v112_).length(0))
            (d_1272_v112_)[index166_] = p1
            index167_ = default__.safeIndex(222, (d_1272_v112_).length(0))
            (d_1272_v112_)[index167_] = p1
            index168_ = default__.safeIndex(222, (d_1272_v112_).length(0))
            (d_1272_v112_)[index168_] = (0) - ((default__.safeDivisionInt(p1, self.f4)) + (p1))
            r0 = not(not((not(p2)) and (not(p2))))
        d_1273_v113_: _dafny.Array
        def lambda56_(d_1274_p3_):
            def lambda57_(d_1275_i10_):
                return d_1274_p3_

            return lambda57_

        init32_ = lambda56_(p3)
        nw210_ = _dafny.Array(None, 11)
        for i0_32_ in range(nw210_.length(0)):
            nw210_[i0_32_] = init32_(i0_32_)
        d_1273_v113_ = nw210_
        d_1276_v114_: _dafny.Set
        d_1276_v114_ = _dafny.Set({p0, p3})
        index169_ = default__.safeIndex(97, (d_1273_v113_).length(0))
        (d_1273_v113_)[index169_] = (d_1276_v114_).issubset(default__.fm17(not(p3), self.f4, globalState))
        d_1277_v115_: _dafny.Seq
        d_1277_v115_ = _dafny.SeqWithoutIsStrInference([p1])
        d_1278_v116_: _dafny.Map
        d_1278_v116_ = _dafny.Map({(d_1277_v115_)[default__.safeIndex(p1, len(d_1277_v115_))]: self.f4})
        d_1279_v117_: C1
        nw211_ = C1()
        nw211_.ctor__(len(d_1278_v116_), self.f4, self.f4, p0)
        d_1279_v117_ = nw211_
        d_1280_v118_: _dafny.MultiSet
        d_1280_v118_ = _dafny.MultiSet([d_1279_v117_, d_1279_v117_])
        d_1281_v119_: D7
        d_1281_v119_ = D7_DC25(p3, default__.fm0(p1, (d_1280_v118_).cardinality, globalState), p1)
        index170_ = default__.safeIndex(97, (d_1273_v113_).length(0))
        (d_1273_v113_)[index170_] = (default__.safeDivisionInt((d_1281_v119_).cf47, d_1279_v117_.f17)) not in (d_1278_v116_)
        r0 = p3
        return r0

    def m4(self, p0, p1, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r2: _dafny.Array = _dafny.Array(None, 0)
        d_1282_v0_: str
        d_1282_v0_ = _dafny.CodePoint('w')
        d_1282_v0_ = d_1282_v0_
        d_1283_v1_: _dafny.Array
        def lambda58_(d_1284_p1_):
            def lambda59_(d_1285_i0_):
                return d_1284_p1_

            return lambda59_

        init33_ = lambda58_(p1)
        nw212_ = _dafny.Array(None, 18)
        for i0_33_ in range(nw212_.length(0)):
            nw212_[i0_33_] = init33_(i0_33_)
        d_1283_v1_ = nw212_
        d_1286_v2_: _dafny.Map
        d_1286_v2_ = _dafny.Map({d_1283_v1_: self.f4})
        d_1286_v2_ = _dafny.Map({d_1283_v1_: default__.safeModuloInt(self.f4, p0)})
        d_1287_v3_: _dafny.Array
        nw213_ = _dafny.Array(int(0), 14)
        d_1287_v3_ = nw213_
        index171_ = default__.safeIndex(960, (d_1287_v3_).length(0))
        (d_1287_v3_)[index171_] = p0
        d_1288_v4_: D8
        d_1288_v4_ = D8_DC29()
        index172_ = default__.safeIndex(960, (d_1287_v3_).length(0))
        rhs144_ = (self.f4) + (self.f4)
        rhs145_ = d_1288_v4_
        rhs146_ = d_1283_v1_
        rhs147_ = (0) - ((self.f4) * (self.f4))
        lhs73_ = d_1287_v3_
        lhs74_ = default__.safeIndex(960, (d_1287_v3_).length(0))
        lhs75_ = self
        lhs73_[lhs74_] = rhs144_
        d_1288_v4_ = rhs145_
        d_1283_v1_ = rhs146_
        lhs75_.f4 = rhs147_
        d_1289_v5_: _dafny.MultiSet
        d_1289_v5_ = _dafny.MultiSet([p1])
        d_1290_v6_: _dafny.Map
        d_1290_v6_ = _dafny.Map({p0: p1})
        d_1291_v7_: D5
        d_1291_v7_ = D5_DC17(d_1289_v5_, ((d_1290_v6_)[len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bsru")))] if (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bsru")))) in (d_1290_v6_) else p1))
        d_1292_v8_: D5
        d_1292_v8_ = D5_DC19(d_1291_v7_)
        d_1293_v9_: D5
        d_1293_v9_ = D5_DC19(d_1292_v8_)
        d_1293_v9_ = d_1293_v9_
        index173_ = default__.safeIndex(960, (d_1287_v3_).length(0))
        (d_1287_v3_)[index173_] = (d_1287_v3_)[default__.safeIndex(960, (d_1287_v3_).length(0))]
        d_1294_v10_: bool
        d_1294_v10_ = True
        d_1295_v11_: _dafny.Seq
        d_1295_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "axuvbeb"))
        d_1294_v10_ = (self).fm5(d_1295_v11_, not (d_1294_v10_) or (p1), globalState)
        d_1296_v12_: _dafny.Seq
        d_1296_v12_ = _dafny.SeqWithoutIsStrInference([(d_1287_v3_)[default__.safeIndex(960, (d_1287_v3_).length(0))], (d_1287_v3_)[default__.safeIndex(960, (d_1287_v3_).length(0))]])
        d_1297_v13_: D0
        d_1297_v13_ = D0_DC0(d_1296_v12_)
        d_1298_v14_: _dafny.Seq
        d_1298_v14_ = _dafny.SeqWithoutIsStrInference([d_1297_v13_])
        r0 = d_1298_v14_
        d_1299_v15_: _dafny.Seq
        d_1299_v15_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1282_v0_ for d_1300_i1_ in range(default__.abs(669))])])
        r1 = (d_1299_v15_)[default__.safeIndex(p0, len(d_1299_v15_))]
        r2 = d_1287_v3_
        return r0, r1, r2

    def m7(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: _dafny.Array = _dafny.Array(None, 0)
        r3: int = int(0)
        d_1301_v0_: _dafny.Map
        d_1301_v0_ = _dafny.Map({p0: _dafny.SeqWithoutIsStrInference([_dafny.Set({p2}) for d_1302_i0_ in range(default__.abs(552))])})
        d_1303_v1_: _dafny.Set
        d_1303_v1_ = _dafny.Set({p2})
        d_1304_v2_: _dafny.Seq
        d_1304_v2_ = _dafny.SeqWithoutIsStrInference([d_1303_v1_, _dafny.Set({p2})])
        d_1305_v3_: _dafny.Seq
        d_1305_v3_ = _dafny.SeqWithoutIsStrInference([len((((d_1301_v0_)[self.f4] if (self.f4) in (d_1301_v0_) else (d_1304_v2_).set(default__.safeIndex(self.f4, len(d_1304_v2_)), d_1303_v1_))).set(default__.safeIndex((0) - (p0), len(((d_1301_v0_)[self.f4] if (self.f4) in (d_1301_v0_) else (d_1304_v2_).set(default__.safeIndex(self.f4, len(d_1304_v2_)), d_1303_v1_)))), d_1303_v1_))])
        d_1306_v4_: C0
        nw214_ = C0()
        nw214_.ctor__(default__.safeDivisionInt(self.f4, len(d_1305_v3_)))
        d_1306_v4_ = nw214_
        d_1307_v5_: _dafny.Array
        nw215_ = _dafny.Array(int(0), 26)
        d_1307_v5_ = nw215_
        d_1308_v6_: _dafny.Map
        d_1308_v6_ = _dafny.Map({d_1307_v5_: p2})
        hi12_ = (len(_dafny.Map({p3: p3}))) - (len(d_1308_v6_))
        for d_1309_i1_ in range(self.f4, hi12_):
            d_1310_v7_: _dafny.Seq
            d_1310_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vemqmita"))
            d_1310_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ukdvqvx"))
            d_1311_v8_: _dafny.Array
            def lambda60_(d_1312_v7_, d_1313_v4_):
                def lambda61_(d_1314_i2_):
                    return _dafny.Map({self.f4: (d_1312_v7_).set(default__.safeIndex(d_1313_v4_.f18, len(d_1312_v7_)), _dafny.CodePoint('m'))})

                return lambda61_

            init34_ = lambda60_(d_1310_v7_, d_1306_v4_)
            nw216_ = _dafny.Array(None, 2)
            for i0_34_ in range(nw216_.length(0)):
                nw216_[i0_34_] = init34_(i0_34_)
            d_1311_v8_ = nw216_
            d_1315_v9_: _dafny.Map
            d_1315_v9_ = _dafny.Map({d_1306_v4_.f18: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mblbnj"))})
            index174_ = default__.safeIndex(932, (d_1311_v8_).length(0))
            (d_1311_v8_)[index174_] = (d_1315_v9_) | (d_1315_v9_)
            index175_ = default__.safeIndex(932, (d_1311_v8_).length(0))
            def iife146_():
                coll78_ = _dafny.Map()
                compr_78_: int
                for compr_78_ in ((_dafny.SeqWithoutIsStrInference([d_1306_v4_.f18 for d_1316_i3_ in range(default__.abs(528))])).set(default__.safeIndex(d_1306_v4_.f18, len(_dafny.SeqWithoutIsStrInference([d_1306_v4_.f18 for d_1317_i3_ in range(default__.abs(528))]))), len(d_1310_v7_))).Elements:
                    d_1318_v10_: int = compr_78_
                    if (d_1318_v10_) in ((_dafny.SeqWithoutIsStrInference([d_1306_v4_.f18 for d_1316_i3_ in range(default__.abs(528))])).set(default__.safeIndex(d_1306_v4_.f18, len(_dafny.SeqWithoutIsStrInference([d_1306_v4_.f18 for d_1317_i3_ in range(default__.abs(528))]))), len(d_1310_v7_))):
                        coll78_[default__.safeModuloInt(d_1318_v10_, 256)] = p1
                return _dafny.Map(coll78_)
            (d_1311_v8_)[index175_] = iife146_()
            
            d_1319_v11_: bool
            d_1319_v11_ = True
            d_1319_v11_ = p3
            d_1320_v12_: str
            d_1320_v12_ = _dafny.CodePoint('l')
            d_1321_v13_: _dafny.Map
            d_1321_v13_ = _dafny.Map({d_1306_v4_.f18: d_1320_v12_})
            d_1322_v14_: _dafny.Set
            d_1322_v14_ = _dafny.Set({p0, (0) - (d_1306_v4_.f18)})
            d_1323_v15_: _dafny.Map
            d_1323_v15_ = _dafny.Map({len((d_1321_v13_) | (d_1321_v13_)): (d_1306_v4_.f18) * (len(d_1322_v14_))})
            d_1323_v15_ = (d_1323_v15_).set(self.f4, 774)
        index176_ = default__.safeIndex(358, (d_1307_v5_).length(0))
        (d_1307_v5_)[index176_] = p0
        index177_ = default__.safeIndex(358, (d_1307_v5_).length(0))
        (d_1307_v5_)[index177_] = d_1306_v4_.f18
        d_1324_v16_: D0
        d_1324_v16_ = D0_DC1(len(default__.fm17(p2, len(p1), globalState)), (d_1307_v5_)[default__.safeIndex(358, (d_1307_v5_).length(0))], p3)
        (self).f4 = (d_1324_v16_).cf2
        (self).f4 = (494) + (p0)
        if p3:
            d_1325_v17_: D14
            d_1325_v17_ = D14_DC41(_dafny.Set({p2}))
            pat_let_tv37_ = p3
            pat_let_tv38_ = p3
            def iife147_(_pat_let34_0):
                def iife148_(d_1326_dt__update__tmp_h0_):
                    def iife149_(_pat_let35_0):
                        def iife150_(d_1327_dt__update_hcf71_h0_):
                            return D14_DC41(d_1327_dt__update_hcf71_h0_)
                        return iife150_(_pat_let35_0)
                    return iife149_(_dafny.Set({not(pat_let_tv37_), pat_let_tv38_}))
                return iife148_(_pat_let34_0)
            source12_ = iife147_(d_1325_v17_)
            if source12_.is_DC42:
                d_1328___mcc_h0_ = source12_.cf72
                d_1329___mcc_h1_ = source12_.cf73
                d_1330___mcc_h2_ = source12_.cf74
                d_1331___mcc_h3_ = source12_.cf75
                d_1332_cf75_ = d_1331___mcc_h3_
                d_1333_cf74_ = d_1330___mcc_h2_
                d_1334_cf73_ = d_1329___mcc_h1_
                d_1335_cf72_ = d_1328___mcc_h0_
                d_1336_v18_: D7
                d_1336_v18_ = D7_DC25(d_1334_cf73_, p2, d_1306_v4_.f18)
                d_1334_cf73_ = default__.fm0((d_1336_v18_).cf47, d_1306_v4_.f18, globalState)
                d_1334_cf73_ = default__.fm0(d_1306_v4_.f18, d_1335_cf72_, globalState)
                d_1337_v19_: _dafny.Array
                nw217_ = _dafny.Array(None, 9)
                d_1337_v19_ = nw217_
                d_1338_v20_: _dafny.Map
                d_1338_v20_ = _dafny.Map({d_1306_v4_.f18: d_1337_v19_})
                d_1339_v21_: _dafny.Map
                d_1339_v21_ = d_1338_v20_
                (d_1306_v4_).f18 = len(((d_1338_v20_).set((0) - (p0), d_1337_v19_)) | ((_dafny.Map({398: d_1337_v19_})) | ((d_1339_v21_))))
                d_1340_v22_: _dafny.MultiSet
                d_1340_v22_ = _dafny.MultiSet([(D12_DC38(p0)).cf68])
                d_1340_v22_ = _dafny.MultiSet([self.f4, d_1306_v4_.f18])
            elif True:
                d_1341___mcc_h4_ = source12_.cf71
                d_1342_cf71_ = d_1341___mcc_h4_
                d_1343_v23_: bool
                d_1343_v23_ = True
                d_1343_v23_ = p3
                d_1344_v24_: _dafny.Array
                def lambda62_(d_1345_v23_):
                    def lambda63_(d_1346_i4_):
                        return d_1345_v23_

                    return lambda63_

                init35_ = lambda62_(d_1343_v23_)
                nw218_ = _dafny.Array(None, 11)
                for i0_35_ in range(nw218_.length(0)):
                    nw218_[i0_35_] = init35_(i0_35_)
                d_1344_v24_ = nw218_
                d_1347_v25_: _dafny.Map
                d_1347_v25_ = _dafny.Map({default__.fm1(globalState): d_1344_v24_})
                d_1348_v26_: D14
                d_1348_v26_ = D14_DC42(default__.safeDivisionInt((d_1307_v5_)[default__.safeIndex(358, (d_1307_v5_).length(0))], self.f4), d_1343_v23_, (d_1342_cf71_) | (d_1342_cf71_), self.f4)
                d_1349_v27_: D5
                d_1349_v27_ = D5_DC18(d_1306_v4_, p3, d_1343_v23_)
                d_1350_v28_: D3
                d_1350_v28_ = D3_DC9(d_1343_v23_, self.f4)
                pat_let_tv39_ = d_1343_v23_
                pat_let_tv40_ = d_1306_v4_
                pat_let_tv41_ = p2
                pat_let_tv42_ = d_1306_v4_
                def iife151_(_pat_let36_0):
                    def iife152_(d_1351_dt__update__tmp_h1_):
                        def iife153_(_pat_let37_0):
                            def iife154_(d_1352_dt__update_hcf38_h0_):
                                def iife155_(_pat_let38_0):
                                    def iife156_(d_1353_dt__update_hcf36_h0_):
                                        return D5_DC18(d_1353_dt__update_hcf36_h0_, (d_1351_dt__update__tmp_h1_).cf37, d_1352_dt__update_hcf38_h0_)
                                    return iife156_(_pat_let38_0)
                                return iife155_((pat_let_tv40_ if pat_let_tv41_ else pat_let_tv42_))
                            return iife154_(_pat_let37_0)
                        return iife153_(not(not(pat_let_tv39_)))
                    return iife152_(_pat_let36_0)
                rhs148_ = (p3) == ((True if d_1343_v23_ else (d_1350_v28_).cf15))
                rhs149_ = d_1347_v25_
                rhs150_ = d_1348_v26_
                rhs151_ = iife151_(d_1349_v27_)
                d_1343_v23_ = rhs148_
                d_1347_v25_ = rhs149_
                d_1348_v26_ = rhs150_
                d_1349_v27_ = rhs151_
                d_1354_v29_: _dafny.Seq
                d_1354_v29_ = _dafny.SeqWithoutIsStrInference([p3, p2, d_1343_v23_])
                d_1355_v30_: D1
                d_1355_v30_ = D1_DC4(-531, d_1354_v29_, d_1306_v4_.f18)
                d_1356_v31_: D4
                d_1356_v31_ = D4_DC13(d_1355_v30_, not(p2), d_1354_v29_)
                d_1357_v32_: D4
                d_1357_v32_ = D4_DC14(d_1356_v31_)
                d_1358_v33_: _dafny.Seq
                d_1358_v33_ = _dafny.SeqWithoutIsStrInference([d_1357_v32_])
                d_1359_v34_: _dafny.Array
                nw219_ = _dafny.Array(None, 2)
                nw219_[int(0)] = default__.fm36(globalState)
                nw219_[int(1)] = (d_1358_v33_) + (d_1358_v33_)
                d_1359_v34_ = nw219_
                d_1359_v34_ = d_1359_v34_
                d_1360_v35_: str
                d_1360_v35_ = _dafny.CodePoint('m')
                d_1354_v29_ = default__.fm2(d_1360_v35_, self.f4, len(p1), globalState)
            index178_ = default__.safeIndex(358, (d_1307_v5_).length(0))
            (d_1307_v5_)[index178_] = d_1306_v4_.f18
            d_1361_v36_: bool
            d_1361_v36_ = True
            d_1361_v36_ = p2
            d_1362_v37_: C2
            nw220_ = C2()
            nw220_.ctor__(p3, self.f4)
            d_1362_v37_ = nw220_
            pat_let_tv43_ = d_1303_v1_
            def iife157_(_pat_let39_0):
                def iife158_(d_1363_dt__update__tmp_h2_):
                    def iife159_(_pat_let40_0):
                        def iife160_(d_1364_dt__update_hcf71_h1_):
                            return D14_DC41(d_1364_dt__update_hcf71_h1_)
                        return iife160_(_pat_let40_0)
                    return iife159_(pat_let_tv43_)
                return iife158_(_pat_let39_0)
            source13_ = iife157_(d_1325_v17_)
            if source13_.is_DC42:
                d_1365___mcc_h5_ = source13_.cf72
                d_1366___mcc_h6_ = source13_.cf73
                d_1367___mcc_h7_ = source13_.cf74
                d_1368___mcc_h8_ = source13_.cf75
                d_1369_cf75_ = d_1368___mcc_h8_
                d_1370_cf74_ = d_1367___mcc_h7_
                d_1371_cf73_ = d_1366___mcc_h6_
                d_1372_cf72_ = d_1365___mcc_h5_
                d_1373_v38_: str
                d_1373_v38_ = _dafny.CodePoint('v')
                d_1374_v39_: C3
                nw221_ = C3()
                nw221_.ctor__((d_1362_v37_).f14, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c'), d_1373_v38_]), p0)
                d_1374_v39_ = nw221_
                d_1375_v40_: _dafny.Seq
                d_1375_v40_ = _dafny.SeqWithoutIsStrInference([True])
                d_1376_v41_: _dafny.Map
                d_1376_v41_ = _dafny.Map({d_1375_v40_: _dafny.CodePoint('s')})
                d_1376_v41_ = d_1376_v41_
                d_1377_v42_: _dafny.Array
                nw222_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 27)
                d_1377_v42_ = nw222_
                index179_ = default__.safeIndex(267, (d_1377_v42_).length(0))
                (d_1377_v42_)[index179_] = (d_1374_v39_).f12
                d_1378_v43_: _dafny.Array
                nw223_ = _dafny.Array(None, 1)
                nw223_[int(0)] = d_1362_v37_.f13
                d_1378_v43_ = nw223_
                d_1379_v44_: _dafny.Set
                d_1379_v44_ = _dafny.Set({d_1306_v4_.f18})
                d_1380_v45_: _dafny.Map
                d_1380_v45_ = _dafny.Map({d_1361_v36_: d_1379_v44_})
                d_1381_v46_: _dafny.MultiSet
                d_1381_v46_ = _dafny.MultiSet([(d_1374_v39_).fm11(d_1380_v45_, d_1362_v37_.f13, globalState), len(p1), self.f4])
                index180_ = default__.safeIndex(793, (d_1378_v43_).length(0))
                (d_1378_v43_)[index180_] = (_dafny.Set({d_1372_cf72_, (d_1381_v46_).cardinality})).issubset(d_1379_v44_)
                index181_ = default__.safeIndex(267, (d_1377_v42_).length(0))
                index182_ = default__.safeIndex(793, (d_1378_v43_).length(0))
                rhs152_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nikmu"))
                rhs153_ = d_1361_v36_
                rhs154_ = (d_1306_v4_.f18) >= (len((d_1374_v39_).f12))
                lhs76_ = d_1377_v42_
                lhs77_ = default__.safeIndex(267, (d_1377_v42_).length(0))
                lhs78_ = d_1378_v43_
                lhs79_ = default__.safeIndex(793, (d_1378_v43_).length(0))
                lhs76_[lhs77_] = rhs152_
                d_1371_cf73_ = rhs153_
                lhs78_[lhs79_] = rhs154_
                r3 = self.f4
            elif True:
                d_1382___mcc_h9_ = source13_.cf71
                d_1383_cf71_ = d_1382___mcc_h9_
                d_1384_v47_: _dafny.Map
                d_1384_v47_ = _dafny.Map({d_1362_v37_.f13: (0) - (d_1306_v4_.f18)})
                d_1385_v48_: _dafny.Map
                d_1385_v48_ = _dafny.Map({d_1307_v5_: ((d_1384_v47_)[True] if (True) in (d_1384_v47_) else p0)})
                d_1385_v48_ = (d_1385_v48_).set(d_1307_v5_, d_1306_v4_.f18)
                d_1386_v49_: C1
                nw224_ = C1()
                nw224_.ctor__(len(p1), default__.safeModuloInt((d_1307_v5_)[default__.safeIndex(358, (d_1307_v5_).length(0))], (d_1307_v5_)[default__.safeIndex(358, (d_1307_v5_).length(0))]), self.f4, p3)
                d_1386_v49_ = nw224_
                d_1387_v50_: C3
                nw225_ = C3()
                nw225_.ctor__(d_1306_v4_.f18, (p1) + (p1), self.f4)
                d_1387_v50_ = nw225_
                (self).f4 = d_1306_v4_.f18
        elif True:
            d_1388_v51_: bool
            d_1388_v51_ = False
            d_1388_v51_ = d_1388_v51_
            (d_1306_v4_).f18 = self.f4
            d_1389_v52_: _dafny.MultiSet
            d_1389_v52_ = _dafny.MultiSet([True, False, False, not(True), d_1388_v51_])
            if not(not(default__.fm0((0) - ((d_1307_v5_)[default__.safeIndex(358, (d_1307_v5_).length(0))]), ((d_1389_v52_)[d_1388_v51_] if (d_1388_v51_) in (d_1389_v52_) else (d_1307_v5_)[default__.safeIndex(358, (d_1307_v5_).length(0))]), globalState))):
                (d_1306_v4_).f18 = default__.fm1(globalState)
                d_1388_v51_ = not((True) and (d_1388_v51_))
                d_1388_v51_ = True
                d_1390_v53_: str
                out35_: str
                out35_ = (self).m9(globalState)
                d_1390_v53_ = out35_
                d_1391_v54_: _dafny.Array
                nw226_ = _dafny.Array(_dafny.Map({}), 3)
                d_1391_v54_ = nw226_
                index183_ = default__.safeIndex(149, (d_1391_v54_).length(0))
                (d_1391_v54_)[index183_] = _dafny.Map({p1: d_1305_v3_})
                d_1392_v55_: _dafny.Map
                d_1392_v55_ = _dafny.Map({p1: d_1305_v3_})
                index184_ = default__.safeIndex(149, (d_1391_v54_).length(0))
                (d_1391_v54_)[index184_] = (d_1392_v55_) | (d_1392_v55_)
            elif True:
                d_1393_v56_: _dafny.Array
                nw227_ = _dafny.Array(None, 21)
                nw227_[int(0)] = d_1388_v51_
                nw227_[int(1)] = d_1388_v51_
                nw227_[int(2)] = p2
                nw227_[int(3)] = d_1388_v51_
                nw227_[int(4)] = p2
                nw227_[int(5)] = p2
                nw227_[int(6)] = p2
                nw227_[int(7)] = p3
                nw227_[int(8)] = d_1388_v51_
                nw227_[int(9)] = d_1388_v51_
                nw227_[int(10)] = p2
                nw227_[int(11)] = False
                nw227_[int(12)] = p2
                nw227_[int(13)] = p2
                nw227_[int(14)] = d_1388_v51_
                nw227_[int(15)] = p3
                nw227_[int(16)] = d_1388_v51_
                nw227_[int(17)] = d_1388_v51_
                nw227_[int(18)] = False
                nw227_[int(19)] = p2
                nw227_[int(20)] = d_1388_v51_
                d_1393_v56_ = nw227_
                d_1394_v57_: _dafny.Seq
                d_1394_v57_ = _dafny.SeqWithoutIsStrInference([d_1393_v56_])
                d_1395_v58_: _dafny.Map
                d_1395_v58_ = _dafny.Map({True: p3})
                d_1396_v59_: _dafny.Map
                d_1396_v59_ = _dafny.Map({d_1307_v5_: (d_1395_v58_).set(p2, d_1388_v51_)})
                d_1397_v60_: _dafny.Map
                d_1397_v60_ = _dafny.Map({(d_1394_v57_) + (d_1394_v57_): d_1396_v59_})
                d_1397_v60_ = (d_1397_v60_).set(_dafny.SeqWithoutIsStrInference([d_1393_v56_, d_1393_v56_]), d_1396_v59_)
                d_1398_v61_: str
                d_1398_v61_ = _dafny.CodePoint('c')
                d_1399_v62_: _dafny.Set
                d_1399_v62_ = _dafny.Set({d_1398_v61_})
                d_1400_v63_: _dafny.MultiSet
                d_1400_v63_ = _dafny.MultiSet([len(d_1399_v62_)])
                d_1401_v64_: _dafny.Map
                d_1401_v64_ = _dafny.Map({p2: self.f4})
                d_1402_v65_: _dafny.Seq
                d_1402_v65_ = _dafny.SeqWithoutIsStrInference([p2])
                d_1403_v66_: D11
                d_1403_v66_ = D11_DC35(len(d_1402_v65_), d_1388_v51_, d_1388_v51_)
                d_1404_v67_: _dafny.Seq
                d_1404_v67_ = _dafny.SeqWithoutIsStrInference([d_1401_v64_])
                d_1405_v68_: _dafny.Array
                nw228_ = _dafny.Array(None, 16)
                nw228_[int(0)] = _dafny.Map({p2: ((d_1400_v63_)[p0] if (p0) in (d_1400_v63_) else d_1306_v4_.f18)})
                nw228_[int(1)] = (_dafny.Map({d_1388_v51_: (d_1400_v63_).cardinality})) | (d_1401_v64_)
                nw228_[int(2)] = d_1401_v64_
                nw228_[int(3)] = (d_1401_v64_) | (_dafny.Map({p3: d_1306_v4_.f18}))
                nw228_[int(4)] = _dafny.Map({d_1388_v51_: d_1306_v4_.f18})
                nw228_[int(5)] = (d_1401_v64_) | (d_1401_v64_)
                nw228_[int(6)] = (d_1401_v64_ if p3 else d_1401_v64_)
                nw228_[int(7)] = ((_dafny.Map({p3: d_1306_v4_.f18})).set((d_1403_v66_).cf63, -63)) | ((d_1404_v67_)[default__.safeIndex(len(p1), len(d_1404_v67_))])
                nw228_[int(8)] = d_1401_v64_
                nw228_[int(9)] = (d_1401_v64_ if p2 else d_1401_v64_)
                nw228_[int(10)] = (d_1401_v64_) | (d_1401_v64_)
                nw228_[int(11)] = d_1401_v64_
                nw228_[int(12)] = d_1401_v64_
                nw228_[int(13)] = d_1401_v64_
                nw228_[int(14)] = (d_1401_v64_) | (d_1401_v64_)
                nw228_[int(15)] = (_dafny.Map({d_1388_v51_: d_1306_v4_.f18})) | (_dafny.Map({p3: ((d_1400_v63_)[len((p1).set(default__.safeIndex(len(p1), len(p1)), d_1398_v61_))] if (len((p1).set(default__.safeIndex(len(p1), len(p1)), d_1398_v61_))) in (d_1400_v63_) else (d_1307_v5_)[default__.safeIndex(358, (d_1307_v5_).length(0))])}))
                d_1405_v68_ = nw228_
                nw229_ = _dafny.Array(_dafny.Map({}), 19)
                d_1405_v68_ = nw229_
                d_1388_v51_ = p2
                d_1406_v69_: _dafny.Set
                d_1406_v69_ = _dafny.Set({(0) - (self.f4), p0, self.f4, ((d_1307_v5_)[default__.safeIndex(358, (d_1307_v5_).length(0))]) - (d_1306_v4_.f18)})
                d_1407_v70_: _dafny.Map
                d_1407_v70_ = _dafny.Map({p3: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cp"))})
                d_1408_v71_: _dafny.Seq
                d_1408_v71_ = _dafny.SeqWithoutIsStrInference([((d_1407_v70_)[False] if (False) in (d_1407_v70_) else _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_1409_i5_ in range(default__.abs(261))]))])
                d_1406_v69_ = default__.fm16((p1) < (p1), d_1408_v71_, p3, p0, globalState)
                d_1410_v72_: _dafny.Map
                d_1410_v72_ = _dafny.Map({(d_1394_v57_)[default__.safeIndex(p0, len(d_1394_v57_))]: len(d_1402_v65_)})
                (d_1306_v4_).f18 = ((0) - ((d_1307_v5_)[default__.safeIndex(358, (d_1307_v5_).length(0))]) if (d_1410_v72_) == (d_1410_v72_) else (512) * (len(d_1406_v69_)))
            d_1411_v73_: str
            d_1411_v73_ = _dafny.CodePoint('k')
            d_1412_v74_: _dafny.Seq
            d_1412_v74_ = _dafny.SeqWithoutIsStrInference([d_1388_v51_, p3, p2, p3, True])
            d_1413_v75_: _dafny.Map
            d_1413_v75_ = _dafny.Map({d_1388_v51_: p3})
            d_1414_v76_: _dafny.Map
            d_1414_v76_ = _dafny.Map({len(d_1413_v75_): (d_1307_v5_)[default__.safeIndex(358, (d_1307_v5_).length(0))]})
            d_1415_v77_: D2
            d_1415_v77_ = D2_DC6(536, p3)
            d_1416_v78_: _dafny.Set
            d_1416_v78_ = _dafny.Set({d_1411_v73_, (self).fm10((d_1307_v5_)[default__.safeIndex(358, (d_1307_v5_).length(0))], p3, (d_1412_v74_)[default__.safeIndex(len(d_1414_v76_), len(d_1412_v74_))], d_1415_v77_, globalState)})
            d_1417_v79_: _dafny.Array
            nw230_ = _dafny.Array(False, 6)
            d_1417_v79_ = nw230_
            index185_ = default__.safeIndex(237, (d_1417_v79_).length(0))
            (d_1417_v79_)[index185_] = p2
            index186_ = default__.safeIndex(358, (d_1307_v5_).length(0))
            index187_ = default__.safeIndex(237, (d_1417_v79_).length(0))
            rhs155_ = d_1306_v4_.f18
            rhs156_ = d_1388_v51_
            rhs157_ = d_1416_v78_
            rhs158_ = not((d_1412_v74_)[default__.safeIndex(len((p1).set(default__.safeIndex(default__.fm1(globalState), len(p1)), d_1411_v73_)), len(d_1412_v74_))])
            lhs80_ = d_1307_v5_
            lhs81_ = default__.safeIndex(358, (d_1307_v5_).length(0))
            lhs82_ = d_1417_v79_
            lhs83_ = default__.safeIndex(237, (d_1417_v79_).length(0))
            lhs80_[lhs81_] = rhs155_
            d_1388_v51_ = rhs156_
            d_1416_v78_ = rhs157_
            lhs82_[lhs83_] = rhs158_
            rhs159_ = d_1303_v1_
            rhs160_ = d_1306_v4_.f18
            lhs84_ = d_1306_v4_
            d_1303_v1_ = rhs159_
            lhs84_.f18 = rhs160_
        r0 = default__.fm1(globalState)
        d_1418_v80_: _dafny.Array
        def lambda64_(d_1419_p2_):
            def lambda65_(d_1420_i6_):
                return d_1419_p2_

            return lambda65_

        init36_ = lambda64_(p2)
        nw231_ = _dafny.Array(None, 18)
        for i0_36_ in range(nw231_.length(0)):
            nw231_[i0_36_] = init36_(i0_36_)
        d_1418_v80_ = nw231_
        d_1421_v81_: _dafny.Seq
        d_1421_v81_ = _dafny.SeqWithoutIsStrInference([d_1418_v80_, d_1418_v80_])
        nw232_ = _dafny.Array(None, 13)
        nw232_[int(0)] = d_1418_v80_
        nw232_[int(1)] = d_1418_v80_
        nw232_[int(2)] = d_1418_v80_
        nw232_[int(3)] = (d_1418_v80_ if p2 else d_1418_v80_)
        nw232_[int(4)] = d_1418_v80_
        nw232_[int(5)] = d_1418_v80_
        nw232_[int(6)] = d_1418_v80_
        nw232_[int(7)] = d_1418_v80_
        nw232_[int(8)] = d_1418_v80_
        nw232_[int(9)] = d_1418_v80_
        nw232_[int(10)] = d_1418_v80_
        nw232_[int(11)] = (d_1421_v81_)[default__.safeIndex((0) - (p0), len(d_1421_v81_))]
        nw232_[int(12)] = (d_1418_v80_ if p3 else d_1418_v80_)
        r1 = nw232_
        r2 = d_1307_v5_
        r3 = p0
        return r0, r1, r2, r3

    def m8(self, p0, globalState):
        d_1422_v0_: bool
        d_1422_v0_ = True
        if d_1422_v0_:
            if d_1422_v0_:
                (self).f4 = default__.fm1(globalState)
                d_1423_v1_: _dafny.MultiSet
                d_1423_v1_ = _dafny.MultiSet([d_1422_v0_, d_1422_v0_, d_1422_v0_])
                d_1424_v2_: _dafny.Seq
                d_1424_v2_ = _dafny.SeqWithoutIsStrInference([d_1422_v0_])
                d_1425_v3_: _dafny.Array
                nw233_ = _dafny.Array(None, 13)
                nw233_[int(0)] = d_1423_v1_
                nw233_[int(1)] = d_1423_v1_
                nw233_[int(2)] = d_1423_v1_
                nw233_[int(3)] = (d_1423_v1_) | (d_1423_v1_)
                nw233_[int(4)] = (d_1423_v1_) - (d_1423_v1_)
                nw233_[int(5)] = d_1423_v1_
                nw233_[int(6)] = d_1423_v1_
                nw233_[int(7)] = (d_1423_v1_).intersection(_dafny.MultiSet(d_1424_v2_))
                nw233_[int(8)] = _dafny.MultiSet([True, d_1422_v0_])
                nw233_[int(9)] = d_1423_v1_
                nw233_[int(10)] = (_dafny.MultiSet([d_1422_v0_])).set((d_1424_v2_)[default__.safeIndex(self.f4, len(d_1424_v2_))], default__.abs(self.f4))
                nw233_[int(11)] = d_1423_v1_
                nw233_[int(12)] = default__.fm25(globalState)
                d_1425_v3_ = nw233_
                index188_ = default__.safeIndex(828, (d_1425_v3_).length(0))
                (d_1425_v3_)[index188_] = (d_1423_v1_) | (_dafny.MultiSet([True]))
                index189_ = default__.safeIndex(828, (d_1425_v3_).length(0))
                (d_1425_v3_)[index189_] = d_1423_v1_
                (self).f4 = ((self.f4 if d_1422_v0_ else self.f4)) + (((0) - (len(_dafny.SeqWithoutIsStrInference([d_1422_v0_, d_1422_v0_])))) + (self.f4))
                d_1426_v4_: C0
                nw234_ = C0()
                nw234_.ctor__(530)
                d_1426_v4_ = nw234_
                index190_ = default__.safeIndex(486, (p0).length(0))
                (p0)[index190_] = (d_1426_v4_.f18) >= (d_1426_v4_.f18)
                index191_ = default__.safeIndex(486, (p0).length(0))
                (p0)[index191_] = (self.f4) <= (self.f4)
            elif True:
                d_1427_v5_: _dafny.Array
                nw235_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 2)
                d_1427_v5_ = nw235_
                d_1428_v6_: _dafny.Array
                nw236_ = _dafny.Array(int(0), 4)
                d_1428_v6_ = nw236_
                d_1429_v7_: _dafny.MultiSet
                d_1429_v7_ = _dafny.MultiSet([d_1422_v0_])
                d_1430_v8_: _dafny.Map
                d_1430_v8_ = _dafny.Map({_dafny.CodePoint('a'): (d_1429_v7_) | (d_1429_v7_)})
                d_1431_v9_: _dafny.Seq
                d_1431_v9_ = _dafny.SeqWithoutIsStrInference([d_1428_v6_, d_1428_v6_])
                d_1432_v10_: D17
                d_1432_v10_ = D17_DC45(d_1430_v8_)
                d_1433_v11_: str
                d_1433_v11_ = _dafny.CodePoint('r')
                rhs161_ = d_1427_v5_
                rhs162_ = (d_1431_v9_)[default__.safeIndex(self.f4, len(d_1431_v9_))]
                rhs163_ = self.f4
                rhs164_ = (d_1432_v10_).cf78
                rhs165_ = ((_dafny.Set({_dafny.CodePoint('w')})).ispropersubset(_dafny.Set({d_1433_v11_, d_1433_v11_}))) == (False)
                lhs85_ = self
                d_1427_v5_ = rhs161_
                d_1428_v6_ = rhs162_
                lhs85_.f4 = rhs163_
                d_1430_v8_ = rhs164_
                d_1422_v0_ = rhs165_
                d_1434_v12_: _dafny.Map
                d_1434_v12_ = _dafny.Map({d_1422_v0_: d_1422_v0_})
                d_1434_v12_ = d_1434_v12_
                d_1435_v13_: _dafny.Seq
                d_1435_v13_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hbh"))
                (self).f4 = len(d_1435_v13_)
                (self).f4 = self.f4
                d_1436_v14_: str
                out36_: str
                out36_ = (self).m9(globalState)
                d_1436_v14_ = out36_
            d_1437_v15_: _dafny.MultiSet
            d_1437_v15_ = _dafny.MultiSet([d_1422_v0_])
            d_1438_v16_: C1
            nw237_ = C1()
            nw237_.ctor__(((d_1437_v15_).cardinality if d_1422_v0_ else self.f4), self.f4, default__.safeDivisionInt(self.f4, self.f4), (-945) > (self.f4))
            d_1438_v16_ = nw237_
            d_1439_v17_: _dafny.Seq
            d_1439_v17_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y"))
            if ((d_1439_v17_) + (d_1439_v17_)) < (d_1439_v17_):
                d_1422_v0_ = d_1422_v0_
                d_1422_v0_ = d_1422_v0_
                d_1440_v18_: str
                d_1440_v18_ = _dafny.CodePoint('v')
                d_1440_v18_ = d_1440_v18_
                d_1441_v19_: _dafny.Seq
                d_1441_v19_ = _dafny.SeqWithoutIsStrInference([default__.fm0(872, self.f4, globalState)])
                d_1442_v20_: _dafny.Array
                nw238_ = _dafny.Array(None, 12)
                nw238_[int(0)] = (d_1438_v16_.f17) * (d_1438_v16_.f17)
                nw238_[int(1)] = (0) - (d_1438_v16_.f17)
                nw238_[int(2)] = d_1438_v16_.f17
                nw238_[int(3)] = default__.fm1(globalState)
                nw238_[int(4)] = (d_1438_v16_.f17) - (d_1438_v16_.f17)
                nw238_[int(5)] = (self.f4) - (len(_dafny.SeqWithoutIsStrInference([p0])))
                nw238_[int(6)] = self.f4
                nw238_[int(7)] = d_1438_v16_.f17
                nw238_[int(8)] = d_1438_v16_.f17
                nw238_[int(9)] = d_1438_v16_.f17
                nw238_[int(10)] = d_1438_v16_.f17
                nw238_[int(11)] = len(default__.fm23((d_1441_v19_)[default__.safeIndex(d_1438_v16_.f17, len(d_1441_v19_))], globalState))
                d_1442_v20_ = nw238_
                index192_ = default__.safeIndex(493, (d_1442_v20_).length(0))
                (d_1442_v20_)[index192_] = self.f4
                index193_ = default__.safeIndex(493, (d_1442_v20_).length(0))
                (d_1442_v20_)[index193_] = default__.safeModuloInt(default__.safeDivisionInt(d_1438_v16_.f17, d_1438_v16_.f17), default__.safeModuloInt(d_1438_v16_.f17, -74))
                index194_ = default__.safeIndex(25, (p0).length(0))
                (p0)[index194_] = d_1422_v0_
                d_1443_v21_: _dafny.Seq
                d_1443_v21_ = _dafny.SeqWithoutIsStrInference([(d_1437_v15_).cardinality])
                index195_ = default__.safeIndex(25, (p0).length(0))
                (p0)[index195_] = default__.fm0((d_1443_v21_)[default__.safeIndex(424, len(d_1443_v21_))], d_1438_v16_.f17, globalState)
            elif True:
                d_1422_v0_ = True
                d_1444_v22_: _dafny.Set
                d_1444_v22_ = _dafny.Set({d_1422_v0_, d_1422_v0_})
                d_1445_v23_: _dafny.Map
                d_1445_v23_ = _dafny.Map({d_1444_v22_: not((self.f4) >= (self.f4))})
                d_1445_v23_ = (d_1445_v23_).set(_dafny.Set({d_1422_v0_, d_1422_v0_, d_1422_v0_}), d_1422_v0_)
                d_1446_v24_: _dafny.Seq
                d_1446_v24_ = _dafny.SeqWithoutIsStrInference([False, d_1422_v0_, d_1422_v0_])
                d_1447_v25_: _dafny.Map
                d_1447_v25_ = _dafny.Map({d_1438_v16_.f17: d_1446_v24_})
                d_1448_v27_: C0
                nw239_ = C0()
                def iife161_():
                    coll79_ = _dafny.Set()
                    compr_79_: int
                    for compr_79_ in (d_1447_v25_).keys.Elements:
                        d_1449_v26_: int = compr_79_
                        if (d_1449_v26_) in (d_1447_v25_):
                            coll79_ = coll79_.union(_dafny.Set([(d_1449_v26_) * (414)]))
                    return _dafny.Set(coll79_)
                nw239_.ctor__(len(iife161_()
                ))
                d_1448_v27_ = nw239_
                d_1437_v15_ = d_1437_v15_
                d_1450_v28_: _dafny.Seq
                d_1450_v28_ = _dafny.SeqWithoutIsStrInference([(0) - (self.f4)])
                d_1422_v0_ = (d_1438_v16_.f17) not in (d_1450_v28_)
            d_1422_v0_ = not(d_1422_v0_)
            (self).f4 = default__.safeModuloInt(310, (d_1438_v16_.f17) + (self.f4))
        elif True:
            d_1422_v0_ = d_1422_v0_
            d_1451_v29_: _dafny.Map
            d_1451_v29_ = _dafny.Map({self.f4: d_1422_v0_})
            d_1452_v30_: C2
            nw240_ = C2()
            nw240_.ctor__(d_1422_v0_, len(d_1451_v29_))
            d_1452_v30_ = nw240_
            d_1453_v31_: _dafny.Seq
            d_1453_v31_ = _dafny.SeqWithoutIsStrInference([d_1422_v0_, d_1422_v0_, d_1422_v0_])
            d_1454_v32_: _dafny.Seq
            d_1454_v32_ = _dafny.SeqWithoutIsStrInference([(d_1453_v31_)[default__.safeIndex((d_1452_v30_).f14, len(d_1453_v31_))]])
            d_1455_v33_: _dafny.Map
            d_1455_v33_ = _dafny.Map({self.f4: d_1454_v32_})
            source14_ = D12_DC36(d_1455_v33_)
            if source14_.is_DC37:
                d_1456___mcc_h0_ = source14_.cf65
                d_1457___mcc_h1_ = source14_.cf66
                d_1458___mcc_h2_ = source14_.cf67
                d_1459_cf67_ = d_1458___mcc_h2_
                d_1460_cf66_ = d_1457___mcc_h1_
                d_1461_cf65_ = d_1456___mcc_h0_
                d_1462_v34_: _dafny.Array
                def lambda66_(d_1463_i0_):
                    return (_dafny.CodePoint('b') if True else _dafny.CodePoint('e'))

                init37_ = lambda66_
                nw241_ = _dafny.Array(None, 1)
                for i0_37_ in range(nw241_.length(0)):
                    nw241_[i0_37_] = init37_(i0_37_)
                d_1462_v34_ = nw241_
                d_1464_v35_: str
                d_1464_v35_ = _dafny.CodePoint('f')
                index196_ = default__.safeIndex(191, (d_1462_v34_).length(0))
                (d_1462_v34_)[index196_] = (_dafny.CodePoint('y') if d_1460_cf66_ else d_1464_v35_)
                index197_ = default__.safeIndex(191, (d_1462_v34_).length(0))
                (d_1462_v34_)[index197_] = _dafny.CodePoint('o')
                (self).f4 = (d_1452_v30_).f14
                d_1465_v36_: _dafny.Seq
                d_1465_v36_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ejjdarfg"))
                d_1466_v37_: _dafny.MultiSet
                d_1466_v37_ = _dafny.MultiSet([True, (self).fm5(d_1465_v36_, d_1452_v30_.f13, globalState)])
                d_1466_v37_ = ((d_1466_v37_) | (d_1466_v37_)) - (_dafny.MultiSet([d_1452_v30_.f13, d_1422_v0_, d_1459_cf67_]))
                d_1465_v36_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ajkahjx"))) + (d_1465_v36_)) + (default__.fm23(d_1460_cf66_, globalState))
            elif source14_.is_DC38:
                d_1467___mcc_h3_ = source14_.cf68
                d_1468_cf68_ = d_1467___mcc_h3_
                index198_ = default__.safeIndex(128, (p0).length(0))
                (p0)[index198_] = True
                d_1469_v38_: _dafny.Seq
                d_1469_v38_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vejy"))
                d_1470_v39_: _dafny.Set
                d_1470_v39_ = _dafny.Set({d_1454_v32_})
                index199_ = default__.safeIndex(128, (p0).length(0))
                rhs166_ = default__.safeModuloInt((d_1452_v30_).f14, (d_1452_v30_).f14)
                rhs167_ = ((d_1470_v39_).intersection(d_1470_v39_)).issubset(d_1470_v39_)
                rhs168_ = d_1468_cf68_
                rhs169_ = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_1471_i1_ in range(default__.abs(904))])) + (default__.fm23(d_1422_v0_, globalState))
                lhs86_ = p0
                lhs87_ = default__.safeIndex(128, (p0).length(0))
                d_1468_cf68_ = rhs166_
                lhs86_[lhs87_] = rhs167_
                d_1468_cf68_ = rhs168_
                d_1469_v38_ = rhs169_
                d_1472_v40_: C2
                nw242_ = C2()
                nw242_.ctor__((p0)[default__.safeIndex(128, (p0).length(0))], (d_1452_v30_).f14)
                d_1472_v40_ = nw242_
                d_1473_v41_: str
                d_1473_v41_ = _dafny.CodePoint('j')
                d_1474_v42_: _dafny.Map
                d_1474_v42_ = _dafny.Map({d_1452_v30_.f13: (d_1469_v38_).set(default__.safeIndex(self.f4, len(d_1469_v38_)), d_1473_v41_)})
                d_1475_v43_: _dafny.Map
                d_1475_v43_ = _dafny.Map({self.f4: d_1469_v38_})
                d_1476_v44_: _dafny.Map
                d_1476_v44_ = _dafny.Map({True: self.f4})
                (d_1452_v30_).f13 = (self).fm5((((d_1474_v42_)[(p0)[default__.safeIndex(128, (p0).length(0))]] if ((p0)[default__.safeIndex(128, (p0).length(0))]) in (d_1474_v42_) else ((d_1475_v43_)[len(d_1469_v38_)] if (len(d_1469_v38_)) in (d_1475_v43_) else d_1469_v38_)) if d_1452_v30_.f13 else d_1469_v38_), ((d_1452_v30_).f14) != (((d_1476_v44_)[False] if (False) in (d_1476_v44_) else (d_1452_v30_).f14)), globalState)
                d_1469_v38_ = ((d_1469_v38_) + (d_1469_v38_)) + (_dafny.SeqWithoutIsStrInference([d_1473_v41_, d_1473_v41_]))
            elif True:
                d_1477___mcc_h4_ = source14_.cf64
                d_1478_cf64_ = d_1477___mcc_h4_
                d_1479_v45_: _dafny.Array
                nw243_ = _dafny.Array(int(0), 21)
                d_1479_v45_ = nw243_
                index200_ = default__.safeIndex(731, (d_1479_v45_).length(0))
                (d_1479_v45_)[index200_] = default__.safeModuloInt(self.f4, (d_1452_v30_).f14)
                index201_ = default__.safeIndex(731, (d_1479_v45_).length(0))
                (d_1479_v45_)[index201_] = default__.safeModuloInt(self.f4, (d_1452_v30_).f14)
                d_1480_v46_: _dafny.Seq
                d_1480_v46_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "namqbvvvj"))
                d_1480_v46_ = d_1480_v46_
                d_1481_v47_: D7
                d_1481_v47_ = D7_DC25(d_1422_v0_, d_1422_v0_, (d_1479_v45_)[default__.safeIndex(731, (d_1479_v45_).length(0))])
                d_1482_v48_: _dafny.Set
                d_1482_v48_ = _dafny.Set({(d_1481_v47_).cf45})
                d_1483_v49_: D4
                d_1483_v49_ = D4_DC12(d_1480_v46_)
                d_1484_v50_: _dafny.Map
                d_1484_v50_ = _dafny.Map({(d_1482_v48_).ispropersubset(_dafny.Set({d_1452_v30_.f13})): (d_1483_v49_).cf23})
                d_1485_v51_: _dafny.Set
                d_1485_v51_ = _dafny.Set({(d_1452_v30_).f14})
                d_1486_v52_: _dafny.Map
                d_1486_v52_ = _dafny.Map({self.f4: d_1485_v51_})
                d_1480_v46_ = ((d_1484_v50_)[(default__.fm16(d_1452_v30_.f13, _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "osodv"))]), d_1422_v0_, 417, globalState)).isdisjoint(((d_1486_v52_)[self.f4] if (self.f4) in (d_1486_v52_) else d_1485_v51_))] if ((default__.fm16(d_1452_v30_.f13, _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "osodv"))]), d_1422_v0_, 417, globalState)).isdisjoint(((d_1486_v52_)[self.f4] if (self.f4) in (d_1486_v52_) else d_1485_v51_))) in (d_1484_v50_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oubareni")))
                d_1487_v53_: _dafny.Map
                d_1487_v53_ = _dafny.Map({d_1422_v0_: len(_dafny.SeqWithoutIsStrInference([d_1452_v30_.f13]))})
                d_1488_v54_: D7
                d_1488_v54_ = D7_DC26(len(d_1487_v53_), (d_1479_v45_)[default__.safeIndex(731, (d_1479_v45_).length(0))])
                d_1489_v55_: D13
                d_1489_v55_ = D13_DC40(d_1452_v30_.f13)
                d_1490_v56_: _dafny.Array
                nw244_ = _dafny.Array(None, 14)
                nw244_[int(0)] = d_1488_v54_
                nw244_[int(1)] = default__.fm32((d_1480_v46_)[default__.safeIndex((d_1452_v30_).f14, len(d_1480_v46_))], d_1489_v55_, globalState)
                nw244_[int(2)] = d_1488_v54_
                nw244_[int(3)] = d_1488_v54_
                nw244_[int(4)] = d_1488_v54_
                nw244_[int(5)] = d_1488_v54_
                nw244_[int(6)] = d_1488_v54_
                nw244_[int(7)] = D7_DC26((d_1479_v45_)[default__.safeIndex(731, (d_1479_v45_).length(0))], 747)
                nw244_[int(8)] = d_1488_v54_
                nw244_[int(9)] = d_1488_v54_
                nw244_[int(10)] = d_1488_v54_
                nw244_[int(11)] = D7_DC26(98, (d_1452_v30_).f14)
                nw244_[int(12)] = d_1488_v54_
                nw244_[int(13)] = d_1488_v54_
                d_1490_v56_ = nw244_
                index202_ = default__.safeIndex(83, (d_1490_v56_).length(0))
                (d_1490_v56_)[index202_] = d_1488_v54_
                index203_ = default__.safeIndex(83, (d_1490_v56_).length(0))
                (d_1490_v56_)[index203_] = d_1488_v54_
            (self).f4 = default__.safeDivisionInt(default__.safeDivisionInt((d_1452_v30_).f14, (d_1452_v30_).f14), (d_1452_v30_).f14)
            d_1491_v57_: _dafny.Array
            def lambda67_(d_1492_v0_, d_1493_v30_):
                def lambda68_(d_1494_i2_):
                    return (d_1494_i2_) + (len(_dafny.Map({len(_dafny.Map({self.f4: not(d_1492_v0_)})): D0_DC2(D0_DC1((d_1493_v30_).f14, (d_1493_v30_).f14, d_1493_v30_.f13))})))

                return lambda68_

            init38_ = lambda67_(d_1422_v0_, d_1452_v30_)
            nw245_ = _dafny.Array(None, 29)
            for i0_38_ in range(nw245_.length(0)):
                nw245_[i0_38_] = init38_(i0_38_)
            d_1491_v57_ = nw245_
            d_1495_v58_: _dafny.Map
            d_1495_v58_ = _dafny.Map({d_1491_v57_: d_1452_v30_.f13})
            d_1495_v58_ = (d_1495_v58_).set(d_1491_v57_, d_1422_v0_)
        if d_1422_v0_:
            d_1496_v59_: _dafny.Seq
            d_1496_v59_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pw"))
            if (self.f4) >= (len(d_1496_v59_)):
                d_1497_v60_: str
                d_1497_v60_ = _dafny.CodePoint('r')
                d_1497_v60_ = d_1497_v60_
                d_1422_v0_ = d_1422_v0_
                d_1498_v61_: _dafny.Seq
                d_1498_v61_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "alysl")), (d_1496_v59_).set(default__.safeIndex((0) - (self.f4), len(d_1496_v59_)), _dafny.CodePoint('g'))])
                d_1499_v62_: _dafny.Map
                d_1499_v62_ = _dafny.Map({self.f4: d_1496_v59_})
                d_1500_v63_: _dafny.Array
                nw246_ = _dafny.Array(None, 23)
                nw246_[int(0)] = d_1496_v59_
                nw246_[int(1)] = (d_1496_v59_ if d_1422_v0_ else d_1496_v59_)
                nw246_[int(2)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yjgfl")))
                nw246_[int(3)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rrhfbmkka"))) + (d_1496_v59_)
                nw246_[int(4)] = d_1496_v59_
                nw246_[int(5)] = d_1496_v59_
                nw246_[int(6)] = d_1496_v59_
                nw246_[int(7)] = (d_1498_v61_)[default__.safeIndex(self.f4, len(d_1498_v61_))]
                nw246_[int(8)] = (d_1496_v59_ if d_1422_v0_ else _dafny.SeqWithoutIsStrInference([d_1497_v60_ for d_1501_i3_ in range(default__.abs(977))]))
                nw246_[int(9)] = d_1496_v59_
                nw246_[int(10)] = d_1496_v59_
                nw246_[int(11)] = d_1496_v59_
                nw246_[int(12)] = default__.fm23(d_1422_v0_, globalState)
                nw246_[int(13)] = ((d_1499_v62_)[self.f4] if (self.f4) in (d_1499_v62_) else d_1496_v59_)
                nw246_[int(14)] = d_1496_v59_
                nw246_[int(15)] = (d_1496_v59_ if False else d_1496_v59_)
                nw246_[int(16)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y"))
                nw246_[int(17)] = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ylipchha"))) + (d_1496_v59_)).set(default__.safeIndex(default__.fm1(globalState), len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ylipchha"))) + (d_1496_v59_))), d_1497_v60_)
                nw246_[int(18)] = d_1496_v59_
                nw246_[int(19)] = d_1496_v59_
                nw246_[int(20)] = d_1496_v59_
                nw246_[int(21)] = (_dafny.SeqWithoutIsStrInference([d_1497_v60_ for d_1502_i4_ in range(default__.abs(744))])) + (d_1496_v59_)
                nw246_[int(22)] = d_1496_v59_
                d_1500_v63_ = nw246_
                index204_ = default__.safeIndex(878, (d_1500_v63_).length(0))
                (d_1500_v63_)[index204_] = d_1496_v59_
                d_1503_v64_: _dafny.Seq
                d_1503_v64_ = _dafny.SeqWithoutIsStrInference([d_1422_v0_, d_1422_v0_])
                d_1504_v65_: _dafny.Map
                d_1504_v65_ = _dafny.Map({d_1422_v0_: d_1422_v0_})
                d_1505_v66_: _dafny.Array
                nw247_ = _dafny.Array(None, 20)
                nw247_[int(0)] = _dafny.SeqWithoutIsStrInference([d_1422_v0_, d_1422_v0_, d_1422_v0_])
                nw247_[int(1)] = d_1503_v64_
                nw247_[int(2)] = d_1503_v64_
                nw247_[int(3)] = _dafny.SeqWithoutIsStrInference([d_1422_v0_, False, d_1422_v0_])
                nw247_[int(4)] = _dafny.SeqWithoutIsStrInference([d_1422_v0_, d_1422_v0_])
                nw247_[int(5)] = d_1503_v64_
                nw247_[int(6)] = d_1503_v64_
                nw247_[int(7)] = d_1503_v64_
                nw247_[int(8)] = d_1503_v64_
                nw247_[int(9)] = d_1503_v64_
                nw247_[int(10)] = _dafny.SeqWithoutIsStrInference([d_1422_v0_, ((d_1504_v65_)[d_1422_v0_] if (d_1422_v0_) in (d_1504_v65_) else d_1422_v0_)])
                nw247_[int(11)] = _dafny.SeqWithoutIsStrInference([d_1422_v0_])
                nw247_[int(12)] = d_1503_v64_
                nw247_[int(13)] = d_1503_v64_
                nw247_[int(14)] = d_1503_v64_
                nw247_[int(15)] = d_1503_v64_
                nw247_[int(16)] = d_1503_v64_
                nw247_[int(17)] = d_1503_v64_
                nw247_[int(18)] = d_1503_v64_
                nw247_[int(19)] = d_1503_v64_
                d_1505_v66_ = nw247_
                d_1506_v67_: D18
                d_1506_v67_ = D18_DC47(d_1505_v66_)
                d_1507_v68_: _dafny.Seq
                d_1507_v68_ = _dafny.SeqWithoutIsStrInference([d_1505_v66_])
                d_1508_v69_: _dafny.Array
                nw248_ = _dafny.Array(None, 25)
                nw248_[int(0)] = d_1505_v66_
                nw248_[int(1)] = d_1505_v66_
                nw248_[int(2)] = d_1505_v66_
                nw248_[int(3)] = d_1505_v66_
                nw248_[int(4)] = d_1505_v66_
                nw248_[int(5)] = d_1505_v66_
                nw248_[int(6)] = d_1505_v66_
                nw248_[int(7)] = d_1505_v66_
                nw248_[int(8)] = d_1505_v66_
                nw248_[int(9)] = d_1505_v66_
                nw248_[int(10)] = (d_1505_v66_ if False else d_1505_v66_)
                nw248_[int(11)] = (d_1506_v67_).cf83
                nw248_[int(12)] = d_1505_v66_
                nw248_[int(13)] = d_1505_v66_
                nw248_[int(14)] = d_1505_v66_
                nw248_[int(15)] = d_1505_v66_
                nw248_[int(16)] = d_1505_v66_
                nw248_[int(17)] = d_1505_v66_
                nw248_[int(18)] = (d_1507_v68_)[default__.safeIndex(-656, len(d_1507_v68_))]
                nw248_[int(19)] = d_1505_v66_
                nw248_[int(20)] = d_1505_v66_
                nw248_[int(21)] = d_1505_v66_
                nw248_[int(22)] = d_1505_v66_
                nw248_[int(23)] = d_1505_v66_
                nw248_[int(24)] = d_1505_v66_
                d_1508_v69_ = nw248_
                index205_ = default__.safeIndex(784, (d_1508_v69_).length(0))
                (d_1508_v69_)[index205_] = d_1505_v66_
                d_1509_v70_: _dafny.Set
                d_1509_v70_ = _dafny.Set({self.f4, (self.f4 if d_1422_v0_ else default__.fm1(globalState)), self.f4})
                d_1510_v71_: _dafny.Array
                def lambda69_(d_1511_v59_):
                    def lambda70_(d_1512_i5_):
                        return (d_1511_v59_)[default__.safeIndex(self.f4, len(d_1511_v59_))]

                    return lambda70_

                init39_ = lambda69_(d_1496_v59_)
                nw249_ = _dafny.Array(None, 13)
                for i0_39_ in range(nw249_.length(0)):
                    nw249_[i0_39_] = init39_(i0_39_)
                d_1510_v71_ = nw249_
                d_1513_v72_: _dafny.MultiSet
                d_1513_v72_ = _dafny.MultiSet([d_1510_v71_, d_1510_v71_])
                d_1514_v73_: D3
                d_1514_v73_ = D3_DC10(self.f4, (0) - (self.f4), d_1422_v0_, self.f4, d_1422_v0_)
                index206_ = default__.safeIndex(878, (d_1500_v63_).length(0))
                index207_ = default__.safeIndex(784, (d_1508_v69_).length(0))
                rhs170_ = (d_1496_v59_) + (d_1496_v59_)
                rhs171_ = d_1422_v0_
                rhs172_ = d_1505_v66_
                rhs173_ = _dafny.Set({(((d_1513_v72_)[d_1510_v71_] if (d_1510_v71_) in (d_1513_v72_) else self.f4)) - ((d_1514_v73_).cf18)})
                lhs88_ = d_1500_v63_
                lhs89_ = default__.safeIndex(878, (d_1500_v63_).length(0))
                lhs90_ = d_1508_v69_
                lhs91_ = default__.safeIndex(784, (d_1508_v69_).length(0))
                lhs88_[lhs89_] = rhs170_
                d_1422_v0_ = rhs171_
                lhs90_[lhs91_] = rhs172_
                d_1509_v70_ = rhs173_
                d_1515_v74_: _dafny.MultiSet
                d_1515_v74_ = _dafny.MultiSet([self.f4, -130, self.f4])
                d_1516_v75_: _dafny.Array
                nw250_ = _dafny.Array(None, 10)
                nw250_[int(0)] = (d_1515_v74_).cardinality
                nw250_[int(1)] = default__.fm1(globalState)
                nw250_[int(2)] = default__.safeDivisionInt(self.f4, self.f4)
                nw250_[int(3)] = default__.safeModuloInt(self.f4, self.f4)
                nw250_[int(4)] = self.f4
                nw250_[int(5)] = self.f4
                nw250_[int(6)] = len(d_1496_v59_)
                nw250_[int(7)] = default__.fm1(globalState)
                nw250_[int(8)] = (self.f4) - (self.f4)
                nw250_[int(9)] = len((d_1496_v59_).set(default__.safeIndex(self.f4, len(d_1496_v59_)), d_1497_v60_))
                d_1516_v75_ = nw250_
                index208_ = default__.safeIndex(326, (d_1516_v75_).length(0))
                (d_1516_v75_)[index208_] = self.f4
                d_1517_v76_: _dafny.Array
                nw251_ = _dafny.Array(None, 8)
                nw251_[int(0)] = False
                nw251_[int(1)] = d_1422_v0_
                nw251_[int(2)] = d_1422_v0_
                nw251_[int(3)] = d_1422_v0_
                nw251_[int(4)] = d_1422_v0_
                nw251_[int(5)] = (d_1503_v64_)[default__.safeIndex(len(d_1509_v70_), len(d_1503_v64_))]
                nw251_[int(6)] = True
                nw251_[int(7)] = d_1422_v0_
                d_1517_v76_ = nw251_
                index209_ = default__.safeIndex(326, (d_1516_v75_).length(0))
                index210_ = default__.safeIndex(878, (d_1500_v63_).length(0))
                rhs174_ = (0) - ((0) - (-448))
                rhs175_ = ((_dafny.SeqWithoutIsStrInference([d_1497_v60_ for d_1518_i6_ in range(default__.abs(663))])).set(default__.safeIndex(self.f4, len(_dafny.SeqWithoutIsStrInference([d_1497_v60_ for d_1519_i6_ in range(default__.abs(663))]))), d_1497_v60_)).set(default__.safeIndex(len((d_1500_v63_)[default__.safeIndex(878, (d_1500_v63_).length(0))]), len((_dafny.SeqWithoutIsStrInference([d_1497_v60_ for d_1520_i6_ in range(default__.abs(663))])).set(default__.safeIndex(self.f4, len(_dafny.SeqWithoutIsStrInference([d_1497_v60_ for d_1521_i6_ in range(default__.abs(663))]))), d_1497_v60_))), d_1497_v60_)
                rhs176_ = (p0 if default__.fm0(self.f4, len(d_1509_v70_), globalState) else d_1517_v76_)
                lhs92_ = d_1516_v75_
                lhs93_ = default__.safeIndex(326, (d_1516_v75_).length(0))
                lhs94_ = d_1500_v63_
                lhs95_ = default__.safeIndex(878, (d_1500_v63_).length(0))
                lhs92_[lhs93_] = rhs174_
                lhs94_[lhs95_] = rhs175_
                d_1517_v76_ = rhs176_
                d_1422_v0_ = not (d_1422_v0_) or (d_1422_v0_)
            elif True:
                d_1522_v77_: C1
                nw252_ = C1()
                nw252_.ctor__(self.f4, self.f4, len((d_1496_v59_) + (default__.fm23(d_1422_v0_, globalState))), d_1422_v0_)
                d_1522_v77_ = nw252_
                d_1422_v0_ = False
                d_1422_v0_ = d_1422_v0_
                (self).f4 = default__.safeModuloInt(self.f4, (self.f4 if d_1422_v0_ else self.f4))
                d_1523_v78_: str
                d_1523_v78_ = _dafny.CodePoint('d')
                d_1524_v79_: _dafny.Array
                def lambda71_(d_1525_v77_):
                    def lambda72_(d_1526_i7_):
                        return default__.safeDivisionInt(d_1526_i7_, d_1525_v77_.f17)

                    return lambda72_

                init40_ = lambda71_(d_1522_v77_)
                nw253_ = _dafny.Array(None, 24)
                for i0_40_ in range(nw253_.length(0)):
                    nw253_[i0_40_] = init40_(i0_40_)
                d_1524_v79_ = nw253_
                index211_ = default__.safeIndex(375, (d_1524_v79_).length(0))
                (d_1524_v79_)[index211_] = -961
                d_1527_v80_: _dafny.MultiSet
                d_1527_v80_ = _dafny.MultiSet([d_1422_v0_])
                d_1528_v81_: D5
                d_1528_v81_ = D5_DC17(d_1527_v80_, d_1422_v0_)
                d_1529_v82_: _dafny.Seq
                d_1529_v82_ = _dafny.SeqWithoutIsStrInference([d_1528_v81_])
                index212_ = default__.safeIndex(375, (d_1524_v79_).length(0))
                rhs177_ = (_dafny.SeqWithoutIsStrInference([d_1528_v81_])) != (d_1529_v82_)
                rhs178_ = self.f4
                rhs179_ = _dafny.CodePoint('t')
                rhs180_ = default__.fm1(globalState)
                rhs181_ = default__.safeDivisionInt(397, ((0) - (d_1522_v77_.f17)) - (self.f4))
                lhs96_ = d_1522_v77_
                lhs97_ = d_1524_v79_
                lhs98_ = default__.safeIndex(375, (d_1524_v79_).length(0))
                lhs99_ = self
                d_1422_v0_ = rhs177_
                lhs96_.f17 = rhs178_
                d_1523_v78_ = rhs179_
                lhs97_[lhs98_] = rhs180_
                lhs99_.f4 = rhs181_
            if d_1422_v0_:
                d_1530_v83_: _dafny.Array
                nw254_ = _dafny.Array(None, 4)
                d_1530_v83_ = nw254_
                d_1531_v84_: _dafny.MultiSet
                d_1531_v84_ = _dafny.MultiSet([d_1422_v0_])
                d_1532_v85_: T1
                nw255_ = C1()
                nw255_.ctor__(((d_1531_v84_)[True] if (True) in (d_1531_v84_) else len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_1533_i8_ in range(default__.abs(459))]))), len(d_1496_v59_), self.f4, d_1422_v0_)
                d_1532_v85_ = nw255_
                index213_ = default__.safeIndex(261, (d_1530_v83_).length(0))
                (d_1530_v83_)[index213_] = d_1532_v85_
                d_1534_v86_: str
                d_1534_v86_ = _dafny.CodePoint('y')
                d_1535_v87_: _dafny.Set
                d_1535_v87_ = _dafny.Set({d_1534_v86_, d_1534_v86_})
                d_1536_v88_: _dafny.Set
                d_1536_v88_ = _dafny.Set({d_1422_v0_, d_1422_v0_})
                index214_ = default__.safeIndex(261, (d_1530_v83_).length(0))
                rhs182_ = d_1422_v0_
                rhs183_ = (d_1496_v59_) + (d_1496_v59_)
                rhs184_ = (d_1535_v87_).issubset(_dafny.Set({d_1534_v86_, d_1534_v86_}))
                rhs185_ = ((self.f4) - (len(d_1536_v88_))) != (self.f4)
                rhs186_ = (d_1532_v85_ if d_1422_v0_ else d_1532_v85_)
                lhs100_ = d_1530_v83_
                lhs101_ = default__.safeIndex(261, (d_1530_v83_).length(0))
                d_1422_v0_ = rhs182_
                d_1496_v59_ = rhs183_
                d_1422_v0_ = rhs184_
                d_1422_v0_ = rhs185_
                lhs100_[lhs101_] = rhs186_
                d_1496_v59_ = ((d_1496_v59_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sdxe")))) + ((d_1496_v59_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vat"))))
                d_1422_v0_ = (default__.safeDivisionInt(default__.fm1(globalState), self.f4)) <= ((0) - (self.f4))
                d_1537_v89_: _dafny.Seq
                d_1537_v89_ = _dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wiu"))))])
                d_1538_v90_: _dafny.Map
                d_1538_v90_ = _dafny.Map({not(not(d_1422_v0_)): d_1422_v0_})
                d_1539_v91_: _dafny.Map
                d_1539_v91_ = _dafny.Map({904: len(d_1538_v90_)})
                d_1540_v92_: _dafny.MultiSet
                d_1540_v92_ = _dafny.MultiSet([self.f4, (0) - (d_1532_v85_.f4)])
                d_1541_v93_: _dafny.Seq
                d_1541_v93_ = _dafny.SeqWithoutIsStrInference([d_1422_v0_])
                d_1542_v94_: D17
                d_1542_v94_ = D17_DC46(d_1422_v0_, self.f4, default__.fm1(globalState), d_1422_v0_)
                d_1543_v95_: T2
                nw256_ = C1()
                nw256_.ctor__(len(d_1541_v93_), (d_1542_v94_).cf81, self.f4, d_1422_v0_)
                d_1543_v95_ = nw256_
                d_1544_v96_: D6
                d_1544_v96_ = D6_DC21(len(d_1496_v59_), True)
                d_1545_v97_: _dafny.Array
                nw257_ = _dafny.Array(None, 23)
                nw257_[int(0)] = d_1532_v85_.f4
                nw257_[int(1)] = d_1532_v85_.f4
                nw257_[int(2)] = d_1532_v85_.f4
                nw257_[int(3)] = (d_1537_v89_)[default__.safeIndex(self.f4, len(d_1537_v89_))]
                nw257_[int(4)] = len(d_1496_v59_)
                nw257_[int(5)] = (0) - ((self.f4) * (len(d_1496_v59_)))
                nw257_[int(6)] = len(d_1539_v91_)
                nw257_[int(7)] = d_1532_v85_.f4
                nw257_[int(8)] = self.f4
                nw257_[int(9)] = default__.safeModuloInt(d_1532_v85_.f4, self.f4)
                nw257_[int(10)] = d_1532_v85_.f4
                nw257_[int(11)] = d_1532_v85_.f4
                nw257_[int(12)] = (0) - (((d_1540_v92_)[len(_dafny.Map({len(_dafny.Set({d_1422_v0_})): d_1543_v95_}))] if (len(_dafny.Map({len(_dafny.Set({d_1422_v0_})): d_1543_v95_}))) in (d_1540_v92_) else d_1532_v85_.f4))
                nw257_[int(13)] = default__.safeModuloInt(self.f4, d_1532_v85_.f4)
                nw257_[int(14)] = default__.safeDivisionInt((0) - (len(d_1496_v59_)), (d_1543_v95_).f15)
                nw257_[int(15)] = self.f4
                nw257_[int(16)] = (d_1543_v95_).f15
                nw257_[int(17)] = d_1532_v85_.f4
                nw257_[int(18)] = default__.safeDivisionInt(526, 524)
                nw257_[int(19)] = len(d_1537_v89_)
                nw257_[int(20)] = d_1532_v85_.f4
                nw257_[int(21)] = (d_1544_v96_).cf41
                nw257_[int(22)] = (d_1543_v95_).f15
                d_1545_v97_ = nw257_
                d_1546_v98_: _dafny.Map
                d_1546_v98_ = _dafny.Map({d_1422_v0_: d_1532_v85_.f4})
                index215_ = default__.safeIndex(9, (d_1545_v97_).length(0))
                (d_1545_v97_)[index215_] = len(d_1546_v98_)
                index216_ = default__.safeIndex(326, (p0).length(0))
                (p0)[index216_] = (d_1541_v93_) < (d_1541_v93_)
                index217_ = default__.safeIndex(9, (d_1545_v97_).length(0))
                index218_ = default__.safeIndex(326, (p0).length(0))
                rhs187_ = default__.safeDivisionInt(d_1532_v85_.f4, self.f4)
                rhs188_ = d_1422_v0_
                lhs102_ = d_1545_v97_
                lhs103_ = default__.safeIndex(9, (d_1545_v97_).length(0))
                lhs104_ = p0
                lhs105_ = default__.safeIndex(326, (p0).length(0))
                lhs102_[lhs103_] = rhs187_
                lhs104_[lhs105_] = rhs188_
                d_1547_v99_: _dafny.Array
                nw258_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 3)
                d_1547_v99_ = nw258_
                d_1548_v100_: _dafny.Array
                nw259_ = _dafny.Array(None, 1)
                nw259_[int(0)] = D8_DC28(d_1547_v99_)
                d_1548_v100_ = nw259_
                d_1549_v101_: _dafny.Seq
                d_1549_v101_ = _dafny.SeqWithoutIsStrInference([d_1548_v100_, d_1548_v100_, d_1548_v100_, d_1548_v100_])
                d_1550_v102_: _dafny.Set
                d_1550_v102_ = _dafny.Set({-901})
                d_1551_v103_: _dafny.Map
                d_1551_v103_ = _dafny.Map({d_1532_v85_.f4: d_1548_v100_})
                d_1552_v104_: D19
                d_1552_v104_ = D19_DC50(d_1548_v100_)
                d_1553_v105_: _dafny.Array
                nw260_ = _dafny.Array(None, 23)
                nw260_[int(0)] = d_1548_v100_
                nw260_[int(1)] = d_1548_v100_
                nw260_[int(2)] = d_1548_v100_
                nw260_[int(3)] = d_1548_v100_
                nw260_[int(4)] = d_1548_v100_
                nw260_[int(5)] = (d_1549_v101_)[default__.safeIndex(len(d_1550_v102_), len(d_1549_v101_))]
                nw260_[int(6)] = d_1548_v100_
                nw260_[int(7)] = d_1548_v100_
                nw260_[int(8)] = d_1548_v100_
                nw260_[int(9)] = ((d_1551_v103_)[(d_1545_v97_)[default__.safeIndex(9, (d_1545_v97_).length(0))]] if ((d_1545_v97_)[default__.safeIndex(9, (d_1545_v97_).length(0))]) in (d_1551_v103_) else d_1548_v100_)
                nw260_[int(10)] = d_1548_v100_
                nw260_[int(11)] = d_1548_v100_
                nw260_[int(12)] = d_1548_v100_
                nw260_[int(13)] = (d_1548_v100_ if (d_1543_v95_).f16 else d_1548_v100_)
                nw260_[int(14)] = (d_1552_v104_).cf89
                nw260_[int(15)] = d_1548_v100_
                nw260_[int(16)] = d_1548_v100_
                nw260_[int(17)] = d_1548_v100_
                nw260_[int(18)] = d_1548_v100_
                nw260_[int(19)] = d_1548_v100_
                nw260_[int(20)] = d_1548_v100_
                nw260_[int(21)] = d_1548_v100_
                nw260_[int(22)] = d_1548_v100_
                d_1553_v105_ = nw260_
                index219_ = default__.safeIndex(338, (d_1553_v105_).length(0))
                (d_1553_v105_)[index219_] = (d_1552_v104_).cf89
                index220_ = default__.safeIndex(338, (d_1553_v105_).length(0))
                (d_1553_v105_)[index220_] = d_1548_v100_
            elif True:
                d_1554_v106_: C0
                nw261_ = C0()
                nw261_.ctor__(self.f4)
                d_1554_v106_ = nw261_
                d_1555_v107_: _dafny.MultiSet
                d_1555_v107_ = _dafny.MultiSet([d_1554_v106_, d_1554_v106_, d_1554_v106_, d_1554_v106_, d_1554_v106_])
                d_1556_v108_: C3
                nw262_ = C3()
                nw262_.ctor__((d_1555_v107_).cardinality, d_1496_v59_, self.f4)
                d_1556_v108_ = nw262_
                d_1557_v109_: _dafny.Seq
                d_1557_v109_ = _dafny.SeqWithoutIsStrInference([(d_1556_v108_).f11])
                d_1558_v110_: _dafny.Set
                d_1558_v110_ = _dafny.Set({d_1422_v0_, d_1422_v0_})
                d_1559_v111_: str
                d_1559_v111_ = _dafny.CodePoint('o')
                d_1560_v112_: _dafny.Array
                nw263_ = _dafny.Array(None, 21)
                nw263_[int(0)] = d_1554_v106_.f18
                nw263_[int(1)] = (0) - ((d_1556_v108_).f11)
                nw263_[int(2)] = -839
                nw263_[int(3)] = (d_1556_v108_).f11
                nw263_[int(4)] = len((d_1557_v109_).set(default__.safeIndex(self.f4, len(d_1557_v109_)), (d_1556_v108_).f11))
                nw263_[int(5)] = len(d_1558_v110_)
                nw263_[int(6)] = (0) - (len(d_1557_v109_))
                nw263_[int(7)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mfumfjre")))
                nw263_[int(8)] = (d_1556_v108_).f11
                nw263_[int(9)] = len(_dafny.SeqWithoutIsStrInference([self.f4 for d_1561_i9_ in range(default__.abs(324))]))
                nw263_[int(10)] = 527
                nw263_[int(11)] = (d_1556_v108_).f11
                nw263_[int(12)] = (0) - ((d_1556_v108_).f11)
                nw263_[int(13)] = len(default__.fm2(d_1559_v111_, d_1554_v106_.f18, len((d_1556_v108_).f12), globalState))
                nw263_[int(14)] = self.f4
                nw263_[int(15)] = self.f4
                nw263_[int(16)] = (d_1556_v108_).f11
                nw263_[int(17)] = (d_1556_v108_).f11
                nw263_[int(18)] = d_1554_v106_.f18
                nw263_[int(19)] = self.f4
                nw263_[int(20)] = (d_1556_v108_).f11
                d_1560_v112_ = nw263_
                d_1562_v113_: _dafny.Array
                nw264_ = _dafny.Array(None, 16)
                nw264_[int(0)] = d_1560_v112_
                nw264_[int(1)] = d_1560_v112_
                nw264_[int(2)] = d_1560_v112_
                nw264_[int(3)] = d_1560_v112_
                nw264_[int(4)] = d_1560_v112_
                nw264_[int(5)] = d_1560_v112_
                nw264_[int(6)] = d_1560_v112_
                nw264_[int(7)] = d_1560_v112_
                nw264_[int(8)] = d_1560_v112_
                nw264_[int(9)] = d_1560_v112_
                nw264_[int(10)] = d_1560_v112_
                nw264_[int(11)] = d_1560_v112_
                nw264_[int(12)] = d_1560_v112_
                nw264_[int(13)] = d_1560_v112_
                nw264_[int(14)] = d_1560_v112_
                nw264_[int(15)] = d_1560_v112_
                d_1562_v113_ = nw264_
                index221_ = default__.safeIndex(977, (d_1562_v113_).length(0))
                (d_1562_v113_)[index221_] = d_1560_v112_
                d_1563_v114_: _dafny.Array
                nw265_ = _dafny.Array(D11.default()(), 1)
                d_1563_v114_ = nw265_
                d_1564_v115_: D11
                d_1564_v115_ = D11_DC35((d_1556_v108_).f11, d_1422_v0_, d_1422_v0_)
                index222_ = default__.safeIndex(927, (d_1563_v114_).length(0))
                (d_1563_v114_)[index222_] = d_1564_v115_
                pat_let_tv44_ = d_1422_v0_
                pat_let_tv45_ = d_1422_v0_
                index223_ = default__.safeIndex(977, (d_1562_v113_).length(0))
                index224_ = default__.safeIndex(927, (d_1563_v114_).length(0))
                def iife162_(_pat_let41_0):
                    def iife163_(d_1565_dt__update__tmp_h0_):
                        def iife164_(_pat_let42_0):
                            def iife165_(d_1566_dt__update_hcf62_h0_):
                                def iife166_(_pat_let43_0):
                                    def iife167_(d_1567_dt__update_hcf63_h0_):
                                        return D11_DC35((d_1565_dt__update__tmp_h0_).cf61, d_1566_dt__update_hcf62_h0_, d_1567_dt__update_hcf63_h0_)
                                    return iife167_(_pat_let43_0)
                                return iife166_(pat_let_tv45_)
                            return iife165_(_pat_let42_0)
                        return iife164_(pat_let_tv44_)
                    return iife163_(_pat_let41_0)
                rhs189_ = not(not(d_1422_v0_))
                rhs190_ = d_1560_v112_
                rhs191_ = (iife162_(d_1564_v115_) if default__.fm0((0) - (self.f4), d_1554_v106_.f18, globalState) else d_1564_v115_)
                rhs192_ = (d_1422_v0_) and (d_1422_v0_)
                lhs106_ = d_1562_v113_
                lhs107_ = default__.safeIndex(977, (d_1562_v113_).length(0))
                lhs108_ = d_1563_v114_
                lhs109_ = default__.safeIndex(927, (d_1563_v114_).length(0))
                d_1422_v0_ = rhs189_
                lhs106_[lhs107_] = rhs190_
                lhs108_[lhs109_] = rhs191_
                d_1422_v0_ = rhs192_
                d_1568_v116_: _dafny.Map
                d_1568_v116_ = _dafny.Map({d_1422_v0_: (d_1562_v113_)[default__.safeIndex(977, (d_1562_v113_).length(0))]})
                d_1560_v112_ = ((d_1568_v116_)[d_1422_v0_] if (d_1422_v0_) in (d_1568_v116_) else (d_1562_v113_)[default__.safeIndex(977, (d_1562_v113_).length(0))])
                d_1569_v117_: _dafny.Map
                d_1569_v117_ = _dafny.Map({d_1422_v0_: d_1422_v0_})
                d_1570_v118_: _dafny.Seq
                d_1570_v118_ = _dafny.SeqWithoutIsStrInference([d_1569_v117_])
                d_1571_v119_: _dafny.Seq
                d_1571_v119_ = _dafny.SeqWithoutIsStrInference([d_1570_v118_, d_1570_v118_])
                d_1572_v120_: C1
                nw266_ = C1()
                nw266_.ctor__(d_1554_v106_.f18, len((d_1571_v119_)[default__.safeIndex((0) - (d_1554_v106_.f18), len(d_1571_v119_))]), default__.safeDivisionInt(d_1554_v106_.f18, 357), True)
                d_1572_v120_ = nw266_
                (d_1572_v120_).f17 = self.f4
            d_1573_v121_: _dafny.Map
            d_1573_v121_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_1574_i10_ in range(default__.abs(700))]): self.f4})
            d_1573_v121_ = (d_1573_v121_).set(d_1496_v59_, self.f4)
            d_1575_v122_: _dafny.Array
            def lambda73_(d_1576_i11_):
                return (d_1576_i11_) + (self.f4)

            init41_ = lambda73_
            nw267_ = _dafny.Array(None, 25)
            for i0_41_ in range(nw267_.length(0)):
                nw267_[i0_41_] = init41_(i0_41_)
            d_1575_v122_ = nw267_
            d_1575_v122_ = d_1575_v122_
            d_1577_v123_: _dafny.Map
            d_1577_v123_ = _dafny.Map({(default__.fm37(d_1422_v0_, d_1422_v0_, self.f4, globalState)).set(self.f4, d_1422_v0_): self.f4})
            d_1578_v124_: _dafny.Map
            d_1578_v124_ = _dafny.Map({self.f4: d_1422_v0_})
            d_1579_v125_: _dafny.Map
            d_1579_v125_ = _dafny.Map({d_1575_v122_: ((d_1577_v123_)[d_1578_v124_] if (d_1578_v124_) in (d_1577_v123_) else self.f4)})
            d_1579_v125_ = d_1579_v125_
        elif True:
            d_1580_v126_: _dafny.Seq
            d_1580_v126_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pdukqcif"))
            d_1581_v127_: D12
            d_1581_v127_ = D12_DC37(True, d_1422_v0_, False)
            d_1580_v126_ = (d_1580_v126_ if (d_1581_v127_).cf66 else d_1580_v126_)
            d_1582_v128_: _dafny.Array
            nw268_ = _dafny.Array(_dafny.Array(None, 0), 3)
            d_1582_v128_ = nw268_
            index225_ = default__.safeIndex(87, (d_1582_v128_).length(0))
            (d_1582_v128_)[index225_] = p0
            d_1583_v129_: _dafny.Map
            d_1583_v129_ = _dafny.Map({self.f4: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i")))})
            d_1584_v130_: _dafny.Map
            d_1584_v130_ = _dafny.Map({self.f4: d_1422_v0_})
            d_1585_v131_: _dafny.Map
            d_1585_v131_ = _dafny.Map({((d_1584_v130_)[self.f4] if (self.f4) in (d_1584_v130_) else d_1422_v0_): self.f4})
            d_1586_v132_: _dafny.Seq
            d_1586_v132_ = _dafny.SeqWithoutIsStrInference([p0])
            index226_ = default__.safeIndex(87, (d_1582_v128_).length(0))
            rhs193_ = (d_1586_v132_)[default__.safeIndex(default__.fm1(globalState), len(d_1586_v132_))]
            rhs194_ = d_1583_v129_
            rhs195_ = d_1585_v131_
            lhs110_ = d_1582_v128_
            lhs111_ = default__.safeIndex(87, (d_1582_v128_).length(0))
            lhs110_[lhs111_] = rhs193_
            d_1583_v129_ = rhs194_
            d_1585_v131_ = rhs195_
            d_1587_v133_: _dafny.Array
            def lambda74_(d_1588_i12_):
                return _dafny.CodePoint('u')

            init42_ = lambda74_
            nw269_ = _dafny.Array(None, 7)
            for i0_42_ in range(nw269_.length(0)):
                nw269_[i0_42_] = init42_(i0_42_)
            d_1587_v133_ = nw269_
            d_1589_v134_: _dafny.Map
            d_1589_v134_ = _dafny.Map({d_1587_v133_: self.f4})
            d_1590_v135_: C1
            nw270_ = C1()
            nw270_.ctor__(self.f4, self.f4, ((d_1589_v134_)[d_1587_v133_] if (d_1587_v133_) in (d_1589_v134_) else self.f4), d_1422_v0_)
            d_1590_v135_ = nw270_
            d_1591_v136_: _dafny.Array
            nw271_ = _dafny.Array(None, 28)
            d_1591_v136_ = nw271_
            d_1591_v136_ = d_1591_v136_
            d_1592_v137_: str
            d_1592_v137_ = _dafny.CodePoint('d')
            d_1593_v138_: C3
            nw272_ = C3()
            nw272_.ctor__(self.f4, _dafny.SeqWithoutIsStrInference([d_1592_v137_, d_1592_v137_]), d_1590_v135_.f17)
            d_1593_v138_ = nw272_
        d_1594_v139_: _dafny.Seq
        d_1594_v139_ = _dafny.SeqWithoutIsStrInference([default__.fm0(-355, self.f4, globalState)])
        d_1595_i13_: int
        d_1595_i13_ = 0
        with _dafny.label("8"):
            while (d_1594_v139_)[default__.safeIndex((0) - (self.f4), len(d_1594_v139_))]:
                with _dafny.c_label("8"):
                    if (d_1595_i13_) >= (100):
                        raise _dafny.Break("8")
                    d_1595_i13_ = (d_1595_i13_) + (1)
                    d_1422_v0_ = True
                    d_1596_v140_: _dafny.Array
                    def lambda75_(d_1597_i14_):
                        return (_dafny.SeqWithoutIsStrInference([self.f4])) + (_dafny.SeqWithoutIsStrInference([self.f4]))

                    init43_ = lambda75_
                    nw273_ = _dafny.Array(None, 9)
                    for i0_43_ in range(nw273_.length(0)):
                        nw273_[i0_43_] = init43_(i0_43_)
                    d_1596_v140_ = nw273_
                    def lambda76_(d_1598_i15_):
                        return _dafny.SeqWithoutIsStrInference([self.f4, self.f4])

                    init44_ = lambda76_
                    nw274_ = _dafny.Array(None, 2)
                    for i0_44_ in range(nw274_.length(0)):
                        nw274_[i0_44_] = init44_(i0_44_)
                    d_1596_v140_ = nw274_
                    (self).f4 = default__.fm1(globalState)
                    (self).f4 = self.f4
                    pass
            pass
        (self).f4 = self.f4
        (self).f4 = (0) - ((self.f4) + (self.f4))
        if (d_1594_v139_)[default__.safeIndex(self.f4, len(d_1594_v139_))]:
            d_1599_v141_: _dafny.Seq
            d_1599_v141_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eijm"))
            index227_ = default__.safeIndex(669, (p0).length(0))
            (p0)[index227_] = d_1422_v0_
            d_1600_v142_: _dafny.Array
            nw275_ = _dafny.Array(int(0), 18)
            d_1600_v142_ = nw275_
            index228_ = default__.safeIndex(891, (d_1600_v142_).length(0))
            (d_1600_v142_)[index228_] = self.f4
            d_1601_v143_: _dafny.Map
            d_1601_v143_ = _dafny.Map({d_1599_v141_: self.f4})
            d_1602_v144_: str
            d_1602_v144_ = _dafny.CodePoint('h')
            d_1603_v145_: _dafny.MultiSet
            d_1603_v145_ = _dafny.MultiSet([d_1602_v144_])
            d_1604_v146_: _dafny.Map
            d_1604_v146_ = _dafny.Map({self.f4: d_1602_v144_})
            index229_ = default__.safeIndex(669, (p0).length(0))
            index230_ = default__.safeIndex(891, (d_1600_v142_).length(0))
            rhs196_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "quc"))) + (d_1599_v141_)
            rhs197_ = 600
            rhs198_ = not((d_1599_v141_) not in (d_1601_v143_))
            rhs199_ = self.f4
            rhs200_ = (len(d_1599_v141_)) + (((d_1603_v145_ if d_1422_v0_ else _dafny.MultiSet([((d_1604_v146_)[self.f4] if (self.f4) in (d_1604_v146_) else _dafny.CodePoint('h'))]))).cardinality)
            lhs112_ = self
            lhs113_ = p0
            lhs114_ = default__.safeIndex(669, (p0).length(0))
            lhs115_ = d_1600_v142_
            lhs116_ = default__.safeIndex(891, (d_1600_v142_).length(0))
            lhs117_ = self
            d_1599_v141_ = rhs196_
            lhs112_.f4 = rhs197_
            lhs113_[lhs114_] = rhs198_
            lhs115_[lhs116_] = rhs199_
            lhs117_.f4 = rhs200_
            d_1605_v147_: _dafny.Map
            d_1605_v147_ = _dafny.Map({default__.safeModuloInt((d_1600_v142_)[default__.safeIndex(891, (d_1600_v142_).length(0))], self.f4): d_1422_v0_})
            d_1422_v0_ = ((d_1605_v147_)[-887] if (-887) in (d_1605_v147_) else (d_1422_v0_ if (p0)[default__.safeIndex(669, (p0).length(0))] else d_1422_v0_))
            index231_ = default__.safeIndex(891, (d_1600_v142_).length(0))
            (d_1600_v142_)[index231_] = (d_1600_v142_)[default__.safeIndex(891, (d_1600_v142_).length(0))]
            index232_ = default__.safeIndex(891, (d_1600_v142_).length(0))
            (d_1600_v142_)[index232_] = ((d_1600_v142_)[default__.safeIndex(891, (d_1600_v142_).length(0))]) - (self.f4)
            d_1606_v148_: _dafny.Array
            nw276_ = _dafny.Array(_dafny.Set({}), 6)
            d_1606_v148_ = nw276_
            d_1607_v149_: _dafny.Array
            nw277_ = _dafny.Array(_dafny.Map({}), 15)
            d_1607_v149_ = nw277_
            d_1608_v151_: _dafny.Map
            d_1608_v151_ = _dafny.Map({self.f4: _dafny.Set({_dafny.CodePoint('t'), d_1602_v144_})})
            d_1609_v152_: _dafny.Set
            d_1609_v152_ = _dafny.Set({d_1602_v144_})
            d_1610_v153_: D17
            def iife168_():
                coll80_ = _dafny.Map()
                compr_80_: str
                for compr_80_ in (((d_1608_v151_)[self.f4] if (self.f4) in (d_1608_v151_) else d_1609_v152_)).Elements:
                    d_1611_v150_: str = compr_80_
                    if (d_1611_v150_) in (((d_1608_v151_)[self.f4] if (self.f4) in (d_1608_v151_) else d_1609_v152_)):
                        coll80_[d_1611_v150_] = _dafny.MultiSet([d_1422_v0_, (p0)[default__.safeIndex(669, (p0).length(0))]])
                return _dafny.Map(coll80_)
            d_1610_v153_ = D17_DC45(iife168_()
)
            d_1612_v154_: _dafny.MultiSet
            d_1612_v154_ = _dafny.MultiSet([(p0)[default__.safeIndex(669, (p0).length(0))], (p0)[default__.safeIndex(669, (p0).length(0))], d_1422_v0_, (p0)[default__.safeIndex(669, (p0).length(0))], False])
            d_1613_v155_: _dafny.Map
            d_1613_v155_ = _dafny.Map({d_1602_v144_: d_1612_v154_})
            pat_let_tv46_ = d_1613_v155_
            d_1614_v156_: _dafny.Map
            def iife169_(_pat_let44_0):
                def iife170_(d_1615_dt__update__tmp_h1_):
                    def iife171_(_pat_let45_0):
                        def iife172_(d_1616_dt__update_hcf78_h0_):
                            return D17_DC45(d_1616_dt__update_hcf78_h0_)
                        return iife172_(_pat_let45_0)
                    return iife171_(pat_let_tv46_)
                return iife170_(_pat_let44_0)
            d_1614_v156_ = _dafny.Map({self.f4: iife169_(d_1610_v153_)})
            index233_ = default__.safeIndex(293, (d_1607_v149_).length(0))
            (d_1607_v149_)[index233_] = (d_1614_v156_) | (d_1614_v156_)
            d_1617_v157_: _dafny.MultiSet
            d_1617_v157_ = _dafny.MultiSet([(d_1600_v142_)[default__.safeIndex(891, (d_1600_v142_).length(0))], (d_1600_v142_)[default__.safeIndex(891, (d_1600_v142_).length(0))], self.f4, (d_1600_v142_)[default__.safeIndex(891, (d_1600_v142_).length(0))], (d_1600_v142_)[default__.safeIndex(891, (d_1600_v142_).length(0))]])
            d_1618_v158_: _dafny.Map
            d_1618_v158_ = _dafny.Map({self.f4: d_1600_v142_})
            d_1619_v159_: _dafny.Map
            d_1619_v159_ = _dafny.Map({(p0)[default__.safeIndex(669, (p0).length(0))]: d_1600_v142_})
            d_1620_v160_: _dafny.Array
            nw278_ = _dafny.Array(None, 22)
            nw278_[int(0)] = d_1600_v142_
            nw278_[int(1)] = d_1600_v142_
            nw278_[int(2)] = d_1600_v142_
            nw278_[int(3)] = ((d_1618_v158_)[(d_1600_v142_)[default__.safeIndex(891, (d_1600_v142_).length(0))]] if ((d_1600_v142_)[default__.safeIndex(891, (d_1600_v142_).length(0))]) in (d_1618_v158_) else d_1600_v142_)
            nw278_[int(4)] = d_1600_v142_
            nw278_[int(5)] = d_1600_v142_
            nw278_[int(6)] = d_1600_v142_
            nw278_[int(7)] = d_1600_v142_
            nw278_[int(8)] = d_1600_v142_
            nw278_[int(9)] = (d_1600_v142_ if (p0)[default__.safeIndex(669, (p0).length(0))] else d_1600_v142_)
            nw278_[int(10)] = d_1600_v142_
            nw278_[int(11)] = d_1600_v142_
            nw278_[int(12)] = ((d_1619_v159_)[(p0)[default__.safeIndex(669, (p0).length(0))]] if ((p0)[default__.safeIndex(669, (p0).length(0))]) in (d_1619_v159_) else d_1600_v142_)
            nw278_[int(13)] = d_1600_v142_
            nw278_[int(14)] = d_1600_v142_
            nw278_[int(15)] = d_1600_v142_
            nw278_[int(16)] = d_1600_v142_
            nw278_[int(17)] = d_1600_v142_
            nw278_[int(18)] = d_1600_v142_
            nw278_[int(19)] = d_1600_v142_
            nw278_[int(20)] = d_1600_v142_
            nw278_[int(21)] = d_1600_v142_
            d_1620_v160_ = nw278_
            d_1621_v161_: _dafny.Map
            d_1621_v161_ = _dafny.Map({d_1422_v0_: False})
            d_1622_v162_: D1
            d_1622_v162_ = D1_DC4(935, d_1594_v139_, (d_1600_v142_)[default__.safeIndex(891, (d_1600_v142_).length(0))])
            d_1623_v163_: C0
            nw279_ = C0()
            nw279_.ctor__((d_1600_v142_)[default__.safeIndex(891, (d_1600_v142_).length(0))])
            d_1623_v163_ = nw279_
            index234_ = default__.safeIndex(293, (d_1607_v149_).length(0))
            rhs201_ = (d_1606_v148_ if (p0)[default__.safeIndex(669, (p0).length(0))] else d_1606_v148_)
            rhs202_ = d_1614_v156_
            rhs203_ = (d_1617_v157_) - ((d_1617_v157_ if ((d_1621_v161_)[False] if (False) in (d_1621_v161_) else (D18_DC48((p0)[default__.safeIndex(669, (p0).length(0))], False, d_1622_v162_, d_1623_v163_)).cf84) else d_1617_v157_))
            rhs204_ = d_1620_v160_
            lhs118_ = d_1607_v149_
            lhs119_ = default__.safeIndex(293, (d_1607_v149_).length(0))
            d_1606_v148_ = rhs201_
            lhs118_[lhs119_] = rhs202_
            d_1617_v157_ = rhs203_
            d_1620_v160_ = rhs204_
        elif True:
            d_1624_v164_: str
            d_1624_v164_ = _dafny.CodePoint('a')
            d_1625_v165_: _dafny.Map
            d_1625_v165_ = _dafny.Map({340: self.f4})
            d_1626_v166_: D2
            d_1626_v166_ = D2_DC6(self.f4, d_1422_v0_)
            d_1624_v164_ = (self).fm10((self.f4 if d_1422_v0_ else len(d_1625_v165_)), False, d_1422_v0_, d_1626_v166_, globalState)
            (self).f4 = ((self.f4) * (self.f4)) - (self.f4)
            d_1627_v167_: _dafny.MultiSet
            d_1627_v167_ = _dafny.MultiSet([d_1422_v0_])
            d_1628_v168_: _dafny.Set
            d_1628_v168_ = _dafny.Set({self.f4})
            d_1629_v169_: _dafny.Seq
            d_1629_v169_ = _dafny.SeqWithoutIsStrInference([d_1628_v168_])
            d_1630_v170_: _dafny.Seq
            d_1630_v170_ = _dafny.SeqWithoutIsStrInference([601, (d_1627_v167_).cardinality, len((d_1629_v169_)[default__.safeIndex(self.f4, len(d_1629_v169_))]), self.f4, -383])
            d_1631_v171_: C1
            nw280_ = C1()
            nw280_.ctor__((d_1630_v170_)[default__.safeIndex(self.f4, len(d_1630_v170_))], self.f4, self.f4, d_1422_v0_)
            d_1631_v171_ = nw280_
            d_1632_v172_: _dafny.Seq
            d_1632_v172_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a')])
            d_1633_v173_: D12
            d_1633_v173_ = D12_DC38(d_1631_v171_.f17)
            d_1634_v174_: _dafny.Seq
            d_1634_v174_ = _dafny.SeqWithoutIsStrInference([d_1633_v173_, D12_DC38(self.f4)])
            d_1635_v175_: C3
            nw281_ = C3()
            nw281_.ctor__(len((d_1629_v169_)[default__.safeIndex(len(d_1594_v139_), len(d_1629_v169_))]), ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g')])).set(default__.safeIndex(d_1631_v171_.f17, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g')]))), d_1624_v164_)) + (d_1632_v172_), (len(d_1634_v174_)) + (self.f4))
            d_1635_v175_ = nw281_
            d_1636_v176_: _dafny.MultiSet
            d_1636_v176_ = _dafny.MultiSet(d_1630_v170_)
            d_1637_v177_: _dafny.MultiSet
            d_1637_v177_ = _dafny.MultiSet([d_1636_v176_, _dafny.MultiSet([(d_1635_v175_).f11])])
            d_1637_v177_ = _dafny.MultiSet([d_1636_v176_, d_1636_v176_])

    def m9(self, globalState):
        r0: str = _dafny.CodePoint('D')
        d_1638_v0_: _dafny.Array
        nw282_ = _dafny.Array(_dafny.Array(None, 0), 18)
        d_1638_v0_ = nw282_
        d_1639_v1_: bool
        d_1639_v1_ = False
        d_1640_v2_: D13
        d_1640_v2_ = D13_DC40(d_1639_v1_)
        d_1641_v3_: _dafny.Array
        nw283_ = _dafny.Array(None, 17)
        nw283_[int(0)] = d_1639_v1_
        nw283_[int(1)] = d_1639_v1_
        nw283_[int(2)] = not(d_1639_v1_)
        nw283_[int(3)] = True
        nw283_[int(4)] = d_1639_v1_
        nw283_[int(5)] = d_1639_v1_
        nw283_[int(6)] = False
        nw283_[int(7)] = d_1639_v1_
        nw283_[int(8)] = (d_1640_v2_).cf70
        nw283_[int(9)] = d_1639_v1_
        nw283_[int(10)] = d_1639_v1_
        nw283_[int(11)] = False
        nw283_[int(12)] = d_1639_v1_
        nw283_[int(13)] = True
        nw283_[int(14)] = d_1639_v1_
        nw283_[int(15)] = d_1639_v1_
        nw283_[int(16)] = d_1639_v1_
        d_1641_v3_ = nw283_
        index235_ = default__.safeIndex(908, (d_1638_v0_).length(0))
        (d_1638_v0_)[index235_] = d_1641_v3_
        index236_ = default__.safeIndex(699, (d_1641_v3_).length(0))
        (d_1641_v3_)[index236_] = d_1639_v1_
        d_1642_v4_: str
        d_1642_v4_ = _dafny.CodePoint('j')
        index237_ = default__.safeIndex(908, (d_1638_v0_).length(0))
        index238_ = default__.safeIndex(699, (d_1641_v3_).length(0))
        rhs205_ = (D3_DC8(d_1642_v4_, d_1641_v3_)).cf14
        rhs206_ = not(d_1639_v1_)
        lhs120_ = d_1638_v0_
        lhs121_ = default__.safeIndex(908, (d_1638_v0_).length(0))
        lhs122_ = d_1641_v3_
        lhs123_ = default__.safeIndex(699, (d_1641_v3_).length(0))
        lhs120_[lhs121_] = rhs205_
        lhs122_[lhs123_] = rhs206_
        d_1643_v5_: _dafny.Set
        d_1643_v5_ = _dafny.Set({d_1639_v1_})
        d_1644_v6_: _dafny.Map
        d_1644_v6_ = _dafny.Map({len(d_1643_v5_): d_1639_v1_})
        d_1645_v7_: _dafny.Seq
        d_1645_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ttanqrjm"))
        d_1646_v8_: _dafny.Map
        d_1646_v8_ = _dafny.Map({len(d_1644_v6_): len(d_1645_v7_)})
        d_1647_v9_: _dafny.Map
        d_1647_v9_ = _dafny.Map({(d_1641_v3_)[default__.safeIndex(699, (d_1641_v3_).length(0))]: (d_1641_v3_)[default__.safeIndex(699, (d_1641_v3_).length(0))]})
        d_1648_v10_: _dafny.Seq
        d_1648_v10_ = _dafny.SeqWithoutIsStrInference([len(d_1647_v9_), len(d_1644_v6_)])
        d_1649_v11_: _dafny.Set
        d_1649_v11_ = _dafny.Set({len(d_1645_v7_), (d_1648_v10_)[default__.safeIndex(self.f4, len(d_1648_v10_))], self.f4, self.f4, self.f4})
        d_1650_v12_: _dafny.Seq
        d_1650_v12_ = _dafny.SeqWithoutIsStrInference([d_1649_v11_])
        d_1651_i0_: int
        d_1651_i0_ = 0
        with _dafny.label("9"):
            while not(((len(d_1646_v8_)) in (d_1646_v8_)) and ((d_1650_v12_) < (d_1650_v12_))):
                with _dafny.c_label("9"):
                    if (d_1651_i0_) >= (100):
                        raise _dafny.Break("9")
                    d_1651_i0_ = (d_1651_i0_) + (1)
                    d_1652_v13_: _dafny.MultiSet
                    d_1652_v13_ = _dafny.MultiSet([(d_1641_v3_)[default__.safeIndex(699, (d_1641_v3_).length(0))]])
                    source15_ = D5_DC17(d_1652_v13_, True)
                    if source15_.is_DC16:
                        d_1653___mcc_h0_ = source15_.cf29
                        d_1654___mcc_h1_ = source15_.cf30
                        d_1655___mcc_h2_ = source15_.cf31
                        d_1656___mcc_h3_ = source15_.cf32
                        d_1657___mcc_h4_ = source15_.cf33
                        d_1658_cf33_ = d_1657___mcc_h4_
                        d_1659_cf32_ = d_1656___mcc_h3_
                        d_1660_cf31_ = d_1655___mcc_h2_
                        d_1661_cf30_ = d_1654___mcc_h1_
                        d_1662_cf29_ = d_1653___mcc_h0_
                        d_1663_v14_: _dafny.Array
                        def lambda77_(d_1664_i1_):
                            return (d_1664_i1_) - (self.f4)

                        init45_ = lambda77_
                        nw284_ = _dafny.Array(None, 9)
                        for i0_45_ in range(nw284_.length(0)):
                            nw284_[i0_45_] = init45_(i0_45_)
                        d_1663_v14_ = nw284_
                        index239_ = default__.safeIndex(180, (d_1663_v14_).length(0))
                        (d_1663_v14_)[index239_] = self.f4
                        index240_ = default__.safeIndex(180, (d_1663_v14_).length(0))
                        (d_1663_v14_)[index240_] = 900
                        d_1665_v15_: _dafny.Map
                        d_1665_v15_ = _dafny.Map({d_1662_cf29_: (d_1663_v14_)[default__.safeIndex(180, (d_1663_v14_).length(0))]})
                        d_1666_v16_: D19
                        d_1666_v16_ = D19_DC51(self.f4, d_1659_cf32_)
                        d_1667_v17_: D12
                        d_1667_v17_ = D12_DC38(d_1659_cf32_)
                        pat_let_tv47_ = d_1645_v7_
                        d_1668_v18_: _dafny.Array
                        nw285_ = _dafny.Array(None, 15)
                        nw285_[int(0)] = D12_DC38(((d_1665_v15_)[d_1639_v1_] if (d_1639_v1_) in (d_1665_v15_) else (d_1663_v14_)[default__.safeIndex(180, (d_1663_v14_).length(0))]))
                        nw285_[int(1)] = default__.fm38(d_1666_v16_, d_1645_v7_, d_1639_v1_, globalState)
                        nw285_[int(2)] = d_1667_v17_
                        nw285_[int(3)] = d_1667_v17_
                        nw285_[int(4)] = (d_1667_v17_ if d_1661_cf30_ else d_1667_v17_)
                        nw285_[int(5)] = d_1667_v17_
                        nw285_[int(6)] = d_1667_v17_
                        nw285_[int(7)] = d_1667_v17_
                        nw285_[int(8)] = d_1667_v17_
                        nw285_[int(9)] = d_1667_v17_
                        nw285_[int(10)] = d_1667_v17_
                        nw285_[int(11)] = d_1667_v17_
                        nw285_[int(12)] = default__.fm38(D19_DC51((0) - (d_1659_cf32_), len(d_1645_v7_)), d_1645_v7_, (d_1641_v3_)[default__.safeIndex(699, (d_1641_v3_).length(0))], globalState)
                        nw285_[int(13)] = default__.fm38(d_1666_v16_, default__.fm23(True, globalState), default__.fm0(self.f4, self.f4, globalState), globalState)
                        def iife173_(_pat_let46_0):
                            def iife174_(d_1669_dt__update__tmp_h0_):
                                def iife175_(_pat_let47_0):
                                    def iife176_(d_1670_dt__update_hcf68_h0_):
                                        return D12_DC38(d_1670_dt__update_hcf68_h0_)
                                    return iife176_(_pat_let47_0)
                                return iife175_(len(pat_let_tv47_))
                            return iife174_(_pat_let46_0)
                        nw285_[int(14)] = iife173_(d_1667_v17_)
                        d_1668_v18_ = nw285_
                        index241_ = default__.safeIndex(343, (d_1668_v18_).length(0))
                        (d_1668_v18_)[index241_] = d_1667_v17_
                        d_1671_v19_: _dafny.MultiSet
                        d_1671_v19_ = _dafny.MultiSet([(d_1663_v14_)[default__.safeIndex(180, (d_1663_v14_).length(0))]])
                        d_1672_v20_: _dafny.Array
                        d_1672_v20_ = d_1663_v14_
                        d_1673_v21_: _dafny.Seq
                        d_1673_v21_ = _dafny.SeqWithoutIsStrInference([d_1672_v20_])
                        d_1674_v22_: _dafny.Seq
                        d_1674_v22_ = _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([(d_1663_v14_)[default__.safeIndex(180, (d_1663_v14_).length(0))], self.f4])).intersection(d_1671_v19_), d_1671_v19_, _dafny.MultiSet([(d_1663_v14_)[default__.safeIndex(180, (d_1663_v14_).length(0))], len(d_1673_v21_), len(d_1644_v6_)])])
                        d_1675_v23_: D2
                        d_1675_v23_ = D2_DC6(-857, d_1658_cf33_)
                        d_1676_v24_: _dafny.Seq
                        d_1676_v24_ = _dafny.SeqWithoutIsStrInference([d_1661_cf30_, (d_1675_v23_).cf11, (d_1641_v3_)[default__.safeIndex(699, (d_1641_v3_).length(0))], d_1658_cf33_])
                        index242_ = default__.safeIndex(343, (d_1668_v18_).length(0))
                        rhs207_ = d_1667_v17_
                        rhs208_ = ((d_1676_v24_) + (d_1676_v24_)) <= (d_1676_v24_)
                        rhs209_ = self.f4
                        rhs210_ = _dafny.SeqWithoutIsStrInference([(d_1671_v19_).intersection(d_1671_v19_) for d_1677_i2_ in range(default__.abs(-862))])
                        lhs124_ = d_1668_v18_
                        lhs125_ = default__.safeIndex(343, (d_1668_v18_).length(0))
                        lhs126_ = self
                        lhs124_[lhs125_] = rhs207_
                        d_1658_cf33_ = rhs208_
                        lhs126_.f4 = rhs209_
                        d_1674_v22_ = rhs210_
                        d_1678_v25_: _dafny.Seq
                        d_1678_v25_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([d_1642_v4_, d_1642_v4_])).cardinality for d_1679_i3_ in range(default__.abs(425))]), d_1648_v10_])
                        d_1680_v26_: _dafny.Map
                        d_1680_v26_ = _dafny.Map({self.f4: d_1678_v25_})
                        d_1681_v27_: D20
                        d_1681_v27_ = D20_DC53(((d_1680_v26_)[self.f4] if (self.f4) in (d_1680_v26_) else d_1678_v25_))
                        d_1678_v25_ = ((d_1678_v25_) + ((d_1681_v27_).cf96)) + (d_1678_v25_)
                        index243_ = default__.safeIndex(180, (d_1663_v14_).length(0))
                        (d_1663_v14_)[index243_] = default__.safeDivisionInt(478, default__.safeDivisionInt((d_1663_v14_)[default__.safeIndex(180, (d_1663_v14_).length(0))], d_1659_cf32_))
                    elif source15_.is_DC17:
                        d_1682___mcc_h5_ = source15_.cf34
                        d_1683___mcc_h6_ = source15_.cf35
                        d_1684_cf35_ = d_1683___mcc_h6_
                        d_1685_cf34_ = d_1682___mcc_h5_
                        d_1639_v1_ = d_1684_cf35_
                        index244_ = default__.safeIndex(699, (d_1641_v3_).length(0))
                        def iife177_():
                            coll81_ = _dafny.Map()
                            compr_81_: int
                            for compr_81_ in _dafny.IntegerRange(621, 496):
                                d_1686_v28_: int = compr_81_
                                if ((621) <= (d_1686_v28_)) and ((d_1686_v28_) < (496)):
                                    coll81_[default__.safeModuloInt(d_1686_v28_, (0) - (self.f4))] = len(d_1649_v11_)
                            return _dafny.Map(coll81_)
                        rhs211_ = d_1684_cf35_
                        rhs212_ = d_1642_v4_
                        rhs213_ = not (d_1684_cf35_) or ((self.f4) > (self.f4))
                        rhs214_ = (d_1646_v8_ if (_dafny.MultiSet([(d_1641_v3_)[default__.safeIndex(699, (d_1641_v3_).length(0))], True])).ispropersubset(d_1652_v13_) else iife177_()
                        )
                        rhs215_ = (d_1648_v10_)[default__.safeIndex(self.f4, len(d_1648_v10_))]
                        lhs127_ = d_1641_v3_
                        lhs128_ = default__.safeIndex(699, (d_1641_v3_).length(0))
                        lhs129_ = self
                        lhs127_[lhs128_] = rhs211_
                        d_1642_v4_ = rhs212_
                        d_1684_cf35_ = rhs213_
                        d_1646_v8_ = rhs214_
                        lhs129_.f4 = rhs215_
                        d_1687_v29_: _dafny.Array
                        def lambda78_(d_1688_v8_):
                            def lambda79_(d_1689_i4_):
                                return (d_1689_i4_) - (len(d_1688_v8_))

                            return lambda79_

                        init46_ = lambda78_(d_1646_v8_)
                        nw286_ = _dafny.Array(None, 27)
                        for i0_46_ in range(nw286_.length(0)):
                            nw286_[i0_46_] = init46_(i0_46_)
                        d_1687_v29_ = nw286_
                        d_1687_v29_ = (d_1687_v29_ if d_1639_v1_ else d_1687_v29_)
                        (self).f4 = 141
                    elif source15_.is_DC18:
                        d_1690___mcc_h7_ = source15_.cf36
                        d_1691___mcc_h8_ = source15_.cf37
                        d_1692___mcc_h9_ = source15_.cf38
                        d_1693_cf38_ = d_1692___mcc_h9_
                        d_1694_cf37_ = d_1691___mcc_h8_
                        d_1695_cf36_ = d_1690___mcc_h7_
                        d_1644_v6_ = (d_1644_v6_).set(d_1695_cf36_.f18, (d_1641_v3_)[default__.safeIndex(699, (d_1641_v3_).length(0))])
                        d_1696_v30_: D3
                        d_1696_v30_ = D3_DC8(d_1642_v4_, (d_1638_v0_)[default__.safeIndex(908, (d_1638_v0_).length(0))])
                        d_1697_v31_: _dafny.Map
                        d_1697_v31_ = _dafny.Map({_dafny.Map({(0) - (self.f4): self.f4}): _dafny.SeqWithoutIsStrInference([d_1695_cf36_.f18 for d_1698_i5_ in range(default__.abs(423))])})
                        d_1699_v32_: _dafny.Map
                        d_1699_v32_ = _dafny.Map({d_1696_v30_: ((d_1697_v31_)[d_1646_v8_] if (d_1646_v8_) in (d_1697_v31_) else d_1648_v10_)})
                        d_1699_v32_ = (d_1699_v32_).set(d_1696_v30_, (d_1648_v10_) + (d_1648_v10_))
                        d_1700_v33_: _dafny.Map
                        d_1700_v33_ = _dafny.Map({d_1695_cf36_.f18: (d_1648_v10_) + (d_1648_v10_)})
                        d_1700_v33_ = _dafny.Map({(d_1695_cf36_.f18) - (self.f4): d_1648_v10_})
                        d_1701_v34_: _dafny.MultiSet
                        d_1701_v34_ = _dafny.MultiSet([len(d_1646_v8_)])
                        d_1702_v35_: _dafny.Seq
                        d_1702_v35_ = _dafny.SeqWithoutIsStrInference([d_1701_v34_, d_1701_v34_])
                        d_1703_v36_: _dafny.Array
                        nw287_ = _dafny.Array(None, 8)
                        nw287_[int(0)] = d_1701_v34_
                        nw287_[int(1)] = (d_1702_v35_)[default__.safeIndex(d_1695_cf36_.f18, len(d_1702_v35_))]
                        nw287_[int(2)] = d_1701_v34_
                        nw287_[int(3)] = d_1701_v34_
                        nw287_[int(4)] = d_1701_v34_
                        nw287_[int(5)] = d_1701_v34_
                        nw287_[int(6)] = d_1701_v34_
                        nw287_[int(7)] = d_1701_v34_
                        d_1703_v36_ = nw287_
                        d_1704_v37_: D5
                        d_1704_v37_ = D5_DC15(d_1703_v36_)
                        d_1705_v38_: D5
                        d_1705_v38_ = D5_DC19(d_1704_v37_)
                        d_1706_v39_: D5
                        d_1706_v39_ = D5_DC19(d_1704_v37_)
                        d_1707_v40_: _dafny.Seq
                        d_1707_v40_ = _dafny.SeqWithoutIsStrInference([d_1706_v39_])
                        d_1707_v40_ = (_dafny.SeqWithoutIsStrInference([d_1706_v39_])).set(default__.safeIndex(d_1695_cf36_.f18, len(_dafny.SeqWithoutIsStrInference([d_1706_v39_]))), d_1706_v39_)
                    elif source15_.is_DC15:
                        d_1708___mcc_h10_ = source15_.cf28
                        d_1709_cf28_ = d_1708___mcc_h10_
                        d_1710_v41_: _dafny.Array
                        nw288_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 4)
                        d_1710_v41_ = nw288_
                        index245_ = default__.safeIndex(185, (d_1710_v41_).length(0))
                        (d_1710_v41_)[index245_] = default__.fm23(False, globalState)
                        d_1711_v42_: _dafny.Map
                        d_1711_v42_ = _dafny.Map({d_1639_v1_: len(d_1645_v7_)})
                        d_1712_v43_: _dafny.Array
                        nw289_ = _dafny.Array(None, 24)
                        d_1712_v43_ = nw289_
                        d_1713_v44_: _dafny.Map
                        d_1713_v44_ = _dafny.Map({len((d_1711_v42_).set(d_1639_v1_, self.f4)): d_1712_v43_})
                        d_1714_v45_: _dafny.Map
                        d_1714_v45_ = (d_1713_v44_).set(self.f4, d_1712_v43_)
                        d_1715_v46_: _dafny.Seq
                        d_1715_v46_ = _dafny.SeqWithoutIsStrInference([d_1713_v44_])
                        d_1716_v47_: _dafny.Array
                        d_1716_v47_ = d_1712_v43_
                        d_1717_v48_: _dafny.Seq
                        d_1717_v48_ = _dafny.SeqWithoutIsStrInference([(d_1716_v47_)])
                        d_1718_v49_: _dafny.Array
                        nw290_ = _dafny.Array(None, 16)
                        nw290_[int(0)] = d_1714_v45_
                        nw290_[int(1)] = d_1714_v45_
                        nw290_[int(2)] = (d_1715_v46_)[default__.safeIndex(len(d_1645_v7_), len(d_1715_v46_))]
                        nw290_[int(3)] = d_1714_v45_
                        nw290_[int(4)] = d_1714_v45_
                        nw290_[int(5)] = d_1714_v45_
                        nw290_[int(6)] = d_1713_v44_
                        nw290_[int(7)] = d_1714_v45_
                        nw290_[int(8)] = _dafny.Map({self.f4: (d_1717_v48_)[default__.safeIndex(self.f4, len(d_1717_v48_))]})
                        nw290_[int(9)] = d_1714_v45_
                        nw290_[int(10)] = d_1714_v45_
                        nw290_[int(11)] = d_1713_v44_
                        nw290_[int(12)] = d_1714_v45_
                        nw290_[int(13)] = d_1714_v45_
                        nw290_[int(14)] = d_1714_v45_
                        nw290_[int(15)] = d_1714_v45_
                        d_1718_v49_ = nw290_
                        index246_ = default__.safeIndex(185, (d_1710_v41_).length(0))
                        rhs216_ = d_1645_v7_
                        rhs217_ = (-379) - (265)
                        rhs218_ = d_1718_v49_
                        lhs130_ = d_1710_v41_
                        lhs131_ = default__.safeIndex(185, (d_1710_v41_).length(0))
                        lhs132_ = self
                        lhs130_[lhs131_] = rhs216_
                        lhs132_.f4 = rhs217_
                        d_1718_v49_ = rhs218_
                        d_1719_v50_: _dafny.Seq
                        d_1719_v50_ = _dafny.SeqWithoutIsStrInference([d_1645_v7_])
                        (self).f4 = len((d_1719_v50_) + (d_1719_v50_))
                        index247_ = default__.safeIndex(699, (d_1641_v3_).length(0))
                        (d_1641_v3_)[index247_] = (d_1641_v3_)[default__.safeIndex(699, (d_1641_v3_).length(0))]
                        d_1720_v51_: C3
                        nw291_ = C3()
                        nw291_.ctor__(self.f4, d_1645_v7_, (0) - ((0) - ((default__.fm1(globalState)) * ((0) - (self.f4)))))
                        d_1720_v51_ = nw291_
                    elif True:
                        d_1721___mcc_h11_ = source15_.cf39
                        d_1722_cf39_ = d_1721___mcc_h11_
                        r0 = d_1642_v4_
                        index248_ = default__.safeIndex(699, (d_1641_v3_).length(0))
                        (d_1641_v3_)[index248_] = d_1639_v1_
                        d_1723_v52_: _dafny.Array
                        nw292_ = _dafny.Array(_dafny.MultiSet({}), 15)
                        d_1723_v52_ = nw292_
                        d_1724_v53_: _dafny.MultiSet
                        d_1724_v53_ = _dafny.MultiSet([d_1645_v7_, d_1645_v7_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yoi")), d_1645_v7_])
                        index249_ = default__.safeIndex(418, (d_1723_v52_).length(0))
                        (d_1723_v52_)[index249_] = d_1724_v53_
                        index250_ = default__.safeIndex(699, (d_1641_v3_).length(0))
                        index251_ = default__.safeIndex(418, (d_1723_v52_).length(0))
                        rhs219_ = (d_1641_v3_)[default__.safeIndex(699, (d_1641_v3_).length(0))]
                        rhs220_ = 531
                        rhs221_ = ((d_1724_v53_) - (default__.fm39(self.f4, globalState))) | ((d_1724_v53_).set(_dafny.SeqWithoutIsStrInference([d_1642_v4_ for d_1725_i6_ in range(default__.abs(-77))]), default__.abs(self.f4)))
                        lhs133_ = d_1641_v3_
                        lhs134_ = default__.safeIndex(699, (d_1641_v3_).length(0))
                        lhs135_ = self
                        lhs136_ = d_1723_v52_
                        lhs137_ = default__.safeIndex(418, (d_1723_v52_).length(0))
                        lhs133_[lhs134_] = rhs219_
                        lhs135_.f4 = rhs220_
                        lhs136_[lhs137_] = rhs221_
                        d_1645_v7_ = (d_1645_v7_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mcpxtxnu")))
                    d_1643_v5_ = (d_1643_v5_) - ((_dafny.Set({(d_1641_v3_)[default__.safeIndex(699, (d_1641_v3_).length(0))]})) - (d_1643_v5_))
                    if d_1639_v1_:
                        d_1726_v54_: _dafny.Array
                        nw293_ = _dafny.Array(int(0), 13)
                        d_1726_v54_ = nw293_
                        index252_ = default__.safeIndex(834, (d_1726_v54_).length(0))
                        (d_1726_v54_)[index252_] = self.f4
                        index253_ = default__.safeIndex(834, (d_1726_v54_).length(0))
                        (d_1726_v54_)[index253_] = len(d_1643_v5_)
                        d_1727_v56_: _dafny.Map
                        def iife178_():
                            coll82_ = _dafny.Map()
                            compr_82_: str
                            for compr_82_ in (d_1645_v7_).Elements:
                                d_1728_v55_: str = compr_82_
                                if (d_1728_v55_) in (d_1645_v7_):
                                    coll82_[d_1728_v55_] = len(d_1648_v10_)
                            return _dafny.Map(coll82_)
                        d_1727_v56_ = _dafny.Map({iife178_()
                        : d_1639_v1_})
                        d_1729_v58_: _dafny.Map
                        d_1729_v58_ = _dafny.Map({d_1642_v4_: 108})
                        d_1730_v59_: _dafny.Set
                        d_1730_v59_ = _dafny.Set({d_1729_v58_})
                        d_1731_v60_: _dafny.Seq
                        def iife179_():
                            coll83_ = _dafny.Set()
                            compr_83_: _dafny.Map
                            for compr_83_ in (d_1727_v56_).keys.Elements:
                                d_1732_v57_: _dafny.Map = compr_83_
                                if (d_1732_v57_) in (d_1727_v56_):
                                    coll83_ = coll83_.union(_dafny.Set([d_1732_v57_]))
                            return _dafny.Set(coll83_)
                        d_1731_v60_ = _dafny.SeqWithoutIsStrInference([(d_1641_v3_)[default__.safeIndex(699, (d_1641_v3_).length(0))], (d_1641_v3_)[default__.safeIndex(699, (d_1641_v3_).length(0))], ((D22_DC58(d_1730_v59_)).cf106).issubset(iife179_()
                        )])
                        rhs222_ = (self).fm5(d_1645_v7_, (d_1641_v3_)[default__.safeIndex(699, (d_1641_v3_).length(0))], globalState)
                        rhs223_ = (d_1638_v0_)[default__.safeIndex(908, (d_1638_v0_).length(0))]
                        rhs224_ = ((_dafny.SeqWithoutIsStrInference([(d_1641_v3_)[default__.safeIndex(699, (d_1641_v3_).length(0))], d_1639_v1_])) + (_dafny.SeqWithoutIsStrInference([d_1639_v1_, d_1639_v1_]))) + (d_1731_v60_)
                        rhs225_ = False
                        rhs226_ = d_1639_v1_
                        d_1639_v1_ = rhs222_
                        d_1641_v3_ = rhs223_
                        d_1731_v60_ = rhs224_
                        d_1639_v1_ = rhs225_
                        d_1639_v1_ = rhs226_
                        d_1641_v3_ = (d_1638_v0_)[default__.safeIndex(908, (d_1638_v0_).length(0))]
                        d_1733_v61_: _dafny.Set
                        d_1733_v61_ = _dafny.Set({(d_1638_v0_)[default__.safeIndex(908, (d_1638_v0_).length(0))]})
                        d_1734_v62_: _dafny.Map
                        d_1734_v62_ = _dafny.Map({d_1726_v54_: d_1733_v61_})
                        d_1734_v62_ = (d_1734_v62_).set(d_1726_v54_, d_1733_v61_)
                        index254_ = default__.safeIndex(834, (d_1726_v54_).length(0))
                        (d_1726_v54_)[index254_] = (0) - ((403) + (((0) - (self.f4)) * ((d_1726_v54_)[default__.safeIndex(834, (d_1726_v54_).length(0))])))
                    elif True:
                        rhs227_ = (d_1641_v3_)[default__.safeIndex(699, (d_1641_v3_).length(0))]
                        rhs228_ = d_1649_v11_
                        rhs229_ = d_1639_v1_
                        rhs230_ = d_1639_v1_
                        rhs231_ = default__.safeModuloInt(self.f4, self.f4)
                        lhs138_ = self
                        d_1639_v1_ = rhs227_
                        d_1649_v11_ = rhs228_
                        d_1639_v1_ = rhs229_
                        d_1639_v1_ = rhs230_
                        lhs138_.f4 = rhs231_
                        d_1644_v6_ = (d_1644_v6_).set(self.f4, (d_1641_v3_)[default__.safeIndex(699, (d_1641_v3_).length(0))])
                        (self).f4 = (0) - (self.f4)
                        d_1735_v63_: _dafny.Seq
                        d_1735_v63_ = _dafny.SeqWithoutIsStrInference([d_1641_v3_])
                        d_1735_v63_ = d_1735_v63_
                        index255_ = default__.safeIndex(699, (d_1641_v3_).length(0))
                        (d_1641_v3_)[index255_] = ((D12_DC37(d_1639_v1_, (d_1641_v3_)[default__.safeIndex(699, (d_1641_v3_).length(0))], d_1639_v1_)).cf65) and ((d_1641_v3_)[default__.safeIndex(699, (d_1641_v3_).length(0))])
                    d_1736_v64_: _dafny.Map
                    d_1736_v64_ = _dafny.Map({d_1639_v1_: _dafny.Map({self.f4: (d_1641_v3_)[default__.safeIndex(699, (d_1641_v3_).length(0))]})})
                    d_1639_v1_ = ((d_1641_v3_)[default__.safeIndex(699, (d_1641_v3_).length(0))]) not in ((d_1736_v64_) | (d_1736_v64_))
                    pass
            pass
        (self).f4 = -137
        d_1737_v65_: _dafny.MultiSet
        d_1737_v65_ = _dafny.MultiSet([self.f4])
        hi13_ = (self.f4) - (self.f4)
        for d_1738_i7_ in range(((d_1737_v65_).intersection(_dafny.MultiSet(d_1648_v10_))).cardinality, hi13_):
            d_1739_v66_: C2
            nw294_ = C2()
            nw294_.ctor__(d_1639_v1_, (0) - (self.f4))
            d_1739_v66_ = nw294_
            (self).f4 = (636) + (default__.fm1(globalState))
            (self).f4 = default__.safeModuloInt(default__.safeModuloInt((d_1739_v66_).f14, self.f4), self.f4)
            d_1740_v67_: C0
            nw295_ = C0()
            nw295_.ctor__(d_1738_i7_)
            d_1740_v67_ = nw295_
        hi14_ = -516
        for d_1741_i8_ in range((self.f4) * (len(d_1646_v8_)), hi14_):
            d_1742_v68_: C2
            nw296_ = C2()
            nw296_.ctor__((d_1641_v3_)[default__.safeIndex(699, (d_1641_v3_).length(0))], self.f4)
            d_1742_v68_ = nw296_
            d_1743_v72_: _dafny.Array
            nw297_ = _dafny.Array(None, 23)
            nw297_[int(0)] = _dafny.Map({613: True})
            nw297_[int(1)] = d_1644_v6_
            nw297_[int(2)] = default__.fm37(d_1639_v1_, d_1742_v68_.f13, 787, globalState)
            nw297_[int(3)] = d_1644_v6_
            nw297_[int(4)] = d_1644_v6_
            nw297_[int(5)] = d_1644_v6_
            nw297_[int(6)] = d_1644_v6_
            nw297_[int(7)] = d_1644_v6_
            nw297_[int(8)] = default__.fm37((d_1641_v3_)[default__.safeIndex(699, (d_1641_v3_).length(0))], (d_1641_v3_)[default__.safeIndex(699, (d_1641_v3_).length(0))], (d_1742_v68_).f14, globalState)
            nw297_[int(9)] = d_1644_v6_
            nw297_[int(10)] = d_1644_v6_
            nw297_[int(11)] = d_1644_v6_
            nw297_[int(12)] = d_1644_v6_
            nw297_[int(13)] = d_1644_v6_
            nw297_[int(14)] = d_1644_v6_
            nw297_[int(15)] = d_1644_v6_
            def iife180_():
                coll84_ = _dafny.Map()
                compr_84_: int
                for compr_84_ in _dafny.IntegerRange(119, 87):
                    d_1744_v69_: int = compr_84_
                    if ((119) <= (d_1744_v69_)) and ((d_1744_v69_) < (87)):
                        coll84_[default__.safeDivisionInt(d_1744_v69_, self.f4)] = True
                return _dafny.Map(coll84_)
            nw297_[int(16)] = iife180_()
            
            def iife181_():
                coll85_ = _dafny.Map()
                compr_85_: int
                for compr_85_ in _dafny.IntegerRange(356, 912):
                    d_1745_v70_: int = compr_85_
                    if ((356) <= (d_1745_v70_)) and ((d_1745_v70_) < (912)):
                        coll85_[default__.safeModuloInt(d_1745_v70_, len(_dafny.SeqWithoutIsStrInference([d_1642_v4_ for d_1746_i9_ in range(default__.abs(349))])))] = d_1742_v68_.f13
                return _dafny.Map(coll85_)
            nw297_[int(17)] = iife181_()
            
            nw297_[int(18)] = d_1644_v6_
            def iife182_():
                coll86_ = _dafny.Map()
                compr_86_: int
                for compr_86_ in (d_1648_v10_).Elements:
                    d_1747_v71_: int = compr_86_
                    if (d_1747_v71_) in (d_1648_v10_):
                        coll86_[default__.safeModuloInt(d_1747_v71_, 410)] = (d_1641_v3_)[default__.safeIndex(699, (d_1641_v3_).length(0))]
                return _dafny.Map(coll86_)
            nw297_[int(19)] = iife182_()
            
            nw297_[int(20)] = d_1644_v6_
            nw297_[int(21)] = d_1644_v6_
            nw297_[int(22)] = d_1644_v6_
            d_1743_v72_ = nw297_
            d_1748_v73_: _dafny.Map
            d_1748_v73_ = _dafny.Map({d_1743_v72_: (d_1641_v3_)[default__.safeIndex(699, (d_1641_v3_).length(0))]})
            index256_ = default__.safeIndex(699, (d_1641_v3_).length(0))
            (d_1641_v3_)[index256_] = ((d_1748_v73_)[d_1743_v72_] if (d_1743_v72_) in (d_1748_v73_) else not (False) or (d_1639_v1_))
            d_1749_v74_: C1
            nw298_ = C1()
            nw298_.ctor__(self.f4, (d_1742_v68_).f14, 865, d_1639_v1_)
            d_1749_v74_ = nw298_
            d_1750_v75_: _dafny.Array
            def lambda80_(d_1751_i10_):
                return default__.safeModuloInt(d_1751_i10_, self.f4)

            init47_ = lambda80_
            nw299_ = _dafny.Array(None, 7)
            for i0_47_ in range(nw299_.length(0)):
                nw299_[i0_47_] = init47_(i0_47_)
            d_1750_v75_ = nw299_
            d_1752_v76_: _dafny.Map
            d_1752_v76_ = _dafny.Map({_dafny.Set({(d_1641_v3_)[default__.safeIndex(699, (d_1641_v3_).length(0))], (d_1641_v3_)[default__.safeIndex(699, (d_1641_v3_).length(0))], d_1639_v1_, d_1639_v1_, True}): d_1750_v75_})
            d_1752_v76_ = d_1752_v76_
        d_1753_v77_: _dafny.Map
        d_1753_v77_ = _dafny.Map({(default__.fm2(d_1642_v4_, self.f4, 135, globalState)) < (_dafny.SeqWithoutIsStrInference([(d_1641_v3_)[default__.safeIndex(699, (d_1641_v3_).length(0))]])): self.f4})
        d_1753_v77_ = _dafny.Map({d_1639_v1_: ((0) - (self.f4)) + (self.f4)})
        r0 = d_1642_v4_
        return r0


class C5(T0):
    def  __init__(self):
        self._f4: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    @property
    def f4(self):
        return self._f4
    @f4.setter
    def f4(self, value):
        self._f4 = value
    def ctor__(self, f4):
        (self).f4 = f4

    def fm5(self, p0, p1, globalState):
        return (D0_DC1(360, len(_dafny.SeqWithoutIsStrInference([self.f4, 122])), False)).cf3

    def fm8(self, globalState):
        return ((_dafny.MultiSet([False, not(False), not(True), True])).intersection(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True, True, False, True])))).ispropersubset((_dafny.MultiSet([True, True])).intersection(_dafny.MultiSet([False])))

    def fm9(self, p0, globalState):
        return D1_DC3((_dafny.Map({True: self.f4})) | (_dafny.Map({False: self.f4})))

    def m3(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        if p3:
            r0 = (not(p3)) == (False)
            d_1754_v0_: _dafny.Map
            d_1754_v0_ = _dafny.Map({p2: p1})
            d_1754_v0_ = (d_1754_v0_).set(p3, p1)
            d_1755_v1_: _dafny.Set
            d_1755_v1_ = _dafny.Set({True})
            if (d_1755_v1_).ispropersubset(d_1755_v1_):
                d_1756_v2_: _dafny.Seq
                d_1756_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tpbmnsdj"))
                d_1757_v3_: _dafny.Seq
                d_1757_v3_ = _dafny.SeqWithoutIsStrInference([p2])
                d_1758_v4_: _dafny.MultiSet
                d_1758_v4_ = _dafny.MultiSet([p0, p0])
                d_1759_v5_: _dafny.Map
                d_1759_v5_ = _dafny.Map({len(d_1757_v3_): d_1758_v4_})
                d_1760_v6_: D2
                d_1760_v6_ = D2_DC5(d_1758_v4_)
                d_1761_v7_: D0
                d_1761_v7_ = D0_DC1((0) - (len(d_1756_v2_)), (((d_1759_v5_)[p1] if (p1) in (d_1759_v5_) else (d_1760_v6_).cf9)).cardinality, p2)
                d_1762_v8_: _dafny.MultiSet
                d_1762_v8_ = _dafny.MultiSet([self.f4])
                d_1763_v9_: _dafny.Seq
                d_1763_v9_ = _dafny.SeqWithoutIsStrInference([d_1762_v8_, _dafny.MultiSet([p1, p1])])
                d_1764_v10_: _dafny.Array
                def lambda81_(d_1765_i0_):
                    return default__.safeModuloInt(d_1765_i0_, self.f4)

                init48_ = lambda81_
                nw300_ = _dafny.Array(None, 16)
                for i0_48_ in range(nw300_.length(0)):
                    nw300_[i0_48_] = init48_(i0_48_)
                d_1764_v10_ = nw300_
                d_1766_v11_: _dafny.Map
                d_1766_v11_ = _dafny.Map({d_1761_v7_: default__.safeDivisionInt(p1, ((d_1763_v9_)[default__.safeIndex(len(_dafny.Set({d_1764_v10_})), len(d_1763_v9_))]).cardinality)})
                d_1766_v11_ = (d_1766_v11_).set(d_1761_v7_, self.f4)
                d_1767_v12_: C2
                nw301_ = C2()
                nw301_.ctor__(True, self.f4)
                d_1767_v12_ = nw301_
                d_1768_v13_: _dafny.Map
                d_1768_v13_ = _dafny.Map({self.f4: p2})
                d_1769_v15_: _dafny.Array
                nw302_ = _dafny.Array(_dafny.CodePoint('D'), 13)
                d_1769_v15_ = nw302_
                d_1770_v16_: _dafny.Map
                d_1770_v16_ = _dafny.Map({d_1769_v15_: p3})
                d_1771_v17_: _dafny.Seq
                d_1771_v17_ = _dafny.SeqWithoutIsStrInference([len(d_1770_v16_)])
                d_1772_v18_: _dafny.Map
                def iife183_():
                    coll87_ = _dafny.Map()
                    compr_87_: int
                    for compr_87_ in (d_1771_v17_).Elements:
                        d_1773_v14_: int = compr_87_
                        if (d_1773_v14_) in (d_1771_v17_):
                            coll87_[(d_1773_v14_) + (744)] = p0
                    return _dafny.Map(coll87_)
                d_1772_v18_ = _dafny.Map({p1: iife183_()
                })
                d_1774_v19_: _dafny.Array
                nw303_ = _dafny.Array(None, 5)
                nw303_[int(0)] = (_dafny.Map({p1: True})) | (d_1768_v13_)
                nw303_[int(1)] = d_1768_v13_
                nw303_[int(2)] = d_1768_v13_
                nw303_[int(3)] = d_1768_v13_
                nw303_[int(4)] = ((d_1772_v18_)[p1] if (p1) in (d_1772_v18_) else d_1768_v13_)
                d_1774_v19_ = nw303_
                index257_ = default__.safeIndex(191, (d_1774_v19_).length(0))
                (d_1774_v19_)[index257_] = d_1768_v13_
                index258_ = default__.safeIndex(191, (d_1774_v19_).length(0))
                (d_1774_v19_)[index258_] = (d_1768_v13_).set(default__.safeModuloInt(p1, -501), (d_1767_v12_.f13 if p3 else p0))
                d_1775_v20_: _dafny.Array
                def lambda82_(d_1776_v2_, d_1777_v12_):
                    def lambda83_(d_1778_i1_):
                        return ((0) - (len(d_1776_v2_))) >= ((d_1777_v12_).f14)

                    return lambda83_

                init49_ = lambda82_(d_1756_v2_, d_1767_v12_)
                nw304_ = _dafny.Array(None, 5)
                for i0_49_ in range(nw304_.length(0)):
                    nw304_[i0_49_] = init49_(i0_49_)
                d_1775_v20_ = nw304_
                index259_ = default__.safeIndex(758, (d_1775_v20_).length(0))
                (d_1775_v20_)[index259_] = p0
                index260_ = default__.safeIndex(758, (d_1775_v20_).length(0))
                (d_1775_v20_)[index260_] = d_1767_v12_.f13
                d_1775_v20_ = (d_1775_v20_ if p3 else d_1775_v20_)
            elif True:
                d_1779_v21_: _dafny.Array
                nw305_ = _dafny.Array(False, 29)
                d_1779_v21_ = nw305_
                d_1780_v22_: _dafny.Map
                d_1780_v22_ = _dafny.Map({p3: d_1754_v0_})
                d_1781_v23_: _dafny.Array
                nw306_ = _dafny.Array(None, 29)
                nw306_[int(0)] = d_1754_v0_
                nw306_[int(1)] = d_1754_v0_
                nw306_[int(2)] = d_1754_v0_
                nw306_[int(3)] = d_1754_v0_
                nw306_[int(4)] = d_1754_v0_
                nw306_[int(5)] = d_1754_v0_
                nw306_[int(6)] = (d_1754_v0_) | (_dafny.Map({p0: self.f4}))
                nw306_[int(7)] = d_1754_v0_
                nw306_[int(8)] = _dafny.Map({p2: len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({p0})) for d_1782_i2_ in range(default__.abs(-128))]))})
                nw306_[int(9)] = d_1754_v0_
                nw306_[int(10)] = d_1754_v0_
                nw306_[int(11)] = d_1754_v0_
                nw306_[int(12)] = d_1754_v0_
                nw306_[int(13)] = (_dafny.Map({p3: (0) - ((_dafny.MultiSet([(0) - (self.f4)])).cardinality)})) | (d_1754_v0_)
                nw306_[int(14)] = (d_1754_v0_).set(not(p0), p1)
                nw306_[int(15)] = d_1754_v0_
                nw306_[int(16)] = ((d_1780_v22_)[p0] if (p0) in (d_1780_v22_) else d_1754_v0_)
                nw306_[int(17)] = (d_1754_v0_).set(not(p0), 245)
                nw306_[int(18)] = (d_1754_v0_) | (d_1754_v0_)
                nw306_[int(19)] = d_1754_v0_
                nw306_[int(20)] = ((d_1754_v0_).set(True, self.f4)) | (d_1754_v0_)
                nw306_[int(21)] = d_1754_v0_
                nw306_[int(22)] = (d_1754_v0_) | (_dafny.Map({p0: self.f4}))
                nw306_[int(23)] = d_1754_v0_
                nw306_[int(24)] = d_1754_v0_
                nw306_[int(25)] = d_1754_v0_
                nw306_[int(26)] = d_1754_v0_
                nw306_[int(27)] = (d_1754_v0_ if p3 else d_1754_v0_)
                nw306_[int(28)] = _dafny.Map({p2: p1})
                d_1781_v23_ = nw306_
                rhs232_ = (self.f4) + (p1)
                rhs233_ = d_1779_v21_
                rhs234_ = d_1781_v23_
                lhs139_ = self
                lhs139_.f4 = rhs232_
                d_1779_v21_ = rhs233_
                d_1781_v23_ = rhs234_
                d_1783_v24_: _dafny.Seq
                d_1783_v24_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lwnavcm"))
                d_1783_v24_ = d_1783_v24_
                d_1784_v25_: _dafny.Array
                nw307_ = _dafny.Array(int(0), 3)
                d_1784_v25_ = nw307_
                d_1785_v26_: _dafny.MultiSet
                d_1785_v26_ = _dafny.MultiSet([d_1784_v25_, d_1784_v25_, (d_1784_v25_ if p3 else d_1784_v25_)])
                d_1785_v26_ = d_1785_v26_
                d_1786_v27_: _dafny.Map
                d_1786_v27_ = _dafny.Map({p3: not(p3)})
                d_1787_v28_: _dafny.Seq
                d_1787_v28_ = _dafny.SeqWithoutIsStrInference([len(d_1786_v27_)])
                d_1788_v29_: _dafny.Seq
                d_1788_v29_ = _dafny.SeqWithoutIsStrInference([p3])
                d_1789_v30_: _dafny.MultiSet
                d_1789_v30_ = _dafny.MultiSet([p0])
                d_1790_v31_: _dafny.Set
                d_1790_v31_ = _dafny.Set({d_1789_v30_})
                d_1791_v32_: _dafny.Seq
                d_1791_v32_ = _dafny.SeqWithoutIsStrInference([d_1787_v28_, default__.fm21(len(d_1788_v29_), p3, False, d_1790_v31_, globalState), d_1787_v28_, d_1787_v28_])
                r0 = not((d_1791_v32_) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.Map({478: _dafny.CodePoint('e')}))]) for d_1792_i3_ in range(default__.abs(483))])))
                (self).f4 = (0) - (self.f4)
            d_1793_v33_: _dafny.MultiSet
            d_1793_v33_ = _dafny.MultiSet([p0])
            d_1794_v34_: _dafny.Seq
            d_1794_v34_ = _dafny.SeqWithoutIsStrInference([p2, True, False])
            d_1795_v35_: _dafny.Set
            d_1795_v35_ = _dafny.Set({default__.safeDivisionInt(((d_1793_v33_)[p3] if (p3) in (d_1793_v33_) else len(d_1794_v34_)), p1)})
            d_1796_v37_: _dafny.Seq
            d_1796_v37_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vkng"))
            d_1797_v38_: str
            d_1797_v38_ = _dafny.CodePoint('f')
            d_1798_v39_: _dafny.Map
            d_1798_v39_ = _dafny.Map({p1: len(default__.fm14(len(d_1796_v37_), self.f4, False, d_1797_v38_, globalState))})
            def iife184_():
                coll88_ = _dafny.Set()
                compr_88_: int
                for compr_88_ in _dafny.IntegerRange(492, 847):
                    d_1799_v36_: int = compr_88_
                    if ((492) <= (d_1799_v36_)) and ((d_1799_v36_) < (847)):
                        coll88_ = coll88_.union(_dafny.Set([default__.safeDivisionInt(d_1799_v36_, p1)]))
                return _dafny.Set(coll88_)
            rhs235_ = iife184_()

            rhs236_ = (self.f4) * ((len(d_1754_v0_)) - (len(d_1798_v39_)))
            lhs140_ = self
            d_1795_v35_ = rhs235_
            lhs140_.f4 = rhs236_
            d_1800_v40_: _dafny.MultiSet
            d_1800_v40_ = _dafny.MultiSet([p1])
            d_1801_v41_: _dafny.MultiSet
            d_1801_v41_ = d_1800_v40_
            r0 = ((d_1801_v41_)).ispropersubset(d_1800_v40_)
        elif True:
            if p3:
                (self).f4 = (p1) + (self.f4)
                d_1802_v42_: C0
                nw308_ = C0()
                nw308_.ctor__((0) - (default__.safeDivisionInt(p1, p1)))
                d_1802_v42_ = nw308_
                d_1803_v43_: _dafny.Seq
                d_1803_v43_ = _dafny.SeqWithoutIsStrInference([self.f4, 695, self.f4, d_1802_v42_.f18])
                d_1804_v44_: D19
                d_1804_v44_ = D19_DC51(self.f4, (d_1803_v43_)[default__.safeIndex(d_1802_v42_.f18, len(d_1803_v43_))])
                d_1805_v45_: str
                d_1805_v45_ = _dafny.CodePoint('a')
                d_1806_v46_: _dafny.Map
                d_1806_v46_ = _dafny.Map({d_1805_v45_: (d_1803_v43_) != (_dafny.SeqWithoutIsStrInference([d_1802_v42_.f18]))})
                d_1807_v47_: _dafny.Array
                def lambda84_(d_1808_i4_):
                    return (d_1808_i4_) + (self.f4)

                init50_ = lambda84_
                nw309_ = _dafny.Array(None, 13)
                for i0_50_ in range(nw309_.length(0)):
                    nw309_[i0_50_] = init50_(i0_50_)
                d_1807_v47_ = nw309_
                d_1809_v48_: _dafny.MultiSet
                d_1809_v48_ = _dafny.MultiSet([p2, True])
                rhs237_ = p3
                rhs238_ = d_1804_v44_
                rhs239_ = ((d_1806_v46_) | (d_1806_v46_)).set(d_1805_v45_, (p1) == ((d_1809_v48_).cardinality))
                rhs240_ = d_1807_v47_
                r0 = rhs237_
                d_1804_v44_ = rhs238_
                d_1806_v46_ = rhs239_
                d_1807_v47_ = rhs240_
                (self).f4 = d_1802_v42_.f18
                (d_1802_v42_).f18 = default__.safeDivisionInt(len(default__.fm40(globalState)), len(d_1803_v43_))
            elif True:
                (self).f4 = default__.safeDivisionInt(p1, (0) - (self.f4))
                d_1810_v49_: _dafny.Array
                nw310_ = _dafny.Array(int(0), 16)
                d_1810_v49_ = nw310_
                index261_ = default__.safeIndex(646, (d_1810_v49_).length(0))
                (d_1810_v49_)[index261_] = self.f4
                d_1811_v50_: _dafny.MultiSet
                d_1811_v50_ = _dafny.MultiSet([self.f4])
                index262_ = default__.safeIndex(646, (d_1810_v49_).length(0))
                rhs241_ = (0) - (self.f4)
                rhs242_ = (d_1811_v50_).issubset(d_1811_v50_)
                rhs243_ = p0
                lhs141_ = d_1810_v49_
                lhs142_ = default__.safeIndex(646, (d_1810_v49_).length(0))
                lhs141_[lhs142_] = rhs241_
                r0 = rhs242_
                r0 = rhs243_
                r0 = (p3 if p2 else p0)
                index263_ = default__.safeIndex(646, (d_1810_v49_).length(0))
                (d_1810_v49_)[index263_] = p1
                d_1812_v51_: _dafny.Seq
                d_1812_v51_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qcx"))
                d_1813_v52_: _dafny.Seq
                d_1813_v52_ = _dafny.SeqWithoutIsStrInference([d_1812_v51_])
                d_1814_v53_: _dafny.Set
                d_1814_v53_ = _dafny.Set({3})
                r0 = (d_1814_v53_).issubset(default__.fm16(not(p3), d_1813_v52_, p3, p1, globalState))
            d_1815_v55_: _dafny.Array
            def lambda85_(d_1816_p2_, d_1817_p3_, d_1818_p1_):
                def lambda86_(d_1819_i5_):
                    def iife185_():
                        coll89_ = _dafny.Map()
                        compr_89_: int
                        for compr_89_ in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([self.f4]))).Elements:
                            d_1820_v54_: int = compr_89_
                            if (d_1820_v54_) in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([self.f4]))):
                                coll89_[default__.safeModuloInt(d_1820_v54_, (_dafny.MultiSet([self.f4, self.f4, d_1818_p1_])).cardinality)] = True
                        return _dafny.Map(coll89_)
                    return (_dafny.Map({self.f4: d_1817_p3_}) if d_1816_p2_ else iife185_()
                    )

                return lambda86_

            init51_ = lambda85_(p2, p3, p1)
            nw311_ = _dafny.Array(None, 1)
            for i0_51_ in range(nw311_.length(0)):
                nw311_[i0_51_] = init51_(i0_51_)
            d_1815_v55_ = nw311_
            def lambda87_(d_1821_p2_):
                def lambda88_(d_1822_i6_):
                    return _dafny.Map({self.f4: d_1821_p2_})

                return lambda88_

            init52_ = lambda87_(p2)
            nw312_ = _dafny.Array(None, 13)
            for i0_52_ in range(nw312_.length(0)):
                nw312_[i0_52_] = init52_(i0_52_)
            d_1815_v55_ = nw312_
            d_1823_v56_: _dafny.Array
            nw313_ = _dafny.Array(_dafny.Seq({}), 15)
            d_1823_v56_ = nw313_
            d_1824_v57_: _dafny.Map
            d_1824_v57_ = _dafny.Map({p2: self.f4})
            d_1825_v58_: D1
            d_1825_v58_ = D1_DC3(d_1824_v57_)
            d_1826_v59_: _dafny.Seq
            d_1826_v59_ = _dafny.SeqWithoutIsStrInference([d_1825_v58_])
            index264_ = default__.safeIndex(816, (d_1823_v56_).length(0))
            (d_1823_v56_)[index264_] = (d_1826_v59_) + (d_1826_v59_)
            index265_ = default__.safeIndex(816, (d_1823_v56_).length(0))
            (d_1823_v56_)[index265_] = default__.fm41(globalState)
            d_1824_v57_ = (d_1824_v57_).set(not(p2), self.f4)
            d_1827_v60_: _dafny.Seq
            d_1827_v60_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qx"))
            d_1828_v61_: str
            d_1828_v61_ = _dafny.CodePoint('u')
            d_1827_v60_ = ((d_1827_v60_) + ((d_1827_v60_).set(default__.safeIndex(self.f4, len(d_1827_v60_)), d_1828_v61_))) + (d_1827_v60_)
        d_1829_v62_: C0
        nw314_ = C0()
        nw314_.ctor__(self.f4)
        d_1829_v62_ = nw314_
        d_1830_v63_: str
        d_1830_v63_ = _dafny.CodePoint('x')
        d_1831_v64_: _dafny.Array
        def lambda89_(d_1832_p0_):
            def lambda90_(d_1833_i7_):
                return d_1832_p0_

            return lambda90_

        init53_ = lambda89_(p0)
        nw315_ = _dafny.Array(None, 14)
        for i0_53_ in range(nw315_.length(0)):
            nw315_[i0_53_] = init53_(i0_53_)
        d_1831_v64_ = nw315_
        d_1834_v65_: _dafny.Map
        d_1834_v65_ = _dafny.Map({d_1830_v63_: d_1831_v64_})
        (d_1829_v62_).f18 = len((d_1834_v65_) | ((d_1834_v65_) | (d_1834_v65_)))
        d_1835_v66_: _dafny.Seq
        d_1835_v66_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uwm"))
        d_1836_v67_: D20
        d_1836_v67_ = D20_DC55(d_1835_v66_, d_1829_v62_, d_1829_v62_.f18)
        source16_ = d_1836_v67_
        if source16_.is_DC54:
            d_1837___mcc_h0_ = source16_.cf97
            d_1838___mcc_h1_ = source16_.cf98
            d_1839___mcc_h2_ = source16_.cf99
            d_1840___mcc_h3_ = source16_.cf100
            d_1841_cf100_ = d_1840___mcc_h3_
            d_1842_cf99_ = d_1839___mcc_h2_
            d_1843_cf98_ = d_1838___mcc_h1_
            d_1844_cf97_ = d_1837___mcc_h0_
            d_1845_v68_: D3
            d_1845_v68_ = D3_DC10(self.f4, self.f4, p2, (0) - (d_1829_v62_.f18), d_1843_cf98_)
            d_1846_v69_: D3
            d_1846_v69_ = D3_DC11(d_1845_v68_)
            d_1847_v70_: D3
            d_1847_v70_ = D3_DC11((d_1846_v69_).cf22)
            source17_ = d_1846_v69_
            if source17_.is_DC8:
                d_1848___mcc_h9_ = source17_.cf13
                d_1849___mcc_h10_ = source17_.cf14
                d_1850_cf14_ = d_1849___mcc_h10_
                d_1851_cf13_ = d_1848___mcc_h9_
                d_1852_v71_: _dafny.Seq
                d_1852_v71_ = _dafny.SeqWithoutIsStrInference([d_1829_v62_.f18, d_1829_v62_.f18])
                d_1853_v72_: _dafny.Array
                nw316_ = _dafny.Array(None, 27)
                nw316_[int(0)] = p1
                nw316_[int(1)] = default__.safeModuloInt(d_1829_v62_.f18, self.f4)
                nw316_[int(2)] = d_1829_v62_.f18
                nw316_[int(3)] = (p1) * (len(d_1835_v66_))
                nw316_[int(4)] = d_1829_v62_.f18
                nw316_[int(5)] = d_1829_v62_.f18
                nw316_[int(6)] = (d_1829_v62_.f18 if p2 else d_1829_v62_.f18)
                nw316_[int(7)] = default__.safeModuloInt(d_1829_v62_.f18, p1)
                nw316_[int(8)] = default__.fm1(globalState)
                nw316_[int(9)] = self.f4
                nw316_[int(10)] = default__.fm1(globalState)
                nw316_[int(11)] = d_1829_v62_.f18
                nw316_[int(12)] = d_1829_v62_.f18
                nw316_[int(13)] = (0) - (d_1829_v62_.f18)
                nw316_[int(14)] = (d_1842_cf99_) + (d_1829_v62_.f18)
                nw316_[int(15)] = (d_1852_v71_)[default__.safeIndex(-466, len(d_1852_v71_))]
                nw316_[int(16)] = default__.safeDivisionInt(default__.fm1(globalState), d_1829_v62_.f18)
                nw316_[int(17)] = default__.safeModuloInt(p1, p1)
                nw316_[int(18)] = d_1829_v62_.f18
                nw316_[int(19)] = p1
                nw316_[int(20)] = d_1829_v62_.f18
                nw316_[int(21)] = d_1842_cf99_
                nw316_[int(22)] = self.f4
                nw316_[int(23)] = self.f4
                nw316_[int(24)] = default__.safeDivisionInt(len(d_1852_v71_), d_1829_v62_.f18)
                nw316_[int(25)] = (0) - (d_1829_v62_.f18)
                nw316_[int(26)] = default__.safeDivisionInt(d_1829_v62_.f18, 883)
                d_1853_v72_ = nw316_
                index266_ = default__.safeIndex(803, (d_1853_v72_).length(0))
                (d_1853_v72_)[index266_] = (p1) - (273)
                index267_ = default__.safeIndex(803, (d_1853_v72_).length(0))
                (d_1853_v72_)[index267_] = (default__.safeModuloInt(d_1829_v62_.f18, default__.fm1(globalState))) * (self.f4)
                d_1854_v73_: _dafny.Map
                d_1854_v73_ = _dafny.Map({d_1851_cf13_: p2})
                d_1855_v74_: _dafny.Map
                d_1855_v74_ = _dafny.Map({d_1841_cf100_: ((d_1854_v73_)[d_1851_cf13_] if (d_1851_cf13_) in (d_1854_v73_) else d_1844_cf97_)})
                d_1856_v75_: _dafny.Map
                d_1856_v75_ = _dafny.Map({False: ((d_1855_v74_)[False] if (False) in (d_1855_v74_) else False)})
                (d_1829_v62_).f18 = (self.f4) - ((D3_DC10(self.f4, p1, d_1841_cf100_, len(d_1852_v71_), not(((d_1855_v74_)[p2] if (p2) in (d_1855_v74_) else default__.fm0(d_1829_v62_.f18, -686, globalState))))).cf17)
                (d_1829_v62_).f18 = p1
                d_1857_v76_: _dafny.Map
                d_1857_v76_ = _dafny.Map({(d_1853_v72_)[default__.safeIndex(803, (d_1853_v72_).length(0))]: d_1843_cf98_})
                d_1858_v77_: _dafny.Map
                d_1858_v77_ = _dafny.Map({d_1829_v62_.f18: d_1853_v72_})
                d_1859_v78_: int
                d_1860_v79_: int
                d_1861_v80_: _dafny.Array
                d_1862_v81_: bool
                out37_: int
                out38_: int
                out39_: _dafny.Array
                out40_: bool
                out37_, out38_, out39_, out40_ = default__.m0((d_1857_v76_) | (_dafny.Map({p1: True})), (d_1852_v71_) + (d_1852_v71_), d_1858_v77_, globalState)
                d_1859_v78_ = out37_
                d_1860_v79_ = out38_
                d_1861_v80_ = out39_
                d_1862_v81_ = out40_
            elif source17_.is_DC9:
                d_1863___mcc_h11_ = source17_.cf15
                d_1864___mcc_h12_ = source17_.cf16
                d_1865_cf16_ = d_1864___mcc_h12_
                d_1866_cf15_ = d_1863___mcc_h11_
                d_1867_v82_: _dafny.Array
                def lambda91_(d_1868_p0_):
                    def lambda92_(d_1869_i8_):
                        return (_dafny.SeqWithoutIsStrInference([False])) + (_dafny.SeqWithoutIsStrInference([d_1868_p0_]))

                    return lambda92_

                init54_ = lambda91_(p0)
                nw317_ = _dafny.Array(None, 3)
                for i0_54_ in range(nw317_.length(0)):
                    nw317_[i0_54_] = init54_(i0_54_)
                d_1867_v82_ = nw317_
                d_1867_v82_ = d_1867_v82_
                d_1870_v83_: _dafny.Array
                def lambda93_(d_1871_i9_):
                    return default__.safeModuloInt(d_1871_i9_, -326)

                init55_ = lambda93_
                nw318_ = _dafny.Array(None, 7)
                for i0_55_ in range(nw318_.length(0)):
                    nw318_[i0_55_] = init55_(i0_55_)
                d_1870_v83_ = nw318_
                nw319_ = _dafny.Array(int(0), 7)
                d_1870_v83_ = nw319_
                d_1843_cf98_ = not(p0)
                d_1872_v84_: _dafny.MultiSet
                d_1872_v84_ = _dafny.MultiSet([d_1829_v62_.f18, d_1829_v62_.f18, d_1829_v62_.f18])
                d_1873_v85_: C1
                nw320_ = C1()
                nw320_.ctor__(self.f4, (d_1829_v62_.f18) + (d_1829_v62_.f18), d_1842_cf99_, (d_1842_cf99_) in (d_1872_v84_))
                d_1873_v85_ = nw320_
            elif source17_.is_DC10:
                d_1874___mcc_h13_ = source17_.cf17
                d_1875___mcc_h14_ = source17_.cf18
                d_1876___mcc_h15_ = source17_.cf19
                d_1877___mcc_h16_ = source17_.cf20
                d_1878___mcc_h17_ = source17_.cf21
                d_1879_cf21_ = d_1878___mcc_h17_
                d_1880_cf20_ = d_1877___mcc_h16_
                d_1881_cf19_ = d_1876___mcc_h15_
                d_1882_cf18_ = d_1875___mcc_h14_
                d_1883_cf17_ = d_1874___mcc_h13_
                d_1880_cf20_ = (0) - (d_1883_cf17_)
                d_1884_v86_: _dafny.Array
                nw321_ = _dafny.Array(int(0), 7)
                d_1884_v86_ = nw321_
                index268_ = default__.safeIndex(628, (d_1884_v86_).length(0))
                (d_1884_v86_)[index268_] = p1
                d_1885_v87_: _dafny.Map
                d_1885_v87_ = _dafny.Map({(-29) * (d_1882_cf18_): (d_1883_cf17_) + (d_1882_cf18_)})
                d_1886_v88_: _dafny.Set
                d_1886_v88_ = _dafny.Set({(_dafny.MultiSet([p3])).cardinality})
                index269_ = default__.safeIndex(628, (d_1884_v86_).length(0))
                rhs244_ = (0) - (((d_1885_v87_)[(0) - (len((d_1886_v88_) - (d_1886_v88_)))] if ((0) - (len((d_1886_v88_) - (d_1886_v88_)))) in (d_1885_v87_) else d_1882_cf18_))
                rhs245_ = _dafny.SeqWithoutIsStrInference([d_1830_v63_ for d_1887_i10_ in range(default__.abs(800))])
                rhs246_ = (0) - (self.f4)
                lhs143_ = d_1884_v86_
                lhs144_ = default__.safeIndex(628, (d_1884_v86_).length(0))
                lhs143_[lhs144_] = rhs244_
                d_1835_v66_ = rhs245_
                d_1880_cf20_ = rhs246_
                d_1888_v89_: _dafny.Map
                d_1888_v89_ = _dafny.Map({self.f4: not(not(d_1841_cf100_))})
                d_1889_v90_: _dafny.Map
                d_1889_v90_ = _dafny.Map({(d_1882_cf18_) > (d_1842_cf99_): D6_DC20(d_1888_v89_)})
                d_1890_v91_: D6
                d_1890_v91_ = D6_DC20(d_1888_v89_)
                d_1889_v90_ = (d_1889_v90_).set((d_1843_cf98_ if p2 else d_1841_cf100_), d_1890_v91_)
                rhs247_ = ((p1 if True else (0) - (self.f4))) <= (len(d_1886_v88_))
                rhs248_ = d_1835_v66_
                rhs249_ = (0) - (d_1883_cf17_)
                rhs250_ = (d_1835_v66_ if p2 else (d_1835_v66_) + (d_1835_v66_))
                lhs145_ = d_1829_v62_
                r0 = rhs247_
                d_1835_v66_ = rhs248_
                lhs145_.f18 = rhs249_
                d_1835_v66_ = rhs250_
            elif source17_.is_DC7:
                d_1891___mcc_h18_ = source17_.cf12
                d_1892_cf12_ = d_1891___mcc_h18_
                d_1893_v93_: _dafny.Map
                def iife186_():
                    coll90_ = _dafny.Map()
                    compr_90_: int
                    for compr_90_ in _dafny.IntegerRange(425, 766):
                        d_1894_v92_: int = compr_90_
                        if ((425) <= (d_1894_v92_)) and ((d_1894_v92_) < (766)):
                            coll90_[(d_1894_v92_) - ((0) - (len(d_1835_v66_)))] = p3
                    return _dafny.Map(coll90_)
                d_1893_v93_ = _dafny.Map({d_1843_cf98_: iife186_()
                })
                d_1895_v94_: _dafny.Set
                d_1895_v94_ = _dafny.Set({d_1843_cf98_, p0})
                rhs251_ = (_dafny.Set({d_1843_cf98_, p2, p3, p3, default__.fm0(p1, len(d_1893_v93_), globalState)})).isdisjoint(d_1895_v94_)
                rhs252_ = default__.safeModuloInt(d_1829_v62_.f18, d_1842_cf99_)
                rhs253_ = d_1830_v63_
                lhs146_ = self
                r0 = rhs251_
                lhs146_.f4 = rhs252_
                d_1830_v63_ = rhs253_
                d_1896_v95_: C2
                nw322_ = C2()
                nw322_.ctor__(False, len(d_1835_v66_))
                d_1896_v95_ = nw322_
                d_1897_v96_: _dafny.Map
                d_1897_v96_ = _dafny.Map({d_1896_v95_: p3})
                d_1898_v97_: _dafny.Seq
                d_1898_v97_ = _dafny.SeqWithoutIsStrInference([d_1897_v96_])
                d_1899_v98_: D13
                d_1899_v98_ = D13_DC39(d_1898_v97_)
                d_1900_v99_: _dafny.Map
                d_1900_v99_ = _dafny.Map({d_1899_v98_: (d_1843_cf98_) or (d_1844_cf97_)})
                pat_let_tv48_ = d_1898_v97_
                def iife187_(_pat_let48_0):
                    def iife188_(d_1901_dt__update__tmp_h0_):
                        def iife189_(_pat_let49_0):
                            def iife190_(d_1902_dt__update_hcf69_h0_):
                                return D13_DC39(d_1902_dt__update_hcf69_h0_)
                            return iife190_(_pat_let49_0)
                        return iife189_(pat_let_tv48_)
                    return iife188_(_pat_let48_0)
                d_1900_v99_ = (d_1900_v99_).set(iife187_(D13_DC39(d_1898_v97_)), d_1841_cf100_)
                (d_1829_v62_).f18 = default__.safeDivisionInt(d_1842_cf99_, (0) - (d_1829_v62_.f18))
                index270_ = default__.safeIndex(710, (d_1831_v64_).length(0))
                (d_1831_v64_)[index270_] = d_1841_cf100_
                index271_ = default__.safeIndex(710, (d_1831_v64_).length(0))
                (d_1831_v64_)[index271_] = d_1844_cf97_
            elif True:
                d_1903___mcc_h19_ = source17_.cf22
                d_1904_cf22_ = d_1903___mcc_h19_
                index272_ = default__.safeIndex(943, (d_1831_v64_).length(0))
                (d_1831_v64_)[index272_] = p2
                d_1905_v100_: _dafny.MultiSet
                d_1905_v100_ = _dafny.MultiSet([d_1830_v63_])
                index273_ = default__.safeIndex(943, (d_1831_v64_).length(0))
                (d_1831_v64_)[index273_] = (_dafny.MultiSet((d_1835_v66_) + (_dafny.SeqWithoutIsStrInference([d_1830_v63_, d_1830_v63_, d_1830_v63_, d_1830_v63_])))).ispropersubset(d_1905_v100_)
                d_1906_v101_: D4
                d_1906_v101_ = D4_DC12(d_1835_v66_)
                d_1907_v102_: _dafny.Array
                nw323_ = _dafny.Array(None, 23)
                nw323_[int(0)] = d_1905_v100_
                nw323_[int(1)] = (d_1905_v100_).set(d_1830_v63_, default__.abs((0) - ((0) - (d_1829_v62_.f18))))
                nw323_[int(2)] = d_1905_v100_
                nw323_[int(3)] = _dafny.MultiSet([d_1830_v63_, d_1830_v63_])
                nw323_[int(4)] = d_1905_v100_
                nw323_[int(5)] = ((d_1905_v100_).set(d_1830_v63_, default__.abs(d_1829_v62_.f18)) if p0 else _dafny.MultiSet(d_1835_v66_))
                nw323_[int(6)] = d_1905_v100_
                nw323_[int(7)] = d_1905_v100_
                nw323_[int(8)] = d_1905_v100_
                nw323_[int(9)] = (d_1905_v100_).intersection(d_1905_v100_)
                nw323_[int(10)] = _dafny.MultiSet([d_1830_v63_])
                nw323_[int(11)] = d_1905_v100_
                nw323_[int(12)] = d_1905_v100_
                nw323_[int(13)] = d_1905_v100_
                nw323_[int(14)] = (_dafny.MultiSet(d_1835_v66_)) - (d_1905_v100_)
                nw323_[int(15)] = d_1905_v100_
                nw323_[int(16)] = _dafny.MultiSet((d_1906_v101_).cf23)
                nw323_[int(17)] = (d_1905_v100_) | (d_1905_v100_)
                nw323_[int(18)] = (d_1905_v100_).intersection(d_1905_v100_)
                nw323_[int(19)] = d_1905_v100_
                nw323_[int(20)] = d_1905_v100_
                nw323_[int(21)] = _dafny.MultiSet([d_1830_v63_])
                nw323_[int(22)] = ((d_1905_v100_).set(d_1830_v63_, default__.abs(len(d_1835_v66_))) if p0 else d_1905_v100_)
                d_1907_v102_ = nw323_
                nw324_ = _dafny.Array(_dafny.MultiSet({}), 19)
                d_1907_v102_ = nw324_
                d_1908_v103_: _dafny.MultiSet
                d_1908_v103_ = _dafny.MultiSet([p1])
                d_1909_v104_: _dafny.Array
                nw325_ = _dafny.Array(int(0), 21)
                d_1909_v104_ = nw325_
                d_1910_v105_: _dafny.MultiSet
                d_1910_v105_ = _dafny.MultiSet([d_1909_v104_, d_1909_v104_])
                d_1911_v106_: _dafny.Map
                d_1911_v106_ = _dafny.Map({d_1843_cf98_: True})
                d_1835_v66_ = (((d_1835_v66_) + (default__.fm23(True, globalState))).set(default__.safeIndex((d_1842_cf99_) * ((_dafny.MultiSet([d_1829_v62_.f18])).cardinality), len((d_1835_v66_) + (default__.fm23(True, globalState)))), d_1830_v63_)).set(default__.safeIndex(default__.safeDivisionInt(d_1829_v62_.f18, ((d_1908_v103_)[self.f4] if (self.f4) in (d_1908_v103_) else ((d_1910_v105_)[d_1909_v104_] if (d_1909_v104_) in (d_1910_v105_) else 737))), len(((d_1835_v66_) + (default__.fm23(True, globalState))).set(default__.safeIndex((d_1842_cf99_) * ((_dafny.MultiSet([d_1829_v62_.f18])).cardinality), len((d_1835_v66_) + (default__.fm23(True, globalState)))), d_1830_v63_))), default__.fm27(((d_1911_v106_)[d_1843_cf98_] if (d_1843_cf98_) in (d_1911_v106_) else d_1841_cf100_), globalState))
                d_1912_v107_: _dafny.Array
                nw326_ = _dafny.Array(_dafny.Seq({}), 17)
                d_1912_v107_ = nw326_
                d_1913_v108_: _dafny.Seq
                d_1913_v108_ = _dafny.SeqWithoutIsStrInference([p2])
                index274_ = default__.safeIndex(8, (d_1912_v107_).length(0))
                (d_1912_v107_)[index274_] = d_1913_v108_
                d_1914_v109_: _dafny.Seq
                d_1914_v109_ = _dafny.SeqWithoutIsStrInference([d_1829_v62_.f18])
                d_1915_v110_: _dafny.Map
                d_1915_v110_ = _dafny.Map({p0: len(d_1914_v109_)})
                d_1916_v111_: D1
                d_1916_v111_ = D1_DC3(d_1915_v110_)
                d_1917_v112_: _dafny.Map
                d_1917_v112_ = _dafny.Map({d_1830_v63_: d_1913_v108_})
                pat_let_tv49_ = d_1915_v110_
                index275_ = default__.safeIndex(8, (d_1912_v107_).length(0))
                def iife191_(_pat_let50_0):
                    def iife192_(d_1918_dt__update__tmp_h1_):
                        def iife193_(_pat_let51_0):
                            def iife194_(d_1919_dt__update_hcf5_h0_):
                                return D1_DC3(d_1919_dt__update_hcf5_h0_)
                            return iife194_(_pat_let51_0)
                        return iife193_(pat_let_tv49_)
                    return iife192_(_pat_let50_0)
                rhs254_ = (((d_1917_v112_)[d_1830_v63_] if (d_1830_v63_) in (d_1917_v112_) else d_1913_v108_)) + (d_1913_v108_)
                rhs255_ = iife191_(d_1916_v111_)
                rhs256_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pevlxxv"))
                lhs147_ = d_1912_v107_
                lhs148_ = default__.safeIndex(8, (d_1912_v107_).length(0))
                lhs147_[lhs148_] = rhs254_
                d_1916_v111_ = rhs255_
                d_1835_v66_ = rhs256_
            d_1835_v66_ = ((_dafny.SeqWithoutIsStrInference([d_1830_v63_ for d_1920_i11_ in range(default__.abs(-288))])) + ((d_1835_v66_) + (_dafny.SeqWithoutIsStrInference([d_1830_v63_ for d_1921_i12_ in range(default__.abs(474))])))).set(default__.safeIndex((0) - ((0) - (d_1829_v62_.f18)), len((_dafny.SeqWithoutIsStrInference([d_1830_v63_ for d_1922_i11_ in range(default__.abs(-288))])) + ((d_1835_v66_) + (_dafny.SeqWithoutIsStrInference([d_1830_v63_ for d_1923_i12_ in range(default__.abs(474))]))))), _dafny.CodePoint('m'))
            (d_1829_v62_).f18 = (0) - ((self.f4) - (self.f4))
            d_1924_v113_: _dafny.Seq
            d_1924_v113_ = _dafny.SeqWithoutIsStrInference([p2])
            d_1925_v114_: _dafny.Set
            d_1925_v114_ = _dafny.Set({self.f4})
            d_1926_v115_: _dafny.Map
            d_1926_v115_ = _dafny.Map({not(not (p0) or (d_1844_cf97_)): default__.fm42((d_1924_v113_)[default__.safeIndex(self.f4, len(d_1924_v113_))], d_1925_v114_, True, d_1843_cf98_, globalState)})
            d_1927_v116_: _dafny.Map
            d_1927_v116_ = _dafny.Map({d_1830_v63_: d_1844_cf97_})
            d_1928_v117_: D5
            d_1928_v117_ = D5_DC16(((d_1927_v116_)[d_1830_v63_] if (d_1830_v63_) in (d_1927_v116_) else p3), default__.fm0(d_1829_v62_.f18, 439, globalState), d_1843_cf98_, p1, p0)
            pat_let_tv50_ = d_1841_cf100_
            def iife195_(_pat_let52_0):
                def iife196_(d_1929_dt__update__tmp_h2_):
                    def iife197_(_pat_let53_0):
                        def iife198_(d_1930_dt__update_hcf31_h0_):
                            return D5_DC16((d_1929_dt__update__tmp_h2_).cf29, (d_1929_dt__update__tmp_h2_).cf30, d_1930_dt__update_hcf31_h0_, (d_1929_dt__update__tmp_h2_).cf32, (d_1929_dt__update__tmp_h2_).cf33)
                        return iife198_(_pat_let53_0)
                    return iife197_(pat_let_tv50_)
                return iife196_(_pat_let52_0)
            d_1926_v115_ = (d_1926_v115_).set(p3, iife195_(d_1928_v117_))
        elif source16_.is_DC55:
            d_1931___mcc_h4_ = source16_.cf101
            d_1932___mcc_h5_ = source16_.cf102
            d_1933___mcc_h6_ = source16_.cf103
            d_1934_cf103_ = d_1933___mcc_h6_
            d_1935_cf102_ = d_1932___mcc_h5_
            d_1936_cf101_ = d_1931___mcc_h4_
            d_1830_v63_ = d_1830_v63_
            d_1937_v118_: _dafny.Seq
            d_1937_v118_ = _dafny.SeqWithoutIsStrInference([p2])
            d_1934_cf103_ = len(d_1937_v118_)
            d_1938_v119_: D1
            d_1938_v119_ = D1_DC4(d_1935_cf102_.f18, d_1937_v118_, len(d_1835_v66_))
            d_1939_v120_: _dafny.MultiSet
            d_1939_v120_ = _dafny.MultiSet([d_1835_v66_])
            d_1940_v121_: _dafny.Array
            nw327_ = _dafny.Array(None, 21)
            nw327_[int(0)] = d_1937_v118_
            nw327_[int(1)] = (d_1937_v118_).set(default__.safeIndex(d_1935_cf102_.f18, len(d_1937_v118_)), p2)
            nw327_[int(2)] = d_1937_v118_
            nw327_[int(3)] = d_1937_v118_
            nw327_[int(4)] = d_1937_v118_
            nw327_[int(5)] = ((d_1938_v119_).cf7).set(default__.safeIndex((d_1939_v120_).cardinality, len((d_1938_v119_).cf7)), p0)
            nw327_[int(6)] = default__.fm2(d_1830_v63_, self.f4, d_1829_v62_.f18, globalState)
            nw327_[int(7)] = d_1937_v118_
            nw327_[int(8)] = (_dafny.SeqWithoutIsStrInference([not(p2)])).set(default__.safeIndex(self.f4, len(_dafny.SeqWithoutIsStrInference([not(p2)]))), p0)
            nw327_[int(9)] = d_1937_v118_
            nw327_[int(10)] = d_1937_v118_
            nw327_[int(11)] = default__.fm2(d_1830_v63_, self.f4, self.f4, globalState)
            nw327_[int(12)] = d_1937_v118_
            nw327_[int(13)] = _dafny.SeqWithoutIsStrInference([p0, p3])
            nw327_[int(14)] = d_1937_v118_
            nw327_[int(15)] = d_1937_v118_
            nw327_[int(16)] = _dafny.SeqWithoutIsStrInference([(d_1937_v118_)[default__.safeIndex(d_1935_cf102_.f18, len(d_1937_v118_))]])
            nw327_[int(17)] = d_1937_v118_
            nw327_[int(18)] = d_1937_v118_
            nw327_[int(19)] = _dafny.SeqWithoutIsStrInference([p0])
            nw327_[int(20)] = d_1937_v118_
            d_1940_v121_ = nw327_
            d_1941_v122_: _dafny.Seq
            d_1941_v122_ = _dafny.SeqWithoutIsStrInference([d_1940_v121_])
            d_1942_v123_: D18
            d_1942_v123_ = D18_DC49(D18_DC47((d_1941_v122_)[default__.safeIndex(p1, len(d_1941_v122_))]))
            d_1943_v124_: _dafny.Map
            d_1943_v124_ = _dafny.Map({self.f4: d_1942_v123_})
            d_1944_v125_: _dafny.Set
            d_1944_v125_ = _dafny.Set({p3, not(p2)})
            d_1943_v124_ = (d_1943_v124_).set(len(d_1944_v125_), d_1942_v123_)
            r0 = not(p0)
        elif source16_.is_DC53:
            d_1945___mcc_h7_ = source16_.cf96
            d_1946_cf96_ = d_1945___mcc_h7_
            (d_1829_v62_).f18 = (p1) + ((self.f4) * (self.f4))
            index276_ = default__.safeIndex(629, (d_1831_v64_).length(0))
            (d_1831_v64_)[index276_] = p3
            index277_ = default__.safeIndex(36, (d_1831_v64_).length(0))
            (d_1831_v64_)[index277_] = p2
            index278_ = default__.safeIndex(629, (d_1831_v64_).length(0))
            index279_ = default__.safeIndex(36, (d_1831_v64_).length(0))
            rhs257_ = d_1829_v62_.f18
            rhs258_ = p1
            rhs259_ = not(False)
            rhs260_ = p3
            lhs149_ = d_1829_v62_
            lhs150_ = d_1829_v62_
            lhs151_ = d_1831_v64_
            lhs152_ = default__.safeIndex(629, (d_1831_v64_).length(0))
            lhs153_ = d_1831_v64_
            lhs154_ = default__.safeIndex(36, (d_1831_v64_).length(0))
            lhs149_.f18 = rhs257_
            lhs150_.f18 = rhs258_
            lhs151_[lhs152_] = rhs259_
            lhs153_[lhs154_] = rhs260_
            d_1947_v126_: _dafny.Array
            nw328_ = _dafny.Array(_dafny.Array(None, 0), 29)
            d_1947_v126_ = nw328_
            d_1948_v127_: D6
            d_1948_v127_ = D6_DC21(self.f4, p3)
            d_1949_v128_: _dafny.Map
            d_1949_v128_ = _dafny.Map({p0: d_1829_v62_.f18})
            d_1950_v129_: _dafny.Set
            d_1950_v129_ = _dafny.Set({d_1949_v128_, (d_1949_v128_).set((d_1831_v64_)[default__.safeIndex(629, (d_1831_v64_).length(0))], p1), _dafny.Map({p0: len(d_1835_v66_)}), default__.fm28(d_1829_v62_.f18, globalState)})
            rhs261_ = (d_1829_v62_.f18) * ((0) - (len((d_1950_v129_) | (d_1950_v129_))))
            rhs262_ = d_1947_v126_
            rhs263_ = d_1948_v127_
            lhs155_ = d_1829_v62_
            lhs155_.f18 = rhs261_
            d_1947_v126_ = rhs262_
            d_1948_v127_ = rhs263_
            d_1951_v130_: _dafny.Seq
            d_1951_v130_ = _dafny.SeqWithoutIsStrInference([p1])
            index280_ = default__.safeIndex(629, (d_1831_v64_).length(0))
            (d_1831_v64_)[index280_] = not (not((d_1951_v130_) <= (d_1951_v130_))) or (not((d_1831_v64_)[default__.safeIndex(629, (d_1831_v64_).length(0))]))
        elif True:
            d_1952___mcc_h8_ = source16_.cf104
            d_1953_cf104_ = d_1952___mcc_h8_
            d_1954_v131_: _dafny.Set
            d_1954_v131_ = _dafny.Set({d_1829_v62_.f18})
            r0 = ((d_1954_v131_) != (_dafny.Set({len(d_1835_v66_), d_1829_v62_.f18, self.f4})) if p0 else p2)
            rhs264_ = d_1954_v131_
            rhs265_ = p0
            rhs266_ = d_1830_v63_
            d_1954_v131_ = rhs264_
            r0 = rhs265_
            d_1830_v63_ = rhs266_
            (d_1829_v62_).f18 = (0) - ((p1) - (p1))
            d_1955_v132_: _dafny.Map
            d_1955_v132_ = _dafny.Map({(p3) and (True): p2})
            d_1955_v132_ = (d_1955_v132_) | (default__.fm40(globalState))
        r0 = p3
        d_1956_v133_: C3
        nw329_ = C3()
        nw329_.ctor__(len(d_1835_v66_), d_1835_v66_, self.f4)
        d_1956_v133_ = nw329_
        r0 = p3
        return r0

    def m4(self, p0, p1, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r2: _dafny.Array = _dafny.Array(None, 0)
        d_1957_v0_: _dafny.MultiSet
        d_1957_v0_ = _dafny.MultiSet([True])
        d_1958_v1_: _dafny.Set
        d_1958_v1_ = _dafny.Set({d_1957_v0_})
        d_1959_v2_: _dafny.Seq
        d_1959_v2_ = _dafny.SeqWithoutIsStrInference([p1, p1])
        d_1960_v3_: D1
        d_1960_v3_ = D1_DC4(len(default__.fm21(self.f4, not(True), p1, d_1958_v1_, globalState)), d_1959_v2_, self.f4)
        pat_let_tv51_ = p1
        pat_let_tv52_ = p1
        pat_let_tv53_ = p1
        pat_let_tv54_ = p1
        def lambda94_(source19_):
            if source19_.is_DC4:
                d_1961___mcc_h9_ = source19_.cf6
                d_1962___mcc_h10_ = source19_.cf7
                d_1963___mcc_h11_ = source19_.cf8
                d_1964_cf8_ = d_1963___mcc_h11_
                d_1965_cf7_ = d_1962___mcc_h10_
                d_1966_cf6_ = d_1961___mcc_h9_
                return D7_DC24(_dafny.Map({pat_let_tv51_: pat_let_tv52_}))
            elif True:
                d_1967___mcc_h12_ = source19_.cf5
                d_1968_cf5_ = d_1967___mcc_h12_
                return D7_DC24(_dafny.Map({pat_let_tv53_: pat_let_tv54_}))

        source18_ = lambda94_(d_1960_v3_)
        if source18_.is_DC25:
            d_1969___mcc_h0_ = source18_.cf45
            d_1970___mcc_h1_ = source18_.cf46
            d_1971___mcc_h2_ = source18_.cf47
            d_1972_cf47_ = d_1971___mcc_h2_
            d_1973_cf46_ = d_1970___mcc_h1_
            d_1974_cf45_ = d_1969___mcc_h0_
            if (p0) < (d_1972_cf47_):
                d_1974_cf45_ = d_1973_cf46_
                d_1975_v4_: _dafny.Array
                def lambda95_(d_1976_cf47_):
                    def lambda96_(d_1977_i0_):
                        return (d_1977_i0_) + (d_1976_cf47_)

                    return lambda96_

                init56_ = lambda95_(d_1972_cf47_)
                nw330_ = _dafny.Array(None, 6)
                for i0_56_ in range(nw330_.length(0)):
                    nw330_[i0_56_] = init56_(i0_56_)
                d_1975_v4_ = nw330_
                index281_ = default__.safeIndex(243, (d_1975_v4_).length(0))
                (d_1975_v4_)[index281_] = p0
                index282_ = default__.safeIndex(243, (d_1975_v4_).length(0))
                (d_1975_v4_)[index282_] = d_1972_cf47_
                d_1972_cf47_ = (d_1975_v4_)[default__.safeIndex(243, (d_1975_v4_).length(0))]
                d_1978_v5_: _dafny.Array
                nw331_ = _dafny.Array(_dafny.MultiSet({}), 6)
                d_1978_v5_ = nw331_
                d_1979_v6_: _dafny.Array
                nw332_ = _dafny.Array(None, 4)
                nw332_[int(0)] = d_1978_v5_
                nw332_[int(1)] = d_1978_v5_
                nw332_[int(2)] = d_1978_v5_
                nw332_[int(3)] = d_1978_v5_
                d_1979_v6_ = nw332_
                index283_ = default__.safeIndex(472, (d_1979_v6_).length(0))
                (d_1979_v6_)[index283_] = d_1978_v5_
                index284_ = default__.safeIndex(472, (d_1979_v6_).length(0))
                (d_1979_v6_)[index284_] = d_1978_v5_
                (self).f4 = d_1972_cf47_
            elif True:
                d_1980_v7_: _dafny.MultiSet
                d_1980_v7_ = _dafny.MultiSet([_dafny.CodePoint('g')])
                d_1981_v8_: str
                d_1981_v8_ = _dafny.CodePoint('c')
                d_1982_v9_: _dafny.Seq
                d_1982_v9_ = _dafny.SeqWithoutIsStrInference([d_1981_v8_])
                d_1983_v10_: _dafny.Map
                d_1983_v10_ = _dafny.Map({(-979) - (d_1972_cf47_): (d_1980_v7_).isdisjoint(_dafny.MultiSet(d_1982_v9_))})
                d_1983_v10_ = (d_1983_v10_).set(402, p1)
                d_1984_v11_: C4
                nw333_ = C4()
                nw333_.ctor__(d_1972_cf47_)
                d_1984_v11_ = nw333_
                d_1985_v12_: _dafny.Seq
                d_1985_v12_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1981_v8_, d_1981_v8_, default__.fm27(False, globalState), d_1981_v8_]), d_1982_v9_, d_1982_v9_, _dafny.SeqWithoutIsStrInference([d_1981_v8_ for d_1986_i1_ in range(default__.abs(368))])])
                d_1987_v13_: C3
                nw334_ = C3()
                nw334_.ctor__(self.f4, (d_1982_v9_) + ((d_1985_v12_)[default__.safeIndex(d_1972_cf47_, len(d_1985_v12_))]), default__.fm1(globalState))
                d_1987_v13_ = nw334_
                d_1973_cf46_ = (self).fm8(globalState)
                d_1988_v14_: _dafny.Array
                def lambda97_(d_1989_cf45_):
                    def lambda98_(d_1990_i2_):
                        return d_1989_cf45_

                    return lambda98_

                init57_ = lambda97_(d_1974_cf45_)
                nw335_ = _dafny.Array(None, 23)
                for i0_57_ in range(nw335_.length(0)):
                    nw335_[i0_57_] = init57_(i0_57_)
                d_1988_v14_ = nw335_
                index285_ = default__.safeIndex(100, (d_1988_v14_).length(0))
                (d_1988_v14_)[index285_] = d_1974_cf45_
                d_1991_v15_: _dafny.Map
                d_1991_v15_ = _dafny.Map({p1: d_1973_cf46_})
                d_1992_v16_: _dafny.Map
                d_1992_v16_ = _dafny.Map({len((d_1987_v13_).f12): d_1972_cf47_})
                index286_ = default__.safeIndex(100, (d_1988_v14_).length(0))
                rhs267_ = default__.safeModuloInt(((d_1987_v13_).f11) + (len(d_1991_v15_)), self.f4)
                rhs268_ = (d_1982_v9_) < (d_1982_v9_)
                rhs269_ = d_1974_cf45_
                rhs270_ = (not (d_1974_cf45_) or (d_1974_cf45_)) == (d_1974_cf45_)
                rhs271_ = default__.safeModuloInt((d_1987_v13_).f11, len(d_1992_v16_))
                lhs156_ = d_1988_v14_
                lhs157_ = default__.safeIndex(100, (d_1988_v14_).length(0))
                lhs158_ = self
                d_1972_cf47_ = rhs267_
                d_1973_cf46_ = rhs268_
                lhs156_[lhs157_] = rhs269_
                d_1974_cf45_ = rhs270_
                lhs158_.f4 = rhs271_
            (self).f4 = self.f4
            d_1993_v17_: _dafny.Array
            def lambda99_(d_1994_p0_):
                def lambda100_(d_1995_i3_):
                    return (d_1995_i3_) - (d_1994_p0_)

                return lambda100_

            init58_ = lambda99_(p0)
            nw336_ = _dafny.Array(None, 26)
            for i0_58_ in range(nw336_.length(0)):
                nw336_[i0_58_] = init58_(i0_58_)
            d_1993_v17_ = nw336_
            index287_ = default__.safeIndex(669, (d_1993_v17_).length(0))
            (d_1993_v17_)[index287_] = -304
            index288_ = default__.safeIndex(669, (d_1993_v17_).length(0))
            (d_1993_v17_)[index288_] = (p0 if p1 else default__.safeModuloInt(self.f4, self.f4))
            d_1996_v18_: _dafny.Array
            nw337_ = _dafny.Array(False, 11)
            d_1996_v18_ = nw337_
            index289_ = default__.safeIndex(292, (d_1996_v18_).length(0))
            (d_1996_v18_)[index289_] = (self.f4) < (p0)
            index290_ = default__.safeIndex(292, (d_1996_v18_).length(0))
            (d_1996_v18_)[index290_] = (p1) or (p1)
        elif source18_.is_DC26:
            d_1997___mcc_h3_ = source18_.cf48
            d_1998___mcc_h4_ = source18_.cf49
            d_1999_cf49_ = d_1998___mcc_h4_
            d_2000_cf48_ = d_1997___mcc_h3_
            d_2001_v19_: _dafny.Array
            nw338_ = _dafny.Array(False, 10)
            d_2001_v19_ = nw338_
            index291_ = default__.safeIndex(857, (d_2001_v19_).length(0))
            (d_2001_v19_)[index291_] = (p1) == (p1)
            index292_ = default__.safeIndex(857, (d_2001_v19_).length(0))
            (d_2001_v19_)[index292_] = not(p1)
            index293_ = default__.safeIndex(857, (d_2001_v19_).length(0))
            (d_2001_v19_)[index293_] = (p1 if (True if (d_2001_v19_)[default__.safeIndex(857, (d_2001_v19_).length(0))] else (d_2001_v19_)[default__.safeIndex(857, (d_2001_v19_).length(0))]) else (d_2001_v19_)[default__.safeIndex(857, (d_2001_v19_).length(0))])
            if (d_2001_v19_)[default__.safeIndex(857, (d_2001_v19_).length(0))]:
                d_1999_cf49_ = self.f4
                d_2002_v20_: _dafny.Seq
                d_2002_v20_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "knpddun"))
                d_2003_v21_: _dafny.MultiSet
                d_2003_v21_ = _dafny.MultiSet([len(d_2002_v20_), p0, -491, (self.f4) * (d_2000_cf48_), d_2000_cf48_])
                d_2004_v22_: _dafny.Map
                d_2004_v22_ = _dafny.Map({d_2002_v20_: True})
                rhs272_ = (d_2003_v21_) | (d_2003_v21_)
                rhs273_ = default__.safeModuloInt(len((d_2004_v22_) | (d_2004_v22_)), len(d_1959_v2_))
                d_2003_v21_ = rhs272_
                d_2000_cf48_ = rhs273_
                d_1999_cf49_ = (len(d_2004_v22_)) + (d_2000_cf48_)
                d_2000_cf48_ = self.f4
                d_2005_v23_: _dafny.Array
                nw339_ = _dafny.Array(int(0), 24)
                d_2005_v23_ = nw339_
                r2 = (d_2005_v23_ if (d_2001_v19_)[default__.safeIndex(857, (d_2001_v19_).length(0))] else d_2005_v23_)
            elif True:
                d_2006_v24_: _dafny.Array
                nw340_ = _dafny.Array(D1.default()(), 3)
                d_2006_v24_ = nw340_
                d_2007_v25_: _dafny.Array
                nw341_ = _dafny.Array(None, 2)
                nw341_[int(0)] = len(d_1959_v2_)
                nw341_[int(1)] = d_1999_cf49_
                d_2007_v25_ = nw341_
                index294_ = default__.safeIndex(940, (d_2007_v25_).length(0))
                (d_2007_v25_)[index294_] = p0
                d_2008_v26_: _dafny.Map
                d_2008_v26_ = _dafny.Map({d_2000_cf48_: True})
                d_2009_v27_: _dafny.Map
                d_2009_v27_ = _dafny.Map({d_1999_cf49_: d_2008_v26_})
                index295_ = default__.safeIndex(940, (d_2007_v25_).length(0))
                rhs274_ = d_2006_v24_
                rhs275_ = d_2001_v19_
                rhs276_ = d_2000_cf48_
                rhs277_ = d_2009_v27_
                lhs159_ = d_2007_v25_
                lhs160_ = default__.safeIndex(940, (d_2007_v25_).length(0))
                d_2006_v24_ = rhs274_
                d_2001_v19_ = rhs275_
                lhs159_[lhs160_] = rhs276_
                d_2009_v27_ = rhs277_
                d_2010_v28_: _dafny.Array
                nw342_ = _dafny.Array(_dafny.Map({}), 26)
                d_2010_v28_ = nw342_
                d_2011_v29_: T0
                nw343_ = C1()
                nw343_.ctor__(308, (d_2007_v25_)[default__.safeIndex(940, (d_2007_v25_).length(0))], len(_dafny.SeqWithoutIsStrInference([d_1959_v2_ for d_2012_i4_ in range(default__.abs(-952))])), p1)
                d_2011_v29_ = nw343_
                d_2013_v30_: _dafny.Map
                d_2013_v30_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bv")): d_2011_v29_})
                d_2014_v31_: _dafny.Seq
                d_2014_v31_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pxdxutayt"))
                index296_ = default__.safeIndex(380, (d_2010_v28_).length(0))
                (d_2010_v28_)[index296_] = (d_2013_v30_).set(d_2014_v31_, d_2011_v29_)
                index297_ = default__.safeIndex(380, (d_2010_v28_).length(0))
                (d_2010_v28_)[index297_] = (d_2013_v30_).set(d_2014_v31_, (d_2011_v29_ if (d_2001_v19_)[default__.safeIndex(857, (d_2001_v19_).length(0))] else d_2011_v29_))
                d_2015_v32_: _dafny.MultiSet
                d_2015_v32_ = _dafny.MultiSet([-179, 324])
                d_2016_v33_: _dafny.Map
                d_2016_v33_ = _dafny.Map({p1: d_2015_v32_})
                index298_ = default__.safeIndex(857, (d_2001_v19_).length(0))
                rhs278_ = len((d_2016_v33_) | (d_2016_v33_))
                rhs279_ = (d_2001_v19_)[default__.safeIndex(857, (d_2001_v19_).length(0))]
                rhs280_ = default__.fm1(globalState)
                rhs281_ = d_2007_v25_
                lhs161_ = d_2011_v29_
                lhs162_ = d_2001_v19_
                lhs163_ = default__.safeIndex(857, (d_2001_v19_).length(0))
                lhs161_.f4 = rhs278_
                lhs162_[lhs163_] = rhs279_
                d_2000_cf48_ = rhs280_
                d_2007_v25_ = rhs281_
                d_2017_v34_: _dafny.Set
                d_2017_v34_ = _dafny.Set({(0) - (d_1999_cf49_)})
                d_2018_v35_: _dafny.Map
                d_2018_v35_ = _dafny.Map({self.f4: d_2017_v34_})
                def iife199_():
                    coll91_ = _dafny.Set()
                    compr_91_: int
                    for compr_91_ in _dafny.IntegerRange(432, 924):
                        d_2019_v36_: int = compr_91_
                        if ((432) <= (d_2019_v36_)) and ((d_2019_v36_) < (924)):
                            coll91_ = coll91_.union(_dafny.Set([(d_2019_v36_) + (115)]))
                    return _dafny.Set(coll91_)
                d_2018_v35_ = (d_2018_v35_).set(385, iife199_()
                )
                index299_ = default__.safeIndex(857, (d_2001_v19_).length(0))
                (d_2001_v19_)[index299_] = p1
            d_2020_v37_: C1
            nw344_ = C1()
            nw344_.ctor__(p0, len(_dafny.Set({not(False)})), len(d_1959_v2_), p1)
            d_2020_v37_ = nw344_
            d_2021_v38_: _dafny.Set
            d_2021_v38_ = _dafny.Set({d_2020_v37_, d_2020_v37_, d_2020_v37_})
            d_2021_v38_ = d_2021_v38_
        elif source18_.is_DC27:
            d_2022___mcc_h5_ = source18_.cf50
            d_2023___mcc_h6_ = source18_.cf51
            d_2024___mcc_h7_ = source18_.cf52
            d_2025_cf52_ = d_2024___mcc_h7_
            d_2026_cf51_ = d_2023___mcc_h6_
            d_2027_cf50_ = d_2022___mcc_h5_
            if p1:
                d_2028_v39_: _dafny.Seq
                d_2028_v39_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b"))
                d_2029_v40_: str
                d_2029_v40_ = _dafny.CodePoint('h')
                d_2030_v41_: _dafny.Map
                d_2030_v41_ = _dafny.Map({p0: d_2029_v40_})
                d_2031_v42_: _dafny.MultiSet
                d_2031_v42_ = _dafny.MultiSet([d_2025_cf52_])
                d_2032_v43_: _dafny.Array
                nw345_ = _dafny.Array(False, 2)
                d_2032_v43_ = nw345_
                d_2033_v44_: D3
                d_2033_v44_ = D3_DC8(((d_2030_v41_)[(d_2031_v42_).cardinality] if ((d_2031_v42_).cardinality) in (d_2030_v41_) else d_2029_v40_), d_2032_v43_)
                d_2026_cf51_ = (len((d_2028_v39_).set(default__.safeIndex(d_2027_cf50_, len(d_2028_v39_)), (d_2033_v44_).cf13))) * (((d_2031_v42_)[(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bsfeide"))))] if ((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bsfeide"))))) in (d_2031_v42_) else d_2025_cf52_))
                d_1959_v2_ = (d_1959_v2_) + (d_1959_v2_)
                d_2027_cf50_ = (0) - (default__.safeDivisionInt(p0, p0))
                d_2034_v45_: bool
                d_2034_v45_ = True
                d_2035_v46_: _dafny.Array
                nw346_ = _dafny.Array(None, 22)
                nw346_[int(0)] = _dafny.SeqWithoutIsStrInference([default__.fm0(d_2025_cf52_, p0, globalState)])
                nw346_[int(1)] = (d_1959_v2_).set(default__.safeIndex(default__.fm1(globalState), len(d_1959_v2_)), p1)
                nw346_[int(2)] = (_dafny.SeqWithoutIsStrInference([p1]) if False else d_1959_v2_)
                nw346_[int(3)] = d_1959_v2_
                nw346_[int(4)] = (_dafny.SeqWithoutIsStrInference([p1])) + (d_1959_v2_)
                nw346_[int(5)] = ((d_1959_v2_).set(default__.safeIndex(d_2026_cf51_, len(d_1959_v2_)), False)) + (default__.fm2(d_2029_v40_, self.f4, p0, globalState))
                nw346_[int(6)] = d_1959_v2_
                nw346_[int(7)] = d_1959_v2_
                nw346_[int(8)] = d_1959_v2_
                nw346_[int(9)] = (d_1959_v2_) + (d_1959_v2_)
                nw346_[int(10)] = d_1959_v2_
                nw346_[int(11)] = ((d_1959_v2_) + (d_1959_v2_)).set(default__.safeIndex(len(d_1959_v2_), len((d_1959_v2_) + (d_1959_v2_))), d_2034_v45_)
                nw346_[int(12)] = (d_1959_v2_) + (default__.fm2(d_2029_v40_, self.f4, p0, globalState))
                nw346_[int(13)] = d_1959_v2_
                nw346_[int(14)] = d_1959_v2_
                nw346_[int(15)] = (d_1959_v2_ if not(p1) else (_dafny.SeqWithoutIsStrInference([False, True])).set(default__.safeIndex(d_2027_cf50_, len(_dafny.SeqWithoutIsStrInference([False, True]))), p1))
                nw346_[int(16)] = d_1959_v2_
                nw346_[int(17)] = d_1959_v2_
                nw346_[int(18)] = d_1959_v2_
                nw346_[int(19)] = d_1959_v2_
                nw346_[int(20)] = d_1959_v2_
                nw346_[int(21)] = d_1959_v2_
                d_2035_v46_ = nw346_
                index300_ = default__.safeIndex(847, (d_2035_v46_).length(0))
                (d_2035_v46_)[index300_] = d_1959_v2_
                d_2036_v47_: C0
                nw347_ = C0()
                nw347_.ctor__(d_2027_cf50_)
                d_2036_v47_ = nw347_
                d_2037_v48_: D20
                d_2037_v48_ = D20_DC55(_dafny.SeqWithoutIsStrInference([d_2029_v40_ for d_2038_i5_ in range(default__.abs(936))]), d_2036_v47_, self.f4)
                index301_ = default__.safeIndex(847, (d_2035_v46_).length(0))
                rhs282_ = not ((len(d_1959_v2_)) != ((0) - (len((d_2037_v48_).cf101)))) or ((p1 if d_2034_v45_ else d_2034_v45_))
                rhs283_ = d_1959_v2_
                lhs164_ = d_2035_v46_
                lhs165_ = default__.safeIndex(847, (d_2035_v46_).length(0))
                d_2034_v45_ = rhs282_
                lhs164_[lhs165_] = rhs283_
                d_2034_v45_ = p1
            elif True:
                d_2039_v49_: _dafny.Map
                d_2039_v49_ = _dafny.Map({len(default__.fm43(globalState)): False})
                d_2040_v50_: _dafny.Map
                d_2040_v50_ = _dafny.Map({p1: default__.fm1(globalState)})
                d_2039_v49_ = (d_2039_v49_).set(d_2026_cf51_, default__.fm0(p0, ((d_2040_v50_)[p1] if (p1) in (d_2040_v50_) else p0), globalState))
                d_2041_v51_: _dafny.Array
                nw348_ = _dafny.Array(_dafny.Map({}), 5)
                d_2041_v51_ = nw348_
                d_2042_v52_: D23
                d_2042_v52_ = D23_DC60(d_2041_v51_)
                d_2041_v51_ = (d_2042_v52_).cf108
                d_2025_cf52_ = (p0) + (default__.fm1(globalState))
                d_2043_v53_: _dafny.Array
                nw349_ = _dafny.Array(None, 15)
                d_2043_v53_ = nw349_
                d_2044_v54_: _dafny.Array
                d_2044_v54_ = d_2043_v53_
                d_2045_v55_: _dafny.MultiSet
                d_2045_v55_ = _dafny.MultiSet([d_2044_v54_])
                d_2046_v56_: _dafny.Array
                nw350_ = _dafny.Array(None, 9)
                nw350_[int(0)] = p1
                nw350_[int(1)] = p1
                nw350_[int(2)] = p1
                nw350_[int(3)] = p1
                nw350_[int(4)] = p1
                nw350_[int(5)] = p1
                nw350_[int(6)] = False
                nw350_[int(7)] = p1
                nw350_[int(8)] = p1
                d_2046_v56_ = nw350_
                d_2047_v57_: _dafny.Map
                d_2047_v57_ = _dafny.Map({p0: d_2046_v56_})
                d_2048_v58_: str
                d_2048_v58_ = _dafny.CodePoint('l')
                d_2049_v59_: _dafny.Map
                d_2049_v59_ = _dafny.Map({(len(default__.fm14(p0, ((d_2045_v55_)[d_2044_v54_] if (d_2044_v54_) in (d_2045_v55_) else len(d_2047_v57_)), p1, d_2048_v58_, globalState))) <= (d_2025_cf52_): p1})
                d_2049_v59_ = (d_2049_v59_).set(default__.fm0(d_2027_cf50_, p0, globalState), p1)
                index302_ = default__.safeIndex(49, (d_2046_v56_).length(0))
                (d_2046_v56_)[index302_] = p1
                index303_ = default__.safeIndex(49, (d_2046_v56_).length(0))
                (d_2046_v56_)[index303_] = not((default__.safeModuloInt(d_2026_cf51_, d_2027_cf50_)) > (p0))
            d_2050_v60_: C0
            nw351_ = C0()
            nw351_.ctor__((d_2027_cf50_) + (d_2027_cf50_))
            d_2050_v60_ = nw351_
            d_2051_v61_: _dafny.Seq
            d_2051_v61_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qovo"))
            d_2052_v62_: _dafny.Map
            d_2052_v62_ = _dafny.Map({d_2051_v61_: True})
            rhs284_ = (len(d_2052_v62_)) * (d_2027_cf50_)
            rhs285_ = (d_2027_cf50_) * (d_2027_cf50_)
            rhs286_ = _dafny.SeqWithoutIsStrInference([p1, (d_1959_v2_)[default__.safeIndex((_dafny.MultiSet([self.f4])).cardinality, len(d_1959_v2_))]])
            d_2026_cf51_ = rhs284_
            d_2025_cf52_ = rhs285_
            d_1959_v2_ = rhs286_
            d_2053_v63_: _dafny.Array
            def lambda101_(d_2054_p1_):
                def lambda102_(d_2055_i6_):
                    return not(d_2054_p1_)

                return lambda102_

            init59_ = lambda101_(p1)
            nw352_ = _dafny.Array(None, 3)
            for i0_59_ in range(nw352_.length(0)):
                nw352_[i0_59_] = init59_(i0_59_)
            d_2053_v63_ = nw352_
            index304_ = default__.safeIndex(786, (d_2053_v63_).length(0))
            (d_2053_v63_)[index304_] = p1
            d_2056_v64_: str
            d_2056_v64_ = _dafny.CodePoint('v')
            d_2057_v65_: _dafny.MultiSet
            d_2057_v65_ = _dafny.MultiSet([d_2056_v64_])
            index305_ = default__.safeIndex(786, (d_2053_v63_).length(0))
            (d_2053_v63_)[index305_] = (self).fm5((d_2051_v61_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_2058_i7_ in range(default__.abs(385))])), (d_2057_v65_).issubset(_dafny.MultiSet([_dafny.CodePoint('q')])), globalState)
        elif True:
            d_2059___mcc_h8_ = source18_.cf44
            d_2060_cf44_ = d_2059___mcc_h8_
            d_2061_v66_: D2
            d_2061_v66_ = D2_DC6(p0, False)
            d_2062_v67_: _dafny.Array
            nw353_ = _dafny.Array(_dafny.Seq({}), 23)
            d_2062_v67_ = nw353_
            index306_ = default__.safeIndex(633, (d_2062_v67_).length(0))
            (d_2062_v67_)[index306_] = (d_1959_v2_) + ((d_1959_v2_).set(default__.safeIndex(self.f4, len(d_1959_v2_)), False))
            d_2063_v68_: str
            d_2063_v68_ = _dafny.CodePoint('n')
            d_2064_v69_: _dafny.Seq
            d_2064_v69_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cdug"))
            d_2065_v70_: _dafny.Array
            nw354_ = _dafny.Array(False, 17)
            d_2065_v70_ = nw354_
            d_2066_v71_: D3
            d_2066_v71_ = D3_DC8(d_2063_v68_, d_2065_v70_)
            d_2067_v72_: _dafny.Array
            nw355_ = _dafny.Array(None, 19)
            nw355_[int(0)] = d_2063_v68_
            nw355_[int(1)] = d_2063_v68_
            nw355_[int(2)] = d_2063_v68_
            nw355_[int(3)] = (_dafny.CodePoint('k') if (self).fm5(d_2064_v69_, p1, globalState) else d_2063_v68_)
            nw355_[int(4)] = (d_2066_v71_).cf13
            nw355_[int(5)] = (d_2063_v68_ if p1 else d_2063_v68_)
            nw355_[int(6)] = default__.fm27(p1, globalState)
            nw355_[int(7)] = d_2063_v68_
            nw355_[int(8)] = (d_2063_v68_ if p1 else d_2063_v68_)
            nw355_[int(9)] = d_2063_v68_
            nw355_[int(10)] = d_2063_v68_
            nw355_[int(11)] = d_2063_v68_
            nw355_[int(12)] = d_2063_v68_
            nw355_[int(13)] = (d_2064_v69_)[default__.safeIndex(p0, len(d_2064_v69_))]
            nw355_[int(14)] = d_2063_v68_
            nw355_[int(15)] = d_2063_v68_
            nw355_[int(16)] = d_2063_v68_
            nw355_[int(17)] = d_2063_v68_
            nw355_[int(18)] = d_2063_v68_
            d_2067_v72_ = nw355_
            index307_ = default__.safeIndex(651, (d_2067_v72_).length(0))
            (d_2067_v72_)[index307_] = d_2063_v68_
            d_2068_v73_: _dafny.Array
            def lambda103_(d_2069_p1_, d_2070_p0_):
                def lambda104_(d_2071_i8_):
                    return (D1_DC3(_dafny.Map({d_2069_p1_: d_2070_p0_}))).cf5

                return lambda104_

            init60_ = lambda103_(p1, p0)
            nw356_ = _dafny.Array(None, 3)
            for i0_60_ in range(nw356_.length(0)):
                nw356_[i0_60_] = init60_(i0_60_)
            d_2068_v73_ = nw356_
            d_2072_v74_: D23
            d_2072_v74_ = D23_DC60(d_2068_v73_)
            pat_let_tv55_ = d_2068_v73_
            d_2073_v75_: _dafny.Array
            nw357_ = _dafny.Array(None, 18)
            nw357_[int(0)] = d_2072_v74_
            nw357_[int(1)] = d_2072_v74_
            def iife200_(_pat_let54_0):
                def iife201_(d_2074_dt__update__tmp_h0_):
                    def iife202_(_pat_let55_0):
                        def iife203_(d_2075_dt__update_hcf108_h0_):
                            return D23_DC60(d_2075_dt__update_hcf108_h0_)
                        return iife203_(_pat_let55_0)
                    return iife202_(pat_let_tv55_)
                return iife201_(_pat_let54_0)
            nw357_[int(2)] = iife200_(d_2072_v74_)
            nw357_[int(3)] = d_2072_v74_
            nw357_[int(4)] = d_2072_v74_
            nw357_[int(5)] = d_2072_v74_
            nw357_[int(6)] = d_2072_v74_
            nw357_[int(7)] = d_2072_v74_
            nw357_[int(8)] = d_2072_v74_
            nw357_[int(9)] = d_2072_v74_
            nw357_[int(10)] = d_2072_v74_
            nw357_[int(11)] = d_2072_v74_
            nw357_[int(12)] = d_2072_v74_
            nw357_[int(13)] = d_2072_v74_
            nw357_[int(14)] = d_2072_v74_
            nw357_[int(15)] = D23_DC60(d_2068_v73_)
            nw357_[int(16)] = D23_DC60((d_2072_v74_).cf108)
            nw357_[int(17)] = d_2072_v74_
            d_2073_v75_ = nw357_
            index308_ = default__.safeIndex(633, (d_2062_v67_).length(0))
            index309_ = default__.safeIndex(651, (d_2067_v72_).length(0))
            rhs287_ = d_2061_v66_
            rhs288_ = d_1959_v2_
            rhs289_ = default__.fm1(globalState)
            rhs290_ = d_2063_v68_
            rhs291_ = d_2073_v75_
            lhs166_ = d_2062_v67_
            lhs167_ = default__.safeIndex(633, (d_2062_v67_).length(0))
            lhs168_ = self
            lhs169_ = d_2067_v72_
            lhs170_ = default__.safeIndex(651, (d_2067_v72_).length(0))
            d_2061_v66_ = rhs287_
            lhs166_[lhs167_] = rhs288_
            lhs168_.f4 = rhs289_
            lhs169_[lhs170_] = rhs290_
            d_2073_v75_ = rhs291_
            d_2076_v76_: _dafny.Map
            d_2076_v76_ = _dafny.Map({p0: p1})
            d_2077_v77_: _dafny.Array
            nw358_ = _dafny.Array(int(0), 1)
            d_2077_v77_ = nw358_
            d_2078_v78_: _dafny.Map
            d_2078_v78_ = _dafny.Map({self.f4: d_2077_v77_})
            d_2079_v79_: int
            d_2080_v80_: int
            d_2081_v81_: _dafny.Array
            d_2082_v82_: bool
            out41_: int
            out42_: int
            out43_: _dafny.Array
            out44_: bool
            out41_, out42_, out43_, out44_ = default__.m0(d_2076_v76_, _dafny.SeqWithoutIsStrInference([(0) - (self.f4), p0]), d_2078_v78_, globalState)
            d_2079_v79_ = out41_
            d_2080_v80_ = out42_
            d_2081_v81_ = out43_
            d_2082_v82_ = out44_
            nw359_ = _dafny.Array(int(0), 24)
            r2 = nw359_
            d_2082_v82_ = d_2082_v82_
        d_2083_v83_: bool
        d_2083_v83_ = False
        d_2083_v83_ = (d_2083_v83_) == (d_2083_v83_)
        d_2084_v84_: _dafny.Array
        def lambda105_(d_2085_p0_):
            def lambda106_(d_2086_i9_):
                return (d_2085_p0_) >= (d_2085_p0_)

            return lambda106_

        init61_ = lambda105_(p0)
        nw360_ = _dafny.Array(None, 17)
        for i0_61_ in range(nw360_.length(0)):
            nw360_[i0_61_] = init61_(i0_61_)
        d_2084_v84_ = nw360_
        index310_ = default__.safeIndex(628, (d_2084_v84_).length(0))
        (d_2084_v84_)[index310_] = p1
        index311_ = default__.safeIndex(628, (d_2084_v84_).length(0))
        (d_2084_v84_)[index311_] = d_2083_v83_
        (self).f4 = p0
        d_2087_v85_: str
        d_2087_v85_ = _dafny.CodePoint('i')
        d_2088_v86_: _dafny.Array
        nw361_ = _dafny.Array(None, 9)
        nw361_[int(0)] = d_2087_v85_
        nw361_[int(1)] = d_2087_v85_
        nw361_[int(2)] = d_2087_v85_
        nw361_[int(3)] = d_2087_v85_
        nw361_[int(4)] = d_2087_v85_
        nw361_[int(5)] = d_2087_v85_
        nw361_[int(6)] = d_2087_v85_
        nw361_[int(7)] = d_2087_v85_
        nw361_[int(8)] = d_2087_v85_
        d_2088_v86_ = nw361_
        d_2089_v87_: _dafny.Map
        d_2089_v87_ = _dafny.Map({self.f4: d_2088_v86_})
        if (p0) not in ((d_2089_v87_) | ((d_2089_v87_).set(self.f4, d_2088_v86_))):
            index312_ = default__.safeIndex(628, (d_2084_v84_).length(0))
            (d_2084_v84_)[index312_] = d_2083_v83_
            d_2090_v88_: _dafny.Map
            d_2090_v88_ = _dafny.Map({p1: d_2083_v83_})
            d_2090_v88_ = (d_2090_v88_).set(True, d_2083_v83_)
            d_2091_v89_: _dafny.Map
            d_2091_v89_ = _dafny.Map({(0) - (default__.safeDivisionInt(p0, len(_dafny.SeqWithoutIsStrInference([d_2087_v85_ for d_2092_i10_ in range(default__.abs(929))])))): self.f4})
            d_2093_v90_: _dafny.Array
            nw362_ = _dafny.Array(int(0), 2)
            d_2093_v90_ = nw362_
            index313_ = default__.safeIndex(897, (d_2093_v90_).length(0))
            (d_2093_v90_)[index313_] = self.f4
            d_2094_v91_: _dafny.Set
            d_2094_v91_ = _dafny.Set({d_2083_v83_, p1, True, p1})
            d_2095_v92_: _dafny.MultiSet
            d_2095_v92_ = _dafny.MultiSet([len(d_2094_v91_)])
            index314_ = default__.safeIndex(897, (d_2093_v90_).length(0))
            rhs292_ = ((d_2091_v89_) | (_dafny.Map({self.f4: p0}))) | ((d_2091_v89_) | (d_2091_v89_))
            rhs293_ = d_2083_v83_
            rhs294_ = d_2093_v90_
            rhs295_ = default__.safeDivisionInt(self.f4, self.f4)
            rhs296_ = d_2095_v92_
            lhs171_ = d_2093_v90_
            lhs172_ = default__.safeIndex(897, (d_2093_v90_).length(0))
            d_2091_v89_ = rhs292_
            d_2083_v83_ = rhs293_
            r2 = rhs294_
            lhs171_[lhs172_] = rhs295_
            d_2095_v92_ = rhs296_
            index315_ = default__.safeIndex(628, (d_2084_v84_).length(0))
            (d_2084_v84_)[index315_] = d_2083_v83_
            if not(p1):
                d_2083_v83_ = False
                index316_ = default__.safeIndex(897, (d_2093_v90_).length(0))
                (d_2093_v90_)[index316_] = p0
                d_2093_v90_ = d_2093_v90_
                d_2096_v93_: _dafny.Map
                d_2096_v93_ = _dafny.Map({p0: d_2093_v90_})
                d_2096_v93_ = (d_2096_v93_).set(178, d_2093_v90_)
                index317_ = default__.safeIndex(628, (d_2084_v84_).length(0))
                (d_2084_v84_)[index317_] = (d_1959_v2_)[default__.safeIndex(662, len(d_1959_v2_))]
            elif True:
                d_2097_v94_: _dafny.Seq
                d_2097_v94_ = _dafny.SeqWithoutIsStrInference([(self.f4) + ((d_2093_v90_)[default__.safeIndex(897, (d_2093_v90_).length(0))]), p0, (p0) * (p0), p0, ((d_2093_v90_)[default__.safeIndex(897, (d_2093_v90_).length(0))]) * (self.f4)])
                d_2097_v94_ = d_2097_v94_
                d_2098_v95_: _dafny.Map
                d_2098_v95_ = _dafny.Map({(d_2093_v90_)[default__.safeIndex(897, (d_2093_v90_).length(0))]: d_2093_v90_})
                d_2099_v96_: int
                d_2100_v97_: int
                d_2101_v98_: _dafny.Array
                d_2102_v99_: bool
                out45_: int
                out46_: int
                out47_: _dafny.Array
                out48_: bool
                out45_, out46_, out47_, out48_ = default__.m0(_dafny.Map({(d_2093_v90_)[default__.safeIndex(897, (d_2093_v90_).length(0))]: d_2083_v83_}), _dafny.SeqWithoutIsStrInference([p0, (d_2093_v90_)[default__.safeIndex(897, (d_2093_v90_).length(0))]]), d_2098_v95_, globalState)
                d_2099_v96_ = out45_
                d_2100_v97_ = out46_
                d_2101_v98_ = out47_
                d_2102_v99_ = out48_
                d_2103_v100_: _dafny.Map
                d_2103_v100_ = _dafny.Map({d_2100_v97_: p1})
                d_2104_v101_: _dafny.MultiSet
                d_2104_v101_ = _dafny.MultiSet([_dafny.Set({((d_2103_v100_)[(d_2093_v90_)[default__.safeIndex(897, (d_2093_v90_).length(0))]] if ((d_2093_v90_)[default__.safeIndex(897, (d_2093_v90_).length(0))]) in (d_2103_v100_) else p1)}), d_2094_v91_])
                d_2099_v96_ = (((d_2104_v101_)[d_2094_v91_] if (d_2094_v91_) in (d_2104_v101_) else (d_2093_v90_)[default__.safeIndex(897, (d_2093_v90_).length(0))])) + (d_2100_v97_)
                d_2105_v102_: D0
                d_2105_v102_ = D0_DC0(d_2097_v94_)
                d_2106_v103_: int
                d_2107_v104_: int
                d_2108_v105_: _dafny.Array
                d_2109_v106_: bool
                out49_: int
                out50_: int
                out51_: _dafny.Array
                out52_: bool
                out49_, out50_, out51_, out52_ = default__.m0(d_2103_v100_, (d_2105_v102_).cf0, _dafny.Map({317: d_2101_v98_}), globalState)
                d_2106_v103_ = out49_
                d_2107_v104_ = out50_
                d_2108_v105_ = out51_
                d_2109_v106_ = out52_
                d_2110_v107_: _dafny.Seq
                d_2110_v107_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u"))
                d_2111_v108_: _dafny.Map
                d_2111_v108_ = _dafny.Map({(0) - (d_2100_v97_): d_2110_v107_})
                d_2112_v109_: _dafny.Map
                d_2112_v109_ = _dafny.Map({((d_2111_v108_)[d_2106_v103_] if (d_2106_v103_) in (d_2111_v108_) else _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_2113_i11_ in range(default__.abs(316))])): _dafny.CodePoint('h')})
                d_2112_v109_ = (d_2112_v109_) | (d_2112_v109_)
        elif True:
            (self).f4 = p0
            d_2114_v110_: _dafny.MultiSet
            d_2114_v110_ = _dafny.MultiSet([633])
            if (d_2114_v110_).ispropersubset(d_2114_v110_):
                d_2115_v111_: D5
                d_2115_v111_ = D5_DC16(p1, p1, p1, p0, p1)
                index318_ = default__.safeIndex(628, (d_2084_v84_).length(0))
                rhs297_ = p0
                rhs298_ = default__.fm27((d_2084_v84_)[default__.safeIndex(628, (d_2084_v84_).length(0))], globalState)
                rhs299_ = (p1 if (default__.fm1(globalState)) < (self.f4) else (d_2115_v111_).cf31)
                lhs173_ = self
                lhs174_ = d_2084_v84_
                lhs175_ = default__.safeIndex(628, (d_2084_v84_).length(0))
                lhs173_.f4 = rhs297_
                d_2087_v85_ = rhs298_
                lhs174_[lhs175_] = rhs299_
                d_2116_v112_: C1
                nw363_ = C1()
                nw363_.ctor__(p0, (0) - (-941), p0, p1)
                d_2116_v112_ = nw363_
                index319_ = default__.safeIndex(628, (d_2084_v84_).length(0))
                (d_2084_v84_)[index319_] = d_2083_v83_
                d_2117_v113_: _dafny.Seq
                d_2117_v113_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dyttigb"))
                d_2118_v114_: _dafny.Seq
                d_2118_v114_ = _dafny.SeqWithoutIsStrInference([self.f4])
                d_2119_v115_: _dafny.Array
                nw364_ = _dafny.Array(None, 27)
                nw364_[int(0)] = p0
                nw364_[int(1)] = len(_dafny.Map({d_1959_v2_: p0}))
                nw364_[int(2)] = p0
                nw364_[int(3)] = p0
                nw364_[int(4)] = (0) - (p0)
                nw364_[int(5)] = ((0) - (d_2116_v112_.f17)) * (p0)
                nw364_[int(6)] = self.f4
                nw364_[int(7)] = default__.safeDivisionInt(self.f4, p0)
                nw364_[int(8)] = p0
                nw364_[int(9)] = self.f4
                nw364_[int(10)] = (0) - (((d_1957_v0_)[(d_2084_v84_)[default__.safeIndex(628, (d_2084_v84_).length(0))]] if ((d_2084_v84_)[default__.safeIndex(628, (d_2084_v84_).length(0))]) in (d_1957_v0_) else d_2116_v112_.f17))
                nw364_[int(11)] = d_2116_v112_.f17
                nw364_[int(12)] = self.f4
                nw364_[int(13)] = (d_2116_v112_.f17) + (p0)
                nw364_[int(14)] = len(d_2117_v113_)
                nw364_[int(15)] = (0) - (p0)
                nw364_[int(16)] = (46) * (len(_dafny.SeqWithoutIsStrInference([self.f4 for d_2120_i12_ in range(default__.abs(757))])))
                nw364_[int(17)] = self.f4
                nw364_[int(18)] = self.f4
                nw364_[int(19)] = d_2116_v112_.f17
                nw364_[int(20)] = (d_2116_v112_.f17) * (p0)
                nw364_[int(21)] = d_2116_v112_.f17
                nw364_[int(22)] = len(default__.fm23(p1, globalState))
                nw364_[int(23)] = (d_2114_v110_).cardinality
                nw364_[int(24)] = d_2116_v112_.f17
                nw364_[int(25)] = p0
                nw364_[int(26)] = (len(d_2118_v114_)) - (p0)
                d_2119_v115_ = nw364_
                index320_ = default__.safeIndex(715, (d_2119_v115_).length(0))
                (d_2119_v115_)[index320_] = self.f4
                index321_ = default__.safeIndex(715, (d_2119_v115_).length(0))
                (d_2119_v115_)[index321_] = default__.fm1(globalState)
                d_2084_v84_ = (D3_DC8(d_2087_v85_, d_2084_v84_)).cf14
            elif True:
                d_2121_v116_: _dafny.Array
                nw365_ = _dafny.Array(int(0), 27)
                d_2121_v116_ = nw365_
                d_2122_v117_: _dafny.Array
                d_2122_v117_ = d_2121_v116_
                d_2123_v118_: _dafny.Array
                nw366_ = _dafny.Array(None, 24)
                nw366_[int(0)] = d_2121_v116_
                nw366_[int(1)] = d_2121_v116_
                nw366_[int(2)] = d_2121_v116_
                nw366_[int(3)] = d_2121_v116_
                nw366_[int(4)] = d_2121_v116_
                nw366_[int(5)] = d_2121_v116_
                nw366_[int(6)] = d_2121_v116_
                nw366_[int(7)] = d_2121_v116_
                nw366_[int(8)] = d_2121_v116_
                nw366_[int(9)] = d_2121_v116_
                nw366_[int(10)] = d_2121_v116_
                nw366_[int(11)] = d_2121_v116_
                nw366_[int(12)] = d_2121_v116_
                nw366_[int(13)] = d_2121_v116_
                nw366_[int(14)] = (d_2122_v117_)
                nw366_[int(15)] = d_2121_v116_
                nw366_[int(16)] = d_2121_v116_
                nw366_[int(17)] = (d_2122_v117_)
                nw366_[int(18)] = d_2121_v116_
                nw366_[int(19)] = d_2121_v116_
                nw366_[int(20)] = d_2121_v116_
                nw366_[int(21)] = d_2121_v116_
                nw366_[int(22)] = (d_2121_v116_ if d_2083_v83_ else d_2121_v116_)
                nw366_[int(23)] = d_2121_v116_
                d_2123_v118_ = nw366_
                d_2124_v119_: _dafny.Seq
                d_2124_v119_ = _dafny.SeqWithoutIsStrInference([d_2121_v116_, (d_2121_v116_ if d_2083_v83_ else d_2121_v116_)])
                d_2125_v120_: _dafny.Array
                nw367_ = _dafny.Array(_dafny.MultiSet({}), 27)
                d_2125_v120_ = nw367_
                d_2126_v121_: D5
                d_2126_v121_ = D5_DC17(d_1957_v0_, (d_1959_v2_)[default__.safeIndex(self.f4, len(d_1959_v2_))])
                index322_ = default__.safeIndex(580, (d_2125_v120_).length(0))
                (d_2125_v120_)[index322_] = ((d_2126_v121_).cf34) - (_dafny.MultiSet([p1]))
                index323_ = default__.safeIndex(174, (d_2121_v116_).length(0))
                (d_2121_v116_)[index323_] = self.f4
                index324_ = default__.safeIndex(580, (d_2125_v120_).length(0))
                index325_ = default__.safeIndex(174, (d_2121_v116_).length(0))
                rhs300_ = d_2087_v85_
                rhs301_ = d_2123_v118_
                rhs302_ = d_2124_v119_
                rhs303_ = (_dafny.MultiSet([(d_2084_v84_)[default__.safeIndex(628, (d_2084_v84_).length(0))], d_2083_v83_, not(p1)])).intersection((d_1957_v0_) - (d_1957_v0_))
                rhs304_ = self.f4
                lhs176_ = d_2125_v120_
                lhs177_ = default__.safeIndex(580, (d_2125_v120_).length(0))
                lhs178_ = d_2121_v116_
                lhs179_ = default__.safeIndex(174, (d_2121_v116_).length(0))
                d_2087_v85_ = rhs300_
                d_2123_v118_ = rhs301_
                d_2124_v119_ = rhs302_
                lhs176_[lhs177_] = rhs303_
                lhs178_[lhs179_] = rhs304_
                d_2127_v122_: _dafny.Seq
                d_2127_v122_ = _dafny.SeqWithoutIsStrInference([((0) - (self.f4)) * ((d_2121_v116_)[default__.safeIndex(174, (d_2121_v116_).length(0))]), (d_2121_v116_)[default__.safeIndex(174, (d_2121_v116_).length(0))], self.f4])
                d_2127_v122_ = d_2127_v122_
                index326_ = default__.safeIndex(174, (d_2121_v116_).length(0))
                (d_2121_v116_)[index326_] = (690 if p1 else (d_2121_v116_)[default__.safeIndex(174, (d_2121_v116_).length(0))])
                d_2128_v123_: C2
                nw368_ = C2()
                nw368_.ctor__(d_2083_v83_, (d_2121_v116_)[default__.safeIndex(174, (d_2121_v116_).length(0))])
                d_2128_v123_ = nw368_
                d_2129_v124_: _dafny.Seq
                d_2129_v124_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u"))
                r1 = d_2129_v124_
            d_2130_v125_: _dafny.Seq
            d_2130_v125_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "psl"))
            d_2131_v126_: _dafny.Set
            d_2131_v126_ = _dafny.Set({d_1959_v2_})
            d_2132_v127_: _dafny.Map
            d_2132_v127_ = _dafny.Map({(d_2130_v125_ if p1 else d_2130_v125_): d_2131_v126_})
            d_2132_v127_ = (d_2132_v127_).set(d_2130_v125_, _dafny.Set({d_1959_v2_}))
            d_2133_v128_: D11
            d_2133_v128_ = D11_DC35(p0, not((d_2084_v84_)[default__.safeIndex(628, (d_2084_v84_).length(0))]), p1)
            index327_ = default__.safeIndex(628, (d_2084_v84_).length(0))
            (d_2084_v84_)[index327_] = (d_2133_v128_).cf63
            d_2134_v129_: _dafny.Seq
            d_2134_v129_ = _dafny.SeqWithoutIsStrInference([p0])
            d_2135_v130_: _dafny.Map
            d_2135_v130_ = _dafny.Map({not(not (d_2083_v83_) or (p1)): d_2134_v129_})
            d_2135_v130_ = d_2135_v130_
        d_2136_v131_: _dafny.Array
        nw369_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 11)
        d_2136_v131_ = nw369_
        guard_loop_8_: int
        for guard_loop_8_ in _dafny.IntegerRange(0, (d_2136_v131_).length(0)):
            d_2137_i13_: int = guard_loop_8_
            if (True) and (((0) <= (d_2137_i13_)) and ((d_2137_i13_) < ((d_2136_v131_).length(0)))):
                (d_2136_v131_)[(d_2137_i13_)] = (((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vouab"))).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vouab")))), d_2087_v85_)) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ufou"))).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ufou")))), _dafny.CodePoint('f')))) + ((_dafny.SeqWithoutIsStrInference([d_2087_v85_ for d_2138_i14_ in range(default__.abs(604))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jihww"))))
        r0 = _dafny.SeqWithoutIsStrInference([(D0_DC0(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([p0])).cardinality for d_2139_i16_ in range(default__.abs(668))])) if (d_2084_v84_)[default__.safeIndex(628, (d_2084_v84_).length(0))] else D0_DC0(_dafny.SeqWithoutIsStrInference([self.f4]))) for d_2140_i15_ in range(default__.abs(904))])
        d_2141_v132_: _dafny.Seq
        d_2141_v132_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pxwsh"))
        r1 = d_2141_v132_
        d_2142_v133_: C1
        nw370_ = C1()
        nw370_.ctor__(p0, (0) - (p0), p0, d_2083_v83_)
        d_2142_v133_ = nw370_
        d_2143_v134_: _dafny.Array
        nw371_ = _dafny.Array(None, 9)
        nw371_[int(0)] = self.f4
        nw371_[int(1)] = p0
        nw371_[int(2)] = p0
        nw371_[int(3)] = 714
        nw371_[int(4)] = (self.f4) + (p0)
        nw371_[int(5)] = len(_dafny.SeqWithoutIsStrInference([d_2142_v133_]))
        nw371_[int(6)] = ((0) - (self.f4)) + (self.f4)
        nw371_[int(7)] = d_2142_v133_.f17
        nw371_[int(8)] = default__.safeDivisionInt(p0, d_2142_v133_.f17)
        d_2143_v134_ = nw371_
        r2 = d_2143_v134_
        return r0, r1, r2


class C6(T0):
    def  __init__(self):
        self._f4: int = int(0)
        self._f8: bool = False
        self._f9: _dafny.MultiSet = _dafny.MultiSet({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C6"
    @property
    def f4(self):
        return self._f4
    @f4.setter
    def f4(self, value):
        self._f4 = value
    def ctor__(self, f8, f9, f4):
        (self)._f8 = f8
        (self)._f9 = f9
        (self).f4 = f4

    def fm5(self, p0, p1, globalState):
        return (self).f8

    def m3(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        d_2144_v0_: _dafny.Set
        d_2144_v0_ = _dafny.Set({True, p3})
        d_2145_v1_: _dafny.Seq
        d_2145_v1_ = _dafny.SeqWithoutIsStrInference([p3])
        d_2146_v2_: D1
        d_2146_v2_ = D1_DC4(len(d_2144_v0_), d_2145_v1_, self.f4)
        source20_ = d_2146_v2_
        if source20_.is_DC4:
            d_2147___mcc_h0_ = source20_.cf6
            d_2148___mcc_h1_ = source20_.cf7
            d_2149___mcc_h2_ = source20_.cf8
            d_2150_cf8_ = d_2149___mcc_h2_
            d_2151_cf7_ = d_2148___mcc_h1_
            d_2152_cf6_ = d_2147___mcc_h0_
            d_2153_v3_: _dafny.Seq
            d_2153_v3_ = _dafny.SeqWithoutIsStrInference([d_2152_cf6_, default__.safeDivisionInt(self.f4, self.f4), -939, self.f4])
            d_2153_v3_ = (((D0_DC0(d_2153_v3_)).cf0) + (d_2153_v3_)) + (d_2153_v3_)
            d_2154_v4_: _dafny.Map
            d_2154_v4_ = _dafny.Map({(self).f8: len(_dafny.SeqWithoutIsStrInference([d_2152_cf6_ for d_2155_i0_ in range(default__.abs(211))]))})
            d_2156_v5_: D1
            d_2156_v5_ = D1_DC3(d_2154_v4_)
            d_2157_v6_: D1
            d_2157_v6_ = D1_DC3(((d_2156_v5_).cf5) | (d_2154_v4_))
            d_2156_v5_ = (d_2157_v6_ if p3 else d_2156_v5_)
            if p2:
                d_2158_v7_: D0
                d_2158_v7_ = D0_DC1((d_2153_v3_)[default__.safeIndex(d_2152_cf6_, len(d_2153_v3_))], (0) - ((829) + (len(_dafny.SeqWithoutIsStrInference([False])))), p3)
                d_2159_v8_: _dafny.Map
                d_2159_v8_ = _dafny.Map({(self).f8: False})
                rhs305_ = default__.fm1(globalState)
                rhs306_ = D0_DC1(default__.safeModuloInt(p1, len(d_2159_v8_)), d_2150_cf8_, True)
                d_2152_cf6_ = rhs305_
                d_2158_v7_ = rhs306_
                d_2160_v9_: C0
                nw372_ = C0()
                nw372_.ctor__(d_2150_cf8_)
                d_2160_v9_ = nw372_
                d_2161_v10_: C0
                nw373_ = C0()
                nw373_.ctor__(d_2160_v9_.f18)
                d_2161_v10_ = nw373_
                r0 = p2
                d_2162_v11_: _dafny.Array
                nw374_ = _dafny.Array(int(0), 26)
                d_2162_v11_ = nw374_
                index328_ = default__.safeIndex(477, (d_2162_v11_).length(0))
                (d_2162_v11_)[index328_] = ((0) - ((0) - (d_2152_cf6_))) * (p1)
                index329_ = default__.safeIndex(477, (d_2162_v11_).length(0))
                (d_2162_v11_)[index329_] = d_2161_v10_.f18
            elif True:
                d_2163_v12_: _dafny.Map
                d_2163_v12_ = _dafny.Map({(0) - (d_2152_cf6_): d_2150_cf8_})
                d_2163_v12_ = (d_2163_v12_).set(d_2152_cf6_, d_2152_cf6_)
                r0 = default__.fm0(default__.safeDivisionInt(p1, -896), d_2150_cf8_, globalState)
                d_2164_v13_: str
                d_2164_v13_ = _dafny.CodePoint('k')
                d_2165_v14_: _dafny.Set
                d_2165_v14_ = _dafny.Set({_dafny.CodePoint('y'), d_2164_v13_, d_2164_v13_})
                d_2166_v15_: C1
                nw375_ = C1()
                nw375_.ctor__(default__.safeModuloInt(len(d_2163_v12_), p1), (len(d_2165_v14_)) + (self.f4), self.f4, p0)
                d_2166_v15_ = nw375_
                r0 = p3
                (d_2166_v15_).f17 = (d_2166_v15_.f17) * (d_2166_v15_.f17)
            d_2167_v16_: _dafny.Array
            def lambda107_(d_2168_i1_):
                return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h')])

            init62_ = lambda107_
            nw376_ = _dafny.Array(None, 15)
            for i0_62_ in range(nw376_.length(0)):
                nw376_[i0_62_] = init62_(i0_62_)
            d_2167_v16_ = nw376_
            d_2169_v17_: str
            d_2169_v17_ = _dafny.CodePoint('g')
            d_2170_v18_: _dafny.Seq
            d_2170_v18_ = _dafny.SeqWithoutIsStrInference([d_2169_v17_, d_2169_v17_])
            index330_ = default__.safeIndex(633, (d_2167_v16_).length(0))
            (d_2167_v16_)[index330_] = (_dafny.SeqWithoutIsStrInference([d_2169_v17_, d_2169_v17_])) + (d_2170_v18_)
            index331_ = default__.safeIndex(633, (d_2167_v16_).length(0))
            (d_2167_v16_)[index331_] = default__.fm23((self).f8, globalState)
        elif True:
            d_2171___mcc_h3_ = source20_.cf5
            d_2172_cf5_ = d_2171___mcc_h3_
            d_2173_v19_: _dafny.Array
            def lambda108_(d_2174_i2_):
                return (d_2174_i2_) * (self.f4)

            init63_ = lambda108_
            nw377_ = _dafny.Array(None, 6)
            for i0_63_ in range(nw377_.length(0)):
                nw377_[i0_63_] = init63_(i0_63_)
            d_2173_v19_ = nw377_
            d_2175_v20_: _dafny.Array
            nw378_ = _dafny.Array(None, 20)
            nw378_[int(0)] = d_2173_v19_
            nw378_[int(1)] = d_2173_v19_
            nw378_[int(2)] = d_2173_v19_
            nw378_[int(3)] = d_2173_v19_
            nw378_[int(4)] = d_2173_v19_
            nw378_[int(5)] = d_2173_v19_
            nw378_[int(6)] = d_2173_v19_
            nw378_[int(7)] = d_2173_v19_
            nw378_[int(8)] = d_2173_v19_
            nw378_[int(9)] = d_2173_v19_
            nw378_[int(10)] = d_2173_v19_
            nw378_[int(11)] = d_2173_v19_
            nw378_[int(12)] = d_2173_v19_
            nw378_[int(13)] = d_2173_v19_
            nw378_[int(14)] = d_2173_v19_
            nw378_[int(15)] = d_2173_v19_
            nw378_[int(16)] = d_2173_v19_
            nw378_[int(17)] = d_2173_v19_
            nw378_[int(18)] = d_2173_v19_
            nw378_[int(19)] = d_2173_v19_
            d_2175_v20_ = nw378_
            d_2176_v21_: _dafny.Map
            d_2176_v21_ = _dafny.Map({self.f4: d_2175_v20_})
            d_2177_v22_: _dafny.Seq
            d_2177_v22_ = _dafny.SeqWithoutIsStrInference([p1, self.f4])
            d_2178_v23_: _dafny.Array
            nw379_ = _dafny.Array(None, 11)
            nw379_[int(0)] = (d_2175_v20_ if p2 else d_2175_v20_)
            nw379_[int(1)] = d_2175_v20_
            nw379_[int(2)] = ((d_2176_v21_)[len(d_2177_v22_)] if (len(d_2177_v22_)) in (d_2176_v21_) else d_2175_v20_)
            nw379_[int(3)] = d_2175_v20_
            nw379_[int(4)] = d_2175_v20_
            nw379_[int(5)] = d_2175_v20_
            nw379_[int(6)] = d_2175_v20_
            nw379_[int(7)] = d_2175_v20_
            nw379_[int(8)] = d_2175_v20_
            nw379_[int(9)] = d_2175_v20_
            nw379_[int(10)] = d_2175_v20_
            d_2178_v23_ = nw379_
            index332_ = default__.safeIndex(369, (d_2178_v23_).length(0))
            (d_2178_v23_)[index332_] = d_2175_v20_
            d_2179_v24_: _dafny.MultiSet
            d_2179_v24_ = _dafny.MultiSet([p3])
            index333_ = default__.safeIndex(369, (d_2178_v23_).length(0))
            rhs307_ = self.f4
            rhs308_ = d_2175_v20_
            rhs309_ = (((d_2179_v24_)[p2] if (p2) in (d_2179_v24_) else -442)) - ((self.f4 if p3 else self.f4))
            lhs180_ = self
            lhs181_ = d_2178_v23_
            lhs182_ = default__.safeIndex(369, (d_2178_v23_).length(0))
            lhs183_ = self
            lhs180_.f4 = rhs307_
            lhs181_[lhs182_] = rhs308_
            lhs183_.f4 = rhs309_
            d_2180_v25_: _dafny.Array
            nw380_ = _dafny.Array(_dafny.Seq({}), 19)
            d_2180_v25_ = nw380_
            d_2181_v26_: D18
            d_2181_v26_ = D18_DC47(d_2180_v25_)
            source21_ = (d_2181_v26_ if p3 else d_2181_v26_)
            if source21_.is_DC48:
                d_2182___mcc_h4_ = source21_.cf84
                d_2183___mcc_h5_ = source21_.cf85
                d_2184___mcc_h6_ = source21_.cf86
                d_2185___mcc_h7_ = source21_.cf87
                d_2186_cf87_ = d_2185___mcc_h7_
                d_2187_cf86_ = d_2184___mcc_h6_
                d_2188_cf85_ = d_2183___mcc_h5_
                d_2189_cf84_ = d_2182___mcc_h4_
                arr0_ = (d_2178_v23_)[default__.safeIndex(369, (d_2178_v23_).length(0))]
                index334_ = default__.safeIndex(745, ((d_2178_v23_)[default__.safeIndex(369, (d_2178_v23_).length(0))]).length(0))
                arr0_[index334_] = d_2173_v19_
                arr1_ = (d_2178_v23_)[default__.safeIndex(369, (d_2178_v23_).length(0))]
                index335_ = default__.safeIndex(745, ((d_2178_v23_)[default__.safeIndex(369, (d_2178_v23_).length(0))]).length(0))
                arr1_[index335_] = d_2173_v19_
                d_2190_v27_: _dafny.Map
                d_2190_v27_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([p3, not(p3), d_2188_cf85_, p3, d_2188_cf85_]): default__.fm0(d_2186_cf87_.f18, p1, globalState)})
                d_2190_v27_ = default__.fm44(self.f4, (d_2172_cf5_) | (d_2172_cf5_), d_2179_v24_, self.f4, globalState)
                d_2191_v28_: _dafny.Set
                d_2191_v28_ = _dafny.Set({p1})
                r0 = (d_2191_v28_).ispropersubset((d_2191_v28_) | (d_2191_v28_))
                (self).f4 = (d_2186_cf87_.f18 if p0 else self.f4)
            elif source21_.is_DC47:
                d_2192___mcc_h8_ = source21_.cf83
                d_2193_cf83_ = d_2192___mcc_h8_
                (self).f4 = (self.f4) - (342)
                r0 = p0
                index336_ = default__.safeIndex(399, (d_2173_v19_).length(0))
                (d_2173_v19_)[index336_] = self.f4
                index337_ = default__.safeIndex(399, (d_2173_v19_).length(0))
                (d_2173_v19_)[index337_] = (default__.fm1(globalState)) * (self.f4)
                d_2194_v29_: _dafny.Seq
                d_2194_v29_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yhnuhtivm"))
                d_2195_v30_: _dafny.Map
                d_2195_v30_ = _dafny.Map({(self).f8: d_2194_v29_})
                d_2196_v31_: _dafny.Map
                d_2196_v31_ = _dafny.Map({(d_2173_v19_)[default__.safeIndex(399, (d_2173_v19_).length(0))]: d_2194_v29_})
                d_2197_v32_: _dafny.Set
                d_2197_v32_ = _dafny.Set({_dafny.MultiSet(d_2145_v1_), d_2179_v24_})
                d_2198_v33_: _dafny.Array
                nw381_ = _dafny.Array(None, 26)
                nw381_[int(0)] = d_2195_v30_
                nw381_[int(1)] = d_2195_v30_
                nw381_[int(2)] = (_dafny.Map({False: ((d_2196_v31_)[(0) - ((d_2173_v19_)[default__.safeIndex(399, (d_2173_v19_).length(0))])] if ((0) - ((d_2173_v19_)[default__.safeIndex(399, (d_2173_v19_).length(0))])) in (d_2196_v31_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r")))})) | (d_2195_v30_)
                nw381_[int(3)] = d_2195_v30_
                nw381_[int(4)] = d_2195_v30_
                nw381_[int(5)] = (d_2195_v30_).set(p0, default__.fm23((self).f8, globalState))
                nw381_[int(6)] = _dafny.Map({p0: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_2199_i3_ in range(default__.abs(711))])})
                nw381_[int(7)] = d_2195_v30_
                nw381_[int(8)] = (d_2195_v30_) | (d_2195_v30_)
                nw381_[int(9)] = d_2195_v30_
                nw381_[int(10)] = _dafny.Map({p3: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oaaxwb"))})
                nw381_[int(11)] = _dafny.Map({(self).f8: d_2194_v29_})
                nw381_[int(12)] = (d_2195_v30_) | (d_2195_v30_)
                nw381_[int(13)] = (d_2195_v30_) | (d_2195_v30_)
                nw381_[int(14)] = d_2195_v30_
                nw381_[int(15)] = d_2195_v30_
                nw381_[int(16)] = d_2195_v30_
                nw381_[int(17)] = d_2195_v30_
                nw381_[int(18)] = d_2195_v30_
                nw381_[int(19)] = d_2195_v30_
                nw381_[int(20)] = (default__.fm19(d_2177_v22_, len(d_2194_v29_), globalState)) | (default__.fm19(d_2177_v22_, self.f4, globalState))
                nw381_[int(21)] = ((_dafny.Map({p2: d_2194_v29_})).set(p2, d_2194_v29_) if (self).f8 else d_2195_v30_)
                nw381_[int(22)] = default__.fm19(default__.fm21(self.f4, p3, p3, d_2197_v32_, globalState), 730, globalState)
                nw381_[int(23)] = d_2195_v30_
                nw381_[int(24)] = _dafny.Map({not(p0): d_2194_v29_})
                nw381_[int(25)] = d_2195_v30_
                d_2198_v33_ = nw381_
                index338_ = default__.safeIndex(883, (d_2198_v33_).length(0))
                (d_2198_v33_)[index338_] = ((d_2195_v30_).set(p3, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t")))) | (d_2195_v30_)
                index339_ = default__.safeIndex(883, (d_2198_v33_).length(0))
                (d_2198_v33_)[index339_] = (d_2195_v30_) | (d_2195_v30_)
            elif True:
                d_2200___mcc_h9_ = source21_.cf88
                d_2201_cf88_ = d_2200___mcc_h9_
                d_2202_v34_: _dafny.Array
                def lambda109_(d_2203_p3_):
                    def lambda110_(d_2204_i4_):
                        return d_2203_p3_

                    return lambda110_

                init64_ = lambda109_(p3)
                nw382_ = _dafny.Array(None, 11)
                for i0_64_ in range(nw382_.length(0)):
                    nw382_[i0_64_] = init64_(i0_64_)
                d_2202_v34_ = nw382_
                index340_ = default__.safeIndex(747, (d_2202_v34_).length(0))
                (d_2202_v34_)[index340_] = False
                d_2205_v35_: _dafny.Map
                d_2205_v35_ = _dafny.Map({default__.fm0(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tbh"))), p1, globalState): p0})
                d_2206_v36_: _dafny.MultiSet
                d_2206_v36_ = _dafny.MultiSet([d_2205_v35_])
                index341_ = default__.safeIndex(747, (d_2202_v34_).length(0))
                (d_2202_v34_)[index341_] = ((p2) == (p2)) == ((d_2206_v36_) != (d_2206_v36_))
                d_2207_v37_: str
                d_2207_v37_ = _dafny.CodePoint('w')
                d_2208_v38_: _dafny.Seq
                d_2208_v38_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "erx"))
                d_2209_v39_: _dafny.Set
                d_2209_v39_ = _dafny.Set({(0) - (p1)})
                d_2210_v40_: D12
                d_2210_v40_ = D12_DC38(self.f4)
                d_2211_v41_: D19
                d_2211_v41_ = D19_DC51(default__.fm1(globalState), self.f4)
                d_2212_v42_: D7
                d_2212_v42_ = D7_DC27(self.f4, -515, self.f4)
                d_2213_v43_: _dafny.Array
                nw383_ = _dafny.Array(None, 18)
                nw383_[int(0)] = D12_DC38(len(d_2209_v39_))
                nw383_[int(1)] = d_2210_v40_
                def iife204_(_pat_let56_0):
                    def iife205_(d_2214_dt__update__tmp_h0_):
                        def iife206_(_pat_let57_0):
                            def iife207_(d_2215_dt__update_hcf68_h0_):
                                return D12_DC38(d_2215_dt__update_hcf68_h0_)
                            return iife207_(_pat_let57_0)
                        return iife206_(self.f4)
                    return iife205_(_pat_let56_0)
                nw383_[int(2)] = iife204_(d_2210_v40_)
                nw383_[int(3)] = d_2210_v40_
                nw383_[int(4)] = d_2210_v40_
                nw383_[int(5)] = d_2210_v40_
                def iife208_(_pat_let58_0):
                    def iife209_(d_2216_dt__update__tmp_h1_):
                        def iife210_(_pat_let59_0):
                            def iife211_(d_2217_dt__update_hcf68_h1_):
                                return D12_DC38(d_2217_dt__update_hcf68_h1_)
                            return iife211_(_pat_let59_0)
                        return iife210_(self.f4)
                    return iife209_(_pat_let58_0)
                nw383_[int(6)] = iife208_(d_2210_v40_)
                def iife212_(_pat_let60_0):
                    def iife213_(d_2218_dt__update__tmp_h2_):
                        def iife214_(_pat_let61_0):
                            def iife215_(d_2219_dt__update_hcf68_h2_):
                                return D12_DC38(d_2219_dt__update_hcf68_h2_)
                            return iife215_(_pat_let61_0)
                        return iife214_(self.f4)
                    return iife213_(_pat_let60_0)
                nw383_[int(7)] = iife212_(d_2210_v40_)
                nw383_[int(8)] = d_2210_v40_
                nw383_[int(9)] = default__.fm38(d_2211_v41_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a")), (d_2145_v1_)[default__.safeIndex((d_2212_v42_).cf50, len(d_2145_v1_))], globalState)
                nw383_[int(10)] = D12_DC38(p1)
                nw383_[int(11)] = d_2210_v40_
                nw383_[int(12)] = d_2210_v40_
                nw383_[int(13)] = d_2210_v40_
                nw383_[int(14)] = d_2210_v40_
                nw383_[int(15)] = d_2210_v40_
                nw383_[int(16)] = d_2210_v40_
                nw383_[int(17)] = d_2210_v40_
                d_2213_v43_ = nw383_
                index342_ = default__.safeIndex(22, (d_2213_v43_).length(0))
                (d_2213_v43_)[index342_] = d_2210_v40_
                index343_ = default__.safeIndex(22, (d_2213_v43_).length(0))
                rhs310_ = d_2207_v37_
                rhs311_ = (d_2208_v38_) + ((d_2208_v38_) + ((_dafny.SeqWithoutIsStrInference([d_2207_v37_ for d_2220_i5_ in range(default__.abs(629))])).set(default__.safeIndex(p1, len(_dafny.SeqWithoutIsStrInference([d_2207_v37_ for d_2221_i5_ in range(default__.abs(629))]))), d_2207_v37_)))
                rhs312_ = d_2210_v40_
                lhs184_ = d_2213_v43_
                lhs185_ = default__.safeIndex(22, (d_2213_v43_).length(0))
                d_2207_v37_ = rhs310_
                d_2208_v38_ = rhs311_
                lhs184_[lhs185_] = rhs312_
                d_2222_v44_: _dafny.MultiSet
                d_2222_v44_ = _dafny.MultiSet([d_2207_v37_, d_2207_v37_])
                d_2223_v45_: _dafny.Seq
                d_2223_v45_ = _dafny.SeqWithoutIsStrInference([d_2222_v44_, (d_2222_v44_ if p3 else d_2222_v44_), d_2222_v44_])
                d_2222_v44_ = (d_2223_v45_)[default__.safeIndex(default__.safeModuloInt(self.f4, (0) - (self.f4)), len(d_2223_v45_))]
                d_2224_v46_: _dafny.Seq
                d_2225_v47_: int
                out53_: _dafny.Seq
                out54_: int
                out53_, out54_ = (self).m6(globalState)
                d_2224_v46_ = out53_
                d_2225_v47_ = out54_
            index344_ = default__.safeIndex(373, (d_2173_v19_).length(0))
            (d_2173_v19_)[index344_] = ((d_2172_cf5_)[p0] if (p0) in (d_2172_cf5_) else p1)
            index345_ = default__.safeIndex(373, (d_2173_v19_).length(0))
            (d_2173_v19_)[index345_] = 93
            index346_ = default__.safeIndex(373, (d_2173_v19_).length(0))
            (d_2173_v19_)[index346_] = p1
        if default__.fm0(self.f4, p1, globalState):
            d_2226_v48_: C5
            nw384_ = C5()
            nw384_.ctor__(self.f4)
            d_2226_v48_ = nw384_
            (self).f4 = p1
            d_2227_v49_: _dafny.Seq
            d_2227_v49_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "me"))
            d_2228_v50_: _dafny.Map
            d_2228_v50_ = _dafny.Map({d_2227_v49_: False})
            d_2229_v51_: _dafny.Seq
            d_2229_v51_ = _dafny.SeqWithoutIsStrInference([d_2145_v1_])
            d_2228_v50_ = (d_2228_v50_).set(d_2227_v49_, ((d_2229_v51_)[default__.safeIndex(len(d_2145_v1_), len(d_2229_v51_))]) != (d_2145_v1_))
            d_2230_v52_: str
            d_2230_v52_ = _dafny.CodePoint('s')
            d_2231_v53_: _dafny.Map
            d_2231_v53_ = _dafny.Map({358: self.f4})
            d_2232_v54_: _dafny.MultiSet
            d_2232_v54_ = _dafny.MultiSet([p1, self.f4])
            d_2145_v1_ = default__.fm2(d_2230_v52_, len((_dafny.Map({p1: self.f4})) | (d_2231_v53_)), (d_2232_v54_).cardinality, globalState)
            d_2233_v55_: C0
            nw385_ = C0()
            nw385_.ctor__(((0) - (p1)) * (p1))
            d_2233_v55_ = nw385_
        elif True:
            d_2234_v56_: _dafny.MultiSet
            d_2234_v56_ = _dafny.MultiSet([self.f4])
            if (p3 if not((d_2234_v56_).ispropersubset(d_2234_v56_)) else (self).f8):
                d_2235_v57_: str
                d_2235_v57_ = _dafny.CodePoint('g')
                d_2236_v58_: _dafny.Map
                d_2236_v58_ = _dafny.Map({(d_2235_v57_ if not(p3) else d_2235_v57_): p2})
                d_2237_v59_: _dafny.Seq
                d_2237_v59_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lw"))
                d_2238_v60_: _dafny.MultiSet
                d_2238_v60_ = _dafny.MultiSet([d_2237_v59_, d_2237_v59_])
                d_2236_v58_ = (d_2236_v58_).set(default__.fm27(True, globalState), ((d_2238_v60_).set(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yjggffb")), default__.abs(p1))).issubset(d_2238_v60_))
                d_2239_v61_: _dafny.Array
                nw386_ = _dafny.Array(_dafny.Seq({}), 16)
                d_2239_v61_ = nw386_
                d_2240_v62_: _dafny.Map
                d_2240_v62_ = _dafny.Map({d_2234_v56_: p3})
                d_2241_v63_: C0
                nw387_ = C0()
                nw387_.ctor__(p1)
                d_2241_v63_ = nw387_
                d_2242_v64_: _dafny.Array
                nw388_ = _dafny.Array(None, 16)
                nw388_[int(0)] = d_2241_v63_
                nw388_[int(1)] = d_2241_v63_
                nw388_[int(2)] = d_2241_v63_
                nw388_[int(3)] = d_2241_v63_
                nw388_[int(4)] = d_2241_v63_
                nw388_[int(5)] = d_2241_v63_
                nw388_[int(6)] = d_2241_v63_
                nw388_[int(7)] = d_2241_v63_
                nw388_[int(8)] = d_2241_v63_
                nw388_[int(9)] = d_2241_v63_
                nw388_[int(10)] = d_2241_v63_
                nw388_[int(11)] = d_2241_v63_
                nw388_[int(12)] = d_2241_v63_
                nw388_[int(13)] = d_2241_v63_
                nw388_[int(14)] = d_2241_v63_
                nw388_[int(15)] = d_2241_v63_
                d_2242_v64_ = nw388_
                d_2243_v65_: _dafny.Seq
                d_2243_v65_ = _dafny.SeqWithoutIsStrInference([d_2242_v64_])
                d_2244_v66_: _dafny.Map
                d_2244_v66_ = _dafny.Map({d_2240_v62_: d_2243_v65_})
                index347_ = default__.safeIndex(790, (d_2239_v61_).length(0))
                (d_2239_v61_)[index347_] = ((d_2244_v66_)[d_2240_v62_] if (d_2240_v62_) in (d_2244_v66_) else d_2243_v65_)
                d_2245_v67_: _dafny.Array
                nw389_ = _dafny.Array(False, 27)
                d_2245_v67_ = nw389_
                index348_ = default__.safeIndex(592, (d_2245_v67_).length(0))
                (d_2245_v67_)[index348_] = ((d_2234_v56_).cardinality) > (self.f4)
                index349_ = default__.safeIndex(790, (d_2239_v61_).length(0))
                index350_ = default__.safeIndex(592, (d_2245_v67_).length(0))
                rhs313_ = p3
                rhs314_ = d_2243_v65_
                rhs315_ = p3
                lhs186_ = d_2239_v61_
                lhs187_ = default__.safeIndex(790, (d_2239_v61_).length(0))
                lhs188_ = d_2245_v67_
                lhs189_ = default__.safeIndex(592, (d_2245_v67_).length(0))
                r0 = rhs313_
                lhs186_[lhs187_] = rhs314_
                lhs188_[lhs189_] = rhs315_
                d_2246_v69_: _dafny.Set
                d_2246_v69_ = _dafny.Set({self.f4})
                def iife216_():
                    coll92_ = _dafny.Map()
                    compr_92_: _dafny.Set
                    for compr_92_ in (_dafny.Set({d_2246_v69_})).Elements:
                        d_2247_v68_: _dafny.Set = compr_92_
                        if (d_2247_v68_) in (_dafny.Set({d_2246_v69_})):
                            coll92_[d_2247_v68_] = (_dafny.MultiSet([_dafny.MultiSet([d_2241_v63_.f18])])).cardinality
                    return _dafny.Map(coll92_)
                (self).f4 = (len(iife216_()
                )) * (d_2241_v63_.f18)
                d_2248_v70_: _dafny.Seq
                d_2248_v70_ = _dafny.SeqWithoutIsStrInference([self.f4, self.f4, len((d_2246_v69_) - (d_2246_v69_)), 599, d_2241_v63_.f18])
                d_2248_v70_ = d_2248_v70_
                r0 = p3
            elif True:
                d_2249_v71_: str
                d_2249_v71_ = _dafny.CodePoint('v')
                d_2250_v72_: _dafny.Seq
                d_2250_v72_ = _dafny.SeqWithoutIsStrInference([d_2249_v71_])
                d_2251_v73_: C3
                nw390_ = C3()
                nw390_.ctor__(p1, d_2250_v72_, default__.safeModuloInt(p1, p1))
                d_2251_v73_ = nw390_
                d_2252_v74_: _dafny.Map
                d_2252_v74_ = _dafny.Map({p0: p1})
                rhs316_ = (d_2145_v1_)[default__.safeIndex(self.f4, len(d_2145_v1_))]
                rhs317_ = (self.f4) + (len(d_2252_v74_))
                lhs190_ = self
                r0 = rhs316_
                lhs190_.f4 = rhs317_
                (self).f4 = (default__.fm1(globalState)) - ((d_2251_v73_).f11)
                (self).f4 = ((default__.fm1(globalState)) - ((0) - (self.f4))) * ((0) - (self.f4))
                d_2253_v75_: _dafny.MultiSet
                d_2253_v75_ = _dafny.MultiSet([d_2249_v71_, d_2249_v71_])
                d_2254_v76_: D24
                d_2254_v76_ = D24_DC63(d_2253_v75_)
                pat_let_tv56_ = d_2253_v75_
                d_2255_v77_: _dafny.Map
                def iife217_(_pat_let62_0):
                    def iife218_(d_2256_dt__update__tmp_h3_):
                        def iife219_(_pat_let63_0):
                            def iife220_(d_2257_dt__update_hcf111_h0_):
                                return D24_DC63(d_2257_dt__update_hcf111_h0_)
                            return iife220_(_pat_let63_0)
                        return iife219_(pat_let_tv56_)
                    return iife218_(_pat_let62_0)
                d_2255_v77_ = _dafny.Map({(iife217_(d_2254_v76_)).cf111: ((d_2251_v73_).f11 if False else (d_2251_v73_).f11)})
                d_2255_v77_ = (d_2255_v77_).set(_dafny.MultiSet([_dafny.CodePoint('g'), d_2249_v71_, d_2249_v71_]), (d_2251_v73_).f11)
            d_2258_v78_: _dafny.Array
            nw391_ = _dafny.Array(None, 5)
            nw391_[int(0)] = p2
            nw391_[int(1)] = (self).f8
            nw391_[int(2)] = p3
            nw391_[int(3)] = p2
            nw391_[int(4)] = True
            d_2258_v78_ = nw391_
            index351_ = default__.safeIndex(595, (d_2258_v78_).length(0))
            (d_2258_v78_)[index351_] = False
            d_2259_v79_: _dafny.Array
            def lambda111_(d_2260_p1_):
                def lambda112_(d_2261_i6_):
                    return (d_2261_i6_) * (d_2260_p1_)

                return lambda112_

            init65_ = lambda111_(p1)
            nw392_ = _dafny.Array(None, 22)
            for i0_65_ in range(nw392_.length(0)):
                nw392_[i0_65_] = init65_(i0_65_)
            d_2259_v79_ = nw392_
            d_2262_v80_: _dafny.Map
            d_2262_v80_ = _dafny.Map({not((self).f8): default__.fm1(globalState)})
            d_2263_v81_: _dafny.Seq
            d_2263_v81_ = _dafny.SeqWithoutIsStrInference([self.f4, self.f4, self.f4])
            index352_ = default__.safeIndex(700, (d_2259_v79_).length(0))
            (d_2259_v79_)[index352_] = ((d_2262_v80_)[p3] if (p3) in (d_2262_v80_) else len(d_2263_v81_))
            d_2264_v82_: D0
            d_2264_v82_ = D0_DC1(self.f4, 283, p2)
            pat_let_tv57_ = p1
            pat_let_tv58_ = d_2145_v1_
            d_2265_v83_: D4
            def iife221_(_pat_let64_0):
                def iife222_(d_2266_dt__update__tmp_h4_):
                    def iife223_(_pat_let65_0):
                        def iife224_(d_2267_dt__update_hcf6_h0_):
                            def iife225_(_pat_let66_0):
                                def iife226_(d_2268_dt__update_hcf7_h0_):
                                    return D1_DC4(d_2267_dt__update_hcf6_h0_, d_2268_dt__update_hcf7_h0_, (d_2266_dt__update__tmp_h4_).cf8)
                                return iife226_(_pat_let66_0)
                            return iife225_(pat_let_tv58_)
                        return iife224_(_pat_let65_0)
                    return iife223_(pat_let_tv57_)
                return iife222_(_pat_let64_0)
            d_2265_v83_ = D4_DC13(iife221_(d_2146_v2_), (self).f8, d_2145_v1_)
            index353_ = default__.safeIndex(595, (d_2258_v78_).length(0))
            index354_ = default__.safeIndex(700, (d_2259_v79_).length(0))
            rhs318_ = (self).f8
            rhs319_ = (d_2264_v82_).cf3
            rhs320_ = (0) - (len((d_2265_v83_).cf26))
            lhs191_ = d_2258_v78_
            lhs192_ = default__.safeIndex(595, (d_2258_v78_).length(0))
            lhs193_ = d_2259_v79_
            lhs194_ = default__.safeIndex(700, (d_2259_v79_).length(0))
            lhs191_[lhs192_] = rhs318_
            r0 = rhs319_
            lhs193_[lhs194_] = rhs320_
            d_2269_v84_: _dafny.Array
            nw393_ = _dafny.Array(_dafny.Array(None, 0), 4)
            d_2269_v84_ = nw393_
            index355_ = default__.safeIndex(445, (d_2269_v84_).length(0))
            (d_2269_v84_)[index355_] = d_2258_v78_
            d_2270_v85_: _dafny.Map
            d_2270_v85_ = _dafny.Map({p2: d_2258_v78_})
            index356_ = default__.safeIndex(445, (d_2269_v84_).length(0))
            (d_2269_v84_)[index356_] = ((d_2270_v85_)[(self).f8] if ((self).f8) in (d_2270_v85_) else d_2258_v78_)
            index357_ = default__.safeIndex(700, (d_2259_v79_).length(0))
            rhs321_ = p2
            rhs322_ = p0
            rhs323_ = p1
            rhs324_ = (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_2271_i7_ in range(default__.abs(-232))]))) - ((d_2259_v79_)[default__.safeIndex(700, (d_2259_v79_).length(0))])
            lhs195_ = self
            lhs196_ = d_2259_v79_
            lhs197_ = default__.safeIndex(700, (d_2259_v79_).length(0))
            r0 = rhs321_
            r0 = rhs322_
            lhs195_.f4 = rhs323_
            lhs196_[lhs197_] = rhs324_
            d_2272_v86_: _dafny.Seq
            d_2272_v86_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uvhcfih"))
            d_2273_v87_: _dafny.Map
            d_2273_v87_ = _dafny.Map({p1: d_2272_v86_})
            d_2274_v88_: _dafny.Map
            d_2274_v88_ = _dafny.Map({len(d_2273_v87_): p2})
            d_2275_v89_: C4
            nw394_ = C4()
            nw394_.ctor__((0) - (len(d_2274_v88_)))
            d_2275_v89_ = nw394_
        if p0:
            d_2276_v90_: str
            d_2276_v90_ = _dafny.CodePoint('j')
            d_2277_v91_: _dafny.Seq
            d_2277_v91_ = _dafny.SeqWithoutIsStrInference([d_2276_v90_, d_2276_v90_, d_2276_v90_, d_2276_v90_])
            d_2278_v92_: C3
            nw395_ = C3()
            nw395_.ctor__(p1, ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h'), d_2276_v90_])) + (d_2277_v91_)).set(default__.safeIndex(p1, len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h'), d_2276_v90_])) + (d_2277_v91_))), d_2276_v90_), self.f4)
            d_2278_v92_ = nw395_
            d_2279_v93_: _dafny.Array
            nw396_ = _dafny.Array(None, 29)
            d_2279_v93_ = nw396_
            d_2280_v94_: _dafny.Map
            d_2280_v94_ = _dafny.Map({self.f4: 797})
            d_2281_v95_: C4
            nw397_ = C4()
            nw397_.ctor__((0) - (len(d_2280_v94_)))
            d_2281_v95_ = nw397_
            d_2282_v96_: D25
            d_2282_v96_ = D25_DC65(d_2281_v95_)
            index358_ = default__.safeIndex(867, (d_2279_v93_).length(0))
            (d_2279_v93_)[index358_] = (d_2282_v96_).cf112
            d_2283_v97_: _dafny.Map
            d_2283_v97_ = _dafny.Map({not(p0): (self).f8})
            d_2284_v98_: _dafny.Set
            d_2284_v98_ = _dafny.Set({self.f4})
            d_2285_v99_: _dafny.MultiSet
            d_2285_v99_ = _dafny.MultiSet([len(d_2284_v98_)])
            index359_ = default__.safeIndex(867, (d_2279_v93_).length(0))
            rhs325_ = p1
            rhs326_ = d_2276_v90_
            rhs327_ = (d_2277_v91_) + (d_2277_v91_)
            rhs328_ = not(((len(default__.fm14(len(d_2283_v97_), self.f4, (self).f8, d_2276_v90_, globalState))) + (len(_dafny.Map({((d_2285_v99_)[p1] if (p1) in (d_2285_v99_) else self.f4): not((self).f8)})))) >= ((d_2278_v92_).f11))
            rhs329_ = d_2281_v95_
            lhs198_ = self
            lhs199_ = d_2279_v93_
            lhs200_ = default__.safeIndex(867, (d_2279_v93_).length(0))
            lhs198_.f4 = rhs325_
            d_2276_v90_ = rhs326_
            d_2277_v91_ = rhs327_
            r0 = rhs328_
            lhs199_[lhs200_] = rhs329_
            d_2286_v100_: C5
            nw398_ = C5()
            nw398_.ctor__(p1)
            d_2286_v100_ = nw398_
            r0 = (d_2286_v100_).fm8(globalState)
            d_2287_v101_: _dafny.Seq
            d_2287_v101_ = _dafny.SeqWithoutIsStrInference([(d_2278_v92_).f11])
            d_2288_v102_: _dafny.MultiSet
            d_2288_v102_ = _dafny.MultiSet([p3])
            d_2289_v103_: C1
            nw399_ = C1()
            nw399_.ctor__(self.f4, len((((d_2287_v101_).set(default__.safeIndex(self.f4, len(d_2287_v101_)), p1)).set(default__.safeIndex(p1, len((d_2287_v101_).set(default__.safeIndex(self.f4, len(d_2287_v101_)), p1))), (d_2288_v102_).cardinality)) + (d_2287_v101_)), (p1) * (len(d_2287_v101_)), (_dafny.SeqWithoutIsStrInference([d_2276_v90_ for d_2290_i8_ in range(default__.abs(534))])) < ((d_2278_v92_).f12))
            d_2289_v103_ = nw399_
        elif True:
            d_2291_v104_: str
            d_2291_v104_ = _dafny.CodePoint('c')
            d_2292_v105_: C3
            nw400_ = C3()
            nw400_.ctor__(self.f4, _dafny.SeqWithoutIsStrInference([d_2291_v104_, d_2291_v104_]), p1)
            d_2292_v105_ = nw400_
            (self).f4 = (((_dafny.MultiSet([d_2292_v105_, d_2292_v105_, d_2292_v105_, d_2292_v105_])).cardinality) * (self.f4) if (self).f8 else default__.fm1(globalState))
            d_2293_v106_: _dafny.Set
            d_2293_v106_ = _dafny.Set({(0) - (self.f4)})
            d_2294_v107_: _dafny.Seq
            d_2294_v107_ = _dafny.SeqWithoutIsStrInference([len(d_2293_v106_), (0) - ((d_2292_v105_).f11)])
            (self).f4 = (0) - (((d_2294_v107_)[default__.safeIndex(566, len(d_2294_v107_))]) - ((d_2292_v105_).f11))
            if p3:
                (self).f4 = (d_2292_v105_).f11
                (self).f4 = p1
                d_2295_v108_: C4
                nw401_ = C4()
                nw401_.ctor__(p1)
                d_2295_v108_ = nw401_
                d_2296_v109_: _dafny.Map
                d_2296_v109_ = _dafny.Map({d_2295_v108_: d_2293_v106_})
                d_2296_v109_ = (d_2296_v109_).set(d_2295_v108_, d_2293_v106_)
                d_2297_v110_: _dafny.Array
                nw402_ = _dafny.Array(None, 11)
                nw402_[int(0)] = ((d_2292_v105_).f12 if p3 else default__.fm23(p2, globalState))
                nw402_[int(1)] = _dafny.SeqWithoutIsStrInference([d_2291_v104_ for d_2298_i9_ in range(default__.abs(484))])
                nw402_[int(2)] = (d_2292_v105_).f12
                nw402_[int(3)] = (d_2292_v105_).f12
                nw402_[int(4)] = (d_2292_v105_).f12
                nw402_[int(5)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ajiofkj"))
                nw402_[int(6)] = _dafny.SeqWithoutIsStrInference([d_2291_v104_ for d_2299_i10_ in range(default__.abs(-470))])
                nw402_[int(7)] = (d_2292_v105_).f12
                nw402_[int(8)] = (d_2292_v105_).f12
                nw402_[int(9)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gjgc"))
                nw402_[int(10)] = (D4_DC12((d_2292_v105_).f12)).cf23
                d_2297_v110_ = nw402_
                index360_ = default__.safeIndex(513, (d_2297_v110_).length(0))
                (d_2297_v110_)[index360_] = (d_2292_v105_).f12
                index361_ = default__.safeIndex(513, (d_2297_v110_).length(0))
                (d_2297_v110_)[index361_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "irbhrutbv"))
                d_2300_v111_: _dafny.Map
                d_2300_v111_ = _dafny.Map({False: self.f4})
                d_2301_v112_: C1
                nw403_ = C1()
                nw403_.ctor__(self.f4, ((d_2292_v105_).f11) + (len((d_2292_v105_).f12)), len(d_2300_v111_), p2)
                d_2301_v112_ = nw403_
            elif True:
                (self).f4 = self.f4
                d_2302_v113_: _dafny.Map
                d_2302_v113_ = _dafny.Map({p1: self.f4})
                d_2303_v114_: _dafny.Map
                d_2303_v114_ = _dafny.Map({p2: (self).f8})
                d_2304_v115_: _dafny.Array
                nw404_ = _dafny.Array(None, 12)
                nw404_[int(0)] = p0
                nw404_[int(1)] = (p2) or ((self).f8)
                nw404_[int(2)] = p3
                nw404_[int(3)] = False
                nw404_[int(4)] = p0
                nw404_[int(5)] = (self).f8
                nw404_[int(6)] = (self).fm5((d_2292_v105_).f12, p0, globalState)
                nw404_[int(7)] = (p1) < (len((d_2302_v113_).set(len(d_2303_v114_), (d_2292_v105_).f11)))
                nw404_[int(8)] = p2
                nw404_[int(9)] = (p2 if p0 else p3)
                nw404_[int(10)] = p2
                nw404_[int(11)] = (self).f8
                d_2304_v115_ = nw404_
                index362_ = default__.safeIndex(899, (d_2304_v115_).length(0))
                (d_2304_v115_)[index362_] = p3
                index363_ = default__.safeIndex(899, (d_2304_v115_).length(0))
                (d_2304_v115_)[index363_] = p0
                d_2305_v116_: _dafny.Array
                def lambda113_(d_2306_v105_):
                    def lambda114_(d_2307_i11_):
                        return (d_2307_i11_) + ((d_2306_v105_).f11)

                    return lambda114_

                init66_ = lambda113_(d_2292_v105_)
                nw405_ = _dafny.Array(None, 17)
                for i0_66_ in range(nw405_.length(0)):
                    nw405_[i0_66_] = init66_(i0_66_)
                d_2305_v116_ = nw405_
                index364_ = default__.safeIndex(277, (d_2305_v116_).length(0))
                (d_2305_v116_)[index364_] = (d_2292_v105_).f11
                index365_ = default__.safeIndex(277, (d_2305_v116_).length(0))
                (d_2305_v116_)[index365_] = (self.f4) * ((p1) - ((d_2292_v105_).f11))
                d_2308_v117_: _dafny.Map
                d_2308_v117_ = _dafny.Map({p1: d_2304_v115_})
                d_2309_v118_: _dafny.Seq
                d_2309_v118_ = _dafny.SeqWithoutIsStrInference([d_2304_v115_, d_2304_v115_])
                d_2310_v119_: D3
                d_2310_v119_ = D3_DC8(d_2291_v104_, ((d_2308_v117_)[self.f4] if (self.f4) in (d_2308_v117_) else (d_2309_v118_)[default__.safeIndex(self.f4, len(d_2309_v118_))]))
                pat_let_tv59_ = d_2304_v115_
                d_2311_v120_: _dafny.Array
                nw406_ = _dafny.Array(None, 9)
                nw406_[int(0)] = d_2291_v104_
                nw406_[int(1)] = d_2291_v104_
                def iife227_(_pat_let67_0):
                    def iife228_(d_2312_dt__update__tmp_h5_):
                        def iife229_(_pat_let68_0):
                            def iife230_(d_2313_dt__update_hcf14_h0_):
                                return D3_DC8((d_2312_dt__update__tmp_h5_).cf13, d_2313_dt__update_hcf14_h0_)
                            return iife230_(_pat_let68_0)
                        return iife229_(pat_let_tv59_)
                    return iife228_(_pat_let67_0)
                nw406_[int(2)] = (iife227_(D3_DC8(d_2291_v104_, d_2304_v115_))).cf13
                nw406_[int(3)] = d_2291_v104_
                nw406_[int(4)] = d_2291_v104_
                nw406_[int(5)] = d_2291_v104_
                nw406_[int(6)] = d_2291_v104_
                nw406_[int(7)] = (d_2310_v119_).cf13
                nw406_[int(8)] = d_2291_v104_
                d_2311_v120_ = nw406_
                d_2311_v120_ = d_2311_v120_
                index366_ = default__.safeIndex(277, (d_2305_v116_).length(0))
                (d_2305_v116_)[index366_] = (d_2305_v116_)[default__.safeIndex(277, (d_2305_v116_).length(0))]
            d_2314_v121_: _dafny.Array
            def lambda115_(d_2315_v105_):
                def lambda116_(d_2316_i12_):
                    return (d_2315_v105_).f12

                return lambda116_

            init67_ = lambda115_(d_2292_v105_)
            nw407_ = _dafny.Array(None, 12)
            for i0_67_ in range(nw407_.length(0)):
                nw407_[i0_67_] = init67_(i0_67_)
            d_2314_v121_ = nw407_
            d_2317_v122_: D8
            d_2317_v122_ = D8_DC30(p1, (D8_DC28(d_2314_v121_)).cf53, p2)
            d_2318_v123_: _dafny.Array
            nw408_ = _dafny.Array(None, 10)
            nw408_[int(0)] = p2
            nw408_[int(1)] = p2
            nw408_[int(2)] = p2
            nw408_[int(3)] = True
            nw408_[int(4)] = p0
            nw408_[int(5)] = (self).f8
            nw408_[int(6)] = False
            nw408_[int(7)] = not(p3)
            nw408_[int(8)] = (d_2317_v122_).cf56
            nw408_[int(9)] = p2
            d_2318_v123_ = nw408_
            d_2319_v124_: D23
            d_2319_v124_ = D23_DC61(d_2318_v123_)
            d_2320_v125_: D23
            d_2320_v125_ = D23_DC62(d_2319_v124_)
            pat_let_tv60_ = d_2319_v124_
            def iife231_(_pat_let69_0):
                def iife232_(d_2321_dt__update__tmp_h6_):
                    def iife233_(_pat_let70_0):
                        def iife234_(d_2322_dt__update_hcf110_h0_):
                            return D23_DC62(d_2322_dt__update_hcf110_h0_)
                        return iife234_(_pat_let70_0)
                    return iife233_(pat_let_tv60_)
                return iife232_(_pat_let69_0)
            source22_ = iife231_(d_2320_v125_)
            if source22_.is_DC61:
                d_2323___mcc_h10_ = source22_.cf109
                d_2324_cf109_ = d_2323___mcc_h10_
                d_2325_v126_: _dafny.Array
                nw409_ = _dafny.Array(int(0), 7)
                d_2325_v126_ = nw409_
                index367_ = default__.safeIndex(165, (d_2325_v126_).length(0))
                (d_2325_v126_)[index367_] = p1
                index368_ = default__.safeIndex(165, (d_2325_v126_).length(0))
                (d_2325_v126_)[index368_] = p1
                index369_ = default__.safeIndex(678, (d_2324_cf109_).length(0))
                (d_2324_cf109_)[index369_] = p3
                index370_ = default__.safeIndex(367, (d_2324_cf109_).length(0))
                (d_2324_cf109_)[index370_] = False
                index371_ = default__.safeIndex(678, (d_2324_cf109_).length(0))
                index372_ = default__.safeIndex(367, (d_2324_cf109_).length(0))
                index373_ = default__.safeIndex(165, (d_2325_v126_).length(0))
                rhs330_ = p0
                rhs331_ = (self).f8
                rhs332_ = 757
                rhs333_ = ((d_2292_v105_).f11) >= ((len(d_2293_v106_)) * ((d_2292_v105_).f11))
                rhs334_ = p3
                lhs201_ = d_2324_cf109_
                lhs202_ = default__.safeIndex(678, (d_2324_cf109_).length(0))
                lhs203_ = d_2324_cf109_
                lhs204_ = default__.safeIndex(367, (d_2324_cf109_).length(0))
                lhs205_ = d_2325_v126_
                lhs206_ = default__.safeIndex(165, (d_2325_v126_).length(0))
                lhs201_[lhs202_] = rhs330_
                lhs203_[lhs204_] = rhs331_
                lhs205_[lhs206_] = rhs332_
                r0 = rhs333_
                r0 = rhs334_
                index374_ = default__.safeIndex(165, (d_2325_v126_).length(0))
                (d_2325_v126_)[index374_] = default__.safeModuloInt(len((d_2292_v105_).f12), self.f4)
                d_2326_v127_: _dafny.Array
                def lambda117_(d_2327_cf109_):
                    def lambda118_(d_2328_i13_):
                        return D12_DC37((self).f8, (d_2327_cf109_)[default__.safeIndex(678, (d_2327_cf109_).length(0))], False)

                    return lambda118_

                init68_ = lambda117_(d_2324_cf109_)
                nw410_ = _dafny.Array(None, 27)
                for i0_68_ in range(nw410_.length(0)):
                    nw410_[i0_68_] = init68_(i0_68_)
                d_2326_v127_ = nw410_
                d_2329_v128_: _dafny.Set
                d_2329_v128_ = _dafny.Set({d_2326_v127_})
                d_2330_v129_: _dafny.Map
                d_2330_v129_ = _dafny.Map({len((_dafny.Set({d_2326_v127_, d_2326_v127_, d_2326_v127_})) | (d_2329_v128_)): self.f4})
                d_2331_v130_: _dafny.Map
                d_2331_v130_ = _dafny.Map({(d_2292_v105_).f11: d_2330_v129_})
                index375_ = default__.safeIndex(678, (d_2324_cf109_).length(0))
                index376_ = default__.safeIndex(165, (d_2325_v126_).length(0))
                rhs335_ = d_2325_v126_
                rhs336_ = p2
                rhs337_ = ((d_2331_v130_)[p1] if (p1) in (d_2331_v130_) else _dafny.Map({667: p1}))
                rhs338_ = p1
                lhs207_ = d_2324_cf109_
                lhs208_ = default__.safeIndex(678, (d_2324_cf109_).length(0))
                lhs209_ = d_2325_v126_
                lhs210_ = default__.safeIndex(165, (d_2325_v126_).length(0))
                d_2325_v126_ = rhs335_
                lhs207_[lhs208_] = rhs336_
                d_2330_v129_ = rhs337_
                lhs209_[lhs210_] = rhs338_
            elif source22_.is_DC60:
                d_2332___mcc_h11_ = source22_.cf108
                d_2333_cf108_ = d_2332___mcc_h11_
                d_2334_v131_: _dafny.MultiSet
                d_2334_v131_ = _dafny.MultiSet([p0, False])
                d_2335_v132_: _dafny.MultiSet
                d_2335_v132_ = _dafny.MultiSet([p0, (_dafny.MultiSet([(self).f8, p2, p3, p3, p2])) == (d_2334_v131_), (not((self).f8)) or (p2)])
                d_2336_v133_: _dafny.Map
                d_2336_v133_ = _dafny.Map({d_2291_v104_: self.f4})
                d_2337_v134_: _dafny.MultiSet
                d_2337_v134_ = _dafny.MultiSet([((d_2336_v133_)[d_2291_v104_] if (d_2291_v104_) in (d_2336_v133_) else len((d_2292_v105_).f12)), 892, len((d_2292_v105_).f12)])
                d_2338_v135_: _dafny.Map
                d_2338_v135_ = _dafny.Map({not(p0): self.f4})
                rhs339_ = ((d_2334_v131_)[p3] if (p3) in (d_2334_v131_) else 806)
                rhs340_ = ((d_2337_v134_)[((default__.fm22(p2, self.f4, (d_2292_v105_).f11, p0, globalState)).cf9).cardinality] if (((default__.fm22(p2, self.f4, (d_2292_v105_).f11, p0, globalState)).cf9).cardinality) in (d_2337_v134_) else ((d_2334_v131_)[p0] if (p0) in (d_2334_v131_) else len(d_2338_v135_)))
                lhs211_ = self
                lhs212_ = self
                lhs211_.f4 = rhs339_
                lhs212_.f4 = rhs340_
                d_2339_v136_: _dafny.Map
                d_2339_v136_ = _dafny.Map({(d_2145_v1_) + (_dafny.SeqWithoutIsStrInference([p2, False, p0, p0, p2])): (d_2145_v1_) < (d_2145_v1_)})
                d_2339_v136_ = (d_2339_v136_).set(d_2145_v1_, not((self).f8))
                d_2340_v137_: _dafny.Set
                d_2340_v137_ = _dafny.Set({d_2336_v133_})
                d_2341_v138_: D22
                d_2341_v138_ = D22_DC58(d_2340_v137_)
                pat_let_tv61_ = d_2340_v137_
                def iife235_(_pat_let71_0):
                    def iife236_(d_2342_dt__update__tmp_h7_):
                        def iife237_(_pat_let72_0):
                            def iife238_(d_2343_dt__update_hcf106_h0_):
                                return D22_DC58(d_2343_dt__update_hcf106_h0_)
                            return iife238_(_pat_let72_0)
                        return iife237_(pat_let_tv61_)
                    return iife236_(_pat_let71_0)
                d_2341_v138_ = iife235_(d_2341_v138_)
                (self).f4 = (0) - ((d_2292_v105_).f11)
            elif True:
                d_2344___mcc_h12_ = source22_.cf110
                d_2345_cf110_ = d_2344___mcc_h12_
                d_2346_v139_: C5
                nw411_ = C5()
                nw411_.ctor__((0) - ((0) - ((0) - (p1))))
                d_2346_v139_ = nw411_
                r0 = p2
                d_2347_v140_: _dafny.Map
                d_2347_v140_ = _dafny.Map({_dafny.CodePoint('q'): p1})
                d_2348_v141_: _dafny.Set
                d_2348_v141_ = _dafny.Set({d_2347_v140_})
                index377_ = default__.safeIndex(558, (d_2318_v123_).length(0))
                def iife239_():
                    coll93_ = _dafny.Set()
                    compr_93_: _dafny.Map
                    for compr_93_ in (d_2348_v141_).Elements:
                        d_2349_v142_: _dafny.Map = compr_93_
                        if (d_2349_v142_) in (d_2348_v141_):
                            coll93_ = coll93_.union(_dafny.Set([d_2349_v142_]))
                    return _dafny.Set(coll93_)
                (d_2318_v123_)[index377_] = (d_2348_v141_) == (iife239_()
                )
                index378_ = default__.safeIndex(558, (d_2318_v123_).length(0))
                (d_2318_v123_)[index378_] = p3
                d_2350_v143_: _dafny.MultiSet
                d_2350_v143_ = _dafny.MultiSet([(d_2318_v123_)[default__.safeIndex(558, (d_2318_v123_).length(0))], (self).f8, p2])
                d_2351_v144_: _dafny.Map
                d_2351_v144_ = _dafny.Map({(d_2350_v143_).isdisjoint(d_2350_v143_): d_2145_v1_})
                d_2352_v145_: _dafny.Seq
                d_2352_v145_ = _dafny.SeqWithoutIsStrInference([d_2351_v144_])
                d_2351_v144_ = ((d_2351_v144_) | (_dafny.Map({p0: _dafny.SeqWithoutIsStrInference([p2])})) if p3 else (d_2352_v145_)[default__.safeIndex(p1, len(d_2352_v145_))])
            d_2353_v146_: _dafny.MultiSet
            d_2353_v146_ = _dafny.MultiSet([True, not(p3)])
            d_2354_v147_: _dafny.Set
            d_2354_v147_ = _dafny.Set({d_2353_v146_, d_2353_v146_})
            d_2355_v148_: D0
            d_2355_v148_ = D0_DC0(d_2294_v107_)
            d_2356_v149_: _dafny.Array
            nw412_ = _dafny.Array(None, 15)
            nw412_[int(0)] = d_2294_v107_
            nw412_[int(1)] = _dafny.SeqWithoutIsStrInference([(d_2292_v105_).f11 for d_2357_i14_ in range(default__.abs(786))])
            nw412_[int(2)] = d_2294_v107_
            nw412_[int(3)] = _dafny.SeqWithoutIsStrInference([(d_2292_v105_).f11 for d_2358_i15_ in range(default__.abs(782))])
            nw412_[int(4)] = (d_2294_v107_) + (d_2294_v107_)
            nw412_[int(5)] = default__.fm21(p1, p2, (self).f8, d_2354_v147_, globalState)
            nw412_[int(6)] = (d_2294_v107_) + (d_2294_v107_)
            nw412_[int(7)] = (d_2355_v148_).cf0
            nw412_[int(8)] = d_2294_v107_
            nw412_[int(9)] = _dafny.SeqWithoutIsStrInference([p1])
            nw412_[int(10)] = (d_2294_v107_) + (d_2294_v107_)
            nw412_[int(11)] = d_2294_v107_
            nw412_[int(12)] = d_2294_v107_
            nw412_[int(13)] = (_dafny.SeqWithoutIsStrInference([(d_2292_v105_).f11])) + (_dafny.SeqWithoutIsStrInference([(d_2292_v105_).f11 for d_2359_i16_ in range(default__.abs(337))]))
            nw412_[int(14)] = d_2294_v107_
            d_2356_v149_ = nw412_
            index379_ = default__.safeIndex(356, (d_2356_v149_).length(0))
            (d_2356_v149_)[index379_] = _dafny.SeqWithoutIsStrInference([43, p1])
            d_2360_v150_: _dafny.Map
            d_2360_v150_ = _dafny.Map({True: (_dafny.MultiSet(d_2294_v107_)).cardinality})
            index380_ = default__.safeIndex(356, (d_2356_v149_).length(0))
            rhs341_ = default__.safeModuloInt(self.f4, len((d_2360_v150_) | (d_2360_v150_)))
            rhs342_ = p3
            rhs343_ = (d_2292_v105_).f11
            rhs344_ = ((d_2294_v107_) + (d_2294_v107_)).set(default__.safeIndex((0) - (p1), len((d_2294_v107_) + (d_2294_v107_))), (d_2292_v105_).f11)
            lhs213_ = self
            lhs214_ = self
            lhs215_ = d_2356_v149_
            lhs216_ = default__.safeIndex(356, (d_2356_v149_).length(0))
            lhs213_.f4 = rhs341_
            r0 = rhs342_
            lhs214_.f4 = rhs343_
            lhs215_[lhs216_] = rhs344_
        d_2361_v151_: _dafny.Array
        nw413_ = _dafny.Array(False, 6)
        d_2361_v151_ = nw413_
        d_2362_v152_: _dafny.Map
        d_2362_v152_ = _dafny.Map({self.f4: d_2361_v151_})
        d_2363_v153_: _dafny.Array
        nw414_ = _dafny.Array(None, 13)
        nw414_[int(0)] = d_2361_v151_
        nw414_[int(1)] = d_2361_v151_
        nw414_[int(2)] = d_2361_v151_
        nw414_[int(3)] = d_2361_v151_
        nw414_[int(4)] = d_2361_v151_
        nw414_[int(5)] = d_2361_v151_
        nw414_[int(6)] = d_2361_v151_
        nw414_[int(7)] = d_2361_v151_
        nw414_[int(8)] = d_2361_v151_
        nw414_[int(9)] = d_2361_v151_
        nw414_[int(10)] = d_2361_v151_
        nw414_[int(11)] = d_2361_v151_
        nw414_[int(12)] = (d_2361_v151_ if p3 else ((d_2362_v152_)[p1] if (p1) in (d_2362_v152_) else d_2361_v151_))
        d_2363_v153_ = nw414_
        index381_ = default__.safeIndex(643, (d_2363_v153_).length(0))
        (d_2363_v153_)[index381_] = d_2361_v151_
        index382_ = default__.safeIndex(643, (d_2363_v153_).length(0))
        (d_2363_v153_)[index382_] = d_2361_v151_
        if p3:
            d_2364_v155_: _dafny.Map
            def iife240_():
                coll94_ = _dafny.Map()
                compr_94_: int
                for compr_94_ in _dafny.IntegerRange(381, 809):
                    d_2365_v154_: int = compr_94_
                    if ((381) <= (d_2365_v154_)) and ((d_2365_v154_) < (809)):
                        coll94_[(d_2365_v154_) * (self.f4)] = _dafny.CodePoint('l')
                return _dafny.Map(coll94_)
            d_2364_v155_ = iife240_()
            
            source23_ = d_2364_v155_
            d_2366___mcc_h13_ = source23_
            d_2367_cf76_ = d_2366___mcc_h13_
            d_2368_v156_: str
            d_2368_v156_ = _dafny.CodePoint('c')
            rhs345_ = self.f4
            rhs346_ = _dafny.CodePoint('m')
            rhs347_ = p3
            lhs217_ = self
            lhs217_.f4 = rhs345_
            d_2368_v156_ = rhs346_
            r0 = rhs347_
            r0 = not(p3)
            d_2369_v157_: _dafny.Array
            def lambda119_(d_2370_v1_):
                def lambda120_(d_2371_i17_):
                    return (d_2371_i17_) + (len(d_2370_v1_))

                return lambda120_

            init69_ = lambda119_(d_2145_v1_)
            nw415_ = _dafny.Array(None, 11)
            for i0_69_ in range(nw415_.length(0)):
                nw415_[i0_69_] = init69_(i0_69_)
            d_2369_v157_ = nw415_
            d_2372_v158_: _dafny.Map
            d_2372_v158_ = _dafny.Map({default__.fm0(self.f4, p1, globalState): (_dafny.SeqWithoutIsStrInference([d_2368_v156_ for d_2373_i18_ in range(default__.abs(-514))])).set(default__.safeIndex(p1, len(_dafny.SeqWithoutIsStrInference([d_2368_v156_ for d_2374_i18_ in range(default__.abs(-514))]))), d_2368_v156_)})
            index383_ = default__.safeIndex(147, (d_2369_v157_).length(0))
            (d_2369_v157_)[index383_] = len(d_2372_v158_)
            d_2375_v159_: _dafny.Map
            d_2375_v159_ = _dafny.Map({False: d_2368_v156_})
            d_2376_v160_: D3
            d_2376_v160_ = D3_DC10(len(d_2145_v1_), p1, (self).f8, p1, p2)
            d_2377_v161_: _dafny.Map
            d_2377_v161_ = _dafny.Map({self.f4: d_2376_v160_})
            d_2378_v162_: _dafny.Seq
            d_2378_v162_ = _dafny.SeqWithoutIsStrInference([(d_2377_v161_).set(self.f4, D3_DC10(p1, (0) - (p1), (D2_DC6(861, False)).cf11, self.f4, p0))])
            d_2379_v163_: D3
            d_2379_v163_ = D3_DC10(len(d_2378_v162_), self.f4, p0, p1, p3)
            index384_ = default__.safeIndex(147, (d_2369_v157_).length(0))
            rhs348_ = p1
            rhs349_ = self.f4
            rhs350_ = d_2368_v156_
            rhs351_ = len(default__.fm2(((d_2375_v159_)[p3] if (p3) in (d_2375_v159_) else default__.fm27((self).f8, globalState)), ((d_2376_v160_).cf20 if p3 else p1), (self.f4) - (p1), globalState))
            lhs218_ = d_2369_v157_
            lhs219_ = default__.safeIndex(147, (d_2369_v157_).length(0))
            lhs220_ = self
            lhs221_ = self
            lhs218_[lhs219_] = rhs348_
            lhs220_.f4 = rhs349_
            d_2368_v156_ = rhs350_
            lhs221_.f4 = rhs351_
            d_2380_v164_: _dafny.Seq
            d_2380_v164_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gss"))
            d_2381_v165_: _dafny.Seq
            d_2381_v165_ = _dafny.SeqWithoutIsStrInference([p1, (0) - (len(d_2380_v164_))])
            d_2382_v166_: _dafny.Seq
            d_2382_v166_ = _dafny.SeqWithoutIsStrInference([(d_2381_v165_)[default__.safeIndex(self.f4, len(d_2381_v165_))], self.f4, (d_2369_v157_)[default__.safeIndex(147, (d_2369_v157_).length(0))], p1, p1])
            index385_ = default__.safeIndex(147, (d_2369_v157_).length(0))
            (d_2369_v157_)[index385_] = (p1) * (len(d_2382_v166_))
            d_2383_v167_: _dafny.Seq
            d_2383_v167_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e')])
            d_2384_v168_: _dafny.Seq
            d_2384_v168_ = _dafny.SeqWithoutIsStrInference([737, p1])
            d_2385_v169_: _dafny.Map
            d_2385_v169_ = _dafny.Map({self.f4: (0) - (len(d_2384_v168_))})
            d_2386_v170_: C3
            nw416_ = C3()
            nw416_.ctor__(default__.safeModuloInt(p1, self.f4), d_2383_v167_, ((d_2385_v169_)[self.f4] if (self.f4) in (d_2385_v169_) else p1))
            d_2386_v170_ = nw416_
            d_2387_v171_: _dafny.Array
            nw417_ = _dafny.Array(_dafny.CodePoint('D'), 8)
            d_2387_v171_ = nw417_
            d_2388_v172_: _dafny.Map
            d_2388_v172_ = _dafny.Map({d_2387_v171_: d_2361_v151_})
            d_2388_v172_ = d_2388_v172_
            d_2389_v173_: _dafny.Array
            nw418_ = _dafny.Array(int(0), 10)
            d_2389_v173_ = nw418_
            nw419_ = _dafny.Array(None, 8)
            nw419_[int(0)] = self.f4
            nw419_[int(1)] = (0) - (len((d_2384_v168_) + (d_2384_v168_)))
            nw419_[int(2)] = default__.safeModuloInt(p1, self.f4)
            nw419_[int(3)] = p1
            nw419_[int(4)] = len(d_2145_v1_)
            nw419_[int(5)] = (0) - (self.f4)
            nw419_[int(6)] = default__.safeDivisionInt((0) - (len((d_2384_v168_).set(default__.safeIndex(p1, len(d_2384_v168_)), (0) - (self.f4)))), (d_2386_v170_).f11)
            nw419_[int(7)] = p1
            d_2389_v173_ = nw419_
            (self).f4 = (p1) + ((d_2386_v170_).f11)
        elif True:
            r0 = (p3) and ((p3) == (p0))
            d_2390_v174_: _dafny.Seq
            d_2390_v174_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oyoltakt"))
            d_2391_v175_: _dafny.Seq
            d_2391_v175_ = _dafny.SeqWithoutIsStrInference([self.f4, self.f4, p1])
            d_2392_v176_: str
            d_2392_v176_ = _dafny.CodePoint('l')
            d_2393_v177_: D4
            d_2393_v177_ = D4_DC12(d_2390_v174_)
            d_2394_v178_: _dafny.Array
            nw420_ = _dafny.Array(None, 12)
            nw420_[int(0)] = (d_2390_v174_).set(default__.safeIndex(len(d_2391_v175_), len(d_2390_v174_)), d_2392_v176_)
            nw420_[int(1)] = (d_2393_v177_).cf23
            nw420_[int(2)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sol"))
            nw420_[int(3)] = d_2390_v174_
            nw420_[int(4)] = d_2390_v174_
            nw420_[int(5)] = d_2390_v174_
            nw420_[int(6)] = default__.fm23(p3, globalState)
            nw420_[int(7)] = _dafny.SeqWithoutIsStrInference([d_2392_v176_ for d_2395_i19_ in range(default__.abs(315))])
            nw420_[int(8)] = d_2390_v174_
            nw420_[int(9)] = (D4_DC12(d_2390_v174_)).cf23
            nw420_[int(10)] = d_2390_v174_
            nw420_[int(11)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_2396_i20_ in range(default__.abs(651))])
            d_2394_v178_ = nw420_
            d_2397_v179_: D8
            d_2397_v179_ = D8_DC30(self.f4, d_2394_v178_, p0)
            d_2398_v180_: D14
            d_2398_v180_ = D14_DC41(d_2144_v0_)
            rhs352_ = default__.safeModuloInt((self.f4) - (self.f4), (self.f4) - ((0) - (self.f4)))
            rhs353_ = ((d_2390_v174_).set(default__.safeIndex((0) - (p1), len(d_2390_v174_)), d_2392_v176_) if p0 else d_2390_v174_)
            rhs354_ = d_2397_v179_
            rhs355_ = len(((d_2398_v180_).cf71).intersection(d_2144_v0_))
            rhs356_ = default__.fm1(globalState)
            lhs222_ = self
            lhs223_ = self
            lhs224_ = self
            lhs222_.f4 = rhs352_
            d_2390_v174_ = rhs353_
            d_2397_v179_ = rhs354_
            lhs223_.f4 = rhs355_
            lhs224_.f4 = rhs356_
            d_2399_v181_: _dafny.Array
            def lambda121_(d_2400_i21_):
                return default__.safeModuloInt(d_2400_i21_, self.f4)

            init70_ = lambda121_
            nw421_ = _dafny.Array(None, 3)
            for i0_70_ in range(nw421_.length(0)):
                nw421_[i0_70_] = init70_(i0_70_)
            d_2399_v181_ = nw421_
            d_2399_v181_ = (d_2399_v181_ if (self).f8 else d_2399_v181_)
            d_2401_v182_: _dafny.Set
            d_2401_v182_ = _dafny.Set({p1})
            d_2402_v183_: _dafny.Map
            d_2402_v183_ = _dafny.Map({p3: default__.fm23(p2, globalState)})
            d_2401_v182_ = default__.fm16((self).f8, _dafny.SeqWithoutIsStrInference([d_2390_v174_, ((d_2402_v183_)[p3] if (p3) in (d_2402_v183_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tutfxc"))), d_2390_v174_]), (self).f8, (p1) * (self.f4), globalState)
            if p2:
                d_2403_v184_: _dafny.Array
                def lambda122_(d_2404_v176_):
                    def lambda123_(d_2405_i22_):
                        return (d_2404_v176_ if False else d_2404_v176_)

                    return lambda123_

                init71_ = lambda122_(d_2392_v176_)
                nw422_ = _dafny.Array(None, 14)
                for i0_71_ in range(nw422_.length(0)):
                    nw422_[i0_71_] = init71_(i0_71_)
                d_2403_v184_ = nw422_
                d_2406_v186_: D6
                def iife241_():
                    coll95_ = _dafny.Map()
                    compr_95_: int
                    for compr_95_ in (d_2391_v175_).Elements:
                        d_2407_v185_: int = compr_95_
                        if (d_2407_v185_) in (d_2391_v175_):
                            coll95_[(d_2407_v185_) - (self.f4)] = (self).f8
                    return _dafny.Map(coll95_)
                d_2406_v186_ = D6_DC20(iife241_()
)
                d_2408_v187_: _dafny.MultiSet
                d_2408_v187_ = _dafny.MultiSet([len((d_2406_v186_).cf40)])
                d_2409_v188_: _dafny.Map
                d_2409_v188_ = _dafny.Map({(d_2408_v187_).intersection(d_2408_v187_): d_2403_v184_})
                d_2403_v184_ = ((d_2409_v188_)[d_2408_v187_] if (d_2408_v187_) in (d_2409_v188_) else d_2403_v184_)
                d_2410_v189_: _dafny.Array
                nw423_ = _dafny.Array(_dafny.Set({}), 4)
                d_2410_v189_ = nw423_
                index386_ = default__.safeIndex(574, (d_2410_v189_).length(0))
                (d_2410_v189_)[index386_] = _dafny.Set({p2, p0})
                index387_ = default__.safeIndex(574, (d_2410_v189_).length(0))
                (d_2410_v189_)[index387_] = _dafny.Set({p2})
                d_2411_v190_: D3
                d_2411_v190_ = D3_DC9(p2, p1)
                index388_ = default__.safeIndex(357, (d_2361_v151_).length(0))
                (d_2361_v151_)[index388_] = p2
                d_2412_v191_: _dafny.Seq
                d_2412_v191_ = _dafny.SeqWithoutIsStrInference([(d_2410_v189_)[default__.safeIndex(574, (d_2410_v189_).length(0))], (d_2410_v189_)[default__.safeIndex(574, (d_2410_v189_).length(0))]])
                d_2413_v192_: _dafny.Map
                d_2413_v192_ = _dafny.Map({d_2145_v1_: (self).f8})
                index389_ = default__.safeIndex(357, (d_2361_v151_).length(0))
                rhs357_ = p3
                rhs358_ = D3_DC9(p0, len(d_2145_v1_))
                rhs359_ = p1
                rhs360_ = ((self).f8 if p3 else (self.f4) < (len((d_2412_v191_)[default__.safeIndex(p1, len(d_2412_v191_))])))
                rhs361_ = ((d_2413_v192_)[(d_2145_v1_) + (d_2145_v1_)] if ((d_2145_v1_) + (d_2145_v1_)) in (d_2413_v192_) else True)
                lhs225_ = self
                lhs226_ = d_2361_v151_
                lhs227_ = default__.safeIndex(357, (d_2361_v151_).length(0))
                r0 = rhs357_
                d_2411_v190_ = rhs358_
                lhs225_.f4 = rhs359_
                lhs226_[lhs227_] = rhs360_
                r0 = rhs361_
                d_2144_v0_ = (d_2410_v189_)[default__.safeIndex(574, (d_2410_v189_).length(0))]
                d_2414_v194_: _dafny.Array
                nw424_ = _dafny.Array(None, 8)
                nw424_[int(0)] = (d_2401_v182_).intersection(d_2401_v182_)
                nw424_[int(1)] = (d_2401_v182_).intersection(d_2401_v182_)
                nw424_[int(2)] = d_2401_v182_
                nw424_[int(3)] = d_2401_v182_
                def iife242_():
                    coll96_ = _dafny.Set()
                    compr_96_: int
                    for compr_96_ in _dafny.IntegerRange(111, 930):
                        d_2415_v193_: int = compr_96_
                        if ((111) <= (d_2415_v193_)) and ((d_2415_v193_) < (930)):
                            coll96_ = coll96_.union(_dafny.Set([(d_2415_v193_) + (self.f4)]))
                    return _dafny.Set(coll96_)
                nw424_[int(4)] = (iife242_()
                ) - (d_2401_v182_)
                nw424_[int(5)] = (d_2401_v182_).intersection(d_2401_v182_)
                nw424_[int(6)] = (d_2401_v182_).intersection(d_2401_v182_)
                nw424_[int(7)] = d_2401_v182_
                d_2414_v194_ = nw424_
                index390_ = default__.safeIndex(604, (d_2414_v194_).length(0))
                (d_2414_v194_)[index390_] = _dafny.Set({self.f4, p1, default__.fm1(globalState), self.f4, (d_2391_v175_)[default__.safeIndex(p1, len(d_2391_v175_))]})
                index391_ = default__.safeIndex(604, (d_2414_v194_).length(0))
                (d_2414_v194_)[index391_] = ((d_2401_v182_).intersection(d_2401_v182_)).intersection(d_2401_v182_)
            elif True:
                d_2401_v182_ = _dafny.Set({self.f4, default__.safeDivisionInt(-823, self.f4)})
                index392_ = default__.safeIndex(373, (d_2361_v151_).length(0))
                (d_2361_v151_)[index392_] = p0
                index393_ = default__.safeIndex(373, (d_2361_v151_).length(0))
                (d_2361_v151_)[index393_] = p3
                (self).f4 = default__.safeModuloInt(self.f4, default__.fm1(globalState))
                index394_ = default__.safeIndex(373, (d_2361_v151_).length(0))
                (d_2361_v151_)[index394_] = not(default__.fm0(default__.safeDivisionInt(p1, 654), (p1) - (p1), globalState))
                d_2416_v195_: C2
                nw425_ = C2()
                nw425_.ctor__(p3, p1)
                d_2416_v195_ = nw425_
                d_2417_v196_: _dafny.Seq
                d_2417_v196_ = _dafny.SeqWithoutIsStrInference([(_dafny.Map({d_2416_v195_: False})).set(d_2416_v195_, p3)])
                d_2418_v197_: D13
                d_2418_v197_ = D13_DC39(d_2417_v196_)
                d_2418_v197_ = (D13_DC39(d_2417_v196_) if p0 else d_2418_v197_)
        d_2419_v198_: D23
        d_2419_v198_ = D23_DC61(d_2361_v151_)
        source24_ = d_2419_v198_
        if source24_.is_DC61:
            d_2420___mcc_h14_ = source24_.cf109
            d_2421_cf109_ = d_2420___mcc_h14_
            d_2422_v199_: D11
            d_2422_v199_ = D11_DC35(self.f4, p3, (self).f8)
            d_2423_v200_: _dafny.Map
            d_2423_v200_ = _dafny.Map({default__.safeModuloInt((d_2422_v199_).cf61, self.f4): (0) - (self.f4)})
            d_2423_v200_ = (d_2423_v200_).set(990, default__.safeModuloInt(p1, p1))
            index395_ = default__.safeIndex(173, (d_2361_v151_).length(0))
            (d_2361_v151_)[index395_] = p2
            index396_ = default__.safeIndex(173, (d_2361_v151_).length(0))
            (d_2361_v151_)[index396_] = p2
            d_2424_v201_: _dafny.Map
            d_2424_v201_ = _dafny.Map({p1: d_2145_v1_})
            d_2425_v202_: D12
            d_2425_v202_ = D12_DC36(d_2424_v201_)
            source25_ = d_2425_v202_
            if source25_.is_DC37:
                d_2426___mcc_h17_ = source25_.cf65
                d_2427___mcc_h18_ = source25_.cf66
                d_2428___mcc_h19_ = source25_.cf67
                d_2429_cf67_ = d_2428___mcc_h19_
                d_2430_cf66_ = d_2427___mcc_h18_
                d_2431_cf65_ = d_2426___mcc_h17_
                d_2432_v203_: _dafny.Seq
                d_2432_v203_ = _dafny.SeqWithoutIsStrInference([d_2421_cf109_, (d_2363_v153_)[default__.safeIndex(643, (d_2363_v153_).length(0))]])
                d_2433_v204_: _dafny.Array
                nw426_ = _dafny.Array(None, 24)
                nw426_[int(0)] = p1
                nw426_[int(1)] = len(d_2432_v203_)
                nw426_[int(2)] = (0) - (-358)
                nw426_[int(3)] = self.f4
                nw426_[int(4)] = self.f4
                nw426_[int(5)] = p1
                nw426_[int(6)] = self.f4
                nw426_[int(7)] = p1
                nw426_[int(8)] = p1
                nw426_[int(9)] = 572
                nw426_[int(10)] = self.f4
                nw426_[int(11)] = p1
                nw426_[int(12)] = p1
                nw426_[int(13)] = self.f4
                nw426_[int(14)] = p1
                nw426_[int(15)] = p1
                nw426_[int(16)] = p1
                nw426_[int(17)] = p1
                nw426_[int(18)] = p1
                nw426_[int(19)] = 597
                nw426_[int(20)] = self.f4
                nw426_[int(21)] = 676
                nw426_[int(22)] = p1
                nw426_[int(23)] = self.f4
                d_2433_v204_ = nw426_
                d_2434_v205_: _dafny.Seq
                d_2434_v205_ = _dafny.SeqWithoutIsStrInference([d_2433_v204_])
                (self).f4 = (len((d_2434_v205_).set(default__.safeIndex(p1, len(d_2434_v205_)), d_2433_v204_)) if not(p0) else 525)
                d_2435_v206_: C5
                nw427_ = C5()
                nw427_.ctor__(925)
                d_2435_v206_ = nw427_
                d_2421_cf109_ = (d_2363_v153_)[default__.safeIndex(643, (d_2363_v153_).length(0))]
                d_2436_v207_: str
                d_2436_v207_ = _dafny.CodePoint('s')
                d_2437_v208_: _dafny.Seq
                d_2437_v208_ = _dafny.SeqWithoutIsStrInference([d_2436_v207_])
                d_2438_v209_: C3
                nw428_ = C3()
                nw428_.ctor__(self.f4, d_2437_v208_, p1)
                d_2438_v209_ = nw428_
            elif source25_.is_DC38:
                d_2439___mcc_h20_ = source25_.cf68
                d_2440_cf68_ = d_2439___mcc_h20_
                d_2441_v210_: _dafny.Seq
                d_2441_v210_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kqfagp"))
                d_2442_v211_: _dafny.MultiSet
                d_2442_v211_ = _dafny.MultiSet([not((len(d_2441_v210_)) == (p1)), True])
                d_2440_cf68_ = ((d_2442_v211_)[p0] if (p0) in (d_2442_v211_) else p1)
                d_2441_v210_ = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_2443_i23_ in range(default__.abs(392))])) + ((d_2441_v210_) + (d_2441_v210_))
                d_2444_v212_: _dafny.Map
                d_2444_v212_ = _dafny.Map({False: p3})
                d_2445_v213_: C3
                nw429_ = C3()
                nw429_.ctor__(p1, d_2441_v210_, (len(d_2444_v212_)) * (d_2440_cf68_))
                d_2445_v213_ = nw429_
                d_2446_v214_: _dafny.Array
                def lambda124_(d_2447_v1_, d_2448_v151_, d_2449_v213_):
                    def lambda125_(d_2450_i24_):
                        return (_dafny.SeqWithoutIsStrInference([len(d_2447_v1_), 290, len(_dafny.Map({(d_2448_v151_)[default__.safeIndex(173, (d_2448_v151_).length(0))]: (d_2449_v213_).f12}))])) + (_dafny.SeqWithoutIsStrInference([(d_2449_v213_).f11, len(_dafny.SeqWithoutIsStrInference([True]))]))

                    return lambda125_

                init72_ = lambda124_(d_2145_v1_, d_2361_v151_, d_2445_v213_)
                nw430_ = _dafny.Array(None, 19)
                for i0_72_ in range(nw430_.length(0)):
                    nw430_[i0_72_] = init72_(i0_72_)
                d_2446_v214_ = nw430_
                d_2451_v215_: _dafny.Array
                def lambda126_(d_2452_v1_):
                    def lambda127_(d_2453_i25_):
                        return d_2452_v1_

                    return lambda127_

                init73_ = lambda126_(d_2145_v1_)
                nw431_ = _dafny.Array(None, 16)
                for i0_73_ in range(nw431_.length(0)):
                    nw431_[i0_73_] = init73_(i0_73_)
                d_2451_v215_ = nw431_
                rhs362_ = d_2446_v214_
                rhs363_ = d_2451_v215_
                d_2446_v214_ = rhs362_
                d_2451_v215_ = rhs363_
            elif True:
                d_2454___mcc_h21_ = source25_.cf64
                d_2455_cf64_ = d_2454___mcc_h21_
                d_2456_v216_: _dafny.MultiSet
                d_2456_v216_ = _dafny.MultiSet([default__.fm1(globalState)])
                d_2457_v217_: _dafny.Map
                d_2457_v217_ = _dafny.Map({(d_2456_v216_).cardinality: (default__.fm0(p1, self.f4, globalState)) and (p2)})
                d_2457_v217_ = (d_2457_v217_).set(self.f4, not (p0) or (p3))
                (self).f4 = p1
                index397_ = default__.safeIndex(173, (d_2361_v151_).length(0))
                (d_2361_v151_)[index397_] = p3
                d_2458_v218_: _dafny.Map
                d_2458_v218_ = _dafny.Map({(d_2361_v151_)[default__.safeIndex(173, (d_2361_v151_).length(0))]: True})
                r0 = not(default__.fm0((0) - ((0) - (p1)), (p1) * (len(d_2458_v218_)), globalState))
            d_2459_v219_: D3
            d_2459_v219_ = D3_DC8(_dafny.CodePoint('p'), d_2421_cf109_)
            d_2460_v220_: str
            d_2460_v220_ = _dafny.CodePoint('b')
            d_2461_v221_: _dafny.Seq
            d_2461_v221_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "anjxjeynf"))
            d_2462_v222_: _dafny.Array
            nw432_ = _dafny.Array(None, 19)
            nw432_[int(0)] = d_2461_v221_
            nw432_[int(1)] = d_2461_v221_
            nw432_[int(2)] = d_2461_v221_
            nw432_[int(3)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wwdemlm"))
            nw432_[int(4)] = d_2461_v221_
            nw432_[int(5)] = _dafny.SeqWithoutIsStrInference([d_2460_v220_ for d_2463_i26_ in range(default__.abs(686))])
            nw432_[int(6)] = d_2461_v221_
            nw432_[int(7)] = d_2461_v221_
            nw432_[int(8)] = d_2461_v221_
            nw432_[int(9)] = _dafny.SeqWithoutIsStrInference([d_2460_v220_ for d_2464_i27_ in range(default__.abs(147))])
            nw432_[int(10)] = d_2461_v221_
            nw432_[int(11)] = _dafny.SeqWithoutIsStrInference([d_2460_v220_ for d_2465_i28_ in range(default__.abs(933))])
            nw432_[int(12)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lufi"))
            nw432_[int(13)] = d_2461_v221_
            nw432_[int(14)] = d_2461_v221_
            nw432_[int(15)] = d_2461_v221_
            nw432_[int(16)] = d_2461_v221_
            nw432_[int(17)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wigod"))
            nw432_[int(18)] = d_2461_v221_
            d_2462_v222_ = nw432_
            d_2466_v223_: D8
            d_2466_v223_ = D8_DC30(self.f4, d_2462_v222_, (self).f8)
            d_2467_v224_: _dafny.Array
            nw433_ = _dafny.Array(None, 28)
            nw433_[int(0)] = (d_2459_v219_).cf13
            nw433_[int(1)] = d_2460_v220_
            nw433_[int(2)] = d_2460_v220_
            nw433_[int(3)] = d_2460_v220_
            nw433_[int(4)] = (d_2460_v220_ if (d_2361_v151_)[default__.safeIndex(173, (d_2361_v151_).length(0))] else d_2460_v220_)
            nw433_[int(5)] = _dafny.CodePoint('u')
            nw433_[int(6)] = (d_2461_v221_)[default__.safeIndex((d_2466_v223_).cf54, len(d_2461_v221_))]
            nw433_[int(7)] = default__.fm27((d_2361_v151_)[default__.safeIndex(173, (d_2361_v151_).length(0))], globalState)
            nw433_[int(8)] = d_2460_v220_
            nw433_[int(9)] = _dafny.CodePoint('r')
            nw433_[int(10)] = d_2460_v220_
            nw433_[int(11)] = d_2460_v220_
            nw433_[int(12)] = d_2460_v220_
            nw433_[int(13)] = d_2460_v220_
            nw433_[int(14)] = d_2460_v220_
            nw433_[int(15)] = d_2460_v220_
            nw433_[int(16)] = (d_2461_v221_)[default__.safeIndex((0) - (p1), len(d_2461_v221_))]
            nw433_[int(17)] = d_2460_v220_
            nw433_[int(18)] = d_2460_v220_
            nw433_[int(19)] = d_2460_v220_
            nw433_[int(20)] = d_2460_v220_
            nw433_[int(21)] = d_2460_v220_
            nw433_[int(22)] = d_2460_v220_
            nw433_[int(23)] = d_2460_v220_
            nw433_[int(24)] = _dafny.CodePoint('c')
            nw433_[int(25)] = _dafny.CodePoint('r')
            nw433_[int(26)] = d_2460_v220_
            nw433_[int(27)] = d_2460_v220_
            d_2467_v224_ = nw433_
            index398_ = default__.safeIndex(205, (d_2467_v224_).length(0))
            (d_2467_v224_)[index398_] = d_2460_v220_
            index399_ = default__.safeIndex(205, (d_2467_v224_).length(0))
            (d_2467_v224_)[index399_] = d_2460_v220_
        elif source24_.is_DC60:
            d_2468___mcc_h15_ = source24_.cf108
            d_2469_cf108_ = d_2468___mcc_h15_
            d_2470_v225_: _dafny.Seq
            d_2470_v225_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fdypxqjbb"))
            d_2471_v226_: _dafny.Seq
            d_2471_v226_ = _dafny.SeqWithoutIsStrInference([self.f4])
            d_2472_v227_: _dafny.Array
            nw434_ = _dafny.Array(None, 24)
            nw434_[int(0)] = self.f4
            nw434_[int(1)] = len(d_2470_v225_)
            nw434_[int(2)] = p1
            nw434_[int(3)] = p1
            nw434_[int(4)] = self.f4
            nw434_[int(5)] = p1
            nw434_[int(6)] = self.f4
            nw434_[int(7)] = p1
            nw434_[int(8)] = len(d_2470_v225_)
            nw434_[int(9)] = p1
            nw434_[int(10)] = p1
            nw434_[int(11)] = p1
            nw434_[int(12)] = len(d_2470_v225_)
            nw434_[int(13)] = (d_2471_v226_)[default__.safeIndex(self.f4, len(d_2471_v226_))]
            nw434_[int(14)] = self.f4
            nw434_[int(15)] = len(_dafny.SeqWithoutIsStrInference([self.f4 for d_2473_i29_ in range(default__.abs(824))]))
            nw434_[int(16)] = self.f4
            nw434_[int(17)] = self.f4
            nw434_[int(18)] = self.f4
            nw434_[int(19)] = (0) - ((_dafny.MultiSet([p1])).cardinality)
            nw434_[int(20)] = p1
            nw434_[int(21)] = p1
            nw434_[int(22)] = p1
            nw434_[int(23)] = p1
            d_2472_v227_ = nw434_
            d_2474_v228_: _dafny.MultiSet
            d_2474_v228_ = _dafny.MultiSet([d_2472_v227_])
            rhs364_ = d_2474_v228_
            rhs365_ = (default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hsxqsgq"))), p1)) * (default__.safeModuloInt(self.f4, -333))
            rhs366_ = p0
            rhs367_ = self.f4
            lhs228_ = self
            lhs229_ = self
            d_2474_v228_ = rhs364_
            lhs228_.f4 = rhs365_
            r0 = rhs366_
            lhs229_.f4 = rhs367_
            d_2475_v229_: _dafny.Map
            d_2475_v229_ = _dafny.Map({True: p1})
            (self).f4 = ((self.f4 if (self).fm5(d_2470_v225_, p0, globalState) else self.f4)) - (((d_2475_v229_)[p2] if (p2) in (d_2475_v229_) else len(d_2471_v226_)))
            index400_ = default__.safeIndex(398, (d_2472_v227_).length(0))
            (d_2472_v227_)[index400_] = (p1) - (p1)
            index401_ = default__.safeIndex(398, (d_2472_v227_).length(0))
            rhs368_ = (0) - (self.f4)
            rhs369_ = ((len(d_2470_v225_)) - (p1) if (self).f8 else default__.fm1(globalState))
            rhs370_ = p3
            lhs230_ = d_2472_v227_
            lhs231_ = default__.safeIndex(398, (d_2472_v227_).length(0))
            lhs232_ = self
            lhs230_[lhs231_] = rhs368_
            lhs232_.f4 = rhs369_
            r0 = rhs370_
            index402_ = default__.safeIndex(398, (d_2472_v227_).length(0))
            (d_2472_v227_)[index402_] = self.f4
        elif True:
            d_2476___mcc_h16_ = source24_.cf110
            d_2477_cf110_ = d_2476___mcc_h16_
            d_2478_v230_: _dafny.Map
            d_2478_v230_ = _dafny.Map({self.f4: True})
            d_2479_v231_: D13
            d_2479_v231_ = D13_DC40(((d_2478_v230_)[len(d_2145_v1_)] if (len(d_2145_v1_)) in (d_2478_v230_) else p2))
            source26_ = d_2479_v231_
            if source26_.is_DC40:
                d_2480___mcc_h22_ = source26_.cf70
                d_2481_cf70_ = d_2480___mcc_h22_
                d_2482_v232_: _dafny.MultiSet
                d_2482_v232_ = _dafny.MultiSet([self.f4, 479])
                d_2483_v233_: D19
                d_2483_v233_ = D19_DC52((0) - (len(_dafny.Map({(d_2482_v232_).cardinality: p0}))), p1, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dmxup"))), p3)
                d_2484_v234_: _dafny.Map
                d_2484_v234_ = _dafny.Map({self.f4: self.f4})
                (self).f4 = default__.safeDivisionInt((d_2483_v233_).cf94, (len(d_2484_v234_) if (self).f8 else p1))
                d_2485_v235_: str
                d_2485_v235_ = _dafny.CodePoint('e')
                d_2486_v236_: D3
                d_2486_v236_ = D3_DC8(d_2485_v235_, (d_2363_v153_)[default__.safeIndex(643, (d_2363_v153_).length(0))])
                d_2485_v235_ = (d_2486_v236_).cf13
                nw435_ = _dafny.Array(False, 28)
                d_2361_v151_ = nw435_
                d_2487_v237_: _dafny.Array
                nw436_ = _dafny.Array(None, 6)
                nw436_[int(0)] = p1
                nw436_[int(1)] = p1
                nw436_[int(2)] = p1
                nw436_[int(3)] = default__.fm1(globalState)
                nw436_[int(4)] = self.f4
                nw436_[int(5)] = p1
                d_2487_v237_ = nw436_
                index403_ = default__.safeIndex(501, (d_2487_v237_).length(0))
                (d_2487_v237_)[index403_] = self.f4
                index404_ = default__.safeIndex(501, (d_2487_v237_).length(0))
                (d_2487_v237_)[index404_] = p1
            elif True:
                d_2488___mcc_h23_ = source26_.cf69
                d_2489_cf69_ = d_2488___mcc_h23_
                d_2490_v238_: D5
                d_2490_v238_ = D5_DC16(True, p3, p3, self.f4, (self).f8)
                d_2490_v238_ = d_2490_v238_
                r0 = ((p1) + (p1)) >= (self.f4)
                d_2491_v239_: _dafny.Array
                nw437_ = _dafny.Array(_dafny.Seq({}), 11)
                d_2491_v239_ = nw437_
                d_2492_v240_: _dafny.Map
                d_2492_v240_ = _dafny.Map({p3: p1})
                d_2493_v241_: D1
                d_2493_v241_ = D1_DC3(d_2492_v240_)
                d_2494_v242_: D1
                d_2494_v242_ = D1_DC3((d_2493_v241_).cf5)
                d_2495_v243_: _dafny.Seq
                d_2495_v243_ = _dafny.SeqWithoutIsStrInference([d_2493_v241_])
                d_2496_v244_: _dafny.Map
                d_2496_v244_ = _dafny.Map({True: d_2495_v243_})
                index405_ = default__.safeIndex(690, (d_2491_v239_).length(0))
                (d_2491_v239_)[index405_] = ((d_2496_v244_)[p3] if (p3) in (d_2496_v244_) else _dafny.SeqWithoutIsStrInference([d_2494_v242_]))
                index406_ = default__.safeIndex(690, (d_2491_v239_).length(0))
                (d_2491_v239_)[index406_] = (d_2495_v243_) + (d_2495_v243_)
                d_2497_v245_: _dafny.Seq
                d_2497_v245_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fxu"))
                d_2498_v246_: D2
                d_2498_v246_ = D2_DC6(len(d_2497_v245_), not(True))
                d_2499_v247_: str
                d_2499_v247_ = _dafny.CodePoint('g')
                d_2500_v248_: _dafny.Seq
                d_2500_v248_ = _dafny.SeqWithoutIsStrInference([((d_2497_v245_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_2501_i30_ in range(default__.abs(541))]))).set(default__.safeIndex(self.f4, len((d_2497_v245_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_2502_i30_ in range(default__.abs(541))])))), d_2499_v247_), (_dafny.SeqWithoutIsStrInference([d_2499_v247_ for d_2503_i31_ in range(default__.abs(795))])) + (d_2497_v245_), d_2497_v245_])
                pat_let_tv62_ = p0
                def iife243_(_pat_let73_0):
                    def iife244_(d_2504_dt__update__tmp_h8_):
                        def iife245_(_pat_let74_0):
                            def iife246_(d_2505_dt__update_hcf11_h0_):
                                return D2_DC6((d_2504_dt__update__tmp_h8_).cf10, d_2505_dt__update_hcf11_h0_)
                            return iife246_(_pat_let74_0)
                        return iife245_(pat_let_tv62_)
                    return iife244_(_pat_let73_0)
                rhs371_ = iife243_(d_2498_v246_)
                rhs372_ = d_2500_v248_
                rhs373_ = p1
                lhs233_ = self
                d_2498_v246_ = rhs371_
                d_2500_v248_ = rhs372_
                lhs233_.f4 = rhs373_
            d_2506_v249_: _dafny.Seq
            d_2506_v249_ = _dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Map({p1: p1})))])
            d_2507_v250_: _dafny.Seq
            d_2507_v250_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fxfq"))
            d_2508_v251_: _dafny.Array
            nw438_ = _dafny.Array(None, 27)
            nw438_[int(0)] = p1
            nw438_[int(1)] = default__.fm1(globalState)
            nw438_[int(2)] = p1
            nw438_[int(3)] = 722
            nw438_[int(4)] = -664
            nw438_[int(5)] = p1
            nw438_[int(6)] = len((d_2506_v249_).set(default__.safeIndex(p1, len(d_2506_v249_)), p1))
            nw438_[int(7)] = self.f4
            nw438_[int(8)] = 527
            nw438_[int(9)] = self.f4
            nw438_[int(10)] = p1
            nw438_[int(11)] = self.f4
            nw438_[int(12)] = p1
            nw438_[int(13)] = 430
            nw438_[int(14)] = self.f4
            nw438_[int(15)] = self.f4
            nw438_[int(16)] = -919
            nw438_[int(17)] = len(d_2507_v250_)
            nw438_[int(18)] = self.f4
            nw438_[int(19)] = self.f4
            nw438_[int(20)] = len(d_2145_v1_)
            nw438_[int(21)] = len(d_2507_v250_)
            nw438_[int(22)] = len(d_2506_v249_)
            nw438_[int(23)] = p1
            nw438_[int(24)] = len(d_2478_v230_)
            nw438_[int(25)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xedjbel")))
            nw438_[int(26)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xllxo")))
            d_2508_v251_ = nw438_
            d_2509_v252_: _dafny.Array
            nw439_ = _dafny.Array(None, 12)
            nw439_[int(0)] = d_2508_v251_
            nw439_[int(1)] = d_2508_v251_
            nw439_[int(2)] = d_2508_v251_
            nw439_[int(3)] = d_2508_v251_
            nw439_[int(4)] = d_2508_v251_
            nw439_[int(5)] = d_2508_v251_
            nw439_[int(6)] = d_2508_v251_
            nw439_[int(7)] = d_2508_v251_
            nw439_[int(8)] = d_2508_v251_
            nw439_[int(9)] = d_2508_v251_
            nw439_[int(10)] = d_2508_v251_
            nw439_[int(11)] = d_2508_v251_
            d_2509_v252_ = nw439_
            d_2510_v253_: _dafny.Seq
            d_2510_v253_ = _dafny.SeqWithoutIsStrInference([d_2509_v252_, d_2509_v252_, d_2509_v252_, d_2509_v252_])
            d_2511_v254_: D11
            d_2511_v254_ = D11_DC34((d_2510_v253_)[default__.safeIndex((0) - (-451), len(d_2510_v253_))])
            pat_let_tv63_ = d_2509_v252_
            def iife247_(_pat_let75_0):
                def iife248_(d_2512_dt__update__tmp_h9_):
                    def iife249_(_pat_let76_0):
                        def iife250_(d_2513_dt__update_hcf60_h0_):
                            return D11_DC34(d_2513_dt__update_hcf60_h0_)
                        return iife250_(_pat_let76_0)
                    return iife249_(pat_let_tv63_)
                return iife248_(_pat_let75_0)
            source27_ = iife247_(d_2511_v254_)
            if source27_.is_DC35:
                d_2514___mcc_h24_ = source27_.cf61
                d_2515___mcc_h25_ = source27_.cf62
                d_2516___mcc_h26_ = source27_.cf63
                d_2517_cf63_ = d_2516___mcc_h26_
                d_2518_cf62_ = d_2515___mcc_h25_
                d_2519_cf61_ = d_2514___mcc_h24_
                d_2520_v255_: _dafny.Array
                nw440_ = _dafny.Array(_dafny.MultiSet({}), 20)
                d_2520_v255_ = nw440_
                d_2521_v256_: str
                d_2521_v256_ = _dafny.CodePoint('p')
                index407_ = default__.safeIndex(604, (d_2520_v255_).length(0))
                (d_2520_v255_)[index407_] = _dafny.MultiSet(default__.fm14((0) - ((D19_DC51(self.f4, len(d_2145_v1_))).cf91), p1, p2, d_2521_v256_, globalState))
                d_2522_v257_: _dafny.Map
                d_2522_v257_ = _dafny.Map({d_2521_v256_: p3})
                d_2523_v258_: _dafny.Map
                d_2523_v258_ = _dafny.Map({(d_2522_v257_).set(d_2521_v256_, False): d_2519_cf61_})
                d_2524_v259_: _dafny.MultiSet
                d_2524_v259_ = _dafny.MultiSet([len(d_2523_v258_)])
                d_2525_v260_: D2
                d_2525_v260_ = D2_DC6(default__.fm1(globalState), d_2518_cf62_)
                d_2526_v261_: D2
                d_2526_v261_ = D2_DC6(default__.fm1(globalState), (d_2525_v260_).cf11)
                index408_ = default__.safeIndex(604, (d_2520_v255_).length(0))
                (d_2520_v255_)[index408_] = (d_2524_v259_).set((d_2526_v261_).cf10, default__.abs((len(d_2507_v250_)) - (self.f4)))
                d_2527_v262_: C0
                nw441_ = C0()
                nw441_.ctor__(714)
                d_2527_v262_ = nw441_
                d_2528_v263_: D5
                d_2528_v263_ = D5_DC18(d_2527_v262_, True, p2)
                d_2529_v264_: D20
                d_2529_v264_ = D20_DC55(d_2507_v250_, (d_2528_v263_).cf36, d_2527_v262_.f18)
                d_2530_v265_: _dafny.MultiSet
                d_2530_v265_ = _dafny.MultiSet([((d_2529_v264_).cf101).set(default__.safeIndex(self.f4, len((d_2529_v264_).cf101)), d_2521_v256_), d_2507_v250_, default__.fm23(d_2518_cf62_, globalState), d_2507_v250_])
                d_2530_v265_ = d_2530_v265_
                d_2531_v266_: _dafny.Set
                d_2531_v266_ = _dafny.Set({p1, self.f4})
                d_2532_v267_: _dafny.Seq
                d_2532_v267_ = _dafny.SeqWithoutIsStrInference([d_2507_v250_, d_2507_v250_])
                d_2531_v266_ = default__.fm16(not(p0), d_2532_v267_, True, d_2527_v262_.f18, globalState)
                d_2519_cf61_ = default__.safeDivisionInt((d_2519_cf61_) + (p1), (d_2506_v249_)[default__.safeIndex(self.f4, len(d_2506_v249_))])
            elif True:
                d_2533___mcc_h27_ = source27_.cf60
                d_2534_cf60_ = d_2533___mcc_h27_
                index409_ = default__.safeIndex(643, (d_2363_v153_).length(0))
                (d_2363_v153_)[index409_] = (d_2363_v153_)[default__.safeIndex(643, (d_2363_v153_).length(0))]
                rhs374_ = (self).fm5((d_2507_v250_) + (d_2507_v250_), (d_2145_v1_) != (d_2145_v1_), globalState)
                rhs375_ = p2
                rhs376_ = self.f4
                lhs234_ = self
                r0 = rhs374_
                r0 = rhs375_
                lhs234_.f4 = rhs376_
                d_2535_v268_: str
                d_2535_v268_ = _dafny.CodePoint('t')
                d_2535_v268_ = d_2535_v268_
                (self).f4 = (0) - (default__.safeDivisionInt(self.f4, 518))
            d_2536_v269_: _dafny.Map
            d_2536_v269_ = _dafny.Map({(self.f4) != (p1): not(p0)})
            if ((d_2536_v269_)[p2] if (p2) in (d_2536_v269_) else (d_2506_v249_) != (d_2506_v249_)):
                d_2537_v270_: _dafny.Map
                d_2537_v270_ = _dafny.Map({(d_2506_v249_)[default__.safeIndex(self.f4, len(d_2506_v249_))]: p1})
                d_2538_v271_: _dafny.Seq
                d_2538_v271_ = _dafny.SeqWithoutIsStrInference([d_2537_v270_, d_2537_v270_])
                d_2539_v272_: _dafny.Array
                nw442_ = _dafny.Array(None, 21)
                nw442_[int(0)] = (d_2538_v271_) + (d_2538_v271_)
                nw442_[int(1)] = d_2538_v271_
                nw442_[int(2)] = _dafny.SeqWithoutIsStrInference([d_2537_v270_ for d_2540_i32_ in range(default__.abs(-227))])
                nw442_[int(3)] = d_2538_v271_
                nw442_[int(4)] = d_2538_v271_
                nw442_[int(5)] = d_2538_v271_
                nw442_[int(6)] = _dafny.SeqWithoutIsStrInference([(_dafny.Map({p1: self.f4})).set(p1, len(_dafny.SeqWithoutIsStrInference([p2]))) for d_2541_i33_ in range(default__.abs(593))])
                nw442_[int(7)] = _dafny.SeqWithoutIsStrInference([d_2537_v270_ for d_2542_i34_ in range(default__.abs(113))])
                nw442_[int(8)] = d_2538_v271_
                nw442_[int(9)] = _dafny.SeqWithoutIsStrInference([_dafny.Map({p1: len(d_2507_v250_)}), d_2537_v270_, d_2537_v270_])
                nw442_[int(10)] = _dafny.SeqWithoutIsStrInference([d_2537_v270_, default__.fm45(p3, p2, globalState), (d_2537_v270_).set(p1, self.f4), d_2537_v270_, d_2537_v270_])
                nw442_[int(11)] = (d_2538_v271_ if p2 else d_2538_v271_)
                nw442_[int(12)] = d_2538_v271_
                nw442_[int(13)] = d_2538_v271_
                nw442_[int(14)] = _dafny.SeqWithoutIsStrInference([d_2537_v270_ for d_2543_i35_ in range(default__.abs(-581))])
                nw442_[int(15)] = (d_2538_v271_ if p0 else d_2538_v271_)
                nw442_[int(16)] = d_2538_v271_
                nw442_[int(17)] = d_2538_v271_
                nw442_[int(18)] = d_2538_v271_
                nw442_[int(19)] = (d_2538_v271_) + (_dafny.SeqWithoutIsStrInference([d_2537_v270_ for d_2544_i36_ in range(default__.abs(15))]))
                nw442_[int(20)] = d_2538_v271_
                d_2539_v272_ = nw442_
                d_2545_v273_: _dafny.Map
                d_2545_v273_ = _dafny.Map({self.f4: d_2538_v271_})
                index410_ = default__.safeIndex(616, (d_2539_v272_).length(0))
                (d_2539_v272_)[index410_] = ((d_2545_v273_)[p1] if (p1) in (d_2545_v273_) else d_2538_v271_)
                index411_ = default__.safeIndex(616, (d_2539_v272_).length(0))
                (d_2539_v272_)[index411_] = d_2538_v271_
                d_2546_v274_: _dafny.MultiSet
                d_2546_v274_ = _dafny.MultiSet([p1])
                d_2547_v275_: _dafny.Map
                d_2547_v275_ = _dafny.Map({d_2546_v274_: d_2508_v251_})
                d_2548_v276_: _dafny.Map
                d_2548_v276_ = _dafny.Map({default__.safeDivisionInt(self.f4, p1): (d_2547_v275_) | (d_2547_v275_)})
                d_2548_v276_ = (d_2548_v276_).set(p1, d_2547_v275_)
                r0 = ((self).f8 if True else p2)
                (self).f4 = self.f4
                d_2362_v152_ = (d_2362_v152_).set(p1, (d_2419_v198_).cf109)
            elif True:
                d_2549_v277_: C0
                nw443_ = C0()
                nw443_.ctor__(p1)
                d_2549_v277_ = nw443_
                d_2550_v278_: D18
                d_2550_v278_ = D18_DC48(p0, p0, d_2146_v2_, d_2549_v277_)
                d_2551_v279_: _dafny.Map
                d_2551_v279_ = _dafny.Map({(len(default__.fm40(globalState))) * (p1): d_2550_v278_})
                d_2551_v279_ = (d_2551_v279_).set(p1, d_2550_v278_)
                index412_ = default__.safeIndex(859, (d_2509_v252_).length(0))
                (d_2509_v252_)[index412_] = d_2508_v251_
                d_2552_v280_: D24
                d_2552_v280_ = D24_DC64()
                index413_ = default__.safeIndex(859, (d_2509_v252_).length(0))
                rhs377_ = p1
                rhs378_ = d_2508_v251_
                rhs379_ = d_2552_v280_
                lhs235_ = self
                lhs236_ = d_2509_v252_
                lhs237_ = default__.safeIndex(859, (d_2509_v252_).length(0))
                lhs235_.f4 = rhs377_
                lhs236_[lhs237_] = rhs378_
                d_2552_v280_ = rhs379_
                d_2553_v281_: C2
                nw444_ = C2()
                nw444_.ctor__((self).f8, (default__.fm1(globalState)) * (p1))
                d_2553_v281_ = nw444_
                d_2554_v282_: _dafny.Array
                nw445_ = _dafny.Array(None, 26)
                nw445_[int(0)] = d_2507_v250_
                nw445_[int(1)] = d_2507_v250_
                nw445_[int(2)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gct"))
                nw445_[int(3)] = d_2507_v250_
                nw445_[int(4)] = d_2507_v250_
                nw445_[int(5)] = d_2507_v250_
                nw445_[int(6)] = d_2507_v250_
                nw445_[int(7)] = d_2507_v250_
                nw445_[int(8)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_2555_i37_ in range(default__.abs(24))])
                nw445_[int(9)] = d_2507_v250_
                nw445_[int(10)] = d_2507_v250_
                nw445_[int(11)] = default__.fm23((self).f8, globalState)
                nw445_[int(12)] = d_2507_v250_
                nw445_[int(13)] = d_2507_v250_
                nw445_[int(14)] = d_2507_v250_
                nw445_[int(15)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_2556_i38_ in range(default__.abs(903))])
                nw445_[int(16)] = d_2507_v250_
                nw445_[int(17)] = d_2507_v250_
                nw445_[int(18)] = d_2507_v250_
                nw445_[int(19)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ebbox"))
                nw445_[int(20)] = d_2507_v250_
                nw445_[int(21)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_2557_i39_ in range(default__.abs(-388))])
                nw445_[int(22)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dlam"))
                nw445_[int(23)] = d_2507_v250_
                nw445_[int(24)] = d_2507_v250_
                nw445_[int(25)] = d_2507_v250_
                d_2554_v282_ = nw445_
                d_2558_v283_: D8
                d_2558_v283_ = D8_DC28(d_2554_v282_)
                d_2559_v284_: _dafny.Array
                nw446_ = _dafny.Array(None, 16)
                nw446_[int(0)] = D8_DC28(d_2554_v282_)
                nw446_[int(1)] = d_2558_v283_
                nw446_[int(2)] = d_2558_v283_
                nw446_[int(3)] = d_2558_v283_
                nw446_[int(4)] = d_2558_v283_
                nw446_[int(5)] = d_2558_v283_
                nw446_[int(6)] = d_2558_v283_
                nw446_[int(7)] = d_2558_v283_
                nw446_[int(8)] = D8_DC28(d_2554_v282_)
                nw446_[int(9)] = d_2558_v283_
                nw446_[int(10)] = D8_DC28(d_2554_v282_)
                nw446_[int(11)] = d_2558_v283_
                nw446_[int(12)] = d_2558_v283_
                nw446_[int(13)] = d_2558_v283_
                nw446_[int(14)] = d_2558_v283_
                nw446_[int(15)] = d_2558_v283_
                d_2559_v284_ = nw446_
                d_2560_v285_: D19
                d_2560_v285_ = D19_DC50(d_2559_v284_)
                d_2561_v286_: _dafny.Set
                d_2561_v286_ = _dafny.Set({d_2560_v285_, d_2560_v285_})
                r0 = not(not(((_dafny.Set({d_2560_v285_})) | (d_2561_v286_)).issubset(d_2561_v286_)))
                d_2562_v287_: C2
                nw447_ = C2()
                nw447_.ctor__(not((self).f8), (d_2553_v281_).f14)
                d_2562_v287_ = nw447_
                nw448_ = C2()
                nw448_.ctor__(p3, (self.f4) - (d_2549_v277_.f18))
                d_2562_v287_ = nw448_
            d_2563_v288_: D14
            d_2563_v288_ = D14_DC42(907, (self).f8, d_2144_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bwdjwbmd"))))
            if not((d_2563_v288_).cf73):
                (self).f4 = p1
                d_2564_v289_: D7
                d_2564_v289_ = D7_DC25(((self).f8) or (p3), ((self).f8 if p0 else p2), (0) - (p1))
                d_2565_v290_: _dafny.Seq
                d_2565_v290_ = _dafny.SeqWithoutIsStrInference([d_2507_v250_])
                d_2566_v291_: _dafny.MultiSet
                d_2566_v291_ = _dafny.MultiSet([d_2507_v250_, d_2507_v250_, d_2507_v250_, (d_2565_v290_)[default__.safeIndex(347, len(d_2565_v290_))], _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_2567_i40_ in range(default__.abs(76))])])
                d_2568_v292_: _dafny.Array
                d_2568_v292_ = d_2363_v153_
                rhs380_ = (d_2568_v292_)
                rhs381_ = ((self.f4) <= (152)) or (p3)
                rhs382_ = d_2564_v289_
                rhs383_ = (d_2566_v291_) | (d_2566_v291_)
                rhs384_ = d_2508_v251_
                d_2363_v153_ = rhs380_
                r0 = rhs381_
                d_2564_v289_ = rhs382_
                d_2566_v291_ = rhs383_
                d_2508_v251_ = rhs384_
                r0 = (default__.fm0((0) - (self.f4), self.f4, globalState)) or (not((p3) or (p2)))
                index414_ = default__.safeIndex(683, (d_2509_v252_).length(0))
                (d_2509_v252_)[index414_] = d_2508_v251_
                index415_ = default__.safeIndex(213, (d_2508_v251_).length(0))
                (d_2508_v251_)[index415_] = (p1) - (p1)
                index416_ = default__.safeIndex(683, (d_2509_v252_).length(0))
                index417_ = default__.safeIndex(213, (d_2508_v251_).length(0))
                rhs385_ = (default__.safeDivisionInt(p1, p1) if (self).fm5(d_2507_v250_, (self).f8, globalState) else default__.fm1(globalState))
                rhs386_ = d_2508_v251_
                rhs387_ = self.f4
                lhs238_ = self
                lhs239_ = d_2509_v252_
                lhs240_ = default__.safeIndex(683, (d_2509_v252_).length(0))
                lhs241_ = d_2508_v251_
                lhs242_ = default__.safeIndex(213, (d_2508_v251_).length(0))
                lhs238_.f4 = rhs385_
                lhs239_[lhs240_] = rhs386_
                lhs241_[lhs242_] = rhs387_
                d_2569_v293_: _dafny.Map
                d_2569_v293_ = _dafny.Map({p2: d_2507_v250_})
                d_2569_v293_ = (d_2569_v293_).set((self).f8, d_2507_v250_)
            elif True:
                (self).f4 = 386
                d_2570_v294_: _dafny.Set
                d_2570_v294_ = _dafny.Set({self.f4, p1})
                d_2571_v295_: _dafny.MultiSet
                d_2571_v295_ = _dafny.MultiSet([p1])
                d_2572_v296_: _dafny.Map
                d_2572_v296_ = _dafny.Map({d_2508_v251_: d_2571_v295_})
                d_2573_v297_: _dafny.Set
                d_2573_v297_ = _dafny.Set({d_2570_v294_, _dafny.Set({((((d_2572_v296_)[d_2508_v251_] if (d_2508_v251_) in (d_2572_v296_) else d_2571_v295_)).set(self.f4, default__.abs(p1))).cardinality})})
                (self).f4 = len((d_2573_v297_) - (d_2573_v297_))
                d_2574_v298_: str
                d_2574_v298_ = _dafny.CodePoint('v')
                rhs388_ = (d_2507_v250_)[default__.safeIndex(len(_dafny.Set({p0})), len(d_2507_v250_))]
                rhs389_ = not(False)
                rhs390_ = (((d_2507_v250_).set(default__.safeIndex(len(d_2507_v250_), len(d_2507_v250_)), d_2574_v298_)) + (d_2507_v250_)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uvuta")))
                d_2574_v298_ = rhs388_
                r0 = rhs389_
                d_2507_v250_ = rhs390_
                d_2575_v299_: D24
                d_2575_v299_ = D24_DC64()
                d_2575_v299_ = d_2575_v299_
                d_2507_v250_ = d_2507_v250_
        r0 = (self).f8
        return r0

    def m4(self, p0, p1, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r2: _dafny.Array = _dafny.Array(None, 0)
        d_2576_v0_: _dafny.Map
        d_2576_v0_ = _dafny.Map({(self).f8: p1})
        d_2577_v1_: _dafny.Seq
        d_2577_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lmp"))
        d_2578_v2_: D7
        d_2578_v2_ = D7_DC25((self).f8, ((d_2576_v0_)[(self).f8] if ((self).f8) in (d_2576_v0_) else False), len(d_2577_v1_))
        source28_ = (d_2578_v2_ if not(not((self).f8)) else d_2578_v2_)
        if source28_.is_DC25:
            d_2579___mcc_h0_ = source28_.cf45
            d_2580___mcc_h1_ = source28_.cf46
            d_2581___mcc_h2_ = source28_.cf47
            d_2582_cf47_ = d_2581___mcc_h2_
            d_2583_cf46_ = d_2580___mcc_h1_
            d_2584_cf45_ = d_2579___mcc_h0_
            d_2585_v3_: _dafny.Map
            d_2585_v3_ = _dafny.Map({self.f4: (p0) > (d_2582_cf47_)})
            d_2586_v4_: _dafny.Set
            d_2586_v4_ = _dafny.Set({True})
            d_2585_v3_ = (d_2585_v3_).set(len((d_2586_v4_) - (d_2586_v4_)), d_2583_cf46_)
            d_2587_v5_: _dafny.Seq
            d_2587_v5_ = _dafny.SeqWithoutIsStrInference([(self).f8, d_2583_cf46_])
            d_2588_v6_: _dafny.Set
            d_2588_v6_ = _dafny.Set({self.f4})
            d_2589_v7_: _dafny.Seq
            d_2589_v7_ = _dafny.SeqWithoutIsStrInference([936, self.f4, p0, len(d_2588_v6_), 797])
            d_2590_v8_: _dafny.Seq
            d_2590_v8_ = _dafny.SeqWithoutIsStrInference([(len(d_2587_v5_)) <= ((0) - (len(d_2589_v7_))), False, p1, (d_2582_cf47_) > (-625), (self).fm5(d_2577_v1_, True, globalState)])
            if (d_2587_v5_)[default__.safeIndex((D3_DC10(self.f4, (0) - (d_2582_cf47_), False, p0, (self).f8)).cf20, len(d_2587_v5_))]:
                d_2591_v9_: _dafny.Array
                nw449_ = _dafny.Array(_dafny.Set({}), 10)
                d_2591_v9_ = nw449_
                d_2591_v9_ = d_2591_v9_
                d_2592_v10_: _dafny.Array
                def lambda128_(d_2593_cf46_):
                    def lambda129_(d_2594_i0_):
                        return d_2593_cf46_

                    return lambda129_

                init74_ = lambda128_(d_2583_cf46_)
                nw450_ = _dafny.Array(None, 20)
                for i0_74_ in range(nw450_.length(0)):
                    nw450_[i0_74_] = init74_(i0_74_)
                d_2592_v10_ = nw450_
                rhs391_ = d_2582_cf47_
                rhs392_ = d_2592_v10_
                rhs393_ = (d_2584_cf45_) == (d_2583_cf46_)
                rhs394_ = d_2582_cf47_
                lhs243_ = self
                lhs243_.f4 = rhs391_
                d_2592_v10_ = rhs392_
                d_2583_cf46_ = rhs393_
                d_2582_cf47_ = rhs394_
                d_2592_v10_ = d_2592_v10_
                d_2595_v11_: C4
                nw451_ = C4()
                nw451_.ctor__(self.f4)
                d_2595_v11_ = nw451_
                d_2596_v12_: _dafny.MultiSet
                d_2596_v12_ = _dafny.MultiSet([d_2584_cf45_, d_2583_cf46_])
                d_2597_v13_: D5
                d_2597_v13_ = D5_DC17(d_2596_v12_, not(True))
                rhs395_ = (d_2588_v6_).intersection(_dafny.Set({self.f4}))
                rhs396_ = (d_2597_v13_).cf35
                d_2588_v6_ = rhs395_
                d_2583_cf46_ = rhs396_
            elif True:
                d_2583_cf46_ = not((self).f8)
                (self).f4 = self.f4
                d_2598_v14_: C4
                nw452_ = C4()
                nw452_.ctor__(d_2582_cf47_)
                d_2598_v14_ = nw452_
                d_2599_v15_: _dafny.Map
                d_2599_v15_ = _dafny.Map({not((d_2586_v4_).issubset(d_2586_v4_)): default__.safeModuloInt(self.f4, d_2582_cf47_)})
                d_2599_v15_ = (d_2599_v15_).set(not (d_2583_cf46_) or (p1), self.f4)
                d_2584_cf45_ = p1
            (self).f4 = p0
            d_2600_v16_: _dafny.Array
            nw453_ = _dafny.Array(int(0), 9)
            d_2600_v16_ = nw453_
            d_2601_v17_: _dafny.Map
            d_2601_v17_ = _dafny.Map({p0: d_2600_v16_})
            d_2602_v18_: int
            d_2603_v19_: int
            d_2604_v20_: _dafny.Array
            d_2605_v21_: bool
            out55_: int
            out56_: int
            out57_: _dafny.Array
            out58_: bool
            out55_, out56_, out57_, out58_ = default__.m0(d_2585_v3_, (d_2589_v7_ if p1 else d_2589_v7_), d_2601_v17_, globalState)
            d_2602_v18_ = out55_
            d_2603_v19_ = out56_
            d_2604_v20_ = out57_
            d_2605_v21_ = out58_
        elif source28_.is_DC26:
            d_2606___mcc_h3_ = source28_.cf48
            d_2607___mcc_h4_ = source28_.cf49
            d_2608_cf49_ = d_2607___mcc_h4_
            d_2609_cf48_ = d_2606___mcc_h3_
            d_2610_v22_: _dafny.Map
            d_2610_v22_ = _dafny.Map({default__.fm1(globalState): (d_2577_v1_) + (default__.fm23(p1, globalState))})
            d_2611_v23_: _dafny.Seq
            d_2611_v23_ = _dafny.SeqWithoutIsStrInference([len(_dafny.Set({p1})), 744])
            d_2612_v24_: _dafny.Map
            d_2612_v24_ = _dafny.Map({(self).f8: d_2611_v23_})
            d_2613_v25_: _dafny.Map
            d_2613_v25_ = _dafny.Map({False: len(d_2612_v24_)})
            d_2610_v22_ = (d_2610_v22_).set((p0) + (((d_2613_v25_)[False] if (False) in (d_2613_v25_) else d_2609_cf48_)), d_2577_v1_)
            r1 = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "matdqkva"))) + (d_2577_v1_)
            d_2614_v26_: bool
            d_2614_v26_ = False
            d_2615_v27_: _dafny.Array
            nw454_ = _dafny.Array(False, 29)
            d_2615_v27_ = nw454_
            d_2616_v28_: _dafny.Seq
            d_2616_v28_ = _dafny.SeqWithoutIsStrInference([p1])
            d_2617_v29_: _dafny.Set
            d_2617_v29_ = _dafny.Set({d_2614_v26_, p1})
            rhs397_ = not (p1) or (d_2614_v26_)
            rhs398_ = d_2615_v27_
            rhs399_ = ((d_2616_v28_) + (d_2616_v28_)) != (d_2616_v28_)
            rhs400_ = not(((len((D14_DC42(d_2609_cf48_, True, d_2617_v29_, d_2608_cf49_)).cf74)) != (d_2608_cf49_) if (self).f8 else d_2614_v26_))
            d_2614_v26_ = rhs397_
            d_2615_v27_ = rhs398_
            d_2614_v26_ = rhs399_
            d_2614_v26_ = rhs400_
            d_2618_v31_: _dafny.MultiSet
            d_2618_v31_ = _dafny.MultiSet([(self).f8, d_2614_v26_])
            def iife251_():
                coll97_ = _dafny.Map()
                compr_97_: int
                for compr_97_ in _dafny.IntegerRange(878, 441):
                    d_2619_v30_: int = compr_97_
                    if ((878) <= (d_2619_v30_)) and ((d_2619_v30_) < (441)):
                        coll97_[default__.safeDivisionInt(d_2619_v30_, d_2609_cf48_)] = False
                return _dafny.Map(coll97_)
            d_2614_v26_ = not((self.f4) < ((len(iife251_()
            )) * ((d_2618_v31_).cardinality)))
        elif source28_.is_DC27:
            d_2620___mcc_h5_ = source28_.cf50
            d_2621___mcc_h6_ = source28_.cf51
            d_2622___mcc_h7_ = source28_.cf52
            d_2623_cf52_ = d_2622___mcc_h7_
            d_2624_cf51_ = d_2621___mcc_h6_
            d_2625_cf50_ = d_2620___mcc_h5_
            d_2626_v32_: C0
            nw455_ = C0()
            nw455_.ctor__(56)
            d_2626_v32_ = nw455_
            d_2627_v33_: D3
            d_2627_v33_ = D3_DC9(p1, 959)
            d_2628_v34_: _dafny.Map
            d_2628_v34_ = _dafny.Map({d_2625_cf50_: d_2625_cf50_})
            d_2627_v33_ = (D3_DC9(False, d_2626_v32_.f18) if (d_2626_v32_.f18) not in (d_2628_v34_) else d_2627_v33_)
            d_2629_v35_: _dafny.Array
            nw456_ = _dafny.Array(_dafny.Array(None, 0), 20)
            d_2629_v35_ = nw456_
            d_2630_v36_: _dafny.Map
            d_2630_v36_ = _dafny.Map({d_2623_cf52_: d_2629_v35_})
            d_2629_v35_ = ((d_2630_v36_)[self.f4] if (self.f4) in (d_2630_v36_) else d_2629_v35_)
            d_2631_v37_: str
            d_2631_v37_ = _dafny.CodePoint('k')
            d_2632_v38_: _dafny.Seq
            d_2632_v38_ = _dafny.SeqWithoutIsStrInference([len(d_2577_v1_), -739])
            d_2633_v39_: _dafny.Map
            d_2633_v39_ = _dafny.Map({d_2631_v37_: (d_2632_v38_)[default__.safeIndex(d_2623_cf52_, len(d_2632_v38_))]})
            d_2634_v40_: _dafny.Set
            d_2634_v40_ = _dafny.Set({len(d_2633_v39_)})
            d_2635_v41_: _dafny.Set
            d_2635_v41_ = _dafny.Set({p1})
            d_2636_v42_: _dafny.Set
            d_2636_v42_ = _dafny.Set({d_2628_v34_, d_2628_v34_})
            d_2637_v43_: _dafny.Array
            def lambda130_(d_2638_i1_):
                return (d_2638_i1_) + (self.f4)

            init75_ = lambda130_
            nw457_ = _dafny.Array(None, 2)
            for i0_75_ in range(nw457_.length(0)):
                nw457_[i0_75_] = init75_(i0_75_)
            d_2637_v43_ = nw457_
            d_2639_v44_: _dafny.MultiSet
            d_2639_v44_ = _dafny.MultiSet([d_2637_v43_])
            d_2640_v45_: _dafny.Map
            d_2640_v45_ = _dafny.Map({(d_2639_v44_).cardinality: False})
            d_2641_v46_: _dafny.Array
            nw458_ = _dafny.Array(None, 24)
            nw458_[int(0)] = not(p1)
            nw458_[int(1)] = (self).f8
            nw458_[int(2)] = (d_2634_v40_) == (_dafny.Set({d_2625_cf50_, d_2625_cf50_}))
            nw458_[int(3)] = p1
            nw458_[int(4)] = (self).f8
            nw458_[int(5)] = p1
            nw458_[int(6)] = p1
            nw458_[int(7)] = p1
            nw458_[int(8)] = not(p1)
            nw458_[int(9)] = (d_2635_v41_).ispropersubset(d_2635_v41_)
            nw458_[int(10)] = p1
            nw458_[int(11)] = not(False)
            nw458_[int(12)] = (d_2636_v42_).ispropersubset(d_2636_v42_)
            nw458_[int(13)] = (self).f8
            nw458_[int(14)] = p1
            nw458_[int(15)] = (self).f8
            nw458_[int(16)] = p1
            nw458_[int(17)] = not((self).fm5(d_2577_v1_, default__.fm0(672, d_2623_cf52_, globalState), globalState))
            nw458_[int(18)] = (self).f8
            nw458_[int(19)] = p1
            nw458_[int(20)] = (self.f4) not in (d_2640_v45_)
            nw458_[int(21)] = p1
            nw458_[int(22)] = True
            nw458_[int(23)] = p1
            d_2641_v46_ = nw458_
            index418_ = default__.safeIndex(744, (d_2641_v46_).length(0))
            (d_2641_v46_)[index418_] = p1
            index419_ = default__.safeIndex(744, (d_2641_v46_).length(0))
            (d_2641_v46_)[index419_] = (self).f8
        elif True:
            d_2642___mcc_h8_ = source28_.cf44
            d_2643_cf44_ = d_2642___mcc_h8_
            d_2644_v47_: _dafny.Array
            nw459_ = _dafny.Array(_dafny.Seq({}), 14)
            d_2644_v47_ = nw459_
            d_2645_v48_: _dafny.Map
            d_2645_v48_ = _dafny.Map({(self).fm5(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lv")), True, globalState): d_2644_v47_})
            d_2646_v49_: _dafny.Array
            nw460_ = _dafny.Array(None, 29)
            nw460_[int(0)] = d_2644_v47_
            nw460_[int(1)] = d_2644_v47_
            nw460_[int(2)] = d_2644_v47_
            nw460_[int(3)] = d_2644_v47_
            nw460_[int(4)] = d_2644_v47_
            nw460_[int(5)] = d_2644_v47_
            nw460_[int(6)] = d_2644_v47_
            nw460_[int(7)] = d_2644_v47_
            nw460_[int(8)] = ((d_2645_v48_)[False] if (False) in (d_2645_v48_) else d_2644_v47_)
            nw460_[int(9)] = d_2644_v47_
            nw460_[int(10)] = d_2644_v47_
            nw460_[int(11)] = d_2644_v47_
            nw460_[int(12)] = d_2644_v47_
            nw460_[int(13)] = d_2644_v47_
            nw460_[int(14)] = (d_2644_v47_ if p1 else d_2644_v47_)
            nw460_[int(15)] = d_2644_v47_
            nw460_[int(16)] = d_2644_v47_
            nw460_[int(17)] = d_2644_v47_
            nw460_[int(18)] = d_2644_v47_
            nw460_[int(19)] = d_2644_v47_
            nw460_[int(20)] = d_2644_v47_
            nw460_[int(21)] = d_2644_v47_
            nw460_[int(22)] = d_2644_v47_
            nw460_[int(23)] = d_2644_v47_
            nw460_[int(24)] = d_2644_v47_
            nw460_[int(25)] = d_2644_v47_
            nw460_[int(26)] = ((d_2645_v48_)[p1] if (p1) in (d_2645_v48_) else d_2644_v47_)
            nw460_[int(27)] = d_2644_v47_
            nw460_[int(28)] = d_2644_v47_
            d_2646_v49_ = nw460_
            index420_ = default__.safeIndex(395, (d_2646_v49_).length(0))
            (d_2646_v49_)[index420_] = d_2644_v47_
            d_2647_v50_: _dafny.Array
            nw461_ = _dafny.Array(int(0), 14)
            d_2647_v50_ = nw461_
            d_2648_v51_: D18
            d_2648_v51_ = D18_DC47(d_2644_v47_)
            index421_ = default__.safeIndex(395, (d_2646_v49_).length(0))
            rhs401_ = d_2647_v50_
            rhs402_ = ((d_2577_v1_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_2649_i2_ in range(default__.abs(597))]))) + ((d_2577_v1_) + (d_2577_v1_))
            rhs403_ = (d_2648_v51_).cf83
            lhs244_ = d_2646_v49_
            lhs245_ = default__.safeIndex(395, (d_2646_v49_).length(0))
            r2 = rhs401_
            r1 = rhs402_
            lhs244_[lhs245_] = rhs403_
            d_2650_v52_: bool
            d_2650_v52_ = False
            d_2651_v53_: _dafny.Seq
            d_2651_v53_ = _dafny.SeqWithoutIsStrInference([d_2650_v52_, (self).f8, p1, d_2650_v52_])
            d_2652_v54_: _dafny.Seq
            d_2652_v54_ = _dafny.SeqWithoutIsStrInference([d_2651_v53_, d_2651_v53_])
            d_2653_v56_: _dafny.Set
            d_2653_v56_ = _dafny.Set({d_2651_v53_, d_2651_v53_, (d_2651_v53_).set(default__.safeIndex(self.f4, len(d_2651_v53_)), (self).f8)})
            d_2654_v57_: _dafny.Map
            d_2654_v57_ = _dafny.Map({p1: d_2653_v56_})
            def iife252_():
                coll98_ = _dafny.Set()
                compr_98_: _dafny.Seq
                for compr_98_ in (d_2652_v54_).Elements:
                    d_2655_v55_: _dafny.Seq = compr_98_
                    if (d_2655_v55_) in (d_2652_v54_):
                        coll98_ = coll98_.union(_dafny.Set([d_2655_v55_]))
                return _dafny.Set(coll98_)
            d_2650_v52_ = (((d_2654_v57_)[(self).f8] if ((self).f8) in (d_2654_v57_) else d_2653_v56_)).issubset((iife252_()
            ) - (d_2653_v56_))
            d_2656_v58_: _dafny.Array
            nw462_ = _dafny.Array(False, 7)
            d_2656_v58_ = nw462_
            index422_ = default__.safeIndex(841, (d_2656_v58_).length(0))
            (d_2656_v58_)[index422_] = p1
            index423_ = default__.safeIndex(841, (d_2656_v58_).length(0))
            (d_2656_v58_)[index423_] = (d_2651_v53_) < (d_2651_v53_)
            d_2657_v59_: D2
            d_2657_v59_ = D2_DC6(self.f4, d_2650_v52_)
            d_2650_v52_ = (d_2657_v59_).cf11
        (self).f4 = default__.fm1(globalState)
        hi15_ = p0
        for d_2658_i3_ in range((len(_dafny.SeqWithoutIsStrInference([p1])) if True else (0) - ((0) - (p0))), hi15_):
            (self).f4 = len(d_2577_v1_)
            d_2659_v60_: _dafny.Array
            nw463_ = _dafny.Array(False, 13)
            d_2659_v60_ = nw463_
            d_2660_v61_: _dafny.Seq
            d_2660_v61_ = _dafny.SeqWithoutIsStrInference([p1, False])
            index424_ = default__.safeIndex(195, (d_2659_v60_).length(0))
            (d_2659_v60_)[index424_] = (d_2660_v61_)[default__.safeIndex(d_2658_i3_, len(d_2660_v61_))]
            index425_ = default__.safeIndex(195, (d_2659_v60_).length(0))
            (d_2659_v60_)[index425_] = (self).f8
            d_2661_v62_: _dafny.Seq
            d_2661_v62_ = _dafny.SeqWithoutIsStrInference([self.f4])
            d_2662_v63_: str
            d_2662_v63_ = _dafny.CodePoint('v')
            d_2663_v64_: C3
            nw464_ = C3()
            nw464_.ctor__((len(d_2660_v61_)) + (len(d_2661_v62_)), _dafny.SeqWithoutIsStrInference([d_2662_v63_]), p0)
            d_2663_v64_ = nw464_
            d_2664_v65_: C3
            nw465_ = C3()
            nw465_.ctor__(((0) - ((D11_DC35((d_2663_v64_).f11, (d_2659_v60_)[default__.safeIndex(195, (d_2659_v60_).length(0))], (self).fm5(d_2577_v1_, True, globalState))).cf61)) * (875), (d_2663_v64_).f12, (d_2663_v64_).f11)
            d_2664_v65_ = nw465_
        d_2665_v66_: _dafny.Seq
        d_2665_v66_ = _dafny.SeqWithoutIsStrInference([p0])
        d_2666_v67_: _dafny.Map
        d_2666_v67_ = _dafny.Map({d_2665_v66_: d_2665_v66_})
        d_2666_v67_ = ((d_2666_v67_) | (d_2666_v67_)) | (d_2666_v67_)
        d_2667_v68_: _dafny.Map
        d_2667_v68_ = _dafny.Map({self.f4: p1})
        d_2668_v69_: _dafny.Map
        d_2668_v69_ = _dafny.Map({p0: len(d_2665_v66_)})
        d_2669_v70_: _dafny.Map
        d_2669_v70_ = _dafny.Map({_dafny.Set({len(d_2667_v68_)}): (d_2668_v69_) | (d_2668_v69_)})
        d_2670_v71_: _dafny.Seq
        d_2670_v71_ = _dafny.SeqWithoutIsStrInference([(self).f8, False, (self).f8, (self).f8, (self).f8])
        d_2671_v72_: _dafny.Set
        d_2671_v72_ = _dafny.Set({len(d_2670_v71_), self.f4})
        d_2672_v73_: _dafny.Map
        d_2672_v73_ = _dafny.Map({p1: d_2668_v69_})
        d_2669_v70_ = (d_2669_v70_).set((d_2671_v72_) | (d_2671_v72_), ((d_2672_v73_)[(self).f8] if ((self).f8) in (d_2672_v73_) else d_2668_v69_))
        (self).f4 = self.f4
        d_2673_v74_: D0
        d_2673_v74_ = D0_DC0(d_2665_v66_)
        d_2674_v75_: _dafny.Seq
        def iife253_(_pat_let77_0):
            def iife254_(d_2675_dt__update__tmp_h0_):
                def iife255_(_pat_let78_0):
                    def iife256_(d_2676_dt__update_hcf0_h0_):
                        return D0_DC0(d_2676_dt__update_hcf0_h0_)
                    return iife256_(_pat_let78_0)
                return iife255_(_dafny.SeqWithoutIsStrInference([self.f4]))
            return iife254_(_pat_let77_0)
        d_2674_v75_ = _dafny.SeqWithoutIsStrInference([iife253_(D0_DC0(d_2665_v66_)), d_2673_v74_])
        r0 = d_2674_v75_
        r1 = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q"))
        d_2677_v76_: _dafny.MultiSet
        d_2677_v76_ = _dafny.MultiSet([self.f4])
        d_2678_v77_: _dafny.Set
        d_2678_v77_ = _dafny.Set({(self).f8})
        d_2679_v78_: _dafny.Array
        nw466_ = _dafny.Array(None, 8)
        nw466_[int(0)] = len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_2680_i4_ in range(default__.abs(946))])) + (d_2577_v1_))
        nw466_[int(1)] = ((d_2677_v76_).cardinality) - (self.f4)
        nw466_[int(2)] = (len(d_2678_v77_)) + (self.f4)
        nw466_[int(3)] = p0
        nw466_[int(4)] = p0
        nw466_[int(5)] = (0) - (self.f4)
        nw466_[int(6)] = len(d_2577_v1_)
        nw466_[int(7)] = 390
        d_2679_v78_ = nw466_
        r2 = d_2679_v78_
        return r0, r1, r2

    def m6(self, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: int = int(0)
        d_2681_v0_: D8
        d_2681_v0_ = D8_DC29()
        source29_ = d_2681_v0_
        if source29_.is_DC29:
            d_2682_v1_: bool
            d_2682_v1_ = True
            d_2683_v2_: _dafny.Array
            nw467_ = _dafny.Array(None, 1)
            nw467_[int(0)] = -556
            d_2683_v2_ = nw467_
            d_2684_v3_: _dafny.Seq
            d_2684_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j"))
            d_2685_v4_: _dafny.Map
            d_2685_v4_ = _dafny.Map({default__.fm23((self).f8, globalState): len(d_2684_v3_)})
            d_2686_v5_: _dafny.MultiSet
            d_2686_v5_ = _dafny.MultiSet([len(d_2685_v4_)])
            d_2687_v6_: _dafny.Map
            d_2687_v6_ = _dafny.Map({self.f4: (self).f8})
            index426_ = default__.safeIndex(7, (d_2683_v2_).length(0))
            (d_2683_v2_)[index426_] = ((d_2686_v5_)[len(d_2687_v6_)] if (len(d_2687_v6_)) in (d_2686_v5_) else self.f4)
            index427_ = default__.safeIndex(7, (d_2683_v2_).length(0))
            rhs404_ = (self).f8
            rhs405_ = (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_2688_i0_ in range(default__.abs(323))]))) * (288)
            lhs246_ = d_2683_v2_
            lhs247_ = default__.safeIndex(7, (d_2683_v2_).length(0))
            d_2682_v1_ = rhs404_
            lhs246_[lhs247_] = rhs405_
            index428_ = default__.safeIndex(7, (d_2683_v2_).length(0))
            (d_2683_v2_)[index428_] = (d_2683_v2_)[default__.safeIndex(7, (d_2683_v2_).length(0))]
            r1 = default__.safeDivisionInt((self.f4) - (self.f4), (d_2683_v2_)[default__.safeIndex(7, (d_2683_v2_).length(0))])
            (self).f4 = -882
        elif source29_.is_DC30:
            d_2689___mcc_h0_ = source29_.cf54
            d_2690___mcc_h1_ = source29_.cf55
            d_2691___mcc_h2_ = source29_.cf56
            d_2692_cf56_ = d_2691___mcc_h2_
            d_2693_cf55_ = d_2690___mcc_h1_
            d_2694_cf54_ = d_2689___mcc_h0_
            d_2692_cf56_ = (self).f8
            d_2695_v7_: _dafny.Seq
            d_2695_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qlevhtgt"))
            d_2692_cf56_ = (self).fm5((d_2695_v7_) + (d_2695_v7_), d_2692_cf56_, globalState)
            d_2696_v8_: _dafny.Array
            def lambda131_(d_2697_cf54_):
                def lambda132_(d_2698_i1_):
                    return default__.safeModuloInt(d_2698_i1_, d_2697_cf54_)

                return lambda132_

            init76_ = lambda131_(d_2694_cf54_)
            nw468_ = _dafny.Array(None, 4)
            for i0_76_ in range(nw468_.length(0)):
                nw468_[i0_76_] = init76_(i0_76_)
            d_2696_v8_ = nw468_
            d_2699_v9_: _dafny.Seq
            d_2699_v9_ = _dafny.SeqWithoutIsStrInference([self.f4, len(d_2695_v7_)])
            d_2700_v10_: _dafny.Map
            d_2700_v10_ = _dafny.Map({(d_2699_v9_)[default__.safeIndex(d_2694_cf54_, len(d_2699_v9_))]: (self).f8})
            index429_ = default__.safeIndex(693, (d_2696_v8_).length(0))
            (d_2696_v8_)[index429_] = default__.safeModuloInt(default__.fm1(globalState), len(d_2700_v10_))
            d_2701_v11_: _dafny.MultiSet
            d_2701_v11_ = _dafny.MultiSet([d_2694_cf54_, self.f4, d_2694_cf54_])
            d_2702_v12_: _dafny.Set
            d_2702_v12_ = _dafny.Set({d_2692_cf56_, (self).f8, d_2692_cf56_, (self).f8, False})
            index430_ = default__.safeIndex(693, (d_2696_v8_).length(0))
            (d_2696_v8_)[index430_] = ((d_2701_v11_).intersection((_dafny.MultiSet([len(d_2695_v7_), len(d_2702_v12_)])).intersection(_dafny.MultiSet(d_2699_v9_)))).cardinality
            d_2693_cf55_ = d_2693_cf55_
        elif source29_.is_DC28:
            d_2703___mcc_h3_ = source29_.cf53
            d_2704_cf53_ = d_2703___mcc_h3_
            d_2705_v13_: bool
            d_2705_v13_ = True
            d_2706_v14_: _dafny.Map
            d_2706_v14_ = _dafny.Map({d_2705_v13_: d_2705_v13_})
            d_2707_v15_: _dafny.Map
            d_2707_v15_ = _dafny.Map({default__.fm1(globalState): d_2706_v14_})
            d_2708_v16_: _dafny.Map
            d_2708_v16_ = _dafny.Map({d_2707_v15_: d_2705_v13_})
            d_2705_v13_ = ((d_2708_v16_)[(d_2707_v15_) | (d_2707_v15_)] if ((d_2707_v15_) | (d_2707_v15_)) in (d_2708_v16_) else (default__.fm1(globalState)) <= (self.f4))
            d_2709_v17_: _dafny.Seq
            d_2709_v17_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cfx"))
            d_2705_v13_ = ((d_2709_v17_) == (d_2709_v17_)) or ((self).f8)
            d_2705_v13_ = ((self.f4) >= (self.f4)) and (d_2705_v13_)
            d_2710_v18_: _dafny.MultiSet
            d_2710_v18_ = _dafny.MultiSet([self.f4, self.f4])
            (self).f4 = ((self.f4) + (self.f4)) * ((d_2710_v18_).cardinality)
        elif True:
            d_2711___mcc_h4_ = source29_.cf57
            d_2712_cf57_ = d_2711___mcc_h4_
            d_2713_v19_: bool
            d_2713_v19_ = True
            d_2713_v19_ = (self.f4) == (self.f4)
            d_2714_v20_: _dafny.Map
            d_2714_v20_ = _dafny.Map({True: d_2713_v19_})
            if (default__.safeModuloInt(len(d_2714_v20_), self.f4)) >= (self.f4):
                d_2715_v21_: _dafny.Array
                nw469_ = _dafny.Array(_dafny.Map({}), 8)
                d_2715_v21_ = nw469_
                index431_ = default__.safeIndex(904, (d_2715_v21_).length(0))
                (d_2715_v21_)[index431_] = _dafny.Map({self.f4: (self).f8})
                d_2716_v22_: _dafny.Map
                d_2716_v22_ = _dafny.Map({self.f4: d_2713_v19_})
                d_2717_v23_: _dafny.MultiSet
                d_2717_v23_ = _dafny.MultiSet([d_2713_v19_])
                d_2718_v24_: D5
                d_2718_v24_ = D5_DC16(d_2713_v19_, d_2713_v19_, (self).f8, self.f4, ((d_2716_v22_)[(d_2717_v23_).cardinality] if ((d_2717_v23_).cardinality) in (d_2716_v22_) else (self).f8))
                index432_ = default__.safeIndex(904, (d_2715_v21_).length(0))
                (d_2715_v21_)[index432_] = ((d_2716_v22_) | (d_2716_v22_) if (124) not in (_dafny.Map({self.f4: len(_dafny.SeqWithoutIsStrInference([self.f4, (d_2718_v24_).cf32]))})) else _dafny.Map({self.f4: d_2713_v19_}))
                d_2713_v19_ = (self).f8
                r1 = ((self.f4) * (self.f4)) + (default__.safeModuloInt(self.f4, self.f4))
                d_2719_v25_: C4
                nw470_ = C4()
                nw470_.ctor__(self.f4)
                d_2719_v25_ = nw470_
                d_2720_v26_: _dafny.Seq
                d_2720_v26_ = _dafny.SeqWithoutIsStrInference([True, not(d_2713_v19_), (self).f8])
                d_2720_v26_ = (d_2720_v26_) + (d_2720_v26_)
            elif True:
                d_2721_v27_: _dafny.Array
                nw471_ = _dafny.Array(False, 15)
                d_2721_v27_ = nw471_
                d_2722_v28_: _dafny.Seq
                d_2722_v28_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hthe"))
                d_2723_v29_: _dafny.Map
                d_2723_v29_ = _dafny.Map({len(d_2722_v28_): self.f4})
                index433_ = default__.safeIndex(159, (d_2721_v27_).length(0))
                (d_2721_v27_)[index433_] = default__.fm0(self.f4, len(d_2723_v29_), globalState)
                d_2724_v30_: C2
                nw472_ = C2()
                nw472_.ctor__((self).f8, self.f4)
                d_2724_v30_ = nw472_
                d_2725_v31_: _dafny.Map
                d_2725_v31_ = _dafny.Map({d_2724_v30_: _dafny.Map({428: (self).f8})})
                d_2726_v32_: _dafny.Map
                d_2726_v32_ = _dafny.Map({self.f4: d_2724_v30_.f13})
                index434_ = default__.safeIndex(159, (d_2721_v27_).length(0))
                (d_2721_v27_)[index434_] = (d_2724_v30_) in ((d_2725_v31_).set(d_2724_v30_, d_2726_v32_))
                d_2727_v33_: str
                d_2727_v33_ = _dafny.CodePoint('p')
                d_2728_v34_: _dafny.Map
                d_2728_v34_ = _dafny.Map({self.f4: d_2727_v33_})
                d_2729_v35_: _dafny.Map
                d_2729_v35_ = _dafny.Map({d_2728_v34_: (d_2724_v30_).f14})
                d_2730_v36_: _dafny.Map
                d_2730_v36_ = _dafny.Map({(d_2724_v30_).f14: d_2727_v33_})
                d_2729_v35_ = (d_2729_v35_).set(d_2730_v36_, self.f4)
                d_2731_v37_: C4
                nw473_ = C4()
                nw473_.ctor__((d_2724_v30_).f14)
                d_2731_v37_ = nw473_
                d_2732_v38_: D25
                d_2732_v38_ = D25_DC65(d_2731_v37_)
                pat_let_tv64_ = d_2731_v37_
                def iife257_(_pat_let79_0):
                    def iife258_(d_2733_dt__update__tmp_h0_):
                        def iife259_(_pat_let80_0):
                            def iife260_(d_2734_dt__update_hcf112_h0_):
                                return D25_DC65(d_2734_dt__update_hcf112_h0_)
                            return iife260_(_pat_let80_0)
                        return iife259_(pat_let_tv64_)
                    return iife258_(_pat_let79_0)
                rhs406_ = -519
                rhs407_ = iife257_(d_2732_v38_)
                lhs248_ = self
                lhs248_.f4 = rhs406_
                d_2732_v38_ = rhs407_
                d_2714_v20_ = (d_2714_v20_).set((d_2721_v27_)[default__.safeIndex(159, (d_2721_v27_).length(0))], d_2713_v19_)
                d_2735_v40_: _dafny.Array
                def lambda133_(d_2736_v32_, d_2737_v30_, d_2738_v19_, d_2739_v27_):
                    def lambda134_(d_2740_i2_):
                        def iife261_():
                            coll99_ = _dafny.Map()
                            compr_99_: D3
                            for compr_99_ in (_dafny.Set({D3_DC10(len(d_2736_v32_), len(_dafny.SeqWithoutIsStrInference([d_2737_v30_.f13])), d_2738_v19_, (d_2737_v30_).f14, (d_2739_v27_)[default__.safeIndex(159, (d_2739_v27_).length(0))])})).Elements:
                                d_2741_v39_: D3 = compr_99_
                                if (d_2741_v39_) in (_dafny.Set({D3_DC10(len(d_2736_v32_), len(_dafny.SeqWithoutIsStrInference([d_2737_v30_.f13])), d_2738_v19_, (d_2737_v30_).f14, (d_2739_v27_)[default__.safeIndex(159, (d_2739_v27_).length(0))])})):
                                    coll99_[d_2741_v39_] = d_2737_v30_.f13
                            return _dafny.Map(coll99_)
                        return (iife261_()
                        ) | (_dafny.Map({D3_DC10((d_2737_v30_).f14, len(_dafny.Set({self.f4})), d_2737_v30_.f13, self.f4, (d_2739_v27_)[default__.safeIndex(159, (d_2739_v27_).length(0))]): (d_2739_v27_)[default__.safeIndex(159, (d_2739_v27_).length(0))]}))

                    return lambda134_

                init77_ = lambda133_(d_2726_v32_, d_2724_v30_, d_2713_v19_, d_2721_v27_)
                nw474_ = _dafny.Array(None, 1)
                for i0_77_ in range(nw474_.length(0)):
                    nw474_[i0_77_] = init77_(i0_77_)
                d_2735_v40_ = nw474_
                d_2742_v41_: _dafny.Map
                d_2742_v41_ = _dafny.Map({(d_2721_v27_)[default__.safeIndex(159, (d_2721_v27_).length(0))]: self.f4})
                d_2743_v42_: D3
                d_2743_v42_ = D3_DC10(len(d_2742_v41_), self.f4, (d_2721_v27_)[default__.safeIndex(159, (d_2721_v27_).length(0))], self.f4, (d_2721_v27_)[default__.safeIndex(159, (d_2721_v27_).length(0))])
                d_2744_v43_: _dafny.Map
                d_2744_v43_ = _dafny.Map({d_2743_v42_: d_2713_v19_})
                index435_ = default__.safeIndex(46, (d_2735_v40_).length(0))
                (d_2735_v40_)[index435_] = d_2744_v43_
                index436_ = default__.safeIndex(46, (d_2735_v40_).length(0))
                (d_2735_v40_)[index436_] = default__.fm46(self.f4, d_2724_v30_.f13, self.f4, globalState)
            d_2713_v19_ = (self).f8
            d_2745_v44_: str
            d_2745_v44_ = _dafny.CodePoint('w')
            d_2746_v45_: _dafny.Seq
            d_2746_v45_ = _dafny.SeqWithoutIsStrInference([d_2745_v44_, d_2745_v44_])
            d_2747_v46_: _dafny.Array
            nw475_ = _dafny.Array(None, 3)
            nw475_[int(0)] = d_2745_v44_
            nw475_[int(1)] = d_2745_v44_
            nw475_[int(2)] = (d_2746_v45_)[default__.safeIndex(self.f4, len(d_2746_v45_))]
            d_2747_v46_ = nw475_
            d_2747_v46_ = d_2747_v46_
        d_2748_v47_: D19
        d_2748_v47_ = D19_DC52(self.f4, self.f4, self.f4, (self).f8)
        source30_ = d_2748_v47_
        if source30_.is_DC51:
            d_2749___mcc_h5_ = source30_.cf90
            d_2750___mcc_h6_ = source30_.cf91
            d_2751_cf91_ = d_2750___mcc_h6_
            d_2752_cf90_ = d_2749___mcc_h5_
            d_2753_v48_: _dafny.Seq
            d_2753_v48_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yvngvydnp"))
            d_2754_v49_: str
            d_2754_v49_ = _dafny.CodePoint('k')
            d_2755_v50_: _dafny.Map
            d_2755_v50_ = _dafny.Map({d_2752_cf90_: d_2754_v49_})
            d_2756_v51_: _dafny.Seq
            d_2756_v51_ = _dafny.SeqWithoutIsStrInference([len(d_2755_v50_)])
            d_2757_v52_: _dafny.MultiSet
            d_2757_v52_ = _dafny.MultiSet([d_2751_cf91_, d_2751_cf91_])
            d_2758_v53_: _dafny.Map
            d_2758_v53_ = _dafny.Map({(self).f8: (self).f8})
            d_2759_v54_: _dafny.Map
            d_2759_v54_ = _dafny.Map({d_2752_cf90_: (self).f8})
            d_2760_v55_: _dafny.Array
            nw476_ = _dafny.Array(None, 17)
            nw476_[int(0)] = default__.safeDivisionInt((_dafny.MultiSet([self.f4])).cardinality, d_2751_cf91_)
            nw476_[int(1)] = len((d_2753_v48_) + (d_2753_v48_))
            nw476_[int(2)] = (D17_DC46((self).f8, self.f4, (0) - (d_2752_cf90_), (self).f8)).cf81
            nw476_[int(3)] = (d_2756_v51_)[default__.safeIndex(d_2752_cf90_, len(d_2756_v51_))]
            nw476_[int(4)] = d_2752_cf90_
            nw476_[int(5)] = len((d_2753_v48_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aupng"))))
            nw476_[int(6)] = d_2751_cf91_
            nw476_[int(7)] = default__.safeDivisionInt(-765, self.f4)
            nw476_[int(8)] = default__.safeModuloInt(d_2751_cf91_, d_2751_cf91_)
            nw476_[int(9)] = (d_2756_v51_)[default__.safeIndex(((d_2757_v52_)[d_2752_cf90_] if (d_2752_cf90_) in (d_2757_v52_) else 958), len(d_2756_v51_))]
            nw476_[int(10)] = self.f4
            nw476_[int(11)] = (d_2751_cf91_) * (235)
            nw476_[int(12)] = self.f4
            nw476_[int(13)] = d_2751_cf91_
            nw476_[int(14)] = default__.safeDivisionInt(len(d_2758_v53_), (0) - (len(d_2753_v48_)))
            nw476_[int(15)] = (0) - (len(d_2759_v54_))
            nw476_[int(16)] = d_2751_cf91_
            d_2760_v55_ = nw476_
            index437_ = default__.safeIndex(633, (d_2760_v55_).length(0))
            (d_2760_v55_)[index437_] = default__.safeModuloInt((0) - (self.f4), self.f4)
            d_2761_v56_: _dafny.MultiSet
            d_2761_v56_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_2762_i3_ in range(default__.abs(180))])])
            index438_ = default__.safeIndex(633, (d_2760_v55_).length(0))
            (d_2760_v55_)[index438_] = (d_2761_v56_).cardinality
            d_2763_v57_: bool
            d_2763_v57_ = True
            d_2763_v57_ = (self).f8
            d_2752_cf90_ = d_2751_cf91_
            d_2764_v58_: _dafny.Seq
            d_2764_v58_ = _dafny.SeqWithoutIsStrInference([((d_2759_v54_)[d_2752_cf90_] if (d_2752_cf90_) in (d_2759_v54_) else d_2763_v57_), d_2763_v57_])
            d_2765_v59_: _dafny.Array
            nw477_ = _dafny.Array(None, 12)
            nw477_[int(0)] = d_2764_v58_
            nw477_[int(1)] = d_2764_v58_
            nw477_[int(2)] = d_2764_v58_
            nw477_[int(3)] = d_2764_v58_
            nw477_[int(4)] = d_2764_v58_
            nw477_[int(5)] = d_2764_v58_
            nw477_[int(6)] = d_2764_v58_
            nw477_[int(7)] = d_2764_v58_
            nw477_[int(8)] = d_2764_v58_
            nw477_[int(9)] = d_2764_v58_
            nw477_[int(10)] = d_2764_v58_
            nw477_[int(11)] = d_2764_v58_
            d_2765_v59_ = nw477_
            d_2766_v60_: D18
            d_2766_v60_ = D18_DC47(d_2765_v59_)
            d_2767_v61_: _dafny.Map
            d_2767_v61_ = _dafny.Map({(d_2760_v55_)[default__.safeIndex(633, (d_2760_v55_).length(0))]: d_2765_v59_})
            d_2768_v62_: _dafny.Array
            nw478_ = _dafny.Array(None, 26)
            nw478_[int(0)] = d_2765_v59_
            nw478_[int(1)] = d_2765_v59_
            nw478_[int(2)] = d_2765_v59_
            nw478_[int(3)] = (d_2766_v60_).cf83
            nw478_[int(4)] = d_2765_v59_
            nw478_[int(5)] = d_2765_v59_
            nw478_[int(6)] = ((d_2767_v61_)[d_2752_cf90_] if (d_2752_cf90_) in (d_2767_v61_) else d_2765_v59_)
            nw478_[int(7)] = d_2765_v59_
            nw478_[int(8)] = d_2765_v59_
            nw478_[int(9)] = d_2765_v59_
            nw478_[int(10)] = d_2765_v59_
            nw478_[int(11)] = d_2765_v59_
            nw478_[int(12)] = d_2765_v59_
            nw478_[int(13)] = (d_2765_v59_ if (self).f8 else d_2765_v59_)
            nw478_[int(14)] = d_2765_v59_
            nw478_[int(15)] = d_2765_v59_
            nw478_[int(16)] = d_2765_v59_
            nw478_[int(17)] = d_2765_v59_
            nw478_[int(18)] = d_2765_v59_
            nw478_[int(19)] = d_2765_v59_
            nw478_[int(20)] = d_2765_v59_
            nw478_[int(21)] = d_2765_v59_
            nw478_[int(22)] = d_2765_v59_
            nw478_[int(23)] = (d_2766_v60_).cf83
            nw478_[int(24)] = d_2765_v59_
            nw478_[int(25)] = d_2765_v59_
            d_2768_v62_ = nw478_
            index439_ = default__.safeIndex(368, (d_2768_v62_).length(0))
            (d_2768_v62_)[index439_] = d_2765_v59_
            d_2769_v63_: _dafny.Map
            d_2769_v63_ = _dafny.Map({(self).f8: d_2765_v59_})
            d_2770_v64_: _dafny.Set
            d_2770_v64_ = _dafny.Set({not(d_2763_v57_), d_2763_v57_})
            d_2771_v65_: _dafny.Seq
            d_2771_v65_ = _dafny.SeqWithoutIsStrInference([d_2770_v64_])
            index440_ = default__.safeIndex(368, (d_2768_v62_).length(0))
            (d_2768_v62_)[index440_] = ((d_2769_v63_)[(_dafny.Set({d_2763_v57_})).ispropersubset((d_2771_v65_)[default__.safeIndex((0) - ((d_2760_v55_)[default__.safeIndex(633, (d_2760_v55_).length(0))]), len(d_2771_v65_))])] if ((_dafny.Set({d_2763_v57_})).ispropersubset((d_2771_v65_)[default__.safeIndex((0) - ((d_2760_v55_)[default__.safeIndex(633, (d_2760_v55_).length(0))]), len(d_2771_v65_))])) in (d_2769_v63_) else d_2765_v59_)
        elif source30_.is_DC52:
            d_2772___mcc_h7_ = source30_.cf92
            d_2773___mcc_h8_ = source30_.cf93
            d_2774___mcc_h9_ = source30_.cf94
            d_2775___mcc_h10_ = source30_.cf95
            d_2776_cf95_ = d_2775___mcc_h10_
            d_2777_cf94_ = d_2774___mcc_h9_
            d_2778_cf93_ = d_2773___mcc_h8_
            d_2779_cf92_ = d_2772___mcc_h7_
            (self).f4 = default__.fm1(globalState)
            d_2780_v66_: _dafny.Seq
            d_2780_v66_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rvvfgyry"))
            d_2779_cf92_ = (0) - (default__.safeModuloInt(len(d_2780_v66_), d_2777_cf94_))
            r1 = self.f4
            if not(d_2776_cf95_):
                d_2781_v67_: _dafny.Seq
                d_2781_v67_ = _dafny.SeqWithoutIsStrInference([d_2776_cf95_, (self).f8, (self).f8, (self).f8, False])
                d_2776_cf95_ = not((d_2781_v67_)[default__.safeIndex(-847, len(d_2781_v67_))])
                d_2776_cf95_ = (self).f8
                d_2782_v68_: _dafny.Map
                d_2782_v68_ = _dafny.Map({_dafny.CodePoint('w'): not((self).f8)})
                d_2783_v69_: T0
                nw479_ = C1()
                nw479_.ctor__((self.f4) * (self.f4), default__.safeModuloInt(len(d_2782_v68_), len(d_2780_v66_)), self.f4, not((self).f8))
                d_2783_v69_ = nw479_
                d_2783_v69_ = (d_2783_v69_ if (self).f8 else d_2783_v69_)
                d_2784_v70_: _dafny.MultiSet
                d_2784_v70_ = _dafny.MultiSet([self.f4, (0) - ((0) - (len(d_2780_v66_)))])
                d_2785_v71_: _dafny.Seq
                d_2785_v71_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_2786_i4_ in range(default__.abs(-642))])), d_2779_cf92_, (d_2784_v70_).cardinality, -339, d_2777_cf94_])
                d_2787_v72_: _dafny.Set
                d_2787_v72_ = _dafny.Set({d_2778_cf93_})
                d_2788_v73_: _dafny.Map
                d_2788_v73_ = _dafny.Map({len(d_2787_v72_): _dafny.CodePoint('r')})
                d_2789_v74_: str
                d_2789_v74_ = _dafny.CodePoint('j')
                d_2790_v75_: _dafny.Map
                d_2790_v75_ = _dafny.Map({(0) - ((0) - (d_2779_cf92_)): d_2776_cf95_})
                d_2791_v76_: _dafny.Map
                d_2791_v76_ = _dafny.Map({(self).f8: not(d_2776_cf95_)})
                d_2776_cf95_ = ((self).f8 if ((d_2780_v66_).set(default__.safeIndex(len(d_2785_v71_), len(d_2780_v66_)), ((d_2788_v73_)[d_2783_v69_.f4] if (d_2783_v69_.f4) in (d_2788_v73_) else d_2789_v74_))) == (d_2780_v66_) else ((d_2790_v75_)[(0) - (len(d_2791_v76_))] if ((0) - (len(d_2791_v76_))) in (d_2790_v75_) else (self).f8))
                d_2792_v77_: C4
                nw480_ = C4()
                nw480_.ctor__(d_2779_cf92_)
                d_2792_v77_ = nw480_
            elif True:
                d_2793_v78_: _dafny.Seq
                d_2793_v78_ = _dafny.SeqWithoutIsStrInference([self.f4])
                d_2794_v79_: _dafny.Array
                def lambda135_(d_2795_i5_):
                    return (d_2795_i5_) * (self.f4)

                init78_ = lambda135_
                nw481_ = _dafny.Array(None, 27)
                for i0_78_ in range(nw481_.length(0)):
                    nw481_[i0_78_] = init78_(i0_78_)
                d_2794_v79_ = nw481_
                d_2796_v80_: _dafny.Map
                d_2796_v80_ = _dafny.Map({824: d_2794_v79_})
                d_2797_v81_: int
                d_2798_v82_: int
                d_2799_v83_: _dafny.Array
                d_2800_v84_: bool
                out59_: int
                out60_: int
                out61_: _dafny.Array
                out62_: bool
                out59_, out60_, out61_, out62_ = default__.m0(_dafny.Map({898: (self).f8}), d_2793_v78_, (d_2796_v80_).set(d_2779_cf92_, d_2794_v79_), globalState)
                d_2797_v81_ = out59_
                d_2798_v82_ = out60_
                d_2799_v83_ = out61_
                d_2800_v84_ = out62_
                (self).f4 = (638 if (self).f8 else default__.safeDivisionInt(self.f4, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_2801_i6_ in range(default__.abs(-544))]))))
                d_2802_v85_: _dafny.Array
                def lambda136_(d_2803_i7_):
                    return D1_DC3(_dafny.Map({True: self.f4}))

                init79_ = lambda136_
                nw482_ = _dafny.Array(None, 22)
                for i0_79_ in range(nw482_.length(0)):
                    nw482_[i0_79_] = init79_(i0_79_)
                d_2802_v85_ = nw482_
                d_2804_v86_: _dafny.Map
                d_2804_v86_ = _dafny.Map({d_2800_v84_: d_2780_v66_})
                d_2805_v87_: _dafny.Map
                d_2805_v87_ = _dafny.Map({(self).fm5(((d_2804_v86_)[d_2800_v84_] if (d_2800_v84_) in (d_2804_v86_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lbqkomvge"))), d_2776_cf95_, globalState): -390})
                d_2806_v88_: D1
                d_2806_v88_ = D1_DC3(d_2805_v87_)
                index441_ = default__.safeIndex(309, (d_2802_v85_).length(0))
                (d_2802_v85_)[index441_] = d_2806_v88_
                index442_ = default__.safeIndex(309, (d_2802_v85_).length(0))
                (d_2802_v85_)[index442_] = d_2806_v88_
                d_2776_cf95_ = (self).f8
                d_2780_v66_ = d_2780_v66_
        elif True:
            d_2807___mcc_h11_ = source30_.cf89
            d_2808_cf89_ = d_2807___mcc_h11_
            d_2809_v89_: str
            d_2809_v89_ = _dafny.CodePoint('u')
            d_2810_v90_: _dafny.Map
            d_2810_v90_ = _dafny.Map({(self.f4) + (self.f4): d_2809_v89_})
            d_2811_v91_: _dafny.Seq
            d_2811_v91_ = _dafny.SeqWithoutIsStrInference([self.f4, self.f4])
            d_2810_v90_ = (d_2810_v90_).set(len((d_2811_v91_ if False else d_2811_v91_)), d_2809_v89_)
            d_2812_v92_: _dafny.Map
            d_2812_v92_ = _dafny.Map({not(((self).f8) == (False)): (self).f8})
            d_2813_v93_: _dafny.Seq
            d_2813_v93_ = _dafny.SeqWithoutIsStrInference([(self).f8])
            d_2814_v94_: D1
            d_2814_v94_ = D1_DC4(self.f4, d_2813_v93_, self.f4)
            d_2815_v95_: _dafny.Seq
            d_2815_v95_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(self).f8])])
            d_2816_v96_: D4
            d_2816_v96_ = D4_DC13(d_2814_v94_, (self).f8, (d_2815_v95_)[default__.safeIndex(self.f4, len(d_2815_v95_))])
            d_2812_v92_ = (d_2812_v92_).set((False) and ((self).f8), (d_2816_v96_).cf25)
            d_2817_v97_: _dafny.Array
            def lambda137_(d_2818_i8_):
                return (self).f8

            init80_ = lambda137_
            nw483_ = _dafny.Array(None, 19)
            for i0_80_ in range(nw483_.length(0)):
                nw483_[i0_80_] = init80_(i0_80_)
            d_2817_v97_ = nw483_
            d_2817_v97_ = d_2817_v97_
            d_2819_v98_: _dafny.Array
            nw484_ = _dafny.Array(_dafny.Array(None, 0), 1)
            d_2819_v98_ = nw484_
            d_2820_v99_: bool
            d_2820_v99_ = False
            d_2821_v100_: D0
            d_2821_v100_ = D0_DC0(_dafny.SeqWithoutIsStrInference([self.f4 for d_2822_i9_ in range(default__.abs(818))]))
            d_2823_v101_: _dafny.Set
            d_2823_v101_ = _dafny.Set({self.f4})
            rhs408_ = ((d_2821_v100_).cf0).set(default__.safeIndex((self.f4) + (self.f4), len((d_2821_v100_).cf0)), self.f4)
            rhs409_ = d_2819_v98_
            rhs410_ = (d_2823_v101_).ispropersubset(d_2823_v101_)
            d_2811_v91_ = rhs408_
            d_2819_v98_ = rhs409_
            d_2820_v99_ = rhs410_
        d_2824_v102_: bool
        d_2824_v102_ = True
        d_2825_v103_: _dafny.MultiSet
        d_2825_v103_ = _dafny.MultiSet([d_2824_v102_, default__.fm0(self.f4, self.f4, globalState), d_2824_v102_])
        d_2826_v104_: _dafny.Set
        d_2826_v104_ = _dafny.Set({d_2825_v103_, d_2825_v103_})
        d_2824_v102_ = ((d_2826_v104_) - (_dafny.Set({_dafny.MultiSet([not((self).f8), d_2824_v102_])}))).isdisjoint(d_2826_v104_)
        if (self).f8:
            d_2827_v105_: C0
            nw485_ = C0()
            nw485_.ctor__(self.f4)
            d_2827_v105_ = nw485_
            d_2828_v106_: _dafny.Seq
            d_2828_v106_ = _dafny.SeqWithoutIsStrInference([d_2827_v105_])
            d_2829_v107_: _dafny.Seq
            d_2829_v107_ = _dafny.SeqWithoutIsStrInference([D5_DC18((d_2828_v106_)[default__.safeIndex(self.f4, len(d_2828_v106_))], d_2824_v102_, (self).f8)])
            d_2829_v107_ = ((d_2829_v107_) + (d_2829_v107_)) + ((d_2829_v107_) + (d_2829_v107_))
            d_2830_v108_: _dafny.Seq
            d_2830_v108_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qigwel"))
            d_2831_v109_: _dafny.Map
            d_2831_v109_ = _dafny.Map({d_2830_v108_: d_2827_v105_.f18})
            d_2832_v110_: _dafny.Map
            d_2832_v110_ = _dafny.Map({d_2831_v109_: (0) - ((0) - (self.f4))})
            d_2833_v111_: _dafny.Seq
            d_2833_v111_ = _dafny.SeqWithoutIsStrInference([d_2827_v105_.f18, d_2827_v105_.f18])
            d_2834_v112_: D7
            d_2834_v112_ = D7_DC27(d_2827_v105_.f18, self.f4, (0) - (len(d_2830_v108_)))
            d_2832_v110_ = (d_2832_v110_).set(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ncm")): self.f4}), (len(d_2833_v111_)) + ((d_2834_v112_).cf50))
            d_2835_v113_: _dafny.Seq
            d_2835_v113_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([True]), d_2825_v103_])
            d_2825_v103_ = (d_2825_v103_ if d_2824_v102_ else (d_2825_v103_) - ((d_2835_v113_)[default__.safeIndex(self.f4, len(d_2835_v113_))]))
            d_2824_v102_ = (self).f8
            if (self).f8:
                d_2836_v114_: _dafny.Array
                nw486_ = _dafny.Array(_dafny.Array(None, 0), 1)
                d_2836_v114_ = nw486_
                d_2837_v115_: _dafny.Set
                d_2837_v115_ = _dafny.Set({d_2824_v102_})
                d_2838_v116_: D14
                d_2838_v116_ = D14_DC41(d_2837_v115_)
                pat_let_tv65_ = d_2837_v115_
                d_2839_v117_: _dafny.Seq
                def iife262_(_pat_let81_0):
                    def iife263_(d_2840_dt__update__tmp_h1_):
                        def iife264_(_pat_let82_0):
                            def iife265_(d_2841_dt__update_hcf71_h0_):
                                return D14_DC41(d_2841_dt__update_hcf71_h0_)
                            return iife265_(_pat_let82_0)
                        return iife264_(pat_let_tv65_)
                    return iife263_(_pat_let81_0)
                d_2839_v117_ = _dafny.SeqWithoutIsStrInference([d_2838_v116_, d_2838_v116_, iife262_(d_2838_v116_)])
                rhs411_ = d_2836_v114_
                rhs412_ = d_2839_v117_
                rhs413_ = d_2830_v108_
                d_2836_v114_ = rhs411_
                d_2839_v117_ = rhs412_
                d_2830_v108_ = rhs413_
                d_2842_v119_: _dafny.Array
                def lambda138_(d_2843_v108_, d_2844_v105_):
                    def lambda139_(d_2845_i10_):
                        def iife266_():
                            coll100_ = _dafny.Map()
                            compr_100_: int
                            for compr_100_ in _dafny.IntegerRange(347, 558):
                                d_2846_v118_: int = compr_100_
                                if ((347) <= (d_2846_v118_)) and ((d_2846_v118_) < (558)):
                                    coll100_[default__.safeDivisionInt(d_2846_v118_, len(d_2843_v108_))] = d_2844_v105_.f18
                            return _dafny.Map(coll100_)
                        return (_dafny.Map({self.f4: self.f4})) | (iife266_()
                        )

                    return lambda139_

                init81_ = lambda138_(d_2830_v108_, d_2827_v105_)
                nw487_ = _dafny.Array(None, 10)
                for i0_81_ in range(nw487_.length(0)):
                    nw487_[i0_81_] = init81_(i0_81_)
                d_2842_v119_ = nw487_
                d_2847_v120_: _dafny.Map
                d_2847_v120_ = _dafny.Map({d_2827_v105_.f18: d_2827_v105_.f18})
                index443_ = default__.safeIndex(101, (d_2842_v119_).length(0))
                (d_2842_v119_)[index443_] = d_2847_v120_
                index444_ = default__.safeIndex(101, (d_2842_v119_).length(0))
                def iife267_():
                    coll101_ = _dafny.Map()
                    compr_101_: int
                    for compr_101_ in (d_2833_v111_).Elements:
                        d_2848_v121_: int = compr_101_
                        if (d_2848_v121_) in (d_2833_v111_):
                            coll101_[(d_2848_v121_) * (d_2827_v105_.f18)] = len(d_2830_v108_)
                    return _dafny.Map(coll101_)
                (d_2842_v119_)[index444_] = (iife267_()
                ) | (d_2847_v120_)
                (d_2827_v105_).f18 = self.f4
                r1 = default__.safeModuloInt(d_2827_v105_.f18, self.f4)
                d_2849_v122_: _dafny.Map
                d_2849_v122_ = _dafny.Map({d_2824_v102_: d_2827_v105_.f18})
                d_2850_v123_: _dafny.Map
                d_2850_v123_ = _dafny.Map({d_2827_v105_.f18: d_2849_v122_})
                d_2851_v124_: _dafny.Array
                nw488_ = _dafny.Array(None, 25)
                nw488_[int(0)] = d_2849_v122_
                nw488_[int(1)] = d_2849_v122_
                nw488_[int(2)] = d_2849_v122_
                nw488_[int(3)] = d_2849_v122_
                nw488_[int(4)] = _dafny.Map({d_2824_v102_: default__.fm1(globalState)})
                nw488_[int(5)] = _dafny.Map({(self).f8: self.f4})
                nw488_[int(6)] = d_2849_v122_
                nw488_[int(7)] = d_2849_v122_
                nw488_[int(8)] = d_2849_v122_
                nw488_[int(9)] = _dafny.Map({default__.fm0(d_2827_v105_.f18, 145, globalState): self.f4})
                nw488_[int(10)] = d_2849_v122_
                nw488_[int(11)] = _dafny.Map({(self).f8: self.f4})
                nw488_[int(12)] = _dafny.Map({(self).f8: self.f4})
                nw488_[int(13)] = d_2849_v122_
                nw488_[int(14)] = d_2849_v122_
                nw488_[int(15)] = d_2849_v122_
                nw488_[int(16)] = d_2849_v122_
                nw488_[int(17)] = d_2849_v122_
                nw488_[int(18)] = d_2849_v122_
                nw488_[int(19)] = (d_2849_v122_).set(d_2824_v102_, len(d_2830_v108_))
                nw488_[int(20)] = (d_2849_v122_).set((self).fm5(d_2830_v108_, (self).f8, globalState), d_2827_v105_.f18)
                nw488_[int(21)] = ((d_2850_v123_)[d_2827_v105_.f18] if (d_2827_v105_.f18) in (d_2850_v123_) else d_2849_v122_)
                nw488_[int(22)] = d_2849_v122_
                nw488_[int(23)] = d_2849_v122_
                nw488_[int(24)] = _dafny.Map({(self).f8: -777})
                d_2851_v124_ = nw488_
                d_2852_v125_: D23
                d_2852_v125_ = D23_DC60(d_2851_v124_)
                d_2852_v125_ = d_2852_v125_
            elif True:
                d_2830_v108_ = d_2830_v108_
                d_2853_v126_: _dafny.MultiSet
                d_2853_v126_ = _dafny.MultiSet([d_2827_v105_.f18])
                d_2854_v127_: _dafny.Array
                nw489_ = _dafny.Array(False, 20)
                d_2854_v127_ = nw489_
                rhs414_ = ((d_2853_v126_) - (d_2853_v126_)) | ((d_2853_v126_) | (d_2853_v126_))
                rhs415_ = len((d_2830_v108_) + (d_2830_v108_))
                rhs416_ = d_2854_v127_
                d_2853_v126_ = rhs414_
                r1 = rhs415_
                d_2854_v127_ = rhs416_
                d_2855_v128_: D3
                d_2855_v128_ = D3_DC11(D3_DC8(_dafny.CodePoint('t'), d_2854_v127_))
                d_2856_v129_: _dafny.Set
                d_2856_v129_ = _dafny.Set({785})
                d_2857_v130_: D3
                d_2857_v130_ = D3_DC11(D3_DC7(d_2856_v129_))
                pat_let_tv66_ = d_2857_v130_
                d_2858_v131_: _dafny.Array
                nw490_ = _dafny.Array(None, 9)
                nw490_[int(0)] = d_2855_v128_
                nw490_[int(1)] = d_2855_v128_
                nw490_[int(2)] = d_2855_v128_
                nw490_[int(3)] = d_2855_v128_
                nw490_[int(4)] = d_2855_v128_
                nw490_[int(5)] = d_2855_v128_
                nw490_[int(6)] = d_2855_v128_
                def iife268_(_pat_let83_0):
                    def iife269_(d_2859_dt__update__tmp_h2_):
                        def iife270_(_pat_let84_0):
                            def iife271_(d_2860_dt__update_hcf22_h0_):
                                return D3_DC11(d_2860_dt__update_hcf22_h0_)
                            return iife271_(_pat_let84_0)
                        return iife270_(pat_let_tv66_)
                    return iife269_(_pat_let83_0)
                nw490_[int(7)] = iife268_(D3_DC11(d_2857_v130_))
                nw490_[int(8)] = d_2855_v128_
                d_2858_v131_ = nw490_
                index445_ = default__.safeIndex(808, (d_2858_v131_).length(0))
                (d_2858_v131_)[index445_] = D3_DC11(d_2857_v130_)
                d_2861_v132_: _dafny.Set
                d_2861_v132_ = _dafny.Set({(d_2830_v108_) + (d_2830_v108_)})
                index446_ = default__.safeIndex(527, (d_2854_v127_).length(0))
                (d_2854_v127_)[index446_] = d_2824_v102_
                d_2862_v133_: _dafny.Map
                d_2862_v133_ = _dafny.Map({516: (d_2853_v126_) - (d_2853_v126_)})
                d_2863_v134_: _dafny.Seq
                d_2863_v134_ = _dafny.SeqWithoutIsStrInference([d_2853_v126_])
                index447_ = default__.safeIndex(808, (d_2858_v131_).length(0))
                index448_ = default__.safeIndex(527, (d_2854_v127_).length(0))
                rhs417_ = d_2855_v128_
                rhs418_ = ((d_2862_v133_)[self.f4] if (self.f4) in (d_2862_v133_) else d_2853_v126_)
                rhs419_ = (d_2861_v132_) | (d_2861_v132_)
                rhs420_ = ((d_2863_v134_)[default__.safeIndex((0) - (d_2827_v105_.f18), len(d_2863_v134_))]).cardinality
                rhs421_ = default__.fm0(self.f4, d_2827_v105_.f18, globalState)
                lhs249_ = d_2858_v131_
                lhs250_ = default__.safeIndex(808, (d_2858_v131_).length(0))
                lhs251_ = self
                lhs252_ = d_2854_v127_
                lhs253_ = default__.safeIndex(527, (d_2854_v127_).length(0))
                lhs249_[lhs250_] = rhs417_
                d_2853_v126_ = rhs418_
                d_2861_v132_ = rhs419_
                lhs251_.f4 = rhs420_
                lhs252_[lhs253_] = rhs421_
                d_2864_v135_: _dafny.Map
                d_2864_v135_ = _dafny.Map({(d_2854_v127_)[default__.safeIndex(527, (d_2854_v127_).length(0))]: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rocniknf"))})
                d_2865_v136_: _dafny.Array
                nw491_ = _dafny.Array(None, 29)
                nw491_[int(0)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hegdluem"))) + (d_2830_v108_)
                nw491_[int(1)] = d_2830_v108_
                nw491_[int(2)] = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_2866_i11_ in range(default__.abs(688))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iytidt")))
                nw491_[int(3)] = d_2830_v108_
                nw491_[int(4)] = d_2830_v108_
                nw491_[int(5)] = d_2830_v108_
                nw491_[int(6)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ja"))
                nw491_[int(7)] = (_dafny.SeqWithoutIsStrInference([(d_2830_v108_)[default__.safeIndex(533, len(d_2830_v108_))] for d_2867_i12_ in range(default__.abs(107))])) + (d_2830_v108_)
                nw491_[int(8)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_2868_i13_ in range(default__.abs(969))])
                nw491_[int(9)] = d_2830_v108_
                nw491_[int(10)] = d_2830_v108_
                nw491_[int(11)] = (d_2830_v108_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rbowqac")))
                nw491_[int(12)] = d_2830_v108_
                nw491_[int(13)] = (d_2830_v108_) + (((d_2864_v135_)[(d_2854_v127_)[default__.safeIndex(527, (d_2854_v127_).length(0))]] if ((d_2854_v127_)[default__.safeIndex(527, (d_2854_v127_).length(0))]) in (d_2864_v135_) else d_2830_v108_))
                nw491_[int(14)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_2869_i14_ in range(default__.abs(307))])
                nw491_[int(15)] = d_2830_v108_
                nw491_[int(16)] = d_2830_v108_
                nw491_[int(17)] = d_2830_v108_
                nw491_[int(18)] = (d_2830_v108_) + (d_2830_v108_)
                nw491_[int(19)] = d_2830_v108_
                nw491_[int(20)] = d_2830_v108_
                nw491_[int(21)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_2870_i15_ in range(default__.abs(327))])
                nw491_[int(22)] = d_2830_v108_
                nw491_[int(23)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_2871_i16_ in range(default__.abs(-658))])
                nw491_[int(24)] = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_2872_i17_ in range(default__.abs(681))])) + (d_2830_v108_)
                nw491_[int(25)] = d_2830_v108_
                nw491_[int(26)] = d_2830_v108_
                nw491_[int(27)] = (d_2830_v108_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_2873_i18_ in range(default__.abs(373))]))
                nw491_[int(28)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_2874_i19_ in range(default__.abs(972))])
                d_2865_v136_ = nw491_
                d_2875_v137_: _dafny.Seq
                d_2875_v137_ = _dafny.SeqWithoutIsStrInference([d_2830_v108_])
                index449_ = default__.safeIndex(593, (d_2865_v136_).length(0))
                (d_2865_v136_)[index449_] = (d_2875_v137_)[default__.safeIndex(self.f4, len(d_2875_v137_))]
                index450_ = default__.safeIndex(593, (d_2865_v136_).length(0))
                (d_2865_v136_)[index450_] = d_2830_v108_
                d_2876_v138_: D2
                d_2876_v138_ = D2_DC6(self.f4, True)
                d_2877_v139_: _dafny.Seq
                d_2877_v139_ = _dafny.SeqWithoutIsStrInference([d_2824_v102_, not(d_2824_v102_)])
                d_2878_v140_: _dafny.Set
                d_2878_v140_ = _dafny.Set({(self).f8, (self).f8})
                d_2879_v141_: T1
                nw492_ = C4()
                nw492_.ctor__((0) - (((d_2853_v126_)[len(d_2878_v140_)] if (len(d_2878_v140_)) in (d_2853_v126_) else self.f4)))
                d_2879_v141_ = nw492_
                d_2880_v142_: _dafny.Map
                d_2880_v142_ = _dafny.Map({d_2879_v141_: (d_2854_v127_)[default__.safeIndex(527, (d_2854_v127_).length(0))]})
                d_2881_v143_: _dafny.Array
                nw493_ = _dafny.Array(None, 17)
                nw493_[int(0)] = d_2876_v138_
                nw493_[int(1)] = d_2876_v138_
                nw493_[int(2)] = d_2876_v138_
                nw493_[int(3)] = d_2876_v138_
                nw493_[int(4)] = d_2876_v138_
                def iife272_(_pat_let85_0):
                    def iife273_(d_2882_dt__update__tmp_h3_):
                        def iife274_(_pat_let86_0):
                            def iife275_(d_2883_dt__update_hcf11_h0_):
                                return D2_DC6((d_2882_dt__update__tmp_h3_).cf10, d_2883_dt__update_hcf11_h0_)
                            return iife275_(_pat_let86_0)
                        return iife274_((self).f8)
                    return iife273_(_pat_let85_0)
                nw493_[int(5)] = iife272_(default__.fm47(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l")), D3_DC10(len(d_2877_v139_), d_2827_v105_.f18, (self).f8, self.f4, True), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r")), d_2824_v102_, globalState))
                nw493_[int(6)] = d_2876_v138_
                nw493_[int(7)] = d_2876_v138_
                nw493_[int(8)] = d_2876_v138_
                nw493_[int(9)] = d_2876_v138_
                nw493_[int(10)] = d_2876_v138_
                nw493_[int(11)] = default__.fm47((d_2865_v136_)[default__.safeIndex(593, (d_2865_v136_).length(0))], default__.fm20(False, d_2856_v129_, globalState), (d_2875_v137_)[default__.safeIndex(d_2827_v105_.f18, len(d_2875_v137_))], (d_2854_v127_)[default__.safeIndex(527, (d_2854_v127_).length(0))], globalState)
                nw493_[int(12)] = d_2876_v138_
                nw493_[int(13)] = d_2876_v138_
                nw493_[int(14)] = d_2876_v138_
                nw493_[int(15)] = D2_DC6(self.f4, not(default__.fm0(len(d_2880_v142_), self.f4, globalState)))
                nw493_[int(16)] = d_2876_v138_
                d_2881_v143_ = nw493_
                index451_ = default__.safeIndex(525, (d_2881_v143_).length(0))
                (d_2881_v143_)[index451_] = d_2876_v138_
                d_2884_v144_: D3
                d_2884_v144_ = D3_DC10(self.f4, self.f4, True, default__.fm1(globalState), (d_2854_v127_)[default__.safeIndex(527, (d_2854_v127_).length(0))])
                index452_ = default__.safeIndex(527, (d_2854_v127_).length(0))
                index453_ = default__.safeIndex(593, (d_2865_v136_).length(0))
                index454_ = default__.safeIndex(525, (d_2881_v143_).length(0))
                rhs422_ = (self).f8
                rhs423_ = (d_2833_v111_) + ((d_2833_v111_).set(default__.safeIndex(d_2879_v141_.f4, len(d_2833_v111_)), (0) - (d_2879_v141_.f4)))
                rhs424_ = ((d_2833_v111_)[default__.safeIndex(default__.fm1(globalState), len(d_2833_v111_))]) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w"))))
                rhs425_ = default__.fm23(d_2824_v102_, globalState)
                rhs426_ = (d_2876_v138_ if (d_2884_v144_).cf21 else default__.fm47(default__.fm23((self).f8, globalState), d_2884_v144_, d_2830_v108_, (self).f8, globalState))
                lhs254_ = d_2854_v127_
                lhs255_ = default__.safeIndex(527, (d_2854_v127_).length(0))
                lhs256_ = d_2827_v105_
                lhs257_ = d_2865_v136_
                lhs258_ = default__.safeIndex(593, (d_2865_v136_).length(0))
                lhs259_ = d_2881_v143_
                lhs260_ = default__.safeIndex(525, (d_2881_v143_).length(0))
                lhs254_[lhs255_] = rhs422_
                d_2833_v111_ = rhs423_
                lhs256_.f18 = rhs424_
                lhs257_[lhs258_] = rhs425_
                lhs259_[lhs260_] = rhs426_
        elif True:
            d_2824_v102_ = (d_2824_v102_) or (d_2824_v102_)
            d_2885_v145_: _dafny.MultiSet
            d_2885_v145_ = _dafny.MultiSet([self.f4])
            d_2886_v146_: _dafny.MultiSet
            d_2886_v146_ = d_2885_v145_
            d_2887_v147_: _dafny.MultiSet
            d_2887_v147_ = _dafny.MultiSet([(d_2886_v146_ if (self).f8 else d_2886_v146_), d_2886_v146_, d_2886_v146_, d_2886_v146_, d_2886_v146_])
            d_2887_v147_ = ((_dafny.MultiSet([d_2886_v146_, d_2885_v145_, d_2885_v145_, d_2885_v145_])) | (d_2887_v147_)).set(default__.fm48(globalState), default__.abs(self.f4))
            r1 = (0) - (default__.fm1(globalState))
            d_2824_v102_ = (self).f8
            d_2888_v148_: _dafny.Array
            nw494_ = _dafny.Array(int(0), 14)
            d_2888_v148_ = nw494_
            index455_ = default__.safeIndex(770, (d_2888_v148_).length(0))
            (d_2888_v148_)[index455_] = default__.fm1(globalState)
            d_2889_v149_: _dafny.Seq
            d_2889_v149_ = _dafny.SeqWithoutIsStrInference([self.f4])
            index456_ = default__.safeIndex(770, (d_2888_v148_).length(0))
            (d_2888_v148_)[index456_] = ((d_2889_v149_)[default__.safeIndex(self.f4, len(d_2889_v149_))]) - (self.f4)
        d_2890_v150_: str
        d_2890_v150_ = _dafny.CodePoint('t')
        d_2891_v151_: _dafny.Seq
        d_2891_v151_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xbfml"))
        d_2892_v152_: _dafny.Map
        d_2892_v152_ = _dafny.Map({self.f4: d_2890_v150_})
        d_2893_v153_: _dafny.Seq
        d_2893_v153_ = _dafny.SeqWithoutIsStrInference([not((self).f8), d_2824_v102_])
        d_2894_v154_: _dafny.Array
        nw495_ = _dafny.Array(None, 29)
        nw495_[int(0)] = d_2890_v150_
        nw495_[int(1)] = d_2890_v150_
        nw495_[int(2)] = _dafny.CodePoint('c')
        nw495_[int(3)] = _dafny.CodePoint('i')
        nw495_[int(4)] = (d_2890_v150_ if d_2824_v102_ else d_2890_v150_)
        nw495_[int(5)] = d_2890_v150_
        nw495_[int(6)] = d_2890_v150_
        nw495_[int(7)] = d_2890_v150_
        nw495_[int(8)] = (d_2891_v151_)[default__.safeIndex(833, len(d_2891_v151_))]
        nw495_[int(9)] = d_2890_v150_
        nw495_[int(10)] = ((d_2892_v152_)[self.f4] if (self.f4) in (d_2892_v152_) else default__.fm27(d_2824_v102_, globalState))
        nw495_[int(11)] = d_2890_v150_
        nw495_[int(12)] = _dafny.CodePoint('i')
        nw495_[int(13)] = d_2890_v150_
        nw495_[int(14)] = d_2890_v150_
        nw495_[int(15)] = _dafny.CodePoint('c')
        nw495_[int(16)] = _dafny.CodePoint('y')
        nw495_[int(17)] = d_2890_v150_
        nw495_[int(18)] = d_2890_v150_
        nw495_[int(19)] = (d_2891_v151_)[default__.safeIndex(len(d_2893_v153_), len(d_2891_v151_))]
        nw495_[int(20)] = d_2890_v150_
        nw495_[int(21)] = d_2890_v150_
        nw495_[int(22)] = d_2890_v150_
        nw495_[int(23)] = d_2890_v150_
        nw495_[int(24)] = d_2890_v150_
        nw495_[int(25)] = _dafny.CodePoint('g')
        nw495_[int(26)] = d_2890_v150_
        nw495_[int(27)] = d_2890_v150_
        nw495_[int(28)] = d_2890_v150_
        d_2894_v154_ = nw495_
        guard_loop_9_: int
        for guard_loop_9_ in _dafny.IntegerRange(0, (d_2894_v154_).length(0)):
            d_2895_i20_: int = guard_loop_9_
            if (True) and (((0) <= (d_2895_i20_)) and ((d_2895_i20_) < ((d_2894_v154_).length(0)))):
                (d_2894_v154_)[(d_2895_i20_)] = ((d_2891_v151_)[default__.safeIndex(self.f4, len(d_2891_v151_))] if (self).f8 else d_2890_v150_)
        d_2896_v155_: _dafny.Array
        nw496_ = _dafny.Array(int(0), 16)
        d_2896_v155_ = nw496_
        d_2897_v156_: _dafny.MultiSet
        d_2897_v156_ = _dafny.MultiSet([d_2896_v155_])
        d_2897_v156_ = d_2897_v156_
        d_2898_v157_: _dafny.Seq
        d_2898_v157_ = _dafny.SeqWithoutIsStrInference([-714])
        d_2899_v158_: _dafny.Seq
        d_2899_v158_ = _dafny.SeqWithoutIsStrInference([d_2898_v157_])
        r0 = (default__.fm49((self).f8, (self).f8, self.f4, globalState) if not(False) else (d_2899_v158_) + (d_2899_v158_))
        r1 = (0) - (default__.safeModuloInt(self.f4, self.f4))
        return r0, r1

    @property
    def f8(self):
        return self._f8
    @property
    def f9(self):
        return self._f9

class C7(T0):
    def  __init__(self):
        self._f4: int = int(0)
        self.f7: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        pass

    def __dafnystr__(self) -> str:
        return "_module.C7"
    @property
    def f4(self):
        return self._f4
    @f4.setter
    def f4(self, value):
        self._f4 = value
    def ctor__(self, f7, f4):
        (self).f7 = f7
        (self).f4 = f4

    def fm5(self, p0, p1, globalState):
        return (self.f4) <= (self.f4)

    def fm7(self, p0, p1, p2, globalState):
        return ((_dafny.MultiSet([_dafny.Map({True: 194}), _dafny.Map({True: len(_dafny.Set({True}))})])) - (_dafny.MultiSet([_dafny.Map({False: self.f4}), _dafny.Map({True: self.f4})]))) | ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(D1_DC3(_dafny.Map({False: 0}))).cf5]))) - (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.Map({False: self.f4})]))))

    def m3(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        hi16_ = p1
        for d_2900_i0_ in range(p1, hi16_):
            d_2901_v0_: str
            d_2901_v0_ = _dafny.CodePoint('t')
            d_2901_v0_ = d_2901_v0_
            d_2902_v1_: _dafny.Array
            def lambda140_(d_2903_i0_, d_2904_p2_):
                def lambda141_(d_2905_i1_):
                    return (_dafny.MultiSet([self.f4, d_2903_i0_])).intersection(_dafny.MultiSet([d_2903_i0_, d_2903_i0_, (_dafny.MultiSet([d_2904_p2_])).cardinality, 35]))

                return lambda141_

            init82_ = lambda140_(d_2900_i0_, p2)
            nw497_ = _dafny.Array(None, 18)
            for i0_82_ in range(nw497_.length(0)):
                nw497_[i0_82_] = init82_(i0_82_)
            d_2902_v1_ = nw497_
            d_2902_v1_ = d_2902_v1_
            d_2906_v2_: _dafny.Seq
            d_2906_v2_ = _dafny.SeqWithoutIsStrInference([not(p3)])
            d_2907_v3_: _dafny.Map
            d_2907_v3_ = _dafny.Map({(self.f4 if True else len(d_2906_v2_)): p0})
            d_2907_v3_ = (d_2907_v3_).set(p1, p3)
            d_2908_v4_: _dafny.Map
            d_2908_v4_ = _dafny.Map({p3: p0})
            d_2909_v5_: _dafny.Array
            nw498_ = _dafny.Array(None, 17)
            nw498_[int(0)] = d_2900_i0_
            nw498_[int(1)] = d_2900_i0_
            nw498_[int(2)] = len(d_2908_v4_)
            nw498_[int(3)] = d_2900_i0_
            nw498_[int(4)] = self.f4
            nw498_[int(5)] = self.f4
            nw498_[int(6)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jyssxtk")))
            nw498_[int(7)] = p1
            nw498_[int(8)] = p1
            nw498_[int(9)] = (0) - (p1)
            nw498_[int(10)] = 617
            nw498_[int(11)] = (0) - (p1)
            nw498_[int(12)] = p1
            nw498_[int(13)] = self.f4
            nw498_[int(14)] = self.f4
            nw498_[int(15)] = 905
            nw498_[int(16)] = p1
            d_2909_v5_ = nw498_
            d_2910_v6_: _dafny.Map
            d_2910_v6_ = _dafny.Map({self.f4: d_2901_v0_})
            d_2911_v7_: _dafny.Map
            d_2911_v7_ = _dafny.Map({d_2909_v5_: d_2910_v6_})
            d_2911_v7_ = d_2911_v7_
        d_2912_v8_: _dafny.Map
        d_2912_v8_ = _dafny.Map({self.f4: p1})
        d_2913_v9_: _dafny.MultiSet
        d_2913_v9_ = _dafny.MultiSet([p1, self.f4, p1, p1, self.f4])
        d_2914_v10_: str
        d_2914_v10_ = _dafny.CodePoint('b')
        d_2915_v11_: T0
        nw499_ = C3()
        nw499_.ctor__(p1, self.f7, 145)
        d_2915_v11_ = nw499_
        d_2916_v12_: _dafny.Map
        d_2916_v12_ = _dafny.Map({d_2915_v11_: p0})
        d_2917_v13_: _dafny.MultiSet
        d_2917_v13_ = _dafny.MultiSet([p3])
        d_2918_v14_: D5
        d_2918_v14_ = D5_DC17(d_2917_v13_, p2)
        d_2919_v15_: _dafny.Seq
        d_2919_v15_ = _dafny.SeqWithoutIsStrInference([p3, p0])
        d_2920_v16_: _dafny.Map
        d_2920_v16_ = _dafny.Map({p2: p1})
        d_2921_v17_: _dafny.Map
        d_2921_v17_ = _dafny.Map({d_2920_v16_: p3})
        pat_let_tv67_ = d_2917_v13_
        d_2922_v18_: _dafny.Array
        nw500_ = _dafny.Array(None, 18)
        nw500_[int(0)] = ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_2923_i2_ in range(default__.abs(698))])).set(default__.safeIndex(((d_2912_v8_)[(d_2913_v9_).cardinality] if ((d_2913_v9_).cardinality) in (d_2912_v8_) else self.f4), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_2924_i2_ in range(default__.abs(698))]))), d_2914_v10_)) != (self.f7)
        nw500_[int(1)] = p2
        nw500_[int(2)] = False
        nw500_[int(3)] = True
        nw500_[int(4)] = (self).fm5(self.f7, p3, globalState)
        nw500_[int(5)] = not((d_2915_v11_) in ((d_2916_v12_).set(d_2915_v11_, p2)))
        def iife276_(_pat_let87_0):
            def iife277_(d_2925_dt__update__tmp_h0_):
                def iife278_(_pat_let88_0):
                    def iife279_(d_2926_dt__update_hcf34_h0_):
                        return D5_DC17(d_2926_dt__update_hcf34_h0_, (d_2925_dt__update__tmp_h0_).cf35)
                    return iife279_(_pat_let88_0)
                return iife278_(pat_let_tv67_)
            return iife277_(_pat_let87_0)
        nw500_[int(6)] = ((iife276_(d_2918_v14_)).cf34).ispropersubset(d_2917_v13_)
        nw500_[int(7)] = p0
        nw500_[int(8)] = not (True) or (p2)
        nw500_[int(9)] = (d_2915_v11_.f4) <= ((0) - (d_2915_v11_.f4))
        nw500_[int(10)] = p3
        nw500_[int(11)] = p3
        nw500_[int(12)] = p0
        nw500_[int(13)] = (d_2915_v11_.f4) == (self.f4)
        nw500_[int(14)] = p3
        nw500_[int(15)] = not((_dafny.SeqWithoutIsStrInference([p0])) <= (d_2919_v15_))
        nw500_[int(16)] = p0
        nw500_[int(17)] = ((d_2921_v17_)[default__.fm28(self.f4, globalState)] if (default__.fm28(self.f4, globalState)) in (d_2921_v17_) else p0)
        d_2922_v18_ = nw500_
        index457_ = default__.safeIndex(82, (d_2922_v18_).length(0))
        (d_2922_v18_)[index457_] = (d_2915_v11_.f4) > (len(self.f7))
        index458_ = default__.safeIndex(82, (d_2922_v18_).length(0))
        (d_2922_v18_)[index458_] = p3
        index459_ = default__.safeIndex(82, (d_2922_v18_).length(0))
        (d_2922_v18_)[index459_] = p0
        d_2912_v8_ = (d_2912_v8_).set(p1, len(_dafny.SeqWithoutIsStrInference([p1])))
        d_2927_i3_: int
        d_2927_i3_ = 0
        with _dafny.label("10"):
            while p0:
                with _dafny.c_label("10"):
                    if (d_2927_i3_) >= (100):
                        raise _dafny.Break("10")
                    d_2927_i3_ = (d_2927_i3_) + (1)
                    d_2928_v19_: _dafny.Seq
                    d_2928_v19_ = _dafny.SeqWithoutIsStrInference([len(_dafny.Set({(d_2922_v18_)[default__.safeIndex(82, (d_2922_v18_).length(0))], p0})), p1])
                    d_2929_v20_: _dafny.Seq
                    d_2929_v20_ = _dafny.SeqWithoutIsStrInference([d_2915_v11_.f4, (0) - (p1), (d_2928_v19_)[default__.safeIndex(d_2915_v11_.f4, len(d_2928_v19_))], len(d_2920_v16_)])
                    d_2930_v21_: D0
                    d_2930_v21_ = D0_DC0(d_2928_v19_)
                    source31_ = d_2930_v21_
                    if source31_.is_DC1:
                        d_2931___mcc_h0_ = source31_.cf1
                        d_2932___mcc_h1_ = source31_.cf2
                        d_2933___mcc_h2_ = source31_.cf3
                        d_2934_cf3_ = d_2933___mcc_h2_
                        d_2935_cf2_ = d_2932___mcc_h1_
                        d_2936_cf1_ = d_2931___mcc_h0_
                        d_2936_cf1_ = ((0) - (self.f4)) + (default__.safeDivisionInt(d_2915_v11_.f4, len(d_2920_v16_)))
                        d_2937_v22_: _dafny.Seq
                        d_2937_v22_ = _dafny.SeqWithoutIsStrInference([self.f7])
                        r0 = default__.fm0(default__.safeDivisionInt((d_2917_v13_).cardinality, len((d_2937_v22_)[default__.safeIndex(d_2936_cf1_, len(d_2937_v22_))])), default__.safeDivisionInt((0) - (d_2915_v11_.f4), d_2915_v11_.f4), globalState)
                        d_2938_v23_: _dafny.Seq
                        d_2938_v23_ = _dafny.SeqWithoutIsStrInference([d_2919_v15_])
                        index460_ = default__.safeIndex(82, (d_2922_v18_).length(0))
                        index461_ = default__.safeIndex(82, (d_2922_v18_).length(0))
                        index462_ = default__.safeIndex(82, (d_2922_v18_).length(0))
                        index463_ = default__.safeIndex(82, (d_2922_v18_).length(0))
                        rhs427_ = False
                        rhs428_ = ((d_2938_v23_)[default__.safeIndex((0) - ((d_2913_v9_).cardinality), len(d_2938_v23_))]) + ((d_2919_v15_).set(default__.safeIndex(self.f4, len(d_2919_v15_)), p3))
                        rhs429_ = p3
                        rhs430_ = not(p3)
                        rhs431_ = not((True if p0 else (d_2922_v18_)[default__.safeIndex(82, (d_2922_v18_).length(0))]))
                        lhs261_ = d_2922_v18_
                        lhs262_ = default__.safeIndex(82, (d_2922_v18_).length(0))
                        lhs263_ = d_2922_v18_
                        lhs264_ = default__.safeIndex(82, (d_2922_v18_).length(0))
                        lhs265_ = d_2922_v18_
                        lhs266_ = default__.safeIndex(82, (d_2922_v18_).length(0))
                        lhs267_ = d_2922_v18_
                        lhs268_ = default__.safeIndex(82, (d_2922_v18_).length(0))
                        lhs261_[lhs262_] = rhs427_
                        d_2919_v15_ = rhs428_
                        lhs263_[lhs264_] = rhs429_
                        lhs265_[lhs266_] = rhs430_
                        lhs267_[lhs268_] = rhs431_
                        d_2914_v10_ = d_2914_v10_
                    elif source31_.is_DC0:
                        d_2939___mcc_h3_ = source31_.cf0
                        d_2940_cf0_ = d_2939___mcc_h3_
                        d_2941_v24_: T2
                        nw501_ = C1()
                        nw501_.ctor__(d_2915_v11_.f4, default__.safeModuloInt(205, d_2915_v11_.f4), self.f4, True)
                        d_2941_v24_ = nw501_
                        d_2941_v24_ = ((d_2941_v24_ if p2 else d_2941_v24_) if not((default__.fm0(365, len(d_2928_v19_), globalState) if p2 else p0)) else d_2941_v24_)
                        (d_2915_v11_).f4 = d_2915_v11_.f4
                        d_2919_v15_ = d_2919_v15_
                        d_2942_v25_: _dafny.Seq
                        d_2942_v25_ = _dafny.SeqWithoutIsStrInference([d_2928_v19_, d_2929_v20_])
                        d_2943_v26_: C2
                        nw502_ = C2()
                        nw502_.ctor__((d_2941_v24_).f16, len(((d_2942_v25_)[default__.safeIndex((d_2941_v24_).f15, len(d_2942_v25_))]) + (_dafny.SeqWithoutIsStrInference([self.f4 for d_2944_i4_ in range(default__.abs(-835))]))))
                        d_2943_v26_ = nw502_
                    elif True:
                        d_2945___mcc_h4_ = source31_.cf4
                        d_2946_cf4_ = d_2945___mcc_h4_
                        d_2947_v27_: _dafny.Seq
                        d_2947_v27_ = _dafny.SeqWithoutIsStrInference([d_2919_v15_, d_2919_v15_, d_2919_v15_])
                        d_2948_v28_: C0
                        nw503_ = C0()
                        nw503_.ctor__(len(((d_2947_v27_)[default__.safeIndex(921, len(d_2947_v27_))]) + (_dafny.SeqWithoutIsStrInference([(d_2922_v18_)[default__.safeIndex(82, (d_2922_v18_).length(0))]]))))
                        d_2948_v28_ = nw503_
                        d_2949_v29_: _dafny.Map
                        d_2949_v29_ = _dafny.Map({p2: d_2914_v10_})
                        index464_ = default__.safeIndex(82, (d_2922_v18_).length(0))
                        (d_2922_v18_)[index464_] = (False) in ((d_2949_v29_) | (d_2949_v29_))
                        d_2950_v30_: _dafny.Map
                        d_2950_v30_ = _dafny.Map({self.f7: default__.fm1(globalState)})
                        d_2951_v31_: _dafny.Set
                        d_2951_v31_ = _dafny.Set({len(d_2919_v15_), len(d_2919_v15_), d_2915_v11_.f4})
                        rhs432_ = d_2950_v30_
                        rhs433_ = default__.safeDivisionInt(((d_2920_v16_)[(d_2922_v18_)[default__.safeIndex(82, (d_2922_v18_).length(0))]] if ((d_2922_v18_)[default__.safeIndex(82, (d_2922_v18_).length(0))]) in (d_2920_v16_) else len(d_2951_v31_)), p1)
                        rhs434_ = p1
                        rhs435_ = p0
                        lhs269_ = d_2915_v11_
                        lhs270_ = d_2915_v11_
                        d_2950_v30_ = rhs432_
                        lhs269_.f4 = rhs433_
                        lhs270_.f4 = rhs434_
                        r0 = rhs435_
                        r0 = p3
                    d_2917_v13_ = d_2917_v13_
                    d_2952_v32_: _dafny.Map
                    d_2952_v32_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cmjfifo")): _dafny.MultiSet(self.f7)})
                    d_2953_v33_: _dafny.Map
                    d_2953_v33_ = _dafny.Map({(d_2922_v18_)[default__.safeIndex(82, (d_2922_v18_).length(0))]: (d_2919_v15_)[default__.safeIndex(len(d_2952_v32_), len(d_2919_v15_))]})
                    d_2954_v34_: _dafny.Map
                    d_2954_v34_ = _dafny.Map({default__.safeDivisionInt(d_2915_v11_.f4, d_2915_v11_.f4): ((d_2953_v33_)[not(False)] if (not(False)) in (d_2953_v33_) else p2)})
                    d_2954_v34_ = (d_2954_v34_).set((self.f4) + (d_2915_v11_.f4), (d_2922_v18_)[default__.safeIndex(82, (d_2922_v18_).length(0))])
                    d_2955_v35_: _dafny.Array
                    nw504_ = _dafny.Array(_dafny.Array(None, 0), 18)
                    d_2955_v35_ = nw504_
                    d_2956_v36_: _dafny.Map
                    d_2956_v36_ = _dafny.Map({d_2914_v10_: d_2915_v11_.f4})
                    d_2957_v38_: _dafny.Array
                    nw505_ = _dafny.Array(None, 10)
                    nw505_[int(0)] = _dafny.Map({d_2914_v10_: len(self.f7)})
                    nw505_[int(1)] = d_2956_v36_
                    nw505_[int(2)] = d_2956_v36_
                    nw505_[int(3)] = _dafny.Map({d_2914_v10_: (0) - (d_2915_v11_.f4)})
                    nw505_[int(4)] = default__.fm50(p0, p2, d_2915_v11_.f4, globalState)
                    nw505_[int(5)] = d_2956_v36_
                    nw505_[int(6)] = (d_2956_v36_).set(d_2914_v10_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qqdnepps"))))
                    nw505_[int(7)] = d_2956_v36_
                    nw505_[int(8)] = d_2956_v36_
                    def iife280_():
                        coll102_ = _dafny.Map()
                        compr_102_: str
                        for compr_102_ in (self.f7).Elements:
                            d_2958_v37_: str = compr_102_
                            if (d_2958_v37_) in (self.f7):
                                coll102_[d_2958_v37_] = d_2915_v11_.f4
                        return _dafny.Map(coll102_)
                    nw505_[int(9)] = (iife280_()
                    ).set(_dafny.CodePoint('t'), d_2915_v11_.f4)
                    d_2957_v38_ = nw505_
                    index465_ = default__.safeIndex(928, (d_2955_v35_).length(0))
                    (d_2955_v35_)[index465_] = d_2957_v38_
                    index466_ = default__.safeIndex(928, (d_2955_v35_).length(0))
                    (d_2955_v35_)[index466_] = d_2957_v38_
                    pass
            pass
        d_2959_v39_: _dafny.Array
        nw506_ = _dafny.Array(None, 2)
        nw506_[int(0)] = len(self.f7)
        nw506_[int(1)] = p1
        d_2959_v39_ = nw506_
        guard_loop_10_: int
        for guard_loop_10_ in _dafny.IntegerRange(0, (d_2959_v39_).length(0)):
            d_2960_i5_: int = guard_loop_10_
            if (True) and (((0) <= (d_2960_i5_)) and ((d_2960_i5_) < ((d_2959_v39_).length(0)))):
                (d_2959_v39_)[(d_2960_i5_)] = default__.safeModuloInt(d_2960_i5_, self.f4)
        r0 = not(((d_2922_v18_)[default__.safeIndex(82, (d_2922_v18_).length(0))]) and (p0))
        return r0

    def m4(self, p0, p1, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r2: _dafny.Array = _dafny.Array(None, 0)
        hi17_ = p0
        for d_2961_i0_ in range(707, hi17_):
            d_2962_v0_: _dafny.Seq
            d_2962_v0_ = _dafny.SeqWithoutIsStrInference([default__.fm1(globalState)])
            if (default__.safeModuloInt(len(d_2962_v0_), self.f4)) < (523):
                d_2963_v1_: _dafny.Array
                nw507_ = _dafny.Array(None, 28)
                nw507_[int(0)] = p1
                nw507_[int(1)] = True
                nw507_[int(2)] = p1
                nw507_[int(3)] = p1
                nw507_[int(4)] = p1
                nw507_[int(5)] = p1
                nw507_[int(6)] = p1
                nw507_[int(7)] = True
                nw507_[int(8)] = p1
                nw507_[int(9)] = p1
                nw507_[int(10)] = p1
                nw507_[int(11)] = p1
                nw507_[int(12)] = p1
                nw507_[int(13)] = p1
                nw507_[int(14)] = p1
                nw507_[int(15)] = p1
                nw507_[int(16)] = p1
                nw507_[int(17)] = p1
                nw507_[int(18)] = p1
                nw507_[int(19)] = p1
                nw507_[int(20)] = True
                nw507_[int(21)] = p1
                nw507_[int(22)] = p1
                nw507_[int(23)] = p1
                nw507_[int(24)] = p1
                nw507_[int(25)] = p1
                nw507_[int(26)] = p1
                nw507_[int(27)] = p1
                d_2963_v1_ = nw507_
                d_2964_v2_: _dafny.Array
                nw508_ = _dafny.Array(None, 5)
                nw508_[int(0)] = d_2963_v1_
                nw508_[int(1)] = d_2963_v1_
                nw508_[int(2)] = d_2963_v1_
                nw508_[int(3)] = d_2963_v1_
                nw508_[int(4)] = d_2963_v1_
                d_2964_v2_ = nw508_
                d_2965_v3_: _dafny.Array
                d_2965_v3_ = d_2964_v2_
                d_2965_v3_ = (d_2965_v3_ if p1 else d_2965_v3_)
                d_2966_v4_: str
                d_2966_v4_ = _dafny.CodePoint('n')
                d_2966_v4_ = (self.f7)[default__.safeIndex(len(self.f7), len(self.f7))]
                d_2967_v5_: _dafny.Array
                nw509_ = _dafny.Array(_dafny.Map({}), 8)
                d_2967_v5_ = nw509_
                d_2968_v6_: _dafny.MultiSet
                d_2968_v6_ = _dafny.MultiSet([p1, (p1) == (p1), p1, (857) in (d_2962_v0_)])
                rhs436_ = (d_2967_v5_ if p1 else d_2967_v5_)
                rhs437_ = (d_2968_v6_).intersection(_dafny.MultiSet([True]))
                rhs438_ = (0) - ((d_2962_v0_)[default__.safeIndex(default__.fm1(globalState), len(d_2962_v0_))])
                lhs271_ = self
                d_2967_v5_ = rhs436_
                d_2968_v6_ = rhs437_
                lhs271_.f4 = rhs438_
                d_2969_v7_: _dafny.MultiSet
                d_2969_v7_ = _dafny.MultiSet([self.f7, (self.f7).set(default__.safeIndex(p0, len(self.f7)), d_2966_v4_)])
                d_2970_v8_: T2
                nw510_ = C1()
                nw510_.ctor__(p0, d_2961_i0_, d_2961_i0_, p1)
                d_2970_v8_ = nw510_
                d_2971_v9_: _dafny.Map
                d_2971_v9_ = _dafny.Map({d_2970_v8_: self.f7})
                d_2972_v10_: _dafny.Map
                d_2972_v10_ = _dafny.Map({(d_2970_v8_).f16: self.f7})
                d_2973_v11_: _dafny.Array
                nw511_ = _dafny.Array(None, 26)
                nw511_[int(0)] = self.f7
                nw511_[int(1)] = ((d_2971_v9_)[d_2970_v8_] if (d_2970_v8_) in (d_2971_v9_) else self.f7)
                nw511_[int(2)] = default__.fm23((d_2970_v8_).f16, globalState)
                nw511_[int(3)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jmdbnb"))
                nw511_[int(4)] = self.f7
                nw511_[int(5)] = self.f7
                nw511_[int(6)] = self.f7
                nw511_[int(7)] = self.f7
                nw511_[int(8)] = self.f7
                nw511_[int(9)] = self.f7
                nw511_[int(10)] = self.f7
                nw511_[int(11)] = self.f7
                nw511_[int(12)] = default__.fm23(p1, globalState)
                nw511_[int(13)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_2974_i1_ in range(default__.abs(18))])
                nw511_[int(14)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gkid"))
                nw511_[int(15)] = self.f7
                nw511_[int(16)] = self.f7
                nw511_[int(17)] = self.f7
                nw511_[int(18)] = _dafny.SeqWithoutIsStrInference([d_2966_v4_ for d_2975_i2_ in range(default__.abs(-958))])
                nw511_[int(19)] = _dafny.SeqWithoutIsStrInference([d_2966_v4_ for d_2976_i3_ in range(default__.abs(-657))])
                nw511_[int(20)] = self.f7
                nw511_[int(21)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lqxsea"))
                nw511_[int(22)] = self.f7
                nw511_[int(23)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))
                nw511_[int(24)] = self.f7
                nw511_[int(25)] = ((d_2972_v10_)[not(False)] if (not(False)) in (d_2972_v10_) else self.f7)
                d_2973_v11_ = nw511_
                d_2977_v12_: D8
                d_2977_v12_ = D8_DC30((d_2969_v7_).cardinality, d_2973_v11_, default__.fm0(self.f4, len(self.f7), globalState))
                d_2978_v13_: D3
                d_2978_v13_ = D3_DC8((d_2966_v4_ if (d_2977_v12_).cf56 else d_2966_v4_), d_2963_v1_)
                d_2978_v13_ = d_2978_v13_
                d_2979_v14_: _dafny.Array
                nw512_ = _dafny.Array(D24.default()(), 6)
                d_2979_v14_ = nw512_
                d_2980_v15_: D24
                d_2980_v15_ = D24_DC64()
                index467_ = default__.safeIndex(327, (d_2979_v14_).length(0))
                (d_2979_v14_)[index467_] = d_2980_v15_
                index468_ = default__.safeIndex(327, (d_2979_v14_).length(0))
                (d_2979_v14_)[index468_] = d_2980_v15_
            elif True:
                (self).f4 = self.f4
                d_2981_v16_: _dafny.Map
                d_2981_v16_ = _dafny.Map({599: self.f4})
                d_2981_v16_ = (d_2981_v16_).set(self.f4, p0)
                d_2982_v17_: bool
                d_2982_v17_ = False
                d_2982_v17_ = (d_2961_i0_) == ((self.f4) + (len(self.f7)))
                (self).f4 = default__.safeModuloInt(p0, d_2961_i0_)
                d_2983_v18_: C1
                nw513_ = C1()
                nw513_.ctor__(p0, default__.fm1(globalState), p0, True)
                d_2983_v18_ = nw513_
            (self).f4 = self.f4
            (self).f4 = (0) - (self.f4)
            d_2984_v19_: _dafny.Map
            d_2984_v19_ = _dafny.Map({False: self.f4})
            d_2985_v20_: _dafny.Map
            d_2985_v20_ = _dafny.Map({d_2961_i0_: d_2984_v19_})
            d_2986_v21_: D20
            d_2986_v21_ = D20_DC54(p1, p1, (d_2962_v0_)[default__.safeIndex(d_2961_i0_, len(d_2962_v0_))], p1)
            d_2987_v22_: _dafny.Array
            nw514_ = _dafny.Array(None, 18)
            nw514_[int(0)] = p1
            nw514_[int(1)] = p1
            nw514_[int(2)] = p1
            nw514_[int(3)] = p1
            nw514_[int(4)] = p1
            nw514_[int(5)] = default__.fm0(d_2961_i0_, len(((d_2985_v20_)[self.f4] if (self.f4) in (d_2985_v20_) else d_2984_v19_)), globalState)
            nw514_[int(6)] = p1
            nw514_[int(7)] = p1
            nw514_[int(8)] = p1
            nw514_[int(9)] = p1
            nw514_[int(10)] = p1
            nw514_[int(11)] = p1
            nw514_[int(12)] = (d_2986_v21_).cf100
            nw514_[int(13)] = p1
            nw514_[int(14)] = not(False)
            nw514_[int(15)] = p1
            nw514_[int(16)] = p1
            nw514_[int(17)] = p1
            d_2987_v22_ = nw514_
            d_2988_v23_: C6
            nw515_ = C6()
            nw515_.ctor__(p1, _dafny.MultiSet([d_2987_v22_]), len(default__.fm45(p1, default__.fm0(d_2961_i0_, 700, globalState), globalState)))
            d_2988_v23_ = nw515_
        d_2989_v24_: _dafny.Array
        nw516_ = _dafny.Array(D23.default()(), 27)
        d_2989_v24_ = nw516_
        d_2990_v25_: _dafny.Array
        def lambda142_(d_2991_p1_):
            def lambda143_(d_2992_i4_):
                return d_2991_p1_

            return lambda143_

        init83_ = lambda142_(p1)
        nw517_ = _dafny.Array(None, 11)
        for i0_83_ in range(nw517_.length(0)):
            nw517_[i0_83_] = init83_(i0_83_)
        d_2990_v25_ = nw517_
        d_2993_v26_: D23
        d_2993_v26_ = D23_DC61(d_2990_v25_)
        index469_ = default__.safeIndex(356, (d_2989_v24_).length(0))
        (d_2989_v24_)[index469_] = (d_2993_v26_ if not(p1) else d_2993_v26_)
        index470_ = default__.safeIndex(356, (d_2989_v24_).length(0))
        (d_2989_v24_)[index470_] = d_2993_v26_
        d_2994_v27_: bool
        d_2994_v27_ = True
        d_2994_v27_ = p1
        d_2995_v28_: _dafny.Array
        def lambda144_(d_2996_i5_):
            return (d_2996_i5_) + (self.f4)

        init84_ = lambda144_
        nw518_ = _dafny.Array(None, 7)
        for i0_84_ in range(nw518_.length(0)):
            nw518_[i0_84_] = init84_(i0_84_)
        d_2995_v28_ = nw518_
        d_2997_v29_: _dafny.Seq
        d_2997_v29_ = _dafny.SeqWithoutIsStrInference([d_2995_v28_])
        d_2998_v30_: _dafny.Map
        d_2998_v30_ = _dafny.Map({p1: d_2994_v27_})
        d_2999_v31_: _dafny.Array
        nw519_ = _dafny.Array(None, 9)
        nw519_[int(0)] = (d_2997_v29_)[default__.safeIndex(len(d_2998_v30_), len(d_2997_v29_))]
        nw519_[int(1)] = d_2995_v28_
        nw519_[int(2)] = d_2995_v28_
        nw519_[int(3)] = d_2995_v28_
        nw519_[int(4)] = d_2995_v28_
        nw519_[int(5)] = d_2995_v28_
        nw519_[int(6)] = d_2995_v28_
        nw519_[int(7)] = d_2995_v28_
        nw519_[int(8)] = d_2995_v28_
        d_2999_v31_ = nw519_
        index471_ = default__.safeIndex(423, (d_2999_v31_).length(0))
        (d_2999_v31_)[index471_] = d_2995_v28_
        index472_ = default__.safeIndex(423, (d_2999_v31_).length(0))
        (d_2999_v31_)[index472_] = d_2995_v28_
        d_3000_v32_: _dafny.Seq
        d_3000_v32_ = _dafny.SeqWithoutIsStrInference([self.f7])
        hi18_ = (p0) + (self.f4)
        for d_3001_i6_ in range(default__.safeDivisionInt(118, len(d_3000_v32_)), hi18_):
            (self).f4 = default__.safeModuloInt(d_3001_i6_, d_3001_i6_)
            if (p1 if d_2994_v27_ else d_2994_v27_):
                (self).f4 = d_3001_i6_
                rhs439_ = d_2998_v30_
                rhs440_ = False
                d_2998_v30_ = rhs439_
                d_2994_v27_ = rhs440_
                d_2994_v27_ = d_2994_v27_
                d_3002_v33_: _dafny.Map
                d_3002_v33_ = _dafny.Map({(d_2999_v31_)[default__.safeIndex(423, (d_2999_v31_).length(0))]: p0})
                d_3003_v34_: _dafny.Map
                d_3003_v34_ = _dafny.Map({p1: self.f4})
                d_3004_v35_: _dafny.Seq
                d_3004_v35_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_2995_v28_: p0}), d_3002_v33_])
                d_3005_v36_: _dafny.Array
                nw520_ = _dafny.Array(None, 28)
                nw520_[int(0)] = d_3002_v33_
                nw520_[int(1)] = (d_3002_v33_) | (_dafny.Map({(d_2999_v31_)[default__.safeIndex(423, (d_2999_v31_).length(0))]: len(d_3003_v34_)}))
                nw520_[int(2)] = d_3002_v33_
                nw520_[int(3)] = d_3002_v33_
                nw520_[int(4)] = _dafny.Map({d_2995_v28_: p0})
                nw520_[int(5)] = d_3002_v33_
                nw520_[int(6)] = d_3002_v33_
                nw520_[int(7)] = d_3002_v33_
                nw520_[int(8)] = (d_3002_v33_) | (d_3002_v33_)
                nw520_[int(9)] = d_3002_v33_
                nw520_[int(10)] = _dafny.Map({d_2995_v28_: d_3001_i6_})
                nw520_[int(11)] = (_dafny.Map({d_2995_v28_: p0})) | (d_3002_v33_)
                nw520_[int(12)] = d_3002_v33_
                nw520_[int(13)] = d_3002_v33_
                nw520_[int(14)] = (d_3002_v33_) | ((d_3004_v35_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([p0 for d_3006_i7_ in range(default__.abs(243))])), len(d_3004_v35_))])
                nw520_[int(15)] = d_3002_v33_
                nw520_[int(16)] = d_3002_v33_
                nw520_[int(17)] = d_3002_v33_
                nw520_[int(18)] = d_3002_v33_
                nw520_[int(19)] = (d_3002_v33_) | (d_3002_v33_)
                nw520_[int(20)] = d_3002_v33_
                nw520_[int(21)] = _dafny.Map({(d_2999_v31_)[default__.safeIndex(423, (d_2999_v31_).length(0))]: -726})
                nw520_[int(22)] = d_3002_v33_
                nw520_[int(23)] = (_dafny.Map({d_2995_v28_: d_3001_i6_})) | (d_3002_v33_)
                nw520_[int(24)] = _dafny.Map({(d_2999_v31_)[default__.safeIndex(423, (d_2999_v31_).length(0))]: p0})
                nw520_[int(25)] = d_3002_v33_
                nw520_[int(26)] = d_3002_v33_
                nw520_[int(27)] = d_3002_v33_
                d_3005_v36_ = nw520_
                index473_ = default__.safeIndex(183, (d_3005_v36_).length(0))
                (d_3005_v36_)[index473_] = d_3002_v33_
                index474_ = default__.safeIndex(594, (d_2990_v25_).length(0))
                (d_2990_v25_)[index474_] = d_2994_v27_
                d_3007_v37_: _dafny.Map
                d_3007_v37_ = _dafny.Map({p0: d_2994_v27_})
                d_3008_v38_: _dafny.Set
                d_3008_v38_ = _dafny.Set({d_2994_v27_, p1, not(((d_3007_v37_)[-904] if (-904) in (d_3007_v37_) else p1)), p1, p1})
                d_3009_v39_: _dafny.Seq
                d_3009_v39_ = _dafny.SeqWithoutIsStrInference([(len(d_3008_v38_)) > (len(self.f7)), (True) in (_dafny.Map({True: p0}))])
                d_3010_v40_: _dafny.Seq
                d_3010_v40_ = _dafny.SeqWithoutIsStrInference([self.f4, (_dafny.MultiSet([self.f4, d_3001_i6_])).cardinality, p0, len(self.f7), d_3001_i6_])
                d_3011_v41_: _dafny.Map
                d_3011_v41_ = _dafny.Map({self.f4: d_3010_v40_})
                d_3012_v42_: _dafny.Map
                d_3012_v42_ = _dafny.Map({len(((d_3011_v41_)[default__.fm1(globalState)] if (default__.fm1(globalState)) in (d_3011_v41_) else d_3010_v40_)): d_3002_v33_})
                index475_ = default__.safeIndex(183, (d_3005_v36_).length(0))
                index476_ = default__.safeIndex(594, (d_2990_v25_).length(0))
                rhs441_ = (d_3009_v39_)[default__.safeIndex(self.f4, len(d_3009_v39_))]
                rhs442_ = p0
                rhs443_ = (((d_3012_v42_)[(d_3010_v40_)[default__.safeIndex(default__.fm1(globalState), len(d_3010_v40_))]] if ((d_3010_v40_)[default__.safeIndex(default__.fm1(globalState), len(d_3010_v40_))]) in (d_3012_v42_) else d_3002_v33_)) | (_dafny.Map({(d_2999_v31_)[default__.safeIndex(423, (d_2999_v31_).length(0))]: -126}))
                rhs444_ = d_2994_v27_
                lhs272_ = self
                lhs273_ = d_3005_v36_
                lhs274_ = default__.safeIndex(183, (d_3005_v36_).length(0))
                lhs275_ = d_2990_v25_
                lhs276_ = default__.safeIndex(594, (d_2990_v25_).length(0))
                d_2994_v27_ = rhs441_
                lhs272_.f4 = rhs442_
                lhs273_[lhs274_] = rhs443_
                lhs275_[lhs276_] = rhs444_
                d_3013_v43_: _dafny.Array
                nw521_ = _dafny.Array(_dafny.Map({}), 10)
                d_3013_v43_ = nw521_
                d_3014_v44_: D23
                d_3014_v44_ = D23_DC60(d_3013_v43_)
                d_3015_v45_: D23
                d_3015_v45_ = D23_DC62(d_3014_v44_)
                d_3016_v46_: _dafny.Map
                d_3016_v46_ = _dafny.Map({d_3015_v45_: self.f4})
                d_3017_v47_: C5
                nw522_ = C5()
                nw522_.ctor__(((d_3016_v46_)[d_3015_v45_] if (d_3015_v45_) in (d_3016_v46_) else p0))
                d_3017_v47_ = nw522_
            elif True:
                d_3018_v48_: _dafny.Seq
                d_3018_v48_ = _dafny.SeqWithoutIsStrInference([d_2994_v27_])
                d_3019_v49_: _dafny.MultiSet
                d_3019_v49_ = _dafny.MultiSet([(d_3018_v48_)[default__.safeIndex(p0, len(d_3018_v48_))]])
                d_3020_v50_: C0
                nw523_ = C0()
                nw523_.ctor__(((d_3019_v49_)[d_2994_v27_] if (d_2994_v27_) in (d_3019_v49_) else default__.fm1(globalState)))
                d_3020_v50_ = nw523_
                index477_ = default__.safeIndex(423, (d_2999_v31_).length(0))
                (d_2999_v31_)[index477_] = d_2995_v28_
                d_3021_v51_: _dafny.Map
                d_3021_v51_ = _dafny.Map({self.f4: not(not(d_2994_v27_))})
                d_3022_v52_: C2
                nw524_ = C2()
                nw524_.ctor__((len(d_3021_v51_)) >= (946), d_3020_v50_.f18)
                d_3022_v52_ = nw524_
                d_3023_v53_: str
                d_3023_v53_ = _dafny.CodePoint('f')
                index478_ = default__.safeIndex(346, (d_2990_v25_).length(0))
                (d_2990_v25_)[index478_] = d_3022_v52_.f13
                index479_ = default__.safeIndex(346, (d_2990_v25_).length(0))
                rhs445_ = d_2994_v27_
                rhs446_ = d_3023_v53_
                rhs447_ = (_dafny.CodePoint('m') if ((d_3021_v51_)[d_3020_v50_.f18] if (d_3020_v50_.f18) in (d_3021_v51_) else d_2994_v27_) else d_3023_v53_)
                rhs448_ = d_3022_v52_.f13
                rhs449_ = ((d_3022_v52_).f14) >= (self.f4)
                lhs277_ = d_3022_v52_
                lhs278_ = d_3022_v52_
                lhs279_ = d_2990_v25_
                lhs280_ = default__.safeIndex(346, (d_2990_v25_).length(0))
                lhs277_.f13 = rhs445_
                d_3023_v53_ = rhs446_
                d_3023_v53_ = rhs447_
                lhs278_.f13 = rhs448_
                lhs279_[lhs280_] = rhs449_
                d_3024_v54_: _dafny.Map
                d_3024_v54_ = _dafny.Map({d_3022_v52_.f13: 740})
                d_3025_v55_: _dafny.Seq
                d_3025_v55_ = _dafny.SeqWithoutIsStrInference([d_3024_v54_, _dafny.Map({True: p0}), d_3024_v54_])
                d_3026_v56_: _dafny.Seq
                d_3026_v56_ = _dafny.SeqWithoutIsStrInference([((d_3019_v49_)[(d_2990_v25_)[default__.safeIndex(346, (d_2990_v25_).length(0))]] if ((d_2990_v25_)[default__.safeIndex(346, (d_2990_v25_).length(0))]) in (d_3019_v49_) else self.f4), self.f4])
                d_2994_v27_ = (p1) not in ((d_3025_v55_)[default__.safeIndex(len(d_3026_v56_), len(d_3025_v55_))])
            d_3027_v57_: _dafny.Seq
            d_3027_v57_ = _dafny.SeqWithoutIsStrInference([d_2994_v27_])
            d_3028_v58_: _dafny.Map
            d_3028_v58_ = _dafny.Map({p1: self.f4})
            d_3029_v59_: _dafny.Map
            d_3029_v59_ = _dafny.Map({d_2990_v25_: ((d_3028_v58_)[p1] if (p1) in (d_3028_v58_) else self.f4)})
            d_3030_v60_: C0
            nw525_ = C0()
            nw525_.ctor__((0) - (len(d_3029_v59_)))
            d_3030_v60_ = nw525_
            d_3031_v61_: _dafny.Map
            d_3031_v61_ = _dafny.Map({(d_3027_v57_)[default__.safeIndex(self.f4, len(d_3027_v57_))]: d_3030_v60_})
            d_3032_v62_: D3
            d_3032_v62_ = D3_DC9(False, len(d_3031_v61_))
            d_3033_v63_: D3
            d_3033_v63_ = D3_DC11(d_3032_v62_)
            d_3034_v64_: D3
            d_3034_v64_ = D3_DC11((d_3033_v63_).cf22)
            source32_ = (d_3033_v63_ if True else D3_DC11(d_3032_v62_))
            if source32_.is_DC8:
                d_3035___mcc_h0_ = source32_.cf13
                d_3036___mcc_h1_ = source32_.cf14
                d_3037_cf14_ = d_3036___mcc_h1_
                d_3038_cf13_ = d_3035___mcc_h0_
                d_2994_v27_ = p1
                (d_3030_v60_).f18 = 884
                d_3039_v65_: C5
                nw526_ = C5()
                nw526_.ctor__((0) - (len((_dafny.SeqWithoutIsStrInference([d_3038_cf13_ for d_3040_i8_ in range(default__.abs(899))])) + (self.f7))))
                d_3039_v65_ = nw526_
                d_3041_v66_: _dafny.MultiSet
                d_3041_v66_ = _dafny.MultiSet([False, (d_3038_cf13_) in (self.f7), d_2994_v27_])
                d_3042_v67_: _dafny.Seq
                d_3042_v67_ = _dafny.SeqWithoutIsStrInference([self.f4, self.f4, default__.safeModuloInt(562, d_3030_v60_.f18)])
                d_3043_v68_: _dafny.Seq
                d_3043_v68_ = _dafny.SeqWithoutIsStrInference([d_3027_v57_, _dafny.SeqWithoutIsStrInference([d_2994_v27_, d_2994_v27_, p1, p1, p1]), d_3027_v57_, d_3027_v57_, d_3027_v57_])
                d_3044_v69_: _dafny.Array
                nw527_ = _dafny.Array(None, 10)
                nw527_[int(0)] = (d_3043_v68_)[default__.safeIndex(p0, len(d_3043_v68_))]
                nw527_[int(1)] = d_3027_v57_
                nw527_[int(2)] = d_3027_v57_
                nw527_[int(3)] = (d_3027_v57_) + (_dafny.SeqWithoutIsStrInference([p1]))
                nw527_[int(4)] = (_dafny.SeqWithoutIsStrInference([not(d_2994_v27_)])).set(default__.safeIndex(len(self.f7), len(_dafny.SeqWithoutIsStrInference([not(d_2994_v27_)]))), False)
                nw527_[int(5)] = (d_3043_v68_)[default__.safeIndex(d_3001_i6_, len(d_3043_v68_))]
                nw527_[int(6)] = d_3027_v57_
                nw527_[int(7)] = (d_3043_v68_)[default__.safeIndex(177, len(d_3043_v68_))]
                nw527_[int(8)] = d_3027_v57_
                nw527_[int(9)] = d_3027_v57_
                d_3044_v69_ = nw527_
                index480_ = default__.safeIndex(268, (d_3044_v69_).length(0))
                (d_3044_v69_)[index480_] = d_3027_v57_
                index481_ = default__.safeIndex(268, (d_3044_v69_).length(0))
                rhs450_ = (d_3041_v66_) | (d_3041_v66_)
                rhs451_ = _dafny.SeqWithoutIsStrInference([885, self.f4])
                rhs452_ = (d_3027_v57_) + (_dafny.SeqWithoutIsStrInference([p1]))
                rhs453_ = d_3001_i6_
                rhs454_ = not (d_2994_v27_) or (d_2994_v27_)
                lhs281_ = d_3044_v69_
                lhs282_ = default__.safeIndex(268, (d_3044_v69_).length(0))
                lhs283_ = self
                d_3041_v66_ = rhs450_
                d_3042_v67_ = rhs451_
                lhs281_[lhs282_] = rhs452_
                lhs283_.f4 = rhs453_
                d_2994_v27_ = rhs454_
            elif source32_.is_DC9:
                d_3045___mcc_h2_ = source32_.cf15
                d_3046___mcc_h3_ = source32_.cf16
                d_3047_cf16_ = d_3046___mcc_h3_
                d_3048_cf15_ = d_3045___mcc_h2_
                d_2994_v27_ = (d_3048_cf15_) == ((self).fm5(self.f7, d_2994_v27_, globalState))
                r1 = (((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_3049_i9_ in range(default__.abs(-14))])).set(default__.safeIndex(d_3001_i6_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_3050_i9_ in range(default__.abs(-14))]))), _dafny.CodePoint('k'))) + (self.f7)) + (self.f7)
                index482_ = default__.safeIndex(414, (d_2995_v28_).length(0))
                (d_2995_v28_)[index482_] = (_dafny.MultiSet([617])).cardinality
                index483_ = default__.safeIndex(414, (d_2995_v28_).length(0))
                (d_2995_v28_)[index483_] = p0
                index484_ = default__.safeIndex(828, (d_2990_v25_).length(0))
                (d_2990_v25_)[index484_] = d_3048_cf15_
                d_3051_v70_: _dafny.Map
                d_3051_v70_ = _dafny.Map({623: ((d_2995_v28_)[default__.safeIndex(414, (d_2995_v28_).length(0))]) <= (len(d_3028_v58_))})
                d_3052_v71_: D6
                d_3052_v71_ = D6_DC22()
                d_3053_v72_: D6
                d_3053_v72_ = D6_DC23(d_3052_v71_)
                d_3054_v73_: _dafny.MultiSet
                d_3054_v73_ = _dafny.MultiSet([d_3053_v72_])
                d_3055_v75_: _dafny.MultiSet
                d_3055_v75_ = _dafny.MultiSet([d_3047_cf16_])
                d_3056_v76_: _dafny.MultiSet
                d_3056_v76_ = _dafny.MultiSet([(d_3055_v75_).cardinality, self.f4])
                d_3057_v77_: _dafny.Map
                d_3057_v77_ = _dafny.Map({d_2994_v27_: d_2990_v25_})
                index485_ = default__.safeIndex(414, (d_2995_v28_).length(0))
                index486_ = default__.safeIndex(828, (d_2990_v25_).length(0))
                def iife281_():
                    coll103_ = _dafny.Map()
                    compr_103_: int
                    for compr_103_ in (d_3055_v75_).Elements:
                        d_3058_v74_: int = compr_103_
                        if (d_3058_v74_) in (d_3055_v75_):
                            coll103_[(d_3058_v74_) + ((d_2995_v28_)[default__.safeIndex(414, (d_2995_v28_).length(0))])] = d_3048_cf15_
                    return _dafny.Map(coll103_)
                rhs455_ = p0
                rhs456_ = (d_3054_v73_).ispropersubset(_dafny.MultiSet([d_3053_v72_, d_3053_v72_]))
                rhs457_ = False
                rhs458_ = (((iife281_()
                ).set(len(d_3057_v77_), d_2994_v27_)) | (d_3051_v70_)).set(self.f4, d_3048_cf15_)
                rhs459_ = (d_3027_v57_)[default__.safeIndex(d_3030_v60_.f18, len(d_3027_v57_))]
                lhs284_ = d_2995_v28_
                lhs285_ = default__.safeIndex(414, (d_2995_v28_).length(0))
                lhs286_ = d_2990_v25_
                lhs287_ = default__.safeIndex(828, (d_2990_v25_).length(0))
                lhs284_[lhs285_] = rhs455_
                d_2994_v27_ = rhs456_
                lhs286_[lhs287_] = rhs457_
                d_3051_v70_ = rhs458_
                d_2994_v27_ = rhs459_
            elif source32_.is_DC10:
                d_3059___mcc_h4_ = source32_.cf17
                d_3060___mcc_h5_ = source32_.cf18
                d_3061___mcc_h6_ = source32_.cf19
                d_3062___mcc_h7_ = source32_.cf20
                d_3063___mcc_h8_ = source32_.cf21
                d_3064_cf21_ = d_3063___mcc_h8_
                d_3065_cf20_ = d_3062___mcc_h7_
                d_3066_cf19_ = d_3061___mcc_h6_
                d_3067_cf18_ = d_3060___mcc_h5_
                d_3068_cf17_ = d_3059___mcc_h4_
                d_3066_cf19_ = True
                (self).f7 = (self.f7) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rlwnanv"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "swltrq"))))
                d_3069_v78_: C3
                nw528_ = C3()
                nw528_.ctor__(d_3068_cf17_, self.f7, d_3068_cf17_)
                d_3069_v78_ = nw528_
                d_3065_cf20_ = d_3065_cf20_
            elif source32_.is_DC7:
                d_3070___mcc_h9_ = source32_.cf12
                d_3071_cf12_ = d_3070___mcc_h9_
                d_3072_v79_: D4
                d_3072_v79_ = D4_DC12(self.f7)
                d_3073_v80_: _dafny.Array
                nw529_ = _dafny.Array(None, 6)
                nw529_[int(0)] = (self.f7) + (self.f7)
                nw529_[int(1)] = self.f7
                nw529_[int(2)] = self.f7
                nw529_[int(3)] = (d_3072_v79_).cf23
                nw529_[int(4)] = (d_3000_v32_)[default__.safeIndex(196, len(d_3000_v32_))]
                nw529_[int(5)] = self.f7
                d_3073_v80_ = nw529_
                index487_ = default__.safeIndex(733, (d_3073_v80_).length(0))
                (d_3073_v80_)[index487_] = self.f7
                index488_ = default__.safeIndex(733, (d_3073_v80_).length(0))
                (d_3073_v80_)[index488_] = self.f7
                d_3074_v81_: D11
                d_3074_v81_ = D11_DC35(d_3001_i6_, d_2994_v27_, p1)
                (self).f4 = (0) - ((d_3074_v81_).cf61)
                d_3075_v82_: C1
                nw530_ = C1()
                nw530_.ctor__(self.f4, self.f4, default__.safeDivisionInt(d_3030_v60_.f18, p0), not(p1))
                d_3075_v82_ = nw530_
                d_3076_v83_: _dafny.Array
                def lambda145_(d_3077_v57_):
                    def lambda146_(d_3078_i10_):
                        return d_3077_v57_

                    return lambda146_

                init85_ = lambda145_(d_3027_v57_)
                nw531_ = _dafny.Array(None, 13)
                for i0_85_ in range(nw531_.length(0)):
                    nw531_[i0_85_] = init85_(i0_85_)
                d_3076_v83_ = nw531_
                d_3079_v84_: D18
                d_3079_v84_ = D18_DC47(d_3076_v83_)
                d_3080_v85_: _dafny.Seq
                d_3080_v85_ = _dafny.SeqWithoutIsStrInference([d_3079_v84_])
                d_3081_v86_: _dafny.Seq
                d_3081_v86_ = _dafny.SeqWithoutIsStrInference([d_3080_v85_])
                index489_ = default__.safeIndex(481, (d_2995_v28_).length(0))
                (d_2995_v28_)[index489_] = len((d_3081_v86_) + (d_3081_v86_))
                d_3082_v87_: _dafny.Map
                d_3082_v87_ = _dafny.Map({p1: d_3076_v83_})
                d_3083_v88_: _dafny.Map
                d_3083_v88_ = _dafny.Map({d_3030_v60_.f18: ((d_3082_v87_)[p1] if (p1) in (d_3082_v87_) else d_3076_v83_)})
                d_3084_v89_: D18
                d_3084_v89_ = D18_DC47(((d_3083_v88_)[d_3030_v60_.f18] if (d_3030_v60_.f18) in (d_3083_v88_) else d_3076_v83_))
                d_3085_v90_: D18
                d_3085_v90_ = D18_DC49(D18_DC49(d_3084_v89_))
                d_3086_v91_: _dafny.Array
                nw532_ = _dafny.Array(_dafny.Map({}), 4)
                d_3086_v91_ = nw532_
                index490_ = default__.safeIndex(249, (d_3086_v91_).length(0))
                (d_3086_v91_)[index490_] = _dafny.Map({p1: self.f4})
                index491_ = default__.safeIndex(481, (d_2995_v28_).length(0))
                index492_ = default__.safeIndex(249, (d_3086_v91_).length(0))
                rhs460_ = (0) - (d_3075_v82_.f17)
                rhs461_ = p1
                rhs462_ = D18_DC49((D18_DC49(d_3084_v89_)).cf88)
                rhs463_ = d_3028_v58_
                rhs464_ = (d_3030_v60_.f18 if d_2994_v27_ else p0)
                lhs288_ = d_2995_v28_
                lhs289_ = default__.safeIndex(481, (d_2995_v28_).length(0))
                lhs290_ = d_3086_v91_
                lhs291_ = default__.safeIndex(249, (d_3086_v91_).length(0))
                lhs292_ = d_3030_v60_
                lhs288_[lhs289_] = rhs460_
                d_2994_v27_ = rhs461_
                d_3085_v90_ = rhs462_
                lhs290_[lhs291_] = rhs463_
                lhs292_.f18 = rhs464_
            elif True:
                d_3087___mcc_h10_ = source32_.cf22
                d_3088_cf22_ = d_3087___mcc_h10_
                d_2994_v27_ = p1
                d_3089_v92_: _dafny.Map
                d_3089_v92_ = _dafny.Map({default__.fm1(globalState): p1})
                d_3090_v93_: _dafny.Seq
                d_3090_v93_ = _dafny.SeqWithoutIsStrInference([len(self.f7), d_3030_v60_.f18, self.f4, d_3001_i6_])
                d_3091_v94_: _dafny.Array
                d_3091_v94_ = (d_2999_v31_)[default__.safeIndex(423, (d_2999_v31_).length(0))]
                d_3092_v95_: int
                d_3093_v96_: int
                d_3094_v97_: _dafny.Array
                d_3095_v98_: bool
                out63_: int
                out64_: int
                out65_: _dafny.Array
                out66_: bool
                out63_, out64_, out65_, out66_ = default__.m0(d_3089_v92_, d_3090_v93_, (_dafny.Map({378: d_2995_v28_})) | (_dafny.Map({d_3030_v60_.f18: (d_3091_v94_)})), globalState)
                d_3092_v95_ = out63_
                d_3093_v96_ = out64_
                d_3094_v97_ = out65_
                d_3095_v98_ = out66_
                d_3096_v99_: _dafny.MultiSet
                d_3096_v99_ = _dafny.MultiSet([d_3090_v93_])
                d_2994_v27_ = (d_3096_v99_).issubset(d_3096_v99_)
                d_3097_v100_: C1
                nw533_ = C1()
                nw533_.ctor__((0) - (d_3030_v60_.f18), (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_3098_i11_ in range(default__.abs(828))]))) - (d_3030_v60_.f18), default__.safeDivisionInt(d_3030_v60_.f18, 949), d_3095_v98_)
                d_3097_v100_ = nw533_
            (d_3030_v60_).f18 = d_3001_i6_
        (self).f4 = p0
        d_3099_v101_: _dafny.Seq
        d_3099_v101_ = _dafny.SeqWithoutIsStrInference([d_2994_v27_])
        d_3100_v102_: _dafny.Seq
        d_3100_v102_ = _dafny.SeqWithoutIsStrInference([d_2990_v25_, d_2990_v25_, d_2990_v25_, d_2990_v25_])
        d_3101_v103_: _dafny.MultiSet
        d_3101_v103_ = _dafny.MultiSet([p0, p0, self.f4])
        d_3102_v104_: _dafny.Seq
        d_3102_v104_ = _dafny.SeqWithoutIsStrInference([p0])
        d_3103_v105_: _dafny.Map
        d_3103_v105_ = _dafny.Map({d_2994_v27_: d_3102_v104_})
        d_3104_v106_: D0
        d_3104_v106_ = D0_DC0(((d_3103_v105_)[p1] if (p1) in (d_3103_v105_) else d_3102_v104_))
        pat_let_tv68_ = d_3102_v104_
        def iife282_(_pat_let89_0):
            def iife283_(d_3105_dt__update__tmp_h0_):
                def iife284_(_pat_let90_0):
                    def iife285_(d_3106_dt__update_hcf0_h0_):
                        return D0_DC0(d_3106_dt__update_hcf0_h0_)
                    return iife285_(_pat_let90_0)
                return iife284_(pat_let_tv68_)
            return iife283_(_pat_let89_0)
        r0 = ((default__.fm51((_dafny.MultiSet([len(d_3099_v101_), self.f4, len(d_3100_v102_)])).isdisjoint(d_3101_v103_), d_2994_v27_, globalState)).set(default__.safeIndex((p0 if d_2994_v27_ else -987), len(default__.fm51((_dafny.MultiSet([len(d_3099_v101_), self.f4, len(d_3100_v102_)])).isdisjoint(d_3101_v103_), d_2994_v27_, globalState))), d_3104_v106_)).set(default__.safeIndex(p0, len((default__.fm51((_dafny.MultiSet([len(d_3099_v101_), self.f4, len(d_3100_v102_)])).isdisjoint(d_3101_v103_), d_2994_v27_, globalState)).set(default__.safeIndex((p0 if d_2994_v27_ else -987), len(default__.fm51((_dafny.MultiSet([len(d_3099_v101_), self.f4, len(d_3100_v102_)])).isdisjoint(d_3101_v103_), d_2994_v27_, globalState))), d_3104_v106_))), iife282_(d_3104_v106_))
        r1 = self.f7
        r2 = (d_2999_v31_)[default__.safeIndex(423, (d_2999_v31_).length(0))]
        return r0, r1, r2


class C8:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C8"
    def ctor__(self):
        pass
        pass

    def fm6(self, p0, p1, p2, p3, globalState):
        return ((D0_DC1(-224, 831, not(False))).cf2) * ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True, False, not(False)]))).cardinality)

    def m5(self, p0, globalState):
        r0: int = int(0)
        r1: bool = False
        d_3107_v0_: _dafny.Seq
        d_3107_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "shro"))
        d_3108_v1_: str
        d_3108_v1_ = _dafny.CodePoint('l')
        d_3109_v2_: _dafny.Set
        d_3109_v2_ = _dafny.Set({d_3108_v1_})
        d_3110_v3_: _dafny.Map
        d_3110_v3_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hci")): d_3109_v2_})
        d_3111_v4_: bool
        d_3111_v4_ = False
        d_3112_v5_: int
        d_3112_v5_ = -541
        d_3113_v6_: _dafny.Set
        d_3113_v6_ = _dafny.Set({False, d_3111_v4_})
        d_3114_v7_: _dafny.Array
        nw534_ = _dafny.Array(None, 12)
        nw534_[int(0)] = (715) in (p0)
        nw534_[int(1)] = (len(d_3107_v0_)) in (p0)
        nw534_[int(2)] = (d_3109_v2_).ispropersubset(((d_3110_v3_)[_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ucvhjfkh"))] if (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ucvhjfkh"))) in (d_3110_v3_) else d_3109_v2_))
        nw534_[int(3)] = d_3111_v4_
        nw534_[int(4)] = d_3111_v4_
        nw534_[int(5)] = not (d_3111_v4_) or (d_3111_v4_)
        nw534_[int(6)] = (d_3112_v5_) < (default__.fm1(globalState))
        nw534_[int(7)] = not (d_3111_v4_) or (not(d_3111_v4_))
        nw534_[int(8)] = d_3111_v4_
        nw534_[int(9)] = not (d_3111_v4_) or (d_3111_v4_)
        nw534_[int(10)] = (d_3113_v6_).isdisjoint(_dafny.Set({d_3111_v4_}))
        nw534_[int(11)] = True
        d_3114_v7_ = nw534_
        d_3114_v7_ = d_3114_v7_
        d_3111_v4_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rdyiu"))) <= (d_3107_v0_)
        hi19_ = default__.safeModuloInt(d_3112_v5_, d_3112_v5_)
        for d_3115_i0_ in range(d_3112_v5_, hi19_):
            r1 = d_3111_v4_
            d_3116_v8_: C2
            nw535_ = C2()
            nw535_.ctor__(d_3111_v4_, d_3112_v5_)
            d_3116_v8_ = nw535_
            d_3117_v9_: _dafny.Array
            nw536_ = _dafny.Array(int(0), 19)
            d_3117_v9_ = nw536_
            index493_ = default__.safeIndex(956, (d_3117_v9_).length(0))
            (d_3117_v9_)[index493_] = default__.safeModuloInt(len(p0), 169)
            index494_ = default__.safeIndex(956, (d_3117_v9_).length(0))
            (d_3117_v9_)[index494_] = d_3112_v5_
            d_3118_v10_: C1
            nw537_ = C1()
            nw537_.ctor__((d_3117_v9_)[default__.safeIndex(956, (d_3117_v9_).length(0))], d_3115_i0_, (d_3116_v8_).f14, d_3111_v4_)
            d_3118_v10_ = nw537_
            d_3119_v11_: _dafny.Array
            nw538_ = _dafny.Array(None, 7)
            nw538_[int(0)] = d_3118_v10_
            nw538_[int(1)] = d_3118_v10_
            nw538_[int(2)] = d_3118_v10_
            nw538_[int(3)] = d_3118_v10_
            nw538_[int(4)] = d_3118_v10_
            nw538_[int(5)] = d_3118_v10_
            nw538_[int(6)] = d_3118_v10_
            d_3119_v11_ = nw538_
            d_3120_v12_: _dafny.Map
            d_3120_v12_ = _dafny.Map({d_3116_v8_.f13: d_3119_v11_})
            d_3120_v12_ = d_3120_v12_
        index495_ = default__.safeIndex(29, (d_3114_v7_).length(0))
        (d_3114_v7_)[index495_] = d_3111_v4_
        index496_ = default__.safeIndex(29, (d_3114_v7_).length(0))
        rhs465_ = default__.fm23(d_3111_v4_, globalState)
        rhs466_ = d_3107_v0_
        rhs467_ = (d_3107_v0_) + ((d_3107_v0_) + (d_3107_v0_))
        rhs468_ = d_3111_v4_
        lhs293_ = d_3114_v7_
        lhs294_ = default__.safeIndex(29, (d_3114_v7_).length(0))
        d_3107_v0_ = rhs465_
        d_3107_v0_ = rhs466_
        d_3107_v0_ = rhs467_
        lhs293_[lhs294_] = rhs468_
        hi20_ = d_3112_v5_
        for d_3121_i1_ in range((d_3112_v5_) * (d_3112_v5_), hi20_):
            d_3122_v13_: C2
            nw539_ = C2()
            nw539_.ctor__(not((d_3114_v7_)[default__.safeIndex(29, (d_3114_v7_).length(0))]), (0) - (d_3112_v5_))
            d_3122_v13_ = nw539_
            r0 = (d_3122_v13_).f14
            r0 = (0) - ((d_3121_i1_) + ((d_3122_v13_).f14))
            d_3123_v14_: _dafny.Seq
            d_3123_v14_ = _dafny.SeqWithoutIsStrInference([default__.fm0(len(d_3113_v6_), len(_dafny.Map({(d_3114_v7_)[default__.safeIndex(29, (d_3114_v7_).length(0))]: (d_3122_v13_).f14})), globalState), d_3122_v13_.f13])
            if not((_dafny.SeqWithoutIsStrInference([d_3111_v4_, (d_3114_v7_)[default__.safeIndex(29, (d_3114_v7_).length(0))]])) != (d_3123_v14_)):
                d_3124_v15_: C4
                nw540_ = C4()
                nw540_.ctor__(d_3112_v5_)
                d_3124_v15_ = nw540_
                d_3125_v16_: _dafny.Seq
                d_3125_v16_ = _dafny.SeqWithoutIsStrInference([d_3124_v15_])
                d_3126_v17_: _dafny.MultiSet
                d_3126_v17_ = _dafny.MultiSet([970, len((d_3125_v16_ if d_3122_v13_.f13 else d_3125_v16_)), (d_3122_v13_).f14, d_3121_i1_, d_3121_i1_])
                d_3112_v5_ = (d_3126_v17_).cardinality
                d_3127_v18_: _dafny.Map
                d_3127_v18_ = _dafny.Map({d_3121_i1_: p0})
                d_3128_v19_: _dafny.Map
                d_3128_v19_ = _dafny.Map({d_3111_v4_: _dafny.Map({(d_3122_v13_).f14: d_3111_v4_})})
                (d_3122_v13_).f13 = not (default__.fm0(len(((d_3127_v18_)[918] if (918) in (d_3127_v18_) else p0)), len(((d_3128_v19_)[d_3122_v13_.f13] if (d_3122_v13_.f13) in (d_3128_v19_) else default__.fm37(d_3122_v13_.f13, d_3122_v13_.f13, (d_3122_v13_).f14, globalState))), globalState)) or (not(d_3111_v4_))
                d_3112_v5_ = (d_3122_v13_).f14
                d_3129_v20_: D3
                d_3129_v20_ = D3_DC10((d_3122_v13_).f14, 846, d_3122_v13_.f13, (0) - (d_3112_v5_), d_3111_v4_)
                d_3130_v21_: _dafny.Map
                d_3130_v21_ = _dafny.Map({d_3122_v13_.f13: d_3107_v0_})
                pat_let_tv69_ = d_3112_v5_
                d_3131_v22_: _dafny.Array
                nw541_ = _dafny.Array(None, 5)
                nw541_[int(0)] = d_3129_v20_
                nw541_[int(1)] = D3_DC10((d_3122_v13_).f14, len(d_3107_v0_), d_3122_v13_.f13, d_3121_i1_, d_3122_v13_.f13)
                def iife286_(_pat_let91_0):
                    def iife287_(d_3132_dt__update__tmp_h0_):
                        def iife288_(_pat_let92_0):
                            def iife289_(d_3133_dt__update_hcf20_h0_):
                                return D3_DC10((d_3132_dt__update__tmp_h0_).cf17, (d_3132_dt__update__tmp_h0_).cf18, (d_3132_dt__update__tmp_h0_).cf19, d_3133_dt__update_hcf20_h0_, (d_3132_dt__update__tmp_h0_).cf21)
                            return iife289_(_pat_let92_0)
                        return iife288_(pat_let_tv69_)
                    return iife287_(_pat_let91_0)
                nw541_[int(2)] = iife286_(D3_DC10(d_3112_v5_, (d_3122_v13_).f14, not(d_3111_v4_), len(d_3130_v21_), d_3122_v13_.f13))
                nw541_[int(3)] = d_3129_v20_
                nw541_[int(4)] = D3_DC10(len(d_3107_v0_), (d_3122_v13_).f14, False, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ks"))), (d_3124_v15_).fm5(d_3107_v0_, d_3111_v4_, globalState))
                d_3131_v22_ = nw541_
                index497_ = default__.safeIndex(459, (d_3131_v22_).length(0))
                (d_3131_v22_)[index497_] = d_3129_v20_
                d_3134_v23_: _dafny.Map
                d_3134_v23_ = _dafny.Map({len(_dafny.Map({d_3123_v14_: d_3122_v13_.f13})): d_3108_v1_})
                index498_ = default__.safeIndex(459, (d_3131_v22_).length(0))
                (d_3131_v22_)[index498_] = (D3_DC10(len((_dafny.SeqWithoutIsStrInference([d_3108_v1_ for d_3135_i2_ in range(default__.abs(855))])).set(default__.safeIndex(d_3121_i1_, len(_dafny.SeqWithoutIsStrInference([d_3108_v1_ for d_3136_i2_ in range(default__.abs(855))]))), d_3108_v1_)), (d_3122_v13_).f14, (d_3114_v7_)[default__.safeIndex(29, (d_3114_v7_).length(0))], len((d_3134_v23_).set((d_3122_v13_).f14, d_3108_v1_)), (d_3114_v7_)[default__.safeIndex(29, (d_3114_v7_).length(0))]) if not((d_3114_v7_)[default__.safeIndex(29, (d_3114_v7_).length(0))]) else d_3129_v20_)
                d_3137_v24_: C0
                nw542_ = C0()
                nw542_.ctor__(216)
                d_3137_v24_ = nw542_
            elif True:
                d_3109_v2_ = _dafny.Set({d_3108_v1_})
                d_3138_v25_: _dafny.Array
                nw543_ = _dafny.Array(None, 4)
                nw543_[int(0)] = d_3107_v0_
                nw543_[int(1)] = (d_3107_v0_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q")))
                nw543_[int(2)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nwkn"))
                nw543_[int(3)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cwvxgkcp"))
                d_3138_v25_ = nw543_
                index499_ = default__.safeIndex(462, (d_3138_v25_).length(0))
                (d_3138_v25_)[index499_] = d_3107_v0_
                index500_ = default__.safeIndex(462, (d_3138_v25_).length(0))
                (d_3138_v25_)[index500_] = d_3107_v0_
                d_3139_v26_: _dafny.MultiSet
                d_3139_v26_ = _dafny.MultiSet([d_3114_v7_])
                d_3140_v27_: C6
                nw544_ = C6()
                nw544_.ctor__(default__.fm0(d_3121_i1_, (d_3122_v13_).f14, globalState), (d_3139_v26_) | (d_3139_v26_), (0) - (default__.safeModuloInt(d_3121_i1_, d_3112_v5_)))
                d_3140_v27_ = nw544_
                d_3141_v28_: C3
                nw545_ = C3()
                nw545_.ctor__((d_3122_v13_).f14, (_dafny.SeqWithoutIsStrInference([d_3108_v1_, d_3108_v1_, d_3108_v1_])).set(default__.safeIndex(d_3112_v5_, len(_dafny.SeqWithoutIsStrInference([d_3108_v1_, d_3108_v1_, d_3108_v1_]))), d_3108_v1_), len(d_3107_v0_))
                d_3141_v28_ = nw545_
                d_3108_v1_ = d_3108_v1_
        d_3142_v29_: _dafny.Array
        def lambda147_(d_3143_v5_):
            def lambda148_(d_3144_i3_):
                return (d_3144_i3_) * (d_3143_v5_)

            return lambda148_

        init86_ = lambda147_(d_3112_v5_)
        nw546_ = _dafny.Array(None, 16)
        for i0_86_ in range(nw546_.length(0)):
            nw546_[i0_86_] = init86_(i0_86_)
        d_3142_v29_ = nw546_
        index501_ = default__.safeIndex(818, (d_3142_v29_).length(0))
        (d_3142_v29_)[index501_] = d_3112_v5_
        index502_ = default__.safeIndex(818, (d_3142_v29_).length(0))
        (d_3142_v29_)[index502_] = d_3112_v5_
        r0 = (d_3142_v29_)[default__.safeIndex(818, (d_3142_v29_).length(0))]
        r1 = d_3111_v4_
        return r0, r1


class C9(T0):
    def  __init__(self):
        self._f4: int = int(0)
        self.f5: int = int(0)
        self.f6: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C9"
    @property
    def f4(self):
        return self._f4
    @f4.setter
    def f4(self, value):
        self._f4 = value
    def ctor__(self, f5, f6, f4):
        (self).f5 = f5
        (self).f6 = f6
        (self).f4 = f4

    def fm5(self, p0, p1, globalState):
        return self.f6

    def m3(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        d_3145_v0_: _dafny.Seq
        d_3145_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lpxfcmc"))
        d_3146_v1_: _dafny.Array
        nw547_ = _dafny.Array(False, 7)
        d_3146_v1_ = nw547_
        d_3147_v2_: C7
        nw548_ = C7()
        nw548_.ctor__(d_3145_v0_, len(_dafny.Map({d_3146_v1_: len((default__.fm28(self.f4, globalState)).set(p0, p1))})))
        d_3147_v2_ = nw548_
        d_3148_i0_: int
        d_3148_i0_ = 0
        with _dafny.label("11"):
            while p0:
                with _dafny.c_label("11"):
                    if (d_3148_i0_) >= (100):
                        raise _dafny.Break("11")
                    d_3148_i0_ = (d_3148_i0_) + (1)
                    (self).f6 = (p3) == ((self.f5) <= (p1))
                    d_3149_v3_: str
                    d_3149_v3_ = _dafny.CodePoint('j')
                    (d_3147_v2_).f7 = (((d_3147_v2_.f7).set(default__.safeIndex(default__.fm1(globalState), len(d_3147_v2_.f7)), d_3149_v3_)) + (d_3147_v2_.f7) if p3 else _dafny.SeqWithoutIsStrInference([d_3149_v3_ for d_3150_i1_ in range(default__.abs(430))]))
                    d_3151_v4_: _dafny.Array
                    def lambda149_(d_3152_i2_):
                        return (d_3152_i2_) + (self.f4)

                    init87_ = lambda149_
                    nw549_ = _dafny.Array(None, 4)
                    for i0_87_ in range(nw549_.length(0)):
                        nw549_[i0_87_] = init87_(i0_87_)
                    d_3151_v4_ = nw549_
                    index503_ = default__.safeIndex(261, (d_3151_v4_).length(0))
                    (d_3151_v4_)[index503_] = len(d_3145_v0_)
                    index504_ = default__.safeIndex(261, (d_3151_v4_).length(0))
                    rhs469_ = self.f5
                    rhs470_ = d_3146_v1_
                    lhs295_ = d_3151_v4_
                    lhs296_ = default__.safeIndex(261, (d_3151_v4_).length(0))
                    lhs295_[lhs296_] = rhs469_
                    d_3146_v1_ = rhs470_
                    d_3153_v5_: _dafny.Seq
                    d_3153_v5_ = _dafny.SeqWithoutIsStrInference([d_3146_v1_])
                    if ((d_3153_v5_) + (d_3153_v5_)) < (d_3153_v5_):
                        d_3154_v6_: _dafny.MultiSet
                        d_3154_v6_ = _dafny.MultiSet([d_3146_v1_])
                        d_3155_v7_: C6
                        nw550_ = C6()
                        nw550_.ctor__(p2, (d_3154_v6_) - (d_3154_v6_), self.f5)
                        d_3155_v7_ = nw550_
                        d_3156_v8_: C4
                        nw551_ = C4()
                        nw551_.ctor__((d_3151_v4_)[default__.safeIndex(261, (d_3151_v4_).length(0))])
                        d_3156_v8_ = nw551_
                        r0 = (p3 if True else self.f6)
                        d_3157_v9_: D5
                        d_3157_v9_ = D5_DC16(self.f6, not(p3), p2, self.f5, not((p1) == (self.f4)))
                        d_3157_v9_ = d_3157_v9_
                        index505_ = default__.safeIndex(380, (d_3146_v1_).length(0))
                        (d_3146_v1_)[index505_] = not(p3)
                        index506_ = default__.safeIndex(380, (d_3146_v1_).length(0))
                        (d_3146_v1_)[index506_] = (d_3155_v7_).f8
                    elif True:
                        d_3158_v10_: _dafny.Array
                        d_3158_v10_ = d_3151_v4_
                        d_3151_v4_ = (d_3158_v10_)
                        d_3159_v11_: D12
                        d_3159_v11_ = D12_DC37(p2, self.f6, p2)
                        index507_ = default__.safeIndex(261, (d_3151_v4_).length(0))
                        (d_3151_v4_)[index507_] = (-197) - (len(default__.fm16(p2, _dafny.SeqWithoutIsStrInference([d_3147_v2_.f7, d_3147_v2_.f7]), (d_3159_v11_).cf65, self.f4, globalState)))
                        (self).f5 = default__.safeDivisionInt(750, p1)
                        index508_ = default__.safeIndex(867, (d_3146_v1_).length(0))
                        (d_3146_v1_)[index508_] = not (p3) or (p0)
                        d_3160_v12_: _dafny.MultiSet
                        d_3160_v12_ = _dafny.MultiSet([p0, p3])
                        d_3161_v13_: _dafny.MultiSet
                        d_3161_v13_ = _dafny.MultiSet([(p1) + (self.f4), (d_3151_v4_)[default__.safeIndex(261, (d_3151_v4_).length(0))], default__.safeDivisionInt((d_3160_v12_).cardinality, self.f4), self.f5, self.f5])
                        index509_ = default__.safeIndex(867, (d_3146_v1_).length(0))
                        index510_ = default__.safeIndex(261, (d_3151_v4_).length(0))
                        index511_ = default__.safeIndex(261, (d_3151_v4_).length(0))
                        rhs471_ = not(p0)
                        rhs472_ = (self.f5) - (self.f4)
                        rhs473_ = p2
                        rhs474_ = ((d_3161_v13_)[self.f4] if (self.f4) in (d_3161_v13_) else p1)
                        rhs475_ = (p1) * (self.f5)
                        lhs297_ = d_3146_v1_
                        lhs298_ = default__.safeIndex(867, (d_3146_v1_).length(0))
                        lhs299_ = d_3151_v4_
                        lhs300_ = default__.safeIndex(261, (d_3151_v4_).length(0))
                        lhs301_ = d_3151_v4_
                        lhs302_ = default__.safeIndex(261, (d_3151_v4_).length(0))
                        lhs303_ = self
                        lhs297_[lhs298_] = rhs471_
                        lhs299_[lhs300_] = rhs472_
                        r0 = rhs473_
                        lhs301_[lhs302_] = rhs474_
                        lhs303_.f5 = rhs475_
                        d_3162_v14_: _dafny.Array
                        nw552_ = _dafny.Array(None, 16)
                        nw552_[int(0)] = self.f6
                        nw552_[int(1)] = (d_3146_v1_)[default__.safeIndex(867, (d_3146_v1_).length(0))]
                        nw552_[int(2)] = p3
                        nw552_[int(3)] = p3
                        nw552_[int(4)] = p3
                        nw552_[int(5)] = p0
                        nw552_[int(6)] = not(self.f6)
                        nw552_[int(7)] = not((d_3146_v1_)[default__.safeIndex(867, (d_3146_v1_).length(0))])
                        nw552_[int(8)] = self.f6
                        nw552_[int(9)] = self.f6
                        nw552_[int(10)] = False
                        nw552_[int(11)] = p2
                        nw552_[int(12)] = True
                        nw552_[int(13)] = (d_3146_v1_)[default__.safeIndex(867, (d_3146_v1_).length(0))]
                        nw552_[int(14)] = p0
                        nw552_[int(15)] = p2
                        d_3162_v14_ = nw552_
                        d_3163_v15_: _dafny.MultiSet
                        d_3163_v15_ = _dafny.MultiSet([d_3162_v14_, d_3162_v14_])
                        d_3164_v16_: C6
                        nw553_ = C6()
                        nw553_.ctor__(False, (d_3163_v15_).intersection(d_3163_v15_), (self.f5) * (p1))
                        d_3164_v16_ = nw553_
                    pass
            pass
        d_3165_v17_: _dafny.Array
        def lambda150_(d_3166_i3_):
            return _dafny.CodePoint('c')

        init88_ = lambda150_
        nw554_ = _dafny.Array(None, 4)
        for i0_88_ in range(nw554_.length(0)):
            nw554_[i0_88_] = init88_(i0_88_)
        d_3165_v17_ = nw554_
        d_3167_v18_: _dafny.MultiSet
        d_3167_v18_ = _dafny.MultiSet([d_3165_v17_])
        d_3168_v19_: _dafny.Map
        d_3168_v19_ = _dafny.Map({self.f5: self.f4})
        d_3169_v20_: _dafny.Set
        d_3169_v20_ = _dafny.Set({p0, p2, (d_3147_v2_).fm5(d_3145_v0_, p2, globalState)})
        d_3170_v21_: _dafny.Seq
        d_3170_v21_ = _dafny.SeqWithoutIsStrInference([d_3168_v19_])
        d_3171_v22_: _dafny.Set
        d_3171_v22_ = _dafny.Set({(d_3168_v19_ if self.f6 else _dafny.Map({len(d_3169_v20_): self.f4})), (d_3170_v21_)[default__.safeIndex(self.f5, len(d_3170_v21_))]})
        d_3172_v23_: _dafny.MultiSet
        d_3172_v23_ = _dafny.MultiSet([self.f4])
        rhs476_ = d_3167_v18_
        rhs477_ = default__.fm52(p1, globalState)
        rhs478_ = default__.fm33(not(p0), globalState)
        rhs479_ = self.f6
        lhs304_ = self
        d_3167_v18_ = rhs476_
        d_3171_v22_ = rhs477_
        d_3172_v23_ = rhs478_
        lhs304_.f6 = rhs479_
        d_3173_v24_: D3
        d_3173_v24_ = D3_DC10(self.f4, self.f5, p3, 685, p2)
        d_3174_v25_: _dafny.Array
        nw555_ = _dafny.Array(int(0), 17)
        d_3174_v25_ = nw555_
        d_3175_v26_: _dafny.Array
        nw556_ = _dafny.Array(_dafny.MultiSet({}), 21)
        d_3175_v26_ = nw556_
        d_3176_v27_: _dafny.MultiSet
        d_3176_v27_ = _dafny.MultiSet([d_3175_v26_, d_3175_v26_, d_3175_v26_, d_3175_v26_])
        d_3177_v28_: _dafny.Array
        nw557_ = _dafny.Array(None, 19)
        nw557_[int(0)] = d_3173_v24_
        nw557_[int(1)] = d_3173_v24_
        nw557_[int(2)] = d_3173_v24_
        nw557_[int(3)] = d_3173_v24_
        nw557_[int(4)] = d_3173_v24_
        nw557_[int(5)] = d_3173_v24_
        nw557_[int(6)] = d_3173_v24_
        nw557_[int(7)] = d_3173_v24_
        nw557_[int(8)] = d_3173_v24_
        nw557_[int(9)] = d_3173_v24_
        nw557_[int(10)] = d_3173_v24_
        nw557_[int(11)] = d_3173_v24_
        nw557_[int(12)] = d_3173_v24_
        nw557_[int(13)] = D3_DC10(p1, len(_dafny.SeqWithoutIsStrInference([d_3174_v25_])), p2, ((d_3176_v27_)[d_3175_v26_] if (d_3175_v26_) in (d_3176_v27_) else self.f5), p0)
        nw557_[int(14)] = d_3173_v24_
        nw557_[int(15)] = d_3173_v24_
        nw557_[int(16)] = d_3173_v24_
        nw557_[int(17)] = d_3173_v24_
        nw557_[int(18)] = d_3173_v24_
        d_3177_v28_ = nw557_
        index512_ = default__.safeIndex(531, (d_3177_v28_).length(0))
        (d_3177_v28_)[index512_] = d_3173_v24_
        index513_ = default__.safeIndex(531, (d_3177_v28_).length(0))
        def iife290_(_pat_let93_0):
            def iife291_(d_3178_dt__update__tmp_h0_):
                def iife292_(_pat_let94_0):
                    def iife293_(d_3179_dt__update_hcf18_h0_):
                        def iife294_(_pat_let95_0):
                            def iife295_(d_3180_dt__update_hcf20_h0_):
                                return D3_DC10((d_3178_dt__update__tmp_h0_).cf17, d_3179_dt__update_hcf18_h0_, (d_3178_dt__update__tmp_h0_).cf19, d_3180_dt__update_hcf20_h0_, (d_3178_dt__update__tmp_h0_).cf21)
                            return iife295_(_pat_let95_0)
                        return iife294_(self.f4)
                    return iife293_(_pat_let94_0)
                return iife292_(self.f5)
            return iife291_(_pat_let93_0)
        (d_3177_v28_)[index513_] = (iife290_(d_3173_v24_) if True else d_3173_v24_)
        if not(self.f6):
            (self).f4 = (0) - (default__.safeDivisionInt((self.f5) + (self.f5), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_3181_i4_ in range(default__.abs(632))]))))
            if (p1) <= (self.f5):
                d_3146_v1_ = d_3146_v1_
                d_3182_v29_: _dafny.Set
                d_3182_v29_ = _dafny.Set({d_3146_v1_})
                r0 = (d_3182_v29_).ispropersubset(_dafny.Set({d_3146_v1_}))
                d_3183_v30_: D12
                d_3183_v30_ = D12_DC38(self.f5)
                d_3184_v31_: str
                d_3184_v31_ = _dafny.CodePoint('q')
                d_3185_v32_: C3
                nw558_ = C3()
                nw558_.ctor__((d_3183_v30_).cf68, (_dafny.SeqWithoutIsStrInference([d_3184_v31_])) + (d_3147_v2_.f7), default__.safeDivisionInt(p1, self.f4))
                d_3185_v32_ = nw558_
                d_3186_v33_: _dafny.Map
                d_3186_v33_ = _dafny.Map({p0: d_3147_v2_.f7})
                d_3186_v33_ = (d_3186_v33_).set(False, (d_3185_v32_).f12)
                d_3187_v34_: _dafny.Map
                d_3187_v34_ = _dafny.Map({(d_3185_v32_).f11: p3})
                d_3188_v35_: _dafny.Map
                d_3188_v35_ = _dafny.Map({p0: (d_3183_v30_).cf68})
                d_3189_v36_: D0
                d_3189_v36_ = D0_DC1((d_3185_v32_).f11, (default__.fm25(globalState)).cardinality, p0)
                d_3190_v37_: _dafny.Set
                d_3190_v37_ = _dafny.Set({p1})
                d_3191_v38_: _dafny.Map
                d_3191_v38_ = _dafny.Map({(0) - (((d_3188_v35_)[(d_3189_v36_).cf3] if ((d_3189_v36_).cf3) in (d_3188_v35_) else len(d_3190_v37_))): d_3174_v25_})
                d_3192_v39_: int
                d_3193_v40_: int
                d_3194_v41_: _dafny.Array
                d_3195_v42_: bool
                out67_: int
                out68_: int
                out69_: _dafny.Array
                out70_: bool
                out67_, out68_, out69_, out70_ = default__.m0((d_3187_v34_).set(len(_dafny.Map({132: p2})), p3), _dafny.SeqWithoutIsStrInference([(d_3185_v32_).f11, 338, p1]), d_3191_v38_, globalState)
                d_3192_v39_ = out67_
                d_3193_v40_ = out68_
                d_3194_v41_ = out69_
                d_3195_v42_ = out70_
            elif True:
                d_3196_v43_: D7
                d_3196_v43_ = D7_DC25(p2, p3, self.f4)
                index514_ = default__.safeIndex(175, (d_3174_v25_).length(0))
                (d_3174_v25_)[index514_] = (d_3196_v43_).cf47
                d_3197_v44_: _dafny.Map
                d_3197_v44_ = _dafny.Map({self.f6: self.f4})
                index515_ = default__.safeIndex(175, (d_3174_v25_).length(0))
                rhs480_ = self.f4
                rhs481_ = ((d_3197_v44_)[p3] if (p3) in (d_3197_v44_) else self.f4)
                lhs305_ = self
                lhs306_ = d_3174_v25_
                lhs307_ = default__.safeIndex(175, (d_3174_v25_).length(0))
                lhs305_.f4 = rhs480_
                lhs306_[lhs307_] = rhs481_
                (self).f5 = (0) - (default__.fm1(globalState))
                d_3198_v45_: _dafny.Map
                d_3198_v45_ = _dafny.Map({self.f4: d_3146_v1_})
                d_3199_v46_: _dafny.MultiSet
                d_3199_v46_ = _dafny.MultiSet([d_3146_v1_, d_3146_v1_])
                d_3200_v47_: T0
                nw559_ = C6()
                nw559_.ctor__(p3, d_3199_v46_, self.f5)
                d_3200_v47_ = nw559_
                d_3201_v48_: _dafny.Map
                d_3201_v48_ = _dafny.Map({d_3200_v47_: p2})
                d_3202_v49_: _dafny.Map
                d_3202_v49_ = _dafny.Map({((d_3198_v45_)[len(d_3201_v48_)] if (len(d_3201_v48_)) in (d_3198_v45_) else d_3146_v1_): _dafny.SeqWithoutIsStrInference([self.f6])})
                d_3203_v50_: _dafny.Seq
                d_3203_v50_ = _dafny.SeqWithoutIsStrInference([p3])
                d_3202_v49_ = (d_3202_v49_).set(d_3146_v1_, (d_3203_v50_) + (d_3203_v50_))
                d_3204_v51_: C5
                nw560_ = C5()
                nw560_.ctor__(d_3200_v47_.f4)
                d_3204_v51_ = nw560_
                d_3205_v52_: _dafny.Map
                d_3205_v52_ = _dafny.Map({self.f5: (d_3204_v51_ if p2 else d_3204_v51_)})
                d_3206_v53_: D4
                d_3206_v53_ = D4_DC12(default__.fm23(False, globalState))
                d_3207_v54_: _dafny.Seq
                d_3207_v54_ = _dafny.SeqWithoutIsStrInference([len(d_3203_v50_), self.f5, len(d_3169_v20_), len((d_3206_v53_).cf23), self.f4])
                d_3205_v52_ = (d_3205_v52_).set((0) - (len(d_3207_v54_)), d_3204_v51_)
                d_3208_v55_: _dafny.Map
                d_3208_v55_ = _dafny.Map({p0: p0})
                d_3209_v56_: _dafny.MultiSet
                d_3209_v56_ = _dafny.MultiSet([((d_3208_v55_)[p3] if (p3) in (d_3208_v55_) else p3)])
                d_3210_v57_: D5
                d_3210_v57_ = D5_DC17(d_3209_v56_, self.f6)
                d_3211_v58_: _dafny.Seq
                d_3211_v58_ = _dafny.SeqWithoutIsStrInference([d_3208_v55_])
                (self).f6 = (not((_dafny.Map({d_3210_v57_: len(_dafny.SeqWithoutIsStrInference([self.f5, len(d_3211_v58_), p1]))})) == (_dafny.Map({d_3210_v57_: d_3200_v47_.f4}))) if p3 else not(not(p2)))
            d_3212_v59_: _dafny.Array
            nw561_ = _dafny.Array(_dafny.Array(None, 0), 9)
            d_3212_v59_ = nw561_
            d_3213_v60_: _dafny.Array
            d_3213_v60_ = d_3212_v59_
            d_3214_v61_: _dafny.Map
            d_3214_v61_ = _dafny.Map({d_3213_v60_: p1})
            d_3215_v62_: _dafny.Map
            d_3215_v62_ = _dafny.Map({(_dafny.Map({d_3214_v61_: not(self.f6)}) if p2 else _dafny.Map({(d_3214_v61_).set(d_3213_v60_, ((d_3168_v19_)[self.f4] if (self.f4) in (d_3168_v19_) else p1)): p3})): p3})
            d_3216_v63_: _dafny.Map
            d_3216_v63_ = _dafny.Map({d_3214_v61_: False})
            d_3217_v64_: _dafny.Map
            d_3217_v64_ = _dafny.Map({False: self.f5})
            d_3218_v65_: _dafny.Seq
            d_3218_v65_ = _dafny.SeqWithoutIsStrInference([d_3217_v64_])
            d_3219_v66_: _dafny.MultiSet
            d_3219_v66_ = _dafny.MultiSet([self.f6])
            d_3220_v67_: _dafny.Map
            d_3220_v67_ = _dafny.Map({p3: d_3217_v64_})
            d_3221_v68_: _dafny.Array
            nw562_ = _dafny.Array(None, 14)
            nw562_[int(0)] = d_3217_v64_
            nw562_[int(1)] = d_3217_v64_
            nw562_[int(2)] = (d_3218_v65_)[default__.safeIndex((d_3219_v66_).cardinality, len(d_3218_v65_))]
            nw562_[int(3)] = d_3217_v64_
            nw562_[int(4)] = ((d_3220_v67_)[True] if (True) in (d_3220_v67_) else d_3217_v64_)
            nw562_[int(5)] = d_3217_v64_
            nw562_[int(6)] = d_3217_v64_
            nw562_[int(7)] = d_3217_v64_
            nw562_[int(8)] = d_3217_v64_
            nw562_[int(9)] = d_3217_v64_
            nw562_[int(10)] = d_3217_v64_
            nw562_[int(11)] = (d_3217_v64_).set(True, p1)
            nw562_[int(12)] = d_3217_v64_
            nw562_[int(13)] = d_3217_v64_
            d_3221_v68_ = nw562_
            d_3222_v69_: D23
            d_3222_v69_ = D23_DC60(d_3221_v68_)
            pat_let_tv70_ = d_3221_v68_
            d_3223_v70_: _dafny.Seq
            def iife296_(_pat_let96_0):
                def iife297_(d_3224_dt__update__tmp_h1_):
                    def iife298_(_pat_let97_0):
                        def iife299_(d_3225_dt__update_hcf108_h0_):
                            return D23_DC60(d_3225_dt__update_hcf108_h0_)
                        return iife299_(_pat_let97_0)
                    return iife298_(pat_let_tv70_)
                return iife297_(_pat_let96_0)
            d_3223_v70_ = _dafny.SeqWithoutIsStrInference([iife296_(d_3222_v69_)])
            if ((d_3215_v62_)[(d_3216_v63_) | (_dafny.Map({d_3214_v61_: p0}))] if ((d_3216_v63_) | (_dafny.Map({d_3214_v61_: p0}))) in (d_3215_v62_) else (d_3222_v69_) in (d_3223_v70_)):
                index516_ = default__.safeIndex(490, (d_3146_v1_).length(0))
                (d_3146_v1_)[index516_] = p2
                index517_ = default__.safeIndex(490, (d_3146_v1_).length(0))
                (d_3146_v1_)[index517_] = True
                d_3226_v71_: _dafny.Array
                nw563_ = _dafny.Array(_dafny.Seq({}), 24)
                d_3226_v71_ = nw563_
                d_3226_v71_ = d_3226_v71_
                d_3227_v72_: D4
                d_3227_v72_ = D4_DC12(d_3147_v2_.f7)
                d_3228_v73_: _dafny.Seq
                d_3228_v73_ = _dafny.SeqWithoutIsStrInference([d_3147_v2_.f7, (d_3227_v72_).cf23])
                d_3229_v74_: _dafny.Array
                nw564_ = _dafny.Array(None, 14)
                nw564_[int(0)] = d_3147_v2_.f7
                nw564_[int(1)] = (d_3147_v2_.f7) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wo")))
                nw564_[int(2)] = (d_3147_v2_.f7) + (d_3145_v0_)
                nw564_[int(3)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_3230_i5_ in range(default__.abs(222))])
                nw564_[int(4)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fuwrnbxif"))) + (d_3147_v2_.f7)
                nw564_[int(5)] = d_3145_v0_
                nw564_[int(6)] = d_3147_v2_.f7
                nw564_[int(7)] = d_3145_v0_
                nw564_[int(8)] = d_3147_v2_.f7
                nw564_[int(9)] = (d_3228_v73_)[default__.safeIndex(self.f5, len(d_3228_v73_))]
                nw564_[int(10)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_3231_i6_ in range(default__.abs(-990))])
                nw564_[int(11)] = (d_3147_v2_.f7) + (d_3147_v2_.f7)
                nw564_[int(12)] = (d_3147_v2_.f7) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qdw")))
                nw564_[int(13)] = default__.fm23(not(False), globalState)
                d_3229_v74_ = nw564_
                index518_ = default__.safeIndex(632, (d_3229_v74_).length(0))
                (d_3229_v74_)[index518_] = d_3147_v2_.f7
                d_3232_v75_: _dafny.Seq
                d_3232_v75_ = _dafny.SeqWithoutIsStrInference([p3, (p2) and (p3), self.f6, p0, p0])
                d_3233_v76_: D11
                d_3233_v76_ = D11_DC35((0) - (default__.safeDivisionInt((0) - (p1), default__.fm1(globalState))), p3, p3)
                d_3234_v77_: _dafny.MultiSet
                d_3234_v77_ = _dafny.MultiSet([d_3168_v19_])
                index519_ = default__.safeIndex(632, (d_3229_v74_).length(0))
                rhs482_ = (d_3145_v0_) + ((d_3147_v2_.f7) + (d_3145_v0_))
                rhs483_ = default__.safeModuloInt((158) + (len(d_3147_v2_.f7)), default__.safeDivisionInt((0) - (self.f5), ((d_3234_v77_)[_dafny.Map({self.f4: self.f4})] if (_dafny.Map({self.f4: self.f4})) in (d_3234_v77_) else self.f4)))
                rhs484_ = (0) - (((len(d_3232_v75_)) + (p1) if p3 else p1))
                rhs485_ = d_3232_v75_
                rhs486_ = D11_DC35(((d_3172_v23_)[len(default__.fm23(p3, globalState))] if (len(default__.fm23(p3, globalState))) in (d_3172_v23_) else p1), (p3) == (self.f6), p2)
                lhs308_ = d_3229_v74_
                lhs309_ = default__.safeIndex(632, (d_3229_v74_).length(0))
                lhs310_ = self
                lhs311_ = self
                lhs308_[lhs309_] = rhs482_
                lhs310_.f5 = rhs483_
                lhs311_.f5 = rhs484_
                d_3232_v75_ = rhs485_
                d_3233_v76_ = rhs486_
                d_3235_v78_: _dafny.Array
                nw565_ = _dafny.Array(False, 19)
                d_3235_v78_ = nw565_
                d_3235_v78_ = d_3235_v78_
                d_3236_v79_: D3
                d_3236_v79_ = D3_DC10(len(d_3147_v2_.f7), self.f5, p2, len(d_3232_v75_), p0)
                d_3237_v80_: D3
                d_3237_v80_ = D3_DC11(d_3236_v79_)
                d_3237_v80_ = d_3237_v80_
            elif True:
                d_3238_v81_: C8
                nw566_ = C8()
                nw566_.ctor__()
                d_3238_v81_ = nw566_
                d_3239_v82_: _dafny.MultiSet
                d_3239_v82_ = _dafny.MultiSet([d_3238_v81_, d_3238_v81_, d_3238_v81_])
                d_3239_v82_ = _dafny.MultiSet([d_3238_v81_, d_3238_v81_, d_3238_v81_])
                d_3240_v83_: C1
                nw567_ = C1()
                nw567_.ctor__(p1, self.f5, p1, (d_3147_v2_).fm5(d_3147_v2_.f7, p0, globalState))
                d_3240_v83_ = nw567_
                (self).f4 = (0) - (p1)
                d_3241_v84_: D7
                d_3241_v84_ = D7_DC27(len(d_3147_v2_.f7), self.f4, 513)
                d_3242_v85_: _dafny.Seq
                d_3242_v85_ = _dafny.SeqWithoutIsStrInference([d_3241_v84_])
                d_3242_v85_ = (d_3242_v85_) + (d_3242_v85_)
                index520_ = default__.safeIndex(254, (d_3146_v1_).length(0))
                (d_3146_v1_)[index520_] = self.f6
                index521_ = default__.safeIndex(254, (d_3146_v1_).length(0))
                (d_3146_v1_)[index521_] = False
            r0 = not(p3)
            if p2:
                d_3243_v86_: _dafny.Seq
                d_3243_v86_ = _dafny.SeqWithoutIsStrInference([p2, p0, p0])
                d_3244_v87_: _dafny.Map
                d_3244_v87_ = _dafny.Map({True: not(self.f6)})
                (self).f6 = (d_3243_v86_)[default__.safeIndex((0) - (((0) - (len(d_3244_v87_))) - (self.f5)), len(d_3243_v86_))]
                d_3245_v88_: str
                d_3245_v88_ = _dafny.CodePoint('u')
                d_3246_v89_: _dafny.Map
                d_3246_v89_ = _dafny.Map({p2: d_3245_v88_})
                rhs487_ = p2
                rhs488_ = ((d_3246_v89_)[p0] if (p0) in (d_3246_v89_) else d_3245_v88_)
                rhs489_ = 523
                rhs490_ = (self.f5 if p3 else self.f4)
                rhs491_ = p1
                lhs312_ = self
                lhs313_ = self
                lhs314_ = self
                lhs315_ = self
                lhs312_.f6 = rhs487_
                d_3245_v88_ = rhs488_
                lhs313_.f5 = rhs489_
                lhs314_.f4 = rhs490_
                lhs315_.f5 = rhs491_
                d_3247_v90_: _dafny.Array
                d_3247_v90_ = d_3174_v25_
                d_3248_v91_: _dafny.Array
                nw568_ = _dafny.Array(None, 12)
                nw568_[int(0)] = d_3174_v25_
                nw568_[int(1)] = d_3174_v25_
                nw568_[int(2)] = d_3174_v25_
                nw568_[int(3)] = d_3174_v25_
                nw568_[int(4)] = d_3174_v25_
                nw568_[int(5)] = d_3174_v25_
                nw568_[int(6)] = (d_3174_v25_ if self.f6 else d_3174_v25_)
                nw568_[int(7)] = d_3174_v25_
                nw568_[int(8)] = d_3174_v25_
                nw568_[int(9)] = d_3174_v25_
                nw568_[int(10)] = d_3174_v25_
                nw568_[int(11)] = (d_3247_v90_)
                d_3248_v91_ = nw568_
                index522_ = default__.safeIndex(944, (d_3248_v91_).length(0))
                (d_3248_v91_)[index522_] = d_3174_v25_
                index523_ = default__.safeIndex(944, (d_3248_v91_).length(0))
                (d_3248_v91_)[index523_] = d_3174_v25_
                (self).f5 = self.f4
                d_3249_v92_: _dafny.Map
                d_3249_v92_ = _dafny.Map({d_3245_v88_: -649})
                d_3250_v93_: _dafny.Seq
                d_3250_v93_ = _dafny.SeqWithoutIsStrInference([d_3249_v92_, d_3249_v92_])
                d_3250_v93_ = (((d_3250_v93_) + (d_3250_v93_)) + ((d_3250_v93_) + (d_3250_v93_))).set(default__.safeIndex(self.f5, len(((d_3250_v93_) + (d_3250_v93_)) + ((d_3250_v93_) + (d_3250_v93_)))), (_dafny.Map({d_3245_v88_: self.f5})) | (d_3249_v92_))
            elif True:
                d_3251_v94_: _dafny.Array
                nw569_ = _dafny.Array(_dafny.Array(None, 0), 27)
                d_3251_v94_ = nw569_
                d_3252_v95_: T0
                nw570_ = C7()
                nw570_.ctor__(d_3147_v2_.f7, self.f5)
                d_3252_v95_ = nw570_
                d_3253_v96_: T0
                d_3253_v96_ = d_3252_v95_
                d_3254_v97_: _dafny.Array
                nw571_ = _dafny.Array(None, 14)
                nw571_[int(0)] = d_3252_v95_
                nw571_[int(1)] = d_3252_v95_
                nw571_[int(2)] = (d_3253_v96_)
                nw571_[int(3)] = d_3252_v95_
                nw571_[int(4)] = d_3252_v95_
                nw571_[int(5)] = d_3252_v95_
                nw571_[int(6)] = d_3252_v95_
                nw571_[int(7)] = d_3252_v95_
                nw571_[int(8)] = d_3252_v95_
                nw571_[int(9)] = d_3252_v95_
                nw571_[int(10)] = d_3252_v95_
                nw571_[int(11)] = d_3252_v95_
                nw571_[int(12)] = d_3252_v95_
                nw571_[int(13)] = d_3252_v95_
                d_3254_v97_ = nw571_
                index524_ = default__.safeIndex(992, (d_3251_v94_).length(0))
                (d_3251_v94_)[index524_] = d_3254_v97_
                index525_ = default__.safeIndex(992, (d_3251_v94_).length(0))
                (d_3251_v94_)[index525_] = d_3254_v97_
                d_3255_v98_: C4
                nw572_ = C4()
                nw572_.ctor__(p1)
                d_3255_v98_ = nw572_
                d_3256_v99_: C0
                nw573_ = C0()
                nw573_.ctor__(self.f4)
                d_3256_v99_ = nw573_
                d_3257_v100_: _dafny.Map
                d_3257_v100_ = _dafny.Map({p1: p2})
                d_3258_v101_: _dafny.Seq
                d_3258_v101_ = _dafny.SeqWithoutIsStrInference([443])
                d_3259_v102_: _dafny.Map
                d_3259_v102_ = _dafny.Map({self.f4: d_3174_v25_})
                d_3260_v103_: int
                d_3261_v104_: int
                d_3262_v105_: _dafny.Array
                d_3263_v106_: bool
                out71_: int
                out72_: int
                out73_: _dafny.Array
                out74_: bool
                out71_, out72_, out73_, out74_ = default__.m0((d_3257_v100_) | (d_3257_v100_), (d_3258_v101_) + (_dafny.SeqWithoutIsStrInference([self.f5, p1, d_3252_v95_.f4, d_3256_v99_.f18, d_3252_v95_.f4])), (d_3259_v102_).set(((d_3219_v66_)[p2] if (p2) in (d_3219_v66_) else p1), d_3174_v25_), globalState)
                d_3260_v103_ = out71_
                d_3261_v104_ = out72_
                d_3262_v105_ = out73_
                d_3263_v106_ = out74_
                d_3260_v103_ = (d_3261_v104_) * (24)
        elif True:
            (self).f6 = p3
            d_3264_v107_: _dafny.Seq
            d_3264_v107_ = _dafny.SeqWithoutIsStrInference([d_3146_v1_, d_3146_v1_, d_3146_v1_, d_3146_v1_])
            d_3264_v107_ = d_3264_v107_
            index526_ = default__.safeIndex(760, (d_3174_v25_).length(0))
            (d_3174_v25_)[index526_] = default__.fm1(globalState)
            d_3265_v108_: _dafny.Seq
            d_3265_v108_ = _dafny.SeqWithoutIsStrInference([d_3145_v0_])
            d_3266_v109_: _dafny.Set
            d_3266_v109_ = _dafny.Set({self.f4})
            index527_ = default__.safeIndex(760, (d_3174_v25_).length(0))
            rhs492_ = (self.f5) * (len((d_3265_v108_) + (d_3265_v108_)))
            rhs493_ = (d_3266_v109_) == ((d_3266_v109_) | (d_3266_v109_))
            rhs494_ = (0) - (default__.safeDivisionInt(self.f5, 495))
            rhs495_ = (p1) - (p1)
            lhs316_ = d_3174_v25_
            lhs317_ = default__.safeIndex(760, (d_3174_v25_).length(0))
            lhs318_ = self
            lhs319_ = self
            lhs316_[lhs317_] = rhs492_
            r0 = rhs493_
            lhs318_.f5 = rhs494_
            lhs319_.f5 = rhs495_
            if False:
                (d_3147_v2_).f7 = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mku"))
                index528_ = default__.safeIndex(760, (d_3174_v25_).length(0))
                (d_3174_v25_)[index528_] = (0) - (self.f4)
                (self).f4 = len((d_3147_v2_.f7) + (d_3147_v2_.f7))
                d_3267_v110_: _dafny.Map
                d_3267_v110_ = _dafny.Map({d_3146_v1_: d_3169_v20_})
                d_3267_v110_ = d_3267_v110_
                d_3268_v111_: _dafny.Map
                d_3268_v111_ = _dafny.Map({self.f4: _dafny.Map({-898: (d_3174_v25_)[default__.safeIndex(760, (d_3174_v25_).length(0))]})})
                d_3269_v112_: _dafny.Map
                d_3269_v112_ = _dafny.Map({d_3145_v0_: len(d_3268_v111_)})
                d_3269_v112_ = d_3269_v112_
            elif True:
                d_3270_v113_: _dafny.Seq
                d_3270_v113_ = _dafny.SeqWithoutIsStrInference([p0])
                d_3271_v114_: _dafny.Map
                d_3271_v114_ = _dafny.Map({d_3270_v113_: d_3146_v1_})
                d_3271_v114_ = d_3271_v114_
                d_3146_v1_ = d_3146_v1_
                d_3146_v1_ = d_3146_v1_
                (d_3147_v2_).f7 = (d_3147_v2_.f7) + (d_3147_v2_.f7)
                d_3272_v115_: _dafny.Array
                def lambda151_(d_3273_i7_):
                    return D22_DC59(self.f4)

                init89_ = lambda151_
                nw574_ = _dafny.Array(None, 10)
                for i0_89_ in range(nw574_.length(0)):
                    nw574_[i0_89_] = init89_(i0_89_)
                d_3272_v115_ = nw574_
                d_3274_v116_: _dafny.MultiSet
                d_3274_v116_ = _dafny.MultiSet([d_3272_v115_])
                d_3168_v19_ = (d_3168_v19_).set(((0) - (p1)) * (self.f4), (d_3274_v116_).cardinality)
            d_3275_v117_: _dafny.Seq
            d_3275_v117_ = _dafny.SeqWithoutIsStrInference([p2, p0, p0, not(p2)])
            d_3276_v118_: C7
            nw575_ = C7()
            nw575_.ctor__(d_3147_v2_.f7, len(_dafny.SeqWithoutIsStrInference([len(d_3275_v117_), self.f4, (d_3174_v25_)[default__.safeIndex(760, (d_3174_v25_).length(0))]])))
            d_3276_v118_ = nw575_
        d_3277_v119_: _dafny.Seq
        d_3277_v119_ = _dafny.SeqWithoutIsStrInference([p1, self.f5])
        d_3278_v120_: _dafny.Seq
        d_3278_v120_ = _dafny.SeqWithoutIsStrInference([self.f6, p2])
        d_3279_v121_: D1
        d_3279_v121_ = D1_DC4(len(d_3277_v119_), d_3278_v120_, p1)
        d_3280_v122_: C0
        nw576_ = C0()
        nw576_.ctor__(len(d_3277_v119_))
        d_3280_v122_ = nw576_
        d_3281_v123_: D18
        d_3281_v123_ = D18_DC48(self.f6, p2, d_3279_v121_, d_3280_v122_)
        d_3282_v124_: D12
        d_3282_v124_ = D12_DC38((self.f4 if (d_3281_v123_).cf85 else self.f4))
        d_3282_v124_ = d_3282_v124_
        d_3283_v125_: _dafny.MultiSet
        d_3283_v125_ = _dafny.MultiSet([d_3168_v19_, d_3168_v19_])
        d_3284_v126_: _dafny.Seq
        d_3284_v126_ = _dafny.SeqWithoutIsStrInference([d_3145_v0_])
        r0 = (d_3283_v125_).issubset(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(_dafny.Map({d_3280_v122_.f18: self.f5})).set(d_3280_v122_.f18, len(d_3284_v126_)), (d_3168_v19_).set(self.f4, self.f4)])))
        return r0

    def m4(self, p0, p1, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r2: _dafny.Array = _dafny.Array(None, 0)
        d_3285_v0_: _dafny.Seq
        d_3285_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gpxm"))
        d_3286_v1_: str
        d_3286_v1_ = _dafny.CodePoint('k')
        d_3287_v2_: _dafny.Map
        d_3287_v2_ = _dafny.Map({d_3285_v0_: d_3286_v1_})
        d_3287_v2_ = (d_3287_v2_).set((d_3285_v0_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q"))), d_3286_v1_)
        (self).f6 = self.f6
        (self).f4 = default__.safeModuloInt((108) + (self.f4), len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_3288_i0_ in range(default__.abs(961))])) + (default__.fm23(p1, globalState))))
        d_3289_v3_: _dafny.Array
        nw577_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 28)
        d_3289_v3_ = nw577_
        index529_ = default__.safeIndex(6, (d_3289_v3_).length(0))
        (d_3289_v3_)[index529_] = (d_3285_v0_) + (d_3285_v0_)
        index530_ = default__.safeIndex(6, (d_3289_v3_).length(0))
        (d_3289_v3_)[index530_] = (d_3285_v0_).set(default__.safeIndex(self.f5, len(d_3285_v0_)), d_3286_v1_)
        (self).f5 = self.f5
        d_3290_v4_: _dafny.Seq
        d_3290_v4_ = _dafny.SeqWithoutIsStrInference([self.f6])
        d_3291_v5_: D1
        d_3291_v5_ = D1_DC4(807, d_3290_v4_, p0)
        d_3292_v6_: C0
        nw578_ = C0()
        nw578_.ctor__(self.f5)
        d_3292_v6_ = nw578_
        d_3293_v7_: _dafny.Seq
        d_3293_v7_ = _dafny.SeqWithoutIsStrInference([False, (D18_DC48(p1, self.f6, d_3291_v5_, d_3292_v6_)).cf85])
        d_3294_v8_: _dafny.Seq
        d_3294_v8_ = _dafny.SeqWithoutIsStrInference([d_3293_v7_, d_3290_v4_, d_3293_v7_, (d_3290_v4_ if self.f6 else d_3290_v4_)])
        d_3294_v8_ = d_3294_v8_
        r0 = default__.fm51((193) <= (self.f4), self.f6, globalState)
        r1 = d_3285_v0_
        d_3295_v9_: _dafny.Map
        d_3295_v9_ = _dafny.Map({self.f6: self.f6})
        d_3296_v10_: _dafny.MultiSet
        d_3296_v10_ = _dafny.MultiSet([not(self.f6)])
        d_3297_v11_: _dafny.Map
        d_3297_v11_ = _dafny.Map({p0: (d_3289_v3_)[default__.safeIndex(6, (d_3289_v3_).length(0))]})
        d_3298_v12_: T1
        nw579_ = C1()
        nw579_.ctor__(p0, len(d_3297_v11_), self.f4, False)
        d_3298_v12_ = nw579_
        d_3299_v13_: _dafny.MultiSet
        d_3299_v13_ = _dafny.MultiSet([d_3298_v12_, d_3298_v12_])
        d_3300_v14_: _dafny.Map
        d_3300_v14_ = _dafny.Map({_dafny.MultiSet([self.f5, len(d_3285_v0_), (d_3296_v10_).cardinality, len(default__.fm23(self.f6, globalState)), self.f4]): (d_3299_v13_).set(d_3298_v12_, default__.abs(self.f4))})
        d_3301_v15_: _dafny.Map
        d_3301_v15_ = _dafny.Map({(d_3296_v10_).cardinality: len(_dafny.Set({p0}))})
        d_3302_v16_: _dafny.Array
        nw580_ = _dafny.Array(None, 17)
        nw580_[int(0)] = self.f5
        nw580_[int(1)] = (len(d_3295_v9_)) - (len(d_3300_v14_))
        nw580_[int(2)] = self.f5
        nw580_[int(3)] = ((d_3296_v10_).cardinality) + (860)
        nw580_[int(4)] = d_3292_v6_.f18
        nw580_[int(5)] = 439
        nw580_[int(6)] = p0
        nw580_[int(7)] = -558
        nw580_[int(8)] = d_3298_v12_.f4
        nw580_[int(9)] = d_3292_v6_.f18
        nw580_[int(10)] = (0) - (d_3298_v12_.f4)
        nw580_[int(11)] = 684
        nw580_[int(12)] = d_3292_v6_.f18
        nw580_[int(13)] = p0
        nw580_[int(14)] = (len(d_3301_v15_)) * (self.f4)
        nw580_[int(15)] = (d_3292_v6_.f18 if self.f6 else d_3292_v6_.f18)
        nw580_[int(16)] = self.f5
        d_3302_v16_ = nw580_
        r2 = d_3302_v16_
        return r0, r1, r2


class C10:
    def  __init__(self):
        self.f3: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C10"
    def ctor__(self, f3):
        (self).f3 = f3

    def m1(self, globalState):
        r0: _dafny.Set = _dafny.Set({})
        r1: int = int(0)
        r2: bool = False
        d_3303_v0_: int
        d_3303_v0_ = 746
        d_3304_v1_: _dafny.Seq
        d_3304_v1_ = _dafny.SeqWithoutIsStrInference([d_3303_v0_])
        d_3305_v2_: D0
        d_3305_v2_ = D0_DC0(d_3304_v1_)
        d_3306_v3_: _dafny.Map
        d_3306_v3_ = _dafny.Map({D0_DC0(_dafny.SeqWithoutIsStrInference([d_3303_v0_])): d_3305_v2_})
        d_3307_v4_: D0
        d_3307_v4_ = D0_DC0(d_3304_v1_)
        d_3308_v5_: D0
        d_3308_v5_ = D0_DC2(((d_3306_v3_)[d_3307_v4_] if (d_3307_v4_) in (d_3306_v3_) else d_3305_v2_))
        d_3309_v7_: _dafny.Seq
        d_3309_v7_ = _dafny.SeqWithoutIsStrInference([d_3308_v5_, d_3308_v5_, d_3308_v5_, d_3308_v5_])
        def iife300_():
            coll104_ = _dafny.Map()
            compr_104_: D0
            for compr_104_ in (d_3309_v7_).Elements:
                d_3310_v6_: D0 = compr_104_
                if (d_3310_v6_) in (d_3309_v7_):
                    coll104_[d_3310_v6_] = d_3303_v0_
            return _dafny.Map(coll104_)
        if not((d_3308_v5_) in (iife300_()
        )):
            d_3311_v8_: _dafny.Map
            d_3311_v8_ = _dafny.Map({self.f3: d_3303_v0_})
            r1 = ((d_3311_v8_)[self.f3] if (self.f3) in (d_3311_v8_) else d_3303_v0_)
            d_3312_v9_: _dafny.Map
            d_3312_v9_ = _dafny.Map({d_3303_v0_: self.f3})
            d_3313_v10_: _dafny.MultiSet
            d_3313_v10_ = _dafny.MultiSet([d_3312_v9_, d_3312_v9_])
            d_3314_v11_: _dafny.MultiSet
            d_3314_v11_ = _dafny.MultiSet([self.f3, self.f3, default__.fm0(d_3303_v0_, d_3303_v0_, globalState)])
            d_3315_v12_: _dafny.Set
            d_3315_v12_ = _dafny.Set({self.f3, True})
            d_3316_v13_: D0
            d_3316_v13_ = D0_DC1(d_3303_v0_, len(d_3315_v12_), not(self.f3))
            rhs496_ = (d_3313_v10_).ispropersubset((d_3313_v10_) - (d_3313_v10_))
            rhs497_ = ((d_3314_v11_)[not((d_3316_v13_).cf3)] if (not((d_3316_v13_).cf3)) in (d_3314_v11_) else 265)
            r2 = rhs496_
            d_3303_v0_ = rhs497_
            d_3317_v14_: _dafny.Map
            d_3317_v14_ = _dafny.Map({433: (d_3303_v0_) - (d_3303_v0_)})
            d_3317_v14_ = (d_3317_v14_).set(585, (0) - (((d_3314_v11_).set(self.f3, default__.abs(d_3303_v0_))).cardinality))
            d_3318_v15_: _dafny.Array
            nw581_ = _dafny.Array(int(0), 15)
            d_3318_v15_ = nw581_
            d_3318_v15_ = d_3318_v15_
            d_3319_v16_: _dafny.Seq
            d_3319_v16_ = _dafny.SeqWithoutIsStrInference([self.f3])
            (self).f3 = not((d_3319_v16_)[default__.safeIndex(d_3303_v0_, len(d_3319_v16_))])
        elif True:
            d_3320_v17_: _dafny.Array
            nw582_ = _dafny.Array(False, 5)
            d_3320_v17_ = nw582_
            d_3321_v18_: str
            d_3321_v18_ = _dafny.CodePoint('l')
            d_3322_v19_: _dafny.Map
            d_3322_v19_ = _dafny.Map({(d_3320_v17_ if self.f3 else d_3320_v17_): d_3321_v18_})
            d_3322_v19_ = (_dafny.Map({d_3320_v17_: d_3321_v18_})) | (d_3322_v19_)
            d_3323_v20_: _dafny.Array
            nw583_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 4)
            d_3323_v20_ = nw583_
            index531_ = default__.safeIndex(588, (d_3323_v20_).length(0))
            (d_3323_v20_)[index531_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yewpd"))
            d_3324_v21_: _dafny.MultiSet
            d_3324_v21_ = _dafny.MultiSet([935, -381])
            d_3325_v22_: _dafny.Array
            nw584_ = _dafny.Array(int(0), 8)
            d_3325_v22_ = nw584_
            index532_ = default__.safeIndex(724, (d_3325_v22_).length(0))
            (d_3325_v22_)[index532_] = d_3303_v0_
            d_3326_v23_: _dafny.Seq
            d_3326_v23_ = _dafny.SeqWithoutIsStrInference([self.f3])
            d_3327_v24_: _dafny.Array
            nw585_ = _dafny.Array(None, 6)
            nw585_[int(0)] = d_3326_v23_
            nw585_[int(1)] = (_dafny.SeqWithoutIsStrInference([self.f3])) + (d_3326_v23_)
            nw585_[int(2)] = d_3326_v23_
            nw585_[int(3)] = d_3326_v23_
            nw585_[int(4)] = (d_3326_v23_).set(default__.safeIndex(len(d_3304_v1_), len(d_3326_v23_)), self.f3)
            nw585_[int(5)] = d_3326_v23_
            d_3327_v24_ = nw585_
            d_3328_v25_: _dafny.Seq
            d_3328_v25_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "juvstfji"))
            index533_ = default__.safeIndex(588, (d_3323_v20_).length(0))
            index534_ = default__.safeIndex(724, (d_3325_v22_).length(0))
            rhs498_ = ((d_3328_v25_) + (d_3328_v25_)) + ((d_3328_v25_) + (d_3328_v25_))
            rhs499_ = d_3324_v21_
            rhs500_ = (0) - (d_3303_v0_)
            rhs501_ = d_3327_v24_
            lhs320_ = d_3323_v20_
            lhs321_ = default__.safeIndex(588, (d_3323_v20_).length(0))
            lhs322_ = d_3325_v22_
            lhs323_ = default__.safeIndex(724, (d_3325_v22_).length(0))
            lhs320_[lhs321_] = rhs498_
            d_3324_v21_ = rhs499_
            lhs322_[lhs323_] = rhs500_
            d_3327_v24_ = rhs501_
            d_3329_v26_: _dafny.Map
            d_3329_v26_ = _dafny.Map({True: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bsd")))})
            d_3330_v27_: _dafny.Set
            d_3330_v27_ = _dafny.Set({(0) - (d_3303_v0_)})
            d_3331_v28_: D0
            d_3331_v28_ = D0_DC1(len(d_3330_v27_), d_3303_v0_, self.f3)
            d_3329_v26_ = (d_3329_v26_ if ((d_3331_v28_).cf3 if self.f3 else self.f3) else (d_3329_v26_) | (d_3329_v26_))
            if (d_3330_v27_).ispropersubset(d_3330_v27_):
                d_3303_v0_ = ((d_3325_v22_)[default__.safeIndex(724, (d_3325_v22_).length(0))] if self.f3 else (d_3325_v22_)[default__.safeIndex(724, (d_3325_v22_).length(0))])
                d_3332_v29_: _dafny.Array
                nw586_ = _dafny.Array(None, 1)
                nw586_[int(0)] = d_3325_v22_
                d_3332_v29_ = nw586_
                d_3332_v29_ = d_3332_v29_
                (self).m2(d_3328_v25_, self.f3, globalState)
                d_3333_v30_: _dafny.Array
                nw587_ = _dafny.Array(_dafny.MultiSet({}), 16)
                d_3333_v30_ = nw587_
                d_3334_v31_: _dafny.Map
                d_3334_v31_ = _dafny.Map({(0) - ((d_3325_v22_)[default__.safeIndex(724, (d_3325_v22_).length(0))]): d_3303_v0_})
                d_3335_v32_: _dafny.MultiSet
                d_3335_v32_ = _dafny.MultiSet([d_3334_v31_, d_3334_v31_, d_3334_v31_])
                index535_ = default__.safeIndex(96, (d_3333_v30_).length(0))
                (d_3333_v30_)[index535_] = d_3335_v32_
                index536_ = default__.safeIndex(96, (d_3333_v30_).length(0))
                (d_3333_v30_)[index536_] = _dafny.MultiSet([_dafny.Map({(d_3325_v22_)[default__.safeIndex(724, (d_3325_v22_).length(0))]: d_3303_v0_}), _dafny.Map({(d_3325_v22_)[default__.safeIndex(724, (d_3325_v22_).length(0))]: (d_3304_v1_)[default__.safeIndex(d_3303_v0_, len(d_3304_v1_))]})])
                (self).f3 = (d_3326_v23_)[default__.safeIndex((d_3325_v22_)[default__.safeIndex(724, (d_3325_v22_).length(0))], len(d_3326_v23_))]
            elif True:
                d_3320_v17_ = d_3320_v17_
                d_3321_v18_ = ((d_3323_v20_)[default__.safeIndex(588, (d_3323_v20_).length(0))])[default__.safeIndex(d_3303_v0_, len((d_3323_v20_)[default__.safeIndex(588, (d_3323_v20_).length(0))]))]
                (self).f3 = not (self.f3) or (((d_3325_v22_)[default__.safeIndex(724, (d_3325_v22_).length(0))]) < (42))
                r1 = default__.safeDivisionInt((d_3325_v22_)[default__.safeIndex(724, (d_3325_v22_).length(0))], d_3303_v0_)
                d_3336_v33_: _dafny.Map
                d_3336_v33_ = _dafny.Map({(0) - (len(d_3304_v1_)): self.f3})
                d_3337_v34_: _dafny.Set
                d_3337_v34_ = _dafny.Set({d_3336_v33_, d_3336_v33_})
                d_3338_v36_: _dafny.Set
                d_3338_v36_ = _dafny.Set({d_3321_v18_, d_3321_v18_})
                index537_ = default__.safeIndex(724, (d_3325_v22_).length(0))
                def iife301_():
                    coll105_ = _dafny.Set()
                    compr_105_: str
                    for compr_105_ in (d_3338_v36_).Elements:
                        d_3339_v37_: str = compr_105_
                        if (d_3339_v37_) in (d_3338_v36_):
                            coll105_ = coll105_.union(_dafny.Set([d_3339_v37_]))
                    return _dafny.Set(coll105_)
                def iife302_():
                    coll106_ = _dafny.Set()
                    compr_106_: str
                    for compr_106_ in (default__.fm3((d_3325_v22_)[default__.safeIndex(724, (d_3325_v22_).length(0))], self.f3, (d_3325_v22_)[default__.safeIndex(724, (d_3325_v22_).length(0))], globalState)).keys.Elements:
                        d_3340_v35_: str = compr_106_
                        if (d_3340_v35_) in (default__.fm3((d_3325_v22_)[default__.safeIndex(724, (d_3325_v22_).length(0))], self.f3, (d_3325_v22_)[default__.safeIndex(724, (d_3325_v22_).length(0))], globalState)):
                            coll106_ = coll106_.union(_dafny.Set([d_3340_v35_]))
                    return _dafny.Set(coll106_)
                rhs502_ = (d_3325_v22_)[default__.safeIndex(724, (d_3325_v22_).length(0))]
                rhs503_ = self.f3
                rhs504_ = ((iife301_()
                ) - (d_3338_v36_)).ispropersubset(iife302_()
                )
                rhs505_ = default__.fm4(globalState)
                lhs324_ = d_3325_v22_
                lhs325_ = default__.safeIndex(724, (d_3325_v22_).length(0))
                lhs326_ = self
                lhs324_[lhs325_] = rhs502_
                r2 = rhs503_
                lhs326_.f3 = rhs504_
                d_3337_v34_ = rhs505_
            index538_ = default__.safeIndex(37, (d_3320_v17_).length(0))
            (d_3320_v17_)[index538_] = self.f3
            index539_ = default__.safeIndex(37, (d_3320_v17_).length(0))
            rhs506_ = d_3325_v22_
            rhs507_ = default__.fm0(d_3303_v0_, (d_3325_v22_)[default__.safeIndex(724, (d_3325_v22_).length(0))], globalState)
            rhs508_ = self.f3
            lhs327_ = d_3320_v17_
            lhs328_ = default__.safeIndex(37, (d_3320_v17_).length(0))
            lhs329_ = self
            d_3325_v22_ = rhs506_
            lhs327_[lhs328_] = rhs507_
            lhs329_.f3 = rhs508_
        d_3341_i0_: int
        d_3341_i0_ = 0
        with _dafny.label("12"):
            while self.f3:
                with _dafny.c_label("12"):
                    if (d_3341_i0_) >= (100):
                        raise _dafny.Break("12")
                    d_3341_i0_ = (d_3341_i0_) + (1)
                    source33_ = d_3307_v4_
                    if source33_.is_DC1:
                        d_3342___mcc_h0_ = source33_.cf1
                        d_3343___mcc_h1_ = source33_.cf2
                        d_3344___mcc_h2_ = source33_.cf3
                        d_3345_cf3_ = d_3344___mcc_h2_
                        d_3346_cf2_ = d_3343___mcc_h1_
                        d_3347_cf1_ = d_3342___mcc_h0_
                        d_3348_v38_: str
                        d_3348_v38_ = _dafny.CodePoint('v')
                        d_3349_v39_: _dafny.Array
                        nw588_ = _dafny.Array(int(0), 21)
                        d_3349_v39_ = nw588_
                        d_3350_v40_: _dafny.Seq
                        d_3350_v40_ = _dafny.SeqWithoutIsStrInference([d_3349_v39_, d_3349_v39_])
                        d_3351_v41_: _dafny.Map
                        d_3351_v41_ = _dafny.Map({d_3348_v38_: (d_3350_v40_) + (d_3350_v40_)})
                        d_3351_v41_ = (d_3351_v41_).set(d_3348_v38_, d_3350_v40_)
                        d_3352_v42_: _dafny.Map
                        d_3352_v42_ = _dafny.Map({d_3345_cf3_: self.f3})
                        d_3352_v42_ = (d_3352_v42_).set(self.f3, self.f3)
                        d_3353_v43_: _dafny.Map
                        d_3353_v43_ = _dafny.Map({self.f3: d_3346_cf2_})
                        d_3353_v43_ = (d_3353_v43_).set(True, default__.fm1(globalState))
                        d_3348_v38_ = (_dafny.CodePoint('j') if self.f3 else d_3348_v38_)
                    elif source33_.is_DC0:
                        d_3354___mcc_h3_ = source33_.cf0
                        d_3355_cf0_ = d_3354___mcc_h3_
                        d_3356_v44_: _dafny.Map
                        d_3356_v44_ = _dafny.Map({d_3303_v0_: self.f3})
                        d_3357_v45_: str
                        d_3357_v45_ = _dafny.CodePoint('o')
                        d_3358_v46_: _dafny.Seq
                        d_3358_v46_ = _dafny.SeqWithoutIsStrInference([d_3357_v45_, d_3357_v45_, d_3357_v45_, d_3357_v45_, d_3357_v45_])
                        d_3359_v47_: C3
                        nw589_ = C3()
                        nw589_.ctor__(len((d_3356_v44_) | (d_3356_v44_)), d_3358_v46_, (0) - (d_3303_v0_))
                        d_3359_v47_ = nw589_
                        d_3360_v48_: _dafny.Array
                        def lambda152_(d_3361_v1_):
                            def lambda153_(d_3362_i1_):
                                return (d_3362_i1_) * (len(d_3361_v1_))

                            return lambda153_

                        init90_ = lambda152_(d_3304_v1_)
                        nw590_ = _dafny.Array(None, 7)
                        for i0_90_ in range(nw590_.length(0)):
                            nw590_[i0_90_] = init90_(i0_90_)
                        d_3360_v48_ = nw590_
                        index540_ = default__.safeIndex(163, (d_3360_v48_).length(0))
                        (d_3360_v48_)[index540_] = d_3303_v0_
                        index541_ = default__.safeIndex(163, (d_3360_v48_).length(0))
                        (d_3360_v48_)[index541_] = (0) - ((d_3359_v47_).f11)
                        d_3358_v46_ = d_3358_v46_
                        d_3363_v49_: _dafny.Array
                        nw591_ = _dafny.Array(None, 5)
                        nw591_[int(0)] = self.f3
                        nw591_[int(1)] = self.f3
                        nw591_[int(2)] = self.f3
                        nw591_[int(3)] = (d_3359_v47_).fm5((d_3359_v47_).f12, self.f3, globalState)
                        nw591_[int(4)] = self.f3
                        d_3363_v49_ = nw591_
                        d_3364_v50_: _dafny.Set
                        d_3364_v50_ = _dafny.Set({d_3363_v49_, d_3363_v49_, d_3363_v49_})
                        d_3364_v50_ = d_3364_v50_
                    elif True:
                        d_3365___mcc_h4_ = source33_.cf4
                        d_3366_cf4_ = d_3365___mcc_h4_
                        d_3367_v51_: _dafny.Seq
                        d_3367_v51_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mso"))
                        d_3368_v52_: _dafny.Map
                        d_3368_v52_ = _dafny.Map({(d_3304_v1_)[default__.safeIndex(d_3303_v0_, len(d_3304_v1_))]: d_3367_v51_})
                        d_3368_v52_ = (d_3368_v52_).set(d_3303_v0_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v")))
                        (self).f3 = default__.fm0(d_3303_v0_, d_3303_v0_, globalState)
                        d_3369_v53_: _dafny.Map
                        d_3369_v53_ = _dafny.Map({not(self.f3): (self.f3) == (self.f3)})
                        d_3369_v53_ = (d_3369_v53_).set(self.f3, self.f3)
                        (self).m2(d_3367_v51_, self.f3, globalState)
                    d_3370_v54_: _dafny.Array
                    def lambda154_(d_3371_i2_):
                        return (not(not(self.f3))) and (self.f3)

                    init91_ = lambda154_
                    nw592_ = _dafny.Array(None, 22)
                    for i0_91_ in range(nw592_.length(0)):
                        nw592_[i0_91_] = init91_(i0_91_)
                    d_3370_v54_ = nw592_
                    d_3370_v54_ = d_3370_v54_
                    d_3372_v55_: _dafny.Map
                    d_3372_v55_ = _dafny.Map({d_3303_v0_: self.f3})
                    d_3373_v56_: _dafny.Map
                    d_3373_v56_ = _dafny.Map({d_3303_v0_: d_3303_v0_})
                    d_3374_v57_: _dafny.Array
                    nw593_ = _dafny.Array(int(0), 25)
                    d_3374_v57_ = nw593_
                    d_3375_v58_: _dafny.Map
                    d_3375_v58_ = _dafny.Map({len(d_3373_v56_): d_3374_v57_})
                    d_3376_v59_: int
                    d_3377_v60_: int
                    d_3378_v61_: _dafny.Array
                    d_3379_v62_: bool
                    out75_: int
                    out76_: int
                    out77_: _dafny.Array
                    out78_: bool
                    out75_, out76_, out77_, out78_ = default__.m0(d_3372_v55_, (D0_DC0(d_3304_v1_)).cf0, (d_3375_v58_) | (d_3375_v58_), globalState)
                    d_3376_v59_ = out75_
                    d_3377_v60_ = out76_
                    d_3378_v61_ = out77_
                    d_3379_v62_ = out78_
                    d_3380_v63_: _dafny.Array
                    nw594_ = _dafny.Array(None, 20)
                    d_3380_v63_ = nw594_
                    d_3381_v64_: T0
                    nw595_ = C4()
                    nw595_.ctor__(d_3377_v60_)
                    d_3381_v64_ = nw595_
                    index542_ = default__.safeIndex(411, (d_3380_v63_).length(0))
                    (d_3380_v63_)[index542_] = d_3381_v64_
                    d_3382_v65_: _dafny.Seq
                    d_3382_v65_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "krgduhkps"))
                    d_3383_v66_: str
                    d_3383_v66_ = _dafny.CodePoint('e')
                    d_3384_v67_: _dafny.Seq
                    d_3384_v67_ = _dafny.SeqWithoutIsStrInference([d_3370_v54_, d_3370_v54_, d_3370_v54_])
                    d_3385_v68_: _dafny.MultiSet
                    d_3385_v68_ = _dafny.MultiSet([d_3370_v54_, d_3370_v54_, (d_3384_v67_)[default__.safeIndex(d_3376_v59_, len(d_3384_v67_))]])
                    index543_ = default__.safeIndex(411, (d_3380_v63_).length(0))
                    nw596_ = C6()
                    nw596_.ctor__((len((d_3382_v65_).set(default__.safeIndex(d_3377_v60_, len(d_3382_v65_)), d_3383_v66_))) <= (d_3377_v60_), d_3385_v68_, d_3376_v59_)
                    (d_3380_v63_)[index543_] = nw596_
                    pass
            pass
        d_3386_v69_: _dafny.MultiSet
        d_3386_v69_ = _dafny.MultiSet([default__.fm0(d_3303_v0_, d_3303_v0_, globalState)])
        d_3387_v70_: _dafny.Seq
        d_3387_v70_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vejkvyf"))
        d_3388_v71_: _dafny.Seq
        d_3388_v71_ = _dafny.SeqWithoutIsStrInference([self.f3, self.f3, self.f3])
        d_3389_v72_: _dafny.Set
        d_3389_v72_ = _dafny.Set({self.f3, self.f3})
        d_3390_v73_: _dafny.Array
        nw597_ = _dafny.Array(None, 25)
        nw597_[int(0)] = (d_3386_v69_).cardinality
        nw597_[int(1)] = default__.fm1(globalState)
        nw597_[int(2)] = d_3303_v0_
        nw597_[int(3)] = len(d_3387_v70_)
        nw597_[int(4)] = d_3303_v0_
        nw597_[int(5)] = d_3303_v0_
        nw597_[int(6)] = d_3303_v0_
        nw597_[int(7)] = d_3303_v0_
        nw597_[int(8)] = d_3303_v0_
        nw597_[int(9)] = default__.fm1(globalState)
        nw597_[int(10)] = d_3303_v0_
        nw597_[int(11)] = default__.fm1(globalState)
        nw597_[int(12)] = d_3303_v0_
        nw597_[int(13)] = d_3303_v0_
        nw597_[int(14)] = d_3303_v0_
        nw597_[int(15)] = d_3303_v0_
        nw597_[int(16)] = (D19_DC51(d_3303_v0_, (0) - (d_3303_v0_))).cf90
        nw597_[int(17)] = d_3303_v0_
        nw597_[int(18)] = len(d_3388_v71_)
        nw597_[int(19)] = d_3303_v0_
        nw597_[int(20)] = (d_3304_v1_)[default__.safeIndex(d_3303_v0_, len(d_3304_v1_))]
        nw597_[int(21)] = d_3303_v0_
        nw597_[int(22)] = len(d_3388_v71_)
        nw597_[int(23)] = d_3303_v0_
        nw597_[int(24)] = len(d_3389_v72_)
        d_3390_v73_ = nw597_
        d_3391_v74_: _dafny.Array
        d_3391_v74_ = d_3390_v73_
        d_3391_v74_ = d_3391_v74_
        d_3392_v75_: str
        d_3392_v75_ = _dafny.CodePoint('m')
        d_3392_v75_ = d_3392_v75_
        d_3392_v75_ = d_3392_v75_
        d_3393_v76_: _dafny.Array
        nw598_ = _dafny.Array(False, 15)
        d_3393_v76_ = nw598_
        d_3394_v77_: D3
        d_3394_v77_ = D3_DC8(d_3392_v75_, d_3393_v76_)
        source34_ = d_3394_v77_
        if source34_.is_DC8:
            d_3395___mcc_h5_ = source34_.cf13
            d_3396___mcc_h6_ = source34_.cf14
            d_3397_cf14_ = d_3396___mcc_h6_
            d_3398_cf13_ = d_3395___mcc_h5_
            d_3399_v78_: _dafny.Array
            nw599_ = _dafny.Array(_dafny.Set({}), 28)
            d_3399_v78_ = nw599_
            d_3400_v79_: T0
            nw600_ = C7()
            nw600_.ctor__(d_3387_v70_, d_3303_v0_)
            d_3400_v79_ = nw600_
            d_3401_v80_: _dafny.Map
            d_3401_v80_ = _dafny.Map({d_3400_v79_.f4: d_3400_v79_.f4})
            d_3402_v81_: _dafny.Map
            d_3402_v81_ = _dafny.Map({d_3400_v79_: d_3401_v80_})
            d_3403_v82_: _dafny.Set
            d_3403_v82_ = _dafny.Set({len(((d_3402_v81_)[d_3400_v79_] if (d_3400_v79_) in (d_3402_v81_) else d_3401_v80_))})
            index544_ = default__.safeIndex(696, (d_3399_v78_).length(0))
            (d_3399_v78_)[index544_] = d_3403_v82_
            index545_ = default__.safeIndex(696, (d_3399_v78_).length(0))
            (d_3399_v78_)[index545_] = d_3403_v82_
            rhs509_ = (d_3387_v70_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rbjxp")))
            rhs510_ = (d_3400_v79_.f4) * ((0) - (d_3400_v79_.f4))
            rhs511_ = (d_3388_v71_)[default__.safeIndex(308, len(d_3388_v71_))]
            lhs330_ = d_3400_v79_
            d_3387_v70_ = rhs509_
            lhs330_.f4 = rhs510_
            r2 = rhs511_
            d_3404_v83_: _dafny.Map
            d_3404_v83_ = _dafny.Map({d_3400_v79_.f4: self.f3})
            d_3404_v83_ = d_3404_v83_
            rhs512_ = self.f3
            rhs513_ = self.f3
            lhs331_ = self
            r2 = rhs512_
            lhs331_.f3 = rhs513_
        elif source34_.is_DC9:
            d_3405___mcc_h7_ = source34_.cf15
            d_3406___mcc_h8_ = source34_.cf16
            d_3407_cf16_ = d_3406___mcc_h8_
            d_3408_cf15_ = d_3405___mcc_h7_
            (self).f3 = (d_3386_v69_).isdisjoint(d_3386_v69_)
            d_3409_v84_: _dafny.Map
            d_3409_v84_ = _dafny.Map({d_3408_cf15_: d_3389_v72_})
            r2 = (d_3408_cf15_) not in (d_3409_v84_)
            def iife303_():
                coll107_ = _dafny.Set()
                compr_107_: D14
                for compr_107_ in (_dafny.Map({D14_DC41(_dafny.Set({d_3408_cf15_, (D0_DC1(len(d_3304_v1_), d_3407_cf16_, d_3408_cf15_)).cf3, d_3408_cf15_})): D17_DC46(False, d_3407_cf16_, d_3303_v0_, self.f3)})).keys.Elements:
                    d_3410_v85_: D14 = compr_107_
                    if (d_3410_v85_) in (_dafny.Map({D14_DC41(_dafny.Set({d_3408_cf15_, (D0_DC1(len(d_3304_v1_), d_3407_cf16_, d_3408_cf15_)).cf3, d_3408_cf15_})): D17_DC46(False, d_3407_cf16_, d_3303_v0_, self.f3)})):
                        coll107_ = coll107_.union(_dafny.Set([d_3410_v85_]))
                return _dafny.Set(coll107_)
            d_3407_cf16_ = len((d_3304_v1_) + (_dafny.SeqWithoutIsStrInference([len(iife303_()
) for d_3411_i3_ in range(default__.abs(-241))])))
            d_3412_v86_: _dafny.Array
            nw601_ = _dafny.Array(D23.default()(), 28)
            d_3412_v86_ = nw601_
            d_3413_v87_: _dafny.Map
            d_3413_v87_ = _dafny.Map({False: (d_3407_cf16_) <= (d_3303_v0_)})
            rhs514_ = ((d_3413_v87_)[d_3408_cf15_] if (d_3408_cf15_) in (d_3413_v87_) else not(self.f3))
            rhs515_ = d_3412_v86_
            rhs516_ = (d_3386_v69_) | ((d_3386_v69_) | (d_3386_v69_))
            rhs517_ = d_3408_cf15_
            d_3408_cf15_ = rhs514_
            d_3412_v86_ = rhs515_
            d_3386_v69_ = rhs516_
            r2 = rhs517_
        elif source34_.is_DC10:
            d_3414___mcc_h9_ = source34_.cf17
            d_3415___mcc_h10_ = source34_.cf18
            d_3416___mcc_h11_ = source34_.cf19
            d_3417___mcc_h12_ = source34_.cf20
            d_3418___mcc_h13_ = source34_.cf21
            d_3419_cf21_ = d_3418___mcc_h13_
            d_3420_cf20_ = d_3417___mcc_h12_
            d_3421_cf19_ = d_3416___mcc_h11_
            d_3422_cf18_ = d_3415___mcc_h10_
            d_3423_cf17_ = d_3414___mcc_h9_
            index546_ = default__.safeIndex(356, (d_3390_v73_).length(0))
            (d_3390_v73_)[index546_] = (701) * (d_3423_cf17_)
            index547_ = default__.safeIndex(356, (d_3390_v73_).length(0))
            rhs518_ = d_3422_cf18_
            rhs519_ = d_3393_v76_
            lhs332_ = d_3390_v73_
            lhs333_ = default__.safeIndex(356, (d_3390_v73_).length(0))
            lhs332_[lhs333_] = rhs518_
            d_3393_v76_ = rhs519_
            d_3424_v88_: D4
            d_3424_v88_ = D4_DC12(d_3387_v70_)
            d_3425_v89_: D4
            d_3425_v89_ = D4_DC14(d_3424_v88_)
            pat_let_tv71_ = d_3387_v70_
            def iife304_(_pat_let98_0):
                def iife305_(d_3426_dt__update__tmp_h0_):
                    def iife306_(_pat_let99_0):
                        def iife307_(d_3427_dt__update_hcf27_h0_):
                            return D4_DC14(d_3427_dt__update_hcf27_h0_)
                        return iife307_(_pat_let99_0)
                    return iife306_(D4_DC12(pat_let_tv71_))
                return iife305_(_pat_let98_0)
            source35_ = iife304_(d_3425_v89_)
            if source35_.is_DC13:
                d_3428___mcc_h16_ = source35_.cf24
                d_3429___mcc_h17_ = source35_.cf25
                d_3430___mcc_h18_ = source35_.cf26
                d_3431_cf26_ = d_3430___mcc_h18_
                d_3432_cf25_ = d_3429___mcc_h17_
                d_3433_cf24_ = d_3428___mcc_h16_
                d_3419_cf21_ = not(d_3432_cf25_)
                index548_ = default__.safeIndex(759, (d_3393_v76_).length(0))
                (d_3393_v76_)[index548_] = (d_3386_v69_).isdisjoint(d_3386_v69_)
                d_3434_v90_: _dafny.Map
                d_3434_v90_ = _dafny.Map({self.f3: (-934) - (d_3420_cf20_)})
                index549_ = default__.safeIndex(759, (d_3393_v76_).length(0))
                rhs520_ = default__.safeModuloInt(916, d_3423_cf17_)
                rhs521_ = self.f3
                rhs522_ = d_3421_cf19_
                rhs523_ = self.f3
                rhs524_ = d_3434_v90_
                lhs334_ = d_3393_v76_
                lhs335_ = default__.safeIndex(759, (d_3393_v76_).length(0))
                lhs336_ = self
                r1 = rhs520_
                lhs334_[lhs335_] = rhs521_
                d_3419_cf21_ = rhs522_
                lhs336_.f3 = rhs523_
                d_3434_v90_ = rhs524_
                index550_ = default__.safeIndex(759, (d_3393_v76_).length(0))
                (d_3393_v76_)[index550_] = (d_3393_v76_)[default__.safeIndex(759, (d_3393_v76_).length(0))]
                d_3435_v91_: C7
                nw602_ = C7()
                nw602_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aax")), (_dafny.MultiSet([d_3432_cf25_, d_3419_cf21_])).cardinality)
                d_3435_v91_ = nw602_
                d_3436_v92_: _dafny.Map
                d_3436_v92_ = _dafny.Map({self.f3: d_3435_v91_})
                d_3437_v93_: _dafny.Seq
                d_3437_v93_ = _dafny.SeqWithoutIsStrInference([d_3435_v91_, ((d_3436_v92_)[(d_3393_v76_)[default__.safeIndex(759, (d_3393_v76_).length(0))]] if ((d_3393_v76_)[default__.safeIndex(759, (d_3393_v76_).length(0))]) in (d_3436_v92_) else d_3435_v91_)])
                index551_ = default__.safeIndex(356, (d_3390_v73_).length(0))
                (d_3390_v73_)[index551_] = (d_3304_v1_)[default__.safeIndex(len((_dafny.SeqWithoutIsStrInference([d_3435_v91_, d_3435_v91_])) + (d_3437_v93_)), len(d_3304_v1_))]
            elif source35_.is_DC12:
                d_3438___mcc_h19_ = source35_.cf23
                d_3439_cf23_ = d_3438___mcc_h19_
                (self).f3 = self.f3
                d_3440_v94_: _dafny.Map
                d_3440_v94_ = _dafny.Map({d_3422_cf18_: d_3303_v0_})
                d_3441_v95_: _dafny.Map
                d_3441_v95_ = _dafny.Map({d_3388_v71_: d_3419_cf21_})
                d_3442_v96_: _dafny.Map
                d_3442_v96_ = _dafny.Map({(len(d_3440_v94_)) > (d_3420_cf20_): d_3441_v95_})
                def iife308_():
                    coll108_ = _dafny.Map()
                    compr_108_: _dafny.Seq
                    for compr_108_ in (_dafny.Set({d_3388_v71_})).Elements:
                        d_3443_v97_: _dafny.Seq = compr_108_
                        if (d_3443_v97_) in (_dafny.Set({d_3388_v71_})):
                            coll108_[d_3443_v97_] = True
                    return _dafny.Map(coll108_)
                d_3442_v96_ = (d_3442_v96_).set(not(d_3419_cf21_), iife308_()
                )
                d_3444_v98_: D6
                d_3444_v98_ = D6_DC21(d_3303_v0_, d_3419_cf21_)
                d_3445_v99_: D6
                d_3445_v99_ = D6_DC23(d_3444_v98_)
                d_3445_v99_ = d_3445_v99_
                d_3387_v70_ = d_3439_cf23_
            elif True:
                d_3446___mcc_h20_ = source35_.cf27
                d_3447_cf27_ = d_3446___mcc_h20_
                d_3423_cf17_ = -855
                (self).f3 = self.f3
                d_3448_v100_: _dafny.Array
                nw603_ = _dafny.Array(_dafny.MultiSet({}), 26)
                d_3448_v100_ = nw603_
                d_3449_v101_: D5
                d_3449_v101_ = D5_DC15(d_3448_v100_)
                d_3450_v102_: _dafny.Map
                d_3450_v102_ = _dafny.Map({d_3449_v101_: d_3419_cf21_})
                d_3451_v103_: T0
                nw604_ = C1()
                nw604_.ctor__(d_3422_cf18_, d_3423_cf17_, 143, d_3419_cf21_)
                d_3451_v103_ = nw604_
                rhs525_ = (len(_dafny.SeqWithoutIsStrInference([d_3451_v103_]))) + (default__.safeDivisionInt((0) - (len(d_3387_v70_)), d_3420_cf20_))
                rhs526_ = d_3450_v102_
                rhs527_ = (0) - (len((d_3387_v70_) + (d_3387_v70_)))
                d_3422_cf18_ = rhs525_
                d_3450_v102_ = rhs526_
                d_3422_cf18_ = rhs527_
                d_3452_v104_: _dafny.Array
                nw605_ = _dafny.Array(int(0), 23)
                d_3452_v104_ = nw605_
                d_3453_v105_: _dafny.Map
                d_3453_v105_ = _dafny.Map({d_3419_cf21_: d_3452_v104_})
                d_3454_v106_: _dafny.Map
                d_3454_v106_ = _dafny.Map({_dafny.Set({_dafny.SeqWithoutIsStrInference([(d_3387_v70_)[default__.safeIndex(d_3451_v103_.f4, len(d_3387_v70_))] for d_3455_i4_ in range(default__.abs(203))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qfpctmww"))}): d_3452_v104_})
                d_3456_v107_: _dafny.Set
                d_3456_v107_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vwrfaf"))})
                d_3457_v108_: _dafny.Array
                nw606_ = _dafny.Array(None, 24)
                nw606_[int(0)] = d_3452_v104_
                nw606_[int(1)] = d_3452_v104_
                nw606_[int(2)] = d_3452_v104_
                nw606_[int(3)] = (d_3452_v104_ if True else d_3452_v104_)
                nw606_[int(4)] = d_3452_v104_
                nw606_[int(5)] = d_3452_v104_
                nw606_[int(6)] = d_3452_v104_
                nw606_[int(7)] = d_3452_v104_
                nw606_[int(8)] = ((d_3453_v105_)[d_3421_cf19_] if (d_3421_cf19_) in (d_3453_v105_) else d_3452_v104_)
                nw606_[int(9)] = d_3452_v104_
                nw606_[int(10)] = (d_3391_v74_)
                nw606_[int(11)] = d_3452_v104_
                nw606_[int(12)] = d_3452_v104_
                nw606_[int(13)] = d_3452_v104_
                nw606_[int(14)] = d_3452_v104_
                nw606_[int(15)] = d_3452_v104_
                nw606_[int(16)] = d_3452_v104_
                nw606_[int(17)] = d_3452_v104_
                nw606_[int(18)] = d_3452_v104_
                nw606_[int(19)] = d_3452_v104_
                nw606_[int(20)] = ((d_3454_v106_)[d_3456_v107_] if (d_3456_v107_) in (d_3454_v106_) else d_3452_v104_)
                nw606_[int(21)] = d_3452_v104_
                nw606_[int(22)] = d_3452_v104_
                nw606_[int(23)] = d_3452_v104_
                d_3457_v108_ = nw606_
                index552_ = default__.safeIndex(702, (d_3457_v108_).length(0))
                (d_3457_v108_)[index552_] = d_3452_v104_
                index553_ = default__.safeIndex(702, (d_3457_v108_).length(0))
                (d_3457_v108_)[index553_] = d_3452_v104_
            d_3458_v109_: _dafny.Map
            d_3458_v109_ = _dafny.Map({d_3419_cf21_: (d_3422_cf18_) > (d_3423_cf17_)})
            d_3458_v109_ = (d_3458_v109_).set(False, (d_3419_cf21_) or (True))
            d_3421_cf19_ = d_3421_cf19_
        elif source34_.is_DC7:
            d_3459___mcc_h14_ = source34_.cf12
            d_3460_cf12_ = d_3459___mcc_h14_
            d_3461_v110_: _dafny.Array
            def lambda155_(d_3462_v70_):
                def lambda156_(d_3463_i5_):
                    return d_3462_v70_

                return lambda156_

            init92_ = lambda155_(d_3387_v70_)
            nw607_ = _dafny.Array(None, 5)
            for i0_92_ in range(nw607_.length(0)):
                nw607_[i0_92_] = init92_(i0_92_)
            d_3461_v110_ = nw607_
            d_3464_v111_: _dafny.Seq
            d_3464_v111_ = _dafny.SeqWithoutIsStrInference([d_3461_v110_])
            d_3465_v112_: D8
            d_3465_v112_ = D8_DC30((d_3304_v1_)[default__.safeIndex(863, len(d_3304_v1_))], (d_3461_v110_ if default__.fm0(-395, d_3303_v0_, globalState) else (d_3464_v111_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xfoc"))), len(d_3464_v111_))]), self.f3)
            d_3466_v113_: _dafny.Seq
            d_3466_v113_ = _dafny.SeqWithoutIsStrInference([d_3388_v71_, d_3388_v71_, (d_3388_v71_) + (_dafny.SeqWithoutIsStrInference([self.f3]))])
            pat_let_tv72_ = d_3303_v0_
            def iife309_(_pat_let100_0):
                def iife310_(d_3467_dt__update__tmp_h1_):
                    def iife311_(_pat_let101_0):
                        def iife312_(d_3468_dt__update_hcf54_h0_):
                            return D8_DC30(d_3468_dt__update_hcf54_h0_, (d_3467_dt__update__tmp_h1_).cf55, (d_3467_dt__update__tmp_h1_).cf56)
                        return iife312_(_pat_let101_0)
                    return iife311_(len(_dafny.SeqWithoutIsStrInference([pat_let_tv72_, 533])))
                return iife310_(_pat_let100_0)
            rhs528_ = default__.fm1(globalState)
            rhs529_ = iife309_(d_3465_v112_)
            rhs530_ = d_3466_v113_
            d_3303_v0_ = rhs528_
            d_3465_v112_ = rhs529_
            d_3466_v113_ = rhs530_
            d_3390_v73_ = (d_3391_v74_)
            d_3469_v114_: C2
            nw608_ = C2()
            nw608_.ctor__(self.f3, d_3303_v0_)
            d_3469_v114_ = nw608_
            (self).f3 = d_3469_v114_.f13
        elif True:
            d_3470___mcc_h15_ = source34_.cf22
            d_3471_cf22_ = d_3470___mcc_h15_
            d_3472_v115_: D4
            d_3472_v115_ = D4_DC12(d_3387_v70_)
            d_3473_v116_: T2
            nw609_ = C1()
            nw609_.ctor__(len(_dafny.Set({self.f3, self.f3, self.f3, self.f3, self.f3})), d_3303_v0_, d_3303_v0_, ((d_3472_v115_).cf23) <= (d_3387_v70_))
            d_3473_v116_ = nw609_
            d_3473_v116_ = d_3473_v116_
            d_3474_v117_: _dafny.Array
            nw610_ = _dafny.Array(_dafny.Map({}), 26)
            d_3474_v117_ = nw610_
            d_3475_v118_: _dafny.Array
            nw611_ = _dafny.Array(None, 25)
            d_3475_v118_ = nw611_
            d_3476_v119_: C3
            nw612_ = C3()
            nw612_.ctor__(d_3303_v0_, d_3387_v70_, (d_3473_v116_).f15)
            d_3476_v119_ = nw612_
            index554_ = default__.safeIndex(99, (d_3475_v118_).length(0))
            (d_3475_v118_)[index554_] = d_3476_v119_
            d_3477_v120_: _dafny.Map
            d_3477_v120_ = _dafny.Map({934: d_3475_v118_})
            d_3478_v121_: _dafny.Seq
            d_3478_v121_ = _dafny.SeqWithoutIsStrInference([d_3475_v118_])
            d_3479_v122_: _dafny.Array
            nw613_ = _dafny.Array(None, 13)
            nw613_[int(0)] = d_3475_v118_
            nw613_[int(1)] = ((d_3477_v120_)[len((d_3476_v119_).f12)] if (len((d_3476_v119_).f12)) in (d_3477_v120_) else d_3475_v118_)
            nw613_[int(2)] = d_3475_v118_
            nw613_[int(3)] = d_3475_v118_
            nw613_[int(4)] = d_3475_v118_
            nw613_[int(5)] = d_3475_v118_
            nw613_[int(6)] = d_3475_v118_
            nw613_[int(7)] = d_3475_v118_
            nw613_[int(8)] = d_3475_v118_
            nw613_[int(9)] = d_3475_v118_
            nw613_[int(10)] = d_3475_v118_
            nw613_[int(11)] = d_3475_v118_
            nw613_[int(12)] = (d_3478_v121_)[default__.safeIndex((d_3473_v116_).f15, len(d_3478_v121_))]
            d_3479_v122_ = nw613_
            index555_ = default__.safeIndex(850, (d_3479_v122_).length(0))
            (d_3479_v122_)[index555_] = d_3475_v118_
            d_3480_v123_: _dafny.Array
            nw614_ = _dafny.Array(_dafny.Set({}), 18)
            d_3480_v123_ = nw614_
            index556_ = default__.safeIndex(99, (d_3475_v118_).length(0))
            index557_ = default__.safeIndex(850, (d_3479_v122_).length(0))
            rhs531_ = d_3474_v117_
            rhs532_ = (d_3476_v119_).f11
            rhs533_ = d_3476_v119_
            rhs534_ = ((d_3477_v120_)[(d_3476_v119_).f11] if ((d_3476_v119_).f11) in (d_3477_v120_) else d_3475_v118_)
            rhs535_ = d_3480_v123_
            lhs337_ = d_3475_v118_
            lhs338_ = default__.safeIndex(99, (d_3475_v118_).length(0))
            lhs339_ = d_3479_v122_
            lhs340_ = default__.safeIndex(850, (d_3479_v122_).length(0))
            d_3474_v117_ = rhs531_
            r1 = rhs532_
            lhs337_[lhs338_] = rhs533_
            lhs339_[lhs340_] = rhs534_
            d_3480_v123_ = rhs535_
            index558_ = default__.safeIndex(828, (d_3393_v76_).length(0))
            (d_3393_v76_)[index558_] = (d_3303_v0_) > (len(d_3387_v70_))
            index559_ = default__.safeIndex(828, (d_3393_v76_).length(0))
            (d_3393_v76_)[index559_] = self.f3
            d_3481_v124_: _dafny.Map
            d_3481_v124_ = _dafny.Map({(d_3476_v119_).f11: True})
            r1 = len(d_3481_v124_)
        r0 = d_3389_v72_
        d_3482_v125_: _dafny.Map
        d_3482_v125_ = _dafny.Map({self.f3: _dafny.SeqWithoutIsStrInference([d_3303_v0_ for d_3483_i6_ in range(default__.abs(-678))])})
        r1 = len((default__.fm53(self.f3, d_3303_v0_, globalState)).set(((d_3482_v125_)[(d_3388_v71_)[default__.safeIndex(d_3303_v0_, len(d_3388_v71_))]] if ((d_3388_v71_)[default__.safeIndex(d_3303_v0_, len(d_3388_v71_))]) in (d_3482_v125_) else _dafny.SeqWithoutIsStrInference([d_3303_v0_])), d_3303_v0_))
        r2 = self.f3
        return r0, r1, r2

    def m2(self, p0, p1, globalState):
        (self).f3 = p1
        d_3484_v0_: str
        d_3484_v0_ = _dafny.CodePoint('g')
        d_3485_v1_: _dafny.Set
        d_3485_v1_ = _dafny.Set({d_3484_v0_})
        if not (p1) or ((d_3484_v0_) not in (d_3485_v1_)):
            d_3486_v2_: int
            d_3486_v2_ = 149
            d_3486_v2_ = ((948) + ((0) - (d_3486_v2_))) * ((d_3486_v2_) * (d_3486_v2_))
            d_3487_v3_: _dafny.Array
            nw615_ = _dafny.Array(_dafny.Array(None, 0), 26)
            d_3487_v3_ = nw615_
            d_3488_v4_: _dafny.Array
            def lambda157_(d_3489_i0_):
                return self.f3

            init93_ = lambda157_
            nw616_ = _dafny.Array(None, 26)
            for i0_93_ in range(nw616_.length(0)):
                nw616_[i0_93_] = init93_(i0_93_)
            d_3488_v4_ = nw616_
            index560_ = default__.safeIndex(421, (d_3487_v3_).length(0))
            (d_3487_v3_)[index560_] = d_3488_v4_
            index561_ = default__.safeIndex(421, (d_3487_v3_).length(0))
            (d_3487_v3_)[index561_] = (d_3488_v4_ if p1 else d_3488_v4_)
            if self.f3:
                d_3490_v5_: C0
                nw617_ = C0()
                nw617_.ctor__(d_3486_v2_)
                d_3490_v5_ = nw617_
                d_3490_v5_ = d_3490_v5_
                d_3491_v6_: C3
                nw618_ = C3()
                nw618_.ctor__(d_3486_v2_, p0, default__.safeModuloInt(d_3486_v2_, d_3486_v2_))
                d_3491_v6_ = nw618_
                d_3492_v7_: _dafny.Array
                def lambda158_(d_3493_v6_):
                    def lambda159_(d_3494_i1_):
                        return default__.safeModuloInt(d_3494_i1_, (d_3493_v6_).f11)

                    return lambda159_

                init94_ = lambda158_(d_3491_v6_)
                nw619_ = _dafny.Array(None, 18)
                for i0_94_ in range(nw619_.length(0)):
                    nw619_[i0_94_] = init94_(i0_94_)
                d_3492_v7_ = nw619_
                index562_ = default__.safeIndex(517, (d_3492_v7_).length(0))
                (d_3492_v7_)[index562_] = d_3486_v2_
                index563_ = default__.safeIndex(517, (d_3492_v7_).length(0))
                (d_3492_v7_)[index563_] = (d_3486_v2_ if self.f3 else default__.fm1(globalState))
                d_3495_v8_: _dafny.MultiSet
                d_3495_v8_ = _dafny.MultiSet([(d_3487_v3_)[default__.safeIndex(421, (d_3487_v3_).length(0))], (d_3487_v3_)[default__.safeIndex(421, (d_3487_v3_).length(0))]])
                d_3496_v9_: C6
                nw620_ = C6()
                nw620_.ctor__(False, d_3495_v8_, d_3486_v2_)
                d_3496_v9_ = nw620_
                d_3497_v10_: _dafny.Map
                d_3497_v10_ = _dafny.Map({(d_3492_v7_)[default__.safeIndex(517, (d_3492_v7_).length(0))]: d_3496_v9_})
                d_3498_v11_: _dafny.Array
                nw621_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 14)
                d_3498_v11_ = nw621_
                index564_ = default__.safeIndex(357, (d_3498_v11_).length(0))
                (d_3498_v11_)[index564_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "om"))
                d_3499_v12_: _dafny.Set
                d_3499_v12_ = _dafny.Set({False})
                d_3500_v13_: _dafny.Map
                d_3500_v13_ = _dafny.Map({len(d_3499_v12_): d_3486_v2_})
                d_3501_v14_: _dafny.Map
                d_3501_v14_ = _dafny.Map({d_3500_v13_: p1})
                index565_ = default__.safeIndex(517, (d_3492_v7_).length(0))
                index566_ = default__.safeIndex(357, (d_3498_v11_).length(0))
                index567_ = default__.safeIndex(517, (d_3492_v7_).length(0))
                rhs536_ = ((d_3491_v6_).f11) * (((d_3491_v6_).f11 if (d_3496_v9_).f8 else d_3486_v2_))
                rhs537_ = d_3497_v10_
                rhs538_ = p1
                rhs539_ = (d_3491_v6_).f12
                rhs540_ = default__.safeModuloInt((d_3492_v7_)[default__.safeIndex(517, (d_3492_v7_).length(0))], len(d_3501_v14_))
                lhs341_ = d_3492_v7_
                lhs342_ = default__.safeIndex(517, (d_3492_v7_).length(0))
                lhs343_ = self
                lhs344_ = d_3498_v11_
                lhs345_ = default__.safeIndex(357, (d_3498_v11_).length(0))
                lhs346_ = d_3492_v7_
                lhs347_ = default__.safeIndex(517, (d_3492_v7_).length(0))
                lhs341_[lhs342_] = rhs536_
                d_3497_v10_ = rhs537_
                lhs343_.f3 = rhs538_
                lhs344_[lhs345_] = rhs539_
                lhs346_[lhs347_] = rhs540_
                d_3502_v15_: _dafny.Array
                def lambda160_(d_3503_v11_, d_3504_v9_):
                    def lambda161_(d_3505_i2_):
                        return _dafny.Map({(d_3503_v11_)[default__.safeIndex(357, (d_3503_v11_).length(0))]: (d_3504_v9_).f8})

                    return lambda161_

                init95_ = lambda160_(d_3498_v11_, d_3496_v9_)
                nw622_ = _dafny.Array(None, 11)
                for i0_95_ in range(nw622_.length(0)):
                    nw622_[i0_95_] = init95_(i0_95_)
                d_3502_v15_ = nw622_
                d_3506_v16_: _dafny.Map
                d_3506_v16_ = _dafny.Map({p0: not(p1)})
                index568_ = default__.safeIndex(103, (d_3502_v15_).length(0))
                (d_3502_v15_)[index568_] = (d_3506_v16_) | (d_3506_v16_)
                d_3507_v17_: _dafny.Seq
                d_3507_v17_ = _dafny.SeqWithoutIsStrInference([(d_3506_v16_) | (d_3506_v16_)])
                d_3508_v18_: _dafny.Seq
                d_3508_v18_ = _dafny.SeqWithoutIsStrInference([(d_3491_v6_).f11, (d_3491_v6_).f11])
                index569_ = default__.safeIndex(103, (d_3502_v15_).length(0))
                (d_3502_v15_)[index569_] = (d_3507_v17_)[default__.safeIndex(len(d_3508_v18_), len(d_3507_v17_))]
            elif True:
                d_3509_v19_: _dafny.Seq
                d_3509_v19_ = _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([d_3484_v0_])).set(d_3484_v0_, default__.abs(d_3486_v2_))])
                d_3510_v20_: D24
                d_3510_v20_ = D24_DC63((d_3509_v19_)[default__.safeIndex(d_3486_v2_, len(d_3509_v19_))])
                rhs541_ = d_3510_v20_
                rhs542_ = ((d_3486_v2_) * (d_3486_v2_)) + (((D7_DC27(d_3486_v2_, d_3486_v2_, d_3486_v2_)).cf51 if self.f3 else d_3486_v2_))
                rhs543_ = d_3487_v3_
                d_3510_v20_ = rhs541_
                d_3486_v2_ = rhs542_
                d_3487_v3_ = rhs543_
                d_3511_v21_: _dafny.MultiSet
                d_3511_v21_ = _dafny.MultiSet([d_3488_v4_])
                d_3512_v22_: _dafny.Array
                def lambda162_(d_3513_v2_):
                    def lambda163_(d_3514_i3_):
                        return default__.safeModuloInt(d_3514_i3_, d_3513_v2_)

                    return lambda163_

                init96_ = lambda162_(d_3486_v2_)
                nw623_ = _dafny.Array(None, 22)
                for i0_96_ in range(nw623_.length(0)):
                    nw623_[i0_96_] = init96_(i0_96_)
                d_3512_v22_ = nw623_
                d_3515_v23_: _dafny.MultiSet
                d_3515_v23_ = _dafny.MultiSet([d_3512_v22_])
                d_3516_v24_: C6
                nw624_ = C6()
                nw624_.ctor__(p1, d_3511_v21_, ((d_3515_v23_)[d_3512_v22_] if (d_3512_v22_) in (d_3515_v23_) else default__.fm1(globalState)))
                d_3516_v24_ = nw624_
                index570_ = default__.safeIndex(319, (d_3488_v4_).length(0))
                (d_3488_v4_)[index570_] = (d_3516_v24_).f8
                d_3517_v25_: _dafny.Seq
                d_3517_v25_ = _dafny.SeqWithoutIsStrInference([d_3486_v2_])
                d_3518_v26_: _dafny.Seq
                d_3518_v26_ = _dafny.SeqWithoutIsStrInference([(d_3516_v24_).f8, p1, (d_3517_v25_) <= (d_3517_v25_), False])
                index571_ = default__.safeIndex(319, (d_3488_v4_).length(0))
                rhs544_ = (p0) < (((p0) + (p0)).set(default__.safeIndex(d_3486_v2_, len((p0) + (p0))), d_3484_v0_))
                rhs545_ = d_3488_v4_
                rhs546_ = p1
                rhs547_ = (0) - (d_3486_v2_)
                rhs548_ = default__.fm2(d_3484_v0_, default__.safeModuloInt(d_3486_v2_, d_3486_v2_), d_3486_v2_, globalState)
                lhs348_ = self
                lhs349_ = d_3488_v4_
                lhs350_ = default__.safeIndex(319, (d_3488_v4_).length(0))
                lhs348_.f3 = rhs544_
                d_3488_v4_ = rhs545_
                lhs349_[lhs350_] = rhs546_
                d_3486_v2_ = rhs547_
                d_3518_v26_ = rhs548_
                index572_ = default__.safeIndex(319, (d_3488_v4_).length(0))
                (d_3488_v4_)[index572_] = (d_3516_v24_).f8
                d_3519_v27_: _dafny.Set
                d_3519_v27_ = _dafny.Set({not((d_3516_v24_).f8)})
                d_3520_v28_: _dafny.Map
                d_3520_v28_ = _dafny.Map({d_3519_v27_: (len(_dafny.SeqWithoutIsStrInference([d_3484_v0_ for d_3521_i4_ in range(default__.abs(877))]))) - (611)})
                d_3520_v28_ = (d_3520_v28_).set(d_3519_v27_, d_3486_v2_)
            d_3522_v29_: D0
            d_3522_v29_ = D0_DC1(833, (0) - (d_3486_v2_), not(self.f3))
            d_3523_v30_: _dafny.MultiSet
            d_3523_v30_ = _dafny.MultiSet([d_3488_v4_])
            d_3524_v31_: _dafny.Seq
            d_3524_v31_ = _dafny.SeqWithoutIsStrInference([653])
            d_3525_v32_: _dafny.Seq
            d_3525_v32_ = _dafny.SeqWithoutIsStrInference([d_3524_v31_, _dafny.SeqWithoutIsStrInference([d_3486_v2_]), _dafny.SeqWithoutIsStrInference([len(_dafny.Set({d_3486_v2_, d_3486_v2_, 708, (0) - (d_3486_v2_)})) for d_3526_i7_ in range(default__.abs(727))]), d_3524_v31_])
            d_3527_v33_: _dafny.Seq
            d_3527_v33_ = _dafny.SeqWithoutIsStrInference([d_3524_v31_, _dafny.SeqWithoutIsStrInference([d_3486_v2_ for d_3528_i5_ in range(default__.abs(816))]), _dafny.SeqWithoutIsStrInference([d_3486_v2_ for d_3529_i6_ in range(default__.abs(568))]), (d_3525_v32_)[default__.safeIndex(d_3486_v2_, len(d_3525_v32_))], d_3524_v31_])
            d_3530_v34_: C6
            nw625_ = C6()
            nw625_.ctor__((d_3522_v29_).cf3, d_3523_v30_, (len((d_3527_v33_)[default__.safeIndex(len(p0), len(d_3527_v33_))])) + (len(_dafny.Set({p1, True}))))
            d_3530_v34_ = nw625_
            (self).f3 = True
        elif True:
            d_3531_v35_: _dafny.Set
            d_3531_v35_ = _dafny.Set({p1})
            d_3532_v36_: _dafny.Map
            d_3532_v36_ = _dafny.Map({self.f3: d_3531_v35_})
            d_3533_v37_: _dafny.Map
            d_3533_v37_ = _dafny.Map({p1: (d_3532_v36_).set(not(self.f3), _dafny.Set({False}))})
            d_3534_v38_: _dafny.Seq
            d_3534_v38_ = _dafny.SeqWithoutIsStrInference([d_3531_v35_, d_3531_v35_])
            d_3535_v39_: int
            d_3535_v39_ = 17
            d_3533_v37_ = (d_3533_v37_).set((p0) < (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "knydbf"))), ((d_3532_v36_).set(p1, (d_3534_v38_)[default__.safeIndex((0) - (d_3535_v39_), len(d_3534_v38_))])).set(p1, _dafny.Set({not(True)})))
            d_3535_v39_ = d_3535_v39_
            (self).f3 = p1
            (self).f3 = False
            rhs549_ = self.f3
            rhs550_ = (513) * (d_3535_v39_)
            rhs551_ = default__.safeModuloInt(default__.safeDivisionInt(d_3535_v39_, d_3535_v39_), d_3535_v39_)
            lhs351_ = self
            lhs351_.f3 = rhs549_
            d_3535_v39_ = rhs550_
            d_3535_v39_ = rhs551_
        d_3536_v41_: int
        d_3536_v41_ = 765
        d_3537_v42_: _dafny.Array
        nw626_ = _dafny.Array(int(0), 5)
        d_3537_v42_ = nw626_
        d_3538_v43_: _dafny.Map
        d_3538_v43_ = _dafny.Map({d_3536_v41_: d_3537_v42_})
        d_3539_v44_: int
        d_3540_v45_: int
        d_3541_v46_: _dafny.Array
        d_3542_v47_: bool
        out79_: int
        out80_: int
        out81_: _dafny.Array
        out82_: bool
        def iife313_():
            coll109_ = _dafny.Map()
            compr_109_: int
            for compr_109_ in _dafny.IntegerRange(-640, 87):
                d_3543_v40_: int = compr_109_
                if ((-640) <= (d_3543_v40_)) and ((d_3543_v40_) < (87)):
                    coll109_[(d_3543_v40_) + (len(p0))] = p1
            return _dafny.Map(coll109_)
        out79_, out80_, out81_, out82_ = default__.m0(iife313_()
        , _dafny.SeqWithoutIsStrInference([189 for d_3544_i8_ in range(default__.abs(-814))]), d_3538_v43_, globalState)
        d_3539_v44_ = out79_
        d_3540_v45_ = out80_
        d_3541_v46_ = out81_
        d_3542_v47_ = out82_
        if (True) and (d_3542_v47_):
            d_3545_v48_: _dafny.Seq
            d_3545_v48_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v"))
            d_3545_v48_ = p0
            d_3546_v49_: _dafny.Array
            nw627_ = _dafny.Array(False, 4)
            d_3546_v49_ = nw627_
            index573_ = default__.safeIndex(362, (d_3546_v49_).length(0))
            (d_3546_v49_)[index573_] = p1
            d_3547_v50_: D11
            d_3547_v50_ = D11_DC35(d_3539_v44_, p1, self.f3)
            d_3548_v51_: _dafny.MultiSet
            d_3548_v51_ = _dafny.MultiSet([not(p1), True])
            index574_ = default__.safeIndex(362, (d_3546_v49_).length(0))
            rhs552_ = d_3539_v44_
            rhs553_ = d_3540_v45_
            rhs554_ = not((d_3548_v51_).ispropersubset(d_3548_v51_))
            rhs555_ = (D11_DC35(d_3536_v41_, d_3542_v47_, self.f3) if (d_3542_v47_) or ((d_3547_v50_).cf62) else d_3547_v50_)
            rhs556_ = not(((d_3536_v41_) - (d_3539_v44_)) == (d_3539_v44_))
            lhs352_ = d_3546_v49_
            lhs353_ = default__.safeIndex(362, (d_3546_v49_).length(0))
            d_3539_v44_ = rhs552_
            d_3536_v41_ = rhs553_
            lhs352_[lhs353_] = rhs554_
            d_3547_v50_ = rhs555_
            d_3542_v47_ = rhs556_
            d_3549_v52_: _dafny.Seq
            d_3549_v52_ = _dafny.SeqWithoutIsStrInference([d_3542_v47_])
            d_3550_v54_: _dafny.Map
            def iife314_():
                coll110_ = _dafny.Map()
                compr_110_: int
                for compr_110_ in _dafny.IntegerRange(206, 986):
                    d_3551_v53_: int = compr_110_
                    if ((206) <= (d_3551_v53_)) and ((d_3551_v53_) < (986)):
                        coll110_[(d_3551_v53_) + ((D0_DC1(d_3536_v41_, len(d_3549_v52_), True)).cf2)] = d_3542_v47_
                return _dafny.Map(coll110_)
            d_3550_v54_ = _dafny.Map({(default__.fm37(d_3542_v47_, p1, (0) - (len(d_3549_v52_)), globalState)) | (iife314_()
            ): (0) - ((0) - (d_3539_v44_))})
            d_3550_v54_ = d_3550_v54_
            d_3540_v45_ = d_3540_v45_
            if not(False):
                d_3552_v55_: C0
                nw628_ = C0()
                nw628_.ctor__(d_3536_v41_)
                d_3552_v55_ = nw628_
                nw629_ = C0()
                nw629_.ctor__(default__.safeDivisionInt(d_3536_v41_, d_3539_v44_))
                d_3552_v55_ = nw629_
                d_3553_v56_: C0
                nw630_ = C0()
                nw630_.ctor__((d_3540_v45_ if default__.fm0(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "grcmeyna"))), len(d_3545_v48_), globalState) else d_3552_v55_.f18))
                d_3553_v56_ = nw630_
                d_3554_v57_: _dafny.MultiSet
                d_3554_v57_ = _dafny.MultiSet([560, ((default__.fm25(globalState)) - (_dafny.MultiSet([p1]))).cardinality, d_3536_v41_])
                d_3539_v44_ = ((d_3554_v57_)[d_3539_v44_] if (d_3539_v44_) in (d_3554_v57_) else (d_3548_v51_).cardinality)
                d_3555_v58_: _dafny.Seq
                d_3555_v58_ = _dafny.SeqWithoutIsStrInference([(d_3545_v48_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vmskjctci")))])
                d_3556_v59_: _dafny.Map
                d_3556_v59_ = _dafny.Map({d_3484_v0_: (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gwwwqmca"))).set(default__.safeIndex(d_3536_v41_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gwwwqmca")))), (p0)[default__.safeIndex(d_3540_v45_, len(p0))])})
                d_3555_v58_ = (default__.fm54((698) + (d_3552_v55_.f18), d_3484_v0_, d_3556_v59_, globalState)).set(default__.safeIndex(d_3536_v41_, len(default__.fm54((698) + (d_3552_v55_.f18), d_3484_v0_, d_3556_v59_, globalState))), (d_3545_v48_ if (d_3549_v52_)[default__.safeIndex(d_3539_v44_, len(d_3549_v52_))] else p0))
                d_3557_v60_: _dafny.Array
                nw631_ = _dafny.Array(_dafny.Map({}), 21)
                d_3557_v60_ = nw631_
                d_3558_v61_: _dafny.Array
                nw632_ = _dafny.Array(None, 18)
                d_3558_v61_ = nw632_
                d_3559_v62_: _dafny.Map
                d_3559_v62_ = _dafny.Map({d_3540_v45_: d_3558_v61_})
                index575_ = default__.safeIndex(39, (d_3557_v60_).length(0))
                (d_3557_v60_)[index575_] = d_3559_v62_
                d_3560_v63_: _dafny.Map
                d_3560_v63_ = d_3559_v62_
                index576_ = default__.safeIndex(39, (d_3557_v60_).length(0))
                (d_3557_v60_)[index576_] = (d_3559_v62_ if d_3542_v47_ else d_3560_v63_)
            elif True:
                index577_ = default__.safeIndex(362, (d_3546_v49_).length(0))
                (d_3546_v49_)[index577_] = d_3542_v47_
                d_3561_v64_: _dafny.Map
                d_3561_v64_ = _dafny.Map({d_3540_v45_: d_3539_v44_})
                def iife315_():
                    coll111_ = _dafny.Map()
                    compr_111_: int
                    for compr_111_ in _dafny.IntegerRange(562, 728):
                        d_3562_v65_: int = compr_111_
                        if ((562) <= (d_3562_v65_)) and ((d_3562_v65_) < (728)):
                            coll111_[default__.safeDivisionInt(d_3562_v65_, (0) - (len(_dafny.Map({self.f3: d_3539_v44_}))))] = d_3540_v45_
                    return _dafny.Map(coll111_)
                d_3561_v64_ = iife315_()
                
                (self).f3 = ((d_3546_v49_)[default__.safeIndex(362, (d_3546_v49_).length(0))]) or ((d_3546_v49_)[default__.safeIndex(362, (d_3546_v49_).length(0))])
                d_3563_v66_: _dafny.Map
                d_3563_v66_ = _dafny.Map({(d_3546_v49_)[default__.safeIndex(362, (d_3546_v49_).length(0))]: d_3542_v47_})
                d_3564_v67_: _dafny.Set
                d_3564_v67_ = _dafny.Set({d_3563_v66_, d_3563_v66_, d_3563_v66_, d_3563_v66_, _dafny.Map({p1: (d_3546_v49_)[default__.safeIndex(362, (d_3546_v49_).length(0))]})})
                d_3565_v68_: _dafny.Set
                d_3565_v68_ = _dafny.Set({(d_3564_v67_).issubset(_dafny.Set({d_3563_v66_})), (d_3484_v0_) in (p0)})
                d_3565_v68_ = d_3565_v68_
                d_3566_v69_: _dafny.Set
                d_3566_v69_ = _dafny.Set({d_3539_v44_})
                d_3567_v70_: _dafny.Map
                d_3567_v70_ = _dafny.Map({(d_3539_v44_) * (648): default__.fm55(len(d_3566_v69_), d_3536_v41_, d_3536_v41_, globalState)})
                d_3568_v71_: _dafny.Map
                d_3568_v71_ = _dafny.Map({p1: d_3484_v0_})
                d_3567_v70_ = (d_3567_v70_).set((d_3539_v44_) * (len(default__.fm56(((d_3568_v71_)[p1] if (p1) in (d_3568_v71_) else d_3484_v0_), globalState))), _dafny.Map({d_3536_v41_: d_3484_v0_}))
        elif True:
            d_3569_v72_: _dafny.Seq
            d_3569_v72_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yrwuvgf"))
            d_3570_v73_: D8
            d_3570_v73_ = D8_DC29()
            rhs557_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "otxoiimtp"))
            rhs558_ = d_3570_v73_
            d_3569_v72_ = rhs557_
            d_3570_v73_ = rhs558_
            d_3571_v74_: _dafny.Map
            d_3571_v74_ = _dafny.Map({(d_3569_v72_) <= (p0): d_3542_v47_})
            d_3572_v75_: _dafny.Array
            nw633_ = _dafny.Array(False, 10)
            d_3572_v75_ = nw633_
            d_3573_v76_: _dafny.Set
            d_3573_v76_ = _dafny.Set({d_3572_v75_})
            index578_ = default__.safeIndex(606, (d_3541_v46_).length(0))
            (d_3541_v46_)[index578_] = (len(d_3573_v76_)) - (d_3539_v44_)
            index579_ = default__.safeIndex(606, (d_3541_v46_).length(0))
            rhs559_ = d_3571_v74_
            rhs560_ = default__.fm1(globalState)
            rhs561_ = d_3484_v0_
            rhs562_ = ((0) - (default__.safeDivisionInt(d_3540_v45_, d_3540_v45_))) > (d_3539_v44_)
            lhs354_ = d_3541_v46_
            lhs355_ = default__.safeIndex(606, (d_3541_v46_).length(0))
            d_3571_v74_ = rhs559_
            lhs354_[lhs355_] = rhs560_
            d_3484_v0_ = rhs561_
            d_3542_v47_ = rhs562_
            d_3574_v77_: _dafny.Array
            def lambda164_(d_3575_v46_, d_3576_v41_, d_3577_v45_, d_3578_p1_):
                def lambda165_(d_3579_i9_):
                    return D19_DC52((d_3575_v46_)[default__.safeIndex(606, (d_3575_v46_).length(0))], d_3576_v41_, (D19_DC52(d_3576_v41_, d_3576_v41_, d_3577_v45_, d_3578_p1_)).cf94, self.f3)

                return lambda165_

            init97_ = lambda164_(d_3541_v46_, d_3536_v41_, d_3540_v45_, p1)
            nw634_ = _dafny.Array(None, 19)
            for i0_97_ in range(nw634_.length(0)):
                nw634_[i0_97_] = init97_(i0_97_)
            d_3574_v77_ = nw634_
            d_3580_v78_: C1
            nw635_ = C1()
            nw635_.ctor__((d_3541_v46_)[default__.safeIndex(606, (d_3541_v46_).length(0))], 387, d_3539_v44_, True)
            d_3580_v78_ = nw635_
            d_3581_v79_: _dafny.Map
            d_3581_v79_ = _dafny.Map({d_3580_v78_: d_3542_v47_})
            d_3582_v80_: D19
            d_3582_v80_ = D19_DC52(633, default__.fm1(globalState), len(d_3581_v79_), not(self.f3))
            index580_ = default__.safeIndex(540, (d_3574_v77_).length(0))
            (d_3574_v77_)[index580_] = d_3582_v80_
            index581_ = default__.safeIndex(540, (d_3574_v77_).length(0))
            (d_3574_v77_)[index581_] = D19_DC52(len((p0).set(default__.safeIndex(len(_dafny.Map({d_3542_v47_: d_3540_v45_})), len(p0)), d_3484_v0_)), len(_dafny.SeqWithoutIsStrInference([d_3540_v45_])), default__.safeModuloInt(d_3580_v78_.f17, d_3536_v41_), p1)
            d_3583_v81_: D0
            d_3583_v81_ = D0_DC1(d_3580_v78_.f17, d_3536_v41_, d_3542_v47_)
            d_3584_v82_: D0
            d_3584_v82_ = D0_DC2(d_3583_v81_)
            d_3585_v83_: D0
            d_3585_v83_ = D0_DC2(d_3584_v82_)
            d_3586_v84_: _dafny.Seq
            d_3586_v84_ = _dafny.SeqWithoutIsStrInference([d_3585_v83_, d_3585_v83_])
            d_3587_v85_: _dafny.Map
            d_3587_v85_ = _dafny.Map({d_3580_v78_.f17: (0) - (d_3536_v41_)})
            rhs563_ = d_3542_v47_
            rhs564_ = len(_dafny.Map({(((d_3587_v85_)[d_3539_v44_] if (d_3539_v44_) in (d_3587_v85_) else default__.fm1(globalState))) <= (d_3580_v78_.f17): (d_3541_v46_)[default__.safeIndex(606, (d_3541_v46_).length(0))]}))
            rhs565_ = _dafny.SeqWithoutIsStrInference([d_3585_v83_ for d_3588_i10_ in range(default__.abs(976))])
            lhs356_ = self
            lhs356_.f3 = rhs563_
            d_3536_v41_ = rhs564_
            d_3586_v84_ = rhs565_
            d_3589_v86_: _dafny.Map
            d_3589_v86_ = _dafny.Map({(d_3541_v46_)[default__.safeIndex(606, (d_3541_v46_).length(0))]: d_3542_v47_})
            d_3589_v86_ = (d_3589_v86_).set(d_3539_v44_, d_3542_v47_)
        d_3590_v87_: _dafny.Map
        d_3590_v87_ = _dafny.Map({d_3542_v47_: (d_3536_v41_) - (d_3536_v41_)})
        d_3591_v88_: _dafny.Seq
        d_3591_v88_ = _dafny.SeqWithoutIsStrInference([d_3539_v44_])
        d_3590_v87_ = (d_3590_v87_).set(((_dafny.MultiSet(d_3591_v88_)).cardinality) > (d_3539_v44_), d_3540_v45_)
        d_3592_v89_: _dafny.Set
        d_3592_v89_ = _dafny.Set({(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r')])).set(default__.safeIndex(d_3540_v45_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r')]))), d_3484_v0_)})
        d_3593_i11_: int
        d_3593_i11_ = 0
        with _dafny.label("13"):
            while ((_dafny.Set({p0})).isdisjoint(d_3592_v89_)) not in ((D1_DC3(_dafny.Map({d_3542_v47_: d_3539_v44_}))).cf5):
                with _dafny.c_label("13"):
                    if (d_3593_i11_) >= (100):
                        raise _dafny.Break("13")
                    d_3593_i11_ = (d_3593_i11_) + (1)
                    d_3594_v90_: _dafny.MultiSet
                    d_3594_v90_ = _dafny.MultiSet([d_3542_v47_, self.f3, self.f3])
                    d_3595_v91_: _dafny.Array
                    def lambda166_(d_3596_v47_):
                        def lambda167_(d_3597_i12_):
                            return d_3596_v47_

                        return lambda167_

                    init98_ = lambda166_(d_3542_v47_)
                    nw636_ = _dafny.Array(None, 10)
                    for i0_98_ in range(nw636_.length(0)):
                        nw636_[i0_98_] = init98_(i0_98_)
                    d_3595_v91_ = nw636_
                    rhs566_ = (default__.fm1(globalState)) * ((d_3539_v44_) * (d_3540_v45_))
                    rhs567_ = d_3594_v90_
                    rhs568_ = default__.fm1(globalState)
                    rhs569_ = d_3595_v91_
                    rhs570_ = len(((d_3485_v1_).intersection(d_3485_v1_)) | (default__.fm57(p1, self.f3, 304, globalState)))
                    d_3536_v41_ = rhs566_
                    d_3594_v90_ = rhs567_
                    d_3539_v44_ = rhs568_
                    d_3595_v91_ = rhs569_
                    d_3540_v45_ = rhs570_
                    d_3598_v92_: _dafny.MultiSet
                    d_3598_v92_ = _dafny.MultiSet([d_3595_v91_, d_3595_v91_])
                    d_3599_v93_: C6
                    nw637_ = C6()
                    nw637_.ctor__(True, d_3598_v92_, len((d_3591_v88_) + (d_3591_v88_)))
                    d_3599_v93_ = nw637_
                    d_3600_v94_: _dafny.Map
                    d_3600_v94_ = _dafny.Map({p0: 648})
                    d_3600_v94_ = (d_3600_v94_).set((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o"))) + (p0), ((0) - (d_3539_v44_)) * (len(p0)))
                    d_3590_v87_ = (d_3590_v87_).set(self.f3, len((p0 if (d_3599_v93_).f8 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qncabrib")))))
                    pass
            pass

