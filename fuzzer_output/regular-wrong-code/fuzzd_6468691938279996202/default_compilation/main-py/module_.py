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
    def fm4(p0, globalState):
        return (_dafny.SeqWithoutIsStrInference([False])) + ((_dafny.SeqWithoutIsStrInference([True])) + (_dafny.SeqWithoutIsStrInference([True, False, False, False, False])))

    @staticmethod
    def fm6(p0, p1, p2, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_0_i0_ in range(default__.abs(984))])) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iwtniohu"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aa"))))

    @staticmethod
    def fm7(p0, globalState):
        return _dafny.CodePoint('l')

    @staticmethod
    def fm10(p0, p1, p2, p3, globalState):
        return (_dafny.MultiSet([(0) - (-104), 964])) - ((_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f')])), 57])).intersection(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rjys")))])))

    @staticmethod
    def fm11(p0, p1, p2, p3, globalState):
        return _dafny.Map({(907) * (len(_dafny.Map({-984: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bsuch")))}))): False})

    @staticmethod
    def fm12(p0, p1, p2, p3, globalState):
        return _dafny.Map({default__.safeDivisionInt((0) - (len(_dafny.Map({True: len(_dafny.Map({(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([False, False]))])])).cardinality: (0) - (len(_dafny.SeqWithoutIsStrInference([970 for d_1_i0_ in range(default__.abs(-992))])))}))}))), 782): (_dafny.CodePoint('u'))})

    @staticmethod
    def fm13(p0, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "frnc"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_2_i0_ in range(default__.abs(-34))]))

    @staticmethod
    def fm14(globalState):
        return _dafny.CodePoint('j')

    @staticmethod
    def fm15(p0, p1, p2, p3, globalState):
        return ((_dafny.MultiSet([True])) - (_dafny.MultiSet([False, True]))) | ((_dafny.MultiSet([False, True])).intersection(_dafny.MultiSet([True, False])))

    @staticmethod
    def fm16(p0, p1, p2, globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.Set({False}), _dafny.Set({False, True, False, True, False})])) + (_dafny.SeqWithoutIsStrInference([_dafny.Set({True}), _dafny.Set({not(not(False))}), _dafny.Set({False})]))) + ((D34_DC84(_dafny.SeqWithoutIsStrInference([_dafny.Set({True, True}) for d_3_i0_ in range(default__.abs(129))]))).cf145)

    @staticmethod
    def fm17(p0, globalState):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: int
            for compr_0_ in _dafny.IntegerRange(316, -425):
                d_4_v0_: int = compr_0_
                if ((316) <= (d_4_v0_)) and ((d_4_v0_) < (-425)):
                    coll0_[(d_4_v0_) * ((0) - (len(_dafny.Map({833: _dafny.CodePoint('b')}))))] = False
            return _dafny.Map(coll0_)
        return ((_dafny.Set({751, (_dafny.MultiSet([False])).cardinality})) | (_dafny.Set({len(_dafny.Map({True: True})), len(_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Map({_dafny.Map({935: True}): True})))])), (_dafny.MultiSet([False])).cardinality}))) - (_dafny.Set({82, (0) - (len(_dafny.Map({len(iife0_()
        ): 874})))}))

    @staticmethod
    def fm20(p0, p1, p2, p3, globalState):
        if True:
            return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "awte"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "skgl")))
        elif True:
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yr"))

    @staticmethod
    def fm21(globalState):
        source0_ = D1_DC4(False, not(True), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jee"))))
        if source0_.is_DC3:
            d_5___mcc_h0_ = source0_.cf5
            d_6_cf5_ = d_5___mcc_h0_
            return D1_DC2(_dafny.SeqWithoutIsStrInference([len(d_6_cf5_) for d_7_i0_ in range(default__.abs(126))]))
        elif source0_.is_DC4:
            d_8___mcc_h1_ = source0_.cf6
            d_9___mcc_h2_ = source0_.cf7
            d_10___mcc_h3_ = source0_.cf8
            d_11_cf8_ = d_10___mcc_h3_
            d_12_cf7_ = d_9___mcc_h2_
            d_13_cf6_ = d_8___mcc_h1_
            return D1_DC2(_dafny.SeqWithoutIsStrInference([246]))
        elif source0_.is_DC2:
            d_14___mcc_h4_ = source0_.cf4
            d_15_cf4_ = d_14___mcc_h4_
            return D1_DC2(d_15_cf4_)
        elif True:
            d_16___mcc_h5_ = source0_.cf9
            d_17_cf9_ = d_16___mcc_h5_
            return D1_DC2(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gkgim"))), -329]))

    @staticmethod
    def fm22(p0, p1, p2, p3, globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fihalg"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cileu")))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "urjn")))

    @staticmethod
    def fm23(globalState):
        if False:
            return D0_DC0(not(True))
        elif True:
            return D0_DC0(False)

    @staticmethod
    def fm24(p0, globalState):
        def iife1_():
            coll1_ = _dafny.Set()
            compr_1_: _dafny.Map
            for compr_1_ in (_dafny.MultiSet([_dafny.Map({True: True})])).Elements:
                d_18_v0_: _dafny.Map = compr_1_
                if (d_18_v0_) in (_dafny.MultiSet([_dafny.Map({True: True})])):
                    coll1_ = coll1_.union(_dafny.Set([d_18_v0_]))
            return _dafny.Set(coll1_)
        return (_dafny.Set({_dafny.Map({False: True}), _dafny.Map({True: False}), _dafny.Map({False: False})})).intersection((_dafny.Set({_dafny.Map({True: True})})).intersection(iife1_()
        ))

    @staticmethod
    def fm25(p0, p1, globalState):
        return (_dafny.Set({510})).intersection((_dafny.Set({239})) | (_dafny.Set({179})))

    @staticmethod
    def fm26(p0, globalState):
        source1_ = (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "alriycn"))): 463}) if True else _dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_19_i0_ in range(default__.abs(896))])): 486}))
        d_20___mcc_h0_ = source1_
        d_21_cf45_ = d_20___mcc_h0_
        return _dafny.CodePoint('c')

    @staticmethod
    def fm29(p0, p1, p2, p3, globalState):
        return ((_dafny.Map({True: _dafny.MultiSet([_dafny.CodePoint('o')])})) | (_dafny.Map({True: _dafny.MultiSet([_dafny.CodePoint('v')])}))) | (_dafny.Map({True: _dafny.MultiSet([_dafny.CodePoint('x'), _dafny.CodePoint('h'), _dafny.CodePoint('o'), _dafny.CodePoint('s')])}))

    @staticmethod
    def fm30(globalState):
        return _dafny.Set({True})

    @staticmethod
    def fm31(p0, p1, globalState):
        return (D6_DC18(_dafny.CodePoint('l'), 486, len(_dafny.Set({False})))).cf34

    @staticmethod
    def fm33(p0, p1, p2, p3, globalState):
        source2_ = D6_DC18(_dafny.CodePoint('p'), len(_dafny.Set({-530})), (_dafny.MultiSet([False])).cardinality)
        if source2_.is_DC17:
            d_22___mcc_h0_ = source2_.cf29
            d_23___mcc_h1_ = source2_.cf30
            d_24___mcc_h2_ = source2_.cf31
            d_25___mcc_h3_ = source2_.cf32
            d_26___mcc_h4_ = source2_.cf33
            d_27_cf33_ = d_26___mcc_h4_
            d_28_cf32_ = d_25___mcc_h3_
            d_29_cf31_ = d_24___mcc_h2_
            d_30_cf30_ = d_23___mcc_h1_
            d_31_cf29_ = d_22___mcc_h0_
            return d_29_cf31_
        elif source2_.is_DC18:
            d_32___mcc_h5_ = source2_.cf34
            d_33___mcc_h6_ = source2_.cf35
            d_34___mcc_h7_ = source2_.cf36
            d_35_cf36_ = d_34___mcc_h7_
            d_36_cf35_ = d_33___mcc_h6_
            d_37_cf34_ = d_32___mcc_h5_
            return d_36_cf35_
        elif source2_.is_DC19:
            d_38___mcc_h8_ = source2_.cf37
            d_39___mcc_h9_ = source2_.cf38
            d_40_cf38_ = d_39___mcc_h9_
            d_41_cf37_ = d_38___mcc_h8_
            def iife2_():
                coll2_ = _dafny.Map()
                compr_2_: int
                for compr_2_ in _dafny.IntegerRange(-815, 685):
                    d_42_v0_: int = compr_2_
                    if ((-815) <= (d_42_v0_)) and ((d_42_v0_) < (685)):
                        coll2_[default__.safeModuloInt(d_42_v0_, 371)] = (D4_DC11(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fv"))), _dafny.Set({len(_dafny.Map({847: d_40_cf38_}))}), d_41_cf37_)).cf18
                return _dafny.Map(coll2_)
            return default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([iife2_()
            ])), 618)
        elif source2_.is_DC16:
            d_43___mcc_h10_ = source2_.cf28
            d_44_cf28_ = d_43___mcc_h10_
            return (0) - (len((_dafny.Map({True: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dgejc"))})) | (_dafny.Map({not(True): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vben"))}))))
        elif True:
            d_45___mcc_h11_ = source2_.cf39
            d_46_cf39_ = d_45___mcc_h11_
            return 408

    @staticmethod
    def fm34(p0, p1, p2, p3, globalState):
        return _dafny.Set({(False if True else False)})

    @staticmethod
    def fm35(p0, globalState):
        def iife3_():
            coll3_ = _dafny.Map()
            compr_3_: _dafny.Seq
            for compr_3_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference([D4_DC11(-429, _dafny.Set({694, 859}), True) for d_47_i0_ in range(default__.abs(281))]): _dafny.CodePoint('q')})).keys.Elements:
                d_48_v0_: _dafny.Seq = compr_3_
                if (d_48_v0_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference([D4_DC11(-429, _dafny.Set({694, 859}), True) for d_47_i0_ in range(default__.abs(281))]): _dafny.CodePoint('q')})):
                    coll3_[d_48_v0_] = -622
            return _dafny.Map(coll3_)
        source3_ = D14_DC34(iife3_()
)
        if source3_.is_DC35:
            d_49___mcc_h0_ = source3_.cf62
            d_50___mcc_h1_ = source3_.cf63
            d_51___mcc_h2_ = source3_.cf64
            d_52_cf64_ = d_51___mcc_h2_
            d_53_cf63_ = d_50___mcc_h1_
            d_54_cf62_ = d_49___mcc_h0_
            return D3_DC8(d_53_cf63_, d_53_cf63_)
        elif source3_.is_DC36:
            d_55___mcc_h3_ = source3_.cf65
            d_56___mcc_h4_ = source3_.cf66
            d_57___mcc_h5_ = source3_.cf67
            d_58_cf67_ = d_57___mcc_h5_
            d_59_cf66_ = d_56___mcc_h4_
            d_60_cf65_ = d_55___mcc_h3_
            return D3_DC8(not(d_60_cf65_), d_60_cf65_)
        elif True:
            d_61___mcc_h6_ = source3_.cf61
            d_62_cf61_ = d_61___mcc_h6_
            return D3_DC8(not(False), not(True))

    @staticmethod
    def fm36(globalState):
        source4_ = D6_DC20(D6_DC18(_dafny.CodePoint('m'), 779, len(_dafny.Map({True: 555}))))
        if source4_.is_DC17:
            d_63___mcc_h0_ = source4_.cf29
            d_64___mcc_h1_ = source4_.cf30
            d_65___mcc_h2_ = source4_.cf31
            d_66___mcc_h3_ = source4_.cf32
            d_67___mcc_h4_ = source4_.cf33
            d_68_cf33_ = d_67___mcc_h4_
            d_69_cf32_ = d_66___mcc_h3_
            d_70_cf31_ = d_65___mcc_h2_
            d_71_cf30_ = d_64___mcc_h1_
            d_72_cf29_ = d_63___mcc_h0_
            return _dafny.CodePoint('o')
        elif source4_.is_DC18:
            d_73___mcc_h5_ = source4_.cf34
            d_74___mcc_h6_ = source4_.cf35
            d_75___mcc_h7_ = source4_.cf36
            d_76_cf36_ = d_75___mcc_h7_
            d_77_cf35_ = d_74___mcc_h6_
            d_78_cf34_ = d_73___mcc_h5_
            return _dafny.CodePoint('x')
        elif source4_.is_DC19:
            d_79___mcc_h8_ = source4_.cf37
            d_80___mcc_h9_ = source4_.cf38
            d_81_cf38_ = d_80___mcc_h9_
            d_82_cf37_ = d_79___mcc_h8_
            return _dafny.CodePoint('g')
        elif source4_.is_DC16:
            d_83___mcc_h10_ = source4_.cf28
            d_84_cf28_ = d_83___mcc_h10_
            return _dafny.CodePoint('s')
        elif True:
            d_85___mcc_h11_ = source4_.cf39
            d_86_cf39_ = d_85___mcc_h11_
            return _dafny.CodePoint('q')

    @staticmethod
    def fm37(globalState):
        return ((_dafny.SeqWithoutIsStrInference([929 for d_87_i0_ in range(default__.abs(972))]) if False else _dafny.SeqWithoutIsStrInference([342, 642]))) + (_dafny.SeqWithoutIsStrInference([208]))

    @staticmethod
    def fm38(p0, globalState):
        return (_dafny.SeqWithoutIsStrInference([False, True, False, not(False)])) + (_dafny.SeqWithoutIsStrInference([False, False, False]))

    @staticmethod
    def fm39(p0, p1, p2, globalState):
        return _dafny.Map({default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([True, False, True, not(True), True])), len(_dafny.Map({335: not(True)}))): ((0) - (len(_dafny.Set({-919}))) if False else 247)})

    @staticmethod
    def fm40(p0, p1, p2, globalState):
        def iife4_():
            coll4_ = _dafny.Map()
            compr_4_: int
            for compr_4_ in (_dafny.Map({(_dafny.MultiSet([631])).cardinality: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_88_i1_ in range(default__.abs(575))]))})).keys.Elements:
                d_89_v0_: int = compr_4_
                if (d_89_v0_) in (_dafny.Map({(_dafny.MultiSet([631])).cardinality: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_88_i1_ in range(default__.abs(575))]))})):
                    coll4_[(d_89_v0_) * (len(_dafny.Map({True: not(True)})))] = 696
            return _dafny.Map(coll4_)
        return (_dafny.SeqWithoutIsStrInference([_dafny.Map({188: -896}), _dafny.Map({894: 511})])) + ((_dafny.SeqWithoutIsStrInference([iife4_()
 for d_90_i0_ in range(default__.abs(698))])) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({587: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cldmyuqk")))}), _dafny.Map({len(_dafny.SeqWithoutIsStrInference([False])): 179})])))

    @staticmethod
    def fm43(p0, globalState):
        if True:
            return _dafny.CodePoint('y')
        elif True:
            return _dafny.CodePoint('b')

    @staticmethod
    def fm44(p0, globalState):
        def iife5_():
            coll5_ = _dafny.Map()
            compr_5_: int
            for compr_5_ in (_dafny.MultiSet([(D17_DC44(357, _dafny.SeqWithoutIsStrInference([False]), _dafny.CodePoint('t'), True, 983)).cf84])).Elements:
                d_91_v0_: int = compr_5_
                if (d_91_v0_) in (_dafny.MultiSet([(D17_DC44(357, _dafny.SeqWithoutIsStrInference([False]), _dafny.CodePoint('t'), True, 983)).cf84])):
                    coll5_[(d_91_v0_) + (-17)] = not(True)
            return _dafny.Map(coll5_)
        return (iife5_()
        ) | ((D35_DC87(_dafny.Map({len((D15_DC38(False, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_92_i0_ in range(default__.abs(331))]), True, not(True), _dafny.CodePoint('i'))).cf70): True}))).cf149)

    @staticmethod
    def fm45(p0, p1, p2, globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wgb")))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lkwgqmac"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_93_i0_ in range(default__.abs(878))])))

    @staticmethod
    def fm46(p0, p1, p2, globalState):
        return (_dafny.MultiSet([347, (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ye"))))])).intersection((_dafny.MultiSet([-658])) - (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([False])), 42])))

    @staticmethod
    def fm47(p0, globalState):
        if not(True):
            return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bvpnsv")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f"))])
        elif True:
            return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_94_i0_ in range(default__.abs(420))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ch"))])

    @staticmethod
    def fm48(p0, globalState):
        if (False) or (True):
            return (_dafny.Map({False: False})) | (_dafny.Map({not(True): False}))
        elif True:
            return (_dafny.Map({True: True})) | (_dafny.Map({not(not(True)): True}))

    @staticmethod
    def fm49(p0, globalState):
        def iife6_():
            def iife8_():
                coll8_ = _dafny.Map()
                compr_8_: _dafny.Seq
                for compr_8_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r')])])).Elements:
                    d_96_v1_: _dafny.Seq = compr_8_
                    if (d_96_v1_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r')])])):
                        coll8_[d_96_v1_] = False
                return _dafny.Map(coll8_)
            coll6_ = _dafny.Map()
            def iife7_():
                coll7_ = _dafny.Map()
                compr_7_: _dafny.Seq
                for compr_7_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r')])])).Elements:
                    d_96_v1_: _dafny.Seq = compr_7_
                    if (d_96_v1_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r')])])):
                        coll7_[d_96_v1_] = False
                return _dafny.Map(coll7_)
            compr_6_: _dafny.Seq
            for compr_6_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_95_i0_ in range(default__.abs(822))]): len(iife7_()
            )})).keys.Elements:
                d_97_v0_: _dafny.Seq = compr_6_
                if (d_97_v0_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_95_i0_ in range(default__.abs(822))]): len(iife8_()
                )})):
                    coll6_[d_97_v0_] = -740
            return _dafny.Map(coll6_)
        return D6_DC17(True, (_dafny.MultiSet([-731])).issubset(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(iife6_()
)]))), 330, (3) * (285), default__.safeDivisionInt(-341, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sn")))))

    @staticmethod
    def fm50(p0, p1, globalState):
        return (_dafny.MultiSet([True])) - (_dafny.MultiSet([not(False)]))

    @staticmethod
    def fm51(globalState):
        if (_dafny.SeqWithoutIsStrInference([-181 for d_98_i0_ in range(default__.abs(106))])) <= (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([(_dafny.CodePoint('d'))])).cardinality])), len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_99_i1_ in range(default__.abs(516))]))).cardinality]))])):
            return D14_DC34(_dafny.Map({_dafny.SeqWithoutIsStrInference([D4_DC11(-556, _dafny.Set({253}), True) for d_100_i2_ in range(default__.abs(-130))]): len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_101_i3_ in range(default__.abs(5))]))}))
        elif True:
            def iife9_():
                coll9_ = _dafny.Set()
                compr_9_: int
                for compr_9_ in _dafny.IntegerRange(712, 679):
                    d_102_v0_: int = compr_9_
                    if ((712) <= (d_102_v0_)) and ((d_102_v0_) < (679)):
                        coll9_ = coll9_.union(_dafny.Set([(d_102_v0_) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iqsvf"))))]))
                return _dafny.Set(coll9_)
            return D14_DC34(_dafny.Map({_dafny.SeqWithoutIsStrInference([D4_DC11(636, iife9_()
, True)]): len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([False, True, True, not(True)]))}))}))

    @staticmethod
    def fm52(p0, globalState):
        def iife10_():
            def iife14_():
                def iife16_():
                    coll16_ = _dafny.Map()
                    compr_16_: int
                    for compr_16_ in _dafny.IntegerRange(419, 579):
                        d_103_v2_: int = compr_16_
                        if ((419) <= (d_103_v2_)) and ((d_103_v2_) < (579)):
                            coll16_[default__.safeDivisionInt(d_103_v2_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sfqi"))))] = 971
                    return _dafny.Map(coll16_)
                coll14_ = _dafny.Map()
                def iife15_():
                    coll15_ = _dafny.Map()
                    compr_15_: int
                    for compr_15_ in _dafny.IntegerRange(419, 579):
                        d_103_v2_: int = compr_15_
                        if ((419) <= (d_103_v2_)) and ((d_103_v2_) < (579)):
                            coll15_[default__.safeDivisionInt(d_103_v2_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sfqi"))))] = 971
                    return _dafny.Map(coll15_)
                compr_14_: int
                for compr_14_ in (iife15_()
                ).keys.Elements:
                    d_104_v1_: int = compr_14_
                    if (d_104_v1_) in (iife16_()
                    ):
                        coll14_[default__.safeModuloInt(d_104_v1_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_105_i0_ in range(default__.abs(689))])))] = True
                return _dafny.Map(coll14_)
            coll10_ = _dafny.Map()
            def iife11_():
                def iife13_():
                    coll13_ = _dafny.Map()
                    compr_13_: int
                    for compr_13_ in _dafny.IntegerRange(419, 579):
                        d_103_v2_: int = compr_13_
                        if ((419) <= (d_103_v2_)) and ((d_103_v2_) < (579)):
                            coll13_[default__.safeDivisionInt(d_103_v2_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sfqi"))))] = 971
                    return _dafny.Map(coll13_)
                coll11_ = _dafny.Map()
                def iife12_():
                    coll12_ = _dafny.Map()
                    compr_12_: int
                    for compr_12_ in _dafny.IntegerRange(419, 579):
                        d_103_v2_: int = compr_12_
                        if ((419) <= (d_103_v2_)) and ((d_103_v2_) < (579)):
                            coll12_[default__.safeDivisionInt(d_103_v2_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sfqi"))))] = 971
                    return _dafny.Map(coll12_)
                compr_11_: int
                for compr_11_ in (iife12_()
                ).keys.Elements:
                    d_104_v1_: int = compr_11_
                    if (d_104_v1_) in (iife13_()
                    ):
                        coll11_[default__.safeModuloInt(d_104_v1_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_105_i0_ in range(default__.abs(689))])))] = True
                return _dafny.Map(coll11_)
            compr_10_: int
            for compr_10_ in (iife11_()
            ).keys.Elements:
                d_106_v0_: int = compr_10_
                if (d_106_v0_) in (iife14_()
                ):
                    coll10_[default__.safeModuloInt(d_106_v0_, (_dafny.MultiSet([not(False)])).cardinality)] = -334
            return _dafny.Map(coll10_)
        if not((len(iife10_()
        )) < (len(_dafny.SeqWithoutIsStrInference([True])))):
            return D8_DC23(_dafny.SeqWithoutIsStrInference([True]))
        elif True:
            return D8_DC23(_dafny.SeqWithoutIsStrInference([not(not(True))]))

    @staticmethod
    def fm53(p0, globalState):
        def iife17_():
            coll17_ = _dafny.Set()
            compr_17_: int
            for compr_17_ in _dafny.IntegerRange(845, 623):
                d_107_v0_: int = compr_17_
                if ((845) <= (d_107_v0_)) and ((d_107_v0_) < (623)):
                    coll17_ = coll17_.union(_dafny.Set([default__.safeDivisionInt(d_107_v0_, -722)]))
            return _dafny.Set(coll17_)
        def iife18_():
            coll18_ = _dafny.Set()
            compr_18_: int
            for compr_18_ in _dafny.IntegerRange(-473, 423):
                d_108_v1_: int = compr_18_
                if ((-473) <= (d_108_v1_)) and ((d_108_v1_) < (423)):
                    coll18_ = coll18_.union(_dafny.Set([(d_108_v1_) + (470)]))
            return _dafny.Set(coll18_)
        return _dafny.Map({(not(True) if not(not(False)) else False): (D4_DC11(859, iife17_()
, False) if False else D4_DC11(len(_dafny.Set({not(False), True, not(not(not(False)))})), iife18_()
, not(True)))})

    @staticmethod
    def fm54(globalState):
        def iife19_():
            coll19_ = _dafny.Set()
            compr_19_: str
            for compr_19_ in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_109_i0_ in range(default__.abs(748))]))).Elements:
                d_110_v0_: str = compr_19_
                if (d_110_v0_) in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_109_i0_ in range(default__.abs(748))]))):
                    coll19_ = coll19_.union(_dafny.Set([d_110_v0_]))
            return _dafny.Set(coll19_)
        return iife19_()
        

    @staticmethod
    def fm55(p0, globalState):
        return D6_DC20((D6_DC19(not(False), _dafny.CodePoint('i')) if True else D6_DC20(D6_DC20(D6_DC19(False, _dafny.CodePoint('q'))))))

    @staticmethod
    def fm56(p0, globalState):
        return ((_dafny.SeqWithoutIsStrInference([D14_DC36(False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mpbldi")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dccddbi"))) for d_111_i0_ in range(default__.abs(-87))])) + (_dafny.SeqWithoutIsStrInference([D14_DC36(False, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_112_i1_ in range(default__.abs(108))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xdxoru"))), D14_DC36(True, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "igbltr")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_113_i2_ in range(default__.abs(982))]))]))) + (_dafny.SeqWithoutIsStrInference([D14_DC36(not(not(True)), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kwqsarg")))]))

    @staticmethod
    def fm57(p0, globalState):
        return ((_dafny.MultiSet([_dafny.CodePoint('g')])) - (_dafny.MultiSet([_dafny.CodePoint('s')]))) - ((_dafny.MultiSet([_dafny.CodePoint('k'), _dafny.CodePoint('q'), _dafny.CodePoint('i'), _dafny.CodePoint('m')])) | (_dafny.MultiSet([_dafny.CodePoint('i'), _dafny.CodePoint('r')])))

    @staticmethod
    def fm58(globalState):
        def iife20_():
            coll20_ = _dafny.Set()
            compr_20_: int
            for compr_20_ in _dafny.IntegerRange(630, 790):
                d_114_v0_: int = compr_20_
                if ((630) <= (d_114_v0_)) and ((d_114_v0_) < (790)):
                    coll20_ = coll20_.union(_dafny.Set([(d_114_v0_) - (len(_dafny.Map({False: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mnhgh"))})))]))
            return _dafny.Set(coll20_)
        return _dafny.SeqWithoutIsStrInference([(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([True]))})) | (_dafny.Set({(_dafny.MultiSet([False])).cardinality})), iife20_()
        , _dafny.Set({len(_dafny.Map({not(True): len(_dafny.Set({False, False}))}))})])

    @staticmethod
    def fm59(p0, p1, globalState):
        return ((_dafny.Map({_dafny.CodePoint('p'): _dafny.SeqWithoutIsStrInference([not(True), True, not(True)])})) | (_dafny.Map({_dafny.CodePoint('b'): (D17_DC44((D34_DC86(True, 906)).cf148, _dafny.SeqWithoutIsStrInference([True]), _dafny.CodePoint('p'), not(True), len(_dafny.Set({False})))).cf85}))) | (_dafny.Map({_dafny.CodePoint('h'): _dafny.SeqWithoutIsStrInference([not(not(False)), True])}))

    @staticmethod
    def fm60(p0, p1, globalState):
        source5_ = _dafny.SeqWithoutIsStrInference([D4_DC11((_dafny.MultiSet([True])).cardinality, _dafny.Set({943, -813, -314}), not(False))])
        d_115___mcc_h0_ = source5_
        d_116_cf60_ = d_115___mcc_h0_
        return D15_DC39(True, True)

    @staticmethod
    def fm61(p0, globalState):
        source6_ = D34_DC85(_dafny.CodePoint('p'))
        if source6_.is_DC85:
            d_117___mcc_h0_ = source6_.cf146
            d_118_cf146_ = d_117___mcc_h0_
            return _dafny.MultiSet([857])
        elif source6_.is_DC86:
            d_119___mcc_h1_ = source6_.cf147
            d_120___mcc_h2_ = source6_.cf148
            d_121_cf148_ = d_120___mcc_h2_
            d_122_cf147_ = d_119___mcc_h1_
            if d_122_cf147_:
                return _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_121_cf148_, len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ey")) for d_123_i0_ in range(default__.abs(214))]))]))
            elif True:
                return _dafny.MultiSet([d_121_cf148_, d_121_cf148_])
        elif True:
            d_124___mcc_h3_ = source6_.cf145
            d_125_cf145_ = d_124___mcc_h3_
            return _dafny.MultiSet([-923, len(_dafny.Set({True})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qgxceec"))), 738])

    @staticmethod
    def Main(noArgsParameter__):
        d_126_v0_: _dafny.Array
        def lambda0_(d_127_i0_):
            return False

        init0_ = lambda0_
        nw0_ = _dafny.Array(None, 18)
        for i0_0_ in range(nw0_.length(0)):
            nw0_[i0_0_] = init0_(i0_0_)
        d_126_v0_ = nw0_
        d_128_v1_: bool
        d_128_v1_ = True
        d_129_v2_: _dafny.Map
        d_129_v2_ = _dafny.Map({4: d_128_v1_})
        d_130_globalState_: GlobalState
        nw1_ = GlobalState()
        nw1_.ctor__(False, False, d_126_v0_, True, _dafny.Map({d_129_v2_: not(d_128_v1_)}), -523, False, -26, False)
        d_130_globalState_ = nw1_
        d_131_v3_: C1
        nw2_ = C1()
        nw2_.ctor__(d_128_v1_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_132_i1_ in range(default__.abs(880))]))
        d_131_v3_ = nw2_
        d_133_v4_: _dafny.Array
        nw3_ = _dafny.Array(_dafny.Array(None, 0), 10)
        d_133_v4_ = nw3_
        d_134_v5_: D17
        d_134_v5_ = D17_DC43(d_133_v4_)
        d_135_v6_: _dafny.MultiSet
        d_135_v6_ = _dafny.MultiSet([d_134_v5_])
        d_136_v7_: _dafny.Map
        d_136_v7_ = _dafny.Map({d_135_v6_: len(_dafny.SeqWithoutIsStrInference([d_128_v1_]))})
        d_137_v8_: int
        d_137_v8_ = 925
        if (d_131_v3_).fm32(((d_136_v7_)[(d_135_v6_).set(d_134_v5_, default__.abs(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rdlthb")))))] if ((d_135_v6_).set(d_134_v5_, default__.abs(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rdlthb")))))) in (d_136_v7_) else d_137_v8_), (_dafny.CodePoint('j') if d_128_v1_ else _dafny.CodePoint('l')), d_130_globalState_):
            d_138_v9_: C0
            nw4_ = C0()
            nw4_.ctor__()
            d_138_v9_ = nw4_
            d_139_v10_: str
            d_139_v10_ = _dafny.CodePoint('v')
            d_137_v8_ = len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pc"))).set(default__.safeIndex(len((d_131_v3_).f13), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pc")))), d_139_v10_))
            (d_130_globalState_).f8 = not((d_137_v8_) != (default__.safeDivisionInt(d_137_v8_, 67)))
            d_140_v11_: _dafny.Map
            d_140_v11_ = _dafny.Map({not(((d_131_v3_).f12 if (d_131_v3_).f12 else (d_131_v3_).f12)): d_137_v8_})
            d_141_v12_: _dafny.Seq
            d_141_v12_ = _dafny.SeqWithoutIsStrInference([d_128_v1_, (d_131_v3_).f12])
            d_142_v13_: _dafny.Set
            d_142_v13_ = _dafny.Set({d_137_v8_, 521})
            d_140_v11_ = (d_140_v11_).set((d_131_v3_).f12, len((_dafny.Map({d_139_v10_: d_141_v12_})) | (default__.fm59(len(d_142_v13_), d_137_v8_, d_130_globalState_))))
            d_143_v14_: D12
            d_143_v14_ = D12_DC31((d_131_v3_).f12, (d_131_v3_).f13, (d_131_v3_).f13)
            d_144_v15_: _dafny.Map
            d_144_v15_ = _dafny.Map({d_137_v8_: (d_131_v3_).f13})
            d_145_v16_: D28
            d_145_v16_ = D28_DC67(default__.fm58(d_130_globalState_))
            d_146_v17_: _dafny.Array
            nw5_ = _dafny.Array(None, 20)
            nw5_[int(0)] = ((d_131_v3_).f13) + ((d_131_v3_).f13)
            nw5_[int(1)] = (d_131_v3_).f13
            nw5_[int(2)] = (d_131_v3_).f13
            nw5_[int(3)] = (d_131_v3_).f13
            nw5_[int(4)] = ((d_131_v3_).f13) + ((d_131_v3_).f13)
            nw5_[int(5)] = ((d_131_v3_).f13) + ((d_131_v3_).f13)
            nw5_[int(6)] = ((d_131_v3_).f13) + ((d_143_v14_).cf56)
            nw5_[int(7)] = (d_131_v3_).f13
            nw5_[int(8)] = ((d_131_v3_).f13) + ((d_131_v3_).f13)
            nw5_[int(9)] = default__.fm22(d_137_v8_, (d_131_v3_).f12, False, d_137_v8_, d_130_globalState_)
            nw5_[int(10)] = (((d_131_v3_).f13).set(default__.safeIndex(d_137_v8_, len((d_131_v3_).f13)), d_139_v10_)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bmidbbbbs")))
            nw5_[int(11)] = (d_131_v3_).f13
            nw5_[int(12)] = (((d_144_v15_)[(0) - (len(_dafny.Map({d_139_v10_: (default__.fm60(d_145_v16_, (d_131_v3_).f12, d_130_globalState_)).cf75})))] if ((0) - (len(_dafny.Map({d_139_v10_: (default__.fm60(d_145_v16_, (d_131_v3_).f12, d_130_globalState_)).cf75})))) in (d_144_v15_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "etfudx")))) + ((d_131_v3_).f13)
            nw5_[int(13)] = (d_131_v3_).f13
            nw5_[int(14)] = (d_131_v3_).f13
            nw5_[int(15)] = _dafny.SeqWithoutIsStrInference([d_139_v10_ for d_147_i2_ in range(default__.abs(13))])
            nw5_[int(16)] = default__.fm13(d_137_v8_, d_130_globalState_)
            nw5_[int(17)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ieoynisy"))
            nw5_[int(18)] = (d_131_v3_).f13
            nw5_[int(19)] = default__.fm45((d_131_v3_).f12, d_137_v8_, (d_131_v3_).f12, d_130_globalState_)
            d_146_v17_ = nw5_
            d_146_v17_ = d_146_v17_
        elif True:
            d_148_v18_: C0
            nw6_ = C0()
            nw6_.ctor__()
            d_148_v18_ = nw6_
            d_149_v19_: _dafny.Map
            d_149_v19_ = _dafny.Map({False: d_137_v8_})
            d_150_v20_: _dafny.Map
            d_150_v20_ = _dafny.Map({(d_131_v3_).f12: d_128_v1_})
            d_149_v19_ = (d_149_v19_).set(d_128_v1_, default__.safeDivisionInt(d_137_v8_, (0) - (len(d_150_v20_))))
            if (d_131_v3_).f12:
                d_137_v8_ = (d_137_v8_) - (default__.safeModuloInt(d_137_v8_, default__.fm33((d_131_v3_).f12, (d_131_v3_).f12, d_137_v8_, (d_131_v3_).f12, d_130_globalState_)))
                d_151_v21_: _dafny.MultiSet
                d_151_v21_ = _dafny.MultiSet([(d_131_v3_).f12])
                d_152_v22_: _dafny.Map
                d_152_v22_ = _dafny.Map({((d_131_v3_).f12 if d_128_v1_ else d_128_v1_): d_151_v21_})
                d_152_v22_ = (d_152_v22_).set((d_137_v8_) < (d_137_v8_), d_151_v21_)
                d_153_v23_: _dafny.Seq
                d_153_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ynqtib"))
                d_154_v24_: _dafny.Set
                d_154_v24_ = _dafny.Set({not((d_131_v3_).f12)})
                d_155_v25_: _dafny.Map
                d_155_v25_ = _dafny.Map({default__.fm44((d_131_v3_).f12, d_130_globalState_): True})
                rhs0_ = len((d_154_v24_) - (default__.fm34(d_137_v8_, (d_131_v3_).f12, d_128_v1_, ((d_155_v25_)[d_129_v2_] if (d_129_v2_) in (d_155_v25_) else (d_131_v3_).f12), d_130_globalState_)))
                rhs1_ = (((d_131_v3_).f13 if (d_131_v3_).f12 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i")))) + (((d_131_v3_).f13) + ((d_131_v3_).f13))
                d_137_v8_ = rhs0_
                d_153_v23_ = rhs1_
                d_156_v26_: C5
                nw7_ = C5()
                nw7_.ctor__(d_150_v20_, (d_131_v3_).f12)
                d_156_v26_ = nw7_
                d_156_v26_ = d_156_v26_
                d_157_v27_: str
                d_157_v27_ = _dafny.CodePoint('f')
                d_153_v23_ = ((d_131_v3_).f13).set(default__.safeIndex(d_137_v8_, len((d_131_v3_).f13)), d_157_v27_)
            elif True:
                d_137_v8_ = -455
                d_158_v28_: T0
                nw8_ = C3()
                nw8_.ctor__()
                d_158_v28_ = nw8_
                d_159_v29_: str
                d_159_v29_ = _dafny.CodePoint('o')
                rhs2_ = not(((d_137_v8_) + (d_137_v8_)) != (default__.safeModuloInt(d_137_v8_, d_137_v8_)))
                rhs3_ = (d_128_v1_ if not((d_131_v3_).fm32(d_137_v8_, d_159_v29_, d_130_globalState_)) else d_128_v1_)
                rhs4_ = d_158_v28_
                rhs5_ = default__.safeDivisionInt(d_137_v8_, d_137_v8_)
                lhs0_ = d_130_globalState_
                lhs0_.f8 = rhs2_
                d_128_v1_ = rhs3_
                d_158_v28_ = rhs4_
                d_137_v8_ = rhs5_
                d_137_v8_ = (0) - (default__.safeModuloInt(d_137_v8_, d_137_v8_))
                d_137_v8_ = default__.fm33((d_131_v3_).f12, (d_131_v3_).f12, default__.fm33((d_131_v3_).f12, d_128_v1_, 590, (d_131_v3_).f12, d_130_globalState_), d_128_v1_, d_130_globalState_)
                d_137_v8_ = d_137_v8_
            d_160_v30_: _dafny.Array
            nw9_ = _dafny.Array(int(0), 19)
            d_160_v30_ = nw9_
            index0_ = default__.safeIndex(402, (d_160_v30_).length(0))
            (d_160_v30_)[index0_] = len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mlch"))) + ((d_131_v3_).f13))
            index1_ = default__.safeIndex(402, (d_160_v30_).length(0))
            (d_160_v30_)[index1_] = (d_137_v8_) - (default__.fm33(d_128_v1_, (d_131_v3_).f12, (0) - (d_137_v8_), not(True), d_130_globalState_))
            d_161_v32_: _dafny.Array
            def lambda1_(d_162_v3_):
                def lambda2_(d_163_i3_):
                    def iife21_():
                        coll21_ = _dafny.Map()
                        compr_21_: _dafny.Seq
                        for compr_21_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_164_i5_ in range(default__.abs(610))]) for d_165_i4_ in range(default__.abs(271))])).Elements:
                            d_166_v31_: _dafny.Seq = compr_21_
                            if (d_166_v31_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_164_i5_ in range(default__.abs(610))]) for d_165_i4_ in range(default__.abs(271))])):
                                coll21_[d_166_v31_] = len((d_162_v3_).f13)
                        return _dafny.Map(coll21_)
                    return iife21_()
                    

                return lambda2_

            init1_ = lambda1_(d_131_v3_)
            nw10_ = _dafny.Array(None, 7)
            for i0_1_ in range(nw10_.length(0)):
                nw10_[i0_1_] = init1_(i0_1_)
            d_161_v32_ = nw10_
            d_167_v33_: _dafny.Map
            d_167_v33_ = _dafny.Map({(d_131_v3_).f13: (d_160_v30_)[default__.safeIndex(402, (d_160_v30_).length(0))]})
            d_168_v34_: str
            d_168_v34_ = _dafny.CodePoint('p')
            index2_ = default__.safeIndex(362, (d_161_v32_).length(0))
            (d_161_v32_)[index2_] = ((d_167_v33_).set(_dafny.SeqWithoutIsStrInference([d_168_v34_, d_168_v34_]), (_dafny.MultiSet([d_137_v8_])).cardinality)).set((d_131_v3_).f13, d_137_v8_)
            d_169_v35_: _dafny.Map
            d_169_v35_ = _dafny.Map({d_160_v30_: d_137_v8_})
            d_170_v36_: _dafny.MultiSet
            d_170_v36_ = _dafny.MultiSet([(d_131_v3_).f12])
            index3_ = default__.safeIndex(362, (d_161_v32_).length(0))
            rhs6_ = (d_137_v8_ if ((d_169_v35_)) == (d_169_v35_) else (_dafny.MultiSet([(d_170_v36_).cardinality])).cardinality)
            rhs7_ = ((d_167_v33_).set((d_131_v3_).f13, (d_160_v30_)[default__.safeIndex(402, (d_160_v30_).length(0))])) | ((_dafny.Map({(d_131_v3_).f13: (d_160_v30_)[default__.safeIndex(402, (d_160_v30_).length(0))]})) | ((d_167_v33_).set(_dafny.SeqWithoutIsStrInference([d_168_v34_ for d_171_i6_ in range(default__.abs(697))]), (d_160_v30_)[default__.safeIndex(402, (d_160_v30_).length(0))])))
            rhs8_ = d_168_v34_
            lhs1_ = d_161_v32_
            lhs2_ = default__.safeIndex(362, (d_161_v32_).length(0))
            d_137_v8_ = rhs6_
            lhs1_[lhs2_] = rhs7_
            d_168_v34_ = rhs8_
        d_172_v37_: str
        d_172_v37_ = _dafny.CodePoint('q')
        d_172_v37_ = d_172_v37_
        d_173_v38_: _dafny.Map
        d_173_v38_ = _dafny.Map({(d_131_v3_).f12: d_128_v1_})
        d_174_v39_: T1
        nw11_ = C6()
        nw11_.ctor__(d_173_v38_, True)
        d_174_v39_ = nw11_
        d_175_v40_: D22
        d_175_v40_ = D22_DC56(d_174_v39_)
        d_176_v41_: _dafny.Array
        nw12_ = _dafny.Array(None, 26)
        nw12_[int(0)] = d_174_v39_
        nw12_[int(1)] = d_174_v39_
        nw12_[int(2)] = d_174_v39_
        nw12_[int(3)] = d_174_v39_
        nw12_[int(4)] = d_174_v39_
        nw12_[int(5)] = d_174_v39_
        nw12_[int(6)] = d_174_v39_
        nw12_[int(7)] = d_174_v39_
        nw12_[int(8)] = d_174_v39_
        nw12_[int(9)] = d_174_v39_
        nw12_[int(10)] = d_174_v39_
        nw12_[int(11)] = d_174_v39_
        nw12_[int(12)] = d_174_v39_
        nw12_[int(13)] = d_174_v39_
        nw12_[int(14)] = d_174_v39_
        nw12_[int(15)] = d_174_v39_
        nw12_[int(16)] = d_174_v39_
        nw12_[int(17)] = d_174_v39_
        nw12_[int(18)] = d_174_v39_
        nw12_[int(19)] = d_174_v39_
        nw12_[int(20)] = (d_175_v40_).cf104
        nw12_[int(21)] = d_174_v39_
        nw12_[int(22)] = d_174_v39_
        nw12_[int(23)] = d_174_v39_
        nw12_[int(24)] = d_174_v39_
        nw12_[int(25)] = (d_174_v39_ if (d_174_v39_).f10 else d_174_v39_)
        d_176_v41_ = nw12_
        index4_ = default__.safeIndex(846, (d_176_v41_).length(0))
        (d_176_v41_)[index4_] = d_174_v39_
        d_177_v42_: _dafny.Seq
        d_177_v42_ = _dafny.SeqWithoutIsStrInference([d_174_v39_, d_174_v39_, d_174_v39_])
        index5_ = default__.safeIndex(846, (d_176_v41_).length(0))
        (d_176_v41_)[index5_] = (d_177_v42_)[default__.safeIndex(default__.safeModuloInt(d_137_v8_, d_137_v8_), len(d_177_v42_))]
        d_137_v8_ = -556
        d_178_v43_: C4
        nw13_ = C4()
        nw13_.ctor__((d_174_v39_).f9, (d_174_v39_).f10)
        d_178_v43_ = nw13_
        d_179_v44_: D30
        d_179_v44_ = D30_DC73(d_178_v43_)
        d_178_v43_ = (d_178_v43_ if (d_131_v3_).f12 else (d_179_v44_).cf127)
        d_180_v45_: _dafny.Seq
        d_180_v45_ = _dafny.SeqWithoutIsStrInference([d_137_v8_])
        d_181_v46_: D1
        d_181_v46_ = D1_DC2(d_180_v45_)
        pat_let_tv0_ = d_131_v3_
        pat_let_tv1_ = d_137_v8_
        pat_let_tv2_ = d_137_v8_
        def lambda3_(source7_):
            if source7_.is_DC3:
                d_182___mcc_h0_ = source7_.cf5
                d_183_cf5_ = d_182___mcc_h0_
                return 138
            elif source7_.is_DC4:
                d_184___mcc_h1_ = source7_.cf6
                d_185___mcc_h2_ = source7_.cf7
                d_186___mcc_h3_ = source7_.cf8
                d_187_cf8_ = d_186___mcc_h3_
                d_188_cf7_ = d_185___mcc_h2_
                d_189_cf6_ = d_184___mcc_h1_
                return len(_dafny.SeqWithoutIsStrInference([d_188_cf7_, not(d_188_cf7_), (pat_let_tv0_).f12]))
            elif source7_.is_DC2:
                d_190___mcc_h4_ = source7_.cf4
                d_191_cf4_ = d_190___mcc_h4_
                return 558
            elif True:
                d_192___mcc_h5_ = source7_.cf9
                d_193_cf9_ = d_192___mcc_h5_
                return default__.safeModuloInt(pat_let_tv1_, pat_let_tv2_)

        d_137_v8_ = lambda3_(d_181_v46_)
        index6_ = default__.safeIndex(701, (d_126_v0_).length(0))
        (d_126_v0_)[index6_] = (d_131_v3_).fm32(len((d_131_v3_).f13), d_172_v37_, d_130_globalState_)
        index7_ = default__.safeIndex(701, (d_126_v0_).length(0))
        (d_126_v0_)[index7_] = (d_131_v3_).f12
        if not (d_128_v1_) or ((d_174_v39_).f10):
            if (d_174_v39_).f10:
                d_194_v47_: _dafny.Array
                nw14_ = _dafny.Array(int(0), 26)
                d_194_v47_ = nw14_
                d_195_v48_: _dafny.Map
                d_195_v48_ = _dafny.Map({d_194_v47_: d_137_v8_})
                d_196_v49_: _dafny.Map
                d_196_v49_ = d_195_v48_
                d_197_v50_: _dafny.Set
                d_197_v50_ = _dafny.Set({d_196_v49_})
                d_198_v51_: _dafny.Seq
                d_198_v51_ = _dafny.SeqWithoutIsStrInference([d_197_v50_])
                d_199_v52_: _dafny.Array
                nw15_ = _dafny.Array(None, 2)
                nw15_[int(0)] = _dafny.SeqWithoutIsStrInference([d_197_v50_])
                nw15_[int(1)] = d_198_v51_
                d_199_v52_ = nw15_
                index8_ = default__.safeIndex(199, (d_199_v52_).length(0))
                (d_199_v52_)[index8_] = d_198_v51_
                index9_ = default__.safeIndex(133, (d_194_v47_).length(0))
                (d_194_v47_)[index9_] = d_137_v8_
                index10_ = default__.safeIndex(199, (d_199_v52_).length(0))
                index11_ = default__.safeIndex(133, (d_194_v47_).length(0))
                index12_ = default__.safeIndex(701, (d_126_v0_).length(0))
                rhs9_ = (d_198_v51_) + (d_198_v51_)
                rhs10_ = (d_137_v8_ if (d_137_v8_) in (_dafny.Map({d_137_v8_: d_172_v37_})) else d_137_v8_)
                rhs11_ = default__.safeModuloInt((d_180_v45_)[default__.safeIndex(len(_dafny.Map({d_137_v8_: d_137_v8_})), len(d_180_v45_))], (D8_DC24(True, _dafny.Set({d_194_v47_, d_194_v47_, d_194_v47_}), d_137_v8_)).cf44)
                rhs12_ = d_128_v1_
                lhs3_ = d_199_v52_
                lhs4_ = default__.safeIndex(199, (d_199_v52_).length(0))
                lhs5_ = d_194_v47_
                lhs6_ = default__.safeIndex(133, (d_194_v47_).length(0))
                lhs7_ = d_126_v0_
                lhs8_ = default__.safeIndex(701, (d_126_v0_).length(0))
                lhs3_[lhs4_] = rhs9_
                lhs5_[lhs6_] = rhs10_
                d_137_v8_ = rhs11_
                lhs7_[lhs8_] = rhs12_
                d_200_v53_: _dafny.Array
                def lambda4_(d_201_v3_):
                    def lambda5_(d_202_i7_):
                        return (d_201_v3_).f13

                    return lambda5_

                init2_ = lambda4_(d_131_v3_)
                nw16_ = _dafny.Array(None, 17)
                for i0_2_ in range(nw16_.length(0)):
                    nw16_[i0_2_] = init2_(i0_2_)
                d_200_v53_ = nw16_
                index13_ = default__.safeIndex(198, (d_200_v53_).length(0))
                (d_200_v53_)[index13_] = (d_131_v3_).f13
                index14_ = default__.safeIndex(198, (d_200_v53_).length(0))
                (d_200_v53_)[index14_] = (d_131_v3_).f13
                d_203_v54_: D30
                d_203_v54_ = D30_DC74((d_131_v3_).f13)
                d_204_v55_: D30
                d_204_v55_ = D30_DC76(D30_DC76(d_203_v54_))
                pat_let_tv3_ = d_203_v54_
                d_205_v56_: _dafny.Array
                nw17_ = _dafny.Array(None, 15)
                nw17_[int(0)] = d_204_v55_
                nw17_[int(1)] = d_204_v55_
                nw17_[int(2)] = d_204_v55_
                nw17_[int(3)] = d_204_v55_
                nw17_[int(4)] = d_204_v55_
                nw17_[int(5)] = d_204_v55_
                nw17_[int(6)] = d_204_v55_
                nw17_[int(7)] = d_204_v55_
                nw17_[int(8)] = d_204_v55_
                nw17_[int(9)] = d_204_v55_
                nw17_[int(10)] = d_204_v55_
                def iife22_(_pat_let0_0):
                    def iife23_(d_206_dt__update__tmp_h0_):
                        def iife24_(_pat_let1_0):
                            def iife25_(d_207_dt__update_hcf130_h0_):
                                return D30_DC76(d_207_dt__update_hcf130_h0_)
                            return iife25_(_pat_let1_0)
                        return iife24_(pat_let_tv3_)
                    return iife23_(_pat_let0_0)
                nw17_[int(11)] = iife22_(d_204_v55_)
                nw17_[int(12)] = d_204_v55_
                nw17_[int(13)] = d_204_v55_
                nw17_[int(14)] = d_204_v55_
                d_205_v56_ = nw17_
                index15_ = default__.safeIndex(307, (d_205_v56_).length(0))
                (d_205_v56_)[index15_] = d_204_v55_
                d_208_v57_: _dafny.Seq
                d_208_v57_ = _dafny.SeqWithoutIsStrInference([(d_174_v39_).f10])
                d_209_v58_: _dafny.MultiSet
                d_209_v58_ = _dafny.MultiSet([not((d_131_v3_).f12), not((d_174_v39_).fm0(d_208_v57_, (d_174_v39_).f10, d_137_v8_, False, d_130_globalState_)), (d_131_v3_).f12])
                d_210_v59_: C2
                nw18_ = C2()
                nw18_.ctor__((d_131_v3_).f12, d_126_v0_, (d_174_v39_).f9, (d_174_v39_).f10)
                d_210_v59_ = nw18_
                d_211_v60_: _dafny.Map
                d_211_v60_ = _dafny.Map({955: d_210_v59_})
                d_212_v61_: _dafny.MultiSet
                d_212_v61_ = _dafny.MultiSet([len(d_211_v60_), (default__.fm15(True, (d_174_v39_).f10, d_137_v8_, d_172_v37_, d_130_globalState_)).cardinality, (d_194_v47_)[default__.safeIndex(133, (d_194_v47_).length(0))]])
                d_213_v62_: _dafny.Seq
                d_213_v62_ = _dafny.SeqWithoutIsStrInference([d_131_v3_])
                d_214_v63_: _dafny.Map
                d_214_v63_ = _dafny.Map({(d_212_v61_).intersection(default__.fm46((d_174_v39_).f10, (d_131_v3_).f12, (d_131_v3_).f12, d_130_globalState_)): d_213_v62_})
                d_215_v64_: _dafny.MultiSet
                d_215_v64_ = d_212_v61_
                index16_ = default__.safeIndex(307, (d_205_v56_).length(0))
                rhs13_ = not (d_128_v1_) or ((d_126_v0_)[default__.safeIndex(701, (d_126_v0_).length(0))])
                rhs14_ = d_204_v55_
                rhs15_ = (d_131_v3_).f12
                rhs16_ = d_209_v58_
                rhs17_ = (_dafny.Map({(d_215_v64_): _dafny.SeqWithoutIsStrInference([d_131_v3_])})) | (_dafny.Map({d_212_v61_: d_213_v62_}))
                lhs9_ = d_130_globalState_
                lhs10_ = d_205_v56_
                lhs11_ = default__.safeIndex(307, (d_205_v56_).length(0))
                lhs9_.f8 = rhs13_
                lhs10_[lhs11_] = rhs14_
                d_128_v1_ = rhs15_
                d_209_v58_ = rhs16_
                d_214_v63_ = rhs17_
                d_216_v65_: _dafny.Map
                d_216_v65_ = _dafny.Map({865: len(((d_200_v53_)[default__.safeIndex(198, (d_200_v53_).length(0))]).set(default__.safeIndex(d_137_v8_, len((d_200_v53_)[default__.safeIndex(198, (d_200_v53_).length(0))])), d_172_v37_))})
                d_217_v66_: _dafny.Map
                d_217_v66_ = _dafny.Map({len((d_131_v3_).f13): ((d_216_v65_)[d_137_v8_] if (d_137_v8_) in (d_216_v65_) else d_137_v8_)})
                d_218_v67_: _dafny.Seq
                d_218_v67_ = _dafny.SeqWithoutIsStrInference([(d_217_v66_) | (d_216_v65_)])
                d_218_v67_ = d_218_v67_
                index17_ = default__.safeIndex(133, (d_194_v47_).length(0))
                (d_194_v47_)[index17_] = (d_194_v47_)[default__.safeIndex(133, (d_194_v47_).length(0))]
            elif True:
                d_219_v68_: _dafny.Map
                d_219_v68_ = _dafny.Map({d_131_v3_: d_126_v0_})
                d_220_v69_: _dafny.Array
                nw19_ = _dafny.Array(None, 27)
                nw19_[int(0)] = d_126_v0_
                nw19_[int(1)] = d_126_v0_
                nw19_[int(2)] = d_126_v0_
                nw19_[int(3)] = d_126_v0_
                nw19_[int(4)] = d_126_v0_
                nw19_[int(5)] = d_126_v0_
                nw19_[int(6)] = d_126_v0_
                nw19_[int(7)] = d_126_v0_
                nw19_[int(8)] = d_126_v0_
                nw19_[int(9)] = d_126_v0_
                nw19_[int(10)] = d_126_v0_
                nw19_[int(11)] = d_126_v0_
                nw19_[int(12)] = d_126_v0_
                nw19_[int(13)] = d_126_v0_
                nw19_[int(14)] = d_126_v0_
                nw19_[int(15)] = d_126_v0_
                nw19_[int(16)] = ((d_219_v68_)[d_131_v3_] if (d_131_v3_) in (d_219_v68_) else d_126_v0_)
                nw19_[int(17)] = d_126_v0_
                nw19_[int(18)] = d_126_v0_
                nw19_[int(19)] = d_126_v0_
                nw19_[int(20)] = d_126_v0_
                nw19_[int(21)] = d_126_v0_
                nw19_[int(22)] = (d_126_v0_ if (d_174_v39_).f10 else d_126_v0_)
                nw19_[int(23)] = d_126_v0_
                nw19_[int(24)] = d_126_v0_
                nw19_[int(25)] = d_126_v0_
                nw19_[int(26)] = d_126_v0_
                d_220_v69_ = nw19_
                d_221_v70_: D32
                d_221_v70_ = D32_DC78(d_220_v69_)
                d_220_v69_ = (d_221_v70_).cf132
                (d_130_globalState_).f8 = True
                d_129_v2_ = (d_129_v2_).set(len((d_131_v3_).f13), (d_131_v3_).f12)
                d_222_v71_: _dafny.Array
                nw20_ = _dafny.Array(int(0), 19)
                d_222_v71_ = nw20_
                index18_ = default__.safeIndex(126, (d_222_v71_).length(0))
                (d_222_v71_)[index18_] = (0) - (d_137_v8_)
                index19_ = default__.safeIndex(126, (d_222_v71_).length(0))
                (d_222_v71_)[index19_] = 299
                d_223_v72_: _dafny.MultiSet
                d_223_v72_ = _dafny.MultiSet([(d_222_v71_)[default__.safeIndex(126, (d_222_v71_).length(0))], (d_222_v71_)[default__.safeIndex(126, (d_222_v71_).length(0))], (d_222_v71_)[default__.safeIndex(126, (d_222_v71_).length(0))], 178])
                d_224_v73_: _dafny.Set
                d_224_v73_ = _dafny.Set({d_137_v8_, ((d_223_v72_).set(len((d_131_v3_).f13), default__.abs((d_222_v71_)[default__.safeIndex(126, (d_222_v71_).length(0))]))).cardinality})
                d_224_v73_ = (d_224_v73_) | (d_224_v73_)
            d_225_v74_: _dafny.Seq
            d_226_v75_: bool
            d_227_v76_: _dafny.Seq
            d_228_v77_: _dafny.Seq
            out0_: _dafny.Seq
            out1_: bool
            out2_: _dafny.Seq
            out3_: _dafny.Seq
            out0_, out1_, out2_, out3_ = (d_178_v43_).m0(d_130_globalState_)
            d_225_v74_ = out0_
            d_226_v75_ = out1_
            d_227_v76_ = out2_
            d_228_v77_ = out3_
            d_229_v78_: _dafny.Seq
            d_230_v79_: bool
            d_231_v80_: _dafny.Seq
            d_232_v81_: _dafny.Seq
            out4_: _dafny.Seq
            out5_: bool
            out6_: _dafny.Seq
            out7_: _dafny.Seq
            out4_, out5_, out6_, out7_ = (d_178_v43_).m0(d_130_globalState_)
            d_229_v78_ = out4_
            d_230_v79_ = out5_
            d_231_v80_ = out6_
            d_232_v81_ = out7_
            (d_178_v43_).m5(d_130_globalState_)
            d_233_v82_: _dafny.Set
            d_233_v82_ = _dafny.Set({default__.fm37(d_130_globalState_)})
            d_233_v82_ = d_233_v82_
        elif True:
            (d_130_globalState_).f8 = d_128_v1_
            d_234_v83_: _dafny.Array
            def lambda6_(d_235_i8_):
                return _dafny.CodePoint('n')

            init3_ = lambda6_
            nw21_ = _dafny.Array(None, 3)
            for i0_3_ in range(nw21_.length(0)):
                nw21_[i0_3_] = init3_(i0_3_)
            d_234_v83_ = nw21_
            d_236_v84_: _dafny.Array
            nw22_ = _dafny.Array(int(0), 2)
            d_236_v84_ = nw22_
            index20_ = default__.safeIndex(825, (d_236_v84_).length(0))
            (d_236_v84_)[index20_] = d_137_v8_
            index21_ = default__.safeIndex(701, (d_126_v0_).length(0))
            index22_ = default__.safeIndex(825, (d_236_v84_).length(0))
            rhs18_ = (d_137_v8_) - (d_137_v8_)
            rhs19_ = (d_174_v39_).f10
            rhs20_ = d_234_v83_
            rhs21_ = d_137_v8_
            rhs22_ = d_137_v8_
            lhs12_ = d_126_v0_
            lhs13_ = default__.safeIndex(701, (d_126_v0_).length(0))
            lhs14_ = d_236_v84_
            lhs15_ = default__.safeIndex(825, (d_236_v84_).length(0))
            d_137_v8_ = rhs18_
            lhs12_[lhs13_] = rhs19_
            d_234_v83_ = rhs20_
            d_137_v8_ = rhs21_
            lhs14_[lhs15_] = rhs22_
            d_237_v85_: _dafny.Seq
            d_237_v85_ = _dafny.SeqWithoutIsStrInference([(D3_DC7(d_236_v84_)).cf11, d_236_v84_])
            d_238_v86_: _dafny.Set
            d_238_v86_ = _dafny.Set({(d_237_v85_)[default__.safeIndex(d_137_v8_, len(d_237_v85_))], d_236_v84_, d_236_v84_})
            if not((d_238_v86_).isdisjoint(d_238_v86_)):
                d_128_v1_ = (((d_174_v39_).f9)[(d_174_v39_).f10] if ((d_174_v39_).f10) in ((d_174_v39_).f9) else ((0) - ((d_236_v84_)[default__.safeIndex(825, (d_236_v84_).length(0))])) > (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nlwuecg")))))
                d_137_v8_ = (0) - (d_137_v8_)
                index23_ = default__.safeIndex(825, (d_236_v84_).length(0))
                (d_236_v84_)[index23_] = len(_dafny.SeqWithoutIsStrInference([d_172_v37_ for d_239_i9_ in range(default__.abs(19))]))
                d_240_v87_: _dafny.Map
                d_240_v87_ = _dafny.Map({not(d_128_v1_): d_137_v8_})
                d_241_v88_: _dafny.Set
                d_241_v88_ = _dafny.Set({(d_131_v3_).f13})
                d_137_v8_ = (0) - ((default__.safeDivisionInt(len(d_240_v87_), (d_236_v84_)[default__.safeIndex(825, (d_236_v84_).length(0))])) + (len(d_241_v88_)))
                d_137_v8_ = default__.safeDivisionInt(d_137_v8_, (d_137_v8_) * (d_137_v8_))
            elif True:
                d_242_v89_: _dafny.Map
                d_242_v89_ = _dafny.Map({d_137_v8_: d_180_v45_})
                d_243_v91_: _dafny.MultiSet
                d_243_v91_ = _dafny.MultiSet([(d_131_v3_).f12, (d_174_v39_).f10, (d_131_v3_).f12])
                d_244_v92_: _dafny.Set
                d_244_v92_ = _dafny.Set({(d_243_v91_).cardinality, d_137_v8_})
                d_245_v96_: _dafny.MultiSet
                d_245_v96_ = _dafny.MultiSet([(d_236_v84_)[default__.safeIndex(825, (d_236_v84_).length(0))]])
                d_246_v98_: _dafny.Array
                nw23_ = _dafny.Array(None, 16)
                nw23_[int(0)] = d_242_v89_
                def iife26_():
                    coll22_ = _dafny.Map()
                    compr_22_: int
                    for compr_22_ in (d_244_v92_).Elements:
                        d_247_v90_: int = compr_22_
                        if (d_247_v90_) in (d_244_v92_):
                            coll22_[(d_247_v90_) - (d_137_v8_)] = d_180_v45_
                    return _dafny.Map(coll22_)
                nw23_[int(1)] = (iife26_()
                ) | (d_242_v89_)
                nw23_[int(2)] = d_242_v89_
                nw23_[int(3)] = _dafny.Map({380: d_180_v45_})
                nw23_[int(4)] = d_242_v89_
                nw23_[int(5)] = (d_242_v89_) | ((D33_DC81(d_242_v89_)).cf139)
                def iife27_():
                    coll23_ = _dafny.Map()
                    compr_23_: int
                    for compr_23_ in (d_180_v45_).Elements:
                        d_248_v93_: int = compr_23_
                        if (d_248_v93_) in (d_180_v45_):
                            coll23_[default__.safeModuloInt(d_248_v93_, (d_236_v84_)[default__.safeIndex(825, (d_236_v84_).length(0))])] = d_180_v45_
                    return _dafny.Map(coll23_)
                nw23_[int(6)] = (iife27_()
                ) | (d_242_v89_)
                def iife28_():
                    coll24_ = _dafny.Map()
                    compr_24_: int
                    for compr_24_ in _dafny.IntegerRange(-33, -393):
                        d_249_v94_: int = compr_24_
                        if ((-33) <= (d_249_v94_)) and ((d_249_v94_) < (-393)):
                            coll24_[(d_249_v94_) + (len(_dafny.SeqWithoutIsStrInference([d_172_v37_ for d_250_i10_ in range(default__.abs(603))])))] = d_180_v45_
                    return _dafny.Map(coll24_)
                nw23_[int(7)] = iife28_()
                
                nw23_[int(8)] = d_242_v89_
                nw23_[int(9)] = d_242_v89_
                def iife29_():
                    coll25_ = _dafny.Map()
                    compr_25_: int
                    for compr_25_ in (d_245_v96_).Elements:
                        d_251_v95_: int = compr_25_
                        if (d_251_v95_) in (d_245_v96_):
                            coll25_[(d_251_v95_) * ((d_236_v84_)[default__.safeIndex(825, (d_236_v84_).length(0))])] = _dafny.SeqWithoutIsStrInference([d_137_v8_, (d_236_v84_)[default__.safeIndex(825, (d_236_v84_).length(0))]])
                    return _dafny.Map(coll25_)
                nw23_[int(10)] = iife29_()
                
                nw23_[int(11)] = d_242_v89_
                nw23_[int(12)] = d_242_v89_
                def iife30_():
                    coll26_ = _dafny.Map()
                    compr_26_: int
                    for compr_26_ in (d_245_v96_).Elements:
                        d_252_v97_: int = compr_26_
                        if (d_252_v97_) in (d_245_v96_):
                            coll26_[default__.safeModuloInt(d_252_v97_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fvnf"))))] = d_180_v45_
                    return _dafny.Map(coll26_)
                nw23_[int(13)] = iife30_()
                
                nw23_[int(14)] = d_242_v89_
                nw23_[int(15)] = d_242_v89_
                d_246_v98_ = nw23_
                index24_ = default__.safeIndex(633, (d_246_v98_).length(0))
                (d_246_v98_)[index24_] = d_242_v89_
                index25_ = default__.safeIndex(633, (d_246_v98_).length(0))
                (d_246_v98_)[index25_] = ((d_242_v89_) | (d_242_v89_) if False else (_dafny.Map({d_137_v8_: d_180_v45_})).set((d_236_v84_)[default__.safeIndex(825, (d_236_v84_).length(0))], (d_180_v45_).set(default__.safeIndex((d_236_v84_)[default__.safeIndex(825, (d_236_v84_).length(0))], len(d_180_v45_)), len(d_177_v42_))))
                rhs23_ = (0) - (d_137_v8_)
                rhs24_ = not((d_131_v3_).f12)
                rhs25_ = d_236_v84_
                lhs16_ = d_130_globalState_
                d_137_v8_ = rhs23_
                lhs16_.f8 = rhs24_
                d_236_v84_ = rhs25_
                index26_ = default__.safeIndex(701, (d_126_v0_).length(0))
                (d_126_v0_)[index26_] = (d_174_v39_).f10
                (d_130_globalState_).f8 = (d_174_v39_).f10
                d_253_v99_: _dafny.Set
                d_253_v99_ = _dafny.Set({(d_174_v39_).f10})
                d_254_v100_: C7
                nw24_ = C7()
                nw24_.ctor__(d_253_v99_, (d_174_v39_).f9, (d_131_v3_).fm32((d_236_v84_)[default__.safeIndex(825, (d_236_v84_).length(0))], d_172_v37_, d_130_globalState_))
                d_254_v100_ = nw24_
            d_255_v101_: _dafny.Seq
            d_255_v101_ = _dafny.SeqWithoutIsStrInference([(d_126_v0_)[default__.safeIndex(701, (d_126_v0_).length(0))]])
            d_256_v102_: _dafny.Seq
            d_256_v102_ = _dafny.SeqWithoutIsStrInference([(d_131_v3_).f13])
            d_137_v8_ = (default__.safeModuloInt(len(d_255_v101_), d_137_v8_)) + (len(d_256_v102_))
            (d_130_globalState_).f8 = (d_131_v3_).f12
        if (d_126_v0_)[default__.safeIndex(701, (d_126_v0_).length(0))]:
            d_137_v8_ = (d_137_v8_) - (-444)
            d_172_v37_ = d_172_v37_
            d_257_v103_: _dafny.Seq
            d_257_v103_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "elcmh"))
            index27_ = default__.safeIndex(701, (d_126_v0_).length(0))
            rhs26_ = not((645) < (d_137_v8_))
            rhs27_ = ((d_131_v3_).f13) + ((d_131_v3_).f13)
            lhs17_ = d_126_v0_
            lhs18_ = default__.safeIndex(701, (d_126_v0_).length(0))
            lhs17_[lhs18_] = rhs26_
            d_257_v103_ = rhs27_
            (d_178_v43_).m5(d_130_globalState_)
            d_258_v104_: _dafny.Seq
            d_258_v104_ = _dafny.SeqWithoutIsStrInference([d_129_v2_])
            d_258_v104_ = _dafny.SeqWithoutIsStrInference([(d_129_v2_) | (d_129_v2_) for d_259_i11_ in range(default__.abs(395))])
        elif True:
            d_137_v8_ = (d_137_v8_) - (d_137_v8_)
            d_260_v105_: C0
            nw25_ = C0()
            nw25_.ctor__()
            d_260_v105_ = nw25_
            d_261_v106_: _dafny.Seq
            d_261_v106_ = _dafny.SeqWithoutIsStrInference([d_126_v0_])
            d_262_v107_: _dafny.Array
            nw26_ = _dafny.Array(None, 26)
            nw26_[int(0)] = d_126_v0_
            nw26_[int(1)] = d_126_v0_
            nw26_[int(2)] = d_126_v0_
            nw26_[int(3)] = d_126_v0_
            nw26_[int(4)] = d_126_v0_
            nw26_[int(5)] = d_126_v0_
            nw26_[int(6)] = d_126_v0_
            nw26_[int(7)] = d_126_v0_
            nw26_[int(8)] = d_126_v0_
            nw26_[int(9)] = d_126_v0_
            nw26_[int(10)] = d_126_v0_
            nw26_[int(11)] = d_126_v0_
            nw26_[int(12)] = d_126_v0_
            nw26_[int(13)] = d_126_v0_
            nw26_[int(14)] = d_126_v0_
            nw26_[int(15)] = (d_261_v106_)[default__.safeIndex((0) - (d_137_v8_), len(d_261_v106_))]
            nw26_[int(16)] = d_126_v0_
            nw26_[int(17)] = d_126_v0_
            nw26_[int(18)] = d_126_v0_
            nw26_[int(19)] = d_126_v0_
            nw26_[int(20)] = d_126_v0_
            nw26_[int(21)] = d_126_v0_
            nw26_[int(22)] = d_126_v0_
            nw26_[int(23)] = d_126_v0_
            nw26_[int(24)] = d_126_v0_
            nw26_[int(25)] = d_126_v0_
            d_262_v107_ = nw26_
            index28_ = default__.safeIndex(934, (d_262_v107_).length(0))
            (d_262_v107_)[index28_] = d_126_v0_
            index29_ = default__.safeIndex(934, (d_262_v107_).length(0))
            (d_262_v107_)[index29_] = d_126_v0_
            d_263_v108_: T0
            nw27_ = C1()
            nw27_.ctor__((d_131_v3_).f12, (d_131_v3_).f13)
            d_263_v108_ = nw27_
            d_263_v108_ = d_263_v108_
            d_264_v109_: _dafny.Array
            nw28_ = _dafny.Array(_dafny.Seq({}), 22)
            d_264_v109_ = nw28_
            index30_ = default__.safeIndex(974, (d_264_v109_).length(0))
            (d_264_v109_)[index30_] = d_261_v106_
            index31_ = default__.safeIndex(974, (d_264_v109_).length(0))
            (d_264_v109_)[index31_] = (_dafny.SeqWithoutIsStrInference([(d_262_v107_)[default__.safeIndex(934, (d_262_v107_).length(0))]])) + ((d_261_v106_) + (d_261_v106_))
        (d_130_globalState_).f8 = not((d_174_v39_).f10)
        d_265_v110_: _dafny.MultiSet
        d_265_v110_ = _dafny.MultiSet([(d_131_v3_).f12, (d_131_v3_).f12])
        d_266_v111_: T0
        nw29_ = C1()
        nw29_.ctor__((948) != ((d_265_v110_).cardinality), ((d_131_v3_).f13) + ((d_131_v3_).f13))
        d_266_v111_ = nw29_
        d_267_v112_: C5
        nw30_ = C5()
        nw30_.ctor__(_dafny.Map({d_128_v1_: True}), (d_126_v0_)[default__.safeIndex(701, (d_126_v0_).length(0))])
        d_267_v112_ = nw30_
        d_268_v113_: _dafny.Map
        d_268_v113_ = _dafny.Map({d_267_v112_: d_137_v8_})
        index32_ = default__.safeIndex(701, (d_126_v0_).length(0))
        rhs28_ = (d_131_v3_).fm32(d_137_v8_, d_172_v37_, d_130_globalState_)
        rhs29_ = d_266_v111_
        rhs30_ = d_128_v1_
        rhs31_ = default__.safeModuloInt(default__.safeModuloInt(d_137_v8_, ((d_268_v113_)[d_267_v112_] if (d_267_v112_) in (d_268_v113_) else d_137_v8_)), d_137_v8_)
        lhs19_ = d_130_globalState_
        lhs20_ = d_126_v0_
        lhs21_ = default__.safeIndex(701, (d_126_v0_).length(0))
        lhs19_.f8 = rhs28_
        d_266_v111_ = rhs29_
        lhs20_[lhs21_] = rhs30_
        d_137_v8_ = rhs31_
        if (d_174_v39_).f10:
            index33_ = default__.safeIndex(701, (d_126_v0_).length(0))
            (d_126_v0_)[index33_] = (d_126_v0_)[default__.safeIndex(701, (d_126_v0_).length(0))]
            d_269_v114_: _dafny.Map
            d_269_v114_ = _dafny.Map({d_128_v1_: _dafny.MultiSet(d_180_v45_)})
            d_270_v115_: _dafny.MultiSet
            d_270_v115_ = _dafny.MultiSet([d_137_v8_, len(_dafny.SeqWithoutIsStrInference([(d_174_v39_).f10, (d_131_v3_).f12]))])
            d_271_v116_: _dafny.Seq
            d_271_v116_ = _dafny.SeqWithoutIsStrInference([d_270_v115_, d_270_v115_, (d_270_v115_).set(d_137_v8_, default__.abs(d_137_v8_)), d_270_v115_, d_270_v115_])
            (d_130_globalState_).f8 = (((d_126_v0_)[default__.safeIndex(701, (d_126_v0_).length(0))] if (d_126_v0_)[default__.safeIndex(701, (d_126_v0_).length(0))] else True)) in ((d_269_v114_).set((d_126_v0_)[default__.safeIndex(701, (d_126_v0_).length(0))], (d_271_v116_)[default__.safeIndex(d_137_v8_, len(d_271_v116_))]))
            d_272_v117_: _dafny.Array
            nw31_ = _dafny.Array(_dafny.Array(None, 0), 24)
            d_272_v117_ = nw31_
            d_273_v118_: D32
            d_273_v118_ = D32_DC78(d_272_v117_)
            source8_ = d_273_v118_
            if source8_.is_DC79:
                d_274___mcc_h6_ = source8_.cf133
                d_275___mcc_h7_ = source8_.cf134
                d_276___mcc_h8_ = source8_.cf135
                d_277_cf135_ = d_276___mcc_h8_
                d_278_cf134_ = d_275___mcc_h7_
                d_279_cf133_ = d_274___mcc_h6_
                d_280_v119_: _dafny.Seq
                d_280_v119_ = _dafny.SeqWithoutIsStrInference([d_128_v1_])
                d_281_v120_: _dafny.Set
                d_281_v120_ = _dafny.Set({(d_131_v3_).f12, (d_174_v39_).fm0(d_280_v119_, (d_131_v3_).f12, d_137_v8_, (d_126_v0_)[default__.safeIndex(701, (d_126_v0_).length(0))], d_130_globalState_)})
                d_281_v120_ = (default__.fm34(d_137_v8_, False, (d_174_v39_).f10, True, d_130_globalState_)) | (d_281_v120_)
                d_282_v121_: D0
                d_282_v121_ = D0_DC1(d_137_v8_, (d_126_v0_)[default__.safeIndex(701, (d_126_v0_).length(0))], d_137_v8_)
                d_283_v122_: _dafny.Seq
                d_283_v122_ = _dafny.SeqWithoutIsStrInference([d_266_v111_, d_266_v111_])
                d_284_v123_: _dafny.Seq
                d_284_v123_ = _dafny.SeqWithoutIsStrInference([d_283_v122_])
                d_285_v124_: _dafny.Map
                d_285_v124_ = _dafny.Map({d_137_v8_: d_172_v37_})
                d_286_v125_: _dafny.Array
                nw32_ = _dafny.Array(None, 18)
                nw32_[int(0)] = d_137_v8_
                nw32_[int(1)] = d_137_v8_
                nw32_[int(2)] = d_137_v8_
                nw32_[int(3)] = d_137_v8_
                nw32_[int(4)] = len(d_284_v123_)
                nw32_[int(5)] = len(d_180_v45_)
                nw32_[int(6)] = len(d_281_v120_)
                nw32_[int(7)] = d_137_v8_
                nw32_[int(8)] = d_137_v8_
                nw32_[int(9)] = len(d_280_v119_)
                nw32_[int(10)] = d_137_v8_
                nw32_[int(11)] = (0) - (d_137_v8_)
                nw32_[int(12)] = d_137_v8_
                nw32_[int(13)] = len(d_285_v124_)
                nw32_[int(14)] = (0) - (d_137_v8_)
                nw32_[int(15)] = d_137_v8_
                nw32_[int(16)] = d_137_v8_
                nw32_[int(17)] = d_137_v8_
                d_286_v125_ = nw32_
                d_287_v126_: _dafny.Set
                d_287_v126_ = _dafny.Set({d_286_v125_, d_286_v125_, d_286_v125_, d_286_v125_, d_286_v125_})
                d_288_v127_: D8
                d_288_v127_ = D8_DC24((d_282_v121_).cf2, (d_287_v126_).intersection(d_287_v126_), (d_267_v112_).fm1((d_174_v39_).f9, d_130_globalState_))
                d_288_v127_ = d_288_v127_
                d_172_v37_ = d_172_v37_
                d_289_v128_: _dafny.Set
                d_289_v128_ = _dafny.Set({921, d_137_v8_})
                index34_ = default__.safeIndex(701, (d_126_v0_).length(0))
                (d_126_v0_)[index34_] = (d_289_v128_).issubset((d_289_v128_).intersection(d_289_v128_))
            elif source8_.is_DC80:
                d_290___mcc_h9_ = source8_.cf136
                d_291___mcc_h10_ = source8_.cf137
                d_292___mcc_h11_ = source8_.cf138
                d_293_cf138_ = d_292___mcc_h11_
                d_294_cf137_ = d_291___mcc_h10_
                d_295_cf136_ = d_290___mcc_h9_
                (d_130_globalState_).f8 = d_128_v1_
                d_272_v117_ = d_272_v117_
                d_295_cf136_ = default__.safeDivisionInt(d_293_cf138_, d_294_cf137_)
                d_296_v129_: _dafny.Set
                d_296_v129_ = _dafny.Set({(d_131_v3_).f12, (d_174_v39_).f10})
                d_297_v130_: _dafny.Map
                d_297_v130_ = _dafny.Map({len(d_296_v129_): (d_131_v3_).f13})
                d_298_v131_: _dafny.Seq
                d_298_v131_ = _dafny.SeqWithoutIsStrInference([d_129_v2_, (default__.fm11((d_270_v115_).cardinality, len(d_180_v45_), (d_174_v39_).f10, len((_dafny.Map({(d_174_v39_).f10: (d_176_v41_)[default__.safeIndex(846, (d_176_v41_).length(0))]})).set((d_131_v3_).f12, (d_176_v41_)[default__.safeIndex(846, (d_176_v41_).length(0))])), d_130_globalState_)).set(d_294_cf137_, d_128_v1_)])
                d_299_v132_: D16
                d_299_v132_ = D16_DC40(d_298_v131_)
                pat_let_tv4_ = d_298_v131_
                d_300_v133_: _dafny.Map
                def iife31_(_pat_let2_0):
                    def iife32_(d_301_dt__update__tmp_h1_):
                        def iife33_(_pat_let3_0):
                            def iife34_(d_302_dt__update_hcf76_h0_):
                                return D16_DC40(d_302_dt__update_hcf76_h0_)
                            return iife34_(_pat_let3_0)
                        return iife33_(pat_let_tv4_)
                    return iife32_(_pat_let2_0)
                d_300_v133_ = _dafny.Map({(_dafny.CodePoint('c')) not in (((d_297_v130_)[(0) - (d_294_cf137_)] if ((0) - (d_294_cf137_)) in (d_297_v130_) else (d_131_v3_).f13)): iife31_(d_299_v132_)})
                d_300_v133_ = (d_300_v133_).set((d_180_v45_) <= (_dafny.SeqWithoutIsStrInference([153])), d_299_v132_)
            elif True:
                d_303___mcc_h12_ = source8_.cf132
                d_304_cf132_ = d_303___mcc_h12_
                d_172_v37_ = d_172_v37_
                d_305_v134_: _dafny.Seq
                d_305_v134_ = _dafny.SeqWithoutIsStrInference([(d_174_v39_).f10])
                (d_130_globalState_).f8 = (d_305_v134_)[default__.safeIndex(d_137_v8_, len(d_305_v134_))]
                d_306_v135_: _dafny.Seq
                d_307_v136_: bool
                d_308_v137_: _dafny.Seq
                d_309_v138_: _dafny.Seq
                out8_: _dafny.Seq
                out9_: bool
                out10_: _dafny.Seq
                out11_: _dafny.Seq
                out8_, out9_, out10_, out11_ = (d_267_v112_).m0(d_130_globalState_)
                d_306_v135_ = out8_
                d_307_v136_ = out9_
                d_308_v137_ = out10_
                d_309_v138_ = out11_
                d_309_v138_ = (d_131_v3_).f13
            d_310_v139_: D25
            d_310_v139_ = D25_DC63((d_131_v3_).f13, d_137_v8_)
            d_311_v140_: _dafny.Set
            d_311_v140_ = _dafny.Set({(d_310_v139_).cf113})
            d_312_v141_: _dafny.Array
            nw33_ = _dafny.Array(None, 25)
            nw33_[int(0)] = d_137_v8_
            nw33_[int(1)] = d_137_v8_
            nw33_[int(2)] = d_137_v8_
            nw33_[int(3)] = d_137_v8_
            nw33_[int(4)] = d_137_v8_
            nw33_[int(5)] = d_137_v8_
            nw33_[int(6)] = d_137_v8_
            nw33_[int(7)] = d_137_v8_
            nw33_[int(8)] = d_137_v8_
            nw33_[int(9)] = d_137_v8_
            nw33_[int(10)] = d_137_v8_
            nw33_[int(11)] = d_137_v8_
            nw33_[int(12)] = d_137_v8_
            nw33_[int(13)] = d_137_v8_
            nw33_[int(14)] = d_137_v8_
            nw33_[int(15)] = len(d_311_v140_)
            nw33_[int(16)] = len((d_131_v3_).f13)
            nw33_[int(17)] = d_137_v8_
            nw33_[int(18)] = d_137_v8_
            nw33_[int(19)] = d_137_v8_
            nw33_[int(20)] = d_137_v8_
            nw33_[int(21)] = default__.fm33(d_128_v1_, (d_174_v39_).f10, d_137_v8_, (d_174_v39_).f10, d_130_globalState_)
            nw33_[int(22)] = d_137_v8_
            nw33_[int(23)] = len((d_131_v3_).f13)
            nw33_[int(24)] = d_137_v8_
            d_312_v141_ = nw33_
            d_313_v142_: _dafny.Set
            d_313_v142_ = _dafny.Set({d_312_v141_})
            d_314_v143_: D8
            d_314_v143_ = D8_DC24((d_174_v39_).f10, d_313_v142_, d_137_v8_)
            d_315_v144_: D16
            d_315_v144_ = D16_DC42(d_128_v1_, (d_314_v143_).cf44, (d_267_v112_).fm19(d_137_v8_, d_137_v8_, d_130_globalState_))
            source9_ = d_315_v144_
            if source9_.is_DC41:
                d_316___mcc_h13_ = source9_.cf77
                d_317___mcc_h14_ = source9_.cf78
                d_318___mcc_h15_ = source9_.cf79
                d_319_cf79_ = d_318___mcc_h15_
                d_320_cf78_ = d_317___mcc_h14_
                d_321_cf77_ = d_316___mcc_h13_
                d_128_v1_ = (len(((d_173_v38_).set((d_174_v39_).f10, (d_126_v0_)[default__.safeIndex(701, (d_126_v0_).length(0))])) | ((d_174_v39_).f9))) != (((0) - (d_137_v8_)) + (d_320_cf78_))
                (d_178_v43_).m5(d_130_globalState_)
                d_322_v145_: D21
                d_322_v145_ = D21_DC53(d_266_v111_)
                pat_let_tv5_ = d_266_v111_
                def iife35_(_pat_let4_0):
                    def iife36_(d_323_dt__update__tmp_h2_):
                        def iife37_(_pat_let5_0):
                            def iife38_(d_324_dt__update_hcf99_h0_):
                                return D21_DC53(d_324_dt__update_hcf99_h0_)
                            return iife38_(_pat_let5_0)
                        return iife37_(pat_let_tv5_)
                    return iife36_(_pat_let4_0)
                d_322_v145_ = iife35_(d_322_v145_)
                (d_130_globalState_).f8 = (d_172_v37_) not in ((d_131_v3_).f13)
            elif source9_.is_DC42:
                d_325___mcc_h16_ = source9_.cf80
                d_326___mcc_h17_ = source9_.cf81
                d_327___mcc_h18_ = source9_.cf82
                d_328_cf82_ = d_327___mcc_h18_
                d_329_cf81_ = d_326___mcc_h17_
                d_330_cf80_ = d_325___mcc_h16_
                d_331_v146_: _dafny.Map
                d_331_v146_ = _dafny.Map({not((d_131_v3_).f12): ((0) - (d_329_cf81_)) - (d_328_cf82_)})
                rhs32_ = (d_126_v0_)[default__.safeIndex(701, (d_126_v0_).length(0))]
                rhs33_ = d_331_v146_
                rhs34_ = not (not ((d_131_v3_).f12) or (not((d_131_v3_).f12))) or (not ((d_126_v0_)[default__.safeIndex(701, (d_126_v0_).length(0))]) or ((d_174_v39_).f10))
                lhs22_ = d_130_globalState_
                lhs23_ = d_130_globalState_
                lhs22_.f8 = rhs32_
                d_331_v146_ = rhs33_
                lhs23_.f8 = rhs34_
                d_137_v8_ = d_328_cf82_
                d_332_v147_: _dafny.Set
                d_332_v147_ = _dafny.Set({d_126_v0_})
                index35_ = default__.safeIndex(701, (d_126_v0_).length(0))
                rhs35_ = (d_174_v39_).f10
                rhs36_ = (d_328_cf82_) + ((d_178_v43_).fm1((d_174_v39_).f9, d_130_globalState_))
                rhs37_ = d_332_v147_
                lhs24_ = d_126_v0_
                lhs25_ = default__.safeIndex(701, (d_126_v0_).length(0))
                lhs24_[lhs25_] = rhs35_
                d_329_cf81_ = rhs36_
                d_332_v147_ = rhs37_
                d_333_v148_: _dafny.Seq
                d_333_v148_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kspfey"))
                d_333_v148_ = ((d_131_v3_).f13) + (_dafny.SeqWithoutIsStrInference([d_172_v37_ for d_334_i12_ in range(default__.abs(784))]))
            elif True:
                d_335___mcc_h19_ = source9_.cf76
                d_336_cf76_ = d_335___mcc_h19_
                index36_ = default__.safeIndex(701, (d_126_v0_).length(0))
                (d_126_v0_)[index36_] = (d_174_v39_).f10
                d_337_v149_: _dafny.Seq
                d_338_v150_: bool
                d_339_v151_: _dafny.Seq
                d_340_v152_: _dafny.Seq
                out12_: _dafny.Seq
                out13_: bool
                out14_: _dafny.Seq
                out15_: _dafny.Seq
                out12_, out13_, out14_, out15_ = (d_178_v43_).m0(d_130_globalState_)
                d_337_v149_ = out12_
                d_338_v150_ = out13_
                d_339_v151_ = out14_
                d_340_v152_ = out15_
                d_341_v153_: _dafny.Seq
                d_342_v154_: bool
                d_343_v155_: _dafny.Seq
                d_344_v156_: _dafny.Seq
                out16_: _dafny.Seq
                out17_: bool
                out18_: _dafny.Seq
                out19_: _dafny.Seq
                out16_, out17_, out18_, out19_ = (d_131_v3_).m0(d_130_globalState_)
                d_341_v153_ = out16_
                d_342_v154_ = out17_
                d_343_v155_ = out18_
                d_344_v156_ = out19_
                (d_130_globalState_).f8 = d_338_v150_
            d_345_v157_: _dafny.Map
            d_345_v157_ = _dafny.Map({d_137_v8_: d_137_v8_})
            d_137_v8_ = default__.safeModuloInt(d_137_v8_, default__.safeModuloInt(len(d_345_v157_), d_137_v8_))
        elif True:
            d_346_v158_: _dafny.Seq
            d_347_v159_: bool
            d_348_v160_: _dafny.Seq
            d_349_v161_: _dafny.Seq
            out20_: _dafny.Seq
            out21_: bool
            out22_: _dafny.Seq
            out23_: _dafny.Seq
            out20_, out21_, out22_, out23_ = (d_267_v112_).m0(d_130_globalState_)
            d_346_v158_ = out20_
            d_347_v159_ = out21_
            d_348_v160_ = out22_
            d_349_v161_ = out23_
            d_350_v162_: C4
            nw34_ = C4()
            nw34_.ctor__(d_173_v38_, ((d_131_v3_).f12) or ((d_174_v39_).f10))
            d_350_v162_ = nw34_
            d_351_v163_: _dafny.Map
            d_351_v163_ = _dafny.Map({d_266_v111_: d_128_v1_})
            d_352_v164_: _dafny.Map
            d_352_v164_ = _dafny.Map({d_351_v163_: d_137_v8_})
            d_353_v165_: D32
            d_353_v165_ = D32_DC80(d_137_v8_, d_137_v8_, 665)
            d_354_v166_: _dafny.MultiSet
            d_354_v166_ = _dafny.MultiSet([879])
            d_355_v167_: _dafny.Array
            nw35_ = _dafny.Array(_dafny.Array(None, 0), 3)
            d_355_v167_ = nw35_
            d_356_v168_: _dafny.Map
            d_356_v168_ = _dafny.Map({d_355_v167_: d_349_v161_})
            d_357_v169_: _dafny.Array
            nw36_ = _dafny.Array(None, 23)
            nw36_[int(0)] = default__.fm22(d_137_v8_, d_128_v1_, (d_174_v39_).f10, d_137_v8_, d_130_globalState_)
            nw36_[int(1)] = d_346_v158_
            nw36_[int(2)] = d_349_v161_
            nw36_[int(3)] = ((d_131_v3_).f13).set(default__.safeIndex(len((d_131_v3_).f13), len((d_131_v3_).f13)), d_172_v37_)
            nw36_[int(4)] = (d_131_v3_).f13
            nw36_[int(5)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yieoik"))
            nw36_[int(6)] = d_349_v161_
            nw36_[int(7)] = ((d_131_v3_).f13) + (d_346_v158_)
            nw36_[int(8)] = ((d_131_v3_).f13).set(default__.safeIndex(len(d_352_v164_), len((d_131_v3_).f13)), d_172_v37_)
            nw36_[int(9)] = default__.fm13((d_353_v165_).cf138, d_130_globalState_)
            nw36_[int(10)] = d_349_v161_
            nw36_[int(11)] = ((d_131_v3_).f13).set(default__.safeIndex(d_137_v8_, len((d_131_v3_).f13)), _dafny.CodePoint('p'))
            nw36_[int(12)] = _dafny.SeqWithoutIsStrInference([d_172_v37_ for d_358_i13_ in range(default__.abs(476))])
            nw36_[int(13)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a"))
            nw36_[int(14)] = default__.fm45((d_131_v3_).f12, ((d_354_v166_)[d_137_v8_] if (d_137_v8_) in (d_354_v166_) else d_137_v8_), (d_131_v3_).f12, d_130_globalState_)
            nw36_[int(15)] = d_346_v158_
            nw36_[int(16)] = ((d_131_v3_).f13).set(default__.safeIndex(len(d_346_v158_), len((d_131_v3_).f13)), _dafny.CodePoint('p'))
            nw36_[int(17)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iybtk"))
            nw36_[int(18)] = default__.fm45(d_347_v159_, default__.fm33(d_128_v1_, d_347_v159_, (0) - (d_137_v8_), (d_126_v0_)[default__.safeIndex(701, (d_126_v0_).length(0))], d_130_globalState_), not(not(not((d_131_v3_).f12))), d_130_globalState_)
            nw36_[int(19)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wgwipapdd")) if (d_174_v39_).f10 else (d_131_v3_).f13)
            nw36_[int(20)] = d_346_v158_
            nw36_[int(21)] = (d_346_v158_) + (d_346_v158_)
            nw36_[int(22)] = ((d_356_v168_)[d_355_v167_] if (d_355_v167_) in (d_356_v168_) else (d_131_v3_).f13)
            d_357_v169_ = nw36_
            d_359_v170_: _dafny.Map
            d_359_v170_ = _dafny.Map({d_137_v8_: d_357_v169_})
            rhs38_ = (d_350_v162_).fm1(_dafny.Map({(d_174_v39_).f10: (d_126_v0_)[default__.safeIndex(701, (d_126_v0_).length(0))]}), d_130_globalState_)
            rhs39_ = (d_131_v3_).f13
            rhs40_ = ((d_359_v170_)[len((d_131_v3_).f13)] if (len((d_131_v3_).f13)) in (d_359_v170_) else d_357_v169_)
            rhs41_ = d_128_v1_
            d_137_v8_ = rhs38_
            d_349_v161_ = rhs39_
            d_357_v169_ = rhs40_
            d_347_v159_ = rhs41_
            d_360_v171_: _dafny.Map
            d_360_v171_ = _dafny.Map({(d_131_v3_).f12: d_126_v0_})
            d_360_v171_ = _dafny.Map({not ((d_174_v39_).f10) or (d_128_v1_): d_126_v0_})
            d_128_v1_ = (d_174_v39_).f10
        d_361_v172_: _dafny.Seq
        d_361_v172_ = _dafny.SeqWithoutIsStrInference([(d_174_v39_).f10, ((d_174_v39_).f10 if d_128_v1_ else not(d_128_v1_)), (_dafny.SeqWithoutIsStrInference([d_172_v37_ for d_362_i14_ in range(default__.abs(534))])) <= ((d_131_v3_).f13)])
        if (d_361_v172_)[default__.safeIndex(d_137_v8_, len(d_361_v172_))]:
            d_363_v173_: _dafny.Map
            d_364_v174_: bool
            out24_: _dafny.Map
            out25_: bool
            out24_, out25_ = (d_266_v111_).m1(d_172_v37_, d_130_globalState_)
            d_363_v173_ = out24_
            d_364_v174_ = out25_
            (d_178_v43_).m5(d_130_globalState_)
            d_365_v175_: _dafny.Seq
            d_366_v176_: bool
            d_367_v177_: _dafny.Seq
            d_368_v178_: _dafny.Seq
            out26_: _dafny.Seq
            out27_: bool
            out28_: _dafny.Seq
            out29_: _dafny.Seq
            out26_, out27_, out28_, out29_ = (d_267_v112_).m0(d_130_globalState_)
            d_365_v175_ = out26_
            d_366_v176_ = out27_
            d_367_v177_ = out28_
            d_368_v178_ = out29_
            d_369_v179_: C0
            nw37_ = C0()
            nw37_.ctor__()
            d_369_v179_ = nw37_
            d_370_v181_: D33
            def iife39_():
                coll27_ = _dafny.Map()
                compr_27_: int
                for compr_27_ in _dafny.IntegerRange(547, 232):
                    d_371_v180_: int = compr_27_
                    if ((547) <= (d_371_v180_)) and ((d_371_v180_) < (232)):
                        coll27_[(d_371_v180_) * (len(_dafny.SeqWithoutIsStrInference([d_172_v37_ for d_372_i15_ in range(default__.abs(7))])))] = d_180_v45_
                return _dafny.Map(coll27_)
            d_370_v181_ = D33_DC81(iife39_()
)
            d_373_v182_: D33
            d_373_v182_ = D33_DC83(d_370_v181_)
            d_374_v183_: D33
            d_374_v183_ = D33_DC83(d_370_v181_)
            d_374_v183_ = d_374_v183_
        elif True:
            d_375_v184_: _dafny.Seq
            d_375_v184_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jytbikv"))
            d_375_v184_ = (((d_131_v3_).f13).set(default__.safeIndex(d_137_v8_, len((d_131_v3_).f13)), d_172_v37_)).set(default__.safeIndex(d_137_v8_, len(((d_131_v3_).f13).set(default__.safeIndex(d_137_v8_, len((d_131_v3_).f13)), d_172_v37_))), d_172_v37_)
            d_376_v185_: _dafny.MultiSet
            d_376_v185_ = _dafny.MultiSet([d_137_v8_])
            d_376_v185_ = (default__.fm61(_dafny.Map({986: d_137_v8_}), d_130_globalState_))
            index37_ = default__.safeIndex(701, (d_126_v0_).length(0))
            (d_126_v0_)[index37_] = not ((d_131_v3_).f12) or (not(((d_131_v3_).f12) == ((d_361_v172_)[default__.safeIndex(d_137_v8_, len(d_361_v172_))])))
            d_137_v8_ = (len((d_375_v184_) + (d_375_v184_))) - (len(_dafny.SeqWithoutIsStrInference([d_172_v37_ for d_377_i16_ in range(default__.abs(-69))])))
            d_378_v186_: D14
            d_378_v186_ = D14_DC36(d_128_v1_, d_375_v184_, (d_131_v3_).f13)
            source10_ = d_378_v186_
            if source10_.is_DC35:
                d_379___mcc_h20_ = source10_.cf62
                d_380___mcc_h21_ = source10_.cf63
                d_381___mcc_h22_ = source10_.cf64
                d_382_cf64_ = d_381___mcc_h22_
                d_383_cf63_ = d_380___mcc_h21_
                d_384_cf62_ = d_379___mcc_h20_
                d_385_v187_: _dafny.Map
                d_385_v187_ = _dafny.Map({d_137_v8_: d_384_cf62_})
                d_384_cf62_ = ((d_385_v187_)[(D5_DC14(len(d_375_v184_), (d_131_v3_).f12)).cf25] if ((D5_DC14(len(d_375_v184_), (d_131_v3_).f12)).cf25) in (d_385_v187_) else d_384_cf62_)
                d_384_cf62_ = d_137_v8_
                d_386_v188_: _dafny.Array
                def lambda7_(d_387_v8_):
                    def lambda8_(d_388_i17_):
                        return (d_388_i17_) - (d_387_v8_)

                    return lambda8_

                init4_ = lambda7_(d_137_v8_)
                nw38_ = _dafny.Array(None, 3)
                for i0_4_ in range(nw38_.length(0)):
                    nw38_[i0_4_] = init4_(i0_4_)
                d_386_v188_ = nw38_
                d_389_v189_: _dafny.Set
                d_389_v189_ = _dafny.Set({d_137_v8_})
                index38_ = default__.safeIndex(7, (d_386_v188_).length(0))
                (d_386_v188_)[index38_] = len(d_389_v189_)
                d_390_v190_: _dafny.Map
                d_390_v190_ = _dafny.Map({d_384_cf62_: d_180_v45_})
                index39_ = default__.safeIndex(701, (d_126_v0_).length(0))
                index40_ = default__.safeIndex(7, (d_386_v188_).length(0))
                def iife40_():
                    coll28_ = _dafny.Map()
                    compr_28_: int
                    for compr_28_ in _dafny.IntegerRange(492, 183):
                        d_391_v191_: int = compr_28_
                        if ((492) <= (d_391_v191_)) and ((d_391_v191_) < (183)):
                            coll28_[default__.safeDivisionInt(d_391_v191_, -771)] = _dafny.SeqWithoutIsStrInference([d_384_cf62_ for d_392_i18_ in range(default__.abs(996))])
                    return _dafny.Map(coll28_)
                rhs42_ = (d_131_v3_).f12
                rhs43_ = (len((d_131_v3_).f13)) not in (d_376_v185_)
                rhs44_ = d_126_v0_
                rhs45_ = (0) - (len((d_390_v190_) | (iife40_()
                )))
                lhs26_ = d_126_v0_
                lhs27_ = default__.safeIndex(701, (d_126_v0_).length(0))
                lhs28_ = d_130_globalState_
                lhs29_ = d_386_v188_
                lhs30_ = default__.safeIndex(7, (d_386_v188_).length(0))
                lhs26_[lhs27_] = rhs42_
                lhs28_.f8 = rhs43_
                d_126_v0_ = rhs44_
                lhs29_[lhs30_] = rhs45_
                d_384_cf62_ = default__.safeDivisionInt(default__.safeModuloInt(d_384_cf62_, 270), d_384_cf62_)
            elif source10_.is_DC36:
                d_393___mcc_h23_ = source10_.cf65
                d_394___mcc_h24_ = source10_.cf66
                d_395___mcc_h25_ = source10_.cf67
                d_396_cf67_ = d_395___mcc_h25_
                d_397_cf66_ = d_394___mcc_h24_
                d_398_cf65_ = d_393___mcc_h23_
                d_399_v192_: _dafny.Map
                d_400_v193_: bool
                out30_: _dafny.Map
                out31_: bool
                out30_, out31_ = (d_266_v111_).m1(d_172_v37_, d_130_globalState_)
                d_399_v192_ = out30_
                d_400_v193_ = out31_
                d_401_v194_: _dafny.Map
                d_402_v195_: bool
                out32_: _dafny.Map
                out33_: bool
                out32_, out33_ = (d_174_v39_).m1(d_172_v37_, d_130_globalState_)
                d_401_v194_ = out32_
                d_402_v195_ = out33_
                d_398_cf65_ = (d_172_v37_) in (d_397_cf66_)
                d_361_v172_ = (d_361_v172_ if (d_174_v39_).f10 else d_361_v172_)
            elif True:
                d_403___mcc_h26_ = source10_.cf61
                d_404_cf61_ = d_403___mcc_h26_
                d_405_v196_: _dafny.Map
                d_405_v196_ = _dafny.Map({d_137_v8_: ((d_131_v3_).f13) + ((d_131_v3_).f13)})
                d_405_v196_ = (d_405_v196_).set(d_137_v8_, d_375_v184_)
                d_406_v197_: _dafny.Map
                d_406_v197_ = _dafny.Map({d_172_v37_: len(_dafny.Map({d_128_v1_: d_137_v8_}))})
                d_407_v198_: _dafny.Set
                d_407_v198_ = _dafny.Set({d_137_v8_, (0) - (d_137_v8_), len((_dafny.SeqWithoutIsStrInference([(d_131_v3_).f12])).set(default__.safeIndex(len(d_406_v197_), len(_dafny.SeqWithoutIsStrInference([(d_131_v3_).f12]))), not((d_131_v3_).f12)))})
                (d_130_globalState_).f8 = (d_407_v198_).isdisjoint(d_407_v198_)
                index41_ = default__.safeIndex(701, (d_126_v0_).length(0))
                (d_126_v0_)[index41_] = (d_137_v8_) < (len(d_407_v198_))
                index42_ = default__.safeIndex(701, (d_126_v0_).length(0))
                rhs46_ = (d_131_v3_).f12
                rhs47_ = d_126_v0_
                lhs31_ = d_126_v0_
                lhs32_ = default__.safeIndex(701, (d_126_v0_).length(0))
                lhs31_[lhs32_] = rhs46_
                d_126_v0_ = rhs47_
        source11_ = D16_DC41(not (d_128_v1_) or ((d_126_v0_)[default__.safeIndex(701, (d_126_v0_).length(0))]), (0) - (d_137_v8_), (d_126_v0_)[default__.safeIndex(701, (d_126_v0_).length(0))])
        if source11_.is_DC41:
            d_408___mcc_h27_ = source11_.cf77
            d_409___mcc_h28_ = source11_.cf78
            d_410___mcc_h29_ = source11_.cf79
            d_411_cf79_ = d_410___mcc_h29_
            d_412_cf78_ = d_409___mcc_h28_
            d_413_cf77_ = d_408___mcc_h27_
            (d_130_globalState_).f8 = True
            d_414_v199_: _dafny.Seq
            d_414_v199_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_128_v1_: len((d_131_v3_).f13)}), _dafny.Map({(d_174_v39_).f10: d_137_v8_})])
            d_415_v200_: _dafny.Map
            d_415_v200_ = _dafny.Map({d_128_v1_: d_412_cf78_})
            d_137_v8_ = len(((d_414_v199_) + (_dafny.SeqWithoutIsStrInference([d_415_v200_]))) + ((d_414_v199_) + (_dafny.SeqWithoutIsStrInference([d_415_v200_]))))
            d_137_v8_ = (0) - ((0) - (d_412_cf78_))
            d_416_v201_: _dafny.Set
            d_416_v201_ = _dafny.Set({d_412_cf78_, len(d_361_v172_)})
            d_417_v202_: _dafny.Set
            d_417_v202_ = _dafny.Set({len(d_416_v201_), d_137_v8_, d_412_cf78_, d_137_v8_, d_412_cf78_})
            index43_ = default__.safeIndex(701, (d_126_v0_).length(0))
            rhs48_ = (d_417_v202_).issubset(d_416_v201_)
            rhs49_ = d_412_cf78_
            rhs50_ = d_413_cf77_
            lhs33_ = d_130_globalState_
            lhs34_ = d_126_v0_
            lhs35_ = default__.safeIndex(701, (d_126_v0_).length(0))
            lhs33_.f8 = rhs48_
            d_412_cf78_ = rhs49_
            lhs34_[lhs35_] = rhs50_
        elif source11_.is_DC42:
            d_418___mcc_h30_ = source11_.cf80
            d_419___mcc_h31_ = source11_.cf81
            d_420___mcc_h32_ = source11_.cf82
            d_421_cf82_ = d_420___mcc_h32_
            d_422_cf81_ = d_419___mcc_h31_
            d_423_cf80_ = d_418___mcc_h30_
            d_424_v203_: _dafny.Array
            nw39_ = _dafny.Array(int(0), 15)
            d_424_v203_ = nw39_
            index44_ = default__.safeIndex(552, (d_424_v203_).length(0))
            (d_424_v203_)[index44_] = d_422_cf81_
            d_425_v204_: _dafny.Array
            nw40_ = _dafny.Array(D12.default()(), 12)
            d_425_v204_ = nw40_
            index45_ = default__.safeIndex(552, (d_424_v203_).length(0))
            rhs51_ = (0) - (d_421_cf82_)
            rhs52_ = d_425_v204_
            lhs36_ = d_424_v203_
            lhs37_ = default__.safeIndex(552, (d_424_v203_).length(0))
            lhs36_[lhs37_] = rhs51_
            d_425_v204_ = rhs52_
            d_426_v205_: _dafny.Seq
            d_427_v206_: bool
            d_428_v207_: _dafny.Seq
            d_429_v208_: _dafny.Seq
            out34_: _dafny.Seq
            out35_: bool
            out36_: _dafny.Seq
            out37_: _dafny.Seq
            out34_, out35_, out36_, out37_ = (d_174_v39_).m0(d_130_globalState_)
            d_426_v205_ = out34_
            d_427_v206_ = out35_
            d_428_v207_ = out36_
            d_429_v208_ = out37_
            d_430_v209_: _dafny.Set
            d_430_v209_ = _dafny.Set({(d_174_v39_).f10, (d_131_v3_).f12})
            d_431_v210_: C7
            nw41_ = C7()
            nw41_.ctor__((d_430_v209_) | (default__.fm34(d_422_cf81_, True, (d_131_v3_).f12, (d_174_v39_).f10, d_130_globalState_)), d_173_v38_, (d_427_v206_ if d_427_v206_ else (d_126_v0_)[default__.safeIndex(701, (d_126_v0_).length(0))]))
            d_431_v210_ = nw41_
            d_432_v211_: _dafny.Array
            nw42_ = _dafny.Array(None, 25)
            d_432_v211_ = nw42_
            d_433_v212_: C6
            nw43_ = C6()
            nw43_.ctor__(d_173_v38_, False)
            d_433_v212_ = nw43_
            index46_ = default__.safeIndex(581, (d_432_v211_).length(0))
            (d_432_v211_)[index46_] = d_433_v212_
            d_434_v213_: _dafny.Array
            nw44_ = _dafny.Array(D1.default()(), 24)
            d_434_v213_ = nw44_
            d_435_v214_: _dafny.Seq
            d_435_v214_ = _dafny.SeqWithoutIsStrInference([d_434_v213_])
            d_436_v215_: D15
            d_436_v215_ = D15_DC37((d_435_v214_)[default__.safeIndex(d_137_v8_, len(d_435_v214_))])
            d_437_v216_: _dafny.Seq
            d_437_v216_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({d_137_v8_, d_137_v8_})])
            d_438_v217_: _dafny.Set
            d_438_v217_ = _dafny.Set({len(_dafny.Set({True, (d_131_v3_).f12})), d_137_v8_})
            d_439_v218_: _dafny.Set
            d_439_v218_ = _dafny.Set({d_426_v205_, (d_131_v3_).f13, (d_131_v3_).f13})
            index47_ = default__.safeIndex(581, (d_432_v211_).length(0))
            index48_ = default__.safeIndex(552, (d_424_v203_).length(0))
            index49_ = default__.safeIndex(701, (d_126_v0_).length(0))
            rhs53_ = d_431_v210_
            rhs54_ = d_433_v212_
            rhs55_ = (0) - (len((d_437_v216_) + (_dafny.SeqWithoutIsStrInference([d_438_v217_]))))
            rhs56_ = D15_DC37(d_434_v213_)
            rhs57_ = not((((d_131_v3_).f13) + ((d_131_v3_).f13)) in (d_439_v218_))
            lhs38_ = d_432_v211_
            lhs39_ = default__.safeIndex(581, (d_432_v211_).length(0))
            lhs40_ = d_424_v203_
            lhs41_ = default__.safeIndex(552, (d_424_v203_).length(0))
            lhs42_ = d_126_v0_
            lhs43_ = default__.safeIndex(701, (d_126_v0_).length(0))
            d_431_v210_ = rhs53_
            lhs38_[lhs39_] = rhs54_
            lhs40_[lhs41_] = rhs55_
            d_436_v215_ = rhs56_
            lhs42_[lhs43_] = rhs57_
            d_423_cf80_ = (d_126_v0_)[default__.safeIndex(701, (d_126_v0_).length(0))]
        elif True:
            d_440___mcc_h33_ = source11_.cf76
            d_441_cf76_ = d_440___mcc_h33_
            d_180_v45_ = d_180_v45_
            d_137_v8_ = d_137_v8_
            d_126_v0_ = d_126_v0_
            d_442_v219_: _dafny.Map
            d_442_v219_ = _dafny.Map({len((d_131_v3_).f13): d_361_v172_})
            d_137_v8_ = default__.fm33((d_126_v0_)[default__.safeIndex(701, (d_126_v0_).length(0))], (len(((d_442_v219_)[-795] if (-795) in (d_442_v219_) else ((d_442_v219_)[d_137_v8_] if (d_137_v8_) in (d_442_v219_) else d_361_v172_)))) <= (d_137_v8_), (d_137_v8_) * (d_137_v8_), (d_131_v3_).f12, d_130_globalState_)
        d_443_v220_: _dafny.Seq
        d_444_v221_: bool
        d_445_v222_: _dafny.Seq
        d_446_v223_: _dafny.Seq
        out38_: _dafny.Seq
        out39_: bool
        out40_: _dafny.Seq
        out41_: _dafny.Seq
        out38_, out39_, out40_, out41_ = (d_174_v39_).m0(d_130_globalState_)
        d_443_v220_ = out38_
        d_444_v221_ = out39_
        d_445_v222_ = out40_
        d_446_v223_ = out41_
        _dafny.print(_dafny.string_of((d_126_v0_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_v0_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_v0_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_v0_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_v0_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_v0_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_v0_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_v0_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_v0_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_v0_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_v0_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_v0_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_v0_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_v0_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_v0_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_v0_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_v0_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_v0_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_128_v1_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_129_v2_) == (_dafny.Map({4: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_130_globalState_).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_130_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_130_globalState_).f2)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_130_globalState_).f2)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_130_globalState_).f2)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_130_globalState_).f2)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_130_globalState_).f2)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_130_globalState_).f2)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_130_globalState_).f2)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_130_globalState_).f2)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_130_globalState_).f2)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_130_globalState_).f2)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_130_globalState_).f2)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_130_globalState_).f2)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_130_globalState_).f2)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_130_globalState_).f2)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_130_globalState_).f2)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_130_globalState_).f2)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_130_globalState_).f2)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_130_globalState_).f2)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_130_globalState_).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_130_globalState_.f4) == (_dafny.Map({_dafny.Map({4: True}): False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_130_globalState_).f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_130_globalState_).f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_130_globalState_).f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_130_globalState_.f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_131_v3_).f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_131_v3_).f13) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_135_v6_).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_136_v7_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_137_v8_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_172_v37_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_173_v38_) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_174_v39_).f9) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_174_v39_).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_175_v40_).cf104).f9) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_175_v40_).cf104).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_176_v41_)[0]).f9) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_176_v41_)[0]).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_176_v41_)[1]).f9) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_176_v41_)[1]).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_176_v41_)[2]).f9) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_176_v41_)[2]).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_176_v41_)[3]).f9) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_176_v41_)[3]).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_176_v41_)[4]).f9) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_176_v41_)[4]).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_176_v41_)[5]).f9) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_176_v41_)[5]).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_176_v41_)[6]).f9) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_176_v41_)[6]).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_176_v41_)[7]).f9) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_176_v41_)[7]).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_176_v41_)[8]).f9) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_176_v41_)[8]).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_176_v41_)[9]).f9) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_176_v41_)[9]).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_176_v41_)[10]).f9) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_176_v41_)[10]).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_176_v41_)[11]).f9) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_176_v41_)[11]).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_176_v41_)[12]).f9) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_176_v41_)[12]).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_176_v41_)[13]).f9) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_176_v41_)[13]).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_176_v41_)[14]).f9) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_176_v41_)[14]).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_176_v41_)[15]).f9) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_176_v41_)[15]).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_176_v41_)[16]).f9) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_176_v41_)[16]).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_176_v41_)[17]).f9) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_176_v41_)[17]).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_176_v41_)[18]).f9) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_176_v41_)[18]).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_176_v41_)[19]).f9) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_176_v41_)[19]).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_176_v41_)[20]).f9) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_176_v41_)[20]).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_176_v41_)[21]).f9) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_176_v41_)[21]).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_176_v41_)[22]).f9) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_176_v41_)[22]).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_176_v41_)[23]).f9) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_176_v41_)[23]).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_176_v41_)[24]).f9) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_176_v41_)[24]).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_176_v41_)[25]).f9) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_176_v41_)[25]).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_177_v42_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_179_v44_).cf127).f9) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_179_v44_).cf127).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_180_v45_) == (_dafny.SeqWithoutIsStrInference([-556]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_181_v46_).cf4) == (_dafny.SeqWithoutIsStrInference([-556]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_265_v110_) == (_dafny.MultiSet([True, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_268_v113_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_361_v172_) == (_dafny.SeqWithoutIsStrInference([True, True, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_443_v220_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_444_v221_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_445_v222_) == (_dafny.SeqWithoutIsStrInference([_dafny.Set({True})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_446_v223_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(int(0), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D0_DC1)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)

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


class D1:
    @classmethod
    def default(cls, ):
        return lambda: D1_DC3(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D1_DC2)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D1_DC5)

class D1_DC3(D1, NamedTuple('DC3', [('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({self.cf5.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC4(D1, NamedTuple('DC4', [('cf6', Any), ('cf7', Any), ('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf6)}, {_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf6 == __o.cf6 and self.cf7 == __o.cf7 and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC2(D1, NamedTuple('DC2', [('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC2({_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC2) and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC5(D1, NamedTuple('DC5', [('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.CodePoint('D')
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)

class D2_DC6(D2, NamedTuple('DC6', [('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC8(False, False)
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

class D3_DC8(D3, NamedTuple('DC8', [('cf12', Any), ('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC8({_dafny.string_of(self.cf12)}, {_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC8) and self.cf12 == __o.cf12 and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC7(D3, NamedTuple('DC7', [('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC7({_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC7) and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC9(D3, NamedTuple('DC9', [('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC11(int(0), _dafny.Set({}), False)
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

class D4_DC11(D4, NamedTuple('DC11', [('cf16', Any), ('cf17', Any), ('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC11({_dafny.string_of(self.cf16)}, {_dafny.string_of(self.cf17)}, {_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC11) and self.cf16 == __o.cf16 and self.cf17 == __o.cf17 and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC12(D4, NamedTuple('DC12', [('cf19', Any), ('cf20', Any), ('cf21', Any), ('cf22', Any), ('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC12({_dafny.string_of(self.cf19)}, {_dafny.string_of(self.cf20)}, {_dafny.string_of(self.cf21)}, {_dafny.string_of(self.cf22)}, {_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC12) and self.cf19 == __o.cf19 and self.cf20 == __o.cf20 and self.cf21 == __o.cf21 and self.cf22 == __o.cf22 and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC10(D4, NamedTuple('DC10', [('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC10({_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC10) and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC14(int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D5_DC14)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D5_DC13)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D5_DC15)

class D5_DC14(D5, NamedTuple('DC14', [('cf25', Any), ('cf26', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC14({_dafny.string_of(self.cf25)}, {_dafny.string_of(self.cf26)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC14) and self.cf25 == __o.cf25 and self.cf26 == __o.cf26
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC13(D5, NamedTuple('DC13', [('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC13({_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC13) and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC15(D5, NamedTuple('DC15', [('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC15({_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC15) and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC17(False, False, int(0), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D6_DC17)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D6_DC18)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D6_DC19)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D6_DC16)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D6_DC20)

class D6_DC17(D6, NamedTuple('DC17', [('cf29', Any), ('cf30', Any), ('cf31', Any), ('cf32', Any), ('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC17({_dafny.string_of(self.cf29)}, {_dafny.string_of(self.cf30)}, {_dafny.string_of(self.cf31)}, {_dafny.string_of(self.cf32)}, {_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC17) and self.cf29 == __o.cf29 and self.cf30 == __o.cf30 and self.cf31 == __o.cf31 and self.cf32 == __o.cf32 and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC18(D6, NamedTuple('DC18', [('cf34', Any), ('cf35', Any), ('cf36', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC18({_dafny.string_of(self.cf34)}, {_dafny.string_of(self.cf35)}, {_dafny.string_of(self.cf36)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC18) and self.cf34 == __o.cf34 and self.cf35 == __o.cf35 and self.cf36 == __o.cf36
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC19(D6, NamedTuple('DC19', [('cf37', Any), ('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC19({_dafny.string_of(self.cf37)}, {_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC19) and self.cf37 == __o.cf37 and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC16(D6, NamedTuple('DC16', [('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC16({_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC16) and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC20(D6, NamedTuple('DC20', [('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC20({_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC20) and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC22()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D7_DC22)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D7_DC21)

class D7_DC22(D7, NamedTuple('DC22', [])):
    def __dafnystr__(self) -> str:
        return f'D7.DC22'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC22)
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC21(D7, NamedTuple('DC21', [('cf40', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC21({_dafny.string_of(self.cf40)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC21) and self.cf40 == __o.cf40
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC24(False, _dafny.Set({}), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D8_DC24)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D8_DC23)

class D8_DC24(D8, NamedTuple('DC24', [('cf42', Any), ('cf43', Any), ('cf44', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC24({_dafny.string_of(self.cf42)}, {_dafny.string_of(self.cf43)}, {_dafny.string_of(self.cf44)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC24) and self.cf42 == __o.cf42 and self.cf43 == __o.cf43 and self.cf44 == __o.cf44
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC23(D8, NamedTuple('DC23', [('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC23({_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC23) and self.cf41 == __o.cf41
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

class D9_DC25(D9, NamedTuple('DC25', [('cf45', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC25({_dafny.string_of(self.cf45)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC25) and self.cf45 == __o.cf45
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D10_DC26)

class D10_DC26(D10, NamedTuple('DC26', [('cf46', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC26({_dafny.string_of(self.cf46)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC26) and self.cf46 == __o.cf46
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC28(int(0), False, _dafny.Map({}), False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D11_DC28)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D11_DC27)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D11_DC29)

class D11_DC28(D11, NamedTuple('DC28', [('cf48', Any), ('cf49', Any), ('cf50', Any), ('cf51', Any), ('cf52', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC28({_dafny.string_of(self.cf48)}, {_dafny.string_of(self.cf49)}, {_dafny.string_of(self.cf50)}, {_dafny.string_of(self.cf51)}, {self.cf52.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC28) and self.cf48 == __o.cf48 and self.cf49 == __o.cf49 and self.cf50 == __o.cf50 and self.cf51 == __o.cf51 and self.cf52 == __o.cf52
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC27(D11, NamedTuple('DC27', [('cf47', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC27({_dafny.string_of(self.cf47)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC27) and self.cf47 == __o.cf47
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC29(D11, NamedTuple('DC29', [('cf53', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC29({_dafny.string_of(self.cf53)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC29) and self.cf53 == __o.cf53
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC31(False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D12_DC31)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D12_DC32)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D12_DC30)

class D12_DC31(D12, NamedTuple('DC31', [('cf55', Any), ('cf56', Any), ('cf57', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC31({_dafny.string_of(self.cf55)}, {self.cf56.VerbatimString(True)}, {self.cf57.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC31) and self.cf55 == __o.cf55 and self.cf56 == __o.cf56 and self.cf57 == __o.cf57
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC32(D12, NamedTuple('DC32', [('cf58', Any), ('cf59', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC32({_dafny.string_of(self.cf58)}, {_dafny.string_of(self.cf59)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC32) and self.cf58 == __o.cf58 and self.cf59 == __o.cf59
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC30(D12, NamedTuple('DC30', [('cf54', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC30({_dafny.string_of(self.cf54)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC30) and self.cf54 == __o.cf54
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D13_DC33)

class D13_DC33(D13, NamedTuple('DC33', [('cf60', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC33({_dafny.string_of(self.cf60)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC33) and self.cf60 == __o.cf60
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: D14_DC35(int(0), False, _dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D14_DC35)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D14_DC36)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D14_DC34)

class D14_DC35(D14, NamedTuple('DC35', [('cf62', Any), ('cf63', Any), ('cf64', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC35({_dafny.string_of(self.cf62)}, {_dafny.string_of(self.cf63)}, {_dafny.string_of(self.cf64)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC35) and self.cf62 == __o.cf62 and self.cf63 == __o.cf63 and self.cf64 == __o.cf64
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC36(D14, NamedTuple('DC36', [('cf65', Any), ('cf66', Any), ('cf67', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC36({_dafny.string_of(self.cf65)}, {self.cf66.VerbatimString(True)}, {self.cf67.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC36) and self.cf65 == __o.cf65 and self.cf66 == __o.cf66 and self.cf67 == __o.cf67
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC34(D14, NamedTuple('DC34', [('cf61', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC34({_dafny.string_of(self.cf61)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC34) and self.cf61 == __o.cf61
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: D15_DC38(False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), False, False, _dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D15_DC38)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D15_DC39)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D15_DC37)

class D15_DC38(D15, NamedTuple('DC38', [('cf69', Any), ('cf70', Any), ('cf71', Any), ('cf72', Any), ('cf73', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC38({_dafny.string_of(self.cf69)}, {self.cf70.VerbatimString(True)}, {_dafny.string_of(self.cf71)}, {_dafny.string_of(self.cf72)}, {_dafny.string_of(self.cf73)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC38) and self.cf69 == __o.cf69 and self.cf70 == __o.cf70 and self.cf71 == __o.cf71 and self.cf72 == __o.cf72 and self.cf73 == __o.cf73
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC39(D15, NamedTuple('DC39', [('cf74', Any), ('cf75', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC39({_dafny.string_of(self.cf74)}, {_dafny.string_of(self.cf75)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC39) and self.cf74 == __o.cf74 and self.cf75 == __o.cf75
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC37(D15, NamedTuple('DC37', [('cf68', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC37({_dafny.string_of(self.cf68)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC37) and self.cf68 == __o.cf68
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: D16_DC41(False, int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC41(self) -> bool:
        return isinstance(self, D16_DC41)
    @property
    def is_DC42(self) -> bool:
        return isinstance(self, D16_DC42)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D16_DC40)

class D16_DC41(D16, NamedTuple('DC41', [('cf77', Any), ('cf78', Any), ('cf79', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC41({_dafny.string_of(self.cf77)}, {_dafny.string_of(self.cf78)}, {_dafny.string_of(self.cf79)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC41) and self.cf77 == __o.cf77 and self.cf78 == __o.cf78 and self.cf79 == __o.cf79
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC42(D16, NamedTuple('DC42', [('cf80', Any), ('cf81', Any), ('cf82', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC42({_dafny.string_of(self.cf80)}, {_dafny.string_of(self.cf81)}, {_dafny.string_of(self.cf82)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC42) and self.cf80 == __o.cf80 and self.cf81 == __o.cf81 and self.cf82 == __o.cf82
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC40(D16, NamedTuple('DC40', [('cf76', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC40({_dafny.string_of(self.cf76)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC40) and self.cf76 == __o.cf76
    def __hash__(self) -> int:
        return super().__hash__()


class D17:
    @classmethod
    def default(cls, ):
        return lambda: D17_DC44(int(0), _dafny.Seq({}), _dafny.CodePoint('D'), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC44(self) -> bool:
        return isinstance(self, D17_DC44)
    @property
    def is_DC45(self) -> bool:
        return isinstance(self, D17_DC45)
    @property
    def is_DC43(self) -> bool:
        return isinstance(self, D17_DC43)

class D17_DC44(D17, NamedTuple('DC44', [('cf84', Any), ('cf85', Any), ('cf86', Any), ('cf87', Any), ('cf88', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC44({_dafny.string_of(self.cf84)}, {_dafny.string_of(self.cf85)}, {_dafny.string_of(self.cf86)}, {_dafny.string_of(self.cf87)}, {_dafny.string_of(self.cf88)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC44) and self.cf84 == __o.cf84 and self.cf85 == __o.cf85 and self.cf86 == __o.cf86 and self.cf87 == __o.cf87 and self.cf88 == __o.cf88
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC45(D17, NamedTuple('DC45', [])):
    def __dafnystr__(self) -> str:
        return f'D17.DC45'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC45)
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC43(D17, NamedTuple('DC43', [('cf83', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC43({_dafny.string_of(self.cf83)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC43) and self.cf83 == __o.cf83
    def __hash__(self) -> int:
        return super().__hash__()


class D18:
    @classmethod
    def default(cls, ):
        return lambda: D18_DC47(False, D6.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC47(self) -> bool:
        return isinstance(self, D18_DC47)
    @property
    def is_DC46(self) -> bool:
        return isinstance(self, D18_DC46)

class D18_DC47(D18, NamedTuple('DC47', [('cf90', Any), ('cf91', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC47({_dafny.string_of(self.cf90)}, {_dafny.string_of(self.cf91)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC47) and self.cf90 == __o.cf90 and self.cf91 == __o.cf91
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC46(D18, NamedTuple('DC46', [('cf89', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC46({_dafny.string_of(self.cf89)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC46) and self.cf89 == __o.cf89
    def __hash__(self) -> int:
        return super().__hash__()


class D19:
    @classmethod
    def default(cls, ):
        return lambda: D19_DC49(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC49(self) -> bool:
        return isinstance(self, D19_DC49)
    @property
    def is_DC50(self) -> bool:
        return isinstance(self, D19_DC50)
    @property
    def is_DC48(self) -> bool:
        return isinstance(self, D19_DC48)

class D19_DC49(D19, NamedTuple('DC49', [('cf93', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC49({_dafny.string_of(self.cf93)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC49) and self.cf93 == __o.cf93
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC50(D19, NamedTuple('DC50', [('cf94', Any), ('cf95', Any), ('cf96', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC50({_dafny.string_of(self.cf94)}, {_dafny.string_of(self.cf95)}, {_dafny.string_of(self.cf96)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC50) and self.cf94 == __o.cf94 and self.cf95 == __o.cf95 and self.cf96 == __o.cf96
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC48(D19, NamedTuple('DC48', [('cf92', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC48({_dafny.string_of(self.cf92)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC48) and self.cf92 == __o.cf92
    def __hash__(self) -> int:
        return super().__hash__()


class D20:
    @classmethod
    def default(cls, ):
        return lambda: D20_DC52(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC52(self) -> bool:
        return isinstance(self, D20_DC52)
    @property
    def is_DC51(self) -> bool:
        return isinstance(self, D20_DC51)

class D20_DC52(D20, NamedTuple('DC52', [('cf98', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC52({_dafny.string_of(self.cf98)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC52) and self.cf98 == __o.cf98
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC51(D20, NamedTuple('DC51', [('cf97', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC51({_dafny.string_of(self.cf97)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC51) and self.cf97 == __o.cf97
    def __hash__(self) -> int:
        return super().__hash__()


class D21:
    @classmethod
    def default(cls, ):
        return lambda: D21_DC54(_dafny.CodePoint('D'), int(0), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC54(self) -> bool:
        return isinstance(self, D21_DC54)
    @property
    def is_DC53(self) -> bool:
        return isinstance(self, D21_DC53)
    @property
    def is_DC55(self) -> bool:
        return isinstance(self, D21_DC55)

class D21_DC54(D21, NamedTuple('DC54', [('cf100', Any), ('cf101', Any), ('cf102', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC54({_dafny.string_of(self.cf100)}, {_dafny.string_of(self.cf101)}, {self.cf102.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC54) and self.cf100 == __o.cf100 and self.cf101 == __o.cf101 and self.cf102 == __o.cf102
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC53(D21, NamedTuple('DC53', [('cf99', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC53({_dafny.string_of(self.cf99)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC53) and self.cf99 == __o.cf99
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC55(D21, NamedTuple('DC55', [('cf103', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC55({_dafny.string_of(self.cf103)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC55) and self.cf103 == __o.cf103
    def __hash__(self) -> int:
        return super().__hash__()


class D22:
    @classmethod
    def default(cls, ):
        return lambda: D22_DC57()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC57(self) -> bool:
        return isinstance(self, D22_DC57)
    @property
    def is_DC56(self) -> bool:
        return isinstance(self, D22_DC56)

class D22_DC57(D22, NamedTuple('DC57', [])):
    def __dafnystr__(self) -> str:
        return f'D22.DC57'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC57)
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC56(D22, NamedTuple('DC56', [('cf104', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC56({_dafny.string_of(self.cf104)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC56) and self.cf104 == __o.cf104
    def __hash__(self) -> int:
        return super().__hash__()


class D23:
    @classmethod
    def default(cls, ):
        return lambda: D23_DC59(D6.default()(), False, int(0), _dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC59(self) -> bool:
        return isinstance(self, D23_DC59)
    @property
    def is_DC58(self) -> bool:
        return isinstance(self, D23_DC58)

class D23_DC59(D23, NamedTuple('DC59', [('cf106', Any), ('cf107', Any), ('cf108', Any), ('cf109', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC59({_dafny.string_of(self.cf106)}, {_dafny.string_of(self.cf107)}, {_dafny.string_of(self.cf108)}, {_dafny.string_of(self.cf109)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC59) and self.cf106 == __o.cf106 and self.cf107 == __o.cf107 and self.cf108 == __o.cf108 and self.cf109 == __o.cf109
    def __hash__(self) -> int:
        return super().__hash__()

class D23_DC58(D23, NamedTuple('DC58', [('cf105', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC58({_dafny.string_of(self.cf105)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC58) and self.cf105 == __o.cf105
    def __hash__(self) -> int:
        return super().__hash__()


class D24:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC60(self) -> bool:
        return isinstance(self, D24_DC60)

class D24_DC60(D24, NamedTuple('DC60', [('cf110', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC60({_dafny.string_of(self.cf110)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC60) and self.cf110 == __o.cf110
    def __hash__(self) -> int:
        return super().__hash__()


class D25:
    @classmethod
    def default(cls, ):
        return lambda: D25_DC62()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC62(self) -> bool:
        return isinstance(self, D25_DC62)
    @property
    def is_DC63(self) -> bool:
        return isinstance(self, D25_DC63)
    @property
    def is_DC61(self) -> bool:
        return isinstance(self, D25_DC61)

class D25_DC62(D25, NamedTuple('DC62', [])):
    def __dafnystr__(self) -> str:
        return f'D25.DC62'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC62)
    def __hash__(self) -> int:
        return super().__hash__()

class D25_DC63(D25, NamedTuple('DC63', [('cf112', Any), ('cf113', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC63({self.cf112.VerbatimString(True)}, {_dafny.string_of(self.cf113)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC63) and self.cf112 == __o.cf112 and self.cf113 == __o.cf113
    def __hash__(self) -> int:
        return super().__hash__()

class D25_DC61(D25, NamedTuple('DC61', [('cf111', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC61({_dafny.string_of(self.cf111)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC61) and self.cf111 == __o.cf111
    def __hash__(self) -> int:
        return super().__hash__()


class D26:
    @classmethod
    def default(cls, ):
        return lambda: D26_DC65(False, _dafny.Array(None, 0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC65(self) -> bool:
        return isinstance(self, D26_DC65)
    @property
    def is_DC64(self) -> bool:
        return isinstance(self, D26_DC64)

class D26_DC65(D26, NamedTuple('DC65', [('cf115', Any), ('cf116', Any), ('cf117', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC65({_dafny.string_of(self.cf115)}, {_dafny.string_of(self.cf116)}, {_dafny.string_of(self.cf117)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC65) and self.cf115 == __o.cf115 and self.cf116 == __o.cf116 and self.cf117 == __o.cf117
    def __hash__(self) -> int:
        return super().__hash__()

class D26_DC64(D26, NamedTuple('DC64', [('cf114', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC64({_dafny.string_of(self.cf114)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC64) and self.cf114 == __o.cf114
    def __hash__(self) -> int:
        return super().__hash__()


class D27:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC66(self) -> bool:
        return isinstance(self, D27_DC66)

class D27_DC66(D27, NamedTuple('DC66', [('cf118', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC66({_dafny.string_of(self.cf118)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC66) and self.cf118 == __o.cf118
    def __hash__(self) -> int:
        return super().__hash__()


class D28:
    @classmethod
    def default(cls, ):
        return lambda: D28_DC68(int(0), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC68(self) -> bool:
        return isinstance(self, D28_DC68)
    @property
    def is_DC69(self) -> bool:
        return isinstance(self, D28_DC69)
    @property
    def is_DC70(self) -> bool:
        return isinstance(self, D28_DC70)
    @property
    def is_DC67(self) -> bool:
        return isinstance(self, D28_DC67)
    @property
    def is_DC71(self) -> bool:
        return isinstance(self, D28_DC71)

class D28_DC68(D28, NamedTuple('DC68', [('cf120', Any), ('cf121', Any), ('cf122', Any)])):
    def __dafnystr__(self) -> str:
        return f'D28.DC68({_dafny.string_of(self.cf120)}, {_dafny.string_of(self.cf121)}, {_dafny.string_of(self.cf122)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D28_DC68) and self.cf120 == __o.cf120 and self.cf121 == __o.cf121 and self.cf122 == __o.cf122
    def __hash__(self) -> int:
        return super().__hash__()

class D28_DC69(D28, NamedTuple('DC69', [])):
    def __dafnystr__(self) -> str:
        return f'D28.DC69'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D28_DC69)
    def __hash__(self) -> int:
        return super().__hash__()

class D28_DC70(D28, NamedTuple('DC70', [('cf123', Any), ('cf124', Any)])):
    def __dafnystr__(self) -> str:
        return f'D28.DC70({_dafny.string_of(self.cf123)}, {_dafny.string_of(self.cf124)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D28_DC70) and self.cf123 == __o.cf123 and self.cf124 == __o.cf124
    def __hash__(self) -> int:
        return super().__hash__()

class D28_DC67(D28, NamedTuple('DC67', [('cf119', Any)])):
    def __dafnystr__(self) -> str:
        return f'D28.DC67({_dafny.string_of(self.cf119)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D28_DC67) and self.cf119 == __o.cf119
    def __hash__(self) -> int:
        return super().__hash__()

class D28_DC71(D28, NamedTuple('DC71', [('cf125', Any)])):
    def __dafnystr__(self) -> str:
        return f'D28.DC71({_dafny.string_of(self.cf125)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D28_DC71) and self.cf125 == __o.cf125
    def __hash__(self) -> int:
        return super().__hash__()


class D29:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC72(self) -> bool:
        return isinstance(self, D29_DC72)

class D29_DC72(D29, NamedTuple('DC72', [('cf126', Any)])):
    def __dafnystr__(self) -> str:
        return f'D29.DC72({_dafny.string_of(self.cf126)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D29_DC72) and self.cf126 == __o.cf126
    def __hash__(self) -> int:
        return super().__hash__()


class D30:
    @classmethod
    def default(cls, ):
        return lambda: D30_DC74(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC74(self) -> bool:
        return isinstance(self, D30_DC74)
    @property
    def is_DC75(self) -> bool:
        return isinstance(self, D30_DC75)
    @property
    def is_DC73(self) -> bool:
        return isinstance(self, D30_DC73)
    @property
    def is_DC76(self) -> bool:
        return isinstance(self, D30_DC76)

class D30_DC74(D30, NamedTuple('DC74', [('cf128', Any)])):
    def __dafnystr__(self) -> str:
        return f'D30.DC74({self.cf128.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D30_DC74) and self.cf128 == __o.cf128
    def __hash__(self) -> int:
        return super().__hash__()

class D30_DC75(D30, NamedTuple('DC75', [('cf129', Any)])):
    def __dafnystr__(self) -> str:
        return f'D30.DC75({_dafny.string_of(self.cf129)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D30_DC75) and self.cf129 == __o.cf129
    def __hash__(self) -> int:
        return super().__hash__()

class D30_DC73(D30, NamedTuple('DC73', [('cf127', Any)])):
    def __dafnystr__(self) -> str:
        return f'D30.DC73({_dafny.string_of(self.cf127)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D30_DC73) and self.cf127 == __o.cf127
    def __hash__(self) -> int:
        return super().__hash__()

class D30_DC76(D30, NamedTuple('DC76', [('cf130', Any)])):
    def __dafnystr__(self) -> str:
        return f'D30.DC76({_dafny.string_of(self.cf130)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D30_DC76) and self.cf130 == __o.cf130
    def __hash__(self) -> int:
        return super().__hash__()


class D31:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC77(self) -> bool:
        return isinstance(self, D31_DC77)

class D31_DC77(D31, NamedTuple('DC77', [('cf131', Any)])):
    def __dafnystr__(self) -> str:
        return f'D31.DC77({_dafny.string_of(self.cf131)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D31_DC77) and self.cf131 == __o.cf131
    def __hash__(self) -> int:
        return super().__hash__()


class D32:
    @classmethod
    def default(cls, ):
        return lambda: D32_DC79(False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.Seq({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC79(self) -> bool:
        return isinstance(self, D32_DC79)
    @property
    def is_DC80(self) -> bool:
        return isinstance(self, D32_DC80)
    @property
    def is_DC78(self) -> bool:
        return isinstance(self, D32_DC78)

class D32_DC79(D32, NamedTuple('DC79', [('cf133', Any), ('cf134', Any), ('cf135', Any)])):
    def __dafnystr__(self) -> str:
        return f'D32.DC79({_dafny.string_of(self.cf133)}, {self.cf134.VerbatimString(True)}, {_dafny.string_of(self.cf135)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D32_DC79) and self.cf133 == __o.cf133 and self.cf134 == __o.cf134 and self.cf135 == __o.cf135
    def __hash__(self) -> int:
        return super().__hash__()

class D32_DC80(D32, NamedTuple('DC80', [('cf136', Any), ('cf137', Any), ('cf138', Any)])):
    def __dafnystr__(self) -> str:
        return f'D32.DC80({_dafny.string_of(self.cf136)}, {_dafny.string_of(self.cf137)}, {_dafny.string_of(self.cf138)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D32_DC80) and self.cf136 == __o.cf136 and self.cf137 == __o.cf137 and self.cf138 == __o.cf138
    def __hash__(self) -> int:
        return super().__hash__()

class D32_DC78(D32, NamedTuple('DC78', [('cf132', Any)])):
    def __dafnystr__(self) -> str:
        return f'D32.DC78({_dafny.string_of(self.cf132)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D32_DC78) and self.cf132 == __o.cf132
    def __hash__(self) -> int:
        return super().__hash__()


class D33:
    @classmethod
    def default(cls, ):
        return lambda: D33_DC82(False, False, None, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC82(self) -> bool:
        return isinstance(self, D33_DC82)
    @property
    def is_DC81(self) -> bool:
        return isinstance(self, D33_DC81)
    @property
    def is_DC83(self) -> bool:
        return isinstance(self, D33_DC83)

class D33_DC82(D33, NamedTuple('DC82', [('cf140', Any), ('cf141', Any), ('cf142', Any), ('cf143', Any)])):
    def __dafnystr__(self) -> str:
        return f'D33.DC82({_dafny.string_of(self.cf140)}, {_dafny.string_of(self.cf141)}, {_dafny.string_of(self.cf142)}, {_dafny.string_of(self.cf143)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D33_DC82) and self.cf140 == __o.cf140 and self.cf141 == __o.cf141 and self.cf142 == __o.cf142 and self.cf143 == __o.cf143
    def __hash__(self) -> int:
        return super().__hash__()

class D33_DC81(D33, NamedTuple('DC81', [('cf139', Any)])):
    def __dafnystr__(self) -> str:
        return f'D33.DC81({_dafny.string_of(self.cf139)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D33_DC81) and self.cf139 == __o.cf139
    def __hash__(self) -> int:
        return super().__hash__()

class D33_DC83(D33, NamedTuple('DC83', [('cf144', Any)])):
    def __dafnystr__(self) -> str:
        return f'D33.DC83({_dafny.string_of(self.cf144)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D33_DC83) and self.cf144 == __o.cf144
    def __hash__(self) -> int:
        return super().__hash__()


class D34:
    @classmethod
    def default(cls, ):
        return lambda: D34_DC85(_dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC85(self) -> bool:
        return isinstance(self, D34_DC85)
    @property
    def is_DC86(self) -> bool:
        return isinstance(self, D34_DC86)
    @property
    def is_DC84(self) -> bool:
        return isinstance(self, D34_DC84)

class D34_DC85(D34, NamedTuple('DC85', [('cf146', Any)])):
    def __dafnystr__(self) -> str:
        return f'D34.DC85({_dafny.string_of(self.cf146)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D34_DC85) and self.cf146 == __o.cf146
    def __hash__(self) -> int:
        return super().__hash__()

class D34_DC86(D34, NamedTuple('DC86', [('cf147', Any), ('cf148', Any)])):
    def __dafnystr__(self) -> str:
        return f'D34.DC86({_dafny.string_of(self.cf147)}, {_dafny.string_of(self.cf148)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D34_DC86) and self.cf147 == __o.cf147 and self.cf148 == __o.cf148
    def __hash__(self) -> int:
        return super().__hash__()

class D34_DC84(D34, NamedTuple('DC84', [('cf145', Any)])):
    def __dafnystr__(self) -> str:
        return f'D34.DC84({_dafny.string_of(self.cf145)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D34_DC84) and self.cf145 == __o.cf145
    def __hash__(self) -> int:
        return super().__hash__()


class D35:
    @classmethod
    def default(cls, ):
        return lambda: D35_DC88()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC88(self) -> bool:
        return isinstance(self, D35_DC88)
    @property
    def is_DC89(self) -> bool:
        return isinstance(self, D35_DC89)
    @property
    def is_DC87(self) -> bool:
        return isinstance(self, D35_DC87)

class D35_DC88(D35, NamedTuple('DC88', [])):
    def __dafnystr__(self) -> str:
        return f'D35.DC88'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D35_DC88)
    def __hash__(self) -> int:
        return super().__hash__()

class D35_DC89(D35, NamedTuple('DC89', [('cf150', Any), ('cf151', Any), ('cf152', Any)])):
    def __dafnystr__(self) -> str:
        return f'D35.DC89({_dafny.string_of(self.cf150)}, {_dafny.string_of(self.cf151)}, {_dafny.string_of(self.cf152)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D35_DC89) and self.cf150 == __o.cf150 and self.cf151 == __o.cf151 and self.cf152 == __o.cf152
    def __hash__(self) -> int:
        return super().__hash__()

class D35_DC87(D35, NamedTuple('DC87', [('cf149', Any)])):
    def __dafnystr__(self) -> str:
        return f'D35.DC87({_dafny.string_of(self.cf149)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D35_DC87) and self.cf149 == __o.cf149
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    def m0(self, globalState):
        pass

    def m1(self, p0, globalState):
        pass


class T1:
    pass

class GlobalState:
    def  __init__(self):
        self.f4: _dafny.Map = _dafny.Map({})
        self.f8: bool = False
        self._f0: bool = False
        self._f1: bool = False
        self._f2: _dafny.Array = _dafny.Array(None, 0)
        self._f3: bool = False
        self._f5: int = int(0)
        self._f6: bool = False
        self._f7: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8):
        (self)._f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2
        (self)._f3 = f3
        (self).f4 = f4
        (self)._f5 = f5
        (self)._f6 = f6
        (self)._f7 = f7
        (self).f8 = f8

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
    def f5(self):
        return self._f5
    @property
    def f6(self):
        return self._f6
    @property
    def f7(self):
        return self._f7

class C0:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self):
        pass
        pass

    def fm5(self, p0, p1, p2, p3, globalState):
        return ((-952) - (834)) >= (default__.safeDivisionInt(len(_dafny.Map({False: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aiohhcmxm"))})), -442))


class C1(T0):
    def  __init__(self):
        self._f12: bool = False
        self._f13: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    def ctor__(self, f12, f13):
        (self)._f12 = f12
        (self)._f13 = f13

    def fm32(self, p0, p1, globalState):
        return (self).f12

    def m0(self, globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r1: bool = False
        r2: _dafny.Seq = _dafny.Seq({})
        r3: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        d_447_v0_: _dafny.Array
        nw45_ = _dafny.Array(int(0), 16)
        d_447_v0_ = nw45_
        d_448_v1_: int
        d_448_v1_ = 261
        index50_ = default__.safeIndex(609, (d_447_v0_).length(0))
        (d_447_v0_)[index50_] = default__.safeDivisionInt(d_448_v1_, d_448_v1_)
        index51_ = default__.safeIndex(609, (d_447_v0_).length(0))
        (d_447_v0_)[index51_] = d_448_v1_
        (globalState).f8 = (self).f12
        d_449_v2_: _dafny.Map
        d_449_v2_ = _dafny.Map({d_448_v1_: (self).f12})
        d_449_v2_ = (d_449_v2_).set((default__.fm33((self).f12, (self).f12, d_448_v1_, True, globalState)) * ((d_447_v0_)[default__.safeIndex(609, (d_447_v0_).length(0))]), ((self).f12) == (not((self).f12)))
        d_450_v3_: str
        d_450_v3_ = _dafny.CodePoint('m')
        if (self).fm32((d_447_v0_)[default__.safeIndex(609, (d_447_v0_).length(0))], d_450_v3_, globalState):
            r0 = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tsoaubs"))
            d_451_v4_: C0
            nw46_ = C0()
            nw46_.ctor__()
            d_451_v4_ = nw46_
            if ((d_447_v0_)[default__.safeIndex(609, (d_447_v0_).length(0))]) != ((d_447_v0_)[default__.safeIndex(609, (d_447_v0_).length(0))]):
                (globalState).f8 = not((self).f12)
                d_452_v5_: _dafny.Array
                nw47_ = _dafny.Array(False, 27)
                d_452_v5_ = nw47_
                index52_ = default__.safeIndex(180, (d_452_v5_).length(0))
                (d_452_v5_)[index52_] = ((d_449_v2_)[(d_447_v0_)[default__.safeIndex(609, (d_447_v0_).length(0))]] if ((d_447_v0_)[default__.safeIndex(609, (d_447_v0_).length(0))]) in (d_449_v2_) else False)
                d_453_v6_: _dafny.MultiSet
                d_453_v6_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([(self).f12])])
                index53_ = default__.safeIndex(180, (d_452_v5_).length(0))
                (d_452_v5_)[index53_] = (d_453_v6_).ispropersubset(d_453_v6_)
                nw48_ = _dafny.Array(int(0), 23)
                d_447_v0_ = nw48_
                (globalState).f8 = (d_452_v5_)[default__.safeIndex(180, (d_452_v5_).length(0))]
                d_454_v7_: _dafny.Array
                nw49_ = _dafny.Array(_dafny.Set({}), 1)
                d_454_v7_ = nw49_
                d_454_v7_ = d_454_v7_
            elif True:
                index54_ = default__.safeIndex(609, (d_447_v0_).length(0))
                (d_447_v0_)[index54_] = 606
                d_455_v8_: C0
                nw50_ = C0()
                nw50_.ctor__()
                d_455_v8_ = nw50_
                d_456_v9_: _dafny.MultiSet
                d_456_v9_ = _dafny.MultiSet([not((self).f12)])
                d_448_v1_ = default__.fm33((self).f12, not((d_455_v8_).fm5((self).f12, d_448_v1_, (d_447_v0_)[default__.safeIndex(609, (d_447_v0_).length(0))], d_456_v9_, globalState)), default__.safeModuloInt(d_448_v1_, d_448_v1_), (self).f12, globalState)
                d_457_v10_: _dafny.Map
                d_457_v10_ = _dafny.Map({d_450_v3_: (self).f13})
                d_458_v11_: _dafny.Map
                d_458_v11_ = _dafny.Map({(d_447_v0_)[default__.safeIndex(609, (d_447_v0_).length(0))]: d_448_v1_})
                rhs58_ = (d_457_v10_) | ((d_457_v10_).set(((self).f13)[default__.safeIndex(len(d_458_v11_), len((self).f13))], (self).f13))
                rhs59_ = (d_448_v1_) + (d_448_v1_)
                rhs60_ = (self).f12
                lhs44_ = globalState
                d_457_v10_ = rhs58_
                d_448_v1_ = rhs59_
                lhs44_.f8 = rhs60_
                d_459_v12_: _dafny.Array
                def lambda9_(d_460_i0_):
                    return (self).f12

                init5_ = lambda9_
                nw51_ = _dafny.Array(None, 5)
                for i0_5_ in range(nw51_.length(0)):
                    nw51_[i0_5_] = init5_(i0_5_)
                d_459_v12_ = nw51_
                d_461_v13_: _dafny.MultiSet
                d_461_v13_ = _dafny.MultiSet([d_459_v12_])
                d_461_v13_ = ((d_461_v13_) | (d_461_v13_)).intersection(d_461_v13_)
            d_462_v14_: _dafny.MultiSet
            d_462_v14_ = _dafny.MultiSet([367, default__.fm33((self).f12, (self).f12, (d_447_v0_)[default__.safeIndex(609, (d_447_v0_).length(0))], (self).f12, globalState)])
            d_463_v15_: _dafny.Array
            nw52_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 14)
            d_463_v15_ = nw52_
            index55_ = default__.safeIndex(272, (d_463_v15_).length(0))
            (d_463_v15_)[index55_] = ((self).f13) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ejfjtu")))
            index56_ = default__.safeIndex(272, (d_463_v15_).length(0))
            rhs61_ = ((d_462_v14_).intersection(_dafny.MultiSet([d_448_v1_, d_448_v1_]))).intersection(d_462_v14_)
            rhs62_ = (d_447_v0_)[default__.safeIndex(609, (d_447_v0_).length(0))]
            rhs63_ = ((self).f13 if True else ((self).f13 if (self).f12 else (self).f13))
            lhs45_ = d_463_v15_
            lhs46_ = default__.safeIndex(272, (d_463_v15_).length(0))
            d_462_v14_ = rhs61_
            d_448_v1_ = rhs62_
            lhs45_[lhs46_] = rhs63_
            r1 = False
        elif True:
            d_464_v16_: D3
            d_464_v16_ = D3_DC7(d_447_v0_)
            d_464_v16_ = ((D3_DC7(d_447_v0_) if (self).f12 else d_464_v16_) if not ((self).f12) or ((self).f12) else d_464_v16_)
            d_465_v17_: _dafny.Set
            d_465_v17_ = _dafny.Set({d_448_v1_, default__.fm33((self).f12, (self).f12, (d_447_v0_)[default__.safeIndex(609, (d_447_v0_).length(0))], (self).f12, globalState)})
            d_466_v18_: _dafny.Map
            d_466_v18_ = _dafny.Map({(d_465_v17_).isdisjoint(d_465_v17_): (d_448_v1_) - ((d_447_v0_)[default__.safeIndex(609, (d_447_v0_).length(0))])})
            d_466_v18_ = (d_466_v18_).set((self).f12, (d_447_v0_)[default__.safeIndex(609, (d_447_v0_).length(0))])
            index57_ = default__.safeIndex(609, (d_447_v0_).length(0))
            (d_447_v0_)[index57_] = (0) - (d_448_v1_)
            d_467_v19_: _dafny.Array
            nw53_ = _dafny.Array(False, 6)
            d_467_v19_ = nw53_
            index58_ = default__.safeIndex(193, (d_467_v19_).length(0))
            (d_467_v19_)[index58_] = (self).f12
            index59_ = default__.safeIndex(193, (d_467_v19_).length(0))
            (d_467_v19_)[index59_] = (self).f12
            d_468_v20_: _dafny.Array
            nw54_ = _dafny.Array(None, 4)
            nw54_[int(0)] = d_447_v0_
            nw54_[int(1)] = d_447_v0_
            nw54_[int(2)] = d_447_v0_
            nw54_[int(3)] = d_447_v0_
            d_468_v20_ = nw54_
            nw55_ = _dafny.Array(None, 28)
            nw55_[int(0)] = (d_464_v16_).cf11
            nw55_[int(1)] = d_447_v0_
            nw55_[int(2)] = d_447_v0_
            nw55_[int(3)] = d_447_v0_
            nw55_[int(4)] = d_447_v0_
            nw55_[int(5)] = d_447_v0_
            nw55_[int(6)] = d_447_v0_
            nw55_[int(7)] = d_447_v0_
            nw55_[int(8)] = d_447_v0_
            nw55_[int(9)] = d_447_v0_
            nw55_[int(10)] = d_447_v0_
            nw55_[int(11)] = d_447_v0_
            nw55_[int(12)] = d_447_v0_
            nw55_[int(13)] = d_447_v0_
            nw55_[int(14)] = (d_447_v0_ if (d_467_v19_)[default__.safeIndex(193, (d_467_v19_).length(0))] else d_447_v0_)
            nw55_[int(15)] = d_447_v0_
            nw55_[int(16)] = d_447_v0_
            nw55_[int(17)] = d_447_v0_
            nw55_[int(18)] = (d_447_v0_ if (d_467_v19_)[default__.safeIndex(193, (d_467_v19_).length(0))] else d_447_v0_)
            nw55_[int(19)] = d_447_v0_
            nw55_[int(20)] = (d_447_v0_ if (d_467_v19_)[default__.safeIndex(193, (d_467_v19_).length(0))] else d_447_v0_)
            nw55_[int(21)] = d_447_v0_
            nw55_[int(22)] = d_447_v0_
            nw55_[int(23)] = d_447_v0_
            nw55_[int(24)] = d_447_v0_
            nw55_[int(25)] = d_447_v0_
            nw55_[int(26)] = d_447_v0_
            nw55_[int(27)] = (D3_DC7(d_447_v0_)).cf11
            d_468_v20_ = nw55_
        index60_ = default__.safeIndex(609, (d_447_v0_).length(0))
        (d_447_v0_)[index60_] = d_448_v1_
        index61_ = default__.safeIndex(609, (d_447_v0_).length(0))
        (d_447_v0_)[index61_] = d_448_v1_
        r0 = (((self).f13) + (((self).f13) + ((self).f13))).set(default__.safeIndex(-935, len(((self).f13) + (((self).f13) + ((self).f13)))), d_450_v3_)
        r1 = ((self).f12 if (self).f12 else (self).f12)
        d_469_v21_: _dafny.Set
        d_469_v21_ = _dafny.Set({(self).fm32(d_448_v1_, d_450_v3_, globalState), (self).f12, False})
        d_470_v22_: _dafny.Seq
        d_470_v22_ = _dafny.SeqWithoutIsStrInference([d_469_v21_])
        d_471_v23_: _dafny.Map
        d_471_v23_ = _dafny.Map({d_447_v0_: (d_447_v0_)[default__.safeIndex(609, (d_447_v0_).length(0))]})
        r2 = (_dafny.SeqWithoutIsStrInference([d_469_v21_, (d_470_v22_)[default__.safeIndex(d_448_v1_, len(d_470_v22_))]])) + (_dafny.SeqWithoutIsStrInference([d_469_v21_, d_469_v21_, default__.fm34(len(d_471_v23_), (self).f12, (self).f12, (self).f12, globalState), _dafny.Set({False})]))
        r3 = ((self).f13) + ((self).f13)
        return r0, r1, r2, r3

    def m1(self, p0, globalState):
        r0: _dafny.Map = _dafny.Map({})
        r1: bool = False
        d_472_v0_: int
        d_472_v0_ = 776
        d_473_v1_: _dafny.Map
        d_473_v1_ = _dafny.Map({d_472_v0_: (self).f12})
        d_474_v2_: D0
        d_474_v2_ = D0_DC1(d_472_v0_, (self).fm32(len(d_473_v1_), p0, globalState), (0) - (d_472_v0_))
        d_475_v4_: _dafny.Map
        def iife41_():
            coll29_ = _dafny.Map()
            compr_29_: int
            for compr_29_ in _dafny.IntegerRange(7, 695):
                d_476_v3_: int = compr_29_
                if ((7) <= (d_476_v3_)) and ((d_476_v3_) < (695)):
                    coll29_[(d_476_v3_) * (d_472_v0_)] = d_472_v0_
            return _dafny.Map(coll29_)
        d_475_v4_ = _dafny.Map({(d_472_v0_) == ((d_474_v2_).cf3): iife41_()
        })
        d_477_v5_: _dafny.Map
        d_477_v5_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([674 for d_478_i0_ in range(default__.abs(620))])): 54})
        d_475_v4_ = _dafny.Map({not((self).fm32(d_472_v0_, p0, globalState)): d_477_v5_})
        d_479_v6_: C0
        nw56_ = C0()
        nw56_.ctor__()
        d_479_v6_ = nw56_
        d_480_v7_: str
        d_480_v7_ = _dafny.CodePoint('i')
        pat_let_tv6_ = p0
        pat_let_tv7_ = p0
        pat_let_tv8_ = p0
        def lambda10_(source12_):
            if source12_.is_DC8:
                d_481___mcc_h0_ = source12_.cf12
                d_482___mcc_h1_ = source12_.cf13
                d_483_cf13_ = d_482___mcc_h1_
                d_484_cf12_ = d_481___mcc_h0_
                if d_483_cf13_:
                    return _dafny.CodePoint('o')
                elif True:
                    return pat_let_tv6_
            elif source12_.is_DC7:
                d_485___mcc_h2_ = source12_.cf11
                d_486_cf11_ = d_485___mcc_h2_
                return pat_let_tv7_
            elif True:
                d_487___mcc_h3_ = source12_.cf14
                d_488_cf14_ = d_487___mcc_h3_
                return (pat_let_tv8_)

        d_480_v7_ = lambda10_(default__.fm35((self).f12, globalState))
        d_489_v8_: _dafny.MultiSet
        d_489_v8_ = _dafny.MultiSet([(self).f12])
        d_490_v9_: D1
        d_490_v9_ = D1_DC4((d_479_v6_).fm5((self).f12, d_472_v0_, d_472_v0_, d_489_v8_, globalState), (self).f12, ((d_489_v8_).cardinality if (self).f12 else d_472_v0_))
        source13_ = d_490_v9_
        if source13_.is_DC3:
            d_491___mcc_h4_ = source13_.cf5
            d_492_cf5_ = d_491___mcc_h4_
            d_493_v10_: _dafny.MultiSet
            d_493_v10_ = _dafny.MultiSet([d_472_v0_])
            d_494_v11_: _dafny.Map
            d_494_v11_ = _dafny.Map({(self).f12: d_493_v10_})
            d_495_v12_: _dafny.Seq
            d_495_v12_ = _dafny.SeqWithoutIsStrInference([d_472_v0_])
            d_494_v11_ = (d_494_v11_).set(not ((self).f12) or (True), (_dafny.MultiSet(d_495_v12_)) - (d_493_v10_))
            (globalState).f8 = (924) != ((d_489_v8_).cardinality)
            d_496_v13_: _dafny.Array
            nw57_ = _dafny.Array(None, 16)
            nw57_[int(0)] = (self).f12
            nw57_[int(1)] = (self).f12
            nw57_[int(2)] = (self).f12
            nw57_[int(3)] = (self).f12
            nw57_[int(4)] = (self).f12
            nw57_[int(5)] = (len((self).f13)) < (d_472_v0_)
            nw57_[int(6)] = (self).f12
            nw57_[int(7)] = False
            nw57_[int(8)] = ((self).f12 if (self).f12 else (self).f12)
            nw57_[int(9)] = (self).f12
            nw57_[int(10)] = (self).f12
            nw57_[int(11)] = (self).f12
            nw57_[int(12)] = (self).f12
            nw57_[int(13)] = not(((self).f12 if (self).f12 else (self).f12))
            nw57_[int(14)] = not(not ((d_479_v6_).fm5((self).f12, d_472_v0_, d_472_v0_, d_489_v8_, globalState)) or ((self).f12))
            nw57_[int(15)] = ((self).f12) or ((self).f12)
            d_496_v13_ = nw57_
            index62_ = default__.safeIndex(257, (d_496_v13_).length(0))
            (d_496_v13_)[index62_] = (self).f12
            index63_ = default__.safeIndex(257, (d_496_v13_).length(0))
            (d_496_v13_)[index63_] = not((self).f12)
            d_497_v14_: _dafny.Seq
            d_497_v14_ = _dafny.SeqWithoutIsStrInference([(d_496_v13_)[default__.safeIndex(257, (d_496_v13_).length(0))], (self).f12])
            d_480_v7_ = (D4_DC12(True, len(d_497_v14_), default__.fm36(globalState), 739, d_472_v0_)).cf21
        elif source13_.is_DC4:
            d_498___mcc_h5_ = source13_.cf6
            d_499___mcc_h6_ = source13_.cf7
            d_500___mcc_h7_ = source13_.cf8
            d_501_cf8_ = d_500___mcc_h7_
            d_502_cf7_ = d_499___mcc_h6_
            d_503_cf6_ = d_498___mcc_h5_
            d_501_cf8_ = -389
            d_504_v15_: _dafny.Set
            d_504_v15_ = _dafny.Set({True, d_503_cf6_})
            d_502_cf7_ = (d_503_cf6_) in (d_504_v15_)
            d_479_v6_ = d_479_v6_
            d_505_v16_: D0
            d_505_v16_ = D0_DC0((self).f12)
            d_506_v17_: _dafny.Seq
            d_506_v17_ = _dafny.SeqWithoutIsStrInference([d_502_cf7_, (d_505_v16_).cf0])
            if (d_503_cf6_) or ((_dafny.MultiSet(d_506_v17_)).issubset(d_489_v8_)):
                r1 = ((d_504_v15_).intersection(_dafny.Set({d_502_cf7_}))).issubset(d_504_v15_)
                (globalState).f8 = d_503_cf6_
                d_507_v18_: _dafny.Seq
                d_507_v18_ = _dafny.SeqWithoutIsStrInference([d_472_v0_, d_472_v0_])
                d_472_v0_ = (d_501_cf8_) - (len(d_507_v18_))
                d_508_v19_: C0
                nw58_ = C0()
                nw58_.ctor__()
                d_508_v19_ = nw58_
                rhs64_ = d_502_cf7_
                rhs65_ = d_502_cf7_
                d_503_cf6_ = rhs64_
                r1 = rhs65_
            elif True:
                d_509_v20_: _dafny.Seq
                d_509_v20_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xrh"))
                d_509_v20_ = d_509_v20_
                d_501_cf8_ = ((0) - (d_501_cf8_)) - (((d_477_v5_)[d_472_v0_] if (d_472_v0_) in (d_477_v5_) else d_501_cf8_))
                d_510_v21_: _dafny.Array
                nw59_ = _dafny.Array(None, 3)
                nw59_[int(0)] = d_479_v6_
                nw59_[int(1)] = d_479_v6_
                nw59_[int(2)] = d_479_v6_
                d_510_v21_ = nw59_
                index64_ = default__.safeIndex(222, (d_510_v21_).length(0))
                (d_510_v21_)[index64_] = d_479_v6_
                index65_ = default__.safeIndex(222, (d_510_v21_).length(0))
                (d_510_v21_)[index65_] = d_479_v6_
                (globalState).f8 = d_503_cf6_
                d_511_v22_: _dafny.Seq
                d_511_v22_ = _dafny.SeqWithoutIsStrInference([d_501_cf8_, (default__.fm33(d_503_cf6_, (self).f12, d_501_cf8_, (self).f12, globalState)) - (d_501_cf8_), default__.safeModuloInt(len(d_509_v20_), d_472_v0_)])
                d_511_v22_ = default__.fm37(globalState)
        elif source13_.is_DC2:
            d_512___mcc_h8_ = source13_.cf4
            d_513_cf4_ = d_512___mcc_h8_
            d_480_v7_ = d_480_v7_
            (globalState).f8 = ((self).f12) == ((self).f12)
            (globalState).f8 = (self).f12
            if (self).f12:
                d_514_v23_: _dafny.Seq
                d_514_v23_ = _dafny.SeqWithoutIsStrInference([True, (self).f12])
                d_515_v24_: _dafny.Map
                d_515_v24_ = _dafny.Map({not((self).f12): (self).f12})
                d_516_v25_: _dafny.Seq
                d_516_v25_ = _dafny.SeqWithoutIsStrInference([d_514_v23_])
                d_517_v26_: _dafny.Array
                nw60_ = _dafny.Array(None, 18)
                nw60_[int(0)] = (d_514_v23_) + (d_514_v23_)
                nw60_[int(1)] = default__.fm38(((d_515_v24_)[(self).f12] if ((self).f12) in (d_515_v24_) else (self).f12), globalState)
                nw60_[int(2)] = d_514_v23_
                nw60_[int(3)] = (d_514_v23_) + ((d_516_v25_)[default__.safeIndex(d_472_v0_, len(d_516_v25_))])
                nw60_[int(4)] = (d_514_v23_ if (self).f12 else d_514_v23_)
                nw60_[int(5)] = d_514_v23_
                nw60_[int(6)] = _dafny.SeqWithoutIsStrInference([(self).f12, not((self).f12)])
                nw60_[int(7)] = d_514_v23_
                nw60_[int(8)] = (d_514_v23_) + (d_514_v23_)
                nw60_[int(9)] = (d_514_v23_) + (d_514_v23_)
                nw60_[int(10)] = (d_514_v23_) + (d_514_v23_)
                nw60_[int(11)] = d_514_v23_
                nw60_[int(12)] = default__.fm38((self).f12, globalState)
                nw60_[int(13)] = d_514_v23_
                nw60_[int(14)] = (d_514_v23_) + (d_514_v23_)
                nw60_[int(15)] = d_514_v23_
                nw60_[int(16)] = _dafny.SeqWithoutIsStrInference([(self).f12, not((d_479_v6_).fm5((self).f12, d_472_v0_, d_472_v0_, d_489_v8_, globalState)), (self).f12])
                nw60_[int(17)] = d_514_v23_
                d_517_v26_ = nw60_
                index66_ = default__.safeIndex(226, (d_517_v26_).length(0))
                (d_517_v26_)[index66_] = (d_514_v23_).set(default__.safeIndex((0) - (d_472_v0_), len(d_514_v23_)), (self).f12)
                index67_ = default__.safeIndex(226, (d_517_v26_).length(0))
                (d_517_v26_)[index67_] = ((d_514_v23_) + (d_514_v23_)).set(default__.safeIndex(d_472_v0_, len((d_514_v23_) + (d_514_v23_))), (self).f12)
                (globalState).f8 = ((d_517_v26_)[default__.safeIndex(226, (d_517_v26_).length(0))])[default__.safeIndex(d_472_v0_, len((d_517_v26_)[default__.safeIndex(226, (d_517_v26_).length(0))]))]
                (globalState).f8 = (self).f12
                d_518_v27_: C0
                nw61_ = C0()
                nw61_.ctor__()
                d_518_v27_ = nw61_
                d_472_v0_ = d_472_v0_
            elif True:
                d_480_v7_ = _dafny.CodePoint('m')
                d_472_v0_ = (d_472_v0_ if (d_472_v0_) == (d_472_v0_) else d_472_v0_)
                d_519_v28_: _dafny.Array
                def lambda11_(d_520_v0_):
                    def lambda12_(d_521_i1_):
                        return (d_521_i1_) + (d_520_v0_)

                    return lambda12_

                init6_ = lambda11_(d_472_v0_)
                nw62_ = _dafny.Array(None, 21)
                for i0_6_ in range(nw62_.length(0)):
                    nw62_[i0_6_] = init6_(i0_6_)
                d_519_v28_ = nw62_
                d_522_v29_: _dafny.Map
                d_522_v29_ = _dafny.Map({(self).f12: d_519_v28_})
                d_523_v30_: _dafny.Array
                nw63_ = _dafny.Array(None, 16)
                nw63_[int(0)] = d_519_v28_
                nw63_[int(1)] = d_519_v28_
                nw63_[int(2)] = d_519_v28_
                nw63_[int(3)] = d_519_v28_
                nw63_[int(4)] = d_519_v28_
                nw63_[int(5)] = d_519_v28_
                nw63_[int(6)] = d_519_v28_
                nw63_[int(7)] = d_519_v28_
                nw63_[int(8)] = d_519_v28_
                nw63_[int(9)] = d_519_v28_
                nw63_[int(10)] = d_519_v28_
                nw63_[int(11)] = d_519_v28_
                nw63_[int(12)] = d_519_v28_
                nw63_[int(13)] = ((d_522_v29_)[(self).f12] if ((self).f12) in (d_522_v29_) else d_519_v28_)
                nw63_[int(14)] = d_519_v28_
                nw63_[int(15)] = d_519_v28_
                d_523_v30_ = nw63_
                index68_ = default__.safeIndex(96, (d_523_v30_).length(0))
                (d_523_v30_)[index68_] = d_519_v28_
                index69_ = default__.safeIndex(96, (d_523_v30_).length(0))
                (d_523_v30_)[index69_] = d_519_v28_
                d_524_v31_: _dafny.Seq
                d_524_v31_ = _dafny.SeqWithoutIsStrInference([(self).f12, not((self).f12)])
                d_525_v32_: _dafny.Map
                d_525_v32_ = _dafny.Map({(self).f12: (self).f12})
                d_526_v33_: D6
                d_526_v33_ = D6_DC16(d_525_v32_)
                d_527_v34_: _dafny.Map
                d_527_v34_ = _dafny.Map({(d_524_v31_)[default__.safeIndex(len((d_526_v33_).cf28), len(d_524_v31_))]: -724})
                d_527_v34_ = (d_527_v34_).set(not((self).f12), d_472_v0_)
                (globalState).f8 = (d_472_v0_) < (d_472_v0_)
        elif True:
            d_528___mcc_h9_ = source13_.cf9
            d_529_cf9_ = d_528___mcc_h9_
            d_472_v0_ = d_472_v0_
            r1 = False
            d_530_v35_: D7
            d_530_v35_ = D7_DC21(_dafny.Set({(self).f12}))
            d_531_v36_: _dafny.Set
            d_531_v36_ = _dafny.Set({(self).f12, (self).f12, (self).f12})
            d_532_v37_: _dafny.MultiSet
            d_532_v37_ = _dafny.MultiSet([d_472_v0_, -245, len(((d_530_v35_).cf40) | (d_531_v36_)), d_472_v0_, d_472_v0_])
            d_532_v37_ = (d_532_v37_) - (d_532_v37_)
            if (d_472_v0_) != (d_472_v0_):
                def iife42_():
                    coll30_ = _dafny.Map()
                    compr_30_: int
                    for compr_30_ in (d_532_v37_).Elements:
                        d_533_v38_: int = compr_30_
                        if (d_533_v38_) in (d_532_v37_):
                            coll30_[default__.safeModuloInt(d_533_v38_, len((self).f13))] = (self).f12
                    return _dafny.Map(coll30_)
                d_473_v1_ = (d_473_v1_) | ((iife42_()
                ).set(d_472_v0_, (self).f12))
                d_534_v39_: _dafny.Array
                nw64_ = _dafny.Array(None, 26)
                nw64_[int(0)] = d_480_v7_
                nw64_[int(1)] = p0
                nw64_[int(2)] = d_480_v7_
                nw64_[int(3)] = d_480_v7_
                nw64_[int(4)] = p0
                nw64_[int(5)] = p0
                nw64_[int(6)] = p0
                nw64_[int(7)] = d_480_v7_
                nw64_[int(8)] = d_480_v7_
                nw64_[int(9)] = p0
                nw64_[int(10)] = d_480_v7_
                nw64_[int(11)] = p0
                nw64_[int(12)] = _dafny.CodePoint('v')
                nw64_[int(13)] = _dafny.CodePoint('m')
                nw64_[int(14)] = d_480_v7_
                nw64_[int(15)] = p0
                nw64_[int(16)] = p0
                nw64_[int(17)] = d_480_v7_
                nw64_[int(18)] = p0
                nw64_[int(19)] = p0
                nw64_[int(20)] = p0
                nw64_[int(21)] = p0
                nw64_[int(22)] = d_480_v7_
                nw64_[int(23)] = p0
                nw64_[int(24)] = d_480_v7_
                nw64_[int(25)] = d_480_v7_
                d_534_v39_ = nw64_
                index70_ = default__.safeIndex(62, (d_534_v39_).length(0))
                (d_534_v39_)[index70_] = d_480_v7_
                index71_ = default__.safeIndex(62, (d_534_v39_).length(0))
                (d_534_v39_)[index71_] = _dafny.CodePoint('b')
                d_535_v40_: _dafny.Seq
                d_535_v40_ = _dafny.SeqWithoutIsStrInference([-832])
                rhs66_ = (self).f12
                rhs67_ = (len(d_531_v36_)) * (default__.safeModuloInt(d_472_v0_, len(d_535_v40_)))
                lhs47_ = globalState
                lhs47_.f8 = rhs66_
                d_472_v0_ = rhs67_
                d_536_v41_: _dafny.Seq
                d_536_v41_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rax"))
                d_536_v41_ = (self).f13
                d_473_v1_ = (d_473_v1_).set(d_472_v0_, True)
            elif True:
                d_537_v42_: C0
                nw65_ = C0()
                nw65_.ctor__()
                d_537_v42_ = nw65_
                d_538_v43_: _dafny.Seq
                d_538_v43_ = _dafny.SeqWithoutIsStrInference([(self).f12])
                d_538_v43_ = d_538_v43_
                d_539_v44_: _dafny.Seq
                d_539_v44_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "feeckpr"))
                d_539_v44_ = _dafny.SeqWithoutIsStrInference([p0 for d_540_i2_ in range(default__.abs(9))])
                d_541_v45_: _dafny.Map
                d_541_v45_ = _dafny.Map({(self).f12: (self).f12})
                d_542_v46_: _dafny.Set
                d_542_v46_ = _dafny.Set({default__.safeModuloInt(-682, default__.fm33((self).f12, (self).f12, len(d_541_v45_), (D1_DC4((self).f12, (self).f12, d_472_v0_)).cf6, globalState))})
                def iife43_():
                    coll31_ = _dafny.Set()
                    compr_31_: int
                    for compr_31_ in _dafny.IntegerRange(825, 781):
                        d_543_v47_: int = compr_31_
                        if ((825) <= (d_543_v47_)) and ((d_543_v47_) < (781)):
                            coll31_ = coll31_.union(_dafny.Set([(d_543_v47_) * (len(d_538_v43_))]))
                    return _dafny.Set(coll31_)
                d_542_v46_ = iife43_()
                
                d_544_v48_: _dafny.Array
                nw66_ = _dafny.Array(False, 14)
                d_544_v48_ = nw66_
                d_544_v48_ = d_544_v48_
        d_545_v49_: C0
        nw67_ = C0()
        nw67_.ctor__()
        d_545_v49_ = nw67_
        if (self).f12:
            if (default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xplooehy"))), d_472_v0_)) < (d_472_v0_):
                d_546_v50_: D3
                d_546_v50_ = D3_DC8((self).f12, (self).f12)
                def iife44_(_pat_let6_0):
                    def iife45_(d_547_dt__update__tmp_h0_):
                        def iife46_(_pat_let7_0):
                            def iife47_(d_548_dt__update_hcf12_h0_):
                                return D3_DC8(d_548_dt__update_hcf12_h0_, (d_547_dt__update__tmp_h0_).cf13)
                            return iife47_(_pat_let7_0)
                        return iife46_((self).f12)
                    return iife45_(_pat_let6_0)
                d_546_v50_ = (d_546_v50_ if (self).f12 else iife44_(d_546_v50_))
                d_549_v51_: _dafny.Seq
                d_549_v51_ = _dafny.SeqWithoutIsStrInference([d_473_v1_])
                d_550_v52_: _dafny.Map
                d_550_v52_ = _dafny.Map({((d_549_v51_)[default__.safeIndex(len((self).f13), len(d_549_v51_))]).set(-24, (self).f12): d_472_v0_})
                d_472_v0_ = len((d_550_v52_).set(d_473_v1_, (d_472_v0_ if (self).f12 else -270)))
                (globalState).f8 = (d_479_v6_).fm5((self).f12, (0) - (d_472_v0_), default__.safeDivisionInt(d_472_v0_, d_472_v0_), d_489_v8_, globalState)
                d_551_v53_: C0
                nw68_ = C0()
                nw68_.ctor__()
                d_551_v53_ = nw68_
                d_552_v55_: _dafny.Set
                d_552_v55_ = _dafny.Set({default__.fm33((self).f12, (self).f12, 953, (self).f12, globalState), d_472_v0_, d_472_v0_, d_472_v0_, d_472_v0_})
                d_553_v56_: _dafny.MultiSet
                def iife48_():
                    coll32_ = _dafny.Set()
                    compr_32_: int
                    for compr_32_ in _dafny.IntegerRange(954, 800):
                        d_554_v54_: int = compr_32_
                        if ((954) <= (d_554_v54_)) and ((d_554_v54_) < (800)):
                            coll32_ = coll32_.union(_dafny.Set([(d_554_v54_) + (((d_477_v5_)[d_472_v0_] if (d_472_v0_) in (d_477_v5_) else d_472_v0_))]))
                    return _dafny.Set(coll32_)
                d_553_v56_ = _dafny.MultiSet([iife48_()
                , d_552_v55_, d_552_v55_])
                d_553_v56_ = (d_553_v56_) | (d_553_v56_)
            elif True:
                d_555_v57_: _dafny.Array
                def lambda13_(d_556_i3_):
                    return True

                init7_ = lambda13_
                nw69_ = _dafny.Array(None, 24)
                for i0_7_ in range(nw69_.length(0)):
                    nw69_[i0_7_] = init7_(i0_7_)
                d_555_v57_ = nw69_
                d_557_v58_: _dafny.Array
                nw70_ = _dafny.Array(int(0), 28)
                d_557_v58_ = nw70_
                d_558_v59_: D8
                d_558_v59_ = D8_DC23(_dafny.SeqWithoutIsStrInference([(self).f12, (self).f12]))
                index72_ = default__.safeIndex(319, (d_557_v58_).length(0))
                (d_557_v58_)[index72_] = (_dafny.MultiSet((d_558_v59_).cf41)).cardinality
                d_559_v60_: _dafny.Array
                nw71_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 1)
                d_559_v60_ = nw71_
                d_560_v61_: _dafny.Map
                d_560_v61_ = _dafny.Map({d_559_v60_: (d_472_v0_) * (default__.fm33((self).f12, (self).fm32(d_472_v0_, d_480_v7_, globalState), d_472_v0_, True, globalState))})
                index73_ = default__.safeIndex(319, (d_557_v58_).length(0))
                rhs68_ = d_555_v57_
                rhs69_ = (self).f12
                rhs70_ = ((d_560_v61_)[d_559_v60_] if (d_559_v60_) in (d_560_v61_) else d_472_v0_)
                lhs48_ = d_557_v58_
                lhs49_ = default__.safeIndex(319, (d_557_v58_).length(0))
                d_555_v57_ = rhs68_
                r1 = rhs69_
                lhs48_[lhs49_] = rhs70_
                d_561_v62_: _dafny.Set
                d_561_v62_ = _dafny.Set({d_557_v58_, d_557_v58_})
                d_562_v63_: D8
                d_562_v63_ = D8_DC24((self).f12, d_561_v62_, (-247) * ((d_557_v58_)[default__.safeIndex(319, (d_557_v58_).length(0))]))
                d_562_v63_ = d_562_v63_
                d_563_v65_: _dafny.MultiSet
                d_563_v65_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "igpmkqt"))), (d_557_v58_)[default__.safeIndex(319, (d_557_v58_).length(0))]])
                d_564_v66_: _dafny.Seq
                def iife49_():
                    coll33_ = _dafny.Map()
                    compr_33_: int
                    for compr_33_ in (d_563_v65_).Elements:
                        d_565_v64_: int = compr_33_
                        if (d_565_v64_) in (d_563_v65_):
                            coll33_[(d_565_v64_) * ((d_557_v58_)[default__.safeIndex(319, (d_557_v58_).length(0))])] = (d_557_v58_)[default__.safeIndex(319, (d_557_v58_).length(0))]
                    return _dafny.Map(coll33_)
                d_564_v66_ = _dafny.SeqWithoutIsStrInference([default__.fm39((self).f12, (0) - ((d_557_v58_)[default__.safeIndex(319, (d_557_v58_).length(0))]), True, globalState), iife49_()
                , d_477_v5_, _dafny.Map({(d_557_v58_)[default__.safeIndex(319, (d_557_v58_).length(0))]: 90})])
                index74_ = default__.safeIndex(478, (d_555_v57_).length(0))
                (d_555_v57_)[index74_] = (False) or (True)
                d_566_v67_: _dafny.Seq
                d_566_v67_ = _dafny.SeqWithoutIsStrInference([d_472_v0_])
                d_567_v68_: D5
                d_567_v68_ = D5_DC14((d_557_v58_)[default__.safeIndex(319, (d_557_v58_).length(0))], (self).f12)
                d_568_v69_: _dafny.Set
                d_568_v69_ = _dafny.Set({(self).f12, (self).f12})
                d_569_v70_: _dafny.Seq
                d_569_v70_ = _dafny.SeqWithoutIsStrInference([(d_545_v49_).fm5((self).f12, (d_557_v58_)[default__.safeIndex(319, (d_557_v58_).length(0))], len(d_568_v69_), d_489_v8_, globalState)])
                d_570_v71_: _dafny.Seq
                d_570_v71_ = _dafny.SeqWithoutIsStrInference([d_569_v70_])
                index75_ = default__.safeIndex(478, (d_555_v57_).length(0))
                rhs71_ = (d_557_v58_)[default__.safeIndex(319, (d_557_v58_).length(0))]
                rhs72_ = default__.fm40(d_566_v67_, d_567_v68_, len((self).f13), globalState)
                rhs73_ = ((d_569_v70_) + (d_569_v70_)) <= ((d_570_v71_)[default__.safeIndex(-736, len(d_570_v71_))])
                lhs50_ = d_555_v57_
                lhs51_ = default__.safeIndex(478, (d_555_v57_).length(0))
                d_472_v0_ = rhs71_
                d_564_v66_ = rhs72_
                lhs50_[lhs51_] = rhs73_
                index76_ = default__.safeIndex(319, (d_557_v58_).length(0))
                (d_557_v58_)[index76_] = default__.safeDivisionInt(135, d_472_v0_)
                d_571_v72_: C0
                nw72_ = C0()
                nw72_.ctor__()
                d_571_v72_ = nw72_
            d_572_v73_: _dafny.Array
            def lambda14_(d_573_i4_):
                return (self).f12

            init8_ = lambda14_
            nw73_ = _dafny.Array(None, 17)
            for i0_8_ in range(nw73_.length(0)):
                nw73_[i0_8_] = init8_(i0_8_)
            d_572_v73_ = nw73_
            index77_ = default__.safeIndex(125, (d_572_v73_).length(0))
            (d_572_v73_)[index77_] = (self).f12
            index78_ = default__.safeIndex(125, (d_572_v73_).length(0))
            (d_572_v73_)[index78_] = (self).f12
            d_574_v74_: C0
            nw74_ = C0()
            nw74_.ctor__()
            d_574_v74_ = nw74_
            d_575_v75_: _dafny.Seq
            d_575_v75_ = _dafny.SeqWithoutIsStrInference([len(((self).f13).set(default__.safeIndex(d_472_v0_, len((self).f13)), p0)), d_472_v0_])
            d_576_v76_: _dafny.Map
            d_576_v76_ = _dafny.Map({d_472_v0_: d_575_v75_})
            d_577_v77_: _dafny.Map
            d_577_v77_ = _dafny.Map({(self).f13: len((d_575_v75_) + (((d_576_v76_)[d_472_v0_] if (d_472_v0_) in (d_576_v76_) else d_575_v75_)))})
            d_577_v77_ = (d_577_v77_).set((self).f13, (d_472_v0_) + (len(d_575_v75_)))
            d_472_v0_ = (0) - ((144) + (d_472_v0_))
        elif True:
            d_578_v78_: _dafny.Seq
            d_578_v78_ = _dafny.SeqWithoutIsStrInference([(self).f12, (self).f12])
            r1 = ((d_578_v78_)[default__.safeIndex(d_472_v0_, len(d_578_v78_))]) or ((p0) not in ((self).f13))
            d_579_v79_: D6
            d_579_v79_ = D6_DC18(p0, d_472_v0_, (d_472_v0_) - (d_472_v0_))
            d_579_v79_ = d_579_v79_
            r1 = ((d_473_v1_)[((d_477_v5_)[d_472_v0_] if (d_472_v0_) in (d_477_v5_) else 142)] if (((d_477_v5_)[d_472_v0_] if (d_472_v0_) in (d_477_v5_) else 142)) in (d_473_v1_) else (self).f12)
            d_580_v80_: C0
            nw75_ = C0()
            nw75_.ctor__()
            d_580_v80_ = nw75_
            d_480_v7_ = d_480_v7_
        d_581_v81_: _dafny.Array
        def lambda15_(d_582_v0_):
            def lambda16_(d_583_i5_):
                return (d_583_i5_) * (d_582_v0_)

            return lambda16_

        init9_ = lambda15_(d_472_v0_)
        nw76_ = _dafny.Array(None, 12)
        for i0_9_ in range(nw76_.length(0)):
            nw76_[i0_9_] = init9_(i0_9_)
        d_581_v81_ = nw76_
        d_584_v82_: _dafny.Map
        d_584_v82_ = _dafny.Map({d_581_v81_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lhau"))})
        r0 = d_584_v82_
        r1 = ((p0 if (self).f12 else p0)) in ((self).f13)
        return r0, r1

    @property
    def f12(self):
        return self._f12
    @property
    def f13(self):
        return self._f13

class C2(T1, T0):
    def  __init__(self):
        self._f9: _dafny.Map = _dafny.Map({})
        self._f10: bool = False
        self.f14: bool = False
        self._f15: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    @property
    def f9(self):
        return self._f9
    @property
    def f10(self):
        return self._f10
    def ctor__(self, f14, f15, f9, f10):
        (self).f14 = f14
        (self)._f15 = f15
        (self)._f9 = f9
        (self)._f10 = f10

    def fm0(self, p0, p1, p2, p3, globalState):
        def iife50_():
            coll34_ = _dafny.Map()
            compr_34_: int
            for compr_34_ in (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dnq"))): (self).f9})).keys.Elements:
                d_585_v0_: int = compr_34_
                if (d_585_v0_) in (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dnq"))): (self).f9})):
                    coll34_[(d_585_v0_) * (688)] = (D0_DC1(len(_dafny.Set({self.f14, (self).f10})), self.f14, (0) - (len(_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_586_i0_ in range(default__.abs(289))]))}))))).cf3
            return _dafny.Map(coll34_)
        return (len((iife50_()
        ))) != (301)

    def fm1(self, p0, globalState):
        if (len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bnhgpf"))): D6_DC17((self).f10, self.f14, 55, len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xy"))): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uyxnksw")))})), len(_dafny.Map({255: self.f14})))}))) != (-987):
            return ((0) - (-395)) - (len(_dafny.Set({self.f14})))
        elif True:
            return ((0) - (len(_dafny.Map({963: self.f14})))) + (-169)

    def fm41(self, p0, p1, p2, globalState):
        def iife51_():
            coll35_ = _dafny.Map()
            compr_35_: int
            for compr_35_ in (_dafny.MultiSet([411])).Elements:
                d_587_v0_: int = compr_35_
                if (d_587_v0_) in (_dafny.MultiSet([411])):
                    coll35_[(d_587_v0_) - (396)] = 585
            return _dafny.Map(coll35_)
        return _dafny.Set({self.f14, self.f14, (_dafny.MultiSet([196, 419, len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kepfnihg"))): len(iife51_()
        )}))])) == (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xeverqrey")))])), not((289) >= (944))})

    def fm42(self, p0, p1, p2, globalState):
        return 24

    def m0(self, globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r1: bool = False
        r2: _dafny.Seq = _dafny.Seq({})
        r3: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        d_588_v0_: _dafny.MultiSet
        d_588_v0_ = _dafny.MultiSet([self.f14])
        d_589_v1_: _dafny.Seq
        d_589_v1_ = _dafny.SeqWithoutIsStrInference([self.f14])
        d_590_v2_: int
        d_590_v2_ = 689
        d_591_v3_: _dafny.Map
        d_591_v3_ = _dafny.Map({(d_588_v0_).isdisjoint(_dafny.MultiSet((d_589_v1_).set(default__.safeIndex((self).fm42(d_590_v2_, d_590_v2_, d_590_v2_, globalState), len(d_589_v1_)), (self).f10))): d_590_v2_})
        d_592_v4_: _dafny.MultiSet
        d_592_v4_ = _dafny.MultiSet([d_590_v2_])
        d_591_v3_ = (d_591_v3_).set((True if False else self.f14), (0) - (((d_592_v4_)[(0) - (d_590_v2_)] if ((0) - (d_590_v2_)) in (d_592_v4_) else d_590_v2_)))
        (self).f14 = (self).f10
        d_593_v5_: _dafny.Array
        def lambda17_(d_594_i0_):
            return _dafny.CodePoint('f')

        init10_ = lambda17_
        nw77_ = _dafny.Array(None, 22)
        for i0_10_ in range(nw77_.length(0)):
            nw77_[i0_10_] = init10_(i0_10_)
        d_593_v5_ = nw77_
        index79_ = default__.safeIndex(353, (d_593_v5_).length(0))
        (d_593_v5_)[index79_] = _dafny.CodePoint('d')
        index80_ = default__.safeIndex(59, ((self).f15).length(0))
        ((self).f15)[index80_] = (self).fm0(d_589_v1_, (self).f10, len(_dafny.SeqWithoutIsStrInference([d_590_v2_ for d_595_i1_ in range(default__.abs(-427))])), self.f14, globalState)
        d_596_v7_: _dafny.Seq
        d_596_v7_ = _dafny.SeqWithoutIsStrInference([d_590_v2_])
        d_597_v8_: D6
        def iife52_():
            coll36_ = _dafny.Map()
            compr_36_: int
            for compr_36_ in (d_596_v7_).Elements:
                d_598_v6_: int = compr_36_
                if (d_598_v6_) in (d_596_v7_):
                    coll36_[default__.safeDivisionInt(d_598_v6_, d_590_v2_)] = self.f14
            return _dafny.Map(coll36_)
        d_597_v8_ = D6_DC17(self.f14, (self).f10, -155, (0) - ((0) - ((0) - (len(iife52_()
)))), (self).fm1((self).f9, globalState))
        d_599_v9_: D5
        d_599_v9_ = D5_DC13((self).f15)
        index81_ = default__.safeIndex(353, (d_593_v5_).length(0))
        index82_ = default__.safeIndex(59, ((self).f15).length(0))
        rhs74_ = (d_597_v8_).cf33
        rhs75_ = default__.fm43((len(_dafny.SeqWithoutIsStrInference([d_599_v9_]))) - (170), globalState)
        rhs76_ = d_591_v3_
        rhs77_ = not ((self).f10) or ((self).f10)
        rhs78_ = d_590_v2_
        lhs52_ = d_593_v5_
        lhs53_ = default__.safeIndex(353, (d_593_v5_).length(0))
        lhs54_ = (self).f15
        lhs55_ = default__.safeIndex(59, ((self).f15).length(0))
        d_590_v2_ = rhs74_
        lhs52_[lhs53_] = rhs75_
        d_591_v3_ = rhs76_
        lhs54_[lhs55_] = rhs77_
        d_590_v2_ = rhs78_
        d_600_i2_: int
        d_600_i2_ = 0
        with _dafny.label("0"):
            while not(((self).f15)[default__.safeIndex(59, ((self).f15).length(0))]):
                with _dafny.c_label("0"):
                    if (d_600_i2_) >= (100):
                        raise _dafny.Break("0")
                    d_600_i2_ = (d_600_i2_) + (1)
                    (globalState).f8 = (d_590_v2_) >= (d_590_v2_)
                    r1 = ((self).f10) or ((self).f10)
                    d_590_v2_ = d_590_v2_
                    index83_ = default__.safeIndex(59, ((self).f15).length(0))
                    ((self).f15)[index83_] = (False) and (self.f14)
                    pass
            pass
        d_601_v10_: _dafny.Map
        d_601_v10_ = _dafny.Map({d_590_v2_: ((self).f15)[default__.safeIndex(59, ((self).f15).length(0))]})
        d_601_v10_ = default__.fm44(False, globalState)
        if ((self).f15)[default__.safeIndex(59, ((self).f15).length(0))]:
            d_602_v11_: _dafny.Seq
            d_602_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wem"))
            d_603_v12_: _dafny.Seq
            d_603_v12_ = _dafny.SeqWithoutIsStrInference([d_602_v11_])
            d_604_v13_: C1
            nw78_ = C1()
            nw78_.ctor__(((self).f15)[default__.safeIndex(59, ((self).f15).length(0))], (d_603_v12_)[default__.safeIndex(len(d_602_v11_), len(d_603_v12_))])
            d_604_v13_ = nw78_
            d_605_v14_: _dafny.Array
            nw79_ = _dafny.Array(_dafny.Array(None, 0), 17)
            d_605_v14_ = nw79_
            d_606_v15_: _dafny.Map
            d_606_v15_ = _dafny.Map({(self).f10: d_589_v1_})
            d_607_v16_: _dafny.Array
            nw80_ = _dafny.Array(None, 22)
            nw80_[int(0)] = (d_604_v13_).f13
            nw80_[int(1)] = d_602_v11_
            nw80_[int(2)] = default__.fm45(not(self.f14), d_590_v2_, (d_604_v13_).f12, globalState)
            nw80_[int(3)] = d_602_v11_
            nw80_[int(4)] = default__.fm45((d_589_v1_)[default__.safeIndex(d_590_v2_, len(d_589_v1_))], d_590_v2_, self.f14, globalState)
            nw80_[int(5)] = (d_604_v13_).f13
            nw80_[int(6)] = d_602_v11_
            nw80_[int(7)] = d_602_v11_
            nw80_[int(8)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cmnxwuk"))).set(default__.safeIndex(d_590_v2_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cmnxwuk")))), (d_593_v5_)[default__.safeIndex(353, (d_593_v5_).length(0))])
            nw80_[int(9)] = d_602_v11_
            nw80_[int(10)] = _dafny.SeqWithoutIsStrInference([(d_593_v5_)[default__.safeIndex(353, (d_593_v5_).length(0))] for d_608_i3_ in range(default__.abs(486))])
            nw80_[int(11)] = (d_604_v13_).f13
            nw80_[int(12)] = (d_604_v13_).f13
            nw80_[int(13)] = ((d_604_v13_).f13).set(default__.safeIndex(len(d_606_v15_), len((d_604_v13_).f13)), (d_593_v5_)[default__.safeIndex(353, (d_593_v5_).length(0))])
            nw80_[int(14)] = (d_604_v13_).f13
            nw80_[int(15)] = _dafny.SeqWithoutIsStrInference([(d_593_v5_)[default__.safeIndex(353, (d_593_v5_).length(0))] for d_609_i4_ in range(default__.abs(202))])
            nw80_[int(16)] = _dafny.SeqWithoutIsStrInference([(d_593_v5_)[default__.safeIndex(353, (d_593_v5_).length(0))] for d_610_i5_ in range(default__.abs(31))])
            nw80_[int(17)] = d_602_v11_
            nw80_[int(18)] = (((d_602_v11_).set(default__.safeIndex(len((d_604_v13_).f13), len(d_602_v11_)), _dafny.CodePoint('o'))).set(default__.safeIndex(d_590_v2_, len((d_602_v11_).set(default__.safeIndex(len((d_604_v13_).f13), len(d_602_v11_)), _dafny.CodePoint('o')))), (d_593_v5_)[default__.safeIndex(353, (d_593_v5_).length(0))])).set(default__.safeIndex(d_590_v2_, len(((d_602_v11_).set(default__.safeIndex(len((d_604_v13_).f13), len(d_602_v11_)), _dafny.CodePoint('o'))).set(default__.safeIndex(d_590_v2_, len((d_602_v11_).set(default__.safeIndex(len((d_604_v13_).f13), len(d_602_v11_)), _dafny.CodePoint('o')))), (d_593_v5_)[default__.safeIndex(353, (d_593_v5_).length(0))]))), (d_593_v5_)[default__.safeIndex(353, (d_593_v5_).length(0))])
            nw80_[int(19)] = d_602_v11_
            nw80_[int(20)] = (d_604_v13_).f13
            nw80_[int(21)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_611_i6_ in range(default__.abs(591))])
            d_607_v16_ = nw80_
            index84_ = default__.safeIndex(918, (d_605_v14_).length(0))
            (d_605_v14_)[index84_] = d_607_v16_
            index85_ = default__.safeIndex(918, (d_605_v14_).length(0))
            (d_605_v14_)[index85_] = d_607_v16_
            (globalState).f8 = (self).fm0((d_589_v1_).set(default__.safeIndex(len(d_589_v1_), len(d_589_v1_)), self.f14), (d_604_v13_).f12, d_590_v2_, (d_589_v1_) == (d_589_v1_), globalState)
            d_612_v17_: C1
            nw81_ = C1()
            nw81_.ctor__(not(self.f14), (d_604_v13_).f13)
            d_612_v17_ = nw81_
            d_613_v18_: D6
            d_613_v18_ = D6_DC18((d_593_v5_)[default__.safeIndex(353, (d_593_v5_).length(0))], d_590_v2_, (d_590_v2_) + (d_590_v2_))
            rhs79_ = d_613_v18_
            rhs80_ = (D6_DC18(((d_612_v17_).f13)[default__.safeIndex(len(d_596_v7_), len((d_612_v17_).f13))], d_590_v2_, d_590_v2_)) not in (_dafny.Set({d_613_v18_}))
            lhs56_ = self
            d_613_v18_ = rhs79_
            lhs56_.f14 = rhs80_
        elif True:
            source14_ = D1_DC5(D1_DC3(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n"))))
            if source14_.is_DC3:
                d_614___mcc_h0_ = source14_.cf5
                d_615_cf5_ = d_614___mcc_h0_
                rhs81_ = (self).f10
                rhs82_ = ((self).f15)[default__.safeIndex(59, ((self).f15).length(0))]
                rhs83_ = d_590_v2_
                lhs57_ = globalState
                lhs58_ = self
                lhs57_.f8 = rhs81_
                lhs58_.f14 = rhs82_
                d_590_v2_ = rhs83_
                d_616_v19_: C1
                nw82_ = C1()
                nw82_.ctor__(self.f14, (d_615_cf5_).set(default__.safeIndex(d_590_v2_, len(d_615_cf5_)), (d_593_v5_)[default__.safeIndex(353, (d_593_v5_).length(0))]))
                d_616_v19_ = nw82_
                d_617_v20_: _dafny.Set
                d_617_v20_ = _dafny.Set({d_590_v2_, d_590_v2_})
                index86_ = default__.safeIndex(59, ((self).f15).length(0))
                index87_ = default__.safeIndex(59, ((self).f15).length(0))
                rhs84_ = self.f14
                rhs85_ = (d_617_v20_).ispropersubset(d_617_v20_)
                lhs59_ = (self).f15
                lhs60_ = default__.safeIndex(59, ((self).f15).length(0))
                lhs61_ = (self).f15
                lhs62_ = default__.safeIndex(59, ((self).f15).length(0))
                lhs59_[lhs60_] = rhs84_
                lhs61_[lhs62_] = rhs85_
                d_618_v21_: _dafny.Array
                nw83_ = _dafny.Array(None, 23)
                nw83_[int(0)] = self.f14
                nw83_[int(1)] = (self).f10
                nw83_[int(2)] = (d_616_v19_).f12
                nw83_[int(3)] = (d_616_v19_).f12
                nw83_[int(4)] = ((self).f15)[default__.safeIndex(59, ((self).f15).length(0))]
                nw83_[int(5)] = (d_616_v19_).f12
                nw83_[int(6)] = (d_616_v19_).fm32(d_590_v2_, _dafny.CodePoint('i'), globalState)
                nw83_[int(7)] = True
                nw83_[int(8)] = ((self).f15)[default__.safeIndex(59, ((self).f15).length(0))]
                nw83_[int(9)] = False
                nw83_[int(10)] = self.f14
                nw83_[int(11)] = not((d_616_v19_).f12)
                nw83_[int(12)] = not((self).f10)
                nw83_[int(13)] = not(True)
                nw83_[int(14)] = self.f14
                nw83_[int(15)] = ((self).f15)[default__.safeIndex(59, ((self).f15).length(0))]
                nw83_[int(16)] = ((self).f15)[default__.safeIndex(59, ((self).f15).length(0))]
                nw83_[int(17)] = (d_616_v19_).f12
                nw83_[int(18)] = self.f14
                nw83_[int(19)] = ((self).f15)[default__.safeIndex(59, ((self).f15).length(0))]
                nw83_[int(20)] = (d_616_v19_).f12
                nw83_[int(21)] = ((self).f15)[default__.safeIndex(59, ((self).f15).length(0))]
                nw83_[int(22)] = (self).f10
                d_618_v21_ = nw83_
                d_619_v22_: _dafny.Map
                d_619_v22_ = _dafny.Map({not((d_616_v19_).f12): d_618_v21_})
                d_619_v22_ = ((d_619_v22_ if ((self).f15)[default__.safeIndex(59, ((self).f15).length(0))] else d_619_v22_)) | ((d_619_v22_))
            elif source14_.is_DC4:
                d_620___mcc_h1_ = source14_.cf6
                d_621___mcc_h2_ = source14_.cf7
                d_622___mcc_h3_ = source14_.cf8
                d_623_cf8_ = d_622___mcc_h3_
                d_624_cf7_ = d_621___mcc_h2_
                d_625_cf6_ = d_620___mcc_h1_
                d_626_v23_: D4
                d_626_v23_ = D4_DC11((self).fm1((self).f9, globalState), _dafny.Set({d_623_cf8_, d_590_v2_}), d_624_cf7_)
                d_623_cf8_ = (d_590_v2_) * ((d_626_v23_).cf16)
                rhs86_ = -571
                rhs87_ = d_624_cf7_
                lhs63_ = globalState
                d_623_cf8_ = rhs86_
                lhs63_.f8 = rhs87_
                d_627_v24_: _dafny.Map
                d_627_v24_ = _dafny.Map({d_590_v2_: default__.safeModuloInt(len(_dafny.Set({d_623_cf8_, d_623_cf8_})), d_623_cf8_)})
                d_627_v24_ = (d_627_v24_).set((self).fm42(d_590_v2_, (0) - (-380), d_590_v2_, globalState), default__.safeModuloInt(d_623_cf8_, d_623_cf8_))
                d_628_v25_: _dafny.Array
                nw84_ = _dafny.Array(int(0), 7)
                d_628_v25_ = nw84_
                index88_ = default__.safeIndex(405, (d_628_v25_).length(0))
                (d_628_v25_)[index88_] = d_623_cf8_
                index89_ = default__.safeIndex(405, (d_628_v25_).length(0))
                (d_628_v25_)[index89_] = (self).fm1((_dafny.Map({(self).fm0(d_589_v1_, not(d_625_cf6_), d_590_v2_, d_624_cf7_, globalState): self.f14})).set(True, d_625_cf6_), globalState)
            elif source14_.is_DC2:
                d_629___mcc_h4_ = source14_.cf4
                d_630_cf4_ = d_629___mcc_h4_
                (globalState).f8 = (d_590_v2_) >= (d_590_v2_)
                d_631_v26_: C0
                nw85_ = C0()
                nw85_.ctor__()
                d_631_v26_ = nw85_
                d_632_v27_: _dafny.Array
                def lambda18_(d_633_v2_):
                    def lambda19_(d_634_i7_):
                        return default__.safeDivisionInt(d_634_i7_, (0) - (d_633_v2_))

                    return lambda19_

                init11_ = lambda18_(d_590_v2_)
                nw86_ = _dafny.Array(None, 8)
                for i0_11_ in range(nw86_.length(0)):
                    nw86_[i0_11_] = init11_(i0_11_)
                d_632_v27_ = nw86_
                d_635_v28_: _dafny.Map
                d_635_v28_ = _dafny.Map({(self).fm1((self).f9, globalState): -190})
                d_636_v29_: _dafny.Seq
                d_636_v29_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rnvidwis"))
                d_637_v30_: _dafny.Map
                d_637_v30_ = _dafny.Map({d_636_v29_: d_590_v2_})
                index90_ = default__.safeIndex(935, (d_632_v27_).length(0))
                (d_632_v27_)[index90_] = (((d_635_v28_)[len(d_636_v29_)] if (len(d_636_v29_)) in (d_635_v28_) else d_590_v2_)) - (((d_637_v30_)[d_636_v29_] if (d_636_v29_) in (d_637_v30_) else d_590_v2_))
                index91_ = default__.safeIndex(935, (d_632_v27_).length(0))
                (d_632_v27_)[index91_] = d_590_v2_
                d_638_v31_: _dafny.Map
                d_638_v31_ = _dafny.Map({(d_590_v2_) != (d_590_v2_): ((self).f15)[default__.safeIndex(59, ((self).f15).length(0))]})
                d_638_v31_ = (d_638_v31_).set(((self).f15)[default__.safeIndex(59, ((self).f15).length(0))], ((self).f15)[default__.safeIndex(59, ((self).f15).length(0))])
            elif True:
                d_639___mcc_h5_ = source14_.cf9
                d_640_cf9_ = d_639___mcc_h5_
                d_588_v0_ = (d_588_v0_) - (d_588_v0_)
                d_641_v32_: D0
                d_641_v32_ = D0_DC0(self.f14)
                (self).f14 = ((d_641_v32_ if (self).f10 else d_641_v32_)).cf0
                d_642_v33_: _dafny.Array
                nw87_ = _dafny.Array(int(0), 22)
                d_642_v33_ = nw87_
                d_643_v34_: _dafny.Set
                d_643_v34_ = _dafny.Set({d_642_v33_})
                d_644_v35_: D8
                d_644_v35_ = D8_DC24((self).f10, d_643_v34_, d_590_v2_)
                d_645_v36_: _dafny.Seq
                d_645_v36_ = _dafny.SeqWithoutIsStrInference([d_644_v35_])
                d_646_v37_: _dafny.MultiSet
                d_646_v37_ = _dafny.MultiSet([(d_645_v36_) + ((d_645_v36_).set(default__.safeIndex(949, len(d_645_v36_)), d_644_v35_))])
                d_646_v37_ = (d_646_v37_).intersection(d_646_v37_)
                d_647_v38_: bool
                d_648_v39_: _dafny.Map
                d_649_v40_: D7
                out42_: bool
                out43_: _dafny.Map
                out44_: D7
                out42_, out43_, out44_ = (self).m7(511, d_590_v2_, globalState)
                d_647_v38_ = out42_
                d_648_v39_ = out43_
                d_649_v40_ = out44_
            d_650_v41_: _dafny.Array
            nw88_ = _dafny.Array(_dafny.Array(None, 0), 15)
            d_650_v41_ = nw88_
            d_651_v42_: _dafny.Array
            def lambda20_(d_652_i8_):
                return ((self).f15)[default__.safeIndex(59, ((self).f15).length(0))]

            init12_ = lambda20_
            nw89_ = _dafny.Array(None, 6)
            for i0_12_ in range(nw89_.length(0)):
                nw89_[i0_12_] = init12_(i0_12_)
            d_651_v42_ = nw89_
            index92_ = default__.safeIndex(651, (d_650_v41_).length(0))
            (d_650_v41_)[index92_] = d_651_v42_
            d_653_v43_: _dafny.Seq
            d_653_v43_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lueejueq"))
            d_654_v44_: _dafny.Array
            nw90_ = _dafny.Array(None, 8)
            nw90_[int(0)] = (default__.fm45(self.f14, (d_588_v0_).cardinality, ((self).f15)[default__.safeIndex(59, ((self).f15).length(0))], globalState)).set(default__.safeIndex((0) - (d_590_v2_), len(default__.fm45(self.f14, (d_588_v0_).cardinality, ((self).f15)[default__.safeIndex(59, ((self).f15).length(0))], globalState))), (d_593_v5_)[default__.safeIndex(353, (d_593_v5_).length(0))])
            nw90_[int(1)] = _dafny.SeqWithoutIsStrInference([(d_593_v5_)[default__.safeIndex(353, (d_593_v5_).length(0))] for d_655_i9_ in range(default__.abs(690))])
            nw90_[int(2)] = d_653_v43_
            nw90_[int(3)] = (d_653_v43_) + (d_653_v43_)
            nw90_[int(4)] = (_dafny.SeqWithoutIsStrInference([(d_593_v5_)[default__.safeIndex(353, (d_593_v5_).length(0))] for d_656_i10_ in range(default__.abs(584))])) + (d_653_v43_)
            nw90_[int(5)] = d_653_v43_
            nw90_[int(6)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cafhh"))
            nw90_[int(7)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wgxv"))) + (d_653_v43_)
            d_654_v44_ = nw90_
            index93_ = default__.safeIndex(86, (d_654_v44_).length(0))
            (d_654_v44_)[index93_] = d_653_v43_
            index94_ = default__.safeIndex(651, (d_650_v41_).length(0))
            index95_ = default__.safeIndex(86, (d_654_v44_).length(0))
            index96_ = default__.safeIndex(59, ((self).f15).length(0))
            rhs88_ = (d_651_v42_ if ((self).f15)[default__.safeIndex(59, ((self).f15).length(0))] else d_651_v42_)
            rhs89_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j"))) + (default__.fm45((self).fm0(d_589_v1_, self.f14, d_590_v2_, False, globalState), d_590_v2_, self.f14, globalState))
            rhs90_ = (self).f10
            rhs91_ = ((0) - (len(d_653_v43_))) + ((d_590_v2_) + (d_590_v2_))
            lhs64_ = d_650_v41_
            lhs65_ = default__.safeIndex(651, (d_650_v41_).length(0))
            lhs66_ = d_654_v44_
            lhs67_ = default__.safeIndex(86, (d_654_v44_).length(0))
            lhs68_ = (self).f15
            lhs69_ = default__.safeIndex(59, ((self).f15).length(0))
            lhs64_[lhs65_] = rhs88_
            lhs66_[lhs67_] = rhs89_
            lhs68_[lhs69_] = rhs90_
            d_590_v2_ = rhs91_
            if ((d_601_v10_)[d_590_v2_] if (d_590_v2_) in (d_601_v10_) else (d_590_v2_) == (d_590_v2_)):
                d_599_v9_ = d_599_v9_
                d_657_v45_: D1
                d_657_v45_ = D1_DC3(d_653_v43_)
                pat_let_tv9_ = d_653_v43_
                def iife53_(_pat_let8_0):
                    def iife54_(d_658_dt__update__tmp_h0_):
                        def iife55_(_pat_let9_0):
                            def iife56_(d_659_dt__update_hcf5_h0_):
                                return D1_DC3(d_659_dt__update_hcf5_h0_)
                            return iife56_(_pat_let9_0)
                        return iife55_(pat_let_tv9_)
                    return iife54_(_pat_let8_0)
                d_657_v45_ = iife53_(D1_DC3(d_653_v43_))
                d_660_v46_: _dafny.Array
                nw91_ = _dafny.Array(int(0), 11)
                d_660_v46_ = nw91_
                index97_ = default__.safeIndex(108, (d_660_v46_).length(0))
                (d_660_v46_)[index97_] = d_590_v2_
                d_661_v47_: _dafny.Array
                def lambda21_(d_662_v2_):
                    def lambda22_(d_663_i11_):
                        return D1_DC4(self.f14, (self).f10, d_662_v2_)

                    return lambda22_

                init13_ = lambda21_(d_590_v2_)
                nw92_ = _dafny.Array(None, 28)
                for i0_13_ in range(nw92_.length(0)):
                    nw92_[i0_13_] = init13_(i0_13_)
                d_661_v47_ = nw92_
                index98_ = default__.safeIndex(636, (d_661_v47_).length(0))
                (d_661_v47_)[index98_] = D1_DC4(not(((self).f15)[default__.safeIndex(59, ((self).f15).length(0))]), self.f14, d_590_v2_)
                d_664_v48_: _dafny.Seq
                d_664_v48_ = _dafny.SeqWithoutIsStrInference([d_651_v42_, (d_650_v41_)[default__.safeIndex(651, (d_650_v41_).length(0))]])
                d_665_v49_: D1
                d_665_v49_ = D1_DC4(((self).f15)[default__.safeIndex(59, ((self).f15).length(0))], (((self).f15)[default__.safeIndex(59, ((self).f15).length(0))]) == ((((self).f9)[((self).f15)[default__.safeIndex(59, ((self).f15).length(0))]] if (((self).f15)[default__.safeIndex(59, ((self).f15).length(0))]) in ((self).f9) else False)), d_590_v2_)
                index99_ = default__.safeIndex(651, (d_650_v41_).length(0))
                index100_ = default__.safeIndex(108, (d_660_v46_).length(0))
                index101_ = default__.safeIndex(636, (d_661_v47_).length(0))
                rhs92_ = (self).fm1((self).f9, globalState)
                rhs93_ = (d_664_v48_)[default__.safeIndex(len(_dafny.Set({(d_593_v5_)[default__.safeIndex(353, (d_593_v5_).length(0))], (d_593_v5_)[default__.safeIndex(353, (d_593_v5_).length(0))], _dafny.CodePoint('r')})), len(d_664_v48_))]
                rhs94_ = d_590_v2_
                rhs95_ = d_665_v49_
                rhs96_ = len(d_589_v1_)
                lhs70_ = d_650_v41_
                lhs71_ = default__.safeIndex(651, (d_650_v41_).length(0))
                lhs72_ = d_660_v46_
                lhs73_ = default__.safeIndex(108, (d_660_v46_).length(0))
                lhs74_ = d_661_v47_
                lhs75_ = default__.safeIndex(636, (d_661_v47_).length(0))
                d_590_v2_ = rhs92_
                lhs70_[lhs71_] = rhs93_
                lhs72_[lhs73_] = rhs94_
                lhs74_[lhs75_] = rhs95_
                d_590_v2_ = rhs96_
                d_590_v2_ = default__.safeDivisionInt(len(d_589_v1_), ((0) - ((self).fm42((d_660_v46_)[default__.safeIndex(108, (d_660_v46_).length(0))], len(_dafny.Map({(0) - (len((d_654_v44_)[default__.safeIndex(86, (d_654_v44_).length(0))])): (d_660_v46_)[default__.safeIndex(108, (d_660_v46_).length(0))]})), -430, globalState))) + (d_590_v2_))
                d_666_v50_: bool
                d_667_v51_: _dafny.Map
                d_668_v52_: D7
                out45_: bool
                out46_: _dafny.Map
                out47_: D7
                out45_, out46_, out47_ = (self).m7(default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ou"))), len((d_654_v44_)[default__.safeIndex(86, (d_654_v44_).length(0))])), d_590_v2_, globalState)
                d_666_v50_ = out45_
                d_667_v51_ = out46_
                d_668_v52_ = out47_
            elif True:
                d_669_v53_: _dafny.Set
                d_669_v53_ = _dafny.Set({(d_593_v5_)[default__.safeIndex(353, (d_593_v5_).length(0))]})
                d_670_v54_: _dafny.Map
                d_670_v54_ = _dafny.Map({d_590_v2_: d_590_v2_})
                d_671_v55_: D1
                d_671_v55_ = D1_DC4((self).fm0(_dafny.SeqWithoutIsStrInference([((self).f15)[default__.safeIndex(59, ((self).f15).length(0))]]), (self).f10, 445, (self).f10, globalState), self.f14, d_590_v2_)
                d_672_v56_: D5
                d_672_v56_ = D5_DC14(d_590_v2_, self.f14)
                d_673_v57_: _dafny.Array
                nw93_ = _dafny.Array(None, 12)
                nw93_[int(0)] = d_590_v2_
                nw93_[int(1)] = len(d_596_v7_)
                nw93_[int(2)] = (d_590_v2_) * (d_590_v2_)
                nw93_[int(3)] = d_590_v2_
                nw93_[int(4)] = default__.safeModuloInt(len(d_670_v54_), d_590_v2_)
                nw93_[int(5)] = (0) - ((d_590_v2_ if self.f14 else d_590_v2_))
                nw93_[int(6)] = d_590_v2_
                nw93_[int(7)] = d_590_v2_
                nw93_[int(8)] = (d_671_v55_).cf8
                nw93_[int(9)] = d_590_v2_
                nw93_[int(10)] = (d_672_v56_).cf25
                nw93_[int(11)] = len(_dafny.SeqWithoutIsStrInference([(d_593_v5_)[default__.safeIndex(353, (d_593_v5_).length(0))] for d_674_i12_ in range(default__.abs(364))]))
                d_673_v57_ = nw93_
                index102_ = default__.safeIndex(698, (d_673_v57_).length(0))
                (d_673_v57_)[index102_] = d_590_v2_
                d_675_v58_: _dafny.Map
                d_675_v58_ = _dafny.Map({(self).f9: ((self).f15)[default__.safeIndex(59, ((self).f15).length(0))]})
                d_676_v59_: _dafny.Map
                d_676_v59_ = _dafny.Map({(_dafny.MultiSet([((self).f15)[default__.safeIndex(59, ((self).f15).length(0))]])).cardinality: d_675_v58_})
                d_677_v61_: _dafny.MultiSet
                d_677_v61_ = _dafny.MultiSet([(self).f9])
                index103_ = default__.safeIndex(698, (d_673_v57_).length(0))
                def iife57_():
                    coll37_ = _dafny.Map()
                    compr_37_: _dafny.Map
                    for compr_37_ in (d_677_v61_).Elements:
                        d_678_v60_: _dafny.Map = compr_37_
                        if (d_678_v60_) in (d_677_v61_):
                            coll37_[d_678_v60_] = ((self).f15)[default__.safeIndex(59, ((self).f15).length(0))]
                    return _dafny.Map(coll37_)
                def iife58_():
                    coll38_ = _dafny.Set()
                    compr_38_: str
                    for compr_38_ in (d_653_v43_).Elements:
                        d_679_v62_: str = compr_38_
                        if (d_679_v62_) in (d_653_v43_):
                            coll38_ = coll38_.union(_dafny.Set([d_679_v62_]))
                    return _dafny.Set(coll38_)
                rhs97_ = default__.safeModuloInt((0) - (d_590_v2_), default__.safeDivisionInt(len(((d_676_v59_)[d_590_v2_] if (d_590_v2_) in (d_676_v59_) else iife57_()
                )), d_590_v2_))
                rhs98_ = ((d_669_v53_) - (d_669_v53_)) | ((_dafny.Set({(d_593_v5_)[default__.safeIndex(353, (d_593_v5_).length(0))]})) - (iife58_()
                ))
                rhs99_ = d_590_v2_
                rhs100_ = (0) - ((d_588_v0_).cardinality)
                lhs76_ = d_673_v57_
                lhs77_ = default__.safeIndex(698, (d_673_v57_).length(0))
                d_590_v2_ = rhs97_
                d_669_v53_ = rhs98_
                d_590_v2_ = rhs99_
                lhs76_[lhs77_] = rhs100_
                r0 = _dafny.SeqWithoutIsStrInference([(d_593_v5_)[default__.safeIndex(353, (d_593_v5_).length(0))] for d_680_i13_ in range(default__.abs(67))])
                d_681_v63_: D8
                d_681_v63_ = D8_DC24(not(((self).f15)[default__.safeIndex(59, ((self).f15).length(0))]), _dafny.Set({d_673_v57_}), (d_673_v57_)[default__.safeIndex(698, (d_673_v57_).length(0))])
                d_682_v64_: _dafny.Seq
                d_682_v64_ = _dafny.SeqWithoutIsStrInference([d_681_v63_, d_681_v63_, d_681_v63_])
                d_683_v65_: _dafny.Set
                d_683_v65_ = _dafny.Set({d_673_v57_})
                d_684_v66_: _dafny.Map
                d_684_v66_ = _dafny.Map({(d_673_v57_)[default__.safeIndex(698, (d_673_v57_).length(0))]: d_682_v64_})
                d_685_v67_: _dafny.Array
                nw94_ = _dafny.Array(None, 21)
                nw94_[int(0)] = d_682_v64_
                nw94_[int(1)] = (d_682_v64_) + (_dafny.SeqWithoutIsStrInference([d_681_v63_]))
                nw94_[int(2)] = d_682_v64_
                nw94_[int(3)] = d_682_v64_
                nw94_[int(4)] = d_682_v64_
                nw94_[int(5)] = d_682_v64_
                nw94_[int(6)] = _dafny.SeqWithoutIsStrInference([d_681_v63_, d_681_v63_, d_681_v63_, D8_DC24(self.f14, d_683_v65_, (d_673_v57_)[default__.safeIndex(698, (d_673_v57_).length(0))]), d_681_v63_])
                nw94_[int(7)] = _dafny.SeqWithoutIsStrInference([d_681_v63_, d_681_v63_])
                nw94_[int(8)] = (d_682_v64_ if True else d_682_v64_)
                nw94_[int(9)] = d_682_v64_
                nw94_[int(10)] = (_dafny.SeqWithoutIsStrInference([d_681_v63_, d_681_v63_, d_681_v63_]) if ((self).f15)[default__.safeIndex(59, ((self).f15).length(0))] else d_682_v64_)
                nw94_[int(11)] = (d_682_v64_).set(default__.safeIndex((d_673_v57_)[default__.safeIndex(698, (d_673_v57_).length(0))], len(d_682_v64_)), d_681_v63_)
                nw94_[int(12)] = d_682_v64_
                nw94_[int(13)] = d_682_v64_
                nw94_[int(14)] = d_682_v64_
                nw94_[int(15)] = _dafny.SeqWithoutIsStrInference([d_681_v63_, d_681_v63_, d_681_v63_])
                nw94_[int(16)] = (d_682_v64_) + (_dafny.SeqWithoutIsStrInference([d_681_v63_, d_681_v63_]))
                nw94_[int(17)] = d_682_v64_
                nw94_[int(18)] = d_682_v64_
                nw94_[int(19)] = ((d_684_v66_)[d_590_v2_] if (d_590_v2_) in (d_684_v66_) else d_682_v64_)
                nw94_[int(20)] = d_682_v64_
                d_685_v67_ = nw94_
                index104_ = default__.safeIndex(663, (d_685_v67_).length(0))
                (d_685_v67_)[index104_] = d_682_v64_
                d_686_v68_: _dafny.Map
                d_686_v68_ = _dafny.Map({d_588_v0_: _dafny.SeqWithoutIsStrInference([d_681_v63_])})
                pat_let_tv10_ = d_589_v1_
                pat_let_tv11_ = d_592_v4_
                pat_let_tv12_ = globalState
                pat_let_tv13_ = d_589_v1_
                pat_let_tv14_ = d_592_v4_
                pat_let_tv15_ = globalState
                index105_ = default__.safeIndex(663, (d_685_v67_).length(0))
                def iife59_(_pat_let10_0):
                    def iife60_(d_687_dt__update__tmp_h2_):
                        def iife61_(_pat_let11_0):
                            def iife62_(d_688_dt__update_hcf7_h1_):
                                def iife63_(_pat_let12_0):
                                    def iife64_(d_689_dt__update_hcf6_h1_):
                                        return D1_DC4(d_689_dt__update_hcf6_h1_, d_688_dt__update_hcf7_h1_, (d_687_dt__update__tmp_h2_).cf8)
                                    return iife64_(_pat_let12_0)
                                return iife63_((self).fm0(pat_let_tv10_, (self).f10, (pat_let_tv11_).cardinality, self.f14, pat_let_tv12_))
                            return iife62_(_pat_let11_0)
                        return iife61_((self).f10)
                    return iife60_(_pat_let10_0)
                def iife65_(_pat_let13_0):
                    def iife66_(d_690_dt__update__tmp_h1_):
                        def iife67_(_pat_let14_0):
                            def iife68_(d_691_dt__update_hcf7_h0_):
                                def iife69_(_pat_let15_0):
                                    def iife70_(d_692_dt__update_hcf6_h0_):
                                        return D1_DC4(d_692_dt__update_hcf6_h0_, d_691_dt__update_hcf7_h0_, (d_690_dt__update__tmp_h1_).cf8)
                                    return iife70_(_pat_let15_0)
                                return iife69_((self).fm0(pat_let_tv13_, (self).f10, (pat_let_tv14_).cardinality, self.f14, pat_let_tv15_))
                            return iife68_(_pat_let14_0)
                        return iife67_((self).f10)
                    return iife66_(_pat_let13_0)
                (d_685_v67_)[index105_] = ((d_686_v68_)[(d_588_v0_).set((iife59_(d_671_v55_)).cf6, default__.abs((self).fm1((self).f9, globalState)))] if ((d_588_v0_).set((iife65_(d_671_v55_)).cf6, default__.abs((self).fm1((self).f9, globalState)))) in (d_686_v68_) else d_682_v64_)
                d_693_v69_: _dafny.Array
                nw95_ = _dafny.Array(D8.default()(), 7)
                d_693_v69_ = nw95_
                index106_ = default__.safeIndex(802, (d_693_v69_).length(0))
                (d_693_v69_)[index106_] = d_681_v63_
                index107_ = default__.safeIndex(802, (d_693_v69_).length(0))
                (d_693_v69_)[index107_] = d_681_v63_
                index108_ = default__.safeIndex(59, ((self).f15).length(0))
                ((self).f15)[index108_] = not((((self).f15)[default__.safeIndex(59, ((self).f15).length(0))]) in (d_589_v1_))
            index109_ = default__.safeIndex(59, ((self).f15).length(0))
            rhs101_ = (self).f10
            rhs102_ = (176) - (len((self).f9))
            lhs78_ = (self).f15
            lhs79_ = default__.safeIndex(59, ((self).f15).length(0))
            lhs78_[lhs79_] = rhs101_
            d_590_v2_ = rhs102_
            d_590_v2_ = (0) - ((default__.safeDivisionInt(d_590_v2_, d_590_v2_)) - ((d_590_v2_ if ((self).f15)[default__.safeIndex(59, ((self).f15).length(0))] else d_590_v2_)))
        d_694_v70_: _dafny.Seq
        d_694_v70_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ysxcsgfq"))
        r0 = d_694_v70_
        r1 = False
        d_695_v71_: _dafny.Set
        d_695_v71_ = _dafny.Set({self.f14})
        d_696_v72_: _dafny.Seq
        d_696_v72_ = _dafny.SeqWithoutIsStrInference([(d_695_v71_) | (d_695_v71_)])
        r2 = d_696_v72_
        r3 = (d_694_v70_).set(default__.safeIndex((len(d_695_v71_)) * (d_590_v2_), len(d_694_v70_)), _dafny.CodePoint('v'))
        return r0, r1, r2, r3

    def m1(self, p0, globalState):
        r0: _dafny.Map = _dafny.Map({})
        r1: bool = False
        d_697_v0_: int
        d_697_v0_ = 125
        d_697_v0_ = 857
        if (self.f14) and (not (self.f14) or (self.f14)):
            index110_ = default__.safeIndex(575, ((self).f15).length(0))
            ((self).f15)[index110_] = (910) > (d_697_v0_)
            d_698_v1_: _dafny.Seq
            d_698_v1_ = _dafny.SeqWithoutIsStrInference([(True) == ((self).f10), False, False])
            d_699_v2_: _dafny.MultiSet
            d_699_v2_ = _dafny.MultiSet([d_697_v0_])
            index111_ = default__.safeIndex(575, ((self).f15).length(0))
            rhs103_ = True
            rhs104_ = d_698_v1_
            rhs105_ = (self).f10
            rhs106_ = (self).fm42(d_697_v0_, ((0) - (d_697_v0_)) + (d_697_v0_), (d_699_v2_).cardinality, globalState)
            lhs80_ = (self).f15
            lhs81_ = default__.safeIndex(575, ((self).f15).length(0))
            lhs82_ = globalState
            lhs80_[lhs81_] = rhs103_
            d_698_v1_ = rhs104_
            lhs82_.f8 = rhs105_
            d_697_v0_ = rhs106_
            d_697_v0_ = (d_697_v0_) - (d_697_v0_)
            d_700_v3_: _dafny.Seq
            d_700_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ll"))
            d_701_v4_: _dafny.Array
            def lambda23_(d_702_v0_):
                def lambda24_(d_703_i0_):
                    return (d_703_i0_) * (d_702_v0_)

                return lambda24_

            init14_ = lambda23_(d_697_v0_)
            nw96_ = _dafny.Array(None, 6)
            for i0_14_ in range(nw96_.length(0)):
                nw96_[i0_14_] = init14_(i0_14_)
            d_701_v4_ = nw96_
            index112_ = default__.safeIndex(727, (d_701_v4_).length(0))
            (d_701_v4_)[index112_] = (0) - (d_697_v0_)
            index113_ = default__.safeIndex(727, (d_701_v4_).length(0))
            rhs107_ = d_697_v0_
            rhs108_ = self.f14
            rhs109_ = (self).f10
            rhs110_ = d_700_v3_
            rhs111_ = ((d_697_v0_) * (840)) * (default__.safeModuloInt(d_697_v0_, d_697_v0_))
            lhs83_ = self
            lhs84_ = globalState
            lhs85_ = d_701_v4_
            lhs86_ = default__.safeIndex(727, (d_701_v4_).length(0))
            d_697_v0_ = rhs107_
            lhs83_.f14 = rhs108_
            lhs84_.f8 = rhs109_
            d_700_v3_ = rhs110_
            lhs85_[lhs86_] = rhs111_
            index114_ = default__.safeIndex(727, (d_701_v4_).length(0))
            (d_701_v4_)[index114_] = ((d_701_v4_)[default__.safeIndex(727, (d_701_v4_).length(0))]) - (d_697_v0_)
            d_704_v5_: C0
            nw97_ = C0()
            nw97_.ctor__()
            d_704_v5_ = nw97_
        elif True:
            d_705_v6_: _dafny.Array
            def lambda25_(d_706_v0_):
                def lambda26_(d_707_i1_):
                    return default__.safeDivisionInt(d_707_i1_, (D5_DC14(d_706_v0_, (self).f10)).cf25)

                return lambda26_

            init15_ = lambda25_(d_697_v0_)
            nw98_ = _dafny.Array(None, 26)
            for i0_15_ in range(nw98_.length(0)):
                nw98_[i0_15_] = init15_(i0_15_)
            d_705_v6_ = nw98_
            d_705_v6_ = d_705_v6_
            d_708_v7_: _dafny.Seq
            d_708_v7_ = _dafny.SeqWithoutIsStrInference([d_697_v0_, d_697_v0_, 858])
            d_709_v8_: _dafny.Map
            d_709_v8_ = _dafny.Map({self.f14: default__.safeDivisionInt(d_697_v0_, len(d_708_v7_))})
            d_709_v8_ = (d_709_v8_).set(True, (0) - ((self).fm1((self).f9, globalState)))
            d_710_v9_: _dafny.Array
            nw99_ = _dafny.Array(_dafny.Array(None, 0), 1)
            d_710_v9_ = nw99_
            index115_ = default__.safeIndex(193, (d_710_v9_).length(0))
            (d_710_v9_)[index115_] = d_705_v6_
            d_711_v10_: D3
            d_711_v10_ = D3_DC7(d_705_v6_)
            index116_ = default__.safeIndex(193, (d_710_v9_).length(0))
            (d_710_v9_)[index116_] = (d_711_v10_).cf11
            d_712_v11_: _dafny.Seq
            d_712_v11_ = _dafny.SeqWithoutIsStrInference([True])
            d_713_v12_: D6
            d_713_v12_ = D6_DC18(p0, len(d_712_v11_), d_697_v0_)
            d_714_v13_: D6
            d_714_v13_ = D6_DC20(D6_DC20(d_713_v12_))
            d_715_v14_: _dafny.Set
            def iife71_(_pat_let16_0):
                def iife72_(d_716_dt__update__tmp_h0_):
                    def iife73_(_pat_let17_0):
                        def iife74_(d_717_dt__update_hcf39_h0_):
                            return D6_DC20(d_717_dt__update_hcf39_h0_)
                        return iife74_(_pat_let17_0)
                    return iife73_(D6_DC16((self).f9))
                return iife72_(_pat_let16_0)
            d_715_v14_ = _dafny.Set({iife71_(d_714_v13_)})
            d_718_v15_: _dafny.Seq
            d_718_v15_ = _dafny.SeqWithoutIsStrInference([d_715_v14_])
            arr0_ = (d_710_v9_)[default__.safeIndex(193, (d_710_v9_).length(0))]
            index117_ = default__.safeIndex(282, ((d_710_v9_)[default__.safeIndex(193, (d_710_v9_).length(0))]).length(0))
            arr0_[index117_] = d_697_v0_
            arr1_ = (d_710_v9_)[default__.safeIndex(193, (d_710_v9_).length(0))]
            index118_ = default__.safeIndex(282, ((d_710_v9_)[default__.safeIndex(193, (d_710_v9_).length(0))]).length(0))
            index119_ = default__.safeIndex(193, (d_710_v9_).length(0))
            rhs112_ = ((0) - ((d_697_v0_) * (d_697_v0_))) * (d_697_v0_)
            rhs113_ = d_718_v15_
            rhs114_ = (195) * (d_697_v0_)
            rhs115_ = (d_711_v10_).cf11
            lhs87_ = (d_710_v9_)[default__.safeIndex(193, (d_710_v9_).length(0))]
            lhs88_ = default__.safeIndex(282, ((d_710_v9_)[default__.safeIndex(193, (d_710_v9_).length(0))]).length(0))
            lhs89_ = d_710_v9_
            lhs90_ = default__.safeIndex(193, (d_710_v9_).length(0))
            d_697_v0_ = rhs112_
            d_718_v15_ = rhs113_
            lhs87_[lhs88_] = rhs114_
            lhs89_[lhs90_] = rhs115_
            (globalState).f8 = (d_697_v0_) < ((((d_710_v9_)[default__.safeIndex(193, (d_710_v9_).length(0))])[default__.safeIndex(282, ((d_710_v9_)[default__.safeIndex(193, (d_710_v9_).length(0))]).length(0))] if (self).f10 else (default__.fm46((self).f10, (self).f10, (self).f10, globalState)).cardinality))
        d_719_v16_: _dafny.Array
        nw100_ = _dafny.Array(_dafny.Set({}), 18)
        d_719_v16_ = nw100_
        d_720_v17_: _dafny.Set
        d_720_v17_ = _dafny.Set({(self).f10})
        index120_ = default__.safeIndex(828, (d_719_v16_).length(0))
        (d_719_v16_)[index120_] = d_720_v17_
        index121_ = default__.safeIndex(828, (d_719_v16_).length(0))
        (d_719_v16_)[index121_] = d_720_v17_
        d_721_v18_: _dafny.Array
        nw101_ = _dafny.Array(_dafny.Seq({}), 12)
        d_721_v18_ = nw101_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_721_v18_).length(0)):
            d_722_i2_: int = guard_loop_0_
            if (True) and (((0) <= (d_722_i2_)) and ((d_722_i2_) < ((d_721_v18_).length(0)))):
                (d_721_v18_)[(d_722_i2_)] = (_dafny.SeqWithoutIsStrInference([(0) - ((0) - (d_697_v0_)), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "upbj"))), d_697_v0_])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kmewrgjgd"))), d_697_v0_, -56]))
        r1 = self.f14
        d_723_v19_: _dafny.Map
        d_723_v19_ = _dafny.Map({(self).f10: d_697_v0_})
        d_724_v20_: str
        d_724_v20_ = p0
        d_725_v21_: _dafny.Map
        d_725_v21_ = _dafny.Map({d_724_v20_: d_723_v19_})
        d_723_v19_ = ((d_725_v21_)[p0] if (p0) in (d_725_v21_) else _dafny.Map({self.f14: 39}))
        d_726_v22_: _dafny.Array
        nw102_ = _dafny.Array(None, 7)
        nw102_[int(0)] = d_697_v0_
        nw102_[int(1)] = d_697_v0_
        nw102_[int(2)] = d_697_v0_
        nw102_[int(3)] = -454
        nw102_[int(4)] = d_697_v0_
        nw102_[int(5)] = d_697_v0_
        nw102_[int(6)] = d_697_v0_
        d_726_v22_ = nw102_
        d_727_v23_: _dafny.Seq
        d_727_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "by"))
        d_728_v24_: _dafny.Map
        d_728_v24_ = _dafny.Map({d_726_v22_: d_727_v23_})
        r0 = (d_728_v24_).set(d_726_v22_, d_727_v23_)
        d_729_v25_: D1
        d_729_v25_ = D1_DC3(_dafny.SeqWithoutIsStrInference([p0 for d_730_i4_ in range(default__.abs(626))]))
        r1 = (_dafny.SeqWithoutIsStrInference([p0 for d_731_i3_ in range(default__.abs(322))])) <= ((d_729_v25_).cf5)
        return r0, r1

    def m7(self, p0, p1, globalState):
        r0: bool = False
        r1: _dafny.Map = _dafny.Map({})
        r2: D7 = D7.default()()
        r0 = (self).f10
        d_732_v0_: C0
        nw103_ = C0()
        nw103_.ctor__()
        d_732_v0_ = nw103_
        d_733_i0_: int
        d_733_i0_ = 0
        with _dafny.label("1"):
            while not(self.f14):
                with _dafny.c_label("1"):
                    if (d_733_i0_) >= (100):
                        raise _dafny.Break("1")
                    d_733_i0_ = (d_733_i0_) + (1)
                    d_734_v1_: int
                    d_734_v1_ = 891
                    d_734_v1_ = p0
                    d_735_v2_: str
                    d_735_v2_ = _dafny.CodePoint('o')
                    rhs116_ = d_735_v2_
                    rhs117_ = (d_734_v1_) * (p1)
                    rhs118_ = p0
                    d_735_v2_ = rhs116_
                    d_734_v1_ = rhs117_
                    d_734_v1_ = rhs118_
                    d_734_v1_ = (0) - (d_734_v1_)
                    d_734_v1_ = (0) - ((self).fm42(p0, p0, p0, globalState))
                    pass
            pass
        d_736_v3_: _dafny.Seq
        d_736_v3_ = _dafny.SeqWithoutIsStrInference([p0])
        d_737_v4_: _dafny.Map
        d_737_v4_ = _dafny.Map({p1: (d_736_v3_)[default__.safeIndex(341, len(d_736_v3_))]})
        d_738_v5_: _dafny.Map
        d_738_v5_ = _dafny.Map({self.f14: d_737_v4_})
        d_739_v6_: _dafny.Map
        d_739_v6_ = _dafny.Map({self.f14: p0})
        d_740_v7_: _dafny.MultiSet
        d_740_v7_ = _dafny.MultiSet([p1])
        d_741_v8_: _dafny.Map
        d_741_v8_ = _dafny.Map({(d_740_v7_).cardinality: (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nixblma"))))})
        d_738_v5_ = (d_738_v5_).set((d_739_v6_) in (_dafny.Set({_dafny.Map({(self).f10: p0})})), d_741_v8_)
        r0 = (((self).f10) or (self.f14)) == (self.f14)
        d_742_v9_: _dafny.Array
        nw104_ = _dafny.Array(None, 13)
        nw104_[int(0)] = self.f14
        nw104_[int(1)] = ((self).f10 if (self).f10 else self.f14)
        nw104_[int(2)] = (self).f10
        nw104_[int(3)] = self.f14
        nw104_[int(4)] = self.f14
        nw104_[int(5)] = self.f14
        nw104_[int(6)] = ((self).f10 if (self).f10 else (self).f10)
        nw104_[int(7)] = (self).f10
        nw104_[int(8)] = (self).f10
        nw104_[int(9)] = (708) > (-800)
        nw104_[int(10)] = not ((self).f10) or (self.f14)
        nw104_[int(11)] = not ((self).f10) or (self.f14)
        nw104_[int(12)] = self.f14
        d_742_v9_ = nw104_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_742_v9_).length(0)):
            d_743_i1_: int = guard_loop_1_
            if (True) and (((0) <= (d_743_i1_)) and ((d_743_i1_) < ((d_742_v9_).length(0)))):
                def iife75_():
                    coll39_ = _dafny.Set()
                    compr_39_: int
                    for compr_39_ in (_dafny.Set({p1})).Elements:
                        d_744_v10_: int = compr_39_
                        if (d_744_v10_) in (_dafny.Set({p1})):
                            coll39_ = coll39_.union(_dafny.Set([(d_744_v10_) - (len(_dafny.Map({not(not(not(True))): (0) - ((0) - ((0) - ((0) - (len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([len(_dafny.Map({False: True}))])).cardinality for d_745_i2_ in range(default__.abs(324))]))))))})))]))
                    return _dafny.Set(coll39_)
                (d_742_v9_)[(d_743_i1_)] = not ((len(iife75_()
                )) == (p0)) or ((self).f10)
        r0 = self.f14
        r1 = d_739_v6_
        d_746_v11_: D7
        d_746_v11_ = D7_DC22()
        r2 = d_746_v11_
        return r0, r1, r2

    @property
    def f15(self):
        return self._f15

class C3(T0):
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    def ctor__(self):
        pass
        pass

    def fm27(self, p0, p1, p2, globalState):
        return (not(False)) and ((-159) == (821))

    def fm28(self, p0, p1, p2, p3, globalState):
        def lambda27_(source15_):
            if source15_.is_DC3:
                d_747___mcc_h0_ = source15_.cf5
                d_748_cf5_ = d_747___mcc_h0_
                return (_dafny.Map({_dafny.CodePoint('f'): not(True)})) | (_dafny.Map({_dafny.CodePoint('j'): False}))
            elif source15_.is_DC4:
                d_749___mcc_h1_ = source15_.cf6
                d_750___mcc_h2_ = source15_.cf7
                d_751___mcc_h3_ = source15_.cf8
                d_752_cf8_ = d_751___mcc_h3_
                d_753_cf7_ = d_750___mcc_h2_
                d_754_cf6_ = d_749___mcc_h1_
                return _dafny.Map({_dafny.CodePoint('y'): not(d_753_cf7_)})
            elif source15_.is_DC2:
                d_755___mcc_h4_ = source15_.cf4
                d_756_cf4_ = d_755___mcc_h4_
                return (_dafny.Map({_dafny.CodePoint('q'): False})) | (_dafny.Map({_dafny.CodePoint('m'): True}))
            elif True:
                d_757___mcc_h5_ = source15_.cf9
                d_758_cf9_ = d_757___mcc_h5_
                return (_dafny.Map({_dafny.CodePoint('m'): not(True)})) | (_dafny.Map({_dafny.CodePoint('p'): False}))

        return len(lambda27_(D1_DC3(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_759_i0_ in range(default__.abs(725))]))))

    def m0(self, globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r1: bool = False
        r2: _dafny.Seq = _dafny.Seq({})
        r3: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        d_760_v0_: bool
        d_760_v0_ = True
        r1 = not (d_760_v0_) or (d_760_v0_)
        d_761_v1_: _dafny.Array
        nw105_ = _dafny.Array(int(0), 21)
        d_761_v1_ = nw105_
        d_762_v2_: _dafny.Map
        d_762_v2_ = _dafny.Map({d_761_v1_: d_760_v0_})
        d_762_v2_ = (d_762_v2_).set(d_761_v1_, (self).fm27(d_760_v0_, d_760_v0_, D3_DC8(d_760_v0_, d_760_v0_), globalState))
        d_763_v3_: int
        d_763_v3_ = 838
        d_764_v4_: str
        d_764_v4_ = _dafny.CodePoint('q')
        d_765_v5_: _dafny.Seq
        d_765_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gkun"))
        d_766_v6_: D1
        d_766_v6_ = D1_DC3(d_765_v5_)
        d_767_v7_: _dafny.Map
        d_767_v7_ = _dafny.Map({d_766_v6_: d_763_v3_})
        index122_ = default__.safeIndex(968, (d_761_v1_).length(0))
        (d_761_v1_)[index122_] = len((d_767_v7_).set(d_766_v6_, d_763_v3_))
        d_768_v8_: D3
        d_768_v8_ = D3_DC8(d_760_v0_, d_760_v0_)
        d_769_v9_: _dafny.MultiSet
        d_769_v9_ = _dafny.MultiSet([(self).fm27(d_760_v0_, d_760_v0_, d_768_v8_, globalState)])
        index123_ = default__.safeIndex(968, (d_761_v1_).length(0))
        rhs119_ = ((d_769_v9_) - (_dafny.MultiSet([True]))).cardinality
        rhs120_ = d_760_v0_
        rhs121_ = (d_765_v5_)[default__.safeIndex(d_763_v3_, len(d_765_v5_))]
        rhs122_ = (d_763_v3_) - (d_763_v3_)
        lhs91_ = globalState
        lhs92_ = d_761_v1_
        lhs93_ = default__.safeIndex(968, (d_761_v1_).length(0))
        d_763_v3_ = rhs119_
        lhs91_.f8 = rhs120_
        d_764_v4_ = rhs121_
        lhs92_[lhs93_] = rhs122_
        d_770_v10_: _dafny.Array
        def lambda28_(d_771_v0_):
            def lambda29_(d_772_i0_):
                return d_771_v0_

            return lambda29_

        init16_ = lambda28_(d_760_v0_)
        nw106_ = _dafny.Array(None, 10)
        for i0_16_ in range(nw106_.length(0)):
            nw106_[i0_16_] = init16_(i0_16_)
        d_770_v10_ = nw106_
        index124_ = default__.safeIndex(470, (d_770_v10_).length(0))
        (d_770_v10_)[index124_] = d_760_v0_
        d_773_v11_: _dafny.MultiSet
        d_773_v11_ = _dafny.MultiSet([(d_761_v1_)[default__.safeIndex(968, (d_761_v1_).length(0))], (d_761_v1_)[default__.safeIndex(968, (d_761_v1_).length(0))]])
        d_774_v12_: _dafny.Seq
        d_774_v12_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([len(d_765_v5_), d_763_v3_]), d_773_v11_])
        index125_ = default__.safeIndex(470, (d_770_v10_).length(0))
        (d_770_v10_)[index125_] = (d_763_v3_) not in ((d_774_v12_)[default__.safeIndex(d_763_v3_, len(d_774_v12_))])
        def iife76_():
            coll40_ = _dafny.Set()
            compr_40_: int
            for compr_40_ in _dafny.IntegerRange(-59, 873):
                d_795_v13_: int = compr_40_
                if ((-59) <= (d_795_v13_)) and ((d_795_v13_) < (873)):
                    coll40_ = coll40_.union(_dafny.Set([(d_795_v13_) * ((0) - ((d_761_v1_)[default__.safeIndex(968, (d_761_v1_).length(0))]))]))
            return _dafny.Set(coll40_)
        hi0_ = (d_763_v3_) + (d_763_v3_)
        for d_775_i1_ in range(((d_761_v1_)[default__.safeIndex(968, (d_761_v1_).length(0))]) - (len(iife76_()
        )), hi0_):
            d_776_v14_: D1
            d_776_v14_ = D1_DC4(d_760_v0_, d_760_v0_, d_775_i1_)
            if (d_776_v14_).cf7:
                index126_ = default__.safeIndex(968, (d_761_v1_).length(0))
                (d_761_v1_)[index126_] = (0) - ((0) - (d_775_i1_))
                d_777_v15_: _dafny.Array
                nw107_ = _dafny.Array(_dafny.Map({}), 20)
                d_777_v15_ = nw107_
                nw108_ = _dafny.Array(_dafny.Map({}), 9)
                d_777_v15_ = nw108_
                d_778_v16_: _dafny.Seq
                d_778_v16_ = _dafny.SeqWithoutIsStrInference([False])
                rhs123_ = d_778_v16_
                rhs124_ = d_760_v0_
                rhs125_ = (0) - (((d_761_v1_)[default__.safeIndex(968, (d_761_v1_).length(0))]) + ((d_761_v1_)[default__.safeIndex(968, (d_761_v1_).length(0))]))
                rhs126_ = (d_765_v5_) + (d_765_v5_)
                d_778_v16_ = rhs123_
                d_760_v0_ = rhs124_
                d_763_v3_ = rhs125_
                d_765_v5_ = rhs126_
                d_779_v17_: C0
                nw109_ = C0()
                nw109_.ctor__()
                d_779_v17_ = nw109_
                index127_ = default__.safeIndex(470, (d_770_v10_).length(0))
                (d_770_v10_)[index127_] = (self).fm27((d_775_i1_) <= ((d_761_v1_)[default__.safeIndex(968, (d_761_v1_).length(0))]), False, d_768_v8_, globalState)
            elif True:
                d_780_v18_: _dafny.Array
                def lambda30_(d_781_v14_):
                    def lambda31_(d_782_i2_):
                        return d_781_v14_

                    return lambda31_

                init17_ = lambda30_(d_776_v14_)
                nw110_ = _dafny.Array(None, 16)
                for i0_17_ in range(nw110_.length(0)):
                    nw110_[i0_17_] = init17_(i0_17_)
                d_780_v18_ = nw110_
                index128_ = default__.safeIndex(903, (d_780_v18_).length(0))
                (d_780_v18_)[index128_] = d_776_v14_
                index129_ = default__.safeIndex(903, (d_780_v18_).length(0))
                (d_780_v18_)[index129_] = d_776_v14_
                d_783_v19_: C0
                nw111_ = C0()
                nw111_.ctor__()
                d_783_v19_ = nw111_
                index130_ = default__.safeIndex(470, (d_770_v10_).length(0))
                (d_770_v10_)[index130_] = d_760_v0_
                d_784_v20_: _dafny.MultiSet
                d_784_v20_ = _dafny.MultiSet([d_764_v4_])
                d_785_v21_: _dafny.Map
                d_785_v21_ = _dafny.Map({d_760_v0_: d_784_v20_})
                d_786_v22_: _dafny.Set
                d_786_v22_ = _dafny.Set({d_775_i1_})
                d_787_v23_: _dafny.Set
                d_787_v23_ = _dafny.Set({(d_770_v10_)[default__.safeIndex(470, (d_770_v10_).length(0))], (d_770_v10_)[default__.safeIndex(470, (d_770_v10_).length(0))]})
                d_788_v24_: _dafny.Seq
                d_788_v24_ = _dafny.SeqWithoutIsStrInference([True, False])
                d_789_v25_: _dafny.Seq
                d_789_v25_ = _dafny.SeqWithoutIsStrInference([(d_788_v24_)[default__.safeIndex(d_775_i1_, len(d_788_v24_))]])
                d_790_v26_: _dafny.Array
                nw112_ = _dafny.Array(None, 19)
                nw112_[int(0)] = (d_785_v21_) | (d_785_v21_)
                nw112_[int(1)] = d_785_v21_
                nw112_[int(2)] = d_785_v21_
                nw112_[int(3)] = (d_785_v21_) | (d_785_v21_)
                nw112_[int(4)] = d_785_v21_
                nw112_[int(5)] = d_785_v21_
                nw112_[int(6)] = d_785_v21_
                nw112_[int(7)] = d_785_v21_
                nw112_[int(8)] = _dafny.Map({d_760_v0_: d_784_v20_})
                nw112_[int(9)] = default__.fm29(d_760_v0_, True, default__.fm30(globalState), d_786_v22_, globalState)
                nw112_[int(10)] = default__.fm29(d_760_v0_, d_760_v0_, d_787_v23_, d_786_v22_, globalState)
                nw112_[int(11)] = (d_785_v21_).set(d_760_v0_, d_784_v20_)
                nw112_[int(12)] = (_dafny.Map({not((d_789_v25_)[default__.safeIndex(d_775_i1_, len(d_789_v25_))]): d_784_v20_}) if (d_770_v10_)[default__.safeIndex(470, (d_770_v10_).length(0))] else _dafny.Map({d_760_v0_: d_784_v20_}))
                nw112_[int(13)] = (default__.fm29((d_770_v10_)[default__.safeIndex(470, (d_770_v10_).length(0))], d_760_v0_, d_787_v23_, d_786_v22_, globalState)) | (d_785_v21_)
                nw112_[int(14)] = (d_785_v21_) | (_dafny.Map({(d_770_v10_)[default__.safeIndex(470, (d_770_v10_).length(0))]: d_784_v20_}))
                nw112_[int(15)] = d_785_v21_
                nw112_[int(16)] = (D4_DC10(_dafny.Map({d_760_v0_: d_784_v20_}))).cf15
                nw112_[int(17)] = d_785_v21_
                nw112_[int(18)] = d_785_v21_
                d_790_v26_ = nw112_
                index131_ = default__.safeIndex(489, (d_790_v26_).length(0))
                (d_790_v26_)[index131_] = _dafny.Map({True: _dafny.MultiSet([d_764_v4_])})
                index132_ = default__.safeIndex(489, (d_790_v26_).length(0))
                (d_790_v26_)[index132_] = d_785_v21_
                d_764_v4_ = d_764_v4_
            d_791_v27_: C0
            nw113_ = C0()
            nw113_.ctor__()
            d_791_v27_ = nw113_
            d_792_v28_: C0
            nw114_ = C0()
            nw114_.ctor__()
            d_792_v28_ = nw114_
            d_793_v29_: _dafny.Seq
            d_793_v29_ = _dafny.SeqWithoutIsStrInference([(d_770_v10_)[default__.safeIndex(470, (d_770_v10_).length(0))]])
            d_794_v30_: _dafny.Map
            d_794_v30_ = _dafny.Map({-256: d_793_v29_})
            d_794_v30_ = (d_794_v30_).set((d_773_v11_).cardinality, d_793_v29_)
        hi1_ = (d_761_v1_)[default__.safeIndex(968, (d_761_v1_).length(0))]
        for d_796_i3_ in range((0) - (len(_dafny.SeqWithoutIsStrInference([d_764_v4_ for d_803_i4_ in range(default__.abs(838))]))), hi1_):
            d_769_v9_ = d_769_v9_
            index133_ = default__.safeIndex(470, (d_770_v10_).length(0))
            (d_770_v10_)[index133_] = (d_770_v10_)[default__.safeIndex(470, (d_770_v10_).length(0))]
            (globalState).f8 = ((0) - (d_796_i3_)) >= ((d_761_v1_)[default__.safeIndex(968, (d_761_v1_).length(0))])
            d_797_v31_: _dafny.Set
            d_797_v31_ = _dafny.Set({d_763_v3_, d_796_i3_})
            d_798_v32_: D4
            d_798_v32_ = D4_DC12((len(_dafny.SeqWithoutIsStrInference([d_764_v4_ for d_799_i5_ in range(default__.abs(-450))]))) not in (d_797_v31_), default__.safeDivisionInt(303, d_796_i3_), d_764_v4_, d_763_v3_, 754)
            d_800_v33_: _dafny.Map
            d_800_v33_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_764_v4_ for d_801_i6_ in range(default__.abs(338))]): False})
            d_802_v34_: D1
            d_802_v34_ = D1_DC4(d_760_v0_, d_760_v0_, (0) - (len(d_800_v33_)))
            index134_ = default__.safeIndex(968, (d_761_v1_).length(0))
            index135_ = default__.safeIndex(968, (d_761_v1_).length(0))
            index136_ = default__.safeIndex(968, (d_761_v1_).length(0))
            rhs127_ = d_798_v32_
            rhs128_ = (d_796_i3_) * ((self).fm28((d_770_v10_)[default__.safeIndex(470, (d_770_v10_).length(0))], d_802_v34_, d_763_v3_, -108, globalState))
            rhs129_ = default__.safeDivisionInt(261, ((self).fm28(d_760_v0_, d_802_v34_, (d_761_v1_)[default__.safeIndex(968, (d_761_v1_).length(0))], d_796_i3_, globalState)) - (d_796_i3_))
            rhs130_ = d_763_v3_
            lhs94_ = d_761_v1_
            lhs95_ = default__.safeIndex(968, (d_761_v1_).length(0))
            lhs96_ = d_761_v1_
            lhs97_ = default__.safeIndex(968, (d_761_v1_).length(0))
            lhs98_ = d_761_v1_
            lhs99_ = default__.safeIndex(968, (d_761_v1_).length(0))
            d_798_v32_ = rhs127_
            lhs94_[lhs95_] = rhs128_
            lhs96_[lhs97_] = rhs129_
            lhs98_[lhs99_] = rhs130_
        r0 = ((d_765_v5_) + (d_765_v5_)).set(default__.safeIndex((d_761_v1_)[default__.safeIndex(968, (d_761_v1_).length(0))], len((d_765_v5_) + (d_765_v5_))), d_764_v4_)
        r1 = d_760_v0_
        d_804_v35_: _dafny.Set
        d_804_v35_ = _dafny.Set({False, (d_770_v10_)[default__.safeIndex(470, (d_770_v10_).length(0))]})
        r2 = _dafny.SeqWithoutIsStrInference([d_804_v35_])
        r3 = (d_765_v5_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "letpy")))
        return r0, r1, r2, r3

    def m1(self, p0, globalState):
        r0: _dafny.Map = _dafny.Map({})
        r1: bool = False
        d_805_v0_: C0
        nw115_ = C0()
        nw115_.ctor__()
        d_805_v0_ = nw115_
        d_806_v1_: bool
        d_806_v1_ = True
        d_807_v2_: int
        d_807_v2_ = 248
        d_808_v3_: _dafny.Map
        d_808_v3_ = _dafny.Map({d_807_v2_: d_806_v1_})
        d_809_v4_: D4
        d_809_v4_ = D4_DC12(d_806_v1_, d_807_v2_, default__.fm31(d_806_v1_, d_808_v3_, globalState), d_807_v2_, d_807_v2_)
        pat_let_tv16_ = d_806_v1_
        pat_let_tv17_ = d_806_v1_
        def lambda32_(source16_):
            if source16_.is_DC11:
                d_810___mcc_h0_ = source16_.cf16
                d_811___mcc_h1_ = source16_.cf17
                d_812___mcc_h2_ = source16_.cf18
                d_813_cf18_ = d_812___mcc_h2_
                d_814_cf17_ = d_811___mcc_h1_
                d_815_cf16_ = d_810___mcc_h0_
                return pat_let_tv16_
            elif source16_.is_DC12:
                d_816___mcc_h3_ = source16_.cf19
                d_817___mcc_h4_ = source16_.cf20
                d_818___mcc_h5_ = source16_.cf21
                d_819___mcc_h6_ = source16_.cf22
                d_820___mcc_h7_ = source16_.cf23
                d_821_cf23_ = d_820___mcc_h7_
                d_822_cf22_ = d_819___mcc_h6_
                d_823_cf21_ = d_818___mcc_h5_
                d_824_cf20_ = d_817___mcc_h4_
                d_825_cf19_ = d_816___mcc_h3_
                return (d_822_cf22_) < (d_824_cf20_)
            elif True:
                d_826___mcc_h8_ = source16_.cf15
                d_827_cf15_ = d_826___mcc_h8_
                return pat_let_tv17_

        if lambda32_(d_809_v4_):
            d_828_v6_: _dafny.Array
            def lambda33_(d_829_v1_, d_830_v2_):
                def lambda34_(d_831_i0_):
                    def iife77_():
                        coll41_ = _dafny.Set()
                        compr_41_: int
                        for compr_41_ in _dafny.IntegerRange(-371, 330):
                            d_832_v5_: int = compr_41_
                            if ((-371) <= (d_832_v5_)) and ((d_832_v5_) < (330)):
                                coll41_ = coll41_.union(_dafny.Set([(d_832_v5_) + (d_830_v2_)]))
                        return _dafny.Set(coll41_)
                    return (D4_DC11(-957, iife77_()
, d_829_v1_)).cf18

                return lambda34_

            init18_ = lambda33_(d_806_v1_, d_807_v2_)
            nw116_ = _dafny.Array(None, 4)
            for i0_18_ in range(nw116_.length(0)):
                nw116_[i0_18_] = init18_(i0_18_)
            d_828_v6_ = nw116_
            index137_ = default__.safeIndex(306, (d_828_v6_).length(0))
            (d_828_v6_)[index137_] = d_806_v1_
            index138_ = default__.safeIndex(306, (d_828_v6_).length(0))
            (d_828_v6_)[index138_] = d_806_v1_
            index139_ = default__.safeIndex(306, (d_828_v6_).length(0))
            (d_828_v6_)[index139_] = d_806_v1_
            d_833_v7_: C0
            nw117_ = C0()
            nw117_.ctor__()
            d_833_v7_ = nw117_
            d_834_v8_: _dafny.Seq
            d_834_v8_ = _dafny.SeqWithoutIsStrInference([d_806_v1_])
            if ((d_834_v8_)[default__.safeIndex(d_807_v2_, len(d_834_v8_))]) or ((d_806_v1_ if (d_828_v6_)[default__.safeIndex(306, (d_828_v6_).length(0))] else (d_828_v6_)[default__.safeIndex(306, (d_828_v6_).length(0))])):
                index140_ = default__.safeIndex(306, (d_828_v6_).length(0))
                (d_828_v6_)[index140_] = not(d_806_v1_)
                d_835_v9_: _dafny.Seq
                d_835_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ewyay"))
                d_835_v9_ = (d_835_v9_) + ((d_835_v9_) + (_dafny.SeqWithoutIsStrInference([p0 for d_836_i1_ in range(default__.abs(-365))])))
                d_837_v10_: _dafny.Map
                d_837_v10_ = _dafny.Map({d_828_v6_: d_806_v1_})
                d_838_v11_: D5
                d_838_v11_ = D5_DC13(d_828_v6_)
                d_837_v10_ = (d_837_v10_).set((d_838_v11_).cf24, (d_828_v6_)[default__.safeIndex(306, (d_828_v6_).length(0))])
                d_839_v12_: _dafny.MultiSet
                d_839_v12_ = _dafny.MultiSet([False, d_806_v1_, True, (d_828_v6_)[default__.safeIndex(306, (d_828_v6_).length(0))], ((0) - (d_807_v2_)) == (d_807_v2_)])
                d_839_v12_ = (_dafny.MultiSet((d_834_v8_) + (d_834_v8_))).intersection(d_839_v12_)
                d_840_v13_: _dafny.Map
                d_840_v13_ = _dafny.Map({d_807_v2_: d_805_v0_})
                d_841_v14_: _dafny.Map
                d_841_v14_ = _dafny.Map({d_807_v2_: len(d_840_v13_)})
                d_842_v15_: D5
                d_842_v15_ = D5_DC14(d_807_v2_, (d_828_v6_)[default__.safeIndex(306, (d_828_v6_).length(0))])
                d_843_v16_: _dafny.Map
                d_843_v16_ = _dafny.Map({_dafny.Map({(d_828_v6_)[default__.safeIndex(306, (d_828_v6_).length(0))]: _dafny.Set({d_833_v7_, d_805_v0_})}): _dafny.Map({(d_842_v15_).cf25: d_807_v2_})})
                d_844_v17_: _dafny.Set
                d_844_v17_ = _dafny.Set({d_805_v0_})
                d_845_v18_: _dafny.Map
                d_845_v18_ = _dafny.Map({(d_828_v6_)[default__.safeIndex(306, (d_828_v6_).length(0))]: d_844_v17_})
                d_846_v19_: D1
                d_846_v19_ = D1_DC4(d_806_v1_, (d_828_v6_)[default__.safeIndex(306, (d_828_v6_).length(0))], d_807_v2_)
                d_847_v20_: _dafny.MultiSet
                d_847_v20_ = _dafny.MultiSet([d_807_v2_, d_807_v2_])
                d_848_v21_: _dafny.Array
                nw118_ = _dafny.Array(None, 13)
                nw118_[int(0)] = d_841_v14_
                nw118_[int(1)] = ((d_843_v16_)[d_845_v18_] if (d_845_v18_) in (d_843_v16_) else d_841_v14_)
                nw118_[int(2)] = d_841_v14_
                nw118_[int(3)] = ((d_841_v14_).set(d_807_v2_, d_807_v2_)).set(d_807_v2_, d_807_v2_)
                nw118_[int(4)] = d_841_v14_
                nw118_[int(5)] = d_841_v14_
                nw118_[int(6)] = (d_841_v14_).set((self).fm28((d_828_v6_)[default__.safeIndex(306, (d_828_v6_).length(0))], d_846_v19_, (d_847_v20_).cardinality, d_807_v2_, globalState), d_807_v2_)
                nw118_[int(7)] = _dafny.Map({d_807_v2_: d_807_v2_})
                nw118_[int(8)] = (d_841_v14_) | (_dafny.Map({136: 457}))
                nw118_[int(9)] = d_841_v14_
                nw118_[int(10)] = d_841_v14_
                nw118_[int(11)] = d_841_v14_
                nw118_[int(12)] = d_841_v14_
                d_848_v21_ = nw118_
                d_848_v21_ = d_848_v21_
            elif True:
                d_849_v22_: _dafny.Seq
                d_849_v22_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nhxyjlt"))
                index141_ = default__.safeIndex(306, (d_828_v6_).length(0))
                (d_828_v6_)[index141_] = (default__.safeModuloInt(len(d_849_v22_), d_807_v2_)) < (len(_dafny.Set({d_807_v2_})))
                d_850_v23_: _dafny.Seq
                d_850_v23_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hpdclvpfa"))])
                d_849_v22_ = (d_849_v22_) + ((d_850_v23_)[default__.safeIndex((0) - (d_807_v2_), len(d_850_v23_))])
                d_851_v24_: _dafny.Set
                d_851_v24_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([p0 for d_852_i2_ in range(default__.abs(739))])})
                d_853_v25_: _dafny.Map
                d_853_v25_ = _dafny.Map({d_834_v8_: False})
                d_854_v26_: _dafny.Map
                d_854_v26_ = _dafny.Map({d_807_v2_: p0})
                d_855_v27_: _dafny.MultiSet
                d_855_v27_ = _dafny.MultiSet([d_807_v2_, d_807_v2_, d_807_v2_, len(d_854_v26_)])
                d_856_v28_: _dafny.MultiSet
                d_856_v28_ = _dafny.MultiSet([False, False])
                d_857_v29_: _dafny.Map
                d_857_v29_ = _dafny.Map({(d_828_v6_)[default__.safeIndex(306, (d_828_v6_).length(0))]: not(d_806_v1_)})
                d_858_v30_: _dafny.Seq
                d_858_v30_ = _dafny.SeqWithoutIsStrInference([d_807_v2_])
                d_859_v31_: _dafny.Array
                nw119_ = _dafny.Array(None, 24)
                nw119_[int(0)] = default__.safeDivisionInt((0) - ((0) - (d_807_v2_)), d_807_v2_)
                nw119_[int(1)] = default__.safeDivisionInt(d_807_v2_, d_807_v2_)
                nw119_[int(2)] = 705
                nw119_[int(3)] = d_807_v2_
                nw119_[int(4)] = 559
                nw119_[int(5)] = d_807_v2_
                nw119_[int(6)] = d_807_v2_
                nw119_[int(7)] = (d_807_v2_) - (d_807_v2_)
                nw119_[int(8)] = len(d_851_v24_)
                nw119_[int(9)] = d_807_v2_
                nw119_[int(10)] = len((d_853_v25_) | (d_853_v25_))
                nw119_[int(11)] = ((d_855_v27_)[(d_856_v28_).cardinality] if ((d_856_v28_).cardinality) in (d_855_v27_) else d_807_v2_)
                nw119_[int(12)] = d_807_v2_
                nw119_[int(13)] = ((0) - (d_807_v2_)) * (d_807_v2_)
                nw119_[int(14)] = 779
                nw119_[int(15)] = len(d_834_v8_)
                nw119_[int(16)] = (0) - ((len(d_857_v29_)) + ((d_858_v30_)[default__.safeIndex(d_807_v2_, len(d_858_v30_))]))
                nw119_[int(17)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qryuduwux")))
                nw119_[int(18)] = d_807_v2_
                nw119_[int(19)] = (0) - (len(d_858_v30_))
                nw119_[int(20)] = d_807_v2_
                nw119_[int(21)] = d_807_v2_
                nw119_[int(22)] = d_807_v2_
                nw119_[int(23)] = default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "loerye"))), d_807_v2_)
                d_859_v31_ = nw119_
                d_860_v32_: _dafny.Seq
                d_860_v32_ = _dafny.SeqWithoutIsStrInference([d_859_v31_, d_859_v31_, d_859_v31_])
                d_859_v31_ = (d_860_v32_)[default__.safeIndex((d_807_v2_) - (d_807_v2_), len(d_860_v32_))]
                index142_ = default__.safeIndex(212, (d_859_v31_).length(0))
                (d_859_v31_)[index142_] = d_807_v2_
                d_861_v33_: _dafny.Set
                d_861_v33_ = _dafny.Set({(d_828_v6_)[default__.safeIndex(306, (d_828_v6_).length(0))], False})
                d_862_v34_: _dafny.Map
                d_862_v34_ = _dafny.Map({d_807_v2_: d_861_v33_})
                index143_ = default__.safeIndex(212, (d_859_v31_).length(0))
                index144_ = default__.safeIndex(306, (d_828_v6_).length(0))
                rhs131_ = d_849_v22_
                rhs132_ = (d_807_v2_) * ((d_807_v2_ if d_806_v1_ else 865))
                rhs133_ = not((d_828_v6_)[default__.safeIndex(306, (d_828_v6_).length(0))])
                rhs134_ = d_862_v34_
                rhs135_ = d_807_v2_
                lhs100_ = d_859_v31_
                lhs101_ = default__.safeIndex(212, (d_859_v31_).length(0))
                lhs102_ = d_828_v6_
                lhs103_ = default__.safeIndex(306, (d_828_v6_).length(0))
                d_849_v22_ = rhs131_
                lhs100_[lhs101_] = rhs132_
                lhs102_[lhs103_] = rhs133_
                d_862_v34_ = rhs134_
                d_807_v2_ = rhs135_
                d_863_v35_: _dafny.Map
                d_863_v35_ = _dafny.Map({p0: d_807_v2_})
                index145_ = default__.safeIndex(212, (d_859_v31_).length(0))
                (d_859_v31_)[index145_] = ((d_863_v35_)[default__.fm31(not(True), d_808_v3_, globalState)] if (default__.fm31(not(True), d_808_v3_, globalState)) in (d_863_v35_) else 111)
            d_864_v36_: C0
            nw120_ = C0()
            nw120_.ctor__()
            d_864_v36_ = nw120_
        elif True:
            (globalState).f8 = d_806_v1_
            (globalState).f8 = d_806_v1_
            d_865_v37_: T0
            nw121_ = C1()
            nw121_.ctor__(d_806_v1_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_866_i3_ in range(default__.abs(173))]))
            d_865_v37_ = nw121_
            d_867_v38_: _dafny.Array
            def lambda35_(d_868_v1_):
                def lambda36_(d_869_i4_):
                    return d_868_v1_

                return lambda36_

            init19_ = lambda35_(d_806_v1_)
            nw122_ = _dafny.Array(None, 13)
            for i0_19_ in range(nw122_.length(0)):
                nw122_[i0_19_] = init19_(i0_19_)
            d_867_v38_ = nw122_
            d_870_v39_: _dafny.Map
            d_870_v39_ = _dafny.Map({True: d_806_v1_})
            d_871_v40_: T1
            nw123_ = C2()
            nw123_.ctor__(d_806_v1_, d_867_v38_, d_870_v39_, d_806_v1_)
            d_871_v40_ = nw123_
            d_872_v41_: _dafny.Map
            d_872_v41_ = _dafny.Map({d_865_v37_: d_871_v40_})
            d_872_v41_ = (d_872_v41_).set(d_865_v37_, d_871_v40_)
            d_807_v2_ = (d_807_v2_) + (752)
            d_873_v42_: _dafny.Map
            d_873_v42_ = _dafny.Map({d_867_v38_: d_807_v2_})
            d_873_v42_ = (d_873_v42_).set(d_867_v38_, (d_807_v2_ if d_806_v1_ else d_807_v2_))
        rhs136_ = not(d_806_v1_)
        rhs137_ = (0) - (d_807_v2_)
        r1 = rhs136_
        d_807_v2_ = rhs137_
        d_874_v43_: _dafny.Set
        d_874_v43_ = _dafny.Set({d_807_v2_, d_807_v2_})
        d_875_v44_: _dafny.Seq
        d_875_v44_ = _dafny.SeqWithoutIsStrInference([(725) >= (len(d_874_v43_)), (d_807_v2_) < (d_807_v2_)])
        r1 = (d_875_v44_)[default__.safeIndex(d_807_v2_, len(d_875_v44_))]
        d_876_v45_: _dafny.Array
        def lambda37_(d_877_v2_, d_878_v44_):
            def lambda38_(d_879_i6_):
                return default__.safeModuloInt(d_879_i6_, (0) - (len(_dafny.Map({d_877_v2_: len(d_878_v44_)}))))

            return lambda38_

        init20_ = lambda37_(d_807_v2_, d_875_v44_)
        nw124_ = _dafny.Array(None, 12)
        for i0_20_ in range(nw124_.length(0)):
            nw124_[i0_20_] = init20_(i0_20_)
        d_876_v45_ = nw124_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_876_v45_).length(0)):
            d_880_i5_: int = guard_loop_2_
            if (True) and (((0) <= (d_880_i5_)) and ((d_880_i5_) < ((d_876_v45_).length(0)))):
                (d_876_v45_)[(d_880_i5_)] = default__.safeModuloInt(d_880_i5_, d_807_v2_)
        (globalState).f8 = d_806_v1_
        d_881_v46_: _dafny.Seq
        d_881_v46_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tgohvumtv"))
        r0 = _dafny.Map({(d_876_v45_ if d_806_v1_ else d_876_v45_): (d_881_v46_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kaeltdutj")))})
        r1 = (d_806_v1_ if d_806_v1_ else d_806_v1_)
        return r0, r1


class C4(T1, T0):
    def  __init__(self):
        self._f9: _dafny.Map = _dafny.Map({})
        self._f10: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    @property
    def f9(self):
        return self._f9
    @property
    def f10(self):
        return self._f10
    def ctor__(self, f9, f10):
        (self)._f9 = f9
        (self)._f10 = f10

    def fm0(self, p0, p1, p2, p3, globalState):
        return not ((self).f10) or ((self).f10)

    def fm1(self, p0, globalState):
        return default__.safeModuloInt(-969, (-383) * (184))

    def m0(self, globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r1: bool = False
        r2: _dafny.Seq = _dafny.Seq({})
        r3: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        d_882_v0_: _dafny.Seq
        d_882_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qlknnmx"))
        d_883_v1_: D1
        d_883_v1_ = D1_DC3(d_882_v0_)
        d_884_v2_: _dafny.Seq
        d_884_v2_ = _dafny.SeqWithoutIsStrInference([(self).f10])
        d_885_v3_: int
        d_885_v3_ = 30
        pat_let_tv18_ = d_884_v2_
        pat_let_tv19_ = globalState
        pat_let_tv20_ = d_885_v3_
        pat_let_tv21_ = globalState
        d_886_v4_: _dafny.Array
        nw125_ = _dafny.Array(None, 19)
        nw125_[int(0)] = d_883_v1_
        nw125_[int(1)] = d_883_v1_
        nw125_[int(2)] = d_883_v1_
        nw125_[int(3)] = d_883_v1_
        nw125_[int(4)] = d_883_v1_
        nw125_[int(5)] = d_883_v1_
        nw125_[int(6)] = d_883_v1_
        nw125_[int(7)] = d_883_v1_
        nw125_[int(8)] = d_883_v1_
        nw125_[int(9)] = d_883_v1_
        def iife78_(_pat_let18_0):
            def iife79_(d_887_dt__update__tmp_h0_):
                def iife80_(_pat_let19_0):
                    def iife81_(d_890_dt__update_hcf5_h0_):
                        return D1_DC3(d_890_dt__update_hcf5_h0_)
                    return iife81_(_pat_let19_0)
                return iife80_(default__.fm22(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_888_i0_ in range(default__.abs(588))])), (self).f10, (self).fm0(pat_let_tv18_, False, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_889_i1_ in range(default__.abs(-736))])), (self).f10, pat_let_tv19_), pat_let_tv20_, pat_let_tv21_))
            return iife79_(_pat_let18_0)
        nw125_[int(10)] = iife78_(d_883_v1_)
        nw125_[int(11)] = d_883_v1_
        nw125_[int(12)] = D1_DC3(d_882_v0_)
        nw125_[int(13)] = d_883_v1_
        nw125_[int(14)] = d_883_v1_
        nw125_[int(15)] = d_883_v1_
        nw125_[int(16)] = d_883_v1_
        nw125_[int(17)] = D1_DC3(d_882_v0_)
        nw125_[int(18)] = d_883_v1_
        d_886_v4_ = nw125_
        d_891_v5_: _dafny.Map
        d_891_v5_ = _dafny.Map({d_885_v3_: d_886_v4_})
        d_886_v4_ = (((d_891_v5_)[67] if (67) in (d_891_v5_) else d_886_v4_) if (self).f10 else d_886_v4_)
        d_882_v0_ = d_882_v0_
        r1 = (self).f10
        (globalState).f8 = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bvcipvb"))) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ruf")))
        d_892_v6_: C0
        nw126_ = C0()
        nw126_.ctor__()
        d_892_v6_ = nw126_
        d_892_v6_ = d_892_v6_
        d_885_v3_ = d_885_v3_
        d_893_v7_: _dafny.MultiSet
        d_893_v7_ = _dafny.MultiSet([(self).f10])
        d_894_v8_: _dafny.Map
        d_894_v8_ = _dafny.Map({(d_893_v7_).cardinality: len(d_882_v0_)})
        r0 = (d_882_v0_).set(default__.safeIndex(d_885_v3_, len(d_882_v0_)), (d_882_v0_)[default__.safeIndex(len(d_894_v8_), len(d_882_v0_))])
        r1 = (self).f10
        r2 = _dafny.SeqWithoutIsStrInference([_dafny.Set({(self).f10, (self).f10, not(True)}) for d_895_i2_ in range(default__.abs(592))])
        r3 = d_882_v0_
        return r0, r1, r2, r3

    def m1(self, p0, globalState):
        r0: _dafny.Map = _dafny.Map({})
        r1: bool = False
        d_896_i0_: int
        d_896_i0_ = 0
        with _dafny.label("2"):
            while not((((self).f10 if (self).f10 else (self).f10)) or ((not((self).f10) if (self).f10 else (self).f10))):
                with _dafny.c_label("2"):
                    if (d_896_i0_) >= (100):
                        raise _dafny.Break("2")
                    d_896_i0_ = (d_896_i0_) + (1)
                    d_897_v0_: _dafny.Seq
                    d_897_v0_ = _dafny.SeqWithoutIsStrInference([(self).f10])
                    d_898_v1_: int
                    d_898_v1_ = 169
                    d_899_v2_: _dafny.Map
                    d_899_v2_ = _dafny.Map({d_898_v1_: not((self).f10)})
                    d_900_v3_: _dafny.Set
                    d_900_v3_ = _dafny.Set({(self).fm0(d_897_v0_, (self).f10, len(d_899_v2_), (self).f10, globalState)})
                    d_901_v4_: _dafny.Map
                    d_901_v4_ = _dafny.Map({(False if (self).f10 else not((self).f10)): d_900_v3_})
                    d_902_v5_: _dafny.Array
                    nw127_ = _dafny.Array(int(0), 14)
                    d_902_v5_ = nw127_
                    index146_ = default__.safeIndex(5, (d_902_v5_).length(0))
                    (d_902_v5_)[index146_] = d_898_v1_
                    index147_ = default__.safeIndex(5, (d_902_v5_).length(0))
                    rhs138_ = (d_901_v4_) | (d_901_v4_)
                    rhs139_ = (0) - (d_898_v1_)
                    lhs104_ = d_902_v5_
                    lhs105_ = default__.safeIndex(5, (d_902_v5_).length(0))
                    d_901_v4_ = rhs138_
                    lhs104_[lhs105_] = rhs139_
                    (globalState).f8 = (self).f10
                    d_903_v6_: _dafny.MultiSet
                    d_903_v6_ = _dafny.MultiSet([False])
                    d_904_v7_: _dafny.Map
                    d_904_v7_ = _dafny.Map({_dafny.Map({(_dafny.MultiSet([d_898_v1_])).cardinality: d_903_v6_}): d_898_v1_})
                    d_905_v8_: _dafny.Map
                    d_905_v8_ = _dafny.Map({(d_902_v5_)[default__.safeIndex(5, (d_902_v5_).length(0))]: d_903_v6_})
                    d_904_v7_ = (d_904_v7_).set((d_905_v8_) | (d_905_v8_), d_898_v1_)
                    d_906_v10_: _dafny.Set
                    d_906_v10_ = _dafny.Set({p0})
                    def iife82_():
                        coll42_ = _dafny.Map()
                        compr_42_: str
                        for compr_42_ in (d_906_v10_).Elements:
                            d_907_v9_: str = compr_42_
                            if (d_907_v9_) in (d_906_v10_):
                                coll42_[d_907_v9_] = len(_dafny.SeqWithoutIsStrInference([p0 for d_908_i1_ in range(default__.abs(186))]))
                        return _dafny.Map(coll42_)
                    d_898_v1_ = ((d_902_v5_)[default__.safeIndex(5, (d_902_v5_).length(0))]) - (((self).fm1((self).f9, globalState)) + (len((iife82_()
                    ).set(p0, (d_902_v5_)[default__.safeIndex(5, (d_902_v5_).length(0))]))))
                    pass
            pass
        d_909_v11_: _dafny.Seq
        d_909_v11_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({(self).f10, True})])
        d_910_v12_: int
        d_910_v12_ = -285
        d_911_v13_: _dafny.Set
        d_911_v13_ = _dafny.Set({(self).f10, (self).f10})
        if ((d_909_v11_)[default__.safeIndex(d_910_v12_, len(d_909_v11_))]) != (d_911_v13_):
            (globalState).f8 = (self).f10
            r1 = (self).f10
            d_912_v14_: D0
            d_912_v14_ = D0_DC0((self).f10)
            d_913_v15_: _dafny.Map
            d_913_v15_ = _dafny.Map({(d_912_v14_ if (self).f10 else d_912_v14_): not(not ((self).f10) or ((self).f10))})
            d_913_v15_ = (d_913_v15_).set(default__.fm23(globalState), (self).f10)
            d_914_v16_: _dafny.Seq
            d_914_v16_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hlsgfllkl"))
            if ((len(d_914_v16_)) != (d_910_v12_)) == ((self).f10):
                d_915_v17_: _dafny.Map
                d_915_v17_ = _dafny.Map({p0: d_910_v12_})
                d_915_v17_ = (d_915_v17_).set(p0, d_910_v12_)
                d_910_v12_ = d_910_v12_
                d_916_v18_: _dafny.Array
                nw128_ = _dafny.Array(False, 29)
                d_916_v18_ = nw128_
                d_916_v18_ = d_916_v18_
                d_917_v19_: _dafny.Array
                nw129_ = _dafny.Array(_dafny.Map({}), 26)
                d_917_v19_ = nw129_
                d_918_v20_: _dafny.Array
                def lambda39_(d_919_i2_):
                    return (d_919_i2_) + (len(_dafny.SeqWithoutIsStrInference([(self).f10])))

                init21_ = lambda39_
                nw130_ = _dafny.Array(None, 29)
                for i0_21_ in range(nw130_.length(0)):
                    nw130_[i0_21_] = init21_(i0_21_)
                d_918_v20_ = nw130_
                d_920_v21_: D3
                d_920_v21_ = D3_DC7(d_918_v20_)
                d_921_v22_: _dafny.Map
                d_921_v22_ = _dafny.Map({(d_920_v21_).cf11: (self).fm1((self).f9, globalState)})
                index148_ = default__.safeIndex(846, (d_917_v19_).length(0))
                (d_917_v19_)[index148_] = d_921_v22_
                index149_ = default__.safeIndex(846, (d_917_v19_).length(0))
                (d_917_v19_)[index149_] = (_dafny.Map({d_918_v20_: (self).fm1((self).f9, globalState)})) | (d_921_v22_)
                d_922_v23_: _dafny.Seq
                d_922_v23_ = _dafny.SeqWithoutIsStrInference([762])
                d_923_v24_: _dafny.Seq
                d_923_v24_ = _dafny.SeqWithoutIsStrInference([(self).f10])
                d_922_v23_ = (d_922_v23_ if (self).fm0(d_923_v24_, (self).f10, len(_dafny.SeqWithoutIsStrInference([d_910_v12_, d_910_v12_])), not((self).f10), globalState) else d_922_v23_)
            elif True:
                d_924_v25_: _dafny.Array
                nw131_ = _dafny.Array(False, 7)
                d_924_v25_ = nw131_
                index150_ = default__.safeIndex(728, (d_924_v25_).length(0))
                (d_924_v25_)[index150_] = not ((self).f10) or ((self).f10)
                d_925_v26_: _dafny.Seq
                d_925_v26_ = _dafny.SeqWithoutIsStrInference([(self).f10])
                d_926_v27_: _dafny.Map
                d_926_v27_ = _dafny.Map({(d_925_v26_) + (d_925_v26_): len(_dafny.Set({d_910_v12_, d_910_v12_}))})
                d_927_v28_: _dafny.Set
                d_927_v28_ = _dafny.Set({(self).f9})
                index151_ = default__.safeIndex(728, (d_924_v25_).length(0))
                rhs140_ = d_924_v25_
                rhs141_ = ((self).fm0(d_925_v26_, False, d_910_v12_, (self).f10, globalState) if (self).f10 else (self).f10)
                rhs142_ = default__.safeModuloInt(d_910_v12_, d_910_v12_)
                rhs143_ = (default__.fm24((self).f10, globalState)).isdisjoint(d_927_v28_)
                rhs144_ = d_926_v27_
                lhs106_ = d_924_v25_
                lhs107_ = default__.safeIndex(728, (d_924_v25_).length(0))
                lhs108_ = globalState
                d_924_v25_ = rhs140_
                lhs106_[lhs107_] = rhs141_
                d_910_v12_ = rhs142_
                lhs108_.f8 = rhs143_
                d_926_v27_ = rhs144_
                d_928_v29_: C0
                nw132_ = C0()
                nw132_.ctor__()
                d_928_v29_ = nw132_
                d_929_v30_: str
                d_929_v30_ = _dafny.CodePoint('h')
                d_930_v32_: _dafny.Map
                def iife83_():
                    coll43_ = _dafny.Map()
                    compr_43_: int
                    for compr_43_ in (_dafny.Map({d_910_v12_: (0) - (d_910_v12_)})).keys.Elements:
                        d_931_v31_: int = compr_43_
                        if (d_931_v31_) in (_dafny.Map({d_910_v12_: (0) - (d_910_v12_)})):
                            coll43_[default__.safeModuloInt(d_931_v31_, d_910_v12_)] = len(_dafny.Map({d_910_v12_: d_910_v12_}))
                    return _dafny.Map(coll43_)
                d_930_v32_ = _dafny.Map({(d_924_v25_)[default__.safeIndex(728, (d_924_v25_).length(0))]: (_dafny.Map({len(default__.fm25((0) - (d_910_v12_), len(d_914_v16_), globalState)): d_910_v12_})) | ((_dafny.Map({d_910_v12_: d_910_v12_})).set(len(d_914_v16_), len(iife83_()
                )))})
                d_932_v33_: _dafny.Array
                nw133_ = _dafny.Array(int(0), 9)
                d_932_v33_ = nw133_
                d_933_v34_: _dafny.Map
                d_933_v34_ = _dafny.Map({323: d_910_v12_})
                rhs145_ = d_929_v30_
                rhs146_ = _dafny.Map({(d_924_v25_)[default__.safeIndex(728, (d_924_v25_).length(0))]: (d_933_v34_) | (d_933_v34_)})
                rhs147_ = d_932_v33_
                d_929_v30_ = rhs145_
                d_930_v32_ = rhs146_
                d_932_v33_ = rhs147_
                d_934_v35_: _dafny.Map
                d_934_v35_ = _dafny.Map({d_929_v30_: (self).f10})
                d_935_v36_: _dafny.Map
                d_935_v36_ = _dafny.Map({(self).fm0(d_925_v26_, False, d_910_v12_, (self).fm0(d_925_v26_, (self).f10, d_910_v12_, (self).f10, globalState), globalState): (D1_DC4((self).f10, ((d_934_v35_)[d_929_v30_] if (d_929_v30_) in (d_934_v35_) else (d_924_v25_)[default__.safeIndex(728, (d_924_v25_).length(0))]), (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pojq")))))).cf7})
                d_935_v36_ = (d_935_v36_).set((self).f10, True)
                d_936_v37_: _dafny.Seq
                d_936_v37_ = _dafny.SeqWithoutIsStrInference([d_914_v16_])
                index152_ = default__.safeIndex(728, (d_924_v25_).length(0))
                (d_924_v25_)[index152_] = ((d_936_v37_).set(default__.safeIndex(d_910_v12_, len(d_936_v37_)), _dafny.SeqWithoutIsStrInference([d_929_v30_ for d_937_i3_ in range(default__.abs(-464))]))) == (d_936_v37_)
            d_938_v38_: _dafny.MultiSet
            d_938_v38_ = _dafny.MultiSet([(0) - (default__.safeModuloInt(d_910_v12_, 391))])
            d_938_v38_ = (d_938_v38_) | (d_938_v38_)
        elif True:
            d_939_v39_: C0
            nw134_ = C0()
            nw134_.ctor__()
            d_939_v39_ = nw134_
            d_940_v40_: _dafny.Seq
            d_940_v40_ = _dafny.SeqWithoutIsStrInference([d_939_v39_, d_939_v39_])
            d_941_v41_: _dafny.Seq
            d_941_v41_ = _dafny.SeqWithoutIsStrInference([(self).f10, (self).f10])
            rhs148_ = (len(d_941_v41_) if (self).f10 else d_910_v12_)
            rhs149_ = (self).f10
            rhs150_ = (self).f10
            rhs151_ = ((d_940_v40_) + (d_940_v40_)) + (d_940_v40_)
            lhs109_ = globalState
            lhs110_ = globalState
            d_910_v12_ = rhs148_
            lhs109_.f8 = rhs149_
            lhs110_.f8 = rhs150_
            d_940_v40_ = rhs151_
            d_942_v42_: _dafny.Array
            nw135_ = _dafny.Array(False, 18)
            d_942_v42_ = nw135_
            d_943_v43_: _dafny.Seq
            d_943_v43_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "coavp"))
            d_944_v44_: _dafny.MultiSet
            d_944_v44_ = _dafny.MultiSet([len(d_943_v43_)])
            index153_ = default__.safeIndex(765, (d_942_v42_).length(0))
            (d_942_v42_)[index153_] = (d_944_v44_) == (d_944_v44_)
            index154_ = default__.safeIndex(765, (d_942_v42_).length(0))
            rhs152_ = (self).f10
            rhs153_ = (self).f10
            rhs154_ = (self).f10
            lhs111_ = globalState
            lhs112_ = d_942_v42_
            lhs113_ = default__.safeIndex(765, (d_942_v42_).length(0))
            lhs114_ = globalState
            lhs111_.f8 = rhs152_
            lhs112_[lhs113_] = rhs153_
            lhs114_.f8 = rhs154_
            d_910_v12_ = default__.safeDivisionInt((d_910_v12_ if (d_942_v42_)[default__.safeIndex(765, (d_942_v42_).length(0))] else d_910_v12_), (0) - (default__.safeDivisionInt(323, d_910_v12_)))
            d_945_v45_: _dafny.Array
            nw136_ = _dafny.Array(None, 12)
            nw136_[int(0)] = p0
            nw136_[int(1)] = p0
            nw136_[int(2)] = p0
            nw136_[int(3)] = p0
            nw136_[int(4)] = p0
            nw136_[int(5)] = p0
            nw136_[int(6)] = p0
            nw136_[int(7)] = p0
            nw136_[int(8)] = _dafny.CodePoint('c')
            nw136_[int(9)] = (default__.fm26(p0, globalState) if (d_942_v42_)[default__.safeIndex(765, (d_942_v42_).length(0))] else p0)
            nw136_[int(10)] = p0
            nw136_[int(11)] = p0
            d_945_v45_ = nw136_
            d_945_v45_ = (d_945_v45_ if not((self).f10) else d_945_v45_)
            d_946_v46_: _dafny.Map
            d_946_v46_ = _dafny.Map({(self).f10: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kxm"))})
            d_947_v47_: _dafny.Seq
            d_947_v47_ = _dafny.SeqWithoutIsStrInference([((d_946_v46_)[(d_942_v42_)[default__.safeIndex(765, (d_942_v42_).length(0))]] if ((d_942_v42_)[default__.safeIndex(765, (d_942_v42_).length(0))]) in (d_946_v46_) else d_943_v43_), d_943_v43_])
            (globalState).f8 = (p0) not in ((d_947_v47_)[default__.safeIndex(d_910_v12_, len(d_947_v47_))])
        d_948_v48_: _dafny.Seq
        d_948_v48_ = _dafny.SeqWithoutIsStrInference([d_910_v12_, (d_910_v12_) - (d_910_v12_), d_910_v12_])
        d_910_v12_ = (d_948_v48_)[default__.safeIndex(d_910_v12_, len(d_948_v48_))]
        r1 = (self).f10
        d_949_v49_: C0
        nw137_ = C0()
        nw137_.ctor__()
        d_949_v49_ = nw137_
        d_950_v50_: _dafny.Seq
        d_950_v50_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f"))
        rhs155_ = not((d_950_v50_) == (d_950_v50_))
        rhs156_ = (self).f10
        r1 = rhs155_
        r1 = rhs156_
        d_951_v51_: _dafny.Array
        nw138_ = _dafny.Array(int(0), 3)
        d_951_v51_ = nw138_
        d_952_v52_: _dafny.Map
        d_952_v52_ = _dafny.Map({d_951_v51_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xxghgxtt"))})
        r0 = (d_952_v52_) | (_dafny.Map({d_951_v51_: d_950_v50_}))
        r1 = (self).f10
        return r0, r1

    def m5(self, globalState):
        d_953_v0_: _dafny.Seq
        d_953_v0_ = _dafny.SeqWithoutIsStrInference([(self).f10, (self).f10])
        (globalState).f8 = (not((self).f10)) in (d_953_v0_)
        (globalState).f8 = (self).f10
        d_954_i0_: int
        d_954_i0_ = 0
        with _dafny.label("3"):
            while (self).f10:
                with _dafny.c_label("3"):
                    if (d_954_i0_) >= (100):
                        raise _dafny.Break("3")
                    d_954_i0_ = (d_954_i0_) + (1)
                    if (self).f10:
                        d_955_v1_: int
                        d_955_v1_ = 762
                        d_955_v1_ = ((0) - (d_955_v1_)) - (d_955_v1_)
                        d_956_v2_: _dafny.Seq
                        d_956_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qqk"))
                        d_957_v3_: _dafny.Array
                        nw139_ = _dafny.Array(int(0), 5)
                        d_957_v3_ = nw139_
                        d_958_v4_: _dafny.Map
                        d_958_v4_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsdcjxlyo")): d_957_v3_})
                        d_959_v5_: _dafny.Seq
                        d_960_v6_: int
                        d_961_v7_: _dafny.Seq
                        d_962_v8_: int
                        out48_: _dafny.Seq
                        out49_: int
                        out50_: _dafny.Seq
                        out51_: int
                        out48_, out49_, out50_, out51_ = (self).m6(len(d_956_v2_), ((d_958_v4_)[d_956_v2_] if (d_956_v2_) in (d_958_v4_) else d_957_v3_), (d_955_v1_) * (d_955_v1_), globalState)
                        d_959_v5_ = out48_
                        d_960_v6_ = out49_
                        d_961_v7_ = out50_
                        d_962_v8_ = out51_
                        d_963_v9_: str
                        d_963_v9_ = _dafny.CodePoint('f')
                        d_963_v9_ = d_963_v9_
                        d_956_v2_ = _dafny.SeqWithoutIsStrInference([d_963_v9_])
                        d_964_v10_: T0
                        nw140_ = C3()
                        nw140_.ctor__()
                        d_964_v10_ = nw140_
                        d_964_v10_ = d_964_v10_
                    elif True:
                        d_965_v11_: int
                        d_965_v11_ = 153
                        d_965_v11_ = d_965_v11_
                        d_966_v12_: _dafny.Array
                        def lambda40_(d_967_i1_):
                            return not((self).f10)

                        init22_ = lambda40_
                        nw141_ = _dafny.Array(None, 7)
                        for i0_22_ in range(nw141_.length(0)):
                            nw141_[i0_22_] = init22_(i0_22_)
                        d_966_v12_ = nw141_
                        index155_ = default__.safeIndex(684, (d_966_v12_).length(0))
                        (d_966_v12_)[index155_] = (self).f10
                        index156_ = default__.safeIndex(684, (d_966_v12_).length(0))
                        (d_966_v12_)[index156_] = not((d_953_v0_)[default__.safeIndex(d_965_v11_, len(d_953_v0_))])
                        d_968_v13_: _dafny.Seq
                        d_968_v13_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xum"))
                        d_968_v13_ = _dafny.SeqWithoutIsStrInference([(D4_DC12((d_966_v12_)[default__.safeIndex(684, (d_966_v12_).length(0))], 357, _dafny.CodePoint('d'), (0) - (d_965_v11_), d_965_v11_)).cf21 for d_969_i2_ in range(default__.abs(-636))])
                        d_970_v14_: C1
                        nw142_ = C1()
                        nw142_.ctor__((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_971_i3_ in range(default__.abs(792))])) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sixbvx"))), d_968_v13_)
                        d_970_v14_ = nw142_
                        d_972_v15_: _dafny.Seq
                        d_972_v15_ = _dafny.SeqWithoutIsStrInference([d_965_v11_, d_965_v11_, d_965_v11_, 369, len(_dafny.Set({d_965_v11_}))])
                        d_973_v16_: _dafny.Seq
                        d_973_v16_ = _dafny.SeqWithoutIsStrInference([d_972_v15_])
                        d_974_v17_: _dafny.Seq
                        d_974_v17_ = _dafny.SeqWithoutIsStrInference([(d_973_v16_)[default__.safeIndex(d_965_v11_, len(d_973_v16_))]])
                        d_975_v18_: _dafny.Seq
                        d_975_v18_ = _dafny.SeqWithoutIsStrInference([(len((d_973_v16_)[default__.safeIndex(len((d_970_v14_).f13), len(d_973_v16_))])) - (650)])
                        d_972_v15_ = d_975_v18_
                    d_976_v19_: int
                    d_976_v19_ = 745
                    d_977_v20_: _dafny.Set
                    d_977_v20_ = _dafny.Set({True})
                    d_976_v19_ = (default__.safeDivisionInt(d_976_v19_, 426) if (d_977_v20_).ispropersubset(d_977_v20_) else d_976_v19_)
                    (globalState).f8 = (self).f10
                    d_978_v21_: _dafny.Array
                    nw143_ = _dafny.Array(False, 25)
                    d_978_v21_ = nw143_
                    index157_ = default__.safeIndex(103, (d_978_v21_).length(0))
                    (d_978_v21_)[index157_] = (self).f10
                    d_979_v22_: _dafny.Seq
                    d_979_v22_ = _dafny.SeqWithoutIsStrInference([d_976_v19_])
                    d_980_v23_: _dafny.Map
                    d_980_v23_ = _dafny.Map({(self).f10: (0) - ((self).fm1((self).f9, globalState))})
                    d_981_v24_: _dafny.Set
                    d_981_v24_ = _dafny.Set({d_980_v23_})
                    index158_ = default__.safeIndex(103, (d_978_v21_).length(0))
                    (d_978_v21_)[index158_] = (d_979_v22_) <= (_dafny.SeqWithoutIsStrInference([((d_980_v23_)[(self).f10] if ((self).f10) in (d_980_v23_) else -128), 411, len(d_981_v24_), d_976_v19_]))
                    pass
            pass
        d_982_i4_: int
        d_982_i4_ = 0
        with _dafny.label("4"):
            while (self).f10:
                with _dafny.c_label("4"):
                    if (d_982_i4_) >= (100):
                        raise _dafny.Break("4")
                    d_982_i4_ = (d_982_i4_) + (1)
                    if ((self).f10) == (not((self).f10)):
                        d_983_v25_: _dafny.Seq
                        d_983_v25_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yuctjfpul"))
                        d_984_v26_: C1
                        nw144_ = C1()
                        nw144_.ctor__((self).f10, d_983_v25_)
                        d_984_v26_ = nw144_
                        d_985_v27_: int
                        d_985_v27_ = 729
                        d_986_v28_: _dafny.Map
                        d_986_v28_ = _dafny.Map({d_985_v27_: (self).f9})
                        d_987_v29_: _dafny.Map
                        d_987_v29_ = _dafny.Map({(d_986_v28_) | (d_986_v28_): (d_985_v27_) + (len((d_984_v26_).f13))})
                        d_988_v30_: _dafny.Map
                        d_988_v30_ = _dafny.Map({(d_984_v26_).f12: d_985_v27_})
                        d_987_v29_ = (d_987_v29_).set(d_986_v28_, default__.safeModuloInt(d_985_v27_, ((d_988_v30_)[(self).f10] if ((self).f10) in (d_988_v30_) else d_985_v27_)))
                        d_989_v31_: _dafny.Seq
                        d_989_v31_ = _dafny.SeqWithoutIsStrInference([(d_984_v26_).f13, (d_984_v26_).f13, (d_984_v26_).f13])
                        d_983_v25_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w"))) + (((d_989_v31_)[default__.safeIndex(d_985_v27_, len(d_989_v31_))]) + (d_983_v25_))
                        d_985_v27_ = len(((self).f9) | ((self).f9))
                        d_990_v32_: _dafny.MultiSet
                        d_990_v32_ = _dafny.MultiSet([d_985_v27_])
                        d_988_v30_ = (d_988_v30_).set((d_990_v32_).ispropersubset(d_990_v32_), d_985_v27_)
                    elif True:
                        d_991_v33_: _dafny.Seq
                        d_991_v33_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cckevci"))
                        d_992_v34_: int
                        d_992_v34_ = 419
                        d_993_v35_: _dafny.Seq
                        d_993_v35_ = _dafny.SeqWithoutIsStrInference([d_992_v34_])
                        d_994_v36_: _dafny.Array
                        def lambda41_(d_995_i5_):
                            return _dafny.CodePoint('n')

                        init23_ = lambda41_
                        nw145_ = _dafny.Array(None, 29)
                        for i0_23_ in range(nw145_.length(0)):
                            nw145_[i0_23_] = init23_(i0_23_)
                        d_994_v36_ = nw145_
                        d_996_v37_: _dafny.MultiSet
                        d_996_v37_ = _dafny.MultiSet([d_992_v34_])
                        rhs157_ = d_991_v33_
                        rhs158_ = ((d_993_v35_) + (_dafny.SeqWithoutIsStrInference([d_992_v34_, d_992_v34_]))) + (_dafny.SeqWithoutIsStrInference([d_992_v34_ for d_997_i6_ in range(default__.abs(750))]))
                        rhs159_ = d_994_v36_
                        rhs160_ = default__.safeDivisionInt(d_992_v34_, d_992_v34_)
                        rhs161_ = default__.safeDivisionInt(d_992_v34_, ((d_996_v37_).intersection(d_996_v37_)).cardinality)
                        d_991_v33_ = rhs157_
                        d_993_v35_ = rhs158_
                        d_994_v36_ = rhs159_
                        d_992_v34_ = rhs160_
                        d_992_v34_ = rhs161_
                        (globalState).f8 = (self).f10
                        d_998_v38_: _dafny.MultiSet
                        d_998_v38_ = _dafny.MultiSet([not((self).f10)])
                        d_999_v39_: _dafny.Seq
                        d_999_v39_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet(d_953_v0_), d_998_v38_])
                        d_1000_v40_: _dafny.Map
                        d_1000_v40_ = _dafny.Map({(self).f10: _dafny.CodePoint('g')})
                        d_1001_v41_: _dafny.MultiSet
                        d_1001_v41_ = _dafny.MultiSet([default__.fm22(len(_dafny.SeqWithoutIsStrInference([d_992_v34_])), (self).f10, (self).fm0(d_953_v0_, (self).f10, (self).fm1((self).f9, globalState), (self).f10, globalState), d_992_v34_, globalState)])
                        d_1002_v42_: _dafny.Array
                        nw146_ = _dafny.Array(None, 5)
                        nw146_[int(0)] = (self).f10
                        nw146_[int(1)] = (self).f10
                        nw146_[int(2)] = (self).f10
                        nw146_[int(3)] = not(False)
                        nw146_[int(4)] = (self).f10
                        d_1002_v42_ = nw146_
                        d_1003_v43_: _dafny.Seq
                        d_1003_v43_ = _dafny.SeqWithoutIsStrInference([d_1002_v42_])
                        d_1004_v44_: _dafny.Array
                        nw147_ = _dafny.Array(None, 25)
                        nw147_[int(0)] = d_992_v34_
                        nw147_[int(1)] = (d_992_v34_) * (len(d_991_v33_))
                        nw147_[int(2)] = d_992_v34_
                        nw147_[int(3)] = d_992_v34_
                        nw147_[int(4)] = ((d_999_v39_)[default__.safeIndex(531, len(d_999_v39_))]).cardinality
                        nw147_[int(5)] = d_992_v34_
                        nw147_[int(6)] = d_992_v34_
                        nw147_[int(7)] = d_992_v34_
                        nw147_[int(8)] = (d_992_v34_) - (915)
                        nw147_[int(9)] = d_992_v34_
                        nw147_[int(10)] = d_992_v34_
                        nw147_[int(11)] = (len(d_1000_v40_) if (self).f10 else d_992_v34_)
                        nw147_[int(12)] = default__.safeModuloInt(d_992_v34_, d_992_v34_)
                        nw147_[int(13)] = (0) - (d_992_v34_)
                        nw147_[int(14)] = ((0) - ((0) - (d_992_v34_)) if (self).f10 else d_992_v34_)
                        nw147_[int(15)] = (d_1001_v41_).cardinality
                        nw147_[int(16)] = (d_992_v34_) * (-334)
                        nw147_[int(17)] = d_992_v34_
                        nw147_[int(18)] = d_992_v34_
                        nw147_[int(19)] = ((d_996_v37_)[(0) - (d_992_v34_)] if ((0) - (d_992_v34_)) in (d_996_v37_) else d_992_v34_)
                        nw147_[int(20)] = (d_992_v34_) * (d_992_v34_)
                        nw147_[int(21)] = (0) - (d_992_v34_)
                        nw147_[int(22)] = len((d_1003_v43_ if (self).f10 else d_1003_v43_))
                        nw147_[int(23)] = d_992_v34_
                        nw147_[int(24)] = 236
                        d_1004_v44_ = nw147_
                        index159_ = default__.safeIndex(163, (d_1004_v44_).length(0))
                        (d_1004_v44_)[index159_] = default__.safeModuloInt(d_992_v34_, d_992_v34_)
                        d_1005_v45_: str
                        d_1005_v45_ = _dafny.CodePoint('u')
                        index160_ = default__.safeIndex(163, (d_1004_v44_).length(0))
                        (d_1004_v44_)[index160_] = ((0) - (len((d_991_v33_).set(default__.safeIndex(284, len(d_991_v33_)), d_1005_v45_)))) - ((0) - (d_992_v34_))
                        index161_ = default__.safeIndex(232, (d_1002_v42_).length(0))
                        (d_1002_v42_)[index161_] = (self).f10
                        index162_ = default__.safeIndex(232, (d_1002_v42_).length(0))
                        rhs162_ = (default__.fm46((self).f10, (d_953_v0_)[default__.safeIndex(d_992_v34_, len(d_953_v0_))], not((self).f10), globalState)).issubset(d_996_v37_)
                        rhs163_ = (self).f10
                        lhs115_ = d_1002_v42_
                        lhs116_ = default__.safeIndex(232, (d_1002_v42_).length(0))
                        lhs117_ = globalState
                        lhs115_[lhs116_] = rhs162_
                        lhs117_.f8 = rhs163_
                        d_1006_v46_: _dafny.Map
                        d_1006_v46_ = _dafny.Map({d_992_v34_: d_992_v34_})
                        d_1006_v46_ = (d_1006_v46_).set(154, (d_1004_v44_)[default__.safeIndex(163, (d_1004_v44_).length(0))])
                    d_1007_v47_: _dafny.Array
                    nw148_ = _dafny.Array(int(0), 16)
                    d_1007_v47_ = nw148_
                    d_1008_v48_: int
                    d_1008_v48_ = 915
                    index163_ = default__.safeIndex(505, (d_1007_v47_).length(0))
                    (d_1007_v47_)[index163_] = d_1008_v48_
                    index164_ = default__.safeIndex(505, (d_1007_v47_).length(0))
                    (d_1007_v47_)[index164_] = default__.safeDivisionInt((d_1008_v48_) - (d_1008_v48_), d_1008_v48_)
                    d_1009_v49_: _dafny.Array
                    nw149_ = _dafny.Array(None, 6)
                    d_1009_v49_ = nw149_
                    d_1010_v50_: _dafny.Array
                    nw150_ = _dafny.Array(False, 3)
                    d_1010_v50_ = nw150_
                    d_1011_v51_: T1
                    nw151_ = C2()
                    nw151_.ctor__(not((self).f10), d_1010_v50_, (self).f9, (self).f10)
                    d_1011_v51_ = nw151_
                    index165_ = default__.safeIndex(479, (d_1009_v49_).length(0))
                    (d_1009_v49_)[index165_] = d_1011_v51_
                    d_1012_v52_: _dafny.Map
                    d_1012_v52_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "krxj"))) for d_1013_i7_ in range(default__.abs(-785))]): d_1011_v51_})
                    d_1014_v54_: _dafny.Set
                    d_1014_v54_ = _dafny.Set({(d_1011_v51_).f10})
                    d_1015_v55_: _dafny.Seq
                    def iife84_():
                        coll44_ = _dafny.Set()
                        compr_44_: int
                        for compr_44_ in _dafny.IntegerRange(525, 392):
                            d_1016_v53_: int = compr_44_
                            if ((525) <= (d_1016_v53_)) and ((d_1016_v53_) < (392)):
                                coll44_ = coll44_.union(_dafny.Set([(d_1016_v53_) * (d_1008_v48_)]))
                        return _dafny.Set(coll44_)
                    d_1015_v55_ = _dafny.SeqWithoutIsStrInference([len(iife84_()
                    ), len(d_1014_v54_)])
                    d_1017_v56_: _dafny.Seq
                    d_1017_v56_ = _dafny.SeqWithoutIsStrInference([d_1015_v55_])
                    index166_ = default__.safeIndex(479, (d_1009_v49_).length(0))
                    (d_1009_v49_)[index166_] = ((d_1012_v52_)[(d_1017_v56_)[default__.safeIndex((d_1007_v47_)[default__.safeIndex(505, (d_1007_v47_).length(0))], len(d_1017_v56_))]] if ((d_1017_v56_)[default__.safeIndex((d_1007_v47_)[default__.safeIndex(505, (d_1007_v47_).length(0))], len(d_1017_v56_))]) in (d_1012_v52_) else d_1011_v51_)
                    d_1018_v57_: _dafny.Array
                    nw152_ = _dafny.Array(_dafny.Seq({}), 8)
                    d_1018_v57_ = nw152_
                    d_1019_v58_: C3
                    nw153_ = C3()
                    nw153_.ctor__()
                    d_1019_v58_ = nw153_
                    d_1020_v59_: _dafny.Seq
                    d_1020_v59_ = _dafny.SeqWithoutIsStrInference([d_1019_v58_])
                    index167_ = default__.safeIndex(984, (d_1018_v57_).length(0))
                    (d_1018_v57_)[index167_] = d_1020_v59_
                    index168_ = default__.safeIndex(984, (d_1018_v57_).length(0))
                    (d_1018_v57_)[index168_] = d_1020_v59_
                    pass
            pass
        if (self).f10:
            d_1021_v60_: _dafny.Seq
            d_1021_v60_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tndwyncs"))
            d_1022_v61_: _dafny.MultiSet
            d_1022_v61_ = _dafny.MultiSet([d_1021_v60_])
            d_1023_v62_: int
            d_1023_v62_ = 469
            d_1024_v63_: D11
            d_1024_v63_ = D11_DC27((d_1022_v61_).set(d_1021_v60_, default__.abs(d_1023_v62_)))
            if (_dafny.MultiSet([d_1021_v60_])).issubset((d_1024_v63_).cf47):
                d_1023_v62_ = d_1023_v62_
                d_1025_v64_: _dafny.Array
                nw154_ = _dafny.Array(int(0), 7)
                d_1025_v64_ = nw154_
                d_1026_v65_: _dafny.Set
                d_1026_v65_ = _dafny.Set({(self).f10, False})
                d_1027_v66_: D8
                d_1027_v66_ = D8_DC24(not((self).f10), _dafny.Set({d_1025_v64_}), len(d_1026_v65_))
                pat_let_tv22_ = globalState
                d_1028_v67_: _dafny.Array
                nw155_ = _dafny.Array(None, 8)
                nw155_[int(0)] = d_1027_v66_
                nw155_[int(1)] = d_1027_v66_
                nw155_[int(2)] = d_1027_v66_
                nw155_[int(3)] = d_1027_v66_
                nw155_[int(4)] = d_1027_v66_
                nw155_[int(5)] = d_1027_v66_
                nw155_[int(6)] = d_1027_v66_
                def iife85_(_pat_let20_0):
                    def iife86_(d_1029_dt__update__tmp_h0_):
                        def iife87_(_pat_let21_0):
                            def iife88_(d_1030_dt__update_hcf44_h0_):
                                return D8_DC24((d_1029_dt__update__tmp_h0_).cf42, (d_1029_dt__update__tmp_h0_).cf43, d_1030_dt__update_hcf44_h0_)
                            return iife88_(_pat_let21_0)
                        return iife87_((self).fm1((self).f9, pat_let_tv22_))
                    return iife86_(_pat_let20_0)
                nw155_[int(7)] = iife85_(d_1027_v66_)
                d_1028_v67_ = nw155_
                d_1031_v68_: _dafny.Seq
                d_1031_v68_ = _dafny.SeqWithoutIsStrInference([d_1028_v67_, d_1028_v67_])
                d_1032_v69_: D12
                d_1032_v69_ = D12_DC30(d_1028_v67_)
                d_1033_v70_: _dafny.Array
                nw156_ = _dafny.Array(None, 26)
                nw156_[int(0)] = d_1028_v67_
                nw156_[int(1)] = d_1028_v67_
                nw156_[int(2)] = d_1028_v67_
                nw156_[int(3)] = d_1028_v67_
                nw156_[int(4)] = d_1028_v67_
                nw156_[int(5)] = d_1028_v67_
                nw156_[int(6)] = d_1028_v67_
                nw156_[int(7)] = d_1028_v67_
                nw156_[int(8)] = d_1028_v67_
                nw156_[int(9)] = (d_1031_v68_)[default__.safeIndex(d_1023_v62_, len(d_1031_v68_))]
                nw156_[int(10)] = (d_1028_v67_ if (self).f10 else d_1028_v67_)
                nw156_[int(11)] = (d_1028_v67_ if False else d_1028_v67_)
                nw156_[int(12)] = d_1028_v67_
                nw156_[int(13)] = (d_1032_v69_).cf54
                nw156_[int(14)] = (d_1028_v67_ if (self).f10 else d_1028_v67_)
                nw156_[int(15)] = d_1028_v67_
                nw156_[int(16)] = d_1028_v67_
                nw156_[int(17)] = d_1028_v67_
                nw156_[int(18)] = d_1028_v67_
                nw156_[int(19)] = d_1028_v67_
                nw156_[int(20)] = d_1028_v67_
                nw156_[int(21)] = d_1028_v67_
                nw156_[int(22)] = (d_1028_v67_ if (self).f10 else d_1028_v67_)
                nw156_[int(23)] = d_1028_v67_
                nw156_[int(24)] = d_1028_v67_
                nw156_[int(25)] = d_1028_v67_
                d_1033_v70_ = nw156_
                d_1033_v70_ = d_1033_v70_
                (globalState).f8 = (self).f10
                d_1034_v71_: _dafny.Array
                nw157_ = _dafny.Array(False, 3)
                d_1034_v71_ = nw157_
                index169_ = default__.safeIndex(722, (d_1034_v71_).length(0))
                (d_1034_v71_)[index169_] = (d_1023_v62_) <= (d_1023_v62_)
                d_1035_v72_: _dafny.Map
                d_1035_v72_ = _dafny.Map({d_1021_v60_: (self).f10})
                d_1036_v73_: D1
                d_1036_v73_ = D1_DC4((self).f10, (self).f10, d_1023_v62_)
                d_1037_v74_: _dafny.Map
                d_1037_v74_ = _dafny.Map({d_1023_v62_: (d_1036_v73_).cf8})
                index170_ = default__.safeIndex(722, (d_1034_v71_).length(0))
                rhs164_ = ((d_1035_v72_)[d_1021_v60_] if (d_1021_v60_) in (d_1035_v72_) else (self).f10)
                rhs165_ = not((self).f10)
                rhs166_ = (self).f10
                rhs167_ = not((len((d_1037_v74_) | (_dafny.Map({d_1023_v62_: d_1023_v62_})))) >= (d_1023_v62_))
                lhs118_ = globalState
                lhs119_ = globalState
                lhs120_ = d_1034_v71_
                lhs121_ = default__.safeIndex(722, (d_1034_v71_).length(0))
                lhs122_ = globalState
                lhs118_.f8 = rhs164_
                lhs119_.f8 = rhs165_
                lhs120_[lhs121_] = rhs166_
                lhs122_.f8 = rhs167_
                d_1038_v75_: _dafny.MultiSet
                d_1038_v75_ = _dafny.MultiSet([(d_1034_v71_)[default__.safeIndex(722, (d_1034_v71_).length(0))], (self).f10])
                index171_ = default__.safeIndex(14, (d_1025_v64_).length(0))
                (d_1025_v64_)[index171_] = ((d_1038_v75_)[(self).fm0(d_953_v0_, (self).f10, d_1023_v62_, (self).f10, globalState)] if ((self).fm0(d_953_v0_, (self).f10, d_1023_v62_, (self).f10, globalState)) in (d_1038_v75_) else 819)
                d_1039_v76_: str
                d_1039_v76_ = _dafny.CodePoint('a')
                d_1040_v77_: _dafny.Map
                d_1040_v77_ = _dafny.Map({d_1025_v64_: d_1023_v62_})
                index172_ = default__.safeIndex(14, (d_1025_v64_).length(0))
                (d_1025_v64_)[index172_] = default__.safeModuloInt((D6_DC18(d_1039_v76_, len(d_1021_v60_), 142)).cf36, len((d_1040_v77_) | (d_1040_v77_)))
            elif True:
                d_1041_v78_: _dafny.Array
                def lambda42_(d_1042_i8_):
                    return (self).f10

                init24_ = lambda42_
                nw158_ = _dafny.Array(None, 27)
                for i0_24_ in range(nw158_.length(0)):
                    nw158_[i0_24_] = init24_(i0_24_)
                d_1041_v78_ = nw158_
                d_1043_v79_: T1
                nw159_ = C2()
                nw159_.ctor__((self).f10, (D5_DC13(d_1041_v78_)).cf24, (self).f9, (self).f10)
                d_1043_v79_ = nw159_
                nw160_ = C2()
                nw160_.ctor__((self).fm0(d_953_v0_, (d_1043_v79_).f10, d_1023_v62_, True, globalState), d_1041_v78_, (d_1043_v79_).f9, (d_1043_v79_).f10)
                d_1043_v79_ = nw160_
                d_1044_v80_: _dafny.Array
                nw161_ = _dafny.Array(None, 7)
                nw161_[int(0)] = d_953_v0_
                nw161_[int(1)] = (d_953_v0_) + (_dafny.SeqWithoutIsStrInference([(self).f10, (d_1043_v79_).f10]))
                nw161_[int(2)] = d_953_v0_
                nw161_[int(3)] = (d_953_v0_).set(default__.safeIndex((0) - (-31), len(d_953_v0_)), not((d_1043_v79_).f10))
                nw161_[int(4)] = (d_953_v0_) + (d_953_v0_)
                nw161_[int(5)] = (d_953_v0_) + (d_953_v0_)
                nw161_[int(6)] = (d_953_v0_) + (d_953_v0_)
                d_1044_v80_ = nw161_
                index173_ = default__.safeIndex(812, (d_1044_v80_).length(0))
                (d_1044_v80_)[index173_] = d_953_v0_
                index174_ = default__.safeIndex(812, (d_1044_v80_).length(0))
                rhs168_ = d_953_v0_
                rhs169_ = d_1021_v60_
                rhs170_ = (False if (self).f10 else (d_1043_v79_).f10)
                lhs123_ = d_1044_v80_
                lhs124_ = default__.safeIndex(812, (d_1044_v80_).length(0))
                lhs125_ = globalState
                lhs123_[lhs124_] = rhs168_
                d_1021_v60_ = rhs169_
                lhs125_.f8 = rhs170_
                d_1045_v81_: _dafny.Seq
                d_1045_v81_ = _dafny.SeqWithoutIsStrInference([(d_1044_v80_)[default__.safeIndex(812, (d_1044_v80_).length(0))]])
                d_1046_v82_: C1
                nw162_ = C1()
                nw162_.ctor__(not((self).f10), d_1021_v60_)
                d_1046_v82_ = nw162_
                d_1047_v83_: str
                d_1047_v83_ = _dafny.CodePoint('b')
                rhs171_ = ((d_953_v0_) + ((d_1045_v81_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_1046_v82_, d_1046_v82_])), len(d_1045_v81_))])) < (((d_1044_v80_)[default__.safeIndex(812, (d_1044_v80_).length(0))]) + (default__.fm38((d_1046_v82_).f12, globalState)))
                rhs172_ = not((self).fm0((d_1044_v80_)[default__.safeIndex(812, (d_1044_v80_).length(0))], (d_1043_v79_).f10, d_1023_v62_, (d_1043_v79_).f10, globalState))
                rhs173_ = (d_1046_v82_).fm32(d_1023_v62_, d_1047_v83_, globalState)
                lhs126_ = globalState
                lhs127_ = globalState
                lhs128_ = globalState
                lhs126_.f8 = rhs171_
                lhs127_.f8 = rhs172_
                lhs128_.f8 = rhs173_
                d_1048_v84_: _dafny.Array
                nw163_ = _dafny.Array(None, 13)
                nw163_[int(0)] = d_1047_v83_
                nw163_[int(1)] = d_1047_v83_
                nw163_[int(2)] = d_1047_v83_
                nw163_[int(3)] = d_1047_v83_
                nw163_[int(4)] = d_1047_v83_
                nw163_[int(5)] = d_1047_v83_
                nw163_[int(6)] = d_1047_v83_
                nw163_[int(7)] = d_1047_v83_
                nw163_[int(8)] = d_1047_v83_
                nw163_[int(9)] = d_1047_v83_
                nw163_[int(10)] = d_1047_v83_
                nw163_[int(11)] = d_1047_v83_
                nw163_[int(12)] = d_1047_v83_
                d_1048_v84_ = nw163_
                d_1049_v85_: _dafny.Array
                nw164_ = _dafny.Array(None, 5)
                nw164_[int(0)] = d_1048_v84_
                nw164_[int(1)] = d_1048_v84_
                nw164_[int(2)] = d_1048_v84_
                nw164_[int(3)] = d_1048_v84_
                nw164_[int(4)] = d_1048_v84_
                d_1049_v85_ = nw164_
                d_1050_v86_: _dafny.Map
                d_1050_v86_ = _dafny.Map({d_1049_v85_: ((d_1043_v79_).f10 if (d_1046_v82_).f12 else (d_1046_v82_).f12)})
                d_1051_v87_: _dafny.MultiSet
                d_1051_v87_ = _dafny.MultiSet([d_1047_v83_])
                d_1050_v86_ = (d_1050_v86_).set(d_1049_v85_, (d_1051_v87_).issubset(d_1051_v87_))
                d_1052_v88_: D5
                d_1052_v88_ = D5_DC13(d_1041_v78_)
                d_1053_v89_: _dafny.Map
                d_1053_v89_ = _dafny.Map({(d_1046_v82_).f12: (d_1052_v88_).cf24})
                d_1054_v90_: _dafny.Map
                d_1054_v90_ = d_1053_v89_
                d_1054_v90_ = d_1054_v90_
            d_1055_v91_: C3
            nw165_ = C3()
            nw165_.ctor__()
            d_1055_v91_ = nw165_
            d_1056_v92_: _dafny.MultiSet
            d_1056_v92_ = _dafny.MultiSet([d_1055_v91_, d_1055_v91_])
            d_1057_v93_: _dafny.MultiSet
            d_1057_v93_ = _dafny.MultiSet([(self).f10, (self).f10, (self).f10])
            d_1058_v94_: _dafny.Map
            d_1058_v94_ = _dafny.Map({(self).f10: d_1057_v93_})
            rhs174_ = d_1056_v92_
            rhs175_ = d_1058_v94_
            d_1056_v92_ = rhs174_
            d_1058_v94_ = rhs175_
            (globalState).f8 = (self).f10
            d_1059_v95_: _dafny.Seq
            d_1060_v96_: bool
            d_1061_v97_: _dafny.Seq
            d_1062_v98_: _dafny.Seq
            out52_: _dafny.Seq
            out53_: bool
            out54_: _dafny.Seq
            out55_: _dafny.Seq
            out52_, out53_, out54_, out55_ = (d_1055_v91_).m0(globalState)
            d_1059_v95_ = out52_
            d_1060_v96_ = out53_
            d_1061_v97_ = out54_
            d_1062_v98_ = out55_
            d_1023_v62_ = len((default__.fm45(d_1060_v96_, d_1023_v62_, d_1060_v96_, globalState)) + ((d_1021_v60_) + (d_1062_v98_)))
        elif True:
            d_1063_v99_: int
            d_1063_v99_ = 295
            (globalState).f8 = ((d_1063_v99_ if (self).f10 else (0) - (d_1063_v99_))) >= ((self).fm1(_dafny.Map({(self).f10: (self).f10}), globalState))
            d_1064_v100_: C3
            nw166_ = C3()
            nw166_.ctor__()
            d_1064_v100_ = nw166_
            d_1065_v101_: _dafny.Seq
            d_1065_v101_ = _dafny.SeqWithoutIsStrInference([d_1064_v100_])
            d_1066_v102_: _dafny.Map
            d_1066_v102_ = _dafny.Map({(self).f10: d_1064_v100_})
            d_1067_v103_: _dafny.Array
            nw167_ = _dafny.Array(None, 15)
            nw167_[int(0)] = d_1064_v100_
            nw167_[int(1)] = d_1064_v100_
            nw167_[int(2)] = (d_1065_v101_)[default__.safeIndex(d_1063_v99_, len(d_1065_v101_))]
            nw167_[int(3)] = d_1064_v100_
            nw167_[int(4)] = d_1064_v100_
            nw167_[int(5)] = d_1064_v100_
            nw167_[int(6)] = d_1064_v100_
            nw167_[int(7)] = (d_1064_v100_ if (self).f10 else d_1064_v100_)
            nw167_[int(8)] = d_1064_v100_
            nw167_[int(9)] = d_1064_v100_
            nw167_[int(10)] = ((d_1066_v102_)[(self).f10] if ((self).f10) in (d_1066_v102_) else d_1064_v100_)
            nw167_[int(11)] = d_1064_v100_
            nw167_[int(12)] = d_1064_v100_
            nw167_[int(13)] = d_1064_v100_
            nw167_[int(14)] = d_1064_v100_
            d_1067_v103_ = nw167_
            index175_ = default__.safeIndex(975, (d_1067_v103_).length(0))
            (d_1067_v103_)[index175_] = d_1064_v100_
            index176_ = default__.safeIndex(975, (d_1067_v103_).length(0))
            (d_1067_v103_)[index176_] = d_1064_v100_
            d_1068_v104_: _dafny.Seq
            d_1068_v104_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "smnhcyhf"))
            (globalState).f8 = (d_1068_v104_) <= (d_1068_v104_)
            d_1069_v105_: C0
            nw168_ = C0()
            nw168_.ctor__()
            d_1069_v105_ = nw168_
            d_1070_v106_: str
            d_1070_v106_ = _dafny.CodePoint('s')
            d_1070_v106_ = d_1070_v106_
        d_1071_v107_: int
        d_1071_v107_ = 521
        d_1071_v107_ = (0) - (d_1071_v107_)

    def m6(self, p0, p1, p2, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: int = int(0)
        r2: _dafny.Seq = _dafny.Seq({})
        r3: int = int(0)
        d_1072_v0_: _dafny.Map
        d_1072_v0_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_1073_i0_ in range(default__.abs(467))]): ((0) - (p2)) + (p0)})
        d_1074_v1_: _dafny.Seq
        d_1074_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "veytmuod"))
        d_1072_v0_ = (d_1072_v0_).set((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kxkg"))) + ((d_1074_v1_).set(default__.safeIndex((self).fm1(_dafny.Map({not((self).f10): (self).f10}), globalState), len(d_1074_v1_)), _dafny.CodePoint('u'))), p2)
        if not(((self).f10 if (self).f10 else (self).f10)):
            (globalState).f8 = (self).f10
            d_1075_v2_: D0
            d_1075_v2_ = D0_DC0(False)
            d_1076_v3_: _dafny.Seq
            d_1076_v3_ = _dafny.SeqWithoutIsStrInference([(self).f10, (self).f10])
            d_1077_v4_: D8
            d_1077_v4_ = D8_DC23(d_1076_v3_)
            d_1078_v5_: _dafny.Map
            d_1078_v5_ = _dafny.Map({d_1075_v2_: d_1077_v4_})
            r3 = len((d_1078_v5_).set((d_1075_v2_ if True else d_1075_v2_), d_1077_v4_))
            d_1079_v6_: str
            d_1079_v6_ = _dafny.CodePoint('s')
            d_1079_v6_ = d_1079_v6_
            d_1080_v7_: _dafny.Map
            d_1080_v7_ = _dafny.Map({not((self).f10): p2})
            d_1080_v7_ = (d_1080_v7_).set((self).f10, len(default__.fm47((self).f10, globalState)))
            d_1081_v8_: D6
            d_1081_v8_ = D6_DC16((self).f9)
            d_1081_v8_ = D6_DC16(default__.fm48((self).f10, globalState))
        elif True:
            d_1082_v9_: C1
            nw169_ = C1()
            nw169_.ctor__((self).f10, d_1074_v1_)
            d_1082_v9_ = nw169_
            d_1083_v10_: _dafny.Seq
            d_1083_v10_ = _dafny.SeqWithoutIsStrInference([not(True), (d_1082_v9_).f12])
            d_1084_v11_: D6
            d_1084_v11_ = D6_DC17((d_1082_v9_).f12, (self).f10, len((d_1082_v9_).f13), len((d_1082_v9_).f13), len(_dafny.SeqWithoutIsStrInference([p2, len(default__.fm45((d_1082_v9_).f12, p2, (d_1082_v9_).f12, globalState)), len(d_1083_v10_), p0, p0])))
            d_1085_v12_: D6
            d_1085_v12_ = D6_DC20(d_1084_v11_)
            source17_ = d_1085_v12_
            if source17_.is_DC17:
                d_1086___mcc_h0_ = source17_.cf29
                d_1087___mcc_h1_ = source17_.cf30
                d_1088___mcc_h2_ = source17_.cf31
                d_1089___mcc_h3_ = source17_.cf32
                d_1090___mcc_h4_ = source17_.cf33
                d_1091_cf33_ = d_1090___mcc_h4_
                d_1092_cf32_ = d_1089___mcc_h3_
                d_1093_cf31_ = d_1088___mcc_h2_
                d_1094_cf30_ = d_1087___mcc_h1_
                d_1095_cf29_ = d_1086___mcc_h0_
                d_1096_v13_: _dafny.Array
                def lambda43_(d_1097_cf29_):
                    def lambda44_(d_1098_i1_):
                        return d_1097_cf29_

                    return lambda44_

                init25_ = lambda43_(d_1095_cf29_)
                nw170_ = _dafny.Array(None, 21)
                for i0_25_ in range(nw170_.length(0)):
                    nw170_[i0_25_] = init25_(i0_25_)
                d_1096_v13_ = nw170_
                d_1099_v14_: _dafny.Map
                d_1099_v14_ = _dafny.Map({not(d_1094_cf30_): d_1096_v13_})
                d_1100_v15_: _dafny.Map
                d_1100_v15_ = _dafny.Map({d_1091_cf33_: d_1095_cf29_})
                d_1099_v14_ = (d_1099_v14_).set(((d_1100_v15_)[(0) - (d_1093_cf31_)] if ((0) - (d_1093_cf31_)) in (d_1100_v15_) else d_1094_cf30_), d_1096_v13_)
                (globalState).f8 = (d_1082_v9_).f12
                d_1101_v16_: _dafny.Map
                d_1101_v16_ = _dafny.Map({len((d_1082_v9_).f13): default__.fm33((d_1082_v9_).f12, d_1095_cf29_, -327, d_1095_cf29_, globalState)})
                d_1102_v17_: _dafny.Seq
                d_1102_v17_ = _dafny.SeqWithoutIsStrInference([d_1101_v16_, d_1101_v16_, _dafny.Map({p0: len(d_1074_v1_)})])
                d_1103_v18_: _dafny.Set
                d_1103_v18_ = _dafny.Set({(d_1082_v9_).f12, False, (self).fm0(d_1083_v10_, d_1095_cf29_, d_1093_cf31_, (d_1082_v9_).f12, globalState), not((d_1082_v9_).f12), True})
                d_1104_v19_: _dafny.Map
                d_1104_v19_ = _dafny.Map({d_1103_v18_: d_1092_cf32_})
                d_1105_v20_: str
                d_1105_v20_ = _dafny.CodePoint('a')
                d_1106_v21_: D4
                d_1106_v21_ = D4_DC12(d_1094_cf30_, len((d_1082_v9_).f13), d_1105_v20_, (0) - (d_1093_cf31_), d_1091_cf33_)
                pat_let_tv23_ = d_1095_cf29_
                pat_let_tv24_ = p2
                def iife89_(_pat_let22_0):
                    def iife90_(d_1107_dt__update__tmp_h0_):
                        def iife91_(_pat_let23_0):
                            def iife92_(d_1108_dt__update_hcf19_h0_):
                                def iife93_(_pat_let24_0):
                                    def iife94_(d_1109_dt__update_hcf23_h0_):
                                        return D4_DC12(d_1108_dt__update_hcf19_h0_, (d_1107_dt__update__tmp_h0_).cf20, (d_1107_dt__update__tmp_h0_).cf21, (d_1107_dt__update__tmp_h0_).cf22, d_1109_dt__update_hcf23_h0_)
                                    return iife94_(_pat_let24_0)
                                return iife93_(pat_let_tv24_)
                            return iife92_(_pat_let23_0)
                        return iife91_(pat_let_tv23_)
                    return iife90_(_pat_let22_0)
                d_1094_cf30_ = ((len((d_1102_v17_)[default__.safeIndex(p0, len(d_1102_v17_))])) * (len(d_1104_v19_))) == ((iife89_(d_1106_v21_)).cf20)
                d_1110_v22_: _dafny.Map
                d_1110_v22_ = _dafny.Map({p1: (default__.fm33((d_1082_v9_).f12, (d_1082_v9_).f12, d_1093_cf31_, (d_1082_v9_).f12, globalState)) <= (d_1092_cf32_)})
                d_1110_v22_ = d_1110_v22_
            elif source17_.is_DC18:
                d_1111___mcc_h5_ = source17_.cf34
                d_1112___mcc_h6_ = source17_.cf35
                d_1113___mcc_h7_ = source17_.cf36
                d_1114_cf36_ = d_1113___mcc_h7_
                d_1115_cf35_ = d_1112___mcc_h6_
                d_1116_cf34_ = d_1111___mcc_h5_
                d_1117_v23_: _dafny.Array
                nw171_ = _dafny.Array(_dafny.Seq({}), 17)
                d_1117_v23_ = nw171_
                nw172_ = _dafny.Array(_dafny.Seq({}), 13)
                d_1117_v23_ = nw172_
                d_1118_v25_: _dafny.Array
                def lambda45_(d_1119_v9_, d_1120_v10_, d_1121_p0_, d_1122_cf36_, d_1123_cf35_):
                    def lambda46_(d_1124_i2_):
                        def iife95_():
                            coll45_ = _dafny.Set()
                            compr_45_: int
                            for compr_45_ in (_dafny.MultiSet([d_1121_p0_, 604, d_1122_cf36_, 916, d_1123_cf35_])).Elements:
                                d_1125_v24_: int = compr_45_
                                if (d_1125_v24_) in (_dafny.MultiSet([d_1121_p0_, 604, d_1122_cf36_, 916, d_1123_cf35_])):
                                    coll45_ = coll45_.union(_dafny.Set([(d_1125_v24_) - (len((D8_DC23(_dafny.SeqWithoutIsStrInference([False]))).cf41))]))
                            return _dafny.Set(coll45_)
                        return (iife95_()
                         if (d_1119_v9_).f12 else _dafny.Set({len(_dafny.Map({_dafny.MultiSet(d_1120_v10_): not((d_1119_v9_).f12)}))}))

                    return lambda46_

                init26_ = lambda45_(d_1082_v9_, d_1083_v10_, p0, d_1114_cf36_, d_1115_cf35_)
                nw173_ = _dafny.Array(None, 2)
                for i0_26_ in range(nw173_.length(0)):
                    nw173_[i0_26_] = init26_(i0_26_)
                d_1118_v25_ = nw173_
                nw174_ = _dafny.Array(_dafny.Set({}), 5)
                d_1118_v25_ = nw174_
                (globalState).f8 = (p0) < ((len(_dafny.SeqWithoutIsStrInference([p0 for d_1126_i3_ in range(default__.abs(331))]))) - (p2))
                d_1074_v1_ = (((default__.fm22(d_1115_cf35_, (d_1082_v9_).f12, (self).f10, p0, globalState)).set(default__.safeIndex(959, len(default__.fm22(d_1115_cf35_, (d_1082_v9_).f12, (self).f10, p0, globalState))), d_1116_cf34_)) + ((d_1082_v9_).f13)) + (_dafny.SeqWithoutIsStrInference([d_1116_cf34_ for d_1127_i4_ in range(default__.abs(429))]))
            elif source17_.is_DC19:
                d_1128___mcc_h8_ = source17_.cf37
                d_1129___mcc_h9_ = source17_.cf38
                d_1130_cf38_ = d_1129___mcc_h9_
                d_1131_cf37_ = d_1128___mcc_h8_
                r1 = p2
                d_1132_v26_: _dafny.Seq
                d_1132_v26_ = _dafny.SeqWithoutIsStrInference([p2])
                (globalState).f8 = ((p0) != ((d_1132_v26_)[default__.safeIndex(len(d_1083_v10_), len(d_1132_v26_))])) or ((d_1082_v9_).fm32((self).fm1(((self).f9).set(d_1131_cf37_, (d_1082_v9_).f12), globalState), d_1130_cf38_, globalState))
                d_1133_v27_: C1
                nw175_ = C1()
                nw175_.ctor__((self).f10, ((d_1082_v9_).f13).set(default__.safeIndex(p0, len((d_1082_v9_).f13)), d_1130_cf38_))
                d_1133_v27_ = nw175_
                (globalState).f8 = (p0) == ((0) - (-486))
            elif source17_.is_DC16:
                d_1134___mcc_h10_ = source17_.cf28
                d_1135_cf28_ = d_1134___mcc_h10_
                d_1136_v28_: D7
                d_1136_v28_ = D7_DC22()
                d_1137_v29_: _dafny.Array
                nw176_ = _dafny.Array(False, 2)
                d_1137_v29_ = nw176_
                d_1138_v30_: _dafny.Set
                d_1138_v30_ = _dafny.Set({(self).f10, (d_1082_v9_).f12})
                d_1139_v31_: _dafny.Map
                d_1139_v31_ = _dafny.Map({p2: (self).f10})
                d_1140_v32_: _dafny.Map
                d_1140_v32_ = _dafny.Map({_dafny.CodePoint('e'): (d_1082_v9_).f12})
                d_1141_v33_: _dafny.Map
                d_1141_v33_ = _dafny.Map({d_1140_v32_: d_1138_v30_})
                d_1142_v34_: _dafny.Seq
                d_1142_v34_ = _dafny.SeqWithoutIsStrInference([p1, p1])
                rhs176_ = d_1136_v28_
                rhs177_ = ((p2) * (p2)) - (p2)
                rhs178_ = d_1137_v29_
                rhs179_ = (default__.fm34(len(d_1139_v31_), (self).f10, (d_1082_v9_).f12, (d_1082_v9_).f12, globalState)) | (((d_1141_v33_)[d_1140_v32_] if (d_1140_v32_) in (d_1141_v33_) else d_1138_v30_))
                rhs180_ = len(d_1142_v34_)
                d_1136_v28_ = rhs176_
                r1 = rhs177_
                d_1137_v29_ = rhs178_
                d_1138_v30_ = rhs179_
                r3 = rhs180_
                d_1143_v35_: _dafny.Seq
                d_1143_v35_ = _dafny.SeqWithoutIsStrInference([p2])
                d_1144_v36_: D1
                d_1144_v36_ = D1_DC2(d_1143_v35_)
                d_1145_v37_: _dafny.Set
                d_1145_v37_ = _dafny.Set({d_1144_v36_})
                d_1146_v38_: _dafny.Map
                d_1146_v38_ = _dafny.Map({((d_1082_v9_).f13) <= ((d_1082_v9_).f13): d_1145_v37_})
                d_1147_v39_: str
                d_1147_v39_ = _dafny.CodePoint('q')
                d_1148_v40_: _dafny.Map
                d_1148_v40_ = _dafny.Map({d_1147_v39_: d_1147_v39_})
                rhs181_ = (d_1147_v39_) not in (d_1148_v40_)
                rhs182_ = (d_1146_v38_) | (d_1146_v38_)
                rhs183_ = (d_1138_v30_).issubset(d_1138_v30_)
                lhs129_ = globalState
                lhs130_ = globalState
                lhs129_.f8 = rhs181_
                d_1146_v38_ = rhs182_
                lhs130_.f8 = rhs183_
                d_1149_v41_: _dafny.Map
                d_1149_v41_ = _dafny.Map({((0) - (len(d_1139_v31_))) != (p0): (0) - ((p2) + (p0))})
                d_1149_v41_ = d_1149_v41_
                index177_ = default__.safeIndex(129, (d_1137_v29_).length(0))
                (d_1137_v29_)[index177_] = (d_1082_v9_).f12
                index178_ = default__.safeIndex(262, (p1).length(0))
                (p1)[index178_] = p0
                index179_ = default__.safeIndex(129, (d_1137_v29_).length(0))
                index180_ = default__.safeIndex(262, (p1).length(0))
                rhs184_ = (p2) < (p0)
                rhs185_ = (0) - ((0) - ((self).fm1(d_1135_cf28_, globalState)))
                rhs186_ = p2
                rhs187_ = p0
                rhs188_ = (d_1082_v9_).f12
                lhs131_ = d_1137_v29_
                lhs132_ = default__.safeIndex(129, (d_1137_v29_).length(0))
                lhs133_ = p1
                lhs134_ = default__.safeIndex(262, (p1).length(0))
                lhs135_ = globalState
                lhs131_[lhs132_] = rhs184_
                r3 = rhs185_
                r3 = rhs186_
                lhs133_[lhs134_] = rhs187_
                lhs135_.f8 = rhs188_
            elif True:
                d_1150___mcc_h11_ = source17_.cf39
                d_1151_cf39_ = d_1150___mcc_h11_
                d_1152_v42_: _dafny.Set
                d_1152_v42_ = _dafny.Set({(d_1082_v9_).f12, (self).f10})
                d_1153_v43_: D3
                d_1153_v43_ = D3_DC8((d_1152_v42_).issubset(d_1152_v42_), (self).f10)
                d_1153_v43_ = default__.fm35(False, globalState)
                d_1154_v44_: _dafny.Map
                d_1154_v44_ = _dafny.Map({(d_1082_v9_).f12: len((self).f9)})
                d_1155_v45_: _dafny.Set
                d_1155_v45_ = _dafny.Set({len(d_1154_v44_)})
                d_1156_v46_: D4
                d_1156_v46_ = D4_DC11(len(d_1083_v10_), d_1155_v45_, not((d_1082_v9_).f12))
                d_1157_v47_: _dafny.Map
                d_1157_v47_ = _dafny.Map({309: p2})
                d_1158_v48_: _dafny.Map
                d_1158_v48_ = _dafny.Map({p0: len(d_1157_v47_)})
                pat_let_tv25_ = d_1155_v45_
                d_1159_v50_: _dafny.Seq
                def iife97_():
                    coll46_ = _dafny.Set()
                    compr_46_: int
                    for compr_46_ in (_dafny.SeqWithoutIsStrInference([p0 for d_1160_i5_ in range(default__.abs(-534))])).Elements:
                        d_1161_v49_: int = compr_46_
                        if (d_1161_v49_) in (_dafny.SeqWithoutIsStrInference([p0 for d_1160_i5_ in range(default__.abs(-534))])):
                            coll46_ = coll46_.union(_dafny.Set([default__.safeModuloInt(d_1161_v49_, -700)]))
                    return _dafny.Set(coll46_)
                def iife96_(_pat_let25_0):
                    def iife98_(d_1162_dt__update__tmp_h1_):
                        def iife99_(_pat_let26_0):
                            def iife100_(d_1163_dt__update_hcf17_h0_):
                                return D4_DC11((d_1162_dt__update__tmp_h1_).cf16, d_1163_dt__update_hcf17_h0_, (d_1162_dt__update__tmp_h1_).cf18)
                            return iife100_(_pat_let26_0)
                        return iife99_(pat_let_tv25_)
                    return iife98_(_pat_let25_0)
                d_1159_v50_ = _dafny.SeqWithoutIsStrInference([d_1156_v46_, iife96_(D4_DC11(len(d_1157_v47_), iife97_()
, (d_1082_v9_).f12)), d_1156_v46_])
                d_1164_v51_: _dafny.Seq
                d_1164_v51_ = d_1159_v50_
                d_1165_v52_: _dafny.Map
                d_1165_v52_ = _dafny.Map({d_1159_v50_: -803})
                rhs189_ = ((d_1164_v51_)) not in ((D14_DC34(d_1165_v52_)).cf61)
                rhs190_ = p2
                lhs136_ = globalState
                lhs136_.f8 = rhs189_
                r1 = rhs190_
                d_1074_v1_ = ((d_1082_v9_).f13) + ((d_1074_v1_) + ((d_1082_v9_).f13))
                d_1166_v53_: _dafny.Map
                d_1166_v53_ = _dafny.Map({d_1164_v51_: d_1155_v45_})
                d_1154_v44_ = (d_1154_v44_).set((p0) <= (len(d_1166_v53_)), default__.safeModuloInt(len(d_1074_v1_), p0))
            d_1167_v54_: C3
            nw177_ = C3()
            nw177_.ctor__()
            d_1167_v54_ = nw177_
            d_1168_v55_: _dafny.Array
            def lambda47_(d_1169_v10_, d_1170_p2_):
                def lambda48_(d_1171_i6_):
                    return (len(d_1169_v10_)) < (d_1170_p2_)

                return lambda48_

            init27_ = lambda47_(d_1083_v10_, p2)
            nw178_ = _dafny.Array(None, 27)
            for i0_27_ in range(nw178_.length(0)):
                nw178_[i0_27_] = init27_(i0_27_)
            d_1168_v55_ = nw178_
            index181_ = default__.safeIndex(107, (d_1168_v55_).length(0))
            (d_1168_v55_)[index181_] = (d_1082_v9_).f12
            index182_ = default__.safeIndex(107, (d_1168_v55_).length(0))
            (d_1168_v55_)[index182_] = (d_1083_v10_)[default__.safeIndex(p0, len(d_1083_v10_))]
            index183_ = default__.safeIndex(197, (p1).length(0))
            (p1)[index183_] = (207) + (p0)
            index184_ = default__.safeIndex(197, (p1).length(0))
            (p1)[index184_] = p2
        r1 = len((default__.fm45(True, p2, True, globalState)) + ((d_1074_v1_) + (d_1074_v1_)))
        d_1172_v56_: _dafny.Array
        nw179_ = _dafny.Array(False, 14)
        d_1172_v56_ = nw179_
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_1172_v56_).length(0)):
            d_1173_i7_: int = guard_loop_3_
            if (True) and (((0) <= (d_1173_i7_)) and ((d_1173_i7_) < ((d_1172_v56_).length(0)))):
                (d_1172_v56_)[(d_1173_i7_)] = not((self).f10)
        d_1174_v57_: _dafny.Map
        d_1174_v57_ = _dafny.Map({p0: (self).f10})
        d_1175_v58_: _dafny.Seq
        d_1175_v58_ = _dafny.SeqWithoutIsStrInference([(self).f10])
        (globalState).f8 = ((d_1174_v57_)[p0] if (p0) in (d_1174_v57_) else ((self).f10) in (d_1175_v58_))
        d_1176_v59_: C2
        nw180_ = C2()
        nw180_.ctor__((self).f10, d_1172_v56_, (_dafny.Map({(self).f10: (self).f10})) | (_dafny.Map({True: not(not((self).f10))})), (self).f10)
        d_1176_v59_ = nw180_
        r0 = _dafny.SeqWithoutIsStrInference([831])
        d_1177_v60_: _dafny.Seq
        d_1177_v60_ = _dafny.SeqWithoutIsStrInference([p0, p0])
        d_1178_v61_: _dafny.Map
        d_1178_v61_ = _dafny.Map({p0: p2})
        d_1179_v62_: _dafny.Map
        d_1179_v62_ = _dafny.Map({(self).f10: (d_1177_v60_) + ((_dafny.SeqWithoutIsStrInference([(0) - (p0), ((d_1178_v61_)[len(d_1177_v60_)] if (len(d_1177_v60_)) in (d_1178_v61_) else p2), p2, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vyqvooyef")))])).set(default__.safeIndex((d_1176_v59_).fm42(p0, p2, p2, globalState), len(_dafny.SeqWithoutIsStrInference([(0) - (p0), ((d_1178_v61_)[len(d_1177_v60_)] if (len(d_1177_v60_)) in (d_1178_v61_) else p2), p2, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vyqvooyef")))]))), p0))})
        r1 = len(d_1179_v62_)
        r2 = default__.fm37(globalState)
        r3 = p2
        return r0, r1, r2, r3


class C5(T1, T0):
    def  __init__(self):
        self._f9: _dafny.Map = _dafny.Map({})
        self._f10: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    @property
    def f9(self):
        return self._f9
    @property
    def f10(self):
        return self._f10
    def ctor__(self, f9, f10):
        (self)._f9 = f9
        (self)._f10 = f10

    def fm0(self, p0, p1, p2, p3, globalState):
        return not((D0_DC1(len(_dafny.Set({(_dafny.MultiSet([(self).f10])).cardinality})), (self).f10, -312)).cf2)

    def fm1(self, p0, globalState):
        return (0) - (default__.safeDivisionInt(138, len(_dafny.SeqWithoutIsStrInference([(self).f10, (self).f10]))))

    def fm18(self, p0, p1, globalState):
        return ((0) - ((-298) - (len(_dafny.Map({133: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n"))}))))) < ((0) - (default__.safeModuloInt(len(_dafny.Map({-447: (self).f10})), 686)))

    def fm19(self, p0, p1, globalState):
        return len(((_dafny.Map({_dafny.SeqWithoutIsStrInference([D1_DC5(D1_DC2(_dafny.SeqWithoutIsStrInference([725]))), D1_DC5(D1_DC4((self).f10, (self).f10, 103))]): _dafny.Map({(self).f10: D1_DC3(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pppcgeiju")))})})) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([D1_DC5(D1_DC3(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wjqjefffv")))) for d_1180_i0_ in range(default__.abs(322))]): _dafny.Map({(self).f10: D1_DC3(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ojakub")))})})) if (self).f10 else (_dafny.Map({_dafny.SeqWithoutIsStrInference([D1_DC5(D1_DC3(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vlvw"))))]): _dafny.Map({(self).f10: D1_DC3(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i")))})})) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([D1_DC5(D1_DC3(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oa")))), D1_DC5(D1_DC3(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lualcqxw"))))]): _dafny.Map({(self).f10: D1_DC3(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jlo")))})}))))

    def m0(self, globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r1: bool = False
        r2: _dafny.Seq = _dafny.Seq({})
        r3: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        d_1181_v0_: _dafny.Seq
        d_1181_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "avx"))
        d_1182_v1_: int
        d_1182_v1_ = 649
        hi2_ = default__.safeModuloInt(d_1182_v1_, d_1182_v1_)
        for d_1183_i0_ in range(len(d_1181_v0_), hi2_):
            d_1184_v2_: _dafny.Array
            nw181_ = _dafny.Array(int(0), 23)
            d_1184_v2_ = nw181_
            index185_ = default__.safeIndex(880, (d_1184_v2_).length(0))
            (d_1184_v2_)[index185_] = d_1183_i0_
            d_1185_v3_: _dafny.Seq
            d_1185_v3_ = _dafny.SeqWithoutIsStrInference([d_1181_v0_, default__.fm20(d_1183_i0_, (self).f10, (self).f10, (self).fm19(d_1182_v1_, d_1183_i0_, globalState), globalState)])
            index186_ = default__.safeIndex(880, (d_1184_v2_).length(0))
            (d_1184_v2_)[index186_] = len((d_1185_v3_)[default__.safeIndex(d_1183_i0_, len(d_1185_v3_))])
            d_1182_v1_ = (0) - (default__.safeDivisionInt((524) - ((d_1184_v2_)[default__.safeIndex(880, (d_1184_v2_).length(0))]), d_1182_v1_))
            d_1186_v4_: _dafny.Array
            d_1187_v5_: _dafny.Seq
            d_1188_v6_: _dafny.Set
            out56_: _dafny.Array
            out57_: _dafny.Seq
            out58_: _dafny.Set
            out56_, out57_, out58_ = (self).m4((self).f10, (self).f10, globalState)
            d_1186_v4_ = out56_
            d_1187_v5_ = out57_
            d_1188_v6_ = out58_
            if (self).f10:
                index187_ = default__.safeIndex(230, (d_1186_v4_).length(0))
                (d_1186_v4_)[index187_] = True
                d_1189_v7_: _dafny.Set
                d_1189_v7_ = _dafny.Set({(self).f10})
                index188_ = default__.safeIndex(230, (d_1186_v4_).length(0))
                (d_1186_v4_)[index188_] = (d_1189_v7_).issubset(d_1189_v7_)
                d_1190_v8_: _dafny.Map
                d_1190_v8_ = _dafny.Map({d_1183_i0_: _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([d_1182_v1_])), d_1182_v1_, d_1183_i0_, d_1182_v1_])})
                d_1191_v9_: _dafny.Seq
                d_1191_v9_ = _dafny.SeqWithoutIsStrInference([d_1183_i0_])
                d_1190_v8_ = (d_1190_v8_).set(d_1183_i0_, d_1191_v9_)
                d_1181_v0_ = d_1181_v0_
                r0 = d_1181_v0_
                r1 = True
            elif True:
                index189_ = default__.safeIndex(331, (d_1186_v4_).length(0))
                (d_1186_v4_)[index189_] = (self).f10
                index190_ = default__.safeIndex(331, (d_1186_v4_).length(0))
                (d_1186_v4_)[index190_] = (self).f10
                d_1192_v10_: _dafny.Map
                d_1192_v10_ = _dafny.Map({default__.fm21(globalState): (d_1183_i0_) + (len(d_1181_v0_))})
                d_1193_v11_: _dafny.Seq
                d_1193_v11_ = _dafny.SeqWithoutIsStrInference([284, len(d_1181_v0_)])
                d_1194_v12_: D1
                d_1194_v12_ = D1_DC2(d_1193_v11_)
                d_1192_v10_ = (d_1192_v10_).set(d_1194_v12_, d_1182_v1_)
                index191_ = default__.safeIndex(880, (d_1184_v2_).length(0))
                (d_1184_v2_)[index191_] = (d_1193_v11_)[default__.safeIndex((d_1184_v2_)[default__.safeIndex(880, (d_1184_v2_).length(0))], len(d_1193_v11_))]
                r0 = (d_1181_v0_) + (d_1181_v0_)
                (globalState).f8 = (d_1186_v4_)[default__.safeIndex(331, (d_1186_v4_).length(0))]
        d_1182_v1_ = d_1182_v1_
        d_1195_v13_: _dafny.Array
        def lambda49_(d_1196_i1_):
            return ((self).f10) not in (_dafny.SeqWithoutIsStrInference([(self).f10, (self).f10, (self).f10]))

        init28_ = lambda49_
        nw182_ = _dafny.Array(None, 24)
        for i0_28_ in range(nw182_.length(0)):
            nw182_[i0_28_] = init28_(i0_28_)
        d_1195_v13_ = nw182_
        index192_ = default__.safeIndex(795, (d_1195_v13_).length(0))
        (d_1195_v13_)[index192_] = (((self).f9)[(self).f10] if ((self).f10) in ((self).f9) else True)
        index193_ = default__.safeIndex(795, (d_1195_v13_).length(0))
        (d_1195_v13_)[index193_] = True
        d_1197_v14_: _dafny.Set
        d_1197_v14_ = _dafny.Set({d_1181_v0_, d_1181_v0_, d_1181_v0_})
        d_1198_v15_: _dafny.Map
        d_1198_v15_ = _dafny.Map({len(d_1197_v14_): (0) - (len(d_1181_v0_))})
        d_1199_v16_: T1
        nw183_ = C2()
        nw183_.ctor__((self).f10, d_1195_v13_, _dafny.Map({(self).f10: (d_1195_v13_)[default__.safeIndex(795, (d_1195_v13_).length(0))]}), (d_1195_v13_)[default__.safeIndex(795, (d_1195_v13_).length(0))])
        d_1199_v16_ = nw183_
        d_1200_v17_: _dafny.Map
        d_1200_v17_ = _dafny.Map({d_1198_v15_: d_1199_v16_})
        d_1201_v19_: _dafny.Map
        d_1201_v19_ = _dafny.Map({d_1182_v1_: d_1199_v16_})
        def iife101_():
            coll47_ = _dafny.Map()
            compr_47_: int
            for compr_47_ in _dafny.IntegerRange(218, -929):
                d_1202_v18_: int = compr_47_
                if ((218) <= (d_1202_v18_)) and ((d_1202_v18_) < (-929)):
                    coll47_[(d_1202_v18_) - (413)] = d_1182_v1_
            return _dafny.Map(coll47_)
        d_1200_v17_ = (d_1200_v17_).set(iife101_()
        , ((d_1201_v19_)[d_1182_v1_] if (d_1182_v1_) in (d_1201_v19_) else d_1199_v16_))
        d_1203_v20_: C1
        nw184_ = C1()
        nw184_.ctor__((d_1195_v13_)[default__.safeIndex(795, (d_1195_v13_).length(0))], d_1181_v0_)
        d_1203_v20_ = nw184_
        index194_ = default__.safeIndex(795, (d_1195_v13_).length(0))
        rhs191_ = not((d_1195_v13_)[default__.safeIndex(795, (d_1195_v13_).length(0))])
        rhs192_ = (d_1182_v1_ if (d_1195_v13_)[default__.safeIndex(795, (d_1195_v13_).length(0))] else d_1182_v1_)
        rhs193_ = (0) - (d_1182_v1_)
        rhs194_ = d_1203_v20_
        rhs195_ = (d_1195_v13_)[default__.safeIndex(795, (d_1195_v13_).length(0))]
        lhs137_ = d_1195_v13_
        lhs138_ = default__.safeIndex(795, (d_1195_v13_).length(0))
        lhs139_ = globalState
        lhs137_[lhs138_] = rhs191_
        d_1182_v1_ = rhs192_
        d_1182_v1_ = rhs193_
        d_1203_v20_ = rhs194_
        lhs139_.f8 = rhs195_
        d_1204_i2_: int
        d_1204_i2_ = 0
        with _dafny.label("5"):
            while (d_1203_v20_).f12:
                with _dafny.c_label("5"):
                    if (d_1204_i2_) >= (100):
                        raise _dafny.Break("5")
                    d_1204_i2_ = (d_1204_i2_) + (1)
                    d_1205_v21_: _dafny.MultiSet
                    d_1205_v21_ = _dafny.MultiSet([d_1181_v0_, d_1181_v0_])
                    d_1206_v22_: D11
                    d_1206_v22_ = D11_DC27(d_1205_v21_)
                    d_1207_v23_: _dafny.Map
                    d_1207_v23_ = _dafny.Map({(self).f10: d_1206_v22_})
                    d_1207_v23_ = d_1207_v23_
                    d_1208_v24_: _dafny.Map
                    d_1208_v24_ = _dafny.Map({(d_1203_v20_).f12: (d_1203_v20_).f12})
                    d_1208_v24_ = ((d_1199_v16_).f9 if (d_1203_v20_).f12 else (d_1199_v16_).f9)
                    if (d_1203_v20_).f12:
                        index195_ = default__.safeIndex(795, (d_1195_v13_).length(0))
                        (d_1195_v13_)[index195_] = ((0) - (d_1182_v1_)) < ((d_1182_v1_) + (d_1182_v1_))
                        d_1209_v25_: _dafny.Seq
                        d_1209_v25_ = _dafny.SeqWithoutIsStrInference([d_1195_v13_, d_1195_v13_, d_1195_v13_, d_1195_v13_])
                        d_1209_v25_ = d_1209_v25_
                        (globalState).f8 = False
                        d_1210_v26_: C4
                        nw185_ = C4()
                        nw185_.ctor__((d_1199_v16_).f9, (d_1199_v16_).f10)
                        d_1210_v26_ = nw185_
                        d_1211_v27_: _dafny.Seq
                        d_1211_v27_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_1212_i3_ in range(default__.abs(113))]))])
                        d_1213_v28_: _dafny.Set
                        d_1213_v28_ = _dafny.Set({len(d_1211_v27_)})
                        d_1214_v29_: D6
                        d_1214_v29_ = D6_DC17((d_1199_v16_).f10, True, d_1182_v1_, d_1182_v1_, len(d_1198_v15_))
                        d_1215_v30_: D4
                        d_1215_v30_ = D4_DC11(len(_dafny.Map({d_1210_v26_: d_1182_v1_})), d_1213_v28_, not((d_1214_v29_).cf30))
                        d_1216_v31_: _dafny.Seq
                        d_1216_v31_ = _dafny.SeqWithoutIsStrInference([d_1215_v30_, d_1215_v30_, D4_DC11(len((d_1199_v16_).f9), d_1213_v28_, (d_1195_v13_)[default__.safeIndex(795, (d_1195_v13_).length(0))])])
                        d_1217_v32_: _dafny.Seq
                        d_1217_v32_ = d_1216_v31_
                        d_1217_v32_ = d_1217_v32_
                        d_1218_v33_: _dafny.Map
                        d_1218_v33_ = _dafny.Map({d_1182_v1_: d_1195_v13_})
                        d_1218_v33_ = (d_1218_v33_).set(d_1182_v1_, ((d_1218_v33_)[d_1182_v1_] if (d_1182_v1_) in (d_1218_v33_) else d_1195_v13_))
                    elif True:
                        d_1219_v34_: _dafny.Array
                        nw186_ = _dafny.Array(D1.default()(), 27)
                        d_1219_v34_ = nw186_
                        d_1220_v35_: D1
                        d_1220_v35_ = D1_DC4((d_1195_v13_)[default__.safeIndex(795, (d_1195_v13_).length(0))], (d_1199_v16_).f10, d_1182_v1_)
                        d_1221_v36_: _dafny.Map
                        d_1221_v36_ = _dafny.Map({d_1219_v34_: (0) - ((d_1220_v35_).cf8)})
                        d_1222_v37_: D15
                        d_1222_v37_ = D15_DC37(d_1219_v34_)
                        d_1221_v36_ = (d_1221_v36_).set((d_1222_v37_).cf68, 207)
                        d_1223_v38_: str
                        d_1223_v38_ = _dafny.CodePoint('l')
                        d_1224_v39_: _dafny.Map
                        d_1225_v40_: bool
                        out59_: _dafny.Map
                        out60_: bool
                        out59_, out60_ = (d_1199_v16_).m1(d_1223_v38_, globalState)
                        d_1224_v39_ = out59_
                        d_1225_v40_ = out60_
                        d_1226_v41_: _dafny.Seq
                        d_1226_v41_ = _dafny.SeqWithoutIsStrInference([d_1225_v40_])
                        d_1227_v42_: _dafny.MultiSet
                        d_1227_v42_ = _dafny.MultiSet([d_1182_v1_])
                        d_1228_v43_: _dafny.Seq
                        d_1228_v43_ = _dafny.SeqWithoutIsStrInference([d_1227_v42_])
                        d_1229_v44_: _dafny.Map
                        d_1229_v44_ = _dafny.Map({(d_1195_v13_)[default__.safeIndex(795, (d_1195_v13_).length(0))]: d_1199_v16_})
                        d_1230_v45_: _dafny.Array
                        nw187_ = _dafny.Array(None, 23)
                        nw187_[int(0)] = d_1182_v1_
                        nw187_[int(1)] = d_1182_v1_
                        nw187_[int(2)] = d_1182_v1_
                        nw187_[int(3)] = (0) - ((self).fm19(d_1182_v1_, d_1182_v1_, globalState))
                        nw187_[int(4)] = d_1182_v1_
                        nw187_[int(5)] = d_1182_v1_
                        nw187_[int(6)] = d_1182_v1_
                        nw187_[int(7)] = d_1182_v1_
                        nw187_[int(8)] = d_1182_v1_
                        nw187_[int(9)] = d_1182_v1_
                        nw187_[int(10)] = ((d_1228_v43_)[default__.safeIndex(d_1182_v1_, len(d_1228_v43_))]).cardinality
                        nw187_[int(11)] = -273
                        nw187_[int(12)] = d_1182_v1_
                        nw187_[int(13)] = len(d_1229_v44_)
                        nw187_[int(14)] = d_1182_v1_
                        nw187_[int(15)] = d_1182_v1_
                        nw187_[int(16)] = d_1182_v1_
                        nw187_[int(17)] = -283
                        nw187_[int(18)] = d_1182_v1_
                        nw187_[int(19)] = len(d_1198_v15_)
                        nw187_[int(20)] = d_1182_v1_
                        nw187_[int(21)] = d_1182_v1_
                        nw187_[int(22)] = 485
                        d_1230_v45_ = nw187_
                        d_1231_v46_: _dafny.Map
                        d_1231_v46_ = _dafny.Map({d_1230_v45_: (d_1199_v16_).f10})
                        d_1232_v47_: _dafny.Map
                        d_1232_v47_ = _dafny.Map({d_1223_v38_: d_1231_v46_})
                        (globalState).f8 = (d_1199_v16_).fm0((d_1226_v41_) + (d_1226_v41_), (default__.fm49(not((d_1199_v16_).f10), globalState)).cf29, len(((d_1232_v47_)[d_1223_v38_] if (d_1223_v38_) in (d_1232_v47_) else d_1231_v46_)), (d_1199_v16_).fm0(d_1226_v41_, (d_1199_v16_).f10, d_1182_v1_, (self).f10, globalState), globalState)
                        d_1233_v48_: _dafny.MultiSet
                        d_1233_v48_ = _dafny.MultiSet([(d_1199_v16_).f10])
                        d_1234_v49_: _dafny.Map
                        d_1234_v49_ = _dafny.Map({d_1233_v48_: d_1195_v13_})
                        (globalState).f8 = (len(d_1234_v49_)) != (d_1182_v1_)
                        (globalState).f8 = False
                    d_1235_v50_: _dafny.Array
                    nw188_ = _dafny.Array(int(0), 5)
                    d_1235_v50_ = nw188_
                    d_1236_v51_: _dafny.Seq
                    d_1236_v51_ = _dafny.SeqWithoutIsStrInference([d_1235_v50_, d_1235_v50_])
                    d_1237_v52_: D12
                    d_1237_v52_ = D12_DC32(d_1235_v50_, d_1236_v51_)
                    d_1238_v53_: _dafny.Map
                    d_1238_v53_ = _dafny.Map({(0) - (d_1182_v1_): d_1237_v52_})
                    d_1239_v54_: _dafny.Map
                    d_1239_v54_ = _dafny.Map({d_1238_v53_: d_1182_v1_})
                    d_1198_v15_ = (d_1198_v15_).set(len((d_1239_v54_) | (_dafny.Map({d_1238_v53_: d_1182_v1_}))), len(d_1181_v0_))
                    pass
            pass
        r0 = d_1181_v0_
        r1 = (d_1199_v16_).f10
        d_1240_v55_: _dafny.Set
        d_1240_v55_ = _dafny.Set({(self).f10, (d_1203_v20_).f12, (D0_DC0((d_1199_v16_).f10)).cf0, not((self).f10)})
        d_1241_v56_: _dafny.Seq
        d_1241_v56_ = _dafny.SeqWithoutIsStrInference([d_1240_v55_])
        d_1242_v57_: _dafny.Seq
        d_1242_v57_ = _dafny.SeqWithoutIsStrInference([(self).f10])
        r2 = ((d_1241_v56_) + (_dafny.SeqWithoutIsStrInference([d_1240_v55_ for d_1243_i4_ in range(default__.abs(956))]))) + ((d_1241_v56_).set(default__.safeIndex(28, len(d_1241_v56_)), _dafny.Set({(d_1199_v16_).fm0(d_1242_v57_, (d_1199_v16_).f10, d_1182_v1_, (d_1195_v13_)[default__.safeIndex(795, (d_1195_v13_).length(0))], globalState)})))
        r3 = default__.fm22(d_1182_v1_, (self).f10, (self).f10, len((d_1203_v20_).f13), globalState)
        return r0, r1, r2, r3

    def m1(self, p0, globalState):
        r0: _dafny.Map = _dafny.Map({})
        r1: bool = False
        d_1244_v0_: int
        d_1244_v0_ = 579
        d_1245_i0_: int
        d_1245_i0_ = 0
        with _dafny.label("6"):
            while (d_1244_v0_) > (d_1244_v0_):
                with _dafny.c_label("6"):
                    if (d_1245_i0_) >= (100):
                        raise _dafny.Break("6")
                    d_1245_i0_ = (d_1245_i0_) + (1)
                    d_1246_v1_: _dafny.Seq
                    d_1246_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nllkwrs"))
                    rhs196_ = (self).f10
                    rhs197_ = 994
                    rhs198_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b")) if (self).f10 else (d_1246_v1_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aogmc"))))
                    lhs140_ = globalState
                    lhs140_.f8 = rhs196_
                    d_1244_v0_ = rhs197_
                    d_1246_v1_ = rhs198_
                    d_1247_v2_: _dafny.Array
                    def lambda50_(d_1248_p0_):
                        def lambda51_(d_1249_i1_):
                            return d_1248_p0_

                        return lambda51_

                    init29_ = lambda50_(p0)
                    nw189_ = _dafny.Array(None, 6)
                    for i0_29_ in range(nw189_.length(0)):
                        nw189_[i0_29_] = init29_(i0_29_)
                    d_1247_v2_ = nw189_
                    index196_ = default__.safeIndex(475, (d_1247_v2_).length(0))
                    (d_1247_v2_)[index196_] = p0
                    index197_ = default__.safeIndex(475, (d_1247_v2_).length(0))
                    (d_1247_v2_)[index197_] = p0
                    d_1250_v3_: _dafny.Array
                    def lambda52_(d_1251_i2_):
                        return (((self).f9)[(self).f10] if ((self).f10) in ((self).f9) else (self).f10)

                    init30_ = lambda52_
                    nw190_ = _dafny.Array(None, 5)
                    for i0_30_ in range(nw190_.length(0)):
                        nw190_[i0_30_] = init30_(i0_30_)
                    d_1250_v3_ = nw190_
                    index198_ = default__.safeIndex(663, (d_1250_v3_).length(0))
                    (d_1250_v3_)[index198_] = (self).f10
                    index199_ = default__.safeIndex(663, (d_1250_v3_).length(0))
                    (d_1250_v3_)[index199_] = (self).f10
                    d_1252_v4_: _dafny.Map
                    d_1252_v4_ = _dafny.Map({-414: (self).f10})
                    d_1252_v4_ = (d_1252_v4_) | (default__.fm44((self).f10, globalState))
                    pass
            pass
        d_1253_v5_: _dafny.MultiSet
        d_1253_v5_ = _dafny.MultiSet([(self).f10])
        d_1244_v0_ = (0) - ((((d_1253_v5_).set(not(True), default__.abs(d_1244_v0_))) - (_dafny.MultiSet([(self).f10]))).cardinality)
        (globalState).f8 = (self).f10
        d_1254_v6_: _dafny.Seq
        d_1254_v6_ = _dafny.SeqWithoutIsStrInference([d_1244_v0_, d_1244_v0_])
        if ((d_1254_v6_) + (_dafny.SeqWithoutIsStrInference([d_1244_v0_]))) <= ((d_1254_v6_) + (d_1254_v6_)):
            d_1244_v0_ = len(((d_1254_v6_) + (d_1254_v6_)) + (d_1254_v6_))
            d_1255_v8_: _dafny.Seq
            d_1255_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "edykaowmo"))
            d_1256_v9_: _dafny.Array
            nw191_ = _dafny.Array(None, 15)
            nw191_[int(0)] = d_1244_v0_
            nw191_[int(1)] = d_1244_v0_
            nw191_[int(2)] = d_1244_v0_
            nw191_[int(3)] = d_1244_v0_
            nw191_[int(4)] = d_1244_v0_
            nw191_[int(5)] = len(_dafny.Map({d_1244_v0_: (self).f10}))
            nw191_[int(6)] = d_1244_v0_
            nw191_[int(7)] = -569
            nw191_[int(8)] = d_1244_v0_
            def iife102_():
                coll48_ = _dafny.Set()
                compr_48_: int
                for compr_48_ in _dafny.IntegerRange(98, 915):
                    d_1257_v7_: int = compr_48_
                    if ((98) <= (d_1257_v7_)) and ((d_1257_v7_) < (915)):
                        coll48_ = coll48_.union(_dafny.Set([default__.safeModuloInt(d_1257_v7_, len(_dafny.SeqWithoutIsStrInference([(self).f10])))]))
                return _dafny.Set(coll48_)
            nw191_[int(9)] = len(iife102_()
            )
            nw191_[int(10)] = d_1244_v0_
            nw191_[int(11)] = len(d_1255_v8_)
            nw191_[int(12)] = d_1244_v0_
            nw191_[int(13)] = d_1244_v0_
            nw191_[int(14)] = d_1244_v0_
            d_1256_v9_ = nw191_
            d_1258_v10_: _dafny.MultiSet
            d_1258_v10_ = _dafny.MultiSet([d_1256_v9_])
            d_1259_v11_: _dafny.Map
            d_1259_v11_ = _dafny.Map({(d_1258_v10_) - (d_1258_v10_): default__.fm33((self).f10, (self).f10, (0) - (d_1244_v0_), not((self).f10), globalState)})
            d_1259_v11_ = (d_1259_v11_).set(d_1258_v10_, d_1244_v0_)
            d_1260_v12_: _dafny.Array
            nw192_ = _dafny.Array(None, 4)
            nw192_[int(0)] = (self).f10
            nw192_[int(1)] = (self).f10
            nw192_[int(2)] = (self).f10
            nw192_[int(3)] = (self).f10
            d_1260_v12_ = nw192_
            d_1261_v13_: C2
            nw193_ = C2()
            nw193_.ctor__((d_1244_v0_) >= (d_1244_v0_), d_1260_v12_, (self).f9, (d_1244_v0_) > (d_1244_v0_))
            d_1261_v13_ = nw193_
            d_1262_v14_: _dafny.Set
            d_1262_v14_ = _dafny.Set({(self).f10})
            (d_1261_v13_).f14 = not ((d_1262_v14_).issubset(d_1262_v14_)) or ((d_1244_v0_) <= (640))
            d_1256_v9_ = d_1256_v9_
        elif True:
            d_1244_v0_ = d_1244_v0_
            d_1263_v15_: _dafny.Array
            def lambda53_(d_1264_i3_):
                return (d_1264_i3_) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uneixc"))))

            init31_ = lambda53_
            nw194_ = _dafny.Array(None, 4)
            for i0_31_ in range(nw194_.length(0)):
                nw194_[i0_31_] = init31_(i0_31_)
            d_1263_v15_ = nw194_
            d_1265_v16_: _dafny.Map
            d_1265_v16_ = _dafny.Map({(self).f10: d_1263_v15_})
            d_1265_v16_ = (d_1265_v16_).set((self).f10, d_1263_v15_)
            index200_ = default__.safeIndex(34, (d_1263_v15_).length(0))
            (d_1263_v15_)[index200_] = d_1244_v0_
            index201_ = default__.safeIndex(34, (d_1263_v15_).length(0))
            (d_1263_v15_)[index201_] = d_1244_v0_
            d_1266_v17_: C0
            nw195_ = C0()
            nw195_.ctor__()
            d_1266_v17_ = nw195_
            r1 = (((d_1263_v15_)[default__.safeIndex(34, (d_1263_v15_).length(0))]) * (d_1244_v0_)) == (d_1244_v0_)
        d_1267_v18_: _dafny.Array
        def lambda54_(d_1268_i5_):
            return default__.safeDivisionInt(d_1268_i5_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kradmkld"))))

        init32_ = lambda54_
        nw196_ = _dafny.Array(None, 15)
        for i0_32_ in range(nw196_.length(0)):
            nw196_[i0_32_] = init32_(i0_32_)
        d_1267_v18_ = nw196_
        guard_loop_4_: int
        for guard_loop_4_ in _dafny.IntegerRange(0, (d_1267_v18_).length(0)):
            d_1269_i4_: int = guard_loop_4_
            if (True) and (((0) <= (d_1269_i4_)) and ((d_1269_i4_) < ((d_1267_v18_).length(0)))):
                (d_1267_v18_)[(d_1269_i4_)] = default__.safeDivisionInt(d_1269_i4_, d_1244_v0_)
        if (self).f10:
            if (self).f10:
                index202_ = default__.safeIndex(160, (d_1267_v18_).length(0))
                (d_1267_v18_)[index202_] = (d_1244_v0_) * (d_1244_v0_)
                d_1270_v19_: _dafny.Seq
                d_1270_v19_ = _dafny.SeqWithoutIsStrInference([(self).f10])
                d_1271_v20_: _dafny.Set
                d_1271_v20_ = _dafny.Set({not((self).f10)})
                d_1272_v21_: _dafny.Seq
                d_1272_v21_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cgysxnmd"))
                d_1273_v22_: _dafny.Array
                nw197_ = _dafny.Array(None, 29)
                nw197_[int(0)] = True
                nw197_[int(1)] = (self).fm18(d_1271_v20_, d_1272_v21_, globalState)
                nw197_[int(2)] = (((self).f9)[(self).f10] if ((self).f10) in ((self).f9) else (self).f10)
                nw197_[int(3)] = True
                nw197_[int(4)] = False
                nw197_[int(5)] = True
                nw197_[int(6)] = True
                nw197_[int(7)] = (self).f10
                nw197_[int(8)] = (self).f10
                nw197_[int(9)] = (self).f10
                nw197_[int(10)] = (self).f10
                nw197_[int(11)] = (self).f10
                nw197_[int(12)] = (self).f10
                nw197_[int(13)] = (self).f10
                nw197_[int(14)] = False
                nw197_[int(15)] = (self).f10
                nw197_[int(16)] = (self).f10
                nw197_[int(17)] = (self).f10
                nw197_[int(18)] = False
                nw197_[int(19)] = (self).f10
                nw197_[int(20)] = (self).f10
                nw197_[int(21)] = (self).f10
                nw197_[int(22)] = (self).f10
                nw197_[int(23)] = (self).f10
                nw197_[int(24)] = (self).f10
                nw197_[int(25)] = (((self).f9)[(self).f10] if ((self).f10) in ((self).f9) else (self).f10)
                nw197_[int(26)] = (self).f10
                nw197_[int(27)] = True
                nw197_[int(28)] = True
                d_1273_v22_ = nw197_
                d_1274_v23_: _dafny.Map
                d_1274_v23_ = _dafny.Map({(self).f10: (d_1254_v6_)[default__.safeIndex(293, len(d_1254_v6_))]})
                d_1275_v24_: T1
                nw198_ = C2()
                nw198_.ctor__((D1_DC4((self).fm0(d_1270_v19_, True, d_1244_v0_, (self).f10, globalState), not((self).f10), len(d_1254_v6_))).cf6, d_1273_v22_, (self).f9, (self).fm0(d_1270_v19_, (self).f10, ((d_1274_v23_)[False] if (False) in (d_1274_v23_) else (d_1254_v6_)[default__.safeIndex(d_1244_v0_, len(d_1254_v6_))]), (self).f10, globalState))
                d_1275_v24_ = nw198_
                index203_ = default__.safeIndex(160, (d_1267_v18_).length(0))
                rhs199_ = (d_1244_v0_) + (d_1244_v0_)
                rhs200_ = (d_1275_v24_ if ((self).f10) == ((self).f10) else d_1275_v24_)
                rhs201_ = (_dafny.SeqWithoutIsStrInference([p0 for d_1276_i6_ in range(default__.abs(7))])) + (d_1272_v21_)
                lhs141_ = d_1267_v18_
                lhs142_ = default__.safeIndex(160, (d_1267_v18_).length(0))
                lhs141_[lhs142_] = rhs199_
                d_1275_v24_ = rhs200_
                d_1272_v21_ = rhs201_
                index204_ = default__.safeIndex(160, (d_1267_v18_).length(0))
                index205_ = default__.safeIndex(160, (d_1267_v18_).length(0))
                def iife103_():
                    def iife105_():
                        coll51_ = _dafny.Map()
                        compr_51_: int
                        for compr_51_ in _dafny.IntegerRange(586, 962):
                            d_1277_v26_: int = compr_51_
                            if ((586) <= (d_1277_v26_)) and ((d_1277_v26_) < (962)):
                                coll51_[(d_1277_v26_) - ((d_1267_v18_)[default__.safeIndex(160, (d_1267_v18_).length(0))])] = False
                        return _dafny.Map(coll51_)
                    coll49_ = _dafny.Map()
                    def iife104_():
                        coll50_ = _dafny.Map()
                        compr_50_: int
                        for compr_50_ in _dafny.IntegerRange(586, 962):
                            d_1277_v26_: int = compr_50_
                            if ((586) <= (d_1277_v26_)) and ((d_1277_v26_) < (962)):
                                coll50_[(d_1277_v26_) - ((d_1267_v18_)[default__.safeIndex(160, (d_1267_v18_).length(0))])] = False
                        return _dafny.Map(coll50_)
                    compr_49_: int
                    for compr_49_ in (iife104_()
                    ).keys.Elements:
                        d_1278_v25_: int = compr_49_
                        if (d_1278_v25_) in (iife105_()
                        ):
                            coll49_[default__.safeModuloInt(d_1278_v25_, d_1244_v0_)] = (d_1275_v24_).f10
                    return _dafny.Map(coll49_)
                rhs202_ = d_1244_v0_
                rhs203_ = (self).f10
                rhs204_ = default__.safeDivisionInt(d_1244_v0_, (0) - ((self).fm19(len(_dafny.SeqWithoutIsStrInference([len(iife103_()
) for d_1279_i7_ in range(default__.abs(347))])), (d_1267_v18_)[default__.safeIndex(160, (d_1267_v18_).length(0))], globalState)))
                rhs205_ = d_1272_v21_
                rhs206_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cxs"))
                lhs143_ = d_1267_v18_
                lhs144_ = default__.safeIndex(160, (d_1267_v18_).length(0))
                lhs145_ = globalState
                lhs146_ = d_1267_v18_
                lhs147_ = default__.safeIndex(160, (d_1267_v18_).length(0))
                lhs143_[lhs144_] = rhs202_
                lhs145_.f8 = rhs203_
                lhs146_[lhs147_] = rhs204_
                d_1272_v21_ = rhs205_
                d_1272_v21_ = rhs206_
                d_1280_v27_: C3
                nw199_ = C3()
                nw199_.ctor__()
                d_1280_v27_ = nw199_
                index206_ = default__.safeIndex(686, (d_1273_v22_).length(0))
                (d_1273_v22_)[index206_] = not ((d_1275_v24_).f10) or ((d_1275_v24_).f10)
                d_1281_v28_: D14
                d_1281_v28_ = D14_DC35(d_1244_v0_, (d_1275_v24_).f10, p0)
                index207_ = default__.safeIndex(686, (d_1273_v22_).length(0))
                (d_1273_v22_)[index207_] = (d_1281_v28_).cf63
                d_1282_v29_: C4
                nw200_ = C4()
                nw200_.ctor__((self).f9, (d_1244_v0_) > (d_1244_v0_))
                d_1282_v29_ = nw200_
            elif True:
                d_1283_v30_: _dafny.Map
                d_1283_v30_ = _dafny.Map({(self).f10: True})
                d_1284_v31_: _dafny.MultiSet
                d_1284_v31_ = _dafny.MultiSet([d_1244_v0_])
                d_1285_v32_: _dafny.Set
                d_1285_v32_ = _dafny.Set({d_1244_v0_, len(_dafny.Map({(d_1253_v5_).cardinality: d_1244_v0_}))})
                d_1283_v30_ = (d_1283_v30_).set((self).f10, (_dafny.MultiSet([len(d_1285_v32_), d_1244_v0_])).ispropersubset(d_1284_v31_))
                d_1286_v33_: _dafny.Seq
                d_1286_v33_ = _dafny.SeqWithoutIsStrInference([(self).f10, (self).f10])
                d_1287_v34_: _dafny.Map
                d_1287_v34_ = _dafny.Map({d_1286_v33_: 953})
                d_1288_v35_: _dafny.Map
                d_1288_v35_ = _dafny.Map({((d_1287_v34_)[d_1286_v33_] if (d_1286_v33_) in (d_1287_v34_) else d_1244_v0_): (self).f10})
                d_1289_v36_: _dafny.Set
                d_1289_v36_ = _dafny.Set({(D14_DC35(len(d_1288_v35_), (self).f10, p0)).cf63})
                (globalState).f8 = not(not((self).fm18(d_1289_v36_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m")), globalState)))
                d_1290_v37_: _dafny.Seq
                d_1290_v37_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x"))
                d_1290_v37_ = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_1291_i8_ in range(default__.abs(-22))])) + (d_1290_v37_)
                d_1292_v38_: _dafny.Array
                nw201_ = _dafny.Array(False, 3)
                d_1292_v38_ = nw201_
                index208_ = default__.safeIndex(295, (d_1292_v38_).length(0))
                (d_1292_v38_)[index208_] = (self).fm18(d_1289_v36_, d_1290_v37_, globalState)
                index209_ = default__.safeIndex(295, (d_1292_v38_).length(0))
                (d_1292_v38_)[index209_] = (self).f10
                index210_ = default__.safeIndex(295, (d_1292_v38_).length(0))
                (d_1292_v38_)[index210_] = ((self).fm18(d_1289_v36_, d_1290_v37_, globalState)) in (default__.fm48((self).f10, globalState))
            (globalState).f8 = (self).f10
            d_1293_v39_: _dafny.Array
            nw202_ = _dafny.Array(None, 5)
            nw202_[int(0)] = (d_1267_v18_ if (self).f10 else d_1267_v18_)
            nw202_[int(1)] = d_1267_v18_
            nw202_[int(2)] = d_1267_v18_
            nw202_[int(3)] = d_1267_v18_
            nw202_[int(4)] = d_1267_v18_
            d_1293_v39_ = nw202_
            index211_ = default__.safeIndex(495, (d_1293_v39_).length(0))
            (d_1293_v39_)[index211_] = d_1267_v18_
            index212_ = default__.safeIndex(495, (d_1293_v39_).length(0))
            (d_1293_v39_)[index212_] = d_1267_v18_
            d_1294_v40_: _dafny.Array
            nw203_ = _dafny.Array(False, 11)
            d_1294_v40_ = nw203_
            d_1295_v41_: _dafny.Set
            d_1295_v41_ = _dafny.Set({(self).f10})
            d_1296_v42_: _dafny.Seq
            d_1296_v42_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ksbootweo"))
            index213_ = default__.safeIndex(946, (d_1294_v40_).length(0))
            (d_1294_v40_)[index213_] = ((self).fm18(d_1295_v41_, d_1296_v42_, globalState)) == ((self).f10)
            d_1297_v43_: _dafny.Seq
            d_1297_v43_ = _dafny.SeqWithoutIsStrInference([False, (self).f10])
            d_1298_v44_: _dafny.MultiSet
            d_1298_v44_ = _dafny.MultiSet([d_1244_v0_, 55, d_1244_v0_])
            d_1299_v45_: _dafny.Map
            d_1299_v45_ = _dafny.Map({(_dafny.MultiSet([d_1244_v0_, d_1244_v0_])).cardinality: p0})
            index214_ = default__.safeIndex(946, (d_1294_v40_).length(0))
            (d_1294_v40_)[index214_] = (len(d_1296_v42_)) <= ((D6_DC17((self).fm0(d_1297_v43_, (self).f10, (d_1298_v44_).cardinality, (self).f10, globalState), True, len(d_1299_v45_), d_1244_v0_, d_1244_v0_)).cf31)
            if (D0_DC1(d_1244_v0_, (((self).f9)[(d_1294_v40_)[default__.safeIndex(946, (d_1294_v40_).length(0))]] if ((d_1294_v40_)[default__.safeIndex(946, (d_1294_v40_).length(0))]) in ((self).f9) else (d_1294_v40_)[default__.safeIndex(946, (d_1294_v40_).length(0))]), d_1244_v0_)).cf2:
                (globalState).f8 = not((d_1294_v40_)[default__.safeIndex(946, (d_1294_v40_).length(0))])
                d_1300_v46_: _dafny.Array
                def lambda55_(d_1301_v40_, d_1302_v5_):
                    def lambda56_(d_1303_i9_):
                        return (d_1302_v5_ if (d_1301_v40_)[default__.safeIndex(946, (d_1301_v40_).length(0))] else _dafny.MultiSet([(d_1301_v40_)[default__.safeIndex(946, (d_1301_v40_).length(0))], (d_1301_v40_)[default__.safeIndex(946, (d_1301_v40_).length(0))]]))

                    return lambda56_

                init33_ = lambda55_(d_1294_v40_, d_1253_v5_)
                nw204_ = _dafny.Array(None, 12)
                for i0_33_ in range(nw204_.length(0)):
                    nw204_[i0_33_] = init33_(i0_33_)
                d_1300_v46_ = nw204_
                d_1304_v47_: _dafny.Array
                def lambda57_(d_1305_v43_):
                    def lambda58_(d_1306_i10_):
                        return d_1305_v43_

                    return lambda58_

                init34_ = lambda57_(d_1297_v43_)
                nw205_ = _dafny.Array(None, 27)
                for i0_34_ in range(nw205_.length(0)):
                    nw205_[i0_34_] = init34_(i0_34_)
                d_1304_v47_ = nw205_
                index215_ = default__.safeIndex(134, (d_1304_v47_).length(0))
                (d_1304_v47_)[index215_] = (d_1297_v43_) + ((d_1297_v43_).set(default__.safeIndex(d_1244_v0_, len(d_1297_v43_)), (d_1294_v40_)[default__.safeIndex(946, (d_1294_v40_).length(0))]))
                d_1307_v48_: _dafny.Array
                nw206_ = _dafny.Array(_dafny.Array(None, 0), 29)
                d_1307_v48_ = nw206_
                index216_ = default__.safeIndex(341, (d_1307_v48_).length(0))
                (d_1307_v48_)[index216_] = d_1294_v40_
                index217_ = default__.safeIndex(134, (d_1304_v47_).length(0))
                index218_ = default__.safeIndex(341, (d_1307_v48_).length(0))
                rhs207_ = d_1300_v46_
                rhs208_ = d_1297_v43_
                rhs209_ = d_1294_v40_
                rhs210_ = ((self).f10) or ((self).f10)
                rhs211_ = d_1244_v0_
                lhs148_ = d_1304_v47_
                lhs149_ = default__.safeIndex(134, (d_1304_v47_).length(0))
                lhs150_ = d_1307_v48_
                lhs151_ = default__.safeIndex(341, (d_1307_v48_).length(0))
                lhs152_ = globalState
                d_1300_v46_ = rhs207_
                lhs148_[lhs149_] = rhs208_
                lhs150_[lhs151_] = rhs209_
                lhs152_.f8 = rhs210_
                d_1244_v0_ = rhs211_
                d_1244_v0_ = (345) * (d_1244_v0_)
                d_1308_v49_: _dafny.Set
                d_1308_v49_ = _dafny.Set({d_1244_v0_, d_1244_v0_})
                d_1309_v50_: D4
                d_1309_v50_ = D4_DC11(d_1244_v0_, d_1308_v49_, (d_1294_v40_)[default__.safeIndex(946, (d_1294_v40_).length(0))])
                d_1310_v51_: _dafny.Seq
                d_1310_v51_ = _dafny.SeqWithoutIsStrInference([d_1309_v50_, d_1309_v50_, d_1309_v50_])
                d_1311_v52_: _dafny.Seq
                d_1311_v52_ = (d_1310_v51_) + (d_1310_v51_)
                d_1311_v52_ = _dafny.SeqWithoutIsStrInference([d_1309_v50_, d_1309_v50_, d_1309_v50_, d_1309_v50_, d_1309_v50_])
                arr2_ = (d_1293_v39_)[default__.safeIndex(495, (d_1293_v39_).length(0))]
                index219_ = default__.safeIndex(998, ((d_1293_v39_)[default__.safeIndex(495, (d_1293_v39_).length(0))]).length(0))
                arr2_[index219_] = default__.safeModuloInt(d_1244_v0_, len(d_1296_v42_))
                arr3_ = (d_1293_v39_)[default__.safeIndex(495, (d_1293_v39_).length(0))]
                index220_ = default__.safeIndex(998, ((d_1293_v39_)[default__.safeIndex(495, (d_1293_v39_).length(0))]).length(0))
                arr3_[index220_] = ((len((self).f9)) * ((0) - (len(d_1296_v42_)))) + (d_1244_v0_)
            elif True:
                d_1312_v53_: _dafny.Seq
                d_1312_v53_ = _dafny.SeqWithoutIsStrInference([(d_1254_v6_).set(default__.safeIndex(d_1244_v0_, len(d_1254_v6_)), d_1244_v0_)])
                d_1313_v54_: C2
                nw207_ = C2()
                nw207_.ctor__((_dafny.SeqWithoutIsStrInference([d_1244_v0_ for d_1314_i11_ in range(default__.abs(263))])) in (d_1312_v53_), d_1294_v40_, ((self).f9) | ((self).f9), (self).f10)
                d_1313_v54_ = nw207_
                rhs212_ = (0) - ((0) - (d_1244_v0_))
                rhs213_ = d_1244_v0_
                rhs214_ = d_1244_v0_
                rhs215_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tm"))) <= (d_1296_v42_)
                rhs216_ = False
                lhs153_ = globalState
                lhs154_ = d_1313_v54_
                d_1244_v0_ = rhs212_
                d_1244_v0_ = rhs213_
                d_1244_v0_ = rhs214_
                lhs153_.f8 = rhs215_
                lhs154_.f14 = rhs216_
                index221_ = default__.safeIndex(469, (d_1267_v18_).length(0))
                (d_1267_v18_)[index221_] = d_1244_v0_
                index222_ = default__.safeIndex(469, (d_1267_v18_).length(0))
                (d_1267_v18_)[index222_] = d_1244_v0_
                d_1315_v55_: _dafny.Seq
                d_1315_v55_ = _dafny.SeqWithoutIsStrInference([default__.fm50((0) - ((_dafny.MultiSet([(d_1267_v18_)[default__.safeIndex(469, (d_1267_v18_).length(0))]])).cardinality), (d_1267_v18_)[default__.safeIndex(469, (d_1267_v18_).length(0))], globalState)])
                d_1315_v55_ = d_1315_v55_
                (globalState).f8 = d_1313_v54_.f14
        elif True:
            d_1316_v56_: _dafny.Seq
            d_1316_v56_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ty"))
            d_1317_v57_: _dafny.Seq
            d_1317_v57_ = _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yof"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uhu"))), (d_1316_v56_) + (d_1316_v56_), d_1316_v56_, (d_1316_v56_) + (d_1316_v56_), d_1316_v56_])
            d_1317_v57_ = _dafny.SeqWithoutIsStrInference([d_1316_v56_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y"))])
            d_1318_v58_: D14
            d_1318_v58_ = D14_DC36(False, d_1316_v56_, d_1316_v56_)
            d_1319_v59_: _dafny.MultiSet
            d_1319_v59_ = _dafny.MultiSet([d_1318_v58_])
            d_1244_v0_ = ((d_1319_v59_).set(d_1318_v58_, default__.abs(d_1244_v0_))).cardinality
            d_1320_v60_: D1
            d_1320_v60_ = D1_DC3(d_1316_v56_)
            d_1321_v61_: D1
            d_1321_v61_ = D1_DC5(d_1320_v60_)
            source18_ = d_1321_v61_
            if source18_.is_DC3:
                d_1322___mcc_h0_ = source18_.cf5
                d_1323_cf5_ = d_1322___mcc_h0_
                d_1324_v62_: _dafny.Seq
                d_1324_v62_ = _dafny.SeqWithoutIsStrInference([d_1267_v18_])
                d_1325_v63_: _dafny.Set
                d_1325_v63_ = _dafny.Set({d_1244_v0_, d_1244_v0_})
                d_1326_v64_: D4
                d_1326_v64_ = D4_DC11(d_1244_v0_, d_1325_v63_, (self).f10)
                d_1327_v65_: _dafny.Map
                d_1327_v65_ = _dafny.Map({d_1324_v62_: ((d_1326_v64_).cf16) - (d_1244_v0_)})
                d_1327_v65_ = (d_1327_v65_).set(d_1324_v62_, (0) - (d_1244_v0_))
                d_1328_v66_: _dafny.Map
                d_1328_v66_ = _dafny.Map({False: _dafny.SeqWithoutIsStrInference([(self).f10, (self).f10])})
                d_1329_v67_: _dafny.Seq
                d_1329_v67_ = _dafny.SeqWithoutIsStrInference([False])
                d_1328_v66_ = (d_1328_v66_).set((self).f10, d_1329_v67_)
                d_1244_v0_ = default__.safeDivisionInt(len(d_1329_v67_), (self).fm1((self).f9, globalState))
                d_1330_v68_: _dafny.Array
                def lambda59_(d_1331_i12_):
                    return _dafny.CodePoint('r')

                init35_ = lambda59_
                nw208_ = _dafny.Array(None, 5)
                for i0_35_ in range(nw208_.length(0)):
                    nw208_[i0_35_] = init35_(i0_35_)
                d_1330_v68_ = nw208_
                index223_ = default__.safeIndex(45, (d_1330_v68_).length(0))
                (d_1330_v68_)[index223_] = p0
                index224_ = default__.safeIndex(45, (d_1330_v68_).length(0))
                (d_1330_v68_)[index224_] = p0
            elif source18_.is_DC4:
                d_1332___mcc_h1_ = source18_.cf6
                d_1333___mcc_h2_ = source18_.cf7
                d_1334___mcc_h3_ = source18_.cf8
                d_1335_cf8_ = d_1334___mcc_h3_
                d_1336_cf7_ = d_1333___mcc_h2_
                d_1337_cf6_ = d_1332___mcc_h1_
                d_1338_v69_: _dafny.Set
                d_1338_v69_ = _dafny.Set({976})
                d_1339_v70_: D4
                d_1339_v70_ = D4_DC11(d_1244_v0_, d_1338_v69_, (self).f10)
                d_1340_v71_: _dafny.Seq
                d_1340_v71_ = _dafny.SeqWithoutIsStrInference([d_1339_v70_])
                d_1341_v72_: _dafny.Map
                d_1341_v72_ = _dafny.Map({d_1340_v71_: d_1244_v0_})
                d_1342_v73_: _dafny.Seq
                d_1342_v73_ = _dafny.SeqWithoutIsStrInference([d_1341_v72_, _dafny.Map({d_1340_v71_: (0) - (d_1244_v0_)}), d_1341_v72_, d_1341_v72_])
                d_1343_v74_: D14
                d_1343_v74_ = D14_DC34((d_1342_v73_)[default__.safeIndex(174, len(d_1342_v73_))])
                d_1343_v74_ = default__.fm51(globalState)
                d_1344_v76_: _dafny.Array
                def lambda60_(d_1345_cf7_, d_1346_v0_):
                    def lambda61_(d_1347_i13_):
                        def iife106_():
                            coll52_ = _dafny.Set()
                            compr_52_: _dafny.Seq
                            for compr_52_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference([d_1345_cf7_]): d_1346_v0_})).keys.Elements:
                                d_1348_v75_: _dafny.Seq = compr_52_
                                if (d_1348_v75_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference([d_1345_cf7_]): d_1346_v0_})):
                                    coll52_ = coll52_.union(_dafny.Set([d_1348_v75_]))
                            return _dafny.Set(coll52_)
                        return iife106_()
                        

                    return lambda61_

                init36_ = lambda60_(d_1336_cf7_, d_1244_v0_)
                nw209_ = _dafny.Array(None, 19)
                for i0_36_ in range(nw209_.length(0)):
                    nw209_[i0_36_] = init36_(i0_36_)
                d_1344_v76_ = nw209_
                d_1349_v77_: _dafny.Seq
                d_1349_v77_ = _dafny.SeqWithoutIsStrInference([d_1336_cf7_])
                d_1350_v78_: _dafny.Set
                d_1350_v78_ = _dafny.Set({d_1349_v77_})
                index225_ = default__.safeIndex(831, (d_1344_v76_).length(0))
                (d_1344_v76_)[index225_] = d_1350_v78_
                index226_ = default__.safeIndex(831, (d_1344_v76_).length(0))
                (d_1344_v76_)[index226_] = d_1350_v78_
                d_1351_v79_: _dafny.Seq
                d_1351_v79_ = _dafny.SeqWithoutIsStrInference([d_1267_v18_, d_1267_v18_, d_1267_v18_, d_1267_v18_, d_1267_v18_])
                d_1267_v18_ = (d_1351_v79_)[default__.safeIndex((self).fm1(((self).f9).set(False, not((((self).f9)[False] if (False) in ((self).f9) else True))), globalState), len(d_1351_v79_))]
                (globalState).f8 = d_1336_cf7_
            elif source18_.is_DC2:
                d_1352___mcc_h4_ = source18_.cf4
                d_1353_cf4_ = d_1352___mcc_h4_
                rhs217_ = (self).f10
                rhs218_ = (self).f10
                lhs155_ = globalState
                lhs155_.f8 = rhs217_
                r1 = rhs218_
                d_1354_v80_: C0
                nw210_ = C0()
                nw210_.ctor__()
                d_1354_v80_ = nw210_
                index227_ = default__.safeIndex(683, (d_1267_v18_).length(0))
                (d_1267_v18_)[index227_] = (0) - ((self).fm19(d_1244_v0_, (0) - (d_1244_v0_), globalState))
                index228_ = default__.safeIndex(683, (d_1267_v18_).length(0))
                (d_1267_v18_)[index228_] = -366
                d_1355_v81_: str
                d_1355_v81_ = _dafny.CodePoint('o')
                d_1355_v81_ = d_1355_v81_
            elif True:
                d_1356___mcc_h5_ = source18_.cf9
                d_1357_cf9_ = d_1356___mcc_h5_
                d_1358_v82_: _dafny.Set
                d_1358_v82_ = _dafny.Set({d_1267_v18_})
                d_1358_v82_ = d_1358_v82_
                (globalState).f8 = (self).f10
                d_1244_v0_ = d_1244_v0_
                d_1359_v83_: _dafny.Array
                nw211_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 3)
                d_1359_v83_ = nw211_
                index229_ = default__.safeIndex(924, (d_1359_v83_).length(0))
                (d_1359_v83_)[index229_] = d_1316_v56_
                index230_ = default__.safeIndex(924, (d_1359_v83_).length(0))
                (d_1359_v83_)[index230_] = d_1316_v56_
            r1 = (self).f10
            index231_ = default__.safeIndex(74, (d_1267_v18_).length(0))
            (d_1267_v18_)[index231_] = default__.safeDivisionInt(len(d_1254_v6_), d_1244_v0_)
            index232_ = default__.safeIndex(74, (d_1267_v18_).length(0))
            (d_1267_v18_)[index232_] = 106
        d_1360_v84_: _dafny.Seq
        d_1360_v84_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dh"))
        d_1361_v85_: _dafny.Map
        d_1361_v85_ = _dafny.Map({d_1267_v18_: (d_1360_v84_) + (d_1360_v84_)})
        r0 = d_1361_v85_
        r1 = (self).f10
        return r0, r1

    def m4(self, p0, p1, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: _dafny.Seq = _dafny.Seq({})
        r2: _dafny.Set = _dafny.Set({})
        d_1362_v0_: int
        d_1362_v0_ = 49
        hi3_ = d_1362_v0_
        for d_1363_i0_ in range(default__.safeDivisionInt(d_1362_v0_, 146), hi3_):
            d_1364_v1_: _dafny.Seq
            d_1364_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "itssvdb"))
            d_1364_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mfa"))
            source19_ = D6_DC17(p1, (self).f10, d_1363_i0_, d_1362_v0_, default__.safeDivisionInt(576, d_1363_i0_))
            if source19_.is_DC17:
                d_1365___mcc_h0_ = source19_.cf29
                d_1366___mcc_h1_ = source19_.cf30
                d_1367___mcc_h2_ = source19_.cf31
                d_1368___mcc_h3_ = source19_.cf32
                d_1369___mcc_h4_ = source19_.cf33
                d_1370_cf33_ = d_1369___mcc_h4_
                d_1371_cf32_ = d_1368___mcc_h3_
                d_1372_cf31_ = d_1367___mcc_h2_
                d_1373_cf30_ = d_1366___mcc_h1_
                d_1374_cf29_ = d_1365___mcc_h0_
                d_1372_cf31_ = (d_1363_i0_) + (d_1362_v0_)
                d_1375_v2_: str
                d_1375_v2_ = _dafny.CodePoint('t')
                d_1375_v2_ = d_1375_v2_
                d_1376_v3_: _dafny.Seq
                d_1376_v3_ = _dafny.SeqWithoutIsStrInference([d_1373_cf30_])
                d_1377_v4_: D8
                d_1377_v4_ = D8_DC23(_dafny.SeqWithoutIsStrInference([(self).f10, p0, (self).fm0(d_1376_v3_, p1, 313, False, globalState)]))
                d_1377_v4_ = d_1377_v4_
                d_1378_v5_: _dafny.Set
                d_1378_v5_ = _dafny.Set({(self).f10})
                d_1379_v6_: _dafny.Array
                nw212_ = _dafny.Array(None, 4)
                nw212_[int(0)] = len(d_1378_v5_)
                nw212_[int(1)] = -973
                nw212_[int(2)] = d_1371_cf32_
                nw212_[int(3)] = d_1372_cf31_
                d_1379_v6_ = nw212_
                d_1380_v7_: _dafny.Set
                d_1380_v7_ = _dafny.Set({d_1379_v6_, d_1379_v6_})
                d_1380_v7_ = d_1380_v7_
            elif source19_.is_DC18:
                d_1381___mcc_h5_ = source19_.cf34
                d_1382___mcc_h6_ = source19_.cf35
                d_1383___mcc_h7_ = source19_.cf36
                d_1384_cf36_ = d_1383___mcc_h7_
                d_1385_cf35_ = d_1382___mcc_h6_
                d_1386_cf34_ = d_1381___mcc_h5_
                d_1387_v8_: D14
                d_1387_v8_ = D14_DC35(d_1384_cf36_, not(True), d_1386_cf34_)
                d_1384_cf36_ = ((d_1387_v8_).cf62) - (d_1385_cf35_)
                d_1388_v9_: _dafny.Array
                def lambda62_(d_1389_p1_):
                    def lambda63_(d_1390_i1_):
                        return d_1389_p1_

                    return lambda63_

                init37_ = lambda62_(p1)
                nw213_ = _dafny.Array(None, 10)
                for i0_37_ in range(nw213_.length(0)):
                    nw213_[i0_37_] = init37_(i0_37_)
                d_1388_v9_ = nw213_
                d_1391_v10_: C2
                nw214_ = C2()
                nw214_.ctor__(p1, d_1388_v9_, (self).f9, p0)
                d_1391_v10_ = nw214_
                d_1392_v11_: _dafny.Map
                d_1393_v12_: bool
                out61_: _dafny.Map
                out62_: bool
                out61_, out62_ = (d_1391_v10_).m1(d_1386_cf34_, globalState)
                d_1392_v11_ = out61_
                d_1393_v12_ = out62_
                d_1394_v13_: D0
                d_1394_v13_ = D0_DC1(d_1363_i0_, not(not(p0)), d_1385_cf35_)
                rhs219_ = ((d_1384_cf36_) - (len(_dafny.SeqWithoutIsStrInference([d_1363_i0_])))) * (d_1363_i0_)
                rhs220_ = (d_1394_v13_).cf2
                rhs221_ = p1
                rhs222_ = (self).f10
                lhs156_ = globalState
                lhs157_ = globalState
                lhs158_ = globalState
                d_1384_cf36_ = rhs219_
                lhs156_.f8 = rhs220_
                lhs157_.f8 = rhs221_
                lhs158_.f8 = rhs222_
            elif source19_.is_DC19:
                d_1395___mcc_h8_ = source19_.cf37
                d_1396___mcc_h9_ = source19_.cf38
                d_1397_cf38_ = d_1396___mcc_h9_
                d_1398_cf37_ = d_1395___mcc_h8_
                d_1362_v0_ = d_1363_i0_
                d_1399_v14_: _dafny.Seq
                d_1399_v14_ = _dafny.SeqWithoutIsStrInference([(self).f10, not((self).f10)])
                d_1400_v15_: _dafny.Set
                d_1400_v15_ = _dafny.Set({d_1399_v14_})
                d_1401_v16_: _dafny.MultiSet
                d_1401_v16_ = _dafny.MultiSet([p1])
                d_1402_v17_: _dafny.Seq
                d_1402_v17_ = _dafny.SeqWithoutIsStrInference([d_1363_i0_])
                d_1403_v18_: _dafny.Set
                d_1403_v18_ = _dafny.Set({((d_1401_v16_)[d_1398_cf37_] if (d_1398_cf37_) in (d_1401_v16_) else len(d_1402_v17_))})
                d_1398_cf37_ = (d_1403_v18_).ispropersubset(_dafny.Set({len(d_1400_v15_)}))
                d_1404_v19_: D8
                d_1404_v19_ = D8_DC23(d_1399_v14_)
                pat_let_tv26_ = d_1399_v14_
                pat_let_tv27_ = d_1399_v14_
                d_1405_v20_: _dafny.Array
                nw215_ = _dafny.Array(None, 28)
                nw215_[int(0)] = d_1404_v19_
                nw215_[int(1)] = default__.fm52(not((self).f10), globalState)
                nw215_[int(2)] = d_1404_v19_
                nw215_[int(3)] = D8_DC23(d_1399_v14_)
                nw215_[int(4)] = d_1404_v19_
                nw215_[int(5)] = d_1404_v19_
                nw215_[int(6)] = d_1404_v19_
                nw215_[int(7)] = d_1404_v19_
                def iife107_(_pat_let27_0):
                    def iife108_(d_1406_dt__update__tmp_h0_):
                        def iife109_(_pat_let28_0):
                            def iife110_(d_1407_dt__update_hcf41_h0_):
                                return D8_DC23(d_1407_dt__update_hcf41_h0_)
                            return iife110_(_pat_let28_0)
                        return iife109_(pat_let_tv26_)
                    return iife108_(_pat_let27_0)
                nw215_[int(8)] = iife107_(d_1404_v19_)
                nw215_[int(9)] = default__.fm52(d_1398_cf37_, globalState)
                nw215_[int(10)] = d_1404_v19_
                nw215_[int(11)] = d_1404_v19_
                nw215_[int(12)] = d_1404_v19_
                nw215_[int(13)] = d_1404_v19_
                nw215_[int(14)] = d_1404_v19_
                nw215_[int(15)] = d_1404_v19_
                nw215_[int(16)] = d_1404_v19_
                nw215_[int(17)] = d_1404_v19_
                nw215_[int(18)] = d_1404_v19_
                nw215_[int(19)] = d_1404_v19_
                nw215_[int(20)] = d_1404_v19_
                def iife111_(_pat_let29_0):
                    def iife112_(d_1408_dt__update__tmp_h1_):
                        def iife113_(_pat_let30_0):
                            def iife114_(d_1409_dt__update_hcf41_h1_):
                                return D8_DC23(d_1409_dt__update_hcf41_h1_)
                            return iife114_(_pat_let30_0)
                        return iife113_(pat_let_tv27_)
                    return iife112_(_pat_let29_0)
                nw215_[int(21)] = iife111_(d_1404_v19_)
                nw215_[int(22)] = D8_DC23(d_1399_v14_)
                nw215_[int(23)] = d_1404_v19_
                nw215_[int(24)] = D8_DC23(d_1399_v14_)
                nw215_[int(25)] = D8_DC23(d_1399_v14_)
                nw215_[int(26)] = d_1404_v19_
                nw215_[int(27)] = d_1404_v19_
                d_1405_v20_ = nw215_
                d_1410_v21_: _dafny.Seq
                d_1410_v21_ = _dafny.SeqWithoutIsStrInference([d_1405_v20_])
                rhs223_ = d_1410_v21_
                rhs224_ = p1
                rhs225_ = d_1363_i0_
                lhs159_ = globalState
                d_1410_v21_ = rhs223_
                lhs159_.f8 = rhs224_
                d_1362_v0_ = rhs225_
                d_1411_v22_: _dafny.Map
                d_1411_v22_ = _dafny.Map({d_1397_cf38_: d_1397_cf38_})
                d_1412_v23_: D11
                d_1412_v23_ = D11_DC28(d_1362_v0_, d_1398_cf37_, _dafny.Map({len(d_1403_v18_): d_1411_v22_}), (self).f10, (d_1364_v1_).set(default__.safeIndex(d_1363_i0_, len(d_1364_v1_)), d_1397_cf38_))
                (globalState).f8 = not ((len(_dafny.Map({len(d_1364_v1_): d_1397_cf38_}))) > (d_1362_v0_)) or ((d_1412_v23_).cf51)
            elif source19_.is_DC16:
                d_1413___mcc_h10_ = source19_.cf28
                d_1414_cf28_ = d_1413___mcc_h10_
                d_1415_v24_: C3
                nw216_ = C3()
                nw216_.ctor__()
                d_1415_v24_ = nw216_
                d_1416_v25_: str
                d_1416_v25_ = _dafny.CodePoint('x')
                d_1417_v26_: _dafny.MultiSet
                d_1417_v26_ = _dafny.MultiSet([d_1363_i0_])
                rhs226_ = (self).f10
                rhs227_ = (d_1415_v24_ if (d_1417_v26_).issubset(_dafny.MultiSet([d_1362_v0_, d_1362_v0_])) else d_1415_v24_)
                rhs228_ = p1
                rhs229_ = d_1416_v25_
                lhs160_ = globalState
                lhs161_ = globalState
                lhs160_.f8 = rhs226_
                d_1415_v24_ = rhs227_
                lhs161_.f8 = rhs228_
                d_1416_v25_ = rhs229_
                (globalState).f8 = not ((len(d_1364_v1_)) > (d_1362_v0_)) or (p1)
                d_1418_v27_: D7
                d_1418_v27_ = D7_DC21(_dafny.Set({p1, False}))
                d_1419_v28_: _dafny.Set
                d_1419_v28_ = _dafny.Set({(self).f10, (self).f10})
                d_1418_v27_ = D7_DC21(d_1419_v28_)
                d_1420_v29_: _dafny.Seq
                d_1420_v29_ = _dafny.SeqWithoutIsStrInference([d_1417_v26_, _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(0) - (d_1362_v0_) for d_1421_i2_ in range(default__.abs(-567))]))])
                d_1422_v30_: _dafny.MultiSet
                d_1422_v30_ = _dafny.MultiSet([p1, p0])
                d_1423_v31_: _dafny.Set
                d_1423_v31_ = _dafny.Set({628, (d_1422_v30_).cardinality})
                d_1424_v32_: D4
                d_1424_v32_ = D4_DC11(914, d_1423_v31_, (self).f10)
                d_1425_v33_: _dafny.Map
                d_1425_v33_ = _dafny.Map({d_1362_v0_: p1})
                pat_let_tv28_ = d_1425_v33_
                pat_let_tv29_ = d_1425_v33_
                pat_let_tv30_ = d_1424_v32_
                def iife115_(_pat_let31_0):
                    def iife116_(d_1426_dt__update__tmp_h2_):
                        def iife117_(_pat_let32_0):
                            def iife118_(d_1427_dt__update_hcf7_h0_):
                                def iife119_(_pat_let33_0):
                                    def iife120_(d_1428_dt__update_hcf8_h0_):
                                        return D1_DC4((d_1426_dt__update__tmp_h2_).cf6, d_1427_dt__update_hcf7_h0_, d_1428_dt__update_hcf8_h0_)
                                    return iife120_(_pat_let33_0)
                                return iife119_((pat_let_tv30_).cf16)
                            return iife118_(_pat_let32_0)
                        return iife117_(((pat_let_tv28_)[d_1363_i0_] if (d_1363_i0_) in (pat_let_tv29_) else (self).f10))
                    return iife116_(_pat_let31_0)
                d_1362_v0_ = (iife115_(D1_DC4(p1, p1, len(d_1420_v29_)))).cf8
            elif True:
                d_1429___mcc_h11_ = source19_.cf39
                d_1430_cf39_ = d_1429___mcc_h11_
                d_1431_v34_: _dafny.Seq
                d_1431_v34_ = _dafny.SeqWithoutIsStrInference([p1])
                (globalState).f8 = ((self).fm0(d_1431_v34_, (d_1431_v34_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "axewyqwhx"))), len(d_1431_v34_))], d_1362_v0_, (self).f10, globalState)) and (p0)
                d_1432_v35_: _dafny.Seq
                d_1432_v35_ = _dafny.SeqWithoutIsStrInference([d_1362_v0_, d_1363_i0_])
                d_1433_v36_: _dafny.MultiSet
                d_1433_v36_ = _dafny.MultiSet([p1, (self).f10, (self).f10])
                d_1434_v37_: _dafny.Map
                d_1434_v37_ = _dafny.Map({(d_1433_v36_).cardinality: d_1363_i0_})
                d_1435_v38_: _dafny.Array
                nw217_ = _dafny.Array(None, 12)
                nw217_[int(0)] = (0) - ((0) - (d_1363_i0_))
                nw217_[int(1)] = (self).fm1((self).f9, globalState)
                nw217_[int(2)] = (len(d_1431_v34_)) * ((_dafny.MultiSet(d_1432_v35_)).cardinality)
                nw217_[int(3)] = (0) - ((self).fm19(d_1362_v0_, d_1363_i0_, globalState))
                nw217_[int(4)] = d_1363_i0_
                nw217_[int(5)] = d_1362_v0_
                nw217_[int(6)] = len((d_1434_v37_) | (d_1434_v37_))
                nw217_[int(7)] = d_1362_v0_
                nw217_[int(8)] = d_1362_v0_
                nw217_[int(9)] = d_1363_i0_
                nw217_[int(10)] = -769
                nw217_[int(11)] = (197) * (d_1363_i0_)
                d_1435_v38_ = nw217_
                index233_ = default__.safeIndex(494, (d_1435_v38_).length(0))
                (d_1435_v38_)[index233_] = default__.safeModuloInt(543, len((d_1434_v37_).set(d_1362_v0_, d_1363_i0_)))
                d_1436_v39_: D5
                d_1436_v39_ = D5_DC14(d_1362_v0_, p1)
                d_1437_v40_: _dafny.Seq
                d_1437_v40_ = _dafny.SeqWithoutIsStrInference([d_1436_v39_])
                index234_ = default__.safeIndex(494, (d_1435_v38_).length(0))
                rhs230_ = p0
                rhs231_ = (self).f10
                rhs232_ = default__.fm33(p0, (_dafny.SeqWithoutIsStrInference([D5_DC14(d_1363_i0_, p0) for d_1438_i3_ in range(default__.abs(868))])) <= (d_1437_v40_), (((d_1433_v36_)[(self).f10] if ((self).f10) in (d_1433_v36_) else d_1362_v0_)) * ((0) - (d_1363_i0_)), (self).fm0(d_1431_v34_, p0, d_1363_i0_, (self).f10, globalState), globalState)
                lhs162_ = globalState
                lhs163_ = globalState
                lhs164_ = d_1435_v38_
                lhs165_ = default__.safeIndex(494, (d_1435_v38_).length(0))
                lhs162_.f8 = rhs230_
                lhs163_.f8 = rhs231_
                lhs164_[lhs165_] = rhs232_
                d_1439_v41_: str
                d_1439_v41_ = _dafny.CodePoint('o')
                d_1439_v41_ = d_1439_v41_
                d_1440_v42_: _dafny.Array
                nw218_ = _dafny.Array(None, 11)
                nw218_[int(0)] = d_1364_v1_
                nw218_[int(1)] = _dafny.SeqWithoutIsStrInference([d_1439_v41_, d_1439_v41_, d_1439_v41_])
                nw218_[int(2)] = d_1364_v1_
                nw218_[int(3)] = _dafny.SeqWithoutIsStrInference([d_1439_v41_])
                nw218_[int(4)] = d_1364_v1_
                nw218_[int(5)] = (d_1364_v1_) + (d_1364_v1_)
                nw218_[int(6)] = d_1364_v1_
                nw218_[int(7)] = default__.fm20(753, p1, not(False), (d_1435_v38_)[default__.safeIndex(494, (d_1435_v38_).length(0))], globalState)
                nw218_[int(8)] = d_1364_v1_
                nw218_[int(9)] = _dafny.SeqWithoutIsStrInference([d_1439_v41_])
                nw218_[int(10)] = d_1364_v1_
                d_1440_v42_ = nw218_
                index235_ = default__.safeIndex(397, (d_1440_v42_).length(0))
                (d_1440_v42_)[index235_] = d_1364_v1_
                index236_ = default__.safeIndex(397, (d_1440_v42_).length(0))
                (d_1440_v42_)[index236_] = (d_1364_v1_) + (d_1364_v1_)
            if False:
                d_1441_v43_: _dafny.Map
                d_1441_v43_ = _dafny.Map({d_1363_i0_: (self).f10})
                d_1442_v45_: _dafny.Seq
                def iife121_():
                    coll53_ = _dafny.Map()
                    compr_53_: int
                    for compr_53_ in _dafny.IntegerRange(185, 321):
                        d_1443_v44_: int = compr_53_
                        if ((185) <= (d_1443_v44_)) and ((d_1443_v44_) < (321)):
                            coll53_[(d_1443_v44_) * (887)] = p0
                    return _dafny.Map(coll53_)
                d_1442_v45_ = _dafny.SeqWithoutIsStrInference([d_1441_v43_, d_1441_v43_, (d_1441_v43_) | (iife121_()
                ), (default__.fm44(p1, globalState)) | (default__.fm44(p1, globalState)), d_1441_v43_])
                d_1444_v46_: D16
                d_1444_v46_ = D16_DC40(d_1442_v45_)
                d_1442_v45_ = (((_dafny.SeqWithoutIsStrInference([d_1441_v43_])) + ((d_1444_v46_).cf76)) + ((d_1442_v45_) + (d_1442_v45_))).set(default__.safeIndex((0) - (d_1363_i0_), len(((_dafny.SeqWithoutIsStrInference([d_1441_v43_])) + ((d_1444_v46_).cf76)) + ((d_1442_v45_) + (d_1442_v45_)))), d_1441_v43_)
                d_1445_v47_: _dafny.Seq
                d_1445_v47_ = _dafny.SeqWithoutIsStrInference([(self).f10, p0, p1])
                (globalState).f8 = not((self).fm0(d_1445_v47_, (p1) == ((self).f10), default__.safeDivisionInt((_dafny.MultiSet([-494])).cardinality, d_1362_v0_), p0, globalState))
                d_1446_v48_: _dafny.Array
                nw219_ = _dafny.Array(int(0), 2)
                d_1446_v48_ = nw219_
                d_1447_v49_: _dafny.Map
                d_1447_v49_ = _dafny.Map({_dafny.CodePoint('n'): d_1446_v48_})
                d_1448_v50_: str
                d_1448_v50_ = _dafny.CodePoint('u')
                d_1449_v51_: C3
                nw220_ = C3()
                nw220_.ctor__()
                d_1449_v51_ = nw220_
                d_1450_v52_: _dafny.Map
                d_1450_v52_ = _dafny.Map({((d_1447_v49_)[d_1448_v50_] if (d_1448_v50_) in (d_1447_v49_) else d_1446_v48_): d_1449_v51_})
                d_1450_v52_ = ((d_1450_v52_) | (d_1450_v52_)) | (d_1450_v52_)
                d_1448_v50_ = d_1448_v50_
                d_1362_v0_ = d_1363_i0_
            elif True:
                d_1451_v53_: D1
                d_1451_v53_ = D1_DC3(d_1364_v1_)
                d_1452_v54_: _dafny.Seq
                d_1452_v54_ = _dafny.SeqWithoutIsStrInference([(d_1362_v0_) + (d_1362_v0_), len((d_1451_v53_).cf5)])
                d_1452_v54_ = d_1452_v54_
                d_1362_v0_ = (0) - (((0) - (len(d_1364_v1_))) + (d_1363_i0_))
                d_1453_v55_: _dafny.Array
                nw221_ = _dafny.Array(int(0), 1)
                d_1453_v55_ = nw221_
                d_1454_v56_: _dafny.Seq
                d_1454_v56_ = _dafny.SeqWithoutIsStrInference([p1, p0])
                d_1455_v57_: _dafny.MultiSet
                d_1455_v57_ = _dafny.MultiSet([d_1454_v56_, d_1454_v56_])
                d_1456_v58_: _dafny.Map
                d_1456_v58_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ayxnih"))): d_1363_i0_})
                rhs233_ = d_1453_v55_
                rhs234_ = (default__.safeDivisionInt(d_1362_v0_, (d_1455_v57_).cardinality) if (d_1362_v0_) >= (d_1363_i0_) else ((d_1456_v58_)[(0) - (d_1362_v0_)] if ((0) - (d_1362_v0_)) in (d_1456_v58_) else d_1363_i0_))
                rhs235_ = (0) - (d_1363_i0_)
                d_1453_v55_ = rhs233_
                d_1362_v0_ = rhs234_
                d_1362_v0_ = rhs235_
                d_1457_v59_: _dafny.Set
                d_1457_v59_ = _dafny.Set({(d_1452_v54_) + (d_1452_v54_)})
                d_1362_v0_ = len(d_1457_v59_)
                d_1458_v60_: _dafny.MultiSet
                d_1458_v60_ = _dafny.MultiSet([False, True, False, p1, p1])
                d_1459_v61_: _dafny.Set
                d_1459_v61_ = _dafny.Set({d_1453_v55_, d_1453_v55_, d_1453_v55_, d_1453_v55_, d_1453_v55_})
                d_1460_v62_: D8
                d_1460_v62_ = D8_DC24((d_1458_v60_).ispropersubset(d_1458_v60_), d_1459_v61_, d_1362_v0_)
                rhs236_ = d_1460_v62_
                rhs237_ = p0
                rhs238_ = p0
                rhs239_ = (d_1362_v0_) + (d_1362_v0_)
                rhs240_ = (self).f10
                lhs166_ = globalState
                lhs167_ = globalState
                lhs168_ = globalState
                d_1460_v62_ = rhs236_
                lhs166_.f8 = rhs237_
                lhs167_.f8 = rhs238_
                d_1362_v0_ = rhs239_
                lhs168_.f8 = rhs240_
            d_1461_v63_: _dafny.Array
            nw222_ = _dafny.Array(_dafny.Map({}), 11)
            d_1461_v63_ = nw222_
            d_1462_v64_: _dafny.Map
            d_1462_v64_ = _dafny.Map({(self).f10: len(d_1364_v1_)})
            d_1463_v65_: _dafny.Seq
            d_1463_v65_ = _dafny.SeqWithoutIsStrInference([d_1363_i0_, d_1363_i0_])
            d_1464_v66_: _dafny.Map
            d_1464_v66_ = _dafny.Map({d_1462_v64_: d_1463_v65_})
            index237_ = default__.safeIndex(396, (d_1461_v63_).length(0))
            (d_1461_v63_)[index237_] = (d_1464_v66_) | (d_1464_v66_)
            d_1465_v67_: _dafny.Seq
            d_1465_v67_ = _dafny.SeqWithoutIsStrInference([d_1464_v66_, d_1464_v66_, _dafny.Map({d_1462_v64_: d_1463_v65_})])
            index238_ = default__.safeIndex(396, (d_1461_v63_).length(0))
            (d_1461_v63_)[index238_] = ((d_1464_v66_) | ((d_1465_v67_)[default__.safeIndex(d_1362_v0_, len(d_1465_v67_))])) | (d_1464_v66_)
        (globalState).f8 = p1
        d_1466_v68_: _dafny.MultiSet
        d_1466_v68_ = _dafny.MultiSet([(d_1362_v0_) - (d_1362_v0_), (self).fm1((self).f9, globalState), 208, d_1362_v0_])
        d_1466_v68_ = (d_1466_v68_).set(default__.safeDivisionInt(d_1362_v0_, d_1362_v0_), default__.abs(d_1362_v0_))
        (globalState).f8 = p0
        d_1467_v69_: D5
        d_1467_v69_ = D5_DC14(75, p1)
        d_1468_v70_: _dafny.Array
        nw223_ = _dafny.Array(int(0), 25)
        d_1468_v70_ = nw223_
        d_1469_v71_: _dafny.Set
        d_1469_v71_ = _dafny.Set({d_1468_v70_})
        d_1470_v72_: _dafny.Seq
        d_1470_v72_ = _dafny.SeqWithoutIsStrInference([d_1362_v0_])
        d_1471_v73_: _dafny.Seq
        d_1471_v73_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jeudkbik"))
        d_1472_v74_: _dafny.Map
        d_1472_v74_ = _dafny.Map({d_1362_v0_: p0})
        d_1473_v75_: str
        d_1473_v75_ = _dafny.CodePoint('e')
        d_1474_v76_: _dafny.Array
        nw224_ = _dafny.Array(None, 27)
        nw224_[int(0)] = d_1362_v0_
        nw224_[int(1)] = len(_dafny.Map({(0) - ((d_1467_v69_).cf25): len(d_1469_v71_)}))
        nw224_[int(2)] = d_1362_v0_
        nw224_[int(3)] = d_1362_v0_
        nw224_[int(4)] = -813
        nw224_[int(5)] = d_1362_v0_
        nw224_[int(6)] = (0) - (len((d_1470_v72_) + (d_1470_v72_)))
        nw224_[int(7)] = d_1362_v0_
        nw224_[int(8)] = len(d_1471_v73_)
        nw224_[int(9)] = len(d_1472_v74_)
        nw224_[int(10)] = (0) - (d_1362_v0_)
        nw224_[int(11)] = (d_1362_v0_) * (184)
        nw224_[int(12)] = d_1362_v0_
        nw224_[int(13)] = d_1362_v0_
        nw224_[int(14)] = d_1362_v0_
        nw224_[int(15)] = d_1362_v0_
        nw224_[int(16)] = (d_1362_v0_) - (d_1362_v0_)
        nw224_[int(17)] = (d_1470_v72_)[default__.safeIndex(d_1362_v0_, len(d_1470_v72_))]
        nw224_[int(18)] = d_1362_v0_
        nw224_[int(19)] = default__.safeModuloInt(d_1362_v0_, d_1362_v0_)
        nw224_[int(20)] = 822
        nw224_[int(21)] = (0) - (default__.safeModuloInt(d_1362_v0_, d_1362_v0_))
        nw224_[int(22)] = len(d_1471_v73_)
        nw224_[int(23)] = 104
        nw224_[int(24)] = d_1362_v0_
        nw224_[int(25)] = (d_1362_v0_) - (464)
        nw224_[int(26)] = len(_dafny.Map({d_1473_v75_: (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nbqhra"))).set(default__.safeIndex(d_1362_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nbqhra")))), d_1473_v75_)}))
        d_1474_v76_ = nw224_
        index239_ = default__.safeIndex(655, (d_1468_v70_).length(0))
        (d_1468_v70_)[index239_] = d_1362_v0_
        d_1475_v77_: D6
        d_1475_v77_ = D6_DC18(_dafny.CodePoint('m'), d_1362_v0_, d_1362_v0_)
        pat_let_tv31_ = p1
        pat_let_tv32_ = d_1362_v0_
        pat_let_tv33_ = d_1362_v0_
        pat_let_tv34_ = d_1471_v73_
        pat_let_tv35_ = d_1471_v73_
        pat_let_tv36_ = d_1471_v73_
        index240_ = default__.safeIndex(655, (d_1468_v70_).length(0))
        def lambda64_(source20_):
            if source20_.is_DC17:
                d_1476___mcc_h12_ = source20_.cf29
                d_1477___mcc_h13_ = source20_.cf30
                d_1478___mcc_h14_ = source20_.cf31
                d_1479___mcc_h15_ = source20_.cf32
                d_1480___mcc_h16_ = source20_.cf33
                d_1481_cf33_ = d_1480___mcc_h16_
                d_1482_cf32_ = d_1479___mcc_h15_
                d_1483_cf31_ = d_1478___mcc_h14_
                d_1484_cf30_ = d_1477___mcc_h13_
                d_1485_cf29_ = d_1476___mcc_h12_
                return len((_dafny.Set({d_1484_cf30_, pat_let_tv31_})).intersection(_dafny.Set({d_1485_cf29_})))
            elif source20_.is_DC18:
                d_1486___mcc_h17_ = source20_.cf34
                d_1487___mcc_h18_ = source20_.cf35
                d_1488___mcc_h19_ = source20_.cf36
                d_1489_cf36_ = d_1488___mcc_h19_
                d_1490_cf35_ = d_1487___mcc_h18_
                d_1491_cf34_ = d_1486___mcc_h17_
                return pat_let_tv32_
            elif source20_.is_DC19:
                d_1492___mcc_h20_ = source20_.cf37
                d_1493___mcc_h21_ = source20_.cf38
                d_1494_cf38_ = d_1493___mcc_h21_
                d_1495_cf37_ = d_1492___mcc_h20_
                return pat_let_tv33_
            elif source20_.is_DC16:
                d_1496___mcc_h22_ = source20_.cf28
                d_1497_cf28_ = d_1496___mcc_h22_
                return (len(pat_let_tv34_)) - ((0) - (len(d_1497_cf28_)))
            elif True:
                d_1498___mcc_h23_ = source20_.cf39
                d_1499_cf39_ = d_1498___mcc_h23_
                return (0) - (len((pat_let_tv35_) + (pat_let_tv36_)))

        (d_1468_v70_)[index240_] = lambda64_(d_1475_v77_)
        guard_loop_5_: int
        for guard_loop_5_ in _dafny.IntegerRange(0, (d_1474_v76_).length(0)):
            d_1500_i4_: int = guard_loop_5_
            if (True) and (((0) <= (d_1500_i4_)) and ((d_1500_i4_) < ((d_1474_v76_).length(0)))):
                (d_1474_v76_)[(d_1500_i4_)] = (d_1500_i4_) * (703)
        d_1501_v78_: _dafny.Array
        nw225_ = _dafny.Array(False, 1)
        d_1501_v78_ = nw225_
        r0 = d_1501_v78_
        d_1502_v79_: D0
        d_1502_v79_ = D0_DC0(p1)
        d_1503_v80_: _dafny.Seq
        d_1503_v80_ = _dafny.SeqWithoutIsStrInference([d_1502_v79_, d_1502_v79_, d_1502_v79_, D0_DC0(p0)])
        r1 = d_1503_v80_
        d_1504_v82_: D4
        def iife122_():
            coll54_ = _dafny.Set()
            compr_54_: int
            for compr_54_ in _dafny.IntegerRange(742, -453):
                d_1505_v81_: int = compr_54_
                if ((742) <= (d_1505_v81_)) and ((d_1505_v81_) < (-453)):
                    coll54_ = coll54_.union(_dafny.Set([(d_1505_v81_) - ((d_1468_v70_)[default__.safeIndex(655, (d_1468_v70_).length(0))])]))
            return _dafny.Set(coll54_)
        d_1504_v82_ = D4_DC11(919, iife122_()
, True)
        r2 = ((D4_DC11(d_1362_v0_, _dafny.Set({len(d_1471_v73_)}), p1) if False else d_1504_v82_)).cf17
        return r0, r1, r2


class C6(T1, T0):
    def  __init__(self):
        self._f9: _dafny.Map = _dafny.Map({})
        self._f10: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C6"
    @property
    def f9(self):
        return self._f9
    @property
    def f10(self):
        return self._f10
    def ctor__(self, f9, f10):
        (self)._f9 = f9
        (self)._f10 = f10

    def fm0(self, p0, p1, p2, p3, globalState):
        return (self).f10

    def fm1(self, p0, globalState):
        return 93

    def fm8(self, globalState):
        source21_ = D1_DC4((self).f10, (self).f10, -971)
        if source21_.is_DC3:
            d_1506___mcc_h0_ = source21_.cf5
            d_1507_cf5_ = d_1506___mcc_h0_
            return (len(d_1507_cf5_)) - (-673)
        elif source21_.is_DC4:
            d_1508___mcc_h1_ = source21_.cf6
            d_1509___mcc_h2_ = source21_.cf7
            d_1510___mcc_h3_ = source21_.cf8
            d_1511_cf8_ = d_1510___mcc_h3_
            d_1512_cf7_ = d_1509___mcc_h2_
            d_1513_cf6_ = d_1508___mcc_h1_
            return d_1511_cf8_
        elif source21_.is_DC2:
            d_1514___mcc_h4_ = source21_.cf4
            d_1515_cf4_ = d_1514___mcc_h4_
            return (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([(self).f10, (D0_DC1(576, (self).f10, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qfvjs"))))).cf2]), _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f10]))])) for d_1516_i0_ in range(default__.abs(862))]))) + (len(_dafny.SeqWithoutIsStrInference([(self).f10, True, (self).f10, (self).f10, (self).f10])))
        elif True:
            d_1517___mcc_h5_ = source21_.cf9
            d_1518_cf9_ = d_1517___mcc_h5_
            return (0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.Set({462, 3}) for d_1519_i1_ in range(default__.abs(753))])))

    def fm9(self, p0, p1, p2, p3, globalState):
        return not((self).f10)

    def m0(self, globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r1: bool = False
        r2: _dafny.Seq = _dafny.Seq({})
        r3: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        if not(not ((self).f10) or (True)):
            d_1520_v0_: _dafny.Array
            def lambda65_(d_1521_i0_):
                return (d_1521_i0_) - (len((self).f9))

            init38_ = lambda65_
            nw226_ = _dafny.Array(None, 21)
            for i0_38_ in range(nw226_.length(0)):
                nw226_[i0_38_] = init38_(i0_38_)
            d_1520_v0_ = nw226_
            d_1522_v1_: _dafny.Set
            d_1522_v1_ = _dafny.Set({(self).f10, (self).f10, (self).f10})
            index241_ = default__.safeIndex(333, (d_1520_v0_).length(0))
            (d_1520_v0_)[index241_] = (len(d_1522_v1_)) - (len((self).f9))
            d_1523_v2_: int
            d_1523_v2_ = -316
            index242_ = default__.safeIndex(333, (d_1520_v0_).length(0))
            (d_1520_v0_)[index242_] = (d_1523_v2_) * (d_1523_v2_)
            d_1524_v4_: _dafny.Array
            def lambda66_(d_1525_v0_):
                def lambda67_(d_1526_i1_):
                    def iife123_():
                        coll55_ = _dafny.Set()
                        compr_55_: int
                        for compr_55_ in _dafny.IntegerRange(249, 61):
                            d_1527_v3_: int = compr_55_
                            if ((249) <= (d_1527_v3_)) and ((d_1527_v3_) < (61)):
                                coll55_ = coll55_.union(_dafny.Set([(d_1527_v3_) * ((d_1525_v0_)[default__.safeIndex(333, (d_1525_v0_).length(0))])]))
                        return _dafny.Set(coll55_)
                    return (iife123_()
                    ) - (_dafny.Set({(d_1525_v0_)[default__.safeIndex(333, (d_1525_v0_).length(0))]}))

                return lambda67_

            init39_ = lambda66_(d_1520_v0_)
            nw227_ = _dafny.Array(None, 19)
            for i0_39_ in range(nw227_.length(0)):
                nw227_[i0_39_] = init39_(i0_39_)
            d_1524_v4_ = nw227_
            d_1524_v4_ = d_1524_v4_
            d_1528_v5_: _dafny.Array
            def lambda68_(d_1529_i2_):
                return D0_DC0((self).f10)

            init40_ = lambda68_
            nw228_ = _dafny.Array(None, 11)
            for i0_40_ in range(nw228_.length(0)):
                nw228_[i0_40_] = init40_(i0_40_)
            d_1528_v5_ = nw228_
            d_1530_v6_: D0
            d_1530_v6_ = D0_DC0((self).f10)
            index243_ = default__.safeIndex(302, (d_1528_v5_).length(0))
            def iife124_(_pat_let34_0):
                def iife125_(d_1531_dt__update__tmp_h0_):
                    def iife126_(_pat_let35_0):
                        def iife127_(d_1532_dt__update_hcf0_h0_):
                            return D0_DC0(d_1532_dt__update_hcf0_h0_)
                        return iife127_(_pat_let35_0)
                    return iife126_(not(not((self).f10)))
                return iife125_(_pat_let34_0)
            (d_1528_v5_)[index243_] = iife124_(d_1530_v6_)
            index244_ = default__.safeIndex(302, (d_1528_v5_).length(0))
            def iife128_(_pat_let36_0):
                def iife129_(d_1533_dt__update__tmp_h1_):
                    def iife130_(_pat_let37_0):
                        def iife131_(d_1534_dt__update_hcf0_h1_):
                            return D0_DC0(d_1534_dt__update_hcf0_h1_)
                        return iife131_(_pat_let37_0)
                    return iife130_((self).f10)
                return iife129_(_pat_let36_0)
            (d_1528_v5_)[index244_] = iife128_(d_1530_v6_)
            d_1535_v7_: C0
            nw229_ = C0()
            nw229_.ctor__()
            d_1535_v7_ = nw229_
            d_1536_v8_: _dafny.Seq
            d_1536_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qbbn"))
            d_1537_v9_: _dafny.Map
            d_1537_v9_ = _dafny.Map({(d_1520_v0_)[default__.safeIndex(333, (d_1520_v0_).length(0))]: d_1523_v2_})
            d_1538_v10_: _dafny.MultiSet
            d_1538_v10_ = _dafny.MultiSet([(d_1520_v0_)[default__.safeIndex(333, (d_1520_v0_).length(0))], 312])
            (globalState).f8 = (default__.fm10((self).f10, d_1523_v2_, d_1536_v8_, d_1537_v9_, globalState)).isdisjoint(d_1538_v10_)
        elif True:
            d_1539_v11_: C0
            nw230_ = C0()
            nw230_.ctor__()
            d_1539_v11_ = nw230_
            d_1540_v12_: _dafny.Array
            def lambda69_(d_1541_i3_):
                return (self).f10

            init41_ = lambda69_
            nw231_ = _dafny.Array(None, 21)
            for i0_41_ in range(nw231_.length(0)):
                nw231_[i0_41_] = init41_(i0_41_)
            d_1540_v12_ = nw231_
            index245_ = default__.safeIndex(294, (d_1540_v12_).length(0))
            (d_1540_v12_)[index245_] = ((self).f10) == ((self).f10)
            d_1542_v13_: int
            d_1542_v13_ = 274
            index246_ = default__.safeIndex(294, (d_1540_v12_).length(0))
            (d_1540_v12_)[index246_] = ((0) - (d_1542_v13_)) > (d_1542_v13_)
            d_1543_v14_: D1
            d_1543_v14_ = D1_DC4((self).f10, False, (self).fm1((self).f9, globalState))
            rhs241_ = d_1543_v14_
            rhs242_ = (self).f10
            rhs243_ = d_1542_v13_
            rhs244_ = False
            lhs169_ = globalState
            lhs170_ = globalState
            d_1543_v14_ = rhs241_
            lhs169_.f8 = rhs242_
            d_1542_v13_ = rhs243_
            lhs170_.f8 = rhs244_
            d_1544_v15_: _dafny.Map
            d_1544_v15_ = _dafny.Map({d_1542_v13_: (self).f9})
            if (((d_1544_v15_)[d_1542_v13_] if (d_1542_v13_) in (d_1544_v15_) else (self).f9)) == ((self).f9):
                d_1545_v16_: _dafny.Map
                d_1545_v16_ = _dafny.Map({d_1542_v13_: d_1542_v13_})
                d_1546_v17_: _dafny.MultiSet
                d_1546_v17_ = _dafny.MultiSet([(self).f10, (d_1540_v12_)[default__.safeIndex(294, (d_1540_v12_).length(0))], (d_1540_v12_)[default__.safeIndex(294, (d_1540_v12_).length(0))], (d_1540_v12_)[default__.safeIndex(294, (d_1540_v12_).length(0))], (d_1540_v12_)[default__.safeIndex(294, (d_1540_v12_).length(0))]])
                d_1542_v13_ = (d_1542_v13_) * (((d_1545_v16_)[(d_1546_v17_).cardinality] if ((d_1546_v17_).cardinality) in (d_1545_v16_) else d_1542_v13_))
                d_1547_v18_: _dafny.Map
                d_1547_v18_ = _dafny.Map({d_1542_v13_: (d_1540_v12_)[default__.safeIndex(294, (d_1540_v12_).length(0))]})
                d_1548_v19_: _dafny.MultiSet
                d_1548_v19_ = _dafny.MultiSet([d_1542_v13_])
                d_1549_v20_: _dafny.Set
                d_1549_v20_ = _dafny.Set({d_1548_v19_})
                d_1550_v22_: _dafny.Seq
                d_1550_v22_ = _dafny.SeqWithoutIsStrInference([d_1542_v13_, d_1542_v13_])
                d_1551_v23_: _dafny.Array
                nw232_ = _dafny.Array(None, 21)
                nw232_[int(0)] = (_dafny.Map({d_1542_v13_: (d_1540_v12_)[default__.safeIndex(294, (d_1540_v12_).length(0))]})) | (d_1547_v18_)
                nw232_[int(1)] = d_1547_v18_
                nw232_[int(2)] = (default__.fm11(len(d_1549_v20_), d_1542_v13_, False, d_1542_v13_, globalState)) | (d_1547_v18_)
                nw232_[int(3)] = (d_1547_v18_) | (d_1547_v18_)
                nw232_[int(4)] = d_1547_v18_
                nw232_[int(5)] = d_1547_v18_
                nw232_[int(6)] = d_1547_v18_
                nw232_[int(7)] = d_1547_v18_
                nw232_[int(8)] = d_1547_v18_
                nw232_[int(9)] = d_1547_v18_
                def iife132_():
                    coll56_ = _dafny.Map()
                    compr_56_: int
                    for compr_56_ in _dafny.IntegerRange(0, 445):
                        d_1552_v21_: int = compr_56_
                        if ((0) <= (d_1552_v21_)) and ((d_1552_v21_) < (445)):
                            coll56_[(d_1552_v21_) * (351)] = (d_1540_v12_)[default__.safeIndex(294, (d_1540_v12_).length(0))]
                    return _dafny.Map(coll56_)
                nw232_[int(10)] = iife132_()
                
                nw232_[int(11)] = default__.fm11((d_1550_v22_)[default__.safeIndex(len(d_1545_v16_), len(d_1550_v22_))], 439, (d_1540_v12_)[default__.safeIndex(294, (d_1540_v12_).length(0))], d_1542_v13_, globalState)
                nw232_[int(12)] = d_1547_v18_
                nw232_[int(13)] = d_1547_v18_
                nw232_[int(14)] = (d_1547_v18_) | (d_1547_v18_)
                nw232_[int(15)] = _dafny.Map({d_1542_v13_: (d_1540_v12_)[default__.safeIndex(294, (d_1540_v12_).length(0))]})
                nw232_[int(16)] = d_1547_v18_
                nw232_[int(17)] = d_1547_v18_
                nw232_[int(18)] = default__.fm11(-251, d_1542_v13_, (self).f10, d_1542_v13_, globalState)
                nw232_[int(19)] = d_1547_v18_
                nw232_[int(20)] = d_1547_v18_
                d_1551_v23_ = nw232_
                index247_ = default__.safeIndex(985, (d_1551_v23_).length(0))
                (d_1551_v23_)[index247_] = (d_1547_v18_ if (self).f10 else d_1547_v18_)
                d_1553_v24_: _dafny.Map
                d_1553_v24_ = _dafny.Map({(self).f10: (d_1547_v18_) | (d_1547_v18_)})
                d_1554_v25_: _dafny.Seq
                d_1554_v25_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tnx"))
                d_1555_v26_: _dafny.Set
                d_1555_v26_ = _dafny.Set({d_1554_v25_})
                index248_ = default__.safeIndex(985, (d_1551_v23_).length(0))
                rhs245_ = -464
                rhs246_ = (d_1540_v12_)[default__.safeIndex(294, (d_1540_v12_).length(0))]
                rhs247_ = ((d_1553_v24_)[not((d_1555_v26_).issubset(d_1555_v26_))] if (not((d_1555_v26_).issubset(d_1555_v26_))) in (d_1553_v24_) else _dafny.Map({d_1542_v13_: (d_1540_v12_)[default__.safeIndex(294, (d_1540_v12_).length(0))]}))
                rhs248_ = (d_1554_v25_) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xiesenboe")))
                lhs171_ = d_1551_v23_
                lhs172_ = default__.safeIndex(985, (d_1551_v23_).length(0))
                d_1542_v13_ = rhs245_
                r1 = rhs246_
                lhs171_[lhs172_] = rhs247_
                r1 = rhs248_
                d_1556_v27_: C0
                nw233_ = C0()
                nw233_.ctor__()
                d_1556_v27_ = nw233_
                d_1557_v28_: C0
                nw234_ = C0()
                nw234_.ctor__()
                d_1557_v28_ = nw234_
                (globalState).f8 = (d_1540_v12_)[default__.safeIndex(294, (d_1540_v12_).length(0))]
            elif True:
                d_1558_v29_: _dafny.Map
                d_1558_v29_ = _dafny.Map({len(default__.fm13(d_1542_v13_, globalState)): default__.fm14(globalState)})
                d_1559_v30_: str
                d_1559_v30_ = _dafny.CodePoint('q')
                d_1560_v33_: _dafny.Seq
                d_1560_v33_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qxxtkeihl"))
                d_1561_v34_: _dafny.Map
                d_1561_v34_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rvcpxn")): d_1542_v13_})
                d_1562_v35_: _dafny.Map
                d_1562_v35_ = _dafny.Map({len(d_1561_v34_): _dafny.Map({d_1542_v13_: d_1559_v30_})})
                d_1563_v36_: _dafny.Seq
                d_1563_v36_ = _dafny.SeqWithoutIsStrInference([d_1542_v13_, d_1542_v13_])
                d_1564_v37_: _dafny.Array
                nw235_ = _dafny.Array(None, 19)
                nw235_[int(0)] = (default__.fm12((d_1540_v12_)[default__.safeIndex(294, (d_1540_v12_).length(0))], d_1542_v13_, (0) - (d_1542_v13_), (d_1540_v12_)[default__.safeIndex(294, (d_1540_v12_).length(0))], globalState)) | (d_1558_v29_)
                nw235_[int(1)] = d_1558_v29_
                nw235_[int(2)] = (d_1558_v29_) | (_dafny.Map({d_1542_v13_: d_1559_v30_}))
                nw235_[int(3)] = d_1558_v29_
                nw235_[int(4)] = (d_1558_v29_) | (_dafny.Map({d_1542_v13_: d_1559_v30_}))
                nw235_[int(5)] = d_1558_v29_
                def iife133_():
                    coll57_ = _dafny.Map()
                    compr_57_: int
                    for compr_57_ in _dafny.IntegerRange(123, -659):
                        d_1565_v31_: int = compr_57_
                        if ((123) <= (d_1565_v31_)) and ((d_1565_v31_) < (-659)):
                            coll57_[(d_1565_v31_) + (d_1542_v13_)] = d_1559_v30_
                    return _dafny.Map(coll57_)
                nw235_[int(6)] = iife133_()
                
                nw235_[int(7)] = (d_1558_v29_) | (d_1558_v29_)
                def iife134_():
                    coll58_ = _dafny.Map()
                    compr_58_: int
                    for compr_58_ in (_dafny.SeqWithoutIsStrInference([360 for d_1566_i4_ in range(default__.abs(528))])).Elements:
                        d_1567_v32_: int = compr_58_
                        if (d_1567_v32_) in (_dafny.SeqWithoutIsStrInference([360 for d_1566_i4_ in range(default__.abs(528))])):
                            coll58_[(d_1567_v32_) * ((_dafny.MultiSet([(self).f10])).cardinality)] = d_1559_v30_
                    return _dafny.Map(coll58_)
                nw235_[int(8)] = iife134_()
                
                nw235_[int(9)] = d_1558_v29_
                nw235_[int(10)] = (default__.fm12(not((self).f10), d_1542_v13_, len(d_1560_v33_), (self).f10, globalState)) | (d_1558_v29_)
                nw235_[int(11)] = d_1558_v29_
                nw235_[int(12)] = (d_1558_v29_) | (_dafny.Map({d_1542_v13_: d_1559_v30_}))
                nw235_[int(13)] = ((d_1558_v29_).set(d_1542_v13_, d_1559_v30_)) | (((d_1562_v35_)[d_1542_v13_] if (d_1542_v13_) in (d_1562_v35_) else default__.fm12(True, d_1542_v13_, d_1542_v13_, (self).f10, globalState)))
                nw235_[int(14)] = default__.fm12((self).f10, len(d_1560_v33_), len(d_1563_v36_), (self).f10, globalState)
                nw235_[int(15)] = _dafny.Map({d_1542_v13_: d_1559_v30_})
                nw235_[int(16)] = d_1558_v29_
                nw235_[int(17)] = d_1558_v29_
                nw235_[int(18)] = d_1558_v29_
                d_1564_v37_ = nw235_
                d_1564_v37_ = d_1564_v37_
                d_1560_v33_ = d_1560_v33_
                d_1568_v38_: _dafny.Array
                def lambda70_(d_1569_v13_):
                    def lambda71_(d_1570_i5_):
                        return (d_1570_i5_) - (d_1569_v13_)

                    return lambda71_

                init42_ = lambda70_(d_1542_v13_)
                nw236_ = _dafny.Array(None, 25)
                for i0_42_ in range(nw236_.length(0)):
                    nw236_[i0_42_] = init42_(i0_42_)
                d_1568_v38_ = nw236_
                index249_ = default__.safeIndex(758, (d_1568_v38_).length(0))
                (d_1568_v38_)[index249_] = 940
                index250_ = default__.safeIndex(758, (d_1568_v38_).length(0))
                (d_1568_v38_)[index250_] = (607) - (d_1542_v13_)
                (globalState).f8 = (d_1540_v12_)[default__.safeIndex(294, (d_1540_v12_).length(0))]
                d_1571_v39_: _dafny.Map
                d_1571_v39_ = _dafny.Map({696: (d_1540_v12_)[default__.safeIndex(294, (d_1540_v12_).length(0))]})
                d_1572_v40_: _dafny.Map
                d_1572_v40_ = _dafny.Map({(self).f10: default__.fm13(d_1542_v13_, globalState)})
                rhs249_ = not (((self).f10) or ((d_1540_v12_)[default__.safeIndex(294, (d_1540_v12_).length(0))])) or ((self).fm9((d_1568_v38_)[default__.safeIndex(758, (d_1568_v38_).length(0))], d_1542_v13_, d_1542_v13_, d_1572_v40_, globalState))
                rhs250_ = (d_1571_v39_ if (self).f10 else d_1571_v39_)
                lhs173_ = globalState
                lhs173_.f8 = rhs249_
                d_1571_v39_ = rhs250_
            (globalState).f8 = not (not((d_1540_v12_)[default__.safeIndex(294, (d_1540_v12_).length(0))])) or ((self).f10)
        d_1573_v41_: _dafny.Seq
        d_1573_v41_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fuy"))
        (globalState).f8 = (d_1573_v41_) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d")))
        d_1574_v42_: int
        d_1574_v42_ = -36
        d_1575_v43_: _dafny.MultiSet
        d_1575_v43_ = _dafny.MultiSet([d_1574_v42_, 94])
        d_1576_v44_: _dafny.MultiSet
        d_1576_v44_ = _dafny.MultiSet([d_1575_v43_, d_1575_v43_])
        d_1577_i6_: int
        d_1577_i6_ = 0
        with _dafny.label("7"):
            while (d_1576_v44_).issubset(d_1576_v44_):
                with _dafny.c_label("7"):
                    if (d_1577_i6_) >= (100):
                        raise _dafny.Break("7")
                    d_1577_i6_ = (d_1577_i6_) + (1)
                    d_1578_v45_: C0
                    nw237_ = C0()
                    nw237_.ctor__()
                    d_1578_v45_ = nw237_
                    d_1574_v42_ = d_1574_v42_
                    d_1579_v46_: _dafny.MultiSet
                    d_1579_v46_ = _dafny.MultiSet([(self).f10, not(False), (self).f10, (self).f10, (self).f10])
                    d_1580_v47_: _dafny.Set
                    d_1580_v47_ = _dafny.Set({(self).f10})
                    d_1581_v48_: _dafny.Seq
                    d_1581_v48_ = _dafny.SeqWithoutIsStrInference([len(d_1580_v47_)])
                    d_1582_v49_: _dafny.Map
                    d_1582_v49_ = _dafny.Map({d_1574_v42_: (d_1581_v48_)[default__.safeIndex(806, len(d_1581_v48_))]})
                    d_1583_v50_: _dafny.Map
                    d_1583_v50_ = _dafny.Map({((self).f10) in (d_1579_v46_): (len(d_1573_v41_)) * (((d_1582_v49_)[len(_dafny.SeqWithoutIsStrInference([d_1574_v42_, 296]))] if (len(_dafny.SeqWithoutIsStrInference([d_1574_v42_, 296]))) in (d_1582_v49_) else d_1574_v42_))})
                    d_1584_v51_: _dafny.Set
                    d_1584_v51_ = _dafny.Set({863})
                    d_1574_v42_ = ((d_1583_v50_)[(self).f10] if ((self).f10) in (d_1583_v50_) else (d_1574_v42_) * (len(d_1584_v51_)))
                    d_1585_v52_: _dafny.Map
                    d_1585_v52_ = _dafny.Map({(self).f10: d_1573_v41_})
                    d_1586_v53_: _dafny.Array
                    nw238_ = _dafny.Array(None, 25)
                    nw238_[int(0)] = (self).f10
                    nw238_[int(1)] = (self).f10
                    nw238_[int(2)] = (self).f10
                    nw238_[int(3)] = True
                    nw238_[int(4)] = (self).f10
                    nw238_[int(5)] = (self).f10
                    nw238_[int(6)] = (self).f10
                    nw238_[int(7)] = not((self).f10)
                    nw238_[int(8)] = (self).f10
                    nw238_[int(9)] = (self).f10
                    nw238_[int(10)] = True
                    nw238_[int(11)] = (self).f10
                    nw238_[int(12)] = (self).f10
                    nw238_[int(13)] = (self).f10
                    nw238_[int(14)] = (self).f10
                    nw238_[int(15)] = (self).f10
                    nw238_[int(16)] = (self).f10
                    nw238_[int(17)] = not((self).f10)
                    nw238_[int(18)] = (self).f10
                    nw238_[int(19)] = (self).fm9((self).fm1((self).f9, globalState), d_1574_v42_, d_1574_v42_, d_1585_v52_, globalState)
                    nw238_[int(20)] = (self).f10
                    nw238_[int(21)] = (self).f10
                    nw238_[int(22)] = not((self).f10)
                    nw238_[int(23)] = (self).f10
                    nw238_[int(24)] = (self).f10
                    d_1586_v53_ = nw238_
                    d_1587_v54_: _dafny.Seq
                    d_1587_v54_ = _dafny.SeqWithoutIsStrInference([d_1586_v53_, d_1586_v53_])
                    d_1588_v55_: _dafny.MultiSet
                    d_1588_v55_ = _dafny.MultiSet([(d_1587_v54_)[default__.safeIndex(len(d_1581_v48_), len(d_1587_v54_))], (d_1586_v53_ if (self).f10 else (d_1587_v54_)[default__.safeIndex(len(_dafny.Map({d_1574_v42_: d_1582_v49_})), len(d_1587_v54_))]), d_1586_v53_, d_1586_v53_, d_1586_v53_])
                    rhs251_ = d_1588_v55_
                    rhs252_ = (d_1574_v42_) != (d_1574_v42_)
                    lhs174_ = globalState
                    d_1588_v55_ = rhs251_
                    lhs174_.f8 = rhs252_
                    pass
            pass
        d_1589_v56_: _dafny.Array
        nw239_ = _dafny.Array(False, 10)
        d_1589_v56_ = nw239_
        guard_loop_6_: int
        for guard_loop_6_ in _dafny.IntegerRange(0, (d_1589_v56_).length(0)):
            d_1590_i7_: int = guard_loop_6_
            if (True) and (((0) <= (d_1590_i7_)) and ((d_1590_i7_) < ((d_1589_v56_).length(0)))):
                (d_1589_v56_)[(d_1590_i7_)] = (len((_dafny.Map({not(not((self).f10)): d_1574_v42_})) | (_dafny.Map({not((self).f10): d_1574_v42_})))) <= (939)
        d_1591_v57_: _dafny.Set
        d_1591_v57_ = _dafny.Set({d_1574_v42_})
        d_1592_v58_: _dafny.MultiSet
        d_1592_v58_ = _dafny.MultiSet([not((self).f10), (self).f10])
        d_1593_v59_: _dafny.Seq
        d_1593_v59_ = _dafny.SeqWithoutIsStrInference([d_1592_v58_, d_1592_v58_, d_1592_v58_])
        d_1594_v60_: _dafny.Map
        d_1594_v60_ = _dafny.Map({d_1591_v57_: d_1593_v59_})
        def iife135_():
            coll59_ = _dafny.Set()
            compr_59_: int
            for compr_59_ in (_dafny.SeqWithoutIsStrInference([d_1574_v42_])).Elements:
                d_1595_v61_: int = compr_59_
                if (d_1595_v61_) in (_dafny.SeqWithoutIsStrInference([d_1574_v42_])):
                    coll59_ = coll59_.union(_dafny.Set([default__.safeDivisionInt(d_1595_v61_, len(_dafny.SeqWithoutIsStrInference([True, True, False, False])))]))
            return _dafny.Set(coll59_)
        d_1594_v60_ = (d_1594_v60_).set((d_1591_v57_ if (self).f10 else iife135_()
        ), (_dafny.SeqWithoutIsStrInference([d_1592_v58_])) + (d_1593_v59_))
        d_1596_v62_: D0
        d_1596_v62_ = D0_DC1(default__.safeModuloInt(765, d_1574_v42_), (self).f10, d_1574_v42_)
        source22_ = d_1596_v62_
        if source22_.is_DC1:
            d_1597___mcc_h0_ = source22_.cf1
            d_1598___mcc_h1_ = source22_.cf2
            d_1599___mcc_h2_ = source22_.cf3
            d_1600_cf3_ = d_1599___mcc_h2_
            d_1601_cf2_ = d_1598___mcc_h1_
            d_1602_cf1_ = d_1597___mcc_h0_
            d_1603_v63_: _dafny.Array
            nw240_ = _dafny.Array(int(0), 4)
            d_1603_v63_ = nw240_
            d_1604_v64_: D1
            d_1604_v64_ = D1_DC4(False, d_1601_cf2_, len(d_1573_v41_))
            index251_ = default__.safeIndex(44, (d_1603_v63_).length(0))
            (d_1603_v63_)[index251_] = (d_1604_v64_).cf8
            index252_ = default__.safeIndex(44, (d_1603_v63_).length(0))
            rhs253_ = d_1600_cf3_
            rhs254_ = (self).f10
            rhs255_ = d_1574_v42_
            lhs175_ = d_1603_v63_
            lhs176_ = default__.safeIndex(44, (d_1603_v63_).length(0))
            d_1574_v42_ = rhs253_
            d_1601_cf2_ = rhs254_
            lhs175_[lhs176_] = rhs255_
            d_1605_v65_: _dafny.Seq
            d_1605_v65_ = _dafny.SeqWithoutIsStrInference([d_1573_v41_])
            d_1606_v66_: D1
            d_1606_v66_ = D1_DC3((d_1605_v65_)[default__.safeIndex(d_1574_v42_, len(d_1605_v65_))])
            d_1606_v66_ = d_1606_v66_
            d_1589_v56_ = d_1589_v56_
            d_1601_cf2_ = d_1601_cf2_
        elif True:
            d_1607___mcc_h3_ = source22_.cf0
            d_1608_cf0_ = d_1607___mcc_h3_
            d_1609_v67_: _dafny.Array
            def lambda72_(d_1610_v58_):
                def lambda73_(d_1611_i8_):
                    return (d_1610_v58_).intersection(d_1610_v58_)

                return lambda73_

            init43_ = lambda72_(d_1592_v58_)
            nw241_ = _dafny.Array(None, 16)
            for i0_43_ in range(nw241_.length(0)):
                nw241_[i0_43_] = init43_(i0_43_)
            d_1609_v67_ = nw241_
            index253_ = default__.safeIndex(723, (d_1609_v67_).length(0))
            (d_1609_v67_)[index253_] = (d_1592_v58_) | (d_1592_v58_)
            index254_ = default__.safeIndex(723, (d_1609_v67_).length(0))
            def iife136_():
                coll60_ = _dafny.Map()
                compr_60_: int
                for compr_60_ in _dafny.IntegerRange(-192, 379):
                    d_1612_v68_: int = compr_60_
                    if ((-192) <= (d_1612_v68_)) and ((d_1612_v68_) < (379)):
                        coll60_[(d_1612_v68_) + (d_1574_v42_)] = d_1608_cf0_
                return _dafny.Map(coll60_)
            (d_1609_v67_)[index254_] = default__.fm15((self).f10, (self).f10, (d_1574_v42_) * (len(iife136_()
            )), _dafny.CodePoint('k'), globalState)
            d_1613_v69_: C0
            nw242_ = C0()
            nw242_.ctor__()
            d_1613_v69_ = nw242_
            d_1614_v70_: _dafny.Array
            def lambda74_(d_1615_cf0_, d_1616_v62_):
                def lambda75_(d_1617_i9_):
                    return (d_1616_v62_ if d_1615_cf0_ else d_1616_v62_)

                return lambda75_

            init44_ = lambda74_(d_1608_cf0_, d_1596_v62_)
            nw243_ = _dafny.Array(None, 25)
            for i0_44_ in range(nw243_.length(0)):
                nw243_[i0_44_] = init44_(i0_44_)
            d_1614_v70_ = nw243_
            index255_ = default__.safeIndex(108, (d_1614_v70_).length(0))
            (d_1614_v70_)[index255_] = d_1596_v62_
            index256_ = default__.safeIndex(108, (d_1614_v70_).length(0))
            (d_1614_v70_)[index256_] = d_1596_v62_
            index257_ = default__.safeIndex(767, (d_1589_v56_).length(0))
            (d_1589_v56_)[index257_] = (self).f10
            index258_ = default__.safeIndex(767, (d_1589_v56_).length(0))
            (d_1589_v56_)[index258_] = True
        r0 = ((d_1573_v41_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_1618_i10_ in range(default__.abs(363))]))).set(default__.safeIndex((d_1574_v42_ if True else (0) - (d_1574_v42_)), len((d_1573_v41_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_1619_i10_ in range(default__.abs(363))])))), (d_1573_v41_)[default__.safeIndex(d_1574_v42_, len(d_1573_v41_))])
        r1 = (self).f10
        d_1620_v71_: _dafny.Set
        d_1620_v71_ = _dafny.Set({(self).f10})
        d_1621_v72_: _dafny.Seq
        d_1621_v72_ = _dafny.SeqWithoutIsStrInference([d_1620_v71_])
        d_1622_v73_: _dafny.Map
        d_1622_v73_ = _dafny.Map({(self).f10: d_1621_v72_})
        d_1623_v74_: _dafny.Seq
        d_1623_v74_ = _dafny.SeqWithoutIsStrInference([d_1574_v42_])
        r2 = ((d_1622_v73_)[(self).f10] if ((self).f10) in (d_1622_v73_) else default__.fm16(len(d_1623_v74_), d_1574_v42_, 913, globalState))
        r3 = d_1573_v41_
        return r0, r1, r2, r3

    def m1(self, p0, globalState):
        r0: _dafny.Map = _dafny.Map({})
        r1: bool = False
        d_1624_v0_: C0
        nw244_ = C0()
        nw244_.ctor__()
        d_1624_v0_ = nw244_
        d_1625_v1_: int
        d_1625_v1_ = 699
        d_1626_v2_: D1
        d_1626_v2_ = D1_DC4((self).f10, (self).f10, d_1625_v1_)
        def lambda76_(source23_):
            if source23_.is_DC3:
                d_1627___mcc_h0_ = source23_.cf5
                d_1628_cf5_ = d_1627___mcc_h0_
                return (self).f10
            elif source23_.is_DC4:
                d_1629___mcc_h1_ = source23_.cf6
                d_1630___mcc_h2_ = source23_.cf7
                d_1631___mcc_h3_ = source23_.cf8
                d_1632_cf8_ = d_1631___mcc_h3_
                d_1633_cf7_ = d_1630___mcc_h2_
                d_1634_cf6_ = d_1629___mcc_h1_
                return (_dafny.MultiSet([d_1634_cf6_, (self).f10, (self).f10, True])).isdisjoint(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True])))
            elif source23_.is_DC2:
                d_1635___mcc_h4_ = source23_.cf4
                d_1636_cf4_ = d_1635___mcc_h4_
                return (self).f10
            elif True:
                d_1637___mcc_h5_ = source23_.cf9
                d_1638_cf9_ = d_1637___mcc_h5_
                return (self).f10

        r1 = lambda76_(D1_DC5(d_1626_v2_))
        d_1639_v3_: _dafny.MultiSet
        d_1639_v3_ = _dafny.MultiSet([d_1625_v1_, d_1625_v1_])
        d_1640_v4_: _dafny.Seq
        d_1640_v4_ = _dafny.SeqWithoutIsStrInference([d_1625_v1_])
        d_1641_v5_: D1
        d_1641_v5_ = D1_DC4((self).f10, (self).f10, d_1625_v1_)
        d_1642_v6_: _dafny.Seq
        d_1642_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qfpc"))
        d_1643_v7_: _dafny.Set
        d_1643_v7_ = _dafny.Set({(self).f10})
        d_1644_v8_: _dafny.Map
        d_1644_v8_ = _dafny.Map({d_1625_v1_: (self).f10})
        d_1645_v9_: _dafny.Array
        nw245_ = _dafny.Array(None, 19)
        nw245_[int(0)] = -961
        nw245_[int(1)] = d_1625_v1_
        nw245_[int(2)] = default__.safeDivisionInt(810, d_1625_v1_)
        nw245_[int(3)] = ((d_1639_v3_) - (d_1639_v3_)).cardinality
        nw245_[int(4)] = 139
        nw245_[int(5)] = len((d_1640_v4_) + (d_1640_v4_))
        nw245_[int(6)] = (-855) + ((d_1641_v5_).cf8)
        nw245_[int(7)] = len(d_1642_v6_)
        nw245_[int(8)] = d_1625_v1_
        nw245_[int(9)] = len(d_1643_v7_)
        nw245_[int(10)] = (-393) - (d_1625_v1_)
        nw245_[int(11)] = d_1625_v1_
        nw245_[int(12)] = (0) - (default__.safeModuloInt(d_1625_v1_, d_1625_v1_))
        nw245_[int(13)] = default__.safeDivisionInt(869, d_1625_v1_)
        nw245_[int(14)] = d_1625_v1_
        nw245_[int(15)] = d_1625_v1_
        nw245_[int(16)] = d_1625_v1_
        nw245_[int(17)] = len(((d_1642_v6_).set(default__.safeIndex(len(d_1644_v8_), len(d_1642_v6_)), p0)) + (d_1642_v6_))
        nw245_[int(18)] = d_1625_v1_
        d_1645_v9_ = nw245_
        guard_loop_7_: int
        for guard_loop_7_ in _dafny.IntegerRange(0, (d_1645_v9_).length(0)):
            d_1646_i0_: int = guard_loop_7_
            if (True) and (((0) <= (d_1646_i0_)) and ((d_1646_i0_) < ((d_1645_v9_).length(0)))):
                (d_1645_v9_)[(d_1646_i0_)] = (d_1646_i0_) + (d_1625_v1_)
        hi4_ = (0) - (d_1625_v1_)
        for d_1647_i1_ in range((d_1625_v1_) - (d_1625_v1_), hi4_):
            (globalState).f8 = True
            d_1648_v10_: C0
            nw246_ = C0()
            nw246_.ctor__()
            d_1648_v10_ = nw246_
            d_1649_v11_: _dafny.Map
            d_1649_v11_ = _dafny.Map({d_1645_v9_: d_1642_v6_})
            d_1650_v12_: _dafny.Array
            nw247_ = _dafny.Array(None, 25)
            nw247_[int(0)] = (d_1642_v6_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_1651_i2_ in range(default__.abs(877))]))
            nw247_[int(1)] = d_1642_v6_
            nw247_[int(2)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "amb"))) + (d_1642_v6_)
            nw247_[int(3)] = d_1642_v6_
            nw247_[int(4)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xmfr"))) + ((d_1642_v6_).set(default__.safeIndex(d_1625_v1_, len(d_1642_v6_)), p0))
            nw247_[int(5)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bk"))
            nw247_[int(6)] = ((d_1642_v6_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bac")))).set(default__.safeIndex(d_1625_v1_, len((d_1642_v6_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bac"))))), p0)
            nw247_[int(7)] = (d_1642_v6_) + (d_1642_v6_)
            nw247_[int(8)] = d_1642_v6_
            nw247_[int(9)] = d_1642_v6_
            nw247_[int(10)] = d_1642_v6_
            nw247_[int(11)] = d_1642_v6_
            nw247_[int(12)] = d_1642_v6_
            nw247_[int(13)] = d_1642_v6_
            nw247_[int(14)] = d_1642_v6_
            nw247_[int(15)] = ((d_1649_v11_)[d_1645_v9_] if (d_1645_v9_) in (d_1649_v11_) else d_1642_v6_)
            nw247_[int(16)] = default__.fm13(592, globalState)
            nw247_[int(17)] = (d_1642_v6_) + (d_1642_v6_)
            nw247_[int(18)] = (d_1642_v6_) + (d_1642_v6_)
            nw247_[int(19)] = (d_1642_v6_) + (_dafny.SeqWithoutIsStrInference([(d_1642_v6_)[default__.safeIndex(len(_dafny.Set({d_1625_v1_, d_1625_v1_})), len(d_1642_v6_))] for d_1652_i3_ in range(default__.abs(644))]))
            nw247_[int(20)] = d_1642_v6_
            nw247_[int(21)] = (d_1642_v6_) + (d_1642_v6_)
            nw247_[int(22)] = (d_1642_v6_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "om")))
            nw247_[int(23)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lsvc"))
            nw247_[int(24)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ldunu"))
            d_1650_v12_ = nw247_
            index259_ = default__.safeIndex(106, (d_1650_v12_).length(0))
            (d_1650_v12_)[index259_] = _dafny.SeqWithoutIsStrInference([p0 for d_1653_i4_ in range(default__.abs(786))])
            index260_ = default__.safeIndex(757, (d_1645_v9_).length(0))
            (d_1645_v9_)[index260_] = d_1647_i1_
            index261_ = default__.safeIndex(106, (d_1650_v12_).length(0))
            index262_ = default__.safeIndex(757, (d_1645_v9_).length(0))
            rhs256_ = d_1625_v1_
            rhs257_ = (d_1642_v6_) + (((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gpkthj"))) + (d_1642_v6_)).set(default__.safeIndex(d_1625_v1_, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gpkthj"))) + (d_1642_v6_))), _dafny.CodePoint('i')))
            rhs258_ = (self).f10
            rhs259_ = (-553) * ((d_1647_i1_) * (583))
            lhs177_ = d_1650_v12_
            lhs178_ = default__.safeIndex(106, (d_1650_v12_).length(0))
            lhs179_ = d_1645_v9_
            lhs180_ = default__.safeIndex(757, (d_1645_v9_).length(0))
            d_1625_v1_ = rhs256_
            lhs177_[lhs178_] = rhs257_
            r1 = rhs258_
            lhs179_[lhs180_] = rhs259_
            d_1654_v13_: C0
            nw248_ = C0()
            nw248_.ctor__()
            d_1654_v13_ = nw248_
        d_1655_v14_: str
        d_1655_v14_ = _dafny.CodePoint('j')
        d_1655_v14_ = d_1655_v14_
        d_1656_v15_: D1
        d_1656_v15_ = D1_DC2((d_1640_v4_).set(default__.safeIndex(d_1625_v1_, len(d_1640_v4_)), (0) - (d_1625_v1_)))
        source24_ = d_1656_v15_
        if source24_.is_DC3:
            d_1657___mcc_h6_ = source24_.cf5
            d_1658_cf5_ = d_1657___mcc_h6_
            d_1659_v16_: _dafny.Map
            d_1659_v16_ = _dafny.Map({_dafny.CodePoint('d'): (self).f10})
            d_1660_v17_: _dafny.Seq
            d_1660_v17_ = _dafny.SeqWithoutIsStrInference([(self).f10])
            d_1640_v4_ = ((_dafny.SeqWithoutIsStrInference([(0) - (d_1625_v1_) for d_1661_i5_ in range(default__.abs(149))])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(self).f10, (self).f10]))])) if (((self).f9)[(self).f10] if ((self).f10) in ((self).f9) else not(((d_1659_v16_)[d_1655_v14_] if (d_1655_v14_) in (d_1659_v16_) else (d_1660_v17_)[default__.safeIndex(d_1625_v1_, len(d_1660_v17_))]))) else d_1640_v4_)
            index263_ = default__.safeIndex(393, (d_1645_v9_).length(0))
            (d_1645_v9_)[index263_] = d_1625_v1_
            index264_ = default__.safeIndex(393, (d_1645_v9_).length(0))
            (d_1645_v9_)[index264_] = (0) - ((0) - (d_1625_v1_))
            d_1644_v8_ = d_1644_v8_
            d_1662_v18_: _dafny.Array
            nw249_ = _dafny.Array(None, 4)
            nw249_[int(0)] = (self).f10
            nw249_[int(1)] = (self).f10
            nw249_[int(2)] = (True if (self).f10 else (self).f10)
            nw249_[int(3)] = (-255) <= (528)
            d_1662_v18_ = nw249_
            index265_ = default__.safeIndex(158, (d_1662_v18_).length(0))
            (d_1662_v18_)[index265_] = True
            d_1663_v19_: _dafny.Set
            d_1663_v19_ = _dafny.Set({d_1625_v1_})
            index266_ = default__.safeIndex(158, (d_1662_v18_).length(0))
            (d_1662_v18_)[index266_] = (default__.fm17(True, globalState)).issubset((d_1663_v19_) - (d_1663_v19_))
        elif source24_.is_DC4:
            d_1664___mcc_h7_ = source24_.cf6
            d_1665___mcc_h8_ = source24_.cf7
            d_1666___mcc_h9_ = source24_.cf8
            d_1667_cf8_ = d_1666___mcc_h9_
            d_1668_cf7_ = d_1665___mcc_h8_
            d_1669_cf6_ = d_1664___mcc_h7_
            d_1670_v20_: D1
            d_1670_v20_ = D1_DC5(d_1626_v2_)
            d_1671_v21_: _dafny.Map
            d_1671_v21_ = _dafny.Map({d_1670_v20_: D0_DC1((0) - (d_1625_v1_), d_1668_cf7_, d_1625_v1_)})
            d_1672_v22_: D0
            d_1672_v22_ = D0_DC0(d_1668_cf7_)
            d_1673_v23_: _dafny.Seq
            d_1673_v23_ = _dafny.SeqWithoutIsStrInference([d_1672_v22_, d_1672_v22_])
            d_1674_v24_: _dafny.Map
            d_1674_v24_ = _dafny.Map({d_1668_cf7_: -583})
            d_1675_v25_: _dafny.Seq
            d_1675_v25_ = _dafny.SeqWithoutIsStrInference([(self).f10])
            rhs260_ = (d_1625_v1_) > (len((d_1673_v23_) + (d_1673_v23_)))
            rhs261_ = d_1671_v21_
            rhs262_ = (d_1667_cf8_ if d_1669_cf6_ else ((d_1674_v24_)[(d_1675_v25_)[default__.safeIndex(d_1625_v1_, len(d_1675_v25_))]] if ((d_1675_v25_)[default__.safeIndex(d_1625_v1_, len(d_1675_v25_))]) in (d_1674_v24_) else len(d_1643_v7_)))
            d_1668_cf7_ = rhs260_
            d_1671_v21_ = rhs261_
            d_1667_cf8_ = rhs262_
            d_1676_v26_: C0
            nw250_ = C0()
            nw250_.ctor__()
            d_1676_v26_ = nw250_
            d_1677_v27_: _dafny.Seq
            d_1677_v27_ = _dafny.SeqWithoutIsStrInference([d_1645_v9_])
            d_1667_cf8_ = (len(d_1677_v27_)) + (d_1625_v1_)
            d_1678_v28_: str
            d_1678_v28_ = d_1655_v14_
            d_1679_v29_: D1
            d_1679_v29_ = D1_DC3((d_1642_v6_).set(default__.safeIndex(d_1625_v1_, len(d_1642_v6_)), (d_1678_v28_)))
            source25_ = d_1679_v29_
            if source25_.is_DC3:
                d_1680___mcc_h12_ = source25_.cf5
                d_1681_cf5_ = d_1680___mcc_h12_
                d_1682_v30_: C0
                nw251_ = C0()
                nw251_.ctor__()
                d_1682_v30_ = nw251_
                d_1678_v28_ = d_1655_v14_
                d_1683_v31_: _dafny.Array
                def lambda77_(d_1684_i6_):
                    return (self).f10

                init45_ = lambda77_
                nw252_ = _dafny.Array(None, 26)
                for i0_45_ in range(nw252_.length(0)):
                    nw252_[i0_45_] = init45_(i0_45_)
                d_1683_v31_ = nw252_
                index267_ = default__.safeIndex(138, (d_1683_v31_).length(0))
                (d_1683_v31_)[index267_] = d_1668_cf7_
                d_1685_v32_: _dafny.Map
                d_1685_v32_ = _dafny.Map({d_1669_cf6_: d_1645_v9_})
                index268_ = default__.safeIndex(138, (d_1683_v31_).length(0))
                rhs263_ = d_1669_cf6_
                rhs264_ = ((d_1685_v32_)[False] if (False) in (d_1685_v32_) else (d_1677_v27_)[default__.safeIndex(d_1625_v1_, len(d_1677_v27_))])
                rhs265_ = len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fnhqhubuo"))) + (d_1642_v6_))
                rhs266_ = default__.fm14(globalState)
                lhs181_ = d_1683_v31_
                lhs182_ = default__.safeIndex(138, (d_1683_v31_).length(0))
                lhs181_[lhs182_] = rhs263_
                d_1645_v9_ = rhs264_
                d_1667_cf8_ = rhs265_
                d_1655_v14_ = rhs266_
                d_1644_v8_ = d_1644_v8_
            elif source25_.is_DC4:
                d_1686___mcc_h13_ = source25_.cf6
                d_1687___mcc_h14_ = source25_.cf7
                d_1688___mcc_h15_ = source25_.cf8
                d_1689_cf8_ = d_1688___mcc_h15_
                d_1690_cf7_ = d_1687___mcc_h14_
                d_1691_cf6_ = d_1686___mcc_h13_
                d_1644_v8_ = (d_1644_v8_).set(default__.safeModuloInt(d_1689_cf8_, 472), d_1691_cf6_)
                d_1692_v33_: _dafny.Set
                d_1692_v33_ = _dafny.Set({len(_dafny.Set({d_1674_v24_}))})
                d_1693_v43_: _dafny.Map
                d_1693_v43_ = _dafny.Map({d_1690_cf7_: d_1624_v0_})
                d_1694_v44_: _dafny.Seq
                def iife137_():
                    coll61_ = _dafny.Set()
                    compr_61_: int
                    for compr_61_ in (d_1640_v4_).Elements:
                        d_1695_v41_: int = compr_61_
                        if (d_1695_v41_) in (d_1640_v4_):
                            coll61_ = coll61_.union(_dafny.Set([(d_1695_v41_) - (754)]))
                    return _dafny.Set(coll61_)
                def iife138_():
                    coll62_ = _dafny.Set()
                    compr_62_: int
                    for compr_62_ in _dafny.IntegerRange(487, 259):
                        d_1696_v42_: int = compr_62_
                        if ((487) <= (d_1696_v42_)) and ((d_1696_v42_) < (259)):
                            coll62_ = coll62_.union(_dafny.Set([(d_1696_v42_) - (len(d_1642_v6_))]))
                    return _dafny.Set(coll62_)
                d_1694_v44_ = _dafny.SeqWithoutIsStrInference([iife137_()
                , d_1692_v33_, iife138_()
                , d_1692_v33_, _dafny.Set({d_1689_cf8_, d_1689_cf8_, -918, d_1625_v1_, len(d_1693_v43_)})])
                d_1697_v45_: _dafny.Array
                nw253_ = _dafny.Array(None, 13)
                nw253_[int(0)] = d_1692_v33_
                nw253_[int(1)] = d_1692_v33_
                nw253_[int(2)] = d_1692_v33_
                def iife139_():
                    coll63_ = _dafny.Set()
                    compr_63_: int
                    for compr_63_ in _dafny.IntegerRange(600, -636):
                        d_1698_v34_: int = compr_63_
                        if ((600) <= (d_1698_v34_)) and ((d_1698_v34_) < (-636)):
                            def iife140_():
                                def iife142_():
                                    coll66_ = _dafny.Map()
                                    compr_66_: _dafny.Seq
                                    for compr_66_ in (_dafny.MultiSet([d_1642_v6_])).Elements:
                                        d_1699_v36_: _dafny.Seq = compr_66_
                                        if (d_1699_v36_) in (_dafny.MultiSet([d_1642_v6_])):
                                            coll66_[d_1699_v36_] = len(d_1642_v6_)
                                    return _dafny.Map(coll66_)
                                coll64_ = _dafny.Map()
                                def iife141_():
                                    coll65_ = _dafny.Map()
                                    compr_65_: _dafny.Seq
                                    for compr_65_ in (_dafny.MultiSet([d_1642_v6_])).Elements:
                                        d_1699_v36_: _dafny.Seq = compr_65_
                                        if (d_1699_v36_) in (_dafny.MultiSet([d_1642_v6_])):
                                            coll65_[d_1699_v36_] = len(d_1642_v6_)
                                    return _dafny.Map(coll65_)
                                compr_64_: _dafny.Seq
                                for compr_64_ in (iife141_()
                                ).keys.Elements:
                                    d_1700_v35_: _dafny.Seq = compr_64_
                                    if (d_1700_v35_) in (iife142_()
                                    ):
                                        coll64_[d_1700_v35_] = d_1691_cf6_
                                return _dafny.Map(coll64_)
                            coll63_ = coll63_.union(_dafny.Set([default__.safeModuloInt(d_1698_v34_, len(iife140_()
))]))
                    return _dafny.Set(coll63_)
                nw253_[int(3)] = iife139_()
                
                def iife143_():
                    coll67_ = _dafny.Set()
                    compr_67_: int
                    for compr_67_ in (d_1639_v3_).Elements:
                        d_1701_v37_: int = compr_67_
                        if (d_1701_v37_) in (d_1639_v3_):
                            coll67_ = coll67_.union(_dafny.Set([(d_1701_v37_) * (485)]))
                    return _dafny.Set(coll67_)
                nw253_[int(4)] = iife143_()
                
                nw253_[int(5)] = (d_1692_v33_ if d_1690_cf7_ else d_1692_v33_)
                nw253_[int(6)] = d_1692_v33_
                nw253_[int(7)] = d_1692_v33_
                nw253_[int(8)] = d_1692_v33_
                def iife144_():
                    coll68_ = _dafny.Set()
                    compr_68_: int
                    for compr_68_ in (d_1640_v4_).Elements:
                        d_1702_v38_: int = compr_68_
                        if (d_1702_v38_) in (d_1640_v4_):
                            def iife145_():
                                coll69_ = _dafny.Map()
                                compr_69_: int
                                for compr_69_ in _dafny.IntegerRange(305, 277):
                                    d_1703_v39_: int = compr_69_
                                    if ((305) <= (d_1703_v39_)) and ((d_1703_v39_) < (277)):
                                        def iife146_():
                                            coll70_ = _dafny.Set()
                                            compr_70_: int
                                            for compr_70_ in _dafny.IntegerRange(390, 962):
                                                d_1704_v40_: int = compr_70_
                                                if ((390) <= (d_1704_v40_)) and ((d_1704_v40_) < (962)):
                                                    coll70_ = coll70_.union(_dafny.Set([(d_1704_v40_) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h"))))]))
                                            return _dafny.Set(coll70_)
                                        coll69_[default__.safeModuloInt(d_1703_v39_, (0) - (len(_dafny.SeqWithoutIsStrInference([809 for d_1705_i7_ in range(default__.abs(330))]))))] = iife146_()

                                return _dafny.Map(coll69_)
                            coll68_ = coll68_.union(_dafny.Set([default__.safeDivisionInt(d_1702_v38_, len(iife145_()
))]))
                    return _dafny.Set(coll68_)
                nw253_[int(9)] = iife144_()
                
                nw253_[int(10)] = (d_1694_v44_)[default__.safeIndex(307, len(d_1694_v44_))]
                nw253_[int(11)] = d_1692_v33_
                nw253_[int(12)] = default__.fm17((self).f10, globalState)
                d_1697_v45_ = nw253_
                index269_ = default__.safeIndex(739, (d_1697_v45_).length(0))
                (d_1697_v45_)[index269_] = d_1692_v33_
                index270_ = default__.safeIndex(739, (d_1697_v45_).length(0))
                (d_1697_v45_)[index270_] = d_1692_v33_
                d_1706_v46_: _dafny.Map
                d_1706_v46_ = _dafny.Map({-407: d_1689_cf8_})
                d_1707_v47_: T1
                nw254_ = C4()
                nw254_.ctor__((self).f9, d_1691_cf6_)
                d_1707_v47_ = nw254_
                d_1708_v48_: _dafny.Seq
                d_1708_v48_ = _dafny.SeqWithoutIsStrInference([d_1707_v47_])
                d_1709_v49_: _dafny.Map
                d_1709_v49_ = _dafny.Map({d_1691_cf6_: _dafny.MultiSet((d_1708_v48_).set(default__.safeIndex(d_1625_v1_, len(d_1708_v48_)), d_1707_v47_))})
                d_1706_v46_ = (d_1706_v46_).set((d_1625_v1_) + (d_1689_cf8_), ((0) - (d_1689_cf8_)) * (len(d_1709_v49_)))
                d_1710_v50_: D6
                d_1710_v50_ = D6_DC17(d_1691_cf6_, (self).f10, d_1625_v1_, d_1689_cf8_, (d_1639_v3_).cardinality)
                d_1711_v51_: _dafny.Array
                nw255_ = _dafny.Array(None, 2)
                nw255_[int(0)] = (d_1707_v47_).f10
                nw255_[int(1)] = not(d_1691_cf6_)
                d_1711_v51_ = nw255_
                d_1712_v52_: _dafny.Map
                d_1712_v52_ = _dafny.Map({d_1625_v1_: d_1711_v51_})
                d_1713_v53_: _dafny.Map
                d_1713_v53_ = _dafny.Map({d_1712_v52_: d_1691_cf6_})
                d_1714_v54_: _dafny.Array
                nw256_ = _dafny.Array(None, 21)
                nw256_[int(0)] = d_1691_cf6_
                nw256_[int(1)] = (self).f10
                nw256_[int(2)] = d_1691_cf6_
                nw256_[int(3)] = not(d_1690_cf7_)
                nw256_[int(4)] = (d_1710_v50_).cf29
                nw256_[int(5)] = (self).f10
                nw256_[int(6)] = True
                nw256_[int(7)] = (self).f10
                nw256_[int(8)] = False
                nw256_[int(9)] = d_1691_cf6_
                nw256_[int(10)] = ((d_1713_v53_)[d_1712_v52_] if (d_1712_v52_) in (d_1713_v53_) else True)
                nw256_[int(11)] = (self).f10
                nw256_[int(12)] = (d_1707_v47_).f10
                nw256_[int(13)] = False
                nw256_[int(14)] = False
                nw256_[int(15)] = d_1690_cf7_
                nw256_[int(16)] = (self).f10
                nw256_[int(17)] = True
                nw256_[int(18)] = d_1691_cf6_
                nw256_[int(19)] = (d_1707_v47_).f10
                nw256_[int(20)] = d_1691_cf6_
                d_1714_v54_ = nw256_
                d_1715_v55_: _dafny.Map
                d_1715_v55_ = _dafny.Map({d_1691_cf6_: d_1714_v54_})
                d_1716_v56_: _dafny.Map
                d_1716_v56_ = d_1715_v55_
                d_1717_v57_: _dafny.Array
                nw257_ = _dafny.Array(None, 21)
                nw257_[int(0)] = d_1716_v56_
                nw257_[int(1)] = d_1716_v56_
                nw257_[int(2)] = d_1716_v56_
                nw257_[int(3)] = d_1716_v56_
                nw257_[int(4)] = d_1715_v55_
                nw257_[int(5)] = d_1716_v56_
                nw257_[int(6)] = d_1716_v56_
                nw257_[int(7)] = d_1716_v56_
                nw257_[int(8)] = d_1716_v56_
                nw257_[int(9)] = d_1716_v56_
                nw257_[int(10)] = d_1716_v56_
                nw257_[int(11)] = d_1716_v56_
                nw257_[int(12)] = d_1716_v56_
                nw257_[int(13)] = d_1716_v56_
                nw257_[int(14)] = d_1716_v56_
                nw257_[int(15)] = d_1716_v56_
                nw257_[int(16)] = d_1716_v56_
                nw257_[int(17)] = d_1716_v56_
                nw257_[int(18)] = d_1716_v56_
                nw257_[int(19)] = d_1716_v56_
                nw257_[int(20)] = (d_1716_v56_ if d_1690_cf7_ else d_1716_v56_)
                d_1717_v57_ = nw257_
                index271_ = default__.safeIndex(463, (d_1717_v57_).length(0))
                (d_1717_v57_)[index271_] = d_1716_v56_
                index272_ = default__.safeIndex(463, (d_1717_v57_).length(0))
                (d_1717_v57_)[index272_] = d_1716_v56_
            elif source25_.is_DC2:
                d_1718___mcc_h16_ = source25_.cf4
                d_1719_cf4_ = d_1718___mcc_h16_
                d_1625_v1_ = d_1667_cf8_
                (globalState).f8 = d_1669_cf6_
                d_1720_v58_: _dafny.Array
                nw258_ = _dafny.Array(None, 8)
                d_1720_v58_ = nw258_
                index273_ = default__.safeIndex(699, (d_1720_v58_).length(0))
                (d_1720_v58_)[index273_] = d_1676_v26_
                index274_ = default__.safeIndex(699, (d_1720_v58_).length(0))
                (d_1720_v58_)[index274_] = d_1676_v26_
                d_1721_v59_: _dafny.Map
                d_1721_v59_ = _dafny.Map({d_1625_v1_: default__.fm13(d_1625_v1_, globalState)})
                d_1721_v59_ = (d_1721_v59_).set(d_1667_cf8_, (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_1722_i8_ in range(default__.abs(51))])) + (d_1642_v6_))
            elif True:
                d_1723___mcc_h17_ = source25_.cf9
                d_1724_cf9_ = d_1723___mcc_h17_
                d_1725_v60_: D3
                d_1725_v60_ = D3_DC7(d_1645_v9_)
                pat_let_tv37_ = d_1645_v9_
                d_1726_v61_: _dafny.Array
                nw259_ = _dafny.Array(None, 8)
                nw259_[int(0)] = d_1725_v60_
                def iife147_(_pat_let38_0):
                    def iife148_(d_1727_dt__update__tmp_h1_):
                        def iife149_(_pat_let39_0):
                            def iife150_(d_1728_dt__update_hcf11_h0_):
                                return D3_DC7(d_1728_dt__update_hcf11_h0_)
                            return iife150_(_pat_let39_0)
                        return iife149_(pat_let_tv37_)
                    return iife148_(_pat_let38_0)
                nw259_[int(1)] = iife147_(d_1725_v60_)
                nw259_[int(2)] = d_1725_v60_
                nw259_[int(3)] = d_1725_v60_
                nw259_[int(4)] = d_1725_v60_
                nw259_[int(5)] = d_1725_v60_
                nw259_[int(6)] = d_1725_v60_
                nw259_[int(7)] = d_1725_v60_
                d_1726_v61_ = nw259_
                d_1729_v62_: _dafny.Map
                d_1729_v62_ = _dafny.Map({d_1726_v61_: False})
                r1 = (d_1669_cf6_ if d_1668_cf7_ else not((d_1729_v62_) == (d_1729_v62_)))
                d_1730_v63_: _dafny.Array
                nw260_ = _dafny.Array(None, 1)
                nw260_[int(0)] = d_1669_cf6_
                d_1730_v63_ = nw260_
                index275_ = default__.safeIndex(304, (d_1730_v63_).length(0))
                (d_1730_v63_)[index275_] = not((self).fm0((D8_DC23(_dafny.SeqWithoutIsStrInference([d_1669_cf6_]))).cf41, False, d_1625_v1_, d_1668_cf7_, globalState))
                index276_ = default__.safeIndex(304, (d_1730_v63_).length(0))
                (d_1730_v63_)[index276_] = not(d_1668_cf7_)
                d_1731_v65_: _dafny.Map
                def iife151_():
                    coll71_ = _dafny.Map()
                    compr_71_: int
                    for compr_71_ in _dafny.IntegerRange(-892, 990):
                        d_1732_v64_: int = compr_71_
                        if ((-892) <= (d_1732_v64_)) and ((d_1732_v64_) < (990)):
                            coll71_[default__.safeDivisionInt(d_1732_v64_, 547)] = (D14_DC35(d_1625_v1_, d_1669_cf6_, p0)).cf62
                    return _dafny.Map(coll71_)
                d_1731_v65_ = _dafny.Map({d_1644_v8_: iife151_()
                })
                d_1625_v1_ = len(d_1731_v65_)
                d_1675_v25_ = ((_dafny.SeqWithoutIsStrInference([(d_1730_v63_)[default__.safeIndex(304, (d_1730_v63_).length(0))], (self).f10])) + (d_1675_v25_)) + (d_1675_v25_)
        elif source24_.is_DC2:
            d_1733___mcc_h10_ = source24_.cf4
            d_1734_cf4_ = d_1733___mcc_h10_
            d_1625_v1_ = d_1625_v1_
            rhs267_ = ((0) - (d_1625_v1_)) < ((0) - ((self).fm1((self).f9, globalState)))
            rhs268_ = True
            rhs269_ = d_1645_v9_
            rhs270_ = (d_1625_v1_) != (len(d_1644_v8_))
            lhs183_ = globalState
            lhs184_ = globalState
            lhs185_ = globalState
            lhs183_.f8 = rhs267_
            lhs184_.f8 = rhs268_
            d_1645_v9_ = rhs269_
            lhs185_.f8 = rhs270_
            d_1735_v66_: _dafny.Seq
            d_1735_v66_ = _dafny.SeqWithoutIsStrInference([(self).f10, False])
            if (d_1735_v66_)[default__.safeIndex(d_1625_v1_, len(d_1735_v66_))]:
                d_1736_v67_: _dafny.Map
                d_1736_v67_ = _dafny.Map({((d_1644_v8_)[len(_dafny.SeqWithoutIsStrInference([d_1645_v9_, d_1645_v9_]))] if (len(_dafny.SeqWithoutIsStrInference([d_1645_v9_, d_1645_v9_]))) in (d_1644_v8_) else (self).f10): d_1625_v1_})
                d_1737_v68_: _dafny.Map
                d_1737_v68_ = _dafny.Map({d_1736_v67_: d_1735_v66_})
                rhs271_ = not ((d_1737_v68_) != (d_1737_v68_)) or (False)
                rhs272_ = (d_1644_v8_) | ((d_1644_v8_) | (d_1644_v8_))
                lhs186_ = globalState
                lhs186_.f8 = rhs271_
                d_1644_v8_ = rhs272_
                d_1738_v69_: _dafny.Array
                def lambda78_(d_1739_p0_):
                    def lambda79_(d_1740_i9_):
                        return d_1739_p0_

                    return lambda79_

                init46_ = lambda78_(p0)
                nw261_ = _dafny.Array(None, 29)
                for i0_46_ in range(nw261_.length(0)):
                    nw261_[i0_46_] = init46_(i0_46_)
                d_1738_v69_ = nw261_
                index277_ = default__.safeIndex(88, (d_1738_v69_).length(0))
                (d_1738_v69_)[index277_] = d_1655_v14_
                d_1741_v70_: C3
                nw262_ = C3()
                nw262_.ctor__()
                d_1741_v70_ = nw262_
                d_1742_v71_: _dafny.Map
                d_1742_v71_ = _dafny.Map({d_1625_v1_: d_1741_v70_})
                d_1743_v72_: C5
                nw263_ = C5()
                nw263_.ctor__(_dafny.Map({(self).f10: (self).f10}), (self).f10)
                d_1743_v72_ = nw263_
                d_1744_v73_: _dafny.Map
                d_1744_v73_ = _dafny.Map({d_1743_v72_: d_1625_v1_})
                d_1745_v74_: _dafny.Map
                d_1745_v74_ = _dafny.Map({(D6_DC17(False, (self).f10, d_1625_v1_, -929, len(d_1742_v71_))).cf32: len(d_1744_v73_)})
                index278_ = default__.safeIndex(88, (d_1738_v69_).length(0))
                rhs273_ = (d_1639_v3_).ispropersubset(default__.fm10((self).f10, (0) - (d_1625_v1_), d_1642_v6_, d_1745_v74_, globalState))
                rhs274_ = p0
                lhs187_ = globalState
                lhs188_ = d_1738_v69_
                lhs189_ = default__.safeIndex(88, (d_1738_v69_).length(0))
                lhs187_.f8 = rhs273_
                lhs188_[lhs189_] = rhs274_
                d_1746_v75_: _dafny.Map
                d_1746_v75_ = _dafny.Map({(self).f10: d_1734_cf4_})
                d_1625_v1_ = len(((d_1746_v75_) | (d_1746_v75_)) | ((d_1746_v75_) | (d_1746_v75_)))
                d_1747_v76_: C3
                nw264_ = C3()
                nw264_.ctor__()
                d_1747_v76_ = nw264_
                d_1748_v77_: _dafny.MultiSet
                d_1748_v77_ = _dafny.MultiSet([(self).f10, False])
                d_1749_v78_: D6
                d_1749_v78_ = D6_DC17((self).f10, (self).f10, ((d_1748_v77_)[(self).f10] if ((self).f10) in (d_1748_v77_) else 667), d_1625_v1_, d_1625_v1_)
                d_1750_v79_: D3
                d_1750_v79_ = D3_DC8((self).f10, (self).f10)
                pat_let_tv38_ = d_1747_v76_
                pat_let_tv39_ = d_1750_v79_
                pat_let_tv40_ = globalState
                pat_let_tv41_ = globalState
                pat_let_tv42_ = d_1625_v1_
                d_1751_v80_: _dafny.Seq
                def iife152_(_pat_let40_0):
                    def iife153_(d_1752_dt__update__tmp_h2_):
                        def iife154_(_pat_let41_0):
                            def iife155_(d_1753_dt__update_hcf33_h0_):
                                def iife156_(_pat_let42_0):
                                    def iife157_(d_1754_dt__update_hcf32_h0_):
                                        return D6_DC17((d_1752_dt__update__tmp_h2_).cf29, (d_1752_dt__update__tmp_h2_).cf30, (d_1752_dt__update__tmp_h2_).cf31, d_1754_dt__update_hcf32_h0_, d_1753_dt__update_hcf33_h0_)
                                    return iife157_(_pat_let42_0)
                                return iife156_(pat_let_tv42_)
                            return iife155_(_pat_let41_0)
                        return iife154_((default__.fm46((self).f10, (pat_let_tv38_).fm27((self).f10, (self).f10, pat_let_tv39_, pat_let_tv40_), (self).f10, pat_let_tv41_)).cardinality)
                    return iife153_(_pat_let40_0)
                d_1751_v80_ = _dafny.SeqWithoutIsStrInference([d_1736_v67_, _dafny.Map({False: (iife152_(d_1749_v78_)).cf31}), d_1736_v67_])
                d_1751_v80_ = _dafny.SeqWithoutIsStrInference([(d_1736_v67_) | (d_1736_v67_)])
            elif True:
                d_1755_v81_: C0
                nw265_ = C0()
                nw265_.ctor__()
                d_1755_v81_ = nw265_
                d_1756_v82_: _dafny.Array
                def lambda80_(d_1757_cf4_):
                    def lambda81_(d_1758_i10_):
                        return d_1757_cf4_

                    return lambda81_

                init47_ = lambda80_(d_1734_cf4_)
                nw266_ = _dafny.Array(None, 16)
                for i0_47_ in range(nw266_.length(0)):
                    nw266_[i0_47_] = init47_(i0_47_)
                d_1756_v82_ = nw266_
                index279_ = default__.safeIndex(941, (d_1756_v82_).length(0))
                (d_1756_v82_)[index279_] = _dafny.SeqWithoutIsStrInference([d_1625_v1_])
                index280_ = default__.safeIndex(941, (d_1756_v82_).length(0))
                (d_1756_v82_)[index280_] = ((d_1734_cf4_).set(default__.safeIndex(591, len(d_1734_cf4_)), d_1625_v1_)) + (default__.fm37(globalState))
                d_1655_v14_ = d_1655_v14_
                d_1759_v83_: _dafny.Array
                nw267_ = _dafny.Array(_dafny.Array(None, 0), 13)
                d_1759_v83_ = nw267_
                d_1760_v84_: D17
                d_1760_v84_ = D17_DC43(d_1759_v83_)
                d_1761_v85_: _dafny.Array
                nw268_ = _dafny.Array(None, 16)
                nw268_[int(0)] = d_1759_v83_
                nw268_[int(1)] = d_1759_v83_
                nw268_[int(2)] = d_1759_v83_
                nw268_[int(3)] = d_1759_v83_
                nw268_[int(4)] = d_1759_v83_
                nw268_[int(5)] = (d_1760_v84_).cf83
                nw268_[int(6)] = d_1759_v83_
                nw268_[int(7)] = d_1759_v83_
                nw268_[int(8)] = d_1759_v83_
                nw268_[int(9)] = d_1759_v83_
                nw268_[int(10)] = d_1759_v83_
                nw268_[int(11)] = d_1759_v83_
                nw268_[int(12)] = d_1759_v83_
                nw268_[int(13)] = d_1759_v83_
                nw268_[int(14)] = (d_1760_v84_).cf83
                nw268_[int(15)] = d_1759_v83_
                d_1761_v85_ = nw268_
                index281_ = default__.safeIndex(874, (d_1761_v85_).length(0))
                (d_1761_v85_)[index281_] = d_1759_v83_
                index282_ = default__.safeIndex(874, (d_1761_v85_).length(0))
                (d_1761_v85_)[index282_] = d_1759_v83_
                d_1762_v86_: _dafny.Array
                def lambda82_(d_1763_v6_, d_1764_v1_):
                    def lambda83_(d_1765_i11_):
                        return _dafny.Set({len(d_1763_v6_), 684, d_1764_v1_})

                    return lambda83_

                init48_ = lambda82_(d_1642_v6_, d_1625_v1_)
                nw269_ = _dafny.Array(None, 17)
                for i0_48_ in range(nw269_.length(0)):
                    nw269_[i0_48_] = init48_(i0_48_)
                d_1762_v86_ = nw269_
                d_1766_v87_: _dafny.Set
                d_1766_v87_ = _dafny.Set({d_1625_v1_, d_1625_v1_})
                d_1767_v88_: D4
                d_1767_v88_ = D4_DC11(d_1625_v1_, d_1766_v87_, (self).f10)
                d_1768_v89_: _dafny.Set
                d_1768_v89_ = _dafny.Set({len(_dafny.Map({(d_1767_v88_).cf16: _dafny.CodePoint('p')})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bjnevp")))})
                index283_ = default__.safeIndex(860, (d_1762_v86_).length(0))
                (d_1762_v86_)[index283_] = (d_1766_v87_) | (d_1766_v87_)
                index284_ = default__.safeIndex(860, (d_1762_v86_).length(0))
                rhs275_ = ((d_1644_v8_)[d_1625_v1_] if (d_1625_v1_) in (d_1644_v8_) else (d_1642_v6_) < (d_1642_v6_))
                rhs276_ = default__.fm25(len(d_1642_v6_), d_1625_v1_, globalState)
                lhs190_ = globalState
                lhs191_ = d_1762_v86_
                lhs192_ = default__.safeIndex(860, (d_1762_v86_).length(0))
                lhs190_.f8 = rhs275_
                lhs191_[lhs192_] = rhs276_
            d_1769_v90_: _dafny.Set
            d_1769_v90_ = _dafny.Set({d_1625_v1_})
            d_1769_v90_ = d_1769_v90_
        elif True:
            d_1770___mcc_h11_ = source24_.cf9
            d_1771_cf9_ = d_1770___mcc_h11_
            index285_ = default__.safeIndex(863, (d_1645_v9_).length(0))
            (d_1645_v9_)[index285_] = d_1625_v1_
            d_1772_v91_: _dafny.Set
            d_1772_v91_ = _dafny.Set({d_1642_v6_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yp")), d_1642_v6_, _dafny.SeqWithoutIsStrInference([d_1655_v14_ for d_1773_i12_ in range(default__.abs(-425))]), d_1642_v6_})
            index286_ = default__.safeIndex(863, (d_1645_v9_).length(0))
            rhs277_ = (len(d_1772_v91_)) - (d_1625_v1_)
            rhs278_ = (d_1625_v1_) + (d_1625_v1_)
            lhs193_ = d_1645_v9_
            lhs194_ = default__.safeIndex(863, (d_1645_v9_).length(0))
            d_1625_v1_ = rhs277_
            lhs193_[lhs194_] = rhs278_
            if (self).f10:
                d_1774_v92_: D8
                d_1774_v92_ = D8_DC23(_dafny.SeqWithoutIsStrInference([(self).f10, (self).f10]))
                d_1775_v93_: _dafny.Array
                nw270_ = _dafny.Array(False, 16)
                d_1775_v93_ = nw270_
                index287_ = default__.safeIndex(442, (d_1775_v93_).length(0))
                (d_1775_v93_)[index287_] = (self).f10
                d_1776_v94_: _dafny.Seq
                d_1776_v94_ = _dafny.SeqWithoutIsStrInference([(self).f10, (d_1625_v1_) == (91)])
                index288_ = default__.safeIndex(442, (d_1775_v93_).length(0))
                index289_ = default__.safeIndex(863, (d_1645_v9_).length(0))
                rhs279_ = d_1774_v92_
                rhs280_ = (d_1625_v1_) < (default__.safeDivisionInt(d_1625_v1_, d_1625_v1_))
                rhs281_ = (d_1625_v1_) in (_dafny.SeqWithoutIsStrInference([d_1625_v1_ for d_1777_i13_ in range(default__.abs(983))]))
                rhs282_ = d_1625_v1_
                rhs283_ = (d_1776_v94_)[default__.safeIndex(len(d_1642_v6_), len(d_1776_v94_))]
                lhs195_ = d_1775_v93_
                lhs196_ = default__.safeIndex(442, (d_1775_v93_).length(0))
                lhs197_ = globalState
                lhs198_ = d_1645_v9_
                lhs199_ = default__.safeIndex(863, (d_1645_v9_).length(0))
                lhs200_ = globalState
                d_1774_v92_ = rhs279_
                lhs195_[lhs196_] = rhs280_
                lhs197_.f8 = rhs281_
                lhs198_[lhs199_] = rhs282_
                lhs200_.f8 = rhs283_
                d_1778_v95_: D14
                d_1778_v95_ = D14_DC35(d_1625_v1_, (self).f10, d_1655_v14_)
                d_1655_v14_ = (d_1778_v95_).cf64
                (globalState).f8 = (d_1642_v6_) != ((d_1642_v6_) + (d_1642_v6_))
                d_1779_v96_: C3
                nw271_ = C3()
                nw271_.ctor__()
                d_1779_v96_ = nw271_
                d_1780_v97_: _dafny.Map
                d_1780_v97_ = _dafny.Map({default__.fm53(d_1625_v1_, globalState): d_1779_v96_})
                index290_ = default__.safeIndex(442, (d_1775_v93_).length(0))
                index291_ = default__.safeIndex(863, (d_1645_v9_).length(0))
                index292_ = default__.safeIndex(863, (d_1645_v9_).length(0))
                rhs284_ = not ((self).f10) or ((self).f10)
                rhs285_ = (default__.safeModuloInt((d_1645_v9_)[default__.safeIndex(863, (d_1645_v9_).length(0))], d_1625_v1_)) < ((d_1645_v9_)[default__.safeIndex(863, (d_1645_v9_).length(0))])
                rhs286_ = (d_1780_v97_) != (d_1780_v97_)
                rhs287_ = ((d_1645_v9_)[default__.safeIndex(863, (d_1645_v9_).length(0))]) - (471)
                rhs288_ = (d_1645_v9_)[default__.safeIndex(863, (d_1645_v9_).length(0))]
                lhs201_ = d_1775_v93_
                lhs202_ = default__.safeIndex(442, (d_1775_v93_).length(0))
                lhs203_ = globalState
                lhs204_ = globalState
                lhs205_ = d_1645_v9_
                lhs206_ = default__.safeIndex(863, (d_1645_v9_).length(0))
                lhs207_ = d_1645_v9_
                lhs208_ = default__.safeIndex(863, (d_1645_v9_).length(0))
                lhs201_[lhs202_] = rhs284_
                lhs203_.f8 = rhs285_
                lhs204_.f8 = rhs286_
                lhs205_[lhs206_] = rhs287_
                lhs207_[lhs208_] = rhs288_
                d_1781_v98_: _dafny.Map
                d_1781_v98_ = _dafny.Map({d_1625_v1_: (self).f9})
                d_1782_v99_: _dafny.Map
                d_1782_v99_ = _dafny.Map({(self).f10: d_1640_v4_})
                d_1783_v100_: C2
                nw272_ = C2()
                nw272_.ctor__((d_1775_v93_)[default__.safeIndex(442, (d_1775_v93_).length(0))], d_1775_v93_, (_dafny.Map({(self).f10: (d_1775_v93_)[default__.safeIndex(442, (d_1775_v93_).length(0))]})) | (((d_1781_v98_)[len(((d_1782_v99_)[not((d_1775_v93_)[default__.safeIndex(442, (d_1775_v93_).length(0))])] if (not((d_1775_v93_)[default__.safeIndex(442, (d_1775_v93_).length(0))])) in (d_1782_v99_) else d_1640_v4_))] if (len(((d_1782_v99_)[not((d_1775_v93_)[default__.safeIndex(442, (d_1775_v93_).length(0))])] if (not((d_1775_v93_)[default__.safeIndex(442, (d_1775_v93_).length(0))])) in (d_1782_v99_) else d_1640_v4_))) in (d_1781_v98_) else (self).f9)), (self).f10)
                d_1783_v100_ = nw272_
            elif True:
                d_1784_v101_: C3
                nw273_ = C3()
                nw273_.ctor__()
                d_1784_v101_ = nw273_
                index293_ = default__.safeIndex(863, (d_1645_v9_).length(0))
                (d_1645_v9_)[index293_] = ((d_1645_v9_)[default__.safeIndex(863, (d_1645_v9_).length(0))]) - ((d_1645_v9_)[default__.safeIndex(863, (d_1645_v9_).length(0))])
                d_1639_v3_ = (d_1639_v3_) | ((d_1639_v3_) | (d_1639_v3_))
                d_1785_v102_: C4
                nw274_ = C4()
                nw274_.ctor__((self).f9, (d_1625_v1_) > ((d_1645_v9_)[default__.safeIndex(863, (d_1645_v9_).length(0))]))
                d_1785_v102_ = nw274_
                d_1785_v102_ = d_1785_v102_
                index294_ = default__.safeIndex(863, (d_1645_v9_).length(0))
                (d_1645_v9_)[index294_] = (d_1645_v9_)[default__.safeIndex(863, (d_1645_v9_).length(0))]
            d_1786_v103_: _dafny.Array
            nw275_ = _dafny.Array(False, 1)
            d_1786_v103_ = nw275_
            index295_ = default__.safeIndex(734, (d_1786_v103_).length(0))
            (d_1786_v103_)[index295_] = (self).f10
            index296_ = default__.safeIndex(734, (d_1786_v103_).length(0))
            (d_1786_v103_)[index296_] = (self).f10
            d_1787_v104_: _dafny.Array
            def lambda84_(d_1788_v6_):
                def lambda85_(d_1789_i14_):
                    return (d_1789_i14_) + ((0) - (len(d_1788_v6_)))

                return lambda85_

            init49_ = lambda84_(d_1642_v6_)
            nw276_ = _dafny.Array(None, 6)
            for i0_49_ in range(nw276_.length(0)):
                nw276_[i0_49_] = init49_(i0_49_)
            d_1787_v104_ = nw276_
            d_1790_v105_: _dafny.Set
            d_1790_v105_ = _dafny.Set({d_1787_v104_, d_1787_v104_})
            d_1791_v106_: _dafny.Map
            d_1791_v106_ = _dafny.Map({d_1790_v105_: (d_1786_v103_)[default__.safeIndex(734, (d_1786_v103_).length(0))]})
            d_1791_v106_ = (d_1791_v106_).set((d_1790_v105_ if (d_1786_v103_)[default__.safeIndex(734, (d_1786_v103_).length(0))] else d_1790_v105_), not((d_1786_v103_)[default__.safeIndex(734, (d_1786_v103_).length(0))]))
        d_1792_v107_: _dafny.Map
        d_1792_v107_ = _dafny.Map({d_1645_v9_: default__.fm20(d_1625_v1_, (self).f10, (self).f10, d_1625_v1_, globalState)})
        r0 = d_1792_v107_
        r1 = not((False if (self).f10 else (self).f10))
        return r0, r1


class C7(T1, T0):
    def  __init__(self):
        self._f9: _dafny.Map = _dafny.Map({})
        self._f10: bool = False
        self._f11: _dafny.Set = _dafny.Set({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C7"
    @property
    def f9(self):
        return self._f9
    @property
    def f10(self):
        return self._f10
    def ctor__(self, f11, f9, f10):
        (self)._f11 = f11
        (self)._f9 = f9
        (self)._f10 = f10

    def fm0(self, p0, p1, p2, p3, globalState):
        return (self).f10

    def fm1(self, p0, globalState):
        return (0) - (len(_dafny.SeqWithoutIsStrInference([(self).f10])))

    def fm2(self, globalState):
        return (default__.safeDivisionInt((0) - (len((self).f11)), (_dafny.MultiSet([len(_dafny.Map({False: 73})), (0) - (-305), (0) - (-142)])).cardinality)) <= (len(_dafny.SeqWithoutIsStrInference([(D0_DC0(not(not((self).f10)))).cf0])))

    def fm3(self, p0, p1, p2, p3, globalState):
        return len((D1_DC2(_dafny.SeqWithoutIsStrInference([len(((self).f9).set(True, (self).f10))]))).cf4)

    def m0(self, globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r1: bool = False
        r2: _dafny.Seq = _dafny.Seq({})
        r3: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        d_1793_v0_: int
        d_1793_v0_ = 233
        d_1794_v1_: _dafny.Set
        d_1794_v1_ = _dafny.Set({d_1793_v0_, d_1793_v0_})
        d_1795_v2_: _dafny.Seq
        d_1795_v2_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({not((self).f10): not((self).f10)})])
        hi5_ = (len(d_1795_v2_) if (self).f10 else (0) - (d_1793_v0_))
        for d_1796_i0_ in range((0) - (default__.safeDivisionInt(len(d_1794_v1_), -93)), hi5_):
            d_1797_v3_: _dafny.Seq
            d_1797_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bt"))
            d_1798_v4_: D0
            d_1798_v4_ = D0_DC1(d_1793_v0_, (self).f10, (len(d_1797_v3_)) + (d_1793_v0_))
            d_1799_v6_: _dafny.Map
            d_1799_v6_ = _dafny.Map({(self).f10: d_1796_i0_})
            d_1800_v7_: _dafny.Array
            nw277_ = _dafny.Array(None, 6)
            def iife158_():
                coll72_ = _dafny.Map()
                compr_72_: int
                for compr_72_ in _dafny.IntegerRange(23, 718):
                    d_1801_v5_: int = compr_72_
                    if ((23) <= (d_1801_v5_)) and ((d_1801_v5_) < (718)):
                        coll72_[(d_1801_v5_) + (d_1793_v0_)] = d_1796_i0_
                return _dafny.Map(coll72_)
            nw277_[int(0)] = (0) - ((0) - (len(iife158_()
            )))
            nw277_[int(1)] = ((0) - (len(d_1797_v3_))) - (((d_1799_v6_)[(self).f10] if ((self).f10) in (d_1799_v6_) else d_1793_v0_))
            nw277_[int(2)] = 169
            nw277_[int(3)] = (0) - (d_1793_v0_)
            nw277_[int(4)] = len((d_1797_v3_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mootujt"))))
            nw277_[int(5)] = d_1793_v0_
            d_1800_v7_ = nw277_
            d_1802_v8_: _dafny.Map
            d_1802_v8_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_1797_v3_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ywcvn"))]))).cardinality for d_1803_i1_ in range(default__.abs(729))]): (self).f10})
            index297_ = default__.safeIndex(414, (d_1800_v7_).length(0))
            (d_1800_v7_)[index297_] = len((d_1802_v8_) | (d_1802_v8_))
            d_1804_v9_: _dafny.Array
            nw278_ = _dafny.Array(_dafny.CodePoint('D'), 10)
            d_1804_v9_ = nw278_
            d_1805_v10_: _dafny.Seq
            d_1805_v10_ = _dafny.SeqWithoutIsStrInference([(self).f10, (self).f10])
            d_1806_v11_: _dafny.Map
            d_1806_v11_ = _dafny.Map({d_1804_v9_: (d_1805_v10_).set(default__.safeIndex((0) - (d_1793_v0_), len(d_1805_v10_)), (self).f10)})
            d_1807_v12_: str
            d_1807_v12_ = _dafny.CodePoint('c')
            d_1808_v13_: _dafny.Seq
            d_1808_v13_ = _dafny.SeqWithoutIsStrInference([(d_1806_v11_) | ((d_1806_v11_).set(d_1804_v9_, default__.fm4(d_1807_v12_, globalState)))])
            d_1809_v14_: _dafny.Seq
            d_1809_v14_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1796_i0_])])
            index298_ = default__.safeIndex(414, (d_1800_v7_).length(0))
            rhs289_ = True
            rhs290_ = d_1798_v4_
            rhs291_ = len((d_1808_v13_)[default__.safeIndex((0) - (len(d_1809_v14_)), len(d_1808_v13_))])
            lhs209_ = globalState
            lhs210_ = d_1800_v7_
            lhs211_ = default__.safeIndex(414, (d_1800_v7_).length(0))
            lhs209_.f8 = rhs289_
            d_1798_v4_ = rhs290_
            lhs210_[lhs211_] = rhs291_
            d_1810_v15_: _dafny.Seq
            d_1810_v15_ = _dafny.SeqWithoutIsStrInference([(d_1800_v7_)[default__.safeIndex(414, (d_1800_v7_).length(0))]])
            d_1811_v16_: D1
            d_1811_v16_ = D1_DC2(d_1810_v15_)
            source26_ = D1_DC5(d_1811_v16_)
            if source26_.is_DC3:
                d_1812___mcc_h0_ = source26_.cf5
                d_1813_cf5_ = d_1812___mcc_h0_
                (globalState).f8 = (d_1794_v1_).issubset((d_1794_v1_) | (d_1794_v1_))
                (globalState).f8 = (default__.safeDivisionInt((0) - (d_1793_v0_), (0) - ((d_1810_v15_)[default__.safeIndex(d_1793_v0_, len(d_1810_v15_))]))) <= (d_1793_v0_)
                d_1814_v17_: C0
                nw279_ = C0()
                nw279_.ctor__()
                d_1814_v17_ = nw279_
                d_1815_v18_: C0
                nw280_ = C0()
                nw280_.ctor__()
                d_1815_v18_ = nw280_
            elif source26_.is_DC4:
                d_1816___mcc_h1_ = source26_.cf6
                d_1817___mcc_h2_ = source26_.cf7
                d_1818___mcc_h3_ = source26_.cf8
                d_1819_cf8_ = d_1818___mcc_h3_
                d_1820_cf7_ = d_1817___mcc_h2_
                d_1821_cf6_ = d_1816___mcc_h1_
                index299_ = default__.safeIndex(414, (d_1800_v7_).length(0))
                (d_1800_v7_)[index299_] = d_1793_v0_
                d_1822_v19_: _dafny.Array
                def lambda86_(d_1823_v3_):
                    def lambda87_(d_1824_i2_):
                        return d_1823_v3_

                    return lambda87_

                init50_ = lambda86_(d_1797_v3_)
                nw281_ = _dafny.Array(None, 15)
                for i0_50_ in range(nw281_.length(0)):
                    nw281_[i0_50_] = init50_(i0_50_)
                d_1822_v19_ = nw281_
                d_1825_v20_: D1
                d_1825_v20_ = D1_DC3(d_1797_v3_)
                index300_ = default__.safeIndex(769, (d_1822_v19_).length(0))
                (d_1822_v19_)[index300_] = (d_1825_v20_).cf5
                index301_ = default__.safeIndex(769, (d_1822_v19_).length(0))
                (d_1822_v19_)[index301_] = ((d_1797_v3_ if not((self).f10) else d_1797_v3_)) + (d_1797_v3_)
                (globalState).f8 = d_1820_cf7_
                d_1826_v21_: _dafny.Map
                d_1826_v21_ = _dafny.Map({d_1793_v0_: d_1793_v0_})
                d_1827_v22_: D0
                d_1827_v22_ = D0_DC0(not(d_1821_cf6_))
                d_1828_v23_: _dafny.Map
                d_1828_v23_ = _dafny.Map({(0) - ((d_1796_i0_) - (len(d_1826_v21_))): d_1827_v22_})
                def iife159_(_pat_let43_0):
                    def iife160_(d_1829_dt__update__tmp_h0_):
                        def iife161_(_pat_let44_0):
                            def iife162_(d_1830_dt__update_hcf0_h0_):
                                return D0_DC0(d_1830_dt__update_hcf0_h0_)
                            return iife162_(_pat_let44_0)
                        return iife161_((self).f10)
                    return iife160_(_pat_let43_0)
                d_1828_v23_ = (d_1828_v23_).set(308, iife159_(d_1827_v22_))
            elif source26_.is_DC2:
                d_1831___mcc_h4_ = source26_.cf4
                d_1832_cf4_ = d_1831___mcc_h4_
                index302_ = default__.safeIndex(414, (d_1800_v7_).length(0))
                (d_1800_v7_)[index302_] = d_1793_v0_
                d_1833_v24_: _dafny.MultiSet
                d_1833_v24_ = _dafny.MultiSet([58])
                index303_ = default__.safeIndex(414, (d_1800_v7_).length(0))
                rhs292_ = (d_1805_v10_)[default__.safeIndex(d_1793_v0_, len(d_1805_v10_))]
                rhs293_ = d_1796_i0_
                rhs294_ = not ((self).f10) or ((len(d_1805_v10_)) < ((d_1833_v24_).cardinality))
                rhs295_ = (d_1793_v0_) * (d_1793_v0_)
                lhs212_ = d_1800_v7_
                lhs213_ = default__.safeIndex(414, (d_1800_v7_).length(0))
                lhs214_ = globalState
                r1 = rhs292_
                lhs212_[lhs213_] = rhs293_
                lhs214_.f8 = rhs294_
                d_1793_v0_ = rhs295_
                d_1834_v25_: D1
                d_1834_v25_ = D1_DC4((self).fm2(globalState), (d_1794_v1_) == (d_1794_v1_), d_1796_i0_)
                d_1834_v25_ = (d_1834_v25_ if (self).f10 else d_1834_v25_)
                index304_ = default__.safeIndex(812, (d_1804_v9_).length(0))
                (d_1804_v9_)[index304_] = d_1807_v12_
                d_1835_v26_: _dafny.Map
                d_1835_v26_ = _dafny.Map({(self).f11: (self).f10})
                index305_ = default__.safeIndex(414, (d_1800_v7_).length(0))
                index306_ = default__.safeIndex(812, (d_1804_v9_).length(0))
                index307_ = default__.safeIndex(414, (d_1800_v7_).length(0))
                rhs296_ = default__.fm6(d_1793_v0_, d_1793_v0_, d_1835_v26_, globalState)
                rhs297_ = default__.safeDivisionInt((d_1800_v7_)[default__.safeIndex(414, (d_1800_v7_).length(0))], ((0) - ((d_1800_v7_)[default__.safeIndex(414, (d_1800_v7_).length(0))])) * (678))
                rhs298_ = d_1796_i0_
                rhs299_ = default__.fm7(d_1793_v0_, globalState)
                rhs300_ = (0) - ((d_1800_v7_)[default__.safeIndex(414, (d_1800_v7_).length(0))])
                lhs215_ = d_1800_v7_
                lhs216_ = default__.safeIndex(414, (d_1800_v7_).length(0))
                lhs217_ = d_1804_v9_
                lhs218_ = default__.safeIndex(812, (d_1804_v9_).length(0))
                lhs219_ = d_1800_v7_
                lhs220_ = default__.safeIndex(414, (d_1800_v7_).length(0))
                r0 = rhs296_
                d_1793_v0_ = rhs297_
                lhs215_[lhs216_] = rhs298_
                lhs217_[lhs218_] = rhs299_
                lhs219_[lhs220_] = rhs300_
            elif True:
                d_1836___mcc_h5_ = source26_.cf9
                d_1837_cf9_ = d_1836___mcc_h5_
                d_1838_v27_: D1
                d_1838_v27_ = D1_DC5(d_1811_v16_)
                d_1838_v27_ = d_1838_v27_
                d_1839_v28_: _dafny.MultiSet
                d_1839_v28_ = _dafny.MultiSet([(self).f10])
                d_1840_v29_: _dafny.Array
                def lambda88_(d_1841_i3_):
                    return (self).f10

                init51_ = lambda88_
                nw282_ = _dafny.Array(None, 5)
                for i0_51_ in range(nw282_.length(0)):
                    nw282_[i0_51_] = init51_(i0_51_)
                d_1840_v29_ = nw282_
                d_1842_v30_: T1
                nw283_ = C2()
                nw283_.ctor__((self).f10, d_1840_v29_, _dafny.Map({(self).f10: (self).f10}), False)
                d_1842_v30_ = nw283_
                d_1843_v31_: _dafny.MultiSet
                d_1843_v31_ = _dafny.MultiSet([d_1842_v30_])
                d_1844_v32_: _dafny.Map
                d_1844_v32_ = _dafny.Map({(d_1843_v31_).set(d_1842_v30_, default__.abs(len(d_1805_v10_))): d_1810_v15_})
                rhs301_ = (len((d_1844_v32_) | ((_dafny.Map({d_1843_v31_: d_1810_v15_})).set(_dafny.MultiSet([d_1842_v30_, d_1842_v30_]), d_1810_v15_)))) + ((0) - ((d_1800_v7_)[default__.safeIndex(414, (d_1800_v7_).length(0))]))
                rhs302_ = (d_1839_v28_).intersection(d_1839_v28_)
                rhs303_ = d_1800_v7_
                rhs304_ = not((d_1842_v30_).f10)
                lhs221_ = globalState
                d_1793_v0_ = rhs301_
                d_1839_v28_ = rhs302_
                d_1800_v7_ = rhs303_
                lhs221_.f8 = rhs304_
                d_1845_v33_: _dafny.Array
                nw284_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 22)
                d_1845_v33_ = nw284_
                d_1846_v34_: _dafny.Map
                d_1846_v34_ = _dafny.Map({(d_1842_v30_).f10: _dafny.Map({(self).f11: not((d_1842_v30_).f10)})})
                d_1847_v35_: _dafny.Map
                d_1847_v35_ = _dafny.Map({(self).f11: (self).f10})
                index308_ = default__.safeIndex(136, (d_1845_v33_).length(0))
                (d_1845_v33_)[index308_] = (default__.fm6(d_1793_v0_, d_1796_i0_, ((d_1846_v34_)[True] if (True) in (d_1846_v34_) else d_1847_v35_), globalState)) + (d_1797_v3_)
                d_1848_v36_: _dafny.Array
                nw285_ = _dafny.Array(_dafny.Map({}), 27)
                d_1848_v36_ = nw285_
                d_1849_v37_: _dafny.Map
                d_1849_v37_ = _dafny.Map({d_1807_v12_: d_1796_i0_})
                index309_ = default__.safeIndex(710, (d_1848_v36_).length(0))
                (d_1848_v36_)[index309_] = d_1849_v37_
                d_1850_v38_: D3
                d_1850_v38_ = D3_DC7(d_1800_v7_)
                d_1851_v39_: D3
                d_1851_v39_ = D3_DC9(d_1850_v38_)
                d_1852_v41_: _dafny.Map
                d_1852_v41_ = _dafny.Map({d_1807_v12_: d_1807_v12_})
                d_1853_v42_: D18
                d_1853_v42_ = D18_DC46(d_1852_v41_)
                index310_ = default__.safeIndex(136, (d_1845_v33_).length(0))
                index311_ = default__.safeIndex(710, (d_1848_v36_).length(0))
                def iife163_():
                    coll73_ = _dafny.Map()
                    compr_73_: str
                    for compr_73_ in ((d_1853_v42_).cf89).keys.Elements:
                        d_1854_v40_: str = compr_73_
                        if (d_1854_v40_) in ((d_1853_v42_).cf89):
                            coll73_[d_1854_v40_] = (d_1796_i0_) + (51)
                    return _dafny.Map(coll73_)
                rhs305_ = d_1797_v3_
                rhs306_ = iife163_()

                rhs307_ = d_1851_v39_
                rhs308_ = (self).f10
                lhs222_ = d_1845_v33_
                lhs223_ = default__.safeIndex(136, (d_1845_v33_).length(0))
                lhs224_ = d_1848_v36_
                lhs225_ = default__.safeIndex(710, (d_1848_v36_).length(0))
                lhs222_[lhs223_] = rhs305_
                lhs224_[lhs225_] = rhs306_
                d_1851_v39_ = rhs307_
                r1 = rhs308_
                d_1855_v43_: str
                d_1855_v43_ = d_1807_v12_
                d_1856_v44_: _dafny.Map
                d_1856_v44_ = _dafny.Map({(d_1855_v43_ if (self).f10 else d_1855_v43_): (d_1845_v33_)[default__.safeIndex(136, (d_1845_v33_).length(0))]})
                d_1856_v44_ = (_dafny.Map({d_1855_v43_: (d_1845_v33_)[default__.safeIndex(136, (d_1845_v33_).length(0))]})) | (d_1856_v44_)
            d_1857_v45_: _dafny.Array
            def lambda89_(d_1858_v10_):
                def lambda90_(d_1859_i4_):
                    return d_1858_v10_

                return lambda90_

            init52_ = lambda89_(d_1805_v10_)
            nw286_ = _dafny.Array(None, 7)
            for i0_52_ in range(nw286_.length(0)):
                nw286_[i0_52_] = init52_(i0_52_)
            d_1857_v45_ = nw286_
            index312_ = default__.safeIndex(503, (d_1857_v45_).length(0))
            (d_1857_v45_)[index312_] = d_1805_v10_
            index313_ = default__.safeIndex(503, (d_1857_v45_).length(0))
            rhs309_ = default__.safeModuloInt(d_1793_v0_, default__.fm33((self).f10, (self).f10, d_1796_i0_, (self).f10, globalState))
            rhs310_ = _dafny.CodePoint('w')
            rhs311_ = (d_1805_v10_) + (_dafny.SeqWithoutIsStrInference([(self).fm2(globalState), (self).f10, (self).f10]))
            lhs226_ = d_1857_v45_
            lhs227_ = default__.safeIndex(503, (d_1857_v45_).length(0))
            d_1793_v0_ = rhs309_
            d_1807_v12_ = rhs310_
            lhs226_[lhs227_] = rhs311_
            d_1860_v46_: C0
            nw287_ = C0()
            nw287_.ctor__()
            d_1860_v46_ = nw287_
        d_1861_v47_: T0
        nw288_ = C3()
        nw288_.ctor__()
        d_1861_v47_ = nw288_
        d_1862_v48_: _dafny.MultiSet
        d_1862_v48_ = _dafny.MultiSet([d_1861_v47_, d_1861_v47_, d_1861_v47_])
        d_1863_v49_: _dafny.Map
        d_1863_v49_ = _dafny.Map({d_1793_v0_: d_1862_v48_})
        hi6_ = len(d_1863_v49_)
        for d_1864_i5_ in range(default__.safeDivisionInt(d_1793_v0_, d_1793_v0_), hi6_):
            r1 = (self).f10
            if ((self).f10) == ((self).f10):
                d_1865_v50_: _dafny.Array
                nw289_ = _dafny.Array(int(0), 1)
                d_1865_v50_ = nw289_
                index314_ = default__.safeIndex(216, (d_1865_v50_).length(0))
                (d_1865_v50_)[index314_] = d_1864_i5_
                index315_ = default__.safeIndex(216, (d_1865_v50_).length(0))
                (d_1865_v50_)[index315_] = d_1864_i5_
                (globalState).f8 = (self).fm2(globalState)
                index316_ = default__.safeIndex(216, (d_1865_v50_).length(0))
                (d_1865_v50_)[index316_] = (0) - ((d_1864_i5_ if not(not ((self).f10) or ((self).f10)) else (d_1865_v50_)[default__.safeIndex(216, (d_1865_v50_).length(0))]))
                d_1866_v51_: _dafny.Seq
                d_1866_v51_ = _dafny.SeqWithoutIsStrInference([(d_1865_v50_)[default__.safeIndex(216, (d_1865_v50_).length(0))]])
                d_1867_v52_: _dafny.Seq
                d_1867_v52_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p"))
                d_1866_v51_ = _dafny.SeqWithoutIsStrInference([(d_1793_v0_) * (len(d_1867_v52_))])
                (globalState).f8 = (self).f10
            elif True:
                d_1868_v53_: str
                d_1868_v53_ = _dafny.CodePoint('x')
                d_1869_v54_: D6
                d_1869_v54_ = D6_DC19((self).f10, d_1868_v53_)
                pat_let_tv43_ = d_1868_v53_
                def iife164_(_pat_let45_0):
                    def iife165_(d_1870_dt__update__tmp_h1_):
                        def iife166_(_pat_let46_0):
                            def iife167_(d_1871_dt__update_hcf38_h0_):
                                return D6_DC19((d_1870_dt__update__tmp_h1_).cf37, d_1871_dt__update_hcf38_h0_)
                            return iife167_(_pat_let46_0)
                        return iife166_(pat_let_tv43_)
                    return iife165_(_pat_let45_0)
                (globalState).f8 = (iife164_(d_1869_v54_)).cf37
                d_1872_v55_: C3
                nw290_ = C3()
                nw290_.ctor__()
                d_1872_v55_ = nw290_
                d_1873_v56_: C0
                nw291_ = C0()
                nw291_.ctor__()
                d_1873_v56_ = nw291_
                d_1793_v0_ = default__.safeDivisionInt(d_1793_v0_, (0) - ((0) - (d_1793_v0_)))
                d_1874_v57_: _dafny.Map
                d_1874_v57_ = _dafny.Map({d_1793_v0_: d_1864_i5_})
                d_1875_v58_: _dafny.Map
                d_1875_v58_ = _dafny.Map({(self).f10: d_1874_v57_})
                d_1876_v59_: _dafny.Seq
                d_1876_v59_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "whn"))
                d_1877_v60_: _dafny.Map
                d_1877_v60_ = _dafny.Map({d_1793_v0_: len(d_1876_v59_)})
                d_1875_v58_ = (d_1875_v58_).set((self).f10, d_1877_v60_)
            if (self).f10:
                r1 = (self).f10
                d_1878_v61_: _dafny.Array
                def lambda91_(d_1879_i6_):
                    return (d_1879_i6_) * (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_1880_i7_ in range(default__.abs(-602))])))

                init53_ = lambda91_
                nw292_ = _dafny.Array(None, 9)
                for i0_53_ in range(nw292_.length(0)):
                    nw292_[i0_53_] = init53_(i0_53_)
                d_1878_v61_ = nw292_
                d_1881_v62_: _dafny.Seq
                d_1881_v62_ = _dafny.SeqWithoutIsStrInference([d_1878_v61_])
                d_1882_v63_: D12
                d_1882_v63_ = D12_DC32(d_1878_v61_, (d_1881_v62_).set(default__.safeIndex(851, len(d_1881_v62_)), d_1878_v61_))
                d_1882_v63_ = d_1882_v63_
                d_1883_v64_: _dafny.MultiSet
                d_1883_v64_ = _dafny.MultiSet([(self).f10])
                d_1884_v65_: _dafny.MultiSet
                d_1884_v65_ = _dafny.MultiSet([len((self).f9)])
                d_1885_v66_: str
                d_1885_v66_ = _dafny.CodePoint('a')
                d_1886_v67_: _dafny.MultiSet
                d_1886_v67_ = _dafny.MultiSet([d_1885_v66_])
                d_1887_v68_: _dafny.Seq
                d_1887_v68_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dokjnmwo"))
                d_1888_v69_: _dafny.Map
                d_1888_v69_ = _dafny.Map({d_1864_i5_: d_1887_v68_})
                d_1889_v70_: _dafny.Seq
                d_1889_v70_ = _dafny.SeqWithoutIsStrInference([d_1887_v68_])
                d_1890_v71_: _dafny.Array
                nw293_ = _dafny.Array(None, 20)
                nw293_[int(0)] = (self).f10
                nw293_[int(1)] = (self).f10
                nw293_[int(2)] = (d_1864_i5_) <= (d_1793_v0_)
                nw293_[int(3)] = (self).f10
                nw293_[int(4)] = not ((self).f10) or ((self).f10)
                nw293_[int(5)] = (self).f10
                nw293_[int(6)] = (d_1864_i5_) >= (((d_1883_v64_)[True] if (True) in (d_1883_v64_) else d_1793_v0_))
                nw293_[int(7)] = (self).f10
                nw293_[int(8)] = (self).f10
                nw293_[int(9)] = (self).f10
                nw293_[int(10)] = not((d_1884_v65_).ispropersubset(d_1884_v65_))
                nw293_[int(11)] = (d_1886_v67_).ispropersubset(d_1886_v67_)
                nw293_[int(12)] = (d_1885_v66_) not in (((d_1888_v69_)[d_1864_i5_] if (d_1864_i5_) in (d_1888_v69_) else (d_1889_v70_)[default__.safeIndex((0) - (d_1793_v0_), len(d_1889_v70_))]))
                nw293_[int(13)] = (self).f10
                nw293_[int(14)] = (self).f10
                nw293_[int(15)] = not(True)
                nw293_[int(16)] = not((self).f10)
                nw293_[int(17)] = (self).f10
                nw293_[int(18)] = True
                nw293_[int(19)] = ((self).f10 if False else (self).f10)
                d_1890_v71_ = nw293_
                index317_ = default__.safeIndex(298, (d_1890_v71_).length(0))
                (d_1890_v71_)[index317_] = (self).f10
                d_1891_v72_: D12
                d_1891_v72_ = D12_DC31((self).f10, d_1887_v68_, _dafny.SeqWithoutIsStrInference([d_1885_v66_ for d_1892_i8_ in range(default__.abs(287))]))
                pat_let_tv44_ = d_1887_v68_
                pat_let_tv45_ = d_1793_v0_
                pat_let_tv46_ = d_1885_v66_
                d_1893_v73_: T1
                nw294_ = C2()
                def iife168_(_pat_let47_0):
                    def iife169_(d_1894_dt__update__tmp_h2_):
                        def iife170_(_pat_let48_0):
                            def iife171_(d_1895_dt__update_hcf57_h0_):
                                def iife172_(_pat_let49_0):
                                    def iife173_(d_1896_dt__update_hcf56_h0_):
                                        return D12_DC31((d_1894_dt__update__tmp_h2_).cf55, d_1896_dt__update_hcf56_h0_, d_1895_dt__update_hcf57_h0_)
                                    return iife173_(_pat_let49_0)
                                return iife172_((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cguqrqn"))).set(default__.safeIndex(pat_let_tv45_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cguqrqn")))), pat_let_tv46_))
                            return iife171_(_pat_let48_0)
                        return iife170_(pat_let_tv44_)
                    return iife169_(_pat_let47_0)
                nw294_.ctor__(True, d_1890_v71_, (self).f9, (iife168_(d_1891_v72_)).cf55)
                d_1893_v73_ = nw294_
                d_1897_v74_: D16
                d_1897_v74_ = D16_DC41((d_1893_v73_).f10, 563, (self).f10)
                d_1898_v75_: _dafny.Seq
                d_1898_v75_ = _dafny.SeqWithoutIsStrInference([576, (d_1897_v74_).cf78])
                d_1899_v76_: _dafny.Seq
                d_1899_v76_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_1898_v75_, _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i"))) for d_1900_i9_ in range(default__.abs(194))])]))])
                d_1901_v77_: _dafny.MultiSet
                d_1901_v77_ = _dafny.MultiSet([d_1898_v75_, d_1898_v75_])
                index318_ = default__.safeIndex(298, (d_1890_v71_).length(0))
                rhs312_ = (d_1901_v77_).ispropersubset((d_1899_v76_)[default__.safeIndex(((d_1884_v65_)[d_1864_i5_] if (d_1864_i5_) in (d_1884_v65_) else d_1793_v0_), len(d_1899_v76_))])
                rhs313_ = d_1893_v73_
                lhs228_ = d_1890_v71_
                lhs229_ = default__.safeIndex(298, (d_1890_v71_).length(0))
                lhs228_[lhs229_] = rhs312_
                d_1893_v73_ = rhs313_
                d_1902_v78_: C2
                nw295_ = C2()
                nw295_.ctor__((d_1890_v71_)[default__.safeIndex(298, (d_1890_v71_).length(0))], d_1890_v71_, (d_1893_v73_).f9, (d_1890_v71_)[default__.safeIndex(298, (d_1890_v71_).length(0))])
                d_1902_v78_ = nw295_
                d_1903_v79_: _dafny.Array
                nw296_ = _dafny.Array(None, 18)
                nw296_[int(0)] = d_1902_v78_
                nw296_[int(1)] = d_1902_v78_
                nw296_[int(2)] = d_1902_v78_
                nw296_[int(3)] = d_1902_v78_
                nw296_[int(4)] = d_1902_v78_
                nw296_[int(5)] = d_1902_v78_
                nw296_[int(6)] = d_1902_v78_
                nw296_[int(7)] = d_1902_v78_
                nw296_[int(8)] = (d_1902_v78_ if d_1902_v78_.f14 else d_1902_v78_)
                nw296_[int(9)] = d_1902_v78_
                nw296_[int(10)] = d_1902_v78_
                nw296_[int(11)] = d_1902_v78_
                nw296_[int(12)] = d_1902_v78_
                nw296_[int(13)] = d_1902_v78_
                nw296_[int(14)] = d_1902_v78_
                nw296_[int(15)] = d_1902_v78_
                nw296_[int(16)] = d_1902_v78_
                nw296_[int(17)] = d_1902_v78_
                d_1903_v79_ = nw296_
                index319_ = default__.safeIndex(115, (d_1903_v79_).length(0))
                (d_1903_v79_)[index319_] = d_1902_v78_
                index320_ = default__.safeIndex(115, (d_1903_v79_).length(0))
                (d_1903_v79_)[index320_] = d_1902_v78_
                d_1883_v64_ = (d_1883_v64_) | ((_dafny.MultiSet([d_1902_v78_.f14, (d_1893_v73_).f10])).intersection(d_1883_v64_))
            elif True:
                r1 = (self).f10
                (globalState).f8 = not((self).f10)
                d_1904_v80_: _dafny.Seq
                d_1904_v80_ = _dafny.SeqWithoutIsStrInference([623])
                d_1905_v81_: _dafny.Seq
                d_1905_v81_ = _dafny.SeqWithoutIsStrInference([d_1864_i5_, len(d_1904_v80_), 227])
                d_1906_v82_: _dafny.Array
                nw297_ = _dafny.Array(None, 8)
                nw297_[int(0)] = (self).f10
                nw297_[int(1)] = (self).f10
                nw297_[int(2)] = (d_1793_v0_) == (d_1793_v0_)
                nw297_[int(3)] = (self).f10
                nw297_[int(4)] = (self).f10
                nw297_[int(5)] = (len(d_1904_v80_)) > (len((self).f11))
                nw297_[int(6)] = (self).f10
                nw297_[int(7)] = (self).f10
                d_1906_v82_ = nw297_
                index321_ = default__.safeIndex(176, (d_1906_v82_).length(0))
                (d_1906_v82_)[index321_] = (self).f10
                index322_ = default__.safeIndex(176, (d_1906_v82_).length(0))
                (d_1906_v82_)[index322_] = ((self).f10) or ((self).f10)
                rhs314_ = d_1864_i5_
                rhs315_ = d_1906_v82_
                rhs316_ = (d_1906_v82_)[default__.safeIndex(176, (d_1906_v82_).length(0))]
                rhs317_ = d_1906_v82_
                lhs230_ = globalState
                d_1793_v0_ = rhs314_
                d_1906_v82_ = rhs315_
                lhs230_.f8 = rhs316_
                d_1906_v82_ = rhs317_
                d_1793_v0_ = len(default__.fm54(globalState))
            d_1907_v83_: str
            d_1907_v83_ = _dafny.CodePoint('f')
            d_1908_v84_: _dafny.Seq
            d_1908_v84_ = _dafny.SeqWithoutIsStrInference([(self).fm1(_dafny.Map({(self).f10: (self).f10}), globalState)])
            d_1909_v85_: _dafny.Array
            nw298_ = _dafny.Array(None, 4)
            nw298_[int(0)] = (d_1908_v84_)[default__.safeIndex(d_1793_v0_, len(d_1908_v84_))]
            nw298_[int(1)] = d_1864_i5_
            nw298_[int(2)] = d_1793_v0_
            nw298_[int(3)] = d_1793_v0_
            d_1909_v85_ = nw298_
            d_1910_v86_: _dafny.Map
            d_1910_v86_ = _dafny.Map({d_1909_v85_: (self).f10})
            d_1911_v87_: _dafny.Seq
            d_1911_v87_ = _dafny.SeqWithoutIsStrInference([d_1910_v86_])
            d_1912_v88_: _dafny.Map
            d_1912_v88_ = _dafny.Map({d_1907_v83_: (d_1911_v87_)[default__.safeIndex(d_1793_v0_, len(d_1911_v87_))]})
            d_1913_v89_: _dafny.Map
            d_1913_v89_ = _dafny.Map({(self).f10: d_1909_v85_})
            d_1914_v90_: _dafny.Seq
            d_1914_v90_ = _dafny.SeqWithoutIsStrInference([not((self).f10), (self).f10, not((self).f10), True, (self).f10])
            d_1915_v91_: _dafny.Seq
            d_1915_v91_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uqmwx"))
            d_1916_v92_: _dafny.Map
            d_1916_v92_ = _dafny.Map({d_1864_i5_: (d_1910_v86_).set(((d_1913_v89_)[(self).f10] if ((self).f10) in (d_1913_v89_) else d_1909_v85_), (d_1914_v90_)[default__.safeIndex(len(d_1915_v91_), len(d_1914_v90_))])})
            d_1917_v93_: C5
            nw299_ = C5()
            nw299_.ctor__((self).f9, (self).f10)
            d_1917_v93_ = nw299_
            d_1918_v94_: C6
            nw300_ = C6()
            nw300_.ctor__((self).f9, True)
            d_1918_v94_ = nw300_
            d_1919_v95_: _dafny.Map
            d_1919_v95_ = _dafny.Map({d_1917_v93_: d_1918_v94_})
            d_1920_v96_: D8
            d_1920_v96_ = D8_DC24((self).f10, _dafny.Set({d_1909_v85_, d_1909_v85_}), len(d_1919_v95_))
            if ((((d_1912_v88_)[d_1907_v83_] if (d_1907_v83_) in (d_1912_v88_) else d_1910_v86_)) | (((d_1916_v92_)[(d_1920_v96_).cf44] if ((d_1920_v96_).cf44) in (d_1916_v92_) else d_1910_v86_))) == ((d_1910_v86_) | (d_1910_v86_)):
                (globalState).f8 = (self).f10
                d_1921_v97_: _dafny.Array
                nw301_ = _dafny.Array(False, 18)
                d_1921_v97_ = nw301_
                index323_ = default__.safeIndex(212, (d_1921_v97_).length(0))
                (d_1921_v97_)[index323_] = not((self).f10)
                d_1922_v98_: _dafny.Map
                d_1922_v98_ = _dafny.Map({_dafny.Set({d_1864_i5_}): (self).f10})
                d_1923_v99_: D16
                d_1923_v99_ = D16_DC42((self).f10, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "im"))), d_1864_i5_)
                index324_ = default__.safeIndex(212, (d_1921_v97_).length(0))
                (d_1921_v97_)[index324_] = ((d_1923_v99_ if ((d_1922_v98_)[d_1794_v1_] if (d_1794_v1_) in (d_1922_v98_) else (self).f10) else d_1923_v99_)).cf80
                d_1924_v101_: C2
                nw302_ = C2()
                def iife174_():
                    coll74_ = _dafny.Set()
                    compr_74_: int
                    for compr_74_ in _dafny.IntegerRange(47, 633):
                        d_1925_v100_: int = compr_74_
                        if ((47) <= (d_1925_v100_)) and ((d_1925_v100_) < (633)):
                            coll74_ = coll74_.union(_dafny.Set([(d_1925_v100_) + (d_1793_v0_)]))
                    return _dafny.Set(coll74_)
                nw302_.ctor__((d_1921_v97_)[default__.safeIndex(212, (d_1921_v97_).length(0))], d_1921_v97_, (self).f9, (d_1914_v90_)[default__.safeIndex((0) - ((self).fm3(len(iife174_()
                ), 113, d_1907_v83_, d_1793_v0_, globalState)), len(d_1914_v90_))])
                d_1924_v101_ = nw302_
                d_1926_v102_: _dafny.Map
                d_1926_v102_ = _dafny.Map({(d_1921_v97_)[default__.safeIndex(212, (d_1921_v97_).length(0))]: d_1924_v101_})
                d_1926_v102_ = d_1926_v102_
                d_1793_v0_ = d_1864_i5_
                d_1793_v0_ = default__.safeDivisionInt(d_1864_i5_, d_1793_v0_)
            elif True:
                d_1927_v103_: _dafny.MultiSet
                d_1927_v103_ = _dafny.MultiSet([(self).f10, True, (self).f10, ((self).f10) and ((self).f10)])
                d_1793_v0_ = ((d_1927_v103_)[(self).f10] if ((self).f10) in (d_1927_v103_) else d_1793_v0_)
                d_1909_v85_ = d_1909_v85_
                d_1928_v104_: _dafny.Array
                nw303_ = _dafny.Array(None, 22)
                d_1928_v104_ = nw303_
                d_1929_v105_: D19
                d_1929_v105_ = D19_DC48(d_1928_v104_)
                d_1930_v106_: _dafny.Array
                nw304_ = _dafny.Array(None, 6)
                nw304_[int(0)] = d_1928_v104_
                nw304_[int(1)] = d_1928_v104_
                nw304_[int(2)] = d_1928_v104_
                nw304_[int(3)] = d_1928_v104_
                nw304_[int(4)] = d_1928_v104_
                nw304_[int(5)] = (d_1929_v105_).cf92
                d_1930_v106_ = nw304_
                d_1930_v106_ = d_1930_v106_
                d_1931_v107_: C1
                nw305_ = C1()
                nw305_.ctor__((self).f10, d_1915_v91_)
                d_1931_v107_ = nw305_
                rhs318_ = d_1931_v107_
                rhs319_ = len((d_1908_v84_) + (_dafny.SeqWithoutIsStrInference([(d_1908_v84_)[default__.safeIndex(d_1793_v0_, len(d_1908_v84_))] for d_1932_i10_ in range(default__.abs(313))])))
                rhs320_ = (self).fm1((self).f9, globalState)
                rhs321_ = _dafny.SeqWithoutIsStrInference([13])
                d_1931_v107_ = rhs318_
                d_1793_v0_ = rhs319_
                d_1793_v0_ = rhs320_
                d_1908_v84_ = rhs321_
                index325_ = default__.safeIndex(136, (d_1909_v85_).length(0))
                (d_1909_v85_)[index325_] = (d_1864_i5_) - (d_1793_v0_)
                index326_ = default__.safeIndex(136, (d_1909_v85_).length(0))
                (d_1909_v85_)[index326_] = (0) - (d_1793_v0_)
        d_1933_v108_: _dafny.Seq
        d_1933_v108_ = _dafny.SeqWithoutIsStrInference([(self).f10])
        d_1934_v109_: _dafny.Seq
        d_1934_v109_ = _dafny.SeqWithoutIsStrInference([((self).fm0(d_1933_v108_, (self).f10, d_1793_v0_, (self).f10, globalState)) or ((self).f10)])
        d_1935_v110_: D15
        d_1935_v110_ = D15_DC39(not((self).f10), True)
        pat_let_tv47_ = d_1933_v108_
        pat_let_tv48_ = d_1934_v109_
        pat_let_tv49_ = d_1934_v109_
        def lambda92_(source27_):
            if source27_.is_DC38:
                d_1936___mcc_h6_ = source27_.cf69
                d_1937___mcc_h7_ = source27_.cf70
                d_1938___mcc_h8_ = source27_.cf71
                d_1939___mcc_h9_ = source27_.cf72
                d_1940___mcc_h10_ = source27_.cf73
                d_1941_cf73_ = d_1940___mcc_h10_
                d_1942_cf72_ = d_1939___mcc_h9_
                d_1943_cf71_ = d_1938___mcc_h8_
                d_1944_cf70_ = d_1937___mcc_h7_
                d_1945_cf69_ = d_1936___mcc_h6_
                return (D8_DC23(pat_let_tv47_)).cf41
            elif source27_.is_DC39:
                d_1946___mcc_h11_ = source27_.cf74
                d_1947___mcc_h12_ = source27_.cf75
                d_1948_cf75_ = d_1947___mcc_h12_
                d_1949_cf74_ = d_1946___mcc_h11_
                return pat_let_tv48_
            elif True:
                d_1950___mcc_h13_ = source27_.cf68
                d_1951_cf68_ = d_1950___mcc_h13_
                return pat_let_tv49_

        d_1934_v109_ = lambda92_((d_1935_v110_ if (self).f10 else D15_DC39((self).f10, (self).f10)))
        d_1952_v111_: _dafny.Array
        nw306_ = _dafny.Array(False, 29)
        d_1952_v111_ = nw306_
        guard_loop_8_: int
        for guard_loop_8_ in _dafny.IntegerRange(0, (d_1952_v111_).length(0)):
            d_1953_i11_: int = guard_loop_8_
            if (True) and (((0) <= (d_1953_i11_)) and ((d_1953_i11_) < ((d_1952_v111_).length(0)))):
                (d_1952_v111_)[(d_1953_i11_)] = ((self).f10) or ((self).f10)
        d_1954_i12_: int
        d_1954_i12_ = 0
        with _dafny.label("8"):
            while (self).f10:
                with _dafny.c_label("8"):
                    if (d_1954_i12_) >= (100):
                        raise _dafny.Break("8")
                    d_1954_i12_ = (d_1954_i12_) + (1)
                    d_1955_v112_: C2
                    nw307_ = C2()
                    nw307_.ctor__((self).f10, d_1952_v111_, (self).f9, (self).f10)
                    d_1955_v112_ = nw307_
                    d_1956_v113_: str
                    d_1956_v113_ = _dafny.CodePoint('q')
                    d_1793_v0_ = ((0) - ((d_1793_v0_) * (d_1793_v0_))) + (((self).fm3(len(_dafny.SeqWithoutIsStrInference([d_1955_v112_])), len((self).f11), d_1956_v113_, d_1793_v0_, globalState) if (self).f10 else d_1793_v0_))
                    d_1957_v114_: _dafny.Array
                    nw308_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 16)
                    d_1957_v114_ = nw308_
                    d_1958_v115_: _dafny.Seq
                    d_1958_v115_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fie"))
                    index327_ = default__.safeIndex(571, (d_1957_v114_).length(0))
                    (d_1957_v114_)[index327_] = d_1958_v115_
                    d_1959_v116_: D14
                    d_1959_v116_ = D14_DC36((self).f10, d_1958_v115_, _dafny.SeqWithoutIsStrInference([d_1956_v113_ for d_1960_i13_ in range(default__.abs(425))]))
                    index328_ = default__.safeIndex(571, (d_1957_v114_).length(0))
                    (d_1957_v114_)[index328_] = (d_1958_v115_) + ((default__.fm20(d_1793_v0_, d_1955_v112_.f14, not((self).f10), d_1793_v0_, globalState) if True else (d_1959_v116_).cf66))
                    d_1793_v0_ = d_1793_v0_
                    if (self).f10:
                        d_1793_v0_ = 609
                        d_1961_v117_: C4
                        nw309_ = C4()
                        nw309_.ctor__((self).f9, True)
                        d_1961_v117_ = nw309_
                        d_1962_v118_: _dafny.Map
                        d_1962_v118_ = _dafny.Map({d_1955_v112_.f14: d_1961_v117_})
                        d_1963_v119_: _dafny.MultiSet
                        d_1963_v119_ = _dafny.MultiSet([d_1793_v0_, len(d_1962_v118_)])
                        index329_ = default__.safeIndex(982, ((d_1955_v112_).f15).length(0))
                        ((d_1955_v112_).f15)[index329_] = (_dafny.MultiSet([d_1793_v0_, d_1793_v0_])).ispropersubset(d_1963_v119_)
                        index330_ = default__.safeIndex(982, ((d_1955_v112_).f15).length(0))
                        ((d_1955_v112_).f15)[index330_] = True
                        d_1964_v120_: _dafny.Array
                        def lambda93_(d_1965_v0_):
                            def lambda94_(d_1966_i14_):
                                return default__.safeDivisionInt(d_1966_i14_, d_1965_v0_)

                            return lambda94_

                        init54_ = lambda93_(d_1793_v0_)
                        nw310_ = _dafny.Array(None, 22)
                        for i0_54_ in range(nw310_.length(0)):
                            nw310_[i0_54_] = init54_(i0_54_)
                        d_1964_v120_ = nw310_
                        index331_ = default__.safeIndex(894, (d_1964_v120_).length(0))
                        (d_1964_v120_)[index331_] = (0) - (len(_dafny.SeqWithoutIsStrInference([d_1956_v113_ for d_1967_i15_ in range(default__.abs(687))])))
                        d_1968_v121_: _dafny.Array
                        def lambda95_(d_1969_v0_, d_1970_v1_, d_1971_v112_, d_1972_v113_):
                            def lambda96_(d_1973_i16_):
                                return _dafny.Map({d_1969_v0_: _dafny.Map({_dafny.SeqWithoutIsStrInference([D4_DC11(d_1969_v0_, d_1970_v1_, d_1971_v112_.f14), D4_DC11(d_1969_v0_, d_1970_v1_, ((d_1971_v112_).f15)[default__.safeIndex(982, ((d_1971_v112_).f15).length(0))]), D4_DC11(len(_dafny.SeqWithoutIsStrInference([d_1972_v113_ for d_1974_i17_ in range(default__.abs(847))])), d_1970_v1_, True)]): d_1969_v0_})})

                            return lambda96_

                        init55_ = lambda95_(d_1793_v0_, d_1794_v1_, d_1955_v112_, d_1956_v113_)
                        nw311_ = _dafny.Array(None, 3)
                        for i0_55_ in range(nw311_.length(0)):
                            nw311_[i0_55_] = init55_(i0_55_)
                        d_1968_v121_ = nw311_
                        d_1975_v122_: _dafny.Seq
                        d_1975_v122_ = _dafny.SeqWithoutIsStrInference([D4_DC11(d_1793_v0_, d_1794_v1_, d_1955_v112_.f14) for d_1976_i18_ in range(default__.abs(-656))])
                        d_1977_v123_: _dafny.Map
                        d_1977_v123_ = _dafny.Map({d_1975_v122_: d_1793_v0_})
                        d_1978_v124_: _dafny.Map
                        d_1978_v124_ = _dafny.Map({d_1793_v0_: d_1977_v123_})
                        index332_ = default__.safeIndex(524, (d_1968_v121_).length(0))
                        (d_1968_v121_)[index332_] = d_1978_v124_
                        index333_ = default__.safeIndex(894, (d_1964_v120_).length(0))
                        index334_ = default__.safeIndex(524, (d_1968_v121_).length(0))
                        rhs322_ = (d_1793_v0_) + (d_1793_v0_)
                        rhs323_ = d_1956_v113_
                        rhs324_ = (d_1978_v124_) | (d_1978_v124_)
                        lhs231_ = d_1964_v120_
                        lhs232_ = default__.safeIndex(894, (d_1964_v120_).length(0))
                        lhs233_ = d_1968_v121_
                        lhs234_ = default__.safeIndex(524, (d_1968_v121_).length(0))
                        lhs231_[lhs232_] = rhs322_
                        d_1956_v113_ = rhs323_
                        lhs233_[lhs234_] = rhs324_
                        d_1979_v125_: _dafny.Map
                        d_1979_v125_ = _dafny.Map({313: d_1793_v0_})
                        d_1980_v126_: _dafny.Seq
                        d_1980_v126_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([((d_1979_v125_)[(d_1964_v120_)[default__.safeIndex(894, (d_1964_v120_).length(0))]] if ((d_1964_v120_)[default__.safeIndex(894, (d_1964_v120_).length(0))]) in (d_1979_v125_) else d_1793_v0_), d_1793_v0_]))])
                        d_1981_v127_: _dafny.Map
                        d_1981_v127_ = _dafny.Map({(d_1980_v126_)[default__.safeIndex(len(d_1933_v108_), len(d_1980_v126_))]: True})
                        (d_1955_v112_).f14 = ((d_1981_v127_) | (d_1981_v127_)) != (d_1981_v127_)
                        d_1982_v128_: D6
                        d_1982_v128_ = D6_DC19((self).f10, d_1956_v113_)
                        d_1983_v129_: _dafny.MultiSet
                        d_1983_v129_ = _dafny.MultiSet([d_1982_v128_, d_1982_v128_, d_1982_v128_])
                        d_1984_v130_: _dafny.Seq
                        d_1984_v130_ = _dafny.SeqWithoutIsStrInference([d_1983_v129_])
                        rhs325_ = (d_1957_v114_)[default__.safeIndex(571, (d_1957_v114_).length(0))]
                        rhs326_ = (d_1984_v130_)[default__.safeIndex(d_1793_v0_, len(d_1984_v130_))]
                        d_1958_v115_ = rhs325_
                        d_1983_v129_ = rhs326_
                    elif True:
                        d_1985_v131_: _dafny.Array
                        nw312_ = _dafny.Array(D1.default()(), 23)
                        d_1985_v131_ = nw312_
                        d_1986_v132_: _dafny.Set
                        d_1986_v132_ = _dafny.Set({d_1985_v131_, d_1985_v131_, d_1985_v131_, d_1985_v131_})
                        d_1986_v132_ = d_1986_v132_
                        d_1987_v133_: C5
                        nw313_ = C5()
                        nw313_.ctor__(default__.fm48(d_1955_v112_.f14, globalState), d_1955_v112_.f14)
                        d_1987_v133_ = nw313_
                        d_1988_v134_: _dafny.Array
                        nw314_ = _dafny.Array(int(0), 2)
                        d_1988_v134_ = nw314_
                        d_1988_v134_ = d_1988_v134_
                        d_1989_v135_: _dafny.Set
                        d_1989_v135_ = _dafny.Set({(d_1957_v114_)[default__.safeIndex(571, (d_1957_v114_).length(0))]})
                        d_1989_v135_ = (d_1989_v135_) - (d_1989_v135_)
                        d_1990_v136_: _dafny.Array
                        nw315_ = _dafny.Array(_dafny.CodePoint('D'), 15)
                        d_1990_v136_ = nw315_
                        index335_ = default__.safeIndex(781, (d_1990_v136_).length(0))
                        (d_1990_v136_)[index335_] = _dafny.CodePoint('a')
                        index336_ = default__.safeIndex(781, (d_1990_v136_).length(0))
                        (d_1990_v136_)[index336_] = d_1956_v113_
                    pass
            pass
        r1 = (self).f10
        d_1991_v137_: _dafny.Seq
        d_1991_v137_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y"))
        r0 = ((d_1991_v137_) + (d_1991_v137_)) + (d_1991_v137_)
        d_1992_v138_: _dafny.Seq
        d_1992_v138_ = _dafny.SeqWithoutIsStrInference([d_1793_v0_])
        d_1993_v139_: str
        d_1993_v139_ = _dafny.CodePoint('f')
        d_1994_v140_: _dafny.Map
        d_1994_v140_ = _dafny.Map({d_1793_v0_: (self).f10})
        d_1995_v141_: D14
        d_1995_v141_ = D14_DC35(len(d_1994_v140_), (self).f10, d_1993_v139_)
        r1 = (((0) - (d_1793_v0_)) + ((self).fm3(len(d_1992_v138_), len(d_1934_v109_), d_1993_v139_, d_1793_v0_, globalState))) != ((0) - ((d_1995_v141_).cf62))
        d_1996_v142_: D7
        d_1996_v142_ = D7_DC21((self).f11)
        d_1997_v143_: _dafny.Seq
        def iife175_(_pat_let50_0):
            def iife176_(d_1998_dt__update__tmp_h3_):
                def iife177_(_pat_let51_0):
                    def iife178_(d_1999_dt__update_hcf40_h0_):
                        return D7_DC21(d_1999_dt__update_hcf40_h0_)
                    return iife178_(_pat_let51_0)
                return iife177_(_dafny.Set({(self).f10}))
            return iife176_(_pat_let50_0)
        d_1997_v143_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({(self).f10, (self).f10, (self).f10}), (self).f11, (iife175_(d_1996_v142_)).cf40])
        d_2000_v144_: _dafny.Seq
        d_2000_v144_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(self).f11 for d_2001_i19_ in range(default__.abs(375))]), ((d_1997_v143_) + (d_1997_v143_)).set(default__.safeIndex(len(_dafny.Map({True: (self).f9})), len((d_1997_v143_) + (d_1997_v143_))), (self).f11), d_1997_v143_, _dafny.SeqWithoutIsStrInference([(self).f11, (self).f11]), _dafny.SeqWithoutIsStrInference([(self).f11, _dafny.Set({(self).f10}), (self).f11, (self).f11])])
        r2 = (d_2000_v144_)[default__.safeIndex(default__.safeModuloInt(d_1793_v0_, len(d_1992_v138_)), len(d_2000_v144_))]
        r3 = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xjnxgcyay"))
        return r0, r1, r2, r3

    def m1(self, p0, globalState):
        r0: _dafny.Map = _dafny.Map({})
        r1: bool = False
        d_2002_v0_: _dafny.Array
        nw316_ = _dafny.Array(_dafny.CodePoint('D'), 2)
        d_2002_v0_ = nw316_
        guard_loop_9_: int
        for guard_loop_9_ in _dafny.IntegerRange(0, (d_2002_v0_).length(0)):
            d_2003_i0_: int = guard_loop_9_
            if (True) and (((0) <= (d_2003_i0_)) and ((d_2003_i0_) < ((d_2002_v0_).length(0)))):
                (d_2002_v0_)[(d_2003_i0_)] = p0
        if (self).f10:
            d_2004_v1_: _dafny.Array
            nw317_ = _dafny.Array(False, 8)
            d_2004_v1_ = nw317_
            d_2004_v1_ = d_2004_v1_
            d_2005_v2_: str
            d_2005_v2_ = _dafny.CodePoint('y')
            d_2005_v2_ = p0
            r1 = ((self).f10 if (self).f10 else True)
            d_2006_v3_: _dafny.Array
            nw318_ = _dafny.Array(_dafny.Array(None, 0), 6)
            d_2006_v3_ = nw318_
            d_2006_v3_ = d_2006_v3_
            d_2007_v4_: int
            d_2007_v4_ = 128
            d_2007_v4_ = (662 if (d_2007_v4_) != (d_2007_v4_) else d_2007_v4_)
        elif True:
            d_2008_v5_: int
            d_2008_v5_ = 817
            d_2008_v5_ = d_2008_v5_
            if (self).f10:
                d_2009_v6_: _dafny.Array
                def lambda97_(d_2010_i1_):
                    return (self).f10

                init56_ = lambda97_
                nw319_ = _dafny.Array(None, 6)
                for i0_56_ in range(nw319_.length(0)):
                    nw319_[i0_56_] = init56_(i0_56_)
                d_2009_v6_ = nw319_
                d_2011_v7_: _dafny.Array
                nw320_ = _dafny.Array(None, 2)
                nw320_[int(0)] = d_2009_v6_
                nw320_[int(1)] = d_2009_v6_
                d_2011_v7_ = nw320_
                index337_ = default__.safeIndex(517, (d_2009_v6_).length(0))
                (d_2009_v6_)[index337_] = (self).f10
                d_2012_v8_: _dafny.Seq
                d_2012_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gfnvwqdiy"))
                d_2013_v9_: D14
                d_2013_v9_ = D14_DC36(not((self).f10), d_2012_v8_, d_2012_v8_)
                index338_ = default__.safeIndex(517, (d_2009_v6_).length(0))
                rhs327_ = d_2011_v7_
                rhs328_ = (d_2013_v9_).cf65
                rhs329_ = not((self).f10)
                rhs330_ = -820
                lhs235_ = d_2009_v6_
                lhs236_ = default__.safeIndex(517, (d_2009_v6_).length(0))
                lhs237_ = globalState
                d_2011_v7_ = rhs327_
                lhs235_[lhs236_] = rhs328_
                lhs237_.f8 = rhs329_
                d_2008_v5_ = rhs330_
                d_2008_v5_ = 557
                r1 = (self).f10
                d_2014_v10_: _dafny.Seq
                d_2014_v10_ = _dafny.SeqWithoutIsStrInference([(d_2009_v6_)[default__.safeIndex(517, (d_2009_v6_).length(0))]])
                (globalState).f8 = (_dafny.Map({d_2008_v5_: (d_2014_v10_).set(default__.safeIndex(d_2008_v5_, len(d_2014_v10_)), (self).fm2(globalState))})) == (_dafny.Map({(self).fm1((self).f9, globalState): (d_2014_v10_).set(default__.safeIndex(d_2008_v5_, len(d_2014_v10_)), (self).f10)}))
                (globalState).f8 = (D15_DC39((self).f10, True)).cf74
            elif True:
                d_2008_v5_ = d_2008_v5_
                (globalState).f8 = not((self).f10)
                (globalState).f8 = (self).f10
                d_2015_v11_: _dafny.Array
                nw321_ = _dafny.Array(int(0), 8)
                d_2015_v11_ = nw321_
                index339_ = default__.safeIndex(655, (d_2015_v11_).length(0))
                (d_2015_v11_)[index339_] = (0) - ((d_2008_v5_) - (d_2008_v5_))
                index340_ = default__.safeIndex(655, (d_2015_v11_).length(0))
                (d_2015_v11_)[index340_] = 224
                d_2016_v12_: _dafny.Seq
                d_2016_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uppopviv"))
                (globalState).f8 = not((default__.safeModuloInt(len(d_2016_v12_), len(_dafny.Set({(self).f10, not((self).f10)})))) < (len(d_2016_v12_)))
            d_2008_v5_ = default__.safeModuloInt(d_2008_v5_, d_2008_v5_)
            rhs331_ = d_2008_v5_
            rhs332_ = d_2008_v5_
            rhs333_ = (self).f10
            lhs238_ = globalState
            d_2008_v5_ = rhs331_
            d_2008_v5_ = rhs332_
            lhs238_.f8 = rhs333_
            d_2017_v13_: _dafny.Seq
            d_2017_v13_ = _dafny.SeqWithoutIsStrInference([(self).f10, True])
            if not (not(not((d_2017_v13_) <= (_dafny.SeqWithoutIsStrInference([(self).f10]))))) or ((self).f10):
                d_2018_v14_: _dafny.Array
                def lambda98_(d_2019_i2_):
                    return default__.safeModuloInt(d_2019_i2_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eqk"))))

                init57_ = lambda98_
                nw322_ = _dafny.Array(None, 13)
                for i0_57_ in range(nw322_.length(0)):
                    nw322_[i0_57_] = init57_(i0_57_)
                d_2018_v14_ = nw322_
                d_2018_v14_ = d_2018_v14_
                d_2020_v15_: D0
                d_2020_v15_ = D0_DC0((self).f10)
                d_2021_v16_: _dafny.Map
                d_2021_v16_ = _dafny.Map({not((d_2020_v15_).cf0): d_2008_v5_})
                d_2022_v17_: _dafny.Map
                d_2022_v17_ = _dafny.Map({d_2008_v5_: True})
                d_2021_v16_ = (d_2021_v16_).set(not(((d_2022_v17_)[d_2008_v5_] if (d_2008_v5_) in (d_2022_v17_) else (self).f10)), default__.fm33((self).f10, (d_2017_v13_)[default__.safeIndex(d_2008_v5_, len(d_2017_v13_))], 704, not((self).f10), globalState))
                d_2023_v18_: C3
                nw323_ = C3()
                nw323_.ctor__()
                d_2023_v18_ = nw323_
                d_2024_v19_: C4
                nw324_ = C4()
                nw324_.ctor__((self).f9, (self).f10)
                d_2024_v19_ = nw324_
                d_2025_v20_: _dafny.Map
                d_2025_v20_ = _dafny.Map({(self).f10: d_2024_v19_})
                d_2026_v21_: _dafny.Map
                d_2026_v21_ = _dafny.Map({d_2023_v18_: d_2025_v20_})
                d_2026_v21_ = (d_2026_v21_).set(d_2023_v18_, d_2025_v20_)
                d_2027_v22_: str
                d_2027_v22_ = _dafny.CodePoint('r')
                d_2027_v22_ = d_2027_v22_
                d_2028_v23_: _dafny.Seq
                d_2028_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sta"))
                d_2029_v24_: _dafny.Seq
                d_2029_v24_ = _dafny.SeqWithoutIsStrInference([d_2028_v23_])
                d_2030_v25_: str
                d_2030_v25_ = d_2027_v22_
                d_2031_v26_: D15
                d_2031_v26_ = D15_DC38(not((self).f10), (d_2029_v24_)[default__.safeIndex(d_2008_v5_, len(d_2029_v24_))], True, (self).f10, d_2030_v25_)
                d_2032_v27_: _dafny.Seq
                d_2032_v27_ = _dafny.SeqWithoutIsStrInference([d_2022_v17_, d_2022_v17_, d_2022_v17_, d_2022_v17_])
                d_2033_v28_: _dafny.Array
                nw325_ = _dafny.Array(False, 5)
                d_2033_v28_ = nw325_
                d_2034_v29_: _dafny.Map
                d_2034_v29_ = _dafny.Map({(self).f10: d_2033_v28_})
                d_2035_v30_: C2
                nw326_ = C2()
                nw326_.ctor__((self).f10, ((d_2034_v29_)[(self).f10] if ((self).f10) in (d_2034_v29_) else d_2033_v28_), _dafny.Map({(self).f10: (self).f10}), True)
                d_2035_v30_ = nw326_
                d_2036_v31_: _dafny.MultiSet
                d_2036_v31_ = _dafny.MultiSet([d_2035_v30_])
                d_2037_v32_: D4
                d_2037_v32_ = D4_DC12((self).f10, d_2008_v5_, default__.fm31((self).f10, (d_2032_v27_)[default__.safeIndex((d_2036_v31_).cardinality, len(d_2032_v27_))], globalState), (0) - (d_2008_v5_), len(d_2028_v23_))
                d_2038_v33_: _dafny.Map
                d_2038_v33_ = _dafny.Map({d_2031_v26_: d_2037_v32_})
                d_2039_v34_: D6
                d_2039_v34_ = D6_DC18(p0, d_2008_v5_, d_2008_v5_)
                d_2040_v35_: D6
                d_2040_v35_ = D6_DC20(d_2039_v34_)
                d_2041_v36_: _dafny.Seq
                d_2041_v36_ = _dafny.SeqWithoutIsStrInference([default__.fm55((self).f10, globalState)])
                d_2042_v37_: _dafny.Array
                nw327_ = _dafny.Array(None, 22)
                nw327_[int(0)] = (self).f10
                nw327_[int(1)] = (self).f10
                nw327_[int(2)] = (self).f10
                nw327_[int(3)] = not((self).f10)
                nw327_[int(4)] = (self).f10
                nw327_[int(5)] = not((self).f10)
                nw327_[int(6)] = (self).f10
                nw327_[int(7)] = (self).f10
                nw327_[int(8)] = (self).f10
                nw327_[int(9)] = (d_2038_v33_) != (_dafny.Map({d_2031_v26_: d_2037_v32_}))
                nw327_[int(10)] = (self).f10
                nw327_[int(11)] = (self).f10
                nw327_[int(12)] = (d_2008_v5_) > (len((self).f9))
                nw327_[int(13)] = not((d_2035_v30_.f14) or ((self).f10))
                nw327_[int(14)] = (self).f10
                nw327_[int(15)] = (d_2035_v30_.f14) or (d_2035_v30_.f14)
                nw327_[int(16)] = (D6_DC20(d_2040_v35_)) not in (d_2041_v36_)
                nw327_[int(17)] = (self).f10
                nw327_[int(18)] = not((False) not in ((_dafny.MultiSet(d_2017_v13_)).set(d_2035_v30_.f14, default__.abs(d_2008_v5_))))
                nw327_[int(19)] = d_2035_v30_.f14
                nw327_[int(20)] = (self).f10
                nw327_[int(21)] = d_2035_v30_.f14
                d_2042_v37_ = nw327_
                d_2043_v38_: _dafny.MultiSet
                d_2043_v38_ = _dafny.MultiSet([d_2008_v5_])
                index341_ = default__.safeIndex(123, (d_2042_v37_).length(0))
                (d_2042_v37_)[index341_] = (d_2043_v38_).isdisjoint(d_2043_v38_)
                index342_ = default__.safeIndex(123, (d_2042_v37_).length(0))
                (d_2042_v37_)[index342_] = (self).fm0(default__.fm4(_dafny.CodePoint('o'), globalState), d_2035_v30_.f14, d_2008_v5_, (self).f10, globalState)
            elif True:
                d_2044_v39_: _dafny.Array
                nw328_ = _dafny.Array(_dafny.Array(None, 0), 16)
                d_2044_v39_ = nw328_
                d_2045_v40_: _dafny.Array
                def lambda99_(d_2046_v5_):
                    def lambda100_(d_2047_i3_):
                        return (d_2047_i3_) * (len(_dafny.Map({d_2046_v5_: (self).f10})))

                    return lambda100_

                init58_ = lambda99_(d_2008_v5_)
                nw329_ = _dafny.Array(None, 26)
                for i0_58_ in range(nw329_.length(0)):
                    nw329_[i0_58_] = init58_(i0_58_)
                d_2045_v40_ = nw329_
                index343_ = default__.safeIndex(893, (d_2044_v39_).length(0))
                (d_2044_v39_)[index343_] = d_2045_v40_
                d_2048_v41_: _dafny.Map
                d_2048_v41_ = _dafny.Map({(self).f10: d_2045_v40_})
                index344_ = default__.safeIndex(893, (d_2044_v39_).length(0))
                (d_2044_v39_)[index344_] = ((d_2048_v41_)[(self).f10] if ((self).f10) in (d_2048_v41_) else d_2045_v40_)
                d_2049_v42_: D6
                d_2049_v42_ = D6_DC18(p0, d_2008_v5_, len((self).f9))
                arr4_ = (d_2044_v39_)[default__.safeIndex(893, (d_2044_v39_).length(0))]
                index345_ = default__.safeIndex(675, ((d_2044_v39_)[default__.safeIndex(893, (d_2044_v39_).length(0))]).length(0))
                arr4_[index345_] = (d_2049_v42_).cf35
                d_2050_v43_: C4
                nw330_ = C4()
                nw330_.ctor__((self).f9, (self).f10)
                d_2050_v43_ = nw330_
                d_2051_v44_: _dafny.Map
                d_2051_v44_ = _dafny.Map({d_2008_v5_: d_2050_v43_})
                d_2052_v45_: _dafny.Array
                nw331_ = _dafny.Array(None, 16)
                nw331_[int(0)] = d_2050_v43_
                nw331_[int(1)] = d_2050_v43_
                nw331_[int(2)] = d_2050_v43_
                nw331_[int(3)] = d_2050_v43_
                nw331_[int(4)] = d_2050_v43_
                nw331_[int(5)] = d_2050_v43_
                nw331_[int(6)] = d_2050_v43_
                nw331_[int(7)] = d_2050_v43_
                nw331_[int(8)] = d_2050_v43_
                nw331_[int(9)] = ((d_2051_v44_)[d_2008_v5_] if (d_2008_v5_) in (d_2051_v44_) else d_2050_v43_)
                nw331_[int(10)] = d_2050_v43_
                nw331_[int(11)] = d_2050_v43_
                nw331_[int(12)] = d_2050_v43_
                nw331_[int(13)] = d_2050_v43_
                nw331_[int(14)] = d_2050_v43_
                nw331_[int(15)] = d_2050_v43_
                d_2052_v45_ = nw331_
                index346_ = default__.safeIndex(202, (d_2052_v45_).length(0))
                (d_2052_v45_)[index346_] = d_2050_v43_
                arr5_ = (d_2044_v39_)[default__.safeIndex(893, (d_2044_v39_).length(0))]
                index347_ = default__.safeIndex(675, ((d_2044_v39_)[default__.safeIndex(893, (d_2044_v39_).length(0))]).length(0))
                index348_ = default__.safeIndex(202, (d_2052_v45_).length(0))
                rhs334_ = (self).f10
                rhs335_ = d_2008_v5_
                rhs336_ = d_2050_v43_
                lhs239_ = globalState
                lhs240_ = (d_2044_v39_)[default__.safeIndex(893, (d_2044_v39_).length(0))]
                lhs241_ = default__.safeIndex(675, ((d_2044_v39_)[default__.safeIndex(893, (d_2044_v39_).length(0))]).length(0))
                lhs242_ = d_2052_v45_
                lhs243_ = default__.safeIndex(202, (d_2052_v45_).length(0))
                lhs239_.f8 = rhs334_
                lhs240_[lhs241_] = rhs335_
                lhs242_[lhs243_] = rhs336_
                d_2053_v46_: _dafny.Array
                def lambda101_(d_2054_i4_):
                    return not((self).f10)

                init59_ = lambda101_
                nw332_ = _dafny.Array(None, 2)
                for i0_59_ in range(nw332_.length(0)):
                    nw332_[i0_59_] = init59_(i0_59_)
                d_2053_v46_ = nw332_
                index349_ = default__.safeIndex(976, (d_2053_v46_).length(0))
                (d_2053_v46_)[index349_] = (self).f10
                index350_ = default__.safeIndex(976, (d_2053_v46_).length(0))
                (d_2053_v46_)[index350_] = (210) <= (((d_2044_v39_)[default__.safeIndex(893, (d_2044_v39_).length(0))])[default__.safeIndex(675, ((d_2044_v39_)[default__.safeIndex(893, (d_2044_v39_).length(0))]).length(0))])
                r1 = (d_2053_v46_)[default__.safeIndex(976, (d_2053_v46_).length(0))]
                def lambda102_(d_2055_v39_, d_2056_v5_):
                    def lambda103_(d_2057_i5_):
                        def iife179_():
                            coll75_ = _dafny.Map()
                            compr_75_: int
                            for compr_75_ in (_dafny.SeqWithoutIsStrInference([d_2056_v5_ for d_2058_i6_ in range(default__.abs(239))])).Elements:
                                d_2059_v47_: int = compr_75_
                                if (d_2059_v47_) in (_dafny.SeqWithoutIsStrInference([d_2056_v5_ for d_2058_i6_ in range(default__.abs(239))])):
                                    coll75_[default__.safeModuloInt(d_2059_v47_, ((d_2055_v39_)[default__.safeIndex(893, (d_2055_v39_).length(0))])[default__.safeIndex(675, ((d_2055_v39_)[default__.safeIndex(893, (d_2055_v39_).length(0))]).length(0))])] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cxlqhjcb")))
                            return _dafny.Map(coll75_)
                        return (671) not in (iife179_()
                        )

                    return lambda103_

                init60_ = lambda102_(d_2044_v39_, d_2008_v5_)
                nw333_ = _dafny.Array(None, 19)
                for i0_60_ in range(nw333_.length(0)):
                    nw333_[i0_60_] = init60_(i0_60_)
                d_2053_v46_ = nw333_
        d_2060_v48_: int
        d_2060_v48_ = -908
        rhs337_ = d_2060_v48_
        rhs338_ = (not ((self).fm2(globalState)) or ((self).f10)) or ((self).f10)
        d_2060_v48_ = rhs337_
        r1 = rhs338_
        d_2061_v49_: _dafny.Seq
        d_2061_v49_ = _dafny.SeqWithoutIsStrInference([(self).f10])
        if (d_2061_v49_)[default__.safeIndex(d_2060_v48_, len(d_2061_v49_))]:
            d_2062_v50_: _dafny.Array
            def lambda104_(d_2063_i7_):
                return (self).f10

            init61_ = lambda104_
            nw334_ = _dafny.Array(None, 1)
            for i0_61_ in range(nw334_.length(0)):
                nw334_[i0_61_] = init61_(i0_61_)
            d_2062_v50_ = nw334_
            d_2064_v51_: _dafny.MultiSet
            d_2064_v51_ = _dafny.MultiSet([(self).f10])
            index351_ = default__.safeIndex(191, (d_2062_v50_).length(0))
            (d_2062_v50_)[index351_] = (d_2064_v51_).issubset(_dafny.MultiSet(d_2061_v49_))
            d_2065_v52_: _dafny.Array
            def lambda105_(d_2066_v49_):
                def lambda106_(d_2067_i8_):
                    return _dafny.Map({(self).f10: d_2066_v49_})

                return lambda106_

            init62_ = lambda105_(d_2061_v49_)
            nw335_ = _dafny.Array(None, 4)
            for i0_62_ in range(nw335_.length(0)):
                nw335_[i0_62_] = init62_(i0_62_)
            d_2065_v52_ = nw335_
            d_2068_v53_: _dafny.Map
            d_2068_v53_ = _dafny.Map({(self).f10: d_2061_v49_})
            index352_ = default__.safeIndex(181, (d_2065_v52_).length(0))
            (d_2065_v52_)[index352_] = (d_2068_v53_) | (d_2068_v53_)
            d_2069_v54_: _dafny.Map
            d_2069_v54_ = _dafny.Map({(self).f10: (0) - (d_2060_v48_)})
            index353_ = default__.safeIndex(191, (d_2062_v50_).length(0))
            index354_ = default__.safeIndex(181, (d_2065_v52_).length(0))
            rhs339_ = not ((self).f10) or ((self).f10)
            rhs340_ = ((self).f10) == ((d_2060_v48_) > (len(d_2069_v54_)))
            rhs341_ = ((d_2068_v53_) | ((d_2068_v53_).set(not((self).fm2(globalState)), d_2061_v49_))) | ((d_2068_v53_).set((self).f10, d_2061_v49_))
            rhs342_ = (d_2060_v48_) - (d_2060_v48_)
            lhs244_ = d_2062_v50_
            lhs245_ = default__.safeIndex(191, (d_2062_v50_).length(0))
            lhs246_ = d_2065_v52_
            lhs247_ = default__.safeIndex(181, (d_2065_v52_).length(0))
            lhs244_[lhs245_] = rhs339_
            r1 = rhs340_
            lhs246_[lhs247_] = rhs341_
            d_2060_v48_ = rhs342_
            d_2070_v55_: D6
            d_2070_v55_ = D6_DC16((self).f9)
            d_2071_v56_: _dafny.Array
            nw336_ = _dafny.Array(None, 15)
            nw336_[int(0)] = (self).f9
            nw336_[int(1)] = ((self).f9) | ((self).f9)
            nw336_[int(2)] = (default__.fm48((d_2062_v50_)[default__.safeIndex(191, (d_2062_v50_).length(0))], globalState)).set((d_2062_v50_)[default__.safeIndex(191, (d_2062_v50_).length(0))], (self).f10)
            nw336_[int(3)] = _dafny.Map({(d_2062_v50_)[default__.safeIndex(191, (d_2062_v50_).length(0))]: (d_2062_v50_)[default__.safeIndex(191, (d_2062_v50_).length(0))]})
            nw336_[int(4)] = ((self).f9).set((d_2062_v50_)[default__.safeIndex(191, (d_2062_v50_).length(0))], (d_2062_v50_)[default__.safeIndex(191, (d_2062_v50_).length(0))])
            nw336_[int(5)] = (self).f9
            nw336_[int(6)] = ((self).f9) | ((self).f9)
            nw336_[int(7)] = (d_2070_v55_).cf28
            nw336_[int(8)] = (self).f9
            nw336_[int(9)] = ((self).f9).set((d_2062_v50_)[default__.safeIndex(191, (d_2062_v50_).length(0))], (self).f10)
            nw336_[int(10)] = (self).f9
            nw336_[int(11)] = (self).f9
            nw336_[int(12)] = ((self).f9) | ((self).f9)
            nw336_[int(13)] = _dafny.Map({not((d_2062_v50_)[default__.safeIndex(191, (d_2062_v50_).length(0))]): (d_2062_v50_)[default__.safeIndex(191, (d_2062_v50_).length(0))]})
            nw336_[int(14)] = (self).f9
            d_2071_v56_ = nw336_
            d_2072_v57_: _dafny.Seq
            d_2072_v57_ = _dafny.SeqWithoutIsStrInference([(self).f9])
            index355_ = default__.safeIndex(801, (d_2071_v56_).length(0))
            (d_2071_v56_)[index355_] = ((d_2072_v57_)[default__.safeIndex(d_2060_v48_, len(d_2072_v57_))]) | ((self).f9)
            index356_ = default__.safeIndex(801, (d_2071_v56_).length(0))
            (d_2071_v56_)[index356_] = (self).f9
            d_2073_v58_: D15
            d_2073_v58_ = D15_DC39(not ((self).f10) or (not((d_2062_v50_)[default__.safeIndex(191, (d_2062_v50_).length(0))])), (d_2062_v50_)[default__.safeIndex(191, (d_2062_v50_).length(0))])
            source28_ = d_2073_v58_
            if source28_.is_DC38:
                d_2074___mcc_h0_ = source28_.cf69
                d_2075___mcc_h1_ = source28_.cf70
                d_2076___mcc_h2_ = source28_.cf71
                d_2077___mcc_h3_ = source28_.cf72
                d_2078___mcc_h4_ = source28_.cf73
                d_2079_cf73_ = d_2078___mcc_h4_
                d_2080_cf72_ = d_2077___mcc_h3_
                d_2081_cf71_ = d_2076___mcc_h2_
                d_2082_cf70_ = d_2075___mcc_h1_
                d_2083_cf69_ = d_2074___mcc_h0_
                d_2084_v59_: _dafny.Map
                d_2084_v59_ = _dafny.Map({d_2062_v50_: d_2080_cf72_})
                d_2084_v59_ = ((d_2084_v59_) | ((d_2084_v59_).set(d_2062_v50_, not((self).f10)))) | (d_2084_v59_)
                d_2081_cf71_ = d_2080_cf72_
                d_2085_v60_: C0
                nw337_ = C0()
                nw337_.ctor__()
                d_2085_v60_ = nw337_
                d_2060_v48_ = default__.safeDivisionInt((0) - (d_2060_v48_), d_2060_v48_)
            elif source28_.is_DC39:
                d_2086___mcc_h5_ = source28_.cf74
                d_2087___mcc_h6_ = source28_.cf75
                d_2088_cf75_ = d_2087___mcc_h6_
                d_2089_cf74_ = d_2086___mcc_h5_
                index357_ = default__.safeIndex(191, (d_2062_v50_).length(0))
                (d_2062_v50_)[index357_] = (self).f10
                d_2090_v61_: _dafny.Seq
                d_2090_v61_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dsijjtbga"))
                d_2091_v62_: C1
                nw338_ = C1()
                nw338_.ctor__((self).fm2(globalState), d_2090_v61_)
                d_2091_v62_ = nw338_
                d_2092_v63_: _dafny.Array
                nw339_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 16)
                d_2092_v63_ = nw339_
                index358_ = default__.safeIndex(295, (d_2092_v63_).length(0))
                (d_2092_v63_)[index358_] = d_2090_v61_
                d_2093_v64_: _dafny.Map
                d_2093_v64_ = _dafny.Map({p0: (d_2091_v62_).f13})
                index359_ = default__.safeIndex(295, (d_2092_v63_).length(0))
                (d_2092_v63_)[index359_] = ((((d_2093_v64_)[p0] if (p0) in (d_2093_v64_) else _dafny.SeqWithoutIsStrInference([p0 for d_2094_i9_ in range(default__.abs(-651))]))) + (d_2090_v61_)) + ((d_2091_v62_).f13)
                d_2095_v65_: _dafny.Seq
                d_2096_v66_: bool
                d_2097_v67_: _dafny.Seq
                d_2098_v68_: _dafny.Seq
                out63_: _dafny.Seq
                out64_: bool
                out65_: _dafny.Seq
                out66_: _dafny.Seq
                out63_, out64_, out65_, out66_ = (d_2091_v62_).m0(globalState)
                d_2095_v65_ = out63_
                d_2096_v66_ = out64_
                d_2097_v67_ = out65_
                d_2098_v68_ = out66_
            elif True:
                d_2099___mcc_h7_ = source28_.cf68
                d_2100_cf68_ = d_2099___mcc_h7_
                d_2101_v69_: str
                d_2101_v69_ = _dafny.CodePoint('r')
                d_2102_v70_: _dafny.Map
                d_2102_v70_ = _dafny.Map({d_2060_v48_: (self).f10})
                d_2101_v69_ = default__.fm31((d_2062_v50_)[default__.safeIndex(191, (d_2062_v50_).length(0))], d_2102_v70_, globalState)
                d_2103_v71_: _dafny.Set
                d_2103_v71_ = _dafny.Set({d_2060_v48_, d_2060_v48_})
                d_2104_v72_: _dafny.Map
                d_2104_v72_ = _dafny.Map({d_2103_v71_: d_2002_v0_})
                d_2060_v48_ = len((d_2104_v72_) | (d_2104_v72_))
                r1 = (d_2062_v50_)[default__.safeIndex(191, (d_2062_v50_).length(0))]
                d_2105_v73_: _dafny.Array
                nw340_ = _dafny.Array(int(0), 18)
                d_2105_v73_ = nw340_
                d_2106_v74_: _dafny.Map
                d_2106_v74_ = _dafny.Map({d_2060_v48_: d_2105_v73_})
                d_2106_v74_ = (d_2106_v74_).set((d_2060_v48_) * (d_2060_v48_), d_2105_v73_)
            d_2060_v48_ = d_2060_v48_
            d_2107_v75_: _dafny.Array
            def lambda107_(d_2108_v48_):
                def lambda108_(d_2109_i10_):
                    return (d_2109_i10_) + (d_2108_v48_)

                return lambda108_

            init63_ = lambda107_(d_2060_v48_)
            nw341_ = _dafny.Array(None, 28)
            for i0_63_ in range(nw341_.length(0)):
                nw341_[i0_63_] = init63_(i0_63_)
            d_2107_v75_ = nw341_
            d_2110_v76_: _dafny.Map
            d_2110_v76_ = _dafny.Map({d_2062_v50_: d_2107_v75_})
            d_2111_v77_: _dafny.Map
            d_2111_v77_ = _dafny.Map({d_2107_v75_: (d_2110_v76_) | (d_2110_v76_)})
            d_2111_v77_ = (d_2111_v77_).set(d_2107_v75_, _dafny.Map({d_2062_v50_: d_2107_v75_}))
        elif True:
            d_2112_v78_: D0
            d_2112_v78_ = D0_DC1(598, (self).f10, (0) - (d_2060_v48_))
            d_2113_v79_: _dafny.Map
            d_2113_v79_ = _dafny.Map({d_2060_v48_: (self).f11})
            d_2114_v80_: _dafny.Array
            nw342_ = _dafny.Array(None, 22)
            nw342_[int(0)] = (d_2112_v78_).cf2
            nw342_[int(1)] = (self).f10
            nw342_[int(2)] = (self).f10
            nw342_[int(3)] = (self).f10
            nw342_[int(4)] = (self).f10
            nw342_[int(5)] = (self).f10
            nw342_[int(6)] = (self).f10
            nw342_[int(7)] = (self).f10
            nw342_[int(8)] = (self).f10
            nw342_[int(9)] = (self).f10
            nw342_[int(10)] = not((self).f10)
            nw342_[int(11)] = (self).fm2(globalState)
            nw342_[int(12)] = (self).f10
            nw342_[int(13)] = not(not((self).f10))
            nw342_[int(14)] = (self).f10
            nw342_[int(15)] = (self).f10
            nw342_[int(16)] = not((self).f10)
            nw342_[int(17)] = (self).f10
            nw342_[int(18)] = ((self).f11).isdisjoint(((d_2113_v79_)[d_2060_v48_] if (d_2060_v48_) in (d_2113_v79_) else (self).f11))
            nw342_[int(19)] = False
            nw342_[int(20)] = (self).f10
            nw342_[int(21)] = (self).f10
            d_2114_v80_ = nw342_
            index360_ = default__.safeIndex(524, (d_2114_v80_).length(0))
            (d_2114_v80_)[index360_] = (self).f10
            index361_ = default__.safeIndex(524, (d_2114_v80_).length(0))
            (d_2114_v80_)[index361_] = ((self).f10) or (False)
            d_2115_v81_: str
            d_2115_v81_ = _dafny.CodePoint('x')
            d_2115_v81_ = p0
            d_2116_v82_: _dafny.Seq
            d_2116_v82_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tf"))
            if (d_2116_v82_) <= (d_2116_v82_):
                d_2117_v83_: _dafny.Array
                nw343_ = _dafny.Array(_dafny.Array(None, 0), 1)
                d_2117_v83_ = nw343_
                index362_ = default__.safeIndex(268, (d_2117_v83_).length(0))
                (d_2117_v83_)[index362_] = d_2114_v80_
                index363_ = default__.safeIndex(268, (d_2117_v83_).length(0))
                (d_2117_v83_)[index363_] = d_2114_v80_
                d_2118_v84_: C3
                nw344_ = C3()
                nw344_.ctor__()
                d_2118_v84_ = nw344_
                d_2060_v48_ = default__.safeDivisionInt(d_2060_v48_, d_2060_v48_)
                d_2119_v85_: _dafny.MultiSet
                d_2119_v85_ = _dafny.MultiSet([997, len(_dafny.Map({(self).f10: d_2060_v48_}))])
                d_2120_v86_: _dafny.Map
                d_2120_v86_ = _dafny.Map({(d_2119_v85_).set(d_2060_v48_, default__.abs(d_2060_v48_)): d_2061_v49_})
                index364_ = default__.safeIndex(524, (d_2114_v80_).length(0))
                (d_2114_v80_)[index364_] = ((d_2060_v48_) <= (d_2060_v48_)) or (not(not((((d_2120_v86_)[d_2119_v85_] if (d_2119_v85_) in (d_2120_v86_) else d_2061_v49_)) <= (d_2061_v49_))))
                d_2121_v87_: _dafny.Seq
                d_2121_v87_ = _dafny.SeqWithoutIsStrInference([d_2060_v48_, d_2060_v48_])
                d_2122_v88_: _dafny.Array
                nw345_ = _dafny.Array(None, 15)
                nw345_[int(0)] = d_2119_v85_
                nw345_[int(1)] = d_2119_v85_
                nw345_[int(2)] = d_2119_v85_
                nw345_[int(3)] = _dafny.MultiSet(d_2121_v87_)
                nw345_[int(4)] = default__.fm46((d_2114_v80_)[default__.safeIndex(524, (d_2114_v80_).length(0))], (self).f10, (d_2114_v80_)[default__.safeIndex(524, (d_2114_v80_).length(0))], globalState)
                nw345_[int(5)] = d_2119_v85_
                nw345_[int(6)] = d_2119_v85_
                nw345_[int(7)] = d_2119_v85_
                nw345_[int(8)] = d_2119_v85_
                nw345_[int(9)] = _dafny.MultiSet([d_2060_v48_])
                nw345_[int(10)] = d_2119_v85_
                nw345_[int(11)] = d_2119_v85_
                nw345_[int(12)] = _dafny.MultiSet([d_2060_v48_, d_2060_v48_])
                nw345_[int(13)] = d_2119_v85_
                nw345_[int(14)] = _dafny.MultiSet(d_2121_v87_)
                d_2122_v88_ = nw345_
                d_2123_v89_: _dafny.Map
                d_2123_v89_ = _dafny.Map({(self).f10: d_2122_v88_})
                d_2124_v90_: D20
                d_2124_v90_ = D20_DC51(d_2122_v88_)
                d_2125_v91_: _dafny.Array
                nw346_ = _dafny.Array(None, 14)
                nw346_[int(0)] = ((d_2123_v89_)[(self).f10] if ((self).f10) in (d_2123_v89_) else d_2122_v88_)
                nw346_[int(1)] = (d_2124_v90_).cf97
                nw346_[int(2)] = d_2122_v88_
                nw346_[int(3)] = ((d_2123_v89_)[True] if (True) in (d_2123_v89_) else d_2122_v88_)
                nw346_[int(4)] = d_2122_v88_
                nw346_[int(5)] = d_2122_v88_
                nw346_[int(6)] = d_2122_v88_
                nw346_[int(7)] = d_2122_v88_
                nw346_[int(8)] = d_2122_v88_
                nw346_[int(9)] = d_2122_v88_
                nw346_[int(10)] = d_2122_v88_
                nw346_[int(11)] = d_2122_v88_
                nw346_[int(12)] = d_2122_v88_
                nw346_[int(13)] = (d_2122_v88_ if (d_2114_v80_)[default__.safeIndex(524, (d_2114_v80_).length(0))] else d_2122_v88_)
                d_2125_v91_ = nw346_
                index365_ = default__.safeIndex(45, (d_2125_v91_).length(0))
                (d_2125_v91_)[index365_] = d_2122_v88_
                d_2126_v92_: _dafny.Seq
                d_2126_v92_ = _dafny.SeqWithoutIsStrInference([d_2122_v88_])
                index366_ = default__.safeIndex(45, (d_2125_v91_).length(0))
                (d_2125_v91_)[index366_] = (d_2126_v92_)[default__.safeIndex((84) - (len(d_2061_v49_)), len(d_2126_v92_))]
            elif True:
                d_2127_v93_: _dafny.Map
                d_2127_v93_ = _dafny.Map({(d_2114_v80_)[default__.safeIndex(524, (d_2114_v80_).length(0))]: True})
                d_2127_v93_ = ((self).f9) | ((D6_DC16((self).f9)).cf28)
                d_2128_v94_: _dafny.Seq
                d_2128_v94_ = _dafny.SeqWithoutIsStrInference([d_2060_v48_])
                d_2129_v95_: _dafny.Array
                def lambda109_(d_2130_v48_):
                    def lambda110_(d_2131_i11_):
                        return (d_2131_i11_) - (d_2130_v48_)

                    return lambda110_

                init64_ = lambda109_(d_2060_v48_)
                nw347_ = _dafny.Array(None, 22)
                for i0_64_ in range(nw347_.length(0)):
                    nw347_[i0_64_] = init64_(i0_64_)
                d_2129_v95_ = nw347_
                d_2132_v96_: _dafny.Map
                d_2132_v96_ = _dafny.Map({d_2060_v48_: d_2060_v48_})
                d_2133_v97_: _dafny.Map
                d_2133_v97_ = _dafny.Map({d_2129_v95_: ((d_2132_v96_)[d_2060_v48_] if (d_2060_v48_) in (d_2132_v96_) else d_2060_v48_)})
                d_2134_v98_: _dafny.Map
                d_2134_v98_ = _dafny.Map({D4_DC12(True, d_2060_v48_, d_2115_v81_, d_2060_v48_, d_2060_v48_): d_2060_v48_})
                d_2135_v99_: D4
                d_2135_v99_ = D4_DC12((self).f10, d_2060_v48_, d_2115_v81_, d_2060_v48_, d_2060_v48_)
                d_2136_v100_: _dafny.Seq
                d_2136_v100_ = _dafny.SeqWithoutIsStrInference([d_2128_v94_, d_2128_v94_])
                d_2137_v101_: _dafny.Array
                nw348_ = _dafny.Array(None, 6)
                nw348_[int(0)] = len(_dafny.Set({default__.fm17((d_2114_v80_)[default__.safeIndex(524, (d_2114_v80_).length(0))], globalState)}))
                nw348_[int(1)] = (d_2128_v94_)[default__.safeIndex(d_2060_v48_, len(d_2128_v94_))]
                nw348_[int(2)] = (len(d_2133_v97_)) * (((d_2134_v98_)[d_2135_v99_] if (d_2135_v99_) in (d_2134_v98_) else d_2060_v48_))
                nw348_[int(3)] = d_2060_v48_
                nw348_[int(4)] = -343
                nw348_[int(5)] = len((_dafny.SeqWithoutIsStrInference([d_2128_v94_ for d_2138_i12_ in range(default__.abs(327))])) + (d_2136_v100_))
                d_2137_v101_ = nw348_
                index367_ = default__.safeIndex(571, (d_2129_v95_).length(0))
                (d_2129_v95_)[index367_] = len((d_2116_v82_) + (d_2116_v82_))
                d_2139_v102_: _dafny.MultiSet
                d_2139_v102_ = _dafny.MultiSet([(d_2114_v80_)[default__.safeIndex(524, (d_2114_v80_).length(0))]])
                d_2140_v103_: _dafny.MultiSet
                d_2140_v103_ = _dafny.MultiSet([d_2060_v48_, len((d_2116_v82_).set(default__.safeIndex((0) - (d_2060_v48_), len(d_2116_v82_)), d_2115_v81_)), d_2060_v48_])
                d_2141_v104_: _dafny.Seq
                d_2141_v104_ = _dafny.SeqWithoutIsStrInference([d_2140_v103_, d_2140_v103_])
                d_2142_v105_: _dafny.Seq
                d_2142_v105_ = _dafny.SeqWithoutIsStrInference([d_2140_v103_, d_2140_v103_, (d_2141_v104_)[default__.safeIndex(len(d_2116_v82_), len(d_2141_v104_))], default__.fm46((d_2114_v80_)[default__.safeIndex(524, (d_2114_v80_).length(0))], True, (d_2114_v80_)[default__.safeIndex(524, (d_2114_v80_).length(0))], globalState)])
                index368_ = default__.safeIndex(524, (d_2114_v80_).length(0))
                index369_ = default__.safeIndex(571, (d_2129_v95_).length(0))
                rhs343_ = (d_2116_v82_) + ((d_2116_v82_) + (d_2116_v82_))
                rhs344_ = True
                rhs345_ = (((d_2132_v96_)[799] if (799) in (d_2132_v96_) else (0) - (len((d_2128_v94_).set(default__.safeIndex(232, len(d_2128_v94_)), (0) - ((d_2139_v102_).cardinality))))) if (self).f10 else default__.safeDivisionInt(d_2060_v48_, (self).fm3(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fscb"))), d_2060_v48_, d_2115_v81_, len(d_2142_v105_), globalState)))
                rhs346_ = p0
                rhs347_ = (d_2060_v48_ if (self).f10 else 81)
                lhs248_ = d_2114_v80_
                lhs249_ = default__.safeIndex(524, (d_2114_v80_).length(0))
                lhs250_ = d_2129_v95_
                lhs251_ = default__.safeIndex(571, (d_2129_v95_).length(0))
                d_2116_v82_ = rhs343_
                lhs248_[lhs249_] = rhs344_
                d_2060_v48_ = rhs345_
                d_2115_v81_ = rhs346_
                lhs250_[lhs251_] = rhs347_
                d_2143_v106_: D5
                d_2143_v106_ = D5_DC13(d_2114_v80_)
                d_2144_v107_: _dafny.Set
                d_2144_v107_ = _dafny.Set({d_2143_v106_, d_2143_v106_, d_2143_v106_})
                d_2145_v108_: _dafny.Map
                d_2145_v108_ = _dafny.Map({d_2115_v81_: _dafny.Set({d_2143_v106_})})
                d_2146_v109_: str
                d_2146_v109_ = d_2115_v81_
                d_2144_v107_ = ((d_2145_v108_)[d_2146_v109_] if (d_2146_v109_) in (d_2145_v108_) else d_2144_v107_)
                d_2060_v48_ = ((d_2129_v95_)[default__.safeIndex(571, (d_2129_v95_).length(0))]) * (d_2060_v48_)
                d_2147_v110_: _dafny.Map
                d_2147_v110_ = _dafny.Map({d_2060_v48_: d_2137_v101_})
                d_2137_v101_ = ((d_2147_v110_)[default__.safeDivisionInt(570, d_2060_v48_)] if (default__.safeDivisionInt(570, d_2060_v48_)) in (d_2147_v110_) else d_2129_v95_)
            d_2148_v111_: _dafny.Map
            d_2148_v111_ = _dafny.Map({(d_2114_v80_)[default__.safeIndex(524, (d_2114_v80_).length(0))]: d_2060_v48_})
            d_2148_v111_ = (d_2148_v111_).set((self).f10, d_2060_v48_)
            index370_ = default__.safeIndex(524, (d_2114_v80_).length(0))
            (d_2114_v80_)[index370_] = (self).f10
        d_2149_v112_: int
        d_2150_v113_: _dafny.Seq
        out67_: int
        out68_: _dafny.Seq
        out67_, out68_ = (self).m2(globalState)
        d_2149_v112_ = out67_
        d_2150_v113_ = out68_
        d_2151_v114_: str
        d_2151_v114_ = _dafny.CodePoint('j')
        d_2151_v114_ = d_2151_v114_
        d_2152_v115_: _dafny.Array
        nw349_ = _dafny.Array(int(0), 17)
        d_2152_v115_ = nw349_
        d_2153_v116_: _dafny.Map
        d_2153_v116_ = _dafny.Map({d_2152_v115_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gwjllvov"))})
        r0 = d_2153_v116_
        r1 = ((self).f10) or (True)
        return r0, r1

    def m2(self, globalState):
        r0: int = int(0)
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        d_2154_v0_: int
        d_2154_v0_ = -474
        d_2155_v1_: _dafny.MultiSet
        d_2155_v1_ = _dafny.MultiSet([d_2154_v0_])
        d_2156_v2_: _dafny.MultiSet
        d_2156_v2_ = _dafny.MultiSet([(d_2155_v1_).cardinality])
        (globalState).f8 = (d_2156_v2_).ispropersubset((_dafny.MultiSet([d_2154_v0_])).intersection(d_2155_v1_))
        if (self).f10:
            d_2157_v3_: C5
            nw350_ = C5()
            nw350_.ctor__(((self).f9).set(False, (self).f10), not((self).f10))
            d_2157_v3_ = nw350_
            d_2158_v4_: _dafny.Seq
            d_2158_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))
            d_2159_v5_: T0
            nw351_ = C1()
            nw351_.ctor__((self).f10, d_2158_v4_)
            d_2159_v5_ = nw351_
            d_2160_v6_: D21
            d_2160_v6_ = D21_DC53(d_2159_v5_)
            d_2161_v7_: _dafny.Array
            nw352_ = _dafny.Array(None, 12)
            nw352_[int(0)] = d_2159_v5_
            nw352_[int(1)] = d_2159_v5_
            nw352_[int(2)] = d_2159_v5_
            nw352_[int(3)] = d_2159_v5_
            nw352_[int(4)] = (d_2160_v6_).cf99
            nw352_[int(5)] = d_2159_v5_
            nw352_[int(6)] = d_2159_v5_
            nw352_[int(7)] = d_2159_v5_
            nw352_[int(8)] = d_2159_v5_
            nw352_[int(9)] = d_2159_v5_
            nw352_[int(10)] = d_2159_v5_
            nw352_[int(11)] = d_2159_v5_
            d_2161_v7_ = nw352_
            index371_ = default__.safeIndex(962, (d_2161_v7_).length(0))
            (d_2161_v7_)[index371_] = d_2159_v5_
            d_2162_v8_: str
            d_2162_v8_ = _dafny.CodePoint('m')
            d_2163_v9_: D6
            d_2163_v9_ = D6_DC19(True, d_2162_v8_)
            index372_ = default__.safeIndex(962, (d_2161_v7_).length(0))
            rhs348_ = (d_2159_v5_ if (d_2154_v0_) >= (d_2154_v0_) else d_2159_v5_)
            rhs349_ = (not((self).f10)) == (False)
            rhs350_ = d_2159_v5_
            rhs351_ = d_2163_v9_
            lhs252_ = globalState
            lhs253_ = d_2161_v7_
            lhs254_ = default__.safeIndex(962, (d_2161_v7_).length(0))
            d_2159_v5_ = rhs348_
            lhs252_.f8 = rhs349_
            lhs253_[lhs254_] = rhs350_
            d_2163_v9_ = rhs351_
            d_2164_v10_: _dafny.Array
            def lambda111_(d_2165_v0_):
                def lambda112_(d_2166_i0_):
                    return (d_2166_i0_) * (d_2165_v0_)

                return lambda112_

            init65_ = lambda111_(d_2154_v0_)
            nw353_ = _dafny.Array(None, 22)
            for i0_65_ in range(nw353_.length(0)):
                nw353_[i0_65_] = init65_(i0_65_)
            d_2164_v10_ = nw353_
            d_2167_v11_: _dafny.MultiSet
            d_2167_v11_ = _dafny.MultiSet([(d_2157_v3_).fm18((self).f11, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "boc")), globalState), True])
            d_2168_v12_: _dafny.Map
            d_2168_v12_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ejnrlofe"))): (d_2167_v11_).cardinality})
            index373_ = default__.safeIndex(347, (d_2164_v10_).length(0))
            (d_2164_v10_)[index373_] = default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([(self).f10, (self).f10, (self).f10])), len(d_2168_v12_))
            d_2169_v13_: T1
            nw354_ = C4()
            nw354_.ctor__((self).f9, (self).f10)
            d_2169_v13_ = nw354_
            d_2170_v14_: _dafny.Array
            nw355_ = _dafny.Array(None, 4)
            nw355_[int(0)] = d_2169_v13_
            nw355_[int(1)] = (D22_DC56(d_2169_v13_)).cf104
            nw355_[int(2)] = (d_2169_v13_ if (self).f10 else d_2169_v13_)
            nw355_[int(3)] = d_2169_v13_
            d_2170_v14_ = nw355_
            index374_ = default__.safeIndex(884, (d_2170_v14_).length(0))
            (d_2170_v14_)[index374_] = d_2169_v13_
            d_2171_v15_: _dafny.Seq
            d_2171_v15_ = _dafny.SeqWithoutIsStrInference([(self).f10])
            index375_ = default__.safeIndex(347, (d_2164_v10_).length(0))
            index376_ = default__.safeIndex(884, (d_2170_v14_).length(0))
            rhs352_ = ((d_2157_v3_).fm1((self).f9, globalState)) * (122)
            rhs353_ = d_2154_v0_
            rhs354_ = (self).fm1(((d_2169_v13_).f9).set((d_2169_v13_).f10, (d_2169_v13_).f10), globalState)
            rhs355_ = (604) == (len(d_2171_v15_))
            rhs356_ = d_2169_v13_
            lhs255_ = d_2164_v10_
            lhs256_ = default__.safeIndex(347, (d_2164_v10_).length(0))
            lhs257_ = globalState
            lhs258_ = d_2170_v14_
            lhs259_ = default__.safeIndex(884, (d_2170_v14_).length(0))
            lhs255_[lhs256_] = rhs352_
            d_2154_v0_ = rhs353_
            r0 = rhs354_
            lhs257_.f8 = rhs355_
            lhs258_[lhs259_] = rhs356_
            (globalState).f8 = False
            d_2158_v4_ = (d_2158_v4_).set(default__.safeIndex((d_2164_v10_)[default__.safeIndex(347, (d_2164_v10_).length(0))], len(d_2158_v4_)), d_2162_v8_)
        elif True:
            d_2172_v16_: _dafny.Array
            def lambda113_(d_2173_v0_):
                def lambda114_(d_2174_i1_):
                    return (d_2174_i1_) * (d_2173_v0_)

                return lambda114_

            init66_ = lambda113_(d_2154_v0_)
            nw356_ = _dafny.Array(None, 19)
            for i0_66_ in range(nw356_.length(0)):
                nw356_[i0_66_] = init66_(i0_66_)
            d_2172_v16_ = nw356_
            index377_ = default__.safeIndex(237, (d_2172_v16_).length(0))
            (d_2172_v16_)[index377_] = d_2154_v0_
            index378_ = default__.safeIndex(237, (d_2172_v16_).length(0))
            (d_2172_v16_)[index378_] = (0) - (d_2154_v0_)
            d_2175_v17_: _dafny.Seq
            d_2175_v17_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f"))
            d_2176_v18_: D14
            d_2176_v18_ = D14_DC36((self).f10, d_2175_v17_, d_2175_v17_)
            d_2177_v19_: _dafny.Seq
            d_2177_v19_ = _dafny.SeqWithoutIsStrInference([d_2176_v18_])
            d_2178_v20_: str
            d_2178_v20_ = _dafny.CodePoint('e')
            d_2179_v21_: _dafny.Array
            nw357_ = _dafny.Array(None, 17)
            nw357_[int(0)] = _dafny.SeqWithoutIsStrInference([d_2176_v18_, d_2176_v18_, d_2176_v18_])
            nw357_[int(1)] = (default__.fm56((self).f10, globalState)) + (default__.fm56(True, globalState))
            nw357_[int(2)] = d_2177_v19_
            nw357_[int(3)] = d_2177_v19_
            nw357_[int(4)] = _dafny.SeqWithoutIsStrInference([d_2176_v18_, D14_DC36((self).f10, d_2175_v17_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_2180_i2_ in range(default__.abs(629))]))])
            nw357_[int(5)] = _dafny.SeqWithoutIsStrInference([d_2176_v18_])
            nw357_[int(6)] = (_dafny.SeqWithoutIsStrInference([d_2176_v18_])) + (d_2177_v19_)
            nw357_[int(7)] = d_2177_v19_
            nw357_[int(8)] = d_2177_v19_
            nw357_[int(9)] = d_2177_v19_
            nw357_[int(10)] = d_2177_v19_
            nw357_[int(11)] = d_2177_v19_
            nw357_[int(12)] = d_2177_v19_
            nw357_[int(13)] = default__.fm56((self).f10, globalState)
            nw357_[int(14)] = d_2177_v19_
            nw357_[int(15)] = d_2177_v19_
            nw357_[int(16)] = _dafny.SeqWithoutIsStrInference([D14_DC36((self).f10, (d_2175_v17_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lmmm"))), len(d_2175_v17_)), d_2178_v20_), d_2175_v17_)])
            d_2179_v21_ = nw357_
            pat_let_tv50_ = d_2179_v21_
            index379_ = default__.safeIndex(237, (d_2172_v16_).length(0))
            def iife180_(_pat_let52_0):
                def iife181_(d_2181_dt__update__tmp_h0_):
                    def iife182_(_pat_let53_0):
                        def iife183_(d_2182_dt__update_hcf105_h0_):
                            return D23_DC58(d_2182_dt__update_hcf105_h0_)
                        return iife183_(_pat_let53_0)
                    return iife182_(pat_let_tv50_)
                return iife181_(_pat_let52_0)
            rhs357_ = (d_2172_v16_)[default__.safeIndex(237, (d_2172_v16_).length(0))]
            rhs358_ = ((iife180_(D23_DC58(d_2179_v21_))).cf105 if (self).f10 else d_2179_v21_)
            lhs260_ = d_2172_v16_
            lhs261_ = default__.safeIndex(237, (d_2172_v16_).length(0))
            lhs260_[lhs261_] = rhs357_
            d_2179_v21_ = rhs358_
            d_2183_v22_: _dafny.Array
            nw358_ = _dafny.Array(D15.default()(), 6)
            d_2183_v22_ = nw358_
            d_2184_v23_: D19
            d_2184_v23_ = D19_DC49(False)
            d_2185_v24_: _dafny.Seq
            d_2185_v24_ = _dafny.SeqWithoutIsStrInference([(self).f10, (d_2184_v23_).cf93, (self).f10, (self).f10, (self).f10])
            d_2186_v25_: D15
            d_2186_v25_ = D15_DC39((self).f10, (d_2185_v24_)[default__.safeIndex((d_2172_v16_)[default__.safeIndex(237, (d_2172_v16_).length(0))], len(d_2185_v24_))])
            index380_ = default__.safeIndex(13, (d_2183_v22_).length(0))
            (d_2183_v22_)[index380_] = d_2186_v25_
            d_2187_v26_: _dafny.MultiSet
            d_2187_v26_ = _dafny.MultiSet([(self).f10, (self).f10, (self).f10, (self).f10])
            pat_let_tv51_ = d_2187_v26_
            pat_let_tv52_ = d_2185_v24_
            index381_ = default__.safeIndex(13, (d_2183_v22_).length(0))
            index382_ = default__.safeIndex(237, (d_2172_v16_).length(0))
            index383_ = default__.safeIndex(237, (d_2172_v16_).length(0))
            def iife184_(_pat_let54_0):
                def iife185_(d_2188_dt__update__tmp_h1_):
                    def iife186_(_pat_let55_0):
                        def iife187_(d_2189_dt__update_hcf74_h0_):
                            return D15_DC39(d_2189_dt__update_hcf74_h0_, (d_2188_dt__update__tmp_h1_).cf75)
                        return iife187_(_pat_let55_0)
                    return iife186_((pat_let_tv51_).isdisjoint(_dafny.MultiSet(pat_let_tv52_)))
                return iife185_(_pat_let54_0)
            rhs359_ = iife184_(d_2186_v25_)
            rhs360_ = (self).f10
            rhs361_ = len((self).f9)
            rhs362_ = d_2154_v0_
            lhs262_ = d_2183_v22_
            lhs263_ = default__.safeIndex(13, (d_2183_v22_).length(0))
            lhs264_ = globalState
            lhs265_ = d_2172_v16_
            lhs266_ = default__.safeIndex(237, (d_2172_v16_).length(0))
            lhs267_ = d_2172_v16_
            lhs268_ = default__.safeIndex(237, (d_2172_v16_).length(0))
            lhs262_[lhs263_] = rhs359_
            lhs264_.f8 = rhs360_
            lhs265_[lhs266_] = rhs361_
            lhs267_[lhs268_] = rhs362_
            d_2190_v27_: C5
            nw359_ = C5()
            nw359_.ctor__((self).f9, (self).f10)
            d_2190_v27_ = nw359_
            d_2191_v28_: _dafny.Array
            nw360_ = _dafny.Array(False, 18)
            d_2191_v28_ = nw360_
            index384_ = default__.safeIndex(533, (d_2191_v28_).length(0))
            (d_2191_v28_)[index384_] = (self).f10
            index385_ = default__.safeIndex(533, (d_2191_v28_).length(0))
            (d_2191_v28_)[index385_] = (self).f10
        d_2192_v29_: _dafny.Seq
        d_2192_v29_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "glcese"))
        if not(not(not(((d_2192_v29_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nhghrw")))) <= (d_2192_v29_)))):
            d_2193_v30_: _dafny.Array
            nw361_ = _dafny.Array(False, 26)
            d_2193_v30_ = nw361_
            d_2194_v31_: _dafny.MultiSet
            d_2194_v31_ = _dafny.MultiSet([(self).f10])
            index386_ = default__.safeIndex(412, (d_2193_v30_).length(0))
            (d_2193_v30_)[index386_] = (d_2194_v31_).ispropersubset(d_2194_v31_)
            index387_ = default__.safeIndex(412, (d_2193_v30_).length(0))
            rhs363_ = (self).f10
            rhs364_ = d_2154_v0_
            rhs365_ = default__.safeDivisionInt(d_2154_v0_, len((self).f11))
            lhs269_ = d_2193_v30_
            lhs270_ = default__.safeIndex(412, (d_2193_v30_).length(0))
            lhs269_[lhs270_] = rhs363_
            r0 = rhs364_
            r0 = rhs365_
            d_2195_v32_: _dafny.MultiSet
            d_2195_v32_ = _dafny.MultiSet([_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kiarxiha")))})])
            (globalState).f8 = not((d_2195_v32_).issubset(d_2195_v32_))
            r0 = d_2154_v0_
            if (d_2193_v30_)[default__.safeIndex(412, (d_2193_v30_).length(0))]:
                d_2196_v33_: _dafny.Array
                nw362_ = _dafny.Array(int(0), 11)
                d_2196_v33_ = nw362_
                d_2197_v34_: _dafny.Seq
                d_2197_v34_ = _dafny.SeqWithoutIsStrInference([d_2154_v0_, d_2154_v0_, d_2154_v0_, 93])
                index388_ = default__.safeIndex(696, (d_2196_v33_).length(0))
                (d_2196_v33_)[index388_] = len(((d_2197_v34_) + (d_2197_v34_)).set(default__.safeIndex((d_2197_v34_)[default__.safeIndex(d_2154_v0_, len(d_2197_v34_))], len((d_2197_v34_) + (d_2197_v34_))), d_2154_v0_))
                index389_ = default__.safeIndex(696, (d_2196_v33_).length(0))
                (d_2196_v33_)[index389_] = (d_2197_v34_)[default__.safeIndex(d_2154_v0_, len(d_2197_v34_))]
                index390_ = default__.safeIndex(696, (d_2196_v33_).length(0))
                (d_2196_v33_)[index390_] = len(((d_2197_v34_) + (d_2197_v34_)) + (d_2197_v34_))
                d_2198_v35_: _dafny.Set
                d_2198_v35_ = _dafny.Set({(d_2196_v33_)[default__.safeIndex(696, (d_2196_v33_).length(0))], 815})
                d_2199_v36_: _dafny.Map
                d_2199_v36_ = _dafny.Map({len(d_2198_v35_): (d_2196_v33_)[default__.safeIndex(696, (d_2196_v33_).length(0))]})
                d_2200_v37_: _dafny.Seq
                d_2200_v37_ = _dafny.SeqWithoutIsStrInference([(d_2193_v30_)[default__.safeIndex(412, (d_2193_v30_).length(0))]])
                d_2201_v38_: str
                d_2201_v38_ = _dafny.CodePoint('v')
                d_2202_v40_: _dafny.Seq
                d_2202_v40_ = _dafny.SeqWithoutIsStrInference([d_2198_v35_])
                def iife188_():
                    coll76_ = _dafny.Map()
                    compr_76_: int
                    for compr_76_ in ((d_2202_v40_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ykixtfsa"))), len(d_2202_v40_))]).Elements:
                        d_2203_v39_: int = compr_76_
                        if (d_2203_v39_) in ((d_2202_v40_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ykixtfsa"))), len(d_2202_v40_))]):
                            coll76_[(d_2203_v39_) + ((d_2196_v33_)[default__.safeIndex(696, (d_2196_v33_).length(0))])] = d_2154_v0_
                    return _dafny.Map(coll76_)
                (globalState).f8 = ((d_2199_v36_).set((0) - ((default__.fm15((self).f10, not((d_2200_v37_)[default__.safeIndex(-274, len(d_2200_v37_))]), len(d_2198_v35_), d_2201_v38_, globalState)).cardinality), d_2154_v0_)) != (iife188_()
                )
                d_2204_v41_: C4
                nw363_ = C4()
                nw363_.ctor__((self).f9, (self).f10)
                d_2204_v41_ = nw363_
                d_2205_v42_: _dafny.Seq
                d_2205_v42_ = _dafny.SeqWithoutIsStrInference([d_2204_v41_])
                d_2206_v43_: _dafny.Map
                d_2206_v43_ = _dafny.Map({((d_2199_v36_)[d_2154_v0_] if (d_2154_v0_) in (d_2199_v36_) else (d_2196_v33_)[default__.safeIndex(696, (d_2196_v33_).length(0))]): (d_2205_v42_)[default__.safeIndex(d_2154_v0_, len(d_2205_v42_))]})
                d_2206_v43_ = d_2206_v43_
                d_2207_v44_: C2
                nw364_ = C2()
                nw364_.ctor__((d_2193_v30_)[default__.safeIndex(412, (d_2193_v30_).length(0))], d_2193_v30_, (self).f9, (self).fm0(d_2200_v37_, (d_2193_v30_)[default__.safeIndex(412, (d_2193_v30_).length(0))], d_2154_v0_, (d_2204_v41_).fm0(d_2200_v37_, (d_2193_v30_)[default__.safeIndex(412, (d_2193_v30_).length(0))], 939, False, globalState), globalState))
                d_2207_v44_ = nw364_
                d_2207_v44_ = d_2207_v44_
            elif True:
                d_2208_v45_: _dafny.Map
                d_2208_v45_ = _dafny.Map({not(((self).f11).ispropersubset((self).f11)): (237) == (d_2154_v0_)})
                d_2208_v45_ = d_2208_v45_
                d_2154_v0_ = d_2154_v0_
                d_2209_v46_: _dafny.Map
                d_2209_v46_ = _dafny.Map({(self).f10: d_2192_v29_})
                d_2210_v47_: C1
                nw365_ = C1()
                nw365_.ctor__((d_2193_v30_)[default__.safeIndex(412, (d_2193_v30_).length(0))], (((d_2209_v46_)[(d_2193_v30_)[default__.safeIndex(412, (d_2193_v30_).length(0))]] if ((d_2193_v30_)[default__.safeIndex(412, (d_2193_v30_).length(0))]) in (d_2209_v46_) else d_2192_v29_)) + (d_2192_v29_))
                d_2210_v47_ = nw365_
                d_2211_v48_: _dafny.Array
                nw366_ = _dafny.Array(int(0), 11)
                d_2211_v48_ = nw366_
                index391_ = default__.safeIndex(767, (d_2211_v48_).length(0))
                (d_2211_v48_)[index391_] = d_2154_v0_
                index392_ = default__.safeIndex(767, (d_2211_v48_).length(0))
                (d_2211_v48_)[index392_] = d_2154_v0_
                d_2212_v49_: C5
                nw367_ = C5()
                nw367_.ctor__((d_2208_v45_) | (_dafny.Map({True: (d_2193_v30_)[default__.safeIndex(412, (d_2193_v30_).length(0))]})), ((d_2211_v48_)[default__.safeIndex(767, (d_2211_v48_).length(0))]) == ((d_2211_v48_)[default__.safeIndex(767, (d_2211_v48_).length(0))]))
                d_2212_v49_ = nw367_
            d_2213_v50_: _dafny.Array
            nw368_ = _dafny.Array(None, 10)
            nw368_[int(0)] = default__.fm33((d_2193_v30_)[default__.safeIndex(412, (d_2193_v30_).length(0))], (self).f10, d_2154_v0_, (self).f10, globalState)
            nw368_[int(1)] = (d_2154_v0_) * (d_2154_v0_)
            nw368_[int(2)] = d_2154_v0_
            nw368_[int(3)] = d_2154_v0_
            nw368_[int(4)] = d_2154_v0_
            nw368_[int(5)] = (0) - (d_2154_v0_)
            nw368_[int(6)] = default__.safeModuloInt(d_2154_v0_, d_2154_v0_)
            nw368_[int(7)] = d_2154_v0_
            nw368_[int(8)] = d_2154_v0_
            nw368_[int(9)] = default__.fm33((d_2193_v30_)[default__.safeIndex(412, (d_2193_v30_).length(0))], (self).f10, d_2154_v0_, True, globalState)
            d_2213_v50_ = nw368_
            d_2213_v50_ = d_2213_v50_
        elif True:
            d_2214_v51_: _dafny.Map
            d_2214_v51_ = _dafny.Map({((self).f10) == ((self).f10): not((self).f10)})
            d_2214_v51_ = (d_2214_v51_).set((self).f10, (self).f10)
            d_2215_v52_: _dafny.MultiSet
            d_2216_v53_: bool
            d_2217_v54_: int
            d_2218_v55_: _dafny.Seq
            out69_: _dafny.MultiSet
            out70_: bool
            out71_: int
            out72_: _dafny.Seq
            out69_, out70_, out71_, out72_ = (self).m3(272, globalState)
            d_2215_v52_ = out69_
            d_2216_v53_ = out70_
            d_2217_v54_ = out71_
            d_2218_v55_ = out72_
            (globalState).f8 = (self).f10
            d_2219_v56_: _dafny.Map
            d_2219_v56_ = _dafny.Map({_dafny.Map({(0) - (d_2217_v54_): (self).f10}): d_2154_v0_})
            d_2220_v57_: _dafny.Seq
            d_2220_v57_ = _dafny.SeqWithoutIsStrInference([len(d_2219_v56_)])
            d_2221_v58_: _dafny.Map
            d_2221_v58_ = _dafny.Map({d_2216_v53_: D1_DC2(d_2220_v57_)})
            d_2222_v59_: str
            d_2222_v59_ = _dafny.CodePoint('c')
            d_2218_v55_ = (((d_2192_v29_) + (d_2192_v29_)).set(default__.safeIndex(len(d_2221_v58_), len((d_2192_v29_) + (d_2192_v29_))), (D14_DC35(d_2154_v0_, d_2216_v53_, d_2222_v59_)).cf64)).set(default__.safeIndex(d_2217_v54_, len(((d_2192_v29_) + (d_2192_v29_)).set(default__.safeIndex(len(d_2221_v58_), len((d_2192_v29_) + (d_2192_v29_))), (D14_DC35(d_2154_v0_, d_2216_v53_, d_2222_v59_)).cf64))), _dafny.CodePoint('a'))
            d_2223_v60_: _dafny.Array
            def lambda115_(d_2224_i3_):
                return (self).f10

            init67_ = lambda115_
            nw369_ = _dafny.Array(None, 1)
            for i0_67_ in range(nw369_.length(0)):
                nw369_[i0_67_] = init67_(i0_67_)
            d_2223_v60_ = nw369_
            d_2225_v61_: _dafny.Seq
            d_2225_v61_ = _dafny.SeqWithoutIsStrInference([False, d_2216_v53_])
            d_2226_v62_: C2
            nw370_ = C2()
            nw370_.ctor__(False, d_2223_v60_, (self).f9, ((self).fm2(globalState)) or ((d_2225_v61_)[default__.safeIndex((0) - (d_2217_v54_), len(d_2225_v61_))]))
            d_2226_v62_ = nw370_
        d_2227_v63_: _dafny.Map
        d_2227_v63_ = _dafny.Map({d_2156_v2_: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_2228_i5_ in range(default__.abs(761))]))})
        hi7_ = len(d_2227_v63_)
        for d_2229_i4_ in range(d_2154_v0_, hi7_):
            (globalState).f8 = (d_2154_v0_) != (d_2229_i4_)
            r0 = (default__.safeModuloInt(len(d_2192_v29_), len(d_2192_v29_))) * (d_2229_i4_)
            r1 = (d_2192_v29_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cemlbkmst")))
            d_2230_v64_: str
            d_2230_v64_ = _dafny.CodePoint('b')
            d_2231_v65_: D1
            d_2231_v65_ = D1_DC3(d_2192_v29_)
            d_2232_v66_: _dafny.Array
            nw371_ = _dafny.Array(None, 19)
            nw371_[int(0)] = d_2192_v29_
            nw371_[int(1)] = d_2192_v29_
            nw371_[int(2)] = d_2192_v29_
            nw371_[int(3)] = d_2192_v29_
            nw371_[int(4)] = d_2192_v29_
            nw371_[int(5)] = ((d_2192_v29_).set(default__.safeIndex(d_2154_v0_, len(d_2192_v29_)), _dafny.CodePoint('q'))).set(default__.safeIndex(d_2229_i4_, len((d_2192_v29_).set(default__.safeIndex(d_2154_v0_, len(d_2192_v29_)), _dafny.CodePoint('q')))), d_2230_v64_)
            nw371_[int(6)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t')])
            nw371_[int(7)] = (d_2192_v29_) + (d_2192_v29_)
            nw371_[int(8)] = d_2192_v29_
            nw371_[int(9)] = (d_2192_v29_) + (d_2192_v29_)
            nw371_[int(10)] = default__.fm13(d_2154_v0_, globalState)
            nw371_[int(11)] = (d_2231_v65_).cf5
            nw371_[int(12)] = (d_2192_v29_) + (d_2192_v29_)
            nw371_[int(13)] = d_2192_v29_
            nw371_[int(14)] = (d_2192_v29_) + (_dafny.SeqWithoutIsStrInference([d_2230_v64_]))
            nw371_[int(15)] = _dafny.SeqWithoutIsStrInference([d_2230_v64_])
            nw371_[int(16)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_2233_i6_ in range(default__.abs(524))])
            nw371_[int(17)] = (d_2192_v29_) + (d_2192_v29_)
            nw371_[int(18)] = d_2192_v29_
            d_2232_v66_ = nw371_
            index393_ = default__.safeIndex(90, (d_2232_v66_).length(0))
            (d_2232_v66_)[index393_] = d_2192_v29_
            d_2234_v67_: T1
            nw372_ = C6()
            nw372_.ctor__(_dafny.Map({(self).f10: False}), not((self).f10))
            d_2234_v67_ = nw372_
            d_2235_v68_: _dafny.Set
            d_2235_v68_ = _dafny.Set({d_2229_i4_, d_2229_i4_})
            index394_ = default__.safeIndex(90, (d_2232_v66_).length(0))
            rhs366_ = _dafny.MultiSet([d_2154_v0_, len(_dafny.SeqWithoutIsStrInference([not((self).f10), (self).f10])), (len(_dafny.Map({d_2234_v67_: not((self).f10)})) if (self).f10 else d_2229_i4_), d_2229_i4_, len(d_2235_v68_)])
            rhs367_ = default__.safeDivisionInt(d_2229_i4_, (0) - (d_2229_i4_))
            rhs368_ = (d_2192_v29_).set(default__.safeIndex(d_2229_i4_, len(d_2192_v29_)), _dafny.CodePoint('x'))
            rhs369_ = (self).f10
            rhs370_ = d_2230_v64_
            lhs271_ = d_2232_v66_
            lhs272_ = default__.safeIndex(90, (d_2232_v66_).length(0))
            lhs273_ = globalState
            d_2156_v2_ = rhs366_
            d_2154_v0_ = rhs367_
            lhs271_[lhs272_] = rhs368_
            lhs273_.f8 = rhs369_
            d_2230_v64_ = rhs370_
        d_2236_v69_: _dafny.Seq
        d_2236_v69_ = _dafny.SeqWithoutIsStrInference([(self).f10])
        d_2237_v70_: _dafny.Array
        nw373_ = _dafny.Array(None, 13)
        nw373_[int(0)] = True
        nw373_[int(1)] = (self).f10
        nw373_[int(2)] = (self).f10
        nw373_[int(3)] = (self).f10
        nw373_[int(4)] = (self).f10
        nw373_[int(5)] = (self).f10
        nw373_[int(6)] = (self).fm0(d_2236_v69_, (self).f10, len(default__.fm25(d_2154_v0_, 435, globalState)), (self).f10, globalState)
        nw373_[int(7)] = (self).f10
        nw373_[int(8)] = (self).f10
        nw373_[int(9)] = (self).f10
        nw373_[int(10)] = (self).f10
        nw373_[int(11)] = (self).f10
        nw373_[int(12)] = (self).f10
        d_2237_v70_ = nw373_
        d_2238_v71_: _dafny.Seq
        d_2238_v71_ = _dafny.SeqWithoutIsStrInference([d_2237_v70_])
        d_2239_v72_: _dafny.Map
        d_2239_v72_ = _dafny.Map({(0) - ((0) - (d_2154_v0_)): (self).f10})
        d_2240_v73_: D5
        d_2240_v73_ = D5_DC13((d_2238_v71_)[default__.safeIndex(len(d_2239_v72_), len(d_2238_v71_))])
        source29_ = d_2240_v73_
        if source29_.is_DC14:
            d_2241___mcc_h0_ = source29_.cf25
            d_2242___mcc_h1_ = source29_.cf26
            d_2243_cf26_ = d_2242___mcc_h1_
            d_2244_cf25_ = d_2241___mcc_h0_
            d_2245_v74_: _dafny.Seq
            d_2245_v74_ = _dafny.SeqWithoutIsStrInference([(self).f11, (self).f11, (self).f11, (self).f11])
            d_2246_v75_: _dafny.Array
            nw374_ = _dafny.Array(None, 17)
            nw374_[int(0)] = default__.fm30(globalState)
            nw374_[int(1)] = (_dafny.Set({d_2243_cf26_, (self).fm0(d_2236_v69_, d_2243_cf26_, 497, (self).f10, globalState), d_2243_cf26_, d_2243_cf26_, d_2243_cf26_})).intersection(_dafny.Set({(self).f10, False, d_2243_cf26_}))
            nw374_[int(2)] = _dafny.Set({d_2243_cf26_, d_2243_cf26_, d_2243_cf26_})
            nw374_[int(3)] = ((d_2245_v74_)[default__.safeIndex(d_2244_cf25_, len(d_2245_v74_))]) - ((self).f11)
            nw374_[int(4)] = ((self).f11) | ((self).f11)
            nw374_[int(5)] = (self).f11
            nw374_[int(6)] = (self).f11
            nw374_[int(7)] = (self).f11
            nw374_[int(8)] = (self).f11
            nw374_[int(9)] = (self).f11
            nw374_[int(10)] = (self).f11
            nw374_[int(11)] = (self).f11
            nw374_[int(12)] = _dafny.Set({(self).f10})
            nw374_[int(13)] = (self).f11
            nw374_[int(14)] = ((self).f11) - ((self).f11)
            nw374_[int(15)] = (self).f11
            nw374_[int(16)] = (self).f11
            d_2246_v75_ = nw374_
            d_2246_v75_ = d_2246_v75_
            d_2247_v76_: C5
            nw375_ = C5()
            nw375_.ctor__((self).f9, d_2243_cf26_)
            d_2247_v76_ = nw375_
            d_2248_v77_: _dafny.Array
            nw376_ = _dafny.Array(int(0), 7)
            d_2248_v77_ = nw376_
            index395_ = default__.safeIndex(543, (d_2248_v77_).length(0))
            (d_2248_v77_)[index395_] = d_2154_v0_
            index396_ = default__.safeIndex(543, (d_2248_v77_).length(0))
            (d_2248_v77_)[index396_] = default__.safeModuloInt(len(d_2192_v29_), d_2244_cf25_)
            d_2244_cf25_ = (0) - ((d_2154_v0_) - (d_2154_v0_))
        elif source29_.is_DC13:
            d_2249___mcc_h2_ = source29_.cf24
            d_2250_cf24_ = d_2249___mcc_h2_
            d_2251_v78_: str
            d_2251_v78_ = _dafny.CodePoint('c')
            d_2252_v79_: _dafny.Map
            d_2252_v79_ = _dafny.Map({default__.fm36(globalState): d_2154_v0_})
            d_2253_v80_: _dafny.Set
            d_2253_v80_ = _dafny.Set({(self).fm3((0) - (d_2154_v0_), d_2154_v0_, d_2251_v78_, len(d_2252_v79_), globalState)})
            d_2254_v81_: _dafny.Seq
            d_2254_v81_ = _dafny.SeqWithoutIsStrInference([(_dafny.Set({d_2154_v0_, d_2154_v0_}) if (self).f10 else _dafny.Set({d_2154_v0_, d_2154_v0_})), _dafny.Set({803, d_2154_v0_}), d_2253_v80_])
            d_2254_v81_ = (d_2254_v81_) + (d_2254_v81_)
            d_2154_v0_ = d_2154_v0_
            if (self).fm2(globalState):
                r0 = d_2154_v0_
                d_2255_v82_: _dafny.Array
                nw377_ = _dafny.Array(D15.default()(), 5)
                d_2255_v82_ = nw377_
                d_2256_v83_: _dafny.Map
                d_2256_v83_ = _dafny.Map({(self).fm3(d_2154_v0_, d_2154_v0_, d_2251_v78_, len(_dafny.Map({d_2251_v78_: not((self).f10)})), globalState): d_2255_v82_})
                d_2257_v84_: _dafny.Array
                nw378_ = _dafny.Array(None, 6)
                nw378_[int(0)] = (d_2256_v83_) | (d_2256_v83_)
                nw378_[int(1)] = (d_2256_v83_) | (_dafny.Map({d_2154_v0_: d_2255_v82_}))
                nw378_[int(2)] = (_dafny.Map({d_2154_v0_: d_2255_v82_})) | (d_2256_v83_)
                nw378_[int(3)] = (d_2256_v83_) | (d_2256_v83_)
                nw378_[int(4)] = d_2256_v83_
                nw378_[int(5)] = d_2256_v83_
                d_2257_v84_ = nw378_
                index397_ = default__.safeIndex(576, (d_2257_v84_).length(0))
                (d_2257_v84_)[index397_] = d_2256_v83_
                d_2258_v85_: _dafny.Array
                nw379_ = _dafny.Array(int(0), 13)
                d_2258_v85_ = nw379_
                index398_ = default__.safeIndex(576, (d_2257_v84_).length(0))
                rhs371_ = (_dafny.Map({d_2154_v0_: d_2255_v82_}) if (self).f10 else d_2256_v83_)
                rhs372_ = d_2251_v78_
                rhs373_ = d_2258_v85_
                rhs374_ = d_2154_v0_
                lhs274_ = d_2257_v84_
                lhs275_ = default__.safeIndex(576, (d_2257_v84_).length(0))
                lhs274_[lhs275_] = rhs371_
                d_2251_v78_ = rhs372_
                d_2258_v85_ = rhs373_
                d_2154_v0_ = rhs374_
                index399_ = default__.safeIndex(690, (d_2250_cf24_).length(0))
                (d_2250_cf24_)[index399_] = False
                d_2259_v86_: _dafny.Map
                d_2259_v86_ = _dafny.Map({d_2258_v85_: True})
                d_2260_v87_: _dafny.Map
                d_2260_v87_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_2251_v78_ for d_2261_i7_ in range(default__.abs(427))]): ((d_2259_v86_)[d_2258_v85_] if (d_2258_v85_) in (d_2259_v86_) else (self).f10)})
                index400_ = default__.safeIndex(690, (d_2250_cf24_).length(0))
                rhs375_ = ((d_2260_v87_)[((D1_DC3(d_2192_v29_)).cf5).set(default__.safeIndex(d_2154_v0_, len((D1_DC3(d_2192_v29_)).cf5)), _dafny.CodePoint('d'))] if (((D1_DC3(d_2192_v29_)).cf5).set(default__.safeIndex(d_2154_v0_, len((D1_DC3(d_2192_v29_)).cf5)), _dafny.CodePoint('d'))) in (d_2260_v87_) else (self).f10)
                rhs376_ = not((not((self).f10)) or ((self).f10))
                rhs377_ = d_2154_v0_
                rhs378_ = (self).f10
                lhs276_ = d_2250_cf24_
                lhs277_ = default__.safeIndex(690, (d_2250_cf24_).length(0))
                lhs278_ = globalState
                lhs279_ = globalState
                lhs276_[lhs277_] = rhs375_
                lhs278_.f8 = rhs376_
                r0 = rhs377_
                lhs279_.f8 = rhs378_
                d_2262_v88_: C6
                nw380_ = C6()
                nw380_.ctor__((self).f9, (self).f10)
                d_2262_v88_ = nw380_
                index401_ = default__.safeIndex(384, (d_2258_v85_).length(0))
                (d_2258_v85_)[index401_] = (0) - (default__.safeDivisionInt(d_2154_v0_, d_2154_v0_))
                d_2263_v89_: _dafny.Map
                d_2263_v89_ = _dafny.Map({False: len(_dafny.Map({(self).f10: d_2192_v29_}))})
                index402_ = default__.safeIndex(384, (d_2258_v85_).length(0))
                (d_2258_v85_)[index402_] = ((d_2263_v89_)[(self).f10] if ((self).f10) in (d_2263_v89_) else d_2154_v0_)
            elif True:
                r1 = (d_2192_v29_ if (self).f10 else (d_2192_v29_) + (d_2192_v29_))
                d_2264_v90_: D22
                d_2264_v90_ = D22_DC57()
                d_2264_v90_ = d_2264_v90_
                r0 = (0) - (((d_2154_v0_ if (self).f10 else d_2154_v0_)) + (len(d_2192_v29_)))
                d_2265_v91_: C1
                nw381_ = C1()
                nw381_.ctor__((self).f10, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v")))
                d_2265_v91_ = nw381_
                d_2266_v92_: _dafny.Seq
                d_2266_v92_ = _dafny.SeqWithoutIsStrInference([d_2265_v91_, d_2265_v91_])
                d_2267_v93_: _dafny.Set
                d_2267_v93_ = _dafny.Set({((d_2266_v92_).set(default__.safeIndex(d_2154_v0_, len(d_2266_v92_)), d_2265_v91_)) + (d_2266_v92_), (d_2266_v92_) + ((d_2266_v92_).set(default__.safeIndex(len(d_2192_v29_), len(d_2266_v92_)), d_2265_v91_))})
                d_2267_v93_ = d_2267_v93_
                d_2154_v0_ = (0) - ((0) - (len((d_2254_v81_) + (d_2254_v81_))))
            (globalState).f8 = (self).f10
        elif True:
            d_2268___mcc_h3_ = source29_.cf27
            d_2269_cf27_ = d_2268___mcc_h3_
            d_2270_v94_: _dafny.Seq
            d_2270_v94_ = _dafny.SeqWithoutIsStrInference([(0) - (d_2154_v0_), d_2154_v0_])
            d_2270_v94_ = d_2270_v94_
            d_2271_v95_: C1
            nw382_ = C1()
            nw382_.ctor__((self).f10, d_2192_v29_)
            d_2271_v95_ = nw382_
            d_2271_v95_ = (d_2271_v95_ if (self).fm2(globalState) else d_2271_v95_)
            d_2272_v96_: str
            d_2272_v96_ = _dafny.CodePoint('i')
            d_2273_v97_: _dafny.Map
            d_2273_v97_ = _dafny.Map({default__.fm36(globalState): d_2272_v96_})
            (globalState).f8 = (len(d_2273_v97_)) != (d_2154_v0_)
            d_2154_v0_ = (d_2154_v0_) - ((33) - (d_2154_v0_))
        d_2274_v98_: C2
        nw383_ = C2()
        nw383_.ctor__((d_2236_v69_)[default__.safeIndex(d_2154_v0_, len(d_2236_v69_))], d_2237_v70_, (_dafny.Map({(self).f10: (self).f10})) | ((self).f9), (self).f10)
        d_2274_v98_ = nw383_
        r0 = (default__.fm46(False, (d_2156_v2_).isdisjoint(d_2156_v2_), d_2274_v98_.f14, globalState)).cardinality
        d_2275_v99_: str
        d_2275_v99_ = _dafny.CodePoint('g')
        r1 = (d_2192_v29_).set(default__.safeIndex(len((self).f11), len(d_2192_v29_)), d_2275_v99_)
        return r0, r1

    def m3(self, p0, globalState):
        r0: _dafny.MultiSet = _dafny.MultiSet({})
        r1: bool = False
        r2: int = int(0)
        r3: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        d_2276_v0_: str
        d_2276_v0_ = _dafny.CodePoint('x')
        d_2277_v1_: D6
        d_2277_v1_ = D6_DC18(d_2276_v0_, p0, p0)
        d_2278_v2_: D18
        d_2278_v2_ = D18_DC47((self).f10, d_2277_v1_)
        d_2279_v3_: _dafny.Map
        d_2279_v3_ = _dafny.Map({(0) - (default__.safeModuloInt(p0, 958)): d_2278_v2_})
        d_2280_v4_: _dafny.Map
        d_2280_v4_ = _dafny.Map({(self).f10: p0})
        d_2281_v5_: _dafny.MultiSet
        d_2281_v5_ = _dafny.MultiSet([(self).f10, (self).f10, (self).f10, (self).f10])
        d_2282_v6_: _dafny.MultiSet
        d_2282_v6_ = _dafny.MultiSet([(d_2281_v5_).cardinality])
        d_2283_v7_: _dafny.Map
        d_2283_v7_ = _dafny.Map({len((d_2280_v4_)): (d_2282_v6_).cardinality})
        rhs379_ = ((_dafny.Map({len(d_2283_v7_): d_2278_v2_})) | (_dafny.Map({p0: d_2278_v2_}))) | ((_dafny.Map({p0: d_2278_v2_})) | (d_2279_v3_))
        rhs380_ = p0
        d_2279_v3_ = rhs379_
        r2 = rhs380_
        d_2284_i0_: int
        d_2284_i0_ = 0
        with _dafny.label("9"):
            while (self).f10:
                with _dafny.c_label("9"):
                    if (d_2284_i0_) >= (100):
                        raise _dafny.Break("9")
                    d_2284_i0_ = (d_2284_i0_) + (1)
                    d_2285_v8_: _dafny.Array
                    def lambda116_(d_2286_i1_):
                        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dvupvsxfh"))

                    init68_ = lambda116_
                    nw384_ = _dafny.Array(None, 9)
                    for i0_68_ in range(nw384_.length(0)):
                        nw384_[i0_68_] = init68_(i0_68_)
                    d_2285_v8_ = nw384_
                    d_2287_v9_: _dafny.Map
                    d_2287_v9_ = _dafny.Map({p0: d_2285_v8_})
                    d_2287_v9_ = (d_2287_v9_).set((0) - (default__.safeModuloInt(p0, p0)), (d_2285_v8_ if (self).f10 else d_2285_v8_))
                    r2 = (-261) - (p0)
                    r1 = (self).f10
                    d_2288_v10_: C5
                    nw385_ = C5()
                    nw385_.ctor__((self).f9, False)
                    d_2288_v10_ = nw385_
                    pass
            pass
        r2 = default__.safeModuloInt((p0 if (self).f10 else p0), default__.safeDivisionInt(p0, p0))
        if (self).f10:
            r2 = (self).fm1((self).f9, globalState)
            d_2289_v11_: _dafny.Array
            nw386_ = _dafny.Array(int(0), 13)
            d_2289_v11_ = nw386_
            d_2290_v12_: _dafny.Array
            nw387_ = _dafny.Array(_dafny.Seq({}), 3)
            d_2290_v12_ = nw387_
            index403_ = default__.safeIndex(484, (d_2290_v12_).length(0))
            (d_2290_v12_)[index403_] = _dafny.SeqWithoutIsStrInference([(self).f10, (self).f10])
            d_2291_v13_: C4
            nw388_ = C4()
            nw388_.ctor__(((self).f9).set((self).f10, (self).f10), (self).f10)
            d_2291_v13_ = nw388_
            index404_ = default__.safeIndex(311, (d_2289_v11_).length(0))
            (d_2289_v11_)[index404_] = p0
            d_2292_v14_: _dafny.MultiSet
            d_2292_v14_ = _dafny.MultiSet([d_2276_v0_])
            d_2293_v15_: _dafny.Seq
            d_2293_v15_ = _dafny.SeqWithoutIsStrInference([(p0) >= ((0) - (p0)), (d_2292_v14_).isdisjoint(d_2292_v14_)])
            index405_ = default__.safeIndex(484, (d_2290_v12_).length(0))
            index406_ = default__.safeIndex(311, (d_2289_v11_).length(0))
            rhs381_ = d_2289_v11_
            rhs382_ = d_2293_v15_
            rhs383_ = d_2291_v13_
            rhs384_ = ((p0 if not(False) else p0)) * (((d_2282_v6_)[p0] if (p0) in (d_2282_v6_) else p0))
            lhs280_ = d_2290_v12_
            lhs281_ = default__.safeIndex(484, (d_2290_v12_).length(0))
            lhs282_ = d_2289_v11_
            lhs283_ = default__.safeIndex(311, (d_2289_v11_).length(0))
            d_2289_v11_ = rhs381_
            lhs280_[lhs281_] = rhs382_
            d_2291_v13_ = rhs383_
            lhs282_[lhs283_] = rhs384_
            d_2294_v16_: _dafny.Map
            d_2294_v16_ = _dafny.Map({(155) <= ((0) - ((d_2289_v11_)[default__.safeIndex(311, (d_2289_v11_).length(0))])): d_2276_v0_})
            d_2276_v0_ = ((d_2294_v16_)[False] if (False) in (d_2294_v16_) else d_2276_v0_)
            d_2295_v17_: _dafny.Map
            d_2295_v17_ = _dafny.Map({((self).f10 if (self).f10 else True): d_2281_v5_})
            d_2295_v17_ = (d_2295_v17_).set((self).f10, d_2281_v5_)
            d_2296_v18_: _dafny.Array
            def lambda117_(d_2297_i2_):
                return (self).f10

            init69_ = lambda117_
            nw389_ = _dafny.Array(None, 20)
            for i0_69_ in range(nw389_.length(0)):
                nw389_[i0_69_] = init69_(i0_69_)
            d_2296_v18_ = nw389_
            d_2296_v18_ = (d_2296_v18_ if True else d_2296_v18_)
        elif True:
            d_2298_v19_: D19
            d_2298_v19_ = D19_DC49(True)
            d_2299_v20_: D1
            d_2299_v20_ = D1_DC4((d_2298_v19_).cf93, (self).f10, p0)
            d_2300_v21_: _dafny.Seq
            d_2300_v21_ = _dafny.SeqWithoutIsStrInference([(self).f10])
            d_2301_v22_: _dafny.Map
            d_2301_v22_ = _dafny.Map({p0: d_2300_v21_})
            d_2302_v23_: _dafny.Set
            d_2302_v23_ = _dafny.Set({_dafny.CodePoint('d')})
            d_2303_v24_: _dafny.Map
            d_2303_v24_ = _dafny.Map({((d_2299_v20_).cf8) - (len(((d_2301_v22_)[p0] if (p0) in (d_2301_v22_) else d_2300_v21_))): (d_2302_v23_).issubset(d_2302_v23_)})
            d_2303_v24_ = (d_2303_v24_).set(p0, (self).f10)
            d_2304_v25_: _dafny.Array
            nw390_ = _dafny.Array(False, 10)
            d_2304_v25_ = nw390_
            d_2305_v26_: D6
            d_2305_v26_ = D6_DC16((self).f9)
            d_2306_v27_: _dafny.Set
            d_2306_v27_ = _dafny.Set({-117})
            d_2307_v28_: D23
            d_2307_v28_ = D23_DC59(d_2305_v26_, (self).f10, len(d_2306_v27_), (D5_DC13((D5_DC13(d_2304_v25_)).cf24)).cf24)
            d_2308_v29_: _dafny.Array
            nw391_ = _dafny.Array(None, 5)
            nw391_[int(0)] = d_2304_v25_
            nw391_[int(1)] = d_2304_v25_
            nw391_[int(2)] = d_2304_v25_
            nw391_[int(3)] = (d_2307_v28_).cf109
            nw391_[int(4)] = d_2304_v25_
            d_2308_v29_ = nw391_
            d_2308_v29_ = d_2308_v29_
            d_2309_v30_: _dafny.Seq
            d_2309_v30_ = _dafny.SeqWithoutIsStrInference([p0])
            d_2310_v31_: _dafny.Array
            nw392_ = _dafny.Array(None, 2)
            nw392_[int(0)] = (d_2309_v30_) + (d_2309_v30_)
            nw392_[int(1)] = d_2309_v30_
            d_2310_v31_ = nw392_
            index407_ = default__.safeIndex(487, (d_2310_v31_).length(0))
            (d_2310_v31_)[index407_] = (d_2309_v30_) + (_dafny.SeqWithoutIsStrInference([p0]))
            index408_ = default__.safeIndex(487, (d_2310_v31_).length(0))
            (d_2310_v31_)[index408_] = (_dafny.SeqWithoutIsStrInference([-744 for d_2311_i3_ in range(default__.abs(423))])) + (d_2309_v30_)
            d_2303_v24_ = (d_2303_v24_).set((0) - (p0), (self).f10)
            d_2300_v21_ = (d_2300_v21_).set(default__.safeIndex(p0, len(d_2300_v21_)), (self).f10)
        hi8_ = 643
        for d_2312_i4_ in range(p0, hi8_):
            d_2313_v32_: _dafny.Map
            d_2313_v32_ = _dafny.Map({(self).f10: d_2312_i4_})
            d_2314_v33_: _dafny.Seq
            d_2314_v33_ = _dafny.SeqWithoutIsStrInference([len(d_2313_v32_)])
            source30_ = D16_DC42((self).f10, len(_dafny.Map({(self).f10: not(False)})), ((_dafny.MultiSet(d_2314_v33_)).cardinality if (self).f10 else d_2312_i4_))
            if source30_.is_DC41:
                d_2315___mcc_h0_ = source30_.cf77
                d_2316___mcc_h1_ = source30_.cf78
                d_2317___mcc_h2_ = source30_.cf79
                d_2318_cf79_ = d_2317___mcc_h2_
                d_2319_cf78_ = d_2316___mcc_h1_
                d_2320_cf77_ = d_2315___mcc_h0_
                d_2321_v34_: _dafny.Array
                nw393_ = _dafny.Array(False, 1)
                d_2321_v34_ = nw393_
                d_2322_v35_: T1
                nw394_ = C2()
                nw394_.ctor__(d_2318_cf79_, d_2321_v34_, (self).f9, d_2320_cf77_)
                d_2322_v35_ = nw394_
                d_2323_v36_: _dafny.Seq
                d_2323_v36_ = _dafny.SeqWithoutIsStrInference([d_2322_v35_, d_2322_v35_])
                d_2323_v36_ = d_2323_v36_
                index409_ = default__.safeIndex(545, (d_2321_v34_).length(0))
                (d_2321_v34_)[index409_] = d_2320_cf77_
                index410_ = default__.safeIndex(545, (d_2321_v34_).length(0))
                (d_2321_v34_)[index410_] = not((self).f10)
                d_2324_v37_: _dafny.Array
                nw395_ = _dafny.Array(None, 11)
                nw395_[int(0)] = d_2276_v0_
                nw395_[int(1)] = d_2276_v0_
                nw395_[int(2)] = d_2276_v0_
                nw395_[int(3)] = d_2276_v0_
                nw395_[int(4)] = d_2276_v0_
                nw395_[int(5)] = d_2276_v0_
                nw395_[int(6)] = d_2276_v0_
                nw395_[int(7)] = d_2276_v0_
                nw395_[int(8)] = d_2276_v0_
                nw395_[int(9)] = d_2276_v0_
                nw395_[int(10)] = _dafny.CodePoint('y')
                d_2324_v37_ = nw395_
                d_2325_v38_: D25
                d_2325_v38_ = D25_DC61(d_2324_v37_)
                d_2326_v39_: _dafny.Seq
                d_2326_v39_ = _dafny.SeqWithoutIsStrInference([d_2324_v37_, d_2324_v37_, d_2324_v37_, d_2324_v37_])
                d_2327_v40_: _dafny.Map
                d_2327_v40_ = _dafny.Map({(d_2321_v34_)[default__.safeIndex(545, (d_2321_v34_).length(0))]: d_2313_v32_})
                d_2328_v41_: _dafny.Array
                nw396_ = _dafny.Array(None, 13)
                nw396_[int(0)] = d_2324_v37_
                nw396_[int(1)] = (d_2325_v38_).cf111
                nw396_[int(2)] = d_2324_v37_
                nw396_[int(3)] = d_2324_v37_
                nw396_[int(4)] = d_2324_v37_
                nw396_[int(5)] = (d_2326_v39_)[default__.safeIndex(len(d_2327_v40_), len(d_2326_v39_))]
                nw396_[int(6)] = d_2324_v37_
                nw396_[int(7)] = d_2324_v37_
                nw396_[int(8)] = d_2324_v37_
                nw396_[int(9)] = d_2324_v37_
                nw396_[int(10)] = d_2324_v37_
                nw396_[int(11)] = d_2324_v37_
                nw396_[int(12)] = (d_2324_v37_ if (d_2322_v35_).f10 else d_2324_v37_)
                d_2328_v41_ = nw396_
                index411_ = default__.safeIndex(758, (d_2328_v41_).length(0))
                (d_2328_v41_)[index411_] = (d_2324_v37_ if d_2318_cf79_ else d_2324_v37_)
                index412_ = default__.safeIndex(758, (d_2328_v41_).length(0))
                rhs385_ = d_2312_i4_
                rhs386_ = d_2324_v37_
                lhs284_ = d_2328_v41_
                lhs285_ = default__.safeIndex(758, (d_2328_v41_).length(0))
                r2 = rhs385_
                lhs284_[lhs285_] = rhs386_
                d_2329_v42_: C4
                nw397_ = C4()
                nw397_.ctor__((d_2322_v35_).f9, (self).f10)
                d_2329_v42_ = nw397_
            elif source30_.is_DC42:
                d_2330___mcc_h3_ = source30_.cf80
                d_2331___mcc_h4_ = source30_.cf81
                d_2332___mcc_h5_ = source30_.cf82
                d_2333_cf82_ = d_2332___mcc_h5_
                d_2334_cf81_ = d_2331___mcc_h4_
                d_2335_cf80_ = d_2330___mcc_h3_
                d_2336_v43_: C1
                nw398_ = C1()
                nw398_.ctor__(d_2335_cf80_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "knxgb")))
                d_2336_v43_ = nw398_
                d_2337_v44_: D26
                d_2337_v44_ = D26_DC64(d_2336_v43_)
                d_2338_v45_: _dafny.Set
                d_2338_v45_ = _dafny.Set({d_2336_v43_, d_2336_v43_, d_2336_v43_, d_2336_v43_, (d_2337_v44_).cf114})
                d_2339_v47_: _dafny.Array
                nw399_ = _dafny.Array(None, 25)
                nw399_[int(0)] = ((self).f10) and ((self).f10)
                nw399_[int(1)] = d_2335_cf80_
                nw399_[int(2)] = d_2335_cf80_
                nw399_[int(3)] = not (d_2335_cf80_) or (d_2335_cf80_)
                nw399_[int(4)] = True
                nw399_[int(5)] = (d_2338_v45_).isdisjoint(_dafny.Set({d_2336_v43_}))
                nw399_[int(6)] = False
                nw399_[int(7)] = (self).f10
                nw399_[int(8)] = (d_2336_v43_).f12
                nw399_[int(9)] = True
                nw399_[int(10)] = (self).f10
                def iife189_():
                    coll77_ = _dafny.Set()
                    compr_77_: int
                    for compr_77_ in _dafny.IntegerRange(707, 683):
                        d_2340_v46_: int = compr_77_
                        if ((707) <= (d_2340_v46_)) and ((d_2340_v46_) < (683)):
                            coll77_ = coll77_.union(_dafny.Set([default__.safeModuloInt(d_2340_v46_, d_2312_i4_)]))
                    return _dafny.Set(coll77_)
                nw399_[int(11)] = (iife189_()
                ).issubset(default__.fm25(d_2334_cf81_, d_2333_cf82_, globalState))
                nw399_[int(12)] = (d_2333_cf82_) <= (p0)
                nw399_[int(13)] = d_2335_cf80_
                nw399_[int(14)] = (d_2336_v43_).f12
                nw399_[int(15)] = not (d_2335_cf80_) or (d_2335_cf80_)
                nw399_[int(16)] = (d_2334_cf81_) == ((self).fm1(_dafny.Map({(d_2336_v43_).f12: d_2335_cf80_}), globalState))
                nw399_[int(17)] = (self).f10
                nw399_[int(18)] = d_2335_cf80_
                nw399_[int(19)] = (d_2336_v43_).f12
                nw399_[int(20)] = d_2335_cf80_
                nw399_[int(21)] = ((self).f10 if d_2335_cf80_ else not((self).f10))
                nw399_[int(22)] = d_2335_cf80_
                nw399_[int(23)] = (self).f10
                nw399_[int(24)] = (d_2335_cf80_) or ((self).f10)
                d_2339_v47_ = nw399_
                d_2339_v47_ = d_2339_v47_
                (globalState).f8 = d_2335_cf80_
                d_2341_v48_: _dafny.Map
                d_2341_v48_ = _dafny.Map({(630 if not((d_2336_v43_).fm32(d_2312_i4_, d_2276_v0_, globalState)) else d_2334_cf81_): d_2339_v47_})
                d_2339_v47_ = ((d_2341_v48_)[default__.safeModuloInt(524, p0)] if (default__.safeModuloInt(524, p0)) in (d_2341_v48_) else d_2339_v47_)
                d_2342_v49_: _dafny.Seq
                d_2342_v49_ = _dafny.SeqWithoutIsStrInference([(self).f10, d_2335_cf80_])
                d_2343_v50_: _dafny.Map
                d_2343_v50_ = _dafny.Map({_dafny.CodePoint('c'): (self).f10})
                d_2344_v51_: C4
                nw400_ = C4()
                nw400_.ctor__(default__.fm48((self).fm0(d_2342_v49_, (d_2336_v43_).f12, len(d_2343_v50_), (self).f10, globalState), globalState), d_2335_cf80_)
                d_2344_v51_ = nw400_
                d_2345_v52_: _dafny.Map
                d_2345_v52_ = _dafny.Map({(d_2336_v43_).f13: len(d_2342_v49_)})
                d_2346_v53_: T0
                nw401_ = C1()
                nw401_.ctor__((d_2334_cf81_) in (_dafny.Set({d_2334_cf81_, (d_2314_v33_)[default__.safeIndex(len(d_2342_v49_), len(d_2314_v33_))], (0) - ((0) - ((0) - (((d_2345_v52_)[(d_2336_v43_).f13] if ((d_2336_v43_).f13) in (d_2345_v52_) else p0))))})), (d_2336_v43_).f13)
                d_2346_v53_ = nw401_
                d_2347_v54_: _dafny.Map
                d_2347_v54_ = _dafny.Map({(d_2336_v43_).f12: d_2339_v47_})
                d_2348_v55_: _dafny.Map
                d_2348_v55_ = _dafny.Map({len((d_2336_v43_).f13): (d_2336_v43_).f13})
                rhs387_ = d_2344_v51_
                rhs388_ = d_2346_v53_
                rhs389_ = ((_dafny.MultiSet([d_2334_cf81_, d_2312_i4_, len(((d_2348_v55_)[(d_2314_v33_)[default__.safeIndex(d_2333_cf82_, len(d_2314_v33_))]] if ((d_2314_v33_)[default__.safeIndex(d_2333_cf82_, len(d_2314_v33_))]) in (d_2348_v55_) else (d_2336_v43_).f13)), d_2312_i4_])) | (_dafny.MultiSet([d_2312_i4_]))).intersection(default__.fm46(d_2335_cf80_, (d_2336_v43_).f12, not((d_2336_v43_).f12), globalState))
                rhs390_ = ((d_2312_i4_ if (d_2344_v51_).fm0(d_2342_v49_, (d_2336_v43_).f12, d_2312_i4_, (d_2336_v43_).f12, globalState) else d_2333_cf82_)) * (default__.safeDivisionInt(p0, len((d_2336_v43_).f13)))
                rhs391_ = d_2347_v54_
                d_2344_v51_ = rhs387_
                d_2346_v53_ = rhs388_
                d_2282_v6_ = rhs389_
                d_2334_cf81_ = rhs390_
                d_2347_v54_ = rhs391_
            elif True:
                d_2349___mcc_h6_ = source30_.cf76
                d_2350_cf76_ = d_2349___mcc_h6_
                d_2351_v56_: C0
                nw402_ = C0()
                nw402_.ctor__()
                d_2351_v56_ = nw402_
                d_2281_v5_ = d_2281_v5_
                d_2352_v57_: C0
                nw403_ = C0()
                nw403_.ctor__()
                d_2352_v57_ = nw403_
                r2 = p0
            d_2282_v6_ = (d_2282_v6_) - (d_2282_v6_)
            if not((self).f10):
                d_2353_v58_: _dafny.MultiSet
                d_2353_v58_ = _dafny.MultiSet([d_2276_v0_, d_2276_v0_])
                d_2354_v59_: _dafny.Seq
                d_2354_v59_ = _dafny.SeqWithoutIsStrInference([d_2276_v0_, d_2276_v0_])
                d_2355_v60_: _dafny.Map
                d_2355_v60_ = _dafny.Map({p0: d_2353_v58_})
                d_2356_v61_: _dafny.Seq
                d_2356_v61_ = _dafny.SeqWithoutIsStrInference([d_2353_v58_])
                d_2357_v62_: _dafny.Array
                nw404_ = _dafny.Array(False, 1)
                d_2357_v62_ = nw404_
                d_2358_v63_: _dafny.MultiSet
                d_2358_v63_ = _dafny.MultiSet([d_2357_v62_])
                d_2359_v64_: _dafny.MultiSet
                d_2359_v64_ = d_2353_v58_
                d_2360_v65_: _dafny.Array
                nw405_ = _dafny.Array(None, 28)
                nw405_[int(0)] = (_dafny.MultiSet([d_2276_v0_, d_2276_v0_, d_2276_v0_])) - (d_2353_v58_)
                nw405_[int(1)] = _dafny.MultiSet(d_2354_v59_)
                nw405_[int(2)] = d_2353_v58_
                nw405_[int(3)] = ((d_2355_v60_)[p0] if (p0) in (d_2355_v60_) else ((d_2355_v60_)[d_2312_i4_] if (d_2312_i4_) in (d_2355_v60_) else d_2353_v58_))
                nw405_[int(4)] = _dafny.MultiSet([d_2276_v0_, d_2276_v0_])
                nw405_[int(5)] = d_2353_v58_
                nw405_[int(6)] = d_2353_v58_
                nw405_[int(7)] = d_2353_v58_
                nw405_[int(8)] = (d_2353_v58_).set(_dafny.CodePoint('i'), default__.abs(d_2312_i4_))
                nw405_[int(9)] = (d_2356_v61_)[default__.safeIndex((d_2358_v63_).cardinality, len(d_2356_v61_))]
                nw405_[int(10)] = (d_2353_v58_).set(d_2276_v0_, default__.abs(274))
                nw405_[int(11)] = d_2353_v58_
                nw405_[int(12)] = d_2353_v58_
                nw405_[int(13)] = (d_2359_v64_)
                nw405_[int(14)] = (d_2353_v58_) - (_dafny.MultiSet(d_2354_v59_))
                nw405_[int(15)] = d_2353_v58_
                nw405_[int(16)] = (d_2353_v58_) | (_dafny.MultiSet(d_2354_v59_))
                nw405_[int(17)] = d_2353_v58_
                nw405_[int(18)] = d_2353_v58_
                nw405_[int(19)] = d_2353_v58_
                nw405_[int(20)] = d_2353_v58_
                nw405_[int(21)] = d_2353_v58_
                nw405_[int(22)] = d_2353_v58_
                nw405_[int(23)] = d_2353_v58_
                nw405_[int(24)] = d_2353_v58_
                nw405_[int(25)] = (d_2353_v58_) | (d_2353_v58_)
                nw405_[int(26)] = d_2353_v58_
                nw405_[int(27)] = d_2353_v58_
                d_2360_v65_ = nw405_
                d_2361_v66_: _dafny.Seq
                d_2361_v66_ = _dafny.SeqWithoutIsStrInference([(self).f10])
                index413_ = default__.safeIndex(999, (d_2360_v65_).length(0))
                (d_2360_v65_)[index413_] = (d_2353_v58_).intersection(default__.fm57(len(d_2361_v66_), globalState))
                index414_ = default__.safeIndex(999, (d_2360_v65_).length(0))
                (d_2360_v65_)[index414_] = ((d_2353_v58_).intersection(d_2353_v58_)) - (((d_2356_v61_)[default__.safeIndex(p0, len(d_2356_v61_))] if (self).f10 else d_2353_v58_))
                d_2362_v67_: D6
                d_2362_v67_ = D6_DC16((self).f9)
                d_2363_v68_: C5
                nw406_ = C5()
                nw406_.ctor__((d_2362_v67_).cf28, (self).f10)
                d_2363_v68_ = nw406_
                d_2364_v69_: _dafny.Array
                nw407_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 10)
                d_2364_v69_ = nw407_
                d_2365_v70_: _dafny.Map
                d_2365_v70_ = _dafny.Map({p0: d_2354_v59_})
                d_2366_v71_: _dafny.Map
                d_2366_v71_ = _dafny.Map({d_2364_v69_: (_dafny.SeqWithoutIsStrInference([d_2276_v0_ for d_2367_i5_ in range(default__.abs(733))])) + (((d_2365_v70_)[p0] if (p0) in (d_2365_v70_) else d_2354_v59_))})
                rhs392_ = ((d_2366_v71_)[d_2364_v69_] if (d_2364_v69_) in (d_2366_v71_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sxeiw")))
                rhs393_ = d_2363_v68_
                d_2354_v59_ = rhs392_
                d_2363_v68_ = rhs393_
                d_2276_v0_ = d_2276_v0_
                nw408_ = _dafny.Array(False, 28)
                d_2357_v62_ = nw408_
                r2 = len(((_dafny.SeqWithoutIsStrInference([d_2276_v0_])) + (_dafny.SeqWithoutIsStrInference([d_2276_v0_, default__.fm36(globalState), d_2276_v0_]))) + ((d_2354_v59_) + (d_2354_v59_)))
            elif True:
                d_2368_v72_: _dafny.Seq
                d_2368_v72_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p"))
                d_2369_v73_: C1
                nw409_ = C1()
                nw409_.ctor__((self).f10, d_2368_v72_)
                d_2369_v73_ = nw409_
                d_2370_v74_: D26
                d_2370_v74_ = D26_DC64(d_2369_v73_)
                d_2371_v75_: C3
                nw410_ = C3()
                nw410_.ctor__()
                d_2371_v75_ = nw410_
                pat_let_tv53_ = d_2369_v73_
                def iife190_(_pat_let56_0):
                    def iife191_(d_2372_dt__update__tmp_h0_):
                        def iife192_(_pat_let57_0):
                            def iife193_(d_2373_dt__update_hcf114_h0_):
                                return D26_DC64(d_2373_dt__update_hcf114_h0_)
                            return iife193_(_pat_let57_0)
                        return iife192_(pat_let_tv53_)
                    return iife191_(_pat_let56_0)
                rhs394_ = d_2312_i4_
                rhs395_ = iife190_(D26_DC64(d_2369_v73_))
                rhs396_ = (d_2369_v73_).f12
                rhs397_ = (d_2371_v75_ if ((self).f10) == ((d_2369_v73_).f12) else d_2371_v75_)
                lhs286_ = globalState
                r2 = rhs394_
                d_2370_v74_ = rhs395_
                lhs286_.f8 = rhs396_
                d_2371_v75_ = rhs397_
                r2 = p0
                d_2374_v76_: C0
                nw411_ = C0()
                nw411_.ctor__()
                d_2374_v76_ = nw411_
                r2 = (0) - (default__.safeModuloInt(d_2312_i4_, len((d_2369_v73_).f13)))
                d_2375_v77_: _dafny.Array
                def lambda118_(d_2376_p0_):
                    def lambda119_(d_2377_i6_):
                        return (d_2377_i6_) * (d_2376_p0_)

                    return lambda119_

                init70_ = lambda118_(p0)
                nw412_ = _dafny.Array(None, 27)
                for i0_70_ in range(nw412_.length(0)):
                    nw412_[i0_70_] = init70_(i0_70_)
                d_2375_v77_ = nw412_
                d_2375_v77_ = d_2375_v77_
            d_2378_v79_: _dafny.Set
            d_2378_v79_ = _dafny.Set({p0})
            d_2379_v80_: _dafny.Map
            def iife194_():
                coll78_ = _dafny.Set()
                compr_78_: int
                for compr_78_ in _dafny.IntegerRange(-767, 994):
                    d_2380_v78_: int = compr_78_
                    if ((-767) <= (d_2380_v78_)) and ((d_2380_v78_) < (994)):
                        coll78_ = coll78_.union(_dafny.Set([default__.safeModuloInt(d_2380_v78_, d_2312_i4_)]))
                return _dafny.Set(coll78_)
            def iife195_():
                coll79_ = _dafny.Set()
                compr_79_: int
                for compr_79_ in _dafny.IntegerRange(-767, 994):
                    d_2381_v78_: int = compr_79_
                    if ((-767) <= (d_2381_v78_)) and ((d_2381_v78_) < (994)):
                        coll79_ = coll79_.union(_dafny.Set([default__.safeModuloInt(d_2381_v78_, d_2312_i4_)]))
                return _dafny.Set(coll79_)
            d_2379_v80_ = _dafny.Map({p0: (D28_DC67((_dafny.SeqWithoutIsStrInference([_dafny.Set({p0}), iife194_()
, d_2378_v79_, d_2378_v79_, d_2378_v79_])).set(default__.safeIndex(d_2312_i4_, len(_dafny.SeqWithoutIsStrInference([_dafny.Set({p0}), iife195_()
, d_2378_v79_, d_2378_v79_, d_2378_v79_]))), d_2378_v79_))).cf119})
            d_2382_v81_: _dafny.Seq
            d_2382_v81_ = _dafny.SeqWithoutIsStrInference([d_2378_v79_])
            if (default__.fm58(globalState)) != (((d_2379_v80_)[-615] if (-615) in (d_2379_v80_) else d_2382_v81_)):
                (globalState).f8 = (self).f10
                d_2383_v82_: _dafny.Seq
                d_2383_v82_ = _dafny.SeqWithoutIsStrInference([(self).f11, (self).f11, (self).f11])
                rhs398_ = True
                rhs399_ = (len(d_2383_v82_)) - (d_2312_i4_)
                rhs400_ = (self).f10
                rhs401_ = d_2312_i4_
                lhs287_ = globalState
                r1 = rhs398_
                r2 = rhs399_
                lhs287_.f8 = rhs400_
                r2 = rhs401_
                d_2384_v83_: _dafny.Seq
                d_2384_v83_ = _dafny.SeqWithoutIsStrInference([(self).f10])
                (globalState).f8 = (p0) <= (len(d_2384_v83_))
                r2 = p0
                d_2385_v84_: T0
                nw413_ = C1()
                nw413_.ctor__(not((self).f10), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ukpsmds")))
                d_2385_v84_ = nw413_
                d_2386_v85_: _dafny.Map
                d_2386_v85_ = _dafny.Map({p0: d_2385_v84_})
                d_2386_v85_ = (d_2386_v85_).set((_dafny.MultiSet((d_2384_v83_ if (self).f10 else d_2384_v83_))).cardinality, d_2385_v84_)
            elif True:
                r2 = (p0) * ((self).fm1((self).f9, globalState))
                d_2387_v86_: _dafny.Seq
                d_2387_v86_ = _dafny.SeqWithoutIsStrInference([(self).f10])
                d_2388_v87_: C0
                nw414_ = C0()
                nw414_.ctor__()
                d_2388_v87_ = nw414_
                d_2389_v88_: _dafny.MultiSet
                d_2389_v88_ = _dafny.MultiSet([d_2388_v87_])
                d_2390_v89_: _dafny.Array
                nw415_ = _dafny.Array(None, 22)
                nw415_[int(0)] = not((self).f10)
                nw415_[int(1)] = (self).f10
                nw415_[int(2)] = (self).f10
                nw415_[int(3)] = (self).f10
                nw415_[int(4)] = (self).f10
                nw415_[int(5)] = (self).f10
                nw415_[int(6)] = not (not((self).f10)) or ((self).f10)
                nw415_[int(7)] = (d_2312_i4_) == (len(d_2387_v86_))
                nw415_[int(8)] = (self).f10
                nw415_[int(9)] = not((self).f10)
                nw415_[int(10)] = (self).f10
                nw415_[int(11)] = (d_2389_v88_).ispropersubset(d_2389_v88_)
                nw415_[int(12)] = (self).f10
                nw415_[int(13)] = not((self).f10)
                nw415_[int(14)] = (self).f10
                nw415_[int(15)] = not((self).f10)
                nw415_[int(16)] = (self).f10
                nw415_[int(17)] = (self).f10
                nw415_[int(18)] = (self).f10
                nw415_[int(19)] = (d_2387_v86_)[default__.safeIndex(d_2312_i4_, len(d_2387_v86_))]
                nw415_[int(20)] = (self).f10
                nw415_[int(21)] = (self).f10
                d_2390_v89_ = nw415_
                index415_ = default__.safeIndex(184, (d_2390_v89_).length(0))
                (d_2390_v89_)[index415_] = (self).f10
                index416_ = default__.safeIndex(184, (d_2390_v89_).length(0))
                (d_2390_v89_)[index416_] = (self).f10
                (globalState).f8 = ((d_2390_v89_)[default__.safeIndex(184, (d_2390_v89_).length(0))]) == ((d_2390_v89_)[default__.safeIndex(184, (d_2390_v89_).length(0))])
                d_2391_v90_: D7
                d_2391_v90_ = D7_DC21((self).f11)
                d_2392_v91_: D7
                d_2392_v91_ = D7_DC21((d_2391_v90_).cf40)
                d_2393_v92_: _dafny.Map
                d_2393_v92_ = _dafny.Map({((self).f10 if (d_2390_v89_)[default__.safeIndex(184, (d_2390_v89_).length(0))] else (self).f10): ((d_2392_v91_).cf40).ispropersubset((self).f11)})
                d_2393_v92_ = d_2393_v92_
                index417_ = default__.safeIndex(184, (d_2390_v89_).length(0))
                (d_2390_v89_)[index417_] = not ((d_2390_v89_)[default__.safeIndex(184, (d_2390_v89_).length(0))]) or ((self).f10)
        d_2394_v93_: _dafny.Seq
        d_2394_v93_ = _dafny.SeqWithoutIsStrInference([p0, -457, p0])
        d_2395_v94_: _dafny.Array
        nw416_ = _dafny.Array(None, 10)
        nw416_[int(0)] = 864
        nw416_[int(1)] = p0
        nw416_[int(2)] = p0
        nw416_[int(3)] = (d_2394_v93_)[default__.safeIndex((0) - (default__.fm33(not((self).f10), (self).f10, p0, (self).f10, globalState)), len(d_2394_v93_))]
        nw416_[int(4)] = p0
        nw416_[int(5)] = p0
        nw416_[int(6)] = p0
        nw416_[int(7)] = p0
        nw416_[int(8)] = p0
        nw416_[int(9)] = p0
        d_2395_v94_ = nw416_
        d_2396_v95_: _dafny.MultiSet
        d_2396_v95_ = _dafny.MultiSet([d_2395_v94_, d_2395_v94_, d_2395_v94_])
        hi9_ = (p0) - (p0)
        for d_2397_i7_ in range((0) - ((d_2396_v95_).cardinality), hi9_):
            d_2398_v96_: C5
            nw417_ = C5()
            nw417_.ctor__((_dafny.Map({(self).f10: (self).f10})).set((self).f10, not((self).f10)), not((self).f10))
            d_2398_v96_ = nw417_
            d_2398_v96_ = d_2398_v96_
            d_2399_v97_: _dafny.Seq
            d_2399_v97_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ax"))
            d_2400_v98_: _dafny.Map
            d_2400_v98_ = _dafny.Map({len(d_2399_v97_): (self).f10})
            rhs402_ = (d_2398_v96_).fm18((self).f11, (d_2399_v97_) + ((d_2399_v97_).set(default__.safeIndex(p0, len(d_2399_v97_)), default__.fm31((self).f10, d_2400_v98_, globalState))), globalState)
            rhs403_ = d_2397_i7_
            rhs404_ = (self).f10
            lhs288_ = globalState
            r1 = rhs402_
            r2 = rhs403_
            lhs288_.f8 = rhs404_
            index418_ = default__.safeIndex(130, (d_2395_v94_).length(0))
            (d_2395_v94_)[index418_] = d_2397_i7_
            d_2401_v99_: _dafny.Map
            d_2401_v99_ = _dafny.Map({(self).f10: (d_2394_v93_)[default__.safeIndex(p0, len(d_2394_v93_))]})
            index419_ = default__.safeIndex(130, (d_2395_v94_).length(0))
            (d_2395_v94_)[index419_] = (d_2397_i7_) - ((len(d_2401_v99_)) - ((0) - (d_2397_i7_)))
            d_2402_v100_: C0
            nw418_ = C0()
            nw418_.ctor__()
            d_2402_v100_ = nw418_
        d_2403_v101_: _dafny.Seq
        d_2403_v101_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gbx"))
        r0 = (default__.fm10((self).f10, len(d_2394_v93_), d_2403_v101_, d_2283_v7_, globalState)) - (d_2282_v6_)
        r1 = True
        d_2404_v102_: _dafny.Seq
        d_2404_v102_ = _dafny.SeqWithoutIsStrInference([(self).f10])
        r2 = len((_dafny.SeqWithoutIsStrInference([(not((self).f10) if (self).f10 else False)])).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference([(not((self).f10) if (self).f10 else False)]))), (not((self).f10)) in (d_2404_v102_)))
        r3 = d_2403_v101_
        return r0, r1, r2, r3

    @property
    def f11(self):
        return self._f11
