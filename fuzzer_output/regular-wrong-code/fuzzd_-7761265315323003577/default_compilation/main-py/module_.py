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
        source0_ = D1_DC3(_dafny.MultiSet([_dafny.CodePoint('s'), _dafny.CodePoint('l'), _dafny.CodePoint('q')]))
        if source0_.is_DC4:
            d_0___mcc_h0_ = source0_.cf7
            d_1___mcc_h1_ = source0_.cf8
            d_2___mcc_h2_ = source0_.cf9
            d_3___mcc_h3_ = source0_.cf10
            d_4___mcc_h4_ = source0_.cf11
            d_5_cf11_ = d_4___mcc_h4_
            d_6_cf10_ = d_3___mcc_h3_
            d_7_cf9_ = d_2___mcc_h2_
            d_8_cf8_ = d_1___mcc_h1_
            d_9_cf7_ = d_0___mcc_h0_
            return (d_7_cf9_)[default__.safeIndex(len(d_7_cf9_), len(d_7_cf9_))]
        elif True:
            d_10___mcc_h5_ = source0_.cf6
            d_11_cf6_ = d_10___mcc_h5_
            return _dafny.CodePoint('n')

    @staticmethod
    def fm1(globalState):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: int
            for compr_0_ in (_dafny.MultiSet([-912])).Elements:
                d_12_v0_: int = compr_0_
                if (d_12_v0_) in (_dafny.MultiSet([-912])):
                    def iife1_():
                        coll1_ = _dafny.Map()
                        compr_1_: int
                        for compr_1_ in _dafny.IntegerRange(676, 222):
                            d_13_v1_: int = compr_1_
                            if ((676) <= (d_13_v1_)) and ((d_13_v1_) < (222)):
                                coll1_[(d_13_v1_) + (-910)] = True
                        return _dafny.Map(coll1_)
                    coll0_[default__.safeModuloInt(d_12_v0_, len(_dafny.Map({False: True})))] = len(iife1_()
                    )
            return _dafny.Map(coll0_)
        return _dafny.Set({(0) - (((_dafny.MultiSet([False, not(not(False)), False])).cardinality if False else len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({len(iife0_()
        )}))])))), (0) - ((len(_dafny.Map({_dafny.CodePoint('k'): False}))) * (-907)), ((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lk"))))) + (802)})

    @staticmethod
    def fm2(p0, globalState):
        return ((-235) - (-236)) * (663)

    @staticmethod
    def fm3(globalState):
        return False

    @staticmethod
    def fm4(p0, p1, p2, globalState):
        return ((_dafny.SeqWithoutIsStrInference([True, False])) + (_dafny.SeqWithoutIsStrInference([False, False]))) + (_dafny.SeqWithoutIsStrInference([False]))

    @staticmethod
    def fm9(globalState):
        def iife2_():
            coll2_ = _dafny.Map()
            compr_2_: int
            for compr_2_ in _dafny.IntegerRange(294, 803):
                d_15_v0_: int = compr_2_
                if ((294) <= (d_15_v0_)) and ((d_15_v0_) < (803)):
                    coll2_[(d_15_v0_) * ((_dafny.MultiSet([_dafny.MultiSet([len(_dafny.Map({len(_dafny.Set({True})): len(_dafny.SeqWithoutIsStrInference([True]))}))])])).cardinality)] = 781
            return _dafny.Map(coll2_)
        return ((_dafny.Map({len(_dafny.Map({(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(D11_DC33(-562, not(False), 369, False, _dafny.Set({True}))).cf65, False]))).cardinality: 412})): 706})) | (_dafny.Map({len(_dafny.Set({len(_dafny.Set({True, False}))})): 906}))) | ((_dafny.Map({len(_dafny.Set({False})): len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_14_i0_ in range(default__.abs(697))]))})) | (iife2_()
        ))

    @staticmethod
    def fm12(p0, p1, p2, globalState):
        return D2_DC8(False, not((_dafny.Set({90})).ispropersubset(_dafny.Set({894, (_dafny.MultiSet([False, True, False])).cardinality}))))

    @staticmethod
    def fm17(p0, p1, globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_16_i0_ in range(default__.abs(357))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fqjluaqn")))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pfpcdepv")))

    @staticmethod
    def fm18(p0, globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_17_i0_ in range(default__.abs(902))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "go")))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ywfd")))

    @staticmethod
    def fm19(p0, globalState):
        return (_dafny.MultiSet([(_dafny.MultiSet([400, (_dafny.MultiSet([False, False, not((D0_DC1(not(True), (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Set({not(True)}))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qxcxo")))])) for d_18_i0_ in range(default__.abs(-171))]))).cardinality, _dafny.CodePoint('h'))).cf2), not(False)])).cardinality])).cardinality, (_dafny.MultiSet([False])).cardinality, (_dafny.MultiSet([False, True, False, True])).cardinality])) - (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([-397 for d_19_i1_ in range(default__.abs(538))])))

    @staticmethod
    def fm20(p0, p1, p2, globalState):
        def iife3_():
            def iife5_():
                coll5_ = _dafny.Set()
                compr_5_: int
                for compr_5_ in _dafny.IntegerRange(-648, 965):
                    d_22_v1_: int = compr_5_
                    if ((-648) <= (d_22_v1_)) and ((d_22_v1_) < (965)):
                        coll5_ = coll5_.union(_dafny.Set([default__.safeDivisionInt(d_22_v1_, 422)]))
                return _dafny.Set(coll5_)
            coll3_ = _dafny.Map()
            def iife4_():
                coll4_ = _dafny.Set()
                compr_4_: int
                for compr_4_ in _dafny.IntegerRange(-648, 965):
                    d_20_v1_: int = compr_4_
                    if ((-648) <= (d_20_v1_)) and ((d_20_v1_) < (965)):
                        coll4_ = coll4_.union(_dafny.Set([default__.safeDivisionInt(d_20_v1_, 422)]))
                return _dafny.Set(coll4_)
            compr_3_: _dafny.Set
            for compr_3_ in (_dafny.SeqWithoutIsStrInference([iife4_()
            ])).Elements:
                d_21_v0_: _dafny.Set = compr_3_
                if (d_21_v0_) in (_dafny.SeqWithoutIsStrInference([iife5_()
                ])):
                    coll3_[d_21_v0_] = True
            return _dafny.Map(coll3_)
        return D0_DC1((-119) >= (len(_dafny.SeqWithoutIsStrInference([(0) - (len(iife3_()
))]))), len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qnlh"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_23_i0_ in range(default__.abs(886))]))), _dafny.CodePoint('r'))

    @staticmethod
    def fm25(p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_24_i0_ in range(default__.abs(761))])

    @staticmethod
    def fm26(p0, p1, p2, p3, globalState):
        return ((_dafny.SeqWithoutIsStrInference([-716, len(_dafny.Set({True}))])) + (_dafny.SeqWithoutIsStrInference([(0) - ((0) - (len(_dafny.Map({not(True): False})))) for d_25_i0_ in range(default__.abs(993))]))) + ((_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_26_i1_ in range(default__.abs(495))]))]) if True else _dafny.SeqWithoutIsStrInference([90, 979])))

    @staticmethod
    def fm27(globalState):
        return ((_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tqkyhgesq")): -595})) | (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bk")): 753}))) | (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mrd")): 673}))

    @staticmethod
    def fm29(p0, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yhclj"))])) + ((D21_DC49(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jcnirnr")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ttxfurlkd"))]))).cf88)

    @staticmethod
    def fm30(p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xty"))

    @staticmethod
    def fm31(p0, p1, p2, p3, globalState):
        return ((_dafny.Map({False: (D9_DC25(-77, (0) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tussinc"))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tk")))]))), len(_dafny.SeqWithoutIsStrInference([False, True])), _dafny.CodePoint('n'))).cf47})) | (_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference([10]))}))) | (_dafny.Map({False: -344}))

    @staticmethod
    def fm32(p0, p1, p2, p3, globalState):
        if not((False) or (False)):
            return D0_DC0(_dafny.SeqWithoutIsStrInference([not(False), False]), 468)
        elif True:
            return D0_DC0(_dafny.SeqWithoutIsStrInference([True]), 501)

    @staticmethod
    def fm33(p0, p1, p2, p3, globalState):
        def iife6_():
            coll6_ = _dafny.Map()
            compr_6_: int
            for compr_6_ in (_dafny.Map({593: -876})).keys.Elements:
                d_27_v0_: int = compr_6_
                if (d_27_v0_) in (_dafny.Map({593: -876})):
                    coll6_[(d_27_v0_) - (505)] = True
            return _dafny.Map(coll6_)
        return (((D22_DC51(iife6_()
)).cf93) | (_dafny.Map({482: not(not(True))}))) | (_dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.Map({False: (_dafny.MultiSet([True])).cardinality}) for d_28_i0_ in range(default__.abs(381))]))): True}))

    @staticmethod
    def fm34(p0, globalState):
        return (_dafny.Set({True})).intersection((_dafny.Set({not(False)})) | (_dafny.Set({False, not(False), False, False, True})))

    @staticmethod
    def fm35(p0, p1, p2, p3, globalState):
        return _dafny.MultiSet([778, (len(_dafny.Map({584: _dafny.SeqWithoutIsStrInference([True])}))) - (-262), 812, 881, (222) * (-684)])

    @staticmethod
    def fm36(globalState):
        return ((_dafny.MultiSet([_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ekwnkte"))), 380, 628, -978, 812]), _dafny.MultiSet([(0) - (-503)]), _dafny.MultiSet([(D0_DC1(False, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xuckykvh"))), _dafny.CodePoint('f'))).cf3, 815])])).intersection(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([377, 315]), _dafny.MultiSet([(0) - ((0) - (len(_dafny.Map({True: -684})))), 826]), _dafny.MultiSet([150, 489])])))) - ((_dafny.MultiSet([_dafny.MultiSet([len(_dafny.Map({326: not(False)}))])])) - (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ijuq")) for d_29_i0_ in range(default__.abs(868))])), len(_dafny.Map({True: _dafny.MultiSet([True])}))])]))))

    @staticmethod
    def fm37(p0, globalState):
        return D6_DC19(543, default__.safeDivisionInt(691, 460), (95) - (len(_dafny.Map({not(True): len(_dafny.Set({(0) - (-772)}))}))))

    @staticmethod
    def fm38(globalState):
        if (True if not(True) else not(not(False))):
            return D5_DC16(_dafny.Set({not(False), False, True}))
        elif True:
            return D5_DC16(_dafny.Set({True}))

    @staticmethod
    def fm39(p0, globalState):
        return (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([True])): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_30_i0_ in range(default__.abs(767))])}))

    @staticmethod
    def fm40(p0, p1, globalState):
        def iife7_():
            coll7_ = _dafny.Map()
            compr_7_: int
            for compr_7_ in _dafny.IntegerRange(12, 391):
                d_31_v0_: int = compr_7_
                if ((12) <= (d_31_v0_)) and ((d_31_v0_) < (391)):
                    coll7_[(d_31_v0_) * (len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([True]))})))] = False
            return _dafny.Map(coll7_)
        return (_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Set({iife7_()
        , _dafny.Map({670: False})}))), -354])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({len(_dafny.Set({True})): not(not(False))}))]))]))

    @staticmethod
    def fm41(globalState):
        return ((_dafny.Map({False: True})) | (_dafny.Map({False: True}))) | ((_dafny.Map({False: not(True)})) | (_dafny.Map({not(True): True})))

    @staticmethod
    def fm42(p0, p1, p2, p3, globalState):
        return _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "emmw"))): (749) + (-156)})

    @staticmethod
    def fm43(p0, p1, globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([-599 for d_32_i1_ in range(default__.abs(344))]) for d_33_i0_ in range(default__.abs(-894))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(0) - (-910)]), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True, not(True)]))])]))) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([375, 729]) for d_34_i3_ in range(default__.abs(28))]))]) for d_35_i2_ in range(default__.abs(792))]))

    @staticmethod
    def fm44(p0, p1, p2, p3, globalState):
        return _dafny.Map({default__.safeModuloInt((0) - (len(_dafny.Map({len(_dafny.Map({True: not(False)})): False}))), len(_dafny.Set({499, 373, len(_dafny.Set({True}))}))): D9_DC25(788, len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))), len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_36_i0_ in range(default__.abs(311))])): not(False)}))])), len(_dafny.Set({False, False})), _dafny.CodePoint('u'))})

    @staticmethod
    def fm45(p0, p1, globalState):
        source1_ = D4_DC15(True, not(False), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vypeoo")))
        if source1_.is_DC13:
            d_37___mcc_h0_ = source1_.cf25
            d_38___mcc_h1_ = source1_.cf26
            d_39___mcc_h2_ = source1_.cf27
            d_40_cf27_ = d_39___mcc_h2_
            d_41_cf26_ = d_38___mcc_h1_
            d_42_cf25_ = d_37___mcc_h0_
            return D4_DC13(_dafny.SeqWithoutIsStrInference([d_40_cf27_]), d_41_cf26_, d_40_cf27_)
        elif source1_.is_DC14:
            d_43___mcc_h3_ = source1_.cf28
            d_44___mcc_h4_ = source1_.cf29
            d_45___mcc_h5_ = source1_.cf30
            d_46___mcc_h6_ = source1_.cf31
            d_47_cf31_ = d_46___mcc_h6_
            d_48_cf30_ = d_45___mcc_h5_
            d_49_cf29_ = d_44___mcc_h4_
            d_50_cf28_ = d_43___mcc_h3_
            return D4_DC13(_dafny.SeqWithoutIsStrInference([652 for d_51_i0_ in range(default__.abs(43))]), d_50_cf28_, d_48_cf30_)
        elif source1_.is_DC15:
            d_52___mcc_h7_ = source1_.cf32
            d_53___mcc_h8_ = source1_.cf33
            d_54___mcc_h9_ = source1_.cf34
            d_55_cf34_ = d_54___mcc_h9_
            d_56_cf33_ = d_53___mcc_h8_
            d_57_cf32_ = d_52___mcc_h7_
            return D4_DC13(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({d_56_cf33_: 34})), (D1_DC4(D0_DC0(_dafny.SeqWithoutIsStrInference([d_57_cf32_, d_57_cf32_, d_56_cf33_, True]), (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_56_cf33_]))).cardinality), len(_dafny.Map({d_56_cf33_: -314})), d_55_cf34_, d_56_cf33_, 222)).cf11]), d_57_cf32_, (0) - (-758))
        elif True:
            d_58___mcc_h10_ = source1_.cf24
            d_59_cf24_ = d_58___mcc_h10_
            return D4_DC13(_dafny.SeqWithoutIsStrInference([368]), False, -642)

    @staticmethod
    def fm46(globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.Map({True: not(False)})])) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({False: not(True)})]))

    @staticmethod
    def fm47(p0, p1, p2, globalState):
        return _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_60_i0_ in range(default__.abs(719))])])

    @staticmethod
    def fm48(globalState):
        return ((_dafny.MultiSet([False]) if False else _dafny.MultiSet([False, True]))) | ((_dafny.MultiSet([False])) - (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True, False, False]))))

    @staticmethod
    def fm49(p0, p1, p2, p3, globalState):
        if (D4_DC13(_dafny.SeqWithoutIsStrInference([-20, 680]), not(not(not(False))), -244)).cf26:
            return D2_DC8(not(True), True)
        elif True:
            return D2_DC8(True, True)

    @staticmethod
    def fm50(p0, globalState):
        return D3_DC11(D3_DC11(D3_DC9(_dafny.SeqWithoutIsStrInference([-215]))))

    @staticmethod
    def fm51(globalState):
        return _dafny.MultiSet([(_dafny.SeqWithoutIsStrInference([-201 for d_61_i0_ in range(default__.abs(577))])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: 491}))])), _dafny.SeqWithoutIsStrInference([986]), (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dyeiakrtw"))), -582, (0) - ((_dafny.MultiSet([755, -79, (0) - (-688)])).cardinality)])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: 23})) for d_62_i1_ in range(default__.abs(122))])), _dafny.SeqWithoutIsStrInference([(D11_DC31(len((D4_DC14(False, _dafny.Map({True: True}), 0, _dafny.Map({False: not(True)}))).cf29), (_dafny.MultiSet([False])).cardinality)).cf59 for d_63_i2_ in range(default__.abs(822))])])

    @staticmethod
    def fm52(p0, globalState):
        return D14_DC37((_dafny.Set({False, True})).issubset(_dafny.Set({False})), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k")))

    @staticmethod
    def fm53(p0, globalState):
        return (_dafny.MultiSet([D1_DC3(_dafny.MultiSet([_dafny.CodePoint('v')]))])) | (_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([D1_DC3(_dafny.MultiSet([_dafny.CodePoint('j')])) for d_64_i0_ in range(default__.abs(-950))])) + (_dafny.SeqWithoutIsStrInference([D1_DC3(_dafny.MultiSet([_dafny.CodePoint('i'), _dafny.CodePoint('k')])), D1_DC3(_dafny.MultiSet([_dafny.CodePoint('l')]))]))))

    @staticmethod
    def fm54(p0, p1, p2, globalState):
        return _dafny.Set({_dafny.CodePoint('f')})

    @staticmethod
    def fm55(globalState):
        return D4_DC15((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iptf")))) <= (-985), not((False) and (True)), (D4_DC15(True, False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uj")))).cf34)

    @staticmethod
    def fm56(p0, p1, p2, globalState):
        return D11_DC31(len((_dafny.SeqWithoutIsStrInference([804, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True, not(True)]))).cardinality])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_65_i0_ in range(default__.abs(316))])), 827]))), (518) + (449))

    @staticmethod
    def Main(noArgsParameter__):
        d_66_v0_: str
        d_66_v0_ = _dafny.CodePoint('s')
        d_67_v1_: _dafny.Seq
        d_67_v1_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f'), d_66_v0_])
        d_68_v2_: int
        d_68_v2_ = 291
        d_69_v3_: _dafny.Seq
        d_69_v3_ = _dafny.SeqWithoutIsStrInference([d_68_v2_])
        d_70_v4_: _dafny.Map
        d_70_v4_ = _dafny.Map({len(d_69_v3_): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vyifsmwe")))})
        d_71_globalState_: GlobalState
        nw0_ = GlobalState()
        nw0_.ctor__(True, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_72_i0_ in range(default__.abs(215))]), d_67_v1_, True, d_70_v4_, 583)
        d_71_globalState_ = nw0_
        d_73_v5_: bool
        d_73_v5_ = False
        d_74_v6_: D0
        d_74_v6_ = D0_DC1(d_73_v5_, default__.safeModuloInt(d_68_v2_, d_68_v2_), default__.fm0(d_68_v2_, default__.fm1(d_71_globalState_), d_73_v5_, d_71_globalState_))
        source2_ = d_74_v6_
        if source2_.is_DC0:
            d_75___mcc_h0_ = source2_.cf0
            d_76___mcc_h1_ = source2_.cf1
            d_77_cf1_ = d_76___mcc_h1_
            d_78_cf0_ = d_75___mcc_h0_
            d_79_v7_: _dafny.Map
            d_79_v7_ = _dafny.Map({default__.fm2(d_68_v2_, d_71_globalState_): d_66_v0_})
            d_79_v7_ = d_79_v7_
            d_80_v8_: _dafny.Array
            nw1_ = _dafny.Array(int(0), 7)
            d_80_v8_ = nw1_
            index0_ = default__.safeIndex(504, (d_80_v8_).length(0))
            (d_80_v8_)[index0_] = -792
            index1_ = default__.safeIndex(504, (d_80_v8_).length(0))
            (d_80_v8_)[index1_] = (d_77_cf1_) - (default__.safeModuloInt(d_77_cf1_, len(d_78_cf0_)))
            index2_ = default__.safeIndex(504, (d_80_v8_).length(0))
            (d_80_v8_)[index2_] = (d_80_v8_)[default__.safeIndex(504, (d_80_v8_).length(0))]
            d_81_v9_: D0
            d_81_v9_ = D0_DC2(D0_DC1(d_73_v5_, d_68_v2_, d_66_v0_))
            d_82_v10_: _dafny.Seq
            d_82_v10_ = _dafny.SeqWithoutIsStrInference([d_67_v1_, d_67_v1_, d_67_v1_, d_67_v1_, d_67_v1_])
            d_83_v11_: D0
            d_83_v11_ = D0_DC0(default__.fm4(d_68_v2_, len(d_82_v10_), (d_80_v8_)[default__.safeIndex(504, (d_80_v8_).length(0))], d_71_globalState_), d_77_cf1_)
            d_84_v12_: D0
            d_84_v12_ = D0_DC2(d_83_v11_)
            d_85_v13_: D0
            d_85_v13_ = D0_DC2(d_83_v11_)
            pat_let_tv0_ = d_83_v11_
            def iife8_(_pat_let0_0):
                def iife9_(d_86_dt__update__tmp_h0_):
                    def iife10_(_pat_let1_0):
                        def iife11_(d_87_dt__update_hcf5_h0_):
                            return D0_DC2(d_87_dt__update_hcf5_h0_)
                        return iife11_(_pat_let1_0)
                    return iife10_(pat_let_tv0_)
                return iife9_(_pat_let0_0)
            source3_ = (d_81_v9_ if default__.fm3(d_71_globalState_) else iife8_(d_81_v9_))
            if source3_.is_DC0:
                d_88___mcc_h6_ = source3_.cf0
                d_89___mcc_h7_ = source3_.cf1
                d_90_cf1_ = d_89___mcc_h7_
                d_91_cf0_ = d_88___mcc_h6_
                d_92_v14_: C1
                nw2_ = C1()
                nw2_.ctor__(d_68_v2_)
                d_92_v14_ = nw2_
                d_73_v5_ = (d_91_cf0_)[default__.safeIndex((0) - (d_90_cf1_), len(d_91_cf0_))]
                d_90_cf1_ = (0) - ((d_80_v8_)[default__.safeIndex(504, (d_80_v8_).length(0))])
                d_93_v15_: T0
                nw3_ = C5()
                nw3_.ctor__(d_90_cf1_, d_90_cf1_)
                d_93_v15_ = nw3_
                d_94_v16_: _dafny.MultiSet
                d_94_v16_ = _dafny.MultiSet([d_93_v15_])
                d_95_v17_: _dafny.Map
                d_95_v17_ = _dafny.Map({d_68_v2_: d_73_v5_})
                d_96_v18_: _dafny.Map
                d_96_v18_ = _dafny.Map({((d_95_v17_)[d_90_cf1_] if (d_90_cf1_) in (d_95_v17_) else not(d_73_v5_)): (d_80_v8_)[default__.safeIndex(504, (d_80_v8_).length(0))]})
                d_97_v19_: _dafny.Array
                def lambda0_(d_98_v5_):
                    def lambda1_(d_99_i1_):
                        return d_98_v5_

                    return lambda1_

                init0_ = lambda0_(d_73_v5_)
                nw4_ = _dafny.Array(None, 26)
                for i0_0_ in range(nw4_.length(0)):
                    nw4_[i0_0_] = init0_(i0_0_)
                d_97_v19_ = nw4_
                d_100_v20_: _dafny.Map
                d_100_v20_ = _dafny.Map({547: d_67_v1_})
                d_101_v21_: _dafny.MultiSet
                d_101_v21_ = _dafny.MultiSet([d_68_v2_])
                d_102_v22_: _dafny.MultiSet
                d_102_v22_ = _dafny.MultiSet([d_101_v21_])
                d_103_v23_: T1
                nw5_ = C4()
                nw5_.ctor__(((d_94_v16_)[d_93_v15_] if (d_93_v15_) in (d_94_v16_) else d_68_v2_), len((d_96_v18_).set(d_73_v5_, len(d_67_v1_))), d_97_v19_, d_100_v20_, d_68_v2_, d_102_v22_)
                d_103_v23_ = nw5_
                d_104_v24_: D6
                d_104_v24_ = D6_DC18(d_103_v23_)
                d_105_v25_: _dafny.Map
                d_105_v25_ = _dafny.Map({d_80_v8_: d_69_v3_})
                d_106_v26_: _dafny.Map
                d_106_v26_ = _dafny.Map({d_104_v24_: d_105_v25_})
                (d_71_globalState_).f5 = len((((d_106_v26_)[d_104_v24_] if (d_104_v24_) in (d_106_v26_) else d_105_v25_)).set(d_80_v8_, d_69_v3_))
            elif source3_.is_DC1:
                d_107___mcc_h8_ = source3_.cf2
                d_108___mcc_h9_ = source3_.cf3
                d_109___mcc_h10_ = source3_.cf4
                d_110_cf4_ = d_109___mcc_h10_
                d_111_cf3_ = d_108___mcc_h9_
                d_112_cf2_ = d_107___mcc_h8_
                d_113_v27_: _dafny.Array
                nw6_ = _dafny.Array(_dafny.Array(None, 0), 17)
                d_113_v27_ = nw6_
                nw7_ = _dafny.Array(_dafny.Array(None, 0), 10)
                d_113_v27_ = nw7_
                d_114_v28_: _dafny.Map
                d_114_v28_ = _dafny.Map({(0) - (d_111_cf3_): d_112_cf2_})
                d_114_v28_ = (_dafny.Map({(d_80_v8_)[default__.safeIndex(504, (d_80_v8_).length(0))]: not(d_112_cf2_)})) | (_dafny.Map({d_77_cf1_: d_112_cf2_}))
                d_115_v29_: D6
                d_115_v29_ = D6_DC19(d_68_v2_, (d_80_v8_)[default__.safeIndex(504, (d_80_v8_).length(0))], 121)
                pat_let_tv1_ = d_80_v8_
                pat_let_tv2_ = d_80_v8_
                index3_ = default__.safeIndex(504, (d_80_v8_).length(0))
                def iife12_(_pat_let2_0):
                    def iife13_(d_116_dt__update__tmp_h1_):
                        def iife14_(_pat_let3_0):
                            def iife15_(d_117_dt__update_hcf41_h0_):
                                return D6_DC19((d_116_dt__update__tmp_h1_).cf39, (d_116_dt__update__tmp_h1_).cf40, d_117_dt__update_hcf41_h0_)
                            return iife15_(_pat_let3_0)
                        return iife14_((pat_let_tv2_)[default__.safeIndex(504, (pat_let_tv1_).length(0))])
                    return iife13_(_pat_let2_0)
                (d_80_v8_)[index3_] = default__.fm2((iife12_(d_115_v29_)).cf41, d_71_globalState_)
                d_118_v30_: C6
                nw8_ = C6()
                nw8_.ctor__(d_68_v2_, d_73_v5_)
                d_118_v30_ = nw8_
                d_119_v31_: _dafny.Set
                d_119_v31_ = _dafny.Set({d_118_v30_, d_118_v30_, d_118_v30_, d_118_v30_, d_118_v30_})
                d_120_v32_: _dafny.Map
                d_120_v32_ = _dafny.Map({d_119_v31_: d_78_cf0_})
                d_121_v33_: _dafny.Seq
                d_121_v33_ = _dafny.SeqWithoutIsStrInference([d_119_v31_, _dafny.Set({d_118_v30_, d_118_v30_, d_118_v30_})])
                d_120_v32_ = (d_120_v32_).set((d_121_v33_)[default__.safeIndex(d_68_v2_, len(d_121_v33_))], d_78_cf0_)
            elif True:
                d_122___mcc_h11_ = source3_.cf5
                d_123_cf5_ = d_122___mcc_h11_
                d_124_v34_: _dafny.Array
                nw9_ = _dafny.Array(False, 27)
                d_124_v34_ = nw9_
                index4_ = default__.safeIndex(866, (d_124_v34_).length(0))
                (d_124_v34_)[index4_] = d_73_v5_
                d_125_v35_: _dafny.Map
                d_125_v35_ = _dafny.Map({d_73_v5_: d_73_v5_})
                index5_ = default__.safeIndex(866, (d_124_v34_).length(0))
                (d_124_v34_)[index5_] = ((d_125_v35_)[d_73_v5_] if (d_73_v5_) in (d_125_v35_) else d_73_v5_)
                d_126_v36_: C9
                nw10_ = C9()
                nw10_.ctor__((d_80_v8_)[default__.safeIndex(504, (d_80_v8_).length(0))])
                d_126_v36_ = nw10_
                d_127_v37_: _dafny.MultiSet
                d_127_v37_ = _dafny.MultiSet([(d_124_v34_)[default__.safeIndex(866, (d_124_v34_).length(0))]])
                d_128_v38_: _dafny.Map
                d_128_v38_ = _dafny.Map({d_73_v5_: (0) - ((d_127_v37_).cardinality)})
                d_68_v2_ = (d_68_v2_) * (default__.safeModuloInt(d_68_v2_, ((d_128_v38_)[True] if (True) in (d_128_v38_) else 858)))
                d_129_v39_: _dafny.Array
                nw11_ = _dafny.Array(_dafny.Array(None, 0), 25)
                d_129_v39_ = nw11_
                d_130_v40_: _dafny.Array
                nw12_ = _dafny.Array(None, 6)
                nw12_[int(0)] = d_129_v39_
                nw12_[int(1)] = d_129_v39_
                nw12_[int(2)] = d_129_v39_
                nw12_[int(3)] = d_129_v39_
                nw12_[int(4)] = (d_129_v39_ if d_73_v5_ else d_129_v39_)
                nw12_[int(5)] = d_129_v39_
                d_130_v40_ = nw12_
                index6_ = default__.safeIndex(369, (d_130_v40_).length(0))
                (d_130_v40_)[index6_] = d_129_v39_
                d_131_v41_: C10
                nw13_ = C10()
                nw13_.ctor__(d_77_cf1_)
                d_131_v41_ = nw13_
                d_132_v42_: _dafny.Map
                d_132_v42_ = _dafny.Map({d_131_v41_: 692})
                d_133_v43_: _dafny.Seq
                d_133_v43_ = _dafny.SeqWithoutIsStrInference([d_125_v35_])
                index7_ = default__.safeIndex(369, (d_130_v40_).length(0))
                rhs0_ = d_129_v39_
                rhs1_ = _dafny.MultiSet((_dafny.SeqWithoutIsStrInference([False])) + (((d_78_cf0_).set(default__.safeIndex((d_127_v37_).cardinality, len(d_78_cf0_)), (d_124_v34_)[default__.safeIndex(866, (d_124_v34_).length(0))])) + (default__.fm4((d_80_v8_)[default__.safeIndex(504, (d_80_v8_).length(0))], d_68_v2_, len(d_132_v42_), d_71_globalState_))))
                rhs2_ = ((d_133_v43_) + (default__.fm46(d_71_globalState_))) == (d_133_v43_)
                rhs3_ = ((d_80_v8_)[default__.safeIndex(504, (d_80_v8_).length(0))]) + ((len(d_67_v1_)) + (d_68_v2_))
                lhs0_ = d_130_v40_
                lhs1_ = default__.safeIndex(369, (d_130_v40_).length(0))
                lhs2_ = d_71_globalState_
                lhs3_ = d_71_globalState_
                lhs0_[lhs1_] = rhs0_
                d_127_v37_ = rhs1_
                lhs2_.f0 = rhs2_
                lhs3_.f5 = rhs3_
        elif source2_.is_DC1:
            d_134___mcc_h2_ = source2_.cf2
            d_135___mcc_h3_ = source2_.cf3
            d_136___mcc_h4_ = source2_.cf4
            d_137_cf4_ = d_136___mcc_h4_
            d_138_cf3_ = d_135___mcc_h3_
            d_139_cf2_ = d_134___mcc_h2_
            d_73_v5_ = d_73_v5_
            d_140_v44_: _dafny.Array
            nw14_ = _dafny.Array(None, 1)
            d_140_v44_ = nw14_
            d_141_v45_: D20
            d_141_v45_ = D20_DC46(d_140_v44_)
            d_142_v46_: _dafny.Map
            d_142_v46_ = _dafny.Map({(d_141_v45_).cf84: d_138_cf3_})
            d_142_v46_ = (d_142_v46_).set(d_140_v44_, d_68_v2_)
            (d_71_globalState_).f0 = d_139_cf2_
            (d_71_globalState_).f5 = d_138_cf3_
        elif True:
            d_143___mcc_h5_ = source2_.cf5
            d_144_cf5_ = d_143___mcc_h5_
            d_145_v47_: _dafny.Array
            nw15_ = _dafny.Array(None, 14)
            nw15_[int(0)] = 563
            nw15_[int(1)] = d_68_v2_
            nw15_[int(2)] = d_68_v2_
            nw15_[int(3)] = d_68_v2_
            nw15_[int(4)] = d_68_v2_
            nw15_[int(5)] = len(d_70_v4_)
            nw15_[int(6)] = -185
            nw15_[int(7)] = d_68_v2_
            nw15_[int(8)] = d_68_v2_
            nw15_[int(9)] = d_68_v2_
            nw15_[int(10)] = (0) - (d_68_v2_)
            nw15_[int(11)] = -959
            nw15_[int(12)] = len(d_69_v3_)
            nw15_[int(13)] = len(d_67_v1_)
            d_145_v47_ = nw15_
            d_146_v48_: _dafny.Array
            def lambda2_(d_147_i2_):
                return True

            init1_ = lambda2_
            nw16_ = _dafny.Array(None, 5)
            for i0_1_ in range(nw16_.length(0)):
                nw16_[i0_1_] = init1_(i0_1_)
            d_146_v48_ = nw16_
            d_148_v49_: _dafny.Seq
            d_148_v49_ = _dafny.SeqWithoutIsStrInference([False])
            d_149_v50_: _dafny.Map
            d_149_v50_ = _dafny.Map({len(d_148_v49_): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vuxo"))})
            d_150_v51_: C7
            nw17_ = C7()
            nw17_.ctor__(d_145_v47_, d_145_v47_, default__.fm2(d_68_v2_, d_71_globalState_), d_146_v48_, d_149_v50_)
            d_150_v51_ = nw17_
            rhs4_ = d_68_v2_
            rhs5_ = (d_68_v2_) - (d_68_v2_)
            lhs4_ = d_71_globalState_
            lhs5_ = d_71_globalState_
            lhs4_.f5 = rhs4_
            lhs5_.f5 = rhs5_
            d_151_v52_: C0
            nw18_ = C0()
            nw18_.ctor__(d_146_v48_, ((d_150_v51_).fm14(d_71_globalState_)) >= (len(d_69_v3_)))
            d_151_v52_ = nw18_
            d_151_v52_ = d_151_v52_
            d_152_v53_: C8
            nw19_ = C8()
            nw19_.ctor__()
            d_152_v53_ = nw19_
        (d_71_globalState_).f5 = default__.safeDivisionInt(d_68_v2_, d_68_v2_)
        d_153_v54_: _dafny.Map
        d_153_v54_ = _dafny.Map({d_73_v5_: default__.fm2(d_68_v2_, d_71_globalState_)})
        d_154_v55_: _dafny.Map
        d_154_v55_ = _dafny.Map({d_153_v54_: _dafny.MultiSet([not(d_73_v5_)])})
        d_154_v55_ = (d_154_v55_) | (d_154_v55_)
        d_73_v5_ = (d_73_v5_ if d_73_v5_ else d_73_v5_)
        d_155_v56_: _dafny.Seq
        d_155_v56_ = _dafny.SeqWithoutIsStrInference([d_73_v5_, d_73_v5_])
        d_156_v57_: _dafny.Set
        d_156_v57_ = _dafny.Set({(0) - (d_68_v2_), default__.fm2(d_68_v2_, d_71_globalState_)})
        d_157_v58_: _dafny.Set
        d_157_v58_ = _dafny.Set({d_73_v5_, d_73_v5_})
        d_158_v59_: _dafny.Array
        nw20_ = _dafny.Array(False, 7)
        d_158_v59_ = nw20_
        d_159_v60_: _dafny.Map
        d_159_v60_ = _dafny.Map({d_68_v2_: d_67_v1_})
        d_160_v61_: T1
        nw21_ = C3()
        nw21_.ctor__(_dafny.Map({d_68_v2_: (0) - (d_68_v2_)}), d_158_v59_, d_159_v60_, d_68_v2_)
        d_160_v61_ = nw21_
        d_161_v62_: _dafny.Map
        d_161_v62_ = _dafny.Map({d_69_v3_: d_160_v61_})
        d_162_v63_: _dafny.Set
        d_162_v63_ = _dafny.Set({d_66_v0_})
        d_163_v64_: _dafny.MultiSet
        d_163_v64_ = _dafny.MultiSet([d_157_v58_])
        d_164_v65_: _dafny.Array
        nw22_ = _dafny.Array(None, 25)
        nw22_[int(0)] = default__.fm2(d_68_v2_, d_71_globalState_)
        nw22_[int(1)] = len(d_155_v56_)
        nw22_[int(2)] = d_68_v2_
        nw22_[int(3)] = d_68_v2_
        nw22_[int(4)] = d_68_v2_
        nw22_[int(5)] = d_68_v2_
        nw22_[int(6)] = d_68_v2_
        nw22_[int(7)] = d_68_v2_
        nw22_[int(8)] = d_68_v2_
        nw22_[int(9)] = -22
        nw22_[int(10)] = d_68_v2_
        nw22_[int(11)] = d_68_v2_
        nw22_[int(12)] = len(d_156_v57_)
        nw22_[int(13)] = len(d_157_v58_)
        nw22_[int(14)] = -993
        nw22_[int(15)] = len(d_161_v62_)
        nw22_[int(16)] = len(d_70_v4_)
        nw22_[int(17)] = d_68_v2_
        nw22_[int(18)] = d_68_v2_
        nw22_[int(19)] = 253
        nw22_[int(20)] = (d_160_v61_).f6
        nw22_[int(21)] = len(d_162_v63_)
        nw22_[int(22)] = (d_160_v61_).f6
        nw22_[int(23)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ibypqjp")))
        nw22_[int(24)] = ((d_163_v64_)[d_157_v58_] if (d_157_v58_) in (d_163_v64_) else (d_160_v61_).f6)
        d_164_v65_ = nw22_
        d_165_v66_: _dafny.Map
        d_165_v66_ = _dafny.Map({d_164_v65_: not(d_73_v5_)})
        d_166_v67_: D10
        d_166_v67_ = D10_DC28(d_165_v66_)
        d_166_v67_ = d_166_v67_
        d_167_v68_: _dafny.Map
        d_167_v68_ = _dafny.Map({d_73_v5_: d_160_v61_.f9})
        d_168_v69_: _dafny.Seq
        d_169_v70_: bool
        out0_: _dafny.Seq
        out1_: bool
        out0_, out1_ = (d_160_v61_).m6(d_73_v5_, (_dafny.SeqWithoutIsStrInference([d_73_v5_, d_73_v5_])) + (default__.fm4((d_160_v61_).f6, d_68_v2_, len(d_167_v68_), d_71_globalState_)), d_68_v2_, d_71_globalState_)
        d_168_v69_ = out0_
        d_169_v70_ = out1_
        (d_71_globalState_).f2 = ((d_67_v1_) + (d_67_v1_)) + (d_67_v1_)
        d_170_v71_: _dafny.Seq
        d_171_v72_: bool
        out2_: _dafny.Seq
        out3_: bool
        out2_, out3_ = (d_160_v61_).m6(d_73_v5_, default__.fm4((d_160_v61_).f6, d_68_v2_, (d_160_v61_).f6, d_71_globalState_), len(_dafny.SeqWithoutIsStrInference([d_73_v5_, d_169_v70_, d_73_v5_, d_73_v5_])), d_71_globalState_)
        d_170_v71_ = out2_
        d_171_v72_ = out3_
        d_172_i3_: int
        d_172_i3_ = 0
        with _dafny.label("0"):
            while not(d_171_v72_):
                with _dafny.c_label("0"):
                    if (d_172_i3_) >= (100):
                        raise _dafny.Break("0")
                    d_172_i3_ = (d_172_i3_) + (1)
                    d_173_v73_: C6
                    nw23_ = C6()
                    nw23_.ctor__((d_160_v61_).f6, d_169_v70_)
                    d_173_v73_ = nw23_
                    d_173_v73_ = d_173_v73_
                    d_174_v74_: _dafny.Seq
                    d_175_v75_: bool
                    out4_: _dafny.Seq
                    out5_: bool
                    out4_, out5_ = (d_160_v61_).m6(((0) - ((d_160_v61_).f6)) <= ((d_160_v61_).f6), (d_155_v56_) + (d_155_v56_), default__.safeDivisionInt(d_68_v2_, (0) - ((d_160_v61_).f6)), d_71_globalState_)
                    d_174_v74_ = out4_
                    d_175_v75_ = out5_
                    d_176_v76_: _dafny.MultiSet
                    d_176_v76_ = _dafny.MultiSet([(d_156_v57_).issubset(d_156_v57_), True, d_73_v5_, (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_177_i4_ in range(default__.abs(311))])) < (d_67_v1_), d_171_v72_])
                    d_176_v76_ = (d_176_v76_).intersection(d_176_v76_)
                    d_69_v3_ = default__.fm40(default__.safeDivisionInt((_dafny.MultiSet(d_155_v56_)).cardinality, (d_160_v61_).f6), d_157_v58_, d_71_globalState_)
                    pass
            pass
        d_73_v5_ = default__.fm3(d_71_globalState_)
        hi0_ = default__.safeModuloInt((d_160_v61_).f6, len(d_156_v57_))
        for d_178_i5_ in range(d_68_v2_, hi0_):
            d_179_v77_: _dafny.Seq
            d_180_v78_: bool
            out6_: _dafny.Seq
            out7_: bool
            out6_, out7_ = (d_160_v61_).m6(d_171_v72_, (_dafny.SeqWithoutIsStrInference([d_169_v70_])) + (d_155_v56_), 481, d_71_globalState_)
            d_179_v77_ = out6_
            d_180_v78_ = out7_
            if not(d_169_v70_):
                d_181_v79_: _dafny.Map
                d_181_v79_ = _dafny.Map({False: d_180_v78_})
                d_182_v80_: _dafny.Set
                d_182_v80_ = _dafny.Set({d_181_v79_})
                d_183_v81_: _dafny.Array
                nw24_ = _dafny.Array(None, 8)
                nw24_[int(0)] = _dafny.SeqWithoutIsStrInference([d_66_v0_ for d_184_i6_ in range(default__.abs(175))])
                nw24_[int(1)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kosnitju"))
                nw24_[int(2)] = default__.fm17((d_160_v61_).f6, len(d_67_v1_), d_71_globalState_)
                nw24_[int(3)] = d_67_v1_
                nw24_[int(4)] = d_67_v1_
                nw24_[int(5)] = d_67_v1_
                nw24_[int(6)] = d_67_v1_
                nw24_[int(7)] = d_67_v1_
                d_183_v81_ = nw24_
                d_185_v82_: D10
                d_185_v82_ = D10_DC29((d_160_v61_).f6, d_183_v81_, d_180_v78_, d_178_i5_)
                rhs6_ = d_164_v65_
                rhs7_ = d_182_v80_
                rhs8_ = ((d_181_v79_)[d_169_v70_] if (d_169_v70_) in (d_181_v79_) else (d_185_v82_).cf55)
                lhs6_ = d_71_globalState_
                d_164_v65_ = rhs6_
                d_182_v80_ = rhs7_
                lhs6_.f0 = rhs8_
                d_186_v83_: _dafny.Array
                d_186_v83_ = d_160_v61_.f9
                d_187_v84_: C7
                nw25_ = C7()
                nw25_.ctor__(d_164_v65_, d_164_v65_, default__.safeDivisionInt((d_160_v61_).f6, len((d_168_v69_)[default__.safeIndex((d_160_v61_).f6, len(d_168_v69_))])), (d_186_v83_), _dafny.Map({d_68_v2_: d_67_v1_}))
                d_187_v84_ = nw25_
                d_188_v85_: C1
                nw26_ = C1()
                nw26_.ctor__((684) - ((0) - (len(d_67_v1_))))
                d_188_v85_ = nw26_
                d_188_v85_ = d_188_v85_
                d_189_v86_: D11
                d_189_v86_ = D11_DC31((d_160_v61_).f6, (0) - ((d_160_v61_).f6))
                d_189_v86_ = default__.fm56(_dafny.SeqWithoutIsStrInference([d_66_v0_ for d_190_i7_ in range(default__.abs(-459))]), d_73_v5_, d_67_v1_, d_71_globalState_)
                d_191_v87_: T0
                nw27_ = C1()
                nw27_.ctor__(len((d_69_v3_).set(default__.safeIndex((d_160_v61_).f6, len(d_69_v3_)), d_178_i5_)))
                d_191_v87_ = nw27_
                d_192_v88_: _dafny.Seq
                d_192_v88_ = _dafny.SeqWithoutIsStrInference([d_191_v87_, d_191_v87_, d_191_v87_, d_191_v87_])
                d_193_v89_: _dafny.Array
                nw28_ = _dafny.Array(None, 22)
                nw28_[int(0)] = d_191_v87_
                nw28_[int(1)] = d_191_v87_
                nw28_[int(2)] = d_191_v87_
                nw28_[int(3)] = d_191_v87_
                nw28_[int(4)] = d_191_v87_
                nw28_[int(5)] = d_191_v87_
                nw28_[int(6)] = d_191_v87_
                nw28_[int(7)] = d_191_v87_
                nw28_[int(8)] = d_191_v87_
                nw28_[int(9)] = d_191_v87_
                nw28_[int(10)] = d_191_v87_
                nw28_[int(11)] = d_191_v87_
                nw28_[int(12)] = d_191_v87_
                nw28_[int(13)] = d_191_v87_
                nw28_[int(14)] = (d_192_v88_)[default__.safeIndex(default__.fm2((d_191_v87_).f6, d_71_globalState_), len(d_192_v88_))]
                nw28_[int(15)] = d_191_v87_
                nw28_[int(16)] = d_191_v87_
                nw28_[int(17)] = d_191_v87_
                nw28_[int(18)] = d_191_v87_
                nw28_[int(19)] = d_191_v87_
                nw28_[int(20)] = d_191_v87_
                nw28_[int(21)] = d_191_v87_
                d_193_v89_ = nw28_
                index8_ = default__.safeIndex(295, (d_193_v89_).length(0))
                (d_193_v89_)[index8_] = d_191_v87_
                d_194_v90_: D2
                d_194_v90_ = D2_DC6((d_160_v61_).f6, d_187_v84_.f12)
                d_195_v91_: _dafny.MultiSet
                d_195_v91_ = _dafny.MultiSet([(d_194_v90_).cf14, d_187_v84_.f12])
                d_196_v92_: _dafny.Map
                d_196_v92_ = _dafny.Map({len((d_70_v4_).set((d_160_v61_).f6, (d_160_v61_).f6)): not(d_73_v5_)})
                index9_ = default__.safeIndex(295, (d_193_v89_).length(0))
                rhs9_ = d_191_v87_
                rhs10_ = ((len(d_196_v92_)) + (default__.fm2((d_160_v61_).f6, d_71_globalState_))) <= ((d_160_v61_).f6)
                rhs11_ = d_155_v56_
                rhs12_ = (d_67_v1_)[default__.safeIndex(d_68_v2_, len(d_67_v1_))]
                rhs13_ = d_195_v91_
                lhs7_ = d_193_v89_
                lhs8_ = default__.safeIndex(295, (d_193_v89_).length(0))
                lhs7_[lhs8_] = rhs9_
                d_171_v72_ = rhs10_
                d_155_v56_ = rhs11_
                d_66_v0_ = rhs12_
                d_195_v91_ = rhs13_
            elif True:
                (d_160_v61_).f9 = d_160_v61_.f9
                d_197_v93_: D4
                d_197_v93_ = D4_DC13(d_69_v3_, not(d_169_v70_), (d_160_v61_).f6)
                d_198_v94_: D14
                d_198_v94_ = D14_DC37(d_73_v5_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_199_i9_ in range(default__.abs(743))]))
                d_197_v93_ = D4_DC13((_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([d_169_v70_])).cardinality for d_200_i8_ in range(default__.abs(569))]) if d_169_v70_ else d_69_v3_), (d_198_v94_).cf70, ((0) - ((d_160_v61_).f6)) * ((d_160_v61_).f6))
                d_201_v95_: _dafny.Array
                nw29_ = _dafny.Array(D2.default()(), 21)
                d_201_v95_ = nw29_
                d_202_v96_: _dafny.Map
                d_202_v96_ = _dafny.Map({d_178_i5_: d_180_v78_})
                d_203_v97_: _dafny.Map
                d_203_v97_ = _dafny.Map({d_202_v96_: d_169_v70_})
                d_204_v98_: D2
                d_204_v98_ = D2_DC6(len(d_203_v97_), d_164_v65_)
                index10_ = default__.safeIndex(683, (d_201_v95_).length(0))
                (d_201_v95_)[index10_] = d_204_v98_
                index11_ = default__.safeIndex(683, (d_201_v95_).length(0))
                (d_201_v95_)[index11_] = D2_DC6(d_178_i5_, d_164_v65_)
                d_155_v56_ = (d_155_v56_) + (d_155_v56_)
                d_66_v0_ = d_66_v0_
            d_205_v99_: _dafny.Array
            def lambda3_(d_206_v0_):
                def lambda4_(d_207_i10_):
                    return d_206_v0_

                return lambda4_

            init2_ = lambda3_(d_66_v0_)
            nw30_ = _dafny.Array(None, 24)
            for i0_2_ in range(nw30_.length(0)):
                nw30_[i0_2_] = init2_(i0_2_)
            d_205_v99_ = nw30_
            index12_ = default__.safeIndex(630, (d_205_v99_).length(0))
            (d_205_v99_)[index12_] = d_66_v0_
            index13_ = default__.safeIndex(630, (d_205_v99_).length(0))
            (d_205_v99_)[index13_] = d_66_v0_
            index14_ = default__.safeIndex(970, (d_164_v65_).length(0))
            (d_164_v65_)[index14_] = (d_160_v61_).f6
            d_208_v100_: _dafny.Array
            nw31_ = _dafny.Array(D1.default()(), 11)
            d_208_v100_ = nw31_
            d_209_v101_: D1
            def iife16_(_pat_let4_0):
                def iife17_(d_210_dt__update__tmp_h2_):
                    def iife18_(_pat_let5_0):
                        def iife19_(d_211_dt__update_hcf1_h0_):
                            return D0_DC0((d_210_dt__update__tmp_h2_).cf0, d_211_dt__update_hcf1_h0_)
                        return iife19_(_pat_let5_0)
                    return iife18_(-120)
                return iife17_(_pat_let4_0)
            d_209_v101_ = D1_DC4(iife16_(D0_DC0(d_155_v56_, d_178_i5_)), len(_dafny.Map({d_171_v72_: d_171_v72_})), d_67_v1_, d_169_v70_, d_68_v2_)
            index15_ = default__.safeIndex(818, (d_208_v100_).length(0))
            (d_208_v100_)[index15_] = d_209_v101_
            index16_ = default__.safeIndex(570, (d_164_v65_).length(0))
            (d_164_v65_)[index16_] = d_178_i5_
            d_212_v102_: _dafny.Map
            d_212_v102_ = _dafny.Map({(d_205_v99_)[default__.safeIndex(630, (d_205_v99_).length(0))]: d_73_v5_})
            index17_ = default__.safeIndex(970, (d_164_v65_).length(0))
            index18_ = default__.safeIndex(818, (d_208_v100_).length(0))
            index19_ = default__.safeIndex(570, (d_164_v65_).length(0))
            rhs14_ = default__.fm2(len((_dafny.SeqWithoutIsStrInference([d_66_v0_ for d_213_i11_ in range(default__.abs(101))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r")))), d_71_globalState_)
            rhs15_ = default__.fm2(d_178_i5_, d_71_globalState_)
            rhs16_ = d_209_v101_
            rhs17_ = default__.safeDivisionInt(len(d_212_v102_), d_68_v2_)
            rhs18_ = d_178_i5_
            lhs9_ = d_164_v65_
            lhs10_ = default__.safeIndex(970, (d_164_v65_).length(0))
            lhs11_ = d_208_v100_
            lhs12_ = default__.safeIndex(818, (d_208_v100_).length(0))
            lhs13_ = d_164_v65_
            lhs14_ = default__.safeIndex(570, (d_164_v65_).length(0))
            lhs9_[lhs10_] = rhs14_
            d_68_v2_ = rhs15_
            lhs11_[lhs12_] = rhs16_
            lhs13_[lhs14_] = rhs17_
            d_68_v2_ = rhs18_
        index20_ = default__.safeIndex(157, (d_164_v65_).length(0))
        (d_164_v65_)[index20_] = default__.safeDivisionInt(d_68_v2_, len(_dafny.Set({d_171_v72_})))
        index21_ = default__.safeIndex(157, (d_164_v65_).length(0))
        (d_164_v65_)[index21_] = (d_160_v61_).f6
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_164_v65_).length(0)):
            d_214_i12_: int = guard_loop_0_
            if (True) and (((0) <= (d_214_i12_)) and ((d_214_i12_) < ((d_164_v65_).length(0)))):
                (d_164_v65_)[(d_214_i12_)] = default__.safeModuloInt(d_214_i12_, len(d_153_v54_))
        arr0_ = d_160_v61_.f9
        index22_ = default__.safeIndex(541, (d_160_v61_.f9).length(0))
        arr0_[index22_] = not (d_171_v72_) or (d_171_v72_)
        arr1_ = d_160_v61_.f9
        index23_ = default__.safeIndex(541, (d_160_v61_.f9).length(0))
        rhs19_ = ((d_164_v65_)[default__.safeIndex(157, (d_164_v65_).length(0))]) > (309)
        rhs20_ = d_73_v5_
        lhs15_ = d_160_v61_.f9
        lhs16_ = default__.safeIndex(541, (d_160_v61_.f9).length(0))
        d_171_v72_ = rhs19_
        lhs15_[lhs16_] = rhs20_
        d_215_v103_: _dafny.MultiSet
        d_215_v103_ = _dafny.MultiSet([d_68_v2_, (d_160_v61_).f6])
        d_216_v104_: str
        d_217_v105_: int
        d_218_v106_: _dafny.Array
        d_219_v107_: int
        out8_: str
        out9_: int
        out10_: _dafny.Array
        out11_: int
        out8_, out9_, out10_, out11_ = (d_160_v61_).m7((d_215_v103_) | (default__.fm19(534, d_71_globalState_)), len((_dafny.SeqWithoutIsStrInference([d_169_v70_, d_169_v70_])) + (d_155_v56_)), d_71_globalState_)
        d_216_v104_ = out8_
        d_217_v105_ = out9_
        d_218_v106_ = out10_
        d_219_v107_ = out11_
        d_220_i13_: int
        d_220_i13_ = 0
        with _dafny.label("1"):
            while default__.fm3(d_71_globalState_):
                with _dafny.c_label("1"):
                    if (d_220_i13_) >= (100):
                        raise _dafny.Break("1")
                    d_220_i13_ = (d_220_i13_) + (1)
                    d_221_v108_: _dafny.Array
                    nw32_ = _dafny.Array(None, 17)
                    d_221_v108_ = nw32_
                    d_222_v109_: C0
                    nw33_ = C0()
                    nw33_.ctor__(d_160_v61_.f9, d_73_v5_)
                    d_222_v109_ = nw33_
                    index24_ = default__.safeIndex(968, (d_221_v108_).length(0))
                    (d_221_v108_)[index24_] = d_222_v109_
                    index25_ = default__.safeIndex(968, (d_221_v108_).length(0))
                    (d_221_v108_)[index25_] = (d_222_v109_ if d_169_v70_ else d_222_v109_)
                    d_68_v2_ = (0) - ((d_160_v61_).f6)
                    d_169_v70_ = d_73_v5_
                    if d_171_v72_:
                        d_223_v110_: str
                        d_224_v111_: int
                        d_225_v112_: _dafny.Array
                        d_226_v113_: int
                        out12_: str
                        out13_: int
                        out14_: _dafny.Array
                        out15_: int
                        out12_, out13_, out14_, out15_ = (d_160_v61_).m7(d_215_v103_, ((d_164_v65_)[default__.safeIndex(157, (d_164_v65_).length(0))] if d_171_v72_ else (d_164_v65_)[default__.safeIndex(157, (d_164_v65_).length(0))]), d_71_globalState_)
                        d_223_v110_ = out12_
                        d_224_v111_ = out13_
                        d_225_v112_ = out14_
                        d_226_v113_ = out15_
                        d_227_v114_: _dafny.Seq
                        d_227_v114_ = _dafny.SeqWithoutIsStrInference([d_164_v65_, d_164_v65_])
                        d_227_v114_ = d_227_v114_
                        d_228_v115_: _dafny.MultiSet
                        d_228_v115_ = _dafny.MultiSet([False, (d_155_v56_)[default__.safeIndex(d_226_v113_, len(d_155_v56_))]])
                        d_229_v116_: _dafny.Map
                        d_229_v116_ = _dafny.Map({d_73_v5_: d_222_v109_.f20})
                        d_219_v107_ = (((d_228_v115_).cardinality) * (d_224_v111_)) - ((len(d_67_v1_)) + (len(d_229_v116_)))
                        d_230_v117_: _dafny.Map
                        d_230_v117_ = _dafny.Map({d_74_v6_: True})
                        d_230_v117_ = (d_230_v117_).set(d_74_v6_, (d_160_v61_.f9)[default__.safeIndex(541, (d_160_v61_.f9).length(0))])
                        d_216_v104_ = default__.fm0(default__.fm2((d_160_v61_).f6, d_71_globalState_), (default__.fm1(d_71_globalState_)) | (d_156_v57_), (d_67_v1_) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eiyyqrevp"))), d_71_globalState_)
                    elif True:
                        d_169_v70_ = (d_68_v2_) != (default__.fm2(-764, d_71_globalState_))
                        d_219_v107_ = (0) - ((d_160_v61_).f6)
                        rhs21_ = (d_222_v109_.f20) or (d_171_v72_)
                        rhs22_ = d_160_v61_
                        rhs23_ = ((d_164_v65_)[default__.safeIndex(157, (d_164_v65_).length(0))]) <= ((d_68_v2_) - ((d_69_v3_)[default__.safeIndex((d_164_v65_)[default__.safeIndex(157, (d_164_v65_).length(0))], len(d_69_v3_))]))
                        rhs24_ = d_219_v107_
                        rhs25_ = d_156_v57_
                        lhs17_ = d_71_globalState_
                        lhs18_ = d_71_globalState_
                        d_73_v5_ = rhs21_
                        d_160_v61_ = rhs22_
                        lhs17_.f0 = rhs23_
                        lhs18_.f5 = rhs24_
                        d_156_v57_ = rhs25_
                        index26_ = default__.safeIndex(157, (d_164_v65_).length(0))
                        (d_164_v65_)[index26_] = d_68_v2_
                        d_155_v56_ = d_155_v56_
                    pass
            pass
        _dafny.print(_dafny.string_of(d_66_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_67_v1_) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f'), _dafny.CodePoint('s')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_68_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_69_v3_) == (_dafny.SeqWithoutIsStrInference([-2, -354, 1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_70_v4_) == (_dafny.Map({1: 8}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_71_globalState_.f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_).f1) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_71_globalState_.f2).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_71_globalState_).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_).f4) == (_dafny.Map({1: 8}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_71_globalState_.f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_73_v5_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_74_v6_).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_74_v6_).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_74_v6_).cf4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_153_v54_) == (_dafny.Map({False: 663}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_154_v55_) == (_dafny.Map({_dafny.Map({False: 663}): _dafny.MultiSet([True])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_155_v56_) == (_dafny.SeqWithoutIsStrInference([False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_156_v57_) == (_dafny.Set({-291, 663}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_157_v58_) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_158_v59_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_158_v59_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_158_v59_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_158_v59_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_159_v60_) == (_dafny.Map({291: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f'), _dafny.CodePoint('s')])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_160_v61_.f9)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_160_v61_.f9)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_160_v61_.f9)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_160_v61_.f9)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_160_v61_).f10) == (_dafny.Map({291: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f'), _dafny.CodePoint('s')])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_160_v61_).f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_161_v62_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_162_v63_) == (_dafny.Set({_dafny.CodePoint('s')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_163_v64_) == (_dafny.MultiSet([_dafny.Set({False})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_v65_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_v65_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_v65_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_v65_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_v65_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_v65_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_v65_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_v65_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_v65_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_v65_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_v65_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_v65_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_v65_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_v65_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_v65_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_v65_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_v65_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_v65_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_v65_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_v65_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_v65_)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_v65_)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_v65_)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_v65_)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_v65_)[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_165_v66_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_166_v67_).cf52)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_167_v68_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v69_) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yhclj")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jcnirnr")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ttxfurlkd"))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_169_v70_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_170_v71_) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yhclj")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jcnirnr")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ttxfurlkd"))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_171_v72_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_172_i3_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_215_v103_) == (_dafny.MultiSet([291, 291]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_216_v104_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_217_v105_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_218_v106_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_218_v106_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_218_v106_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_218_v106_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_218_v106_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_218_v106_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_218_v106_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_219_v107_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_220_i13_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC0(_dafny.Seq({}), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D0_DC1)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D0_DC2)

class D0_DC0(D0, NamedTuple('DC0', [('cf0', Any), ('cf1', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC0({_dafny.string_of(self.cf0)}, {_dafny.string_of(self.cf1)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC0) and self.cf0 == __o.cf0 and self.cf1 == __o.cf1
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC1(D0, NamedTuple('DC1', [('cf2', Any), ('cf3', Any), ('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)}, {_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf2 == __o.cf2 and self.cf3 == __o.cf3 and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC2(D0, NamedTuple('DC2', [('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2({_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2) and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()


class D1:
    @classmethod
    def default(cls, ):
        return lambda: D1_DC4(D0.default()(), int(0), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)

class D1_DC4(D1, NamedTuple('DC4', [('cf7', Any), ('cf8', Any), ('cf9', Any), ('cf10', Any), ('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)}, {self.cf9.VerbatimString(True)}, {_dafny.string_of(self.cf10)}, {_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf7 == __o.cf7 and self.cf8 == __o.cf8 and self.cf9 == __o.cf9 and self.cf10 == __o.cf10 and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC3(D1, NamedTuple('DC3', [('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC6(int(0), _dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D2_DC8)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D2_DC5)

class D2_DC6(D2, NamedTuple('DC6', [('cf13', Any), ('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf13)}, {_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf13 == __o.cf13 and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC7(D2, NamedTuple('DC7', [('cf15', Any), ('cf16', Any), ('cf17', Any), ('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf15)}, {_dafny.string_of(self.cf16)}, {_dafny.string_of(self.cf17)}, {_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf15 == __o.cf15 and self.cf16 == __o.cf16 and self.cf17 == __o.cf17 and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC8(D2, NamedTuple('DC8', [('cf19', Any), ('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC8({_dafny.string_of(self.cf19)}, {_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC8) and self.cf19 == __o.cf19 and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC5(D2, NamedTuple('DC5', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC5({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC5) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC10(_dafny.Map({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D3_DC10)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D3_DC9)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D3_DC11)

class D3_DC10(D3, NamedTuple('DC10', [('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC10({_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC10) and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC9(D3, NamedTuple('DC9', [('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC11(D3, NamedTuple('DC11', [('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC11({_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC11) and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC13(_dafny.Seq({}), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D4_DC13)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D4_DC14)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D4_DC15)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D4_DC12)

class D4_DC13(D4, NamedTuple('DC13', [('cf25', Any), ('cf26', Any), ('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC13({_dafny.string_of(self.cf25)}, {_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC13) and self.cf25 == __o.cf25 and self.cf26 == __o.cf26 and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC14(D4, NamedTuple('DC14', [('cf28', Any), ('cf29', Any), ('cf30', Any), ('cf31', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC14({_dafny.string_of(self.cf28)}, {_dafny.string_of(self.cf29)}, {_dafny.string_of(self.cf30)}, {_dafny.string_of(self.cf31)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC14) and self.cf28 == __o.cf28 and self.cf29 == __o.cf29 and self.cf30 == __o.cf30 and self.cf31 == __o.cf31
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC15(D4, NamedTuple('DC15', [('cf32', Any), ('cf33', Any), ('cf34', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC15({_dafny.string_of(self.cf32)}, {_dafny.string_of(self.cf33)}, {self.cf34.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC15) and self.cf32 == __o.cf32 and self.cf33 == __o.cf33 and self.cf34 == __o.cf34
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC12(D4, NamedTuple('DC12', [('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC12({_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC12) and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC17(_dafny.Seq({}), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D5_DC17)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D5_DC16)

class D5_DC17(D5, NamedTuple('DC17', [('cf36', Any), ('cf37', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC17({_dafny.string_of(self.cf36)}, {_dafny.string_of(self.cf37)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC17) and self.cf36 == __o.cf36 and self.cf37 == __o.cf37
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC16(D5, NamedTuple('DC16', [('cf35', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC16({_dafny.string_of(self.cf35)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC16) and self.cf35 == __o.cf35
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC19(int(0), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D6_DC19)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D6_DC18)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D6_DC20)

class D6_DC19(D6, NamedTuple('DC19', [('cf39', Any), ('cf40', Any), ('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC19({_dafny.string_of(self.cf39)}, {_dafny.string_of(self.cf40)}, {_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC19) and self.cf39 == __o.cf39 and self.cf40 == __o.cf40 and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC18(D6, NamedTuple('DC18', [('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC18({_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC18) and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC20(D6, NamedTuple('DC20', [('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC20({_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC20) and self.cf42 == __o.cf42
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

class D7_DC21(D7, NamedTuple('DC21', [('cf43', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC21({_dafny.string_of(self.cf43)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC21) and self.cf43 == __o.cf43
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D8_DC23)

class D8_DC23(D8, NamedTuple('DC23', [('cf44', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC23({_dafny.string_of(self.cf44)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC23) and self.cf44 == __o.cf44
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC25(int(0), int(0), int(0), _dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D9_DC25)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D9_DC26)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D9_DC24)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D9_DC27)

class D9_DC25(D9, NamedTuple('DC25', [('cf46', Any), ('cf47', Any), ('cf48', Any), ('cf49', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC25({_dafny.string_of(self.cf46)}, {_dafny.string_of(self.cf47)}, {_dafny.string_of(self.cf48)}, {_dafny.string_of(self.cf49)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC25) and self.cf46 == __o.cf46 and self.cf47 == __o.cf47 and self.cf48 == __o.cf48 and self.cf49 == __o.cf49
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC26(D9, NamedTuple('DC26', [('cf50', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC26({_dafny.string_of(self.cf50)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC26) and self.cf50 == __o.cf50
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC24(D9, NamedTuple('DC24', [('cf45', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC24({_dafny.string_of(self.cf45)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC24) and self.cf45 == __o.cf45
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC27(D9, NamedTuple('DC27', [('cf51', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC27({_dafny.string_of(self.cf51)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC27) and self.cf51 == __o.cf51
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: D10_DC29(int(0), _dafny.Array(None, 0), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D10_DC29)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D10_DC28)

class D10_DC29(D10, NamedTuple('DC29', [('cf53', Any), ('cf54', Any), ('cf55', Any), ('cf56', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC29({_dafny.string_of(self.cf53)}, {_dafny.string_of(self.cf54)}, {_dafny.string_of(self.cf55)}, {_dafny.string_of(self.cf56)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC29) and self.cf53 == __o.cf53 and self.cf54 == __o.cf54 and self.cf55 == __o.cf55 and self.cf56 == __o.cf56
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC28(D10, NamedTuple('DC28', [('cf52', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC28({_dafny.string_of(self.cf52)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC28) and self.cf52 == __o.cf52
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC31(int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D11_DC31)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D11_DC32)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D11_DC33)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D11_DC30)

class D11_DC31(D11, NamedTuple('DC31', [('cf58', Any), ('cf59', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC31({_dafny.string_of(self.cf58)}, {_dafny.string_of(self.cf59)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC31) and self.cf58 == __o.cf58 and self.cf59 == __o.cf59
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC32(D11, NamedTuple('DC32', [('cf60', Any), ('cf61', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC32({_dafny.string_of(self.cf60)}, {self.cf61.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC32) and self.cf60 == __o.cf60 and self.cf61 == __o.cf61
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC33(D11, NamedTuple('DC33', [('cf62', Any), ('cf63', Any), ('cf64', Any), ('cf65', Any), ('cf66', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC33({_dafny.string_of(self.cf62)}, {_dafny.string_of(self.cf63)}, {_dafny.string_of(self.cf64)}, {_dafny.string_of(self.cf65)}, {_dafny.string_of(self.cf66)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC33) and self.cf62 == __o.cf62 and self.cf63 == __o.cf63 and self.cf64 == __o.cf64 and self.cf65 == __o.cf65 and self.cf66 == __o.cf66
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC30(D11, NamedTuple('DC30', [('cf57', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC30({_dafny.string_of(self.cf57)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC30) and self.cf57 == __o.cf57
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D12_DC34)

class D12_DC34(D12, NamedTuple('DC34', [('cf67', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC34({_dafny.string_of(self.cf67)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC34) and self.cf67 == __o.cf67
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D13_DC35)

class D13_DC35(D13, NamedTuple('DC35', [('cf68', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC35({_dafny.string_of(self.cf68)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC35) and self.cf68 == __o.cf68
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: D14_DC37(False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D14_DC37)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D14_DC36)

class D14_DC37(D14, NamedTuple('DC37', [('cf70', Any), ('cf71', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC37({_dafny.string_of(self.cf70)}, {self.cf71.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC37) and self.cf70 == __o.cf70 and self.cf71 == __o.cf71
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC36(D14, NamedTuple('DC36', [('cf69', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC36({_dafny.string_of(self.cf69)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC36) and self.cf69 == __o.cf69
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D15_DC38)

class D15_DC38(D15, NamedTuple('DC38', [('cf72', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC38({_dafny.string_of(self.cf72)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC38) and self.cf72 == __o.cf72
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D16_DC39)

class D16_DC39(D16, NamedTuple('DC39', [('cf73', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC39({_dafny.string_of(self.cf73)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC39) and self.cf73 == __o.cf73
    def __hash__(self) -> int:
        return super().__hash__()


class D17:
    @classmethod
    def default(cls, ):
        return lambda: D17_DC41(False, _dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC41(self) -> bool:
        return isinstance(self, D17_DC41)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D17_DC40)

class D17_DC41(D17, NamedTuple('DC41', [('cf75', Any), ('cf76', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC41({_dafny.string_of(self.cf75)}, {_dafny.string_of(self.cf76)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC41) and self.cf75 == __o.cf75 and self.cf76 == __o.cf76
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC40(D17, NamedTuple('DC40', [('cf74', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC40({_dafny.string_of(self.cf74)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC40) and self.cf74 == __o.cf74
    def __hash__(self) -> int:
        return super().__hash__()


class D18:
    @classmethod
    def default(cls, ):
        return lambda: D18_DC43(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC43(self) -> bool:
        return isinstance(self, D18_DC43)
    @property
    def is_DC44(self) -> bool:
        return isinstance(self, D18_DC44)
    @property
    def is_DC42(self) -> bool:
        return isinstance(self, D18_DC42)

class D18_DC43(D18, NamedTuple('DC43', [('cf78', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC43({_dafny.string_of(self.cf78)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC43) and self.cf78 == __o.cf78
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC44(D18, NamedTuple('DC44', [('cf79', Any), ('cf80', Any), ('cf81', Any), ('cf82', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC44({_dafny.string_of(self.cf79)}, {_dafny.string_of(self.cf80)}, {_dafny.string_of(self.cf81)}, {_dafny.string_of(self.cf82)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC44) and self.cf79 == __o.cf79 and self.cf80 == __o.cf80 and self.cf81 == __o.cf81 and self.cf82 == __o.cf82
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC42(D18, NamedTuple('DC42', [('cf77', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC42({_dafny.string_of(self.cf77)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC42) and self.cf77 == __o.cf77
    def __hash__(self) -> int:
        return super().__hash__()


class D19:
    @classmethod
    def default(cls, ):
        return lambda: None
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC45(self) -> bool:
        return isinstance(self, D19_DC45)

class D19_DC45(D19, NamedTuple('DC45', [('cf83', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC45({_dafny.string_of(self.cf83)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC45) and self.cf83 == __o.cf83
    def __hash__(self) -> int:
        return super().__hash__()


class D20:
    @classmethod
    def default(cls, ):
        return lambda: D20_DC47(False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC47(self) -> bool:
        return isinstance(self, D20_DC47)
    @property
    def is_DC48(self) -> bool:
        return isinstance(self, D20_DC48)
    @property
    def is_DC46(self) -> bool:
        return isinstance(self, D20_DC46)

class D20_DC47(D20, NamedTuple('DC47', [('cf85', Any), ('cf86', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC47({_dafny.string_of(self.cf85)}, {_dafny.string_of(self.cf86)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC47) and self.cf85 == __o.cf85 and self.cf86 == __o.cf86
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC48(D20, NamedTuple('DC48', [('cf87', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC48({_dafny.string_of(self.cf87)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC48) and self.cf87 == __o.cf87
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC46(D20, NamedTuple('DC46', [('cf84', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC46({_dafny.string_of(self.cf84)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC46) and self.cf84 == __o.cf84
    def __hash__(self) -> int:
        return super().__hash__()


class D21:
    @classmethod
    def default(cls, ):
        return lambda: D21_DC50(False, int(0), False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC50(self) -> bool:
        return isinstance(self, D21_DC50)
    @property
    def is_DC49(self) -> bool:
        return isinstance(self, D21_DC49)

class D21_DC50(D21, NamedTuple('DC50', [('cf89', Any), ('cf90', Any), ('cf91', Any), ('cf92', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC50({_dafny.string_of(self.cf89)}, {_dafny.string_of(self.cf90)}, {_dafny.string_of(self.cf91)}, {_dafny.string_of(self.cf92)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC50) and self.cf89 == __o.cf89 and self.cf90 == __o.cf90 and self.cf91 == __o.cf91 and self.cf92 == __o.cf92
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC49(D21, NamedTuple('DC49', [('cf88', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC49({_dafny.string_of(self.cf88)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC49) and self.cf88 == __o.cf88
    def __hash__(self) -> int:
        return super().__hash__()


class D22:
    @classmethod
    def default(cls, ):
        return lambda: D22_DC52(_dafny.Map({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC52(self) -> bool:
        return isinstance(self, D22_DC52)
    @property
    def is_DC53(self) -> bool:
        return isinstance(self, D22_DC53)
    @property
    def is_DC51(self) -> bool:
        return isinstance(self, D22_DC51)
    @property
    def is_DC54(self) -> bool:
        return isinstance(self, D22_DC54)

class D22_DC52(D22, NamedTuple('DC52', [('cf94', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC52({_dafny.string_of(self.cf94)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC52) and self.cf94 == __o.cf94
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC53(D22, NamedTuple('DC53', [('cf95', Any), ('cf96', Any), ('cf97', Any), ('cf98', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC53({_dafny.string_of(self.cf95)}, {_dafny.string_of(self.cf96)}, {_dafny.string_of(self.cf97)}, {_dafny.string_of(self.cf98)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC53) and self.cf95 == __o.cf95 and self.cf96 == __o.cf96 and self.cf97 == __o.cf97 and self.cf98 == __o.cf98
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC51(D22, NamedTuple('DC51', [('cf93', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC51({_dafny.string_of(self.cf93)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC51) and self.cf93 == __o.cf93
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC54(D22, NamedTuple('DC54', [('cf99', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC54({_dafny.string_of(self.cf99)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC54) and self.cf99 == __o.cf99
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass

class T1:
    pass
    @property
    def f9(self):
        return self._f9
    @f9.setter
    def f9(self, value):
        self._f9 = value
    def m6(self, p0, p1, p2, globalState):
        pass

    def m7(self, p0, p1, globalState):
        pass


class T2:
    pass
    @property
    def f16(self):
        return self._f16
    @f16.setter
    def f16(self, value):
        self._f16 = value

class GlobalState:
    def  __init__(self):
        self.f0: bool = False
        self.f2: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self.f5: int = int(0)
        self._f1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f3: bool = False
        self._f4: _dafny.Map = _dafny.Map({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5):
        (self).f0 = f0
        (self)._f1 = f1
        (self).f2 = f2
        (self)._f3 = f3
        (self)._f4 = f4
        (self).f5 = f5

    @property
    def f1(self):
        return self._f1
    @property
    def f3(self):
        return self._f3
    @property
    def f4(self):
        return self._f4

class C0:
    def  __init__(self):
        self.f19: _dafny.Array = _dafny.Array(None, 0)
        self.f20: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f19, f20):
        (self).f19 = f19
        (self).f20 = f20

    def fm23(self, p0, globalState):
        return 638

    def fm24(self, globalState):
        return _dafny.SeqWithoutIsStrInference([not((len(_dafny.Map({self.f20: len(_dafny.Map({_dafny.Map({True: len(_dafny.Map({len(_dafny.Map({len(_dafny.Set({False})): self.f20})): self.f20}))}): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_231_i0_ in range(default__.abs(-605))])}))}))) < (779)), self.f20, self.f20, self.f20])


class C1(T0):
    def  __init__(self):
        self._f6: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    @property
    def f6(self):
        return self._f6
    def ctor__(self, f6):
        (self)._f6 = f6

    def fm5(self, p0, globalState):
        source4_ = (D3_DC10(_dafny.Map({_dafny.Map({(self).f6: False}): False})) if False else D3_DC10(_dafny.Map({_dafny.Map({214: False}): False})))
        if source4_.is_DC10:
            d_232___mcc_h0_ = source4_.cf22
            d_233_cf22_ = d_232___mcc_h0_
            return _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([_dafny.CodePoint('h'), _dafny.CodePoint('f'), _dafny.CodePoint('r')])])
        elif source4_.is_DC9:
            d_234___mcc_h1_ = source4_.cf21
            d_235_cf21_ = d_234___mcc_h1_
            return (D4_DC12(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([_dafny.CodePoint('g')]), _dafny.MultiSet([_dafny.CodePoint('m'), _dafny.CodePoint('k'), _dafny.CodePoint('s'), _dafny.CodePoint('t')])]))).cf24
        elif True:
            d_236___mcc_h2_ = source4_.cf23
            d_237_cf23_ = d_236___mcc_h2_
            return (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([_dafny.CodePoint('r'), _dafny.CodePoint('q')])])) + (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([_dafny.CodePoint('k'), _dafny.CodePoint('i')])]))

    def fm6(self, p0, globalState):
        def iife20_():
            coll8_ = _dafny.Map()
            compr_8_: int
            for compr_8_ in ((_dafny.MultiSet([(self).f6, (self).f6, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pxn"))), (self).f6])) | (_dafny.MultiSet([(self).f6]))).Elements:
                d_238_v0_: int = compr_8_
                if (d_238_v0_) in ((_dafny.MultiSet([(self).f6, (self).f6, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pxn"))), (self).f6])) | (_dafny.MultiSet([(self).f6]))):
                    coll8_[(d_238_v0_) * (len(_dafny.Map({272: (self).f6})))] = (_dafny.SeqWithoutIsStrInference([(self).f6 for d_239_i0_ in range(default__.abs(256))])) + (_dafny.SeqWithoutIsStrInference([(self).f6]))
            return _dafny.Map(coll8_)
        return iife20_()
        

    def fm28(self, globalState):
        return ((0) - (len((_dafny.SeqWithoutIsStrInference([(self).f6])) + (_dafny.SeqWithoutIsStrInference([(self).f6, (self).f6]))))) <= (618)


class C2(T0, T2, T1):
    def  __init__(self):
        self._f16: _dafny.MultiSet = _dafny.MultiSet({})
        self._f9: _dafny.Array = _dafny.Array(None, 0)
        self._f6: int = int(0)
        self._f10: _dafny.Map = _dafny.Map({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    @property
    def f16(self):
        return self._f16
    @f16.setter
    def f16(self, value):
        self._f16 = value
    @property
    def f9(self):
        return self._f9
    @f9.setter
    def f9(self, value):
        self._f9 = value
    @property
    def f6(self):
        return self._f6
    @property
    def f10(self):
        return self._f10
    def ctor__(self, f6, f16, f9, f10):
        (self)._f6 = f6
        (self).f16 = f16
        (self).f9 = f9
        (self)._f10 = f10

    def fm5(self, p0, globalState):
        source5_ = D2_DC8(not(not(False)), False)
        if source5_.is_DC6:
            d_240___mcc_h0_ = source5_.cf13
            d_241___mcc_h1_ = source5_.cf14
            d_242_cf14_ = d_241___mcc_h1_
            d_243_cf13_ = d_240___mcc_h0_
            return (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s'), _dafny.CodePoint('h')]))])) + (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([_dafny.CodePoint('d')])]))
        elif source5_.is_DC7:
            d_244___mcc_h2_ = source5_.cf15
            d_245___mcc_h3_ = source5_.cf16
            d_246___mcc_h4_ = source5_.cf17
            d_247___mcc_h5_ = source5_.cf18
            d_248_cf18_ = d_247___mcc_h5_
            d_249_cf17_ = d_246___mcc_h4_
            d_250_cf16_ = d_245___mcc_h3_
            d_251_cf15_ = d_244___mcc_h2_
            return (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([_dafny.CodePoint('r'), _dafny.CodePoint('h')]), _dafny.MultiSet([_dafny.CodePoint('r')]), _dafny.MultiSet([_dafny.CodePoint('e'), _dafny.CodePoint('u')])])) + (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([_dafny.CodePoint('h')]), _dafny.MultiSet([_dafny.CodePoint('y')]), _dafny.MultiSet([_dafny.CodePoint('s')])]))
        elif source5_.is_DC8:
            d_252___mcc_h6_ = source5_.cf19
            d_253___mcc_h7_ = source5_.cf20
            d_254_cf20_ = d_253___mcc_h7_
            d_255_cf19_ = d_252___mcc_h6_
            return _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([_dafny.CodePoint('m'), _dafny.CodePoint('d'), _dafny.CodePoint('i'), _dafny.CodePoint('k'), _dafny.CodePoint('o')]), _dafny.MultiSet([_dafny.CodePoint('h')])])
        elif True:
            d_256___mcc_h8_ = source5_.cf12
            d_257_cf12_ = d_256___mcc_h8_
            return (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([_dafny.CodePoint('r'), _dafny.CodePoint('q')])])) + (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([_dafny.CodePoint('u')]) for d_258_i0_ in range(default__.abs(523))]))

    def fm6(self, p0, globalState):
        return _dafny.Map({(self).f6: _dafny.SeqWithoutIsStrInference([(self).f6])})

    def fm21(self, p0, globalState):
        return ((self).f6) - ((self).f6)

    def fm22(self, p0, p1, p2, globalState):
        if (_dafny.CodePoint('d')) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m'), _dafny.CodePoint('d')])):
            return True
        elif True:
            return ((0) - ((self).f6)) <= (248)

    def fm13(self, p0, p1, p2, p3, globalState):
        return (_dafny.Map({not(not(not(not(not(not(True)))))): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_259_i0_ in range(default__.abs(-918))])})) | ((_dafny.Map({not(False): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_260_i1_ in range(default__.abs(2))])})) | (_dafny.Map({True: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rxhrrtyr"))})))

    def m6(self, p0, p1, p2, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: bool = False
        d_261_v0_: _dafny.Map
        d_261_v0_ = _dafny.Map({(self).f6: not(p0)})
        d_262_v1_: _dafny.Set
        d_262_v1_ = _dafny.Set({(d_261_v0_) | (d_261_v0_)})
        d_263_v2_: _dafny.Map
        d_263_v2_ = _dafny.Map({91: (self).f6})
        d_264_v3_: D2
        d_264_v3_ = D2_DC5(d_263_v2_)
        d_265_v4_: _dafny.Set
        d_265_v4_ = _dafny.Set({p0})
        d_266_v5_: _dafny.MultiSet
        d_266_v5_ = _dafny.MultiSet([p0])
        d_267_v6_: _dafny.MultiSet
        d_267_v6_ = _dafny.MultiSet([(self).f6, (self).f6])
        d_268_v7_: _dafny.Map
        d_268_v7_ = _dafny.Map({p0: (d_267_v6_).set((self).f6, default__.abs(p2))})
        rhs26_ = ((len(d_265_v4_)) - (648)) >= ((self).f6)
        rhs27_ = ((d_262_v1_ if (self).fm22((d_266_v5_).set(p0, default__.abs((self).f6)), p0, (self).f6, globalState) else d_262_v1_)) - ((d_262_v1_ if p0 else d_262_v1_))
        rhs28_ = D2_DC5((d_263_v2_) | (d_263_v2_))
        rhs29_ = default__.safeModuloInt((self).f6, (((d_268_v7_)[p0] if (p0) in (d_268_v7_) else d_267_v6_)).cardinality)
        rhs30_ = p0
        lhs19_ = globalState
        lhs20_ = globalState
        lhs19_.f0 = rhs26_
        d_262_v1_ = rhs27_
        d_264_v3_ = rhs28_
        lhs20_.f5 = rhs29_
        r1 = rhs30_
        d_269_v8_: _dafny.Array
        nw34_ = _dafny.Array(None, 3)
        nw34_[int(0)] = (self).f6
        nw34_[int(1)] = (self).f6
        nw34_[int(2)] = (self).f6
        d_269_v8_ = nw34_
        d_270_v9_: _dafny.Set
        d_270_v9_ = _dafny.Set({(self).f6})
        d_271_v10_: _dafny.Seq
        d_271_v10_ = _dafny.SeqWithoutIsStrInference([d_270_v9_, d_270_v9_, d_270_v9_, d_270_v9_])
        index27_ = default__.safeIndex(48, (d_269_v8_).length(0))
        (d_269_v8_)[index27_] = len(((d_271_v10_).set(default__.safeIndex((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([p2]))).cardinality, len(d_271_v10_)), d_270_v9_)) + (d_271_v10_))
        index28_ = default__.safeIndex(48, (d_269_v8_).length(0))
        rhs31_ = ((p2 if (self).fm22(d_266_v5_, p0, (self).f6, globalState) else (self).f6)) >= (default__.safeModuloInt((self).f6, (self).f6))
        rhs32_ = True
        rhs33_ = (len(d_265_v4_) if (d_267_v6_).isdisjoint(d_267_v6_) else ((self).f6) + ((self).f6))
        rhs34_ = p0
        lhs21_ = globalState
        lhs22_ = globalState
        lhs23_ = d_269_v8_
        lhs24_ = default__.safeIndex(48, (d_269_v8_).length(0))
        lhs25_ = globalState
        lhs21_.f0 = rhs31_
        lhs22_.f0 = rhs32_
        lhs23_[lhs24_] = rhs33_
        lhs25_.f0 = rhs34_
        d_272_v11_: C1
        nw35_ = C1()
        nw35_.ctor__(-380)
        d_272_v11_ = nw35_
        (globalState).f0 = not ((p2) >= ((d_269_v8_)[default__.safeIndex(48, (d_269_v8_).length(0))])) or (p0)
        d_273_v12_: _dafny.Array
        nw36_ = _dafny.Array(_dafny.MultiSet({}), 24)
        d_273_v12_ = nw36_
        d_274_v13_: D2
        d_274_v13_ = D2_DC7(p1, d_273_v12_, (self).f6, len(d_265_v4_))
        d_275_v14_: _dafny.Map
        d_275_v14_ = _dafny.Map({d_274_v13_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gtvimpmrf"))})
        d_276_v15_: _dafny.Seq
        d_276_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ohcxxhxh"))
        d_277_v16_: str
        d_277_v16_ = _dafny.CodePoint('t')
        (globalState).f5 = ((d_269_v8_)[default__.safeIndex(48, (d_269_v8_).length(0))]) * ((len((((d_275_v14_)[d_274_v13_] if (d_274_v13_) in (d_275_v14_) else d_276_v15_)).set(default__.safeIndex((self).f6, len(((d_275_v14_)[d_274_v13_] if (d_274_v13_) in (d_275_v14_) else d_276_v15_))), d_277_v16_))) + (len(d_276_v15_)))
        arr2_ = self.f9
        index29_ = default__.safeIndex(375, (self.f9).length(0))
        arr2_[index29_] = p0
        arr3_ = self.f9
        index30_ = default__.safeIndex(375, (self.f9).length(0))
        arr3_[index30_] = not(p0)
        r0 = (default__.fm29(d_270_v9_, globalState)).set(default__.safeIndex(((self).f6 if p0 else (d_269_v8_)[default__.safeIndex(48, (d_269_v8_).length(0))]), len(default__.fm29(d_270_v9_, globalState))), (d_276_v15_) + (d_276_v15_))
        d_278_v17_: _dafny.Seq
        d_278_v17_ = _dafny.SeqWithoutIsStrInference([len(d_263_v2_)])
        r1 = (p2) >= ((len(d_278_v17_)) * ((d_269_v8_)[default__.safeIndex(48, (d_269_v8_).length(0))]))
        return r0, r1

    def m7(self, p0, p1, globalState):
        r0: str = _dafny.CodePoint('D')
        r1: int = int(0)
        r2: _dafny.Array = _dafny.Array(None, 0)
        r3: int = int(0)
        d_279_v0_: bool
        d_279_v0_ = True
        if not (d_279_v0_) or (not(not(d_279_v0_))):
            d_280_v1_: _dafny.Set
            d_280_v1_ = _dafny.Set({d_279_v0_})
            d_281_v2_: D5
            d_281_v2_ = D5_DC16(d_280_v1_)
            d_282_v3_: _dafny.Seq
            d_282_v3_ = _dafny.SeqWithoutIsStrInference([(d_281_v2_).cf35, (d_280_v1_) - (d_280_v1_), (d_280_v1_) - (d_280_v1_)])
            d_283_v4_: _dafny.Array
            nw37_ = _dafny.Array(_dafny.Set({}), 12)
            d_283_v4_ = nw37_
            index31_ = default__.safeIndex(282, (d_283_v4_).length(0))
            (d_283_v4_)[index31_] = (_dafny.Set({d_279_v0_})) - ((d_281_v2_).cf35)
            d_284_v5_: str
            d_284_v5_ = _dafny.CodePoint('v')
            d_285_v6_: _dafny.Map
            d_285_v6_ = _dafny.Map({not(d_279_v0_): (self).f6})
            index32_ = default__.safeIndex(282, (d_283_v4_).length(0))
            rhs35_ = d_284_v5_
            rhs36_ = (d_282_v3_ if d_279_v0_ else d_282_v3_)
            rhs37_ = (self).fm21(len((d_285_v6_) | (d_285_v6_)), globalState)
            rhs38_ = p1
            rhs39_ = ((d_280_v1_).intersection(d_280_v1_)).intersection(_dafny.Set({d_279_v0_, True, d_279_v0_}))
            lhs26_ = globalState
            lhs27_ = d_283_v4_
            lhs28_ = default__.safeIndex(282, (d_283_v4_).length(0))
            r0 = rhs35_
            d_282_v3_ = rhs36_
            lhs26_.f5 = rhs37_
            r3 = rhs38_
            lhs27_[lhs28_] = rhs39_
            d_286_v7_: _dafny.Map
            d_286_v7_ = _dafny.Map({d_279_v0_: not(d_279_v0_)})
            d_287_v8_: D4
            d_287_v8_ = D4_DC14(d_279_v0_, d_286_v7_, (self).f6, _dafny.Map({d_279_v0_: d_279_v0_}))
            d_288_v9_: _dafny.Seq
            d_288_v9_ = _dafny.SeqWithoutIsStrInference([False, d_279_v0_])
            d_289_v10_: _dafny.Map
            d_289_v10_ = _dafny.Map({d_287_v8_: d_288_v9_})
            d_289_v10_ = (d_289_v10_).set(d_287_v8_, (_dafny.SeqWithoutIsStrInference([True])).set(default__.safeIndex(p1, len(_dafny.SeqWithoutIsStrInference([True]))), not((d_288_v9_)[default__.safeIndex(p1, len(d_288_v9_))])))
            r3 = len((_dafny.Map({d_279_v0_: d_279_v0_})) | ((d_286_v7_) | (d_286_v7_)))
            d_290_v11_: _dafny.Seq
            d_290_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cidjjdvpa"))
            (globalState).f2 = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ik"))) + ((d_290_v11_) + (d_290_v11_))).set(default__.safeIndex(default__.safeModuloInt((self).f6, p1), len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ik"))) + ((d_290_v11_) + (d_290_v11_)))), d_284_v5_)
            (globalState).f0 = default__.fm3(globalState)
        elif True:
            (globalState).f0 = d_279_v0_
            if not(d_279_v0_):
                (globalState).f0 = d_279_v0_
                d_291_v12_: _dafny.Array
                nw38_ = _dafny.Array(_dafny.Array(None, 0), 5)
                d_291_v12_ = nw38_
                d_292_v13_: _dafny.Array
                nw39_ = _dafny.Array(_dafny.Seq({}), 23)
                d_292_v13_ = nw39_
                index33_ = default__.safeIndex(772, (d_291_v12_).length(0))
                (d_291_v12_)[index33_] = d_292_v13_
                index34_ = default__.safeIndex(772, (d_291_v12_).length(0))
                (d_291_v12_)[index34_] = d_292_v13_
                d_293_v14_: _dafny.Seq
                d_293_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rdc"))
                d_279_v0_ = (d_293_v14_) == (d_293_v14_)
                d_294_v15_: _dafny.Array
                nw40_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 12)
                d_294_v15_ = nw40_
                index35_ = default__.safeIndex(485, (d_294_v15_).length(0))
                (d_294_v15_)[index35_] = d_293_v14_
                d_295_v16_: _dafny.Array
                nw41_ = _dafny.Array(int(0), 20)
                d_295_v16_ = nw41_
                d_296_v17_: _dafny.Seq
                d_296_v17_ = _dafny.SeqWithoutIsStrInference([p1])
                index36_ = default__.safeIndex(246, (d_295_v16_).length(0))
                (d_295_v16_)[index36_] = (0) - ((p1 if d_279_v0_ else len(d_296_v17_)))
                d_297_v18_: _dafny.Seq
                d_297_v18_ = _dafny.SeqWithoutIsStrInference([True])
                d_298_v19_: D4
                d_298_v19_ = D4_DC13(_dafny.SeqWithoutIsStrInference([-406]), d_279_v0_, p1)
                d_299_v20_: _dafny.Map
                d_299_v20_ = _dafny.Map({d_298_v19_: len(d_296_v17_)})
                d_300_v21_: _dafny.Map
                d_300_v21_ = _dafny.Map({((d_299_v20_)[d_298_v19_] if (d_298_v19_) in (d_299_v20_) else p1): (self).f6})
                d_301_v22_: str
                d_301_v22_ = _dafny.CodePoint('c')
                index37_ = default__.safeIndex(485, (d_294_v15_).length(0))
                index38_ = default__.safeIndex(246, (d_295_v16_).length(0))
                rhs40_ = len(d_296_v17_)
                rhs41_ = default__.safeModuloInt((self).f6, default__.safeDivisionInt(p1, len(d_297_v18_)))
                rhs42_ = (d_293_v14_).set(default__.safeIndex(((d_300_v21_)[(0) - (len(d_300_v21_))] if ((0) - (len(d_300_v21_))) in (d_300_v21_) else p1), len(d_293_v14_)), d_301_v22_)
                rhs43_ = (p1) * (p1)
                rhs44_ = p1
                lhs29_ = globalState
                lhs30_ = globalState
                lhs31_ = d_294_v15_
                lhs32_ = default__.safeIndex(485, (d_294_v15_).length(0))
                lhs33_ = d_295_v16_
                lhs34_ = default__.safeIndex(246, (d_295_v16_).length(0))
                lhs35_ = globalState
                lhs29_.f5 = rhs40_
                lhs30_.f5 = rhs41_
                lhs31_[lhs32_] = rhs42_
                lhs33_[lhs34_] = rhs43_
                lhs35_.f5 = rhs44_
                index39_ = default__.safeIndex(246, (d_295_v16_).length(0))
                (d_295_v16_)[index39_] = p1
            elif True:
                d_302_v23_: str
                d_302_v23_ = _dafny.CodePoint('j')
                d_303_v24_: _dafny.Seq
                d_303_v24_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tgruuabrh"))
                d_304_v25_: _dafny.Seq
                d_304_v25_ = _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yjxblq"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kfxtlw"))), (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_305_i0_ in range(default__.abs(293))]) if d_279_v0_ else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yiaoosnib"))), default__.fm30(True, d_302_v23_, globalState), d_303_v24_])
                (globalState).f2 = (d_304_v25_)[default__.safeIndex((self).f6, len(d_304_v25_))]
                (globalState).f0 = not(not(d_279_v0_))
                d_279_v0_ = ((self).f6) != ((self).f6)
                (globalState).f2 = ((d_303_v24_) + (d_303_v24_)) + (d_303_v24_)
                (globalState).f0 = d_279_v0_
            d_306_v26_: _dafny.Map
            d_306_v26_ = _dafny.Map({d_279_v0_: (self).f6})
            d_307_v27_: _dafny.Seq
            d_307_v27_ = _dafny.SeqWithoutIsStrInference([(d_306_v26_) | (d_306_v26_)])
            rhs45_ = (d_306_v26_) | (default__.fm31(d_279_v0_, (self).f6, default__.fm2(p1, globalState), _dafny.SeqWithoutIsStrInference([d_279_v0_]), globalState))
            rhs46_ = _dafny.SeqWithoutIsStrInference([d_306_v26_, (d_306_v26_) | (d_306_v26_), d_306_v26_, d_306_v26_])
            d_306_v26_ = rhs45_
            d_307_v27_ = rhs46_
            d_308_v28_: _dafny.Seq
            d_308_v28_ = _dafny.SeqWithoutIsStrInference([d_279_v0_])
            d_309_v29_: _dafny.Map
            d_309_v29_ = _dafny.Map({(d_308_v28_) + (d_308_v28_): d_279_v0_})
            d_309_v29_ = d_309_v29_
            d_310_v30_: _dafny.Map
            d_310_v30_ = _dafny.Map({d_279_v0_: d_279_v0_})
            d_311_v31_: str
            d_311_v31_ = _dafny.CodePoint('i')
            d_312_v32_: _dafny.MultiSet
            d_312_v32_ = _dafny.MultiSet([d_311_v31_, d_311_v31_, d_311_v31_, d_311_v31_, _dafny.CodePoint('m')])
            d_313_v33_: _dafny.MultiSet
            d_313_v33_ = _dafny.MultiSet([not(d_279_v0_), not(d_279_v0_)])
            d_314_v34_: D4
            d_314_v34_ = D4_DC14(True, _dafny.Map({d_279_v0_: True}), -428, d_310_v30_)
            d_315_v35_: _dafny.Set
            d_315_v35_ = _dafny.Set({True, not(d_279_v0_), (d_314_v34_).cf28, d_279_v0_, True})
            d_316_v36_: _dafny.Array
            nw42_ = _dafny.Array(None, 14)
            nw42_[int(0)] = len(d_310_v30_)
            nw42_[int(1)] = len(_dafny.Map({d_311_v31_: d_279_v0_}))
            nw42_[int(2)] = ((self).f6) - (p1)
            nw42_[int(3)] = ((d_312_v32_)[d_311_v31_] if (d_311_v31_) in (d_312_v32_) else p1)
            nw42_[int(4)] = len(d_308_v28_)
            nw42_[int(5)] = default__.fm2((self).f6, globalState)
            nw42_[int(6)] = (0) - (p1)
            nw42_[int(7)] = len(_dafny.Map({(self).f6: d_279_v0_}))
            nw42_[int(8)] = (default__.fm2(p1, globalState)) - ((d_313_v33_).cardinality)
            nw42_[int(9)] = p1
            nw42_[int(10)] = (self).f6
            nw42_[int(11)] = (self).fm21(p1, globalState)
            nw42_[int(12)] = p1
            nw42_[int(13)] = (len(d_315_v35_) if d_279_v0_ else (self).f6)
            d_316_v36_ = nw42_
            index40_ = default__.safeIndex(981, (d_316_v36_).length(0))
            (d_316_v36_)[index40_] = -446
            d_317_v37_: D0
            d_317_v37_ = D0_DC0(d_308_v28_, 593)
            d_318_v38_: D1
            d_318_v38_ = D1_DC4(d_317_v37_, (self).fm21(p1, globalState), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rdeikwgw")), d_279_v0_, p1)
            d_319_v39_: _dafny.Seq
            d_319_v39_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s"))
            d_320_v40_: _dafny.Map
            d_320_v40_ = _dafny.Map({d_279_v0_: (d_319_v39_)[default__.safeIndex(len(d_319_v39_), len(d_319_v39_))]})
            pat_let_tv3_ = d_320_v40_
            pat_let_tv4_ = d_319_v39_
            d_321_v42_: D1
            def iife21_(_pat_let6_0):
                def iife22_(d_322_dt__update__tmp_h0_):
                    def iife23_(_pat_let7_0):
                        def iife24_(d_323_dt__update_hcf8_h0_):
                            def iife25_(_pat_let8_0):
                                def iife26_(d_324_dt__update_hcf9_h0_):
                                    return D1_DC4((d_322_dt__update__tmp_h0_).cf7, d_323_dt__update_hcf8_h0_, d_324_dt__update_hcf9_h0_, (d_322_dt__update__tmp_h0_).cf10, (d_322_dt__update__tmp_h0_).cf11)
                                return iife26_(_pat_let8_0)
                            return iife25_(pat_let_tv4_)
                        return iife24_(_pat_let7_0)
                    return iife23_(len(pat_let_tv3_))
                return iife22_(_pat_let6_0)
            def iife27_():
                coll9_ = _dafny.Map()
                compr_9_: int
                for compr_9_ in _dafny.IntegerRange(401, 516):
                    d_325_v41_: int = compr_9_
                    if ((401) <= (d_325_v41_)) and ((d_325_v41_) < (516)):
                        coll9_[(d_325_v41_) + (p1)] = p1
                return _dafny.Map(coll9_)
            d_321_v42_ = D1_DC4(d_317_v37_, (self).f6, (iife21_(d_318_v38_)).cf9, d_279_v0_, len(iife27_()
))
            d_326_v43_: _dafny.Seq
            d_326_v43_ = _dafny.SeqWithoutIsStrInference([d_321_v42_, d_321_v42_])
            index41_ = default__.safeIndex(24, (d_316_v36_).length(0))
            (d_316_v36_)[index41_] = p1
            d_327_v44_: _dafny.Set
            d_327_v44_ = _dafny.Set({d_316_v36_})
            d_328_v45_: _dafny.Map
            d_328_v45_ = _dafny.Map({_dafny.Map({(self).f6: d_279_v0_}): d_279_v0_})
            d_329_v46_: D3
            d_329_v46_ = D3_DC10(d_328_v45_)
            d_330_v47_: _dafny.Set
            d_330_v47_ = _dafny.Set({_dafny.MultiSet([d_329_v46_, d_329_v46_, D3_DC10(d_328_v45_)])})
            index42_ = default__.safeIndex(981, (d_316_v36_).length(0))
            index43_ = default__.safeIndex(24, (d_316_v36_).length(0))
            rhs47_ = (self).f6
            rhs48_ = _dafny.SeqWithoutIsStrInference([D1_DC4(D0_DC0(d_308_v28_, -852), (self).f6, d_319_v39_, d_279_v0_, (self).f6)])
            rhs49_ = len(((d_330_v47_) - (d_330_v47_)) | (d_330_v47_))
            rhs50_ = _dafny.Set({d_316_v36_, d_316_v36_, d_316_v36_, d_316_v36_})
            lhs36_ = d_316_v36_
            lhs37_ = default__.safeIndex(981, (d_316_v36_).length(0))
            lhs38_ = d_316_v36_
            lhs39_ = default__.safeIndex(24, (d_316_v36_).length(0))
            lhs36_[lhs37_] = rhs47_
            d_326_v43_ = rhs48_
            lhs38_[lhs39_] = rhs49_
            d_327_v44_ = rhs50_
        d_331_v48_: _dafny.Array
        def lambda5_(d_332_i1_):
            return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_333_i2_ in range(default__.abs(651))])

        init3_ = lambda5_
        nw43_ = _dafny.Array(None, 11)
        for i0_3_ in range(nw43_.length(0)):
            nw43_[i0_3_] = init3_(i0_3_)
        d_331_v48_ = nw43_
        d_334_v49_: _dafny.Seq
        d_334_v49_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aabonynb"))
        index44_ = default__.safeIndex(739, (d_331_v48_).length(0))
        (d_331_v48_)[index44_] = (d_334_v49_) + (d_334_v49_)
        d_335_v50_: _dafny.Map
        d_335_v50_ = _dafny.Map({default__.safeDivisionInt((self).f6, (self).f6): len(d_334_v49_)})
        d_336_v51_: _dafny.Map
        d_336_v51_ = _dafny.Map({d_279_v0_: True})
        index45_ = default__.safeIndex(739, (d_331_v48_).length(0))
        rhs51_ = ((d_335_v50_)[(731) - ((self).f6)] if ((731) - ((self).f6)) in (d_335_v50_) else (self).f6)
        rhs52_ = d_279_v0_
        rhs53_ = (p1) <= (len(d_336_v51_))
        rhs54_ = (d_334_v49_) + (d_334_v49_)
        lhs40_ = d_331_v48_
        lhs41_ = default__.safeIndex(739, (d_331_v48_).length(0))
        r1 = rhs51_
        d_279_v0_ = rhs52_
        d_279_v0_ = rhs53_
        lhs40_[lhs41_] = rhs54_
        index46_ = default__.safeIndex(739, (d_331_v48_).length(0))
        (d_331_v48_)[index46_] = d_334_v49_
        hi1_ = default__.safeModuloInt((self).f6, p1)
        for d_337_i3_ in range(p1, hi1_):
            d_338_v52_: C0
            nw44_ = C0()
            nw44_.ctor__(self.f9, (d_279_v0_ if d_279_v0_ else d_279_v0_))
            d_338_v52_ = nw44_
            d_339_v53_: str
            d_339_v53_ = _dafny.CodePoint('t')
            d_340_v54_: _dafny.Map
            d_340_v54_ = _dafny.Map({d_339_v53_: d_338_v52_.f20})
            (d_338_v52_).f20 = ((d_340_v54_) | (d_340_v54_)) not in (_dafny.SeqWithoutIsStrInference([_dafny.Map({d_339_v53_: d_279_v0_})]))
            r3 = 15
            d_341_v55_: _dafny.MultiSet
            d_341_v55_ = _dafny.MultiSet([d_338_v52_.f20])
            d_342_v56_: D1
            def iife28_(_pat_let9_0):
                def iife29_(d_343_dt__update__tmp_h1_):
                    def iife30_(_pat_let10_0):
                        def iife31_(d_344_dt__update_hcf1_h0_):
                            return D0_DC0((d_343_dt__update__tmp_h1_).cf0, d_344_dt__update_hcf1_h0_)
                        return iife31_(_pat_let10_0)
                    return iife30_(811)
                return iife29_(_pat_let9_0)
            d_342_v56_ = D1_DC4(iife28_(default__.fm32(d_279_v0_, d_338_v52_.f20, d_337_i3_, d_279_v0_, globalState)), ((d_341_v55_)[d_338_v52_.f20] if (d_338_v52_.f20) in (d_341_v55_) else (self).f6), d_334_v49_, d_279_v0_, (0) - (p1))
            d_345_v57_: _dafny.Set
            d_345_v57_ = _dafny.Set({not(d_279_v0_)})
            d_346_v58_: D5
            d_346_v58_ = D5_DC16(d_345_v57_)
            pat_let_tv5_ = d_346_v58_
            d_347_v59_: _dafny.Map
            def iife32_(_pat_let11_0):
                def iife33_(d_348_dt__update__tmp_h2_):
                    def iife34_(_pat_let12_0):
                        def iife35_(d_349_dt__update_hcf8_h1_):
                            return D1_DC4((d_348_dt__update__tmp_h2_).cf7, d_349_dt__update_hcf8_h1_, (d_348_dt__update__tmp_h2_).cf9, (d_348_dt__update__tmp_h2_).cf10, (d_348_dt__update__tmp_h2_).cf11)
                        return iife35_(_pat_let12_0)
                    return iife34_(len((pat_let_tv5_).cf35))
                return iife33_(_pat_let11_0)
            d_347_v59_ = _dafny.Map({d_338_v52_.f20: iife32_(d_342_v56_)})
            d_347_v59_ = (d_347_v59_).set(not(d_279_v0_), d_342_v56_)
        nw45_ = _dafny.Array(False, 3)
        (self).f9 = nw45_
        d_350_v60_: _dafny.Seq
        d_350_v60_ = _dafny.SeqWithoutIsStrInference([(self).f6])
        d_351_v61_: D3
        d_351_v61_ = D3_DC9(_dafny.SeqWithoutIsStrInference([(self).f6, (self).f6]))
        d_352_v62_: _dafny.Seq
        d_352_v62_ = _dafny.SeqWithoutIsStrInference([self.f9, (self.f9 if d_279_v0_ else self.f9), self.f9, self.f9])
        d_353_v64_: D3
        def iife36_():
            coll10_ = _dafny.Map()
            compr_10_: int
            for compr_10_ in _dafny.IntegerRange(620, 163):
                d_354_v63_: int = compr_10_
                if ((620) <= (d_354_v63_)) and ((d_354_v63_) < (163)):
                    coll10_[(d_354_v63_) * (len(d_350_v60_))] = d_279_v0_
            return _dafny.Map(coll10_)
        d_353_v64_ = D3_DC10(_dafny.Map({iife36_()
: d_279_v0_}))
        pat_let_tv6_ = p1
        pat_let_tv7_ = p1
        pat_let_tv8_ = p1
        pat_let_tv9_ = d_279_v0_
        pat_let_tv10_ = d_279_v0_
        def lambda6_(source6_):
            if source6_.is_DC10:
                d_355___mcc_h0_ = source6_.cf22
                d_356_cf22_ = d_355___mcc_h0_
                return not(not(((0) - (pat_let_tv6_)) != (pat_let_tv7_)))
            elif source6_.is_DC9:
                d_357___mcc_h1_ = source6_.cf21
                d_358_cf21_ = d_357___mcc_h1_
                return ((0) - (pat_let_tv8_)) != ((self).f6)
            elif True:
                d_359___mcc_h2_ = source6_.cf23
                d_360_cf23_ = d_359___mcc_h2_
                return not (pat_let_tv9_) or (pat_let_tv10_)

        rhs55_ = (len((d_334_v49_) + (d_334_v49_))) - (default__.safeDivisionInt((self).f6, len(d_350_v60_)))
        rhs56_ = lambda6_(D3_DC11(D3_DC11(D3_DC11(D3_DC11(d_351_v61_)))))
        rhs57_ = (d_352_v62_)[default__.safeIndex((0) - (((self).f6) * ((self).f6)), len(d_352_v62_))]
        rhs58_ = not(d_279_v0_)
        rhs59_ = default__.safeModuloInt((self).f6, (len(default__.fm33(d_353_v64_, 969, d_279_v0_, d_279_v0_, globalState))) + ((self).f6))
        lhs42_ = self
        lhs43_ = globalState
        lhs44_ = globalState
        r3 = rhs55_
        d_279_v0_ = rhs56_
        lhs42_.f9 = rhs57_
        lhs43_.f0 = rhs58_
        lhs44_.f5 = rhs59_
        r0 = _dafny.CodePoint('j')
        r1 = default__.fm2(p1, globalState)
        r2 = self.f9
        r3 = (((self).f6) + (len(d_334_v49_))) - ((self).f6)
        return r0, r1, r2, r3

    def m15(self, p0, p1, p2, globalState):
        (globalState).f0 = not(p0)
        hi2_ = (442) - (553)
        for d_361_i0_ in range(832, hi2_):
            d_362_v0_: C1
            nw46_ = C1()
            nw46_.ctor__((self).f6)
            d_362_v0_ = nw46_
            d_363_v1_: _dafny.Array
            nw47_ = _dafny.Array(int(0), 6)
            d_363_v1_ = nw47_
            index47_ = default__.safeIndex(640, (d_363_v1_).length(0))
            (d_363_v1_)[index47_] = (d_361_i0_) + (d_361_i0_)
            index48_ = default__.safeIndex(374, (d_363_v1_).length(0))
            (d_363_v1_)[index48_] = d_361_i0_
            d_364_v2_: _dafny.Array
            def lambda7_(d_365_i1_):
                return (_dafny.Set({len(_dafny.SeqWithoutIsStrInference([(self).f6]))})) - (_dafny.Set({(self).f6}))

            init4_ = lambda7_
            nw48_ = _dafny.Array(None, 26)
            for i0_4_ in range(nw48_.length(0)):
                nw48_[i0_4_] = init4_(i0_4_)
            d_364_v2_ = nw48_
            d_366_v3_: _dafny.Set
            d_366_v3_ = _dafny.Set({d_361_i0_})
            index49_ = default__.safeIndex(386, (d_364_v2_).length(0))
            (d_364_v2_)[index49_] = (d_366_v3_).intersection(d_366_v3_)
            d_367_v4_: _dafny.Seq
            d_367_v4_ = _dafny.SeqWithoutIsStrInference([p1])
            index50_ = default__.safeIndex(640, (d_363_v1_).length(0))
            index51_ = default__.safeIndex(374, (d_363_v1_).length(0))
            index52_ = default__.safeIndex(386, (d_364_v2_).length(0))
            rhs60_ = (self).f6
            rhs61_ = (default__.safeDivisionInt(682, -15) if (d_367_v4_) == (d_367_v4_) else ((self).f6) - ((self).f6))
            rhs62_ = d_366_v3_
            lhs45_ = d_363_v1_
            lhs46_ = default__.safeIndex(640, (d_363_v1_).length(0))
            lhs47_ = d_363_v1_
            lhs48_ = default__.safeIndex(374, (d_363_v1_).length(0))
            lhs49_ = d_364_v2_
            lhs50_ = default__.safeIndex(386, (d_364_v2_).length(0))
            lhs45_[lhs46_] = rhs60_
            lhs47_[lhs48_] = rhs61_
            lhs49_[lhs50_] = rhs62_
            d_368_v5_: _dafny.Seq
            d_368_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fwds"))
            if not (True) or ((d_367_v4_)[default__.safeIndex(len(d_368_v5_), len(d_367_v4_))]):
                d_369_v6_: _dafny.Array
                def lambda8_(d_370_v5_):
                    def lambda9_(d_371_i2_):
                        return d_370_v5_

                    return lambda9_

                init5_ = lambda8_(d_368_v5_)
                nw49_ = _dafny.Array(None, 23)
                for i0_5_ in range(nw49_.length(0)):
                    nw49_[i0_5_] = init5_(i0_5_)
                d_369_v6_ = nw49_
                index53_ = default__.safeIndex(212, (d_369_v6_).length(0))
                (d_369_v6_)[index53_] = (d_368_v5_) + (d_368_v5_)
                index54_ = default__.safeIndex(212, (d_369_v6_).length(0))
                (d_369_v6_)[index54_] = d_368_v5_
                d_372_v7_: _dafny.Map
                d_372_v7_ = _dafny.Map({p1: True})
                d_373_v8_: _dafny.Seq
                d_373_v8_ = _dafny.SeqWithoutIsStrInference([len(d_372_v7_), d_361_i0_])
                d_374_v9_: D4
                d_374_v9_ = D4_DC13(d_373_v8_, p1, (self).f6)
                d_375_v10_: _dafny.Map
                d_375_v10_ = _dafny.Map({(self).fm21(d_361_i0_, globalState): (d_374_v9_).cf26})
                d_376_v11_: _dafny.Set
                d_376_v11_ = _dafny.Set({p0, False, p1})
                d_377_v12_: _dafny.Set
                d_377_v12_ = _dafny.Set({d_376_v11_, d_376_v11_})
                d_375_v10_ = (d_375_v10_).set((d_363_v1_)[default__.safeIndex(640, (d_363_v1_).length(0))], not((d_376_v11_) in (d_377_v12_)))
                (globalState).f5 = d_361_i0_
                index55_ = default__.safeIndex(640, (d_363_v1_).length(0))
                (d_363_v1_)[index55_] = (d_363_v1_)[default__.safeIndex(640, (d_363_v1_).length(0))]
                d_378_v13_: _dafny.Map
                d_378_v13_ = _dafny.Map({(d_363_v1_)[default__.safeIndex(640, (d_363_v1_).length(0))]: (d_363_v1_)[default__.safeIndex(640, (d_363_v1_).length(0))]})
                d_378_v13_ = (d_378_v13_).set(default__.safeModuloInt((self).f6, -29), (d_363_v1_)[default__.safeIndex(640, (d_363_v1_).length(0))])
            elif True:
                d_379_v14_: str
                d_379_v14_ = _dafny.CodePoint('p')
                d_380_v15_: _dafny.Map
                d_380_v15_ = _dafny.Map({(d_363_v1_)[default__.safeIndex(640, (d_363_v1_).length(0))]: p1})
                d_381_v16_: _dafny.Seq
                d_381_v16_ = _dafny.SeqWithoutIsStrInference([(d_363_v1_)[default__.safeIndex(640, (d_363_v1_).length(0))], len(d_380_v15_)])
                d_382_v17_: _dafny.Map
                d_382_v17_ = _dafny.Map({(D4_DC13(d_381_v16_, p1, d_361_i0_)).cf27: p0})
                d_383_v18_: _dafny.Map
                d_383_v18_ = _dafny.Map({d_379_v14_: len((d_382_v17_).set((self).f6, p0))})
                d_383_v18_ = (d_383_v18_).set(d_379_v14_, 43)
                (globalState).f0 = ((d_363_v1_)[default__.safeIndex(640, (d_363_v1_).length(0))]) in (_dafny.SeqWithoutIsStrInference([(self).f6 for d_384_i3_ in range(default__.abs(32))]))
                (globalState).f5 = (0) - ((len(d_367_v4_)) * ((0) - ((d_363_v1_)[default__.safeIndex(640, (d_363_v1_).length(0))])))
                index56_ = default__.safeIndex(640, (d_363_v1_).length(0))
                (d_363_v1_)[index56_] = (d_363_v1_)[default__.safeIndex(640, (d_363_v1_).length(0))]
                (globalState).f0 = p0
            d_385_v19_: _dafny.Map
            d_385_v19_ = _dafny.Map({d_361_i0_: p0})
            d_386_v20_: _dafny.Map
            d_386_v20_ = _dafny.Map({(self).f6: ((d_363_v1_)[default__.safeIndex(640, (d_363_v1_).length(0))] if not(((d_385_v19_)[947] if (947) in (d_385_v19_) else p1)) else 229)})
            d_386_v20_ = d_386_v20_
        d_387_v21_: _dafny.Set
        d_387_v21_ = _dafny.Set({self.f9})
        d_388_v22_: _dafny.Map
        d_388_v22_ = _dafny.Map({p0: len(d_387_v21_)})
        d_389_v23_: _dafny.Seq
        d_389_v23_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([False]))])
        hi3_ = (len((d_388_v22_).set(p0, (self).f6))) - ((d_389_v23_)[default__.safeIndex((0) - ((self).f6), len(d_389_v23_))])
        for d_390_i4_ in range((self).f6, hi3_):
            d_391_v24_: C0
            nw50_ = C0()
            nw50_.ctor__(self.f9, p1)
            d_391_v24_ = nw50_
            d_392_v25_: C1
            nw51_ = C1()
            nw51_.ctor__((self).f6)
            d_392_v25_ = nw51_
            d_393_v26_: _dafny.Map
            d_393_v26_ = _dafny.Map({(self).f6: d_391_v24_.f19})
            d_394_v27_: str
            d_394_v27_ = _dafny.CodePoint('c')
            d_395_v28_: _dafny.Map
            d_395_v28_ = _dafny.Map({d_394_v27_: (self).f6})
            d_396_v29_: _dafny.Map
            d_396_v29_ = _dafny.Map({p0: _dafny.Set({len(_dafny.Set({d_390_i4_}))})})
            (d_391_v24_).f19 = ((d_393_v26_)[default__.safeModuloInt(len(d_395_v28_), len(d_396_v29_))] if (default__.safeModuloInt(len(d_395_v28_), len(d_396_v29_))) in (d_393_v26_) else self.f9)
            d_397_v30_: C1
            nw52_ = C1()
            nw52_.ctor__((self).f6)
            d_397_v30_ = nw52_
        d_398_v31_: _dafny.MultiSet
        d_398_v31_ = _dafny.MultiSet([d_389_v23_])
        d_399_v32_: _dafny.Seq
        d_399_v32_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "waawxeny"))
        d_400_v33_: _dafny.Map
        d_400_v33_ = _dafny.Map({d_399_v32_: _dafny.SeqWithoutIsStrInference([default__.fm2(len(d_399_v32_), globalState)])})
        d_401_v34_: _dafny.Map
        d_401_v34_ = _dafny.Map({p0: p1})
        d_402_v35_: D4
        d_402_v35_ = D4_DC15(p1, p1, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_403_i6_ in range(default__.abs(53))]))
        d_404_v36_: _dafny.MultiSet
        d_404_v36_ = _dafny.MultiSet([D4_DC15(p1, p0, d_399_v32_), d_402_v35_])
        d_405_v37_: _dafny.Seq
        d_405_v37_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(self).f6])])
        d_406_v38_: _dafny.Array
        nw53_ = _dafny.Array(None, 11)
        nw53_[int(0)] = ((d_398_v31_)[((d_400_v33_)[d_399_v32_] if (d_399_v32_) in (d_400_v33_) else d_389_v23_)] if (((d_400_v33_)[d_399_v32_] if (d_399_v32_) in (d_400_v33_) else d_389_v23_)) in (d_398_v31_) else len(d_401_v34_))
        nw53_[int(1)] = ((d_388_v22_)[p0] if (p0) in (d_388_v22_) else (self).f6)
        nw53_[int(2)] = ((d_404_v36_).cardinality) * (383)
        nw53_[int(3)] = (self).f6
        nw53_[int(4)] = default__.safeDivisionInt((self).f6, len(d_399_v32_))
        nw53_[int(5)] = len(d_389_v23_)
        nw53_[int(6)] = (len(_dafny.Map({(self).f6: (self).f6})) if p1 else (self).f6)
        nw53_[int(7)] = (self).f6
        nw53_[int(8)] = (self).f6
        nw53_[int(9)] = ((self).f6) * (len(default__.fm34(len((d_405_v37_)[default__.safeIndex((self).f6, len(d_405_v37_))]), globalState)))
        nw53_[int(10)] = (d_389_v23_)[default__.safeIndex(39, len(d_389_v23_))]
        d_406_v38_ = nw53_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_406_v38_).length(0)):
            d_407_i5_: int = guard_loop_1_
            if (True) and (((0) <= (d_407_i5_)) and ((d_407_i5_) < ((d_406_v38_).length(0)))):
                (d_406_v38_)[(d_407_i5_)] = (d_407_i5_) * (404)
        hi4_ = (self).f6
        for d_408_i7_ in range((0) - ((self).f6), hi4_):
            (globalState).f0 = p0
            d_409_v39_: _dafny.Seq
            d_409_v39_ = _dafny.SeqWithoutIsStrInference([p1])
            d_409_v39_ = d_409_v39_
            d_410_v40_: _dafny.Array
            nw54_ = _dafny.Array(None, 9)
            nw54_[int(0)] = _dafny.SeqWithoutIsStrInference([p1])
            nw54_[int(1)] = d_409_v39_
            nw54_[int(2)] = (d_409_v39_) + (d_409_v39_)
            nw54_[int(3)] = d_409_v39_
            nw54_[int(4)] = d_409_v39_
            nw54_[int(5)] = (d_409_v39_) + (d_409_v39_)
            nw54_[int(6)] = (_dafny.SeqWithoutIsStrInference([default__.fm3(globalState), p1, p0, p0, p1])) + (_dafny.SeqWithoutIsStrInference([p0, p1]))
            nw54_[int(7)] = d_409_v39_
            nw54_[int(8)] = (d_409_v39_) + (_dafny.SeqWithoutIsStrInference([True, p1]))
            d_410_v40_ = nw54_
            index57_ = default__.safeIndex(324, (d_410_v40_).length(0))
            (d_410_v40_)[index57_] = d_409_v39_
            index58_ = default__.safeIndex(324, (d_410_v40_).length(0))
            (d_410_v40_)[index58_] = d_409_v39_
            d_411_v41_: _dafny.Array
            nw55_ = _dafny.Array(_dafny.Array(None, 0), 1)
            d_411_v41_ = nw55_
            d_412_v42_: _dafny.Array
            nw56_ = _dafny.Array(None, 4)
            nw56_[int(0)] = self.f9
            nw56_[int(1)] = self.f9
            nw56_[int(2)] = self.f9
            nw56_[int(3)] = self.f9
            d_412_v42_ = nw56_
            index59_ = default__.safeIndex(291, (d_411_v41_).length(0))
            (d_411_v41_)[index59_] = d_412_v42_
            index60_ = default__.safeIndex(291, (d_411_v41_).length(0))
            (d_411_v41_)[index60_] = d_412_v42_
        (globalState).f0 = p0


class C3(T1, T0):
    def  __init__(self):
        self._f9: _dafny.Array = _dafny.Array(None, 0)
        self._f6: int = int(0)
        self._f10: _dafny.Map = _dafny.Map({})
        self._f21: _dafny.Map = _dafny.Map({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f9(self):
        return self._f9
    @f9.setter
    def f9(self, value):
        self._f9 = value
    @property
    def f6(self):
        return self._f6
    @property
    def f10(self):
        return self._f10
    def ctor__(self, f21, f9, f10, f6):
        (self)._f21 = f21
        (self).f9 = f9
        (self)._f10 = f10
        (self)._f6 = f6

    def fm13(self, p0, p1, p2, p3, globalState):
        return _dafny.Map({not(True): (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qhkakwoox")))})

    def fm5(self, p0, globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b')])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d'), _dafny.CodePoint('b'), _dafny.CodePoint('n'), _dafny.CodePoint('b')]))), (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m'), _dafny.CodePoint('g'), _dafny.CodePoint('f'), _dafny.CodePoint('p'), _dafny.CodePoint('e')]))).intersection(_dafny.MultiSet([_dafny.CodePoint('n')])), _dafny.MultiSet([_dafny.CodePoint('s')]), _dafny.MultiSet([_dafny.CodePoint('p'), _dafny.CodePoint('k'), _dafny.CodePoint('x'), _dafny.CodePoint('g'), _dafny.CodePoint('y')]), _dafny.MultiSet([_dafny.CodePoint('i')])])

    def fm6(self, p0, globalState):
        return _dafny.Map({(self).f6: _dafny.SeqWithoutIsStrInference([-371, (self).f6, (self).f6, (self).f6, (0) - ((self).f6)])})

    def m6(self, p0, p1, p2, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: bool = False
        d_413_v0_: _dafny.Seq
        d_413_v0_ = _dafny.SeqWithoutIsStrInference([(self).f6, -821])
        d_414_v1_: D3
        d_414_v1_ = D3_DC9(d_413_v0_)
        pat_let_tv11_ = d_413_v0_
        def iife37_(_pat_let13_0):
            def iife38_(d_415_dt__update__tmp_h0_):
                def iife39_(_pat_let14_0):
                    def iife40_(d_416_dt__update_hcf21_h0_):
                        return D3_DC9(d_416_dt__update_hcf21_h0_)
                    return iife40_(_pat_let14_0)
                return iife39_(pat_let_tv11_)
            return iife38_(_pat_let13_0)
        (globalState).f5 = (0) - (len((iife37_(d_414_v1_)).cf21))
        d_417_v2_: _dafny.Array
        nw57_ = _dafny.Array(int(0), 6)
        d_417_v2_ = nw57_
        index61_ = default__.safeIndex(281, (d_417_v2_).length(0))
        (d_417_v2_)[index61_] = p2
        index62_ = default__.safeIndex(281, (d_417_v2_).length(0))
        (d_417_v2_)[index62_] = (self).f6
        (globalState).f2 = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "djjjt"))
        d_413_v0_ = d_413_v0_
        d_418_v3_: _dafny.Array
        def lambda10_(d_419_v0_):
            def lambda11_(d_420_i0_):
                return _dafny.MultiSet(d_419_v0_)

            return lambda11_

        init6_ = lambda10_(d_413_v0_)
        nw58_ = _dafny.Array(None, 8)
        for i0_6_ in range(nw58_.length(0)):
            nw58_[i0_6_] = init6_(i0_6_)
        d_418_v3_ = nw58_
        d_421_v4_: D2
        d_421_v4_ = D2_DC7(p1, d_418_v3_, ((self).f6) - ((d_417_v2_)[default__.safeIndex(281, (d_417_v2_).length(0))]), ((self).f6) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "olpetvk")))))
        source7_ = d_421_v4_
        if source7_.is_DC6:
            d_422___mcc_h0_ = source7_.cf13
            d_423___mcc_h1_ = source7_.cf14
            d_424_cf14_ = d_423___mcc_h1_
            d_425_cf13_ = d_422___mcc_h0_
            arr4_ = self.f9
            index63_ = default__.safeIndex(633, (self.f9).length(0))
            arr4_[index63_] = p0
            d_426_v5_: _dafny.Map
            d_426_v5_ = _dafny.Map({self.f9: p0})
            d_427_v6_: D0
            d_427_v6_ = D0_DC0(p1, len(p1))
            d_428_v7_: _dafny.Set
            d_428_v7_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hjgjnji"))})
            pat_let_tv12_ = d_417_v2_
            pat_let_tv13_ = d_417_v2_
            d_429_v8_: D1
            def iife41_(_pat_let15_0):
                def iife42_(d_430_dt__update__tmp_h1_):
                    def iife43_(_pat_let16_0):
                        def iife44_(d_431_dt__update_hcf1_h0_):
                            return D0_DC0((d_430_dt__update__tmp_h1_).cf0, d_431_dt__update_hcf1_h0_)
                        return iife44_(_pat_let16_0)
                    return iife43_((pat_let_tv13_)[default__.safeIndex(281, (pat_let_tv12_).length(0))])
                return iife42_(_pat_let15_0)
            d_429_v8_ = D1_DC4(iife41_(d_427_v6_), d_425_cf13_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_432_i1_ in range(default__.abs(921))]), True, len(d_428_v7_))
            arr5_ = self.f9
            index64_ = default__.safeIndex(633, (self.f9).length(0))
            rhs63_ = len(((_dafny.Map({self.f9: p0})) | (d_426_v5_)) | (d_426_v5_))
            rhs64_ = (d_429_v8_).cf10
            rhs65_ = p0
            lhs51_ = globalState
            lhs52_ = globalState
            lhs53_ = self.f9
            lhs54_ = default__.safeIndex(633, (self.f9).length(0))
            lhs51_.f5 = rhs63_
            lhs52_.f0 = rhs64_
            lhs53_[lhs54_] = rhs65_
            index65_ = default__.safeIndex(281, (d_417_v2_).length(0))
            (d_417_v2_)[index65_] = d_425_cf13_
            d_433_v9_: _dafny.Map
            d_433_v9_ = _dafny.Map({default__.fm2(d_425_cf13_, globalState): p0})
            d_434_v10_: _dafny.Map
            d_434_v10_ = _dafny.Map({(self).f6: ((d_433_v9_)[(self).f6] if ((self).f6) in (d_433_v9_) else (self.f9)[default__.safeIndex(633, (self.f9).length(0))])})
            d_433_v9_ = (d_433_v9_).set(d_425_cf13_, p0)
            arr6_ = self.f9
            index66_ = default__.safeIndex(633, (self.f9).length(0))
            arr6_[index66_] = p0
        elif source7_.is_DC7:
            d_435___mcc_h2_ = source7_.cf15
            d_436___mcc_h3_ = source7_.cf16
            d_437___mcc_h4_ = source7_.cf17
            d_438___mcc_h5_ = source7_.cf18
            d_439_cf18_ = d_438___mcc_h5_
            d_440_cf17_ = d_437___mcc_h4_
            d_441_cf16_ = d_436___mcc_h3_
            d_442_cf15_ = d_435___mcc_h2_
            d_443_v11_: _dafny.Map
            d_443_v11_ = _dafny.Map({self.f9: p0})
            d_443_v11_ = (d_443_v11_).set(self.f9, p0)
            d_444_v12_: _dafny.MultiSet
            d_444_v12_ = _dafny.MultiSet([p0, p0, p0, not(True)])
            d_445_v13_: _dafny.Seq
            d_445_v13_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o"))
            rhs66_ = default__.safeDivisionInt((p2) + (((d_444_v12_).set(p0, default__.abs(d_439_cf18_))).cardinality), p2)
            rhs67_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lcenstrw"))) + (d_445_v13_)
            rhs68_ = (self).f6
            rhs69_ = p0
            lhs55_ = globalState
            lhs56_ = globalState
            lhs57_ = globalState
            lhs58_ = globalState
            lhs55_.f5 = rhs66_
            lhs56_.f2 = rhs67_
            lhs57_.f5 = rhs68_
            lhs58_.f0 = rhs69_
            d_446_v14_: _dafny.Set
            d_446_v14_ = _dafny.Set({False, p0})
            d_447_v15_: _dafny.Map
            d_447_v15_ = _dafny.Map({not(p0): d_446_v14_})
            d_447_v15_ = (d_447_v15_).set(p0, d_446_v14_)
            d_448_v16_: T0
            nw59_ = C1()
            nw59_.ctor__(default__.safeDivisionInt(d_440_cf17_, 67))
            d_448_v16_ = nw59_
            d_449_v17_: str
            d_449_v17_ = _dafny.CodePoint('o')
            d_450_v18_: _dafny.Map
            d_450_v18_ = _dafny.Map({(D0_DC1(p0, (d_448_v16_).f6, d_449_v17_)).cf2: (_dafny.MultiSet([p0, p0]) if p0 else _dafny.MultiSet(d_442_cf15_))})
            rhs70_ = d_448_v16_
            rhs71_ = (d_450_v18_ if p0 else d_450_v18_)
            d_448_v16_ = rhs70_
            d_450_v18_ = rhs71_
        elif source7_.is_DC8:
            d_451___mcc_h6_ = source7_.cf19
            d_452___mcc_h7_ = source7_.cf20
            d_453_cf20_ = d_452___mcc_h7_
            d_454_cf19_ = d_451___mcc_h6_
            def iife45_():
                coll11_ = _dafny.Set()
                compr_11_: _dafny.Seq
                for compr_11_ in (_dafny.SeqWithoutIsStrInference([(p1).set(default__.safeIndex(339, len(p1)), d_453_cf20_) for d_455_i2_ in range(default__.abs(619))])).Elements:
                    d_456_v19_: _dafny.Seq = compr_11_
                    if (d_456_v19_) in (_dafny.SeqWithoutIsStrInference([(p1).set(default__.safeIndex(339, len(p1)), d_453_cf20_) for d_455_i2_ in range(default__.abs(619))])):
                        coll11_ = coll11_.union(_dafny.Set([d_456_v19_]))
                return _dafny.Set(coll11_)
            if ((d_417_v2_)[default__.safeIndex(281, (d_417_v2_).length(0))]) == (len(iife45_()
            )):
                d_457_v20_: _dafny.Seq
                d_457_v20_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hgq"))
                (globalState).f2 = d_457_v20_
                d_458_v21_: int
                d_459_v22_: str
                d_460_v23_: bool
                out16_: int
                out17_: str
                out18_: bool
                out16_, out17_, out18_ = (self).m14(default__.fm3(globalState), True, globalState)
                d_458_v21_ = out16_
                d_459_v22_ = out17_
                d_460_v23_ = out18_
                (globalState).f5 = (p2) - ((self).f6)
                d_454_cf19_ = not((d_460_v23_) == (p0))
                d_461_v24_: C0
                nw60_ = C0()
                nw60_.ctor__(self.f9, p0)
                d_461_v24_ = nw60_
            elif True:
                d_462_v26_: str
                d_462_v26_ = _dafny.CodePoint('r')
                d_463_v27_: _dafny.Map
                d_463_v27_ = _dafny.Map({d_462_v26_: len(_dafny.SeqWithoutIsStrInference([d_462_v26_ for d_464_i3_ in range(default__.abs(-134))]))})
                d_465_v28_: _dafny.Map
                d_465_v28_ = _dafny.Map({p0: d_453_cf20_})
                d_466_v29_: _dafny.MultiSet
                d_466_v29_ = _dafny.MultiSet([len(d_465_v28_), p2, (self).f6])
                d_467_v30_: _dafny.Map
                d_467_v30_ = _dafny.Map({(d_417_v2_)[default__.safeIndex(281, (d_417_v2_).length(0))]: p0})
                def iife46_():
                    coll12_ = _dafny.Map()
                    compr_12_: str
                    for compr_12_ in ((d_463_v27_).set(d_462_v26_, (d_417_v2_)[default__.safeIndex(281, (d_417_v2_).length(0))])).keys.Elements:
                        d_468_v25_: str = compr_12_
                        if (d_468_v25_) in ((d_463_v27_).set(d_462_v26_, (d_417_v2_)[default__.safeIndex(281, (d_417_v2_).length(0))])):
                            coll12_[d_468_v25_] = p0
                    return _dafny.Map(coll12_)
                (globalState).f0 = (len(iife46_()
                )) not in ((d_466_v29_).set(len(d_467_v30_), default__.abs((d_417_v2_)[default__.safeIndex(281, (d_417_v2_).length(0))])))
                d_469_v31_: _dafny.Array
                nw61_ = _dafny.Array(_dafny.Seq({}), 2)
                d_469_v31_ = nw61_
                d_469_v31_ = d_469_v31_
                d_470_v32_: _dafny.Map
                d_470_v32_ = _dafny.Map({d_454_cf19_: self.f9})
                d_471_v33_: _dafny.Set
                d_471_v33_ = _dafny.Set({790})
                d_472_v34_: _dafny.MultiSet
                d_472_v34_ = _dafny.MultiSet([d_454_cf19_, d_454_cf19_])
                d_473_v35_: _dafny.MultiSet
                d_473_v35_ = _dafny.MultiSet([default__.fm35(p2, len(d_471_v33_), d_454_cf19_, len(p1), globalState), _dafny.MultiSet([((d_472_v34_).set(False, default__.abs((self).f6))).cardinality]), d_466_v29_])
                d_474_v36_: T1
                nw62_ = C2()
                nw62_.ctor__(p2, d_473_v35_, self.f9, (self).f10)
                d_474_v36_ = nw62_
                d_475_v37_: _dafny.Array
                def lambda12_(d_476_p1_):
                    def lambda13_(d_477_i4_):
                        return d_476_p1_

                    return lambda13_

                init7_ = lambda12_(p1)
                nw63_ = _dafny.Array(None, 5)
                for i0_7_ in range(nw63_.length(0)):
                    nw63_[i0_7_] = init7_(i0_7_)
                d_475_v37_ = nw63_
                index67_ = default__.safeIndex(992, (d_475_v37_).length(0))
                (d_475_v37_)[index67_] = p1
                index68_ = default__.safeIndex(992, (d_475_v37_).length(0))
                rhs72_ = d_470_v32_
                rhs73_ = (D6_DC18(d_474_v36_)).cf38
                rhs74_ = p1
                lhs59_ = d_475_v37_
                lhs60_ = default__.safeIndex(992, (d_475_v37_).length(0))
                d_470_v32_ = rhs72_
                d_474_v36_ = rhs73_
                lhs59_[lhs60_] = rhs74_
                index69_ = default__.safeIndex(281, (d_417_v2_).length(0))
                (d_417_v2_)[index69_] = (d_417_v2_)[default__.safeIndex(281, (d_417_v2_).length(0))]
                d_478_v38_: C2
                nw64_ = C2()
                nw64_.ctor__((d_474_v36_).f6, (d_473_v35_) | (d_473_v35_), self.f9, (self).f10)
                d_478_v38_ = nw64_
            d_453_cf20_ = (d_454_cf19_ if (_dafny.SeqWithoutIsStrInference([p2, (self).f6])) < (d_413_v0_) else d_453_cf20_)
            d_479_v39_: _dafny.Map
            d_479_v39_ = _dafny.Map({False: p0})
            d_480_v40_: D4
            d_480_v40_ = D4_DC13(d_413_v0_, False, len(d_479_v39_))
            d_480_v40_ = d_480_v40_
            d_481_v41_: D4
            d_481_v41_ = D4_DC14(d_453_cf20_, d_479_v39_, (self).f6, _dafny.Map({True: p0}))
            d_482_v42_: _dafny.Map
            d_482_v42_ = _dafny.Map({d_481_v41_: ((d_417_v2_)[default__.safeIndex(281, (d_417_v2_).length(0))]) < (len(_dafny.Map({d_453_cf20_: d_453_cf20_})))})
            d_482_v42_ = (d_482_v42_).set(d_481_v41_, d_453_cf20_)
        elif True:
            d_483___mcc_h8_ = source7_.cf12
            d_484_cf12_ = d_483___mcc_h8_
            d_485_v43_: D2
            d_485_v43_ = D2_DC8(p0, p0)
            d_486_v44_: _dafny.Map
            d_486_v44_ = _dafny.Map({len((d_413_v0_) + (_dafny.SeqWithoutIsStrInference([(d_417_v2_)[default__.safeIndex(281, (d_417_v2_).length(0))] for d_487_i5_ in range(default__.abs(-880))]))): (d_485_v43_).cf20})
            d_488_v45_: _dafny.Map
            d_488_v45_ = _dafny.Map({(p1)[default__.safeIndex((self).f6, len(p1))]: p0})
            d_486_v44_ = (d_486_v44_).set(len(d_488_v45_), True)
            d_489_v46_: _dafny.Array
            def lambda14_(d_490_i6_):
                return _dafny.Set({(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_491_i7_ in range(default__.abs(-810))])).set(default__.safeIndex(238, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_492_i7_ in range(default__.abs(-810))]))), _dafny.CodePoint('l'))})

            init8_ = lambda14_
            nw65_ = _dafny.Array(None, 27)
            for i0_8_ in range(nw65_.length(0)):
                nw65_[i0_8_] = init8_(i0_8_)
            d_489_v46_ = nw65_
            d_493_v47_: str
            d_493_v47_ = _dafny.CodePoint('b')
            d_494_v48_: _dafny.Seq
            d_494_v48_ = _dafny.SeqWithoutIsStrInference([d_493_v47_, d_493_v47_])
            d_495_v49_: _dafny.Set
            d_495_v49_ = _dafny.Set({d_494_v48_})
            index70_ = default__.safeIndex(599, (d_489_v46_).length(0))
            (d_489_v46_)[index70_] = d_495_v49_
            index71_ = default__.safeIndex(599, (d_489_v46_).length(0))
            (d_489_v46_)[index71_] = d_495_v49_
            if False:
                d_493_v47_ = d_493_v47_
                d_496_v50_: _dafny.MultiSet
                d_496_v50_ = _dafny.MultiSet([_dafny.CodePoint('l'), d_493_v47_])
                d_497_v51_: _dafny.MultiSet
                d_497_v51_ = _dafny.MultiSet([_dafny.MultiSet([619, (d_417_v2_)[default__.safeIndex(281, (d_417_v2_).length(0))]])])
                d_498_v52_: C2
                nw66_ = C2()
                nw66_.ctor__((0) - ((d_496_v50_).cardinality), (default__.fm36(globalState)) - (d_497_v51_), self.f9, _dafny.Map({len(d_494_v48_): _dafny.SeqWithoutIsStrInference([d_493_v47_ for d_499_i8_ in range(default__.abs(670))])}))
                d_498_v52_ = nw66_
                d_500_v53_: _dafny.Map
                d_500_v53_ = _dafny.Map({default__.safeModuloInt(p2, (self).f6): d_417_v2_})
                d_500_v53_ = d_500_v53_
                arr7_ = self.f9
                index72_ = default__.safeIndex(918, (self.f9).length(0))
                arr7_[index72_] = p0
                d_501_v54_: _dafny.Map
                d_501_v54_ = _dafny.Map({p0: d_494_v48_})
                d_502_v55_: _dafny.MultiSet
                d_502_v55_ = _dafny.MultiSet([p0])
                d_503_v56_: _dafny.MultiSet
                d_503_v56_ = _dafny.MultiSet([(self).f6])
                d_504_v57_: _dafny.Set
                d_504_v57_ = _dafny.Set({739, (0) - ((d_417_v2_)[default__.safeIndex(281, (d_417_v2_).length(0))]), (d_417_v2_)[default__.safeIndex(281, (d_417_v2_).length(0))]})
                pat_let_tv14_ = p1
                d_505_v58_: _dafny.Array
                nw67_ = _dafny.Array(None, 16)
                nw67_[int(0)] = d_501_v54_
                nw67_[int(1)] = (d_501_v54_).set(p0, d_494_v48_)
                nw67_[int(2)] = d_501_v54_
                nw67_[int(3)] = _dafny.Map({p0: d_494_v48_})
                nw67_[int(4)] = d_501_v54_
                nw67_[int(5)] = d_501_v54_
                nw67_[int(6)] = (self).fm13(_dafny.CodePoint('s'), d_502_v55_, p0, (d_503_v56_).cardinality, globalState)
                nw67_[int(7)] = d_501_v54_
                nw67_[int(8)] = d_501_v54_
                nw67_[int(9)] = _dafny.Map({p0: d_494_v48_})
                nw67_[int(10)] = d_501_v54_
                nw67_[int(11)] = d_501_v54_
                nw67_[int(12)] = d_501_v54_
                nw67_[int(13)] = d_501_v54_
                nw67_[int(14)] = d_501_v54_
                def iife47_(_pat_let17_0):
                    def iife48_(d_506_dt__update__tmp_h2_):
                        def iife49_(_pat_let18_0):
                            def iife50_(d_507_dt__update_hcf0_h0_):
                                return D0_DC0(d_507_dt__update_hcf0_h0_, (d_506_dt__update__tmp_h2_).cf1)
                            return iife50_(_pat_let18_0)
                        return iife49_(pat_let_tv14_)
                    return iife48_(_pat_let17_0)
                nw67_[int(15)] = _dafny.Map({p0: (D1_DC4(iife47_(D0_DC0(p1, (self).f6)), (self).f6, d_494_v48_, p0, len(d_504_v57_))).cf9})
                d_505_v58_ = nw67_
                index73_ = default__.safeIndex(795, (d_505_v58_).length(0))
                (d_505_v58_)[index73_] = d_501_v54_
                d_508_v59_: D0
                d_508_v59_ = D0_DC1(p0, (self).f6, d_493_v47_)
                pat_let_tv15_ = p0
                arr8_ = self.f9
                index74_ = default__.safeIndex(918, (self.f9).length(0))
                index75_ = default__.safeIndex(795, (d_505_v58_).length(0))
                index76_ = default__.safeIndex(281, (d_417_v2_).length(0))
                def iife51_(_pat_let19_0):
                    def iife52_(d_509_dt__update__tmp_h3_):
                        def iife53_(_pat_let20_0):
                            def iife54_(d_510_dt__update_hcf2_h0_):
                                return D0_DC1(d_510_dt__update_hcf2_h0_, (d_509_dt__update__tmp_h3_).cf3, (d_509_dt__update__tmp_h3_).cf4)
                            return iife54_(_pat_let20_0)
                        return iife53_(pat_let_tv15_)
                    return iife52_(_pat_let19_0)
                rhs75_ = p0
                rhs76_ = d_501_v54_
                rhs77_ = len((d_494_v48_).set(default__.safeIndex((p2) * (p2), len(d_494_v48_)), d_493_v47_))
                rhs78_ = iife51_(d_508_v59_)
                rhs79_ = (p0) in (p1)
                lhs61_ = self.f9
                lhs62_ = default__.safeIndex(918, (self.f9).length(0))
                lhs63_ = d_505_v58_
                lhs64_ = default__.safeIndex(795, (d_505_v58_).length(0))
                lhs65_ = d_417_v2_
                lhs66_ = default__.safeIndex(281, (d_417_v2_).length(0))
                lhs61_[lhs62_] = rhs75_
                lhs63_[lhs64_] = rhs76_
                lhs65_[lhs66_] = rhs77_
                d_508_v59_ = rhs78_
                r1 = rhs79_
                d_511_v60_: _dafny.Set
                d_511_v60_ = _dafny.Set({p0, (self.f9)[default__.safeIndex(918, (self.f9).length(0))]})
                d_512_v61_: D0
                d_512_v61_ = D0_DC0((p1).set(default__.safeIndex(len(d_511_v60_), len(p1)), p0), -92)
                d_513_v62_: D1
                d_513_v62_ = D1_DC4(d_512_v61_, (self).f6, d_494_v48_, p0, p2)
                d_514_v63_: _dafny.Set
                d_514_v63_ = _dafny.Set({((d_486_v44_)[(self).f6] if ((self).f6) in (d_486_v44_) else (self.f9)[default__.safeIndex(918, (self.f9).length(0))]), (d_513_v62_).cf10})
                d_515_v64_: _dafny.Set
                d_515_v64_ = _dafny.Set({d_511_v60_, d_514_v63_, _dafny.Set({((d_488_v45_)[False] if (False) in (d_488_v45_) else (self.f9)[default__.safeIndex(918, (self.f9).length(0))])})})
                d_516_v65_: _dafny.Seq
                d_516_v65_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_514_v63_, d_511_v60_])])
                d_517_v67_: _dafny.Map
                d_517_v67_ = _dafny.Map({d_514_v63_: (self).f6})
                d_518_v69_: _dafny.Map
                d_518_v69_ = _dafny.Map({_dafny.Set({p0, (self.f9)[default__.safeIndex(918, (self.f9).length(0))]}): d_511_v60_})
                d_519_v71_: _dafny.Seq
                d_519_v71_ = _dafny.SeqWithoutIsStrInference([d_515_v64_])
                d_520_v72_: _dafny.Map
                d_520_v72_ = _dafny.Map({d_514_v63_: (self.f9)[default__.safeIndex(918, (self.f9).length(0))]})
                d_521_v74_: _dafny.Array
                nw68_ = _dafny.Array(None, 27)
                nw68_[int(0)] = d_515_v64_
                nw68_[int(1)] = d_515_v64_
                nw68_[int(2)] = d_515_v64_
                nw68_[int(3)] = d_515_v64_
                nw68_[int(4)] = d_515_v64_
                nw68_[int(5)] = (d_515_v64_).intersection(d_515_v64_)
                nw68_[int(6)] = _dafny.Set({d_514_v63_, d_514_v63_, d_511_v60_})
                nw68_[int(7)] = d_515_v64_
                nw68_[int(8)] = (d_515_v64_ if p0 else d_515_v64_)
                def iife55_():
                    coll13_ = _dafny.Set()
                    compr_13_: _dafny.Set
                    for compr_13_ in ((d_516_v65_)[default__.safeIndex(735, len(d_516_v65_))]).Elements:
                        d_522_v66_: _dafny.Set = compr_13_
                        if (d_522_v66_) in ((d_516_v65_)[default__.safeIndex(735, len(d_516_v65_))]):
                            coll13_ = coll13_.union(_dafny.Set([d_522_v66_]))
                    return _dafny.Set(coll13_)
                nw68_[int(9)] = iife55_()
                
                nw68_[int(10)] = (d_515_v64_) | (d_515_v64_)
                def iife56_():
                    coll14_ = _dafny.Set()
                    compr_14_: _dafny.Set
                    for compr_14_ in (d_517_v67_).keys.Elements:
                        d_523_v68_: _dafny.Set = compr_14_
                        if (d_523_v68_) in (d_517_v67_):
                            coll14_ = coll14_.union(_dafny.Set([d_523_v68_]))
                    return _dafny.Set(coll14_)
                nw68_[int(11)] = iife56_()
                
                nw68_[int(12)] = d_515_v64_
                nw68_[int(13)] = _dafny.Set({_dafny.Set({(p1)[default__.safeIndex((d_417_v2_)[default__.safeIndex(281, (d_417_v2_).length(0))], len(p1))], (self.f9)[default__.safeIndex(918, (self.f9).length(0))]}), d_511_v60_})
                nw68_[int(14)] = d_515_v64_
                def iife57_():
                    coll15_ = _dafny.Set()
                    compr_15_: _dafny.Set
                    for compr_15_ in (d_518_v69_).keys.Elements:
                        d_524_v70_: _dafny.Set = compr_15_
                        if (d_524_v70_) in (d_518_v69_):
                            coll15_ = coll15_.union(_dafny.Set([d_524_v70_]))
                    return _dafny.Set(coll15_)
                nw68_[int(15)] = (d_515_v64_) - (iife57_()
                )
                nw68_[int(16)] = (d_519_v71_)[default__.safeIndex((d_502_v55_).cardinality, len(d_519_v71_))]
                def iife58_():
                    coll16_ = _dafny.Set()
                    compr_16_: _dafny.Set
                    for compr_16_ in (d_520_v72_).keys.Elements:
                        d_525_v73_: _dafny.Set = compr_16_
                        if (d_525_v73_) in (d_520_v72_):
                            coll16_ = coll16_.union(_dafny.Set([d_525_v73_]))
                    return _dafny.Set(coll16_)
                nw68_[int(17)] = iife58_()
                
                nw68_[int(18)] = (d_519_v71_)[default__.safeIndex((self).f6, len(d_519_v71_))]
                nw68_[int(19)] = (d_515_v64_).intersection(_dafny.Set({d_511_v60_, d_511_v60_}))
                nw68_[int(20)] = _dafny.Set({d_514_v63_})
                nw68_[int(21)] = (_dafny.Set({d_514_v63_, _dafny.Set({p0}), d_514_v63_, d_511_v60_, d_511_v60_})).intersection(d_515_v64_)
                nw68_[int(22)] = d_515_v64_
                nw68_[int(23)] = d_515_v64_
                nw68_[int(24)] = d_515_v64_
                nw68_[int(25)] = (d_515_v64_ if p0 else d_515_v64_)
                nw68_[int(26)] = d_515_v64_
                d_521_v74_ = nw68_
                index77_ = default__.safeIndex(14, (d_521_v74_).length(0))
                (d_521_v74_)[index77_] = _dafny.Set({d_514_v63_})
                index78_ = default__.safeIndex(14, (d_521_v74_).length(0))
                rhs80_ = (p0) and ((self.f9)[default__.safeIndex(918, (self.f9).length(0))])
                rhs81_ = (self.f9)[default__.safeIndex(918, (self.f9).length(0))]
                rhs82_ = (p0) and ((self.f9)[default__.safeIndex(918, (self.f9).length(0))])
                rhs83_ = default__.fm2(((d_417_v2_)[default__.safeIndex(281, (d_417_v2_).length(0))]) - (p2), globalState)
                rhs84_ = d_515_v64_
                lhs67_ = globalState
                lhs68_ = globalState
                lhs69_ = globalState
                lhs70_ = globalState
                lhs71_ = d_521_v74_
                lhs72_ = default__.safeIndex(14, (d_521_v74_).length(0))
                lhs67_.f0 = rhs80_
                lhs68_.f0 = rhs81_
                lhs69_.f0 = rhs82_
                lhs70_.f5 = rhs83_
                lhs71_[lhs72_] = rhs84_
            elif True:
                index79_ = default__.safeIndex(281, (d_417_v2_).length(0))
                (d_417_v2_)[index79_] = (0) - (((d_417_v2_)[default__.safeIndex(281, (d_417_v2_).length(0))]) - ((0) - ((0) - ((self).f6))))
                d_526_v75_: _dafny.MultiSet
                d_526_v75_ = _dafny.MultiSet([(self).f6])
                d_527_v76_: _dafny.MultiSet
                d_527_v76_ = _dafny.MultiSet([d_526_v75_])
                d_528_v77_: C2
                nw69_ = C2()
                nw69_.ctor__((d_417_v2_)[default__.safeIndex(281, (d_417_v2_).length(0))], (d_527_v76_) - (d_527_v76_), self.f9, _dafny.Map({(self).f6: d_494_v48_}))
                d_528_v77_ = nw69_
                d_529_v78_: _dafny.Set
                d_529_v78_ = _dafny.Set({p0})
                d_530_v79_: D5
                d_530_v79_ = D5_DC16(d_529_v78_)
                d_530_v79_ = d_530_v79_
                (globalState).f5 = (d_528_v77_).fm21(len(p1), globalState)
                d_417_v2_ = d_417_v2_
            (globalState).f5 = (d_417_v2_)[default__.safeIndex(281, (d_417_v2_).length(0))]
        d_531_v80_: D5
        d_531_v80_ = D5_DC17(d_413_v0_, p0)
        source8_ = d_531_v80_
        if source8_.is_DC17:
            d_532___mcc_h9_ = source8_.cf36
            d_533___mcc_h10_ = source8_.cf37
            d_534_cf37_ = d_533___mcc_h10_
            d_535_cf36_ = d_532___mcc_h9_
            d_536_v81_: _dafny.Map
            d_536_v81_ = _dafny.Map({(self).f6: (d_417_v2_)[default__.safeIndex(281, (d_417_v2_).length(0))]})
            d_537_v82_: _dafny.Map
            d_537_v82_ = _dafny.Map({d_534_cf37_: d_534_cf37_})
            d_538_v83_: _dafny.Map
            d_538_v83_ = _dafny.Map({(default__.fm37(d_534_cf37_, globalState)).cf41: ((d_537_v82_)[p0] if (p0) in (d_537_v82_) else d_534_cf37_)})
            d_536_v81_ = (d_536_v81_).set(len((p1) + (p1)), len(d_538_v83_))
            d_539_v84_: C1
            nw70_ = C1()
            nw70_.ctor__(p2)
            d_539_v84_ = nw70_
            d_540_v85_: _dafny.Seq
            d_540_v85_ = _dafny.SeqWithoutIsStrInference([d_539_v84_, d_539_v84_])
            d_541_v86_: _dafny.Map
            d_541_v86_ = _dafny.Map({(d_540_v85_)[default__.safeIndex(len(p1), len(d_540_v85_))]: d_534_cf37_})
            arr9_ = self.f9
            index80_ = default__.safeIndex(102, (self.f9).length(0))
            arr9_[index80_] = ((d_541_v86_)[d_539_v84_] if (d_539_v84_) in (d_541_v86_) else d_534_cf37_)
            d_542_v87_: str
            d_542_v87_ = _dafny.CodePoint('r')
            d_543_v88_: _dafny.MultiSet
            d_543_v88_ = _dafny.MultiSet([d_542_v87_])
            d_544_v89_: _dafny.Map
            d_544_v89_ = _dafny.Map({d_534_cf37_: d_543_v88_})
            d_545_v90_: _dafny.Seq
            d_545_v90_ = _dafny.SeqWithoutIsStrInference([d_543_v88_, _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v'), d_542_v87_, d_542_v87_])), d_543_v88_, ((d_544_v89_)[d_534_cf37_] if (d_534_cf37_) in (d_544_v89_) else (d_543_v88_).set(_dafny.CodePoint('c'), default__.abs(p2)))])
            d_546_v91_: D4
            d_546_v91_ = D4_DC12(d_545_v90_)
            arr10_ = self.f9
            index81_ = default__.safeIndex(102, (self.f9).length(0))
            rhs85_ = False
            rhs86_ = (len(((_dafny.SeqWithoutIsStrInference([p1])).set(default__.safeIndex(237, len(_dafny.SeqWithoutIsStrInference([p1]))), _dafny.SeqWithoutIsStrInference([p0, not(d_534_cf37_), True, p0]))) + (_dafny.SeqWithoutIsStrInference([p1])))) > ((182) - (default__.fm2(565, globalState)))
            rhs87_ = d_546_v91_
            lhs73_ = self.f9
            lhs74_ = default__.safeIndex(102, (self.f9).length(0))
            lhs73_[lhs74_] = rhs85_
            d_534_cf37_ = rhs86_
            d_546_v91_ = rhs87_
            d_547_v92_: _dafny.MultiSet
            d_547_v92_ = _dafny.MultiSet([-930])
            d_548_v93_: _dafny.Array
            nw71_ = _dafny.Array(None, 10)
            d_548_v93_ = nw71_
            d_549_v94_: D0
            d_549_v94_ = D0_DC1(d_534_cf37_, (self).f6, d_542_v87_)
            d_550_v95_: _dafny.Seq
            d_550_v95_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vst"))
            d_551_v96_: _dafny.Array
            def lambda15_(d_552_p1_):
                def lambda16_(d_553_i11_):
                    return D0_DC0(d_552_p1_, (self).f6)

                return lambda16_

            init9_ = lambda15_(p1)
            nw72_ = _dafny.Array(None, 15)
            for i0_9_ in range(nw72_.length(0)):
                nw72_[i0_9_] = init9_(i0_9_)
            d_551_v96_ = nw72_
            d_554_v97_: D7
            d_554_v97_ = D7_DC21(d_551_v96_)
            d_555_v98_: _dafny.Set
            d_555_v98_ = _dafny.Set({(d_554_v97_).cf43})
            rhs88_ = ((d_549_v94_).cf2 if d_534_cf37_ else ((self.f9)[default__.safeIndex(102, (self.f9).length(0))]) == (d_534_cf37_))
            rhs89_ = default__.fm35((d_417_v2_)[default__.safeIndex(281, (d_417_v2_).length(0))], ((self).f6) * ((0) - ((self).f6)), False, p2, globalState)
            rhs90_ = d_417_v2_
            rhs91_ = len((_dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_556_i10_ in range(default__.abs(921))])).set(default__.safeIndex((self).f6, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_557_i10_ in range(default__.abs(921))]))), d_542_v87_) for d_558_i9_ in range(default__.abs(928))])).set(default__.safeIndex(-96, len(_dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_559_i10_ in range(default__.abs(921))])).set(default__.safeIndex((self).f6, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_560_i10_ in range(default__.abs(921))]))), d_542_v87_) for d_561_i9_ in range(default__.abs(928))]))), d_550_v95_))
            rhs92_ = (d_548_v93_ if (d_555_v98_) != (d_555_v98_) else d_548_v93_)
            lhs75_ = globalState
            d_534_cf37_ = rhs88_
            d_547_v92_ = rhs89_
            d_417_v2_ = rhs90_
            lhs75_.f5 = rhs91_
            d_548_v93_ = rhs92_
            arr11_ = self.f9
            index82_ = default__.safeIndex(102, (self.f9).length(0))
            arr11_[index82_] = p0
        elif True:
            d_562___mcc_h11_ = source8_.cf35
            d_563_cf35_ = d_562___mcc_h11_
            d_564_v99_: _dafny.Map
            d_564_v99_ = _dafny.Map({(_dafny.MultiSet(p1)).cardinality: _dafny.SeqWithoutIsStrInference([(self).f6 for d_565_i14_ in range(default__.abs(416))])})
            d_566_v100_: _dafny.MultiSet
            d_566_v100_ = _dafny.MultiSet([p0, True])
            d_567_v101_: str
            d_567_v101_ = _dafny.CodePoint('h')
            d_568_v102_: _dafny.MultiSet
            d_568_v102_ = _dafny.MultiSet([_dafny.CodePoint('u'), d_567_v101_])
            d_569_v103_: _dafny.Map
            d_569_v103_ = _dafny.Map({not(p0): (d_413_v0_).set(default__.safeIndex(782, len(d_413_v0_)), ((d_566_v100_)[p0] if (p0) in (d_566_v100_) else (d_568_v102_).cardinality))})
            d_570_v104_: _dafny.MultiSet
            d_570_v104_ = _dafny.MultiSet([(self).f6, (d_417_v2_)[default__.safeIndex(281, (d_417_v2_).length(0))]])
            d_571_v105_: _dafny.Seq
            d_571_v105_ = _dafny.SeqWithoutIsStrInference([d_413_v0_])
            d_572_v106_: _dafny.Seq
            d_572_v106_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nhw"))
            d_573_v107_: _dafny.Seq
            d_573_v107_ = _dafny.SeqWithoutIsStrInference([d_413_v0_, d_413_v0_, (d_571_v105_)[default__.safeIndex(default__.fm2(len(d_572_v106_), globalState), len(d_571_v105_))]])
            d_574_v108_: _dafny.Array
            nw73_ = _dafny.Array(None, 14)
            nw73_[int(0)] = (_dafny.SeqWithoutIsStrInference([(d_417_v2_)[default__.safeIndex(281, (d_417_v2_).length(0))] for d_575_i12_ in range(default__.abs(49))])) + (d_413_v0_)
            nw73_[int(1)] = d_413_v0_
            nw73_[int(2)] = (d_413_v0_) + (_dafny.SeqWithoutIsStrInference([p2 for d_576_i13_ in range(default__.abs(332))]))
            nw73_[int(3)] = ((d_564_v99_)[len(d_413_v0_)] if (len(d_413_v0_)) in (d_564_v99_) else _dafny.SeqWithoutIsStrInference([p2, (self).f6]))
            nw73_[int(4)] = (d_413_v0_).set(default__.safeIndex((d_417_v2_)[default__.safeIndex(281, (d_417_v2_).length(0))], len(d_413_v0_)), default__.fm2((self).f6, globalState))
            nw73_[int(5)] = (d_413_v0_) + (_dafny.SeqWithoutIsStrInference([(self).f6 for d_577_i15_ in range(default__.abs(640))]))
            nw73_[int(6)] = d_413_v0_
            nw73_[int(7)] = d_413_v0_
            nw73_[int(8)] = ((d_569_v103_)[p0] if (p0) in (d_569_v103_) else (d_413_v0_).set(default__.safeIndex((d_570_v104_).cardinality, len(d_413_v0_)), (d_417_v2_)[default__.safeIndex(281, (d_417_v2_).length(0))]))
            nw73_[int(9)] = d_413_v0_
            nw73_[int(10)] = (d_413_v0_) + (d_413_v0_)
            nw73_[int(11)] = (d_571_v105_)[default__.safeIndex((self).f6, len(d_571_v105_))]
            nw73_[int(12)] = (d_413_v0_ if not(not(not(p0))) else d_413_v0_)
            nw73_[int(13)] = (_dafny.SeqWithoutIsStrInference([(d_417_v2_)[default__.safeIndex(281, (d_417_v2_).length(0))], (self).f6, (d_417_v2_)[default__.safeIndex(281, (d_417_v2_).length(0))], 953, p2])) + (_dafny.SeqWithoutIsStrInference([(d_417_v2_)[default__.safeIndex(281, (d_417_v2_).length(0))]]))
            d_574_v108_ = nw73_
            def lambda17_(d_578_v0_, d_579_p1_):
                def lambda18_(d_580_i16_):
                    return (d_578_v0_) + ((d_578_v0_).set(default__.safeIndex((self).f6, len(d_578_v0_)), len(d_579_p1_)))

                return lambda18_

            init10_ = lambda17_(d_413_v0_, p1)
            nw74_ = _dafny.Array(None, 28)
            for i0_10_ in range(nw74_.length(0)):
                nw74_[i0_10_] = init10_(i0_10_)
            d_574_v108_ = nw74_
            d_581_v109_: D5
            d_581_v109_ = D5_DC16(d_563_cf35_)
            d_582_v110_: D5
            d_582_v110_ = D5_DC16((d_581_v109_).cf35)
            pat_let_tv16_ = globalState
            pat_let_tv17_ = d_563_cf35_
            def iife59_(_pat_let21_0):
                def iife60_(d_583_dt__update__tmp_h4_):
                    def iife61_(_pat_let22_0):
                        def iife62_(d_584_dt__update_hcf35_h0_):
                            return D5_DC16(d_584_dt__update_hcf35_h0_)
                        return iife62_(_pat_let22_0)
                    return iife61_(((default__.fm38(pat_let_tv16_)).cf35) | (pat_let_tv17_))
                return iife60_(_pat_let21_0)
            rhs93_ = p2
            rhs94_ = iife59_(d_582_v110_)
            lhs76_ = globalState
            lhs76_.f5 = rhs93_
            d_581_v109_ = rhs94_
            (globalState).f5 = ((self).f6 if p0 else p2)
            d_585_v111_: _dafny.Seq
            d_585_v111_ = _dafny.SeqWithoutIsStrInference([d_572_v106_, d_572_v106_])
            (globalState).f5 = default__.safeDivisionInt(len((d_585_v111_) + (_dafny.SeqWithoutIsStrInference([d_572_v106_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f"))]))), len(d_572_v106_))
        d_586_v112_: _dafny.Set
        d_586_v112_ = _dafny.Set({-87})
        r0 = default__.fm29(d_586_v112_, globalState)
        r1 = p0
        return r0, r1

    def m7(self, p0, p1, globalState):
        r0: str = _dafny.CodePoint('D')
        r1: int = int(0)
        r2: _dafny.Array = _dafny.Array(None, 0)
        r3: int = int(0)
        d_587_v0_: str
        d_587_v0_ = _dafny.CodePoint('j')
        r0 = d_587_v0_
        (self).f9 = ((self.f9))
        d_588_v1_: _dafny.Array
        def lambda19_(d_589_i0_):
            return (d_589_i0_) + ((self).f6)

        init11_ = lambda19_
        nw75_ = _dafny.Array(None, 9)
        for i0_11_ in range(nw75_.length(0)):
            nw75_[i0_11_] = init11_(i0_11_)
        d_588_v1_ = nw75_
        index83_ = default__.safeIndex(211, (d_588_v1_).length(0))
        (d_588_v1_)[index83_] = (self).f6
        d_590_v2_: bool
        d_590_v2_ = False
        index84_ = default__.safeIndex(211, (d_588_v1_).length(0))
        rhs95_ = -180
        rhs96_ = 84
        rhs97_ = d_590_v2_
        lhs77_ = d_588_v1_
        lhs78_ = default__.safeIndex(211, (d_588_v1_).length(0))
        lhs79_ = globalState
        r3 = rhs95_
        lhs77_[lhs78_] = rhs96_
        lhs79_.f0 = rhs97_
        (globalState).f0 = d_590_v2_
        d_591_v3_: _dafny.Map
        d_591_v3_ = _dafny.Map({d_590_v2_: p1})
        d_592_v5_: _dafny.Seq
        d_592_v5_ = _dafny.SeqWithoutIsStrInference([(d_588_v1_)[default__.safeIndex(211, (d_588_v1_).length(0))], p1, (d_588_v1_)[default__.safeIndex(211, (d_588_v1_).length(0))]])
        def iife63_():
            coll17_ = _dafny.Map()
            compr_17_: _dafny.Seq
            for compr_17_ in (_dafny.Map({d_592_v5_: d_590_v2_})).keys.Elements:
                d_593_v4_: _dafny.Seq = compr_17_
                if (d_593_v4_) in (_dafny.Map({d_592_v5_: d_590_v2_})):
                    coll17_[d_593_v4_] = d_590_v2_
            return _dafny.Map(coll17_)
        if (len(default__.fm1(globalState))) <= (len((_dafny.Set({d_591_v3_, _dafny.Map({not(d_590_v2_): len(iife63_()
        )})})) | (_dafny.Set({d_591_v3_, d_591_v3_})))):
            d_594_v6_: C0
            nw76_ = C0()
            nw76_.ctor__(self.f9, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gxtawrg"))) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qi"))))
            d_594_v6_ = nw76_
            (globalState).f0 = default__.fm3(globalState)
            d_595_v7_: _dafny.Seq
            d_595_v7_ = _dafny.SeqWithoutIsStrInference([d_594_v6_.f20])
            d_596_v8_: D0
            d_596_v8_ = D0_DC0(d_595_v7_, (d_588_v1_)[default__.safeIndex(211, (d_588_v1_).length(0))])
            d_597_v9_: _dafny.Seq
            d_597_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fsmbn"))
            d_598_v10_: D1
            d_598_v10_ = D1_DC4(d_596_v8_, p1, d_597_v9_, d_590_v2_, (self).f6)
            d_599_v11_: int
            d_600_v12_: str
            d_601_v13_: bool
            out19_: int
            out20_: str
            out21_: bool
            out19_, out20_, out21_ = (self).m14((d_598_v10_).cf10, d_594_v6_.f20, globalState)
            d_599_v11_ = out19_
            d_600_v12_ = out20_
            d_601_v13_ = out21_
            d_602_v14_: _dafny.Array
            nw77_ = _dafny.Array(_dafny.Seq({}), 18)
            d_602_v14_ = nw77_
            d_603_v15_: _dafny.Map
            d_603_v15_ = _dafny.Map({d_602_v14_: d_590_v2_})
            d_603_v15_ = (d_603_v15_).set(d_602_v14_, default__.fm3(globalState))
            d_604_v16_: _dafny.MultiSet
            d_604_v16_ = _dafny.MultiSet([d_588_v1_])
            (d_594_v6_).f20 = (d_604_v16_).ispropersubset((_dafny.MultiSet([d_588_v1_])).intersection(d_604_v16_))
        elif True:
            d_605_v17_: _dafny.Array
            nw78_ = _dafny.Array(D2.default()(), 13)
            d_605_v17_ = nw78_
            nw79_ = _dafny.Array(D2.default()(), 25)
            d_605_v17_ = nw79_
            d_606_v18_: _dafny.MultiSet
            d_606_v18_ = _dafny.MultiSet([(p0).set((d_588_v1_)[default__.safeIndex(211, (d_588_v1_).length(0))], default__.abs((self).f6))])
            d_607_v19_: C2
            nw80_ = C2()
            nw80_.ctor__((d_588_v1_)[default__.safeIndex(211, (d_588_v1_).length(0))], (d_606_v18_).set(p0, default__.abs((d_588_v1_)[default__.safeIndex(211, (d_588_v1_).length(0))])), self.f9, default__.fm39(d_587_v0_, globalState))
            d_607_v19_ = nw80_
            d_608_v20_: _dafny.Set
            d_608_v20_ = _dafny.Set({(d_588_v1_)[default__.safeIndex(211, (d_588_v1_).length(0))], (self).f6, (p1) + (default__.fm2((self).f6, globalState))})
            d_608_v20_ = d_608_v20_
            d_609_v21_: _dafny.Array
            nw81_ = _dafny.Array(_dafny.Array(None, 0), 24)
            d_609_v21_ = nw81_
            d_610_v22_: _dafny.Array
            nw82_ = _dafny.Array(_dafny.Set({}), 2)
            d_610_v22_ = nw82_
            index85_ = default__.safeIndex(554, (d_609_v21_).length(0))
            (d_609_v21_)[index85_] = d_610_v22_
            index86_ = default__.safeIndex(554, (d_609_v21_).length(0))
            (d_609_v21_)[index86_] = d_610_v22_
            if not (d_590_v2_) or (d_590_v2_):
                d_611_v23_: int
                d_612_v24_: str
                d_613_v25_: bool
                out22_: int
                out23_: str
                out24_: bool
                out22_, out23_, out24_ = (self).m14(True, d_590_v2_, globalState)
                d_611_v23_ = out22_
                d_612_v24_ = out23_
                d_613_v25_ = out24_
                d_614_v26_: _dafny.Array
                d_614_v26_ = self.f9
                d_614_v26_ = d_614_v26_
                d_615_v27_: _dafny.Map
                d_615_v27_ = _dafny.Map({p0: p1})
                d_590_v2_ = (d_611_v23_) != (((d_615_v27_)[p0] if (p0) in (d_615_v27_) else (self).f6))
                d_616_v28_: _dafny.Seq
                d_616_v28_ = _dafny.SeqWithoutIsStrInference([d_613_v25_])
                d_617_v29_: _dafny.Map
                d_617_v29_ = _dafny.Map({d_613_v25_: d_613_v25_})
                d_618_v30_: _dafny.Map
                d_618_v30_ = _dafny.Map({not((d_616_v28_)[default__.safeIndex(255, len(d_616_v28_))]): ((d_617_v29_)[d_613_v25_] if (d_613_v25_) in (d_617_v29_) else True)})
                r1 = ((0) - (((d_591_v3_)[d_590_v2_] if (d_590_v2_) in (d_591_v3_) else len(d_618_v30_)))) + (116)
                d_608_v20_ = d_608_v20_
            elif True:
                d_619_v31_: D9
                d_619_v31_ = D9_DC24(d_606_v18_)
                d_620_v32_: C2
                nw83_ = C2()
                nw83_.ctor__(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ytxvws"))), (d_606_v18_) | ((d_619_v31_).cf45), self.f9, (self).f10)
                d_620_v32_ = nw83_
                d_621_v33_: _dafny.Array
                nw84_ = _dafny.Array(_dafny.CodePoint('D'), 7)
                d_621_v33_ = nw84_
                d_621_v33_ = d_621_v33_
                d_588_v1_ = d_588_v1_
                d_622_v34_: _dafny.Set
                d_622_v34_ = _dafny.Set({d_590_v2_, d_590_v2_, True})
                (globalState).f5 = len(d_622_v34_)
                d_623_v35_: C0
                nw85_ = C0()
                nw85_.ctor__(self.f9, not(d_590_v2_))
                d_623_v35_ = nw85_
        d_624_v36_: _dafny.Array
        nw86_ = _dafny.Array(None, 6)
        nw86_[int(0)] = (self).f10
        nw86_[int(1)] = (self).f10
        nw86_[int(2)] = ((self).f10) | ((self).f10)
        nw86_[int(3)] = (self).f10
        nw86_[int(4)] = (self).f10
        nw86_[int(5)] = (self).f10
        d_624_v36_ = nw86_
        d_625_v37_: _dafny.Seq
        d_625_v37_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rtg"))
        index87_ = default__.safeIndex(273, (d_624_v36_).length(0))
        (d_624_v36_)[index87_] = _dafny.Map({(d_588_v1_)[default__.safeIndex(211, (d_588_v1_).length(0))]: d_625_v37_})
        index88_ = default__.safeIndex(273, (d_624_v36_).length(0))
        (d_624_v36_)[index88_] = (self).f10
        r0 = d_587_v0_
        d_626_v38_: _dafny.Seq
        d_626_v38_ = _dafny.SeqWithoutIsStrInference([d_590_v2_, default__.fm3(globalState)])
        r1 = ((d_588_v1_)[default__.safeIndex(211, (d_588_v1_).length(0))]) - (len(d_626_v38_))
        nw87_ = _dafny.Array(None, 7)
        nw87_[int(0)] = d_590_v2_
        nw87_[int(1)] = False
        nw87_[int(2)] = (d_626_v38_) <= ((d_626_v38_).set(default__.safeIndex((self).f6, len(d_626_v38_)), default__.fm3(globalState)))
        nw87_[int(3)] = (-757) == (len(d_592_v5_))
        nw87_[int(4)] = (d_590_v2_ if d_590_v2_ else d_590_v2_)
        nw87_[int(5)] = d_590_v2_
        nw87_[int(6)] = not (d_590_v2_) or (True)
        r2 = nw87_
        r3 = (0) - (default__.safeDivisionInt(len(d_625_v37_), (self).f6))
        return r0, r1, r2, r3

    def m14(self, p0, p1, globalState):
        r0: int = int(0)
        r1: str = _dafny.CodePoint('D')
        r2: bool = False
        if p1:
            r0 = (self).f6
            (globalState).f5 = (self).f6
            d_627_v0_: C0
            nw88_ = C0()
            nw88_.ctor__(self.f9, p0)
            d_627_v0_ = nw88_
            d_628_v1_: _dafny.Seq
            d_628_v1_ = _dafny.SeqWithoutIsStrInference([((self).f6) - ((self).f6), (self).f6])
            d_628_v1_ = (d_628_v1_) + (d_628_v1_)
            d_629_v2_: _dafny.Map
            d_629_v2_ = _dafny.Map({(self).f6: p0})
            d_629_v2_ = (d_629_v2_).set(-864, d_627_v0_.f20)
        elif True:
            rhs98_ = p0
            rhs99_ = (self).f6
            lhs80_ = globalState
            lhs81_ = globalState
            lhs80_.f0 = rhs98_
            lhs81_.f5 = rhs99_
            arr12_ = self.f9
            index89_ = default__.safeIndex(199, (self.f9).length(0))
            arr12_[index89_] = p0
            arr13_ = self.f9
            index90_ = default__.safeIndex(199, (self.f9).length(0))
            arr13_[index90_] = p0
            (globalState).f5 = ((self).f6 if p1 else (self).f6)
            if not(p0):
                d_630_v3_: _dafny.Seq
                d_630_v3_ = _dafny.SeqWithoutIsStrInference([p1])
                d_631_v4_: D0
                d_631_v4_ = D0_DC0(d_630_v3_, (self).f6)
                d_632_v5_: str
                d_632_v5_ = _dafny.CodePoint('v')
                d_633_v6_: D1
                d_633_v6_ = D1_DC4(d_631_v4_, (self).f6, default__.fm30(p0, d_632_v5_, globalState), (self.f9)[default__.safeIndex(199, (self.f9).length(0))], (self).f6)
                (globalState).f2 = (d_633_v6_).cf9
                d_634_v7_: _dafny.Set
                d_634_v7_ = _dafny.Set({d_632_v5_, d_632_v5_, d_632_v5_})
                d_635_v8_: _dafny.MultiSet
                d_635_v8_ = _dafny.MultiSet([d_634_v7_, d_634_v7_, d_634_v7_])
                d_636_v9_: _dafny.Set
                d_636_v9_ = _dafny.Set({(self.f9)[default__.safeIndex(199, (self.f9).length(0))], p0})
                d_637_v10_: _dafny.Map
                d_637_v10_ = _dafny.Map({len(d_636_v9_): p0})
                d_638_v11_: _dafny.Map
                d_638_v11_ = _dafny.Map({d_637_v10_: p1})
                d_639_v12_: D3
                d_639_v12_ = D3_DC10(d_638_v11_)
                d_640_v13_: _dafny.Seq
                d_640_v13_ = _dafny.SeqWithoutIsStrInference([len(_dafny.Map({p0: len(default__.fm33(d_639_v12_, (self).f6, False, p1, globalState))}))])
                d_641_v14_: _dafny.Seq
                d_641_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b"))
                d_642_v15_: _dafny.Array
                nw89_ = _dafny.Array(None, 22)
                nw89_[int(0)] = (self).f6
                nw89_[int(1)] = (self).f6
                nw89_[int(2)] = (0) - ((self).f6)
                nw89_[int(3)] = (self).f6
                nw89_[int(4)] = (self).f6
                nw89_[int(5)] = (self).f6
                nw89_[int(6)] = (self).f6
                nw89_[int(7)] = ((self).f6) - ((self).f6)
                nw89_[int(8)] = 247
                nw89_[int(9)] = (self).f6
                nw89_[int(10)] = ((self).f6) * (558)
                nw89_[int(11)] = ((self).f6) - (default__.fm2((0) - ((self).f6), globalState))
                nw89_[int(12)] = (self).f6
                nw89_[int(13)] = (self).f6
                nw89_[int(14)] = (self).f6
                nw89_[int(15)] = (self).f6
                nw89_[int(16)] = (d_635_v8_).cardinality
                nw89_[int(17)] = ((self).f6) + (len(d_640_v13_))
                nw89_[int(18)] = (self).f6
                nw89_[int(19)] = len((d_641_v14_) + (d_641_v14_))
                nw89_[int(20)] = (self).f6
                nw89_[int(21)] = (self).f6
                d_642_v15_ = nw89_
                index91_ = default__.safeIndex(13, (d_642_v15_).length(0))
                (d_642_v15_)[index91_] = (self).f6
                index92_ = default__.safeIndex(13, (d_642_v15_).length(0))
                (d_642_v15_)[index92_] = len(_dafny.SeqWithoutIsStrInference([d_632_v5_]))
                d_643_v16_: _dafny.Array
                nw90_ = _dafny.Array(_dafny.Map({}), 3)
                d_643_v16_ = nw90_
                d_644_v17_: _dafny.Seq
                d_644_v17_ = _dafny.SeqWithoutIsStrInference([d_643_v16_, d_643_v16_])
                d_645_v18_: D9
                d_645_v18_ = D9_DC26((d_644_v17_)[default__.safeIndex((self).f6, len(d_644_v17_))])
                index93_ = default__.safeIndex(13, (d_642_v15_).length(0))
                arr14_ = self.f9
                index94_ = default__.safeIndex(199, (self.f9).length(0))
                rhs100_ = d_645_v18_
                rhs101_ = default__.safeModuloInt((d_642_v15_)[default__.safeIndex(13, (d_642_v15_).length(0))], len((d_640_v13_) + (d_640_v13_)))
                rhs102_ = p0
                lhs82_ = d_642_v15_
                lhs83_ = default__.safeIndex(13, (d_642_v15_).length(0))
                lhs84_ = self.f9
                lhs85_ = default__.safeIndex(199, (self.f9).length(0))
                d_645_v18_ = rhs100_
                lhs82_[lhs83_] = rhs101_
                lhs84_[lhs85_] = rhs102_
                d_646_v19_: _dafny.Set
                d_646_v19_ = _dafny.Set({default__.safeDivisionInt((d_642_v15_)[default__.safeIndex(13, (d_642_v15_).length(0))], len(d_641_v14_))})
                def iife64_():
                    coll18_ = _dafny.Set()
                    compr_18_: int
                    for compr_18_ in (d_640_v13_).Elements:
                        d_647_v20_: int = compr_18_
                        if (d_647_v20_) in (d_640_v13_):
                            coll18_ = coll18_.union(_dafny.Set([default__.safeDivisionInt(d_647_v20_, (len(_dafny.Map({_dafny.Set({-986}): D4_DC12(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([_dafny.CodePoint('n'), _dafny.CodePoint('h')]), _dafny.MultiSet([_dafny.CodePoint('c'), _dafny.CodePoint('m')])]))}))) * ((_dafny.MultiSet([(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({not(not(not(False)))})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ehssnh"))), len(_dafny.Map({True: -954}))]))).cardinality])).cardinality))]))
                    return _dafny.Set(coll18_)
                d_646_v19_ = iife64_()
                
                d_648_v21_: _dafny.Array
                nw91_ = _dafny.Array(None, 28)
                nw91_[int(0)] = d_632_v5_
                nw91_[int(1)] = d_632_v5_
                nw91_[int(2)] = d_632_v5_
                nw91_[int(3)] = d_632_v5_
                nw91_[int(4)] = _dafny.CodePoint('j')
                nw91_[int(5)] = d_632_v5_
                nw91_[int(6)] = d_632_v5_
                nw91_[int(7)] = d_632_v5_
                nw91_[int(8)] = d_632_v5_
                nw91_[int(9)] = d_632_v5_
                nw91_[int(10)] = _dafny.CodePoint('p')
                nw91_[int(11)] = d_632_v5_
                nw91_[int(12)] = d_632_v5_
                nw91_[int(13)] = (D0_DC1(p1, (self).f6, d_632_v5_)).cf4
                nw91_[int(14)] = d_632_v5_
                nw91_[int(15)] = d_632_v5_
                nw91_[int(16)] = d_632_v5_
                nw91_[int(17)] = d_632_v5_
                nw91_[int(18)] = d_632_v5_
                nw91_[int(19)] = d_632_v5_
                nw91_[int(20)] = d_632_v5_
                nw91_[int(21)] = d_632_v5_
                nw91_[int(22)] = d_632_v5_
                nw91_[int(23)] = d_632_v5_
                nw91_[int(24)] = d_632_v5_
                nw91_[int(25)] = d_632_v5_
                nw91_[int(26)] = _dafny.CodePoint('r')
                nw91_[int(27)] = d_632_v5_
                d_648_v21_ = nw91_
                index95_ = default__.safeIndex(121, (d_648_v21_).length(0))
                (d_648_v21_)[index95_] = d_632_v5_
                d_649_v22_: _dafny.Map
                d_649_v22_ = _dafny.Map({not((self.f9)[default__.safeIndex(199, (self.f9).length(0))]): p1})
                index96_ = default__.safeIndex(121, (d_648_v21_).length(0))
                (d_648_v21_)[index96_] = default__.fm0(len(d_630_v3_), _dafny.Set({(self).f6, len(d_630_v3_)}), ((d_649_v22_)[p0] if (p0) in (d_649_v22_) else (self.f9)[default__.safeIndex(199, (self.f9).length(0))]), globalState)
            elif True:
                d_650_v23_: _dafny.MultiSet
                d_650_v23_ = _dafny.MultiSet([(self).f6, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "idv")))])
                (globalState).f5 = (((self).f6) - ((d_650_v23_).cardinality)) + (default__.safeDivisionInt((self).f6, (self).f6))
                d_651_v24_: _dafny.Map
                d_651_v24_ = _dafny.Map({p0: (self).f6})
                d_651_v24_ = (d_651_v24_).set(p0, (self).f6)
                d_652_v25_: _dafny.Array
                def lambda20_(d_653_i0_):
                    return _dafny.SeqWithoutIsStrInference([(self).f6, (self).f6])

                init12_ = lambda20_
                nw92_ = _dafny.Array(None, 13)
                for i0_12_ in range(nw92_.length(0)):
                    nw92_[i0_12_] = init12_(i0_12_)
                d_652_v25_ = nw92_
                d_652_v25_ = d_652_v25_
                d_654_v26_: _dafny.Array
                nw93_ = _dafny.Array(int(0), 9)
                d_654_v26_ = nw93_
                index97_ = default__.safeIndex(339, (d_654_v26_).length(0))
                (d_654_v26_)[index97_] = (self).f6
                index98_ = default__.safeIndex(339, (d_654_v26_).length(0))
                rhs103_ = p1
                rhs104_ = (self).f6
                lhs86_ = globalState
                lhs87_ = d_654_v26_
                lhs88_ = default__.safeIndex(339, (d_654_v26_).length(0))
                lhs86_.f0 = rhs103_
                lhs87_[lhs88_] = rhs104_
                d_655_v27_: _dafny.Seq
                d_655_v27_ = _dafny.SeqWithoutIsStrInference([(self.f9)[default__.safeIndex(199, (self.f9).length(0))], default__.fm3(globalState), p1])
                d_656_v28_: _dafny.Array
                nw94_ = _dafny.Array(False, 17)
                d_656_v28_ = nw94_
                d_657_v29_: _dafny.Map
                d_657_v29_ = _dafny.Map({(self).f6: d_656_v28_})
                d_658_v30_: _dafny.Set
                d_658_v30_ = _dafny.Set({default__.safeDivisionInt((d_654_v26_)[default__.safeIndex(339, (d_654_v26_).length(0))], (d_654_v26_)[default__.safeIndex(339, (d_654_v26_).length(0))]), (263) * (len(d_657_v29_))})
                d_659_v31_: _dafny.Seq
                d_659_v31_ = _dafny.SeqWithoutIsStrInference([(self).f6])
                d_660_v32_: D2
                d_660_v32_ = D2_DC6((d_654_v26_)[default__.safeIndex(339, (d_654_v26_).length(0))], d_654_v26_)
                rhs105_ = d_654_v26_
                rhs106_ = len((d_659_v31_) + ((d_659_v31_) + (d_659_v31_)))
                rhs107_ = (d_655_v27_) + ((d_655_v27_).set(default__.safeIndex((self).f6, len(d_655_v27_)), (self.f9)[default__.safeIndex(199, (self.f9).length(0))]))
                rhs108_ = (d_658_v30_).intersection(_dafny.Set({len((self).f21)}))
                rhs109_ = (d_660_v32_).cf13
                lhs89_ = globalState
                d_654_v26_ = rhs105_
                r0 = rhs106_
                d_655_v27_ = rhs107_
                d_658_v30_ = rhs108_
                lhs89_.f5 = rhs109_
            d_661_v33_: _dafny.Map
            d_661_v33_ = _dafny.Map({(self).f6: True})
            d_662_v34_: _dafny.Seq
            d_662_v34_ = _dafny.SeqWithoutIsStrInference([(self).f6, (self).f6, len(d_661_v33_)])
            d_663_v35_: _dafny.Set
            d_663_v35_ = _dafny.Set({(0) - ((self).f6), (self).f6, (0) - ((self).f6)})
            d_664_v36_: _dafny.Seq
            d_664_v36_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jfb"))
            d_665_v37_: _dafny.Map
            d_665_v37_ = _dafny.Map({default__.fm3(globalState): p0})
            d_666_v38_: D4
            d_666_v38_ = D4_DC14(p1, _dafny.Map({p1: (self.f9)[default__.safeIndex(199, (self.f9).length(0))]}), default__.fm2((self).f6, globalState), d_665_v37_)
            d_667_v39_: str
            d_667_v39_ = _dafny.CodePoint('i')
            d_668_v40_: _dafny.MultiSet
            d_668_v40_ = _dafny.MultiSet([(self).f6])
            d_669_v41_: _dafny.MultiSet
            d_669_v41_ = _dafny.MultiSet([(d_668_v40_).cardinality, 877])
            d_670_v42_: D0
            d_670_v42_ = D0_DC1((self.f9)[default__.safeIndex(199, (self.f9).length(0))], (self).f6, d_667_v39_)
            d_671_v43_: _dafny.Seq
            d_671_v43_ = _dafny.SeqWithoutIsStrInference([d_664_v36_])
            d_672_v44_: _dafny.Array
            nw95_ = _dafny.Array(None, 24)
            nw95_[int(0)] = -382
            nw95_[int(1)] = (self).f6
            nw95_[int(2)] = (self).f6
            nw95_[int(3)] = (self).f6
            nw95_[int(4)] = len(d_664_v36_)
            nw95_[int(5)] = (0) - ((0) - ((0) - ((_dafny.MultiSet([default__.fm0((self).f6, d_663_v35_, p0, globalState), d_667_v39_])).cardinality)))
            nw95_[int(6)] = len(d_664_v36_)
            nw95_[int(7)] = (self).f6
            nw95_[int(8)] = (self).f6
            nw95_[int(9)] = len((_dafny.SeqWithoutIsStrInference([d_667_v39_])).set(default__.safeIndex((self).f6, len(_dafny.SeqWithoutIsStrInference([d_667_v39_]))), _dafny.CodePoint('m')))
            nw95_[int(10)] = len(d_662_v34_)
            nw95_[int(11)] = 850
            nw95_[int(12)] = 216
            nw95_[int(13)] = (d_668_v40_).cardinality
            nw95_[int(14)] = -913
            nw95_[int(15)] = (0) - ((self).f6)
            nw95_[int(16)] = (self).f6
            nw95_[int(17)] = (self).f6
            nw95_[int(18)] = (d_670_v42_).cf3
            nw95_[int(19)] = (self).f6
            nw95_[int(20)] = 227
            nw95_[int(21)] = (0) - ((self).f6)
            nw95_[int(22)] = ((d_669_v41_)[(self).f6] if ((self).f6) in (d_669_v41_) else len((d_671_v43_)[default__.safeIndex((self).f6, len(d_671_v43_))]))
            nw95_[int(23)] = (self).f6
            d_672_v44_ = nw95_
            d_673_v45_: D2
            d_673_v45_ = D2_DC6(181, d_672_v44_)
            d_674_v46_: _dafny.Array
            nw96_ = _dafny.Array(None, 27)
            nw96_[int(0)] = ((self).f6) + ((self).f6)
            nw96_[int(1)] = (0) - (-882)
            nw96_[int(2)] = -131
            nw96_[int(3)] = (self).f6
            nw96_[int(4)] = (self).f6
            nw96_[int(5)] = ((self).f6) + (-829)
            nw96_[int(6)] = 21
            nw96_[int(7)] = (self).f6
            nw96_[int(8)] = default__.safeDivisionInt(len(d_662_v34_), len(d_663_v35_))
            nw96_[int(9)] = (0) - (len(d_664_v36_))
            nw96_[int(10)] = (self).f6
            nw96_[int(11)] = (self).f6
            nw96_[int(12)] = (self).f6
            nw96_[int(13)] = (0) - ((d_666_v38_).cf30)
            nw96_[int(14)] = (self).f6
            nw96_[int(15)] = len(d_664_v36_)
            nw96_[int(16)] = (self).f6
            nw96_[int(17)] = (self).f6
            nw96_[int(18)] = (self).f6
            nw96_[int(19)] = (self).f6
            nw96_[int(20)] = (self).f6
            nw96_[int(21)] = (self).f6
            nw96_[int(22)] = (self).f6
            nw96_[int(23)] = (self).f6
            nw96_[int(24)] = len((default__.fm30(p1, d_667_v39_, globalState)).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_667_v39_ for d_675_i1_ in range(default__.abs(543))])), len(default__.fm30(p1, d_667_v39_, globalState))), d_667_v39_))
            nw96_[int(25)] = (self).f6
            nw96_[int(26)] = (d_673_v45_).cf13
            d_674_v46_ = nw96_
            d_676_v47_: _dafny.Map
            d_676_v47_ = _dafny.Map({(self).f6: _dafny.Set({default__.fm2(len(d_664_v36_), globalState), (self).f6, (self).f6})})
            d_677_v48_: _dafny.Array
            nw97_ = _dafny.Array(None, 23)
            nw97_[int(0)] = D2_DC6((self).f6, d_672_v44_)
            nw97_[int(1)] = d_673_v45_
            nw97_[int(2)] = d_673_v45_
            nw97_[int(3)] = d_673_v45_
            nw97_[int(4)] = d_673_v45_
            nw97_[int(5)] = d_673_v45_
            nw97_[int(6)] = d_673_v45_
            nw97_[int(7)] = d_673_v45_
            nw97_[int(8)] = D2_DC6((self).f6, d_672_v44_)
            nw97_[int(9)] = d_673_v45_
            nw97_[int(10)] = d_673_v45_
            nw97_[int(11)] = d_673_v45_
            nw97_[int(12)] = d_673_v45_
            nw97_[int(13)] = d_673_v45_
            nw97_[int(14)] = d_673_v45_
            nw97_[int(15)] = d_673_v45_
            nw97_[int(16)] = D2_DC6(len(((d_676_v47_)[(self).f6] if ((self).f6) in (d_676_v47_) else d_663_v35_)), d_674_v46_)
            nw97_[int(17)] = d_673_v45_
            nw97_[int(18)] = d_673_v45_
            nw97_[int(19)] = d_673_v45_
            nw97_[int(20)] = d_673_v45_
            nw97_[int(21)] = d_673_v45_
            nw97_[int(22)] = d_673_v45_
            d_677_v48_ = nw97_
            index99_ = default__.safeIndex(947, (d_677_v48_).length(0))
            (d_677_v48_)[index99_] = d_673_v45_
            index100_ = default__.safeIndex(947, (d_677_v48_).length(0))
            rhs110_ = ((0) - ((self).f6) if p0 else (self).f6)
            rhs111_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i"))
            rhs112_ = (self).f6
            rhs113_ = d_672_v44_
            rhs114_ = d_673_v45_
            lhs90_ = globalState
            lhs91_ = globalState
            lhs92_ = globalState
            lhs93_ = d_677_v48_
            lhs94_ = default__.safeIndex(947, (d_677_v48_).length(0))
            lhs90_.f5 = rhs110_
            lhs91_.f2 = rhs111_
            lhs92_.f5 = rhs112_
            d_674_v46_ = rhs113_
            lhs93_[lhs94_] = rhs114_
        d_678_v49_: _dafny.Array
        nw98_ = _dafny.Array(int(0), 7)
        d_678_v49_ = nw98_
        d_679_v50_: _dafny.Map
        d_679_v50_ = _dafny.Map({d_678_v49_: p1})
        d_680_v51_: D10
        d_680_v51_ = D10_DC28(d_679_v50_)
        d_679_v50_ = (d_680_v51_).cf52
        d_681_v52_: _dafny.Set
        d_681_v52_ = _dafny.Set({(self).f6, (self).f6})
        index101_ = default__.safeIndex(450, (d_678_v49_).length(0))
        (d_678_v49_)[index101_] = len(d_681_v52_)
        index102_ = default__.safeIndex(450, (d_678_v49_).length(0))
        (d_678_v49_)[index102_] = (self).f6
        d_682_v53_: _dafny.Seq
        d_682_v53_ = _dafny.SeqWithoutIsStrInference([409])
        d_683_v54_: _dafny.MultiSet
        d_683_v54_ = _dafny.MultiSet([(_dafny.MultiSet([665, 647])).cardinality, (self).f6])
        d_684_v55_: _dafny.MultiSet
        d_684_v55_ = _dafny.MultiSet([d_683_v54_])
        d_685_v56_: T0
        nw99_ = C2()
        nw99_.ctor__((0) - ((d_682_v53_)[default__.safeIndex((self).f6, len(d_682_v53_))]), d_684_v55_, self.f9, (self).f10)
        d_685_v56_ = nw99_
        d_686_v57_: _dafny.Map
        d_686_v57_ = _dafny.Map({d_685_v56_: p0})
        d_687_v58_: _dafny.Map
        d_687_v58_ = _dafny.Map({d_686_v57_: d_682_v53_})
        d_688_v59_: _dafny.Map
        d_688_v59_ = _dafny.Map({(self).f6: d_687_v58_})
        d_688_v59_ = (d_688_v59_).set((d_678_v49_)[default__.safeIndex(450, (d_678_v49_).length(0))], d_687_v58_)
        d_689_i2_: int
        d_689_i2_ = 0
        with _dafny.label("2"):
            while p0:
                with _dafny.c_label("2"):
                    if (d_689_i2_) >= (100):
                        raise _dafny.Break("2")
                    d_689_i2_ = (d_689_i2_) + (1)
                    r2 = p1
                    r0 = 839
                    arr15_ = self.f9
                    index103_ = default__.safeIndex(28, (self.f9).length(0))
                    arr15_[index103_] = p1
                    arr16_ = self.f9
                    index104_ = default__.safeIndex(28, (self.f9).length(0))
                    arr16_[index104_] = (p1 if (p0) or (p0) else True)
                    d_690_v60_: _dafny.Array
                    nw100_ = _dafny.Array(None, 6)
                    nw100_[int(0)] = d_678_v49_
                    nw100_[int(1)] = d_678_v49_
                    nw100_[int(2)] = d_678_v49_
                    nw100_[int(3)] = d_678_v49_
                    nw100_[int(4)] = d_678_v49_
                    nw100_[int(5)] = d_678_v49_
                    d_690_v60_ = nw100_
                    index105_ = default__.safeIndex(762, (d_690_v60_).length(0))
                    (d_690_v60_)[index105_] = d_678_v49_
                    d_691_v61_: _dafny.Seq
                    d_691_v61_ = _dafny.SeqWithoutIsStrInference([d_682_v53_, _dafny.SeqWithoutIsStrInference([(d_685_v56_).f6 for d_692_i3_ in range(default__.abs(640))]), d_682_v53_])
                    d_693_v62_: _dafny.Array
                    def lambda21_(d_694_p0_):
                        def lambda22_(d_695_i4_):
                            return d_694_p0_

                        return lambda22_

                    init13_ = lambda21_(p0)
                    nw101_ = _dafny.Array(None, 9)
                    for i0_13_ in range(nw101_.length(0)):
                        nw101_[i0_13_] = init13_(i0_13_)
                    d_693_v62_ = nw101_
                    d_696_v63_: _dafny.Map
                    d_696_v63_ = _dafny.Map({(d_691_v61_)[default__.safeIndex((self).f6, len(d_691_v61_))]: d_693_v62_})
                    d_697_v64_: _dafny.Set
                    d_697_v64_ = _dafny.Set({(self.f9)[default__.safeIndex(28, (self.f9).length(0))], True})
                    index106_ = default__.safeIndex(450, (d_678_v49_).length(0))
                    index107_ = default__.safeIndex(762, (d_690_v60_).length(0))
                    rhs115_ = (default__.safeModuloInt((self).f6, len(default__.fm40((self).f6, d_697_v64_, globalState)))) * ((0) - ((d_685_v56_).f6))
                    rhs116_ = (d_678_v49_)[default__.safeIndex(450, (d_678_v49_).length(0))]
                    rhs117_ = d_678_v49_
                    rhs118_ = ((d_696_v63_) | (d_696_v63_)) | (d_696_v63_)
                    rhs119_ = p1
                    lhs95_ = d_678_v49_
                    lhs96_ = default__.safeIndex(450, (d_678_v49_).length(0))
                    lhs97_ = d_690_v60_
                    lhs98_ = default__.safeIndex(762, (d_690_v60_).length(0))
                    lhs99_ = globalState
                    r0 = rhs115_
                    lhs95_[lhs96_] = rhs116_
                    lhs97_[lhs98_] = rhs117_
                    d_696_v63_ = rhs118_
                    lhs99_.f0 = rhs119_
                    pass
            pass
        rhs120_ = p0
        rhs121_ = (d_678_v49_ if p0 else d_678_v49_)
        r2 = rhs120_
        d_678_v49_ = rhs121_
        r0 = default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wwrq"))), (d_678_v49_)[default__.safeIndex(450, (d_678_v49_).length(0))])
        d_698_v65_: str
        d_698_v65_ = _dafny.CodePoint('y')
        r1 = d_698_v65_
        d_699_v66_: _dafny.Set
        d_699_v66_ = _dafny.Set({p0, p1, True})
        r2 = (d_699_v66_).ispropersubset(_dafny.Set({p0}))
        return r0, r1, r2

    @property
    def f21(self):
        return self._f21

class C4(T1, T0, T2):
    def  __init__(self):
        self._f9: _dafny.Array = _dafny.Array(None, 0)
        self._f16: _dafny.MultiSet = _dafny.MultiSet({})
        self._f6: int = int(0)
        self._f10: _dafny.Map = _dafny.Map({})
        self._f17: int = int(0)
        self._f18: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    @property
    def f9(self):
        return self._f9
    @f9.setter
    def f9(self, value):
        self._f9 = value
    @property
    def f16(self):
        return self._f16
    @f16.setter
    def f16(self, value):
        self._f16 = value
    @property
    def f6(self):
        return self._f6
    @property
    def f10(self):
        return self._f10
    def ctor__(self, f17, f18, f9, f10, f6, f16):
        (self)._f17 = f17
        (self)._f18 = f18
        (self).f9 = f9
        (self)._f10 = f10
        (self)._f6 = f6
        (self).f16 = f16

    def fm13(self, p0, p1, p2, p3, globalState):
        source9_ = D1_DC3(_dafny.MultiSet([_dafny.CodePoint('n'), _dafny.CodePoint('o'), _dafny.CodePoint('p'), _dafny.CodePoint('p')]))
        if source9_.is_DC4:
            d_700___mcc_h0_ = source9_.cf7
            d_701___mcc_h1_ = source9_.cf8
            d_702___mcc_h2_ = source9_.cf9
            d_703___mcc_h3_ = source9_.cf10
            d_704___mcc_h4_ = source9_.cf11
            d_705_cf11_ = d_704___mcc_h4_
            d_706_cf10_ = d_703___mcc_h3_
            d_707_cf9_ = d_702___mcc_h2_
            d_708_cf8_ = d_701___mcc_h1_
            d_709_cf7_ = d_700___mcc_h0_
            return (_dafny.Map({False: d_707_cf9_})) | (_dafny.Map({d_706_cf10_: d_707_cf9_}))
        elif True:
            d_710___mcc_h5_ = source9_.cf6
            d_711_cf6_ = d_710___mcc_h5_
            return (_dafny.Map({True: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kg"))})) | (_dafny.Map({False: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xmpjtp"))}))

    def fm5(self, p0, globalState):
        source10_ = D0_DC0(_dafny.SeqWithoutIsStrInference([False, True]), (self).f6)
        if source10_.is_DC0:
            d_712___mcc_h0_ = source10_.cf0
            d_713___mcc_h1_ = source10_.cf1
            d_714_cf1_ = d_713___mcc_h1_
            d_715_cf0_ = d_712___mcc_h0_
            return _dafny.SeqWithoutIsStrInference([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a')])), _dafny.MultiSet([_dafny.CodePoint('w'), _dafny.CodePoint('f')]), _dafny.MultiSet([_dafny.CodePoint('d')]), _dafny.MultiSet([_dafny.CodePoint('u')]), _dafny.MultiSet([_dafny.CodePoint('h'), _dafny.CodePoint('d')])])
        elif source10_.is_DC1:
            d_716___mcc_h2_ = source10_.cf2
            d_717___mcc_h3_ = source10_.cf3
            d_718___mcc_h4_ = source10_.cf4
            d_719_cf4_ = d_718___mcc_h4_
            d_720_cf3_ = d_717___mcc_h3_
            d_721_cf2_ = d_716___mcc_h2_
            return (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_719_cf4_]), _dafny.MultiSet([d_719_cf4_])])) + (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_719_cf4_]) for d_722_i0_ in range(default__.abs(448))]))
        elif True:
            d_723___mcc_h5_ = source10_.cf5
            d_724_cf5_ = d_723___mcc_h5_
            return (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([_dafny.CodePoint('e'), _dafny.CodePoint('g')]) for d_725_i1_ in range(default__.abs(409))])) + (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_726_i3_ in range(default__.abs(237))])) for d_727_i2_ in range(default__.abs(613))]))

    def fm6(self, p0, globalState):
        return (_dafny.Map({(self).f17: _dafny.SeqWithoutIsStrInference([(self).f18])})) | ((_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_728_i0_ in range(default__.abs(669))])): _dafny.SeqWithoutIsStrInference([(self).f18])})) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([False, True])): _dafny.SeqWithoutIsStrInference([(self).f18])})))

    def fm21(self, p0, globalState):
        return default__.safeDivisionInt((self).f18, ((self).f17) - (len(_dafny.Map({True: not(False)}))))

    def fm22(self, p0, p1, p2, globalState):
        source11_ = (D0_DC1(False, (self).f6, _dafny.CodePoint('y')) if not(False) else D0_DC1(True, (self).f18, _dafny.CodePoint('n')))
        if source11_.is_DC0:
            d_729___mcc_h0_ = source11_.cf0
            d_730___mcc_h1_ = source11_.cf1
            d_731_cf1_ = d_730___mcc_h1_
            d_732_cf0_ = d_729___mcc_h0_
            return not(True)
        elif source11_.is_DC1:
            d_733___mcc_h2_ = source11_.cf2
            d_734___mcc_h3_ = source11_.cf3
            d_735___mcc_h4_ = source11_.cf4
            d_736_cf4_ = d_735___mcc_h4_
            d_737_cf3_ = d_734___mcc_h3_
            d_738_cf2_ = d_733___mcc_h2_
            return d_738_cf2_
        elif True:
            d_739___mcc_h5_ = source11_.cf5
            d_740_cf5_ = d_739___mcc_h5_
            return (_dafny.Set({False})).ispropersubset(_dafny.Set({True}))

    def m6(self, p0, p1, p2, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: bool = False
        d_741_v0_: _dafny.Map
        d_741_v0_ = _dafny.Map({18: p0})
        hi5_ = (p2) - (p2)
        for d_742_i0_ in range(len((d_741_v0_) | (d_741_v0_)), hi5_):
            (globalState).f5 = p2
            d_743_v1_: D0
            d_743_v1_ = D0_DC0(p1, (self).f18)
            d_744_v2_: _dafny.Seq
            d_744_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jnbsg"))
            d_745_v3_: D1
            d_745_v3_ = D1_DC4(d_743_v1_, (self).f6, (d_744_v2_ if p0 else d_744_v2_), p0, (self).f18)
            d_745_v3_ = D1_DC4(d_743_v1_, 403, d_744_v2_, p0, (self).f18)
            d_744_v2_ = ((d_744_v2_) + (d_744_v2_)) + ((d_744_v2_) + (d_744_v2_))
            d_746_v4_: _dafny.Array
            def lambda23_(d_747_i1_):
                return D1_DC3(_dafny.MultiSet([_dafny.CodePoint('g')]))

            init14_ = lambda23_
            nw102_ = _dafny.Array(None, 21)
            for i0_14_ in range(nw102_.length(0)):
                nw102_[i0_14_] = init14_(i0_14_)
            d_746_v4_ = nw102_
            d_748_v5_: _dafny.Array
            nw103_ = _dafny.Array(_dafny.MultiSet({}), 26)
            d_748_v5_ = nw103_
            d_749_v6_: _dafny.Map
            d_749_v6_ = _dafny.Map({d_746_v4_: d_748_v5_})
            d_750_v7_: _dafny.Map
            d_750_v7_ = _dafny.Map({p0: (p1)[default__.safeIndex(-465, len(p1))]})
            d_751_v8_: _dafny.MultiSet
            d_751_v8_ = _dafny.MultiSet([p0, ((d_750_v7_)[not(p0)] if (not(p0)) in (d_750_v7_) else p0)])
            d_752_v9_: D2
            d_752_v9_ = D2_DC7((p1) + (default__.fm4(d_742_i0_, (0) - (d_742_i0_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m"))), globalState)), ((d_749_v6_)[d_746_v4_] if (d_746_v4_) in (d_749_v6_) else d_748_v5_), default__.fm2((self).f17, globalState), ((self).f18) + ((d_751_v8_).cardinality))
            d_752_v9_ = d_752_v9_
        (globalState).f5 = (self).f18
        d_753_i2_: int
        d_753_i2_ = 0
        with _dafny.label("3"):
            while p0:
                with _dafny.c_label("3"):
                    if (d_753_i2_) >= (100):
                        raise _dafny.Break("3")
                    d_753_i2_ = (d_753_i2_) + (1)
                    d_754_v10_: C0
                    nw104_ = C0()
                    nw104_.ctor__(self.f9, default__.fm3(globalState))
                    d_754_v10_ = nw104_
                    (globalState).f5 = (self).f18
                    d_755_v11_: D0
                    d_755_v11_ = D0_DC0(_dafny.SeqWithoutIsStrInference([p0]), (self).f6)
                    (globalState).f5 = default__.safeDivisionInt(len(default__.fm25(p2, d_754_v10_.f20, globalState)), (d_755_v11_).cf1)
                    d_756_v12_: _dafny.Array
                    def lambda24_(d_757_i3_):
                        return default__.safeDivisionInt(d_757_i3_, (self).f18)

                    init15_ = lambda24_
                    nw105_ = _dafny.Array(None, 9)
                    for i0_15_ in range(nw105_.length(0)):
                        nw105_[i0_15_] = init15_(i0_15_)
                    d_756_v12_ = nw105_
                    index108_ = default__.safeIndex(602, (d_756_v12_).length(0))
                    (d_756_v12_)[index108_] = (self).f6
                    d_758_v13_: _dafny.Seq
                    d_758_v13_ = _dafny.SeqWithoutIsStrInference([8, (self).f17])
                    d_759_v14_: D2
                    d_759_v14_ = D2_DC8(d_754_v10_.f20, p0)
                    d_760_v15_: _dafny.MultiSet
                    d_760_v15_ = _dafny.MultiSet([(self).f6])
                    d_761_v16_: str
                    d_761_v16_ = _dafny.CodePoint('l')
                    d_762_v17_: _dafny.Map
                    d_762_v17_ = _dafny.Map({d_761_v16_: (self).f6})
                    index109_ = default__.safeIndex(602, (d_756_v12_).length(0))
                    rhs122_ = d_754_v10_.f20
                    rhs123_ = (self).f6
                    rhs124_ = (d_758_v13_) <= (d_758_v13_)
                    rhs125_ = len(_dafny.Map({d_754_v10_.f20: d_759_v14_}))
                    rhs126_ = default__.safeDivisionInt(len(default__.fm4((0) - (((d_760_v15_).set((self).f6, default__.abs((0) - ((self).f18)))).cardinality), 621, (self).f18, globalState)), default__.safeModuloInt(((d_762_v17_)[_dafny.CodePoint('g')] if (_dafny.CodePoint('g')) in (d_762_v17_) else (self).f6), (self).f6))
                    lhs100_ = d_756_v12_
                    lhs101_ = default__.safeIndex(602, (d_756_v12_).length(0))
                    lhs102_ = globalState
                    lhs103_ = globalState
                    lhs104_ = globalState
                    r1 = rhs122_
                    lhs100_[lhs101_] = rhs123_
                    lhs102_.f0 = rhs124_
                    lhs103_.f5 = rhs125_
                    lhs104_.f5 = rhs126_
                    pass
            pass
        if p0:
            d_763_v18_: D0
            d_763_v18_ = D0_DC0(p1, p2)
            d_764_v19_: D0
            d_764_v19_ = D0_DC2(d_763_v18_)
            d_765_v20_: _dafny.Map
            d_765_v20_ = _dafny.Map({(self).f18: d_764_v19_})
            d_766_v21_: _dafny.Set
            d_766_v21_ = _dafny.Set({d_765_v20_})
            def iife65_():
                coll19_ = _dafny.Set()
                compr_19_: _dafny.Map
                for compr_19_ in (_dafny.SeqWithoutIsStrInference([d_765_v20_ for d_767_i4_ in range(default__.abs(-885))])).Elements:
                    d_768_v22_: _dafny.Map = compr_19_
                    if (d_768_v22_) in (_dafny.SeqWithoutIsStrInference([d_765_v20_ for d_767_i4_ in range(default__.abs(-885))])):
                        coll19_ = coll19_.union(_dafny.Set([d_768_v22_]))
                return _dafny.Set(coll19_)
            (globalState).f0 = not(((iife65_()
            ) | (d_766_v21_)).ispropersubset(d_766_v21_))
            d_769_v23_: str
            d_769_v23_ = _dafny.CodePoint('f')
            d_769_v23_ = d_769_v23_
            d_770_v24_: C0
            nw106_ = C0()
            nw106_.ctor__(self.f9, p0)
            d_770_v24_ = nw106_
            d_770_v24_ = d_770_v24_
            d_771_v25_: _dafny.Array
            nw107_ = _dafny.Array(_dafny.CodePoint('D'), 2)
            d_771_v25_ = nw107_
            index110_ = default__.safeIndex(656, (d_771_v25_).length(0))
            (d_771_v25_)[index110_] = d_769_v23_
            index111_ = default__.safeIndex(656, (d_771_v25_).length(0))
            (d_771_v25_)[index111_] = d_769_v23_
            d_772_v26_: _dafny.Seq
            d_772_v26_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wjkiaxif"))
            d_773_v27_: _dafny.MultiSet
            d_773_v27_ = _dafny.MultiSet([len(d_772_v26_)])
            (globalState).f0 = ((self).fm21((0) - (len(d_741_v0_)), globalState)) < (default__.safeModuloInt((self).f17, (d_773_v27_).cardinality))
        elif True:
            if ((self).f18) != (p2):
                d_774_v28_: _dafny.Array
                def lambda25_(d_775_i5_):
                    return (d_775_i5_) - ((self).f6)

                init16_ = lambda25_
                nw108_ = _dafny.Array(None, 17)
                for i0_16_ in range(nw108_.length(0)):
                    nw108_[i0_16_] = init16_(i0_16_)
                d_774_v28_ = nw108_
                d_776_v29_: _dafny.Map
                d_776_v29_ = _dafny.Map({d_774_v28_: len((p1) + (p1))})
                d_776_v29_ = d_776_v29_
                d_777_v30_: C0
                nw109_ = C0()
                nw109_.ctor__(self.f9, p0)
                d_777_v30_ = nw109_
                d_778_v31_: C0
                nw110_ = C0()
                nw110_.ctor__(self.f9, p0)
                d_778_v31_ = nw110_
                d_779_v32_: _dafny.Seq
                d_779_v32_ = _dafny.SeqWithoutIsStrInference([d_778_v31_.f20])
                rhs127_ = p1
                rhs128_ = True
                lhs105_ = d_778_v31_
                d_779_v32_ = rhs127_
                lhs105_.f20 = rhs128_
                d_780_v33_: C0
                nw111_ = C0()
                nw111_.ctor__(self.f9, p0)
                d_780_v33_ = nw111_
            elif True:
                (globalState).f0 = (_dafny.Set({p0})).issubset((_dafny.Set({p0, p0, False})) - (_dafny.Set({p0, p0})))
                (globalState).f5 = (self).f17
                (globalState).f0 = p0
                d_781_v34_: C0
                nw112_ = C0()
                nw112_.ctor__((self.f9 if p0 else self.f9), p0)
                d_781_v34_ = nw112_
                d_782_v35_: _dafny.Array
                nw113_ = _dafny.Array(int(0), 28)
                d_782_v35_ = nw113_
                d_783_v36_: _dafny.Seq
                d_783_v36_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "koabsqd"))
                d_784_v37_: str
                d_784_v37_ = _dafny.CodePoint('v')
                index112_ = default__.safeIndex(788, (d_782_v35_).length(0))
                (d_782_v35_)[index112_] = (len(d_783_v36_)) * (len(default__.fm26((self).f6, d_784_v37_, d_781_v34_.f20, d_784_v37_, globalState)))
                index113_ = default__.safeIndex(788, (d_782_v35_).length(0))
                (d_782_v35_)[index113_] = (self).f6
            d_785_v38_: _dafny.Seq
            d_785_v38_ = _dafny.SeqWithoutIsStrInference([(self).f17])
            d_786_v39_: C0
            nw114_ = C0()
            nw114_.ctor__(self.f9, ((self).f17) > ((d_785_v38_)[default__.safeIndex((self).f18, len(d_785_v38_))]))
            d_786_v39_ = nw114_
            d_787_v40_: _dafny.Seq
            d_787_v40_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "snhg"))
            d_788_v41_: _dafny.Set
            d_788_v41_ = _dafny.Set({d_786_v39_.f20})
            d_789_v42_: str
            d_789_v42_ = _dafny.CodePoint('b')
            d_790_v43_: _dafny.Map
            d_790_v43_ = _dafny.Map({(_dafny.MultiSet([len(d_788_v41_)])).cardinality: d_789_v42_})
            d_791_v45_: _dafny.MultiSet
            d_791_v45_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_789_v42_ for d_792_i6_ in range(default__.abs(18))])])
            d_793_v46_: _dafny.Map
            d_793_v46_ = _dafny.Map({d_787_v40_: -128})
            d_794_v48_: _dafny.Array
            nw115_ = _dafny.Array(None, 13)
            def iife66_():
                coll20_ = _dafny.Map()
                compr_20_: _dafny.Seq
                for compr_20_ in (d_791_v45_).Elements:
                    d_795_v44_: _dafny.Seq = compr_20_
                    if (d_795_v44_) in (d_791_v45_):
                        coll20_[d_795_v44_] = 313
                return _dafny.Map(coll20_)
            nw115_[int(0)] = (_dafny.Map({d_787_v40_: len(_dafny.SeqWithoutIsStrInference([(d_787_v40_)[default__.safeIndex(p2, len(d_787_v40_))], ((d_790_v43_)[(self).f17] if ((self).f17) in (d_790_v43_) else (d_787_v40_)[default__.safeIndex(p2, len(d_787_v40_))])]))}) if p0 else iife66_()
            )
            nw115_[int(1)] = d_793_v46_
            nw115_[int(2)] = d_793_v46_
            nw115_[int(3)] = d_793_v46_
            nw115_[int(4)] = d_793_v46_
            nw115_[int(5)] = default__.fm27(globalState)
            nw115_[int(6)] = d_793_v46_
            nw115_[int(7)] = d_793_v46_
            nw115_[int(8)] = d_793_v46_
            nw115_[int(9)] = d_793_v46_
            nw115_[int(10)] = (d_793_v46_) | (_dafny.Map({d_787_v40_: (0) - (-528)}))
            nw115_[int(11)] = d_793_v46_
            def iife67_():
                coll21_ = _dafny.Map()
                compr_21_: _dafny.Seq
                for compr_21_ in ((_dafny.MultiSet([d_787_v40_])).set(d_787_v40_, default__.abs((self).f17))).Elements:
                    d_796_v47_: _dafny.Seq = compr_21_
                    if (d_796_v47_) in ((_dafny.MultiSet([d_787_v40_])).set(d_787_v40_, default__.abs((self).f17))):
                        coll21_[d_796_v47_] = (self).f17
                return _dafny.Map(coll21_)
            nw115_[int(12)] = (iife67_()
            ) | (d_793_v46_)
            d_794_v48_ = nw115_
            index114_ = default__.safeIndex(179, (d_794_v48_).length(0))
            (d_794_v48_)[index114_] = d_793_v46_
            index115_ = default__.safeIndex(179, (d_794_v48_).length(0))
            (d_794_v48_)[index115_] = d_793_v46_
            d_797_v49_: _dafny.Set
            d_797_v49_ = _dafny.Set({(self).f17, (0) - (p2), (0) - ((self).f6)})
            d_798_v50_: _dafny.MultiSet
            d_798_v50_ = _dafny.MultiSet([(self).f18, p2, len(d_797_v49_)])
            rhs129_ = (d_798_v50_ if not((p0) == (p0)) else d_798_v50_)
            rhs130_ = d_786_v39_.f20
            lhs106_ = d_786_v39_
            d_798_v50_ = rhs129_
            lhs106_.f20 = rhs130_
            arr17_ = d_786_v39_.f19
            index116_ = default__.safeIndex(313, (d_786_v39_.f19).length(0))
            arr17_[index116_] = p0
            arr18_ = d_786_v39_.f19
            index117_ = default__.safeIndex(313, (d_786_v39_.f19).length(0))
            arr18_[index117_] = (d_798_v50_).issubset((d_798_v50_) | (d_798_v50_))
        d_799_i7_: int
        d_799_i7_ = 0
        with _dafny.label("4"):
            while p0:
                with _dafny.c_label("4"):
                    if (d_799_i7_) >= (100):
                        raise _dafny.Break("4")
                    d_799_i7_ = (d_799_i7_) + (1)
                    d_800_v51_: _dafny.MultiSet
                    d_800_v51_ = _dafny.MultiSet([len(_dafny.Map({self.f9: (self).f18}))])
                    d_741_v0_ = (d_741_v0_).set(default__.safeModuloInt(p2, (d_800_v51_).cardinality), True)
                    arr19_ = self.f9
                    index118_ = default__.safeIndex(921, (self.f9).length(0))
                    arr19_[index118_] = p0
                    d_801_v52_: _dafny.Array
                    def lambda26_(d_802_i8_):
                        return (d_802_i8_) - ((self).f17)

                    init17_ = lambda26_
                    nw116_ = _dafny.Array(None, 8)
                    for i0_17_ in range(nw116_.length(0)):
                        nw116_[i0_17_] = init17_(i0_17_)
                    d_801_v52_ = nw116_
                    index119_ = default__.safeIndex(522, (d_801_v52_).length(0))
                    (d_801_v52_)[index119_] = p2
                    d_803_v53_: _dafny.Seq
                    d_803_v53_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wmae"))
                    arr20_ = self.f9
                    index120_ = default__.safeIndex(921, (self.f9).length(0))
                    index121_ = default__.safeIndex(522, (d_801_v52_).length(0))
                    rhs131_ = p0
                    rhs132_ = (self).f6
                    rhs133_ = p0
                    rhs134_ = (len(d_803_v53_)) + ((self).f17)
                    lhs107_ = globalState
                    lhs108_ = globalState
                    lhs109_ = self.f9
                    lhs110_ = default__.safeIndex(921, (self.f9).length(0))
                    lhs111_ = d_801_v52_
                    lhs112_ = default__.safeIndex(522, (d_801_v52_).length(0))
                    lhs107_.f0 = rhs131_
                    lhs108_.f5 = rhs132_
                    lhs109_[lhs110_] = rhs133_
                    lhs111_[lhs112_] = rhs134_
                    (globalState).f5 = p2
                    d_804_v54_: _dafny.MultiSet
                    d_804_v54_ = _dafny.MultiSet([d_801_v52_])
                    rhs135_ = ((d_801_v52_ if p0 else d_801_v52_) if (self.f9)[default__.safeIndex(921, (self.f9).length(0))] else d_801_v52_)
                    rhs136_ = d_804_v54_
                    rhs137_ = _dafny.SeqWithoutIsStrInference([(d_803_v53_)[default__.safeIndex((self).f17, len(d_803_v53_))] for d_805_i9_ in range(default__.abs(751))])
                    rhs138_ = default__.fm3(globalState)
                    lhs113_ = globalState
                    lhs114_ = globalState
                    d_801_v52_ = rhs135_
                    d_804_v54_ = rhs136_
                    lhs113_.f2 = rhs137_
                    lhs114_.f0 = rhs138_
                    pass
            pass
        d_806_v55_: _dafny.Array
        nw117_ = _dafny.Array(_dafny.Set({}), 21)
        d_806_v55_ = nw117_
        index122_ = default__.safeIndex(275, (d_806_v55_).length(0))
        (d_806_v55_)[index122_] = _dafny.Set({(self).f6, (self).f17})
        d_807_v56_: _dafny.Map
        d_807_v56_ = _dafny.Map({(0) - (p2): (self).f17})
        d_808_v57_: D2
        d_808_v57_ = D2_DC5(d_807_v56_)
        d_809_v58_: T1
        nw118_ = C3()
        nw118_.ctor__((d_808_v57_).cf12, self.f9, (self).f10, 13)
        d_809_v58_ = nw118_
        d_810_v59_: _dafny.Seq
        d_810_v59_ = _dafny.SeqWithoutIsStrInference([d_809_v58_])
        d_811_v60_: _dafny.Map
        d_811_v60_ = _dafny.Map({p0: p0})
        d_812_v61_: _dafny.Seq
        d_812_v61_ = _dafny.SeqWithoutIsStrInference([d_811_v60_, d_811_v60_])
        d_813_v62_: _dafny.Set
        d_813_v62_ = _dafny.Set({p2, (_dafny.MultiSet(d_810_v59_)).cardinality, len(d_812_v61_)})
        d_814_v63_: D11
        d_814_v63_ = D11_DC30(d_813_v62_)
        d_815_v64_: _dafny.Seq
        d_815_v64_ = _dafny.SeqWithoutIsStrInference([(d_813_v62_) - ((d_814_v63_).cf57)])
        index123_ = default__.safeIndex(275, (d_806_v55_).length(0))
        rhs139_ = ((self).f18) - ((self).f18)
        rhs140_ = self.f9
        rhs141_ = (d_815_v64_)[default__.safeIndex(len(_dafny.Map({p0: False})), len(d_815_v64_))]
        lhs115_ = globalState
        lhs116_ = self
        lhs117_ = d_806_v55_
        lhs118_ = default__.safeIndex(275, (d_806_v55_).length(0))
        lhs115_.f5 = rhs139_
        lhs116_.f9 = rhs140_
        lhs117_[lhs118_] = rhs141_
        r0 = default__.fm29(((d_806_v55_)[default__.safeIndex(275, (d_806_v55_).length(0))] if p0 else (d_806_v55_)[default__.safeIndex(275, (d_806_v55_).length(0))]), globalState)
        r1 = p0
        return r0, r1

    def m7(self, p0, p1, globalState):
        r0: str = _dafny.CodePoint('D')
        r1: int = int(0)
        r2: _dafny.Array = _dafny.Array(None, 0)
        r3: int = int(0)
        d_816_v0_: str
        d_816_v0_ = _dafny.CodePoint('v')
        d_817_v1_: _dafny.MultiSet
        d_817_v1_ = _dafny.MultiSet([False])
        d_818_v2_: _dafny.Seq
        d_818_v2_ = _dafny.SeqWithoutIsStrInference([(self).f6])
        d_819_v3_: _dafny.MultiSet
        d_819_v3_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([(self).f17, (d_817_v1_).cardinality]), d_818_v2_])
        d_820_v5_: _dafny.Map
        def iife68_():
            coll22_ = _dafny.Set()
            compr_22_: _dafny.Seq
            for compr_22_ in (d_819_v3_).Elements:
                d_821_v4_: _dafny.Seq = compr_22_
                if (d_821_v4_) in (d_819_v3_):
                    coll22_ = coll22_.union(_dafny.Set([d_821_v4_]))
            return _dafny.Set(coll22_)
        d_820_v5_ = _dafny.Map({d_816_v0_: len(iife68_()
        )})
        def iife69_():
            coll23_ = _dafny.Map()
            compr_23_: int
            for compr_23_ in _dafny.IntegerRange(35, 958):
                d_822_v6_: int = compr_23_
                if ((35) <= (d_822_v6_)) and ((d_822_v6_) < (958)):
                    coll23_[(d_822_v6_) - ((self).f18)] = True
            return _dafny.Map(coll23_)
        hi6_ = ((d_820_v5_)[d_816_v0_] if (d_816_v0_) in (d_820_v5_) else len(iife69_()
        ))
        for d_823_i0_ in range((self).f6, hi6_):
            d_824_v7_: _dafny.Seq
            d_824_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rrvwuy"))
            d_825_v8_: bool
            d_825_v8_ = False
            d_826_v9_: _dafny.Map
            d_826_v9_ = _dafny.Map({d_825_v8_: p1})
            d_827_v10_: C3
            nw119_ = C3()
            nw119_.ctor__(_dafny.Map({d_823_i0_: (self).f6}), self.f9, (self).f10, d_823_i0_)
            d_827_v10_ = nw119_
            d_828_v11_: _dafny.Map
            d_828_v11_ = _dafny.Map({d_827_v10_: (self).f17})
            d_829_v12_: _dafny.Array
            nw120_ = _dafny.Array(None, 10)
            nw120_[int(0)] = p1
            nw120_[int(1)] = -173
            nw120_[int(2)] = len(d_824_v7_)
            nw120_[int(3)] = len(d_826_v9_)
            nw120_[int(4)] = len(d_828_v11_)
            nw120_[int(5)] = -505
            nw120_[int(6)] = (self).f6
            nw120_[int(7)] = p1
            nw120_[int(8)] = (self).f18
            nw120_[int(9)] = -558
            d_829_v12_ = nw120_
            d_830_v13_: D5
            d_830_v13_ = D5_DC17(d_818_v2_, not(d_825_v8_))
            d_831_v14_: _dafny.Map
            d_831_v14_ = _dafny.Map({d_829_v12_: (d_825_v8_ if d_825_v8_ else (d_830_v13_).cf37)})
            d_831_v14_ = (d_831_v14_).set(d_829_v12_, d_825_v8_)
            if d_825_v8_:
                d_832_v15_: _dafny.Map
                d_832_v15_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([(self).f18 for d_833_i1_ in range(default__.abs(199))]): not (d_825_v8_) or (d_825_v8_)})
                d_832_v15_ = d_832_v15_
                (globalState).f0 = d_825_v8_
                d_834_v16_: D4
                d_834_v16_ = D4_DC14(d_825_v8_, default__.fm41(globalState), d_823_i0_, _dafny.Map({d_825_v8_: d_825_v8_}))
                d_835_v17_: _dafny.Map
                d_835_v17_ = _dafny.Map({p1: (d_834_v16_).cf28})
                d_835_v17_ = (d_835_v17_).set(len((d_835_v17_) | (d_835_v17_)), (-185) > ((self).f18))
                d_836_v18_: _dafny.Array
                def lambda27_(d_837_v8_):
                    def lambda28_(d_838_i2_):
                        return _dafny.Map({d_837_v8_: d_837_v8_})

                    return lambda28_

                init18_ = lambda27_(d_825_v8_)
                nw121_ = _dafny.Array(None, 28)
                for i0_18_ in range(nw121_.length(0)):
                    nw121_[i0_18_] = init18_(i0_18_)
                d_836_v18_ = nw121_
                d_839_v19_: _dafny.Set
                d_839_v19_ = _dafny.Set({D9_DC26(d_836_v18_)})
                d_840_v20_: _dafny.Map
                d_840_v20_ = _dafny.Map({d_825_v8_: d_839_v19_})
                d_841_v21_: _dafny.Array
                nw122_ = _dafny.Array(_dafny.MultiSet({}), 13)
                d_841_v21_ = nw122_
                d_842_v22_: D10
                d_842_v22_ = D10_DC28(d_831_v14_)
                rhs142_ = ((d_840_v20_) | (d_840_v20_)) | (d_840_v20_)
                rhs143_ = d_841_v21_
                rhs144_ = d_842_v22_
                rhs145_ = d_825_v8_
                lhs119_ = globalState
                d_840_v20_ = rhs142_
                d_841_v21_ = rhs143_
                d_842_v22_ = rhs144_
                lhs119_.f0 = rhs145_
                d_816_v0_ = d_816_v0_
            elif True:
                arr21_ = self.f9
                index124_ = default__.safeIndex(921, (self.f9).length(0))
                arr21_[index124_] = d_825_v8_
                d_843_v23_: _dafny.Map
                d_843_v23_ = _dafny.Map({not(d_825_v8_): d_825_v8_})
                d_844_v24_: D4
                d_844_v24_ = D4_DC14(d_825_v8_, d_843_v23_, p1, d_843_v23_)
                arr22_ = self.f9
                index125_ = default__.safeIndex(921, (self.f9).length(0))
                arr22_[index125_] = (not(d_825_v8_)) and ((d_844_v24_).cf28)
                d_845_v25_: _dafny.Array
                def lambda29_(d_846_v10_):
                    def lambda30_(d_847_i3_):
                        return ((d_846_v10_).f21) | ((d_846_v10_).f21)

                    return lambda30_

                init19_ = lambda29_(d_827_v10_)
                nw123_ = _dafny.Array(None, 7)
                for i0_19_ in range(nw123_.length(0)):
                    nw123_[i0_19_] = init19_(i0_19_)
                d_845_v25_ = nw123_
                d_845_v25_ = d_845_v25_
                d_848_v26_: _dafny.Seq
                d_848_v26_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_816_v0_])])
                d_849_v27_: D4
                d_849_v27_ = D4_DC12(d_848_v26_)
                d_850_v28_: _dafny.Array
                nw124_ = _dafny.Array(_dafny.MultiSet({}), 21)
                d_850_v28_ = nw124_
                index126_ = default__.safeIndex(569, (d_850_v28_).length(0))
                (d_850_v28_)[index126_] = (_dafny.MultiSet([d_818_v2_])).set(_dafny.SeqWithoutIsStrInference([766]), default__.abs(p1))
                d_851_v29_: _dafny.Array
                nw125_ = _dafny.Array(_dafny.Seq({}), 10)
                d_851_v29_ = nw125_
                index127_ = default__.safeIndex(569, (d_850_v28_).length(0))
                rhs146_ = d_849_v27_
                rhs147_ = d_819_v3_
                rhs148_ = d_851_v29_
                lhs120_ = d_850_v28_
                lhs121_ = default__.safeIndex(569, (d_850_v28_).length(0))
                d_849_v27_ = rhs146_
                lhs120_[lhs121_] = rhs147_
                d_851_v29_ = rhs148_
                r3 = len(((d_843_v23_ if d_825_v8_ else d_843_v23_)) | (d_843_v23_))
                arr23_ = self.f9
                index128_ = default__.safeIndex(921, (self.f9).length(0))
                arr24_ = self.f9
                index129_ = default__.safeIndex(921, (self.f9).length(0))
                arr25_ = self.f9
                index130_ = default__.safeIndex(921, (self.f9).length(0))
                rhs149_ = d_825_v8_
                rhs150_ = not((self.f9)[default__.safeIndex(921, (self.f9).length(0))])
                rhs151_ = (D0_DC1(d_825_v8_, p1, d_816_v0_)).cf4
                rhs152_ = d_825_v8_
                lhs122_ = self.f9
                lhs123_ = default__.safeIndex(921, (self.f9).length(0))
                lhs124_ = self.f9
                lhs125_ = default__.safeIndex(921, (self.f9).length(0))
                lhs126_ = self.f9
                lhs127_ = default__.safeIndex(921, (self.f9).length(0))
                lhs122_[lhs123_] = rhs149_
                lhs124_[lhs125_] = rhs150_
                d_816_v0_ = rhs151_
                lhs126_[lhs127_] = rhs152_
            d_852_v30_: _dafny.Array
            nw126_ = _dafny.Array(_dafny.Seq({}), 29)
            d_852_v30_ = nw126_
            index131_ = default__.safeIndex(466, (d_852_v30_).length(0))
            (d_852_v30_)[index131_] = d_818_v2_
            d_853_v31_: _dafny.Map
            d_853_v31_ = _dafny.Map({False: d_825_v8_})
            d_854_v32_: D4
            d_854_v32_ = D4_DC14(d_825_v8_, d_853_v31_, (self).f17, d_853_v31_)
            d_855_v33_: _dafny.Seq
            d_855_v33_ = _dafny.SeqWithoutIsStrInference([(self).fm22((_dafny.MultiSet([d_825_v8_])).set((d_854_v32_).cf28, default__.abs((self).f6)), False, (self).f6, globalState)])
            d_856_v34_: _dafny.Set
            d_856_v34_ = _dafny.Set({p1, len(d_855_v33_), (d_817_v1_).cardinality, (self).f6})
            index132_ = default__.safeIndex(466, (d_852_v30_).length(0))
            (d_852_v30_)[index132_] = ((d_818_v2_) + (d_818_v2_)) + ((d_818_v2_) + (_dafny.SeqWithoutIsStrInference([len(d_856_v34_), (self).f17, (self).f18])))
            r3 = ((p1) * ((self).f18)) + ((self).f6)
        d_857_v37_: _dafny.Array
        def lambda31_(d_858_v1_):
            def lambda32_(d_859_i4_):
                def iife70_():
                    def iife72_():
                        coll26_ = _dafny.Set()
                        compr_26_: _dafny.Seq
                        for compr_26_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d")): True})).keys.Elements:
                            d_862_v36_: _dafny.Seq = compr_26_
                            if (d_862_v36_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d")): True})):
                                coll26_ = coll26_.union(_dafny.Set([d_862_v36_]))
                        return _dafny.Set(coll26_)
                    coll24_ = _dafny.Map()
                    def iife71_():
                        coll25_ = _dafny.Set()
                        compr_25_: _dafny.Seq
                        for compr_25_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d")): True})).keys.Elements:
                            d_860_v36_: _dafny.Seq = compr_25_
                            if (d_860_v36_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d")): True})):
                                coll25_ = coll25_.union(_dafny.Set([d_860_v36_]))
                        return _dafny.Set(coll25_)
                    compr_24_: _dafny.Seq
                    for compr_24_ in (iife71_()
                    ).Elements:
                        d_861_v35_: _dafny.Seq = compr_24_
                        if (d_861_v35_) in (iife72_()
                        ):
                            coll24_[d_861_v35_] = (d_858_v1_).cardinality
                    return _dafny.Map(coll24_)
                return iife70_()
                

            return lambda32_

        init20_ = lambda31_(d_817_v1_)
        nw127_ = _dafny.Array(None, 3)
        for i0_20_ in range(nw127_.length(0)):
            nw127_[i0_20_] = init20_(i0_20_)
        d_857_v37_ = nw127_
        d_863_v38_: _dafny.Seq
        d_863_v38_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jgfwo"))
        d_864_v39_: _dafny.Map
        d_864_v39_ = _dafny.Map({d_863_v38_: (self).f17})
        index133_ = default__.safeIndex(962, (d_857_v37_).length(0))
        (d_857_v37_)[index133_] = d_864_v39_
        index134_ = default__.safeIndex(962, (d_857_v37_).length(0))
        (d_857_v37_)[index134_] = d_864_v39_
        d_817_v1_ = d_817_v1_
        d_865_v40_: bool
        d_865_v40_ = True
        d_866_v41_: _dafny.Set
        d_866_v41_ = _dafny.Set({d_865_v40_})
        d_866_v41_ = ((d_866_v41_) - (_dafny.Set({False}))) - ((d_866_v41_) - (_dafny.Set({d_865_v40_})))
        hi7_ = len(d_863_v38_)
        for d_867_i5_ in range((self).f17, hi7_):
            d_868_v42_: _dafny.Map
            d_868_v42_ = _dafny.Map({p1: 115})
            d_869_v43_: _dafny.Seq
            d_869_v43_ = _dafny.SeqWithoutIsStrInference([self.f9, self.f9])
            d_870_v44_: _dafny.Seq
            d_870_v44_ = _dafny.SeqWithoutIsStrInference([self.f9, (d_869_v43_)[default__.safeIndex(797, len(d_869_v43_))], self.f9, self.f9])
            d_871_v45_: D4
            d_871_v45_ = D4_DC13(_dafny.SeqWithoutIsStrInference([(self).f6 for d_872_i6_ in range(default__.abs(-484))]), d_865_v40_, (self).f18)
            d_873_v46_: _dafny.Map
            d_873_v46_ = _dafny.Map({d_865_v40_: (d_871_v45_).cf26})
            d_874_v47_: C3
            nw128_ = C3()
            nw128_.ctor__(d_868_v42_, (d_870_v44_)[default__.safeIndex(len(d_873_v46_), len(d_870_v44_))], (self).f10, len(d_863_v38_))
            d_874_v47_ = nw128_
            r1 = (self).f6
            d_875_v48_: bool
            d_876_v49_: bool
            d_877_v50_: _dafny.Array
            out25_: bool
            out26_: bool
            out27_: _dafny.Array
            out25_, out26_, out27_ = (self).m12(d_865_v40_, (_dafny.Set({True})).issubset(d_866_v41_), globalState)
            d_875_v48_ = out25_
            d_876_v49_ = out26_
            d_877_v50_ = out27_
            d_878_v51_: _dafny.Set
            d_878_v51_ = _dafny.Set({8, (0) - ((0) - (p1))})
            if (d_878_v51_) != (d_878_v51_):
                d_879_v52_: _dafny.Array
                d_879_v52_ = self.f9
                d_880_v53_: C0
                nw129_ = C0()
                nw129_.ctor__((d_879_v52_), (default__.fm30(d_865_v40_, d_816_v0_, globalState)) <= (d_863_v38_))
                d_880_v53_ = nw129_
                rhs153_ = p1
                rhs154_ = _dafny.CodePoint('o')
                r3 = rhs153_
                d_816_v0_ = rhs154_
                d_881_v55_: C3
                nw130_ = C3()
                def iife73_():
                    coll27_ = _dafny.Map()
                    compr_27_: int
                    for compr_27_ in _dafny.IntegerRange(978, 814):
                        d_882_v54_: int = compr_27_
                        if ((978) <= (d_882_v54_)) and ((d_882_v54_) < (814)):
                            coll27_[(d_882_v54_) * (456)] = d_863_v38_
                    return _dafny.Map(coll27_)
                nw130_.ctor__(_dafny.Map({(self).f18: (self).f17}), d_880_v53_.f19, ((self).f10) | (iife73_()
                ), (0) - ((self).f6))
                d_881_v55_ = nw130_
                d_883_v56_: _dafny.Map
                d_883_v56_ = _dafny.Map({d_881_v55_: len(d_863_v38_)})
                d_868_v42_ = (d_868_v42_).set(len(d_883_v56_), (self).f6)
                d_884_v57_: _dafny.Map
                d_884_v57_ = _dafny.Map({not(d_875_v48_): d_877_v50_})
                d_885_v58_: _dafny.Map
                d_885_v58_ = _dafny.Map({d_884_v57_: (self).f18})
                d_886_v59_: _dafny.Map
                d_886_v59_ = _dafny.Map({p1: d_884_v57_})
                d_885_v58_ = (d_885_v58_).set(((d_886_v59_)[(self).f6] if ((self).f6) in (d_886_v59_) else d_884_v57_), ((self).f6) - ((d_818_v2_)[default__.safeIndex((self).f6, len(d_818_v2_))]))
            elif True:
                index135_ = default__.safeIndex(492, (d_877_v50_).length(0))
                (d_877_v50_)[index135_] = (d_817_v1_).cardinality
                index136_ = default__.safeIndex(492, (d_877_v50_).length(0))
                (d_877_v50_)[index136_] = (self).f17
                d_887_v60_: _dafny.Set
                d_887_v60_ = _dafny.Set({d_816_v0_})
                d_888_v61_: _dafny.Map
                d_888_v61_ = _dafny.Map({(self).f17: (_dafny.Set({d_816_v0_, d_816_v0_, d_816_v0_})) | (d_887_v60_)})
                def iife74_():
                    def iife75_():
                        coll29_ = _dafny.Map()
                        compr_29_: D0
                        for compr_29_ in (_dafny.MultiSet([D0_DC1(d_876_v49_, 265, d_816_v0_), D0_DC1(d_875_v48_, len(d_863_v38_), d_816_v0_), D0_DC1(d_865_v40_, (self).f18, _dafny.CodePoint('g')), D0_DC1(d_875_v48_, 376, d_816_v0_), D0_DC1(d_876_v49_, (self).f6, d_816_v0_)])).Elements:
                            d_890_v63_: D0 = compr_29_
                            if (d_890_v63_) in (_dafny.MultiSet([D0_DC1(d_876_v49_, 265, d_816_v0_), D0_DC1(d_875_v48_, len(d_863_v38_), d_816_v0_), D0_DC1(d_865_v40_, (self).f18, _dafny.CodePoint('g')), D0_DC1(d_875_v48_, 376, d_816_v0_), D0_DC1(d_876_v49_, (self).f6, d_816_v0_)])):
                                coll29_[d_890_v63_] = (0) - ((d_877_v50_)[default__.safeIndex(492, (d_877_v50_).length(0))])
                        return _dafny.Map(coll29_)
                    coll28_ = _dafny.Map()
                    compr_28_: int
                    for compr_28_ in (((d_874_v47_).f21) | (default__.fm42(d_876_v49_, p0, (0) - ((self).f6), d_816_v0_, globalState))).keys.Elements:
                        d_889_v62_: int = compr_28_
                        if (d_889_v62_) in (((d_874_v47_).f21) | (default__.fm42(d_876_v49_, p0, (0) - ((self).f6), d_816_v0_, globalState))):
                            coll28_[(d_889_v62_) * (len(iife75_()
                            ))] = _dafny.Set({d_816_v0_, d_816_v0_})
                    return _dafny.Map(coll28_)
                d_888_v61_ = iife74_()
                
                d_891_v64_: _dafny.Array
                nw131_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 20)
                d_891_v64_ = nw131_
                index137_ = default__.safeIndex(386, (d_891_v64_).length(0))
                (d_891_v64_)[index137_] = _dafny.SeqWithoutIsStrInference([(d_863_v38_)[default__.safeIndex((d_877_v50_)[default__.safeIndex(492, (d_877_v50_).length(0))], len(d_863_v38_))] for d_892_i7_ in range(default__.abs(948))])
                index138_ = default__.safeIndex(386, (d_891_v64_).length(0))
                (d_891_v64_)[index138_] = d_863_v38_
                d_893_v65_: C0
                nw132_ = C0()
                nw132_.ctor__(self.f9, d_865_v40_)
                d_893_v65_ = nw132_
                d_894_v66_: _dafny.Array
                nw133_ = _dafny.Array(_dafny.Array(None, 0), 8)
                d_894_v66_ = nw133_
                d_895_v67_: _dafny.Array
                nw134_ = _dafny.Array(None, 2)
                nw134_[int(0)] = (self).f18
                nw134_[int(1)] = (self).f6
                d_895_v67_ = nw134_
                index139_ = default__.safeIndex(472, (d_894_v66_).length(0))
                (d_894_v66_)[index139_] = d_895_v67_
                d_896_v68_: _dafny.Array
                def lambda33_(d_897_v65_, d_898_v49_):
                    def lambda34_(d_899_i8_):
                        return _dafny.Map({d_897_v65_.f20: d_898_v49_})

                    return lambda34_

                init21_ = lambda33_(d_893_v65_, d_876_v49_)
                nw135_ = _dafny.Array(None, 4)
                for i0_21_ in range(nw135_.length(0)):
                    nw135_[i0_21_] = init21_(i0_21_)
                d_896_v68_ = nw135_
                d_900_v69_: D9
                d_900_v69_ = D9_DC26(d_896_v68_)
                d_901_v70_: _dafny.Seq
                d_901_v70_ = _dafny.SeqWithoutIsStrInference([d_893_v65_.f20, d_876_v49_])
                d_902_v71_: _dafny.Map
                d_902_v71_ = _dafny.Map({d_900_v69_: (d_901_v70_).set(default__.safeIndex((self).f18, len(d_901_v70_)), d_865_v40_)})
                index140_ = default__.safeIndex(472, (d_894_v66_).length(0))
                nw136_ = _dafny.Array(None, 12)
                nw136_[int(0)] = p1
                nw136_[int(1)] = len(d_902_v71_)
                nw136_[int(2)] = (self).f18
                nw136_[int(3)] = 566
                nw136_[int(4)] = ((self).f6 if d_875_v48_ else (0) - ((self).f17))
                nw136_[int(5)] = default__.safeDivisionInt(114, d_867_i5_)
                nw136_[int(6)] = -915
                nw136_[int(7)] = (self).f17
                nw136_[int(8)] = (self).f6
                nw136_[int(9)] = p1
                nw136_[int(10)] = (d_877_v50_)[default__.safeIndex(492, (d_877_v50_).length(0))]
                nw136_[int(11)] = (self).f6
                (d_894_v66_)[index140_] = nw136_
        d_903_v72_: _dafny.Map
        d_903_v72_ = _dafny.Map({self.f9: d_863_v38_})
        d_904_i9_: int
        d_904_i9_ = 0
        with _dafny.label("5"):
            while (d_903_v72_) != (d_903_v72_):
                with _dafny.c_label("5"):
                    if (d_904_i9_) >= (100):
                        raise _dafny.Break("5")
                    d_904_i9_ = (d_904_i9_) + (1)
                    d_905_v73_: _dafny.Seq
                    d_905_v73_ = _dafny.SeqWithoutIsStrInference([d_865_v40_])
                    d_906_v74_: _dafny.Map
                    d_906_v74_ = _dafny.Map({len(d_863_v38_): d_905_v73_})
                    d_907_v75_: _dafny.Map
                    d_907_v75_ = _dafny.Map({p1: d_865_v40_})
                    d_908_v76_: D5
                    d_908_v76_ = D5_DC17(d_818_v2_, d_865_v40_)
                    d_909_v77_: _dafny.Seq
                    d_909_v77_ = _dafny.SeqWithoutIsStrInference([D5_DC17((_dafny.SeqWithoutIsStrInference([len(_dafny.Map({d_865_v40_: (self).f18})) for d_910_i10_ in range(default__.abs(868))])).set(default__.safeIndex(len(d_906_v74_), len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({d_865_v40_: (self).f18})) for d_911_i10_ in range(default__.abs(868))]))), len(d_907_v75_)), not(d_865_v40_)), D5_DC17(d_818_v2_, d_865_v40_), d_908_v76_])
                    d_912_v78_: _dafny.Seq
                    d_912_v78_ = _dafny.SeqWithoutIsStrInference([default__.fm35((self).f17, len(_dafny.Map({(0) - (p1): self.f9})), d_865_v40_, len(d_866_v41_), globalState)])
                    d_913_v79_: _dafny.Map
                    d_913_v79_ = _dafny.Map({d_909_v77_: ((self).f17 if d_865_v40_ else (((d_912_v78_)[default__.safeIndex((self).f17, len(d_912_v78_))]).set((self).f17, default__.abs((self).f6))).cardinality)})
                    d_913_v79_ = d_913_v79_
                    d_914_v80_: _dafny.Set
                    d_914_v80_ = _dafny.Set({(self).f6})
                    d_915_v81_: _dafny.Seq
                    d_915_v81_ = _dafny.SeqWithoutIsStrInference([(d_863_v38_).set(default__.safeIndex(p1, len(d_863_v38_)), default__.fm0((self).f6, d_914_v80_, (d_905_v73_)[default__.safeIndex((self).f18, len(d_905_v73_))], globalState))])
                    d_916_v82_: _dafny.Array
                    nw137_ = _dafny.Array(None, 17)
                    nw137_[int(0)] = (d_863_v38_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dpxwvvn")))
                    nw137_[int(1)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oplylgwjk"))
                    nw137_[int(2)] = _dafny.SeqWithoutIsStrInference([d_816_v0_ for d_917_i11_ in range(default__.abs(439))])
                    nw137_[int(3)] = d_863_v38_
                    nw137_[int(4)] = d_863_v38_
                    nw137_[int(5)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "huf"))
                    nw137_[int(6)] = default__.fm30(d_865_v40_, _dafny.CodePoint('n'), globalState)
                    nw137_[int(7)] = d_863_v38_
                    nw137_[int(8)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a"))) + (d_863_v38_)
                    nw137_[int(9)] = d_863_v38_
                    nw137_[int(10)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xkaku"))
                    nw137_[int(11)] = d_863_v38_
                    nw137_[int(12)] = ((d_915_v81_)[default__.safeIndex(len(d_905_v73_), len(d_915_v81_))]).set(default__.safeIndex((self).f18, len((d_915_v81_)[default__.safeIndex(len(d_905_v73_), len(d_915_v81_))])), d_816_v0_)
                    nw137_[int(13)] = default__.fm25((self).f17, d_865_v40_, globalState)
                    nw137_[int(14)] = default__.fm30(d_865_v40_, d_816_v0_, globalState)
                    nw137_[int(15)] = d_863_v38_
                    nw137_[int(16)] = d_863_v38_
                    d_916_v82_ = nw137_
                    d_918_v83_: _dafny.Seq
                    d_918_v83_ = _dafny.SeqWithoutIsStrInference([d_916_v82_, d_916_v82_, d_916_v82_, d_916_v82_])
                    d_916_v82_ = (d_918_v83_)[default__.safeIndex(91, len(d_918_v83_))]
                    d_865_v40_ = (-609) <= ((self).f18)
                    if not(d_865_v40_):
                        d_919_v84_: _dafny.Seq
                        d_919_v84_ = _dafny.SeqWithoutIsStrInference([D5_DC16(d_866_v41_)])
                        d_920_v85_: _dafny.Set
                        d_920_v85_ = _dafny.Set({d_919_v84_})
                        d_921_v86_: _dafny.Map
                        d_921_v86_ = _dafny.Map({d_919_v84_: (self).f6})
                        def iife76_():
                            coll30_ = _dafny.Set()
                            compr_30_: _dafny.Seq
                            for compr_30_ in (d_921_v86_).keys.Elements:
                                d_922_v87_: _dafny.Seq = compr_30_
                                if (d_922_v87_) in (d_921_v86_):
                                    coll30_ = coll30_.union(_dafny.Set([d_922_v87_]))
                            return _dafny.Set(coll30_)
                        d_920_v85_ = iife76_()
                        
                        d_923_v88_: D9
                        d_923_v88_ = D9_DC25((p1 if d_865_v40_ else (self).f6), ((d_817_v1_)[d_865_v40_] if (d_865_v40_) in (d_817_v1_) else 139), (len(_dafny.SeqWithoutIsStrInference([(self).f18 for d_924_i12_ in range(default__.abs(-255))])) if d_865_v40_ else (0) - ((self).f18)), d_816_v0_)
                        d_923_v88_ = d_923_v88_
                        d_925_v89_: D0
                        d_925_v89_ = D0_DC0(d_905_v73_, (self).f18)
                        rhs155_ = (d_863_v38_) + (d_863_v38_)
                        rhs156_ = (((d_925_v89_).cf1 if d_865_v40_ else (self).f18)) + (((self).f17) * ((self).f17))
                        lhs128_ = globalState
                        lhs128_.f2 = rhs155_
                        r3 = rhs156_
                        d_926_v90_: _dafny.Array
                        nw138_ = _dafny.Array(int(0), 17)
                        d_926_v90_ = nw138_
                        index141_ = default__.safeIndex(606, (d_926_v90_).length(0))
                        (d_926_v90_)[index141_] = (self).f18
                        index142_ = default__.safeIndex(606, (d_926_v90_).length(0))
                        rhs157_ = d_865_v40_
                        rhs158_ = default__.safeModuloInt(-824, p1)
                        rhs159_ = (_dafny.Set({d_865_v40_})) - (d_866_v41_)
                        rhs160_ = (D4_DC13(d_818_v2_, d_865_v40_, (0) - ((self).f17))).cf27
                        lhs129_ = globalState
                        lhs130_ = d_926_v90_
                        lhs131_ = default__.safeIndex(606, (d_926_v90_).length(0))
                        lhs129_.f0 = rhs157_
                        r3 = rhs158_
                        d_866_v41_ = rhs159_
                        lhs130_[lhs131_] = rhs160_
                        r3 = len((d_863_v38_) + (d_863_v38_))
                    elif True:
                        (globalState).f0 = d_865_v40_
                        arr26_ = self.f9
                        index143_ = default__.safeIndex(257, (self.f9).length(0))
                        arr26_[index143_] = d_865_v40_
                        d_927_v91_: _dafny.Seq
                        d_927_v91_ = _dafny.SeqWithoutIsStrInference([d_818_v2_, d_818_v2_])
                        arr27_ = self.f9
                        index144_ = default__.safeIndex(257, (self.f9).length(0))
                        rhs161_ = (len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "afv"))) + (d_863_v38_))) > (default__.safeDivisionInt((self).f18, (0) - (p1)))
                        rhs162_ = d_927_v91_
                        rhs163_ = (self).f6
                        lhs132_ = self.f9
                        lhs133_ = default__.safeIndex(257, (self.f9).length(0))
                        lhs134_ = globalState
                        lhs132_[lhs133_] = rhs161_
                        d_927_v91_ = rhs162_
                        lhs134_.f5 = rhs163_
                        arr28_ = self.f9
                        index145_ = default__.safeIndex(257, (self.f9).length(0))
                        arr28_[index145_] = (self.f9)[default__.safeIndex(257, (self.f9).length(0))]
                        d_928_v92_: _dafny.MultiSet
                        d_928_v92_ = d_817_v1_
                        d_929_v93_: _dafny.Map
                        d_929_v93_ = _dafny.Map({(self).f6: (d_928_v92_)})
                        d_817_v1_ = (d_817_v1_) | (((d_817_v1_).set(False, default__.abs((self).f17))) - (((d_929_v93_)[(self).fm21(len(d_905_v73_), globalState)] if ((self).fm21(len(d_905_v73_), globalState)) in (d_929_v93_) else _dafny.MultiSet([d_865_v40_, (self.f9)[default__.safeIndex(257, (self.f9).length(0))]]))))
                        (globalState).f5 = (0) - ((self).f6)
                    pass
            pass
        d_930_v94_: D7
        d_930_v94_ = D7_DC22()
        pat_let_tv18_ = d_816_v0_
        def lambda35_(source12_):
            if source12_.is_DC22:
                return pat_let_tv18_
            elif True:
                d_931___mcc_h0_ = source12_.cf43
                d_932_cf43_ = d_931___mcc_h0_
                return _dafny.CodePoint('r')

        r0 = lambda35_((d_930_v94_ if d_865_v40_ else d_930_v94_))
        d_933_v95_: _dafny.MultiSet
        d_933_v95_ = _dafny.MultiSet([self.f9])
        r1 = ((d_933_v95_)[self.f9] if (self.f9) in (d_933_v95_) else (self).f18)
        r2 = self.f9
        r3 = (self).f17
        return r0, r1, r2, r3

    def m12(self, p0, p1, globalState):
        r0: bool = False
        r1: bool = False
        r2: _dafny.Array = _dafny.Array(None, 0)
        d_934_v0_: _dafny.Seq
        d_934_v0_ = _dafny.SeqWithoutIsStrInference([(self).f6])
        d_935_v1_: _dafny.Seq
        d_935_v1_ = _dafny.SeqWithoutIsStrInference([((self).f18) - ((self).f17), (d_934_v0_)[default__.safeIndex((self).f18, len(d_934_v0_))], (self).f18, default__.safeDivisionInt(880, (self).f6), ((self).f6) + ((self).f18)])
        d_936_v2_: _dafny.Seq
        d_936_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xxnkqk"))
        rhs164_ = d_936_v2_
        rhs165_ = d_934_v0_
        rhs166_ = (899) - ((self).f17)
        lhs135_ = globalState
        lhs136_ = globalState
        lhs135_.f2 = rhs164_
        d_934_v0_ = rhs165_
        lhs136_.f5 = rhs166_
        arr29_ = self.f9
        index146_ = default__.safeIndex(202, (self.f9).length(0))
        arr29_[index146_] = p0
        d_937_v3_: _dafny.Map
        d_937_v3_ = _dafny.Map({240: p0})
        d_938_v4_: _dafny.Map
        d_938_v4_ = _dafny.Map({(self).fm22((_dafny.MultiSet([not(True)])), ((d_937_v3_)[(self).f17] if ((self).f17) in (d_937_v3_) else p1), len(_dafny.SeqWithoutIsStrInference([(self).f17, (self).f17, (self).f17])), globalState): not (p1) or (p1)})
        d_939_v5_: _dafny.Map
        d_939_v5_ = _dafny.Map({(self).f17: (self).f6})
        d_940_v6_: _dafny.MultiSet
        d_940_v6_ = _dafny.MultiSet([(self).f18])
        d_941_v8_: _dafny.Set
        d_941_v8_ = _dafny.Set({(self).f6})
        d_942_v9_: C2
        nw139_ = C2()
        def iife77_():
            coll31_ = _dafny.Map()
            compr_31_: int
            for compr_31_ in (d_941_v8_).Elements:
                d_943_v7_: int = compr_31_
                if (d_943_v7_) in (d_941_v8_):
                    coll31_[default__.safeModuloInt(d_943_v7_, -615)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m"))
            return _dafny.Map(coll31_)
        nw139_.ctor__(len((d_939_v5_) | (d_939_v5_)), _dafny.MultiSet([d_940_v6_]), self.f9, (iife77_()
        ).set((self).f17, d_936_v2_))
        d_942_v9_ = nw139_
        d_944_v10_: _dafny.Seq
        d_944_v10_ = _dafny.SeqWithoutIsStrInference([d_937_v3_])
        d_945_v11_: _dafny.Seq
        d_945_v11_ = _dafny.SeqWithoutIsStrInference([True])
        d_946_v12_: _dafny.Map
        d_946_v12_ = _dafny.Map({self.f9: _dafny.MultiSet((d_945_v11_).set(default__.safeIndex(len(d_936_v2_), len(d_945_v11_)), p1))})
        d_947_v13_: _dafny.Seq
        d_947_v13_ = _dafny.SeqWithoutIsStrInference([((d_944_v10_)[default__.safeIndex((self).f6, len(d_944_v10_))]).set(len(d_946_v12_), p1), d_937_v3_])
        d_948_v14_: _dafny.Set
        d_948_v14_ = _dafny.Set({not(False), ((self).f6) > ((self).f17), (d_944_v10_) == (d_947_v13_)})
        d_949_v15_: str
        d_949_v15_ = _dafny.CodePoint('y')
        d_950_v16_: _dafny.Map
        d_950_v16_ = _dafny.Map({d_949_v15_: not((d_945_v11_)[default__.safeIndex((self).f6, len(d_945_v11_))])})
        arr30_ = self.f9
        index147_ = default__.safeIndex(202, (self.f9).length(0))
        rhs167_ = ((d_950_v16_)[d_949_v15_] if (d_949_v15_) in (d_950_v16_) else default__.fm3(globalState))
        rhs168_ = d_938_v4_
        rhs169_ = d_942_v9_
        rhs170_ = d_948_v14_
        rhs171_ = p1
        lhs137_ = self.f9
        lhs138_ = default__.safeIndex(202, (self.f9).length(0))
        lhs137_[lhs138_] = rhs167_
        d_938_v4_ = rhs168_
        d_942_v9_ = rhs169_
        d_948_v14_ = rhs170_
        r1 = rhs171_
        hi8_ = (len(d_945_v11_) if not(p1) else (self).f6)
        for d_951_i0_ in range((self).f18, hi8_):
            if (((0) - ((d_942_v9_).fm21((d_940_v6_).cardinality, globalState))) != ((0) - ((self).f17))) in (d_945_v11_):
                (globalState).f5 = (d_951_i0_ if p0 else d_951_i0_)
                (globalState).f0 = p0
                d_952_v17_: _dafny.Array
                nw140_ = _dafny.Array(None, 28)
                d_952_v17_ = nw140_
                d_953_v18_: _dafny.Map
                d_953_v18_ = _dafny.Map({d_952_v17_: _dafny.Map({len(_dafny.Map({p1: (self).f17})): (self).f17})})
                d_953_v18_ = d_953_v18_
                (globalState).f0 = not(not(p1))
                d_954_v19_: _dafny.Seq
                d_954_v19_ = _dafny.SeqWithoutIsStrInference([d_934_v0_, d_934_v0_, d_934_v0_, _dafny.SeqWithoutIsStrInference([(self).f6])])
                d_955_v20_: _dafny.Map
                d_955_v20_ = _dafny.Map({(d_954_v19_) + (default__.fm43(d_949_v15_, (self).f17, globalState)): (self).f6})
                d_955_v20_ = (d_955_v20_).set(d_954_v19_, (self).f17)
            elif True:
                d_956_v21_: _dafny.Map
                d_956_v21_ = _dafny.Map({(self.f9)[default__.safeIndex(202, (self.f9).length(0))]: (self).f18})
                (globalState).f5 = (D0_DC1(p1, len(_dafny.Set({(self).f6, len(d_956_v21_)})), d_949_v15_)).cf3
                d_957_v22_: _dafny.Array
                nw141_ = _dafny.Array(False, 23)
                d_957_v22_ = nw141_
                d_958_v23_: C2
                nw142_ = C2()
                nw142_.ctor__((0) - ((self).f18), (self.f16).intersection(self.f16), d_957_v22_, _dafny.Map({(self).f18: (d_936_v2_).set(default__.safeIndex((self).f18, len(d_936_v2_)), d_949_v15_)}))
                d_958_v23_ = nw142_
                d_959_v24_: D3
                d_959_v24_ = D3_DC9((d_935_v1_).set(default__.safeIndex((self).f6, len(d_935_v1_)), default__.fm2(d_951_i0_, globalState)))
                d_960_v25_: C2
                nw143_ = C2()
                nw143_.ctor__((self).f18, (self.f16).set(_dafny.MultiSet((d_959_v24_).cf21), default__.abs(len(d_936_v2_))), (d_957_v22_), (self).f10)
                d_960_v25_ = nw143_
                (globalState).f0 = p1
                d_949_v15_ = (d_949_v15_ if (self.f9)[default__.safeIndex(202, (self.f9).length(0))] else d_949_v15_)
            (globalState).f5 = d_951_i0_
            d_961_v26_: _dafny.Map
            d_961_v26_ = _dafny.Map({p0: d_951_i0_})
            d_961_v26_ = (d_961_v26_).set(((self).f18) > ((self).f18), d_951_i0_)
            d_962_v27_: D2
            d_963_v28_: _dafny.Seq
            d_964_v29_: _dafny.Seq
            out28_: D2
            out29_: _dafny.Seq
            out30_: _dafny.Seq
            out28_, out29_, out30_ = (self).m13((self).f18, globalState)
            d_962_v27_ = out28_
            d_963_v28_ = out29_
            d_964_v29_ = out30_
        d_965_v30_: _dafny.Array
        nw144_ = _dafny.Array(False, 5)
        d_965_v30_ = nw144_
        d_966_v31_: C0
        nw145_ = C0()
        nw145_.ctor__(d_965_v30_, (self.f9)[default__.safeIndex(202, (self.f9).length(0))])
        d_966_v31_ = nw145_
        d_967_v32_: _dafny.Map
        d_967_v32_ = _dafny.Map({p0: (self).f17})
        d_968_v33_: _dafny.MultiSet
        d_968_v33_ = _dafny.MultiSet([p1])
        d_940_v6_ = default__.fm35(len(d_967_v32_), (self).f17, ((self).f18) == ((d_940_v6_).cardinality), ((d_968_v33_)[p1] if (p1) in (d_968_v33_) else (self).f17), globalState)
        arr31_ = self.f9
        index148_ = default__.safeIndex(202, (self.f9).length(0))
        rhs172_ = False
        rhs173_ = d_949_v15_
        lhs139_ = self.f9
        lhs140_ = default__.safeIndex(202, (self.f9).length(0))
        lhs139_[lhs140_] = rhs172_
        d_949_v15_ = rhs173_
        r0 = p1
        r1 = p0
        d_969_v34_: _dafny.Array
        def lambda36_(d_970_v33_):
            def lambda37_(d_971_i1_):
                return default__.safeDivisionInt(d_971_i1_, (d_970_v33_).cardinality)

            return lambda37_

        init22_ = lambda36_(d_968_v33_)
        nw146_ = _dafny.Array(None, 24)
        for i0_22_ in range(nw146_.length(0)):
            nw146_[i0_22_] = init22_(i0_22_)
        d_969_v34_ = nw146_
        r2 = d_969_v34_
        return r0, r1, r2

    def m13(self, p0, globalState):
        r0: D2 = D2.default()()
        r1: _dafny.Seq = _dafny.Seq({})
        r2: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        if ((self).f17) == (default__.safeDivisionInt(p0, p0)):
            d_972_v0_: bool
            d_972_v0_ = False
            d_973_v1_: str
            d_973_v1_ = _dafny.CodePoint('j')
            d_974_v2_: _dafny.Seq
            d_974_v2_ = _dafny.SeqWithoutIsStrInference([(self).f18])
            d_975_v3_: D4
            d_975_v3_ = D4_DC13(d_974_v2_, d_972_v0_, (self).f18)
            d_976_v4_: _dafny.MultiSet
            d_976_v4_ = _dafny.MultiSet([(d_975_v3_).cf26])
            d_977_v5_: _dafny.Set
            d_977_v5_ = _dafny.Set({d_972_v0_, True, d_972_v0_})
            d_978_v6_: _dafny.Map
            d_978_v6_ = _dafny.Map({(d_973_v1_ if d_972_v0_ else _dafny.CodePoint('a')): default__.fm44(p0, (d_976_v4_).cardinality, d_972_v0_, len(d_977_v5_), globalState)})
            d_979_v7_: D9
            d_979_v7_ = D9_DC25((0) - ((self).f17), len(d_974_v2_), (self).f17, d_973_v1_)
            d_978_v6_ = (d_978_v6_).set(d_973_v1_, _dafny.Map({p0: d_979_v7_}))
            if False:
                d_980_v8_: C2
                nw147_ = C2()
                nw147_.ctor__(-796, self.f16, self.f9, (self).f10)
                d_980_v8_ = nw147_
                d_981_v9_: _dafny.Map
                d_981_v9_ = _dafny.Map({d_980_v8_: d_972_v0_})
                d_982_v10_: _dafny.Seq
                d_982_v10_ = _dafny.SeqWithoutIsStrInference([not(d_972_v0_)])
                d_983_v11_: _dafny.Map
                d_983_v11_ = d_981_v9_
                d_984_v12_: _dafny.Array
                nw148_ = _dafny.Array(None, 28)
                nw148_[int(0)] = d_981_v9_
                nw148_[int(1)] = d_981_v9_
                nw148_[int(2)] = d_981_v9_
                nw148_[int(3)] = d_981_v9_
                nw148_[int(4)] = (_dafny.Map({d_980_v8_: d_972_v0_}) if d_972_v0_ else d_981_v9_)
                nw148_[int(5)] = d_981_v9_
                nw148_[int(6)] = d_981_v9_
                nw148_[int(7)] = (d_981_v9_) | (d_981_v9_)
                nw148_[int(8)] = d_981_v9_
                nw148_[int(9)] = d_981_v9_
                nw148_[int(10)] = d_981_v9_
                nw148_[int(11)] = (d_981_v9_ if True else d_981_v9_)
                nw148_[int(12)] = (d_981_v9_).set(d_980_v8_, d_972_v0_)
                nw148_[int(13)] = (d_981_v9_).set(d_980_v8_, True)
                nw148_[int(14)] = d_981_v9_
                nw148_[int(15)] = (_dafny.Map({d_980_v8_: d_972_v0_})) | (d_981_v9_)
                nw148_[int(16)] = (d_981_v9_) | ((d_981_v9_).set(d_980_v8_, (d_982_v10_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tk"))), len(d_982_v10_))]))
                nw148_[int(17)] = (_dafny.Map({d_980_v8_: d_972_v0_})) | (d_981_v9_)
                nw148_[int(18)] = (d_983_v11_)
                nw148_[int(19)] = (d_981_v9_).set(d_980_v8_, d_972_v0_)
                nw148_[int(20)] = d_981_v9_
                nw148_[int(21)] = d_981_v9_
                nw148_[int(22)] = (d_981_v9_) | (d_981_v9_)
                nw148_[int(23)] = d_981_v9_
                nw148_[int(24)] = d_981_v9_
                nw148_[int(25)] = d_981_v9_
                nw148_[int(26)] = d_981_v9_
                nw148_[int(27)] = d_981_v9_
                d_984_v12_ = nw148_
                index149_ = default__.safeIndex(608, (d_984_v12_).length(0))
                (d_984_v12_)[index149_] = (d_981_v9_) | (d_981_v9_)
                d_985_v13_: _dafny.Map
                d_985_v13_ = _dafny.Map({d_972_v0_: _dafny.Map({d_972_v0_: True})})
                d_986_v14_: _dafny.Map
                d_986_v14_ = _dafny.Map({d_972_v0_: False})
                index150_ = default__.safeIndex(608, (d_984_v12_).length(0))
                rhs174_ = d_981_v9_
                rhs175_ = default__.fm45((len(d_974_v2_)) - ((self).f17), (self).f6, globalState)
                rhs176_ = (d_982_v10_) + ((d_982_v10_) + (default__.fm4((self).f18, p0, len(((d_985_v13_)[d_972_v0_] if (d_972_v0_) in (d_985_v13_) else d_986_v14_)), globalState)))
                rhs177_ = d_972_v0_
                rhs178_ = d_972_v0_
                lhs141_ = d_984_v12_
                lhs142_ = default__.safeIndex(608, (d_984_v12_).length(0))
                lhs143_ = globalState
                lhs144_ = globalState
                lhs141_[lhs142_] = rhs174_
                d_975_v3_ = rhs175_
                d_982_v10_ = rhs176_
                lhs143_.f0 = rhs177_
                lhs144_.f0 = rhs178_
                d_987_v15_: C1
                nw149_ = C1()
                nw149_.ctor__(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uemcv"))))
                d_987_v15_ = nw149_
                d_988_v16_: _dafny.Array
                def lambda38_(d_989_i0_):
                    return (d_989_i0_) * ((0) - ((self).f17))

                init23_ = lambda38_
                nw150_ = _dafny.Array(None, 23)
                for i0_23_ in range(nw150_.length(0)):
                    nw150_[i0_23_] = init23_(i0_23_)
                d_988_v16_ = nw150_
                d_990_v17_: D2
                d_990_v17_ = D2_DC6((self).f17, d_988_v16_)
                d_988_v16_ = (d_990_v17_).cf14
                d_991_v18_: C1
                nw151_ = C1()
                nw151_.ctor__((self).f18)
                d_991_v18_ = nw151_
                (self).f9 = self.f9
            elif True:
                d_992_v19_: T1
                nw152_ = C2()
                nw152_.ctor__((self).f17, self.f16, self.f9, _dafny.Map({(self).f18: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nbgkv"))}))
                d_992_v19_ = nw152_
                d_993_v20_: D6
                d_993_v20_ = D6_DC18(d_992_v19_)
                d_994_v21_: D6
                d_994_v21_ = D6_DC20(d_993_v20_)
                d_994_v21_ = D6_DC20(d_993_v20_)
                d_995_v22_: _dafny.Map
                d_995_v22_ = _dafny.Map({d_972_v0_: d_972_v0_})
                d_996_v23_: _dafny.Seq
                d_996_v23_ = _dafny.SeqWithoutIsStrInference([d_972_v0_, ((d_995_v22_)[False] if (False) in (d_995_v22_) else True)])
                d_997_v24_: D6
                d_997_v24_ = D6_DC19(len(d_996_v23_), len(d_974_v2_), (d_992_v19_).f6)
                d_998_v25_: _dafny.MultiSet
                d_998_v25_ = _dafny.MultiSet([d_997_v24_, d_997_v24_, d_997_v24_, d_997_v24_])
                (globalState).f5 = ((d_998_v25_)[d_997_v24_] if (d_997_v24_) in (d_998_v25_) else (self).f6)
                d_999_v26_: _dafny.Array
                nw153_ = _dafny.Array(_dafny.CodePoint('D'), 7)
                d_999_v26_ = nw153_
                d_1000_v27_: _dafny.Map
                d_1000_v27_ = _dafny.Map({not (d_972_v0_) or (d_972_v0_): d_999_v26_})
                d_1001_v28_: _dafny.Seq
                d_1001_v28_ = _dafny.SeqWithoutIsStrInference([d_999_v26_, d_999_v26_])
                d_1000_v27_ = (d_1000_v27_).set(d_972_v0_, ((d_1000_v27_)[d_972_v0_] if (d_972_v0_) in (d_1000_v27_) else (d_1001_v28_)[default__.safeIndex((self).f17, len(d_1001_v28_))]))
                d_1002_v29_: _dafny.Array
                nw154_ = _dafny.Array(int(0), 6)
                d_1002_v29_ = nw154_
                d_1003_v30_: _dafny.Map
                d_1003_v30_ = _dafny.Map({d_972_v0_: p0})
                index151_ = default__.safeIndex(768, (d_1002_v29_).length(0))
                (d_1002_v29_)[index151_] = default__.safeDivisionInt(264, len(d_1003_v30_))
                d_1004_v31_: _dafny.Map
                d_1004_v31_ = _dafny.Map({914: d_972_v0_})
                index152_ = default__.safeIndex(768, (d_1002_v29_).length(0))
                (d_1002_v29_)[index152_] = default__.fm2(((d_1003_v30_)[d_972_v0_] if (d_972_v0_) in (d_1003_v30_) else len(d_1004_v31_)), globalState)
                d_1005_v32_: _dafny.Seq
                d_1005_v32_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lnnj"))
                d_1006_v33_: _dafny.Seq
                d_1006_v33_ = _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hieeu"))) + (d_1005_v32_), d_1005_v32_, _dafny.SeqWithoutIsStrInference([d_973_v1_ for d_1007_i1_ in range(default__.abs(130))])])
                rhs179_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xrfic")) for d_1008_i2_ in range(default__.abs(704))])
                rhs180_ = (self).f18
                lhs145_ = globalState
                d_1006_v33_ = rhs179_
                lhs145_.f5 = rhs180_
            d_1009_v34_: _dafny.Map
            d_1009_v34_ = _dafny.Map({(self).fm22(_dafny.MultiSet([d_972_v0_, not(d_972_v0_)]), d_972_v0_, 66, globalState): (self).f6})
            d_1010_v35_: _dafny.Array
            def lambda39_(d_1011_i3_):
                return default__.safeDivisionInt(d_1011_i3_, -228)

            init24_ = lambda39_
            nw155_ = _dafny.Array(None, 12)
            for i0_24_ in range(nw155_.length(0)):
                nw155_[i0_24_] = init24_(i0_24_)
            d_1010_v35_ = nw155_
            d_1012_v36_: _dafny.MultiSet
            d_1012_v36_ = _dafny.MultiSet([(self).f17])
            index153_ = default__.safeIndex(784, (d_1010_v35_).length(0))
            (d_1010_v35_)[index153_] = (p0 if d_972_v0_ else (d_1012_v36_).cardinality)
            d_1013_v37_: _dafny.Seq
            d_1013_v37_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w"))
            index154_ = default__.safeIndex(784, (d_1010_v35_).length(0))
            rhs181_ = len((d_1013_v37_) + (d_1013_v37_))
            rhs182_ = d_1009_v34_
            rhs183_ = (0) - (-547)
            rhs184_ = (self).f18
            rhs185_ = (d_1010_v35_ if d_972_v0_ else (d_1010_v35_ if d_972_v0_ else d_1010_v35_))
            lhs146_ = globalState
            lhs147_ = d_1010_v35_
            lhs148_ = default__.safeIndex(784, (d_1010_v35_).length(0))
            lhs149_ = globalState
            lhs146_.f5 = rhs181_
            d_1009_v34_ = rhs182_
            lhs147_[lhs148_] = rhs183_
            lhs149_.f5 = rhs184_
            d_1010_v35_ = rhs185_
            if d_972_v0_:
                (globalState).f5 = default__.safeDivisionInt((self).f6, p0)
                d_1014_v38_: C2
                nw156_ = C2()
                nw156_.ctor__((p0) + (len(d_1013_v37_)), self.f16, self.f9, ((self).f10) | (_dafny.Map({(self).f17: (d_1013_v37_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_973_v1_ for d_1015_i4_ in range(default__.abs(939))])), len(d_1013_v37_)), d_973_v1_)})))
                d_1014_v38_ = nw156_
                d_1016_v39_: C1
                nw157_ = C1()
                nw157_.ctor__(default__.safeDivisionInt((self).f17, (self).f17))
                d_1016_v39_ = nw157_
                (globalState).f0 = False
                d_1009_v34_ = (d_1009_v34_).set(((self).f18) <= (-227), 68)
            elif True:
                d_1017_v40_: _dafny.Map
                d_1017_v40_ = _dafny.Map({d_972_v0_: d_972_v0_})
                d_1017_v40_ = (d_1017_v40_).set(d_972_v0_, d_972_v0_)
                d_1018_v41_: _dafny.Array
                nw158_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 7)
                d_1018_v41_ = nw158_
                d_1019_v42_: _dafny.Array
                nw159_ = _dafny.Array(None, 6)
                nw159_[int(0)] = d_1018_v41_
                nw159_[int(1)] = d_1018_v41_
                nw159_[int(2)] = d_1018_v41_
                nw159_[int(3)] = d_1018_v41_
                nw159_[int(4)] = d_1018_v41_
                nw159_[int(5)] = d_1018_v41_
                d_1019_v42_ = nw159_
                index155_ = default__.safeIndex(807, (d_1019_v42_).length(0))
                (d_1019_v42_)[index155_] = d_1018_v41_
                index156_ = default__.safeIndex(807, (d_1019_v42_).length(0))
                rhs186_ = d_1018_v41_
                rhs187_ = (self).f17
                rhs188_ = not(d_972_v0_)
                lhs150_ = d_1019_v42_
                lhs151_ = default__.safeIndex(807, (d_1019_v42_).length(0))
                lhs152_ = globalState
                lhs153_ = globalState
                lhs150_[lhs151_] = rhs186_
                lhs152_.f5 = rhs187_
                lhs153_.f0 = rhs188_
                d_1020_v43_: _dafny.Seq
                d_1020_v43_ = _dafny.SeqWithoutIsStrInference([d_972_v0_])
                r1 = ((_dafny.SeqWithoutIsStrInference([d_972_v0_])) + (d_1020_v43_)) + (d_1020_v43_)
                (globalState).f5 = (self).f17
                arr32_ = self.f9
                index157_ = default__.safeIndex(245, (self.f9).length(0))
                arr32_[index157_] = d_972_v0_
                arr33_ = self.f9
                index158_ = default__.safeIndex(245, (self.f9).length(0))
                arr33_[index158_] = (not (d_972_v0_) or (d_972_v0_)) == (d_972_v0_)
            d_1021_v44_: _dafny.Array
            def lambda40_(d_1022_v0_, d_1023_v5_):
                def lambda41_(d_1024_i5_):
                    return (D5_DC16(_dafny.Set({d_1022_v0_})) if d_1022_v0_ else D5_DC16(d_1023_v5_))

                return lambda41_

            init25_ = lambda40_(d_972_v0_, d_977_v5_)
            nw160_ = _dafny.Array(None, 10)
            for i0_25_ in range(nw160_.length(0)):
                nw160_[i0_25_] = init25_(i0_25_)
            d_1021_v44_ = nw160_
            d_1025_v45_: D5
            d_1025_v45_ = D5_DC16(d_977_v5_)
            index159_ = default__.safeIndex(84, (d_1021_v44_).length(0))
            (d_1021_v44_)[index159_] = d_1025_v45_
            index160_ = default__.safeIndex(84, (d_1021_v44_).length(0))
            (d_1021_v44_)[index160_] = d_1025_v45_
        elif True:
            d_1026_v46_: bool
            d_1026_v46_ = False
            d_1027_v47_: _dafny.Seq
            d_1027_v47_ = _dafny.SeqWithoutIsStrInference([p0])
            d_1028_v48_: _dafny.Map
            d_1028_v48_ = _dafny.Map({d_1026_v46_: d_1027_v47_})
            d_1029_v49_: str
            d_1029_v49_ = _dafny.CodePoint('j')
            d_1030_v50_: _dafny.Map
            d_1030_v50_ = _dafny.Map({(self).f18: d_1026_v46_})
            d_1031_v51_: _dafny.MultiSet
            d_1031_v51_ = _dafny.MultiSet([d_1030_v50_])
            d_1028_v48_ = (d_1028_v48_).set((p0) > (p0), (default__.fm26((self).f6, d_1029_v49_, not(d_1026_v46_), d_1029_v49_, globalState)) + (_dafny.SeqWithoutIsStrInference([default__.fm2((d_1031_v51_).cardinality, globalState), 904])))
            d_1032_v52_: D6
            d_1032_v52_ = D6_DC19(313, p0, (self).f17)
            d_1033_v53_: D6
            d_1033_v53_ = D6_DC20(d_1032_v52_)
            source13_ = d_1033_v53_
            if source13_.is_DC19:
                d_1034___mcc_h0_ = source13_.cf39
                d_1035___mcc_h1_ = source13_.cf40
                d_1036___mcc_h2_ = source13_.cf41
                d_1037_cf41_ = d_1036___mcc_h2_
                d_1038_cf40_ = d_1035___mcc_h1_
                d_1039_cf39_ = d_1034___mcc_h0_
                arr34_ = self.f9
                index161_ = default__.safeIndex(900, (self.f9).length(0))
                arr34_[index161_] = d_1026_v46_
                arr35_ = self.f9
                index162_ = default__.safeIndex(900, (self.f9).length(0))
                arr35_[index162_] = d_1026_v46_
                d_1040_v54_: _dafny.Array
                nw161_ = _dafny.Array(int(0), 20)
                d_1040_v54_ = nw161_
                index163_ = default__.safeIndex(560, (d_1040_v54_).length(0))
                (d_1040_v54_)[index163_] = (self).f6
                d_1041_v55_: _dafny.Seq
                d_1041_v55_ = _dafny.SeqWithoutIsStrInference([d_1040_v54_, d_1040_v54_, d_1040_v54_, d_1040_v54_, d_1040_v54_])
                d_1042_v56_: _dafny.Map
                d_1042_v56_ = _dafny.Map({d_1029_v49_: (self).f10})
                d_1043_v57_: D3
                d_1043_v57_ = D3_DC9(d_1027_v47_)
                d_1044_v58_: _dafny.MultiSet
                d_1044_v58_ = _dafny.MultiSet([(d_1043_v57_).cf21])
                d_1045_v59_: D0
                d_1045_v59_ = D0_DC0(default__.fm4(p0, d_1038_cf40_, (d_1044_v58_).cardinality, globalState), 644)
                arr36_ = self.f9
                index164_ = default__.safeIndex(900, (self.f9).length(0))
                index165_ = default__.safeIndex(560, (d_1040_v54_).length(0))
                rhs189_ = ((0) - (p0)) in (((d_1042_v56_)[d_1029_v49_] if (d_1029_v49_) in (d_1042_v56_) else (self).f10))
                rhs190_ = (d_1045_v59_).cf1
                rhs191_ = default__.safeDivisionInt((0) - (d_1039_cf39_), 481)
                rhs192_ = d_1041_v55_
                lhs154_ = self.f9
                lhs155_ = default__.safeIndex(900, (self.f9).length(0))
                lhs156_ = d_1040_v54_
                lhs157_ = default__.safeIndex(560, (d_1040_v54_).length(0))
                lhs158_ = globalState
                lhs154_[lhs155_] = rhs189_
                lhs156_[lhs157_] = rhs190_
                lhs158_.f5 = rhs191_
                d_1041_v55_ = rhs192_
                d_1046_v60_: _dafny.Array
                def lambda42_(d_1047_v46_):
                    def lambda43_(d_1048_i6_):
                        return _dafny.Map({False: d_1047_v46_})

                    return lambda43_

                init26_ = lambda42_(d_1026_v46_)
                nw162_ = _dafny.Array(None, 6)
                for i0_26_ in range(nw162_.length(0)):
                    nw162_[i0_26_] = init26_(i0_26_)
                d_1046_v60_ = nw162_
                d_1049_v61_: D9
                d_1049_v61_ = D9_DC26(d_1046_v60_)
                d_1049_v61_ = d_1049_v61_
                (globalState).f0 = (self.f9)[default__.safeIndex(900, (self.f9).length(0))]
            elif source13_.is_DC18:
                d_1050___mcc_h3_ = source13_.cf38
                d_1051_cf38_ = d_1050___mcc_h3_
                d_1052_v62_: _dafny.Seq
                d_1052_v62_ = _dafny.SeqWithoutIsStrInference([d_1027_v47_, d_1027_v47_, d_1027_v47_])
                d_1052_v62_ = d_1052_v62_
                (globalState).f5 = (self).f17
                d_1053_v63_: _dafny.Map
                d_1053_v63_ = _dafny.Map({p0: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_1054_i7_ in range(default__.abs(59))])})
                d_1055_v64_: _dafny.Seq
                d_1055_v64_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oaqqtxvm"))
                d_1056_v65_: _dafny.Map
                d_1056_v65_ = _dafny.Map({d_1026_v46_: d_1055_v64_})
                d_1053_v63_ = _dafny.Map({p0: ((d_1056_v65_)[d_1026_v46_] if (d_1026_v46_) in (d_1056_v65_) else d_1055_v64_)})
                d_1057_v66_: _dafny.Array
                nw163_ = _dafny.Array(_dafny.Array(None, 0), 24)
                d_1057_v66_ = nw163_
                d_1058_v67_: _dafny.Array
                def lambda44_(d_1059_i8_):
                    return default__.safeDivisionInt(d_1059_i8_, (self).f6)

                init27_ = lambda44_
                nw164_ = _dafny.Array(None, 7)
                for i0_27_ in range(nw164_.length(0)):
                    nw164_[i0_27_] = init27_(i0_27_)
                d_1058_v67_ = nw164_
                index166_ = default__.safeIndex(475, (d_1057_v66_).length(0))
                (d_1057_v66_)[index166_] = (d_1058_v67_ if d_1026_v46_ else d_1058_v67_)
                d_1060_v68_: _dafny.Seq
                d_1060_v68_ = _dafny.SeqWithoutIsStrInference([d_1026_v46_, d_1026_v46_, (p0) <= (p0)])
                index167_ = default__.safeIndex(475, (d_1057_v66_).length(0))
                rhs193_ = (d_1027_v47_)[default__.safeIndex((self).f17, len(d_1027_v47_))]
                rhs194_ = (self).f18
                rhs195_ = (d_1060_v68_)[default__.safeIndex((self).f6, len(d_1060_v68_))]
                rhs196_ = d_1058_v67_
                lhs159_ = globalState
                lhs160_ = globalState
                lhs161_ = d_1057_v66_
                lhs162_ = default__.safeIndex(475, (d_1057_v66_).length(0))
                lhs159_.f5 = rhs193_
                lhs160_.f5 = rhs194_
                d_1026_v46_ = rhs195_
                lhs161_[lhs162_] = rhs196_
            elif True:
                d_1061___mcc_h4_ = source13_.cf42
                d_1062_cf42_ = d_1061___mcc_h4_
                d_1063_v69_: _dafny.Seq
                d_1063_v69_ = _dafny.SeqWithoutIsStrInference([d_1026_v46_])
                d_1064_v70_: _dafny.Map
                d_1064_v70_ = _dafny.Map({(d_1063_v69_)[default__.safeIndex((self).f17, len(d_1063_v69_))]: not(d_1026_v46_)})
                d_1065_v71_: _dafny.Seq
                d_1065_v71_ = _dafny.SeqWithoutIsStrInference([d_1064_v70_])
                d_1065_v71_ = ((_dafny.SeqWithoutIsStrInference([d_1064_v70_, d_1064_v70_])) + (default__.fm46(globalState)) if d_1026_v46_ else (d_1065_v71_) + (d_1065_v71_))
                (globalState).f0 = False
                d_1066_v72_: _dafny.MultiSet
                d_1066_v72_ = _dafny.MultiSet([(self).f17])
                d_1067_v74_: _dafny.Array
                def lambda45_(d_1068_p0_):
                    def lambda46_(d_1069_i9_):
                        return (d_1069_i9_) + (d_1068_p0_)

                    return lambda46_

                init28_ = lambda45_(p0)
                nw165_ = _dafny.Array(None, 19)
                for i0_28_ in range(nw165_.length(0)):
                    nw165_[i0_28_] = init28_(i0_28_)
                d_1067_v74_ = nw165_
                d_1070_v75_: _dafny.Map
                d_1070_v75_ = _dafny.Map({d_1026_v46_: d_1067_v74_})
                pat_let_tv19_ = d_1066_v72_
                pat_let_tv20_ = d_1066_v72_
                d_1071_v76_: _dafny.Map
                def iife78_(_pat_let23_0):
                    def iife79_(d_1072_dt__update__tmp_h0_):
                        def iife81_():
                            coll32_ = _dafny.Set()
                            compr_32_: int
                            for compr_32_ in (pat_let_tv19_).Elements:
                                d_1073_v73_: int = compr_32_
                                if (d_1073_v73_) in (pat_let_tv20_):
                                    coll32_ = coll32_.union(_dafny.Set([default__.safeDivisionInt(d_1073_v73_, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True]))).cardinality)]))
                            return _dafny.Set(coll32_)
                        def iife80_(_pat_let24_0):
                            def iife82_(d_1074_dt__update_hcf57_h0_):
                                return D11_DC30(d_1074_dt__update_hcf57_h0_)
                            return iife82_(_pat_let24_0)
                        return iife80_(iife81_()
                        )
                    return iife79_(_pat_let23_0)
                d_1071_v76_ = _dafny.Map({iife78_(D11_DC30(_dafny.Set({(self).f17}))): d_1070_v75_})
                d_1075_v77_: _dafny.Map
                d_1075_v77_ = _dafny.Map({(self).f17: 931})
                d_1076_v78_: C3
                nw166_ = C3()
                nw166_.ctor__(d_1075_v77_, self.f9, (self).f10, p0)
                d_1076_v78_ = nw166_
                d_1077_v79_: _dafny.Seq
                d_1077_v79_ = _dafny.SeqWithoutIsStrInference([d_1076_v78_, d_1076_v78_])
                d_1078_v80_: _dafny.Map
                d_1078_v80_ = _dafny.Map({(d_1077_v79_)[default__.safeIndex((self).f6, len(d_1077_v79_))]: False})
                d_1079_v81_: _dafny.Set
                d_1079_v81_ = _dafny.Set({len(d_1078_v80_)})
                d_1080_v82_: D11
                d_1080_v82_ = D11_DC30(d_1079_v81_)
                d_1071_v76_ = (d_1071_v76_).set(d_1080_v82_, d_1070_v75_)
                rhs197_ = (d_1064_v70_).set(d_1026_v46_, (p0) <= (p0))
                rhs198_ = d_1026_v46_
                rhs199_ = default__.safeModuloInt(default__.safeDivisionInt((self).f18, (self).f6), (self).f6)
                lhs163_ = globalState
                lhs164_ = globalState
                d_1064_v70_ = rhs197_
                lhs163_.f0 = rhs198_
                lhs164_.f5 = rhs199_
            arr37_ = self.f9
            index168_ = default__.safeIndex(373, (self.f9).length(0))
            arr37_[index168_] = d_1026_v46_
            d_1081_v83_: _dafny.Set
            d_1081_v83_ = _dafny.Set({d_1026_v46_})
            arr38_ = self.f9
            index169_ = default__.safeIndex(373, (self.f9).length(0))
            arr38_[index169_] = (d_1081_v83_).ispropersubset(d_1081_v83_)
            arr39_ = self.f9
            index170_ = default__.safeIndex(373, (self.f9).length(0))
            arr39_[index170_] = (self.f9)[default__.safeIndex(373, (self.f9).length(0))]
            arr40_ = self.f9
            index171_ = default__.safeIndex(373, (self.f9).length(0))
            arr40_[index171_] = (self.f9)[default__.safeIndex(373, (self.f9).length(0))]
        d_1082_v84_: _dafny.Array
        def lambda47_(d_1083_i10_):
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rkntbsman"))

        init29_ = lambda47_
        nw167_ = _dafny.Array(None, 4)
        for i0_29_ in range(nw167_.length(0)):
            nw167_[i0_29_] = init29_(i0_29_)
        d_1082_v84_ = nw167_
        d_1084_v85_: bool
        d_1084_v85_ = True
        d_1085_v86_: D10
        d_1085_v86_ = D10_DC29(p0, d_1082_v84_, d_1084_v85_, (self).f18)
        source14_ = d_1085_v86_
        if source14_.is_DC29:
            d_1086___mcc_h5_ = source14_.cf53
            d_1087___mcc_h6_ = source14_.cf54
            d_1088___mcc_h7_ = source14_.cf55
            d_1089___mcc_h8_ = source14_.cf56
            d_1090_cf56_ = d_1089___mcc_h8_
            d_1091_cf55_ = d_1088___mcc_h7_
            d_1092_cf54_ = d_1087___mcc_h6_
            d_1093_cf53_ = d_1086___mcc_h5_
            (globalState).f2 = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "na"))
            arr41_ = self.f9
            index172_ = default__.safeIndex(104, (self.f9).length(0))
            arr41_[index172_] = d_1084_v85_
            d_1094_v87_: str
            d_1094_v87_ = _dafny.CodePoint('d')
            d_1095_v88_: _dafny.Set
            d_1095_v88_ = _dafny.Set({754, (self).f18})
            d_1096_v89_: _dafny.Array
            nw168_ = _dafny.Array(None, 21)
            nw168_[int(0)] = d_1094_v87_
            nw168_[int(1)] = d_1094_v87_
            nw168_[int(2)] = d_1094_v87_
            nw168_[int(3)] = d_1094_v87_
            nw168_[int(4)] = d_1094_v87_
            nw168_[int(5)] = default__.fm0((self).f18, d_1095_v88_, d_1084_v85_, globalState)
            nw168_[int(6)] = d_1094_v87_
            nw168_[int(7)] = d_1094_v87_
            nw168_[int(8)] = d_1094_v87_
            nw168_[int(9)] = d_1094_v87_
            nw168_[int(10)] = d_1094_v87_
            nw168_[int(11)] = d_1094_v87_
            nw168_[int(12)] = d_1094_v87_
            nw168_[int(13)] = d_1094_v87_
            nw168_[int(14)] = d_1094_v87_
            nw168_[int(15)] = d_1094_v87_
            nw168_[int(16)] = d_1094_v87_
            nw168_[int(17)] = d_1094_v87_
            nw168_[int(18)] = d_1094_v87_
            nw168_[int(19)] = d_1094_v87_
            nw168_[int(20)] = d_1094_v87_
            d_1096_v89_ = nw168_
            d_1097_v90_: _dafny.Seq
            d_1097_v90_ = _dafny.SeqWithoutIsStrInference([d_1096_v89_, d_1096_v89_, d_1096_v89_])
            arr42_ = self.f9
            index173_ = default__.safeIndex(104, (self.f9).length(0))
            arr42_[index173_] = ((d_1097_v90_)[default__.safeIndex((self).f18, len(d_1097_v90_))]) not in (_dafny.Map({d_1096_v89_: d_1094_v87_}))
            d_1098_v91_: _dafny.Seq
            d_1098_v91_ = _dafny.SeqWithoutIsStrInference([(self.f9)[default__.safeIndex(104, (self.f9).length(0))]])
            d_1099_v92_: _dafny.Seq
            d_1099_v92_ = _dafny.SeqWithoutIsStrInference([d_1098_v91_, d_1098_v91_])
            d_1100_v93_: _dafny.Map
            d_1100_v93_ = _dafny.Map({(0) - (len((d_1099_v92_) + (d_1099_v92_))): (self).f17})
            d_1100_v93_ = (d_1100_v93_).set(len(d_1095_v88_), p0)
            d_1101_v94_: D0
            d_1101_v94_ = D0_DC0(d_1098_v91_, d_1093_cf53_)
            d_1102_v95_: D1
            d_1102_v95_ = D1_DC4(d_1101_v94_, default__.safeModuloInt((self).f6, (self).f6), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pceere")), d_1091_cf55_, p0)
            source15_ = d_1102_v95_
            if source15_.is_DC4:
                d_1103___mcc_h10_ = source15_.cf7
                d_1104___mcc_h11_ = source15_.cf8
                d_1105___mcc_h12_ = source15_.cf9
                d_1106___mcc_h13_ = source15_.cf10
                d_1107___mcc_h14_ = source15_.cf11
                d_1108_cf11_ = d_1107___mcc_h14_
                d_1109_cf10_ = d_1106___mcc_h13_
                d_1110_cf9_ = d_1105___mcc_h12_
                d_1111_cf8_ = d_1104___mcc_h11_
                d_1112_cf7_ = d_1103___mcc_h10_
                d_1113_v96_: _dafny.Map
                d_1113_v96_ = _dafny.Map({False: d_1082_v84_})
                d_1114_v97_: _dafny.Array
                nw169_ = _dafny.Array(None, 21)
                nw169_[int(0)] = d_1092_cf54_
                nw169_[int(1)] = d_1082_v84_
                nw169_[int(2)] = ((d_1113_v96_)[(self.f9)[default__.safeIndex(104, (self.f9).length(0))]] if ((self.f9)[default__.safeIndex(104, (self.f9).length(0))]) in (d_1113_v96_) else d_1082_v84_)
                nw169_[int(3)] = d_1082_v84_
                nw169_[int(4)] = d_1092_cf54_
                nw169_[int(5)] = d_1082_v84_
                nw169_[int(6)] = d_1082_v84_
                nw169_[int(7)] = d_1082_v84_
                nw169_[int(8)] = d_1092_cf54_
                nw169_[int(9)] = d_1092_cf54_
                nw169_[int(10)] = d_1092_cf54_
                nw169_[int(11)] = d_1092_cf54_
                nw169_[int(12)] = (D10_DC29((self).f6, d_1082_v84_, d_1084_v85_, default__.fm2(534, globalState))).cf54
                nw169_[int(13)] = d_1082_v84_
                nw169_[int(14)] = d_1082_v84_
                nw169_[int(15)] = d_1092_cf54_
                nw169_[int(16)] = d_1082_v84_
                nw169_[int(17)] = d_1082_v84_
                nw169_[int(18)] = d_1082_v84_
                nw169_[int(19)] = d_1082_v84_
                nw169_[int(20)] = d_1082_v84_
                d_1114_v97_ = nw169_
                index174_ = default__.safeIndex(959, (d_1114_v97_).length(0))
                (d_1114_v97_)[index174_] = d_1092_cf54_
                index175_ = default__.safeIndex(959, (d_1114_v97_).length(0))
                (d_1114_v97_)[index175_] = d_1092_cf54_
                arr43_ = self.f9
                index176_ = default__.safeIndex(104, (self.f9).length(0))
                arr43_[index176_] = (_dafny.SeqWithoutIsStrInference([d_1098_v91_, _dafny.SeqWithoutIsStrInference([d_1109_cf10_, d_1091_cf55_])])) < (d_1099_v92_)
                d_1093_cf53_ = default__.fm2(615, globalState)
                d_1115_v98_: _dafny.Array
                nw170_ = _dafny.Array(_dafny.Set({}), 20)
                d_1115_v98_ = nw170_
                d_1116_v99_: _dafny.Array
                nw171_ = _dafny.Array(int(0), 13)
                d_1116_v99_ = nw171_
                rhs200_ = d_1115_v98_
                rhs201_ = d_1116_v99_
                rhs202_ = default__.fm25(d_1108_cf11_, ((self).f6) not in (_dafny.SeqWithoutIsStrInference([d_1108_cf11_ for d_1117_i11_ in range(default__.abs(-665))])), globalState)
                d_1115_v98_ = rhs200_
                d_1116_v99_ = rhs201_
                r2 = rhs202_
            elif True:
                d_1118___mcc_h15_ = source15_.cf6
                d_1119_cf6_ = d_1118___mcc_h15_
                d_1120_v100_: _dafny.Array
                nw172_ = _dafny.Array(False, 8)
                d_1120_v100_ = nw172_
                d_1121_v101_: C0
                nw173_ = C0()
                nw173_.ctor__(d_1120_v100_, not(not((self.f9)[default__.safeIndex(104, (self.f9).length(0))])))
                d_1121_v101_ = nw173_
                d_1122_v102_: C1
                nw174_ = C1()
                nw174_.ctor__((self).f6)
                d_1122_v102_ = nw174_
                d_1123_v103_: _dafny.Map
                d_1123_v103_ = _dafny.Map({187: d_1122_v102_})
                d_1124_v104_: _dafny.MultiSet
                d_1124_v104_ = _dafny.MultiSet([d_1090_cf56_, len(d_1123_v103_)])
                d_1124_v104_ = (d_1124_v104_).set((self).f18, default__.abs((self).f18))
                d_1125_v105_: _dafny.Seq
                d_1125_v105_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uyva"))
                d_1126_v106_: _dafny.MultiSet
                d_1126_v106_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jxfruork")), d_1125_v105_, d_1125_v105_])
                d_1127_v108_: D0
                d_1127_v108_ = D0_DC1(d_1121_v101_.f20, d_1090_cf56_, d_1094_v87_)
                d_1128_v109_: _dafny.Set
                d_1128_v109_ = _dafny.Set({d_1084_v85_, d_1091_cf55_})
                def iife83_():
                    coll33_ = _dafny.Map()
                    compr_33_: int
                    for compr_33_ in _dafny.IntegerRange(355, 690):
                        d_1129_v107_: int = compr_33_
                        if ((355) <= (d_1129_v107_)) and ((d_1129_v107_) < (690)):
                            coll33_[default__.safeModuloInt(d_1129_v107_, (self).f6)] = d_1084_v85_
                    return _dafny.Map(coll33_)
                d_1126_v106_ = ((default__.fm47(d_1094_v87_, iife83_()
                , d_1091_cf55_, globalState)) - (_dafny.MultiSet([d_1125_v105_]))) - ((d_1126_v106_).set(default__.fm30(d_1084_v85_, (d_1127_v108_).cf4, globalState), default__.abs(len(d_1128_v109_))))
                d_1094_v87_ = d_1094_v87_
        elif True:
            d_1130___mcc_h9_ = source14_.cf52
            d_1131_cf52_ = d_1130___mcc_h9_
            d_1132_v110_: C0
            nw175_ = C0()
            nw175_.ctor__(self.f9, d_1084_v85_)
            d_1132_v110_ = nw175_
            (globalState).f2 = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_1133_i12_ in range(default__.abs(-170))])
            d_1134_v111_: _dafny.Seq
            d_1134_v111_ = _dafny.SeqWithoutIsStrInference([d_1132_v110_.f20])
            d_1135_v112_: _dafny.MultiSet
            d_1135_v112_ = _dafny.MultiSet([(self).f18, default__.fm2(len(d_1134_v111_), globalState), (self).f17, (self).f18])
            d_1136_v113_: _dafny.Map
            d_1136_v113_ = _dafny.Map({d_1084_v85_: (d_1135_v112_).intersection(d_1135_v112_)})
            d_1137_v114_: _dafny.Seq
            d_1137_v114_ = _dafny.SeqWithoutIsStrInference([(self).f6, p0, (self).f17, (self).f17])
            d_1136_v113_ = (d_1136_v113_).set(d_1132_v110_.f20, _dafny.MultiSet((d_1137_v114_).set(default__.safeIndex((self).f18, len(d_1137_v114_)), (self).f17)))
            arr44_ = d_1132_v110_.f19
            index177_ = default__.safeIndex(998, (d_1132_v110_.f19).length(0))
            arr44_[index177_] = (d_1084_v85_ if d_1132_v110_.f20 else d_1084_v85_)
            d_1138_v115_: _dafny.Array
            def lambda48_(d_1139_i13_):
                return default__.safeModuloInt(d_1139_i13_, (self).f17)

            init30_ = lambda48_
            nw176_ = _dafny.Array(None, 1)
            for i0_30_ in range(nw176_.length(0)):
                nw176_[i0_30_] = init30_(i0_30_)
            d_1138_v115_ = nw176_
            index178_ = default__.safeIndex(256, (d_1138_v115_).length(0))
            (d_1138_v115_)[index178_] = 134
            d_1140_v116_: _dafny.Seq
            d_1140_v116_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tl"))
            d_1141_v117_: _dafny.Seq
            d_1141_v117_ = _dafny.SeqWithoutIsStrInference([d_1140_v116_])
            arr45_ = d_1132_v110_.f19
            index179_ = default__.safeIndex(998, (d_1132_v110_.f19).length(0))
            index180_ = default__.safeIndex(256, (d_1138_v115_).length(0))
            rhs203_ = not(d_1084_v85_)
            rhs204_ = default__.safeDivisionInt(p0, 367)
            rhs205_ = p0
            rhs206_ = (D4_DC13(d_1137_v114_, False, p0)).cf26
            rhs207_ = len((_dafny.SeqWithoutIsStrInference([(self).f17, len((d_1141_v117_)[default__.safeIndex(547, len(d_1141_v117_))]), (self).f17, (self).f17])) + (d_1137_v114_))
            lhs165_ = d_1132_v110_.f19
            lhs166_ = default__.safeIndex(998, (d_1132_v110_.f19).length(0))
            lhs167_ = globalState
            lhs168_ = d_1138_v115_
            lhs169_ = default__.safeIndex(256, (d_1138_v115_).length(0))
            lhs170_ = globalState
            lhs171_ = globalState
            lhs165_[lhs166_] = rhs203_
            lhs167_.f5 = rhs204_
            lhs168_[lhs169_] = rhs205_
            lhs170_.f0 = rhs206_
            lhs171_.f5 = rhs207_
        d_1142_v118_: str
        d_1142_v118_ = _dafny.CodePoint('i')
        d_1143_v119_: _dafny.Seq
        d_1143_v119_ = _dafny.SeqWithoutIsStrInference([d_1142_v118_])
        d_1144_i14_: int
        d_1144_i14_ = 0
        with _dafny.label("6"):
            while not(((self).f6) < (default__.safeDivisionInt(981, default__.fm2(len(d_1143_v119_), globalState)))):
                with _dafny.c_label("6"):
                    if (d_1144_i14_) >= (100):
                        raise _dafny.Break("6")
                    d_1144_i14_ = (d_1144_i14_) + (1)
                    (globalState).f5 = default__.safeDivisionInt((self).f17, default__.safeDivisionInt(len(d_1143_v119_), (self).f18))
                    d_1145_v121_: _dafny.Seq
                    d_1145_v121_ = _dafny.SeqWithoutIsStrInference([p0])
                    d_1146_v122_: _dafny.Set
                    d_1146_v122_ = _dafny.Set({len(d_1145_v121_), (self).f6})
                    d_1147_v123_: C2
                    nw177_ = C2()
                    def iife84_():
                        coll34_ = _dafny.Map()
                        compr_34_: int
                        for compr_34_ in (d_1146_v122_).Elements:
                            d_1148_v120_: int = compr_34_
                            if (d_1148_v120_) in (d_1146_v122_):
                                coll34_[default__.safeModuloInt(d_1148_v120_, 138)] = d_1143_v119_
                        return _dafny.Map(coll34_)
                    nw177_.ctor__(p0, self.f16, self.f9, iife84_()
                    )
                    d_1147_v123_ = nw177_
                    d_1146_v122_ = ((d_1146_v122_) - (d_1146_v122_) if not(d_1084_v85_) else (d_1146_v122_) | (d_1146_v122_))
                    d_1149_v124_: _dafny.Map
                    d_1149_v124_ = _dafny.Map({d_1084_v85_: d_1084_v85_})
                    d_1150_v125_: _dafny.MultiSet
                    d_1150_v125_ = _dafny.MultiSet([d_1084_v85_])
                    d_1151_v126_: _dafny.Map
                    d_1151_v126_ = _dafny.Map({(self).f18: (self).f18})
                    rhs208_ = ((p0) - (p0)) > ((self).f6)
                    rhs209_ = _dafny.Map({d_1084_v85_: (d_1147_v123_).fm22(d_1150_v125_, d_1084_v85_, len(d_1151_v126_), globalState)})
                    lhs172_ = globalState
                    lhs172_.f0 = rhs208_
                    d_1149_v124_ = rhs209_
                    pass
            pass
        d_1152_v127_: _dafny.MultiSet
        d_1152_v127_ = _dafny.MultiSet([d_1142_v118_])
        d_1153_v128_: _dafny.Map
        d_1153_v128_ = _dafny.Map({p0: d_1152_v127_})
        d_1154_v129_: _dafny.Array
        nw178_ = _dafny.Array(int(0), 8)
        d_1154_v129_ = nw178_
        d_1155_v130_: _dafny.MultiSet
        d_1155_v130_ = _dafny.MultiSet([d_1154_v129_, d_1154_v129_, d_1154_v129_, d_1154_v129_])
        d_1156_v131_: _dafny.Seq
        d_1156_v131_ = _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([p0, (d_1155_v130_).cardinality])).cardinality])
        d_1153_v128_ = (d_1153_v128_).set(len(d_1156_v131_), d_1152_v127_)
        d_1157_v132_: _dafny.Array
        nw179_ = _dafny.Array(_dafny.Seq({}), 17)
        d_1157_v132_ = nw179_
        d_1158_v133_: _dafny.Seq
        d_1158_v133_ = _dafny.SeqWithoutIsStrInference([d_1084_v85_])
        d_1159_v134_: _dafny.Seq
        d_1159_v134_ = _dafny.SeqWithoutIsStrInference([d_1158_v133_, d_1158_v133_, _dafny.SeqWithoutIsStrInference([d_1084_v85_])])
        index181_ = default__.safeIndex(924, (d_1157_v132_).length(0))
        (d_1157_v132_)[index181_] = d_1159_v134_
        index182_ = default__.safeIndex(924, (d_1157_v132_).length(0))
        (d_1157_v132_)[index182_] = d_1159_v134_
        d_1160_v135_: D0
        d_1160_v135_ = D0_DC1(d_1084_v85_, p0, (d_1143_v119_)[default__.safeIndex((self).f6, len(d_1143_v119_))])
        d_1161_v136_: _dafny.Map
        d_1161_v136_ = _dafny.Map({self.f9: d_1084_v85_})
        d_1162_v137_: _dafny.Map
        d_1162_v137_ = _dafny.Map({d_1084_v85_: (d_1161_v136_ if (d_1160_v135_).cf2 else d_1161_v136_)})
        d_1162_v137_ = (d_1162_v137_).set(not (d_1084_v85_) or (d_1084_v85_), d_1161_v136_)
        r0 = D2_DC8(True, not(((self).f6) < (default__.fm2(len(d_1143_v119_), globalState))))
        r1 = _dafny.SeqWithoutIsStrInference([not (default__.fm3(globalState)) or (False)])
        r2 = (d_1143_v119_) + (d_1143_v119_)
        return r0, r1, r2

    @property
    def f17(self):
        return self._f17
    @property
    def f18(self):
        return self._f18

class C5(T0):
    def  __init__(self):
        self._f6: int = int(0)
        self._f15: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    @property
    def f6(self):
        return self._f6
    def ctor__(self, f15, f6):
        (self)._f15 = f15
        (self)._f6 = f6

    def fm5(self, p0, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([_dafny.CodePoint('h')]), _dafny.MultiSet([_dafny.CodePoint('j'), _dafny.CodePoint('g')])])) + ((_dafny.SeqWithoutIsStrInference([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r'), _dafny.CodePoint('q'), _dafny.CodePoint('d')]))])) + (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([_dafny.CodePoint('c')])])))

    def fm6(self, p0, globalState):
        def iife85_():
            coll35_ = _dafny.Map()
            compr_35_: int
            for compr_35_ in _dafny.IntegerRange(94, 657):
                d_1165_v0_: int = compr_35_
                if ((94) <= (d_1165_v0_)) and ((d_1165_v0_) < (657)):
                    coll35_[(d_1165_v0_) + ((self).f6)] = _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([True])).cardinality])
            return _dafny.Map(coll35_)
        return ((_dafny.Map({(self).f15: _dafny.SeqWithoutIsStrInference([(self).f6])})) | (_dafny.Map({(D0_DC0(_dafny.SeqWithoutIsStrInference([False, (D2_DC8(not(False), False)).cf20]), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ig"))))).cf1: _dafny.SeqWithoutIsStrInference([(self).f15 for d_1163_i0_ in range(default__.abs(174))])}))) | ((_dafny.Map({(self).f15: _dafny.SeqWithoutIsStrInference([(D1_DC4(D0_DC0(_dafny.SeqWithoutIsStrInference([not(True), True]), (self).f6), (self).f6, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_1164_i1_ in range(default__.abs(108))]), False, (self).f15)).cf11, (self).f6])})) | (iife85_()
        ))

    def m10(self, p0, p1, globalState):
        r0: _dafny.Map = _dafny.Map({})
        r1: bool = False
        d_1166_v0_: _dafny.Array
        nw180_ = _dafny.Array(int(0), 18)
        d_1166_v0_ = nw180_
        d_1167_v1_: D2
        d_1167_v1_ = D2_DC6(p1, d_1166_v0_)
        source16_ = (D2_DC6((d_1167_v1_).cf13, d_1166_v0_) if p0 else d_1167_v1_)
        if source16_.is_DC6:
            d_1168___mcc_h0_ = source16_.cf13
            d_1169___mcc_h1_ = source16_.cf14
            d_1170_cf14_ = d_1169___mcc_h1_
            d_1171_cf13_ = d_1168___mcc_h0_
            d_1172_v2_: _dafny.Seq
            d_1172_v2_ = _dafny.SeqWithoutIsStrInference([(p1) != (p1), False, p0])
            d_1172_v2_ = ((_dafny.SeqWithoutIsStrInference([p0])) + (d_1172_v2_)) + (d_1172_v2_)
            (globalState).f0 = p0
            r1 = p0
            d_1173_v3_: str
            d_1173_v3_ = _dafny.CodePoint('m')
            d_1174_v4_: _dafny.MultiSet
            d_1174_v4_ = _dafny.MultiSet([d_1173_v3_, d_1173_v3_])
            d_1175_v5_: D1
            d_1175_v5_ = D1_DC3(d_1174_v4_)
            d_1176_v6_: _dafny.Set
            d_1176_v6_ = _dafny.Set({d_1175_v5_})
            d_1177_v7_: _dafny.Set
            d_1177_v7_ = _dafny.Set({d_1176_v6_, d_1176_v6_, d_1176_v6_, d_1176_v6_, d_1176_v6_})
            d_1178_v8_: _dafny.Map
            d_1178_v8_ = _dafny.Map({p0: d_1177_v7_})
            d_1178_v8_ = (d_1178_v8_).set(p0, d_1177_v7_)
        elif source16_.is_DC7:
            d_1179___mcc_h2_ = source16_.cf15
            d_1180___mcc_h3_ = source16_.cf16
            d_1181___mcc_h4_ = source16_.cf17
            d_1182___mcc_h5_ = source16_.cf18
            d_1183_cf18_ = d_1182___mcc_h5_
            d_1184_cf17_ = d_1181___mcc_h4_
            d_1185_cf16_ = d_1180___mcc_h3_
            d_1186_cf15_ = d_1179___mcc_h2_
            d_1187_v9_: _dafny.Seq
            d_1187_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y"))
            d_1188_v10_: _dafny.MultiSet
            d_1188_v10_ = _dafny.MultiSet([(d_1187_v9_) != (d_1187_v9_)])
            d_1188_v10_ = (d_1188_v10_) | (_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([p0])) + (d_1186_cf15_)))
            d_1187_v9_ = (d_1187_v9_) + (d_1187_v9_)
            d_1189_v11_: _dafny.Map
            d_1189_v11_ = _dafny.Map({p0: d_1187_v9_})
            d_1190_v12_: str
            d_1190_v12_ = _dafny.CodePoint('i')
            d_1191_v13_: _dafny.Map
            d_1191_v13_ = _dafny.Map({(self).f6: d_1187_v9_})
            d_1189_v11_ = (d_1189_v11_).set((d_1190_v12_) in (default__.fm18(p0, globalState)), ((d_1191_v13_)[(self).f6] if ((self).f6) in (d_1191_v13_) else d_1187_v9_))
            if p0:
                d_1192_v14_: _dafny.MultiSet
                d_1192_v14_ = _dafny.MultiSet([d_1183_cf18_, (self).f15])
                d_1193_v15_: _dafny.Map
                d_1193_v15_ = _dafny.Map({(d_1192_v14_).ispropersubset(default__.fm19(p1, globalState)): (self).f15})
                d_1193_v15_ = (d_1193_v15_).set((d_1187_v9_) <= ((d_1187_v9_).set(default__.safeIndex(len(d_1187_v9_), len(d_1187_v9_)), _dafny.CodePoint('x'))), (0) - ((self).f15))
                d_1194_v16_: _dafny.MultiSet
                d_1194_v16_ = _dafny.MultiSet([d_1190_v12_])
                d_1195_v17_: _dafny.Map
                d_1195_v17_ = _dafny.Map({((d_1194_v16_)[d_1190_v12_] if (d_1190_v12_) in (d_1194_v16_) else (self).f6): (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dgs"))) == (d_1187_v9_)})
                d_1196_v18_: _dafny.Set
                d_1196_v18_ = _dafny.Set({p0, p0})
                index183_ = default__.safeIndex(799, (d_1166_v0_).length(0))
                (d_1166_v0_)[index183_] = len((d_1196_v18_) | (d_1196_v18_))
                d_1197_v19_: _dafny.Map
                d_1197_v19_ = _dafny.Map({False: d_1190_v12_})
                d_1198_v20_: _dafny.Set
                d_1198_v20_ = _dafny.Set({58, (self).f6})
                d_1199_v21_: D0
                d_1199_v21_ = D0_DC1(not((((d_1197_v19_)[p0] if (p0) in (d_1197_v19_) else default__.fm0((self).f15, d_1198_v20_, not(p0), globalState))) not in (d_1187_v9_)), (0) - ((0) - ((0) - (d_1183_cf18_))), d_1190_v12_)
                d_1200_v22_: _dafny.Array
                nw181_ = _dafny.Array(False, 7)
                d_1200_v22_ = nw181_
                index184_ = default__.safeIndex(799, (d_1166_v0_).length(0))
                rhs210_ = d_1195_v17_
                rhs211_ = p1
                rhs212_ = default__.fm20(p0, (p0) or (p0), p0, globalState)
                rhs213_ = d_1200_v22_
                lhs173_ = d_1166_v0_
                lhs174_ = default__.safeIndex(799, (d_1166_v0_).length(0))
                d_1195_v17_ = rhs210_
                lhs173_[lhs174_] = rhs211_
                d_1199_v21_ = rhs212_
                d_1200_v22_ = rhs213_
                r1 = default__.fm3(globalState)
                d_1183_cf18_ = (((d_1166_v0_)[default__.safeIndex(799, (d_1166_v0_).length(0))]) * ((self).f6)) * (d_1183_cf18_)
                d_1201_v23_: _dafny.Array
                nw182_ = _dafny.Array(None, 11)
                nw182_[int(0)] = d_1186_cf15_
                nw182_[int(1)] = d_1186_cf15_
                nw182_[int(2)] = d_1186_cf15_
                nw182_[int(3)] = _dafny.SeqWithoutIsStrInference([p0])
                nw182_[int(4)] = default__.fm4((self).f6, p1, d_1183_cf18_, globalState)
                nw182_[int(5)] = d_1186_cf15_
                nw182_[int(6)] = d_1186_cf15_
                nw182_[int(7)] = _dafny.SeqWithoutIsStrInference([p0])
                nw182_[int(8)] = d_1186_cf15_
                nw182_[int(9)] = _dafny.SeqWithoutIsStrInference([p0, p0])
                nw182_[int(10)] = d_1186_cf15_
                d_1201_v23_ = nw182_
                d_1202_v24_: _dafny.Map
                d_1202_v24_ = _dafny.Map({default__.safeModuloInt(len(d_1198_v20_), len(d_1187_v9_)): d_1201_v23_})
                d_1202_v24_ = (d_1202_v24_).set((0) - ((self).f6), d_1201_v23_)
            elif True:
                d_1203_v25_: C1
                nw183_ = C1()
                nw183_.ctor__(-736)
                d_1203_v25_ = nw183_
                index185_ = default__.safeIndex(220, (d_1166_v0_).length(0))
                (d_1166_v0_)[index185_] = d_1184_cf17_
                index186_ = default__.safeIndex(220, (d_1166_v0_).length(0))
                (d_1166_v0_)[index186_] = d_1184_cf17_
                d_1190_v12_ = d_1190_v12_
                (globalState).f0 = p0
                r1 = False
        elif source16_.is_DC8:
            d_1204___mcc_h6_ = source16_.cf19
            d_1205___mcc_h7_ = source16_.cf20
            d_1206_cf20_ = d_1205___mcc_h7_
            d_1207_cf19_ = d_1204___mcc_h6_
            d_1208_v26_: _dafny.Array
            nw184_ = _dafny.Array(None, 5)
            nw184_[int(0)] = d_1207_cf19_
            nw184_[int(1)] = False
            nw184_[int(2)] = d_1207_cf19_
            nw184_[int(3)] = False
            nw184_[int(4)] = p0
            d_1208_v26_ = nw184_
            d_1209_v27_: _dafny.Seq
            d_1209_v27_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dusfea"))
            d_1210_v28_: _dafny.Seq
            d_1210_v28_ = _dafny.SeqWithoutIsStrInference([d_1209_v27_, d_1209_v27_])
            d_1211_v29_: _dafny.Map
            d_1211_v29_ = _dafny.Map({622: (d_1210_v28_)[default__.safeIndex(p1, len(d_1210_v28_))]})
            d_1212_v30_: C3
            nw185_ = C3()
            nw185_.ctor__(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wbwsr"))): p1}), (d_1208_v26_ if d_1206_cf20_ else d_1208_v26_), d_1211_v29_, (self).f6)
            d_1212_v30_ = nw185_
            d_1213_v31_: C1
            nw186_ = C1()
            nw186_.ctor__(p1)
            d_1213_v31_ = nw186_
            d_1214_v32_: _dafny.Map
            d_1214_v32_ = _dafny.Map({True: 427})
            d_1215_v33_: str
            d_1215_v33_ = _dafny.CodePoint('j')
            d_1216_v34_: _dafny.Map
            d_1216_v34_ = _dafny.Map({d_1215_v33_: p0})
            r0 = (d_1214_v32_).set(p0, len(d_1216_v34_))
            d_1217_v35_: _dafny.Seq
            d_1217_v35_ = _dafny.SeqWithoutIsStrInference([d_1207_cf19_, d_1207_cf19_])
            d_1218_v36_: _dafny.Map
            d_1218_v36_ = _dafny.Map({p0: d_1217_v35_})
            d_1219_v37_: _dafny.Map
            d_1219_v37_ = _dafny.Map({(self).f15: d_1206_cf20_})
            d_1220_v38_: _dafny.Seq
            d_1220_v38_ = _dafny.SeqWithoutIsStrInference([d_1219_v37_])
            d_1221_v39_: _dafny.Map
            d_1221_v39_ = _dafny.Map({(len(((d_1218_v36_)[not(True)] if (not(True)) in (d_1218_v36_) else (d_1217_v35_).set(default__.safeIndex((self).f15, len(d_1217_v35_)), p0)))) == (87): (d_1220_v38_) != (d_1220_v38_)})
            rhs214_ = d_1221_v39_
            rhs215_ = d_1166_v0_
            rhs216_ = p1
            lhs175_ = globalState
            d_1221_v39_ = rhs214_
            d_1166_v0_ = rhs215_
            lhs175_.f5 = rhs216_
        elif True:
            d_1222___mcc_h8_ = source16_.cf12
            d_1223_cf12_ = d_1222___mcc_h8_
            d_1224_v40_: _dafny.Set
            d_1224_v40_ = _dafny.Set({p0})
            d_1224_v40_ = d_1224_v40_
            d_1225_v41_: _dafny.MultiSet
            d_1225_v41_ = _dafny.MultiSet([p1])
            d_1226_v42_: str
            d_1226_v42_ = _dafny.CodePoint('d')
            d_1227_v43_: _dafny.Set
            d_1227_v43_ = _dafny.Set({default__.fm42(p0, d_1225_v41_, p1, d_1226_v42_, globalState)})
            d_1228_v44_: _dafny.Seq
            d_1228_v44_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s"))
            d_1229_v45_: _dafny.MultiSet
            d_1229_v45_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([(self).f15, len(d_1228_v44_)])])
            d_1230_v46_: _dafny.Map
            d_1230_v46_ = _dafny.Map({d_1227_v43_: ((d_1229_v45_)[(_dafny.SeqWithoutIsStrInference([default__.fm2(-138, globalState)])).set(default__.safeIndex(p1, len(_dafny.SeqWithoutIsStrInference([default__.fm2(-138, globalState)]))), (self).f6)] if ((_dafny.SeqWithoutIsStrInference([default__.fm2(-138, globalState)])).set(default__.safeIndex(p1, len(_dafny.SeqWithoutIsStrInference([default__.fm2(-138, globalState)]))), (self).f6)) in (d_1229_v45_) else len(d_1228_v44_))})
            def iife86_():
                coll36_ = _dafny.Set()
                compr_36_: _dafny.Map
                for compr_36_ in (_dafny.Map({d_1223_cf12_: p1})).keys.Elements:
                    d_1231_v47_: _dafny.Map = compr_36_
                    if (d_1231_v47_) in (_dafny.Map({d_1223_cf12_: p1})):
                        coll36_ = coll36_.union(_dafny.Set([d_1231_v47_]))
                return _dafny.Set(coll36_)
            d_1230_v46_ = (d_1230_v46_).set((d_1227_v43_).intersection(iife86_()
            ), (self).f6)
            d_1232_v48_: _dafny.Seq
            d_1232_v48_ = _dafny.SeqWithoutIsStrInference([p0])
            d_1233_v49_: _dafny.Map
            d_1233_v49_ = _dafny.Map({d_1232_v48_: p0})
            d_1234_v50_: _dafny.Set
            d_1234_v50_ = _dafny.Set({d_1166_v0_})
            d_1235_v51_: _dafny.Map
            d_1235_v51_ = _dafny.Map({d_1233_v49_: d_1234_v50_})
            d_1235_v51_ = (d_1235_v51_).set(_dafny.Map({_dafny.SeqWithoutIsStrInference([p0]): default__.fm3(globalState)}), d_1234_v50_)
            d_1236_v52_: _dafny.Map
            d_1236_v52_ = _dafny.Map({p0: ((self).f15) == (p1)})
            d_1236_v52_ = (d_1236_v52_).set(p0, p0)
        d_1237_v53_: _dafny.Array
        nw187_ = _dafny.Array(False, 2)
        d_1237_v53_ = nw187_
        d_1238_v54_: _dafny.Array
        d_1238_v54_ = d_1237_v53_
        d_1239_v55_: _dafny.Map
        d_1239_v55_ = _dafny.Map({d_1238_v54_: p0})
        index187_ = default__.safeIndex(609, (d_1166_v0_).length(0))
        (d_1166_v0_)[index187_] = len(d_1239_v55_)
        index188_ = default__.safeIndex(609, (d_1166_v0_).length(0))
        (d_1166_v0_)[index188_] = p1
        d_1240_v56_: _dafny.Map
        d_1240_v56_ = _dafny.Map({d_1166_v0_: ((d_1166_v0_)[default__.safeIndex(609, (d_1166_v0_).length(0))]) * ((d_1166_v0_)[default__.safeIndex(609, (d_1166_v0_).length(0))])})
        d_1241_v57_: _dafny.Map
        d_1241_v57_ = _dafny.Map({p0: d_1166_v0_})
        d_1240_v56_ = (d_1240_v56_).set(((d_1241_v57_)[False] if (False) in (d_1241_v57_) else d_1166_v0_), (self).f15)
        d_1242_v58_: str
        d_1242_v58_ = _dafny.CodePoint('n')
        d_1243_v59_: _dafny.Map
        d_1243_v59_ = _dafny.Map({d_1242_v58_: d_1237_v53_})
        d_1243_v59_ = (d_1243_v59_).set(d_1242_v58_, d_1237_v53_)
        d_1244_v60_: _dafny.Map
        d_1244_v60_ = _dafny.Map({p0: (self).f15})
        d_1245_v61_: _dafny.Seq
        d_1245_v61_ = _dafny.SeqWithoutIsStrInference([len(d_1244_v60_), (self).f15])
        d_1246_v62_: _dafny.Seq
        d_1246_v62_ = _dafny.SeqWithoutIsStrInference([p0])
        d_1247_v63_: _dafny.Map
        d_1247_v63_ = _dafny.Map({(self).f6: (d_1246_v62_)[default__.safeIndex(len(default__.fm1(globalState)), len(d_1246_v62_))]})
        d_1248_v64_: _dafny.Map
        d_1248_v64_ = _dafny.Map({d_1247_v63_: p0})
        r1 = ((len(d_1245_v61_)) < (len(default__.fm33(D3_DC10(d_1248_v64_), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_1249_i0_ in range(default__.abs(863))])), not(p0), ((d_1247_v63_)[(d_1166_v0_)[default__.safeIndex(609, (d_1166_v0_).length(0))]] if ((d_1166_v0_)[default__.safeIndex(609, (d_1166_v0_).length(0))]) in (d_1247_v63_) else False), globalState))) if p0 else p0)
        d_1250_v65_: _dafny.Set
        d_1250_v65_ = _dafny.Set({False, p0})
        d_1251_v66_: D11
        d_1251_v66_ = D11_DC33((self).f15, p0, (self).f15, p0, d_1250_v65_)
        d_1252_v67_: _dafny.MultiSet
        d_1252_v67_ = _dafny.MultiSet([(d_1251_v66_).cf64, (self).f15, 112, (self).f6, 477])
        index189_ = default__.safeIndex(609, (d_1166_v0_).length(0))
        (d_1166_v0_)[index189_] = ((p1) - ((_dafny.MultiSet([(d_1166_v0_)[default__.safeIndex(609, (d_1166_v0_).length(0))]])).cardinality)) * (len(default__.fm42(p0, d_1252_v67_, (d_1252_v67_).cardinality, d_1242_v58_, globalState)))
        r0 = (D14_DC36(d_1244_v60_)).cf69
        r1 = p0
        return r0, r1

    def m11(self, p0, p1, p2, globalState):
        r0: int = int(0)
        r1: _dafny.Seq = _dafny.Seq({})
        r2: bool = False
        r3: int = int(0)
        d_1253_v0_: _dafny.Seq
        d_1253_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sc"))
        d_1254_v1_: _dafny.Seq
        d_1254_v1_ = _dafny.SeqWithoutIsStrInference([(self).f15])
        d_1255_v2_: str
        d_1255_v2_ = _dafny.CodePoint('o')
        d_1256_v3_: D9
        d_1256_v3_ = D9_DC25((self).f6, (self).f6, (self).f15, d_1255_v2_)
        d_1257_v4_: D9
        d_1257_v4_ = D9_DC27(d_1256_v3_)
        d_1258_v5_: D9
        d_1258_v5_ = D9_DC27(d_1257_v4_)
        d_1259_v6_: _dafny.Array
        nw188_ = _dafny.Array(None, 6)
        nw188_[int(0)] = D9_DC27(D9_DC25(len(d_1253_v0_), len(d_1254_v1_), (self).f6, d_1255_v2_))
        nw188_[int(1)] = d_1258_v5_
        nw188_[int(2)] = d_1258_v5_
        nw188_[int(3)] = d_1258_v5_
        nw188_[int(4)] = d_1258_v5_
        nw188_[int(5)] = d_1258_v5_
        d_1259_v6_ = nw188_
        index190_ = default__.safeIndex(307, (d_1259_v6_).length(0))
        (d_1259_v6_)[index190_] = d_1258_v5_
        d_1260_v7_: _dafny.MultiSet
        d_1260_v7_ = _dafny.MultiSet([True])
        d_1261_v8_: _dafny.Map
        d_1261_v8_ = _dafny.Map({p0: d_1260_v7_})
        d_1262_v9_: _dafny.Seq
        d_1262_v9_ = _dafny.SeqWithoutIsStrInference([(d_1260_v7_ if p1 else ((d_1261_v8_)[p1] if (p1) in (d_1261_v8_) else d_1260_v7_))])
        index191_ = default__.safeIndex(307, (d_1259_v6_).length(0))
        rhs217_ = ((self).f6) + (10)
        rhs218_ = d_1258_v5_
        rhs219_ = (d_1262_v9_)[default__.safeIndex(((self).f15) + ((d_1254_v1_)[default__.safeIndex((self).f6, len(d_1254_v1_))]), len(d_1262_v9_))]
        lhs176_ = d_1259_v6_
        lhs177_ = default__.safeIndex(307, (d_1259_v6_).length(0))
        r3 = rhs217_
        lhs176_[lhs177_] = rhs218_
        d_1260_v7_ = rhs219_
        hi9_ = (d_1254_v1_)[default__.safeIndex((self).f15, len(d_1254_v1_))]
        for d_1263_i0_ in range(417, hi9_):
            r0 = ((0) - ((self).f15)) - (((d_1260_v7_)[p0] if (p0) in (d_1260_v7_) else (self).f6))
            d_1264_v10_: _dafny.Map
            d_1264_v10_ = _dafny.Map({(863) == ((0) - ((self).f15)): p1})
            d_1264_v10_ = (d_1264_v10_).set(p1, default__.fm3(globalState))
            d_1265_v11_: _dafny.MultiSet
            d_1265_v11_ = _dafny.MultiSet([len(d_1253_v0_)])
            d_1266_v12_: _dafny.MultiSet
            d_1266_v12_ = _dafny.MultiSet([d_1265_v11_, d_1265_v11_])
            d_1267_v13_: D4
            d_1267_v13_ = D4_DC13(d_1254_v1_, p0, (d_1254_v1_)[default__.safeIndex((self).f15, len(d_1254_v1_))])
            d_1268_v14_: _dafny.Array
            nw189_ = _dafny.Array(None, 18)
            nw189_[int(0)] = p1
            nw189_[int(1)] = p1
            nw189_[int(2)] = p0
            nw189_[int(3)] = p0
            nw189_[int(4)] = p0
            nw189_[int(5)] = p0
            nw189_[int(6)] = p0
            nw189_[int(7)] = (d_1267_v13_).cf26
            nw189_[int(8)] = p1
            nw189_[int(9)] = False
            nw189_[int(10)] = p1
            nw189_[int(11)] = p0
            nw189_[int(12)] = True
            nw189_[int(13)] = p1
            nw189_[int(14)] = p1
            nw189_[int(15)] = p1
            nw189_[int(16)] = p1
            nw189_[int(17)] = not(p0)
            d_1268_v14_ = nw189_
            d_1269_v15_: _dafny.Map
            d_1269_v15_ = _dafny.Map({d_1255_v2_: d_1268_v14_})
            d_1270_v16_: C2
            nw190_ = C2()
            nw190_.ctor__((self).f6, d_1266_v12_, ((d_1269_v15_)[d_1255_v2_] if (d_1255_v2_) in (d_1269_v15_) else d_1268_v14_), _dafny.Map({(self).f15: default__.fm18(p1, globalState)}))
            d_1270_v16_ = nw190_
            (globalState).f0 = False
        source17_ = D5_DC17(d_1254_v1_, p0)
        if source17_.is_DC17:
            d_1271___mcc_h0_ = source17_.cf36
            d_1272___mcc_h1_ = source17_.cf37
            d_1273_cf37_ = d_1272___mcc_h1_
            d_1274_cf36_ = d_1271___mcc_h0_
            d_1275_v17_: _dafny.Array
            def lambda49_(d_1276_p1_):
                def lambda50_(d_1277_i1_):
                    return (True if d_1276_p1_ else d_1276_p1_)

                return lambda50_

            init31_ = lambda49_(p1)
            nw191_ = _dafny.Array(None, 21)
            for i0_31_ in range(nw191_.length(0)):
                nw191_[i0_31_] = init31_(i0_31_)
            d_1275_v17_ = nw191_
            index192_ = default__.safeIndex(166, (d_1275_v17_).length(0))
            (d_1275_v17_)[index192_] = not(p0)
            index193_ = default__.safeIndex(166, (d_1275_v17_).length(0))
            (d_1275_v17_)[index193_] = p0
            d_1278_v18_: _dafny.Seq
            d_1278_v18_ = _dafny.SeqWithoutIsStrInference([d_1273_cf37_])
            d_1279_v19_: _dafny.Map
            d_1279_v19_ = _dafny.Map({(self).f15: (d_1275_v17_)[default__.safeIndex(166, (d_1275_v17_).length(0))]})
            d_1280_v20_: _dafny.Set
            d_1280_v20_ = _dafny.Set({len(d_1279_v19_), 667, (self).f6})
            d_1281_v21_: _dafny.Seq
            d_1281_v21_ = _dafny.SeqWithoutIsStrInference([d_1280_v20_, d_1280_v20_])
            d_1282_v22_: _dafny.Array
            nw192_ = _dafny.Array(None, 22)
            nw192_[int(0)] = len(_dafny.Map({d_1273_cf37_: True}))
            nw192_[int(1)] = (self).f15
            nw192_[int(2)] = len((d_1278_v18_).set(default__.safeIndex(len(d_1274_cf36_), len(d_1278_v18_)), p0))
            nw192_[int(3)] = (self).f15
            nw192_[int(4)] = len((d_1281_v21_)[default__.safeIndex((self).f6, len(d_1281_v21_))])
            nw192_[int(5)] = len(d_1253_v0_)
            nw192_[int(6)] = (self).f6
            nw192_[int(7)] = default__.fm2((self).f6, globalState)
            nw192_[int(8)] = (self).f15
            nw192_[int(9)] = (self).f6
            nw192_[int(10)] = (self).f6
            nw192_[int(11)] = (self).f15
            nw192_[int(12)] = (0) - ((self).f15)
            nw192_[int(13)] = (self).f6
            nw192_[int(14)] = (0) - ((self).f15)
            nw192_[int(15)] = (d_1254_v1_)[default__.safeIndex((self).f15, len(d_1254_v1_))]
            nw192_[int(16)] = default__.fm2((self).f6, globalState)
            nw192_[int(17)] = 697
            nw192_[int(18)] = (self).f15
            nw192_[int(19)] = (self).f15
            nw192_[int(20)] = 576
            nw192_[int(21)] = (self).f15
            d_1282_v22_ = nw192_
            d_1283_v23_: _dafny.Map
            d_1283_v23_ = _dafny.Map({d_1282_v22_: (d_1275_v17_)[default__.safeIndex(166, (d_1275_v17_).length(0))]})
            d_1273_cf37_ = not(((d_1283_v23_)[d_1282_v22_] if (d_1282_v22_) in (d_1283_v23_) else p0))
            r2 = (d_1253_v0_) <= (d_1253_v0_)
            d_1279_v19_ = (d_1279_v19_).set((self).f6, (d_1275_v17_)[default__.safeIndex(166, (d_1275_v17_).length(0))])
        elif True:
            d_1284___mcc_h2_ = source17_.cf35
            d_1285_cf35_ = d_1284___mcc_h2_
            d_1286_v24_: _dafny.Array
            nw193_ = _dafny.Array(False, 17)
            d_1286_v24_ = nw193_
            d_1286_v24_ = d_1286_v24_
            (globalState).f5 = (self).f15
            d_1287_v25_: _dafny.Seq
            d_1287_v25_ = _dafny.SeqWithoutIsStrInference([p1])
            d_1288_v26_: _dafny.Map
            d_1288_v26_ = _dafny.Map({d_1287_v25_: (self).f15})
            d_1288_v26_ = (d_1288_v26_).set(d_1287_v25_, -172)
            r2 = not (not(p1)) or (((self).f15) == ((self).f15))
        d_1289_i2_: int
        d_1289_i2_ = 0
        with _dafny.label("7"):
            while p0:
                with _dafny.c_label("7"):
                    if (d_1289_i2_) >= (100):
                        raise _dafny.Break("7")
                    d_1289_i2_ = (d_1289_i2_) + (1)
                    d_1290_v27_: _dafny.Seq
                    d_1290_v27_ = _dafny.SeqWithoutIsStrInference([(p1 if p0 else p1)])
                    d_1291_v28_: _dafny.Map
                    d_1291_v28_ = _dafny.Map({(self).f15: d_1290_v27_})
                    d_1290_v27_ = (_dafny.SeqWithoutIsStrInference([default__.fm3(globalState)])) + (((d_1291_v28_)[(self).f6] if ((self).f6) in (d_1291_v28_) else d_1290_v27_))
                    d_1292_v29_: _dafny.Array
                    def lambda51_(d_1293_i3_):
                        return default__.safeModuloInt(d_1293_i3_, (self).f15)

                    init32_ = lambda51_
                    nw194_ = _dafny.Array(None, 14)
                    for i0_32_ in range(nw194_.length(0)):
                        nw194_[i0_32_] = init32_(i0_32_)
                    d_1292_v29_ = nw194_
                    d_1294_v30_: D2
                    d_1294_v30_ = D2_DC6((self).f6, d_1292_v29_)
                    d_1295_v31_: _dafny.Seq
                    d_1295_v31_ = _dafny.SeqWithoutIsStrInference([d_1292_v29_, d_1292_v29_])
                    pat_let_tv21_ = d_1292_v29_
                    pat_let_tv22_ = d_1292_v29_
                    d_1296_v32_: _dafny.Array
                    nw195_ = _dafny.Array(None, 11)
                    nw195_[int(0)] = d_1292_v29_
                    nw195_[int(1)] = d_1292_v29_
                    nw195_[int(2)] = d_1292_v29_
                    nw195_[int(3)] = d_1292_v29_
                    nw195_[int(4)] = d_1292_v29_
                    def iife88_(_pat_let26_0):
                        def iife89_(d_1297_dt__update__tmp_h0_):
                            def iife90_(_pat_let27_0):
                                def iife91_(d_1298_dt__update_hcf14_h0_):
                                    return D2_DC6((d_1297_dt__update__tmp_h0_).cf13, d_1298_dt__update_hcf14_h0_)
                                return iife91_(_pat_let27_0)
                            return iife90_(pat_let_tv21_)
                        return iife89_(_pat_let26_0)
                    def iife87_(_pat_let25_0):
                        def iife92_(d_1299_dt__update__tmp_h1_):
                            def iife93_(_pat_let28_0):
                                def iife94_(d_1300_dt__update_hcf14_h1_):
                                    return D2_DC6((d_1299_dt__update__tmp_h1_).cf13, d_1300_dt__update_hcf14_h1_)
                                return iife94_(_pat_let28_0)
                            return iife93_(pat_let_tv22_)
                        return iife92_(_pat_let25_0)
                    nw195_[int(5)] = (iife87_(iife88_(d_1294_v30_))).cf14
                    nw195_[int(6)] = (d_1295_v31_)[default__.safeIndex(((d_1260_v7_).set(p0, default__.abs((self).f15))).cardinality, len(d_1295_v31_))]
                    nw195_[int(7)] = d_1292_v29_
                    nw195_[int(8)] = d_1292_v29_
                    nw195_[int(9)] = d_1292_v29_
                    nw195_[int(10)] = d_1292_v29_
                    d_1296_v32_ = nw195_
                    index194_ = default__.safeIndex(790, (d_1296_v32_).length(0))
                    (d_1296_v32_)[index194_] = d_1292_v29_
                    index195_ = default__.safeIndex(790, (d_1296_v32_).length(0))
                    (d_1296_v32_)[index195_] = d_1292_v29_
                    (globalState).f5 = (self).f15
                    d_1301_v33_: C1
                    nw196_ = C1()
                    nw196_.ctor__((self).f15)
                    d_1301_v33_ = nw196_
                    pass
            pass
        d_1302_v34_: _dafny.Seq
        d_1302_v34_ = _dafny.SeqWithoutIsStrInference([p0])
        d_1303_v35_: _dafny.Seq
        d_1303_v35_ = _dafny.SeqWithoutIsStrInference([(d_1302_v34_)[default__.safeIndex((0) - (len(_dafny.Set({p1}))), len(d_1302_v34_))], False])
        d_1304_v37_: _dafny.Map
        def iife95_():
            coll37_ = _dafny.Map()
            compr_37_: int
            for compr_37_ in (d_1254_v1_).Elements:
                d_1305_v36_: int = compr_37_
                if (d_1305_v36_) in (d_1254_v1_):
                    coll37_[(d_1305_v36_) * ((0) - ((self).f15))] = (self).f6
            return _dafny.Map(coll37_)
        d_1304_v37_ = _dafny.Map({len(iife95_()
        ): p0})
        d_1306_v38_: _dafny.Map
        d_1306_v38_ = _dafny.Map({d_1304_v37_: p0})
        d_1307_v39_: D3
        d_1307_v39_ = D3_DC10(d_1306_v38_)
        pat_let_tv23_ = d_1306_v38_
        d_1308_v40_: _dafny.Array
        nw197_ = _dafny.Array(None, 3)
        def iife96_(_pat_let29_0):
            def iife97_(d_1309_dt__update__tmp_h2_):
                def iife98_(_pat_let30_0):
                    def iife99_(d_1310_dt__update_hcf22_h0_):
                        return D3_DC10(d_1310_dt__update_hcf22_h0_)
                    return iife99_(_pat_let30_0)
                return iife98_(pat_let_tv23_)
            return iife97_(_pat_let29_0)
        nw197_[int(0)] = (d_1307_v39_ if (d_1302_v34_)[default__.safeIndex((self).f15, len(d_1302_v34_))] else iife96_(d_1307_v39_))
        nw197_[int(1)] = D3_DC10(d_1306_v38_)
        nw197_[int(2)] = d_1307_v39_
        d_1308_v40_ = nw197_
        index196_ = default__.safeIndex(680, (d_1308_v40_).length(0))
        (d_1308_v40_)[index196_] = d_1307_v39_
        index197_ = default__.safeIndex(680, (d_1308_v40_).length(0))
        (d_1308_v40_)[index197_] = (d_1307_v39_ if (p1) == (p1) else d_1307_v39_)
        if p1:
            if True:
                (globalState).f2 = default__.fm30(p1, d_1255_v2_, globalState)
                d_1311_v41_: _dafny.Array
                def lambda52_(d_1312_p1_):
                    def lambda53_(d_1313_i4_):
                        return not(d_1312_p1_)

                    return lambda53_

                init33_ = lambda52_(p1)
                nw198_ = _dafny.Array(None, 10)
                for i0_33_ in range(nw198_.length(0)):
                    nw198_[i0_33_] = init33_(i0_33_)
                d_1311_v41_ = nw198_
                d_1314_v42_: _dafny.Map
                d_1314_v42_ = _dafny.Map({(self).f15: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "amxb"))})
                d_1315_v43_: _dafny.MultiSet
                d_1315_v43_ = _dafny.MultiSet([_dafny.MultiSet((d_1254_v1_).set(default__.safeIndex((self).f6, len(d_1254_v1_)), 680))])
                d_1316_v44_: C4
                nw199_ = C4()
                nw199_.ctor__(default__.fm2(default__.fm2((self).f6, globalState), globalState), (self).f15, (d_1311_v41_ if p0 else d_1311_v41_), (d_1314_v42_) | ((_dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_1255_v2_ for d_1317_i5_ in range(default__.abs(141))])): d_1253_v0_})).set(default__.fm2((d_1260_v7_).cardinality, globalState), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "coc")))), (self).f6, d_1315_v43_)
                d_1316_v44_ = nw199_
                (globalState).f5 = default__.safeModuloInt(len(d_1254_v1_), (self).f6)
                d_1318_v45_: D0
                d_1318_v45_ = D0_DC1(p1, (d_1316_v44_).f18, d_1255_v2_)
                d_1319_v46_: _dafny.Map
                d_1319_v46_ = _dafny.Map({p0: False})
                d_1320_v47_: D4
                d_1320_v47_ = D4_DC14(p1, d_1319_v46_, (self).f15, d_1319_v46_)
                d_1321_v48_: _dafny.Seq
                d_1321_v48_ = _dafny.SeqWithoutIsStrInference([d_1320_v47_, D4_DC14(p1, d_1319_v46_, (self).f6, d_1319_v46_)])
                d_1322_v49_: _dafny.Seq
                d_1322_v49_ = _dafny.SeqWithoutIsStrInference([d_1321_v48_, _dafny.SeqWithoutIsStrInference([d_1320_v47_ for d_1323_i6_ in range(default__.abs(710))]), d_1321_v48_, ((d_1321_v48_).set(default__.safeIndex((self).f6, len(d_1321_v48_)), d_1320_v47_)).set(default__.safeIndex((self).f6, len((d_1321_v48_).set(default__.safeIndex((self).f6, len(d_1321_v48_)), d_1320_v47_))), D4_DC14(p1, d_1319_v46_, (0) - ((d_1316_v44_).f17), _dafny.Map({p0: p0}))), d_1321_v48_])
                rhs220_ = (d_1318_v45_).cf2
                rhs221_ = (len(d_1322_v49_)) * ((self).f6)
                rhs222_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kx"))) != (d_1253_v0_)) or (p0)
                rhs223_ = _dafny.CodePoint('t')
                lhs178_ = globalState
                lhs179_ = globalState
                lhs180_ = globalState
                lhs178_.f0 = rhs220_
                lhs179_.f5 = rhs221_
                lhs180_.f0 = rhs222_
                d_1255_v2_ = rhs223_
                d_1324_v50_: _dafny.Set
                d_1324_v50_ = _dafny.Set({p0})
                d_1325_v51_: C1
                nw200_ = C1()
                nw200_.ctor__(default__.safeModuloInt(548, (d_1316_v44_).fm21(len(d_1324_v50_), globalState)))
                d_1325_v51_ = nw200_
            elif True:
                (globalState).f0 = p0
                d_1326_v52_: _dafny.MultiSet
                d_1326_v52_ = _dafny.MultiSet([28, (self).f15])
                d_1327_v53_: _dafny.MultiSet
                d_1327_v53_ = _dafny.MultiSet([d_1326_v52_])
                d_1328_v54_: _dafny.Array
                def lambda54_(d_1329_i7_):
                    return True

                init34_ = lambda54_
                nw201_ = _dafny.Array(None, 28)
                for i0_34_ in range(nw201_.length(0)):
                    nw201_[i0_34_] = init34_(i0_34_)
                d_1328_v54_ = nw201_
                d_1330_v55_: _dafny.Map
                d_1330_v55_ = _dafny.Map({(0) - ((self).f6): d_1253_v0_})
                d_1331_v57_: C2
                nw202_ = C2()
                def iife100_():
                    coll38_ = _dafny.Map()
                    compr_38_: int
                    for compr_38_ in _dafny.IntegerRange(-58, 555):
                        d_1332_v56_: int = compr_38_
                        if ((-58) <= (d_1332_v56_)) and ((d_1332_v56_) < (555)):
                            coll38_[default__.safeModuloInt(d_1332_v56_, len(_dafny.Map({False: (d_1326_v52_).cardinality})))] = d_1253_v0_
                    return _dafny.Map(coll38_)
                nw202_.ctor__(-676, (d_1327_v53_) | (_dafny.MultiSet([d_1326_v52_, d_1326_v52_])), d_1328_v54_, (d_1330_v55_) | (iife100_()
                ))
                d_1331_v57_ = nw202_
                d_1333_v59_: _dafny.Array
                nw203_ = _dafny.Array(None, 26)
                nw203_[int(0)] = (d_1260_v7_).cardinality
                nw203_[int(1)] = (self).f6
                nw203_[int(2)] = (self).f15
                nw203_[int(3)] = (self).f6
                nw203_[int(4)] = (self).f15
                nw203_[int(5)] = (self).f15
                nw203_[int(6)] = (self).f15
                nw203_[int(7)] = (self).f6
                nw203_[int(8)] = len(d_1253_v0_)
                nw203_[int(9)] = (self).f6
                nw203_[int(10)] = (self).f15
                nw203_[int(11)] = (self).f6
                nw203_[int(12)] = (d_1326_v52_).cardinality
                nw203_[int(13)] = 99
                nw203_[int(14)] = (self).f15
                nw203_[int(15)] = -832
                nw203_[int(16)] = (self).f6
                nw203_[int(17)] = (self).f6
                nw203_[int(18)] = (d_1331_v57_).fm21((self).f15, globalState)
                nw203_[int(19)] = len(p2)
                nw203_[int(20)] = (self).f6
                nw203_[int(21)] = (self).f15
                nw203_[int(22)] = (self).f15
                nw203_[int(23)] = (self).f15
                nw203_[int(24)] = (self).f6
                nw203_[int(25)] = (self).f15
                d_1333_v59_ = nw203_
                d_1334_v60_: D2
                d_1334_v60_ = D2_DC6(len((d_1253_v0_).set(default__.safeIndex((self).f15, len(d_1253_v0_)), d_1255_v2_)), d_1333_v59_)
                d_1335_v61_: C2
                nw204_ = C2()
                def iife101_():
                    coll39_ = _dafny.Map()
                    compr_39_: int
                    for compr_39_ in (default__.fm35((d_1334_v60_).cf13, (self).f6, p1, (self).f6, globalState)).Elements:
                        d_1336_v58_: int = compr_39_
                        if (d_1336_v58_) in (default__.fm35((d_1334_v60_).cf13, (self).f6, p1, (self).f6, globalState)):
                            coll39_[(d_1336_v58_) + (len(d_1302_v34_))] = d_1253_v0_
                    return _dafny.Map(coll39_)
                nw204_.ctor__((self).f15, d_1327_v53_, d_1328_v54_, iife101_()
                )
                d_1335_v61_ = nw204_
                d_1337_v62_: _dafny.Map
                d_1337_v62_ = _dafny.Map({651: d_1333_v59_})
                d_1337_v62_ = (d_1337_v62_).set((self).f6, d_1333_v59_)
                d_1338_v63_: _dafny.Map
                d_1338_v63_ = _dafny.Map({d_1255_v2_: p0})
                d_1339_v64_: _dafny.Set
                d_1339_v64_ = _dafny.Set({p0, p1})
                d_1338_v63_ = (d_1338_v63_).set(d_1255_v2_, ((_dafny.SeqWithoutIsStrInference([(d_1335_v61_).fm21((self).f6, globalState), (default__.fm35((self).f6, (self).f6, p1, (self).f6, globalState)).cardinality, (self).f6, 991, (self).f15])).set(default__.safeIndex(len(default__.fm40((self).f6, d_1339_v64_, globalState)), len(_dafny.SeqWithoutIsStrInference([(d_1335_v61_).fm21((self).f6, globalState), (default__.fm35((self).f6, (self).f6, p1, (self).f6, globalState)).cardinality, (self).f6, 991, (self).f15]))), 409)) != (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tvotwjin"))) for d_1340_i8_ in range(default__.abs(761))])))
            d_1341_v65_: _dafny.Seq
            d_1341_v65_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1255_v2_ for d_1342_i9_ in range(default__.abs(-371))]), d_1253_v0_, _dafny.SeqWithoutIsStrInference([d_1255_v2_ for d_1343_i10_ in range(default__.abs(714))])])
            d_1341_v65_ = (d_1341_v65_ if (p1) or (p1) else d_1341_v65_)
            d_1344_v66_: _dafny.Array
            nw205_ = _dafny.Array(False, 6)
            d_1344_v66_ = nw205_
            index198_ = default__.safeIndex(861, (d_1344_v66_).length(0))
            (d_1344_v66_)[index198_] = (d_1260_v7_).issubset(default__.fm48(globalState))
            d_1345_v67_: _dafny.Array
            nw206_ = _dafny.Array(D1.default()(), 11)
            d_1345_v67_ = nw206_
            d_1346_v68_: _dafny.Map
            d_1346_v68_ = _dafny.Map({d_1345_v67_: p0})
            index199_ = default__.safeIndex(861, (d_1344_v66_).length(0))
            rhs224_ = p1
            rhs225_ = ((self).f15) > ((self).f15)
            rhs226_ = d_1346_v68_
            rhs227_ = (self).f15
            rhs228_ = (d_1303_v35_)[default__.safeIndex((self).f6, len(d_1303_v35_))]
            lhs181_ = d_1344_v66_
            lhs182_ = default__.safeIndex(861, (d_1344_v66_).length(0))
            lhs183_ = globalState
            lhs181_[lhs182_] = rhs224_
            r2 = rhs225_
            d_1346_v68_ = rhs226_
            r3 = rhs227_
            lhs183_.f0 = rhs228_
            d_1347_v69_: _dafny.Map
            d_1347_v69_ = _dafny.Map({p0: (self).f15})
            d_1348_v70_: _dafny.Map
            d_1348_v70_ = _dafny.Map({d_1344_v66_: (((d_1347_v69_)[True] if (True) in (d_1347_v69_) else (self).f6)) > (default__.fm2((self).f15, globalState))})
            d_1348_v70_ = (_dafny.Map({d_1344_v66_: (d_1344_v66_)[default__.safeIndex(861, (d_1344_v66_).length(0))]})) | (d_1348_v70_)
            d_1349_v71_: _dafny.MultiSet
            d_1349_v71_ = _dafny.MultiSet([(self).f15, (self).f15])
            d_1350_v72_: _dafny.MultiSet
            d_1350_v72_ = _dafny.MultiSet([d_1349_v71_, _dafny.MultiSet(d_1254_v1_)])
            d_1351_v73_: _dafny.Map
            d_1351_v73_ = _dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sigwl")))): d_1253_v0_})
            d_1352_v74_: _dafny.Seq
            d_1352_v74_ = _dafny.SeqWithoutIsStrInference([d_1351_v73_])
            d_1353_v75_: C2
            nw207_ = C2()
            nw207_.ctor__((self).f15, d_1350_v72_, d_1344_v66_, (d_1352_v74_)[default__.safeIndex((self).f15, len(d_1352_v74_))])
            d_1353_v75_ = nw207_
        elif True:
            d_1354_v76_: _dafny.Set
            d_1354_v76_ = _dafny.Set({d_1255_v2_})
            d_1355_v78_: _dafny.Seq
            def iife102_():
                coll40_ = _dafny.Set()
                compr_40_: str
                for compr_40_ in (_dafny.Map({d_1255_v2_: p0})).keys.Elements:
                    d_1356_v77_: str = compr_40_
                    if (d_1356_v77_) in (_dafny.Map({d_1255_v2_: p0})):
                        coll40_ = coll40_.union(_dafny.Set([d_1356_v77_]))
                return _dafny.Set(coll40_)
            d_1355_v78_ = _dafny.SeqWithoutIsStrInference([d_1354_v76_, iife102_()
            ])
            (globalState).f5 = default__.safeDivisionInt(len(d_1355_v78_), (988 if p0 else len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_1357_i11_ in range(default__.abs(699))]))))
            d_1358_v79_: _dafny.Map
            d_1358_v79_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference([-145 for d_1359_i12_ in range(default__.abs(554))])) != (d_1254_v1_): p0})
            d_1358_v79_ = (d_1358_v79_).set(p0, p0)
            (globalState).f0 = p1
            d_1260_v7_ = (d_1260_v7_) | (_dafny.MultiSet([p0]))
            d_1360_v80_: _dafny.Set
            d_1360_v80_ = _dafny.Set({(self).f15})
            d_1361_v81_: _dafny.Array
            def lambda55_(d_1362_i13_):
                return (d_1362_i13_) - ((0) - ((self).f15))

            init35_ = lambda55_
            nw208_ = _dafny.Array(None, 5)
            for i0_35_ in range(nw208_.length(0)):
                nw208_[i0_35_] = init35_(i0_35_)
            d_1361_v81_ = nw208_
            d_1363_v82_: _dafny.Map
            d_1363_v82_ = _dafny.Map({(self).f15: d_1361_v81_})
            d_1364_v83_: _dafny.MultiSet
            d_1364_v83_ = _dafny.MultiSet([(0) - ((self).f6)])
            d_1365_v84_: _dafny.MultiSet
            d_1365_v84_ = _dafny.MultiSet([(d_1254_v1_)[default__.safeIndex(default__.fm2((d_1364_v83_).cardinality, globalState), len(d_1254_v1_))]])
            d_1366_v85_: _dafny.Map
            d_1366_v85_ = _dafny.Map({p1: _dafny.CodePoint('a')})
            d_1367_v86_: _dafny.Array
            nw209_ = _dafny.Array(None, 16)
            nw209_[int(0)] = (self).f15
            nw209_[int(1)] = len(d_1360_v80_)
            nw209_[int(2)] = (self).f15
            nw209_[int(3)] = (self).f6
            nw209_[int(4)] = len(d_1363_v82_)
            nw209_[int(5)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "loqmtebnn")))
            nw209_[int(6)] = (self).f15
            nw209_[int(7)] = (self).f6
            nw209_[int(8)] = (self).f15
            nw209_[int(9)] = (self).f6
            nw209_[int(10)] = len(d_1253_v0_)
            nw209_[int(11)] = len(_dafny.Map({p0: p1}))
            nw209_[int(12)] = (self).f6
            nw209_[int(13)] = (self).f15
            nw209_[int(14)] = ((d_1365_v84_)[len(d_1366_v85_)] if (len(d_1366_v85_)) in (d_1365_v84_) else (0) - (len(_dafny.SeqWithoutIsStrInference([p0]))))
            nw209_[int(15)] = (self).f15
            d_1367_v86_ = nw209_
            d_1368_v87_: D2
            d_1368_v87_ = D2_DC6((len(_dafny.SeqWithoutIsStrInference([d_1303_v35_]))) + ((self).f15), d_1367_v86_)
            source18_ = d_1368_v87_
            if source18_.is_DC6:
                d_1369___mcc_h3_ = source18_.cf13
                d_1370___mcc_h4_ = source18_.cf14
                d_1371_cf14_ = d_1370___mcc_h4_
                d_1372_cf13_ = d_1369___mcc_h3_
                (globalState).f5 = (default__.safeModuloInt((self).f6, len(d_1253_v0_))) * ((self).f6)
                r3 = ((self).f6) + ((0) - (((self).f15) + (337)))
                (globalState).f2 = d_1253_v0_
                (globalState).f0 = (p0 if p0 else ((self).f15) >= ((self).f15))
            elif source18_.is_DC7:
                d_1373___mcc_h5_ = source18_.cf15
                d_1374___mcc_h6_ = source18_.cf16
                d_1375___mcc_h7_ = source18_.cf17
                d_1376___mcc_h8_ = source18_.cf18
                d_1377_cf18_ = d_1376___mcc_h8_
                d_1378_cf17_ = d_1375___mcc_h7_
                d_1379_cf16_ = d_1374___mcc_h6_
                d_1380_cf15_ = d_1373___mcc_h5_
                (globalState).f0 = (p1 if (p0 if p0 else p0) else (d_1380_cf15_)[default__.safeIndex((self).f15, len(d_1380_cf15_))])
                d_1381_v88_: D0
                d_1381_v88_ = D0_DC1(not(p0), 88, _dafny.CodePoint('f'))
                d_1382_v89_: _dafny.Map
                d_1382_v89_ = _dafny.Map({p0: (d_1260_v7_).cardinality})
                pat_let_tv24_ = d_1255_v2_
                d_1383_v90_: _dafny.Array
                nw210_ = _dafny.Array(None, 26)
                nw210_[int(0)] = d_1381_v88_
                nw210_[int(1)] = d_1381_v88_
                nw210_[int(2)] = d_1381_v88_
                nw210_[int(3)] = d_1381_v88_
                nw210_[int(4)] = (d_1381_v88_ if p1 else D0_DC1(p1, d_1377_cf18_, d_1255_v2_))
                nw210_[int(5)] = d_1381_v88_
                nw210_[int(6)] = d_1381_v88_
                nw210_[int(7)] = default__.fm20(p1, not(p0), True, globalState)
                def iife103_(_pat_let31_0):
                    def iife104_(d_1384_dt__update__tmp_h3_):
                        def iife105_(_pat_let32_0):
                            def iife106_(d_1385_dt__update_hcf4_h0_):
                                return D0_DC1((d_1384_dt__update__tmp_h3_).cf2, (d_1384_dt__update__tmp_h3_).cf3, d_1385_dt__update_hcf4_h0_)
                            return iife106_(_pat_let32_0)
                        return iife105_(pat_let_tv24_)
                    return iife104_(_pat_let31_0)
                nw210_[int(8)] = iife103_(d_1381_v88_)
                nw210_[int(9)] = d_1381_v88_
                nw210_[int(10)] = d_1381_v88_
                nw210_[int(11)] = D0_DC1((d_1303_v35_)[default__.safeIndex(((d_1382_v89_)[p1] if (p1) in (d_1382_v89_) else (self).f6), len(d_1303_v35_))], (self).f15, d_1255_v2_)
                nw210_[int(12)] = d_1381_v88_
                nw210_[int(13)] = (d_1381_v88_ if p1 else d_1381_v88_)
                nw210_[int(14)] = default__.fm20(p0, p1, p1, globalState)
                nw210_[int(15)] = d_1381_v88_
                nw210_[int(16)] = d_1381_v88_
                nw210_[int(17)] = d_1381_v88_
                nw210_[int(18)] = D0_DC1(p1, d_1378_cf17_, d_1255_v2_)
                nw210_[int(19)] = d_1381_v88_
                nw210_[int(20)] = d_1381_v88_
                nw210_[int(21)] = d_1381_v88_
                nw210_[int(22)] = d_1381_v88_
                nw210_[int(23)] = d_1381_v88_
                nw210_[int(24)] = d_1381_v88_
                nw210_[int(25)] = d_1381_v88_
                d_1383_v90_ = nw210_
                d_1383_v90_ = d_1383_v90_
                d_1386_v91_: _dafny.Array
                nw211_ = _dafny.Array(None, 8)
                nw211_[int(0)] = True
                nw211_[int(1)] = not(p1)
                nw211_[int(2)] = p0
                nw211_[int(3)] = p0
                nw211_[int(4)] = not(False)
                nw211_[int(5)] = p1
                nw211_[int(6)] = p1
                nw211_[int(7)] = p1
                d_1386_v91_ = nw211_
                d_1387_v92_: C0
                nw212_ = C0()
                nw212_.ctor__(d_1386_v91_, True)
                d_1387_v92_ = nw212_
                (globalState).f2 = (d_1253_v0_) + ((d_1253_v0_) + (d_1253_v0_))
            elif source18_.is_DC8:
                d_1388___mcc_h9_ = source18_.cf19
                d_1389___mcc_h10_ = source18_.cf20
                d_1390_cf20_ = d_1389___mcc_h10_
                d_1391_cf19_ = d_1388___mcc_h9_
                d_1255_v2_ = _dafny.CodePoint('l')
                (globalState).f5 = (self).f15
                d_1392_v93_: _dafny.Seq
                d_1392_v93_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({(self).f6, len(d_1303_v35_)})])
                d_1393_v94_: _dafny.Map
                d_1393_v94_ = _dafny.Map({d_1253_v0_: d_1256_v3_})
                pat_let_tv25_ = d_1393_v94_
                pat_let_tv26_ = d_1253_v0_
                pat_let_tv27_ = d_1253_v0_
                pat_let_tv28_ = d_1393_v94_
                pat_let_tv29_ = d_1256_v3_
                index200_ = default__.safeIndex(307, (d_1259_v6_).length(0))
                def iife107_(_pat_let33_0):
                    def iife108_(d_1394_dt__update__tmp_h4_):
                        def iife109_(_pat_let34_0):
                            def iife110_(d_1395_dt__update_hcf51_h0_):
                                return D9_DC27(d_1395_dt__update_hcf51_h0_)
                            return iife110_(_pat_let34_0)
                        return iife109_(((pat_let_tv25_)[pat_let_tv26_] if (pat_let_tv27_) in (pat_let_tv28_) else D9_DC27(pat_let_tv29_)))
                    return iife108_(_pat_let33_0)
                rhs229_ = (d_1392_v93_)[default__.safeIndex((self).f6, len(d_1392_v93_))]
                rhs230_ = iife107_(d_1258_v5_)
                rhs231_ = p0
                lhs184_ = d_1259_v6_
                lhs185_ = default__.safeIndex(307, (d_1259_v6_).length(0))
                lhs186_ = globalState
                d_1360_v80_ = rhs229_
                lhs184_[lhs185_] = rhs230_
                lhs186_.f0 = rhs231_
                d_1396_v95_: D9
                d_1396_v95_ = D9_DC25((self).f15, (self).f15, (self).f6, d_1255_v2_)
                d_1397_v96_: _dafny.Map
                d_1397_v96_ = _dafny.Map({(d_1368_v87_).cf14: d_1396_v95_})
                d_1397_v96_ = (d_1397_v96_).set(d_1367_v86_, d_1396_v95_)
            elif True:
                d_1398___mcc_h11_ = source18_.cf12
                d_1399_cf12_ = d_1398___mcc_h11_
                d_1400_v97_: _dafny.Set
                d_1400_v97_ = _dafny.Set({d_1399_cf12_})
                d_1400_v97_ = d_1400_v97_
                d_1401_v98_: C1
                nw213_ = C1()
                nw213_.ctor__(((self).f6) - ((self).f15))
                d_1401_v98_ = nw213_
                rhs232_ = p0
                rhs233_ = True
                lhs187_ = globalState
                r2 = rhs232_
                lhs187_.f0 = rhs233_
                d_1361_v81_ = d_1367_v86_
        r0 = (default__.safeModuloInt(len(d_1253_v0_), 352)) * ((self).f15)
        r1 = d_1254_v1_
        r2 = not (p0) or (((self).f6) > (len(_dafny.SeqWithoutIsStrInference([(self).f6, 895, (self).f6, (self).f6, (self).f6]))))
        d_1402_v99_: _dafny.Array
        nw214_ = _dafny.Array(_dafny.MultiSet({}), 19)
        d_1402_v99_ = nw214_
        d_1403_v100_: D2
        d_1403_v100_ = D2_DC7(d_1302_v34_, d_1402_v99_, len(d_1253_v0_), (0) - ((self).f6))
        r3 = (d_1403_v100_).cf18
        return r0, r1, r2, r3

    @property
    def f15(self):
        return self._f15

class C6:
    def  __init__(self):
        self._f13: int = int(0)
        self._f14: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C6"
    def ctor__(self, f13, f14):
        (self)._f13 = f13
        (self)._f14 = f14

    def fm15(self, p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference([903])) + ((_dafny.SeqWithoutIsStrInference([(self).f13 for d_1404_i0_ in range(default__.abs(-992))])) + (_dafny.SeqWithoutIsStrInference([741 for d_1405_i1_ in range(default__.abs(20))])))

    def fm16(self, p0, p1, globalState):
        return ((self).f13) * (len(_dafny.SeqWithoutIsStrInference([_dafny.Map({(self).f14: (self).f14})])))

    def m8(self, globalState):
        r0: int = int(0)
        r1: D2 = D2.default()()
        r2: bool = False
        d_1406_v0_: _dafny.Map
        d_1406_v0_ = _dafny.Map({(self).f13: (self).f13})
        d_1407_v1_: _dafny.Map
        d_1407_v1_ = _dafny.Map({((d_1406_v0_)[(self).f13] if ((self).f13) in (d_1406_v0_) else 956): (self).f14})
        d_1407_v1_ = (d_1407_v1_).set((self).f13, (self).f14)
        (globalState).f0 = default__.fm3(globalState)
        d_1408_i0_: int
        d_1408_i0_ = 0
        with _dafny.label("8"):
            while (self).f14:
                with _dafny.c_label("8"):
                    if (d_1408_i0_) >= (100):
                        raise _dafny.Break("8")
                    d_1408_i0_ = (d_1408_i0_) + (1)
                    d_1409_v2_: _dafny.Array
                    nw215_ = _dafny.Array(False, 25)
                    d_1409_v2_ = nw215_
                    d_1410_v3_: _dafny.MultiSet
                    d_1410_v3_ = _dafny.MultiSet([(self).f13])
                    d_1411_v4_: _dafny.Seq
                    d_1411_v4_ = _dafny.SeqWithoutIsStrInference([(d_1410_v3_).cardinality, (self).f13])
                    d_1412_v5_: _dafny.Map
                    d_1412_v5_ = _dafny.Map({d_1409_v2_: d_1411_v4_})
                    d_1412_v5_ = (d_1412_v5_).set(d_1409_v2_, d_1411_v4_)
                    d_1413_v6_: _dafny.Map
                    d_1413_v6_ = _dafny.Map({(self).f14: (0) - (-212)})
                    d_1413_v6_ = (d_1413_v6_).set((self).f14, (self).f13)
                    r2 = (self).f14
                    if (self).f14:
                        (globalState).f0 = not(False)
                        (globalState).f5 = (self).f13
                        d_1414_v7_: _dafny.Seq
                        d_1414_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cfytvh"))
                        d_1415_v8_: _dafny.Seq
                        d_1415_v8_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cevep")), (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_1416_i1_ in range(default__.abs(447))]) if (self).f14 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o"))), (d_1414_v7_) + (d_1414_v7_)])
                        (globalState).f2 = (d_1415_v8_)[default__.safeIndex(((0) - ((self).f13)) + (97), len(d_1415_v8_))]
                        d_1417_v9_: str
                        d_1417_v9_ = _dafny.CodePoint('e')
                        d_1417_v9_ = d_1417_v9_
                        (globalState).f5 = (self).f13
                    elif True:
                        d_1418_v10_: _dafny.Set
                        d_1418_v10_ = _dafny.Set({d_1409_v2_, d_1409_v2_, d_1409_v2_, d_1409_v2_, d_1409_v2_})
                        d_1418_v10_ = d_1418_v10_
                        d_1419_v11_: _dafny.Seq
                        d_1419_v11_ = _dafny.SeqWithoutIsStrInference([(self).f14, (self).f14])
                        d_1420_v12_: _dafny.Seq
                        d_1420_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "snuxhfgl"))
                        d_1421_v13_: _dafny.Set
                        d_1421_v13_ = _dafny.Set({d_1420_v12_})
                        d_1422_v14_: D0
                        d_1422_v14_ = D0_DC0(d_1419_v11_, len(d_1421_v13_))
                        d_1423_v15_: D1
                        d_1423_v15_ = D1_DC4(d_1422_v14_, 923, default__.fm17((self).f13, (self).f13, globalState), (self).f14, (self).f13)
                        (globalState).f2 = (d_1423_v15_).cf9
                        d_1424_v16_: _dafny.Array
                        nw216_ = _dafny.Array(int(0), 12)
                        d_1424_v16_ = nw216_
                        d_1425_v17_: _dafny.MultiSet
                        d_1425_v17_ = _dafny.MultiSet([d_1424_v16_, d_1424_v16_, d_1424_v16_, d_1424_v16_])
                        d_1426_v18_: _dafny.Seq
                        d_1426_v18_ = _dafny.SeqWithoutIsStrInference([d_1411_v4_])
                        d_1427_v19_: str
                        d_1427_v19_ = _dafny.CodePoint('f')
                        d_1428_v20_: _dafny.Map
                        d_1428_v20_ = _dafny.Map({d_1427_v19_: d_1424_v16_})
                        d_1429_v21_: _dafny.Seq
                        def iife111_(_pat_let35_0):
                            def iife112_(d_1430_dt__update__tmp_h0_):
                                def iife113_(_pat_let36_0):
                                    def iife114_(d_1431_dt__update_hcf1_h0_):
                                        return D0_DC0((d_1430_dt__update__tmp_h0_).cf0, d_1431_dt__update_hcf1_h0_)
                                    return iife114_(_pat_let36_0)
                                return iife113_((self).f13)
                            return iife112_(_pat_let35_0)
                        d_1429_v21_ = _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([d_1424_v16_])).set(((d_1428_v20_)[_dafny.CodePoint('e')] if (_dafny.CodePoint('e')) in (d_1428_v20_) else d_1424_v16_), default__.abs(len(_dafny.SeqWithoutIsStrInference([iife111_(d_1422_v14_), d_1422_v14_, d_1422_v14_, d_1422_v14_])))), d_1425_v17_])
                        d_1425_v17_ = ((((d_1425_v17_).set(d_1424_v16_, default__.abs(774))).set(d_1424_v16_, default__.abs(len((d_1426_v18_)[default__.safeIndex((self).f13, len(d_1426_v18_))])))) | (d_1425_v17_)) | ((d_1425_v17_) - ((d_1429_v21_)[default__.safeIndex((self).f13, len(d_1429_v21_))]))
                        index201_ = default__.safeIndex(101, (d_1424_v16_).length(0))
                        (d_1424_v16_)[index201_] = (d_1410_v3_).cardinality
                        d_1432_v22_: _dafny.Set
                        d_1432_v22_ = _dafny.Set({len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vp"))).set(default__.safeIndex((self).f13, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vp")))), _dafny.CodePoint('q')))})
                        index202_ = default__.safeIndex(101, (d_1424_v16_).length(0))
                        rhs234_ = default__.safeDivisionInt((self).f13, len(d_1432_v22_))
                        rhs235_ = (0) - (((self).f13) * (((d_1410_v3_)[(self).f13] if ((self).f13) in (d_1410_v3_) else (self).f13)))
                        lhs188_ = d_1424_v16_
                        lhs189_ = default__.safeIndex(101, (d_1424_v16_).length(0))
                        lhs190_ = globalState
                        lhs188_[lhs189_] = rhs234_
                        lhs190_.f5 = rhs235_
                        index203_ = default__.safeIndex(101, (d_1424_v16_).length(0))
                        (d_1424_v16_)[index203_] = (self).f13
                    pass
            pass
        hi10_ = len(_dafny.Map({(self).f14: (self).f14}))
        for d_1433_i2_ in range((self).f13, hi10_):
            d_1434_v23_: _dafny.Map
            d_1434_v23_ = _dafny.Map({not((self).f14): (self).f13})
            d_1435_v24_: D2
            d_1435_v24_ = D2_DC5(_dafny.Map({((d_1434_v23_)[(self).f14] if ((self).f14) in (d_1434_v23_) else (0) - (d_1433_i2_)): (0) - (d_1433_i2_)}))
            source19_ = d_1435_v24_
            if source19_.is_DC6:
                d_1436___mcc_h0_ = source19_.cf13
                d_1437___mcc_h1_ = source19_.cf14
                d_1438_cf14_ = d_1437___mcc_h1_
                d_1439_cf13_ = d_1436___mcc_h0_
                d_1440_v25_: _dafny.Map
                d_1440_v25_ = _dafny.Map({(self).f14: (self).f14})
                d_1440_v25_ = _dafny.Map({(self).f14: default__.fm3(globalState)})
                d_1441_v26_: _dafny.Seq
                d_1441_v26_ = _dafny.SeqWithoutIsStrInference([d_1434_v23_, d_1434_v23_, _dafny.Map({not((self).f14): d_1439_cf13_})])
                d_1442_v27_: _dafny.Seq
                d_1442_v27_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qk"))
                d_1443_v28_: _dafny.Seq
                d_1443_v28_ = _dafny.SeqWithoutIsStrInference([(self).f14])
                r2 = ((d_1441_v26_).set(default__.safeIndex(len(d_1442_v27_), len(d_1441_v26_)), ((_dafny.Map({(self).f14: len(d_1443_v28_)})).set((self).f14, 83)).set((self).f14, d_1439_cf13_))) < (d_1441_v26_)
                d_1440_v25_ = (d_1440_v25_).set(True, ((d_1440_v25_)[(self).f14] if ((self).f14) in (d_1440_v25_) else not((self).f14)))
                (globalState).f0 = (self).f14
            elif source19_.is_DC7:
                d_1444___mcc_h2_ = source19_.cf15
                d_1445___mcc_h3_ = source19_.cf16
                d_1446___mcc_h4_ = source19_.cf17
                d_1447___mcc_h5_ = source19_.cf18
                d_1448_cf18_ = d_1447___mcc_h5_
                d_1449_cf17_ = d_1446___mcc_h4_
                d_1450_cf16_ = d_1445___mcc_h3_
                d_1451_cf15_ = d_1444___mcc_h2_
                d_1452_v29_: C1
                nw217_ = C1()
                nw217_.ctor__((self).f13)
                d_1452_v29_ = nw217_
                d_1453_v30_: _dafny.Seq
                d_1453_v30_ = _dafny.SeqWithoutIsStrInference([d_1451_cf15_, d_1451_cf15_, d_1451_cf15_, d_1451_cf15_, d_1451_cf15_])
                d_1453_v30_ = d_1453_v30_
                (globalState).f5 = (self).f13
                (globalState).f5 = (self).f13
            elif source19_.is_DC8:
                d_1454___mcc_h6_ = source19_.cf19
                d_1455___mcc_h7_ = source19_.cf20
                d_1456_cf20_ = d_1455___mcc_h7_
                d_1457_cf19_ = d_1454___mcc_h6_
                (globalState).f5 = default__.safeModuloInt(default__.safeDivisionInt((self).f13, (self).f13), (self).f13)
                d_1458_v31_: _dafny.Array
                nw218_ = _dafny.Array(_dafny.Array(None, 0), 13)
                d_1458_v31_ = nw218_
                d_1459_v32_: _dafny.Array
                nw219_ = _dafny.Array(int(0), 12)
                d_1459_v32_ = nw219_
                d_1460_v33_: _dafny.MultiSet
                d_1460_v33_ = _dafny.MultiSet([d_1459_v32_, d_1459_v32_])
                d_1461_v34_: str
                d_1461_v34_ = _dafny.CodePoint('g')
                d_1462_v35_: _dafny.MultiSet
                d_1462_v35_ = _dafny.MultiSet([d_1461_v34_])
                d_1463_v36_: _dafny.Seq
                d_1463_v36_ = _dafny.SeqWithoutIsStrInference([(self).f14, default__.fm3(globalState)])
                d_1464_v37_: _dafny.Array
                nw220_ = _dafny.Array(None, 9)
                nw220_[int(0)] = d_1434_v23_
                nw220_[int(1)] = d_1434_v23_
                nw220_[int(2)] = _dafny.Map({d_1457_cf19_: d_1433_i2_})
                nw220_[int(3)] = _dafny.Map({d_1456_cf20_: (d_1460_v33_).cardinality})
                nw220_[int(4)] = _dafny.Map({(self).f14: (d_1462_v35_).cardinality})
                nw220_[int(5)] = (_dafny.Map({d_1457_cf19_: (self).f13})).set(d_1457_cf19_, (0) - (d_1433_i2_))
                nw220_[int(6)] = d_1434_v23_
                nw220_[int(7)] = default__.fm31(d_1457_cf19_, (self).f13, d_1433_i2_, d_1463_v36_, globalState)
                nw220_[int(8)] = d_1434_v23_
                d_1464_v37_ = nw220_
                index204_ = default__.safeIndex(990, (d_1458_v31_).length(0))
                (d_1458_v31_)[index204_] = d_1464_v37_
                index205_ = default__.safeIndex(990, (d_1458_v31_).length(0))
                (d_1458_v31_)[index205_] = d_1464_v37_
                (globalState).f5 = (0) - (((self).f13) - (((self).f13) - ((self).f13)))
                d_1465_v38_: _dafny.Array
                def lambda56_(d_1466_i3_):
                    return (_dafny.CodePoint('c')) in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tsfvbtav")))

                init36_ = lambda56_
                nw221_ = _dafny.Array(None, 9)
                for i0_36_ in range(nw221_.length(0)):
                    nw221_[i0_36_] = init36_(i0_36_)
                d_1465_v38_ = nw221_
                d_1467_v39_: _dafny.Seq
                d_1467_v39_ = _dafny.SeqWithoutIsStrInference([d_1465_v38_])
                d_1465_v38_ = (d_1467_v39_)[default__.safeIndex((-871) - (len(default__.fm4(d_1433_i2_, (self).f13, (self).f13, globalState))), len(d_1467_v39_))]
            elif True:
                d_1468___mcc_h8_ = source19_.cf12
                d_1469_cf12_ = d_1468___mcc_h8_
                d_1470_v40_: _dafny.Array
                def lambda57_(d_1471_i2_):
                    def lambda58_(d_1472_i4_):
                        return default__.safeDivisionInt(d_1472_i4_, (0) - (d_1471_i2_))

                    return lambda58_

                init37_ = lambda57_(d_1433_i2_)
                nw222_ = _dafny.Array(None, 22)
                for i0_37_ in range(nw222_.length(0)):
                    nw222_[i0_37_] = init37_(i0_37_)
                d_1470_v40_ = nw222_
                d_1473_v41_: _dafny.MultiSet
                d_1473_v41_ = _dafny.MultiSet([d_1433_i2_])
                d_1474_v42_: _dafny.MultiSet
                d_1474_v42_ = _dafny.MultiSet([d_1473_v41_])
                index206_ = default__.safeIndex(298, (d_1470_v40_).length(0))
                (d_1470_v40_)[index206_] = default__.safeModuloInt((d_1474_v42_).cardinality, d_1433_i2_)
                index207_ = default__.safeIndex(298, (d_1470_v40_).length(0))
                (d_1470_v40_)[index207_] = (self).f13
                d_1475_v43_: _dafny.Seq
                d_1475_v43_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ecuurfubj"))
                index208_ = default__.safeIndex(298, (d_1470_v40_).length(0))
                (d_1470_v40_)[index208_] = (0) - (len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gyxgc"))) + (d_1475_v43_)))
                d_1476_v44_: str
                d_1476_v44_ = _dafny.CodePoint('t')
                d_1477_v45_: _dafny.Set
                d_1477_v45_ = _dafny.Set({(d_1475_v43_).set(default__.safeIndex((d_1470_v40_)[default__.safeIndex(298, (d_1470_v40_).length(0))], len(d_1475_v43_)), d_1476_v44_), d_1475_v43_, d_1475_v43_})
                d_1478_v46_: _dafny.Seq
                d_1478_v46_ = _dafny.SeqWithoutIsStrInference([(d_1477_v45_).intersection(d_1477_v45_)])
                d_1477_v45_ = (d_1478_v46_)[default__.safeIndex(d_1433_i2_, len(d_1478_v46_))]
                d_1479_v47_: _dafny.Array
                def lambda59_(d_1480_i5_):
                    return (self).f14

                init38_ = lambda59_
                nw223_ = _dafny.Array(None, 12)
                for i0_38_ in range(nw223_.length(0)):
                    nw223_[i0_38_] = init38_(i0_38_)
                d_1479_v47_ = nw223_
                d_1481_v48_: _dafny.Map
                d_1481_v48_ = _dafny.Map({(self).f13: d_1475_v43_})
                d_1482_v49_: C3
                nw224_ = C3()
                nw224_.ctor__(d_1406_v0_, d_1479_v47_, d_1481_v48_, (self).f13)
                d_1482_v49_ = nw224_
                d_1483_v50_: _dafny.Seq
                d_1483_v50_ = _dafny.SeqWithoutIsStrInference([d_1482_v49_])
                d_1484_v51_: _dafny.Set
                d_1484_v51_ = _dafny.Set({(d_1483_v50_)[default__.safeIndex((self).f13, len(d_1483_v50_))]})
                rhs236_ = (self).f14
                rhs237_ = (d_1484_v51_) | ((_dafny.Set({d_1482_v49_, d_1482_v49_})) - (d_1484_v51_))
                lhs191_ = globalState
                lhs191_.f0 = rhs236_
                d_1484_v51_ = rhs237_
            d_1485_v52_: _dafny.Seq
            d_1485_v52_ = _dafny.SeqWithoutIsStrInference([d_1433_i2_, d_1433_i2_])
            d_1486_v53_: _dafny.Set
            d_1486_v53_ = _dafny.Set({(d_1485_v52_) < (d_1485_v52_), ((self).f13) >= ((0) - ((self).f13)), (self).f14, (self).f14, (self).f14})
            d_1486_v53_ = (d_1486_v53_).intersection(d_1486_v53_)
            d_1487_v54_: _dafny.Map
            d_1487_v54_ = _dafny.Map({(self).f14: not((self).f14)})
            d_1488_v55_: D4
            d_1488_v55_ = D4_DC14((self).f14, d_1487_v54_, (self).f13, d_1487_v54_)
            d_1489_v56_: _dafny.Seq
            d_1489_v56_ = _dafny.SeqWithoutIsStrInference([(d_1488_v55_).cf28, (self).f14])
            d_1490_v57_: _dafny.Map
            d_1490_v57_ = _dafny.Map({d_1433_i2_: d_1489_v56_})
            d_1491_v58_: _dafny.Seq
            d_1491_v58_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wdgqi"))
            d_1492_v59_: _dafny.Map
            d_1492_v59_ = _dafny.Map({_dafny.Set({(self).f14, (self).f14}): (_dafny.Map({d_1433_i2_: d_1433_i2_})).set(len(d_1491_v58_), d_1433_i2_)})
            d_1490_v57_ = (d_1490_v57_).set(len(d_1492_v59_), d_1489_v56_)
            (globalState).f5 = d_1433_i2_
        d_1493_v60_: _dafny.Array
        nw225_ = _dafny.Array(int(0), 24)
        d_1493_v60_ = nw225_
        d_1494_v61_: D2
        d_1494_v61_ = D2_DC6((self).f13, d_1493_v60_)
        d_1494_v61_ = d_1494_v61_
        d_1495_v62_: D2
        d_1495_v62_ = D2_DC8(False, (self).f14)
        d_1496_v63_: str
        d_1496_v63_ = _dafny.CodePoint('h')
        d_1497_v64_: D0
        d_1497_v64_ = D0_DC1((self).f14, (self).f13, d_1496_v63_)
        hi11_ = default__.safeModuloInt((self).f13, (self).fm16(default__.fm49((self).f13, (self).fm16(d_1495_v62_, d_1497_v64_, globalState), (self).f14, (self).f13, globalState), d_1497_v64_, globalState))
        for d_1498_i6_ in range((self).f13, hi11_):
            (globalState).f0 = (self).f14
            d_1406_v0_ = (d_1406_v0_).set((self).f13, (-670) + ((self).f13))
            d_1499_v65_: _dafny.MultiSet
            d_1499_v65_ = _dafny.MultiSet([(self).f14, False])
            r0 = ((d_1499_v65_).cardinality) * (default__.fm2(len(_dafny.Set({d_1498_i6_})), globalState))
            d_1500_v66_: _dafny.Seq
            d_1500_v66_ = _dafny.SeqWithoutIsStrInference([(self).f13])
            (globalState).f0 = (len((d_1500_v66_) + (d_1500_v66_))) <= (-581)
        r0 = (self).f13
        d_1501_v67_: D2
        d_1501_v67_ = D2_DC5(d_1406_v0_)
        r1 = d_1501_v67_
        r2 = (self).f14
        return r0, r1, r2

    def m9(self, p0, p1, globalState):
        r0: _dafny.Set = _dafny.Set({})
        r1: int = int(0)
        r2: _dafny.MultiSet = _dafny.MultiSet({})
        d_1502_v0_: _dafny.Map
        d_1502_v0_ = _dafny.Map({(self).f14: (self).f13})
        d_1503_v1_: _dafny.MultiSet
        d_1503_v1_ = _dafny.MultiSet([(self).f13])
        d_1504_v2_: _dafny.MultiSet
        d_1504_v2_ = _dafny.MultiSet([d_1503_v1_])
        d_1505_v3_: _dafny.Array
        def lambda60_(d_1506_i0_):
            return (self).f14

        init39_ = lambda60_
        nw226_ = _dafny.Array(None, 22)
        for i0_39_ in range(nw226_.length(0)):
            nw226_[i0_39_] = init39_(i0_39_)
        d_1505_v3_ = nw226_
        d_1507_v4_: _dafny.Map
        d_1507_v4_ = _dafny.Map({default__.fm2((self).f13, globalState): p0})
        d_1508_v5_: C2
        nw227_ = C2()
        nw227_.ctor__((self).f13, d_1504_v2_, d_1505_v3_, d_1507_v4_)
        d_1508_v5_ = nw227_
        d_1509_v6_: _dafny.Seq
        d_1509_v6_ = _dafny.SeqWithoutIsStrInference([(self).f14, (self).f14])
        d_1510_v7_: _dafny.Seq
        d_1510_v7_ = _dafny.SeqWithoutIsStrInference([len(d_1509_v6_)])
        d_1511_v8_: _dafny.Map
        d_1511_v8_ = _dafny.Map({d_1508_v5_: d_1510_v7_})
        d_1502_v0_ = (d_1502_v0_).set(((self).f13) in (((d_1511_v8_)[d_1508_v5_] if (d_1508_v5_) in (d_1511_v8_) else d_1510_v7_)), (self).f13)
        d_1512_i1_: int
        d_1512_i1_ = 0
        with _dafny.label("9"):
            while (self).f14:
                with _dafny.c_label("9"):
                    if (d_1512_i1_) >= (100):
                        raise _dafny.Break("9")
                    d_1512_i1_ = (d_1512_i1_) + (1)
                    r1 = default__.safeDivisionInt((0) - ((self).f13), (self).f13)
                    d_1513_v9_: _dafny.Set
                    d_1513_v9_ = _dafny.Set({d_1505_v3_, d_1505_v3_, d_1505_v3_, d_1505_v3_})
                    d_1513_v9_ = d_1513_v9_
                    (globalState).f5 = (self).f13
                    d_1514_v11_: _dafny.Map
                    d_1514_v11_ = _dafny.Map({(self).f13: (self).f14})
                    d_1515_v12_: C2
                    nw228_ = C2()
                    def iife115_():
                        coll41_ = _dafny.Map()
                        compr_41_: int
                        for compr_41_ in (d_1514_v11_).keys.Elements:
                            d_1516_v10_: int = compr_41_
                            if (d_1516_v10_) in (d_1514_v11_):
                                coll41_[(d_1516_v10_) + ((self).f13)] = p1
                        return _dafny.Map(coll41_)
                    nw228_.ctor__((self).f13, d_1504_v2_, d_1505_v3_, (iife115_()
                    ) | (_dafny.Map({-554: p1})))
                    d_1515_v12_ = nw228_
                    pass
            pass
        d_1517_v13_: D2
        d_1517_v13_ = D2_DC8((self).f14, (self).f14)
        d_1518_v14_: str
        d_1518_v14_ = _dafny.CodePoint('d')
        d_1519_v15_: _dafny.Map
        d_1519_v15_ = _dafny.Map({(self).f13: _dafny.Set({(self).f13})})
        d_1520_v16_: _dafny.Set
        d_1520_v16_ = _dafny.Set({len(d_1502_v0_), (self).f13, (self).f13})
        def iife116_(_pat_let37_0):
            def iife117_(d_1521_dt__update__tmp_h0_):
                def iife118_(_pat_let38_0):
                    def iife119_(d_1522_dt__update_hcf19_h0_):
                        return D2_DC8(d_1522_dt__update_hcf19_h0_, (d_1521_dt__update__tmp_h0_).cf20)
                    return iife119_(_pat_let38_0)
                return iife118_((self).f14)
            return iife117_(_pat_let37_0)
        rhs238_ = ((449 if False else (self).f13)) + (((self).f13) + ((self).fm16(iife116_(d_1517_v13_), D0_DC1(True, (self).f13, d_1518_v14_), globalState)))
        rhs239_ = d_1505_v3_
        rhs240_ = d_1509_v6_
        rhs241_ = ((self).f13) < ((0) - ((self).f13))
        rhs242_ = (d_1520_v16_).issubset(((d_1519_v15_)[(self).f13] if ((self).f13) in (d_1519_v15_) else d_1520_v16_))
        lhs192_ = globalState
        lhs193_ = globalState
        r1 = rhs238_
        d_1505_v3_ = rhs239_
        d_1509_v6_ = rhs240_
        lhs192_.f0 = rhs241_
        lhs193_.f0 = rhs242_
        if ((self).f14 if ((self).f13) == ((self).f13) else (d_1518_v14_) in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q")))):
            d_1523_v17_: _dafny.Set
            d_1523_v17_ = _dafny.Set({(self).f14, (self).f14, (self).f14})
            (globalState).f0 = (d_1523_v17_).isdisjoint(d_1523_v17_)
            d_1524_v18_: _dafny.Map
            d_1524_v18_ = _dafny.Map({(self).f13: (self).f13})
            d_1525_v19_: D2
            d_1525_v19_ = D2_DC5(d_1524_v18_)
            rhs243_ = D2_DC5(d_1524_v18_)
            rhs244_ = (d_1509_v6_)[default__.safeIndex(-56, len(d_1509_v6_))]
            rhs245_ = ((self).f13) - ((d_1510_v7_)[default__.safeIndex(978, len(d_1510_v7_))])
            lhs194_ = globalState
            lhs195_ = globalState
            d_1525_v19_ = rhs243_
            lhs194_.f0 = rhs244_
            lhs195_.f5 = rhs245_
            r1 = ((self).f13) + ((self).f13)
            d_1526_v20_: _dafny.Map
            d_1526_v20_ = _dafny.Map({(d_1509_v6_)[default__.safeIndex((self).f13, len(d_1509_v6_))]: False})
            (globalState).f0 = ((d_1526_v20_)[((self).f14) or (False)] if (((self).f14) or (False)) in (d_1526_v20_) else not((d_1520_v16_).isdisjoint(d_1520_v16_)))
            d_1527_v21_: _dafny.Array
            def lambda61_(d_1528_i2_):
                return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jfhhroih"))

            init40_ = lambda61_
            nw229_ = _dafny.Array(None, 28)
            for i0_40_ in range(nw229_.length(0)):
                nw229_[i0_40_] = init40_(i0_40_)
            d_1527_v21_ = nw229_
            index209_ = default__.safeIndex(805, (d_1527_v21_).length(0))
            (d_1527_v21_)[index209_] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pli"))) + (p0)
            index210_ = default__.safeIndex(805, (d_1527_v21_).length(0))
            (d_1527_v21_)[index210_] = p0
        elif True:
            d_1520_v16_ = d_1520_v16_
            d_1529_v22_: _dafny.Array
            nw230_ = _dafny.Array(None, 11)
            nw230_[int(0)] = d_1505_v3_
            nw230_[int(1)] = (d_1505_v3_ if (self).f14 else d_1505_v3_)
            nw230_[int(2)] = d_1505_v3_
            nw230_[int(3)] = d_1505_v3_
            nw230_[int(4)] = d_1505_v3_
            nw230_[int(5)] = d_1505_v3_
            nw230_[int(6)] = d_1505_v3_
            nw230_[int(7)] = d_1505_v3_
            nw230_[int(8)] = d_1505_v3_
            nw230_[int(9)] = d_1505_v3_
            nw230_[int(10)] = d_1505_v3_
            d_1529_v22_ = nw230_
            index211_ = default__.safeIndex(247, (d_1529_v22_).length(0))
            (d_1529_v22_)[index211_] = d_1505_v3_
            index212_ = default__.safeIndex(247, (d_1529_v22_).length(0))
            (d_1529_v22_)[index212_] = d_1505_v3_
            d_1530_v23_: _dafny.Map
            d_1530_v23_ = _dafny.Map({len(d_1510_v7_): (self).f14})
            (globalState).f5 = (len(d_1530_v23_)) + (((self).f13) * ((self).f13))
            index213_ = default__.safeIndex(936, (d_1505_v3_).length(0))
            (d_1505_v3_)[index213_] = (self).f14
            index214_ = default__.safeIndex(936, (d_1505_v3_).length(0))
            (d_1505_v3_)[index214_] = False
            d_1510_v7_ = d_1510_v7_
        d_1531_v24_: _dafny.Set
        d_1531_v24_ = _dafny.Set({True})
        (globalState).f5 = len(d_1531_v24_)
        d_1532_v25_: D1
        d_1532_v25_ = D1_DC4(D0_DC0(_dafny.SeqWithoutIsStrInference([(self).f14, (self).f14]), (self).f13), (0) - ((self).f13), p1, (self).f14, (self).f13)
        d_1533_i3_: int
        d_1533_i3_ = 0
        with _dafny.label("10"):
            while not (((self).f13) != (623)) or (((self).f13) >= ((d_1532_v25_).cf8)):
                with _dafny.c_label("10"):
                    if (d_1533_i3_) >= (100):
                        raise _dafny.Break("10")
                    d_1533_i3_ = (d_1533_i3_) + (1)
                    d_1534_v26_: _dafny.Array
                    nw231_ = _dafny.Array(_dafny.Seq({}), 11)
                    d_1534_v26_ = nw231_
                    index215_ = default__.safeIndex(372, (d_1534_v26_).length(0))
                    (d_1534_v26_)[index215_] = (d_1510_v7_) + ((d_1510_v7_).set(default__.safeIndex((self).f13, len(d_1510_v7_)), len(_dafny.SeqWithoutIsStrInference([not((self).f14)]))))
                    d_1535_v27_: _dafny.Map
                    d_1535_v27_ = _dafny.Map({(self).f13: (self).f13})
                    d_1536_v28_: D0
                    d_1536_v28_ = D0_DC0(default__.fm4((self).f13, (self).f13, (self).f13, globalState), (self).f13)
                    d_1537_v29_: _dafny.Set
                    d_1537_v29_ = _dafny.Set({d_1505_v3_})
                    pat_let_tv30_ = d_1536_v28_
                    pat_let_tv31_ = globalState
                    d_1538_v30_: _dafny.Array
                    nw232_ = _dafny.Array(None, 5)
                    nw232_[int(0)] = (len(d_1509_v6_)) * ((self).f13)
                    nw232_[int(1)] = (self).f13
                    nw232_[int(2)] = len((d_1535_v27_) | (d_1535_v27_))
                    def iife120_(_pat_let39_0):
                        def iife121_(d_1539_dt__update__tmp_h1_):
                            def iife122_(_pat_let40_0):
                                def iife123_(d_1540_dt__update_hcf7_h0_):
                                    def iife124_(_pat_let41_0):
                                        def iife125_(d_1541_dt__update_hcf9_h0_):
                                            def iife126_(_pat_let42_0):
                                                def iife127_(d_1542_dt__update_hcf10_h0_):
                                                    return D1_DC4(d_1540_dt__update_hcf7_h0_, (d_1539_dt__update__tmp_h1_).cf8, d_1541_dt__update_hcf9_h0_, d_1542_dt__update_hcf10_h0_, (d_1539_dt__update__tmp_h1_).cf11)
                                                return iife127_(_pat_let42_0)
                                            return iife126_((self).f14)
                                        return iife125_(_pat_let41_0)
                                    return iife124_(default__.fm18((self).f14, pat_let_tv31_))
                                return iife123_(_pat_let40_0)
                            return iife122_(pat_let_tv30_)
                        return iife121_(_pat_let39_0)
                    nw232_[int(3)] = (0) - ((iife120_(d_1532_v25_)).cf11)
                    nw232_[int(4)] = len(d_1537_v29_)
                    d_1538_v30_ = nw232_
                    index216_ = default__.safeIndex(420, (d_1538_v30_).length(0))
                    (d_1538_v30_)[index216_] = default__.fm2((self).f13, globalState)
                    d_1543_v31_: _dafny.Map
                    d_1543_v31_ = _dafny.Map({(self).f14: p0})
                    index217_ = default__.safeIndex(372, (d_1534_v26_).length(0))
                    index218_ = default__.safeIndex(420, (d_1538_v30_).length(0))
                    rhs246_ = (default__.fm26(default__.safeModuloInt(len((d_1543_v31_).set((self).f14, p1)), len(d_1535_v27_)), d_1518_v14_, (self).f14, d_1518_v14_, globalState)).set(default__.safeIndex((self).f13, len(default__.fm26(default__.safeModuloInt(len((d_1543_v31_).set((self).f14, p1)), len(d_1535_v27_)), d_1518_v14_, (self).f14, d_1518_v14_, globalState))), (self).f13)
                    rhs247_ = ((default__.fm34(len(p0), globalState)).intersection(d_1531_v24_)).ispropersubset(d_1531_v24_)
                    rhs248_ = d_1510_v7_
                    rhs249_ = (self).f13
                    rhs250_ = default__.fm3(globalState)
                    lhs196_ = globalState
                    lhs197_ = d_1534_v26_
                    lhs198_ = default__.safeIndex(372, (d_1534_v26_).length(0))
                    lhs199_ = d_1538_v30_
                    lhs200_ = default__.safeIndex(420, (d_1538_v30_).length(0))
                    lhs201_ = globalState
                    d_1510_v7_ = rhs246_
                    lhs196_.f0 = rhs247_
                    lhs197_[lhs198_] = rhs248_
                    lhs199_[lhs200_] = rhs249_
                    lhs201_.f0 = rhs250_
                    r1 = (self).f13
                    if (self).f14:
                        d_1544_v32_: _dafny.Array
                        nw233_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 28)
                        d_1544_v32_ = nw233_
                        d_1544_v32_ = d_1544_v32_
                        d_1545_v33_: _dafny.Map
                        d_1545_v33_ = _dafny.Map({(d_1534_v26_)[default__.safeIndex(372, (d_1534_v26_).length(0))]: (d_1538_v30_)[default__.safeIndex(420, (d_1538_v30_).length(0))]})
                        d_1546_v34_: _dafny.Map
                        d_1546_v34_ = _dafny.Map({(d_1538_v30_)[default__.safeIndex(420, (d_1538_v30_).length(0))]: d_1535_v27_})
                        index219_ = default__.safeIndex(420, (d_1538_v30_).length(0))
                        (d_1538_v30_)[index219_] = ((d_1545_v33_)[((d_1534_v26_)[default__.safeIndex(372, (d_1534_v26_).length(0))]) + (d_1510_v7_)] if (((d_1534_v26_)[default__.safeIndex(372, (d_1534_v26_).length(0))]) + (d_1510_v7_)) in (d_1545_v33_) else len(d_1546_v34_))
                        index220_ = default__.safeIndex(372, (d_1534_v26_).length(0))
                        (d_1534_v26_)[index220_] = (default__.fm40((self).f13, _dafny.Set({(self).f14, (self).f14}), globalState)) + ((d_1534_v26_)[default__.safeIndex(372, (d_1534_v26_).length(0))])
                        index221_ = default__.safeIndex(420, (d_1538_v30_).length(0))
                        rhs251_ = (0) - ((d_1538_v30_)[default__.safeIndex(420, (d_1538_v30_).length(0))])
                        rhs252_ = default__.fm2(len(_dafny.SeqWithoutIsStrInference([d_1531_v24_, d_1531_v24_])), globalState)
                        lhs202_ = d_1538_v30_
                        lhs203_ = default__.safeIndex(420, (d_1538_v30_).length(0))
                        lhs204_ = globalState
                        lhs202_[lhs203_] = rhs251_
                        lhs204_.f5 = rhs252_
                        d_1547_v35_: _dafny.MultiSet
                        d_1547_v35_ = _dafny.MultiSet([default__.fm3(globalState)])
                        (globalState).f0 = (d_1508_v5_).fm22(d_1547_v35_, ((self).f13) >= ((self).f13), (self).f13, globalState)
                    elif True:
                        index222_ = default__.safeIndex(981, (d_1505_v3_).length(0))
                        (d_1505_v3_)[index222_] = (self).f14
                        d_1548_v36_: _dafny.Map
                        d_1548_v36_ = _dafny.Map({(d_1538_v30_)[default__.safeIndex(420, (d_1538_v30_).length(0))]: d_1509_v6_})
                        d_1549_v38_: _dafny.Map
                        d_1549_v38_ = _dafny.Map({(d_1538_v30_)[default__.safeIndex(420, (d_1538_v30_).length(0))]: (self).f14})
                        d_1550_v40_: D11
                        def iife128_():
                            coll42_ = _dafny.Set()
                            compr_42_: int
                            for compr_42_ in (d_1549_v38_).keys.Elements:
                                d_1551_v39_: int = compr_42_
                                if (d_1551_v39_) in (d_1549_v38_):
                                    coll42_ = coll42_.union(_dafny.Set([(d_1551_v39_) * (342)]))
                            return _dafny.Set(coll42_)
                        d_1550_v40_ = D11_DC30(iife128_()
)
                        index223_ = default__.safeIndex(981, (d_1505_v3_).length(0))
                        def iife129_():
                            coll43_ = _dafny.Set()
                            compr_43_: int
                            for compr_43_ in (d_1548_v36_).keys.Elements:
                                d_1553_v37_: int = compr_43_
                                if (d_1553_v37_) in (d_1548_v36_):
                                    coll43_ = coll43_.union(_dafny.Set([(d_1553_v37_) + (744)]))
                            return _dafny.Set(coll43_)
                        (d_1505_v3_)[index223_] = (((d_1550_v40_).cf57).intersection(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([d_1518_v14_ for d_1552_i4_ in range(default__.abs(-74))])), (d_1538_v30_)[default__.safeIndex(420, (d_1538_v30_).length(0))]}))).ispropersubset((d_1520_v16_).intersection(iife129_()
                        ))
                        index224_ = default__.safeIndex(372, (d_1534_v26_).length(0))
                        (d_1534_v26_)[index224_] = ((_dafny.SeqWithoutIsStrInference([(0) - (len(d_1520_v16_)) for d_1554_i5_ in range(default__.abs(401))])) + (d_1510_v7_)).set(default__.safeIndex((self).f13, len((_dafny.SeqWithoutIsStrInference([(0) - (len(d_1520_v16_)) for d_1555_i5_ in range(default__.abs(401))])) + (d_1510_v7_))), (771) * ((d_1538_v30_)[default__.safeIndex(420, (d_1538_v30_).length(0))]))
                        (globalState).f5 = (self).f13
                        d_1556_v41_: D11
                        d_1556_v41_ = D11_DC33(((self).f13) - (len(d_1535_v27_)), ((d_1538_v30_)[default__.safeIndex(420, (d_1538_v30_).length(0))]) >= ((self).f13), (d_1538_v30_)[default__.safeIndex(420, (d_1538_v30_).length(0))], False, d_1531_v24_)
                        d_1557_v43_: D0
                        def iife130_():
                            coll44_ = _dafny.Map()
                            compr_44_: str
                            for compr_44_ in (p1).Elements:
                                d_1558_v42_: str = compr_44_
                                if (d_1558_v42_) in (p1):
                                    coll44_[d_1558_v42_] = d_1509_v6_
                            return _dafny.Map(coll44_)
                        d_1557_v43_ = D0_DC1((self).f14, len(iife130_()
), d_1518_v14_)
                        index225_ = default__.safeIndex(420, (d_1538_v30_).length(0))
                        index226_ = default__.safeIndex(420, (d_1538_v30_).length(0))
                        index227_ = default__.safeIndex(981, (d_1505_v3_).length(0))
                        rhs253_ = d_1556_v41_
                        rhs254_ = (d_1538_v30_)[default__.safeIndex(420, (d_1538_v30_).length(0))]
                        rhs255_ = (self).fm16(d_1517_v13_, d_1557_v43_, globalState)
                        rhs256_ = ((d_1538_v30_)[default__.safeIndex(420, (d_1538_v30_).length(0))]) >= ((d_1538_v30_)[default__.safeIndex(420, (d_1538_v30_).length(0))])
                        rhs257_ = d_1537_v29_
                        lhs205_ = d_1538_v30_
                        lhs206_ = default__.safeIndex(420, (d_1538_v30_).length(0))
                        lhs207_ = d_1538_v30_
                        lhs208_ = default__.safeIndex(420, (d_1538_v30_).length(0))
                        lhs209_ = d_1505_v3_
                        lhs210_ = default__.safeIndex(981, (d_1505_v3_).length(0))
                        d_1556_v41_ = rhs253_
                        lhs205_[lhs206_] = rhs254_
                        lhs207_[lhs208_] = rhs255_
                        lhs209_[lhs210_] = rhs256_
                        d_1537_v29_ = rhs257_
                        d_1559_v44_: _dafny.MultiSet
                        d_1559_v44_ = _dafny.MultiSet([(d_1505_v3_)[default__.safeIndex(981, (d_1505_v3_).length(0))], (self).f14, (d_1505_v3_)[default__.safeIndex(981, (d_1505_v3_).length(0))]])
                        d_1560_v45_: D5
                        d_1560_v45_ = D5_DC17(d_1510_v7_, not((self).f14))
                        rhs258_ = (d_1508_v5_).fm22(d_1559_v44_, ((self).f14) in (d_1543_v31_), ((self).f13) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uuiikup")))), globalState)
                        rhs259_ = ((d_1520_v16_).intersection(_dafny.Set({len(d_1509_v6_), (self).f13, (self).f13, (self).f13}))).ispropersubset(d_1520_v16_)
                        rhs260_ = ((d_1549_v38_)[len(p1)] if (len(p1)) in (d_1549_v38_) else (self).f14)
                        rhs261_ = ((0) - ((d_1538_v30_)[default__.safeIndex(420, (d_1538_v30_).length(0))]) if (self).f14 else default__.safeDivisionInt((self).f13, (self).f13))
                        rhs262_ = (d_1560_v45_).cf37
                        lhs211_ = globalState
                        lhs212_ = globalState
                        lhs213_ = globalState
                        lhs214_ = globalState
                        lhs215_ = globalState
                        lhs211_.f0 = rhs258_
                        lhs212_.f0 = rhs259_
                        lhs213_.f0 = rhs260_
                        lhs214_.f5 = rhs261_
                        lhs215_.f0 = rhs262_
                    (globalState).f2 = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xpbyxlev"))) + (default__.fm18((self).f14, globalState))
                    pass
            pass
        d_1561_v46_: _dafny.Array
        def lambda62_(d_1562_i6_):
            return (d_1562_i6_) * ((self).f13)

        init41_ = lambda62_
        nw234_ = _dafny.Array(None, 2)
        for i0_41_ in range(nw234_.length(0)):
            nw234_[i0_41_] = init41_(i0_41_)
        d_1561_v46_ = nw234_
        d_1563_v47_: _dafny.Set
        d_1563_v47_ = _dafny.Set({d_1561_v46_, d_1561_v46_, d_1561_v46_, d_1561_v46_})
        r0 = d_1563_v47_
        r1 = ((self).f13) + (default__.safeModuloInt((0) - ((self).f13), (self).f13))
        r2 = (d_1503_v1_).intersection((d_1503_v1_) - (d_1503_v1_))
        return r0, r1, r2

    @property
    def f13(self):
        return self._f13
    @property
    def f14(self):
        return self._f14

class C7(T0, T1):
    def  __init__(self):
        self._f9: _dafny.Array = _dafny.Array(None, 0)
        self._f6: int = int(0)
        self._f10: _dafny.Map = _dafny.Map({})
        self.f11: _dafny.Array = _dafny.Array(None, 0)
        self.f12: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C7"
    @property
    def f9(self):
        return self._f9
    @f9.setter
    def f9(self, value):
        self._f9 = value
    @property
    def f6(self):
        return self._f6
    @property
    def f10(self):
        return self._f10
    def ctor__(self, f11, f12, f6, f9, f10):
        (self).f11 = f11
        (self).f12 = f12
        (self)._f6 = f6
        (self).f9 = f9
        (self)._f10 = f10

    def fm5(self, p0, globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([_dafny.CodePoint('u')]) for d_1564_i0_ in range(default__.abs(-609))])) + (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([_dafny.CodePoint('u'), _dafny.CodePoint('l'), (D0_DC1(False, 711, _dafny.CodePoint('g'))).cf4, _dafny.CodePoint('k')]) for d_1565_i1_ in range(default__.abs(596))]))) + (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([_dafny.CodePoint('l'), _dafny.CodePoint('b')])]))

    def fm6(self, p0, globalState):
        return (_dafny.Map({(0) - ((self).f6): _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(self).f6]))])})) | (_dafny.Map({(self).f6: _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(self).f6, (self).f6])), (self).f6, (self).f6, (self).f6])}))

    def fm13(self, p0, p1, p2, p3, globalState):
        return ((_dafny.Map({True: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j"))}) if False else _dafny.Map({True: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ggjslaqas"))}))) | ((_dafny.Map({True: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kdhe"))})) | (_dafny.Map({not(True): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_1566_i0_ in range(default__.abs(298))])})))

    def fm14(self, globalState):
        return (0) - (((D0_DC0(_dafny.SeqWithoutIsStrInference([False, False]), (self).f6)).cf1) - (((self).f6) + ((self).f6)))

    def m6(self, p0, p1, p2, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: bool = False
        hi12_ = (self).fm14(globalState)
        for d_1567_i0_ in range(len((_dafny.SeqWithoutIsStrInference([p0, p0])) + (p1)), hi12_):
            (globalState).f5 = (0) - (default__.safeDivisionInt((p2 if p0 else d_1567_i0_), d_1567_i0_))
            (globalState).f2 = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bksq"))
            d_1568_v0_: _dafny.Seq
            d_1568_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rwrd"))
            d_1569_v1_: D0
            d_1569_v1_ = D0_DC0(default__.fm4(d_1567_i0_, d_1567_i0_, (self).f6, globalState), len((d_1568_v0_) + (d_1568_v0_)))
            d_1569_v1_ = d_1569_v1_
            d_1570_v2_: _dafny.MultiSet
            d_1570_v2_ = _dafny.MultiSet([self.f9])
            if (d_1570_v2_).ispropersubset(d_1570_v2_):
                d_1571_v3_: C0
                nw235_ = C0()
                nw235_.ctor__(self.f9, p0)
                d_1571_v3_ = nw235_
                (globalState).f5 = (self).f6
                (globalState).f5 = ((p2) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vsvuef"))))) + (p2)
                d_1572_v4_: _dafny.Array
                def lambda63_(d_1573_v3_):
                    def lambda64_(d_1574_i1_):
                        return _dafny.SeqWithoutIsStrInference([d_1573_v3_.f20])

                    return lambda64_

                init42_ = lambda63_(d_1571_v3_)
                nw236_ = _dafny.Array(None, 16)
                for i0_42_ in range(nw236_.length(0)):
                    nw236_[i0_42_] = init42_(i0_42_)
                d_1572_v4_ = nw236_
                index228_ = default__.safeIndex(778, (d_1572_v4_).length(0))
                (d_1572_v4_)[index228_] = (_dafny.SeqWithoutIsStrInference([d_1571_v3_.f20, d_1571_v3_.f20])) + (p1)
                d_1575_v5_: _dafny.Seq
                d_1575_v5_ = _dafny.SeqWithoutIsStrInference([p0])
                d_1576_v6_: _dafny.Map
                d_1576_v6_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_1577_i2_ in range(default__.abs(551))])): _dafny.SeqWithoutIsStrInference([not(d_1571_v3_.f20), d_1571_v3_.f20])})
                index229_ = default__.safeIndex(778, (d_1572_v4_).length(0))
                rhs263_ = default__.fm3(globalState)
                rhs264_ = ((d_1576_v6_)[(self).f6] if ((self).f6) in (d_1576_v6_) else d_1575_v5_)
                rhs265_ = p1
                rhs266_ = self.f12
                rhs267_ = d_1567_i0_
                lhs216_ = globalState
                lhs217_ = d_1572_v4_
                lhs218_ = default__.safeIndex(778, (d_1572_v4_).length(0))
                lhs219_ = self
                lhs220_ = globalState
                lhs216_.f0 = rhs263_
                lhs217_[lhs218_] = rhs264_
                d_1575_v5_ = rhs265_
                lhs219_.f11 = rhs266_
                lhs220_.f5 = rhs267_
                d_1578_v8_: _dafny.Set
                d_1578_v8_ = _dafny.Set({False, d_1571_v3_.f20, d_1571_v3_.f20})
                d_1579_v9_: _dafny.Map
                def iife131_():
                    coll45_ = _dafny.Map()
                    compr_45_: _dafny.Set
                    for compr_45_ in (_dafny.SeqWithoutIsStrInference([d_1578_v8_, d_1578_v8_])).Elements:
                        d_1580_v7_: _dafny.Set = compr_45_
                        if (d_1580_v7_) in (_dafny.SeqWithoutIsStrInference([d_1578_v8_, d_1578_v8_])):
                            coll45_[d_1580_v7_] = d_1571_v3_.f20
                    return _dafny.Map(coll45_)
                d_1579_v9_ = _dafny.Map({930: (default__.fm2(d_1567_i0_, globalState)) == (len(iife131_()
                ))})
                (d_1571_v3_).f20 = ((d_1579_v9_)[default__.safeModuloInt(p2, len(d_1568_v0_))] if (default__.safeModuloInt(p2, len(d_1568_v0_))) in (d_1579_v9_) else p0)
            elif True:
                d_1581_v10_: _dafny.MultiSet
                d_1581_v10_ = _dafny.MultiSet([p2])
                d_1582_v11_: _dafny.Map
                d_1582_v11_ = _dafny.Map({d_1581_v10_: p2})
                d_1582_v11_ = d_1582_v11_
                d_1583_v12_: D2
                d_1583_v12_ = D2_DC6(835, self.f11)
                d_1584_v13_: _dafny.Array
                def lambda65_(d_1585_i3_):
                    return _dafny.SeqWithoutIsStrInference([(self).f6 for d_1586_i4_ in range(default__.abs(957))])

                init43_ = lambda65_
                nw237_ = _dafny.Array(None, 28)
                for i0_43_ in range(nw237_.length(0)):
                    nw237_[i0_43_] = init43_(i0_43_)
                d_1584_v13_ = nw237_
                d_1587_v14_: _dafny.Seq
                d_1587_v14_ = _dafny.SeqWithoutIsStrInference([(d_1581_v10_).cardinality])
                index230_ = default__.safeIndex(944, (d_1584_v13_).length(0))
                (d_1584_v13_)[index230_] = d_1587_v14_
                d_1588_v15_: _dafny.Array
                nw238_ = _dafny.Array(D7.default()(), 20)
                d_1588_v15_ = nw238_
                pat_let_tv32_ = p1
                d_1589_v16_: _dafny.Array
                nw239_ = _dafny.Array(None, 16)
                nw239_[int(0)] = d_1569_v1_
                nw239_[int(1)] = D0_DC0(p1, len(d_1568_v0_))
                nw239_[int(2)] = d_1569_v1_
                nw239_[int(3)] = d_1569_v1_
                nw239_[int(4)] = d_1569_v1_
                nw239_[int(5)] = default__.fm32(p0, p0, p2, p0, globalState)
                nw239_[int(6)] = d_1569_v1_
                nw239_[int(7)] = d_1569_v1_
                nw239_[int(8)] = d_1569_v1_
                nw239_[int(9)] = default__.fm32(p0, p0, d_1567_i0_, p0, globalState)
                nw239_[int(10)] = d_1569_v1_
                nw239_[int(11)] = D0_DC0(_dafny.SeqWithoutIsStrInference([p0]), d_1567_i0_)
                nw239_[int(12)] = d_1569_v1_
                def iife132_(_pat_let43_0):
                    def iife133_(d_1590_dt__update__tmp_h0_):
                        def iife134_(_pat_let44_0):
                            def iife135_(d_1591_dt__update_hcf0_h0_):
                                return D0_DC0(d_1591_dt__update_hcf0_h0_, (d_1590_dt__update__tmp_h0_).cf1)
                            return iife135_(_pat_let44_0)
                        return iife134_(pat_let_tv32_)
                    return iife133_(_pat_let43_0)
                nw239_[int(13)] = iife132_(d_1569_v1_)
                nw239_[int(14)] = d_1569_v1_
                nw239_[int(15)] = d_1569_v1_
                d_1589_v16_ = nw239_
                d_1592_v17_: D7
                d_1592_v17_ = D7_DC21(d_1589_v16_)
                index231_ = default__.safeIndex(495, (d_1588_v15_).length(0))
                (d_1588_v15_)[index231_] = d_1592_v17_
                d_1593_v18_: _dafny.Seq
                d_1593_v18_ = _dafny.SeqWithoutIsStrInference([self.f12, self.f11])
                index232_ = default__.safeIndex(944, (d_1584_v13_).length(0))
                index233_ = default__.safeIndex(495, (d_1588_v15_).length(0))
                rhs268_ = d_1583_v12_
                rhs269_ = ((_dafny.SeqWithoutIsStrInference([len(_dafny.Map({p0: p0}))])).set(default__.safeIndex(65, len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({p0: p0}))]))), 848)) + ((d_1587_v14_) + (d_1587_v14_))
                rhs270_ = d_1592_v17_
                rhs271_ = p0
                rhs272_ = (d_1593_v18_)[default__.safeIndex(p2, len(d_1593_v18_))]
                lhs221_ = d_1584_v13_
                lhs222_ = default__.safeIndex(944, (d_1584_v13_).length(0))
                lhs223_ = d_1588_v15_
                lhs224_ = default__.safeIndex(495, (d_1588_v15_).length(0))
                lhs225_ = globalState
                lhs226_ = self
                d_1583_v12_ = rhs268_
                lhs221_[lhs222_] = rhs269_
                lhs223_[lhs224_] = rhs270_
                lhs225_.f0 = rhs271_
                lhs226_.f11 = rhs272_
                d_1594_v19_: _dafny.MultiSet
                d_1594_v19_ = _dafny.MultiSet([d_1581_v10_])
                d_1595_v20_: C2
                nw240_ = C2()
                nw240_.ctor__(((self).f6) + (d_1567_i0_), (d_1594_v19_ if not(p0) else d_1594_v19_), self.f9, (self).f10)
                d_1595_v20_ = nw240_
                d_1596_v21_: _dafny.Map
                d_1596_v21_ = _dafny.Map({(self).f6: d_1568_v0_})
                d_1596_v21_ = _dafny.Map({p2: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dkcllx"))})
                (globalState).f0 = (p0 if p0 else p0)
        arr46_ = self.f9
        index234_ = default__.safeIndex(449, (self.f9).length(0))
        arr46_[index234_] = p0
        d_1597_v22_: _dafny.Map
        d_1597_v22_ = _dafny.Map({p0: (self).f6})
        arr47_ = self.f9
        index235_ = default__.safeIndex(449, (self.f9).length(0))
        arr47_[index235_] = (not((len(_dafny.SeqWithoutIsStrInference([_dafny.Map({p0: p0}), _dafny.Map({p0: p0})]))) == (len(d_1597_v22_))) if p0 else (p0 if p0 else p0))
        if p0:
            d_1598_v23_: _dafny.Map
            d_1598_v23_ = _dafny.Map({(self).f6: not(p0)})
            d_1599_v24_: _dafny.Map
            d_1599_v24_ = _dafny.Map({d_1598_v23_: (self.f9)[default__.safeIndex(449, (self.f9).length(0))]})
            d_1600_v25_: D3
            d_1600_v25_ = D3_DC10(d_1599_v24_)
            d_1601_v26_: D3
            d_1601_v26_ = D3_DC11(d_1600_v25_)
            d_1602_v27_: _dafny.Map
            d_1602_v27_ = _dafny.Map({False: p1})
            d_1601_v26_ = default__.fm50(((d_1602_v27_)[(self.f9)[default__.safeIndex(449, (self.f9).length(0))]] if ((self.f9)[default__.safeIndex(449, (self.f9).length(0))]) in (d_1602_v27_) else p1), globalState)
            (globalState).f5 = (0) - ((self).f6)
            d_1603_v28_: _dafny.Seq
            d_1603_v28_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bvftkwuvg"))
            (globalState).f5 = default__.safeModuloInt((len(d_1603_v28_)) + ((self).f6), p2)
            if (p2) > (((self).f6) * ((self).f6)):
                d_1604_v29_: _dafny.Seq
                d_1604_v29_ = _dafny.SeqWithoutIsStrInference([self.f11])
                d_1605_v30_: _dafny.Map
                d_1605_v30_ = _dafny.Map({not((self.f9)[default__.safeIndex(449, (self.f9).length(0))]): p0})
                d_1606_v31_: D4
                d_1606_v31_ = D4_DC14((self.f9)[default__.safeIndex(449, (self.f9).length(0))], d_1605_v30_, (self).f6, d_1605_v30_)
                d_1607_v32_: _dafny.MultiSet
                d_1607_v32_ = _dafny.MultiSet([(d_1606_v31_).cf28])
                d_1608_v33_: _dafny.Seq
                d_1608_v33_ = _dafny.SeqWithoutIsStrInference([((d_1607_v32_)[(self.f9)[default__.safeIndex(449, (self.f9).length(0))]] if ((self.f9)[default__.safeIndex(449, (self.f9).length(0))]) in (d_1607_v32_) else (self).f6), (self).f6])
                rhs273_ = d_1604_v29_
                rhs274_ = ((d_1608_v33_)[default__.safeIndex((self).f6, len(d_1608_v33_))]) < (default__.fm2(p2, globalState))
                lhs227_ = globalState
                d_1604_v29_ = rhs273_
                lhs227_.f0 = rhs274_
                d_1609_v34_: _dafny.Array
                nw241_ = _dafny.Array(False, 15)
                d_1609_v34_ = nw241_
                d_1610_v35_: _dafny.Map
                d_1610_v35_ = _dafny.Map({d_1609_v34_: (self).f6})
                (globalState).f5 = ((d_1610_v35_)[d_1609_v34_] if (d_1609_v34_) in (d_1610_v35_) else p2)
                d_1608_v33_ = ((d_1608_v33_) + (_dafny.SeqWithoutIsStrInference([default__.fm2(p2, globalState)]))).set(default__.safeIndex(default__.fm2(len(p1), globalState), len((d_1608_v33_) + (_dafny.SeqWithoutIsStrInference([default__.fm2(p2, globalState)])))), (745) * (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_1611_i5_ in range(default__.abs(94))]))))
                arr48_ = self.f9
                index236_ = default__.safeIndex(449, (self.f9).length(0))
                arr48_[index236_] = p0
                d_1612_v36_: _dafny.MultiSet
                d_1612_v36_ = _dafny.MultiSet([(0) - ((self).f6)])
                d_1613_v37_: _dafny.MultiSet
                d_1613_v37_ = _dafny.MultiSet([d_1612_v36_, d_1612_v36_, d_1612_v36_])
                d_1614_v38_: C2
                nw242_ = C2()
                nw242_.ctor__((0) - (len(d_1598_v23_)), (d_1613_v37_) - (d_1613_v37_), d_1609_v34_, (_dafny.Map({p2: (((self).f10)[len(d_1603_v28_)] if (len(d_1603_v28_)) in ((self).f10) else d_1603_v28_)})) | (_dafny.Map({len(_dafny.Map({p0: (self).f6})): d_1603_v28_})))
                d_1614_v38_ = nw242_
            elif True:
                d_1615_v39_: C6
                nw243_ = C6()
                nw243_.ctor__(default__.safeModuloInt((0) - (p2), (self).f6), ((d_1598_v23_)[p2] if (p2) in (d_1598_v23_) else (self.f9)[default__.safeIndex(449, (self.f9).length(0))]))
                d_1615_v39_ = nw243_
                d_1615_v39_ = d_1615_v39_
                d_1616_v40_: _dafny.Map
                d_1616_v40_ = _dafny.Map({(self).f6: -125})
                d_1617_v41_: _dafny.Seq
                d_1617_v41_ = _dafny.SeqWithoutIsStrInference([(d_1615_v39_).f13])
                d_1618_v42_: _dafny.Array
                nw244_ = _dafny.Array(None, 16)
                nw244_[int(0)] = (self.f9)[default__.safeIndex(449, (self.f9).length(0))]
                nw244_[int(1)] = p0
                nw244_[int(2)] = p0
                nw244_[int(3)] = (self.f9)[default__.safeIndex(449, (self.f9).length(0))]
                nw244_[int(4)] = (p1)[default__.safeIndex(len(d_1617_v41_), len(p1))]
                nw244_[int(5)] = (self.f9)[default__.safeIndex(449, (self.f9).length(0))]
                nw244_[int(6)] = (self.f9)[default__.safeIndex(449, (self.f9).length(0))]
                nw244_[int(7)] = p0
                nw244_[int(8)] = p0
                nw244_[int(9)] = p0
                nw244_[int(10)] = p0
                nw244_[int(11)] = (self.f9)[default__.safeIndex(449, (self.f9).length(0))]
                nw244_[int(12)] = (self.f9)[default__.safeIndex(449, (self.f9).length(0))]
                nw244_[int(13)] = (self.f9)[default__.safeIndex(449, (self.f9).length(0))]
                nw244_[int(14)] = (self.f9)[default__.safeIndex(449, (self.f9).length(0))]
                nw244_[int(15)] = (self.f9)[default__.safeIndex(449, (self.f9).length(0))]
                d_1618_v42_ = nw244_
                d_1619_v43_: C3
                nw245_ = C3()
                nw245_.ctor__(_dafny.Map({(self).f6: len(d_1616_v40_)}), d_1618_v42_, (self).f10, (d_1617_v41_)[default__.safeIndex((d_1615_v39_).f13, len(d_1617_v41_))])
                d_1619_v43_ = nw245_
                d_1620_v44_: _dafny.Set
                d_1620_v44_ = _dafny.Set({(d_1615_v39_).f13, (self).f6})
                d_1621_v45_: D11
                d_1621_v45_ = D11_DC30(d_1620_v44_)
                d_1621_v45_ = d_1621_v45_
                d_1622_v46_: C5
                nw246_ = C5()
                nw246_.ctor__(((self).f6) * ((d_1615_v39_).f13), 690)
                d_1622_v46_ = nw246_
                d_1623_v47_: _dafny.Array
                nw247_ = _dafny.Array(D3.default()(), 23)
                d_1623_v47_ = nw247_
                pat_let_tv33_ = d_1600_v25_
                index237_ = default__.safeIndex(466, (d_1623_v47_).length(0))
                def iife136_(_pat_let45_0):
                    def iife137_(d_1624_dt__update__tmp_h1_):
                        def iife138_(_pat_let46_0):
                            def iife139_(d_1625_dt__update_hcf23_h0_):
                                return D3_DC11(d_1625_dt__update_hcf23_h0_)
                            return iife139_(_pat_let46_0)
                        return iife138_(pat_let_tv33_)
                    return iife137_(_pat_let45_0)
                (d_1623_v47_)[index237_] = iife136_(d_1601_v26_)
                index238_ = default__.safeIndex(466, (d_1623_v47_).length(0))
                (d_1623_v47_)[index238_] = d_1601_v26_
            (globalState).f5 = (D1_DC4(default__.fm32(p0, p0, (self).f6, (self.f9)[default__.safeIndex(449, (self.f9).length(0))], globalState), p2, d_1603_v28_, (self.f9)[default__.safeIndex(449, (self.f9).length(0))], p2)).cf8
        elif True:
            d_1626_v48_: str
            d_1626_v48_ = _dafny.CodePoint('u')
            d_1627_v49_: _dafny.Map
            d_1627_v49_ = _dafny.Map({d_1626_v48_: not(p0)})
            d_1627_v49_ = (d_1627_v49_).set(d_1626_v48_, p0)
            d_1628_v50_: _dafny.MultiSet
            d_1628_v50_ = _dafny.MultiSet([p0])
            d_1629_v51_: _dafny.Set
            d_1629_v51_ = _dafny.Set({(self.f9)[default__.safeIndex(449, (self.f9).length(0))]})
            (globalState).f5 = (0) - (((p2 if (self.f9)[default__.safeIndex(449, (self.f9).length(0))] else ((d_1628_v50_)[(self.f9)[default__.safeIndex(449, (self.f9).length(0))]] if ((self.f9)[default__.safeIndex(449, (self.f9).length(0))]) in (d_1628_v50_) else (self).f6))) - (((self).f6 if True else len(d_1629_v51_))))
            d_1630_v52_: D0
            d_1630_v52_ = D0_DC1((self.f9)[default__.safeIndex(449, (self.f9).length(0))], p2, d_1626_v48_)
            rhs275_ = (d_1630_v52_).cf4
            rhs276_ = (0) - ((0) - ((0) - (p2)))
            lhs228_ = globalState
            d_1626_v48_ = rhs275_
            lhs228_.f5 = rhs276_
            rhs277_ = default__.safeDivisionInt(default__.safeModuloInt(p2, (self).fm14(globalState)), (self).f6)
            rhs278_ = ((self).f6) - (((self).f6) * (p2))
            lhs229_ = globalState
            lhs230_ = globalState
            lhs229_.f5 = rhs277_
            lhs230_.f5 = rhs278_
            d_1631_v53_: _dafny.MultiSet
            d_1631_v53_ = _dafny.MultiSet([d_1626_v48_])
            d_1632_v54_: _dafny.Seq
            d_1632_v54_ = _dafny.SeqWithoutIsStrInference([d_1631_v53_, d_1631_v53_, d_1631_v53_])
            d_1633_v55_: D4
            d_1633_v55_ = D4_DC12(d_1632_v54_)
            d_1634_v56_: _dafny.Map
            d_1634_v56_ = _dafny.Map({d_1633_v55_: p2})
            d_1635_v57_: _dafny.Seq
            d_1635_v57_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ogh"))
            arr49_ = self.f9
            index239_ = default__.safeIndex(449, (self.f9).length(0))
            rhs279_ = ((not(not(p0))) == (default__.fm3(globalState)) if (default__.fm2((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f6, (self).f6]))).cardinality, globalState)) >= (len(d_1634_v56_)) else not(not(not((self.f9)[default__.safeIndex(449, (self.f9).length(0))]))))
            rhs280_ = d_1635_v57_
            lhs231_ = self.f9
            lhs232_ = default__.safeIndex(449, (self.f9).length(0))
            lhs233_ = globalState
            lhs231_[lhs232_] = rhs279_
            lhs233_.f2 = rhs280_
        d_1636_v58_: _dafny.MultiSet
        d_1636_v58_ = _dafny.MultiSet([(self.f9)[default__.safeIndex(449, (self.f9).length(0))]])
        d_1637_v59_: D6
        d_1637_v59_ = D6_DC19((self).f6, (self).f6, 428)
        d_1638_v60_: _dafny.Map
        d_1638_v60_ = _dafny.Map({((d_1636_v58_)[(self.f9)[default__.safeIndex(449, (self.f9).length(0))]] if ((self.f9)[default__.safeIndex(449, (self.f9).length(0))]) in (d_1636_v58_) else (0) - (p2)): (d_1637_v59_).cf41})
        def iife140_():
            coll46_ = _dafny.Map()
            compr_46_: int
            for compr_46_ in ((_dafny.SeqWithoutIsStrInference([(self).f6])).set(default__.safeIndex(p2, len(_dafny.SeqWithoutIsStrInference([(self).f6]))), len(d_1638_v60_))).Elements:
                d_1639_v61_: int = compr_46_
                if (d_1639_v61_) in ((_dafny.SeqWithoutIsStrInference([(self).f6])).set(default__.safeIndex(p2, len(_dafny.SeqWithoutIsStrInference([(self).f6]))), len(d_1638_v60_))):
                    coll46_[default__.safeDivisionInt(d_1639_v61_, (self).f6)] = p2
            return _dafny.Map(coll46_)
        d_1638_v60_ = (iife140_()
        ).set(p2, (self).f6)
        d_1640_v62_: str
        d_1640_v62_ = _dafny.CodePoint('y')
        (globalState).f5 = (len(_dafny.SeqWithoutIsStrInference([d_1640_v62_]))) + ((self).f6)
        d_1641_v63_: _dafny.Array
        def lambda66_(d_1642_p0_):
            def lambda67_(d_1643_i6_):
                return _dafny.Map({d_1642_p0_: (self.f9)[default__.safeIndex(449, (self.f9).length(0))]})

            return lambda67_

        init44_ = lambda66_(p0)
        nw248_ = _dafny.Array(None, 19)
        for i0_44_ in range(nw248_.length(0)):
            nw248_[i0_44_] = init44_(i0_44_)
        d_1641_v63_ = nw248_
        d_1644_v64_: D9
        d_1644_v64_ = D9_DC26(d_1641_v63_)
        source20_ = d_1644_v64_
        if source20_.is_DC25:
            d_1645___mcc_h0_ = source20_.cf46
            d_1646___mcc_h1_ = source20_.cf47
            d_1647___mcc_h2_ = source20_.cf48
            d_1648___mcc_h3_ = source20_.cf49
            d_1649_cf49_ = d_1648___mcc_h3_
            d_1650_cf48_ = d_1647___mcc_h2_
            d_1651_cf47_ = d_1646___mcc_h1_
            d_1652_cf46_ = d_1645___mcc_h0_
            d_1653_v65_: _dafny.Seq
            d_1653_v65_ = _dafny.SeqWithoutIsStrInference([d_1652_cf46_])
            d_1654_v66_: _dafny.MultiSet
            d_1654_v66_ = _dafny.MultiSet([len(d_1653_v65_), d_1651_cf47_])
            d_1655_v67_: _dafny.Seq
            d_1655_v67_ = _dafny.SeqWithoutIsStrInference([d_1649_cf49_])
            d_1656_v68_: _dafny.Map
            d_1656_v68_ = _dafny.Map({((d_1654_v66_).cardinality) + (d_1651_cf47_): (_dafny.SeqWithoutIsStrInference([d_1649_cf49_ for d_1657_i7_ in range(default__.abs(50))])) + (d_1655_v67_)})
            d_1656_v68_ = (d_1656_v68_).set((d_1653_v65_)[default__.safeIndex((0) - (-709), len(d_1653_v65_))], d_1655_v67_)
            if p0:
                d_1658_v69_: D2
                d_1658_v69_ = D2_DC6(d_1650_cf48_, self.f11)
                d_1659_v70_: _dafny.Seq
                d_1659_v70_ = _dafny.SeqWithoutIsStrInference([(d_1658_v69_).cf14, self.f11])
                d_1659_v70_ = _dafny.SeqWithoutIsStrInference([self.f12, self.f12, self.f11])
                (globalState).f0 = p0
                d_1660_v71_: _dafny.Seq
                d_1660_v71_ = _dafny.SeqWithoutIsStrInference([d_1653_v65_])
                d_1661_v72_: _dafny.Set
                d_1661_v72_ = _dafny.Set({(self.f9)[default__.safeIndex(449, (self.f9).length(0))]})
                d_1662_v73_: _dafny.Array
                nw249_ = _dafny.Array(None, 28)
                nw249_[int(0)] = d_1653_v65_
                nw249_[int(1)] = d_1653_v65_
                nw249_[int(2)] = d_1653_v65_
                nw249_[int(3)] = d_1653_v65_
                nw249_[int(4)] = _dafny.SeqWithoutIsStrInference([(self).f6 for d_1663_i8_ in range(default__.abs(-787))])
                nw249_[int(5)] = d_1653_v65_
                nw249_[int(6)] = d_1653_v65_
                nw249_[int(7)] = (d_1653_v65_).set(default__.safeIndex(default__.fm2(p2, globalState), len(d_1653_v65_)), (self).f6)
                nw249_[int(8)] = d_1653_v65_
                nw249_[int(9)] = d_1653_v65_
                nw249_[int(10)] = _dafny.SeqWithoutIsStrInference([-939 for d_1664_i9_ in range(default__.abs(-670))])
                nw249_[int(11)] = d_1653_v65_
                nw249_[int(12)] = d_1653_v65_
                nw249_[int(13)] = d_1653_v65_
                nw249_[int(14)] = d_1653_v65_
                nw249_[int(15)] = d_1653_v65_
                nw249_[int(16)] = d_1653_v65_
                nw249_[int(17)] = d_1653_v65_
                nw249_[int(18)] = _dafny.SeqWithoutIsStrInference([-584])
                nw249_[int(19)] = _dafny.SeqWithoutIsStrInference([d_1651_cf47_, p2])
                nw249_[int(20)] = d_1653_v65_
                nw249_[int(21)] = d_1653_v65_
                nw249_[int(22)] = d_1653_v65_
                nw249_[int(23)] = ((d_1660_v71_)[default__.safeIndex(p2, len(d_1660_v71_))]).set(default__.safeIndex(d_1651_cf47_, len((d_1660_v71_)[default__.safeIndex(p2, len(d_1660_v71_))])), d_1651_cf47_)
                nw249_[int(24)] = d_1653_v65_
                nw249_[int(25)] = d_1653_v65_
                nw249_[int(26)] = (d_1653_v65_).set(default__.safeIndex(d_1652_cf46_, len(d_1653_v65_)), len(d_1661_v72_))
                nw249_[int(27)] = _dafny.SeqWithoutIsStrInference([(self).f6, p2])
                d_1662_v73_ = nw249_
                d_1665_v74_: _dafny.Map
                d_1665_v74_ = _dafny.Map({d_1662_v73_: d_1655_v67_})
                d_1665_v74_ = (d_1665_v74_).set(d_1662_v73_, d_1655_v67_)
                d_1666_v75_: _dafny.Map
                d_1666_v75_ = _dafny.Map({d_1640_v62_: -779})
                d_1666_v75_ = (d_1666_v75_) | (d_1666_v75_)
                (globalState).f0 = p0
            elif True:
                d_1667_v76_: _dafny.Map
                d_1667_v76_ = _dafny.Map({p2: p0})
                d_1668_v77_: D4
                d_1668_v77_ = D4_DC15(((d_1667_v76_)[(0) - (d_1650_cf48_)] if ((0) - (d_1650_cf48_)) in (d_1667_v76_) else p0), p0, d_1655_v67_)
                r1 = not((d_1668_v77_).cf33)
                d_1669_v78_: _dafny.Array
                nw250_ = _dafny.Array(None, 15)
                nw250_[int(0)] = True
                nw250_[int(1)] = False
                nw250_[int(2)] = p0
                nw250_[int(3)] = not(p0)
                nw250_[int(4)] = not((self.f9)[default__.safeIndex(449, (self.f9).length(0))])
                nw250_[int(5)] = default__.fm3(globalState)
                nw250_[int(6)] = p0
                nw250_[int(7)] = (d_1668_v77_).cf33
                nw250_[int(8)] = (self.f9)[default__.safeIndex(449, (self.f9).length(0))]
                nw250_[int(9)] = p0
                nw250_[int(10)] = (self.f9)[default__.safeIndex(449, (self.f9).length(0))]
                nw250_[int(11)] = p0
                nw250_[int(12)] = p0
                nw250_[int(13)] = p0
                nw250_[int(14)] = (self.f9)[default__.safeIndex(449, (self.f9).length(0))]
                d_1669_v78_ = nw250_
                d_1670_v79_: C3
                nw251_ = C3()
                nw251_.ctor__((d_1638_v60_) | (d_1638_v60_), d_1669_v78_, (d_1656_v68_) | (_dafny.Map({(self).f6: (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "thytt"))).set(default__.safeIndex(d_1652_cf46_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "thytt")))), d_1649_cf49_)})), d_1650_cf48_)
                d_1670_v79_ = nw251_
                (globalState).f5 = d_1651_cf47_
                d_1671_v80_: _dafny.Map
                d_1671_v80_ = _dafny.Map({(self.f9)[default__.safeIndex(449, (self.f9).length(0))]: d_1649_cf49_})
                d_1672_v81_: _dafny.Set
                d_1672_v81_ = _dafny.Set({d_1651_cf47_, (0) - (len(_dafny.SeqWithoutIsStrInference([d_1652_cf46_ for d_1673_i10_ in range(default__.abs(752))])))})
                d_1649_cf49_ = default__.fm0(len(d_1671_v80_), d_1672_v81_, (self.f9)[default__.safeIndex(449, (self.f9).length(0))], globalState)
                d_1674_v82_: _dafny.MultiSet
                d_1674_v82_ = _dafny.MultiSet([(d_1654_v66_).set(772, default__.abs((0) - (len(_dafny.SeqWithoutIsStrInference([d_1649_cf49_ for d_1675_i11_ in range(default__.abs(734))]))))), d_1654_v66_])
                d_1676_v83_: T0
                nw252_ = C2()
                nw252_.ctor__(d_1652_cf46_, d_1674_v82_, d_1669_v78_, d_1656_v68_)
                d_1676_v83_ = nw252_
                d_1677_v84_: _dafny.Seq
                d_1677_v84_ = _dafny.SeqWithoutIsStrInference([d_1676_v83_, d_1676_v83_, d_1676_v83_])
                d_1678_v85_: T1
                nw253_ = C3()
                nw253_.ctor__((d_1638_v60_) | (d_1638_v60_), (d_1669_v78_ if p0 else d_1669_v78_), ((self).f10 if False else _dafny.Map({len(d_1655_v67_): (d_1655_v67_).set(default__.safeIndex(len(d_1677_v84_), len(d_1655_v67_)), _dafny.CodePoint('j'))})), (0) - ((d_1636_v58_).cardinality))
                d_1678_v85_ = nw253_
                nw254_ = C2()
                nw254_.ctor__((d_1676_v83_).f6, d_1674_v82_, d_1678_v85_.f9, ((self).f10) | ((d_1678_v85_).f10))
                d_1678_v85_ = nw254_
            arr50_ = self.f9
            index240_ = default__.safeIndex(449, (self.f9).length(0))
            arr50_[index240_] = ((93 if default__.fm3(globalState) else p2)) != (d_1651_cf47_)
            arr51_ = self.f9
            index241_ = default__.safeIndex(449, (self.f9).length(0))
            arr51_[index241_] = not(not(((len(p1)) - (d_1650_cf48_)) == ((d_1636_v58_).cardinality)))
        elif source20_.is_DC26:
            d_1679___mcc_h4_ = source20_.cf50
            d_1680_cf50_ = d_1679___mcc_h4_
            d_1681_v86_: _dafny.Seq
            d_1681_v86_ = _dafny.SeqWithoutIsStrInference([d_1640_v62_])
            d_1682_v87_: _dafny.Set
            d_1682_v87_ = _dafny.Set({p0, True})
            d_1683_v88_: D11
            d_1683_v88_ = D11_DC33(p2, p0, len(d_1681_v86_), (self.f9)[default__.safeIndex(449, (self.f9).length(0))], d_1682_v87_)
            if (d_1683_v88_).cf65:
                (globalState).f5 = (self).f6
                d_1684_v89_: _dafny.Map
                d_1684_v89_ = _dafny.Map({len(d_1681_v86_): d_1640_v62_})
                d_1685_v90_: _dafny.Map
                d_1685_v90_ = _dafny.Map({len(d_1684_v89_): (self.f9)[default__.safeIndex(449, (self.f9).length(0))]})
                d_1686_v91_: _dafny.Map
                d_1686_v91_ = _dafny.Map({(self.f9)[default__.safeIndex(449, (self.f9).length(0))]: (self.f9)[default__.safeIndex(449, (self.f9).length(0))]})
                d_1685_v90_ = (d_1685_v90_).set(p2, ((self.f9)[default__.safeIndex(449, (self.f9).length(0))]) in (d_1686_v91_))
                (globalState).f0 = (p0) and (((self.f9)[default__.safeIndex(449, (self.f9).length(0))]) not in (p1))
                rhs281_ = p2
                rhs282_ = ((D11_DC31(-330, len(_dafny.Map({(self).f6: (self.f9)[default__.safeIndex(449, (self.f9).length(0))]})))).cf59) * ((0) - (len(d_1682_v87_)))
                lhs234_ = globalState
                lhs235_ = globalState
                lhs234_.f5 = rhs281_
                lhs235_.f5 = rhs282_
                (globalState).f5 = (self).f6
            elif True:
                d_1687_v92_: _dafny.Array
                nw255_ = _dafny.Array(D2.default()(), 26)
                d_1687_v92_ = nw255_
                d_1688_v93_: D2
                d_1688_v93_ = D2_DC8((self.f9)[default__.safeIndex(449, (self.f9).length(0))], (self.f9)[default__.safeIndex(449, (self.f9).length(0))])
                index242_ = default__.safeIndex(386, (d_1687_v92_).length(0))
                (d_1687_v92_)[index242_] = d_1688_v93_
                index243_ = default__.safeIndex(386, (d_1687_v92_).length(0))
                (d_1687_v92_)[index243_] = (d_1688_v93_ if p0 else d_1688_v93_)
                (globalState).f5 = (p2) + ((self).f6)
                d_1689_v94_: C1
                nw256_ = C1()
                nw256_.ctor__(len(d_1681_v86_))
                d_1689_v94_ = nw256_
                d_1690_v95_: _dafny.Array
                nw257_ = _dafny.Array(_dafny.Array(None, 0), 9)
                d_1690_v95_ = nw257_
                arr52_ = self.f11
                index244_ = default__.safeIndex(471, (self.f11).length(0))
                arr52_[index244_] = (0) - (p2)
                d_1691_v96_: _dafny.MultiSet
                d_1691_v96_ = _dafny.MultiSet([p1, (p1).set(default__.safeIndex(len((d_1681_v86_).set(default__.safeIndex(p2, len(d_1681_v86_)), d_1640_v62_)), len(p1)), (self.f9)[default__.safeIndex(449, (self.f9).length(0))]), p1, p1])
                arr53_ = self.f11
                index245_ = default__.safeIndex(471, (self.f11).length(0))
                rhs283_ = d_1690_v95_
                rhs284_ = (((d_1691_v96_).set(p1, default__.abs(569)) if default__.fm3(globalState) else d_1691_v96_)).cardinality
                lhs236_ = self.f11
                lhs237_ = default__.safeIndex(471, (self.f11).length(0))
                d_1690_v95_ = rhs283_
                lhs236_[lhs237_] = rhs284_
                d_1692_v97_: _dafny.Seq
                d_1692_v97_ = _dafny.SeqWithoutIsStrInference([self.f12])
                d_1693_v98_: _dafny.Seq
                d_1693_v98_ = _dafny.SeqWithoutIsStrInference([(self).f6, p2])
                d_1694_v99_: _dafny.Array
                nw258_ = _dafny.Array(None, 9)
                nw258_[int(0)] = self.f12
                nw258_[int(1)] = self.f12
                nw258_[int(2)] = self.f12
                nw258_[int(3)] = self.f12
                nw258_[int(4)] = self.f12
                nw258_[int(5)] = self.f12
                nw258_[int(6)] = self.f12
                nw258_[int(7)] = self.f12
                nw258_[int(8)] = (d_1692_v97_)[default__.safeIndex((0) - ((0) - ((0) - (len(d_1693_v98_)))), len(d_1692_v97_))]
                d_1694_v99_ = nw258_
                d_1695_v101_: _dafny.Set
                d_1695_v101_ = _dafny.Set({(self.f11)[default__.safeIndex(471, (self.f11).length(0))]})
                d_1696_v102_: _dafny.Seq
                def iife141_():
                    coll47_ = _dafny.Set()
                    compr_47_: int
                    for compr_47_ in _dafny.IntegerRange(-932, -282):
                        d_1697_v100_: int = compr_47_
                        if ((-932) <= (d_1697_v100_)) and ((d_1697_v100_) < (-282)):
                            coll47_ = coll47_.union(_dafny.Set([default__.safeModuloInt(d_1697_v100_, p2)]))
                    return _dafny.Set(coll47_)
                d_1696_v102_ = _dafny.SeqWithoutIsStrInference([(iife141_()
                ).ispropersubset(d_1695_v101_), (self.f9)[default__.safeIndex(449, (self.f9).length(0))], ((self).f6) != (-21)])
                d_1698_v103_: _dafny.Array
                nw259_ = _dafny.Array(_dafny.Seq({}), 2)
                d_1698_v103_ = nw259_
                arr54_ = self.f11
                index246_ = default__.safeIndex(471, (self.f11).length(0))
                rhs285_ = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "psadyloiw")))
                rhs286_ = d_1694_v99_
                rhs287_ = default__.fm4(p2, (-846) * (17), ((self.f11)[default__.safeIndex(471, (self.f11).length(0))] if p0 else (self).f6), globalState)
                rhs288_ = d_1698_v103_
                lhs238_ = self.f11
                lhs239_ = default__.safeIndex(471, (self.f11).length(0))
                lhs238_[lhs239_] = rhs285_
                d_1694_v99_ = rhs286_
                d_1696_v102_ = rhs287_
                d_1698_v103_ = rhs288_
            d_1699_v104_: _dafny.MultiSet
            d_1699_v104_ = _dafny.MultiSet([105])
            d_1700_v105_: _dafny.MultiSet
            d_1700_v105_ = _dafny.MultiSet([d_1699_v104_, d_1699_v104_])
            d_1701_v106_: _dafny.Array
            nw260_ = _dafny.Array(None, 8)
            nw260_[int(0)] = p0
            nw260_[int(1)] = p0
            nw260_[int(2)] = p0
            nw260_[int(3)] = (self.f9)[default__.safeIndex(449, (self.f9).length(0))]
            nw260_[int(4)] = False
            nw260_[int(5)] = p0
            nw260_[int(6)] = p0
            nw260_[int(7)] = (self.f9)[default__.safeIndex(449, (self.f9).length(0))]
            d_1701_v106_ = nw260_
            d_1702_v107_: T0
            nw261_ = C2()
            nw261_.ctor__((self).f6, d_1700_v105_, d_1701_v106_, (self).f10)
            d_1702_v107_ = nw261_
            d_1703_v108_: _dafny.Set
            d_1703_v108_ = _dafny.Set({d_1702_v107_, d_1702_v107_, d_1702_v107_})
            d_1704_v109_: _dafny.Map
            d_1704_v109_ = _dafny.Map({d_1703_v108_: _dafny.CodePoint('g')})
            d_1704_v109_ = (d_1704_v109_ if default__.fm3(globalState) else d_1704_v109_)
            (globalState).f5 = p2
            arr55_ = self.f9
            index247_ = default__.safeIndex(449, (self.f9).length(0))
            arr55_[index247_] = not(p0)
        elif source20_.is_DC24:
            d_1705___mcc_h5_ = source20_.cf45
            d_1706_cf45_ = d_1705___mcc_h5_
            arr56_ = self.f9
            index248_ = default__.safeIndex(449, (self.f9).length(0))
            arr56_[index248_] = p0
            d_1707_v110_: _dafny.Set
            d_1707_v110_ = _dafny.Set({p0, p0, p0})
            d_1708_v111_: _dafny.Array
            nw262_ = _dafny.Array(None, 6)
            nw262_[int(0)] = False
            nw262_[int(1)] = (self.f9)[default__.safeIndex(449, (self.f9).length(0))]
            nw262_[int(2)] = p0
            nw262_[int(3)] = not((d_1707_v110_).ispropersubset(d_1707_v110_))
            nw262_[int(4)] = False
            nw262_[int(5)] = (len(p1)) == ((self).f6)
            d_1708_v111_ = nw262_
            d_1709_v112_: _dafny.Seq
            d_1709_v112_ = _dafny.SeqWithoutIsStrInference([(d_1708_v111_ if p0 else d_1708_v111_)])
            d_1708_v111_ = (d_1709_v112_)[default__.safeIndex((self).f6, len(d_1709_v112_))]
            (globalState).f5 = (self).f6
            if ((self).f6) <= ((self).f6):
                d_1710_v113_: C0
                nw263_ = C0()
                nw263_.ctor__(d_1708_v111_, (self.f9)[default__.safeIndex(449, (self.f9).length(0))])
                d_1710_v113_ = nw263_
                d_1711_v114_: _dafny.Seq
                d_1711_v114_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hqae"))
                d_1712_v115_: D9
                d_1712_v115_ = D9_DC25(p2, default__.safeDivisionInt(len(d_1711_v114_), p2), (0) - (len(_dafny.SeqWithoutIsStrInference([D3_DC9(_dafny.SeqWithoutIsStrInference([p2 for d_1713_i13_ in range(default__.abs(266))])) for d_1714_i12_ in range(default__.abs(125))]))), d_1640_v62_)
                d_1712_v115_ = d_1712_v115_
                d_1715_v116_: _dafny.Seq
                d_1715_v116_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({p0})])
                d_1716_v117_: _dafny.Map
                d_1716_v117_ = _dafny.Map({((d_1715_v116_)[default__.safeIndex(len(d_1711_v114_), len(d_1715_v116_))]) | (d_1707_v110_): not((self.f9)[default__.safeIndex(449, (self.f9).length(0))])})
                d_1716_v117_ = (d_1716_v117_).set(d_1707_v110_, d_1710_v113_.f20)
                d_1717_v118_: _dafny.Map
                d_1717_v118_ = _dafny.Map({p2: p0})
                def iife142_():
                    coll48_ = _dafny.Map()
                    compr_48_: int
                    for compr_48_ in (_dafny.Map({p2: len(d_1638_v60_)})).keys.Elements:
                        d_1718_v119_: int = compr_48_
                        if (d_1718_v119_) in (_dafny.Map({p2: len(d_1638_v60_)})):
                            coll48_[default__.safeDivisionInt(d_1718_v119_, p2)] = p0
                    return _dafny.Map(coll48_)
                d_1717_v118_ = (d_1717_v118_) | ((d_1717_v118_) | (iife142_()
                ))
                d_1719_v120_: _dafny.Seq
                d_1719_v120_ = _dafny.SeqWithoutIsStrInference([p2])
                d_1720_v121_: _dafny.Map
                d_1720_v121_ = _dafny.Map({d_1710_v113_.f20: d_1710_v113_.f20})
                (globalState).f5 = (((d_1719_v120_)[default__.safeIndex(len(d_1720_v121_), len(d_1719_v120_))]) * ((self).f6)) * (p2)
            elif True:
                d_1721_v122_: _dafny.MultiSet
                d_1721_v122_ = _dafny.MultiSet([p2, p2])
                arr57_ = self.f12
                index249_ = default__.safeIndex(45, (self.f12).length(0))
                arr57_[index249_] = (((d_1721_v122_)[len(_dafny.SeqWithoutIsStrInference([(self).f6 for d_1722_i14_ in range(default__.abs(998))]))] if (len(_dafny.SeqWithoutIsStrInference([(self).f6 for d_1723_i14_ in range(default__.abs(998))]))) in (d_1721_v122_) else p2)) - (p2)
                arr58_ = self.f12
                index250_ = default__.safeIndex(45, (self.f12).length(0))
                arr58_[index250_] = (self).f6
                d_1724_v123_: C6
                nw264_ = C6()
                nw264_.ctor__((self).f6, ((self.f9)[default__.safeIndex(449, (self.f9).length(0))] if (self.f9)[default__.safeIndex(449, (self.f9).length(0))] else (self.f9)[default__.safeIndex(449, (self.f9).length(0))]))
                d_1724_v123_ = nw264_
                d_1725_v124_: _dafny.Array
                nw265_ = _dafny.Array(D6.default()(), 26)
                d_1725_v124_ = nw265_
                d_1726_v125_: _dafny.Array
                d_1726_v125_ = d_1725_v124_
                d_1725_v124_ = (d_1726_v125_)
                d_1727_v126_: _dafny.Map
                d_1727_v126_ = (self).f10
                d_1728_v127_: C3
                nw266_ = C3()
                nw266_.ctor__(d_1638_v60_, d_1708_v111_, ((d_1727_v126_)) | ((self).f10), len((((self).f10)[(d_1724_v123_).f13] if ((d_1724_v123_).f13) in ((self).f10) else _dafny.SeqWithoutIsStrInference([d_1640_v62_ for d_1729_i15_ in range(default__.abs(646))]))))
                d_1728_v127_ = nw266_
                d_1730_v128_: C5
                nw267_ = C5()
                nw267_.ctor__((self.f12)[default__.safeIndex(45, (self.f12).length(0))], (d_1724_v123_).f13)
                d_1730_v128_ = nw267_
        elif True:
            d_1731___mcc_h6_ = source20_.cf51
            d_1732_cf51_ = d_1731___mcc_h6_
            d_1733_v129_: _dafny.Array
            def lambda68_(d_1734_p1_, d_1735_p2_):
                def lambda69_(d_1736_i16_):
                    return D0_DC0(d_1734_p1_, d_1735_p2_)

                return lambda69_

            init45_ = lambda68_(p1, p2)
            nw268_ = _dafny.Array(None, 16)
            for i0_45_ in range(nw268_.length(0)):
                nw268_[i0_45_] = init45_(i0_45_)
            d_1733_v129_ = nw268_
            d_1737_v130_: D7
            d_1737_v130_ = D7_DC21(d_1733_v129_)
            d_1737_v130_ = d_1737_v130_
            d_1738_v131_: _dafny.Seq
            d_1738_v131_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dcbeliug"))
            d_1739_v132_: C5
            nw269_ = C5()
            nw269_.ctor__((0) - (len(d_1738_v131_)), 510)
            d_1739_v132_ = nw269_
            d_1740_v133_: _dafny.Map
            d_1740_v133_ = _dafny.Map({d_1739_v132_: self.f12})
            rhs289_ = (p2) * ((self).f6)
            rhs290_ = (self).f6
            rhs291_ = (d_1740_v133_) | (d_1740_v133_)
            lhs240_ = globalState
            lhs241_ = globalState
            lhs240_.f5 = rhs289_
            lhs241_.f5 = rhs290_
            d_1740_v133_ = rhs291_
            d_1741_v134_: _dafny.Map
            d_1741_v134_ = _dafny.Map({p2: (not((self.f9)[default__.safeIndex(449, (self.f9).length(0))]) if (self.f9)[default__.safeIndex(449, (self.f9).length(0))] else p0)})
            d_1741_v134_ = (d_1741_v134_).set((d_1739_v132_).f15, (self.f9)[default__.safeIndex(449, (self.f9).length(0))])
            arr59_ = self.f11
            index251_ = default__.safeIndex(990, (self.f11).length(0))
            arr59_[index251_] = (d_1739_v132_).f15
            d_1742_v135_: _dafny.Seq
            d_1742_v135_ = _dafny.SeqWithoutIsStrInference([self.f12, self.f11])
            d_1743_v136_: _dafny.Array
            nw270_ = _dafny.Array(False, 14)
            d_1743_v136_ = nw270_
            d_1744_v137_: _dafny.Seq
            d_1744_v137_ = _dafny.SeqWithoutIsStrInference([(d_1739_v132_).f15, (self).f6, (0) - ((d_1739_v132_).f15)])
            d_1745_v138_: _dafny.MultiSet
            d_1745_v138_ = _dafny.MultiSet([d_1744_v137_, d_1744_v137_])
            d_1746_v139_: _dafny.Seq
            d_1746_v139_ = _dafny.SeqWithoutIsStrInference([d_1744_v137_])
            arr60_ = self.f11
            index252_ = default__.safeIndex(990, (self.f11).length(0))
            rhs292_ = (0) - ((d_1739_v132_).f15)
            rhs293_ = (d_1738_v131_ if p0 else (d_1738_v131_).set(default__.safeIndex((self).fm14(globalState), len(d_1738_v131_)), d_1640_v62_))
            rhs294_ = (d_1742_v135_ if (_dafny.MultiSet(d_1746_v139_)).issubset(d_1745_v138_) else d_1742_v135_)
            rhs295_ = (d_1739_v132_).f15
            rhs296_ = d_1743_v136_
            lhs242_ = self.f11
            lhs243_ = default__.safeIndex(990, (self.f11).length(0))
            lhs244_ = globalState
            lhs245_ = globalState
            lhs242_[lhs243_] = rhs292_
            lhs244_.f2 = rhs293_
            d_1742_v135_ = rhs294_
            lhs245_.f5 = rhs295_
            d_1743_v136_ = rhs296_
        r0 = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qvp")) for d_1747_i17_ in range(default__.abs(780))])
        d_1748_v140_: _dafny.Set
        d_1748_v140_ = _dafny.Set({(self.f9)[default__.safeIndex(449, (self.f9).length(0))], (self.f9)[default__.safeIndex(449, (self.f9).length(0))]})
        d_1749_v141_: _dafny.MultiSet
        d_1749_v141_ = _dafny.MultiSet([976, len(d_1748_v140_)])
        d_1750_v142_: D11
        d_1750_v142_ = D11_DC31((self).f6, (self).f6)
        d_1751_v143_: _dafny.Map
        d_1751_v143_ = _dafny.Map({d_1749_v141_: (p1)[default__.safeIndex((d_1750_v142_).cf58, len(p1))]})
        d_1752_v144_: D5
        d_1752_v144_ = D5_DC16(d_1748_v140_)
        d_1753_v145_: D11
        d_1753_v145_ = D11_DC33(len(d_1751_v143_), p0, (self).f6, (self.f9)[default__.safeIndex(449, (self.f9).length(0))], (d_1752_v144_).cf35)
        r1 = (d_1753_v145_).cf63
        return r0, r1

    def m7(self, p0, p1, globalState):
        r0: str = _dafny.CodePoint('D')
        r1: int = int(0)
        r2: _dafny.Array = _dafny.Array(None, 0)
        r3: int = int(0)
        (globalState).f5 = (self).f6
        d_1754_v0_: bool
        d_1754_v0_ = False
        d_1755_v1_: _dafny.Seq
        d_1755_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ldvfwsx"))
        (globalState).f0 = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gsjgptuvl"))) <= (d_1755_v1_) if d_1754_v0_ else d_1754_v0_)
        (globalState).f0 = default__.fm3(globalState)
        d_1754_v0_ = not(d_1754_v0_)
        d_1756_v2_: D7
        d_1756_v2_ = D7_DC22()
        d_1757_v3_: _dafny.Seq
        d_1757_v3_ = _dafny.SeqWithoutIsStrInference([d_1754_v0_])
        d_1758_v4_: _dafny.Map
        d_1758_v4_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_1754_v0_]): d_1754_v0_})
        d_1759_v5_: D1
        d_1759_v5_ = D1_DC4(D0_DC0(d_1757_v3_, len(d_1758_v4_)), default__.fm2((self).f6, globalState), default__.fm18(d_1754_v0_, globalState), d_1754_v0_, (0) - ((self).f6))
        rhs297_ = (d_1759_v5_).cf10
        rhs298_ = d_1756_v2_
        rhs299_ = d_1754_v0_
        rhs300_ = default__.fm3(globalState)
        lhs246_ = globalState
        lhs247_ = globalState
        lhs248_ = globalState
        lhs246_.f0 = rhs297_
        d_1756_v2_ = rhs298_
        lhs247_.f0 = rhs299_
        lhs248_.f0 = rhs300_
        d_1760_v6_: str
        d_1760_v6_ = _dafny.CodePoint('l')
        d_1761_v7_: _dafny.Map
        d_1761_v7_ = _dafny.Map({287: (0) - (len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ljwkjfnhn"))).set(default__.safeIndex(458, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ljwkjfnhn")))), d_1760_v6_)))})
        d_1762_v9_: _dafny.Map
        d_1762_v9_ = _dafny.Map({p1: _dafny.Set({d_1760_v6_})})
        d_1763_v10_: C3
        nw271_ = C3()
        def iife143_():
            coll49_ = _dafny.Map()
            compr_49_: int
            for compr_49_ in (d_1762_v9_).keys.Elements:
                d_1764_v8_: int = compr_49_
                if (d_1764_v8_) in (d_1762_v9_):
                    coll49_[(d_1764_v8_) + (p1)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f"))
            return _dafny.Map(coll49_)
        nw271_.ctor__(d_1761_v7_, self.f9, iife143_()
        , p1)
        d_1763_v10_ = nw271_
        r0 = d_1760_v6_
        d_1765_v11_: _dafny.MultiSet
        d_1765_v11_ = _dafny.MultiSet([d_1754_v0_])
        r1 = (default__.safeModuloInt((self).f6, (self).f6) if d_1754_v0_ else ((self).f6) + ((0) - (((d_1765_v11_)[False] if (False) in (d_1765_v11_) else 788))))
        nw272_ = _dafny.Array(False, 12)
        r2 = nw272_
        r3 = p1
        return r0, r1, r2, r3


class C8:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C8"
    def ctor__(self):
        pass
        pass

    def fm10(self, p0, p1, globalState):
        return (_dafny.MultiSet([697, 243, 884, -12])) - (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([False])), 835]))

    def fm11(self, globalState):
        return default__.safeDivisionInt((len(_dafny.SeqWithoutIsStrInference([D1_DC4(D0_DC0(_dafny.SeqWithoutIsStrInference([not(True), False]), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "skr")))), 887, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lqerk")), False, 739)]))) - (-290), len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ctlfngak"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "snnrnqplc")))))

    def m4(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Map = _dafny.Map({})
        r1: _dafny.Map = _dafny.Map({})
        r2: int = int(0)
        d_1766_v0_: _dafny.Map
        d_1766_v0_ = _dafny.Map({p3: p0})
        d_1767_v1_: _dafny.Array
        nw273_ = _dafny.Array(None, 5)
        nw273_[int(0)] = p1
        nw273_[int(1)] = default__.fm3(globalState)
        nw273_[int(2)] = (p0) and (p3)
        nw273_[int(3)] = p0
        nw273_[int(4)] = (((d_1766_v0_)[p0] if (p0) in (d_1766_v0_) else p0) if not(((d_1766_v0_)[p3] if (p3) in (d_1766_v0_) else default__.fm3(globalState))) else p3)
        d_1767_v1_ = nw273_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_1767_v1_).length(0)):
            d_1768_i0_: int = guard_loop_2_
            if (True) and (((0) <= (d_1768_i0_)) and ((d_1768_i0_) < ((d_1767_v1_).length(0)))):
                (d_1767_v1_)[(d_1768_i0_)] = (True if p3 else (p0) == (True))
        d_1769_v2_: _dafny.Seq
        d_1769_v2_ = _dafny.SeqWithoutIsStrInference([p2])
        d_1770_v3_: _dafny.MultiSet
        d_1770_v3_ = _dafny.MultiSet([d_1769_v2_, d_1769_v2_])
        if (d_1770_v3_).isdisjoint(d_1770_v3_):
            d_1771_v4_: str
            d_1771_v4_ = _dafny.CodePoint('x')
            d_1772_v5_: D0
            d_1772_v5_ = D0_DC1(not(not(p0)), 172, d_1771_v4_)
            d_1773_v6_: _dafny.Set
            d_1773_v6_ = _dafny.Set({(self).fm11(globalState), p2, p2})
            d_1774_v8_: _dafny.Seq
            d_1774_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mjcroura"))
            d_1775_v9_: _dafny.Map
            d_1775_v9_ = _dafny.Map({d_1771_v4_: len(d_1774_v8_)})
            d_1776_v10_: _dafny.Seq
            d_1776_v10_ = _dafny.SeqWithoutIsStrInference([d_1775_v9_])
            def iife144_():
                coll50_ = _dafny.Set()
                compr_50_: int
                for compr_50_ in _dafny.IntegerRange(-492, 679):
                    d_1777_v7_: int = compr_50_
                    if ((-492) <= (d_1777_v7_)) and ((d_1777_v7_) < (679)):
                        coll50_ = coll50_.union(_dafny.Set([(d_1777_v7_) - (len(_dafny.SeqWithoutIsStrInference([(d_1772_v5_).cf2, p1])))]))
                return _dafny.Set(coll50_)
            rhs301_ = (d_1773_v6_).isdisjoint(iife144_()
            )
            rhs302_ = D0_DC1(p3, p2, (d_1772_v5_).cf4)
            rhs303_ = (_dafny.CodePoint('k')) in ((d_1776_v10_)[default__.safeIndex(p2, len(d_1776_v10_))])
            rhs304_ = p0
            lhs249_ = globalState
            lhs250_ = globalState
            lhs251_ = globalState
            lhs249_.f0 = rhs301_
            d_1772_v5_ = rhs302_
            lhs250_.f0 = rhs303_
            lhs251_.f0 = rhs304_
            d_1778_v11_: _dafny.Array
            nw274_ = _dafny.Array(int(0), 26)
            d_1778_v11_ = nw274_
            index253_ = default__.safeIndex(2, (d_1778_v11_).length(0))
            (d_1778_v11_)[index253_] = p2
            index254_ = default__.safeIndex(2, (d_1778_v11_).length(0))
            (d_1778_v11_)[index254_] = (p2) * (240)
            pat_let_tv34_ = p3
            def iife145_(_pat_let47_0):
                def iife146_(d_1779_dt__update__tmp_h0_):
                    def iife147_(_pat_let48_0):
                        def iife148_(d_1780_dt__update_hcf2_h0_):
                            return D0_DC1(d_1780_dt__update_hcf2_h0_, (d_1779_dt__update__tmp_h0_).cf3, (d_1779_dt__update__tmp_h0_).cf4)
                        return iife148_(_pat_let48_0)
                    return iife147_(pat_let_tv34_)
                return iife146_(_pat_let47_0)
            d_1771_v4_ = (iife145_(D0_DC1(p1, (d_1778_v11_)[default__.safeIndex(2, (d_1778_v11_).length(0))], d_1771_v4_))).cf4
            (globalState).f5 = (d_1778_v11_)[default__.safeIndex(2, (d_1778_v11_).length(0))]
            index255_ = default__.safeIndex(2, (d_1778_v11_).length(0))
            (d_1778_v11_)[index255_] = (self).fm11(globalState)
        elif True:
            d_1781_v12_: _dafny.Array
            def lambda70_(d_1782_v2_):
                def lambda71_(d_1783_i1_):
                    return d_1782_v2_

                return lambda71_

            init46_ = lambda70_(d_1769_v2_)
            nw275_ = _dafny.Array(None, 6)
            for i0_46_ in range(nw275_.length(0)):
                nw275_[i0_46_] = init46_(i0_46_)
            d_1781_v12_ = nw275_
            d_1781_v12_ = (d_1781_v12_ if not(p0) else d_1781_v12_)
            d_1784_v13_: _dafny.Array
            nw276_ = _dafny.Array(int(0), 12)
            d_1784_v13_ = nw276_
            d_1784_v13_ = d_1784_v13_
            d_1785_v14_: D2
            d_1785_v14_ = D2_DC6(11, d_1784_v13_)
            d_1785_v14_ = d_1785_v14_
            d_1769_v2_ = d_1769_v2_
            index256_ = default__.safeIndex(373, (d_1767_v1_).length(0))
            (d_1767_v1_)[index256_] = p1
            index257_ = default__.safeIndex(373, (d_1767_v1_).length(0))
            (d_1767_v1_)[index257_] = p0
        (globalState).f2 = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hrnahy"))
        d_1786_v15_: D2
        d_1786_v15_ = D2_DC8(p0, not(False))
        pat_let_tv35_ = d_1766_v0_
        pat_let_tv36_ = d_1766_v0_
        pat_let_tv37_ = p1
        pat_let_tv38_ = d_1766_v0_
        def lambda72_(source21_):
            if source21_.is_DC6:
                d_1787___mcc_h0_ = source21_.cf13
                d_1788___mcc_h1_ = source21_.cf14
                d_1789_cf14_ = d_1788___mcc_h1_
                d_1790_cf13_ = d_1787___mcc_h0_
                return pat_let_tv35_
            elif source21_.is_DC7:
                d_1791___mcc_h2_ = source21_.cf15
                d_1792___mcc_h3_ = source21_.cf16
                d_1793___mcc_h4_ = source21_.cf17
                d_1794___mcc_h5_ = source21_.cf18
                d_1795_cf18_ = d_1794___mcc_h5_
                d_1796_cf17_ = d_1793___mcc_h4_
                d_1797_cf16_ = d_1792___mcc_h3_
                d_1798_cf15_ = d_1791___mcc_h2_
                return pat_let_tv36_
            elif source21_.is_DC8:
                d_1799___mcc_h6_ = source21_.cf19
                d_1800___mcc_h7_ = source21_.cf20
                d_1801_cf20_ = d_1800___mcc_h7_
                d_1802_cf19_ = d_1799___mcc_h6_
                return _dafny.Map({pat_let_tv37_: d_1801_cf20_})
            elif True:
                d_1803___mcc_h8_ = source21_.cf12
                d_1804_cf12_ = d_1803___mcc_h8_
                return pat_let_tv38_

        d_1766_v0_ = lambda72_(d_1786_v15_)
        d_1805_v16_: str
        d_1805_v16_ = _dafny.CodePoint('i')
        d_1806_v17_: _dafny.Seq
        d_1806_v17_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m"))
        source22_ = default__.fm12(p2, not((d_1805_v16_) in (d_1806_v17_)), not(not(not(p0))), globalState)
        if source22_.is_DC6:
            d_1807___mcc_h9_ = source22_.cf13
            d_1808___mcc_h10_ = source22_.cf14
            d_1809_cf14_ = d_1808___mcc_h10_
            d_1810_cf13_ = d_1807___mcc_h9_
            d_1811_v18_: _dafny.Map
            d_1811_v18_ = _dafny.Map({p1: p2})
            d_1812_v19_: _dafny.Map
            d_1812_v19_ = _dafny.Map({d_1805_v16_: False})
            d_1813_v21_: _dafny.MultiSet
            d_1813_v21_ = _dafny.MultiSet([d_1805_v16_])
            d_1814_v22_: _dafny.Seq
            def iife149_():
                coll51_ = _dafny.Map()
                compr_51_: str
                for compr_51_ in (d_1813_v21_).Elements:
                    d_1815_v20_: str = compr_51_
                    if (d_1815_v20_) in (d_1813_v21_):
                        coll51_[d_1815_v20_] = not(p1)
                return _dafny.Map(coll51_)
            d_1814_v22_ = _dafny.SeqWithoutIsStrInference([d_1812_v19_, iife149_()
            ])
            d_1811_v18_ = (d_1811_v18_).set(p0, len(d_1814_v22_))
            d_1810_cf13_ = (default__.safeDivisionInt(p2, 650)) * (91)
            d_1816_v23_: _dafny.Array
            nw277_ = _dafny.Array(_dafny.Map({}), 16)
            d_1816_v23_ = nw277_
            d_1817_v24_: _dafny.Map
            d_1817_v24_ = _dafny.Map({len(d_1806_v17_): p1})
            d_1818_v25_: _dafny.Seq
            d_1818_v25_ = _dafny.SeqWithoutIsStrInference([d_1817_v24_])
            index258_ = default__.safeIndex(479, (d_1816_v23_).length(0))
            (d_1816_v23_)[index258_] = (d_1818_v25_)[default__.safeIndex(p2, len(d_1818_v25_))]
            index259_ = default__.safeIndex(479, (d_1816_v23_).length(0))
            (d_1816_v23_)[index259_] = d_1817_v24_
            d_1819_v26_: _dafny.Set
            d_1819_v26_ = _dafny.Set({True, p1})
            if (_dafny.Set({p1, not(p3)})).issubset(d_1819_v26_):
                d_1820_v27_: _dafny.Map
                d_1820_v27_ = _dafny.Map({d_1810_cf13_: d_1805_v16_})
                d_1820_v27_ = (d_1820_v27_).set(d_1810_cf13_, d_1805_v16_)
                d_1821_v28_: _dafny.Array
                def lambda73_(d_1822_v2_, d_1823_p2_):
                    def lambda74_(d_1824_i2_):
                        return (d_1822_v2_) + (_dafny.SeqWithoutIsStrInference([d_1823_p2_ for d_1825_i3_ in range(default__.abs(-622))]))

                    return lambda74_

                init47_ = lambda73_(d_1769_v2_, p2)
                nw278_ = _dafny.Array(None, 7)
                for i0_47_ in range(nw278_.length(0)):
                    nw278_[i0_47_] = init47_(i0_47_)
                d_1821_v28_ = nw278_
                rhs305_ = ((d_1810_cf13_) * (d_1810_cf13_)) < (p2)
                rhs306_ = (d_1810_cf13_) - (len(d_1806_v17_))
                rhs307_ = d_1821_v28_
                lhs252_ = globalState
                lhs252_.f0 = rhs305_
                r2 = rhs306_
                d_1821_v28_ = rhs307_
                (globalState).f5 = (d_1810_cf13_) * ((0) - (d_1810_cf13_))
                (globalState).f0 = p3
                d_1826_v29_: C0
                nw279_ = C0()
                nw279_.ctor__(d_1767_v1_, p1)
                d_1826_v29_ = nw279_
            elif True:
                (globalState).f2 = d_1806_v17_
                (globalState).f0 = p1
                (globalState).f0 = p3
                (globalState).f0 = (p1 if p3 else p0)
                d_1827_v30_: C6
                nw280_ = C6()
                nw280_.ctor__((p2 if False else p2), p0)
                d_1827_v30_ = nw280_
        elif source22_.is_DC7:
            d_1828___mcc_h11_ = source22_.cf15
            d_1829___mcc_h12_ = source22_.cf16
            d_1830___mcc_h13_ = source22_.cf17
            d_1831___mcc_h14_ = source22_.cf18
            d_1832_cf18_ = d_1831___mcc_h14_
            d_1833_cf17_ = d_1830___mcc_h13_
            d_1834_cf16_ = d_1829___mcc_h12_
            d_1835_cf15_ = d_1828___mcc_h11_
            if p0:
                d_1836_v31_: _dafny.Map
                d_1836_v31_ = _dafny.Map({537: p2})
                d_1837_v32_: C3
                nw281_ = C3()
                nw281_.ctor__(_dafny.Map({len(d_1806_v17_): p2}), d_1767_v1_, _dafny.Map({426: d_1806_v17_}), -162)
                d_1837_v32_ = nw281_
                d_1838_v33_: _dafny.Seq
                d_1838_v33_ = _dafny.SeqWithoutIsStrInference([d_1837_v32_, d_1837_v32_])
                d_1839_v34_: _dafny.Map
                d_1839_v34_ = _dafny.Map({len((d_1836_v31_).set(len(d_1838_v33_), -92)): p3})
                d_1840_v35_: _dafny.Map
                d_1840_v35_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dppy")): d_1832_cf18_})
                index260_ = default__.safeIndex(479, (d_1767_v1_).length(0))
                (d_1767_v1_)[index260_] = ((d_1839_v34_)[len(d_1840_v35_)] if (len(d_1840_v35_)) in (d_1839_v34_) else not(p1))
                d_1841_v36_: _dafny.Seq
                d_1841_v36_ = _dafny.SeqWithoutIsStrInference([default__.fm51(globalState)])
                d_1842_v37_: _dafny.Set
                d_1842_v37_ = _dafny.Set({len(d_1766_v0_)})
                index261_ = default__.safeIndex(479, (d_1767_v1_).length(0))
                rhs308_ = (d_1841_v36_)[default__.safeIndex(p2, len(d_1841_v36_))]
                rhs309_ = d_1832_cf18_
                rhs310_ = (d_1842_v37_).issubset((d_1842_v37_) - (_dafny.Set({d_1833_cf17_})))
                lhs253_ = globalState
                lhs254_ = d_1767_v1_
                lhs255_ = default__.safeIndex(479, (d_1767_v1_).length(0))
                d_1770_v3_ = rhs308_
                lhs253_.f5 = rhs309_
                lhs254_[lhs255_] = rhs310_
                d_1843_v38_: C5
                nw282_ = C5()
                nw282_.ctor__(len(d_1835_cf15_), d_1832_cf18_)
                d_1843_v38_ = nw282_
                index262_ = default__.safeIndex(479, (d_1767_v1_).length(0))
                (d_1767_v1_)[index262_] = p0
                d_1844_v39_: _dafny.Map
                d_1844_v39_ = _dafny.Map({len((d_1769_v2_ if p0 else _dafny.SeqWithoutIsStrInference([d_1832_cf18_]))): default__.fm0((d_1843_v38_).f15, _dafny.Set({d_1833_cf17_}), (d_1767_v1_)[default__.safeIndex(479, (d_1767_v1_).length(0))], globalState)})
                d_1844_v39_ = (d_1844_v39_).set(((d_1843_v38_).f15 if p1 else (d_1843_v38_).f15), d_1805_v16_)
                d_1845_v40_: C5
                nw283_ = C5()
                nw283_.ctor__(d_1832_cf18_, (0) - ((self).fm11(globalState)))
                d_1845_v40_ = nw283_
            elif True:
                index263_ = default__.safeIndex(87, (d_1767_v1_).length(0))
                (d_1767_v1_)[index263_] = p1
                index264_ = default__.safeIndex(701, (d_1767_v1_).length(0))
                (d_1767_v1_)[index264_] = p0
                index265_ = default__.safeIndex(87, (d_1767_v1_).length(0))
                index266_ = default__.safeIndex(701, (d_1767_v1_).length(0))
                rhs311_ = (p0) and (p1)
                rhs312_ = p1
                rhs313_ = not(p1)
                rhs314_ = p1
                lhs256_ = globalState
                lhs257_ = d_1767_v1_
                lhs258_ = default__.safeIndex(87, (d_1767_v1_).length(0))
                lhs259_ = globalState
                lhs260_ = d_1767_v1_
                lhs261_ = default__.safeIndex(701, (d_1767_v1_).length(0))
                lhs256_.f0 = rhs311_
                lhs257_[lhs258_] = rhs312_
                lhs259_.f0 = rhs313_
                lhs260_[lhs261_] = rhs314_
                d_1846_v41_: _dafny.MultiSet
                d_1846_v41_ = _dafny.MultiSet([p3, (d_1767_v1_)[default__.safeIndex(87, (d_1767_v1_).length(0))], p3])
                d_1847_v42_: C5
                nw284_ = C5()
                nw284_.ctor__((self).fm11(globalState), (0) - ((0) - (((d_1846_v41_).cardinality) + (-241))))
                d_1847_v42_ = nw284_
                d_1848_v43_: _dafny.Array
                nw285_ = _dafny.Array(None, 28)
                nw285_[int(0)] = p0
                nw285_[int(1)] = not(not(p0))
                nw285_[int(2)] = p1
                nw285_[int(3)] = (d_1767_v1_)[default__.safeIndex(87, (d_1767_v1_).length(0))]
                nw285_[int(4)] = default__.fm3(globalState)
                nw285_[int(5)] = (d_1767_v1_)[default__.safeIndex(87, (d_1767_v1_).length(0))]
                nw285_[int(6)] = (-985) <= (d_1833_cf17_)
                nw285_[int(7)] = p3
                nw285_[int(8)] = not(True)
                nw285_[int(9)] = p0
                nw285_[int(10)] = p0
                nw285_[int(11)] = p3
                nw285_[int(12)] = True
                nw285_[int(13)] = p3
                nw285_[int(14)] = p0
                nw285_[int(15)] = not(p0)
                nw285_[int(16)] = True
                nw285_[int(17)] = not (p3) or (p3)
                nw285_[int(18)] = (d_1767_v1_)[default__.safeIndex(87, (d_1767_v1_).length(0))]
                nw285_[int(19)] = (d_1767_v1_)[default__.safeIndex(87, (d_1767_v1_).length(0))]
                nw285_[int(20)] = (d_1767_v1_)[default__.safeIndex(87, (d_1767_v1_).length(0))]
                nw285_[int(21)] = p3
                nw285_[int(22)] = p3
                nw285_[int(23)] = not (p0) or (p1)
                nw285_[int(24)] = (d_1767_v1_)[default__.safeIndex(87, (d_1767_v1_).length(0))]
                nw285_[int(25)] = True
                nw285_[int(26)] = (d_1835_cf15_) != (d_1835_cf15_)
                nw285_[int(27)] = default__.fm3(globalState)
                d_1848_v43_ = nw285_
                d_1849_v44_: _dafny.Array
                def lambda75_(d_1850_p2_):
                    def lambda76_(d_1851_i4_):
                        return (d_1851_i4_) * (d_1850_p2_)

                    return lambda76_

                init48_ = lambda75_(p2)
                nw286_ = _dafny.Array(None, 27)
                for i0_48_ in range(nw286_.length(0)):
                    nw286_[i0_48_] = init48_(i0_48_)
                d_1849_v44_ = nw286_
                index267_ = default__.safeIndex(87, (d_1767_v1_).length(0))
                rhs315_ = not(p3)
                rhs316_ = (d_1767_v1_)[default__.safeIndex(87, (d_1767_v1_).length(0))]
                rhs317_ = d_1848_v43_
                rhs318_ = d_1849_v44_
                lhs262_ = d_1767_v1_
                lhs263_ = default__.safeIndex(87, (d_1767_v1_).length(0))
                lhs264_ = globalState
                lhs262_[lhs263_] = rhs315_
                lhs264_.f0 = rhs316_
                d_1848_v43_ = rhs317_
                d_1849_v44_ = rhs318_
                d_1852_v45_: _dafny.Array
                def lambda77_(d_1853_v2_):
                    def lambda78_(d_1854_i5_):
                        return d_1853_v2_

                    return lambda78_

                init49_ = lambda77_(d_1769_v2_)
                nw287_ = _dafny.Array(None, 23)
                for i0_49_ in range(nw287_.length(0)):
                    nw287_[i0_49_] = init49_(i0_49_)
                d_1852_v45_ = nw287_
                index268_ = default__.safeIndex(302, (d_1852_v45_).length(0))
                (d_1852_v45_)[index268_] = d_1769_v2_
                d_1855_v46_: _dafny.Map
                d_1855_v46_ = _dafny.Map({len(d_1769_v2_): d_1833_cf17_})
                d_1856_v47_: _dafny.Map
                d_1856_v47_ = _dafny.Map({d_1855_v46_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dgimpf"))})
                d_1857_v48_: D14
                d_1857_v48_ = D14_DC37(default__.fm3(globalState), ((d_1856_v47_)[d_1855_v46_] if (d_1855_v46_) in (d_1856_v47_) else d_1806_v17_))
                d_1858_v49_: D5
                d_1858_v49_ = D5_DC17(d_1769_v2_, p3)
                pat_let_tv39_ = p0
                pat_let_tv40_ = d_1767_v1_
                pat_let_tv41_ = d_1767_v1_
                pat_let_tv42_ = d_1806_v17_
                d_1859_v50_: _dafny.Array
                nw288_ = _dafny.Array(None, 29)
                nw288_[int(0)] = d_1857_v48_
                nw288_[int(1)] = D14_DC37(((d_1766_v0_)[p1] if (p1) in (d_1766_v0_) else (d_1767_v1_)[default__.safeIndex(87, (d_1767_v1_).length(0))]), d_1806_v17_)
                nw288_[int(2)] = D14_DC37(default__.fm3(globalState), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "im")))
                nw288_[int(3)] = d_1857_v48_
                nw288_[int(4)] = d_1857_v48_
                nw288_[int(5)] = d_1857_v48_
                nw288_[int(6)] = d_1857_v48_
                nw288_[int(7)] = d_1857_v48_
                nw288_[int(8)] = d_1857_v48_
                nw288_[int(9)] = d_1857_v48_
                nw288_[int(10)] = d_1857_v48_
                nw288_[int(11)] = d_1857_v48_
                nw288_[int(12)] = (D14_DC37(p3, d_1806_v17_) if True else d_1857_v48_)
                def iife150_(_pat_let49_0):
                    def iife151_(d_1860_dt__update__tmp_h1_):
                        def iife152_(_pat_let50_0):
                            def iife153_(d_1861_dt__update_hcf70_h0_):
                                return D14_DC37(d_1861_dt__update_hcf70_h0_, (d_1860_dt__update__tmp_h1_).cf71)
                            return iife153_(_pat_let50_0)
                        return iife152_(pat_let_tv39_)
                    return iife151_(_pat_let49_0)
                nw288_[int(13)] = iife150_(d_1857_v48_)
                nw288_[int(14)] = D14_DC37((d_1767_v1_)[default__.safeIndex(87, (d_1767_v1_).length(0))], d_1806_v17_)
                nw288_[int(15)] = d_1857_v48_
                nw288_[int(16)] = D14_DC37(p3, d_1806_v17_)
                nw288_[int(17)] = d_1857_v48_
                nw288_[int(18)] = d_1857_v48_
                nw288_[int(19)] = d_1857_v48_
                nw288_[int(20)] = d_1857_v48_
                nw288_[int(21)] = d_1857_v48_
                def iife154_(_pat_let51_0):
                    def iife155_(d_1862_dt__update__tmp_h2_):
                        def iife156_(_pat_let52_0):
                            def iife157_(d_1863_dt__update_hcf70_h1_):
                                return D14_DC37(d_1863_dt__update_hcf70_h1_, (d_1862_dt__update__tmp_h2_).cf71)
                            return iife157_(_pat_let52_0)
                        return iife156_((pat_let_tv41_)[default__.safeIndex(87, (pat_let_tv40_).length(0))])
                    return iife155_(_pat_let51_0)
                nw288_[int(22)] = iife154_(d_1857_v48_)
                nw288_[int(23)] = d_1857_v48_
                nw288_[int(24)] = d_1857_v48_
                nw288_[int(25)] = d_1857_v48_
                nw288_[int(26)] = (d_1857_v48_ if (d_1858_v49_).cf37 else d_1857_v48_)
                nw288_[int(27)] = d_1857_v48_
                def iife158_(_pat_let53_0):
                    def iife159_(d_1864_dt__update__tmp_h3_):
                        def iife160_(_pat_let54_0):
                            def iife161_(d_1865_dt__update_hcf71_h0_):
                                return D14_DC37((d_1864_dt__update__tmp_h3_).cf70, d_1865_dt__update_hcf71_h0_)
                            return iife161_(_pat_let54_0)
                        return iife160_(pat_let_tv42_)
                    return iife159_(_pat_let53_0)
                nw288_[int(28)] = iife158_(default__.fm52(p1, globalState))
                d_1859_v50_ = nw288_
                index269_ = default__.safeIndex(388, (d_1859_v50_).length(0))
                (d_1859_v50_)[index269_] = D14_DC37(p0, d_1806_v17_)
                index270_ = default__.safeIndex(302, (d_1852_v45_).length(0))
                index271_ = default__.safeIndex(388, (d_1859_v50_).length(0))
                rhs319_ = d_1805_v16_
                rhs320_ = (d_1769_v2_).set(default__.safeIndex(default__.safeModuloInt((self).fm11(globalState), d_1832_cf18_), len(d_1769_v2_)), d_1832_cf18_)
                rhs321_ = d_1834_cf16_
                rhs322_ = p0
                rhs323_ = d_1857_v48_
                lhs265_ = d_1852_v45_
                lhs266_ = default__.safeIndex(302, (d_1852_v45_).length(0))
                lhs267_ = globalState
                lhs268_ = d_1859_v50_
                lhs269_ = default__.safeIndex(388, (d_1859_v50_).length(0))
                d_1805_v16_ = rhs319_
                lhs265_[lhs266_] = rhs320_
                d_1834_cf16_ = rhs321_
                lhs267_.f0 = rhs322_
                lhs268_[lhs269_] = rhs323_
                d_1866_v51_: _dafny.Map
                d_1866_v51_ = _dafny.Map({d_1833_cf17_: p1})
                d_1866_v51_ = (d_1866_v51_).set(p2, p1)
            source23_ = d_1786_v15_
            if source23_.is_DC6:
                d_1867___mcc_h18_ = source23_.cf13
                d_1868___mcc_h19_ = source23_.cf14
                d_1869_cf14_ = d_1868___mcc_h19_
                d_1870_cf13_ = d_1867___mcc_h18_
                d_1871_v52_: _dafny.MultiSet
                d_1871_v52_ = _dafny.MultiSet([_dafny.CodePoint('h'), d_1805_v16_])
                d_1872_v53_: D1
                d_1872_v53_ = D1_DC3(d_1871_v52_)
                d_1873_v54_: _dafny.MultiSet
                d_1873_v54_ = _dafny.MultiSet([d_1872_v53_, d_1872_v53_, d_1872_v53_, d_1872_v53_, d_1872_v53_])
                rhs324_ = _dafny.SeqWithoutIsStrInference([d_1805_v16_ for d_1874_i6_ in range(default__.abs(61))])
                rhs325_ = (default__.fm53(d_1805_v16_, globalState)).isdisjoint(d_1873_v54_)
                lhs270_ = globalState
                lhs271_ = globalState
                lhs270_.f2 = rhs324_
                lhs271_.f0 = rhs325_
                (globalState).f5 = d_1833_cf17_
                (globalState).f0 = p0
                d_1875_v55_: _dafny.Array
                nw289_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 12)
                d_1875_v55_ = nw289_
                d_1876_v56_: D10
                d_1876_v56_ = D10_DC29(534, d_1875_v55_, True, p2)
                d_1877_v57_: _dafny.Map
                d_1877_v57_ = _dafny.Map({p0: 525})
                pat_let_tv43_ = d_1875_v55_
                pat_let_tv44_ = d_1870_cf13_
                pat_let_tv45_ = p3
                d_1878_v58_: _dafny.Array
                nw290_ = _dafny.Array(None, 23)
                nw290_[int(0)] = d_1876_v56_
                nw290_[int(1)] = d_1876_v56_
                nw290_[int(2)] = d_1876_v56_
                nw290_[int(3)] = d_1876_v56_
                nw290_[int(4)] = d_1876_v56_
                def iife162_(_pat_let55_0):
                    def iife163_(d_1879_dt__update__tmp_h4_):
                        def iife164_(_pat_let56_0):
                            def iife165_(d_1880_dt__update_hcf54_h0_):
                                def iife166_(_pat_let57_0):
                                    def iife167_(d_1881_dt__update_hcf56_h0_):
                                        def iife168_(_pat_let58_0):
                                            def iife169_(d_1882_dt__update_hcf55_h0_):
                                                return D10_DC29((d_1879_dt__update__tmp_h4_).cf53, d_1880_dt__update_hcf54_h0_, d_1882_dt__update_hcf55_h0_, d_1881_dt__update_hcf56_h0_)
                                            return iife169_(_pat_let58_0)
                                        return iife168_(pat_let_tv45_)
                                    return iife167_(_pat_let57_0)
                                return iife166_(pat_let_tv44_)
                            return iife165_(_pat_let56_0)
                        return iife164_(pat_let_tv43_)
                    return iife163_(_pat_let55_0)
                nw290_[int(5)] = (iife162_(D10_DC29(len(d_1806_v17_), d_1875_v55_, p1, -766)) if p3 else d_1876_v56_)
                nw290_[int(6)] = d_1876_v56_
                nw290_[int(7)] = d_1876_v56_
                nw290_[int(8)] = d_1876_v56_
                nw290_[int(9)] = d_1876_v56_
                nw290_[int(10)] = d_1876_v56_
                nw290_[int(11)] = D10_DC29(d_1832_cf18_, d_1875_v55_, p0, d_1833_cf17_)
                nw290_[int(12)] = d_1876_v56_
                nw290_[int(13)] = D10_DC29(d_1833_cf17_, d_1875_v55_, p3, d_1870_cf13_)
                nw290_[int(14)] = d_1876_v56_
                nw290_[int(15)] = d_1876_v56_
                nw290_[int(16)] = d_1876_v56_
                nw290_[int(17)] = d_1876_v56_
                nw290_[int(18)] = d_1876_v56_
                nw290_[int(19)] = d_1876_v56_
                nw290_[int(20)] = D10_DC29(len((d_1877_v57_).set(p0, d_1832_cf18_)), d_1875_v55_, False, p2)
                nw290_[int(21)] = d_1876_v56_
                nw290_[int(22)] = d_1876_v56_
                d_1878_v58_ = nw290_
                d_1878_v58_ = d_1878_v58_
            elif source23_.is_DC7:
                d_1883___mcc_h20_ = source23_.cf15
                d_1884___mcc_h21_ = source23_.cf16
                d_1885___mcc_h22_ = source23_.cf17
                d_1886___mcc_h23_ = source23_.cf18
                d_1887_cf18_ = d_1886___mcc_h23_
                d_1888_cf17_ = d_1885___mcc_h22_
                d_1889_cf16_ = d_1884___mcc_h21_
                d_1890_cf15_ = d_1883___mcc_h20_
                d_1891_v59_: _dafny.MultiSet
                d_1891_v59_ = _dafny.MultiSet([len(d_1769_v2_), p2, p2, d_1888_cf17_, 534])
                d_1892_v60_: _dafny.Seq
                d_1892_v60_ = _dafny.SeqWithoutIsStrInference([d_1769_v2_, d_1769_v2_])
                d_1893_v61_: _dafny.Set
                d_1893_v61_ = _dafny.Set({len((d_1892_v60_)[default__.safeIndex(d_1888_cf17_, len(d_1892_v60_))])})
                d_1805_v16_ = default__.fm0((d_1891_v59_).cardinality, d_1893_v61_, p0, globalState)
                d_1894_v62_: _dafny.Map
                d_1894_v62_ = _dafny.Map({d_1889_cf16_: d_1888_cf17_})
                d_1894_v62_ = (d_1894_v62_).set(d_1889_cf16_, d_1887_cf18_)
                rhs326_ = not (p3) or ((p1) in (_dafny.Map({p3: d_1887_cf18_})))
                rhs327_ = d_1887_cf18_
                lhs272_ = globalState
                lhs272_.f0 = rhs326_
                r2 = rhs327_
                d_1895_v63_: _dafny.Map
                d_1895_v63_ = _dafny.Map({d_1888_cf17_: default__.fm0(d_1888_cf17_, d_1893_v61_, p0, globalState)})
                d_1896_v64_: _dafny.MultiSet
                d_1896_v64_ = _dafny.MultiSet([d_1895_v63_, d_1895_v63_])
                index272_ = default__.safeIndex(898, (d_1767_v1_).length(0))
                (d_1767_v1_)[index272_] = (d_1896_v64_).isdisjoint(d_1896_v64_)
                index273_ = default__.safeIndex(898, (d_1767_v1_).length(0))
                (d_1767_v1_)[index273_] = p0
            elif source23_.is_DC8:
                d_1897___mcc_h24_ = source23_.cf19
                d_1898___mcc_h25_ = source23_.cf20
                d_1899_cf20_ = d_1898___mcc_h25_
                d_1900_cf19_ = d_1897___mcc_h24_
                d_1901_v65_: C1
                nw291_ = C1()
                nw291_.ctor__(d_1833_cf17_)
                d_1901_v65_ = nw291_
                d_1902_v66_: _dafny.Map
                d_1902_v66_ = _dafny.Map({p2: p0})
                d_1903_v67_: _dafny.Set
                d_1903_v67_ = _dafny.Set({d_1899_cf20_})
                d_1904_v68_: _dafny.Array
                nw292_ = _dafny.Array(None, 9)
                nw292_[int(0)] = p2
                nw292_[int(1)] = len(d_1902_v66_)
                nw292_[int(2)] = d_1833_cf17_
                nw292_[int(3)] = p2
                nw292_[int(4)] = len(d_1903_v67_)
                nw292_[int(5)] = d_1832_cf18_
                nw292_[int(6)] = p2
                nw292_[int(7)] = d_1832_cf18_
                nw292_[int(8)] = len(d_1806_v17_)
                d_1904_v68_ = nw292_
                d_1905_v69_: _dafny.Map
                d_1905_v69_ = _dafny.Map({d_1904_v68_: d_1899_cf20_})
                d_1906_v70_: _dafny.Set
                d_1906_v70_ = _dafny.Set({len(d_1905_v69_)})
                d_1907_v71_: _dafny.Map
                d_1907_v71_ = _dafny.Map({len(d_1906_v70_): p1})
                d_1908_v72_: _dafny.Map
                d_1908_v72_ = _dafny.Map({d_1833_cf17_: default__.fm4(len(d_1902_v66_), p2, d_1833_cf17_, globalState)})
                d_1909_v73_: _dafny.Array
                nw293_ = _dafny.Array(None, 11)
                nw293_[int(0)] = p2
                nw293_[int(1)] = d_1832_cf18_
                nw293_[int(2)] = d_1833_cf17_
                nw293_[int(3)] = d_1832_cf18_
                nw293_[int(4)] = d_1833_cf17_
                nw293_[int(5)] = (len(d_1908_v72_)) + (p2)
                nw293_[int(6)] = d_1833_cf17_
                nw293_[int(7)] = d_1833_cf17_
                nw293_[int(8)] = 902
                nw293_[int(9)] = d_1833_cf17_
                nw293_[int(10)] = 905
                d_1909_v73_ = nw293_
                index274_ = default__.safeIndex(201, (d_1909_v73_).length(0))
                (d_1909_v73_)[index274_] = d_1832_cf18_
                d_1910_v74_: _dafny.MultiSet
                d_1910_v74_ = _dafny.MultiSet([p0])
                d_1911_v75_: _dafny.MultiSet
                d_1911_v75_ = d_1910_v74_
                index275_ = default__.safeIndex(201, (d_1909_v73_).length(0))
                rhs328_ = (0) - ((d_1832_cf18_) - (d_1833_cf17_))
                rhs329_ = p2
                rhs330_ = p2
                rhs331_ = ((d_1910_v74_).intersection((d_1911_v75_))).intersection((d_1910_v74_) - (d_1910_v74_))
                lhs273_ = d_1909_v73_
                lhs274_ = default__.safeIndex(201, (d_1909_v73_).length(0))
                lhs275_ = globalState
                d_1832_cf18_ = rhs328_
                lhs273_[lhs274_] = rhs329_
                lhs275_.f5 = rhs330_
                d_1910_v74_ = rhs331_
                d_1912_v76_: _dafny.Set
                d_1912_v76_ = _dafny.Set({d_1906_v70_, d_1906_v70_})
                d_1913_v77_: _dafny.Map
                d_1913_v77_ = _dafny.Map({(d_1833_cf17_ if d_1900_cf19_ else d_1832_cf18_): (0) - (default__.safeDivisionInt(len(d_1912_v76_), (0) - ((d_1909_v73_)[default__.safeIndex(201, (d_1909_v73_).length(0))])))})
                d_1914_v78_: _dafny.Map
                d_1914_v78_ = _dafny.Map({d_1767_v1_: p1})
                d_1915_v79_: D14
                d_1915_v79_ = D14_DC37(d_1899_cf20_, d_1806_v17_)
                d_1916_v80_: D0
                d_1916_v80_ = D0_DC1(True, d_1832_cf18_, default__.fm0((d_1909_v73_)[default__.safeIndex(201, (d_1909_v73_).length(0))], d_1906_v70_, d_1900_cf19_, globalState))
                pat_let_tv46_ = p3
                pat_let_tv47_ = d_1805_v16_
                def iife170_(_pat_let59_0):
                    def iife171_(d_1917_dt__update__tmp_h5_):
                        def iife172_(_pat_let60_0):
                            def iife173_(d_1918_dt__update_hcf2_h1_):
                                def iife174_(_pat_let61_0):
                                    def iife175_(d_1919_dt__update_hcf4_h0_):
                                        return D0_DC1(d_1918_dt__update_hcf2_h1_, (d_1917_dt__update__tmp_h5_).cf3, d_1919_dt__update_hcf4_h0_)
                                    return iife175_(_pat_let61_0)
                                return iife174_(pat_let_tv47_)
                            return iife173_(_pat_let60_0)
                        return iife172_(pat_let_tv46_)
                    return iife171_(_pat_let59_0)
                rhs332_ = (d_1913_v77_) | (_dafny.Map({len(d_1835_cf15_): (iife170_(d_1916_v80_)).cf3}))
                rhs333_ = (d_1914_v78_ if d_1899_cf20_ else d_1914_v78_)
                rhs334_ = d_1915_v79_
                d_1913_v77_ = rhs332_
                d_1914_v78_ = rhs333_
                d_1915_v79_ = rhs334_
                d_1920_v82_: C3
                nw294_ = C3()
                def iife176_():
                    coll52_ = _dafny.Map()
                    compr_52_: int
                    for compr_52_ in (d_1902_v66_).keys.Elements:
                        d_1921_v81_: int = compr_52_
                        if (d_1921_v81_) in (d_1902_v66_):
                            coll52_[(d_1921_v81_) * (d_1832_cf18_)] = _dafny.SeqWithoutIsStrInference([d_1805_v16_ for d_1922_i7_ in range(default__.abs(517))])
                    return _dafny.Map(coll52_)
                nw294_.ctor__(d_1913_v77_, d_1767_v1_, iife176_()
                , d_1832_cf18_)
                d_1920_v82_ = nw294_
                d_1923_v83_: _dafny.Seq
                d_1923_v83_ = _dafny.SeqWithoutIsStrInference([d_1920_v82_])
                d_1924_v84_: D17
                d_1924_v84_ = D17_DC40(d_1920_v82_)
                d_1925_v85_: _dafny.Array
                nw295_ = _dafny.Array(None, 20)
                nw295_[int(0)] = d_1920_v82_
                nw295_[int(1)] = d_1920_v82_
                nw295_[int(2)] = d_1920_v82_
                nw295_[int(3)] = (d_1920_v82_ if p0 else d_1920_v82_)
                nw295_[int(4)] = d_1920_v82_
                nw295_[int(5)] = d_1920_v82_
                nw295_[int(6)] = d_1920_v82_
                nw295_[int(7)] = d_1920_v82_
                nw295_[int(8)] = (d_1923_v83_)[default__.safeIndex(d_1832_cf18_, len(d_1923_v83_))]
                nw295_[int(9)] = d_1920_v82_
                nw295_[int(10)] = d_1920_v82_
                nw295_[int(11)] = d_1920_v82_
                nw295_[int(12)] = d_1920_v82_
                nw295_[int(13)] = d_1920_v82_
                nw295_[int(14)] = d_1920_v82_
                nw295_[int(15)] = d_1920_v82_
                nw295_[int(16)] = d_1920_v82_
                nw295_[int(17)] = (d_1924_v84_).cf74
                nw295_[int(18)] = d_1920_v82_
                nw295_[int(19)] = d_1920_v82_
                d_1925_v85_ = nw295_
                index276_ = default__.safeIndex(155, (d_1925_v85_).length(0))
                (d_1925_v85_)[index276_] = d_1920_v82_
                d_1926_v86_: _dafny.Map
                d_1926_v86_ = _dafny.Map({d_1833_cf17_: d_1920_v82_})
                index277_ = default__.safeIndex(155, (d_1925_v85_).length(0))
                (d_1925_v85_)[index277_] = ((d_1926_v86_)[(0) - ((0) - (((d_1909_v73_)[default__.safeIndex(201, (d_1909_v73_).length(0))]) + ((d_1909_v73_)[default__.safeIndex(201, (d_1909_v73_).length(0))])))] if ((0) - ((0) - (((d_1909_v73_)[default__.safeIndex(201, (d_1909_v73_).length(0))]) + ((d_1909_v73_)[default__.safeIndex(201, (d_1909_v73_).length(0))])))) in (d_1926_v86_) else d_1920_v82_)
            elif True:
                d_1927___mcc_h26_ = source23_.cf12
                d_1928_cf12_ = d_1927___mcc_h26_
                d_1929_v87_: _dafny.Array
                nw296_ = _dafny.Array(int(0), 24)
                d_1929_v87_ = nw296_
                d_1930_v88_: _dafny.MultiSet
                d_1930_v88_ = _dafny.MultiSet([d_1929_v87_])
                d_1931_v89_: _dafny.Map
                d_1931_v89_ = _dafny.Map({p0: _dafny.SeqWithoutIsStrInference([d_1832_cf18_ for d_1932_i8_ in range(default__.abs(403))])})
                d_1933_v90_: _dafny.MultiSet
                d_1933_v90_ = _dafny.MultiSet([p1])
                d_1934_v91_: _dafny.Map
                d_1934_v91_ = _dafny.Map({284: ((d_1806_v17_).set(default__.safeIndex(850, len(d_1806_v17_)), d_1805_v16_)) + (_dafny.SeqWithoutIsStrInference([d_1805_v16_ for d_1935_i9_ in range(default__.abs(396))]))})
                rhs335_ = (d_1930_v88_).isdisjoint(d_1930_v88_)
                rhs336_ = (d_1833_cf17_) != (d_1833_cf17_)
                rhs337_ = (d_1832_cf18_) <= (default__.safeDivisionInt(len(((d_1931_v89_)[p0] if (p0) in (d_1931_v89_) else d_1769_v2_)), (d_1933_v90_).cardinality))
                rhs338_ = ((d_1934_v91_)[d_1832_cf18_] if (d_1832_cf18_) in (d_1934_v91_) else d_1806_v17_)
                lhs276_ = globalState
                lhs277_ = globalState
                lhs278_ = globalState
                lhs276_.f0 = rhs335_
                lhs277_.f0 = rhs336_
                lhs278_.f0 = rhs337_
                d_1806_v17_ = rhs338_
                d_1767_v1_ = d_1767_v1_
                d_1805_v16_ = d_1805_v16_
                index278_ = default__.safeIndex(912, (d_1929_v87_).length(0))
                (d_1929_v87_)[index278_] = d_1833_cf17_
                d_1936_v92_: _dafny.Seq
                d_1936_v92_ = _dafny.SeqWithoutIsStrInference([d_1929_v87_])
                d_1937_v93_: D4
                d_1937_v93_ = D4_DC15(p1, p0, (d_1806_v17_).set(default__.safeIndex(d_1833_cf17_, len(d_1806_v17_)), d_1805_v16_))
                d_1938_v94_: _dafny.Map
                d_1938_v94_ = _dafny.Map({d_1937_v93_: d_1929_v87_})
                d_1939_v95_: _dafny.Map
                d_1939_v95_ = _dafny.Map({p2: ((d_1938_v94_)[d_1937_v93_] if (d_1937_v93_) in (d_1938_v94_) else d_1929_v87_)})
                d_1940_v96_: D2
                d_1940_v96_ = D2_DC6(p2, d_1929_v87_)
                d_1941_v97_: _dafny.Array
                nw297_ = _dafny.Array(None, 28)
                nw297_[int(0)] = d_1929_v87_
                nw297_[int(1)] = (d_1936_v92_)[default__.safeIndex(len(d_1939_v95_), len(d_1936_v92_))]
                nw297_[int(2)] = d_1929_v87_
                nw297_[int(3)] = d_1929_v87_
                nw297_[int(4)] = d_1929_v87_
                nw297_[int(5)] = d_1929_v87_
                nw297_[int(6)] = d_1929_v87_
                nw297_[int(7)] = d_1929_v87_
                nw297_[int(8)] = (d_1929_v87_ if p1 else d_1929_v87_)
                nw297_[int(9)] = d_1929_v87_
                nw297_[int(10)] = (d_1929_v87_ if p0 else (d_1936_v92_)[default__.safeIndex(d_1833_cf17_, len(d_1936_v92_))])
                nw297_[int(11)] = d_1929_v87_
                nw297_[int(12)] = d_1929_v87_
                nw297_[int(13)] = d_1929_v87_
                nw297_[int(14)] = d_1929_v87_
                nw297_[int(15)] = (d_1940_v96_).cf14
                nw297_[int(16)] = d_1929_v87_
                nw297_[int(17)] = d_1929_v87_
                nw297_[int(18)] = d_1929_v87_
                nw297_[int(19)] = d_1929_v87_
                nw297_[int(20)] = (d_1929_v87_ if p3 else d_1929_v87_)
                nw297_[int(21)] = d_1929_v87_
                nw297_[int(22)] = d_1929_v87_
                nw297_[int(23)] = d_1929_v87_
                nw297_[int(24)] = d_1929_v87_
                nw297_[int(25)] = d_1929_v87_
                nw297_[int(26)] = d_1929_v87_
                nw297_[int(27)] = d_1929_v87_
                d_1941_v97_ = nw297_
                index279_ = default__.safeIndex(26, (d_1941_v97_).length(0))
                (d_1941_v97_)[index279_] = d_1929_v87_
                d_1942_v98_: _dafny.Set
                d_1942_v98_ = _dafny.Set({(d_1806_v17_)[default__.safeIndex(d_1832_cf18_, len(d_1806_v17_))], d_1805_v16_, d_1805_v16_})
                d_1943_v99_: _dafny.Map
                d_1943_v99_ = _dafny.Map({default__.fm2(d_1832_cf18_, globalState): d_1769_v2_})
                index280_ = default__.safeIndex(912, (d_1929_v87_).length(0))
                index281_ = default__.safeIndex(26, (d_1941_v97_).length(0))
                rhs339_ = d_1833_cf17_
                rhs340_ = d_1929_v87_
                rhs341_ = (d_1942_v98_).intersection(_dafny.Set({d_1805_v16_, _dafny.CodePoint('y')}))
                rhs342_ = (0) - (len(((d_1943_v99_)[len(d_1835_cf15_)] if (len(d_1835_cf15_)) in (d_1943_v99_) else _dafny.SeqWithoutIsStrInference([p2 for d_1944_i10_ in range(default__.abs(836))]))))
                lhs279_ = d_1929_v87_
                lhs280_ = default__.safeIndex(912, (d_1929_v87_).length(0))
                lhs281_ = d_1941_v97_
                lhs282_ = default__.safeIndex(26, (d_1941_v97_).length(0))
                lhs283_ = globalState
                lhs279_[lhs280_] = rhs339_
                lhs281_[lhs282_] = rhs340_
                d_1942_v98_ = rhs341_
                lhs283_.f5 = rhs342_
            d_1945_v100_: _dafny.Array
            nw298_ = _dafny.Array(_dafny.Seq({}), 25)
            d_1945_v100_ = nw298_
            index282_ = default__.safeIndex(240, (d_1945_v100_).length(0))
            (d_1945_v100_)[index282_] = d_1769_v2_
            index283_ = default__.safeIndex(240, (d_1945_v100_).length(0))
            (d_1945_v100_)[index283_] = d_1769_v2_
            (globalState).f5 = (d_1833_cf17_) + (d_1833_cf17_)
        elif source22_.is_DC8:
            d_1946___mcc_h15_ = source22_.cf19
            d_1947___mcc_h16_ = source22_.cf20
            d_1948_cf20_ = d_1947___mcc_h16_
            d_1949_cf19_ = d_1946___mcc_h15_
            d_1950_v101_: _dafny.Array
            nw299_ = _dafny.Array(int(0), 19)
            d_1950_v101_ = nw299_
            d_1951_v102_: _dafny.Map
            d_1951_v102_ = _dafny.Map({d_1950_v101_: d_1948_cf20_})
            d_1952_v103_: D10
            d_1952_v103_ = D10_DC28(d_1951_v102_)
            source24_ = d_1952_v103_
            if source24_.is_DC29:
                d_1953___mcc_h27_ = source24_.cf53
                d_1954___mcc_h28_ = source24_.cf54
                d_1955___mcc_h29_ = source24_.cf55
                d_1956___mcc_h30_ = source24_.cf56
                d_1957_cf56_ = d_1956___mcc_h30_
                d_1958_cf55_ = d_1955___mcc_h29_
                d_1959_cf54_ = d_1954___mcc_h28_
                d_1960_cf53_ = d_1953___mcc_h27_
                d_1961_v104_: _dafny.Map
                d_1961_v104_ = _dafny.Map({d_1960_cf53_: _dafny.SeqWithoutIsStrInference([d_1805_v16_ for d_1962_i11_ in range(default__.abs(958))])})
                d_1963_v105_: C7
                nw300_ = C7()
                nw300_.ctor__(d_1950_v101_, d_1950_v101_, d_1960_cf53_, d_1767_v1_, d_1961_v104_)
                d_1963_v105_ = nw300_
                d_1964_v106_: _dafny.Set
                d_1964_v106_ = _dafny.Set({p3})
                d_1965_v107_: D11
                d_1965_v107_ = D11_DC33(d_1960_cf53_, d_1958_cf55_, 655, False, d_1964_v106_)
                d_1966_v108_: _dafny.Map
                d_1966_v108_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([(d_1965_v107_).cf65, d_1949_cf19_, d_1949_cf19_, p1, d_1949_cf19_]): d_1961_v104_})
                d_1967_v109_: _dafny.Seq
                d_1967_v109_ = _dafny.SeqWithoutIsStrInference([p0, d_1958_cf55_])
                nw301_ = C7()
                nw301_.ctor__(d_1963_v105_.f11, d_1950_v101_, d_1960_cf53_, d_1767_v1_, ((d_1966_v108_)[d_1967_v109_] if (d_1967_v109_) in (d_1966_v108_) else d_1961_v104_))
                d_1963_v105_ = nw301_
                d_1806_v17_ = ((d_1806_v17_) + (d_1806_v17_)) + (d_1806_v17_)
                index284_ = default__.safeIndex(923, (d_1959_cf54_).length(0))
                (d_1959_cf54_)[index284_] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "btcbtiuf"))) + (d_1806_v17_)
                index285_ = default__.safeIndex(923, (d_1959_cf54_).length(0))
                rhs343_ = p2
                rhs344_ = (d_1806_v17_) + (d_1806_v17_)
                lhs284_ = d_1959_cf54_
                lhs285_ = default__.safeIndex(923, (d_1959_cf54_).length(0))
                d_1960_cf53_ = rhs343_
                lhs284_[lhs285_] = rhs344_
                (globalState).f5 = (p2) + (d_1960_cf53_)
            elif True:
                d_1968___mcc_h31_ = source24_.cf52
                d_1969_cf52_ = d_1968___mcc_h31_
                d_1970_v110_: _dafny.MultiSet
                d_1970_v110_ = _dafny.MultiSet([default__.fm0(192, _dafny.Set({-474}), d_1949_cf19_, globalState), d_1805_v16_])
                d_1971_v111_: D1
                d_1971_v111_ = D1_DC3(d_1970_v110_)
                pat_let_tv48_ = d_1970_v110_
                d_1972_v112_: _dafny.Map
                def iife177_(_pat_let62_0):
                    def iife178_(d_1973_dt__update__tmp_h6_):
                        def iife179_(_pat_let63_0):
                            def iife180_(d_1974_dt__update_hcf6_h0_):
                                return D1_DC3(d_1974_dt__update_hcf6_h0_)
                            return iife180_(_pat_let63_0)
                        return iife179_(pat_let_tv48_)
                    return iife178_(_pat_let62_0)
                d_1972_v112_ = _dafny.Map({p2: iife177_(d_1971_v111_)})
                d_1972_v112_ = (d_1972_v112_).set(p2, d_1971_v111_)
                d_1975_v113_: _dafny.MultiSet
                d_1975_v113_ = _dafny.MultiSet([p2, len(d_1806_v17_)])
                d_1976_v114_: _dafny.Seq
                d_1976_v114_ = _dafny.SeqWithoutIsStrInference([d_1975_v113_, d_1975_v113_, d_1975_v113_])
                d_1977_v115_: D11
                d_1977_v115_ = D11_DC31(p2, p2)
                pat_let_tv49_ = p2
                pat_let_tv50_ = p2
                pat_let_tv51_ = d_1806_v17_
                d_1978_v116_: _dafny.Map
                def iife183_(_pat_let66_0):
                    def iife184_(d_1979_dt__update__tmp_h7_):
                        def iife185_(_pat_let67_0):
                            def iife186_(d_1980_dt__update_hcf59_h0_):
                                return D11_DC31((d_1979_dt__update__tmp_h7_).cf58, d_1980_dt__update_hcf59_h0_)
                            return iife186_(_pat_let67_0)
                        return iife185_(pat_let_tv49_)
                    return iife184_(_pat_let66_0)
                def iife182_(_pat_let65_0):
                    def iife187_(d_1981_dt__update__tmp_h8_):
                        def iife188_(_pat_let68_0):
                            def iife189_(d_1982_dt__update_hcf59_h1_):
                                return D11_DC31((d_1981_dt__update__tmp_h8_).cf58, d_1982_dt__update_hcf59_h1_)
                            return iife189_(_pat_let68_0)
                        return iife188_(pat_let_tv50_)
                    return iife187_(_pat_let65_0)
                def iife181_(_pat_let64_0):
                    def iife190_(d_1983_dt__update__tmp_h9_):
                        def iife191_(_pat_let69_0):
                            def iife192_(d_1984_dt__update_hcf58_h0_):
                                return D11_DC31(d_1984_dt__update_hcf58_h0_, (d_1983_dt__update__tmp_h9_).cf59)
                            return iife192_(_pat_let69_0)
                        return iife191_(len(pat_let_tv51_))
                    return iife190_(_pat_let64_0)
                d_1978_v116_ = _dafny.Map({len(d_1976_v114_): (iife181_(iife182_(iife183_(d_1977_v115_)))).cf59})
                d_1978_v116_ = ((_dafny.Map({p2: p2})) | (d_1978_v116_)) | (_dafny.Map({p2: p2}))
                d_1949_cf19_ = d_1948_cf20_
                (globalState).f5 = (self).fm11(globalState)
            d_1985_v117_: _dafny.Array
            nw302_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 2)
            d_1985_v117_ = nw302_
            index286_ = default__.safeIndex(517, (d_1985_v117_).length(0))
            (d_1985_v117_)[index286_] = d_1806_v17_
            d_1986_v118_: _dafny.Seq
            d_1986_v118_ = _dafny.SeqWithoutIsStrInference([d_1769_v2_])
            d_1987_v119_: _dafny.Seq
            d_1987_v119_ = _dafny.SeqWithoutIsStrInference([(d_1806_v17_) + (d_1806_v17_)])
            d_1988_v120_: _dafny.Set
            d_1988_v120_ = _dafny.Set({d_1950_v101_, d_1950_v101_})
            d_1989_v121_: _dafny.MultiSet
            d_1989_v121_ = _dafny.MultiSet([len(d_1988_v120_)])
            index287_ = default__.safeIndex(517, (d_1985_v117_).length(0))
            rhs345_ = (d_1986_v118_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_1806_v17_ for d_1990_i12_ in range(default__.abs(291))])), len(d_1986_v118_))]
            rhs346_ = (d_1987_v119_)[default__.safeIndex(((d_1989_v121_)[974] if (974) in (d_1989_v121_) else p2), len(d_1987_v119_))]
            lhs286_ = d_1985_v117_
            lhs287_ = default__.safeIndex(517, (d_1985_v117_).length(0))
            d_1769_v2_ = rhs345_
            lhs286_[lhs287_] = rhs346_
            (globalState).f5 = default__.safeModuloInt(p2, -255)
            d_1991_v122_: _dafny.MultiSet
            d_1991_v122_ = _dafny.MultiSet([d_1948_cf20_])
            d_1992_v123_: _dafny.Map
            d_1992_v123_ = _dafny.Map({d_1991_v122_: p1})
            d_1992_v123_ = (d_1992_v123_).set(_dafny.MultiSet([d_1948_cf20_, d_1948_cf20_, ((d_1766_v0_)[p3] if (p3) in (d_1766_v0_) else p0)]), p3)
        elif True:
            d_1993___mcc_h17_ = source22_.cf12
            d_1994_cf12_ = d_1993___mcc_h17_
            d_1995_v124_: _dafny.Array
            def lambda79_(d_1996_p2_):
                def lambda80_(d_1997_i13_):
                    return (d_1997_i13_) - ((0) - (len(_dafny.Set({d_1996_p2_, d_1996_p2_, d_1996_p2_}))))

                return lambda80_

            init50_ = lambda79_(p2)
            nw303_ = _dafny.Array(None, 18)
            for i0_50_ in range(nw303_.length(0)):
                nw303_[i0_50_] = init50_(i0_50_)
            d_1995_v124_ = nw303_
            index288_ = default__.safeIndex(332, (d_1995_v124_).length(0))
            (d_1995_v124_)[index288_] = p2
            index289_ = default__.safeIndex(332, (d_1995_v124_).length(0))
            (d_1995_v124_)[index289_] = p2
            d_1998_v125_: _dafny.Map
            d_1998_v125_ = _dafny.Map({p3: default__.fm2(p2, globalState)})
            d_1999_v126_: _dafny.Map
            d_1999_v126_ = _dafny.Map({d_1998_v125_: default__.safeModuloInt((d_1995_v124_)[default__.safeIndex(332, (d_1995_v124_).length(0))], p2)})
            index290_ = default__.safeIndex(332, (d_1995_v124_).length(0))
            (d_1995_v124_)[index290_] = ((d_1999_v126_)[d_1998_v125_] if (d_1998_v125_) in (d_1999_v126_) else default__.safeDivisionInt((d_1995_v124_)[default__.safeIndex(332, (d_1995_v124_).length(0))], (d_1995_v124_)[default__.safeIndex(332, (d_1995_v124_).length(0))]))
            d_2000_v127_: D4
            d_2000_v127_ = D4_DC13(_dafny.SeqWithoutIsStrInference([len(d_1806_v17_), (d_1995_v124_)[default__.safeIndex(332, (d_1995_v124_).length(0))], len(default__.fm17((0) - (p2), p2, globalState)), (0) - ((d_1995_v124_)[default__.safeIndex(332, (d_1995_v124_).length(0))]), p2]), False, len(d_1769_v2_))
            d_2000_v127_ = d_2000_v127_
            d_2001_v128_: _dafny.Array
            nw304_ = _dafny.Array(None, 2)
            nw304_[int(0)] = (d_1806_v17_).set(default__.safeIndex((0) - ((d_1995_v124_)[default__.safeIndex(332, (d_1995_v124_).length(0))]), len(d_1806_v17_)), d_1805_v16_)
            nw304_[int(1)] = d_1806_v17_
            d_2001_v128_ = nw304_
            d_2002_v129_: D10
            d_2002_v129_ = D10_DC29(p2, d_2001_v128_, p3, p2)
            d_2003_v130_: _dafny.Map
            d_2003_v130_ = _dafny.Map({not(p3): d_2002_v129_})
            index291_ = default__.safeIndex(299, (d_1767_v1_).length(0))
            (d_1767_v1_)[index291_] = (p2) > ((d_1995_v124_)[default__.safeIndex(332, (d_1995_v124_).length(0))])
            d_2004_v131_: _dafny.Map
            d_2004_v131_ = _dafny.Map({d_1805_v16_: (p1 if p1 else p0)})
            d_2005_v132_: _dafny.Seq
            d_2005_v132_ = _dafny.SeqWithoutIsStrInference([d_2003_v130_])
            d_2006_v133_: _dafny.Map
            d_2006_v133_ = _dafny.Map({(d_1995_v124_)[default__.safeIndex(332, (d_1995_v124_).length(0))]: d_2003_v130_})
            d_2007_v134_: _dafny.Map
            d_2007_v134_ = _dafny.Map({p2: True})
            index292_ = default__.safeIndex(299, (d_1767_v1_).length(0))
            rhs347_ = 326
            rhs348_ = d_1995_v124_
            rhs349_ = ((((d_2005_v132_)[default__.safeIndex((d_1995_v124_)[default__.safeIndex(332, (d_1995_v124_).length(0))], len(d_2005_v132_))]).set(p0, d_2002_v129_) if p3 else d_2003_v130_)) | (((d_2006_v133_)[p2] if (p2) in (d_2006_v133_) else _dafny.Map({p0: d_2002_v129_})))
            rhs350_ = (D2_DC8(((d_2007_v134_)[-284] if (-284) in (d_2007_v134_) else p1), p3)).cf19
            rhs351_ = _dafny.Map({d_1805_v16_: p0})
            lhs288_ = d_1767_v1_
            lhs289_ = default__.safeIndex(299, (d_1767_v1_).length(0))
            r2 = rhs347_
            d_1995_v124_ = rhs348_
            d_2003_v130_ = rhs349_
            lhs288_[lhs289_] = rhs350_
            d_2004_v131_ = rhs351_
        d_2008_v135_: _dafny.Set
        d_2008_v135_ = _dafny.Set({not(p0), p3})
        d_2009_v136_: D6
        d_2009_v136_ = D6_DC19(p2, (self).fm11(globalState), len((d_2008_v135_) | (d_2008_v135_)))
        d_2010_v137_: _dafny.Array
        nw305_ = _dafny.Array(None, 21)
        d_2010_v137_ = nw305_
        d_2011_v138_: _dafny.MultiSet
        d_2011_v138_ = _dafny.MultiSet([d_2010_v137_, d_2010_v137_, d_2010_v137_])
        rhs352_ = d_2009_v136_
        rhs353_ = ((d_2011_v138_) | (d_2011_v138_)).cardinality
        lhs290_ = globalState
        d_2009_v136_ = rhs352_
        lhs290_.f5 = rhs353_
        d_2012_v139_: _dafny.Array
        nw306_ = _dafny.Array(int(0), 13)
        d_2012_v139_ = nw306_
        d_2013_v140_: D2
        d_2013_v140_ = D2_DC6(p2, d_2012_v139_)
        d_2014_v141_: _dafny.Map
        d_2014_v141_ = _dafny.Map({d_2013_v140_: p0})
        d_2015_v142_: _dafny.Seq
        d_2015_v142_ = _dafny.SeqWithoutIsStrInference([d_2014_v141_])
        r0 = ((d_2014_v141_) | (d_2014_v141_)) | ((d_2014_v141_ if False else (d_2015_v142_)[default__.safeIndex(p2, len(d_2015_v142_))]))
        r1 = d_1766_v0_
        r2 = (d_1769_v2_)[default__.safeIndex(p2, len(d_1769_v2_))]
        return r0, r1, r2

    def m5(self, p0, p1, globalState):
        r0: bool = False
        r1: bool = False
        r2: int = int(0)
        d_2016_v0_: int
        d_2016_v0_ = 372
        d_2017_v1_: _dafny.Map
        d_2017_v1_ = _dafny.Map({d_2016_v0_: 100})
        d_2018_v2_: _dafny.Map
        d_2018_v2_ = _dafny.Map({d_2016_v0_: d_2017_v1_})
        r2 = len((_dafny.Map({d_2016_v0_: d_2017_v1_})) | (d_2018_v2_))
        d_2016_v0_ = d_2016_v0_
        d_2019_v3_: _dafny.MultiSet
        d_2019_v3_ = _dafny.MultiSet([_dafny.CodePoint('h')])
        d_2020_v4_: D1
        d_2020_v4_ = D1_DC3(d_2019_v3_)
        d_2021_v5_: str
        d_2021_v5_ = _dafny.CodePoint('o')
        d_2022_v6_: _dafny.Array
        nw307_ = _dafny.Array(None, 23)
        nw307_[int(0)] = d_2020_v4_
        nw307_[int(1)] = d_2020_v4_
        nw307_[int(2)] = d_2020_v4_
        nw307_[int(3)] = d_2020_v4_
        nw307_[int(4)] = d_2020_v4_
        nw307_[int(5)] = d_2020_v4_
        nw307_[int(6)] = d_2020_v4_
        nw307_[int(7)] = d_2020_v4_
        nw307_[int(8)] = D1_DC3(d_2019_v3_)
        nw307_[int(9)] = d_2020_v4_
        nw307_[int(10)] = D1_DC3(d_2019_v3_)
        nw307_[int(11)] = d_2020_v4_
        nw307_[int(12)] = (d_2020_v4_ if p0 else d_2020_v4_)
        nw307_[int(13)] = d_2020_v4_
        nw307_[int(14)] = d_2020_v4_
        nw307_[int(15)] = D1_DC3(_dafny.MultiSet([d_2021_v5_]))
        nw307_[int(16)] = d_2020_v4_
        nw307_[int(17)] = d_2020_v4_
        nw307_[int(18)] = d_2020_v4_
        nw307_[int(19)] = d_2020_v4_
        nw307_[int(20)] = (d_2020_v4_ if p0 else d_2020_v4_)
        nw307_[int(21)] = d_2020_v4_
        nw307_[int(22)] = d_2020_v4_
        d_2022_v6_ = nw307_
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_2022_v6_).length(0)):
            d_2023_i0_: int = guard_loop_3_
            if (True) and (((0) <= (d_2023_i0_)) and ((d_2023_i0_) < ((d_2022_v6_).length(0)))):
                (d_2022_v6_)[(d_2023_i0_)] = (d_2020_v4_ if p0 else d_2020_v4_)
        d_2024_v7_: _dafny.Array
        def lambda81_(d_2025_i1_):
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sggnxrdx"))

        init51_ = lambda81_
        nw308_ = _dafny.Array(None, 10)
        for i0_51_ in range(nw308_.length(0)):
            nw308_[i0_51_] = init51_(i0_51_)
        d_2024_v7_ = nw308_
        d_2026_v8_: D10
        d_2026_v8_ = D10_DC29(d_2016_v0_, d_2024_v7_, p0, d_2016_v0_)
        d_2027_v9_: _dafny.Array
        nw309_ = _dafny.Array(None, 1)
        nw309_[int(0)] = (d_2026_v8_).cf53
        d_2027_v9_ = nw309_
        d_2028_v10_: _dafny.Map
        d_2028_v10_ = _dafny.Map({d_2027_v9_: p0})
        d_2028_v10_ = (d_2028_v10_).set(d_2027_v9_, p0)
        d_2029_v11_: _dafny.Set
        d_2029_v11_ = _dafny.Set({p1, p1, (p0) or (not(False))})
        d_2029_v11_ = (D11_DC33(-74, p0, (self).fm11(globalState), p0, d_2029_v11_)).cf66
        hi13_ = d_2016_v0_
        for d_2030_i2_ in range((0) - (d_2016_v0_), hi13_):
            d_2031_v12_: _dafny.Seq
            d_2031_v12_ = _dafny.SeqWithoutIsStrInference([428, d_2016_v0_])
            d_2032_v13_: D11
            d_2032_v13_ = D11_DC31(67, (d_2031_v12_)[default__.safeIndex(d_2030_i2_, len(d_2031_v12_))])
            pat_let_tv52_ = d_2017_v1_
            def iife193_(_pat_let70_0):
                def iife194_(d_2033_dt__update__tmp_h0_):
                    def iife195_(_pat_let71_0):
                        def iife196_(d_2034_dt__update_hcf59_h0_):
                            return D11_DC31((d_2033_dt__update__tmp_h0_).cf58, d_2034_dt__update_hcf59_h0_)
                        return iife196_(_pat_let71_0)
                    return iife195_((d_2030_i2_) + (len(pat_let_tv52_)))
                return iife194_(_pat_let70_0)
            rhs354_ = iife193_(d_2032_v13_)
            rhs355_ = (d_2016_v0_) * (d_2016_v0_)
            d_2032_v13_ = rhs354_
            r2 = rhs355_
            d_2035_v14_: _dafny.Seq
            d_2035_v14_ = _dafny.SeqWithoutIsStrInference([d_2019_v3_, d_2019_v3_])
            d_2036_v15_: _dafny.Seq
            d_2036_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iguyj"))
            d_2037_v16_: D4
            d_2037_v16_ = D4_DC12((d_2035_v14_).set(default__.safeIndex(d_2030_i2_, len(d_2035_v14_)), (d_2019_v3_).set(d_2021_v5_, default__.abs(len(d_2036_v15_)))))
            d_2037_v16_ = d_2037_v16_
            d_2038_v17_: _dafny.Array
            nw310_ = _dafny.Array(False, 23)
            d_2038_v17_ = nw310_
            d_2039_v18_: _dafny.Array
            d_2039_v18_ = d_2038_v17_
            d_2040_v19_: _dafny.Map
            d_2040_v19_ = _dafny.Map({d_2016_v0_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "by"))})
            d_2041_v20_: C3
            nw311_ = C3()
            nw311_.ctor__(_dafny.Map({87: d_2016_v0_}), (d_2039_v18_), d_2040_v19_, d_2030_i2_)
            d_2041_v20_ = nw311_
            d_2016_v0_ = d_2016_v0_
        r0 = p1
        d_2042_v21_: _dafny.Map
        d_2042_v21_ = _dafny.Map({d_2016_v0_: p1})
        d_2043_v22_: _dafny.Map
        d_2043_v22_ = _dafny.Map({d_2042_v21_: p1})
        d_2044_v23_: D3
        d_2044_v23_ = D3_DC10(d_2043_v22_)
        pat_let_tv53_ = p0
        pat_let_tv54_ = p1
        pat_let_tv55_ = p0
        def lambda82_(source25_):
            if source25_.is_DC10:
                d_2045___mcc_h0_ = source25_.cf22
                d_2046_cf22_ = d_2045___mcc_h0_
                return pat_let_tv53_
            elif source25_.is_DC9:
                d_2047___mcc_h1_ = source25_.cf21
                d_2048_cf21_ = d_2047___mcc_h1_
                return pat_let_tv54_
            elif True:
                d_2049___mcc_h2_ = source25_.cf23
                d_2050_cf23_ = d_2049___mcc_h2_
                return pat_let_tv55_

        r1 = not(lambda82_(d_2044_v23_))
        r2 = (0) - (d_2016_v0_)
        return r0, r1, r2


class C9(T0):
    def  __init__(self):
        self._f6: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C9"
    @property
    def f6(self):
        return self._f6
    def ctor__(self, f6):
        (self)._f6 = f6

    def fm5(self, p0, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([_dafny.CodePoint('c'), _dafny.CodePoint('v'), _dafny.CodePoint('w')])])) + ((_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([_dafny.CodePoint('q'), _dafny.CodePoint('e')])])) + (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_2051_i0_ in range(default__.abs(26))])), _dafny.MultiSet([_dafny.CodePoint('u'), _dafny.CodePoint('a')]), _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r'), _dafny.CodePoint('a')])), (D1_DC3(_dafny.MultiSet([_dafny.CodePoint('v'), _dafny.CodePoint('w')]))).cf6, _dafny.MultiSet([_dafny.CodePoint('p')])])))

    def fm6(self, p0, globalState):
        return _dafny.Map({(self).f6: _dafny.SeqWithoutIsStrInference([(self).f6])})

    def m3(self, p0, p1, p2, globalState):
        r0: bool = False
        d_2052_v0_: _dafny.Array
        nw312_ = _dafny.Array(_dafny.Seq({}), 5)
        d_2052_v0_ = nw312_
        d_2053_v1_: _dafny.Seq
        d_2053_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hbqce"))
        d_2054_v2_: _dafny.Set
        d_2054_v2_ = _dafny.Set({(self).f6, len(d_2053_v1_)})
        d_2055_v3_: _dafny.Set
        d_2055_v3_ = _dafny.Set({False})
        d_2056_v4_: _dafny.Seq
        d_2056_v4_ = _dafny.SeqWithoutIsStrInference([len(d_2055_v3_)])
        index293_ = default__.safeIndex(841, (d_2052_v0_).length(0))
        (d_2052_v0_)[index293_] = (_dafny.SeqWithoutIsStrInference([len(d_2054_v2_)])) + ((d_2056_v4_).set(default__.safeIndex((self).f6, len(d_2056_v4_)), (self).f6))
        index294_ = default__.safeIndex(841, (d_2052_v0_).length(0))
        (d_2052_v0_)[index294_] = (((d_2056_v4_) + (d_2056_v4_)) + (d_2056_v4_)).set(default__.safeIndex((self).f6, len(((d_2056_v4_) + (d_2056_v4_)) + (d_2056_v4_))), len(d_2053_v1_))
        d_2057_v5_: _dafny.Seq
        d_2057_v5_ = _dafny.SeqWithoutIsStrInference([p2, p2])
        (globalState).f0 = not(((self).f6) >= ((len(d_2057_v5_) if p0 else (self).f6)))
        d_2058_v6_: D0
        d_2058_v6_ = D0_DC0(d_2057_v5_, (self).f6)
        source26_ = d_2058_v6_
        if source26_.is_DC0:
            d_2059___mcc_h0_ = source26_.cf0
            d_2060___mcc_h1_ = source26_.cf1
            d_2061_cf1_ = d_2060___mcc_h1_
            d_2062_cf0_ = d_2059___mcc_h0_
            d_2063_v7_: str
            d_2063_v7_ = _dafny.CodePoint('e')
            d_2064_v8_: _dafny.Map
            d_2064_v8_ = _dafny.Map({(d_2061_cf1_ if p0 else (self).f6): (self).f6})
            d_2065_v9_: _dafny.Map
            d_2065_v9_ = _dafny.Map({p2: _dafny.CodePoint('v')})
            rhs356_ = True
            rhs357_ = p0
            rhs358_ = ((d_2065_v9_)[(d_2053_v1_) <= (d_2053_v1_)] if ((d_2053_v1_) <= (d_2053_v1_)) in (d_2065_v9_) else d_2063_v7_)
            rhs359_ = (d_2064_v8_).set(default__.fm2(d_2061_cf1_, globalState), (self).f6)
            lhs291_ = globalState
            lhs292_ = globalState
            lhs291_.f0 = rhs356_
            lhs292_.f0 = rhs357_
            d_2063_v7_ = rhs358_
            d_2064_v8_ = rhs359_
            (globalState).f0 = False
            if p0:
                d_2066_v10_: _dafny.Seq
                d_2066_v10_ = _dafny.SeqWithoutIsStrInference([d_2053_v1_, _dafny.SeqWithoutIsStrInference([d_2063_v7_ for d_2067_i1_ in range(default__.abs(-641))]), (d_2053_v1_).set(default__.safeIndex(d_2061_cf1_, len(d_2053_v1_)), d_2063_v7_)])
                def iife197_():
                    def iife199_():
                        coll55_ = _dafny.Map()
                        compr_55_: int
                        for compr_55_ in _dafny.IntegerRange(381, -819):
                            d_2069_v12_: int = compr_55_
                            if ((381) <= (d_2069_v12_)) and ((d_2069_v12_) < (-819)):
                                coll55_[(d_2069_v12_) - ((self).f6)] = (_dafny.MultiSet([not(p0), p2])).cardinality
                        return _dafny.Map(coll55_)
                    coll53_ = _dafny.Map()
                    def iife198_():
                        coll54_ = _dafny.Map()
                        compr_54_: int
                        for compr_54_ in _dafny.IntegerRange(381, -819):
                            d_2069_v12_: int = compr_54_
                            if ((381) <= (d_2069_v12_)) and ((d_2069_v12_) < (-819)):
                                coll54_[(d_2069_v12_) - ((self).f6)] = (_dafny.MultiSet([not(p0), p2])).cardinality
                        return _dafny.Map(coll54_)
                    compr_53_: int
                    for compr_53_ in (iife198_()
                    ).keys.Elements:
                        d_2070_v11_: int = compr_53_
                        if (d_2070_v11_) in (iife199_()
                        ):
                            coll53_[(d_2070_v11_) - (d_2061_cf1_)] = d_2061_cf1_
                    return _dafny.Map(coll53_)
                d_2053_v1_ = (_dafny.SeqWithoutIsStrInference([d_2063_v7_ for d_2068_i0_ in range(default__.abs(885))])) + ((d_2066_v10_)[default__.safeIndex(len(iife197_()
                ), len(d_2066_v10_))])
                (globalState).f0 = p2
                d_2071_v13_: _dafny.MultiSet
                d_2071_v13_ = _dafny.MultiSet([p1, p1])
                d_2072_v14_: _dafny.MultiSet
                d_2072_v14_ = _dafny.MultiSet([d_2053_v1_])
                d_2073_v15_: _dafny.Map
                d_2073_v15_ = _dafny.Map({p2: d_2072_v14_})
                d_2074_v16_: _dafny.Map
                d_2074_v16_ = _dafny.Map({(d_2071_v13_).ispropersubset(d_2071_v13_): ((d_2073_v15_)[p0] if (p0) in (d_2073_v15_) else _dafny.MultiSet([d_2053_v1_, d_2053_v1_, d_2053_v1_]))})
                d_2074_v16_ = (d_2074_v16_).set((default__.fm3(globalState)) == (p0), d_2072_v14_)
                (globalState).f0 = (p0) or (False)
                index295_ = default__.safeIndex(540, (p1).length(0))
                (p1)[index295_] = (self).f6
                index296_ = default__.safeIndex(540, (p1).length(0))
                (p1)[index296_] = default__.fm2(default__.fm2(d_2061_cf1_, globalState), globalState)
            elif True:
                d_2075_v17_: _dafny.MultiSet
                d_2075_v17_ = _dafny.MultiSet([len(_dafny.Set({d_2061_cf1_, (self).f6})), (self).f6, (self).f6, d_2061_cf1_])
                (globalState).f0 = (_dafny.MultiSet([(self).f6, (self).f6])).isdisjoint(d_2075_v17_)
                d_2058_v6_ = d_2058_v6_
                d_2064_v8_ = (d_2064_v8_).set((self).f6, (self).f6)
                d_2076_v18_: _dafny.Array
                def lambda83_(d_2077_i2_):
                    return False

                init52_ = lambda83_
                nw313_ = _dafny.Array(None, 1)
                for i0_52_ in range(nw313_.length(0)):
                    nw313_[i0_52_] = init52_(i0_52_)
                d_2076_v18_ = nw313_
                index297_ = default__.safeIndex(903, (d_2076_v18_).length(0))
                (d_2076_v18_)[index297_] = not(p0)
                d_2078_v19_: _dafny.Map
                d_2078_v19_ = _dafny.Map({d_2061_cf1_: False})
                index298_ = default__.safeIndex(903, (d_2076_v18_).length(0))
                (d_2076_v18_)[index298_] = ((d_2078_v19_) | (_dafny.Map({d_2061_cf1_: p0}))) == ((d_2078_v19_) | (d_2078_v19_))
                d_2079_v20_: _dafny.MultiSet
                d_2079_v20_ = _dafny.MultiSet([p2, ((d_2078_v19_)[158] if (158) in (d_2078_v19_) else p0), p2])
                d_2080_v21_: D1
                d_2080_v21_ = D1_DC4(d_2058_v6_, (d_2079_v20_).cardinality, d_2053_v1_, (d_2057_v5_)[default__.safeIndex(d_2061_cf1_, len(d_2057_v5_))], d_2061_cf1_)
                d_2081_v22_: _dafny.Map
                d_2081_v22_ = _dafny.Map({(0) - (default__.fm2(971, globalState)): d_2080_v21_})
                d_2081_v22_ = d_2081_v22_
            d_2082_v23_: _dafny.Array
            nw314_ = _dafny.Array(_dafny.Array(None, 0), 2)
            d_2082_v23_ = nw314_
            d_2083_v24_: _dafny.Array
            nw315_ = _dafny.Array(D0.default()(), 6)
            d_2083_v24_ = nw315_
            index299_ = default__.safeIndex(322, (d_2082_v23_).length(0))
            (d_2082_v23_)[index299_] = d_2083_v24_
            index300_ = default__.safeIndex(322, (d_2082_v23_).length(0))
            def lambda84_(d_2084_v6_):
                def lambda85_(d_2085_i3_):
                    return d_2084_v6_

                return lambda85_

            init53_ = lambda84_(d_2058_v6_)
            nw316_ = _dafny.Array(None, 6)
            for i0_53_ in range(nw316_.length(0)):
                nw316_[i0_53_] = init53_(i0_53_)
            (d_2082_v23_)[index300_] = nw316_
        elif source26_.is_DC1:
            d_2086___mcc_h2_ = source26_.cf2
            d_2087___mcc_h3_ = source26_.cf3
            d_2088___mcc_h4_ = source26_.cf4
            d_2089_cf4_ = d_2088___mcc_h4_
            d_2090_cf3_ = d_2087___mcc_h3_
            d_2091_cf2_ = d_2086___mcc_h2_
            d_2090_cf3_ = (self).f6
            r0 = not (p0) or (p0)
            d_2092_v25_: _dafny.Array
            def lambda86_(d_2093_v1_):
                def lambda87_(d_2094_i4_):
                    return d_2093_v1_

                return lambda87_

            init54_ = lambda86_(d_2053_v1_)
            nw317_ = _dafny.Array(None, 10)
            for i0_54_ in range(nw317_.length(0)):
                nw317_[i0_54_] = init54_(i0_54_)
            d_2092_v25_ = nw317_
            d_2095_v26_: _dafny.Map
            d_2095_v26_ = _dafny.Map({d_2092_v25_: True})
            d_2095_v26_ = (d_2095_v26_).set(d_2092_v25_, p0)
            d_2090_cf3_ = (default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([p2, p2])), d_2090_cf3_) if p0 else d_2090_cf3_)
        elif True:
            d_2096___mcc_h5_ = source26_.cf5
            d_2097_cf5_ = d_2096___mcc_h5_
            (globalState).f5 = (self).f6
            (globalState).f5 = (self).f6
            d_2098_v27_: _dafny.Map
            d_2098_v27_ = _dafny.Map({(self).f6: (self).f6})
            d_2099_v31_: _dafny.Array
            nw318_ = _dafny.Array(None, 16)
            nw318_[int(0)] = d_2098_v27_
            nw318_[int(1)] = (d_2098_v27_).set((self).f6, -510)
            nw318_[int(2)] = d_2098_v27_
            nw318_[int(3)] = d_2098_v27_
            nw318_[int(4)] = (d_2098_v27_) | (d_2098_v27_)
            nw318_[int(5)] = d_2098_v27_
            nw318_[int(6)] = (d_2098_v27_ if p2 else d_2098_v27_)
            def iife200_():
                coll56_ = _dafny.Map()
                compr_56_: int
                for compr_56_ in _dafny.IntegerRange(-1, 458):
                    d_2100_v28_: int = compr_56_
                    if ((-1) <= (d_2100_v28_)) and ((d_2100_v28_) < (458)):
                        coll56_[default__.safeModuloInt(d_2100_v28_, 76)] = (self).f6
                return _dafny.Map(coll56_)
            nw318_[int(7)] = iife200_()
            
            def iife201_():
                coll57_ = _dafny.Map()
                compr_57_: int
                for compr_57_ in _dafny.IntegerRange(52, 438):
                    d_2101_v29_: int = compr_57_
                    if ((52) <= (d_2101_v29_)) and ((d_2101_v29_) < (438)):
                        coll57_[(d_2101_v29_) + ((self).f6)] = (self).f6
                return _dafny.Map(coll57_)
            nw318_[int(8)] = iife201_()
            
            nw318_[int(9)] = default__.fm9(globalState)
            nw318_[int(10)] = (d_2098_v27_) | ((d_2098_v27_).set(len(d_2057_v5_), (self).f6))
            def iife202_():
                coll58_ = _dafny.Map()
                compr_58_: int
                for compr_58_ in (d_2056_v4_).Elements:
                    d_2102_v30_: int = compr_58_
                    if (d_2102_v30_) in (d_2056_v4_):
                        coll58_[(d_2102_v30_) + (len((((d_2052_v0_)[default__.safeIndex(841, (d_2052_v0_).length(0))]).set(default__.safeIndex((self).f6, len((d_2052_v0_)[default__.safeIndex(841, (d_2052_v0_).length(0))])), len(_dafny.SeqWithoutIsStrInference([p2, p2])))).set(default__.safeIndex(len(_dafny.Map({(self).f6: p0})), len(((d_2052_v0_)[default__.safeIndex(841, (d_2052_v0_).length(0))]).set(default__.safeIndex((self).f6, len((d_2052_v0_)[default__.safeIndex(841, (d_2052_v0_).length(0))])), len(_dafny.SeqWithoutIsStrInference([p2, p2]))))), (self).f6)))] = (self).f6
                return _dafny.Map(coll58_)
            nw318_[int(11)] = iife202_()
            
            nw318_[int(12)] = d_2098_v27_
            nw318_[int(13)] = (D2_DC5(d_2098_v27_)).cf12
            nw318_[int(14)] = d_2098_v27_
            nw318_[int(15)] = (d_2098_v27_) | (d_2098_v27_)
            d_2099_v31_ = nw318_
            d_2099_v31_ = (d_2099_v31_ if (p2) or (not(p2)) else d_2099_v31_)
            rhs360_ = (d_2056_v4_)[default__.safeIndex((0) - ((self).f6), len(d_2056_v4_))]
            rhs361_ = p2
            rhs362_ = d_2053_v1_
            rhs363_ = (self).f6
            lhs293_ = globalState
            lhs294_ = globalState
            lhs293_.f5 = rhs360_
            r0 = rhs361_
            d_2053_v1_ = rhs362_
            lhs294_.f5 = rhs363_
        d_2103_v32_: str
        d_2103_v32_ = _dafny.CodePoint('f')
        d_2104_v33_: D0
        d_2104_v33_ = D0_DC1(p0, (self).f6, d_2103_v32_)
        source27_ = (d_2104_v33_ if p0 else d_2104_v33_)
        if source27_.is_DC0:
            d_2105___mcc_h6_ = source27_.cf0
            d_2106___mcc_h7_ = source27_.cf1
            d_2107_cf1_ = d_2106___mcc_h7_
            d_2108_cf0_ = d_2105___mcc_h6_
            d_2107_cf1_ = (self).f6
            d_2109_v34_: _dafny.MultiSet
            d_2109_v34_ = _dafny.MultiSet([d_2103_v32_])
            d_2110_v35_: _dafny.Array
            nw319_ = _dafny.Array(None, 28)
            nw319_[int(0)] = p2
            nw319_[int(1)] = not(p2)
            nw319_[int(2)] = p2
            nw319_[int(3)] = p0
            nw319_[int(4)] = p2
            nw319_[int(5)] = p0
            nw319_[int(6)] = p2
            nw319_[int(7)] = p2
            nw319_[int(8)] = False
            nw319_[int(9)] = p2
            nw319_[int(10)] = default__.fm3(globalState)
            nw319_[int(11)] = default__.fm3(globalState)
            nw319_[int(12)] = p2
            nw319_[int(13)] = False
            nw319_[int(14)] = p0
            nw319_[int(15)] = p2
            nw319_[int(16)] = p0
            nw319_[int(17)] = p2
            nw319_[int(18)] = p2
            nw319_[int(19)] = p0
            nw319_[int(20)] = p0
            nw319_[int(21)] = p0
            nw319_[int(22)] = p2
            nw319_[int(23)] = p0
            nw319_[int(24)] = p2
            nw319_[int(25)] = p2
            nw319_[int(26)] = p0
            nw319_[int(27)] = p2
            d_2110_v35_ = nw319_
            d_2111_v36_: _dafny.Map
            d_2111_v36_ = _dafny.Map({(self).f6: d_2053_v1_})
            d_2112_v37_: _dafny.MultiSet
            d_2112_v37_ = _dafny.MultiSet([_dafny.MultiSet(d_2056_v4_)])
            d_2113_v38_: C4
            nw320_ = C4()
            nw320_.ctor__(d_2107_cf1_, ((d_2109_v34_)[d_2103_v32_] if (d_2103_v32_) in (d_2109_v34_) else len(d_2053_v1_)), (d_2110_v35_ if p0 else d_2110_v35_), (d_2111_v36_) | (d_2111_v36_), d_2107_cf1_, d_2112_v37_)
            d_2113_v38_ = nw320_
            d_2114_v39_: _dafny.Map
            d_2114_v39_ = _dafny.Map({(self).f6: p2})
            d_2115_v40_: D1
            d_2115_v40_ = D1_DC4(d_2058_v6_, default__.fm2(len(d_2114_v39_), globalState), d_2053_v1_, False, (d_2113_v38_).f18)
            d_2116_v41_: _dafny.Map
            d_2116_v41_ = _dafny.Map({d_2115_v40_: d_2110_v35_})
            (globalState).f5 = (0) - (len(d_2116_v41_))
            (globalState).f0 = (p2) == (p2)
        elif source27_.is_DC1:
            d_2117___mcc_h8_ = source27_.cf2
            d_2118___mcc_h9_ = source27_.cf3
            d_2119___mcc_h10_ = source27_.cf4
            d_2120_cf4_ = d_2119___mcc_h10_
            d_2121_cf3_ = d_2118___mcc_h9_
            d_2122_cf2_ = d_2117___mcc_h8_
            d_2123_v42_: _dafny.Map
            d_2123_v42_ = _dafny.Map({d_2104_v33_: not(d_2122_cf2_)})
            (globalState).f5 = ((0) - ((self).f6)) * (len(d_2123_v42_))
            d_2121_cf3_ = (((self).f6) - (169) if (d_2057_v5_)[default__.safeIndex((self).f6, len(d_2057_v5_))] else (self).f6)
            d_2124_v43_: _dafny.Array
            nw321_ = _dafny.Array(int(0), 19)
            d_2124_v43_ = nw321_
            d_2124_v43_ = p1
            d_2125_v44_: _dafny.Array
            def lambda88_(d_2126_p2_, d_2127_p0_):
                def lambda89_(d_2128_i5_):
                    return _dafny.Map({not(not(d_2126_p2_)): d_2127_p0_})

                return lambda89_

            init55_ = lambda88_(p2, p0)
            nw322_ = _dafny.Array(None, 21)
            for i0_55_ in range(nw322_.length(0)):
                nw322_[i0_55_] = init55_(i0_55_)
            d_2125_v44_ = nw322_
            d_2129_v45_: _dafny.Map
            d_2129_v45_ = _dafny.Map({default__.fm3(globalState): p2})
            index301_ = default__.safeIndex(86, (d_2125_v44_).length(0))
            (d_2125_v44_)[index301_] = d_2129_v45_
            index302_ = default__.safeIndex(86, (d_2125_v44_).length(0))
            (d_2125_v44_)[index302_] = d_2129_v45_
        elif True:
            d_2130___mcc_h11_ = source27_.cf5
            d_2131_cf5_ = d_2130___mcc_h11_
            (globalState).f5 = ((self).f6) + ((self).f6)
            d_2132_v46_: _dafny.Map
            d_2132_v46_ = _dafny.Map({(self).f6: len((d_2053_v1_).set(default__.safeIndex((self).f6, len(d_2053_v1_)), d_2103_v32_))})
            d_2132_v46_ = (d_2132_v46_).set((self).f6, (self).f6)
            d_2133_v47_: _dafny.Map
            d_2133_v47_ = _dafny.Map({789: not(p0)})
            rhs364_ = (_dafny.Set({(self).f6, (self).f6}) if ((d_2133_v47_)[(self).f6] if ((self).f6) in (d_2133_v47_) else p2) else _dafny.Set({(self).f6}))
            rhs365_ = ((self).f6) >= ((self).f6)
            d_2054_v2_ = rhs364_
            r0 = rhs365_
            (globalState).f0 = p2
        d_2134_v48_: _dafny.Map
        d_2134_v48_ = _dafny.Map({p2: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "icjgt")))})
        d_2135_v49_: D14
        d_2135_v49_ = D14_DC36(d_2134_v48_)
        d_2136_v50_: _dafny.Array
        nw323_ = _dafny.Array(None, 2)
        nw323_[int(0)] = (d_2135_v49_ if not(p2) else d_2135_v49_)
        nw323_[int(1)] = d_2135_v49_
        d_2136_v50_ = nw323_
        d_2137_v51_: _dafny.Seq
        d_2137_v51_ = _dafny.SeqWithoutIsStrInference([default__.fm41(globalState), _dafny.Map({p0: p0})])
        d_2138_v52_: _dafny.Map
        d_2138_v52_ = _dafny.Map({p2: p0})
        d_2139_v53_: _dafny.MultiSet
        d_2139_v53_ = _dafny.MultiSet([d_2138_v52_, d_2138_v52_])
        rhs366_ = 6
        rhs367_ = (d_2139_v53_).issubset((_dafny.MultiSet(d_2137_v51_)) - (d_2139_v53_))
        rhs368_ = d_2136_v50_
        lhs295_ = globalState
        lhs296_ = globalState
        lhs295_.f5 = rhs366_
        lhs296_.f0 = rhs367_
        d_2136_v50_ = rhs368_
        d_2140_v54_: _dafny.MultiSet
        d_2140_v54_ = _dafny.MultiSet([p0, default__.fm3(globalState)])
        d_2141_v55_: _dafny.Map
        d_2141_v55_ = _dafny.Map({d_2140_v54_: _dafny.CodePoint('a')})
        d_2142_v56_: _dafny.Seq
        d_2142_v56_ = _dafny.SeqWithoutIsStrInference([d_2141_v55_])
        if (default__.safeModuloInt(len((d_2142_v56_)[default__.safeIndex((self).f6, len(d_2142_v56_))]), (self).f6)) == ((self).f6):
            d_2143_v57_: C6
            nw324_ = C6()
            nw324_.ctor__((self).f6, True)
            d_2143_v57_ = nw324_
            d_2144_v58_: D6
            d_2144_v58_ = D6_DC19((self).f6, (self).f6, len(d_2054_v2_))
            pat_let_tv56_ = d_2143_v57_
            def iife203_(_pat_let72_0):
                def iife204_(d_2145_dt__update__tmp_h0_):
                    def iife205_(_pat_let73_0):
                        def iife206_(d_2146_dt__update_hcf41_h0_):
                            return D6_DC19((d_2145_dt__update__tmp_h0_).cf39, (d_2145_dt__update__tmp_h0_).cf40, d_2146_dt__update_hcf41_h0_)
                        return iife206_(_pat_let73_0)
                    return iife205_((pat_let_tv56_).f13)
                return iife204_(_pat_let72_0)
            d_2144_v58_ = iife203_(D6_DC19(len(d_2053_v1_), (d_2143_v57_).f13, (d_2143_v57_).f13))
            d_2147_v59_: _dafny.Map
            d_2147_v59_ = _dafny.Map({((d_2143_v57_).f14) and (p2): (_dafny.Map({not(not(p2)): -942})).set(p2, (self).f6)})
            d_2147_v59_ = (d_2147_v59_).set(not((d_2143_v57_).f14), ((d_2134_v48_).set(True, (d_2143_v57_).f13)).set(False, (self).f6))
            d_2148_v60_: _dafny.Array
            nw325_ = _dafny.Array(_dafny.Array(None, 0), 2)
            d_2148_v60_ = nw325_
            d_2148_v60_ = d_2148_v60_
            d_2149_v61_: _dafny.Array
            def lambda90_(d_2150_v6_):
                def lambda91_(d_2151_i6_):
                    return d_2150_v6_

                return lambda91_

            init56_ = lambda90_(d_2058_v6_)
            nw326_ = _dafny.Array(None, 22)
            for i0_56_ in range(nw326_.length(0)):
                nw326_[i0_56_] = init56_(i0_56_)
            d_2149_v61_ = nw326_
            d_2152_v62_: D7
            d_2152_v62_ = D7_DC21(d_2149_v61_)
            source28_ = d_2152_v62_
            if source28_.is_DC22:
                index303_ = default__.safeIndex(11, (p1).length(0))
                (p1)[index303_] = ((d_2143_v57_).f13) - ((self).f6)
                index304_ = default__.safeIndex(11, (p1).length(0))
                (p1)[index304_] = (self).f6
                rhs369_ = p2
                rhs370_ = (0) - ((d_2056_v4_)[default__.safeIndex((p1)[default__.safeIndex(11, (p1).length(0))], len(d_2056_v4_))])
                rhs371_ = d_2053_v1_
                rhs372_ = (self).f6
                lhs297_ = globalState
                lhs298_ = globalState
                lhs299_ = globalState
                lhs300_ = globalState
                lhs297_.f0 = rhs369_
                lhs298_.f5 = rhs370_
                lhs299_.f2 = rhs371_
                lhs300_.f5 = rhs372_
                d_2153_v63_: C1
                nw327_ = C1()
                nw327_.ctor__((self).f6)
                d_2153_v63_ = nw327_
                (globalState).f0 = p0
            elif True:
                d_2154___mcc_h12_ = source28_.cf43
                d_2155_cf43_ = d_2154___mcc_h12_
                (globalState).f0 = (d_2143_v57_).f14
                d_2156_v64_: C1
                nw328_ = C1()
                nw328_.ctor__(469)
                d_2156_v64_ = nw328_
                d_2157_v65_: _dafny.Map
                d_2157_v65_ = _dafny.Map({d_2134_v48_: d_2143_v57_})
                d_2158_v66_: _dafny.Map
                d_2158_v66_ = _dafny.Map({p2: d_2143_v57_})
                d_2159_v67_: _dafny.Array
                nw329_ = _dafny.Array(None, 14)
                nw329_[int(0)] = ((d_2157_v65_)[d_2134_v48_] if (d_2134_v48_) in (d_2157_v65_) else d_2143_v57_)
                nw329_[int(1)] = d_2143_v57_
                nw329_[int(2)] = d_2143_v57_
                nw329_[int(3)] = d_2143_v57_
                nw329_[int(4)] = d_2143_v57_
                nw329_[int(5)] = d_2143_v57_
                nw329_[int(6)] = d_2143_v57_
                nw329_[int(7)] = d_2143_v57_
                nw329_[int(8)] = d_2143_v57_
                nw329_[int(9)] = d_2143_v57_
                nw329_[int(10)] = d_2143_v57_
                nw329_[int(11)] = d_2143_v57_
                nw329_[int(12)] = ((d_2158_v66_)[(d_2156_v64_).fm28(globalState)] if ((d_2156_v64_).fm28(globalState)) in (d_2158_v66_) else d_2143_v57_)
                nw329_[int(13)] = d_2143_v57_
                d_2159_v67_ = nw329_
                d_2159_v67_ = d_2159_v67_
                d_2138_v52_ = (d_2138_v52_).set((d_2143_v57_).f14, ((d_2143_v57_).f13) >= ((d_2143_v57_).f13))
        elif True:
            d_2160_v68_: _dafny.Map
            d_2160_v68_ = _dafny.Map({-195: d_2138_v52_})
            d_2161_v69_: _dafny.MultiSet
            d_2161_v69_ = _dafny.MultiSet([p1, p1])
            d_2162_v71_: _dafny.Seq
            d_2162_v71_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({(self).f6: d_2138_v52_}), d_2160_v68_])
            d_2163_v72_: _dafny.Array
            nw330_ = _dafny.Array(None, 29)
            nw330_[int(0)] = d_2160_v68_
            nw330_[int(1)] = d_2160_v68_
            nw330_[int(2)] = (d_2160_v68_) | (d_2160_v68_)
            nw330_[int(3)] = d_2160_v68_
            nw330_[int(4)] = d_2160_v68_
            nw330_[int(5)] = d_2160_v68_
            nw330_[int(6)] = (d_2160_v68_).set((d_2161_v69_).cardinality, d_2138_v52_)
            nw330_[int(7)] = d_2160_v68_
            nw330_[int(8)] = d_2160_v68_
            nw330_[int(9)] = (d_2160_v68_) | (d_2160_v68_)
            nw330_[int(10)] = d_2160_v68_
            nw330_[int(11)] = d_2160_v68_
            def iife207_():
                coll59_ = _dafny.Map()
                compr_59_: int
                for compr_59_ in _dafny.IntegerRange(604, -336):
                    d_2164_v70_: int = compr_59_
                    if ((604) <= (d_2164_v70_)) and ((d_2164_v70_) < (-336)):
                        coll59_[(d_2164_v70_) + ((self).f6)] = d_2138_v52_
                return _dafny.Map(coll59_)
            nw330_[int(12)] = iife207_()
            
            nw330_[int(13)] = d_2160_v68_
            nw330_[int(14)] = (d_2160_v68_).set(len(d_2057_v5_), d_2138_v52_)
            nw330_[int(15)] = d_2160_v68_
            nw330_[int(16)] = d_2160_v68_
            nw330_[int(17)] = (d_2160_v68_) | (d_2160_v68_)
            nw330_[int(18)] = (d_2160_v68_) | ((d_2162_v71_)[default__.safeIndex(len((d_2052_v0_)[default__.safeIndex(841, (d_2052_v0_).length(0))]), len(d_2162_v71_))])
            nw330_[int(19)] = (d_2160_v68_) | (d_2160_v68_)
            nw330_[int(20)] = d_2160_v68_
            nw330_[int(21)] = d_2160_v68_
            nw330_[int(22)] = d_2160_v68_
            nw330_[int(23)] = _dafny.Map({854: _dafny.Map({p0: p2})})
            nw330_[int(24)] = d_2160_v68_
            nw330_[int(25)] = d_2160_v68_
            nw330_[int(26)] = d_2160_v68_
            nw330_[int(27)] = _dafny.Map({(self).f6: (d_2137_v51_)[default__.safeIndex(len(d_2053_v1_), len(d_2137_v51_))]})
            nw330_[int(28)] = d_2160_v68_
            d_2163_v72_ = nw330_
            index305_ = default__.safeIndex(896, (d_2163_v72_).length(0))
            (d_2163_v72_)[index305_] = _dafny.Map({(self).f6: _dafny.Map({p0: p0})})
            d_2165_v73_: D4
            d_2165_v73_ = D4_DC14(False, d_2138_v52_, (self).f6, d_2138_v52_)
            index306_ = default__.safeIndex(896, (d_2163_v72_).length(0))
            (d_2163_v72_)[index306_] = (_dafny.Map({(self).f6: (d_2165_v73_).cf29})) | (d_2160_v68_)
            (globalState).f5 = (self).f6
            if (p0) or (p0):
                d_2166_v74_: _dafny.Seq
                d_2166_v74_ = _dafny.SeqWithoutIsStrInference([d_2057_v5_, d_2057_v5_, _dafny.SeqWithoutIsStrInference([p2]), d_2057_v5_])
                d_2166_v74_ = d_2166_v74_
                (globalState).f0 = p0
                d_2167_v75_: C5
                nw331_ = C5()
                nw331_.ctor__((self).f6, (self).f6)
                d_2167_v75_ = nw331_
                d_2168_v76_: _dafny.Map
                d_2168_v76_ = _dafny.Map({((self).f6) not in (default__.fm26((self).f6, d_2103_v32_, (d_2057_v5_)[default__.safeIndex((self).f6, len(d_2057_v5_))], d_2103_v32_, globalState)): d_2053_v1_})
                d_2168_v76_ = (d_2168_v76_).set(p0, d_2053_v1_)
                d_2169_v77_: _dafny.Seq
                d_2169_v77_ = _dafny.SeqWithoutIsStrInference([d_2140_v54_, d_2140_v54_, d_2140_v54_, _dafny.MultiSet(d_2057_v5_), (d_2140_v54_).set(p2, default__.abs((0) - ((self).f6)))])
                d_2170_v78_: _dafny.Set
                d_2170_v78_ = _dafny.Set({d_2056_v4_, d_2056_v4_})
                d_2171_v79_: _dafny.Map
                d_2171_v79_ = _dafny.Map({d_2170_v78_: d_2169_v77_})
                d_2169_v77_ = ((_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([p0])).set(p0, default__.abs(683))])) + (d_2169_v77_)) + (((d_2171_v79_)[d_2170_v78_] if (d_2170_v78_) in (d_2171_v79_) else d_2169_v77_))
            elif True:
                d_2172_v80_: _dafny.Array
                nw332_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 28)
                d_2172_v80_ = nw332_
                d_2173_v81_: _dafny.Map
                d_2173_v81_ = _dafny.Map({d_2103_v32_: d_2172_v80_})
                d_2173_v81_ = (d_2173_v81_).set(default__.fm0((self).f6, d_2054_v2_, p2, globalState), d_2172_v80_)
                d_2174_v82_: _dafny.Map
                d_2174_v82_ = _dafny.Map({(self).f6: p2})
                d_2175_v83_: _dafny.MultiSet
                d_2175_v83_ = _dafny.MultiSet([(_dafny.MultiSet([p2, True, ((d_2174_v82_)[(self).f6] if ((self).f6) in (d_2174_v82_) else not(p0))])).cardinality, (843) - (len(d_2053_v1_)), ((self).f6) * ((self).f6), (self).f6, ((self).f6) + (758)])
                d_2176_v84_: _dafny.Seq
                d_2176_v84_ = _dafny.SeqWithoutIsStrInference([d_2053_v1_, d_2053_v1_])
                rhs373_ = p2
                rhs374_ = (default__.safeModuloInt(-848, (self).f6)) >= ((self).f6)
                rhs375_ = ((self).f6) + ((self).f6)
                rhs376_ = ((d_2176_v84_) + (d_2176_v84_)) <= (d_2176_v84_)
                rhs377_ = d_2175_v83_
                lhs301_ = globalState
                lhs302_ = globalState
                lhs303_ = globalState
                lhs304_ = globalState
                lhs301_.f0 = rhs373_
                lhs302_.f0 = rhs374_
                lhs303_.f5 = rhs375_
                lhs304_.f0 = rhs376_
                d_2175_v83_ = rhs377_
                d_2174_v82_ = (d_2174_v82_).set((self).f6, ((self).f6) < ((self).f6))
                (globalState).f0 = p2
                index307_ = default__.safeIndex(274, (p1).length(0))
                (p1)[index307_] = (self).f6
                d_2177_v85_: _dafny.Seq
                d_2177_v85_ = _dafny.SeqWithoutIsStrInference([d_2104_v33_])
                index308_ = default__.safeIndex(274, (p1).length(0))
                (p1)[index308_] = len((d_2177_v85_) + (d_2177_v85_))
            d_2178_v86_: _dafny.Seq
            d_2178_v86_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p2])])
            d_2179_v87_: _dafny.Seq
            d_2179_v87_ = _dafny.SeqWithoutIsStrInference([d_2178_v86_, d_2178_v86_, d_2178_v86_, d_2178_v86_])
            d_2180_v88_: _dafny.Map
            d_2180_v88_ = _dafny.Map({(d_2179_v87_)[default__.safeIndex((self).f6, len(d_2179_v87_))]: p0})
            d_2180_v88_ = (d_2180_v88_).set(_dafny.SeqWithoutIsStrInference([d_2057_v5_]), default__.fm3(globalState))
            d_2181_v89_: _dafny.Array
            nw333_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 4)
            d_2181_v89_ = nw333_
            index309_ = default__.safeIndex(24, (d_2181_v89_).length(0))
            (d_2181_v89_)[index309_] = (d_2053_v1_) + (d_2053_v1_)
            index310_ = default__.safeIndex(24, (d_2181_v89_).length(0))
            (d_2181_v89_)[index310_] = d_2053_v1_
        r0 = p0
        return r0


class C10(T0):
    def  __init__(self):
        self._f6: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C10"
    @property
    def f6(self):
        return self._f6
    def ctor__(self, f6):
        (self)._f6 = f6

    def fm5(self, p0, globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o')]))])) + (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet((D1_DC4(D0_DC0(_dafny.SeqWithoutIsStrInference([True, not(True)]), 796), (self).f6, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lxfutcy")), False, (self).f6)).cf9) for d_2182_i0_ in range(default__.abs(771))]))) + ((_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([_dafny.CodePoint('y'), _dafny.CodePoint('l')])]) if False else _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([_dafny.CodePoint('c')]), _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p')]))])))

    def fm6(self, p0, globalState):
        def iife208_():
            coll60_ = _dafny.Map()
            compr_60_: int
            for compr_60_ in _dafny.IntegerRange(858, 166):
                d_2183_v0_: int = compr_60_
                if ((858) <= (d_2183_v0_)) and ((d_2183_v0_) < (166)):
                    coll60_[default__.safeModuloInt(d_2183_v0_, (self).f6)] = _dafny.SeqWithoutIsStrInference([(self).f6, (self).f6])
            return _dafny.Map(coll60_)
        return (_dafny.Map({(self).f6: _dafny.SeqWithoutIsStrInference([471, len(_dafny.SeqWithoutIsStrInference([False, False, False, False, False]))])})) | (iife208_()
        )

    def m1(self, p0, globalState):
        r0: bool = False
        d_2184_v0_: _dafny.Seq
        d_2184_v0_ = _dafny.SeqWithoutIsStrInference([len(p0), (self).f6])
        d_2185_v1_: bool
        d_2185_v1_ = False
        d_2186_v2_: _dafny.Seq
        d_2186_v2_ = _dafny.SeqWithoutIsStrInference([d_2185_v1_, d_2185_v1_])
        d_2187_v3_: D0
        d_2187_v3_ = D0_DC0(d_2186_v2_, (self).f6)
        d_2188_v4_: _dafny.Seq
        d_2188_v4_ = _dafny.SeqWithoutIsStrInference([((self).f6) * ((d_2184_v0_)[default__.safeIndex((d_2187_v3_).cf1, len(d_2184_v0_))])])
        (globalState).f5 = (d_2184_v0_)[default__.safeIndex((self).f6, len(d_2184_v0_))]
        d_2189_v5_: _dafny.Array
        def lambda92_(d_2190_v1_):
            def lambda93_(d_2191_i0_):
                return d_2190_v1_

            return lambda93_

        init57_ = lambda92_(d_2185_v1_)
        nw334_ = _dafny.Array(None, 9)
        for i0_57_ in range(nw334_.length(0)):
            nw334_[i0_57_] = init57_(i0_57_)
        d_2189_v5_ = nw334_
        d_2189_v5_ = d_2189_v5_
        d_2192_i1_: int
        d_2192_i1_ = 0
        with _dafny.label("11"):
            while d_2185_v1_:
                with _dafny.c_label("11"):
                    if (d_2192_i1_) >= (100):
                        raise _dafny.Break("11")
                    d_2192_i1_ = (d_2192_i1_) + (1)
                    d_2193_v6_: _dafny.Map
                    d_2193_v6_ = _dafny.Map({7: d_2185_v1_})
                    d_2194_v7_: C6
                    nw335_ = C6()
                    nw335_.ctor__(len(d_2193_v6_), d_2185_v1_)
                    d_2194_v7_ = nw335_
                    d_2195_v8_: _dafny.Map
                    d_2195_v8_ = _dafny.Map({(self).f6: d_2194_v7_})
                    d_2196_v9_: D18
                    d_2196_v9_ = D18_DC42(d_2195_v8_)
                    d_2197_v10_: _dafny.MultiSet
                    d_2197_v10_ = _dafny.MultiSet([len((d_2196_v9_).cf77)])
                    d_2198_v11_: _dafny.Map
                    d_2198_v11_ = _dafny.Map({(self).f6: (0) - (((d_2197_v10_)[(self).f6] if ((self).f6) in (d_2197_v10_) else -417))})
                    d_2199_v12_: _dafny.Map
                    d_2199_v12_ = _dafny.Map({(self).f6: p0})
                    d_2200_v13_: C3
                    nw336_ = C3()
                    nw336_.ctor__(d_2198_v11_, d_2189_v5_, (d_2199_v12_) | (d_2199_v12_), (d_2194_v7_).f13)
                    d_2200_v13_ = nw336_
                    d_2201_v14_: _dafny.Map
                    d_2201_v14_ = _dafny.Map({(d_2194_v7_).f14: (self).f6})
                    d_2202_v15_: _dafny.Seq
                    d_2202_v15_ = _dafny.SeqWithoutIsStrInference([d_2188_v4_])
                    d_2203_v16_: _dafny.Set
                    d_2203_v16_ = _dafny.Set({(d_2194_v7_).f14})
                    d_2204_v17_: _dafny.Array
                    nw337_ = _dafny.Array(None, 28)
                    nw337_[int(0)] = -254
                    nw337_[int(1)] = (self).f6
                    nw337_[int(2)] = ((d_2197_v10_)[(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f6, 953, 72]))).cardinality] if ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f6, 953, 72]))).cardinality) in (d_2197_v10_) else (self).f6)
                    nw337_[int(3)] = (self).f6
                    nw337_[int(4)] = (d_2194_v7_).f13
                    nw337_[int(5)] = (d_2194_v7_).f13
                    nw337_[int(6)] = (0) - ((0) - ((((d_2197_v10_)[(d_2188_v4_)[default__.safeIndex((d_2194_v7_).f13, len(d_2188_v4_))]] if ((d_2188_v4_)[default__.safeIndex((d_2194_v7_).f13, len(d_2188_v4_))]) in (d_2197_v10_) else (0) - (len(d_2186_v2_)))) + (len(d_2201_v14_))))
                    nw337_[int(7)] = (self).f6
                    nw337_[int(8)] = (self).f6
                    nw337_[int(9)] = ((self).f6 if (d_2194_v7_).f14 else (0) - ((self).f6))
                    nw337_[int(10)] = (self).f6
                    nw337_[int(11)] = (d_2194_v7_).f13
                    nw337_[int(12)] = (d_2194_v7_).f13
                    nw337_[int(13)] = (d_2194_v7_).f13
                    nw337_[int(14)] = len((d_2184_v0_) + ((d_2202_v15_)[default__.safeIndex(len(d_2203_v16_), len(d_2202_v15_))]))
                    nw337_[int(15)] = 175
                    nw337_[int(16)] = default__.safeDivisionInt(509, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_2205_i2_ in range(default__.abs(352))])))
                    nw337_[int(17)] = default__.safeDivisionInt((self).f6, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yl"))))
                    nw337_[int(18)] = (self).f6
                    nw337_[int(19)] = (0) - ((0) - ((self).f6))
                    nw337_[int(20)] = len(p0)
                    nw337_[int(21)] = ((d_2194_v7_).f13) + ((default__.fm19((d_2194_v7_).f13, globalState)).cardinality)
                    nw337_[int(22)] = (d_2194_v7_).f13
                    nw337_[int(23)] = default__.safeDivisionInt((d_2194_v7_).f13, (0) - (len(_dafny.SeqWithoutIsStrInference([(d_2194_v7_).f13]))))
                    nw337_[int(24)] = (d_2194_v7_).f13
                    nw337_[int(25)] = 118
                    nw337_[int(26)] = (0) - (default__.safeDivisionInt((0) - (len(p0)), len((D5_DC17(d_2188_v4_, d_2185_v1_)).cf36)))
                    nw337_[int(27)] = default__.safeModuloInt((0) - ((self).f6), (d_2194_v7_).f13)
                    d_2204_v17_ = nw337_
                    d_2206_v18_: D2
                    d_2206_v18_ = D2_DC8(d_2185_v1_, False)
                    d_2207_v19_: str
                    d_2207_v19_ = _dafny.CodePoint('f')
                    index311_ = default__.safeIndex(736, (d_2204_v17_).length(0))
                    (d_2204_v17_)[index311_] = (d_2194_v7_).fm16(d_2206_v18_, D0_DC1(d_2185_v1_, (self).f6, d_2207_v19_), globalState)
                    index312_ = default__.safeIndex(736, (d_2204_v17_).length(0))
                    (d_2204_v17_)[index312_] = default__.safeDivisionInt(default__.safeModuloInt((self).f6, -30), (d_2194_v7_).f13)
                    d_2208_v20_: _dafny.Map
                    d_2208_v20_ = _dafny.Map({p0: not(not(d_2185_v1_))})
                    d_2209_v21_: _dafny.Seq
                    d_2209_v21_ = _dafny.SeqWithoutIsStrInference([p0])
                    d_2210_v22_: C8
                    nw338_ = C8()
                    nw338_.ctor__()
                    d_2210_v22_ = nw338_
                    d_2211_v23_: _dafny.Seq
                    d_2211_v23_ = _dafny.SeqWithoutIsStrInference([d_2210_v22_])
                    d_2208_v20_ = (d_2208_v20_).set(((d_2209_v21_)[default__.safeIndex(len(d_2211_v23_), len(d_2209_v21_))]) + (p0), (d_2194_v7_).f14)
                    d_2212_v24_: D9
                    d_2212_v24_ = D9_DC25(765, ((d_2194_v7_).f13 if (d_2194_v7_).f14 else (self).f6), (d_2204_v17_)[default__.safeIndex(736, (d_2204_v17_).length(0))], d_2207_v19_)
                    d_2213_v25_: _dafny.MultiSet
                    d_2213_v25_ = _dafny.MultiSet([d_2197_v10_])
                    d_2214_v26_: _dafny.Seq
                    d_2214_v26_ = _dafny.SeqWithoutIsStrInference([default__.fm4((self).f6, (d_2204_v17_)[default__.safeIndex(736, (d_2204_v17_).length(0))], (d_2194_v7_).f13, globalState)])
                    d_2215_v27_: _dafny.Seq
                    d_2215_v27_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet((d_2184_v0_).set(default__.safeIndex((d_2204_v17_)[default__.safeIndex(736, (d_2204_v17_).length(0))], len(d_2184_v0_)), (d_2204_v17_)[default__.safeIndex(736, (d_2204_v17_).length(0))])), d_2197_v10_])
                    rhs378_ = ((d_2186_v2_)[default__.safeIndex((self).f6, len(d_2186_v2_))]) and ((d_2186_v2_)[default__.safeIndex((d_2212_v24_).cf48, len(d_2186_v2_))])
                    rhs379_ = d_2189_v5_
                    rhs380_ = d_2212_v24_
                    rhs381_ = (_dafny.SeqWithoutIsStrInference([d_2186_v2_])) < (d_2214_v26_)
                    rhs382_ = ((d_2213_v25_) - (_dafny.MultiSet(d_2215_v27_))) | (d_2213_v25_)
                    lhs305_ = globalState
                    lhs306_ = globalState
                    lhs305_.f0 = rhs378_
                    d_2189_v5_ = rhs379_
                    d_2212_v24_ = rhs380_
                    lhs306_.f0 = rhs381_
                    d_2213_v25_ = rhs382_
                    pass
            pass
        d_2216_v28_: str
        d_2216_v28_ = _dafny.CodePoint('q')
        d_2216_v28_ = d_2216_v28_
        (globalState).f2 = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "henupki"))
        d_2217_v29_: _dafny.Map
        d_2217_v29_ = _dafny.Map({(self).f6: (self).f6})
        d_2218_v30_: D6
        d_2218_v30_ = D6_DC19(len(d_2217_v29_), len(_dafny.SeqWithoutIsStrInference([(self).f6, (self).f6, (self).f6])), (self).f6)
        d_2219_v31_: D6
        d_2219_v31_ = D6_DC20(d_2218_v30_)
        pat_let_tv57_ = d_2218_v30_
        d_2220_v32_: _dafny.Array
        nw339_ = _dafny.Array(None, 29)
        nw339_[int(0)] = D6_DC20(d_2218_v30_)
        nw339_[int(1)] = D6_DC20(d_2218_v30_)
        nw339_[int(2)] = d_2219_v31_
        nw339_[int(3)] = d_2219_v31_
        def iife209_(_pat_let74_0):
            def iife210_(d_2221_dt__update__tmp_h0_):
                def iife211_(_pat_let75_0):
                    def iife212_(d_2222_dt__update_hcf42_h0_):
                        return D6_DC20(d_2222_dt__update_hcf42_h0_)
                    return iife212_(_pat_let75_0)
                return iife211_(pat_let_tv57_)
            return iife210_(_pat_let74_0)
        nw339_[int(4)] = iife209_(d_2219_v31_)
        nw339_[int(5)] = d_2219_v31_
        nw339_[int(6)] = d_2219_v31_
        nw339_[int(7)] = D6_DC20(D6_DC20(d_2218_v30_))
        nw339_[int(8)] = D6_DC20(d_2218_v30_)
        nw339_[int(9)] = d_2219_v31_
        nw339_[int(10)] = d_2219_v31_
        nw339_[int(11)] = d_2219_v31_
        nw339_[int(12)] = d_2219_v31_
        nw339_[int(13)] = d_2219_v31_
        nw339_[int(14)] = d_2219_v31_
        nw339_[int(15)] = d_2219_v31_
        nw339_[int(16)] = d_2219_v31_
        nw339_[int(17)] = d_2219_v31_
        nw339_[int(18)] = d_2219_v31_
        nw339_[int(19)] = d_2219_v31_
        nw339_[int(20)] = d_2219_v31_
        nw339_[int(21)] = d_2219_v31_
        nw339_[int(22)] = d_2219_v31_
        nw339_[int(23)] = d_2219_v31_
        nw339_[int(24)] = d_2219_v31_
        nw339_[int(25)] = d_2219_v31_
        nw339_[int(26)] = d_2219_v31_
        nw339_[int(27)] = d_2219_v31_
        nw339_[int(28)] = d_2219_v31_
        d_2220_v32_ = nw339_
        d_2223_v33_: _dafny.Array
        d_2223_v33_ = d_2220_v32_
        source29_ = d_2223_v33_
        d_2224___mcc_h0_ = source29_
        d_2225_cf72_ = d_2224___mcc_h0_
        d_2226_v34_: D1
        d_2226_v34_ = D1_DC4(d_2187_v3_, (self).f6, p0, d_2185_v1_, default__.fm2((self).f6, globalState))
        source30_ = d_2226_v34_
        if source30_.is_DC4:
            d_2227___mcc_h1_ = source30_.cf7
            d_2228___mcc_h2_ = source30_.cf8
            d_2229___mcc_h3_ = source30_.cf9
            d_2230___mcc_h4_ = source30_.cf10
            d_2231___mcc_h5_ = source30_.cf11
            d_2232_cf11_ = d_2231___mcc_h5_
            d_2233_cf10_ = d_2230___mcc_h4_
            d_2234_cf9_ = d_2229___mcc_h3_
            d_2235_cf8_ = d_2228___mcc_h2_
            d_2236_cf7_ = d_2227___mcc_h1_
            (globalState).f5 = d_2235_cf8_
            d_2237_v35_: _dafny.MultiSet
            d_2237_v35_ = _dafny.MultiSet([d_2232_cf11_])
            d_2238_v36_: _dafny.Seq
            d_2238_v36_ = _dafny.SeqWithoutIsStrInference([d_2237_v35_])
            d_2239_v37_: _dafny.MultiSet
            d_2239_v37_ = _dafny.MultiSet([d_2185_v1_])
            d_2240_v38_: _dafny.Seq
            d_2240_v38_ = _dafny.SeqWithoutIsStrInference([d_2237_v35_, d_2237_v35_, d_2237_v35_, d_2237_v35_, (d_2238_v36_)[default__.safeIndex((d_2239_v37_).cardinality, len(d_2238_v36_))]])
            d_2241_v39_: _dafny.Array
            nw340_ = _dafny.Array(None, 15)
            nw340_[int(0)] = default__.fm2(d_2232_cf11_, globalState)
            nw340_[int(1)] = -393
            nw340_[int(2)] = d_2232_cf11_
            nw340_[int(3)] = -989
            nw340_[int(4)] = ((self).f6) + (999)
            nw340_[int(5)] = default__.safeModuloInt(d_2235_cf8_, d_2232_cf11_)
            nw340_[int(6)] = (0) - ((((d_2217_v29_)[len(d_2234_cf9_)] if (len(d_2234_cf9_)) in (d_2217_v29_) else (d_2237_v35_).cardinality)) * (len(d_2238_v36_)))
            nw340_[int(7)] = (len(p0)) * (d_2232_cf11_)
            nw340_[int(8)] = len(_dafny.Map({d_2235_cf8_: (d_2186_v2_)[default__.safeIndex(d_2232_cf11_, len(d_2186_v2_))]}))
            nw340_[int(9)] = (self).f6
            nw340_[int(10)] = ((d_2237_v35_)[(self).f6] if ((self).f6) in (d_2237_v35_) else d_2232_cf11_)
            nw340_[int(11)] = (d_2235_cf8_) - ((0) - (d_2232_cf11_))
            nw340_[int(12)] = 664
            nw340_[int(13)] = (d_2239_v37_).cardinality
            nw340_[int(14)] = d_2235_cf8_
            d_2241_v39_ = nw340_
            index313_ = default__.safeIndex(356, (d_2241_v39_).length(0))
            (d_2241_v39_)[index313_] = d_2235_cf8_
            d_2242_v40_: D2
            d_2242_v40_ = D2_DC6(d_2232_cf11_, d_2241_v39_)
            d_2243_v41_: _dafny.MultiSet
            d_2243_v41_ = _dafny.MultiSet([(d_2242_v40_).cf14])
            index314_ = default__.safeIndex(356, (d_2241_v39_).length(0))
            (d_2241_v39_)[index314_] = ((d_2243_v41_) | (_dafny.MultiSet([d_2241_v39_, d_2241_v39_]))).cardinality
            (globalState).f0 = d_2233_cf10_
            d_2244_v42_: _dafny.Map
            d_2244_v42_ = _dafny.Map({d_2232_cf11_: p0})
            d_2245_v43_: C7
            nw341_ = C7()
            nw341_.ctor__(d_2241_v39_, d_2241_v39_, default__.fm2((self).f6, globalState), d_2189_v5_, d_2244_v42_)
            d_2245_v43_ = nw341_
            rhs383_ = d_2245_v43_
            rhs384_ = d_2232_cf11_
            d_2245_v43_ = rhs383_
            d_2235_cf8_ = rhs384_
        elif True:
            d_2246___mcc_h6_ = source30_.cf6
            d_2247_cf6_ = d_2246___mcc_h6_
            (globalState).f0 = not((d_2216_v28_) not in (p0))
            d_2187_v3_ = d_2187_v3_
            (globalState).f0 = ((self).f6) > (322)
            d_2248_v44_: C1
            nw342_ = C1()
            nw342_.ctor__((self).f6)
            d_2248_v44_ = nw342_
        index315_ = default__.safeIndex(467, (d_2189_v5_).length(0))
        (d_2189_v5_)[index315_] = d_2185_v1_
        index316_ = default__.safeIndex(467, (d_2189_v5_).length(0))
        (d_2189_v5_)[index316_] = (len(p0)) < ((self).f6)
        d_2249_v45_: _dafny.Set
        d_2249_v45_ = _dafny.Set({(self).f6})
        d_2250_v46_: _dafny.MultiSet
        d_2250_v46_ = _dafny.MultiSet([len(d_2249_v45_), 951])
        d_2251_v47_: _dafny.MultiSet
        d_2251_v47_ = _dafny.MultiSet([d_2250_v46_])
        d_2252_v48_: _dafny.Array
        nw343_ = _dafny.Array(False, 23)
        d_2252_v48_ = nw343_
        d_2253_v50_: T2
        nw344_ = C2()
        def iife213_():
            coll61_ = _dafny.Map()
            compr_61_: int
            for compr_61_ in _dafny.IntegerRange(993, -151):
                d_2254_v49_: int = compr_61_
                if ((993) <= (d_2254_v49_)) and ((d_2254_v49_) < (-151)):
                    coll61_[(d_2254_v49_) - ((self).f6)] = p0
            return _dafny.Map(coll61_)
        nw344_.ctor__((self).f6, d_2251_v47_, d_2252_v48_, iife213_()
        )
        d_2253_v50_ = nw344_
        d_2255_v51_: _dafny.MultiSet
        d_2255_v51_ = _dafny.MultiSet([d_2253_v50_])
        d_2256_v52_: _dafny.Map
        d_2256_v52_ = _dafny.Map({(d_2253_v50_).f6: p0})
        d_2257_v53_: _dafny.Map
        d_2257_v53_ = d_2256_v52_
        d_2258_v54_: _dafny.MultiSet
        d_2258_v54_ = _dafny.MultiSet([d_2257_v53_])
        d_2259_v55_: _dafny.MultiSet
        d_2259_v55_ = _dafny.MultiSet([not(d_2185_v1_)])
        d_2260_v56_: _dafny.Array
        nw345_ = _dafny.Array(None, 17)
        nw345_[int(0)] = (self).f6
        nw345_[int(1)] = ((self).f6) + ((0) - ((self).f6))
        nw345_[int(2)] = (self).f6
        nw345_[int(3)] = (self).f6
        nw345_[int(4)] = ((self).f6) + ((self).f6)
        nw345_[int(5)] = (self).f6
        nw345_[int(6)] = (self).f6
        nw345_[int(7)] = (d_2255_v51_).cardinality
        nw345_[int(8)] = (d_2253_v50_).f6
        nw345_[int(9)] = ((d_2253_v50_).f6) - (default__.fm2((self).f6, globalState))
        nw345_[int(10)] = (self).f6
        nw345_[int(11)] = ((d_2253_v50_).f6) - ((d_2253_v50_).f6)
        nw345_[int(12)] = ((d_2253_v50_).f6) - ((self).f6)
        nw345_[int(13)] = (self).f6
        nw345_[int(14)] = (self).f6
        nw345_[int(15)] = (d_2253_v50_).f6
        nw345_[int(16)] = ((((d_2258_v54_).set(d_2257_v53_, default__.abs((d_2259_v55_).cardinality))).set(d_2257_v53_, default__.abs((self).f6))).intersection(d_2258_v54_)).cardinality
        d_2260_v56_ = nw345_
        index317_ = default__.safeIndex(866, (d_2260_v56_).length(0))
        (d_2260_v56_)[index317_] = (d_2253_v50_).f6
        index318_ = default__.safeIndex(866, (d_2260_v56_).length(0))
        (d_2260_v56_)[index318_] = default__.fm2((d_2253_v50_).f6, globalState)
        if default__.fm3(globalState):
            d_2261_v57_: _dafny.Seq
            d_2261_v57_ = _dafny.SeqWithoutIsStrInference([d_2253_v50_])
            d_2262_v58_: _dafny.Map
            d_2262_v58_ = _dafny.Map({(d_2186_v2_).set(default__.safeIndex((d_2259_v55_).cardinality, len(d_2186_v2_)), False): len(d_2261_v57_)})
            d_2262_v58_ = (d_2262_v58_).set(d_2186_v2_, ((d_2260_v56_)[default__.safeIndex(866, (d_2260_v56_).length(0))]) - (804))
            d_2263_v59_: bool
            out31_: bool
            out31_ = (self).m2(globalState)
            d_2263_v59_ = out31_
            d_2264_v60_: _dafny.MultiSet
            d_2264_v60_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([(self).f6]), d_2188_v4_])
            d_2265_v61_: _dafny.Map
            d_2265_v61_ = _dafny.Map({(d_2260_v56_)[default__.safeIndex(866, (d_2260_v56_).length(0))]: d_2264_v60_})
            def iife214_():
                coll62_ = _dafny.Set()
                compr_62_: _dafny.Seq
                for compr_62_ in (((d_2265_v61_)[(d_2253_v50_).f6] if ((d_2253_v50_).f6) in (d_2265_v61_) else d_2264_v60_)).Elements:
                    d_2266_v62_: _dafny.Seq = compr_62_
                    if (d_2266_v62_) in (((d_2265_v61_)[(d_2253_v50_).f6] if ((d_2253_v50_).f6) in (d_2265_v61_) else d_2264_v60_)):
                        coll62_ = coll62_.union(_dafny.Set([d_2266_v62_]))
                return _dafny.Set(coll62_)
            d_2184_v0_ = _dafny.SeqWithoutIsStrInference([((d_2253_v50_).f6) + (len(iife214_()
            ))])
            d_2267_v63_: _dafny.Array
            nw346_ = _dafny.Array(D11.default()(), 18)
            d_2267_v63_ = nw346_
            index319_ = default__.safeIndex(467, (d_2189_v5_).length(0))
            index320_ = default__.safeIndex(866, (d_2260_v56_).length(0))
            rhs385_ = d_2267_v63_
            rhs386_ = not(default__.fm3(globalState))
            rhs387_ = d_2185_v1_
            rhs388_ = default__.safeDivisionInt((d_2253_v50_).f6, len(p0))
            rhs389_ = d_2260_v56_
            lhs307_ = d_2189_v5_
            lhs308_ = default__.safeIndex(467, (d_2189_v5_).length(0))
            lhs309_ = d_2260_v56_
            lhs310_ = default__.safeIndex(866, (d_2260_v56_).length(0))
            d_2267_v63_ = rhs385_
            d_2263_v59_ = rhs386_
            lhs307_[lhs308_] = rhs387_
            lhs309_[lhs310_] = rhs388_
            d_2260_v56_ = rhs389_
            (globalState).f5 = (0) - ((((d_2253_v50_).f6) * ((d_2253_v50_).f6)) * ((d_2184_v0_)[default__.safeIndex((d_2260_v56_)[default__.safeIndex(866, (d_2260_v56_).length(0))], len(d_2184_v0_))]))
        elif True:
            index321_ = default__.safeIndex(866, (d_2260_v56_).length(0))
            (d_2260_v56_)[index321_] = (self).f6
            d_2268_v64_: C9
            nw347_ = C9()
            nw347_.ctor__(((d_2253_v50_).f6) + ((d_2260_v56_)[default__.safeIndex(866, (d_2260_v56_).length(0))]))
            d_2268_v64_ = nw347_
            index322_ = default__.safeIndex(866, (d_2260_v56_).length(0))
            rhs390_ = (d_2186_v2_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([(d_2253_v50_).f6, (0) - ((d_2253_v50_).f6)])), len(d_2186_v2_))]
            rhs391_ = d_2268_v64_
            rhs392_ = (0) - (default__.safeModuloInt((d_2253_v50_).f6, (d_2253_v50_).f6))
            rhs393_ = (0) - ((self).f6)
            lhs311_ = d_2260_v56_
            lhs312_ = default__.safeIndex(866, (d_2260_v56_).length(0))
            lhs313_ = globalState
            d_2185_v1_ = rhs390_
            d_2268_v64_ = rhs391_
            lhs311_[lhs312_] = rhs392_
            lhs313_.f5 = rhs393_
            d_2269_v65_: _dafny.Map
            d_2269_v65_ = _dafny.Map({(d_2253_v50_).f6: (d_2189_v5_)[default__.safeIndex(467, (d_2189_v5_).length(0))]})
            (globalState).f2 = (p0).set(default__.safeIndex(len(d_2269_v65_), len(p0)), d_2216_v28_)
            (globalState).f5 = (d_2253_v50_).f6
            d_2270_v66_: bool
            out32_: bool
            out32_ = (self).m2(globalState)
            d_2270_v66_ = out32_
        r0 = ((self).f6) == ((self).f6)
        return r0

    def m2(self, globalState):
        r0: bool = False
        d_2271_v0_: bool
        d_2271_v0_ = False
        (globalState).f2 = default__.fm18(d_2271_v0_, globalState)
        d_2272_v1_: _dafny.Array
        nw348_ = _dafny.Array(None, 5)
        nw348_[int(0)] = (False) or (d_2271_v0_)
        nw348_[int(1)] = d_2271_v0_
        nw348_[int(2)] = d_2271_v0_
        nw348_[int(3)] = (d_2271_v0_) or (d_2271_v0_)
        nw348_[int(4)] = d_2271_v0_
        d_2272_v1_ = nw348_
        index323_ = default__.safeIndex(145, (d_2272_v1_).length(0))
        (d_2272_v1_)[index323_] = d_2271_v0_
        d_2273_v2_: _dafny.Seq
        d_2273_v2_ = _dafny.SeqWithoutIsStrInference([d_2271_v0_])
        d_2274_v3_: _dafny.MultiSet
        d_2274_v3_ = _dafny.MultiSet([d_2273_v2_, d_2273_v2_])
        d_2275_v4_: _dafny.Seq
        d_2275_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rpavxdbs"))
        d_2276_v5_: _dafny.Map
        d_2276_v5_ = _dafny.Map({(self).f6: d_2275_v4_})
        d_2277_v6_: _dafny.Map
        d_2277_v6_ = _dafny.Map({(self).f6: (self).f6})
        d_2278_v7_: _dafny.MultiSet
        d_2278_v7_ = _dafny.MultiSet([264])
        d_2279_v8_: _dafny.MultiSet
        d_2279_v8_ = _dafny.MultiSet([(d_2278_v7_).cardinality])
        d_2280_v9_: _dafny.MultiSet
        d_2280_v9_ = _dafny.MultiSet([_dafny.MultiSet([(self).f6]), d_2279_v8_, d_2279_v8_, d_2279_v8_, d_2278_v7_])
        d_2281_v10_: C4
        nw349_ = C4()
        nw349_.ctor__(len(d_2275_v4_), (self).f6, d_2272_v1_, d_2276_v5_, ((d_2277_v6_)[len(_dafny.Set({(self).f6}))] if (len(_dafny.Set({(self).f6}))) in (d_2277_v6_) else (self).f6), d_2280_v9_)
        d_2281_v10_ = nw349_
        d_2282_v11_: D18
        d_2282_v11_ = D18_DC44(d_2271_v0_, d_2271_v0_, d_2274_v3_, d_2281_v10_)
        index324_ = default__.safeIndex(145, (d_2272_v1_).length(0))
        rhs394_ = (d_2282_v11_).cf80
        rhs395_ = (d_2281_v10_).f17
        rhs396_ = d_2271_v0_
        rhs397_ = d_2271_v0_
        rhs398_ = d_2272_v1_
        lhs314_ = globalState
        lhs315_ = d_2272_v1_
        lhs316_ = default__.safeIndex(145, (d_2272_v1_).length(0))
        lhs317_ = globalState
        r0 = rhs394_
        lhs314_.f5 = rhs395_
        lhs315_[lhs316_] = rhs396_
        lhs317_.f0 = rhs397_
        d_2272_v1_ = rhs398_
        if False:
            (globalState).f5 = (d_2281_v10_).fm21((d_2281_v10_).f17, globalState)
            d_2283_v12_: str
            d_2283_v12_ = _dafny.CodePoint('f')
            d_2284_v13_: _dafny.Set
            d_2284_v13_ = _dafny.Set({(d_2281_v10_).f17, (0) - ((d_2281_v10_).f18)})
            d_2285_v14_: D18
            d_2285_v14_ = D18_DC43(266)
            d_2286_v15_: _dafny.Seq
            d_2286_v15_ = _dafny.SeqWithoutIsStrInference([(d_2281_v10_).f17, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xsvnce")))])
            d_2287_v17_: _dafny.Set
            def iife215_():
                coll63_ = _dafny.Set()
                compr_63_: int
                for compr_63_ in (d_2286_v15_).Elements:
                    d_2288_v16_: int = compr_63_
                    if (d_2288_v16_) in (d_2286_v15_):
                        coll63_ = coll63_.union(_dafny.Set([(d_2288_v16_) * (2)]))
                return _dafny.Set(coll63_)
            d_2287_v17_ = _dafny.Set({(d_2283_v12_ if False else (d_2275_v4_)[default__.safeIndex(len(d_2284_v13_), len(d_2275_v4_))]), d_2283_v12_, default__.fm0(((d_2277_v6_)[(d_2285_v14_).cf78] if ((d_2285_v14_).cf78) in (d_2277_v6_) else len(d_2286_v15_)), iife215_()
            , (d_2272_v1_)[default__.safeIndex(145, (d_2272_v1_).length(0))], globalState), d_2283_v12_})
            d_2289_v18_: _dafny.Array
            nw350_ = _dafny.Array(int(0), 5)
            d_2289_v18_ = nw350_
            d_2290_v19_: _dafny.Map
            d_2290_v19_ = _dafny.Map({d_2289_v18_: d_2287_v17_})
            d_2291_v20_: _dafny.Seq
            d_2291_v20_ = _dafny.SeqWithoutIsStrInference([d_2273_v2_, d_2273_v2_, d_2273_v2_, d_2273_v2_])
            d_2292_v21_: D9
            d_2292_v21_ = D9_DC25(len((d_2291_v20_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q"))), len(d_2291_v20_))]), (d_2281_v10_).f18, (d_2281_v10_).f18, (d_2275_v4_)[default__.safeIndex((d_2281_v10_).f18, len(d_2275_v4_))])
            d_2293_v22_: _dafny.Seq
            d_2293_v22_ = _dafny.SeqWithoutIsStrInference([D9_DC25(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qqakq"))), (d_2281_v10_).f18, len(default__.fm9(globalState)), d_2283_v12_), d_2292_v21_])
            d_2287_v17_ = ((d_2290_v19_)[d_2289_v18_] if (d_2289_v18_) in (d_2290_v19_) else default__.fm54((d_2272_v1_)[default__.safeIndex(145, (d_2272_v1_).length(0))], default__.fm0((d_2281_v10_).f17, d_2284_v13_, (d_2272_v1_)[default__.safeIndex(145, (d_2272_v1_).length(0))], globalState), d_2293_v22_, globalState))
            d_2283_v12_ = d_2283_v12_
            d_2294_v23_: _dafny.Array
            nw351_ = _dafny.Array(_dafny.Array(None, 0), 7)
            d_2294_v23_ = nw351_
            d_2295_v24_: _dafny.Array
            nw352_ = _dafny.Array(None, 22)
            d_2295_v24_ = nw352_
            index325_ = default__.safeIndex(875, (d_2294_v23_).length(0))
            (d_2294_v23_)[index325_] = d_2295_v24_
            d_2296_v25_: T0
            nw353_ = C9()
            nw353_.ctor__((d_2281_v10_).f18)
            d_2296_v25_ = nw353_
            d_2297_v26_: _dafny.Map
            d_2297_v26_ = _dafny.Map({(0) - ((d_2296_v25_).f6): d_2296_v25_})
            d_2298_v27_: T0
            d_2298_v27_ = d_2296_v25_
            index326_ = default__.safeIndex(875, (d_2294_v23_).length(0))
            nw354_ = _dafny.Array(None, 20)
            nw354_[int(0)] = d_2296_v25_
            nw354_[int(1)] = d_2296_v25_
            nw354_[int(2)] = d_2296_v25_
            nw354_[int(3)] = d_2296_v25_
            nw354_[int(4)] = d_2296_v25_
            nw354_[int(5)] = d_2296_v25_
            nw354_[int(6)] = d_2296_v25_
            nw354_[int(7)] = d_2296_v25_
            nw354_[int(8)] = d_2296_v25_
            nw354_[int(9)] = d_2296_v25_
            nw354_[int(10)] = ((d_2297_v26_)[(d_2281_v10_).f17] if ((d_2281_v10_).f17) in (d_2297_v26_) else d_2296_v25_)
            nw354_[int(11)] = d_2296_v25_
            nw354_[int(12)] = d_2296_v25_
            nw354_[int(13)] = d_2296_v25_
            nw354_[int(14)] = d_2296_v25_
            nw354_[int(15)] = (d_2298_v27_)
            nw354_[int(16)] = d_2296_v25_
            nw354_[int(17)] = d_2296_v25_
            nw354_[int(18)] = d_2296_v25_
            nw354_[int(19)] = (d_2296_v25_)
            (d_2294_v23_)[index326_] = nw354_
            d_2299_v28_: _dafny.Map
            d_2299_v28_ = _dafny.Map({d_2271_v0_: (d_2281_v10_).f17})
            d_2300_v29_: _dafny.Map
            d_2300_v29_ = _dafny.Map({not((d_2272_v1_)[default__.safeIndex(145, (d_2272_v1_).length(0))]): len((d_2299_v28_) | (d_2299_v28_))})
            d_2299_v28_ = (d_2299_v28_).set((d_2275_v4_) < (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ywxdsc"))), len((d_2286_v15_) + (d_2286_v15_)))
        elif True:
            d_2301_v30_: D0
            d_2301_v30_ = D0_DC0(_dafny.SeqWithoutIsStrInference([d_2271_v0_, False, (d_2272_v1_)[default__.safeIndex(145, (d_2272_v1_).length(0))], (d_2272_v1_)[default__.safeIndex(145, (d_2272_v1_).length(0))]]), (d_2281_v10_).f18)
            d_2302_v31_: _dafny.Array
            def lambda94_(d_2303_i0_):
                return default__.safeDivisionInt(d_2303_i0_, (self).f6)

            init58_ = lambda94_
            nw355_ = _dafny.Array(None, 19)
            for i0_58_ in range(nw355_.length(0)):
                nw355_[i0_58_] = init58_(i0_58_)
            d_2302_v31_ = nw355_
            d_2304_v32_: _dafny.MultiSet
            d_2304_v32_ = _dafny.MultiSet([d_2302_v31_, d_2302_v31_])
            d_2305_v33_: _dafny.Map
            d_2305_v33_ = _dafny.Map({not((d_2272_v1_)[default__.safeIndex(145, (d_2272_v1_).length(0))]): d_2302_v31_})
            d_2306_v34_: _dafny.Map
            d_2306_v34_ = _dafny.Map({d_2301_v30_: (d_2304_v32_).set(((d_2305_v33_)[(d_2272_v1_)[default__.safeIndex(145, (d_2272_v1_).length(0))]] if ((d_2272_v1_)[default__.safeIndex(145, (d_2272_v1_).length(0))]) in (d_2305_v33_) else d_2302_v31_), default__.abs((self).f6))})
            d_2306_v34_ = (d_2306_v34_).set(d_2301_v30_, (d_2304_v32_) - (d_2304_v32_))
            d_2307_v35_: _dafny.Seq
            d_2307_v35_ = _dafny.SeqWithoutIsStrInference([(d_2281_v10_).f17, (d_2281_v10_).f18, -387, (d_2281_v10_).f18])
            d_2308_v36_: _dafny.Map
            d_2308_v36_ = _dafny.Map({(d_2307_v35_).set(default__.safeIndex(-334, len(d_2307_v35_)), (d_2281_v10_).f18): _dafny.SeqWithoutIsStrInference([default__.fm55(globalState)])})
            d_2308_v36_ = d_2308_v36_
            (globalState).f0 = ((d_2307_v35_) + (_dafny.SeqWithoutIsStrInference([(d_2281_v10_).f17 for d_2309_i1_ in range(default__.abs(249))]))) <= (d_2307_v35_)
            d_2310_v37_: _dafny.Map
            d_2310_v37_ = _dafny.Map({d_2271_v0_: 788})
            d_2311_v38_: _dafny.Seq
            d_2311_v38_ = _dafny.SeqWithoutIsStrInference([d_2310_v37_])
            (globalState).f5 = (len(((d_2311_v38_)[default__.safeIndex(len(d_2307_v35_), len(d_2311_v38_))]) | (d_2310_v37_))) * (((0) - ((d_2281_v10_).f17)) + ((d_2281_v10_).f18))
            d_2312_v39_: T2
            nw356_ = C4()
            nw356_.ctor__((d_2281_v10_).f18, 82, d_2272_v1_, (d_2276_v5_).set((d_2281_v10_).f18, d_2275_v4_), (self).f6, d_2280_v9_)
            d_2312_v39_ = nw356_
            d_2313_v40_: _dafny.Map
            d_2313_v40_ = _dafny.Map({(d_2281_v10_).f17: d_2312_v39_})
            d_2314_v41_: _dafny.Map
            d_2314_v41_ = _dafny.Map({d_2313_v40_: (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m")))) > ((d_2281_v10_).f18)})
            d_2314_v41_ = (d_2314_v41_).set(d_2313_v40_, d_2271_v0_)
        d_2315_v42_: _dafny.Array
        nw357_ = _dafny.Array(_dafny.CodePoint('D'), 28)
        d_2315_v42_ = nw357_
        d_2316_v43_: str
        d_2316_v43_ = _dafny.CodePoint('t')
        index327_ = default__.safeIndex(893, (d_2315_v42_).length(0))
        (d_2315_v42_)[index327_] = d_2316_v43_
        d_2317_v44_: D14
        d_2317_v44_ = D14_DC37((d_2272_v1_)[default__.safeIndex(145, (d_2272_v1_).length(0))], d_2275_v4_)
        pat_let_tv58_ = d_2281_v10_
        pat_let_tv59_ = d_2271_v0_
        pat_let_tv60_ = d_2281_v10_
        pat_let_tv61_ = d_2281_v10_
        pat_let_tv62_ = d_2316_v43_
        pat_let_tv63_ = d_2272_v1_
        pat_let_tv64_ = d_2272_v1_
        pat_let_tv65_ = d_2272_v1_
        pat_let_tv66_ = d_2272_v1_
        pat_let_tv67_ = d_2316_v43_
        pat_let_tv68_ = d_2316_v43_
        index328_ = default__.safeIndex(893, (d_2315_v42_).length(0))
        def lambda95_(source31_):
            if source31_.is_DC37:
                d_2318___mcc_h0_ = source31_.cf70
                d_2319___mcc_h1_ = source31_.cf71
                d_2320_cf71_ = d_2319___mcc_h1_
                d_2321_cf70_ = d_2318___mcc_h0_
                return (D9_DC25((pat_let_tv58_).f17, len(_dafny.Set({len(_dafny.Map({pat_let_tv59_: len(_dafny.SeqWithoutIsStrInference([(pat_let_tv60_).f17]))})), (pat_let_tv61_).f18})), -684, pat_let_tv62_)).cf49
            elif True:
                d_2322___mcc_h2_ = source31_.cf69
                d_2323_cf69_ = d_2322___mcc_h2_
                if (pat_let_tv64_)[default__.safeIndex(145, (pat_let_tv63_).length(0))]:
                    return (D0_DC1((pat_let_tv66_)[default__.safeIndex(145, (pat_let_tv65_).length(0))], (self).f6, pat_let_tv67_)).cf4
                elif True:
                    return pat_let_tv68_

        (d_2315_v42_)[index328_] = lambda95_(d_2317_v44_)
        d_2324_v45_: _dafny.Seq
        d_2324_v45_ = _dafny.SeqWithoutIsStrInference([701])
        rhs399_ = (d_2324_v45_)[default__.safeIndex((d_2281_v10_).f18, len(d_2324_v45_))]
        rhs400_ = (d_2324_v45_)[default__.safeIndex(((default__.fm56(d_2275_v4_, d_2271_v0_, d_2275_v4_, globalState)).cf59) - ((self).f6), len(d_2324_v45_))]
        rhs401_ = (d_2281_v10_).f17
        lhs318_ = globalState
        lhs319_ = globalState
        lhs320_ = globalState
        lhs318_.f5 = rhs399_
        lhs319_.f5 = rhs400_
        lhs320_.f5 = rhs401_
        d_2325_v46_: _dafny.Map
        d_2325_v46_ = _dafny.Map({False: d_2271_v0_})
        d_2326_v47_: _dafny.Set
        d_2326_v47_ = _dafny.Set({d_2325_v46_, (d_2325_v46_).set((d_2272_v1_)[default__.safeIndex(145, (d_2272_v1_).length(0))], True), d_2325_v46_})
        (globalState).f5 = len(d_2326_v47_)
        r0 = (d_2272_v1_)[default__.safeIndex(145, (d_2272_v1_).length(0))]
        return r0


class C11(T0):
    def  __init__(self):
        self._f6: int = int(0)
        self.f8: bool = False
        self._f7: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C11"
    @property
    def f6(self):
        return self._f6
    def ctor__(self, f7, f8, f6):
        (self)._f7 = f7
        (self).f8 = f8
        (self)._f6 = f6

    def fm5(self, p0, globalState):
        source32_ = (D0_DC2(D0_DC1(self.f8, (self).f6, _dafny.CodePoint('l'))) if self.f8 else D0_DC2(D0_DC0(_dafny.SeqWithoutIsStrInference([self.f8]), (self).f6)))
        if source32_.is_DC0:
            d_2327___mcc_h0_ = source32_.cf0
            d_2328___mcc_h1_ = source32_.cf1
            d_2329_cf1_ = d_2328___mcc_h1_
            d_2330_cf0_ = d_2327___mcc_h0_
            return _dafny.SeqWithoutIsStrInference([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v'), _dafny.CodePoint('y')]))])
        elif source32_.is_DC1:
            d_2331___mcc_h2_ = source32_.cf2
            d_2332___mcc_h3_ = source32_.cf3
            d_2333___mcc_h4_ = source32_.cf4
            d_2334_cf4_ = d_2333___mcc_h4_
            d_2335_cf3_ = d_2332___mcc_h3_
            d_2336_cf2_ = d_2331___mcc_h2_
            return (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_2334_cf4_])])) + (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_2334_cf4_, d_2334_cf4_])]))
        elif True:
            d_2337___mcc_h5_ = source32_.cf5
            d_2338_cf5_ = d_2337___mcc_h5_
            return (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f')])) for d_2339_i0_ in range(default__.abs(538))])) + (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([_dafny.CodePoint('g')]), (D1_DC3(_dafny.MultiSet([_dafny.CodePoint('j')]))).cf6, _dafny.MultiSet([_dafny.CodePoint('k')])]))

    def fm6(self, p0, globalState):
        return (_dafny.Map({(self).f6: _dafny.SeqWithoutIsStrInference([(self).f6])})) | (_dafny.Map({(self).f6: _dafny.SeqWithoutIsStrInference([len(_dafny.Set({self.f8, self.f8})), len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([self.f8]))})), (self).f6, (self).f6])}))

    def fm7(self, p0, p1, p2, p3, globalState):
        return self.f8

    def fm8(self, globalState):
        return _dafny.Map({(self).f6: (D0_DC0(_dafny.SeqWithoutIsStrInference([False]), (self).f6) if self.f8 else D0_DC0(_dafny.SeqWithoutIsStrInference([self.f8]), (0) - ((self).f6)))})

    def m0(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        d_2340_v0_: _dafny.Seq
        d_2340_v0_ = _dafny.SeqWithoutIsStrInference([(self).f6])
        d_2341_v1_: _dafny.Set
        d_2341_v1_ = _dafny.Set({self.f8, self.f8})
        d_2342_v2_: _dafny.Seq
        d_2342_v2_ = _dafny.SeqWithoutIsStrInference([p1, self.f8, False, self.f8])
        d_2343_v3_: _dafny.Seq
        d_2343_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tbon"))
        d_2344_v4_: D0
        d_2344_v4_ = D0_DC1(self.f8, (self).f6, _dafny.CodePoint('p'))
        d_2345_v5_: _dafny.MultiSet
        d_2345_v5_ = _dafny.MultiSet([(self).f6])
        d_2346_v6_: _dafny.MultiSet
        d_2346_v6_ = _dafny.MultiSet([d_2345_v5_])
        d_2347_v7_: _dafny.Map
        d_2347_v7_ = _dafny.Map({(self).f6: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ogtymm"))})
        d_2348_v8_: T0
        nw358_ = C2()
        nw358_.ctor__(len(d_2343_v3_), d_2346_v6_, p2, d_2347_v7_)
        d_2348_v8_ = nw358_
        d_2349_v9_: _dafny.Seq
        d_2349_v9_ = _dafny.SeqWithoutIsStrInference([d_2348_v8_])
        d_2350_v10_: D9
        d_2350_v10_ = D9_DC24(d_2346_v6_)
        d_2351_v11_: _dafny.Seq
        d_2351_v11_ = _dafny.SeqWithoutIsStrInference([d_2350_v10_, d_2350_v10_, d_2350_v10_])
        d_2352_v12_: _dafny.Array
        nw359_ = _dafny.Array(None, 20)
        nw359_[int(0)] = not (p1) or (self.f8)
        nw359_[int(1)] = self.f8
        nw359_[int(2)] = (365) >= ((self).f6)
        nw359_[int(3)] = self.f8
        nw359_[int(4)] = True
        nw359_[int(5)] = (d_2340_v0_) == (d_2340_v0_)
        nw359_[int(6)] = ((0) - (len(d_2341_v1_))) == ((0) - ((self).f6))
        nw359_[int(7)] = (p1) not in (d_2342_v2_)
        nw359_[int(8)] = p1
        nw359_[int(9)] = False
        nw359_[int(10)] = not((len((d_2343_v3_).set(default__.safeIndex((d_2344_v4_).cf3, len(d_2343_v3_)), p0))) <= ((self).f6))
        nw359_[int(11)] = p1
        nw359_[int(12)] = (d_2349_v9_) <= (d_2349_v9_)
        nw359_[int(13)] = p1
        nw359_[int(14)] = (d_2351_v11_) < (d_2351_v11_)
        nw359_[int(15)] = (d_2342_v2_)[default__.safeIndex((self).f6, len(d_2342_v2_))]
        nw359_[int(16)] = ((d_2348_v8_).f6) < (746)
        nw359_[int(17)] = not(self.f8)
        nw359_[int(18)] = p1
        nw359_[int(19)] = (len(d_2343_v3_)) >= ((d_2348_v8_).f6)
        d_2352_v12_ = nw359_
        d_2352_v12_ = d_2352_v12_
        (globalState).f5 = (d_2348_v8_).f6
        (globalState).f5 = (d_2348_v8_).f6
        d_2353_v13_: C4
        nw360_ = C4()
        nw360_.ctor__((d_2348_v8_).f6, (self).f6, p2, d_2347_v7_, (self).f6, d_2346_v6_)
        d_2353_v13_ = nw360_
        d_2354_v14_: _dafny.MultiSet
        d_2354_v14_ = _dafny.MultiSet([d_2353_v13_, d_2353_v13_])
        hi14_ = (0) - ((len(d_2341_v1_) if p1 else (d_2348_v8_).f6))
        for d_2355_i0_ in range(((d_2354_v14_)[d_2353_v13_] if (d_2353_v13_) in (d_2354_v14_) else (d_2353_v13_).f17), hi14_):
            d_2356_v15_: _dafny.Map
            d_2356_v15_ = _dafny.Map({((d_2353_v13_).f18) + ((d_2353_v13_).f17): (0) - ((d_2353_v13_).f18)})
            def iife216_():
                coll64_ = _dafny.Map()
                compr_64_: int
                for compr_64_ in _dafny.IntegerRange(-162, 475):
                    d_2357_v16_: int = compr_64_
                    if ((-162) <= (d_2357_v16_)) and ((d_2357_v16_) < (475)):
                        coll64_[(d_2357_v16_) + ((_dafny.MultiSet([(d_2353_v13_).f18, (d_2348_v8_).f6])).cardinality)] = (self).f6
                return _dafny.Map(coll64_)
            d_2356_v15_ = (iife216_()
            ) | (d_2356_v15_)
            if self.f8:
                d_2358_v17_: _dafny.Array
                def lambda96_(d_2359_v3_):
                    def lambda97_(d_2360_i1_):
                        return (d_2360_i1_) + (len(d_2359_v3_))

                    return lambda97_

                init59_ = lambda96_(d_2343_v3_)
                nw361_ = _dafny.Array(None, 29)
                for i0_59_ in range(nw361_.length(0)):
                    nw361_[i0_59_] = init59_(i0_59_)
                d_2358_v17_ = nw361_
                index329_ = default__.safeIndex(854, (d_2358_v17_).length(0))
                (d_2358_v17_)[index329_] = (291) * ((d_2353_v13_).f17)
                index330_ = default__.safeIndex(854, (d_2358_v17_).length(0))
                rhs402_ = len(((d_2343_v3_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_2361_i2_ in range(default__.abs(750))]))).set(default__.safeIndex((d_2348_v8_).f6, len((d_2343_v3_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_2362_i2_ in range(default__.abs(750))])))), p0))
                rhs403_ = self.f8
                rhs404_ = d_2343_v3_
                rhs405_ = p1
                lhs321_ = d_2358_v17_
                lhs322_ = default__.safeIndex(854, (d_2358_v17_).length(0))
                lhs323_ = globalState
                lhs324_ = globalState
                lhs321_[lhs322_] = rhs402_
                lhs323_.f0 = rhs403_
                d_2343_v3_ = rhs404_
                lhs324_.f0 = rhs405_
                d_2363_v18_: D14
                d_2363_v18_ = D14_DC37(self.f8, d_2343_v3_)
                (globalState).f0 = (d_2363_v18_).cf70
                (globalState).f5 = (d_2340_v0_)[default__.safeIndex((d_2353_v13_).f18, len(d_2340_v0_))]
                d_2364_v19_: _dafny.Map
                d_2364_v19_ = _dafny.Map({d_2348_v8_: d_2352_v12_})
                d_2352_v12_ = ((d_2364_v19_)[d_2348_v8_] if (d_2348_v8_) in (d_2364_v19_) else p2)
                d_2365_v20_: str
                d_2366_v21_: int
                d_2367_v22_: _dafny.Array
                d_2368_v23_: int
                out33_: str
                out34_: int
                out35_: _dafny.Array
                out36_: int
                out33_, out34_, out35_, out36_ = (d_2353_v13_).m7((d_2345_v5_) | (d_2345_v5_), (d_2353_v13_).f18, globalState)
                d_2365_v20_ = out33_
                d_2366_v21_ = out34_
                d_2367_v22_ = out35_
                d_2368_v23_ = out36_
            elif True:
                d_2369_v24_: _dafny.Array
                nw362_ = _dafny.Array(int(0), 20)
                d_2369_v24_ = nw362_
                index331_ = default__.safeIndex(590, (d_2369_v24_).length(0))
                (d_2369_v24_)[index331_] = (len(d_2342_v2_)) * ((d_2353_v13_).f17)
                index332_ = default__.safeIndex(590, (d_2369_v24_).length(0))
                rhs406_ = (d_2353_v13_).f18
                rhs407_ = (d_2353_v13_).f17
                lhs325_ = globalState
                lhs326_ = d_2369_v24_
                lhs327_ = default__.safeIndex(590, (d_2369_v24_).length(0))
                lhs325_.f5 = rhs406_
                lhs326_[lhs327_] = rhs407_
                rhs408_ = ((d_2345_v5_).set((d_2348_v8_).f6, default__.abs(default__.fm2((d_2348_v8_).f6, globalState)))) - (d_2345_v5_)
                rhs409_ = (d_2356_v15_) | (d_2356_v15_)
                rhs410_ = not(((d_2348_v8_).f6) >= (default__.safeModuloInt((d_2353_v13_).f18, 747)))
                rhs411_ = (d_2348_v8_).f6
                rhs412_ = d_2352_v12_
                lhs328_ = globalState
                lhs329_ = globalState
                d_2345_v5_ = rhs408_
                d_2356_v15_ = rhs409_
                lhs328_.f0 = rhs410_
                lhs329_.f5 = rhs411_
                d_2352_v12_ = rhs412_
                d_2370_v25_: _dafny.Array
                nw363_ = _dafny.Array(None, 23)
                nw363_[int(0)] = d_2348_v8_
                nw363_[int(1)] = d_2348_v8_
                nw363_[int(2)] = d_2348_v8_
                nw363_[int(3)] = d_2348_v8_
                nw363_[int(4)] = d_2348_v8_
                nw363_[int(5)] = d_2348_v8_
                nw363_[int(6)] = d_2348_v8_
                nw363_[int(7)] = d_2348_v8_
                nw363_[int(8)] = d_2348_v8_
                nw363_[int(9)] = d_2348_v8_
                nw363_[int(10)] = d_2348_v8_
                nw363_[int(11)] = d_2348_v8_
                nw363_[int(12)] = d_2348_v8_
                nw363_[int(13)] = d_2348_v8_
                nw363_[int(14)] = d_2348_v8_
                nw363_[int(15)] = d_2348_v8_
                nw363_[int(16)] = d_2348_v8_
                nw363_[int(17)] = d_2348_v8_
                nw363_[int(18)] = d_2348_v8_
                nw363_[int(19)] = d_2348_v8_
                nw363_[int(20)] = d_2348_v8_
                nw363_[int(21)] = d_2348_v8_
                nw363_[int(22)] = d_2348_v8_
                d_2370_v25_ = nw363_
                d_2371_v26_: _dafny.Map
                d_2371_v26_ = _dafny.Map({d_2370_v25_: d_2369_v24_})
                d_2371_v26_ = (d_2371_v26_).set(d_2370_v25_, d_2369_v24_)
                d_2372_v27_: D2
                d_2373_v28_: _dafny.Seq
                d_2374_v29_: _dafny.Seq
                out37_: D2
                out38_: _dafny.Seq
                out39_: _dafny.Seq
                out37_, out38_, out39_ = (d_2353_v13_).m13((d_2353_v13_).f18, globalState)
                d_2372_v27_ = out37_
                d_2373_v28_ = out38_
                d_2374_v29_ = out39_
                d_2375_v30_: _dafny.MultiSet
                d_2375_v30_ = _dafny.MultiSet([d_2342_v2_])
                d_2376_v31_: D18
                d_2376_v31_ = D18_DC44(True, p1, (d_2375_v30_).intersection(d_2375_v30_), d_2353_v13_)
                d_2376_v31_ = d_2376_v31_
            d_2377_v32_: C5
            nw364_ = C5()
            nw364_.ctor__((d_2353_v13_).f18, default__.safeModuloInt((self).f6, (d_2348_v8_).f6))
            d_2377_v32_ = nw364_
            d_2378_v33_: _dafny.Map
            d_2378_v33_ = _dafny.Map({(self).f6: _dafny.CodePoint('n')})
            rhs413_ = (d_2353_v13_).f17
            rhs414_ = len((_dafny.Map({len(((d_2343_v3_).set(default__.safeIndex((d_2377_v32_).f15, len(d_2343_v3_)), p0)).set(default__.safeIndex((d_2348_v8_).f6, len((d_2343_v3_).set(default__.safeIndex((d_2377_v32_).f15, len(d_2343_v3_)), p0))), p0)): p0}) if self.f8 else d_2378_v33_))
            rhs415_ = (d_2353_v13_).f18
            rhs416_ = d_2377_v32_
            lhs330_ = globalState
            lhs331_ = globalState
            lhs332_ = globalState
            lhs330_.f5 = rhs413_
            lhs331_.f5 = rhs414_
            lhs332_.f5 = rhs415_
            d_2377_v32_ = rhs416_
            (globalState).f0 = p1
        d_2379_v34_: _dafny.Array
        def lambda98_(d_2380_i3_):
            return (d_2380_i3_) * ((self).f6)

        init60_ = lambda98_
        nw365_ = _dafny.Array(None, 2)
        for i0_60_ in range(nw365_.length(0)):
            nw365_[i0_60_] = init60_(i0_60_)
        d_2379_v34_ = nw365_
        d_2381_v35_: C7
        nw366_ = C7()
        nw366_.ctor__(d_2379_v34_, d_2379_v34_, (d_2348_v8_).f6, d_2352_v12_, d_2347_v7_)
        d_2381_v35_ = nw366_
        d_2382_v36_: C2
        nw367_ = C2()
        nw367_.ctor__(len(d_2342_v2_), _dafny.MultiSet([_dafny.MultiSet([(d_2353_v13_).f17, (d_2348_v8_).f6]), d_2345_v5_]), p2, d_2347_v7_)
        d_2382_v36_ = nw367_
        d_2383_v37_: _dafny.Map
        d_2383_v37_ = _dafny.Map({d_2382_v36_: True})
        source33_ = _dafny.Map({d_2382_v36_: default__.fm3(globalState)})
        d_2384___mcc_h0_ = source33_
        d_2385_cf68_ = d_2384___mcc_h0_
        d_2386_v38_: str
        d_2386_v38_ = _dafny.CodePoint('a')
        rhs417_ = d_2342_v2_
        rhs418_ = p0
        d_2342_v2_ = rhs417_
        d_2386_v38_ = rhs418_
        d_2387_v39_: _dafny.Array
        nw368_ = _dafny.Array(_dafny.CodePoint('D'), 10)
        d_2387_v39_ = nw368_
        d_2388_v40_: _dafny.Array
        nw369_ = _dafny.Array(None, 27)
        nw369_[int(0)] = d_2387_v39_
        nw369_[int(1)] = d_2387_v39_
        nw369_[int(2)] = d_2387_v39_
        nw369_[int(3)] = d_2387_v39_
        nw369_[int(4)] = d_2387_v39_
        nw369_[int(5)] = d_2387_v39_
        nw369_[int(6)] = d_2387_v39_
        nw369_[int(7)] = d_2387_v39_
        nw369_[int(8)] = d_2387_v39_
        nw369_[int(9)] = d_2387_v39_
        nw369_[int(10)] = d_2387_v39_
        nw369_[int(11)] = d_2387_v39_
        nw369_[int(12)] = d_2387_v39_
        nw369_[int(13)] = d_2387_v39_
        nw369_[int(14)] = d_2387_v39_
        nw369_[int(15)] = d_2387_v39_
        nw369_[int(16)] = d_2387_v39_
        nw369_[int(17)] = d_2387_v39_
        nw369_[int(18)] = d_2387_v39_
        nw369_[int(19)] = d_2387_v39_
        nw369_[int(20)] = d_2387_v39_
        nw369_[int(21)] = d_2387_v39_
        nw369_[int(22)] = d_2387_v39_
        nw369_[int(23)] = d_2387_v39_
        nw369_[int(24)] = d_2387_v39_
        nw369_[int(25)] = d_2387_v39_
        nw369_[int(26)] = d_2387_v39_
        d_2388_v40_ = nw369_
        index333_ = default__.safeIndex(455, (d_2388_v40_).length(0))
        (d_2388_v40_)[index333_] = d_2387_v39_
        index334_ = default__.safeIndex(455, (d_2388_v40_).length(0))
        rhs419_ = d_2387_v39_
        rhs420_ = (default__.safeModuloInt(977, (d_2348_v8_).f6)) < ((d_2353_v13_).f18)
        lhs333_ = d_2388_v40_
        lhs334_ = default__.safeIndex(455, (d_2388_v40_).length(0))
        lhs335_ = self
        lhs333_[lhs334_] = rhs419_
        lhs335_.f8 = rhs420_
        (self).f8 = True
        d_2389_v41_: _dafny.Array
        nw370_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 24)
        d_2389_v41_ = nw370_
        index335_ = default__.safeIndex(429, (d_2389_v41_).length(0))
        (d_2389_v41_)[index335_] = d_2343_v3_
        index336_ = default__.safeIndex(429, (d_2389_v41_).length(0))
        (d_2389_v41_)[index336_] = (d_2343_v3_) + ((d_2343_v3_) + (d_2343_v3_))
        r0 = d_2343_v3_
        return r0

    @property
    def f7(self):
        return self._f7
