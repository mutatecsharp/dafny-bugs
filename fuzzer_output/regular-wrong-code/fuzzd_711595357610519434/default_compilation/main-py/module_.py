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
    def fm0(p0, p1, p2, globalState):
        source0_ = _dafny.Map({True: False})
        d_0___mcc_h0_ = source0_
        d_1_cf18_ = d_0___mcc_h0_
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: int
            for compr_0_ in _dafny.IntegerRange(538, 46):
                d_2_v0_: int = compr_0_
                if ((538) <= (d_2_v0_)) and ((d_2_v0_) < (46)):
                    coll0_[(d_2_v0_) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kdylahw"))))] = _dafny.Set({True})
            return _dafny.Map(coll0_)
        return iife0_()
        

    @staticmethod
    def fm1(p0, p1, p2, p3, globalState):
        if False:
            return (_dafny.Set({_dafny.Map({True: not(False)}), _dafny.Map({True: False})})).ispropersubset(_dafny.Set({_dafny.Map({False: False})}))
        elif True:
            return not (True) or (False)

    @staticmethod
    def fm2(p0, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jkyhcsar"))

    @staticmethod
    def fm4(p0, p1, p2, globalState):
        source1_ = D17_DC41(-267, not(True))
        if source1_.is_DC41:
            d_3___mcc_h0_ = source1_.cf77
            d_4___mcc_h1_ = source1_.cf78
            d_5_cf78_ = d_4___mcc_h1_
            d_6_cf77_ = d_3___mcc_h0_
            return _dafny.CodePoint('b')
        elif source1_.is_DC40:
            d_7___mcc_h2_ = source1_.cf76
            d_8_cf76_ = d_7___mcc_h2_
            return _dafny.CodePoint('u')
        elif True:
            d_9___mcc_h3_ = source1_.cf79
            d_10_cf79_ = d_9___mcc_h3_
            return _dafny.CodePoint('y')

    @staticmethod
    def fm7(globalState):
        return _dafny.SeqWithoutIsStrInference([548, -508])

    @staticmethod
    def fm8(p0, p1, p2, p3, globalState):
        source2_ = D14_DC36(True, 110, (0) - ((_dafny.MultiSet([len(_dafny.Map({True: len(_dafny.Set({True, not(not(False))}))}))])).cardinality), _dafny.Map({888: True}))
        if source2_.is_DC36:
            d_11___mcc_h0_ = source2_.cf68
            d_12___mcc_h1_ = source2_.cf69
            d_13___mcc_h2_ = source2_.cf70
            d_14___mcc_h3_ = source2_.cf71
            d_15_cf71_ = d_14___mcc_h3_
            d_16_cf70_ = d_13___mcc_h2_
            d_17_cf69_ = d_12___mcc_h1_
            d_18_cf68_ = d_11___mcc_h0_
            return _dafny.CodePoint('q')
        elif True:
            d_19___mcc_h4_ = source2_.cf67
            d_20_cf67_ = d_19___mcc_h4_
            return _dafny.CodePoint('h')

    @staticmethod
    def fm9(globalState):
        source3_ = D1_DC5()
        if source3_.is_DC5:
            def iife1_():
                coll1_ = _dafny.Set()
                compr_1_: _dafny.MultiSet
                for compr_1_ in (_dafny.Map({_dafny.MultiSet([True]): 213})).keys.Elements:
                    d_21_v0_: _dafny.MultiSet = compr_1_
                    if (d_21_v0_) in (_dafny.Map({_dafny.MultiSet([True]): 213})):
                        coll1_ = coll1_.union(_dafny.Set([d_21_v0_]))
                return _dafny.Set(coll1_)
            return (_dafny.Set({_dafny.MultiSet([not(True), not(False)]), _dafny.MultiSet([True, False])})) - (iife1_()
            )
        elif True:
            d_22___mcc_h0_ = source3_.cf11
            d_23_cf11_ = d_22___mcc_h0_
            return (_dafny.Set({_dafny.MultiSet([True, True, False, True, False])})) - (_dafny.Set({_dafny.MultiSet([True]), _dafny.MultiSet([False])}))

    @staticmethod
    def fm10(p0, p1, p2, globalState):
        return ((_dafny.Set({-989})) | (_dafny.Set({922}))) - ((_dafny.Set({len(_dafny.Map({True: (0) - (len(_dafny.SeqWithoutIsStrInference([True, True])))}))})) | (_dafny.Set({len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([17, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bbpgrxd")))])): False})), len(_dafny.Map({len(_dafny.Map({791: not(not(True))})): 854})), 439})))

    @staticmethod
    def fm12(globalState):
        source4_ = D1_DC4(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({False}))]))
        if source4_.is_DC5:
            if True:
                return _dafny.CodePoint('u')
            elif True:
                return _dafny.CodePoint('u')
        elif True:
            d_24___mcc_h0_ = source4_.cf11
            d_25_cf11_ = d_24___mcc_h0_
            if not(not(False)):
                return _dafny.CodePoint('r')
            elif True:
                return _dafny.CodePoint('b')

    @staticmethod
    def fm13(p0, p1, globalState):
        source5_ = D5_DC13(True, not(not(False)), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_26_i0_ in range(default__.abs(-59))]), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_27_i1_ in range(default__.abs(921))])))
        if source5_.is_DC12:
            d_28___mcc_h0_ = source5_.cf20
            d_29_cf20_ = d_28___mcc_h0_
            return D0_DC1(d_29_cf20_, False, d_29_cf20_, not(True))
        elif source5_.is_DC13:
            d_30___mcc_h1_ = source5_.cf21
            d_31___mcc_h2_ = source5_.cf22
            d_32___mcc_h3_ = source5_.cf23
            d_33___mcc_h4_ = source5_.cf24
            d_34_cf24_ = d_33___mcc_h4_
            d_35_cf23_ = d_32___mcc_h3_
            d_36_cf22_ = d_31___mcc_h2_
            d_37_cf21_ = d_30___mcc_h1_
            return D0_DC1(d_34_cf24_, d_36_cf22_, 842, d_36_cf22_)
        elif source5_.is_DC11:
            d_38___mcc_h5_ = source5_.cf19
            d_39_cf19_ = d_38___mcc_h5_
            return D0_DC1(651, not(False), -746, False)
        elif True:
            d_40___mcc_h6_ = source5_.cf25
            d_41_cf25_ = d_40___mcc_h6_
            return D0_DC1(114, True, -419, not(True))

    @staticmethod
    def fm14(p0, p1, p2, globalState):
        if not(not((True if True else True))):
            return _dafny.Map({(D17_DC41(377, True)).cf78: True})
        elif True:
            return (_dafny.Map({True: False})) | (_dafny.Map({True: True}))

    @staticmethod
    def fm15(p0, p1, globalState):
        return (default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([False, False])), 183)) - ((879) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "raqgm")))))

    @staticmethod
    def fm19(p0, p1, p2, p3, globalState):
        def iife2_():
            coll2_ = _dafny.Set()
            compr_2_: int
            for compr_2_ in _dafny.IntegerRange(180, 64):
                d_42_v0_: int = compr_2_
                if ((180) <= (d_42_v0_)) and ((d_42_v0_) < (64)):
                    coll2_ = coll2_.union(_dafny.Set([(d_42_v0_) * (170)]))
            return _dafny.Set(coll2_)
        return D0_DC2(len((_dafny.Set({(_dafny.MultiSet([-412])).cardinality})).intersection(iife2_()
)), _dafny.CodePoint('f'), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "li"))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mhhv"))), len((_dafny.SeqWithoutIsStrInference([len(_dafny.Set({True})), 938])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([False, False, True, True]))])) for d_43_i0_ in range(default__.abs(288))]))))

    @staticmethod
    def fm21(p0, globalState):
        return ((_dafny.SeqWithoutIsStrInference([True])) + (_dafny.SeqWithoutIsStrInference([not(True)]))) + (_dafny.SeqWithoutIsStrInference([not(not(False))]))

    @staticmethod
    def fm22(globalState):
        def iife3_():
            coll3_ = _dafny.Set()
            compr_3_: int
            for compr_3_ in (_dafny.Map({-362: _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ksdltj")))])})).keys.Elements:
                d_44_v0_: int = compr_3_
                if (d_44_v0_) in (_dafny.Map({-362: _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ksdltj")))])})):
                    coll3_ = coll3_.union(_dafny.Set([(d_44_v0_) - (default__.safeModuloInt(428, len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bdnraex")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lktqteah")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_45_i0_ in range(default__.abs(589))])]))))]))
            return _dafny.Set(coll3_)
        return iife3_()
        

    @staticmethod
    def fm24(p0, globalState):
        if (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y")))) > (814):
            return default__.safeDivisionInt(-666, (D17_DC41(94, True)).cf77)
        elif True:
            return default__.safeModuloInt((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hhew")))), 885)

    @staticmethod
    def fm25(p0, p1, p2, globalState):
        return _dafny.SeqWithoutIsStrInference([-408, default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([not(True)])), 572), (0) - (len((_dafny.SeqWithoutIsStrInference([True])) + (_dafny.SeqWithoutIsStrInference([False])))), 547])

    @staticmethod
    def fm26(p0, globalState):
        return (_dafny.Set({False})).intersection(_dafny.Set({False, True, False, True}))

    @staticmethod
    def fm27(globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([not(False), False]), _dafny.SeqWithoutIsStrInference([True, False]), _dafny.SeqWithoutIsStrInference([False]), (_dafny.SeqWithoutIsStrInference([True])) + (_dafny.SeqWithoutIsStrInference([False]))])

    @staticmethod
    def fm28(p0, p1, p2, globalState):
        def iife4_():
            coll4_ = _dafny.Set()
            compr_4_: int
            for compr_4_ in _dafny.IntegerRange(-713, -163):
                d_46_v0_: int = compr_4_
                if ((-713) <= (d_46_v0_)) and ((d_46_v0_) < (-163)):
                    coll4_ = coll4_.union(_dafny.Set([default__.safeDivisionInt(d_46_v0_, len(_dafny.SeqWithoutIsStrInference([False])))]))
            return _dafny.Set(coll4_)
        if (_dafny.Map({D7_DC19(_dafny.MultiSet([False, not(not(False))]), 330): len(_dafny.Map({False: not(not(True))}))})) != (_dafny.Map({D7_DC19(_dafny.MultiSet([not(True)]), len(iife4_()
)): (_dafny.MultiSet([26, 942])).cardinality})):
            if False:
                return _dafny.CodePoint('x')
            elif True:
                return _dafny.CodePoint('r')
        elif True:
            return _dafny.CodePoint('q')
        elif True:
            return _dafny.CodePoint('p')

    @staticmethod
    def fm29(p0, p1, globalState):
        def iife5_():
            coll5_ = _dafny.Map()
            compr_5_: int
            for compr_5_ in (_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False]))).cardinality])).Elements:
                d_47_v0_: int = compr_5_
                if (d_47_v0_) in (_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False]))).cardinality])):
                    coll5_[(d_47_v0_) * (-611)] = False
            return _dafny.Map(coll5_)
        source6_ = D13_DC31(iife5_()
)
        if source6_.is_DC32:
            d_48___mcc_h0_ = source6_.cf56
            d_49___mcc_h1_ = source6_.cf57
            d_50___mcc_h2_ = source6_.cf58
            d_51___mcc_h3_ = source6_.cf59
            d_52___mcc_h4_ = source6_.cf60
            d_53_cf60_ = d_52___mcc_h4_
            d_54_cf59_ = d_51___mcc_h3_
            d_55_cf58_ = d_50___mcc_h2_
            d_56_cf57_ = d_49___mcc_h1_
            d_57_cf56_ = d_48___mcc_h0_
            return D6_DC16((0) - (d_57_cf56_))
        elif source6_.is_DC33:
            d_58___mcc_h5_ = source6_.cf61
            d_59___mcc_h6_ = source6_.cf62
            d_60___mcc_h7_ = source6_.cf63
            d_61___mcc_h8_ = source6_.cf64
            d_62___mcc_h9_ = source6_.cf65
            d_63_cf65_ = d_62___mcc_h9_
            d_64_cf64_ = d_61___mcc_h8_
            d_65_cf63_ = d_60___mcc_h7_
            d_66_cf62_ = d_59___mcc_h6_
            d_67_cf61_ = d_58___mcc_h5_
            return D6_DC16((d_63_cf65_).f9)
        elif source6_.is_DC31:
            d_68___mcc_h10_ = source6_.cf55
            d_69_cf55_ = d_68___mcc_h10_
            return D6_DC16((0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_70_i0_ in range(default__.abs(-504))]))))
        elif True:
            d_71___mcc_h11_ = source6_.cf66
            d_72_cf66_ = d_71___mcc_h11_
            if False:
                return D6_DC16(len(_dafny.Map({33: False})))
            elif True:
                return D6_DC16(746)

    @staticmethod
    def fm30(p0, p1, p2, globalState):
        return _dafny.MultiSet([True])

    @staticmethod
    def fm31(p0, p1, p2, p3, globalState):
        return D6_DC15((_dafny.Set({True})) | (_dafny.Set({True})))

    @staticmethod
    def fm32(globalState):
        return D0_DC3(D0_DC0(_dafny.Map({(_dafny.MultiSet([-345])).cardinality: _dafny.Set({True})})))

    @staticmethod
    def fm33(globalState):
        source7_ = D17_DC41((_dafny.MultiSet([True])).cardinality, False)
        if source7_.is_DC41:
            d_73___mcc_h0_ = source7_.cf77
            d_74___mcc_h1_ = source7_.cf78
            d_75_cf78_ = d_74___mcc_h1_
            d_76_cf77_ = d_73___mcc_h0_
            if d_75_cf78_:
                return D1_DC4(_dafny.SeqWithoutIsStrInference([d_76_cf77_, len(_dafny.SeqWithoutIsStrInference([False]))]))
            elif True:
                return D1_DC4(_dafny.SeqWithoutIsStrInference([d_76_cf77_, d_76_cf77_, d_76_cf77_]))
        elif source7_.is_DC40:
            d_77___mcc_h2_ = source7_.cf76
            d_78_cf76_ = d_77___mcc_h2_
            return D1_DC4(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_79_i0_ in range(default__.abs(458))])), (0) - ((_dafny.MultiSet([False, True, True, False, False])).cardinality)]))
        elif True:
            d_80___mcc_h3_ = source7_.cf79
            d_81_cf79_ = d_80___mcc_h3_
            return D1_DC4(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({128, len(_dafny.SeqWithoutIsStrInference([False])), 321})), 658, 195, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vhqc"))), 888]))

    @staticmethod
    def fm34(globalState):
        def iife6_():
            def iife8_():
                coll8_ = _dafny.Map()
                compr_8_: int
                for compr_8_ in _dafny.IntegerRange(246, 118):
                    d_82_v1_: int = compr_8_
                    if ((246) <= (d_82_v1_)) and ((d_82_v1_) < (118)):
                        coll8_[(d_82_v1_) * (len(_dafny.Map({False: True})))] = not(True)
                return _dafny.Map(coll8_)
            coll6_ = _dafny.Map()
            def iife7_():
                coll7_ = _dafny.Map()
                compr_7_: int
                for compr_7_ in _dafny.IntegerRange(246, 118):
                    d_82_v1_: int = compr_7_
                    if ((246) <= (d_82_v1_)) and ((d_82_v1_) < (118)):
                        coll7_[(d_82_v1_) * (len(_dafny.Map({False: True})))] = not(True)
                return _dafny.Map(coll7_)
            compr_6_: _dafny.Map
            for compr_6_ in (_dafny.SeqWithoutIsStrInference([_dafny.Map({394: True}), iife7_()
            , _dafny.Map({70: not(False)})])).Elements:
                d_83_v0_: _dafny.Map = compr_6_
                if (d_83_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.Map({394: True}), iife8_()
                , _dafny.Map({70: not(False)})])):
                    coll6_[d_83_v0_] = 999
            return _dafny.Map(coll6_)
        def iife9_():
            coll9_ = _dafny.Map()
            compr_9_: int
            for compr_9_ in _dafny.IntegerRange(133, 439):
                d_85_v2_: int = compr_9_
                if ((133) <= (d_85_v2_)) and ((d_85_v2_) < (439)):
                    coll9_[(d_85_v2_) - (214)] = len(_dafny.Set({False}))
            return _dafny.Map(coll9_)
        return _dafny.Map({((D13_DC31(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False]), _dafny.SeqWithoutIsStrInference([True, False])])): True}))).cf55) not in (iife6_()
        ): (0) - (default__.safeModuloInt(len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([720 for d_84_i0_ in range(default__.abs(-291))])): len(_dafny.SeqWithoutIsStrInference([False]))})), len(_dafny.Map({not(False): len(_dafny.Map({len(_dafny.Set({-82, len(iife9_()
        )})): not(True)}))}))))})

    @staticmethod
    def fm35(p0, globalState):
        def iife10_():
            coll10_ = _dafny.Map()
            compr_10_: int
            for compr_10_ in _dafny.IntegerRange(735, -829):
                d_86_v0_: int = compr_10_
                if ((735) <= (d_86_v0_)) and ((d_86_v0_) < (-829)):
                    coll10_[default__.safeModuloInt(d_86_v0_, 520)] = False
            return _dafny.Map(coll10_)
        def iife11_():
            coll11_ = _dafny.Map()
            compr_11_: int
            for compr_11_ in _dafny.IntegerRange(120, -129):
                d_88_v1_: int = compr_11_
                if ((120) <= (d_88_v1_)) and ((d_88_v1_) < (-129)):
                    coll11_[(d_88_v1_) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True]))])))] = False
            return _dafny.Map(coll11_)
        return ((iife10_()
        ) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True]))])): False}))) | ((_dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_87_i0_ in range(default__.abs(679))]))): True})) | (iife11_()
        ))

    @staticmethod
    def fm36(p0, p1, p2, p3, globalState):
        def iife12_():
            coll12_ = _dafny.Map()
            compr_12_: str
            for compr_12_ in (_dafny.Set({_dafny.CodePoint('u'), _dafny.CodePoint('g'), _dafny.CodePoint('k')})).Elements:
                d_89_v0_: str = compr_12_
                if (d_89_v0_) in (_dafny.Set({_dafny.CodePoint('u'), _dafny.CodePoint('g'), _dafny.CodePoint('k')})):
                    coll12_[d_89_v0_] = (not(True)) in (_dafny.Set({True, True}))
            return _dafny.Map(coll12_)
        return iife12_()
        

    @staticmethod
    def fm37(p0, p1, p2, globalState):
        return D12_DC28(_dafny.Map({False: -806}))

    @staticmethod
    def fm38(p0, globalState):
        return ((_dafny.Map({False: _dafny.SeqWithoutIsStrInference([False])})) | (_dafny.Map({True: _dafny.SeqWithoutIsStrInference([False])}))) | (_dafny.Map({True: _dafny.SeqWithoutIsStrInference([False])}))

    @staticmethod
    def fm39(p0, globalState):
        return (D16_DC39(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ok"))), _dafny.Map({len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_90_i0_ in range(default__.abs(57))])), (_dafny.MultiSet([True])).cardinality, -821])): 181}))).cf75

    @staticmethod
    def fm40(globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([-31, 889])])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([not(True)]))]), _dafny.SeqWithoutIsStrInference([len(_dafny.Map({_dafny.Map({False: 677}): False}))])]))

    @staticmethod
    def m7(p0, globalState):
        d_91_v0_: _dafny.Seq
        d_91_v0_ = _dafny.SeqWithoutIsStrInference([p0])
        d_92_v1_: C3
        nw0_ = C3()
        nw0_.ctor__(823, d_91_v0_, (0) - (p0))
        d_92_v1_ = nw0_
        d_93_v2_: bool
        d_93_v2_ = False
        d_94_v3_: _dafny.Map
        d_94_v3_ = _dafny.Map({d_93_v2_: (d_92_v1_).f9})
        d_95_v4_: _dafny.Map
        d_95_v4_ = _dafny.Map({False: d_94_v3_})
        d_96_v5_: _dafny.Seq
        d_96_v5_ = _dafny.SeqWithoutIsStrInference([d_93_v2_, d_93_v2_])
        d_97_v6_: _dafny.Seq
        d_97_v6_ = _dafny.SeqWithoutIsStrInference([not((d_96_v5_)[default__.safeIndex(p0, len(d_96_v5_))])])
        d_98_v7_: D13
        d_98_v7_ = D13_DC33(d_93_v2_, len(d_95_v4_), len(d_96_v5_), d_93_v2_, d_92_v1_)
        d_99_v8_: _dafny.Array
        nw1_ = _dafny.Array(None, 10)
        nw1_[int(0)] = d_92_v1_
        nw1_[int(1)] = d_92_v1_
        nw1_[int(2)] = (d_92_v1_ if d_93_v2_ else d_92_v1_)
        nw1_[int(3)] = (d_98_v7_).cf65
        nw1_[int(4)] = d_92_v1_
        nw1_[int(5)] = d_92_v1_
        nw1_[int(6)] = d_92_v1_
        nw1_[int(7)] = d_92_v1_
        nw1_[int(8)] = d_92_v1_
        nw1_[int(9)] = d_92_v1_
        d_99_v8_ = nw1_
        index0_ = default__.safeIndex(551, (d_99_v8_).length(0))
        (d_99_v8_)[index0_] = d_92_v1_
        index1_ = default__.safeIndex(551, (d_99_v8_).length(0))
        (d_99_v8_)[index1_] = d_92_v1_
        d_97_v6_ = d_97_v6_
        d_100_v9_: _dafny.Array
        nw2_ = _dafny.Array(None, 26)
        d_100_v9_ = nw2_
        d_101_v10_: C0
        nw3_ = C0()
        nw3_.ctor__(not(not(d_93_v2_)))
        d_101_v10_ = nw3_
        index2_ = default__.safeIndex(796, (d_100_v9_).length(0))
        (d_100_v9_)[index2_] = d_101_v10_
        index3_ = default__.safeIndex(796, (d_100_v9_).length(0))
        nw4_ = C0()
        nw4_.ctor__((d_101_v10_).f13)
        (d_100_v9_)[index3_] = nw4_
        d_102_v11_: C0
        nw5_ = C0()
        nw5_.ctor__((d_101_v10_).f13)
        d_102_v11_ = nw5_
        d_103_v12_: _dafny.Array
        def lambda0_(d_104_v10_):
            def lambda1_(d_105_i0_):
                return (d_104_v10_).f13

            return lambda1_

        init0_ = lambda0_(d_101_v10_)
        nw6_ = _dafny.Array(None, 5)
        for i0_0_ in range(nw6_.length(0)):
            nw6_[i0_0_] = init0_(i0_0_)
        d_103_v12_ = nw6_
        index4_ = default__.safeIndex(177, (d_103_v12_).length(0))
        (d_103_v12_)[index4_] = d_93_v2_
        d_106_v13_: D5
        d_106_v13_ = D5_DC12(p0)
        pat_let_tv0_ = d_92_v1_
        index5_ = default__.safeIndex(177, (d_103_v12_).length(0))
        def iife13_(_pat_let0_0):
            def iife14_(d_107_dt__update__tmp_h0_):
                def iife15_(_pat_let1_0):
                    def iife16_(d_108_dt__update_hcf20_h0_):
                        return D5_DC12(d_108_dt__update_hcf20_h0_)
                    return iife16_(_pat_let1_0)
                return iife15_((pat_let_tv0_).f9)
            return iife14_(_pat_let0_0)
        rhs0_ = (d_102_v11_).f13
        rhs1_ = not (d_93_v2_) or (d_93_v2_)
        rhs2_ = iife13_(d_106_v13_)
        lhs0_ = d_103_v12_
        lhs1_ = default__.safeIndex(177, (d_103_v12_).length(0))
        lhs2_ = globalState
        lhs0_[lhs1_] = rhs0_
        lhs2_.f5 = rhs1_
        d_106_v13_ = rhs2_
        hi0_ = (d_92_v1_).f9
        for d_109_i1_ in range(p0, hi0_):
            d_110_v14_: _dafny.Seq
            d_110_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ode"))
            d_111_v15_: _dafny.Map
            d_111_v15_ = _dafny.Map({(d_92_v1_).f9: (d_110_v14_ if (d_103_v12_)[default__.safeIndex(177, (d_103_v12_).length(0))] else d_110_v14_)})
            d_111_v15_ = (d_111_v15_).set(default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([(d_92_v1_).f9 for d_112_i2_ in range(default__.abs(479))])), (d_92_v1_).f9), d_110_v14_)
            index6_ = default__.safeIndex(177, (d_103_v12_).length(0))
            (d_103_v12_)[index6_] = (d_96_v5_)[default__.safeIndex(d_109_i1_, len(d_96_v5_))]
            if not(d_93_v2_):
                d_113_v16_: _dafny.Map
                d_113_v16_ = _dafny.Map({d_103_v12_: (d_92_v1_).fm16(_dafny.MultiSet([(d_101_v10_).f13]), globalState)})
                d_113_v16_ = d_113_v16_
                d_93_v2_ = not(d_93_v2_)
                index7_ = default__.safeIndex(177, (d_103_v12_).length(0))
                (d_103_v12_)[index7_] = (d_101_v10_).f13
                d_114_v17_: C1
                nw7_ = C1()
                nw7_.ctor__()
                d_114_v17_ = nw7_
                d_115_v18_: _dafny.MultiSet
                d_115_v18_ = _dafny.MultiSet([442, (0) - ((0) - ((d_92_v1_).f9)), (0) - (p0)])
                d_116_v19_: _dafny.Seq
                d_116_v19_ = _dafny.SeqWithoutIsStrInference([d_115_v18_, _dafny.MultiSet((d_92_v1_).f10)])
                d_117_v20_: _dafny.Set
                d_117_v20_ = _dafny.Set({(d_102_v11_).f13, (d_102_v11_).f13})
                d_118_v21_: _dafny.Map
                d_118_v21_ = _dafny.Map({((d_116_v19_)[default__.safeIndex(len(d_117_v20_), len(d_116_v19_))]).ispropersubset(d_115_v18_): (d_101_v10_).f13})
                d_119_v22_: _dafny.Map
                d_119_v22_ = d_118_v21_
                d_118_v21_ = (d_118_v21_) | (((d_119_v22_)))
            elif True:
                d_120_v23_: _dafny.Array
                nw8_ = _dafny.Array(int(0), 10)
                d_120_v23_ = nw8_
                index8_ = default__.safeIndex(532, (d_120_v23_).length(0))
                (d_120_v23_)[index8_] = p0
                index9_ = default__.safeIndex(532, (d_120_v23_).length(0))
                (d_120_v23_)[index9_] = d_109_i1_
                index10_ = default__.safeIndex(177, (d_103_v12_).length(0))
                (d_103_v12_)[index10_] = (d_101_v10_).f13
                d_121_v24_: str
                d_121_v24_ = _dafny.CodePoint('w')
                d_122_v25_: D0
                d_122_v25_ = D0_DC2((0) - ((0) - (d_109_i1_)), d_121_v24_, (0) - (len((d_97_v6_).set(default__.safeIndex((d_120_v23_)[default__.safeIndex(532, (d_120_v23_).length(0))], len(d_97_v6_)), (d_101_v10_).f13))), d_109_i1_, (d_92_v1_).f9)
                d_121_v24_ = (d_122_v25_).cf6
                d_123_v26_: _dafny.Set
                d_123_v26_ = _dafny.Set({not ((d_101_v10_).f13) or ((d_101_v10_).f13), (d_110_v14_) <= (d_110_v14_), (d_102_v11_).f13})
                index11_ = default__.safeIndex(177, (d_103_v12_).length(0))
                rhs3_ = d_103_v12_
                rhs4_ = (d_102_v11_).f13
                rhs5_ = (d_123_v26_) - ((d_123_v26_).intersection(d_123_v26_))
                lhs3_ = d_103_v12_
                lhs4_ = default__.safeIndex(177, (d_103_v12_).length(0))
                d_103_v12_ = rhs3_
                lhs3_[lhs4_] = rhs4_
                d_123_v26_ = rhs5_
                index12_ = default__.safeIndex(532, (d_120_v23_).length(0))
                (d_120_v23_)[index12_] = (d_92_v1_).f9
            d_124_v27_: _dafny.MultiSet
            d_124_v27_ = _dafny.MultiSet([d_103_v12_, d_103_v12_, d_103_v12_, d_103_v12_, d_103_v12_])
            (globalState).f1 = not(((d_124_v27_) | (d_124_v27_)).issubset(d_124_v27_))

    @staticmethod
    def Main(noArgsParameter__):
        d_125_v0_: _dafny.Seq
        d_125_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oeg"))
        d_126_globalState_: GlobalState
        nw9_ = GlobalState()
        nw9_.ctor__(-971, False, 560, True, d_125_v0_, False)
        d_126_globalState_ = nw9_
        d_127_v1_: bool
        d_127_v1_ = False
        d_128_v2_: _dafny.Set
        d_128_v2_ = _dafny.Set({d_127_v1_})
        d_129_v3_: _dafny.Map
        d_129_v3_ = _dafny.Map({277: d_128_v2_})
        d_130_v5_: int
        d_130_v5_ = 371
        d_131_v7_: str
        d_131_v7_ = _dafny.CodePoint('g')
        d_132_v8_: _dafny.MultiSet
        d_132_v8_ = _dafny.MultiSet([d_130_v5_])
        d_133_v9_: D0
        d_133_v9_ = D0_DC2(d_130_v5_, d_131_v7_, d_130_v5_, d_130_v5_, (d_132_v8_).cardinality)
        pat_let_tv1_ = d_130_v5_
        d_134_v10_: _dafny.Array
        nw10_ = _dafny.Array(None, 17)
        nw10_[int(0)] = (d_129_v3_ if True else d_129_v3_)
        nw10_[int(1)] = d_129_v3_
        nw10_[int(2)] = d_129_v3_
        nw10_[int(3)] = d_129_v3_
        nw10_[int(4)] = d_129_v3_
        nw10_[int(5)] = (d_129_v3_) | (d_129_v3_)
        nw10_[int(6)] = d_129_v3_
        nw10_[int(7)] = d_129_v3_
        nw10_[int(8)] = d_129_v3_
        nw10_[int(9)] = d_129_v3_
        def iife17_():
            coll13_ = _dafny.Map()
            compr_13_: int
            for compr_13_ in (_dafny.MultiSet([d_130_v5_])).Elements:
                d_135_v4_: int = compr_13_
                if (d_135_v4_) in (_dafny.MultiSet([d_130_v5_])):
                    coll13_[(d_135_v4_) + (d_130_v5_)] = d_128_v2_
            return _dafny.Map(coll13_)
        nw10_[int(10)] = iife17_()
        
        nw10_[int(11)] = _dafny.Map({d_130_v5_: d_128_v2_})
        nw10_[int(12)] = _dafny.Map({d_130_v5_: _dafny.Set({d_127_v1_, d_127_v1_})})
        nw10_[int(13)] = (d_129_v3_) | (d_129_v3_)
        def iife18_():
            coll14_ = _dafny.Map()
            compr_14_: int
            for compr_14_ in _dafny.IntegerRange(181, 229):
                d_136_v6_: int = compr_14_
                if ((181) <= (d_136_v6_)) and ((d_136_v6_) < (229)):
                    coll14_[default__.safeDivisionInt(d_136_v6_, d_130_v5_)] = d_128_v2_
            return _dafny.Map(coll14_)
        nw10_[int(14)] = (D0_DC0(iife18_()
)).cf0
        nw10_[int(15)] = d_129_v3_
        def iife19_(_pat_let2_0):
            def iife20_(d_137_dt__update__tmp_h0_):
                def iife21_(_pat_let3_0):
                    def iife22_(d_138_dt__update_hcf9_h0_):
                        return D0_DC2((d_137_dt__update__tmp_h0_).cf5, (d_137_dt__update__tmp_h0_).cf6, (d_137_dt__update__tmp_h0_).cf7, (d_137_dt__update__tmp_h0_).cf8, d_138_dt__update_hcf9_h0_)
                    return iife22_(_pat_let3_0)
                return iife21_(pat_let_tv1_)
            return iife20_(_pat_let2_0)
        nw10_[int(16)] = default__.fm0(d_127_v1_, default__.fm1(D0_DC0(d_129_v3_), d_127_v1_, d_130_v5_, d_127_v1_, d_126_globalState_), iife19_(d_133_v9_), d_126_globalState_)
        d_134_v10_ = nw10_
        index13_ = default__.safeIndex(207, (d_134_v10_).length(0))
        (d_134_v10_)[index13_] = d_129_v3_
        d_139_v11_: _dafny.Map
        d_139_v11_ = _dafny.Map({d_130_v5_: d_129_v3_})
        d_140_v12_: D0
        d_140_v12_ = D0_DC0(d_129_v3_)
        index14_ = default__.safeIndex(207, (d_134_v10_).length(0))
        (d_134_v10_)[index14_] = ((d_139_v11_)[d_130_v5_] if (d_130_v5_) in (d_139_v11_) else default__.fm0(default__.fm1(d_140_v12_, d_127_v1_, (0) - (len(d_128_v2_)), d_127_v1_, d_126_globalState_), d_127_v1_, d_133_v9_, d_126_globalState_))
        if d_127_v1_:
            d_141_v13_: _dafny.Map
            d_141_v13_ = _dafny.Map({default__.safeDivisionInt(d_130_v5_, d_130_v5_): d_130_v5_})
            d_141_v13_ = d_141_v13_
            d_142_v14_: _dafny.Array
            nw11_ = _dafny.Array(False, 26)
            d_142_v14_ = nw11_
            index15_ = default__.safeIndex(28, (d_142_v14_).length(0))
            (d_142_v14_)[index15_] = d_127_v1_
            index16_ = default__.safeIndex(28, (d_142_v14_).length(0))
            (d_142_v14_)[index16_] = d_127_v1_
            d_143_v15_: _dafny.Map
            d_143_v15_ = _dafny.Map({(d_142_v14_)[default__.safeIndex(28, (d_142_v14_).length(0))]: d_130_v5_})
            d_144_v16_: _dafny.Map
            d_144_v16_ = _dafny.Map({d_125_v0_: len(d_143_v15_)})
            d_145_v17_: _dafny.Map
            d_145_v17_ = _dafny.Map({d_130_v5_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e"))})
            d_146_v19_: _dafny.Seq
            def iife23_():
                coll15_ = _dafny.Set()
                compr_15_: int
                for compr_15_ in ((d_145_v17_).set(d_130_v5_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k")))).keys.Elements:
                    d_147_v18_: int = compr_15_
                    if (d_147_v18_) in ((d_145_v17_).set(d_130_v5_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k")))):
                        coll15_ = coll15_.union(_dafny.Set([default__.safeDivisionInt(d_147_v18_, (_dafny.MultiSet([len(_dafny.Map({True: len(_dafny.Map({511: 38}))})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "edejin")))])).cardinality)]))
                return _dafny.Set(coll15_)
            d_146_v19_ = _dafny.SeqWithoutIsStrInference([len(iife23_()
            )])
            d_144_v16_ = (d_144_v16_).set((d_125_v0_) + (d_125_v0_), len(d_146_v19_))
            d_125_v0_ = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_148_i0_ in range(default__.abs(-612))])) + (default__.fm2(d_125_v0_, d_126_globalState_))
            d_149_v20_: D0
            d_149_v20_ = D0_DC1(len(d_125_v0_), (d_142_v14_)[default__.safeIndex(28, (d_142_v14_).length(0))], d_130_v5_, (d_142_v14_)[default__.safeIndex(28, (d_142_v14_).length(0))])
            d_150_v21_: _dafny.Seq
            d_150_v21_ = _dafny.SeqWithoutIsStrInference([d_127_v1_])
            d_151_v22_: _dafny.MultiSet
            d_151_v22_ = _dafny.MultiSet([(d_142_v14_)[default__.safeIndex(28, (d_142_v14_).length(0))], d_127_v1_, False])
            nw12_ = _dafny.Array(None, 15)
            nw12_[int(0)] = not((d_149_v20_).cf2)
            nw12_[int(1)] = (d_142_v14_)[default__.safeIndex(28, (d_142_v14_).length(0))]
            nw12_[int(2)] = (d_142_v14_)[default__.safeIndex(28, (d_142_v14_).length(0))]
            nw12_[int(3)] = default__.fm1(d_140_v12_, d_127_v1_, d_130_v5_, (d_142_v14_)[default__.safeIndex(28, (d_142_v14_).length(0))], d_126_globalState_)
            nw12_[int(4)] = d_127_v1_
            nw12_[int(5)] = (d_130_v5_) < (d_130_v5_)
            nw12_[int(6)] = not ((d_142_v14_)[default__.safeIndex(28, (d_142_v14_).length(0))]) or ((d_142_v14_)[default__.safeIndex(28, (d_142_v14_).length(0))])
            nw12_[int(7)] = d_127_v1_
            nw12_[int(8)] = (d_142_v14_)[default__.safeIndex(28, (d_142_v14_).length(0))]
            nw12_[int(9)] = (d_150_v21_) <= (d_150_v21_)
            nw12_[int(10)] = (d_142_v14_)[default__.safeIndex(28, (d_142_v14_).length(0))]
            nw12_[int(11)] = d_127_v1_
            nw12_[int(12)] = d_127_v1_
            nw12_[int(13)] = (d_142_v14_)[default__.safeIndex(28, (d_142_v14_).length(0))]
            nw12_[int(14)] = (d_151_v22_).issubset(d_151_v22_)
            d_142_v14_ = nw12_
        elif True:
            (d_126_globalState_).f5 = default__.fm1(D0_DC0(_dafny.Map({len(d_125_v0_): d_128_v2_})), d_127_v1_, (d_130_v5_ if d_127_v1_ else d_130_v5_), d_127_v1_, d_126_globalState_)
            if d_127_v1_:
                d_152_v23_: C0
                nw13_ = C0()
                nw13_.ctor__((d_127_v1_ if True else d_127_v1_))
                d_152_v23_ = nw13_
                d_153_v24_: _dafny.Map
                d_153_v24_ = _dafny.Map({d_130_v5_: d_130_v5_})
                rhs6_ = (d_130_v5_) - (((d_153_v24_)[d_130_v5_] if (d_130_v5_) in (d_153_v24_) else len(d_125_v0_)))
                rhs7_ = (d_152_v23_).f13
                rhs8_ = d_130_v5_
                lhs5_ = d_126_globalState_
                lhs6_ = d_126_globalState_
                lhs7_ = d_126_globalState_
                lhs5_.f2 = rhs6_
                lhs6_.f1 = rhs7_
                lhs7_.f2 = rhs8_
                pat_let_tv2_ = d_130_v5_
                pat_let_tv3_ = d_130_v5_
                pat_let_tv4_ = d_128_v2_
                pat_let_tv5_ = d_130_v5_
                pat_let_tv6_ = d_130_v5_
                pat_let_tv7_ = d_128_v2_
                d_154_v27_: D13
                def iife24_(_pat_let4_0):
                    def iife25_(d_155_dt__update__tmp_h1_):
                        def iife27_():
                            def iife29_():
                                coll18_ = _dafny.Set()
                                compr_18_: int
                                for compr_18_ in _dafny.IntegerRange(598, 884):
                                    d_158_v26_: int = compr_18_
                                    if ((598) <= (d_158_v26_)) and ((d_158_v26_) < (884)):
                                        coll18_ = coll18_.union(_dafny.Set([(d_158_v26_) + (pat_let_tv3_)]))
                                return _dafny.Set(coll18_)
                            coll16_ = _dafny.Map()
                            def iife28_():
                                coll17_ = _dafny.Set()
                                compr_17_: int
                                for compr_17_ in _dafny.IntegerRange(598, 884):
                                    d_156_v26_: int = compr_17_
                                    if ((598) <= (d_156_v26_)) and ((d_156_v26_) < (884)):
                                        coll17_ = coll17_.union(_dafny.Set([(d_156_v26_) + (pat_let_tv2_)]))
                                return _dafny.Set(coll17_)
                            compr_16_: int
                            for compr_16_ in (iife28_()
                            ).Elements:
                                d_157_v25_: int = compr_16_
                                if (d_157_v25_) in (iife29_()
                                ):
                                    coll16_[default__.safeModuloInt(d_157_v25_, pat_let_tv5_)] = pat_let_tv4_
                            return _dafny.Map(coll16_)
                        def iife26_(_pat_let5_0):
                            def iife30_(d_159_dt__update_hcf0_h0_):
                                return D0_DC0(d_159_dt__update_hcf0_h0_)
                            return iife30_(_pat_let5_0)
                        return iife26_((iife27_()
                        ).set(pat_let_tv6_, pat_let_tv7_))
                    return iife25_(_pat_let4_0)
                d_154_v27_ = D13_DC32(d_130_v5_, True, default__.fm1(iife24_(d_140_v12_), d_127_v1_, d_130_v5_, d_127_v1_, d_126_globalState_), (d_152_v23_).f13, d_127_v1_)
                pat_let_tv8_ = d_127_v1_
                pat_let_tv9_ = d_152_v23_
                pat_let_tv10_ = d_152_v23_
                pat_let_tv11_ = d_127_v1_
                d_160_v28_: _dafny.Map
                def iife32_(_pat_let7_0):
                    def iife33_(d_161_dt__update__tmp_h2_):
                        def iife34_(_pat_let8_0):
                            def iife35_(d_162_dt__update_hcf59_h0_):
                                def iife36_(_pat_let9_0):
                                    def iife37_(d_163_dt__update_hcf58_h0_):
                                        def iife38_(_pat_let10_0):
                                            def iife39_(d_164_dt__update_hcf57_h0_):
                                                return D13_DC32((d_161_dt__update__tmp_h2_).cf56, d_164_dt__update_hcf57_h0_, d_163_dt__update_hcf58_h0_, d_162_dt__update_hcf59_h0_, (d_161_dt__update__tmp_h2_).cf60)
                                            return iife39_(_pat_let10_0)
                                        return iife38_((pat_let_tv10_).f13)
                                    return iife37_(_pat_let9_0)
                                return iife36_((pat_let_tv9_).f13)
                            return iife35_(_pat_let8_0)
                        return iife34_(pat_let_tv8_)
                    return iife33_(_pat_let7_0)
                def iife31_(_pat_let6_0):
                    def iife40_(d_165_dt__update__tmp_h3_):
                        def iife41_(_pat_let11_0):
                            def iife42_(d_166_dt__update_hcf58_h1_):
                                return D13_DC32((d_165_dt__update__tmp_h3_).cf56, (d_165_dt__update__tmp_h3_).cf57, d_166_dt__update_hcf58_h1_, (d_165_dt__update__tmp_h3_).cf59, (d_165_dt__update__tmp_h3_).cf60)
                            return iife42_(_pat_let11_0)
                        return iife41_(pat_let_tv11_)
                    return iife40_(_pat_let6_0)
                d_160_v28_ = _dafny.Map({(iife31_(iife32_(d_154_v27_))).cf58: len(_dafny.Set({(d_152_v23_).f13}))})
                d_160_v28_ = (d_160_v28_) | ((d_160_v28_) | (_dafny.Map({not(d_127_v1_): d_130_v5_})))
                d_125_v0_ = (default__.fm2(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ioebdfr")), d_126_globalState_)) + (d_125_v0_)
                d_130_v5_ = (d_133_v9_).cf5
            elif True:
                default__.m7(d_130_v5_, d_126_globalState_)
                d_167_v29_: _dafny.Array
                nw14_ = _dafny.Array(_dafny.MultiSet({}), 17)
                d_167_v29_ = nw14_
                d_167_v29_ = (d_167_v29_ if (not(d_127_v1_) if not(d_127_v1_) else d_127_v1_) else d_167_v29_)
                d_168_v30_: _dafny.Seq
                d_168_v30_ = _dafny.SeqWithoutIsStrInference([d_127_v1_, d_127_v1_])
                d_169_v31_: _dafny.Seq
                d_169_v31_ = _dafny.SeqWithoutIsStrInference([d_130_v5_, d_130_v5_])
                d_170_v32_: C4
                nw15_ = C4()
                nw15_.ctor__(len(d_169_v31_))
                d_170_v32_ = nw15_
                d_171_v33_: _dafny.Seq
                d_171_v33_ = _dafny.SeqWithoutIsStrInference([d_170_v32_, d_170_v32_])
                d_172_v34_: _dafny.Array
                nw16_ = _dafny.Array(None, 24)
                nw16_[int(0)] = d_130_v5_
                nw16_[int(1)] = d_130_v5_
                nw16_[int(2)] = d_130_v5_
                nw16_[int(3)] = d_130_v5_
                nw16_[int(4)] = default__.safeDivisionInt(len(d_168_v30_), d_130_v5_)
                nw16_[int(5)] = default__.safeModuloInt(806, d_130_v5_)
                nw16_[int(6)] = d_130_v5_
                nw16_[int(7)] = (d_130_v5_) * (442)
                nw16_[int(8)] = (694) * (len(d_171_v33_))
                nw16_[int(9)] = d_130_v5_
                nw16_[int(10)] = d_130_v5_
                nw16_[int(11)] = 660
                nw16_[int(12)] = d_130_v5_
                nw16_[int(13)] = (d_130_v5_) * (len(d_168_v30_))
                nw16_[int(14)] = d_130_v5_
                nw16_[int(15)] = d_130_v5_
                nw16_[int(16)] = d_130_v5_
                nw16_[int(17)] = default__.fm24(d_130_v5_, d_126_globalState_)
                nw16_[int(18)] = (0) - (d_130_v5_)
                nw16_[int(19)] = (d_130_v5_) * (d_130_v5_)
                nw16_[int(20)] = d_130_v5_
                nw16_[int(21)] = len(_dafny.SeqWithoutIsStrInference([d_127_v1_, d_127_v1_]))
                nw16_[int(22)] = (0) - (-329)
                nw16_[int(23)] = d_130_v5_
                d_172_v34_ = nw16_
                index17_ = default__.safeIndex(597, (d_172_v34_).length(0))
                (d_172_v34_)[index17_] = default__.safeDivisionInt(d_130_v5_, d_130_v5_)
                d_173_v35_: _dafny.MultiSet
                d_173_v35_ = _dafny.MultiSet([d_172_v34_, d_172_v34_])
                d_174_v36_: _dafny.Map
                d_174_v36_ = _dafny.Map({d_127_v1_: d_130_v5_})
                d_175_v37_: _dafny.MultiSet
                d_175_v37_ = _dafny.MultiSet([(d_130_v5_) != (((d_173_v35_)[d_172_v34_] if (d_172_v34_) in (d_173_v35_) else (0) - (len(d_174_v36_)))), default__.fm1(d_140_v12_, d_127_v1_, (0) - (d_130_v5_), d_127_v1_, d_126_globalState_), False])
                index18_ = default__.safeIndex(597, (d_172_v34_).length(0))
                (d_172_v34_)[index18_] = ((d_175_v37_)[(d_127_v1_) or (d_127_v1_)] if ((d_127_v1_) or (d_127_v1_)) in (d_175_v37_) else 83)
                (d_126_globalState_).f2 = (d_130_v5_ if (d_128_v2_).issubset(d_128_v2_) else d_130_v5_)
                (d_126_globalState_).f0 = ((d_172_v34_)[default__.safeIndex(597, (d_172_v34_).length(0))] if not((len(d_128_v2_)) == ((d_172_v34_)[default__.safeIndex(597, (d_172_v34_).length(0))])) else default__.fm15(d_130_v5_, d_127_v1_, d_126_globalState_))
            d_176_v38_: _dafny.MultiSet
            d_176_v38_ = _dafny.MultiSet([d_127_v1_, d_127_v1_])
            (d_126_globalState_).f2 = (d_130_v5_) + ((d_176_v38_).cardinality)
            (d_126_globalState_).f0 = ((d_130_v5_) + ((0) - (d_130_v5_))) * (d_130_v5_)
            if d_127_v1_:
                default__.m7(d_130_v5_, d_126_globalState_)
                d_177_v39_: _dafny.Array
                def lambda2_(d_178_v1_):
                    def lambda3_(d_179_i1_):
                        return d_178_v1_

                    return lambda3_

                init1_ = lambda2_(d_127_v1_)
                nw17_ = _dafny.Array(None, 25)
                for i0_1_ in range(nw17_.length(0)):
                    nw17_[i0_1_] = init1_(i0_1_)
                d_177_v39_ = nw17_
                index19_ = default__.safeIndex(957, (d_177_v39_).length(0))
                (d_177_v39_)[index19_] = not(d_127_v1_)
                d_180_v40_: _dafny.Seq
                d_180_v40_ = _dafny.SeqWithoutIsStrInference([d_130_v5_, d_130_v5_, d_130_v5_])
                d_181_v41_: _dafny.Seq
                d_181_v41_ = _dafny.SeqWithoutIsStrInference([not(True), d_127_v1_, default__.fm1(D0_DC0(d_129_v3_), d_127_v1_, (d_180_v40_)[default__.safeIndex(900, len(d_180_v40_))], d_127_v1_, d_126_globalState_), d_127_v1_, d_127_v1_])
                d_182_v42_: _dafny.Map
                d_182_v42_ = _dafny.Map({len(d_125_v0_): d_127_v1_})
                d_183_v43_: D14
                d_183_v43_ = D14_DC36(d_127_v1_, (0) - (d_130_v5_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ncd"))), _dafny.Map({-972: ((d_182_v42_)[d_130_v5_] if (d_130_v5_) in (d_182_v42_) else d_127_v1_)}))
                d_184_v44_: _dafny.Map
                d_184_v44_ = _dafny.Map({d_131_v7_: d_183_v43_})
                index20_ = default__.safeIndex(957, (d_177_v39_).length(0))
                rhs9_ = d_130_v5_
                rhs10_ = d_127_v1_
                rhs11_ = (d_181_v41_)[default__.safeIndex((0) - (len(d_184_v44_)), len(d_181_v41_))]
                lhs8_ = d_126_globalState_
                lhs9_ = d_177_v39_
                lhs10_ = default__.safeIndex(957, (d_177_v39_).length(0))
                lhs11_ = d_126_globalState_
                lhs8_.f2 = rhs9_
                lhs9_[lhs10_] = rhs10_
                lhs11_.f1 = rhs11_
                (d_126_globalState_).f5 = not((d_130_v5_) < (d_130_v5_))
                d_180_v40_ = d_180_v40_
                d_185_v45_: _dafny.Map
                d_185_v45_ = _dafny.Map({d_127_v1_: _dafny.SeqWithoutIsStrInference([False, (d_177_v39_)[default__.safeIndex(957, (d_177_v39_).length(0))], (d_177_v39_)[default__.safeIndex(957, (d_177_v39_).length(0))]])})
                (d_126_globalState_).f1 = not((d_185_v45_) == ((default__.fm38((d_177_v39_)[default__.safeIndex(957, (d_177_v39_).length(0))], d_126_globalState_)) | (d_185_v45_)))
            elif True:
                d_186_v46_: _dafny.Map
                d_186_v46_ = _dafny.Map({d_130_v5_: (d_127_v1_ if True else d_127_v1_)})
                d_187_v47_: _dafny.Map
                d_187_v47_ = _dafny.Map({d_131_v7_: d_130_v5_})
                d_186_v46_ = (d_186_v46_).set(d_130_v5_, (d_132_v8_).issubset(_dafny.MultiSet(default__.fm25(d_130_v5_, d_187_v47_, d_131_v7_, d_126_globalState_))))
                d_125_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dcti"))
                d_188_v48_: T0
                nw18_ = C6()
                nw18_.ctor__(d_130_v5_)
                d_188_v48_ = nw18_
                d_189_v49_: T1
                nw19_ = C1()
                nw19_.ctor__()
                d_189_v49_ = nw19_
                d_190_v50_: _dafny.MultiSet
                d_190_v50_ = _dafny.MultiSet([d_189_v49_, d_189_v49_, d_189_v49_, d_189_v49_])
                d_191_v51_: D11
                d_191_v51_ = D11_DC26(d_188_v48_, d_190_v50_, d_130_v5_, d_125_v0_)
                d_192_v52_: _dafny.Set
                d_192_v52_ = _dafny.Set({d_188_v48_})
                d_193_v53_: _dafny.Array
                nw20_ = _dafny.Array(None, 5)
                nw20_[int(0)] = d_127_v1_
                nw20_[int(1)] = (d_130_v5_) != (d_130_v5_)
                nw20_[int(2)] = (_dafny.Set({(d_191_v51_).cf45, d_188_v48_})).isdisjoint(d_192_v52_)
                nw20_[int(3)] = d_127_v1_
                nw20_[int(4)] = d_127_v1_
                d_193_v53_ = nw20_
                index21_ = default__.safeIndex(896, (d_193_v53_).length(0))
                (d_193_v53_)[index21_] = d_127_v1_
                d_194_v54_: C0
                nw21_ = C0()
                nw21_.ctor__(d_127_v1_)
                d_194_v54_ = nw21_
                d_195_v55_: _dafny.Map
                d_195_v55_ = _dafny.Map({516: d_194_v54_})
                d_196_v56_: _dafny.Set
                d_196_v56_ = _dafny.Set({len(_dafny.SeqWithoutIsStrInference([len(d_195_v55_), d_130_v5_]))})
                d_197_v57_: _dafny.Seq
                d_197_v57_ = _dafny.SeqWithoutIsStrInference([(d_188_v48_).f6, d_130_v5_])
                index22_ = default__.safeIndex(896, (d_193_v53_).length(0))
                (d_193_v53_)[index22_] = (default__.fm10(len(d_197_v57_), (d_194_v54_).f13, d_127_v1_, d_126_globalState_)).issubset(d_196_v56_)
                d_127_v1_ = d_127_v1_
                d_198_v58_: _dafny.Map
                d_198_v58_ = _dafny.Map({(d_194_v54_).f13: len(d_197_v57_)})
                (d_126_globalState_).f2 = ((0) - ((len(d_198_v58_)) * (826)) if d_127_v1_ else default__.safeModuloInt(d_130_v5_, 964))
        d_199_v59_: C0
        nw22_ = C0()
        nw22_.ctor__(d_127_v1_)
        d_199_v59_ = nw22_
        source8_ = d_199_v59_
        d_200___mcc_h0_ = source8_
        d_201_cf72_ = d_200___mcc_h0_
        d_202_v60_: _dafny.Array
        nw23_ = _dafny.Array(False, 6)
        d_202_v60_ = nw23_
        index23_ = default__.safeIndex(782, (d_202_v60_).length(0))
        (d_202_v60_)[index23_] = d_127_v1_
        d_203_v61_: _dafny.Set
        d_203_v61_ = _dafny.Set({d_130_v5_})
        d_204_v62_: _dafny.MultiSet
        d_204_v62_ = _dafny.MultiSet([d_203_v61_])
        index24_ = default__.safeIndex(782, (d_202_v60_).length(0))
        (d_202_v60_)[index24_] = (d_203_v61_) in (d_204_v62_)
        d_205_v63_: _dafny.Seq
        d_205_v63_ = _dafny.SeqWithoutIsStrInference([d_130_v5_])
        d_132_v8_ = _dafny.MultiSet([(0) - ((len(d_205_v63_) if d_127_v1_ else default__.fm15(d_130_v5_, d_127_v1_, d_126_globalState_))), (d_130_v5_) + (len(d_203_v61_)), d_130_v5_, d_130_v5_])
        (d_126_globalState_).f1 = not((d_199_v59_).f13)
        if ((d_199_v59_).f13 if (d_199_v59_).f13 else not ((d_201_cf72_).f13) or (d_127_v1_)):
            d_206_v64_: _dafny.Seq
            d_206_v64_ = _dafny.SeqWithoutIsStrInference([(d_201_cf72_).f13])
            d_206_v64_ = (d_206_v64_) + (_dafny.SeqWithoutIsStrInference([(d_199_v59_).f13]))
            d_207_v65_: C2
            nw24_ = C2()
            nw24_.ctor__(d_125_v0_, (d_201_cf72_).f13)
            d_207_v65_ = nw24_
            d_208_v66_: _dafny.Map
            d_208_v66_ = _dafny.Map({d_203_v61_: d_206_v64_})
            (d_207_v65_).f12 = (_dafny.Map({_dafny.Set({len(d_203_v61_)}): default__.fm21((d_207_v65_).f11, d_126_globalState_)})) != (d_208_v66_)
            d_209_v67_: _dafny.Array
            def lambda4_(d_210_v5_):
                def lambda5_(d_211_i2_):
                    return (d_211_i2_) * (d_210_v5_)

                return lambda5_

            init2_ = lambda4_(d_130_v5_)
            nw25_ = _dafny.Array(None, 16)
            for i0_2_ in range(nw25_.length(0)):
                nw25_[i0_2_] = init2_(i0_2_)
            d_209_v67_ = nw25_
            def lambda6_(d_212_v1_):
                def lambda7_(d_213_i3_):
                    return (d_213_i3_) * ((0) - ((_dafny.MultiSet([d_212_v1_])).cardinality))

                return lambda7_

            init3_ = lambda6_(d_127_v1_)
            nw26_ = _dafny.Array(None, 26)
            for i0_3_ in range(nw26_.length(0)):
                nw26_[i0_3_] = init3_(i0_3_)
            d_209_v67_ = nw26_
            d_214_v68_: _dafny.Map
            d_214_v68_ = _dafny.Map({d_131_v7_: default__.fm1(d_140_v12_, d_207_v65_.f12, (0) - (d_130_v5_), (d_201_cf72_).f13, d_126_globalState_)})
            d_215_v69_: _dafny.Seq
            d_215_v69_ = _dafny.SeqWithoutIsStrInference([d_214_v68_, d_214_v68_, (d_214_v68_) | (d_214_v68_)])
            d_215_v69_ = d_215_v69_
        elif True:
            default__.m7(d_130_v5_, d_126_globalState_)
            (d_126_globalState_).f5 = (d_199_v59_).f13
            d_125_v0_ = d_125_v0_
            (d_126_globalState_).f5 = True
            d_216_v70_: _dafny.Seq
            d_216_v70_ = _dafny.SeqWithoutIsStrInference([(d_199_v59_).f13, (d_199_v59_).f13])
            d_217_v71_: _dafny.Map
            d_217_v71_ = _dafny.Map({not(not((len(d_216_v70_)) == (d_130_v5_))): ((d_199_v59_).f13 if (d_201_cf72_).f13 else (d_201_cf72_).f13)})
            d_217_v71_ = (d_217_v71_).set(((d_217_v71_)[True] if (True) in (d_217_v71_) else d_127_v1_), (d_132_v8_).ispropersubset(d_132_v8_))
        d_218_v72_: _dafny.Map
        d_218_v72_ = _dafny.Map({d_127_v1_: d_130_v5_})
        d_219_v73_: _dafny.Map
        d_219_v73_ = _dafny.Map({d_130_v5_: (d_199_v59_).f13})
        d_220_v74_: _dafny.Set
        d_220_v74_ = _dafny.Set({default__.safeModuloInt(d_130_v5_, len((_dafny.SeqWithoutIsStrInference([d_218_v72_, d_218_v72_])).set(default__.safeIndex(d_130_v5_, len(_dafny.SeqWithoutIsStrInference([d_218_v72_, d_218_v72_]))), _dafny.Map({d_127_v1_: len(d_219_v73_)}))))})
        d_220_v74_ = d_220_v74_
        d_127_v1_ = False
        (d_126_globalState_).f0 = d_130_v5_
        rhs12_ = d_130_v5_
        rhs13_ = d_127_v1_
        lhs12_ = d_126_globalState_
        lhs13_ = d_126_globalState_
        lhs12_.f0 = rhs12_
        lhs13_.f5 = rhs13_
        d_221_v75_: _dafny.Map
        d_221_v75_ = _dafny.Map({len(default__.fm39(d_130_v5_, d_126_globalState_)): d_130_v5_})
        d_222_v76_: D16
        d_222_v76_ = D16_DC39((0) - (len(d_125_v0_)), d_221_v75_)
        pat_let_tv12_ = d_221_v75_
        def iife43_(_pat_let12_0):
            def iife44_(d_223_dt__update__tmp_h4_):
                def iife45_(_pat_let13_0):
                    def iife46_(d_224_dt__update_hcf75_h0_):
                        return D16_DC39((d_223_dt__update__tmp_h4_).cf74, d_224_dt__update_hcf75_h0_)
                    return iife46_(_pat_let13_0)
                return iife45_(pat_let_tv12_)
            return iife44_(_pat_let12_0)
        source9_ = iife43_(d_222_v76_)
        if source9_.is_DC39:
            d_225___mcc_h1_ = source9_.cf74
            d_226___mcc_h2_ = source9_.cf75
            d_227_cf75_ = d_226___mcc_h2_
            d_228_cf74_ = d_225___mcc_h1_
            d_229_v77_: D1
            d_229_v77_ = D1_DC5()
            source10_ = d_229_v77_
            if source10_.is_DC5:
                d_230_v78_: _dafny.MultiSet
                d_230_v78_ = _dafny.MultiSet([(d_199_v59_).f13])
                d_231_v79_: _dafny.Seq
                d_231_v79_ = _dafny.SeqWithoutIsStrInference([d_228_cf74_])
                d_232_v80_: C3
                nw27_ = C3()
                nw27_.ctor__((d_230_v78_).cardinality, d_231_v79_, d_228_cf74_)
                d_232_v80_ = nw27_
                d_233_v81_: D13
                d_233_v81_ = D13_DC33((d_199_v59_).f13, len(d_128_v2_), 11, (d_199_v59_).f13, d_232_v80_)
                d_234_v82_: _dafny.Array
                nw28_ = _dafny.Array(None, 6)
                nw28_[int(0)] = d_228_cf74_
                nw28_[int(1)] = len(d_221_v75_)
                nw28_[int(2)] = default__.safeDivisionInt(d_130_v5_, d_228_cf74_)
                nw28_[int(3)] = d_130_v5_
                nw28_[int(4)] = (len(d_220_v74_) if not(default__.fm1(d_140_v12_, (d_233_v81_).cf64, d_130_v5_, not((d_199_v59_).f13), d_126_globalState_)) else d_130_v5_)
                nw28_[int(5)] = d_130_v5_
                d_234_v82_ = nw28_
                index25_ = default__.safeIndex(649, (d_234_v82_).length(0))
                (d_234_v82_)[index25_] = (0) - ((0) - ((d_232_v80_).f9))
                index26_ = default__.safeIndex(649, (d_234_v82_).length(0))
                (d_234_v82_)[index26_] = (d_232_v80_).f9
                d_235_v83_: C5
                nw29_ = C5()
                nw29_.ctor__((d_199_v59_).f13, d_131_v7_)
                d_235_v83_ = nw29_
                d_235_v83_ = d_235_v83_
                d_236_v84_: _dafny.Array
                nw30_ = _dafny.Array(_dafny.Seq({}), 12)
                d_236_v84_ = nw30_
                d_237_v85_: _dafny.Array
                def lambda8_(d_238_v83_):
                    def lambda9_(d_239_i4_):
                        return (d_238_v83_).f7

                    return lambda9_

                init4_ = lambda8_(d_235_v83_)
                nw31_ = _dafny.Array(None, 26)
                for i0_4_ in range(nw31_.length(0)):
                    nw31_[i0_4_] = init4_(i0_4_)
                d_237_v85_ = nw31_
                d_240_v86_: _dafny.Seq
                d_240_v86_ = _dafny.SeqWithoutIsStrInference([d_237_v85_, d_237_v85_])
                index27_ = default__.safeIndex(360, (d_236_v84_).length(0))
                (d_236_v84_)[index27_] = d_240_v86_
                index28_ = default__.safeIndex(360, (d_236_v84_).length(0))
                (d_236_v84_)[index28_] = d_240_v86_
                (d_126_globalState_).f2 = (d_230_v78_).cardinality
            elif True:
                d_241___mcc_h4_ = source10_.cf11
                d_242_cf11_ = d_241___mcc_h4_
                d_243_v87_: _dafny.Array
                nw32_ = _dafny.Array(False, 26)
                d_243_v87_ = nw32_
                index29_ = default__.safeIndex(671, (d_243_v87_).length(0))
                (d_243_v87_)[index29_] = (d_199_v59_).f13
                index30_ = default__.safeIndex(671, (d_243_v87_).length(0))
                (d_243_v87_)[index30_] = d_127_v1_
                (d_126_globalState_).f1 = (d_199_v59_).f13
                d_244_v88_: _dafny.Array
                def lambda10_(d_245_v5_):
                    def lambda11_(d_246_i5_):
                        return (d_246_i5_) + (d_245_v5_)

                    return lambda11_

                init5_ = lambda10_(d_130_v5_)
                nw33_ = _dafny.Array(None, 2)
                for i0_5_ in range(nw33_.length(0)):
                    nw33_[i0_5_] = init5_(i0_5_)
                d_244_v88_ = nw33_
                d_244_v88_ = d_244_v88_
                default__.m7(d_228_cf74_, d_126_globalState_)
            if (d_199_v59_).f13:
                d_247_v89_: _dafny.Array
                nw34_ = _dafny.Array(False, 10)
                d_247_v89_ = nw34_
                d_247_v89_ = d_247_v89_
                d_248_v90_: _dafny.Map
                d_248_v90_ = _dafny.Map({d_247_v89_: d_127_v1_})
                (d_126_globalState_).f5 = (False if (d_247_v89_) in (d_248_v90_) else (d_199_v59_).f13)
                d_249_v91_: D0
                d_249_v91_ = D0_DC1(d_130_v5_, d_127_v1_, d_228_cf74_, (d_199_v59_).f13)
                d_250_v92_: _dafny.Map
                d_250_v92_ = _dafny.Map({default__.fm1(d_140_v12_, not((d_199_v59_).f13), len(d_125_v0_), default__.fm1(d_140_v12_, (d_199_v59_).f13, -926, (d_249_v91_).cf4, d_126_globalState_), d_126_globalState_): d_125_v0_})
                d_250_v92_ = (d_250_v92_).set((d_199_v59_).f13, d_125_v0_)
                d_251_v93_: _dafny.Array
                nw35_ = _dafny.Array(None, 19)
                nw35_[int(0)] = d_247_v89_
                nw35_[int(1)] = d_247_v89_
                nw35_[int(2)] = d_247_v89_
                nw35_[int(3)] = d_247_v89_
                nw35_[int(4)] = d_247_v89_
                nw35_[int(5)] = d_247_v89_
                nw35_[int(6)] = d_247_v89_
                nw35_[int(7)] = d_247_v89_
                nw35_[int(8)] = d_247_v89_
                nw35_[int(9)] = d_247_v89_
                nw35_[int(10)] = d_247_v89_
                nw35_[int(11)] = d_247_v89_
                nw35_[int(12)] = d_247_v89_
                nw35_[int(13)] = d_247_v89_
                nw35_[int(14)] = d_247_v89_
                nw35_[int(15)] = d_247_v89_
                nw35_[int(16)] = d_247_v89_
                nw35_[int(17)] = d_247_v89_
                nw35_[int(18)] = d_247_v89_
                d_251_v93_ = nw35_
                index31_ = default__.safeIndex(944, (d_251_v93_).length(0))
                (d_251_v93_)[index31_] = d_247_v89_
                d_252_v94_: _dafny.Array
                d_252_v94_ = d_247_v89_
                index32_ = default__.safeIndex(944, (d_251_v93_).length(0))
                (d_251_v93_)[index32_] = ((d_252_v94_))
                d_253_v95_: C0
                nw36_ = C0()
                nw36_.ctor__((d_228_cf74_) == (d_228_cf74_))
                d_253_v95_ = nw36_
            elif True:
                d_254_v96_: T0
                nw37_ = C4()
                nw37_.ctor__(d_228_cf74_)
                d_254_v96_ = nw37_
                d_255_v97_: _dafny.Seq
                d_255_v97_ = _dafny.SeqWithoutIsStrInference([d_254_v96_, d_254_v96_])
                d_256_v98_: _dafny.Map
                d_256_v98_ = _dafny.Map({(d_199_v59_).f13: _dafny.MultiSet([d_255_v97_])})
                d_257_v99_: _dafny.Seq
                d_257_v99_ = _dafny.SeqWithoutIsStrInference([d_255_v97_])
                d_256_v98_ = (d_256_v98_).set((d_199_v59_).f13, _dafny.MultiSet(d_257_v99_))
                d_258_v100_: _dafny.Seq
                d_258_v100_ = _dafny.SeqWithoutIsStrInference([d_228_cf74_, (d_254_v96_).f6])
                d_258_v100_ = d_258_v100_
                (d_126_globalState_).f5 = ((d_199_v59_).f13) == ((d_199_v59_).f13)
                d_259_v101_: D8
                d_259_v101_ = D8_DC21((d_199_v59_).f13, default__.fm1(d_140_v12_, (d_199_v59_).f13, (d_254_v96_).f6, False, d_126_globalState_), d_130_v5_)
                (d_126_globalState_).f5 = not((d_259_v101_).cf37)
                d_260_v102_: _dafny.Seq
                d_260_v102_ = _dafny.SeqWithoutIsStrInference([(d_199_v59_).f13, d_127_v1_, False])
                (d_126_globalState_).f0 = (len(d_260_v102_)) - (d_130_v5_)
            d_261_v103_: _dafny.Array
            nw38_ = _dafny.Array(_dafny.Map({}), 20)
            d_261_v103_ = nw38_
            d_262_v104_: D14
            d_262_v104_ = D14_DC35(d_261_v103_)
            source11_ = d_262_v104_
            if source11_.is_DC36:
                d_263___mcc_h5_ = source11_.cf68
                d_264___mcc_h6_ = source11_.cf69
                d_265___mcc_h7_ = source11_.cf70
                d_266___mcc_h8_ = source11_.cf71
                d_267_cf71_ = d_266___mcc_h8_
                d_268_cf70_ = d_265___mcc_h7_
                d_269_cf69_ = d_264___mcc_h6_
                d_270_cf68_ = d_263___mcc_h5_
                (d_126_globalState_).f1 = default__.fm1(d_140_v12_, not(d_270_cf68_), d_268_cf70_, (d_199_v59_).f13, d_126_globalState_)
                d_127_v1_ = ((d_268_cf70_) * (len(d_125_v0_))) != ((len(d_125_v0_)) - (d_228_cf74_))
                d_271_v105_: _dafny.Seq
                d_271_v105_ = _dafny.SeqWithoutIsStrInference([D1_DC5(), d_229_v77_, (D1_DC5() if d_270_cf68_ else d_229_v77_), D1_DC5(), d_229_v77_])
                d_271_v105_ = (d_271_v105_ if d_127_v1_ else d_271_v105_)
                default__.m7(d_130_v5_, d_126_globalState_)
            elif True:
                d_272___mcc_h9_ = source11_.cf67
                d_273_cf67_ = d_272___mcc_h9_
                d_274_v106_: _dafny.Array
                nw39_ = _dafny.Array(None, 3)
                nw39_[int(0)] = (d_199_v59_).f13
                nw39_[int(1)] = (d_199_v59_).f13
                nw39_[int(2)] = (d_199_v59_).f13
                d_274_v106_ = nw39_
                index33_ = default__.safeIndex(482, (d_274_v106_).length(0))
                (d_274_v106_)[index33_] = (d_199_v59_).f13
                d_275_v107_: _dafny.Set
                d_275_v107_ = _dafny.Set({(d_125_v0_).set(default__.safeIndex(d_130_v5_, len(d_125_v0_)), d_131_v7_)})
                d_276_v108_: _dafny.Map
                d_276_v108_ = _dafny.Map({d_130_v5_: default__.fm8(d_131_v7_, d_127_v1_, -237, default__.fm24(d_130_v5_, d_126_globalState_), d_126_globalState_)})
                index34_ = default__.safeIndex(482, (d_274_v106_).length(0))
                rhs14_ = not((len((d_128_v2_) - (d_128_v2_))) < (285))
                rhs15_ = ((d_276_v108_)[d_130_v5_] if (d_130_v5_) in (d_276_v108_) else d_131_v7_)
                rhs16_ = (d_199_v59_).f13
                rhs17_ = _dafny.Set({d_125_v0_, _dafny.SeqWithoutIsStrInference([d_131_v7_ for d_277_i6_ in range(default__.abs(569))]), (d_125_v0_) + (d_125_v0_)})
                lhs14_ = d_126_globalState_
                lhs15_ = d_274_v106_
                lhs16_ = default__.safeIndex(482, (d_274_v106_).length(0))
                lhs14_.f1 = rhs14_
                d_131_v7_ = rhs15_
                lhs15_[lhs16_] = rhs16_
                d_275_v107_ = rhs17_
                index35_ = default__.safeIndex(482, (d_274_v106_).length(0))
                (d_274_v106_)[index35_] = default__.fm1(d_140_v12_, (d_130_v5_) >= (695), (((d_132_v8_)[d_228_cf74_] if (d_228_cf74_) in (d_132_v8_) else d_130_v5_)) * (-56), (d_199_v59_).f13, d_126_globalState_)
                d_278_v109_: _dafny.Array
                nw40_ = _dafny.Array(int(0), 2)
                d_278_v109_ = nw40_
                rhs18_ = d_278_v109_
                rhs19_ = default__.fm24(d_130_v5_, d_126_globalState_)
                rhs20_ = (d_130_v5_) > (len((d_125_v0_) + (d_125_v0_)))
                d_278_v109_ = rhs18_
                d_130_v5_ = rhs19_
                d_127_v1_ = rhs20_
                d_279_v110_: D8
                d_279_v110_ = D8_DC21((d_274_v106_)[default__.safeIndex(482, (d_274_v106_).length(0))], (d_199_v59_).f13, d_228_cf74_)
                (d_126_globalState_).f2 = (d_279_v110_).cf38
            d_280_v111_: _dafny.Map
            d_280_v111_ = _dafny.Map({d_229_v77_: (d_199_v59_).f13})
            rhs21_ = default__.safeModuloInt(d_130_v5_, d_228_cf74_)
            rhs22_ = _dafny.Map({d_229_v77_: False})
            d_130_v5_ = rhs21_
            d_280_v111_ = rhs22_
        elif True:
            d_281___mcc_h3_ = source9_.cf73
            d_282_cf73_ = d_281___mcc_h3_
            d_125_v0_ = d_125_v0_
            d_283_v112_: _dafny.Array
            def lambda12_(d_284_v59_):
                def lambda13_(d_285_i7_):
                    return (d_284_v59_).f13

                return lambda13_

            init6_ = lambda12_(d_199_v59_)
            nw41_ = _dafny.Array(None, 22)
            for i0_6_ in range(nw41_.length(0)):
                nw41_[i0_6_] = init6_(i0_6_)
            d_283_v112_ = nw41_
            d_286_v113_: _dafny.Seq
            d_286_v113_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([-204, d_130_v5_, d_130_v5_, d_130_v5_, d_130_v5_])])
            index36_ = default__.safeIndex(421, (d_283_v112_).length(0))
            (d_283_v112_)[index36_] = (d_199_v59_).f13
            d_287_v114_: _dafny.Seq
            d_287_v114_ = _dafny.SeqWithoutIsStrInference([d_128_v2_, d_128_v2_, d_128_v2_])
            index37_ = default__.safeIndex(421, (d_283_v112_).length(0))
            rhs23_ = (0) - (len(((d_287_v114_)[default__.safeIndex(d_130_v5_, len(d_287_v114_))]) - (_dafny.Set({(d_199_v59_).f13}))))
            rhs24_ = ((_dafny.Map({False: len(d_125_v0_)})) | (d_218_v72_)) | (_dafny.Map({(d_199_v59_).f13: len(d_125_v0_)}))
            rhs25_ = d_283_v112_
            rhs26_ = default__.fm40(d_126_globalState_)
            rhs27_ = d_127_v1_
            lhs17_ = d_126_globalState_
            lhs18_ = d_283_v112_
            lhs19_ = default__.safeIndex(421, (d_283_v112_).length(0))
            lhs17_.f2 = rhs23_
            d_218_v72_ = rhs24_
            d_283_v112_ = rhs25_
            d_286_v113_ = rhs26_
            lhs18_[lhs19_] = rhs27_
            d_288_v115_: T1
            nw42_ = C1()
            nw42_.ctor__()
            d_288_v115_ = nw42_
            d_289_v116_: _dafny.Map
            d_289_v116_ = _dafny.Map({d_288_v115_: (d_283_v112_)[default__.safeIndex(421, (d_283_v112_).length(0))]})
            def iife47_():
                coll19_ = _dafny.Set()
                compr_19_: int
                for compr_19_ in (d_221_v75_).keys.Elements:
                    d_290_v117_: int = compr_19_
                    if (d_290_v117_) in (d_221_v75_):
                        coll19_ = coll19_.union(_dafny.Set([(d_290_v117_) + (734)]))
                return _dafny.Set(coll19_)
            def iife48_():
                coll20_ = _dafny.Set()
                compr_20_: int
                for compr_20_ in _dafny.IntegerRange(-261, 593):
                    d_291_v118_: int = compr_20_
                    if ((-261) <= (d_291_v118_)) and ((d_291_v118_) < (593)):
                        coll20_ = coll20_.union(_dafny.Set([(d_291_v118_) * (len(d_125_v0_))]))
                return _dafny.Set(coll20_)
            rhs28_ = d_289_v116_
            rhs29_ = d_128_v2_
            rhs30_ = d_199_v59_
            rhs31_ = (0) - ((d_130_v5_) + ((0) - (len(d_218_v72_))))
            rhs32_ = ((iife47_()
            ) - (d_220_v74_) if (d_283_v112_)[default__.safeIndex(421, (d_283_v112_).length(0))] else iife48_()
            )
            lhs20_ = d_126_globalState_
            d_289_v116_ = rhs28_
            d_128_v2_ = rhs29_
            d_199_v59_ = rhs30_
            lhs20_.f0 = rhs31_
            d_220_v74_ = rhs32_
            d_292_v119_: C2
            nw43_ = C2()
            nw43_.ctor__(d_125_v0_, True)
            d_292_v119_ = nw43_
            d_292_v119_ = d_292_v119_
        d_293_v120_: _dafny.Set
        d_293_v120_ = _dafny.Set({d_199_v59_})
        d_294_v121_: C0
        d_294_v121_ = d_199_v59_
        d_293_v120_ = _dafny.Set({d_199_v59_, (d_199_v59_), d_199_v59_, d_199_v59_})
        d_295_v122_: _dafny.MultiSet
        d_295_v122_ = _dafny.MultiSet([d_127_v1_, d_127_v1_, d_127_v1_, d_127_v1_, False])
        d_296_v123_: _dafny.Map
        d_296_v123_ = _dafny.Map({((d_295_v122_).cardinality) != (len(d_128_v2_)): d_128_v2_})
        d_296_v123_ = (d_296_v123_) | ((d_296_v123_) | (d_296_v123_))
        hi1_ = (d_130_v5_) - (d_130_v5_)
        for d_297_i8_ in range(len(d_125_v0_), hi1_):
            rhs33_ = d_131_v7_
            rhs34_ = (0) - (d_297_i8_)
            lhs21_ = d_126_globalState_
            d_131_v7_ = rhs33_
            lhs21_.f2 = rhs34_
            d_298_v124_: _dafny.Array
            def lambda14_(d_299_v59_):
                def lambda15_(d_300_i9_):
                    return (d_299_v59_).f13

                return lambda15_

            init7_ = lambda14_(d_199_v59_)
            nw44_ = _dafny.Array(None, 29)
            for i0_7_ in range(nw44_.length(0)):
                nw44_[i0_7_] = init7_(i0_7_)
            d_298_v124_ = nw44_
            index38_ = default__.safeIndex(752, (d_298_v124_).length(0))
            (d_298_v124_)[index38_] = d_127_v1_
            d_301_v125_: _dafny.Seq
            d_301_v125_ = _dafny.SeqWithoutIsStrInference([d_297_i8_, d_130_v5_])
            d_302_v126_: _dafny.Map
            d_302_v126_ = _dafny.Map({(d_199_v59_).f13: default__.fm8(d_131_v7_, (d_199_v59_).f13, d_130_v5_, d_130_v5_, d_126_globalState_)})
            index39_ = default__.safeIndex(752, (d_298_v124_).length(0))
            (d_298_v124_)[index39_] = default__.fm1(d_140_v12_, True, (d_301_v125_)[default__.safeIndex(d_130_v5_, len(d_301_v125_))], (len(d_302_v126_)) < (d_297_i8_), d_126_globalState_)
            (d_126_globalState_).f1 = True
            index40_ = default__.safeIndex(752, (d_298_v124_).length(0))
            (d_298_v124_)[index40_] = d_127_v1_
        default__.m7(d_130_v5_, d_126_globalState_)
        d_303_v127_: _dafny.Seq
        d_303_v127_ = _dafny.SeqWithoutIsStrInference([d_199_v59_])
        d_304_v128_: _dafny.Array
        nw45_ = _dafny.Array(None, 13)
        nw45_[int(0)] = (773) + (d_130_v5_)
        nw45_[int(1)] = default__.fm15((0) - (d_130_v5_), d_127_v1_, d_126_globalState_)
        nw45_[int(2)] = (d_130_v5_) * ((d_133_v9_).cf9)
        nw45_[int(3)] = d_130_v5_
        nw45_[int(4)] = d_130_v5_
        nw45_[int(5)] = (d_130_v5_) + (len(d_303_v127_))
        nw45_[int(6)] = (d_222_v76_).cf74
        nw45_[int(7)] = d_130_v5_
        nw45_[int(8)] = d_130_v5_
        nw45_[int(9)] = d_130_v5_
        nw45_[int(10)] = d_130_v5_
        nw45_[int(11)] = (d_130_v5_) - (d_130_v5_)
        nw45_[int(12)] = (0) - ((d_130_v5_ if (d_199_v59_).f13 else len(_dafny.SeqWithoutIsStrInference([not((d_199_v59_).f13), d_127_v1_, d_127_v1_]))))
        d_304_v128_ = nw45_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_304_v128_).length(0)):
            d_305_i10_: int = guard_loop_0_
            if (True) and (((0) <= (d_305_i10_)) and ((d_305_i10_) < ((d_304_v128_).length(0)))):
                (d_304_v128_)[(d_305_i10_)] = default__.safeModuloInt(d_305_i10_, d_130_v5_)
        hi2_ = (d_130_v5_) + (d_130_v5_)
        for d_306_i11_ in range(d_130_v5_, hi2_):
            d_307_v129_: T1
            nw46_ = C1()
            nw46_.ctor__()
            d_307_v129_ = nw46_
            d_308_v130_: C6
            nw47_ = C6()
            nw47_.ctor__(d_130_v5_)
            d_308_v130_ = nw47_
            d_309_v131_: _dafny.Map
            d_309_v131_ = _dafny.Map({d_307_v129_: d_308_v130_})
            d_309_v131_ = d_309_v131_
            (d_126_globalState_).f5 = (d_199_v59_).f13
            d_310_v132_: _dafny.Array
            def lambda16_(d_311_v59_):
                def lambda17_(d_312_i12_):
                    return (d_311_v59_).f13

                return lambda17_

            init8_ = lambda16_(d_199_v59_)
            nw48_ = _dafny.Array(None, 25)
            for i0_8_ in range(nw48_.length(0)):
                nw48_[i0_8_] = init8_(i0_8_)
            d_310_v132_ = nw48_
            index41_ = default__.safeIndex(699, (d_310_v132_).length(0))
            (d_310_v132_)[index41_] = (d_125_v0_) <= (d_125_v0_)
            index42_ = default__.safeIndex(699, (d_310_v132_).length(0))
            (d_310_v132_)[index42_] = (d_199_v59_).f13
            d_125_v0_ = d_125_v0_
        if (d_199_v59_).f13:
            if (d_199_v59_).f13:
                d_313_v133_: _dafny.Map
                d_313_v133_ = _dafny.Map({(d_199_v59_).fm23(d_130_v5_, d_130_v5_, d_131_v7_, (d_199_v59_).f13, d_126_globalState_): ((_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vky"))) for d_314_i13_ in range(default__.abs(-126))])) + (_dafny.SeqWithoutIsStrInference([d_130_v5_]))).set(default__.safeIndex(d_130_v5_, len((_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vky"))) for d_315_i13_ in range(default__.abs(-126))])) + (_dafny.SeqWithoutIsStrInference([d_130_v5_])))), 224)})
                d_316_v134_: _dafny.Seq
                d_316_v134_ = _dafny.SeqWithoutIsStrInference([d_130_v5_])
                d_313_v133_ = (d_313_v133_).set(d_130_v5_, (_dafny.SeqWithoutIsStrInference([len((D12_DC29(d_125_v0_, (d_199_v59_).f13)).cf52) for d_317_i14_ in range(default__.abs(860))])) + (d_316_v134_))
                d_318_v135_: _dafny.Array
                nw49_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 19)
                d_318_v135_ = nw49_
                d_318_v135_ = d_318_v135_
                (d_126_globalState_).f0 = default__.safeDivisionInt(d_130_v5_, d_130_v5_)
                (d_126_globalState_).f0 = ((d_316_v134_)[default__.safeIndex((d_199_v59_).fm23(d_130_v5_, len(d_221_v75_), d_131_v7_, d_127_v1_, d_126_globalState_), len(d_316_v134_))]) - (len((d_128_v2_).intersection(_dafny.Set({not((d_199_v59_).f13)}))))
                d_319_v136_: _dafny.Seq
                d_319_v136_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_131_v7_ for d_320_i15_ in range(default__.abs(-864))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c")), d_125_v0_, d_125_v0_])
                default__.m7(default__.safeModuloInt(d_130_v5_, len(d_319_v136_)), d_126_globalState_)
            elif True:
                d_321_v137_: _dafny.Array
                def lambda18_(d_322_v1_):
                    def lambda19_(d_323_i16_):
                        return d_322_v1_

                    return lambda19_

                init9_ = lambda18_(d_127_v1_)
                nw50_ = _dafny.Array(None, 20)
                for i0_9_ in range(nw50_.length(0)):
                    nw50_[i0_9_] = init9_(i0_9_)
                d_321_v137_ = nw50_
                index43_ = default__.safeIndex(331, (d_321_v137_).length(0))
                (d_321_v137_)[index43_] = not(d_127_v1_)
                index44_ = default__.safeIndex(785, (d_304_v128_).length(0))
                (d_304_v128_)[index44_] = d_130_v5_
                d_324_v138_: _dafny.Seq
                d_324_v138_ = _dafny.SeqWithoutIsStrInference([d_130_v5_])
                d_325_v139_: _dafny.Map
                d_325_v139_ = _dafny.Map({d_324_v138_: d_304_v128_})
                d_326_v140_: D8
                d_326_v140_ = D8_DC20(d_325_v139_)
                d_327_v141_: _dafny.Set
                d_327_v141_ = _dafny.Set({d_326_v140_, d_326_v140_})
                index45_ = default__.safeIndex(331, (d_321_v137_).length(0))
                index46_ = default__.safeIndex(785, (d_304_v128_).length(0))
                rhs35_ = (0) - ((d_324_v138_)[default__.safeIndex((d_130_v5_) * (d_130_v5_), len(d_324_v138_))])
                rhs36_ = (d_132_v8_).issubset(d_132_v8_)
                rhs37_ = not(((d_327_v141_).intersection(d_327_v141_)).issubset(d_327_v141_))
                rhs38_ = (d_199_v59_).f13
                rhs39_ = default__.safeModuloInt(len(d_125_v0_), default__.safeModuloInt(d_130_v5_, d_130_v5_))
                lhs22_ = d_126_globalState_
                lhs23_ = d_321_v137_
                lhs24_ = default__.safeIndex(331, (d_321_v137_).length(0))
                lhs25_ = d_126_globalState_
                lhs26_ = d_126_globalState_
                lhs27_ = d_304_v128_
                lhs28_ = default__.safeIndex(785, (d_304_v128_).length(0))
                lhs22_.f0 = rhs35_
                lhs23_[lhs24_] = rhs36_
                lhs25_.f1 = rhs37_
                lhs26_.f5 = rhs38_
                lhs27_[lhs28_] = rhs39_
                d_328_v142_: D5
                d_328_v142_ = D5_DC13(d_127_v1_, d_127_v1_, d_125_v0_, -978)
                d_329_v143_: D3
                d_329_v143_ = D3_DC7((d_328_v142_).cf23)
                index47_ = default__.safeIndex(331, (d_321_v137_).length(0))
                rhs40_ = d_329_v143_
                rhs41_ = (len(d_125_v0_)) + (((d_304_v128_)[default__.safeIndex(785, (d_304_v128_).length(0))] if True else d_130_v5_))
                rhs42_ = (d_199_v59_).f13
                lhs29_ = d_126_globalState_
                lhs30_ = d_321_v137_
                lhs31_ = default__.safeIndex(331, (d_321_v137_).length(0))
                d_329_v143_ = rhs40_
                lhs29_.f2 = rhs41_
                lhs30_[lhs31_] = rhs42_
                d_330_v144_: D8
                d_330_v144_ = D8_DC21((d_199_v59_).f13, False, (len(d_125_v0_)) * ((_dafny.MultiSet([d_127_v1_, d_127_v1_, (d_199_v59_).f13, d_127_v1_])).cardinality))
                d_330_v144_ = d_330_v144_
                d_331_v145_: C4
                nw51_ = C4()
                nw51_.ctor__((d_304_v128_)[default__.safeIndex(785, (d_304_v128_).length(0))])
                d_331_v145_ = nw51_
                d_332_v146_: _dafny.Array
                def lambda20_(d_333_v5_):
                    def lambda21_(d_334_i17_):
                        return default__.safeDivisionInt(d_334_i17_, d_333_v5_)

                    return lambda21_

                init10_ = lambda20_(d_130_v5_)
                nw52_ = _dafny.Array(None, 15)
                for i0_10_ in range(nw52_.length(0)):
                    nw52_[i0_10_] = init10_(i0_10_)
                d_332_v146_ = nw52_
                d_335_v147_: _dafny.Map
                d_335_v147_ = _dafny.Map({d_332_v146_: (0) - (d_130_v5_)})
                (d_126_globalState_).f5 = (d_332_v146_) in (d_335_v147_)
            (d_126_globalState_).f5 = not((d_199_v59_).f13)
            d_336_v148_: _dafny.Array
            nw53_ = _dafny.Array(_dafny.Array(None, 0), 15)
            d_336_v148_ = nw53_
            d_337_v149_: _dafny.Array
            nw54_ = _dafny.Array(_dafny.Array(None, 0), 2)
            d_337_v149_ = nw54_
            index48_ = default__.safeIndex(896, (d_336_v148_).length(0))
            (d_336_v148_)[index48_] = d_337_v149_
            d_338_v150_: T1
            nw55_ = C1()
            nw55_.ctor__()
            d_338_v150_ = nw55_
            d_339_v151_: _dafny.Array
            nw56_ = _dafny.Array(None, 22)
            nw56_[int(0)] = (d_199_v59_).f13
            nw56_[int(1)] = (d_199_v59_).f13
            nw56_[int(2)] = (d_199_v59_).f13
            nw56_[int(3)] = d_127_v1_
            nw56_[int(4)] = (d_199_v59_).f13
            nw56_[int(5)] = (d_199_v59_).f13
            nw56_[int(6)] = (d_199_v59_).f13
            nw56_[int(7)] = True
            nw56_[int(8)] = (d_199_v59_).f13
            nw56_[int(9)] = not(d_127_v1_)
            nw56_[int(10)] = d_127_v1_
            nw56_[int(11)] = not((d_199_v59_).f13)
            nw56_[int(12)] = (d_199_v59_).f13
            nw56_[int(13)] = False
            nw56_[int(14)] = (d_199_v59_).f13
            nw56_[int(15)] = (d_199_v59_).f13
            nw56_[int(16)] = (d_199_v59_).f13
            nw56_[int(17)] = (d_199_v59_).f13
            nw56_[int(18)] = not((d_199_v59_).f13)
            nw56_[int(19)] = not((d_199_v59_).f13)
            nw56_[int(20)] = not(True)
            nw56_[int(21)] = d_127_v1_
            d_339_v151_ = nw56_
            d_340_v152_: D11
            d_340_v152_ = D11_DC25(d_339_v151_, d_338_v150_, len(d_220_v74_))
            d_341_v153_: _dafny.Array
            nw57_ = _dafny.Array(None, 27)
            nw57_[int(0)] = d_338_v150_
            nw57_[int(1)] = d_338_v150_
            nw57_[int(2)] = d_338_v150_
            nw57_[int(3)] = (d_338_v150_ if (d_199_v59_).f13 else d_338_v150_)
            nw57_[int(4)] = d_338_v150_
            nw57_[int(5)] = d_338_v150_
            nw57_[int(6)] = d_338_v150_
            nw57_[int(7)] = d_338_v150_
            nw57_[int(8)] = d_338_v150_
            nw57_[int(9)] = d_338_v150_
            nw57_[int(10)] = (d_340_v152_).cf43
            nw57_[int(11)] = d_338_v150_
            nw57_[int(12)] = d_338_v150_
            nw57_[int(13)] = d_338_v150_
            nw57_[int(14)] = d_338_v150_
            nw57_[int(15)] = d_338_v150_
            nw57_[int(16)] = d_338_v150_
            nw57_[int(17)] = d_338_v150_
            nw57_[int(18)] = d_338_v150_
            nw57_[int(19)] = d_338_v150_
            nw57_[int(20)] = d_338_v150_
            nw57_[int(21)] = d_338_v150_
            nw57_[int(22)] = d_338_v150_
            nw57_[int(23)] = d_338_v150_
            nw57_[int(24)] = d_338_v150_
            nw57_[int(25)] = d_338_v150_
            nw57_[int(26)] = d_338_v150_
            d_341_v153_ = nw57_
            index49_ = default__.safeIndex(892, (d_341_v153_).length(0))
            (d_341_v153_)[index49_] = d_338_v150_
            index50_ = default__.safeIndex(896, (d_336_v148_).length(0))
            index51_ = default__.safeIndex(892, (d_341_v153_).length(0))
            rhs43_ = d_337_v149_
            rhs44_ = d_338_v150_
            lhs32_ = d_336_v148_
            lhs33_ = default__.safeIndex(896, (d_336_v148_).length(0))
            lhs34_ = d_341_v153_
            lhs35_ = default__.safeIndex(892, (d_341_v153_).length(0))
            lhs32_[lhs33_] = rhs43_
            lhs34_[lhs35_] = rhs44_
            def iife49_():
                coll21_ = _dafny.Set()
                compr_21_: int
                for compr_21_ in _dafny.IntegerRange(306, -263):
                    d_342_v154_: int = compr_21_
                    if ((306) <= (d_342_v154_)) and ((d_342_v154_) < (-263)):
                        coll21_ = coll21_.union(_dafny.Set([(d_342_v154_) - (d_130_v5_)]))
                return _dafny.Set(coll21_)
            d_220_v74_ = iife49_()
            
            if not (d_127_v1_) or ((default__.fm1(d_140_v12_, (d_199_v59_).f13, d_130_v5_, (d_199_v59_).f13, d_126_globalState_) if (d_199_v59_).f13 else (d_199_v59_).f13)):
                d_343_v155_: C5
                nw58_ = C5()
                nw58_.ctor__(not(not(d_127_v1_)), d_131_v7_)
                d_343_v155_ = nw58_
                d_344_v156_: C1
                nw59_ = C1()
                nw59_.ctor__()
                d_344_v156_ = nw59_
                d_345_v157_: D18
                d_345_v157_ = D18_DC43(d_344_v156_)
                d_346_v158_: _dafny.Map
                d_346_v158_ = _dafny.Map({d_130_v5_: d_344_v156_})
                d_347_v159_: _dafny.Seq
                d_347_v159_ = _dafny.SeqWithoutIsStrInference([d_130_v5_, d_130_v5_])
                d_348_v160_: T0
                nw60_ = C3()
                nw60_.ctor__(d_130_v5_, d_347_v159_, 507)
                d_348_v160_ = nw60_
                d_349_v161_: _dafny.Map
                d_349_v161_ = _dafny.Map({d_348_v160_: _dafny.CodePoint('k')})
                d_350_v162_: _dafny.MultiSet
                d_350_v162_ = _dafny.MultiSet([d_349_v161_, d_349_v161_])
                d_351_v163_: _dafny.Array
                nw61_ = _dafny.Array(None, 24)
                nw61_[int(0)] = d_344_v156_
                nw61_[int(1)] = d_344_v156_
                nw61_[int(2)] = d_344_v156_
                nw61_[int(3)] = d_344_v156_
                nw61_[int(4)] = d_344_v156_
                nw61_[int(5)] = d_344_v156_
                nw61_[int(6)] = d_344_v156_
                nw61_[int(7)] = d_344_v156_
                nw61_[int(8)] = d_344_v156_
                nw61_[int(9)] = d_344_v156_
                nw61_[int(10)] = d_344_v156_
                nw61_[int(11)] = d_344_v156_
                nw61_[int(12)] = d_344_v156_
                nw61_[int(13)] = (d_345_v157_).cf80
                nw61_[int(14)] = d_344_v156_
                nw61_[int(15)] = d_344_v156_
                nw61_[int(16)] = d_344_v156_
                nw61_[int(17)] = d_344_v156_
                nw61_[int(18)] = d_344_v156_
                nw61_[int(19)] = d_344_v156_
                nw61_[int(20)] = ((d_346_v158_)[(d_350_v162_).cardinality] if ((d_350_v162_).cardinality) in (d_346_v158_) else d_344_v156_)
                nw61_[int(21)] = d_344_v156_
                nw61_[int(22)] = d_344_v156_
                nw61_[int(23)] = d_344_v156_
                d_351_v163_ = nw61_
                index52_ = default__.safeIndex(305, (d_351_v163_).length(0))
                (d_351_v163_)[index52_] = d_344_v156_
                index53_ = default__.safeIndex(305, (d_351_v163_).length(0))
                (d_351_v163_)[index53_] = d_344_v156_
                (d_126_globalState_).f5 = (d_199_v59_).f13
                d_131_v7_ = d_131_v7_
                d_352_v164_: _dafny.Array
                nw62_ = _dafny.Array(_dafny.CodePoint('D'), 17)
                d_352_v164_ = nw62_
                d_353_v165_: D19
                d_353_v165_ = D19_DC45(d_352_v164_)
                d_352_v164_ = (d_353_v165_).cf86
            elif True:
                d_354_v166_: _dafny.Map
                d_354_v166_ = _dafny.Map({(d_341_v153_)[default__.safeIndex(892, (d_341_v153_).length(0))]: d_127_v1_})
                d_355_v167_: D13
                d_355_v167_ = D13_DC32(d_130_v5_, False, (d_199_v59_).f13, (d_199_v59_).f13, (d_199_v59_).f13)
                d_354_v166_ = (d_354_v166_).set((d_341_v153_)[default__.safeIndex(892, (d_341_v153_).length(0))], (d_355_v167_).cf57)
                d_356_v168_: _dafny.Seq
                d_356_v168_ = _dafny.SeqWithoutIsStrInference([d_127_v1_])
                (d_126_globalState_).f1 = ((d_356_v168_) + (d_356_v168_)) < ((d_356_v168_) + (d_356_v168_))
                d_357_v169_: _dafny.Array
                def lambda22_(d_358_v7_):
                    def lambda23_(d_359_i18_):
                        return d_358_v7_

                    return lambda23_

                init11_ = lambda22_(d_131_v7_)
                nw63_ = _dafny.Array(None, 29)
                for i0_11_ in range(nw63_.length(0)):
                    nw63_[i0_11_] = init11_(i0_11_)
                d_357_v169_ = nw63_
                index54_ = default__.safeIndex(492, (d_357_v169_).length(0))
                (d_357_v169_)[index54_] = d_131_v7_
                d_360_v170_: _dafny.Map
                d_360_v170_ = _dafny.Map({(d_199_v59_).f13: (d_199_v59_).f13})
                index55_ = default__.safeIndex(904, (d_339_v151_).length(0))
                (d_339_v151_)[index55_] = ((d_360_v170_)[d_127_v1_] if (d_127_v1_) in (d_360_v170_) else d_127_v1_)
                index56_ = default__.safeIndex(492, (d_357_v169_).length(0))
                index57_ = default__.safeIndex(904, (d_339_v151_).length(0))
                rhs45_ = (default__.fm19((d_199_v59_).f13, d_131_v7_, _dafny.Map({((d_218_v72_)[d_127_v1_] if (d_127_v1_) in (d_218_v72_) else d_130_v5_): True}), d_127_v1_, d_126_globalState_)).cf6
                rhs46_ = not ((d_199_v59_).f13) or (not((d_125_v0_) < ((d_125_v0_).set(default__.safeIndex(d_130_v5_, len(d_125_v0_)), d_131_v7_))))
                lhs36_ = d_357_v169_
                lhs37_ = default__.safeIndex(492, (d_357_v169_).length(0))
                lhs38_ = d_339_v151_
                lhs39_ = default__.safeIndex(904, (d_339_v151_).length(0))
                lhs36_[lhs37_] = rhs45_
                lhs38_[lhs39_] = rhs46_
                default__.m7(d_130_v5_, d_126_globalState_)
                (d_126_globalState_).f0 = (0) - ((default__.safeDivisionInt(d_130_v5_, len(d_125_v0_))) * ((d_199_v59_).fm23(d_130_v5_, d_130_v5_, _dafny.CodePoint('n'), (D8_DC21(d_127_v1_, (d_199_v59_).f13, d_130_v5_)).cf37, d_126_globalState_)))
        elif True:
            d_361_v171_: C4
            nw64_ = C4()
            nw64_.ctor__(d_130_v5_)
            d_361_v171_ = nw64_
            d_362_v172_: _dafny.Seq
            d_362_v172_ = _dafny.SeqWithoutIsStrInference([771, d_130_v5_, d_130_v5_])
            d_363_v173_: _dafny.Map
            d_363_v173_ = _dafny.Map({d_130_v5_: d_125_v0_})
            d_364_v174_: _dafny.MultiSet
            d_364_v174_ = _dafny.MultiSet([d_362_v172_, _dafny.SeqWithoutIsStrInference([len(((d_363_v173_)[d_130_v5_] if (d_130_v5_) in (d_363_v173_) else d_125_v0_)), len(default__.fm21(d_125_v0_, d_126_globalState_))]), d_362_v172_])
            d_365_v175_: _dafny.Seq
            d_365_v175_ = _dafny.SeqWithoutIsStrInference([d_304_v128_, d_304_v128_])
            d_366_v176_: _dafny.Map
            d_366_v176_ = _dafny.Map({(_dafny.MultiSet(default__.fm40(d_126_globalState_)) if (d_199_v59_).f13 else d_364_v174_): default__.fm15(default__.fm15((0) - (len(d_365_v175_)), (d_199_v59_).f13, d_126_globalState_), (d_199_v59_).f13, d_126_globalState_)})
            d_366_v176_ = (d_366_v176_).set(d_364_v174_, d_130_v5_)
            d_367_v177_: _dafny.Array
            nw65_ = _dafny.Array(D14.default()(), 21)
            d_367_v177_ = nw65_
            d_367_v177_ = d_367_v177_
            d_368_v178_: _dafny.Map
            d_368_v178_ = _dafny.Map({d_130_v5_: d_131_v7_})
            d_369_v179_: C3
            nw66_ = C3()
            nw66_.ctor__(default__.safeModuloInt(d_130_v5_, len(_dafny.SeqWithoutIsStrInference([(d_199_v59_).f13, (d_199_v59_).f13, d_127_v1_, d_127_v1_]))), (_dafny.SeqWithoutIsStrInference([len(d_368_v178_)])).set(default__.safeIndex(len(default__.fm2(d_125_v0_, d_126_globalState_)), len(_dafny.SeqWithoutIsStrInference([len(d_368_v178_)]))), d_130_v5_), d_130_v5_)
            d_369_v179_ = nw66_
            d_127_v1_ = not (d_127_v1_) or (False)
        (d_126_globalState_).f0 = len(((d_125_v0_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_370_i19_ in range(default__.abs(788))]))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oobul"))))
        _dafny.print((d_125_v0_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_126_globalState_.f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_126_globalState_.f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_126_globalState_.f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_globalState_).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_126_globalState_).f4).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_126_globalState_.f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_127_v1_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_128_v2_) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_129_v3_) == (_dafny.Map({277: _dafny.Set({False})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_130_v5_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_131_v7_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_132_v8_) == (_dafny.MultiSet([872, 372, 371, 371]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_133_v9_).cf5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_133_v9_).cf6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_133_v9_).cf7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_133_v9_).cf8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_133_v9_).cf9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_134_v10_)[0]) == (_dafny.Map({277: _dafny.Set({False})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_134_v10_)[1]) == (_dafny.Map({277: _dafny.Set({False})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_134_v10_)[2]) == (_dafny.Map({277: _dafny.Set({False})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_134_v10_)[3]) == (_dafny.Map({277: _dafny.Set({False})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_134_v10_)[4]) == (_dafny.Map({277: _dafny.Set({False})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_134_v10_)[5]) == (_dafny.Map({277: _dafny.Set({False})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_134_v10_)[6]) == (_dafny.Map({277: _dafny.Set({False})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_134_v10_)[7]) == (_dafny.Map({277: _dafny.Set({False})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_134_v10_)[8]) == (_dafny.Map({277: _dafny.Set({False})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_134_v10_)[9]) == (_dafny.Map({277: _dafny.Set({False})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_134_v10_)[10]) == (_dafny.Map({742: _dafny.Set({False})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_134_v10_)[11]) == (_dafny.Map({371: _dafny.Set({False})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_134_v10_)[12]) == (_dafny.Map({371: _dafny.Set({False})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_134_v10_)[13]) == (_dafny.Map({277: _dafny.Set({False})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_134_v10_)[14]) == (_dafny.Map({0: _dafny.Set({False})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_134_v10_)[15]) == (_dafny.Map({277: _dafny.Set({False})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_134_v10_)[16]) == (_dafny.Map({}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_139_v11_) == (_dafny.Map({371: _dafny.Map({277: _dafny.Set({False})})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_140_v12_).cf0) == (_dafny.Map({277: _dafny.Set({False})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_199_v59_).f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_218_v72_) == (_dafny.Map({False: 371}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_219_v73_) == (_dafny.Map({371: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_220_v74_) == (_dafny.Set({1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_221_v75_) == (_dafny.Map({1: 371}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_222_v76_).cf74))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_222_v76_).cf75) == (_dafny.Map({1: 371}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_293_v120_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_294_v121_)).f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_295_v122_) == (_dafny.MultiSet([True, True, True, True, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_296_v123_) == (_dafny.Map({True: _dafny.Set({False})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_303_v127_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_304_v128_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_304_v128_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_304_v128_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_304_v128_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_304_v128_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_304_v128_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_304_v128_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_304_v128_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_304_v128_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_304_v128_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_304_v128_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_304_v128_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_304_v128_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(int(0), False, int(0), False)
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

class D0_DC2(D0, NamedTuple('DC2', [('cf5', Any), ('cf6', Any), ('cf7', Any), ('cf8', Any), ('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2({_dafny.string_of(self.cf5)}, {_dafny.string_of(self.cf6)}, {_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)}, {_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2) and self.cf5 == __o.cf5 and self.cf6 == __o.cf6 and self.cf7 == __o.cf7 and self.cf8 == __o.cf8 and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC0(D0, NamedTuple('DC0', [('cf0', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC0({_dafny.string_of(self.cf0)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC0) and self.cf0 == __o.cf0
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC3(D0, NamedTuple('DC3', [('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC3({_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC3) and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()


class D1:
    @classmethod
    def default(cls, ):
        return lambda: D1_DC5()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D1_DC5)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)

class D1_DC5(D1, NamedTuple('DC5', [])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5)
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC4(D1, NamedTuple('DC4', [('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)

class D2_DC6(D2, NamedTuple('DC6', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC8(D1.default()(), _dafny.Array(None, 0), int(0))
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

class D3_DC8(D3, NamedTuple('DC8', [('cf14', Any), ('cf15', Any), ('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC8({_dafny.string_of(self.cf14)}, {_dafny.string_of(self.cf15)}, {_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC8) and self.cf14 == __o.cf14 and self.cf15 == __o.cf15 and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC7(D3, NamedTuple('DC7', [('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC7({self.cf13.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC7) and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC9(D3, NamedTuple('DC9', [('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D4_DC10)

class D4_DC10(D4, NamedTuple('DC10', [('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC10({_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC10) and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC12(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D5_DC12)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D5_DC13)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D5_DC11)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D5_DC14)

class D5_DC12(D5, NamedTuple('DC12', [('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC12({_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC12) and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC13(D5, NamedTuple('DC13', [('cf21', Any), ('cf22', Any), ('cf23', Any), ('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC13({_dafny.string_of(self.cf21)}, {_dafny.string_of(self.cf22)}, {self.cf23.VerbatimString(True)}, {_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC13) and self.cf21 == __o.cf21 and self.cf22 == __o.cf22 and self.cf23 == __o.cf23 and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC11(D5, NamedTuple('DC11', [('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC11({_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC11) and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC14(D5, NamedTuple('DC14', [('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC14({_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC14) and self.cf25 == __o.cf25
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

class D6_DC16(D6, NamedTuple('DC16', [('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC16({_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC16) and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC15(D6, NamedTuple('DC15', [('cf26', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC15({_dafny.string_of(self.cf26)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC15) and self.cf26 == __o.cf26
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC18(False, False, int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D7_DC18)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D7_DC19)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D7_DC17)

class D7_DC18(D7, NamedTuple('DC18', [('cf29', Any), ('cf30', Any), ('cf31', Any), ('cf32', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC18({_dafny.string_of(self.cf29)}, {_dafny.string_of(self.cf30)}, {_dafny.string_of(self.cf31)}, {_dafny.string_of(self.cf32)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC18) and self.cf29 == __o.cf29 and self.cf30 == __o.cf30 and self.cf31 == __o.cf31 and self.cf32 == __o.cf32
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC19(D7, NamedTuple('DC19', [('cf33', Any), ('cf34', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC19({_dafny.string_of(self.cf33)}, {_dafny.string_of(self.cf34)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC19) and self.cf33 == __o.cf33 and self.cf34 == __o.cf34
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC17(D7, NamedTuple('DC17', [('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC17({_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC17) and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC21(False, False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D8_DC21)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D8_DC20)

class D8_DC21(D8, NamedTuple('DC21', [('cf36', Any), ('cf37', Any), ('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC21({_dafny.string_of(self.cf36)}, {_dafny.string_of(self.cf37)}, {_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC21) and self.cf36 == __o.cf36 and self.cf37 == __o.cf37 and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC20(D8, NamedTuple('DC20', [('cf35', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC20({_dafny.string_of(self.cf35)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC20) and self.cf35 == __o.cf35
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D9_DC22)

class D9_DC22(D9, NamedTuple('DC22', [('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC22({_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC22) and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: None
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D10_DC23)

class D10_DC23(D10, NamedTuple('DC23', [('cf40', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC23({_dafny.string_of(self.cf40)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC23) and self.cf40 == __o.cf40
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC25(_dafny.Array(None, 0), None, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D11_DC25)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D11_DC26)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D11_DC27)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D11_DC24)

class D11_DC25(D11, NamedTuple('DC25', [('cf42', Any), ('cf43', Any), ('cf44', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC25({_dafny.string_of(self.cf42)}, {_dafny.string_of(self.cf43)}, {_dafny.string_of(self.cf44)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC25) and self.cf42 == __o.cf42 and self.cf43 == __o.cf43 and self.cf44 == __o.cf44
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC26(D11, NamedTuple('DC26', [('cf45', Any), ('cf46', Any), ('cf47', Any), ('cf48', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC26({_dafny.string_of(self.cf45)}, {_dafny.string_of(self.cf46)}, {_dafny.string_of(self.cf47)}, {self.cf48.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC26) and self.cf45 == __o.cf45 and self.cf46 == __o.cf46 and self.cf47 == __o.cf47 and self.cf48 == __o.cf48
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC27(D11, NamedTuple('DC27', [('cf49', Any), ('cf50', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC27({_dafny.string_of(self.cf49)}, {_dafny.string_of(self.cf50)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC27) and self.cf49 == __o.cf49 and self.cf50 == __o.cf50
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC24(D11, NamedTuple('DC24', [('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC24({_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC24) and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC29(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D12_DC29)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D12_DC28)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D12_DC30)

class D12_DC29(D12, NamedTuple('DC29', [('cf52', Any), ('cf53', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC29({self.cf52.VerbatimString(True)}, {_dafny.string_of(self.cf53)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC29) and self.cf52 == __o.cf52 and self.cf53 == __o.cf53
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC28(D12, NamedTuple('DC28', [('cf51', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC28({_dafny.string_of(self.cf51)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC28) and self.cf51 == __o.cf51
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
        return lambda: D13_DC32(int(0), False, False, False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D13_DC32)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D13_DC33)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D13_DC31)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D13_DC34)

class D13_DC32(D13, NamedTuple('DC32', [('cf56', Any), ('cf57', Any), ('cf58', Any), ('cf59', Any), ('cf60', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC32({_dafny.string_of(self.cf56)}, {_dafny.string_of(self.cf57)}, {_dafny.string_of(self.cf58)}, {_dafny.string_of(self.cf59)}, {_dafny.string_of(self.cf60)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC32) and self.cf56 == __o.cf56 and self.cf57 == __o.cf57 and self.cf58 == __o.cf58 and self.cf59 == __o.cf59 and self.cf60 == __o.cf60
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC33(D13, NamedTuple('DC33', [('cf61', Any), ('cf62', Any), ('cf63', Any), ('cf64', Any), ('cf65', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC33({_dafny.string_of(self.cf61)}, {_dafny.string_of(self.cf62)}, {_dafny.string_of(self.cf63)}, {_dafny.string_of(self.cf64)}, {_dafny.string_of(self.cf65)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC33) and self.cf61 == __o.cf61 and self.cf62 == __o.cf62 and self.cf63 == __o.cf63 and self.cf64 == __o.cf64 and self.cf65 == __o.cf65
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC31(D13, NamedTuple('DC31', [('cf55', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC31({_dafny.string_of(self.cf55)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC31) and self.cf55 == __o.cf55
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC34(D13, NamedTuple('DC34', [('cf66', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC34({_dafny.string_of(self.cf66)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC34) and self.cf66 == __o.cf66
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: D14_DC36(False, int(0), int(0), _dafny.Map({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D14_DC36)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D14_DC35)

class D14_DC36(D14, NamedTuple('DC36', [('cf68', Any), ('cf69', Any), ('cf70', Any), ('cf71', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC36({_dafny.string_of(self.cf68)}, {_dafny.string_of(self.cf69)}, {_dafny.string_of(self.cf70)}, {_dafny.string_of(self.cf71)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC36) and self.cf68 == __o.cf68 and self.cf69 == __o.cf69 and self.cf70 == __o.cf70 and self.cf71 == __o.cf71
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC35(D14, NamedTuple('DC35', [('cf67', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC35({_dafny.string_of(self.cf67)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC35) and self.cf67 == __o.cf67
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: None
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D15_DC37)

class D15_DC37(D15, NamedTuple('DC37', [('cf72', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC37({_dafny.string_of(self.cf72)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC37) and self.cf72 == __o.cf72
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: D16_DC39(int(0), _dafny.Map({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D16_DC39)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D16_DC38)

class D16_DC39(D16, NamedTuple('DC39', [('cf74', Any), ('cf75', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC39({_dafny.string_of(self.cf74)}, {_dafny.string_of(self.cf75)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC39) and self.cf74 == __o.cf74 and self.cf75 == __o.cf75
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC38(D16, NamedTuple('DC38', [('cf73', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC38({_dafny.string_of(self.cf73)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC38) and self.cf73 == __o.cf73
    def __hash__(self) -> int:
        return super().__hash__()


class D17:
    @classmethod
    def default(cls, ):
        return lambda: D17_DC41(int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC41(self) -> bool:
        return isinstance(self, D17_DC41)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D17_DC40)
    @property
    def is_DC42(self) -> bool:
        return isinstance(self, D17_DC42)

class D17_DC41(D17, NamedTuple('DC41', [('cf77', Any), ('cf78', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC41({_dafny.string_of(self.cf77)}, {_dafny.string_of(self.cf78)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC41) and self.cf77 == __o.cf77 and self.cf78 == __o.cf78
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC40(D17, NamedTuple('DC40', [('cf76', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC40({_dafny.string_of(self.cf76)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC40) and self.cf76 == __o.cf76
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC42(D17, NamedTuple('DC42', [('cf79', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC42({_dafny.string_of(self.cf79)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC42) and self.cf79 == __o.cf79
    def __hash__(self) -> int:
        return super().__hash__()


class D18:
    @classmethod
    def default(cls, ):
        return lambda: D18_DC44(_dafny.Set({}), int(0), int(0), False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC44(self) -> bool:
        return isinstance(self, D18_DC44)
    @property
    def is_DC43(self) -> bool:
        return isinstance(self, D18_DC43)

class D18_DC44(D18, NamedTuple('DC44', [('cf81', Any), ('cf82', Any), ('cf83', Any), ('cf84', Any), ('cf85', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC44({_dafny.string_of(self.cf81)}, {_dafny.string_of(self.cf82)}, {_dafny.string_of(self.cf83)}, {_dafny.string_of(self.cf84)}, {_dafny.string_of(self.cf85)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC44) and self.cf81 == __o.cf81 and self.cf82 == __o.cf82 and self.cf83 == __o.cf83 and self.cf84 == __o.cf84 and self.cf85 == __o.cf85
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC43(D18, NamedTuple('DC43', [('cf80', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC43({_dafny.string_of(self.cf80)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC43) and self.cf80 == __o.cf80
    def __hash__(self) -> int:
        return super().__hash__()


class D19:
    @classmethod
    def default(cls, ):
        return lambda: D19_DC46(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), int(0), int(0), int(0), _dafny.Set({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC46(self) -> bool:
        return isinstance(self, D19_DC46)
    @property
    def is_DC45(self) -> bool:
        return isinstance(self, D19_DC45)

class D19_DC46(D19, NamedTuple('DC46', [('cf87', Any), ('cf88', Any), ('cf89', Any), ('cf90', Any), ('cf91', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC46({self.cf87.VerbatimString(True)}, {_dafny.string_of(self.cf88)}, {_dafny.string_of(self.cf89)}, {_dafny.string_of(self.cf90)}, {_dafny.string_of(self.cf91)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC46) and self.cf87 == __o.cf87 and self.cf88 == __o.cf88 and self.cf89 == __o.cf89 and self.cf90 == __o.cf90 and self.cf91 == __o.cf91
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC45(D19, NamedTuple('DC45', [('cf86', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC45({_dafny.string_of(self.cf86)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC45) and self.cf86 == __o.cf86
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    def m0(self, globalState):
        pass


class T1:
    pass

class GlobalState:
    def  __init__(self):
        self.f0: int = int(0)
        self.f1: bool = False
        self.f2: int = int(0)
        self.f5: bool = False
        self._f3: bool = False
        self._f4: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5):
        (self).f0 = f0
        (self).f1 = f1
        (self).f2 = f2
        (self)._f3 = f3
        (self)._f4 = f4
        (self).f5 = f5

    @property
    def f3(self):
        return self._f3
    @property
    def f4(self):
        return self._f4

class C0:
    def  __init__(self):
        self._f13: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f13):
        (self)._f13 = f13

    def fm23(self, p0, p1, p2, p3, globalState):
        return default__.safeDivisionInt(-986, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sicrkla"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nyetvlcd")))))

    @property
    def f13(self):
        return self._f13

class C1(T1):
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    def ctor__(self):
        pass
        pass

    def fm16(self, p0, globalState):
        return not(not ((_dafny.CodePoint('d')) in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lfuq")))) or ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k'), _dafny.CodePoint('g'), _dafny.CodePoint('x'), _dafny.CodePoint('j'), _dafny.CodePoint('l')])) < (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j'), _dafny.CodePoint('i')]))))

    def fm17(self, p0, p1, globalState):
        return ((0) - ((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "igmm"))) if True else 913))) + (828)

    def fm20(self, p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_371_i0_ in range(default__.abs(699))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qodmkq")))

    def m6(self, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: bool = False
        d_372_v0_: int
        d_372_v0_ = 477
        (globalState).f2 = d_372_v0_
        hi3_ = d_372_v0_
        for d_373_i0_ in range(d_372_v0_, hi3_):
            d_374_v1_: _dafny.Map
            d_374_v1_ = _dafny.Map({65: d_373_i0_})
            d_375_v2_: bool
            d_375_v2_ = True
            d_376_v4_: _dafny.Seq
            d_376_v4_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({d_372_v0_})])
            d_377_v5_: _dafny.Seq
            d_377_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a"))
            d_378_v6_: str
            d_378_v6_ = _dafny.CodePoint('t')
            d_379_v7_: _dafny.Set
            d_379_v7_ = _dafny.Set({(self).fm17(-88, _dafny.Set({len((d_377_v5_).set(default__.safeIndex(d_372_v0_, len(d_377_v5_)), d_378_v6_))}), globalState)})
            d_380_v9_: _dafny.Map
            d_380_v9_ = _dafny.Map({17: d_375_v2_})
            d_381_v10_: _dafny.MultiSet
            d_381_v10_ = _dafny.MultiSet([len(d_380_v9_), d_372_v0_])
            d_382_v12_: _dafny.Array
            nw67_ = _dafny.Array(None, 14)
            nw67_[int(0)] = d_374_v1_
            def iife50_():
                coll22_ = _dafny.Map()
                compr_22_: int
                for compr_22_ in ((d_376_v4_)[default__.safeIndex(d_372_v0_, len(d_376_v4_))]).Elements:
                    d_383_v3_: int = compr_22_
                    if (d_383_v3_) in ((d_376_v4_)[default__.safeIndex(d_372_v0_, len(d_376_v4_))]):
                        coll22_[default__.safeModuloInt(d_383_v3_, d_373_i0_)] = d_373_i0_
                return _dafny.Map(coll22_)
            nw67_[int(1)] = (d_374_v1_ if d_375_v2_ else iife50_()
            )
            nw67_[int(2)] = _dafny.Map({d_372_v0_: d_372_v0_})
            nw67_[int(3)] = d_374_v1_
            nw67_[int(4)] = (d_374_v1_) | (_dafny.Map({d_372_v0_: len(d_379_v7_)}))
            nw67_[int(5)] = d_374_v1_
            def iife51_():
                coll23_ = _dafny.Map()
                compr_23_: int
                for compr_23_ in (d_381_v10_).Elements:
                    d_384_v8_: int = compr_23_
                    if (d_384_v8_) in (d_381_v10_):
                        coll23_[(d_384_v8_) * (d_373_i0_)] = len(_dafny.SeqWithoutIsStrInference([d_373_i0_]))
                return _dafny.Map(coll23_)
            nw67_[int(6)] = iife51_()
            
            nw67_[int(7)] = d_374_v1_
            nw67_[int(8)] = (d_374_v1_).set(d_372_v0_, d_372_v0_)
            nw67_[int(9)] = d_374_v1_
            nw67_[int(10)] = d_374_v1_
            nw67_[int(11)] = d_374_v1_
            nw67_[int(12)] = d_374_v1_
            def iife52_():
                coll24_ = _dafny.Map()
                compr_24_: int
                for compr_24_ in _dafny.IntegerRange(318, 205):
                    d_385_v11_: int = compr_24_
                    if ((318) <= (d_385_v11_)) and ((d_385_v11_) < (205)):
                        coll24_[(d_385_v11_) + (d_373_i0_)] = d_372_v0_
                return _dafny.Map(coll24_)
            nw67_[int(13)] = iife52_()
            
            d_382_v12_ = nw67_
            index58_ = default__.safeIndex(989, (d_382_v12_).length(0))
            (d_382_v12_)[index58_] = d_374_v1_
            index59_ = default__.safeIndex(989, (d_382_v12_).length(0))
            (d_382_v12_)[index59_] = d_374_v1_
            d_386_v13_: _dafny.Map
            d_386_v13_ = _dafny.Map({True: 754})
            d_386_v13_ = (d_386_v13_).set(not(d_375_v2_), d_372_v0_)
            d_387_v14_: _dafny.Array
            nw68_ = _dafny.Array(False, 8)
            d_387_v14_ = nw68_
            d_388_v15_: _dafny.Map
            d_388_v15_ = _dafny.Map({d_375_v2_: d_387_v14_})
            d_389_v16_: D0
            d_389_v16_ = D0_DC2(d_372_v0_, d_378_v6_, len(d_388_v15_), d_373_i0_, (0) - (len(default__.fm21(d_377_v5_, globalState))))
            d_378_v6_ = (d_389_v16_).cf6
            def iife53_():
                coll25_ = _dafny.Set()
                compr_25_: int
                for compr_25_ in _dafny.IntegerRange(873, 595):
                    d_390_v17_: int = compr_25_
                    if ((873) <= (d_390_v17_)) and ((d_390_v17_) < (595)):
                        coll25_ = coll25_.union(_dafny.Set([default__.safeDivisionInt(d_390_v17_, d_372_v0_)]))
                return _dafny.Set(coll25_)
            d_379_v7_ = ((d_379_v7_) - (default__.fm22(globalState))) - (iife53_()
            )
        d_391_v18_: _dafny.Array
        nw69_ = _dafny.Array(False, 25)
        d_391_v18_ = nw69_
        nw70_ = _dafny.Array(False, 19)
        d_391_v18_ = nw70_
        d_392_i1_: int
        d_392_i1_ = 0
        with _dafny.label("0"):
            while not(False):
                with _dafny.c_label("0"):
                    if (d_392_i1_) >= (100):
                        raise _dafny.Break("0")
                    d_392_i1_ = (d_392_i1_) + (1)
                    d_393_v19_: bool
                    d_393_v19_ = True
                    d_394_v20_: _dafny.MultiSet
                    d_394_v20_ = _dafny.MultiSet([d_393_v19_, d_393_v19_])
                    d_395_v21_: _dafny.Seq
                    d_395_v21_ = _dafny.SeqWithoutIsStrInference([d_393_v19_, d_393_v19_, (self).fm16(d_394_v20_, globalState)])
                    d_396_v22_: str
                    d_396_v22_ = _dafny.CodePoint('b')
                    d_397_v23_: _dafny.Map
                    d_397_v23_ = _dafny.Map({d_396_v22_: d_396_v22_})
                    (globalState).f5 = (d_395_v21_) == ((d_395_v21_).set(default__.safeIndex(len(d_397_v23_), len(d_395_v21_)), True))
                    d_398_v24_: _dafny.Seq
                    d_398_v24_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "whpm"))
                    d_398_v24_ = d_398_v24_
                    d_399_v25_: _dafny.MultiSet
                    d_399_v25_ = _dafny.MultiSet([default__.safeModuloInt(d_372_v0_, d_372_v0_)])
                    d_399_v25_ = d_399_v25_
                    d_400_v26_: D0
                    d_400_v26_ = D0_DC2(955, _dafny.CodePoint('y'), d_372_v0_, d_372_v0_, d_372_v0_)
                    d_401_v27_: _dafny.Map
                    d_401_v27_ = _dafny.Map({d_372_v0_: d_400_v26_})
                    d_402_v28_: _dafny.Map
                    d_402_v28_ = _dafny.Map({d_393_v19_: len(d_401_v27_)})
                    d_403_v29_: _dafny.Array
                    def lambda24_(d_404_v0_):
                        def lambda25_(d_405_i2_):
                            return (d_405_i2_) + (d_404_v0_)

                        return lambda25_

                    init12_ = lambda24_(d_372_v0_)
                    nw71_ = _dafny.Array(None, 20)
                    for i0_12_ in range(nw71_.length(0)):
                        nw71_[i0_12_] = init12_(i0_12_)
                    d_403_v29_ = nw71_
                    d_406_v30_: _dafny.Map
                    d_406_v30_ = _dafny.Map({_dafny.MultiSet([(0) - (len(d_402_v28_))]): d_403_v29_})
                    d_407_v31_: C0
                    nw72_ = C0()
                    nw72_.ctor__((d_406_v30_) == (d_406_v30_))
                    d_407_v31_ = nw72_
                    pass
            pass
        d_408_v32_: bool
        d_408_v32_ = False
        d_409_v33_: _dafny.Set
        d_409_v33_ = _dafny.Set({d_408_v32_})
        d_410_v34_: D6
        d_410_v34_ = D6_DC15(d_409_v33_)
        d_411_v35_: C0
        nw73_ = C0()
        nw73_.ctor__(((d_410_v34_).cf26).issubset(d_409_v33_))
        d_411_v35_ = nw73_
        d_412_v36_: _dafny.Seq
        d_412_v36_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "le"))
        d_413_v37_: _dafny.Map
        d_413_v37_ = _dafny.Map({d_408_v32_: (d_412_v36_) + (default__.fm2(d_412_v36_, globalState))})
        d_413_v37_ = (d_413_v37_).set(not(d_408_v32_), d_412_v36_)
        r0 = d_372_v0_
        d_414_v38_: str
        d_414_v38_ = _dafny.CodePoint('q')
        d_415_v39_: _dafny.Map
        d_415_v39_ = _dafny.Map({237: _dafny.Set({True})})
        d_416_v40_: D0
        d_416_v40_ = D0_DC0(d_415_v39_)
        d_417_v41_: _dafny.Seq
        d_417_v41_ = _dafny.SeqWithoutIsStrInference([d_408_v32_, (d_411_v35_).f13])
        d_418_v43_: _dafny.Array
        nw74_ = _dafny.Array(None, 19)
        nw74_[int(0)] = d_372_v0_
        nw74_[int(1)] = d_372_v0_
        def iife54_():
            coll26_ = _dafny.Set()
            compr_26_: int
            for compr_26_ in _dafny.IntegerRange(381, 786):
                d_419_v42_: int = compr_26_
                if ((381) <= (d_419_v42_)) and ((d_419_v42_) < (786)):
                    coll26_ = coll26_.union(_dafny.Set([(d_419_v42_) + (330)]))
            return _dafny.Set(coll26_)
        nw74_[int(2)] = (self).fm17(len(d_412_v36_), iife54_()
        , globalState)
        nw74_[int(3)] = d_372_v0_
        nw74_[int(4)] = d_372_v0_
        nw74_[int(5)] = d_372_v0_
        nw74_[int(6)] = d_372_v0_
        nw74_[int(7)] = d_372_v0_
        nw74_[int(8)] = d_372_v0_
        nw74_[int(9)] = d_372_v0_
        nw74_[int(10)] = d_372_v0_
        nw74_[int(11)] = d_372_v0_
        nw74_[int(12)] = d_372_v0_
        nw74_[int(13)] = d_372_v0_
        nw74_[int(14)] = d_372_v0_
        nw74_[int(15)] = len(d_409_v33_)
        nw74_[int(16)] = len(d_409_v33_)
        nw74_[int(17)] = d_372_v0_
        nw74_[int(18)] = len(d_412_v36_)
        d_418_v43_ = nw74_
        d_420_v44_: _dafny.Seq
        d_420_v44_ = _dafny.SeqWithoutIsStrInference([d_418_v43_, d_418_v43_, d_418_v43_, d_418_v43_])
        r1 = ((0) - ((len(d_412_v36_) if (d_411_v35_).f13 else 231))) + ((d_411_v35_).fm23(d_372_v0_, d_372_v0_, d_414_v38_, default__.fm1(d_416_v40_, (d_411_v35_).f13, d_372_v0_, (d_417_v41_)[default__.safeIndex(len(d_420_v44_), len(d_417_v41_))], globalState), globalState))
        r2 = not ((d_411_v35_).f13) or ((d_417_v41_)[default__.safeIndex(d_372_v0_, len(d_417_v41_))])
        return r0, r1, r2


class C2:
    def  __init__(self):
        self.f12: bool = False
        self._f11: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    def ctor__(self, f11, f12):
        (self)._f11 = f11
        (self).f12 = f12

    def m5(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        if (p1) and ((not(p1) if False else p1)):
            d_421_v0_: int
            d_421_v0_ = -405
            d_422_v1_: _dafny.MultiSet
            d_422_v1_ = _dafny.MultiSet([(803) - (d_421_v0_)])
            d_423_v2_: _dafny.Seq
            d_423_v2_ = _dafny.SeqWithoutIsStrInference([d_421_v0_, 432])
            d_424_v3_: _dafny.Seq
            d_424_v3_ = _dafny.SeqWithoutIsStrInference([d_422_v1_])
            d_422_v1_ = ((_dafny.MultiSet(d_423_v2_)).intersection(d_422_v1_)) | ((d_424_v3_)[default__.safeIndex(d_421_v0_, len(d_424_v3_))])
            d_425_v4_: _dafny.MultiSet
            d_425_v4_ = _dafny.MultiSet([self.f12, p0])
            d_426_v5_: _dafny.Set
            d_426_v5_ = _dafny.Set({d_425_v4_, d_425_v4_})
            if (len(d_426_v5_)) >= (d_421_v0_):
                index60_ = default__.safeIndex(782, (p2).length(0))
                (p2)[index60_] = p1
                index61_ = default__.safeIndex(782, (p2).length(0))
                rhs47_ = p0
                rhs48_ = p0
                rhs49_ = (True if (p1 if self.f12 else self.f12) else True)
                lhs40_ = self
                lhs41_ = p2
                lhs42_ = default__.safeIndex(782, (p2).length(0))
                lhs43_ = self
                lhs40_.f12 = rhs47_
                lhs41_[lhs42_] = rhs48_
                lhs43_.f12 = rhs49_
                d_427_v6_: _dafny.Set
                d_427_v6_ = _dafny.Set({d_421_v0_, (d_422_v1_).cardinality})
                (globalState).f2 = len((d_427_v6_) | ((d_427_v6_).intersection(d_427_v6_)))
                (globalState).f0 = 409
                d_428_v7_: _dafny.Map
                d_428_v7_ = _dafny.Map({False: d_421_v0_})
                d_429_v8_: _dafny.Map
                d_429_v8_ = _dafny.Map({True: (p2)[default__.safeIndex(782, (p2).length(0))]})
                rhs50_ = ((d_428_v7_) | (_dafny.Map({p1: d_421_v0_}))) | ((_dafny.Map({p1: len(d_429_v8_)})) | (d_428_v7_))
                rhs51_ = p0
                d_428_v7_ = rhs50_
                r0 = rhs51_
                d_421_v0_ = d_421_v0_
            elif True:
                d_430_v9_: C0
                nw75_ = C0()
                nw75_.ctor__(not(p0))
                d_430_v9_ = nw75_
                d_431_v10_: D5
                d_431_v10_ = D5_DC12(d_421_v0_)
                (globalState).f2 = (d_431_v10_).cf20
                (globalState).f0 = d_421_v0_
                d_432_v11_: _dafny.Seq
                d_432_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ts"))
                d_432_v11_ = (self).f11
                (globalState).f0 = d_421_v0_
            if self.f12:
                d_433_v12_: _dafny.Array
                nw76_ = _dafny.Array(int(0), 24)
                d_433_v12_ = nw76_
                index62_ = default__.safeIndex(714, (d_433_v12_).length(0))
                (d_433_v12_)[index62_] = default__.fm24(d_421_v0_, globalState)
                index63_ = default__.safeIndex(714, (d_433_v12_).length(0))
                (d_433_v12_)[index63_] = -795
                (globalState).f0 = default__.safeDivisionInt((d_421_v0_) - (d_421_v0_), d_421_v0_)
                (globalState).f0 = d_421_v0_
                d_434_v13_: _dafny.Set
                d_434_v13_ = _dafny.Set({((d_425_v4_)[p1] if (p1) in (d_425_v4_) else d_421_v0_), (d_433_v12_)[default__.safeIndex(714, (d_433_v12_).length(0))], (d_433_v12_)[default__.safeIndex(714, (d_433_v12_).length(0))]})
                d_435_v14_: _dafny.Map
                d_435_v14_ = _dafny.Map({d_425_v4_: d_421_v0_})
                d_436_v15_: str
                d_436_v15_ = _dafny.CodePoint('m')
                d_437_v16_: _dafny.Map
                d_437_v16_ = _dafny.Map({d_436_v15_: (d_433_v12_)[default__.safeIndex(714, (d_433_v12_).length(0))]})
                index64_ = default__.safeIndex(714, (d_433_v12_).length(0))
                (d_433_v12_)[index64_] = (len((d_434_v13_) | (d_434_v13_))) * (len((default__.fm25(len(d_435_v14_), (d_437_v16_).set(d_436_v15_, d_421_v0_), d_436_v15_, globalState) if not(p0) else d_423_v2_)))
                (globalState).f1 = ((d_433_v12_)[default__.safeIndex(714, (d_433_v12_).length(0))]) != ((0) - ((d_433_v12_)[default__.safeIndex(714, (d_433_v12_).length(0))]))
            elif True:
                d_438_v17_: D1
                d_438_v17_ = D1_DC5()
                d_439_v18_: str
                d_439_v18_ = _dafny.CodePoint('o')
                d_440_v19_: _dafny.Array
                nw77_ = _dafny.Array(None, 10)
                nw77_[int(0)] = d_439_v18_
                nw77_[int(1)] = ((self).f11)[default__.safeIndex(d_421_v0_, len((self).f11))]
                nw77_[int(2)] = d_439_v18_
                nw77_[int(3)] = d_439_v18_
                nw77_[int(4)] = d_439_v18_
                nw77_[int(5)] = d_439_v18_
                nw77_[int(6)] = d_439_v18_
                nw77_[int(7)] = d_439_v18_
                nw77_[int(8)] = d_439_v18_
                nw77_[int(9)] = d_439_v18_
                d_440_v19_ = nw77_
                index65_ = default__.safeIndex(458, (d_440_v19_).length(0))
                (d_440_v19_)[index65_] = d_439_v18_
                d_441_v20_: _dafny.Seq
                d_441_v20_ = _dafny.SeqWithoutIsStrInference([(d_425_v4_).set(False, default__.abs(d_421_v0_))])
                index66_ = default__.safeIndex(458, (d_440_v19_).length(0))
                rhs52_ = D1_DC5()
                rhs53_ = d_439_v18_
                rhs54_ = ((self).f11)[default__.safeIndex((0) - (d_421_v0_), len((self).f11))]
                rhs55_ = d_421_v0_
                rhs56_ = (d_421_v0_) > (len(d_441_v20_))
                lhs44_ = d_440_v19_
                lhs45_ = default__.safeIndex(458, (d_440_v19_).length(0))
                lhs46_ = globalState
                lhs47_ = self
                d_438_v17_ = rhs52_
                d_439_v18_ = rhs53_
                lhs44_[lhs45_] = rhs54_
                lhs46_.f2 = rhs55_
                lhs47_.f12 = rhs56_
                d_442_v21_: _dafny.Seq
                d_442_v21_ = _dafny.SeqWithoutIsStrInference([p1])
                rhs57_ = d_421_v0_
                rhs58_ = d_442_v21_
                lhs48_ = globalState
                lhs48_.f2 = rhs57_
                d_442_v21_ = rhs58_
                d_443_v22_: _dafny.Array
                nw78_ = _dafny.Array(int(0), 2)
                d_443_v22_ = nw78_
                d_443_v22_ = d_443_v22_
                (globalState).f1 = False
                d_423_v2_ = d_423_v2_
            d_444_v23_: str
            d_444_v23_ = _dafny.CodePoint('s')
            d_445_v24_: _dafny.Set
            d_445_v24_ = _dafny.Set({d_421_v0_})
            d_446_v25_: D0
            d_446_v25_ = D0_DC2(len(_dafny.Set({(d_425_v4_).cardinality})), d_444_v23_, d_421_v0_, d_421_v0_, len(d_445_v24_))
            pat_let_tv13_ = d_421_v0_
            d_447_v26_: D0
            def iife55_(_pat_let14_0):
                def iife56_(d_448_dt__update__tmp_h0_):
                    def iife57_(_pat_let15_0):
                        def iife58_(d_449_dt__update_hcf5_h0_):
                            return D0_DC2(d_449_dt__update_hcf5_h0_, (d_448_dt__update__tmp_h0_).cf6, (d_448_dt__update__tmp_h0_).cf7, (d_448_dt__update__tmp_h0_).cf8, (d_448_dt__update__tmp_h0_).cf9)
                        return iife58_(_pat_let15_0)
                    return iife57_(pat_let_tv13_)
                return iife56_(_pat_let14_0)
            d_447_v26_ = D0_DC0(default__.fm0(self.f12, p1, iife55_(d_446_v25_), globalState))
            d_450_v27_: _dafny.Set
            d_450_v27_ = _dafny.Set({self.f12, True, p0})
            d_451_v28_: _dafny.Map
            d_451_v28_ = _dafny.Map({len(_dafny.Map({d_421_v0_: 168})): d_450_v27_})
            d_447_v26_ = D0_DC0((d_451_v28_) | (d_451_v28_))
            d_452_v29_: _dafny.Map
            d_452_v29_ = _dafny.Map({p0: len(d_423_v2_)})
            d_452_v29_ = (d_452_v29_).set(not(p0), d_421_v0_)
        elif True:
            d_453_v30_: int
            d_453_v30_ = 12
            d_454_v31_: _dafny.Map
            d_454_v31_ = _dafny.Map({len(_dafny.Map({self.f12: p0})): d_453_v30_})
            d_455_v32_: _dafny.Array
            nw79_ = _dafny.Array(None, 7)
            nw79_[int(0)] = 914
            nw79_[int(1)] = len(d_454_v31_)
            nw79_[int(2)] = d_453_v30_
            nw79_[int(3)] = len((self).f11)
            nw79_[int(4)] = d_453_v30_
            nw79_[int(5)] = d_453_v30_
            nw79_[int(6)] = d_453_v30_
            d_455_v32_ = nw79_
            d_456_v33_: _dafny.Set
            d_456_v33_ = _dafny.Set({d_455_v32_})
            d_457_v34_: D5
            d_457_v34_ = D5_DC11(d_456_v33_)
            d_458_v35_: _dafny.Seq
            d_458_v35_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lwclh"))
            d_459_v36_: _dafny.Seq
            d_459_v36_ = _dafny.SeqWithoutIsStrInference([self.f12])
            d_460_v37_: _dafny.Map
            d_460_v37_ = _dafny.Map({d_453_v30_: default__.fm26(p1, globalState)})
            d_461_v38_: D0
            d_461_v38_ = D0_DC0(d_460_v37_)
            rhs59_ = d_457_v34_
            rhs60_ = default__.fm1((d_461_v38_ if (d_459_v36_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "losi"))), len(d_459_v36_))] else d_461_v38_), (False) or (p0), d_453_v30_, p1, globalState)
            rhs61_ = default__.fm2((self).f11, globalState)
            rhs62_ = p0
            rhs63_ = not((p3).cf2)
            lhs49_ = globalState
            lhs50_ = globalState
            d_457_v34_ = rhs59_
            lhs49_.f1 = rhs60_
            d_458_v35_ = rhs61_
            r0 = rhs62_
            lhs50_.f1 = rhs63_
            (globalState).f2 = default__.safeModuloInt(d_453_v30_, d_453_v30_)
            d_462_v39_: _dafny.Seq
            d_462_v39_ = _dafny.SeqWithoutIsStrInference([len(d_458_v35_)])
            d_463_v40_: D3
            d_463_v40_ = D3_DC8(D1_DC4(d_462_v39_), d_455_v32_, len((self).f11))
            d_464_v41_: D3
            d_464_v41_ = D3_DC9((d_463_v40_ if p0 else d_463_v40_))
            d_465_v42_: _dafny.Map
            d_465_v42_ = _dafny.Map({default__.fm27(globalState): d_462_v39_})
            d_466_v43_: _dafny.Seq
            d_466_v43_ = _dafny.SeqWithoutIsStrInference([d_459_v36_, d_459_v36_, _dafny.SeqWithoutIsStrInference([default__.fm1(D0_DC0(d_460_v37_), False, 842, p1, globalState), p1])])
            d_467_v44_: _dafny.Map
            d_467_v44_ = _dafny.Map({p1: self.f12})
            rhs64_ = d_464_v41_
            rhs65_ = ((d_465_v42_)[d_466_v43_] if (d_466_v43_) in (d_465_v42_) else _dafny.SeqWithoutIsStrInference([len(d_467_v44_)]))
            rhs66_ = (d_453_v30_) - ((0) - ((0) - ((d_453_v30_ if self.f12 else d_453_v30_))))
            d_464_v41_ = rhs64_
            d_462_v39_ = rhs65_
            d_453_v30_ = rhs66_
            r0 = self.f12
            d_454_v31_ = (d_454_v31_).set(d_453_v30_, d_453_v30_)
        d_468_v45_: int
        d_468_v45_ = 997
        d_469_v46_: _dafny.Seq
        d_469_v46_ = _dafny.SeqWithoutIsStrInference([d_468_v45_])
        d_469_v46_ = d_469_v46_
        if p1:
            d_470_v47_: _dafny.Seq
            d_470_v47_ = _dafny.SeqWithoutIsStrInference([not(self.f12)])
            (globalState).f0 = default__.safeDivisionInt((0) - ((d_468_v45_) - (len(d_470_v47_))), 921)
            r0 = self.f12
            d_471_v48_: _dafny.Array
            nw80_ = _dafny.Array(None, 23)
            d_471_v48_ = nw80_
            d_472_v49_: _dafny.Seq
            d_472_v49_ = _dafny.SeqWithoutIsStrInference([d_471_v48_, d_471_v48_])
            (globalState).f2 = len(d_472_v49_)
            (globalState).f1 = not(p0)
            d_473_v50_: C1
            nw81_ = C1()
            nw81_.ctor__()
            d_473_v50_ = nw81_
        elif True:
            d_474_v51_: C0
            nw82_ = C0()
            nw82_.ctor__(p0)
            d_474_v51_ = nw82_
            if ((self).f11) != ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cyofsjs"))) + ((self).f11)):
                index67_ = default__.safeIndex(889, (p2).length(0))
                (p2)[index67_] = (p1 if (d_474_v51_).f13 else p1)
                index68_ = default__.safeIndex(889, (p2).length(0))
                (p2)[index68_] = p0
                d_475_v52_: _dafny.MultiSet
                d_475_v52_ = _dafny.MultiSet([p0])
                d_476_v53_: _dafny.Set
                d_476_v53_ = _dafny.Set({p1})
                d_477_v54_: D0
                d_477_v54_ = D0_DC0(_dafny.Map({((d_475_v52_)[(d_474_v51_).f13] if ((d_474_v51_).f13) in (d_475_v52_) else len(d_469_v46_)): d_476_v53_}))
                d_478_v55_: _dafny.Map
                d_478_v55_ = _dafny.Map({(d_474_v51_).f13: -292})
                (globalState).f5 = default__.fm1(d_477_v54_, (d_474_v51_).f13, (len(d_478_v55_)) * (d_468_v45_), (d_474_v51_).f13, globalState)
                d_479_v56_: _dafny.Array
                nw83_ = _dafny.Array(None, 6)
                nw83_[int(0)] = p1
                nw83_[int(1)] = ((d_474_v51_).f13) in (d_475_v52_)
                nw83_[int(2)] = (d_474_v51_).f13
                nw83_[int(3)] = (d_474_v51_).f13
                nw83_[int(4)] = (p0 if (d_474_v51_).f13 else default__.fm1(d_477_v54_, self.f12, d_468_v45_, p1, globalState))
                nw83_[int(5)] = self.f12
                d_479_v56_ = nw83_
                d_480_v57_: _dafny.Array
                d_480_v57_ = d_479_v56_
                d_481_v58_: _dafny.Seq
                d_481_v58_ = _dafny.SeqWithoutIsStrInference([(d_480_v57_), d_479_v56_])
                d_479_v56_ = (d_481_v58_)[default__.safeIndex(d_468_v45_, len(d_481_v58_))]
                index69_ = default__.safeIndex(889, (p2).length(0))
                (p2)[index69_] = ((self.f12 if (d_474_v51_).f13 else (p2)[default__.safeIndex(889, (p2).length(0))]) if self.f12 else (d_474_v51_).f13)
                d_482_v59_: _dafny.Map
                d_482_v59_ = _dafny.Map({d_468_v45_: _dafny.Set({p0, (p2)[default__.safeIndex(889, (p2).length(0))], p1, p0})})
                (globalState).f1 = (p0 if default__.fm1(D0_DC0(d_482_v59_), self.f12, len(d_476_v53_), (d_474_v51_).f13, globalState) else (d_474_v51_).f13)
            elif True:
                d_483_v61_: _dafny.Map
                d_483_v61_ = _dafny.Map({self.f12: d_468_v45_})
                d_484_v62_: _dafny.Array
                nw84_ = _dafny.Array(None, 20)
                nw84_[int(0)] = (_dafny.SeqWithoutIsStrInference([d_468_v45_])) + (d_469_v46_)
                nw84_[int(1)] = d_469_v46_
                nw84_[int(2)] = (_dafny.SeqWithoutIsStrInference([d_468_v45_ for d_485_i0_ in range(default__.abs(637))])) + (d_469_v46_)
                nw84_[int(3)] = (_dafny.SeqWithoutIsStrInference([d_468_v45_ for d_486_i1_ in range(default__.abs(906))])) + (d_469_v46_)
                nw84_[int(4)] = d_469_v46_
                nw84_[int(5)] = (d_469_v46_) + (d_469_v46_)
                nw84_[int(6)] = (D1_DC4(d_469_v46_)).cf11
                def iife59_():
                    coll27_ = _dafny.Set()
                    compr_27_: int
                    for compr_27_ in (d_469_v46_).Elements:
                        d_487_v60_: int = compr_27_
                        if (d_487_v60_) in (d_469_v46_):
                            coll27_ = coll27_.union(_dafny.Set([(d_487_v60_) - ((0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_488_i2_ in range(default__.abs(-190))]))))]))
                    return _dafny.Set(coll27_)
                nw84_[int(7)] = (d_469_v46_).set(default__.safeIndex(d_468_v45_, len(d_469_v46_)), (d_469_v46_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([iife59_()
                ])), len(d_469_v46_))])
                nw84_[int(8)] = (d_469_v46_) + (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({d_483_v61_: (d_474_v51_).f13})), d_468_v45_, d_468_v45_]))
                nw84_[int(9)] = d_469_v46_
                nw84_[int(10)] = d_469_v46_
                nw84_[int(11)] = (d_469_v46_) + (d_469_v46_)
                nw84_[int(12)] = d_469_v46_
                nw84_[int(13)] = (d_469_v46_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "usfonr"))), len(d_469_v46_)), d_468_v45_)
                nw84_[int(14)] = (d_469_v46_ if (d_474_v51_).f13 else d_469_v46_)
                nw84_[int(15)] = _dafny.SeqWithoutIsStrInference([d_468_v45_])
                nw84_[int(16)] = d_469_v46_
                nw84_[int(17)] = d_469_v46_
                nw84_[int(18)] = d_469_v46_
                nw84_[int(19)] = (d_469_v46_ if (d_474_v51_).f13 else d_469_v46_)
                d_484_v62_ = nw84_
                index70_ = default__.safeIndex(178, (d_484_v62_).length(0))
                (d_484_v62_)[index70_] = (d_469_v46_) + (d_469_v46_)
                index71_ = default__.safeIndex(178, (d_484_v62_).length(0))
                (d_484_v62_)[index71_] = d_469_v46_
                d_489_v63_: D3
                d_489_v63_ = D3_DC7(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x")))
                d_490_v64_: _dafny.Array
                nw85_ = _dafny.Array(None, 15)
                nw85_[int(0)] = (self).f11
                nw85_[int(1)] = default__.fm2((self).f11, globalState)
                nw85_[int(2)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_491_i3_ in range(default__.abs(728))])
                nw85_[int(3)] = (self).f11
                nw85_[int(4)] = (self).f11
                nw85_[int(5)] = (self).f11
                nw85_[int(6)] = (self).f11
                nw85_[int(7)] = default__.fm2((self).f11, globalState)
                nw85_[int(8)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_492_i4_ in range(default__.abs(382))])
                nw85_[int(9)] = ((self).f11) + ((self).f11)
                nw85_[int(10)] = (self).f11
                nw85_[int(11)] = (self).f11
                nw85_[int(12)] = (d_489_v63_).cf13
                nw85_[int(13)] = (self).f11
                nw85_[int(14)] = (self).f11
                d_490_v64_ = nw85_
                d_490_v64_ = d_490_v64_
                d_493_v67_: _dafny.Seq
                d_493_v67_ = _dafny.SeqWithoutIsStrInference([True])
                d_494_v68_: _dafny.Set
                d_494_v68_ = _dafny.Set({d_468_v45_, d_468_v45_, 162, len(d_493_v67_)})
                d_495_v69_: _dafny.Seq
                d_495_v69_ = _dafny.SeqWithoutIsStrInference([d_494_v68_, d_494_v68_, d_494_v68_, d_494_v68_, d_494_v68_])
                d_496_v70_: str
                d_496_v70_ = _dafny.CodePoint('p')
                d_497_v71_: _dafny.Map
                d_497_v71_ = _dafny.Map({d_496_v70_: d_468_v45_})
                d_498_v72_: D5
                d_498_v72_ = D5_DC12(d_468_v45_)
                d_499_v73_: _dafny.Map
                d_499_v73_ = _dafny.Map({_dafny.Set({d_468_v45_, len((self).f11), len(d_497_v71_), d_468_v45_, d_468_v45_}): d_498_v72_})
                d_500_v74_: _dafny.Map
                d_500_v74_ = _dafny.Map({p1: p2})
                d_501_v76_: _dafny.Map
                d_501_v76_ = _dafny.Map({d_494_v68_: d_468_v45_})
                d_502_v77_: _dafny.Seq
                def iife60_():
                    coll28_ = _dafny.Map()
                    compr_28_: _dafny.Set
                    for compr_28_ in (d_501_v76_).keys.Elements:
                        d_503_v75_: _dafny.Set = compr_28_
                        if (d_503_v75_) in (d_501_v76_):
                            coll28_[d_503_v75_] = d_498_v72_
                    return _dafny.Map(coll28_)
                d_502_v77_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({_dafny.Set({d_468_v45_, len(d_500_v74_)}): d_498_v72_}), iife60_()
                , _dafny.Map({(d_495_v69_)[default__.safeIndex(d_468_v45_, len(d_495_v69_))]: d_498_v72_}), _dafny.Map({d_494_v68_: d_498_v72_})])
                d_504_v79_: _dafny.Set
                d_504_v79_ = _dafny.Set({d_494_v68_, d_494_v68_})
                pat_let_tv14_ = d_468_v45_
                d_505_v80_: _dafny.Array
                nw86_ = _dafny.Array(None, 17)
                def iife61_():
                    def iife63_():
                        coll31_ = _dafny.Map()
                        compr_31_: _dafny.Set
                        for compr_31_ in (d_495_v69_).Elements:
                            d_506_v66_: _dafny.Set = compr_31_
                            if (d_506_v66_) in (d_495_v69_):
                                coll31_[d_506_v66_] = p0
                        return _dafny.Map(coll31_)
                    coll29_ = _dafny.Map()
                    def iife62_():
                        coll30_ = _dafny.Map()
                        compr_30_: _dafny.Set
                        for compr_30_ in (d_495_v69_).Elements:
                            d_506_v66_: _dafny.Set = compr_30_
                            if (d_506_v66_) in (d_495_v69_):
                                coll30_[d_506_v66_] = p0
                        return _dafny.Map(coll30_)
                    compr_29_: _dafny.Set
                    for compr_29_ in (iife62_()
                    ).keys.Elements:
                        d_507_v65_: _dafny.Set = compr_29_
                        if (d_507_v65_) in (iife63_()
                        ):
                            coll29_[d_507_v65_] = D5_DC12(d_468_v45_)
                    return _dafny.Map(coll29_)
                nw86_[int(0)] = iife61_()
                
                nw86_[int(1)] = d_499_v73_
                nw86_[int(2)] = (d_502_v77_)[default__.safeIndex((0) - (d_468_v45_), len(d_502_v77_))]
                def iife64_():
                    coll32_ = _dafny.Map()
                    compr_32_: _dafny.Set
                    for compr_32_ in (d_504_v79_).Elements:
                        d_508_v78_: _dafny.Set = compr_32_
                        if (d_508_v78_) in (d_504_v79_):
                            coll32_[d_508_v78_] = d_498_v72_
                    return _dafny.Map(coll32_)
                nw86_[int(3)] = iife64_()
                
                nw86_[int(4)] = d_499_v73_
                nw86_[int(5)] = d_499_v73_
                nw86_[int(6)] = d_499_v73_
                nw86_[int(7)] = d_499_v73_
                nw86_[int(8)] = d_499_v73_
                nw86_[int(9)] = d_499_v73_
                nw86_[int(10)] = d_499_v73_
                nw86_[int(11)] = d_499_v73_
                nw86_[int(12)] = d_499_v73_
                nw86_[int(13)] = d_499_v73_
                nw86_[int(14)] = d_499_v73_
                def iife65_(_pat_let16_0):
                    def iife66_(d_509_dt__update__tmp_h1_):
                        def iife67_(_pat_let17_0):
                            def iife68_(d_510_dt__update_hcf20_h0_):
                                return D5_DC12(d_510_dt__update_hcf20_h0_)
                            return iife68_(_pat_let17_0)
                        return iife67_(pat_let_tv14_)
                    return iife66_(_pat_let16_0)
                nw86_[int(15)] = _dafny.Map({d_494_v68_: iife65_(d_498_v72_)})
                nw86_[int(16)] = d_499_v73_
                d_505_v80_ = nw86_
                index72_ = default__.safeIndex(511, (d_505_v80_).length(0))
                (d_505_v80_)[index72_] = d_499_v73_
                index73_ = default__.safeIndex(511, (d_505_v80_).length(0))
                (d_505_v80_)[index73_] = (d_499_v73_) | (_dafny.Map({d_494_v68_: d_498_v72_}))
                d_511_v81_: _dafny.Map
                d_511_v81_ = _dafny.Map({d_468_v45_: d_468_v45_})
                d_512_v82_: D1
                d_512_v82_ = D1_DC4(_dafny.SeqWithoutIsStrInference([len(d_511_v81_), d_468_v45_, len(d_493_v67_)]))
                d_513_v83_: _dafny.Array
                def lambda26_(d_514_v45_):
                    def lambda27_(d_515_i5_):
                        return default__.safeModuloInt(d_515_i5_, d_514_v45_)

                    return lambda27_

                init13_ = lambda26_(d_468_v45_)
                nw87_ = _dafny.Array(None, 8)
                for i0_13_ in range(nw87_.length(0)):
                    nw87_[i0_13_] = init13_(i0_13_)
                d_513_v83_ = nw87_
                d_516_v84_: D3
                d_516_v84_ = D3_DC8(d_512_v82_, d_513_v83_, len((self).f11))
                (globalState).f2 = ((default__.fm24((d_516_v84_).cf16, globalState) if p0 else d_468_v45_)) * (d_468_v45_)
                d_517_v85_: _dafny.Seq
                d_517_v85_ = _dafny.SeqWithoutIsStrInference([(self).f11])
                d_517_v85_ = (((d_517_v85_) + (d_517_v85_)).set(default__.safeIndex(d_468_v45_, len((d_517_v85_) + (d_517_v85_))), (self).f11)) + (d_517_v85_)
            d_518_v86_: str
            d_518_v86_ = _dafny.CodePoint('p')
            d_518_v86_ = d_518_v86_
            d_519_v87_: _dafny.Seq
            d_519_v87_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bdx"))
            d_520_v88_: D3
            d_520_v88_ = D3_DC7(d_519_v87_)
            d_519_v87_ = (d_520_v88_).cf13
            d_521_v90_: D0
            def iife69_():
                coll33_ = _dafny.Map()
                compr_33_: int
                for compr_33_ in _dafny.IntegerRange(-760, -629):
                    d_522_v89_: int = compr_33_
                    if ((-760) <= (d_522_v89_)) and ((d_522_v89_) < (-629)):
                        coll33_[default__.safeModuloInt(d_522_v89_, len(_dafny.Set({False, (d_474_v51_).f13})))] = _dafny.Set({False})
                return _dafny.Map(coll33_)
            d_521_v90_ = D0_DC0(iife69_()
)
            d_523_v91_: _dafny.MultiSet
            d_523_v91_ = _dafny.MultiSet([default__.fm1(d_521_v90_, not((d_474_v51_).f13), d_468_v45_, not((d_474_v51_).f13), globalState)])
            d_524_v92_: C1
            nw88_ = C1()
            nw88_.ctor__()
            d_524_v92_ = nw88_
            d_525_v93_: _dafny.Array
            nw89_ = _dafny.Array(None, 10)
            nw89_[int(0)] = d_468_v45_
            nw89_[int(1)] = d_468_v45_
            nw89_[int(2)] = ((_dafny.MultiSet([p0])) - (d_523_v91_)).cardinality
            nw89_[int(3)] = default__.safeDivisionInt(d_468_v45_, len((self).f11))
            nw89_[int(4)] = d_468_v45_
            nw89_[int(5)] = d_468_v45_
            nw89_[int(6)] = len(d_519_v87_)
            nw89_[int(7)] = default__.safeDivisionInt((0) - ((0) - (d_468_v45_)), d_468_v45_)
            nw89_[int(8)] = default__.safeModuloInt(len(_dafny.Set({d_524_v92_, d_524_v92_, d_524_v92_, d_524_v92_, d_524_v92_})), d_468_v45_)
            nw89_[int(9)] = len((d_520_v88_).cf13)
            d_525_v93_ = nw89_
            index74_ = default__.safeIndex(792, (d_525_v93_).length(0))
            (d_525_v93_)[index74_] = d_468_v45_
            d_526_v94_: _dafny.MultiSet
            d_526_v94_ = _dafny.MultiSet([(0) - (d_468_v45_)])
            d_527_v95_: D3
            d_527_v95_ = D3_DC8(D1_DC4(d_469_v46_), d_525_v93_, (d_526_v94_).cardinality)
            index75_ = default__.safeIndex(792, (d_525_v93_).length(0))
            rhs67_ = (d_527_v95_).cf16
            rhs68_ = p1
            lhs51_ = d_525_v93_
            lhs52_ = default__.safeIndex(792, (d_525_v93_).length(0))
            lhs53_ = globalState
            lhs51_[lhs52_] = rhs67_
            lhs53_.f5 = rhs68_
        d_528_v96_: _dafny.Array
        def lambda28_(d_529_p0_):
            def lambda29_(d_530_i7_):
                return default__.safeModuloInt(d_530_i7_, len(_dafny.Map({self.f12: d_529_p0_})))

            return lambda29_

        init14_ = lambda28_(p0)
        nw90_ = _dafny.Array(None, 20)
        for i0_14_ in range(nw90_.length(0)):
            nw90_[i0_14_] = init14_(i0_14_)
        d_528_v96_ = nw90_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_528_v96_).length(0)):
            d_531_i6_: int = guard_loop_1_
            if (True) and (((0) <= (d_531_i6_)) and ((d_531_i6_) < ((d_528_v96_).length(0)))):
                (d_528_v96_)[(d_531_i6_)] = (d_531_i6_) + (len(_dafny.Set({d_468_v45_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_532_i8_ in range(default__.abs(249))]))})))
        d_533_v97_: _dafny.Seq
        d_533_v97_ = _dafny.SeqWithoutIsStrInference([self.f12, p0, self.f12, False])
        d_534_v98_: _dafny.Set
        d_534_v98_ = _dafny.Set({d_533_v97_, d_533_v97_})
        if (d_534_v98_).issubset(d_534_v98_):
            d_535_v99_: str
            d_535_v99_ = _dafny.CodePoint('t')
            d_536_v100_: _dafny.Map
            d_536_v100_ = _dafny.Map({p0: d_468_v45_})
            d_537_v101_: D0
            d_537_v101_ = D0_DC2(d_468_v45_, d_535_v99_, d_468_v45_, d_468_v45_, ((d_536_v100_)[p0] if (p0) in (d_536_v100_) else d_468_v45_))
            d_538_v102_: D0
            d_538_v102_ = D0_DC3(d_537_v101_)
            d_539_v103_: _dafny.Set
            d_539_v103_ = _dafny.Set({not(not(self.f12)), p1})
            d_538_v102_ = (d_538_v102_ if (d_539_v103_).issubset(d_539_v103_) else d_538_v102_)
            d_540_v104_: C0
            nw91_ = C0()
            nw91_.ctor__(self.f12)
            d_540_v104_ = nw91_
            d_541_v105_: C1
            nw92_ = C1()
            nw92_.ctor__()
            d_541_v105_ = nw92_
            (self).f12 = self.f12
            d_535_v99_ = d_535_v99_
        elif True:
            d_542_v106_: _dafny.Set
            d_542_v106_ = _dafny.Set({d_468_v45_})
            d_543_v107_: D7
            d_543_v107_ = D7_DC17(d_533_v97_)
            d_544_v108_: _dafny.Map
            d_544_v108_ = _dafny.Map({d_468_v45_: d_542_v106_})
            d_545_v112_: _dafny.Array
            nw93_ = _dafny.Array(None, 27)
            nw93_[int(0)] = d_542_v106_
            nw93_[int(1)] = (d_542_v106_) | (d_542_v106_)
            nw93_[int(2)] = d_542_v106_
            nw93_[int(3)] = _dafny.Set({len((d_543_v107_).cf28)})
            def iife70_():
                coll34_ = _dafny.Set()
                compr_34_: int
                for compr_34_ in _dafny.IntegerRange(170, 858):
                    d_546_v109_: int = compr_34_
                    if ((170) <= (d_546_v109_)) and ((d_546_v109_) < (858)):
                        coll34_ = coll34_.union(_dafny.Set([default__.safeModuloInt(d_546_v109_, d_468_v45_)]))
                return _dafny.Set(coll34_)
            nw93_[int(4)] = ((d_544_v108_)[d_468_v45_] if (d_468_v45_) in (d_544_v108_) else iife70_()
            )
            nw93_[int(5)] = d_542_v106_
            nw93_[int(6)] = (d_542_v106_) | (d_542_v106_)
            nw93_[int(7)] = _dafny.Set({d_468_v45_})
            nw93_[int(8)] = (d_542_v106_).intersection(d_542_v106_)
            nw93_[int(9)] = d_542_v106_
            nw93_[int(10)] = (_dafny.Set({len(_dafny.Map({d_468_v45_: p0}))})).intersection(_dafny.Set({-132}))
            def iife71_():
                coll35_ = _dafny.Set()
                compr_35_: int
                for compr_35_ in _dafny.IntegerRange(790, 697):
                    d_547_v110_: int = compr_35_
                    if ((790) <= (d_547_v110_)) and ((d_547_v110_) < (697)):
                        coll35_ = coll35_.union(_dafny.Set([(d_547_v110_) + ((d_469_v46_)[default__.safeIndex(d_468_v45_, len(d_469_v46_))])]))
                return _dafny.Set(coll35_)
            nw93_[int(11)] = iife71_()
            
            nw93_[int(12)] = (d_542_v106_) | (d_542_v106_)
            nw93_[int(13)] = d_542_v106_
            nw93_[int(14)] = d_542_v106_
            nw93_[int(15)] = d_542_v106_
            nw93_[int(16)] = d_542_v106_
            def iife72_():
                coll36_ = _dafny.Set()
                compr_36_: int
                for compr_36_ in _dafny.IntegerRange(-491, 281):
                    d_548_v111_: int = compr_36_
                    if ((-491) <= (d_548_v111_)) and ((d_548_v111_) < (281)):
                        coll36_ = coll36_.union(_dafny.Set([default__.safeDivisionInt(d_548_v111_, len(_dafny.Map({len(d_469_v46_): d_468_v45_})))]))
                return _dafny.Set(coll36_)
            nw93_[int(17)] = (iife72_()
             if p0 else d_542_v106_)
            nw93_[int(18)] = (default__.fm22(globalState)) | (d_542_v106_)
            nw93_[int(19)] = d_542_v106_
            nw93_[int(20)] = d_542_v106_
            nw93_[int(21)] = d_542_v106_
            nw93_[int(22)] = d_542_v106_
            nw93_[int(23)] = d_542_v106_
            nw93_[int(24)] = d_542_v106_
            nw93_[int(25)] = _dafny.Set({d_468_v45_, len(d_533_v97_), d_468_v45_, d_468_v45_})
            nw93_[int(26)] = d_542_v106_
            d_545_v112_ = nw93_
            index76_ = default__.safeIndex(740, (d_545_v112_).length(0))
            (d_545_v112_)[index76_] = d_542_v106_
            index77_ = default__.safeIndex(740, (d_545_v112_).length(0))
            def iife73_():
                coll37_ = _dafny.Set()
                compr_37_: int
                for compr_37_ in (default__.fm22(globalState)).Elements:
                    d_549_v113_: int = compr_37_
                    if (d_549_v113_) in (default__.fm22(globalState)):
                        coll37_ = coll37_.union(_dafny.Set([(d_549_v113_) + ((D0_DC1(-491, not(True), len(_dafny.SeqWithoutIsStrInference([True])), False)).cf3)]))
                return _dafny.Set(coll37_)
            (d_545_v112_)[index77_] = (d_542_v106_) | (iife73_()
            )
            (globalState).f1 = p0
            index78_ = default__.safeIndex(445, (d_528_v96_).length(0))
            (d_528_v96_)[index78_] = len((d_533_v97_) + (d_533_v97_))
            d_550_v114_: _dafny.Seq
            d_550_v114_ = _dafny.SeqWithoutIsStrInference([((self).f11) + ((self).f11)])
            index79_ = default__.safeIndex(445, (d_528_v96_).length(0))
            (d_528_v96_)[index79_] = len(((d_550_v114_)[default__.safeIndex(default__.safeDivisionInt(d_468_v45_, d_468_v45_), len(d_550_v114_))]).set(default__.safeIndex(d_468_v45_, len((d_550_v114_)[default__.safeIndex(default__.safeDivisionInt(d_468_v45_, d_468_v45_), len(d_550_v114_))])), default__.fm28((d_469_v46_)[default__.safeIndex(d_468_v45_, len(d_469_v46_))], (0) - (default__.fm24(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xvcevpcui"))), globalState)), self.f12, globalState)))
            d_551_v115_: C0
            nw94_ = C0()
            nw94_.ctor__(False)
            d_551_v115_ = nw94_
            index80_ = default__.safeIndex(251, (p2).length(0))
            (p2)[index80_] = (d_533_v97_) <= (d_533_v97_)
            index81_ = default__.safeIndex(251, (p2).length(0))
            (p2)[index81_] = (d_551_v115_).f13
        d_552_v116_: D6
        d_552_v116_ = D6_DC16(d_468_v45_)
        (globalState).f2 = ((default__.fm29(p0, d_468_v45_, globalState) if self.f12 else d_552_v116_)).cf27
        r0 = p1
        return r0

    @property
    def f11(self):
        return self._f11

class C3(T0, T1):
    def  __init__(self):
        self._f6: int = int(0)
        self._f9: int = int(0)
        self._f10: _dafny.Seq = _dafny.Seq({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f6(self):
        return self._f6
    def ctor__(self, f9, f10, f6):
        (self)._f9 = f9
        (self)._f10 = f10
        (self)._f6 = f6

    def fm16(self, p0, globalState):
        return ((self).f6) != ((self).f6)

    def fm17(self, p0, p1, globalState):
        return default__.safeModuloInt((self).f9, len((_dafny.Set({not(True)})) | (_dafny.Set({False}))))

    def fm18(self, p0, p1, p2, globalState):
        return ((_dafny.Map({True: True})) | (_dafny.Map({True: True}))) | (_dafny.Map({False: True}))

    def m0(self, globalState):
        r0: int = int(0)
        r1: bool = False
        r2: bool = False
        d_553_v0_: _dafny.Array
        nw95_ = _dafny.Array(int(0), 24)
        d_553_v0_ = nw95_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_553_v0_).length(0)):
            d_554_i0_: int = guard_loop_2_
            if (True) and (((0) <= (d_554_i0_)) and ((d_554_i0_) < ((d_553_v0_).length(0)))):
                (d_553_v0_)[(d_554_i0_)] = (d_554_i0_) - (default__.safeModuloInt(-943, (self).f6))
        d_555_v1_: D5
        d_555_v1_ = D5_DC11(_dafny.Set({d_553_v0_, d_553_v0_, d_553_v0_}))
        d_556_v2_: _dafny.Set
        d_556_v2_ = _dafny.Set({d_553_v0_, d_553_v0_, d_553_v0_})
        d_557_v3_: _dafny.Seq
        d_557_v3_ = _dafny.SeqWithoutIsStrInference([True, ((self).f9) > ((self).f6), (d_556_v2_).ispropersubset((d_555_v1_).cf19)])
        d_557_v3_ = d_557_v3_
        d_558_v4_: bool
        d_558_v4_ = True
        d_559_v5_: str
        d_559_v5_ = _dafny.CodePoint('v')
        d_560_v6_: _dafny.Map
        d_560_v6_ = _dafny.Map({27: d_559_v5_})
        d_561_v7_: _dafny.Map
        d_561_v7_ = _dafny.Map({(self).f9: d_558_v4_})
        d_562_i1_: int
        d_562_i1_ = 0
        with _dafny.label("1"):
            pat_let_tv15_ = d_558_v4_
            pat_let_tv16_ = d_557_v3_
            pat_let_tv17_ = d_557_v3_
            pat_let_tv18_ = d_558_v4_
            pat_let_tv19_ = d_558_v4_
            pat_let_tv20_ = d_558_v4_
            def lambda30_(source12_):
                if source12_.is_DC1:
                    d_568___mcc_h0_ = source12_.cf1
                    d_569___mcc_h1_ = source12_.cf2
                    d_570___mcc_h2_ = source12_.cf3
                    d_571___mcc_h3_ = source12_.cf4
                    d_572_cf4_ = d_571___mcc_h3_
                    d_573_cf3_ = d_570___mcc_h2_
                    d_574_cf2_ = d_569___mcc_h1_
                    d_575_cf1_ = d_568___mcc_h0_
                    return d_574_cf2_
                elif source12_.is_DC2:
                    d_576___mcc_h4_ = source12_.cf5
                    d_577___mcc_h5_ = source12_.cf6
                    d_578___mcc_h6_ = source12_.cf7
                    d_579___mcc_h7_ = source12_.cf8
                    d_580___mcc_h8_ = source12_.cf9
                    d_581_cf9_ = d_580___mcc_h8_
                    d_582_cf8_ = d_579___mcc_h7_
                    d_583_cf7_ = d_578___mcc_h6_
                    d_584_cf6_ = d_577___mcc_h5_
                    d_585_cf5_ = d_576___mcc_h4_
                    return pat_let_tv15_
                elif source12_.is_DC0:
                    d_586___mcc_h9_ = source12_.cf0
                    d_587_cf0_ = d_586___mcc_h9_
                    return (pat_let_tv16_) != (pat_let_tv17_)
                elif True:
                    d_588___mcc_h10_ = source12_.cf10
                    d_589_cf10_ = d_588___mcc_h10_
                    if pat_let_tv18_:
                        return pat_let_tv19_
                    elif True:
                        return pat_let_tv20_

            while lambda30_(default__.fm19(d_558_v4_, ((d_560_v6_)[(self).f6] if ((self).f6) in (d_560_v6_) else d_559_v5_), d_561_v7_, True, globalState)):
                with _dafny.c_label("1"):
                    if (d_562_i1_) >= (100):
                        raise _dafny.Break("1")
                    d_562_i1_ = (d_562_i1_) + (1)
                    d_563_v8_: _dafny.Map
                    d_563_v8_ = _dafny.Map({d_558_v4_: 969})
                    d_564_v9_: _dafny.Seq
                    d_564_v9_ = _dafny.SeqWithoutIsStrInference([d_563_v8_])
                    d_564_v9_ = d_564_v9_
                    d_565_v10_: C1
                    nw96_ = C1()
                    nw96_.ctor__()
                    d_565_v10_ = nw96_
                    d_566_v11_: D7
                    d_566_v11_ = D7_DC18(not(d_558_v4_), not(d_558_v4_), (self).f6, True)
                    (globalState).f2 = (0) - ((d_566_v11_).cf31)
                    d_567_v12_: C1
                    nw97_ = C1()
                    nw97_.ctor__()
                    d_567_v12_ = nw97_
                    pass
            pass
        d_590_v13_: _dafny.Array
        nw98_ = _dafny.Array(_dafny.Seq({}), 14)
        d_590_v13_ = nw98_
        d_591_v14_: _dafny.Seq
        d_591_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ftx"))
        d_592_v15_: _dafny.Seq
        d_592_v15_ = _dafny.SeqWithoutIsStrInference([d_591_v14_])
        index82_ = default__.safeIndex(672, (d_590_v13_).length(0))
        (d_590_v13_)[index82_] = (d_592_v15_).set(default__.safeIndex((self).f9, len(d_592_v15_)), d_591_v14_)
        index83_ = default__.safeIndex(672, (d_590_v13_).length(0))
        (d_590_v13_)[index83_] = _dafny.SeqWithoutIsStrInference([d_591_v14_ for d_593_i2_ in range(default__.abs(359))])
        d_594_v16_: C0
        nw99_ = C0()
        nw99_.ctor__(d_558_v4_)
        d_594_v16_ = nw99_
        hi4_ = ((self).f9) + ((self).f9)
        for d_595_i3_ in range((self).f9, hi4_):
            d_591_v14_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ubvcc"))) + (d_591_v14_)
            r0 = (self).f6
            d_596_v17_: _dafny.Array
            def lambda31_(d_597_v5_):
                def lambda32_(d_598_i4_):
                    return _dafny.SeqWithoutIsStrInference([d_597_v5_ for d_599_i5_ in range(default__.abs(-871))])

                return lambda32_

            init15_ = lambda31_(d_559_v5_)
            nw100_ = _dafny.Array(None, 29)
            for i0_15_ in range(nw100_.length(0)):
                nw100_[i0_15_] = init15_(i0_15_)
            d_596_v17_ = nw100_
            index84_ = default__.safeIndex(335, (d_596_v17_).length(0))
            (d_596_v17_)[index84_] = d_591_v14_
            index85_ = default__.safeIndex(335, (d_596_v17_).length(0))
            (d_596_v17_)[index85_] = (d_591_v14_) + (_dafny.SeqWithoutIsStrInference([d_559_v5_ for d_600_i6_ in range(default__.abs(190))]))
            d_601_v18_: _dafny.Set
            d_601_v18_ = _dafny.Set({d_595_i3_, d_595_i3_, default__.safeDivisionInt(d_595_i3_, d_595_i3_)})
            rhs69_ = (not((d_594_v16_).f13) if (d_594_v16_).f13 else (d_594_v16_).f13)
            rhs70_ = d_601_v18_
            rhs71_ = d_553_v0_
            lhs54_ = globalState
            lhs54_.f1 = rhs69_
            d_601_v18_ = rhs70_
            d_553_v0_ = rhs71_
        r0 = ((self).f6 if d_558_v4_ else (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jlx")))))
        d_602_v20_: _dafny.Set
        d_602_v20_ = _dafny.Set({_dafny.CodePoint('g'), d_559_v5_})
        d_603_v21_: _dafny.Seq
        d_603_v21_ = _dafny.SeqWithoutIsStrInference([d_602_v20_, d_602_v20_])
        def iife74_():
            coll38_ = _dafny.Set()
            compr_38_: str
            for compr_38_ in (_dafny.Map({d_559_v5_: (self).f6})).keys.Elements:
                d_604_v19_: str = compr_38_
                if (d_604_v19_) in (_dafny.Map({d_559_v5_: (self).f6})):
                    coll38_ = coll38_.union(_dafny.Set([d_604_v19_]))
            return _dafny.Set(coll38_)
        r1 = (_dafny.SeqWithoutIsStrInference([iife74_()
 for d_605_i7_ in range(default__.abs(712))])) <= (d_603_v21_)
        d_606_v22_: _dafny.Set
        d_606_v22_ = _dafny.Set({True})
        d_607_v23_: _dafny.Map
        d_607_v23_ = _dafny.Map({(self).f9: d_606_v22_})
        r2 = (((self).f9) <= (len(d_560_v6_))) == (default__.fm1(D0_DC0(d_607_v23_), (d_594_v16_).f13, (self).f9, (d_594_v16_).f13, globalState))
        return r0, r1, r2

    @property
    def f9(self):
        return self._f9
    @property
    def f10(self):
        return self._f10

class C4(T0):
    def  __init__(self):
        self._f6: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    @property
    def f6(self):
        return self._f6
    def ctor__(self, f6):
        (self)._f6 = f6

    def fm11(self, globalState):
        def iife75_():
            coll39_ = _dafny.Set()
            compr_39_: _dafny.Seq
            for compr_39_ in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wcshiaqcw"))])).Elements:
                d_610_v0_: _dafny.Seq = compr_39_
                if (d_610_v0_) in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wcshiaqcw"))])):
                    coll39_ = coll39_.union(_dafny.Set([d_610_v0_]))
            return _dafny.Set(coll39_)
        return ((_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tp")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_608_i0_ in range(default__.abs(-574))])})).intersection(_dafny.Set({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_609_i1_ in range(default__.abs(622))])}))).isdisjoint((iife75_()
        ) | (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qhbibqr")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jiutwlf")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tmcd"))})))

    def m0(self, globalState):
        r0: int = int(0)
        r1: bool = False
        r2: bool = False
        d_611_v0_: _dafny.Array
        nw101_ = _dafny.Array(False, 15)
        d_611_v0_ = nw101_
        index86_ = default__.safeIndex(880, (d_611_v0_).length(0))
        (d_611_v0_)[index86_] = False
        index87_ = default__.safeIndex(880, (d_611_v0_).length(0))
        (d_611_v0_)[index87_] = False
        d_612_v1_: str
        d_612_v1_ = _dafny.CodePoint('c')
        d_613_v2_: _dafny.Seq
        d_613_v2_ = _dafny.SeqWithoutIsStrInference([d_612_v1_, _dafny.CodePoint('a')])
        d_614_v3_: _dafny.Array
        nw102_ = _dafny.Array(None, 28)
        nw102_[int(0)] = (d_612_v1_ if (d_611_v0_)[default__.safeIndex(880, (d_611_v0_).length(0))] else d_612_v1_)
        nw102_[int(1)] = (d_613_v2_)[default__.safeIndex(-956, len(d_613_v2_))]
        nw102_[int(2)] = d_612_v1_
        nw102_[int(3)] = d_612_v1_
        nw102_[int(4)] = d_612_v1_
        nw102_[int(5)] = d_612_v1_
        nw102_[int(6)] = d_612_v1_
        nw102_[int(7)] = d_612_v1_
        nw102_[int(8)] = d_612_v1_
        nw102_[int(9)] = (d_613_v2_)[default__.safeIndex(953, len(d_613_v2_))]
        nw102_[int(10)] = d_612_v1_
        nw102_[int(11)] = d_612_v1_
        nw102_[int(12)] = d_612_v1_
        nw102_[int(13)] = d_612_v1_
        nw102_[int(14)] = d_612_v1_
        nw102_[int(15)] = d_612_v1_
        nw102_[int(16)] = d_612_v1_
        nw102_[int(17)] = d_612_v1_
        nw102_[int(18)] = d_612_v1_
        nw102_[int(19)] = d_612_v1_
        nw102_[int(20)] = d_612_v1_
        nw102_[int(21)] = (_dafny.CodePoint('y') if (d_611_v0_)[default__.safeIndex(880, (d_611_v0_).length(0))] else d_612_v1_)
        nw102_[int(22)] = d_612_v1_
        nw102_[int(23)] = d_612_v1_
        nw102_[int(24)] = default__.fm12(globalState)
        nw102_[int(25)] = d_612_v1_
        nw102_[int(26)] = _dafny.CodePoint('s')
        nw102_[int(27)] = d_612_v1_
        d_614_v3_ = nw102_
        d_615_v4_: _dafny.Set
        d_615_v4_ = _dafny.Set({(d_611_v0_)[default__.safeIndex(880, (d_611_v0_).length(0))]})
        d_616_v5_: D0
        d_616_v5_ = D0_DC1((self).f6, ((d_611_v0_)[default__.safeIndex(880, (d_611_v0_).length(0))]) and ((d_611_v0_)[default__.safeIndex(880, (d_611_v0_).length(0))]), (self).f6, (d_615_v4_).ispropersubset(d_615_v4_))
        d_617_v6_: D3
        d_617_v6_ = D3_DC7(d_613_v2_)
        pat_let_tv21_ = d_613_v2_
        d_618_v7_: D0
        def iife76_(_pat_let18_0):
            def iife77_(d_619_dt__update__tmp_h0_):
                def iife78_(_pat_let19_0):
                    def iife79_(d_620_dt__update_hcf13_h0_):
                        return D3_DC7(d_620_dt__update_hcf13_h0_)
                    return iife79_(_pat_let19_0)
                return iife78_(pat_let_tv21_)
            return iife77_(_pat_let18_0)
        d_618_v7_ = D0_DC2(542, d_612_v1_, (self).f6, (self).f6, len((iife76_(d_617_v6_)).cf13))
        rhs72_ = (d_614_v3_ if (d_611_v0_)[default__.safeIndex(880, (d_611_v0_).length(0))] else d_614_v3_)
        rhs73_ = (d_618_v7_).cf8
        rhs74_ = default__.fm13(d_612_v1_, d_613_v2_, globalState)
        lhs55_ = globalState
        d_614_v3_ = rhs72_
        lhs55_.f2 = rhs73_
        d_616_v5_ = rhs74_
        index88_ = default__.safeIndex(880, (d_611_v0_).length(0))
        (d_611_v0_)[index88_] = (d_611_v0_)[default__.safeIndex(880, (d_611_v0_).length(0))]
        hi5_ = (self).f6
        for d_621_i0_ in range((self).f6, hi5_):
            if (d_611_v0_)[default__.safeIndex(880, (d_611_v0_).length(0))]:
                d_622_v8_: D0
                d_622_v8_ = D0_DC0(_dafny.Map({d_621_i0_: d_615_v4_}))
                r2 = not (not (default__.fm1(d_622_v8_, (d_611_v0_)[default__.safeIndex(880, (d_611_v0_).length(0))], d_621_i0_, not((d_611_v0_)[default__.safeIndex(880, (d_611_v0_).length(0))]), globalState)) or ((d_611_v0_)[default__.safeIndex(880, (d_611_v0_).length(0))])) or (not((d_611_v0_)[default__.safeIndex(880, (d_611_v0_).length(0))]))
                d_623_v9_: _dafny.Array
                def lambda33_(d_624_i0_):
                    def lambda34_(d_625_i1_):
                        return (d_625_i1_) + (d_624_i0_)

                    return lambda34_

                init16_ = lambda33_(d_621_i0_)
                nw103_ = _dafny.Array(None, 5)
                for i0_16_ in range(nw103_.length(0)):
                    nw103_[i0_16_] = init16_(i0_16_)
                d_623_v9_ = nw103_
                def lambda35_(d_626_i2_):
                    return (d_626_i2_) - ((self).f6)

                init17_ = lambda35_
                nw104_ = _dafny.Array(None, 14)
                for i0_17_ in range(nw104_.length(0)):
                    nw104_[i0_17_] = init17_(i0_17_)
                d_623_v9_ = nw104_
                d_612_v1_ = (d_618_v7_).cf6
                d_627_v11_: _dafny.Array
                def lambda36_(d_628_v0_):
                    def lambda37_(d_629_i3_):
                        def iife80_():
                            coll40_ = _dafny.Set()
                            compr_40_: _dafny.Map
                            for compr_40_ in (_dafny.SeqWithoutIsStrInference([_dafny.Map({(d_628_v0_)[default__.safeIndex(880, (d_628_v0_).length(0))]: (d_628_v0_)[default__.safeIndex(880, (d_628_v0_).length(0))]})])).Elements:
                                d_630_v10_: _dafny.Map = compr_40_
                                if (d_630_v10_) in (_dafny.SeqWithoutIsStrInference([_dafny.Map({(d_628_v0_)[default__.safeIndex(880, (d_628_v0_).length(0))]: (d_628_v0_)[default__.safeIndex(880, (d_628_v0_).length(0))]})])):
                                    coll40_ = coll40_.union(_dafny.Set([d_630_v10_]))
                            return _dafny.Set(coll40_)
                        return iife80_()
                        

                    return lambda37_

                init18_ = lambda36_(d_611_v0_)
                nw105_ = _dafny.Array(None, 22)
                for i0_18_ in range(nw105_.length(0)):
                    nw105_[i0_18_] = init18_(i0_18_)
                d_627_v11_ = nw105_
                d_631_v12_: _dafny.Map
                d_631_v12_ = _dafny.Map({(d_611_v0_)[default__.safeIndex(880, (d_611_v0_).length(0))]: (d_611_v0_)[default__.safeIndex(880, (d_611_v0_).length(0))]})
                d_632_v13_: _dafny.Seq
                d_632_v13_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({d_631_v12_, d_631_v12_}), _dafny.Set({_dafny.Map({(d_611_v0_)[default__.safeIndex(880, (d_611_v0_).length(0))]: True})})])
                d_633_v14_: _dafny.MultiSet
                d_633_v14_ = _dafny.MultiSet([(self).f6])
                d_634_v15_: _dafny.Seq
                d_634_v15_ = _dafny.SeqWithoutIsStrInference([(d_633_v14_).cardinality, d_621_i0_])
                index89_ = default__.safeIndex(825, (d_627_v11_).length(0))
                (d_627_v11_)[index89_] = ((d_632_v13_)[default__.safeIndex((self).f6, len(d_632_v13_))]) | (_dafny.Set({default__.fm14(702, d_621_i0_, len(d_634_v15_), globalState)}))
                d_635_v16_: _dafny.Map
                d_635_v16_ = _dafny.Map({(d_611_v0_)[default__.safeIndex(880, (d_611_v0_).length(0))]: d_631_v12_})
                d_636_v17_: _dafny.Map
                d_636_v17_ = _dafny.Map({((d_635_v16_)[(d_611_v0_)[default__.safeIndex(880, (d_611_v0_).length(0))]] if ((d_611_v0_)[default__.safeIndex(880, (d_611_v0_).length(0))]) in (d_635_v16_) else d_631_v12_): d_621_i0_})
                index90_ = default__.safeIndex(825, (d_627_v11_).length(0))
                index91_ = default__.safeIndex(880, (d_611_v0_).length(0))
                index92_ = default__.safeIndex(880, (d_611_v0_).length(0))
                def iife81_():
                    coll41_ = _dafny.Set()
                    compr_41_: _dafny.Map
                    for compr_41_ in ((d_636_v17_) | (d_636_v17_)).keys.Elements:
                        d_637_v18_: _dafny.Map = compr_41_
                        if (d_637_v18_) in ((d_636_v17_) | (d_636_v17_)):
                            coll41_ = coll41_.union(_dafny.Set([d_637_v18_]))
                    return _dafny.Set(coll41_)
                rhs75_ = iife81_()

                rhs76_ = (self).fm11(globalState)
                rhs77_ = (d_616_v5_).cf2
                rhs78_ = (d_611_v0_)[default__.safeIndex(880, (d_611_v0_).length(0))]
                lhs56_ = d_627_v11_
                lhs57_ = default__.safeIndex(825, (d_627_v11_).length(0))
                lhs58_ = d_611_v0_
                lhs59_ = default__.safeIndex(880, (d_611_v0_).length(0))
                lhs60_ = d_611_v0_
                lhs61_ = default__.safeIndex(880, (d_611_v0_).length(0))
                lhs62_ = globalState
                lhs56_[lhs57_] = rhs75_
                lhs58_[lhs59_] = rhs76_
                lhs60_[lhs61_] = rhs77_
                lhs62_.f1 = rhs78_
                d_638_v19_: _dafny.Array
                nw106_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 22)
                d_638_v19_ = nw106_
                index93_ = default__.safeIndex(860, (d_638_v19_).length(0))
                (d_638_v19_)[index93_] = d_613_v2_
                index94_ = default__.safeIndex(860, (d_638_v19_).length(0))
                (d_638_v19_)[index94_] = (d_613_v2_) + (d_613_v2_)
            elif True:
                d_639_v20_: _dafny.Seq
                d_639_v20_ = _dafny.SeqWithoutIsStrInference([False, (d_611_v0_)[default__.safeIndex(880, (d_611_v0_).length(0))]])
                d_640_v21_: _dafny.Seq
                d_640_v21_ = _dafny.SeqWithoutIsStrInference([default__.fm15(d_621_i0_, True, globalState), len(d_639_v20_), (self).f6, d_621_i0_, (self).f6])
                d_641_v22_: int
                d_642_v23_: D1
                d_643_v24_: int
                d_644_v25_: _dafny.Set
                out0_: int
                out1_: D1
                out2_: int
                out3_: _dafny.Set
                out0_, out1_, out2_, out3_ = (self).m3(len(d_640_v21_), (d_611_v0_)[default__.safeIndex(880, (d_611_v0_).length(0))], d_621_i0_, (d_621_i0_) == (231), globalState)
                d_641_v22_ = out0_
                d_642_v23_ = out1_
                d_643_v24_ = out2_
                d_644_v25_ = out3_
                d_645_v26_: _dafny.Map
                d_645_v26_ = _dafny.Map({d_643_v24_: (self).f6})
                d_646_v27_: _dafny.Set
                d_646_v27_ = _dafny.Set({((d_645_v26_)[-238] if (-238) in (d_645_v26_) else len(d_640_v21_))})
                rhs79_ = len(((d_646_v27_) | (_dafny.Set({562})) if (d_611_v0_)[default__.safeIndex(880, (d_611_v0_).length(0))] else _dafny.Set({(self).f6, d_643_v24_})))
                rhs80_ = (d_611_v0_)[default__.safeIndex(880, (d_611_v0_).length(0))]
                rhs81_ = (d_611_v0_)[default__.safeIndex(880, (d_611_v0_).length(0))]
                rhs82_ = 370
                lhs63_ = globalState
                lhs64_ = globalState
                lhs65_ = globalState
                r0 = rhs79_
                lhs63_.f1 = rhs80_
                lhs64_.f1 = rhs81_
                lhs65_.f2 = rhs82_
                d_647_v28_: _dafny.MultiSet
                d_647_v28_ = _dafny.MultiSet([(d_611_v0_)[default__.safeIndex(880, (d_611_v0_).length(0))], (d_611_v0_)[default__.safeIndex(880, (d_611_v0_).length(0))]])
                d_648_v29_: _dafny.Array
                nw107_ = _dafny.Array(None, 12)
                nw107_[int(0)] = (0) - (d_643_v24_)
                nw107_[int(1)] = default__.fm15(d_643_v24_, (d_611_v0_)[default__.safeIndex(880, (d_611_v0_).length(0))], globalState)
                nw107_[int(2)] = len(d_646_v27_)
                nw107_[int(3)] = (self).f6
                nw107_[int(4)] = len((d_640_v21_) + (d_640_v21_))
                nw107_[int(5)] = (d_641_v22_ if (d_611_v0_)[default__.safeIndex(880, (d_611_v0_).length(0))] else (self).f6)
                nw107_[int(6)] = default__.safeModuloInt(d_643_v24_, d_621_i0_)
                nw107_[int(7)] = ((d_647_v28_)[(d_611_v0_)[default__.safeIndex(880, (d_611_v0_).length(0))]] if ((d_611_v0_)[default__.safeIndex(880, (d_611_v0_).length(0))]) in (d_647_v28_) else -667)
                nw107_[int(8)] = 962
                nw107_[int(9)] = d_643_v24_
                nw107_[int(10)] = d_641_v22_
                nw107_[int(11)] = (0) - (d_621_i0_)
                d_648_v29_ = nw107_
                index95_ = default__.safeIndex(220, (d_648_v29_).length(0))
                (d_648_v29_)[index95_] = default__.safeModuloInt((self).f6, d_621_i0_)
                index96_ = default__.safeIndex(220, (d_648_v29_).length(0))
                (d_648_v29_)[index96_] = default__.safeDivisionInt((d_641_v22_) - ((self).f6), (0) - ((d_616_v5_).cf1))
                d_612_v1_ = (d_613_v2_)[default__.safeIndex(d_621_i0_, len(d_613_v2_))]
                (globalState).f5 = (d_611_v0_)[default__.safeIndex(880, (d_611_v0_).length(0))]
            d_649_v30_: _dafny.Map
            d_649_v30_ = _dafny.Map({(d_611_v0_)[default__.safeIndex(880, (d_611_v0_).length(0))]: (self).f6})
            d_650_v31_: _dafny.Map
            d_650_v31_ = _dafny.Map({not((d_611_v0_)[default__.safeIndex(880, (d_611_v0_).length(0))]): (d_611_v0_)[default__.safeIndex(880, (d_611_v0_).length(0))]})
            d_651_v32_: _dafny.Map
            d_651_v32_ = d_650_v31_
            d_649_v30_ = (d_649_v30_).set((d_611_v0_)[default__.safeIndex(880, (d_611_v0_).length(0))], (0) - (len((d_650_v31_))))
            d_612_v1_ = d_612_v1_
            d_652_v33_: _dafny.Map
            d_652_v33_ = _dafny.Map({d_616_v5_: -138})
            d_652_v33_ = (d_652_v33_).set(d_616_v5_, (d_618_v7_).cf7)
        d_653_v34_: _dafny.Map
        d_653_v34_ = _dafny.Map({(d_611_v0_)[default__.safeIndex(880, (d_611_v0_).length(0))]: (d_611_v0_)[default__.safeIndex(880, (d_611_v0_).length(0))]})
        d_654_v35_: _dafny.Map
        d_654_v35_ = _dafny.Map({d_653_v34_: (self).fm11(globalState)})
        d_654_v35_ = (d_654_v35_).set((d_653_v34_).set((d_611_v0_)[default__.safeIndex(880, (d_611_v0_).length(0))], (d_611_v0_)[default__.safeIndex(880, (d_611_v0_).length(0))]), (d_611_v0_)[default__.safeIndex(880, (d_611_v0_).length(0))])
        d_655_v36_: C2
        nw108_ = C2()
        nw108_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iyanqn")), (d_611_v0_)[default__.safeIndex(880, (d_611_v0_).length(0))])
        d_655_v36_ = nw108_
        r0 = (self).f6
        r1 = True
        d_656_v37_: _dafny.MultiSet
        d_656_v37_ = _dafny.MultiSet([d_655_v36_.f12])
        d_657_v38_: D7
        d_657_v38_ = D7_DC19(d_656_v37_, (self).f6)
        r2 = (d_656_v37_) == ((d_657_v38_).cf33)
        return r0, r1, r2

    def m3(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: D1 = D1.default()()
        r2: int = int(0)
        r3: _dafny.Set = _dafny.Set({})
        (globalState).f5 = p3
        hi6_ = (self).f6
        for d_658_i0_ in range(p0, hi6_):
            d_659_v0_: _dafny.Array
            nw109_ = _dafny.Array(int(0), 14)
            d_659_v0_ = nw109_
            d_660_v1_: _dafny.Seq
            d_660_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gpng"))
            index97_ = default__.safeIndex(959, (d_659_v0_).length(0))
            (d_659_v0_)[index97_] = default__.safeDivisionInt((self).f6, len(d_660_v1_))
            index98_ = default__.safeIndex(959, (d_659_v0_).length(0))
            (d_659_v0_)[index98_] = (0) - (d_658_i0_)
            d_661_v2_: str
            d_661_v2_ = _dafny.CodePoint('p')
            d_662_v3_: _dafny.Map
            d_662_v3_ = _dafny.Map({d_661_v2_: len(_dafny.Set({p3}))})
            d_663_v4_: _dafny.Seq
            d_663_v4_ = _dafny.SeqWithoutIsStrInference([p3, False])
            d_664_v5_: _dafny.Map
            d_664_v5_ = _dafny.Map({len(d_663_v4_): -529})
            d_665_v6_: D7
            d_665_v6_ = D7_DC18(p3, p1, len(d_664_v5_), True)
            d_666_v7_: D3
            d_666_v7_ = D3_DC7(d_660_v1_)
            d_667_v8_: D5
            d_667_v8_ = D5_DC13((d_665_v6_).cf30, p3, (d_666_v7_).cf13, p0)
            d_668_v9_: C3
            nw110_ = C3()
            nw110_.ctor__(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "poee"))), default__.fm25(p2, d_662_v3_, d_661_v2_, globalState), (d_667_v8_).cf24)
            d_668_v9_ = nw110_
            d_669_v10_: _dafny.Map
            d_669_v10_ = _dafny.Map({688: d_663_v4_})
            d_669_v10_ = (d_669_v10_).set((d_659_v0_)[default__.safeIndex(959, (d_659_v0_).length(0))], d_663_v4_)
            d_670_v11_: _dafny.MultiSet
            d_670_v11_ = _dafny.MultiSet([p3, p3, p1, p3])
            d_671_v12_: _dafny.Seq
            d_671_v12_ = _dafny.SeqWithoutIsStrInference([d_670_v11_])
            d_672_v13_: _dafny.Array
            nw111_ = _dafny.Array(None, 11)
            nw111_[int(0)] = True
            nw111_[int(1)] = ((d_671_v12_)[default__.safeIndex((self).f6, len(d_671_v12_))]).isdisjoint(_dafny.MultiSet(d_663_v4_))
            nw111_[int(2)] = p3
            nw111_[int(3)] = p1
            nw111_[int(4)] = p1
            nw111_[int(5)] = (p1) or (p3)
            nw111_[int(6)] = p1
            nw111_[int(7)] = p3
            nw111_[int(8)] = p3
            nw111_[int(9)] = (not(False) if p1 else p3)
            nw111_[int(10)] = (d_660_v1_) == (d_660_v1_)
            d_672_v13_ = nw111_
            index99_ = default__.safeIndex(749, (d_672_v13_).length(0))
            (d_672_v13_)[index99_] = p1
            index100_ = default__.safeIndex(749, (d_672_v13_).length(0))
            (d_672_v13_)[index100_] = p1
        if True:
            (globalState).f5 = p1
            d_673_v14_: _dafny.Array
            nw112_ = _dafny.Array(False, 22)
            d_673_v14_ = nw112_
            d_674_v15_: _dafny.Seq
            d_674_v15_ = _dafny.SeqWithoutIsStrInference([d_673_v14_, d_673_v14_, d_673_v14_])
            d_674_v15_ = (d_674_v15_) + (d_674_v15_)
            d_675_v16_: D0
            d_675_v16_ = D0_DC1(default__.safeModuloInt(p0, 251), p1, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fr"))), p1)
            pat_let_tv22_ = p0
            def iife82_(_pat_let20_0):
                def iife83_(d_676_dt__update__tmp_h0_):
                    def iife84_(_pat_let21_0):
                        def iife85_(d_677_dt__update_hcf1_h0_):
                            def iife86_(_pat_let22_0):
                                def iife87_(d_678_dt__update_hcf4_h0_):
                                    def iife88_(_pat_let23_0):
                                        def iife89_(d_679_dt__update_hcf3_h0_):
                                            return D0_DC1(d_677_dt__update_hcf1_h0_, (d_676_dt__update__tmp_h0_).cf2, d_679_dt__update_hcf3_h0_, d_678_dt__update_hcf4_h0_)
                                        return iife89_(_pat_let23_0)
                                    return iife88_((self).f6)
                                return iife87_(_pat_let22_0)
                            return iife86_(True)
                        return iife85_(_pat_let21_0)
                    return iife84_(pat_let_tv22_)
                return iife83_(_pat_let20_0)
            d_675_v16_ = iife82_(d_675_v16_)
            d_680_v17_: _dafny.Array
            nw113_ = _dafny.Array(int(0), 4)
            d_680_v17_ = nw113_
            index101_ = default__.safeIndex(410, (d_680_v17_).length(0))
            (d_680_v17_)[index101_] = (-761) * ((self).f6)
            d_681_v18_: _dafny.Map
            d_681_v18_ = _dafny.Map({(0) - (p2): d_673_v14_})
            index102_ = default__.safeIndex(410, (d_680_v17_).length(0))
            rhs83_ = default__.safeDivisionInt(362, (p2) - (667))
            rhs84_ = ((d_681_v18_)[(self).f6] if ((self).f6) in (d_681_v18_) else d_673_v14_)
            rhs85_ = d_673_v14_
            rhs86_ = (self).f6
            lhs66_ = d_680_v17_
            lhs67_ = default__.safeIndex(410, (d_680_v17_).length(0))
            lhs68_ = globalState
            lhs66_[lhs67_] = rhs83_
            d_673_v14_ = rhs84_
            d_673_v14_ = rhs85_
            lhs68_.f2 = rhs86_
            d_682_v19_: _dafny.Seq
            d_682_v19_ = _dafny.SeqWithoutIsStrInference([p0])
            d_683_v20_: _dafny.Map
            d_683_v20_ = _dafny.Map({p3: p3})
            d_684_v21_: C3
            nw114_ = C3()
            nw114_.ctor__(p0, d_682_v19_, (len((D1_DC4(d_682_v19_)).cf11)) + (len(d_683_v20_)))
            d_684_v21_ = nw114_
        elif True:
            d_685_v22_: _dafny.Map
            d_685_v22_ = _dafny.Map({p1: p2})
            if ((d_685_v22_) | (d_685_v22_)) != (d_685_v22_):
                d_686_v23_: _dafny.Seq
                d_686_v23_ = _dafny.SeqWithoutIsStrInference([p0])
                d_687_v24_: _dafny.Seq
                d_687_v24_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aofqlqta"))
                d_688_v25_: T0
                nw115_ = C3()
                nw115_.ctor__((self).f6, d_686_v23_, len(d_687_v24_))
                d_688_v25_ = nw115_
                d_688_v25_ = d_688_v25_
                d_689_v26_: _dafny.Map
                d_689_v26_ = _dafny.Map({default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([len(d_685_v22_), len(_dafny.SeqWithoutIsStrInference([_dafny.Map({p0: p3}) for d_690_i1_ in range(default__.abs(256))]))])), (d_688_v25_).f6): p0})
                d_689_v26_ = (d_689_v26_) | (d_689_v26_)
                d_691_v27_: _dafny.Map
                d_691_v27_ = _dafny.Map({default__.safeDivisionInt(p0, (d_688_v25_).f6): not (p3) or (p3)})
                (globalState).f5 = ((d_691_v27_)[(0) - (len((d_687_v24_ if not(p3) else d_687_v24_)))] if ((0) - (len((d_687_v24_ if not(p3) else d_687_v24_)))) in (d_691_v27_) else (len(d_687_v24_)) < (p2))
                d_692_v28_: _dafny.Set
                d_692_v28_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cdiwerm"))})
                d_693_v29_: _dafny.Seq
                d_693_v29_ = _dafny.SeqWithoutIsStrInference([d_687_v24_, d_687_v24_])
                d_694_v30_: _dafny.Set
                d_694_v30_ = _dafny.Set({-403, (self).f6, (len(d_692_v28_)) * (len(_dafny.SeqWithoutIsStrInference([len(d_693_v29_), default__.fm24((d_688_v25_).f6, globalState)])))})
                d_694_v30_ = d_694_v30_
                (globalState).f5 = p1
            elif True:
                d_695_v31_: _dafny.Array
                nw116_ = _dafny.Array(int(0), 10)
                d_695_v31_ = nw116_
                d_696_v32_: _dafny.Map
                d_696_v32_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([default__.fm24(811, globalState)]): d_695_v31_})
                d_697_v33_: _dafny.Seq
                d_697_v33_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dmvccnxx"))
                d_698_v34_: _dafny.Seq
                d_698_v34_ = _dafny.SeqWithoutIsStrInference([(self).f6, len(d_697_v33_)])
                d_699_v35_: D8
                d_699_v35_ = D8_DC20(_dafny.Map({d_698_v34_: d_695_v31_}))
                d_700_v36_: D1
                d_700_v36_ = D1_DC4(d_698_v34_)
                d_701_v37_: _dafny.Array
                nw117_ = _dafny.Array(_dafny.Array(None, 0), 14)
                d_701_v37_ = nw117_
                d_702_v38_: _dafny.Map
                d_702_v38_ = _dafny.Map({d_701_v37_: d_696_v32_})
                pat_let_tv23_ = p2
                pat_let_tv24_ = d_695_v31_
                d_703_v39_: _dafny.Array
                nw118_ = _dafny.Array(None, 24)
                nw118_[int(0)] = d_696_v32_
                nw118_[int(1)] = d_696_v32_
                nw118_[int(2)] = (d_696_v32_) | (d_696_v32_)
                nw118_[int(3)] = d_696_v32_
                nw118_[int(4)] = _dafny.Map({_dafny.SeqWithoutIsStrInference([(self).f6]): d_695_v31_})
                def iife90_(_pat_let24_0):
                    def iife91_(d_704_dt__update__tmp_h1_):
                        def iife92_(_pat_let25_0):
                            def iife93_(d_706_dt__update_hcf35_h0_):
                                return D8_DC20(d_706_dt__update_hcf35_h0_)
                            return iife93_(_pat_let25_0)
                        return iife92_(_dafny.Map({_dafny.SeqWithoutIsStrInference([pat_let_tv23_ for d_705_i2_ in range(default__.abs(-35))]): pat_let_tv24_}))
                    return iife91_(_pat_let24_0)
                nw118_[int(5)] = (iife90_(d_699_v35_)).cf35
                nw118_[int(6)] = d_696_v32_
                nw118_[int(7)] = d_696_v32_
                nw118_[int(8)] = d_696_v32_
                nw118_[int(9)] = d_696_v32_
                nw118_[int(10)] = d_696_v32_
                nw118_[int(11)] = _dafny.Map({((d_700_v36_).cf11).set(default__.safeIndex(p2, len((d_700_v36_).cf11)), p0): d_695_v31_})
                nw118_[int(12)] = ((d_702_v38_)[d_701_v37_] if (d_701_v37_) in (d_702_v38_) else _dafny.Map({d_698_v34_: d_695_v31_}))
                nw118_[int(13)] = d_696_v32_
                nw118_[int(14)] = _dafny.Map({_dafny.SeqWithoutIsStrInference([(self).f6]): d_695_v31_})
                nw118_[int(15)] = d_696_v32_
                nw118_[int(16)] = d_696_v32_
                nw118_[int(17)] = d_696_v32_
                nw118_[int(18)] = (d_696_v32_) | (d_696_v32_)
                nw118_[int(19)] = d_696_v32_
                nw118_[int(20)] = d_696_v32_
                nw118_[int(21)] = (d_699_v35_).cf35
                nw118_[int(22)] = d_696_v32_
                nw118_[int(23)] = d_696_v32_
                d_703_v39_ = nw118_
                index103_ = default__.safeIndex(778, (d_703_v39_).length(0))
                (d_703_v39_)[index103_] = d_696_v32_
                d_707_v40_: _dafny.Map
                d_707_v40_ = _dafny.Map({default__.fm12(globalState): d_696_v32_})
                d_708_v41_: str
                d_708_v41_ = _dafny.CodePoint('v')
                index104_ = default__.safeIndex(778, (d_703_v39_).length(0))
                (d_703_v39_)[index104_] = (((d_707_v40_)[d_708_v41_] if (d_708_v41_) in (d_707_v40_) else d_696_v32_)).set(_dafny.SeqWithoutIsStrInference([(self).f6]), d_695_v31_)
                d_709_v42_: _dafny.Seq
                d_709_v42_ = _dafny.SeqWithoutIsStrInference([p3])
                (globalState).f0 = default__.safeModuloInt(len((d_709_v42_ if True else _dafny.SeqWithoutIsStrInference([p1]))), 435)
                (globalState).f1 = p3
                (globalState).f0 = p0
                r0 = p2
            r2 = p0
            d_710_v43_: _dafny.MultiSet
            d_710_v43_ = _dafny.MultiSet([p1, p3])
            d_711_v44_: _dafny.Seq
            d_711_v44_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "prilxjot"))
            d_712_v45_: _dafny.Map
            d_712_v45_ = _dafny.Map({d_711_v44_: p2})
            d_710_v43_ = default__.fm30((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xadt"))) + (d_711_v44_), d_712_v45_, p3, globalState)
            d_713_v46_: _dafny.Array
            def lambda38_(d_714_i3_):
                return default__.safeDivisionInt(d_714_i3_, (self).f6)

            init19_ = lambda38_
            nw119_ = _dafny.Array(None, 13)
            for i0_19_ in range(nw119_.length(0)):
                nw119_[i0_19_] = init19_(i0_19_)
            d_713_v46_ = nw119_
            index105_ = default__.safeIndex(823, (d_713_v46_).length(0))
            (d_713_v46_)[index105_] = (self).f6
            d_715_v47_: _dafny.MultiSet
            d_715_v47_ = _dafny.MultiSet([(self).f6])
            d_716_v48_: _dafny.Map
            d_716_v48_ = _dafny.Map({p0: p3})
            d_717_v49_: _dafny.Map
            d_717_v49_ = _dafny.Map({((d_715_v47_)[len(d_716_v48_)] if (len(d_716_v48_)) in (d_715_v47_) else (0) - (p0)): (self).f6})
            index106_ = default__.safeIndex(823, (d_713_v46_).length(0))
            (d_713_v46_)[index106_] = ((d_717_v49_)[(self).f6] if ((self).f6) in (d_717_v49_) else p2)
            if not(not (p1) or (p3)):
                d_718_v50_: T1
                nw120_ = C1()
                nw120_.ctor__()
                d_718_v50_ = nw120_
                d_719_v51_: _dafny.MultiSet
                d_719_v51_ = _dafny.MultiSet([d_718_v50_])
                d_720_v52_: _dafny.Map
                d_720_v52_ = _dafny.Map({((d_719_v51_)).cardinality: d_716_v48_})
                d_721_v53_: D5
                d_721_v53_ = D5_DC13(p1, not(p3), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_722_i4_ in range(default__.abs(478))]), (d_713_v46_)[default__.safeIndex(823, (d_713_v46_).length(0))])
                d_723_v54_: _dafny.Seq
                d_723_v54_ = _dafny.SeqWithoutIsStrInference([d_721_v53_])
                d_724_v56_: _dafny.Seq
                d_724_v56_ = _dafny.SeqWithoutIsStrInference([p2])
                d_725_v57_: _dafny.Seq
                d_725_v57_ = _dafny.SeqWithoutIsStrInference([len(d_724_v56_), 567])
                def iife94_():
                    coll42_ = _dafny.Map()
                    compr_42_: int
                    for compr_42_ in (d_725_v57_).Elements:
                        d_726_v55_: int = compr_42_
                        if (d_726_v55_) in (d_725_v57_):
                            coll42_[(d_726_v55_) * (p2)] = False
                    return _dafny.Map(coll42_)
                def iife95_():
                    coll43_ = _dafny.Map()
                    compr_43_: int
                    for compr_43_ in _dafny.IntegerRange(476, 297):
                        d_727_v58_: int = compr_43_
                        if ((476) <= (d_727_v58_)) and ((d_727_v58_) < (297)):
                            coll43_[(d_727_v58_) - (p0)] = p1
                    return _dafny.Map(coll43_)
                d_720_v52_ = (d_720_v52_).set(len(_dafny.Set({_dafny.SeqWithoutIsStrInference([d_721_v53_, D5_DC13(p1, p1, d_711_v44_, (self).f6)]), d_723_v54_})), (iife94_()
                 if p1 else iife95_()
                ))
                d_728_v59_: _dafny.Set
                d_728_v59_ = _dafny.Set({p1})
                d_717_v49_ = (d_717_v49_).set(len(d_728_v59_), p0)
                d_729_v60_: D6
                d_729_v60_ = D6_DC15((d_728_v59_ if p3 else _dafny.Set({True, p3})))
                d_730_v61_: _dafny.Seq
                d_730_v61_ = _dafny.SeqWithoutIsStrInference([p3])
                d_729_v60_ = default__.fm31((0) - ((d_713_v46_)[default__.safeIndex(823, (d_713_v46_).length(0))]), p1, p3, _dafny.MultiSet([d_730_v61_, d_730_v61_, d_730_v61_, d_730_v61_]), globalState)
                d_731_v62_: _dafny.Array
                nw121_ = _dafny.Array(False, 27)
                d_731_v62_ = nw121_
                index107_ = default__.safeIndex(166, (d_731_v62_).length(0))
                (d_731_v62_)[index107_] = p1
                index108_ = default__.safeIndex(166, (d_731_v62_).length(0))
                (d_731_v62_)[index108_] = p1
                d_732_v63_: _dafny.Array
                def lambda39_(d_733_v53_):
                    def lambda40_(d_734_i5_):
                        return d_733_v53_

                    return lambda40_

                init20_ = lambda39_(d_721_v53_)
                nw122_ = _dafny.Array(None, 9)
                for i0_20_ in range(nw122_.length(0)):
                    nw122_[i0_20_] = init20_(i0_20_)
                d_732_v63_ = nw122_
                index109_ = default__.safeIndex(35, (d_732_v63_).length(0))
                (d_732_v63_)[index109_] = d_721_v53_
                d_735_v64_: str
                d_735_v64_ = _dafny.CodePoint('r')
                pat_let_tv25_ = d_711_v44_
                pat_let_tv26_ = d_711_v44_
                pat_let_tv27_ = d_735_v64_
                index110_ = default__.safeIndex(35, (d_732_v63_).length(0))
                def iife96_(_pat_let26_0):
                    def iife97_(d_736_dt__update__tmp_h2_):
                        def iife98_(_pat_let27_0):
                            def iife99_(d_737_dt__update_hcf23_h0_):
                                return D5_DC13((d_736_dt__update__tmp_h2_).cf21, (d_736_dt__update__tmp_h2_).cf22, d_737_dt__update_hcf23_h0_, (d_736_dt__update__tmp_h2_).cf24)
                            return iife99_(_pat_let27_0)
                        return iife98_((pat_let_tv25_).set(default__.safeIndex(211, len(pat_let_tv26_)), pat_let_tv27_))
                    return iife97_(_pat_let26_0)
                rhs87_ = (p2 if p1 else len(_dafny.Set({(d_731_v62_)[default__.safeIndex(166, (d_731_v62_).length(0))], not(p1)})))
                rhs88_ = (d_731_v62_)[default__.safeIndex(166, (d_731_v62_).length(0))]
                rhs89_ = iife96_(d_721_v53_)
                lhs69_ = globalState
                lhs70_ = globalState
                lhs71_ = d_732_v63_
                lhs72_ = default__.safeIndex(35, (d_732_v63_).length(0))
                lhs69_.f0 = rhs87_
                lhs70_.f5 = rhs88_
                lhs71_[lhs72_] = rhs89_
            elif True:
                d_711_v44_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_738_i6_ in range(default__.abs(533))])
                d_739_v65_: _dafny.Map
                d_739_v65_ = _dafny.Map({default__.fm15(p0, p1, globalState): _dafny.Set({p3, not(p3), p1, p3})})
                d_740_v66_: D0
                d_740_v66_ = D0_DC0(d_739_v65_)
                (globalState).f5 = default__.fm1(d_740_v66_, (p1 if p3 else not(p1)), p0, (p1 if False else not(p3)), globalState)
                index111_ = default__.safeIndex(823, (d_713_v46_).length(0))
                (d_713_v46_)[index111_] = default__.safeDivisionInt((self).f6, (352) - ((d_713_v46_)[default__.safeIndex(823, (d_713_v46_).length(0))]))
                d_741_v67_: str
                d_741_v67_ = _dafny.CodePoint('w')
                d_742_v68_: D0
                d_742_v68_ = D0_DC2(5, d_741_v67_, (self).f6, p0, p2)
                d_743_v69_: D0
                d_743_v69_ = D0_DC3(d_742_v68_)
                d_743_v69_ = default__.fm32(globalState)
                d_744_v70_: _dafny.Seq
                d_744_v70_ = _dafny.SeqWithoutIsStrInference([(d_713_v46_)[default__.safeIndex(823, (d_713_v46_).length(0))]])
                d_744_v70_ = (default__.fm33(globalState)).cf11
        d_745_v71_: _dafny.Seq
        d_745_v71_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tvo"))
        d_745_v71_ = d_745_v71_
        hi7_ = (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_746_i8_ in range(default__.abs(421))]))) + (p0)
        for d_747_i7_ in range(default__.fm24((self).f6, globalState), hi7_):
            d_748_v72_: _dafny.Seq
            d_748_v72_ = _dafny.SeqWithoutIsStrInference([p2])
            d_749_v73_: _dafny.Array
            nw123_ = _dafny.Array(None, 14)
            nw123_[int(0)] = d_747_i7_
            nw123_[int(1)] = (0) - (((0) - (p0)) * ((d_748_v72_)[default__.safeIndex(p0, len(d_748_v72_))]))
            nw123_[int(2)] = d_747_i7_
            nw123_[int(3)] = (self).f6
            nw123_[int(4)] = p0
            nw123_[int(5)] = (d_748_v72_)[default__.safeIndex(p0, len(d_748_v72_))]
            nw123_[int(6)] = (self).f6
            nw123_[int(7)] = (0) - (len(default__.fm2(d_745_v71_, globalState)))
            nw123_[int(8)] = p2
            nw123_[int(9)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qqajrefo")))
            nw123_[int(10)] = len(d_745_v71_)
            nw123_[int(11)] = len(d_745_v71_)
            nw123_[int(12)] = default__.safeModuloInt(default__.fm15(p2, p1, globalState), (0) - (d_747_i7_))
            nw123_[int(13)] = 898
            d_749_v73_ = nw123_
            rhs90_ = p2
            rhs91_ = p2
            rhs92_ = d_749_v73_
            lhs73_ = globalState
            lhs74_ = globalState
            lhs73_.f2 = rhs90_
            lhs74_.f2 = rhs91_
            d_749_v73_ = rhs92_
            d_750_v74_: _dafny.Array
            nw124_ = _dafny.Array(_dafny.Map({}), 3)
            d_750_v74_ = nw124_
            d_751_v75_: _dafny.Map
            d_751_v75_ = _dafny.Map({_dafny.Set({p0, d_747_i7_, d_747_i7_, d_747_i7_, (0) - (p0)}): p1})
            index112_ = default__.safeIndex(539, (d_750_v74_).length(0))
            (d_750_v74_)[index112_] = d_751_v75_
            d_752_v76_: _dafny.Array
            nw125_ = _dafny.Array(False, 12)
            d_752_v76_ = nw125_
            d_753_v77_: _dafny.Set
            d_753_v77_ = _dafny.Set({d_752_v76_})
            d_754_v78_: D0
            d_754_v78_ = D0_DC1(117, p3, (self).f6, p1)
            index113_ = default__.safeIndex(539, (d_750_v74_).length(0))
            rhs93_ = (d_753_v77_).issubset(d_753_v77_)
            rhs94_ = d_751_v75_
            rhs95_ = (d_754_v78_).cf2
            rhs96_ = (0) - (p0)
            rhs97_ = p1
            lhs75_ = globalState
            lhs76_ = d_750_v74_
            lhs77_ = default__.safeIndex(539, (d_750_v74_).length(0))
            lhs78_ = globalState
            lhs79_ = globalState
            lhs75_.f5 = rhs93_
            lhs76_[lhs77_] = rhs94_
            lhs78_.f5 = rhs95_
            r0 = rhs96_
            lhs79_.f1 = rhs97_
            d_755_v79_: T0
            nw126_ = C3()
            nw126_.ctor__(p2, d_748_v72_, (self).f6)
            d_755_v79_ = nw126_
            d_756_v80_: _dafny.MultiSet
            d_756_v80_ = _dafny.MultiSet([d_755_v79_])
            rhs98_ = ((d_756_v80_).set(d_755_v79_, default__.abs((0) - (d_747_i7_)))).cardinality
            rhs99_ = (default__.fm24((self).f6, globalState)) - (p2)
            rhs100_ = 600
            rhs101_ = (d_755_v79_).f6
            lhs80_ = globalState
            lhs81_ = globalState
            lhs82_ = globalState
            lhs80_.f2 = rhs98_
            lhs81_.f0 = rhs99_
            lhs82_.f2 = rhs100_
            r2 = rhs101_
            d_757_v81_: str
            d_757_v81_ = _dafny.CodePoint('o')
            d_758_v82_: _dafny.Set
            d_758_v82_ = _dafny.Set({d_757_v81_, d_757_v81_})
            d_759_v83_: _dafny.Map
            d_759_v83_ = _dafny.Map({p1: p0})
            d_760_v84_: _dafny.Set
            d_760_v84_ = _dafny.Set({len(d_758_v82_), len(d_745_v71_), p0, len(d_759_v83_)})
            if (len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ngk"))))) <= (len(d_760_v84_)):
                d_761_v85_: _dafny.Map
                d_761_v85_ = _dafny.Map({p2: p1})
                d_762_v86_: _dafny.Map
                d_762_v86_ = _dafny.Map({d_761_v85_: d_752_v76_})
                d_763_v87_: _dafny.Map
                d_763_v87_ = _dafny.Map({(d_755_v79_).f6: (d_755_v79_).f6})
                d_764_v88_: C3
                nw127_ = C3()
                nw127_.ctor__((0) - ((d_755_v79_).f6), _dafny.SeqWithoutIsStrInference([p0, d_747_i7_, (d_755_v79_).f6]), ((d_763_v87_)[p2] if (p2) in (d_763_v87_) else p0))
                d_764_v88_ = nw127_
                index114_ = default__.safeIndex(895, (d_749_v73_).length(0))
                (d_749_v73_)[index114_] = len(d_759_v83_)
                index115_ = default__.safeIndex(895, (d_749_v73_).length(0))
                rhs102_ = (d_762_v86_) | (d_762_v86_)
                rhs103_ = d_749_v73_
                rhs104_ = d_764_v88_
                rhs105_ = default__.safeModuloInt(((0) - (len(d_761_v85_))) * ((d_755_v79_).f6), ((self).f6) - (p0))
                rhs106_ = default__.safeModuloInt((0) - (p2), ((d_764_v88_).f10)[default__.safeIndex((d_755_v79_).f6, len((d_764_v88_).f10))])
                lhs83_ = d_749_v73_
                lhs84_ = default__.safeIndex(895, (d_749_v73_).length(0))
                d_762_v86_ = rhs102_
                d_749_v73_ = rhs103_
                d_764_v88_ = rhs104_
                lhs83_[lhs84_] = rhs105_
                r2 = rhs106_
                d_765_v89_: _dafny.MultiSet
                d_765_v89_ = _dafny.MultiSet([p3, p3])
                (globalState).f2 = default__.safeModuloInt((0) - (d_747_i7_), ((d_765_v89_).cardinality) - (d_747_i7_))
                d_766_v90_: D8
                d_766_v90_ = D8_DC21(True, p3, (d_755_v79_).f6)
                d_767_v91_: C0
                nw128_ = C0()
                nw128_.ctor__((d_766_v90_).cf37)
                d_767_v91_ = nw128_
                d_767_v91_ = (d_767_v91_ if (p1 if p1 else p1) else d_767_v91_)
                d_768_v92_: C0
                nw129_ = C0()
                nw129_.ctor__(p1)
                d_768_v92_ = nw129_
                d_769_v93_: C3
                nw130_ = C3()
                nw130_.ctor__(-724, d_748_v72_, (self).f6)
                d_769_v93_ = nw130_
            elif True:
                d_749_v73_ = d_749_v73_
                d_770_v94_: _dafny.Array
                nw131_ = _dafny.Array(_dafny.MultiSet({}), 27)
                d_770_v94_ = nw131_
                d_771_v95_: _dafny.MultiSet
                d_771_v95_ = _dafny.MultiSet([(0) - ((d_755_v79_).f6)])
                index116_ = default__.safeIndex(857, (d_770_v94_).length(0))
                (d_770_v94_)[index116_] = d_771_v95_
                index117_ = default__.safeIndex(857, (d_770_v94_).length(0))
                (d_770_v94_)[index117_] = d_771_v95_
                d_772_v96_: _dafny.Set
                d_772_v96_ = _dafny.Set({(d_755_v79_)})
                d_773_v97_: _dafny.MultiSet
                d_773_v97_ = _dafny.MultiSet([p1])
                r3 = _dafny.Set({(d_772_v96_).ispropersubset(d_772_v96_), (d_773_v97_).isdisjoint(d_773_v97_)})
                d_774_v98_: D7
                d_774_v98_ = D7_DC18(not(p1), p1, d_747_i7_, p3)
                d_775_v99_: _dafny.Map
                d_775_v99_ = _dafny.Map({d_774_v98_: d_752_v76_})
                d_776_v100_: D11
                d_776_v100_ = D11_DC24(d_775_v99_)
                d_775_v99_ = ((d_776_v100_).cf41) | ((d_775_v99_) | (d_775_v99_))
                d_777_v101_: _dafny.Set
                d_777_v101_ = _dafny.Set({p1})
                d_778_v102_: D0
                d_778_v102_ = D0_DC0(_dafny.Map({p2: d_777_v101_}))
                d_779_v103_: _dafny.Seq
                d_779_v103_ = _dafny.SeqWithoutIsStrInference([p3])
                d_780_v104_: D3
                d_780_v104_ = D3_DC7(d_745_v71_)
                d_781_v105_: D12
                d_781_v105_ = D12_DC28(default__.fm34(globalState))
                rhs107_ = default__.fm1(d_778_v102_, p1, default__.fm15(388, (d_779_v103_)[default__.safeIndex(p2, len(d_779_v103_))], globalState), (not(p3)) or (True), globalState)
                rhs108_ = (d_757_v81_) not in ((d_780_v104_).cf13)
                rhs109_ = _dafny.SeqWithoutIsStrInference([d_757_v81_ for d_782_i9_ in range(default__.abs(601))])
                rhs110_ = (d_779_v103_)[default__.safeIndex(p0, len(d_779_v103_))]
                rhs111_ = (d_781_v105_).cf51
                lhs85_ = globalState
                lhs86_ = globalState
                lhs87_ = globalState
                lhs85_.f5 = rhs107_
                lhs86_.f1 = rhs108_
                d_745_v71_ = rhs109_
                lhs87_.f5 = rhs110_
                d_759_v83_ = rhs111_
        d_783_v106_: _dafny.Set
        d_783_v106_ = _dafny.Set({p1, p3})
        d_784_v107_: _dafny.Map
        d_784_v107_ = _dafny.Map({p2: d_783_v106_})
        d_785_v108_: D0
        d_785_v108_ = D0_DC0(d_784_v107_)
        d_786_v109_: _dafny.Set
        d_786_v109_ = _dafny.Set({default__.fm2(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "et")), globalState), d_745_v71_})
        d_787_v110_: _dafny.Seq
        d_787_v110_ = _dafny.SeqWithoutIsStrInference([p1])
        d_788_v111_: _dafny.Seq
        d_788_v111_ = _dafny.SeqWithoutIsStrInference([((d_787_v110_).set(default__.safeIndex(p0, len(d_787_v110_)), False)).set(default__.safeIndex(p0, len((d_787_v110_).set(default__.safeIndex(p0, len(d_787_v110_)), False))), p3)])
        def iife100_():
            coll44_ = _dafny.Set()
            compr_44_: _dafny.Seq
            for compr_44_ in (d_786_v109_).Elements:
                d_789_v112_: _dafny.Seq = compr_44_
                if (d_789_v112_) in (d_786_v109_):
                    coll44_ = coll44_.union(_dafny.Set([d_789_v112_]))
            return _dafny.Set(coll44_)
        rhs112_ = d_785_v108_
        rhs113_ = (d_788_v111_) < (_dafny.SeqWithoutIsStrInference([d_787_v110_]))
        rhs114_ = (iife100_()
        ) - ((d_786_v109_) - (d_786_v109_))
        rhs115_ = p3
        lhs88_ = globalState
        lhs89_ = globalState
        d_785_v108_ = rhs112_
        lhs88_.f5 = rhs113_
        d_786_v109_ = rhs114_
        lhs89_.f1 = rhs115_
        r0 = (self).f6
        r1 = D1_DC5()
        d_790_v113_: _dafny.Map
        d_790_v113_ = _dafny.Map({p0: p1})
        r2 = len(d_790_v113_)
        r3 = d_783_v106_
        return r0, r1, r2, r3

    def m4(self, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: int = int(0)
        r3: str = _dafny.CodePoint('D')
        d_791_i0_: int
        d_791_i0_ = 0
        with _dafny.label("2"):
            while False:
                with _dafny.c_label("2"):
                    if (d_791_i0_) >= (100):
                        raise _dafny.Break("2")
                    d_791_i0_ = (d_791_i0_) + (1)
                    r2 = (0) - ((self).f6)
                    d_792_v0_: bool
                    d_792_v0_ = False
                    d_793_v1_: _dafny.Seq
                    d_793_v1_ = _dafny.SeqWithoutIsStrInference([True, d_792_v0_, d_792_v0_, d_792_v0_, d_792_v0_])
                    d_794_v2_: D7
                    d_794_v2_ = D7_DC19((_dafny.MultiSet([False, d_792_v0_]) if d_792_v0_ else (_dafny.MultiSet((d_793_v1_).set(default__.safeIndex((self).f6, len(d_793_v1_)), d_792_v0_))).set(d_792_v0_, default__.abs((self).f6))), 415)
                    source13_ = d_794_v2_
                    if source13_.is_DC18:
                        d_795___mcc_h0_ = source13_.cf29
                        d_796___mcc_h1_ = source13_.cf30
                        d_797___mcc_h2_ = source13_.cf31
                        d_798___mcc_h3_ = source13_.cf32
                        d_799_cf32_ = d_798___mcc_h3_
                        d_800_cf31_ = d_797___mcc_h2_
                        d_801_cf30_ = d_796___mcc_h1_
                        d_802_cf29_ = d_795___mcc_h0_
                        d_803_v3_: _dafny.Set
                        d_803_v3_ = _dafny.Set({d_799_cf32_})
                        r1 = (len(_dafny.SeqWithoutIsStrInference([d_800_cf31_ for d_804_i1_ in range(default__.abs(454))]))) - ((len(d_803_v3_)) - ((0) - ((self).f6)))
                        d_805_v4_: _dafny.Seq
                        d_805_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mqmrskldp"))
                        d_806_v5_: C2
                        nw132_ = C2()
                        nw132_.ctor__(d_805_v4_, (d_801_cf30_ if d_799_cf32_ else d_801_cf30_))
                        d_806_v5_ = nw132_
                        (globalState).f2 = len(d_793_v1_)
                        d_807_v6_: D3
                        d_807_v6_ = D3_DC7(((d_806_v5_).f11) + ((d_806_v5_).f11))
                        d_808_v7_: _dafny.Array
                        nw133_ = _dafny.Array(int(0), 2)
                        d_808_v7_ = nw133_
                        d_809_v8_: _dafny.Array
                        def lambda41_(d_810_v5_):
                            def lambda42_(d_811_i2_):
                                return (_dafny.CodePoint('g') if d_810_v5_.f12 else _dafny.CodePoint('p'))

                            return lambda42_

                        init21_ = lambda41_(d_806_v5_)
                        nw134_ = _dafny.Array(None, 2)
                        for i0_21_ in range(nw134_.length(0)):
                            nw134_[i0_21_] = init21_(i0_21_)
                        d_809_v8_ = nw134_
                        d_812_v9_: str
                        d_812_v9_ = _dafny.CodePoint('b')
                        index118_ = default__.safeIndex(628, (d_809_v8_).length(0))
                        (d_809_v8_)[index118_] = d_812_v9_
                        d_813_v10_: _dafny.MultiSet
                        d_813_v10_ = _dafny.MultiSet([d_800_cf31_])
                        pat_let_tv28_ = d_805_v4_
                        index119_ = default__.safeIndex(628, (d_809_v8_).length(0))
                        def iife102_(_pat_let29_0):
                            def iife103_(d_814_dt__update__tmp_h0_):
                                def iife104_(_pat_let30_0):
                                    def iife105_(d_815_dt__update_hcf13_h0_):
                                        return D3_DC7(d_815_dt__update_hcf13_h0_)
                                    return iife105_(_pat_let30_0)
                                return iife104_(pat_let_tv28_)
                            return iife103_(_pat_let29_0)
                        def iife101_(_pat_let28_0):
                            def iife106_(d_816_dt__update__tmp_h1_):
                                def iife107_(_pat_let31_0):
                                    def iife108_(d_817_dt__update_hcf13_h1_):
                                        return D3_DC7(d_817_dt__update_hcf13_h1_)
                                    return iife108_(_pat_let31_0)
                                return iife107_(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cwj")))
                            return iife106_(_pat_let28_0)
                        rhs116_ = iife101_(iife102_(d_807_v6_))
                        rhs117_ = d_808_v7_
                        rhs118_ = d_812_v9_
                        rhs119_ = d_813_v10_
                        rhs120_ = not((d_812_v9_) in ((d_806_v5_).f11))
                        lhs90_ = d_809_v8_
                        lhs91_ = default__.safeIndex(628, (d_809_v8_).length(0))
                        d_807_v6_ = rhs116_
                        d_808_v7_ = rhs117_
                        lhs90_[lhs91_] = rhs118_
                        d_813_v10_ = rhs119_
                        d_799_cf32_ = rhs120_
                    elif source13_.is_DC19:
                        d_818___mcc_h4_ = source13_.cf33
                        d_819___mcc_h5_ = source13_.cf34
                        d_820_cf34_ = d_819___mcc_h5_
                        d_821_cf33_ = d_818___mcc_h4_
                        d_822_v11_: str
                        d_822_v11_ = _dafny.CodePoint('i')
                        d_823_v12_: _dafny.MultiSet
                        d_823_v12_ = _dafny.MultiSet([d_822_v11_])
                        d_824_v13_: _dafny.Map
                        d_824_v13_ = _dafny.Map({d_820_cf34_: d_792_v0_})
                        d_825_v14_: _dafny.Map
                        d_825_v14_ = _dafny.Map({default__.fm15(((d_823_v12_)[d_822_v11_] if (d_822_v11_) in (d_823_v12_) else d_820_cf34_), ((d_824_v13_)[(self).f6] if ((self).f6) in (d_824_v13_) else True), globalState): d_792_v0_})
                        d_826_v15_: D13
                        d_826_v15_ = D13_DC32((self).f6, d_792_v0_, d_792_v0_, d_792_v0_, (d_793_v1_)[default__.safeIndex(d_820_cf34_, len(d_793_v1_))])
                        d_827_v16_: _dafny.Seq
                        d_827_v16_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ohdkp"))
                        d_828_v17_: _dafny.Seq
                        d_828_v17_ = _dafny.SeqWithoutIsStrInference([(0) - (len(d_827_v16_))])
                        d_829_v18_: _dafny.Set
                        d_829_v18_ = _dafny.Set({d_792_v0_})
                        d_830_v19_: _dafny.Array
                        def lambda43_(d_831_v0_):
                            def lambda44_(d_832_i3_):
                                return d_831_v0_

                            return lambda44_

                        init22_ = lambda43_(d_792_v0_)
                        nw135_ = _dafny.Array(None, 4)
                        for i0_22_ in range(nw135_.length(0)):
                            nw135_[i0_22_] = init22_(i0_22_)
                        d_830_v19_ = nw135_
                        d_833_v20_: _dafny.Map
                        d_833_v20_ = _dafny.Map({d_830_v19_: d_825_v14_})
                        d_834_v21_: _dafny.MultiSet
                        d_834_v21_ = _dafny.MultiSet([d_820_cf34_])
                        d_835_v22_: _dafny.Array
                        nw136_ = _dafny.Array(None, 28)
                        nw136_[int(0)] = d_792_v0_
                        nw136_[int(1)] = d_792_v0_
                        nw136_[int(2)] = True
                        nw136_[int(3)] = not(d_792_v0_)
                        nw136_[int(4)] = not(d_792_v0_)
                        nw136_[int(5)] = False
                        nw136_[int(6)] = d_792_v0_
                        nw136_[int(7)] = (d_824_v13_) == (((D13_DC31(d_825_v14_)).cf55).set(d_820_cf34_, False))
                        nw136_[int(8)] = d_792_v0_
                        nw136_[int(9)] = (d_826_v15_).cf59
                        nw136_[int(10)] = (len(d_828_v17_)) > ((self).f6)
                        nw136_[int(11)] = d_792_v0_
                        nw136_[int(12)] = ((0) - (len(d_829_v18_))) > ((self).f6)
                        nw136_[int(13)] = not (d_792_v0_) or (d_792_v0_)
                        nw136_[int(14)] = (d_792_v0_) and (d_792_v0_)
                        nw136_[int(15)] = d_792_v0_
                        nw136_[int(16)] = d_792_v0_
                        nw136_[int(17)] = (_dafny.MultiSet([True, d_792_v0_, d_792_v0_, d_792_v0_, ((d_824_v13_)[len(d_833_v20_)] if (len(d_833_v20_)) in (d_824_v13_) else d_792_v0_)])).issubset(d_821_cf33_)
                        nw136_[int(18)] = d_792_v0_
                        nw136_[int(19)] = d_792_v0_
                        nw136_[int(20)] = (d_828_v17_) != (d_828_v17_)
                        nw136_[int(21)] = d_792_v0_
                        nw136_[int(22)] = d_792_v0_
                        nw136_[int(23)] = d_792_v0_
                        nw136_[int(24)] = d_792_v0_
                        nw136_[int(25)] = (d_820_cf34_) == ((self).f6)
                        nw136_[int(26)] = not((_dafny.MultiSet(d_828_v17_)).ispropersubset(d_834_v21_))
                        nw136_[int(27)] = not (False) or (d_792_v0_)
                        d_835_v22_ = nw136_
                        d_835_v22_ = d_830_v19_
                        d_836_v23_: _dafny.Array
                        nw137_ = _dafny.Array(None, 14)
                        nw137_[int(0)] = d_827_v16_
                        nw137_[int(1)] = d_827_v16_
                        nw137_[int(2)] = _dafny.SeqWithoutIsStrInference([d_822_v11_ for d_837_i4_ in range(default__.abs(896))])
                        nw137_[int(3)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "khfcae"))
                        nw137_[int(4)] = d_827_v16_
                        nw137_[int(5)] = d_827_v16_
                        nw137_[int(6)] = d_827_v16_
                        nw137_[int(7)] = d_827_v16_
                        nw137_[int(8)] = d_827_v16_
                        nw137_[int(9)] = d_827_v16_
                        nw137_[int(10)] = (d_827_v16_) + (d_827_v16_)
                        nw137_[int(11)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yareb"))
                        nw137_[int(12)] = d_827_v16_
                        nw137_[int(13)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aqkxf"))
                        d_836_v23_ = nw137_
                        index120_ = default__.safeIndex(152, (d_836_v23_).length(0))
                        (d_836_v23_)[index120_] = (d_827_v16_) + (d_827_v16_)
                        index121_ = default__.safeIndex(152, (d_836_v23_).length(0))
                        (d_836_v23_)[index121_] = ((d_827_v16_) + ((_dafny.SeqWithoutIsStrInference([d_822_v11_ for d_838_i5_ in range(default__.abs(-243))])) + (d_827_v16_))).set(default__.safeIndex(d_820_cf34_, len((d_827_v16_) + ((_dafny.SeqWithoutIsStrInference([d_822_v11_ for d_839_i5_ in range(default__.abs(-243))])) + (d_827_v16_)))), d_822_v11_)
                        (globalState).f5 = d_792_v0_
                        d_828_v17_ = d_828_v17_
                    elif True:
                        d_840___mcc_h6_ = source13_.cf28
                        d_841_cf28_ = d_840___mcc_h6_
                        d_842_v24_: _dafny.Array
                        def lambda45_(d_843_v0_):
                            def lambda46_(d_844_i6_):
                                return d_843_v0_

                            return lambda46_

                        init23_ = lambda45_(d_792_v0_)
                        nw138_ = _dafny.Array(None, 8)
                        for i0_23_ in range(nw138_.length(0)):
                            nw138_[i0_23_] = init23_(i0_23_)
                        d_842_v24_ = nw138_
                        d_845_v25_: _dafny.Seq
                        d_845_v25_ = _dafny.SeqWithoutIsStrInference([d_842_v24_, d_842_v24_])
                        d_845_v25_ = d_845_v25_
                        (globalState).f1 = d_792_v0_
                        d_846_v26_: str
                        d_846_v26_ = _dafny.CodePoint('o')
                        r3 = d_846_v26_
                        d_847_v27_: _dafny.Set
                        d_847_v27_ = _dafny.Set({d_792_v0_, d_792_v0_, d_792_v0_})
                        d_848_v28_: _dafny.Map
                        d_848_v28_ = _dafny.Map({(d_847_v27_) != (d_847_v27_): (self).f6})
                        d_848_v28_ = (d_848_v28_).set(d_792_v0_, (self).f6)
                    d_849_v29_: _dafny.Seq
                    d_849_v29_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lropjewiy"))
                    d_850_v30_: _dafny.Set
                    d_850_v30_ = _dafny.Set({d_792_v0_})
                    d_851_v31_: _dafny.Set
                    d_851_v31_ = _dafny.Set({len(d_850_v30_), (default__.fm15((self).f6, d_792_v0_, globalState)) + ((self).f6), (self).f6, (0) - ((self).f6)})
                    d_852_v32_: _dafny.Array
                    def lambda47_(d_853_v0_):
                        def lambda48_(d_854_i7_):
                            return d_853_v0_

                        return lambda48_

                    init24_ = lambda47_(d_792_v0_)
                    nw139_ = _dafny.Array(None, 7)
                    for i0_24_ in range(nw139_.length(0)):
                        nw139_[i0_24_] = init24_(i0_24_)
                    d_852_v32_ = nw139_
                    d_855_v33_: _dafny.Array
                    nw140_ = _dafny.Array(int(0), 4)
                    d_855_v33_ = nw140_
                    index122_ = default__.safeIndex(381, (d_855_v33_).length(0))
                    (d_855_v33_)[index122_] = (self).f6
                    d_856_v34_: D7
                    d_856_v34_ = D7_DC18(d_792_v0_, d_792_v0_, (0) - ((self).f6), d_792_v0_)
                    index123_ = default__.safeIndex(381, (d_855_v33_).length(0))
                    rhs121_ = (D3_DC7(d_849_v29_)).cf13
                    rhs122_ = _dafny.Set({(self).f6})
                    rhs123_ = d_852_v32_
                    rhs124_ = (self).f6
                    rhs125_ = (d_856_v34_).cf31
                    lhs92_ = d_855_v33_
                    lhs93_ = default__.safeIndex(381, (d_855_v33_).length(0))
                    lhs94_ = globalState
                    d_849_v29_ = rhs121_
                    d_851_v31_ = rhs122_
                    d_852_v32_ = rhs123_
                    lhs92_[lhs93_] = rhs124_
                    lhs94_.f0 = rhs125_
                    d_857_v35_: _dafny.Map
                    d_857_v35_ = _dafny.Map({(self).f6: (len(d_793_v1_) if d_792_v0_ else (d_855_v33_)[default__.safeIndex(381, (d_855_v33_).length(0))])})
                    d_858_v36_: _dafny.Seq
                    d_858_v36_ = _dafny.SeqWithoutIsStrInference([(d_855_v33_)[default__.safeIndex(381, (d_855_v33_).length(0))], (self).f6])
                    d_859_v37_: str
                    d_859_v37_ = _dafny.CodePoint('p')
                    d_860_v38_: _dafny.Map
                    d_860_v38_ = _dafny.Map({default__.safeModuloInt(len(_dafny.Map({len((d_858_v36_).set(default__.safeIndex((0) - (len(d_793_v1_)), len(d_858_v36_)), (d_855_v33_)[default__.safeIndex(381, (d_855_v33_).length(0))])): -615})), default__.fm24((d_855_v33_)[default__.safeIndex(381, (d_855_v33_).length(0))], globalState)): d_859_v37_})
                    def iife109_():
                        coll45_ = _dafny.Map()
                        compr_45_: str
                        for compr_45_ in (d_849_v29_).Elements:
                            d_861_v39_: str = compr_45_
                            if (d_861_v39_) in (d_849_v29_):
                                coll45_[d_861_v39_] = d_792_v0_
                        return _dafny.Map(coll45_)
                    def iife110_():
                        coll46_ = _dafny.Map()
                        compr_46_: str
                        for compr_46_ in (d_849_v29_).Elements:
                            d_862_v39_: str = compr_46_
                            if (d_862_v39_) in (d_849_v29_):
                                coll46_[d_862_v39_] = d_792_v0_
                        return _dafny.Map(coll46_)
                    rhs126_ = default__.fm15((d_855_v33_)[default__.safeIndex(381, (d_855_v33_).length(0))], d_792_v0_, globalState)
                    rhs127_ = ((d_860_v38_)[len(iife109_()
                    )] if (len(iife110_()
                    )) in (d_860_v38_) else d_859_v37_)
                    rhs128_ = d_857_v35_
                    r1 = rhs126_
                    r3 = rhs127_
                    d_857_v35_ = rhs128_
                    pass
            pass
        d_863_v40_: bool
        d_863_v40_ = True
        if d_863_v40_:
            (globalState).f5 = d_863_v40_
            d_864_v41_: _dafny.Array
            def lambda49_(d_865_v40_):
                def lambda50_(d_866_i8_):
                    return d_865_v40_

                return lambda50_

            init25_ = lambda49_(d_863_v40_)
            nw141_ = _dafny.Array(None, 24)
            for i0_25_ in range(nw141_.length(0)):
                nw141_[i0_25_] = init25_(i0_25_)
            d_864_v41_ = nw141_
            d_867_v42_: _dafny.Map
            d_867_v42_ = _dafny.Map({d_863_v40_: d_864_v41_})
            rhs129_ = d_867_v42_
            rhs130_ = default__.safeDivisionInt((0) - ((self).f6), (self).f6)
            d_867_v42_ = rhs129_
            r1 = rhs130_
            d_868_v43_: str
            d_868_v43_ = _dafny.CodePoint('j')
            d_869_v44_: _dafny.Map
            d_869_v44_ = _dafny.Map({d_868_v43_: (self).f6})
            d_870_v45_: _dafny.Seq
            d_870_v45_ = _dafny.SeqWithoutIsStrInference([((d_869_v44_)[d_868_v43_] if (d_868_v43_) in (d_869_v44_) else (self).f6)])
            d_870_v45_ = (D1_DC4(d_870_v45_)).cf11
            rhs131_ = ((d_867_v42_)[d_863_v40_] if (d_863_v40_) in (d_867_v42_) else d_864_v41_)
            rhs132_ = default__.fm24((self).f6, globalState)
            lhs95_ = globalState
            d_864_v41_ = rhs131_
            lhs95_.f0 = rhs132_
            d_871_v46_: _dafny.Seq
            d_871_v46_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nemkv"))
            d_872_v47_: _dafny.Map
            d_872_v47_ = _dafny.Map({d_863_v40_: d_871_v46_})
            d_873_v48_: _dafny.MultiSet
            d_873_v48_ = _dafny.MultiSet([d_863_v40_])
            d_874_v49_: T0
            nw142_ = C3()
            nw142_.ctor__(186, d_870_v45_, (d_873_v48_).cardinality)
            d_874_v49_ = nw142_
            d_875_v50_: T1
            nw143_ = C3()
            nw143_.ctor__((self).f6, d_870_v45_, len(default__.fm2(d_871_v46_, globalState)))
            d_875_v50_ = nw143_
            d_876_v51_: _dafny.MultiSet
            d_876_v51_ = _dafny.MultiSet([d_875_v50_])
            d_877_v52_: _dafny.MultiSet
            d_877_v52_ = d_876_v51_
            d_878_v53_: D11
            d_878_v53_ = D11_DC26(d_874_v49_, d_877_v52_, len(d_869_v44_), d_871_v46_)
            d_879_v54_: _dafny.Array
            nw144_ = _dafny.Array(None, 17)
            nw144_[int(0)] = (d_871_v46_).set(default__.safeIndex((d_870_v45_)[default__.safeIndex((self).f6, len(d_870_v45_))], len(d_871_v46_)), d_868_v43_)
            nw144_[int(1)] = _dafny.SeqWithoutIsStrInference([d_868_v43_ for d_880_i9_ in range(default__.abs(199))])
            nw144_[int(2)] = _dafny.SeqWithoutIsStrInference([d_868_v43_ for d_881_i10_ in range(default__.abs(947))])
            nw144_[int(3)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "npv"))
            nw144_[int(4)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ev"))) + (d_871_v46_)
            nw144_[int(5)] = (d_871_v46_) + (d_871_v46_)
            nw144_[int(6)] = d_871_v46_
            nw144_[int(7)] = (d_871_v46_) + (d_871_v46_)
            nw144_[int(8)] = d_871_v46_
            nw144_[int(9)] = d_871_v46_
            nw144_[int(10)] = (d_871_v46_ if d_863_v40_ else (_dafny.SeqWithoutIsStrInference([d_868_v43_ for d_882_i11_ in range(default__.abs(489))])).set(default__.safeIndex((self).f6, len(_dafny.SeqWithoutIsStrInference([d_868_v43_ for d_883_i11_ in range(default__.abs(489))]))), _dafny.CodePoint('p')))
            nw144_[int(11)] = d_871_v46_
            nw144_[int(12)] = ((d_871_v46_).set(default__.safeIndex((self).f6, len(d_871_v46_)), d_868_v43_)) + (d_871_v46_)
            nw144_[int(13)] = (d_871_v46_) + (((d_872_v47_)[d_863_v40_] if (d_863_v40_) in (d_872_v47_) else d_871_v46_))
            nw144_[int(14)] = d_871_v46_
            nw144_[int(15)] = (d_871_v46_ if d_863_v40_ else d_871_v46_)
            nw144_[int(16)] = ((d_878_v53_).cf48).set(default__.safeIndex((d_874_v49_).f6, len((d_878_v53_).cf48)), d_868_v43_)
            d_879_v54_ = nw144_
            index124_ = default__.safeIndex(946, (d_879_v54_).length(0))
            (d_879_v54_)[index124_] = d_871_v46_
            index125_ = default__.safeIndex(946, (d_879_v54_).length(0))
            (d_879_v54_)[index125_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nqeesu"))
        elif True:
            d_884_v55_: _dafny.Seq
            d_884_v55_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rgltcl"))
            d_884_v55_ = default__.fm2(d_884_v55_, globalState)
            d_885_v56_: _dafny.Array
            nw145_ = _dafny.Array(_dafny.Map({}), 9)
            d_885_v56_ = nw145_
            d_886_v57_: _dafny.Map
            d_886_v57_ = _dafny.Map({d_863_v40_: len(_dafny.SeqWithoutIsStrInference([(self).f6, (self).f6]))})
            index126_ = default__.safeIndex(84, (d_885_v56_).length(0))
            (d_885_v56_)[index126_] = d_886_v57_
            d_887_v58_: _dafny.Map
            d_887_v58_ = _dafny.Map({d_863_v40_: _dafny.CodePoint('o')})
            d_888_v59_: str
            d_888_v59_ = _dafny.CodePoint('t')
            d_889_v60_: _dafny.Seq
            d_889_v60_ = _dafny.SeqWithoutIsStrInference([(_dafny.Map({d_863_v40_: len(_dafny.SeqWithoutIsStrInference([((d_887_v58_)[not(d_863_v40_)] if (not(d_863_v40_)) in (d_887_v58_) else d_888_v59_)]))})) | (d_886_v57_), d_886_v57_])
            d_890_v61_: _dafny.MultiSet
            d_890_v61_ = _dafny.MultiSet([d_863_v40_])
            index127_ = default__.safeIndex(84, (d_885_v56_).length(0))
            (d_885_v56_)[index127_] = (d_889_v60_)[default__.safeIndex(((d_890_v61_).cardinality) * ((self).f6), len(d_889_v60_))]
            r0 = 575
            d_891_v62_: _dafny.Map
            d_891_v62_ = _dafny.Map({(0) - ((self).f6): _dafny.Set({d_863_v40_})})
            d_892_v63_: D0
            d_892_v63_ = D0_DC0(d_891_v62_)
            d_893_v64_: C1
            nw146_ = C1()
            nw146_.ctor__()
            d_893_v64_ = nw146_
            d_894_v65_: _dafny.Seq
            d_894_v65_ = _dafny.SeqWithoutIsStrInference([d_863_v40_])
            d_895_v66_: _dafny.Array
            nw147_ = _dafny.Array(None, 20)
            nw147_[int(0)] = default__.fm1(d_892_v63_, d_863_v40_, (self).f6, d_863_v40_, globalState)
            nw147_[int(1)] = d_863_v40_
            nw147_[int(2)] = (d_894_v65_)[default__.safeIndex(136, len(d_894_v65_))]
            nw147_[int(3)] = d_863_v40_
            nw147_[int(4)] = d_863_v40_
            nw147_[int(5)] = d_863_v40_
            nw147_[int(6)] = d_863_v40_
            nw147_[int(7)] = d_863_v40_
            nw147_[int(8)] = True
            nw147_[int(9)] = d_863_v40_
            nw147_[int(10)] = d_863_v40_
            nw147_[int(11)] = True
            nw147_[int(12)] = d_863_v40_
            nw147_[int(13)] = not(d_863_v40_)
            nw147_[int(14)] = not(d_863_v40_)
            nw147_[int(15)] = d_863_v40_
            nw147_[int(16)] = False
            nw147_[int(17)] = d_863_v40_
            nw147_[int(18)] = d_863_v40_
            nw147_[int(19)] = d_863_v40_
            d_895_v66_ = nw147_
            d_896_v67_: _dafny.MultiSet
            d_896_v67_ = _dafny.MultiSet([d_895_v66_])
            d_897_v68_: _dafny.Map
            d_897_v68_ = _dafny.Map({(self).f6: d_893_v64_})
            d_898_v69_: _dafny.Seq
            d_898_v69_ = _dafny.SeqWithoutIsStrInference([d_893_v64_])
            r2 = len(((_dafny.SeqWithoutIsStrInference([d_893_v64_]) if default__.fm1(d_892_v63_, True, default__.fm24(731, globalState), d_863_v40_, globalState) else (_dafny.SeqWithoutIsStrInference([d_893_v64_, d_893_v64_])).set(default__.safeIndex((d_896_v67_).cardinality, len(_dafny.SeqWithoutIsStrInference([d_893_v64_, d_893_v64_]))), ((d_897_v68_)[72] if (72) in (d_897_v68_) else d_893_v64_)))) + ((d_898_v69_) + (d_898_v69_)))
            d_899_v70_: _dafny.Map
            d_899_v70_ = _dafny.Map({(self).f6: (self).f6})
            r1 = default__.safeDivisionInt(default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([d_863_v40_, d_863_v40_, (d_894_v65_)[default__.safeIndex((self).f6, len(d_894_v65_))], d_863_v40_])), (self).f6), (((d_899_v70_)[(self).f6] if ((self).f6) in (d_899_v70_) else (0) - (len(d_884_v55_)))) * ((self).f6))
        d_900_v71_: _dafny.Array
        def lambda51_(d_901_i13_):
            return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "igm"))) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sgxxap")))

        init26_ = lambda51_
        nw148_ = _dafny.Array(None, 24)
        for i0_26_ in range(nw148_.length(0)):
            nw148_[i0_26_] = init26_(i0_26_)
        d_900_v71_ = nw148_
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_900_v71_).length(0)):
            d_902_i12_: int = guard_loop_3_
            if (True) and (((0) <= (d_902_i12_)) and ((d_902_i12_) < ((d_900_v71_).length(0)))):
                (d_900_v71_)[(d_902_i12_)] = ((_dafny.Set({D7_DC17(_dafny.SeqWithoutIsStrInference([d_863_v40_, d_863_v40_])), D7_DC17(_dafny.SeqWithoutIsStrInference([d_863_v40_, d_863_v40_]))})) | (_dafny.Set({D7_DC17(_dafny.SeqWithoutIsStrInference([d_863_v40_]))}))) != (_dafny.Set({D7_DC17(_dafny.SeqWithoutIsStrInference([d_863_v40_]))}))
        d_903_v72_: _dafny.Seq
        d_903_v72_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gvljalgu"))
        d_904_v73_: D3
        d_904_v73_ = D3_DC7(d_903_v72_)
        source14_ = D3_DC9((D3_DC9(d_904_v73_) if d_863_v40_ else d_904_v73_))
        if source14_.is_DC8:
            d_905___mcc_h7_ = source14_.cf14
            d_906___mcc_h8_ = source14_.cf15
            d_907___mcc_h9_ = source14_.cf16
            d_908_cf16_ = d_907___mcc_h9_
            d_909_cf15_ = d_906___mcc_h8_
            d_910_cf14_ = d_905___mcc_h7_
            r1 = default__.safeDivisionInt(d_908_cf16_, d_908_cf16_)
            d_911_v74_: _dafny.Array
            d_911_v74_ = d_900_v71_
            rhs133_ = (d_911_v74_)
            rhs134_ = d_909_cf15_
            d_900_v71_ = rhs133_
            d_909_cf15_ = rhs134_
            index128_ = default__.safeIndex(981, (d_900_v71_).length(0))
            (d_900_v71_)[index128_] = d_863_v40_
            index129_ = default__.safeIndex(981, (d_900_v71_).length(0))
            (d_900_v71_)[index129_] = d_863_v40_
            d_912_v75_: C1
            nw149_ = C1()
            nw149_.ctor__()
            d_912_v75_ = nw149_
        elif source14_.is_DC7:
            d_913___mcc_h10_ = source14_.cf13
            d_914_cf13_ = d_913___mcc_h10_
            d_915_v76_: _dafny.Array
            nw150_ = _dafny.Array(int(0), 11)
            d_915_v76_ = nw150_
            d_916_v77_: D3
            d_916_v77_ = D3_DC7(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dfgoq")))
            index130_ = default__.safeIndex(412, (d_915_v76_).length(0))
            (d_915_v76_)[index130_] = len((d_916_v77_).cf13)
            index131_ = default__.safeIndex(412, (d_915_v76_).length(0))
            (d_915_v76_)[index131_] = (self).f6
            (globalState).f5 = d_863_v40_
            (globalState).f5 = (d_863_v40_) or (d_863_v40_)
            d_917_v78_: _dafny.Array
            def lambda52_(d_918_i14_):
                return _dafny.Map({_dafny.CodePoint('j'): False})

            init27_ = lambda52_
            nw151_ = _dafny.Array(None, 11)
            for i0_27_ in range(nw151_.length(0)):
                nw151_[i0_27_] = init27_(i0_27_)
            d_917_v78_ = nw151_
            d_919_v79_: D14
            d_919_v79_ = D14_DC35(d_917_v78_)
            d_920_v80_: _dafny.Array
            nw152_ = _dafny.Array(None, 16)
            nw152_[int(0)] = d_917_v78_
            nw152_[int(1)] = d_917_v78_
            nw152_[int(2)] = d_917_v78_
            nw152_[int(3)] = d_917_v78_
            nw152_[int(4)] = d_917_v78_
            nw152_[int(5)] = (d_917_v78_ if d_863_v40_ else d_917_v78_)
            nw152_[int(6)] = d_917_v78_
            nw152_[int(7)] = d_917_v78_
            nw152_[int(8)] = (d_919_v79_).cf67
            nw152_[int(9)] = d_917_v78_
            nw152_[int(10)] = d_917_v78_
            nw152_[int(11)] = d_917_v78_
            nw152_[int(12)] = d_917_v78_
            nw152_[int(13)] = d_917_v78_
            nw152_[int(14)] = d_917_v78_
            nw152_[int(15)] = (d_917_v78_ if d_863_v40_ else d_917_v78_)
            d_920_v80_ = nw152_
            index132_ = default__.safeIndex(201, (d_920_v80_).length(0))
            (d_920_v80_)[index132_] = d_917_v78_
            index133_ = default__.safeIndex(201, (d_920_v80_).length(0))
            index134_ = default__.safeIndex(412, (d_915_v76_).length(0))
            rhs135_ = not(d_863_v40_)
            rhs136_ = (d_915_v76_)[default__.safeIndex(412, (d_915_v76_).length(0))]
            rhs137_ = d_917_v78_
            rhs138_ = (self).f6
            lhs96_ = globalState
            lhs97_ = globalState
            lhs98_ = d_920_v80_
            lhs99_ = default__.safeIndex(201, (d_920_v80_).length(0))
            lhs100_ = d_915_v76_
            lhs101_ = default__.safeIndex(412, (d_915_v76_).length(0))
            lhs96_.f1 = rhs135_
            lhs97_.f2 = rhs136_
            lhs98_[lhs99_] = rhs137_
            lhs100_[lhs101_] = rhs138_
        elif True:
            d_921___mcc_h11_ = source14_.cf17
            d_922_cf17_ = d_921___mcc_h11_
            d_923_v81_: _dafny.Seq
            d_923_v81_ = _dafny.SeqWithoutIsStrInference([d_863_v40_, not(d_863_v40_)])
            d_924_v82_: D12
            d_924_v82_ = D12_DC28(_dafny.Map({d_863_v40_: len(d_923_v81_)}))
            d_924_v82_ = d_924_v82_
            d_925_v83_: _dafny.Set
            d_925_v83_ = _dafny.Set({d_863_v40_, d_863_v40_})
            d_926_v84_: C0
            nw153_ = C0()
            nw153_.ctor__(d_863_v40_)
            d_926_v84_ = nw153_
            d_927_v85_: C0
            d_927_v85_ = d_926_v84_
            d_928_v86_: _dafny.Array
            nw154_ = _dafny.Array(None, 18)
            nw154_[int(0)] = d_926_v84_
            nw154_[int(1)] = d_926_v84_
            nw154_[int(2)] = d_926_v84_
            nw154_[int(3)] = d_926_v84_
            nw154_[int(4)] = d_926_v84_
            nw154_[int(5)] = d_926_v84_
            nw154_[int(6)] = d_926_v84_
            nw154_[int(7)] = d_926_v84_
            nw154_[int(8)] = d_926_v84_
            nw154_[int(9)] = d_926_v84_
            nw154_[int(10)] = d_926_v84_
            nw154_[int(11)] = d_926_v84_
            nw154_[int(12)] = d_926_v84_
            nw154_[int(13)] = d_926_v84_
            nw154_[int(14)] = (d_927_v85_)
            nw154_[int(15)] = d_926_v84_
            nw154_[int(16)] = d_926_v84_
            nw154_[int(17)] = d_926_v84_
            d_928_v86_ = nw154_
            d_929_v87_: _dafny.Map
            d_929_v87_ = _dafny.Map({len(d_925_v83_): d_928_v86_})
            d_929_v87_ = (d_929_v87_).set((self).f6, d_928_v86_)
            d_930_v88_: _dafny.Set
            d_930_v88_ = _dafny.Set({(self).f6})
            d_931_v89_: _dafny.Seq
            d_931_v89_ = _dafny.SeqWithoutIsStrInference([(self).f6, (0) - ((self).f6), (self).f6, (self).f6, (self).f6])
            d_932_v90_: C3
            nw155_ = C3()
            nw155_.ctor__(default__.safeModuloInt((self).f6, len(d_930_v88_)), (_dafny.SeqWithoutIsStrInference([(self).f6])) + (d_931_v89_), 918)
            d_932_v90_ = nw155_
            d_933_v91_: C0
            nw156_ = C0()
            nw156_.ctor__((d_863_v40_ if d_863_v40_ else d_863_v40_))
            d_933_v91_ = nw156_
        d_934_v92_: _dafny.Set
        d_934_v92_ = _dafny.Set({d_863_v40_})
        d_935_v93_: D6
        d_935_v93_ = D6_DC15(d_934_v92_)
        d_936_i15_: int
        d_936_i15_ = 0
        with _dafny.label("3"):
            pat_let_tv29_ = d_863_v40_
            pat_let_tv30_ = d_863_v40_
            def lambda55_(source15_):
                if source15_.is_DC16:
                    d_946___mcc_h12_ = source15_.cf27
                    d_947_cf27_ = d_946___mcc_h12_
                    return ((self).f6) > (len(_dafny.Map({pat_let_tv29_: _dafny.SeqWithoutIsStrInference([13])})))
                elif True:
                    d_948___mcc_h13_ = source15_.cf26
                    d_949_cf26_ = d_948___mcc_h13_
                    return pat_let_tv30_

            while lambda55_(d_935_v93_):
                with _dafny.c_label("3"):
                    if (d_936_i15_) >= (100):
                        raise _dafny.Break("3")
                    d_936_i15_ = (d_936_i15_) + (1)
                    d_937_v94_: _dafny.Seq
                    d_937_v94_ = _dafny.SeqWithoutIsStrInference([(self).f6])
                    d_938_v95_: D13
                    d_938_v95_ = D13_DC32(len(d_937_v94_), d_863_v40_, d_863_v40_, d_863_v40_, d_863_v40_)
                    d_939_v96_: _dafny.Seq
                    d_939_v96_ = _dafny.SeqWithoutIsStrInference([(d_938_v95_).cf56])
                    d_940_v97_: C3
                    nw157_ = C3()
                    nw157_.ctor__((self).f6, d_937_v94_, (self).f6)
                    d_940_v97_ = nw157_
                    (globalState).f0 = default__.safeDivisionInt((self).f6, (len(d_903_v72_)) * ((d_940_v97_).f9))
                    d_941_v98_: _dafny.Array
                    def lambda53_(d_942_v92_):
                        def lambda54_(d_943_i16_):
                            return default__.safeDivisionInt(d_943_i16_, len(d_942_v92_))

                        return lambda54_

                    init28_ = lambda53_(d_934_v92_)
                    nw158_ = _dafny.Array(None, 23)
                    for i0_28_ in range(nw158_.length(0)):
                        nw158_[i0_28_] = init28_(i0_28_)
                    d_941_v98_ = nw158_
                    index135_ = default__.safeIndex(473, (d_941_v98_).length(0))
                    (d_941_v98_)[index135_] = (self).f6
                    d_944_v99_: _dafny.Map
                    d_944_v99_ = _dafny.Map({d_863_v40_: (self).f6})
                    d_945_v100_: D12
                    d_945_v100_ = D12_DC28(d_944_v99_)
                    index136_ = default__.safeIndex(473, (d_941_v98_).length(0))
                    (d_941_v98_)[index136_] = len((d_945_v100_).cf51)
                    (globalState).f1 = d_863_v40_
                    pass
            pass
        d_950_i17_: int
        d_950_i17_ = 0
        with _dafny.label("4"):
            while False:
                with _dafny.c_label("4"):
                    if (d_950_i17_) >= (100):
                        raise _dafny.Break("4")
                    d_950_i17_ = (d_950_i17_) + (1)
                    d_951_v101_: str
                    d_951_v101_ = _dafny.CodePoint('j')
                    d_952_v102_: _dafny.Array
                    nw159_ = _dafny.Array(None, 19)
                    nw159_[int(0)] = d_951_v101_
                    nw159_[int(1)] = _dafny.CodePoint('l')
                    nw159_[int(2)] = d_951_v101_
                    nw159_[int(3)] = d_951_v101_
                    nw159_[int(4)] = d_951_v101_
                    nw159_[int(5)] = d_951_v101_
                    nw159_[int(6)] = d_951_v101_
                    nw159_[int(7)] = d_951_v101_
                    nw159_[int(8)] = d_951_v101_
                    nw159_[int(9)] = d_951_v101_
                    nw159_[int(10)] = d_951_v101_
                    nw159_[int(11)] = d_951_v101_
                    nw159_[int(12)] = (d_951_v101_ if False else d_951_v101_)
                    nw159_[int(13)] = d_951_v101_
                    nw159_[int(14)] = d_951_v101_
                    nw159_[int(15)] = d_951_v101_
                    nw159_[int(16)] = d_951_v101_
                    nw159_[int(17)] = d_951_v101_
                    nw159_[int(18)] = d_951_v101_
                    d_952_v102_ = nw159_
                    index137_ = default__.safeIndex(37, (d_952_v102_).length(0))
                    (d_952_v102_)[index137_] = _dafny.CodePoint('m')
                    index138_ = default__.safeIndex(37, (d_952_v102_).length(0))
                    (d_952_v102_)[index138_] = d_951_v101_
                    d_953_v103_: _dafny.Seq
                    d_953_v103_ = _dafny.SeqWithoutIsStrInference([d_863_v40_, d_863_v40_])
                    d_954_v104_: _dafny.MultiSet
                    d_954_v104_ = _dafny.MultiSet([d_863_v40_])
                    d_955_v105_: _dafny.Seq
                    d_955_v105_ = _dafny.SeqWithoutIsStrInference([(self).f6])
                    d_956_v106_: D1
                    d_956_v106_ = D1_DC4(d_955_v105_)
                    d_957_v107_: _dafny.Array
                    nw160_ = _dafny.Array(int(0), 11)
                    d_957_v107_ = nw160_
                    d_958_v108_: D3
                    d_958_v108_ = D3_DC8(d_956_v106_, d_957_v107_, (self).f6)
                    d_959_v109_: _dafny.Map
                    d_959_v109_ = _dafny.Map({d_863_v40_: d_863_v40_})
                    d_960_v110_: _dafny.Map
                    d_960_v110_ = _dafny.Map({False: len(d_959_v109_)})
                    d_961_v111_: _dafny.MultiSet
                    d_961_v111_ = _dafny.MultiSet([375, len(d_959_v109_)])
                    pat_let_tv31_ = d_961_v111_
                    d_962_v112_: _dafny.Array
                    nw161_ = _dafny.Array(None, 23)
                    nw161_[int(0)] = (self).f6
                    nw161_[int(1)] = len(d_953_v103_)
                    nw161_[int(2)] = (self).f6
                    nw161_[int(3)] = (self).f6
                    nw161_[int(4)] = 364
                    nw161_[int(5)] = (d_954_v104_).cardinality
                    nw161_[int(6)] = (self).f6
                    nw161_[int(7)] = (_dafny.MultiSet(d_955_v105_)).cardinality
                    nw161_[int(8)] = (self).f6
                    nw161_[int(9)] = (self).f6
                    nw161_[int(10)] = (self).f6
                    nw161_[int(11)] = ((self).f6) * ((self).f6)
                    nw161_[int(12)] = default__.fm15((d_958_v108_).cf16, d_863_v40_, globalState)
                    nw161_[int(13)] = (self).f6
                    nw161_[int(14)] = (len(d_960_v110_)) - ((d_954_v104_).cardinality)
                    nw161_[int(15)] = len(d_953_v103_)
                    def iife111_(_pat_let32_0):
                        def iife112_(d_963_dt__update__tmp_h2_):
                            def iife113_(_pat_let33_0):
                                def iife114_(d_964_dt__update_hcf73_h0_):
                                    return D16_DC38(d_964_dt__update_hcf73_h0_)
                                return iife114_(_pat_let33_0)
                            return iife113_(pat_let_tv31_)
                        return iife112_(_pat_let32_0)
                    nw161_[int(16)] = default__.safeModuloInt((self).f6, ((iife111_(D16_DC38(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f6]))))).cf73).cardinality)
                    nw161_[int(17)] = (self).f6
                    nw161_[int(18)] = (self).f6
                    nw161_[int(19)] = len((d_955_v105_) + (_dafny.SeqWithoutIsStrInference([(self).f6 for d_965_i18_ in range(default__.abs(235))])))
                    nw161_[int(20)] = default__.safeDivisionInt((self).f6, (self).f6)
                    nw161_[int(21)] = (self).f6
                    nw161_[int(22)] = (0) - ((self).f6)
                    d_962_v112_ = nw161_
                    index139_ = default__.safeIndex(478, (d_962_v112_).length(0))
                    (d_962_v112_)[index139_] = ((self).f6) - ((0) - ((self).f6))
                    index140_ = default__.safeIndex(478, (d_962_v112_).length(0))
                    (d_962_v112_)[index140_] = (d_954_v104_).cardinality
                    d_966_v113_: _dafny.Array
                    def lambda56_(d_967_v102_, d_968_v40_):
                        def lambda57_(d_969_i19_):
                            return _dafny.Map({(d_967_v102_)[default__.safeIndex(37, (d_967_v102_).length(0))]: d_968_v40_})

                        return lambda57_

                    init29_ = lambda56_(d_952_v102_, d_863_v40_)
                    nw162_ = _dafny.Array(None, 27)
                    for i0_29_ in range(nw162_.length(0)):
                        nw162_[i0_29_] = init29_(i0_29_)
                    d_966_v113_ = nw162_
                    d_970_v114_: D14
                    d_970_v114_ = D14_DC35(d_966_v113_)
                    d_971_v115_: _dafny.Seq
                    d_971_v115_ = _dafny.SeqWithoutIsStrInference([d_970_v114_])
                    if ((d_971_v115_ if False else d_971_v115_)) != (_dafny.SeqWithoutIsStrInference([d_970_v114_])):
                        d_972_v117_: _dafny.Map
                        d_972_v117_ = _dafny.Map({(d_952_v102_)[default__.safeIndex(37, (d_952_v102_).length(0))]: (d_952_v102_)[default__.safeIndex(37, (d_952_v102_).length(0))]})
                        def iife115_():
                            coll47_ = _dafny.Map()
                            compr_47_: str
                            for compr_47_ in (d_972_v117_).keys.Elements:
                                d_973_v116_: str = compr_47_
                                if (d_973_v116_) in (d_972_v117_):
                                    coll47_[d_973_v116_] = d_863_v40_
                            return _dafny.Map(coll47_)
                        (globalState).f1 = (37) != ((0) - ((len(iife115_()
                        )) - (788)))
                        d_903_v72_ = d_903_v72_
                        d_974_v118_: _dafny.Map
                        d_974_v118_ = _dafny.Map({(d_962_v112_)[default__.safeIndex(478, (d_962_v112_).length(0))]: d_953_v103_})
                        d_974_v118_ = d_974_v118_
                        index141_ = default__.safeIndex(529, (d_900_v71_).length(0))
                        (d_900_v71_)[index141_] = d_863_v40_
                        index142_ = default__.safeIndex(529, (d_900_v71_).length(0))
                        (d_900_v71_)[index142_] = d_863_v40_
                        d_975_v120_: D0
                        def iife116_():
                            coll48_ = _dafny.Map()
                            compr_48_: int
                            for compr_48_ in _dafny.IntegerRange(934, 861):
                                d_976_v119_: int = compr_48_
                                if ((934) <= (d_976_v119_)) and ((d_976_v119_) < (861)):
                                    coll48_[(d_976_v119_) + ((d_962_v112_)[default__.safeIndex(478, (d_962_v112_).length(0))])] = d_934_v92_
                            return _dafny.Map(coll48_)
                        d_975_v120_ = D0_DC0(iife116_()
)
                        d_977_v121_: _dafny.Array
                        def lambda58_(d_978_v71_):
                            def lambda59_(d_979_i20_):
                                return (d_978_v71_)[default__.safeIndex(529, (d_978_v71_).length(0))]

                            return lambda59_

                        init30_ = lambda58_(d_900_v71_)
                        nw163_ = _dafny.Array(None, 14)
                        for i0_30_ in range(nw163_.length(0)):
                            nw163_[i0_30_] = init30_(i0_30_)
                        d_977_v121_ = nw163_
                        d_980_v122_: _dafny.Seq
                        d_980_v122_ = _dafny.SeqWithoutIsStrInference([d_977_v121_])
                        index143_ = default__.safeIndex(478, (d_962_v112_).length(0))
                        (d_962_v112_)[index143_] = ((0) - ((d_962_v112_)[default__.safeIndex(478, (d_962_v112_).length(0))])) - ((0) - (len((default__.fm35((d_962_v112_)[default__.safeIndex(478, (d_962_v112_).length(0))], globalState)) | (_dafny.Map({(d_962_v112_)[default__.safeIndex(478, (d_962_v112_).length(0))]: default__.fm1(d_975_v120_, (d_900_v71_)[default__.safeIndex(529, (d_900_v71_).length(0))], len(d_980_v122_), (d_900_v71_)[default__.safeIndex(529, (d_900_v71_).length(0))], globalState)})))))
                    elif True:
                        (globalState).f5 = d_863_v40_
                        d_955_v105_ = d_955_v105_
                        index144_ = default__.safeIndex(478, (d_962_v112_).length(0))
                        rhs139_ = d_863_v40_
                        rhs140_ = ((d_962_v112_)[default__.safeIndex(478, (d_962_v112_).length(0))] if ((self).f6) != (len(d_903_v72_)) else ((self).f6) - ((self).f6))
                        rhs141_ = (d_934_v92_).ispropersubset((d_934_v92_) | (d_934_v92_))
                        rhs142_ = len((d_934_v92_).intersection(d_934_v92_))
                        rhs143_ = ((0) - ((self).f6)) + ((self).f6)
                        lhs102_ = d_962_v112_
                        lhs103_ = default__.safeIndex(478, (d_962_v112_).length(0))
                        lhs104_ = globalState
                        d_863_v40_ = rhs139_
                        lhs102_[lhs103_] = rhs140_
                        d_863_v40_ = rhs141_
                        r2 = rhs142_
                        lhs104_.f0 = rhs143_
                        (globalState).f2 = (self).f6
                        index145_ = default__.safeIndex(478, (d_962_v112_).length(0))
                        (d_962_v112_)[index145_] = (self).f6
                    d_981_v123_: _dafny.Seq
                    d_981_v123_ = _dafny.SeqWithoutIsStrInference([(d_956_v106_).cf11])
                    d_981_v123_ = d_981_v123_
                    pass
            pass
        r0 = (self).f6
        r1 = (0) - ((self).f6)
        r2 = -773
        d_982_v124_: str
        d_982_v124_ = _dafny.CodePoint('a')
        r3 = d_982_v124_
        return r0, r1, r2, r3


class C5:
    def  __init__(self):
        self._f7: bool = False
        self._f8: str = _dafny.CodePoint('D')
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    def ctor__(self, f7, f8):
        (self)._f7 = f7
        (self)._f8 = f8

    def fm5(self, globalState):
        def lambda60_(source16_):
            if source16_.is_DC1:
                d_983___mcc_h0_ = source16_.cf1
                d_984___mcc_h1_ = source16_.cf2
                d_985___mcc_h2_ = source16_.cf3
                d_986___mcc_h3_ = source16_.cf4
                d_987_cf4_ = d_986___mcc_h3_
                d_988_cf3_ = d_985___mcc_h2_
                d_989_cf2_ = d_984___mcc_h1_
                d_990_cf1_ = d_983___mcc_h0_
                return ((D1_DC4(_dafny.SeqWithoutIsStrInference([d_988_cf3_, d_990_cf1_]))).cf11) + (_dafny.SeqWithoutIsStrInference([(0) - (d_988_cf3_) for d_991_i0_ in range(default__.abs(581))]))
            elif source16_.is_DC2:
                d_992___mcc_h4_ = source16_.cf5
                d_993___mcc_h5_ = source16_.cf6
                d_994___mcc_h6_ = source16_.cf7
                d_995___mcc_h7_ = source16_.cf8
                d_996___mcc_h8_ = source16_.cf9
                d_997_cf9_ = d_996___mcc_h8_
                d_998_cf8_ = d_995___mcc_h7_
                d_999_cf7_ = d_994___mcc_h6_
                d_1000_cf6_ = d_993___mcc_h5_
                d_1001_cf5_ = d_992___mcc_h4_
                return _dafny.SeqWithoutIsStrInference([len(_dafny.Set({(self).f7})), d_1001_cf5_])
            elif source16_.is_DC0:
                d_1002___mcc_h9_ = source16_.cf0
                d_1003_cf0_ = d_1002___mcc_h9_
                return (D1_DC4(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({(self).f7}))]))).cf11
            elif True:
                d_1004___mcc_h10_ = source16_.cf10
                d_1005_cf10_ = d_1004___mcc_h10_
                return (_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([(self).f7]))), 943, len(_dafny.Map({(self).f7: (D0_DC1(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ejdijcrpi"))), (self).f7, 757, (self).f7)).cf3}))])) + (_dafny.SeqWithoutIsStrInference([-20, len(_dafny.SeqWithoutIsStrInference([(self).f7]))]))

        def iife117_():
            coll49_ = _dafny.Map()
            compr_49_: int
            for compr_49_ in (_dafny.Set({714})).Elements:
                d_1006_v0_: int = compr_49_
                if (d_1006_v0_) in (_dafny.Set({714})):
                    coll49_[(d_1006_v0_) * (322)] = _dafny.Set({(self).f7})
            return _dafny.Map(coll49_)
        return len(lambda60_(D0_DC0(iife117_()
)))

    def fm6(self, p0, p1, p2, globalState):
        return len(_dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rj"))) + (_dafny.SeqWithoutIsStrInference([(self).f8 for d_1007_i1_ in range(default__.abs(846))])) for d_1008_i0_ in range(default__.abs(-724))]))

    def m1(self, p0, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: bool = False
        r3: int = int(0)
        d_1009_v0_: _dafny.Array
        nw164_ = _dafny.Array(int(0), 12)
        d_1009_v0_ = nw164_
        d_1009_v0_ = d_1009_v0_
        d_1010_v1_: _dafny.Array
        nw165_ = _dafny.Array(False, 17)
        d_1010_v1_ = nw165_
        d_1011_v2_: int
        d_1011_v2_ = 441
        d_1012_v3_: _dafny.Map
        d_1012_v3_ = _dafny.Map({d_1011_v2_: d_1010_v1_})
        d_1013_v4_: _dafny.Set
        d_1013_v4_ = _dafny.Set({d_1011_v2_, d_1011_v2_})
        d_1014_v5_: _dafny.Seq
        d_1014_v5_ = _dafny.SeqWithoutIsStrInference([d_1011_v2_, d_1011_v2_, len(d_1013_v4_)])
        d_1015_v6_: _dafny.Array
        nw166_ = _dafny.Array(None, 18)
        nw166_[int(0)] = d_1010_v1_
        nw166_[int(1)] = d_1010_v1_
        nw166_[int(2)] = d_1010_v1_
        nw166_[int(3)] = d_1010_v1_
        nw166_[int(4)] = d_1010_v1_
        nw166_[int(5)] = d_1010_v1_
        nw166_[int(6)] = d_1010_v1_
        nw166_[int(7)] = d_1010_v1_
        nw166_[int(8)] = d_1010_v1_
        nw166_[int(9)] = d_1010_v1_
        nw166_[int(10)] = d_1010_v1_
        nw166_[int(11)] = d_1010_v1_
        nw166_[int(12)] = d_1010_v1_
        nw166_[int(13)] = d_1010_v1_
        nw166_[int(14)] = d_1010_v1_
        nw166_[int(15)] = ((d_1012_v3_)[(d_1014_v5_)[default__.safeIndex(d_1011_v2_, len(d_1014_v5_))]] if ((d_1014_v5_)[default__.safeIndex(d_1011_v2_, len(d_1014_v5_))]) in (d_1012_v3_) else d_1010_v1_)
        nw166_[int(16)] = d_1010_v1_
        nw166_[int(17)] = d_1010_v1_
        d_1015_v6_ = nw166_
        index146_ = default__.safeIndex(163, (d_1015_v6_).length(0))
        (d_1015_v6_)[index146_] = d_1010_v1_
        d_1016_v7_: D1
        d_1016_v7_ = D1_DC4(default__.fm7(globalState))
        index147_ = default__.safeIndex(425, (d_1009_v0_).length(0))
        (d_1009_v0_)[index147_] = d_1011_v2_
        d_1017_v8_: _dafny.Seq
        d_1017_v8_ = _dafny.SeqWithoutIsStrInference([False])
        d_1018_v9_: _dafny.Array
        d_1018_v9_ = d_1010_v1_
        d_1019_v10_: D1
        d_1019_v10_ = D1_DC5()
        pat_let_tv32_ = d_1011_v2_
        pat_let_tv33_ = d_1011_v2_
        pat_let_tv34_ = d_1011_v2_
        index148_ = default__.safeIndex(163, (d_1015_v6_).length(0))
        index149_ = default__.safeIndex(425, (d_1009_v0_).length(0))
        def lambda61_(source17_):
            if source17_.is_DC5:
                return (pat_let_tv32_) - (pat_let_tv33_)
            elif True:
                d_1020___mcc_h0_ = source17_.cf11
                d_1021_cf11_ = d_1020___mcc_h0_
                return pat_let_tv34_

        rhs144_ = (d_1011_v2_) + (d_1011_v2_)
        rhs145_ = ((self).fm6(d_1011_v2_, p0, 816, globalState)) + ((0) - (len(_dafny.Map({True: len(d_1017_v8_)}))))
        rhs146_ = (d_1018_v9_)
        rhs147_ = D1_DC4(d_1014_v5_)
        rhs148_ = lambda61_(d_1019_v10_)
        lhs105_ = globalState
        lhs106_ = d_1015_v6_
        lhs107_ = default__.safeIndex(163, (d_1015_v6_).length(0))
        lhs108_ = d_1009_v0_
        lhs109_ = default__.safeIndex(425, (d_1009_v0_).length(0))
        r3 = rhs144_
        lhs105_.f0 = rhs145_
        lhs106_[lhs107_] = rhs146_
        d_1016_v7_ = rhs147_
        lhs108_[lhs109_] = rhs148_
        d_1022_v11_: _dafny.Array
        nw167_ = _dafny.Array(_dafny.Map({}), 14)
        d_1022_v11_ = nw167_
        d_1023_v12_: _dafny.Map
        d_1023_v12_ = _dafny.Map({d_1022_v11_: d_1009_v0_})
        d_1023_v12_ = (d_1023_v12_).set(d_1022_v11_, d_1009_v0_)
        d_1024_v13_: _dafny.Set
        d_1024_v13_ = _dafny.Set({p0})
        d_1024_v13_ = (d_1024_v13_ if False else d_1024_v13_)
        d_1025_v14_: _dafny.Array
        nw168_ = _dafny.Array(_dafny.Map({}), 11)
        d_1025_v14_ = nw168_
        d_1026_v15_: D0
        d_1026_v15_ = D0_DC1((self).fm5(globalState), p0, (d_1009_v0_)[default__.safeIndex(425, (d_1009_v0_).length(0))], p0)
        d_1027_v16_: D0
        d_1027_v16_ = D0_DC3(d_1026_v15_)
        d_1028_v17_: D0
        d_1028_v17_ = D0_DC3(d_1027_v16_)
        d_1029_v18_: _dafny.Map
        d_1029_v18_ = _dafny.Map({_dafny.Map({(d_1009_v0_)[default__.safeIndex(425, (d_1009_v0_).length(0))]: (self).f7}): d_1028_v17_})
        index150_ = default__.safeIndex(711, (d_1025_v14_).length(0))
        (d_1025_v14_)[index150_] = d_1029_v18_
        d_1030_v19_: str
        d_1030_v19_ = _dafny.CodePoint('d')
        index151_ = default__.safeIndex(711, (d_1025_v14_).length(0))
        rhs149_ = d_1029_v18_
        rhs150_ = default__.fm8((self).f8, (p0 if p0 else (self).f7), (d_1009_v0_)[default__.safeIndex(425, (d_1009_v0_).length(0))], (d_1009_v0_)[default__.safeIndex(425, (d_1009_v0_).length(0))], globalState)
        lhs110_ = d_1025_v14_
        lhs111_ = default__.safeIndex(711, (d_1025_v14_).length(0))
        lhs110_[lhs111_] = rhs149_
        d_1030_v19_ = rhs150_
        index152_ = default__.safeIndex(163, (d_1015_v6_).length(0))
        (d_1015_v6_)[index152_] = (d_1015_v6_)[default__.safeIndex(163, (d_1015_v6_).length(0))]
        r0 = d_1011_v2_
        d_1031_v21_: _dafny.MultiSet
        d_1031_v21_ = _dafny.MultiSet([(self).f7])
        d_1032_v22_: _dafny.Map
        d_1032_v22_ = _dafny.Map({d_1031_v21_: d_1011_v2_})
        d_1033_v24_: _dafny.Seq
        def iife118_():
            def iife120_():
                coll52_ = _dafny.Map()
                compr_52_: _dafny.MultiSet
                for compr_52_ in (d_1032_v22_).keys.Elements:
                    d_1034_v20_: _dafny.MultiSet = compr_52_
                    if (d_1034_v20_) in (d_1032_v22_):
                        coll52_[d_1034_v20_] = d_1017_v8_
                return _dafny.Map(coll52_)
            coll50_ = _dafny.Set()
            def iife119_():
                coll51_ = _dafny.Map()
                compr_51_: _dafny.MultiSet
                for compr_51_ in (d_1032_v22_).keys.Elements:
                    d_1034_v20_: _dafny.MultiSet = compr_51_
                    if (d_1034_v20_) in (d_1032_v22_):
                        coll51_[d_1034_v20_] = d_1017_v8_
                return _dafny.Map(coll51_)
            compr_50_: _dafny.MultiSet
            for compr_50_ in (iife119_()
            ).keys.Elements:
                d_1035_v23_: _dafny.MultiSet = compr_50_
                if (d_1035_v23_) in (iife120_()
                ):
                    coll50_ = coll50_.union(_dafny.Set([d_1035_v23_]))
            return _dafny.Set(coll50_)
        d_1033_v24_ = _dafny.SeqWithoutIsStrInference([(default__.fm9(globalState) if (self).f7 else iife118_()
        )])
        r1 = len((d_1033_v24_)[default__.safeIndex(682, len(d_1033_v24_))])
        r2 = p0
        r3 = d_1011_v2_
        return r0, r1, r2, r3

    def m2(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        d_1036_v0_: _dafny.Map
        d_1036_v0_ = _dafny.Map({p0: _dafny.Set({(self).f7})})
        d_1037_v1_: _dafny.Map
        d_1037_v1_ = _dafny.Map({p1: d_1036_v0_})
        d_1038_v2_: D0
        d_1038_v2_ = D0_DC0(((d_1037_v1_)[p2] if (p2) in (d_1037_v1_) else d_1036_v0_))
        pat_let_tv35_ = p0
        pat_let_tv36_ = p3
        pat_let_tv37_ = p0
        pat_let_tv38_ = p2
        pat_let_tv39_ = p0
        pat_let_tv40_ = p0
        pat_let_tv41_ = p1
        pat_let_tv42_ = p0
        pat_let_tv43_ = p0
        def lambda62_(source19_):
            if source19_.is_DC1:
                d_1039___mcc_h11_ = source19_.cf1
                d_1040___mcc_h12_ = source19_.cf2
                d_1041___mcc_h13_ = source19_.cf3
                d_1042___mcc_h14_ = source19_.cf4
                d_1043_cf4_ = d_1042___mcc_h14_
                d_1044_cf3_ = d_1041___mcc_h13_
                d_1045_cf2_ = d_1040___mcc_h12_
                d_1046_cf1_ = d_1039___mcc_h11_
                return D0_DC1(pat_let_tv35_, True, len(pat_let_tv36_), d_1045_cf2_)
            elif source19_.is_DC2:
                d_1047___mcc_h15_ = source19_.cf5
                d_1048___mcc_h16_ = source19_.cf6
                d_1049___mcc_h17_ = source19_.cf7
                d_1050___mcc_h18_ = source19_.cf8
                d_1051___mcc_h19_ = source19_.cf9
                d_1052_cf9_ = d_1051___mcc_h19_
                d_1053_cf8_ = d_1050___mcc_h18_
                d_1054_cf7_ = d_1049___mcc_h17_
                d_1055_cf6_ = d_1048___mcc_h16_
                d_1056_cf5_ = d_1047___mcc_h15_
                return D0_DC1(d_1053_cf8_, (self).f7, d_1053_cf8_, (self).f7)
            elif source19_.is_DC0:
                d_1057___mcc_h20_ = source19_.cf0
                d_1058_cf0_ = d_1057___mcc_h20_
                def iife121_():
                    coll53_ = _dafny.Map()
                    compr_53_: int
                    for compr_53_ in (_dafny.SeqWithoutIsStrInference([pat_let_tv39_])).Elements:
                        d_1059_v3_: int = compr_53_
                        if (d_1059_v3_) in (_dafny.SeqWithoutIsStrInference([pat_let_tv40_])):
                            coll53_[(d_1059_v3_) * (len(_dafny.Set({(self).f7, (self).f7})))] = pat_let_tv41_
                    return _dafny.Map(coll53_)
                return D0_DC1(len(_dafny.SeqWithoutIsStrInference([pat_let_tv37_])), pat_let_tv38_, len(iife121_()
), (self).f7)
            elif True:
                d_1060___mcc_h21_ = source19_.cf10
                d_1061_cf10_ = d_1060___mcc_h21_
                return D0_DC1(pat_let_tv42_, (self).f7, pat_let_tv43_, (self).f7)

        source18_ = lambda62_(d_1038_v2_)
        if source18_.is_DC1:
            d_1062___mcc_h0_ = source18_.cf1
            d_1063___mcc_h1_ = source18_.cf2
            d_1064___mcc_h2_ = source18_.cf3
            d_1065___mcc_h3_ = source18_.cf4
            d_1066_cf4_ = d_1065___mcc_h3_
            d_1067_cf3_ = d_1064___mcc_h2_
            d_1068_cf2_ = d_1063___mcc_h1_
            d_1069_cf1_ = d_1062___mcc_h0_
            d_1070_v4_: _dafny.MultiSet
            d_1070_v4_ = _dafny.MultiSet([d_1066_cf4_])
            (globalState).f2 = ((d_1067_cf3_) - ((d_1070_v4_).cardinality)) - (p0)
            d_1071_v5_: _dafny.Map
            d_1071_v5_ = _dafny.Map({(not(p2) if (self).f7 else d_1068_cf2_): p1})
            d_1071_v5_ = d_1071_v5_
            d_1072_v6_: _dafny.Seq
            d_1072_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ja"))
            d_1072_v6_ = (_dafny.SeqWithoutIsStrInference([(self).f8 for d_1073_i0_ in range(default__.abs(129))])) + ((d_1072_v6_) + (d_1072_v6_))
            d_1074_v7_: _dafny.Set
            d_1074_v7_ = _dafny.Set({d_1069_cf1_})
            d_1071_v5_ = (d_1071_v5_).set(d_1066_cf4_, (d_1074_v7_).isdisjoint(_dafny.Set({d_1067_cf3_})))
        elif source18_.is_DC2:
            d_1075___mcc_h4_ = source18_.cf5
            d_1076___mcc_h5_ = source18_.cf6
            d_1077___mcc_h6_ = source18_.cf7
            d_1078___mcc_h7_ = source18_.cf8
            d_1079___mcc_h8_ = source18_.cf9
            d_1080_cf9_ = d_1079___mcc_h8_
            d_1081_cf8_ = d_1078___mcc_h7_
            d_1082_cf7_ = d_1077___mcc_h6_
            d_1083_cf6_ = d_1076___mcc_h5_
            d_1084_cf5_ = d_1075___mcc_h4_
            d_1085_v8_: _dafny.Array
            nw169_ = _dafny.Array(False, 26)
            d_1085_v8_ = nw169_
            d_1086_v9_: _dafny.Array
            d_1086_v9_ = d_1085_v8_
            source20_ = (d_1086_v9_ if (self).f7 else d_1085_v8_)
            d_1087___mcc_h22_ = source20_
            d_1088_cf12_ = d_1087___mcc_h22_
            d_1089_v10_: _dafny.Map
            d_1089_v10_ = _dafny.Map({687: p2})
            index153_ = default__.safeIndex(161, (d_1085_v8_).length(0))
            (d_1085_v8_)[index153_] = ((d_1089_v10_)[d_1080_cf9_] if (d_1080_cf9_) in (d_1089_v10_) else not(p1))
            d_1090_v11_: _dafny.Array
            nw170_ = _dafny.Array(_dafny.Array(None, 0), 5)
            d_1090_v11_ = nw170_
            index154_ = default__.safeIndex(205, (d_1090_v11_).length(0))
            (d_1090_v11_)[index154_] = d_1088_cf12_
            d_1091_v12_: _dafny.Map
            d_1091_v12_ = _dafny.Map({(self).f7: p1})
            d_1092_v13_: _dafny.Map
            d_1092_v13_ = _dafny.Map({(self).fm5(globalState): d_1080_cf9_})
            index155_ = default__.safeIndex(161, (d_1085_v8_).length(0))
            index156_ = default__.safeIndex(205, (d_1090_v11_).length(0))
            rhs151_ = len(d_1091_v12_)
            rhs152_ = p2
            rhs153_ = p1
            rhs154_ = d_1088_cf12_
            rhs155_ = (p1 if (((d_1092_v13_)[82] if (82) in (d_1092_v13_) else len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qng"))))) < (d_1082_cf7_) else (D0_DC1(d_1082_cf7_, default__.fm1(d_1038_v2_, not(p2), (self).fm5(globalState), p2, globalState), d_1081_cf8_, (self).f7)).cf4)
            lhs112_ = globalState
            lhs113_ = d_1085_v8_
            lhs114_ = default__.safeIndex(161, (d_1085_v8_).length(0))
            lhs115_ = globalState
            lhs116_ = d_1090_v11_
            lhs117_ = default__.safeIndex(205, (d_1090_v11_).length(0))
            lhs118_ = globalState
            lhs112_.f0 = rhs151_
            lhs113_[lhs114_] = rhs152_
            lhs115_.f1 = rhs153_
            lhs116_[lhs117_] = rhs154_
            lhs118_.f1 = rhs155_
            d_1093_v14_: int
            d_1094_v15_: int
            d_1095_v16_: bool
            d_1096_v17_: int
            out4_: int
            out5_: int
            out6_: bool
            out7_: int
            out4_, out5_, out6_, out7_ = (self).m1(p1, globalState)
            d_1093_v14_ = out4_
            d_1094_v15_ = out5_
            d_1095_v16_ = out6_
            d_1096_v17_ = out7_
            d_1094_v15_ = d_1084_cf5_
            d_1097_v18_: _dafny.Array
            def lambda63_(d_1098_cf7_):
                def lambda64_(d_1099_i1_):
                    return default__.safeModuloInt(d_1099_i1_, d_1098_cf7_)

                return lambda64_

            init31_ = lambda63_(d_1082_cf7_)
            nw171_ = _dafny.Array(None, 18)
            for i0_31_ in range(nw171_.length(0)):
                nw171_[i0_31_] = init31_(i0_31_)
            d_1097_v18_ = nw171_
            index157_ = default__.safeIndex(690, (d_1097_v18_).length(0))
            (d_1097_v18_)[index157_] = p0
            d_1100_v19_: _dafny.Seq
            d_1100_v19_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hjhrdvw"))
            index158_ = default__.safeIndex(690, (d_1097_v18_).length(0))
            rhs156_ = d_1095_v16_
            rhs157_ = d_1093_v14_
            rhs158_ = (0) - (len(d_1100_v19_))
            rhs159_ = (self).f7
            rhs160_ = d_1080_cf9_
            lhs119_ = globalState
            lhs120_ = d_1097_v18_
            lhs121_ = default__.safeIndex(690, (d_1097_v18_).length(0))
            lhs122_ = globalState
            lhs123_ = globalState
            lhs119_.f1 = rhs156_
            d_1081_cf8_ = rhs157_
            lhs120_[lhs121_] = rhs158_
            lhs122_.f5 = rhs159_
            lhs123_.f0 = rhs160_
            d_1081_cf8_ = (d_1084_cf5_) * (892)
            d_1101_v20_: _dafny.Seq
            d_1101_v20_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qtau"))
            d_1102_v21_: _dafny.Map
            d_1102_v21_ = _dafny.Map({p2: d_1101_v20_})
            d_1103_v22_: _dafny.MultiSet
            d_1103_v22_ = _dafny.MultiSet([((d_1102_v21_)[(self).f7] if ((self).f7) in (d_1102_v21_) else d_1101_v20_), d_1101_v20_])
            d_1103_v22_ = (d_1103_v22_) | (_dafny.MultiSet([d_1101_v20_]))
            index159_ = default__.safeIndex(499, (d_1085_v8_).length(0))
            (d_1085_v8_)[index159_] = p2
            index160_ = default__.safeIndex(499, (d_1085_v8_).length(0))
            (d_1085_v8_)[index160_] = (self).f7
        elif source18_.is_DC0:
            d_1104___mcc_h9_ = source18_.cf0
            d_1105_cf0_ = d_1104___mcc_h9_
            d_1106_v23_: _dafny.Array
            nw172_ = _dafny.Array(None, 7)
            nw172_[int(0)] = p0
            nw172_[int(1)] = p0
            nw172_[int(2)] = p0
            nw172_[int(3)] = p0
            nw172_[int(4)] = 813
            nw172_[int(5)] = p0
            nw172_[int(6)] = p0
            d_1106_v23_ = nw172_
            index161_ = default__.safeIndex(708, (d_1106_v23_).length(0))
            (d_1106_v23_)[index161_] = (0) - (p0)
            d_1107_v24_: D0
            d_1107_v24_ = D0_DC1(p0, p2, p0, (self).f7)
            index162_ = default__.safeIndex(708, (d_1106_v23_).length(0))
            rhs161_ = (d_1107_v24_).cf2
            rhs162_ = p0
            lhs124_ = globalState
            lhs125_ = d_1106_v23_
            lhs126_ = default__.safeIndex(708, (d_1106_v23_).length(0))
            lhs124_.f5 = rhs161_
            lhs125_[lhs126_] = rhs162_
            d_1108_v25_: _dafny.Array
            nw173_ = _dafny.Array(None, 1)
            d_1108_v25_ = nw173_
            d_1109_v26_: _dafny.Seq
            d_1109_v26_ = _dafny.SeqWithoutIsStrInference([default__.fm1(D0_DC0(d_1036_v0_), p1, -909, p1, globalState), (p1 if not(p2) else False)])
            rhs163_ = d_1108_v25_
            rhs164_ = _dafny.SeqWithoutIsStrInference([p2])
            rhs165_ = p0
            lhs127_ = globalState
            d_1108_v25_ = rhs163_
            d_1109_v26_ = rhs164_
            lhs127_.f2 = rhs165_
            d_1110_v27_: _dafny.Set
            d_1110_v27_ = _dafny.Set({default__.safeModuloInt(p0, (d_1106_v23_)[default__.safeIndex(708, (d_1106_v23_).length(0))])})
            d_1110_v27_ = d_1110_v27_
            d_1111_v28_: _dafny.MultiSet
            d_1111_v28_ = _dafny.MultiSet([(d_1106_v23_)[default__.safeIndex(708, (d_1106_v23_).length(0))], (d_1106_v23_)[default__.safeIndex(708, (d_1106_v23_).length(0))], ((d_1106_v23_)[default__.safeIndex(708, (d_1106_v23_).length(0))] if not(p2) else len(_dafny.Set({d_1110_v27_, d_1110_v27_, d_1110_v27_, d_1110_v27_, default__.fm10(802, p1, p2, globalState)}))), (self).fm5(globalState)])
            d_1112_v29_: _dafny.Array
            nw174_ = _dafny.Array(False, 26)
            d_1112_v29_ = nw174_
            index163_ = default__.safeIndex(708, (d_1106_v23_).length(0))
            rhs166_ = (d_1106_v23_)[default__.safeIndex(708, (d_1106_v23_).length(0))]
            rhs167_ = (default__.fm1(d_1038_v2_, False, p0, (self).f7, globalState)) or ((self).f7)
            rhs168_ = _dafny.MultiSet([p0, (d_1106_v23_)[default__.safeIndex(708, (d_1106_v23_).length(0))]])
            rhs169_ = d_1112_v29_
            lhs128_ = d_1106_v23_
            lhs129_ = default__.safeIndex(708, (d_1106_v23_).length(0))
            lhs130_ = globalState
            lhs128_[lhs129_] = rhs166_
            lhs130_.f5 = rhs167_
            d_1111_v28_ = rhs168_
            d_1112_v29_ = rhs169_
        elif True:
            d_1113___mcc_h10_ = source18_.cf10
            d_1114_cf10_ = d_1113___mcc_h10_
            d_1115_v30_: _dafny.Map
            d_1115_v30_ = _dafny.Map({p1: (len(_dafny.SeqWithoutIsStrInference([(self).f8 for d_1116_i2_ in range(default__.abs(575))]))) - (p0)})
            d_1115_v30_ = (d_1115_v30_).set(not(p1), (self).fm5(globalState))
            d_1117_v31_: _dafny.Seq
            d_1117_v31_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xp"))
            d_1118_v32_: _dafny.Map
            d_1118_v32_ = _dafny.Map({(len(d_1117_v31_)) * (len(d_1115_v30_)): (p3) != (p3)})
            d_1118_v32_ = (d_1118_v32_).set(p0, (p0) > (723))
            d_1119_v33_: _dafny.Map
            d_1119_v33_ = _dafny.Map({p2: False})
            d_1119_v33_ = (d_1119_v33_).set((p0) == (len(d_1117_v31_)), p1)
            d_1120_v34_: _dafny.Array
            nw175_ = _dafny.Array(_dafny.Array(None, 0), 23)
            d_1120_v34_ = nw175_
            d_1121_v35_: _dafny.Array
            nw176_ = _dafny.Array(False, 1)
            d_1121_v35_ = nw176_
            index164_ = default__.safeIndex(93, (d_1120_v34_).length(0))
            (d_1120_v34_)[index164_] = d_1121_v35_
            index165_ = default__.safeIndex(93, (d_1120_v34_).length(0))
            (d_1120_v34_)[index165_] = d_1121_v35_
        if (self).f7:
            (globalState).f0 = (701) + (p0)
            d_1122_v36_: _dafny.Seq
            d_1122_v36_ = _dafny.SeqWithoutIsStrInference([p3])
            (globalState).f5 = ((d_1122_v36_) + (d_1122_v36_)) <= (d_1122_v36_)
            d_1122_v36_ = (d_1122_v36_) + (_dafny.SeqWithoutIsStrInference([p3 for d_1123_i3_ in range(default__.abs(966))]))
            d_1124_v37_: _dafny.Array
            nw177_ = _dafny.Array(int(0), 24)
            d_1124_v37_ = nw177_
            index166_ = default__.safeIndex(612, (d_1124_v37_).length(0))
            (d_1124_v37_)[index166_] = p0
            d_1125_v38_: _dafny.Map
            d_1125_v38_ = _dafny.Map({p1: p2})
            d_1126_v39_: D0
            d_1126_v39_ = D0_DC1(p0, p1, 883, ((d_1125_v38_)[p2] if (p2) in (d_1125_v38_) else not((self).f7)))
            pat_let_tv44_ = p0
            pat_let_tv45_ = p0
            index167_ = default__.safeIndex(612, (d_1124_v37_).length(0))
            def iife122_(_pat_let34_0):
                def iife123_(d_1127_dt__update__tmp_h0_):
                    def iife124_(_pat_let35_0):
                        def iife125_(d_1128_dt__update_hcf3_h0_):
                            def iife126_(_pat_let36_0):
                                def iife127_(d_1129_dt__update_hcf1_h0_):
                                    def iife128_(_pat_let37_0):
                                        def iife129_(d_1130_dt__update_hcf2_h0_):
                                            return D0_DC1(d_1129_dt__update_hcf1_h0_, d_1130_dt__update_hcf2_h0_, d_1128_dt__update_hcf3_h0_, (d_1127_dt__update__tmp_h0_).cf4)
                                        return iife129_(_pat_let37_0)
                                    return iife128_((self).f7)
                                return iife127_(_pat_let36_0)
                            return iife126_(pat_let_tv45_)
                        return iife125_(_pat_let35_0)
                    return iife124_(pat_let_tv44_)
                return iife123_(_pat_let34_0)
            rhs170_ = 721
            rhs171_ = ((((d_1125_v38_)[(self).f7] if ((self).f7) in (d_1125_v38_) else p1) if p1 else (iife122_(d_1126_v39_)).cf2)) and ((self).f7)
            lhs131_ = d_1124_v37_
            lhs132_ = default__.safeIndex(612, (d_1124_v37_).length(0))
            lhs133_ = globalState
            lhs131_[lhs132_] = rhs170_
            lhs133_.f5 = rhs171_
            d_1131_v40_: _dafny.Seq
            d_1131_v40_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))
            d_1132_v41_: _dafny.Seq
            d_1132_v41_ = _dafny.SeqWithoutIsStrInference([len(_dafny.Set({len(d_1131_v40_)}))])
            d_1133_v42_: _dafny.Set
            d_1133_v42_ = _dafny.Set({p1, False})
            d_1134_v43_: _dafny.Map
            d_1134_v43_ = _dafny.Map({p2: ((_dafny.MultiSet([len(d_1132_v41_), len(d_1133_v42_)])).set(261, default__.abs((d_1124_v37_)[default__.safeIndex(612, (d_1124_v37_).length(0))]))).cardinality})
            d_1135_v44_: T0
            nw178_ = C4()
            nw178_.ctor__(((d_1134_v43_)[p2] if (p2) in (d_1134_v43_) else p0))
            d_1135_v44_ = nw178_
            d_1136_v45_: _dafny.Map
            d_1136_v45_ = _dafny.Map({p2: (self).f8})
            d_1137_v46_: _dafny.Map
            d_1137_v46_ = _dafny.Map({d_1135_v44_: d_1136_v45_})
            (globalState).f2 = len((d_1137_v46_) | ((_dafny.Map({d_1135_v44_: d_1136_v45_})) | (d_1137_v46_)))
        elif True:
            d_1138_v47_: D13
            d_1138_v47_ = D13_DC32(p0, (self).f7, p1, p1, p2)
            d_1139_v48_: D13
            d_1139_v48_ = D13_DC34(d_1138_v47_)
            d_1140_v49_: _dafny.Map
            d_1140_v49_ = _dafny.Map({(self).f7: d_1139_v48_})
            d_1140_v49_ = d_1140_v49_
            d_1141_v50_: _dafny.Array
            def lambda65_(d_1142_p1_):
                def lambda66_(d_1143_i4_):
                    return d_1142_p1_

                return lambda66_

            init32_ = lambda65_(p1)
            nw179_ = _dafny.Array(None, 12)
            for i0_32_ in range(nw179_.length(0)):
                nw179_[i0_32_] = init32_(i0_32_)
            d_1141_v50_ = nw179_
            d_1144_v51_: _dafny.Map
            d_1144_v51_ = _dafny.Map({p0: p0})
            d_1145_v52_: _dafny.Seq
            d_1145_v52_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gomydnwq"))
            d_1146_v53_: T1
            nw180_ = C3()
            nw180_.ctor__(p0, _dafny.SeqWithoutIsStrInference([len(d_1144_v51_), p0, len(d_1145_v52_), p0, p0]), len(p3))
            d_1146_v53_ = nw180_
            d_1147_v54_: D11
            d_1147_v54_ = D11_DC25(d_1141_v50_, d_1146_v53_, (p0) * (p0))
            d_1147_v54_ = d_1147_v54_
            d_1148_v55_: _dafny.Set
            d_1148_v55_ = _dafny.Set({p1, False, (self).f7, p1, False})
            d_1149_v56_: _dafny.Map
            d_1149_v56_ = _dafny.Map({351: not((self).f7)})
            (globalState).f5 = (D14_DC36((self).f7, len(d_1148_v55_), len(d_1145_v52_), d_1149_v56_)).cf68
            (globalState).f0 = 909
            d_1150_v57_: _dafny.Seq
            d_1150_v57_ = _dafny.SeqWithoutIsStrInference([p0])
            d_1151_v58_: _dafny.Array
            nw181_ = _dafny.Array(int(0), 4)
            d_1151_v58_ = nw181_
            d_1152_v59_: _dafny.Map
            d_1152_v59_ = _dafny.Map({d_1150_v57_: (D3_DC8(D1_DC4(d_1150_v57_), d_1151_v58_, p0)).cf15})
            d_1152_v59_ = d_1152_v59_
        d_1153_v60_: _dafny.Array
        nw182_ = _dafny.Array(D6.default()(), 4)
        d_1153_v60_ = nw182_
        guard_loop_4_: int
        for guard_loop_4_ in _dafny.IntegerRange(0, (d_1153_v60_).length(0)):
            d_1154_i5_: int = guard_loop_4_
            if (True) and (((0) <= (d_1154_i5_)) and ((d_1154_i5_) < ((d_1153_v60_).length(0)))):
                (d_1153_v60_)[(d_1154_i5_)] = D6_DC15(_dafny.Set({p2}))
        d_1155_v61_: _dafny.Map
        d_1155_v61_ = _dafny.Map({p0: (self).f8})
        d_1155_v61_ = (d_1155_v61_).set(p0, (self).f8)
        d_1156_v62_: _dafny.Array
        def lambda67_(d_1157_i7_):
            return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xbox"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rvl")))

        init33_ = lambda67_
        nw183_ = _dafny.Array(None, 29)
        for i0_33_ in range(nw183_.length(0)):
            nw183_[i0_33_] = init33_(i0_33_)
        d_1156_v62_ = nw183_
        guard_loop_5_: int
        for guard_loop_5_ in _dafny.IntegerRange(0, (d_1156_v62_).length(0)):
            d_1158_i6_: int = guard_loop_5_
            if (True) and (((0) <= (d_1158_i6_)) and ((d_1158_i6_) < ((d_1156_v62_).length(0)))):
                (d_1156_v62_)[(d_1158_i6_)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jkqigycsa"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ysjgmd")))
        d_1159_v63_: D13
        d_1159_v63_ = D13_DC32(default__.safeModuloInt(975, p0), p1, (self).f7, p2, p1)
        source21_ = d_1159_v63_
        if source21_.is_DC32:
            d_1160___mcc_h23_ = source21_.cf56
            d_1161___mcc_h24_ = source21_.cf57
            d_1162___mcc_h25_ = source21_.cf58
            d_1163___mcc_h26_ = source21_.cf59
            d_1164___mcc_h27_ = source21_.cf60
            d_1165_cf60_ = d_1164___mcc_h27_
            d_1166_cf59_ = d_1163___mcc_h26_
            d_1167_cf58_ = d_1162___mcc_h25_
            d_1168_cf57_ = d_1161___mcc_h24_
            d_1169_cf56_ = d_1160___mcc_h23_
            d_1170_v64_: D8
            d_1170_v64_ = D8_DC21(d_1166_cf59_, d_1167_cf58_, len((p3).set(default__.safeIndex(p0, len(p3)), d_1167_cf58_)))
            (globalState).f5 = (d_1170_v64_).cf36
            d_1171_v65_: _dafny.Array
            def lambda68_(d_1172_i8_):
                return default__.safeDivisionInt(d_1172_i8_, 630)

            init34_ = lambda68_
            nw184_ = _dafny.Array(None, 19)
            for i0_34_ in range(nw184_.length(0)):
                nw184_[i0_34_] = init34_(i0_34_)
            d_1171_v65_ = nw184_
            index168_ = default__.safeIndex(335, (d_1171_v65_).length(0))
            (d_1171_v65_)[index168_] = -844
            d_1173_v67_: _dafny.Array
            def lambda69_(d_1174_p2_, d_1175_cf56_):
                def lambda70_(d_1176_i9_):
                    def iife130_():
                        coll54_ = _dafny.Set()
                        compr_54_: int
                        for compr_54_ in (_dafny.MultiSet([len(_dafny.Map({not(d_1174_p2_): d_1175_cf56_}))])).Elements:
                            d_1177_v66_: int = compr_54_
                            if (d_1177_v66_) in (_dafny.MultiSet([len(_dafny.Map({not(d_1174_p2_): d_1175_cf56_}))])):
                                coll54_ = coll54_.union(_dafny.Set([(d_1177_v66_) + (372)]))
                        return _dafny.Set(coll54_)
                    return (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kxvatj"))), len(iife130_()
                    )])) + (_dafny.SeqWithoutIsStrInference([d_1175_cf56_]))

                return lambda70_

            init35_ = lambda69_(p2, d_1169_cf56_)
            nw185_ = _dafny.Array(None, 6)
            for i0_35_ in range(nw185_.length(0)):
                nw185_[i0_35_] = init35_(i0_35_)
            d_1173_v67_ = nw185_
            d_1178_v68_: _dafny.Seq
            d_1178_v68_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gaq"))
            index169_ = default__.safeIndex(335, (d_1171_v65_).length(0))
            rhs172_ = d_1169_cf56_
            rhs173_ = not(True)
            rhs174_ = d_1173_v67_
            rhs175_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eakxl"))) + ((d_1178_v68_) + (d_1178_v68_))
            rhs176_ = (self).f7
            lhs134_ = d_1171_v65_
            lhs135_ = default__.safeIndex(335, (d_1171_v65_).length(0))
            lhs136_ = globalState
            lhs134_[lhs135_] = rhs172_
            r0 = rhs173_
            d_1173_v67_ = rhs174_
            d_1178_v68_ = rhs175_
            lhs136_.f5 = rhs176_
            d_1179_v69_: _dafny.Map
            d_1179_v69_ = _dafny.Map({d_1171_v65_: True})
            d_1180_v70_: int
            d_1181_v71_: int
            d_1182_v72_: bool
            d_1183_v73_: int
            out8_: int
            out9_: int
            out10_: bool
            out11_: int
            out8_, out9_, out10_, out11_ = (self).m1(not(default__.fm1(d_1038_v2_, (p3)[default__.safeIndex(d_1169_cf56_, len(p3))], (0) - (len(d_1179_v69_)), False, globalState)), globalState)
            d_1180_v70_ = out8_
            d_1181_v71_ = out9_
            d_1182_v72_ = out10_
            d_1183_v73_ = out11_
            d_1171_v65_ = d_1171_v65_
        elif source21_.is_DC33:
            d_1184___mcc_h28_ = source21_.cf61
            d_1185___mcc_h29_ = source21_.cf62
            d_1186___mcc_h30_ = source21_.cf63
            d_1187___mcc_h31_ = source21_.cf64
            d_1188___mcc_h32_ = source21_.cf65
            d_1189_cf65_ = d_1188___mcc_h32_
            d_1190_cf64_ = d_1187___mcc_h31_
            d_1191_cf63_ = d_1186___mcc_h30_
            d_1192_cf62_ = d_1185___mcc_h29_
            d_1193_cf61_ = d_1184___mcc_h28_
            d_1194_v74_: str
            d_1194_v74_ = _dafny.CodePoint('u')
            d_1194_v74_ = d_1194_v74_
            d_1191_cf63_ = default__.safeDivisionInt(d_1192_cf62_, 566)
            d_1195_v75_: _dafny.Array
            nw186_ = _dafny.Array(None, 8)
            d_1195_v75_ = nw186_
            d_1196_v76_: _dafny.Seq
            d_1196_v76_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n"))
            d_1197_v77_: C2
            nw187_ = C2()
            nw187_.ctor__(d_1196_v76_, d_1193_cf61_)
            d_1197_v77_ = nw187_
            index170_ = default__.safeIndex(356, (d_1195_v75_).length(0))
            (d_1195_v75_)[index170_] = d_1197_v77_
            d_1198_v78_: _dafny.MultiSet
            d_1198_v78_ = _dafny.MultiSet([(self).f7])
            d_1199_v79_: _dafny.Seq
            d_1199_v79_ = _dafny.SeqWithoutIsStrInference([d_1197_v77_, d_1197_v77_])
            index171_ = default__.safeIndex(356, (d_1195_v75_).length(0))
            rhs177_ = (self).f8
            rhs178_ = d_1197_v77_
            rhs179_ = (d_1199_v79_)[default__.safeIndex(((d_1189_cf65_).f9) + (len(_dafny.Set({(d_1189_cf65_).f9, p0}))), len(d_1199_v79_))]
            rhs180_ = (d_1198_v78_).intersection(d_1198_v78_)
            lhs137_ = d_1195_v75_
            lhs138_ = default__.safeIndex(356, (d_1195_v75_).length(0))
            d_1194_v74_ = rhs177_
            lhs137_[lhs138_] = rhs178_
            d_1197_v77_ = rhs179_
            d_1198_v78_ = rhs180_
            if not(d_1190_cf64_):
                d_1200_v80_: C4
                nw188_ = C4()
                nw188_.ctor__((d_1189_cf65_).f9)
                d_1200_v80_ = nw188_
                d_1201_v81_: _dafny.Set
                d_1201_v81_ = _dafny.Set({d_1200_v80_, d_1200_v80_})
                (globalState).f5 = (d_1201_v81_).issubset(d_1201_v81_)
                d_1202_v82_: _dafny.Map
                d_1202_v82_ = _dafny.Map({d_1197_v77_.f12: (d_1197_v77_).f11})
                d_1203_v83_: _dafny.MultiSet
                d_1203_v83_ = _dafny.MultiSet([d_1189_cf65_])
                d_1204_v84_: _dafny.MultiSet
                d_1204_v84_ = _dafny.MultiSet([(d_1203_v83_).cardinality])
                d_1205_v85_: _dafny.Set
                d_1205_v85_ = _dafny.Set({(d_1204_v84_).cardinality, (self).fm6(d_1191_cf63_, d_1197_v77_.f12, d_1191_cf63_, globalState), d_1191_cf63_})
                d_1206_v86_: _dafny.Array
                nw189_ = _dafny.Array(None, 13)
                nw189_[int(0)] = p0
                nw189_[int(1)] = d_1191_cf63_
                nw189_[int(2)] = p0
                nw189_[int(3)] = (d_1189_cf65_).f9
                nw189_[int(4)] = len(p3)
                nw189_[int(5)] = ((self).fm5(globalState)) - (d_1191_cf63_)
                nw189_[int(6)] = 588
                nw189_[int(7)] = 882
                nw189_[int(8)] = p0
                nw189_[int(9)] = (d_1191_cf63_) + (d_1192_cf62_)
                nw189_[int(10)] = (d_1189_cf65_).f9
                nw189_[int(11)] = default__.safeDivisionInt(len(((d_1202_v82_)[d_1197_v77_.f12] if (d_1197_v77_.f12) in (d_1202_v82_) else ((d_1202_v82_)[p2] if (p2) in (d_1202_v82_) else (d_1197_v77_).f11))), d_1192_cf62_)
                nw189_[int(12)] = (d_1189_cf65_).fm17(((d_1189_cf65_).f10)[default__.safeIndex(d_1191_cf63_, len((d_1189_cf65_).f10))], d_1205_v85_, globalState)
                d_1206_v86_ = nw189_
                index172_ = default__.safeIndex(300, (d_1206_v86_).length(0))
                (d_1206_v86_)[index172_] = (d_1189_cf65_).f9
                index173_ = default__.safeIndex(300, (d_1206_v86_).length(0))
                (d_1206_v86_)[index173_] = len(d_1196_v76_)
                d_1207_v87_: _dafny.Map
                d_1207_v87_ = _dafny.Map({(0) - (d_1192_cf62_): (d_1197_v77_).f11})
                d_1196_v76_ = ((d_1197_v77_).f11) + (((d_1207_v87_)[(0) - ((d_1189_cf65_).f9)] if ((0) - ((d_1189_cf65_).f9)) in (d_1207_v87_) else (d_1197_v77_).f11))
                index174_ = default__.safeIndex(300, (d_1206_v86_).length(0))
                (d_1206_v86_)[index174_] = (p0) * ((d_1206_v86_)[default__.safeIndex(300, (d_1206_v86_).length(0))])
                (globalState).f5 = d_1193_cf61_
            elif True:
                d_1208_v88_: _dafny.Array
                def lambda71_(d_1209_cf65_):
                    def lambda72_(d_1210_i10_):
                        return default__.safeModuloInt(d_1210_i10_, (d_1209_cf65_).f9)

                    return lambda72_

                init36_ = lambda71_(d_1189_cf65_)
                nw190_ = _dafny.Array(None, 19)
                for i0_36_ in range(nw190_.length(0)):
                    nw190_[i0_36_] = init36_(i0_36_)
                d_1208_v88_ = nw190_
                d_1208_v88_ = d_1208_v88_
                (globalState).f5 = d_1197_v77_.f12
                d_1211_v89_: _dafny.Set
                d_1211_v89_ = _dafny.Set({(p0) + ((d_1189_cf65_).f9), p0, default__.fm15(-984, d_1193_cf61_, globalState)})
                d_1211_v89_ = d_1211_v89_
                d_1212_v90_: C1
                nw191_ = C1()
                nw191_.ctor__()
                d_1212_v90_ = nw191_
                d_1213_v91_: C3
                nw192_ = C3()
                nw192_.ctor__(d_1191_cf63_, (d_1189_cf65_).f10, (d_1189_cf65_).f9)
                d_1213_v91_ = nw192_
        elif source21_.is_DC31:
            d_1214___mcc_h33_ = source21_.cf55
            d_1215_cf55_ = d_1214___mcc_h33_
            d_1216_v92_: _dafny.Seq
            d_1216_v92_ = _dafny.SeqWithoutIsStrInference([default__.fm36((self).f7, (self).f7, (self).f7, _dafny.CodePoint('x'), globalState)])
            (globalState).f1 = not((d_1216_v92_) == (d_1216_v92_))
            d_1217_v93_: _dafny.Array
            def lambda73_(d_1218_i12_):
                return default__.safeModuloInt(d_1218_i12_, -109)

            init37_ = lambda73_
            nw193_ = _dafny.Array(None, 17)
            for i0_37_ in range(nw193_.length(0)):
                nw193_[i0_37_] = init37_(i0_37_)
            d_1217_v93_ = nw193_
            d_1219_v94_: _dafny.Map
            d_1219_v94_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([len(_dafny.Set({p1, (self).f7})) for d_1220_i11_ in range(default__.abs(910))]): d_1217_v93_})
            d_1221_v95_: D8
            d_1221_v95_ = D8_DC20((d_1219_v94_) | (d_1219_v94_))
            source22_ = d_1221_v95_
            if source22_.is_DC21:
                d_1222___mcc_h35_ = source22_.cf36
                d_1223___mcc_h36_ = source22_.cf37
                d_1224___mcc_h37_ = source22_.cf38
                d_1225_cf38_ = d_1224___mcc_h37_
                d_1226_cf37_ = d_1223___mcc_h36_
                d_1227_cf36_ = d_1222___mcc_h35_
                d_1228_v96_: _dafny.MultiSet
                d_1228_v96_ = _dafny.MultiSet([d_1227_cf36_])
                d_1229_v97_: _dafny.Seq
                d_1229_v97_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ngnbjrkg"))
                d_1230_v98_: _dafny.Map
                d_1230_v98_ = _dafny.Map({d_1229_v97_: d_1225_cf38_})
                d_1231_v99_: _dafny.Map
                d_1231_v99_ = _dafny.Map({d_1227_cf36_: d_1228_v96_})
                d_1232_v100_: _dafny.Map
                d_1232_v100_ = _dafny.Map({p1: p1})
                d_1233_v101_: _dafny.Seq
                d_1233_v101_ = _dafny.SeqWithoutIsStrInference([d_1228_v96_])
                d_1234_v102_: _dafny.Array
                nw194_ = _dafny.Array(None, 26)
                nw194_[int(0)] = d_1228_v96_
                nw194_[int(1)] = d_1228_v96_
                nw194_[int(2)] = (_dafny.MultiSet(p3)) | (d_1228_v96_)
                nw194_[int(3)] = d_1228_v96_
                nw194_[int(4)] = d_1228_v96_
                nw194_[int(5)] = (d_1228_v96_) | (d_1228_v96_)
                nw194_[int(6)] = _dafny.MultiSet([d_1227_cf36_, d_1226_cf37_, d_1226_cf37_])
                nw194_[int(7)] = d_1228_v96_
                nw194_[int(8)] = d_1228_v96_
                nw194_[int(9)] = default__.fm30(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ae")), (_dafny.Map({d_1229_v97_: d_1225_cf38_})).set(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "apq")), p0), (self).f7, globalState)
                nw194_[int(10)] = default__.fm30((d_1229_v97_).set(default__.safeIndex(d_1225_cf38_, len(d_1229_v97_)), _dafny.CodePoint('k')), d_1230_v98_, p2, globalState)
                nw194_[int(11)] = d_1228_v96_
                nw194_[int(12)] = d_1228_v96_
                nw194_[int(13)] = d_1228_v96_
                nw194_[int(14)] = d_1228_v96_
                nw194_[int(15)] = (d_1228_v96_) | (((d_1231_v99_)[((d_1232_v100_)[(self).f7] if ((self).f7) in (d_1232_v100_) else False)] if (((d_1232_v100_)[(self).f7] if ((self).f7) in (d_1232_v100_) else False)) in (d_1231_v99_) else d_1228_v96_))
                nw194_[int(16)] = (d_1228_v96_).intersection(_dafny.MultiSet(p3))
                nw194_[int(17)] = d_1228_v96_
                nw194_[int(18)] = d_1228_v96_
                nw194_[int(19)] = _dafny.MultiSet(p3)
                nw194_[int(20)] = d_1228_v96_
                nw194_[int(21)] = d_1228_v96_
                nw194_[int(22)] = d_1228_v96_
                nw194_[int(23)] = d_1228_v96_
                nw194_[int(24)] = (d_1233_v101_)[default__.safeIndex(d_1225_cf38_, len(d_1233_v101_))]
                nw194_[int(25)] = d_1228_v96_
                d_1234_v102_ = nw194_
                index175_ = default__.safeIndex(165, (d_1234_v102_).length(0))
                (d_1234_v102_)[index175_] = d_1228_v96_
                index176_ = default__.safeIndex(165, (d_1234_v102_).length(0))
                (d_1234_v102_)[index176_] = d_1228_v96_
                d_1235_v103_: C2
                nw195_ = C2()
                nw195_.ctor__((d_1229_v97_).set(default__.safeIndex(d_1225_cf38_, len(d_1229_v97_)), (self).f8), d_1227_cf36_)
                d_1235_v103_ = nw195_
                d_1236_v104_: _dafny.Map
                d_1236_v104_ = _dafny.Map({p2: d_1235_v103_})
                (globalState).f0 = len(d_1236_v104_)
                d_1237_v105_: C1
                nw196_ = C1()
                nw196_.ctor__()
                d_1237_v105_ = nw196_
                d_1238_v106_: _dafny.Seq
                d_1238_v106_ = _dafny.SeqWithoutIsStrInference([(d_1237_v105_ if p2 else d_1237_v105_)])
                d_1237_v105_ = (d_1238_v106_)[default__.safeIndex(909, len(d_1238_v106_))]
                d_1239_v107_: _dafny.Array
                def lambda74_(d_1240_v103_):
                    def lambda75_(d_1241_i13_):
                        return d_1240_v103_.f12

                    return lambda75_

                init38_ = lambda74_(d_1235_v103_)
                nw197_ = _dafny.Array(None, 6)
                for i0_38_ in range(nw197_.length(0)):
                    nw197_[i0_38_] = init38_(i0_38_)
                d_1239_v107_ = nw197_
                d_1239_v107_ = d_1239_v107_
            elif True:
                d_1242___mcc_h38_ = source22_.cf35
                d_1243_cf35_ = d_1242___mcc_h38_
                d_1244_v108_: D1
                d_1244_v108_ = D1_DC4(_dafny.SeqWithoutIsStrInference([p0 for d_1245_i14_ in range(default__.abs(986))]))
                d_1246_v109_: C3
                nw198_ = C3()
                nw198_.ctor__(p0, (d_1244_v108_).cf11, default__.safeModuloInt(412, (0) - (-262)))
                d_1246_v109_ = nw198_
                index177_ = default__.safeIndex(179, (d_1217_v93_).length(0))
                (d_1217_v93_)[index177_] = (d_1246_v109_).f9
                index178_ = default__.safeIndex(179, (d_1217_v93_).length(0))
                (d_1217_v93_)[index178_] = len(_dafny.Set({366, (d_1246_v109_).f9, ((d_1246_v109_).f10)[default__.safeIndex((d_1246_v109_).f9, len((d_1246_v109_).f10))], p0, (d_1246_v109_).f9}))
                d_1247_v110_: C3
                nw199_ = C3()
                nw199_.ctor__((d_1246_v109_).f9, _dafny.SeqWithoutIsStrInference([(d_1246_v109_).f9]), (d_1217_v93_)[default__.safeIndex(179, (d_1217_v93_).length(0))])
                d_1247_v110_ = nw199_
                d_1248_v111_: _dafny.Map
                d_1248_v111_ = _dafny.Map({(0) - ((d_1217_v93_)[default__.safeIndex(179, (d_1217_v93_).length(0))]): len(p3)})
                d_1248_v111_ = d_1248_v111_
            d_1217_v93_ = (d_1217_v93_ if p1 else d_1217_v93_)
            (globalState).f5 = p2
        elif True:
            d_1249___mcc_h34_ = source21_.cf66
            d_1250_cf66_ = d_1249___mcc_h34_
            d_1251_v112_: _dafny.Seq
            d_1251_v112_ = _dafny.SeqWithoutIsStrInference([77])
            d_1252_v113_: D1
            d_1252_v113_ = D1_DC4(d_1251_v112_)
            d_1253_v114_: _dafny.Array
            nw200_ = _dafny.Array(int(0), 27)
            d_1253_v114_ = nw200_
            d_1254_v115_: D3
            d_1254_v115_ = D3_DC8(d_1252_v113_, d_1253_v114_, p0)
            d_1255_v116_: _dafny.MultiSet
            d_1255_v116_ = _dafny.MultiSet([p0, len(_dafny.SeqWithoutIsStrInference([p2, not(True), p2]))])
            d_1256_v117_: C3
            nw201_ = C3()
            nw201_.ctor__((d_1254_v115_).cf16, _dafny.SeqWithoutIsStrInference([p0, (0) - ((d_1255_v116_).cardinality), p0]), p0)
            d_1256_v117_ = nw201_
            (globalState).f5 = ((0) - ((d_1256_v117_).f9)) != (p0)
            d_1257_v118_: _dafny.Seq
            d_1257_v118_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lhbgnik"))
            rhs181_ = ((d_1251_v112_) + ((d_1256_v117_).f10)) + (_dafny.SeqWithoutIsStrInference([(d_1256_v117_).f9]))
            rhs182_ = p0
            rhs183_ = True
            rhs184_ = ((d_1256_v117_).f9) != (len(_dafny.SeqWithoutIsStrInference([(self).f8 for d_1258_i15_ in range(default__.abs(193))])))
            rhs185_ = (default__.fm2(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f")), globalState)) != ((d_1257_v118_) + (d_1257_v118_))
            lhs139_ = globalState
            lhs140_ = globalState
            lhs141_ = globalState
            lhs142_ = globalState
            d_1251_v112_ = rhs181_
            lhs139_.f2 = rhs182_
            lhs140_.f1 = rhs183_
            lhs141_.f5 = rhs184_
            lhs142_.f5 = rhs185_
            (globalState).f0 = (((0) - (len(d_1257_v118_))) + ((d_1256_v117_).f9)) + (((d_1255_v116_).cardinality) + (108))
        r0 = p2
        return r0

    @property
    def f7(self):
        return self._f7
    @property
    def f8(self):
        return self._f8

class C6(T0):
    def  __init__(self):
        self._f6: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C6"
    @property
    def f6(self):
        return self._f6
    def ctor__(self, f6):
        (self)._f6 = f6

    def fm3(self, p0, globalState):
        return default__.safeDivisionInt(default__.safeDivisionInt((self).f6, len(_dafny.SeqWithoutIsStrInference([_dafny.Map({True: False})]))), ((self).f6) * (len(_dafny.SeqWithoutIsStrInference([166]))))

    def m0(self, globalState):
        r0: int = int(0)
        r1: bool = False
        r2: bool = False
        d_1259_v0_: str
        d_1259_v0_ = _dafny.CodePoint('a')
        d_1260_v1_: _dafny.Seq
        d_1260_v1_ = _dafny.SeqWithoutIsStrInference([d_1259_v0_, d_1259_v0_])
        d_1261_v2_: _dafny.Map
        d_1261_v2_ = _dafny.Map({(self).f6: (0) - (len(_dafny.SeqWithoutIsStrInference([(self).f6, len(d_1260_v1_)])))})
        d_1262_v3_: bool
        d_1262_v3_ = True
        d_1263_v4_: _dafny.Seq
        d_1263_v4_ = _dafny.SeqWithoutIsStrInference([(d_1261_v2_) != (d_1261_v2_), d_1262_v3_, d_1262_v3_, d_1262_v3_])
        d_1263_v4_ = (d_1263_v4_).set(default__.safeIndex((self).f6, len(d_1263_v4_)), not((d_1263_v4_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_1262_v3_])), len(d_1263_v4_))]))
        hi8_ = (self).f6
        for d_1264_i0_ in range((self).f6, hi8_):
            d_1265_v5_: D0
            d_1265_v5_ = D0_DC2(d_1264_i0_, d_1259_v0_, 618, (self).f6, d_1264_i0_)
            d_1266_v6_: _dafny.Set
            d_1266_v6_ = _dafny.Set({d_1262_v3_, d_1262_v3_})
            d_1267_v7_: _dafny.Array
            nw202_ = _dafny.Array(None, 14)
            nw202_[int(0)] = d_1259_v0_
            nw202_[int(1)] = _dafny.CodePoint('o')
            nw202_[int(2)] = d_1259_v0_
            nw202_[int(3)] = d_1259_v0_
            nw202_[int(4)] = d_1259_v0_
            nw202_[int(5)] = d_1259_v0_
            nw202_[int(6)] = d_1259_v0_
            nw202_[int(7)] = d_1259_v0_
            nw202_[int(8)] = d_1259_v0_
            nw202_[int(9)] = d_1259_v0_
            nw202_[int(10)] = (d_1259_v0_ if d_1262_v3_ else d_1259_v0_)
            nw202_[int(11)] = d_1259_v0_
            nw202_[int(12)] = d_1259_v0_
            nw202_[int(13)] = default__.fm4(_dafny.Map({_dafny.SeqWithoutIsStrInference([(0) - (d_1264_i0_)]): d_1265_v5_}), (0) - (d_1264_i0_), len(d_1266_v6_), globalState)
            d_1267_v7_ = nw202_
            index179_ = default__.safeIndex(387, (d_1267_v7_).length(0))
            (d_1267_v7_)[index179_] = d_1259_v0_
            index180_ = default__.safeIndex(387, (d_1267_v7_).length(0))
            (d_1267_v7_)[index180_] = d_1259_v0_
            (globalState).f5 = False
            if True:
                (globalState).f0 = ((self).f6 if d_1262_v3_ else (self).f6)
                d_1268_v8_: _dafny.Set
                d_1268_v8_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([664, d_1264_i0_, d_1264_i0_, (self).f6]), _dafny.SeqWithoutIsStrInference([d_1264_i0_ for d_1269_i1_ in range(default__.abs(144))])})
                d_1270_v9_: _dafny.Map
                d_1270_v9_ = _dafny.Map({not((d_1268_v8_).issubset(d_1268_v8_)): d_1266_v6_})
                d_1270_v9_ = (d_1270_v9_) | (d_1270_v9_)
                d_1271_v10_: _dafny.Map
                d_1271_v10_ = _dafny.Map({d_1264_i0_: d_1262_v3_})
                d_1272_v11_: _dafny.Map
                d_1272_v11_ = _dafny.Map({(187) + (len(d_1271_v10_)): d_1262_v3_})
                d_1272_v11_ = (d_1272_v11_).set(default__.safeDivisionInt(d_1264_i0_, (self).f6), not(d_1262_v3_))
                d_1271_v10_ = (d_1271_v10_).set((d_1265_v5_).cf9, ((self).f6) == ((self).f6))
                d_1273_v12_: _dafny.Array
                nw203_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 26)
                d_1273_v12_ = nw203_
                d_1273_v12_ = d_1273_v12_
            elif True:
                d_1274_v13_: C1
                nw204_ = C1()
                nw204_.ctor__()
                d_1274_v13_ = nw204_
                d_1275_v14_: C0
                nw205_ = C0()
                nw205_.ctor__(True)
                d_1275_v14_ = nw205_
                d_1276_v15_: _dafny.Map
                d_1276_v15_ = _dafny.Map({d_1262_v3_: d_1275_v14_})
                d_1277_v16_: _dafny.Map
                d_1277_v16_ = _dafny.Map({d_1276_v15_: (0) - (d_1264_i0_)})
                d_1278_v17_: _dafny.Map
                d_1278_v17_ = _dafny.Map({len(d_1277_v16_): (d_1275_v14_).f13})
                d_1278_v17_ = (d_1278_v17_).set((self).f6, (d_1275_v14_).f13)
                (globalState).f2 = len((_dafny.Map({not(d_1262_v3_): _dafny.CodePoint('y')})).set((d_1275_v14_).f13, (d_1267_v7_)[default__.safeIndex(387, (d_1267_v7_).length(0))]))
                d_1279_v18_: _dafny.Set
                d_1279_v18_ = _dafny.Set({d_1264_i0_})
                d_1280_v19_: _dafny.Seq
                d_1280_v19_ = _dafny.SeqWithoutIsStrInference([d_1264_i0_, len(d_1279_v18_)])
                d_1281_v20_: C3
                nw206_ = C3()
                nw206_.ctor__(d_1264_i0_, d_1280_v19_, (self).f6)
                d_1281_v20_ = nw206_
                rhs186_ = (D13_DC33((d_1275_v14_).f13, (0) - (len(d_1261_v2_)), d_1264_i0_, (d_1275_v14_).f13, d_1281_v20_)).cf64
                rhs187_ = (self).f6
                lhs143_ = globalState
                lhs144_ = globalState
                lhs143_.f5 = rhs186_
                lhs144_.f0 = rhs187_
                (globalState).f1 = not((d_1275_v14_).f13)
            d_1260_v1_ = d_1260_v1_
        d_1282_v21_: _dafny.Map
        d_1282_v21_ = _dafny.Map({d_1262_v3_: _dafny.SeqWithoutIsStrInference([d_1259_v0_ for d_1283_i2_ in range(default__.abs(707))])})
        d_1284_v22_: _dafny.MultiSet
        d_1284_v22_ = _dafny.MultiSet([((d_1282_v21_)[d_1262_v3_] if (d_1262_v3_) in (d_1282_v21_) else d_1260_v1_), (d_1260_v1_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gfyelbune")))])
        d_1284_v22_ = ((d_1284_v22_) - (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_1259_v0_ for d_1285_i3_ in range(default__.abs(471))])]))).set(d_1260_v1_, default__.abs((self).f6))
        d_1286_v23_: _dafny.MultiSet
        d_1286_v23_ = _dafny.MultiSet([d_1262_v3_, d_1262_v3_, not(d_1262_v3_), d_1262_v3_])
        if (d_1286_v23_).isdisjoint((d_1286_v23_) | (d_1286_v23_)):
            d_1287_v24_: _dafny.Map
            d_1287_v24_ = _dafny.Map({d_1262_v3_: d_1262_v3_})
            d_1288_v25_: _dafny.Seq
            d_1288_v25_ = _dafny.SeqWithoutIsStrInference([(self).f6, len(d_1287_v24_)])
            d_1289_v26_: _dafny.Set
            d_1289_v26_ = _dafny.Set({len(d_1288_v25_), (self).f6, (self).f6})
            d_1290_v27_: C1
            nw207_ = C1()
            nw207_.ctor__()
            d_1290_v27_ = nw207_
            d_1291_v28_: _dafny.Map
            d_1291_v28_ = _dafny.Map({(_dafny.Set({(d_1288_v25_)[default__.safeIndex((self).f6, len(d_1288_v25_))]})).ispropersubset(d_1289_v26_): d_1290_v27_})
            d_1291_v28_ = (d_1291_v28_).set(d_1262_v3_, d_1290_v27_)
            d_1292_v29_: _dafny.Map
            d_1292_v29_ = _dafny.Map({d_1262_v3_: (self).f6})
            d_1293_v30_: _dafny.Map
            d_1293_v30_ = _dafny.Map({d_1292_v29_: d_1262_v3_})
            rhs188_ = ((self).f6) * (len(d_1293_v30_))
            rhs189_ = (self).f6
            rhs190_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "npsdnb"))).set(default__.safeIndex(954, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "npsdnb")))), d_1259_v0_)
            lhs145_ = globalState
            lhs146_ = globalState
            lhs145_.f0 = rhs188_
            lhs146_.f0 = rhs189_
            d_1260_v1_ = rhs190_
            d_1294_v32_: D0
            d_1294_v32_ = D0_DC2((self).f6, d_1259_v0_, (self).f6, 101, ((d_1261_v2_)[-905] if (-905) in (d_1261_v2_) else -20))
            d_1295_v33_: _dafny.Set
            d_1295_v33_ = _dafny.Set({d_1259_v0_, (d_1294_v32_).cf6})
            def iife131_():
                coll55_ = _dafny.Map()
                compr_55_: str
                for compr_55_ in (d_1295_v33_).Elements:
                    d_1296_v31_: str = compr_55_
                    if (d_1296_v31_) in (d_1295_v33_):
                        coll55_[d_1296_v31_] = (self).f6
                return _dafny.Map(coll55_)
            (globalState).f5 = ((d_1259_v0_ if not(not(not(d_1262_v3_))) else d_1259_v0_)) not in ((iife131_()
            ) | (_dafny.Map({d_1259_v0_: (self).f6})))
            d_1287_v24_ = (d_1287_v24_).set(d_1262_v3_, d_1262_v3_)
            d_1297_v34_: int
            d_1298_v35_: int
            d_1299_v36_: bool
            out12_: int
            out13_: int
            out14_: bool
            out12_, out13_, out14_ = (d_1290_v27_).m6(globalState)
            d_1297_v34_ = out12_
            d_1298_v35_ = out13_
            d_1299_v36_ = out14_
        elif True:
            r0 = 569
            if (d_1262_v3_) == (d_1262_v3_):
                d_1300_v37_: _dafny.Map
                d_1300_v37_ = _dafny.Map({d_1262_v3_: (self).f6})
                d_1301_v38_: _dafny.Map
                d_1301_v38_ = _dafny.Map({len(d_1300_v37_): d_1262_v3_})
                d_1302_v39_: _dafny.MultiSet
                d_1302_v39_ = _dafny.MultiSet([d_1301_v38_, d_1301_v38_])
                d_1303_v40_: _dafny.Seq
                d_1303_v40_ = _dafny.SeqWithoutIsStrInference([d_1302_v39_])
                (globalState).f5 = (d_1301_v38_) in ((d_1303_v40_)[default__.safeIndex((self).f6, len(d_1303_v40_))])
                d_1300_v37_ = (d_1300_v37_).set(d_1262_v3_, (self).f6)
                (globalState).f0 = default__.fm24(len(d_1260_v1_), globalState)
                d_1304_v41_: D12
                d_1304_v41_ = D12_DC30(default__.fm37(d_1262_v3_, (self).f6, d_1286_v23_, globalState))
                d_1305_v42_: D12
                d_1305_v42_ = D12_DC28(d_1300_v37_)
                d_1304_v41_ = D12_DC30(d_1305_v42_)
                d_1306_v43_: _dafny.Array
                nw208_ = _dafny.Array(_dafny.Array(None, 0), 28)
                d_1306_v43_ = nw208_
                d_1307_v44_: _dafny.Array
                nw209_ = _dafny.Array(_dafny.MultiSet({}), 7)
                d_1307_v44_ = nw209_
                index181_ = default__.safeIndex(324, (d_1306_v43_).length(0))
                (d_1306_v43_)[index181_] = d_1307_v44_
                index182_ = default__.safeIndex(324, (d_1306_v43_).length(0))
                nw210_ = _dafny.Array(_dafny.MultiSet({}), 23)
                (d_1306_v43_)[index182_] = nw210_
            elif True:
                d_1308_v45_: _dafny.Set
                d_1308_v45_ = _dafny.Set({d_1262_v3_, d_1262_v3_})
                d_1309_v46_: _dafny.Map
                d_1309_v46_ = _dafny.Map({(self).f6: d_1308_v45_})
                d_1310_v47_: D0
                d_1310_v47_ = D0_DC0(d_1309_v46_)
                d_1311_v48_: C1
                nw211_ = C1()
                nw211_.ctor__()
                d_1311_v48_ = nw211_
                d_1312_v49_: _dafny.Set
                d_1312_v49_ = _dafny.Set({d_1259_v0_})
                d_1313_v50_: _dafny.Map
                d_1313_v50_ = _dafny.Map({d_1311_v48_: len(d_1312_v49_)})
                d_1314_v51_: _dafny.MultiSet
                d_1314_v51_ = _dafny.MultiSet([_dafny.Set({not(d_1262_v3_), d_1262_v3_})])
                d_1315_v52_: _dafny.Map
                d_1315_v52_ = _dafny.Map({default__.fm1(d_1310_v47_, d_1262_v3_, len(d_1313_v50_), d_1262_v3_, globalState): (d_1314_v51_) | (d_1314_v51_)})
                d_1315_v52_ = (d_1315_v52_).set(d_1262_v3_, d_1314_v51_)
                d_1316_v53_: _dafny.Set
                d_1316_v53_ = _dafny.Set({(default__.fm2(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bq")), globalState)) + (d_1260_v1_)})
                d_1317_v54_: _dafny.Map
                d_1317_v54_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_1259_v0_ for d_1318_i4_ in range(default__.abs(930))]): -965})
                def iife132_():
                    coll56_ = _dafny.Set()
                    compr_56_: _dafny.Seq
                    for compr_56_ in (d_1317_v54_).keys.Elements:
                        d_1319_v55_: _dafny.Seq = compr_56_
                        if (d_1319_v55_) in (d_1317_v54_):
                            coll56_ = coll56_.union(_dafny.Set([d_1319_v55_]))
                    return _dafny.Set(coll56_)
                d_1316_v53_ = (iife132_()
                 if d_1262_v3_ else d_1316_v53_)
                (globalState).f1 = d_1262_v3_
                d_1320_v56_: _dafny.Seq
                d_1320_v56_ = _dafny.SeqWithoutIsStrInference([(self).f6, (0) - (default__.safeModuloInt((self).f6, (self).f6))])
                d_1321_v57_: _dafny.Array
                nw212_ = _dafny.Array(None, 6)
                nw212_[int(0)] = d_1259_v0_
                nw212_[int(1)] = d_1259_v0_
                nw212_[int(2)] = d_1259_v0_
                nw212_[int(3)] = d_1259_v0_
                nw212_[int(4)] = d_1259_v0_
                nw212_[int(5)] = d_1259_v0_
                d_1321_v57_ = nw212_
                d_1322_v58_: _dafny.MultiSet
                d_1322_v58_ = _dafny.MultiSet([d_1321_v57_, d_1321_v57_])
                d_1323_v59_: _dafny.Map
                d_1323_v59_ = _dafny.Map({(d_1263_v4_)[default__.safeIndex((self).f6, len(d_1263_v4_))]: d_1262_v3_})
                rhs191_ = (d_1320_v56_).set(default__.safeIndex(len(((d_1260_v1_).set(default__.safeIndex((0) - ((d_1322_v58_).cardinality), len(d_1260_v1_)), d_1259_v0_)) + ((d_1260_v1_).set(default__.safeIndex(len(d_1320_v56_), len(d_1260_v1_)), d_1259_v0_))), len(d_1320_v56_)), len(d_1323_v59_))
                rhs192_ = ((self).f6 if d_1262_v3_ else (self).f6)
                lhs147_ = globalState
                d_1320_v56_ = rhs191_
                lhs147_.f2 = rhs192_
                d_1324_v60_: _dafny.Array
                nw213_ = _dafny.Array(_dafny.Array(None, 0), 19)
                d_1324_v60_ = nw213_
                d_1325_v61_: _dafny.Array
                nw214_ = _dafny.Array(int(0), 1)
                d_1325_v61_ = nw214_
                index183_ = default__.safeIndex(426, (d_1324_v60_).length(0))
                (d_1324_v60_)[index183_] = d_1325_v61_
                index184_ = default__.safeIndex(426, (d_1324_v60_).length(0))
                nw215_ = _dafny.Array(int(0), 10)
                (d_1324_v60_)[index184_] = nw215_
            d_1326_v62_: _dafny.Map
            d_1326_v62_ = _dafny.Map({d_1262_v3_: _dafny.SeqWithoutIsStrInference([d_1262_v3_])})
            d_1327_v63_: _dafny.Map
            d_1327_v63_ = _dafny.Map({(((d_1326_v62_)[d_1262_v3_] if (d_1262_v3_) in (d_1326_v62_) else _dafny.SeqWithoutIsStrInference([d_1262_v3_]))).set(default__.safeIndex((self).f6, len(((d_1326_v62_)[d_1262_v3_] if (d_1262_v3_) in (d_1326_v62_) else _dafny.SeqWithoutIsStrInference([d_1262_v3_])))), d_1262_v3_): d_1262_v3_})
            d_1327_v63_ = (d_1327_v63_).set(_dafny.SeqWithoutIsStrInference([d_1262_v3_, (d_1263_v4_)[default__.safeIndex((self).f6, len(d_1263_v4_))], d_1262_v3_, d_1262_v3_, d_1262_v3_]), d_1262_v3_)
            if d_1262_v3_:
                (globalState).f1 = d_1262_v3_
                r2 = d_1262_v3_
                d_1328_v64_: _dafny.MultiSet
                d_1328_v64_ = _dafny.MultiSet([(self).f6, 201])
                r1 = ((default__.fm15((self).f6, d_1262_v3_, globalState)) * (-22)) > (((d_1328_v64_)[(0) - ((self).f6)] if ((0) - ((self).f6)) in (d_1328_v64_) else (self).f6))
                d_1329_v65_: T1
                nw216_ = C1()
                nw216_.ctor__()
                d_1329_v65_ = nw216_
                d_1329_v65_ = (d_1329_v65_ if not(d_1262_v3_) else d_1329_v65_)
                d_1330_v66_: C2
                nw217_ = C2()
                nw217_.ctor__(d_1260_v1_, d_1262_v3_)
                d_1330_v66_ = nw217_
            elif True:
                (globalState).f5 = d_1262_v3_
                d_1331_v67_: _dafny.Map
                d_1331_v67_ = _dafny.Map({d_1259_v0_: (self).f6})
                d_1332_v68_: D17
                d_1332_v68_ = D17_DC40(d_1331_v67_)
                d_1333_v69_: _dafny.MultiSet
                d_1333_v69_ = _dafny.MultiSet([((d_1332_v68_).cf76) | (d_1331_v67_), d_1331_v67_])
                d_1333_v69_ = (d_1333_v69_).set(d_1331_v67_, default__.abs(152))
                d_1334_v70_: _dafny.MultiSet
                d_1334_v70_ = _dafny.MultiSet([(self).f6])
                d_1335_v71_: _dafny.Seq
                d_1335_v71_ = _dafny.SeqWithoutIsStrInference([d_1334_v70_])
                (globalState).f1 = (_dafny.SeqWithoutIsStrInference([d_1334_v70_])) < (d_1335_v71_)
                d_1336_v72_: _dafny.Map
                d_1336_v72_ = _dafny.Map({(0) - ((self).f6): True})
                d_1336_v72_ = d_1336_v72_
                d_1337_v73_: _dafny.Set
                d_1337_v73_ = _dafny.Set({(self).f6, len(d_1260_v1_), (self).f6})
                d_1338_v74_: _dafny.Seq
                d_1338_v74_ = _dafny.SeqWithoutIsStrInference([(self).f6])
                d_1339_v75_: T0
                nw218_ = C3()
                nw218_.ctor__(len(d_1337_v73_), d_1338_v74_, (self).f6)
                d_1339_v75_ = nw218_
                d_1340_v76_: _dafny.Map
                d_1340_v76_ = _dafny.Map({d_1262_v3_: d_1339_v75_})
                d_1341_v77_: _dafny.Map
                d_1341_v77_ = _dafny.Map({d_1340_v76_: (d_1339_v75_).f6})
                d_1342_v78_: _dafny.Map
                d_1342_v78_ = _dafny.Map({d_1262_v3_: d_1262_v3_})
                (globalState).f1 = ((len(d_1341_v77_)) + ((d_1339_v75_).f6)) < (len(d_1342_v78_))
            d_1259_v0_ = d_1259_v0_
        d_1343_v79_: _dafny.Seq
        d_1343_v79_ = _dafny.SeqWithoutIsStrInference([(self).f6])
        d_1344_v80_: _dafny.Set
        d_1344_v80_ = _dafny.Set({(self).f6, (self).f6, len((D1_DC4(d_1343_v79_)).cf11), 621, -963})
        d_1262_v3_ = not((True) or ((d_1344_v80_) == (_dafny.Set({(self).f6, (0) - ((self).f6)}))))
        d_1345_v81_: _dafny.Map
        d_1345_v81_ = _dafny.Map({d_1262_v3_: d_1343_v79_})
        d_1346_v82_: C3
        nw219_ = C3()
        nw219_.ctor__(default__.fm15((self).f6, d_1262_v3_, globalState), ((d_1345_v81_)[d_1262_v3_] if (d_1262_v3_) in (d_1345_v81_) else _dafny.SeqWithoutIsStrInference([(self).f6, (self).f6])), (self).f6)
        d_1346_v82_ = nw219_
        d_1347_v83_: _dafny.Set
        d_1347_v83_ = _dafny.Set({d_1262_v3_})
        d_1348_v84_: _dafny.Map
        d_1348_v84_ = _dafny.Map({(self).f6: d_1347_v83_})
        d_1349_v85_: D0
        d_1349_v85_ = D0_DC0(d_1348_v84_)
        d_1350_v86_: _dafny.Map
        d_1350_v86_ = _dafny.Map({d_1262_v3_: d_1262_v3_})
        d_1351_v87_: _dafny.Array
        nw220_ = _dafny.Array(None, 26)
        nw220_[int(0)] = d_1262_v3_
        nw220_[int(1)] = d_1262_v3_
        nw220_[int(2)] = (d_1262_v3_ if d_1262_v3_ else not(d_1262_v3_))
        nw220_[int(3)] = False
        nw220_[int(4)] = (D13_DC33(True, len(d_1261_v2_), (self).f6, d_1262_v3_, d_1346_v82_)).cf61
        nw220_[int(5)] = d_1262_v3_
        nw220_[int(6)] = (d_1286_v23_) == (d_1286_v23_)
        nw220_[int(7)] = d_1262_v3_
        nw220_[int(8)] = d_1262_v3_
        nw220_[int(9)] = (d_1344_v80_).isdisjoint(d_1344_v80_)
        nw220_[int(10)] = (d_1262_v3_) == (d_1262_v3_)
        nw220_[int(11)] = d_1262_v3_
        nw220_[int(12)] = d_1262_v3_
        nw220_[int(13)] = (d_1263_v4_)[default__.safeIndex(default__.fm24((self).f6, globalState), len(d_1263_v4_))]
        nw220_[int(14)] = (d_1346_v82_).fm16(d_1286_v23_, globalState)
        nw220_[int(15)] = (d_1262_v3_ if not(d_1262_v3_) else d_1262_v3_)
        nw220_[int(16)] = d_1262_v3_
        nw220_[int(17)] = d_1262_v3_
        nw220_[int(18)] = d_1262_v3_
        nw220_[int(19)] = d_1262_v3_
        nw220_[int(20)] = default__.fm1(d_1349_v85_, d_1262_v3_, (d_1346_v82_).f9, d_1262_v3_, globalState)
        nw220_[int(21)] = not (d_1262_v3_) or (not(d_1262_v3_))
        nw220_[int(22)] = (d_1262_v3_ if default__.fm1(d_1349_v85_, d_1262_v3_, (d_1346_v82_).f9, True, globalState) else True)
        nw220_[int(23)] = (d_1263_v4_)[default__.safeIndex((self).f6, len(d_1263_v4_))]
        nw220_[int(24)] = d_1262_v3_
        nw220_[int(25)] = (((d_1350_v86_)[d_1262_v3_] if (d_1262_v3_) in (d_1350_v86_) else False) if d_1262_v3_ else d_1262_v3_)
        d_1351_v87_ = nw220_
        d_1352_v88_: _dafny.Array
        nw221_ = _dafny.Array(None, 24)
        d_1352_v88_ = nw221_
        d_1353_v89_: _dafny.MultiSet
        d_1353_v89_ = _dafny.MultiSet([d_1352_v88_, d_1352_v88_, d_1352_v88_])
        index185_ = default__.safeIndex(235, (d_1351_v87_).length(0))
        (d_1351_v87_)[index185_] = (((d_1353_v89_)[d_1352_v88_] if (d_1352_v88_) in (d_1353_v89_) else (d_1346_v82_).f9)) != (len(d_1260_v1_))
        index186_ = default__.safeIndex(235, (d_1351_v87_).length(0))
        (d_1351_v87_)[index186_] = d_1262_v3_
        d_1354_v90_: _dafny.MultiSet
        d_1354_v90_ = _dafny.MultiSet([(d_1346_v82_).fm17((d_1346_v82_).f9, d_1344_v80_, globalState)])
        d_1355_v91_: _dafny.Map
        d_1355_v91_ = _dafny.Map({D16_DC38(d_1354_v90_): d_1344_v80_})
        r0 = default__.safeModuloInt((0) - ((d_1346_v82_).f9), len(d_1355_v91_))
        r1 = d_1262_v3_
        r2 = (d_1286_v23_).ispropersubset(d_1286_v23_)
        return r0, r1, r2

