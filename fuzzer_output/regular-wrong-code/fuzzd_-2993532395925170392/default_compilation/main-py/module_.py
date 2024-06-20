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
    def fm0(p0, p1, p2, p3, globalState):
        return default__.safeDivisionInt(639, 701)

    @staticmethod
    def fm1(p0, p1, p2, globalState):
        return ((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tpfhnsqey")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "txfv")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_0_i1_ in range(default__.abs(732))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "utuau")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ay"))])).intersection(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([(D7_DC15(_dafny.CodePoint('a'), False)).cf28 for d_1_i2_ in range(default__.abs(109))])]))).issubset((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cjuvb"))])) | (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_2_i0_ in range(default__.abs(309))])])))

    @staticmethod
    def fm2(p0, p1, p2, p3, globalState):
        source0_ = D26_DC62(_dafny.SeqWithoutIsStrInference([False]))
        if source0_.is_DC61:
            d_3___mcc_h0_ = source0_.cf97
            d_4_cf97_ = d_3___mcc_h0_
            return (_dafny.Map({d_4_cf97_: d_4_cf97_})) | (_dafny.Map({d_4_cf97_: d_4_cf97_}))
        elif source0_.is_DC62:
            d_5___mcc_h1_ = source0_.cf98
            d_6_cf98_ = d_5___mcc_h1_
            return _dafny.Map({True: True})
        elif source0_.is_DC63:
            d_7___mcc_h2_ = source0_.cf99
            d_8___mcc_h3_ = source0_.cf100
            d_9___mcc_h4_ = source0_.cf101
            d_10___mcc_h5_ = source0_.cf102
            d_11___mcc_h6_ = source0_.cf103
            d_12_cf103_ = d_11___mcc_h6_
            d_13_cf102_ = d_10___mcc_h5_
            d_14_cf101_ = d_9___mcc_h4_
            d_15_cf100_ = d_8___mcc_h3_
            d_16_cf99_ = d_7___mcc_h2_
            return (_dafny.Map({d_12_cf103_: d_12_cf103_})) | (_dafny.Map({d_13_cf102_: False}))
        elif True:
            d_17___mcc_h7_ = source0_.cf96
            d_18_cf96_ = d_17___mcc_h7_
            return _dafny.Map({not(True): True})

    @staticmethod
    def fm3(p0, p1, p2, p3, globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vsc"))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cov")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vx"))]))) + ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gk")) for d_19_i0_ in range(default__.abs(542))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ywtrdjrjo"))])))

    @staticmethod
    def fm9(p0, p1, p2, p3, globalState):
        source1_ = _dafny.Map({False: D3_DC8(False, -843, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jlenwnrbf"))))})
        d_20___mcc_h0_ = source1_
        d_21_cf53_ = d_20___mcc_h0_
        return _dafny.SeqWithoutIsStrInference([False, not(True), True, not(not(True)), True])

    @staticmethod
    def fm10(p0, globalState):
        return _dafny.SeqWithoutIsStrInference([(0) - (default__.safeModuloInt(93, -950))])

    @staticmethod
    def fm11(globalState):
        return _dafny.CodePoint('l')

    @staticmethod
    def fm14(globalState):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: D1
            for compr_0_ in (_dafny.Map({D1_DC1(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eftqndnpe"))): True})).keys.Elements:
                d_22_v0_: D1 = compr_0_
                if (d_22_v0_) in (_dafny.Map({D1_DC1(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eftqndnpe"))): True})):
                    coll0_[d_22_v0_] = -671
            return _dafny.Map(coll0_)
        return _dafny.Map({(iife0_()
        ) | (_dafny.Map({D1_DC1(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yejysvd"))): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iywaymc")))})): (False if True else not(not(True)))})

    @staticmethod
    def fm16(p0, p1, p2, p3, globalState):
        source2_ = D11_DC26(False, False, 811, True)
        if source2_.is_DC26:
            d_23___mcc_h0_ = source2_.cf47
            d_24___mcc_h1_ = source2_.cf48
            d_25___mcc_h2_ = source2_.cf49
            d_26___mcc_h3_ = source2_.cf50
            d_27_cf50_ = d_26___mcc_h3_
            d_28_cf49_ = d_25___mcc_h2_
            d_29_cf48_ = d_24___mcc_h1_
            d_30_cf47_ = d_23___mcc_h0_
            return (_dafny.MultiSet([d_30_cf47_])).intersection(_dafny.MultiSet([True, d_27_cf50_]))
        elif source2_.is_DC27:
            return _dafny.MultiSet([not(True), False])
        elif source2_.is_DC28:
            return (_dafny.MultiSet([True])) - (_dafny.MultiSet([True]))
        elif True:
            d_31___mcc_h4_ = source2_.cf46
            d_32_cf46_ = d_31___mcc_h4_
            return (_dafny.MultiSet([(d_32_cf46_).f32])) - (_dafny.MultiSet([(d_32_cf46_).f32]))

    @staticmethod
    def fm20(globalState):
        return ((_dafny.MultiSet([_dafny.CodePoint('p'), _dafny.CodePoint('g')])).intersection(_dafny.MultiSet([_dafny.CodePoint('p')]))) - (_dafny.MultiSet([_dafny.CodePoint('h'), _dafny.CodePoint('e')]))

    @staticmethod
    def fm21(p0, p1, p2, p3, globalState):
        return _dafny.CodePoint('n')

    @staticmethod
    def fm26(p0, p1, p2, globalState):
        return _dafny.SeqWithoutIsStrInference([(_dafny.CodePoint('q') if not(False) else _dafny.CodePoint('l')) for d_33_i0_ in range(default__.abs(491))])

    @staticmethod
    def fm27(p0, globalState):
        return ((_dafny.MultiSet([not(not(True))])).intersection(_dafny.MultiSet([False]))).intersection(_dafny.MultiSet([(D7_DC16(False, len(_dafny.Map({False: (D7_DC15(_dafny.CodePoint('a'), False)).cf28})), False, False, False)).cf33]))

    @staticmethod
    def fm28(globalState):
        return ((_dafny.SeqWithoutIsStrInference([True]) if not(True) else _dafny.SeqWithoutIsStrInference([False, True, True]))) + (_dafny.SeqWithoutIsStrInference([False, not((D11_DC26(not(True), not(True), (_dafny.MultiSet([(_dafny.MultiSet([False])).cardinality])).cardinality, True)).cf47), False, False]))

    @staticmethod
    def fm29(globalState):
        source3_ = D19_DC41(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_34_i0_ in range(default__.abs(169))])]))
        if source3_.is_DC42:
            d_35___mcc_h0_ = source3_.cf63
            d_36___mcc_h1_ = source3_.cf64
            d_37_cf64_ = d_36___mcc_h1_
            d_38_cf63_ = d_35___mcc_h0_
            return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "krsbrr"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_39_i1_ in range(default__.abs(73))]))
        elif source3_.is_DC43:
            d_40___mcc_h2_ = source3_.cf65
            d_41___mcc_h3_ = source3_.cf66
            d_42___mcc_h4_ = source3_.cf67
            d_43___mcc_h5_ = source3_.cf68
            d_44_cf68_ = d_43___mcc_h5_
            d_45_cf67_ = d_42___mcc_h4_
            d_46_cf66_ = d_41___mcc_h3_
            d_47_cf65_ = d_40___mcc_h2_
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dgeq"))
        elif True:
            d_48___mcc_h6_ = source3_.cf62
            d_49_cf62_ = d_48___mcc_h6_
            return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wkuahmwrp"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_50_i2_ in range(default__.abs(985))]))

    @staticmethod
    def fm30(p0, p1, p2, p3, globalState):
        def iife1_():
            coll1_ = _dafny.Map()
            compr_1_: int
            for compr_1_ in (_dafny.SeqWithoutIsStrInference([-720, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_51_i0_ in range(default__.abs(939))]))])).Elements:
                d_52_v0_: int = compr_1_
                if (d_52_v0_) in (_dafny.SeqWithoutIsStrInference([-720, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_51_i0_ in range(default__.abs(939))]))])):
                    coll1_[(d_52_v0_) * (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_53_i1_ in range(default__.abs(-265))])))] = not(False)
            return _dafny.Map(coll1_)
        return _dafny.MultiSet([(_dafny.Map({-425: False})) | (_dafny.Map({(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([not(True), False, True, True, True]))).cardinality: True})), (iife1_()
        )])

    @staticmethod
    def fm31(globalState):
        if (True) == (False):
            return _dafny.Map({False: _dafny.CodePoint('v')})
        elif True:
            return _dafny.Map({False: _dafny.CodePoint('y')})

    @staticmethod
    def fm32(globalState):
        def iife2_():
            coll2_ = _dafny.Map()
            compr_2_: int
            for compr_2_ in _dafny.IntegerRange(-576, 677):
                d_54_v0_: int = compr_2_
                if ((-576) <= (d_54_v0_)) and ((d_54_v0_) < (677)):
                    coll2_[(d_54_v0_) + (((_dafny.MultiSet([False, True]))).cardinality)] = len(_dafny.Set({True, True}))
            return _dafny.Map(coll2_)
        return _dafny.Map({True: default__.safeDivisionInt(len(iife2_()
        ), len(_dafny.Set({False})))})

    @staticmethod
    def fm33(p0, p1, p2, p3, globalState):
        return ((_dafny.SeqWithoutIsStrInference([not(True), False])) + (_dafny.SeqWithoutIsStrInference([True, not(False)]))) + ((_dafny.SeqWithoutIsStrInference([False, True])) + (_dafny.SeqWithoutIsStrInference([False, not(True), True, False, not(False)])))

    @staticmethod
    def fm34(p0, globalState):
        return _dafny.MultiSet([_dafny.CodePoint('u')])

    @staticmethod
    def fm35(p0, p1, globalState):
        if False:
            def iife3_():
                coll3_ = _dafny.Map()
                compr_3_: int
                for compr_3_ in (_dafny.MultiSet([601])).Elements:
                    d_55_v0_: int = compr_3_
                    if (d_55_v0_) in (_dafny.MultiSet([601])):
                        coll3_[default__.safeModuloInt(d_55_v0_, 214)] = 289
                return _dafny.Map(coll3_)
            def iife4_():
                coll4_ = _dafny.Map()
                compr_4_: int
                for compr_4_ in _dafny.IntegerRange(-765, 482):
                    d_56_v1_: int = compr_4_
                    if ((-765) <= (d_56_v1_)) and ((d_56_v1_) < (482)):
                        coll4_[(d_56_v1_) * (len(_dafny.SeqWithoutIsStrInference([True])))] = -619
                return _dafny.Map(coll4_)
            return (iife3_()
            ) | (iife4_()
            )
        elif True:
            def iife5_():
                coll5_ = _dafny.Map()
                compr_5_: int
                for compr_5_ in _dafny.IntegerRange(-548, 783):
                    d_57_v2_: int = compr_5_
                    if ((-548) <= (d_57_v2_)) and ((d_57_v2_) < (783)):
                        coll5_[(d_57_v2_) + ((_dafny.MultiSet([-449, -325])).cardinality)] = 952
                return _dafny.Map(coll5_)
            return (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dshb"))): -196})) | (iife5_()
            )

    @staticmethod
    def fm36(p0, p1, p2, globalState):
        return _dafny.CodePoint('n')

    @staticmethod
    def fm38(p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_58_i0_ in range(default__.abs(39))])) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vbu"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nlkkuwfva"))))

    @staticmethod
    def fm41(p0, p1, p2, globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cxrxnd"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "irrwukcqv")))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "snynaanoq"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_59_i0_ in range(default__.abs(287))])))

    @staticmethod
    def fm42(globalState):
        return ((_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: _dafny.Set({(0) - (len(_dafny.Map({114: len(_dafny.SeqWithoutIsStrInference([False, False]))}))), 729, len(_dafny.Map({964: 701})), 985, len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([-812]))]))})})) for d_60_i0_ in range(default__.abs(928))]))).cardinality])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "skeep"))), 405]))) + (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({_dafny.CodePoint('n'): False})), -156]))

    @staticmethod
    def fm43(globalState):
        return (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: 456}))])]))).intersection((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True])]))])])) - (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uaj")))]), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "igtvf")))]), _dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jk"))))]), _dafny.SeqWithoutIsStrInference([925 for d_61_i0_ in range(default__.abs(89))])])))

    @staticmethod
    def fm44(p0, p1, globalState):
        return _dafny.Set({False, False, not((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vobys")) for d_62_i0_ in range(default__.abs(253))])) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gxrprxa")) for d_63_i1_ in range(default__.abs(374))])))})

    @staticmethod
    def fm45(p0, globalState):
        return ((_dafny.Map({False: len(_dafny.Map({not(False): False}))})) | (_dafny.Map({False: 118}))) | ((_dafny.Map({False: len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lloe"))): len(_dafny.SeqWithoutIsStrInference([16, (0) - (-725)]))}))})) | (_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference([True]))})))

    @staticmethod
    def fm46(globalState):
        return D2_DC5((_dafny.CodePoint('e') if False else _dafny.CodePoint('b')), not((428) <= (916)), -881, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rotjhxp"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_64_i0_ in range(default__.abs(959))]))))

    @staticmethod
    def fm47(globalState):
        return _dafny.CodePoint('y')

    @staticmethod
    def fm48(globalState):
        if (False) or (True):
            return ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.Set({len(_dafny.Set({878})), len(_dafny.SeqWithoutIsStrInference([not(False)]))}))])])))
        elif True:
            return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.Set({True, True, True, True}))]) for d_65_i0_ in range(default__.abs(31))])

    @staticmethod
    def fm49(p0, p1, p2, p3, globalState):
        return ((_dafny.SeqWithoutIsStrInference([False])) + (_dafny.SeqWithoutIsStrInference([False]))) + (_dafny.SeqWithoutIsStrInference([False]))

    @staticmethod
    def fm50(p0, p1, p2, globalState):
        def iife6_():
            coll6_ = _dafny.Set()
            compr_6_: int
            for compr_6_ in _dafny.IntegerRange(267, 77):
                d_66_v0_: int = compr_6_
                if ((267) <= (d_66_v0_)) and ((d_66_v0_) < (77)):
                    coll6_ = coll6_.union(_dafny.Set([default__.safeModuloInt(d_66_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wowcjj"))))]))
            return _dafny.Set(coll6_)
        return (_dafny.Map({len(_dafny.Map({366: not(True)})): 406})) | ((_dafny.Map({len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({True, True, not(False)}))])): len(_dafny.Map({len(iife6_()
        ): True}))}) if not(False) else _dafny.Map({len(_dafny.Map({_dafny.SeqWithoutIsStrInference([False, False]): 858})): 31})))

    @staticmethod
    def fm51(p0, globalState):
        def lambda0_(source4_):
            if source4_.is_DC51:
                d_67___mcc_h0_ = source4_.cf79
                d_68___mcc_h1_ = source4_.cf80
                d_69_cf80_ = d_68___mcc_h1_
                d_70_cf79_ = d_67___mcc_h0_
                return (_dafny.SeqWithoutIsStrInference([True, d_69_cf80_])) + (_dafny.SeqWithoutIsStrInference([d_69_cf80_, d_70_cf79_, d_69_cf80_, d_70_cf79_, False]))
            elif source4_.is_DC52:
                d_71___mcc_h2_ = source4_.cf81
                d_72___mcc_h3_ = source4_.cf82
                d_73___mcc_h4_ = source4_.cf83
                d_74_cf83_ = d_73___mcc_h4_
                d_75_cf82_ = d_72___mcc_h3_
                d_76_cf81_ = d_71___mcc_h2_
                return _dafny.SeqWithoutIsStrInference([d_74_cf83_, d_75_cf82_, d_75_cf82_, d_74_cf83_])
            elif source4_.is_DC50:
                d_77___mcc_h5_ = source4_.cf78
                d_78_cf78_ = d_77___mcc_h5_
                return (_dafny.SeqWithoutIsStrInference([not(False), True])) + (_dafny.SeqWithoutIsStrInference([False, not(not(True))]))
            elif True:
                d_79___mcc_h6_ = source4_.cf84
                d_80_cf84_ = d_79___mcc_h6_
                return (_dafny.SeqWithoutIsStrInference([False])) + (_dafny.SeqWithoutIsStrInference([True, True, True, False, False]))

        return _dafny.MultiSet(lambda0_(D22_DC52(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fn"))), True, True)))

    @staticmethod
    def fm52(p0, globalState):
        return (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_81_i1_ in range(default__.abs(798))])): True})) for d_82_i0_ in range(default__.abs(79))]))) - (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([False])).cardinality for d_83_i2_ in range(default__.abs(-328))])))

    @staticmethod
    def fm53(p0, p1, globalState):
        return (_dafny.Set({610, 586, len(_dafny.SeqWithoutIsStrInference([False, True]))})) | (_dafny.Set({80}))

    @staticmethod
    def fm54(p0, globalState):
        return (_dafny.SeqWithoutIsStrInference([D7_DC15(_dafny.CodePoint('r'), True), D7_DC15(_dafny.CodePoint('f'), True), D7_DC15(_dafny.CodePoint('r'), False)])) + (_dafny.SeqWithoutIsStrInference([D7_DC15(_dafny.CodePoint('m'), not(False))]))

    @staticmethod
    def fm55(p0, p1, p2, p3, globalState):
        source5_ = D23_DC54(_dafny.Set({False}))
        if source5_.is_DC55:
            d_84___mcc_h0_ = source5_.cf86
            d_85___mcc_h1_ = source5_.cf87
            d_86___mcc_h2_ = source5_.cf88
            d_87_cf88_ = d_86___mcc_h2_
            d_88_cf87_ = d_85___mcc_h1_
            d_89_cf86_ = d_84___mcc_h0_
            return (_dafny.Map({d_87_cf88_: _dafny.SeqWithoutIsStrInference([D0_DC0(True, d_89_cf86_, d_89_cf86_) for d_90_i0_ in range(default__.abs(465))])})) | (_dafny.Map({d_87_cf88_: _dafny.SeqWithoutIsStrInference([D0_DC0(d_88_cf87_, d_89_cf86_, len(_dafny.Set({d_88_cf87_}))), D0_DC0(d_88_cf87_, len(_dafny.SeqWithoutIsStrInference([d_89_cf86_ for d_91_i1_ in range(default__.abs(-693))])), d_89_cf86_)])}))
        elif True:
            d_92___mcc_h3_ = source5_.cf85
            d_93_cf85_ = d_92___mcc_h3_
            return (_dafny.Map({not(True): _dafny.SeqWithoutIsStrInference([D0_DC0(True, 774, len(_dafny.Set({286, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_94_i2_ in range(default__.abs(904))]))})))])})) | (_dafny.Map({True: _dafny.SeqWithoutIsStrInference([D0_DC0(False, (0) - ((0) - (len(d_93_cf85_))), 903)])}))

    @staticmethod
    def fm56(globalState):
        def iife7_():
            def iife9_():
                coll9_ = _dafny.Map()
                compr_9_: _dafny.Seq
                for compr_9_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lll"))) for d_95_i0_ in range(default__.abs(148))]), _dafny.SeqWithoutIsStrInference([17]), _dafny.SeqWithoutIsStrInference([len(_dafny.Set({(0) - (-719), -391, -329})) for d_96_i1_ in range(default__.abs(200))]), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([620, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "caiutom")))]))])])).Elements:
                    d_97_v0_: _dafny.Seq = compr_9_
                    if (d_97_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lll"))) for d_95_i0_ in range(default__.abs(148))]), _dafny.SeqWithoutIsStrInference([17]), _dafny.SeqWithoutIsStrInference([len(_dafny.Set({(0) - (-719), -391, -329})) for d_96_i1_ in range(default__.abs(200))]), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([620, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "caiutom")))]))])])):
                        coll9_[d_97_v0_] = False
                return _dafny.Map(coll9_)
            coll7_ = _dafny.Set()
            def iife8_():
                coll8_ = _dafny.Map()
                compr_8_: _dafny.Seq
                for compr_8_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lll"))) for d_95_i0_ in range(default__.abs(148))]), _dafny.SeqWithoutIsStrInference([17]), _dafny.SeqWithoutIsStrInference([len(_dafny.Set({(0) - (-719), -391, -329})) for d_96_i1_ in range(default__.abs(200))]), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([620, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "caiutom")))]))])])).Elements:
                    d_97_v0_: _dafny.Seq = compr_8_
                    if (d_97_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lll"))) for d_95_i0_ in range(default__.abs(148))]), _dafny.SeqWithoutIsStrInference([17]), _dafny.SeqWithoutIsStrInference([len(_dafny.Set({(0) - (-719), -391, -329})) for d_96_i1_ in range(default__.abs(200))]), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([620, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "caiutom")))]))])])):
                        coll8_[d_97_v0_] = False
                return _dafny.Map(coll8_)
            compr_7_: _dafny.Seq
            for compr_7_ in (iife8_()
            ).keys.Elements:
                d_98_v1_: _dafny.Seq = compr_7_
                if (d_98_v1_) in (iife9_()
                ):
                    coll7_ = coll7_.union(_dafny.Set([d_98_v1_]))
            return _dafny.Set(coll7_)
        return (iife7_()
        ).intersection(_dafny.Set({_dafny.SeqWithoutIsStrInference([765, 590]), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([-195, len(_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference([True]))})), (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([False, True])), 901])).cardinality])), len(_dafny.SeqWithoutIsStrInference([False, not(True)])), (0) - (-526), len(_dafny.Map({253: (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([774]))).cardinality})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xvljky")))])}))

    @staticmethod
    def fm57(p0, p1, p2, globalState):
        def iife10_():
            coll10_ = _dafny.Map()
            compr_10_: int
            for compr_10_ in _dafny.IntegerRange(159, 372):
                d_99_v0_: int = compr_10_
                if ((159) <= (d_99_v0_)) and ((d_99_v0_) < (372)):
                    coll10_[(d_99_v0_) * ((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o")))))] = True
            return _dafny.Map(coll10_)
        return iife10_()
        

    @staticmethod
    def fm58(p0, p1, p2, globalState):
        source6_ = D18_DC39(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ofia")): True}))
        if source6_.is_DC38:
            d_100___mcc_h0_ = source6_.cf59
            d_101_cf59_ = d_100___mcc_h0_
            return D1_DC1(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g")))
        elif source6_.is_DC39:
            d_102___mcc_h1_ = source6_.cf60
            d_103_cf60_ = d_102___mcc_h1_
            return D1_DC1(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "emqidpefj")))
        elif source6_.is_DC37:
            d_104___mcc_h2_ = source6_.cf58
            d_105_cf58_ = d_104___mcc_h2_
            return D1_DC1(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ycb")))
        elif True:
            d_106___mcc_h3_ = source6_.cf61
            d_107_cf61_ = d_106___mcc_h3_
            if True:
                return D1_DC1(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_108_i0_ in range(default__.abs(801))]))
            elif True:
                return D1_DC1(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bcuy")))

    @staticmethod
    def fm59(p0, p1, p2, globalState):
        return (_dafny.SeqWithoutIsStrInference([D0_DC0(True, 649, 770)])) + ((_dafny.SeqWithoutIsStrInference([D0_DC0(not(False), 123, 10)])) + (_dafny.SeqWithoutIsStrInference([D0_DC0(False, 47, 478)])))

    @staticmethod
    def fm60(p0, globalState):
        return _dafny.Set({_dafny.CodePoint('p')})

    @staticmethod
    def fm61(p0, p1, p2, p3, globalState):
        return D11_DC26(False, (442) not in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({D5_DC11(D2_DC6(D2_DC5(_dafny.CodePoint('r'), True, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kvtd"))), 790)), True)}))]))])), len(_dafny.SeqWithoutIsStrInference([751, -702])), (True) in (_dafny.Map({True: _dafny.Map({351: False})})))

    @staticmethod
    def fm62(p0, p1, p2, globalState):
        return D7_DC16((True if True else not(False)), (-915) - (len(_dafny.SeqWithoutIsStrInference([-194]))), (_dafny.SeqWithoutIsStrInference([False, not(not(False))])) == (_dafny.SeqWithoutIsStrInference([False])), (True) not in (_dafny.SeqWithoutIsStrInference([True, False])), not((_dafny.MultiSet([800])) == (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([71 for d_109_i0_ in range(default__.abs(-213))]))]))))

    @staticmethod
    def fm63(globalState):
        return _dafny.Map({len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fnrthabr")))])): _dafny.CodePoint('r')})

    @staticmethod
    def fm64(globalState):
        return D10_DC22(len(_dafny.SeqWithoutIsStrInference([683])), not((_dafny.Set({512})).isdisjoint(_dafny.Set({-327}))))

    @staticmethod
    def fm65(p0, p1, p2, p3, globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([827])) + (_dafny.SeqWithoutIsStrInference([38, -867]))) for d_110_i0_ in range(default__.abs(865))])

    @staticmethod
    def fm66(p0, p1, globalState):
        if not((_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cs")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y"))})).ispropersubset(_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fsxy")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dpujhkh"))}))):
            return D11_DC27()
        elif True:
            return D11_DC27()

    @staticmethod
    def fm67(p0, p1, p2, globalState):
        return (_dafny.Map({not(True): D3_DC8(True, (_dafny.MultiSet([True])).cardinality, 549)})) | ((_dafny.Map({not(False): D3_DC8(not(not(True)), (_dafny.MultiSet([True, False])).cardinality, 590)})) | (_dafny.Map({False: D3_DC8(False, -117, 367)})))

    @staticmethod
    def fm68(globalState):
        source7_ = D7_DC17(D7_DC16(True, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jdresf"))), True, True, True))
        if source7_.is_DC15:
            d_111___mcc_h0_ = source7_.cf28
            d_112___mcc_h1_ = source7_.cf29
            d_113_cf29_ = d_112___mcc_h1_
            d_114_cf28_ = d_111___mcc_h0_
            def iife11_():
                def iife13_():
                    coll13_ = _dafny.Map()
                    compr_13_: _dafny.Seq
                    for compr_13_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([-763 for d_115_i0_ in range(default__.abs(298))]), _dafny.SeqWithoutIsStrInference([756, (D24_DC57(len(_dafny.Map({d_113_cf29_: len(_dafny.SeqWithoutIsStrInference([d_114_cf28_ for d_116_i1_ in range(default__.abs(746))]))})), 839)).cf90, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uqbpxmqu"))) for d_117_i2_ in range(default__.abs(-876))]))).cardinality, 650]), _dafny.SeqWithoutIsStrInference([187, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({len(_dafny.Set({d_114_cf28_})): d_113_cf29_}))]))).cardinality])])).Elements:
                        d_118_v1_: _dafny.Seq = compr_13_
                        if (d_118_v1_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([-763 for d_115_i0_ in range(default__.abs(298))]), _dafny.SeqWithoutIsStrInference([756, (D24_DC57(len(_dafny.Map({d_113_cf29_: len(_dafny.SeqWithoutIsStrInference([d_114_cf28_ for d_116_i1_ in range(default__.abs(746))]))})), 839)).cf90, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uqbpxmqu"))) for d_117_i2_ in range(default__.abs(-876))]))).cardinality, 650]), _dafny.SeqWithoutIsStrInference([187, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({len(_dafny.Set({d_114_cf28_})): d_113_cf29_}))]))).cardinality])])):
                            coll13_[d_118_v1_] = _dafny.CodePoint('r')
                    return _dafny.Map(coll13_)
                coll11_ = _dafny.Map()
                def iife12_():
                    coll12_ = _dafny.Map()
                    compr_12_: _dafny.Seq
                    for compr_12_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([-763 for d_115_i0_ in range(default__.abs(298))]), _dafny.SeqWithoutIsStrInference([756, (D24_DC57(len(_dafny.Map({d_113_cf29_: len(_dafny.SeqWithoutIsStrInference([d_114_cf28_ for d_116_i1_ in range(default__.abs(746))]))})), 839)).cf90, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uqbpxmqu"))) for d_117_i2_ in range(default__.abs(-876))]))).cardinality, 650]), _dafny.SeqWithoutIsStrInference([187, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({len(_dafny.Set({d_114_cf28_})): d_113_cf29_}))]))).cardinality])])).Elements:
                        d_118_v1_: _dafny.Seq = compr_12_
                        if (d_118_v1_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([-763 for d_115_i0_ in range(default__.abs(298))]), _dafny.SeqWithoutIsStrInference([756, (D24_DC57(len(_dafny.Map({d_113_cf29_: len(_dafny.SeqWithoutIsStrInference([d_114_cf28_ for d_116_i1_ in range(default__.abs(746))]))})), 839)).cf90, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uqbpxmqu"))) for d_117_i2_ in range(default__.abs(-876))]))).cardinality, 650]), _dafny.SeqWithoutIsStrInference([187, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({len(_dafny.Set({d_114_cf28_})): d_113_cf29_}))]))).cardinality])])):
                            coll12_[d_118_v1_] = _dafny.CodePoint('r')
                    return _dafny.Map(coll12_)
                compr_11_: _dafny.Seq
                for compr_11_ in (iife12_()
                ).keys.Elements:
                    d_119_v0_: _dafny.Seq = compr_11_
                    if (d_119_v0_) in (iife13_()
                    ):
                        coll11_[d_119_v0_] = d_113_cf29_
                return _dafny.Map(coll11_)
            return (_dafny.Map({_dafny.SeqWithoutIsStrInference([284]): not(d_113_cf29_)})) | (iife11_()
            )
        elif source7_.is_DC16:
            d_120___mcc_h2_ = source7_.cf30
            d_121___mcc_h3_ = source7_.cf31
            d_122___mcc_h4_ = source7_.cf32
            d_123___mcc_h5_ = source7_.cf33
            d_124___mcc_h6_ = source7_.cf34
            d_125_cf34_ = d_124___mcc_h6_
            d_126_cf33_ = d_123___mcc_h5_
            d_127_cf32_ = d_122___mcc_h4_
            d_128_cf31_ = d_121___mcc_h3_
            d_129_cf30_ = d_120___mcc_h2_
            def iife14_():
                coll14_ = _dafny.Map()
                compr_14_: _dafny.Seq
                for compr_14_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(0) - ((0) - (d_128_cf31_)), d_128_cf31_])])).Elements:
                    d_130_v2_: _dafny.Seq = compr_14_
                    if (d_130_v2_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(0) - ((0) - (d_128_cf31_)), d_128_cf31_])])):
                        coll14_[d_130_v2_] = d_125_cf34_
                return _dafny.Map(coll14_)
            return (iife14_()
            ) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([d_128_cf31_]): d_127_cf32_}))
        elif source7_.is_DC14:
            d_131___mcc_h7_ = source7_.cf27
            d_132_cf27_ = d_131___mcc_h7_
            def iife15_():
                coll15_ = _dafny.Map()
                compr_15_: _dafny.Seq
                for compr_15_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([False, True])).cardinality])])).Elements:
                    d_133_v3_: _dafny.Seq = compr_15_
                    if (d_133_v3_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([False, True])).cardinality])])):
                        coll15_[d_133_v3_] = False
                return _dafny.Map(coll15_)
            return iife15_()
            
        elif True:
            d_134___mcc_h8_ = source7_.cf35
            d_135_cf35_ = d_134___mcc_h8_
            return (_dafny.Map({_dafny.SeqWithoutIsStrInference([758, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r")))]): True})) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: False}))]): True}))

    @staticmethod
    def m10(p0, p1, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: bool = False
        d_136_v0_: _dafny.Array
        nw0_ = _dafny.Array(None, 1)
        nw0_[int(0)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yiqwrp"))
        d_136_v0_ = nw0_
        d_137_v1_: _dafny.Array
        d_137_v1_ = d_136_v0_
        d_137_v1_ = d_137_v1_
        def lambda1_(d_138_p1_):
            def lambda2_(d_139_i0_):
                def iife16_():
                    coll16_ = _dafny.Set()
                    compr_16_: _dafny.Map
                    for compr_16_ in (_dafny.MultiSet([_dafny.Map({(D10_DC23(_dafny.SeqWithoutIsStrInference([d_138_p1_]), d_138_p1_)).cf44: 400})])).Elements:
                        d_140_v2_: _dafny.Map = compr_16_
                        if (d_140_v2_) in (_dafny.MultiSet([_dafny.Map({(D10_DC23(_dafny.SeqWithoutIsStrInference([d_138_p1_]), d_138_p1_)).cf44: 400})])):
                            coll16_ = coll16_.union(_dafny.Set([d_140_v2_]))
                    return _dafny.Set(coll16_)
                return (iife16_()
                ).ispropersubset(_dafny.Set({_dafny.Map({d_138_p1_: len(_dafny.Set({d_138_p1_}))}), _dafny.Map({d_138_p1_: 428})}))

            return lambda2_

        init0_ = lambda1_(p1)
        nw1_ = _dafny.Array(None, 19)
        for i0_0_ in range(nw1_.length(0)):
            nw1_[i0_0_] = init0_(i0_0_)
        (globalState).f2 = nw1_
        d_141_v4_: _dafny.Seq
        d_141_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i"))
        d_142_v5_: _dafny.Seq
        d_142_v5_ = _dafny.SeqWithoutIsStrInference([False])
        d_143_v6_: int
        d_143_v6_ = 660
        d_144_v7_: _dafny.Map
        d_144_v7_ = _dafny.Map({len(d_141_v4_): (d_142_v5_)[default__.safeIndex(d_143_v6_, len(d_142_v5_))]})
        d_145_v8_: _dafny.Set
        def iife17_():
            coll17_ = _dafny.Set()
            compr_17_: int
            for compr_17_ in (p0).Elements:
                d_146_v3_: int = compr_17_
                if (d_146_v3_) in (p0):
                    coll17_ = coll17_.union(_dafny.Set([default__.safeModuloInt(d_146_v3_, 549)]))
            return _dafny.Set(coll17_)
        d_145_v8_ = _dafny.Set({default__.fm0(len(iife17_()
        ), p1, d_144_v7_, d_143_v6_, globalState)})
        d_147_i1_: int
        d_147_i1_ = 0
        with _dafny.label("0"):
            while (d_145_v8_).issubset(d_145_v8_):
                with _dafny.c_label("0"):
                    if (d_147_i1_) >= (100):
                        raise _dafny.Break("0")
                    d_147_i1_ = (d_147_i1_) + (1)
                    d_148_v9_: _dafny.Map
                    d_148_v9_ = _dafny.Map({d_143_v6_: len(d_142_v5_)})
                    d_149_v10_: _dafny.Seq
                    d_149_v10_ = _dafny.SeqWithoutIsStrInference([default__.safeModuloInt(-266, len(d_148_v9_))])
                    d_149_v10_ = (D6_DC12(p0)).cf23
                    d_150_v11_: D1
                    d_150_v11_ = D1_DC1(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hofyrvjtu")))
                    d_151_v12_: C8
                    nw2_ = C8()
                    nw2_.ctor__(p1, p1)
                    d_151_v12_ = nw2_
                    d_152_v13_: _dafny.Set
                    d_152_v13_ = _dafny.Set({d_151_v12_})
                    rhs0_ = d_150_v11_
                    rhs1_ = (d_152_v13_).isdisjoint(d_152_v13_)
                    lhs0_ = globalState
                    d_150_v11_ = rhs0_
                    lhs0_.f10 = rhs1_
                    d_143_v6_ = d_143_v6_
                    d_153_v14_: _dafny.Array
                    nw3_ = _dafny.Array(_dafny.Array(None, 0), 11)
                    d_153_v14_ = nw3_
                    d_154_v15_: str
                    d_154_v15_ = _dafny.CodePoint('m')
                    d_155_v16_: C6
                    nw4_ = C6()
                    nw4_.ctor__(p1, p1, p1, d_154_v15_, p1)
                    d_155_v16_ = nw4_
                    d_156_v17_: _dafny.Map
                    d_156_v17_ = _dafny.Map({d_155_v16_: d_143_v6_})
                    d_157_v18_: _dafny.Map
                    d_157_v18_ = _dafny.Map({p1: 784})
                    d_158_v19_: D8
                    d_158_v19_ = D8_DC19((d_155_v16_).f28, (p0)[default__.safeIndex(d_143_v6_, len(p0))])
                    d_159_v20_: _dafny.MultiSet
                    d_159_v20_ = _dafny.MultiSet([(d_155_v16_).f28, (d_155_v16_).f28])
                    d_160_v21_: _dafny.Array
                    nw5_ = _dafny.Array(None, 22)
                    nw5_[int(0)] = d_143_v6_
                    nw5_[int(1)] = d_143_v6_
                    nw5_[int(2)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wefgrtm")))
                    nw5_[int(3)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ertxgegm")))
                    nw5_[int(4)] = ((d_156_v17_)[d_155_v16_] if (d_155_v16_) in (d_156_v17_) else d_143_v6_)
                    nw5_[int(5)] = 618
                    nw5_[int(6)] = d_143_v6_
                    nw5_[int(7)] = 381
                    nw5_[int(8)] = d_143_v6_
                    nw5_[int(9)] = d_143_v6_
                    nw5_[int(10)] = d_143_v6_
                    nw5_[int(11)] = d_143_v6_
                    nw5_[int(12)] = d_143_v6_
                    nw5_[int(13)] = d_143_v6_
                    nw5_[int(14)] = len(d_157_v18_)
                    nw5_[int(15)] = (d_158_v19_).cf38
                    nw5_[int(16)] = d_143_v6_
                    nw5_[int(17)] = d_143_v6_
                    nw5_[int(18)] = 483
                    nw5_[int(19)] = d_143_v6_
                    nw5_[int(20)] = d_143_v6_
                    nw5_[int(21)] = (d_159_v20_).cardinality
                    d_160_v21_ = nw5_
                    index0_ = default__.safeIndex(561, (d_153_v14_).length(0))
                    (d_153_v14_)[index0_] = d_160_v21_
                    index1_ = default__.safeIndex(561, (d_153_v14_).length(0))
                    (d_153_v14_)[index1_] = d_160_v21_
                    pass
            pass
        d_161_v22_: str
        d_161_v22_ = _dafny.CodePoint('x')
        d_162_v23_: C6
        nw6_ = C6()
        nw6_.ctor__(p1, p1, p1, d_161_v22_, p1)
        d_162_v23_ = nw6_
        d_163_v24_: D27
        d_163_v24_ = D27_DC64(d_162_v23_)
        d_164_v25_: _dafny.Map
        d_164_v25_ = _dafny.Map({(d_163_v24_).cf104: False})
        d_165_i2_: int
        d_165_i2_ = 0
        with _dafny.label("1"):
            while ((636) - ((0) - (d_143_v6_))) != (((0) - (len(d_164_v25_)) if p1 else d_143_v6_)):
                with _dafny.c_label("1"):
                    if (d_165_i2_) >= (100):
                        raise _dafny.Break("1")
                    d_165_i2_ = (d_165_i2_) + (1)
                    d_166_v27_: _dafny.Array
                    def lambda3_(d_167_v6_):
                        def lambda4_(d_168_i3_):
                            def iife18_():
                                coll18_ = _dafny.Map()
                                compr_18_: int
                                for compr_18_ in _dafny.IntegerRange(-540, 160):
                                    d_169_v26_: int = compr_18_
                                    if ((-540) <= (d_169_v26_)) and ((d_169_v26_) < (160)):
                                        coll18_[(d_169_v26_) - (d_167_v6_)] = d_167_v6_
                                return _dafny.Map(coll18_)
                            return (iife18_()
                            ) | (_dafny.Map({d_167_v6_: d_167_v6_}))

                        return lambda4_

                    init1_ = lambda3_(d_143_v6_)
                    nw7_ = _dafny.Array(None, 13)
                    for i0_1_ in range(nw7_.length(0)):
                        nw7_[i0_1_] = init1_(i0_1_)
                    d_166_v27_ = nw7_
                    d_170_v28_: _dafny.Map
                    d_170_v28_ = _dafny.Map({False: d_143_v6_})
                    d_171_v29_: C0
                    nw8_ = C0()
                    nw8_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qnyihlu")))
                    d_171_v29_ = nw8_
                    d_172_v30_: _dafny.Map
                    d_172_v30_ = _dafny.Map({d_171_v29_: d_143_v6_})
                    d_173_v31_: _dafny.Map
                    d_173_v31_ = _dafny.Map({d_143_v6_: len(d_172_v30_)})
                    index2_ = default__.safeIndex(300, (d_166_v27_).length(0))
                    (d_166_v27_)[index2_] = (_dafny.Map({((d_170_v28_)[p1] if (p1) in (d_170_v28_) else len(p0)): len(d_141_v4_)})) | (d_173_v31_)
                    index3_ = default__.safeIndex(300, (d_166_v27_).length(0))
                    (d_166_v27_)[index3_] = d_173_v31_
                    d_174_v32_: _dafny.Array
                    nw9_ = _dafny.Array(False, 12)
                    d_174_v32_ = nw9_
                    index4_ = default__.safeIndex(658, (d_174_v32_).length(0))
                    (d_174_v32_)[index4_] = (d_162_v23_).f28
                    d_175_v33_: _dafny.MultiSet
                    d_175_v33_ = _dafny.MultiSet([d_143_v6_])
                    d_176_v34_: _dafny.Map
                    d_176_v34_ = _dafny.Map({False: d_171_v29_})
                    d_177_v35_: _dafny.Map
                    d_177_v35_ = _dafny.Map({(0) - (d_143_v6_): d_176_v34_})
                    index5_ = default__.safeIndex(658, (d_174_v32_).length(0))
                    rhs2_ = p1
                    rhs3_ = p1
                    rhs4_ = (((d_175_v33_)[d_143_v6_] if (d_143_v6_) in (d_175_v33_) else d_143_v6_)) - (((0) - (len(d_177_v35_))) - (561))
                    lhs1_ = d_174_v32_
                    lhs2_ = default__.safeIndex(658, (d_174_v32_).length(0))
                    lhs3_ = globalState
                    r1 = rhs2_
                    lhs1_[lhs2_] = rhs3_
                    lhs3_.f11 = rhs4_
                    (globalState).f11 = 415
                    d_178_v36_: _dafny.Array
                    def lambda5_(d_179_v6_):
                        def lambda6_(d_180_i4_):
                            return (d_180_i4_) - (d_179_v6_)

                        return lambda6_

                    init2_ = lambda5_(d_143_v6_)
                    nw10_ = _dafny.Array(None, 18)
                    for i0_2_ in range(nw10_.length(0)):
                        nw10_[i0_2_] = init2_(i0_2_)
                    d_178_v36_ = nw10_
                    d_181_v37_: D1
                    d_181_v37_ = D1_DC2(d_178_v36_, ((d_171_v29_).f31)[default__.safeIndex(857, len((d_171_v29_).f31))], (((d_166_v27_)[default__.safeIndex(300, (d_166_v27_).length(0))])[d_143_v6_] if (d_143_v6_) in ((d_166_v27_)[default__.safeIndex(300, (d_166_v27_).length(0))]) else d_143_v6_), not(default__.fm1((d_162_v23_).f28, (d_162_v23_).fm7(d_170_v28_, globalState), d_143_v6_, globalState)))
                    rhs5_ = (d_174_v32_)[default__.safeIndex(658, (d_174_v32_).length(0))]
                    rhs6_ = ((d_141_v4_) + ((d_171_v29_).f31)) <= ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w"))) + (d_141_v4_))
                    rhs7_ = d_142_v5_
                    rhs8_ = (d_181_v37_).cf5
                    r1 = rhs5_
                    r1 = rhs6_
                    d_142_v5_ = rhs7_
                    d_161_v22_ = rhs8_
                    pass
            pass
        d_182_v38_: C7
        nw11_ = C7()
        nw11_.ctor__((d_162_v23_).f28)
        d_182_v38_ = nw11_
        d_183_v39_: _dafny.Array
        nw12_ = _dafny.Array(_dafny.Map({}), 19)
        d_183_v39_ = nw12_
        d_184_v40_: _dafny.Seq
        d_184_v40_ = _dafny.SeqWithoutIsStrInference([d_183_v39_, d_183_v39_])
        d_185_v41_: _dafny.Set
        d_185_v41_ = _dafny.Set({p1})
        d_186_v42_: _dafny.Map
        d_186_v42_ = _dafny.Map({p1: len(d_185_v41_)})
        d_187_v43_: C2
        nw13_ = C2()
        nw13_.ctor__(d_141_v4_, default__.fm38((d_162_v23_).f28, d_143_v6_, _dafny.MultiSet(p0), d_186_v42_, globalState))
        d_187_v43_ = nw13_
        d_188_v44_: D8
        d_188_v44_ = D8_DC18((d_184_v40_)[default__.safeIndex(default__.fm0(d_143_v6_, p1, d_144_v7_, len(_dafny.Map({d_187_v43_: d_143_v6_})), globalState), len(d_184_v40_))])
        d_189_v45_: _dafny.Map
        d_189_v45_ = _dafny.Map({d_188_v44_: d_143_v6_})
        d_190_v46_: _dafny.Map
        d_190_v46_ = _dafny.Map({(d_143_v6_) + (d_143_v6_): (d_189_v45_) | (d_189_v45_)})
        d_190_v46_ = d_190_v46_
        r0 = d_136_v0_
        r1 = p1
        return r0, r1

    @staticmethod
    def Main(noArgsParameter__):
        d_191_v0_: _dafny.Array
        nw14_ = _dafny.Array(False, 4)
        d_191_v0_ = nw14_
        d_192_v1_: int
        d_192_v1_ = -248
        d_193_v2_: _dafny.Seq
        d_193_v2_ = _dafny.SeqWithoutIsStrInference([False])
        d_194_v3_: bool
        d_194_v3_ = True
        d_195_v4_: _dafny.Map
        d_195_v4_ = _dafny.Map({d_194_v3_: len(d_193_v2_)})
        d_196_v5_: _dafny.Array
        nw15_ = _dafny.Array(None, 9)
        nw15_[int(0)] = d_192_v1_
        nw15_[int(1)] = d_192_v1_
        nw15_[int(2)] = d_192_v1_
        nw15_[int(3)] = d_192_v1_
        nw15_[int(4)] = len(d_193_v2_)
        nw15_[int(5)] = d_192_v1_
        nw15_[int(6)] = d_192_v1_
        nw15_[int(7)] = ((d_195_v4_)[d_194_v3_] if (d_194_v3_) in (d_195_v4_) else len(d_193_v2_))
        nw15_[int(8)] = d_192_v1_
        d_196_v5_ = nw15_
        d_197_v6_: str
        d_197_v6_ = _dafny.CodePoint('x')
        d_198_v7_: _dafny.Set
        d_198_v7_ = _dafny.Set({d_197_v6_, d_197_v6_})
        d_199_v8_: _dafny.Map
        d_199_v8_ = _dafny.Map({d_194_v3_: d_198_v7_})
        d_200_v9_: _dafny.Seq
        d_200_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ylhvnb"))
        d_201_v10_: _dafny.Map
        d_201_v10_ = _dafny.Map({392: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_202_i0_ in range(default__.abs(837))])})
        d_203_v11_: _dafny.Array
        nw16_ = _dafny.Array(None, 22)
        nw16_[int(0)] = d_200_v9_
        nw16_[int(1)] = d_200_v9_
        nw16_[int(2)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g"))
        nw16_[int(3)] = d_200_v9_
        nw16_[int(4)] = d_200_v9_
        nw16_[int(5)] = d_200_v9_
        nw16_[int(6)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hi"))
        nw16_[int(7)] = d_200_v9_
        nw16_[int(8)] = d_200_v9_
        nw16_[int(9)] = d_200_v9_
        nw16_[int(10)] = d_200_v9_
        nw16_[int(11)] = d_200_v9_
        nw16_[int(12)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "npevt"))
        nw16_[int(13)] = d_200_v9_
        nw16_[int(14)] = d_200_v9_
        nw16_[int(15)] = d_200_v9_
        nw16_[int(16)] = ((d_201_v10_)[698] if (698) in (d_201_v10_) else (d_200_v9_).set(default__.safeIndex(d_192_v1_, len(d_200_v9_)), d_197_v6_))
        nw16_[int(17)] = d_200_v9_
        nw16_[int(18)] = d_200_v9_
        nw16_[int(19)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dh"))
        nw16_[int(20)] = d_200_v9_
        nw16_[int(21)] = d_200_v9_
        d_203_v11_ = nw16_
        d_204_globalState_: GlobalState
        nw17_ = GlobalState()
        nw17_.ctor__(False, 33, d_191_v0_, True, -999, 553, d_191_v0_, d_196_v5_, True, d_199_v8_, True, -891, True, 896, 252, d_193_v2_, 737, False, True, 808, d_203_v11_, -557, 226, False)
        d_204_globalState_ = nw17_
        d_194_v3_ = not(not(d_194_v3_))
        (d_204_globalState_).f10 = (d_200_v9_) != (_dafny.SeqWithoutIsStrInference([d_197_v6_ for d_205_i1_ in range(default__.abs(-755))]))
        d_206_i2_: int
        d_206_i2_ = 0
        with _dafny.label("2"):
            while not (d_194_v3_) or (d_194_v3_):
                with _dafny.c_label("2"):
                    if (d_206_i2_) >= (100):
                        raise _dafny.Break("2")
                    d_206_i2_ = (d_206_i2_) + (1)
                    d_207_v12_: D0
                    d_207_v12_ = D0_DC0(d_194_v3_, d_192_v1_, d_192_v1_)
                    source8_ = d_207_v12_
                    d_208___mcc_h0_ = source8_.cf0
                    d_209___mcc_h1_ = source8_.cf1
                    d_210___mcc_h2_ = source8_.cf2
                    d_211_cf2_ = d_210___mcc_h2_
                    d_212_cf1_ = d_209___mcc_h1_
                    d_213_cf0_ = d_208___mcc_h0_
                    d_197_v6_ = d_197_v6_
                    d_192_v1_ = d_192_v1_
                    index6_ = default__.safeIndex(18, (d_196_v5_).length(0))
                    (d_196_v5_)[index6_] = d_192_v1_
                    d_214_v13_: _dafny.Array
                    nw18_ = _dafny.Array(_dafny.Array(None, 0), 22)
                    d_214_v13_ = nw18_
                    index7_ = default__.safeIndex(748, (d_214_v13_).length(0))
                    (d_214_v13_)[index7_] = d_196_v5_
                    d_215_v14_: _dafny.Array
                    def lambda7_(d_216_v6_):
                        def lambda8_(d_217_i3_):
                            return d_216_v6_

                        return lambda8_

                    init3_ = lambda7_(d_197_v6_)
                    nw19_ = _dafny.Array(None, 20)
                    for i0_3_ in range(nw19_.length(0)):
                        nw19_[i0_3_] = init3_(i0_3_)
                    d_215_v14_ = nw19_
                    d_218_v15_: _dafny.Map
                    d_218_v15_ = _dafny.Map({d_215_v14_: d_213_cf0_})
                    pat_let_tv0_ = d_213_cf0_
                    d_219_v16_: _dafny.Map
                    def iife19_(_pat_let0_0):
                        def iife20_(d_220_dt__update__tmp_h0_):
                            def iife21_(_pat_let1_0):
                                def iife22_(d_221_dt__update_hcf0_h0_):
                                    return D0_DC0(d_221_dt__update_hcf0_h0_, (d_220_dt__update__tmp_h0_).cf1, (d_220_dt__update__tmp_h0_).cf2)
                                return iife22_(_pat_let1_0)
                            return iife21_(pat_let_tv0_)
                        return iife20_(_pat_let0_0)
                    d_219_v16_ = _dafny.Map({iife19_(D0_DC0(d_194_v3_, d_192_v1_, d_212_cf1_)): d_196_v5_})
                    index8_ = default__.safeIndex(18, (d_196_v5_).length(0))
                    index9_ = default__.safeIndex(748, (d_214_v13_).length(0))
                    rhs9_ = len(d_218_v15_)
                    rhs10_ = default__.safeDivisionInt((d_211_cf2_) + (d_211_cf2_), d_212_cf1_)
                    rhs11_ = ((d_219_v16_)[d_207_v12_] if (d_207_v12_) in (d_219_v16_) else d_196_v5_)
                    lhs4_ = d_196_v5_
                    lhs5_ = default__.safeIndex(18, (d_196_v5_).length(0))
                    lhs6_ = d_214_v13_
                    lhs7_ = default__.safeIndex(748, (d_214_v13_).length(0))
                    lhs4_[lhs5_] = rhs9_
                    d_192_v1_ = rhs10_
                    lhs6_[lhs7_] = rhs11_
                    d_222_v17_: _dafny.Map
                    d_222_v17_ = _dafny.Map({(d_196_v5_)[default__.safeIndex(18, (d_196_v5_).length(0))]: (d_196_v5_)[default__.safeIndex(18, (d_196_v5_).length(0))]})
                    d_223_v18_: _dafny.Set
                    d_223_v18_ = _dafny.Set({_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "umttmamld"))): d_192_v1_}), (d_222_v17_) | (d_222_v17_), (_dafny.Map({(d_196_v5_)[default__.safeIndex(18, (d_196_v5_).length(0))]: d_212_cf1_})) | (d_222_v17_), d_222_v17_})
                    d_223_v18_ = (d_223_v18_) - (d_223_v18_)
                    def iife23_():
                        coll19_ = _dafny.Map()
                        compr_19_: int
                        for compr_19_ in (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({d_192_v1_: d_192_v1_})) for d_225_i5_ in range(default__.abs(766))])): d_192_v1_})).keys.Elements:
                            d_226_v19_: int = compr_19_
                            if (d_226_v19_) in (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({d_192_v1_: d_192_v1_})) for d_225_i5_ in range(default__.abs(766))])): d_192_v1_})):
                                coll19_[(d_226_v19_) + (d_192_v1_)] = d_194_v3_
                        return _dafny.Map(coll19_)
                    if (d_192_v1_) == (default__.fm0(d_192_v1_, default__.fm1(False, d_194_v3_, len(_dafny.SeqWithoutIsStrInference([d_192_v1_ for d_224_i4_ in range(default__.abs(-112))])), d_204_globalState_), iife23_()
                    , d_192_v1_, d_204_globalState_)):
                        d_227_v20_: _dafny.MultiSet
                        d_227_v20_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([d_192_v1_ for d_228_i6_ in range(default__.abs(175))]))])
                        d_229_v21_: _dafny.Map
                        d_229_v21_ = _dafny.Map({False: d_227_v20_})
                        d_230_v22_: _dafny.Seq
                        d_230_v22_ = _dafny.SeqWithoutIsStrInference([len(d_229_v21_), d_192_v1_, d_192_v1_])
                        (d_204_globalState_).f11 = (d_230_v22_)[default__.safeIndex(d_192_v1_, len(d_230_v22_))]
                        d_231_v23_: _dafny.Map
                        d_231_v23_ = _dafny.Map({len(d_195_v4_): d_194_v3_})
                        d_232_v24_: _dafny.Seq
                        d_232_v24_ = _dafny.SeqWithoutIsStrInference([d_231_v23_])
                        d_233_v25_: _dafny.Map
                        d_233_v25_ = _dafny.Map({d_194_v3_: d_194_v3_})
                        d_234_v26_: _dafny.Map
                        d_234_v26_ = _dafny.Map({(default__.fm2(_dafny.SeqWithoutIsStrInference([d_194_v3_, d_194_v3_, d_194_v3_]), d_194_v3_, (0) - (len(d_200_v9_)), default__.fm0(d_192_v1_, not(d_194_v3_), (d_232_v24_)[default__.safeIndex(d_192_v1_, len(d_232_v24_))], d_192_v1_, d_204_globalState_), d_204_globalState_)) | (d_233_v25_): len((d_200_v9_).set(default__.safeIndex(default__.fm0(d_192_v1_, not(d_194_v3_), d_231_v23_, d_192_v1_, d_204_globalState_), len(d_200_v9_)), d_197_v6_))})
                        rhs12_ = d_197_v6_
                        rhs13_ = d_194_v3_
                        rhs14_ = (d_234_v26_) | (d_234_v26_)
                        rhs15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ekekwwlsn"))
                        rhs16_ = len((_dafny.SeqWithoutIsStrInference([d_200_v9_ for d_235_i7_ in range(default__.abs(183))])) + (default__.fm3(d_193_v2_, d_192_v1_, d_197_v6_, d_194_v3_, d_204_globalState_)))
                        lhs8_ = d_204_globalState_
                        d_197_v6_ = rhs12_
                        d_194_v3_ = rhs13_
                        d_234_v26_ = rhs14_
                        d_200_v9_ = rhs15_
                        lhs8_.f22 = rhs16_
                        (d_204_globalState_).f10 = ((d_230_v22_) + (d_230_v22_)) == (d_230_v22_)
                        d_236_v27_: _dafny.Map
                        d_236_v27_ = _dafny.Map({(0) - (d_192_v1_): len(d_193_v2_)})
                        d_237_v28_: C3
                        nw20_ = C3()
                        nw20_.ctor__((((d_236_v27_)[d_192_v1_] if (d_192_v1_) in (d_236_v27_) else len(d_200_v9_))) < (d_192_v1_), d_200_v9_, d_194_v3_, d_197_v6_, d_194_v3_, d_194_v3_)
                        d_237_v28_ = nw20_
                        d_238_v29_: bool
                        d_239_v30_: bool
                        d_240_v31_: _dafny.Map
                        out0_: bool
                        out1_: bool
                        out2_: _dafny.Map
                        out0_, out1_, out2_ = (d_237_v28_).m3((0) - (d_192_v1_), (d_227_v20_).issubset(default__.fm52((d_237_v28_).f32, d_204_globalState_)), d_191_v0_, d_204_globalState_)
                        d_238_v29_ = out0_
                        d_239_v30_ = out1_
                        d_240_v31_ = out2_
                    elif True:
                        d_241_v32_: _dafny.Set
                        d_241_v32_ = _dafny.Set({d_194_v3_})
                        d_242_v33_: _dafny.Map
                        d_242_v33_ = _dafny.Map({d_241_v32_: d_200_v9_})
                        d_243_v34_: C2
                        nw21_ = C2()
                        nw21_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rb")), ((d_242_v33_)[d_241_v32_] if (d_241_v32_) in (d_242_v33_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gxgemu"))))
                        d_243_v34_ = nw21_
                        d_244_v35_: _dafny.MultiSet
                        d_244_v35_ = _dafny.MultiSet([d_192_v1_, len(d_193_v2_)])
                        d_245_v36_: _dafny.Seq
                        d_245_v36_ = _dafny.SeqWithoutIsStrInference([(d_244_v35_).cardinality, d_192_v1_])
                        d_246_v37_: _dafny.Array
                        nw22_ = _dafny.Array(None, 26)
                        nw22_[int(0)] = _dafny.SeqWithoutIsStrInference([d_194_v3_, False])
                        nw22_[int(1)] = _dafny.SeqWithoutIsStrInference([d_194_v3_, d_194_v3_])
                        nw22_[int(2)] = d_193_v2_
                        nw22_[int(3)] = d_193_v2_
                        nw22_[int(4)] = _dafny.SeqWithoutIsStrInference([d_194_v3_])
                        nw22_[int(5)] = d_193_v2_
                        nw22_[int(6)] = d_193_v2_
                        nw22_[int(7)] = d_193_v2_
                        nw22_[int(8)] = d_193_v2_
                        nw22_[int(9)] = _dafny.SeqWithoutIsStrInference([d_194_v3_])
                        nw22_[int(10)] = d_193_v2_
                        nw22_[int(11)] = _dafny.SeqWithoutIsStrInference([d_194_v3_])
                        nw22_[int(12)] = d_193_v2_
                        nw22_[int(13)] = d_193_v2_
                        nw22_[int(14)] = d_193_v2_
                        nw22_[int(15)] = (d_193_v2_).set(default__.safeIndex(d_192_v1_, len(d_193_v2_)), d_194_v3_)
                        nw22_[int(16)] = _dafny.SeqWithoutIsStrInference([d_194_v3_, d_194_v3_])
                        nw22_[int(17)] = d_193_v2_
                        nw22_[int(18)] = _dafny.SeqWithoutIsStrInference([d_194_v3_])
                        nw22_[int(19)] = d_193_v2_
                        nw22_[int(20)] = d_193_v2_
                        nw22_[int(21)] = d_193_v2_
                        nw22_[int(22)] = d_193_v2_
                        nw22_[int(23)] = d_193_v2_
                        nw22_[int(24)] = default__.fm9(_dafny.Map({d_194_v3_: len(d_245_v36_)}), d_192_v1_, d_197_v6_, len((d_243_v34_).f30), d_204_globalState_)
                        nw22_[int(25)] = d_193_v2_
                        d_246_v37_ = nw22_
                        d_247_v38_: _dafny.Map
                        d_247_v38_ = _dafny.Map({d_243_v34_: d_246_v37_})
                        d_247_v38_ = (d_247_v38_).set(d_243_v34_, d_246_v37_)
                        index10_ = default__.safeIndex(284, (d_196_v5_).length(0))
                        (d_196_v5_)[index10_] = d_192_v1_
                        index11_ = default__.safeIndex(284, (d_196_v5_).length(0))
                        (d_196_v5_)[index11_] = d_192_v1_
                        (d_204_globalState_).f10 = d_194_v3_
                        d_248_v39_: C5
                        nw23_ = C5()
                        nw23_.ctor__(d_194_v3_, d_197_v6_, False, d_194_v3_)
                        d_248_v39_ = nw23_
                        d_249_v40_: D22
                        d_249_v40_ = D22_DC52(d_192_v1_, d_194_v3_, False)
                        nw24_ = C5()
                        nw24_.ctor__((d_249_v40_).cf82, d_197_v6_, not(d_194_v3_), (d_243_v34_).fm18((0) - (d_192_v1_), d_204_globalState_))
                        d_248_v39_ = nw24_
                        d_250_v41_: bool
                        d_251_v42_: _dafny.Array
                        out3_: bool
                        out4_: _dafny.Array
                        out3_, out4_ = (d_248_v39_).m0(d_204_globalState_)
                        d_250_v41_ = out3_
                        d_251_v42_ = out4_
                    d_252_v43_: _dafny.MultiSet
                    d_252_v43_ = _dafny.MultiSet([d_192_v1_])
                    (d_204_globalState_).f10 = default__.fm1(d_194_v3_, d_194_v3_, ((d_252_v43_)[d_192_v1_] if (d_192_v1_) in (d_252_v43_) else d_192_v1_), d_204_globalState_)
                    d_253_v44_: C9
                    nw25_ = C9()
                    nw25_.ctor__(d_194_v3_, d_197_v6_, d_194_v3_, d_194_v3_)
                    d_253_v44_ = nw25_
                    d_254_v45_: _dafny.Seq
                    d_254_v45_ = _dafny.SeqWithoutIsStrInference([d_253_v44_])
                    (d_204_globalState_).f16 = default__.safeDivisionInt(d_192_v1_, (len(d_254_v45_)) - (d_192_v1_))
                    pass
            pass
        d_255_v47_: _dafny.Array
        def lambda9_(d_256_v6_, d_257_v1_):
            def lambda10_(d_258_i8_):
                def iife24_():
                    coll20_ = _dafny.Set()
                    compr_20_: int
                    for compr_20_ in (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([d_256_v6_ for d_259_i9_ in range(default__.abs(242))])), d_257_v1_])).Elements:
                        d_260_v46_: int = compr_20_
                        if (d_260_v46_) in (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([d_256_v6_ for d_259_i9_ in range(default__.abs(242))])), d_257_v1_])):
                            coll20_ = coll20_.union(_dafny.Set([(d_260_v46_) * (436)]))
                    return _dafny.Set(coll20_)
                return iife24_()
                

            return lambda10_

        init4_ = lambda9_(d_197_v6_, d_192_v1_)
        nw26_ = _dafny.Array(None, 23)
        for i0_4_ in range(nw26_.length(0)):
            nw26_[i0_4_] = init4_(i0_4_)
        d_255_v47_ = nw26_
        d_255_v47_ = d_255_v47_
        if (d_192_v1_) > ((d_192_v1_) - (d_192_v1_)):
            d_261_v48_: C0
            nw27_ = C0()
            nw27_.ctor__(d_200_v9_)
            d_261_v48_ = nw27_
            d_262_v49_: D25
            d_262_v49_ = D25_DC58(d_261_v48_)
            d_263_v50_: _dafny.Seq
            d_263_v50_ = _dafny.SeqWithoutIsStrInference([(d_262_v49_).cf92, d_261_v48_])
            d_264_v51_: _dafny.Map
            d_264_v51_ = _dafny.Map({d_194_v3_: (d_263_v50_) < (_dafny.SeqWithoutIsStrInference([d_261_v48_, d_261_v48_]))})
            d_264_v51_ = (d_264_v51_).set(not (d_194_v3_) or (d_194_v3_), (_dafny.SeqWithoutIsStrInference([d_194_v3_])) <= (d_193_v2_))
            d_265_v52_: _dafny.Seq
            d_265_v52_ = _dafny.SeqWithoutIsStrInference([d_191_v0_])
            d_266_v53_: D21
            d_266_v53_ = D21_DC47(d_265_v52_)
            source9_ = d_266_v53_
            if source9_.is_DC48:
                d_267_v54_: C1
                nw28_ = C1()
                nw28_.ctor__(d_194_v3_)
                d_267_v54_ = nw28_
                d_268_v55_: _dafny.Set
                d_268_v55_ = _dafny.Set({d_192_v1_, d_192_v1_, d_192_v1_, d_192_v1_, d_192_v1_})
                d_269_v56_: _dafny.Map
                d_269_v56_ = _dafny.Map({len(d_268_v55_): d_267_v54_})
                d_270_v57_: _dafny.MultiSet
                d_270_v57_ = _dafny.MultiSet([d_192_v1_, d_192_v1_])
                d_271_v58_: D19
                d_271_v58_ = D19_DC43(d_192_v1_, ((d_269_v56_)[d_192_v1_] if (d_192_v1_) in (d_269_v56_) else d_267_v54_), d_270_v57_, d_192_v1_)
                d_272_v59_: D19
                d_272_v59_ = D19_DC43(d_192_v1_, d_267_v54_, (d_271_v58_).cf67, d_192_v1_)
                d_273_v60_: _dafny.Map
                d_273_v60_ = _dafny.Map({d_271_v58_: d_194_v3_})
                d_274_v61_: _dafny.MultiSet
                d_274_v61_ = _dafny.MultiSet([d_196_v5_])
                (d_204_globalState_).f10 = not((((d_273_v60_)[d_271_v58_] if (d_271_v58_) in (d_273_v60_) else True)) or (not((d_274_v61_).issubset(d_274_v61_))))
                (d_204_globalState_).f10 = d_194_v3_
                (d_204_globalState_).f11 = d_192_v1_
                d_275_v62_: _dafny.Array
                nw29_ = _dafny.Array(D20.default()(), 29)
                d_275_v62_ = nw29_
                d_276_v63_: C8
                nw30_ = C8()
                nw30_.ctor__(d_194_v3_, d_194_v3_)
                d_276_v63_ = nw30_
                d_277_v64_: _dafny.Seq
                d_277_v64_ = _dafny.SeqWithoutIsStrInference([d_276_v63_])
                d_278_v65_: D20
                d_278_v65_ = D20_DC44((d_277_v64_)[default__.safeIndex(len(d_195_v4_), len(d_277_v64_))])
                index12_ = default__.safeIndex(456, (d_275_v62_).length(0))
                (d_275_v62_)[index12_] = d_278_v65_
                index13_ = default__.safeIndex(456, (d_275_v62_).length(0))
                (d_275_v62_)[index13_] = d_278_v65_
            elif source9_.is_DC49:
                d_279___mcc_h3_ = source9_.cf73
                d_280___mcc_h4_ = source9_.cf74
                d_281___mcc_h5_ = source9_.cf75
                d_282___mcc_h6_ = source9_.cf76
                d_283___mcc_h7_ = source9_.cf77
                d_284_cf77_ = d_283___mcc_h7_
                d_285_cf76_ = d_282___mcc_h6_
                d_286_cf75_ = d_281___mcc_h5_
                d_287_cf74_ = d_280___mcc_h4_
                d_288_cf73_ = d_279___mcc_h3_
                index14_ = default__.safeIndex(243, (d_191_v0_).length(0))
                (d_191_v0_)[index14_] = (d_288_cf73_) > (229)
                d_289_v66_: _dafny.Set
                d_289_v66_ = _dafny.Set({d_192_v1_})
                index15_ = default__.safeIndex(243, (d_191_v0_).length(0))
                (d_191_v0_)[index15_] = ((d_289_v66_ if d_286_cf75_ else d_289_v66_)).isdisjoint(d_289_v66_)
                (d_204_globalState_).f16 = (0) - (d_192_v1_)
                d_290_v67_: C7
                nw31_ = C7()
                nw31_.ctor__(d_194_v3_)
                d_290_v67_ = nw31_
                index16_ = default__.safeIndex(243, (d_191_v0_).length(0))
                rhs17_ = d_203_v11_
                rhs18_ = _dafny.CodePoint('o')
                rhs19_ = ((d_197_v6_ if d_194_v3_ else d_197_v6_)) not in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_291_i10_ in range(default__.abs(535))]))
                rhs20_ = d_290_v67_
                rhs21_ = (d_261_v48_).fm22(not((_dafny.MultiSet([(d_261_v48_).fm22(d_285_cf76_, d_204_globalState_), False])).isdisjoint(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_286_cf75_])))), d_204_globalState_)
                lhs9_ = d_191_v0_
                lhs10_ = default__.safeIndex(243, (d_191_v0_).length(0))
                d_203_v11_ = rhs17_
                d_197_v6_ = rhs18_
                d_285_cf76_ = rhs19_
                d_290_v67_ = rhs20_
                lhs9_[lhs10_] = rhs21_
                d_288_cf73_ = (0) - (d_288_cf73_)
            elif True:
                d_292___mcc_h8_ = source9_.cf72
                d_293_cf72_ = d_292___mcc_h8_
                d_294_v68_: D0
                d_294_v68_ = D0_DC0((_dafny.Map({True: d_192_v1_})) != (default__.fm32(d_204_globalState_)), d_192_v1_, 786)
                d_295_v69_: _dafny.Map
                d_295_v69_ = _dafny.Map({d_192_v1_: d_194_v3_})
                d_294_v68_ = (d_294_v68_ if ((d_295_v69_)[277] if (277) in (d_295_v69_) else True) else d_294_v68_)
                (d_204_globalState_).f10 = not((d_261_v48_).fm22(d_194_v3_, d_204_globalState_))
                rhs22_ = (d_261_v48_).fm22(d_194_v3_, d_204_globalState_)
                rhs23_ = True
                lhs11_ = d_204_globalState_
                lhs12_ = d_204_globalState_
                lhs11_.f10 = rhs22_
                lhs12_.f10 = rhs23_
                d_296_v70_: _dafny.Set
                d_296_v70_ = _dafny.Set({d_194_v3_, not(d_194_v3_)})
                (d_204_globalState_).f22 = default__.safeDivisionInt(default__.safeModuloInt(d_192_v1_, d_192_v1_), (len(d_296_v70_)) + (len(d_193_v2_)))
            d_297_v71_: _dafny.Seq
            d_297_v71_ = _dafny.SeqWithoutIsStrInference([d_192_v1_, d_192_v1_, d_192_v1_, 968])
            d_298_v72_: _dafny.Array
            d_299_v73_: bool
            out5_: _dafny.Array
            out6_: bool
            out5_, out6_ = default__.m10(d_297_v71_, d_194_v3_, d_204_globalState_)
            d_298_v72_ = out5_
            d_299_v73_ = out6_
            d_300_v74_: _dafny.MultiSet
            d_300_v74_ = _dafny.MultiSet([d_299_v73_, d_194_v3_, d_194_v3_])
            d_301_v75_: _dafny.Map
            d_301_v75_ = _dafny.Map({((d_300_v74_)[d_194_v3_] if (d_194_v3_) in (d_300_v74_) else d_192_v1_): d_194_v3_})
            d_302_v76_: _dafny.Array
            d_303_v77_: bool
            out7_: _dafny.Array
            out8_: bool
            out7_, out8_ = default__.m10((_dafny.SeqWithoutIsStrInference([d_192_v1_, default__.fm0(d_192_v1_, d_299_v73_, d_301_v75_, d_192_v1_, d_204_globalState_), d_192_v1_, d_192_v1_, 602])) + (d_297_v71_), d_299_v73_, d_204_globalState_)
            d_302_v76_ = out7_
            d_303_v77_ = out8_
            d_304_v78_: _dafny.Array
            d_305_v79_: bool
            out9_: _dafny.Array
            out10_: bool
            out9_, out10_ = default__.m10(_dafny.SeqWithoutIsStrInference([d_192_v1_, (0) - (len(d_193_v2_))]), d_303_v77_, d_204_globalState_)
            d_304_v78_ = out9_
            d_305_v79_ = out10_
        elif True:
            d_306_v80_: _dafny.Array
            nw32_ = _dafny.Array(_dafny.Map({}), 3)
            d_306_v80_ = nw32_
            d_307_v81_: _dafny.Map
            d_307_v81_ = _dafny.Map({857: d_192_v1_})
            index17_ = default__.safeIndex(31, (d_306_v80_).length(0))
            (d_306_v80_)[index17_] = d_307_v81_
            index18_ = default__.safeIndex(31, (d_306_v80_).length(0))
            def iife25_():
                coll21_ = _dafny.Map()
                compr_21_: int
                for compr_21_ in _dafny.IntegerRange(62, 420):
                    d_308_v82_: int = compr_21_
                    if ((62) <= (d_308_v82_)) and ((d_308_v82_) < (420)):
                        coll21_[default__.safeDivisionInt(d_308_v82_, d_192_v1_)] = d_192_v1_
                return _dafny.Map(coll21_)
            (d_306_v80_)[index18_] = (iife25_()
            ).set(d_192_v1_, d_192_v1_)
            d_309_v84_: D6
            def iife26_():
                coll22_ = _dafny.Map()
                compr_22_: int
                for compr_22_ in _dafny.IntegerRange(-679, 999):
                    d_310_v83_: int = compr_22_
                    if ((-679) <= (d_310_v83_)) and ((d_310_v83_) < (999)):
                        coll22_[(d_310_v83_) - (d_192_v1_)] = d_194_v3_
                return _dafny.Map(coll22_)
            d_309_v84_ = D6_DC13(d_191_v0_, default__.fm0(d_192_v1_, d_194_v3_, iife26_()
, d_192_v1_, d_204_globalState_), False)
            d_311_v85_: _dafny.Array
            nw33_ = _dafny.Array(None, 15)
            nw33_[int(0)] = d_191_v0_
            nw33_[int(1)] = d_191_v0_
            nw33_[int(2)] = d_191_v0_
            nw33_[int(3)] = d_191_v0_
            nw33_[int(4)] = d_191_v0_
            nw33_[int(5)] = d_191_v0_
            nw33_[int(6)] = d_191_v0_
            nw33_[int(7)] = d_191_v0_
            nw33_[int(8)] = d_191_v0_
            nw33_[int(9)] = d_191_v0_
            nw33_[int(10)] = d_191_v0_
            nw33_[int(11)] = d_191_v0_
            nw33_[int(12)] = (d_309_v84_).cf24
            nw33_[int(13)] = d_191_v0_
            nw33_[int(14)] = d_191_v0_
            d_311_v85_ = nw33_
            index19_ = default__.safeIndex(307, (d_311_v85_).length(0))
            (d_311_v85_)[index19_] = d_191_v0_
            d_312_v86_: _dafny.Array
            def lambda11_(d_313_v1_):
                def lambda12_(d_314_i11_):
                    return (_dafny.MultiSet([d_313_v1_])) - (_dafny.MultiSet([d_313_v1_]))

                return lambda12_

            init5_ = lambda11_(d_192_v1_)
            nw34_ = _dafny.Array(None, 27)
            for i0_5_ in range(nw34_.length(0)):
                nw34_[i0_5_] = init5_(i0_5_)
            d_312_v86_ = nw34_
            d_315_v87_: _dafny.MultiSet
            d_315_v87_ = _dafny.MultiSet([d_192_v1_])
            index20_ = default__.safeIndex(497, (d_312_v86_).length(0))
            (d_312_v86_)[index20_] = d_315_v87_
            index21_ = default__.safeIndex(307, (d_311_v85_).length(0))
            index22_ = default__.safeIndex(497, (d_312_v86_).length(0))
            rhs24_ = d_191_v0_
            rhs25_ = d_315_v87_
            lhs13_ = d_311_v85_
            lhs14_ = default__.safeIndex(307, (d_311_v85_).length(0))
            lhs15_ = d_312_v86_
            lhs16_ = default__.safeIndex(497, (d_312_v86_).length(0))
            lhs13_[lhs14_] = rhs24_
            lhs15_[lhs16_] = rhs25_
            index23_ = default__.safeIndex(643, (d_196_v5_).length(0))
            (d_196_v5_)[index23_] = -934
            index24_ = default__.safeIndex(643, (d_196_v5_).length(0))
            (d_196_v5_)[index24_] = default__.safeModuloInt(d_192_v1_, (0) - (((d_312_v86_)[default__.safeIndex(497, (d_312_v86_).length(0))]).cardinality))
            d_316_v88_: _dafny.MultiSet
            d_316_v88_ = _dafny.MultiSet([d_194_v3_, d_194_v3_])
            rhs26_ = ((d_194_v3_) and (default__.fm1(False, True, d_192_v1_, d_204_globalState_))) in (d_316_v88_)
            rhs27_ = d_197_v6_
            lhs17_ = d_204_globalState_
            lhs17_.f10 = rhs26_
            d_197_v6_ = rhs27_
            d_317_v89_: D22
            d_317_v89_ = D22_DC52(len((d_306_v80_)[default__.safeIndex(31, (d_306_v80_).length(0))]), d_194_v3_, d_194_v3_)
            pat_let_tv1_ = d_194_v3_
            pat_let_tv2_ = d_194_v3_
            d_318_v90_: _dafny.Map
            def iife27_(_pat_let2_0):
                def iife28_(d_319_dt__update__tmp_h1_):
                    def iife29_(_pat_let3_0):
                        def iife30_(d_320_dt__update_hcf82_h0_):
                            def iife31_(_pat_let4_0):
                                def iife32_(d_321_dt__update_hcf83_h0_):
                                    return D22_DC52((d_319_dt__update__tmp_h1_).cf81, d_320_dt__update_hcf82_h0_, d_321_dt__update_hcf83_h0_)
                                return iife32_(_pat_let4_0)
                            return iife31_(pat_let_tv2_)
                        return iife30_(_pat_let3_0)
                    return iife29_(pat_let_tv1_)
                return iife28_(_pat_let2_0)
            d_318_v90_ = _dafny.Map({default__.safeModuloInt(569, (d_196_v5_)[default__.safeIndex(643, (d_196_v5_).length(0))]): iife27_(d_317_v89_)})
            d_318_v90_ = (d_318_v90_).set((d_196_v5_)[default__.safeIndex(643, (d_196_v5_).length(0))], d_317_v89_)
        d_322_v91_: _dafny.MultiSet
        d_322_v91_ = _dafny.MultiSet([d_194_v3_, d_194_v3_, d_194_v3_, False, d_194_v3_])
        if ((d_322_v91_) | (d_322_v91_)).isdisjoint(d_322_v91_):
            d_323_v92_: C1
            nw35_ = C1()
            nw35_.ctor__(False)
            d_323_v92_ = nw35_
            d_324_v93_: _dafny.Map
            d_324_v93_ = _dafny.Map({(len(d_193_v2_)) + (d_192_v1_): len(_dafny.Set({d_323_v92_}))})
            d_325_v94_: T0
            nw36_ = C2()
            nw36_.ctor__(d_200_v9_, d_200_v9_)
            d_325_v94_ = nw36_
            d_326_v95_: _dafny.Seq
            d_326_v95_ = _dafny.SeqWithoutIsStrInference([d_325_v94_, d_325_v94_])
            d_327_v96_: _dafny.Seq
            d_327_v96_ = _dafny.SeqWithoutIsStrInference([d_326_v95_])
            d_324_v93_ = (d_324_v93_).set(len(d_327_v96_), d_192_v1_)
            d_328_v97_: _dafny.Array
            nw37_ = _dafny.Array(_dafny.CodePoint('D'), 15)
            d_328_v97_ = nw37_
            index25_ = default__.safeIndex(117, (d_328_v97_).length(0))
            (d_328_v97_)[index25_] = default__.fm36(d_192_v1_, not(default__.fm1(d_194_v3_, not(False), d_192_v1_, d_204_globalState_)), d_194_v3_, d_204_globalState_)
            d_329_v98_: _dafny.Set
            d_329_v98_ = _dafny.Set({d_192_v1_})
            d_330_v99_: _dafny.Seq
            d_330_v99_ = _dafny.SeqWithoutIsStrInference([d_192_v1_])
            index26_ = default__.safeIndex(117, (d_328_v97_).length(0))
            rhs28_ = default__.safeDivisionInt(885, default__.safeModuloInt(850, len(d_329_v98_)))
            rhs29_ = not (d_194_v3_) or ((d_192_v1_) not in (d_330_v99_))
            rhs30_ = d_197_v6_
            lhs18_ = d_204_globalState_
            lhs19_ = d_328_v97_
            lhs20_ = default__.safeIndex(117, (d_328_v97_).length(0))
            lhs18_.f22 = rhs28_
            d_194_v3_ = rhs29_
            lhs19_[lhs20_] = rhs30_
            d_331_v100_: D23
            d_331_v100_ = D23_DC55(d_192_v1_, d_194_v3_, d_194_v3_)
            d_331_v100_ = D23_DC55(len(d_200_v9_), d_194_v3_, True)
            (d_204_globalState_).f22 = 319
            d_197_v6_ = (d_197_v6_ if d_194_v3_ else (d_328_v97_)[default__.safeIndex(117, (d_328_v97_).length(0))])
        elif True:
            d_332_v101_: _dafny.Seq
            d_332_v101_ = _dafny.SeqWithoutIsStrInference([d_192_v1_, d_192_v1_])
            d_333_v102_: _dafny.Array
            d_334_v103_: bool
            out11_: _dafny.Array
            out12_: bool
            out11_, out12_ = default__.m10((d_332_v101_) + (d_332_v101_), d_194_v3_, d_204_globalState_)
            d_333_v102_ = out11_
            d_334_v103_ = out12_
            d_335_v104_: T1
            nw38_ = C1()
            nw38_.ctor__(d_334_v103_)
            d_335_v104_ = nw38_
            d_336_v105_: _dafny.Array
            nw39_ = _dafny.Array(None, 21)
            nw39_[int(0)] = d_335_v104_
            nw39_[int(1)] = d_335_v104_
            nw39_[int(2)] = d_335_v104_
            nw39_[int(3)] = d_335_v104_
            nw39_[int(4)] = d_335_v104_
            nw39_[int(5)] = d_335_v104_
            nw39_[int(6)] = d_335_v104_
            nw39_[int(7)] = d_335_v104_
            nw39_[int(8)] = d_335_v104_
            nw39_[int(9)] = d_335_v104_
            nw39_[int(10)] = d_335_v104_
            nw39_[int(11)] = d_335_v104_
            nw39_[int(12)] = d_335_v104_
            nw39_[int(13)] = d_335_v104_
            nw39_[int(14)] = d_335_v104_
            nw39_[int(15)] = d_335_v104_
            nw39_[int(16)] = d_335_v104_
            nw39_[int(17)] = d_335_v104_
            nw39_[int(18)] = d_335_v104_
            nw39_[int(19)] = d_335_v104_
            nw39_[int(20)] = d_335_v104_
            d_336_v105_ = nw39_
            index27_ = default__.safeIndex(479, (d_336_v105_).length(0))
            (d_336_v105_)[index27_] = d_335_v104_
            d_337_v106_: _dafny.Array
            nw40_ = _dafny.Array(None, 19)
            d_337_v106_ = nw40_
            d_338_v107_: T3
            nw41_ = C6()
            nw41_.ctor__(d_335_v104_.f24, d_334_v103_, d_334_v103_, d_197_v6_, d_334_v103_)
            d_338_v107_ = nw41_
            index28_ = default__.safeIndex(718, (d_337_v106_).length(0))
            (d_337_v106_)[index28_] = d_338_v107_
            index29_ = default__.safeIndex(479, (d_336_v105_).length(0))
            index30_ = default__.safeIndex(718, (d_337_v106_).length(0))
            rhs31_ = d_335_v104_
            rhs32_ = (d_338_v107_ if (d_338_v107_).f25 else d_338_v107_)
            lhs21_ = d_336_v105_
            lhs22_ = default__.safeIndex(479, (d_336_v105_).length(0))
            lhs23_ = d_337_v106_
            lhs24_ = default__.safeIndex(718, (d_337_v106_).length(0))
            lhs21_[lhs22_] = rhs31_
            lhs23_[lhs24_] = rhs32_
            (d_204_globalState_).f22 = (407) * (682)
            d_339_v108_: _dafny.Array
            nw42_ = _dafny.Array(_dafny.Array(None, 0), 23)
            d_339_v108_ = nw42_
            index31_ = default__.safeIndex(14, (d_339_v108_).length(0))
            (d_339_v108_)[index31_] = d_196_v5_
            index32_ = default__.safeIndex(14, (d_339_v108_).length(0))
            (d_339_v108_)[index32_] = d_196_v5_
            d_340_v109_: _dafny.Map
            d_340_v109_ = _dafny.Map({default__.fm1(d_194_v3_, d_194_v3_, len(d_200_v9_), d_204_globalState_): d_322_v91_})
            d_341_v110_: _dafny.Map
            d_341_v110_ = _dafny.Map({((d_340_v109_)[(d_338_v107_).f25] if ((d_338_v107_).f25) in (d_340_v109_) else _dafny.MultiSet([(d_338_v107_).f26, (d_338_v107_).f26])): d_338_v107_.f24})
            d_342_v111_: D26
            d_342_v111_ = D26_DC60(d_341_v110_)
            index33_ = default__.safeIndex(850, (d_191_v0_).length(0))
            (d_191_v0_)[index33_] = (len((d_342_v111_).cf96)) > (d_192_v1_)
            index34_ = default__.safeIndex(850, (d_191_v0_).length(0))
            (d_191_v0_)[index34_] = (d_338_v107_).f26
        d_343_v112_: _dafny.Array
        nw43_ = _dafny.Array(_dafny.Array(None, 0), 13)
        d_343_v112_ = nw43_
        index35_ = default__.safeIndex(85, (d_343_v112_).length(0))
        (d_343_v112_)[index35_] = d_203_v11_
        d_344_v113_: _dafny.MultiSet
        d_344_v113_ = _dafny.MultiSet([len(d_200_v9_), d_192_v1_, d_192_v1_])
        index36_ = default__.safeIndex(85, (d_343_v112_).length(0))
        rhs33_ = (default__.safeModuloInt((0) - (d_192_v1_), 598)) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kpsdwsoss"))))
        rhs34_ = d_203_v11_
        rhs35_ = (d_194_v3_ if d_194_v3_ else (d_193_v2_)[default__.safeIndex((d_344_v113_).cardinality, len(d_193_v2_))])
        lhs25_ = d_204_globalState_
        lhs26_ = d_343_v112_
        lhs27_ = default__.safeIndex(85, (d_343_v112_).length(0))
        lhs28_ = d_204_globalState_
        lhs25_.f11 = rhs33_
        lhs26_[lhs27_] = rhs34_
        lhs28_.f10 = rhs35_
        if (293) == ((d_192_v1_) - (d_192_v1_)):
            d_345_v114_: D15
            d_345_v114_ = D15_DC33((d_192_v1_) * (d_192_v1_))
            source10_ = d_345_v114_
            if source10_.is_DC33:
                d_346___mcc_h9_ = source10_.cf55
                d_347_cf55_ = d_346___mcc_h9_
                d_348_v115_: _dafny.Map
                d_348_v115_ = _dafny.Map({not (d_194_v3_) or (d_194_v3_): (d_347_cf55_) < (336)})
                d_348_v115_ = (d_348_v115_).set(False, d_194_v3_)
                d_349_v116_: _dafny.Array
                nw44_ = _dafny.Array(_dafny.Array(None, 0), 27)
                d_349_v116_ = nw44_
                d_350_v117_: C7
                nw45_ = C7()
                nw45_.ctor__(d_194_v3_)
                d_350_v117_ = nw45_
                d_351_v118_: _dafny.Seq
                d_351_v118_ = _dafny.SeqWithoutIsStrInference([d_347_cf55_])
                d_352_v119_: _dafny.Map
                d_352_v119_ = _dafny.Map({d_350_v117_: (d_351_v118_)[default__.safeIndex(d_192_v1_, len(d_351_v118_))]})
                d_353_v120_: _dafny.Map
                d_353_v120_ = _dafny.Map({((d_352_v119_)[d_350_v117_] if (d_350_v117_) in (d_352_v119_) else d_347_cf55_): d_349_v116_})
                d_354_v121_: _dafny.Map
                d_354_v121_ = _dafny.Map({357: d_194_v3_})
                d_349_v116_ = ((d_353_v120_)[default__.fm0(d_192_v1_, d_194_v3_, d_354_v121_, d_347_cf55_, d_204_globalState_)] if (default__.fm0(d_192_v1_, d_194_v3_, d_354_v121_, d_347_cf55_, d_204_globalState_)) in (d_353_v120_) else d_349_v116_)
                (d_204_globalState_).f7 = d_196_v5_
                index37_ = default__.safeIndex(916, (d_349_v116_).length(0))
                (d_349_v116_)[index37_] = d_191_v0_
                index38_ = default__.safeIndex(916, (d_349_v116_).length(0))
                (d_349_v116_)[index38_] = d_191_v0_
            elif source10_.is_DC34:
                index39_ = default__.safeIndex(433, (d_191_v0_).length(0))
                (d_191_v0_)[index39_] = d_194_v3_
                index40_ = default__.safeIndex(433, (d_191_v0_).length(0))
                (d_191_v0_)[index40_] = d_194_v3_
                d_355_v122_: _dafny.Map
                d_355_v122_ = _dafny.Map({d_192_v1_: d_194_v3_})
                (d_204_globalState_).f11 = default__.safeDivisionInt(default__.fm0(d_192_v1_, (d_191_v0_)[default__.safeIndex(433, (d_191_v0_).length(0))], d_355_v122_, d_192_v1_, d_204_globalState_), len(d_200_v9_))
                d_356_v123_: _dafny.Set
                d_356_v123_ = _dafny.Set({670})
                index41_ = default__.safeIndex(469, (d_196_v5_).length(0))
                (d_196_v5_)[index41_] = len((d_356_v123_).intersection(d_356_v123_))
                index42_ = default__.safeIndex(469, (d_196_v5_).length(0))
                (d_196_v5_)[index42_] = d_192_v1_
                d_357_v124_: _dafny.Seq
                d_357_v124_ = _dafny.SeqWithoutIsStrInference([d_200_v9_])
                d_358_v125_: C0
                nw46_ = C0()
                nw46_.ctor__((d_357_v124_)[default__.safeIndex(len(d_193_v2_), len(d_357_v124_))])
                d_358_v125_ = nw46_
            elif True:
                d_359___mcc_h10_ = source10_.cf54
                d_360_cf54_ = d_359___mcc_h10_
                d_361_v126_: T0
                nw47_ = C2()
                nw47_.ctor__(d_200_v9_, d_200_v9_)
                d_361_v126_ = nw47_
                d_362_v127_: _dafny.Map
                d_362_v127_ = _dafny.Map({True: d_361_v126_})
                d_363_v128_: _dafny.MultiSet
                d_363_v128_ = _dafny.MultiSet([d_361_v126_, d_361_v126_, d_361_v126_, d_361_v126_, ((d_362_v127_)[(d_193_v2_)[default__.safeIndex(d_192_v1_, len(d_193_v2_))]] if ((d_193_v2_)[default__.safeIndex(d_192_v1_, len(d_193_v2_))]) in (d_362_v127_) else d_361_v126_)])
                index43_ = default__.safeIndex(345, (d_196_v5_).length(0))
                (d_196_v5_)[index43_] = (d_363_v128_).cardinality
                index44_ = default__.safeIndex(345, (d_196_v5_).length(0))
                (d_196_v5_)[index44_] = d_192_v1_
                d_193_v2_ = (d_193_v2_) + ((_dafny.SeqWithoutIsStrInference([d_194_v3_, d_194_v3_, d_194_v3_, not(d_194_v3_), d_194_v3_]) if d_194_v3_ else d_193_v2_))
                index45_ = default__.safeIndex(379, (d_203_v11_).length(0))
                (d_203_v11_)[index45_] = d_200_v9_
                index46_ = default__.safeIndex(379, (d_203_v11_).length(0))
                (d_203_v11_)[index46_] = d_200_v9_
                d_364_v129_: _dafny.Array
                nw48_ = _dafny.Array(int(0), 9)
                d_364_v129_ = nw48_
                d_365_v130_: _dafny.Map
                d_365_v130_ = _dafny.Map({d_364_v129_: d_191_v0_})
                d_366_v131_: _dafny.Map
                d_366_v131_ = _dafny.Map({d_365_v130_: d_197_v6_})
                d_367_v132_: _dafny.Seq
                d_367_v132_ = _dafny.SeqWithoutIsStrInference([d_191_v0_])
                d_366_v131_ = (d_366_v131_).set((_dafny.Map({d_364_v129_: (d_367_v132_)[default__.safeIndex((d_196_v5_)[default__.safeIndex(345, (d_196_v5_).length(0))], len(d_367_v132_))]})) | (d_365_v130_), d_197_v6_)
            d_368_v133_: T0
            nw49_ = C2()
            nw49_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ryvrbqnqv")), d_200_v9_)
            d_368_v133_ = nw49_
            d_369_v135_: _dafny.Map
            def iife33_():
                coll23_ = _dafny.Map()
                compr_23_: int
                for compr_23_ in _dafny.IntegerRange(35, 11):
                    d_370_v134_: int = compr_23_
                    if ((35) <= (d_370_v134_)) and ((d_370_v134_) < (11)):
                        coll23_[default__.safeDivisionInt(d_370_v134_, d_192_v1_)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oka"))
                return _dafny.Map(coll23_)
            d_369_v135_ = _dafny.Map({d_368_v133_: len(iife33_()
            )})
            d_369_v135_ = (d_369_v135_).set(d_368_v133_, (d_192_v1_) * (d_192_v1_))
            d_200_v9_ = (d_200_v9_ if d_194_v3_ else d_200_v9_)
            (d_204_globalState_).f11 = d_192_v1_
            index47_ = default__.safeIndex(827, (d_203_v11_).length(0))
            (d_203_v11_)[index47_] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ocfk"))) + (d_200_v9_)
            index48_ = default__.safeIndex(827, (d_203_v11_).length(0))
            (d_203_v11_)[index48_] = (d_200_v9_) + (d_200_v9_)
        elif True:
            d_371_v136_: D10
            d_371_v136_ = D10_DC22(d_192_v1_, d_194_v3_)
            d_372_v137_: _dafny.MultiSet
            d_372_v137_ = _dafny.MultiSet([_dafny.MultiSet([d_194_v3_]), (_dafny.MultiSet([d_194_v3_])) - (_dafny.MultiSet([d_194_v3_, not(not((d_371_v136_).cf42))])), d_322_v91_, (d_322_v91_) - (d_322_v91_), (d_322_v91_ if d_194_v3_ else d_322_v91_)])
            d_372_v137_ = (d_372_v137_ if (d_194_v3_ if d_194_v3_ else d_194_v3_) else d_372_v137_)
            (d_204_globalState_).f10 = not(not (d_194_v3_) or (d_194_v3_))
            index49_ = default__.safeIndex(313, (d_191_v0_).length(0))
            (d_191_v0_)[index49_] = not((d_200_v9_) != (_dafny.SeqWithoutIsStrInference([d_197_v6_ for d_373_i12_ in range(default__.abs(331))])))
            index50_ = default__.safeIndex(313, (d_191_v0_).length(0))
            (d_191_v0_)[index50_] = d_194_v3_
            d_374_v138_: _dafny.Map
            d_374_v138_ = _dafny.Map({(d_344_v113_).isdisjoint(d_344_v113_): d_200_v9_})
            d_374_v138_ = (d_374_v138_).set((d_200_v9_) <= (d_200_v9_), d_200_v9_)
            (d_204_globalState_).f7 = d_196_v5_
        d_375_i13_: int
        d_375_i13_ = 0
        with _dafny.label("3"):
            while not((not(d_194_v3_)) and (True)):
                with _dafny.c_label("3"):
                    if (d_375_i13_) >= (100):
                        raise _dafny.Break("3")
                    d_375_i13_ = (d_375_i13_) + (1)
                    d_376_v139_: _dafny.Set
                    d_376_v139_ = _dafny.Set({(d_192_v1_) <= (d_192_v1_)})
                    d_377_v140_: _dafny.Seq
                    d_377_v140_ = _dafny.SeqWithoutIsStrInference([d_196_v5_])
                    d_378_v141_: _dafny.Map
                    d_378_v141_ = _dafny.Map({d_192_v1_: d_194_v3_})
                    d_379_v142_: T3
                    nw50_ = C6()
                    nw50_.ctor__(d_194_v3_, not(d_194_v3_), True, d_197_v6_, d_194_v3_)
                    d_379_v142_ = nw50_
                    d_380_v143_: _dafny.Seq
                    d_380_v143_ = _dafny.SeqWithoutIsStrInference([d_379_v142_])
                    d_381_v144_: _dafny.Set
                    d_381_v144_ = _dafny.Set({(d_380_v143_)[default__.safeIndex(d_192_v1_, len(d_380_v143_))], d_379_v142_, d_379_v142_})
                    d_382_v145_: _dafny.Seq
                    d_382_v145_ = _dafny.SeqWithoutIsStrInference([len(d_377_v140_), default__.fm0(693, d_194_v3_, d_378_v141_, (0) - (len(d_381_v144_)), d_204_globalState_), d_192_v1_])
                    d_383_v146_: D23
                    d_383_v146_ = D23_DC54(d_376_v139_)
                    rhs36_ = (d_382_v145_)[default__.safeIndex((0) - (default__.safeModuloInt(d_192_v1_, d_192_v1_)), len(d_382_v145_))]
                    rhs37_ = (d_383_v146_).cf85
                    rhs38_ = ((0) - ((d_192_v1_) - (d_192_v1_))) * (d_192_v1_)
                    lhs29_ = d_204_globalState_
                    lhs30_ = d_204_globalState_
                    lhs29_.f16 = rhs36_
                    d_376_v139_ = rhs37_
                    lhs30_.f11 = rhs38_
                    index51_ = default__.safeIndex(425, (d_191_v0_).length(0))
                    (d_191_v0_)[index51_] = d_194_v3_
                    index52_ = default__.safeIndex(425, (d_191_v0_).length(0))
                    (d_191_v0_)[index52_] = d_194_v3_
                    def iife34_():
                        coll24_ = _dafny.Set()
                        compr_24_: int
                        for compr_24_ in _dafny.IntegerRange(530, 621):
                            d_384_v147_: int = compr_24_
                            if ((530) <= (d_384_v147_)) and ((d_384_v147_) < (621)):
                                coll24_ = coll24_.union(_dafny.Set([default__.safeModuloInt(d_384_v147_, d_192_v1_)]))
                        return _dafny.Set(coll24_)
                    (d_204_globalState_).f11 = len(iife34_()
                    )
                    if (d_379_v142_).f25:
                        d_385_v148_: C4
                        nw51_ = C4()
                        nw51_.ctor__(d_379_v142_.f24, (d_191_v0_)[default__.safeIndex(425, (d_191_v0_).length(0))])
                        d_385_v148_ = nw51_
                        d_386_v149_: _dafny.Map
                        d_386_v149_ = _dafny.Map({(d_379_v142_).f26: d_385_v148_})
                        d_387_v150_: _dafny.Array
                        nw52_ = _dafny.Array(None, 17)
                        nw52_[int(0)] = d_385_v148_
                        nw52_[int(1)] = d_385_v148_
                        nw52_[int(2)] = d_385_v148_
                        nw52_[int(3)] = d_385_v148_
                        nw52_[int(4)] = d_385_v148_
                        nw52_[int(5)] = d_385_v148_
                        nw52_[int(6)] = d_385_v148_
                        nw52_[int(7)] = d_385_v148_
                        nw52_[int(8)] = d_385_v148_
                        nw52_[int(9)] = ((d_386_v149_)[d_194_v3_] if (d_194_v3_) in (d_386_v149_) else d_385_v148_)
                        nw52_[int(10)] = d_385_v148_
                        nw52_[int(11)] = d_385_v148_
                        nw52_[int(12)] = d_385_v148_
                        nw52_[int(13)] = d_385_v148_
                        nw52_[int(14)] = d_385_v148_
                        nw52_[int(15)] = d_385_v148_
                        nw52_[int(16)] = d_385_v148_
                        d_387_v150_ = nw52_
                        index53_ = default__.safeIndex(308, (d_387_v150_).length(0))
                        (d_387_v150_)[index53_] = d_385_v148_
                        d_388_v151_: T0
                        nw53_ = C2()
                        nw53_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "daltn")), d_200_v9_)
                        d_388_v151_ = nw53_
                        d_389_v152_: _dafny.Map
                        d_389_v152_ = _dafny.Map({d_388_v151_: d_385_v148_})
                        index54_ = default__.safeIndex(308, (d_387_v150_).length(0))
                        (d_387_v150_)[index54_] = ((d_389_v152_)[d_388_v151_] if (d_388_v151_) in (d_389_v152_) else d_385_v148_)
                        d_194_v3_ = (d_379_v142_).f25
                        (d_204_globalState_).f16 = (0) - (((0) - (((d_344_v113_)[d_192_v1_] if (d_192_v1_) in (d_344_v113_) else d_192_v1_))) * ((0) - (default__.safeDivisionInt(d_192_v1_, (_dafny.MultiSet([(d_379_v142_).f26])).cardinality))))
                        d_390_v153_: _dafny.Map
                        d_390_v153_ = _dafny.Map({d_197_v6_: (d_200_v9_) + (d_200_v9_)})
                        (d_204_globalState_).f22 = (0) - (len(d_390_v153_))
                        (d_204_globalState_).f16 = 562
                    elif True:
                        d_391_v154_: C7
                        nw54_ = C7()
                        nw54_.ctor__(d_379_v142_.f24)
                        d_391_v154_ = nw54_
                        index55_ = default__.safeIndex(170, (d_196_v5_).length(0))
                        (d_196_v5_)[index55_] = d_192_v1_
                        d_392_v155_: _dafny.Set
                        d_392_v155_ = _dafny.Set({838, d_192_v1_})
                        index56_ = default__.safeIndex(170, (d_196_v5_).length(0))
                        index57_ = default__.safeIndex(425, (d_191_v0_).length(0))
                        rhs39_ = (0) - (default__.safeModuloInt(d_192_v1_, (d_192_v1_ if d_379_v142_.f24 else d_192_v1_)))
                        rhs40_ = ((len(d_378_v141_) if True else (0) - (d_192_v1_)) if not(d_379_v142_.f24) else d_192_v1_)
                        rhs41_ = not(default__.fm1((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ff"))) < (d_200_v9_), ((d_379_v142_).f26) == ((d_379_v142_).f25), len(d_392_v155_), d_204_globalState_))
                        lhs31_ = d_196_v5_
                        lhs32_ = default__.safeIndex(170, (d_196_v5_).length(0))
                        lhs33_ = d_191_v0_
                        lhs34_ = default__.safeIndex(425, (d_191_v0_).length(0))
                        d_192_v1_ = rhs39_
                        lhs31_[lhs32_] = rhs40_
                        lhs33_[lhs34_] = rhs41_
                        d_393_v156_: _dafny.Map
                        d_393_v156_ = _dafny.Map({(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_192_v1_]))).issubset(d_344_v113_): (d_191_v0_)[default__.safeIndex(425, (d_191_v0_).length(0))]})
                        d_393_v156_ = (d_393_v156_).set(d_379_v142_.f24, True)
                        d_394_v157_: _dafny.Array
                        d_395_v158_: bool
                        out13_: _dafny.Array
                        out14_: bool
                        out13_, out14_ = default__.m10((d_382_v145_) + (d_382_v145_), (d_379_v142_).f25, d_204_globalState_)
                        d_394_v157_ = out13_
                        d_395_v158_ = out14_
                        d_396_v159_: C9
                        nw55_ = C9()
                        nw55_.ctor__((d_191_v0_)[default__.safeIndex(425, (d_191_v0_).length(0))], d_197_v6_, (d_192_v1_) > ((_dafny.MultiSet([(d_379_v142_).f25])).cardinality), (d_379_v142_).f26)
                        d_396_v159_ = nw55_
                        d_396_v159_ = d_396_v159_
                    pass
            pass
        d_397_v160_: C4
        nw56_ = C4()
        nw56_.ctor__(d_194_v3_, d_194_v3_)
        d_397_v160_ = nw56_
        d_398_v161_: _dafny.Set
        d_398_v161_ = _dafny.Set({d_397_v160_, d_397_v160_})
        d_399_v162_: _dafny.Set
        d_399_v162_ = _dafny.Set({d_194_v3_, d_194_v3_, d_194_v3_})
        hi0_ = len(d_399_v162_)
        for d_400_i14_ in range(len(d_398_v161_), hi0_):
            d_401_v163_: _dafny.Array
            nw57_ = _dafny.Array(_dafny.CodePoint('D'), 22)
            d_401_v163_ = nw57_
            d_402_v164_: _dafny.Seq
            d_402_v164_ = _dafny.SeqWithoutIsStrInference([d_401_v163_, d_401_v163_])
            d_402_v164_ = (d_402_v164_).set(default__.safeIndex(len(d_200_v9_), len(d_402_v164_)), d_401_v163_)
            d_403_v165_: _dafny.Map
            d_403_v165_ = _dafny.Map({d_400_i14_: d_194_v3_})
            d_404_v166_: D10
            d_404_v166_ = D10_DC22(default__.fm0(d_192_v1_, d_194_v3_, d_403_v165_, len(d_200_v9_), d_204_globalState_), default__.fm1(d_194_v3_, True, d_192_v1_, d_204_globalState_))
            (d_204_globalState_).f16 = ((d_404_v166_ if d_194_v3_ else d_404_v166_)).cf41
            d_405_v167_: bool
            d_406_v168_: _dafny.Array
            out15_: bool
            out16_: _dafny.Array
            out15_, out16_ = (d_397_v160_).m0(d_204_globalState_)
            d_405_v167_ = out15_
            d_406_v168_ = out16_
            d_407_v169_: _dafny.Map
            d_407_v169_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_197_v6_ for d_408_i15_ in range(default__.abs(955))])): d_192_v1_})
            d_409_v170_: _dafny.MultiSet
            d_409_v170_ = _dafny.MultiSet([d_407_v169_, d_407_v169_])
            d_410_v171_: _dafny.Seq
            d_410_v171_ = _dafny.SeqWithoutIsStrInference([d_400_i14_])
            if (d_409_v170_).isdisjoint((d_409_v170_ if d_405_v167_ else (_dafny.MultiSet([_dafny.Map({d_400_i14_: default__.fm0(len(d_200_v9_), d_194_v3_, d_403_v165_, d_400_i14_, d_204_globalState_)})])).set(_dafny.Map({(d_410_v171_)[default__.safeIndex(d_192_v1_, len(d_410_v171_))]: len(d_410_v171_)}), default__.abs(len(d_195_v4_))))):
                def lambda13_(d_411_i14_):
                    def lambda14_(d_412_i16_):
                        return default__.safeModuloInt(d_412_i16_, d_411_i14_)

                    return lambda14_

                init6_ = lambda13_(d_400_i14_)
                nw58_ = _dafny.Array(None, 23)
                for i0_6_ in range(nw58_.length(0)):
                    nw58_[i0_6_] = init6_(i0_6_)
                (d_204_globalState_).f7 = nw58_
                d_413_v172_: _dafny.Array
                d_414_v173_: bool
                out17_: _dafny.Array
                out18_: bool
                out17_, out18_ = default__.m10(d_410_v171_, d_194_v3_, d_204_globalState_)
                d_413_v172_ = out17_
                d_414_v173_ = out18_
                d_415_v174_: C0
                nw59_ = C0()
                nw59_.ctor__((d_200_v9_) + (d_200_v9_))
                d_415_v174_ = nw59_
                d_410_v171_ = (_dafny.SeqWithoutIsStrInference([85 for d_416_i17_ in range(default__.abs(23))]) if d_414_v173_ else (d_410_v171_) + (d_410_v171_))
                d_417_v175_: C2
                nw60_ = C2()
                nw60_.ctor__((d_415_v174_).f31, (d_415_v174_).f31)
                d_417_v175_ = nw60_
                d_418_v176_: _dafny.Array
                nw61_ = _dafny.Array(_dafny.Seq({}), 10)
                d_418_v176_ = nw61_
                d_419_v177_: _dafny.Map
                d_419_v177_ = _dafny.Map({d_417_v175_: (d_418_v176_ if d_414_v173_ else d_418_v176_)})
                d_419_v177_ = (d_419_v177_).set(d_417_v175_, d_418_v176_)
            elif True:
                d_420_v178_: bool
                d_421_v179_: _dafny.Array
                out19_: bool
                out20_: _dafny.Array
                out19_, out20_ = (d_397_v160_).m0(d_204_globalState_)
                d_420_v178_ = out19_
                d_421_v179_ = out20_
                d_422_v180_: D6
                d_422_v180_ = D6_DC12(d_410_v171_)
                pat_let_tv3_ = d_405_v167_
                def iife35_(_pat_let5_0):
                    def iife36_(d_423_dt__update__tmp_h2_):
                        def iife37_(_pat_let6_0):
                            def iife38_(d_425_dt__update_hcf23_h0_):
                                return D6_DC12(d_425_dt__update_hcf23_h0_)
                            return iife38_(_pat_let6_0)
                        return iife37_(_dafny.SeqWithoutIsStrInference([(D0_DC0(pat_let_tv3_, d_400_i14_, d_400_i14_)).cf2 for d_424_i18_ in range(default__.abs(904))]))
                    return iife36_(_pat_let5_0)
                d_422_v180_ = iife35_(D6_DC12(d_410_v171_))
                (d_204_globalState_).f16 = ((d_410_v171_)[default__.safeIndex(d_192_v1_, len(d_410_v171_))]) * (d_192_v1_)
                index58_ = default__.safeIndex(337, (d_406_v168_).length(0))
                (d_406_v168_)[index58_] = d_405_v167_
                index59_ = default__.safeIndex(337, (d_406_v168_).length(0))
                (d_406_v168_)[index59_] = d_420_v178_
                d_426_v181_: C1
                nw62_ = C1()
                nw62_.ctor__(d_405_v167_)
                d_426_v181_ = nw62_
        d_427_v182_: D0
        d_427_v182_ = D0_DC0(d_194_v3_, d_192_v1_, d_192_v1_)
        d_428_v183_: _dafny.Map
        d_428_v183_ = _dafny.Map({d_192_v1_: d_194_v3_})
        d_429_v184_: _dafny.Seq
        d_429_v184_ = _dafny.SeqWithoutIsStrInference([d_192_v1_, d_192_v1_])
        d_430_v185_: _dafny.Set
        d_430_v185_ = _dafny.Set({d_192_v1_, len(d_429_v184_), d_192_v1_})
        (d_204_globalState_).f22 = (default__.safeModuloInt((d_397_v160_).fm6(d_427_v182_, d_204_globalState_), d_192_v1_)) + (default__.fm0(d_192_v1_, d_194_v3_, default__.fm57(d_192_v1_, d_428_v183_, d_430_v185_, d_204_globalState_), d_192_v1_, d_204_globalState_))
        d_431_v186_: C8
        nw63_ = C8()
        nw63_.ctor__(d_194_v3_, not(((0) - (d_192_v1_)) >= (len(d_193_v2_))))
        d_431_v186_ = nw63_
        d_431_v186_ = d_431_v186_
        d_432_v187_: _dafny.Map
        d_432_v187_ = _dafny.Map({d_429_v184_: d_194_v3_})
        d_433_v188_: _dafny.Map
        d_433_v188_ = _dafny.Map({(d_322_v91_).isdisjoint(d_322_v91_): (d_432_v187_).set(d_429_v184_, not(d_194_v3_))})
        d_434_v189_: D11
        d_434_v189_ = D11_DC28()
        d_435_v190_: _dafny.Map
        d_435_v190_ = _dafny.Map({d_194_v3_: d_434_v189_})
        d_433_v188_ = (d_433_v188_).set((d_193_v2_)[default__.safeIndex(len(d_435_v190_), len(d_193_v2_))], default__.fm68(d_204_globalState_))
        d_436_v191_: C9
        nw64_ = C9()
        nw64_.ctor__(default__.fm1(d_194_v3_, default__.fm1(d_194_v3_, d_194_v3_, d_192_v1_, d_204_globalState_), 835, d_204_globalState_), d_197_v6_, d_194_v3_, d_194_v3_)
        d_436_v191_ = nw64_
        d_437_v192_: _dafny.MultiSet
        d_437_v192_ = _dafny.MultiSet([d_436_v191_])
        d_438_v193_: _dafny.Map
        d_438_v193_ = _dafny.Map({((d_437_v192_) - (_dafny.MultiSet([d_436_v191_, d_436_v191_]))).cardinality: D23_DC54(d_399_v162_)})
        d_439_v194_: C7
        nw65_ = C7()
        nw65_.ctor__(True)
        d_439_v194_ = nw65_
        d_440_v195_: _dafny.MultiSet
        d_440_v195_ = _dafny.MultiSet([d_439_v194_])
        d_438_v193_ = (d_438_v193_).set(((d_440_v195_).cardinality) - (d_192_v1_), D23_DC54(d_399_v162_))
        (d_204_globalState_).f10 = True
        _dafny.print(_dafny.string_of((d_191_v0_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_191_v0_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_192_v1_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_193_v2_) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_194_v3_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_195_v4_) == (_dafny.Map({True: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_196_v5_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_196_v5_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_196_v5_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_196_v5_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_196_v5_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_196_v5_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_196_v5_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_196_v5_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_196_v5_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_197_v6_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_198_v7_) == (_dafny.Set({_dafny.CodePoint('x')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_199_v8_) == (_dafny.Map({True: _dafny.Set({_dafny.CodePoint('x')})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_200_v9_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_201_v10_) == (_dafny.Map({392: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j')])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_203_v11_)[0]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_203_v11_)[1]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_203_v11_)[2]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_203_v11_)[3]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_203_v11_)[4]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_203_v11_)[5]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_203_v11_)[6]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_203_v11_)[7]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_203_v11_)[8]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_203_v11_)[9]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_203_v11_)[10]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_203_v11_)[11]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_203_v11_)[12]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_203_v11_)[13]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_203_v11_)[14]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_203_v11_)[15]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_203_v11_)[16]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_203_v11_)[17]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_203_v11_)[18]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_203_v11_)[19]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_203_v11_)[20]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_203_v11_)[21]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_globalState_).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_globalState_.f2)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_globalState_.f2)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_globalState_.f2)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_globalState_.f2)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_globalState_.f2)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_globalState_.f2)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_globalState_.f2)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_globalState_.f2)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_globalState_.f2)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_globalState_.f2)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_globalState_.f2)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_globalState_.f2)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_globalState_.f2)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_globalState_.f2)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_globalState_.f2)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_globalState_.f2)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_globalState_.f2)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_globalState_.f2)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_globalState_.f2)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_globalState_).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_globalState_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_globalState_).f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_204_globalState_).f6)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_204_globalState_).f6)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_globalState_.f7)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_globalState_.f7)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_globalState_.f7)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_globalState_.f7)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_globalState_.f7)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_globalState_.f7)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_globalState_.f7)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_globalState_.f7)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_globalState_.f7)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_globalState_).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_204_globalState_).f9) == (_dafny.Map({True: _dafny.Set({_dafny.CodePoint('x')})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_204_globalState_.f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_204_globalState_.f11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_globalState_).f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_globalState_).f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_globalState_).f14))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_204_globalState_).f15) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_204_globalState_.f16))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_globalState_).f17))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_globalState_).f18))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_globalState_).f19))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_204_globalState_).f20)[0]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_204_globalState_).f20)[1]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_204_globalState_).f20)[2]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_204_globalState_).f20)[3]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_204_globalState_).f20)[4]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_204_globalState_).f20)[5]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_204_globalState_).f20)[6]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_204_globalState_).f20)[7]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_204_globalState_).f20)[8]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_204_globalState_).f20)[9]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_204_globalState_).f20)[10]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_204_globalState_).f20)[11]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_204_globalState_).f20)[12]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_204_globalState_).f20)[13]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_204_globalState_).f20)[14]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_204_globalState_).f20)[15]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_204_globalState_).f20)[16]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_204_globalState_).f20)[17]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_204_globalState_).f20)[18]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_204_globalState_).f20)[19]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_204_globalState_).f20)[20]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_204_globalState_).f20)[21]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_globalState_).f21))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_204_globalState_.f22))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_globalState_).f23))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_206_i2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_255_v47_)[0]) == (_dafny.Set({105512, 872}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_255_v47_)[1]) == (_dafny.Set({105512, 872}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_255_v47_)[2]) == (_dafny.Set({105512, 872}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_255_v47_)[3]) == (_dafny.Set({105512, 872}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_255_v47_)[4]) == (_dafny.Set({105512, 872}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_255_v47_)[5]) == (_dafny.Set({105512, 872}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_255_v47_)[6]) == (_dafny.Set({105512, 872}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_255_v47_)[7]) == (_dafny.Set({105512, 872}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_255_v47_)[8]) == (_dafny.Set({105512, 872}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_255_v47_)[9]) == (_dafny.Set({105512, 872}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_255_v47_)[10]) == (_dafny.Set({105512, 872}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_255_v47_)[11]) == (_dafny.Set({105512, 872}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_255_v47_)[12]) == (_dafny.Set({105512, 872}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_255_v47_)[13]) == (_dafny.Set({105512, 872}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_255_v47_)[14]) == (_dafny.Set({105512, 872}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_255_v47_)[15]) == (_dafny.Set({105512, 872}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_255_v47_)[16]) == (_dafny.Set({105512, 872}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_255_v47_)[17]) == (_dafny.Set({105512, 872}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_255_v47_)[18]) == (_dafny.Set({105512, 872}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_255_v47_)[19]) == (_dafny.Set({105512, 872}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_255_v47_)[20]) == (_dafny.Set({105512, 872}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_255_v47_)[21]) == (_dafny.Set({105512, 872}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_255_v47_)[22]) == (_dafny.Set({105512, 872}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_322_v91_) == (_dafny.MultiSet([True, True, True, True, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_343_v112_)[7])[0]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_343_v112_)[7])[1]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_343_v112_)[7])[2]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_343_v112_)[7])[3]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_343_v112_)[7])[4]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_343_v112_)[7])[5]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_343_v112_)[7])[6]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_343_v112_)[7])[7]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_343_v112_)[7])[8]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_343_v112_)[7])[9]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_343_v112_)[7])[10]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_343_v112_)[7])[11]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_343_v112_)[7])[12]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_343_v112_)[7])[13]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_343_v112_)[7])[14]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_343_v112_)[7])[15]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_343_v112_)[7])[16]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_343_v112_)[7])[17]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_343_v112_)[7])[18]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_343_v112_)[7])[19]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_343_v112_)[7])[20]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_343_v112_)[7])[21]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_344_v113_) == (_dafny.MultiSet([6, 2, 2]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_375_i13_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_398_v161_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_399_v162_) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_427_v182_).cf0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_427_v182_).cf1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_427_v182_).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_428_v183_) == (_dafny.Map({2: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_429_v184_) == (_dafny.SeqWithoutIsStrInference([2, 2]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_430_v185_) == (_dafny.Set({2}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_432_v187_) == (_dafny.Map({_dafny.SeqWithoutIsStrInference([2, 2]): True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_433_v188_) == (_dafny.Map({False: _dafny.Map({_dafny.SeqWithoutIsStrInference([758, 1]): True, _dafny.SeqWithoutIsStrInference([1]): True})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_435_v190_) == (_dafny.Map({True: D11_DC28()}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_437_v192_).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_438_v193_) == (_dafny.Map({0: D23_DC54(_dafny.Set({True})), -1: D23_DC54(_dafny.Set({True}))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_440_v195_).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC0(False, int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)

class D0_DC0(D0, NamedTuple('DC0', [('cf0', Any), ('cf1', Any), ('cf2', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC0({_dafny.string_of(self.cf0)}, {_dafny.string_of(self.cf1)}, {_dafny.string_of(self.cf2)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC0) and self.cf0 == __o.cf0 and self.cf1 == __o.cf1 and self.cf2 == __o.cf2
    def __hash__(self) -> int:
        return super().__hash__()


class D1:
    @classmethod
    def default(cls, ):
        return lambda: D1_DC2(_dafny.Array(None, 0), _dafny.CodePoint('D'), int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D1_DC2)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D1_DC1)

class D1_DC2(D1, NamedTuple('DC2', [('cf4', Any), ('cf5', Any), ('cf6', Any), ('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC2({_dafny.string_of(self.cf4)}, {_dafny.string_of(self.cf5)}, {_dafny.string_of(self.cf6)}, {_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC2) and self.cf4 == __o.cf4 and self.cf5 == __o.cf5 and self.cf6 == __o.cf6 and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC1(D1, NamedTuple('DC1', [('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC1({self.cf3.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC1) and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC4(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D2_DC4)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D2_DC5)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D2_DC3)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)

class D2_DC4(D2, NamedTuple('DC4', [('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC4({_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC4) and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC5(D2, NamedTuple('DC5', [('cf10', Any), ('cf11', Any), ('cf12', Any), ('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC5({_dafny.string_of(self.cf10)}, {_dafny.string_of(self.cf11)}, {_dafny.string_of(self.cf12)}, {_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC5) and self.cf10 == __o.cf10 and self.cf11 == __o.cf11 and self.cf12 == __o.cf12 and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC3(D2, NamedTuple('DC3', [('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC3({_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC3) and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC6(D2, NamedTuple('DC6', [('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC8(False, int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D3_DC8)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D3_DC7)

class D3_DC8(D3, NamedTuple('DC8', [('cf16', Any), ('cf17', Any), ('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC8({_dafny.string_of(self.cf16)}, {_dafny.string_of(self.cf17)}, {_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC8) and self.cf16 == __o.cf16 and self.cf17 == __o.cf17 and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC7(D3, NamedTuple('DC7', [('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC7({_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC7) and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D4_DC9)

class D4_DC9(D4, NamedTuple('DC9', [('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC9({_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC9) and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC11(D2.default()(), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D5_DC11)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D5_DC10)

class D5_DC11(D5, NamedTuple('DC11', [('cf21', Any), ('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC11({_dafny.string_of(self.cf21)}, {_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC11) and self.cf21 == __o.cf21 and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC10(D5, NamedTuple('DC10', [('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC10({_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC10) and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC13(_dafny.Array(None, 0), int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D6_DC13)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D6_DC12)

class D6_DC13(D6, NamedTuple('DC13', [('cf24', Any), ('cf25', Any), ('cf26', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC13({_dafny.string_of(self.cf24)}, {_dafny.string_of(self.cf25)}, {_dafny.string_of(self.cf26)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC13) and self.cf24 == __o.cf24 and self.cf25 == __o.cf25 and self.cf26 == __o.cf26
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC12(D6, NamedTuple('DC12', [('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC12({_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC12) and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC15(_dafny.CodePoint('D'), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D7_DC15)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D7_DC16)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D7_DC14)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D7_DC17)

class D7_DC15(D7, NamedTuple('DC15', [('cf28', Any), ('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC15({_dafny.string_of(self.cf28)}, {_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC15) and self.cf28 == __o.cf28 and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC16(D7, NamedTuple('DC16', [('cf30', Any), ('cf31', Any), ('cf32', Any), ('cf33', Any), ('cf34', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC16({_dafny.string_of(self.cf30)}, {_dafny.string_of(self.cf31)}, {_dafny.string_of(self.cf32)}, {_dafny.string_of(self.cf33)}, {_dafny.string_of(self.cf34)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC16) and self.cf30 == __o.cf30 and self.cf31 == __o.cf31 and self.cf32 == __o.cf32 and self.cf33 == __o.cf33 and self.cf34 == __o.cf34
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC14(D7, NamedTuple('DC14', [('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC14({_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC14) and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC17(D7, NamedTuple('DC17', [('cf35', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC17({_dafny.string_of(self.cf35)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC17) and self.cf35 == __o.cf35
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC19(False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D8_DC19)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D8_DC18)

class D8_DC19(D8, NamedTuple('DC19', [('cf37', Any), ('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC19({_dafny.string_of(self.cf37)}, {_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC19) and self.cf37 == __o.cf37 and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC18(D8, NamedTuple('DC18', [('cf36', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC18({_dafny.string_of(self.cf36)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC18) and self.cf36 == __o.cf36
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D9_DC20)

class D9_DC20(D9, NamedTuple('DC20', [('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC20({_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC20) and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: D10_DC22(int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D10_DC22)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D10_DC23)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D10_DC21)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D10_DC24)

class D10_DC22(D10, NamedTuple('DC22', [('cf41', Any), ('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC22({_dafny.string_of(self.cf41)}, {_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC22) and self.cf41 == __o.cf41 and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC23(D10, NamedTuple('DC23', [('cf43', Any), ('cf44', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC23({_dafny.string_of(self.cf43)}, {_dafny.string_of(self.cf44)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC23) and self.cf43 == __o.cf43 and self.cf44 == __o.cf44
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC21(D10, NamedTuple('DC21', [('cf40', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC21({_dafny.string_of(self.cf40)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC21) and self.cf40 == __o.cf40
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC24(D10, NamedTuple('DC24', [('cf45', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC24({_dafny.string_of(self.cf45)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC24) and self.cf45 == __o.cf45
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC26(False, False, int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D11_DC26)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D11_DC27)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D11_DC28)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D11_DC25)

class D11_DC26(D11, NamedTuple('DC26', [('cf47', Any), ('cf48', Any), ('cf49', Any), ('cf50', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC26({_dafny.string_of(self.cf47)}, {_dafny.string_of(self.cf48)}, {_dafny.string_of(self.cf49)}, {_dafny.string_of(self.cf50)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC26) and self.cf47 == __o.cf47 and self.cf48 == __o.cf48 and self.cf49 == __o.cf49 and self.cf50 == __o.cf50
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC27(D11, NamedTuple('DC27', [])):
    def __dafnystr__(self) -> str:
        return f'D11.DC27'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC27)
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC28(D11, NamedTuple('DC28', [])):
    def __dafnystr__(self) -> str:
        return f'D11.DC28'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC28)
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC25(D11, NamedTuple('DC25', [('cf46', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC25({_dafny.string_of(self.cf46)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC25) and self.cf46 == __o.cf46
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D12_DC29)

class D12_DC29(D12, NamedTuple('DC29', [('cf51', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC29({_dafny.string_of(self.cf51)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC29) and self.cf51 == __o.cf51
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: None
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D13_DC30)

class D13_DC30(D13, NamedTuple('DC30', [('cf52', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC30({_dafny.string_of(self.cf52)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC30) and self.cf52 == __o.cf52
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D14_DC31)

class D14_DC31(D14, NamedTuple('DC31', [('cf53', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC31({_dafny.string_of(self.cf53)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC31) and self.cf53 == __o.cf53
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: D15_DC33(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D15_DC33)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D15_DC34)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D15_DC32)

class D15_DC33(D15, NamedTuple('DC33', [('cf55', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC33({_dafny.string_of(self.cf55)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC33) and self.cf55 == __o.cf55
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC34(D15, NamedTuple('DC34', [])):
    def __dafnystr__(self) -> str:
        return f'D15.DC34'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC34)
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC32(D15, NamedTuple('DC32', [('cf54', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC32({_dafny.string_of(self.cf54)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC32) and self.cf54 == __o.cf54
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D16_DC35)

class D16_DC35(D16, NamedTuple('DC35', [('cf56', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC35({_dafny.string_of(self.cf56)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC35) and self.cf56 == __o.cf56
    def __hash__(self) -> int:
        return super().__hash__()


class D17:
    @classmethod
    def default(cls, ):
        return lambda: None
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D17_DC36)

class D17_DC36(D17, NamedTuple('DC36', [('cf57', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC36({_dafny.string_of(self.cf57)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC36) and self.cf57 == __o.cf57
    def __hash__(self) -> int:
        return super().__hash__()


class D18:
    @classmethod
    def default(cls, ):
        return lambda: D18_DC38(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D18_DC38)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D18_DC39)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D18_DC37)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D18_DC40)

class D18_DC38(D18, NamedTuple('DC38', [('cf59', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC38({_dafny.string_of(self.cf59)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC38) and self.cf59 == __o.cf59
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC39(D18, NamedTuple('DC39', [('cf60', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC39({_dafny.string_of(self.cf60)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC39) and self.cf60 == __o.cf60
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC37(D18, NamedTuple('DC37', [('cf58', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC37({_dafny.string_of(self.cf58)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC37) and self.cf58 == __o.cf58
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC40(D18, NamedTuple('DC40', [('cf61', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC40({_dafny.string_of(self.cf61)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC40) and self.cf61 == __o.cf61
    def __hash__(self) -> int:
        return super().__hash__()


class D19:
    @classmethod
    def default(cls, ):
        return lambda: D19_DC42(False, _dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC42(self) -> bool:
        return isinstance(self, D19_DC42)
    @property
    def is_DC43(self) -> bool:
        return isinstance(self, D19_DC43)
    @property
    def is_DC41(self) -> bool:
        return isinstance(self, D19_DC41)

class D19_DC42(D19, NamedTuple('DC42', [('cf63', Any), ('cf64', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC42({_dafny.string_of(self.cf63)}, {_dafny.string_of(self.cf64)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC42) and self.cf63 == __o.cf63 and self.cf64 == __o.cf64
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC43(D19, NamedTuple('DC43', [('cf65', Any), ('cf66', Any), ('cf67', Any), ('cf68', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC43({_dafny.string_of(self.cf65)}, {_dafny.string_of(self.cf66)}, {_dafny.string_of(self.cf67)}, {_dafny.string_of(self.cf68)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC43) and self.cf65 == __o.cf65 and self.cf66 == __o.cf66 and self.cf67 == __o.cf67 and self.cf68 == __o.cf68
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC41(D19, NamedTuple('DC41', [('cf62', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC41({_dafny.string_of(self.cf62)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC41) and self.cf62 == __o.cf62
    def __hash__(self) -> int:
        return super().__hash__()


class D20:
    @classmethod
    def default(cls, ):
        return lambda: D20_DC45(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC45(self) -> bool:
        return isinstance(self, D20_DC45)
    @property
    def is_DC44(self) -> bool:
        return isinstance(self, D20_DC44)
    @property
    def is_DC46(self) -> bool:
        return isinstance(self, D20_DC46)

class D20_DC45(D20, NamedTuple('DC45', [('cf70', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC45({_dafny.string_of(self.cf70)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC45) and self.cf70 == __o.cf70
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC44(D20, NamedTuple('DC44', [('cf69', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC44({_dafny.string_of(self.cf69)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC44) and self.cf69 == __o.cf69
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC46(D20, NamedTuple('DC46', [('cf71', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC46({_dafny.string_of(self.cf71)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC46) and self.cf71 == __o.cf71
    def __hash__(self) -> int:
        return super().__hash__()


class D21:
    @classmethod
    def default(cls, ):
        return lambda: D21_DC48()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC48(self) -> bool:
        return isinstance(self, D21_DC48)
    @property
    def is_DC49(self) -> bool:
        return isinstance(self, D21_DC49)
    @property
    def is_DC47(self) -> bool:
        return isinstance(self, D21_DC47)

class D21_DC48(D21, NamedTuple('DC48', [])):
    def __dafnystr__(self) -> str:
        return f'D21.DC48'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC48)
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC49(D21, NamedTuple('DC49', [('cf73', Any), ('cf74', Any), ('cf75', Any), ('cf76', Any), ('cf77', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC49({_dafny.string_of(self.cf73)}, {_dafny.string_of(self.cf74)}, {_dafny.string_of(self.cf75)}, {_dafny.string_of(self.cf76)}, {_dafny.string_of(self.cf77)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC49) and self.cf73 == __o.cf73 and self.cf74 == __o.cf74 and self.cf75 == __o.cf75 and self.cf76 == __o.cf76 and self.cf77 == __o.cf77
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC47(D21, NamedTuple('DC47', [('cf72', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC47({_dafny.string_of(self.cf72)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC47) and self.cf72 == __o.cf72
    def __hash__(self) -> int:
        return super().__hash__()


class D22:
    @classmethod
    def default(cls, ):
        return lambda: D22_DC51(False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC51(self) -> bool:
        return isinstance(self, D22_DC51)
    @property
    def is_DC52(self) -> bool:
        return isinstance(self, D22_DC52)
    @property
    def is_DC50(self) -> bool:
        return isinstance(self, D22_DC50)
    @property
    def is_DC53(self) -> bool:
        return isinstance(self, D22_DC53)

class D22_DC51(D22, NamedTuple('DC51', [('cf79', Any), ('cf80', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC51({_dafny.string_of(self.cf79)}, {_dafny.string_of(self.cf80)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC51) and self.cf79 == __o.cf79 and self.cf80 == __o.cf80
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC52(D22, NamedTuple('DC52', [('cf81', Any), ('cf82', Any), ('cf83', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC52({_dafny.string_of(self.cf81)}, {_dafny.string_of(self.cf82)}, {_dafny.string_of(self.cf83)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC52) and self.cf81 == __o.cf81 and self.cf82 == __o.cf82 and self.cf83 == __o.cf83
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC50(D22, NamedTuple('DC50', [('cf78', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC50({_dafny.string_of(self.cf78)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC50) and self.cf78 == __o.cf78
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC53(D22, NamedTuple('DC53', [('cf84', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC53({_dafny.string_of(self.cf84)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC53) and self.cf84 == __o.cf84
    def __hash__(self) -> int:
        return super().__hash__()


class D23:
    @classmethod
    def default(cls, ):
        return lambda: D23_DC55(int(0), False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC55(self) -> bool:
        return isinstance(self, D23_DC55)
    @property
    def is_DC54(self) -> bool:
        return isinstance(self, D23_DC54)

class D23_DC55(D23, NamedTuple('DC55', [('cf86', Any), ('cf87', Any), ('cf88', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC55({_dafny.string_of(self.cf86)}, {_dafny.string_of(self.cf87)}, {_dafny.string_of(self.cf88)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC55) and self.cf86 == __o.cf86 and self.cf87 == __o.cf87 and self.cf88 == __o.cf88
    def __hash__(self) -> int:
        return super().__hash__()

class D23_DC54(D23, NamedTuple('DC54', [('cf85', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC54({_dafny.string_of(self.cf85)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC54) and self.cf85 == __o.cf85
    def __hash__(self) -> int:
        return super().__hash__()


class D24:
    @classmethod
    def default(cls, ):
        return lambda: D24_DC57(int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC57(self) -> bool:
        return isinstance(self, D24_DC57)
    @property
    def is_DC56(self) -> bool:
        return isinstance(self, D24_DC56)

class D24_DC57(D24, NamedTuple('DC57', [('cf90', Any), ('cf91', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC57({_dafny.string_of(self.cf90)}, {_dafny.string_of(self.cf91)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC57) and self.cf90 == __o.cf90 and self.cf91 == __o.cf91
    def __hash__(self) -> int:
        return super().__hash__()

class D24_DC56(D24, NamedTuple('DC56', [('cf89', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC56({_dafny.string_of(self.cf89)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC56) and self.cf89 == __o.cf89
    def __hash__(self) -> int:
        return super().__hash__()


class D25:
    @classmethod
    def default(cls, ):
        return lambda: D25_DC59(False, False, None)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC59(self) -> bool:
        return isinstance(self, D25_DC59)
    @property
    def is_DC58(self) -> bool:
        return isinstance(self, D25_DC58)

class D25_DC59(D25, NamedTuple('DC59', [('cf93', Any), ('cf94', Any), ('cf95', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC59({_dafny.string_of(self.cf93)}, {_dafny.string_of(self.cf94)}, {_dafny.string_of(self.cf95)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC59) and self.cf93 == __o.cf93 and self.cf94 == __o.cf94 and self.cf95 == __o.cf95
    def __hash__(self) -> int:
        return super().__hash__()

class D25_DC58(D25, NamedTuple('DC58', [('cf92', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC58({_dafny.string_of(self.cf92)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC58) and self.cf92 == __o.cf92
    def __hash__(self) -> int:
        return super().__hash__()


class D26:
    @classmethod
    def default(cls, ):
        return lambda: D26_DC61(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC61(self) -> bool:
        return isinstance(self, D26_DC61)
    @property
    def is_DC62(self) -> bool:
        return isinstance(self, D26_DC62)
    @property
    def is_DC63(self) -> bool:
        return isinstance(self, D26_DC63)
    @property
    def is_DC60(self) -> bool:
        return isinstance(self, D26_DC60)

class D26_DC61(D26, NamedTuple('DC61', [('cf97', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC61({_dafny.string_of(self.cf97)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC61) and self.cf97 == __o.cf97
    def __hash__(self) -> int:
        return super().__hash__()

class D26_DC62(D26, NamedTuple('DC62', [('cf98', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC62({_dafny.string_of(self.cf98)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC62) and self.cf98 == __o.cf98
    def __hash__(self) -> int:
        return super().__hash__()

class D26_DC63(D26, NamedTuple('DC63', [('cf99', Any), ('cf100', Any), ('cf101', Any), ('cf102', Any), ('cf103', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC63({_dafny.string_of(self.cf99)}, {_dafny.string_of(self.cf100)}, {_dafny.string_of(self.cf101)}, {_dafny.string_of(self.cf102)}, {_dafny.string_of(self.cf103)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC63) and self.cf99 == __o.cf99 and self.cf100 == __o.cf100 and self.cf101 == __o.cf101 and self.cf102 == __o.cf102 and self.cf103 == __o.cf103
    def __hash__(self) -> int:
        return super().__hash__()

class D26_DC60(D26, NamedTuple('DC60', [('cf96', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC60({_dafny.string_of(self.cf96)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC60) and self.cf96 == __o.cf96
    def __hash__(self) -> int:
        return super().__hash__()


class D27:
    @classmethod
    def default(cls, ):
        return lambda: D27_DC65(False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC65(self) -> bool:
        return isinstance(self, D27_DC65)
    @property
    def is_DC64(self) -> bool:
        return isinstance(self, D27_DC64)

class D27_DC65(D27, NamedTuple('DC65', [('cf105', Any), ('cf106', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC65({_dafny.string_of(self.cf105)}, {_dafny.string_of(self.cf106)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC65) and self.cf105 == __o.cf105 and self.cf106 == __o.cf106
    def __hash__(self) -> int:
        return super().__hash__()

class D27_DC64(D27, NamedTuple('DC64', [('cf104', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC64({_dafny.string_of(self.cf104)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC64) and self.cf104 == __o.cf104
    def __hash__(self) -> int:
        return super().__hash__()


class D28:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC66(self) -> bool:
        return isinstance(self, D28_DC66)

class D28_DC66(D28, NamedTuple('DC66', [('cf107', Any)])):
    def __dafnystr__(self) -> str:
        return f'D28.DC66({_dafny.string_of(self.cf107)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D28_DC66) and self.cf107 == __o.cf107
    def __hash__(self) -> int:
        return super().__hash__()


class D29:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC67(self) -> bool:
        return isinstance(self, D29_DC67)

class D29_DC67(D29, NamedTuple('DC67', [('cf108', Any)])):
    def __dafnystr__(self) -> str:
        return f'D29.DC67({_dafny.string_of(self.cf108)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D29_DC67) and self.cf108 == __o.cf108
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass

class T1:
    pass
    @property
    def f24(self):
        return self._f24
    @f24.setter
    def f24(self, value):
        self._f24 = value

class T2:
    pass
    def m0(self, globalState):
        pass

    def m1(self, p0, globalState):
        pass


class T3:
    pass
    @property
    def f27(self):
        return self._f27
    @f27.setter
    def f27(self, value):
        self._f27 = value
    def m2(self, p0, p1, p2, globalState):
        pass

    def m3(self, p0, p1, p2, globalState):
        pass


class GlobalState:
    def  __init__(self):
        self.f2: _dafny.Array = _dafny.Array(None, 0)
        self.f7: _dafny.Array = _dafny.Array(None, 0)
        self.f10: bool = False
        self.f11: int = int(0)
        self.f16: int = int(0)
        self.f22: int = int(0)
        self._f0: bool = False
        self._f1: int = int(0)
        self._f3: bool = False
        self._f4: int = int(0)
        self._f5: int = int(0)
        self._f6: _dafny.Array = _dafny.Array(None, 0)
        self._f8: bool = False
        self._f9: _dafny.Map = _dafny.Map({})
        self._f12: bool = False
        self._f13: int = int(0)
        self._f14: int = int(0)
        self._f15: _dafny.Seq = _dafny.Seq({})
        self._f17: bool = False
        self._f18: bool = False
        self._f19: int = int(0)
        self._f20: _dafny.Array = _dafny.Array(None, 0)
        self._f21: int = int(0)
        self._f23: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22, f23):
        (self)._f0 = f0
        (self)._f1 = f1
        (self).f2 = f2
        (self)._f3 = f3
        (self)._f4 = f4
        (self)._f5 = f5
        (self)._f6 = f6
        (self).f7 = f7
        (self)._f8 = f8
        (self)._f9 = f9
        (self).f10 = f10
        (self).f11 = f11
        (self)._f12 = f12
        (self)._f13 = f13
        (self)._f14 = f14
        (self)._f15 = f15
        (self).f16 = f16
        (self)._f17 = f17
        (self)._f18 = f18
        (self)._f19 = f19
        (self)._f20 = f20
        (self)._f21 = f21
        (self).f22 = f22
        (self)._f23 = f23

    @property
    def f0(self):
        return self._f0
    @property
    def f1(self):
        return self._f1
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
    def f6(self):
        return self._f6
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
    def f13(self):
        return self._f13
    @property
    def f14(self):
        return self._f14
    @property
    def f15(self):
        return self._f15
    @property
    def f17(self):
        return self._f17
    @property
    def f18(self):
        return self._f18
    @property
    def f19(self):
        return self._f19
    @property
    def f20(self):
        return self._f20
    @property
    def f21(self):
        return self._f21
    @property
    def f23(self):
        return self._f23

class C0:
    def  __init__(self):
        self._f31: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f31):
        (self)._f31 = f31

    def fm22(self, p0, globalState):
        return False

    def fm23(self, p0, p1, p2, p3, globalState):
        return (0) - ((0) - ((0) - (((0) - ((953) + (473))) - (-380))))

    @property
    def f31(self):
        return self._f31

class C1(T1, T0):
    def  __init__(self):
        self._f24: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    @property
    def f24(self):
        return self._f24
    @f24.setter
    def f24(self, value):
        self._f24 = value
    def ctor__(self, f24):
        (self).f24 = f24

    def fm4(self, p0, p1, globalState):
        return D0_DC0(False, (796) - (487), default__.safeModuloInt(265, 290))

    def fm5(self, p0, p1, p2, globalState):
        return (-281) * (len((_dafny.Map({self.f24: self.f24})) | (_dafny.Map({self.f24: self.f24}))))

    def fm24(self, p0, p1, p2, p3, globalState):
        return self.f24

    def fm25(self, p0, p1, globalState):
        return (0) - (default__.safeModuloInt((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xnc")))), (D3_DC8(self.f24, 498, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jxrwpl"))))).cf18))


class C2(T0):
    def  __init__(self):
        self.f29: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f30: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    def ctor__(self, f29, f30):
        (self).f29 = f29
        (self)._f30 = f30

    def fm18(self, p0, globalState):
        def iife39_():
            coll25_ = _dafny.Map()
            compr_25_: _dafny.MultiSet
            for compr_25_ in (_dafny.Map({_dafny.MultiSet([823, -334]): True})).keys.Elements:
                d_443_v0_: _dafny.MultiSet = compr_25_
                if (d_443_v0_) in (_dafny.Map({_dafny.MultiSet([823, -334]): True})):
                    coll25_[d_443_v0_] = _dafny.SeqWithoutIsStrInference([True])
            return _dafny.Map(coll25_)
        return (_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([869 for d_441_i0_ in range(default__.abs(-626))])) + (_dafny.SeqWithoutIsStrInference([(D3_DC8(False, len(_dafny.SeqWithoutIsStrInference([False])), (0) - (len(_dafny.Map({_dafny.CodePoint('r'): _dafny.CodePoint('j')}))))).cf17 for d_442_i1_ in range(default__.abs(936))])))) not in (iife39_()
        )

    def fm19(self, p0, globalState):
        return True

    def m7(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: bool = False
        r2: str = _dafny.CodePoint('D')
        r1 = p1
        d_444_v0_: _dafny.Array
        nw66_ = _dafny.Array(False, 6)
        d_444_v0_ = nw66_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_444_v0_).length(0)):
            d_445_i0_: int = guard_loop_0_
            if (True) and (((0) <= (d_445_i0_)) and ((d_445_i0_) < ((d_444_v0_).length(0)))):
                (d_444_v0_)[(d_445_i0_)] = (len((_dafny.Set({_dafny.CodePoint('q')}) if p1 else _dafny.Set({_dafny.CodePoint('h')})))) <= (355)
        d_446_i1_: int
        d_446_i1_ = 0
        with _dafny.label("4"):
            while not (not(not(p1))) or (p1):
                with _dafny.c_label("4"):
                    if (d_446_i1_) >= (100):
                        raise _dafny.Break("4")
                    d_446_i1_ = (d_446_i1_) + (1)
                    d_447_v1_: _dafny.Array
                    nw67_ = _dafny.Array(int(0), 7)
                    d_447_v1_ = nw67_
                    index60_ = default__.safeIndex(76, (d_447_v1_).length(0))
                    (d_447_v1_)[index60_] = p2
                    index61_ = default__.safeIndex(76, (d_447_v1_).length(0))
                    (d_447_v1_)[index61_] = p2
                    d_448_v2_: str
                    d_448_v2_ = _dafny.CodePoint('w')
                    d_449_v3_: D2
                    d_449_v3_ = D2_DC6(D2_DC5(d_448_v2_, True, p2, (d_447_v1_)[default__.safeIndex(76, (d_447_v1_).length(0))]))
                    d_450_v4_: _dafny.Seq
                    d_450_v4_ = _dafny.SeqWithoutIsStrInference([p1])
                    d_451_v5_: _dafny.Set
                    d_451_v5_ = _dafny.Set({p1})
                    d_452_v6_: _dafny.Map
                    d_452_v6_ = _dafny.Map({len(d_450_v4_): len(d_451_v5_)})
                    pat_let_tv4_ = d_452_v6_
                    def iife40_(_pat_let7_0):
                        def iife41_(d_453_dt__update__tmp_h0_):
                            def iife42_(_pat_let8_0):
                                def iife43_(d_454_dt__update_hcf14_h0_):
                                    return D2_DC6(d_454_dt__update_hcf14_h0_)
                                return iife43_(_pat_let8_0)
                            return iife42_(D2_DC3(pat_let_tv4_))
                        return iife41_(_pat_let7_0)
                    d_449_v3_ = iife40_(d_449_v3_)
                    d_455_v7_: D2
                    d_455_v7_ = D2_DC5(d_448_v2_, default__.fm1(p1, not(default__.fm1(p1, p1, p2, globalState)), (d_447_v1_)[default__.safeIndex(76, (d_447_v1_).length(0))], globalState), -869, 488)
                    d_456_v8_: _dafny.MultiSet
                    d_456_v8_ = _dafny.MultiSet([d_452_v6_, d_452_v6_])
                    d_457_v9_: _dafny.Map
                    d_457_v9_ = _dafny.Map({True: p2})
                    pat_let_tv5_ = d_456_v8_
                    pat_let_tv6_ = d_457_v9_
                    pat_let_tv7_ = d_448_v2_
                    def iife44_(_pat_let9_0):
                        def iife45_(d_458_dt__update__tmp_h1_):
                            def iife46_(_pat_let10_0):
                                def iife47_(d_459_dt__update_hcf13_h0_):
                                    def iife48_(_pat_let11_0):
                                        def iife49_(d_460_dt__update_hcf10_h0_):
                                            return D2_DC5(d_460_dt__update_hcf10_h0_, (d_458_dt__update__tmp_h1_).cf11, (d_458_dt__update__tmp_h1_).cf12, d_459_dt__update_hcf13_h0_)
                                        return iife49_(_pat_let11_0)
                                    return iife48_(pat_let_tv7_)
                                return iife47_(_pat_let10_0)
                            return iife46_((0) - (((pat_let_tv5_).cardinality) * (len(pat_let_tv6_))))
                        return iife45_(_pat_let9_0)
                    source11_ = iife44_(d_455_v7_)
                    if source11_.is_DC4:
                        d_461___mcc_h0_ = source11_.cf9
                        d_462_cf9_ = d_461___mcc_h0_
                        (globalState).f22 = (d_447_v1_)[default__.safeIndex(76, (d_447_v1_).length(0))]
                        r0 = d_462_cf9_
                        d_463_v10_: _dafny.Map
                        d_463_v10_ = _dafny.Map({(0) - (p2): (self).fm19(_dafny.CodePoint('o'), globalState)})
                        d_464_v11_: _dafny.Seq
                        d_464_v11_ = _dafny.SeqWithoutIsStrInference([(0) - ((d_447_v1_)[default__.safeIndex(76, (d_447_v1_).length(0))]), default__.safeDivisionInt(len((p0).set(p1, d_462_cf9_)), (d_447_v1_)[default__.safeIndex(76, (d_447_v1_).length(0))]), default__.fm0(p2, p1, d_463_v10_, len(d_452_v6_), globalState)])
                        (globalState).f11 = (d_464_v11_)[default__.safeIndex(p2, len(d_464_v11_))]
                        d_465_v12_: _dafny.Array
                        nw68_ = _dafny.Array(_dafny.Array(None, 0), 9)
                        d_465_v12_ = nw68_
                        d_466_v13_: D2
                        d_466_v13_ = D2_DC5(d_448_v2_, p1, 351, len(self.f29))
                        rhs42_ = d_465_v12_
                        rhs43_ = not(d_462_cf9_)
                        rhs44_ = D2_DC6(d_466_v13_)
                        d_465_v12_ = rhs42_
                        d_462_cf9_ = rhs43_
                        d_449_v3_ = rhs44_
                    elif source11_.is_DC5:
                        d_467___mcc_h1_ = source11_.cf10
                        d_468___mcc_h2_ = source11_.cf11
                        d_469___mcc_h3_ = source11_.cf12
                        d_470___mcc_h4_ = source11_.cf13
                        d_471_cf13_ = d_470___mcc_h4_
                        d_472_cf12_ = d_469___mcc_h3_
                        d_473_cf11_ = d_468___mcc_h2_
                        d_474_cf10_ = d_467___mcc_h1_
                        d_447_v1_ = (d_447_v1_ if (d_450_v4_)[default__.safeIndex((d_447_v1_)[default__.safeIndex(76, (d_447_v1_).length(0))], len(d_450_v4_))] else d_447_v1_)
                        d_475_v14_: _dafny.Map
                        d_475_v14_ = _dafny.Map({d_474_cf10_: (d_447_v1_)[default__.safeIndex(76, (d_447_v1_).length(0))]})
                        d_476_v16_: _dafny.Seq
                        d_476_v16_ = _dafny.SeqWithoutIsStrInference([len(d_452_v6_), d_471_cf13_])
                        d_477_v17_: _dafny.Map
                        d_477_v17_ = _dafny.Map({(d_447_v1_)[default__.safeIndex(76, (d_447_v1_).length(0))]: d_473_cf11_})
                        d_478_v19_: _dafny.Array
                        nw69_ = _dafny.Array(None, 25)
                        nw69_[int(0)] = (_dafny.Map({d_448_v2_: d_472_cf12_})) | (d_475_v14_)
                        nw69_[int(1)] = d_475_v14_
                        nw69_[int(2)] = d_475_v14_
                        def iife50_():
                            coll26_ = _dafny.Map()
                            compr_26_: str
                            for compr_26_ in (default__.fm20(globalState)).Elements:
                                d_479_v15_: str = compr_26_
                                if (d_479_v15_) in (default__.fm20(globalState)):
                                    coll26_[d_479_v15_] = len(self.f29)
                            return _dafny.Map(coll26_)
                        nw69_[int(3)] = iife50_()
                        
                        nw69_[int(4)] = d_475_v14_
                        nw69_[int(5)] = (d_475_v14_) | (_dafny.Map({d_448_v2_: -713}))
                        nw69_[int(6)] = d_475_v14_
                        nw69_[int(7)] = (d_475_v14_).set(d_448_v2_, len(d_476_v16_))
                        nw69_[int(8)] = d_475_v14_
                        nw69_[int(9)] = d_475_v14_
                        nw69_[int(10)] = (d_475_v14_).set(d_474_cf10_, default__.fm0(d_472_cf12_, p1, d_477_v17_, d_472_cf12_, globalState))
                        nw69_[int(11)] = d_475_v14_
                        nw69_[int(12)] = d_475_v14_
                        nw69_[int(13)] = d_475_v14_
                        nw69_[int(14)] = d_475_v14_
                        nw69_[int(15)] = d_475_v14_
                        nw69_[int(16)] = (d_475_v14_).set(d_448_v2_, (0) - (d_471_cf13_))
                        nw69_[int(17)] = (d_475_v14_) | (d_475_v14_)
                        nw69_[int(18)] = d_475_v14_
                        nw69_[int(19)] = d_475_v14_
                        nw69_[int(20)] = d_475_v14_
                        def iife51_():
                            coll27_ = _dafny.Map()
                            compr_27_: str
                            for compr_27_ in (self.f29).Elements:
                                d_480_v18_: str = compr_27_
                                if (d_480_v18_) in (self.f29):
                                    coll27_[d_480_v18_] = p2
                            return _dafny.Map(coll27_)
                        nw69_[int(21)] = iife51_()
                        
                        nw69_[int(22)] = _dafny.Map({d_448_v2_: p2})
                        nw69_[int(23)] = (d_475_v14_).set(d_474_cf10_, p2)
                        nw69_[int(24)] = d_475_v14_
                        d_478_v19_ = nw69_
                        index62_ = default__.safeIndex(22, (d_478_v19_).length(0))
                        (d_478_v19_)[index62_] = d_475_v14_
                        index63_ = default__.safeIndex(22, (d_478_v19_).length(0))
                        (d_478_v19_)[index63_] = d_475_v14_
                        d_481_v20_: _dafny.MultiSet
                        d_481_v20_ = _dafny.MultiSet([p1, (p1) or (d_473_cf11_), True, ((self).f30) != (self.f29), p1])
                        d_481_v20_ = (_dafny.MultiSet((d_450_v4_) + (d_450_v4_))).intersection(d_481_v20_)
                        d_482_v21_: _dafny.Set
                        d_482_v21_ = _dafny.Set({default__.fm0(len((self).f30), d_473_cf11_, d_477_v17_, d_472_cf12_, globalState)})
                        (globalState).f11 = (0) - (len(d_482_v21_))
                    elif source11_.is_DC3:
                        d_483___mcc_h5_ = source11_.cf8
                        d_484_cf8_ = d_483___mcc_h5_
                        (globalState).f22 = default__.safeDivisionInt(629, p2)
                        (globalState).f10 = p1
                        d_448_v2_ = default__.fm21(self.f29, p1, (p2) - (187), ((d_447_v1_)[default__.safeIndex(76, (d_447_v1_).length(0))]) + ((0) - ((d_447_v1_)[default__.safeIndex(76, (d_447_v1_).length(0))])), globalState)
                        (globalState).f11 = (d_447_v1_)[default__.safeIndex(76, (d_447_v1_).length(0))]
                    elif True:
                        d_485___mcc_h6_ = source11_.cf14
                        d_486_cf14_ = d_485___mcc_h6_
                        d_457_v9_ = (d_457_v9_).set(True, (p2) + (p2))
                        index64_ = default__.safeIndex(92, (d_444_v0_).length(0))
                        (d_444_v0_)[index64_] = not((p1) == (p1))
                        d_487_v22_: _dafny.Seq
                        d_487_v22_ = _dafny.SeqWithoutIsStrInference([-384])
                        index65_ = default__.safeIndex(92, (d_444_v0_).length(0))
                        rhs45_ = p1
                        rhs46_ = (0) - (default__.fm0((_dafny.MultiSet(d_487_v22_)).cardinality, p1, _dafny.Map({(0) - (p2): p1}), p2, globalState))
                        lhs35_ = d_444_v0_
                        lhs36_ = default__.safeIndex(92, (d_444_v0_).length(0))
                        lhs37_ = globalState
                        lhs35_[lhs36_] = rhs45_
                        lhs37_.f22 = rhs46_
                        index66_ = default__.safeIndex(76, (d_447_v1_).length(0))
                        (d_447_v1_)[index66_] = (p2) * ((_dafny.MultiSet(d_450_v4_)).cardinality)
                        (globalState).f16 = (default__.safeModuloInt(296, len(d_452_v6_))) * (-790)
                    index67_ = default__.safeIndex(76, (d_447_v1_).length(0))
                    (d_447_v1_)[index67_] = default__.safeDivisionInt(p2, (d_447_v1_)[default__.safeIndex(76, (d_447_v1_).length(0))])
                    pass
            pass
        d_488_v23_: str
        d_488_v23_ = _dafny.CodePoint('o')
        r2 = d_488_v23_
        (globalState).f10 = p1
        if not(p1):
            d_489_v24_: _dafny.Array
            nw70_ = _dafny.Array(None, 22)
            nw70_[int(0)] = d_444_v0_
            nw70_[int(1)] = d_444_v0_
            nw70_[int(2)] = d_444_v0_
            nw70_[int(3)] = d_444_v0_
            nw70_[int(4)] = d_444_v0_
            nw70_[int(5)] = d_444_v0_
            nw70_[int(6)] = d_444_v0_
            nw70_[int(7)] = d_444_v0_
            nw70_[int(8)] = d_444_v0_
            nw70_[int(9)] = d_444_v0_
            nw70_[int(10)] = d_444_v0_
            nw70_[int(11)] = d_444_v0_
            nw70_[int(12)] = d_444_v0_
            nw70_[int(13)] = d_444_v0_
            nw70_[int(14)] = d_444_v0_
            nw70_[int(15)] = d_444_v0_
            nw70_[int(16)] = d_444_v0_
            nw70_[int(17)] = d_444_v0_
            nw70_[int(18)] = d_444_v0_
            nw70_[int(19)] = d_444_v0_
            nw70_[int(20)] = d_444_v0_
            nw70_[int(21)] = d_444_v0_
            d_489_v24_ = nw70_
            source12_ = d_489_v24_
            d_490___mcc_h7_ = source12_
            d_491_cf19_ = d_490___mcc_h7_
            d_492_v25_: _dafny.Set
            d_492_v25_ = _dafny.Set({p1, p1})
            d_493_v26_: _dafny.Set
            d_493_v26_ = _dafny.Set({p1, p1, (p2) < ((D0_DC0(p1, p2, len(d_492_v25_))).cf2)})
            d_493_v26_ = _dafny.Set({False})
            (globalState).f16 = default__.safeModuloInt(856, 298)
            d_494_v27_: C0
            nw71_ = C0()
            nw71_.ctor__(self.f29)
            d_494_v27_ = nw71_
            (globalState).f10 = p1
            d_495_v28_: _dafny.Array
            nw72_ = _dafny.Array(_dafny.Map({}), 26)
            d_495_v28_ = nw72_
            d_496_v29_: _dafny.Map
            d_496_v29_ = _dafny.Map({not(p1): p2})
            d_497_v30_: _dafny.Map
            d_497_v30_ = _dafny.Map({d_495_v28_: default__.safeDivisionInt(len(d_496_v29_), p2)})
            d_497_v30_ = (d_497_v30_).set(d_495_v28_, p2)
            (globalState).f22 = p2
            d_498_v31_: _dafny.Array
            d_498_v31_ = d_489_v24_
            d_498_v31_ = d_498_v31_
            d_499_v32_: _dafny.Map
            d_499_v32_ = _dafny.Map({p2: (len(_dafny.Map({671: p2}))) > (p2)})
            d_499_v32_ = d_499_v32_
        elif True:
            r0 = default__.fm1(p1, p1, p2, globalState)
            d_500_v33_: D1
            d_500_v33_ = D1_DC1((self).f30)
            d_501_v34_: _dafny.Map
            d_501_v34_ = _dafny.Map({len(self.f29): d_500_v33_})
            d_502_v35_: _dafny.Map
            d_502_v35_ = _dafny.Map({True: _dafny.SeqWithoutIsStrInference([d_488_v23_ for d_503_i2_ in range(default__.abs(622))])})
            d_504_v36_: _dafny.Map
            d_504_v36_ = _dafny.Map({d_488_v23_: p2})
            d_505_v37_: _dafny.MultiSet
            d_505_v37_ = _dafny.MultiSet([len(((d_502_v35_)[p1] if (p1) in (d_502_v35_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mr")))), p2, len(d_504_v36_), p2])
            d_501_v34_ = (d_501_v34_) | (_dafny.Map({(d_505_v37_).cardinality: d_500_v33_}))
            d_506_v38_: _dafny.Map
            d_506_v38_ = _dafny.Map({p2: p1})
            d_507_v39_: T1
            nw73_ = C1()
            nw73_.ctor__(((d_506_v38_)[p2] if (p2) in (d_506_v38_) else p1))
            d_507_v39_ = nw73_
            d_508_v40_: _dafny.Map
            d_508_v40_ = _dafny.Map({d_507_v39_: p2})
            d_509_v41_: _dafny.Map
            d_509_v41_ = _dafny.Map({len(d_508_v40_): p2})
            d_510_v42_: D0
            d_510_v42_ = D0_DC0(d_507_v39_.f24, len(_dafny.SeqWithoutIsStrInference([d_488_v23_ for d_511_i3_ in range(default__.abs(86))])), p2)
            d_512_v43_: _dafny.MultiSet
            d_512_v43_ = _dafny.MultiSet([d_507_v39_.f24])
            d_513_v44_: _dafny.Map
            d_513_v44_ = _dafny.Map({d_507_v39_.f24: _dafny.SeqWithoutIsStrInference([d_510_v42_ for d_514_i4_ in range(default__.abs(-430))])})
            d_515_v45_: _dafny.Array
            nw74_ = _dafny.Array(None, 17)
            nw74_[int(0)] = p2
            nw74_[int(1)] = p2
            nw74_[int(2)] = p2
            nw74_[int(3)] = p2
            nw74_[int(4)] = len((d_509_v41_) | (d_509_v41_))
            nw74_[int(5)] = p2
            nw74_[int(6)] = p2
            nw74_[int(7)] = (p2) + (p2)
            nw74_[int(8)] = (d_510_v42_).cf2
            nw74_[int(9)] = (0) - ((0) - ((p2) * (p2)))
            nw74_[int(10)] = (p2) - (825)
            nw74_[int(11)] = p2
            nw74_[int(12)] = 25
            nw74_[int(13)] = ((d_512_v43_)[(self).fm19(d_488_v23_, globalState)] if ((self).fm19(d_488_v23_, globalState)) in (d_512_v43_) else (0) - (p2))
            nw74_[int(14)] = p2
            nw74_[int(15)] = (p2) - ((d_505_v37_).cardinality)
            nw74_[int(16)] = (d_507_v39_).fm5(d_513_v44_, p1, d_507_v39_.f24, globalState)
            d_515_v45_ = nw74_
            d_516_v46_: _dafny.Seq
            d_516_v46_ = _dafny.SeqWithoutIsStrInference([(0) - (p2), p2, (d_512_v43_).cardinality, p2, 913])
            index68_ = default__.safeIndex(256, (d_515_v45_).length(0))
            (d_515_v45_)[index68_] = len(d_516_v46_)
            index69_ = default__.safeIndex(256, (d_515_v45_).length(0))
            (d_515_v45_)[index69_] = (d_516_v46_)[default__.safeIndex(len((self).f30), len(d_516_v46_))]
            d_517_v47_: D3
            d_517_v47_ = D3_DC7(_dafny.SeqWithoutIsStrInference([p1]))
            d_518_v48_: _dafny.Map
            d_518_v48_ = _dafny.Map({d_517_v47_: False})
            d_519_v49_: _dafny.Seq
            d_519_v49_ = _dafny.SeqWithoutIsStrInference([d_507_v39_.f24, p1])
            d_518_v48_ = ((_dafny.Map({D3_DC7(d_519_v49_): not(d_507_v39_.f24)})) | (d_518_v48_)).set(d_517_v47_, not(d_507_v39_.f24))
            (globalState).f10 = p1
        r0 = p1
        r1 = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "auktfj"))) != (((self).f30) + ((self).f30))
        r2 = d_488_v23_
        return r0, r1, r2

    @property
    def f30(self):
        return self._f30

class C3(T0, T3, T2, T1):
    def  __init__(self):
        self._f24: bool = False
        self._f27: str = _dafny.CodePoint('D')
        self._f25: bool = False
        self._f26: bool = False
        self._f32: bool = False
        self._f33: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f24(self):
        return self._f24
    @f24.setter
    def f24(self, value):
        self._f24 = value
    @property
    def f27(self):
        return self._f27
    @f27.setter
    def f27(self, value):
        self._f27 = value
    @property
    def f25(self):
        return self._f25
    @property
    def f26(self):
        return self._f26
    def ctor__(self, f32, f33, f26, f27, f25, f24):
        (self)._f32 = f32
        (self)._f33 = f33
        (self)._f26 = f26
        (self).f27 = f27
        (self)._f25 = f25
        (self).f24 = f24

    def fm7(self, p0, globalState):
        return ((self).f33) < ((self).f33)

    def fm6(self, p0, globalState):
        return (len(_dafny.Map({len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([self.f27 for d_520_i0_ in range(default__.abs(421))])): self.f24})): (self).f25}))) * ((-966) - (len(_dafny.Set({(self).f26}))))

    def fm4(self, p0, p1, globalState):
        return D0_DC0(not((len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([-18])): (self).f26}))) != ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f25, True, False]))).cardinality)), 553, (0) - (((0) - (-530)) * ((0) - (len((self).f33)))))

    def fm5(self, p0, p1, p2, globalState):
        return ((0) - (len((_dafny.Set({(self).f26, (self).f25})).intersection(_dafny.Set({(self).f26}))))) - (321)

    def fm39(self, p0, p1, globalState):
        return (self).f26

    def fm40(self, p0, p1, globalState):
        return _dafny.Map({(False) and ((self).f32): self.f24})

    def m2(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: bool = False
        r2: int = int(0)
        r3: int = int(0)
        d_521_v0_: _dafny.Array
        def lambda15_(d_522_i0_):
            return not((self).f32)

        init7_ = lambda15_
        nw75_ = _dafny.Array(None, 25)
        for i0_7_ in range(nw75_.length(0)):
            nw75_[i0_7_] = init7_(i0_7_)
        d_521_v0_ = nw75_
        index70_ = default__.safeIndex(214, (d_521_v0_).length(0))
        (d_521_v0_)[index70_] = (self).f32
        d_523_v1_: _dafny.MultiSet
        d_523_v1_ = _dafny.MultiSet([p0, p0, p0])
        d_524_v2_: _dafny.Array
        nw76_ = _dafny.Array(int(0), 8)
        d_524_v2_ = nw76_
        d_525_v3_: _dafny.Seq
        d_525_v3_ = _dafny.SeqWithoutIsStrInference([d_524_v2_])
        index71_ = default__.safeIndex(214, (d_521_v0_).length(0))
        (d_521_v0_)[index71_] = not(not((_dafny.MultiSet([len(d_525_v3_), (0) - (p0)])).issubset((d_523_v1_) - (d_523_v1_))))
        d_526_i1_: int
        d_526_i1_ = 0
        with _dafny.label("5"):
            while (d_521_v0_)[default__.safeIndex(214, (d_521_v0_).length(0))]:
                with _dafny.c_label("5"):
                    if (d_526_i1_) >= (100):
                        raise _dafny.Break("5")
                    d_526_i1_ = (d_526_i1_) + (1)
                    d_527_v4_: _dafny.Map
                    d_527_v4_ = _dafny.Map({p0: default__.fm1(self.f24, (self).f26, p0, globalState)})
                    d_528_v5_: _dafny.Map
                    d_528_v5_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([p0])): p0})
                    index72_ = default__.safeIndex(214, (d_521_v0_).length(0))
                    (d_521_v0_)[index72_] = not(((d_527_v4_)[len((d_528_v5_) | (d_528_v5_))] if (len((d_528_v5_) | (d_528_v5_))) in (d_527_v4_) else (d_521_v0_)[default__.safeIndex(214, (d_521_v0_).length(0))]))
                    d_529_v6_: C2
                    nw77_ = C2()
                    nw77_.ctor__(((self).f33) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xgggjekrm"))), ((self).f33) + (default__.fm41((self).f32, self.f27, not((self).f32), globalState)))
                    d_529_v6_ = nw77_
                    (globalState).f22 = (p0) - (p0)
                    (self).f27 = self.f27
                    pass
            pass
        d_530_v7_: _dafny.Seq
        d_530_v7_ = _dafny.SeqWithoutIsStrInference([p0])
        d_531_v8_: _dafny.Seq
        d_531_v8_ = _dafny.SeqWithoutIsStrInference([(default__.fm42(globalState)) == (d_530_v7_)])
        d_532_v9_: _dafny.Map
        d_532_v9_ = _dafny.Map({(self).f26: self.f27})
        d_533_v10_: _dafny.Map
        d_533_v10_ = _dafny.Map({326: _dafny.MultiSet([len(d_532_v9_)])})
        r0 = (d_531_v8_)[default__.safeIndex(default__.safeDivisionInt(p0, (((d_533_v10_)[p0] if (p0) in (d_533_v10_) else d_523_v1_)).cardinality), len(d_531_v8_))]
        (self).m9(globalState)
        index73_ = default__.safeIndex(611, (d_524_v2_).length(0))
        (d_524_v2_)[index73_] = p0
        index74_ = default__.safeIndex(611, (d_524_v2_).length(0))
        (d_524_v2_)[index74_] = (self).fm6(p1, globalState)
        d_534_v11_: _dafny.Map
        d_534_v11_ = _dafny.Map({d_524_v2_: False})
        hi1_ = ((d_524_v2_)[default__.safeIndex(611, (d_524_v2_).length(0))]) * (len(d_534_v11_))
        for d_535_i2_ in range((len((self).f33) if False else (0) - ((d_524_v2_)[default__.safeIndex(611, (d_524_v2_).length(0))])), hi1_):
            d_536_v12_: _dafny.MultiSet
            d_536_v12_ = _dafny.MultiSet([d_530_v7_, d_530_v7_])
            d_536_v12_ = ((d_536_v12_) | (d_536_v12_)).intersection((default__.fm43(globalState)) | (default__.fm43(globalState)))
            if (self).f32:
                d_537_v13_: _dafny.Map
                d_537_v13_ = _dafny.Map({d_531_v8_: len(_dafny.SeqWithoutIsStrInference([self.f27 for d_538_i3_ in range(default__.abs(572))]))})
                (globalState).f10 = (((0) - (len((self).f33))) - (d_535_i2_)) != (len(d_537_v13_))
                d_539_v14_: _dafny.Map
                d_539_v14_ = _dafny.Map({(75) >= (583): d_521_v0_})
                d_539_v14_ = (d_539_v14_).set(((d_521_v0_)[default__.safeIndex(214, (d_521_v0_).length(0))] if not((self).f25) else (d_521_v0_)[default__.safeIndex(214, (d_521_v0_).length(0))]), d_521_v0_)
                d_540_v15_: _dafny.Map
                d_540_v15_ = _dafny.Map({d_535_i2_: d_535_i2_})
                d_541_v16_: _dafny.Map
                d_541_v16_ = _dafny.Map({d_521_v0_: (d_530_v7_)[default__.safeIndex(((d_540_v15_)[-301] if (-301) in (d_540_v15_) else p0), len(d_530_v7_))]})
                index75_ = default__.safeIndex(214, (d_521_v0_).length(0))
                index76_ = default__.safeIndex(611, (d_524_v2_).length(0))
                rhs47_ = ((d_524_v2_)[default__.safeIndex(611, (d_524_v2_).length(0))]) <= ((d_524_v2_)[default__.safeIndex(611, (d_524_v2_).length(0))])
                rhs48_ = (p0) + (len(d_541_v16_))
                rhs49_ = ((self).f33) < ((self).f33)
                rhs50_ = default__.safeModuloInt(p0, (p0) - ((d_524_v2_)[default__.safeIndex(611, (d_524_v2_).length(0))]))
                lhs38_ = d_521_v0_
                lhs39_ = default__.safeIndex(214, (d_521_v0_).length(0))
                lhs40_ = globalState
                lhs41_ = self
                lhs42_ = d_524_v2_
                lhs43_ = default__.safeIndex(611, (d_524_v2_).length(0))
                lhs38_[lhs39_] = rhs47_
                lhs40_.f16 = rhs48_
                lhs41_.f24 = rhs49_
                lhs42_[lhs43_] = rhs50_
                (globalState).f10 = (d_521_v0_)[default__.safeIndex(214, (d_521_v0_).length(0))]
                (self).f27 = self.f27
            elif True:
                d_542_v17_: _dafny.Map
                d_542_v17_ = _dafny.Map({p2: (self).f33})
                d_542_v17_ = (d_542_v17_).set((self).f32, (self).f33)
                (globalState).f10 = (self).f32
                (globalState).f16 = d_535_i2_
                (globalState).f22 = (d_524_v2_)[default__.safeIndex(611, (d_524_v2_).length(0))]
                index77_ = default__.safeIndex(214, (d_521_v0_).length(0))
                (d_521_v0_)[index77_] = (self).f32
            d_543_v18_: _dafny.Map
            d_543_v18_ = _dafny.Map({(d_524_v2_)[default__.safeIndex(611, (d_524_v2_).length(0))]: (d_524_v2_)[default__.safeIndex(611, (d_524_v2_).length(0))]})
            d_544_v19_: _dafny.Map
            d_544_v19_ = _dafny.Map({d_535_i2_: d_543_v18_})
            d_545_v20_: _dafny.MultiSet
            d_545_v20_ = _dafny.MultiSet([(self).f26])
            d_546_v21_: _dafny.Set
            d_546_v21_ = _dafny.Set({True, (self).f25})
            r1 = ((default__.fm44(len(d_544_v19_), _dafny.SeqWithoutIsStrInference([p0, (d_524_v2_)[default__.safeIndex(611, (d_524_v2_).length(0))], (self).fm6((self).fm4(d_545_v20_, (self).f33, globalState), globalState), (d_524_v2_)[default__.safeIndex(611, (d_524_v2_).length(0))]]), globalState)) | (d_546_v21_)).isdisjoint(d_546_v21_)
            d_547_v22_: _dafny.Map
            d_547_v22_ = _dafny.Map({(d_521_v0_ if self.f24 else d_521_v0_): (self).f32})
            d_547_v22_ = (d_547_v22_ if True else d_547_v22_)
        r0 = (d_531_v8_) == (_dafny.SeqWithoutIsStrInference([p2, (d_521_v0_)[default__.safeIndex(214, (d_521_v0_).length(0))], p2, (d_521_v0_)[default__.safeIndex(214, (d_521_v0_).length(0))]]))
        d_548_v23_: D6
        d_548_v23_ = D6_DC12(d_530_v7_)
        pat_let_tv8_ = p2
        def lambda16_(source13_):
            if source13_.is_DC13:
                d_549___mcc_h0_ = source13_.cf24
                d_550___mcc_h1_ = source13_.cf25
                d_551___mcc_h2_ = source13_.cf26
                d_552_cf26_ = d_551___mcc_h2_
                d_553_cf25_ = d_550___mcc_h1_
                d_554_cf24_ = d_549___mcc_h0_
                return pat_let_tv8_
            elif True:
                d_555___mcc_h3_ = source13_.cf23
                d_556_cf23_ = d_555___mcc_h3_
                return (self).f26

        r1 = lambda16_(d_548_v23_)
        r2 = (0) - (p0)
        d_557_v25_: _dafny.Map
        d_557_v25_ = _dafny.Map({self.f27: p0})
        def iife52_():
            coll28_ = _dafny.Map()
            compr_28_: str
            for compr_28_ in (d_557_v25_).keys.Elements:
                d_558_v24_: str = compr_28_
                if (d_558_v24_) in (d_557_v25_):
                    coll28_[d_558_v24_] = p0
            return _dafny.Map(coll28_)
        r3 = (len(iife52_()
        )) + ((d_524_v2_)[default__.safeIndex(611, (d_524_v2_).length(0))])
        return r0, r1, r2, r3

    def m3(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: bool = False
        r2: _dafny.Map = _dafny.Map({})
        d_559_v0_: _dafny.Set
        d_559_v0_ = _dafny.Set({(self).f32})
        d_560_v1_: _dafny.Set
        d_560_v1_ = _dafny.Set({d_559_v0_, _dafny.Set({self.f24, (self).f25, (self).f25}), d_559_v0_, d_559_v0_})
        (globalState).f22 = len(d_560_v1_)
        d_561_v2_: _dafny.Map
        d_561_v2_ = _dafny.Map({p0: p0})
        d_562_v3_: D2
        d_562_v3_ = D2_DC6(D2_DC3(d_561_v2_))
        d_563_v4_: D5
        d_563_v4_ = D5_DC11(d_562_v3_, p1)
        d_564_v5_: _dafny.Map
        d_564_v5_ = _dafny.Map({p0: default__.fm0(p0, (d_563_v4_).cf22, _dafny.Map({p0: p1}), p0, globalState)})
        d_565_v6_: D0
        d_565_v6_ = D0_DC0(True, p0, p0)
        d_566_v7_: _dafny.Map
        d_566_v7_ = _dafny.Map({(self).f32: _dafny.SeqWithoutIsStrInference([d_565_v6_])})
        d_561_v2_ = (d_561_v2_).set((self).fm5(d_566_v7_, (self).f32, (self).f25, globalState), p0)
        d_567_v8_: _dafny.Seq
        d_567_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ilp"))
        d_567_v8_ = (d_567_v8_) + ((d_567_v8_) + ((self).f33))
        d_568_v9_: _dafny.Array
        def lambda17_(d_569_p1_):
            def lambda18_(d_570_i1_):
                return (_dafny.Set({D7_DC15(self.f27, d_569_p1_)})) - (_dafny.Set({D7_DC15(self.f27, (self).f26)}))

            return lambda18_

        init8_ = lambda17_(p1)
        nw78_ = _dafny.Array(None, 6)
        for i0_8_ in range(nw78_.length(0)):
            nw78_[i0_8_] = init8_(i0_8_)
        d_568_v9_ = nw78_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_568_v9_).length(0)):
            d_571_i0_: int = guard_loop_1_
            if (True) and (((0) <= (d_571_i0_)) and ((d_571_i0_) < ((d_568_v9_).length(0)))):
                def iife53_():
                    coll29_ = _dafny.Set()
                    compr_29_: D7
                    for compr_29_ in (_dafny.Map({D7_DC15(self.f27, (self).f25): p1})).keys.Elements:
                        d_572_v10_: D7 = compr_29_
                        if (d_572_v10_) in (_dafny.Map({D7_DC15(self.f27, (self).f25): p1})):
                            coll29_ = coll29_.union(_dafny.Set([d_572_v10_]))
                    return _dafny.Set(coll29_)
                (d_568_v9_)[(d_571_i0_)] = ((_dafny.Set({D7_DC15(self.f27, (self).f25), D7_DC15(((self).f33)[default__.safeIndex(p0, len((self).f33))], (self).f32), D7_DC15(self.f27, True), D7_DC15(self.f27, self.f24)})) - (_dafny.Set({D7_DC15(self.f27, (self).f25)})) if (self).f26 else (_dafny.Set({D7_DC15(self.f27, (self).f32), D7_DC15(self.f27, (self).f26)})) | (iife53_()
                ))
        d_573_v11_: _dafny.Map
        d_573_v11_ = _dafny.Map({default__.safeModuloInt(-495, p0): (self).f26})
        r0 = ((d_573_v11_)[-875] if (-875) in (d_573_v11_) else not((self).f32))
        d_574_v12_: _dafny.Array
        def lambda19_(d_575_p0_):
            def lambda20_(d_576_i3_):
                return (d_576_i3_) * ((0) - (d_575_p0_))

            return lambda20_

        init9_ = lambda19_(p0)
        nw79_ = _dafny.Array(None, 16)
        for i0_9_ in range(nw79_.length(0)):
            nw79_[i0_9_] = init9_(i0_9_)
        d_574_v12_ = nw79_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_574_v12_).length(0)):
            d_577_i2_: int = guard_loop_2_
            if (True) and (((0) <= (d_577_i2_)) and ((d_577_i2_) < ((d_574_v12_).length(0)))):
                (d_574_v12_)[(d_577_i2_)] = default__.safeModuloInt(d_577_i2_, p0)
        d_578_v13_: D2
        d_578_v13_ = D2_DC5(self.f27, p1, p0, (0) - (p0))
        def lambda21_(source14_):
            if source14_.is_DC4:
                d_579___mcc_h0_ = source14_.cf9
                d_580_cf9_ = d_579___mcc_h0_
                return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fetbv"))) < ((self).f33)
            elif source14_.is_DC5:
                d_581___mcc_h1_ = source14_.cf10
                d_582___mcc_h2_ = source14_.cf11
                d_583___mcc_h3_ = source14_.cf12
                d_584___mcc_h4_ = source14_.cf13
                d_585_cf13_ = d_584___mcc_h4_
                d_586_cf12_ = d_583___mcc_h3_
                d_587_cf11_ = d_582___mcc_h2_
                d_588_cf10_ = d_581___mcc_h1_
                return (self).f26
            elif source14_.is_DC3:
                d_589___mcc_h5_ = source14_.cf8
                d_590_cf8_ = d_589___mcc_h5_
                return (self).f32
            elif True:
                d_591___mcc_h6_ = source14_.cf14
                d_592_cf14_ = d_591___mcc_h6_
                return (self).f26

        r0 = lambda21_(d_578_v13_)
        r1 = (self).f25
        d_593_v14_: _dafny.Map
        d_593_v14_ = _dafny.Map({d_567_v8_: (p0) + (p0)})
        r2 = d_593_v14_
        return r0, r1, r2

    def m0(self, globalState):
        r0: bool = False
        r1: _dafny.Array = _dafny.Array(None, 0)
        d_594_v0_: C1
        nw80_ = C1()
        nw80_.ctor__((self).f26)
        d_594_v0_ = nw80_
        d_595_v1_: int
        d_595_v1_ = 638
        hi2_ = (0) - (d_595_v1_)
        for d_596_i0_ in range(d_595_v1_, hi2_):
            (self).f24 = (self).f26
            d_597_v2_: D7
            d_597_v2_ = D7_DC15(self.f27, self.f24)
            source15_ = d_597_v2_
            if source15_.is_DC15:
                d_598___mcc_h0_ = source15_.cf28
                d_599___mcc_h1_ = source15_.cf29
                d_600_cf29_ = d_599___mcc_h1_
                d_601_cf28_ = d_598___mcc_h0_
                d_602_v3_: _dafny.Map
                d_602_v3_ = _dafny.Map({self.f24: d_595_v1_})
                d_603_v4_: _dafny.Array
                nw81_ = _dafny.Array(None, 10)
                nw81_[int(0)] = ((self).f33) + ((self).f33)
                nw81_[int(1)] = ((self).f33 if (self).fm7(d_602_v3_, globalState) else (self).f33)
                nw81_[int(2)] = (self).f33
                nw81_[int(3)] = (self).f33
                nw81_[int(4)] = (self).f33
                nw81_[int(5)] = (self).f33
                nw81_[int(6)] = (self).f33
                nw81_[int(7)] = (self).f33
                nw81_[int(8)] = (self).f33
                nw81_[int(9)] = ((self).f33 if self.f24 else (self).f33)
                d_603_v4_ = nw81_
                index78_ = default__.safeIndex(34, (d_603_v4_).length(0))
                (d_603_v4_)[index78_] = (self).f33
                index79_ = default__.safeIndex(34, (d_603_v4_).length(0))
                (d_603_v4_)[index79_] = _dafny.SeqWithoutIsStrInference([self.f27, d_601_cf28_])
                d_604_v5_: _dafny.Map
                d_604_v5_ = _dafny.Map({(self).f26: (self).f26})
                d_604_v5_ = (d_604_v5_).set((self).f25, False)
                d_605_v6_: _dafny.MultiSet
                d_605_v6_ = _dafny.MultiSet([(self).f26, (self).f25])
                d_606_v7_: D0
                d_606_v7_ = D0_DC0((self).f25, d_596_i0_, (d_605_v6_).cardinality)
                d_607_v8_: _dafny.Map
                d_607_v8_ = _dafny.Map({default__.safeDivisionInt(len((self).f33), d_595_v1_): d_606_v7_})
                d_608_v9_: _dafny.Seq
                d_608_v9_ = _dafny.SeqWithoutIsStrInference([(self).f25])
                d_607_v8_ = (d_607_v8_).set(d_596_i0_, D0_DC0(not((self).f26), d_595_v1_, len(d_608_v9_)))
                d_609_v10_: _dafny.Array
                def lambda22_(d_610_i0_):
                    def lambda23_(d_611_i2_):
                        return (d_611_i2_) - (d_610_i0_)

                    return lambda23_

                init10_ = lambda22_(d_596_i0_)
                nw82_ = _dafny.Array(None, 8)
                for i0_10_ in range(nw82_.length(0)):
                    nw82_[i0_10_] = init10_(i0_10_)
                d_609_v10_ = nw82_
                d_612_v11_: _dafny.Set
                d_612_v11_ = _dafny.Set({_dafny.CodePoint('f')})
                d_613_v12_: _dafny.Map
                d_613_v12_ = _dafny.Map({len(d_612_v11_): d_595_v1_})
                d_614_v13_: _dafny.Seq
                d_614_v13_ = _dafny.SeqWithoutIsStrInference([len((self).f33), d_595_v1_, len((d_613_v12_).set(d_595_v1_, 202)), len(d_608_v9_)])
                d_615_v14_: D1
                d_615_v14_ = D1_DC2(d_609_v10_, d_601_cf28_, len(d_614_v13_), True)
                d_616_v15_: _dafny.Seq
                d_616_v15_ = _dafny.SeqWithoutIsStrInference([(d_615_v14_).cf6])
                d_617_v16_: _dafny.MultiSet
                d_617_v16_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_596_i0_ for d_618_i1_ in range(default__.abs(-894))]), (_dafny.SeqWithoutIsStrInference([d_595_v1_])) + (d_614_v13_), (d_616_v15_ if d_600_cf29_ else d_616_v15_)])
                d_619_v17_: _dafny.Seq
                d_619_v17_ = _dafny.SeqWithoutIsStrInference([d_616_v15_])
                d_617_v16_ = (d_617_v16_) - ((d_617_v16_) - (_dafny.MultiSet(d_619_v17_)))
            elif source15_.is_DC16:
                d_620___mcc_h2_ = source15_.cf30
                d_621___mcc_h3_ = source15_.cf31
                d_622___mcc_h4_ = source15_.cf32
                d_623___mcc_h5_ = source15_.cf33
                d_624___mcc_h6_ = source15_.cf34
                d_625_cf34_ = d_624___mcc_h6_
                d_626_cf33_ = d_623___mcc_h5_
                d_627_cf32_ = d_622___mcc_h4_
                d_628_cf31_ = d_621___mcc_h3_
                d_629_cf30_ = d_620___mcc_h2_
                d_630_v18_: _dafny.Seq
                d_630_v18_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "btfxa"))
                d_630_v18_ = ((d_630_v18_) + (_dafny.SeqWithoutIsStrInference([self.f27 for d_631_i3_ in range(default__.abs(514))])) if d_625_cf34_ else (self).f33)
                (globalState).f16 = d_596_i0_
                d_632_v19_: C2
                nw83_ = C2()
                nw83_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "skabkujph")), d_630_v18_)
                d_632_v19_ = nw83_
                d_633_v20_: _dafny.Map
                d_633_v20_ = _dafny.Map({d_627_cf32_: d_596_i0_})
                d_634_v21_: _dafny.Seq
                d_634_v21_ = _dafny.SeqWithoutIsStrInference([(self).f26])
                d_635_v22_: _dafny.Seq
                d_635_v22_ = _dafny.SeqWithoutIsStrInference([len(d_634_v21_)])
                def iife54_():
                    coll30_ = _dafny.Map()
                    compr_30_: int
                    for compr_30_ in _dafny.IntegerRange(217, 893):
                        d_636_v23_: int = compr_30_
                        if ((217) <= (d_636_v23_)) and ((d_636_v23_) < (893)):
                            coll30_[(d_636_v23_) - (d_628_cf31_)] = d_626_cf33_
                    return _dafny.Map(coll30_)
                d_633_v20_ = (d_633_v20_).set(d_629_cf30_, (len(d_635_v22_) if (self).f25 else default__.fm0(d_596_i0_, d_627_cf32_, iife54_()
                , len(_dafny.SeqWithoutIsStrInference([self.f27 for d_637_i4_ in range(default__.abs(112))])), globalState)))
            elif source15_.is_DC14:
                d_638___mcc_h7_ = source15_.cf27
                d_639_cf27_ = d_638___mcc_h7_
                d_640_v24_: _dafny.Seq
                d_640_v24_ = _dafny.SeqWithoutIsStrInference([True, (self).f32])
                d_641_v25_: D3
                d_641_v25_ = D3_DC7(d_640_v24_)
                (globalState).f11 = len(((_dafny.SeqWithoutIsStrInference([(self).f26, (self).f32, (self).f26])) + (d_640_v24_) if (d_595_v1_) >= (len((self).f33)) else (d_641_v25_).cf15))
                rhs51_ = default__.safeDivisionInt(d_595_v1_, d_595_v1_)
                rhs52_ = (0) - (default__.safeModuloInt((d_596_i0_) + (d_596_i0_), len((self).f33)))
                rhs53_ = not((d_596_i0_) <= (d_595_v1_))
                lhs44_ = globalState
                lhs45_ = globalState
                lhs46_ = globalState
                lhs44_.f22 = rhs51_
                lhs45_.f11 = rhs52_
                lhs46_.f10 = rhs53_
                d_642_v26_: _dafny.Array
                def lambda24_(d_643_v24_):
                    def lambda25_(d_644_i5_):
                        return d_643_v24_

                    return lambda25_

                init11_ = lambda24_(d_640_v24_)
                nw84_ = _dafny.Array(None, 1)
                for i0_11_ in range(nw84_.length(0)):
                    nw84_[i0_11_] = init11_(i0_11_)
                d_642_v26_ = nw84_
                d_642_v26_ = d_642_v26_
                d_645_v27_: _dafny.Array
                nw85_ = _dafny.Array(_dafny.Set({}), 16)
                d_645_v27_ = nw85_
                d_646_v28_: D7
                d_646_v28_ = D7_DC16((self).f25, d_596_i0_, (self).f25, (self).f26, self.f24)
                rhs54_ = d_596_i0_
                rhs55_ = d_645_v27_
                rhs56_ = d_646_v28_
                rhs57_ = (self).f26
                rhs58_ = (862) == ((len((self).f33)) * (d_595_v1_))
                lhs47_ = globalState
                lhs48_ = globalState
                lhs47_.f22 = rhs54_
                d_645_v27_ = rhs55_
                d_646_v28_ = rhs56_
                lhs48_.f10 = rhs57_
                r0 = rhs58_
            elif True:
                d_647___mcc_h8_ = source15_.cf35
                d_648_cf35_ = d_647___mcc_h8_
                r0 = self.f24
                d_649_v29_: _dafny.Set
                d_649_v29_ = _dafny.Set({self.f24, (self).f32, (self).f26, self.f24})
                d_650_v30_: _dafny.Array
                def lambda26_(d_651_i0_):
                    def lambda27_(d_652_i6_):
                        return default__.safeDivisionInt(d_652_i6_, d_651_i0_)

                    return lambda27_

                init12_ = lambda26_(d_596_i0_)
                nw86_ = _dafny.Array(None, 11)
                for i0_12_ in range(nw86_.length(0)):
                    nw86_[i0_12_] = init12_(i0_12_)
                d_650_v30_ = nw86_
                d_653_v31_: _dafny.MultiSet
                d_653_v31_ = _dafny.MultiSet([d_650_v30_, d_650_v30_, d_650_v30_, d_650_v30_, d_650_v30_])
                d_654_v32_: D2
                d_654_v32_ = D2_DC5(self.f27, (self).f25, d_596_i0_, d_595_v1_)
                d_655_v33_: _dafny.Map
                d_655_v33_ = _dafny.Map({(d_654_v32_).cf12: d_596_i0_})
                d_656_v34_: D2
                d_656_v34_ = D2_DC3(d_655_v33_)
                d_657_v35_: D2
                d_657_v35_ = D2_DC6(d_656_v34_)
                d_658_v36_: D5
                d_658_v36_ = D5_DC11(d_657_v35_, (self).f26)
                d_659_v37_: _dafny.Seq
                d_659_v37_ = _dafny.SeqWithoutIsStrInference([d_595_v1_])
                pat_let_tv9_ = d_657_v35_
                d_660_v38_: _dafny.Array
                nw87_ = _dafny.Array(None, 20)
                nw87_[int(0)] = self.f24
                nw87_[int(1)] = (_dafny.Set({(self).f32, (self).f32, (self).f25})) == (d_649_v29_)
                nw87_[int(2)] = (self).f26
                nw87_[int(3)] = (self).f32
                nw87_[int(4)] = default__.fm1((self).f32, not(self.f24), d_595_v1_, globalState)
                nw87_[int(5)] = (d_653_v31_).isdisjoint(d_653_v31_)
                nw87_[int(6)] = True
                nw87_[int(7)] = not((self).f26)
                nw87_[int(8)] = self.f24
                nw87_[int(9)] = self.f24
                nw87_[int(10)] = False
                nw87_[int(11)] = True
                nw87_[int(12)] = (self).f25
                nw87_[int(13)] = (self).f25
                def iife55_(_pat_let12_0):
                    def iife56_(d_661_dt__update__tmp_h0_):
                        def iife57_(_pat_let13_0):
                            def iife58_(d_662_dt__update_hcf21_h0_):
                                return D5_DC11(d_662_dt__update_hcf21_h0_, (d_661_dt__update__tmp_h0_).cf22)
                            return iife58_(_pat_let13_0)
                        return iife57_(pat_let_tv9_)
                    return iife56_(_pat_let12_0)
                nw87_[int(14)] = (iife55_(d_658_v36_)).cf22
                nw87_[int(15)] = not ((self).f32) or ((self).fm39(d_659_v37_, (self).fm39(d_659_v37_, (self).f32, globalState), globalState))
                nw87_[int(16)] = (self).f25
                nw87_[int(17)] = False
                nw87_[int(18)] = not ((self).f26) or ((self).f26)
                nw87_[int(19)] = self.f24
                d_660_v38_ = nw87_
                index80_ = default__.safeIndex(710, (d_660_v38_).length(0))
                (d_660_v38_)[index80_] = not(True)
                index81_ = default__.safeIndex(710, (d_660_v38_).length(0))
                (d_660_v38_)[index81_] = (d_594_v0_).fm24(d_596_i0_, (self).fm40((self).f26, d_659_v37_, globalState), default__.safeDivisionInt(d_596_i0_, -917), len(default__.fm41((self).f26, self.f27, not((self).fm39(d_659_v37_, (self).f32, globalState)), globalState)), globalState)
                (globalState).f7 = d_650_v30_
                pat_let_tv10_ = d_657_v35_
                def iife59_(_pat_let14_0):
                    def iife60_(d_663_dt__update__tmp_h1_):
                        def iife61_(_pat_let15_0):
                            def iife62_(d_664_dt__update_hcf21_h1_):
                                return D5_DC11(d_664_dt__update_hcf21_h1_, (d_663_dt__update__tmp_h1_).cf22)
                            return iife62_(_pat_let15_0)
                        return iife61_(pat_let_tv10_)
                    return iife60_(_pat_let14_0)
                d_658_v36_ = iife59_(d_658_v36_)
            d_665_v39_: _dafny.Map
            d_665_v39_ = _dafny.Map({len(default__.fm45(self.f24, globalState)): self.f27})
            d_666_v40_: _dafny.Seq
            d_666_v40_ = _dafny.SeqWithoutIsStrInference([d_665_v39_])
            d_666_v40_ = d_666_v40_
            d_667_v41_: _dafny.Seq
            d_667_v41_ = _dafny.SeqWithoutIsStrInference([d_595_v1_, d_596_i0_, d_595_v1_])
            if default__.fm1(((self).f32) and ((self).f32), (d_595_v1_) > (d_596_i0_), default__.safeModuloInt((d_667_v41_)[default__.safeIndex(d_596_i0_, len(d_667_v41_))], (0) - (-708)), globalState):
                d_668_v42_: _dafny.Array
                def lambda28_(d_669_i0_):
                    def lambda29_(d_670_i7_):
                        return default__.safeDivisionInt(d_670_i7_, d_669_i0_)

                    return lambda29_

                init13_ = lambda28_(d_596_i0_)
                nw88_ = _dafny.Array(None, 15)
                for i0_13_ in range(nw88_.length(0)):
                    nw88_[i0_13_] = init13_(i0_13_)
                d_668_v42_ = nw88_
                index82_ = default__.safeIndex(950, (d_668_v42_).length(0))
                (d_668_v42_)[index82_] = (d_596_i0_) + (d_596_i0_)
                d_671_v43_: D2
                d_671_v43_ = D2_DC5(_dafny.CodePoint('g'), (False) and (self.f24), d_595_v1_, (0) - ((d_667_v41_)[default__.safeIndex(len((self).f33), len(d_667_v41_))]))
                index83_ = default__.safeIndex(950, (d_668_v42_).length(0))
                rhs59_ = (982) - (default__.safeDivisionInt(d_596_i0_, d_595_v1_))
                rhs60_ = (0) - ((0) - (d_595_v1_))
                rhs61_ = default__.fm46(globalState)
                rhs62_ = (default__.safeModuloInt(d_596_i0_, d_595_v1_)) - (d_596_i0_)
                lhs49_ = d_668_v42_
                lhs50_ = default__.safeIndex(950, (d_668_v42_).length(0))
                lhs51_ = globalState
                lhs52_ = globalState
                lhs49_[lhs50_] = rhs59_
                lhs51_.f11 = rhs60_
                d_671_v43_ = rhs61_
                lhs52_.f11 = rhs62_
                d_672_v44_: _dafny.Array
                nw89_ = _dafny.Array(_dafny.Seq({}), 4)
                d_672_v44_ = nw89_
                index84_ = default__.safeIndex(2, (d_672_v44_).length(0))
                (d_672_v44_)[index84_] = d_667_v41_
                d_673_v45_: _dafny.Seq
                d_673_v45_ = _dafny.SeqWithoutIsStrInference([(self).f33, (self).f33, (self).f33, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "irfek")), (self).f33])
                index85_ = default__.safeIndex(2, (d_672_v44_).length(0))
                rhs63_ = (self).f32
                rhs64_ = _dafny.SeqWithoutIsStrInference([-300, d_595_v1_, (0) - ((d_668_v42_)[default__.safeIndex(950, (d_668_v42_).length(0))]), len(d_673_v45_), (d_668_v42_)[default__.safeIndex(950, (d_668_v42_).length(0))]])
                rhs65_ = default__.fm47(globalState)
                rhs66_ = self.f27
                rhs67_ = ((self).f25 if self.f24 else (d_595_v1_) <= (d_596_i0_))
                lhs53_ = globalState
                lhs54_ = d_672_v44_
                lhs55_ = default__.safeIndex(2, (d_672_v44_).length(0))
                lhs56_ = self
                lhs57_ = self
                lhs53_.f10 = rhs63_
                lhs54_[lhs55_] = rhs64_
                lhs56_.f27 = rhs65_
                lhs57_.f27 = rhs66_
                r0 = rhs67_
                d_674_v46_: _dafny.Seq
                d_674_v46_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dreorfuot"))
                d_674_v46_ = (default__.fm41(False, self.f27, self.f24, globalState)) + ((self).f33)
                d_675_v47_: D0
                d_675_v47_ = D0_DC0((self).f32, d_596_i0_, (d_668_v42_)[default__.safeIndex(950, (d_668_v42_).length(0))])
                d_676_v48_: _dafny.Seq
                d_676_v48_ = _dafny.SeqWithoutIsStrInference([d_675_v47_])
                d_677_v49_: _dafny.Map
                d_677_v49_ = _dafny.Map({False: d_676_v48_})
                (globalState).f11 = (((self).fm5(d_677_v49_, (self).f32, (self).f26, globalState)) + ((self).fm5(d_677_v49_, True, self.f24, globalState))) * ((d_668_v42_)[default__.safeIndex(950, (d_668_v42_).length(0))])
                def lambda30_(d_678_i8_):
                    return ((self).f26 if (self).f26 else (self).f26)

                init14_ = lambda30_
                nw90_ = _dafny.Array(None, 7)
                for i0_14_ in range(nw90_.length(0)):
                    nw90_[i0_14_] = init14_(i0_14_)
                r1 = nw90_
            elif True:
                d_679_v50_: _dafny.Array
                def lambda31_(d_680_i9_):
                    return self.f27

                init15_ = lambda31_
                nw91_ = _dafny.Array(None, 8)
                for i0_15_ in range(nw91_.length(0)):
                    nw91_[i0_15_] = init15_(i0_15_)
                d_679_v50_ = nw91_
                index86_ = default__.safeIndex(167, (d_679_v50_).length(0))
                (d_679_v50_)[index86_] = (default__.fm47(globalState) if (self).f32 else self.f27)
                index87_ = default__.safeIndex(167, (d_679_v50_).length(0))
                (d_679_v50_)[index87_] = _dafny.CodePoint('l')
                (globalState).f10 = not(not((self).f25))
                d_681_v51_: C2
                nw92_ = C2()
                nw92_.ctor__(_dafny.SeqWithoutIsStrInference([(D2_DC5((d_679_v50_)[default__.safeIndex(167, (d_679_v50_).length(0))], (self).f32, d_596_i0_, d_595_v1_)).cf10 for d_682_i10_ in range(default__.abs(81))]), (self).f33)
                d_681_v51_ = nw92_
                d_683_v52_: _dafny.Seq
                d_683_v52_ = _dafny.SeqWithoutIsStrInference([(self).f26])
                d_684_v53_: _dafny.Map
                d_684_v53_ = _dafny.Map({default__.fm1(not(not((self).f26)), (self).f25, d_596_i0_, globalState): (self).f26})
                d_685_v54_: _dafny.Map
                d_685_v54_ = _dafny.Map({(d_594_v0_).fm24(d_596_i0_, d_684_v53_, 720, d_596_i0_, globalState): d_596_i0_})
                d_686_v55_: _dafny.MultiSet
                d_686_v55_ = _dafny.MultiSet([_dafny.Map({(self).f25: (self).f32})])
                d_687_v56_: _dafny.Map
                d_687_v56_ = _dafny.Map({(self).f32: self.f27})
                d_688_v57_: T0
                nw93_ = C2()
                nw93_.ctor__(_dafny.SeqWithoutIsStrInference([(d_679_v50_)[default__.safeIndex(167, (d_679_v50_).length(0))] for d_689_i11_ in range(default__.abs(-118))]), _dafny.SeqWithoutIsStrInference([(d_679_v50_)[default__.safeIndex(167, (d_679_v50_).length(0))] for d_690_i12_ in range(default__.abs(266))]))
                d_688_v57_ = nw93_
                d_691_v58_: _dafny.Map
                d_691_v58_ = _dafny.Map({d_688_v57_: self.f24})
                d_692_v59_: _dafny.Array
                nw94_ = _dafny.Array(None, 28)
                nw94_[int(0)] = d_596_i0_
                nw94_[int(1)] = d_595_v1_
                nw94_[int(2)] = len(d_683_v52_)
                nw94_[int(3)] = ((d_685_v54_)[not((self).f25)] if (not((self).f25)) in (d_685_v54_) else d_595_v1_)
                nw94_[int(4)] = (0) - (d_596_i0_)
                nw94_[int(5)] = (d_686_v55_).cardinality
                nw94_[int(6)] = d_595_v1_
                nw94_[int(7)] = d_595_v1_
                nw94_[int(8)] = d_595_v1_
                nw94_[int(9)] = len(d_687_v56_)
                nw94_[int(10)] = d_596_i0_
                nw94_[int(11)] = d_596_i0_
                nw94_[int(12)] = (d_667_v41_)[default__.safeIndex(369, len(d_667_v41_))]
                nw94_[int(13)] = d_595_v1_
                nw94_[int(14)] = 159
                nw94_[int(15)] = d_595_v1_
                nw94_[int(16)] = len(_dafny.Map({d_596_i0_: d_595_v1_}))
                nw94_[int(17)] = d_595_v1_
                nw94_[int(18)] = len(d_691_v58_)
                nw94_[int(19)] = d_595_v1_
                nw94_[int(20)] = 307
                nw94_[int(21)] = (0) - (d_596_i0_)
                nw94_[int(22)] = d_596_i0_
                nw94_[int(23)] = d_595_v1_
                nw94_[int(24)] = d_595_v1_
                nw94_[int(25)] = d_595_v1_
                nw94_[int(26)] = d_595_v1_
                nw94_[int(27)] = d_595_v1_
                d_692_v59_ = nw94_
                d_693_v60_: _dafny.Seq
                d_693_v60_ = _dafny.SeqWithoutIsStrInference([d_692_v59_, d_692_v59_])
                d_694_v61_: _dafny.Seq
                d_694_v61_ = _dafny.SeqWithoutIsStrInference([(d_693_v60_)[default__.safeIndex(d_596_i0_, len(d_693_v60_))], d_692_v59_])
                (globalState).f7 = (d_694_v61_)[default__.safeIndex((0) - (d_595_v1_), len(d_694_v61_))]
                d_695_v62_: _dafny.Map
                d_695_v62_ = _dafny.Map({D7_DC17(D7_DC15(self.f27, (self).f32)): (d_683_v52_) + (d_683_v52_)})
                d_696_v63_: D7
                d_696_v63_ = D7_DC15(self.f27, (self).f25)
                d_697_v64_: D7
                d_697_v64_ = D7_DC17(d_696_v63_)
                d_695_v62_ = (d_695_v62_).set(d_697_v64_, d_683_v52_)
        (globalState).f10 = ((0) - ((0) - (d_595_v1_))) >= (d_595_v1_)
        (globalState).f22 = (600) + (d_595_v1_)
        d_698_i13_: int
        d_698_i13_ = 0
        with _dafny.label("6"):
            while (self).f32:
                with _dafny.c_label("6"):
                    if (d_698_i13_) >= (100):
                        raise _dafny.Break("6")
                    d_698_i13_ = (d_698_i13_) + (1)
                    (globalState).f22 = d_595_v1_
                    d_699_v65_: _dafny.MultiSet
                    d_699_v65_ = _dafny.MultiSet([(self).f32])
                    d_700_v66_: _dafny.Seq
                    d_700_v66_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y"))
                    d_701_v67_: _dafny.Array
                    nw95_ = _dafny.Array(_dafny.Array(None, 0), 5)
                    d_701_v67_ = nw95_
                    d_702_v68_: _dafny.Array
                    def lambda32_(d_703_v1_):
                        def lambda33_(d_704_i14_):
                            return default__.safeDivisionInt(d_704_i14_, d_703_v1_)

                        return lambda33_

                    init16_ = lambda32_(d_595_v1_)
                    nw96_ = _dafny.Array(None, 9)
                    for i0_16_ in range(nw96_.length(0)):
                        nw96_[i0_16_] = init16_(i0_16_)
                    d_702_v68_ = nw96_
                    index88_ = default__.safeIndex(578, (d_701_v67_).length(0))
                    (d_701_v67_)[index88_] = d_702_v68_
                    d_705_v69_: _dafny.Array
                    nw97_ = _dafny.Array(False, 23)
                    d_705_v69_ = nw97_
                    d_706_v70_: _dafny.Map
                    d_706_v70_ = _dafny.Map({d_705_v69_: self.f27})
                    d_707_v71_: D2
                    d_707_v71_ = D2_DC5(self.f27, (self).f32, (0) - (len(d_706_v70_)), d_595_v1_)
                    d_708_v72_: _dafny.MultiSet
                    d_708_v72_ = _dafny.MultiSet([d_699_v65_])
                    index89_ = default__.safeIndex(578, (d_701_v67_).length(0))
                    rhs68_ = (d_699_v65_).set((d_707_v71_).cf11, default__.abs(-635))
                    rhs69_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ep"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gdeamwx")))
                    rhs70_ = (_dafny.Set({(d_708_v72_).cardinality})).ispropersubset(_dafny.Set({d_595_v1_}))
                    rhs71_ = d_702_v68_
                    rhs72_ = (self).f32
                    lhs58_ = globalState
                    lhs59_ = d_701_v67_
                    lhs60_ = default__.safeIndex(578, (d_701_v67_).length(0))
                    lhs61_ = self
                    d_699_v65_ = rhs68_
                    d_700_v66_ = rhs69_
                    lhs58_.f10 = rhs70_
                    lhs59_[lhs60_] = rhs71_
                    lhs61_.f24 = rhs72_
                    (globalState).f10 = (d_595_v1_) > (6)
                    d_709_v73_: _dafny.Set
                    d_709_v73_ = _dafny.Set({self.f27})
                    d_710_v74_: _dafny.Map
                    d_710_v74_ = _dafny.Map({self.f27: d_709_v73_})
                    d_710_v74_ = (d_710_v74_).set(_dafny.CodePoint('f'), d_709_v73_)
                    pass
            pass
        d_711_v75_: D7
        d_711_v75_ = D7_DC16(not(self.f24), d_595_v1_, (self).f26, not((self).f32), (self).f25)
        d_712_i15_: int
        d_712_i15_ = 0
        with _dafny.label("7"):
            while (-931) > ((d_711_v75_).cf31):
                with _dafny.c_label("7"):
                    if (d_712_i15_) >= (100):
                        raise _dafny.Break("7")
                    d_712_i15_ = (d_712_i15_) + (1)
                    (globalState).f16 = d_595_v1_
                    d_713_v76_: _dafny.Array
                    nw98_ = _dafny.Array(None, 5)
                    nw98_[int(0)] = (d_595_v1_) - (d_595_v1_)
                    nw98_[int(1)] = default__.safeModuloInt((0) - (d_595_v1_), d_595_v1_)
                    nw98_[int(2)] = 31
                    nw98_[int(3)] = d_595_v1_
                    nw98_[int(4)] = d_595_v1_
                    d_713_v76_ = nw98_
                    index90_ = default__.safeIndex(731, (d_713_v76_).length(0))
                    (d_713_v76_)[index90_] = d_595_v1_
                    d_714_v77_: _dafny.Array
                    nw99_ = _dafny.Array(_dafny.Array(None, 0), 16)
                    d_714_v77_ = nw99_
                    index91_ = default__.safeIndex(896, (d_714_v77_).length(0))
                    (d_714_v77_)[index91_] = d_713_v76_
                    d_715_v78_: _dafny.MultiSet
                    d_715_v78_ = _dafny.MultiSet([self.f24])
                    d_716_v79_: _dafny.Seq
                    d_716_v79_ = _dafny.SeqWithoutIsStrInference([self.f24, not(not((d_715_v78_).isdisjoint(d_715_v78_))), not((self).f25)])
                    d_717_v80_: _dafny.Seq
                    d_717_v80_ = _dafny.SeqWithoutIsStrInference([d_595_v1_, d_595_v1_, 31])
                    index92_ = default__.safeIndex(731, (d_713_v76_).length(0))
                    index93_ = default__.safeIndex(896, (d_714_v77_).length(0))
                    rhs73_ = (d_717_v80_)[default__.safeIndex((d_717_v80_)[default__.safeIndex(d_595_v1_, len(d_717_v80_))], len(d_717_v80_))]
                    rhs74_ = (d_595_v1_) + (len((self).f33))
                    rhs75_ = d_713_v76_
                    rhs76_ = d_716_v79_
                    lhs62_ = d_713_v76_
                    lhs63_ = default__.safeIndex(731, (d_713_v76_).length(0))
                    lhs64_ = globalState
                    lhs65_ = d_714_v77_
                    lhs66_ = default__.safeIndex(896, (d_714_v77_).length(0))
                    lhs62_[lhs63_] = rhs73_
                    lhs64_.f11 = rhs74_
                    lhs65_[lhs66_] = rhs75_
                    d_716_v79_ = rhs76_
                    (globalState).f22 = (d_594_v0_).fm25((self).f33, (self).f32, globalState)
                    index94_ = default__.safeIndex(731, (d_713_v76_).length(0))
                    (d_713_v76_)[index94_] = 164
                    pass
            pass
        r0 = (d_595_v1_) == (default__.safeModuloInt(len((self).f33), 42))
        d_718_v81_: _dafny.Array
        def lambda34_(d_719_i16_):
            return (self).f26

        init17_ = lambda34_
        nw100_ = _dafny.Array(None, 24)
        for i0_17_ in range(nw100_.length(0)):
            nw100_[i0_17_] = init17_(i0_17_)
        d_718_v81_ = nw100_
        r1 = d_718_v81_
        return r0, r1

    def m1(self, p0, globalState):
        r0: str = _dafny.CodePoint('D')
        r1: _dafny.Seq = _dafny.Seq({})
        d_720_v0_: int
        d_720_v0_ = 222
        d_721_v1_: _dafny.Seq
        d_721_v1_ = _dafny.SeqWithoutIsStrInference([d_720_v0_, d_720_v0_, d_720_v0_, d_720_v0_])
        d_722_v2_: _dafny.Seq
        d_722_v2_ = _dafny.SeqWithoutIsStrInference([d_721_v1_, _dafny.SeqWithoutIsStrInference([d_720_v0_, len(_dafny.SeqWithoutIsStrInference([(self).f25]))])])
        d_723_v3_: _dafny.Set
        d_723_v3_ = _dafny.Set({(self).f26})
        d_724_i0_: int
        d_724_i0_ = 0
        with _dafny.label("8"):
            while ((d_722_v2_)[default__.safeIndex(len(d_723_v3_), len(d_722_v2_))]) <= (d_721_v1_):
                with _dafny.c_label("8"):
                    if (d_724_i0_) >= (100):
                        raise _dafny.Break("8")
                    d_724_i0_ = (d_724_i0_) + (1)
                    d_725_v4_: _dafny.Array
                    def lambda35_(d_726_i1_):
                        return True

                    init18_ = lambda35_
                    nw101_ = _dafny.Array(None, 8)
                    for i0_18_ in range(nw101_.length(0)):
                        nw101_[i0_18_] = init18_(i0_18_)
                    d_725_v4_ = nw101_
                    index95_ = default__.safeIndex(839, (d_725_v4_).length(0))
                    (d_725_v4_)[index95_] = self.f24
                    index96_ = default__.safeIndex(839, (d_725_v4_).length(0))
                    (d_725_v4_)[index96_] = (d_720_v0_) == (default__.safeModuloInt(d_720_v0_, d_720_v0_))
                    d_727_v5_: _dafny.Seq
                    d_727_v5_ = _dafny.SeqWithoutIsStrInference([(self).f26, self.f24])
                    if ((0) - (d_720_v0_)) > (len(d_727_v5_)):
                        d_728_v6_: C0
                        nw102_ = C0()
                        nw102_.ctor__((self).f33)
                        d_728_v6_ = nw102_
                        d_729_v7_: _dafny.Map
                        d_729_v7_ = _dafny.Map({(self).f25: d_720_v0_})
                        d_730_v8_: D2
                        d_730_v8_ = D2_DC5(self.f27, self.f24, len((d_728_v6_).f31), d_720_v0_)
                        d_731_v9_: D2
                        d_731_v9_ = D2_DC6(D2_DC6(d_730_v8_))
                        d_732_v10_: D5
                        d_732_v10_ = D5_DC11(d_731_v9_, (self).f26)
                        d_733_v11_: _dafny.MultiSet
                        d_733_v11_ = _dafny.MultiSet([(d_725_v4_)[default__.safeIndex(839, (d_725_v4_).length(0))], (self).f25, (self).fm7(d_729_v7_, globalState), ((self).f26 if (d_725_v4_)[default__.safeIndex(839, (d_725_v4_).length(0))] else (d_732_v10_).cf22)])
                        rhs77_ = (d_733_v11_).cardinality
                        rhs78_ = d_720_v0_
                        lhs67_ = globalState
                        lhs68_ = globalState
                        lhs67_.f16 = rhs77_
                        lhs68_.f22 = rhs78_
                        index97_ = default__.safeIndex(839, (d_725_v4_).length(0))
                        rhs79_ = d_720_v0_
                        rhs80_ = (0) - (d_720_v0_)
                        rhs81_ = (d_720_v0_) >= ((0) - (d_720_v0_))
                        rhs82_ = d_720_v0_
                        rhs83_ = not(not (self.f24) or ((self).f32))
                        lhs69_ = globalState
                        lhs70_ = self
                        lhs71_ = globalState
                        lhs72_ = d_725_v4_
                        lhs73_ = default__.safeIndex(839, (d_725_v4_).length(0))
                        lhs69_.f16 = rhs79_
                        d_720_v0_ = rhs80_
                        lhs70_.f24 = rhs81_
                        lhs71_.f16 = rhs82_
                        lhs72_[lhs73_] = rhs83_
                        d_734_v12_: D6
                        d_734_v12_ = D6_DC12(d_721_v1_)
                        d_735_v13_: _dafny.Map
                        d_735_v13_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kx"))) == ((self).f33): d_734_v12_})
                        d_735_v13_ = (d_735_v13_).set(not ((self).f25) or ((self).f26), d_734_v12_)
                        d_722_v2_ = default__.fm48(globalState)
                    elif True:
                        d_736_v14_: _dafny.Array
                        nw103_ = _dafny.Array(_dafny.CodePoint('D'), 21)
                        d_736_v14_ = nw103_
                        index98_ = default__.safeIndex(886, (d_736_v14_).length(0))
                        (d_736_v14_)[index98_] = self.f27
                        index99_ = default__.safeIndex(886, (d_736_v14_).length(0))
                        (d_736_v14_)[index99_] = self.f27
                        (self).f24 = ((d_725_v4_)[default__.safeIndex(839, (d_725_v4_).length(0))] if not((self).f26) else (d_721_v1_) == (d_721_v1_))
                        d_737_v15_: _dafny.Seq
                        d_737_v15_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "esq"))])
                        d_738_v16_: _dafny.Seq
                        d_738_v16_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gq"))]), _dafny.SeqWithoutIsStrInference([(self).f33])])
                        d_739_v17_: _dafny.Map
                        d_739_v17_ = _dafny.Map({self.f24: 407})
                        d_740_v18_: _dafny.Array
                        nw104_ = _dafny.Array(None, 16)
                        nw104_[int(0)] = _dafny.SeqWithoutIsStrInference([(self).f33 for d_741_i2_ in range(default__.abs(15))])
                        nw104_[int(1)] = d_737_v15_
                        nw104_[int(2)] = (_dafny.SeqWithoutIsStrInference([((self).f33).set(default__.safeIndex(993, len((self).f33)), self.f27) for d_742_i3_ in range(default__.abs(319))])) + (_dafny.SeqWithoutIsStrInference([(self).f33 for d_743_i4_ in range(default__.abs(989))]))
                        nw104_[int(3)] = (d_738_v16_)[default__.safeIndex(d_720_v0_, len(d_738_v16_))]
                        nw104_[int(4)] = (d_737_v15_) + (d_737_v15_)
                        nw104_[int(5)] = d_737_v15_
                        nw104_[int(6)] = (d_737_v15_) + (d_737_v15_)
                        nw104_[int(7)] = _dafny.SeqWithoutIsStrInference([(self).f33])
                        nw104_[int(8)] = d_737_v15_
                        nw104_[int(9)] = (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(d_736_v14_)[default__.safeIndex(886, (d_736_v14_).length(0))] for d_744_i5_ in range(default__.abs(306))])])).set(default__.safeIndex(d_720_v0_, len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(d_736_v14_)[default__.safeIndex(886, (d_736_v14_).length(0))] for d_745_i5_ in range(default__.abs(306))])]))), (self).f33)
                        nw104_[int(10)] = default__.fm3((_dafny.SeqWithoutIsStrInference([True, not((self).fm7(d_739_v17_, globalState))])).set(default__.safeIndex(d_720_v0_, len(_dafny.SeqWithoutIsStrInference([True, not((self).fm7(d_739_v17_, globalState))]))), (self).f25), 369, _dafny.CodePoint('m'), (self).f26, globalState)
                        nw104_[int(11)] = d_737_v15_
                        nw104_[int(12)] = (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ch"))])) + (d_737_v15_)
                        nw104_[int(13)] = d_737_v15_
                        nw104_[int(14)] = d_737_v15_
                        nw104_[int(15)] = d_737_v15_
                        d_740_v18_ = nw104_
                        index100_ = default__.safeIndex(503, (d_740_v18_).length(0))
                        (d_740_v18_)[index100_] = d_737_v15_
                        index101_ = default__.safeIndex(503, (d_740_v18_).length(0))
                        (d_740_v18_)[index101_] = d_737_v15_
                        (globalState).f22 = (len((self).f33)) * (d_720_v0_)
                        (globalState).f16 = 495
                    (globalState).f10 = self.f24
                    d_746_v19_: _dafny.Set
                    d_746_v19_ = _dafny.Set({self.f27, self.f27})
                    r1 = ((default__.fm49(True, not((self).f32), len(d_746_v19_), (d_725_v4_)[default__.safeIndex(839, (d_725_v4_).length(0))], globalState) if (self).f32 else d_727_v5_)) + (d_727_v5_)
                    pass
            pass
        d_747_v20_: _dafny.Array
        nw105_ = _dafny.Array(D6.default()(), 24)
        d_747_v20_ = nw105_
        d_748_v21_: _dafny.Array
        nw106_ = _dafny.Array(False, 8)
        d_748_v21_ = nw106_
        d_749_v22_: D6
        d_749_v22_ = D6_DC13(d_748_v21_, len(_dafny.SeqWithoutIsStrInference([(self).f33])), (self).f32)
        index102_ = default__.safeIndex(347, (d_747_v20_).length(0))
        (d_747_v20_)[index102_] = d_749_v22_
        d_750_v23_: _dafny.Array
        nw107_ = _dafny.Array(_dafny.Array(None, 0), 29)
        d_750_v23_ = nw107_
        d_751_v24_: _dafny.Array
        def lambda36_(d_752_i6_):
            return _dafny.SeqWithoutIsStrInference([(self).f25, True])

        init19_ = lambda36_
        nw108_ = _dafny.Array(None, 24)
        for i0_19_ in range(nw108_.length(0)):
            nw108_[i0_19_] = init19_(i0_19_)
        d_751_v24_ = nw108_
        index103_ = default__.safeIndex(584, (d_750_v23_).length(0))
        (d_750_v23_)[index103_] = d_751_v24_
        d_753_v25_: _dafny.Array
        nw109_ = _dafny.Array(_dafny.Array(None, 0), 23)
        d_753_v25_ = nw109_
        d_754_v26_: _dafny.Array
        nw110_ = _dafny.Array(_dafny.Array(None, 0), 10)
        d_754_v26_ = nw110_
        index104_ = default__.safeIndex(988, (d_753_v25_).length(0))
        (d_753_v25_)[index104_] = d_754_v26_
        d_755_v27_: _dafny.Seq
        d_755_v27_ = _dafny.SeqWithoutIsStrInference([d_751_v24_, d_751_v24_, d_751_v24_])
        index105_ = default__.safeIndex(347, (d_747_v20_).length(0))
        index106_ = default__.safeIndex(584, (d_750_v23_).length(0))
        index107_ = default__.safeIndex(988, (d_753_v25_).length(0))
        rhs84_ = d_749_v22_
        rhs85_ = (d_755_v27_)[default__.safeIndex((0) - (d_720_v0_), len(d_755_v27_))]
        rhs86_ = d_754_v26_
        rhs87_ = d_720_v0_
        lhs74_ = d_747_v20_
        lhs75_ = default__.safeIndex(347, (d_747_v20_).length(0))
        lhs76_ = d_750_v23_
        lhs77_ = default__.safeIndex(584, (d_750_v23_).length(0))
        lhs78_ = d_753_v25_
        lhs79_ = default__.safeIndex(988, (d_753_v25_).length(0))
        lhs80_ = globalState
        lhs74_[lhs75_] = rhs84_
        lhs76_[lhs77_] = rhs85_
        lhs78_[lhs79_] = rhs86_
        lhs80_.f11 = rhs87_
        d_756_i7_: int
        d_756_i7_ = 0
        with _dafny.label("9"):
            while (self).f26:
                with _dafny.c_label("9"):
                    if (d_756_i7_) >= (100):
                        raise _dafny.Break("9")
                    d_756_i7_ = (d_756_i7_) + (1)
                    if (self.f27) in ((((self).f33).set(default__.safeIndex(d_720_v0_, len((self).f33)), self.f27)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wbjcugj")))):
                        d_757_v28_: D0
                        d_757_v28_ = D0_DC0((self).f25, 262, d_720_v0_)
                        d_758_v29_: _dafny.Seq
                        d_758_v29_ = _dafny.SeqWithoutIsStrInference([D0_DC0(self.f24, d_720_v0_, (d_721_v1_)[default__.safeIndex(d_720_v0_, len(d_721_v1_))]), d_757_v28_])
                        d_759_v30_: _dafny.Map
                        d_759_v30_ = _dafny.Map({(self).f32: d_758_v29_})
                        d_760_v31_: _dafny.Seq
                        d_760_v31_ = _dafny.SeqWithoutIsStrInference([False, (self).f26])
                        d_761_v32_: D3
                        d_761_v32_ = D3_DC7(d_760_v31_)
                        d_762_v33_: D3
                        d_762_v33_ = D3_DC8(not(not((self).f25)), (_dafny.MultiSet((d_761_v32_).cf15)).cardinality, len(_dafny.SeqWithoutIsStrInference([(self).fm5(d_759_v30_, self.f24, (self).f25, globalState), 774])))
                        d_763_v34_: _dafny.Map
                        d_763_v34_ = _dafny.Map({(self).fm5(d_759_v30_, (d_762_v33_).cf16, False, globalState): self.f24})
                        (globalState).f11 = len(((d_721_v1_).set(default__.safeIndex(len((self).f33), len(d_721_v1_)), d_720_v0_)).set(default__.safeIndex(len(d_763_v34_), len((d_721_v1_).set(default__.safeIndex(len((self).f33), len(d_721_v1_)), d_720_v0_))), d_720_v0_))
                        d_764_v35_: C0
                        nw111_ = C0()
                        nw111_.ctor__((self).f33)
                        d_764_v35_ = nw111_
                        d_765_v36_: _dafny.MultiSet
                        d_765_v36_ = _dafny.MultiSet([(self).f25])
                        (globalState).f16 = default__.safeDivisionInt(d_720_v0_, ((d_765_v36_)[(self).f32] if ((self).f32) in (d_765_v36_) else len(d_721_v1_)))
                        (globalState).f22 = d_720_v0_
                        d_766_v37_: _dafny.MultiSet
                        d_766_v37_ = _dafny.MultiSet([d_720_v0_])
                        d_767_v38_: _dafny.Map
                        d_767_v38_ = _dafny.Map({d_720_v0_: d_766_v37_})
                        d_767_v38_ = d_767_v38_
                    elif True:
                        d_768_v39_: _dafny.Map
                        d_768_v39_ = _dafny.Map({default__.safeDivisionInt((d_721_v1_)[default__.safeIndex(d_720_v0_, len(d_721_v1_))], len(default__.fm50(D2_DC6(D2_DC4(self.f24)), d_720_v0_, False, globalState))): (self).f32})
                        d_768_v39_ = (d_768_v39_).set(d_720_v0_, self.f24)
                        d_769_v40_: _dafny.Map
                        d_769_v40_ = _dafny.Map({(self).f26: _dafny.CodePoint('d')})
                        (globalState).f10 = not((False) == ((d_769_v40_) != (d_769_v40_)))
                        index108_ = default__.safeIndex(777, (d_748_v21_).length(0))
                        (d_748_v21_)[index108_] = (self).f25
                        index109_ = default__.safeIndex(777, (d_748_v21_).length(0))
                        (d_748_v21_)[index109_] = (self).f32
                        d_770_v41_: _dafny.MultiSet
                        d_770_v41_ = _dafny.MultiSet([(self).f32])
                        (globalState).f11 = (self).fm6((self).fm4(d_770_v41_, (self).f33, globalState), globalState)
                        d_771_v42_: C1
                        nw112_ = C1()
                        nw112_.ctor__((d_748_v21_)[default__.safeIndex(777, (d_748_v21_).length(0))])
                        d_771_v42_ = nw112_
                    index110_ = default__.safeIndex(897, (d_748_v21_).length(0))
                    (d_748_v21_)[index110_] = (self).f32
                    d_772_v43_: _dafny.Map
                    d_772_v43_ = _dafny.Map({(self).f32: d_720_v0_})
                    index111_ = default__.safeIndex(897, (d_748_v21_).length(0))
                    (d_748_v21_)[index111_] = (self).fm7((_dafny.Map({(self).f25: d_720_v0_})) | (d_772_v43_), globalState)
                    if (d_748_v21_)[default__.safeIndex(897, (d_748_v21_).length(0))]:
                        d_773_v44_: _dafny.Seq
                        d_773_v44_ = _dafny.SeqWithoutIsStrInference([(self).f26, (self).f32, self.f24, (self).f25])
                        (globalState).f11 = len(((default__.fm49((self).f32, (self).f25, d_720_v0_, False, globalState)) + (d_773_v44_)) + (d_773_v44_))
                        d_774_v45_: C0
                        nw113_ = C0()
                        nw113_.ctor__((self).f33)
                        d_774_v45_ = nw113_
                        d_775_v46_: _dafny.Array
                        nw114_ = _dafny.Array(False, 8)
                        d_775_v46_ = nw114_
                        d_776_v47_: _dafny.Map
                        d_776_v47_ = _dafny.Map({self.f24: d_775_v46_})
                        (globalState).f11 = len((d_776_v47_) | (d_776_v47_))
                        d_777_v48_: _dafny.Seq
                        d_777_v48_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "leerl"))
                        d_777_v48_ = ((d_777_v48_ if (d_748_v21_)[default__.safeIndex(897, (d_748_v21_).length(0))] else (self).f33)) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uvggjar")) if (self).f32 else (self).f33))
                        (self).f24 = (default__.fm1((self).f32, (self).f25, d_720_v0_, globalState)) == ((d_748_v21_)[default__.safeIndex(897, (d_748_v21_).length(0))])
                    elif True:
                        index112_ = default__.safeIndex(897, (d_748_v21_).length(0))
                        (d_748_v21_)[index112_] = (self).f26
                        d_778_v49_: _dafny.Map
                        d_778_v49_ = _dafny.Map({self.f24: (d_748_v21_)[default__.safeIndex(897, (d_748_v21_).length(0))]})
                        d_779_v50_: _dafny.Map
                        d_779_v50_ = _dafny.Map({(self).f32: d_778_v49_})
                        (globalState).f10 = (d_720_v0_) <= (len(_dafny.Set({d_778_v49_, ((d_779_v50_)[not(True)] if (not(True)) in (d_779_v50_) else d_778_v49_)})))
                        d_780_v51_: D0
                        d_780_v51_ = D0_DC0((self).f32, d_720_v0_, d_720_v0_)
                        d_781_v52_: _dafny.Seq
                        d_781_v52_ = _dafny.SeqWithoutIsStrInference([d_780_v51_])
                        d_782_v53_: _dafny.Map
                        d_782_v53_ = _dafny.Map({False: d_781_v52_})
                        d_783_v54_: _dafny.Array
                        nw115_ = _dafny.Array(int(0), 25)
                        d_783_v54_ = nw115_
                        d_784_v55_: D1
                        d_784_v55_ = D1_DC2(d_783_v54_, self.f27, d_720_v0_, (self).f26)
                        d_785_v56_: _dafny.MultiSet
                        d_785_v56_ = _dafny.MultiSet([d_784_v55_])
                        d_786_v58_: _dafny.MultiSet
                        d_786_v58_ = _dafny.MultiSet([self.f27])
                        d_787_v59_: _dafny.Array
                        nw116_ = _dafny.Array(None, 17)
                        nw116_[int(0)] = d_720_v0_
                        nw116_[int(1)] = (self).fm5(d_782_v53_, (self).f32, (self).f26, globalState)
                        nw116_[int(2)] = ((d_785_v56_)[d_784_v55_] if (d_784_v55_) in (d_785_v56_) else d_720_v0_)
                        def iife63_():
                            coll31_ = _dafny.Map()
                            compr_31_: int
                            for compr_31_ in _dafny.IntegerRange(895, 996):
                                d_788_v57_: int = compr_31_
                                if ((895) <= (d_788_v57_)) and ((d_788_v57_) < (996)):
                                    coll31_[(d_788_v57_) + (d_720_v0_)] = (self).f26
                            return _dafny.Map(coll31_)
                        nw116_[int(3)] = len(iife63_()
                        )
                        nw116_[int(4)] = d_720_v0_
                        nw116_[int(5)] = d_720_v0_
                        nw116_[int(6)] = (d_784_v55_).cf6
                        nw116_[int(7)] = d_720_v0_
                        nw116_[int(8)] = d_720_v0_
                        nw116_[int(9)] = len(d_772_v43_)
                        nw116_[int(10)] = ((d_786_v58_)[self.f27] if (self.f27) in (d_786_v58_) else d_720_v0_)
                        nw116_[int(11)] = d_720_v0_
                        nw116_[int(12)] = d_720_v0_
                        nw116_[int(13)] = d_720_v0_
                        nw116_[int(14)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y")))
                        nw116_[int(15)] = d_720_v0_
                        nw116_[int(16)] = d_720_v0_
                        d_787_v59_ = nw116_
                        d_789_v60_: _dafny.Map
                        d_789_v60_ = _dafny.Map({d_720_v0_: d_783_v54_})
                        d_790_v61_: _dafny.Array
                        nw117_ = _dafny.Array(None, 23)
                        nw117_[int(0)] = d_720_v0_
                        nw117_[int(1)] = d_720_v0_
                        nw117_[int(2)] = 706
                        nw117_[int(3)] = d_720_v0_
                        nw117_[int(4)] = (0) - ((0) - (d_720_v0_))
                        nw117_[int(5)] = d_720_v0_
                        nw117_[int(6)] = d_720_v0_
                        nw117_[int(7)] = d_720_v0_
                        nw117_[int(8)] = d_720_v0_
                        nw117_[int(9)] = d_720_v0_
                        nw117_[int(10)] = -287
                        nw117_[int(11)] = d_720_v0_
                        nw117_[int(12)] = d_720_v0_
                        nw117_[int(13)] = d_720_v0_
                        nw117_[int(14)] = d_720_v0_
                        nw117_[int(15)] = d_720_v0_
                        nw117_[int(16)] = d_720_v0_
                        nw117_[int(17)] = d_720_v0_
                        nw117_[int(18)] = d_720_v0_
                        nw117_[int(19)] = len(d_789_v60_)
                        nw117_[int(20)] = d_720_v0_
                        nw117_[int(21)] = d_720_v0_
                        nw117_[int(22)] = 858
                        d_790_v61_ = nw117_
                        d_791_v62_: _dafny.Array
                        nw118_ = _dafny.Array(None, 1)
                        nw118_[int(0)] = d_787_v59_
                        d_791_v62_ = nw118_
                        index113_ = default__.safeIndex(552, (d_791_v62_).length(0))
                        (d_791_v62_)[index113_] = d_783_v54_
                        index114_ = default__.safeIndex(552, (d_791_v62_).length(0))
                        (d_791_v62_)[index114_] = d_790_v61_
                        d_792_v63_: _dafny.Map
                        d_792_v63_ = _dafny.Map({(self).f33: self.f27})
                        d_792_v63_ = (d_792_v63_).set((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_793_i8_ in range(default__.abs(273))])).set(default__.safeIndex(d_720_v0_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_794_i8_ in range(default__.abs(273))]))), self.f27), self.f27)
                        (self).f24 = (self).f25
                    d_795_v64_: _dafny.Seq
                    d_795_v64_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "owfhgc"))
                    d_796_v65_: _dafny.Array
                    nw119_ = _dafny.Array(_dafny.Array(None, 0), 15)
                    d_796_v65_ = nw119_
                    d_797_v66_: _dafny.Seq
                    d_797_v66_ = _dafny.SeqWithoutIsStrInference([False])
                    d_798_v67_: D0
                    d_798_v67_ = D0_DC0((d_748_v21_)[default__.safeIndex(897, (d_748_v21_).length(0))], len(d_795_v64_), len(d_797_v66_))
                    d_799_v68_: _dafny.Seq
                    d_799_v68_ = _dafny.SeqWithoutIsStrInference([d_798_v67_])
                    d_800_v69_: _dafny.Map
                    d_800_v69_ = _dafny.Map({(self).f25: d_799_v68_})
                    d_801_v70_: _dafny.Map
                    d_801_v70_ = _dafny.Map({_dafny.CodePoint('m'): (self).f32})
                    d_802_v71_: _dafny.MultiSet
                    d_802_v71_ = _dafny.MultiSet([self.f24])
                    d_803_v72_: _dafny.MultiSet
                    d_803_v72_ = _dafny.MultiSet([d_802_v71_, d_802_v71_])
                    d_804_v73_: D2
                    d_804_v73_ = D2_DC5(self.f27, (self).f32, d_720_v0_, d_720_v0_)
                    d_805_v74_: _dafny.Map
                    d_805_v74_ = _dafny.Map({(d_804_v73_).cf10: (d_803_v72_).set(d_802_v71_, default__.abs((_dafny.MultiSet([d_720_v0_])).cardinality))})
                    index115_ = default__.safeIndex(897, (d_748_v21_).length(0))
                    rhs88_ = (self).fm5(d_800_v69_, (d_720_v0_) <= ((d_721_v1_)[default__.safeIndex(default__.fm0(d_720_v0_, (self).f26, _dafny.Map({len((_dafny.Map({(self).f26: d_720_v0_})).set(not(((d_801_v70_)[self.f27] if (self.f27) in (d_801_v70_) else (d_748_v21_)[default__.safeIndex(897, (d_748_v21_).length(0))])), d_720_v0_)): (d_748_v21_)[default__.safeIndex(897, (d_748_v21_).length(0))]}), d_720_v0_, globalState), len(d_721_v1_))]), (self).fm39(d_721_v1_, (d_748_v21_)[default__.safeIndex(897, (d_748_v21_).length(0))], globalState), globalState)
                    rhs89_ = d_723_v3_
                    rhs90_ = (((d_805_v74_)[self.f27] if (self.f27) in (d_805_v74_) else d_803_v72_)).issubset((d_803_v72_).set(d_802_v71_, default__.abs(d_720_v0_)))
                    rhs91_ = ((d_795_v64_ if (d_748_v21_)[default__.safeIndex(897, (d_748_v21_).length(0))] else (self).f33)) + ((self).f33)
                    rhs92_ = d_796_v65_
                    lhs81_ = globalState
                    lhs82_ = d_748_v21_
                    lhs83_ = default__.safeIndex(897, (d_748_v21_).length(0))
                    lhs81_.f11 = rhs88_
                    d_723_v3_ = rhs89_
                    lhs82_[lhs83_] = rhs90_
                    d_795_v64_ = rhs91_
                    d_796_v65_ = rhs92_
                    pass
            pass
        d_806_v75_: _dafny.Map
        d_806_v75_ = _dafny.Map({(123) <= (d_720_v0_): self.f24})
        d_807_v76_: D1
        d_807_v76_ = D1_DC1((self).f33)
        d_806_v75_ = (d_806_v75_).set((self).f26, ((d_807_v76_).cf3) == ((self).f33))
        d_808_v77_: D6
        d_808_v77_ = D6_DC12(d_721_v1_)
        d_808_v77_ = d_808_v77_
        if (self).f25:
            d_809_v78_: int
            d_810_v79_: _dafny.Seq
            d_811_v80_: _dafny.Array
            d_812_v81_: _dafny.Seq
            out21_: int
            out22_: _dafny.Seq
            out23_: _dafny.Array
            out24_: _dafny.Seq
            out21_, out22_, out23_, out24_ = (self).m8(True, globalState)
            d_809_v78_ = out21_
            d_810_v79_ = out22_
            d_811_v80_ = out23_
            d_812_v81_ = out24_
            index116_ = default__.safeIndex(436, (d_748_v21_).length(0))
            (d_748_v21_)[index116_] = self.f24
            d_813_v82_: _dafny.Seq
            d_813_v82_ = _dafny.SeqWithoutIsStrInference([(self).f26, (self).f26])
            index117_ = default__.safeIndex(436, (d_748_v21_).length(0))
            (d_748_v21_)[index117_] = not(not(((self).f25 if (self).f25 else ((default__.fm51(self.f27, globalState)).cardinality) != (len(d_813_v82_)))))
            d_814_v83_: C0
            nw120_ = C0()
            nw120_.ctor__((self).f33)
            d_814_v83_ = nw120_
            d_811_v80_ = d_811_v80_
            d_815_v84_: _dafny.Map
            d_815_v84_ = _dafny.Map({(d_748_v21_)[default__.safeIndex(436, (d_748_v21_).length(0))]: d_720_v0_})
            d_816_v85_: D0
            d_816_v85_ = D0_DC0((self).fm7(d_815_v84_, globalState), -600, d_809_v78_)
            source16_ = d_816_v85_
            d_817___mcc_h0_ = source16_.cf0
            d_818___mcc_h1_ = source16_.cf1
            d_819___mcc_h2_ = source16_.cf2
            d_820_cf2_ = d_819___mcc_h2_
            d_821_cf1_ = d_818___mcc_h1_
            d_822_cf0_ = d_817___mcc_h0_
            d_823_v86_: C1
            nw121_ = C1()
            nw121_.ctor__(self.f24)
            d_823_v86_ = nw121_
            d_824_v87_: _dafny.MultiSet
            d_824_v87_ = _dafny.MultiSet([(False if (self).f26 else not(True)), (self).f26, (self).f25])
            d_809_v78_ = ((d_824_v87_)[(self).f25] if ((self).f25) in (d_824_v87_) else ((d_721_v1_)[default__.safeIndex(d_809_v78_, len(d_721_v1_))]) - (d_809_v78_))
            d_825_v88_: _dafny.Seq
            d_825_v88_ = _dafny.SeqWithoutIsStrInference([d_824_v87_])
            d_826_v89_: T0
            nw122_ = C2()
            nw122_.ctor__(d_810_v79_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sx")))
            d_826_v89_ = nw122_
            d_827_v90_: _dafny.Map
            d_827_v90_ = _dafny.Map({d_826_v89_: self.f24})
            d_824_v87_ = (d_825_v88_)[default__.safeIndex((len(d_827_v90_) if d_822_cf0_ else d_720_v0_), len(d_825_v88_))]
            d_828_v91_: _dafny.Seq
            d_828_v91_ = _dafny.SeqWithoutIsStrInference([d_813_v82_, _dafny.SeqWithoutIsStrInference([True])])
            r1 = (d_828_v91_)[default__.safeIndex(d_720_v0_, len(d_828_v91_))]
        elif True:
            d_720_v0_ = (0) - (d_720_v0_)
            (globalState).f10 = self.f24
            d_829_v92_: _dafny.MultiSet
            d_829_v92_ = _dafny.MultiSet([(self).f32, ((self).f26 if (self).f32 else (self).f25), self.f24, (self).f32])
            d_829_v92_ = d_829_v92_
            d_830_v93_: _dafny.Seq
            d_830_v93_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "srnc"))
            d_830_v93_ = (d_830_v93_) + ((self).f33)
            (globalState).f11 = (d_720_v0_) - ((0) - (default__.safeModuloInt(d_720_v0_, len((self).f33))))
        r0 = self.f27
        d_831_v94_: _dafny.Seq
        d_831_v94_ = _dafny.SeqWithoutIsStrInference([not(not((self).f26)), self.f24, (self).f25])
        r1 = ((d_831_v94_) + (d_831_v94_)) + ((d_831_v94_) + (_dafny.SeqWithoutIsStrInference([self.f24])))
        return r0, r1

    def m8(self, p0, globalState):
        r0: int = int(0)
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r2: _dafny.Array = _dafny.Array(None, 0)
        r3: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        d_832_v0_: int
        d_832_v0_ = 455
        d_833_v1_: _dafny.Map
        d_833_v1_ = _dafny.Map({(self).f25: 753})
        hi3_ = (d_832_v0_) - (len(d_833_v1_))
        for d_834_i0_ in range(len((self).f33), hi3_):
            d_835_v2_: _dafny.MultiSet
            d_835_v2_ = _dafny.MultiSet([d_832_v0_])
            d_836_v3_: _dafny.Array
            nw123_ = _dafny.Array(None, 7)
            nw123_[int(0)] = _dafny.MultiSet([d_832_v0_])
            nw123_[int(1)] = d_835_v2_
            nw123_[int(2)] = d_835_v2_
            nw123_[int(3)] = _dafny.MultiSet([(0) - (d_834_i0_)])
            nw123_[int(4)] = (d_835_v2_) - (d_835_v2_)
            nw123_[int(5)] = (d_835_v2_).intersection(default__.fm52((self).f26, globalState))
            nw123_[int(6)] = d_835_v2_
            d_836_v3_ = nw123_
            index118_ = default__.safeIndex(860, (d_836_v3_).length(0))
            (d_836_v3_)[index118_] = d_835_v2_
            d_837_v4_: _dafny.Set
            d_837_v4_ = _dafny.Set({d_832_v0_, d_834_i0_})
            d_838_v6_: _dafny.Map
            d_838_v6_ = _dafny.Map({len(((self).f33).set(default__.safeIndex(len(d_837_v4_), len((self).f33)), self.f27)): d_837_v4_})
            d_839_v7_: _dafny.Set
            d_839_v7_ = _dafny.Set({False, (self).f26, (self).f25, self.f24})
            d_840_v9_: _dafny.Seq
            def iife64_():
                coll32_ = _dafny.Set()
                compr_32_: int
                for compr_32_ in _dafny.IntegerRange(672, 729):
                    d_841_v5_: int = compr_32_
                    if ((672) <= (d_841_v5_)) and ((d_841_v5_) < (729)):
                        coll32_ = coll32_.union(_dafny.Set([default__.safeModuloInt(d_841_v5_, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False, not((self).f32)]))).cardinality)]))
                return _dafny.Set(coll32_)
            def iife65_():
                coll33_ = _dafny.Set()
                compr_33_: int
                for compr_33_ in _dafny.IntegerRange(475, 211):
                    d_842_v8_: int = compr_33_
                    if ((475) <= (d_842_v8_)) and ((d_842_v8_) < (211)):
                        coll33_ = coll33_.union(_dafny.Set([(d_842_v8_) - (d_832_v0_)]))
                return _dafny.Set(coll33_)
            d_840_v9_ = _dafny.SeqWithoutIsStrInference([d_837_v4_, iife64_()
            , ((d_838_v6_)[len(d_839_v7_)] if (len(d_839_v7_)) in (d_838_v6_) else iife65_()
            ), d_837_v4_, d_837_v4_])
            d_843_v10_: _dafny.Array
            def lambda37_(d_844_i1_):
                return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ekgjc"))

            init20_ = lambda37_
            nw124_ = _dafny.Array(None, 8)
            for i0_20_ in range(nw124_.length(0)):
                nw124_[i0_20_] = init20_(i0_20_)
            d_843_v10_ = nw124_
            index119_ = default__.safeIndex(577, (d_843_v10_).length(0))
            (d_843_v10_)[index119_] = (self).f33
            d_845_v11_: _dafny.Array
            nw125_ = _dafny.Array(False, 26)
            d_845_v11_ = nw125_
            index120_ = default__.safeIndex(860, (d_836_v3_).length(0))
            index121_ = default__.safeIndex(577, (d_843_v10_).length(0))
            rhs93_ = (d_835_v2_) - ((d_835_v2_).set(len(d_833_v1_), default__.abs(d_832_v0_)))
            rhs94_ = d_840_v9_
            rhs95_ = d_845_v11_
            rhs96_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ymb"))
            lhs84_ = d_836_v3_
            lhs85_ = default__.safeIndex(860, (d_836_v3_).length(0))
            lhs86_ = globalState
            lhs87_ = d_843_v10_
            lhs88_ = default__.safeIndex(577, (d_843_v10_).length(0))
            lhs84_[lhs85_] = rhs93_
            d_840_v9_ = rhs94_
            lhs86_.f2 = rhs95_
            lhs87_[lhs88_] = rhs96_
            (globalState).f10 = True
            d_846_v12_: _dafny.Seq
            d_846_v12_ = _dafny.SeqWithoutIsStrInference([d_832_v0_])
            (globalState).f16 = default__.safeModuloInt((((d_836_v3_)[default__.safeIndex(860, (d_836_v3_).length(0))]) | (_dafny.MultiSet(d_846_v12_))).cardinality, d_834_i0_)
            d_847_v13_: _dafny.Set
            d_847_v13_ = _dafny.Set({(d_843_v10_)[default__.safeIndex(577, (d_843_v10_).length(0))]})
            d_848_v15_: _dafny.Map
            def iife66_():
                coll34_ = _dafny.Set()
                compr_34_: int
                for compr_34_ in (d_846_v12_).Elements:
                    d_849_v14_: int = compr_34_
                    if (d_849_v14_) in (d_846_v12_):
                        coll34_ = coll34_.union(_dafny.Set([(d_849_v14_) + (-440)]))
                return _dafny.Set(coll34_)
            d_848_v15_ = _dafny.Map({iife66_()
            : -625})
            d_847_v13_ = (d_847_v13_ if (d_837_v4_) not in (d_848_v15_) else d_847_v13_)
        d_850_v16_: _dafny.MultiSet
        d_850_v16_ = _dafny.MultiSet([d_832_v0_, (((d_833_v1_)[p0] if (p0) in (d_833_v1_) else d_832_v0_) if (self).f32 else d_832_v0_)])
        d_850_v16_ = d_850_v16_
        (globalState).f10 = (self).f32
        d_851_v17_: _dafny.Array
        nw126_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 18)
        d_851_v17_ = nw126_
        index122_ = default__.safeIndex(826, (d_851_v17_).length(0))
        (d_851_v17_)[index122_] = (self).f33
        index123_ = default__.safeIndex(826, (d_851_v17_).length(0))
        (d_851_v17_)[index123_] = (((self).f33) + (_dafny.SeqWithoutIsStrInference([self.f27 for d_852_i2_ in range(default__.abs(-754))]))).set(default__.safeIndex(d_832_v0_, len(((self).f33) + (_dafny.SeqWithoutIsStrInference([self.f27 for d_853_i2_ in range(default__.abs(-754))])))), self.f27)
        (globalState).f16 = (0) - ((d_832_v0_) + ((len(d_833_v1_)) + (d_832_v0_)))
        d_854_v18_: _dafny.Seq
        d_854_v18_ = _dafny.SeqWithoutIsStrInference([(self).f26])
        d_855_v19_: D1
        d_855_v19_ = D1_DC1((d_851_v17_)[default__.safeIndex(826, (d_851_v17_).length(0))])
        d_856_v20_: _dafny.Set
        d_856_v20_ = _dafny.Set({(self).f25})
        d_857_v22_: D0
        d_857_v22_ = D0_DC0((self).f26, d_832_v0_, len((d_851_v17_)[default__.safeIndex(826, (d_851_v17_).length(0))]))
        d_858_v23_: _dafny.Seq
        d_858_v23_ = _dafny.SeqWithoutIsStrInference([d_832_v0_, d_832_v0_, len((self).f33)])
        d_859_v24_: _dafny.Map
        d_859_v24_ = _dafny.Map({d_832_v0_: (self).f25})
        d_860_v25_: _dafny.Array
        nw127_ = _dafny.Array(None, 24)
        nw127_[int(0)] = d_832_v0_
        nw127_[int(1)] = default__.safeModuloInt(d_832_v0_, d_832_v0_)
        nw127_[int(2)] = (d_832_v0_) - (len(d_854_v18_))
        nw127_[int(3)] = d_832_v0_
        nw127_[int(4)] = (-463) - (d_832_v0_)
        nw127_[int(5)] = d_832_v0_
        nw127_[int(6)] = d_832_v0_
        nw127_[int(7)] = d_832_v0_
        nw127_[int(8)] = ((0) - (len((d_855_v19_).cf3))) * (d_832_v0_)
        nw127_[int(9)] = len(d_856_v20_)
        nw127_[int(10)] = default__.safeDivisionInt(d_832_v0_, d_832_v0_)
        nw127_[int(11)] = (d_832_v0_) + (-633)
        def iife67_():
            coll35_ = _dafny.Set()
            compr_35_: int
            for compr_35_ in _dafny.IntegerRange(778, 505):
                d_861_v21_: int = compr_35_
                if ((778) <= (d_861_v21_)) and ((d_861_v21_) < (505)):
                    coll35_ = coll35_.union(_dafny.Set([(d_861_v21_) - (968)]))
            return _dafny.Set(coll35_)
        nw127_[int(12)] = len((default__.fm53(self.f27, self.f27, globalState)) | (iife67_()
        ))
        nw127_[int(13)] = (0) - (d_832_v0_)
        nw127_[int(14)] = d_832_v0_
        nw127_[int(15)] = (self).fm6(d_857_v22_, globalState)
        nw127_[int(16)] = d_832_v0_
        nw127_[int(17)] = d_832_v0_
        nw127_[int(18)] = len(d_858_v23_)
        nw127_[int(19)] = len((d_859_v24_) | (d_859_v24_))
        nw127_[int(20)] = d_832_v0_
        nw127_[int(21)] = 377
        nw127_[int(22)] = (0) - ((len(_dafny.SeqWithoutIsStrInference([self.f27 for d_862_i3_ in range(default__.abs(-419))]))) - (d_832_v0_))
        nw127_[int(23)] = d_832_v0_
        d_860_v25_ = nw127_
        index124_ = default__.safeIndex(9, (d_860_v25_).length(0))
        (d_860_v25_)[index124_] = (d_832_v0_) * (len(_dafny.SeqWithoutIsStrInference([self.f27 for d_863_i4_ in range(default__.abs(-89))])))
        index125_ = default__.safeIndex(9, (d_860_v25_).length(0))
        (d_860_v25_)[index125_] = d_832_v0_
        r0 = (d_832_v0_ if (self.f27) in (_dafny.SeqWithoutIsStrInference([self.f27 for d_864_i5_ in range(default__.abs(649))])) else d_832_v0_)
        r1 = ((d_851_v17_)[default__.safeIndex(826, (d_851_v17_).length(0))]) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ohus")))
        r2 = d_860_v25_
        r3 = _dafny.SeqWithoutIsStrInference([self.f27 for d_865_i6_ in range(default__.abs(-574))])
        return r0, r1, r2, r3

    def m9(self, globalState):
        d_866_v0_: int
        d_866_v0_ = 615
        d_867_v1_: _dafny.Set
        d_867_v1_ = _dafny.Set({d_866_v0_})
        d_868_v2_: _dafny.Seq
        d_868_v2_ = _dafny.SeqWithoutIsStrInference([len((self).f33), len(d_867_v1_), (0) - ((0) - (d_866_v0_)), d_866_v0_])
        d_869_v3_: _dafny.Set
        d_869_v3_ = _dafny.Set({d_868_v2_})
        if ((_dafny.Set({d_868_v2_})).intersection(d_869_v3_)).ispropersubset(d_869_v3_):
            d_870_v4_: _dafny.MultiSet
            d_870_v4_ = _dafny.MultiSet([not((self).f25)])
            (globalState).f10 = (d_870_v4_).ispropersubset(d_870_v4_)
            d_871_v5_: C2
            nw128_ = C2()
            nw128_.ctor__(((self).f33 if self.f24 else (self).f33), (self).f33)
            d_871_v5_ = nw128_
            rhs97_ = (self).f25
            rhs98_ = (self).f26
            lhs89_ = globalState
            lhs90_ = globalState
            lhs89_.f10 = rhs97_
            lhs90_.f10 = rhs98_
            (d_871_v5_).f29 = (d_871_v5_).f30
            (self).f24 = ((self).f25) or ((self).f32)
        elif True:
            d_872_v6_: _dafny.Array
            nw129_ = _dafny.Array(_dafny.Array(None, 0), 14)
            d_872_v6_ = nw129_
            d_873_v7_: _dafny.Array
            d_873_v7_ = d_872_v6_
            source17_ = (d_873_v7_ if (self).f25 else d_873_v7_)
            d_874___mcc_h0_ = source17_
            d_875_cf19_ = d_874___mcc_h0_
            d_876_v8_: _dafny.Array
            nw130_ = _dafny.Array(_dafny.Seq({}), 21)
            d_876_v8_ = nw130_
            d_877_v9_: _dafny.Seq
            d_877_v9_ = _dafny.SeqWithoutIsStrInference([(self).f26, (self).f25])
            index126_ = default__.safeIndex(166, (d_876_v8_).length(0))
            (d_876_v8_)[index126_] = d_877_v9_
            d_878_v10_: D3
            d_878_v10_ = D3_DC7(d_877_v9_)
            index127_ = default__.safeIndex(166, (d_876_v8_).length(0))
            (d_876_v8_)[index127_] = ((d_878_v10_).cf15) + ((default__.fm49((self).f26, (self).f32, d_866_v0_, (self).f25, globalState)) + (default__.fm49(self.f24, self.f24, len(d_877_v9_), (self).f32, globalState)))
            (self).f27 = default__.fm47(globalState)
            d_879_v11_: _dafny.Map
            d_879_v11_ = _dafny.Map({d_866_v0_: (self).f32})
            d_880_v12_: _dafny.Map
            d_880_v12_ = _dafny.Map({self.f24: d_866_v0_})
            d_881_v13_: _dafny.Map
            d_881_v13_ = _dafny.Map({(d_876_v8_)[default__.safeIndex(166, (d_876_v8_).length(0))]: True})
            (globalState).f10 = ((d_879_v11_)[((d_880_v12_)[(self).f32] if ((self).f32) in (d_880_v12_) else len(d_868_v2_))] if (((d_880_v12_)[(self).f32] if ((self).f32) in (d_880_v12_) else len(d_868_v2_))) in (d_879_v11_) else (d_866_v0_) != (len(d_881_v13_)))
            (globalState).f11 = default__.safeDivisionInt(len((self).f33), (0) - ((0) - (((0) - (d_866_v0_)) - ((0) - (d_866_v0_)))))
            (self).f27 = _dafny.CodePoint('r')
            d_882_v14_: T1
            nw131_ = C1()
            nw131_.ctor__(not((self).f26))
            d_882_v14_ = nw131_
            d_883_v15_: _dafny.MultiSet
            d_883_v15_ = _dafny.MultiSet([d_882_v14_])
            d_883_v15_ = d_883_v15_
            d_884_v16_: _dafny.Array
            nw132_ = _dafny.Array(_dafny.Map({}), 18)
            d_884_v16_ = nw132_
            d_885_v17_: _dafny.Array
            nw133_ = _dafny.Array(None, 3)
            nw133_[int(0)] = (self).f32
            nw133_[int(1)] = (self).f26
            nw133_[int(2)] = self.f24
            d_885_v17_ = nw133_
            rhs99_ = (d_885_v17_ if ((self).f25 if (self).f26 else d_882_v14_.f24) else d_885_v17_)
            rhs100_ = d_884_v16_
            rhs101_ = self.f27
            lhs91_ = globalState
            lhs92_ = self
            lhs91_.f2 = rhs99_
            d_884_v16_ = rhs100_
            lhs92_.f27 = rhs101_
            d_886_v18_: _dafny.Seq
            d_886_v18_ = _dafny.SeqWithoutIsStrInference([d_867_v1_])
            d_886_v18_ = d_886_v18_
        (self).f24 = ((self).f26) == (not((d_866_v0_) == (d_866_v0_)))
        d_887_i0_: int
        d_887_i0_ = 0
        with _dafny.label("10"):
            while not ((self).f32) or ((self).f26):
                with _dafny.c_label("10"):
                    if (d_887_i0_) >= (100):
                        raise _dafny.Break("10")
                    d_887_i0_ = (d_887_i0_) + (1)
                    d_888_v19_: _dafny.Map
                    d_888_v19_ = _dafny.Map({(self).f32: len((self).f33)})
                    d_888_v19_ = (d_888_v19_).set((self).fm39(d_868_v2_, self.f24, globalState), -65)
                    (globalState).f16 = d_866_v0_
                    d_889_v20_: _dafny.Map
                    d_889_v20_ = _dafny.Map({(self).f32: (self).f32})
                    d_890_v21_: T1
                    nw134_ = C1()
                    nw134_.ctor__(False)
                    d_890_v21_ = nw134_
                    d_891_v22_: _dafny.Map
                    d_891_v22_ = _dafny.Map({d_890_v21_: d_868_v2_})
                    d_892_v23_: _dafny.MultiSet
                    d_892_v23_ = _dafny.MultiSet([d_866_v0_, (len((self).f33)) + (len(d_889_v20_)), d_866_v0_, len(d_891_v22_), (len(d_888_v19_)) * (d_866_v0_)])
                    d_892_v23_ = _dafny.MultiSet([default__.safeDivisionInt(d_866_v0_, d_866_v0_)])
                    d_893_v24_: int
                    d_894_v25_: _dafny.Seq
                    d_895_v26_: _dafny.Array
                    d_896_v27_: _dafny.Seq
                    out25_: int
                    out26_: _dafny.Seq
                    out27_: _dafny.Array
                    out28_: _dafny.Seq
                    out25_, out26_, out27_, out28_ = (self).m8((self).f25, globalState)
                    d_893_v24_ = out25_
                    d_894_v25_ = out26_
                    d_895_v26_ = out27_
                    d_896_v27_ = out28_
                    pass
            pass
        d_897_v28_: _dafny.MultiSet
        d_897_v28_ = _dafny.MultiSet([self.f24])
        (globalState).f10 = not(((d_897_v28_).set((self).f32, default__.abs(d_866_v0_))).ispropersubset(d_897_v28_))
        hi4_ = d_866_v0_
        for d_898_i1_ in range((self).fm6(D0_DC0(self.f24, (0) - (d_866_v0_), d_866_v0_), globalState), hi4_):
            d_899_v30_: _dafny.Seq
            d_899_v30_ = _dafny.SeqWithoutIsStrInference([(self).f32])
            d_900_v31_: _dafny.Map
            d_900_v31_ = _dafny.Map({d_899_v30_: (self).f26})
            d_901_v32_: _dafny.Map
            d_901_v32_ = _dafny.Map({len(d_900_v31_): (self).f26})
            d_902_v33_: _dafny.Array
            def lambda38_(d_903_i2_):
                return (self).f26

            init21_ = lambda38_
            nw135_ = _dafny.Array(None, 1)
            for i0_21_ in range(nw135_.length(0)):
                nw135_[i0_21_] = init21_(i0_21_)
            d_902_v33_ = nw135_
            d_904_v34_: D6
            d_904_v34_ = D6_DC13(d_902_v33_, d_898_i1_, True)
            d_905_v35_: _dafny.Set
            d_905_v35_ = _dafny.Set({True})
            d_906_v36_: _dafny.Map
            d_906_v36_ = _dafny.Map({(self).f25: d_898_i1_})
            d_907_v37_: _dafny.Array
            nw136_ = _dafny.Array(None, 23)
            nw136_[int(0)] = d_866_v0_
            def iife68_():
                coll36_ = _dafny.Map()
                compr_36_: int
                for compr_36_ in (d_867_v1_).Elements:
                    d_908_v29_: int = compr_36_
                    if (d_908_v29_) in (d_867_v1_):
                        coll36_[default__.safeDivisionInt(d_908_v29_, d_898_i1_)] = (self).f25
                return _dafny.Map(coll36_)
            nw136_[int(1)] = len((iife68_()
            ) | (d_901_v32_))
            nw136_[int(2)] = (d_898_i1_) + (d_898_i1_)
            nw136_[int(3)] = (d_904_v34_).cf25
            nw136_[int(4)] = (0) - (len((_dafny.Set({False})).intersection(d_905_v35_)))
            nw136_[int(5)] = d_866_v0_
            nw136_[int(6)] = len((self).f33)
            nw136_[int(7)] = (d_868_v2_)[default__.safeIndex(len((d_906_v36_).set(self.f24, d_866_v0_)), len(d_868_v2_))]
            nw136_[int(8)] = -41
            nw136_[int(9)] = len((self).f33)
            nw136_[int(10)] = (len(d_901_v32_)) - (d_866_v0_)
            nw136_[int(11)] = d_866_v0_
            nw136_[int(12)] = (D6_DC13(d_902_v33_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wgnjr"))), (self).f26)).cf25
            nw136_[int(13)] = (0) - ((self).fm6(D0_DC0((self).f32, d_866_v0_, d_866_v0_), globalState))
            nw136_[int(14)] = default__.safeDivisionInt(d_898_i1_, 27)
            nw136_[int(15)] = 179
            nw136_[int(16)] = d_866_v0_
            nw136_[int(17)] = d_866_v0_
            nw136_[int(18)] = default__.safeDivisionInt(d_866_v0_, d_898_i1_)
            nw136_[int(19)] = d_898_i1_
            nw136_[int(20)] = len(d_906_v36_)
            nw136_[int(21)] = d_866_v0_
            nw136_[int(22)] = d_898_i1_
            d_907_v37_ = nw136_
            index128_ = default__.safeIndex(477, (d_907_v37_).length(0))
            (d_907_v37_)[index128_] = d_866_v0_
            index129_ = default__.safeIndex(477, (d_907_v37_).length(0))
            (d_907_v37_)[index129_] = d_866_v0_
            d_909_v38_: _dafny.Array
            nw137_ = _dafny.Array(_dafny.Map({}), 23)
            d_909_v38_ = nw137_
            d_910_v39_: D8
            d_910_v39_ = D8_DC18(d_909_v38_)
            d_909_v38_ = (d_910_v39_).cf36
            d_911_v40_: D1
            d_911_v40_ = D1_DC2(d_907_v37_, self.f27, d_898_i1_, (self).f26)
            d_912_v41_: _dafny.Map
            d_912_v41_ = _dafny.Map({not((self).f26): d_911_v40_})
            d_913_v42_: _dafny.Seq
            d_913_v42_ = _dafny.SeqWithoutIsStrInference([d_912_v41_, _dafny.Map({(self).f25: d_911_v40_})])
            (globalState).f10 = ((d_913_v42_) + (d_913_v42_)) != (d_913_v42_)
            if default__.fm1(not((self).f25), False, default__.safeDivisionInt((d_907_v37_)[default__.safeIndex(477, (d_907_v37_).length(0))], d_866_v0_), globalState):
                index130_ = default__.safeIndex(695, (d_902_v33_).length(0))
                (d_902_v33_)[index130_] = (d_902_v33_) not in (_dafny.Map({d_902_v33_: d_866_v0_}))
                index131_ = default__.safeIndex(695, (d_902_v33_).length(0))
                rhs102_ = default__.fm44((0) - (len(d_868_v2_)), d_868_v2_, globalState)
                rhs103_ = d_910_v39_
                rhs104_ = (self).f32
                rhs105_ = ((self).f26 if (self).f25 else (self).f32)
                lhs93_ = d_902_v33_
                lhs94_ = default__.safeIndex(695, (d_902_v33_).length(0))
                lhs95_ = globalState
                d_905_v35_ = rhs102_
                d_910_v39_ = rhs103_
                lhs93_[lhs94_] = rhs104_
                lhs95_.f10 = rhs105_
                (globalState).f10 = (((self).f33) + (_dafny.SeqWithoutIsStrInference([self.f27 for d_914_i3_ in range(default__.abs(925))]))) <= ((self).f33)
                d_915_v43_: _dafny.Array
                nw138_ = _dafny.Array(_dafny.Array(None, 0), 18)
                d_915_v43_ = nw138_
                d_916_v44_: _dafny.Array
                nw139_ = _dafny.Array(_dafny.Array(None, 0), 9)
                d_916_v44_ = nw139_
                index132_ = default__.safeIndex(61, (d_915_v43_).length(0))
                (d_915_v43_)[index132_] = d_916_v44_
                index133_ = default__.safeIndex(61, (d_915_v43_).length(0))
                rhs106_ = (self.f27) not in ((self).f33)
                rhs107_ = ((d_898_i1_ if self.f24 else (d_907_v37_)[default__.safeIndex(477, (d_907_v37_).length(0))])) < (d_866_v0_)
                rhs108_ = (default__.safeModuloInt((d_907_v37_)[default__.safeIndex(477, (d_907_v37_).length(0))], d_898_i1_)) * ((d_868_v2_)[default__.safeIndex(736, len(d_868_v2_))])
                rhs109_ = (d_916_v44_ if self.f24 else d_916_v44_)
                lhs96_ = self
                lhs97_ = globalState
                lhs98_ = globalState
                lhs99_ = d_915_v43_
                lhs100_ = default__.safeIndex(61, (d_915_v43_).length(0))
                lhs96_.f24 = rhs106_
                lhs97_.f10 = rhs107_
                lhs98_.f11 = rhs108_
                lhs99_[lhs100_] = rhs109_
                (globalState).f10 = (self).f32
                (globalState).f7 = d_907_v37_
            elif True:
                rhs110_ = self.f27
                rhs111_ = d_907_v37_
                rhs112_ = (self).f26
                lhs101_ = self
                lhs102_ = globalState
                lhs103_ = globalState
                lhs101_.f27 = rhs110_
                lhs102_.f7 = rhs111_
                lhs103_.f10 = rhs112_
                d_917_v45_: _dafny.Seq
                d_917_v45_ = _dafny.SeqWithoutIsStrInference([(self).f33, (self).f33])
                (globalState).f10 = ((self).f25) or ((d_917_v45_) == (d_917_v45_))
                d_918_v46_: _dafny.Array
                nw140_ = _dafny.Array(_dafny.MultiSet({}), 20)
                d_918_v46_ = nw140_
                d_919_v47_: _dafny.MultiSet
                d_919_v47_ = _dafny.MultiSet([len(d_899_v30_), d_898_i1_])
                index134_ = default__.safeIndex(437, (d_918_v46_).length(0))
                (d_918_v46_)[index134_] = (d_919_v47_) - (d_919_v47_)
                d_920_v48_: _dafny.Array
                def lambda39_(d_921_i4_):
                    return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qvn"))

                init22_ = lambda39_
                nw141_ = _dafny.Array(None, 28)
                for i0_22_ in range(nw141_.length(0)):
                    nw141_[i0_22_] = init22_(i0_22_)
                d_920_v48_ = nw141_
                index135_ = default__.safeIndex(182, (d_920_v48_).length(0))
                (d_920_v48_)[index135_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aus"))
                d_922_v49_: _dafny.Map
                d_922_v49_ = _dafny.Map({(self).fm7(d_906_v36_, globalState): d_897_v28_})
                index136_ = default__.safeIndex(437, (d_918_v46_).length(0))
                index137_ = default__.safeIndex(182, (d_920_v48_).length(0))
                rhs113_ = (self).f26
                rhs114_ = d_919_v47_
                rhs115_ = ((self).f33 if not (not(self.f24)) or ((self).f26) else ((self).f33).set(default__.safeIndex((((d_922_v49_)[(self).f32] if ((self).f32) in (d_922_v49_) else d_897_v28_)).cardinality, len((self).f33)), self.f27))
                lhs104_ = self
                lhs105_ = d_918_v46_
                lhs106_ = default__.safeIndex(437, (d_918_v46_).length(0))
                lhs107_ = d_920_v48_
                lhs108_ = default__.safeIndex(182, (d_920_v48_).length(0))
                lhs104_.f24 = rhs113_
                lhs105_[lhs106_] = rhs114_
                lhs107_[lhs108_] = rhs115_
                (globalState).f11 = (len(d_868_v2_)) - (600)
                d_923_v50_: _dafny.Map
                d_923_v50_ = _dafny.Map({(self).f25: not((self).f26)})
                d_924_v51_: _dafny.Seq
                d_924_v51_ = _dafny.SeqWithoutIsStrInference([d_923_v50_, _dafny.Map({self.f24: (self).f32})])
                d_925_v52_: _dafny.Seq
                d_925_v52_ = _dafny.SeqWithoutIsStrInference([(d_924_v51_)[default__.safeIndex((d_907_v37_)[default__.safeIndex(477, (d_907_v37_).length(0))], len(d_924_v51_))], _dafny.Map({(d_899_v30_)[default__.safeIndex(d_866_v0_, len(d_899_v30_))]: (self).f32})])
                d_926_v53_: C2
                nw142_ = C2()
                nw142_.ctor__((d_920_v48_)[default__.safeIndex(182, (d_920_v48_).length(0))], ((d_917_v45_)[default__.safeIndex(((d_897_v28_)[not(not((self).f32))] if (not(not((self).f32))) in (d_897_v28_) else len(d_924_v51_)), len(d_917_v45_))]) + ((d_920_v48_)[default__.safeIndex(182, (d_920_v48_).length(0))]))
                d_926_v53_ = nw142_
        d_927_v54_: _dafny.MultiSet
        d_927_v54_ = _dafny.MultiSet([d_866_v0_])
        d_928_v55_: _dafny.Map
        d_928_v55_ = _dafny.Map({(d_927_v54_).cardinality: d_866_v0_})
        d_929_v56_: _dafny.Map
        d_929_v56_ = _dafny.Map({(self).f25: d_928_v55_})
        d_930_v57_: _dafny.Map
        d_930_v57_ = _dafny.Map({d_929_v56_: d_866_v0_})
        d_931_v58_: D2
        d_931_v58_ = D2_DC5(self.f27, True, ((d_930_v57_)[d_929_v56_] if (d_929_v56_) in (d_930_v57_) else d_866_v0_), d_866_v0_)
        d_932_i5_: int
        d_932_i5_ = 0
        with _dafny.label("11"):
            while (d_931_v58_).cf11:
                with _dafny.c_label("11"):
                    if (d_932_i5_) >= (100):
                        raise _dafny.Break("11")
                    d_932_i5_ = (d_932_i5_) + (1)
                    d_933_v59_: _dafny.Map
                    d_933_v59_ = _dafny.Map({(self).f25: False})
                    (self).f27 = ((self).f33)[default__.safeIndex(default__.safeModuloInt(len(d_933_v59_), d_866_v0_), len((self).f33))]
                    d_934_v60_: int
                    d_935_v61_: _dafny.Seq
                    d_936_v62_: _dafny.Array
                    d_937_v63_: _dafny.Seq
                    out29_: int
                    out30_: _dafny.Seq
                    out31_: _dafny.Array
                    out32_: _dafny.Seq
                    out29_, out30_, out31_, out32_ = (self).m8((self).f26, globalState)
                    d_934_v60_ = out29_
                    d_935_v61_ = out30_
                    d_936_v62_ = out31_
                    d_937_v63_ = out32_
                    (globalState).f16 = d_866_v0_
                    index138_ = default__.safeIndex(268, (d_936_v62_).length(0))
                    (d_936_v62_)[index138_] = d_866_v0_
                    index139_ = default__.safeIndex(268, (d_936_v62_).length(0))
                    rhs116_ = d_934_v60_
                    rhs117_ = default__.safeDivisionInt(648, d_934_v60_)
                    lhs109_ = globalState
                    lhs110_ = d_936_v62_
                    lhs111_ = default__.safeIndex(268, (d_936_v62_).length(0))
                    lhs109_.f16 = rhs116_
                    lhs110_[lhs111_] = rhs117_
                    pass
            pass

    @property
    def f32(self):
        return self._f32
    @property
    def f33(self):
        return self._f33

class C4(T2, T1, T0):
    def  __init__(self):
        self._f24: bool = False
        self._f25: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    @property
    def f24(self):
        return self._f24
    @f24.setter
    def f24(self, value):
        self._f24 = value
    @property
    def f25(self):
        return self._f25
    def ctor__(self, f25, f24):
        (self)._f25 = f25
        (self).f24 = f24

    def fm6(self, p0, globalState):
        return -901

    def fm4(self, p0, p1, globalState):
        return D0_DC0(True, (653) - (len(_dafny.Map({476: len(_dafny.SeqWithoutIsStrInference([False]))}))), len(_dafny.Map({(self).f25: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "df")))})))

    def fm5(self, p0, p1, p2, globalState):
        return len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u")))

    def fm37(self, p0, globalState):
        def iife69_():
            coll37_ = _dafny.Set()
            compr_37_: str
            for compr_37_ in (_dafny.Map({_dafny.CodePoint('r'): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iwaa")))})).keys.Elements:
                d_939_v0_: str = compr_37_
                if (d_939_v0_) in (_dafny.Map({_dafny.CodePoint('r'): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iwaa")))})):
                    coll37_ = coll37_.union(_dafny.Set([d_939_v0_]))
            return _dafny.Set(coll37_)
        def iife70_():
            coll38_ = _dafny.Set()
            compr_38_: str
            for compr_38_ in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f'), _dafny.CodePoint('p')]))).Elements:
                d_940_v1_: str = compr_38_
                if (d_940_v1_) in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f'), _dafny.CodePoint('p')]))):
                    coll38_ = coll38_.union(_dafny.Set([d_940_v1_]))
            return _dafny.Set(coll38_)
        return ((len(_dafny.Map({False: -836})) if (self).f25 else len(_dafny.SeqWithoutIsStrInference([_dafny.Set({(_dafny.MultiSet([self.f24])).cardinality}) for d_938_i0_ in range(default__.abs(-564))])))) * (len((iife69_()
        ).intersection(iife70_()
        )))

    def m0(self, globalState):
        r0: bool = False
        r1: _dafny.Array = _dafny.Array(None, 0)
        (self).f24 = not (self.f24) or (((self).f25) or (self.f24))
        d_941_v0_: _dafny.Array
        nw143_ = _dafny.Array(int(0), 6)
        d_941_v0_ = nw143_
        d_942_v1_: str
        d_942_v1_ = _dafny.CodePoint('n')
        d_943_v2_: int
        d_943_v2_ = 336
        d_944_v3_: D1
        d_944_v3_ = D1_DC2(d_941_v0_, d_942_v1_, d_943_v2_, self.f24)
        source18_ = (d_944_v3_ if self.f24 else D1_DC2(d_941_v0_, d_942_v1_, d_943_v2_, self.f24))
        if source18_.is_DC2:
            d_945___mcc_h0_ = source18_.cf4
            d_946___mcc_h1_ = source18_.cf5
            d_947___mcc_h2_ = source18_.cf6
            d_948___mcc_h3_ = source18_.cf7
            d_949_cf7_ = d_948___mcc_h3_
            d_950_cf6_ = d_947___mcc_h2_
            d_951_cf5_ = d_946___mcc_h1_
            d_952_cf4_ = d_945___mcc_h0_
            (globalState).f11 = d_943_v2_
            d_953_v4_: _dafny.Array
            nw144_ = _dafny.Array(_dafny.CodePoint('D'), 29)
            d_953_v4_ = nw144_
            d_954_v5_: _dafny.Array
            nw145_ = _dafny.Array(None, 13)
            nw145_[int(0)] = d_953_v4_
            nw145_[int(1)] = d_953_v4_
            nw145_[int(2)] = d_953_v4_
            nw145_[int(3)] = d_953_v4_
            nw145_[int(4)] = (d_953_v4_ if d_949_cf7_ else d_953_v4_)
            nw145_[int(5)] = d_953_v4_
            nw145_[int(6)] = d_953_v4_
            nw145_[int(7)] = d_953_v4_
            nw145_[int(8)] = d_953_v4_
            nw145_[int(9)] = d_953_v4_
            nw145_[int(10)] = d_953_v4_
            nw145_[int(11)] = (d_953_v4_ if False else d_953_v4_)
            nw145_[int(12)] = d_953_v4_
            d_954_v5_ = nw145_
            d_955_v6_: _dafny.Map
            d_955_v6_ = _dafny.Map({not(self.f24): d_953_v4_})
            d_956_v7_: _dafny.Seq
            d_956_v7_ = _dafny.SeqWithoutIsStrInference([d_953_v4_, d_953_v4_, d_953_v4_, ((d_955_v6_)[d_949_cf7_] if (d_949_cf7_) in (d_955_v6_) else d_953_v4_), d_953_v4_])
            index140_ = default__.safeIndex(431, (d_954_v5_).length(0))
            (d_954_v5_)[index140_] = (d_956_v7_)[default__.safeIndex(d_950_cf6_, len(d_956_v7_))]
            index141_ = default__.safeIndex(431, (d_954_v5_).length(0))
            (d_954_v5_)[index141_] = d_953_v4_
            d_957_v8_: _dafny.Seq
            d_957_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m"))
            if (len(d_957_v8_)) > (default__.safeModuloInt(d_943_v2_, d_950_cf6_)):
                d_958_v9_: _dafny.MultiSet
                d_958_v9_ = _dafny.MultiSet([d_949_cf7_])
                d_959_v10_: C1
                nw146_ = C1()
                nw146_.ctor__((d_958_v9_).ispropersubset(d_958_v9_))
                d_959_v10_ = nw146_
                d_960_v11_: _dafny.Set
                d_960_v11_ = _dafny.Set({_dafny.CodePoint('q'), d_942_v1_})
                d_961_v12_: _dafny.Map
                d_961_v12_ = _dafny.Map({self.f24: self.f24})
                d_962_v13_: _dafny.Map
                d_962_v13_ = _dafny.Map({len(d_960_v11_): len(d_961_v12_)})
                d_963_v14_: _dafny.Set
                d_963_v14_ = _dafny.Set({self.f24})
                d_964_v15_: _dafny.Seq
                d_964_v15_ = _dafny.SeqWithoutIsStrInference([not(d_949_cf7_), (self).f25])
                d_962_v13_ = (d_962_v13_).set((d_943_v2_) * (d_943_v2_), default__.safeDivisionInt(len(d_963_v14_), len(d_964_v15_)))
                (globalState).f10 = d_949_cf7_
                (globalState).f11 = ((d_943_v2_) + ((d_958_v9_).cardinality)) + ((d_950_cf6_ if d_949_cf7_ else 433))
                d_965_v16_: _dafny.Array
                nw147_ = _dafny.Array(False, 6)
                d_965_v16_ = nw147_
                d_966_v17_: _dafny.Map
                d_966_v17_ = _dafny.Map({d_965_v16_: (d_954_v5_)[default__.safeIndex(431, (d_954_v5_).length(0))]})
                index142_ = default__.safeIndex(431, (d_954_v5_).length(0))
                (d_954_v5_)[index142_] = ((d_966_v17_)[d_965_v16_] if (d_965_v16_) in (d_966_v17_) else (d_954_v5_)[default__.safeIndex(431, (d_954_v5_).length(0))])
            elif True:
                d_967_v18_: D7
                d_967_v18_ = D7_DC15(_dafny.CodePoint('g'), d_949_cf7_)
                d_968_v19_: _dafny.Map
                d_968_v19_ = _dafny.Map({(self).f25: d_967_v18_})
                d_969_v20_: D7
                d_969_v20_ = D7_DC17(((d_968_v19_)[(self).f25] if ((self).f25) in (d_968_v19_) else d_967_v18_))
                d_970_v21_: _dafny.Set
                d_970_v21_ = _dafny.Set({d_969_v20_})
                (globalState).f16 = len(d_970_v21_)
                d_971_v22_: _dafny.Array
                nw148_ = _dafny.Array(None, 9)
                nw148_[int(0)] = (len(d_957_v8_)) != (d_943_v2_)
                nw148_[int(1)] = default__.fm1(d_949_cf7_, True, d_950_cf6_, globalState)
                nw148_[int(2)] = (self).f25
                nw148_[int(3)] = True
                nw148_[int(4)] = (d_950_cf6_) == (215)
                nw148_[int(5)] = (self).f25
                nw148_[int(6)] = True
                nw148_[int(7)] = not(self.f24)
                nw148_[int(8)] = not((self).f25)
                d_971_v22_ = nw148_
                index143_ = default__.safeIndex(92, (d_971_v22_).length(0))
                (d_971_v22_)[index143_] = True
                index144_ = default__.safeIndex(92, (d_971_v22_).length(0))
                (d_971_v22_)[index144_] = (self).f25
                d_972_v23_: _dafny.MultiSet
                d_972_v23_ = _dafny.MultiSet([d_949_cf7_, (354) > (d_950_cf6_)])
                d_973_v24_: _dafny.Seq
                d_973_v24_ = _dafny.SeqWithoutIsStrInference([(d_971_v22_)[default__.safeIndex(92, (d_971_v22_).length(0))]])
                d_972_v23_ = ((d_972_v23_) | (_dafny.MultiSet(d_973_v24_))) | (d_972_v23_)
                rhs118_ = ((d_971_v22_)[default__.safeIndex(92, (d_971_v22_).length(0))] if self.f24 else not(d_949_cf7_))
                rhs119_ = d_950_cf6_
                lhs112_ = globalState
                d_949_cf7_ = rhs118_
                lhs112_.f16 = rhs119_
                d_974_v25_: _dafny.Map
                d_974_v25_ = _dafny.Map({202: self.f24})
                d_975_v26_: D6
                d_975_v26_ = D6_DC13(d_971_v22_, d_950_cf6_, (d_971_v22_)[default__.safeIndex(92, (d_971_v22_).length(0))])
                d_976_v27_: _dafny.Map
                d_976_v27_ = _dafny.Map({self.f24: d_957_v8_})
                d_977_v28_: _dafny.Map
                d_977_v28_ = _dafny.Map({(d_971_v22_)[default__.safeIndex(92, (d_971_v22_).length(0))]: _dafny.SeqWithoutIsStrInference([D0_DC0(default__.fm1(self.f24, d_949_cf7_, d_943_v2_, globalState), len(((d_976_v27_)[(d_971_v22_)[default__.safeIndex(92, (d_971_v22_).length(0))]] if ((d_971_v22_)[default__.safeIndex(92, (d_971_v22_).length(0))]) in (d_976_v27_) else d_957_v8_)), len(d_957_v8_))])})
                d_978_v29_: _dafny.Map
                d_978_v29_ = _dafny.Map({((d_974_v25_)[(d_975_v26_).cf25] if ((d_975_v26_).cf25) in (d_974_v25_) else self.f24): (d_943_v2_) - ((self).fm5(d_977_v28_, self.f24, self.f24, globalState))})
                d_979_v30_: _dafny.Seq
                d_979_v30_ = _dafny.SeqWithoutIsStrInference([d_973_v24_, d_973_v24_, d_973_v24_])
                d_978_v29_ = (d_978_v29_).set(not(not(not(d_949_cf7_))), default__.safeDivisionInt(len(d_979_v30_), 282))
            (self).f24 = False
        elif True:
            d_980___mcc_h4_ = source18_.cf3
            d_981_cf3_ = d_980___mcc_h4_
            d_982_v31_: D0
            d_982_v31_ = D0_DC0(False, d_943_v2_, (0) - (d_943_v2_))
            rhs120_ = d_943_v2_
            rhs121_ = default__.fm1((self).f25, False, (self).fm6(d_982_v31_, globalState), globalState)
            lhs113_ = globalState
            lhs114_ = globalState
            lhs113_.f22 = rhs120_
            lhs114_.f10 = rhs121_
            d_983_v32_: _dafny.MultiSet
            d_983_v32_ = _dafny.MultiSet([d_943_v2_, 522, d_943_v2_])
            d_984_v33_: _dafny.Seq
            d_984_v33_ = _dafny.SeqWithoutIsStrInference([-219])
            d_985_v34_: C2
            nw149_ = C2()
            nw149_.ctor__(d_981_cf3_, (default__.fm38((self).f25, d_943_v2_, d_983_v32_, _dafny.Map({(self).f25: len(d_984_v33_)}), globalState)) + (d_981_cf3_))
            d_985_v34_ = nw149_
            d_986_v35_: _dafny.Array
            def lambda40_(d_987_v1_, d_988_v2_):
                def lambda41_(d_989_i0_):
                    return _dafny.Map({self.f24: D2_DC5(d_987_v1_, (self).f25, d_988_v2_, d_988_v2_)})

                return lambda41_

            init23_ = lambda40_(d_942_v1_, d_943_v2_)
            nw150_ = _dafny.Array(None, 8)
            for i0_23_ in range(nw150_.length(0)):
                nw150_[i0_23_] = init23_(i0_23_)
            d_986_v35_ = nw150_
            d_990_v36_: D2
            d_990_v36_ = D2_DC5(d_942_v1_, (self).f25, d_943_v2_, 751)
            d_991_v37_: _dafny.Map
            d_991_v37_ = _dafny.Map({(self).f25: d_990_v36_})
            index145_ = default__.safeIndex(396, (d_986_v35_).length(0))
            (d_986_v35_)[index145_] = (d_991_v37_ if False else d_991_v37_)
            index146_ = default__.safeIndex(396, (d_986_v35_).length(0))
            rhs122_ = len(_dafny.SeqWithoutIsStrInference([self.f24, (self).f25, not((self).f25)]))
            rhs123_ = (self).f25
            rhs124_ = d_942_v1_
            rhs125_ = (d_943_v2_ if (self).f25 else d_943_v2_)
            rhs126_ = d_991_v37_
            lhs115_ = globalState
            lhs116_ = globalState
            lhs117_ = d_986_v35_
            lhs118_ = default__.safeIndex(396, (d_986_v35_).length(0))
            lhs115_.f11 = rhs122_
            r0 = rhs123_
            d_942_v1_ = rhs124_
            lhs116_.f16 = rhs125_
            lhs117_[lhs118_] = rhs126_
            nw151_ = _dafny.Array(int(0), 26)
            (globalState).f7 = nw151_
        d_992_v38_: _dafny.Seq
        d_992_v38_ = _dafny.SeqWithoutIsStrInference([d_943_v2_])
        hi5_ = d_943_v2_
        for d_993_i1_ in range((d_992_v38_)[default__.safeIndex(d_943_v2_, len(d_992_v38_))], hi5_):
            source19_ = D1_DC2((d_944_v3_).cf4, d_942_v1_, d_943_v2_, self.f24)
            if source19_.is_DC2:
                d_994___mcc_h5_ = source19_.cf4
                d_995___mcc_h6_ = source19_.cf5
                d_996___mcc_h7_ = source19_.cf6
                d_997___mcc_h8_ = source19_.cf7
                d_998_cf7_ = d_997___mcc_h8_
                d_999_cf6_ = d_996___mcc_h7_
                d_1000_cf5_ = d_995___mcc_h6_
                d_1001_cf4_ = d_994___mcc_h5_
                (globalState).f10 = d_998_cf7_
                d_1002_v39_: _dafny.Seq
                d_1002_v39_ = _dafny.SeqWithoutIsStrInference([self.f24])
                d_1003_v40_: _dafny.Map
                d_1003_v40_ = _dafny.Map({len((d_1002_v39_).set(default__.safeIndex(d_999_cf6_, len(d_1002_v39_)), self.f24)): d_999_cf6_})
                d_1004_v42_: _dafny.Seq
                d_1004_v42_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aiu"))
                d_1005_v43_: _dafny.Map
                d_1005_v43_ = _dafny.Map({d_999_cf6_: d_1004_v42_})
                d_1006_v44_: _dafny.Seq
                d_1006_v44_ = _dafny.SeqWithoutIsStrInference([d_1004_v42_, d_1004_v42_, d_1004_v42_, d_1004_v42_])
                def iife71_():
                    coll39_ = _dafny.Map()
                    compr_39_: int
                    for compr_39_ in (d_1005_v43_).keys.Elements:
                        d_1007_v41_: int = compr_39_
                        if (d_1007_v41_) in (d_1005_v43_):
                            coll39_[(d_1007_v41_) + (-536)] = d_993_i1_
                    return _dafny.Map(coll39_)
                d_1003_v40_ = (iife71_()
                ) | (_dafny.Map({len(d_1004_v42_): len(d_1006_v44_)}))
                d_1008_v45_: _dafny.Array
                nw152_ = _dafny.Array(_dafny.Seq({}), 26)
                d_1008_v45_ = nw152_
                d_1008_v45_ = d_1008_v45_
                d_1009_v46_: _dafny.Array
                def lambda42_(d_1010_i2_):
                    return (self).f25

                init24_ = lambda42_
                nw153_ = _dafny.Array(None, 26)
                for i0_24_ in range(nw153_.length(0)):
                    nw153_[i0_24_] = init24_(i0_24_)
                d_1009_v46_ = nw153_
                d_1011_v47_: _dafny.MultiSet
                d_1011_v47_ = _dafny.MultiSet([d_1009_v46_, d_1009_v46_, d_1009_v46_, d_1009_v46_, d_1009_v46_])
                d_1012_v48_: _dafny.Seq
                d_1012_v48_ = _dafny.SeqWithoutIsStrInference([d_1011_v47_, d_1011_v47_])
                rhs127_ = self.f24
                rhs128_ = ((d_1012_v48_)[default__.safeIndex(d_999_cf6_, len(d_1012_v48_))] if d_998_cf7_ else d_1011_v47_)
                rhs129_ = not ((self).f25) or ((self).f25)
                rhs130_ = _dafny.SeqWithoutIsStrInference([d_1000_cf5_ for d_1013_i3_ in range(default__.abs(874))])
                lhs119_ = globalState
                lhs119_.f10 = rhs127_
                d_1011_v47_ = rhs128_
                d_998_cf7_ = rhs129_
                d_1004_v42_ = rhs130_
            elif True:
                d_1014___mcc_h9_ = source19_.cf3
                d_1015_cf3_ = d_1014___mcc_h9_
                d_1016_v49_: D2
                d_1016_v49_ = D2_DC5(d_942_v1_, self.f24, d_943_v2_, d_943_v2_)
                (self).f24 = (d_1016_v49_).cf11
                d_1017_v50_: _dafny.Map
                d_1017_v50_ = _dafny.Map({d_943_v2_: d_993_i1_})
                d_1017_v50_ = (d_1017_v50_).set(default__.safeModuloInt(d_943_v2_, (0) - (len(_dafny.SeqWithoutIsStrInference([(self).f25])))), len((d_1015_cf3_) + (d_1015_cf3_)))
                d_1018_v51_: _dafny.Map
                d_1018_v51_ = _dafny.Map({219: (self).f25})
                d_1018_v51_ = d_1018_v51_
                d_1019_v52_: C0
                nw154_ = C0()
                nw154_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bviylf")))
                d_1019_v52_ = nw154_
                d_1020_v53_: _dafny.Array
                nw155_ = _dafny.Array(False, 23)
                d_1020_v53_ = nw155_
                d_1021_v54_: _dafny.Map
                d_1021_v54_ = _dafny.Map({d_1020_v53_: d_1019_v52_})
                d_1022_v55_: _dafny.Seq
                d_1022_v55_ = _dafny.SeqWithoutIsStrInference([d_1020_v53_])
                d_1023_v56_: _dafny.Array
                nw156_ = _dafny.Array(None, 19)
                nw156_[int(0)] = d_1019_v52_
                nw156_[int(1)] = d_1019_v52_
                nw156_[int(2)] = d_1019_v52_
                nw156_[int(3)] = d_1019_v52_
                nw156_[int(4)] = d_1019_v52_
                nw156_[int(5)] = d_1019_v52_
                nw156_[int(6)] = d_1019_v52_
                nw156_[int(7)] = d_1019_v52_
                nw156_[int(8)] = d_1019_v52_
                nw156_[int(9)] = d_1019_v52_
                nw156_[int(10)] = d_1019_v52_
                nw156_[int(11)] = ((d_1021_v54_)[(d_1022_v55_)[default__.safeIndex(d_943_v2_, len(d_1022_v55_))]] if ((d_1022_v55_)[default__.safeIndex(d_943_v2_, len(d_1022_v55_))]) in (d_1021_v54_) else d_1019_v52_)
                nw156_[int(12)] = d_1019_v52_
                nw156_[int(13)] = d_1019_v52_
                nw156_[int(14)] = d_1019_v52_
                nw156_[int(15)] = d_1019_v52_
                nw156_[int(16)] = d_1019_v52_
                nw156_[int(17)] = d_1019_v52_
                nw156_[int(18)] = d_1019_v52_
                d_1023_v56_ = nw156_
                rhs131_ = d_1023_v56_
                rhs132_ = -209
                rhs133_ = (self).f25
                rhs134_ = (0) - (d_993_i1_)
                lhs120_ = globalState
                lhs121_ = globalState
                lhs122_ = globalState
                d_1023_v56_ = rhs131_
                lhs120_.f22 = rhs132_
                lhs121_.f10 = rhs133_
                lhs122_.f11 = rhs134_
            d_1024_v57_: _dafny.Array
            nw157_ = _dafny.Array(False, 2)
            d_1024_v57_ = nw157_
            index147_ = default__.safeIndex(506, (d_1024_v57_).length(0))
            (d_1024_v57_)[index147_] = self.f24
            index148_ = default__.safeIndex(506, (d_1024_v57_).length(0))
            (d_1024_v57_)[index148_] = False
            d_1025_v58_: _dafny.Map
            d_1025_v58_ = _dafny.Map({(-698 if not(True) else d_993_i1_): (d_1024_v57_)[default__.safeIndex(506, (d_1024_v57_).length(0))]})
            d_1025_v58_ = d_1025_v58_
            d_1026_v59_: _dafny.Set
            d_1026_v59_ = _dafny.Set({d_941_v0_})
            index149_ = default__.safeIndex(506, (d_1024_v57_).length(0))
            (d_1024_v57_)[index149_] = not ((d_1026_v59_).ispropersubset(d_1026_v59_)) or (self.f24)
        d_1027_v60_: T3
        nw158_ = C3()
        nw158_.ctor__((self).f25, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k")), not(self.f24), d_942_v1_, (self).f25, (self).f25)
        d_1027_v60_ = nw158_
        d_1028_v61_: _dafny.Seq
        d_1028_v61_ = _dafny.SeqWithoutIsStrInference([d_1027_v60_, d_1027_v60_])
        d_1029_v62_: D2
        d_1029_v62_ = D2_DC4((d_1027_v60_).f26)
        d_1030_v63_: _dafny.Map
        d_1030_v63_ = _dafny.Map({(d_1028_v61_)[default__.safeIndex(d_943_v2_, len(d_1028_v61_))]: d_1029_v62_})
        d_1030_v63_ = d_1030_v63_
        d_1031_v64_: _dafny.Array
        nw159_ = _dafny.Array(False, 24)
        d_1031_v64_ = nw159_
        d_1032_v65_: _dafny.Seq
        d_1032_v65_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cxyt"))
        d_1033_v66_: T0
        nw160_ = C2()
        nw160_.ctor__(d_1032_v65_, d_1032_v65_)
        d_1033_v66_ = nw160_
        d_1034_v67_: _dafny.Set
        d_1034_v67_ = _dafny.Set({d_1033_v66_, d_1033_v66_, d_1033_v66_, d_1033_v66_})
        index150_ = default__.safeIndex(223, (d_1031_v64_).length(0))
        (d_1031_v64_)[index150_] = (d_1034_v67_).ispropersubset(d_1034_v67_)
        index151_ = default__.safeIndex(223, (d_1031_v64_).length(0))
        (d_1031_v64_)[index151_] = (d_1027_v60_).f26
        d_1035_v68_: _dafny.MultiSet
        d_1035_v68_ = _dafny.MultiSet([d_943_v2_])
        hi6_ = (d_943_v2_) - (d_943_v2_)
        for d_1036_i4_ in range((d_1035_v68_).cardinality, hi6_):
            d_1037_v69_: _dafny.Array
            def lambda43_(d_1038_i4_, d_1039_v2_):
                def lambda44_(d_1040_i5_):
                    return D2_DC3(_dafny.Map({d_1038_i4_: d_1039_v2_}))

                return lambda44_

            init25_ = lambda43_(d_1036_i4_, d_943_v2_)
            nw161_ = _dafny.Array(None, 12)
            for i0_25_ in range(nw161_.length(0)):
                nw161_[i0_25_] = init25_(i0_25_)
            d_1037_v69_ = nw161_
            d_1037_v69_ = d_1037_v69_
            (globalState).f16 = default__.safeModuloInt(d_943_v2_, d_943_v2_)
            d_1041_v70_: C0
            nw162_ = C0()
            nw162_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ekqp")))
            d_1041_v70_ = nw162_
            d_1042_v71_: _dafny.Array
            nw163_ = _dafny.Array(_dafny.MultiSet({}), 14)
            d_1042_v71_ = nw163_
            nw164_ = _dafny.Array(_dafny.MultiSet({}), 22)
            d_1042_v71_ = nw164_
        r0 = d_1027_v60_.f24
        r1 = d_1031_v64_
        return r0, r1

    def m1(self, p0, globalState):
        r0: str = _dafny.CodePoint('D')
        r1: _dafny.Seq = _dafny.Seq({})
        d_1043_v0_: int
        d_1043_v0_ = 381
        d_1044_v1_: _dafny.Seq
        d_1044_v1_ = _dafny.SeqWithoutIsStrInference([d_1043_v0_])
        d_1044_v1_ = (d_1044_v1_) + (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_1045_i1_ in range(default__.abs(810))])) for d_1046_i0_ in range(default__.abs(-546))]))
        d_1047_v2_: _dafny.Seq
        d_1047_v2_ = _dafny.SeqWithoutIsStrInference([(self).f25, self.f24])
        d_1048_v3_: D0
        d_1048_v3_ = D0_DC0((self).f25, 551, d_1043_v0_)
        d_1049_v4_: _dafny.Seq
        d_1049_v4_ = _dafny.SeqWithoutIsStrInference([d_1048_v3_, d_1048_v3_])
        r1 = ((d_1047_v2_).set(default__.safeIndex((self).fm5(_dafny.Map({(self).f25: d_1049_v4_}), self.f24, self.f24, globalState), len(d_1047_v2_)), self.f24)) + (_dafny.SeqWithoutIsStrInference([(self).f25]))
        hi7_ = len((p0).intersection(_dafny.Set({len(d_1047_v2_)})))
        for d_1050_i2_ in range((0) - ((0) - (d_1043_v0_)), hi7_):
            d_1051_v5_: _dafny.Array
            def lambda45_(d_1052_i3_):
                return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_1053_i4_ in range(default__.abs(277))])

            init26_ = lambda45_
            nw165_ = _dafny.Array(None, 25)
            for i0_26_ in range(nw165_.length(0)):
                nw165_[i0_26_] = init26_(i0_26_)
            d_1051_v5_ = nw165_
            d_1054_v6_: _dafny.Map
            d_1054_v6_ = _dafny.Map({d_1051_v5_: d_1043_v0_})
            d_1054_v6_ = (d_1054_v6_).set(d_1051_v5_, d_1050_i2_)
            d_1055_v7_: _dafny.Array
            def lambda46_(d_1056_p0_):
                def lambda47_(d_1057_i5_):
                    return (d_1056_p0_) | (d_1056_p0_)

                return lambda47_

            init27_ = lambda46_(p0)
            nw166_ = _dafny.Array(None, 24)
            for i0_27_ in range(nw166_.length(0)):
                nw166_[i0_27_] = init27_(i0_27_)
            d_1055_v7_ = nw166_
            index152_ = default__.safeIndex(613, (d_1055_v7_).length(0))
            (d_1055_v7_)[index152_] = (p0) - (p0)
            d_1058_v8_: _dafny.Map
            d_1058_v8_ = _dafny.Map({(0) - (d_1050_i2_): _dafny.SeqWithoutIsStrInference([self.f24, (self).f25, True])})
            d_1059_v9_: _dafny.Seq
            d_1059_v9_ = _dafny.SeqWithoutIsStrInference([d_1047_v2_])
            index153_ = default__.safeIndex(613, (d_1055_v7_).length(0))
            (d_1055_v7_)[index153_] = (p0) | ((_dafny.Set({d_1050_i2_})) | (_dafny.Set({len(((d_1058_v8_)[d_1050_i2_] if (d_1050_i2_) in (d_1058_v8_) else (d_1059_v9_)[default__.safeIndex(d_1043_v0_, len(d_1059_v9_))]))})))
            d_1060_v10_: _dafny.Seq
            d_1060_v10_ = _dafny.SeqWithoutIsStrInference([(default__.fm42(globalState)).set(default__.safeIndex(d_1050_i2_, len(default__.fm42(globalState))), d_1050_i2_)])
            d_1061_v11_: _dafny.Seq
            d_1061_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hueldcmmw"))
            d_1044_v1_ = (d_1060_v10_)[default__.safeIndex(len((d_1061_v11_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "opovedu")))), len(d_1060_v10_))]
            d_1062_v12_: _dafny.Array
            nw167_ = _dafny.Array(None, 2)
            nw167_[int(0)] = 89
            nw167_[int(1)] = (len(d_1044_v1_)) * (d_1050_i2_)
            d_1062_v12_ = nw167_
            index154_ = default__.safeIndex(118, (d_1062_v12_).length(0))
            (d_1062_v12_)[index154_] = d_1043_v0_
            index155_ = default__.safeIndex(118, (d_1062_v12_).length(0))
            (d_1062_v12_)[index155_] = d_1050_i2_
        if (d_1043_v0_) != ((d_1043_v0_ if self.f24 else d_1043_v0_)):
            d_1063_v13_: _dafny.Seq
            d_1063_v13_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v"))
            d_1063_v13_ = d_1063_v13_
            d_1064_v14_: _dafny.Array
            def lambda48_(d_1065_v2_, d_1066_v0_):
                def lambda49_(d_1067_i6_):
                    return (_dafny.Set({(self).f25})) - (_dafny.Set({(d_1065_v2_)[default__.safeIndex(d_1066_v0_, len(d_1065_v2_))]}))

                return lambda49_

            init28_ = lambda48_(d_1047_v2_, d_1043_v0_)
            nw168_ = _dafny.Array(None, 23)
            for i0_28_ in range(nw168_.length(0)):
                nw168_[i0_28_] = init28_(i0_28_)
            d_1064_v14_ = nw168_
            d_1068_v15_: _dafny.Set
            d_1068_v15_ = _dafny.Set({(self).f25, False})
            index156_ = default__.safeIndex(545, (d_1064_v14_).length(0))
            (d_1064_v14_)[index156_] = d_1068_v15_
            index157_ = default__.safeIndex(545, (d_1064_v14_).length(0))
            (d_1064_v14_)[index157_] = (d_1068_v15_).intersection((default__.fm44(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fitm"))), d_1044_v1_, globalState)) | (d_1068_v15_))
            d_1069_v16_: _dafny.Array
            def lambda50_(d_1070_i7_):
                return (self).f25

            init29_ = lambda50_
            nw169_ = _dafny.Array(None, 4)
            for i0_29_ in range(nw169_.length(0)):
                nw169_[i0_29_] = init29_(i0_29_)
            d_1069_v16_ = nw169_
            index158_ = default__.safeIndex(149, (d_1069_v16_).length(0))
            (d_1069_v16_)[index158_] = self.f24
            index159_ = default__.safeIndex(149, (d_1069_v16_).length(0))
            rhs135_ = (self).f25
            rhs136_ = d_1043_v0_
            lhs123_ = d_1069_v16_
            lhs124_ = default__.safeIndex(149, (d_1069_v16_).length(0))
            lhs123_[lhs124_] = rhs135_
            d_1043_v0_ = rhs136_
            d_1071_v17_: str
            d_1071_v17_ = _dafny.CodePoint('p')
            d_1063_v13_ = (d_1063_v13_).set(default__.safeIndex(d_1043_v0_, len(d_1063_v13_)), d_1071_v17_)
            d_1072_v18_: _dafny.Array
            def lambda51_(d_1073_v1_):
                def lambda52_(d_1074_i8_):
                    return (d_1073_v1_) + (d_1073_v1_)

                return lambda52_

            init30_ = lambda51_(d_1044_v1_)
            nw170_ = _dafny.Array(None, 26)
            for i0_30_ in range(nw170_.length(0)):
                nw170_[i0_30_] = init30_(i0_30_)
            d_1072_v18_ = nw170_
            d_1075_v19_: _dafny.MultiSet
            d_1075_v19_ = _dafny.MultiSet([d_1047_v2_])
            d_1076_v20_: C0
            nw171_ = C0()
            nw171_.ctor__((d_1063_v13_) + (d_1063_v13_))
            d_1076_v20_ = nw171_
            d_1077_v21_: D3
            d_1077_v21_ = D3_DC7(d_1047_v2_)
            pat_let_tv11_ = d_1047_v2_
            d_1078_v22_: _dafny.Array
            nw172_ = _dafny.Array(None, 20)
            nw172_[int(0)] = d_1047_v2_
            nw172_[int(1)] = d_1047_v2_
            nw172_[int(2)] = (d_1047_v2_) + (d_1047_v2_)
            nw172_[int(3)] = (_dafny.SeqWithoutIsStrInference([self.f24])) + (d_1047_v2_)
            nw172_[int(4)] = default__.fm49(True, self.f24, d_1043_v0_, (self).f25, globalState)
            nw172_[int(5)] = (d_1047_v2_).set(default__.safeIndex(d_1043_v0_, len(d_1047_v2_)), (d_1069_v16_)[default__.safeIndex(149, (d_1069_v16_).length(0))])
            def iife72_(_pat_let16_0):
                def iife73_(d_1079_dt__update__tmp_h0_):
                    def iife74_(_pat_let17_0):
                        def iife75_(d_1080_dt__update_hcf15_h0_):
                            return D3_DC7(d_1080_dt__update_hcf15_h0_)
                        return iife75_(_pat_let17_0)
                    return iife74_(pat_let_tv11_)
                return iife73_(_pat_let16_0)
            nw172_[int(6)] = (iife72_(d_1077_v21_)).cf15
            nw172_[int(7)] = d_1047_v2_
            nw172_[int(8)] = d_1047_v2_
            nw172_[int(9)] = d_1047_v2_
            nw172_[int(10)] = default__.fm49((d_1069_v16_)[default__.safeIndex(149, (d_1069_v16_).length(0))], self.f24, d_1043_v0_, self.f24, globalState)
            nw172_[int(11)] = _dafny.SeqWithoutIsStrInference([not((self).f25), (self).f25])
            nw172_[int(12)] = (d_1047_v2_ if (d_1069_v16_)[default__.safeIndex(149, (d_1069_v16_).length(0))] else d_1047_v2_)
            nw172_[int(13)] = d_1047_v2_
            nw172_[int(14)] = _dafny.SeqWithoutIsStrInference([(d_1048_v3_).cf0, self.f24, self.f24])
            nw172_[int(15)] = d_1047_v2_
            nw172_[int(16)] = d_1047_v2_
            nw172_[int(17)] = _dafny.SeqWithoutIsStrInference([(self).f25, (self).f25])
            nw172_[int(18)] = d_1047_v2_
            nw172_[int(19)] = d_1047_v2_
            d_1078_v22_ = nw172_
            index160_ = default__.safeIndex(51, (d_1078_v22_).length(0))
            (d_1078_v22_)[index160_] = d_1047_v2_
            index161_ = default__.safeIndex(51, (d_1078_v22_).length(0))
            rhs137_ = d_1072_v18_
            rhs138_ = _dafny.MultiSet([d_1047_v2_])
            rhs139_ = d_1076_v20_
            rhs140_ = _dafny.SeqWithoutIsStrInference([(d_1069_v16_)[default__.safeIndex(149, (d_1069_v16_).length(0))], (d_1047_v2_)[default__.safeIndex(len(d_1063_v13_), len(d_1047_v2_))]])
            lhs125_ = d_1078_v22_
            lhs126_ = default__.safeIndex(51, (d_1078_v22_).length(0))
            d_1072_v18_ = rhs137_
            d_1075_v19_ = rhs138_
            d_1076_v20_ = rhs139_
            lhs125_[lhs126_] = rhs140_
        elif True:
            d_1081_v23_: _dafny.Seq
            d_1081_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kgqjtvxw"))
            d_1082_v24_: _dafny.Map
            d_1082_v24_ = _dafny.Map({len(d_1081_v23_): self.f24})
            d_1083_v25_: _dafny.Set
            d_1083_v25_ = _dafny.Set({d_1082_v24_, d_1082_v24_})
            rhs141_ = default__.safeDivisionInt(d_1043_v0_, (d_1043_v0_) + (d_1043_v0_))
            rhs142_ = not ((_dafny.Set({d_1082_v24_})).issubset(d_1083_v25_)) or (self.f24)
            lhs127_ = globalState
            lhs128_ = self
            lhs127_.f11 = rhs141_
            lhs128_.f24 = rhs142_
            d_1084_v26_: C1
            nw173_ = C1()
            nw173_.ctor__(self.f24)
            d_1084_v26_ = nw173_
            d_1085_v27_: _dafny.Array
            nw174_ = _dafny.Array(False, 15)
            d_1085_v27_ = nw174_
            d_1086_v28_: _dafny.Array
            nw175_ = _dafny.Array(None, 12)
            nw175_[int(0)] = d_1085_v27_
            nw175_[int(1)] = d_1085_v27_
            nw175_[int(2)] = d_1085_v27_
            nw175_[int(3)] = d_1085_v27_
            nw175_[int(4)] = d_1085_v27_
            nw175_[int(5)] = d_1085_v27_
            nw175_[int(6)] = d_1085_v27_
            nw175_[int(7)] = d_1085_v27_
            nw175_[int(8)] = d_1085_v27_
            nw175_[int(9)] = d_1085_v27_
            nw175_[int(10)] = d_1085_v27_
            nw175_[int(11)] = (d_1085_v27_ if True else d_1085_v27_)
            d_1086_v28_ = nw175_
            index162_ = default__.safeIndex(212, (d_1086_v28_).length(0))
            (d_1086_v28_)[index162_] = d_1085_v27_
            index163_ = default__.safeIndex(212, (d_1086_v28_).length(0))
            rhs143_ = d_1085_v27_
            rhs144_ = d_1043_v0_
            rhs145_ = default__.fm1(self.f24, (self).f25, len(d_1047_v2_), globalState)
            rhs146_ = ((d_1043_v0_) - (d_1043_v0_)) * ((0) - (default__.safeDivisionInt(d_1043_v0_, d_1043_v0_)))
            lhs129_ = d_1086_v28_
            lhs130_ = default__.safeIndex(212, (d_1086_v28_).length(0))
            lhs131_ = globalState
            lhs132_ = globalState
            lhs129_[lhs130_] = rhs143_
            d_1043_v0_ = rhs144_
            lhs131_.f10 = rhs145_
            lhs132_.f22 = rhs146_
            d_1087_v29_: _dafny.MultiSet
            d_1087_v29_ = _dafny.MultiSet([d_1043_v0_, d_1043_v0_])
            d_1088_v30_: _dafny.Map
            d_1088_v30_ = _dafny.Map({d_1043_v0_: d_1043_v0_})
            d_1089_v31_: _dafny.Set
            d_1089_v31_ = _dafny.Set({d_1088_v30_, d_1088_v30_})
            d_1082_v24_ = (d_1082_v24_).set(((default__.fm52((self).f25, globalState)) | (d_1087_v29_)).cardinality, (len(d_1089_v31_)) >= (d_1043_v0_))
            (globalState).f11 = d_1043_v0_
        d_1090_v32_: _dafny.Map
        d_1090_v32_ = _dafny.Map({(self).f25: d_1043_v0_})
        d_1091_v33_: _dafny.Map
        d_1091_v33_ = _dafny.Map({d_1090_v32_: True})
        (globalState).f16 = ((len(d_1091_v33_)) + (d_1043_v0_)) + (d_1043_v0_)
        if (self).f25:
            d_1092_v34_: _dafny.Set
            d_1092_v34_ = _dafny.Set({d_1044_v1_})
            d_1093_v35_: D2
            d_1093_v35_ = D2_DC4((799) != ((0) - (len(d_1092_v34_))))
            source20_ = d_1093_v35_
            if source20_.is_DC4:
                d_1094___mcc_h0_ = source20_.cf9
                d_1095_cf9_ = d_1094___mcc_h0_
                (globalState).f22 = d_1043_v0_
                d_1096_v36_: _dafny.Array
                def lambda53_(d_1097_v2_):
                    def lambda54_(d_1098_i9_):
                        return d_1097_v2_

                    return lambda54_

                init31_ = lambda53_(d_1047_v2_)
                nw176_ = _dafny.Array(None, 22)
                for i0_31_ in range(nw176_.length(0)):
                    nw176_[i0_31_] = init31_(i0_31_)
                d_1096_v36_ = nw176_
                d_1096_v36_ = d_1096_v36_
                rhs147_ = (0) - (default__.safeDivisionInt((self).fm5(_dafny.Map({d_1095_cf9_: d_1049_v4_}), self.f24, d_1095_cf9_, globalState), d_1043_v0_))
                rhs148_ = (d_1043_v0_) + ((0) - ((819 if (self).f25 else (0) - (d_1043_v0_))))
                lhs133_ = globalState
                lhs134_ = globalState
                lhs133_.f16 = rhs147_
                lhs134_.f11 = rhs148_
                (globalState).f16 = (0) - (d_1043_v0_)
            elif source20_.is_DC5:
                d_1099___mcc_h1_ = source20_.cf10
                d_1100___mcc_h2_ = source20_.cf11
                d_1101___mcc_h3_ = source20_.cf12
                d_1102___mcc_h4_ = source20_.cf13
                d_1103_cf13_ = d_1102___mcc_h4_
                d_1104_cf12_ = d_1101___mcc_h3_
                d_1105_cf11_ = d_1100___mcc_h2_
                d_1106_cf10_ = d_1099___mcc_h1_
                d_1107_v37_: _dafny.Array
                def lambda55_(d_1108_cf13_):
                    def lambda56_(d_1109_i10_):
                        return (d_1109_i10_) * ((0) - (d_1108_cf13_))

                    return lambda56_

                init32_ = lambda55_(d_1103_cf13_)
                nw177_ = _dafny.Array(None, 29)
                for i0_32_ in range(nw177_.length(0)):
                    nw177_[i0_32_] = init32_(i0_32_)
                d_1107_v37_ = nw177_
                index164_ = default__.safeIndex(100, (d_1107_v37_).length(0))
                (d_1107_v37_)[index164_] = (d_1104_cf12_) + (len(_dafny.Map({self.f24: d_1043_v0_})))
                d_1110_v38_: _dafny.Array
                nw178_ = _dafny.Array(_dafny.CodePoint('D'), 29)
                d_1110_v38_ = nw178_
                d_1111_v39_: _dafny.Map
                d_1111_v39_ = _dafny.Map({d_1044_v1_: self.f24})
                d_1112_v40_: _dafny.Seq
                d_1112_v40_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ok"))
                index165_ = default__.safeIndex(100, (d_1107_v37_).length(0))
                rhs149_ = len(d_1111_v39_)
                rhs150_ = default__.safeModuloInt(d_1104_cf12_, default__.safeModuloInt(len(d_1112_v40_), d_1103_cf13_))
                rhs151_ = d_1103_cf13_
                rhs152_ = (self).f25
                rhs153_ = d_1110_v38_
                lhs135_ = d_1107_v37_
                lhs136_ = default__.safeIndex(100, (d_1107_v37_).length(0))
                lhs137_ = globalState
                lhs138_ = globalState
                lhs135_[lhs136_] = rhs149_
                lhs137_.f11 = rhs150_
                d_1043_v0_ = rhs151_
                lhs138_.f10 = rhs152_
                d_1110_v38_ = rhs153_
                d_1113_v41_: C2
                nw179_ = C2()
                nw179_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ufhiybsa")), d_1112_v40_)
                d_1113_v41_ = nw179_
                d_1114_v42_: _dafny.Array
                nw180_ = _dafny.Array(False, 7)
                d_1114_v42_ = nw180_
                d_1115_v43_: _dafny.Array
                nw181_ = _dafny.Array(None, 6)
                nw181_[int(0)] = d_1114_v42_
                nw181_[int(1)] = d_1114_v42_
                nw181_[int(2)] = d_1114_v42_
                nw181_[int(3)] = d_1114_v42_
                nw181_[int(4)] = d_1114_v42_
                nw181_[int(5)] = d_1114_v42_
                d_1115_v43_ = nw181_
                d_1116_v44_: D7
                d_1116_v44_ = D7_DC15(d_1106_cf10_, default__.fm1(d_1105_cf11_, self.f24, 425, globalState))
                d_1117_v45_: _dafny.Seq
                d_1117_v45_ = _dafny.SeqWithoutIsStrInference([d_1116_v44_, d_1116_v44_])
                d_1118_v46_: _dafny.Map
                d_1118_v46_ = _dafny.Map({d_1043_v0_: self.f24})
                d_1119_v47_: D6
                d_1119_v47_ = D6_DC12(_dafny.SeqWithoutIsStrInference([len(d_1118_v46_), d_1103_cf13_]))
                d_1120_v48_: _dafny.MultiSet
                d_1120_v48_ = _dafny.MultiSet([(d_1117_v45_) + (default__.fm54(d_1119_v47_, globalState)), _dafny.SeqWithoutIsStrInference([d_1116_v44_]), (default__.fm54(d_1119_v47_, globalState)) + (d_1117_v45_)])
                rhs154_ = d_1115_v43_
                rhs155_ = (self).fm5(default__.fm55(d_1106_cf10_, d_1105_cf11_, (d_1113_v41_).fm19(d_1106_cf10_, globalState), self.f24, globalState), d_1105_cf11_, True, globalState)
                rhs156_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_1116_v44_ for d_1121_i11_ in range(default__.abs(418))]), d_1117_v45_, (d_1117_v45_) + (d_1117_v45_), (d_1117_v45_).set(default__.safeIndex(len(d_1047_v2_), len(d_1117_v45_)), D7_DC15(d_1106_cf10_, (self).f25))])
                lhs139_ = globalState
                d_1115_v43_ = rhs154_
                lhs139_.f11 = rhs155_
                d_1120_v48_ = rhs156_
                (globalState).f10 = d_1105_cf11_
            elif source20_.is_DC3:
                d_1122___mcc_h5_ = source20_.cf8
                d_1123_cf8_ = d_1122___mcc_h5_
                d_1124_v49_: D2
                d_1124_v49_ = D2_DC4(False)
                d_1125_v50_: D2
                d_1125_v50_ = D2_DC6(d_1124_v49_)
                d_1126_v51_: D5
                d_1126_v51_ = D5_DC11(d_1125_v50_, (self).f25)
                d_1127_v52_: _dafny.Array
                def lambda57_(d_1128_i12_):
                    return (self).f25

                init33_ = lambda57_
                nw182_ = _dafny.Array(None, 20)
                for i0_33_ in range(nw182_.length(0)):
                    nw182_[i0_33_] = init33_(i0_33_)
                d_1127_v52_ = nw182_
                d_1129_v53_: _dafny.Map
                d_1129_v53_ = _dafny.Map({d_1126_v51_: d_1127_v52_})
                (globalState).f2 = ((d_1129_v53_)[d_1126_v51_] if (d_1126_v51_) in (d_1129_v53_) else d_1127_v52_)
                d_1130_v54_: _dafny.Seq
                d_1130_v54_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tpcn"))])
                d_1131_v55_: _dafny.Seq
                d_1131_v55_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nkekndhk"))
                d_1132_v56_: C2
                nw183_ = C2()
                nw183_.ctor__(d_1131_v55_, d_1131_v55_)
                d_1132_v56_ = nw183_
                d_1133_v57_: _dafny.Map
                d_1133_v57_ = _dafny.Map({d_1043_v0_: d_1132_v56_})
                d_1134_v58_: _dafny.Set
                d_1134_v58_ = _dafny.Set({_dafny.CodePoint('m')})
                d_1135_v59_: _dafny.Array
                nw184_ = _dafny.Array(None, 21)
                nw184_[int(0)] = d_1043_v0_
                nw184_[int(1)] = d_1043_v0_
                nw184_[int(2)] = len(d_1133_v57_)
                nw184_[int(3)] = d_1043_v0_
                nw184_[int(4)] = d_1043_v0_
                nw184_[int(5)] = d_1043_v0_
                nw184_[int(6)] = d_1043_v0_
                nw184_[int(7)] = d_1043_v0_
                nw184_[int(8)] = d_1043_v0_
                nw184_[int(9)] = 654
                nw184_[int(10)] = d_1043_v0_
                nw184_[int(11)] = -299
                nw184_[int(12)] = d_1043_v0_
                nw184_[int(13)] = ((d_1090_v32_)[self.f24] if (self.f24) in (d_1090_v32_) else len((d_1132_v56_).f30))
                nw184_[int(14)] = d_1043_v0_
                nw184_[int(15)] = (0) - (len(p0))
                nw184_[int(16)] = d_1043_v0_
                nw184_[int(17)] = len(d_1134_v58_)
                nw184_[int(18)] = 660
                nw184_[int(19)] = d_1043_v0_
                nw184_[int(20)] = d_1043_v0_
                d_1135_v59_ = nw184_
                d_1136_v60_: _dafny.Map
                d_1136_v60_ = _dafny.Map({(self).f25: (self).f25})
                d_1137_v61_: _dafny.MultiSet
                d_1137_v61_ = _dafny.MultiSet([d_1043_v0_, 857])
                d_1138_v62_: _dafny.Seq
                d_1138_v62_ = _dafny.SeqWithoutIsStrInference([d_1127_v52_])
                d_1139_v63_: _dafny.MultiSet
                d_1139_v63_ = _dafny.MultiSet([(self).f25])
                d_1140_v64_: D6
                d_1140_v64_ = D6_DC13(d_1127_v52_, (d_1139_v63_).cardinality, (self).f25)
                d_1141_v65_: _dafny.Map
                d_1141_v65_ = _dafny.Map({d_1043_v0_: (d_1140_v64_).cf26})
                d_1142_v66_: _dafny.Array
                nw185_ = _dafny.Array(None, 29)
                nw185_[int(0)] = 348
                nw185_[int(1)] = d_1043_v0_
                nw185_[int(2)] = (_dafny.MultiSet([d_1135_v59_, d_1135_v59_])).cardinality
                nw185_[int(3)] = d_1043_v0_
                nw185_[int(4)] = len((d_1132_v56_).f30)
                nw185_[int(5)] = d_1043_v0_
                nw185_[int(6)] = (0) - (len(d_1044_v1_))
                nw185_[int(7)] = d_1043_v0_
                nw185_[int(8)] = d_1043_v0_
                nw185_[int(9)] = len(d_1136_v60_)
                nw185_[int(10)] = (d_1137_v61_).cardinality
                nw185_[int(11)] = len(d_1138_v62_)
                nw185_[int(12)] = len(d_1047_v2_)
                nw185_[int(13)] = d_1043_v0_
                nw185_[int(14)] = -239
                nw185_[int(15)] = d_1043_v0_
                nw185_[int(16)] = d_1043_v0_
                nw185_[int(17)] = 778
                nw185_[int(18)] = len((d_1132_v56_).f30)
                nw185_[int(19)] = d_1043_v0_
                nw185_[int(20)] = len(d_1141_v65_)
                nw185_[int(21)] = d_1043_v0_
                nw185_[int(22)] = len(d_1131_v55_)
                nw185_[int(23)] = d_1043_v0_
                nw185_[int(24)] = d_1043_v0_
                nw185_[int(25)] = d_1043_v0_
                nw185_[int(26)] = (d_1139_v63_).cardinality
                nw185_[int(27)] = d_1043_v0_
                nw185_[int(28)] = len(d_1044_v1_)
                d_1142_v66_ = nw185_
                d_1143_v67_: _dafny.Map
                d_1143_v67_ = _dafny.Map({d_1043_v0_: d_1135_v59_})
                d_1144_v68_: _dafny.Set
                d_1144_v68_ = _dafny.Set({False})
                d_1145_v69_: _dafny.Array
                nw186_ = _dafny.Array(None, 17)
                nw186_[int(0)] = d_1043_v0_
                nw186_[int(1)] = len(_dafny.SeqWithoutIsStrInference([d_1043_v0_, 160, (0) - (d_1043_v0_)]))
                nw186_[int(2)] = d_1043_v0_
                nw186_[int(3)] = d_1043_v0_
                nw186_[int(4)] = d_1043_v0_
                nw186_[int(5)] = d_1043_v0_
                nw186_[int(6)] = 182
                nw186_[int(7)] = len(d_1143_v67_)
                nw186_[int(8)] = d_1043_v0_
                nw186_[int(9)] = -549
                nw186_[int(10)] = len(_dafny.Map({d_1043_v0_: len(d_1047_v2_)}))
                nw186_[int(11)] = d_1043_v0_
                nw186_[int(12)] = ((d_1123_cf8_)[len(d_1144_v68_)] if (len(d_1144_v68_)) in (d_1123_cf8_) else (self).fm6(d_1048_v3_, globalState))
                nw186_[int(13)] = len(_dafny.Map({_dafny.Map({d_1043_v0_: d_1043_v0_}): self.f24}))
                nw186_[int(14)] = d_1043_v0_
                nw186_[int(15)] = (_dafny.MultiSet([d_1090_v32_, d_1090_v32_, d_1090_v32_, d_1090_v32_, d_1090_v32_])).cardinality
                nw186_[int(16)] = d_1043_v0_
                d_1145_v69_ = nw186_
                index166_ = default__.safeIndex(450, (d_1145_v69_).length(0))
                (d_1145_v69_)[index166_] = d_1043_v0_
                d_1146_v70_: _dafny.Map
                d_1146_v70_ = _dafny.Map({not(False): d_1047_v2_})
                index167_ = default__.safeIndex(450, (d_1145_v69_).length(0))
                rhs157_ = (d_1130_v54_).set(default__.safeIndex(d_1043_v0_, len(d_1130_v54_)), d_1132_v56_.f29)
                rhs158_ = (d_1127_v52_ if self.f24 else d_1127_v52_)
                rhs159_ = (self).f25
                rhs160_ = len(((d_1146_v70_)[(d_1144_v68_).isdisjoint(d_1144_v68_)] if ((d_1144_v68_).isdisjoint(d_1144_v68_)) in (d_1146_v70_) else d_1047_v2_))
                rhs161_ = 789
                lhs140_ = globalState
                lhs141_ = globalState
                lhs142_ = d_1145_v69_
                lhs143_ = default__.safeIndex(450, (d_1145_v69_).length(0))
                d_1130_v54_ = rhs157_
                d_1127_v52_ = rhs158_
                lhs140_.f10 = rhs159_
                lhs141_.f22 = rhs160_
                lhs142_[lhs143_] = rhs161_
                (self).f24 = (d_1138_v62_) < (d_1138_v62_)
                d_1147_v71_: str
                d_1147_v71_ = _dafny.CodePoint('i')
                d_1148_v72_: C3
                nw187_ = C3()
                nw187_.ctor__((self).f25, d_1131_v55_, (d_1043_v0_) != (d_1043_v0_), d_1147_v71_, (self).f25, (self).f25)
                d_1148_v72_ = nw187_
            elif True:
                d_1149___mcc_h6_ = source20_.cf14
                d_1150_cf14_ = d_1149___mcc_h6_
                d_1151_v73_: str
                d_1151_v73_ = _dafny.CodePoint('j')
                d_1152_v74_: T3
                nw188_ = C3()
                nw188_.ctor__(self.f24, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_1153_i13_ in range(default__.abs(188))]), False, d_1151_v73_, True, (self).f25)
                d_1152_v74_ = nw188_
                d_1154_v75_: _dafny.Array
                nw189_ = _dafny.Array(None, 8)
                nw189_[int(0)] = d_1152_v74_
                nw189_[int(1)] = d_1152_v74_
                nw189_[int(2)] = d_1152_v74_
                nw189_[int(3)] = d_1152_v74_
                nw189_[int(4)] = d_1152_v74_
                nw189_[int(5)] = d_1152_v74_
                nw189_[int(6)] = d_1152_v74_
                nw189_[int(7)] = d_1152_v74_
                d_1154_v75_ = nw189_
                d_1155_v76_: _dafny.Map
                d_1155_v76_ = _dafny.Map({d_1154_v75_: (0) - ((0) - (d_1043_v0_))})
                d_1155_v76_ = ((d_1155_v76_) | (d_1155_v76_) if not((d_1152_v74_).f26) else d_1155_v76_)
                (self).f24 = not(True)
                (globalState).f22 = ((d_1043_v0_ if d_1152_v74_.f24 else 194)) - (d_1043_v0_)
                d_1156_v77_: _dafny.Array
                nw190_ = _dafny.Array(_dafny.Array(None, 0), 20)
                d_1156_v77_ = nw190_
                d_1157_v78_: _dafny.Array
                nw191_ = _dafny.Array(int(0), 24)
                d_1157_v78_ = nw191_
                index168_ = default__.safeIndex(322, (d_1156_v77_).length(0))
                (d_1156_v77_)[index168_] = d_1157_v78_
                index169_ = default__.safeIndex(322, (d_1156_v77_).length(0))
                (d_1156_v77_)[index169_] = d_1157_v78_
            (self).f24 = default__.fm1(not(self.f24), self.f24, 786, globalState)
            d_1158_v79_: _dafny.Map
            d_1158_v79_ = _dafny.Map({975: d_1043_v0_})
            d_1159_v80_: _dafny.Seq
            d_1159_v80_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iec"))
            d_1160_v81_: D1
            d_1160_v81_ = D1_DC1(d_1159_v80_)
            d_1161_v82_: _dafny.Map
            d_1161_v82_ = _dafny.Map({d_1160_v81_: d_1043_v0_})
            d_1162_v84_: _dafny.MultiSet
            d_1162_v84_ = _dafny.MultiSet([(self).f25, self.f24])
            def iife76_():
                coll40_ = _dafny.Map()
                compr_40_: int
                for compr_40_ in (p0).Elements:
                    d_1163_v83_: int = compr_40_
                    if (d_1163_v83_) in (p0):
                        coll40_[(d_1163_v83_) * (d_1043_v0_)] = (self).f25
                return _dafny.Map(coll40_)
            rhs162_ = (((d_1158_v79_)[len(d_1159_v80_)] if (len(d_1159_v80_)) in (d_1158_v79_) else len((d_1161_v82_).set(D1_DC1(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cihbog"))), d_1043_v0_)))) == (d_1043_v0_)
            rhs163_ = (d_1043_v0_) > (default__.fm0(d_1043_v0_, (self).f25, iife76_()
            , d_1043_v0_, globalState))
            rhs164_ = len((d_1090_v32_) | ((d_1090_v32_).set((self).f25, ((d_1162_v84_)[not((self).f25)] if (not((self).f25)) in (d_1162_v84_) else d_1043_v0_))))
            lhs144_ = globalState
            lhs145_ = globalState
            lhs146_ = globalState
            lhs144_.f10 = rhs162_
            lhs145_.f10 = rhs163_
            lhs146_.f22 = rhs164_
            d_1164_v85_: _dafny.Map
            d_1164_v85_ = _dafny.Map({d_1043_v0_: self.f24})
            d_1164_v85_ = (d_1164_v85_).set(len(d_1158_v79_), (d_1047_v2_) <= (d_1047_v2_))
            d_1165_v86_: _dafny.Array
            nw192_ = _dafny.Array(int(0), 10)
            d_1165_v86_ = nw192_
            (globalState).f7 = d_1165_v86_
        elif True:
            d_1166_v87_: _dafny.Seq
            d_1166_v87_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sjggemedc"))
            d_1167_v88_: _dafny.Seq
            d_1167_v88_ = _dafny.SeqWithoutIsStrInference([d_1166_v87_, d_1166_v87_])
            d_1168_v89_: str
            d_1168_v89_ = _dafny.CodePoint('m')
            d_1169_v90_: _dafny.Map
            d_1169_v90_ = _dafny.Map({d_1043_v0_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k"))})
            d_1170_v91_: _dafny.Array
            nw193_ = _dafny.Array(None, 12)
            nw193_[int(0)] = (((d_1167_v88_)[default__.safeIndex(d_1043_v0_, len(d_1167_v88_))]) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "psi")))).set(default__.safeIndex(d_1043_v0_, len(((d_1167_v88_)[default__.safeIndex(d_1043_v0_, len(d_1167_v88_))]) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "psi"))))), d_1168_v89_)
            nw193_[int(1)] = d_1166_v87_
            nw193_[int(2)] = d_1166_v87_
            nw193_[int(3)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ormdo"))).set(default__.safeIndex(len(d_1047_v2_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ormdo")))), d_1168_v89_)
            nw193_[int(4)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "drm"))
            nw193_[int(5)] = d_1166_v87_
            nw193_[int(6)] = (d_1166_v87_) + (d_1166_v87_)
            nw193_[int(7)] = (d_1166_v87_) + (d_1166_v87_)
            nw193_[int(8)] = d_1166_v87_
            nw193_[int(9)] = ((d_1169_v90_)[d_1043_v0_] if (d_1043_v0_) in (d_1169_v90_) else d_1166_v87_)
            nw193_[int(10)] = d_1166_v87_
            nw193_[int(11)] = d_1166_v87_
            d_1170_v91_ = nw193_
            d_1170_v91_ = d_1170_v91_
            d_1171_v92_: _dafny.Array
            def lambda58_(d_1172_v89_, d_1173_v0_):
                def lambda59_(d_1174_i14_):
                    return ((D2_DC5(d_1172_v89_, False, d_1173_v0_, d_1173_v0_)).cf13) != (d_1173_v0_)

                return lambda59_

            init34_ = lambda58_(d_1168_v89_, d_1043_v0_)
            nw194_ = _dafny.Array(None, 9)
            for i0_34_ in range(nw194_.length(0)):
                nw194_[i0_34_] = init34_(i0_34_)
            d_1171_v92_ = nw194_
            index170_ = default__.safeIndex(582, (d_1171_v92_).length(0))
            (d_1171_v92_)[index170_] = (self).f25
            index171_ = default__.safeIndex(582, (d_1171_v92_).length(0))
            (d_1171_v92_)[index171_] = (self).f25
            (globalState).f10 = (self).f25
            d_1175_v93_: D3
            d_1175_v93_ = D3_DC8((self).f25, d_1043_v0_, d_1043_v0_)
            source21_ = d_1175_v93_
            if source21_.is_DC8:
                d_1176___mcc_h7_ = source21_.cf16
                d_1177___mcc_h8_ = source21_.cf17
                d_1178___mcc_h9_ = source21_.cf18
                d_1179_cf18_ = d_1178___mcc_h9_
                d_1180_cf17_ = d_1177___mcc_h8_
                d_1181_cf16_ = d_1176___mcc_h7_
                (globalState).f10 = (d_1047_v2_)[default__.safeIndex(d_1043_v0_, len(d_1047_v2_))]
                index172_ = default__.safeIndex(505, (d_1170_v91_).length(0))
                (d_1170_v91_)[index172_] = d_1166_v87_
                index173_ = default__.safeIndex(505, (d_1170_v91_).length(0))
                (d_1170_v91_)[index173_] = (d_1166_v87_) + (d_1166_v87_)
                d_1182_v94_: D8
                d_1182_v94_ = D8_DC19((self).f25, d_1180_cf17_)
                d_1183_v95_: _dafny.Map
                d_1183_v95_ = _dafny.Map({d_1182_v94_: self.f24})
                d_1183_v95_ = d_1183_v95_
                index174_ = default__.safeIndex(582, (d_1171_v92_).length(0))
                rhs165_ = d_1180_cf17_
                rhs166_ = (d_1171_v92_)[default__.safeIndex(582, (d_1171_v92_).length(0))]
                rhs167_ = d_1168_v89_
                lhs147_ = globalState
                lhs148_ = d_1171_v92_
                lhs149_ = default__.safeIndex(582, (d_1171_v92_).length(0))
                lhs147_.f11 = rhs165_
                lhs148_[lhs149_] = rhs166_
                r0 = rhs167_
            elif True:
                d_1184___mcc_h10_ = source21_.cf15
                d_1185_cf15_ = d_1184___mcc_h10_
                d_1166_v87_ = (d_1166_v87_) + (d_1166_v87_)
                d_1186_v96_: _dafny.Array
                nw195_ = _dafny.Array(_dafny.MultiSet({}), 10)
                d_1186_v96_ = nw195_
                d_1187_v97_: _dafny.MultiSet
                d_1187_v97_ = _dafny.MultiSet([self.f24])
                index175_ = default__.safeIndex(722, (d_1186_v96_).length(0))
                (d_1186_v96_)[index175_] = d_1187_v97_
                index176_ = default__.safeIndex(722, (d_1186_v96_).length(0))
                (d_1186_v96_)[index176_] = d_1187_v97_
                d_1188_v98_: _dafny.Set
                d_1188_v98_ = _dafny.Set({d_1166_v87_, d_1166_v87_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y"))})
                (globalState).f10 = (d_1188_v98_).issubset(d_1188_v98_)
                (globalState).f22 = (d_1043_v0_) + (d_1043_v0_)
            d_1189_v99_: _dafny.MultiSet
            d_1189_v99_ = _dafny.MultiSet([d_1043_v0_])
            d_1190_v100_: _dafny.MultiSet
            d_1190_v100_ = _dafny.MultiSet([(d_1043_v0_) != ((d_1189_v99_).cardinality)])
            d_1190_v100_ = ((d_1190_v100_) | (d_1190_v100_)).intersection(d_1190_v100_)
        d_1191_v101_: str
        d_1191_v101_ = _dafny.CodePoint('c')
        r0 = (_dafny.CodePoint('m') if (self).f25 else d_1191_v101_)
        r1 = (d_1047_v2_).set(default__.safeIndex(d_1043_v0_, len(d_1047_v2_)), (self).f25)
        return r0, r1


class C5(T3, T2, T1, T0):
    def  __init__(self):
        self._f24: bool = False
        self._f27: str = _dafny.CodePoint('D')
        self._f25: bool = False
        self._f26: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    @property
    def f24(self):
        return self._f24
    @f24.setter
    def f24(self, value):
        self._f24 = value
    @property
    def f27(self):
        return self._f27
    @f27.setter
    def f27(self, value):
        self._f27 = value
    @property
    def f25(self):
        return self._f25
    @property
    def f26(self):
        return self._f26
    def ctor__(self, f26, f27, f25, f24):
        (self)._f26 = f26
        (self).f27 = f27
        (self)._f25 = f25
        (self).f24 = f24

    def fm7(self, p0, globalState):
        return (362) == (default__.safeModuloInt((0) - (-693), 394))

    def fm6(self, p0, globalState):
        return default__.safeDivisionInt(913, (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gotglday")))) - (599))

    def fm4(self, p0, p1, globalState):
        return D0_DC0((730) > (819), 91, default__.safeDivisionInt(len(_dafny.Map({614: self.f24})), 691))

    def fm5(self, p0, p1, p2, globalState):
        return (default__.safeDivisionInt(-233, len(_dafny.Set({(self).f26})))) * (default__.safeModuloInt(len(_dafny.Map({_dafny.SeqWithoutIsStrInference([-402, len(_dafny.SeqWithoutIsStrInference([self.f27 for d_1192_i0_ in range(default__.abs(328))])), 912]): True})), 323))

    def m2(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: bool = False
        r2: int = int(0)
        r3: int = int(0)
        d_1193_v0_: _dafny.Array
        def lambda60_(d_1194_i0_):
            return self.f24

        init35_ = lambda60_
        nw196_ = _dafny.Array(None, 9)
        for i0_35_ in range(nw196_.length(0)):
            nw196_[i0_35_] = init35_(i0_35_)
        d_1193_v0_ = nw196_
        d_1195_v1_: _dafny.Array
        nw197_ = _dafny.Array(None, 13)
        nw197_[int(0)] = d_1193_v0_
        nw197_[int(1)] = d_1193_v0_
        nw197_[int(2)] = d_1193_v0_
        nw197_[int(3)] = d_1193_v0_
        nw197_[int(4)] = d_1193_v0_
        nw197_[int(5)] = d_1193_v0_
        nw197_[int(6)] = d_1193_v0_
        nw197_[int(7)] = d_1193_v0_
        nw197_[int(8)] = d_1193_v0_
        nw197_[int(9)] = d_1193_v0_
        nw197_[int(10)] = (d_1193_v0_ if (self).f26 else d_1193_v0_)
        nw197_[int(11)] = d_1193_v0_
        nw197_[int(12)] = d_1193_v0_
        d_1195_v1_ = nw197_
        index177_ = default__.safeIndex(838, (d_1195_v1_).length(0))
        (d_1195_v1_)[index177_] = d_1193_v0_
        d_1196_v2_: _dafny.Seq
        d_1196_v2_ = _dafny.SeqWithoutIsStrInference([p1])
        d_1197_v3_: _dafny.Map
        d_1197_v3_ = _dafny.Map({(self).f26: d_1196_v2_})
        d_1198_v4_: _dafny.Map
        d_1198_v4_ = _dafny.Map({p0: (self).fm5(d_1197_v3_, p2, False, globalState)})
        d_1199_v5_: _dafny.Seq
        d_1199_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rmsenfy"))
        index178_ = default__.safeIndex(838, (d_1195_v1_).length(0))
        rhs168_ = (((d_1198_v4_)[p0] if (p0) in (d_1198_v4_) else p0)) * (((self).fm6(p1, globalState)) - (len(d_1199_v5_)))
        rhs169_ = self.f27
        rhs170_ = not(p2)
        rhs171_ = d_1193_v0_
        rhs172_ = self.f27
        lhs150_ = self
        lhs151_ = d_1195_v1_
        lhs152_ = default__.safeIndex(838, (d_1195_v1_).length(0))
        lhs153_ = self
        r3 = rhs168_
        lhs150_.f27 = rhs169_
        r0 = rhs170_
        lhs151_[lhs152_] = rhs171_
        lhs153_.f27 = rhs172_
        d_1200_v6_: D3
        d_1200_v6_ = D3_DC8(p2, p0, p0)
        (self).f24 = (d_1200_v6_).cf16
        d_1201_v7_: C0
        nw198_ = C0()
        nw198_.ctor__(d_1199_v5_)
        d_1201_v7_ = nw198_
        d_1202_v8_: _dafny.Array
        def lambda61_(d_1203_p0_):
            def lambda62_(d_1204_i1_):
                return default__.safeModuloInt(d_1204_i1_, d_1203_p0_)

            return lambda62_

        init36_ = lambda61_(p0)
        nw199_ = _dafny.Array(None, 14)
        for i0_36_ in range(nw199_.length(0)):
            nw199_[i0_36_] = init36_(i0_36_)
        d_1202_v8_ = nw199_
        d_1205_v9_: _dafny.Seq
        d_1205_v9_ = _dafny.SeqWithoutIsStrInference([d_1202_v8_, d_1202_v8_])
        d_1206_v10_: _dafny.MultiSet
        d_1206_v10_ = _dafny.MultiSet([(d_1205_v9_)[default__.safeIndex(p0, len(d_1205_v9_))]])
        d_1207_v11_: _dafny.Map
        d_1207_v11_ = _dafny.Map({d_1206_v10_: (d_1200_v6_).cf17})
        (self).f27 = (d_1199_v5_)[default__.safeIndex(((d_1207_v11_)[d_1206_v10_] if (d_1206_v10_) in (d_1207_v11_) else p0), len(d_1199_v5_))]
        d_1208_v12_: _dafny.Seq
        d_1208_v12_ = _dafny.SeqWithoutIsStrInference([self.f24])
        d_1209_v13_: _dafny.MultiSet
        d_1209_v13_ = _dafny.MultiSet([p2])
        d_1210_i2_: int
        d_1210_i2_ = 0
        with _dafny.label("12"):
            while (d_1209_v13_).ispropersubset(_dafny.MultiSet(d_1208_v12_)):
                with _dafny.c_label("12"):
                    if (d_1210_i2_) >= (100):
                        raise _dafny.Break("12")
                    d_1210_i2_ = (d_1210_i2_) + (1)
                    index179_ = default__.safeIndex(733, (d_1202_v8_).length(0))
                    (d_1202_v8_)[index179_] = p0
                    arr0_ = (d_1195_v1_)[default__.safeIndex(838, (d_1195_v1_).length(0))]
                    index180_ = default__.safeIndex(804, ((d_1195_v1_)[default__.safeIndex(838, (d_1195_v1_).length(0))]).length(0))
                    arr0_[index180_] = p2
                    index181_ = default__.safeIndex(733, (d_1202_v8_).length(0))
                    arr1_ = (d_1195_v1_)[default__.safeIndex(838, (d_1195_v1_).length(0))]
                    index182_ = default__.safeIndex(804, ((d_1195_v1_)[default__.safeIndex(838, (d_1195_v1_).length(0))]).length(0))
                    rhs173_ = default__.fm0(p0, (self).f25, _dafny.Map({p0: (self).f26}), default__.safeDivisionInt(p0, -876), globalState)
                    rhs174_ = (0) - (p0)
                    rhs175_ = p0
                    rhs176_ = (d_1208_v12_)[default__.safeIndex(385, len(d_1208_v12_))]
                    lhs154_ = globalState
                    lhs155_ = d_1202_v8_
                    lhs156_ = default__.safeIndex(733, (d_1202_v8_).length(0))
                    lhs157_ = globalState
                    lhs158_ = (d_1195_v1_)[default__.safeIndex(838, (d_1195_v1_).length(0))]
                    lhs159_ = default__.safeIndex(804, ((d_1195_v1_)[default__.safeIndex(838, (d_1195_v1_).length(0))]).length(0))
                    lhs154_.f11 = rhs173_
                    lhs155_[lhs156_] = rhs174_
                    lhs157_.f16 = rhs175_
                    lhs158_[lhs159_] = rhs176_
                    d_1211_v14_: _dafny.Map
                    d_1211_v14_ = _dafny.Map({not(False): (d_1202_v8_)[default__.safeIndex(733, (d_1202_v8_).length(0))]})
                    r2 = ((d_1211_v14_)[not (True) or (True)] if (not (True) or (True)) in (d_1211_v14_) else (d_1202_v8_)[default__.safeIndex(733, (d_1202_v8_).length(0))])
                    d_1212_v15_: _dafny.Set
                    d_1212_v15_ = _dafny.Set({p0})
                    d_1213_v16_: _dafny.Map
                    d_1213_v16_ = _dafny.Map({(d_1212_v15_).issubset(d_1212_v15_): (self).f25})
                    d_1214_v17_: _dafny.MultiSet
                    d_1214_v17_ = _dafny.MultiSet([(d_1202_v8_)[default__.safeIndex(733, (d_1202_v8_).length(0))]])
                    d_1213_v16_ = (d_1213_v16_).set((d_1214_v17_).issubset(d_1214_v17_), (self.f24 if ((d_1195_v1_)[default__.safeIndex(838, (d_1195_v1_).length(0))])[default__.safeIndex(804, ((d_1195_v1_)[default__.safeIndex(838, (d_1195_v1_).length(0))]).length(0))] else (self).f25))
                    d_1199_v5_ = (default__.fm29(globalState)) + (d_1199_v5_)
                    pass
            pass
        d_1215_v18_: _dafny.Array
        d_1215_v18_ = d_1195_v1_
        d_1216_v19_: D0
        d_1216_v19_ = D0_DC0(False, 356, p0)
        d_1217_v20_: _dafny.Array
        nw200_ = _dafny.Array(_dafny.Seq({}), 3)
        d_1217_v20_ = nw200_
        d_1218_v21_: _dafny.Map
        d_1218_v21_ = _dafny.Map({d_1202_v8_: (_dafny.SeqWithoutIsStrInference([p0 for d_1219_i3_ in range(default__.abs(682))])).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference([p0 for d_1220_i3_ in range(default__.abs(682))]))), p0)})
        d_1221_v22_: _dafny.Seq
        d_1221_v22_ = _dafny.SeqWithoutIsStrInference([p0, p0])
        index183_ = default__.safeIndex(366, (d_1217_v20_).length(0))
        (d_1217_v20_)[index183_] = ((d_1218_v21_)[d_1202_v8_] if (d_1202_v8_) in (d_1218_v21_) else d_1221_v22_)
        index184_ = default__.safeIndex(366, (d_1217_v20_).length(0))
        rhs177_ = d_1215_v18_
        rhs178_ = (self).fm5(d_1197_v3_, (self).f26, (self).f25, globalState)
        rhs179_ = p0
        rhs180_ = p1
        rhs181_ = ((d_1221_v22_) + (_dafny.SeqWithoutIsStrInference([p0 for d_1222_i4_ in range(default__.abs(-171))]))) + (_dafny.SeqWithoutIsStrInference([p0, p0, p0]))
        lhs160_ = globalState
        lhs161_ = globalState
        lhs162_ = d_1217_v20_
        lhs163_ = default__.safeIndex(366, (d_1217_v20_).length(0))
        d_1215_v18_ = rhs177_
        lhs160_.f11 = rhs178_
        lhs161_.f16 = rhs179_
        d_1216_v19_ = rhs180_
        lhs162_[lhs163_] = rhs181_
        r0 = (self).f25
        r1 = p2
        r2 = p0
        r3 = p0
        return r0, r1, r2, r3

    def m3(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: bool = False
        r2: _dafny.Map = _dafny.Map({})
        d_1223_i0_: int
        d_1223_i0_ = 0
        with _dafny.label("13"):
            while p1:
                with _dafny.c_label("13"):
                    if (d_1223_i0_) >= (100):
                        raise _dafny.Break("13")
                    d_1223_i0_ = (d_1223_i0_) + (1)
                    d_1224_v0_: D0
                    d_1224_v0_ = D0_DC0(True, (0) - (p0), p0)
                    d_1225_v1_: _dafny.Map
                    d_1225_v1_ = _dafny.Map({(self).fm6(d_1224_v0_, globalState): (self).f25})
                    d_1226_v2_: _dafny.Seq
                    d_1226_v2_ = _dafny.SeqWithoutIsStrInference([(self).f25])
                    d_1227_v3_: _dafny.Map
                    d_1227_v3_ = _dafny.Map({p1: len(d_1226_v2_)})
                    rhs182_ = (self).fm7(d_1227_v3_, globalState)
                    rhs183_ = d_1225_v1_
                    lhs164_ = globalState
                    lhs164_.f10 = rhs182_
                    d_1225_v1_ = rhs183_
                    d_1228_v4_: _dafny.Set
                    d_1228_v4_ = _dafny.Set({_dafny.CodePoint('d'), self.f27})
                    d_1229_v5_: _dafny.Seq
                    d_1229_v5_ = _dafny.SeqWithoutIsStrInference([len(d_1228_v4_), p0, p0])
                    d_1230_v6_: C1
                    nw201_ = C1()
                    nw201_.ctor__(((d_1229_v5_)[default__.safeIndex(p0, len(d_1229_v5_))]) == (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s")))))
                    d_1230_v6_ = nw201_
                    d_1231_v7_: _dafny.MultiSet
                    d_1231_v7_ = _dafny.MultiSet([d_1225_v1_, _dafny.Map({p0: (self).f26})])
                    d_1232_v8_: _dafny.Map
                    d_1232_v8_ = _dafny.Map({not((self).f25): not((self).f26)})
                    r0 = (default__.fm30(662, ((d_1232_v8_)[(self).f26] if ((self).f26) in (d_1232_v8_) else (self).f26), 749, p0, globalState)).ispropersubset(d_1231_v7_)
                    index185_ = default__.safeIndex(132, (p2).length(0))
                    (p2)[index185_] = p1
                    d_1233_v9_: _dafny.Seq
                    d_1233_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hfdsbjpl"))
                    d_1234_v10_: _dafny.Map
                    d_1234_v10_ = _dafny.Map({(self).f25: (d_1233_v9_)[default__.safeIndex(p0, len(d_1233_v9_))]})
                    d_1235_v11_: _dafny.Map
                    d_1235_v11_ = _dafny.Map({self.f27: (self).f26})
                    d_1236_v12_: _dafny.Array
                    nw202_ = _dafny.Array(None, 23)
                    nw202_[int(0)] = d_1234_v10_
                    nw202_[int(1)] = _dafny.Map({True: (d_1233_v9_)[default__.safeIndex(len(d_1233_v9_), len(d_1233_v9_))]})
                    nw202_[int(2)] = d_1234_v10_
                    nw202_[int(3)] = d_1234_v10_
                    nw202_[int(4)] = (d_1234_v10_) | (d_1234_v10_)
                    nw202_[int(5)] = d_1234_v10_
                    nw202_[int(6)] = d_1234_v10_
                    nw202_[int(7)] = (_dafny.Map({(self).f25: self.f27})) | (d_1234_v10_)
                    nw202_[int(8)] = d_1234_v10_
                    nw202_[int(9)] = default__.fm31(globalState)
                    nw202_[int(10)] = default__.fm31(globalState)
                    nw202_[int(11)] = d_1234_v10_
                    nw202_[int(12)] = d_1234_v10_
                    nw202_[int(13)] = default__.fm31(globalState)
                    nw202_[int(14)] = (_dafny.Map({(self).f26: self.f27})) | (d_1234_v10_)
                    nw202_[int(15)] = (d_1234_v10_).set(self.f24, self.f27)
                    nw202_[int(16)] = d_1234_v10_
                    nw202_[int(17)] = d_1234_v10_
                    nw202_[int(18)] = (d_1234_v10_ if ((d_1235_v11_)[self.f27] if (self.f27) in (d_1235_v11_) else self.f24) else d_1234_v10_)
                    nw202_[int(19)] = (d_1234_v10_) | (d_1234_v10_)
                    nw202_[int(20)] = d_1234_v10_
                    nw202_[int(21)] = d_1234_v10_
                    nw202_[int(22)] = d_1234_v10_
                    d_1236_v12_ = nw202_
                    d_1237_v13_: _dafny.Set
                    d_1237_v13_ = _dafny.Set({-406})
                    index186_ = default__.safeIndex(132, (p2).length(0))
                    rhs184_ = p0
                    rhs185_ = p2
                    rhs186_ = p0
                    rhs187_ = (-902) > (len((d_1237_v13_).intersection(d_1237_v13_)))
                    rhs188_ = d_1236_v12_
                    lhs165_ = globalState
                    lhs166_ = globalState
                    lhs167_ = globalState
                    lhs168_ = p2
                    lhs169_ = default__.safeIndex(132, (p2).length(0))
                    lhs165_.f22 = rhs184_
                    lhs166_.f2 = rhs185_
                    lhs167_.f22 = rhs186_
                    lhs168_[lhs169_] = rhs187_
                    d_1236_v12_ = rhs188_
                    pass
            pass
        r1 = p1
        d_1238_v14_: D2
        d_1238_v14_ = D2_DC5(self.f27, (self).f25, 756, p0)
        d_1239_v15_: _dafny.Map
        d_1239_v15_ = _dafny.Map({p0: D2_DC5(_dafny.CodePoint('o'), self.f24, p0, p0)})
        d_1240_v16_: _dafny.Seq
        d_1240_v16_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({p0: d_1238_v14_}), d_1239_v15_])
        d_1241_v17_: _dafny.Seq
        d_1241_v17_ = _dafny.SeqWithoutIsStrInference([d_1240_v16_])
        d_1242_v18_: D0
        d_1242_v18_ = D0_DC0((p0) <= (len((d_1241_v17_)[default__.safeIndex(p0, len(d_1241_v17_))])), p0, len(default__.fm32(globalState)))
        source22_ = d_1242_v18_
        d_1243___mcc_h0_ = source22_.cf0
        d_1244___mcc_h1_ = source22_.cf1
        d_1245___mcc_h2_ = source22_.cf2
        d_1246_cf2_ = d_1245___mcc_h2_
        d_1247_cf1_ = d_1244___mcc_h1_
        d_1248_cf0_ = d_1243___mcc_h0_
        d_1249_v19_: C1
        nw203_ = C1()
        nw203_.ctor__((self).f25)
        d_1249_v19_ = nw203_
        d_1250_v20_: _dafny.Array
        nw204_ = _dafny.Array(int(0), 18)
        d_1250_v20_ = nw204_
        index187_ = default__.safeIndex(816, (d_1250_v20_).length(0))
        (d_1250_v20_)[index187_] = p0
        d_1251_v21_: _dafny.Map
        d_1251_v21_ = _dafny.Map({d_1246_cf2_: self.f24})
        index188_ = default__.safeIndex(816, (d_1250_v20_).length(0))
        (d_1250_v20_)[index188_] = default__.fm0(p0, self.f24, (d_1251_v21_) | (d_1251_v21_), d_1246_cf2_, globalState)
        d_1252_v22_: _dafny.Seq
        d_1252_v22_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pqnv"))
        d_1252_v22_ = d_1252_v22_
        d_1253_v23_: _dafny.Array
        def lambda63_(d_1254_i1_):
            return self.f27

        init37_ = lambda63_
        nw205_ = _dafny.Array(None, 6)
        for i0_37_ in range(nw205_.length(0)):
            nw205_[i0_37_] = init37_(i0_37_)
        d_1253_v23_ = nw205_
        d_1255_v24_: _dafny.Set
        d_1255_v24_ = _dafny.Set({d_1253_v23_, d_1253_v23_})
        d_1256_v25_: D1
        d_1256_v25_ = D1_DC2(d_1250_v20_, self.f27, 648, (d_1255_v24_).issubset(d_1255_v24_))
        d_1256_v25_ = d_1256_v25_
        hi8_ = len(_dafny.SeqWithoutIsStrInference([(self).f26, (self).f26]))
        for d_1257_i2_ in range(p0, hi8_):
            d_1258_v26_: _dafny.Seq
            d_1258_v26_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u"))
            d_1259_v27_: C2
            nw206_ = C2()
            nw206_.ctor__(d_1258_v26_, ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "krn"))) + (_dafny.SeqWithoutIsStrInference([self.f27 for d_1260_i3_ in range(default__.abs(838))]))).set(default__.safeIndex(p0, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "krn"))) + (_dafny.SeqWithoutIsStrInference([self.f27 for d_1261_i3_ in range(default__.abs(838))])))), self.f27))
            d_1259_v27_ = nw206_
            source23_ = d_1238_v14_
            if source23_.is_DC4:
                d_1262___mcc_h3_ = source23_.cf9
                d_1263_cf9_ = d_1262___mcc_h3_
                (self).f24 = ((d_1259_v27_.f29) <= ((d_1258_v26_).set(default__.safeIndex(p0, len(d_1258_v26_)), self.f27))) and ((d_1257_i2_) < (d_1257_i2_))
                d_1264_v28_: _dafny.Array
                nw207_ = _dafny.Array(_dafny.Map({}), 27)
                d_1264_v28_ = nw207_
                index189_ = default__.safeIndex(785, (d_1264_v28_).length(0))
                def iife77_():
                    coll41_ = _dafny.Map()
                    compr_41_: int
                    for compr_41_ in _dafny.IntegerRange(496, 854):
                        d_1265_v29_: int = compr_41_
                        if ((496) <= (d_1265_v29_)) and ((d_1265_v29_) < (854)):
                            coll41_[(d_1265_v29_) + (d_1257_i2_)] = self.f27
                    return _dafny.Map(coll41_)
                (d_1264_v28_)[index189_] = iife77_()
                
                d_1266_v30_: _dafny.Map
                d_1266_v30_ = _dafny.Map({d_1257_i2_: self.f27})
                index190_ = default__.safeIndex(785, (d_1264_v28_).length(0))
                (d_1264_v28_)[index190_] = d_1266_v30_
                index191_ = default__.safeIndex(821, (p2).length(0))
                (p2)[index191_] = default__.fm1(p1, d_1263_cf9_, p0, globalState)
                index192_ = default__.safeIndex(821, (p2).length(0))
                (p2)[index192_] = (self).f26
                d_1267_v31_: C0
                nw208_ = C0()
                nw208_.ctor__(d_1259_v27_.f29)
                d_1267_v31_ = nw208_
            elif source23_.is_DC5:
                d_1268___mcc_h4_ = source23_.cf10
                d_1269___mcc_h5_ = source23_.cf11
                d_1270___mcc_h6_ = source23_.cf12
                d_1271___mcc_h7_ = source23_.cf13
                d_1272_cf13_ = d_1271___mcc_h7_
                d_1273_cf12_ = d_1270___mcc_h6_
                d_1274_cf11_ = d_1269___mcc_h5_
                d_1275_cf10_ = d_1268___mcc_h4_
                (globalState).f11 = ((0) - ((0) - (d_1272_cf13_)) if (d_1274_cf11_) or ((self).f26) else p0)
                d_1276_v32_: _dafny.Map
                d_1276_v32_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ednsxkv"))): d_1272_cf13_})
                d_1277_v33_: _dafny.Set
                d_1277_v33_ = _dafny.Set({d_1276_v32_, d_1276_v32_})
                r1 = (len(d_1277_v33_)) < (len(((d_1259_v27_).f30) + ((d_1259_v27_).f30)))
                d_1278_v34_: _dafny.Seq
                d_1278_v34_ = _dafny.SeqWithoutIsStrInference([d_1242_v18_])
                d_1278_v34_ = (d_1278_v34_) + (d_1278_v34_)
                r1 = d_1274_cf11_
            elif source23_.is_DC3:
                d_1279___mcc_h8_ = source23_.cf8
                d_1280_cf8_ = d_1279___mcc_h8_
                d_1281_v35_: _dafny.Map
                d_1281_v35_ = _dafny.Map({d_1257_i2_: (p0) < (d_1257_i2_)})
                d_1282_v36_: _dafny.Seq
                d_1282_v36_ = _dafny.SeqWithoutIsStrInference([p0, p0, p0, d_1257_i2_])
                r1 = ((d_1281_v35_)[(d_1282_v36_)[default__.safeIndex(p0, len(d_1282_v36_))]] if ((d_1282_v36_)[default__.safeIndex(p0, len(d_1282_v36_))]) in (d_1281_v35_) else not(self.f24))
                d_1283_v37_: _dafny.Map
                d_1283_v37_ = _dafny.Map({(d_1238_v14_).cf10: d_1257_i2_})
                d_1284_v38_: _dafny.Map
                d_1284_v38_ = _dafny.Map({self.f24: not(False)})
                d_1285_v39_: _dafny.Map
                d_1285_v39_ = _dafny.Map({self.f24: ((d_1284_v38_)[(self).f26] if ((self).f26) in (d_1284_v38_) else (self).f25)})
                d_1283_v37_ = (d_1283_v37_).set((d_1258_v26_)[default__.safeIndex(p0, len(d_1258_v26_))], default__.safeModuloInt(d_1257_i2_, len(d_1284_v38_)))
                d_1286_v40_: _dafny.Array
                nw209_ = _dafny.Array(int(0), 7)
                d_1286_v40_ = nw209_
                d_1287_v41_: D1
                d_1287_v41_ = D1_DC2(d_1286_v40_, self.f27, p0, p1)
                d_1288_v42_: _dafny.Seq
                d_1288_v42_ = _dafny.SeqWithoutIsStrInference([d_1242_v18_, d_1242_v18_])
                d_1289_v43_: _dafny.Map
                d_1289_v43_ = _dafny.Map({(self).f25: d_1288_v42_})
                d_1290_v44_: _dafny.Map
                d_1290_v44_ = _dafny.Map({self.f24: (self).fm5(d_1289_v43_, (self).f25, (self).f26, globalState)})
                d_1291_v45_: _dafny.Array
                nw210_ = _dafny.Array(_dafny.Array(None, 0), 27)
                d_1291_v45_ = nw210_
                d_1292_v46_: _dafny.Array
                d_1292_v46_ = d_1291_v45_
                d_1293_v47_: _dafny.Seq
                d_1293_v47_ = _dafny.SeqWithoutIsStrInference([d_1292_v46_, d_1292_v46_, d_1292_v46_])
                d_1294_v48_: D5
                d_1294_v48_ = D5_DC10(_dafny.MultiSet([d_1290_v44_]))
                d_1295_v49_: _dafny.MultiSet
                d_1295_v49_ = _dafny.MultiSet([True])
                d_1296_v50_: _dafny.Set
                d_1296_v50_ = _dafny.Set({self.f27})
                d_1297_v51_: _dafny.Set
                d_1297_v51_ = _dafny.Set({p0, ((d_1290_v44_)[True] if (True) in (d_1290_v44_) else len(d_1258_v26_)), d_1257_i2_})
                nw211_ = _dafny.Array(None, 26)
                nw211_[int(0)] = 469
                nw211_[int(1)] = d_1257_i2_
                def iife78_(_pat_let18_0):
                    def iife79_(d_1298_dt__update__tmp_h0_):
                        def iife80_(_pat_let19_0):
                            def iife81_(d_1299_dt__update_hcf7_h0_):
                                return D1_DC2((d_1298_dt__update__tmp_h0_).cf4, (d_1298_dt__update__tmp_h0_).cf5, (d_1298_dt__update__tmp_h0_).cf6, d_1299_dt__update_hcf7_h0_)
                            return iife81_(_pat_let19_0)
                        return iife80_((self).f26)
                    return iife79_(_pat_let18_0)
                nw211_[int(2)] = (iife78_(d_1287_v41_)).cf6
                nw211_[int(3)] = len((d_1290_v44_) | (d_1290_v44_))
                nw211_[int(4)] = (d_1257_i2_) + (p0)
                nw211_[int(5)] = (d_1282_v36_)[default__.safeIndex(258, len(d_1282_v36_))]
                nw211_[int(6)] = -467
                nw211_[int(7)] = (d_1257_i2_) + (d_1257_i2_)
                nw211_[int(8)] = len(d_1293_v47_)
                nw211_[int(9)] = 78
                nw211_[int(10)] = (d_1257_i2_) - (d_1257_i2_)
                nw211_[int(11)] = (((d_1280_cf8_)[d_1257_i2_] if (d_1257_i2_) in (d_1280_cf8_) else d_1257_i2_) if p1 else d_1257_i2_)
                nw211_[int(12)] = p0
                nw211_[int(13)] = (0) - (d_1257_i2_)
                nw211_[int(14)] = d_1257_i2_
                nw211_[int(15)] = (((d_1294_v48_).cf20).cardinality) * ((d_1295_v49_).cardinality)
                nw211_[int(16)] = (p0 if (self).f25 else len(d_1296_v50_))
                nw211_[int(17)] = default__.safeDivisionInt(p0, len(_dafny.Map({(self).f26: p0})))
                nw211_[int(18)] = default__.safeDivisionInt(d_1257_i2_, (d_1282_v36_)[default__.safeIndex(p0, len(d_1282_v36_))])
                nw211_[int(19)] = d_1257_i2_
                nw211_[int(20)] = 491
                nw211_[int(21)] = default__.safeModuloInt((d_1295_v49_).cardinality, 632)
                nw211_[int(22)] = d_1257_i2_
                nw211_[int(23)] = len((d_1297_v51_) - (d_1297_v51_))
                nw211_[int(24)] = default__.safeModuloInt(p0, d_1257_i2_)
                nw211_[int(25)] = 972
                (globalState).f7 = nw211_
                d_1300_v52_: _dafny.Seq
                d_1300_v52_ = _dafny.SeqWithoutIsStrInference([self.f24, (self).f26])
                rhs189_ = default__.safeDivisionInt(len(d_1300_v52_), d_1257_i2_)
                rhs190_ = d_1296_v50_
                rhs191_ = d_1282_v36_
                lhs170_ = globalState
                lhs170_.f22 = rhs189_
                d_1296_v50_ = rhs190_
                d_1282_v36_ = rhs191_
            elif True:
                d_1301___mcc_h9_ = source23_.cf14
                d_1302_cf14_ = d_1301___mcc_h9_
                (globalState).f11 = (d_1242_v18_).cf1
                d_1303_v53_: _dafny.MultiSet
                d_1303_v53_ = _dafny.MultiSet([(self).f26])
                d_1304_v54_: _dafny.Map
                d_1304_v54_ = _dafny.Map({default__.fm1((self).f25, (self).f25, d_1257_i2_, globalState): d_1257_i2_})
                d_1305_v55_: _dafny.Map
                d_1305_v55_ = _dafny.Map({((d_1303_v53_)[(self).fm7(d_1304_v54_, globalState)] if ((self).fm7(d_1304_v54_, globalState)) in (d_1303_v53_) else d_1257_i2_): self.f27})
                d_1306_v56_: _dafny.Seq
                d_1306_v56_ = _dafny.SeqWithoutIsStrInference([self.f24])
                rhs192_ = (d_1305_v55_) | (_dafny.Map({d_1257_i2_: self.f27}))
                rhs193_ = ((len((d_1259_v27_).f30)) * (len(d_1306_v56_)) if (self).f26 else d_1257_i2_)
                rhs194_ = p0
                lhs171_ = globalState
                lhs172_ = globalState
                d_1305_v55_ = rhs192_
                lhs171_.f11 = rhs193_
                lhs172_.f16 = rhs194_
                r0 = (d_1257_i2_) > (p0)
                d_1307_v57_: D1
                d_1307_v57_ = D1_DC1((d_1259_v27_).f30)
                d_1307_v57_ = d_1307_v57_
            d_1308_v58_: _dafny.Seq
            d_1308_v58_ = _dafny.SeqWithoutIsStrInference([d_1257_i2_])
            d_1309_v59_: D6
            d_1309_v59_ = D6_DC12(d_1308_v58_)
            pat_let_tv12_ = d_1308_v58_
            d_1310_v60_: C1
            nw212_ = C1()
            def iife82_(_pat_let20_0):
                def iife83_(d_1311_dt__update__tmp_h1_):
                    def iife84_(_pat_let21_0):
                        def iife85_(d_1312_dt__update_hcf23_h0_):
                            return D6_DC12(d_1312_dt__update_hcf23_h0_)
                        return iife85_(_pat_let21_0)
                    return iife84_(pat_let_tv12_)
                return iife83_(_pat_let20_0)
            nw212_.ctor__((d_1308_v58_) < ((iife82_(d_1309_v59_)).cf23))
            d_1310_v60_ = nw212_
            d_1313_v61_: _dafny.Map
            d_1313_v61_ = _dafny.Map({self.f24: self.f24})
            d_1314_v62_: _dafny.Map
            d_1314_v62_ = _dafny.Map({default__.safeModuloInt(p0, (d_1310_v60_).fm25((d_1259_v27_).f30, (d_1310_v60_).fm24(d_1257_i2_, d_1313_v61_, d_1257_i2_, 958, globalState), globalState)): d_1257_i2_})
            d_1314_v62_ = (d_1314_v62_).set(p0, d_1257_i2_)
        d_1315_v63_: _dafny.Seq
        d_1315_v63_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u"))
        hi9_ = default__.safeModuloInt((0) - (len(d_1315_v63_)), p0)
        for d_1316_i4_ in range(p0, hi9_):
            (globalState).f10 = (default__.safeModuloInt(-898, len(d_1315_v63_))) > (d_1316_i4_)
            d_1317_v65_: _dafny.Array
            def lambda64_(d_1318_p0_):
                def lambda65_(d_1319_i5_):
                    def iife86_():
                        coll42_ = _dafny.Set()
                        compr_42_: int
                        for compr_42_ in _dafny.IntegerRange(309, 386):
                            d_1320_v64_: int = compr_42_
                            if ((309) <= (d_1320_v64_)) and ((d_1320_v64_) < (386)):
                                coll42_ = coll42_.union(_dafny.Set([(d_1320_v64_) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pf"))))]))
                        return _dafny.Set(coll42_)
                    return (_dafny.Set({d_1318_p0_, d_1318_p0_, d_1318_p0_}) if self.f24 else iife86_()
                    )

                return lambda65_

            init38_ = lambda64_(p0)
            nw213_ = _dafny.Array(None, 21)
            for i0_38_ in range(nw213_.length(0)):
                nw213_[i0_38_] = init38_(i0_38_)
            d_1317_v65_ = nw213_
            d_1321_v66_: _dafny.Set
            d_1321_v66_ = _dafny.Set({p0})
            index193_ = default__.safeIndex(933, (d_1317_v65_).length(0))
            (d_1317_v65_)[index193_] = d_1321_v66_
            index194_ = default__.safeIndex(933, (d_1317_v65_).length(0))
            (d_1317_v65_)[index194_] = d_1321_v66_
            d_1322_v67_: _dafny.Map
            d_1322_v67_ = _dafny.Map({p0: d_1316_i4_})
            d_1323_v68_: _dafny.Map
            d_1323_v68_ = _dafny.Map({d_1322_v67_: not(True)})
            d_1324_v70_: _dafny.Set
            d_1324_v70_ = _dafny.Set({d_1322_v67_})
            index195_ = default__.safeIndex(417, (p2).length(0))
            def iife87_():
                coll43_ = _dafny.Set()
                compr_43_: _dafny.Map
                for compr_43_ in (d_1323_v68_).keys.Elements:
                    d_1325_v69_: _dafny.Map = compr_43_
                    if (d_1325_v69_) in (d_1323_v68_):
                        coll43_ = coll43_.union(_dafny.Set([d_1325_v69_]))
                return _dafny.Set(coll43_)
            (p2)[index195_] = (iife87_()
            ).issubset(d_1324_v70_)
            index196_ = default__.safeIndex(417, (p2).length(0))
            (p2)[index196_] = (self).f26
            d_1326_v71_: _dafny.Seq
            d_1326_v71_ = _dafny.SeqWithoutIsStrInference([(self).f25])
            d_1327_v72_: _dafny.MultiSet
            d_1327_v72_ = _dafny.MultiSet([d_1316_i4_, d_1316_i4_])
            d_1328_v73_: _dafny.Array
            nw214_ = _dafny.Array(None, 4)
            nw214_[int(0)] = d_1326_v71_
            nw214_[int(1)] = (d_1326_v71_) + (d_1326_v71_)
            nw214_[int(2)] = (d_1326_v71_) + (d_1326_v71_)
            nw214_[int(3)] = default__.fm33(self.f27, (d_1327_v72_).cardinality, (self).f25, (self).f26, globalState)
            d_1328_v73_ = nw214_
            index197_ = default__.safeIndex(394, (d_1328_v73_).length(0))
            (d_1328_v73_)[index197_] = d_1326_v71_
            d_1329_v74_: _dafny.MultiSet
            d_1329_v74_ = _dafny.MultiSet([default__.fm1((self).f25, (self).f26, d_1316_i4_, globalState), (self).f25, self.f24, (self).f25, (p2)[default__.safeIndex(417, (p2).length(0))]])
            index198_ = default__.safeIndex(394, (d_1328_v73_).length(0))
            (d_1328_v73_)[index198_] = (d_1326_v71_).set(default__.safeIndex((d_1329_v74_).cardinality, len(d_1326_v71_)), (self).f25)
        if (self).f25:
            d_1330_v75_: _dafny.Map
            d_1330_v75_ = _dafny.Map({not (p1) or ((self).f25): (self).f25})
            d_1331_v76_: _dafny.Array
            nw215_ = _dafny.Array(int(0), 7)
            d_1331_v76_ = nw215_
            d_1332_v77_: _dafny.Seq
            d_1332_v77_ = _dafny.SeqWithoutIsStrInference([d_1242_v18_])
            d_1333_v78_: _dafny.Map
            d_1333_v78_ = _dafny.Map({p1: (d_1332_v77_).set(default__.safeIndex(p0, len(d_1332_v77_)), d_1242_v18_)})
            d_1334_v79_: _dafny.Seq
            d_1334_v79_ = _dafny.SeqWithoutIsStrInference([(self).f26])
            d_1335_v81_: _dafny.Map
            d_1335_v81_ = _dafny.Map({p0: p1})
            d_1336_v82_: _dafny.MultiSet
            d_1336_v82_ = _dafny.MultiSet([p0])
            d_1337_v83_: D1
            def iife88_():
                coll44_ = _dafny.Map()
                compr_44_: int
                for compr_44_ in (d_1335_v81_).keys.Elements:
                    d_1338_v80_: int = compr_44_
                    if (d_1338_v80_) in (d_1335_v81_):
                        coll44_[default__.safeDivisionInt(d_1338_v80_, 183)] = self.f24
                return _dafny.Map(coll44_)
            d_1337_v83_ = D1_DC2(d_1331_v76_, self.f27, default__.fm0((self).fm5(d_1333_v78_, not((d_1334_v79_)[default__.safeIndex(p0, len(d_1334_v79_))]), self.f24, globalState), False, iife88_()
, len(_dafny.Map({not((self).f25): d_1336_v82_})), globalState), (self).f26)
            d_1330_v75_ = (d_1330_v75_).set(not((p0) >= (p0)), (d_1337_v83_).cf7)
            d_1339_v84_: _dafny.Map
            d_1339_v84_ = _dafny.Map({(0) - (p0): p0})
            d_1340_v85_: _dafny.Seq
            d_1340_v85_ = _dafny.SeqWithoutIsStrInference([p0, p0, (d_1337_v83_).cf6])
            d_1341_v86_: D3
            d_1341_v86_ = D3_DC8((not((self).f25)) and (((d_1335_v81_)[len(d_1339_v84_)] if (len(d_1339_v84_)) in (d_1335_v81_) else False)), p0, len((_dafny.SeqWithoutIsStrInference([len(d_1315_v63_) for d_1342_i6_ in range(default__.abs(160))])) + (d_1340_v85_)))
            source24_ = d_1341_v86_
            if source24_.is_DC8:
                d_1343___mcc_h10_ = source24_.cf16
                d_1344___mcc_h11_ = source24_.cf17
                d_1345___mcc_h12_ = source24_.cf18
                d_1346_cf18_ = d_1345___mcc_h12_
                d_1347_cf17_ = d_1344___mcc_h11_
                d_1348_cf16_ = d_1343___mcc_h10_
                def iife89_(_pat_let22_0):
                    def iife90_(d_1349_dt__update__tmp_h2_):
                        def iife91_(_pat_let23_0):
                            def iife92_(d_1350_dt__update_hcf10_h0_):
                                return D2_DC5(d_1350_dt__update_hcf10_h0_, (d_1349_dt__update__tmp_h2_).cf11, (d_1349_dt__update__tmp_h2_).cf12, (d_1349_dt__update__tmp_h2_).cf13)
                            return iife92_(_pat_let23_0)
                        return iife91_(self.f27)
                    return iife90_(_pat_let22_0)
                rhs195_ = (p0) * (d_1347_cf17_)
                rhs196_ = _dafny.CodePoint('u')
                rhs197_ = iife89_(d_1238_v14_)
                lhs173_ = globalState
                lhs174_ = self
                lhs173_.f22 = rhs195_
                lhs174_.f27 = rhs196_
                d_1238_v14_ = rhs197_
                d_1351_v87_: C0
                nw216_ = C0()
                nw216_.ctor__((d_1315_v63_ if (self).f26 else d_1315_v63_))
                d_1351_v87_ = nw216_
                d_1351_v87_ = d_1351_v87_
                d_1352_v88_: _dafny.Seq
                d_1352_v88_ = _dafny.SeqWithoutIsStrInference([(d_1351_v87_).f31, d_1315_v63_, _dafny.SeqWithoutIsStrInference([self.f27]), d_1315_v63_])
                d_1353_v89_: _dafny.Set
                d_1353_v89_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([self.f27]), (d_1352_v88_)[default__.safeIndex((d_1242_v18_).cf2, len(d_1352_v88_))]})
                d_1354_v90_: _dafny.MultiSet
                d_1354_v90_ = _dafny.MultiSet([not((d_1353_v89_).issubset(d_1353_v89_))])
                d_1354_v90_ = (d_1354_v90_) | ((d_1354_v90_).intersection(d_1354_v90_))
                d_1355_v91_: _dafny.Map
                d_1355_v91_ = _dafny.Map({_dafny.MultiSet([self.f27, self.f27, self.f27, self.f27]): _dafny.MultiSet([(self).f25, self.f24, (self).f25])})
                d_1355_v91_ = _dafny.Map({default__.fm34((self).f25, globalState): d_1354_v90_})
            elif True:
                d_1356___mcc_h13_ = source24_.cf15
                d_1357_cf15_ = d_1356___mcc_h13_
                d_1358_v92_: _dafny.Seq
                d_1358_v92_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p0, p0])])
                d_1358_v92_ = d_1358_v92_
                d_1359_v93_: _dafny.Map
                d_1359_v93_ = _dafny.Map({d_1315_v63_: p0})
                d_1359_v93_ = (d_1359_v93_).set(d_1315_v63_, default__.safeDivisionInt(default__.fm0(-910, (self).f26, d_1335_v81_, p0, globalState), 868))
                d_1360_v94_: _dafny.Array
                nw217_ = _dafny.Array(_dafny.Seq({}), 25)
                d_1360_v94_ = nw217_
                rhs198_ = self.f27
                rhs199_ = True
                rhs200_ = d_1360_v94_
                lhs175_ = self
                lhs176_ = globalState
                lhs175_.f27 = rhs198_
                lhs176_.f10 = rhs199_
                d_1360_v94_ = rhs200_
                (globalState).f11 = ((0) - ((p0) + (667))) + (p0)
            r0 = p1
            d_1361_v95_: _dafny.Set
            d_1361_v95_ = _dafny.Set({((d_1336_v82_).set(p0, default__.abs(len(d_1315_v63_)))).cardinality, len(_dafny.Map({(d_1340_v85_).set(default__.safeIndex(p0, len(d_1340_v85_)), p0): False}))})
            d_1362_v97_: D7
            d_1362_v97_ = D7_DC14(d_1361_v95_)
            d_1363_v100_: T1
            nw218_ = C1()
            nw218_.ctor__((self).f26)
            d_1363_v100_ = nw218_
            d_1364_v101_: _dafny.Map
            d_1364_v101_ = _dafny.Map({(self).f26: d_1363_v100_})
            d_1365_v102_: _dafny.Seq
            d_1365_v102_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({len(d_1364_v101_): p0}), d_1339_v84_])
            d_1366_v103_: _dafny.Map
            d_1366_v103_ = _dafny.Map({d_1363_v100_.f24: (d_1315_v63_)[default__.safeIndex(p0, len(d_1315_v63_))]})
            d_1367_v108_: _dafny.Array
            nw219_ = _dafny.Array(None, 28)
            nw219_[int(0)] = d_1339_v84_
            def iife93_():
                coll45_ = _dafny.Map()
                compr_45_: int
                for compr_45_ in ((d_1362_v97_).cf27).Elements:
                    d_1368_v96_: int = compr_45_
                    if (d_1368_v96_) in ((d_1362_v97_).cf27):
                        def iife94_():
                            coll46_ = _dafny.Map()
                            compr_46_: _dafny.Map
                            for compr_46_ in (_dafny.Set({_dafny.Map({p0: p0})})).Elements:
                                d_1369_v98_: _dafny.Map = compr_46_
                                if (d_1369_v98_) in (_dafny.Set({_dafny.Map({p0: p0})})):
                                    coll46_[d_1369_v98_] = len(d_1334_v79_)
                            return _dafny.Map(coll46_)
                        coll45_[default__.safeModuloInt(d_1368_v96_, p0)] = len(iife94_()
                        )
                return _dafny.Map(coll45_)
            nw219_[int(1)] = iife93_()
            
            def iife95_():
                coll47_ = _dafny.Map()
                compr_47_: int
                for compr_47_ in (d_1335_v81_).keys.Elements:
                    d_1370_v99_: int = compr_47_
                    if (d_1370_v99_) in (d_1335_v81_):
                        coll47_[(d_1370_v99_) - (p0)] = (0) - (p0)
                return _dafny.Map(coll47_)
            nw219_[int(2)] = iife95_()
            
            nw219_[int(3)] = (d_1365_v102_)[default__.safeIndex(p0, len(d_1365_v102_))]
            nw219_[int(4)] = d_1339_v84_
            def iife96_():
                coll48_ = _dafny.Map()
                compr_48_: int
                for compr_48_ in (d_1335_v81_).keys.Elements:
                    d_1371_v104_: int = compr_48_
                    if (d_1371_v104_) in (d_1335_v81_):
                        coll48_[default__.safeModuloInt(d_1371_v104_, p0)] = p0
                return _dafny.Map(coll48_)
            nw219_[int(5)] = (default__.fm35(False, d_1366_v103_, globalState)) | (iife96_()
            )
            nw219_[int(6)] = d_1339_v84_
            nw219_[int(7)] = d_1339_v84_
            def iife97_():
                coll49_ = _dafny.Map()
                compr_49_: int
                for compr_49_ in (d_1340_v85_).Elements:
                    d_1372_v105_: int = compr_49_
                    if (d_1372_v105_) in (d_1340_v85_):
                        coll49_[(d_1372_v105_) * (p0)] = len(d_1334_v79_)
                return _dafny.Map(coll49_)
            nw219_[int(8)] = iife97_()
            
            nw219_[int(9)] = d_1339_v84_
            nw219_[int(10)] = d_1339_v84_
            nw219_[int(11)] = (d_1339_v84_).set(p0, len(_dafny.SeqWithoutIsStrInference([len(d_1334_v79_), p0, len(_dafny.SeqWithoutIsStrInference([p2])), p0, p0])))
            def iife98_():
                coll50_ = _dafny.Map()
                compr_50_: int
                for compr_50_ in _dafny.IntegerRange(604, 624):
                    d_1373_v106_: int = compr_50_
                    if ((604) <= (d_1373_v106_)) and ((d_1373_v106_) < (624)):
                        coll50_[(d_1373_v106_) + (p0)] = len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([p0]))]))
                return _dafny.Map(coll50_)
            nw219_[int(12)] = (iife98_()
            ).set((0) - (len(d_1361_v95_)), len(_dafny.SeqWithoutIsStrInference([self.f24, d_1363_v100_.f24])))
            nw219_[int(13)] = d_1339_v84_
            nw219_[int(14)] = (d_1339_v84_) | (d_1339_v84_)
            nw219_[int(15)] = d_1339_v84_
            nw219_[int(16)] = d_1339_v84_
            nw219_[int(17)] = d_1339_v84_
            nw219_[int(18)] = d_1339_v84_
            nw219_[int(19)] = d_1339_v84_
            nw219_[int(20)] = d_1339_v84_
            nw219_[int(21)] = (d_1339_v84_) | (_dafny.Map({-90: len(default__.fm31(globalState))}))
            nw219_[int(22)] = d_1339_v84_
            def iife99_():
                coll51_ = _dafny.Map()
                compr_51_: int
                for compr_51_ in (d_1361_v95_).Elements:
                    d_1374_v107_: int = compr_51_
                    if (d_1374_v107_) in (d_1361_v95_):
                        coll51_[default__.safeModuloInt(d_1374_v107_, p0)] = p0
                return _dafny.Map(coll51_)
            nw219_[int(23)] = iife99_()
            
            nw219_[int(24)] = (d_1339_v84_).set(p0, p0)
            nw219_[int(25)] = d_1339_v84_
            nw219_[int(26)] = d_1339_v84_
            nw219_[int(27)] = default__.fm35((self).f25, d_1366_v103_, globalState)
            d_1367_v108_ = nw219_
            index199_ = default__.safeIndex(327, (p2).length(0))
            (p2)[index199_] = False
            d_1375_v109_: _dafny.Map
            d_1375_v109_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([(0) - (p0) for d_1376_i7_ in range(default__.abs(353))]): d_1361_v95_})
            d_1377_v110_: _dafny.Map
            d_1377_v110_ = _dafny.Map({(self).f25: 382})
            index200_ = default__.safeIndex(327, (p2).length(0))
            rhs201_ = ((d_1375_v109_)[d_1340_v85_] if (d_1340_v85_) in (d_1375_v109_) else d_1361_v95_)
            rhs202_ = (self).fm7(d_1377_v110_, globalState)
            rhs203_ = d_1367_v108_
            rhs204_ = ((p0) - (default__.fm0(p0, not(d_1363_v100_.f24), d_1335_v81_, (d_1336_v82_).cardinality, globalState))) != (len(((d_1334_v79_) + (_dafny.SeqWithoutIsStrInference([d_1363_v100_.f24]))).set(default__.safeIndex(p0, len((d_1334_v79_) + (_dafny.SeqWithoutIsStrInference([d_1363_v100_.f24])))), (self).f25)))
            lhs177_ = p2
            lhs178_ = default__.safeIndex(327, (p2).length(0))
            d_1361_v95_ = rhs201_
            r0 = rhs202_
            d_1367_v108_ = rhs203_
            lhs177_[lhs178_] = rhs204_
            index201_ = default__.safeIndex(327, (p2).length(0))
            (p2)[index201_] = d_1363_v100_.f24
        elif True:
            d_1378_v111_: D6
            d_1378_v111_ = D6_DC13(p2, p0, (self).f26)
            (globalState).f16 = (d_1378_v111_).cf25
            d_1379_v112_: _dafny.Map
            d_1379_v112_ = _dafny.Map({self.f27: p0})
            d_1380_v113_: _dafny.Map
            d_1380_v113_ = _dafny.Map({p0: p0})
            d_1381_v114_: _dafny.Map
            d_1381_v114_ = _dafny.Map({p0: d_1315_v63_})
            d_1382_v115_: _dafny.MultiSet
            d_1382_v115_ = _dafny.MultiSet([default__.fm36(p0, (self).f26, self.f24, globalState), self.f27])
            d_1383_v116_: _dafny.MultiSet
            d_1383_v116_ = _dafny.MultiSet([p0])
            d_1384_v117_: _dafny.Seq
            d_1384_v117_ = _dafny.SeqWithoutIsStrInference([d_1383_v116_])
            d_1385_v118_: _dafny.Array
            nw220_ = _dafny.Array(None, 22)
            nw220_[int(0)] = -653
            nw220_[int(1)] = p0
            nw220_[int(2)] = p0
            nw220_[int(3)] = p0
            nw220_[int(4)] = ((d_1379_v112_)[_dafny.CodePoint('d')] if (_dafny.CodePoint('d')) in (d_1379_v112_) else (0) - ((0) - (len(d_1380_v113_))))
            nw220_[int(5)] = p0
            nw220_[int(6)] = p0
            nw220_[int(7)] = p0
            nw220_[int(8)] = p0
            nw220_[int(9)] = p0
            nw220_[int(10)] = p0
            nw220_[int(11)] = p0
            nw220_[int(12)] = len(d_1381_v114_)
            nw220_[int(13)] = p0
            nw220_[int(14)] = p0
            nw220_[int(15)] = p0
            nw220_[int(16)] = p0
            nw220_[int(17)] = (0) - ((0) - (((d_1382_v115_)[self.f27] if (self.f27) in (d_1382_v115_) else p0)))
            nw220_[int(18)] = len(d_1384_v117_)
            nw220_[int(19)] = p0
            nw220_[int(20)] = p0
            nw220_[int(21)] = p0
            d_1385_v118_ = nw220_
            d_1386_v119_: _dafny.Map
            d_1386_v119_ = _dafny.Map({d_1385_v118_: (self).f26})
            d_1387_v120_: _dafny.Map
            d_1387_v120_ = _dafny.Map({self.f27: d_1386_v119_})
            d_1386_v119_ = (d_1386_v119_) | (((d_1387_v120_)[_dafny.CodePoint('d')] if (_dafny.CodePoint('d')) in (d_1387_v120_) else d_1386_v119_))
            d_1388_v121_: _dafny.Seq
            d_1388_v121_ = _dafny.SeqWithoutIsStrInference([self.f24, p1])
            index202_ = default__.safeIndex(656, (p2).length(0))
            (p2)[index202_] = (d_1388_v121_)[default__.safeIndex(len(d_1315_v63_), len(d_1388_v121_))]
            index203_ = default__.safeIndex(656, (p2).length(0))
            (p2)[index203_] = p1
            d_1389_v122_: _dafny.Map
            d_1389_v122_ = _dafny.Map({(self).f25: self.f24})
            d_1390_v123_: _dafny.Seq
            d_1390_v123_ = _dafny.SeqWithoutIsStrInference([d_1389_v122_])
            d_1391_v124_: _dafny.MultiSet
            d_1391_v124_ = _dafny.MultiSet([p1, (self).f25, self.f24])
            d_1392_v125_: _dafny.Map
            d_1392_v125_ = _dafny.Map({len(d_1390_v123_): d_1391_v124_})
            d_1392_v125_ = (d_1392_v125_).set(p0, d_1391_v124_)
            d_1393_v126_: _dafny.Map
            d_1393_v126_ = _dafny.Map({(self).f25: p0})
            (globalState).f10 = (self).fm7(d_1393_v126_, globalState)
        d_1394_v127_: _dafny.Map
        d_1394_v127_ = _dafny.Map({p0: p1})
        d_1395_v128_: _dafny.MultiSet
        d_1395_v128_ = _dafny.MultiSet([default__.fm1(self.f24, (self).f25, p0, globalState), ((d_1394_v127_)[p0] if (p0) in (d_1394_v127_) else (self).f25)])
        d_1396_v129_: _dafny.Set
        d_1396_v129_ = _dafny.Set({p0, p0, (d_1395_v128_).cardinality, len(d_1315_v63_)})
        r0 = (d_1396_v129_).issubset(d_1396_v129_)
        r1 = (p0) == (p0)
        r2 = _dafny.Map({d_1315_v63_: (p0 if p1 else p0)})
        return r0, r1, r2

    def m0(self, globalState):
        r0: bool = False
        r1: _dafny.Array = _dafny.Array(None, 0)
        d_1397_v0_: int
        d_1397_v0_ = 822
        if (d_1397_v0_) > (d_1397_v0_):
            (globalState).f10 = True
            d_1398_v1_: _dafny.Seq
            d_1398_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hunkhrwn"))
            d_1399_v2_: D0
            d_1399_v2_ = D0_DC0((self).f25, len(d_1398_v1_), d_1397_v0_)
            d_1400_v3_: _dafny.Map
            d_1400_v3_ = _dafny.Map({(self).fm5(_dafny.Map({True: _dafny.SeqWithoutIsStrInference([d_1399_v2_])}), self.f24, False, globalState): (self).f26})
            d_1400_v3_ = (d_1400_v3_).set((0) - (len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pbww"))) + (d_1398_v1_))), (self).f25)
            d_1401_v4_: _dafny.Array
            def lambda66_(d_1402_i0_):
                return False

            init39_ = lambda66_
            nw221_ = _dafny.Array(None, 26)
            for i0_39_ in range(nw221_.length(0)):
                nw221_[i0_39_] = init39_(i0_39_)
            d_1401_v4_ = nw221_
            d_1403_v5_: _dafny.MultiSet
            d_1403_v5_ = _dafny.MultiSet([(0) - (len(d_1398_v1_)), 91, d_1397_v0_])
            d_1404_v6_: _dafny.Seq
            d_1404_v6_ = _dafny.SeqWithoutIsStrInference([len(d_1398_v1_)])
            d_1405_v7_: T2
            nw222_ = C4()
            nw222_.ctor__((self).f26, self.f24)
            d_1405_v7_ = nw222_
            d_1406_v8_: _dafny.MultiSet
            d_1406_v8_ = _dafny.MultiSet([d_1405_v7_, d_1405_v7_])
            d_1407_v9_: _dafny.Map
            d_1407_v9_ = _dafny.Map({d_1406_v8_: d_1397_v0_})
            d_1408_v10_: _dafny.Map
            d_1408_v10_ = _dafny.Map({(d_1405_v7_).f25: self.f24})
            d_1409_v11_: _dafny.Set
            d_1409_v11_ = _dafny.Set({d_1403_v5_, d_1403_v5_, _dafny.MultiSet([d_1397_v0_, len(d_1404_v6_), len(d_1407_v9_), d_1397_v0_, len(d_1408_v10_)]), d_1403_v5_})
            d_1410_v12_: _dafny.Map
            d_1410_v12_ = _dafny.Map({d_1401_v4_: d_1409_v11_})
            d_1410_v12_ = (d_1410_v12_).set(d_1401_v4_, d_1409_v11_)
            (globalState).f10 = (d_1397_v0_) <= (default__.safeDivisionInt(d_1397_v0_, 573))
            (globalState).f16 = (0) - (len(d_1408_v10_))
        elif True:
            d_1411_v13_: _dafny.Seq
            d_1411_v13_ = _dafny.SeqWithoutIsStrInference([898])
            d_1412_v15_: _dafny.Set
            d_1412_v15_ = _dafny.Set({d_1397_v0_})
            def iife100_():
                coll52_ = _dafny.Set()
                compr_52_: int
                for compr_52_ in (d_1411_v13_).Elements:
                    d_1413_v14_: int = compr_52_
                    if (d_1413_v14_) in (d_1411_v13_):
                        coll52_ = coll52_.union(_dafny.Set([(d_1413_v14_) * ((0) - (len(_dafny.Map({len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ibum"))): 486})): True}))))]))
                return _dafny.Set(coll52_)
            rhs205_ = (self).f25
            rhs206_ = ((self).f25 if (self).f25 else (iife100_()
            ) != (d_1412_v15_))
            lhs179_ = globalState
            lhs180_ = globalState
            lhs179_.f10 = rhs205_
            lhs180_.f10 = rhs206_
            d_1414_v16_: _dafny.Set
            d_1414_v16_ = _dafny.Set({(self).f26, (self).f25, (self).f25, self.f24, self.f24})
            d_1415_v17_: _dafny.Map
            d_1415_v17_ = _dafny.Map({self.f27: len(d_1411_v13_)})
            d_1416_v18_: _dafny.MultiSet
            d_1416_v18_ = _dafny.MultiSet([145, len(_dafny.SeqWithoutIsStrInference([d_1397_v0_])), d_1397_v0_])
            d_1417_v19_: _dafny.MultiSet
            d_1417_v19_ = _dafny.MultiSet([(self).f25])
            d_1418_v20_: _dafny.Map
            d_1418_v20_ = _dafny.Map({self.f24: ((d_1416_v18_)[(0) - (d_1397_v0_)] if ((0) - (d_1397_v0_)) in (d_1416_v18_) else (d_1417_v19_).cardinality)})
            d_1419_v21_: _dafny.MultiSet
            d_1419_v21_ = _dafny.MultiSet([d_1418_v20_])
            d_1420_v22_: D5
            d_1420_v22_ = D5_DC10(d_1419_v21_)
            source25_ = (d_1420_v22_ if (default__.fm36(len(d_1414_v16_), False, (self).f25, globalState)) in (d_1415_v17_) else D5_DC10(d_1419_v21_))
            if source25_.is_DC11:
                d_1421___mcc_h0_ = source25_.cf21
                d_1422___mcc_h1_ = source25_.cf22
                d_1423_cf22_ = d_1422___mcc_h1_
                d_1424_cf21_ = d_1421___mcc_h0_
                d_1397_v0_ = d_1397_v0_
                d_1425_v23_: _dafny.Array
                nw223_ = _dafny.Array(_dafny.Array(None, 0), 18)
                d_1425_v23_ = nw223_
                d_1426_v24_: _dafny.Array
                nw224_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 9)
                d_1426_v24_ = nw224_
                d_1427_v25_: _dafny.Array
                d_1427_v25_ = d_1426_v24_
                index204_ = default__.safeIndex(210, (d_1425_v23_).length(0))
                (d_1425_v23_)[index204_] = (d_1427_v25_)
                index205_ = default__.safeIndex(210, (d_1425_v23_).length(0))
                (d_1425_v23_)[index205_] = d_1426_v24_
                d_1428_v26_: _dafny.Seq
                d_1428_v26_ = _dafny.SeqWithoutIsStrInference([d_1423_cf22_])
                d_1429_v27_: _dafny.Set
                d_1429_v27_ = _dafny.Set({d_1411_v13_, _dafny.SeqWithoutIsStrInference([d_1397_v0_ for d_1430_i1_ in range(default__.abs(596))])})
                d_1431_v28_: D7
                d_1431_v28_ = D7_DC16((self).f26, d_1397_v0_, (self).f26, (d_1428_v26_)[default__.safeIndex(d_1397_v0_, len(d_1428_v26_))], (d_1429_v27_).issubset(default__.fm56(globalState)))
                d_1431_v28_ = D7_DC16(d_1423_cf22_, d_1397_v0_, not(self.f24), (self).f26, (d_1397_v0_) > (d_1397_v0_))
                r0 = False
            elif True:
                d_1432___mcc_h2_ = source25_.cf20
                d_1433_cf20_ = d_1432___mcc_h2_
                (globalState).f10 = (d_1416_v18_).ispropersubset(d_1416_v18_)
                r0 = (self.f24 if (self).f26 else self.f24)
                d_1434_v29_: _dafny.Array
                nw225_ = _dafny.Array(_dafny.CodePoint('D'), 15)
                d_1434_v29_ = nw225_
                index206_ = default__.safeIndex(760, (d_1434_v29_).length(0))
                (d_1434_v29_)[index206_] = self.f27
                index207_ = default__.safeIndex(760, (d_1434_v29_).length(0))
                (d_1434_v29_)[index207_] = self.f27
                d_1435_v30_: _dafny.Seq
                d_1435_v30_ = _dafny.SeqWithoutIsStrInference([self.f24, (self).f26, (self).f26, self.f24, (self).f25])
                d_1436_v31_: _dafny.Seq
                d_1436_v31_ = _dafny.SeqWithoutIsStrInference([(d_1435_v30_)[default__.safeIndex(d_1397_v0_, len(d_1435_v30_))]])
                (self).f24 = (d_1436_v31_) != (d_1435_v30_)
            (self).f24 = (self).f25
            d_1437_v32_: C0
            nw226_ = C0()
            nw226_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hlntshv")))
            d_1437_v32_ = nw226_
            d_1438_v33_: _dafny.Set
            d_1438_v33_ = _dafny.Set({default__.fm36(len(d_1412_v15_), (self).f26, self.f24, globalState), self.f27, self.f27, self.f27})
            d_1438_v33_ = _dafny.Set({(default__.fm36(d_1397_v0_, self.f24, True, globalState) if self.f24 else default__.fm36(d_1397_v0_, (self).f25, self.f24, globalState))})
        d_1439_v34_: _dafny.Set
        d_1439_v34_ = _dafny.Set({not((self).f26), not(False), self.f24, True})
        d_1440_v35_: C4
        nw227_ = C4()
        nw227_.ctor__((self).f25, (d_1439_v34_).ispropersubset(d_1439_v34_))
        d_1440_v35_ = nw227_
        d_1441_v36_: D0
        d_1441_v36_ = D0_DC0(self.f24, d_1397_v0_, d_1397_v0_)
        if (((d_1440_v35_).fm6(d_1441_v36_, globalState)) * (310)) != (d_1397_v0_):
            d_1442_v37_: _dafny.MultiSet
            d_1442_v37_ = _dafny.MultiSet([(self).f25])
            d_1443_v38_: _dafny.Seq
            d_1443_v38_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eeusm"))
            d_1444_v39_: _dafny.MultiSet
            d_1444_v39_ = _dafny.MultiSet([(d_1442_v37_).cardinality, (0) - (len(d_1443_v38_)), d_1397_v0_])
            d_1445_v40_: _dafny.Set
            d_1445_v40_ = _dafny.Set({394, d_1397_v0_, d_1397_v0_})
            d_1446_v41_: _dafny.Map
            d_1446_v41_ = _dafny.Map({(d_1397_v0_) != (d_1397_v0_): ((d_1444_v39_)[d_1397_v0_] if (d_1397_v0_) in (d_1444_v39_) else len(d_1445_v40_))})
            (globalState).f16 = ((d_1446_v41_)[(self).f25] if ((self).f25) in (d_1446_v41_) else (d_1397_v0_) * (d_1397_v0_))
            d_1447_v42_: _dafny.Set
            d_1447_v42_ = _dafny.Set({d_1443_v38_})
            d_1448_v43_: _dafny.Map
            d_1448_v43_ = _dafny.Map({(d_1443_v38_) + (d_1443_v38_): default__.safeDivisionInt(len(d_1447_v42_), d_1397_v0_)})
            d_1449_v44_: _dafny.Map
            d_1449_v44_ = _dafny.Map({d_1397_v0_: True})
            d_1450_v46_: _dafny.Map
            def iife101_():
                coll53_ = _dafny.Map()
                compr_53_: int
                for compr_53_ in (d_1449_v44_).keys.Elements:
                    d_1451_v45_: int = compr_53_
                    if (d_1451_v45_) in (d_1449_v44_):
                        coll53_[default__.safeDivisionInt(d_1451_v45_, d_1397_v0_)] = d_1443_v38_
                return _dafny.Map(coll53_)
            d_1450_v46_ = _dafny.Map({len(iife101_()
            ): d_1397_v0_})
            def iife102_():
                coll54_ = _dafny.Set()
                compr_54_: int
                for compr_54_ in (d_1450_v46_).keys.Elements:
                    d_1454_v47_: int = compr_54_
                    if (d_1454_v47_) in (d_1450_v46_):
                        coll54_ = coll54_.union(_dafny.Set([(d_1454_v47_) - (839)]))
                return _dafny.Set(coll54_)
            d_1448_v43_ = (d_1448_v43_).set((_dafny.SeqWithoutIsStrInference([(D2_DC5(self.f27, (self).f26, d_1397_v0_, d_1397_v0_)).cf10 for d_1452_i2_ in range(default__.abs(405))])).set(default__.safeIndex(d_1397_v0_, len(_dafny.SeqWithoutIsStrInference([(D2_DC5(self.f27, (self).f26, d_1397_v0_, d_1397_v0_)).cf10 for d_1453_i2_ in range(default__.abs(405))]))), _dafny.CodePoint('p')), len(default__.fm57((0) - (d_1397_v0_), d_1449_v44_, iife102_()
            , globalState)))
            d_1455_v48_: C4
            nw228_ = C4()
            nw228_.ctor__(((self).f26) or ((self).f25), (self).f26)
            d_1455_v48_ = nw228_
            d_1456_v49_: _dafny.Array
            nw229_ = _dafny.Array(_dafny.MultiSet({}), 4)
            d_1456_v49_ = nw229_
            index208_ = default__.safeIndex(362, (d_1456_v49_).length(0))
            (d_1456_v49_)[index208_] = d_1444_v39_
            d_1457_v50_: _dafny.Array
            nw230_ = _dafny.Array(int(0), 8)
            d_1457_v50_ = nw230_
            d_1458_v51_: _dafny.Set
            d_1458_v51_ = _dafny.Set({d_1457_v50_})
            d_1459_v52_: _dafny.Seq
            d_1459_v52_ = _dafny.SeqWithoutIsStrInference([len(d_1458_v51_)])
            d_1460_v53_: _dafny.Map
            d_1460_v53_ = _dafny.Map({d_1397_v0_: _dafny.MultiSet(d_1459_v52_)})
            index209_ = default__.safeIndex(362, (d_1456_v49_).length(0))
            rhs207_ = (((d_1460_v53_)[d_1397_v0_] if (d_1397_v0_) in (d_1460_v53_) else d_1444_v39_)) - (d_1444_v39_)
            rhs208_ = d_1397_v0_
            rhs209_ = (0) - ((0) - (((len(_dafny.Map({d_1397_v0_: (self).f25}))) - (d_1397_v0_)) * (len((_dafny.Map({(self).f25: d_1397_v0_})) | (d_1446_v41_)))))
            lhs181_ = d_1456_v49_
            lhs182_ = default__.safeIndex(362, (d_1456_v49_).length(0))
            lhs183_ = globalState
            lhs184_ = globalState
            lhs181_[lhs182_] = rhs207_
            lhs183_.f11 = rhs208_
            lhs184_.f11 = rhs209_
            d_1461_v54_: D2
            d_1461_v54_ = D2_DC4((self).f26)
            d_1462_v55_: _dafny.MultiSet
            d_1462_v55_ = _dafny.MultiSet([d_1461_v54_, d_1461_v54_])
            d_1463_v56_: D3
            d_1463_v56_ = D3_DC8((self).f26, 634, (d_1462_v55_).cardinality)
            d_1464_v57_: _dafny.Seq
            d_1464_v57_ = _dafny.SeqWithoutIsStrInference([d_1441_v36_, d_1441_v36_])
            d_1465_v58_: _dafny.Map
            d_1465_v58_ = _dafny.Map({(self).f25: d_1464_v57_})
            (globalState).f11 = default__.safeDivisionInt((d_1463_v56_).cf18, ((d_1440_v35_).fm5((d_1465_v58_).set(False, _dafny.SeqWithoutIsStrInference([d_1441_v36_ for d_1466_i3_ in range(default__.abs(697))])), (self).f26, (self).f26, globalState)) * (d_1397_v0_))
        elif True:
            d_1467_v59_: _dafny.Array
            nw231_ = _dafny.Array(int(0), 16)
            d_1467_v59_ = nw231_
            d_1468_v60_: D1
            d_1468_v60_ = D1_DC2(d_1467_v59_, self.f27, d_1397_v0_, (self).f25)
            d_1469_v61_: _dafny.MultiSet
            d_1469_v61_ = _dafny.MultiSet([d_1397_v0_])
            d_1470_v62_: _dafny.Seq
            d_1470_v62_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nuswicdwa"))
            d_1471_v63_: _dafny.Map
            d_1471_v63_ = _dafny.Map({(self).f25: len(d_1470_v62_)})
            source26_ = default__.fm58(D1_DC1(default__.fm38((self).f26, (d_1468_v60_).cf6, (d_1469_v61_).set(len(_dafny.SeqWithoutIsStrInference([545 for d_1472_i4_ in range(default__.abs(480))])), default__.abs(d_1397_v0_)), d_1471_v63_, globalState)), d_1470_v62_, (self).f26, globalState)
            if source26_.is_DC2:
                d_1473___mcc_h3_ = source26_.cf4
                d_1474___mcc_h4_ = source26_.cf5
                d_1475___mcc_h5_ = source26_.cf6
                d_1476___mcc_h6_ = source26_.cf7
                d_1477_cf7_ = d_1476___mcc_h6_
                d_1478_cf6_ = d_1475___mcc_h5_
                d_1479_cf5_ = d_1474___mcc_h4_
                d_1480_cf4_ = d_1473___mcc_h3_
                d_1481_v64_: _dafny.Map
                d_1481_v64_ = _dafny.Map({(self).f25: d_1470_v62_})
                d_1481_v64_ = (d_1481_v64_).set(d_1477_cf7_, d_1470_v62_)
                index210_ = default__.safeIndex(541, (d_1467_v59_).length(0))
                (d_1467_v59_)[index210_] = (d_1440_v35_).fm6(d_1441_v36_, globalState)
                d_1482_v65_: _dafny.Set
                d_1482_v65_ = _dafny.Set({d_1479_cf5_, self.f27, d_1479_cf5_, self.f27, d_1479_cf5_})
                index211_ = default__.safeIndex(541, (d_1467_v59_).length(0))
                (d_1467_v59_)[index211_] = (d_1397_v0_) + (len((d_1482_v65_) - (d_1482_v65_)))
                d_1483_v66_: C3
                nw232_ = C3()
                nw232_.ctor__((d_1469_v61_).ispropersubset(d_1469_v61_), _dafny.SeqWithoutIsStrInference([d_1479_cf5_ for d_1484_i5_ in range(default__.abs(507))]), d_1477_cf7_, self.f27, d_1477_cf7_, d_1477_cf7_)
                d_1483_v66_ = nw232_
                d_1483_v66_ = d_1483_v66_
                (self).f27 = self.f27
            elif True:
                d_1485___mcc_h7_ = source26_.cf3
                d_1486_cf3_ = d_1485___mcc_h7_
                d_1487_v67_: _dafny.Array
                nw233_ = _dafny.Array(_dafny.Seq({}), 1)
                d_1487_v67_ = nw233_
                d_1488_v68_: _dafny.Seq
                d_1488_v68_ = _dafny.SeqWithoutIsStrInference([(self).f25, self.f24, self.f24, self.f24, (self).f25])
                index212_ = default__.safeIndex(581, (d_1487_v67_).length(0))
                (d_1487_v67_)[index212_] = d_1488_v68_
                index213_ = default__.safeIndex(581, (d_1487_v67_).length(0))
                (d_1487_v67_)[index213_] = d_1488_v68_
                d_1489_v69_: _dafny.MultiSet
                d_1489_v69_ = _dafny.MultiSet([self.f24, (self).f26])
                d_1490_v70_: D3
                d_1490_v70_ = D3_DC8((self).f25, (d_1489_v69_).cardinality, -662)
                d_1491_v71_: _dafny.Array
                nw234_ = _dafny.Array(None, 27)
                nw234_[int(0)] = (self).f26
                nw234_[int(1)] = self.f24
                nw234_[int(2)] = True
                nw234_[int(3)] = False
                nw234_[int(4)] = (self).f25
                nw234_[int(5)] = False
                nw234_[int(6)] = (self).f25
                nw234_[int(7)] = (self).f26
                nw234_[int(8)] = False
                nw234_[int(9)] = True
                nw234_[int(10)] = (self).f26
                nw234_[int(11)] = self.f24
                nw234_[int(12)] = (self).f25
                nw234_[int(13)] = (d_1490_v70_).cf16
                nw234_[int(14)] = (self).f26
                nw234_[int(15)] = (self).f26
                nw234_[int(16)] = self.f24
                nw234_[int(17)] = (self).f25
                nw234_[int(18)] = self.f24
                nw234_[int(19)] = self.f24
                nw234_[int(20)] = (self).f26
                nw234_[int(21)] = (self).f26
                nw234_[int(22)] = True
                nw234_[int(23)] = (self).f25
                nw234_[int(24)] = self.f24
                nw234_[int(25)] = self.f24
                nw234_[int(26)] = (self).f25
                d_1491_v71_ = nw234_
                d_1492_v72_: _dafny.Map
                d_1492_v72_ = _dafny.Map({d_1491_v71_: ((self).f26 if True else self.f24)})
                rhs210_ = d_1492_v72_
                rhs211_ = (self).f26
                lhs185_ = self
                d_1492_v72_ = rhs210_
                lhs185_.f24 = rhs211_
                index214_ = default__.safeIndex(929, (d_1467_v59_).length(0))
                (d_1467_v59_)[index214_] = d_1397_v0_
                index215_ = default__.safeIndex(929, (d_1467_v59_).length(0))
                (d_1467_v59_)[index215_] = d_1397_v0_
                index216_ = default__.safeIndex(581, (d_1487_v67_).length(0))
                (d_1487_v67_)[index216_] = (d_1488_v68_) + ((d_1487_v67_)[default__.safeIndex(581, (d_1487_v67_).length(0))])
            d_1493_v73_: _dafny.Seq
            d_1493_v73_ = _dafny.SeqWithoutIsStrInference([self.f24])
            d_1494_v74_: C3
            nw235_ = C3()
            nw235_.ctor__((self).f25, d_1470_v62_, not((self).f25), (self.f27 if (self).f26 else self.f27), default__.fm1((d_1493_v73_)[default__.safeIndex(len(_dafny.Set({self.f24, False})), len(d_1493_v73_))], (self).f25, d_1397_v0_, globalState), self.f24)
            d_1494_v74_ = nw235_
            d_1494_v74_ = d_1494_v74_
            d_1495_v75_: _dafny.MultiSet
            d_1495_v75_ = _dafny.MultiSet([self.f27])
            d_1496_v76_: _dafny.Seq
            d_1496_v76_ = _dafny.SeqWithoutIsStrInference([d_1397_v0_, d_1397_v0_])
            d_1495_v75_ = default__.fm34((d_1494_v74_).fm39(d_1496_v76_, self.f24, globalState), globalState)
            (globalState).f10 = (not(self.f24)) or ((d_1469_v61_).isdisjoint(d_1469_v61_))
            d_1497_v77_: _dafny.Seq
            d_1497_v77_ = _dafny.SeqWithoutIsStrInference([d_1496_v76_])
            d_1496_v76_ = (d_1497_v77_)[default__.safeIndex(d_1397_v0_, len(d_1497_v77_))]
        if not((self).f25):
            (globalState).f11 = 73
            d_1498_v78_: _dafny.Seq
            d_1498_v78_ = _dafny.SeqWithoutIsStrInference([d_1397_v0_])
            d_1499_v79_: _dafny.Map
            d_1499_v79_ = _dafny.Map({len(d_1498_v78_): not(self.f24)})
            d_1500_v80_: _dafny.Map
            d_1500_v80_ = _dafny.Map({d_1397_v0_: (0) - (d_1397_v0_)})
            d_1501_v81_: D2
            d_1501_v81_ = D2_DC3(d_1500_v80_)
            d_1502_v82_: _dafny.Map
            d_1502_v82_ = _dafny.Map({d_1501_v81_: self.f24})
            d_1503_v83_: _dafny.Map
            d_1503_v83_ = _dafny.Map({not(False): (self).f25})
            d_1504_v84_: _dafny.Array
            nw236_ = _dafny.Array(None, 23)
            nw236_[int(0)] = (self).f25
            nw236_[int(1)] = (self).f25
            nw236_[int(2)] = (self).f25
            nw236_[int(3)] = (self.f24 if (self).f25 else True)
            nw236_[int(4)] = False
            nw236_[int(5)] = (self).f26
            nw236_[int(6)] = (d_1397_v0_) > (d_1397_v0_)
            nw236_[int(7)] = ((self).f26 if (self).f25 else self.f24)
            nw236_[int(8)] = (d_1397_v0_) > (d_1397_v0_)
            nw236_[int(9)] = (self).f26
            nw236_[int(10)] = (d_1397_v0_) not in (d_1499_v79_)
            nw236_[int(11)] = (self).f25
            nw236_[int(12)] = ((d_1502_v82_)[d_1501_v81_] if (d_1501_v81_) in (d_1502_v82_) else True)
            nw236_[int(13)] = True
            nw236_[int(14)] = (self).f26
            nw236_[int(15)] = self.f24
            nw236_[int(16)] = self.f24
            nw236_[int(17)] = (self).f25
            nw236_[int(18)] = (self).f25
            nw236_[int(19)] = (self).f25
            nw236_[int(20)] = not(((d_1498_v78_)[default__.safeIndex(default__.fm0(d_1397_v0_, (self).f26, d_1499_v79_, d_1397_v0_, globalState), len(d_1498_v78_))]) == (len(d_1498_v78_)))
            nw236_[int(21)] = ((d_1503_v83_)[False] if (False) in (d_1503_v83_) else (self).f26)
            nw236_[int(22)] = (_dafny.Set({self.f24, (self).f25, self.f24, True})) != (d_1439_v34_)
            d_1504_v84_ = nw236_
            index217_ = default__.safeIndex(252, (d_1504_v84_).length(0))
            (d_1504_v84_)[index217_] = False
            index218_ = default__.safeIndex(252, (d_1504_v84_).length(0))
            (d_1504_v84_)[index218_] = (self).f25
            (globalState).f16 = (len(d_1439_v34_)) + ((0) - (d_1397_v0_))
            d_1505_v85_: _dafny.Map
            d_1505_v85_ = _dafny.Map({not(not((self).f26)): d_1397_v0_})
            d_1506_v86_: D5
            d_1506_v86_ = D5_DC10(_dafny.MultiSet([d_1505_v85_]))
            d_1507_v87_: _dafny.Map
            d_1507_v87_ = _dafny.Map({d_1506_v86_: (d_1504_v84_)[default__.safeIndex(252, (d_1504_v84_).length(0))]})
            d_1508_v89_: _dafny.MultiSet
            d_1508_v89_ = _dafny.MultiSet([d_1506_v86_, d_1506_v86_])
            d_1509_v90_: _dafny.MultiSet
            d_1509_v90_ = _dafny.MultiSet([d_1505_v85_, d_1505_v85_])
            def iife103_():
                coll55_ = _dafny.Map()
                compr_55_: D5
                for compr_55_ in ((d_1508_v89_).set(D5_DC10(d_1509_v90_), default__.abs(d_1397_v0_))).Elements:
                    d_1510_v88_: D5 = compr_55_
                    if (d_1510_v88_) in ((d_1508_v89_).set(D5_DC10(d_1509_v90_), default__.abs(d_1397_v0_))):
                        coll55_[d_1510_v88_] = (d_1504_v84_)[default__.safeIndex(252, (d_1504_v84_).length(0))]
                return _dafny.Map(coll55_)
            d_1507_v87_ = iife103_()
            
            if not((self).f25):
                (self).f24 = ((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xvtxs")))) != (d_1397_v0_) if (self).fm7(default__.fm45((self).f25, globalState), globalState) else (d_1504_v84_)[default__.safeIndex(252, (d_1504_v84_).length(0))])
                d_1511_v91_: _dafny.Set
                d_1511_v91_ = _dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sc")))})
                d_1512_v92_: str
                d_1513_v93_: _dafny.Seq
                out33_: str
                out34_: _dafny.Seq
                out33_, out34_ = (d_1440_v35_).m1(d_1511_v91_, globalState)
                d_1512_v92_ = out33_
                d_1513_v93_ = out34_
                d_1514_v94_: _dafny.Array
                def lambda67_(d_1515_v0_):
                    def lambda68_(d_1516_i6_):
                        return (d_1516_i6_) * (d_1515_v0_)

                    return lambda68_

                init40_ = lambda67_(d_1397_v0_)
                nw237_ = _dafny.Array(None, 27)
                for i0_40_ in range(nw237_.length(0)):
                    nw237_[i0_40_] = init40_(i0_40_)
                d_1514_v94_ = nw237_
                d_1517_v95_: _dafny.Map
                d_1517_v95_ = _dafny.Map({d_1514_v94_: d_1506_v86_})
                d_1518_v96_: _dafny.Seq
                d_1518_v96_ = _dafny.SeqWithoutIsStrInference([d_1517_v95_, d_1517_v95_, (d_1517_v95_) | (_dafny.Map({d_1514_v94_: d_1506_v86_}))])
                (globalState).f11 = (0) - (len((d_1518_v96_)[default__.safeIndex((d_1397_v0_ if (self).f26 else d_1397_v0_), len(d_1518_v96_))]))
                d_1439_v34_ = d_1439_v34_
                d_1519_v97_: _dafny.Array
                def lambda69_(d_1520_v85_, d_1521_v78_, d_1522_v0_):
                    def lambda70_(d_1523_i7_):
                        return _dafny.Map({D5_DC10(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_1520_v85_ for d_1524_i8_ in range(default__.abs(386))]))): (d_1521_v78_)[default__.safeIndex(d_1522_v0_, len(d_1521_v78_))]})

                    return lambda70_

                init41_ = lambda69_(d_1505_v85_, d_1498_v78_, d_1397_v0_)
                nw238_ = _dafny.Array(None, 23)
                for i0_41_ in range(nw238_.length(0)):
                    nw238_[i0_41_] = init41_(i0_41_)
                d_1519_v97_ = nw238_
                d_1525_v98_: _dafny.Map
                d_1525_v98_ = _dafny.Map({d_1506_v86_: d_1397_v0_})
                index219_ = default__.safeIndex(318, (d_1519_v97_).length(0))
                (d_1519_v97_)[index219_] = d_1525_v98_
                index220_ = default__.safeIndex(318, (d_1519_v97_).length(0))
                (d_1519_v97_)[index220_] = d_1525_v98_
            elif True:
                d_1526_v99_: _dafny.Seq
                d_1526_v99_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vrgln"))
                d_1526_v99_ = d_1526_v99_
                d_1527_v100_: _dafny.Array
                nw239_ = _dafny.Array(_dafny.Array(None, 0), 14)
                d_1527_v100_ = nw239_
                d_1528_v101_: _dafny.Array
                d_1528_v101_ = d_1527_v100_
                d_1528_v101_ = d_1528_v101_
                (globalState).f22 = d_1397_v0_
                d_1529_v102_: D2
                d_1529_v102_ = D2_DC4(self.f24)
                d_1530_v103_: D2
                d_1530_v103_ = D2_DC6(d_1529_v102_)
                d_1531_v104_: D2
                d_1531_v104_ = D2_DC6(D2_DC6(d_1530_v103_))
                r1 = (d_1504_v84_ if default__.fm1((D5_DC11(d_1531_v104_, True)).cf22, (self).f26, d_1397_v0_, globalState) else d_1504_v84_)
                (globalState).f10 = ((d_1498_v78_).set(default__.safeIndex((0) - (d_1397_v0_), len(d_1498_v78_)), d_1397_v0_)) == (d_1498_v78_)
        elif True:
            d_1532_v105_: _dafny.Seq
            d_1532_v105_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xludmu"))
            d_1532_v105_ = d_1532_v105_
            d_1533_v106_: D2
            d_1533_v106_ = D2_DC4((d_1397_v0_) < ((0) - (d_1397_v0_)))
            source27_ = d_1533_v106_
            if source27_.is_DC4:
                d_1534___mcc_h8_ = source27_.cf9
                d_1535_cf9_ = d_1534___mcc_h8_
                d_1536_v107_: _dafny.Map
                d_1536_v107_ = _dafny.Map({self.f24: d_1397_v0_})
                d_1537_v108_: _dafny.MultiSet
                d_1537_v108_ = _dafny.MultiSet([not((self).fm7(d_1536_v107_, globalState)), not(not(self.f24))])
                d_1538_v109_: _dafny.Map
                d_1538_v109_ = _dafny.Map({d_1397_v0_: d_1535_cf9_})
                d_1539_v111_: _dafny.Seq
                def iife104_():
                    coll56_ = _dafny.Map()
                    compr_56_: int
                    for compr_56_ in (default__.fm53(self.f27, default__.fm47(globalState), globalState)).Elements:
                        d_1540_v110_: int = compr_56_
                        if (d_1540_v110_) in (default__.fm53(self.f27, default__.fm47(globalState), globalState)):
                            coll56_[(d_1540_v110_) - (len(d_1532_v105_))] = not((self).f26)
                    return _dafny.Map(coll56_)
                d_1539_v111_ = _dafny.SeqWithoutIsStrInference([d_1538_v109_, iife104_()
                , d_1538_v109_])
                d_1541_v112_: _dafny.Array
                nw240_ = _dafny.Array(None, 1)
                nw240_[int(0)] = d_1532_v105_
                d_1541_v112_ = nw240_
                index221_ = default__.safeIndex(66, (d_1541_v112_).length(0))
                (d_1541_v112_)[index221_] = d_1532_v105_
                index222_ = default__.safeIndex(66, (d_1541_v112_).length(0))
                rhs212_ = d_1537_v108_
                rhs213_ = (self).f26
                rhs214_ = (d_1539_v111_) + (d_1539_v111_)
                rhs215_ = (d_1532_v105_) + ((d_1532_v105_) + (d_1532_v105_))
                rhs216_ = d_1397_v0_
                lhs186_ = globalState
                lhs187_ = d_1541_v112_
                lhs188_ = default__.safeIndex(66, (d_1541_v112_).length(0))
                d_1537_v108_ = rhs212_
                lhs186_.f10 = rhs213_
                d_1539_v111_ = rhs214_
                lhs187_[lhs188_] = rhs215_
                d_1397_v0_ = rhs216_
                d_1542_v114_: _dafny.Array
                def lambda71_(d_1543_v0_, d_1544_v112_, d_1545_v105_):
                    def lambda72_(d_1546_i9_):
                        def iife105_():
                            coll57_ = _dafny.Map()
                            compr_57_: int
                            for compr_57_ in _dafny.IntegerRange(920, 361):
                                d_1547_v113_: int = compr_57_
                                if ((920) <= (d_1547_v113_)) and ((d_1547_v113_) < (361)):
                                    coll57_[(d_1547_v113_) - (d_1543_v0_)] = d_1545_v105_
                            return _dafny.Map(coll57_)
                        return (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_1543_v0_])): (d_1544_v112_)[default__.safeIndex(66, (d_1544_v112_).length(0))]})) | (iife105_()
                        )

                    return lambda72_

                init42_ = lambda71_(d_1397_v0_, d_1541_v112_, d_1532_v105_)
                nw241_ = _dafny.Array(None, 12)
                for i0_42_ in range(nw241_.length(0)):
                    nw241_[i0_42_] = init42_(i0_42_)
                d_1542_v114_ = nw241_
                d_1548_v115_: _dafny.Seq
                d_1548_v115_ = _dafny.SeqWithoutIsStrInference([d_1532_v105_])
                d_1549_v116_: _dafny.MultiSet
                d_1549_v116_ = _dafny.MultiSet([d_1397_v0_, d_1397_v0_])
                d_1550_v117_: _dafny.Map
                d_1550_v117_ = _dafny.Map({293: (d_1548_v115_)[default__.safeIndex((d_1549_v116_).cardinality, len(d_1548_v115_))]})
                index223_ = default__.safeIndex(850, (d_1542_v114_).length(0))
                (d_1542_v114_)[index223_] = (d_1550_v117_) | (d_1550_v117_)
                index224_ = default__.safeIndex(850, (d_1542_v114_).length(0))
                (d_1542_v114_)[index224_] = d_1550_v117_
                d_1551_v118_: D8
                d_1551_v118_ = D8_DC19((self).f25, d_1397_v0_)
                (globalState).f10 = (d_1551_v118_).cf37
                (globalState).f11 = d_1397_v0_
            elif source27_.is_DC5:
                d_1552___mcc_h9_ = source27_.cf10
                d_1553___mcc_h10_ = source27_.cf11
                d_1554___mcc_h11_ = source27_.cf12
                d_1555___mcc_h12_ = source27_.cf13
                d_1556_cf13_ = d_1555___mcc_h12_
                d_1557_cf12_ = d_1554___mcc_h11_
                d_1558_cf11_ = d_1553___mcc_h10_
                d_1559_cf10_ = d_1552___mcc_h9_
                d_1560_v119_: _dafny.Map
                d_1560_v119_ = _dafny.Map({d_1558_cf11_: False})
                d_1561_v120_: D7
                d_1561_v120_ = D7_DC16(d_1558_cf11_, d_1556_cf13_, d_1558_cf11_, (self).f26, (self).f26)
                d_1562_v121_: _dafny.Map
                d_1562_v121_ = _dafny.Map({self.f27: self.f27})
                d_1563_v122_: _dafny.Map
                d_1563_v122_ = _dafny.Map({d_1557_cf12_: False})
                d_1564_v123_: _dafny.Map
                d_1564_v123_ = _dafny.Map({len(d_1439_v34_): d_1556_cf13_})
                d_1565_v124_: _dafny.Array
                nw242_ = _dafny.Array(None, 26)
                nw242_[int(0)] = 297
                nw242_[int(1)] = len(d_1532_v105_)
                nw242_[int(2)] = d_1557_cf12_
                nw242_[int(3)] = d_1397_v0_
                nw242_[int(4)] = d_1556_cf13_
                nw242_[int(5)] = d_1557_cf12_
                nw242_[int(6)] = d_1556_cf13_
                nw242_[int(7)] = 194
                nw242_[int(8)] = d_1557_cf12_
                nw242_[int(9)] = d_1556_cf13_
                nw242_[int(10)] = d_1397_v0_
                nw242_[int(11)] = len(d_1560_v119_)
                nw242_[int(12)] = d_1556_cf13_
                nw242_[int(13)] = (0) - (len(d_1532_v105_))
                nw242_[int(14)] = (d_1561_v120_).cf31
                nw242_[int(15)] = d_1556_cf13_
                nw242_[int(16)] = d_1557_cf12_
                nw242_[int(17)] = d_1557_cf12_
                nw242_[int(18)] = len(d_1562_v121_)
                nw242_[int(19)] = len(d_1563_v122_)
                nw242_[int(20)] = d_1397_v0_
                nw242_[int(21)] = d_1556_cf13_
                nw242_[int(22)] = d_1556_cf13_
                nw242_[int(23)] = len(d_1532_v105_)
                nw242_[int(24)] = len(d_1564_v123_)
                nw242_[int(25)] = d_1556_cf13_
                d_1565_v124_ = nw242_
                d_1566_v125_: D1
                d_1566_v125_ = D1_DC2(d_1565_v124_, d_1559_cf10_, d_1557_cf12_, (self).f25)
                d_1567_v126_: _dafny.Seq
                d_1567_v126_ = _dafny.SeqWithoutIsStrInference([d_1558_cf11_, (d_1566_v125_).cf7, self.f24])
                d_1568_v127_: C3
                nw243_ = C3()
                nw243_.ctor__(d_1558_cf11_, (d_1532_v105_) + (_dafny.SeqWithoutIsStrInference([d_1559_cf10_ for d_1569_i10_ in range(default__.abs(-101))])), False, self.f27, not(self.f24), (d_1567_v126_) <= (d_1567_v126_))
                d_1568_v127_ = nw243_
                d_1570_v128_: _dafny.Seq
                d_1570_v128_ = _dafny.SeqWithoutIsStrInference([d_1568_v127_, d_1568_v127_, d_1568_v127_])
                d_1568_v127_ = (d_1570_v128_)[default__.safeIndex((0) - (d_1397_v0_), len(d_1570_v128_))]
                d_1571_v129_: C1
                nw244_ = C1()
                nw244_.ctor__(self.f24)
                d_1571_v129_ = nw244_
                index225_ = default__.safeIndex(417, (d_1565_v124_).length(0))
                (d_1565_v124_)[index225_] = d_1397_v0_
                index226_ = default__.safeIndex(417, (d_1565_v124_).length(0))
                (d_1565_v124_)[index226_] = (d_1568_v127_).fm5((_dafny.Map({(self).f26: _dafny.SeqWithoutIsStrInference([d_1441_v36_])})) | (default__.fm55(d_1559_cf10_, False, (self).f25, (self).f26, globalState)), self.f24, (d_1567_v126_)[default__.safeIndex(d_1556_cf13_, len(d_1567_v126_))], globalState)
                index227_ = default__.safeIndex(417, (d_1565_v124_).length(0))
                rhs217_ = (len(((d_1568_v127_).f33) + (_dafny.SeqWithoutIsStrInference([d_1559_cf10_ for d_1572_i11_ in range(default__.abs(-837))])))) - ((d_1565_v124_)[default__.safeIndex(417, (d_1565_v124_).length(0))])
                rhs218_ = default__.safeModuloInt((d_1565_v124_)[default__.safeIndex(417, (d_1565_v124_).length(0))], d_1556_cf13_)
                lhs189_ = d_1565_v124_
                lhs190_ = default__.safeIndex(417, (d_1565_v124_).length(0))
                d_1557_cf12_ = rhs217_
                lhs189_[lhs190_] = rhs218_
            elif source27_.is_DC3:
                d_1573___mcc_h13_ = source27_.cf8
                d_1574_cf8_ = d_1573___mcc_h13_
                d_1575_v130_: C2
                nw245_ = C2()
                nw245_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "egvtbhsk")), (d_1532_v105_) + ((_dafny.SeqWithoutIsStrInference([self.f27 for d_1576_i12_ in range(default__.abs(146))])).set(default__.safeIndex(d_1397_v0_, len(_dafny.SeqWithoutIsStrInference([self.f27 for d_1577_i12_ in range(default__.abs(146))]))), self.f27)))
                d_1575_v130_ = nw245_
                d_1578_v132_: D2
                def iife106_():
                    coll58_ = _dafny.Map()
                    compr_58_: int
                    for compr_58_ in _dafny.IntegerRange(768, 262):
                        d_1579_v131_: int = compr_58_
                        if ((768) <= (d_1579_v131_)) and ((d_1579_v131_) < (262)):
                            coll58_[(d_1579_v131_) * (d_1397_v0_)] = d_1397_v0_
                    return _dafny.Map(coll58_)
                d_1578_v132_ = D2_DC3(iife106_()
)
                d_1580_v133_: D2
                d_1580_v133_ = D2_DC6(d_1578_v132_)
                d_1581_v134_: _dafny.Map
                d_1581_v134_ = _dafny.Map({(len((d_1575_v130_).f30) if self.f24 else len(d_1574_cf8_)): d_1580_v133_})
                d_1581_v134_ = (d_1581_v134_).set((d_1397_v0_) - (d_1397_v0_), d_1580_v133_)
                d_1582_v136_: _dafny.Set
                d_1582_v136_ = _dafny.Set({self.f27})
                d_1583_v137_: _dafny.Seq
                d_1583_v137_ = _dafny.SeqWithoutIsStrInference([(self).f26])
                d_1584_v138_: _dafny.Seq
                def iife107_(_pat_let24_0):
                    def iife108_(d_1585_dt__update__tmp_h0_):
                        def iife109_(_pat_let25_0):
                            def iife110_(d_1586_dt__update_hcf1_h0_):
                                def iife111_(_pat_let26_0):
                                    def iife112_(d_1587_dt__update_hcf0_h0_):
                                        return D0_DC0(d_1587_dt__update_hcf0_h0_, d_1586_dt__update_hcf1_h0_, (d_1585_dt__update__tmp_h0_).cf2)
                                    return iife112_(_pat_let26_0)
                                return iife111_((self).f26)
                            return iife110_(_pat_let25_0)
                        return iife109_(118)
                    return iife108_(_pat_let24_0)
                d_1584_v138_ = _dafny.SeqWithoutIsStrInference([iife107_(d_1441_v36_)])
                d_1588_v139_: _dafny.Seq
                d_1588_v139_ = _dafny.SeqWithoutIsStrInference([d_1397_v0_, d_1397_v0_])
                d_1589_v140_: _dafny.MultiSet
                d_1589_v140_ = _dafny.MultiSet([(self).f26, self.f24])
                d_1590_v141_: _dafny.Array
                nw246_ = _dafny.Array(None, 28)
                nw246_[int(0)] = (self).f26
                nw246_[int(1)] = self.f24
                nw246_[int(2)] = not (self.f24) or ((self).f26)
                def iife113_():
                    coll59_ = _dafny.Set()
                    compr_59_: str
                    for compr_59_ in (d_1532_v105_).Elements:
                        d_1591_v135_: str = compr_59_
                        if (d_1591_v135_) in (d_1532_v105_):
                            coll59_ = coll59_.union(_dafny.Set([d_1591_v135_]))
                    return _dafny.Set(coll59_)
                nw246_[int(3)] = (d_1582_v136_).ispropersubset(iife113_()
                )
                nw246_[int(4)] = (self).f26
                nw246_[int(5)] = (self).f26
                nw246_[int(6)] = (self).f26
                nw246_[int(7)] = self.f24
                nw246_[int(8)] = (self).f25
                nw246_[int(9)] = self.f24
                nw246_[int(10)] = (self).f25
                nw246_[int(11)] = (d_1583_v137_)[default__.safeIndex(d_1397_v0_, len(d_1583_v137_))]
                nw246_[int(12)] = (self).f25
                nw246_[int(13)] = (d_1397_v0_) < (len(default__.fm59((self).fm5(_dafny.Map({(self).f25: d_1584_v138_}), self.f24, self.f24, globalState), (self).f25, self.f24, globalState)))
                nw246_[int(14)] = (self).f26
                nw246_[int(15)] = (d_1583_v137_)[default__.safeIndex((d_1588_v139_)[default__.safeIndex(661, len(d_1588_v139_))], len(d_1583_v137_))]
                nw246_[int(16)] = not(self.f24)
                nw246_[int(17)] = (len(d_1582_v136_)) in (_dafny.SeqWithoutIsStrInference([d_1397_v0_]))
                nw246_[int(18)] = (d_1575_v130_).fm19(self.f27, globalState)
                nw246_[int(19)] = (self).f25
                nw246_[int(20)] = ((self).f25) or (not(self.f24))
                nw246_[int(21)] = (self).f25
                nw246_[int(22)] = (default__.fm51(self.f27, globalState)).issubset(d_1589_v140_)
                nw246_[int(23)] = (d_1588_v139_) < (d_1588_v139_)
                nw246_[int(24)] = (self).f26
                nw246_[int(25)] = self.f24
                nw246_[int(26)] = (d_1397_v0_) not in (d_1588_v139_)
                nw246_[int(27)] = (self).f25
                d_1590_v141_ = nw246_
                d_1592_v142_: _dafny.MultiSet
                d_1592_v142_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([d_1397_v0_, d_1397_v0_])), d_1397_v0_, len(d_1575_v130_.f29)])
                index228_ = default__.safeIndex(2, (d_1590_v141_).length(0))
                (d_1590_v141_)[index228_] = not((d_1592_v142_).isdisjoint(d_1592_v142_))
                index229_ = default__.safeIndex(2, (d_1590_v141_).length(0))
                (d_1590_v141_)[index229_] = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iw"))) + (d_1532_v105_)) == ((d_1575_v130_).f30)
                d_1593_v143_: C4
                nw247_ = C4()
                nw247_.ctor__(self.f24, (self).f26)
                d_1593_v143_ = nw247_
            elif True:
                d_1594___mcc_h14_ = source27_.cf14
                d_1595_cf14_ = d_1594___mcc_h14_
                d_1596_v144_: C1
                nw248_ = C1()
                nw248_.ctor__(not((self).f25))
                d_1596_v144_ = nw248_
                d_1597_v145_: _dafny.Seq
                d_1597_v145_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([False]))])
                d_1598_v146_: _dafny.Array
                nw249_ = _dafny.Array(None, 3)
                nw249_[int(0)] = (0) - (d_1397_v0_)
                nw249_[int(1)] = len(d_1597_v145_)
                nw249_[int(2)] = d_1397_v0_
                d_1598_v146_ = nw249_
                d_1599_v147_: D1
                d_1599_v147_ = D1_DC2(d_1598_v146_, self.f27, d_1397_v0_, (self).f25)
                d_1600_v148_: _dafny.Seq
                d_1600_v148_ = _dafny.SeqWithoutIsStrInference([(self).f26])
                d_1601_v149_: _dafny.Seq
                d_1601_v149_ = _dafny.SeqWithoutIsStrInference([d_1598_v146_, d_1598_v146_])
                d_1602_v150_: D2
                d_1602_v150_ = D2_DC3(_dafny.Map({d_1397_v0_: d_1397_v0_}))
                d_1603_v151_: D2
                d_1603_v151_ = D2_DC6(d_1602_v150_)
                d_1604_v152_: D5
                d_1604_v152_ = D5_DC11(d_1603_v151_, (self).f26)
                d_1605_v153_: D6
                d_1605_v153_ = D6_DC12(_dafny.SeqWithoutIsStrInference([d_1397_v0_]))
                d_1606_v154_: _dafny.Seq
                d_1606_v154_ = _dafny.SeqWithoutIsStrInference([d_1605_v153_])
                d_1607_v155_: _dafny.Set
                d_1607_v155_ = _dafny.Set({d_1397_v0_})
                d_1608_v156_: _dafny.MultiSet
                d_1608_v156_ = _dafny.MultiSet([d_1440_v35_, d_1440_v35_])
                d_1609_v157_: _dafny.Array
                nw250_ = _dafny.Array(None, 22)
                nw250_[int(0)] = (d_1599_v147_).cf7
                nw250_[int(1)] = not((d_1600_v148_)[default__.safeIndex(len(d_1597_v145_), len(d_1600_v148_))])
                nw250_[int(2)] = self.f24
                nw250_[int(3)] = (_dafny.SeqWithoutIsStrInference([d_1598_v146_, d_1598_v146_])) != (d_1601_v149_)
                nw250_[int(4)] = (self).f26
                nw250_[int(5)] = (self).f25
                nw250_[int(6)] = ((self).f26 if (self).f26 else (d_1604_v152_).cf22)
                nw250_[int(7)] = (self).f25
                nw250_[int(8)] = (d_1600_v148_) <= (d_1600_v148_)
                nw250_[int(9)] = (d_1532_v105_) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gtqtq")))
                nw250_[int(10)] = self.f24
                nw250_[int(11)] = (self).f25
                nw250_[int(12)] = not((d_1606_v154_) < (_dafny.SeqWithoutIsStrInference([d_1605_v153_ for d_1610_i13_ in range(default__.abs(-342))])))
                nw250_[int(13)] = (self).f26
                nw250_[int(14)] = (self).f26
                nw250_[int(15)] = ((self).f26) or (self.f24)
                nw250_[int(16)] = (self).f25
                nw250_[int(17)] = (len(d_1532_v105_)) <= (len(d_1607_v155_))
                nw250_[int(18)] = (d_1608_v156_).issubset(d_1608_v156_)
                nw250_[int(19)] = (self.f24) == (self.f24)
                nw250_[int(20)] = (False if (self).f25 else (self).f25)
                nw250_[int(21)] = (self).f25
                d_1609_v157_ = nw250_
                index230_ = default__.safeIndex(418, (d_1609_v157_).length(0))
                (d_1609_v157_)[index230_] = (self).f25
                index231_ = default__.safeIndex(418, (d_1609_v157_).length(0))
                (d_1609_v157_)[index231_] = not((self).f25)
                d_1611_v158_: _dafny.Map
                d_1611_v158_ = _dafny.Map({368: default__.safeDivisionInt(d_1397_v0_, d_1397_v0_)})
                d_1611_v158_ = (d_1611_v158_).set(d_1397_v0_, (d_1397_v0_ if self.f24 else 462))
                (globalState).f11 = (default__.safeDivisionInt(d_1397_v0_, d_1397_v0_)) - (d_1397_v0_)
            (self).f27 = (self.f27 if (self).f25 else self.f27)
            d_1612_v159_: D8
            d_1612_v159_ = D8_DC19(default__.fm1(False, (self).f26, d_1397_v0_, globalState), d_1397_v0_)
            d_1613_v160_: _dafny.Set
            def iife114_(_pat_let27_0):
                def iife115_(d_1614_dt__update__tmp_h1_):
                    def iife116_(_pat_let28_0):
                        def iife117_(d_1615_dt__update_hcf37_h0_):
                            return D8_DC19(d_1615_dt__update_hcf37_h0_, (d_1614_dt__update__tmp_h1_).cf38)
                        return iife117_(_pat_let28_0)
                    return iife116_(self.f24)
                return iife115_(_pat_let27_0)
            d_1613_v160_ = _dafny.Set({iife114_(d_1612_v159_), d_1612_v159_, D8_DC19(self.f24, d_1397_v0_), d_1612_v159_})
            d_1613_v160_ = d_1613_v160_
            rhs219_ = True
            rhs220_ = not((self).f25)
            lhs191_ = self
            lhs192_ = globalState
            lhs191_.f24 = rhs219_
            lhs192_.f10 = rhs220_
        d_1616_v161_: _dafny.Array
        nw251_ = _dafny.Array(_dafny.CodePoint('D'), 24)
        d_1616_v161_ = nw251_
        d_1617_v162_: _dafny.MultiSet
        d_1617_v162_ = _dafny.MultiSet([(self).f25])
        d_1618_v163_: _dafny.Seq
        d_1618_v163_ = _dafny.SeqWithoutIsStrInference([d_1617_v162_])
        d_1619_v164_: _dafny.Map
        d_1619_v164_ = _dafny.Map({d_1616_v161_: (((d_1618_v163_)[default__.safeIndex(392, len(d_1618_v163_))]).cardinality) - (d_1397_v0_)})
        d_1619_v164_ = (d_1619_v164_).set(d_1616_v161_, d_1397_v0_)
        d_1620_v165_: _dafny.Map
        d_1620_v165_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wbv"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yubckxhrt"))): (not((self).f25) if False else (D3_DC8(self.f24, d_1397_v0_, d_1397_v0_)).cf16)})
        rhs221_ = ((self).f25 if (self).f26 else self.f24)
        rhs222_ = ((d_1620_v165_)[(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "puuwhjb"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ocvfbl")))] if ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "puuwhjb"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ocvfbl")))) in (d_1620_v165_) else self.f24)
        rhs223_ = (d_1440_v35_).fm6(D0_DC0(self.f24, d_1397_v0_, d_1397_v0_), globalState)
        rhs224_ = (self).f25
        lhs193_ = self
        lhs194_ = globalState
        lhs193_.f24 = rhs221_
        r0 = rhs222_
        d_1397_v0_ = rhs223_
        lhs194_.f10 = rhs224_
        r0 = (d_1397_v0_) == ((d_1397_v0_) * (-424))
        d_1621_v166_: _dafny.Array
        def lambda73_(d_1622_i14_):
            return (self).f25

        init43_ = lambda73_
        nw252_ = _dafny.Array(None, 25)
        for i0_43_ in range(nw252_.length(0)):
            nw252_[i0_43_] = init43_(i0_43_)
        d_1621_v166_ = nw252_
        r1 = d_1621_v166_
        return r0, r1

    def m1(self, p0, globalState):
        r0: str = _dafny.CodePoint('D')
        r1: _dafny.Seq = _dafny.Seq({})
        d_1623_v0_: _dafny.Map
        d_1623_v0_ = _dafny.Map({self.f24: self.f24})
        d_1624_v1_: _dafny.Set
        d_1624_v1_ = _dafny.Set({d_1623_v0_, d_1623_v0_, _dafny.Map({(self).f26: (self).f25}), (d_1623_v0_).set(self.f24, (self).f26), _dafny.Map({(self).f26: (self).f25})})
        if (d_1624_v1_).issubset(d_1624_v1_):
            (globalState).f10 = not(False)
            d_1625_v2_: int
            d_1625_v2_ = 279
            (globalState).f16 = d_1625_v2_
            d_1626_v3_: _dafny.Array
            nw253_ = _dafny.Array(int(0), 19)
            d_1626_v3_ = nw253_
            d_1627_v4_: _dafny.Seq
            d_1627_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rfk"))
            d_1628_v5_: _dafny.MultiSet
            d_1628_v5_ = _dafny.MultiSet([d_1627_v4_])
            d_1629_v6_: C2
            nw254_ = C2()
            nw254_.ctor__(d_1627_v4_, d_1627_v4_)
            d_1629_v6_ = nw254_
            d_1630_v7_: _dafny.Seq
            d_1630_v7_ = _dafny.SeqWithoutIsStrInference([d_1629_v6_])
            d_1631_v8_: _dafny.Map
            d_1631_v8_ = _dafny.Map({d_1630_v7_: (self).f25})
            index232_ = default__.safeIndex(637, (d_1626_v3_).length(0))
            (d_1626_v3_)[index232_] = (((d_1628_v5_)[_dafny.SeqWithoutIsStrInference([self.f27 for d_1632_i0_ in range(default__.abs(373))])] if (_dafny.SeqWithoutIsStrInference([self.f27 for d_1633_i0_ in range(default__.abs(373))])) in (d_1628_v5_) else d_1625_v2_) if self.f24 else len(d_1631_v8_))
            index233_ = default__.safeIndex(637, (d_1626_v3_).length(0))
            (d_1626_v3_)[index233_] = d_1625_v2_
            d_1634_v9_: _dafny.Seq
            d_1634_v9_ = _dafny.SeqWithoutIsStrInference([(d_1629_v6_).f30])
            (globalState).f11 = (0) - (((0) - (len((d_1634_v9_) + ((d_1634_v9_).set(default__.safeIndex((d_1626_v3_)[default__.safeIndex(637, (d_1626_v3_).length(0))], len(d_1634_v9_)), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hujqf"))))))) * ((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "begj"))))))
            (globalState).f10 = self.f24
        elif True:
            d_1635_v10_: _dafny.Array
            nw255_ = _dafny.Array(int(0), 6)
            d_1635_v10_ = nw255_
            d_1636_v11_: int
            d_1636_v11_ = 768
            index234_ = default__.safeIndex(147, (d_1635_v10_).length(0))
            (d_1635_v10_)[index234_] = d_1636_v11_
            index235_ = default__.safeIndex(147, (d_1635_v10_).length(0))
            (d_1635_v10_)[index235_] = d_1636_v11_
            (globalState).f10 = (self).f26
            d_1637_v12_: _dafny.Map
            d_1637_v12_ = _dafny.Map({(d_1636_v11_) - (len(d_1623_v0_)): (d_1635_v10_)[default__.safeIndex(147, (d_1635_v10_).length(0))]})
            d_1638_v13_: _dafny.Seq
            d_1638_v13_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "neiyfvffm"))
            d_1639_v14_: C0
            nw256_ = C0()
            nw256_.ctor__(d_1638_v13_)
            d_1639_v14_ = nw256_
            d_1640_v15_: _dafny.Array
            nw257_ = _dafny.Array(None, 24)
            nw257_[int(0)] = self.f27
            nw257_[int(1)] = self.f27
            nw257_[int(2)] = self.f27
            nw257_[int(3)] = self.f27
            nw257_[int(4)] = self.f27
            nw257_[int(5)] = _dafny.CodePoint('h')
            nw257_[int(6)] = ((d_1639_v14_).f31)[default__.safeIndex(d_1636_v11_, len((d_1639_v14_).f31))]
            nw257_[int(7)] = _dafny.CodePoint('v')
            nw257_[int(8)] = self.f27
            nw257_[int(9)] = self.f27
            nw257_[int(10)] = self.f27
            nw257_[int(11)] = self.f27
            nw257_[int(12)] = self.f27
            nw257_[int(13)] = self.f27
            nw257_[int(14)] = self.f27
            nw257_[int(15)] = self.f27
            nw257_[int(16)] = _dafny.CodePoint('v')
            nw257_[int(17)] = _dafny.CodePoint('c')
            nw257_[int(18)] = self.f27
            nw257_[int(19)] = self.f27
            nw257_[int(20)] = self.f27
            nw257_[int(21)] = self.f27
            nw257_[int(22)] = self.f27
            nw257_[int(23)] = self.f27
            d_1640_v15_ = nw257_
            d_1641_v16_: _dafny.Seq
            d_1641_v16_ = _dafny.SeqWithoutIsStrInference([d_1640_v15_])
            d_1642_v17_: _dafny.MultiSet
            d_1642_v17_ = _dafny.MultiSet([len(d_1641_v16_)])
            d_1643_v18_: _dafny.Map
            d_1643_v18_ = _dafny.Map({d_1639_v14_: (d_1642_v17_).cardinality})
            d_1637_v12_ = (d_1637_v12_).set((d_1635_v10_)[default__.safeIndex(147, (d_1635_v10_).length(0))], len((d_1643_v18_) | (_dafny.Map({d_1639_v14_: (d_1635_v10_)[default__.safeIndex(147, (d_1635_v10_).length(0))]}))))
            d_1644_v19_: _dafny.Array
            def lambda74_(d_1645_i1_):
                return _dafny.Map({False: self.f24})

            init44_ = lambda74_
            nw258_ = _dafny.Array(None, 23)
            for i0_44_ in range(nw258_.length(0)):
                nw258_[i0_44_] = init44_(i0_44_)
            d_1644_v19_ = nw258_
            d_1646_v20_: _dafny.Map
            d_1646_v20_ = _dafny.Map({(self).f25: (0) - (d_1636_v11_)})
            d_1647_v21_: _dafny.MultiSet
            d_1647_v21_ = _dafny.MultiSet([(d_1639_v14_).f31])
            d_1648_v22_: D2
            d_1648_v22_ = D2_DC5(self.f27, ((self).fm7(d_1646_v20_, globalState) if not(self.f24) else (self).f26), ((d_1647_v21_)[d_1638_v13_] if (d_1638_v13_) in (d_1647_v21_) else (d_1635_v10_)[default__.safeIndex(147, (d_1635_v10_).length(0))]), (d_1635_v10_)[default__.safeIndex(147, (d_1635_v10_).length(0))])
            d_1649_v23_: _dafny.Array
            def lambda75_(d_1650_v14_):
                def lambda76_(d_1651_i2_):
                    return (d_1650_v14_).f31

                return lambda76_

            init45_ = lambda75_(d_1639_v14_)
            nw259_ = _dafny.Array(None, 13)
            for i0_45_ in range(nw259_.length(0)):
                nw259_[i0_45_] = init45_(i0_45_)
            d_1649_v23_ = nw259_
            d_1652_v24_: _dafny.Map
            d_1652_v24_ = _dafny.Map({(d_1635_v10_)[default__.safeIndex(147, (d_1635_v10_).length(0))]: d_1649_v23_})
            d_1653_v25_: _dafny.Array
            d_1653_v25_ = ((d_1652_v24_)[(d_1635_v10_)[default__.safeIndex(147, (d_1635_v10_).length(0))]] if ((d_1635_v10_)[default__.safeIndex(147, (d_1635_v10_).length(0))]) in (d_1652_v24_) else d_1649_v23_)
            rhs225_ = d_1644_v19_
            rhs226_ = (d_1648_v22_ if (d_1636_v11_) > (len(d_1638_v13_)) else d_1648_v22_)
            rhs227_ = d_1636_v11_
            rhs228_ = d_1649_v23_
            lhs195_ = globalState
            d_1644_v19_ = rhs225_
            d_1648_v22_ = rhs226_
            lhs195_.f22 = rhs227_
            d_1653_v25_ = rhs228_
            d_1654_v26_: _dafny.Array
            nw260_ = _dafny.Array(False, 4)
            d_1654_v26_ = nw260_
            (globalState).f2 = d_1654_v26_
        d_1655_v27_: int
        d_1655_v27_ = 363
        hi10_ = d_1655_v27_
        for d_1656_i3_ in range(default__.safeDivisionInt(d_1655_v27_, -188), hi10_):
            d_1657_v28_: _dafny.Array
            nw261_ = _dafny.Array(int(0), 5)
            d_1657_v28_ = nw261_
            d_1658_v29_: _dafny.MultiSet
            d_1658_v29_ = _dafny.MultiSet([self.f24, (self).f26])
            index236_ = default__.safeIndex(676, (d_1657_v28_).length(0))
            (d_1657_v28_)[index236_] = ((d_1658_v29_)[(self).f26] if ((self).f26) in (d_1658_v29_) else d_1655_v27_)
            d_1659_v30_: _dafny.MultiSet
            d_1659_v30_ = _dafny.MultiSet([self.f27, self.f27])
            index237_ = default__.safeIndex(676, (d_1657_v28_).length(0))
            (d_1657_v28_)[index237_] = (935) + (default__.safeDivisionInt(((d_1659_v30_)[self.f27] if (self.f27) in (d_1659_v30_) else d_1655_v27_), d_1656_i3_))
            d_1660_v31_: _dafny.Array
            def lambda77_(d_1661_i4_):
                return (self).f25

            init46_ = lambda77_
            nw262_ = _dafny.Array(None, 17)
            for i0_46_ in range(nw262_.length(0)):
                nw262_[i0_46_] = init46_(i0_46_)
            d_1660_v31_ = nw262_
            d_1662_v32_: _dafny.Seq
            d_1662_v32_ = _dafny.SeqWithoutIsStrInference([d_1660_v31_, d_1660_v31_])
            d_1663_v33_: _dafny.Seq
            d_1663_v33_ = _dafny.SeqWithoutIsStrInference([self.f24, (self).f25])
            source28_ = D6_DC13((d_1662_v32_)[default__.safeIndex(-498, len(d_1662_v32_))], (0) - (len(d_1663_v33_)), self.f24)
            if source28_.is_DC13:
                d_1664___mcc_h0_ = source28_.cf24
                d_1665___mcc_h1_ = source28_.cf25
                d_1666___mcc_h2_ = source28_.cf26
                d_1667_cf26_ = d_1666___mcc_h2_
                d_1668_cf25_ = d_1665___mcc_h1_
                d_1669_cf24_ = d_1664___mcc_h0_
                d_1655_v27_ = default__.safeModuloInt(default__.safeDivisionInt(d_1656_i3_, len(d_1662_v32_)), ((d_1657_v28_)[default__.safeIndex(676, (d_1657_v28_).length(0))]) * ((0) - (d_1655_v27_)))
                d_1670_v34_: _dafny.Seq
                d_1670_v34_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xf"))
                d_1670_v34_ = (d_1670_v34_).set(default__.safeIndex((d_1657_v28_)[default__.safeIndex(676, (d_1657_v28_).length(0))], len(d_1670_v34_)), self.f27)
                index238_ = default__.safeIndex(241, (d_1660_v31_).length(0))
                (d_1660_v31_)[index238_] = not (d_1667_cf26_) or (not(d_1667_cf26_))
                index239_ = default__.safeIndex(763, (d_1660_v31_).length(0))
                (d_1660_v31_)[index239_] = ((d_1657_v28_)[default__.safeIndex(676, (d_1657_v28_).length(0))]) in (_dafny.MultiSet([405]))
                d_1671_v35_: _dafny.Map
                d_1671_v35_ = _dafny.Map({d_1655_v27_: d_1663_v33_})
                index240_ = default__.safeIndex(241, (d_1660_v31_).length(0))
                index241_ = default__.safeIndex(763, (d_1660_v31_).length(0))
                rhs229_ = (d_1663_v33_)[default__.safeIndex(d_1668_cf25_, len(d_1663_v33_))]
                rhs230_ = (d_1656_i3_ if (self).f26 else len((_dafny.Map({818: d_1663_v33_})) | (d_1671_v35_)))
                rhs231_ = ((d_1670_v34_) + (d_1670_v34_)) == (d_1670_v34_)
                rhs232_ = self.f24
                lhs196_ = globalState
                lhs197_ = globalState
                lhs198_ = d_1660_v31_
                lhs199_ = default__.safeIndex(241, (d_1660_v31_).length(0))
                lhs200_ = d_1660_v31_
                lhs201_ = default__.safeIndex(763, (d_1660_v31_).length(0))
                lhs196_.f10 = rhs229_
                lhs197_.f11 = rhs230_
                lhs198_[lhs199_] = rhs231_
                lhs200_[lhs201_] = rhs232_
                d_1672_v36_: _dafny.Array
                nw263_ = _dafny.Array(_dafny.Map({}), 6)
                d_1672_v36_ = nw263_
                d_1673_v37_: D10
                d_1673_v37_ = D10_DC21(d_1672_v36_)
                d_1672_v36_ = (d_1673_v37_).cf40
            elif True:
                d_1674___mcc_h3_ = source28_.cf23
                d_1675_cf23_ = d_1674___mcc_h3_
                d_1676_v38_: _dafny.MultiSet
                d_1676_v38_ = _dafny.MultiSet([(d_1657_v28_)[default__.safeIndex(676, (d_1657_v28_).length(0))], (_dafny.MultiSet([(d_1657_v28_)[default__.safeIndex(676, (d_1657_v28_).length(0))]])).cardinality, (d_1657_v28_)[default__.safeIndex(676, (d_1657_v28_).length(0))]])
                d_1677_v39_: _dafny.Seq
                d_1677_v39_ = _dafny.SeqWithoutIsStrInference([self.f27])
                rhs233_ = d_1656_i3_
                rhs234_ = d_1657_v28_
                rhs235_ = default__.safeDivisionInt((0) - (default__.safeModuloInt(((d_1676_v38_)[571] if (571) in (d_1676_v38_) else len(d_1677_v39_)), (d_1657_v28_)[default__.safeIndex(676, (d_1657_v28_).length(0))])), d_1656_i3_)
                lhs202_ = globalState
                lhs203_ = globalState
                lhs204_ = globalState
                lhs202_.f16 = rhs233_
                lhs203_.f7 = rhs234_
                lhs204_.f16 = rhs235_
                d_1678_v40_: _dafny.Map
                d_1678_v40_ = _dafny.Map({self.f24: _dafny.Set({_dafny.MultiSet([True, (self).f25])})})
                d_1679_v41_: _dafny.Map
                d_1679_v41_ = _dafny.Map({(_dafny.MultiSet([d_1656_i3_])).cardinality: d_1658_v29_})
                d_1680_v42_: _dafny.Set
                d_1680_v42_ = _dafny.Set({((d_1679_v41_)[(d_1657_v28_)[default__.safeIndex(676, (d_1657_v28_).length(0))]] if ((d_1657_v28_)[default__.safeIndex(676, (d_1657_v28_).length(0))]) in (d_1679_v41_) else (_dafny.MultiSet([self.f24, (self).f26])).set((self).f26, default__.abs(-438)))})
                d_1678_v40_ = (d_1678_v40_).set(self.f24, d_1680_v42_)
                (globalState).f11 = (d_1657_v28_)[default__.safeIndex(676, (d_1657_v28_).length(0))]
                d_1677_v39_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gusxkuor"))) + (d_1677_v39_)).set(default__.safeIndex((d_1655_v27_) - (d_1656_i3_), len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gusxkuor"))) + (d_1677_v39_))), (self.f27 if (self).f25 else self.f27))
            d_1681_v43_: _dafny.Map
            d_1681_v43_ = _dafny.Map({(d_1657_v28_)[default__.safeIndex(676, (d_1657_v28_).length(0))]: (self).f25})
            d_1682_v44_: _dafny.Map
            d_1682_v44_ = _dafny.Map({True: (d_1657_v28_)[default__.safeIndex(676, (d_1657_v28_).length(0))]})
            d_1683_v45_: _dafny.MultiSet
            d_1683_v45_ = _dafny.MultiSet([(0) - ((0) - (len((d_1681_v43_ if (self).f25 else d_1681_v43_)))), default__.safeDivisionInt(d_1655_v27_, (0) - (default__.fm0((d_1657_v28_)[default__.safeIndex(676, (d_1657_v28_).length(0))], (self).f26, (d_1681_v43_).set(875, (self).fm7(d_1682_v44_, globalState)), (0) - (d_1656_i3_), globalState)))])
            d_1684_v46_: _dafny.Seq
            d_1684_v46_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fieqtkeqr"))
            d_1683_v45_ = (d_1683_v45_) - (((d_1683_v45_).set((0) - ((d_1657_v28_)[default__.safeIndex(676, (d_1657_v28_).length(0))]), default__.abs(len(d_1684_v46_)))) - (d_1683_v45_))
            d_1685_v47_: C4
            nw264_ = C4()
            nw264_.ctor__(default__.fm1((self).f25, not(self.f24), (d_1657_v28_)[default__.safeIndex(676, (d_1657_v28_).length(0))], globalState), (self).f26)
            d_1685_v47_ = nw264_
        d_1686_v48_: _dafny.MultiSet
        d_1686_v48_ = _dafny.MultiSet([d_1655_v27_, d_1655_v27_, 71])
        d_1687_v49_: _dafny.Map
        d_1687_v49_ = _dafny.Map({d_1655_v27_: default__.safeDivisionInt(d_1655_v27_, (d_1686_v48_).cardinality)})
        d_1688_v50_: _dafny.Set
        d_1688_v50_ = _dafny.Set({(self).f26, self.f24})
        d_1689_v51_: D2
        d_1689_v51_ = D2_DC5(self.f27, (self).f26, d_1655_v27_, len(d_1688_v50_))
        d_1690_v52_: D2
        d_1690_v52_ = D2_DC6(d_1689_v51_)
        d_1691_v53_: D2
        d_1691_v53_ = D2_DC6(d_1689_v51_)
        d_1687_v49_ = default__.fm50(d_1691_v53_, d_1655_v27_, (self).f26, globalState)
        d_1692_i5_: int
        d_1692_i5_ = 0
        with _dafny.label("14"):
            while not((default__.safeModuloInt(d_1655_v27_, 707)) >= (d_1655_v27_)):
                with _dafny.c_label("14"):
                    if (d_1692_i5_) >= (100):
                        raise _dafny.Break("14")
                    d_1692_i5_ = (d_1692_i5_) + (1)
                    d_1693_v54_: _dafny.Array
                    def lambda78_(d_1694_v27_):
                        def lambda79_(d_1695_i6_):
                            return (d_1695_i6_) * (d_1694_v27_)

                        return lambda79_

                    init47_ = lambda78_(d_1655_v27_)
                    nw265_ = _dafny.Array(None, 20)
                    for i0_47_ in range(nw265_.length(0)):
                        nw265_[i0_47_] = init47_(i0_47_)
                    d_1693_v54_ = nw265_
                    d_1696_v55_: _dafny.MultiSet
                    d_1696_v55_ = _dafny.MultiSet([(_dafny.MultiSet([d_1655_v27_])).set(d_1655_v27_, default__.abs(d_1655_v27_)), d_1686_v48_, d_1686_v48_, d_1686_v48_])
                    d_1697_v57_: _dafny.Seq
                    d_1697_v57_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ksrgpvm"))
                    d_1698_v58_: _dafny.MultiSet
                    d_1698_v58_ = _dafny.MultiSet([d_1697_v57_, d_1697_v57_])
                    index242_ = default__.safeIndex(174, (d_1693_v54_).length(0))
                    def iife118_():
                        coll60_ = _dafny.Map()
                        compr_60_: int
                        for compr_60_ in (d_1687_v49_).keys.Elements:
                            d_1699_v56_: int = compr_60_
                            if (d_1699_v56_) in (d_1687_v49_):
                                coll60_[(d_1699_v56_) * ((d_1686_v48_).cardinality)] = (D2_DC5(self.f27, (self).f25, d_1655_v27_, d_1655_v27_)).cf11
                        return _dafny.Map(coll60_)
                    (d_1693_v54_)[index242_] = (0) - (default__.fm0(d_1655_v27_, not(default__.fm1(not((self).f25), (self).f26, (d_1696_v55_).cardinality, globalState)), iife118_()
                    , (d_1698_v58_).cardinality, globalState))
                    index243_ = default__.safeIndex(174, (d_1693_v54_).length(0))
                    (d_1693_v54_)[index243_] = 810
                    (globalState).f16 = (309) - (len(d_1697_v57_))
                    if self.f24:
                        (self).f24 = (self).f26
                        (self).f24 = (d_1655_v27_) > (d_1655_v27_)
                        d_1700_v59_: _dafny.Seq
                        d_1700_v59_ = _dafny.SeqWithoutIsStrInference([938])
                        d_1701_v60_: _dafny.Map
                        d_1701_v60_ = _dafny.Map({len(d_1700_v59_): (self).f26})
                        d_1655_v27_ = default__.fm0((0) - ((d_1693_v54_)[default__.safeIndex(174, (d_1693_v54_).length(0))]), not ((self).f25) or (self.f24), d_1701_v60_, 896, globalState)
                        d_1702_v61_: _dafny.Seq
                        d_1702_v61_ = _dafny.SeqWithoutIsStrInference([(d_1697_v57_).set(default__.safeIndex((d_1686_v48_).cardinality, len(d_1697_v57_)), self.f27)])
                        d_1703_v62_: _dafny.Seq
                        d_1703_v62_ = _dafny.SeqWithoutIsStrInference([(self).f26, not(self.f24)])
                        d_1704_v63_: C3
                        nw266_ = C3()
                        nw266_.ctor__((d_1688_v50_).issubset(d_1688_v50_), (d_1702_v61_)[default__.safeIndex(len(d_1702_v61_), len(d_1702_v61_))], (d_1703_v62_)[default__.safeIndex(d_1655_v27_, len(d_1703_v62_))], self.f27, False, (d_1688_v50_).issubset(d_1688_v50_))
                        d_1704_v63_ = nw266_
                        d_1705_v64_: D11
                        d_1705_v64_ = D11_DC25(d_1704_v63_)
                        d_1704_v63_ = (d_1705_v64_).cf46
                        d_1697_v57_ = ((d_1704_v63_).f33) + (d_1697_v57_)
                    elif True:
                        d_1706_v65_: _dafny.Seq
                        d_1706_v65_ = _dafny.SeqWithoutIsStrInference([(self).f25])
                        r1 = d_1706_v65_
                        (globalState).f10 = False
                        d_1707_v66_: _dafny.Set
                        d_1707_v66_ = _dafny.Set({d_1655_v27_})
                        d_1707_v66_ = p0
                        d_1708_v67_: _dafny.Map
                        d_1708_v67_ = _dafny.Map({self.f27: 407})
                        d_1708_v67_ = (d_1708_v67_).set(self.f27, (d_1693_v54_)[default__.safeIndex(174, (d_1693_v54_).length(0))])
                        d_1709_v68_: C2
                        nw267_ = C2()
                        nw267_.ctor__(d_1697_v57_, d_1697_v57_)
                        d_1709_v68_ = nw267_
                    if ((0) - (default__.safeDivisionInt((d_1693_v54_)[default__.safeIndex(174, (d_1693_v54_).length(0))], d_1655_v27_))) == (d_1655_v27_):
                        (globalState).f10 = (d_1655_v27_) == (((d_1693_v54_)[default__.safeIndex(174, (d_1693_v54_).length(0))]) - ((d_1693_v54_)[default__.safeIndex(174, (d_1693_v54_).length(0))]))
                        d_1687_v49_ = (d_1687_v49_).set((d_1693_v54_)[default__.safeIndex(174, (d_1693_v54_).length(0))], 212)
                        d_1710_v69_: C3
                        nw268_ = C3()
                        nw268_.ctor__((self).f25, d_1697_v57_, False, _dafny.CodePoint('g'), self.f24, False)
                        d_1710_v69_ = nw268_
                        d_1711_v70_: D11
                        d_1711_v70_ = D11_DC25(d_1710_v69_)
                        d_1712_v71_: _dafny.Map
                        d_1712_v71_ = _dafny.Map({d_1711_v70_: (d_1693_v54_)[default__.safeIndex(174, (d_1693_v54_).length(0))]})
                        rhs236_ = ((d_1712_v71_)[D11_DC25(d_1710_v69_)] if (D11_DC25(d_1710_v69_)) in (d_1712_v71_) else d_1655_v27_)
                        rhs237_ = (d_1693_v54_)[default__.safeIndex(174, (d_1693_v54_).length(0))]
                        lhs205_ = globalState
                        lhs206_ = globalState
                        lhs205_.f11 = rhs236_
                        lhs206_.f16 = rhs237_
                        index244_ = default__.safeIndex(174, (d_1693_v54_).length(0))
                        (d_1693_v54_)[index244_] = (d_1693_v54_)[default__.safeIndex(174, (d_1693_v54_).length(0))]
                        d_1713_v72_: C3
                        nw269_ = C3()
                        nw269_.ctor__((self).f25, (d_1710_v69_).f33, ((self).f25) and ((self).f25), self.f27, self.f24, True)
                        d_1713_v72_ = nw269_
                    elif True:
                        d_1714_v73_: _dafny.Array
                        nw270_ = _dafny.Array(_dafny.Array(None, 0), 20)
                        d_1714_v73_ = nw270_
                        d_1715_v74_: _dafny.Array
                        nw271_ = _dafny.Array(None, 2)
                        nw271_[int(0)] = self.f24
                        nw271_[int(1)] = self.f24
                        d_1715_v74_ = nw271_
                        index245_ = default__.safeIndex(905, (d_1714_v73_).length(0))
                        (d_1714_v73_)[index245_] = d_1715_v74_
                        index246_ = default__.safeIndex(905, (d_1714_v73_).length(0))
                        (d_1714_v73_)[index246_] = d_1715_v74_
                        d_1716_v75_: _dafny.Array
                        def lambda80_(d_1717_i7_):
                            return _dafny.Map({(self).f26: self.f27})

                        init48_ = lambda80_
                        nw272_ = _dafny.Array(None, 20)
                        for i0_48_ in range(nw272_.length(0)):
                            nw272_[i0_48_] = init48_(i0_48_)
                        d_1716_v75_ = nw272_
                        d_1718_v76_: D10
                        d_1718_v76_ = D10_DC21(d_1716_v75_)
                        d_1719_v77_: _dafny.Map
                        d_1719_v77_ = _dafny.Map({(d_1693_v54_)[default__.safeIndex(174, (d_1693_v54_).length(0))]: (d_1714_v73_)[default__.safeIndex(905, (d_1714_v73_).length(0))]})
                        index247_ = default__.safeIndex(905, (d_1714_v73_).length(0))
                        rhs238_ = ((d_1714_v73_)[default__.safeIndex(905, (d_1714_v73_).length(0))] if not((self).f26) else ((d_1719_v77_)[d_1655_v27_] if (d_1655_v27_) in (d_1719_v77_) else d_1715_v74_))
                        rhs239_ = (self).f25
                        rhs240_ = D10_DC21(d_1716_v75_)
                        rhs241_ = (self).f25
                        lhs207_ = d_1714_v73_
                        lhs208_ = default__.safeIndex(905, (d_1714_v73_).length(0))
                        lhs209_ = globalState
                        lhs210_ = self
                        lhs207_[lhs208_] = rhs238_
                        lhs209_.f10 = rhs239_
                        d_1718_v76_ = rhs240_
                        lhs210_.f24 = rhs241_
                        (globalState).f11 = len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([d_1655_v27_, (d_1693_v54_)[default__.safeIndex(174, (d_1693_v54_).length(0))]])) for d_1720_i8_ in range(default__.abs(723))]))
                        d_1721_v78_: D10
                        d_1721_v78_ = D10_DC22(d_1655_v27_, (self).f26)
                        d_1722_v79_: D10
                        d_1722_v79_ = D10_DC24(d_1721_v78_)
                        index248_ = default__.safeIndex(174, (d_1693_v54_).length(0))
                        rhs242_ = (self).f26
                        rhs243_ = (0) - (d_1655_v27_)
                        rhs244_ = D10_DC24(d_1721_v78_)
                        rhs245_ = d_1655_v27_
                        lhs211_ = self
                        lhs212_ = globalState
                        lhs213_ = d_1693_v54_
                        lhs214_ = default__.safeIndex(174, (d_1693_v54_).length(0))
                        lhs211_.f24 = rhs242_
                        lhs212_.f11 = rhs243_
                        d_1722_v79_ = rhs244_
                        lhs213_[lhs214_] = rhs245_
                        d_1723_v80_: _dafny.Array
                        nw273_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 25)
                        d_1723_v80_ = nw273_
                        d_1724_v81_: _dafny.Map
                        d_1724_v81_ = _dafny.Map({(d_1714_v73_)[default__.safeIndex(905, (d_1714_v73_).length(0))]: d_1723_v80_})
                        d_1724_v81_ = (d_1724_v81_).set((d_1714_v73_)[default__.safeIndex(905, (d_1714_v73_).length(0))], d_1723_v80_)
                    pass
            pass
        d_1725_v82_: C4
        nw274_ = C4()
        nw274_.ctor__(self.f24, self.f24)
        d_1725_v82_ = nw274_
        d_1687_v49_ = (d_1687_v49_).set(d_1655_v27_, default__.safeModuloInt(d_1655_v27_, d_1655_v27_))
        r0 = default__.fm36(d_1655_v27_, (self).f26, (self).f26, globalState)
        d_1726_v83_: D10
        d_1726_v83_ = D10_DC23(_dafny.SeqWithoutIsStrInference([self.f24]), self.f24)
        r1 = (d_1726_v83_).cf43
        return r0, r1


class C6(T1, T3, T0, T2):
    def  __init__(self):
        self._f24: bool = False
        self._f27: str = _dafny.CodePoint('D')
        self._f25: bool = False
        self._f26: bool = False
        self._f28: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C6"
    @property
    def f24(self):
        return self._f24
    @f24.setter
    def f24(self, value):
        self._f24 = value
    @property
    def f27(self):
        return self._f27
    @f27.setter
    def f27(self, value):
        self._f27 = value
    @property
    def f25(self):
        return self._f25
    @property
    def f26(self):
        return self._f26
    def ctor__(self, f28, f24, f26, f27, f25):
        (self)._f28 = f28
        (self).f24 = f24
        (self)._f26 = f26
        (self).f27 = f27
        (self)._f25 = f25

    def fm4(self, p0, p1, globalState):
        return D0_DC0((self).f25, (0) - (default__.safeModuloInt(41, 42)), -264)

    def fm5(self, p0, p1, p2, globalState):
        source29_ = D2_DC6(D2_DC5(self.f27, self.f24, len((D3_DC7(_dafny.SeqWithoutIsStrInference([(self).f26, (self).f25, (self).f25]))).cf15), 975))
        if source29_.is_DC4:
            d_1727___mcc_h0_ = source29_.cf9
            d_1728_cf9_ = d_1727___mcc_h0_
            return (_dafny.MultiSet([self.f24, self.f24])).cardinality
        elif source29_.is_DC5:
            d_1729___mcc_h1_ = source29_.cf10
            d_1730___mcc_h2_ = source29_.cf11
            d_1731___mcc_h3_ = source29_.cf12
            d_1732___mcc_h4_ = source29_.cf13
            d_1733_cf13_ = d_1732___mcc_h4_
            d_1734_cf12_ = d_1731___mcc_h3_
            d_1735_cf11_ = d_1730___mcc_h2_
            d_1736_cf10_ = d_1729___mcc_h1_
            return (0) - (d_1734_cf12_)
        elif source29_.is_DC3:
            d_1737___mcc_h5_ = source29_.cf8
            d_1738_cf8_ = d_1737___mcc_h5_
            def iife119_():
                coll61_ = _dafny.Set()
                compr_61_: int
                for compr_61_ in _dafny.IntegerRange(137, 979):
                    d_1739_v0_: int = compr_61_
                    if ((137) <= (d_1739_v0_)) and ((d_1739_v0_) < (979)):
                        coll61_ = coll61_.union(_dafny.Set([(d_1739_v0_) - (-500)]))
                return _dafny.Set(coll61_)
            def iife120_():
                coll62_ = _dafny.Set()
                compr_62_: int
                for compr_62_ in _dafny.IntegerRange(77, -901):
                    d_1740_v1_: int = compr_62_
                    if ((77) <= (d_1740_v1_)) and ((d_1740_v1_) < (-901)):
                        coll62_ = coll62_.union(_dafny.Set([(d_1740_v1_) - (-978)]))
                return _dafny.Set(coll62_)
            return len((iife119_()
            ) - (iife120_()
            ))
        elif True:
            d_1741___mcc_h6_ = source29_.cf14
            d_1742_cf14_ = d_1741___mcc_h6_
            return 86

    def fm7(self, p0, globalState):
        return self.f24

    def fm6(self, p0, globalState):
        return len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fss")))

    def fm17(self, p0, p1, globalState):
        def lambda81_(source30_):
            if source30_.is_DC4:
                d_1743___mcc_h0_ = source30_.cf9
                d_1744_cf9_ = d_1743___mcc_h0_
                return (d_1744_cf9_) and (True)
            elif source30_.is_DC5:
                d_1745___mcc_h1_ = source30_.cf10
                d_1746___mcc_h2_ = source30_.cf11
                d_1747___mcc_h3_ = source30_.cf12
                d_1748___mcc_h4_ = source30_.cf13
                d_1749_cf13_ = d_1748___mcc_h4_
                d_1750_cf12_ = d_1747___mcc_h3_
                d_1751_cf11_ = d_1746___mcc_h2_
                d_1752_cf10_ = d_1745___mcc_h1_
                return (d_1749_cf13_) < (d_1749_cf13_)
            elif source30_.is_DC3:
                d_1753___mcc_h5_ = source30_.cf8
                d_1754_cf8_ = d_1753___mcc_h5_
                def iife121_():
                    coll63_ = _dafny.Map()
                    compr_63_: int
                    for compr_63_ in _dafny.IntegerRange(435, 821):
                        d_1755_v0_: int = compr_63_
                        if ((435) <= (d_1755_v0_)) and ((d_1755_v0_) < (821)):
                            coll63_[(d_1755_v0_) + (len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([(_dafny.MultiSet([(self).f25, (self).f26])).cardinality])).cardinality, 484])))] = 223
                    return _dafny.Map(coll63_)
                return (_dafny.Map({len(_dafny.Map({D2_DC4(self.f24): (self).f26})): (self).f28})) == (_dafny.Map({(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([636, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(iife121_()
                )]))).cardinality]))).cardinality: False}))
            elif True:
                d_1756___mcc_h6_ = source30_.cf14
                d_1757_cf14_ = d_1756___mcc_h6_
                return (self).f28

        return not(lambda81_(D2_DC3(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iy"))): -143}))))

    def m2(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: bool = False
        r2: int = int(0)
        r3: int = int(0)
        d_1758_v0_: _dafny.MultiSet
        d_1758_v0_ = _dafny.MultiSet([(self).f28])
        d_1759_v1_: _dafny.Map
        d_1759_v1_ = _dafny.Map({default__.safeDivisionInt(p0, p0): (d_1758_v0_).cardinality})
        d_1760_v2_: _dafny.Map
        d_1760_v2_ = _dafny.Map({p0: p2})
        d_1761_v3_: _dafny.Set
        d_1761_v3_ = _dafny.Set({p0, p0, default__.fm0(p0, p2, (d_1760_v2_).set(p0, (self).f28), p0, globalState)})
        d_1759_v1_ = (d_1759_v1_).set(len(d_1761_v3_), p0)
        d_1762_v4_: _dafny.Array
        def lambda82_(d_1763_i0_):
            return not(self.f24)

        init49_ = lambda82_
        nw275_ = _dafny.Array(None, 7)
        for i0_49_ in range(nw275_.length(0)):
            nw275_[i0_49_] = init49_(i0_49_)
        d_1762_v4_ = nw275_
        d_1764_v5_: _dafny.Array
        nw276_ = _dafny.Array(None, 7)
        nw276_[int(0)] = d_1762_v4_
        nw276_[int(1)] = d_1762_v4_
        nw276_[int(2)] = d_1762_v4_
        nw276_[int(3)] = d_1762_v4_
        nw276_[int(4)] = d_1762_v4_
        nw276_[int(5)] = d_1762_v4_
        nw276_[int(6)] = d_1762_v4_
        d_1764_v5_ = nw276_
        d_1764_v5_ = d_1764_v5_
        d_1765_v6_: _dafny.Map
        d_1765_v6_ = _dafny.Map({d_1764_v5_: not(self.f24)})
        d_1766_v7_: _dafny.Array
        d_1766_v7_ = d_1764_v5_
        d_1765_v6_ = (d_1765_v6_).set((d_1766_v7_), not(p2))
        d_1767_i1_: int
        d_1767_i1_ = 0
        with _dafny.label("15"):
            while ((self).f28 if (self).f28 else (self).f26):
                with _dafny.c_label("15"):
                    if (d_1767_i1_) >= (100):
                        raise _dafny.Break("15")
                    d_1767_i1_ = (d_1767_i1_) + (1)
                    d_1768_v8_: _dafny.Seq
                    d_1768_v8_ = _dafny.SeqWithoutIsStrInference([(self).f25, (self).f25, False, (d_1758_v0_) != (d_1758_v0_), not(p2)])
                    d_1769_v9_: _dafny.MultiSet
                    d_1769_v9_ = _dafny.MultiSet([p0])
                    rhs246_ = 406
                    rhs247_ = (d_1768_v8_).set(default__.safeIndex((0) - (p0), len(d_1768_v8_)), self.f24)
                    rhs248_ = ((d_1768_v8_) <= (d_1768_v8_)) or ((d_1768_v8_)[default__.safeIndex(((d_1769_v9_)[-810] if (-810) in (d_1769_v9_) else 427), len(d_1768_v8_))])
                    rhs249_ = (self).fm17(self.f24, (self).f28, globalState)
                    lhs215_ = self
                    r2 = rhs246_
                    d_1768_v8_ = rhs247_
                    lhs215_.f24 = rhs248_
                    r0 = rhs249_
                    d_1770_v10_: _dafny.Array
                    nw277_ = _dafny.Array(int(0), 10)
                    d_1770_v10_ = nw277_
                    index249_ = default__.safeIndex(799, (d_1770_v10_).length(0))
                    (d_1770_v10_)[index249_] = p0
                    index250_ = default__.safeIndex(799, (d_1770_v10_).length(0))
                    (d_1770_v10_)[index250_] = (default__.safeDivisionInt(p0, p0)) * ((0) - (-921))
                    if (p2) and ((self).f28):
                        d_1771_v11_: _dafny.Map
                        d_1771_v11_ = _dafny.Map({p2: _dafny.SeqWithoutIsStrInference([p1 for d_1772_i2_ in range(default__.abs(391))])})
                        d_1773_v12_: _dafny.Seq
                        d_1773_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mi"))
                        d_1774_v13_: _dafny.MultiSet
                        d_1774_v13_ = _dafny.MultiSet([d_1773_v12_, d_1773_v12_, d_1773_v12_, d_1773_v12_, d_1773_v12_])
                        d_1775_v14_: _dafny.Map
                        d_1775_v14_ = _dafny.Map({(self).f25: len(_dafny.SeqWithoutIsStrInference([self.f27 for d_1776_i3_ in range(default__.abs(-581))]))})
                        d_1777_v15_: _dafny.Seq
                        d_1777_v15_ = _dafny.SeqWithoutIsStrInference([p0])
                        rhs250_ = (d_1770_v10_)[default__.safeIndex(799, (d_1770_v10_).length(0))]
                        rhs251_ = (default__.safeModuloInt((self).fm6(D0_DC0(p2, (0) - (p0), (d_1770_v10_)[default__.safeIndex(799, (d_1770_v10_).length(0))]), globalState), (d_1770_v10_)[default__.safeIndex(799, (d_1770_v10_).length(0))]) if (self).f25 else default__.safeModuloInt((self).fm5(d_1771_v11_, p2, (self).f28, globalState), (d_1774_v13_).cardinality))
                        rhs252_ = ((d_1775_v14_)[(d_1777_v15_) <= (d_1777_v15_)] if ((d_1777_v15_) <= (d_1777_v15_)) in (d_1775_v14_) else p0)
                        lhs216_ = globalState
                        lhs217_ = globalState
                        lhs216_.f22 = rhs250_
                        r2 = rhs251_
                        lhs217_.f22 = rhs252_
                        r0 = False
                        d_1778_v16_: D2
                        d_1778_v16_ = D2_DC4(not(((self).f25) or (p2)))
                        d_1778_v16_ = D2_DC4((d_1758_v0_).isdisjoint(d_1758_v0_))
                        r1 = ((p0 if self.f24 else len(d_1760_v2_))) in (d_1777_v15_)
                        r0 = (self).f28
                    elif True:
                        d_1779_v17_: C0
                        nw278_ = C0()
                        nw278_.ctor__(_dafny.SeqWithoutIsStrInference([self.f27 for d_1780_i4_ in range(default__.abs(-141))]))
                        d_1779_v17_ = nw278_
                        (globalState).f10 = p2
                        d_1759_v1_ = (d_1759_v1_).set(len((_dafny.SeqWithoutIsStrInference([(self).f26])) + (d_1768_v8_)), len(default__.fm26(p2, p0, len(d_1760_v2_), globalState)))
                        d_1781_v18_: _dafny.Map
                        d_1781_v18_ = _dafny.Map({self.f24: False})
                        d_1782_v19_: C1
                        nw279_ = C1()
                        nw279_.ctor__(((d_1781_v18_)[self.f24] if (self.f24) in (d_1781_v18_) else self.f24))
                        d_1782_v19_ = nw279_
                        d_1783_v20_: _dafny.Seq
                        d_1783_v20_ = _dafny.SeqWithoutIsStrInference([p1, p1])
                        d_1784_v21_: _dafny.Map
                        d_1784_v21_ = _dafny.Map({(self).f25: d_1783_v20_})
                        r3 = (d_1782_v19_).fm5(((d_1784_v21_).set(p2, d_1783_v20_)) | (_dafny.Map({True: _dafny.SeqWithoutIsStrInference([p1 for d_1785_i5_ in range(default__.abs(83))])})), p2, p2, globalState)
                    d_1786_v22_: _dafny.Seq
                    d_1786_v22_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tatdnps"))
                    d_1787_v23_: D1
                    d_1787_v23_ = D1_DC1(d_1786_v22_)
                    d_1788_v24_: _dafny.Map
                    d_1788_v24_ = _dafny.Map({d_1787_v23_: d_1770_v10_})
                    d_1789_v25_: _dafny.MultiSet
                    d_1789_v25_ = _dafny.MultiSet([self.f27])
                    d_1788_v24_ = ((d_1788_v24_) | (d_1788_v24_) if (_dafny.MultiSet([self.f27])).issubset(d_1789_v25_) else (_dafny.Map({d_1787_v23_: d_1770_v10_})).set(d_1787_v23_, d_1770_v10_))
                    pass
            pass
        d_1790_v26_: _dafny.Set
        d_1790_v26_ = _dafny.Set({self.f27, self.f27})
        d_1791_v27_: _dafny.Map
        d_1791_v27_ = _dafny.Map({p2: d_1790_v26_})
        d_1791_v27_ = (d_1791_v27_).set((self).f28, d_1790_v26_)
        d_1792_v28_: D3
        d_1792_v28_ = D3_DC8((self).f28, 290, p0)
        (self).f24 = ((self).f28 if (d_1792_v28_).cf16 else self.f24)
        r0 = (self).f26
        r1 = (self).f28
        r2 = -358
        r3 = (self).fm6(p1, globalState)
        return r0, r1, r2, r3

    def m3(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: bool = False
        r2: _dafny.Map = _dafny.Map({})
        d_1793_v0_: _dafny.Array
        nw280_ = _dafny.Array(_dafny.Seq({}), 25)
        d_1793_v0_ = nw280_
        d_1794_v1_: _dafny.Map
        d_1794_v1_ = _dafny.Map({(self).f26: p0})
        d_1795_v2_: _dafny.Map
        d_1795_v2_ = _dafny.Map({817: True})
        d_1796_v3_: _dafny.Map
        d_1796_v3_ = _dafny.Map({self.f24: (self).f28})
        d_1797_v4_: _dafny.Seq
        d_1797_v4_ = _dafny.SeqWithoutIsStrInference([p0, p0, default__.fm0(624, default__.fm1((self).f28, (self).f28, len(_dafny.Map({d_1794_v1_: (self).f26})), globalState), d_1795_v2_, len(d_1796_v3_), globalState)])
        index251_ = default__.safeIndex(500, (d_1793_v0_).length(0))
        (d_1793_v0_)[index251_] = d_1797_v4_
        d_1798_v5_: _dafny.Set
        d_1798_v5_ = _dafny.Set({self.f27})
        d_1799_v6_: _dafny.Array
        nw281_ = _dafny.Array(None, 7)
        nw281_[int(0)] = p0
        nw281_[int(1)] = p0
        nw281_[int(2)] = p0
        nw281_[int(3)] = 255
        nw281_[int(4)] = p0
        nw281_[int(5)] = len(d_1798_v5_)
        nw281_[int(6)] = 956
        d_1799_v6_ = nw281_
        d_1800_v7_: D1
        d_1800_v7_ = D1_DC2(d_1799_v6_, self.f27, (_dafny.MultiSet([p0])).cardinality, (self).f25)
        index252_ = default__.safeIndex(500, (d_1793_v0_).length(0))
        rhs253_ = (self).f25
        rhs254_ = (d_1797_v4_) + (d_1797_v4_)
        rhs255_ = ((self).fm7(d_1794_v1_, globalState)) and ((d_1800_v7_).cf7)
        lhs218_ = d_1793_v0_
        lhs219_ = default__.safeIndex(500, (d_1793_v0_).length(0))
        lhs220_ = globalState
        r1 = rhs253_
        lhs218_[lhs219_] = rhs254_
        lhs220_.f10 = rhs255_
        (self).f24 = (self).f25
        source31_ = D1_DC1(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yxq")))
        if source31_.is_DC2:
            d_1801___mcc_h0_ = source31_.cf4
            d_1802___mcc_h1_ = source31_.cf5
            d_1803___mcc_h2_ = source31_.cf6
            d_1804___mcc_h3_ = source31_.cf7
            d_1805_cf7_ = d_1804___mcc_h3_
            d_1806_cf6_ = d_1803___mcc_h2_
            d_1807_cf5_ = d_1802___mcc_h1_
            d_1808_cf4_ = d_1801___mcc_h0_
            d_1795_v2_ = d_1795_v2_
            d_1809_v8_: _dafny.Seq
            d_1809_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nbvpuduui"))
            d_1810_v9_: C2
            nw282_ = C2()
            nw282_.ctor__(d_1809_v8_, d_1809_v8_)
            d_1810_v9_ = nw282_
            (self).f24 = (self).fm17((self).f28, (self).f26, globalState)
            index253_ = default__.safeIndex(952, (p2).length(0))
            (p2)[index253_] = d_1805_cf7_
            index254_ = default__.safeIndex(952, (p2).length(0))
            (p2)[index254_] = ((self).f26) or ((self).f26)
        elif True:
            d_1811___mcc_h4_ = source31_.cf3
            d_1812_cf3_ = d_1811___mcc_h4_
            index255_ = default__.safeIndex(597, (d_1799_v6_).length(0))
            (d_1799_v6_)[index255_] = p0
            index256_ = default__.safeIndex(597, (d_1799_v6_).length(0))
            (d_1799_v6_)[index256_] = p0
            d_1813_v10_: D1
            d_1813_v10_ = D1_DC1(d_1812_cf3_)
            d_1814_v11_: _dafny.Array
            nw283_ = _dafny.Array(None, 7)
            nw283_[int(0)] = (d_1813_v10_).cf3
            nw283_[int(1)] = d_1812_cf3_
            nw283_[int(2)] = d_1812_cf3_
            nw283_[int(3)] = d_1812_cf3_
            nw283_[int(4)] = (d_1812_cf3_ if p1 else (d_1812_cf3_).set(default__.safeIndex((d_1799_v6_)[default__.safeIndex(597, (d_1799_v6_).length(0))], len(d_1812_cf3_)), self.f27))
            nw283_[int(5)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "usp"))).set(default__.safeIndex((d_1799_v6_)[default__.safeIndex(597, (d_1799_v6_).length(0))], len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "usp")))), self.f27)
            nw283_[int(6)] = d_1812_cf3_
            d_1814_v11_ = nw283_
            index257_ = default__.safeIndex(319, (d_1814_v11_).length(0))
            (d_1814_v11_)[index257_] = d_1812_cf3_
            index258_ = default__.safeIndex(319, (d_1814_v11_).length(0))
            rhs256_ = p1
            rhs257_ = ((_dafny.SeqWithoutIsStrInference([self.f27 for d_1815_i0_ in range(default__.abs(787))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "napupek")))) + ((d_1812_cf3_) + (d_1812_cf3_))
            lhs221_ = self
            lhs222_ = d_1814_v11_
            lhs223_ = default__.safeIndex(319, (d_1814_v11_).length(0))
            lhs221_.f24 = rhs256_
            lhs222_[lhs223_] = rhs257_
            (globalState).f10 = (self).f28
            d_1812_cf3_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rdauad"))).set(default__.safeIndex(default__.fm0(-105, (self).f28, d_1795_v2_, (d_1799_v6_)[default__.safeIndex(597, (d_1799_v6_).length(0))], globalState), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rdauad")))), self.f27)
        d_1816_v12_: D3
        d_1816_v12_ = D3_DC8(True, (0) - (((d_1794_v1_)[p1] if (p1) in (d_1794_v1_) else p0)), p0)
        def lambda83_(source33_):
            if source33_.is_DC8:
                d_1817___mcc_h10_ = source33_.cf16
                d_1818___mcc_h11_ = source33_.cf17
                d_1819___mcc_h12_ = source33_.cf18
                d_1820_cf18_ = d_1819___mcc_h12_
                d_1821_cf17_ = d_1818___mcc_h11_
                d_1822_cf16_ = d_1817___mcc_h10_
                return D1_DC1(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "egasiyaw")))
            elif True:
                d_1823___mcc_h13_ = source33_.cf15
                d_1824_cf15_ = d_1823___mcc_h13_
                if (self).f28:
                    return D1_DC1(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "otoouwy")))
                elif True:
                    return D1_DC1(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wpqa")))

        source32_ = lambda83_(d_1816_v12_)
        if source32_.is_DC2:
            d_1825___mcc_h5_ = source32_.cf4
            d_1826___mcc_h6_ = source32_.cf5
            d_1827___mcc_h7_ = source32_.cf6
            d_1828___mcc_h8_ = source32_.cf7
            d_1829_cf7_ = d_1828___mcc_h8_
            d_1830_cf6_ = d_1827___mcc_h7_
            d_1831_cf5_ = d_1826___mcc_h6_
            d_1832_cf4_ = d_1825___mcc_h5_
            index259_ = default__.safeIndex(500, (d_1793_v0_).length(0))
            (d_1793_v0_)[index259_] = d_1797_v4_
            if not(self.f24):
                index260_ = default__.safeIndex(624, (d_1832_cf4_).length(0))
                (d_1832_cf4_)[index260_] = p0
                index261_ = default__.safeIndex(624, (d_1832_cf4_).length(0))
                (d_1832_cf4_)[index261_] = default__.safeModuloInt(p0, p0)
                d_1833_v13_: _dafny.MultiSet
                d_1833_v13_ = _dafny.MultiSet([(self).f26, d_1829_cf7_])
                d_1834_v14_: _dafny.Seq
                d_1834_v14_ = _dafny.SeqWithoutIsStrInference([p1])
                d_1835_v15_: _dafny.Map
                d_1835_v15_ = _dafny.Map({(d_1832_cf4_)[default__.safeIndex(624, (d_1832_cf4_).length(0))]: d_1833_v13_})
                d_1836_v16_: _dafny.Seq
                d_1836_v16_ = _dafny.SeqWithoutIsStrInference([((d_1835_v15_)[p0] if (p0) in (d_1835_v15_) else d_1833_v13_), d_1833_v13_])
                d_1837_v17_: _dafny.Array
                nw284_ = _dafny.Array(None, 23)
                nw284_[int(0)] = d_1833_v13_
                nw284_[int(1)] = d_1833_v13_
                nw284_[int(2)] = d_1833_v13_
                nw284_[int(3)] = _dafny.MultiSet(d_1834_v14_)
                nw284_[int(4)] = _dafny.MultiSet([not(self.f24), False])
                nw284_[int(5)] = (d_1833_v13_) | (default__.fm27((self).f25, globalState))
                nw284_[int(6)] = default__.fm27(d_1829_cf7_, globalState)
                nw284_[int(7)] = d_1833_v13_
                nw284_[int(8)] = (d_1836_v16_)[default__.safeIndex(len(default__.fm28(globalState)), len(d_1836_v16_))]
                nw284_[int(9)] = d_1833_v13_
                nw284_[int(10)] = (default__.fm27((self).f25, globalState)) | (d_1833_v13_)
                nw284_[int(11)] = (d_1833_v13_ if False else d_1833_v13_)
                nw284_[int(12)] = (d_1833_v13_) | (d_1833_v13_)
                nw284_[int(13)] = d_1833_v13_
                nw284_[int(14)] = (_dafny.MultiSet([self.f24])).set(self.f24, default__.abs(p0))
                nw284_[int(15)] = (_dafny.MultiSet(d_1834_v14_)).intersection(_dafny.MultiSet([d_1829_cf7_, p1, (self).f25, p1]))
                nw284_[int(16)] = d_1833_v13_
                nw284_[int(17)] = (_dafny.MultiSet([d_1829_cf7_])).intersection((d_1833_v13_).set((self).f28, default__.abs((d_1832_cf4_)[default__.safeIndex(624, (d_1832_cf4_).length(0))])))
                nw284_[int(18)] = d_1833_v13_
                nw284_[int(19)] = (d_1833_v13_ if self.f24 else d_1833_v13_)
                nw284_[int(20)] = _dafny.MultiSet([p1])
                nw284_[int(21)] = d_1833_v13_
                nw284_[int(22)] = (d_1833_v13_).intersection(d_1833_v13_)
                d_1837_v17_ = nw284_
                index262_ = default__.safeIndex(283, (d_1837_v17_).length(0))
                (d_1837_v17_)[index262_] = (d_1833_v13_).intersection(_dafny.MultiSet([p1, True]))
                index263_ = default__.safeIndex(283, (d_1837_v17_).length(0))
                (d_1837_v17_)[index263_] = (d_1833_v13_).intersection((d_1833_v13_) - (d_1833_v13_))
                (globalState).f22 = p0
                index264_ = default__.safeIndex(689, (p2).length(0))
                (p2)[index264_] = p1
                index265_ = default__.safeIndex(689, (p2).length(0))
                (p2)[index265_] = not((d_1834_v14_)[default__.safeIndex(default__.safeModuloInt(p0, p0), len(d_1834_v14_))])
                d_1838_v18_: _dafny.Set
                d_1838_v18_ = _dafny.Set({self.f24, (p2)[default__.safeIndex(689, (p2).length(0))]})
                (globalState).f16 = default__.safeModuloInt((len(d_1838_v18_)) + (p0), len(_dafny.Map({True: (p2)[default__.safeIndex(689, (p2).length(0))]})))
            elif True:
                d_1839_v19_: _dafny.Seq
                d_1839_v19_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dvxphtst"))
                d_1840_v20_: C0
                nw285_ = C0()
                nw285_.ctor__((_dafny.SeqWithoutIsStrInference([self.f27 for d_1841_i1_ in range(default__.abs(-157))]) if p1 else d_1839_v19_))
                d_1840_v20_ = nw285_
                d_1842_v21_: _dafny.Array
                def lambda84_(d_1843_v2_):
                    def lambda85_(d_1844_i2_):
                        return (d_1843_v2_) | (d_1843_v2_)

                    return lambda85_

                init50_ = lambda84_(d_1795_v2_)
                nw286_ = _dafny.Array(None, 15)
                for i0_50_ in range(nw286_.length(0)):
                    nw286_[i0_50_] = init50_(i0_50_)
                d_1842_v21_ = nw286_
                index266_ = default__.safeIndex(499, (d_1842_v21_).length(0))
                (d_1842_v21_)[index266_] = d_1795_v2_
                index267_ = default__.safeIndex(499, (d_1842_v21_).length(0))
                rhs258_ = (d_1794_v1_) | ((d_1794_v1_) | (d_1794_v1_))
                rhs259_ = d_1840_v20_
                rhs260_ = (d_1795_v2_) | (d_1795_v2_)
                lhs224_ = d_1842_v21_
                lhs225_ = default__.safeIndex(499, (d_1842_v21_).length(0))
                d_1794_v1_ = rhs258_
                d_1840_v20_ = rhs259_
                lhs224_[lhs225_] = rhs260_
                index268_ = default__.safeIndex(934, (d_1799_v6_).length(0))
                (d_1799_v6_)[index268_] = p0
                d_1845_v22_: T3
                nw287_ = C5()
                nw287_.ctor__((self).f25, self.f27, d_1829_cf7_, d_1829_cf7_)
                d_1845_v22_ = nw287_
                d_1846_v23_: _dafny.Map
                d_1846_v23_ = _dafny.Map({True: d_1845_v22_})
                index269_ = default__.safeIndex(934, (d_1799_v6_).length(0))
                def iife122_():
                    coll64_ = _dafny.Map()
                    compr_64_: int
                    for compr_64_ in _dafny.IntegerRange(955, 313):
                        d_1847_v24_: int = compr_64_
                        if ((955) <= (d_1847_v24_)) and ((d_1847_v24_) < (313)):
                            coll64_[(d_1847_v24_) - (p0)] = d_1845_v22_.f24
                    return _dafny.Map(coll64_)
                (d_1799_v6_)[index269_] = default__.fm0(len(d_1846_v23_), (not(self.f24) if False else (self).f26), ((iife122_()
                ).set(((d_1793_v0_)[default__.safeIndex(500, (d_1793_v0_).length(0))])[default__.safeIndex(d_1830_cf6_, len((d_1793_v0_)[default__.safeIndex(500, (d_1793_v0_).length(0))]))], d_1845_v22_.f24)) | (d_1795_v2_), p0, globalState)
                d_1848_v25_: _dafny.Seq
                d_1848_v25_ = _dafny.SeqWithoutIsStrInference([d_1839_v19_, d_1839_v19_])
                def iife123_():
                    coll65_ = _dafny.Map()
                    compr_65_: _dafny.Seq
                    for compr_65_ in (d_1848_v25_).Elements:
                        d_1849_v26_: _dafny.Seq = compr_65_
                        if (d_1849_v26_) in (d_1848_v25_):
                            coll65_[d_1849_v26_] = ((d_1794_v1_)[(self).f25] if ((self).f25) in (d_1794_v1_) else p0)
                    return _dafny.Map(coll65_)
                (globalState).f10 = (((d_1840_v20_).f31) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "trgbqclqk")))) <= ((d_1848_v25_)[default__.safeIndex(len(iife123_()
                ), len(d_1848_v25_))])
                r1 = not((297) != (p0))
                d_1850_v27_: D0
                d_1850_v27_ = D0_DC0(not(p1), default__.fm0((d_1799_v6_)[default__.safeIndex(934, (d_1799_v6_).length(0))], self.f24, _dafny.Map({d_1830_cf6_: p1}), d_1830_cf6_, globalState), d_1830_cf6_)
                d_1851_v28_: bool
                d_1852_v29_: bool
                d_1853_v30_: int
                d_1854_v31_: int
                out35_: bool
                out36_: bool
                out37_: int
                out38_: int
                out35_, out36_, out37_, out38_ = (d_1845_v22_).m2(d_1830_cf6_, d_1850_v27_, (d_1845_v22_).f26, globalState)
                d_1851_v28_ = out35_
                d_1852_v29_ = out36_
                d_1853_v30_ = out37_
                d_1854_v31_ = out38_
            d_1855_v32_: _dafny.Set
            d_1855_v32_ = _dafny.Set({p0})
            (globalState).f22 = ((0) - ((len(d_1855_v32_)) * (p0))) * (d_1830_cf6_)
            (globalState).f22 = (p0) + (p0)
        elif True:
            d_1856___mcc_h9_ = source32_.cf3
            d_1857_cf3_ = d_1856___mcc_h9_
            (globalState).f10 = not(not((self).f25))
            r1 = (self).f25
            (globalState).f11 = ((d_1793_v0_)[default__.safeIndex(500, (d_1793_v0_).length(0))])[default__.safeIndex(628, len((d_1793_v0_)[default__.safeIndex(500, (d_1793_v0_).length(0))]))]
            d_1858_v33_: _dafny.Set
            d_1858_v33_ = _dafny.Set({p0, len(d_1857_cf3_), p0})
            (globalState).f10 = not(((self).f28 if (d_1858_v33_).isdisjoint(default__.fm53(self.f27, self.f27, globalState)) else ((self).f28 if (self).f25 else (self).f26)))
        d_1859_i3_: int
        d_1859_i3_ = 0
        with _dafny.label("16"):
            while (default__.fm60(688, globalState)) == (d_1798_v5_):
                with _dafny.c_label("16"):
                    if (d_1859_i3_) >= (100):
                        raise _dafny.Break("16")
                    d_1859_i3_ = (d_1859_i3_) + (1)
                    d_1860_v34_: _dafny.Set
                    d_1860_v34_ = _dafny.Set({(p1) or ((self).f25), not(p1), (self).f28})
                    d_1861_v35_: _dafny.MultiSet
                    d_1861_v35_ = _dafny.MultiSet([p1])
                    rhs261_ = 123
                    rhs262_ = (d_1860_v34_) | (d_1860_v34_)
                    rhs263_ = ((d_1861_v35_)[self.f24] if (self.f24) in (d_1861_v35_) else p0)
                    lhs226_ = globalState
                    lhs227_ = globalState
                    lhs226_.f16 = rhs261_
                    d_1860_v34_ = rhs262_
                    lhs227_.f22 = rhs263_
                    (globalState).f10 = (p0) > (p0)
                    d_1860_v34_ = d_1860_v34_
                    d_1862_v37_: _dafny.Array
                    def lambda86_(d_1863_p1_, d_1864_v4_, d_1865_p0_):
                        def lambda87_(d_1866_i4_):
                            def iife124_():
                                coll66_ = _dafny.Set()
                                compr_66_: int
                                for compr_66_ in _dafny.IntegerRange(272, 290):
                                    d_1867_v36_: int = compr_66_
                                    if ((272) <= (d_1867_v36_)) and ((d_1867_v36_) < (290)):
                                        coll66_ = coll66_.union(_dafny.Set([(d_1867_v36_) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qjjn"))))]))
                                return _dafny.Set(coll66_)
                            return (_dafny.Map({_dafny.SeqWithoutIsStrInference([(D3_DC8(d_1863_p1_, len(d_1864_v4_), d_1865_p0_)).cf16]): _dafny.Set({d_1865_p0_})})) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([self.f24]): iife124_()
                            }))

                        return lambda87_

                    init51_ = lambda86_(p1, d_1797_v4_, p0)
                    nw288_ = _dafny.Array(None, 3)
                    for i0_51_ in range(nw288_.length(0)):
                        nw288_[i0_51_] = init51_(i0_51_)
                    d_1862_v37_ = nw288_
                    d_1868_v38_: _dafny.Seq
                    d_1868_v38_ = _dafny.SeqWithoutIsStrInference([p1, (self).f25])
                    d_1869_v39_: _dafny.MultiSet
                    d_1869_v39_ = _dafny.MultiSet([p0])
                    d_1870_v40_: _dafny.MultiSet
                    d_1870_v40_ = (d_1869_v39_).set(p0, default__.abs(len(_dafny.SeqWithoutIsStrInference([self.f27 for d_1871_i5_ in range(default__.abs(373))]))))
                    d_1872_v41_: _dafny.Map
                    d_1872_v41_ = _dafny.Map({p0: default__.fm52(True, globalState)})
                    d_1873_v42_: _dafny.Set
                    d_1873_v42_ = _dafny.Set({p0, ((d_1870_v40_)).cardinality, p0, p0, (((d_1872_v41_)[p0] if (p0) in (d_1872_v41_) else _dafny.MultiSet([p0, p0]))).cardinality})
                    d_1874_v43_: _dafny.Map
                    d_1874_v43_ = _dafny.Map({d_1868_v38_: d_1873_v42_})
                    index270_ = default__.safeIndex(734, (d_1862_v37_).length(0))
                    (d_1862_v37_)[index270_] = d_1874_v43_
                    index271_ = default__.safeIndex(734, (d_1862_v37_).length(0))
                    (d_1862_v37_)[index271_] = d_1874_v43_
                    pass
            pass
        d_1875_v45_: _dafny.Map
        def iife125_():
            coll67_ = _dafny.Map()
            compr_67_: int
            for compr_67_ in _dafny.IntegerRange(238, -705):
                d_1876_v44_: int = compr_67_
                if ((238) <= (d_1876_v44_)) and ((d_1876_v44_) < (-705)):
                    coll67_[(d_1876_v44_) * (p0)] = p0
            return _dafny.Map(coll67_)
        d_1875_v45_ = _dafny.Map({iife125_()
        : d_1799_v6_})
        d_1875_v45_ = (d_1875_v45_).set(_dafny.Map({p0: -343}), d_1799_v6_)
        r0 = (self).f28
        r1 = (self).f28
        d_1877_v46_: D2
        d_1877_v46_ = D2_DC4((self).f26)
        d_1878_v47_: D2
        d_1878_v47_ = D2_DC6(d_1877_v46_)
        pat_let_tv13_ = p0
        pat_let_tv14_ = p1
        pat_let_tv15_ = p1
        pat_let_tv16_ = p0
        pat_let_tv17_ = p0
        pat_let_tv18_ = p0
        def lambda88_(source34_):
            if source34_.is_DC4:
                d_1879___mcc_h14_ = source34_.cf9
                d_1880_cf9_ = d_1879___mcc_h14_
                return _dafny.Map({_dafny.SeqWithoutIsStrInference([self.f27 for d_1881_i6_ in range(default__.abs(206))]): pat_let_tv13_})
            elif source34_.is_DC5:
                d_1882___mcc_h15_ = source34_.cf10
                d_1883___mcc_h16_ = source34_.cf11
                d_1884___mcc_h17_ = source34_.cf12
                d_1885___mcc_h18_ = source34_.cf13
                d_1886_cf13_ = d_1885___mcc_h18_
                d_1887_cf12_ = d_1884___mcc_h17_
                d_1888_cf11_ = d_1883___mcc_h16_
                d_1889_cf10_ = d_1882___mcc_h15_
                return _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "an")): d_1886_cf13_})
            elif source34_.is_DC3:
                d_1890___mcc_h19_ = source34_.cf8
                d_1891_cf8_ = d_1890___mcc_h19_
                def iife126_():
                    coll68_ = _dafny.Map()
                    compr_68_: _dafny.Seq
                    for compr_68_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference([self.f27 for d_1892_i7_ in range(default__.abs(376))]): pat_let_tv14_})).keys.Elements:
                        d_1893_v48_: _dafny.Seq = compr_68_
                        if (d_1893_v48_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference([self.f27 for d_1892_i7_ in range(default__.abs(376))]): pat_let_tv15_})):
                            coll68_[d_1893_v48_] = pat_let_tv16_
                    return _dafny.Map(coll68_)
                return (iife126_()
                ) | (_dafny.Map({(D1_DC1(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oek")))).cf3: pat_let_tv17_}))
            elif True:
                d_1894___mcc_h20_ = source34_.cf14
                d_1895_cf14_ = d_1894___mcc_h20_
                return (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n")): (_dafny.MultiSet([(self).f28, True])).cardinality})) | (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jjh")): pat_let_tv18_}))

        r2 = lambda88_(d_1878_v47_)
        return r0, r1, r2

    def m0(self, globalState):
        r0: bool = False
        r1: _dafny.Array = _dafny.Array(None, 0)
        d_1896_v0_: int
        d_1896_v0_ = 442
        d_1897_v1_: _dafny.Set
        d_1897_v1_ = _dafny.Set({d_1896_v0_})
        d_1898_v2_: _dafny.Seq
        d_1898_v2_ = _dafny.SeqWithoutIsStrInference([(d_1897_v1_).issubset(d_1897_v1_), (self).f26])
        d_1898_v2_ = d_1898_v2_
        hi11_ = d_1896_v0_
        for d_1899_i0_ in range(d_1896_v0_, hi11_):
            d_1900_v3_: _dafny.Seq
            d_1900_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jujcsfde"))
            d_1901_v6_: _dafny.Array
            def lambda89_(d_1902_v2_, d_1903_v0_, d_1904_i0_):
                def lambda90_(d_1905_i1_):
                    def iife127_():
                        coll69_ = _dafny.Map()
                        compr_69_: _dafny.Seq
                        for compr_69_ in (_dafny.SeqWithoutIsStrInference([d_1902_v2_])).Elements:
                            d_1906_v4_: _dafny.Seq = compr_69_
                            if (d_1906_v4_) in (_dafny.SeqWithoutIsStrInference([d_1902_v2_])):
                                def iife128_():
                                    coll70_ = _dafny.Map()
                                    compr_70_: int
                                    for compr_70_ in _dafny.IntegerRange(714, 910):
                                        d_1907_v5_: int = compr_70_
                                        if ((714) <= (d_1907_v5_)) and ((d_1907_v5_) < (910)):
                                            coll70_[(d_1907_v5_) + (d_1903_v0_)] = d_1904_i0_
                                    return _dafny.Map(coll70_)
                                coll69_[d_1906_v4_] = iife128_()

                        return _dafny.Map(coll69_)
                    return iife127_()
                    

                return lambda90_

            init52_ = lambda89_(d_1898_v2_, d_1896_v0_, d_1899_i0_)
            nw289_ = _dafny.Array(None, 19)
            for i0_52_ in range(nw289_.length(0)):
                nw289_[i0_52_] = init52_(i0_52_)
            d_1901_v6_ = nw289_
            d_1908_v7_: _dafny.MultiSet
            d_1908_v7_ = _dafny.MultiSet([d_1896_v0_, d_1896_v0_])
            d_1909_v8_: _dafny.Map
            d_1909_v8_ = _dafny.Map({d_1899_i0_: (self).f25})
            d_1910_v9_: _dafny.Map
            d_1910_v9_ = _dafny.Map({default__.fm0((d_1908_v7_).cardinality, (self).f26, d_1909_v8_, d_1899_i0_, globalState): d_1896_v0_})
            d_1911_v10_: _dafny.Map
            d_1911_v10_ = _dafny.Map({d_1898_v2_: d_1910_v9_})
            index272_ = default__.safeIndex(308, (d_1901_v6_).length(0))
            (d_1901_v6_)[index272_] = d_1911_v10_
            d_1912_v11_: _dafny.Map
            d_1912_v11_ = _dafny.Map({not((self).f26): _dafny.SeqWithoutIsStrInference([(self).f25])})
            index273_ = default__.safeIndex(308, (d_1901_v6_).length(0))
            rhs264_ = d_1900_v3_
            rhs265_ = (D2_DC5(self.f27, (self).f25, d_1899_i0_, d_1899_i0_)).cf11
            rhs266_ = d_1911_v10_
            rhs267_ = default__.safeModuloInt(597, -922)
            rhs268_ = ((d_1912_v11_) | (d_1912_v11_)) | (d_1912_v11_)
            lhs228_ = globalState
            lhs229_ = d_1901_v6_
            lhs230_ = default__.safeIndex(308, (d_1901_v6_).length(0))
            lhs231_ = globalState
            d_1900_v3_ = rhs264_
            lhs228_.f10 = rhs265_
            lhs229_[lhs230_] = rhs266_
            lhs231_.f11 = rhs267_
            d_1912_v11_ = rhs268_
            d_1913_v12_: _dafny.Array
            def lambda91_(d_1914_v0_):
                def lambda92_(d_1915_i2_):
                    return (_dafny.SeqWithoutIsStrInference([519 for d_1916_i3_ in range(default__.abs(225))])) + (_dafny.SeqWithoutIsStrInference([d_1914_v0_]))

                return lambda92_

            init53_ = lambda91_(d_1896_v0_)
            nw290_ = _dafny.Array(None, 1)
            for i0_53_ in range(nw290_.length(0)):
                nw290_[i0_53_] = init53_(i0_53_)
            d_1913_v12_ = nw290_
            d_1917_v13_: C3
            nw291_ = C3()
            nw291_.ctor__(self.f24, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bylgvmw")), not((self).f28), self.f27, self.f24, True)
            d_1917_v13_ = nw291_
            d_1918_v14_: D11
            d_1918_v14_ = D11_DC25(d_1917_v13_)
            d_1919_v15_: _dafny.Map
            d_1919_v15_ = _dafny.Map({d_1918_v14_: (self).f28})
            rhs269_ = (d_1900_v3_) <= ((d_1917_v13_).f33)
            rhs270_ = not ((d_1917_v13_).f32) or ((self).f28)
            rhs271_ = (d_1913_v12_ if (self).f28 else d_1913_v12_)
            rhs272_ = (d_1919_v15_) | ((d_1919_v15_) | (_dafny.Map({d_1918_v14_: (d_1917_v13_).f32})))
            rhs273_ = default__.safeDivisionInt((d_1896_v0_) + (d_1896_v0_), d_1896_v0_)
            lhs232_ = globalState
            lhs233_ = self
            lhs234_ = globalState
            lhs232_.f10 = rhs269_
            lhs233_.f24 = rhs270_
            d_1913_v12_ = rhs271_
            d_1919_v15_ = rhs272_
            lhs234_.f22 = rhs273_
            r0 = self.f24
            d_1920_v16_: D0
            d_1920_v16_ = D0_DC0((self).f25, d_1899_i0_, d_1896_v0_)
            d_1920_v16_ = d_1920_v16_
        d_1921_v17_: _dafny.Array
        nw292_ = _dafny.Array(False, 29)
        d_1921_v17_ = nw292_
        r1 = d_1921_v17_
        d_1922_i4_: int
        d_1922_i4_ = 0
        with _dafny.label("17"):
            while (self).f28:
                with _dafny.c_label("17"):
                    if (d_1922_i4_) >= (100):
                        raise _dafny.Break("17")
                    d_1922_i4_ = (d_1922_i4_) + (1)
                    d_1923_v18_: _dafny.Map
                    d_1923_v18_ = _dafny.Map({d_1921_v17_: (self).f26})
                    d_1924_v19_: _dafny.Map
                    d_1924_v19_ = _dafny.Map({(self).f26: d_1896_v0_})
                    d_1923_v18_ = (d_1923_v18_).set(d_1921_v17_, (self).fm7(d_1924_v19_, globalState))
                    d_1925_v20_: _dafny.Seq
                    d_1925_v20_ = _dafny.SeqWithoutIsStrInference([d_1896_v0_])
                    d_1926_v21_: D8
                    d_1926_v21_ = D8_DC19(True, 419)
                    d_1927_v22_: _dafny.Map
                    d_1927_v22_ = _dafny.Map({len(d_1925_v20_): d_1926_v21_})
                    d_1928_v23_: _dafny.Map
                    d_1928_v23_ = _dafny.Map({(0) - (d_1896_v0_): (self).f26})
                    d_1929_v24_: _dafny.Map
                    d_1929_v24_ = _dafny.Map({not((self).f25): (self).f25})
                    d_1927_v22_ = (d_1927_v22_).set(d_1896_v0_, D8_DC19((self).f25, default__.fm0((0) - (d_1896_v0_), self.f24, d_1928_v23_, len(d_1929_v24_), globalState)))
                    d_1930_v25_: _dafny.Array
                    nw293_ = _dafny.Array(None, 5)
                    nw293_[int(0)] = (0) - (d_1896_v0_)
                    nw293_[int(1)] = (0) - (-27)
                    nw293_[int(2)] = d_1896_v0_
                    nw293_[int(3)] = d_1896_v0_
                    nw293_[int(4)] = d_1896_v0_
                    d_1930_v25_ = nw293_
                    d_1931_v26_: _dafny.Seq
                    d_1931_v26_ = _dafny.SeqWithoutIsStrInference([d_1930_v25_, d_1930_v25_, d_1930_v25_])
                    (globalState).f11 = ((0) - (len(d_1931_v26_))) + ((len(d_1898_v2_) if (self).f25 else (self).fm5(default__.fm55(self.f27, (self).f25, (self).f25, self.f24, globalState), (self).f28, False, globalState)))
                    r0 = (_dafny.Set({not(self.f24), (self).fm17((self).f28, False, globalState), (self).f25})).issubset(_dafny.Set({self.f24, (self).f28}))
                    pass
            pass
        d_1932_v27_: _dafny.Set
        d_1932_v27_ = _dafny.Set({self.f24})
        d_1933_v28_: _dafny.Map
        d_1933_v28_ = _dafny.Map({d_1896_v0_: (self).f28})
        hi12_ = default__.safeDivisionInt(default__.fm0(len(d_1932_v27_), ((d_1933_v28_)[d_1896_v0_] if (d_1896_v0_) in (d_1933_v28_) else (self).f26), d_1933_v28_, d_1896_v0_, globalState), d_1896_v0_)
        for d_1934_i5_ in range(d_1896_v0_, hi12_):
            d_1935_v29_: _dafny.Seq
            d_1935_v29_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iw"))
            d_1936_v30_: _dafny.Array
            nw294_ = _dafny.Array(None, 8)
            nw294_[int(0)] = d_1934_i5_
            nw294_[int(1)] = 85
            nw294_[int(2)] = d_1896_v0_
            nw294_[int(3)] = ((self).fm6(D0_DC0((self).f25, d_1934_i5_, d_1934_i5_), globalState) if (self).f26 else 704)
            nw294_[int(4)] = (d_1896_v0_ if True else 847)
            nw294_[int(5)] = len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fpjnuvr")) if not((self).f28) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b"))))
            nw294_[int(6)] = 847
            nw294_[int(7)] = len(d_1935_v29_)
            d_1936_v30_ = nw294_
            index274_ = default__.safeIndex(467, (d_1936_v30_).length(0))
            (d_1936_v30_)[index274_] = default__.safeDivisionInt(d_1896_v0_, (0) - (len(d_1898_v2_)))
            d_1937_v31_: _dafny.Array
            nw295_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 23)
            d_1937_v31_ = nw295_
            index275_ = default__.safeIndex(155, (d_1937_v31_).length(0))
            (d_1937_v31_)[index275_] = d_1935_v29_
            d_1938_v32_: _dafny.Array
            nw296_ = _dafny.Array(None, 7)
            nw296_[int(0)] = d_1921_v17_
            nw296_[int(1)] = (d_1921_v17_ if (self).f25 else d_1921_v17_)
            nw296_[int(2)] = d_1921_v17_
            nw296_[int(3)] = d_1921_v17_
            nw296_[int(4)] = d_1921_v17_
            nw296_[int(5)] = d_1921_v17_
            nw296_[int(6)] = d_1921_v17_
            d_1938_v32_ = nw296_
            index276_ = default__.safeIndex(467, (d_1936_v30_).length(0))
            index277_ = default__.safeIndex(155, (d_1937_v31_).length(0))
            rhs274_ = (len((d_1932_v27_ if self.f24 else d_1932_v27_)) if (d_1934_i5_) != (d_1934_i5_) else d_1896_v0_)
            rhs275_ = d_1935_v29_
            rhs276_ = d_1938_v32_
            rhs277_ = (self).f26
            rhs278_ = (self).f26
            lhs235_ = d_1936_v30_
            lhs236_ = default__.safeIndex(467, (d_1936_v30_).length(0))
            lhs237_ = d_1937_v31_
            lhs238_ = default__.safeIndex(155, (d_1937_v31_).length(0))
            lhs239_ = self
            lhs240_ = globalState
            lhs235_[lhs236_] = rhs274_
            lhs237_[lhs238_] = rhs275_
            d_1938_v32_ = rhs276_
            lhs239_.f24 = rhs277_
            lhs240_.f10 = rhs278_
            d_1939_v33_: _dafny.MultiSet
            d_1939_v33_ = _dafny.MultiSet([(self).f25])
            if not(((d_1939_v33_).cardinality) <= (d_1934_i5_)):
                d_1940_v34_: _dafny.Map
                d_1940_v34_ = _dafny.Map({self.f24: d_1896_v0_})
                d_1941_v35_: T1
                nw297_ = C1()
                nw297_.ctor__((self).fm7(d_1940_v34_, globalState))
                d_1941_v35_ = nw297_
                d_1941_v35_ = (d_1941_v35_)
                d_1942_v36_: _dafny.Seq
                d_1942_v36_ = _dafny.SeqWithoutIsStrInference([d_1896_v0_, d_1896_v0_])
                d_1942_v36_ = d_1942_v36_
                d_1943_v37_: _dafny.Map
                d_1943_v37_ = _dafny.Map({(d_1934_i5_) * (d_1934_i5_): -945})
                index278_ = default__.safeIndex(11, (d_1921_v17_).length(0))
                (d_1921_v17_)[index278_] = (len(_dafny.SeqWithoutIsStrInference([len(d_1932_v27_)]))) != (d_1934_i5_)
                d_1944_v38_: D0
                d_1944_v38_ = D0_DC0(self.f24, len(d_1935_v29_), d_1896_v0_)
                d_1945_v39_: _dafny.Seq
                d_1945_v39_ = _dafny.SeqWithoutIsStrInference([d_1944_v38_])
                d_1946_v40_: _dafny.Map
                d_1946_v40_ = _dafny.Map({(self).f28: d_1945_v39_})
                d_1947_v41_: D2
                d_1947_v41_ = D2_DC3(d_1943_v37_)
                index279_ = default__.safeIndex(11, (d_1921_v17_).length(0))
                rhs279_ = (d_1934_i5_) * (len(default__.fm60((d_1941_v35_).fm5((d_1946_v40_).set((self).f25, d_1945_v39_), False, d_1941_v35_.f24, globalState), globalState)))
                rhs280_ = d_1936_v30_
                rhs281_ = (d_1947_v41_).cf8
                rhs282_ = (default__.fm0(d_1896_v0_, (self).f26, d_1933_v28_, d_1896_v0_, globalState)) < (d_1896_v0_)
                rhs283_ = d_1921_v17_
                lhs241_ = globalState
                lhs242_ = globalState
                lhs243_ = d_1921_v17_
                lhs244_ = default__.safeIndex(11, (d_1921_v17_).length(0))
                lhs245_ = globalState
                lhs241_.f11 = rhs279_
                lhs242_.f7 = rhs280_
                d_1943_v37_ = rhs281_
                lhs243_[lhs244_] = rhs282_
                lhs245_.f2 = rhs283_
                d_1948_v43_: _dafny.Array
                def lambda93_(d_1949_v0_):
                    def lambda94_(d_1950_i6_):
                        def iife129_():
                            coll71_ = _dafny.Map()
                            compr_71_: _dafny.Seq
                            for compr_71_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vq")): d_1949_v0_})).keys.Elements:
                                d_1951_v42_: _dafny.Seq = compr_71_
                                if (d_1951_v42_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vq")): d_1949_v0_})):
                                    coll71_[d_1951_v42_] = 709
                            return _dafny.Map(coll71_)
                        return iife129_()
                        

                    return lambda94_

                init54_ = lambda93_(d_1896_v0_)
                nw298_ = _dafny.Array(None, 20)
                for i0_54_ in range(nw298_.length(0)):
                    nw298_[i0_54_] = init54_(i0_54_)
                d_1948_v43_ = nw298_
                d_1952_v44_: _dafny.Map
                d_1952_v44_ = _dafny.Map({(d_1937_v31_)[default__.safeIndex(155, (d_1937_v31_).length(0))]: (d_1936_v30_)[default__.safeIndex(467, (d_1936_v30_).length(0))]})
                index280_ = default__.safeIndex(305, (d_1948_v43_).length(0))
                (d_1948_v43_)[index280_] = d_1952_v44_
                d_1953_v45_: _dafny.Set
                d_1953_v45_ = _dafny.Set({d_1898_v2_, (default__.fm33(_dafny.CodePoint('d'), (d_1936_v30_)[default__.safeIndex(467, (d_1936_v30_).length(0))], d_1941_v35_.f24, False, globalState) if (self).f25 else d_1898_v2_)})
                d_1954_v46_: _dafny.Set
                d_1954_v46_ = _dafny.Set({d_1936_v30_})
                index281_ = default__.safeIndex(305, (d_1948_v43_).length(0))
                index282_ = default__.safeIndex(11, (d_1921_v17_).length(0))
                index283_ = default__.safeIndex(467, (d_1936_v30_).length(0))
                rhs284_ = d_1936_v30_
                rhs285_ = d_1952_v44_
                rhs286_ = d_1953_v45_
                rhs287_ = (_dafny.Set({d_1936_v30_})).isdisjoint((d_1954_v46_) | (d_1954_v46_))
                rhs288_ = len(_dafny.Map({d_1940_v34_: d_1896_v0_}))
                lhs246_ = globalState
                lhs247_ = d_1948_v43_
                lhs248_ = default__.safeIndex(305, (d_1948_v43_).length(0))
                lhs249_ = d_1921_v17_
                lhs250_ = default__.safeIndex(11, (d_1921_v17_).length(0))
                lhs251_ = d_1936_v30_
                lhs252_ = default__.safeIndex(467, (d_1936_v30_).length(0))
                lhs246_.f7 = rhs284_
                lhs247_[lhs248_] = rhs285_
                d_1953_v45_ = rhs286_
                lhs249_[lhs250_] = rhs287_
                lhs251_[lhs252_] = rhs288_
                d_1955_v48_: _dafny.Array
                def lambda95_(d_1956_v2_):
                    def lambda96_(d_1957_i7_):
                        def iife130_():
                            coll72_ = _dafny.Map()
                            compr_72_: _dafny.Seq
                            for compr_72_ in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_1956_v2_]))).Elements:
                                d_1958_v47_: _dafny.Seq = compr_72_
                                if (d_1958_v47_) in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_1956_v2_]))):
                                    coll72_[d_1958_v47_] = _dafny.Map({(self).f25: self.f27})
                            return _dafny.Map(coll72_)
                        return (_dafny.Map({d_1956_v2_: _dafny.Map({(self).f26: self.f27})}) if self.f24 else iife130_()
                        )

                    return lambda96_

                init55_ = lambda95_(d_1898_v2_)
                nw299_ = _dafny.Array(None, 28)
                for i0_55_ in range(nw299_.length(0)):
                    nw299_[i0_55_] = init55_(i0_55_)
                d_1955_v48_ = nw299_
                d_1959_v49_: _dafny.Map
                d_1959_v49_ = _dafny.Map({d_1898_v2_: _dafny.Map({self.f24: self.f27})})
                index284_ = default__.safeIndex(821, (d_1955_v48_).length(0))
                (d_1955_v48_)[index284_] = d_1959_v49_
                index285_ = default__.safeIndex(821, (d_1955_v48_).length(0))
                rhs289_ = d_1959_v49_
                rhs290_ = (d_1896_v0_) + (d_1934_i5_)
                rhs291_ = len((d_1937_v31_)[default__.safeIndex(155, (d_1937_v31_).length(0))])
                rhs292_ = not(not (self.f24) or (not((d_1921_v17_)[default__.safeIndex(11, (d_1921_v17_).length(0))])))
                lhs253_ = d_1955_v48_
                lhs254_ = default__.safeIndex(821, (d_1955_v48_).length(0))
                lhs255_ = globalState
                lhs256_ = globalState
                lhs253_[lhs254_] = rhs289_
                d_1896_v0_ = rhs290_
                lhs255_.f22 = rhs291_
                lhs256_.f10 = rhs292_
            elif True:
                d_1960_v50_: _dafny.Seq
                d_1960_v50_ = _dafny.SeqWithoutIsStrInference([d_1896_v0_, (0) - ((0) - ((d_1936_v30_)[default__.safeIndex(467, (d_1936_v30_).length(0))])), (d_1936_v30_)[default__.safeIndex(467, (d_1936_v30_).length(0))], default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "clgdsdkx"))), -852), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fg")))])
                d_1960_v50_ = (d_1960_v50_) + (d_1960_v50_)
                d_1961_v51_: _dafny.Array
                def lambda97_(d_1962_i8_):
                    return _dafny.Map({(self).f25: (self).f26})

                init56_ = lambda97_
                nw300_ = _dafny.Array(None, 5)
                for i0_56_ in range(nw300_.length(0)):
                    nw300_[i0_56_] = init56_(i0_56_)
                d_1961_v51_ = nw300_
                d_1963_v52_: D8
                d_1963_v52_ = D8_DC18(d_1961_v51_)
                d_1964_v53_: _dafny.Map
                d_1964_v53_ = _dafny.Map({d_1921_v17_: d_1963_v52_})
                d_1965_v54_: _dafny.Map
                d_1965_v54_ = _dafny.Map({d_1964_v53_: self.f27})
                d_1966_v55_: D1
                d_1966_v55_ = D1_DC2(d_1936_v30_, self.f27, d_1896_v0_, (self).f25)
                d_1967_v56_: C3
                nw301_ = C3()
                nw301_.ctor__((self).f28, d_1935_v29_, (d_1939_v33_).issubset(d_1939_v33_), ((d_1965_v54_)[d_1964_v53_] if (d_1964_v53_) in (d_1965_v54_) else (d_1966_v55_).cf5), self.f24, (self).f26)
                d_1967_v56_ = nw301_
                d_1896_v0_ = d_1896_v0_
                (globalState).f16 = (d_1936_v30_)[default__.safeIndex(467, (d_1936_v30_).length(0))]
                d_1968_v57_: _dafny.Map
                d_1968_v57_ = _dafny.Map({(self).f26: d_1934_i5_})
                d_1969_v58_: C5
                nw302_ = C5()
                nw302_.ctor__((d_1967_v56_).f32, self.f27, (d_1898_v2_)[default__.safeIndex(d_1934_i5_, len(d_1898_v2_))], (self).fm7(d_1968_v57_, globalState))
                d_1969_v58_ = nw302_
                d_1970_v59_: _dafny.Seq
                d_1970_v59_ = _dafny.SeqWithoutIsStrInference([d_1969_v58_, d_1969_v58_])
                d_1971_v60_: _dafny.MultiSet
                d_1971_v60_ = _dafny.MultiSet([(d_1970_v59_) + (d_1970_v59_), (d_1970_v59_ if ((d_1933_v28_)[d_1896_v0_] if (d_1896_v0_) in (d_1933_v28_) else (d_1967_v56_).f32) else _dafny.SeqWithoutIsStrInference([d_1969_v58_]))])
                d_1896_v0_ = ((d_1971_v60_)[d_1970_v59_] if (d_1970_v59_) in (d_1971_v60_) else d_1896_v0_)
            (self).f24 = (self).f26
            if (self).f25:
                d_1972_v61_: _dafny.Map
                d_1972_v61_ = _dafny.Map({self.f27: d_1934_i5_})
                d_1972_v61_ = (d_1972_v61_).set(self.f27, d_1896_v0_)
                d_1973_v62_: _dafny.Map
                d_1973_v62_ = _dafny.Map({d_1921_v17_: d_1896_v0_})
                d_1974_v63_: _dafny.Array
                nw303_ = _dafny.Array(_dafny.MultiSet({}), 24)
                d_1974_v63_ = nw303_
                d_1975_v64_: D7
                d_1975_v64_ = D7_DC14(d_1897_v1_)
                index286_ = default__.safeIndex(467, (d_1936_v30_).length(0))
                rhs293_ = len(d_1935_v29_)
                rhs294_ = (d_1975_v64_).cf27
                rhs295_ = _dafny.Map({d_1921_v17_: (d_1936_v30_)[default__.safeIndex(467, (d_1936_v30_).length(0))]})
                rhs296_ = d_1896_v0_
                rhs297_ = d_1974_v63_
                lhs257_ = globalState
                lhs258_ = d_1936_v30_
                lhs259_ = default__.safeIndex(467, (d_1936_v30_).length(0))
                lhs257_.f22 = rhs293_
                d_1897_v1_ = rhs294_
                d_1973_v62_ = rhs295_
                lhs258_[lhs259_] = rhs296_
                d_1974_v63_ = rhs297_
                d_1976_v65_: _dafny.Map
                d_1976_v65_ = _dafny.Map({(self).f25: d_1934_i5_})
                d_1977_v66_: _dafny.MultiSet
                d_1977_v66_ = _dafny.MultiSet([d_1976_v65_, d_1976_v65_, d_1976_v65_])
                d_1978_v67_: D5
                d_1978_v67_ = D5_DC10(d_1977_v66_)
                d_1979_v68_: _dafny.MultiSet
                d_1979_v68_ = _dafny.MultiSet([d_1978_v67_, d_1978_v67_, d_1978_v67_])
                d_1980_v69_: _dafny.Map
                d_1980_v69_ = _dafny.Map({d_1936_v30_: d_1979_v68_})
                d_1980_v69_ = (d_1980_v69_).set(d_1936_v30_, _dafny.MultiSet([d_1978_v67_]))
                (globalState).f10 = (self).fm17((d_1932_v27_).isdisjoint(d_1932_v27_), (self).f25, globalState)
                (self).f24 = ((d_1936_v30_)[default__.safeIndex(467, (d_1936_v30_).length(0))]) != ((d_1936_v30_)[default__.safeIndex(467, (d_1936_v30_).length(0))])
            elif True:
                d_1981_v70_: _dafny.Array
                nw304_ = _dafny.Array(_dafny.Array(None, 0), 10)
                d_1981_v70_ = nw304_
                d_1981_v70_ = d_1981_v70_
                d_1982_v71_: _dafny.Seq
                d_1982_v71_ = _dafny.SeqWithoutIsStrInference([d_1936_v30_, d_1936_v30_])
                d_1982_v71_ = (d_1982_v71_) + ((d_1982_v71_) + (d_1982_v71_))
                index287_ = default__.safeIndex(991, (d_1921_v17_).length(0))
                (d_1921_v17_)[index287_] = (self).f25
                d_1983_v72_: _dafny.Seq
                d_1983_v72_ = _dafny.SeqWithoutIsStrInference([d_1921_v17_])
                d_1984_v73_: _dafny.Map
                d_1984_v73_ = _dafny.Map({(d_1983_v72_)[default__.safeIndex(d_1896_v0_, len(d_1983_v72_))]: (self).f28})
                index288_ = default__.safeIndex(991, (d_1921_v17_).length(0))
                (d_1921_v17_)[index288_] = ((d_1984_v73_)[d_1921_v17_] if (d_1921_v17_) in (d_1984_v73_) else (self).f28)
                d_1985_v74_: D0
                d_1985_v74_ = D0_DC0((d_1921_v17_)[default__.safeIndex(991, (d_1921_v17_).length(0))], 942, d_1896_v0_)
                d_1986_v75_: D0
                d_1986_v75_ = D0_DC0(not((self).f25), (self).fm6(d_1985_v74_, globalState), d_1934_i5_)
                d_1987_v76_: _dafny.Map
                d_1987_v76_ = _dafny.Map({not ((self).f26) or ((self).f25): d_1986_v75_})
                d_1988_v77_: _dafny.Seq
                d_1988_v77_ = _dafny.SeqWithoutIsStrInference([d_1896_v0_, d_1934_i5_])
                d_1989_v78_: _dafny.Map
                d_1989_v78_ = _dafny.Map({(self).f25: d_1896_v0_})
                d_1990_v79_: _dafny.Seq
                d_1990_v79_ = _dafny.SeqWithoutIsStrInference([(d_1989_v78_).set((self).f25, d_1934_i5_)])
                pat_let_tv19_ = d_1936_v30_
                pat_let_tv20_ = d_1936_v30_
                pat_let_tv21_ = d_1988_v77_
                pat_let_tv22_ = d_1990_v79_
                pat_let_tv23_ = d_1936_v30_
                pat_let_tv24_ = d_1936_v30_
                pat_let_tv25_ = d_1990_v79_
                pat_let_tv26_ = globalState
                def iife132_(_pat_let30_0):
                    def iife133_(d_1991_dt__update__tmp_h0_):
                        def iife134_(_pat_let31_0):
                            def iife135_(d_1992_dt__update_hcf2_h0_):
                                return D0_DC0((d_1991_dt__update__tmp_h0_).cf0, (d_1991_dt__update__tmp_h0_).cf1, d_1992_dt__update_hcf2_h0_)
                            return iife135_(_pat_let31_0)
                        return iife134_(25)
                    return iife133_(_pat_let30_0)
                def iife131_(_pat_let29_0):
                    def iife136_(d_1993_dt__update__tmp_h1_):
                        def iife137_(_pat_let32_0):
                            def iife138_(d_1994_dt__update_hcf1_h0_):
                                return D0_DC0((d_1993_dt__update__tmp_h1_).cf0, d_1994_dt__update_hcf1_h0_, (d_1993_dt__update__tmp_h1_).cf2)
                            return iife138_(_pat_let32_0)
                        return iife137_(len(default__.fm38((self).f25, (pat_let_tv20_)[default__.safeIndex(467, (pat_let_tv19_).length(0))], _dafny.MultiSet(pat_let_tv21_), (pat_let_tv22_)[default__.safeIndex((pat_let_tv24_)[default__.safeIndex(467, (pat_let_tv23_).length(0))], len(pat_let_tv25_))], pat_let_tv26_)))
                    return iife136_(_pat_let29_0)
                d_1987_v76_ = (d_1987_v76_).set((self).f25, (iife131_(iife132_(d_1986_v75_)) if self.f24 else d_1985_v74_))
                d_1995_v80_: _dafny.Array
                nw305_ = _dafny.Array(_dafny.Seq({}), 7)
                d_1995_v80_ = nw305_
                index289_ = default__.safeIndex(692, (d_1995_v80_).length(0))
                (d_1995_v80_)[index289_] = (d_1983_v72_) + (d_1983_v72_)
                index290_ = default__.safeIndex(692, (d_1995_v80_).length(0))
                (d_1995_v80_)[index290_] = d_1983_v72_
        d_1996_v81_: _dafny.MultiSet
        d_1996_v81_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vyu")))])
        index291_ = default__.safeIndex(590, (d_1921_v17_).length(0))
        (d_1921_v17_)[index291_] = (d_1996_v81_).issubset(d_1996_v81_)
        d_1997_v82_: _dafny.Map
        d_1997_v82_ = _dafny.Map({d_1896_v0_: d_1896_v0_})
        d_1998_v83_: _dafny.Seq
        d_1998_v83_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vtpguopsq"))
        d_1999_v84_: _dafny.Seq
        d_1999_v84_ = _dafny.SeqWithoutIsStrInference([((d_1997_v82_)[(0) - (len(d_1998_v83_))] if ((0) - (len(d_1998_v83_))) in (d_1997_v82_) else d_1896_v0_), d_1896_v0_, (0) - (d_1896_v0_), d_1896_v0_, ((d_1996_v81_)[d_1896_v0_] if (d_1896_v0_) in (d_1996_v81_) else d_1896_v0_)])
        d_2000_v85_: _dafny.Set
        d_2000_v85_ = _dafny.Set({(d_1996_v81_).intersection(_dafny.MultiSet(d_1999_v84_)), d_1996_v81_})
        index292_ = default__.safeIndex(590, (d_1921_v17_).length(0))
        rhs298_ = d_1896_v0_
        rhs299_ = default__.safeDivisionInt(d_1896_v0_, (d_1896_v0_) * (d_1896_v0_))
        rhs300_ = (self).f25
        rhs301_ = (_dafny.Set({d_1996_v81_})).intersection(d_2000_v85_)
        lhs260_ = globalState
        lhs261_ = globalState
        lhs262_ = d_1921_v17_
        lhs263_ = default__.safeIndex(590, (d_1921_v17_).length(0))
        lhs260_.f16 = rhs298_
        lhs261_.f16 = rhs299_
        lhs262_[lhs263_] = rhs300_
        d_2000_v85_ = rhs301_
        r0 = (self).f25
        r1 = d_1921_v17_
        return r0, r1

    def m1(self, p0, globalState):
        r0: str = _dafny.CodePoint('D')
        r1: _dafny.Seq = _dafny.Seq({})
        d_2001_v0_: C1
        nw306_ = C1()
        nw306_.ctor__((self).f25)
        d_2001_v0_ = nw306_
        d_2002_v1_: int
        d_2002_v1_ = 472
        d_2003_v2_: _dafny.Seq
        d_2003_v2_ = _dafny.SeqWithoutIsStrInference([d_2002_v1_])
        hi13_ = d_2002_v1_
        for d_2004_i0_ in range((0) - ((0) - ((d_2003_v2_)[default__.safeIndex(d_2002_v1_, len(d_2003_v2_))])), hi13_):
            d_2005_v3_: _dafny.Map
            d_2005_v3_ = _dafny.Map({d_2002_v1_: d_2003_v2_})
            rhs302_ = d_2005_v3_
            rhs303_ = (self).f26
            lhs264_ = self
            d_2005_v3_ = rhs302_
            lhs264_.f24 = rhs303_
            (self).f24 = not((self).f25)
            (globalState).f10 = self.f24
            d_2006_v4_: _dafny.Seq
            d_2006_v4_ = _dafny.SeqWithoutIsStrInference([self.f24, (self).f28, (self).f26])
            d_2007_v5_: _dafny.Seq
            d_2007_v5_ = _dafny.SeqWithoutIsStrInference([(d_2006_v4_)[default__.safeIndex(d_2004_i0_, len(d_2006_v4_))]])
            d_2008_v6_: D10
            d_2008_v6_ = D10_DC23((d_2007_v5_ if not((self).f28) else d_2006_v4_), self.f24)
            source35_ = d_2008_v6_
            if source35_.is_DC22:
                d_2009___mcc_h0_ = source35_.cf41
                d_2010___mcc_h1_ = source35_.cf42
                d_2011_cf42_ = d_2010___mcc_h1_
                d_2012_cf41_ = d_2009___mcc_h0_
                d_2011_cf42_ = (self.f24) and (d_2011_cf42_)
                d_2013_v7_: _dafny.Array
                nw307_ = _dafny.Array(_dafny.Map({}), 17)
                d_2013_v7_ = nw307_
                d_2014_v8_: _dafny.Array
                nw308_ = _dafny.Array(_dafny.Array(None, 0), 1)
                d_2014_v8_ = nw308_
                d_2015_v9_: _dafny.Array
                d_2015_v9_ = d_2014_v8_
                d_2016_v10_: _dafny.Map
                d_2016_v10_ = _dafny.Map({d_2015_v9_: self.f27})
                index293_ = default__.safeIndex(100, (d_2013_v7_).length(0))
                (d_2013_v7_)[index293_] = d_2016_v10_
                index294_ = default__.safeIndex(100, (d_2013_v7_).length(0))
                (d_2013_v7_)[index294_] = d_2016_v10_
                (globalState).f16 = d_2012_cf41_
                d_2017_v11_: _dafny.Seq
                d_2017_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rsvrarqnn"))
                d_2018_v12_: _dafny.Map
                d_2018_v12_ = _dafny.Map({d_2002_v1_: d_2017_v11_})
                d_2019_v13_: C0
                nw309_ = C0()
                nw309_.ctor__(((d_2018_v12_)[d_2004_i0_] if (d_2004_i0_) in (d_2018_v12_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "slk"))))
                d_2019_v13_ = nw309_
            elif source35_.is_DC23:
                d_2020___mcc_h2_ = source35_.cf43
                d_2021___mcc_h3_ = source35_.cf44
                d_2022_cf44_ = d_2021___mcc_h3_
                d_2023_cf43_ = d_2020___mcc_h2_
                d_2024_v14_: _dafny.Seq
                d_2024_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))
                d_2025_v15_: _dafny.Set
                d_2025_v15_ = _dafny.Set({self.f27})
                d_2026_v16_: C2
                nw310_ = C2()
                nw310_.ctor__(default__.fm41(not((self).f26), (d_2024_v14_)[default__.safeIndex(len(d_2025_v15_), len(d_2024_v14_))], (d_2006_v4_)[default__.safeIndex((_dafny.MultiSet([d_2002_v1_])).cardinality, len(d_2006_v4_))], globalState), d_2024_v14_)
                d_2026_v16_ = nw310_
                d_2027_v17_: _dafny.Array
                def lambda98_(d_2028_v1_, d_2029_cf44_):
                    def lambda99_(d_2030_i1_):
                        return (D10_DC22(d_2028_v1_, d_2029_cf44_)).cf42

                    return lambda99_

                init57_ = lambda98_(d_2002_v1_, d_2022_cf44_)
                nw311_ = _dafny.Array(None, 15)
                for i0_57_ in range(nw311_.length(0)):
                    nw311_[i0_57_] = init57_(i0_57_)
                d_2027_v17_ = nw311_
                d_2031_v18_: _dafny.Map
                d_2031_v18_ = _dafny.Map({d_2004_i0_: d_2027_v17_})
                d_2031_v18_ = ((d_2031_v18_ if self.f24 else d_2031_v18_)) | ((d_2031_v18_) | (d_2031_v18_))
                d_2032_v19_: _dafny.Array
                def lambda100_(d_2033_i2_):
                    return self.f27

                init58_ = lambda100_
                nw312_ = _dafny.Array(None, 27)
                for i0_58_ in range(nw312_.length(0)):
                    nw312_[i0_58_] = init58_(i0_58_)
                d_2032_v19_ = nw312_
                index295_ = default__.safeIndex(801, (d_2032_v19_).length(0))
                (d_2032_v19_)[index295_] = _dafny.CodePoint('i')
                index296_ = default__.safeIndex(801, (d_2032_v19_).length(0))
                (d_2032_v19_)[index296_] = self.f27
                d_2034_v20_: C2
                nw313_ = C2()
                nw313_.ctor__(d_2026_v16_.f29, ((d_2026_v16_).f30) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "onbgsyaww"))))
                d_2034_v20_ = nw313_
                d_2035_v21_: _dafny.Set
                d_2035_v21_ = _dafny.Set({(self).f25, d_2022_cf44_, d_2022_cf44_})
                d_2036_v22_: _dafny.Seq
                d_2036_v22_ = _dafny.SeqWithoutIsStrInference([d_2026_v16_])
                rhs304_ = (d_2035_v21_).isdisjoint(_dafny.Set({(self).f25, d_2022_cf44_, (self).f25, (self).f28, (self).f25}))
                rhs305_ = (default__.fm61((self).f25, d_2004_i0_, default__.fm42(globalState), (self).f28, globalState)).cf49
                rhs306_ = ((d_2036_v22_)[default__.safeIndex(d_2002_v1_, len(d_2036_v22_))] if (self).f26 else d_2026_v16_)
                lhs265_ = globalState
                lhs266_ = globalState
                lhs265_.f10 = rhs304_
                lhs266_.f11 = rhs305_
                d_2034_v20_ = rhs306_
            elif source35_.is_DC21:
                d_2037___mcc_h4_ = source35_.cf40
                d_2038_cf40_ = d_2037___mcc_h4_
                d_2039_v23_: _dafny.Map
                d_2039_v23_ = _dafny.Map({True: not((self).f28)})
                d_2040_v25_: _dafny.Map
                d_2040_v25_ = _dafny.Map({d_2004_i0_: not((self).f28)})
                d_2041_v26_: _dafny.MultiSet
                d_2041_v26_ = _dafny.MultiSet([d_2002_v1_, d_2002_v1_])
                d_2042_v27_: _dafny.Array
                nw314_ = _dafny.Array(None, 20)
                nw314_[int(0)] = (self).f26
                nw314_[int(1)] = (self).f28
                nw314_[int(2)] = self.f24
                nw314_[int(3)] = self.f24
                nw314_[int(4)] = (self).f26
                nw314_[int(5)] = (self).f25
                nw314_[int(6)] = (self).f25
                nw314_[int(7)] = (self).f26
                nw314_[int(8)] = not(self.f24)
                nw314_[int(9)] = self.f24
                nw314_[int(10)] = (self).f25
                nw314_[int(11)] = self.f24
                nw314_[int(12)] = (self).f28
                nw314_[int(13)] = (self).f26
                nw314_[int(14)] = False
                nw314_[int(15)] = self.f24
                nw314_[int(16)] = (self).f28
                nw314_[int(17)] = (self).f28
                nw314_[int(18)] = not((self).f28)
                nw314_[int(19)] = not((self).f26)
                d_2042_v27_ = nw314_
                d_2043_v28_: D6
                d_2043_v28_ = D6_DC13(d_2042_v27_, d_2004_i0_, False)
                d_2044_v29_: _dafny.Array
                nw315_ = _dafny.Array(None, 19)
                nw315_[int(0)] = len(d_2039_v23_)
                nw315_[int(1)] = (len(default__.fm53(self.f27, self.f27, globalState))) - (len(_dafny.Map({d_2004_i0_: (self).f25})))
                nw315_[int(2)] = d_2004_i0_
                def iife139_():
                    coll73_ = _dafny.Set()
                    compr_73_: int
                    for compr_73_ in _dafny.IntegerRange(643, -926):
                        d_2045_v24_: int = compr_73_
                        if ((643) <= (d_2045_v24_)) and ((d_2045_v24_) < (-926)):
                            coll73_ = coll73_.union(_dafny.Set([(d_2045_v24_) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bvaem"))))]))
                    return _dafny.Set(coll73_)
                nw315_[int(3)] = len(iife139_()
                )
                nw315_[int(4)] = default__.safeModuloInt((0) - (d_2004_i0_), d_2004_i0_)
                nw315_[int(5)] = d_2002_v1_
                nw315_[int(6)] = d_2002_v1_
                nw315_[int(7)] = ((0) - (d_2002_v1_)) * (d_2002_v1_)
                nw315_[int(8)] = len(d_2040_v25_)
                nw315_[int(9)] = d_2002_v1_
                nw315_[int(10)] = d_2004_i0_
                nw315_[int(11)] = d_2002_v1_
                nw315_[int(12)] = (((d_2041_v26_)[d_2004_i0_] if (d_2004_i0_) in (d_2041_v26_) else (d_2043_v28_).cf25)) - (d_2004_i0_)
                nw315_[int(13)] = d_2004_i0_
                nw315_[int(14)] = (0) - (d_2002_v1_)
                nw315_[int(15)] = (745) - ((0) - (len(d_2039_v23_)))
                nw315_[int(16)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lnuul")))
                nw315_[int(17)] = d_2002_v1_
                nw315_[int(18)] = default__.safeDivisionInt(d_2002_v1_, d_2004_i0_)
                d_2044_v29_ = nw315_
                index297_ = default__.safeIndex(561, (d_2044_v29_).length(0))
                (d_2044_v29_)[index297_] = 841
                index298_ = default__.safeIndex(561, (d_2044_v29_).length(0))
                (d_2044_v29_)[index298_] = d_2002_v1_
                d_2046_v30_: _dafny.Set
                d_2046_v30_ = _dafny.Set({(self).f25, self.f24})
                d_2047_v31_: _dafny.Seq
                d_2047_v31_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eplk"))
                d_2048_v32_: C3
                nw316_ = C3()
                nw316_.ctor__((d_2046_v30_).ispropersubset(d_2046_v30_), d_2047_v31_, self.f24, self.f27, (self).f25, (self).f25)
                d_2048_v32_ = nw316_
                d_2049_v33_: _dafny.Set
                d_2049_v33_ = _dafny.Set({_dafny.Map({(_dafny.MultiSet(d_2047_v31_)).cardinality: (self).f26}), d_2040_v25_, d_2040_v25_, d_2040_v25_, d_2040_v25_})
                rhs307_ = (d_2004_i0_) > (len(d_2049_v33_))
                rhs308_ = d_2048_v32_
                lhs267_ = globalState
                lhs267_.f10 = rhs307_
                d_2048_v32_ = rhs308_
                (globalState).f10 = (d_2048_v32_).f32
                index299_ = default__.safeIndex(715, (d_2042_v27_).length(0))
                (d_2042_v27_)[index299_] = self.f24
                d_2050_v34_: _dafny.Seq
                d_2050_v34_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([852, d_2004_i0_, len(p0), (d_2044_v29_)[default__.safeIndex(561, (d_2044_v29_).length(0))]]), (d_2003_v2_).set(default__.safeIndex((d_2044_v29_)[default__.safeIndex(561, (d_2044_v29_).length(0))], len(d_2003_v2_)), len(default__.fm41((d_2048_v32_).f32, self.f27, (self).f25, globalState)))])
                d_2051_v35_: _dafny.Seq
                d_2051_v35_ = _dafny.SeqWithoutIsStrInference([d_2050_v34_])
                d_2052_v36_: _dafny.Array
                nw317_ = _dafny.Array(None, 9)
                nw317_[int(0)] = d_2050_v34_
                nw317_[int(1)] = (d_2051_v35_)[default__.safeIndex((d_2044_v29_)[default__.safeIndex(561, (d_2044_v29_).length(0))], len(d_2051_v35_))]
                nw317_[int(2)] = (d_2050_v34_) + (d_2050_v34_)
                nw317_[int(3)] = d_2050_v34_
                nw317_[int(4)] = d_2050_v34_
                nw317_[int(5)] = d_2050_v34_
                nw317_[int(6)] = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(d_2044_v29_)[default__.safeIndex(561, (d_2044_v29_).length(0))], d_2004_i0_])])
                nw317_[int(7)] = d_2050_v34_
                nw317_[int(8)] = d_2050_v34_
                d_2052_v36_ = nw317_
                index300_ = default__.safeIndex(483, (d_2052_v36_).length(0))
                (d_2052_v36_)[index300_] = (d_2050_v34_) + (d_2050_v34_)
                d_2053_v37_: _dafny.Seq
                d_2053_v37_ = _dafny.SeqWithoutIsStrInference([d_2044_v29_, d_2044_v29_])
                d_2054_v38_: _dafny.Array
                nw318_ = _dafny.Array(None, 15)
                nw318_[int(0)] = d_2044_v29_
                nw318_[int(1)] = d_2044_v29_
                nw318_[int(2)] = d_2044_v29_
                nw318_[int(3)] = d_2044_v29_
                nw318_[int(4)] = d_2044_v29_
                nw318_[int(5)] = (d_2053_v37_)[default__.safeIndex((d_2044_v29_)[default__.safeIndex(561, (d_2044_v29_).length(0))], len(d_2053_v37_))]
                nw318_[int(6)] = d_2044_v29_
                nw318_[int(7)] = (d_2053_v37_)[default__.safeIndex(len(d_2047_v31_), len(d_2053_v37_))]
                nw318_[int(8)] = d_2044_v29_
                nw318_[int(9)] = d_2044_v29_
                nw318_[int(10)] = d_2044_v29_
                nw318_[int(11)] = d_2044_v29_
                nw318_[int(12)] = d_2044_v29_
                nw318_[int(13)] = d_2044_v29_
                nw318_[int(14)] = d_2044_v29_
                d_2054_v38_ = nw318_
                index301_ = default__.safeIndex(84, (d_2054_v38_).length(0))
                (d_2054_v38_)[index301_] = d_2044_v29_
                index302_ = default__.safeIndex(715, (d_2042_v27_).length(0))
                index303_ = default__.safeIndex(483, (d_2052_v36_).length(0))
                index304_ = default__.safeIndex(84, (d_2054_v38_).length(0))
                rhs309_ = self.f24
                rhs310_ = d_2050_v34_
                rhs311_ = d_2044_v29_
                rhs312_ = d_2044_v29_
                lhs268_ = d_2042_v27_
                lhs269_ = default__.safeIndex(715, (d_2042_v27_).length(0))
                lhs270_ = d_2052_v36_
                lhs271_ = default__.safeIndex(483, (d_2052_v36_).length(0))
                lhs272_ = d_2054_v38_
                lhs273_ = default__.safeIndex(84, (d_2054_v38_).length(0))
                lhs268_[lhs269_] = rhs309_
                lhs270_[lhs271_] = rhs310_
                lhs272_[lhs273_] = rhs311_
                d_2044_v29_ = rhs312_
            elif True:
                d_2055___mcc_h5_ = source35_.cf45
                d_2056_cf45_ = d_2055___mcc_h5_
                d_2057_v39_: _dafny.Array
                def lambda101_(d_2058_i3_):
                    return self.f24

                init59_ = lambda101_
                nw319_ = _dafny.Array(None, 7)
                for i0_59_ in range(nw319_.length(0)):
                    nw319_[i0_59_] = init59_(i0_59_)
                d_2057_v39_ = nw319_
                index305_ = default__.safeIndex(249, (d_2057_v39_).length(0))
                (d_2057_v39_)[index305_] = not (self.f24) or ((self).f25)
                index306_ = default__.safeIndex(249, (d_2057_v39_).length(0))
                (d_2057_v39_)[index306_] = not((self).f26)
                d_2059_v40_: _dafny.Map
                d_2059_v40_ = _dafny.Map({self.f27: d_2003_v2_})
                d_2059_v40_ = (d_2059_v40_) | ((d_2059_v40_) | (d_2059_v40_))
                d_2060_v41_: _dafny.Seq
                d_2060_v41_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ohc"))
                d_2061_v42_: D11
                d_2061_v42_ = D11_DC27()
                d_2062_v43_: _dafny.Map
                d_2062_v43_ = _dafny.Map({d_2060_v41_: d_2061_v42_})
                d_2062_v43_ = (d_2062_v43_).set(d_2060_v41_, d_2061_v42_)
                d_2063_v44_: _dafny.Array
                nw320_ = _dafny.Array(int(0), 20)
                d_2063_v44_ = nw320_
                d_2064_v45_: _dafny.Map
                d_2064_v45_ = _dafny.Map({d_2006_v4_: (d_2057_v39_)[default__.safeIndex(249, (d_2057_v39_).length(0))]})
                d_2065_v46_: _dafny.MultiSet
                d_2065_v46_ = _dafny.MultiSet([(d_2057_v39_)[default__.safeIndex(249, (d_2057_v39_).length(0))], self.f24, (self).f25, (d_2057_v39_)[default__.safeIndex(249, (d_2057_v39_).length(0))], self.f24])
                d_2066_v47_: _dafny.Map
                d_2066_v47_ = _dafny.Map({((d_2064_v45_)[d_2006_v4_] if (d_2006_v4_) in (d_2064_v45_) else False): d_2065_v46_})
                index307_ = default__.safeIndex(78, (d_2063_v44_).length(0))
                (d_2063_v44_)[index307_] = len(d_2066_v47_)
                index308_ = default__.safeIndex(78, (d_2063_v44_).length(0))
                (d_2063_v44_)[index308_] = default__.safeModuloInt(d_2004_i0_, d_2004_i0_)
        d_2002_v1_ = d_2002_v1_
        if (self).f28:
            d_2067_v48_: _dafny.Array
            def lambda102_(d_2068_i4_):
                return (self).f28

            init60_ = lambda102_
            nw321_ = _dafny.Array(None, 5)
            for i0_60_ in range(nw321_.length(0)):
                nw321_[i0_60_] = init60_(i0_60_)
            d_2067_v48_ = nw321_
            index309_ = default__.safeIndex(436, (d_2067_v48_).length(0))
            (d_2067_v48_)[index309_] = (self).f28
            d_2069_v49_: _dafny.Map
            d_2069_v49_ = _dafny.Map({self.f24: d_2002_v1_})
            d_2070_v50_: _dafny.Seq
            d_2070_v50_ = _dafny.SeqWithoutIsStrInference([d_2069_v49_, d_2069_v49_, d_2069_v49_])
            index310_ = default__.safeIndex(436, (d_2067_v48_).length(0))
            (d_2067_v48_)[index310_] = (default__.fm1((self).f25, (self).f26, len((d_2070_v50_)[default__.safeIndex(d_2002_v1_, len(d_2070_v50_))]), globalState) if (self).f25 else self.f24)
            d_2067_v48_ = d_2067_v48_
            d_2071_v51_: _dafny.Seq
            d_2071_v51_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jearwwh"))
            d_2071_v51_ = (d_2071_v51_) + (d_2071_v51_)
            d_2072_v52_: T1
            nw322_ = C1()
            nw322_.ctor__(True)
            d_2072_v52_ = nw322_
            d_2073_v53_: D2
            d_2073_v53_ = D2_DC5(self.f27, True, len(_dafny.Map({self.f27: len(_dafny.Map({d_2072_v52_: (self).f28}))})), len(d_2071_v51_))
            (globalState).f11 = (len(d_2071_v51_)) + ((d_2073_v53_).cf12)
            d_2074_v54_: _dafny.Map
            d_2074_v54_ = _dafny.Map({(d_2002_v1_ if (d_2067_v48_)[default__.safeIndex(436, (d_2067_v48_).length(0))] else 391): d_2002_v1_})
            d_2075_v55_: D6
            d_2075_v55_ = D6_DC12(d_2003_v2_)
            d_2076_v56_: D0
            d_2076_v56_ = D0_DC0(self.f24, d_2002_v1_, (_dafny.MultiSet((d_2075_v55_).cf23)).cardinality)
            pat_let_tv27_ = d_2002_v1_
            d_2077_v57_: _dafny.Seq
            def iife140_(_pat_let33_0):
                def iife141_(d_2078_dt__update__tmp_h0_):
                    def iife142_(_pat_let34_0):
                        def iife143_(d_2079_dt__update_hcf1_h0_):
                            return D0_DC0((d_2078_dt__update__tmp_h0_).cf0, d_2079_dt__update_hcf1_h0_, (d_2078_dt__update__tmp_h0_).cf2)
                        return iife143_(_pat_let34_0)
                    return iife142_(pat_let_tv27_)
                return iife141_(_pat_let33_0)
            d_2077_v57_ = _dafny.SeqWithoutIsStrInference([iife140_(d_2076_v56_)])
            d_2080_v58_: _dafny.Map
            d_2080_v58_ = _dafny.Map({(self).f26: d_2077_v57_})
            def iife144_():
                coll74_ = _dafny.Map()
                compr_74_: int
                for compr_74_ in (_dafny.SeqWithoutIsStrInference([d_2002_v1_])).Elements:
                    d_2081_v59_: int = compr_74_
                    if (d_2081_v59_) in (_dafny.SeqWithoutIsStrInference([d_2002_v1_])):
                        coll74_[(d_2081_v59_) + (950)] = len(_dafny.SeqWithoutIsStrInference([(D7_DC15((d_2071_v51_)[default__.safeIndex(-828, len(d_2071_v51_))], d_2072_v52_.f24)).cf28 for d_2082_i5_ in range(default__.abs(841))]))
                return _dafny.Map(coll74_)
            d_2074_v54_ = (d_2074_v54_).set((default__.fm62((0) - (d_2002_v1_), (self).fm5(d_2080_v58_, True, (self).f25, globalState), iife144_()
            , globalState)).cf31, d_2002_v1_)
        elif True:
            d_2083_v60_: C4
            nw323_ = C4()
            nw323_.ctor__((self).f28, self.f24)
            d_2083_v60_ = nw323_
            d_2084_v61_: _dafny.MultiSet
            d_2084_v61_ = _dafny.MultiSet([d_2083_v60_])
            d_2085_v62_: _dafny.Seq
            d_2085_v62_ = _dafny.SeqWithoutIsStrInference([(self).f25])
            r1 = (d_2085_v62_ if (d_2084_v61_).ispropersubset(d_2084_v61_) else d_2085_v62_)
            d_2086_v63_: _dafny.Seq
            d_2086_v63_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kywsexvtu"))
            d_2087_v64_: C2
            nw324_ = C2()
            nw324_.ctor__(d_2086_v63_, d_2086_v63_)
            d_2087_v64_ = nw324_
            d_2088_v65_: D11
            d_2088_v65_ = D11_DC26((d_2087_v64_).fm19(self.f27, globalState), (self).f28, d_2002_v1_, not(self.f24))
            d_2089_v66_: _dafny.Map
            d_2089_v66_ = _dafny.Map({(d_2088_v65_ if self.f24 else d_2088_v65_): default__.safeDivisionInt((0) - ((d_2003_v2_)[default__.safeIndex(d_2002_v1_, len(d_2003_v2_))]), d_2002_v1_)})
            d_2089_v66_ = d_2089_v66_
            (globalState).f22 = 651
            (globalState).f10 = ((self).f28) or (self.f24)
        d_2090_v67_: _dafny.Seq
        d_2090_v67_ = _dafny.SeqWithoutIsStrInference([(self).f26, (self).f25])
        d_2091_v68_: _dafny.Map
        d_2091_v68_ = _dafny.Map({False: d_2002_v1_})
        d_2092_v69_: _dafny.Seq
        d_2092_v69_ = _dafny.SeqWithoutIsStrInference([d_2091_v68_])
        d_2093_v70_: C3
        nw325_ = C3()
        nw325_.ctor__((d_2090_v67_)[default__.safeIndex((0) - (d_2002_v1_), len(d_2090_v67_))], default__.fm38(self.f24, d_2002_v1_, default__.fm52(self.f24, globalState), (d_2092_v69_)[default__.safeIndex(d_2002_v1_, len(d_2092_v69_))], globalState), (self).f26, default__.fm47(globalState), (self).f26, (self).f26)
        d_2093_v70_ = nw325_
        d_2094_v71_: _dafny.Array
        nw326_ = _dafny.Array(None, 25)
        nw326_[int(0)] = (self).f25
        nw326_[int(1)] = (self).f28
        nw326_[int(2)] = (self).f28
        nw326_[int(3)] = (self).f26
        nw326_[int(4)] = (self).f28
        nw326_[int(5)] = self.f24
        nw326_[int(6)] = self.f24
        nw326_[int(7)] = (self).f25
        nw326_[int(8)] = (d_2093_v70_).f32
        nw326_[int(9)] = not(self.f24)
        nw326_[int(10)] = self.f24
        nw326_[int(11)] = (self).f26
        nw326_[int(12)] = (self).f28
        nw326_[int(13)] = self.f24
        nw326_[int(14)] = not((self).f28)
        nw326_[int(15)] = (d_2093_v70_).f32
        nw326_[int(16)] = (self).f25
        nw326_[int(17)] = True
        nw326_[int(18)] = (self).f28
        nw326_[int(19)] = (self).f26
        nw326_[int(20)] = (self).f26
        nw326_[int(21)] = self.f24
        nw326_[int(22)] = (self).f25
        nw326_[int(23)] = (self).f25
        nw326_[int(24)] = False
        d_2094_v71_ = nw326_
        d_2095_v72_: D6
        d_2095_v72_ = D6_DC13(d_2094_v71_, d_2002_v1_, (self).f25)
        d_2096_v73_: D7
        d_2096_v73_ = D7_DC16((d_2093_v70_).f32, d_2002_v1_, (d_2095_v72_).cf26, (self).f25, False)
        d_2097_i6_: int
        d_2097_i6_ = 0
        with _dafny.label("18"):
            pat_let_tv31_ = d_2093_v70_
            pat_let_tv32_ = d_2093_v70_
            def lambda105_(source36_):
                if source36_.is_DC15:
                    d_2131___mcc_h6_ = source36_.cf28
                    d_2132___mcc_h7_ = source36_.cf29
                    d_2133_cf29_ = d_2132___mcc_h7_
                    d_2134_cf28_ = d_2131___mcc_h6_
                    return False
                elif source36_.is_DC16:
                    d_2135___mcc_h8_ = source36_.cf30
                    d_2136___mcc_h9_ = source36_.cf31
                    d_2137___mcc_h10_ = source36_.cf32
                    d_2138___mcc_h11_ = source36_.cf33
                    d_2139___mcc_h12_ = source36_.cf34
                    d_2140_cf34_ = d_2139___mcc_h12_
                    d_2141_cf33_ = d_2138___mcc_h11_
                    d_2142_cf32_ = d_2137___mcc_h10_
                    d_2143_cf31_ = d_2136___mcc_h9_
                    d_2144_cf30_ = d_2135___mcc_h8_
                    return d_2141_cf33_
                elif source36_.is_DC14:
                    d_2145___mcc_h13_ = source36_.cf27
                    d_2146_cf27_ = d_2145___mcc_h13_
                    return (pat_let_tv31_).f32
                elif True:
                    d_2147___mcc_h14_ = source36_.cf35
                    d_2148_cf35_ = d_2147___mcc_h14_
                    return (pat_let_tv32_).f32

            while lambda105_(d_2096_v73_):
                with _dafny.c_label("18"):
                    if (d_2097_i6_) >= (100):
                        raise _dafny.Break("18")
                    d_2097_i6_ = (d_2097_i6_) + (1)
                    if (self).f25:
                        d_2098_v74_: _dafny.MultiSet
                        d_2098_v74_ = _dafny.MultiSet([d_2002_v1_])
                        d_2099_v75_: _dafny.Map
                        d_2099_v75_ = _dafny.Map({self.f24: (self).f26})
                        d_2100_v76_: D2
                        d_2100_v76_ = D2_DC5(self.f27, (d_2093_v70_).f32, d_2002_v1_, (0) - (d_2002_v1_))
                        d_2101_v77_: _dafny.Array
                        nw327_ = _dafny.Array(None, 18)
                        nw327_[int(0)] = _dafny.Map({self.f24: True})
                        nw327_[int(1)] = (d_2099_v75_).set(True, (d_2093_v70_).f32)
                        nw327_[int(2)] = d_2099_v75_
                        nw327_[int(3)] = _dafny.Map({(d_2100_v76_).cf11: (d_2093_v70_).f32})
                        nw327_[int(4)] = d_2099_v75_
                        nw327_[int(5)] = _dafny.Map({self.f24: (self).f26})
                        nw327_[int(6)] = _dafny.Map({(d_2093_v70_).f32: (d_2093_v70_).f32})
                        nw327_[int(7)] = d_2099_v75_
                        nw327_[int(8)] = d_2099_v75_
                        nw327_[int(9)] = d_2099_v75_
                        nw327_[int(10)] = d_2099_v75_
                        nw327_[int(11)] = d_2099_v75_
                        nw327_[int(12)] = d_2099_v75_
                        nw327_[int(13)] = (d_2093_v70_).fm40(False, d_2003_v2_, globalState)
                        nw327_[int(14)] = d_2099_v75_
                        nw327_[int(15)] = d_2099_v75_
                        nw327_[int(16)] = d_2099_v75_
                        nw327_[int(17)] = d_2099_v75_
                        d_2101_v77_ = nw327_
                        d_2102_v78_: _dafny.Map
                        d_2102_v78_ = _dafny.Map({(d_2098_v74_).cardinality: D8_DC18(d_2101_v77_)})
                        d_2103_v79_: D8
                        d_2103_v79_ = D8_DC18(d_2101_v77_)
                        d_2104_v80_: _dafny.Map
                        d_2104_v80_ = _dafny.Map({self.f27: d_2101_v77_})
                        pat_let_tv28_ = d_2104_v80_
                        pat_let_tv29_ = d_2104_v80_
                        pat_let_tv30_ = d_2101_v77_
                        def iife145_(_pat_let35_0):
                            def iife146_(d_2105_dt__update__tmp_h1_):
                                def iife147_(_pat_let36_0):
                                    def iife148_(d_2106_dt__update_hcf36_h0_):
                                        return D8_DC18(d_2106_dt__update_hcf36_h0_)
                                    return iife148_(_pat_let36_0)
                                return iife147_(((pat_let_tv28_)[self.f27] if (self.f27) in (pat_let_tv29_) else pat_let_tv30_))
                            return iife146_(_pat_let35_0)
                        d_2102_v78_ = (d_2102_v78_).set(d_2002_v1_, iife145_(d_2103_v79_))
                        d_2107_v81_: D3
                        d_2107_v81_ = D3_DC8((self).f28, d_2002_v1_, 323)
                        d_2108_v82_: _dafny.Map
                        d_2108_v82_ = _dafny.Map({not(self.f24): d_2107_v81_})
                        d_2109_v83_: _dafny.Map
                        d_2109_v83_ = d_2108_v82_
                        d_2110_v84_: _dafny.Set
                        d_2110_v84_ = _dafny.Set({(d_2109_v83_), d_2108_v82_, d_2108_v82_})
                        d_2110_v84_ = d_2110_v84_
                        d_2111_v85_: _dafny.Map
                        d_2111_v85_ = _dafny.Map({self.f27: (d_2093_v70_).f33})
                        d_2112_v86_: D1
                        d_2112_v86_ = D1_DC1((d_2093_v70_).f33)
                        d_2113_v87_: _dafny.Array
                        nw328_ = _dafny.Array(None, 26)
                        nw328_[int(0)] = ((d_2093_v70_).f33) + ((d_2093_v70_).f33)
                        nw328_[int(1)] = (d_2093_v70_).f33
                        nw328_[int(2)] = (d_2093_v70_).f33
                        nw328_[int(3)] = (d_2093_v70_).f33
                        nw328_[int(4)] = (d_2093_v70_).f33
                        nw328_[int(5)] = (d_2093_v70_).f33
                        nw328_[int(6)] = ((d_2093_v70_).f33) + (_dafny.SeqWithoutIsStrInference([self.f27 for d_2114_i7_ in range(default__.abs(767))]))
                        nw328_[int(7)] = ((d_2111_v85_)[_dafny.CodePoint('v')] if (_dafny.CodePoint('v')) in (d_2111_v85_) else (d_2093_v70_).f33)
                        nw328_[int(8)] = (d_2093_v70_).f33
                        nw328_[int(9)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wpcjvwum"))) + ((d_2093_v70_).f33)
                        nw328_[int(10)] = ((d_2112_v86_).cf3) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hjtpq")))
                        nw328_[int(11)] = (d_2093_v70_).f33
                        nw328_[int(12)] = (d_2093_v70_).f33
                        nw328_[int(13)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kttic"))
                        nw328_[int(14)] = ((d_2093_v70_).f33) + ((d_2093_v70_).f33)
                        nw328_[int(15)] = (d_2093_v70_).f33
                        nw328_[int(16)] = (d_2093_v70_).f33
                        nw328_[int(17)] = (d_2093_v70_).f33
                        nw328_[int(18)] = (d_2093_v70_).f33
                        nw328_[int(19)] = ((d_2093_v70_).f33) + ((d_2093_v70_).f33)
                        nw328_[int(20)] = (d_2093_v70_).f33
                        nw328_[int(21)] = ((d_2093_v70_).f33) + ((d_2093_v70_).f33)
                        nw328_[int(22)] = ((d_2093_v70_).f33) + ((d_2093_v70_).f33)
                        nw328_[int(23)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "axgf"))
                        nw328_[int(24)] = ((d_2093_v70_).f33 if (d_2093_v70_).f32 else ((d_2093_v70_).f33).set(default__.safeIndex(d_2002_v1_, len((d_2093_v70_).f33)), self.f27))
                        nw328_[int(25)] = ((d_2093_v70_).f33) + ((d_2093_v70_).f33)
                        d_2113_v87_ = nw328_
                        d_2113_v87_ = d_2113_v87_
                        d_2115_v88_: _dafny.Map
                        d_2115_v88_ = _dafny.Map({d_2002_v1_: (d_2093_v70_).f32})
                        d_2116_v90_: _dafny.Array
                        nw329_ = _dafny.Array(None, 22)
                        nw329_[int(0)] = d_2115_v88_
                        nw329_[int(1)] = d_2115_v88_
                        nw329_[int(2)] = d_2115_v88_
                        nw329_[int(3)] = _dafny.Map({d_2002_v1_: (self).f26})
                        nw329_[int(4)] = _dafny.Map({d_2002_v1_: (self).f26})
                        nw329_[int(5)] = (_dafny.Map({d_2002_v1_: (self).fm7(d_2091_v68_, globalState)})).set(d_2002_v1_, True)
                        nw329_[int(6)] = d_2115_v88_
                        nw329_[int(7)] = d_2115_v88_
                        nw329_[int(8)] = d_2115_v88_
                        nw329_[int(9)] = _dafny.Map({d_2002_v1_: self.f24})
                        nw329_[int(10)] = d_2115_v88_
                        nw329_[int(11)] = d_2115_v88_
                        nw329_[int(12)] = d_2115_v88_
                        nw329_[int(13)] = _dafny.Map({73: (self).f28})
                        nw329_[int(14)] = d_2115_v88_
                        nw329_[int(15)] = d_2115_v88_
                        nw329_[int(16)] = d_2115_v88_
                        nw329_[int(17)] = d_2115_v88_
                        nw329_[int(18)] = d_2115_v88_
                        nw329_[int(19)] = d_2115_v88_
                        def iife149_():
                            coll75_ = _dafny.Map()
                            compr_75_: int
                            for compr_75_ in (d_2098_v74_).Elements:
                                d_2117_v89_: int = compr_75_
                                if (d_2117_v89_) in (d_2098_v74_):
                                    coll75_[(d_2117_v89_) + (d_2002_v1_)] = (d_2093_v70_).f32
                            return _dafny.Map(coll75_)
                        nw329_[int(20)] = iife149_()
                        
                        nw329_[int(21)] = d_2115_v88_
                        d_2116_v90_ = nw329_
                        d_2118_v91_: _dafny.Map
                        d_2118_v91_ = _dafny.Map({(d_2093_v70_).f32: d_2116_v90_})
                        d_2119_v92_: _dafny.Array
                        nw330_ = _dafny.Array(None, 11)
                        nw330_[int(0)] = (d_2116_v90_ if (d_2093_v70_).f32 else d_2116_v90_)
                        nw330_[int(1)] = d_2116_v90_
                        nw330_[int(2)] = d_2116_v90_
                        nw330_[int(3)] = d_2116_v90_
                        nw330_[int(4)] = ((d_2118_v91_)[(d_2093_v70_).f32] if ((d_2093_v70_).f32) in (d_2118_v91_) else d_2116_v90_)
                        nw330_[int(5)] = d_2116_v90_
                        nw330_[int(6)] = d_2116_v90_
                        nw330_[int(7)] = d_2116_v90_
                        nw330_[int(8)] = d_2116_v90_
                        nw330_[int(9)] = d_2116_v90_
                        nw330_[int(10)] = d_2116_v90_
                        d_2119_v92_ = nw330_
                        index311_ = default__.safeIndex(293, (d_2119_v92_).length(0))
                        (d_2119_v92_)[index311_] = d_2116_v90_
                        d_2120_v93_: _dafny.Set
                        d_2120_v93_ = _dafny.Set({(self).f25})
                        d_2121_v94_: _dafny.Map
                        d_2121_v94_ = _dafny.Map({len((_dafny.SeqWithoutIsStrInference([d_2120_v93_, d_2120_v93_, d_2120_v93_, d_2120_v93_, d_2120_v93_])) + (_dafny.SeqWithoutIsStrInference([d_2120_v93_ for d_2122_i8_ in range(default__.abs(279))]))): len(p0)})
                        d_2123_v95_: _dafny.Seq
                        d_2123_v95_ = _dafny.SeqWithoutIsStrInference([D2_DC4((d_2093_v70_).f32)])
                        index312_ = default__.safeIndex(293, (d_2119_v92_).length(0))
                        rhs313_ = d_2116_v90_
                        rhs314_ = d_2121_v94_
                        rhs315_ = (len(d_2123_v95_)) != ((0) - (d_2002_v1_))
                        lhs274_ = d_2119_v92_
                        lhs275_ = default__.safeIndex(293, (d_2119_v92_).length(0))
                        lhs276_ = globalState
                        lhs274_[lhs275_] = rhs313_
                        d_2121_v94_ = rhs314_
                        lhs276_.f10 = rhs315_
                        d_2124_v96_: _dafny.Array
                        def lambda103_(d_2125_v1_):
                            def lambda104_(d_2126_i9_):
                                return (d_2126_i9_) + (d_2125_v1_)

                            return lambda104_

                        init61_ = lambda103_(d_2002_v1_)
                        nw331_ = _dafny.Array(None, 22)
                        for i0_61_ in range(nw331_.length(0)):
                            nw331_[i0_61_] = init61_(i0_61_)
                        d_2124_v96_ = nw331_
                        index313_ = default__.safeIndex(151, (d_2124_v96_).length(0))
                        (d_2124_v96_)[index313_] = d_2002_v1_
                        d_2127_v97_: _dafny.Map
                        d_2127_v97_ = _dafny.Map({(d_2093_v70_).f32: default__.fm21((d_2093_v70_).f33, (d_2093_v70_).f32, d_2002_v1_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dpebkoc"))), globalState)})
                        index314_ = default__.safeIndex(151, (d_2124_v96_).length(0))
                        (d_2124_v96_)[index314_] = (((d_2098_v74_) - (_dafny.MultiSet([761]))).set((d_2002_v1_) + (d_2002_v1_), default__.abs(default__.safeModuloInt(len(d_2127_v97_), d_2002_v1_)))).cardinality
                    elif True:
                        (globalState).f10 = False
                        (globalState).f11 = len(((d_2093_v70_).f33) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uojlmr"))))
                        index315_ = default__.safeIndex(545, (d_2094_v71_).length(0))
                        (d_2094_v71_)[index315_] = (d_2093_v70_).f32
                        index316_ = default__.safeIndex(545, (d_2094_v71_).length(0))
                        (d_2094_v71_)[index316_] = not((d_2095_v72_).cf26)
                        d_2090_v67_ = d_2090_v67_
                        (d_2093_v70_).m9(globalState)
                    d_2128_v98_: C4
                    nw332_ = C4()
                    nw332_.ctor__((d_2093_v70_).f32, (d_2093_v70_).f32)
                    d_2128_v98_ = nw332_
                    d_2129_v99_: _dafny.Array
                    nw333_ = _dafny.Array(int(0), 3)
                    d_2129_v99_ = nw333_
                    index317_ = default__.safeIndex(849, (d_2129_v99_).length(0))
                    (d_2129_v99_)[index317_] = d_2002_v1_
                    index318_ = default__.safeIndex(849, (d_2129_v99_).length(0))
                    (d_2129_v99_)[index318_] = default__.safeModuloInt(d_2002_v1_, (d_2002_v1_) - (d_2002_v1_))
                    d_2130_v100_: _dafny.Set
                    d_2130_v100_ = _dafny.Set({(self).f25})
                    (self).f24 = (_dafny.Set({not((self).f28), (d_2093_v70_).f32})).ispropersubset((_dafny.Set({self.f24, not((self).f28), (self).f26, (self).f26, (self).f25})).intersection(d_2130_v100_))
                    pass
            pass
        r0 = default__.fm47(globalState)
        r1 = d_2090_v67_
        return r0, r1

    @property
    def f28(self):
        return self._f28

class C7(T1, T0):
    def  __init__(self):
        self._f24: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C7"
    @property
    def f24(self):
        return self._f24
    @f24.setter
    def f24(self, value):
        self._f24 = value
    def ctor__(self, f24):
        (self).f24 = f24

    def fm4(self, p0, p1, globalState):
        return D0_DC0((649) != (-51), (0) - ((757) - (372)), len(_dafny.Map({252: self.f24})))

    def fm5(self, p0, p1, p2, globalState):
        def iife150_():
            coll76_ = _dafny.Set()
            compr_76_: _dafny.Map
            for compr_76_ in (_dafny.SeqWithoutIsStrInference([_dafny.Map({-81: 584}), _dafny.Map({(0) - ((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bcufgk"))))): len(_dafny.Map({self.f24: False}))})])).Elements:
                d_2149_v0_: _dafny.Map = compr_76_
                if (d_2149_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.Map({-81: 584}), _dafny.Map({(0) - ((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bcufgk"))))): len(_dafny.Map({self.f24: False}))})])):
                    coll76_ = coll76_.union(_dafny.Set([d_2149_v0_]))
            return _dafny.Set(coll76_)
        return (default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ehhsxxmh")))])), -176)) * ((len(iife150_()
        )) - (721))

    def fm15(self, p0, globalState):
        return ((_dafny.Map({len(_dafny.Map({self.f24: self.f24})): self.f24})) | (_dafny.Map({len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([(0) - (-854), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gdglbboxy")))])): (D2_DC5(_dafny.CodePoint('j'), self.f24, len(_dafny.SeqWithoutIsStrInference([self.f24])), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_2150_i0_ in range(default__.abs(380))])))).cf11})): self.f24}))) | ((_dafny.Map({705: self.f24})) | (_dafny.Map({(D0_DC0(self.f24, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "spnokcg"))), 700)).cf2: self.f24})))

    def m5(self, p0, p1, p2, globalState):
        r0: int = int(0)
        r1: _dafny.MultiSet = _dafny.MultiSet({})
        d_2151_v0_: _dafny.MultiSet
        d_2151_v0_ = _dafny.MultiSet([p0, self.f24])
        d_2152_v1_: int
        d_2152_v1_ = -986
        d_2153_v2_: _dafny.MultiSet
        d_2153_v2_ = _dafny.MultiSet([d_2152_v1_])
        d_2154_v3_: _dafny.Seq
        d_2154_v3_ = _dafny.SeqWithoutIsStrInference([d_2152_v1_, (0) - (d_2152_v1_)])
        d_2155_v4_: _dafny.Map
        d_2155_v4_ = _dafny.Map({((d_2153_v2_)[d_2152_v1_] if (d_2152_v1_) in (d_2153_v2_) else 190): d_2154_v3_})
        d_2156_v5_: _dafny.Seq
        d_2156_v5_ = _dafny.SeqWithoutIsStrInference([p0])
        d_2157_v6_: _dafny.Map
        d_2157_v6_ = _dafny.Map({d_2152_v1_: self.f24})
        d_2158_v7_: _dafny.Map
        d_2158_v7_ = _dafny.Map({self.f24: default__.fm0(d_2152_v1_, self.f24, d_2157_v6_, d_2152_v1_, globalState)})
        d_2159_v8_: _dafny.Array
        nw334_ = _dafny.Array(None, 16)
        nw334_[int(0)] = d_2151_v0_
        nw334_[int(1)] = d_2151_v0_
        nw334_[int(2)] = d_2151_v0_
        nw334_[int(3)] = (d_2151_v0_) | (d_2151_v0_)
        nw334_[int(4)] = default__.fm16(374, d_2155_v4_, self.f24, _dafny.Set({d_2152_v1_, ((d_2151_v0_)[self.f24] if (self.f24) in (d_2151_v0_) else (0) - (d_2152_v1_))}), globalState)
        nw334_[int(5)] = d_2151_v0_
        nw334_[int(6)] = d_2151_v0_
        nw334_[int(7)] = d_2151_v0_
        nw334_[int(8)] = (_dafny.MultiSet(d_2156_v5_)) - (d_2151_v0_)
        nw334_[int(9)] = d_2151_v0_
        nw334_[int(10)] = ((d_2151_v0_).set(self.f24, default__.abs(d_2152_v1_))).set(self.f24, default__.abs(len(d_2158_v7_)))
        nw334_[int(11)] = (_dafny.MultiSet([self.f24])).set(self.f24, default__.abs(d_2152_v1_))
        nw334_[int(12)] = d_2151_v0_
        nw334_[int(13)] = (_dafny.MultiSet([p0])) | (d_2151_v0_)
        nw334_[int(14)] = d_2151_v0_
        nw334_[int(15)] = (d_2151_v0_) - (d_2151_v0_)
        d_2159_v8_ = nw334_
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_2159_v8_).length(0)):
            d_2160_i0_: int = guard_loop_3_
            if (True) and (((0) <= (d_2160_i0_)) and ((d_2160_i0_) < ((d_2159_v8_).length(0)))):
                (d_2159_v8_)[(d_2160_i0_)] = (d_2151_v0_).intersection((d_2151_v0_).intersection(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([self.f24]))))
        (self).f24 = p0
        (globalState).f10 = not(False)
        d_2161_v9_: str
        d_2161_v9_ = _dafny.CodePoint('l')
        d_2161_v9_ = _dafny.CodePoint('t')
        d_2162_v10_: _dafny.Seq
        d_2162_v10_ = _dafny.SeqWithoutIsStrInference([d_2156_v5_, d_2156_v5_])
        d_2163_v11_: _dafny.Map
        d_2163_v11_ = _dafny.Map({self.f24: _dafny.SeqWithoutIsStrInference([p0, False])})
        d_2164_v12_: _dafny.Seq
        d_2164_v12_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([False]), d_2151_v0_, d_2151_v0_, _dafny.MultiSet(((d_2162_v10_)[default__.safeIndex(d_2152_v1_, len(d_2162_v10_))]) + (((d_2163_v11_)[self.f24] if (self.f24) in (d_2163_v11_) else d_2156_v5_))), d_2151_v0_])
        r0 = (0) - (len(d_2164_v12_))
        d_2165_i1_: int
        d_2165_i1_ = 0
        with _dafny.label("19"):
            while self.f24:
                with _dafny.c_label("19"):
                    if (d_2165_i1_) >= (100):
                        raise _dafny.Break("19")
                    d_2165_i1_ = (d_2165_i1_) + (1)
                    r0 = (d_2151_v0_).cardinality
                    d_2166_v13_: _dafny.Seq
                    d_2166_v13_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pwqgxafk"))
                    d_2167_v14_: T0
                    nw335_ = C2()
                    nw335_.ctor__(d_2166_v13_, d_2166_v13_)
                    d_2167_v14_ = nw335_
                    d_2168_v15_: _dafny.Map
                    d_2168_v15_ = _dafny.Map({d_2167_v14_: 192})
                    d_2169_v16_: C5
                    nw336_ = C5()
                    nw336_.ctor__((689) > (d_2152_v1_), _dafny.CodePoint('l'), p0, (-880) < (len(d_2168_v15_)))
                    d_2169_v16_ = nw336_
                    d_2170_v17_: _dafny.Array
                    nw337_ = _dafny.Array(_dafny.Array(None, 0), 25)
                    d_2170_v17_ = nw337_
                    d_2171_v18_: _dafny.Array
                    def lambda106_(d_2172_v13_):
                        def lambda107_(d_2173_i2_):
                            return d_2172_v13_

                        return lambda107_

                    init62_ = lambda106_(d_2166_v13_)
                    nw338_ = _dafny.Array(None, 5)
                    for i0_62_ in range(nw338_.length(0)):
                        nw338_[i0_62_] = init62_(i0_62_)
                    d_2171_v18_ = nw338_
                    index319_ = default__.safeIndex(846, (d_2170_v17_).length(0))
                    (d_2170_v17_)[index319_] = d_2171_v18_
                    d_2174_v19_: _dafny.Seq
                    d_2174_v19_ = _dafny.SeqWithoutIsStrInference([d_2166_v13_])
                    d_2175_v20_: _dafny.Map
                    d_2175_v20_ = _dafny.Map({len(d_2174_v19_): d_2152_v1_})
                    d_2176_v21_: _dafny.MultiSet
                    d_2176_v21_ = _dafny.MultiSet([d_2166_v13_])
                    d_2177_v22_: _dafny.Map
                    d_2177_v22_ = _dafny.Map({self.f24: p2})
                    d_2178_v23_: _dafny.Array
                    nw339_ = _dafny.Array(None, 16)
                    nw339_[int(0)] = ((d_2175_v20_)[d_2152_v1_] if (d_2152_v1_) in (d_2175_v20_) else len(default__.fm41(self.f24, d_2161_v9_, default__.fm1(self.f24, default__.fm1(self.f24, p0, d_2152_v1_, globalState), d_2152_v1_, globalState), globalState)))
                    nw339_[int(1)] = default__.safeModuloInt((d_2151_v0_).cardinality, d_2152_v1_)
                    nw339_[int(2)] = (0) - ((d_2176_v21_).cardinality)
                    nw339_[int(3)] = d_2152_v1_
                    nw339_[int(4)] = default__.safeDivisionInt(d_2152_v1_, d_2152_v1_)
                    nw339_[int(5)] = len(d_2158_v7_)
                    nw339_[int(6)] = d_2152_v1_
                    nw339_[int(7)] = d_2152_v1_
                    nw339_[int(8)] = d_2152_v1_
                    nw339_[int(9)] = len(d_2158_v7_)
                    nw339_[int(10)] = len(((d_2177_v22_)[p0] if (p0) in (d_2177_v22_) else p2))
                    nw339_[int(11)] = (d_2152_v1_) * (d_2152_v1_)
                    nw339_[int(12)] = 60
                    nw339_[int(13)] = default__.safeDivisionInt(993, d_2152_v1_)
                    nw339_[int(14)] = d_2152_v1_
                    nw339_[int(15)] = (0) - (-380)
                    d_2178_v23_ = nw339_
                    index320_ = default__.safeIndex(880, (d_2178_v23_).length(0))
                    (d_2178_v23_)[index320_] = d_2152_v1_
                    index321_ = default__.safeIndex(846, (d_2170_v17_).length(0))
                    index322_ = default__.safeIndex(880, (d_2178_v23_).length(0))
                    rhs316_ = d_2152_v1_
                    rhs317_ = d_2152_v1_
                    rhs318_ = d_2171_v18_
                    rhs319_ = d_2152_v1_
                    rhs320_ = ((len(_dafny.Map({self.f24: len(d_2156_v5_)}))) + (len(d_2166_v13_))) != (d_2152_v1_)
                    lhs277_ = globalState
                    lhs278_ = d_2170_v17_
                    lhs279_ = default__.safeIndex(846, (d_2170_v17_).length(0))
                    lhs280_ = d_2178_v23_
                    lhs281_ = default__.safeIndex(880, (d_2178_v23_).length(0))
                    lhs282_ = globalState
                    d_2152_v1_ = rhs316_
                    lhs277_.f11 = rhs317_
                    lhs278_[lhs279_] = rhs318_
                    lhs280_[lhs281_] = rhs319_
                    lhs282_.f10 = rhs320_
                    d_2179_v25_: _dafny.Set
                    d_2179_v25_ = _dafny.Set({(0) - (-218), d_2152_v1_, (d_2178_v23_)[default__.safeIndex(880, (d_2178_v23_).length(0))], -705, (d_2178_v23_)[default__.safeIndex(880, (d_2178_v23_).length(0))]})
                    d_2180_v26_: _dafny.Array
                    nw340_ = _dafny.Array(_dafny.CodePoint('D'), 17)
                    d_2180_v26_ = nw340_
                    d_2181_v27_: _dafny.MultiSet
                    d_2181_v27_ = _dafny.MultiSet([d_2180_v26_])
                    d_2182_v28_: C3
                    nw341_ = C3()
                    def iife151_():
                        coll77_ = _dafny.Set()
                        compr_77_: int
                        for compr_77_ in (d_2157_v6_).keys.Elements:
                            d_2183_v24_: int = compr_77_
                            if (d_2183_v24_) in (d_2157_v6_):
                                coll77_ = coll77_.union(_dafny.Set([default__.safeDivisionInt(d_2183_v24_, 410)]))
                        return _dafny.Set(coll77_)
                    nw341_.ctor__((iife151_()
                    ).ispropersubset(d_2179_v25_), _dafny.SeqWithoutIsStrInference([d_2161_v9_ for d_2184_i3_ in range(default__.abs(447))]), ((d_2178_v23_)[default__.safeIndex(880, (d_2178_v23_).length(0))]) < (d_2152_v1_), d_2161_v9_, self.f24, (d_2181_v27_).ispropersubset(d_2181_v27_))
                    d_2182_v28_ = nw341_
                    pass
            pass
        r0 = ((d_2158_v7_)[p0] if (p0) in (d_2158_v7_) else (len(d_2156_v5_)) + (len(d_2156_v5_)))
        r1 = (d_2153_v2_).set(164, default__.abs(d_2152_v1_))
        return r0, r1

    def m6(self, p0, p1, globalState):
        r0: _dafny.Set = _dafny.Set({})
        d_2185_v0_: _dafny.Seq
        d_2185_v0_ = _dafny.SeqWithoutIsStrInference([self.f24])
        hi14_ = (len(d_2185_v0_)) - (p0)
        for d_2186_i0_ in range(p0, hi14_):
            d_2187_v1_: str
            d_2187_v1_ = _dafny.CodePoint('l')
            d_2188_v2_: _dafny.Seq
            d_2188_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u"))
            d_2189_v3_: _dafny.Set
            d_2189_v3_ = _dafny.Set({p0, p0})
            d_2190_v4_: _dafny.Map
            d_2190_v4_ = _dafny.Map({len(d_2189_v3_): p0})
            d_2191_v5_: D3
            d_2191_v5_ = D3_DC8((d_2187_v1_) not in (d_2188_v2_), p0, ((0) - (p0) if True else ((d_2190_v4_)[d_2186_i0_] if (d_2186_i0_) in (d_2190_v4_) else p0)))
            d_2191_v5_ = d_2191_v5_
            d_2192_v6_: _dafny.MultiSet
            d_2192_v6_ = _dafny.MultiSet([d_2186_i0_])
            d_2193_v7_: _dafny.Array
            nw342_ = _dafny.Array(None, 8)
            nw342_[int(0)] = d_2192_v6_
            nw342_[int(1)] = d_2192_v6_
            nw342_[int(2)] = d_2192_v6_
            nw342_[int(3)] = d_2192_v6_
            nw342_[int(4)] = d_2192_v6_
            nw342_[int(5)] = d_2192_v6_
            nw342_[int(6)] = d_2192_v6_
            nw342_[int(7)] = d_2192_v6_
            d_2193_v7_ = nw342_
            d_2194_v8_: _dafny.Map
            d_2194_v8_ = _dafny.Map({d_2193_v7_: self.f24})
            d_2195_v9_: C6
            nw343_ = C6()
            nw343_.ctor__((d_2186_i0_) < (p0), (self.f24) and (self.f24), ((d_2194_v8_)[d_2193_v7_] if (d_2193_v7_) in (d_2194_v8_) else self.f24), d_2187_v1_, default__.fm1(self.f24, True, p0, globalState))
            d_2195_v9_ = nw343_
            d_2196_v10_: _dafny.Array
            nw344_ = _dafny.Array(_dafny.Array(None, 0), 9)
            d_2196_v10_ = nw344_
            d_2197_v11_: _dafny.Array
            nw345_ = _dafny.Array(int(0), 1)
            d_2197_v11_ = nw345_
            index323_ = default__.safeIndex(868, (d_2196_v10_).length(0))
            (d_2196_v10_)[index323_] = d_2197_v11_
            index324_ = default__.safeIndex(868, (d_2196_v10_).length(0))
            (d_2196_v10_)[index324_] = d_2197_v11_
            (globalState).f10 = (d_2195_v9_).f28
        d_2198_v12_: _dafny.MultiSet
        d_2198_v12_ = _dafny.MultiSet([self.f24, self.f24, default__.fm1(self.f24, self.f24, p0, globalState)])
        if ((d_2198_v12_).cardinality) == ((0) - (p0)):
            (self).f24 = not(self.f24)
            d_2199_v13_: _dafny.Array
            nw346_ = _dafny.Array(_dafny.Map({}), 18)
            d_2199_v13_ = nw346_
            d_2200_v14_: D10
            d_2200_v14_ = D10_DC24(D10_DC21(d_2199_v13_))
            d_2201_v15_: D10
            d_2201_v15_ = D10_DC24(d_2200_v14_)
            source37_ = d_2201_v15_
            if source37_.is_DC22:
                d_2202___mcc_h0_ = source37_.cf41
                d_2203___mcc_h1_ = source37_.cf42
                d_2204_cf42_ = d_2203___mcc_h1_
                d_2205_cf41_ = d_2202___mcc_h0_
                d_2206_v16_: _dafny.Map
                d_2206_v16_ = _dafny.Map({False: True})
                d_2207_v17_: _dafny.Seq
                d_2207_v17_ = _dafny.SeqWithoutIsStrInference([d_2206_v16_, _dafny.Map({d_2204_cf42_: d_2204_cf42_}), (d_2206_v16_).set(default__.fm1(d_2204_cf42_, self.f24, d_2205_cf41_, globalState), self.f24), _dafny.Map({self.f24: d_2204_cf42_})])
                (self).f24 = (_dafny.Set({(d_2206_v16_).set(d_2204_cf42_, d_2204_cf42_)})) != (_dafny.Set({(d_2207_v17_)[default__.safeIndex(p0, len(d_2207_v17_))], d_2206_v16_, d_2206_v16_, d_2206_v16_, d_2206_v16_}))
                d_2208_v18_: T1
                nw347_ = C1()
                nw347_.ctor__(d_2204_cf42_)
                d_2208_v18_ = nw347_
                d_2209_v19_: T1
                d_2209_v19_ = d_2208_v18_
                d_2210_v20_: _dafny.Map
                d_2210_v20_ = _dafny.Map({d_2209_v19_: self.f24})
                d_2210_v20_ = (d_2210_v20_).set(d_2209_v19_, self.f24)
                d_2211_v21_: D15
                d_2211_v21_ = D15_DC32(d_2206_v16_)
                d_2212_v22_: _dafny.Array
                nw348_ = _dafny.Array(None, 17)
                nw348_[int(0)] = d_2206_v16_
                nw348_[int(1)] = _dafny.Map({self.f24: default__.fm1(d_2208_v18_.f24, d_2208_v18_.f24, d_2205_cf41_, globalState)})
                nw348_[int(2)] = (d_2206_v16_) | (_dafny.Map({self.f24: d_2204_cf42_}))
                nw348_[int(3)] = d_2206_v16_
                nw348_[int(4)] = (default__.fm2(d_2185_v0_, False, (d_2198_v12_).cardinality, d_2205_cf41_, globalState) if self.f24 else _dafny.Map({self.f24: True}))
                nw348_[int(5)] = (d_2211_v21_).cf54
                nw348_[int(6)] = d_2206_v16_
                nw348_[int(7)] = d_2206_v16_
                nw348_[int(8)] = d_2206_v16_
                nw348_[int(9)] = d_2206_v16_
                nw348_[int(10)] = d_2206_v16_
                nw348_[int(11)] = d_2206_v16_
                nw348_[int(12)] = d_2206_v16_
                nw348_[int(13)] = d_2206_v16_
                nw348_[int(14)] = _dafny.Map({d_2204_cf42_: not(False)})
                nw348_[int(15)] = d_2206_v16_
                nw348_[int(16)] = _dafny.Map({False: d_2204_cf42_})
                d_2212_v22_ = nw348_
                index325_ = default__.safeIndex(523, (d_2212_v22_).length(0))
                (d_2212_v22_)[index325_] = (d_2206_v16_) | (d_2206_v16_)
                d_2213_v23_: str
                d_2213_v23_ = _dafny.CodePoint('f')
                d_2214_v24_: _dafny.Seq
                d_2214_v24_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cwwsaoci"))
                d_2215_v25_: _dafny.Map
                d_2215_v25_ = _dafny.Map({p0: self.f24})
                index326_ = default__.safeIndex(523, (d_2212_v22_).length(0))
                (d_2212_v22_)[index326_] = (_dafny.Map({True: ((d_2215_v25_)[d_2205_cf41_] if (d_2205_cf41_) in (d_2215_v25_) else d_2204_cf42_)}) if (d_2213_v23_) not in (d_2214_v24_) else _dafny.Map({not(not(self.f24)): self.f24}))
                d_2216_v26_: C2
                nw349_ = C2()
                nw349_.ctor__(d_2214_v24_, d_2214_v24_)
                d_2216_v26_ = nw349_
            elif source37_.is_DC23:
                d_2217___mcc_h2_ = source37_.cf43
                d_2218___mcc_h3_ = source37_.cf44
                d_2219_cf44_ = d_2218___mcc_h3_
                d_2220_cf43_ = d_2217___mcc_h2_
                (self).f24 = self.f24
                d_2221_v27_: _dafny.Seq
                d_2221_v27_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fnkwaoaon"))
                d_2222_v28_: D0
                d_2222_v28_ = D0_DC0(d_2219_cf44_, p0, p0)
                d_2223_v29_: _dafny.Seq
                d_2223_v29_ = _dafny.SeqWithoutIsStrInference([D0_DC0(self.f24, p0, p0), d_2222_v28_, d_2222_v28_])
                d_2224_v30_: _dafny.Map
                d_2224_v30_ = _dafny.Map({d_2219_cf44_: d_2223_v29_})
                d_2225_v31_: _dafny.MultiSet
                d_2225_v31_ = _dafny.MultiSet([p0])
                d_2226_v32_: _dafny.Seq
                d_2226_v32_ = _dafny.SeqWithoutIsStrInference([p0])
                d_2227_v33_: _dafny.Seq
                d_2227_v33_ = _dafny.SeqWithoutIsStrInference([d_2226_v32_, _dafny.SeqWithoutIsStrInference([p0])])
                nw350_ = _dafny.Array(None, 28)
                nw350_[int(0)] = p0
                nw350_[int(1)] = (p0) + (p0)
                nw350_[int(2)] = len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ifvsn"))) + (d_2221_v27_))
                nw350_[int(3)] = p0
                nw350_[int(4)] = p0
                nw350_[int(5)] = len((d_2221_v27_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ppwatrie"))))
                nw350_[int(6)] = (self).fm5(d_2224_v30_, self.f24, self.f24, globalState)
                nw350_[int(7)] = (373) + (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_2228_i1_ in range(default__.abs(-95))])))
                nw350_[int(8)] = p0
                nw350_[int(9)] = p0
                nw350_[int(10)] = (((d_2225_v31_).set(p0, default__.abs(p0))) - (_dafny.MultiSet((d_2227_v33_)[default__.safeIndex(p0, len(d_2227_v33_))]))).cardinality
                nw350_[int(11)] = p0
                nw350_[int(12)] = p0
                nw350_[int(13)] = len(d_2220_cf43_)
                nw350_[int(14)] = p0
                nw350_[int(15)] = p0
                nw350_[int(16)] = p0
                nw350_[int(17)] = p0
                nw350_[int(18)] = p0
                nw350_[int(19)] = p0
                nw350_[int(20)] = (0) - (p0)
                nw350_[int(21)] = p0
                nw350_[int(22)] = (d_2225_v31_).cardinality
                nw350_[int(23)] = len(_dafny.SeqWithoutIsStrInference([p0, p0]))
                nw350_[int(24)] = (0) - (p0)
                nw350_[int(25)] = p0
                nw350_[int(26)] = p0
                nw350_[int(27)] = p0
                (globalState).f7 = nw350_
                d_2229_v34_: str
                d_2229_v34_ = _dafny.CodePoint('a')
                d_2229_v34_ = _dafny.CodePoint('h')
                d_2230_v35_: C6
                nw351_ = C6()
                nw351_.ctor__(True, d_2219_cf44_, d_2219_cf44_, d_2229_v34_, (p0) < ((0) - (p0)))
                d_2230_v35_ = nw351_
            elif source37_.is_DC21:
                d_2231___mcc_h4_ = source37_.cf40
                d_2232_cf40_ = d_2231___mcc_h4_
                d_2233_v36_: _dafny.Seq
                d_2233_v36_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xnwor"))
                d_2233_v36_ = ((d_2233_v36_) + (d_2233_v36_)) + (d_2233_v36_)
                (globalState).f11 = default__.safeDivisionInt(p0, (p0) * (p0))
                d_2234_v37_: _dafny.MultiSet
                d_2234_v37_ = _dafny.MultiSet([p0])
                d_2235_v38_: str
                d_2235_v38_ = _dafny.CodePoint('e')
                d_2236_v39_: _dafny.Set
                d_2236_v39_ = _dafny.Set({(_dafny.SeqWithoutIsStrInference([d_2235_v38_ for d_2237_i2_ in range(default__.abs(28))])) < (d_2233_v36_)})
                d_2238_v40_: _dafny.Map
                d_2238_v40_ = _dafny.Map({p0: self.f24})
                d_2239_v42_: _dafny.Set
                d_2239_v42_ = _dafny.Set({p0, p0})
                d_2240_v44_: _dafny.Map
                d_2240_v44_ = _dafny.Map({default__.fm0((d_2198_v12_).cardinality, self.f24, d_2238_v40_, p0, globalState): (0) - (p0)})
                d_2241_v45_: _dafny.Array
                nw352_ = _dafny.Array(None, 25)
                nw352_[int(0)] = p0
                nw352_[int(1)] = (len(d_2233_v36_) if self.f24 else len(d_2185_v0_))
                nw352_[int(2)] = -183
                nw352_[int(3)] = 581
                nw352_[int(4)] = (len(d_2236_v39_)) * (p0)
                nw352_[int(5)] = len(d_2238_v40_)
                nw352_[int(6)] = (p0) * (p0)
                nw352_[int(7)] = (d_2234_v37_).cardinality
                nw352_[int(8)] = default__.safeModuloInt((0) - ((0) - (len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qcahaefbq"))).set(default__.safeIndex(124, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qcahaefbq")))), d_2235_v38_)))), p0)
                nw352_[int(9)] = p0
                def iife152_():
                    coll78_ = _dafny.Map()
                    compr_78_: int
                    for compr_78_ in (default__.fm63(globalState)).keys.Elements:
                        d_2242_v41_: int = compr_78_
                        if (d_2242_v41_) in (default__.fm63(globalState)):
                            coll78_[(d_2242_v41_) + ((D2_DC5(d_2235_v38_, self.f24, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))), len(_dafny.Set({p0})))).cf13)] = p0
                    return _dafny.Map(coll78_)
                nw352_[int(10)] = len((D2_DC3(iife152_()
)).cf8)
                nw352_[int(11)] = default__.safeDivisionInt(p0, p0)
                nw352_[int(12)] = (0) - (p0)
                nw352_[int(13)] = (p0) - (p0)
                nw352_[int(14)] = p0
                nw352_[int(15)] = p0
                nw352_[int(16)] = p0
                nw352_[int(17)] = (0) - (p0)
                nw352_[int(18)] = p0
                nw352_[int(19)] = p0
                nw352_[int(20)] = len((d_2233_v36_ if self.f24 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bx"))))
                nw352_[int(21)] = 404
                nw352_[int(22)] = len(d_2239_v42_)
                def iife153_():
                    coll79_ = _dafny.Map()
                    compr_79_: int
                    for compr_79_ in _dafny.IntegerRange(739, 911):
                        d_2243_v43_: int = compr_79_
                        if ((739) <= (d_2243_v43_)) and ((d_2243_v43_) < (911)):
                            coll79_[(d_2243_v43_) + (p0)] = p0
                    return _dafny.Map(coll79_)
                nw352_[int(23)] = len((iife153_()
                ) | (d_2240_v44_))
                nw352_[int(24)] = p0
                d_2241_v45_ = nw352_
                index327_ = default__.safeIndex(271, (d_2241_v45_).length(0))
                (d_2241_v45_)[index327_] = p0
                index328_ = default__.safeIndex(271, (d_2241_v45_).length(0))
                rhs321_ = (d_2234_v37_) - (d_2234_v37_)
                rhs322_ = d_2235_v38_
                rhs323_ = d_2236_v39_
                rhs324_ = p0
                rhs325_ = (self.f24 if self.f24 else self.f24)
                lhs283_ = d_2241_v45_
                lhs284_ = default__.safeIndex(271, (d_2241_v45_).length(0))
                lhs285_ = self
                d_2234_v37_ = rhs321_
                d_2235_v38_ = rhs322_
                d_2236_v39_ = rhs323_
                lhs283_[lhs284_] = rhs324_
                lhs285_.f24 = rhs325_
                rhs326_ = ((d_2241_v45_)[default__.safeIndex(271, (d_2241_v45_).length(0))]) == (((0) - ((d_2241_v45_)[default__.safeIndex(271, (d_2241_v45_).length(0))])) + (802))
                rhs327_ = self.f24
                lhs286_ = self
                lhs287_ = self
                lhs286_.f24 = rhs326_
                lhs287_.f24 = rhs327_
            elif True:
                d_2244___mcc_h5_ = source37_.cf45
                d_2245_cf45_ = d_2244___mcc_h5_
                d_2246_v46_: _dafny.Array
                nw353_ = _dafny.Array(_dafny.Array(None, 0), 23)
                d_2246_v46_ = nw353_
                d_2246_v46_ = d_2246_v46_
                d_2247_v47_: _dafny.Array
                nw354_ = _dafny.Array(int(0), 12)
                d_2247_v47_ = nw354_
                d_2248_v48_: str
                d_2248_v48_ = _dafny.CodePoint('m')
                d_2249_v49_: D1
                d_2249_v49_ = D1_DC2(d_2247_v47_, d_2248_v48_, p0, self.f24)
                d_2250_v50_: _dafny.Seq
                d_2250_v50_ = _dafny.SeqWithoutIsStrInference([d_2248_v48_])
                d_2251_v51_: _dafny.Map
                d_2251_v51_ = _dafny.Map({p0: self.f24})
                d_2252_v52_: _dafny.Set
                d_2252_v52_ = _dafny.Set({p0, 501})
                d_2253_v53_: _dafny.Array
                def lambda108_(d_2254_i3_):
                    return self.f24

                init63_ = lambda108_
                nw355_ = _dafny.Array(None, 29)
                for i0_63_ in range(nw355_.length(0)):
                    nw355_[i0_63_] = init63_(i0_63_)
                d_2253_v53_ = nw355_
                d_2255_v54_: D6
                d_2255_v54_ = D6_DC13(d_2253_v53_, -490, True)
                d_2256_v55_: _dafny.MultiSet
                d_2256_v55_ = _dafny.MultiSet([default__.fm0(len(d_2252_v52_), (d_2255_v54_).cf26, d_2251_v51_, p0, globalState), -552, 52])
                pat_let_tv33_ = d_2248_v48_
                pat_let_tv34_ = p0
                d_2257_v56_: _dafny.Array
                nw356_ = _dafny.Array(None, 20)
                nw356_[int(0)] = p0
                nw356_[int(1)] = p0
                nw356_[int(2)] = p0
                nw356_[int(3)] = p0
                def iife154_(_pat_let37_0):
                    def iife155_(d_2258_dt__update__tmp_h0_):
                        def iife156_(_pat_let38_0):
                            def iife157_(d_2259_dt__update_hcf7_h0_):
                                def iife158_(_pat_let39_0):
                                    def iife159_(d_2260_dt__update_hcf5_h0_):
                                        def iife160_(_pat_let40_0):
                                            def iife161_(d_2261_dt__update_hcf6_h0_):
                                                return D1_DC2((d_2258_dt__update__tmp_h0_).cf4, d_2260_dt__update_hcf5_h0_, d_2261_dt__update_hcf6_h0_, d_2259_dt__update_hcf7_h0_)
                                            return iife161_(_pat_let40_0)
                                        return iife160_(pat_let_tv34_)
                                    return iife159_(_pat_let39_0)
                                return iife158_(pat_let_tv33_)
                            return iife157_(_pat_let38_0)
                        return iife156_(self.f24)
                    return iife155_(_pat_let37_0)
                nw356_[int(4)] = default__.safeDivisionInt(p0, (iife154_(d_2249_v49_)).cf6)
                nw356_[int(5)] = p0
                nw356_[int(6)] = p0
                nw356_[int(7)] = (745 if self.f24 else p0)
                nw356_[int(8)] = len(default__.fm26(self.f24, len(d_2185_v0_), p0, globalState))
                nw356_[int(9)] = len((d_2250_v50_) + (d_2250_v50_))
                nw356_[int(10)] = p0
                nw356_[int(11)] = p0
                nw356_[int(12)] = p0
                nw356_[int(13)] = (p0 if self.f24 else default__.fm0(p0, not(True), d_2251_v51_, p0, globalState))
                nw356_[int(14)] = p0
                nw356_[int(15)] = (p0) - (len(d_2252_v52_))
                nw356_[int(16)] = p0
                nw356_[int(17)] = len(d_2251_v51_)
                nw356_[int(18)] = ((d_2256_v55_)[p0] if (p0) in (d_2256_v55_) else p0)
                nw356_[int(19)] = p0
                d_2257_v56_ = nw356_
                index329_ = default__.safeIndex(10, (d_2247_v47_).length(0))
                (d_2247_v47_)[index329_] = p0
                d_2262_v57_: _dafny.Map
                d_2262_v57_ = _dafny.Map({(p0) - (p0): len(d_2250_v50_)})
                d_2263_v58_: _dafny.Seq
                d_2263_v58_ = _dafny.SeqWithoutIsStrInference([p0, 713])
                d_2264_v59_: _dafny.Map
                d_2264_v59_ = _dafny.Map({p0: d_2253_v53_})
                index330_ = default__.safeIndex(10, (d_2247_v47_).length(0))
                rhs328_ = (0) - (len(d_2262_v57_))
                rhs329_ = default__.safeDivisionInt(len((d_2263_v58_ if self.f24 else _dafny.SeqWithoutIsStrInference([-412 for d_2265_i4_ in range(default__.abs(838))]))), (p0) * (p0))
                rhs330_ = ((d_2264_v59_)[p0] if (p0) in (d_2264_v59_) else (d_2253_v53_ if self.f24 else d_2253_v53_))
                rhs331_ = 820
                lhs288_ = d_2247_v47_
                lhs289_ = default__.safeIndex(10, (d_2247_v47_).length(0))
                lhs290_ = globalState
                lhs291_ = globalState
                lhs288_[lhs289_] = rhs328_
                lhs290_.f11 = rhs329_
                d_2253_v53_ = rhs330_
                lhs291_.f11 = rhs331_
                d_2266_v60_: _dafny.Array
                nw357_ = _dafny.Array(_dafny.MultiSet({}), 5)
                d_2266_v60_ = nw357_
                rhs332_ = default__.fm0(len((d_2250_v50_) + (default__.fm41(self.f24, d_2248_v48_, self.f24, globalState))), self.f24, (d_2251_v51_).set(p0, self.f24), (0) - (p0), globalState)
                rhs333_ = d_2266_v60_
                lhs292_ = globalState
                lhs292_.f16 = rhs332_
                d_2266_v60_ = rhs333_
                index331_ = default__.safeIndex(10, (d_2247_v47_).length(0))
                (d_2247_v47_)[index331_] = p0
            if self.f24:
                d_2267_v61_: _dafny.Map
                d_2267_v61_ = _dafny.Map({p0: self.f24})
                d_2268_v62_: _dafny.MultiSet
                d_2268_v62_ = _dafny.MultiSet([p0])
                (globalState).f11 = ((0) - (default__.safeDivisionInt(p0, p0))) * (default__.fm0(p0, False, d_2267_v61_, (d_2268_v62_).cardinality, globalState))
                d_2269_v63_: D2
                d_2269_v63_ = D2_DC4(((d_2267_v61_)[(_dafny.MultiSet([self.f24])).cardinality] if ((_dafny.MultiSet([self.f24])).cardinality) in (d_2267_v61_) else self.f24))
                d_2270_v64_: D2
                d_2270_v64_ = D2_DC6(d_2269_v63_)
                d_2271_v65_: D2
                d_2271_v65_ = D2_DC6(d_2270_v64_)
                d_2272_v66_: _dafny.Set
                d_2272_v66_ = _dafny.Set({self.f24, self.f24, not(self.f24)})
                d_2273_v67_: int
                d_2274_v68_: _dafny.MultiSet
                out39_: int
                out40_: _dafny.MultiSet
                out39_, out40_ = (self).m5(self.f24, D2_DC6(d_2270_v64_), d_2272_v66_, globalState)
                d_2273_v67_ = out39_
                d_2274_v68_ = out40_
                d_2275_v69_: _dafny.Array
                def lambda109_(d_2276_i5_):
                    return self.f24

                init64_ = lambda109_
                nw358_ = _dafny.Array(None, 27)
                for i0_64_ in range(nw358_.length(0)):
                    nw358_[i0_64_] = init64_(i0_64_)
                d_2275_v69_ = nw358_
                d_2277_v70_: D6
                d_2277_v70_ = D6_DC13(d_2275_v69_, d_2273_v67_, self.f24)
                (globalState).f16 = ((p0) * ((d_2277_v70_).cf25)) - (d_2273_v67_)
                d_2278_v71_: _dafny.Array
                nw359_ = _dafny.Array(int(0), 16)
                d_2278_v71_ = nw359_
                d_2279_v72_: _dafny.MultiSet
                d_2279_v72_ = _dafny.MultiSet([d_2278_v71_])
                d_2280_v73_: str
                d_2280_v73_ = _dafny.CodePoint('k')
                d_2281_v74_: _dafny.Map
                d_2281_v74_ = _dafny.Map({d_2280_v73_: d_2279_v72_})
                d_2282_v75_: _dafny.Array
                nw360_ = _dafny.Array(None, 29)
                nw360_[int(0)] = _dafny.MultiSet([d_2278_v71_, d_2278_v71_])
                nw360_[int(1)] = d_2279_v72_
                nw360_[int(2)] = _dafny.MultiSet([d_2278_v71_, d_2278_v71_])
                nw360_[int(3)] = d_2279_v72_
                nw360_[int(4)] = ((d_2279_v72_).set(d_2278_v71_, default__.abs(p0))).intersection(d_2279_v72_)
                nw360_[int(5)] = d_2279_v72_
                nw360_[int(6)] = (d_2279_v72_) - (d_2279_v72_)
                nw360_[int(7)] = d_2279_v72_
                nw360_[int(8)] = d_2279_v72_
                nw360_[int(9)] = d_2279_v72_
                nw360_[int(10)] = (d_2279_v72_) - (d_2279_v72_)
                nw360_[int(11)] = d_2279_v72_
                nw360_[int(12)] = (d_2279_v72_).set(d_2278_v71_, default__.abs(p0))
                nw360_[int(13)] = ((d_2281_v74_)[d_2280_v73_] if (d_2280_v73_) in (d_2281_v74_) else d_2279_v72_)
                nw360_[int(14)] = d_2279_v72_
                nw360_[int(15)] = (d_2279_v72_).intersection(d_2279_v72_)
                nw360_[int(16)] = _dafny.MultiSet([d_2278_v71_])
                nw360_[int(17)] = (d_2279_v72_ if self.f24 else d_2279_v72_)
                nw360_[int(18)] = d_2279_v72_
                nw360_[int(19)] = (_dafny.MultiSet([d_2278_v71_])).intersection(d_2279_v72_)
                nw360_[int(20)] = d_2279_v72_
                nw360_[int(21)] = d_2279_v72_
                nw360_[int(22)] = d_2279_v72_
                nw360_[int(23)] = d_2279_v72_
                nw360_[int(24)] = (d_2279_v72_) | (d_2279_v72_)
                nw360_[int(25)] = d_2279_v72_
                nw360_[int(26)] = d_2279_v72_
                nw360_[int(27)] = d_2279_v72_
                nw360_[int(28)] = d_2279_v72_
                d_2282_v75_ = nw360_
                index332_ = default__.safeIndex(584, (d_2282_v75_).length(0))
                (d_2282_v75_)[index332_] = (d_2279_v72_).intersection(_dafny.MultiSet([d_2278_v71_]))
                index333_ = default__.safeIndex(584, (d_2282_v75_).length(0))
                (d_2282_v75_)[index333_] = (d_2279_v72_).set(d_2278_v71_, default__.abs((p0) * ((0) - (d_2273_v67_))))
                rhs334_ = p0
                rhs335_ = 765
                lhs293_ = globalState
                lhs294_ = globalState
                lhs293_.f16 = rhs334_
                lhs294_.f22 = rhs335_
            elif True:
                (self).f24 = ((p0 if False else p0)) <= (p0)
                (globalState).f10 = (self.f24 if self.f24 else self.f24)
                d_2283_v76_: _dafny.Map
                d_2283_v76_ = _dafny.Map({self.f24: p0})
                d_2284_v77_: _dafny.Seq
                d_2284_v77_ = _dafny.SeqWithoutIsStrInference([p0])
                d_2283_v76_ = (d_2283_v76_).set(self.f24, len(d_2284_v77_))
                d_2285_v78_: _dafny.Seq
                d_2285_v78_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mtplgs"))
                d_2286_v79_: D1
                d_2286_v79_ = D1_DC1(d_2285_v78_)
                d_2287_v80_: str
                d_2287_v80_ = _dafny.CodePoint('e')
                d_2288_v81_: C3
                nw361_ = C3()
                nw361_.ctor__(self.f24, (D1_DC1((d_2286_v79_).cf3)).cf3, self.f24, d_2287_v80_, (self.f24) == (self.f24), self.f24)
                d_2288_v81_ = nw361_
                (globalState).f10 = not (False) or ((d_2288_v81_).f32)
            d_2289_v82_: _dafny.Seq
            d_2289_v82_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v"))
            d_2290_v83_: _dafny.Array
            nw362_ = _dafny.Array(False, 27)
            d_2290_v83_ = nw362_
            d_2291_v84_: D6
            d_2291_v84_ = D6_DC13(d_2290_v83_, p0, self.f24)
            d_2292_v85_: C2
            nw363_ = C2()
            nw363_.ctor__(d_2289_v82_, default__.fm38(self.f24, p0, _dafny.MultiSet([(d_2291_v84_).cf25, 799]), _dafny.Map({not(not(not(self.f24))): p0}), globalState))
            d_2292_v85_ = nw363_
            (globalState).f10 = not (not((712) > (p0))) or ((not(self.f24) if self.f24 else self.f24))
        elif True:
            (globalState).f16 = p0
            d_2293_v86_: _dafny.Map
            d_2293_v86_ = _dafny.Map({p0: True})
            d_2294_v87_: _dafny.Array
            nw364_ = _dafny.Array(None, 4)
            nw364_[int(0)] = (d_2185_v0_)[default__.safeIndex(p0, len(d_2185_v0_))]
            nw364_[int(1)] = (p0) <= (len(d_2293_v86_))
            nw364_[int(2)] = self.f24
            nw364_[int(3)] = self.f24
            d_2294_v87_ = nw364_
            index334_ = default__.safeIndex(396, (d_2294_v87_).length(0))
            (d_2294_v87_)[index334_] = default__.fm1(not(True), default__.fm1(self.f24, self.f24, p0, globalState), p0, globalState)
            index335_ = default__.safeIndex(396, (d_2294_v87_).length(0))
            (d_2294_v87_)[index335_] = (self.f24) == ((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "umpkhd")))) >= (p0))
            d_2295_v88_: _dafny.Seq
            d_2295_v88_ = _dafny.SeqWithoutIsStrInference([p0])
            rhs336_ = (0) - (len(((d_2295_v88_) + (d_2295_v88_)) + (d_2295_v88_)))
            rhs337_ = d_2294_v87_
            lhs295_ = globalState
            lhs296_ = globalState
            lhs295_.f22 = rhs336_
            lhs296_.f2 = rhs337_
            (globalState).f16 = (p0) * ((p0 if self.f24 else (0) - (p0)))
            (globalState).f16 = p0
        d_2296_v89_: _dafny.Seq
        d_2296_v89_ = _dafny.SeqWithoutIsStrInference([p0, p0, default__.safeModuloInt(p0, p0)])
        d_2296_v89_ = ((d_2296_v89_).set(default__.safeIndex(p0, len(d_2296_v89_)), p0) if (self.f24) == (not(self.f24)) else _dafny.SeqWithoutIsStrInference([p0, 341]))
        (self).f24 = self.f24
        d_2296_v89_ = d_2296_v89_
        d_2297_v90_: _dafny.Array
        nw365_ = _dafny.Array(_dafny.Array(None, 0), 4)
        d_2297_v90_ = nw365_
        d_2298_v91_: _dafny.Array
        nw366_ = _dafny.Array(_dafny.Array(None, 0), 5)
        d_2298_v91_ = nw366_
        d_2299_v92_: D10
        d_2299_v92_ = D10_DC22(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_2300_i6_ in range(default__.abs(-533))])), False)
        pat_let_tv35_ = p0
        d_2301_v93_: _dafny.Array
        nw367_ = _dafny.Array(None, 19)
        nw367_[int(0)] = d_2299_v92_
        nw367_[int(1)] = d_2299_v92_
        nw367_[int(2)] = D10_DC22(p0, self.f24)
        nw367_[int(3)] = d_2299_v92_
        nw367_[int(4)] = d_2299_v92_
        nw367_[int(5)] = default__.fm64(globalState)
        nw367_[int(6)] = D10_DC22(p0, self.f24)
        nw367_[int(7)] = d_2299_v92_
        nw367_[int(8)] = d_2299_v92_
        nw367_[int(9)] = D10_DC22(p0, self.f24)
        def iife162_(_pat_let41_0):
            def iife163_(d_2302_dt__update__tmp_h1_):
                def iife164_(_pat_let42_0):
                    def iife165_(d_2303_dt__update_hcf41_h0_):
                        return D10_DC22(d_2303_dt__update_hcf41_h0_, (d_2302_dt__update__tmp_h1_).cf42)
                    return iife165_(_pat_let42_0)
                return iife164_(pat_let_tv35_)
            return iife163_(_pat_let41_0)
        nw367_[int(10)] = iife162_(d_2299_v92_)
        nw367_[int(11)] = d_2299_v92_
        nw367_[int(12)] = d_2299_v92_
        nw367_[int(13)] = d_2299_v92_
        nw367_[int(14)] = D10_DC22(p0, self.f24)
        nw367_[int(15)] = d_2299_v92_
        def iife166_(_pat_let43_0):
            def iife167_(d_2304_dt__update__tmp_h2_):
                def iife168_(_pat_let44_0):
                    def iife169_(d_2305_dt__update_hcf42_h0_):
                        return D10_DC22((d_2304_dt__update__tmp_h2_).cf41, d_2305_dt__update_hcf42_h0_)
                    return iife169_(_pat_let44_0)
                return iife168_(not(False))
            return iife167_(_pat_let43_0)
        nw367_[int(16)] = iife166_(d_2299_v92_)
        nw367_[int(17)] = d_2299_v92_
        nw367_[int(18)] = d_2299_v92_
        d_2301_v93_ = nw367_
        index336_ = default__.safeIndex(384, (d_2298_v91_).length(0))
        (d_2298_v91_)[index336_] = d_2301_v93_
        index337_ = default__.safeIndex(384, (d_2298_v91_).length(0))
        rhs338_ = d_2297_v90_
        rhs339_ = d_2301_v93_
        lhs297_ = d_2298_v91_
        lhs298_ = default__.safeIndex(384, (d_2298_v91_).length(0))
        d_2297_v90_ = rhs338_
        lhs297_[lhs298_] = rhs339_
        d_2306_v94_: _dafny.Map
        d_2306_v94_ = _dafny.Map({p0: False})
        d_2307_v95_: str
        d_2307_v95_ = _dafny.CodePoint('e')
        d_2308_v96_: _dafny.Seq
        d_2308_v96_ = _dafny.SeqWithoutIsStrInference([d_2307_v95_])
        d_2309_v97_: _dafny.Set
        d_2309_v97_ = _dafny.Set({d_2185_v0_, d_2185_v0_})
        r0 = (d_2309_v97_ if ((d_2306_v94_)[len(d_2308_v96_)] if (len(d_2308_v96_)) in (d_2306_v94_) else self.f24) else (d_2309_v97_).intersection(_dafny.Set({d_2185_v0_})))
        return r0


class C8(T2, T1, T0):
    def  __init__(self):
        self._f24: bool = False
        self._f25: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C8"
    @property
    def f24(self):
        return self._f24
    @f24.setter
    def f24(self, value):
        self._f24 = value
    @property
    def f25(self):
        return self._f25
    def ctor__(self, f25, f24):
        (self)._f25 = f25
        (self).f24 = f24

    def fm6(self, p0, globalState):
        return 66

    def fm4(self, p0, p1, globalState):
        return D0_DC0((self).f25, default__.safeDivisionInt(474, -189), (273 if self.f24 else 984))

    def fm5(self, p0, p1, p2, globalState):
        source38_ = D2_DC3(_dafny.Map({-657: len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wow"))): _dafny.SeqWithoutIsStrInference([self.f24, self.f24])}))}))
        if source38_.is_DC4:
            d_2310___mcc_h0_ = source38_.cf9
            d_2311_cf9_ = d_2310___mcc_h0_
            return (331) - (599)
        elif source38_.is_DC5:
            d_2312___mcc_h1_ = source38_.cf10
            d_2313___mcc_h2_ = source38_.cf11
            d_2314___mcc_h3_ = source38_.cf12
            d_2315___mcc_h4_ = source38_.cf13
            d_2316_cf13_ = d_2315___mcc_h4_
            d_2317_cf12_ = d_2314___mcc_h3_
            d_2318_cf11_ = d_2313___mcc_h2_
            d_2319_cf10_ = d_2312___mcc_h1_
            return d_2317_cf12_
        elif source38_.is_DC3:
            d_2320___mcc_h5_ = source38_.cf8
            d_2321_cf8_ = d_2320___mcc_h5_
            if (self).f25:
                return (_dafny.MultiSet([654, 171, (_dafny.MultiSet([831, -255])).cardinality])).cardinality
            elif True:
                return 277
        elif True:
            d_2322___mcc_h6_ = source38_.cf14
            d_2323_cf14_ = d_2322___mcc_h6_
            return (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "durgk")))) + (306)

    def fm12(self, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "apmpv"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tfvnef")))

    def fm13(self, p0, p1, globalState):
        return ((D2_DC5(_dafny.CodePoint('u'), (self).f25, len(_dafny.SeqWithoutIsStrInference([-898])), len(_dafny.Map({408: 641})))).cf10) not in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_2324_i0_ in range(default__.abs(-520))])))

    def m0(self, globalState):
        r0: bool = False
        r1: _dafny.Array = _dafny.Array(None, 0)
        d_2325_v0_: int
        d_2325_v0_ = 28
        (globalState).f16 = d_2325_v0_
        d_2326_v1_: _dafny.Array
        nw368_ = _dafny.Array(_dafny.Map({}), 8)
        d_2326_v1_ = nw368_
        index338_ = default__.safeIndex(587, (d_2326_v1_).length(0))
        (d_2326_v1_)[index338_] = default__.fm14(globalState)
        d_2327_v2_: _dafny.Seq
        d_2327_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mwgmksent"))
        d_2328_v3_: D1
        d_2328_v3_ = D1_DC1(d_2327_v2_)
        d_2329_v4_: _dafny.Map
        d_2329_v4_ = _dafny.Map({d_2328_v3_: d_2325_v0_})
        index339_ = default__.safeIndex(587, (d_2326_v1_).length(0))
        (d_2326_v1_)[index339_] = _dafny.Map({d_2329_v4_: (self).f25})
        d_2330_v5_: _dafny.MultiSet
        d_2330_v5_ = _dafny.MultiSet([d_2325_v0_])
        d_2331_v6_: _dafny.Seq
        d_2331_v6_ = _dafny.SeqWithoutIsStrInference([(d_2330_v5_).cardinality, d_2325_v0_])
        d_2332_v7_: _dafny.Map
        d_2332_v7_ = _dafny.Map({(d_2331_v6_)[default__.safeIndex(d_2325_v0_, len(d_2331_v6_))]: d_2325_v0_})
        d_2333_v8_: D2
        d_2333_v8_ = D2_DC3(d_2332_v7_)
        d_2334_v9_: _dafny.Map
        d_2334_v9_ = _dafny.Map({(d_2333_v8_ if self.f24 else d_2333_v8_): D2_DC4(self.f24)})
        d_2334_v9_ = d_2334_v9_
        d_2335_v10_: _dafny.Seq
        d_2335_v10_ = _dafny.SeqWithoutIsStrInference([not(self.f24)])
        d_2336_v11_: D2
        d_2336_v11_ = D2_DC4((d_2335_v10_)[default__.safeIndex(d_2325_v0_, len(d_2335_v10_))])
        d_2337_v12_: _dafny.Array
        nw369_ = _dafny.Array(int(0), 12)
        d_2337_v12_ = nw369_
        d_2338_v13_: str
        d_2338_v13_ = _dafny.CodePoint('j')
        d_2339_v14_: D2
        d_2339_v14_ = D2_DC5(d_2338_v13_, (self).f25, d_2325_v0_, len(d_2331_v6_))
        d_2340_v15_: _dafny.Set
        d_2340_v15_ = _dafny.Set({self.f24})
        d_2341_v16_: _dafny.Seq
        d_2341_v16_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({self.f24, (self).f25}), d_2340_v15_, d_2340_v15_, d_2340_v15_])
        d_2342_v17_: _dafny.Map
        d_2342_v17_ = _dafny.Map({d_2341_v16_: self.f24})
        d_2343_v18_: _dafny.Array
        nw370_ = _dafny.Array(None, 28)
        nw370_[int(0)] = (d_2336_v11_).cf9
        nw370_[int(1)] = self.f24
        nw370_[int(2)] = (self).f25
        nw370_[int(3)] = True
        nw370_[int(4)] = (self).f25
        nw370_[int(5)] = (self).f25
        nw370_[int(6)] = self.f24
        nw370_[int(7)] = False
        nw370_[int(8)] = (self).f25
        nw370_[int(9)] = self.f24
        nw370_[int(10)] = (self).f25
        nw370_[int(11)] = self.f24
        nw370_[int(12)] = (self).f25
        nw370_[int(13)] = (self).f25
        nw370_[int(14)] = (D1_DC2(d_2337_v12_, d_2338_v13_, (d_2339_v14_).cf12, (self).f25)).cf7
        nw370_[int(15)] = (self).f25
        nw370_[int(16)] = (self).f25
        nw370_[int(17)] = (d_2335_v10_)[default__.safeIndex(d_2325_v0_, len(d_2335_v10_))]
        nw370_[int(18)] = (self).f25
        nw370_[int(19)] = False
        nw370_[int(20)] = ((d_2342_v17_)[d_2341_v16_] if (d_2341_v16_) in (d_2342_v17_) else (self).f25)
        nw370_[int(21)] = self.f24
        nw370_[int(22)] = self.f24
        nw370_[int(23)] = True
        nw370_[int(24)] = self.f24
        nw370_[int(25)] = (self).f25
        nw370_[int(26)] = not((self).f25)
        nw370_[int(27)] = (self).f25
        d_2343_v18_ = nw370_
        r1 = (d_2343_v18_ if self.f24 else d_2343_v18_)
        d_2344_v19_: D1
        d_2344_v19_ = D1_DC2((d_2337_v12_ if not((self).f25) else d_2337_v12_), d_2338_v13_, d_2325_v0_, True)
        source39_ = d_2344_v19_
        if source39_.is_DC2:
            d_2345___mcc_h0_ = source39_.cf4
            d_2346___mcc_h1_ = source39_.cf5
            d_2347___mcc_h2_ = source39_.cf6
            d_2348___mcc_h3_ = source39_.cf7
            d_2349_cf7_ = d_2348___mcc_h3_
            d_2350_cf6_ = d_2347___mcc_h2_
            d_2351_cf5_ = d_2346___mcc_h1_
            d_2352_cf4_ = d_2345___mcc_h0_
            d_2353_v20_: _dafny.Map
            d_2353_v20_ = _dafny.Map({d_2349_cf7_: d_2350_cf6_})
            d_2354_v21_: C4
            nw371_ = C4()
            nw371_.ctor__((self.f24) not in (d_2353_v20_), self.f24)
            d_2354_v21_ = nw371_
            r0 = d_2349_cf7_
            d_2355_v22_: _dafny.Seq
            d_2355_v22_ = _dafny.SeqWithoutIsStrInference([d_2327_v2_])
            d_2356_v23_: _dafny.Set
            d_2356_v23_ = _dafny.Set({d_2325_v0_})
            d_2357_v24_: _dafny.Map
            d_2357_v24_ = _dafny.Map({d_2349_cf7_: d_2349_cf7_})
            d_2358_v25_: C3
            nw372_ = C3()
            nw372_.ctor__((self).f25, (d_2355_v22_)[default__.safeIndex(len(d_2356_v23_), len(d_2355_v22_))], ((d_2357_v24_)[(self).f25] if ((self).f25) in (d_2357_v24_) else False), d_2351_cf5_, d_2349_cf7_, not((((d_2357_v24_)[self.f24] if (self.f24) in (d_2357_v24_) else (self).f25)) or ((self).fm13(d_2350_cf6_, self.f24, globalState))))
            d_2358_v25_ = nw372_
            index340_ = default__.safeIndex(930, (d_2343_v18_).length(0))
            (d_2343_v18_)[index340_] = d_2349_cf7_
            d_2359_v26_: _dafny.Map
            d_2359_v26_ = _dafny.Map({47: d_2337_v12_})
            index341_ = default__.safeIndex(930, (d_2343_v18_).length(0))
            (d_2343_v18_)[index341_] = (d_2359_v26_) != (d_2359_v26_)
        elif True:
            d_2360___mcc_h4_ = source39_.cf3
            d_2361_cf3_ = d_2360___mcc_h4_
            if ((self).f25) and (self.f24):
                d_2362_v27_: D3
                d_2362_v27_ = D3_DC7(d_2335_v10_)
                d_2363_v28_: _dafny.Seq
                d_2363_v28_ = _dafny.SeqWithoutIsStrInference([d_2362_v27_])
                d_2363_v28_ = _dafny.SeqWithoutIsStrInference([D3_DC7(d_2335_v10_) for d_2364_i0_ in range(default__.abs(-959))])
                d_2365_v29_: _dafny.Seq
                d_2365_v29_ = _dafny.SeqWithoutIsStrInference([d_2361_cf3_])
                (globalState).f11 = ((d_2330_v5_)[len(d_2365_v29_)] if (len(d_2365_v29_)) in (d_2330_v5_) else default__.safeModuloInt((0) - (d_2325_v0_), d_2325_v0_))
                d_2366_v30_: _dafny.MultiSet
                d_2366_v30_ = _dafny.MultiSet([d_2343_v18_])
                d_2367_v31_: _dafny.Map
                d_2367_v31_ = _dafny.Map({((d_2361_cf3_).set(default__.safeIndex(((d_2366_v30_)[d_2343_v18_] if (d_2343_v18_) in (d_2366_v30_) else 883), len(d_2361_cf3_)), d_2338_v13_)) + (d_2361_cf3_): (0) - (d_2325_v0_)})
                d_2368_v32_: C6
                nw373_ = C6()
                nw373_.ctor__((self).f25, self.f24, self.f24, d_2338_v13_, not ((self).f25) or (True))
                d_2368_v32_ = nw373_
                rhs340_ = d_2367_v31_
                rhs341_ = d_2368_v32_
                d_2367_v31_ = rhs340_
                d_2368_v32_ = rhs341_
                d_2369_v33_: _dafny.Map
                d_2369_v33_ = _dafny.Map({(self).f25: d_2325_v0_})
                d_2370_v34_: _dafny.Map
                d_2370_v34_ = _dafny.Map({d_2325_v0_: (d_2368_v32_).f28})
                d_2371_v35_: D0
                d_2371_v35_ = D0_DC0(((d_2370_v34_)[d_2325_v0_] if (d_2325_v0_) in (d_2370_v34_) else True), d_2325_v0_, d_2325_v0_)
                d_2372_v36_: _dafny.Seq
                d_2372_v36_ = _dafny.SeqWithoutIsStrInference([d_2371_v35_])
                d_2373_v37_: _dafny.Map
                d_2373_v37_ = _dafny.Map({(self).fm13(d_2325_v0_, (d_2368_v32_).fm7(d_2369_v33_, globalState), globalState): d_2372_v36_})
                d_2374_v38_: _dafny.Map
                d_2374_v38_ = _dafny.Map({(self).f25: (d_2368_v32_).fm5(d_2373_v37_, (d_2368_v32_).fm17((self).f25, self.f24, globalState), self.f24, globalState)})
                d_2375_v39_: _dafny.Set
                d_2375_v39_ = _dafny.Set({(0) - ((0) - (d_2325_v0_)), d_2325_v0_, d_2325_v0_, d_2325_v0_})
                d_2376_v40_: _dafny.Map
                d_2376_v40_ = _dafny.Map({(self).f25: d_2370_v34_})
                rhs342_ = d_2369_v33_
                rhs343_ = default__.fm53(d_2338_v13_, d_2338_v13_, globalState)
                rhs344_ = (647 if ((self).f25 if self.f24 else self.f24) else default__.safeDivisionInt(default__.fm0(d_2325_v0_, (d_2368_v32_).f28, ((d_2376_v40_)[(self).f25] if ((self).f25) in (d_2376_v40_) else d_2370_v34_), d_2325_v0_, globalState), d_2325_v0_))
                lhs299_ = globalState
                d_2369_v33_ = rhs342_
                d_2375_v39_ = rhs343_
                lhs299_.f22 = rhs344_
                index342_ = default__.safeIndex(339, (d_2337_v12_).length(0))
                (d_2337_v12_)[index342_] = d_2325_v0_
                index343_ = default__.safeIndex(339, (d_2337_v12_).length(0))
                (d_2337_v12_)[index343_] = d_2325_v0_
            elif True:
                (globalState).f10 = (779) == (d_2325_v0_)
                d_2377_v41_: _dafny.Map
                d_2377_v41_ = _dafny.Map({(d_2325_v0_) <= (len(d_2327_v2_)): self.f24})
                d_2377_v41_ = (d_2377_v41_).set(self.f24, (self).f25)
                d_2378_v42_: _dafny.Map
                d_2378_v42_ = _dafny.Map({d_2325_v0_: (self).f25})
                d_2379_v43_: _dafny.Set
                d_2379_v43_ = _dafny.Set({(0) - (default__.fm0(((d_2332_v7_)[d_2325_v0_] if (d_2325_v0_) in (d_2332_v7_) else d_2325_v0_), True, d_2378_v42_, 117, globalState)), d_2325_v0_, d_2325_v0_})
                (globalState).f16 = default__.safeDivisionInt((d_2325_v0_ if (self).f25 else d_2325_v0_), len(d_2379_v43_))
                d_2380_v44_: C1
                nw374_ = C1()
                nw374_.ctor__(self.f24)
                d_2380_v44_ = nw374_
                d_2380_v44_ = d_2380_v44_
                d_2381_v45_: C0
                nw375_ = C0()
                nw375_.ctor__((self).fm12(globalState))
                d_2381_v45_ = nw375_
            d_2382_v46_: D8
            d_2382_v46_ = D8_DC19((self).f25, d_2325_v0_)
            if (d_2382_v46_).cf37:
                d_2383_v47_: _dafny.Map
                d_2383_v47_ = _dafny.Map({106: d_2361_cf3_})
                d_2384_v48_: _dafny.Map
                d_2384_v48_ = _dafny.Map({(self).f25: d_2325_v0_})
                d_2385_v49_: _dafny.Map
                d_2385_v49_ = _dafny.Map({((d_2383_v47_)[((d_2384_v48_)[self.f24] if (self.f24) in (d_2384_v48_) else d_2325_v0_)] if (((d_2384_v48_)[self.f24] if (self.f24) in (d_2384_v48_) else d_2325_v0_)) in (d_2383_v47_) else default__.fm29(globalState)): d_2332_v7_})
                d_2385_v49_ = d_2385_v49_
                d_2386_v50_: D2
                d_2386_v50_ = D2_DC6(D2_DC4((d_2335_v10_)[default__.safeIndex(d_2325_v0_, len(d_2335_v10_))]))
                d_2387_v52_: D5
                def iife170_():
                    coll80_ = _dafny.Map()
                    compr_80_: int
                    for compr_80_ in (d_2330_v5_).Elements:
                        d_2388_v51_: int = compr_80_
                        if (d_2388_v51_) in (d_2330_v5_):
                            coll80_[default__.safeDivisionInt(d_2388_v51_, d_2325_v0_)] = d_2325_v0_
                    return _dafny.Map(coll80_)
                d_2387_v52_ = D5_DC11(d_2386_v50_, (d_2335_v10_)[default__.safeIndex(len(iife170_()
), len(d_2335_v10_))])
                d_2389_v53_: _dafny.Array
                def lambda110_(d_2390_v13_):
                    def lambda111_(d_2391_i1_):
                        return d_2390_v13_

                    return lambda111_

                init65_ = lambda110_(d_2338_v13_)
                nw376_ = _dafny.Array(None, 27)
                for i0_65_ in range(nw376_.length(0)):
                    nw376_[i0_65_] = init65_(i0_65_)
                d_2389_v53_ = nw376_
                index344_ = default__.safeIndex(38, (d_2389_v53_).length(0))
                (d_2389_v53_)[index344_] = d_2338_v13_
                index345_ = default__.safeIndex(894, (d_2343_v18_).length(0))
                (d_2343_v18_)[index345_] = self.f24
                index346_ = default__.safeIndex(38, (d_2389_v53_).length(0))
                index347_ = default__.safeIndex(894, (d_2343_v18_).length(0))
                rhs345_ = d_2361_cf3_
                rhs346_ = d_2387_v52_
                rhs347_ = (self).fm13(((0) - (d_2325_v0_)) * (d_2325_v0_), (self).f25, globalState)
                rhs348_ = d_2338_v13_
                rhs349_ = not(default__.fm1(self.f24, True, d_2325_v0_, globalState))
                lhs300_ = globalState
                lhs301_ = d_2389_v53_
                lhs302_ = default__.safeIndex(38, (d_2389_v53_).length(0))
                lhs303_ = d_2343_v18_
                lhs304_ = default__.safeIndex(894, (d_2343_v18_).length(0))
                d_2327_v2_ = rhs345_
                d_2387_v52_ = rhs346_
                lhs300_.f10 = rhs347_
                lhs301_[lhs302_] = rhs348_
                lhs303_[lhs304_] = rhs349_
                d_2325_v0_ = (d_2325_v0_) - (d_2325_v0_)
                d_2392_v54_: _dafny.MultiSet
                d_2392_v54_ = _dafny.MultiSet([(self).f25, self.f24])
                r0 = (d_2392_v54_).ispropersubset(d_2392_v54_)
                (globalState).f16 = d_2325_v0_
            elif True:
                (globalState).f2 = d_2343_v18_
                (globalState).f22 = default__.safeModuloInt(d_2325_v0_, d_2325_v0_)
                d_2393_v55_: C0
                nw377_ = C0()
                nw377_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "humyuha")))
                d_2393_v55_ = nw377_
                d_2394_v56_: _dafny.MultiSet
                d_2394_v56_ = _dafny.MultiSet([not(self.f24), (self).f25])
                d_2395_v57_: _dafny.Seq
                d_2395_v57_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(d_2394_v56_).cardinality])])
                d_2396_v58_: _dafny.Set
                d_2396_v58_ = _dafny.Set({_dafny.CodePoint('j')})
                d_2395_v57_ = (d_2395_v57_ if (d_2396_v58_).issubset(d_2396_v58_) else d_2395_v57_)
                d_2397_v59_: _dafny.Set
                d_2397_v59_ = _dafny.Set({d_2394_v56_})
                d_2398_v60_: C3
                nw378_ = C3()
                nw378_.ctor__((self).f25, (d_2393_v55_).f31, (d_2394_v56_).isdisjoint(d_2394_v56_), ((d_2393_v55_).f31)[default__.safeIndex(d_2325_v0_, len((d_2393_v55_).f31))], self.f24, (d_2397_v59_).ispropersubset(d_2397_v59_))
                d_2398_v60_ = nw378_
                nw379_ = C3()
                nw379_.ctor__((d_2398_v60_).f32, default__.fm29(globalState), not (self.f24) or (self.f24), d_2338_v13_, self.f24, (self).f25)
                d_2398_v60_ = nw379_
            d_2399_v61_: _dafny.Map
            d_2399_v61_ = _dafny.Map({(0) - (len(d_2361_cf3_)): d_2361_cf3_})
            index348_ = default__.safeIndex(501, (d_2337_v12_).length(0))
            (d_2337_v12_)[index348_] = len((d_2399_v61_) | (d_2399_v61_))
            index349_ = default__.safeIndex(16, (d_2343_v18_).length(0))
            (d_2343_v18_)[index349_] = self.f24
            index350_ = default__.safeIndex(501, (d_2337_v12_).length(0))
            index351_ = default__.safeIndex(16, (d_2343_v18_).length(0))
            rhs350_ = (self).f25
            rhs351_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "chp"))
            rhs352_ = len(d_2340_v15_)
            rhs353_ = (d_2333_v8_).cf8
            rhs354_ = ((d_2325_v0_) + (d_2325_v0_)) == (((d_2330_v5_)[d_2325_v0_] if (d_2325_v0_) in (d_2330_v5_) else 270))
            lhs305_ = globalState
            lhs306_ = d_2337_v12_
            lhs307_ = default__.safeIndex(501, (d_2337_v12_).length(0))
            lhs308_ = d_2343_v18_
            lhs309_ = default__.safeIndex(16, (d_2343_v18_).length(0))
            lhs305_.f10 = rhs350_
            d_2361_cf3_ = rhs351_
            lhs306_[lhs307_] = rhs352_
            d_2332_v7_ = rhs353_
            lhs308_[lhs309_] = rhs354_
            d_2400_v62_: _dafny.Set
            d_2400_v62_ = _dafny.Set({(0) - (d_2325_v0_)})
            d_2401_v63_: C3
            nw380_ = C3()
            nw380_.ctor__(False, (d_2327_v2_).set(default__.safeIndex(len(d_2400_v62_), len(d_2327_v2_)), d_2338_v13_), (self).f25, d_2338_v13_, (d_2343_v18_)[default__.safeIndex(16, (d_2343_v18_).length(0))], self.f24)
            d_2401_v63_ = nw380_
            d_2402_v64_: _dafny.MultiSet
            d_2402_v64_ = _dafny.MultiSet([d_2401_v63_])
            d_2403_v65_: _dafny.Seq
            d_2403_v65_ = _dafny.SeqWithoutIsStrInference([d_2401_v63_, d_2401_v63_, d_2401_v63_])
            d_2402_v64_ = (d_2402_v64_) | (_dafny.MultiSet([(d_2403_v65_)[default__.safeIndex((d_2337_v12_)[default__.safeIndex(501, (d_2337_v12_).length(0))], len(d_2403_v65_))]]))
        d_2404_v66_: _dafny.Array
        def lambda112_(d_2405_v0_):
            def lambda113_(d_2406_i2_):
                return D3_DC8((self).f25, 250, d_2405_v0_)

            return lambda113_

        init66_ = lambda112_(d_2325_v0_)
        nw381_ = _dafny.Array(None, 1)
        for i0_66_ in range(nw381_.length(0)):
            nw381_[i0_66_] = init66_(i0_66_)
        d_2404_v66_ = nw381_
        d_2407_v67_: _dafny.Map
        d_2407_v67_ = _dafny.Map({d_2404_v66_: d_2325_v0_})
        d_2407_v67_ = (_dafny.Map({d_2404_v66_: d_2325_v0_})) | (_dafny.Map({d_2404_v66_: len(_dafny.Set({(d_2335_v10_)[default__.safeIndex(d_2325_v0_, len(d_2335_v10_))], not((self).f25)}))}))
        r0 = default__.fm1(self.f24, (d_2340_v15_).ispropersubset(d_2340_v15_), default__.safeModuloInt(d_2325_v0_, (0) - (d_2325_v0_)), globalState)
        def lambda114_(d_2408_i3_):
            return (self).f25

        init67_ = lambda114_
        nw382_ = _dafny.Array(None, 14)
        for i0_67_ in range(nw382_.length(0)):
            nw382_[i0_67_] = init67_(i0_67_)
        r1 = nw382_
        return r0, r1

    def m1(self, p0, globalState):
        r0: str = _dafny.CodePoint('D')
        r1: _dafny.Seq = _dafny.Seq({})
        d_2409_v0_: int
        d_2409_v0_ = 217
        d_2410_v1_: _dafny.Seq
        d_2410_v1_ = _dafny.SeqWithoutIsStrInference([d_2409_v0_])
        d_2411_v2_: _dafny.Seq
        d_2411_v2_ = _dafny.SeqWithoutIsStrInference([d_2410_v1_])
        hi15_ = len(d_2411_v2_)
        for d_2412_i0_ in range(d_2409_v0_, hi15_):
            d_2413_v3_: str
            d_2413_v3_ = _dafny.CodePoint('c')
            d_2414_v4_: D7
            d_2414_v4_ = D7_DC17(D7_DC15(d_2413_v3_, self.f24))
            d_2414_v4_ = d_2414_v4_
            d_2415_v5_: _dafny.Array
            nw383_ = _dafny.Array(int(0), 23)
            d_2415_v5_ = nw383_
            d_2416_v6_: D0
            d_2416_v6_ = D0_DC0((self).f25, d_2409_v0_, d_2412_i0_)
            d_2417_v7_: _dafny.Seq
            d_2417_v7_ = _dafny.SeqWithoutIsStrInference([d_2416_v6_])
            d_2418_v8_: _dafny.Map
            d_2418_v8_ = _dafny.Map({False: d_2417_v7_})
            d_2419_v9_: _dafny.Map
            d_2419_v9_ = _dafny.Map({(self).fm5(d_2418_v8_, (self).f25, (self).f25, globalState): d_2413_v3_})
            index352_ = default__.safeIndex(554, (d_2415_v5_).length(0))
            (d_2415_v5_)[index352_] = len(d_2419_v9_)
            d_2420_v11_: _dafny.Map
            d_2420_v11_ = _dafny.Map({(0) - (-41): (self).f25})
            d_2421_v12_: _dafny.Seq
            d_2421_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ycyxxta"))
            d_2422_v13_: D15
            d_2422_v13_ = D15_DC33(d_2409_v0_)
            index353_ = default__.safeIndex(554, (d_2415_v5_).length(0))
            def iife171_():
                coll81_ = _dafny.Set()
                compr_81_: int
                for compr_81_ in _dafny.IntegerRange(-953, 560):
                    d_2423_v10_: int = compr_81_
                    if ((-953) <= (d_2423_v10_)) and ((d_2423_v10_) < (560)):
                        coll81_ = coll81_.union(_dafny.Set([default__.safeModuloInt(d_2423_v10_, d_2412_i0_)]))
                return _dafny.Set(coll81_)
            def iife172_():
                coll82_ = _dafny.Map()
                compr_82_: int
                for compr_82_ in _dafny.IntegerRange(767, -982):
                    d_2424_v14_: int = compr_82_
                    if ((767) <= (d_2424_v14_)) and ((d_2424_v14_) < (-982)):
                        coll82_[(d_2424_v14_) - (d_2412_i0_)] = self.f24
                return _dafny.Map(coll82_)
            rhs355_ = d_2413_v3_
            rhs356_ = default__.fm0((0) - (d_2412_i0_), (D11_DC26(self.f24, True, len(iife171_()
), (self).f25)).cf50, d_2420_v11_, d_2412_i0_, globalState)
            rhs357_ = (104) * (len((_dafny.SeqWithoutIsStrInference([len(d_2421_v12_)])).set(default__.safeIndex((0) - (d_2412_i0_), len(_dafny.SeqWithoutIsStrInference([len(d_2421_v12_)]))), 96)))
            rhs358_ = ((d_2422_v13_).cf55) > ((self).fm5(d_2418_v8_, (self).f25, self.f24, globalState))
            rhs359_ = default__.fm0(d_2412_i0_, ((self).f25) == (not(self.f24)), iife172_()
            , d_2409_v0_, globalState)
            lhs310_ = d_2415_v5_
            lhs311_ = default__.safeIndex(554, (d_2415_v5_).length(0))
            lhs312_ = globalState
            lhs313_ = globalState
            lhs314_ = globalState
            d_2413_v3_ = rhs355_
            lhs310_[lhs311_] = rhs356_
            lhs312_.f16 = rhs357_
            lhs313_.f10 = rhs358_
            lhs314_.f11 = rhs359_
            d_2425_v15_: _dafny.Array
            nw384_ = _dafny.Array(_dafny.Map({}), 27)
            d_2425_v15_ = nw384_
            d_2426_v16_: _dafny.Map
            d_2426_v16_ = _dafny.Map({_dafny.CodePoint('p'): d_2410_v1_})
            index354_ = default__.safeIndex(26, (d_2425_v15_).length(0))
            (d_2425_v15_)[index354_] = d_2426_v16_
            d_2427_v17_: _dafny.Array
            def lambda115_(d_2428_i1_):
                return self.f24

            init68_ = lambda115_
            nw385_ = _dafny.Array(None, 5)
            for i0_68_ in range(nw385_.length(0)):
                nw385_[i0_68_] = init68_(i0_68_)
            d_2427_v17_ = nw385_
            index355_ = default__.safeIndex(631, (d_2427_v17_).length(0))
            (d_2427_v17_)[index355_] = default__.fm1(False, self.f24, (0) - ((self).fm5(d_2418_v8_, (self).f25, self.f24, globalState)), globalState)
            index356_ = default__.safeIndex(26, (d_2425_v15_).length(0))
            index357_ = default__.safeIndex(554, (d_2415_v5_).length(0))
            index358_ = default__.safeIndex(631, (d_2427_v17_).length(0))
            rhs360_ = d_2426_v16_
            rhs361_ = self.f24
            rhs362_ = ((d_2415_v5_)[default__.safeIndex(554, (d_2415_v5_).length(0))]) <= (len(d_2421_v12_))
            rhs363_ = (0) - (d_2412_i0_)
            rhs364_ = self.f24
            lhs315_ = d_2425_v15_
            lhs316_ = default__.safeIndex(26, (d_2425_v15_).length(0))
            lhs317_ = globalState
            lhs318_ = self
            lhs319_ = d_2415_v5_
            lhs320_ = default__.safeIndex(554, (d_2415_v5_).length(0))
            lhs321_ = d_2427_v17_
            lhs322_ = default__.safeIndex(631, (d_2427_v17_).length(0))
            lhs315_[lhs316_] = rhs360_
            lhs317_.f10 = rhs361_
            lhs318_.f24 = rhs362_
            lhs319_[lhs320_] = rhs363_
            lhs321_[lhs322_] = rhs364_
            (globalState).f10 = ((812) - ((d_2415_v5_)[default__.safeIndex(554, (d_2415_v5_).length(0))])) <= ((len(_dafny.Set({d_2412_i0_}))) + (43))
        d_2429_i2_: int
        d_2429_i2_ = 0
        with _dafny.label("20"):
            while self.f24:
                with _dafny.c_label("20"):
                    if (d_2429_i2_) >= (100):
                        raise _dafny.Break("20")
                    d_2429_i2_ = (d_2429_i2_) + (1)
                    d_2430_v18_: _dafny.Seq
                    d_2430_v18_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jbegb"))
                    d_2431_v19_: C0
                    nw386_ = C0()
                    nw386_.ctor__(d_2430_v18_)
                    d_2431_v19_ = nw386_
                    d_2432_v20_: _dafny.Array
                    nw387_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 21)
                    d_2432_v20_ = nw387_
                    index359_ = default__.safeIndex(695, (d_2432_v20_).length(0))
                    (d_2432_v20_)[index359_] = d_2430_v18_
                    d_2433_v21_: _dafny.Map
                    d_2433_v21_ = _dafny.Map({d_2409_v0_: (self).f25})
                    index360_ = default__.safeIndex(695, (d_2432_v20_).length(0))
                    rhs365_ = (727) * ((0) - (d_2409_v0_))
                    rhs366_ = self.f24
                    rhs367_ = self.f24
                    rhs368_ = (d_2431_v19_).f31
                    rhs369_ = d_2433_v21_
                    lhs323_ = self
                    lhs324_ = globalState
                    lhs325_ = d_2432_v20_
                    lhs326_ = default__.safeIndex(695, (d_2432_v20_).length(0))
                    d_2409_v0_ = rhs365_
                    lhs323_.f24 = rhs366_
                    lhs324_.f10 = rhs367_
                    lhs325_[lhs326_] = rhs368_
                    d_2433_v21_ = rhs369_
                    d_2434_v22_: _dafny.Array
                    def lambda116_(d_2435_v0_):
                        def lambda117_(d_2436_i3_):
                            return default__.safeDivisionInt(d_2436_i3_, d_2435_v0_)

                        return lambda117_

                    init69_ = lambda116_(d_2409_v0_)
                    nw388_ = _dafny.Array(None, 16)
                    for i0_69_ in range(nw388_.length(0)):
                        nw388_[i0_69_] = init69_(i0_69_)
                    d_2434_v22_ = nw388_
                    index361_ = default__.safeIndex(230, (d_2434_v22_).length(0))
                    (d_2434_v22_)[index361_] = d_2409_v0_
                    index362_ = default__.safeIndex(230, (d_2434_v22_).length(0))
                    (d_2434_v22_)[index362_] = default__.safeDivisionInt(d_2409_v0_, d_2409_v0_)
                    d_2437_v23_: _dafny.Array
                    def lambda118_(d_2438_i4_):
                        return (_dafny.CodePoint('g') if self.f24 else _dafny.CodePoint('s'))

                    init70_ = lambda118_
                    nw389_ = _dafny.Array(None, 3)
                    for i0_70_ in range(nw389_.length(0)):
                        nw389_[i0_70_] = init70_(i0_70_)
                    d_2437_v23_ = nw389_
                    d_2439_v24_: str
                    d_2439_v24_ = _dafny.CodePoint('n')
                    index363_ = default__.safeIndex(248, (d_2437_v23_).length(0))
                    (d_2437_v23_)[index363_] = d_2439_v24_
                    d_2440_v25_: _dafny.Set
                    d_2440_v25_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([d_2439_v24_ for d_2441_i5_ in range(default__.abs(731))])})
                    index364_ = default__.safeIndex(230, (d_2434_v22_).length(0))
                    index365_ = default__.safeIndex(695, (d_2432_v20_).length(0))
                    index366_ = default__.safeIndex(230, (d_2434_v22_).length(0))
                    index367_ = default__.safeIndex(248, (d_2437_v23_).length(0))
                    rhs370_ = d_2409_v0_
                    rhs371_ = (d_2432_v20_)[default__.safeIndex(695, (d_2432_v20_).length(0))]
                    rhs372_ = (d_2440_v25_) != (_dafny.Set({d_2430_v18_}))
                    rhs373_ = d_2409_v0_
                    rhs374_ = d_2439_v24_
                    lhs327_ = d_2434_v22_
                    lhs328_ = default__.safeIndex(230, (d_2434_v22_).length(0))
                    lhs329_ = d_2432_v20_
                    lhs330_ = default__.safeIndex(695, (d_2432_v20_).length(0))
                    lhs331_ = globalState
                    lhs332_ = d_2434_v22_
                    lhs333_ = default__.safeIndex(230, (d_2434_v22_).length(0))
                    lhs334_ = d_2437_v23_
                    lhs335_ = default__.safeIndex(248, (d_2437_v23_).length(0))
                    lhs327_[lhs328_] = rhs370_
                    lhs329_[lhs330_] = rhs371_
                    lhs331_.f10 = rhs372_
                    lhs332_[lhs333_] = rhs373_
                    lhs334_[lhs335_] = rhs374_
                    d_2442_v26_: _dafny.Array
                    nw390_ = _dafny.Array(_dafny.Array(None, 0), 8)
                    d_2442_v26_ = nw390_
                    d_2443_v27_: _dafny.Array
                    nw391_ = _dafny.Array(None, 17)
                    nw391_[int(0)] = (self).f25
                    nw391_[int(1)] = self.f24
                    nw391_[int(2)] = self.f24
                    nw391_[int(3)] = (self).f25
                    nw391_[int(4)] = self.f24
                    nw391_[int(5)] = True
                    nw391_[int(6)] = (self).f25
                    nw391_[int(7)] = (self).f25
                    nw391_[int(8)] = self.f24
                    nw391_[int(9)] = self.f24
                    nw391_[int(10)] = (self).f25
                    nw391_[int(11)] = not((self).f25)
                    nw391_[int(12)] = not(self.f24)
                    nw391_[int(13)] = (self).f25
                    nw391_[int(14)] = (self).f25
                    nw391_[int(15)] = (self).f25
                    nw391_[int(16)] = (self).f25
                    d_2443_v27_ = nw391_
                    d_2444_v28_: _dafny.Map
                    d_2444_v28_ = _dafny.Map({(self).f25: d_2443_v27_})
                    d_2445_v29_: _dafny.Array
                    nw392_ = _dafny.Array(None, 17)
                    nw392_[int(0)] = d_2443_v27_
                    nw392_[int(1)] = d_2443_v27_
                    nw392_[int(2)] = d_2443_v27_
                    nw392_[int(3)] = d_2443_v27_
                    nw392_[int(4)] = d_2443_v27_
                    nw392_[int(5)] = d_2443_v27_
                    nw392_[int(6)] = d_2443_v27_
                    nw392_[int(7)] = d_2443_v27_
                    nw392_[int(8)] = d_2443_v27_
                    nw392_[int(9)] = ((d_2444_v28_)[self.f24] if (self.f24) in (d_2444_v28_) else d_2443_v27_)
                    nw392_[int(10)] = d_2443_v27_
                    nw392_[int(11)] = d_2443_v27_
                    nw392_[int(12)] = d_2443_v27_
                    nw392_[int(13)] = d_2443_v27_
                    nw392_[int(14)] = d_2443_v27_
                    nw392_[int(15)] = d_2443_v27_
                    nw392_[int(16)] = d_2443_v27_
                    d_2445_v29_ = nw392_
                    index368_ = default__.safeIndex(668, (d_2442_v26_).length(0))
                    (d_2442_v26_)[index368_] = d_2445_v29_
                    index369_ = default__.safeIndex(668, (d_2442_v26_).length(0))
                    (d_2442_v26_)[index369_] = d_2445_v29_
                    pass
            pass
        d_2446_v30_: _dafny.Array
        def lambda119_(d_2447_i6_):
            return self.f24

        init71_ = lambda119_
        nw393_ = _dafny.Array(None, 2)
        for i0_71_ in range(nw393_.length(0)):
            nw393_[i0_71_] = init71_(i0_71_)
        d_2446_v30_ = nw393_
        d_2448_v31_: _dafny.Array
        nw394_ = _dafny.Array(None, 4)
        nw394_[int(0)] = d_2446_v30_
        nw394_[int(1)] = d_2446_v30_
        nw394_[int(2)] = d_2446_v30_
        nw394_[int(3)] = d_2446_v30_
        d_2448_v31_ = nw394_
        d_2449_v32_: _dafny.Array
        d_2449_v32_ = d_2448_v31_
        source40_ = d_2449_v32_
        d_2450___mcc_h0_ = source40_
        d_2451_cf19_ = d_2450___mcc_h0_
        d_2452_v33_: _dafny.Seq
        d_2452_v33_ = _dafny.SeqWithoutIsStrInference([self.f24, self.f24])
        d_2453_v34_: _dafny.Map
        d_2453_v34_ = _dafny.Map({d_2452_v33_: len(_dafny.SeqWithoutIsStrInference([self.f24, self.f24]))})
        d_2454_v35_: _dafny.Map
        d_2454_v35_ = _dafny.Map({(0) - (d_2409_v0_): ((d_2453_v34_)[d_2452_v33_] if (d_2452_v33_) in (d_2453_v34_) else d_2409_v0_)})
        d_2455_v36_: str
        d_2455_v36_ = _dafny.CodePoint('e')
        d_2456_v37_: _dafny.Seq
        d_2456_v37_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hgxp"))
        d_2457_v38_: _dafny.Map
        d_2457_v38_ = _dafny.Map({d_2409_v0_: self.f24})
        d_2458_v39_: _dafny.Map
        d_2458_v39_ = _dafny.Map({(self).f25: default__.fm21(d_2456_v37_, ((d_2457_v38_)[d_2409_v0_] if (d_2409_v0_) in (d_2457_v38_) else not((self).f25)), d_2409_v0_, d_2409_v0_, globalState)})
        d_2459_v40_: _dafny.Array
        nw395_ = _dafny.Array(None, 6)
        nw395_[int(0)] = _dafny.CodePoint('n')
        nw395_[int(1)] = d_2455_v36_
        nw395_[int(2)] = d_2455_v36_
        nw395_[int(3)] = default__.fm47(globalState)
        nw395_[int(4)] = ((d_2458_v39_)[(self).f25] if ((self).f25) in (d_2458_v39_) else (d_2456_v37_)[default__.safeIndex(len(d_2410_v1_), len(d_2456_v37_))])
        nw395_[int(5)] = d_2455_v36_
        d_2459_v40_ = nw395_
        index370_ = default__.safeIndex(1, (d_2459_v40_).length(0))
        (d_2459_v40_)[index370_] = d_2455_v36_
        d_2460_v41_: C7
        nw396_ = C7()
        nw396_.ctor__((self).f25)
        d_2460_v41_ = nw396_
        index371_ = default__.safeIndex(1, (d_2459_v40_).length(0))
        rhs375_ = d_2454_v35_
        rhs376_ = d_2455_v36_
        rhs377_ = (d_2409_v0_) * (-112)
        rhs378_ = ((_dafny.SeqWithoutIsStrInference([d_2409_v0_, d_2409_v0_, d_2409_v0_, d_2409_v0_, d_2409_v0_])) + (_dafny.SeqWithoutIsStrInference([d_2409_v0_]))) <= ((D6_DC12(d_2410_v1_)).cf23)
        rhs379_ = d_2460_v41_
        lhs336_ = d_2459_v40_
        lhs337_ = default__.safeIndex(1, (d_2459_v40_).length(0))
        lhs338_ = globalState
        lhs339_ = globalState
        d_2454_v35_ = rhs375_
        lhs336_[lhs337_] = rhs376_
        lhs338_.f16 = rhs377_
        lhs339_.f10 = rhs378_
        d_2460_v41_ = rhs379_
        d_2461_v42_: _dafny.Array
        nw397_ = _dafny.Array(D2.default()(), 24)
        d_2461_v42_ = nw397_
        d_2461_v42_ = (d_2461_v42_ if self.f24 else (d_2461_v42_ if (self).f25 else d_2461_v42_))
        d_2462_v43_: _dafny.Set
        d_2462_v43_ = _dafny.Set({self.f24})
        d_2462_v43_ = d_2462_v43_
        (globalState).f22 = len((d_2456_v37_) + (d_2456_v37_))
        (globalState).f22 = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_2463_i7_ in range(default__.abs(308))]))
        d_2464_v44_: D10
        d_2464_v44_ = D10_DC22(((0) - ((0) - (d_2409_v0_)) if default__.fm1(self.f24, self.f24, d_2409_v0_, globalState) else d_2409_v0_), self.f24)
        source41_ = d_2464_v44_
        if source41_.is_DC22:
            d_2465___mcc_h1_ = source41_.cf41
            d_2466___mcc_h2_ = source41_.cf42
            d_2467_cf42_ = d_2466___mcc_h2_
            d_2468_cf41_ = d_2465___mcc_h1_
            d_2469_v45_: _dafny.Seq
            d_2469_v45_ = _dafny.SeqWithoutIsStrInference([self.f24, self.f24])
            d_2470_v46_: _dafny.Map
            d_2470_v46_ = _dafny.Map({(self).f25: d_2409_v0_})
            rhs380_ = ((_dafny.SeqWithoutIsStrInference([(self).f25, not(self.f24)])) + (d_2469_v45_)) + ((d_2469_v45_) + ((D3_DC7(_dafny.SeqWithoutIsStrInference([d_2467_cf42_]))).cf15))
            rhs381_ = (d_2409_v0_) * (default__.safeDivisionInt(d_2468_cf41_, len(d_2470_v46_)))
            rhs382_ = d_2410_v1_
            rhs383_ = (False) == (d_2467_cf42_)
            rhs384_ = d_2410_v1_
            lhs340_ = globalState
            lhs341_ = self
            r1 = rhs380_
            lhs340_.f22 = rhs381_
            d_2410_v1_ = rhs382_
            lhs341_.f24 = rhs383_
            d_2410_v1_ = rhs384_
            if d_2467_cf42_:
                d_2471_v47_: str
                d_2471_v47_ = _dafny.CodePoint('k')
                d_2472_v48_: _dafny.Map
                d_2472_v48_ = _dafny.Map({d_2471_v47_: self.f24})
                d_2472_v48_ = (d_2472_v48_).set(d_2471_v47_, self.f24)
                d_2473_v49_: _dafny.Map
                d_2473_v49_ = _dafny.Map({d_2467_cf42_: not((d_2468_cf41_) < (d_2468_cf41_))})
                (globalState).f11 = len(d_2473_v49_)
                d_2474_v50_: _dafny.Seq
                d_2474_v50_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gcndkvk"))
                d_2475_v51_: _dafny.Map
                d_2475_v51_ = _dafny.Map({(self).f25: d_2474_v50_})
                d_2476_v52_: _dafny.Set
                d_2476_v52_ = _dafny.Set({len(d_2475_v51_), d_2409_v0_, (0) - ((d_2409_v0_ if (self).f25 else 776))})
                d_2477_v53_: _dafny.Array
                nw398_ = _dafny.Array(None, 2)
                nw398_[int(0)] = d_2474_v50_
                nw398_[int(1)] = (d_2474_v50_).set(default__.safeIndex(290, len(d_2474_v50_)), d_2471_v47_)
                d_2477_v53_ = nw398_
                d_2478_v54_: _dafny.Array
                nw399_ = _dafny.Array(_dafny.Seq({}), 4)
                d_2478_v54_ = nw399_
                d_2479_v55_: _dafny.MultiSet
                d_2479_v55_ = _dafny.MultiSet([(self).f25, self.f24, (self).f25])
                d_2480_v56_: _dafny.MultiSet
                d_2480_v56_ = _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True, d_2467_cf42_]))
                d_2481_v57_: _dafny.Seq
                d_2481_v57_ = _dafny.SeqWithoutIsStrInference([d_2479_v55_, d_2479_v55_, (d_2479_v55_)])
                index372_ = default__.safeIndex(396, (d_2478_v54_).length(0))
                (d_2478_v54_)[index372_] = (d_2481_v57_) + (d_2481_v57_)
                d_2482_v58_: _dafny.MultiSet
                d_2482_v58_ = _dafny.MultiSet([d_2468_cf41_])
                d_2483_v59_: _dafny.Array
                nw400_ = _dafny.Array(D11.default()(), 20)
                d_2483_v59_ = nw400_
                d_2484_v60_: _dafny.Map
                d_2484_v60_ = _dafny.Map({d_2483_v59_: (d_2479_v55_).cardinality})
                d_2485_v61_: _dafny.Array
                nw401_ = _dafny.Array(_dafny.Map({}), 12)
                d_2485_v61_ = nw401_
                d_2486_v62_: D8
                d_2486_v62_ = D8_DC18(d_2485_v61_)
                pat_let_tv36_ = d_2485_v61_
                d_2487_v63_: _dafny.Seq
                def iife173_(_pat_let45_0):
                    def iife174_(d_2488_dt__update__tmp_h1_):
                        def iife175_(_pat_let46_0):
                            def iife176_(d_2489_dt__update_hcf36_h0_):
                                return D8_DC18(d_2489_dt__update_hcf36_h0_)
                            return iife176_(_pat_let46_0)
                        return iife175_(pat_let_tv36_)
                    return iife174_(_pat_let45_0)
                d_2487_v63_ = _dafny.SeqWithoutIsStrInference([d_2486_v62_, iife173_(D8_DC18(d_2485_v61_))])
                index373_ = default__.safeIndex(396, (d_2478_v54_).length(0))
                rhs385_ = (default__.fm53(d_2471_v47_, d_2471_v47_, globalState)) - ((d_2476_v52_) - (_dafny.Set({(d_2482_v58_).cardinality, d_2468_cf41_})))
                rhs386_ = ((d_2484_v60_)[d_2483_v59_] if (d_2483_v59_) in (d_2484_v60_) else default__.safeModuloInt(len(d_2487_v63_), -311))
                rhs387_ = d_2477_v53_
                rhs388_ = not((d_2409_v0_) >= (d_2468_cf41_))
                rhs389_ = d_2481_v57_
                lhs342_ = globalState
                lhs343_ = globalState
                lhs344_ = d_2478_v54_
                lhs345_ = default__.safeIndex(396, (d_2478_v54_).length(0))
                d_2476_v52_ = rhs385_
                lhs342_.f22 = rhs386_
                d_2477_v53_ = rhs387_
                lhs343_.f10 = rhs388_
                lhs344_[lhs345_] = rhs389_
                (globalState).f10 = d_2467_cf42_
                d_2490_v64_: _dafny.Map
                d_2490_v64_ = _dafny.Map({len(d_2410_v1_): False})
                d_2491_v65_: D0
                d_2491_v65_ = D0_DC0(True, d_2409_v0_, d_2468_cf41_)
                d_2492_v66_: _dafny.Map
                d_2492_v66_ = _dafny.Map({not(False): (d_2410_v1_).set(default__.safeIndex(len(d_2474_v50_), len(d_2410_v1_)), (self).fm6(d_2491_v65_, globalState))})
                d_2493_v67_: _dafny.Set
                d_2493_v67_ = _dafny.Set({self.f24})
                d_2494_v68_: _dafny.Array
                nw402_ = _dafny.Array(None, 19)
                nw402_[int(0)] = (d_2410_v1_) + (d_2410_v1_)
                nw402_[int(1)] = d_2410_v1_
                nw402_[int(2)] = d_2410_v1_
                nw402_[int(3)] = d_2410_v1_
                nw402_[int(4)] = (d_2410_v1_) + (d_2410_v1_)
                nw402_[int(5)] = (d_2410_v1_) + ((d_2410_v1_).set(default__.safeIndex(d_2468_cf41_, len(d_2410_v1_)), d_2409_v0_))
                nw402_[int(6)] = _dafny.SeqWithoutIsStrInference([(0) - (d_2409_v0_), len(d_2490_v64_)])
                nw402_[int(7)] = (_dafny.SeqWithoutIsStrInference([d_2409_v0_])) + (d_2410_v1_)
                nw402_[int(8)] = _dafny.SeqWithoutIsStrInference([d_2468_cf41_ for d_2495_i8_ in range(default__.abs(866))])
                nw402_[int(9)] = ((d_2492_v66_)[(self).f25] if ((self).f25) in (d_2492_v66_) else d_2410_v1_)
                nw402_[int(10)] = d_2410_v1_
                nw402_[int(11)] = d_2410_v1_
                nw402_[int(12)] = d_2410_v1_
                nw402_[int(13)] = d_2410_v1_
                nw402_[int(14)] = ((d_2410_v1_).set(default__.safeIndex(d_2409_v0_, len(d_2410_v1_)), 348)) + (d_2410_v1_)
                nw402_[int(15)] = d_2410_v1_
                nw402_[int(16)] = (_dafny.SeqWithoutIsStrInference([len(d_2493_v67_), d_2409_v0_, d_2468_cf41_, d_2409_v0_, ((d_2470_v46_)[(self).fm13(len(d_2410_v1_), d_2467_cf42_, globalState)] if ((self).fm13(len(d_2410_v1_), d_2467_cf42_, globalState)) in (d_2470_v46_) else d_2468_cf41_)])).set(default__.safeIndex(752, len(_dafny.SeqWithoutIsStrInference([len(d_2493_v67_), d_2409_v0_, d_2468_cf41_, d_2409_v0_, ((d_2470_v46_)[(self).fm13(len(d_2410_v1_), d_2467_cf42_, globalState)] if ((self).fm13(len(d_2410_v1_), d_2467_cf42_, globalState)) in (d_2470_v46_) else d_2468_cf41_)]))), d_2409_v0_)
                nw402_[int(17)] = d_2410_v1_
                nw402_[int(18)] = _dafny.SeqWithoutIsStrInference([d_2409_v0_])
                d_2494_v68_ = nw402_
                index374_ = default__.safeIndex(125, (d_2494_v68_).length(0))
                (d_2494_v68_)[index374_] = (d_2410_v1_) + (d_2410_v1_)
                index375_ = default__.safeIndex(125, (d_2494_v68_).length(0))
                (d_2494_v68_)[index375_] = d_2410_v1_
            elif True:
                d_2496_v69_: _dafny.Map
                d_2496_v69_ = _dafny.Map({220: self.f24})
                d_2497_v70_: str
                d_2497_v70_ = _dafny.CodePoint('p')
                d_2498_v71_: _dafny.Map
                d_2498_v71_ = _dafny.Map({d_2497_v70_: d_2468_cf41_})
                d_2499_v72_: _dafny.Map
                d_2499_v72_ = _dafny.Map({d_2467_cf42_: d_2496_v69_})
                d_2500_v73_: _dafny.Map
                d_2500_v73_ = _dafny.Map({d_2409_v0_: len(d_2499_v72_)})
                d_2501_v74_: _dafny.Map
                d_2501_v74_ = _dafny.Map({d_2500_v73_: True})
                d_2502_v75_: _dafny.Array
                nw403_ = _dafny.Array(None, 19)
                nw403_[int(0)] = len(d_2469_v45_)
                nw403_[int(1)] = d_2409_v0_
                nw403_[int(2)] = d_2468_cf41_
                nw403_[int(3)] = len((d_2410_v1_).set(default__.safeIndex(d_2409_v0_, len(d_2410_v1_)), len(_dafny.Set({d_2468_cf41_, default__.fm0(-913, (self).f25, d_2496_v69_, d_2409_v0_, globalState)}))))
                nw403_[int(4)] = (0) - (len(d_2498_v71_))
                nw403_[int(5)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rulfk")))
                nw403_[int(6)] = d_2468_cf41_
                nw403_[int(7)] = len(_dafny.SeqWithoutIsStrInference([816]))
                nw403_[int(8)] = d_2468_cf41_
                nw403_[int(9)] = d_2409_v0_
                nw403_[int(10)] = d_2409_v0_
                nw403_[int(11)] = len(d_2501_v74_)
                nw403_[int(12)] = d_2468_cf41_
                nw403_[int(13)] = d_2468_cf41_
                nw403_[int(14)] = 122
                nw403_[int(15)] = d_2468_cf41_
                nw403_[int(16)] = -237
                nw403_[int(17)] = d_2468_cf41_
                nw403_[int(18)] = d_2409_v0_
                d_2502_v75_ = nw403_
                d_2503_v76_: _dafny.Seq
                d_2503_v76_ = _dafny.SeqWithoutIsStrInference([d_2497_v70_, d_2497_v70_])
                d_2504_v77_: D1
                d_2504_v77_ = D1_DC2(d_2502_v75_, (d_2503_v76_)[default__.safeIndex(226, len(d_2503_v76_))], d_2409_v0_, self.f24)
                r0 = (d_2504_v77_).cf5
                d_2505_v78_: _dafny.Set
                d_2505_v78_ = _dafny.Set({d_2502_v75_})
                d_2506_v79_: C4
                nw404_ = C4()
                nw404_.ctor__(d_2467_cf42_, (d_2505_v78_).ispropersubset(d_2505_v78_))
                d_2506_v79_ = nw404_
                index376_ = default__.safeIndex(985, (d_2502_v75_).length(0))
                def iife177_():
                    coll83_ = _dafny.Map()
                    compr_83_: int
                    for compr_83_ in (default__.fm53(d_2497_v70_, d_2497_v70_, globalState)).Elements:
                        d_2507_v80_: int = compr_83_
                        if (d_2507_v80_) in (default__.fm53(d_2497_v70_, d_2497_v70_, globalState)):
                            coll83_[(d_2507_v80_) - (757)] = True
                    return _dafny.Map(coll83_)
                (d_2502_v75_)[index376_] = (len(iife177_()
                )) * (d_2468_cf41_)
                index377_ = default__.safeIndex(985, (d_2502_v75_).length(0))
                (d_2502_v75_)[index377_] = d_2409_v0_
                d_2508_v81_: _dafny.MultiSet
                d_2508_v81_ = _dafny.MultiSet([d_2503_v76_])
                d_2509_v82_: _dafny.Array
                def lambda120_(d_2510_v70_):
                    def lambda121_(d_2511_i9_):
                        return d_2510_v70_

                    return lambda121_

                init72_ = lambda120_(d_2497_v70_)
                nw405_ = _dafny.Array(None, 4)
                for i0_72_ in range(nw405_.length(0)):
                    nw405_[i0_72_] = init72_(i0_72_)
                d_2509_v82_ = nw405_
                index378_ = default__.safeIndex(320, (d_2509_v82_).length(0))
                (d_2509_v82_)[index378_] = d_2497_v70_
                index379_ = default__.safeIndex(320, (d_2509_v82_).length(0))
                rhs390_ = d_2508_v81_
                rhs391_ = d_2467_cf42_
                rhs392_ = d_2467_cf42_
                rhs393_ = d_2497_v70_
                lhs346_ = self
                lhs347_ = globalState
                lhs348_ = d_2509_v82_
                lhs349_ = default__.safeIndex(320, (d_2509_v82_).length(0))
                d_2508_v81_ = rhs390_
                lhs346_.f24 = rhs391_
                lhs347_.f10 = rhs392_
                lhs348_[lhs349_] = rhs393_
                (self).f24 = not ((self).f25) or ((self).f25)
            d_2512_v83_: _dafny.Array
            nw406_ = _dafny.Array(_dafny.Map({}), 23)
            d_2512_v83_ = nw406_
            d_2512_v83_ = d_2512_v83_
            if not (self.f24) or (d_2467_cf42_):
                d_2513_v84_: _dafny.Seq
                d_2513_v84_ = _dafny.SeqWithoutIsStrInference([d_2470_v46_, d_2470_v46_, d_2470_v46_, d_2470_v46_])
                d_2514_v85_: _dafny.Map
                d_2514_v85_ = _dafny.Map({(self).f25: (d_2470_v46_) | ((d_2513_v84_)[default__.safeIndex(d_2409_v0_, len(d_2513_v84_))])})
                d_2514_v85_ = (d_2514_v85_).set(self.f24, d_2470_v46_)
                d_2515_v86_: _dafny.Seq
                d_2515_v86_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_2516_i10_ in range(default__.abs(868))])])
                d_2517_v87_: _dafny.Map
                d_2517_v87_ = _dafny.Map({d_2515_v86_: p0})
                d_2518_v88_: _dafny.Seq
                d_2518_v88_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x"))
                d_2519_v89_: str
                d_2519_v89_ = _dafny.CodePoint('v')
                d_2520_v90_: C3
                nw407_ = C3()
                nw407_.ctor__((p0).issubset(((d_2517_v87_)[d_2515_v86_] if (d_2515_v86_) in (d_2517_v87_) else p0)), d_2518_v88_, self.f24, d_2519_v89_, self.f24, False)
                d_2520_v90_ = nw407_
                d_2521_v91_: C0
                nw408_ = C0()
                nw408_.ctor__((d_2520_v90_).f33)
                d_2521_v91_ = nw408_
                d_2522_v92_: _dafny.Array
                nw409_ = _dafny.Array(_dafny.MultiSet({}), 11)
                d_2522_v92_ = nw409_
                d_2523_v93_: _dafny.Map
                d_2523_v93_ = _dafny.Map({d_2522_v92_: d_2409_v0_})
                (globalState).f10 = (d_2522_v92_) not in (d_2523_v93_)
                d_2409_v0_ = (0) - (default__.safeDivisionInt(d_2409_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "itiiyhl")))))
            elif True:
                pat_let_tv37_ = d_2468_cf41_
                def iife178_(_pat_let47_0):
                    def iife179_(d_2524_dt__update__tmp_h2_):
                        def iife180_(_pat_let48_0):
                            def iife181_(d_2525_dt__update_hcf41_h0_):
                                return D10_DC22(d_2525_dt__update_hcf41_h0_, (d_2524_dt__update__tmp_h2_).cf42)
                            return iife181_(_pat_let48_0)
                        return iife180_(pat_let_tv37_)
                    return iife179_(_pat_let47_0)
                d_2464_v44_ = iife178_(d_2464_v44_)
                d_2526_v94_: _dafny.MultiSet
                d_2526_v94_ = _dafny.MultiSet([d_2467_cf42_, self.f24])
                d_2527_v95_: _dafny.Map
                d_2527_v95_ = _dafny.Map({default__.fm1((self).f25, (self).f25, d_2409_v0_, globalState): (d_2526_v94_) - (d_2526_v94_)})
                d_2527_v95_ = (d_2527_v95_).set(self.f24, (d_2526_v94_).set(False, default__.abs(d_2409_v0_)))
                rhs394_ = (d_2409_v0_ if d_2467_cf42_ else 369)
                rhs395_ = (_dafny.MultiSet([d_2467_cf42_])).cardinality
                lhs350_ = globalState
                lhs351_ = globalState
                lhs350_.f11 = rhs394_
                lhs351_.f22 = rhs395_
                d_2528_v96_: _dafny.Array
                def lambda122_(d_2529_i11_):
                    return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vqrej"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sadsjaej")))

                init73_ = lambda122_
                nw410_ = _dafny.Array(None, 25)
                for i0_73_ in range(nw410_.length(0)):
                    nw410_[i0_73_] = init73_(i0_73_)
                d_2528_v96_ = nw410_
                d_2530_v97_: _dafny.Seq
                d_2530_v97_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mrbutxxs"))
                index380_ = default__.safeIndex(726, (d_2528_v96_).length(0))
                (d_2528_v96_)[index380_] = d_2530_v97_
                d_2531_v98_: str
                d_2531_v98_ = _dafny.CodePoint('y')
                d_2532_v99_: _dafny.Set
                d_2532_v99_ = _dafny.Set({d_2531_v98_, d_2531_v98_})
                index381_ = default__.safeIndex(726, (d_2528_v96_).length(0))
                rhs396_ = (d_2468_cf41_) > (d_2409_v0_)
                rhs397_ = len(d_2530_v97_)
                rhs398_ = (d_2530_v97_ if True else (d_2530_v97_) + (d_2530_v97_))
                rhs399_ = len(p0)
                rhs400_ = (default__.safeDivisionInt(d_2409_v0_, len(d_2532_v99_))) - (d_2468_cf41_)
                lhs352_ = globalState
                lhs353_ = globalState
                lhs354_ = d_2528_v96_
                lhs355_ = default__.safeIndex(726, (d_2528_v96_).length(0))
                lhs356_ = globalState
                lhs357_ = globalState
                lhs352_.f10 = rhs396_
                lhs353_.f16 = rhs397_
                lhs354_[lhs355_] = rhs398_
                lhs356_.f22 = rhs399_
                lhs357_.f16 = rhs400_
                index382_ = default__.safeIndex(726, (d_2528_v96_).length(0))
                (d_2528_v96_)[index382_] = (d_2530_v97_).set(default__.safeIndex(121, len(d_2530_v97_)), d_2531_v98_)
        elif source41_.is_DC23:
            d_2533___mcc_h3_ = source41_.cf43
            d_2534___mcc_h4_ = source41_.cf44
            d_2535_cf44_ = d_2534___mcc_h4_
            d_2536_cf43_ = d_2533___mcc_h3_
            if (self).f25:
                d_2537_v100_: D0
                d_2537_v100_ = D0_DC0(d_2535_cf44_, (0) - (d_2409_v0_), d_2409_v0_)
                (globalState).f16 = (self).fm6(d_2537_v100_, globalState)
                d_2538_v102_: D6
                d_2538_v102_ = D6_DC12(d_2410_v1_)
                d_2539_v103_: _dafny.Array
                nw411_ = _dafny.Array(None, 15)
                nw411_[int(0)] = ((d_2410_v1_) + (d_2410_v1_)).set(default__.safeIndex((0) - (d_2409_v0_), len((d_2410_v1_) + (d_2410_v1_))), d_2409_v0_)
                nw411_[int(1)] = d_2410_v1_
                nw411_[int(2)] = d_2410_v1_
                nw411_[int(3)] = ((d_2410_v1_).set(default__.safeIndex(d_2409_v0_, len(d_2410_v1_)), d_2409_v0_) if self.f24 else _dafny.SeqWithoutIsStrInference([d_2409_v0_, d_2409_v0_]))
                nw411_[int(4)] = (d_2410_v1_ if self.f24 else d_2410_v1_)
                nw411_[int(5)] = (d_2410_v1_) + (default__.fm42(globalState))
                nw411_[int(6)] = d_2410_v1_
                nw411_[int(7)] = (default__.fm42(globalState)) + (d_2410_v1_)
                nw411_[int(8)] = d_2410_v1_
                nw411_[int(9)] = ((d_2410_v1_) + (d_2410_v1_)).set(default__.safeIndex(d_2409_v0_, len((d_2410_v1_) + (d_2410_v1_))), d_2409_v0_)
                def iife182_():
                    coll84_ = _dafny.Map()
                    compr_84_: int
                    for compr_84_ in ((D2_DC3(_dafny.Map({d_2409_v0_: d_2409_v0_}))).cf8).keys.Elements:
                        d_2540_v101_: int = compr_84_
                        if (d_2540_v101_) in ((D2_DC3(_dafny.Map({d_2409_v0_: d_2409_v0_}))).cf8):
                            coll84_[(d_2540_v101_) * (d_2409_v0_)] = True
                    return _dafny.Map(coll84_)
                nw411_[int(10)] = _dafny.SeqWithoutIsStrInference([len(iife182_()
                )])
                nw411_[int(11)] = d_2410_v1_
                nw411_[int(12)] = d_2410_v1_
                nw411_[int(13)] = d_2410_v1_
                nw411_[int(14)] = (d_2538_v102_).cf23
                d_2539_v103_ = nw411_
                d_2539_v103_ = d_2539_v103_
                d_2541_v104_: _dafny.Map
                d_2541_v104_ = _dafny.Map({(self).f25: d_2409_v0_})
                d_2542_v105_: _dafny.Seq
                d_2542_v105_ = _dafny.SeqWithoutIsStrInference([d_2541_v104_])
                d_2543_v106_: str
                d_2543_v106_ = _dafny.CodePoint('m')
                d_2544_v107_: _dafny.Seq
                d_2544_v107_ = _dafny.SeqWithoutIsStrInference([d_2536_cf43_, d_2536_cf43_])
                d_2545_v108_: _dafny.MultiSet
                d_2545_v108_ = _dafny.MultiSet([d_2536_cf43_, ((d_2536_cf43_) + (d_2536_cf43_)).set(default__.safeIndex(len(d_2542_v105_), len((d_2536_cf43_) + (d_2536_cf43_))), (self).f25), (default__.fm33(d_2543_v106_, d_2409_v0_, self.f24, self.f24, globalState)) + (d_2536_cf43_), d_2536_cf43_, (d_2536_cf43_) + ((d_2544_v107_)[default__.safeIndex(len(d_2541_v104_), len(d_2544_v107_))])])
                d_2545_v108_ = (d_2545_v108_).set(_dafny.SeqWithoutIsStrInference([self.f24, (self).f25, (self).f25]), default__.abs(default__.safeDivisionInt((d_2464_v44_).cf41, d_2409_v0_)))
                d_2535_cf44_ = d_2535_cf44_
                d_2546_v109_: _dafny.Array
                nw412_ = _dafny.Array(int(0), 27)
                d_2546_v109_ = nw412_
                (globalState).f7 = d_2546_v109_
            elif True:
                d_2547_v110_: C4
                nw413_ = C4()
                nw413_.ctor__(self.f24, (d_2409_v0_) > (d_2409_v0_))
                d_2547_v110_ = nw413_
                d_2548_v111_: _dafny.Set
                d_2548_v111_ = _dafny.Set({d_2410_v1_})
                d_2549_v112_: _dafny.Seq
                d_2549_v112_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rrpsm"))
                d_2550_v113_: str
                d_2550_v113_ = _dafny.CodePoint('d')
                d_2551_v114_: T3
                nw414_ = C3()
                nw414_.ctor__(self.f24, d_2549_v112_, True, d_2550_v113_, self.f24, (self).f25)
                d_2551_v114_ = nw414_
                d_2552_v115_: _dafny.Map
                d_2552_v115_ = _dafny.Map({d_2551_v114_: d_2548_v111_})
                (globalState).f10 = (d_2548_v111_).issubset((d_2548_v111_) | (((d_2552_v115_)[d_2551_v114_] if (d_2551_v114_) in (d_2552_v115_) else d_2548_v111_)))
                (globalState).f11 = default__.safeDivisionInt(-38, (d_2409_v0_) - (d_2409_v0_))
                d_2553_v116_: C2
                nw415_ = C2()
                nw415_.ctor__(d_2549_v112_, d_2549_v112_)
                d_2553_v116_ = nw415_
                d_2553_v116_ = d_2553_v116_
                (globalState).f11 = d_2409_v0_
            d_2536_cf43_ = (d_2536_cf43_ if (self).f25 else d_2536_cf43_)
            d_2554_v117_: _dafny.Array
            nw416_ = _dafny.Array(_dafny.Map({}), 5)
            d_2554_v117_ = nw416_
            d_2555_v118_: D10
            d_2555_v118_ = D10_DC21(d_2554_v117_)
            d_2556_v119_: _dafny.Map
            d_2556_v119_ = _dafny.Map({(self).f25: d_2555_v118_})
            d_2557_v120_: D10
            d_2557_v120_ = D10_DC24(((d_2556_v119_)[d_2535_cf44_] if (d_2535_cf44_) in (d_2556_v119_) else d_2555_v118_))
            d_2558_v121_: D10
            d_2558_v121_ = D10_DC24((d_2557_v120_ if (self).f25 else d_2557_v120_))
            d_2559_v122_: _dafny.Array
            def lambda123_(d_2560_i12_):
                return (d_2560_i12_) * (len(_dafny.Map({True: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_2561_i13_ in range(default__.abs(88))])})))

            init74_ = lambda123_
            nw417_ = _dafny.Array(None, 26)
            for i0_74_ in range(nw417_.length(0)):
                nw417_[i0_74_] = init74_(i0_74_)
            d_2559_v122_ = nw417_
            d_2562_v123_: _dafny.Seq
            d_2562_v123_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kledxy"))
            d_2563_v124_: _dafny.Map
            d_2563_v124_ = _dafny.Map({d_2409_v0_: (self).f25})
            d_2564_v125_: str
            d_2564_v125_ = _dafny.CodePoint('g')
            index383_ = default__.safeIndex(486, (d_2559_v122_).length(0))
            (d_2559_v122_)[index383_] = len(((d_2562_v123_ if d_2535_cf44_ else d_2562_v123_)).set(default__.safeIndex(len(d_2563_v124_), len((d_2562_v123_ if d_2535_cf44_ else d_2562_v123_))), d_2564_v125_))
            d_2565_v126_: _dafny.MultiSet
            d_2565_v126_ = _dafny.MultiSet([(p0).ispropersubset(p0), d_2535_cf44_, (self).f25])
            d_2566_v127_: _dafny.MultiSet
            d_2566_v127_ = _dafny.MultiSet([(0) - (d_2409_v0_), d_2409_v0_])
            index384_ = default__.safeIndex(486, (d_2559_v122_).length(0))
            rhs401_ = (self).f25
            rhs402_ = d_2558_v121_
            rhs403_ = (d_2409_v0_) - ((d_2566_v127_).cardinality)
            rhs404_ = ((d_2565_v126_).intersection(d_2565_v126_)).intersection(d_2565_v126_)
            lhs358_ = globalState
            lhs359_ = d_2559_v122_
            lhs360_ = default__.safeIndex(486, (d_2559_v122_).length(0))
            lhs358_.f10 = rhs401_
            d_2558_v121_ = rhs402_
            lhs359_[lhs360_] = rhs403_
            d_2565_v126_ = rhs404_
            d_2567_v128_: _dafny.Array
            nw418_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 13)
            d_2567_v128_ = nw418_
            d_2568_v129_: _dafny.Array
            d_2568_v129_ = d_2567_v128_
            source42_ = d_2568_v129_
            d_2569___mcc_h7_ = source42_
            d_2570_cf39_ = d_2569___mcc_h7_
            (self).f24 = not((d_2566_v127_) == (d_2566_v127_))
            d_2571_v130_: C5
            nw419_ = C5()
            nw419_.ctor__(d_2535_cf44_, d_2564_v125_, not(((self).f25 if d_2535_cf44_ else (self).f25)), self.f24)
            d_2571_v130_ = nw419_
            d_2572_v131_: C2
            nw420_ = C2()
            nw420_.ctor__(d_2562_v123_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ylses")))
            d_2572_v131_ = nw420_
            r0 = d_2564_v125_
        elif source41_.is_DC21:
            d_2573___mcc_h5_ = source41_.cf40
            d_2574_cf40_ = d_2573___mcc_h5_
            (globalState).f11 = d_2409_v0_
            (globalState).f22 = (d_2409_v0_) - (len(_dafny.Set({(d_2410_v1_)[default__.safeIndex(d_2409_v0_, len(d_2410_v1_))], d_2409_v0_, d_2409_v0_})))
            d_2575_v132_: _dafny.Seq
            d_2575_v132_ = _dafny.SeqWithoutIsStrInference([d_2446_v30_, d_2446_v30_, d_2446_v30_, d_2446_v30_, d_2446_v30_])
            (globalState).f2 = (d_2575_v132_)[default__.safeIndex(default__.safeModuloInt(d_2409_v0_, d_2409_v0_), len(d_2575_v132_))]
            d_2576_v133_: _dafny.Map
            d_2576_v133_ = _dafny.Map({d_2409_v0_: d_2448_v31_})
            d_2576_v133_ = (d_2576_v133_).set(d_2409_v0_, d_2448_v31_)
        elif True:
            d_2577___mcc_h6_ = source41_.cf45
            d_2578_cf45_ = d_2577___mcc_h6_
            d_2579_v134_: D0
            d_2579_v134_ = D0_DC0(self.f24, d_2409_v0_, d_2409_v0_)
            (globalState).f11 = (self).fm6(d_2579_v134_, globalState)
            index385_ = default__.safeIndex(62, (d_2446_v30_).length(0))
            (d_2446_v30_)[index385_] = self.f24
            d_2580_v135_: _dafny.Seq
            d_2580_v135_ = _dafny.SeqWithoutIsStrInference([self.f24, self.f24])
            d_2581_v136_: _dafny.Map
            d_2581_v136_ = _dafny.Map({d_2580_v135_: True})
            d_2582_v137_: _dafny.Array
            nw421_ = _dafny.Array(int(0), 13)
            d_2582_v137_ = nw421_
            d_2583_v138_: _dafny.Map
            d_2583_v138_ = _dafny.Map({self.f24: (self).f25})
            index386_ = default__.safeIndex(62, (d_2446_v30_).length(0))
            rhs405_ = d_2582_v137_
            rhs406_ = (d_2583_v138_) == (d_2583_v138_)
            rhs407_ = (self).f25
            rhs408_ = (d_2409_v0_) * (d_2409_v0_)
            rhs409_ = d_2581_v136_
            lhs361_ = globalState
            lhs362_ = d_2446_v30_
            lhs363_ = default__.safeIndex(62, (d_2446_v30_).length(0))
            lhs364_ = self
            lhs365_ = globalState
            lhs361_.f7 = rhs405_
            lhs362_[lhs363_] = rhs406_
            lhs364_.f24 = rhs407_
            lhs365_.f16 = rhs408_
            d_2581_v136_ = rhs409_
            d_2584_v139_: _dafny.Array
            nw422_ = _dafny.Array(_dafny.Array(None, 0), 10)
            d_2584_v139_ = nw422_
            d_2585_v140_: D7
            d_2585_v140_ = D7_DC16((self).f25, d_2409_v0_, (d_2446_v30_)[default__.safeIndex(62, (d_2446_v30_).length(0))], self.f24, (d_2446_v30_)[default__.safeIndex(62, (d_2446_v30_).length(0))])
            index387_ = default__.safeIndex(62, (d_2446_v30_).length(0))
            rhs410_ = (d_2446_v30_)[default__.safeIndex(62, (d_2446_v30_).length(0))]
            rhs411_ = ((0) - (d_2409_v0_)) + (((d_2585_v140_).cf31) + (d_2409_v0_))
            rhs412_ = (d_2409_v0_) + (default__.safeModuloInt(d_2409_v0_, d_2409_v0_))
            rhs413_ = (d_2584_v139_ if (self).f25 else d_2584_v139_)
            lhs366_ = d_2446_v30_
            lhs367_ = default__.safeIndex(62, (d_2446_v30_).length(0))
            lhs368_ = globalState
            lhs369_ = globalState
            lhs366_[lhs367_] = rhs410_
            lhs368_.f22 = rhs411_
            lhs369_.f11 = rhs412_
            d_2584_v139_ = rhs413_
            d_2411_v2_ = d_2411_v2_
        d_2586_v141_: _dafny.Seq
        d_2586_v141_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f"))
        d_2587_v142_: str
        d_2587_v142_ = _dafny.CodePoint('b')
        d_2588_v143_: C3
        nw423_ = C3()
        nw423_.ctor__((self).f25, d_2586_v141_, False, d_2587_v142_, (self).f25, (self).f25)
        d_2588_v143_ = nw423_
        d_2589_v144_: _dafny.Seq
        d_2589_v144_ = _dafny.SeqWithoutIsStrInference([d_2588_v143_])
        d_2590_v145_: _dafny.Set
        d_2590_v145_ = _dafny.Set({d_2589_v144_})
        d_2591_v146_: _dafny.Map
        d_2591_v146_ = _dafny.Map({d_2409_v0_: d_2590_v145_})
        hi16_ = default__.safeModuloInt((D15_DC33(d_2409_v0_)).cf55, len(((d_2591_v146_)[d_2409_v0_] if (d_2409_v0_) in (d_2591_v146_) else _dafny.Set({_dafny.SeqWithoutIsStrInference([d_2588_v143_, d_2588_v143_, d_2588_v143_, d_2588_v143_])}))))
        for d_2592_i14_ in range(d_2409_v0_, hi16_):
            (globalState).f22 = (d_2409_v0_) + (555)
            (globalState).f11 = d_2409_v0_
            d_2593_v147_: _dafny.Map
            d_2593_v147_ = _dafny.Map({d_2409_v0_: (self).f25})
            d_2593_v147_ = (d_2593_v147_).set(d_2592_i14_, (d_2592_i14_) == (d_2409_v0_))
            index388_ = default__.safeIndex(260, (d_2446_v30_).length(0))
            (d_2446_v30_)[index388_] = (d_2588_v143_).f32
            index389_ = default__.safeIndex(260, (d_2446_v30_).length(0))
            (d_2446_v30_)[index389_] = (d_2409_v0_) <= (default__.safeModuloInt(d_2409_v0_, d_2592_i14_))
        r0 = default__.fm36(d_2409_v0_, True, (self).f25, globalState)
        d_2594_v148_: _dafny.Map
        d_2594_v148_ = _dafny.Map({d_2409_v0_: d_2409_v0_})
        d_2595_v149_: D2
        d_2595_v149_ = D2_DC3(d_2594_v148_)
        d_2596_v150_: D2
        d_2596_v150_ = D2_DC6(d_2595_v149_)
        d_2597_v151_: D5
        d_2597_v151_ = D5_DC11(d_2596_v150_, not(self.f24))
        pat_let_tv38_ = d_2588_v143_
        pat_let_tv39_ = d_2409_v0_
        pat_let_tv40_ = d_2588_v143_
        pat_let_tv41_ = d_2588_v143_
        def lambda124_(source43_):
            if source43_.is_DC11:
                d_2598___mcc_h8_ = source43_.cf21
                d_2599___mcc_h9_ = source43_.cf22
                d_2600_cf22_ = d_2599___mcc_h9_
                d_2601_cf21_ = d_2598___mcc_h8_
                return ((D3_DC7(_dafny.SeqWithoutIsStrInference([True]))).cf15) + (_dafny.SeqWithoutIsStrInference([d_2600_cf22_]))
            elif True:
                d_2602___mcc_h10_ = source43_.cf20
                d_2603_cf20_ = d_2602___mcc_h10_
                return ((_dafny.SeqWithoutIsStrInference([not((self).f25), self.f24])) + (_dafny.SeqWithoutIsStrInference([(pat_let_tv38_).f32]))).set(default__.safeIndex(pat_let_tv39_, len((_dafny.SeqWithoutIsStrInference([not((self).f25), self.f24])) + (_dafny.SeqWithoutIsStrInference([(pat_let_tv40_).f32])))), (pat_let_tv41_).f32)

        r1 = lambda124_(d_2597_v151_)
        return r0, r1


class C9(T3, T2, T1, T0):
    def  __init__(self):
        self._f24: bool = False
        self._f27: str = _dafny.CodePoint('D')
        self._f25: bool = False
        self._f26: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C9"
    @property
    def f24(self):
        return self._f24
    @f24.setter
    def f24(self, value):
        self._f24 = value
    @property
    def f27(self):
        return self._f27
    @f27.setter
    def f27(self, value):
        self._f27 = value
    @property
    def f25(self):
        return self._f25
    @property
    def f26(self):
        return self._f26
    def ctor__(self, f26, f27, f25, f24):
        (self)._f26 = f26
        (self).f27 = f27
        (self)._f25 = f25
        (self).f24 = f24

    def fm7(self, p0, globalState):
        return self.f24

    def fm6(self, p0, globalState):
        return (0) - ((0) - (default__.safeDivisionInt((-880) + ((D0_DC0((self).f25, 194, 863)).cf2), 675)))

    def fm4(self, p0, p1, globalState):
        return D0_DC0(True, (0) - (-170), (0) - ((len(_dafny.SeqWithoutIsStrInference([-595]))) + (937)))

    def fm5(self, p0, p1, p2, globalState):
        return ((_dafny.MultiSet([False])) - (_dafny.MultiSet([self.f24]))).cardinality

    def fm8(self, p0, globalState):
        return ((_dafny.SeqWithoutIsStrInference([self.f24])) != (_dafny.SeqWithoutIsStrInference([not((self).f26)]))) or (not((self).f25))

    def m2(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: bool = False
        r2: int = int(0)
        r3: int = int(0)
        d_2604_v0_: _dafny.Array
        nw424_ = _dafny.Array(_dafny.Set({}), 4)
        d_2604_v0_ = nw424_
        d_2604_v0_ = d_2604_v0_
        d_2605_v1_: _dafny.Map
        d_2605_v1_ = _dafny.Map({(self).f25: (self).f25})
        d_2606_v2_: _dafny.Seq
        d_2606_v2_ = _dafny.SeqWithoutIsStrInference([(D0_DC0((self).f25, p0, len(d_2605_v1_))).cf0, False, False, (self).f26])
        d_2607_v3_: _dafny.Map
        d_2607_v3_ = _dafny.Map({d_2606_v2_: d_2606_v2_})
        d_2608_v4_: _dafny.Set
        d_2608_v4_ = _dafny.Set({p0, default__.safeDivisionInt(p0, p0), (p0) - (len(((d_2607_v3_)[d_2606_v2_] if (d_2606_v2_) in (d_2607_v3_) else d_2606_v2_)))})
        d_2608_v4_ = _dafny.Set({(857) + (p0), p0})
        d_2609_v5_: _dafny.Array
        def lambda125_(d_2610_p0_):
            def lambda126_(d_2611_i0_):
                return default__.safeModuloInt(d_2611_i0_, d_2610_p0_)

            return lambda126_

        init75_ = lambda125_(p0)
        nw425_ = _dafny.Array(None, 4)
        for i0_75_ in range(nw425_.length(0)):
            nw425_[i0_75_] = init75_(i0_75_)
        d_2609_v5_ = nw425_
        index390_ = default__.safeIndex(241, (d_2609_v5_).length(0))
        (d_2609_v5_)[index390_] = (p0 if (self).f25 else p0)
        d_2612_v6_: _dafny.Map
        d_2612_v6_ = _dafny.Map({(d_2606_v2_) + (d_2606_v2_): p0})
        d_2613_v7_: _dafny.Seq
        d_2613_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ctpyoif"))
        d_2614_v8_: _dafny.Map
        d_2614_v8_ = _dafny.Map({(self).f26: p0})
        d_2615_v9_: _dafny.Map
        d_2615_v9_ = _dafny.Map({len(d_2613_v7_): default__.fm9(d_2614_v8_, p0, self.f27, p0, globalState)})
        index391_ = default__.safeIndex(241, (d_2609_v5_).length(0))
        (d_2609_v5_)[index391_] = (0) - (((d_2612_v6_)[(d_2606_v2_) + (((d_2615_v9_)[p0] if (p0) in (d_2615_v9_) else d_2606_v2_))] if ((d_2606_v2_) + (((d_2615_v9_)[p0] if (p0) in (d_2615_v9_) else d_2606_v2_))) in (d_2612_v6_) else p0))
        d_2616_v10_: _dafny.MultiSet
        d_2616_v10_ = _dafny.MultiSet([self.f24, not((self).f25), (self).f26])
        if (((d_2616_v10_).cardinality) * (p0)) < ((d_2609_v5_)[default__.safeIndex(241, (d_2609_v5_).length(0))]):
            d_2617_v11_: _dafny.MultiSet
            d_2617_v11_ = _dafny.MultiSet([(d_2609_v5_)[default__.safeIndex(241, (d_2609_v5_).length(0))]])
            d_2617_v11_ = d_2617_v11_
            index392_ = default__.safeIndex(241, (d_2609_v5_).length(0))
            (d_2609_v5_)[index392_] = (d_2609_v5_)[default__.safeIndex(241, (d_2609_v5_).length(0))]
            if (self).f26:
                (globalState).f10 = (d_2606_v2_)[default__.safeIndex((d_2609_v5_)[default__.safeIndex(241, (d_2609_v5_).length(0))], len(d_2606_v2_))]
                (globalState).f16 = p0
                d_2605_v1_ = (d_2605_v1_).set(p2, p2)
                (globalState).f11 = (0) - (((d_2614_v8_)[(self).f26] if ((self).f26) in (d_2614_v8_) else (p0) - (p0)))
                d_2618_v12_: _dafny.Array
                nw426_ = _dafny.Array(_dafny.Seq({}), 25)
                d_2618_v12_ = nw426_
                index393_ = default__.safeIndex(197, (d_2618_v12_).length(0))
                (d_2618_v12_)[index393_] = d_2606_v2_
                index394_ = default__.safeIndex(197, (d_2618_v12_).length(0))
                (d_2618_v12_)[index394_] = d_2606_v2_
            elif True:
                rhs414_ = default__.safeDivisionInt((d_2609_v5_)[default__.safeIndex(241, (d_2609_v5_).length(0))], (p0) * (len(_dafny.SeqWithoutIsStrInference([self.f27 for d_2619_i1_ in range(default__.abs(-777))]))))
                rhs415_ = (len((d_2613_v7_ if not(True) else d_2613_v7_))) != (p0)
                lhs370_ = globalState
                lhs371_ = self
                lhs370_.f22 = rhs414_
                lhs371_.f24 = rhs415_
                d_2620_v13_: _dafny.Map
                d_2620_v13_ = _dafny.Map({p0: p0})
                r0 = default__.fm1(self.f24, (p0) < (((d_2617_v11_)[p0] if (p0) in (d_2617_v11_) else ((d_2614_v8_)[(self).f25] if ((self).f25) in (d_2614_v8_) else (d_2609_v5_)[default__.safeIndex(241, (d_2609_v5_).length(0))]))), len(d_2620_v13_), globalState)
                r0 = (self).fm8(p2, globalState)
                index395_ = default__.safeIndex(241, (d_2609_v5_).length(0))
                (d_2609_v5_)[index395_] = default__.safeDivisionInt(p0, p0)
                d_2621_v14_: _dafny.Set
                d_2621_v14_ = _dafny.Set({not(self.f24), (self).f26, p2, default__.fm1((self).f26, self.f24, p0, globalState), (self).f26})
                d_2622_v15_: _dafny.Map
                d_2622_v15_ = _dafny.Map({(self).f26: d_2620_v13_})
                d_2623_v16_: _dafny.Seq
                d_2623_v16_ = _dafny.SeqWithoutIsStrInference([p0])
                index396_ = default__.safeIndex(241, (d_2609_v5_).length(0))
                rhs416_ = (d_2609_v5_)[default__.safeIndex(241, (d_2609_v5_).length(0))]
                rhs417_ = (d_2606_v2_)[default__.safeIndex(default__.safeDivisionInt(len(d_2621_v14_), (0) - ((d_2609_v5_)[default__.safeIndex(241, (d_2609_v5_).length(0))])), len(d_2606_v2_))]
                rhs418_ = d_2609_v5_
                rhs419_ = (len(((d_2622_v15_)[self.f24] if (self.f24) in (d_2622_v15_) else d_2620_v13_))) > ((d_2623_v16_)[default__.safeIndex((d_2609_v5_)[default__.safeIndex(241, (d_2609_v5_).length(0))], len(d_2623_v16_))])
                lhs372_ = d_2609_v5_
                lhs373_ = default__.safeIndex(241, (d_2609_v5_).length(0))
                lhs374_ = globalState
                lhs375_ = globalState
                lhs372_[lhs373_] = rhs416_
                r0 = rhs417_
                lhs374_.f7 = rhs418_
                lhs375_.f10 = rhs419_
            rhs420_ = (self).f25
            rhs421_ = (self).f26
            rhs422_ = ((d_2605_v1_)[(self).f25] if ((self).f25) in (d_2605_v1_) else (self).f26)
            rhs423_ = ((d_2606_v2_) + ((d_2606_v2_).set(default__.safeIndex(-962, len(d_2606_v2_)), (self).f25))) + (d_2606_v2_)
            rhs424_ = _dafny.MultiSet([-192, len(d_2613_v7_)])
            lhs376_ = self
            r1 = rhs420_
            r1 = rhs421_
            lhs376_.f24 = rhs422_
            d_2606_v2_ = rhs423_
            d_2617_v11_ = rhs424_
            source44_ = p1
            d_2624___mcc_h0_ = source44_.cf0
            d_2625___mcc_h1_ = source44_.cf1
            d_2626___mcc_h2_ = source44_.cf2
            d_2627_cf2_ = d_2626___mcc_h2_
            d_2628_cf1_ = d_2625___mcc_h1_
            d_2629_cf0_ = d_2624___mcc_h0_
            d_2630_v17_: _dafny.Map
            d_2630_v17_ = _dafny.Map({d_2627_cf2_: self.f24})
            (globalState).f10 = (not((self).f25) if (self).f25 else ((d_2630_v17_)[p0] if (p0) in (d_2630_v17_) else d_2629_cf0_))
            d_2631_v18_: _dafny.Seq
            d_2631_v18_ = _dafny.SeqWithoutIsStrInference([(d_2609_v5_)[default__.safeIndex(241, (d_2609_v5_).length(0))], d_2627_cf2_])
            d_2632_v19_: D0
            d_2632_v19_ = D0_DC0(d_2629_cf0_, ((d_2609_v5_)[default__.safeIndex(241, (d_2609_v5_).length(0))]) + ((0) - (d_2627_cf2_)), len(d_2631_v18_))
            d_2632_v19_ = p1
            r0 = ((d_2606_v2_)[default__.safeIndex(p0, len(d_2606_v2_))] if d_2629_cf0_ else (d_2606_v2_)[default__.safeIndex(p0, len(d_2606_v2_))])
            d_2627_cf2_ = (0) - (d_2627_cf2_)
        elif True:
            r0 = p2
            d_2633_v20_: _dafny.Map
            d_2633_v20_ = _dafny.Map({self.f27: d_2613_v7_})
            r2 = len(((d_2633_v20_) | (_dafny.Map({self.f27: d_2613_v7_}))) | (d_2633_v20_))
            d_2634_v21_: _dafny.Map
            d_2634_v21_ = _dafny.Map({217: d_2613_v7_})
            d_2635_v22_: _dafny.Set
            d_2635_v22_ = _dafny.Set({(D1_DC1(((d_2634_v21_)[724] if (724) in (d_2634_v21_) else d_2613_v7_))).cf3, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xpldpnrw")), d_2613_v7_, d_2613_v7_})
            d_2636_v23_: _dafny.Array
            nw427_ = _dafny.Array(_dafny.Array(None, 0), 9)
            d_2636_v23_ = nw427_
            d_2637_v24_: _dafny.Map
            d_2637_v24_ = _dafny.Map({(d_2609_v5_)[default__.safeIndex(241, (d_2609_v5_).length(0))]: p0})
            d_2638_v25_: D2
            d_2638_v25_ = D2_DC3(d_2637_v24_)
            (self).m4(d_2635_v22_, (self).fm8(self.f24, globalState), d_2636_v23_, len((d_2638_v25_).cf8), globalState)
            (self).m4(d_2635_v22_, self.f24, d_2636_v23_, p0, globalState)
            d_2639_v26_: _dafny.Array
            nw428_ = _dafny.Array(False, 5)
            d_2639_v26_ = nw428_
            index397_ = default__.safeIndex(78, (d_2639_v26_).length(0))
            (d_2639_v26_)[index397_] = (self).f26
            index398_ = default__.safeIndex(78, (d_2639_v26_).length(0))
            (d_2639_v26_)[index398_] = ((self).f25 if (self).f25 else (d_2613_v7_) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fvsq"))))
        (globalState).f22 = (0) - (default__.safeDivisionInt(p0, p0))
        d_2640_v27_: D2
        d_2640_v27_ = D2_DC5(self.f27, (self).f26, p0, p0)
        (globalState).f10 = (d_2640_v27_).cf11
        r0 = (self).f25
        d_2641_v28_: D2
        d_2641_v28_ = D2_DC4(True)
        pat_let_tv42_ = p2
        pat_let_tv43_ = p2
        def lambda127_(source45_):
            if source45_.is_DC4:
                d_2642___mcc_h3_ = source45_.cf9
                d_2643_cf9_ = d_2642___mcc_h3_
                return (self).f26
            elif source45_.is_DC5:
                d_2644___mcc_h4_ = source45_.cf10
                d_2645___mcc_h5_ = source45_.cf11
                d_2646___mcc_h6_ = source45_.cf12
                d_2647___mcc_h7_ = source45_.cf13
                d_2648_cf13_ = d_2647___mcc_h7_
                d_2649_cf12_ = d_2646___mcc_h6_
                d_2650_cf11_ = d_2645___mcc_h5_
                d_2651_cf10_ = d_2644___mcc_h4_
                return self.f24
            elif source45_.is_DC3:
                d_2652___mcc_h8_ = source45_.cf8
                d_2653_cf8_ = d_2652___mcc_h8_
                return pat_let_tv42_
            elif True:
                d_2654___mcc_h9_ = source45_.cf14
                d_2655_cf14_ = d_2654___mcc_h9_
                return (self).f25

        def iife183_(_pat_let49_0):
            def iife184_(d_2656_dt__update__tmp_h0_):
                def iife185_(_pat_let50_0):
                    def iife186_(d_2657_dt__update_hcf9_h0_):
                        return D2_DC4(d_2657_dt__update_hcf9_h0_)
                    return iife186_(_pat_let50_0)
                return iife185_(pat_let_tv43_)
            return iife184_(_pat_let49_0)
        r1 = lambda127_(iife183_(d_2641_v28_))
        d_2658_v29_: _dafny.Map
        d_2658_v29_ = _dafny.Map({p0: p0})
        d_2659_v30_: _dafny.Map
        d_2659_v30_ = _dafny.Map({p0: (self).f25})
        r2 = ((d_2658_v29_)[default__.fm0((d_2609_v5_)[default__.safeIndex(241, (d_2609_v5_).length(0))], True, d_2659_v30_, (d_2609_v5_)[default__.safeIndex(241, (d_2609_v5_).length(0))], globalState)] if (default__.fm0((d_2609_v5_)[default__.safeIndex(241, (d_2609_v5_).length(0))], True, d_2659_v30_, (d_2609_v5_)[default__.safeIndex(241, (d_2609_v5_).length(0))], globalState)) in (d_2658_v29_) else (d_2609_v5_)[default__.safeIndex(241, (d_2609_v5_).length(0))])
        r3 = (d_2609_v5_)[default__.safeIndex(241, (d_2609_v5_).length(0))]
        return r0, r1, r2, r3

    def m3(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: bool = False
        r2: _dafny.Map = _dafny.Map({})
        (globalState).f11 = (default__.safeDivisionInt(p0, p0)) + (p0)
        if self.f24:
            d_2660_v0_: _dafny.Seq
            d_2660_v0_ = _dafny.SeqWithoutIsStrInference([p0, p0, p0])
            (globalState).f10 = (p0) > (len((d_2660_v0_) + (default__.fm10(p0, globalState))))
            (globalState).f11 = p0
            rhs425_ = p0
            rhs426_ = self.f27
            rhs427_ = p0
            lhs377_ = globalState
            lhs378_ = self
            lhs379_ = globalState
            lhs377_.f22 = rhs425_
            lhs378_.f27 = rhs426_
            lhs379_.f22 = rhs427_
            d_2661_v1_: _dafny.Array
            def lambda128_(d_2662_i0_):
                return default__.safeModuloInt(d_2662_i0_, len(_dafny.SeqWithoutIsStrInference([self.f24, not((self).f26)])))

            init76_ = lambda128_
            nw429_ = _dafny.Array(None, 5)
            for i0_76_ in range(nw429_.length(0)):
                nw429_[i0_76_] = init76_(i0_76_)
            d_2661_v1_ = nw429_
            d_2663_v2_: _dafny.Map
            d_2663_v2_ = _dafny.Map({p0: d_2661_v1_})
            d_2664_v3_: D1
            d_2664_v3_ = D1_DC2(d_2661_v1_, self.f27, p0, (self).f25)
            d_2665_v4_: _dafny.Array
            nw430_ = _dafny.Array(None, 13)
            nw430_[int(0)] = d_2661_v1_
            nw430_[int(1)] = d_2661_v1_
            nw430_[int(2)] = d_2661_v1_
            nw430_[int(3)] = ((d_2663_v2_)[p0] if (p0) in (d_2663_v2_) else (d_2664_v3_).cf4)
            nw430_[int(4)] = d_2661_v1_
            nw430_[int(5)] = d_2661_v1_
            nw430_[int(6)] = d_2661_v1_
            nw430_[int(7)] = d_2661_v1_
            nw430_[int(8)] = d_2661_v1_
            nw430_[int(9)] = d_2661_v1_
            nw430_[int(10)] = d_2661_v1_
            nw430_[int(11)] = d_2661_v1_
            nw430_[int(12)] = (d_2661_v1_ if p1 else d_2661_v1_)
            d_2665_v4_ = nw430_
            rhs428_ = (d_2665_v4_ if (self).f26 else d_2665_v4_)
            rhs429_ = p0
            lhs380_ = globalState
            d_2665_v4_ = rhs428_
            lhs380_.f16 = rhs429_
            d_2666_v5_: _dafny.Set
            d_2666_v5_ = _dafny.Set({(self).f25})
            d_2667_v6_: D2
            d_2667_v6_ = D2_DC5(self.f27, p1, p0, len(d_2666_v5_))
            d_2668_v7_: _dafny.MultiSet
            d_2668_v7_ = _dafny.MultiSet([d_2667_v6_])
            pat_let_tv44_ = p0
            d_2669_v8_: _dafny.Seq
            def iife187_(_pat_let51_0):
                def iife188_(d_2670_dt__update__tmp_h0_):
                    def iife189_(_pat_let52_0):
                        def iife190_(d_2671_dt__update_hcf6_h0_):
                            return D1_DC2((d_2670_dt__update__tmp_h0_).cf4, (d_2670_dt__update__tmp_h0_).cf5, d_2671_dt__update_hcf6_h0_, (d_2670_dt__update__tmp_h0_).cf7)
                        return iife190_(_pat_let52_0)
                    return iife189_(pat_let_tv44_)
                return iife188_(_pat_let51_0)
            d_2669_v8_ = _dafny.SeqWithoutIsStrInference([iife187_(d_2664_v3_)])
            d_2668_v7_ = (d_2668_v7_).set(D2_DC5(self.f27, (self).f26, p0, p0), default__.abs((len(d_2669_v8_) if (self).f26 else p0)))
        elif True:
            d_2672_v9_: T2
            nw431_ = C4()
            nw431_.ctor__((self).f25, self.f24)
            d_2672_v9_ = nw431_
            d_2673_v10_: _dafny.Map
            d_2673_v10_ = _dafny.Map({default__.fm11(globalState): _dafny.Map({d_2672_v9_: _dafny.Map({p0: (d_2672_v9_).f25})})})
            d_2674_v11_: _dafny.Map
            d_2674_v11_ = _dafny.Map({56: (d_2672_v9_).f25})
            d_2675_v12_: _dafny.Map
            d_2675_v12_ = _dafny.Map({len(d_2674_v11_): p1})
            d_2673_v10_ = (d_2673_v10_).set(self.f27, _dafny.Map({d_2672_v9_: d_2674_v11_}))
            d_2676_v13_: _dafny.Seq
            d_2676_v13_ = _dafny.SeqWithoutIsStrInference([(self).f25, (self).f26])
            d_2677_v14_: _dafny.Seq
            d_2677_v14_ = _dafny.SeqWithoutIsStrInference([(d_2676_v13_)[default__.safeIndex(p0, len(d_2676_v13_))], (self).f26, (self.f24 if (self).f25 else True), (p1) == ((d_2672_v9_).f25), (d_2672_v9_).f25])
            if (d_2677_v14_)[default__.safeIndex(p0, len(d_2677_v14_))]:
                index399_ = default__.safeIndex(972, (p2).length(0))
                (p2)[index399_] = d_2672_v9_.f24
                d_2678_v15_: _dafny.Seq
                d_2678_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a"))
                d_2679_v16_: D0
                d_2679_v16_ = D0_DC0((self).f25, p0, (0) - ((0) - (len(d_2678_v15_))))
                index400_ = default__.safeIndex(972, (p2).length(0))
                (p2)[index400_] = (p0) > ((self).fm6(d_2679_v16_, globalState))
                (d_2672_v9_).f24 = p1
                d_2680_v17_: _dafny.Array
                nw432_ = _dafny.Array(None, 7)
                d_2680_v17_ = nw432_
                d_2681_v18_: C4
                nw433_ = C4()
                nw433_.ctor__(default__.fm1((self).f26, d_2672_v9_.f24, p0, globalState), (p2)[default__.safeIndex(972, (p2).length(0))])
                d_2681_v18_ = nw433_
                index401_ = default__.safeIndex(5, (d_2680_v17_).length(0))
                (d_2680_v17_)[index401_] = d_2681_v18_
                d_2682_v19_: _dafny.MultiSet
                d_2682_v19_ = _dafny.MultiSet([(d_2672_v9_).f25])
                d_2683_v20_: _dafny.Set
                d_2683_v20_ = _dafny.Set({(d_2682_v19_).cardinality})
                d_2684_v21_: _dafny.MultiSet
                d_2684_v21_ = _dafny.MultiSet([559, len(d_2683_v20_)])
                d_2685_v22_: _dafny.MultiSet
                d_2685_v22_ = d_2684_v21_
                d_2686_v23_: _dafny.Array
                nw434_ = _dafny.Array(int(0), 13)
                d_2686_v23_ = nw434_
                index402_ = default__.safeIndex(593, (d_2686_v23_).length(0))
                (d_2686_v23_)[index402_] = default__.safeDivisionInt(p0, 35)
                d_2687_v24_: _dafny.Map
                d_2687_v24_ = _dafny.Map({(d_2672_v9_).f25: self.f27})
                d_2688_v25_: _dafny.Seq
                d_2688_v25_ = _dafny.SeqWithoutIsStrInference([d_2686_v23_])
                d_2689_v26_: _dafny.Seq
                d_2689_v26_ = _dafny.SeqWithoutIsStrInference([(d_2688_v25_)[default__.safeIndex(p0, len(d_2688_v25_))]])
                index403_ = default__.safeIndex(5, (d_2680_v17_).length(0))
                index404_ = default__.safeIndex(593, (d_2686_v23_).length(0))
                rhs430_ = ((d_2687_v24_)[(True if (d_2672_v9_).f25 else d_2672_v9_.f24)] if ((True if (d_2672_v9_).f25 else d_2672_v9_.f24)) in (d_2687_v24_) else self.f27)
                rhs431_ = d_2681_v18_
                rhs432_ = (d_2688_v25_)[default__.safeIndex(default__.safeModuloInt(p0, 299), len(d_2688_v25_))]
                rhs433_ = d_2685_v22_
                rhs434_ = p0
                lhs381_ = self
                lhs382_ = d_2680_v17_
                lhs383_ = default__.safeIndex(5, (d_2680_v17_).length(0))
                lhs384_ = globalState
                lhs385_ = d_2686_v23_
                lhs386_ = default__.safeIndex(593, (d_2686_v23_).length(0))
                lhs381_.f27 = rhs430_
                lhs382_[lhs383_] = rhs431_
                lhs384_.f7 = rhs432_
                d_2685_v22_ = rhs433_
                lhs385_[lhs386_] = rhs434_
                (globalState).f11 = (d_2686_v23_)[default__.safeIndex(593, (d_2686_v23_).length(0))]
                d_2690_v27_: D10
                d_2690_v27_ = D10_DC22(p0, self.f24)
                rhs435_ = (d_2672_v9_).f25
                rhs436_ = (D1_DC2(d_2686_v23_, _dafny.CodePoint('f'), (d_2690_v27_).cf41, d_2672_v9_.f24)).cf5
                lhs387_ = self
                r0 = rhs435_
                lhs387_.f27 = rhs436_
            elif True:
                d_2691_v28_: _dafny.Array
                def lambda129_(d_2692_i1_):
                    return default__.safeDivisionInt(d_2692_i1_, 212)

                init77_ = lambda129_
                nw435_ = _dafny.Array(None, 18)
                for i0_77_ in range(nw435_.length(0)):
                    nw435_[i0_77_] = init77_(i0_77_)
                d_2691_v28_ = nw435_
                index405_ = default__.safeIndex(952, (d_2691_v28_).length(0))
                (d_2691_v28_)[index405_] = default__.safeDivisionInt(871, p0)
                index406_ = default__.safeIndex(952, (d_2691_v28_).length(0))
                (d_2691_v28_)[index406_] = p0
                (globalState).f10 = d_2672_v9_.f24
                d_2693_v29_: _dafny.Seq
                d_2693_v29_ = _dafny.SeqWithoutIsStrInference([self.f27, self.f27, self.f27, self.f27, self.f27])
                d_2694_v30_: _dafny.Set
                d_2694_v30_ = _dafny.Set({_dafny.MultiSet(d_2693_v29_)})
                d_2695_v31_: T3
                nw436_ = C6()
                nw436_.ctor__((d_2672_v9_).f25, (self).f25, d_2672_v9_.f24, self.f27, d_2672_v9_.f24)
                d_2695_v31_ = nw436_
                d_2696_v32_: _dafny.Map
                d_2696_v32_ = _dafny.Map({d_2695_v31_: p0})
                d_2697_v33_: _dafny.Map
                d_2697_v33_ = _dafny.Map({d_2672_v9_.f24: d_2695_v31_})
                d_2698_v34_: _dafny.Map
                d_2698_v34_ = _dafny.Map({p2: d_2695_v31_.f24})
                d_2699_v35_: D18
                d_2699_v35_ = D18_DC37(d_2696_v32_)
                d_2700_v36_: _dafny.Seq
                d_2700_v36_ = _dafny.SeqWithoutIsStrInference([d_2696_v32_])
                d_2701_v37_: D8
                d_2701_v37_ = D8_DC19(p1, (d_2691_v28_)[default__.safeIndex(952, (d_2691_v28_).length(0))])
                d_2702_v38_: _dafny.Map
                d_2702_v38_ = _dafny.Map({(self).f26: d_2696_v32_})
                d_2703_v39_: _dafny.Array
                nw437_ = _dafny.Array(None, 29)
                nw437_[int(0)] = d_2696_v32_
                nw437_[int(1)] = d_2696_v32_
                nw437_[int(2)] = (d_2696_v32_) | (d_2696_v32_)
                nw437_[int(3)] = d_2696_v32_
                nw437_[int(4)] = (d_2696_v32_ if d_2672_v9_.f24 else d_2696_v32_)
                nw437_[int(5)] = (d_2696_v32_).set((d_2695_v31_), (d_2691_v28_)[default__.safeIndex(952, (d_2691_v28_).length(0))])
                nw437_[int(6)] = d_2696_v32_
                nw437_[int(7)] = d_2696_v32_
                nw437_[int(8)] = (d_2696_v32_) | (d_2696_v32_)
                nw437_[int(9)] = d_2696_v32_
                nw437_[int(10)] = d_2696_v32_
                nw437_[int(11)] = d_2696_v32_
                nw437_[int(12)] = (_dafny.Map({d_2695_v31_: (d_2691_v28_)[default__.safeIndex(952, (d_2691_v28_).length(0))]})) | (d_2696_v32_)
                nw437_[int(13)] = _dafny.Map({((d_2697_v33_)[False] if (False) in (d_2697_v33_) else d_2695_v31_): (0) - (len(d_2698_v34_))})
                nw437_[int(14)] = d_2696_v32_
                nw437_[int(15)] = d_2696_v32_
                nw437_[int(16)] = d_2696_v32_
                nw437_[int(17)] = (_dafny.Map({d_2695_v31_: len(d_2693_v29_)})) | (d_2696_v32_)
                nw437_[int(18)] = (d_2696_v32_) | (d_2696_v32_)
                nw437_[int(19)] = d_2696_v32_
                nw437_[int(20)] = (d_2699_v35_).cf58
                nw437_[int(21)] = d_2696_v32_
                nw437_[int(22)] = (d_2700_v36_)[default__.safeIndex((d_2701_v37_).cf38, len(d_2700_v36_))]
                nw437_[int(23)] = (d_2696_v32_) | (d_2696_v32_)
                nw437_[int(24)] = d_2696_v32_
                nw437_[int(25)] = d_2696_v32_
                nw437_[int(26)] = d_2696_v32_
                nw437_[int(27)] = _dafny.Map({d_2695_v31_: -377})
                nw437_[int(28)] = ((d_2702_v38_)[False] if (False) in (d_2702_v38_) else d_2696_v32_)
                d_2703_v39_ = nw437_
                index407_ = default__.safeIndex(773, (d_2703_v39_).length(0))
                (d_2703_v39_)[index407_] = d_2696_v32_
                d_2704_v40_: D0
                d_2704_v40_ = D0_DC0(d_2695_v31_.f24, p0, p0)
                index408_ = default__.safeIndex(952, (d_2691_v28_).length(0))
                index409_ = default__.safeIndex(773, (d_2703_v39_).length(0))
                rhs437_ = (d_2695_v31_).fm6(d_2704_v40_, globalState)
                rhs438_ = d_2694_v30_
                rhs439_ = (d_2696_v32_) | (d_2696_v32_)
                lhs388_ = d_2691_v28_
                lhs389_ = default__.safeIndex(952, (d_2691_v28_).length(0))
                lhs390_ = d_2703_v39_
                lhs391_ = default__.safeIndex(773, (d_2703_v39_).length(0))
                lhs388_[lhs389_] = rhs437_
                d_2694_v30_ = rhs438_
                lhs390_[lhs391_] = rhs439_
                d_2705_v41_: C3
                nw438_ = C3()
                nw438_.ctor__((d_2676_v13_)[default__.safeIndex(len(_dafny.Map({(d_2695_v31_).f26: d_2672_v9_.f24})), len(d_2676_v13_))], d_2693_v29_, (d_2677_v14_)[default__.safeIndex(50, len(d_2677_v14_))], self.f27, (d_2695_v31_).f25, d_2695_v31_.f24)
                d_2705_v41_ = nw438_
                d_2706_v42_: D11
                d_2706_v42_ = D11_DC25(d_2705_v41_)
                d_2707_v43_: _dafny.Map
                d_2707_v43_ = _dafny.Map({d_2706_v42_: not((d_2695_v31_).f25)})
                d_2707_v43_ = (d_2707_v43_).set(d_2706_v42_, (d_2695_v31_).f25)
                d_2708_v44_: _dafny.Seq
                d_2708_v44_ = _dafny.SeqWithoutIsStrInference([d_2695_v31_])
                d_2709_v45_: T3
                d_2709_v45_ = (d_2708_v44_)[default__.safeIndex((0) - ((d_2691_v28_)[default__.safeIndex(952, (d_2691_v28_).length(0))]), len(d_2708_v44_))]
                d_2709_v45_ = d_2695_v31_
            if self.f24:
                (self).f24 = p1
                d_2710_v46_: _dafny.MultiSet
                d_2710_v46_ = _dafny.MultiSet([p0, p0])
                d_2711_v47_: _dafny.Seq
                d_2711_v47_ = _dafny.SeqWithoutIsStrInference([d_2710_v46_])
                d_2712_v48_: _dafny.Set
                d_2712_v48_ = _dafny.Set({(d_2711_v47_)[default__.safeIndex(p0, len(d_2711_v47_))]})
                d_2713_v50_: _dafny.Seq
                d_2713_v50_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nrl"))
                d_2714_v51_: _dafny.Seq
                d_2714_v51_ = _dafny.SeqWithoutIsStrInference([d_2713_v50_])
                d_2715_v52_: C6
                nw439_ = C6()
                def iife191_():
                    coll85_ = _dafny.Set()
                    compr_85_: _dafny.MultiSet
                    for compr_85_ in (d_2711_v47_).Elements:
                        d_2716_v49_: _dafny.MultiSet = compr_85_
                        if (d_2716_v49_) in (d_2711_v47_):
                            coll85_ = coll85_.union(_dafny.Set([d_2716_v49_]))
                    return _dafny.Set(coll85_)
                nw439_.ctor__(not((d_2712_v48_).isdisjoint(iife191_()
                )), (self.f24) and (True), (len(d_2714_v51_)) < (p0), default__.fm11(globalState), ((d_2675_v12_)[(default__.fm27(not(False), globalState)).cardinality] if ((default__.fm27(not(False), globalState)).cardinality) in (d_2675_v12_) else (self).f25))
                d_2715_v52_ = nw439_
                (globalState).f10 = (self).f26
                d_2717_v53_: _dafny.Array
                def lambda130_(d_2718_p0_):
                    def lambda131_(d_2719_i2_):
                        return D15_DC33(d_2718_p0_)

                    return lambda131_

                init78_ = lambda130_(p0)
                nw440_ = _dafny.Array(None, 19)
                for i0_78_ in range(nw440_.length(0)):
                    nw440_[i0_78_] = init78_(i0_78_)
                d_2717_v53_ = nw440_
                d_2720_v54_: D15
                d_2720_v54_ = D15_DC33(p0)
                index410_ = default__.safeIndex(114, (d_2717_v53_).length(0))
                (d_2717_v53_)[index410_] = d_2720_v54_
                index411_ = default__.safeIndex(114, (d_2717_v53_).length(0))
                (d_2717_v53_)[index411_] = d_2720_v54_
                (globalState).f10 = self.f24
            elif True:
                (globalState).f16 = default__.safeModuloInt(p0, p0)
                (globalState).f10 = self.f24
                index412_ = default__.safeIndex(343, (p2).length(0))
                (p2)[index412_] = self.f24
                index413_ = default__.safeIndex(343, (p2).length(0))
                (p2)[index413_] = (self).f26
                d_2721_v55_: _dafny.Seq
                d_2721_v55_ = _dafny.SeqWithoutIsStrInference([p0])
                d_2722_v56_: _dafny.Array
                nw441_ = _dafny.Array(None, 23)
                nw441_[int(0)] = d_2672_v9_.f24
                nw441_[int(1)] = (d_2672_v9_).f25
                nw441_[int(2)] = (self).f25
                nw441_[int(3)] = d_2672_v9_.f24
                nw441_[int(4)] = (d_2672_v9_).f25
                nw441_[int(5)] = (d_2672_v9_).f25
                nw441_[int(6)] = d_2672_v9_.f24
                nw441_[int(7)] = False
                nw441_[int(8)] = self.f24
                nw441_[int(9)] = (self).f26
                nw441_[int(10)] = (d_2672_v9_).f25
                nw441_[int(11)] = self.f24
                nw441_[int(12)] = True
                nw441_[int(13)] = (self).f26
                nw441_[int(14)] = d_2672_v9_.f24
                nw441_[int(15)] = True
                nw441_[int(16)] = default__.fm1((p2)[default__.safeIndex(343, (p2).length(0))], not(p1), p0, globalState)
                nw441_[int(17)] = (d_2676_v13_)[default__.safeIndex((0) - ((d_2721_v55_)[default__.safeIndex(p0, len(d_2721_v55_))]), len(d_2676_v13_))]
                nw441_[int(18)] = (self).f25
                nw441_[int(19)] = p1
                nw441_[int(20)] = self.f24
                nw441_[int(21)] = d_2672_v9_.f24
                nw441_[int(22)] = (d_2672_v9_).f25
                d_2722_v56_ = nw441_
                d_2723_v57_: _dafny.Map
                d_2723_v57_ = _dafny.Map({d_2722_v56_: (d_2672_v9_).f25})
                d_2724_v58_: C4
                nw442_ = C4()
                nw442_.ctor__((d_2672_v9_).f25, self.f24)
                d_2724_v58_ = nw442_
                d_2725_v59_: _dafny.Seq
                d_2725_v59_ = _dafny.SeqWithoutIsStrInference([d_2724_v58_, d_2724_v58_])
                d_2726_v60_: D7
                d_2726_v60_ = D7_DC16(d_2672_v9_.f24, len(d_2725_v59_), (self).f25, (self).f25, (self).f26)
                d_2727_v61_: D0
                d_2727_v61_ = D0_DC0((d_2726_v60_).cf33, p0, p0)
                d_2728_v62_: _dafny.Set
                d_2728_v62_ = _dafny.Set({p0})
                d_2729_v63_: _dafny.Array
                nw443_ = _dafny.Array(None, 25)
                nw443_[int(0)] = d_2672_v9_
                nw443_[int(1)] = d_2672_v9_
                nw443_[int(2)] = d_2672_v9_
                nw443_[int(3)] = d_2672_v9_
                nw443_[int(4)] = d_2672_v9_
                nw443_[int(5)] = d_2672_v9_
                nw443_[int(6)] = d_2672_v9_
                nw443_[int(7)] = d_2672_v9_
                nw443_[int(8)] = d_2672_v9_
                nw443_[int(9)] = d_2672_v9_
                nw443_[int(10)] = d_2672_v9_
                nw443_[int(11)] = d_2672_v9_
                nw443_[int(12)] = d_2672_v9_
                nw443_[int(13)] = d_2672_v9_
                nw443_[int(14)] = d_2672_v9_
                nw443_[int(15)] = d_2672_v9_
                nw443_[int(16)] = d_2672_v9_
                nw443_[int(17)] = d_2672_v9_
                nw443_[int(18)] = d_2672_v9_
                nw443_[int(19)] = d_2672_v9_
                nw443_[int(20)] = d_2672_v9_
                nw443_[int(21)] = d_2672_v9_
                nw443_[int(22)] = d_2672_v9_
                nw443_[int(23)] = d_2672_v9_
                nw443_[int(24)] = d_2672_v9_
                d_2729_v63_ = nw443_
                d_2730_v64_: _dafny.Set
                d_2730_v64_ = _dafny.Set({d_2729_v63_})
                index414_ = default__.safeIndex(343, (p2).length(0))
                index415_ = default__.safeIndex(343, (p2).length(0))
                rhs440_ = not (not(((d_2723_v57_)[d_2722_v56_] if (d_2722_v56_) in (d_2723_v57_) else (self).f26))) or ((p2)[default__.safeIndex(343, (p2).length(0))])
                rhs441_ = p0
                rhs442_ = default__.safeModuloInt(((self).fm6(d_2727_v61_, globalState) if (d_2672_v9_).f25 else len(_dafny.SeqWithoutIsStrInference([p1]))), p0)
                rhs443_ = (self).f25
                rhs444_ = not((_dafny.Set({p0, len(d_2730_v64_), 58, p0})).issubset(d_2728_v62_))
                lhs392_ = p2
                lhs393_ = default__.safeIndex(343, (p2).length(0))
                lhs394_ = globalState
                lhs395_ = globalState
                lhs396_ = globalState
                lhs397_ = p2
                lhs398_ = default__.safeIndex(343, (p2).length(0))
                lhs392_[lhs393_] = rhs440_
                lhs394_.f11 = rhs441_
                lhs395_.f16 = rhs442_
                lhs396_.f10 = rhs443_
                lhs397_[lhs398_] = rhs444_
                d_2731_v65_: C8
                nw444_ = C8()
                nw444_.ctor__((self).f25, (_dafny.SeqWithoutIsStrInference([p0])) < (d_2721_v55_))
                d_2731_v65_ = nw444_
            if (d_2672_v9_).f25:
                (globalState).f22 = (p0) * (p0)
                d_2732_v66_: _dafny.Array
                nw445_ = _dafny.Array(_dafny.Array(None, 0), 28)
                d_2732_v66_ = nw445_
                d_2733_v67_: _dafny.Set
                d_2733_v67_ = _dafny.Set({(0) - (p0)})
                d_2734_v68_: _dafny.Array
                nw446_ = _dafny.Array(None, 3)
                nw446_[int(0)] = d_2733_v67_
                nw446_[int(1)] = d_2733_v67_
                nw446_[int(2)] = d_2733_v67_
                d_2734_v68_ = nw446_
                index416_ = default__.safeIndex(628, (d_2732_v66_).length(0))
                (d_2732_v66_)[index416_] = d_2734_v68_
                index417_ = default__.safeIndex(628, (d_2732_v66_).length(0))
                (d_2732_v66_)[index417_] = d_2734_v68_
                d_2735_v69_: D0
                d_2735_v69_ = D0_DC0(True, p0, p0)
                (globalState).f11 = default__.safeDivisionInt((p0) - ((d_2672_v9_).fm6(d_2735_v69_, globalState)), p0)
                d_2736_v70_: D8
                d_2736_v70_ = D8_DC19(d_2672_v9_.f24, p0)
                (globalState).f16 = default__.safeDivisionInt((p0) * (p0), (d_2736_v70_).cf38)
                d_2737_v71_: D3
                d_2737_v71_ = D3_DC8((d_2672_v9_).f25, p0, p0)
                d_2738_v72_: _dafny.Map
                d_2738_v72_ = _dafny.Map({not((d_2672_v9_).f25): d_2737_v71_})
                d_2739_v73_: _dafny.Array
                def lambda132_(d_2740_i3_):
                    return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rgqehf")), _dafny.SeqWithoutIsStrInference([self.f27 for d_2741_i4_ in range(default__.abs(198))])])

                init79_ = lambda132_
                nw447_ = _dafny.Array(None, 29)
                for i0_79_ in range(nw447_.length(0)):
                    nw447_[i0_79_] = init79_(i0_79_)
                d_2739_v73_ = nw447_
                d_2742_v74_: _dafny.Seq
                d_2742_v74_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kldidyrjv"))
                d_2743_v75_: _dafny.Seq
                d_2743_v75_ = _dafny.SeqWithoutIsStrInference([d_2742_v74_, d_2742_v74_])
                index418_ = default__.safeIndex(228, (d_2739_v73_).length(0))
                (d_2739_v73_)[index418_] = ((d_2743_v75_) + (_dafny.SeqWithoutIsStrInference([d_2742_v74_ for d_2744_i5_ in range(default__.abs(507))]))).set(default__.safeIndex(p0, len((d_2743_v75_) + (_dafny.SeqWithoutIsStrInference([d_2742_v74_ for d_2745_i5_ in range(default__.abs(507))])))), d_2742_v74_)
                d_2746_v76_: _dafny.MultiSet
                d_2746_v76_ = _dafny.MultiSet([d_2672_v9_.f24, ((d_2675_v12_)[p0] if (p0) in (d_2675_v12_) else p1)])
                d_2747_v77_: _dafny.Map
                d_2747_v77_ = _dafny.Map({p0: d_2746_v76_})
                index419_ = default__.safeIndex(228, (d_2739_v73_).length(0))
                rhs445_ = _dafny.Map({not(d_2672_v9_.f24): d_2737_v71_})
                rhs446_ = default__.fm3(d_2676_v13_, p0, (self.f27 if p1 else self.f27), p1, globalState)
                rhs447_ = p0
                rhs448_ = ((d_2747_v77_)[p0] if (p0) in (d_2747_v77_) else (d_2746_v76_) | (_dafny.MultiSet([d_2672_v9_.f24, p1, p1])))
                rhs449_ = False
                lhs399_ = d_2739_v73_
                lhs400_ = default__.safeIndex(228, (d_2739_v73_).length(0))
                lhs401_ = globalState
                lhs402_ = globalState
                d_2738_v72_ = rhs445_
                lhs399_[lhs400_] = rhs446_
                lhs401_.f16 = rhs447_
                d_2746_v76_ = rhs448_
                lhs402_.f10 = rhs449_
            elif True:
                d_2748_v78_: _dafny.MultiSet
                d_2748_v78_ = _dafny.MultiSet([(self).f25, (d_2672_v9_).f25])
                (globalState).f11 = (d_2748_v78_).cardinality
                (globalState).f2 = p2
                d_2749_v79_: _dafny.Array
                nw448_ = _dafny.Array(None, 11)
                d_2749_v79_ = nw448_
                d_2750_v80_: T3
                nw449_ = C5()
                nw449_.ctor__(d_2672_v9_.f24, self.f27, d_2672_v9_.f24, d_2672_v9_.f24)
                d_2750_v80_ = nw449_
                index420_ = default__.safeIndex(994, (d_2749_v79_).length(0))
                (d_2749_v79_)[index420_] = d_2750_v80_
                index421_ = default__.safeIndex(994, (d_2749_v79_).length(0))
                (d_2749_v79_)[index421_] = d_2750_v80_
                (d_2750_v80_).f27 = self.f27
                d_2751_v81_: _dafny.Array
                def lambda133_(d_2752_p0_):
                    def lambda134_(d_2753_i6_):
                        return (d_2753_i6_) - (d_2752_p0_)

                    return lambda134_

                init80_ = lambda133_(p0)
                nw450_ = _dafny.Array(None, 10)
                for i0_80_ in range(nw450_.length(0)):
                    nw450_[i0_80_] = init80_(i0_80_)
                d_2751_v81_ = nw450_
                d_2754_v82_: _dafny.Map
                d_2754_v82_ = _dafny.Map({d_2751_v81_: (-642) <= (p0)})
                d_2754_v82_ = (d_2754_v82_).set(d_2751_v81_, not ((d_2750_v80_).f25) or ((self).f25))
            if d_2672_v9_.f24:
                d_2755_v83_: _dafny.Map
                d_2755_v83_ = _dafny.Map({p0: p2})
                d_2756_v84_: _dafny.Array
                nw451_ = _dafny.Array(None, 24)
                nw451_[int(0)] = p2
                nw451_[int(1)] = p2
                nw451_[int(2)] = p2
                nw451_[int(3)] = p2
                nw451_[int(4)] = p2
                nw451_[int(5)] = p2
                nw451_[int(6)] = p2
                nw451_[int(7)] = p2
                nw451_[int(8)] = p2
                nw451_[int(9)] = p2
                nw451_[int(10)] = ((d_2755_v83_)[(0) - (p0)] if ((0) - (p0)) in (d_2755_v83_) else p2)
                nw451_[int(11)] = p2
                nw451_[int(12)] = p2
                nw451_[int(13)] = p2
                nw451_[int(14)] = p2
                nw451_[int(15)] = p2
                nw451_[int(16)] = p2
                nw451_[int(17)] = p2
                nw451_[int(18)] = p2
                nw451_[int(19)] = (p2 if self.f24 else p2)
                nw451_[int(20)] = p2
                nw451_[int(21)] = p2
                nw451_[int(22)] = p2
                nw451_[int(23)] = p2
                d_2756_v84_ = nw451_
                index422_ = default__.safeIndex(638, (d_2756_v84_).length(0))
                (d_2756_v84_)[index422_] = p2
                index423_ = default__.safeIndex(638, (d_2756_v84_).length(0))
                (d_2756_v84_)[index423_] = p2
                d_2675_v12_ = (d_2675_v12_).set((_dafny.MultiSet([p0, len(_dafny.SeqWithoutIsStrInference([self.f27 for d_2757_i7_ in range(default__.abs(-440))]))])).cardinality, p1)
                rhs450_ = (p0 if p1 else (p0) - (p0))
                rhs451_ = p0
                lhs403_ = globalState
                lhs404_ = globalState
                lhs403_.f11 = rhs450_
                lhs404_.f16 = rhs451_
                (globalState).f22 = p0
                d_2758_v85_: _dafny.Seq
                d_2758_v85_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nnm"))
                d_2758_v85_ = (d_2758_v85_) + (d_2758_v85_)
            elif True:
                d_2759_v86_: C3
                nw452_ = C3()
                nw452_.ctor__(d_2672_v9_.f24, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jbwivhej")), d_2672_v9_.f24, self.f27, self.f24, (d_2672_v9_).f25)
                d_2759_v86_ = nw452_
                d_2760_v87_: D11
                d_2760_v87_ = D11_DC25(d_2759_v86_)
                d_2760_v87_ = d_2760_v87_
                d_2761_v88_: C7
                nw453_ = C7()
                nw453_.ctor__((self).f26)
                d_2761_v88_ = nw453_
                d_2762_v89_: _dafny.Set
                d_2762_v89_ = _dafny.Set({d_2761_v88_, d_2761_v88_, d_2761_v88_, d_2761_v88_, d_2761_v88_})
                d_2763_v90_: D0
                d_2763_v90_ = D0_DC0((d_2672_v9_).f25, p0, p0)
                d_2764_v91_: _dafny.MultiSet
                d_2764_v91_ = _dafny.MultiSet([p0])
                d_2765_v92_: _dafny.Array
                nw454_ = _dafny.Array(None, 17)
                nw454_[int(0)] = (self).fm6(d_2763_v90_, globalState)
                nw454_[int(1)] = (0) - (len((d_2759_v86_).f33))
                nw454_[int(2)] = p0
                nw454_[int(3)] = p0
                nw454_[int(4)] = p0
                nw454_[int(5)] = p0
                nw454_[int(6)] = p0
                nw454_[int(7)] = len((d_2759_v86_).f33)
                nw454_[int(8)] = p0
                nw454_[int(9)] = (0) - (len(d_2675_v12_))
                nw454_[int(10)] = p0
                nw454_[int(11)] = p0
                nw454_[int(12)] = (d_2764_v91_).cardinality
                nw454_[int(13)] = p0
                nw454_[int(14)] = p0
                nw454_[int(15)] = p0
                nw454_[int(16)] = p0
                d_2765_v92_ = nw454_
                d_2766_v93_: _dafny.Map
                d_2766_v93_ = _dafny.Map({d_2762_v89_: d_2765_v92_})
                d_2766_v93_ = d_2766_v93_
                d_2767_v94_: bool
                d_2768_v95_: _dafny.Array
                out41_: bool
                out42_: _dafny.Array
                out41_, out42_ = (d_2672_v9_).m0(globalState)
                d_2767_v94_ = out41_
                d_2768_v95_ = out42_
                d_2769_v96_: _dafny.Array
                def lambda135_(d_2770_p0_, d_2771_v86_, d_2772_v9_, d_2773_p1_):
                    def lambda136_(d_2774_i8_):
                        return _dafny.SeqWithoutIsStrInference([d_2770_p0_, len((d_2771_v86_).f33), len(_dafny.Map({d_2770_p0_: _dafny.Map({d_2772_v9_.f24: d_2773_p1_})}))])

                    return lambda136_

                init81_ = lambda135_(p0, d_2759_v86_, d_2672_v9_, p1)
                nw455_ = _dafny.Array(None, 7)
                for i0_81_ in range(nw455_.length(0)):
                    nw455_[i0_81_] = init81_(i0_81_)
                d_2769_v96_ = nw455_
                d_2775_v97_: _dafny.Map
                d_2775_v97_ = _dafny.Map({p0: d_2769_v96_})
                d_2776_v98_: _dafny.Array
                nw456_ = _dafny.Array(None, 23)
                nw456_[int(0)] = d_2769_v96_
                nw456_[int(1)] = d_2769_v96_
                nw456_[int(2)] = d_2769_v96_
                nw456_[int(3)] = d_2769_v96_
                nw456_[int(4)] = d_2769_v96_
                nw456_[int(5)] = d_2769_v96_
                nw456_[int(6)] = d_2769_v96_
                nw456_[int(7)] = d_2769_v96_
                nw456_[int(8)] = d_2769_v96_
                nw456_[int(9)] = d_2769_v96_
                nw456_[int(10)] = d_2769_v96_
                nw456_[int(11)] = d_2769_v96_
                nw456_[int(12)] = d_2769_v96_
                nw456_[int(13)] = d_2769_v96_
                nw456_[int(14)] = d_2769_v96_
                nw456_[int(15)] = d_2769_v96_
                nw456_[int(16)] = ((d_2775_v97_)[p0] if (p0) in (d_2775_v97_) else d_2769_v96_)
                nw456_[int(17)] = d_2769_v96_
                nw456_[int(18)] = d_2769_v96_
                nw456_[int(19)] = (d_2769_v96_ if (self).f25 else d_2769_v96_)
                nw456_[int(20)] = d_2769_v96_
                nw456_[int(21)] = d_2769_v96_
                nw456_[int(22)] = d_2769_v96_
                d_2776_v98_ = nw456_
                index424_ = default__.safeIndex(422, (d_2776_v98_).length(0))
                (d_2776_v98_)[index424_] = d_2769_v96_
                d_2777_v99_: _dafny.Seq
                d_2777_v99_ = _dafny.SeqWithoutIsStrInference([len(default__.fm53(default__.fm36((default__.fm51(self.f27, globalState)).cardinality, d_2672_v9_.f24, d_2767_v94_, globalState), self.f27, globalState))])
                d_2778_v100_: _dafny.Seq
                d_2778_v100_ = _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference([p0, p0])).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference([p0, p0]))), p0)])
                index425_ = default__.safeIndex(422, (d_2776_v98_).length(0))
                nw457_ = _dafny.Array(None, 7)
                nw457_[int(0)] = d_2777_v99_
                nw457_[int(1)] = d_2777_v99_
                nw457_[int(2)] = d_2777_v99_
                nw457_[int(3)] = (_dafny.SeqWithoutIsStrInference([(0) - (p0) for d_2779_i9_ in range(default__.abs(-91))])) + (d_2777_v99_)
                nw457_[int(4)] = d_2777_v99_
                nw457_[int(5)] = (d_2778_v100_)[default__.safeIndex(p0, len(d_2778_v100_))]
                nw457_[int(6)] = d_2777_v99_
                (d_2776_v98_)[index425_] = nw457_
                r1 = (((d_2759_v86_).f33) + ((d_2759_v86_).f33)) != (((d_2759_v86_).f33) + ((d_2759_v86_).f33))
        if self.f24:
            r1 = p1
            r0 = (self).f25
            d_2780_v101_: _dafny.Set
            d_2780_v101_ = _dafny.Set({(0) - (p0)})
            if (d_2780_v101_).ispropersubset(d_2780_v101_):
                d_2781_v102_: _dafny.Seq
                d_2781_v102_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xmyh"))
                d_2782_v103_: C3
                nw458_ = C3()
                nw458_.ctor__(p1, (d_2781_v102_ if (self).f25 else _dafny.SeqWithoutIsStrInference([self.f27 for d_2783_i10_ in range(default__.abs(233))])), self.f24, _dafny.CodePoint('j'), (p1) and (not(self.f24)), (self).f25)
                d_2782_v103_ = nw458_
                d_2784_v104_: _dafny.Map
                d_2784_v104_ = _dafny.Map({(self).f25: p0})
                d_2785_v105_: _dafny.Map
                d_2785_v105_ = _dafny.Map({(p0) + (len(d_2784_v104_)): default__.safeDivisionInt(p0, p0)})
                d_2785_v105_ = (d_2785_v105_).set(p0, p0)
                d_2786_v106_: _dafny.Seq
                d_2786_v106_ = _dafny.SeqWithoutIsStrInference([(self).fm8(self.f24, globalState)])
                d_2786_v106_ = d_2786_v106_
                index426_ = default__.safeIndex(872, (p2).length(0))
                (p2)[index426_] = ((d_2782_v103_).f32) and (self.f24)
                index427_ = default__.safeIndex(872, (p2).length(0))
                (p2)[index427_] = not((self).f25)
                d_2787_v107_: _dafny.Array
                nw459_ = _dafny.Array(_dafny.Map({}), 13)
                d_2787_v107_ = nw459_
                index428_ = default__.safeIndex(782, (d_2787_v107_).length(0))
                def iife192_():
                    coll86_ = _dafny.Map()
                    compr_86_: int
                    for compr_86_ in _dafny.IntegerRange(637, 708):
                        d_2788_v108_: int = compr_86_
                        if ((637) <= (d_2788_v108_)) and ((d_2788_v108_) < (708)):
                            coll86_[default__.safeModuloInt(d_2788_v108_, -100)] = d_2784_v104_
                    return _dafny.Map(coll86_)
                (d_2787_v107_)[index428_] = iife192_()
                
                d_2789_v109_: _dafny.Map
                d_2789_v109_ = _dafny.Map({p0: d_2784_v104_})
                index429_ = default__.safeIndex(782, (d_2787_v107_).length(0))
                (d_2787_v107_)[index429_] = (d_2789_v109_) | (d_2789_v109_)
            elif True:
                d_2790_v110_: _dafny.Array
                nw460_ = _dafny.Array(int(0), 11)
                d_2790_v110_ = nw460_
                (globalState).f7 = d_2790_v110_
                d_2790_v110_ = d_2790_v110_
                index430_ = default__.safeIndex(85, (p2).length(0))
                (p2)[index430_] = not ((self).f26) or ((self).fm8((self).f26, globalState))
                d_2791_v111_: _dafny.Seq
                d_2791_v111_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xqnxyikmi"))
                d_2792_v112_: _dafny.Seq
                d_2792_v112_ = _dafny.SeqWithoutIsStrInference([d_2791_v111_])
                d_2793_v113_: _dafny.Set
                d_2793_v113_ = _dafny.Set({self.f24})
                index431_ = default__.safeIndex(85, (p2).length(0))
                rhs452_ = not ((self).f25) or (not(p1))
                rhs453_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lfjmvuh"))
                rhs454_ = ((D19_DC41(d_2792_v112_)).cf62).set(default__.safeIndex(len(d_2793_v113_), len((D19_DC41(d_2792_v112_)).cf62)), (d_2791_v111_) + (d_2791_v111_))
                rhs455_ = p0
                lhs405_ = p2
                lhs406_ = default__.safeIndex(85, (p2).length(0))
                lhs407_ = globalState
                lhs405_[lhs406_] = rhs452_
                d_2791_v111_ = rhs453_
                d_2792_v112_ = rhs454_
                lhs407_.f16 = rhs455_
                (globalState).f11 = p0
                (globalState).f11 = ((0) - (p0)) - (len(_dafny.SeqWithoutIsStrInference([p0])))
            d_2794_v114_: _dafny.Array
            nw461_ = _dafny.Array(_dafny.MultiSet({}), 20)
            d_2794_v114_ = nw461_
            d_2795_v115_: C8
            nw462_ = C8()
            nw462_.ctor__(p1, (self).f26)
            d_2795_v115_ = nw462_
            d_2796_v116_: _dafny.Seq
            d_2796_v116_ = _dafny.SeqWithoutIsStrInference([d_2795_v115_])
            d_2797_v117_: C5
            nw463_ = C5()
            nw463_.ctor__((self).f26, self.f27, (self).f26, (self).f26)
            d_2797_v117_ = nw463_
            d_2798_v118_: _dafny.Map
            d_2798_v118_ = _dafny.Map({d_2797_v117_: p0})
            d_2799_v119_: D20
            d_2799_v119_ = D20_DC44((d_2796_v116_)[default__.safeIndex(len(d_2798_v118_), len(d_2796_v116_))])
            d_2800_v120_: _dafny.MultiSet
            d_2800_v120_ = _dafny.MultiSet([d_2795_v115_, (d_2799_v119_).cf69])
            index432_ = default__.safeIndex(829, (d_2794_v114_).length(0))
            (d_2794_v114_)[index432_] = (d_2800_v120_) | (d_2800_v120_)
            index433_ = default__.safeIndex(829, (d_2794_v114_).length(0))
            (d_2794_v114_)[index433_] = _dafny.MultiSet(d_2796_v116_)
            d_2801_v121_: _dafny.Map
            d_2801_v121_ = _dafny.Map({-42: False})
            def iife193_():
                coll87_ = _dafny.Map()
                compr_87_: int
                for compr_87_ in _dafny.IntegerRange(882, 608):
                    d_2802_v122_: int = compr_87_
                    if ((882) <= (d_2802_v122_)) and ((d_2802_v122_) < (608)):
                        coll87_[(d_2802_v122_) - (p0)] = (self).f26
                return _dafny.Map(coll87_)
            def iife194_():
                coll88_ = _dafny.Map()
                compr_88_: int
                for compr_88_ in _dafny.IntegerRange(942, 742):
                    d_2803_v123_: int = compr_88_
                    if ((942) <= (d_2803_v123_)) and ((d_2803_v123_) < (742)):
                        coll88_[(d_2803_v123_) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rgytxpfax"))))] = p1
                return _dafny.Map(coll88_)
            def iife195_():
                coll89_ = _dafny.Map()
                compr_89_: int
                for compr_89_ in (_dafny.Set({p0})).Elements:
                    d_2804_v124_: int = compr_89_
                    if (d_2804_v124_) in (_dafny.Set({p0})):
                        coll89_[(d_2804_v124_) * ((_dafny.MultiSet([p0, p0, -204, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oys"))), p0])).cardinality)] = (self).f26
                return _dafny.Map(coll89_)
            d_2801_v121_ = (d_2801_v121_).set((0) - ((_dafny.MultiSet([d_2801_v121_, iife193_()
            , iife194_()
            , iife195_()
            , d_2801_v121_])).cardinality), p1)
        elif True:
            (globalState).f10 = (self).f26
            d_2805_v125_: _dafny.Map
            d_2805_v125_ = _dafny.Map({368: p0})
            d_2806_v126_: _dafny.Map
            d_2806_v126_ = _dafny.Map({len(d_2805_v125_): self.f24})
            d_2807_v127_: _dafny.Seq
            d_2807_v127_ = _dafny.SeqWithoutIsStrInference([p1, p1])
            d_2808_v128_: _dafny.Map
            d_2808_v128_ = _dafny.Map({self.f24: len(_dafny.SeqWithoutIsStrInference([True]))})
            d_2809_v129_: _dafny.MultiSet
            d_2809_v129_ = _dafny.MultiSet([p0, (0) - (p0), p0, p0, p0])
            d_2810_v130_: _dafny.Seq
            d_2810_v130_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dgpevljab"))
            d_2811_v131_: _dafny.Array
            nw464_ = _dafny.Array(None, 17)
            nw464_[int(0)] = (0) - ((len(d_2806_v126_)) * (p0))
            nw464_[int(1)] = p0
            nw464_[int(2)] = p0
            nw464_[int(3)] = p0
            nw464_[int(4)] = p0
            nw464_[int(5)] = p0
            nw464_[int(6)] = len((_dafny.Map({(self).f26: len(d_2807_v127_)})) | (d_2808_v128_))
            nw464_[int(7)] = p0
            nw464_[int(8)] = (p0) * ((d_2809_v129_).cardinality)
            nw464_[int(9)] = default__.safeModuloInt(p0, p0)
            nw464_[int(10)] = len(d_2807_v127_)
            nw464_[int(11)] = p0
            nw464_[int(12)] = p0
            nw464_[int(13)] = p0
            nw464_[int(14)] = len(d_2810_v130_)
            nw464_[int(15)] = default__.fm0(p0, self.f24, d_2806_v126_, 497, globalState)
            nw464_[int(16)] = p0
            d_2811_v131_ = nw464_
            index434_ = default__.safeIndex(205, (d_2811_v131_).length(0))
            (d_2811_v131_)[index434_] = p0
            d_2812_v132_: _dafny.Map
            d_2812_v132_ = _dafny.Map({default__.fm1(p1, (self).f25, p0, globalState): d_2810_v130_})
            d_2813_v133_: _dafny.Seq
            d_2813_v133_ = _dafny.SeqWithoutIsStrInference([p2, p2, p2])
            d_2814_v134_: D21
            d_2814_v134_ = D21_DC47(d_2813_v133_)
            index435_ = default__.safeIndex(205, (d_2811_v131_).length(0))
            rhs456_ = (((d_2814_v134_).cf72) + (d_2813_v133_)) != (d_2813_v133_)
            rhs457_ = p0
            rhs458_ = d_2812_v132_
            rhs459_ = ((p0) + ((0) - (default__.fm0(p0, False, d_2806_v126_, p0, globalState)))) == (p0)
            rhs460_ = (self).f25
            lhs408_ = globalState
            lhs409_ = d_2811_v131_
            lhs410_ = default__.safeIndex(205, (d_2811_v131_).length(0))
            lhs411_ = globalState
            lhs412_ = self
            lhs408_.f10 = rhs456_
            lhs409_[lhs410_] = rhs457_
            d_2812_v132_ = rhs458_
            lhs411_.f10 = rhs459_
            lhs412_.f24 = rhs460_
            d_2815_v135_: T0
            nw465_ = C2()
            nw465_.ctor__(d_2810_v130_, d_2810_v130_)
            d_2815_v135_ = nw465_
            d_2816_v136_: _dafny.MultiSet
            d_2816_v136_ = _dafny.MultiSet([d_2815_v135_])
            d_2817_v137_: _dafny.Seq
            d_2817_v137_ = _dafny.SeqWithoutIsStrInference([d_2816_v136_])
            (globalState).f10 = (_dafny.MultiSet([d_2815_v135_])).issubset((d_2817_v137_)[default__.safeIndex(406, len(d_2817_v137_))])
            d_2818_v138_: _dafny.Array
            nw466_ = _dafny.Array(_dafny.Seq({}), 20)
            d_2818_v138_ = nw466_
            d_2819_v139_: _dafny.Seq
            d_2819_v139_ = _dafny.SeqWithoutIsStrInference([p0])
            index436_ = default__.safeIndex(104, (d_2818_v138_).length(0))
            (d_2818_v138_)[index436_] = d_2819_v139_
            d_2820_v140_: D6
            d_2820_v140_ = D6_DC12(d_2819_v139_)
            d_2821_v141_: _dafny.Map
            d_2821_v141_ = _dafny.Map({self.f24: (d_2820_v140_).cf23})
            index437_ = default__.safeIndex(104, (d_2818_v138_).length(0))
            (d_2818_v138_)[index437_] = ((d_2821_v141_)[p1] if (p1) in (d_2821_v141_) else d_2819_v139_)
            r1 = not((p0) > ((d_2811_v131_)[default__.safeIndex(205, (d_2811_v131_).length(0))]))
        d_2822_v142_: T1
        nw467_ = C6()
        nw467_.ctor__((self).f26, (self).f25, self.f24, self.f27, p1)
        d_2822_v142_ = nw467_
        d_2823_v143_: _dafny.MultiSet
        d_2823_v143_ = _dafny.MultiSet([d_2822_v142_])
        d_2824_v144_: _dafny.Seq
        d_2824_v144_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qck"))
        d_2825_v145_: _dafny.Map
        d_2825_v145_ = _dafny.Map({p0: len(d_2824_v144_)})
        d_2826_v146_: _dafny.Map
        d_2826_v146_ = _dafny.Map({(_dafny.MultiSet([d_2822_v142_, d_2822_v142_])).issubset(d_2823_v143_): D2_DC3(d_2825_v145_)})
        d_2827_v147_: _dafny.Seq
        d_2827_v147_ = _dafny.SeqWithoutIsStrInference([(self).f26, False])
        d_2828_v148_: D2
        d_2828_v148_ = D2_DC3(d_2825_v145_)
        d_2826_v146_ = (d_2826_v146_).set((d_2827_v147_)[default__.safeIndex(554, len(d_2827_v147_))], d_2828_v148_)
        d_2829_v149_: D11
        d_2829_v149_ = D11_DC28()
        pat_let_tv45_ = d_2824_v144_
        pat_let_tv46_ = p0
        def lambda137_(source46_):
            if source46_.is_DC26:
                d_2830___mcc_h0_ = source46_.cf47
                d_2831___mcc_h1_ = source46_.cf48
                d_2832___mcc_h2_ = source46_.cf49
                d_2833___mcc_h3_ = source46_.cf50
                d_2834_cf50_ = d_2833___mcc_h3_
                d_2835_cf49_ = d_2832___mcc_h2_
                d_2836_cf48_ = d_2831___mcc_h1_
                d_2837_cf47_ = d_2830___mcc_h0_
                return d_2835_cf49_
            elif source46_.is_DC27:
                return len(_dafny.SeqWithoutIsStrInference([pat_let_tv45_]))
            elif source46_.is_DC28:
                return (0) - (pat_let_tv46_)
            elif True:
                d_2838___mcc_h4_ = source46_.cf46
                d_2839_cf46_ = d_2838___mcc_h4_
                return (0) - (-55)

        (globalState).f11 = lambda137_(d_2829_v149_)
        hi17_ = default__.fm0(p0, default__.fm1(self.f24, p1, len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hbmmyrqy"))])), globalState), _dafny.Map({p0: p1}), p0, globalState)
        for d_2840_i11_ in range(p0, hi17_):
            d_2841_v150_: _dafny.Array
            nw468_ = _dafny.Array(_dafny.Array(None, 0), 20)
            d_2841_v150_ = nw468_
            d_2841_v150_ = d_2841_v150_
            d_2842_v151_: _dafny.Map
            d_2842_v151_ = _dafny.Map({(self).f25: d_2840_i11_})
            d_2842_v151_ = (d_2842_v151_).set((self).f25, d_2840_i11_)
            if not(((p0) >= (p0) if p1 else (self).f26)):
                d_2843_v152_: _dafny.Array
                nw469_ = _dafny.Array(D19.default()(), 20)
                d_2843_v152_ = nw469_
                d_2844_v153_: _dafny.Array
                nw470_ = _dafny.Array(None, 8)
                nw470_[int(0)] = d_2843_v152_
                nw470_[int(1)] = d_2843_v152_
                nw470_[int(2)] = d_2843_v152_
                nw470_[int(3)] = d_2843_v152_
                nw470_[int(4)] = d_2843_v152_
                nw470_[int(5)] = d_2843_v152_
                nw470_[int(6)] = d_2843_v152_
                nw470_[int(7)] = d_2843_v152_
                d_2844_v153_ = nw470_
                d_2844_v153_ = d_2844_v153_
                d_2845_v154_: _dafny.Map
                d_2845_v154_ = _dafny.Map({d_2840_i11_: d_2822_v142_.f24})
                (self).f24 = ((d_2845_v154_)[300] if (300) in (d_2845_v154_) else p1)
                d_2846_v155_: _dafny.Map
                d_2846_v155_ = _dafny.Map({_dafny.CodePoint('g'): not((_dafny.SeqWithoutIsStrInference([self.f27 for d_2847_i12_ in range(default__.abs(-348))])) <= (_dafny.SeqWithoutIsStrInference([self.f27 for d_2848_i13_ in range(default__.abs(-450))])))})
                d_2849_v156_: T3
                nw471_ = C6()
                nw471_.ctor__((self).f25, d_2822_v142_.f24, d_2822_v142_.f24, self.f27, False)
                d_2849_v156_ = nw471_
                d_2850_v157_: _dafny.Map
                d_2850_v157_ = _dafny.Map({d_2849_v156_: d_2849_v156_.f27})
                d_2846_v155_ = (d_2846_v155_).set((d_2824_v144_)[default__.safeIndex(len(d_2850_v157_), len(d_2824_v144_))], d_2822_v142_.f24)
                d_2851_v158_: C4
                nw472_ = C4()
                nw472_.ctor__(self.f24, d_2822_v142_.f24)
                d_2851_v158_ = nw472_
                d_2852_v159_: _dafny.Set
                d_2852_v159_ = _dafny.Set({d_2851_v158_})
                d_2842_v151_ = (d_2842_v151_).set(p1, (d_2840_i11_ if self.f24 else (0) - (len(d_2852_v159_))))
                d_2853_v160_: _dafny.Seq
                d_2853_v160_ = _dafny.SeqWithoutIsStrInference([p0])
                d_2854_v161_: _dafny.Map
                d_2854_v161_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference([len(d_2853_v160_)])) + (_dafny.SeqWithoutIsStrInference([987 for d_2855_i14_ in range(default__.abs(896))])): p0})
                d_2854_v161_ = (d_2854_v161_).set(d_2853_v160_, p0)
            elif True:
                d_2828_v148_ = d_2828_v148_
                d_2842_v151_ = (d_2842_v151_).set((d_2827_v147_)[default__.safeIndex(945, len(d_2827_v147_))], -754)
                d_2856_v162_: _dafny.Seq
                d_2856_v162_ = _dafny.SeqWithoutIsStrInference([d_2840_i11_])
                d_2857_v163_: _dafny.Array
                nw473_ = _dafny.Array(None, 21)
                nw473_[int(0)] = d_2824_v144_
                nw473_[int(1)] = (d_2824_v144_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vbrt")))
                nw473_[int(2)] = (d_2824_v144_) + (d_2824_v144_)
                nw473_[int(3)] = (_dafny.SeqWithoutIsStrInference([self.f27 for d_2858_i15_ in range(default__.abs(-149))])) + (d_2824_v144_)
                nw473_[int(4)] = d_2824_v144_
                nw473_[int(5)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ehjcihrpi"))
                nw473_[int(6)] = (d_2824_v144_) + (d_2824_v144_)
                nw473_[int(7)] = default__.fm29(globalState)
                nw473_[int(8)] = (d_2824_v144_) + (d_2824_v144_)
                nw473_[int(9)] = d_2824_v144_
                nw473_[int(10)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yhr"))
                nw473_[int(11)] = ((d_2824_v144_).set(default__.safeIndex(d_2840_i11_, len(d_2824_v144_)), self.f27)) + (d_2824_v144_)
                nw473_[int(12)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gj"))
                nw473_[int(13)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ahwwypl"))
                nw473_[int(14)] = d_2824_v144_
                nw473_[int(15)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "haka"))
                nw473_[int(16)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uwnexvq"))
                nw473_[int(17)] = d_2824_v144_
                nw473_[int(18)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bgyypakx"))
                nw473_[int(19)] = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u"))).set(default__.safeIndex(len(d_2856_v162_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u")))), self.f27)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hmh")))
                nw473_[int(20)] = d_2824_v144_
                d_2857_v163_ = nw473_
                index438_ = default__.safeIndex(469, (d_2857_v163_).length(0))
                (d_2857_v163_)[index438_] = (d_2824_v144_) + (d_2824_v144_)
                index439_ = default__.safeIndex(469, (d_2857_v163_).length(0))
                (d_2857_v163_)[index439_] = d_2824_v144_
                d_2859_v164_: _dafny.Map
                d_2859_v164_ = _dafny.Map({-115: self.f24})
                d_2860_v165_: D6
                d_2860_v165_ = D6_DC12(_dafny.SeqWithoutIsStrInference([p0, d_2840_i11_, 975]))
                d_2861_v166_: _dafny.Map
                d_2861_v166_ = _dafny.Map({default__.fm0(len(d_2842_v151_), (self).f25, d_2859_v164_, d_2840_i11_, globalState): d_2860_v165_})
                d_2861_v166_ = (d_2861_v166_).set(p0, d_2860_v165_)
                d_2862_v167_: C7
                nw474_ = C7()
                nw474_.ctor__(d_2822_v142_.f24)
                d_2862_v167_ = nw474_
            (globalState).f11 = len((_dafny.Map({p1: _dafny.CodePoint('m')})).set((self).f25, self.f27))
        r0 = self.f24
        d_2863_v168_: _dafny.Map
        d_2863_v168_ = _dafny.Map({106: d_2822_v142_.f24})
        d_2864_v169_: _dafny.Map
        d_2864_v169_ = _dafny.Map({d_2822_v142_.f24: self.f27})
        d_2865_v170_: D19
        d_2865_v170_ = D19_DC41(default__.fm3(d_2827_v147_, default__.fm0((self).fm6(D0_DC0((self).f25, p0, -394), globalState), p1, d_2863_v168_, p0, globalState), ((d_2864_v169_)[p1] if (p1) in (d_2864_v169_) else self.f27), (self).f26, globalState))
        pat_let_tv47_ = d_2822_v142_
        pat_let_tv48_ = d_2822_v142_
        pat_let_tv49_ = p1
        pat_let_tv50_ = p0
        def lambda138_(source47_):
            if source47_.is_DC42:
                d_2866___mcc_h5_ = source47_.cf63
                d_2867___mcc_h6_ = source47_.cf64
                d_2868_cf64_ = d_2867___mcc_h6_
                d_2869_cf63_ = d_2866___mcc_h5_
                return pat_let_tv47_.f24
            elif source47_.is_DC43:
                d_2870___mcc_h7_ = source47_.cf65
                d_2871___mcc_h8_ = source47_.cf66
                d_2872___mcc_h9_ = source47_.cf67
                d_2873___mcc_h10_ = source47_.cf68
                d_2874_cf68_ = d_2873___mcc_h10_
                d_2875_cf67_ = d_2872___mcc_h9_
                d_2876_cf66_ = d_2871___mcc_h8_
                d_2877_cf65_ = d_2870___mcc_h7_
                return pat_let_tv48_.f24
            elif True:
                d_2878___mcc_h11_ = source47_.cf62
                d_2879_cf62_ = d_2878___mcc_h11_
                return (_dafny.Set({(self).f26})) not in (_dafny.Map({_dafny.Set({pat_let_tv49_, not(not((self).f25))}): pat_let_tv50_}))

        r1 = lambda138_(d_2865_v170_)
        d_2880_v171_: _dafny.Map
        d_2880_v171_ = _dafny.Map({d_2824_v144_: p0})
        r2 = ((d_2880_v171_) | (d_2880_v171_)).set(d_2824_v144_, default__.safeModuloInt(873, p0))
        return r0, r1, r2

    def m0(self, globalState):
        r0: bool = False
        r1: _dafny.Array = _dafny.Array(None, 0)
        d_2881_v0_: int
        d_2881_v0_ = -522
        d_2882_v1_: _dafny.Seq
        d_2882_v1_ = _dafny.SeqWithoutIsStrInference([157, (0) - (d_2881_v0_)])
        d_2882_v1_ = d_2882_v1_
        d_2883_v2_: _dafny.Map
        d_2883_v2_ = _dafny.Map({(self).f26: d_2881_v0_})
        if (self).fm7(d_2883_v2_, globalState):
            if (self).f26:
                d_2884_v3_: C0
                nw475_ = C0()
                nw475_.ctor__(_dafny.SeqWithoutIsStrInference([self.f27 for d_2885_i0_ in range(default__.abs(258))]))
                d_2884_v3_ = nw475_
                d_2886_v4_: _dafny.Seq
                d_2886_v4_ = _dafny.SeqWithoutIsStrInference([True, (self).f25])
                d_2887_v5_: C1
                nw476_ = C1()
                nw476_.ctor__((d_2886_v4_)[default__.safeIndex((default__.fm27((self).f26, globalState)).cardinality, len(d_2886_v4_))])
                d_2887_v5_ = nw476_
                d_2887_v5_ = d_2887_v5_
                d_2888_v6_: _dafny.Map
                d_2888_v6_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference([(self).f26])).set(default__.safeIndex(len(d_2882_v1_), len(_dafny.SeqWithoutIsStrInference([(self).f26]))), self.f24): _dafny.SeqWithoutIsStrInference([((d_2884_v3_).f31)[default__.safeIndex(d_2881_v0_, len((d_2884_v3_).f31))] for d_2889_i1_ in range(default__.abs(477))])})
                d_2890_v7_: D0
                d_2890_v7_ = D0_DC0(self.f24, d_2881_v0_, (0) - (len(d_2882_v1_)))
                pat_let_tv51_ = d_2881_v0_
                pat_let_tv52_ = d_2881_v0_
                d_2891_v8_: _dafny.Seq
                def iife196_(_pat_let53_0):
                    def iife197_(d_2892_dt__update__tmp_h0_):
                        def iife198_(_pat_let54_0):
                            def iife199_(d_2893_dt__update_hcf1_h0_):
                                def iife200_(_pat_let55_0):
                                    def iife201_(d_2894_dt__update_hcf2_h0_):
                                        return D0_DC0((d_2892_dt__update__tmp_h0_).cf0, d_2893_dt__update_hcf1_h0_, d_2894_dt__update_hcf2_h0_)
                                    return iife201_(_pat_let55_0)
                                return iife200_(pat_let_tv52_)
                            return iife199_(_pat_let54_0)
                        return iife198_(pat_let_tv51_)
                    return iife197_(_pat_let53_0)
                d_2891_v8_ = _dafny.SeqWithoutIsStrInference([iife196_(d_2890_v7_), d_2890_v7_])
                d_2895_v9_: _dafny.Map
                d_2895_v9_ = _dafny.Map({(self).f26: d_2891_v8_})
                d_2896_v10_: _dafny.Map
                d_2896_v10_ = _dafny.Map({(d_2887_v5_).fm5(d_2895_v9_, self.f24, (self).f25, globalState): (self).f25})
                rhs461_ = ((len(((d_2888_v6_)[d_2886_v4_] if (d_2886_v4_) in (d_2888_v6_) else (d_2884_v3_).f31))) * (len(d_2896_v10_))) - (d_2881_v0_)
                rhs462_ = d_2881_v0_
                lhs413_ = globalState
                lhs414_ = globalState
                lhs413_.f11 = rhs461_
                lhs414_.f16 = rhs462_
                rhs463_ = (self).f25
                rhs464_ = d_2883_v2_
                rhs465_ = default__.safeDivisionInt(494, d_2881_v0_)
                lhs415_ = self
                lhs416_ = globalState
                lhs415_.f24 = rhs463_
                d_2883_v2_ = rhs464_
                lhs416_.f11 = rhs465_
                d_2897_v11_: _dafny.Array
                nw477_ = _dafny.Array(None, 18)
                nw477_[int(0)] = d_2881_v0_
                nw477_[int(1)] = d_2881_v0_
                nw477_[int(2)] = 435
                nw477_[int(3)] = d_2881_v0_
                nw477_[int(4)] = d_2881_v0_
                nw477_[int(5)] = len(((d_2886_v4_).set(default__.safeIndex(d_2881_v0_, len(d_2886_v4_)), (self).f25)) + (d_2886_v4_))
                nw477_[int(6)] = d_2881_v0_
                nw477_[int(7)] = default__.safeDivisionInt(d_2881_v0_, len(d_2883_v2_))
                nw477_[int(8)] = d_2881_v0_
                nw477_[int(9)] = d_2881_v0_
                nw477_[int(10)] = (d_2881_v0_) + (d_2881_v0_)
                nw477_[int(11)] = d_2881_v0_
                nw477_[int(12)] = d_2881_v0_
                nw477_[int(13)] = (0) - ((d_2882_v1_)[default__.safeIndex(d_2881_v0_, len(d_2882_v1_))])
                nw477_[int(14)] = d_2881_v0_
                nw477_[int(15)] = d_2881_v0_
                nw477_[int(16)] = (d_2887_v5_).fm5(d_2895_v9_, self.f24, (self).f25, globalState)
                nw477_[int(17)] = (d_2881_v0_ if False else (d_2882_v1_)[default__.safeIndex(d_2881_v0_, len(d_2882_v1_))])
                d_2897_v11_ = nw477_
                index440_ = default__.safeIndex(525, (d_2897_v11_).length(0))
                (d_2897_v11_)[index440_] = (len((d_2884_v3_).f31)) + (d_2881_v0_)
                index441_ = default__.safeIndex(525, (d_2897_v11_).length(0))
                (d_2897_v11_)[index441_] = d_2881_v0_
            elif True:
                (globalState).f10 = False
                (globalState).f11 = d_2881_v0_
                d_2898_v12_: _dafny.Seq
                d_2898_v12_ = _dafny.SeqWithoutIsStrInference([(self).f25])
                d_2899_v13_: T2
                nw478_ = C4()
                nw478_.ctor__((self).f25, (self).f25)
                d_2899_v13_ = nw478_
                d_2900_v14_: _dafny.Seq
                d_2900_v14_ = _dafny.SeqWithoutIsStrInference([d_2899_v13_, d_2899_v13_])
                d_2901_v15_: C6
                nw479_ = C6()
                nw479_.ctor__((self).fm8((self).f25, globalState), ((d_2898_v12_).set(default__.safeIndex(d_2881_v0_, len(d_2898_v12_)), True)) != (d_2898_v12_), not(((0) - (len(d_2900_v14_))) >= (d_2881_v0_)), self.f27, (self.f24 if False else (self).f25))
                d_2901_v15_ = nw479_
                d_2901_v15_ = (d_2901_v15_ if self.f24 else d_2901_v15_)
                d_2902_v16_: C5
                nw480_ = C5()
                nw480_.ctor__((d_2901_v15_).f28, self.f27, (self).f25, (d_2899_v13_).f25)
                d_2902_v16_ = nw480_
                d_2902_v16_ = d_2902_v16_
                (globalState).f16 = d_2881_v0_
            r0 = not(not((self).f25))
            (globalState).f10 = self.f24
            (globalState).f10 = False
            (self).f24 = ((self).f26 if (self).f26 else (self).fm7(d_2883_v2_, globalState))
        elif True:
            d_2903_v17_: _dafny.Map
            d_2903_v17_ = _dafny.Map({d_2881_v0_: (self).f25})
            d_2904_v18_: _dafny.Map
            d_2904_v18_ = _dafny.Map({True: d_2903_v17_})
            d_2904_v18_ = (d_2904_v18_).set((d_2881_v0_) <= (186), d_2903_v17_)
            d_2905_v19_: _dafny.Map
            d_2905_v19_ = _dafny.Map({(self).f25: (self).f25})
            d_2906_v20_: _dafny.Seq
            d_2906_v20_ = _dafny.SeqWithoutIsStrInference([self.f24, ((d_2905_v19_)[True] if (True) in (d_2905_v19_) else False)])
            d_2907_v21_: _dafny.MultiSet
            d_2907_v21_ = _dafny.MultiSet([d_2906_v20_, (d_2906_v20_) + (d_2906_v20_), d_2906_v20_])
            d_2908_v22_: _dafny.Seq
            d_2908_v22_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "puixrq"))
            rhs466_ = (d_2907_v21_) - (d_2907_v21_)
            rhs467_ = d_2908_v22_
            d_2907_v21_ = rhs466_
            d_2908_v22_ = rhs467_
            d_2909_v23_: C3
            nw481_ = C3()
            nw481_.ctor__(not((self).f26), d_2908_v22_, (self).f25, self.f27, (self).f25, self.f24)
            d_2909_v23_ = nw481_
            d_2910_v24_: D11
            d_2910_v24_ = D11_DC25(d_2909_v23_)
            source48_ = d_2910_v24_
            if source48_.is_DC26:
                d_2911___mcc_h0_ = source48_.cf47
                d_2912___mcc_h1_ = source48_.cf48
                d_2913___mcc_h2_ = source48_.cf49
                d_2914___mcc_h3_ = source48_.cf50
                d_2915_cf50_ = d_2914___mcc_h3_
                d_2916_cf49_ = d_2913___mcc_h2_
                d_2917_cf48_ = d_2912___mcc_h1_
                d_2918_cf47_ = d_2911___mcc_h0_
                d_2919_v25_: _dafny.Array
                def lambda139_(d_2920_v20_, d_2921_v0_):
                    def lambda140_(d_2922_i2_):
                        return (d_2920_v20_)[default__.safeIndex(d_2921_v0_, len(d_2920_v20_))]

                    return lambda140_

                init82_ = lambda139_(d_2906_v20_, d_2881_v0_)
                nw482_ = _dafny.Array(None, 21)
                for i0_82_ in range(nw482_.length(0)):
                    nw482_[i0_82_] = init82_(i0_82_)
                d_2919_v25_ = nw482_
                index442_ = default__.safeIndex(960, (d_2919_v25_).length(0))
                (d_2919_v25_)[index442_] = self.f24
                d_2923_v26_: D0
                d_2923_v26_ = D0_DC0(d_2918_cf47_, d_2881_v0_, d_2916_cf49_)
                d_2924_v27_: D0
                d_2924_v27_ = D0_DC0(d_2915_cf50_, d_2916_cf49_, (self).fm6(d_2923_v26_, globalState))
                index443_ = default__.safeIndex(960, (d_2919_v25_).length(0))
                (d_2919_v25_)[index443_] = (d_2882_v1_) <= (((_dafny.SeqWithoutIsStrInference([-558 for d_2925_i3_ in range(default__.abs(293))])).set(default__.safeIndex((0) - (d_2881_v0_), len(_dafny.SeqWithoutIsStrInference([-558 for d_2926_i3_ in range(default__.abs(293))]))), (0) - (len((d_2909_v23_).f33)))).set(default__.safeIndex(d_2916_cf49_, len((_dafny.SeqWithoutIsStrInference([-558 for d_2927_i3_ in range(default__.abs(293))])).set(default__.safeIndex((0) - (d_2881_v0_), len(_dafny.SeqWithoutIsStrInference([-558 for d_2928_i3_ in range(default__.abs(293))]))), (0) - (len((d_2909_v23_).f33))))), (d_2909_v23_).fm6(d_2924_v27_, globalState)))
                (globalState).f10 = (d_2919_v25_)[default__.safeIndex(960, (d_2919_v25_).length(0))]
                (globalState).f16 = (0) - (d_2916_cf49_)
                d_2929_v28_: D10
                d_2929_v28_ = D10_DC23(d_2906_v20_, (d_2919_v25_)[default__.safeIndex(960, (d_2919_v25_).length(0))])
                d_2930_v29_: D10
                d_2930_v29_ = D10_DC24(d_2929_v28_)
                d_2931_v30_: _dafny.Map
                d_2931_v30_ = _dafny.Map({d_2930_v29_: d_2916_cf49_})
                d_2932_v31_: _dafny.MultiSet
                d_2932_v31_ = _dafny.MultiSet([d_2931_v30_])
                d_2933_v32_: _dafny.Map
                d_2933_v32_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_2906_v20_ for d_2934_i4_ in range(default__.abs(241))]): d_2932_v31_})
                d_2935_v33_: _dafny.Seq
                d_2935_v33_ = _dafny.SeqWithoutIsStrInference([d_2906_v20_])
                d_2936_v34_: _dafny.Map
                d_2936_v34_ = _dafny.Map({d_2881_v0_: d_2935_v33_})
                d_2937_v35_: _dafny.Seq
                d_2937_v35_ = _dafny.SeqWithoutIsStrInference([d_2931_v30_, _dafny.Map({d_2930_v29_: -123})])
                d_2933_v32_ = (d_2933_v32_).set((((d_2936_v34_)[(0) - (d_2916_cf49_)] if ((0) - (d_2916_cf49_)) in (d_2936_v34_) else d_2935_v33_) if (self).f25 else d_2935_v33_), (_dafny.MultiSet(d_2937_v35_)) | (d_2932_v31_))
            elif source48_.is_DC27:
                d_2938_v36_: C0
                nw483_ = C0()
                nw483_.ctor__((d_2909_v23_).f33)
                d_2938_v36_ = nw483_
                d_2938_v36_ = d_2938_v36_
                d_2939_v37_: T1
                nw484_ = C1()
                nw484_.ctor__((d_2909_v23_).f32)
                d_2939_v37_ = nw484_
                d_2940_v38_: _dafny.Array
                nw485_ = _dafny.Array(None, 19)
                nw485_[int(0)] = d_2939_v37_
                nw485_[int(1)] = d_2939_v37_
                nw485_[int(2)] = d_2939_v37_
                nw485_[int(3)] = d_2939_v37_
                nw485_[int(4)] = d_2939_v37_
                nw485_[int(5)] = d_2939_v37_
                nw485_[int(6)] = d_2939_v37_
                nw485_[int(7)] = d_2939_v37_
                nw485_[int(8)] = d_2939_v37_
                nw485_[int(9)] = d_2939_v37_
                nw485_[int(10)] = d_2939_v37_
                nw485_[int(11)] = d_2939_v37_
                nw485_[int(12)] = d_2939_v37_
                nw485_[int(13)] = d_2939_v37_
                nw485_[int(14)] = d_2939_v37_
                nw485_[int(15)] = d_2939_v37_
                nw485_[int(16)] = d_2939_v37_
                nw485_[int(17)] = d_2939_v37_
                nw485_[int(18)] = d_2939_v37_
                d_2940_v38_ = nw485_
                d_2941_v39_: _dafny.Array
                def lambda141_(d_2942_i5_):
                    return self.f27

                init83_ = lambda141_
                nw486_ = _dafny.Array(None, 5)
                for i0_83_ in range(nw486_.length(0)):
                    nw486_[i0_83_] = init83_(i0_83_)
                d_2941_v39_ = nw486_
                d_2943_v40_: _dafny.Map
                d_2943_v40_ = _dafny.Map({d_2940_v38_: d_2941_v39_})
                r0 = (d_2943_v40_) == (d_2943_v40_)
                (globalState).f22 = (625) + (d_2881_v0_)
                d_2908_v22_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jwx"))) + ((d_2938_v36_).f31)
            elif source48_.is_DC28:
                d_2944_v41_: _dafny.Map
                d_2944_v41_ = _dafny.Map({(self).f25: (_dafny.SeqWithoutIsStrInference([self.f27 for d_2945_i6_ in range(default__.abs(235))])).set(default__.safeIndex(d_2881_v0_, len(_dafny.SeqWithoutIsStrInference([self.f27 for d_2946_i6_ in range(default__.abs(235))]))), self.f27)})
                d_2947_v42_: _dafny.Seq
                d_2947_v42_ = _dafny.SeqWithoutIsStrInference([d_2908_v22_, (d_2909_v23_).f33, (d_2909_v23_).f33, (d_2909_v23_).f33, _dafny.SeqWithoutIsStrInference([self.f27 for d_2948_i7_ in range(default__.abs(63))])])
                d_2908_v22_ = (((d_2944_v41_)[(d_2909_v23_).f32] if ((d_2909_v23_).f32) in (d_2944_v41_) else (d_2909_v23_).f33)) + ((d_2947_v42_)[default__.safeIndex(d_2881_v0_, len(d_2947_v42_))])
                d_2949_v43_: _dafny.Map
                d_2949_v43_ = _dafny.Map({self.f27: (d_2881_v0_) > (d_2881_v0_)})
                d_2949_v43_ = (d_2949_v43_).set(self.f27, self.f24)
                d_2950_v44_: _dafny.MultiSet
                d_2950_v44_ = _dafny.MultiSet([len((d_2909_v23_).f33)])
                d_2951_v45_: _dafny.Set
                d_2951_v45_ = _dafny.Set({((0) - ((d_2950_v44_).cardinality)) + (d_2881_v0_), d_2881_v0_, d_2881_v0_})
                d_2952_v46_: _dafny.Map
                d_2952_v46_ = _dafny.Map({d_2881_v0_: d_2881_v0_})
                def iife202_():
                    coll90_ = _dafny.Set()
                    compr_90_: int
                    for compr_90_ in (d_2952_v46_).keys.Elements:
                        d_2953_v47_: int = compr_90_
                        if (d_2953_v47_) in (d_2952_v46_):
                            coll90_ = coll90_.union(_dafny.Set([(d_2953_v47_) * (303)]))
                    return _dafny.Set(coll90_)
                d_2951_v45_ = iife202_()
                
                (self).f24 = (self).f25
            elif True:
                d_2954___mcc_h4_ = source48_.cf46
                d_2955_cf46_ = d_2954___mcc_h4_
                d_2956_v48_: _dafny.Set
                d_2956_v48_ = _dafny.Set({(d_2909_v23_).f32, (self).fm8((self).f26, globalState)})
                d_2957_v49_: _dafny.Map
                d_2957_v49_ = _dafny.Map({d_2956_v48_: self.f27})
                (globalState).f22 = default__.safeModuloInt(default__.safeDivisionInt(len(d_2957_v49_), d_2881_v0_), d_2881_v0_)
                d_2958_v50_: _dafny.Array
                nw487_ = _dafny.Array(None, 9)
                d_2958_v50_ = nw487_
                index444_ = default__.safeIndex(681, (d_2958_v50_).length(0))
                (d_2958_v50_)[index444_] = d_2955_cf46_
                index445_ = default__.safeIndex(681, (d_2958_v50_).length(0))
                (d_2958_v50_)[index445_] = d_2955_cf46_
                d_2959_v51_: _dafny.MultiSet
                d_2959_v51_ = _dafny.MultiSet([d_2881_v0_])
                d_2960_v52_: _dafny.MultiSet
                d_2960_v52_ = _dafny.MultiSet([((d_2959_v51_)[d_2881_v0_] if (d_2881_v0_) in (d_2959_v51_) else d_2881_v0_), len(((d_2909_v23_).f33) + (d_2908_v22_)), (len(_dafny.SeqWithoutIsStrInference([_dafny.Set({d_2881_v0_, d_2881_v0_}) for d_2961_i8_ in range(default__.abs(481))]))) * (d_2881_v0_), (0) - ((d_2881_v0_ if (d_2909_v23_).f32 else d_2881_v0_)), d_2881_v0_])
                d_2959_v51_ = d_2959_v51_
                d_2962_v53_: C1
                nw488_ = C1()
                nw488_.ctor__((d_2909_v23_).f32)
                d_2962_v53_ = nw488_
            d_2963_v54_: _dafny.Array
            nw489_ = _dafny.Array(int(0), 3)
            d_2963_v54_ = nw489_
            d_2964_v55_: _dafny.MultiSet
            d_2964_v55_ = _dafny.MultiSet([((d_2883_v2_)[(self).f25] if ((self).f25) in (d_2883_v2_) else (d_2882_v1_)[default__.safeIndex(len((d_2909_v23_).f33), len(d_2882_v1_))]), d_2881_v0_, d_2881_v0_, d_2881_v0_])
            index446_ = default__.safeIndex(631, (d_2963_v54_).length(0))
            (d_2963_v54_)[index446_] = ((d_2964_v55_)[d_2881_v0_] if (d_2881_v0_) in (d_2964_v55_) else (0) - (d_2881_v0_))
            d_2965_v56_: _dafny.Array
            nw490_ = _dafny.Array(_dafny.CodePoint('D'), 24)
            d_2965_v56_ = nw490_
            index447_ = default__.safeIndex(92, (d_2965_v56_).length(0))
            (d_2965_v56_)[index447_] = self.f27
            d_2966_v57_: _dafny.Set
            d_2966_v57_ = _dafny.Set({(self).f25})
            d_2967_v58_: D0
            d_2967_v58_ = D0_DC0((d_2909_v23_).f32, len(d_2966_v57_), d_2881_v0_)
            pat_let_tv53_ = d_2882_v1_
            pat_let_tv54_ = d_2909_v23_
            d_2968_v59_: _dafny.Seq
            def iife203_(_pat_let56_0):
                def iife204_(d_2969_dt__update__tmp_h1_):
                    def iife205_(_pat_let57_0):
                        def iife206_(d_2970_dt__update_hcf2_h1_):
                            def iife207_(_pat_let58_0):
                                def iife208_(d_2971_dt__update_hcf0_h0_):
                                    return D0_DC0(d_2971_dt__update_hcf0_h0_, (d_2969_dt__update__tmp_h1_).cf1, d_2970_dt__update_hcf2_h1_)
                                return iife208_(_pat_let58_0)
                            return iife207_((pat_let_tv54_).f32)
                        return iife206_(_pat_let57_0)
                    return iife205_(len(pat_let_tv53_))
                return iife204_(_pat_let56_0)
            d_2968_v59_ = _dafny.SeqWithoutIsStrInference([iife203_(d_2967_v58_)])
            index448_ = default__.safeIndex(631, (d_2963_v54_).length(0))
            index449_ = default__.safeIndex(92, (d_2965_v56_).length(0))
            rhs468_ = len(d_2882_v1_)
            rhs469_ = (0) - ((d_2881_v0_) * ((self).fm5(_dafny.Map({(self).fm8((self).f26, globalState): d_2968_v59_}), (d_2909_v23_).f32, (self).f26, globalState)))
            rhs470_ = (default__.safeModuloInt(d_2881_v0_, len(_dafny.SeqWithoutIsStrInference([d_2881_v0_])))) == (d_2881_v0_)
            rhs471_ = self.f27
            lhs417_ = d_2963_v54_
            lhs418_ = default__.safeIndex(631, (d_2963_v54_).length(0))
            lhs419_ = globalState
            lhs420_ = self
            lhs421_ = d_2965_v56_
            lhs422_ = default__.safeIndex(92, (d_2965_v56_).length(0))
            lhs417_[lhs418_] = rhs468_
            lhs419_.f22 = rhs469_
            lhs420_.f24 = rhs470_
            lhs421_[lhs422_] = rhs471_
            d_2972_v60_: _dafny.Array
            nw491_ = _dafny.Array(D18.default()(), 16)
            d_2972_v60_ = nw491_
            d_2973_v61_: _dafny.Map
            d_2973_v61_ = _dafny.Map({(d_2909_v23_).f33: (self).f25})
            d_2974_v62_: D18
            d_2974_v62_ = D18_DC39(d_2973_v61_)
            index450_ = default__.safeIndex(551, (d_2972_v60_).length(0))
            (d_2972_v60_)[index450_] = d_2974_v62_
            index451_ = default__.safeIndex(551, (d_2972_v60_).length(0))
            (d_2972_v60_)[index451_] = d_2974_v62_
        d_2975_v63_: _dafny.Array
        nw492_ = _dafny.Array(None, 5)
        nw492_[int(0)] = _dafny.SeqWithoutIsStrInference([len(d_2882_v1_), d_2881_v0_])
        nw492_[int(1)] = (d_2882_v1_) + (d_2882_v1_)
        nw492_[int(2)] = (d_2882_v1_) + (d_2882_v1_)
        nw492_[int(3)] = d_2882_v1_
        nw492_[int(4)] = d_2882_v1_
        d_2975_v63_ = nw492_
        index452_ = default__.safeIndex(862, (d_2975_v63_).length(0))
        (d_2975_v63_)[index452_] = d_2882_v1_
        d_2976_v64_: D6
        d_2976_v64_ = D6_DC12(d_2882_v1_)
        index453_ = default__.safeIndex(862, (d_2975_v63_).length(0))
        (d_2975_v63_)[index453_] = ((d_2976_v64_).cf23) + (_dafny.SeqWithoutIsStrInference([-709 for d_2977_i9_ in range(default__.abs(417))]))
        d_2978_v65_: _dafny.Array
        nw493_ = _dafny.Array(False, 15)
        d_2978_v65_ = nw493_
        d_2979_v66_: _dafny.Seq
        d_2979_v66_ = _dafny.SeqWithoutIsStrInference([(self).f25])
        index454_ = default__.safeIndex(249, (d_2978_v65_).length(0))
        (d_2978_v65_)[index454_] = (d_2979_v66_) == ((d_2979_v66_).set(default__.safeIndex(d_2881_v0_, len(d_2979_v66_)), self.f24))
        index455_ = default__.safeIndex(249, (d_2978_v65_).length(0))
        (d_2978_v65_)[index455_] = (self).f26
        d_2980_v67_: _dafny.Seq
        d_2980_v67_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "frqjc"))
        d_2981_v68_: D10
        d_2981_v68_ = D10_DC22((0) - (len(_dafny.Set({d_2881_v0_, len(d_2980_v67_)}))), True)
        d_2982_v69_: _dafny.Set
        d_2982_v69_ = _dafny.Set({(self.f24 if (d_2978_v65_)[default__.safeIndex(249, (d_2978_v65_).length(0))] else (d_2981_v68_).cf42)})
        d_2983_v70_: _dafny.Map
        d_2983_v70_ = _dafny.Map({not((self).f26): d_2982_v69_})
        d_2982_v69_ = (d_2982_v69_) | (((d_2983_v70_)[(self).f26] if ((self).f26) in (d_2983_v70_) else d_2982_v69_))
        d_2984_v71_: _dafny.Map
        d_2984_v71_ = _dafny.Map({(d_2978_v65_)[default__.safeIndex(249, (d_2978_v65_).length(0))]: (self).f25})
        d_2985_v72_: _dafny.Map
        d_2985_v72_ = _dafny.Map({d_2881_v0_: (0) - (d_2881_v0_)})
        d_2986_v73_: _dafny.Set
        d_2986_v73_ = _dafny.Set({d_2985_v72_})
        d_2984_v71_ = (d_2984_v71_).set((d_2986_v73_).ispropersubset(d_2986_v73_), self.f24)
        r0 = (d_2978_v65_)[default__.safeIndex(249, (d_2978_v65_).length(0))]
        nw494_ = _dafny.Array(False, 3)
        r1 = nw494_
        return r0, r1

    def m1(self, p0, globalState):
        r0: str = _dafny.CodePoint('D')
        r1: _dafny.Seq = _dafny.Seq({})
        d_2987_v0_: _dafny.Seq
        d_2987_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "unkspfxy"))
        d_2988_v1_: C2
        nw495_ = C2()
        nw495_.ctor__(d_2987_v0_, d_2987_v0_)
        d_2988_v1_ = nw495_
        d_2989_i0_: int
        d_2989_i0_ = 0
        with _dafny.label("21"):
            while not ((self).f26) or (not(self.f24)):
                with _dafny.c_label("21"):
                    if (d_2989_i0_) >= (100):
                        raise _dafny.Break("21")
                    d_2989_i0_ = (d_2989_i0_) + (1)
                    d_2990_v2_: D21
                    d_2990_v2_ = D21_DC48()
                    d_2991_v3_: _dafny.Array
                    def lambda142_(d_2992_i1_):
                        return ((D2_DC5(self.f27, (self).f25, 386, (0) - (-277))).cf11 if (self).f25 else (self).f25)

                    init84_ = lambda142_
                    nw496_ = _dafny.Array(None, 27)
                    for i0_84_ in range(nw496_.length(0)):
                        nw496_[i0_84_] = init84_(i0_84_)
                    d_2991_v3_ = nw496_
                    index456_ = default__.safeIndex(892, (d_2991_v3_).length(0))
                    (d_2991_v3_)[index456_] = (self).f25
                    d_2993_v4_: int
                    d_2993_v4_ = 144
                    d_2994_v5_: _dafny.Set
                    d_2994_v5_ = _dafny.Set({self.f24, (self).f26})
                    d_2995_v6_: _dafny.Map
                    d_2995_v6_ = _dafny.Map({default__.safeModuloInt(d_2993_v4_, d_2993_v4_): d_2994_v5_})
                    index457_ = default__.safeIndex(892, (d_2991_v3_).length(0))
                    rhs472_ = (d_2993_v4_) != (d_2993_v4_)
                    rhs473_ = len(d_2995_v6_)
                    rhs474_ = d_2990_v2_
                    rhs475_ = (self).f26
                    lhs423_ = globalState
                    lhs424_ = globalState
                    lhs425_ = d_2991_v3_
                    lhs426_ = default__.safeIndex(892, (d_2991_v3_).length(0))
                    lhs423_.f10 = rhs472_
                    lhs424_.f16 = rhs473_
                    d_2990_v2_ = rhs474_
                    lhs425_[lhs426_] = rhs475_
                    if (self).f25:
                        (globalState).f10 = self.f24
                        d_2996_v7_: _dafny.Array
                        nw497_ = _dafny.Array(_dafny.Seq({}), 14)
                        d_2996_v7_ = nw497_
                        d_2997_v8_: T1
                        nw498_ = C7()
                        nw498_.ctor__(not((self).f26))
                        d_2997_v8_ = nw498_
                        d_2998_v9_: _dafny.Seq
                        d_2998_v9_ = _dafny.SeqWithoutIsStrInference([(d_2997_v8_)])
                        d_2999_v10_: _dafny.Seq
                        d_2999_v10_ = _dafny.SeqWithoutIsStrInference([d_2998_v9_, d_2998_v9_, d_2998_v9_, d_2998_v9_])
                        index458_ = default__.safeIndex(206, (d_2996_v7_).length(0))
                        (d_2996_v7_)[index458_] = d_2999_v10_
                        index459_ = default__.safeIndex(206, (d_2996_v7_).length(0))
                        (d_2996_v7_)[index459_] = (d_2999_v10_) + ((d_2999_v10_) + (d_2999_v10_))
                        d_3000_v11_: _dafny.Array
                        nw499_ = _dafny.Array(int(0), 24)
                        d_3000_v11_ = nw499_
                        index460_ = default__.safeIndex(362, (d_3000_v11_).length(0))
                        (d_3000_v11_)[index460_] = d_2993_v4_
                        index461_ = default__.safeIndex(362, (d_3000_v11_).length(0))
                        (d_3000_v11_)[index461_] = d_2993_v4_
                        d_3001_v12_: _dafny.Map
                        d_3001_v12_ = _dafny.Map({d_2997_v8_.f24: self.f27})
                        d_3002_v13_: D0
                        d_3002_v13_ = D0_DC0(False, len(d_3001_v12_), d_2993_v4_)
                        d_3003_v14_: _dafny.Seq
                        d_3003_v14_ = _dafny.SeqWithoutIsStrInference([self.f24])
                        d_3004_v15_: _dafny.Map
                        d_3004_v15_ = _dafny.Map({True: _dafny.SeqWithoutIsStrInference([d_3002_v13_, d_3002_v13_, d_3002_v13_, D0_DC0(True, d_2993_v4_, (_dafny.MultiSet(d_3003_v14_)).cardinality)])})
                        (globalState).f11 = default__.safeDivisionInt((self).fm5(d_3004_v15_, d_2997_v8_.f24, self.f24, globalState), d_2993_v4_)
                        d_2994_v5_ = d_2994_v5_
                    elif True:
                        d_3005_v16_: _dafny.Array
                        nw500_ = _dafny.Array(D5.default()(), 22)
                        d_3005_v16_ = nw500_
                        d_3006_v17_: D2
                        d_3006_v17_ = D2_DC5(self.f27, True, d_2993_v4_, d_2993_v4_)
                        d_3007_v18_: _dafny.Seq
                        d_3007_v18_ = _dafny.SeqWithoutIsStrInference([(self).f26, (self).f25])
                        d_3008_v19_: _dafny.Seq
                        d_3008_v19_ = _dafny.SeqWithoutIsStrInference([d_2993_v4_, d_2993_v4_])
                        d_3009_v20_: D5
                        d_3009_v20_ = D5_DC11(D2_DC6(d_3006_v17_), (d_3007_v18_)[default__.safeIndex((0) - (len(d_3008_v19_)), len(d_3007_v18_))])
                        index462_ = default__.safeIndex(641, (d_3005_v16_).length(0))
                        (d_3005_v16_)[index462_] = d_3009_v20_
                        index463_ = default__.safeIndex(641, (d_3005_v16_).length(0))
                        (d_3005_v16_)[index463_] = d_3009_v20_
                        d_3010_v21_: D15
                        d_3010_v21_ = D15_DC33(d_2993_v4_)
                        (globalState).f11 = ((len(p0)) + (d_2993_v4_)) * ((d_3010_v21_).cf55)
                        d_3011_v22_: _dafny.Array
                        nw501_ = _dafny.Array(int(0), 9)
                        d_3011_v22_ = nw501_
                        d_3012_v23_: _dafny.Map
                        d_3012_v23_ = _dafny.Map({(self).f26: d_3011_v22_})
                        (globalState).f22 = (0) - ((len(d_3012_v23_)) * (d_2993_v4_))
                        d_3013_v24_: _dafny.MultiSet
                        d_3013_v24_ = _dafny.MultiSet([len(d_3012_v23_)])
                        d_3014_v25_: C1
                        nw502_ = C1()
                        nw502_.ctor__(((d_3013_v24_).cardinality) != (d_2993_v4_))
                        d_3014_v25_ = nw502_
                        d_3014_v25_ = d_3014_v25_
                        d_3015_v26_: _dafny.Seq
                        d_3015_v26_ = _dafny.SeqWithoutIsStrInference([(d_2988_v1_).f30])
                        d_3015_v26_ = d_3015_v26_
                    (globalState).f2 = d_2991_v3_
                    d_3016_v27_: _dafny.Seq
                    d_3016_v27_ = _dafny.SeqWithoutIsStrInference([self.f24])
                    d_3017_v28_: _dafny.Map
                    d_3017_v28_ = _dafny.Map({d_2993_v4_: (self).f25})
                    (globalState).f22 = (d_2993_v4_ if (d_2991_v3_)[default__.safeIndex(892, (d_2991_v3_).length(0))] else default__.safeDivisionInt(d_2993_v4_, default__.fm0(len(d_3016_v27_), True, d_3017_v28_, d_2993_v4_, globalState)))
                    pass
            pass
        d_3018_i2_: int
        d_3018_i2_ = 0
        with _dafny.label("22"):
            while (_dafny.SeqWithoutIsStrInference([(self).f25, self.f24])) < (_dafny.SeqWithoutIsStrInference([False])):
                with _dafny.c_label("22"):
                    if (d_3018_i2_) >= (100):
                        raise _dafny.Break("22")
                    d_3018_i2_ = (d_3018_i2_) + (1)
                    d_3019_v29_: int
                    d_3019_v29_ = 439
                    d_3020_v30_: _dafny.Map
                    d_3020_v30_ = _dafny.Map({d_3019_v29_: self.f24})
                    d_3020_v30_ = (d_3020_v30_).set(d_3019_v29_, self.f24)
                    d_3021_v31_: _dafny.Map
                    d_3021_v31_ = _dafny.Map({self.f24: (self).f26})
                    d_3022_v32_: D10
                    d_3022_v32_ = D10_DC22(len(d_3021_v31_), not((self).f25))
                    if (d_3022_v32_).cf42:
                        d_3023_v33_: C3
                        nw503_ = C3()
                        nw503_.ctor__(True, (d_2988_v1_).f30, (self).f25, self.f27, (d_3019_v29_) > (-615), (self.f24 if (self).f25 else self.f24))
                        d_3023_v33_ = nw503_
                        d_3024_v34_: T2
                        nw504_ = C8()
                        nw504_.ctor__(self.f24, (self).f25)
                        d_3024_v34_ = nw504_
                        d_3025_v35_: _dafny.Map
                        d_3025_v35_ = _dafny.Map({self.f24: (d_3023_v33_).f33})
                        d_3026_v36_: _dafny.Map
                        d_3026_v36_ = _dafny.Map({d_3024_v34_: (d_2988_v1_.f29) != (((d_3025_v35_)[(d_3024_v34_).f25] if ((d_3024_v34_).f25) in (d_3025_v35_) else d_2988_v1_.f29))})
                        d_3026_v36_ = (d_3026_v36_).set(d_3024_v34_, self.f24)
                        d_3027_v37_: C0
                        nw505_ = C0()
                        nw505_.ctor__((d_2987_v0_ if False else d_2988_v1_.f29))
                        d_3027_v37_ = nw505_
                        d_3028_v38_: _dafny.Array
                        nw506_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 20)
                        d_3028_v38_ = nw506_
                        index464_ = default__.safeIndex(597, (d_3028_v38_).length(0))
                        (d_3028_v38_)[index464_] = ((d_3027_v37_).f31) + ((d_2988_v1_).f30)
                        d_3029_v39_: _dafny.Seq
                        d_3029_v39_ = _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vosc"))) + ((d_2988_v1_).f30)])
                        d_3030_v40_: D18
                        d_3030_v40_ = D18_DC38(d_3019_v29_)
                        index465_ = default__.safeIndex(597, (d_3028_v38_).length(0))
                        (d_3028_v38_)[index465_] = (d_3029_v39_)[default__.safeIndex((d_3030_v40_).cf59, len(d_3029_v39_))]
                        d_3031_v41_: D2
                        d_3031_v41_ = D2_DC5(self.f27, d_3024_v34_.f24, len(default__.fm57(746, d_3020_v30_, p0, globalState)), d_3019_v29_)
                        d_3032_v42_: _dafny.Array
                        nw507_ = _dafny.Array(None, 22)
                        nw507_[int(0)] = d_3024_v34_.f24
                        nw507_[int(1)] = (d_3023_v33_).f32
                        nw507_[int(2)] = (self).f25
                        nw507_[int(3)] = (self).f26
                        nw507_[int(4)] = (d_3023_v33_).f32
                        nw507_[int(5)] = True
                        nw507_[int(6)] = (self).f26
                        nw507_[int(7)] = (d_3031_v41_).cf11
                        nw507_[int(8)] = (d_3023_v33_).f32
                        nw507_[int(9)] = (d_3023_v33_).f32
                        nw507_[int(10)] = self.f24
                        nw507_[int(11)] = (self).f26
                        nw507_[int(12)] = (self).f25
                        nw507_[int(13)] = (d_3023_v33_).f32
                        nw507_[int(14)] = (self).f26
                        nw507_[int(15)] = (self).f26
                        nw507_[int(16)] = False
                        nw507_[int(17)] = (d_3023_v33_).f32
                        nw507_[int(18)] = d_3024_v34_.f24
                        nw507_[int(19)] = self.f24
                        nw507_[int(20)] = not((d_3024_v34_).f25)
                        nw507_[int(21)] = (d_3023_v33_).f32
                        d_3032_v42_ = nw507_
                        d_3033_v43_: _dafny.Seq
                        d_3033_v43_ = _dafny.SeqWithoutIsStrInference([d_3032_v42_, d_3032_v42_, d_3032_v42_])
                        d_3034_v44_: D21
                        d_3034_v44_ = D21_DC47((d_3033_v43_ if True else d_3033_v43_))
                        index466_ = default__.safeIndex(581, (d_3032_v42_).length(0))
                        (d_3032_v42_)[index466_] = (d_3024_v34_).f25
                        d_3035_v45_: _dafny.Seq
                        d_3035_v45_ = _dafny.SeqWithoutIsStrInference([d_3019_v29_])
                        index467_ = default__.safeIndex(581, (d_3032_v42_).length(0))
                        rhs476_ = (default__.safeDivisionInt(d_3019_v29_, (0) - (d_3019_v29_))) + ((0) - (d_3019_v29_))
                        rhs477_ = d_3034_v44_
                        rhs478_ = ((0) - (d_3019_v29_)) not in (d_3035_v45_)
                        lhs427_ = globalState
                        lhs428_ = d_3032_v42_
                        lhs429_ = default__.safeIndex(581, (d_3032_v42_).length(0))
                        lhs427_.f22 = rhs476_
                        d_3034_v44_ = rhs477_
                        lhs428_[lhs429_] = rhs478_
                    elif True:
                        (globalState).f16 = d_3019_v29_
                        d_3036_v46_: C7
                        nw508_ = C7()
                        nw508_.ctor__(((self).f26) == ((self).f25))
                        d_3036_v46_ = nw508_
                        (globalState).f11 = (0) - ((d_3019_v29_) * ((0) - ((411) - (d_3019_v29_))))
                        d_3037_v47_: D0
                        d_3037_v47_ = D0_DC0(self.f24, d_3019_v29_, d_3019_v29_)
                        d_3038_v48_: _dafny.Seq
                        d_3038_v48_ = _dafny.SeqWithoutIsStrInference([d_3037_v47_, d_3037_v47_])
                        d_3039_v49_: _dafny.Map
                        d_3039_v49_ = _dafny.Map({not(self.f24): d_3038_v48_})
                        (globalState).f16 = (d_3036_v46_).fm5(d_3039_v49_, not((self).f26), default__.fm1(False, (self).f25, d_3019_v29_, globalState), globalState)
                        r0 = self.f27
                    d_3040_v50_: _dafny.Map
                    d_3040_v50_ = _dafny.Map({d_3020_v30_: (self).f25})
                    d_3041_v51_: _dafny.Seq
                    d_3041_v51_ = _dafny.SeqWithoutIsStrInference([(self).f25, self.f24])
                    d_3042_v52_: _dafny.Array
                    nw509_ = _dafny.Array(None, 8)
                    nw509_[int(0)] = not((self).f25)
                    nw509_[int(1)] = self.f24
                    nw509_[int(2)] = (self).f25
                    nw509_[int(3)] = self.f24
                    nw509_[int(4)] = (self).f25
                    nw509_[int(5)] = (self).f26
                    nw509_[int(6)] = (d_2988_v1_).fm18(len(d_3040_v50_), globalState)
                    nw509_[int(7)] = (d_3041_v51_)[default__.safeIndex(d_3019_v29_, len(d_3041_v51_))]
                    d_3042_v52_ = nw509_
                    d_3043_v53_: _dafny.Set
                    d_3043_v53_ = _dafny.Set({d_3042_v52_})
                    if (d_3043_v53_) != (d_3043_v53_):
                        d_3044_v54_: T1
                        nw510_ = C7()
                        nw510_.ctor__((self).f26)
                        d_3044_v54_ = nw510_
                        index468_ = default__.safeIndex(5, (d_3042_v52_).length(0))
                        (d_3042_v52_)[index468_] = d_3044_v54_.f24
                        index469_ = default__.safeIndex(5, (d_3042_v52_).length(0))
                        rhs479_ = d_3019_v29_
                        rhs480_ = len(p0)
                        rhs481_ = d_3044_v54_
                        rhs482_ = d_3019_v29_
                        rhs483_ = (self).f25
                        lhs430_ = globalState
                        lhs431_ = globalState
                        lhs432_ = d_3042_v52_
                        lhs433_ = default__.safeIndex(5, (d_3042_v52_).length(0))
                        d_3019_v29_ = rhs479_
                        lhs430_.f16 = rhs480_
                        d_3044_v54_ = rhs481_
                        lhs431_.f11 = rhs482_
                        lhs432_[lhs433_] = rhs483_
                        d_3045_v55_: _dafny.Array
                        nw511_ = _dafny.Array(_dafny.Array(None, 0), 2)
                        d_3045_v55_ = nw511_
                        d_3046_v56_: _dafny.Array
                        nw512_ = _dafny.Array(int(0), 21)
                        d_3046_v56_ = nw512_
                        index470_ = default__.safeIndex(630, (d_3045_v55_).length(0))
                        (d_3045_v55_)[index470_] = d_3046_v56_
                        d_3047_v57_: _dafny.Seq
                        d_3047_v57_ = _dafny.SeqWithoutIsStrInference([d_3019_v29_, d_3019_v29_, d_3019_v29_])
                        index471_ = default__.safeIndex(630, (d_3045_v55_).length(0))
                        (d_3045_v55_)[index471_] = (d_3046_v56_ if (d_3047_v57_) < (d_3047_v57_) else d_3046_v56_)
                        d_3048_v58_: D7
                        d_3048_v58_ = D7_DC15(self.f27, (self).f25)
                        d_3049_v59_: D11
                        d_3049_v59_ = D11_DC26(d_3044_v54_.f24, (d_3042_v52_)[default__.safeIndex(5, (d_3042_v52_).length(0))], d_3019_v29_, d_3044_v54_.f24)
                        d_3050_v60_: _dafny.Set
                        d_3050_v60_ = _dafny.Set({(self).f26})
                        r0 = (default__.fm21(d_2988_v1_.f29, (d_3049_v59_).cf47, d_3019_v29_, len(d_3050_v60_), globalState) if not ((d_2988_v1_).fm18(default__.fm0(d_3019_v29_, d_3044_v54_.f24, d_3020_v30_, d_3019_v29_, globalState), globalState)) or ((d_3048_v58_).cf29) else self.f27)
                        d_3051_v61_: _dafny.Seq
                        d_3051_v61_ = _dafny.SeqWithoutIsStrInference([d_3050_v60_])
                        d_3052_v62_: _dafny.Map
                        d_3052_v62_ = _dafny.Map({107: len(d_3051_v61_)})
                        d_3052_v62_ = (d_3052_v62_).set((0) - (default__.safeModuloInt(d_3019_v29_, d_3019_v29_)), d_3019_v29_)
                        (globalState).f10 = self.f24
                    elif True:
                        (globalState).f2 = d_3042_v52_
                        d_3053_v63_: _dafny.Array
                        nw513_ = _dafny.Array(int(0), 13)
                        d_3053_v63_ = nw513_
                        index472_ = default__.safeIndex(92, (d_3053_v63_).length(0))
                        (d_3053_v63_)[index472_] = 876
                        d_3054_v64_: _dafny.Map
                        d_3054_v64_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ggrkfuf")): d_3042_v52_})
                        d_3055_v65_: _dafny.Seq
                        d_3055_v65_ = _dafny.SeqWithoutIsStrInference([d_3042_v52_, d_3042_v52_, ((d_3054_v64_)[(d_2988_v1_).f30] if ((d_2988_v1_).f30) in (d_3054_v64_) else d_3042_v52_)])
                        d_3056_v66_: _dafny.MultiSet
                        d_3056_v66_ = _dafny.MultiSet([(self).f26])
                        index473_ = default__.safeIndex(92, (d_3053_v63_).length(0))
                        rhs484_ = (0) - (d_3019_v29_)
                        rhs485_ = d_3019_v29_
                        rhs486_ = (d_3055_v65_) + (d_3055_v65_)
                        rhs487_ = ((0) - ((((d_3056_v66_)[(self).f26] if ((self).f26) in (d_3056_v66_) else 598) if self.f24 else d_3019_v29_))) - (((0) - (d_3019_v29_)) * (d_3019_v29_))
                        lhs434_ = d_3053_v63_
                        lhs435_ = default__.safeIndex(92, (d_3053_v63_).length(0))
                        lhs436_ = globalState
                        lhs437_ = globalState
                        lhs434_[lhs435_] = rhs484_
                        lhs436_.f11 = rhs485_
                        d_3055_v65_ = rhs486_
                        lhs437_.f22 = rhs487_
                        d_3057_v67_: _dafny.Array
                        nw514_ = _dafny.Array(_dafny.Seq({}), 25)
                        d_3057_v67_ = nw514_
                        d_3058_v68_: _dafny.Seq
                        d_3058_v68_ = _dafny.SeqWithoutIsStrInference([d_2987_v0_])
                        index474_ = default__.safeIndex(414, (d_3057_v67_).length(0))
                        (d_3057_v67_)[index474_] = d_3058_v68_
                        index475_ = default__.safeIndex(414, (d_3057_v67_).length(0))
                        (d_3057_v67_)[index475_] = (default__.fm3(d_3041_v51_, d_3019_v29_, self.f27, self.f24, globalState)) + (d_3058_v68_)
                        d_3059_v69_: _dafny.Map
                        d_3059_v69_ = _dafny.Map({d_3056_v66_: (d_2988_v1_).f30})
                        d_3060_v70_: _dafny.Array
                        nw515_ = _dafny.Array(None, 13)
                        nw515_[int(0)] = (d_3058_v68_)[default__.safeIndex(d_3019_v29_, len(d_3058_v68_))]
                        nw515_[int(1)] = d_2987_v0_
                        nw515_[int(2)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rsqsmtehn"))
                        nw515_[int(3)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "beprrpet"))
                        nw515_[int(4)] = (d_2988_v1_).f30
                        nw515_[int(5)] = ((d_3059_v69_)[_dafny.MultiSet(d_3041_v51_)] if (_dafny.MultiSet(d_3041_v51_)) in (d_3059_v69_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "enogwbmp")))
                        nw515_[int(6)] = d_2988_v1_.f29
                        nw515_[int(7)] = (d_2988_v1_).f30
                        nw515_[int(8)] = (d_2988_v1_).f30
                        nw515_[int(9)] = d_2988_v1_.f29
                        nw515_[int(10)] = (d_2988_v1_).f30
                        nw515_[int(11)] = (d_2988_v1_).f30
                        nw515_[int(12)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tpmhlluw"))
                        d_3060_v70_ = nw515_
                        d_3061_v71_: _dafny.Map
                        d_3061_v71_ = _dafny.Map({d_3060_v70_: (d_3053_v63_)[default__.safeIndex(92, (d_3053_v63_).length(0))]})
                        d_3061_v71_ = (d_3061_v71_).set(d_3060_v70_, (222) - ((0) - ((d_3053_v63_)[default__.safeIndex(92, (d_3053_v63_).length(0))])))
                        d_3062_v72_: C0
                        nw516_ = C0()
                        nw516_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cu")))
                        d_3062_v72_ = nw516_
                        d_3063_v73_: _dafny.Map
                        d_3063_v73_ = _dafny.Map({self.f24: d_3062_v72_})
                        rhs488_ = True
                        rhs489_ = d_3063_v73_
                        rhs490_ = 647
                        lhs438_ = self
                        lhs438_.f24 = rhs488_
                        d_3063_v73_ = rhs489_
                        d_3019_v29_ = rhs490_
                    d_3064_v74_: D2
                    d_3064_v74_ = D2_DC6(default__.fm46(globalState))
                    d_3065_v75_: D5
                    d_3065_v75_ = D5_DC11(d_3064_v74_, (d_3041_v51_)[default__.safeIndex(d_3019_v29_, len(d_3041_v51_))])
                    pat_let_tv55_ = d_2988_v1_
                    pat_let_tv56_ = globalState
                    d_3066_v76_: C6
                    nw517_ = C6()
                    def iife209_(_pat_let59_0):
                        def iife210_(d_3067_dt__update__tmp_h0_):
                            def iife211_(_pat_let60_0):
                                def iife212_(d_3068_dt__update_hcf22_h0_):
                                    return D5_DC11((d_3067_dt__update__tmp_h0_).cf21, d_3068_dt__update_hcf22_h0_)
                                return iife212_(_pat_let60_0)
                            return iife211_((pat_let_tv55_).fm19(self.f27, pat_let_tv56_))
                        return iife210_(_pat_let59_0)
                    nw517_.ctor__(((self).f25) and ((self).f25), (self).f26, (iife209_(d_3065_v75_)).cf22, _dafny.CodePoint('q'), self.f24)
                    d_3066_v76_ = nw517_
                    pass
            pass
        d_3069_v77_: int
        d_3069_v77_ = 659
        d_3070_v78_: _dafny.Map
        d_3070_v78_ = _dafny.Map({d_3069_v77_: (self).f26})
        d_3070_v78_ = (d_3070_v78_).set(default__.safeDivisionInt(d_3069_v77_, d_3069_v77_), (self).f25)
        hi18_ = d_3069_v77_
        for d_3071_i3_ in range(default__.safeModuloInt(d_3069_v77_, d_3069_v77_), hi18_):
            d_3072_v79_: T2
            nw518_ = C8()
            nw518_.ctor__(False, (self).f25)
            d_3072_v79_ = nw518_
            d_3073_v80_: _dafny.Map
            d_3073_v80_ = _dafny.Map({d_3072_v79_: 573})
            d_3074_v81_: _dafny.Map
            d_3074_v81_ = _dafny.Map({self.f24: self.f24})
            d_3075_v82_: _dafny.Seq
            d_3075_v82_ = _dafny.SeqWithoutIsStrInference([d_3072_v79_.f24])
            d_3076_v83_: D0
            d_3076_v83_ = D0_DC0(((d_3074_v81_)[((d_3070_v78_)[(_dafny.MultiSet(d_3075_v82_)).cardinality] if ((_dafny.MultiSet(d_3075_v82_)).cardinality) in (d_3070_v78_) else (self).f25)] if (((d_3070_v78_)[(_dafny.MultiSet(d_3075_v82_)).cardinality] if ((_dafny.MultiSet(d_3075_v82_)).cardinality) in (d_3070_v78_) else (self).f25)) in (d_3074_v81_) else self.f24), d_3071_i3_, d_3071_i3_)
            d_3077_v84_: _dafny.MultiSet
            d_3077_v84_ = _dafny.MultiSet([(d_3076_v83_).cf2])
            d_3078_v86_: _dafny.Seq
            def iife213_():
                coll91_ = _dafny.Set()
                compr_91_: int
                for compr_91_ in (d_3077_v84_).Elements:
                    d_3079_v85_: int = compr_91_
                    if (d_3079_v85_) in (d_3077_v84_):
                        coll91_ = coll91_.union(_dafny.Set([(d_3079_v85_) - (38)]))
                return _dafny.Set(coll91_)
            d_3078_v86_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({d_3069_v77_, d_3069_v77_, d_3069_v77_}), iife213_()
            , p0])
            (globalState).f10 = ((p0) - (_dafny.Set({((d_3073_v80_)[d_3072_v79_] if (d_3072_v79_) in (d_3073_v80_) else d_3069_v77_)}))) == ((d_3078_v86_)[default__.safeIndex(-794, len(d_3078_v86_))])
            (globalState).f10 = (self).f26
            d_3080_v87_: _dafny.Array
            def lambda143_(d_3081_i3_, d_3082_v77_):
                def lambda144_(d_3083_i4_):
                    return (d_3083_i4_) - (len(_dafny.SeqWithoutIsStrInference([d_3081_i3_, d_3082_v77_, d_3081_i3_])))

                return lambda144_

            init85_ = lambda143_(d_3071_i3_, d_3069_v77_)
            nw519_ = _dafny.Array(None, 25)
            for i0_85_ in range(nw519_.length(0)):
                nw519_[i0_85_] = init85_(i0_85_)
            d_3080_v87_ = nw519_
            d_3084_v88_: _dafny.MultiSet
            d_3084_v88_ = _dafny.MultiSet([d_3080_v87_, d_3080_v87_, d_3080_v87_, d_3080_v87_, d_3080_v87_])
            d_3085_v89_: _dafny.Seq
            d_3085_v89_ = _dafny.SeqWithoutIsStrInference([495])
            d_3086_v90_: _dafny.MultiSet
            d_3086_v90_ = _dafny.MultiSet([d_3085_v89_])
            d_3087_v91_: C6
            nw520_ = C6()
            nw520_.ctor__((d_3080_v87_) in (d_3084_v88_), (d_3075_v82_)[default__.safeIndex(((d_3086_v90_)[d_3085_v89_] if (d_3085_v89_) in (d_3086_v90_) else d_3071_i3_), len(d_3075_v82_))], (self).f25, self.f27, (len(p0)) >= (d_3071_i3_))
            d_3087_v91_ = nw520_
            (d_3072_v79_).f24 = (self).f26
        d_3088_v92_: D10
        d_3088_v92_ = D10_DC22(d_3069_v77_, self.f24)
        d_3089_v93_: D2
        d_3089_v93_ = D2_DC5(self.f27, self.f24, d_3069_v77_, d_3069_v77_)
        d_3090_v94_: _dafny.Array
        nw521_ = _dafny.Array(None, 2)
        nw521_[int(0)] = (d_3088_v92_).cf41
        nw521_[int(1)] = default__.safeModuloInt((d_3089_v93_).cf12, d_3069_v77_)
        d_3090_v94_ = nw521_
        index476_ = default__.safeIndex(847, (d_3090_v94_).length(0))
        (d_3090_v94_)[index476_] = (0) - ((0) - (len(d_2988_v1_.f29)))
        d_3091_v95_: C6
        nw522_ = C6()
        nw522_.ctor__((D1_DC2(d_3090_v94_, self.f27, 985, True)).cf7, self.f24, (self).f26, self.f27, (d_3069_v77_) == (d_3069_v77_))
        d_3091_v95_ = nw522_
        d_3092_v96_: _dafny.Seq
        d_3092_v96_ = _dafny.SeqWithoutIsStrInference([(self).f26])
        d_3093_v97_: _dafny.Map
        d_3093_v97_ = _dafny.Map({(d_3092_v96_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([self.f27 for d_3094_i5_ in range(default__.abs(719))])), len(d_3092_v96_))]: d_3091_v95_})
        d_3095_v98_: _dafny.Seq
        d_3095_v98_ = _dafny.SeqWithoutIsStrInference([d_3091_v95_, ((d_3093_v97_)[(d_3091_v95_).f28] if ((d_3091_v95_).f28) in (d_3093_v97_) else d_3091_v95_)])
        d_3096_v99_: D7
        d_3096_v99_ = D7_DC14(p0)
        d_3097_v100_: D7
        d_3097_v100_ = D7_DC17(d_3096_v99_)
        pat_let_tv57_ = d_3091_v95_
        pat_let_tv58_ = d_3069_v77_
        pat_let_tv59_ = d_3091_v95_
        index477_ = default__.safeIndex(847, (d_3090_v94_).length(0))
        def lambda145_(source49_):
            if source49_.is_DC15:
                d_3098___mcc_h0_ = source49_.cf28
                d_3099___mcc_h1_ = source49_.cf29
                d_3100_cf29_ = d_3099___mcc_h1_
                d_3101_cf28_ = d_3098___mcc_h0_
                return not((pat_let_tv57_).f28)
            elif source49_.is_DC16:
                d_3102___mcc_h2_ = source49_.cf30
                d_3103___mcc_h3_ = source49_.cf31
                d_3104___mcc_h4_ = source49_.cf32
                d_3105___mcc_h5_ = source49_.cf33
                d_3106___mcc_h6_ = source49_.cf34
                d_3107_cf34_ = d_3106___mcc_h6_
                d_3108_cf33_ = d_3105___mcc_h5_
                d_3109_cf32_ = d_3104___mcc_h4_
                d_3110_cf31_ = d_3103___mcc_h3_
                d_3111_cf30_ = d_3102___mcc_h2_
                return (pat_let_tv58_) > (d_3110_cf31_)
            elif source49_.is_DC14:
                d_3112___mcc_h7_ = source49_.cf27
                d_3113_cf27_ = d_3112___mcc_h7_
                return (pat_let_tv59_).f28
            elif True:
                d_3114___mcc_h8_ = source49_.cf35
                d_3115_cf35_ = d_3114___mcc_h8_
                return self.f24

        rhs491_ = (default__.safeModuloInt(d_3069_v77_, len(d_3092_v96_))) - (len(d_2987_v0_))
        rhs492_ = default__.safeDivisionInt(d_3069_v77_, (d_3069_v77_) + (d_3069_v77_))
        rhs493_ = ((d_3095_v98_)[default__.safeIndex((0) - (d_3069_v77_), len(d_3095_v98_))] if (self).f25 else d_3091_v95_)
        rhs494_ = lambda145_(d_3097_v100_)
        lhs439_ = d_3090_v94_
        lhs440_ = default__.safeIndex(847, (d_3090_v94_).length(0))
        lhs441_ = globalState
        lhs442_ = globalState
        lhs439_[lhs440_] = rhs491_
        lhs441_.f16 = rhs492_
        d_3091_v95_ = rhs493_
        lhs442_.f10 = rhs494_
        r0 = self.f27
        r1 = d_3092_v96_
        return r0, r1

    def m4(self, p0, p1, p2, p3, globalState):
        d_3116_v0_: _dafny.Map
        d_3116_v0_ = _dafny.Map({p3: p1})
        d_3117_v1_: _dafny.MultiSet
        d_3117_v1_ = _dafny.MultiSet([True])
        d_3118_v2_: _dafny.Seq
        d_3118_v2_ = _dafny.SeqWithoutIsStrInference([((d_3117_v1_)[(self).f26] if ((self).f26) in (d_3117_v1_) else p3), p3])
        d_3119_v3_: _dafny.MultiSet
        d_3119_v3_ = _dafny.MultiSet([(p3) * (len(_dafny.Map({p3: (self).fm8(p1, globalState)}))), default__.fm0(-316, not(self.f24), d_3116_v0_, len(default__.fm44(p3, d_3118_v2_, globalState)), globalState), (331) - (963)])
        (globalState).f16 = ((d_3119_v3_)[len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "knmqv"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m"))))] if (len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "knmqv"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m"))))) in (d_3119_v3_) else (p3) + (257))
        d_3120_v4_: _dafny.Seq
        d_3120_v4_ = _dafny.SeqWithoutIsStrInference([False, False, p1])
        if (True) == ((d_3120_v4_)[default__.safeIndex(p3, len(d_3120_v4_))]):
            d_3121_v6_: _dafny.Seq
            d_3121_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a"))
            d_3122_v7_: _dafny.Seq
            d_3122_v7_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oo")), d_3121_v6_])
            d_3123_v8_: _dafny.Array
            nw523_ = _dafny.Array(None, 13)
            nw523_[int(0)] = 482
            nw523_[int(1)] = (p3 if (self).f25 else p3)
            nw523_[int(2)] = p3
            nw523_[int(3)] = p3
            nw523_[int(4)] = len(d_3120_v4_)
            nw523_[int(5)] = (0) - ((0) - (p3))
            nw523_[int(6)] = p3
            nw523_[int(7)] = p3
            def iife214_():
                coll92_ = _dafny.Set()
                compr_92_: int
                for compr_92_ in (_dafny.SeqWithoutIsStrInference([p3 for d_3124_i0_ in range(default__.abs(925))])).Elements:
                    d_3125_v5_: int = compr_92_
                    if (d_3125_v5_) in (_dafny.SeqWithoutIsStrInference([p3 for d_3124_i0_ in range(default__.abs(925))])):
                        coll92_ = coll92_.union(_dafny.Set([default__.safeDivisionInt(d_3125_v5_, 923)]))
                return _dafny.Set(coll92_)
            nw523_[int(8)] = (len(iife214_()
            )) * (p3)
            nw523_[int(9)] = default__.fm0(p3, (self).f25, d_3116_v0_, len(d_3118_v2_), globalState)
            nw523_[int(10)] = p3
            nw523_[int(11)] = default__.safeDivisionInt(len(d_3122_v7_), len(d_3121_v6_))
            nw523_[int(12)] = p3
            d_3123_v8_ = nw523_
            index478_ = default__.safeIndex(863, (d_3123_v8_).length(0))
            (d_3123_v8_)[index478_] = 664
            index479_ = default__.safeIndex(863, (d_3123_v8_).length(0))
            (d_3123_v8_)[index479_] = (p3) + (p3)
            d_3126_v9_: _dafny.Map
            d_3126_v9_ = _dafny.Map({p1: (self).fm8(p1, globalState)})
            d_3127_v10_: _dafny.Set
            d_3127_v10_ = _dafny.Set({len(d_3126_v9_), p3})
            (self).f24 = (d_3127_v10_).isdisjoint(d_3127_v10_)
            d_3128_v11_: _dafny.MultiSet
            d_3128_v11_ = d_3119_v3_
            d_3129_v12_: _dafny.MultiSet
            d_3129_v12_ = _dafny.MultiSet([d_3128_v11_, d_3128_v11_, d_3128_v11_, d_3128_v11_])
            def iife215_():
                coll93_ = _dafny.Map()
                compr_93_: int
                for compr_93_ in _dafny.IntegerRange(-477, 322):
                    d_3130_v13_: int = compr_93_
                    if ((-477) <= (d_3130_v13_)) and ((d_3130_v13_) < (322)):
                        coll93_[(d_3130_v13_) * (64)] = (self).f25
                return _dafny.Map(coll93_)
            (globalState).f11 = default__.safeModuloInt(((d_3129_v12_).intersection(d_3129_v12_)).cardinality, default__.fm0(((d_3117_v1_)[(self).f25] if ((self).f25) in (d_3117_v1_) else (0) - (p3)), (self).f25, (d_3116_v0_).set(default__.fm0(len(d_3121_v6_), default__.fm1(p1, (self).f26, (d_3123_v8_)[default__.safeIndex(863, (d_3123_v8_).length(0))], globalState), iife215_()
            , p3, globalState), p1), len(d_3121_v6_), globalState))
            source50_ = (self).fm4(d_3117_v1_, d_3121_v6_, globalState)
            d_3131___mcc_h0_ = source50_.cf0
            d_3132___mcc_h1_ = source50_.cf1
            d_3133___mcc_h2_ = source50_.cf2
            d_3134_cf2_ = d_3133___mcc_h2_
            d_3135_cf1_ = d_3132___mcc_h1_
            d_3136_cf0_ = d_3131___mcc_h0_
            d_3137_v14_: _dafny.Array
            def lambda146_(d_3138_i1_):
                return (self).f26

            init86_ = lambda146_
            nw524_ = _dafny.Array(None, 22)
            for i0_86_ in range(nw524_.length(0)):
                nw524_[i0_86_] = init86_(i0_86_)
            d_3137_v14_ = nw524_
            index480_ = default__.safeIndex(577, (d_3137_v14_).length(0))
            (d_3137_v14_)[index480_] = (self.f27) in (d_3121_v6_)
            d_3139_v15_: D0
            d_3139_v15_ = D0_DC0(p1, len(default__.fm28(globalState)), p3)
            index481_ = default__.safeIndex(577, (d_3137_v14_).length(0))
            (d_3137_v14_)[index481_] = ((d_3136_cf0_ if (self).f25 else (d_3139_v15_).cf0)) or ((len(d_3120_v4_)) < ((d_3123_v8_)[default__.safeIndex(863, (d_3123_v8_).length(0))]))
            d_3140_v16_: C4
            nw525_ = C4()
            nw525_.ctor__(not(((d_3137_v14_)[default__.safeIndex(577, (d_3137_v14_).length(0))] if self.f24 else (self).f25)), (d_3119_v3_).issubset(d_3119_v3_))
            d_3140_v16_ = nw525_
            d_3141_v17_: _dafny.Array
            nw526_ = _dafny.Array(None, 11)
            d_3141_v17_ = nw526_
            d_3142_v18_: C5
            nw527_ = C5()
            nw527_.ctor__((self).f25, self.f27, (self).f26, (self).f25)
            d_3142_v18_ = nw527_
            index482_ = default__.safeIndex(28, (d_3141_v17_).length(0))
            (d_3141_v17_)[index482_] = d_3142_v18_
            index483_ = default__.safeIndex(28, (d_3141_v17_).length(0))
            rhs495_ = (d_3135_cf1_) >= ((d_3123_v8_)[default__.safeIndex(863, (d_3123_v8_).length(0))])
            rhs496_ = d_3142_v18_
            rhs497_ = d_3134_cf2_
            lhs443_ = d_3141_v17_
            lhs444_ = default__.safeIndex(28, (d_3141_v17_).length(0))
            lhs445_ = globalState
            d_3136_cf0_ = rhs495_
            lhs443_[lhs444_] = rhs496_
            lhs445_.f11 = rhs497_
            (self).f24 = (default__.fm26(((d_3116_v0_)[-235] if (-235) in (d_3116_v0_) else (d_3137_v14_)[default__.safeIndex(577, (d_3137_v14_).length(0))]), (d_3123_v8_)[default__.safeIndex(863, (d_3123_v8_).length(0))], len(d_3121_v6_), globalState)) <= (d_3121_v6_)
            d_3143_v19_: D0
            d_3143_v19_ = D0_DC0((self).f26, (d_3123_v8_)[default__.safeIndex(863, (d_3123_v8_).length(0))], (d_3123_v8_)[default__.safeIndex(863, (d_3123_v8_).length(0))])
            d_3144_v20_: _dafny.Seq
            d_3144_v20_ = _dafny.SeqWithoutIsStrInference([d_3143_v19_])
            d_3145_v21_: _dafny.Map
            d_3145_v21_ = _dafny.Map({(self).f26: d_3144_v20_})
            if ((self).fm5(d_3145_v21_, (self).f26, p1, globalState)) <= (p3):
                d_3146_v22_: C3
                nw528_ = C3()
                nw528_.ctor__((154) >= (default__.fm0(917, (self).f25, d_3116_v0_, len(d_3121_v6_), globalState)), d_3121_v6_, (91) >= (p3), self.f27, self.f24, p1)
                d_3146_v22_ = nw528_
                d_3147_v23_: _dafny.Map
                d_3147_v23_ = _dafny.Map({(d_3146_v22_).f33: d_3121_v6_})
                d_3148_v24_: _dafny.Map
                d_3148_v24_ = _dafny.Map({(self).f25: d_3147_v23_})
                d_3148_v24_ = (d_3148_v24_).set((self).f26, (d_3147_v23_) | (d_3147_v23_))
                d_3149_v25_: C6
                nw529_ = C6()
                nw529_.ctor__(False, p1, (self).f26, self.f27, (self).f25)
                d_3149_v25_ = nw529_
                d_3149_v25_ = d_3149_v25_
                d_3150_v26_: C5
                nw530_ = C5()
                nw530_.ctor__(not(False), self.f27, (self).f25, not((p1) and ((self).f26)))
                d_3150_v26_ = nw530_
                d_3151_v27_: _dafny.Array
                def lambda147_(d_3152_p3_, d_3153_v3_):
                    def lambda148_(d_3154_i2_):
                        return _dafny.Map({d_3152_p3_: (d_3153_v3_).cardinality})

                    return lambda148_

                init87_ = lambda147_(p3, d_3119_v3_)
                nw531_ = _dafny.Array(None, 18)
                for i0_87_ in range(nw531_.length(0)):
                    nw531_[i0_87_] = init87_(i0_87_)
                d_3151_v27_ = nw531_
                d_3155_v28_: _dafny.Map
                d_3155_v28_ = _dafny.Map({((_dafny.MultiSet(d_3120_v4_)).set((d_3146_v22_).f32, default__.abs((d_3123_v8_)[default__.safeIndex(863, (d_3123_v8_).length(0))]))).cardinality: d_3151_v27_})
                d_3156_v29_: _dafny.Array
                nw532_ = _dafny.Array(None, 21)
                nw532_[int(0)] = d_3151_v27_
                nw532_[int(1)] = d_3151_v27_
                nw532_[int(2)] = d_3151_v27_
                nw532_[int(3)] = (d_3151_v27_ if not(not(False)) else d_3151_v27_)
                nw532_[int(4)] = d_3151_v27_
                nw532_[int(5)] = d_3151_v27_
                nw532_[int(6)] = d_3151_v27_
                nw532_[int(7)] = d_3151_v27_
                nw532_[int(8)] = d_3151_v27_
                nw532_[int(9)] = d_3151_v27_
                nw532_[int(10)] = d_3151_v27_
                nw532_[int(11)] = d_3151_v27_
                nw532_[int(12)] = d_3151_v27_
                nw532_[int(13)] = d_3151_v27_
                nw532_[int(14)] = d_3151_v27_
                nw532_[int(15)] = d_3151_v27_
                nw532_[int(16)] = (((d_3155_v28_)[(d_3123_v8_)[default__.safeIndex(863, (d_3123_v8_).length(0))]] if ((d_3123_v8_)[default__.safeIndex(863, (d_3123_v8_).length(0))]) in (d_3155_v28_) else d_3151_v27_) if (d_3120_v4_)[default__.safeIndex(p3, len(d_3120_v4_))] else d_3151_v27_)
                nw532_[int(17)] = d_3151_v27_
                nw532_[int(18)] = d_3151_v27_
                nw532_[int(19)] = d_3151_v27_
                nw532_[int(20)] = d_3151_v27_
                d_3156_v29_ = nw532_
                index484_ = default__.safeIndex(318, (d_3156_v29_).length(0))
                (d_3156_v29_)[index484_] = d_3151_v27_
                index485_ = default__.safeIndex(318, (d_3156_v29_).length(0))
                rhs498_ = True
                rhs499_ = d_3151_v27_
                lhs446_ = self
                lhs447_ = d_3156_v29_
                lhs448_ = default__.safeIndex(318, (d_3156_v29_).length(0))
                lhs446_.f24 = rhs498_
                lhs447_[lhs448_] = rhs499_
            elif True:
                d_3157_v30_: _dafny.Array
                def lambda149_(d_3158_i3_):
                    return self.f24

                init88_ = lambda149_
                nw533_ = _dafny.Array(None, 28)
                for i0_88_ in range(nw533_.length(0)):
                    nw533_[i0_88_] = init88_(i0_88_)
                d_3157_v30_ = nw533_
                d_3159_v31_: _dafny.Set
                d_3159_v31_ = _dafny.Set({d_3157_v30_, d_3157_v30_})
                d_3160_v32_: _dafny.Map
                d_3160_v32_ = _dafny.Map({self.f24: (d_3159_v31_) | (d_3159_v31_)})
                d_3160_v32_ = (d_3160_v32_).set(p1, d_3159_v31_)
                (globalState).f11 = (d_3123_v8_)[default__.safeIndex(863, (d_3123_v8_).length(0))]
                index486_ = default__.safeIndex(863, (d_3123_v8_).length(0))
                (d_3123_v8_)[index486_] = p3
                d_3161_v34_: _dafny.Array
                def lambda150_(d_3162_v8_):
                    def lambda151_(d_3163_i4_):
                        def iife216_():
                            coll94_ = _dafny.Map()
                            compr_94_: D3
                            for compr_94_ in (_dafny.SeqWithoutIsStrInference([D3_DC8(self.f24, len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({-555: (d_3162_v8_)[default__.safeIndex(863, (d_3162_v8_).length(0))]}))])), 477) for d_3164_i5_ in range(default__.abs(577))])).Elements:
                                d_3165_v33_: D3 = compr_94_
                                if (d_3165_v33_) in (_dafny.SeqWithoutIsStrInference([D3_DC8(self.f24, len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({-555: (d_3162_v8_)[default__.safeIndex(863, (d_3162_v8_).length(0))]}))])), 477) for d_3164_i5_ in range(default__.abs(577))])):
                                    coll94_[d_3165_v33_] = self.f27
                            return _dafny.Map(coll94_)
                        return (iife216_()
                        ) | (_dafny.Map({D3_DC8((self).f25, len(_dafny.Map({self.f27: (d_3162_v8_)[default__.safeIndex(863, (d_3162_v8_).length(0))]})), (d_3162_v8_)[default__.safeIndex(863, (d_3162_v8_).length(0))]): self.f27}))

                    return lambda151_

                init89_ = lambda150_(d_3123_v8_)
                nw534_ = _dafny.Array(None, 4)
                for i0_89_ in range(nw534_.length(0)):
                    nw534_[i0_89_] = init89_(i0_89_)
                d_3161_v34_ = nw534_
                d_3166_v35_: _dafny.Array
                nw535_ = _dafny.Array(_dafny.CodePoint('D'), 18)
                d_3166_v35_ = nw535_
                index487_ = default__.safeIndex(889, (d_3166_v35_).length(0))
                (d_3166_v35_)[index487_] = (self.f27 if self.f24 else self.f27)
                d_3167_v36_: _dafny.Array
                nw536_ = _dafny.Array(_dafny.Seq({}), 29)
                d_3167_v36_ = nw536_
                d_3168_v37_: D19
                d_3168_v37_ = D19_DC42((self).f26, d_3167_v36_)
                index488_ = default__.safeIndex(889, (d_3166_v35_).length(0))
                index489_ = default__.safeIndex(863, (d_3123_v8_).length(0))
                rhs500_ = d_3161_v34_
                rhs501_ = self.f27
                rhs502_ = d_3168_v37_
                rhs503_ = (p3) - ((d_3123_v8_)[default__.safeIndex(863, (d_3123_v8_).length(0))])
                lhs449_ = d_3166_v35_
                lhs450_ = default__.safeIndex(889, (d_3166_v35_).length(0))
                lhs451_ = d_3123_v8_
                lhs452_ = default__.safeIndex(863, (d_3123_v8_).length(0))
                d_3161_v34_ = rhs500_
                lhs449_[lhs450_] = rhs501_
                d_3168_v37_ = rhs502_
                lhs451_[lhs452_] = rhs503_
                index490_ = default__.safeIndex(863, (d_3123_v8_).length(0))
                (d_3123_v8_)[index490_] = ((d_3118_v2_)[default__.safeIndex(p3, len(d_3118_v2_))]) - (p3)
        elif True:
            (self).f24 = (default__.fm29(globalState)) == (_dafny.SeqWithoutIsStrInference([self.f27 for d_3169_i6_ in range(default__.abs(816))]))
            rhs504_ = -649
            rhs505_ = p3
            lhs453_ = globalState
            lhs454_ = globalState
            lhs453_.f16 = rhs504_
            lhs454_.f11 = rhs505_
            rhs506_ = p3
            rhs507_ = (self).f25
            lhs455_ = globalState
            lhs456_ = self
            lhs455_.f16 = rhs506_
            lhs456_.f24 = rhs507_
            d_3170_v38_: C1
            nw537_ = C1()
            nw537_.ctor__(not(p1))
            d_3170_v38_ = nw537_
            d_3171_v39_: _dafny.MultiSet
            d_3171_v39_ = d_3119_v3_
            source51_ = D19_DC43(default__.safeDivisionInt(423, p3), d_3170_v38_, (d_3171_v39_), default__.safeModuloInt(p3, p3))
            if source51_.is_DC42:
                d_3172___mcc_h3_ = source51_.cf63
                d_3173___mcc_h4_ = source51_.cf64
                d_3174_cf64_ = d_3173___mcc_h4_
                d_3175_cf63_ = d_3172___mcc_h3_
                d_3176_v40_: _dafny.Array
                nw538_ = _dafny.Array(None, 6)
                nw538_[int(0)] = False
                nw538_[int(1)] = True
                nw538_[int(2)] = p1
                nw538_[int(3)] = True
                nw538_[int(4)] = (d_3120_v4_)[default__.safeIndex(p3, len(d_3120_v4_))]
                nw538_[int(5)] = (self).f25
                d_3176_v40_ = nw538_
                index491_ = default__.safeIndex(554, (d_3176_v40_).length(0))
                (d_3176_v40_)[index491_] = not(self.f24)
                index492_ = default__.safeIndex(554, (d_3176_v40_).length(0))
                (d_3176_v40_)[index492_] = d_3175_cf63_
                (globalState).f11 = p3
                d_3177_v41_: _dafny.Array
                nw539_ = _dafny.Array(None, 7)
                nw539_[int(0)] = d_3176_v40_
                nw539_[int(1)] = d_3176_v40_
                nw539_[int(2)] = d_3176_v40_
                nw539_[int(3)] = d_3176_v40_
                nw539_[int(4)] = d_3176_v40_
                nw539_[int(5)] = d_3176_v40_
                nw539_[int(6)] = d_3176_v40_
                d_3177_v41_ = nw539_
                d_3178_v42_: _dafny.Array
                d_3178_v42_ = d_3177_v41_
                d_3178_v42_ = d_3177_v41_
                d_3179_v43_: _dafny.Array
                def lambda152_(d_3180_v2_):
                    def lambda153_(d_3181_i7_):
                        return _dafny.MultiSet(d_3180_v2_)

                    return lambda153_

                init90_ = lambda152_(d_3118_v2_)
                nw540_ = _dafny.Array(None, 4)
                for i0_90_ in range(nw540_.length(0)):
                    nw540_[i0_90_] = init90_(i0_90_)
                d_3179_v43_ = nw540_
                d_3182_v44_: D10
                d_3182_v44_ = D10_DC23(d_3120_v4_, (self).f26)
                index493_ = default__.safeIndex(894, (d_3179_v43_).length(0))
                (d_3179_v43_)[index493_] = default__.fm52(not((self).fm8((d_3182_v44_).cf44, globalState)), globalState)
                index494_ = default__.safeIndex(894, (d_3179_v43_).length(0))
                (d_3179_v43_)[index494_] = (d_3119_v3_).set(default__.safeModuloInt(p3, p3), default__.abs(default__.safeDivisionInt(p3, 582)))
            elif source51_.is_DC43:
                d_3183___mcc_h5_ = source51_.cf65
                d_3184___mcc_h6_ = source51_.cf66
                d_3185___mcc_h7_ = source51_.cf67
                d_3186___mcc_h8_ = source51_.cf68
                d_3187_cf68_ = d_3186___mcc_h8_
                d_3188_cf67_ = d_3185___mcc_h7_
                d_3189_cf66_ = d_3184___mcc_h6_
                d_3190_cf65_ = d_3183___mcc_h5_
                d_3118_v2_ = ((d_3118_v2_).set(default__.safeIndex(p3, len(d_3118_v2_)), p3)) + (d_3118_v2_)
                d_3191_v45_: T1
                nw541_ = C1()
                nw541_.ctor__((self).f26)
                d_3191_v45_ = nw541_
                d_3192_v46_: _dafny.Map
                d_3192_v46_ = _dafny.Map({d_3191_v45_: p3})
                d_3190_cf65_ = ((d_3192_v46_)[d_3191_v45_] if (d_3191_v45_) in (d_3192_v46_) else d_3187_cf68_)
                d_3119_v3_ = _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([p3 for d_3193_i8_ in range(default__.abs(534))]))
                d_3194_v47_: _dafny.Seq
                d_3194_v47_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([(0) - (d_3190_cf65_)])])
                d_3195_v48_: _dafny.Map
                d_3195_v48_ = _dafny.Map({default__.fm1(d_3191_v45_.f24, d_3191_v45_.f24, 914, globalState): (self).f25})
                d_3196_v49_: _dafny.Set
                d_3196_v49_ = _dafny.Set({d_3187_cf68_, d_3190_cf65_, len(_dafny.SeqWithoutIsStrInference([self.f27 for d_3197_i9_ in range(default__.abs(210))])), p3, len(d_3195_v48_)})
                d_3198_v50_: D7
                d_3198_v50_ = D7_DC14(d_3196_v49_)
                d_3199_v51_: D7
                d_3199_v51_ = D7_DC17(d_3198_v50_)
                d_3200_v52_: _dafny.Map
                d_3200_v52_ = _dafny.Map({d_3187_cf68_: d_3117_v1_})
                rhs508_ = default__.fm65(d_3191_v45_.f24, d_3200_v52_, d_3190_cf65_, 955, globalState)
                rhs509_ = d_3199_v51_
                d_3194_v47_ = rhs508_
                d_3199_v51_ = rhs509_
            elif True:
                d_3201___mcc_h9_ = source51_.cf62
                d_3202_cf62_ = d_3201___mcc_h9_
                (globalState).f10 = False
                d_3203_v53_: _dafny.Seq
                d_3203_v53_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mmtwdebdf"))
                d_3203_v53_ = d_3203_v53_
                (globalState).f10 = p1
                d_3204_v54_: C4
                nw542_ = C4()
                nw542_.ctor__((self).f26, self.f24)
                d_3204_v54_ = nw542_
            d_3205_v55_: _dafny.Array
            nw543_ = _dafny.Array(False, 16)
            d_3205_v55_ = nw543_
            d_3206_v56_: _dafny.Map
            d_3206_v56_ = _dafny.Map({d_3205_v55_: p3})
            d_3207_v57_: _dafny.Seq
            d_3207_v57_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qchbi"))
            d_3208_v58_: _dafny.Map
            d_3208_v58_ = _dafny.Map({p3: len(d_3207_v57_)})
            d_3209_v59_: _dafny.Map
            d_3209_v59_ = _dafny.Map({(d_3206_v56_) | ((d_3206_v56_).set(d_3205_v55_, len(d_3207_v57_))): d_3208_v58_})
            d_3210_v60_: _dafny.Map
            d_3210_v60_ = _dafny.Map({p1: self.f27})
            d_3209_v59_ = (d_3209_v59_).set(_dafny.Map({d_3205_v55_: p3}), default__.fm35((self).f26, d_3210_v60_, globalState))
        d_3211_v61_: C5
        nw544_ = C5()
        nw544_.ctor__(p1, self.f27, (self).f26, p1)
        d_3211_v61_ = nw544_
        d_3212_v62_: _dafny.Map
        d_3212_v62_ = _dafny.Map({d_3211_v61_: (self).f25})
        if (d_3211_v61_) in (d_3212_v62_):
            d_3213_v63_: _dafny.Array
            nw545_ = _dafny.Array(None, 14)
            d_3213_v63_ = nw545_
            d_3214_v64_: _dafny.Set
            d_3214_v64_ = _dafny.Set({d_3213_v63_, d_3213_v63_})
            d_3214_v64_ = ((d_3214_v64_) | (_dafny.Set({d_3213_v63_, d_3213_v63_}))) | (d_3214_v64_)
            d_3215_v65_: _dafny.Map
            d_3215_v65_ = _dafny.Map({(self).f25: p3})
            d_3216_v66_: _dafny.MultiSet
            d_3216_v66_ = _dafny.MultiSet([_dafny.Map({(self).f26: p3}), d_3215_v65_])
            d_3217_v67_: D5
            d_3217_v67_ = D5_DC10(_dafny.MultiSet([d_3215_v65_, d_3215_v65_, d_3215_v65_, d_3215_v65_, d_3215_v65_]))
            source52_ = (D5_DC10(d_3216_v66_) if (self).f25 else d_3217_v67_)
            if source52_.is_DC11:
                d_3218___mcc_h10_ = source52_.cf21
                d_3219___mcc_h11_ = source52_.cf22
                d_3220_cf22_ = d_3219___mcc_h11_
                d_3221_cf21_ = d_3218___mcc_h10_
                d_3222_v68_: _dafny.Array
                def lambda154_(d_3223_i10_):
                    return True

                init91_ = lambda154_
                nw546_ = _dafny.Array(None, 28)
                for i0_91_ in range(nw546_.length(0)):
                    nw546_[i0_91_] = init91_(i0_91_)
                d_3222_v68_ = nw546_
                index495_ = default__.safeIndex(143, (d_3222_v68_).length(0))
                (d_3222_v68_)[index495_] = self.f24
                d_3224_v69_: _dafny.Set
                d_3224_v69_ = _dafny.Set({d_3116_v0_, d_3116_v0_, d_3116_v0_})
                d_3225_v70_: _dafny.Set
                d_3225_v70_ = _dafny.Set({d_3222_v68_, d_3222_v68_})
                d_3226_v71_: _dafny.Seq
                d_3226_v71_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tp"))
                d_3227_v72_: T1
                nw547_ = C1()
                nw547_.ctor__(self.f24)
                d_3227_v72_ = nw547_
                d_3228_v73_: _dafny.MultiSet
                d_3228_v73_ = _dafny.MultiSet([d_3227_v72_])
                d_3229_v74_: D8
                d_3229_v74_ = D8_DC19((self).f25, p3)
                d_3230_v75_: _dafny.MultiSet
                d_3230_v75_ = _dafny.MultiSet([d_3226_v71_])
                d_3231_v76_: _dafny.Map
                d_3231_v76_ = _dafny.Map({p1: not(self.f24)})
                d_3232_v77_: D0
                d_3232_v77_ = D0_DC0(d_3220_cf22_, (0) - (((d_3230_v75_)[_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d"))] if (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d"))) in (d_3230_v75_) else len(d_3231_v76_))), (0) - (p3))
                d_3233_v78_: _dafny.Array
                nw548_ = _dafny.Array(None, 17)
                nw548_[int(0)] = default__.safeModuloInt(p3, len(d_3224_v69_))
                nw548_[int(1)] = ((d_3117_v1_).cardinality if (self).f25 else 971)
                nw548_[int(2)] = p3
                nw548_[int(3)] = len((d_3225_v70_) | (d_3225_v70_))
                nw548_[int(4)] = 790
                nw548_[int(5)] = (p3) - (p3)
                nw548_[int(6)] = p3
                nw548_[int(7)] = default__.safeDivisionInt(len(d_3226_v71_), (d_3228_v73_).cardinality)
                nw548_[int(8)] = (d_3229_v74_).cf38
                nw548_[int(9)] = len(d_3226_v71_)
                nw548_[int(10)] = 610
                nw548_[int(11)] = p3
                nw548_[int(12)] = p3
                nw548_[int(13)] = p3
                nw548_[int(14)] = (d_3211_v61_).fm6(d_3232_v77_, globalState)
                nw548_[int(15)] = p3
                nw548_[int(16)] = p3
                d_3233_v78_ = nw548_
                index496_ = default__.safeIndex(53, (d_3233_v78_).length(0))
                (d_3233_v78_)[index496_] = p3
                d_3234_v79_: _dafny.Array
                nw549_ = _dafny.Array(D6.default()(), 25)
                d_3234_v79_ = nw549_
                d_3235_v80_: D7
                d_3235_v80_ = D7_DC16(self.f24, p3, d_3220_cf22_, p1, d_3220_cf22_)
                d_3236_v81_: C2
                nw550_ = C2()
                nw550_.ctor__(d_3226_v71_, d_3226_v71_)
                d_3236_v81_ = nw550_
                pat_let_tv60_ = p3
                pat_let_tv61_ = d_3227_v72_
                pat_let_tv62_ = d_3120_v4_
                pat_let_tv63_ = d_3120_v4_
                index497_ = default__.safeIndex(143, (d_3222_v68_).length(0))
                index498_ = default__.safeIndex(53, (d_3233_v78_).length(0))
                def iife217_(_pat_let61_0):
                    def iife218_(d_3237_dt__update__tmp_h0_):
                        def iife219_(_pat_let62_0):
                            def iife220_(d_3238_dt__update_hcf31_h0_):
                                def iife221_(_pat_let63_0):
                                    def iife222_(d_3239_dt__update_hcf33_h0_):
                                        def iife223_(_pat_let64_0):
                                            def iife224_(d_3240_dt__update_hcf32_h0_):
                                                return D7_DC16((d_3237_dt__update__tmp_h0_).cf30, d_3238_dt__update_hcf31_h0_, d_3240_dt__update_hcf32_h0_, d_3239_dt__update_hcf33_h0_, (d_3237_dt__update__tmp_h0_).cf34)
                                            return iife224_(_pat_let64_0)
                                        return iife223_((pat_let_tv62_) not in (_dafny.SeqWithoutIsStrInference([pat_let_tv63_])))
                                    return iife222_(_pat_let63_0)
                                return iife221_(pat_let_tv61_.f24)
                            return iife220_(_pat_let62_0)
                        return iife219_(pat_let_tv60_)
                    return iife218_(_pat_let61_0)
                rhs510_ = (self.f24 if d_3220_cf22_ else p1)
                rhs511_ = (default__.fm0((0) - (p3), p1, d_3116_v0_, p3, globalState)) - (p3)
                rhs512_ = d_3234_v79_
                rhs513_ = iife217_(d_3235_v80_)
                rhs514_ = d_3236_v81_
                lhs457_ = d_3222_v68_
                lhs458_ = default__.safeIndex(143, (d_3222_v68_).length(0))
                lhs459_ = d_3233_v78_
                lhs460_ = default__.safeIndex(53, (d_3233_v78_).length(0))
                lhs457_[lhs458_] = rhs510_
                lhs459_[lhs460_] = rhs511_
                d_3234_v79_ = rhs512_
                d_3235_v80_ = rhs513_
                d_3236_v81_ = rhs514_
                d_3241_v82_: _dafny.Map
                d_3241_v82_ = _dafny.Map({(d_3233_v78_)[default__.safeIndex(53, (d_3233_v78_).length(0))]: (d_3233_v78_)[default__.safeIndex(53, (d_3233_v78_).length(0))]})
                d_3242_v83_: C8
                nw551_ = C8()
                nw551_.ctor__(((d_3116_v0_)[(d_3233_v78_)[default__.safeIndex(53, (d_3233_v78_).length(0))]] if ((d_3233_v78_)[default__.safeIndex(53, (d_3233_v78_).length(0))]) in (d_3116_v0_) else d_3227_v72_.f24), True)
                d_3242_v83_ = nw551_
                d_3243_v84_: _dafny.MultiSet
                d_3243_v84_ = _dafny.MultiSet([d_3242_v83_])
                d_3244_v85_: _dafny.MultiSet
                d_3244_v85_ = _dafny.MultiSet([d_3243_v84_, _dafny.MultiSet([d_3242_v83_, d_3242_v83_, d_3242_v83_, d_3242_v83_]), d_3243_v84_, d_3243_v84_, d_3243_v84_])
                d_3241_v82_ = (d_3241_v82_).set(default__.safeModuloInt(149, p3), (0) - ((d_3244_v85_).cardinality))
                (d_3236_v81_).f29 = d_3226_v71_
                d_3245_v87_: _dafny.Map
                d_3245_v87_ = _dafny.Map({d_3220_cf22_: d_3118_v2_})
                d_3246_v88_: _dafny.Map
                d_3246_v88_ = _dafny.Map({len(d_3215_v65_): d_3120_v4_})
                d_3247_v89_: _dafny.Map
                def iife225_():
                    coll95_ = _dafny.Map()
                    compr_95_: int
                    for compr_95_ in (((d_3245_v87_)[True] if (True) in (d_3245_v87_) else _dafny.SeqWithoutIsStrInference([(d_3233_v78_)[default__.safeIndex(53, (d_3233_v78_).length(0))] for d_3248_i11_ in range(default__.abs(286))]))).Elements:
                        d_3249_v86_: int = compr_95_
                        if (d_3249_v86_) in (((d_3245_v87_)[True] if (True) in (d_3245_v87_) else _dafny.SeqWithoutIsStrInference([(d_3233_v78_)[default__.safeIndex(53, (d_3233_v78_).length(0))] for d_3248_i11_ in range(default__.abs(286))]))):
                            coll95_[default__.safeModuloInt(d_3249_v86_, p3)] = p3
                    return _dafny.Map(coll95_)
                d_3247_v89_ = _dafny.Map({(iife225_()
                ) | (d_3241_v82_): d_3246_v88_})
                d_3247_v89_ = (d_3247_v89_).set(d_3241_v82_, d_3246_v88_)
            elif True:
                d_3250___mcc_h12_ = source52_.cf20
                d_3251_cf20_ = d_3250___mcc_h12_
                d_3252_v90_: _dafny.Map
                d_3252_v90_ = _dafny.Map({(d_3211_v61_).fm7(_dafny.Map({p1: p3}), globalState): p1})
                d_3253_v91_: _dafny.Array
                def lambda155_(d_3254_i14_):
                    return (d_3254_i14_) * (151)

                init92_ = lambda155_
                nw552_ = _dafny.Array(None, 12)
                for i0_92_ in range(nw552_.length(0)):
                    nw552_[i0_92_] = init92_(i0_92_)
                d_3253_v91_ = nw552_
                rhs515_ = ((_dafny.SeqWithoutIsStrInference([_dafny.MultiSet(d_3120_v4_) for d_3255_i12_ in range(default__.abs(334))])).set(default__.safeIndex(len(d_3252_v90_), len(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet(d_3120_v4_) for d_3256_i12_ in range(default__.abs(334))]))), d_3117_v1_)) <= (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f26])) for d_3257_i13_ in range(default__.abs(860))]))
                rhs516_ = d_3253_v91_
                lhs461_ = globalState
                lhs462_ = globalState
                lhs461_.f10 = rhs515_
                lhs462_.f7 = rhs516_
                (self).f24 = (d_3118_v2_) <= (_dafny.SeqWithoutIsStrInference([p3]))
                (self).f24 = (self).f25
                (globalState).f10 = ((p3 if False else p3)) > (p3)
            d_3258_v92_: C7
            nw553_ = C7()
            nw553_.ctor__((self).f26)
            d_3258_v92_ = nw553_
            d_3259_v93_: _dafny.Map
            d_3259_v93_ = _dafny.Map({(self).f26: d_3258_v92_})
            d_3260_v94_: _dafny.MultiSet
            d_3260_v94_ = _dafny.MultiSet([d_3259_v93_])
            d_3261_v95_: _dafny.Map
            d_3261_v95_ = _dafny.Map({(self).f25: _dafny.SeqWithoutIsStrInference([D0_DC0((self).f25, p3, p3) for d_3262_i15_ in range(default__.abs(112))])})
            d_3263_v96_: _dafny.Map
            d_3263_v96_ = _dafny.Map({(d_3211_v61_).fm5(d_3261_v95_, p1, (self).f26, globalState): p3})
            rhs517_ = d_3213_v63_
            rhs518_ = ((d_3260_v94_)[d_3259_v93_] if (d_3259_v93_) in (d_3260_v94_) else len(d_3263_v96_))
            lhs463_ = globalState
            d_3213_v63_ = rhs517_
            lhs463_.f22 = rhs518_
            if not (p1) or ((self).f26):
                d_3264_v98_: D11
                def iife226_():
                    coll96_ = _dafny.Map()
                    compr_96_: int
                    for compr_96_ in _dafny.IntegerRange(958, 442):
                        d_3265_v97_: int = compr_96_
                        if ((958) <= (d_3265_v97_)) and ((d_3265_v97_) < (442)):
                            coll96_[(d_3265_v97_) * (431)] = (self).f25
                    return _dafny.Map(coll96_)
                d_3264_v98_ = D11_DC26(p1, False, len(iife226_()
), self.f24)
                d_3266_v99_: T1
                nw554_ = C6()
                nw554_.ctor__(p1, (self).f25, p1, self.f27, (d_3264_v98_).cf50)
                d_3266_v99_ = nw554_
                d_3267_v100_: _dafny.Seq
                d_3267_v100_ = _dafny.SeqWithoutIsStrInference([d_3266_v99_])
                d_3268_v102_: _dafny.Set
                d_3268_v102_ = _dafny.Set({p1})
                d_3269_v103_: _dafny.Array
                def lambda156_(d_3270_i16_):
                    return (self).f26

                init93_ = lambda156_
                nw555_ = _dafny.Array(None, 28)
                for i0_93_ in range(nw555_.length(0)):
                    nw555_[i0_93_] = init93_(i0_93_)
                d_3269_v103_ = nw555_
                d_3271_v104_: _dafny.Seq
                d_3271_v104_ = _dafny.SeqWithoutIsStrInference([d_3269_v103_])
                d_3272_v105_: D6
                d_3272_v105_ = D6_DC13((d_3271_v104_)[default__.safeIndex(p3, len(d_3271_v104_))], p3, d_3266_v99_.f24)
                d_3273_v106_: _dafny.Set
                d_3273_v106_ = _dafny.Set({p3, len(d_3268_v102_), (d_3272_v105_).cf25, p3, p3})
                def iife227_():
                    coll97_ = _dafny.Set()
                    compr_97_: int
                    for compr_97_ in _dafny.IntegerRange(590, 521):
                        d_3274_v101_: int = compr_97_
                        if ((590) <= (d_3274_v101_)) and ((d_3274_v101_) < (521)):
                            coll97_ = coll97_.union(_dafny.Set([default__.safeModuloInt(d_3274_v101_, p3)]))
                    return _dafny.Set(coll97_)
                (globalState).f16 = (len((D22_DC50(d_3267_v100_)).cf78)) + (len((iife227_()
                ) | (d_3273_v106_)))
                d_3275_v107_: D11
                d_3275_v107_ = D11_DC27()
                d_3276_v108_: D15
                d_3276_v108_ = D15_DC33(p3)
                d_3277_v109_: _dafny.Seq
                d_3277_v109_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "patlt"))
                d_3275_v107_ = default__.fm66(d_3276_v108_, d_3277_v109_, globalState)
                (globalState).f16 = (0) - ((d_3258_v92_).fm5(d_3261_v95_, (p3) > (p3), (p1) and ((self).f26), globalState))
                d_3278_v110_: D21
                d_3278_v110_ = D21_DC49(p3, (0) - (p3), False, p1, 689)
                (globalState).f10 = ((d_3278_v110_).cf76) == ((d_3266_v99_.f24 if self.f24 else d_3266_v99_.f24))
                (globalState).f2 = d_3269_v103_
            elif True:
                d_3279_v111_: _dafny.Seq
                d_3279_v111_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ipsci"))
                d_3279_v111_ = d_3279_v111_
                d_3280_v112_: _dafny.Array
                def lambda157_(d_3281_v65_):
                    def lambda158_(d_3282_i17_):
                        return d_3281_v65_

                    return lambda158_

                init94_ = lambda157_(d_3215_v65_)
                nw556_ = _dafny.Array(None, 18)
                for i0_94_ in range(nw556_.length(0)):
                    nw556_[i0_94_] = init94_(i0_94_)
                d_3280_v112_ = nw556_
                d_3283_v113_: _dafny.Set
                out43_: _dafny.Set
                out43_ = (d_3258_v92_).m6(len(p0), d_3280_v112_, globalState)
                d_3283_v113_ = out43_
                d_3284_v114_: _dafny.Set
                d_3284_v114_ = _dafny.Set({p3, 154, len(d_3118_v2_)})
                d_3285_v115_: _dafny.Array
                nw557_ = _dafny.Array(int(0), 23)
                d_3285_v115_ = nw557_
                d_3286_v116_: _dafny.Map
                d_3286_v116_ = _dafny.Map({len(d_3284_v114_): _dafny.Set({d_3285_v115_, d_3285_v115_, d_3285_v115_})})
                d_3287_v117_: _dafny.Set
                d_3287_v117_ = _dafny.Set({d_3285_v115_, d_3285_v115_, d_3285_v115_, d_3285_v115_, d_3285_v115_})
                d_3286_v116_ = (d_3286_v116_).set(p3, d_3287_v117_)
                (self).f24 = not((self).fm8(p1, globalState))
                (globalState).f22 = -517
            d_3288_v118_: _dafny.Array
            nw558_ = _dafny.Array(None, 25)
            nw558_[int(0)] = d_3120_v4_
            nw558_[int(1)] = d_3120_v4_
            nw558_[int(2)] = d_3120_v4_
            nw558_[int(3)] = d_3120_v4_
            nw558_[int(4)] = d_3120_v4_
            nw558_[int(5)] = d_3120_v4_
            nw558_[int(6)] = d_3120_v4_
            nw558_[int(7)] = d_3120_v4_
            nw558_[int(8)] = d_3120_v4_
            nw558_[int(9)] = d_3120_v4_
            nw558_[int(10)] = d_3120_v4_
            nw558_[int(11)] = _dafny.SeqWithoutIsStrInference([((d_3116_v0_)[p3] if (p3) in (d_3116_v0_) else (self).f26), (self).f25])
            nw558_[int(12)] = d_3120_v4_
            nw558_[int(13)] = d_3120_v4_
            nw558_[int(14)] = d_3120_v4_
            nw558_[int(15)] = (d_3120_v4_).set(default__.safeIndex(988, len(d_3120_v4_)), (self).f26)
            nw558_[int(16)] = d_3120_v4_
            nw558_[int(17)] = d_3120_v4_
            nw558_[int(18)] = d_3120_v4_
            nw558_[int(19)] = _dafny.SeqWithoutIsStrInference([(self).f25, True])
            nw558_[int(20)] = d_3120_v4_
            nw558_[int(21)] = default__.fm9((d_3215_v65_).set((self).f26, p3), p3, default__.fm21(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t")), False, p3, len(d_3118_v2_), globalState), p3, globalState)
            nw558_[int(22)] = d_3120_v4_
            nw558_[int(23)] = d_3120_v4_
            nw558_[int(24)] = _dafny.SeqWithoutIsStrInference([p1])
            d_3288_v118_ = nw558_
            d_3289_v119_: D19
            d_3289_v119_ = D19_DC42(p1, d_3288_v118_)
            def iife228_(_pat_let65_0):
                def iife229_(d_3290_dt__update__tmp_h1_):
                    def iife230_(_pat_let66_0):
                        def iife231_(d_3291_dt__update_hcf63_h0_):
                            return D19_DC42(d_3291_dt__update_hcf63_h0_, (d_3290_dt__update__tmp_h1_).cf64)
                        return iife231_(_pat_let66_0)
                    return iife230_(not((self).f26))
                return iife229_(_pat_let65_0)
            (globalState).f10 = (iife228_(d_3289_v119_)) in (_dafny.SeqWithoutIsStrInference([d_3289_v119_, d_3289_v119_, d_3289_v119_]))
        elif True:
            d_3292_v120_: C7
            nw559_ = C7()
            nw559_.ctor__(self.f24)
            d_3292_v120_ = nw559_
            d_3293_v121_: _dafny.Map
            d_3293_v121_ = _dafny.Map({(self).f26: d_3292_v120_})
            d_3294_v122_: _dafny.Array
            nw560_ = _dafny.Array(None, 15)
            nw560_[int(0)] = d_3292_v120_
            nw560_[int(1)] = d_3292_v120_
            nw560_[int(2)] = d_3292_v120_
            nw560_[int(3)] = d_3292_v120_
            nw560_[int(4)] = d_3292_v120_
            nw560_[int(5)] = d_3292_v120_
            nw560_[int(6)] = d_3292_v120_
            nw560_[int(7)] = d_3292_v120_
            nw560_[int(8)] = d_3292_v120_
            nw560_[int(9)] = d_3292_v120_
            nw560_[int(10)] = ((d_3293_v121_)[p1] if (p1) in (d_3293_v121_) else d_3292_v120_)
            nw560_[int(11)] = d_3292_v120_
            nw560_[int(12)] = d_3292_v120_
            nw560_[int(13)] = d_3292_v120_
            nw560_[int(14)] = d_3292_v120_
            d_3294_v122_ = nw560_
            index499_ = default__.safeIndex(564, (d_3294_v122_).length(0))
            (d_3294_v122_)[index499_] = d_3292_v120_
            index500_ = default__.safeIndex(564, (d_3294_v122_).length(0))
            (d_3294_v122_)[index500_] = (d_3292_v120_ if self.f24 else d_3292_v120_)
            d_3295_v123_: _dafny.Array
            nw561_ = _dafny.Array(None, 6)
            nw561_[int(0)] = p3
            nw561_[int(1)] = 64
            nw561_[int(2)] = (0) - (p3)
            nw561_[int(3)] = p3
            nw561_[int(4)] = p3
            nw561_[int(5)] = p3
            d_3295_v123_ = nw561_
            index501_ = default__.safeIndex(724, (d_3295_v123_).length(0))
            (d_3295_v123_)[index501_] = p3
            d_3296_v124_: T3
            nw562_ = C5()
            nw562_.ctor__((self).f25, self.f27, p1, False)
            d_3296_v124_ = nw562_
            d_3297_v125_: _dafny.Set
            d_3297_v125_ = _dafny.Set({d_3296_v124_})
            d_3298_v126_: _dafny.Seq
            d_3298_v126_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "olfbm"))
            d_3299_v127_: T0
            nw563_ = C2()
            nw563_.ctor__(d_3298_v126_, d_3298_v126_)
            d_3299_v127_ = nw563_
            d_3300_v128_: _dafny.Map
            d_3300_v128_ = _dafny.Map({d_3295_v123_: d_3299_v127_})
            d_3301_v129_: _dafny.Map
            d_3301_v129_ = _dafny.Map({((d_3300_v128_)[d_3295_v123_] if (d_3295_v123_) in (d_3300_v128_) else d_3299_v127_): (self).f26})
            d_3302_v130_: _dafny.Set
            d_3302_v130_ = _dafny.Set({d_3296_v124_.f27})
            d_3303_v131_: D0
            d_3303_v131_ = D0_DC0((self).f25, len(d_3302_v130_), p3)
            d_3304_v132_: _dafny.Array
            nw564_ = _dafny.Array(None, 12)
            nw564_[int(0)] = d_3118_v2_
            nw564_[int(1)] = d_3118_v2_
            nw564_[int(2)] = d_3118_v2_
            nw564_[int(3)] = d_3118_v2_
            nw564_[int(4)] = default__.fm42(globalState)
            nw564_[int(5)] = (d_3118_v2_ if self.f24 else (d_3118_v2_).set(default__.safeIndex(p3, len(d_3118_v2_)), p3))
            nw564_[int(6)] = _dafny.SeqWithoutIsStrInference([535 for d_3305_i18_ in range(default__.abs(278))])
            nw564_[int(7)] = d_3118_v2_
            nw564_[int(8)] = (d_3118_v2_) + (d_3118_v2_)
            nw564_[int(9)] = (d_3118_v2_).set(default__.safeIndex(p3, len(d_3118_v2_)), p3)
            nw564_[int(10)] = _dafny.SeqWithoutIsStrInference([len(d_3297_v125_), ((d_3119_v3_)[len((d_3301_v129_).set(d_3299_v127_, p1))] if (len((d_3301_v129_).set(d_3299_v127_, p1))) in (d_3119_v3_) else p3), p3, (self).fm6(d_3303_v131_, globalState)])
            nw564_[int(11)] = _dafny.SeqWithoutIsStrInference([p3, p3])
            d_3304_v132_ = nw564_
            index502_ = default__.safeIndex(482, (d_3304_v132_).length(0))
            (d_3304_v132_)[index502_] = d_3118_v2_
            index503_ = default__.safeIndex(724, (d_3295_v123_).length(0))
            index504_ = default__.safeIndex(482, (d_3304_v132_).length(0))
            rhs519_ = self.f24
            rhs520_ = ((d_3118_v2_)[default__.safeIndex(867, len(d_3118_v2_))]) + (default__.safeDivisionInt(len(_dafny.Map({len(d_3118_v2_): p3})), len((_dafny.SeqWithoutIsStrInference([self.f27 for d_3306_i19_ in range(default__.abs(468))])).set(default__.safeIndex(p3, len(_dafny.SeqWithoutIsStrInference([self.f27 for d_3307_i19_ in range(default__.abs(468))]))), default__.fm47(globalState)))))
            rhs521_ = (d_3118_v2_) + (d_3118_v2_)
            rhs522_ = ((d_3298_v126_).set(default__.safeIndex(len(d_3298_v126_), len(d_3298_v126_)), d_3296_v124_.f27)).set(default__.safeIndex(p3, len((d_3298_v126_).set(default__.safeIndex(len(d_3298_v126_), len(d_3298_v126_)), d_3296_v124_.f27))), d_3296_v124_.f27)
            rhs523_ = ((d_3296_v124_).f26 if (self).f26 else (d_3118_v2_) < (d_3118_v2_))
            lhs464_ = globalState
            lhs465_ = d_3295_v123_
            lhs466_ = default__.safeIndex(724, (d_3295_v123_).length(0))
            lhs467_ = d_3304_v132_
            lhs468_ = default__.safeIndex(482, (d_3304_v132_).length(0))
            lhs469_ = globalState
            lhs464_.f10 = rhs519_
            lhs465_[lhs466_] = rhs520_
            lhs467_[lhs468_] = rhs521_
            d_3298_v126_ = rhs522_
            lhs469_.f10 = rhs523_
            d_3308_v133_: _dafny.Set
            d_3308_v133_ = _dafny.Set({(d_3296_v124_).f26})
            d_3309_v134_: C6
            nw565_ = C6()
            nw565_.ctor__((d_3308_v133_).isdisjoint(d_3308_v133_), p1, (_dafny.MultiSet([(d_3296_v124_).f26, (self).f26, (self).f25])).isdisjoint(d_3117_v1_), self.f27, True)
            d_3309_v134_ = nw565_
            index505_ = default__.safeIndex(564, (d_3294_v122_).length(0))
            (d_3294_v122_)[index505_] = d_3292_v120_
            d_3310_v135_: D11
            d_3310_v135_ = D11_DC28()
            d_3310_v135_ = d_3310_v135_
        d_3311_i20_: int
        d_3311_i20_ = 0
        with _dafny.label("23"):
            while (self).f25:
                with _dafny.c_label("23"):
                    if (d_3311_i20_) >= (100):
                        raise _dafny.Break("23")
                    d_3311_i20_ = (d_3311_i20_) + (1)
                    d_3312_v136_: _dafny.Map
                    d_3312_v136_ = _dafny.Map({p1: (0) - (p3)})
                    d_3313_v137_: _dafny.Map
                    d_3313_v137_ = _dafny.Map({p3: ((d_3312_v136_)[True] if (True) in (d_3312_v136_) else p3)})
                    d_3314_v138_: D2
                    d_3314_v138_ = D2_DC3(d_3313_v137_)
                    d_3314_v138_ = d_3314_v138_
                    d_3315_v139_: _dafny.Set
                    d_3315_v139_ = _dafny.Set({(self).f26})
                    d_3316_v140_: D23
                    d_3316_v140_ = D23_DC54(_dafny.Set({p1, (self).f26}))
                    d_3317_v141_: _dafny.Seq
                    d_3317_v141_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "juv"))
                    d_3318_v142_: D2
                    d_3318_v142_ = D2_DC5(self.f27, p1, p3, len(d_3317_v141_))
                    d_3319_v143_: D2
                    d_3319_v143_ = D2_DC6(d_3318_v142_)
                    d_3320_v144_: _dafny.Seq
                    d_3320_v144_ = _dafny.SeqWithoutIsStrInference([d_3315_v139_])
                    d_3321_v145_: _dafny.Array
                    nw566_ = _dafny.Array(None, 18)
                    nw566_[int(0)] = d_3315_v139_
                    nw566_[int(1)] = default__.fm44(p3, d_3118_v2_, globalState)
                    nw566_[int(2)] = d_3315_v139_
                    nw566_[int(3)] = _dafny.Set({False})
                    nw566_[int(4)] = d_3315_v139_
                    nw566_[int(5)] = d_3315_v139_
                    nw566_[int(6)] = d_3315_v139_
                    nw566_[int(7)] = (d_3315_v139_).intersection(d_3315_v139_)
                    nw566_[int(8)] = d_3315_v139_
                    nw566_[int(9)] = d_3315_v139_
                    nw566_[int(10)] = d_3315_v139_
                    nw566_[int(11)] = (d_3315_v139_).intersection(_dafny.Set({(self).f25, p1}))
                    nw566_[int(12)] = ((d_3316_v140_).cf85) | (_dafny.Set({self.f24, (D5_DC11(d_3319_v143_, p1)).cf22}))
                    nw566_[int(13)] = (d_3315_v139_) - (d_3315_v139_)
                    nw566_[int(14)] = d_3315_v139_
                    nw566_[int(15)] = d_3315_v139_
                    nw566_[int(16)] = (d_3320_v144_)[default__.safeIndex(p3, len(d_3320_v144_))]
                    nw566_[int(17)] = _dafny.Set({(self).f26})
                    d_3321_v145_ = nw566_
                    d_3321_v145_ = (d_3321_v145_ if p1 else d_3321_v145_)
                    if p1:
                        d_3322_v146_: _dafny.Map
                        d_3322_v146_ = _dafny.Map({self.f24: (self).f26})
                        d_3323_v147_: D15
                        d_3323_v147_ = D15_DC32((d_3322_v146_ if (self).f26 else d_3322_v146_))
                        rhs524_ = (d_3120_v4_) == (d_3120_v4_)
                        rhs525_ = (p3) != (p3)
                        rhs526_ = d_3323_v147_
                        lhs470_ = globalState
                        lhs471_ = globalState
                        lhs470_.f10 = rhs524_
                        lhs471_.f10 = rhs525_
                        d_3323_v147_ = rhs526_
                        d_3324_v149_: _dafny.Array
                        nw567_ = _dafny.Array(False, 2)
                        d_3324_v149_ = nw567_
                        d_3325_v150_: D6
                        d_3325_v150_ = D6_DC13(d_3324_v149_, 834, False)
                        d_3326_v151_: D0
                        d_3326_v151_ = D0_DC0((self).f25, (d_3325_v150_).cf25, p3)
                        d_3327_v152_: _dafny.Set
                        def iife232_():
                            coll98_ = _dafny.Map()
                            compr_98_: int
                            for compr_98_ in _dafny.IntegerRange(543, -527):
                                d_3328_v148_: int = compr_98_
                                if ((543) <= (d_3328_v148_)) and ((d_3328_v148_) < (-527)):
                                    coll98_[default__.safeDivisionInt(d_3328_v148_, p3)] = (self).f25
                            return _dafny.Map(coll98_)
                        d_3327_v152_ = _dafny.Set({len(iife232_()
                        ), default__.fm0(default__.fm0(p3, (self).f25, d_3116_v0_, p3, globalState), p1, _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rqabx"))): self.f24}), p3, globalState), p3, ((d_3117_v1_).intersection(d_3117_v1_)).cardinality, (d_3211_v61_).fm6(d_3326_v151_, globalState)})
                        def iife233_():
                            coll99_ = _dafny.Set()
                            compr_99_: int
                            for compr_99_ in (d_3118_v2_).Elements:
                                d_3329_v153_: int = compr_99_
                                if (d_3329_v153_) in (d_3118_v2_):
                                    coll99_ = coll99_.union(_dafny.Set([default__.safeModuloInt(d_3329_v153_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lplftua"))))]))
                            return _dafny.Set(coll99_)
                        d_3327_v152_ = ((d_3327_v152_).intersection(_dafny.Set({p3}))).intersection(iife233_()
                        )
                        d_3330_v154_: _dafny.Seq
                        d_3330_v154_ = _dafny.SeqWithoutIsStrInference([d_3317_v141_])
                        d_3331_v155_: _dafny.Array
                        nw568_ = _dafny.Array(None, 24)
                        nw568_[int(0)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sdf"))
                        nw568_[int(1)] = d_3317_v141_
                        nw568_[int(2)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j")) if p1 else d_3317_v141_)
                        nw568_[int(3)] = (d_3330_v154_)[default__.safeIndex(p3, len(d_3330_v154_))]
                        nw568_[int(4)] = (_dafny.SeqWithoutIsStrInference([self.f27 for d_3332_i21_ in range(default__.abs(324))])) + (default__.fm29(globalState))
                        nw568_[int(5)] = (d_3317_v141_) + ((d_3330_v154_)[default__.safeIndex(-180, len(d_3330_v154_))])
                        nw568_[int(6)] = _dafny.SeqWithoutIsStrInference([self.f27 for d_3333_i22_ in range(default__.abs(531))])
                        nw568_[int(7)] = default__.fm26(p1, (0) - (p3), p3, globalState)
                        nw568_[int(8)] = _dafny.SeqWithoutIsStrInference([self.f27 for d_3334_i23_ in range(default__.abs(892))])
                        nw568_[int(9)] = d_3317_v141_
                        nw568_[int(10)] = _dafny.SeqWithoutIsStrInference([self.f27 for d_3335_i24_ in range(default__.abs(0))])
                        nw568_[int(11)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a"))) + ((d_3317_v141_).set(default__.safeIndex(p3, len(d_3317_v141_)), self.f27))
                        nw568_[int(12)] = d_3317_v141_
                        nw568_[int(13)] = _dafny.SeqWithoutIsStrInference([self.f27 for d_3336_i25_ in range(default__.abs(15))])
                        nw568_[int(14)] = d_3317_v141_
                        nw568_[int(15)] = d_3317_v141_
                        nw568_[int(16)] = (d_3317_v141_ if (self).f25 else d_3317_v141_)
                        nw568_[int(17)] = d_3317_v141_
                        nw568_[int(18)] = d_3317_v141_
                        nw568_[int(19)] = (d_3330_v154_)[default__.safeIndex((self).fm5(default__.fm55(_dafny.CodePoint('x'), p1, self.f24, (self).f26, globalState), p1, p1, globalState), len(d_3330_v154_))]
                        nw568_[int(20)] = d_3317_v141_
                        nw568_[int(21)] = d_3317_v141_
                        nw568_[int(22)] = d_3317_v141_
                        nw568_[int(23)] = (d_3317_v141_) + (d_3317_v141_)
                        d_3331_v155_ = nw568_
                        index506_ = default__.safeIndex(101, (d_3331_v155_).length(0))
                        (d_3331_v155_)[index506_] = d_3317_v141_
                        index507_ = default__.safeIndex(101, (d_3331_v155_).length(0))
                        rhs527_ = _dafny.CodePoint('j')
                        rhs528_ = d_3323_v147_
                        rhs529_ = (self).f25
                        rhs530_ = len((d_3118_v2_) + ((d_3118_v2_) + (d_3118_v2_)))
                        rhs531_ = d_3317_v141_
                        lhs472_ = self
                        lhs473_ = globalState
                        lhs474_ = globalState
                        lhs475_ = d_3331_v155_
                        lhs476_ = default__.safeIndex(101, (d_3331_v155_).length(0))
                        lhs472_.f27 = rhs527_
                        d_3323_v147_ = rhs528_
                        lhs473_.f10 = rhs529_
                        lhs474_.f16 = rhs530_
                        lhs475_[lhs476_] = rhs531_
                        d_3116_v0_ = (d_3116_v0_).set(p3, (self).f25)
                        d_3337_v156_: D7
                        d_3337_v156_ = D7_DC15(self.f27, ((0) - (p3)) > (549))
                        d_3337_v156_ = d_3337_v156_
                    elif True:
                        d_3338_v157_: C0
                        nw569_ = C0()
                        nw569_.ctor__(d_3317_v141_)
                        d_3338_v157_ = nw569_
                        d_3339_v158_: _dafny.Array
                        nw570_ = _dafny.Array(_dafny.Seq({}), 11)
                        d_3339_v158_ = nw570_
                        index508_ = default__.safeIndex(80, (d_3339_v158_).length(0))
                        (d_3339_v158_)[index508_] = (d_3120_v4_) + (_dafny.SeqWithoutIsStrInference([(self).f26, self.f24]))
                        index509_ = default__.safeIndex(80, (d_3339_v158_).length(0))
                        (d_3339_v158_)[index509_] = (d_3120_v4_).set(default__.safeIndex(default__.safeDivisionInt(p3, 118), len(d_3120_v4_)), (self).f26)
                        d_3340_v159_: _dafny.Array
                        def lambda159_(d_3341_v141_):
                            def lambda160_(d_3342_i26_):
                                return (d_3342_i26_) - (len(d_3341_v141_))

                            return lambda160_

                        init95_ = lambda159_(d_3317_v141_)
                        nw571_ = _dafny.Array(None, 5)
                        for i0_95_ in range(nw571_.length(0)):
                            nw571_[i0_95_] = init95_(i0_95_)
                        d_3340_v159_ = nw571_
                        d_3343_v160_: C8
                        nw572_ = C8()
                        nw572_.ctor__(True, (self).f25)
                        d_3343_v160_ = nw572_
                        d_3344_v161_: _dafny.Seq
                        d_3344_v161_ = _dafny.SeqWithoutIsStrInference([d_3343_v160_])
                        d_3345_v162_: _dafny.MultiSet
                        d_3345_v162_ = _dafny.MultiSet([(d_3344_v161_)[default__.safeIndex(p3, len(d_3344_v161_))]])
                        rhs532_ = not(not((p3) not in ((d_3118_v2_).set(default__.safeIndex(p3, len(d_3118_v2_)), p3))))
                        rhs533_ = d_3340_v159_
                        rhs534_ = (self.f24) not in (_dafny.MultiSet((d_3339_v158_)[default__.safeIndex(80, (d_3339_v158_).length(0))]))
                        rhs535_ = (d_3338_v157_).fm23(p1, self.f27, (p3) != (p3), len(d_3315_v139_), globalState)
                        rhs536_ = ((_dafny.MultiSet([d_3343_v160_, d_3343_v160_, d_3343_v160_, d_3343_v160_])).issubset(d_3345_v162_)) and ((self).f25)
                        lhs477_ = globalState
                        lhs478_ = globalState
                        lhs479_ = self
                        lhs480_ = globalState
                        lhs481_ = globalState
                        lhs477_.f10 = rhs532_
                        lhs478_.f7 = rhs533_
                        lhs479_.f24 = rhs534_
                        lhs480_.f22 = rhs535_
                        lhs481_.f10 = rhs536_
                        (globalState).f22 = p3
                        d_3346_v163_: _dafny.Array
                        nw573_ = _dafny.Array(False, 27)
                        d_3346_v163_ = nw573_
                        index510_ = default__.safeIndex(980, (d_3346_v163_).length(0))
                        (d_3346_v163_)[index510_] = (d_3120_v4_)[default__.safeIndex(p3, len(d_3120_v4_))]
                        index511_ = default__.safeIndex(980, (d_3346_v163_).length(0))
                        (d_3346_v163_)[index511_] = self.f24
                    (self).f27 = default__.fm21(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hp")), p1, p3, (p3) * (895), globalState)
                    pass
            pass
        if (self).f25:
            d_3347_v164_: D7
            d_3347_v164_ = D7_DC16(p1, p3, (self).f26, True, (self).f26)
            d_3348_v165_: _dafny.Array
            nw574_ = _dafny.Array(int(0), 4)
            d_3348_v165_ = nw574_
            d_3349_v166_: _dafny.Map
            d_3349_v166_ = _dafny.Map({p3: d_3348_v165_})
            d_3350_v167_: D24
            d_3350_v167_ = D24_DC56(d_3349_v166_)
            pat_let_tv64_ = d_3350_v167_
            d_3351_v168_: _dafny.Map
            def iife234_(_pat_let67_0):
                def iife235_(d_3352_dt__update__tmp_h2_):
                    def iife236_(_pat_let68_0):
                        def iife237_(d_3353_dt__update_hcf30_h0_):
                            def iife238_(_pat_let69_0):
                                def iife239_(d_3354_dt__update_hcf33_h1_):
                                    def iife240_(_pat_let70_0):
                                        def iife241_(d_3355_dt__update_hcf31_h1_):
                                            return D7_DC16(d_3353_dt__update_hcf30_h0_, d_3355_dt__update_hcf31_h1_, (d_3352_dt__update__tmp_h2_).cf32, d_3354_dt__update_hcf33_h1_, (d_3352_dt__update__tmp_h2_).cf34)
                                        return iife241_(_pat_let70_0)
                                    return iife240_(len((pat_let_tv64_).cf89))
                                return iife239_(_pat_let69_0)
                            return iife238_(self.f24)
                        return iife237_(_pat_let68_0)
                    return iife236_((self).f26)
                return iife235_(_pat_let67_0)
            d_3351_v168_ = _dafny.Map({self.f24: iife234_(d_3347_v164_)})
            d_3351_v168_ = (d_3351_v168_).set((self).f25, D7_DC16((self).f25, 162, (self).f25, (self).f25, self.f24))
            d_3356_v169_: D3
            d_3356_v169_ = D3_DC7(d_3120_v4_)
            d_3357_v170_: _dafny.Map
            d_3357_v170_ = _dafny.Map({(self).f25: len((d_3356_v169_).cf15)})
            d_3358_v171_: _dafny.MultiSet
            d_3358_v171_ = _dafny.MultiSet([d_3357_v170_])
            d_3359_v172_: D5
            d_3359_v172_ = D5_DC10(d_3358_v171_)
            d_3360_v173_: _dafny.Map
            d_3360_v173_ = _dafny.Map({p3: d_3359_v172_})
            def iife242_():
                coll100_ = _dafny.Map()
                compr_100_: int
                for compr_100_ in _dafny.IntegerRange(170, 47):
                    d_3361_v174_: int = compr_100_
                    if ((170) <= (d_3361_v174_)) and ((d_3361_v174_) < (47)):
                        coll100_[default__.safeModuloInt(d_3361_v174_, p3)] = d_3359_v172_
                return _dafny.Map(coll100_)
            d_3360_v173_ = iife242_()
            
            d_3362_v175_: _dafny.Map
            d_3362_v175_ = default__.fm67(p3, d_3118_v2_, p3, globalState)
            d_3363_v176_: _dafny.Set
            d_3363_v176_ = _dafny.Set({d_3362_v175_})
            d_3364_v177_: _dafny.MultiSet
            d_3364_v177_ = _dafny.MultiSet([d_3362_v175_])
            def iife243_():
                coll101_ = _dafny.Set()
                compr_101_: _dafny.Map
                for compr_101_ in ((d_3364_v177_ if p1 else d_3364_v177_)).Elements:
                    d_3365_v178_: _dafny.Map = compr_101_
                    if (d_3365_v178_) in ((d_3364_v177_ if p1 else d_3364_v177_)):
                        coll101_ = coll101_.union(_dafny.Set([d_3365_v178_]))
                return _dafny.Set(coll101_)
            d_3363_v176_ = iife243_()
            
            if True:
                (globalState).f11 = (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hyfpojit")))) - (p3)
                d_3366_v179_: _dafny.Array
                nw575_ = _dafny.Array(_dafny.Map({}), 15)
                d_3366_v179_ = nw575_
                d_3366_v179_ = d_3366_v179_
                (self).f24 = default__.fm1(not (((d_3116_v0_)[len(_dafny.SeqWithoutIsStrInference([813 for d_3367_i27_ in range(default__.abs(367))]))] if (len(_dafny.SeqWithoutIsStrInference([813 for d_3368_i27_ in range(default__.abs(367))]))) in (d_3116_v0_) else not(self.f24))) or (not((self).f25)), self.f24, p3, globalState)
                d_3369_v180_: _dafny.Seq
                d_3369_v180_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xpfnvs"))
                (globalState).f10 = (p3) == (len(d_3369_v180_))
                d_3118_v2_ = d_3118_v2_
            elif True:
                d_3370_v181_: _dafny.Map
                d_3370_v181_ = _dafny.Map({(0) - (p3): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fjsnq"))})
                rhs537_ = len(d_3370_v181_)
                rhs538_ = len(d_3357_v170_)
                lhs482_ = globalState
                lhs483_ = globalState
                lhs482_.f22 = rhs537_
                lhs483_.f16 = rhs538_
                d_3371_v182_: _dafny.Array
                nw576_ = _dafny.Array(False, 3)
                d_3371_v182_ = nw576_
                (globalState).f2 = d_3371_v182_
                (globalState).f11 = p3
                (globalState).f10 = True
                d_3372_v183_: C6
                nw577_ = C6()
                nw577_.ctor__(p1, self.f24, p1, self.f27, (self).f26)
                d_3372_v183_ = nw577_
                d_3373_v184_: _dafny.Map
                d_3373_v184_ = _dafny.Map({(self).f25: d_3372_v183_})
                d_3374_v185_: _dafny.Seq
                d_3374_v185_ = _dafny.SeqWithoutIsStrInference([d_3373_v184_])
                (globalState).f16 = len(_dafny.Map({p3: (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dq"))) != ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qjcwnh"))).set(default__.safeIndex(len(d_3374_v185_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qjcwnh")))), self.f27))}))
            (globalState).f10 = not(self.f24)
        elif True:
            d_3375_v186_: _dafny.Set
            d_3375_v186_ = _dafny.Set({(self).f26, self.f24})
            (globalState).f10 = (self.f24 if self.f24 else (len(d_3375_v186_)) == (227))
            (globalState).f16 = (0) - (p3)
            d_3376_v187_: _dafny.MultiSet
            d_3376_v187_ = _dafny.MultiSet([(d_3116_v0_) | (d_3116_v0_)])
            d_3376_v187_ = d_3376_v187_
            d_3377_v188_: _dafny.Array
            nw578_ = _dafny.Array(D21.default()(), 17)
            d_3377_v188_ = nw578_
            d_3378_v189_: D21
            d_3378_v189_ = D21_DC48()
            index512_ = default__.safeIndex(500, (d_3377_v188_).length(0))
            (d_3377_v188_)[index512_] = d_3378_v189_
            index513_ = default__.safeIndex(500, (d_3377_v188_).length(0))
            (d_3377_v188_)[index513_] = d_3378_v189_
            d_3379_v190_: _dafny.Array
            nw579_ = _dafny.Array(int(0), 16)
            d_3379_v190_ = nw579_
            index514_ = default__.safeIndex(256, (d_3379_v190_).length(0))
            (d_3379_v190_)[index514_] = (-301) - ((0) - (p3))
            index515_ = default__.safeIndex(256, (d_3379_v190_).length(0))
            (d_3379_v190_)[index515_] = default__.safeModuloInt(-196, p3)
        d_3380_i28_: int
        d_3380_i28_ = 0
        with _dafny.label("24"):
            while (p3) < (p3):
                with _dafny.c_label("24"):
                    if (d_3380_i28_) >= (100):
                        raise _dafny.Break("24")
                    d_3380_i28_ = (d_3380_i28_) + (1)
                    d_3381_v191_: C7
                    nw580_ = C7()
                    nw580_.ctor__(self.f24)
                    d_3381_v191_ = nw580_
                    d_3382_v192_: _dafny.Map
                    d_3382_v192_ = _dafny.Map({(self).f25: p3})
                    d_3383_v193_: _dafny.Set
                    d_3383_v193_ = _dafny.Set({p3, len(d_3382_v192_), p3})
                    d_3384_v194_: _dafny.Map
                    d_3384_v194_ = _dafny.Map({self.f27: p1})
                    d_3385_v195_: _dafny.Seq
                    d_3385_v195_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ght"))
                    d_3386_v196_: _dafny.Set
                    d_3386_v196_ = _dafny.Set({p1, p1, ((d_3384_v194_)[(d_3385_v195_)[default__.safeIndex(p3, len(d_3385_v195_))]] if ((d_3385_v195_)[default__.safeIndex(p3, len(d_3385_v195_))]) in (d_3384_v194_) else (self).f26)})
                    d_3387_v197_: _dafny.Seq
                    d_3387_v197_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({len(d_3386_v196_)})])
                    d_3388_v203_: _dafny.Map
                    def iife244_():
                        coll102_ = _dafny.Set()
                        compr_102_: int
                        for compr_102_ in _dafny.IntegerRange(945, 200):
                            d_3389_v202_: int = compr_102_
                            if ((945) <= (d_3389_v202_)) and ((d_3389_v202_) < (200)):
                                coll102_ = coll102_.union(_dafny.Set([(d_3389_v202_) - (p3)]))
                        return _dafny.Set(coll102_)
                    d_3388_v203_ = _dafny.Map({p3: iife244_()
                    })
                    d_3390_v204_: _dafny.Seq
                    d_3390_v204_ = _dafny.SeqWithoutIsStrInference([D0_DC0(p1, p3, len(_dafny.SeqWithoutIsStrInference([p3])))])
                    d_3391_v206_: _dafny.Map
                    d_3391_v206_ = _dafny.Map({not(self.f24): d_3390_v204_})
                    d_3392_v208_: D23
                    d_3392_v208_ = D23_DC54(_dafny.Set({(self).f26, p1, True, (self).f25}))
                    d_3393_v209_: _dafny.Map
                    d_3393_v209_ = _dafny.Map({p3: (d_3392_v208_).cf85})
                    d_3394_v210_: _dafny.Array
                    nw581_ = _dafny.Array(None, 26)
                    nw581_[int(0)] = d_3383_v193_
                    nw581_[int(1)] = d_3383_v193_
                    nw581_[int(2)] = d_3383_v193_
                    nw581_[int(3)] = (d_3383_v193_) | (d_3383_v193_)
                    nw581_[int(4)] = d_3383_v193_
                    nw581_[int(5)] = d_3383_v193_
                    nw581_[int(6)] = (d_3387_v197_)[default__.safeIndex(p3, len(d_3387_v197_))]
                    def iife245_():
                        coll103_ = _dafny.Set()
                        compr_103_: int
                        for compr_103_ in (d_3119_v3_).Elements:
                            d_3395_v198_: int = compr_103_
                            if (d_3395_v198_) in (d_3119_v3_):
                                coll103_ = coll103_.union(_dafny.Set([(d_3395_v198_) + (len(_dafny.Set({True, False})))]))
                        return _dafny.Set(coll103_)
                    nw581_[int(7)] = (iife245_()
                    ).intersection(d_3383_v193_)
                    def iife246_():
                        coll104_ = _dafny.Set()
                        compr_104_: int
                        for compr_104_ in _dafny.IntegerRange(-332, 650):
                            d_3396_v199_: int = compr_104_
                            if ((-332) <= (d_3396_v199_)) and ((d_3396_v199_) < (650)):
                                coll104_ = coll104_.union(_dafny.Set([(d_3396_v199_) + (p3)]))
                        return _dafny.Set(coll104_)
                    nw581_[int(8)] = iife246_()
                    
                    nw581_[int(9)] = d_3383_v193_
                    nw581_[int(10)] = (d_3383_v193_) | (d_3383_v193_)
                    nw581_[int(11)] = _dafny.Set({len(d_3118_v2_), 22})
                    def iife247_():
                        coll105_ = _dafny.Set()
                        compr_105_: int
                        for compr_105_ in _dafny.IntegerRange(854, -817):
                            d_3397_v200_: int = compr_105_
                            if ((854) <= (d_3397_v200_)) and ((d_3397_v200_) < (-817)):
                                coll105_ = coll105_.union(_dafny.Set([(d_3397_v200_) * (p3)]))
                        return _dafny.Set(coll105_)
                    nw581_[int(12)] = (d_3383_v193_).intersection(iife247_()
                    )
                    nw581_[int(13)] = d_3383_v193_
                    def iife248_():
                        coll106_ = _dafny.Set()
                        compr_106_: int
                        for compr_106_ in _dafny.IntegerRange(51, 242):
                            d_3398_v201_: int = compr_106_
                            if ((51) <= (d_3398_v201_)) and ((d_3398_v201_) < (242)):
                                coll106_ = coll106_.union(_dafny.Set([(d_3398_v201_) - (p3)]))
                        return _dafny.Set(coll106_)
                    nw581_[int(14)] = (d_3383_v193_).intersection(iife248_()
                    )
                    nw581_[int(15)] = d_3383_v193_
                    nw581_[int(16)] = (d_3383_v193_).intersection(((d_3388_v203_)[(d_3381_v191_).fm5((_dafny.Map({(self).f26: d_3390_v204_})).set(p1, d_3390_v204_), True, p1, globalState)] if ((d_3381_v191_).fm5((_dafny.Map({(self).f26: d_3390_v204_})).set(p1, d_3390_v204_), True, p1, globalState)) in (d_3388_v203_) else d_3383_v193_))
                    def iife249_():
                        coll107_ = _dafny.Set()
                        compr_107_: int
                        for compr_107_ in (d_3119_v3_).Elements:
                            d_3399_v205_: int = compr_107_
                            if (d_3399_v205_) in (d_3119_v3_):
                                coll107_ = coll107_.union(_dafny.Set([(d_3399_v205_) + (134)]))
                        return _dafny.Set(coll107_)
                    nw581_[int(17)] = (iife249_()
                    ).intersection(_dafny.Set({p3, p3, (self).fm5(d_3391_v206_, self.f24, (self).f26, globalState), len(d_3385_v195_)}))
                    nw581_[int(18)] = d_3383_v193_
                    nw581_[int(19)] = (d_3383_v193_ if self.f24 else d_3383_v193_)
                    nw581_[int(20)] = ((D7_DC14(d_3383_v193_)).cf27).intersection(d_3383_v193_)
                    nw581_[int(21)] = _dafny.Set({p3})
                    def iife250_():
                        coll108_ = _dafny.Set()
                        compr_108_: int
                        for compr_108_ in _dafny.IntegerRange(787, 496):
                            d_3400_v207_: int = compr_108_
                            if ((787) <= (d_3400_v207_)) and ((d_3400_v207_) < (496)):
                                coll108_ = coll108_.union(_dafny.Set([(d_3400_v207_) - (p3)]))
                        return _dafny.Set(coll108_)
                    nw581_[int(22)] = iife250_()
                    
                    nw581_[int(23)] = d_3383_v193_
                    nw581_[int(24)] = d_3383_v193_
                    nw581_[int(25)] = _dafny.Set({p3, 162, len(((d_3393_v209_)[195] if (195) in (d_3393_v209_) else d_3386_v196_))})
                    d_3394_v210_ = nw581_
                    index516_ = default__.safeIndex(430, (d_3394_v210_).length(0))
                    (d_3394_v210_)[index516_] = d_3383_v193_
                    index517_ = default__.safeIndex(430, (d_3394_v210_).length(0))
                    (d_3394_v210_)[index517_] = (_dafny.Set({p3})) | ((d_3387_v197_)[default__.safeIndex(p3, len(d_3387_v197_))])
                    d_3401_v211_: _dafny.Array
                    def lambda161_(d_3402_p3_):
                        def lambda162_(d_3403_i29_):
                            return (d_3403_i29_) - (d_3402_p3_)

                        return lambda162_

                    init96_ = lambda161_(p3)
                    nw582_ = _dafny.Array(None, 9)
                    for i0_96_ in range(nw582_.length(0)):
                        nw582_[i0_96_] = init96_(i0_96_)
                    d_3401_v211_ = nw582_
                    d_3404_v212_: _dafny.Map
                    d_3404_v212_ = _dafny.Map({548: d_3401_v211_})
                    d_3405_v213_: _dafny.Map
                    d_3405_v213_ = _dafny.Map({_dafny.Map({p3: (D6_DC12(_dafny.SeqWithoutIsStrInference([p3]))).cf23}): d_3404_v212_})
                    d_3406_v214_: D0
                    d_3406_v214_ = D0_DC0(p1, len(d_3382_v192_), len(d_3385_v195_))
                    d_3407_v215_: _dafny.Map
                    d_3407_v215_ = _dafny.Map({(d_3211_v61_).fm6(d_3406_v214_, globalState): d_3118_v2_})
                    d_3405_v213_ = (d_3405_v213_).set(d_3407_v215_, _dafny.Map({p3: d_3401_v211_}))
                    d_3408_v216_: _dafny.Seq
                    d_3408_v216_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ntadwiem"))])
                    d_3408_v216_ = ((d_3408_v216_) + (d_3408_v216_)) + (d_3408_v216_)
                    pass
            pass

