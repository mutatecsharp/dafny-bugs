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
        return -214

    @staticmethod
    def fm1(p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference([(_dafny.Map({not(False): (_dafny.MultiSet([420, len(_dafny.SeqWithoutIsStrInference([False, (D0_DC2(_dafny.MultiSet([_dafny.CodePoint('k'), _dafny.CodePoint('h')]), not(not(True)))).cf6])), 630, -203, (0) - ((_dafny.MultiSet([False, True])).cardinality)])).cardinality})) != (_dafny.Map({False: -421}))])

    @staticmethod
    def fm2(p0, globalState):
        return (-770) >= ((0) - ((0) - (default__.safeDivisionInt(-551, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "apvmft")))))))

    @staticmethod
    def fm4(p0, p1, p2, p3, globalState):
        if (False if True else False):
            return _dafny.SeqWithoutIsStrInference([811])
        elif True:
            return (_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Map({927: (0) - ((_dafny.MultiSet([True, False])).cardinality)}))) for d_0_i0_ in range(default__.abs(552))])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([-559, -695, 903])), (0) - (len(_dafny.SeqWithoutIsStrInference([False, False])))]))

    @staticmethod
    def fm6(p0, p1, globalState):
        source0_ = D10_DC29(False, _dafny.CodePoint('f'))
        if source0_.is_DC29:
            d_1___mcc_h0_ = source0_.cf42
            d_2___mcc_h1_ = source0_.cf43
            d_3_cf43_ = d_2___mcc_h1_
            d_4_cf42_ = d_1___mcc_h0_
            return _dafny.MultiSet((_dafny.SeqWithoutIsStrInference([-730])) + (_dafny.SeqWithoutIsStrInference([14, 18])))
        elif source0_.is_DC30:
            d_5___mcc_h2_ = source0_.cf44
            d_6___mcc_h3_ = source0_.cf45
            d_7___mcc_h4_ = source0_.cf46
            d_8___mcc_h5_ = source0_.cf47
            d_9___mcc_h6_ = source0_.cf48
            d_10_cf48_ = d_9___mcc_h6_
            d_11_cf47_ = d_8___mcc_h5_
            d_12_cf46_ = d_7___mcc_h4_
            d_13_cf45_ = d_6___mcc_h3_
            d_14_cf44_ = d_5___mcc_h2_
            if not(d_12_cf46_):
                return _dafny.MultiSet([d_10_cf48_, 212, d_11_cf47_])
            elif True:
                return _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([d_12_cf46_, d_12_cf46_])), 848]))
        elif True:
            d_15___mcc_h7_ = source0_.cf41
            d_16_cf41_ = d_15___mcc_h7_
            return _dafny.MultiSet([723])

    @staticmethod
    def fm7(p0, p1, p2, globalState):
        return ((_dafny.Set({False})) - (_dafny.Set({not(False)}))).intersection((_dafny.Set({True, True, True, True, False})) | (_dafny.Set({True})))

    @staticmethod
    def fm8(p0, p1, p2, p3, globalState):
        return _dafny.SeqWithoutIsStrInference([(_dafny.CodePoint('k') if True else _dafny.CodePoint('a')) for d_17_i0_ in range(default__.abs(468))])

    @staticmethod
    def fm9(p0, p1, globalState):
        return ((_dafny.Map({False: True})) | (_dafny.Map({False: False}))) | (_dafny.Map({not(True): True}))

    @staticmethod
    def fm10(p0, p1, p2, p3, globalState):
        source1_ = D13_DC38(_dafny.SeqWithoutIsStrInference([_dafny.Map({False: False})]))
        if source1_.is_DC39:
            d_18___mcc_h0_ = source1_.cf69
            d_19_cf69_ = d_18___mcc_h0_
            return (_dafny.SeqWithoutIsStrInference([-396])) + (_dafny.SeqWithoutIsStrInference([367, 184, len(_dafny.Map({838: (0) - (len(_dafny.SeqWithoutIsStrInference([-647 for d_20_i0_ in range(default__.abs(-492))])))}))]))
        elif source1_.is_DC40:
            d_21___mcc_h1_ = source1_.cf70
            d_22___mcc_h2_ = source1_.cf71
            d_23_cf71_ = d_22___mcc_h2_
            d_24_cf70_ = d_21___mcc_h1_
            return _dafny.SeqWithoutIsStrInference([d_23_cf71_])
        elif source1_.is_DC38:
            d_25___mcc_h3_ = source1_.cf68
            d_26_cf68_ = d_25___mcc_h3_
            return _dafny.SeqWithoutIsStrInference([len(_dafny.Set({True})) for d_27_i1_ in range(default__.abs(557))])
        elif True:
            d_28___mcc_h4_ = source1_.cf72
            d_29_cf72_ = d_28___mcc_h4_
            return (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({len(_dafny.Map({_dafny.CodePoint('g'): 603})): 497})) for d_30_i2_ in range(default__.abs(-849))])) + (_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.Map({_dafny.CodePoint('f'): False}), _dafny.Map({_dafny.CodePoint('u'): True})]))), 212]))

    @staticmethod
    def fm11(p0, globalState):
        if (331) == ((_dafny.MultiSet([(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ao"))))])).cardinality):
            return (_dafny.Map({len(_dafny.Map({not(not(False)): not(False)})): (0) - (-592)}))
        elif True:
            return _dafny.Map({len(_dafny.Map({104: len(_dafny.SeqWithoutIsStrInference([-845, len(_dafny.Map({not(False): False})), (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rdjmrl")))])).cardinality]))})): 371})

    @staticmethod
    def fm14(p0, p1, globalState):
        return ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False]))).intersection(_dafny.MultiSet([True, not(True)]))) - (_dafny.MultiSet([False]))

    @staticmethod
    def fm15(p0, p1, p2, globalState):
        return ((D16_DC46(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wrs")), _dafny.Set({False}), False) if not(False) else D16_DC46(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kbcxn")), _dafny.Set({False}), True))).cf79

    @staticmethod
    def fm16(p0, p1, globalState):
        return ((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([False, False, not(not(False))]), _dafny.SeqWithoutIsStrInference([True])]) if False else _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([True]), _dafny.SeqWithoutIsStrInference([True])]))) - ((_dafny.MultiSet([(D6_DC17(_dafny.SeqWithoutIsStrInference([True, True, False, not(False)]))).cf23, _dafny.SeqWithoutIsStrInference([False, False]), _dafny.SeqWithoutIsStrInference([False, True])])).intersection(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([False, False, False, False])])))

    @staticmethod
    def fm17(p0, globalState):
        return _dafny.Map({_dafny.CodePoint('r'): True})

    @staticmethod
    def fm18(p0, p1, p2, globalState):
        return (_dafny.Set({(D8_DC23(True, not(False), False, _dafny.CodePoint('x'))).cf35})) - (_dafny.Set({_dafny.CodePoint('s')}))

    @staticmethod
    def fm19(p0, globalState):
        return (_dafny.SeqWithoutIsStrInference([513, (0) - ((_dafny.MultiSet([False])).cardinality), 420, 698])) + ((_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Map({not(True): len((_dafny.Map({_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qa")))}): len(_dafny.Map({True: (0) - (-314)}))})))})))])) + (_dafny.SeqWithoutIsStrInference([48, 460])))

    @staticmethod
    def fm20(p0, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lt"))

    @staticmethod
    def fm21(p0, globalState):
        if True:
            return _dafny.Set({D0_DC3(), D0_DC3(), D0_DC3()})
        elif True:
            return _dafny.Set({D0_DC3()})

    @staticmethod
    def fm22(p0, p1, globalState):
        return ((_dafny.Set({True, True})) - (_dafny.Set({not(True)}))) | ((_dafny.Set({not(True)})) - (_dafny.Set({True, False, True})))

    @staticmethod
    def fm23(globalState):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: int
            for compr_0_ in _dafny.IntegerRange(134, 361):
                d_32_v0_: int = compr_0_
                if ((134) <= (d_32_v0_)) and ((d_32_v0_) < (361)):
                    coll0_[(d_32_v0_) - (len(_dafny.SeqWithoutIsStrInference([False])))] = (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vu"))))
            return _dafny.Map(coll0_)
        return ((_dafny.Map({210: 454})) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([D8_DC23(not(False), True, False, _dafny.CodePoint('i')) for d_31_i0_ in range(default__.abs(417))])): 794}))) | ((iife0_()
        ) | (_dafny.Map({666: -331})))

    @staticmethod
    def fm24(globalState):
        return _dafny.CodePoint('h')

    @staticmethod
    def fm25(p0, globalState):
        def iife1_():
            coll1_ = _dafny.Set()
            compr_1_: _dafny.Seq
            for compr_1_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([428]) for d_33_i0_ in range(default__.abs(249))])).Elements:
                d_34_v0_: _dafny.Seq = compr_1_
                if (d_34_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([428]) for d_33_i0_ in range(default__.abs(249))])):
                    coll1_ = coll1_.union(_dafny.Set([d_34_v0_]))
            return _dafny.Set(coll1_)
        return D0_DC0((iife1_()
).isdisjoint(_dafny.Set({_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({_dafny.CodePoint('i'): False})) for d_35_i1_ in range(default__.abs(-168))])), (0) - (len(_dafny.Set({False, (D16_DC46(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_36_i2_ in range(default__.abs(210))]), _dafny.Set({False}), False)).cf81})))])})))

    @staticmethod
    def fm26(globalState):
        return D8_DC24((D8_DC23(False, not(False), not(True), _dafny.CodePoint('f')) if True else D8_DC23(True, True, True, _dafny.CodePoint('t'))))

    @staticmethod
    def fm27(globalState):
        return (_dafny.Map({186: False})) | ((_dafny.Map({(_dafny.MultiSet([True])).cardinality: False})) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ogslplvme"))): True})))

    @staticmethod
    def fm28(p0, p1, p2, globalState):
        def iife2_():
            coll2_ = _dafny.Map()
            compr_2_: _dafny.Set
            for compr_2_ in (_dafny.Map({_dafny.Set({False, False, False}): True})).keys.Elements:
                d_37_v0_: _dafny.Set = compr_2_
                if (d_37_v0_) in (_dafny.Map({_dafny.Set({False, False, False}): True})):
                    coll2_[d_37_v0_] = len(_dafny.SeqWithoutIsStrInference([True, True, False]))
            return _dafny.Map(coll2_)
        if (len(iife2_()
        )) < ((D11_DC33(False, len(_dafny.Map({False: False})), -513)).cf54):
            def iife3_():
                coll3_ = _dafny.Map()
                compr_3_: _dafny.Seq
                for compr_3_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([True, True])))]): _dafny.CodePoint('m')})).keys.Elements:
                    d_38_v1_: _dafny.Seq = compr_3_
                    if (d_38_v1_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([True, True])))]): _dafny.CodePoint('m')})):
                        coll3_[d_38_v1_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mkjgpgh"))
                return _dafny.Map(coll3_)
            return (iife3_()
            ) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([(0) - (-199)]): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_39_i0_ in range(default__.abs(659))])}))
        elif True:
            def iife4_():
                coll4_ = _dafny.Map()
                compr_4_: _dafny.Seq
                for compr_4_ in (_dafny.Set({_dafny.SeqWithoutIsStrInference([-778, 723])})).Elements:
                    d_40_v2_: _dafny.Seq = compr_4_
                    if (d_40_v2_) in (_dafny.Set({_dafny.SeqWithoutIsStrInference([-778, 723])})):
                        coll4_[d_40_v2_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ivuqtjs"))
                return _dafny.Map(coll4_)
            def iife5_():
                coll5_ = _dafny.Map()
                compr_5_: _dafny.Seq
                for compr_5_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([196])])).Elements:
                    d_41_v3_: _dafny.Seq = compr_5_
                    if (d_41_v3_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([196])])):
                        coll5_[d_41_v3_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qpbh"))
                return _dafny.Map(coll5_)
            return (iife4_()
            ) | (iife5_()
            )

    @staticmethod
    def fm29(globalState):
        if (_dafny.Map({True: _dafny.CodePoint('g')})) == (_dafny.Map({True: _dafny.CodePoint('i')})):
            return D6_DC18(True, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a"))))
        elif True:
            return D6_DC18(not(True), 760)

    @staticmethod
    def fm30(p0, p1, p2, globalState):
        return D12_DC37(_dafny.Map({False: False}), (len(_dafny.Map({len(_dafny.Map({len(_dafny.Map({False: (_dafny.MultiSet([(D12_DC36(False, False, len(_dafny.SeqWithoutIsStrInference([486])))).cf61, not(True)])).cardinality})): False})): 15}))) != (len(_dafny.Map({D0_DC3(): len(_dafny.Set({True, not(False)}))}))), 713, 217)

    @staticmethod
    def fm31(globalState):
        source2_ = D12_DC36(False, False, (D12_DC37(_dafny.Map({not(False): True}), True, 527, 89)).cf67)
        if source2_.is_DC35:
            d_42___mcc_h0_ = source2_.cf56
            d_43___mcc_h1_ = source2_.cf57
            d_44___mcc_h2_ = source2_.cf58
            d_45___mcc_h3_ = source2_.cf59
            d_46___mcc_h4_ = source2_.cf60
            d_47_cf60_ = d_46___mcc_h4_
            d_48_cf59_ = d_45___mcc_h3_
            d_49_cf58_ = d_44___mcc_h2_
            d_50_cf57_ = d_43___mcc_h1_
            d_51_cf56_ = d_42___mcc_h0_
            return _dafny.Set({401})
        elif source2_.is_DC36:
            d_52___mcc_h5_ = source2_.cf61
            d_53___mcc_h6_ = source2_.cf62
            d_54___mcc_h7_ = source2_.cf63
            d_55_cf63_ = d_54___mcc_h7_
            d_56_cf62_ = d_53___mcc_h6_
            d_57_cf61_ = d_52___mcc_h5_
            return _dafny.Set({len(_dafny.Set({d_56_cf62_, d_57_cf61_}))})
        elif source2_.is_DC37:
            d_58___mcc_h8_ = source2_.cf64
            d_59___mcc_h9_ = source2_.cf65
            d_60___mcc_h10_ = source2_.cf66
            d_61___mcc_h11_ = source2_.cf67
            d_62_cf67_ = d_61___mcc_h11_
            d_63_cf66_ = d_60___mcc_h10_
            d_64_cf65_ = d_59___mcc_h9_
            d_65_cf64_ = d_58___mcc_h8_
            return _dafny.Set({d_62_cf67_})
        elif True:
            d_66___mcc_h12_ = source2_.cf55
            d_67_cf55_ = d_66___mcc_h12_
            def iife6_():
                coll6_ = _dafny.Set()
                compr_6_: int
                for compr_6_ in _dafny.IntegerRange(784, -950):
                    d_68_v0_: int = compr_6_
                    if ((784) <= (d_68_v0_)) and ((d_68_v0_) < (-950)):
                        coll6_ = coll6_.union(_dafny.Set([(d_68_v0_) - ((_dafny.MultiSet([_dafny.CodePoint('o'), _dafny.CodePoint('e'), _dafny.CodePoint('d'), _dafny.CodePoint('u'), _dafny.CodePoint('t')])).cardinality)]))
                return _dafny.Set(coll6_)
            return (_dafny.Set({(0) - (len(_dafny.SeqWithoutIsStrInference([True, (D16_DC46(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yoa")), _dafny.Set({False, True}), False)).cf81]))), 725})) | ((D4_DC11(iife6_()
)).cf18)

    @staticmethod
    def Main(noArgsParameter__):
        d_69_v0_: _dafny.Array
        def lambda0_(d_70_i0_):
            return _dafny.CodePoint('c')

        init0_ = lambda0_
        nw0_ = _dafny.Array(None, 26)
        for i0_0_ in range(nw0_.length(0)):
            nw0_[i0_0_] = init0_(i0_0_)
        d_69_v0_ = nw0_
        d_71_globalState_: GlobalState
        nw1_ = GlobalState()
        nw1_.ctor__(d_69_v0_, -562, False)
        d_71_globalState_ = nw1_
        d_72_v1_: bool
        d_72_v1_ = False
        if d_72_v1_:
            d_73_v2_: _dafny.Array
            def lambda1_(d_74_v1_):
                def lambda2_(d_75_i1_):
                    return d_74_v1_

                return lambda2_

            init1_ = lambda1_(d_72_v1_)
            nw2_ = _dafny.Array(None, 13)
            for i0_1_ in range(nw2_.length(0)):
                nw2_[i0_1_] = init1_(i0_1_)
            d_73_v2_ = nw2_
            d_76_v3_: int
            d_76_v3_ = -419
            d_77_v4_: _dafny.Map
            d_77_v4_ = _dafny.Map({d_76_v3_: d_73_v2_})
            d_73_v2_ = ((d_77_v4_)[d_76_v3_] if (d_76_v3_) in (d_77_v4_) else d_73_v2_)
            d_76_v3_ = 587
            (d_71_globalState_).f2 = d_72_v1_
            d_78_v5_: str
            d_78_v5_ = _dafny.CodePoint('v')
            d_78_v5_ = d_78_v5_
            d_79_v6_: _dafny.MultiSet
            d_79_v6_ = _dafny.MultiSet([d_78_v5_, d_78_v5_])
            d_80_v7_: _dafny.Seq
            d_80_v7_ = _dafny.SeqWithoutIsStrInference([43, 206, (d_76_v3_ if d_72_v1_ else 862), ((d_79_v6_).set(d_78_v5_, default__.abs(d_76_v3_))).cardinality, d_76_v3_])
            d_81_v8_: _dafny.Map
            d_81_v8_ = _dafny.Map({349: d_72_v1_})
            d_82_v9_: D0
            d_82_v9_ = D0_DC2(d_79_v6_, d_72_v1_)
            d_83_v10_: _dafny.Map
            d_83_v10_ = _dafny.Map({len(d_81_v8_): (d_82_v9_).cf6})
            d_84_v11_: _dafny.Map
            d_84_v11_ = _dafny.Map({d_72_v1_: d_81_v8_})
            d_80_v7_ = ((d_80_v7_) + (d_80_v7_)) + (_dafny.SeqWithoutIsStrInference([d_76_v3_, len(d_84_v11_), d_76_v3_, 140, -53]))
        elif True:
            d_85_v12_: _dafny.Array
            nw3_ = _dafny.Array(int(0), 6)
            d_85_v12_ = nw3_
            d_86_v13_: int
            d_86_v13_ = 744
            index0_ = default__.safeIndex(164, (d_85_v12_).length(0))
            (d_85_v12_)[index0_] = d_86_v13_
            index1_ = default__.safeIndex(164, (d_85_v12_).length(0))
            (d_85_v12_)[index1_] = default__.fm0(d_71_globalState_)
            d_87_v14_: _dafny.Map
            d_87_v14_ = _dafny.Map({d_86_v13_: default__.fm0(d_71_globalState_)})
            d_88_v15_: _dafny.Seq
            d_88_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))
            d_89_v16_: _dafny.Map
            d_89_v16_ = _dafny.Map({d_86_v13_: d_88_v15_})
            d_90_v17_: _dafny.Seq
            d_90_v17_ = _dafny.SeqWithoutIsStrInference([d_86_v13_])
            d_91_v18_: _dafny.Map
            d_91_v18_ = _dafny.Map({default__.safeDivisionInt((d_85_v12_)[default__.safeIndex(164, (d_85_v12_).length(0))], (d_85_v12_)[default__.safeIndex(164, (d_85_v12_).length(0))]): _dafny.Map({len(d_89_v16_): (d_90_v17_)[default__.safeIndex((d_85_v12_)[default__.safeIndex(164, (d_85_v12_).length(0))], len(d_90_v17_))]})})
            d_87_v14_ = ((d_91_v18_)[d_86_v13_] if (d_86_v13_) in (d_91_v18_) else (d_87_v14_).set(-505, d_86_v13_))
            d_92_v19_: _dafny.Seq
            d_92_v19_ = _dafny.SeqWithoutIsStrInference([d_85_v12_, d_85_v12_, d_85_v12_])
            d_93_v20_: _dafny.Seq
            d_93_v20_ = _dafny.SeqWithoutIsStrInference([not(d_72_v1_)])
            d_94_v21_: _dafny.Array
            nw4_ = _dafny.Array(None, 16)
            nw4_[int(0)] = d_85_v12_
            nw4_[int(1)] = d_85_v12_
            nw4_[int(2)] = (d_85_v12_ if d_72_v1_ else (d_92_v19_)[default__.safeIndex(d_86_v13_, len(d_92_v19_))])
            nw4_[int(3)] = d_85_v12_
            nw4_[int(4)] = d_85_v12_
            nw4_[int(5)] = d_85_v12_
            nw4_[int(6)] = d_85_v12_
            nw4_[int(7)] = d_85_v12_
            nw4_[int(8)] = d_85_v12_
            nw4_[int(9)] = d_85_v12_
            nw4_[int(10)] = d_85_v12_
            nw4_[int(11)] = (d_92_v19_)[default__.safeIndex(len(d_93_v20_), len(d_92_v19_))]
            nw4_[int(12)] = d_85_v12_
            nw4_[int(13)] = d_85_v12_
            nw4_[int(14)] = d_85_v12_
            nw4_[int(15)] = d_85_v12_
            d_94_v21_ = nw4_
            index2_ = default__.safeIndex(223, (d_94_v21_).length(0))
            (d_94_v21_)[index2_] = d_85_v12_
            d_95_v22_: _dafny.Map
            d_95_v22_ = _dafny.Map({(d_85_v12_)[default__.safeIndex(164, (d_85_v12_).length(0))]: False})
            d_96_v23_: str
            d_96_v23_ = _dafny.CodePoint('m')
            d_97_v24_: _dafny.Seq
            d_97_v24_ = _dafny.SeqWithoutIsStrInference([(d_88_v15_).set(default__.safeIndex(len(d_95_v22_), len(d_88_v15_)), d_96_v23_)])
            d_98_v25_: _dafny.Set
            d_98_v25_ = _dafny.Set({default__.fm0(d_71_globalState_), (d_85_v12_)[default__.safeIndex(164, (d_85_v12_).length(0))], d_86_v13_})
            d_99_v26_: _dafny.MultiSet
            d_99_v26_ = _dafny.MultiSet([d_72_v1_])
            d_100_v27_: _dafny.Map
            d_100_v27_ = _dafny.Map({(d_99_v26_).isdisjoint(_dafny.MultiSet(default__.fm1(d_96_v23_, d_72_v1_, d_71_globalState_))): not(default__.fm2(d_86_v13_, d_71_globalState_))})
            d_101_v28_: _dafny.Array
            def lambda3_(d_102_v1_):
                def lambda4_(d_103_i2_):
                    return d_102_v1_

                return lambda4_

            init2_ = lambda3_(d_72_v1_)
            nw5_ = _dafny.Array(None, 21)
            for i0_2_ in range(nw5_.length(0)):
                nw5_[i0_2_] = init2_(i0_2_)
            d_101_v28_ = nw5_
            index3_ = default__.safeIndex(223, (d_94_v21_).length(0))
            rhs0_ = ((0) - ((d_85_v12_)[default__.safeIndex(164, (d_85_v12_).length(0))])) in (d_98_v25_)
            rhs1_ = d_85_v12_
            rhs2_ = (d_97_v24_) + ((d_97_v24_) + (d_97_v24_))
            rhs3_ = ((d_100_v27_)[not((default__.fm0(d_71_globalState_)) != (len(_dafny.Map({(d_99_v26_).cardinality: d_101_v28_}))))] if (not((default__.fm0(d_71_globalState_)) != (len(_dafny.Map({(d_99_v26_).cardinality: d_101_v28_}))))) in (d_100_v27_) else d_72_v1_)
            rhs4_ = d_88_v15_
            lhs0_ = d_94_v21_
            lhs1_ = default__.safeIndex(223, (d_94_v21_).length(0))
            d_72_v1_ = rhs0_
            lhs0_[lhs1_] = rhs1_
            d_97_v24_ = rhs2_
            d_72_v1_ = rhs3_
            d_88_v15_ = rhs4_
            d_104_v29_: C1
            nw6_ = C1()
            nw6_.ctor__(d_86_v13_, default__.safeDivisionInt((d_85_v12_)[default__.safeIndex(164, (d_85_v12_).length(0))], d_86_v13_), default__.fm2(132, d_71_globalState_))
            d_104_v29_ = nw6_
            d_105_v30_: int
            out0_: int
            out0_ = (d_104_v29_).m1((d_85_v12_)[default__.safeIndex(164, (d_85_v12_).length(0))], d_71_globalState_)
            d_105_v30_ = out0_
        d_106_v31_: _dafny.Array
        nw7_ = _dafny.Array(None, 11)
        nw7_[int(0)] = d_72_v1_
        nw7_[int(1)] = d_72_v1_
        nw7_[int(2)] = d_72_v1_
        nw7_[int(3)] = d_72_v1_
        nw7_[int(4)] = d_72_v1_
        nw7_[int(5)] = d_72_v1_
        nw7_[int(6)] = d_72_v1_
        nw7_[int(7)] = d_72_v1_
        nw7_[int(8)] = not(d_72_v1_)
        nw7_[int(9)] = d_72_v1_
        nw7_[int(10)] = d_72_v1_
        d_106_v31_ = nw7_
        d_107_v32_: int
        d_107_v32_ = 714
        d_108_v33_: _dafny.Set
        d_108_v33_ = _dafny.Set({d_72_v1_})
        d_109_v34_: C5
        nw8_ = C5()
        nw8_.ctor__(d_106_v31_, ((0) - (d_107_v32_)) - (len(d_108_v33_)), (d_107_v32_) - (default__.fm0(d_71_globalState_)), (d_108_v33_).ispropersubset(d_108_v33_))
        d_109_v34_ = nw8_
        d_110_v35_: str
        d_110_v35_ = _dafny.CodePoint('s')
        d_111_v36_: _dafny.Seq
        d_111_v36_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ndf"))
        d_112_v37_: _dafny.Array
        nw9_ = _dafny.Array(int(0), 1)
        d_112_v37_ = nw9_
        d_113_v38_: C3
        nw10_ = C3()
        nw10_.ctor__(d_110_v35_, d_111_v36_, d_107_v32_, d_72_v1_, (D10_DC30(len(default__.fm1(d_110_v35_, d_72_v1_, d_71_globalState_)), d_112_v37_, d_72_v1_, d_107_v32_, d_107_v32_)).cf44)
        d_113_v38_ = nw10_
        d_114_v39_: _dafny.Map
        d_114_v39_ = _dafny.Map({d_113_v38_: d_107_v32_})
        d_115_v40_: _dafny.Seq
        d_115_v40_ = _dafny.SeqWithoutIsStrInference([d_72_v1_, d_72_v1_, d_72_v1_, d_72_v1_])
        d_116_v41_: _dafny.Seq
        d_116_v41_ = _dafny.SeqWithoutIsStrInference([(d_114_v39_).set(d_113_v38_, (_dafny.MultiSet(d_115_v40_)).cardinality)])
        hi0_ = (d_107_v32_ if d_72_v1_ else d_107_v32_)
        for d_117_i3_ in range(len(d_116_v41_), hi0_):
            d_118_v42_: _dafny.Array
            nw11_ = _dafny.Array(_dafny.Map({}), 27)
            d_118_v42_ = nw11_
            d_119_v43_: _dafny.Array
            nw12_ = _dafny.Array(None, 25)
            d_119_v43_ = nw12_
            d_120_v44_: _dafny.Map
            d_120_v44_ = _dafny.Map({d_107_v32_: d_119_v43_})
            index4_ = default__.safeIndex(952, (d_118_v42_).length(0))
            (d_118_v42_)[index4_] = (d_120_v44_) | (d_120_v44_)
            d_121_v45_: T1
            nw13_ = C4()
            nw13_.ctor__(472, len(default__.fm31(d_71_globalState_)), d_72_v1_)
            d_121_v45_ = nw13_
            d_122_v46_: _dafny.Map
            d_122_v46_ = _dafny.Map({default__.safeModuloInt(d_117_i3_, len(_dafny.SeqWithoutIsStrInference([d_121_v45_, d_121_v45_]))): d_120_v44_})
            index5_ = default__.safeIndex(952, (d_118_v42_).length(0))
            (d_118_v42_)[index5_] = ((d_122_v46_)[(0) - ((d_107_v32_) + (d_117_i3_))] if ((0) - ((d_107_v32_) + (d_117_i3_))) in (d_122_v46_) else d_120_v44_)
            d_123_v47_: _dafny.Map
            d_123_v47_ = _dafny.Map({d_111_v36_: d_113_v38_.f13})
            d_124_v48_: _dafny.Array
            nw14_ = _dafny.Array(None, 29)
            nw14_[int(0)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wiusmh"))
            nw14_[int(1)] = default__.fm20((d_121_v45_).f3, d_71_globalState_)
            nw14_[int(2)] = (_dafny.SeqWithoutIsStrInference([(d_113_v38_).f12 for d_125_i4_ in range(default__.abs(584))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tr")))
            nw14_[int(3)] = d_113_v38_.f13
            nw14_[int(4)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bybpuk"))
            nw14_[int(5)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qkhmicrt"))
            nw14_[int(6)] = d_111_v36_
            nw14_[int(7)] = _dafny.SeqWithoutIsStrInference([d_110_v35_ for d_126_i5_ in range(default__.abs(-349))])
            nw14_[int(8)] = (d_111_v36_) + ((_dafny.SeqWithoutIsStrInference([(d_113_v38_).f12 for d_127_i6_ in range(default__.abs(293))])).set(default__.safeIndex(53, len(_dafny.SeqWithoutIsStrInference([(d_113_v38_).f12 for d_128_i6_ in range(default__.abs(293))]))), (D8_DC23(d_121_v45_.f4, d_72_v1_, True, (d_113_v38_).f12)).cf35))
            nw14_[int(9)] = (default__.fm20(-245, d_71_globalState_)) + (d_111_v36_)
            nw14_[int(10)] = d_113_v38_.f13
            nw14_[int(11)] = (d_113_v38_.f13) + (d_111_v36_)
            nw14_[int(12)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "crcifcx"))
            nw14_[int(13)] = d_113_v38_.f13
            nw14_[int(14)] = d_113_v38_.f13
            nw14_[int(15)] = d_113_v38_.f13
            nw14_[int(16)] = (d_111_v36_) + (d_113_v38_.f13)
            nw14_[int(17)] = default__.fm20((d_121_v45_).f7, d_71_globalState_)
            nw14_[int(18)] = (d_113_v38_.f13) + (d_111_v36_)
            nw14_[int(19)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jluu"))
            nw14_[int(20)] = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wagffgkjj"))) + (d_113_v38_.f13)).set(default__.safeIndex(d_107_v32_, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wagffgkjj"))) + (d_113_v38_.f13))), (d_113_v38_).f12)
            nw14_[int(21)] = d_111_v36_
            nw14_[int(22)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "agrk"))
            nw14_[int(23)] = (d_113_v38_.f13 if d_72_v1_ else ((d_123_v47_)[_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x"))] if (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x"))) in (d_123_v47_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k"))))
            nw14_[int(24)] = _dafny.SeqWithoutIsStrInference([d_110_v35_ for d_129_i7_ in range(default__.abs(-793))])
            nw14_[int(25)] = d_113_v38_.f13
            nw14_[int(26)] = d_113_v38_.f13
            nw14_[int(27)] = (d_111_v36_) + (d_113_v38_.f13)
            nw14_[int(28)] = (_dafny.SeqWithoutIsStrInference([d_110_v35_ for d_130_i8_ in range(default__.abs(854))]) if d_121_v45_.f4 else d_111_v36_)
            d_124_v48_ = nw14_
            d_124_v48_ = d_124_v48_
            if not (d_121_v45_.f4) or (not (True) or (d_72_v1_)):
                d_131_v49_: _dafny.Map
                d_131_v49_ = _dafny.Map({((d_121_v45_).f3) + ((d_121_v45_).f3): d_110_v35_})
                d_131_v49_ = (d_131_v49_).set((d_113_v38_).fm12(not(d_121_v45_.f4), d_121_v45_.f4, d_71_globalState_), d_110_v35_)
                (d_113_v38_).m5(d_71_globalState_)
                (d_71_globalState_).f1 = d_117_i3_
                d_132_v50_: _dafny.Array
                nw15_ = _dafny.Array(None, 16)
                d_132_v50_ = nw15_
                index6_ = default__.safeIndex(202, (d_132_v50_).length(0))
                (d_132_v50_)[index6_] = d_121_v45_
                index7_ = default__.safeIndex(202, (d_132_v50_).length(0))
                nw16_ = C5()
                nw16_.ctor__((d_109_v34_).f9, (d_121_v45_).f3, -406, d_121_v45_.f4)
                (d_132_v50_)[index7_] = nw16_
                d_133_v51_: C7
                nw17_ = C7()
                nw17_.ctor__(d_121_v45_.f4, (d_121_v45_).f7, (d_121_v45_).f7, (d_108_v33_).ispropersubset(d_108_v33_))
                d_133_v51_ = nw17_
            elif True:
                (d_121_v45_).f4 = d_121_v45_.f4
                d_134_v52_: C7
                nw18_ = C7()
                nw18_.ctor__(d_121_v45_.f4, (d_121_v45_).f7, len(d_111_v36_), d_72_v1_)
                d_134_v52_ = nw18_
                d_135_v53_: _dafny.Seq
                d_135_v53_ = _dafny.SeqWithoutIsStrInference([d_134_v52_, d_134_v52_, d_134_v52_])
                index8_ = default__.safeIndex(608, (d_112_v37_).length(0))
                (d_112_v37_)[index8_] = len(d_135_v53_)
                index9_ = default__.safeIndex(608, (d_112_v37_).length(0))
                (d_112_v37_)[index9_] = (d_134_v52_).f6
                d_136_v54_: _dafny.Set
                d_136_v54_ = _dafny.Set({d_109_v34_, d_109_v34_})
                d_136_v54_ = d_136_v54_
                d_137_v55_: _dafny.Array
                nw19_ = _dafny.Array(int(0), 1)
                d_137_v55_ = nw19_
                d_138_v56_: T0
                nw20_ = C2()
                nw20_.ctor__((d_121_v45_).f7, d_137_v55_, d_107_v32_, d_72_v1_)
                d_138_v56_ = nw20_
                d_139_v57_: _dafny.Set
                d_139_v57_ = _dafny.Set({d_138_v56_})
                d_140_v58_: _dafny.Map
                d_140_v58_ = _dafny.Map({d_121_v45_.f4: len(d_139_v57_)})
                d_141_v59_: _dafny.Map
                d_141_v59_ = _dafny.Map({len((d_140_v58_) | (d_140_v58_)): d_121_v45_.f4})
                rhs5_ = not(((d_141_v59_)[(0) - (d_117_i3_)] if ((0) - (d_117_i3_)) in (d_141_v59_) else d_121_v45_.f4))
                rhs6_ = d_121_v45_.f4
                rhs7_ = (d_121_v45_).f7
                lhs2_ = d_134_v52_
                lhs3_ = d_71_globalState_
                d_72_v1_ = rhs5_
                lhs2_.f5 = rhs6_
                lhs3_.f1 = rhs7_
                d_137_v55_ = d_137_v55_
            d_142_v60_: _dafny.Map
            d_142_v60_ = _dafny.Map({d_121_v45_.f4: d_111_v36_})
            (d_121_v45_).f4 = ((d_121_v45_).f3) >= (((d_113_v38_).fm12(default__.fm2(len(((d_142_v60_)[(d_113_v38_).fm13(d_71_globalState_)] if ((d_113_v38_).fm13(d_71_globalState_)) in (d_142_v60_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "epeew")))), d_71_globalState_), d_72_v1_, d_71_globalState_)) - ((d_121_v45_).f7))
        hi1_ = d_107_v32_
        for d_143_i9_ in range(d_107_v32_, hi1_):
            d_144_v61_: D8
            d_144_v61_ = D8_DC23(d_72_v1_, not(d_72_v1_), False, _dafny.CodePoint('e'))
            d_145_v62_: D8
            d_145_v62_ = D8_DC24(d_144_v61_)
            d_146_v63_: D8
            d_146_v63_ = D8_DC24(d_144_v61_)
            d_147_v64_: D13
            d_147_v64_ = D13_DC40(d_146_v63_, (0) - (d_107_v32_))
            d_148_v65_: D12
            d_148_v65_ = D12_DC36(d_72_v1_, not(d_72_v1_), default__.fm0(d_71_globalState_))
            d_149_v66_: D12
            d_149_v66_ = D12_DC36(d_72_v1_, d_72_v1_, (d_148_v65_).cf63)
            rhs8_ = (d_107_v32_) >= ((d_148_v65_).cf63)
            rhs9_ = d_72_v1_
            rhs10_ = d_147_v64_
            lhs4_ = d_71_globalState_
            d_72_v1_ = rhs8_
            lhs4_.f2 = rhs9_
            d_147_v64_ = rhs10_
            d_150_v67_: _dafny.MultiSet
            d_150_v67_ = _dafny.MultiSet([(d_113_v38_).f12])
            d_151_v68_: C4
            nw21_ = C4()
            nw21_.ctor__(d_107_v32_, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_72_v1_]))).cardinality, (D0_DC2(d_150_v67_, not(d_72_v1_))).cf6)
            d_151_v68_ = nw21_
            rhs11_ = d_113_v38_
            rhs12_ = d_151_v68_
            d_113_v38_ = rhs11_
            d_151_v68_ = rhs12_
            d_107_v32_ = d_107_v32_
            nw22_ = _dafny.Array(False, 10)
            d_106_v31_ = nw22_
        d_152_v69_: C1
        nw23_ = C1()
        nw23_.ctor__(len(d_111_v36_), 840, (d_115_v40_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p"))), len(d_115_v40_))])
        d_152_v69_ = nw23_
        d_153_v70_: _dafny.Seq
        d_153_v70_ = _dafny.SeqWithoutIsStrInference([d_152_v69_])
        d_154_v71_: _dafny.Set
        d_154_v71_ = _dafny.Set({d_152_v69_, d_152_v69_, d_152_v69_, (d_153_v70_)[default__.safeIndex(d_107_v32_, len(d_153_v70_))]})
        d_154_v71_ = (d_154_v71_).intersection(d_154_v71_)
        hi2_ = (273) - (d_107_v32_)
        for d_155_i10_ in range(d_107_v32_, hi2_):
            d_156_v72_: _dafny.MultiSet
            d_156_v72_ = _dafny.MultiSet([d_72_v1_, d_72_v1_])
            d_157_v73_: _dafny.MultiSet
            d_157_v73_ = _dafny.MultiSet([d_155_i10_])
            (d_113_v38_).f13 = default__.fm15((d_72_v1_) or (d_72_v1_), default__.safeModuloInt(((d_156_v72_)[not(d_72_v1_)] if (not(d_72_v1_)) in (d_156_v72_) else -565), d_155_i10_), ((d_157_v73_)[d_155_i10_] if (d_155_i10_) in (d_157_v73_) else d_107_v32_), d_71_globalState_)
            index10_ = default__.safeIndex(840, (d_106_v31_).length(0))
            (d_106_v31_)[index10_] = d_72_v1_
            index11_ = default__.safeIndex(840, (d_106_v31_).length(0))
            (d_106_v31_)[index11_] = (len(_dafny.Map({d_72_v1_: d_155_i10_}))) < (-570)
            d_158_v74_: _dafny.Set
            d_158_v74_ = _dafny.Set({d_113_v38_, d_113_v38_, d_113_v38_, d_113_v38_, d_113_v38_})
            d_158_v74_ = d_158_v74_
            d_159_v75_: _dafny.Array
            def lambda5_(d_160_v31_):
                def lambda6_(d_161_i11_):
                    return not ((d_160_v31_)[default__.safeIndex(840, (d_160_v31_).length(0))]) or ((d_160_v31_)[default__.safeIndex(840, (d_160_v31_).length(0))])

                return lambda6_

            init3_ = lambda5_(d_106_v31_)
            nw24_ = _dafny.Array(None, 1)
            for i0_3_ in range(nw24_.length(0)):
                nw24_[i0_3_] = init3_(i0_3_)
            d_159_v75_ = nw24_
            index12_ = default__.safeIndex(947, (d_69_v0_).length(0))
            (d_69_v0_)[index12_] = (d_113_v38_).f12
            index13_ = default__.safeIndex(947, (d_69_v0_).length(0))
            rhs13_ = d_159_v75_
            rhs14_ = default__.safeModuloInt((d_107_v32_) - (d_107_v32_), d_107_v32_)
            rhs15_ = d_155_i10_
            rhs16_ = (d_113_v38_).f12
            rhs17_ = d_72_v1_
            lhs5_ = d_71_globalState_
            lhs6_ = d_71_globalState_
            lhs7_ = d_69_v0_
            lhs8_ = default__.safeIndex(947, (d_69_v0_).length(0))
            d_159_v75_ = rhs13_
            lhs5_.f1 = rhs14_
            lhs6_.f1 = rhs15_
            lhs7_[lhs8_] = rhs16_
            d_72_v1_ = rhs17_
        d_72_v1_ = d_72_v1_
        d_162_v76_: _dafny.Set
        d_162_v76_ = _dafny.Set({d_107_v32_})
        d_163_i12_: int
        d_163_i12_ = 0
        with _dafny.label("0"):
            while not((d_107_v32_) != (default__.safeModuloInt(len(d_162_v76_), (0) - (d_107_v32_)))):
                with _dafny.c_label("0"):
                    if (d_163_i12_) >= (100):
                        raise _dafny.Break("0")
                    d_163_i12_ = (d_163_i12_) + (1)
                    d_107_v32_ = d_107_v32_
                    index14_ = default__.safeIndex(699, (d_112_v37_).length(0))
                    (d_112_v37_)[index14_] = default__.safeDivisionInt(d_107_v32_, default__.fm0(d_71_globalState_))
                    d_164_v77_: _dafny.Array
                    nw25_ = _dafny.Array(_dafny.Array(None, 0), 17)
                    d_164_v77_ = nw25_
                    d_165_v78_: _dafny.Array
                    nw26_ = _dafny.Array(_dafny.Map({}), 18)
                    d_165_v78_ = nw26_
                    d_166_v79_: _dafny.Set
                    d_166_v79_ = _dafny.Set({(d_113_v38_).f12, _dafny.CodePoint('k')})
                    d_167_v80_: T1
                    nw27_ = C5()
                    nw27_.ctor__((d_109_v34_).f9, len(d_113_v38_.f13), len(d_166_v79_), d_72_v1_)
                    d_167_v80_ = nw27_
                    d_168_v81_: _dafny.Set
                    d_168_v81_ = _dafny.Set({d_167_v80_})
                    d_169_v82_: _dafny.Map
                    d_169_v82_ = _dafny.Map({d_72_v1_: d_168_v81_})
                    index15_ = default__.safeIndex(388, (d_165_v78_).length(0))
                    (d_165_v78_)[index15_] = d_169_v82_
                    index16_ = default__.safeIndex(699, (d_112_v37_).length(0))
                    index17_ = default__.safeIndex(388, (d_165_v78_).length(0))
                    rhs18_ = len(d_162_v76_)
                    rhs19_ = d_164_v77_
                    rhs20_ = (((d_169_v82_).set(d_167_v80_.f4, d_168_v81_)) | (d_169_v82_)).set(d_72_v1_, d_168_v81_)
                    lhs9_ = d_112_v37_
                    lhs10_ = default__.safeIndex(699, (d_112_v37_).length(0))
                    lhs11_ = d_165_v78_
                    lhs12_ = default__.safeIndex(388, (d_165_v78_).length(0))
                    lhs9_[lhs10_] = rhs18_
                    d_164_v77_ = rhs19_
                    lhs11_[lhs12_] = rhs20_
                    if d_167_v80_.f4:
                        d_170_v83_: _dafny.Array
                        def lambda7_(d_171_v80_):
                            def lambda8_(d_172_i13_):
                                return (d_172_i13_) * ((d_171_v80_).f7)

                            return lambda8_

                        init4_ = lambda7_(d_167_v80_)
                        nw28_ = _dafny.Array(None, 9)
                        for i0_4_ in range(nw28_.length(0)):
                            nw28_[i0_4_] = init4_(i0_4_)
                        d_170_v83_ = nw28_
                        d_173_v84_: _dafny.Array
                        nw29_ = _dafny.Array(None, 29)
                        d_173_v84_ = nw29_
                        d_174_v85_: _dafny.Map
                        d_174_v85_ = _dafny.Map({d_173_v84_: d_113_v38_})
                        rhs21_ = d_170_v83_
                        rhs22_ = ((d_174_v85_)[d_173_v84_] if (d_173_v84_) in (d_174_v85_) else d_113_v38_)
                        d_170_v83_ = rhs21_
                        d_113_v38_ = rhs22_
                        d_175_v86_: C2
                        nw30_ = C2()
                        nw30_.ctor__(d_107_v32_, d_170_v83_, (d_167_v80_).f7, d_167_v80_.f4)
                        d_175_v86_ = nw30_
                        d_176_v87_: _dafny.Map
                        d_176_v87_ = _dafny.Map({d_167_v80_.f4: d_167_v80_.f4})
                        d_177_v88_: D12
                        d_177_v88_ = D12_DC37(d_176_v87_, d_72_v1_, (d_112_v37_)[default__.safeIndex(699, (d_112_v37_).length(0))], (d_167_v80_).f7)
                        d_178_v89_: _dafny.Seq
                        d_178_v89_ = _dafny.SeqWithoutIsStrInference([(d_167_v80_).f7, default__.safeModuloInt((d_167_v80_).f7, (d_167_v80_).f7), (d_177_v88_).cf66])
                        d_179_v90_: _dafny.Map
                        d_179_v90_ = _dafny.Map({(d_175_v86_).f14: _dafny.SeqWithoutIsStrInference([d_167_v80_.f4, d_167_v80_.f4])})
                        d_180_v91_: _dafny.MultiSet
                        d_180_v91_ = _dafny.MultiSet([d_113_v38_.f13, d_111_v36_])
                        d_181_v92_: _dafny.MultiSet
                        d_181_v92_ = _dafny.MultiSet([(d_167_v80_).f3])
                        index18_ = default__.safeIndex(699, (d_112_v37_).length(0))
                        rhs23_ = d_175_v86_
                        rhs24_ = _dafny.SeqWithoutIsStrInference([((d_167_v80_).f7) == ((0) - (len(((d_179_v90_)[(d_167_v80_).f7] if ((d_167_v80_).f7) in (d_179_v90_) else d_115_v40_)))), ((d_180_v91_).cardinality) != ((d_167_v80_).f3), d_167_v80_.f4])
                        rhs25_ = default__.fm0(d_71_globalState_)
                        rhs26_ = d_178_v89_
                        rhs27_ = default__.fm4(d_167_v80_.f4, d_72_v1_, default__.fm2((0) - ((d_167_v80_).f3), d_71_globalState_), (d_107_v32_) < ((d_181_v92_).cardinality), d_71_globalState_)
                        lhs13_ = d_112_v37_
                        lhs14_ = default__.safeIndex(699, (d_112_v37_).length(0))
                        d_175_v86_ = rhs23_
                        d_115_v40_ = rhs24_
                        lhs13_[lhs14_] = rhs25_
                        d_178_v89_ = rhs26_
                        d_178_v89_ = rhs27_
                        d_182_v93_: _dafny.Map
                        d_182_v93_ = _dafny.Map({(d_167_v80_).f7: d_167_v80_.f4})
                        (d_71_globalState_).f2 = ((d_182_v93_)[(d_175_v86_).f14] if ((d_175_v86_).f14) in (d_182_v93_) else d_167_v80_.f4)
                        (d_167_v80_).f4 = default__.fm2((0) - (default__.safeDivisionInt((d_167_v80_).f7, ((d_181_v92_)[(d_167_v80_).f7] if ((d_167_v80_).f7) in (d_181_v92_) else (d_167_v80_).f3))), d_71_globalState_)
                        d_183_v94_: C7
                        nw31_ = C7()
                        nw31_.ctor__(not(False), default__.fm0(d_71_globalState_), d_107_v32_, d_167_v80_.f4)
                        d_183_v94_ = nw31_
                        d_184_v95_: _dafny.Array
                        nw32_ = _dafny.Array(None, 5)
                        nw32_[int(0)] = d_183_v94_
                        nw32_[int(1)] = d_183_v94_
                        nw32_[int(2)] = d_183_v94_
                        nw32_[int(3)] = d_183_v94_
                        nw32_[int(4)] = d_183_v94_
                        d_184_v95_ = nw32_
                        index19_ = default__.safeIndex(150, (d_184_v95_).length(0))
                        (d_184_v95_)[index19_] = d_183_v94_
                        index20_ = default__.safeIndex(699, (d_112_v37_).length(0))
                        index21_ = default__.safeIndex(150, (d_184_v95_).length(0))
                        rhs28_ = (d_175_v86_).f14
                        rhs29_ = default__.fm0(d_71_globalState_)
                        rhs30_ = (d_167_v80_).f7
                        rhs31_ = (d_175_v86_).f15
                        rhs32_ = d_183_v94_
                        lhs15_ = d_71_globalState_
                        lhs16_ = d_112_v37_
                        lhs17_ = default__.safeIndex(699, (d_112_v37_).length(0))
                        lhs18_ = d_71_globalState_
                        lhs19_ = d_184_v95_
                        lhs20_ = default__.safeIndex(150, (d_184_v95_).length(0))
                        lhs15_.f1 = rhs28_
                        lhs16_[lhs17_] = rhs29_
                        lhs18_.f1 = rhs30_
                        d_170_v83_ = rhs31_
                        lhs19_[lhs20_] = rhs32_
                    elif True:
                        d_185_v96_: _dafny.Map
                        d_185_v96_ = _dafny.Map({(d_167_v80_).f7: d_69_v0_})
                        d_186_v97_: _dafny.MultiSet
                        d_186_v97_ = _dafny.MultiSet([((d_185_v96_)[(d_167_v80_).f7] if ((d_167_v80_).f7) in (d_185_v96_) else d_69_v0_)])
                        d_186_v97_ = ((d_186_v97_) - (_dafny.MultiSet([d_69_v0_]))) | ((d_186_v97_) - (d_186_v97_))
                        d_187_v98_: _dafny.Map
                        d_187_v98_ = _dafny.Map({d_109_v34_: d_111_v36_})
                        d_188_v99_: _dafny.MultiSet
                        d_188_v99_ = _dafny.MultiSet([True])
                        d_189_v100_: _dafny.Map
                        d_189_v100_ = _dafny.Map({((d_187_v98_).set(d_109_v34_, d_113_v38_.f13)) | (_dafny.Map({d_109_v34_: d_113_v38_.f13})): (len(d_108_v33_)) * ((d_188_v99_).cardinality)})
                        d_189_v100_ = (d_189_v100_).set(d_187_v98_, (d_167_v80_).f3)
                        index22_ = default__.safeIndex(699, (d_112_v37_).length(0))
                        (d_112_v37_)[index22_] = (d_113_v38_).fm12(default__.fm2(d_107_v32_, d_71_globalState_), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jd"))) == (d_111_v36_), d_71_globalState_)
                        (d_167_v80_).f4 = (((d_111_v36_) + (d_113_v38_.f13)).set(default__.safeIndex((d_167_v80_).f3, len((d_111_v36_) + (d_113_v38_.f13))), (d_113_v38_).f12)) < (d_111_v36_)
                        d_190_v101_: D12
                        d_190_v101_ = D12_DC35((d_109_v34_).f9, True, _dafny.SeqWithoutIsStrInference([d_72_v1_]), _dafny.SeqWithoutIsStrInference([(d_113_v38_).f12 for d_191_i14_ in range(default__.abs(-514))]), (d_167_v80_).f7)
                        pat_let_tv0_ = d_113_v38_
                        pat_let_tv1_ = d_167_v80_
                        pat_let_tv2_ = d_115_v40_
                        pat_let_tv3_ = d_106_v31_
                        def iife7_(_pat_let0_0):
                            def iife8_(d_192_dt__update__tmp_h0_):
                                def iife9_(_pat_let1_0):
                                    def iife10_(d_193_dt__update_hcf59_h0_):
                                        def iife11_(_pat_let2_0):
                                            def iife12_(d_194_dt__update_hcf60_h0_):
                                                def iife13_(_pat_let3_0):
                                                    def iife14_(d_195_dt__update_hcf58_h0_):
                                                        def iife15_(_pat_let4_0):
                                                            def iife16_(d_196_dt__update_hcf56_h0_):
                                                                return D12_DC35(d_196_dt__update_hcf56_h0_, (d_192_dt__update__tmp_h0_).cf57, d_195_dt__update_hcf58_h0_, d_193_dt__update_hcf59_h0_, d_194_dt__update_hcf60_h0_)
                                                            return iife16_(_pat_let4_0)
                                                        return iife15_(pat_let_tv3_)
                                                    return iife14_(_pat_let3_0)
                                                return iife13_(pat_let_tv2_)
                                            return iife12_(_pat_let2_0)
                                        return iife11_((pat_let_tv1_).f7)
                                    return iife10_(_pat_let1_0)
                                return iife9_(pat_let_tv0_.f13)
                            return iife8_(_pat_let0_0)
                        d_111_v36_ = (iife7_(d_190_v101_)).cf59
                    (d_113_v38_).m5(d_71_globalState_)
                    pass
            pass
        d_197_v102_: _dafny.Map
        d_197_v102_ = _dafny.Map({d_72_v1_: d_110_v35_})
        d_198_v103_: _dafny.Seq
        d_198_v103_ = _dafny.SeqWithoutIsStrInference([d_107_v32_])
        d_199_i15_: int
        d_199_i15_ = 0
        with _dafny.label("1"):
            while not ((_dafny.SeqWithoutIsStrInference([d_107_v32_ for d_210_i16_ in range(default__.abs(-792))])) < (d_198_v103_)) or (((d_113_v38_.f13).set(default__.safeIndex(d_107_v32_, len(d_113_v38_.f13)), ((d_197_v102_)[d_72_v1_] if (d_72_v1_) in (d_197_v102_) else (d_113_v38_).f12))) == (d_111_v36_)):
                with _dafny.c_label("1"):
                    if (d_199_i15_) >= (100):
                        raise _dafny.Break("1")
                    d_199_i15_ = (d_199_i15_) + (1)
                    d_200_v104_: _dafny.MultiSet
                    d_200_v104_ = _dafny.MultiSet([d_72_v1_])
                    (d_71_globalState_).f1 = ((len(d_113_v38_.f13)) - ((0) - (d_107_v32_))) - ((d_200_v104_).cardinality)
                    d_201_v105_: _dafny.Map
                    d_201_v105_ = _dafny.Map({(d_72_v1_) and (d_72_v1_): d_72_v1_})
                    d_201_v105_ = (d_201_v105_).set((d_107_v32_) <= (d_107_v32_), d_72_v1_)
                    d_202_v106_: C4
                    nw33_ = C4()
                    nw33_.ctor__((d_200_v104_).cardinality, (0) - (len(d_198_v103_)), default__.fm2(d_107_v32_, d_71_globalState_))
                    d_202_v106_ = nw33_
                    d_203_v107_: _dafny.Map
                    d_203_v107_ = _dafny.Map({d_72_v1_: d_202_v106_})
                    d_204_v108_: _dafny.Seq
                    d_204_v108_ = _dafny.SeqWithoutIsStrInference([d_202_v106_, d_202_v106_])
                    d_205_v109_: _dafny.Array
                    nw34_ = _dafny.Array(None, 14)
                    nw34_[int(0)] = d_202_v106_
                    nw34_[int(1)] = d_202_v106_
                    nw34_[int(2)] = d_202_v106_
                    nw34_[int(3)] = d_202_v106_
                    nw34_[int(4)] = d_202_v106_
                    nw34_[int(5)] = d_202_v106_
                    nw34_[int(6)] = d_202_v106_
                    nw34_[int(7)] = ((d_203_v107_)[True] if (True) in (d_203_v107_) else d_202_v106_)
                    nw34_[int(8)] = d_202_v106_
                    nw34_[int(9)] = d_202_v106_
                    nw34_[int(10)] = d_202_v106_
                    nw34_[int(11)] = d_202_v106_
                    nw34_[int(12)] = (d_204_v108_)[default__.safeIndex(d_107_v32_, len(d_204_v108_))]
                    nw34_[int(13)] = d_202_v106_
                    d_205_v109_ = nw34_
                    index23_ = default__.safeIndex(199, (d_205_v109_).length(0))
                    (d_205_v109_)[index23_] = d_202_v106_
                    index24_ = default__.safeIndex(199, (d_205_v109_).length(0))
                    (d_205_v109_)[index24_] = d_202_v106_
                    d_206_v110_: _dafny.MultiSet
                    d_207_v111_: str
                    d_208_v112_: _dafny.Map
                    d_209_v113_: bool
                    out1_: _dafny.MultiSet
                    out2_: str
                    out3_: _dafny.Map
                    out4_: bool
                    out1_, out2_, out3_, out4_ = (d_109_v34_).m2((d_72_v1_) == (True), d_112_v37_, (d_201_v105_) | (d_201_v105_), d_72_v1_, d_71_globalState_)
                    d_206_v110_ = out1_
                    d_207_v111_ = out2_
                    d_208_v112_ = out3_
                    d_209_v113_ = out4_
                    pass
            pass
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_106_v31_).length(0)):
            d_211_i17_: int = guard_loop_0_
            if (True) and (((0) <= (d_211_i17_)) and ((d_211_i17_) < ((d_106_v31_).length(0)))):
                (d_106_v31_)[(d_211_i17_)] = not(not(not(d_72_v1_)))
        d_212_v114_: _dafny.Array
        nw35_ = _dafny.Array(None, 4)
        nw35_[int(0)] = d_113_v38_
        nw35_[int(1)] = (d_113_v38_ if d_72_v1_ else d_113_v38_)
        nw35_[int(2)] = d_113_v38_
        nw35_[int(3)] = d_113_v38_
        d_212_v114_ = nw35_
        index25_ = default__.safeIndex(998, (d_212_v114_).length(0))
        (d_212_v114_)[index25_] = d_113_v38_
        index26_ = default__.safeIndex(998, (d_212_v114_).length(0))
        (d_212_v114_)[index26_] = d_113_v38_
        d_213_v115_: _dafny.Array
        nw36_ = _dafny.Array(_dafny.Map({}), 8)
        d_213_v115_ = nw36_
        d_214_v116_: T0
        nw37_ = C3()
        nw37_.ctor__((d_113_v38_).f12, d_111_v36_, len(d_113_v38_.f13), True, d_107_v32_)
        d_214_v116_ = nw37_
        d_215_v117_: C2
        nw38_ = C2()
        nw38_.ctor__((d_214_v116_).f3, d_112_v37_, (d_214_v116_).f3, d_72_v1_)
        d_215_v117_ = nw38_
        d_216_v118_: _dafny.Map
        d_216_v118_ = _dafny.Map({d_214_v116_: d_215_v117_})
        index27_ = default__.safeIndex(91, (d_213_v115_).length(0))
        (d_213_v115_)[index27_] = (d_216_v118_) | (d_216_v118_)
        index28_ = default__.safeIndex(91, (d_213_v115_).length(0))
        (d_213_v115_)[index28_] = d_216_v118_
        d_217_v119_: _dafny.Seq
        d_217_v119_ = _dafny.SeqWithoutIsStrInference([d_112_v37_, d_112_v37_])
        d_218_v120_: _dafny.Seq
        d_218_v120_ = _dafny.SeqWithoutIsStrInference([d_217_v119_, (D15_DC43(d_217_v119_)).cf74])
        index29_ = default__.safeIndex(209, (d_106_v31_).length(0))
        (d_106_v31_)[index29_] = ((d_215_v117_).f15) not in ((d_218_v120_)[default__.safeIndex((d_215_v117_).f14, len(d_218_v120_))])
        index30_ = default__.safeIndex(209, (d_106_v31_).length(0))
        (d_106_v31_)[index30_] = d_72_v1_
        d_219_i18_: int
        d_219_i18_ = 0
        with _dafny.label("2"):
            while d_72_v1_:
                with _dafny.c_label("2"):
                    if (d_219_i18_) >= (100):
                        raise _dafny.Break("2")
                    d_219_i18_ = (d_219_i18_) + (1)
                    d_220_v121_: _dafny.Map
                    d_220_v121_ = _dafny.Map({d_214_v116_.f4: d_214_v116_.f4})
                    d_220_v121_ = (d_220_v121_).set((d_106_v31_)[default__.safeIndex(209, (d_106_v31_).length(0))], d_214_v116_.f4)
                    d_221_v122_: _dafny.MultiSet
                    d_221_v122_ = _dafny.MultiSet([d_107_v32_])
                    (d_71_globalState_).f2 = default__.fm2(((d_221_v122_).intersection(d_221_v122_)).cardinality, d_71_globalState_)
                    d_222_v123_: C0
                    nw39_ = C0()
                    nw39_.ctor__(d_214_v116_.f4, (d_109_v34_).f9)
                    d_222_v123_ = nw39_
                    (d_214_v116_).f4 = not(d_214_v116_.f4)
                    pass
            pass
        index31_ = default__.safeIndex(911, (d_112_v37_).length(0))
        (d_112_v37_)[index31_] = (d_214_v116_).f3
        index32_ = default__.safeIndex(911, (d_112_v37_).length(0))
        (d_112_v37_)[index32_] = d_107_v32_
        d_223_v124_: _dafny.Map
        d_223_v124_ = _dafny.Map({d_72_v1_: False})
        d_224_v125_: D13
        d_224_v125_ = D13_DC38((_dafny.SeqWithoutIsStrInference([_dafny.Map({d_214_v116_.f4: (d_106_v31_)[default__.safeIndex(209, (d_106_v31_).length(0))]}) for d_225_i19_ in range(default__.abs(144))])).set(default__.safeIndex((d_215_v117_).f14, len(_dafny.SeqWithoutIsStrInference([_dafny.Map({d_214_v116_.f4: (d_106_v31_)[default__.safeIndex(209, (d_106_v31_).length(0))]}) for d_226_i19_ in range(default__.abs(144))]))), d_223_v124_))
        source3_ = d_224_v125_
        if source3_.is_DC39:
            d_227___mcc_h0_ = source3_.cf69
            d_228_cf69_ = d_227___mcc_h0_
            (d_71_globalState_).f2 = d_214_v116_.f4
            d_111_v36_ = (default__.fm15((d_106_v31_)[default__.safeIndex(209, (d_106_v31_).length(0))], (d_215_v117_).f14, (d_214_v116_).f3, d_71_globalState_)) + (_dafny.SeqWithoutIsStrInference([(d_113_v38_).f12 for d_229_i20_ in range(default__.abs(-238))]))
            d_230_v126_: _dafny.MultiSet
            d_230_v126_ = _dafny.MultiSet([d_214_v116_.f4])
            d_231_v127_: _dafny.Map
            d_231_v127_ = _dafny.Map({(d_215_v117_).f14: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sqq")))})
            (d_71_globalState_).f1 = (((d_230_v126_ if (d_113_v38_).fm13(d_71_globalState_) else d_230_v126_)).set((default__.fm2(d_107_v32_, d_71_globalState_)) and (d_214_v116_.f4), default__.abs(((d_231_v127_)[len(_dafny.Set({d_109_v34_}))] if (len(_dafny.Set({d_109_v34_}))) in (d_231_v127_) else default__.fm0(d_71_globalState_))))).cardinality
            d_232_v128_: T1
            nw40_ = C3()
            nw40_.ctor__((d_113_v38_).f12, d_113_v38_.f13, len(d_162_v76_), d_214_v116_.f4, (d_215_v117_).f14)
            d_232_v128_ = nw40_
            d_233_v129_: _dafny.MultiSet
            d_233_v129_ = _dafny.MultiSet([d_232_v128_])
            d_234_v130_: D15
            d_234_v130_ = D15_DC44(d_232_v128_, d_110_v35_, d_107_v32_)
            d_235_v131_: _dafny.Seq
            d_235_v131_ = _dafny.SeqWithoutIsStrInference([d_232_v128_])
            d_233_v129_ = (_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([(d_234_v130_).cf75])) + (d_235_v131_))) - (_dafny.MultiSet(d_235_v131_))
        elif source3_.is_DC40:
            d_236___mcc_h1_ = source3_.cf70
            d_237___mcc_h2_ = source3_.cf71
            d_238_cf71_ = d_237___mcc_h2_
            d_239_cf70_ = d_236___mcc_h1_
            d_240_v132_: _dafny.Seq
            d_240_v132_ = _dafny.SeqWithoutIsStrInference([d_115_v40_, _dafny.SeqWithoutIsStrInference([d_214_v116_.f4, d_72_v1_, (d_106_v31_)[default__.safeIndex(209, (d_106_v31_).length(0))], d_214_v116_.f4])])
            d_241_v133_: C6
            nw41_ = C6()
            nw41_.ctor__((d_215_v117_).f14, d_238_cf71_, d_214_v116_.f4, len((d_240_v132_)[default__.safeIndex(d_107_v32_, len(d_240_v132_))]))
            d_241_v133_ = nw41_
            d_242_v134_: _dafny.Seq
            d_242_v134_ = _dafny.SeqWithoutIsStrInference([d_241_v133_, d_241_v133_, d_241_v133_])
            d_243_v135_: _dafny.MultiSet
            d_243_v135_ = _dafny.MultiSet([(_dafny.MultiSet(d_198_v103_)).cardinality, (d_215_v117_).f14, 844])
            d_244_v136_: D3
            d_244_v136_ = D3_DC10((d_243_v135_).cardinality, (d_109_v34_).f9, (_dafny.MultiSet([(0) - (((d_243_v135_)[(d_215_v117_).f14] if ((d_215_v117_).f14) in (d_243_v135_) else (d_112_v37_)[default__.safeIndex(911, (d_112_v37_).length(0))])), (d_215_v117_).f14, (d_214_v116_).f3])).cardinality)
            d_245_v137_: _dafny.Array
            nw42_ = _dafny.Array(None, 24)
            nw42_[int(0)] = d_241_v133_
            nw42_[int(1)] = d_241_v133_
            nw42_[int(2)] = d_241_v133_
            nw42_[int(3)] = d_241_v133_
            nw42_[int(4)] = d_241_v133_
            nw42_[int(5)] = d_241_v133_
            nw42_[int(6)] = d_241_v133_
            nw42_[int(7)] = d_241_v133_
            nw42_[int(8)] = d_241_v133_
            nw42_[int(9)] = (d_241_v133_ if d_72_v1_ else d_241_v133_)
            nw42_[int(10)] = (d_242_v134_)[default__.safeIndex((d_244_v136_).cf15, len(d_242_v134_))]
            nw42_[int(11)] = d_241_v133_
            nw42_[int(12)] = d_241_v133_
            nw42_[int(13)] = d_241_v133_
            nw42_[int(14)] = d_241_v133_
            nw42_[int(15)] = d_241_v133_
            nw42_[int(16)] = d_241_v133_
            nw42_[int(17)] = d_241_v133_
            nw42_[int(18)] = d_241_v133_
            nw42_[int(19)] = d_241_v133_
            nw42_[int(20)] = d_241_v133_
            nw42_[int(21)] = d_241_v133_
            nw42_[int(22)] = d_241_v133_
            nw42_[int(23)] = d_241_v133_
            d_245_v137_ = nw42_
            rhs33_ = len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qxab"))) + (d_113_v38_.f13))
            rhs34_ = d_245_v137_
            d_107_v32_ = rhs33_
            d_245_v137_ = rhs34_
            d_246_v138_: _dafny.Map
            d_246_v138_ = _dafny.Map({False: (d_214_v116_).f3})
            d_247_v139_: D12
            d_247_v139_ = D12_DC37(d_223_v124_, d_72_v1_, len(d_111_v36_), len(d_246_v138_))
            d_248_v140_: D1
            d_248_v140_ = D1_DC6(d_214_v116_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pvomejghm")), (d_106_v31_)[default__.safeIndex(209, (d_106_v31_).length(0))], d_214_v116_.f4)
            source4_ = (d_247_v139_ if d_214_v116_.f4 else D12_DC37(d_223_v124_, (d_248_v140_).cf12, (d_214_v116_).f3, (d_215_v117_).f14))
            if source4_.is_DC35:
                d_249___mcc_h5_ = source4_.cf56
                d_250___mcc_h6_ = source4_.cf57
                d_251___mcc_h7_ = source4_.cf58
                d_252___mcc_h8_ = source4_.cf59
                d_253___mcc_h9_ = source4_.cf60
                d_254_cf60_ = d_253___mcc_h9_
                d_255_cf59_ = d_252___mcc_h8_
                d_256_cf58_ = d_251___mcc_h7_
                d_257_cf57_ = d_250___mcc_h6_
                d_258_cf56_ = d_249___mcc_h5_
                d_259_v141_: C4
                nw43_ = C4()
                nw43_.ctor__(len(d_108_v33_), (d_241_v133_).f8, d_214_v116_.f4)
                d_259_v141_ = nw43_
                d_260_v142_: _dafny.Map
                d_260_v142_ = _dafny.Map({d_107_v32_: d_259_v141_})
                d_261_v143_: _dafny.Map
                d_261_v143_ = _dafny.Map({True: (((d_260_v142_)[d_238_cf71_] if (d_238_cf71_) in (d_260_v142_) else d_259_v141_) if d_214_v116_.f4 else d_259_v141_)})
                d_261_v143_ = d_261_v143_
                d_262_v144_: _dafny.Array
                def lambda9_(d_263_v40_):
                    def lambda10_(d_264_i21_):
                        return d_263_v40_

                    return lambda10_

                init5_ = lambda9_(d_115_v40_)
                nw44_ = _dafny.Array(None, 12)
                for i0_5_ in range(nw44_.length(0)):
                    nw44_[i0_5_] = init5_(i0_5_)
                d_262_v144_ = nw44_
                d_265_v145_: _dafny.Map
                d_265_v145_ = _dafny.Map({d_112_v37_: (d_241_v133_).f8})
                (d_215_v117_).m6(d_262_v144_, d_238_cf71_, (len(d_265_v145_)) + ((d_215_v117_).f14), d_71_globalState_)
                d_223_v124_ = (d_223_v124_).set(not(not(True)), d_214_v116_.f4)
                d_266_v146_: C7
                nw45_ = C7()
                nw45_.ctor__(d_257_cf57_, default__.fm0(d_71_globalState_), (len(d_162_v76_)) * ((d_215_v117_).f14), True)
                d_266_v146_ = nw45_
            elif source4_.is_DC36:
                d_267___mcc_h10_ = source4_.cf61
                d_268___mcc_h11_ = source4_.cf62
                d_269___mcc_h12_ = source4_.cf63
                d_270_cf63_ = d_269___mcc_h12_
                d_271_cf62_ = d_268___mcc_h11_
                d_272_cf61_ = d_267___mcc_h10_
                d_273_v147_: _dafny.Array
                nw46_ = _dafny.Array(_dafny.Array(None, 0), 2)
                d_273_v147_ = nw46_
                d_274_v148_: C0
                nw47_ = C0()
                nw47_.ctor__(False, (d_109_v34_).f9)
                d_274_v148_ = nw47_
                d_275_v149_: _dafny.Array
                nw48_ = _dafny.Array(None, 3)
                nw48_[int(0)] = d_274_v148_
                nw48_[int(1)] = d_274_v148_
                nw48_[int(2)] = d_274_v148_
                d_275_v149_ = nw48_
                index33_ = default__.safeIndex(524, (d_273_v147_).length(0))
                (d_273_v147_)[index33_] = d_275_v149_
                d_276_v150_: D12
                d_276_v150_ = D12_DC35((d_109_v34_).f9, d_72_v1_, d_115_v40_, d_113_v38_.f13, len(_dafny.SeqWithoutIsStrInference([(d_113_v38_).f12 for d_277_i22_ in range(default__.abs(-843))])))
                d_278_v151_: D4
                d_278_v151_ = D4_DC11(_dafny.Set({737, (d_276_v150_).cf60}))
                d_279_v152_: _dafny.Map
                d_279_v152_ = _dafny.Map({d_278_v151_: d_243_v135_})
                index34_ = default__.safeIndex(524, (d_273_v147_).length(0))
                rhs35_ = ((d_279_v152_)[d_278_v151_] if (d_278_v151_) in (d_279_v152_) else (d_243_v135_) - (d_243_v135_))
                rhs36_ = (0) - ((_dafny.MultiSet(d_198_v103_)).cardinality)
                rhs37_ = d_113_v38_.f13
                rhs38_ = d_115_v40_
                rhs39_ = (d_275_v149_ if d_214_v116_.f4 else d_275_v149_)
                lhs21_ = d_71_globalState_
                lhs22_ = d_113_v38_
                lhs23_ = d_273_v147_
                lhs24_ = default__.safeIndex(524, (d_273_v147_).length(0))
                d_243_v135_ = rhs35_
                lhs21_.f1 = rhs36_
                lhs22_.f13 = rhs37_
                d_115_v40_ = rhs38_
                lhs23_[lhs24_] = rhs39_
                d_280_v153_: _dafny.Set
                d_280_v153_ = _dafny.Set({d_112_v37_, (d_215_v117_).f15, (d_215_v117_).f15, (d_215_v117_).f15})
                d_280_v153_ = ((d_280_v153_).intersection(_dafny.Set({(d_215_v117_).f15}))) - ((d_280_v153_).intersection(d_280_v153_))
                d_281_v155_: _dafny.MultiSet
                d_281_v155_ = _dafny.MultiSet([d_110_v35_, (d_113_v38_).f12, (d_113_v38_).f12, (d_113_v38_).f12, (d_113_v38_).f12])
                d_282_v156_: _dafny.Seq
                def iife17_():
                    coll7_ = _dafny.Map()
                    compr_7_: str
                    for compr_7_ in (d_281_v155_).Elements:
                        d_283_v154_: str = compr_7_
                        if (d_283_v154_) in (d_281_v155_):
                            coll7_[d_283_v154_] = d_274_v148_.f10
                    return _dafny.Map(coll7_)
                d_282_v156_ = _dafny.SeqWithoutIsStrInference([iife17_()
                ])
                d_284_v157_: _dafny.MultiSet
                d_284_v157_ = _dafny.MultiSet([(d_282_v156_)[default__.safeIndex((d_214_v116_).f3, len(d_282_v156_))]])
                d_285_v158_: D10
                d_285_v158_ = D10_DC29(not(d_72_v1_), (d_113_v38_).f12)
                d_286_v159_: _dafny.Map
                d_286_v159_ = _dafny.Map({(d_285_v158_).cf43: d_214_v116_.f4})
                d_238_cf71_ = (0) - (((d_284_v157_)[d_286_v159_] if (d_286_v159_) in (d_284_v157_) else d_107_v32_))
                (d_274_v148_).f10 = (((0) - (default__.fm0(d_71_globalState_))) + (690)) != ((d_214_v116_).f3)
            elif source4_.is_DC37:
                d_287___mcc_h13_ = source4_.cf64
                d_288___mcc_h14_ = source4_.cf65
                d_289___mcc_h15_ = source4_.cf66
                d_290___mcc_h16_ = source4_.cf67
                d_291_cf67_ = d_290___mcc_h16_
                d_292_cf66_ = d_289___mcc_h15_
                d_293_cf65_ = d_288___mcc_h14_
                d_294_cf64_ = d_287___mcc_h13_
                d_106_v31_ = d_106_v31_
                (d_71_globalState_).f2 = default__.fm2(d_107_v32_, d_71_globalState_)
                (d_214_v116_).f4 = d_214_v116_.f4
                d_295_v160_: _dafny.Map
                d_295_v160_ = _dafny.Map({d_72_v1_: d_113_v38_.f13})
                d_296_v161_: D16
                d_296_v161_ = D16_DC45(d_295_v160_)
                d_297_v162_: _dafny.Array
                nw49_ = _dafny.Array(None, 4)
                nw49_[int(0)] = d_295_v160_
                nw49_[int(1)] = (d_295_v160_) | (d_295_v160_)
                nw49_[int(2)] = (d_295_v160_) | ((_dafny.Map({d_293_cf65_: d_111_v36_})).set(False, d_111_v36_))
                nw49_[int(3)] = (d_296_v161_).cf78
                d_297_v162_ = nw49_
                index35_ = default__.safeIndex(349, (d_297_v162_).length(0))
                (d_297_v162_)[index35_] = ((D16_DC45(_dafny.Map({not(d_72_v1_): d_111_v36_}))).cf78) | (d_295_v160_)
                index36_ = default__.safeIndex(349, (d_297_v162_).length(0))
                (d_297_v162_)[index36_] = d_295_v160_
            elif True:
                d_298___mcc_h17_ = source4_.cf55
                d_299_cf55_ = d_298___mcc_h17_
                d_300_v163_: _dafny.Array
                nw50_ = _dafny.Array(None, 26)
                d_300_v163_ = nw50_
                d_301_v164_: _dafny.Set
                d_301_v164_ = _dafny.Set({d_300_v163_, d_300_v163_, d_300_v163_})
                d_301_v164_ = (d_301_v164_) | (_dafny.Set({d_300_v163_, d_300_v163_}))
                (d_71_globalState_).f1 = len(d_113_v38_.f13)
                d_302_v165_: _dafny.MultiSet
                d_303_v166_: str
                d_304_v167_: _dafny.Map
                d_305_v168_: bool
                out5_: _dafny.MultiSet
                out6_: str
                out7_: _dafny.Map
                out8_: bool
                out5_, out6_, out7_, out8_ = (d_109_v34_).m2((d_106_v31_)[default__.safeIndex(209, (d_106_v31_).length(0))], (d_215_v117_).f15, (d_223_v124_) | (d_223_v124_), d_72_v1_, d_71_globalState_)
                d_302_v165_ = out5_
                d_303_v166_ = out6_
                d_304_v167_ = out7_
                d_305_v168_ = out8_
                d_306_v169_: int
                out9_: int
                out9_ = (d_109_v34_).m1(424, d_71_globalState_)
                d_306_v169_ = out9_
            d_307_v170_: C5
            nw51_ = C5()
            nw51_.ctor__((d_109_v34_).f9, (d_215_v117_).f14, (d_215_v117_).f14, d_72_v1_)
            d_307_v170_ = nw51_
            if d_214_v116_.f4:
                d_308_v171_: _dafny.Array
                nw52_ = _dafny.Array(_dafny.Array(None, 0), 1)
                d_308_v171_ = nw52_
                d_309_v172_: _dafny.Array
                nw53_ = _dafny.Array(D8.default()(), 4)
                d_309_v172_ = nw53_
                index37_ = default__.safeIndex(771, (d_308_v171_).length(0))
                (d_308_v171_)[index37_] = (d_309_v172_ if False else d_309_v172_)
                index38_ = default__.safeIndex(771, (d_308_v171_).length(0))
                (d_308_v171_)[index38_] = d_309_v172_
                index39_ = default__.safeIndex(911, (d_112_v37_).length(0))
                (d_112_v37_)[index39_] = 185
                d_310_v173_: C0
                nw54_ = C0()
                nw54_.ctor__(d_214_v116_.f4, (d_109_v34_).f9)
                d_310_v173_ = nw54_
                d_310_v173_ = d_310_v173_
                d_311_v174_: T1
                nw55_ = C5()
                nw55_.ctor__((d_109_v34_).f9, d_107_v32_, d_107_v32_, True)
                d_311_v174_ = nw55_
                d_312_v175_: _dafny.Seq
                d_312_v175_ = _dafny.SeqWithoutIsStrInference([d_311_v174_, d_311_v174_])
                d_313_v176_: _dafny.Map
                d_313_v176_ = _dafny.Map({d_312_v175_: (d_241_v133_).f8})
                d_313_v176_ = d_313_v176_
                index40_ = default__.safeIndex(209, (d_106_v31_).length(0))
                index41_ = default__.safeIndex(209, (d_106_v31_).length(0))
                index42_ = default__.safeIndex(209, (d_106_v31_).length(0))
                rhs40_ = d_311_v174_.f4
                rhs41_ = (False) or ((d_310_v173_).fm5((d_241_v133_).f8, (d_311_v174_).f3, d_110_v35_, d_72_v1_, d_71_globalState_))
                rhs42_ = len(((d_198_v103_) + (_dafny.SeqWithoutIsStrInference([len(d_113_v38_.f13) for d_314_i23_ in range(default__.abs(-229))]))) + (d_198_v103_))
                rhs43_ = d_214_v116_.f4
                lhs25_ = d_106_v31_
                lhs26_ = default__.safeIndex(209, (d_106_v31_).length(0))
                lhs27_ = d_106_v31_
                lhs28_ = default__.safeIndex(209, (d_106_v31_).length(0))
                lhs29_ = d_71_globalState_
                lhs30_ = d_106_v31_
                lhs31_ = default__.safeIndex(209, (d_106_v31_).length(0))
                lhs25_[lhs26_] = rhs40_
                lhs27_[lhs28_] = rhs41_
                lhs29_.f1 = rhs42_
                lhs30_[lhs31_] = rhs43_
            elif True:
                d_72_v1_ = d_214_v116_.f4
                d_315_v177_: C6
                nw56_ = C6()
                nw56_.ctor__((0) - (len(d_111_v36_)), (d_215_v117_).f14, d_214_v116_.f4, d_238_cf71_)
                d_315_v177_ = nw56_
                d_316_v178_: _dafny.Array
                def lambda11_(d_317_v177_):
                    def lambda12_(d_318_i24_):
                        return _dafny.SeqWithoutIsStrInference([(d_317_v177_).f8])

                    return lambda12_

                init6_ = lambda11_(d_315_v177_)
                nw57_ = _dafny.Array(None, 6)
                for i0_6_ in range(nw57_.length(0)):
                    nw57_[i0_6_] = init6_(i0_6_)
                d_316_v178_ = nw57_
                nw58_ = _dafny.Array(_dafny.Seq({}), 6)
                d_316_v178_ = nw58_
                index43_ = default__.safeIndex(911, (d_112_v37_).length(0))
                (d_112_v37_)[index43_] = (d_214_v116_).f3
                d_319_v179_: _dafny.Map
                d_319_v179_ = _dafny.Map({(d_109_v34_).f9: _dafny.Set({d_106_v31_})})
                d_320_v180_: _dafny.Set
                d_320_v180_ = _dafny.Set({(d_109_v34_).f9, (d_109_v34_).f9})
                d_319_v179_ = (d_319_v179_).set((d_109_v34_).f9, d_320_v180_)
        elif source3_.is_DC38:
            d_321___mcc_h3_ = source3_.cf68
            d_322_cf68_ = d_321___mcc_h3_
            index44_ = default__.safeIndex(911, (d_112_v37_).length(0))
            (d_112_v37_)[index44_] = (d_214_v116_).f3
            d_323_v181_: _dafny.Map
            d_323_v181_ = _dafny.Map({d_215_v117_: -124})
            d_323_v181_ = (d_323_v181_).set(d_215_v117_, (d_112_v37_)[default__.safeIndex(911, (d_112_v37_).length(0))])
            d_324_v182_: C7
            nw59_ = C7()
            nw59_.ctor__(d_214_v116_.f4, d_107_v32_, (d_215_v117_).f14, d_72_v1_)
            d_324_v182_ = nw59_
            d_325_v183_: _dafny.Map
            d_325_v183_ = _dafny.Map({d_324_v182_: d_115_v40_})
            d_326_v184_: _dafny.MultiSet
            d_326_v184_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_72_v1_, d_72_v1_]), ((d_325_v183_)[d_324_v182_] if (d_324_v182_) in (d_325_v183_) else (d_115_v40_).set(default__.safeIndex(len(d_111_v36_), len(d_115_v40_)), False)), d_115_v40_])
            d_326_v184_ = ((d_326_v184_).set(d_115_v40_, default__.abs(len(d_162_v76_)))) | ((D17_DC48(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_214_v116_.f4])]))).cf83)
            index45_ = default__.safeIndex(209, (d_106_v31_).length(0))
            (d_106_v31_)[index45_] = not ((d_106_v31_)[default__.safeIndex(209, (d_106_v31_).length(0))]) or (d_324_v182_.f5)
        elif True:
            d_327___mcc_h4_ = source3_.cf72
            d_328_cf72_ = d_327___mcc_h4_
            d_112_v37_ = (d_215_v117_).f15
            (d_71_globalState_).f1 = (d_214_v116_).f3
            (d_71_globalState_).f2 = d_72_v1_
            d_329_v185_: _dafny.Set
            d_329_v185_ = _dafny.Set({d_69_v0_})
            rhs44_ = not((d_329_v185_).ispropersubset((d_329_v185_ if (d_106_v31_)[default__.safeIndex(209, (d_106_v31_).length(0))] else d_329_v185_)))
            rhs45_ = d_72_v1_
            rhs46_ = d_113_v38_.f13
            lhs32_ = d_71_globalState_
            lhs33_ = d_113_v38_
            lhs32_.f2 = rhs44_
            d_72_v1_ = rhs45_
            lhs33_.f13 = rhs46_
        _dafny.print(_dafny.string_of((d_69_v0_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_69_v0_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_69_v0_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_69_v0_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_69_v0_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_69_v0_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_69_v0_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_69_v0_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_69_v0_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_69_v0_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_69_v0_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_69_v0_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_69_v0_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_69_v0_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_69_v0_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_69_v0_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_69_v0_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_69_v0_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_69_v0_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_69_v0_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_69_v0_)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_69_v0_)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_69_v0_)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_69_v0_)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_69_v0_)[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_69_v0_)[25]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_).f0)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_).f0)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_).f0)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_).f0)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_).f0)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_).f0)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_).f0)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_).f0)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_).f0)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_).f0)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_).f0)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_).f0)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_).f0)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_).f0)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_).f0)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_).f0)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_).f0)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_).f0)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_).f0)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_).f0)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_).f0)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_).f0)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_).f0)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_).f0)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_).f0)[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_).f0)[25]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_71_globalState_.f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_71_globalState_.f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_72_v1_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_106_v31_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_106_v31_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_106_v31_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_106_v31_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_106_v31_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_106_v31_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_106_v31_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_106_v31_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_106_v31_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_106_v31_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_106_v31_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_107_v32_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_108_v33_) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_109_v34_).f9)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_109_v34_).f9)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_109_v34_).f9)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_109_v34_).f9)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_109_v34_).f9)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_109_v34_).f9)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_109_v34_).f9)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_109_v34_).f9)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_109_v34_).f9)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_109_v34_).f9)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_109_v34_).f9)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_110_v35_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_111_v36_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_112_v37_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_113_v38_).f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_113_v38_.f13).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_114_v39_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_115_v40_) == (_dafny.SeqWithoutIsStrInference([False, False, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_116_v41_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_153_v70_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_154_v71_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_162_v76_) == (_dafny.Set({714}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_163_i12_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_197_v102_) == (_dafny.Map({False: _dafny.CodePoint('s')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_198_v103_) == (_dafny.SeqWithoutIsStrInference([714]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_199_i15_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_212_v114_)[0]).f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_212_v114_)[0].f13).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_212_v114_)[0]).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_212_v114_)[0].f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_212_v114_)[0]).f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_212_v114_)[1]).f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_212_v114_)[1].f13).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_212_v114_)[1]).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_212_v114_)[1].f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_212_v114_)[1]).f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_212_v114_)[2]).f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_212_v114_)[2].f13).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_212_v114_)[2]).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_212_v114_)[2].f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_212_v114_)[2]).f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_212_v114_)[3]).f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_212_v114_)[3].f13).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_212_v114_)[3]).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_212_v114_)[3].f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_212_v114_)[3]).f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_213_v115_)[3])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_214_v116_).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_214_v116_.f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_215_v117_).f14))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_215_v117_).f15)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_216_v118_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_217_v119_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_218_v120_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_219_i18_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_223_v124_) == (_dafny.Map({False: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_224_v125_).cf68) == (_dafny.SeqWithoutIsStrInference([_dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({False: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False}), _dafny.Map({True: False})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(False, False, _dafny.CodePoint('D'), _dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D0_DC1)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D0_DC2)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D0_DC3)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D0_DC4)

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

class D0_DC3(D0, NamedTuple('DC3', [])):
    def __dafnystr__(self) -> str:
        return f'D0.DC3'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC3)
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC0(D0, NamedTuple('DC0', [('cf0', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC0({_dafny.string_of(self.cf0)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC0) and self.cf0 == __o.cf0
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC4(D0, NamedTuple('DC4', [('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC4({_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC4) and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()


class D1:
    @classmethod
    def default(cls, ):
        return lambda: D1_DC6(None, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D1_DC6)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D1_DC7)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D1_DC5)

class D1_DC6(D1, NamedTuple('DC6', [('cf9', Any), ('cf10', Any), ('cf11', Any), ('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC6({_dafny.string_of(self.cf9)}, {self.cf10.VerbatimString(True)}, {_dafny.string_of(self.cf11)}, {_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC6) and self.cf9 == __o.cf9 and self.cf10 == __o.cf10 and self.cf11 == __o.cf11 and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC7(D1, NamedTuple('DC7', [])):
    def __dafnystr__(self) -> str:
        return f'D1.DC7'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC7)
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC5(D1, NamedTuple('DC5', [('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D2_DC8)

class D2_DC8(D2, NamedTuple('DC8', [('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC8({_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC8) and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC10(int(0), _dafny.Array(None, 0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D3_DC10)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D3_DC9)

class D3_DC10(D3, NamedTuple('DC10', [('cf15', Any), ('cf16', Any), ('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC10({_dafny.string_of(self.cf15)}, {_dafny.string_of(self.cf16)}, {_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC10) and self.cf15 == __o.cf15 and self.cf16 == __o.cf16 and self.cf17 == __o.cf17
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
        return lambda: D4_DC12()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D4_DC12)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D4_DC11)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D4_DC13)

class D4_DC12(D4, NamedTuple('DC12', [])):
    def __dafnystr__(self) -> str:
        return f'D4.DC12'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC12)
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC11(D4, NamedTuple('DC11', [('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC11({_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC11) and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC13(D4, NamedTuple('DC13', [('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC13({_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC13) and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC15(D0.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D5_DC15)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D5_DC14)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D5_DC16)

class D5_DC15(D5, NamedTuple('DC15', [('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC15({_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC15) and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC14(D5, NamedTuple('DC14', [('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC14({_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC14) and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC16(D5, NamedTuple('DC16', [('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC16({_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC16) and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC18(False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D6_DC18)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D6_DC19)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D6_DC17)

class D6_DC18(D6, NamedTuple('DC18', [('cf24', Any), ('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC18({_dafny.string_of(self.cf24)}, {_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC18) and self.cf24 == __o.cf24 and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC19(D6, NamedTuple('DC19', [('cf26', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC19({_dafny.string_of(self.cf26)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC19) and self.cf26 == __o.cf26
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC17(D6, NamedTuple('DC17', [('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC17({_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC17) and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC21(_dafny.Array(None, 0), False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D7_DC21)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D7_DC20)

class D7_DC21(D7, NamedTuple('DC21', [('cf28', Any), ('cf29', Any), ('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC21({_dafny.string_of(self.cf28)}, {_dafny.string_of(self.cf29)}, {_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC21) and self.cf28 == __o.cf28 and self.cf29 == __o.cf29 and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC20(D7, NamedTuple('DC20', [('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC20({_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC20) and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC23(False, False, False, _dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D8_DC23)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D8_DC22)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D8_DC24)

class D8_DC23(D8, NamedTuple('DC23', [('cf32', Any), ('cf33', Any), ('cf34', Any), ('cf35', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC23({_dafny.string_of(self.cf32)}, {_dafny.string_of(self.cf33)}, {_dafny.string_of(self.cf34)}, {_dafny.string_of(self.cf35)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC23) and self.cf32 == __o.cf32 and self.cf33 == __o.cf33 and self.cf34 == __o.cf34 and self.cf35 == __o.cf35
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC22(D8, NamedTuple('DC22', [('cf31', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC22({_dafny.string_of(self.cf31)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC22) and self.cf31 == __o.cf31
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC24(D8, NamedTuple('DC24', [('cf36', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC24({_dafny.string_of(self.cf36)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC24) and self.cf36 == __o.cf36
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC26(int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D9_DC26)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D9_DC25)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D9_DC27)

class D9_DC26(D9, NamedTuple('DC26', [('cf38', Any), ('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC26({_dafny.string_of(self.cf38)}, {_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC26) and self.cf38 == __o.cf38 and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC25(D9, NamedTuple('DC25', [('cf37', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC25({_dafny.string_of(self.cf37)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC25) and self.cf37 == __o.cf37
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC27(D9, NamedTuple('DC27', [('cf40', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC27({_dafny.string_of(self.cf40)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC27) and self.cf40 == __o.cf40
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: D10_DC29(False, _dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D10_DC29)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D10_DC30)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D10_DC28)

class D10_DC29(D10, NamedTuple('DC29', [('cf42', Any), ('cf43', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC29({_dafny.string_of(self.cf42)}, {_dafny.string_of(self.cf43)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC29) and self.cf42 == __o.cf42 and self.cf43 == __o.cf43
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC30(D10, NamedTuple('DC30', [('cf44', Any), ('cf45', Any), ('cf46', Any), ('cf47', Any), ('cf48', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC30({_dafny.string_of(self.cf44)}, {_dafny.string_of(self.cf45)}, {_dafny.string_of(self.cf46)}, {_dafny.string_of(self.cf47)}, {_dafny.string_of(self.cf48)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC30) and self.cf44 == __o.cf44 and self.cf45 == __o.cf45 and self.cf46 == __o.cf46 and self.cf47 == __o.cf47 and self.cf48 == __o.cf48
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC28(D10, NamedTuple('DC28', [('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC28({_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC28) and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC32(False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D11_DC32)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D11_DC33)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D11_DC31)

class D11_DC32(D11, NamedTuple('DC32', [('cf50', Any), ('cf51', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC32({_dafny.string_of(self.cf50)}, {_dafny.string_of(self.cf51)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC32) and self.cf50 == __o.cf50 and self.cf51 == __o.cf51
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC33(D11, NamedTuple('DC33', [('cf52', Any), ('cf53', Any), ('cf54', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC33({_dafny.string_of(self.cf52)}, {_dafny.string_of(self.cf53)}, {_dafny.string_of(self.cf54)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC33) and self.cf52 == __o.cf52 and self.cf53 == __o.cf53 and self.cf54 == __o.cf54
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC31(D11, NamedTuple('DC31', [('cf49', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC31({_dafny.string_of(self.cf49)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC31) and self.cf49 == __o.cf49
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC35(_dafny.Array(None, 0), False, _dafny.Seq({}), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), int(0))
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

class D12_DC35(D12, NamedTuple('DC35', [('cf56', Any), ('cf57', Any), ('cf58', Any), ('cf59', Any), ('cf60', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC35({_dafny.string_of(self.cf56)}, {_dafny.string_of(self.cf57)}, {_dafny.string_of(self.cf58)}, {self.cf59.VerbatimString(True)}, {_dafny.string_of(self.cf60)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC35) and self.cf56 == __o.cf56 and self.cf57 == __o.cf57 and self.cf58 == __o.cf58 and self.cf59 == __o.cf59 and self.cf60 == __o.cf60
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC36(D12, NamedTuple('DC36', [('cf61', Any), ('cf62', Any), ('cf63', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC36({_dafny.string_of(self.cf61)}, {_dafny.string_of(self.cf62)}, {_dafny.string_of(self.cf63)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC36) and self.cf61 == __o.cf61 and self.cf62 == __o.cf62 and self.cf63 == __o.cf63
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC37(D12, NamedTuple('DC37', [('cf64', Any), ('cf65', Any), ('cf66', Any), ('cf67', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC37({_dafny.string_of(self.cf64)}, {_dafny.string_of(self.cf65)}, {_dafny.string_of(self.cf66)}, {_dafny.string_of(self.cf67)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC37) and self.cf64 == __o.cf64 and self.cf65 == __o.cf65 and self.cf66 == __o.cf66 and self.cf67 == __o.cf67
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC34(D12, NamedTuple('DC34', [('cf55', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC34({_dafny.string_of(self.cf55)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC34) and self.cf55 == __o.cf55
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: D13_DC39(_dafny.Map({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D13_DC39)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D13_DC40)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D13_DC38)
    @property
    def is_DC41(self) -> bool:
        return isinstance(self, D13_DC41)

class D13_DC39(D13, NamedTuple('DC39', [('cf69', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC39({_dafny.string_of(self.cf69)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC39) and self.cf69 == __o.cf69
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC40(D13, NamedTuple('DC40', [('cf70', Any), ('cf71', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC40({_dafny.string_of(self.cf70)}, {_dafny.string_of(self.cf71)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC40) and self.cf70 == __o.cf70 and self.cf71 == __o.cf71
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC38(D13, NamedTuple('DC38', [('cf68', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC38({_dafny.string_of(self.cf68)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC38) and self.cf68 == __o.cf68
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC41(D13, NamedTuple('DC41', [('cf72', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC41({_dafny.string_of(self.cf72)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC41) and self.cf72 == __o.cf72
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC42(self) -> bool:
        return isinstance(self, D14_DC42)

class D14_DC42(D14, NamedTuple('DC42', [('cf73', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC42({_dafny.string_of(self.cf73)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC42) and self.cf73 == __o.cf73
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: D15_DC44(None, _dafny.CodePoint('D'), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC44(self) -> bool:
        return isinstance(self, D15_DC44)
    @property
    def is_DC43(self) -> bool:
        return isinstance(self, D15_DC43)

class D15_DC44(D15, NamedTuple('DC44', [('cf75', Any), ('cf76', Any), ('cf77', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC44({_dafny.string_of(self.cf75)}, {_dafny.string_of(self.cf76)}, {_dafny.string_of(self.cf77)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC44) and self.cf75 == __o.cf75 and self.cf76 == __o.cf76 and self.cf77 == __o.cf77
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC43(D15, NamedTuple('DC43', [('cf74', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC43({_dafny.string_of(self.cf74)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC43) and self.cf74 == __o.cf74
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: D16_DC46(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.Set({}), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC46(self) -> bool:
        return isinstance(self, D16_DC46)
    @property
    def is_DC45(self) -> bool:
        return isinstance(self, D16_DC45)
    @property
    def is_DC47(self) -> bool:
        return isinstance(self, D16_DC47)

class D16_DC46(D16, NamedTuple('DC46', [('cf79', Any), ('cf80', Any), ('cf81', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC46({self.cf79.VerbatimString(True)}, {_dafny.string_of(self.cf80)}, {_dafny.string_of(self.cf81)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC46) and self.cf79 == __o.cf79 and self.cf80 == __o.cf80 and self.cf81 == __o.cf81
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC45(D16, NamedTuple('DC45', [('cf78', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC45({_dafny.string_of(self.cf78)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC45) and self.cf78 == __o.cf78
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC47(D16, NamedTuple('DC47', [('cf82', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC47({_dafny.string_of(self.cf82)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC47) and self.cf82 == __o.cf82
    def __hash__(self) -> int:
        return super().__hash__()


class D17:
    @classmethod
    def default(cls, ):
        return lambda: D17_DC49(int(0), int(0), None)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC49(self) -> bool:
        return isinstance(self, D17_DC49)
    @property
    def is_DC50(self) -> bool:
        return isinstance(self, D17_DC50)
    @property
    def is_DC48(self) -> bool:
        return isinstance(self, D17_DC48)

class D17_DC49(D17, NamedTuple('DC49', [('cf84', Any), ('cf85', Any), ('cf86', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC49({_dafny.string_of(self.cf84)}, {_dafny.string_of(self.cf85)}, {_dafny.string_of(self.cf86)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC49) and self.cf84 == __o.cf84 and self.cf85 == __o.cf85 and self.cf86 == __o.cf86
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC50(D17, NamedTuple('DC50', [('cf87', Any), ('cf88', Any), ('cf89', Any), ('cf90', Any), ('cf91', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC50({_dafny.string_of(self.cf87)}, {_dafny.string_of(self.cf88)}, {_dafny.string_of(self.cf89)}, {_dafny.string_of(self.cf90)}, {_dafny.string_of(self.cf91)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC50) and self.cf87 == __o.cf87 and self.cf88 == __o.cf88 and self.cf89 == __o.cf89 and self.cf90 == __o.cf90 and self.cf91 == __o.cf91
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC48(D17, NamedTuple('DC48', [('cf83', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC48({_dafny.string_of(self.cf83)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC48) and self.cf83 == __o.cf83
    def __hash__(self) -> int:
        return super().__hash__()


class D18:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC51(self) -> bool:
        return isinstance(self, D18_DC51)

class D18_DC51(D18, NamedTuple('DC51', [('cf92', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC51({_dafny.string_of(self.cf92)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC51) and self.cf92 == __o.cf92
    def __hash__(self) -> int:
        return super().__hash__()


class D19:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC52(self) -> bool:
        return isinstance(self, D19_DC52)

class D19_DC52(D19, NamedTuple('DC52', [('cf93', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC52({_dafny.string_of(self.cf93)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC52) and self.cf93 == __o.cf93
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

class T1:
    pass
    def m1(self, p0, globalState):
        pass


class GlobalState:
    def  __init__(self):
        self.f1: int = int(0)
        self.f2: bool = False
        self._f0: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2):
        (self)._f0 = f0
        (self).f1 = f1
        (self).f2 = f2

    @property
    def f0(self):
        return self._f0

class C0:
    def  __init__(self):
        self.f10: bool = False
        self._f11: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f10, f11):
        (self).f10 = f10
        (self)._f11 = f11

    def fm5(self, p0, p1, p2, p3, globalState):
        return (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cdnmimuc")))) != (len(_dafny.Set({not(self.f10)})))

    @property
    def f11(self):
        return self._f11

class C1(T1, T0):
    def  __init__(self):
        self._f4: bool = False
        self._f7: int = int(0)
        self._f3: int = int(0)
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
    def f7(self):
        return self._f7
    @property
    def f3(self):
        return self._f3
    def ctor__(self, f7, f3, f4):
        (self)._f7 = f7
        (self)._f3 = f3
        (self).f4 = f4

    def m1(self, p0, globalState):
        r0: int = int(0)
        if not(self.f4):
            d_330_v0_: _dafny.Array
            nw60_ = _dafny.Array(False, 19)
            d_330_v0_ = nw60_
            d_331_v1_: C0
            nw61_ = C0()
            nw61_.ctor__(not(default__.fm2((self).f7, globalState)), d_330_v0_)
            d_331_v1_ = nw61_
            d_332_v2_: C0
            nw62_ = C0()
            nw62_.ctor__(d_331_v1_.f10, (d_331_v1_).f11)
            d_332_v2_ = nw62_
            index46_ = default__.safeIndex(845, (d_330_v0_).length(0))
            (d_330_v0_)[index46_] = (d_332_v2_.f10 if self.f4 else self.f4)
            d_333_v3_: _dafny.Seq
            d_333_v3_ = _dafny.SeqWithoutIsStrInference([d_331_v1_.f10, d_331_v1_.f10])
            d_334_v4_: _dafny.Seq
            d_334_v4_ = _dafny.SeqWithoutIsStrInference([len(default__.fm19(p0, globalState))])
            d_335_v5_: _dafny.Seq
            d_335_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hxpnn"))
            d_336_v6_: D3
            d_336_v6_ = D3_DC10(len(d_333_v3_), (d_331_v1_).f11, (d_334_v4_)[default__.safeIndex(len((d_334_v4_).set(default__.safeIndex((0) - (len(d_335_v5_)), len(d_334_v4_)), p0)), len(d_334_v4_))])
            index47_ = default__.safeIndex(845, (d_330_v0_).length(0))
            rhs47_ = (d_336_v6_).cf17
            rhs48_ = not(self.f4)
            lhs34_ = globalState
            lhs35_ = d_330_v0_
            lhs36_ = default__.safeIndex(845, (d_330_v0_).length(0))
            lhs34_.f1 = rhs47_
            lhs35_[lhs36_] = rhs48_
            d_337_v7_: str
            d_337_v7_ = _dafny.CodePoint('j')
            d_337_v7_ = _dafny.CodePoint('c')
            (globalState).f1 = (self).f7
        elif True:
            d_338_v8_: _dafny.Seq
            d_338_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xpxqoyf"))
            d_339_v9_: _dafny.Array
            def lambda13_(d_340_i0_):
                return self.f4

            init7_ = lambda13_
            nw63_ = _dafny.Array(None, 20)
            for i0_7_ in range(nw63_.length(0)):
                nw63_[i0_7_] = init7_(i0_7_)
            d_339_v9_ = nw63_
            d_341_v10_: D3
            d_341_v10_ = D3_DC10(len(d_338_v8_), d_339_v9_, (self).f7)
            d_338_v8_ = default__.fm20(((d_341_v10_).cf15) * ((self).f3), globalState)
            (globalState).f1 = default__.safeModuloInt((0) - ((self).f3), len(_dafny.Map({(self).f7: self.f4})))
            if (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ukt"))) == (default__.fm20((0) - (len(d_338_v8_)), globalState)):
                (globalState).f1 = (self).f3
                d_342_v11_: _dafny.MultiSet
                d_342_v11_ = _dafny.MultiSet([self.f4])
                d_342_v11_ = d_342_v11_
                (globalState).f2 = ((p0) > ((0) - ((self).f7)) if self.f4 else False)
                (globalState).f1 = default__.safeModuloInt(791, len(_dafny.Map({(self).f3: self.f4})))
                d_341_v10_ = d_341_v10_
            elif True:
                (globalState).f2 = default__.fm2(default__.safeModuloInt(p0, (self).f3), globalState)
                d_343_v12_: D0
                d_343_v12_ = D0_DC0(self.f4)
                d_344_v13_: _dafny.MultiSet
                d_344_v13_ = _dafny.MultiSet([self.f4, self.f4, self.f4, (d_343_v12_).cf0, self.f4])
                d_345_v14_: _dafny.Map
                d_345_v14_ = _dafny.Map({d_338_v8_: d_344_v13_})
                d_345_v14_ = (d_345_v14_).set((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_346_i1_ in range(default__.abs(-476))])).set(default__.safeIndex((self).f7, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_347_i1_ in range(default__.abs(-476))]))), _dafny.CodePoint('q')), d_344_v13_)
                (globalState).f1 = (0) - ((self).f7)
                d_348_v15_: _dafny.Array
                nw64_ = _dafny.Array(None, 11)
                nw64_[int(0)] = p0
                nw64_[int(1)] = (self).f3
                nw64_[int(2)] = (self).f7
                nw64_[int(3)] = default__.fm0(globalState)
                nw64_[int(4)] = (self).f3
                nw64_[int(5)] = len(d_338_v8_)
                nw64_[int(6)] = p0
                nw64_[int(7)] = (0) - ((self).f7)
                nw64_[int(8)] = (self).f7
                nw64_[int(9)] = p0
                nw64_[int(10)] = -233
                d_348_v15_ = nw64_
                index48_ = default__.safeIndex(868, (d_348_v15_).length(0))
                (d_348_v15_)[index48_] = ((0) - ((self).f7)) * (p0)
                index49_ = default__.safeIndex(868, (d_348_v15_).length(0))
                (d_348_v15_)[index49_] = default__.fm0(globalState)
                (globalState).f2 = True
            d_349_v16_: _dafny.Set
            d_349_v16_ = _dafny.Set({p0})
            r0 = len(d_349_v16_)
            d_339_v9_ = d_339_v9_
        hi3_ = default__.safeModuloInt(338, (self).f3)
        for d_350_i2_ in range(len(_dafny.SeqWithoutIsStrInference([(self).f3 for d_363_i3_ in range(default__.abs(40))])), hi3_):
            d_351_v17_: _dafny.Array
            def lambda14_(d_352_i4_):
                return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kjyvte"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ydnlt")))

            init8_ = lambda14_
            nw65_ = _dafny.Array(None, 20)
            for i0_8_ in range(nw65_.length(0)):
                nw65_[i0_8_] = init8_(i0_8_)
            d_351_v17_ = nw65_
            d_353_v18_: _dafny.Seq
            d_353_v18_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wjj"))
            index50_ = default__.safeIndex(329, (d_351_v17_).length(0))
            (d_351_v17_)[index50_] = d_353_v18_
            index51_ = default__.safeIndex(329, (d_351_v17_).length(0))
            (d_351_v17_)[index51_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hyb"))
            d_354_v19_: _dafny.Array
            def lambda15_(d_355_i5_):
                return (d_355_i5_) * (len(_dafny.SeqWithoutIsStrInference([self.f4])))

            init9_ = lambda15_
            nw66_ = _dafny.Array(None, 20)
            for i0_9_ in range(nw66_.length(0)):
                nw66_[i0_9_] = init9_(i0_9_)
            d_354_v19_ = nw66_
            index52_ = default__.safeIndex(812, (d_354_v19_).length(0))
            (d_354_v19_)[index52_] = d_350_i2_
            index53_ = default__.safeIndex(812, (d_354_v19_).length(0))
            (d_354_v19_)[index53_] = (self).f3
            r0 = default__.fm0(globalState)
            d_356_v20_: _dafny.Array
            nw67_ = _dafny.Array(_dafny.CodePoint('D'), 17)
            d_356_v20_ = nw67_
            d_357_v21_: str
            d_357_v21_ = _dafny.CodePoint('u')
            index54_ = default__.safeIndex(584, (d_356_v20_).length(0))
            (d_356_v20_)[index54_] = d_357_v21_
            d_358_v22_: _dafny.Map
            d_358_v22_ = _dafny.Map({self.f4: self.f4})
            d_359_v23_: _dafny.Seq
            d_359_v23_ = _dafny.SeqWithoutIsStrInference([d_358_v22_])
            d_360_v24_: _dafny.Seq
            d_360_v24_ = _dafny.SeqWithoutIsStrInference([default__.fm21(d_359_v23_, globalState)])
            d_361_v25_: _dafny.Seq
            d_361_v25_ = _dafny.SeqWithoutIsStrInference([self.f4])
            d_362_v26_: _dafny.Seq
            d_362_v26_ = _dafny.SeqWithoutIsStrInference([len(d_361_v25_)])
            index55_ = default__.safeIndex(584, (d_356_v20_).length(0))
            index56_ = default__.safeIndex(812, (d_354_v19_).length(0))
            rhs49_ = self.f4
            rhs50_ = d_357_v21_
            rhs51_ = len((d_360_v24_)[default__.safeIndex(default__.safeDivisionInt((self).f7, len(d_362_v26_)), len(d_360_v24_))])
            lhs37_ = globalState
            lhs38_ = d_356_v20_
            lhs39_ = default__.safeIndex(584, (d_356_v20_).length(0))
            lhs40_ = d_354_v19_
            lhs41_ = default__.safeIndex(812, (d_354_v19_).length(0))
            lhs37_.f2 = rhs49_
            lhs38_[lhs39_] = rhs50_
            lhs40_[lhs41_] = rhs51_
        d_364_v27_: _dafny.Seq
        d_364_v27_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s"))
        d_364_v27_ = (d_364_v27_) + ((d_364_v27_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "asujtxoi"))))
        d_365_v28_: _dafny.Map
        d_365_v28_ = _dafny.Map({len(d_364_v27_): 235})
        d_366_v29_: _dafny.Seq
        d_366_v29_ = _dafny.SeqWithoutIsStrInference([d_365_v28_, _dafny.Map({(self).f3: len(d_364_v27_)}), d_365_v28_])
        hi4_ = ((self).f3) + (len(d_366_v29_))
        for d_367_i6_ in range(default__.safeDivisionInt((self).f7, 357), hi4_):
            d_368_v30_: _dafny.MultiSet
            d_368_v30_ = _dafny.MultiSet([(0) - ((0) - (p0))])
            (self).f4 = (self.f4 if (d_368_v30_).issubset(d_368_v30_) else ((self).f7) < (default__.fm0(globalState)))
            if default__.fm2((self).f3, globalState):
                d_369_v31_: _dafny.Array
                def lambda16_(d_370_i7_):
                    return self.f4

                init10_ = lambda16_
                nw68_ = _dafny.Array(None, 27)
                for i0_10_ in range(nw68_.length(0)):
                    nw68_[i0_10_] = init10_(i0_10_)
                d_369_v31_ = nw68_
                d_371_v32_: _dafny.Array
                nw69_ = _dafny.Array(None, 11)
                nw69_[int(0)] = d_369_v31_
                nw69_[int(1)] = d_369_v31_
                nw69_[int(2)] = d_369_v31_
                nw69_[int(3)] = d_369_v31_
                nw69_[int(4)] = d_369_v31_
                nw69_[int(5)] = d_369_v31_
                nw69_[int(6)] = d_369_v31_
                nw69_[int(7)] = (d_369_v31_ if self.f4 else d_369_v31_)
                nw69_[int(8)] = d_369_v31_
                nw69_[int(9)] = d_369_v31_
                nw69_[int(10)] = d_369_v31_
                d_371_v32_ = nw69_
                d_372_v33_: _dafny.Array
                nw70_ = _dafny.Array(_dafny.CodePoint('D'), 22)
                d_372_v33_ = nw70_
                d_373_v34_: str
                d_373_v34_ = _dafny.CodePoint('r')
                index57_ = default__.safeIndex(463, (d_372_v33_).length(0))
                (d_372_v33_)[index57_] = d_373_v34_
                index58_ = default__.safeIndex(463, (d_372_v33_).length(0))
                rhs52_ = d_371_v32_
                rhs53_ = default__.safeModuloInt(default__.fm0(globalState), d_367_i6_)
                rhs54_ = self.f4
                rhs55_ = not(self.f4)
                rhs56_ = d_373_v34_
                lhs42_ = globalState
                lhs43_ = self
                lhs44_ = self
                lhs45_ = d_372_v33_
                lhs46_ = default__.safeIndex(463, (d_372_v33_).length(0))
                d_371_v32_ = rhs52_
                lhs42_.f1 = rhs53_
                lhs43_.f4 = rhs54_
                lhs44_.f4 = rhs55_
                lhs45_[lhs46_] = rhs56_
                d_374_v35_: _dafny.Seq
                d_374_v35_ = _dafny.SeqWithoutIsStrInference([d_364_v27_])
                d_375_v36_: _dafny.Map
                d_375_v36_ = _dafny.Map({(self).f7: d_374_v35_})
                d_375_v36_ = (d_375_v36_).set((0) - (p0), (d_374_v35_) + (d_374_v35_))
                d_376_v37_: _dafny.Array
                def lambda17_(d_377_i6_):
                    def lambda18_(d_378_i8_):
                        return (d_378_i8_) - (d_377_i6_)

                    return lambda18_

                init11_ = lambda17_(d_367_i6_)
                nw71_ = _dafny.Array(None, 26)
                for i0_11_ in range(nw71_.length(0)):
                    nw71_[i0_11_] = init11_(i0_11_)
                d_376_v37_ = nw71_
                d_376_v37_ = d_376_v37_
                (globalState).f1 = (self).f3
                d_379_v38_: _dafny.Seq
                d_379_v38_ = _dafny.SeqWithoutIsStrInference([self.f4, self.f4, self.f4])
                d_380_v39_: _dafny.MultiSet
                d_380_v39_ = _dafny.MultiSet([(d_372_v33_)[default__.safeIndex(463, (d_372_v33_).length(0))]])
                d_381_v40_: D0
                d_381_v40_ = D0_DC2(d_380_v39_, self.f4)
                d_382_v41_: D0
                d_382_v41_ = D0_DC2((d_381_v40_).cf5, self.f4)
                d_383_v42_: C0
                nw72_ = C0()
                nw72_.ctor__(not((d_381_v40_).cf6), d_369_v31_)
                d_383_v42_ = nw72_
                d_384_v43_: D6
                d_384_v43_ = D6_DC17(d_379_v38_)
                rhs57_ = (self).f3
                rhs58_ = ((d_384_v43_).cf23) + (d_379_v38_)
                rhs59_ = self.f4
                rhs60_ = p0
                rhs61_ = d_383_v42_
                lhs47_ = globalState
                lhs48_ = self
                lhs49_ = globalState
                lhs47_.f1 = rhs57_
                d_379_v38_ = rhs58_
                lhs48_.f4 = rhs59_
                lhs49_.f1 = rhs60_
                d_383_v42_ = rhs61_
            elif True:
                d_385_v44_: _dafny.Array
                nw73_ = _dafny.Array(_dafny.Array(None, 0), 14)
                d_385_v44_ = nw73_
                d_385_v44_ = d_385_v44_
                (self).f4 = (d_364_v27_) <= (d_364_v27_)
                d_386_v45_: str
                d_386_v45_ = _dafny.CodePoint('o')
                d_386_v45_ = d_386_v45_
                (globalState).f2 = default__.fm2(-802, globalState)
                d_387_v46_: _dafny.Array
                nw74_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 6)
                d_387_v46_ = nw74_
                index59_ = default__.safeIndex(612, (d_387_v46_).length(0))
                (d_387_v46_)[index59_] = (d_364_v27_) + (_dafny.SeqWithoutIsStrInference([d_386_v45_ for d_388_i9_ in range(default__.abs(940))]))
                d_389_v47_: D1
                d_389_v47_ = D1_DC5(d_367_i6_)
                d_390_v48_: _dafny.Map
                d_390_v48_ = _dafny.Map({d_389_v47_: 220})
                d_391_v49_: _dafny.Map
                d_391_v49_ = _dafny.Map({self.f4: self.f4})
                d_392_v50_: _dafny.Map
                d_392_v50_ = _dafny.Map({d_391_v49_: d_367_i6_})
                index60_ = default__.safeIndex(612, (d_387_v46_).length(0))
                rhs62_ = ((d_390_v48_)[d_389_v47_] if (d_389_v47_) in (d_390_v48_) else len(d_392_v50_))
                rhs63_ = default__.fm20((self).f3, globalState)
                lhs50_ = globalState
                lhs51_ = d_387_v46_
                lhs52_ = default__.safeIndex(612, (d_387_v46_).length(0))
                lhs50_.f1 = rhs62_
                lhs51_[lhs52_] = rhs63_
            (globalState).f2 = self.f4
            (globalState).f2 = True
        d_393_i10_: int
        d_393_i10_ = 0
        with _dafny.label("3"):
            while not (self.f4) or (self.f4):
                with _dafny.c_label("3"):
                    if (d_393_i10_) >= (100):
                        raise _dafny.Break("3")
                    d_393_i10_ = (d_393_i10_) + (1)
                    d_394_v51_: _dafny.Array
                    nw75_ = _dafny.Array(None, 29)
                    nw75_[int(0)] = self.f4
                    nw75_[int(1)] = self.f4
                    nw75_[int(2)] = self.f4
                    nw75_[int(3)] = self.f4
                    nw75_[int(4)] = True
                    nw75_[int(5)] = self.f4
                    nw75_[int(6)] = self.f4
                    nw75_[int(7)] = self.f4
                    nw75_[int(8)] = self.f4
                    nw75_[int(9)] = self.f4
                    nw75_[int(10)] = self.f4
                    nw75_[int(11)] = self.f4
                    nw75_[int(12)] = self.f4
                    nw75_[int(13)] = self.f4
                    nw75_[int(14)] = self.f4
                    nw75_[int(15)] = True
                    nw75_[int(16)] = self.f4
                    nw75_[int(17)] = self.f4
                    nw75_[int(18)] = self.f4
                    nw75_[int(19)] = self.f4
                    nw75_[int(20)] = self.f4
                    nw75_[int(21)] = True
                    nw75_[int(22)] = default__.fm2((0) - ((0) - (len(d_364_v27_))), globalState)
                    nw75_[int(23)] = self.f4
                    nw75_[int(24)] = self.f4
                    nw75_[int(25)] = self.f4
                    nw75_[int(26)] = default__.fm2(p0, globalState)
                    nw75_[int(27)] = self.f4
                    nw75_[int(28)] = True
                    d_394_v51_ = nw75_
                    d_395_v52_: C0
                    nw76_ = C0()
                    nw76_.ctor__(self.f4, d_394_v51_)
                    d_395_v52_ = nw76_
                    d_396_v53_: _dafny.Set
                    d_396_v53_ = _dafny.Set({p0, (self).f7, (self).f3})
                    d_397_v54_: _dafny.Map
                    d_397_v54_ = _dafny.Map({((self).f7) < (p0): d_364_v27_})
                    d_398_v55_: str
                    d_398_v55_ = _dafny.CodePoint('f')
                    rhs64_ = ((d_397_v54_)[((d_395_v52_).fm5(p0, (self).f7, d_398_v55_, d_395_v52_.f10, globalState) if default__.fm2((self).f7, globalState) else self.f4)] if (((d_395_v52_).fm5(p0, (self).f7, d_398_v55_, d_395_v52_.f10, globalState) if default__.fm2((self).f7, globalState) else self.f4)) in (d_397_v54_) else d_364_v27_)
                    rhs65_ = d_396_v53_
                    rhs66_ = ((d_366_v29_)[default__.safeIndex(p0, len(d_366_v29_))]) | (d_365_v28_)
                    rhs67_ = (d_394_v51_ if self.f4 else d_394_v51_)
                    d_364_v27_ = rhs64_
                    d_396_v53_ = rhs65_
                    d_365_v28_ = rhs66_
                    d_394_v51_ = rhs67_
                    d_399_v56_: _dafny.Seq
                    d_399_v56_ = _dafny.SeqWithoutIsStrInference([d_395_v52_.f10, self.f4])
                    d_400_v57_: _dafny.Array
                    nw77_ = _dafny.Array(None, 4)
                    nw77_[int(0)] = d_399_v56_
                    nw77_[int(1)] = (d_399_v56_) + (_dafny.SeqWithoutIsStrInference([d_395_v52_.f10]))
                    nw77_[int(2)] = d_399_v56_
                    nw77_[int(3)] = d_399_v56_
                    d_400_v57_ = nw77_
                    index61_ = default__.safeIndex(598, (d_400_v57_).length(0))
                    (d_400_v57_)[index61_] = d_399_v56_
                    d_401_v58_: _dafny.Array
                    nw78_ = _dafny.Array(int(0), 27)
                    d_401_v58_ = nw78_
                    d_402_v59_: _dafny.Set
                    d_402_v59_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vfxj"))})
                    index62_ = default__.safeIndex(910, (d_401_v58_).length(0))
                    (d_401_v58_)[index62_] = ((self).f3) * (len(d_402_v59_))
                    d_403_v60_: _dafny.Seq
                    d_403_v60_ = _dafny.SeqWithoutIsStrInference([d_396_v53_, d_396_v53_])
                    d_404_v61_: _dafny.Map
                    d_404_v61_ = _dafny.Map({d_394_v51_: d_398_v55_})
                    index63_ = default__.safeIndex(598, (d_400_v57_).length(0))
                    index64_ = default__.safeIndex(910, (d_401_v58_).length(0))
                    rhs68_ = d_399_v56_
                    rhs69_ = len((d_403_v60_) + (d_403_v60_))
                    rhs70_ = len(d_404_v61_)
                    rhs71_ = (_dafny.Set({p0, p0})).intersection(d_396_v53_)
                    lhs53_ = d_400_v57_
                    lhs54_ = default__.safeIndex(598, (d_400_v57_).length(0))
                    lhs55_ = globalState
                    lhs56_ = d_401_v58_
                    lhs57_ = default__.safeIndex(910, (d_401_v58_).length(0))
                    lhs53_[lhs54_] = rhs68_
                    lhs55_.f1 = rhs69_
                    lhs56_[lhs57_] = rhs70_
                    d_396_v53_ = rhs71_
                    (globalState).f1 = (self).f3
                    pass
            pass
        d_405_v62_: _dafny.Set
        d_405_v62_ = _dafny.Set({not(self.f4)})
        d_406_i11_: int
        d_406_i11_ = 0
        with _dafny.label("4"):
            while (d_405_v62_).ispropersubset((default__.fm22(self.f4, self.f4, globalState)).intersection(_dafny.Set({self.f4}))):
                with _dafny.c_label("4"):
                    if (d_406_i11_) >= (100):
                        raise _dafny.Break("4")
                    d_406_i11_ = (d_406_i11_) + (1)
                    d_407_v63_: _dafny.Array
                    nw79_ = _dafny.Array(int(0), 6)
                    d_407_v63_ = nw79_
                    d_408_v64_: _dafny.Array
                    nw80_ = _dafny.Array(None, 11)
                    nw80_[int(0)] = d_407_v63_
                    nw80_[int(1)] = d_407_v63_
                    nw80_[int(2)] = d_407_v63_
                    nw80_[int(3)] = d_407_v63_
                    nw80_[int(4)] = d_407_v63_
                    nw80_[int(5)] = d_407_v63_
                    nw80_[int(6)] = d_407_v63_
                    nw80_[int(7)] = d_407_v63_
                    nw80_[int(8)] = d_407_v63_
                    nw80_[int(9)] = (d_407_v63_ if self.f4 else d_407_v63_)
                    nw80_[int(10)] = d_407_v63_
                    d_408_v64_ = nw80_
                    index65_ = default__.safeIndex(441, (d_408_v64_).length(0))
                    (d_408_v64_)[index65_] = d_407_v63_
                    d_409_v65_: _dafny.Array
                    def lambda19_(d_410_v27_):
                        def lambda20_(d_411_i12_):
                            return _dafny.Map({(self).f7: d_410_v27_})

                        return lambda20_

                    init12_ = lambda19_(d_364_v27_)
                    nw81_ = _dafny.Array(None, 1)
                    for i0_12_ in range(nw81_.length(0)):
                        nw81_[i0_12_] = init12_(i0_12_)
                    d_409_v65_ = nw81_
                    d_412_v66_: D7
                    d_412_v66_ = D7_DC20(_dafny.Map({(self).f3: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tfjfsyru"))}))
                    index66_ = default__.safeIndex(196, (d_409_v65_).length(0))
                    (d_409_v65_)[index66_] = (d_412_v66_).cf27
                    d_413_v67_: str
                    d_413_v67_ = _dafny.CodePoint('y')
                    d_414_v68_: _dafny.Map
                    d_414_v68_ = _dafny.Map({(0) - (p0): d_364_v27_})
                    index67_ = default__.safeIndex(441, (d_408_v64_).length(0))
                    index68_ = default__.safeIndex(196, (d_409_v65_).length(0))
                    rhs72_ = d_407_v63_
                    rhs73_ = d_414_v68_
                    rhs74_ = (self).f3
                    rhs75_ = d_413_v67_
                    lhs58_ = d_408_v64_
                    lhs59_ = default__.safeIndex(441, (d_408_v64_).length(0))
                    lhs60_ = d_409_v65_
                    lhs61_ = default__.safeIndex(196, (d_409_v65_).length(0))
                    lhs62_ = globalState
                    lhs58_[lhs59_] = rhs72_
                    lhs60_[lhs61_] = rhs73_
                    lhs62_.f1 = rhs74_
                    d_413_v67_ = rhs75_
                    r0 = default__.fm0(globalState)
                    d_415_v69_: _dafny.Array
                    def lambda21_(d_416_i13_):
                        return self.f4

                    init13_ = lambda21_
                    nw82_ = _dafny.Array(None, 16)
                    for i0_13_ in range(nw82_.length(0)):
                        nw82_[i0_13_] = init13_(i0_13_)
                    d_415_v69_ = nw82_
                    d_417_v70_: _dafny.Map
                    d_417_v70_ = _dafny.Map({self.f4: self.f4})
                    index69_ = default__.safeIndex(813, (d_415_v69_).length(0))
                    (d_415_v69_)[index69_] = ((d_417_v70_)[self.f4] if (self.f4) in (d_417_v70_) else self.f4)
                    index70_ = default__.safeIndex(813, (d_415_v69_).length(0))
                    (d_415_v69_)[index70_] = (d_413_v67_) not in (_dafny.SeqWithoutIsStrInference([d_413_v67_ for d_418_i14_ in range(default__.abs(-39))]))
                    d_419_v71_: D0
                    d_419_v71_ = D0_DC3()
                    source5_ = d_419_v71_
                    if source5_.is_DC1:
                        d_420___mcc_h0_ = source5_.cf1
                        d_421___mcc_h1_ = source5_.cf2
                        d_422___mcc_h2_ = source5_.cf3
                        d_423___mcc_h3_ = source5_.cf4
                        d_424_cf4_ = d_423___mcc_h3_
                        d_425_cf3_ = d_422___mcc_h2_
                        d_426_cf2_ = d_421___mcc_h1_
                        d_427_cf1_ = d_420___mcc_h0_
                        index71_ = default__.safeIndex(193, (d_407_v63_).length(0))
                        (d_407_v63_)[index71_] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cnff")))
                        index72_ = default__.safeIndex(193, (d_407_v63_).length(0))
                        index73_ = default__.safeIndex(813, (d_415_v69_).length(0))
                        rhs76_ = len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_428_i15_ in range(default__.abs(909))])) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hmi"))).set(default__.safeIndex((self).f7, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hmi")))), d_413_v67_)))
                        rhs77_ = (d_415_v69_)[default__.safeIndex(813, (d_415_v69_).length(0))]
                        rhs78_ = (d_415_v69_)[default__.safeIndex(813, (d_415_v69_).length(0))]
                        lhs63_ = d_407_v63_
                        lhs64_ = default__.safeIndex(193, (d_407_v63_).length(0))
                        lhs65_ = d_415_v69_
                        lhs66_ = default__.safeIndex(813, (d_415_v69_).length(0))
                        lhs67_ = globalState
                        lhs63_[lhs64_] = rhs76_
                        lhs65_[lhs66_] = rhs77_
                        lhs67_.f2 = rhs78_
                        d_429_v72_: _dafny.Seq
                        d_429_v72_ = _dafny.SeqWithoutIsStrInference([((0) - (((d_365_v28_)[p0] if (p0) in (d_365_v28_) else p0))) <= ((d_407_v63_)[default__.safeIndex(193, (d_407_v63_).length(0))]), not(d_426_cf2_)])
                        index74_ = default__.safeIndex(813, (d_415_v69_).length(0))
                        (d_415_v69_)[index74_] = (d_429_v72_)[default__.safeIndex((0) - ((p0) - ((self).f3)), len(d_429_v72_))]
                        d_430_v73_: D6
                        d_430_v73_ = D6_DC17(default__.fm1(d_413_v67_, d_426_cf2_, globalState))
                        d_431_v74_: D5
                        d_431_v74_ = D5_DC14(_dafny.Map({d_426_cf2_: _dafny.MultiSet((d_430_v73_).cf23)}))
                        d_431_v74_ = d_431_v74_
                        index75_ = default__.safeIndex(813, (d_415_v69_).length(0))
                        (d_415_v69_)[index75_] = False
                    elif source5_.is_DC2:
                        d_432___mcc_h4_ = source5_.cf5
                        d_433___mcc_h5_ = source5_.cf6
                        d_434_cf6_ = d_433___mcc_h5_
                        d_435_cf5_ = d_432___mcc_h4_
                        index76_ = default__.safeIndex(813, (d_415_v69_).length(0))
                        (d_415_v69_)[index76_] = not((len(d_364_v27_)) < ((self).f3))
                        (globalState).f2 = not ((d_415_v69_)[default__.safeIndex(813, (d_415_v69_).length(0))]) or (d_434_cf6_)
                        d_364_v27_ = d_364_v27_
                        d_436_v75_: _dafny.Seq
                        d_436_v75_ = _dafny.SeqWithoutIsStrInference([(self.f4 if (d_415_v69_)[default__.safeIndex(813, (d_415_v69_).length(0))] else (d_415_v69_)[default__.safeIndex(813, (d_415_v69_).length(0))])])
                        d_437_v76_: _dafny.Seq
                        d_437_v76_ = _dafny.SeqWithoutIsStrInference([d_436_v75_, d_436_v75_, _dafny.SeqWithoutIsStrInference([self.f4]), d_436_v75_])
                        d_438_v77_: _dafny.MultiSet
                        d_438_v77_ = _dafny.MultiSet([(self).f3, len(d_436_v75_)])
                        d_436_v75_ = (d_437_v76_)[default__.safeIndex((0) - (default__.safeModuloInt(((d_438_v77_)[(self).f7] if ((self).f7) in (d_438_v77_) else p0), (self).f3)), len(d_437_v76_))]
                    elif source5_.is_DC3:
                        d_439_v78_: _dafny.Array
                        nw83_ = _dafny.Array(_dafny.Array(None, 0), 27)
                        d_439_v78_ = nw83_
                        d_440_v79_: _dafny.Array
                        nw84_ = _dafny.Array(_dafny.Map({}), 11)
                        d_440_v79_ = nw84_
                        index77_ = default__.safeIndex(158, (d_439_v78_).length(0))
                        (d_439_v78_)[index77_] = d_440_v79_
                        index78_ = default__.safeIndex(158, (d_439_v78_).length(0))
                        index79_ = default__.safeIndex(813, (d_415_v69_).length(0))
                        rhs79_ = d_440_v79_
                        rhs80_ = d_413_v67_
                        rhs81_ = ((d_405_v62_).isdisjoint(d_405_v62_)) and (self.f4)
                        lhs68_ = d_439_v78_
                        lhs69_ = default__.safeIndex(158, (d_439_v78_).length(0))
                        lhs70_ = d_415_v69_
                        lhs71_ = default__.safeIndex(813, (d_415_v69_).length(0))
                        lhs68_[lhs69_] = rhs79_
                        d_413_v67_ = rhs80_
                        lhs70_[lhs71_] = rhs81_
                        d_441_v80_: _dafny.Map
                        d_441_v80_ = _dafny.Map({d_364_v27_: d_415_v69_})
                        d_441_v80_ = (d_441_v80_).set((d_364_v27_) + (d_364_v27_), d_415_v69_)
                        d_442_v81_: _dafny.Seq
                        d_442_v81_ = _dafny.SeqWithoutIsStrInference([560])
                        d_443_v82_: _dafny.Map
                        d_443_v82_ = _dafny.Map({len(d_442_v81_): not ((d_415_v69_)[default__.safeIndex(813, (d_415_v69_).length(0))]) or ((d_415_v69_)[default__.safeIndex(813, (d_415_v69_).length(0))])})
                        d_443_v82_ = (d_443_v82_).set((self).f3, (d_415_v69_)[default__.safeIndex(813, (d_415_v69_).length(0))])
                        d_444_v83_: _dafny.MultiSet
                        d_444_v83_ = _dafny.MultiSet([(d_415_v69_)[default__.safeIndex(813, (d_415_v69_).length(0))], (p0) != (p0), (d_415_v69_)[default__.safeIndex(813, (d_415_v69_).length(0))]])
                        d_444_v83_ = d_444_v83_
                    elif source5_.is_DC0:
                        d_445___mcc_h6_ = source5_.cf0
                        d_446_cf0_ = d_445___mcc_h6_
                        d_364_v27_ = (d_364_v27_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uclqne")))
                        d_447_v84_: _dafny.Array
                        d_447_v84_ = d_415_v69_
                        d_448_v85_: C0
                        nw85_ = C0()
                        nw85_.ctor__(default__.fm2(p0, globalState), (d_447_v84_))
                        d_448_v85_ = nw85_
                        d_449_v86_: _dafny.Map
                        d_449_v86_ = _dafny.Map({d_448_v85_: (d_415_v69_)[default__.safeIndex(813, (d_415_v69_).length(0))]})
                        d_449_v86_ = (d_449_v86_).set(d_448_v85_, False)
                        d_450_v87_: _dafny.MultiSet
                        d_450_v87_ = _dafny.MultiSet([d_413_v67_])
                        d_451_v88_: D0
                        d_451_v88_ = D0_DC2(d_450_v87_, (d_415_v69_)[default__.safeIndex(813, (d_415_v69_).length(0))])
                        (d_448_v85_).f10 = ((d_451_v88_ if d_446_cf0_ else d_451_v88_)).cf6
                        d_452_v89_: _dafny.MultiSet
                        d_452_v89_ = _dafny.MultiSet([False, default__.fm2((self).f7, globalState), d_448_v85_.f10])
                        d_453_v90_: _dafny.Seq
                        d_453_v90_ = _dafny.SeqWithoutIsStrInference([d_452_v89_, d_452_v89_, (d_452_v89_) | (d_452_v89_), ((d_452_v89_).set(self.f4, default__.abs((0) - ((self).f7)))) - (d_452_v89_), (_dafny.MultiSet([d_448_v85_.f10])) - (d_452_v89_)])
                        d_454_v91_: _dafny.MultiSet
                        d_454_v91_ = _dafny.MultiSet([p0])
                        d_453_v90_ = (d_453_v90_).set(default__.safeIndex(((d_454_v91_)[p0] if (p0) in (d_454_v91_) else (self).f3), len(d_453_v90_)), d_452_v89_)
                    elif True:
                        d_455___mcc_h7_ = source5_.cf7
                        d_456_cf7_ = d_455___mcc_h7_
                        d_457_v92_: _dafny.Seq
                        d_457_v92_ = _dafny.SeqWithoutIsStrInference([not(False)])
                        (globalState).f2 = (self.f4) or (((d_415_v69_)[default__.safeIndex(813, (d_415_v69_).length(0))]) not in (d_457_v92_))
                        d_458_v93_: _dafny.Seq
                        d_458_v93_ = _dafny.SeqWithoutIsStrInference([len(_dafny.Set({(self).f7, (self).f7}))])
                        index80_ = default__.safeIndex(813, (d_415_v69_).length(0))
                        (d_415_v69_)[index80_] = (563) < ((d_458_v93_)[default__.safeIndex((self).f7, len(d_458_v93_))])
                        (globalState).f2 = (d_415_v69_)[default__.safeIndex(813, (d_415_v69_).length(0))]
                        index81_ = default__.safeIndex(813, (d_415_v69_).length(0))
                        (d_415_v69_)[index81_] = self.f4
                    pass
            pass
        r0 = (self).f7
        return r0


class C2(T0):
    def  __init__(self):
        self._f4: bool = False
        self._f3: int = int(0)
        self._f14: int = int(0)
        self._f15: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    @property
    def f4(self):
        return self._f4
    @f4.setter
    def f4(self, value):
        self._f4 = value
    @property
    def f3(self):
        return self._f3
    def ctor__(self, f14, f15, f3, f4):
        (self)._f14 = f14
        (self)._f15 = f15
        (self)._f3 = f3
        (self).f4 = f4

    def m6(self, p0, p1, p2, globalState):
        if self.f4:
            (self).f4 = self.f4
            (globalState).f2 = (p2) >= (p1)
            if True:
                d_459_v0_: _dafny.Array
                def lambda22_(d_460_i0_):
                    return (D0_DC0(False)).cf0

                init14_ = lambda22_
                nw86_ = _dafny.Array(None, 27)
                for i0_14_ in range(nw86_.length(0)):
                    nw86_[i0_14_] = init14_(i0_14_)
                d_459_v0_ = nw86_
                d_461_v1_: C0
                nw87_ = C0()
                nw87_.ctor__(self.f4, d_459_v0_)
                d_461_v1_ = nw87_
                d_462_v2_: C0
                nw88_ = C0()
                nw88_.ctor__(((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ajxdmcbun"))))) == ((self).f3), d_459_v0_)
                d_462_v2_ = nw88_
                index82_ = default__.safeIndex(460, ((self).f15).length(0))
                ((self).f15)[index82_] = p1
                index83_ = default__.safeIndex(460, ((self).f15).length(0))
                ((self).f15)[index83_] = (self).f3
                d_463_v3_: _dafny.Seq
                d_463_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "di"))
                d_464_v4_: _dafny.Map
                d_464_v4_ = _dafny.Map({True: (d_463_v3_) + (d_463_v3_)})
                d_464_v4_ = (d_464_v4_).set(self.f4, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oeaya")))
                d_465_v5_: _dafny.Seq
                d_465_v5_ = _dafny.SeqWithoutIsStrInference([self.f4, d_461_v1_.f10, d_462_v2_.f10])
                d_466_v6_: _dafny.Map
                d_466_v6_ = _dafny.Map({(_dafny.MultiSet(d_465_v5_)).cardinality: (p2) + ((self).f14)})
                d_467_v7_: _dafny.MultiSet
                d_467_v7_ = _dafny.MultiSet([(self).f14])
                index84_ = default__.safeIndex(460, ((self).f15).length(0))
                rhs82_ = ((d_466_v6_)[(435 if d_461_v1_.f10 else (0) - ((self).f14))] if ((435 if d_461_v1_.f10 else (0) - ((self).f14))) in (d_466_v6_) else p1)
                rhs83_ = default__.safeDivisionInt((self).f14, ((0) - (p2)) * (p2))
                rhs84_ = ((d_467_v7_).intersection(d_467_v7_)).ispropersubset(d_467_v7_)
                lhs72_ = (self).f15
                lhs73_ = default__.safeIndex(460, ((self).f15).length(0))
                lhs74_ = globalState
                lhs75_ = d_462_v2_
                lhs72_[lhs73_] = rhs82_
                lhs74_.f1 = rhs83_
                lhs75_.f10 = rhs84_
            elif True:
                d_468_v8_: _dafny.Seq
                d_468_v8_ = _dafny.SeqWithoutIsStrInference([-723, (self).f14, (self).f3, p2, (self).f3])
                (globalState).f1 = (d_468_v8_)[default__.safeIndex(p2, len(d_468_v8_))]
                (globalState).f1 = p1
                d_469_v9_: _dafny.Array
                nw89_ = _dafny.Array(False, 2)
                d_469_v9_ = nw89_
                index85_ = default__.safeIndex(116, (d_469_v9_).length(0))
                (d_469_v9_)[index85_] = self.f4
                index86_ = default__.safeIndex(116, (d_469_v9_).length(0))
                (d_469_v9_)[index86_] = self.f4
                d_470_v10_: _dafny.Map
                d_470_v10_ = _dafny.Map({(d_469_v9_)[default__.safeIndex(116, (d_469_v9_).length(0))]: (self).f3})
                (globalState).f1 = len((d_470_v10_ if (d_469_v9_)[default__.safeIndex(116, (d_469_v9_).length(0))] else d_470_v10_))
                d_471_v11_: C0
                nw90_ = C0()
                nw90_.ctor__(((self).f3) > (p1), d_469_v9_)
                d_471_v11_ = nw90_
            d_472_v12_: _dafny.Array
            nw91_ = _dafny.Array(_dafny.MultiSet({}), 11)
            d_472_v12_ = nw91_
            d_473_v13_: _dafny.Array
            def lambda23_(d_474_i1_):
                return self.f4

            init15_ = lambda23_
            nw92_ = _dafny.Array(None, 15)
            for i0_15_ in range(nw92_.length(0)):
                nw92_[i0_15_] = init15_(i0_15_)
            d_473_v13_ = nw92_
            d_475_v14_: _dafny.MultiSet
            d_475_v14_ = _dafny.MultiSet([d_473_v13_, d_473_v13_])
            index87_ = default__.safeIndex(404, (d_472_v12_).length(0))
            (d_472_v12_)[index87_] = d_475_v14_
            index88_ = default__.safeIndex(404, (d_472_v12_).length(0))
            (d_472_v12_)[index88_] = _dafny.MultiSet([d_473_v13_])
            (globalState).f2 = False
        elif True:
            d_476_v15_: _dafny.Array
            nw93_ = _dafny.Array(_dafny.MultiSet({}), 11)
            d_476_v15_ = nw93_
            d_477_v16_: _dafny.MultiSet
            d_477_v16_ = _dafny.MultiSet([(0) - ((self).f14)])
            index89_ = default__.safeIndex(319, (d_476_v15_).length(0))
            (d_476_v15_)[index89_] = d_477_v16_
            d_478_v17_: str
            d_478_v17_ = _dafny.CodePoint('f')
            d_479_v18_: _dafny.Map
            d_479_v18_ = _dafny.Map({self.f4: d_478_v17_})
            index90_ = default__.safeIndex(319, (d_476_v15_).length(0))
            (d_476_v15_)[index90_] = (d_477_v16_) | (_dafny.MultiSet([len(d_479_v18_)]))
            rhs85_ = (self).f14
            rhs86_ = self.f4
            lhs76_ = globalState
            lhs77_ = globalState
            lhs76_.f1 = rhs85_
            lhs77_.f2 = rhs86_
            d_480_v19_: _dafny.Array
            def lambda24_(d_481_i2_):
                return _dafny.Map({self.f4: _dafny.MultiSet([self.f4, self.f4, self.f4])})

            init16_ = lambda24_
            nw94_ = _dafny.Array(None, 4)
            for i0_16_ in range(nw94_.length(0)):
                nw94_[i0_16_] = init16_(i0_16_)
            d_480_v19_ = nw94_
            d_482_v20_: _dafny.MultiSet
            d_482_v20_ = _dafny.MultiSet([self.f4, False])
            d_483_v21_: _dafny.Map
            d_483_v21_ = _dafny.Map({default__.fm2((self).f14, globalState): d_482_v20_})
            index91_ = default__.safeIndex(931, (d_480_v19_).length(0))
            (d_480_v19_)[index91_] = d_483_v21_
            d_484_v22_: D5
            d_484_v22_ = D5_DC14(d_483_v21_)
            index92_ = default__.safeIndex(931, (d_480_v19_).length(0))
            (d_480_v19_)[index92_] = ((d_484_v22_ if self.f4 else d_484_v22_)).cf20
            (globalState).f2 = not(self.f4)
            d_485_v23_: _dafny.Seq
            d_485_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "arajbrv"))
            if (d_485_v23_) < (d_485_v23_):
                d_486_v24_: _dafny.Array
                def lambda25_(d_487_i3_):
                    return not(self.f4)

                init17_ = lambda25_
                nw95_ = _dafny.Array(None, 5)
                for i0_17_ in range(nw95_.length(0)):
                    nw95_[i0_17_] = init17_(i0_17_)
                d_486_v24_ = nw95_
                d_488_v25_: _dafny.Array
                nw96_ = _dafny.Array(None, 25)
                nw96_[int(0)] = (d_486_v24_ if self.f4 else d_486_v24_)
                nw96_[int(1)] = d_486_v24_
                nw96_[int(2)] = d_486_v24_
                nw96_[int(3)] = d_486_v24_
                nw96_[int(4)] = d_486_v24_
                nw96_[int(5)] = d_486_v24_
                nw96_[int(6)] = d_486_v24_
                nw96_[int(7)] = d_486_v24_
                nw96_[int(8)] = d_486_v24_
                nw96_[int(9)] = d_486_v24_
                nw96_[int(10)] = d_486_v24_
                nw96_[int(11)] = d_486_v24_
                nw96_[int(12)] = d_486_v24_
                nw96_[int(13)] = d_486_v24_
                nw96_[int(14)] = (d_486_v24_ if not(self.f4) else d_486_v24_)
                nw96_[int(15)] = d_486_v24_
                nw96_[int(16)] = d_486_v24_
                nw96_[int(17)] = d_486_v24_
                nw96_[int(18)] = d_486_v24_
                nw96_[int(19)] = d_486_v24_
                nw96_[int(20)] = d_486_v24_
                nw96_[int(21)] = d_486_v24_
                nw96_[int(22)] = d_486_v24_
                nw96_[int(23)] = d_486_v24_
                nw96_[int(24)] = d_486_v24_
                d_488_v25_ = nw96_
                index93_ = default__.safeIndex(325, (d_488_v25_).length(0))
                (d_488_v25_)[index93_] = d_486_v24_
                index94_ = default__.safeIndex(325, (d_488_v25_).length(0))
                rhs87_ = d_486_v24_
                rhs88_ = default__.fm2(((self).f14) + (p1), globalState)
                lhs78_ = d_488_v25_
                lhs79_ = default__.safeIndex(325, (d_488_v25_).length(0))
                lhs80_ = globalState
                lhs78_[lhs79_] = rhs87_
                lhs80_.f2 = rhs88_
                (globalState).f1 = (p2) * (p2)
                index95_ = default__.safeIndex(319, (d_476_v15_).length(0))
                rhs89_ = (d_476_v15_)[default__.safeIndex(319, (d_476_v15_).length(0))]
                rhs90_ = self.f4
                lhs81_ = d_476_v15_
                lhs82_ = default__.safeIndex(319, (d_476_v15_).length(0))
                lhs83_ = globalState
                lhs81_[lhs82_] = rhs89_
                lhs83_.f2 = rhs90_
                d_485_v23_ = d_485_v23_
                d_489_v26_: _dafny.Set
                d_489_v26_ = _dafny.Set({((d_476_v15_)[default__.safeIndex(319, (d_476_v15_).length(0))]).ispropersubset(d_477_v16_), self.f4, self.f4, self.f4})
                d_490_v27_: _dafny.Seq
                d_490_v27_ = _dafny.SeqWithoutIsStrInference([d_489_v26_])
                d_491_v28_: _dafny.Seq
                d_491_v28_ = _dafny.SeqWithoutIsStrInference([p2, p2])
                d_489_v26_ = (d_490_v27_)[default__.safeIndex(len(d_491_v28_), len(d_490_v27_))]
            elif True:
                d_492_v29_: _dafny.Set
                d_492_v29_ = _dafny.Set({(self).f15, (self).f15, (self).f15, (self).f15, (self).f15})
                d_493_v30_: _dafny.Map
                d_493_v30_ = _dafny.Map({self.f4: d_492_v29_})
                d_493_v30_ = (d_493_v30_).set(False, d_492_v29_)
                d_494_v31_: _dafny.Map
                d_494_v31_ = _dafny.Map({214: (p1) == (p2)})
                d_494_v31_ = (d_494_v31_).set(p1, self.f4)
                (globalState).f1 = default__.safeDivisionInt(default__.safeDivisionInt((self).f14, (self).f14), (self).f14)
                d_495_v32_: _dafny.Array
                def lambda26_(d_496_v21_):
                    def lambda27_(d_497_i4_):
                        return D5_DC16(D5_DC16(D5_DC16(D5_DC16(D5_DC14((d_496_v21_).set(self.f4, _dafny.MultiSet([self.f4, self.f4])))))))

                    return lambda27_

                init18_ = lambda26_(d_483_v21_)
                nw97_ = _dafny.Array(None, 4)
                for i0_18_ in range(nw97_.length(0)):
                    nw97_[i0_18_] = init18_(i0_18_)
                d_495_v32_ = nw97_
                index96_ = default__.safeIndex(114, ((self).f15).length(0))
                ((self).f15)[index96_] = (d_482_v20_).cardinality
                index97_ = default__.safeIndex(114, ((self).f15).length(0))
                rhs91_ = d_495_v32_
                rhs92_ = not (self.f4) or (self.f4)
                rhs93_ = self.f4
                rhs94_ = default__.safeModuloInt((self).f14, p1)
                rhs95_ = ((0) - ((self).f14)) + (((self).f14) - (p1))
                lhs84_ = globalState
                lhs85_ = globalState
                lhs86_ = globalState
                lhs87_ = (self).f15
                lhs88_ = default__.safeIndex(114, ((self).f15).length(0))
                d_495_v32_ = rhs91_
                lhs84_.f2 = rhs92_
                lhs85_.f2 = rhs93_
                lhs86_.f1 = rhs94_
                lhs87_[lhs88_] = rhs95_
                d_498_v33_: _dafny.Array
                def lambda28_(d_499_i5_):
                    return (self.f4) == (self.f4)

                init19_ = lambda28_
                nw98_ = _dafny.Array(None, 15)
                for i0_19_ in range(nw98_.length(0)):
                    nw98_[i0_19_] = init19_(i0_19_)
                d_498_v33_ = nw98_
                nw99_ = _dafny.Array(False, 19)
                d_498_v33_ = nw99_
        d_500_i6_: int
        d_500_i6_ = 0
        with _dafny.label("5"):
            while default__.fm2((0) - ((self).f14), globalState):
                with _dafny.c_label("5"):
                    if (d_500_i6_) >= (100):
                        raise _dafny.Break("5")
                    d_500_i6_ = (d_500_i6_) + (1)
                    d_501_v34_: str
                    d_501_v34_ = _dafny.CodePoint('x')
                    d_502_v35_: _dafny.MultiSet
                    d_502_v35_ = _dafny.MultiSet([d_501_v34_, d_501_v34_])
                    d_503_v36_: D0
                    d_503_v36_ = D0_DC2(d_502_v35_, self.f4)
                    d_504_v37_: D0
                    d_504_v37_ = D0_DC4(d_503_v36_)
                    rhs96_ = p2
                    rhs97_ = (default__.fm0(globalState)) == (776)
                    rhs98_ = self.f4
                    rhs99_ = (d_504_v37_ if self.f4 else d_504_v37_)
                    lhs89_ = globalState
                    lhs90_ = globalState
                    lhs91_ = globalState
                    lhs89_.f1 = rhs96_
                    lhs90_.f2 = rhs97_
                    lhs91_.f2 = rhs98_
                    d_504_v37_ = rhs99_
                    d_505_v38_: _dafny.Array
                    nw100_ = _dafny.Array(_dafny.Array(None, 0), 8)
                    d_505_v38_ = nw100_
                    d_506_v39_: _dafny.Array
                    nw101_ = _dafny.Array(None, 18)
                    nw101_[int(0)] = self.f4
                    nw101_[int(1)] = self.f4
                    nw101_[int(2)] = self.f4
                    nw101_[int(3)] = self.f4
                    nw101_[int(4)] = self.f4
                    nw101_[int(5)] = self.f4
                    nw101_[int(6)] = self.f4
                    nw101_[int(7)] = self.f4
                    nw101_[int(8)] = self.f4
                    nw101_[int(9)] = self.f4
                    nw101_[int(10)] = self.f4
                    nw101_[int(11)] = self.f4
                    nw101_[int(12)] = not(self.f4)
                    nw101_[int(13)] = self.f4
                    nw101_[int(14)] = self.f4
                    nw101_[int(15)] = self.f4
                    nw101_[int(16)] = self.f4
                    nw101_[int(17)] = default__.fm2(p1, globalState)
                    d_506_v39_ = nw101_
                    index98_ = default__.safeIndex(909, (d_505_v38_).length(0))
                    (d_505_v38_)[index98_] = d_506_v39_
                    d_507_v40_: _dafny.Seq
                    d_507_v40_ = _dafny.SeqWithoutIsStrInference([(self.f4 if self.f4 else self.f4), False, self.f4, self.f4])
                    index99_ = default__.safeIndex(221, (d_506_v39_).length(0))
                    (d_506_v39_)[index99_] = self.f4
                    index100_ = default__.safeIndex(909, (d_505_v38_).length(0))
                    index101_ = default__.safeIndex(221, (d_506_v39_).length(0))
                    rhs100_ = d_506_v39_
                    rhs101_ = d_507_v40_
                    rhs102_ = (p1) == ((self).f3)
                    lhs92_ = d_505_v38_
                    lhs93_ = default__.safeIndex(909, (d_505_v38_).length(0))
                    lhs94_ = d_506_v39_
                    lhs95_ = default__.safeIndex(221, (d_506_v39_).length(0))
                    lhs92_[lhs93_] = rhs100_
                    d_507_v40_ = rhs101_
                    lhs94_[lhs95_] = rhs102_
                    d_508_v41_: _dafny.Map
                    d_508_v41_ = _dafny.Map({(d_506_v39_)[default__.safeIndex(221, (d_506_v39_).length(0))]: (self).f3})
                    d_509_v42_: _dafny.MultiSet
                    d_509_v42_ = _dafny.MultiSet([850])
                    d_510_v43_: _dafny.MultiSet
                    d_510_v43_ = _dafny.MultiSet([False, (d_506_v39_)[default__.safeIndex(221, (d_506_v39_).length(0))]])
                    d_511_v44_: _dafny.Map
                    d_511_v44_ = _dafny.Map({(d_502_v35_).cardinality: self.f4})
                    d_512_v45_: _dafny.Map
                    d_512_v45_ = _dafny.Map({(self).f14: len(d_511_v44_)})
                    d_513_v46_: _dafny.Array
                    nw102_ = _dafny.Array(None, 25)
                    nw102_[int(0)] = default__.safeModuloInt(p1, (self).f3)
                    nw102_[int(1)] = (len(d_508_v41_)) + (688)
                    nw102_[int(2)] = p1
                    nw102_[int(3)] = (p1) * ((self).f3)
                    nw102_[int(4)] = ((self).f14) * ((self).f14)
                    nw102_[int(5)] = default__.fm0(globalState)
                    nw102_[int(6)] = default__.safeModuloInt((self).f3, (self).f3)
                    nw102_[int(7)] = (p2) - (-388)
                    nw102_[int(8)] = (0) - ((self).f14)
                    nw102_[int(9)] = ((D1_DC5(default__.fm0(globalState))).cf8) - (p1)
                    nw102_[int(10)] = ((d_509_v42_)[p2] if (p2) in (d_509_v42_) else (self).f14)
                    nw102_[int(11)] = (self).f14
                    nw102_[int(12)] = (self).f14
                    nw102_[int(13)] = p2
                    nw102_[int(14)] = p2
                    nw102_[int(15)] = p2
                    nw102_[int(16)] = (self).f3
                    nw102_[int(17)] = p2
                    nw102_[int(18)] = p1
                    nw102_[int(19)] = p1
                    nw102_[int(20)] = (self).f3
                    nw102_[int(21)] = p1
                    nw102_[int(22)] = ((d_510_v43_)[(d_506_v39_)[default__.safeIndex(221, (d_506_v39_).length(0))]] if ((d_506_v39_)[default__.safeIndex(221, (d_506_v39_).length(0))]) in (d_510_v43_) else len(d_512_v45_))
                    nw102_[int(23)] = (self).f3
                    nw102_[int(24)] = -440
                    d_513_v46_ = nw102_
                    d_514_v47_: _dafny.Array
                    nw103_ = _dafny.Array(None, 19)
                    d_514_v47_ = nw103_
                    d_515_v48_: D8
                    d_515_v48_ = D8_DC22(d_511_v44_)
                    d_516_v49_: T1
                    nw104_ = C1()
                    nw104_.ctor__((self).f3, len((d_515_v48_).cf31), False)
                    d_516_v49_ = nw104_
                    index102_ = default__.safeIndex(327, (d_514_v47_).length(0))
                    (d_514_v47_)[index102_] = d_516_v49_
                    d_517_v50_: _dafny.Map
                    d_517_v50_ = _dafny.Map({(d_512_v45_).set(p2, (self).f3): d_513_v46_})
                    index103_ = default__.safeIndex(327, (d_514_v47_).length(0))
                    rhs103_ = (d_506_v39_)[default__.safeIndex(221, (d_506_v39_).length(0))]
                    rhs104_ = p1
                    rhs105_ = ((d_517_v50_)[(default__.fm23(globalState)).set((d_516_v49_).f7, (0) - ((d_516_v49_).f7))] if ((default__.fm23(globalState)).set((d_516_v49_).f7, (0) - ((d_516_v49_).f7))) in (d_517_v50_) else d_513_v46_)
                    rhs106_ = d_516_v49_
                    rhs107_ = (d_516_v49_).f3
                    lhs96_ = self
                    lhs97_ = globalState
                    lhs98_ = d_514_v47_
                    lhs99_ = default__.safeIndex(327, (d_514_v47_).length(0))
                    lhs100_ = globalState
                    lhs96_.f4 = rhs103_
                    lhs97_.f1 = rhs104_
                    d_513_v46_ = rhs105_
                    lhs98_[lhs99_] = rhs106_
                    lhs100_.f1 = rhs107_
                    d_518_v51_: _dafny.Array
                    nw105_ = _dafny.Array(None, 17)
                    nw105_[int(0)] = d_501_v34_
                    nw105_[int(1)] = d_501_v34_
                    nw105_[int(2)] = d_501_v34_
                    nw105_[int(3)] = d_501_v34_
                    nw105_[int(4)] = d_501_v34_
                    nw105_[int(5)] = d_501_v34_
                    nw105_[int(6)] = d_501_v34_
                    nw105_[int(7)] = d_501_v34_
                    nw105_[int(8)] = (_dafny.CodePoint('k') if self.f4 else d_501_v34_)
                    nw105_[int(9)] = d_501_v34_
                    nw105_[int(10)] = d_501_v34_
                    nw105_[int(11)] = d_501_v34_
                    nw105_[int(12)] = d_501_v34_
                    nw105_[int(13)] = d_501_v34_
                    nw105_[int(14)] = d_501_v34_
                    nw105_[int(15)] = d_501_v34_
                    nw105_[int(16)] = d_501_v34_
                    d_518_v51_ = nw105_
                    index104_ = default__.safeIndex(317, (d_518_v51_).length(0))
                    (d_518_v51_)[index104_] = d_501_v34_
                    d_519_v52_: _dafny.Map
                    d_519_v52_ = _dafny.Map({default__.safeModuloInt(p1, p1): _dafny.CodePoint('n')})
                    index105_ = default__.safeIndex(317, (d_518_v51_).length(0))
                    (d_518_v51_)[index105_] = ((d_519_v52_)[693] if (693) in (d_519_v52_) else _dafny.CodePoint('g'))
                    pass
            pass
        d_520_v53_: _dafny.Array
        def lambda29_(d_521_i8_):
            return False

        init20_ = lambda29_
        nw106_ = _dafny.Array(None, 3)
        for i0_20_ in range(nw106_.length(0)):
            nw106_[i0_20_] = init20_(i0_20_)
        d_520_v53_ = nw106_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_520_v53_).length(0)):
            d_522_i7_: int = guard_loop_1_
            if (True) and (((0) <= (d_522_i7_)) and ((d_522_i7_) < ((d_520_v53_).length(0)))):
                (d_520_v53_)[(d_522_i7_)] = self.f4
        d_523_v54_: _dafny.Seq
        d_523_v54_ = _dafny.SeqWithoutIsStrInference([self.f4, self.f4])
        d_524_v55_: D7
        d_524_v55_ = D7_DC21(d_520_v53_, (d_523_v54_) < (_dafny.SeqWithoutIsStrInference([self.f4, self.f4, self.f4, self.f4])), not(self.f4))
        d_524_v55_ = d_524_v55_
        d_525_v56_: _dafny.Map
        d_525_v56_ = _dafny.Map({self.f4: (p1) + (p2)})
        d_526_v57_: str
        d_526_v57_ = _dafny.CodePoint('v')
        d_527_v58_: _dafny.MultiSet
        d_527_v58_ = _dafny.MultiSet([_dafny.CodePoint('w'), d_526_v57_, _dafny.CodePoint('r')])
        d_525_v56_ = (d_525_v56_).set((D0_DC2(d_527_v58_, self.f4)).cf6, default__.safeDivisionInt((self).f3, (self).f3))
        if not ((self.f4) or (True)) or (self.f4):
            d_528_v59_: D8
            d_528_v59_ = D8_DC23(self.f4, self.f4, self.f4, d_526_v57_)
            d_529_v60_: _dafny.Seq
            d_529_v60_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dakcdomv"))
            pat_let_tv4_ = globalState
            def iife18_(_pat_let5_0):
                def iife19_(d_530_dt__update__tmp_h0_):
                    def iife20_(_pat_let6_0):
                        def iife21_(d_531_dt__update_hcf35_h0_):
                            def iife22_(_pat_let7_0):
                                def iife23_(d_532_dt__update_hcf32_h0_):
                                    return D8_DC23(d_532_dt__update_hcf32_h0_, (d_530_dt__update__tmp_h0_).cf33, (d_530_dt__update__tmp_h0_).cf34, d_531_dt__update_hcf35_h0_)
                                return iife23_(_pat_let7_0)
                            return iife22_(self.f4)
                        return iife21_(_pat_let6_0)
                    return iife20_(default__.fm24(pat_let_tv4_))
                return iife19_(_pat_let5_0)
            rhs108_ = iife18_((d_528_v59_ if True else d_528_v59_))
            rhs109_ = self.f4
            rhs110_ = (d_526_v57_) in (d_529_v60_)
            lhs101_ = globalState
            lhs102_ = globalState
            d_528_v59_ = rhs108_
            lhs101_.f2 = rhs109_
            lhs102_.f2 = rhs110_
            d_526_v57_ = d_526_v57_
            d_533_v61_: _dafny.Map
            d_533_v61_ = _dafny.Map({(self).f3: (self).f14})
            index106_ = default__.safeIndex(271, ((self).f15).length(0))
            ((self).f15)[index106_] = ((d_533_v61_)[(self).f3] if ((self).f3) in (d_533_v61_) else (0) - ((self).f3))
            d_534_v62_: _dafny.Set
            d_534_v62_ = _dafny.Set({self.f4})
            index107_ = default__.safeIndex(271, ((self).f15).length(0))
            ((self).f15)[index107_] = (p1) + (len((d_534_v62_) | (d_534_v62_)))
            d_523_v54_ = d_523_v54_
            d_535_v63_: _dafny.Seq
            d_535_v63_ = _dafny.SeqWithoutIsStrInference([p1, p2])
            (globalState).f1 = len((d_535_v63_) + (((_dafny.SeqWithoutIsStrInference([p2 for d_536_i9_ in range(default__.abs(222))])) + (d_535_v63_)).set(default__.safeIndex(((self).f15)[default__.safeIndex(271, ((self).f15).length(0))], len((_dafny.SeqWithoutIsStrInference([p2 for d_537_i9_ in range(default__.abs(222))])) + (d_535_v63_))), p1)))
        elif True:
            d_538_v64_: _dafny.Map
            d_538_v64_ = _dafny.Map({d_523_v54_: (self).f3})
            (globalState).f1 = ((d_538_v64_)[((d_523_v54_) + (d_523_v54_)).set(default__.safeIndex(default__.fm0(globalState), len((d_523_v54_) + (d_523_v54_))), True)] if (((d_523_v54_) + (d_523_v54_)).set(default__.safeIndex(default__.fm0(globalState), len((d_523_v54_) + (d_523_v54_))), True)) in (d_538_v64_) else p1)
            d_539_v65_: D8
            d_539_v65_ = D8_DC23(self.f4, self.f4, self.f4, d_526_v57_)
            d_540_v66_: D8
            d_540_v66_ = D8_DC24(d_539_v65_)
            d_541_v67_: D8
            d_541_v67_ = D8_DC24(d_540_v66_)
            d_542_v68_: D8
            d_542_v68_ = D8_DC24(d_541_v67_)
            d_543_v69_: _dafny.Map
            d_543_v69_ = _dafny.Map({True: d_542_v68_})
            d_543_v69_ = ((d_543_v69_) | (d_543_v69_)) | (_dafny.Map({self.f4: d_542_v68_}))
            d_544_v70_: C1
            nw107_ = C1()
            nw107_.ctor__((self).f14, p2, self.f4)
            d_544_v70_ = nw107_
            d_545_v71_: _dafny.Seq
            d_545_v71_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "axxyv"))
            d_545_v71_ = d_545_v71_
            d_546_v72_: _dafny.Array
            def lambda30_(d_547_v58_):
                def lambda31_(d_548_i10_):
                    return _dafny.SeqWithoutIsStrInference([(d_547_v58_).cardinality])

                return lambda31_

            init21_ = lambda30_(d_527_v58_)
            nw108_ = _dafny.Array(None, 2)
            for i0_21_ in range(nw108_.length(0)):
                nw108_[i0_21_] = init21_(i0_21_)
            d_546_v72_ = nw108_
            d_549_v73_: _dafny.Set
            d_549_v73_ = _dafny.Set({(self).f3})
            d_550_v74_: _dafny.Map
            d_550_v74_ = _dafny.Map({True: True})
            rhs111_ = -860
            rhs112_ = not((self.f4) and ((d_549_v73_) == (d_549_v73_)))
            rhs113_ = default__.fm2(len((d_550_v74_) | ((d_550_v74_).set(self.f4, self.f4))), globalState)
            rhs114_ = default__.fm2(p2, globalState)
            rhs115_ = d_546_v72_
            lhs103_ = globalState
            lhs104_ = globalState
            lhs105_ = globalState
            lhs106_ = globalState
            lhs103_.f1 = rhs111_
            lhs104_.f2 = rhs112_
            lhs105_.f2 = rhs113_
            lhs106_.f2 = rhs114_
            d_546_v72_ = rhs115_

    def m7(self, p0, p1, globalState):
        r0: int = int(0)
        r1: bool = False
        r0 = (self).f14
        d_551_v0_: _dafny.Seq
        d_551_v0_ = _dafny.SeqWithoutIsStrInference([-638])
        d_552_v1_: _dafny.Seq
        d_552_v1_ = _dafny.SeqWithoutIsStrInference([len(d_551_v0_), (p0) + (len(_dafny.Map({True: (self).f3}))), (-673) + ((self).f14)])
        d_552_v1_ = d_551_v0_
        hi5_ = default__.fm0(globalState)
        for d_553_i0_ in range((self).f14, hi5_):
            d_554_v2_: str
            d_554_v2_ = _dafny.CodePoint('i')
            d_555_v3_: _dafny.Set
            d_555_v3_ = _dafny.Set({d_554_v2_})
            d_556_v4_: _dafny.Array
            nw109_ = _dafny.Array(None, 6)
            nw109_[int(0)] = d_555_v3_
            nw109_[int(1)] = d_555_v3_
            nw109_[int(2)] = (d_555_v3_).intersection(d_555_v3_)
            nw109_[int(3)] = d_555_v3_
            nw109_[int(4)] = _dafny.Set({default__.fm24(globalState), d_554_v2_, d_554_v2_})
            nw109_[int(5)] = (d_555_v3_ if self.f4 else d_555_v3_)
            d_556_v4_ = nw109_
            d_556_v4_ = d_556_v4_
            source6_ = default__.fm25((default__.fm2(-606, globalState) if self.f4 else self.f4), globalState)
            if source6_.is_DC1:
                d_557___mcc_h0_ = source6_.cf1
                d_558___mcc_h1_ = source6_.cf2
                d_559___mcc_h2_ = source6_.cf3
                d_560___mcc_h3_ = source6_.cf4
                d_561_cf4_ = d_560___mcc_h3_
                d_562_cf3_ = d_559___mcc_h2_
                d_563_cf2_ = d_558___mcc_h1_
                d_564_cf1_ = d_557___mcc_h0_
                d_565_v5_: D0
                d_565_v5_ = D0_DC1(False, d_564_cf1_, d_562_cf3_, d_561_cf4_)
                (globalState).f1 = len(_dafny.SeqWithoutIsStrInference([(d_565_v5_).cf3]))
                d_566_v6_: _dafny.Map
                d_566_v6_ = _dafny.Map({False: self.f4})
                d_567_v7_: _dafny.Set
                d_567_v7_ = _dafny.Set({d_564_cf1_})
                d_568_v8_: _dafny.Map
                d_568_v8_ = _dafny.Map({d_564_cf1_: d_567_v7_})
                d_569_v9_: _dafny.Seq
                d_569_v9_ = _dafny.SeqWithoutIsStrInference([d_567_v7_, _dafny.Set({False})])
                d_570_v10_: _dafny.Seq
                d_570_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ltfp"))
                d_571_v11_: _dafny.MultiSet
                d_571_v11_ = _dafny.MultiSet([self.f4, True])
                d_572_v12_: _dafny.Map
                d_572_v12_ = _dafny.Map({d_563_cf2_: d_571_v11_})
                d_573_v13_: _dafny.Array
                nw110_ = _dafny.Array(None, 14)
                nw110_[int(0)] = self.f4
                nw110_[int(1)] = self.f4
                nw110_[int(2)] = ((d_566_v6_)[default__.fm2((self).f3, globalState)] if (default__.fm2((self).f3, globalState)) in (d_566_v6_) else self.f4)
                nw110_[int(3)] = (default__.fm20(len(((d_568_v8_)[default__.fm2(p0, globalState)] if (default__.fm2(p0, globalState)) in (d_568_v8_) else (d_569_v9_)[default__.safeIndex(d_553_i0_, len(d_569_v9_))])), globalState)) < (d_570_v10_)
                nw110_[int(4)] = (self.f4) or (default__.fm2(len(d_572_v12_), globalState))
                nw110_[int(5)] = False
                nw110_[int(6)] = d_563_cf2_
                nw110_[int(7)] = not(self.f4)
                nw110_[int(8)] = ((self).f14) < (d_553_i0_)
                nw110_[int(9)] = (d_564_cf1_ if d_563_cf2_ else d_563_cf2_)
                nw110_[int(10)] = ((d_566_v6_)[d_563_cf2_] if (d_563_cf2_) in (d_566_v6_) else d_563_cf2_)
                nw110_[int(11)] = not(self.f4)
                nw110_[int(12)] = self.f4
                nw110_[int(13)] = (default__.fm0(globalState)) != (p0)
                d_573_v13_ = nw110_
                d_574_v14_: _dafny.MultiSet
                d_574_v14_ = _dafny.MultiSet([d_571_v11_, d_571_v11_])
                index108_ = default__.safeIndex(435, (d_573_v13_).length(0))
                (d_573_v13_)[index108_] = (d_574_v14_) == (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_571_v11_])))
                index109_ = default__.safeIndex(435, (d_573_v13_).length(0))
                (d_573_v13_)[index109_] = (d_553_i0_) < (d_553_i0_)
                d_575_v15_: C0
                nw111_ = C0()
                nw111_.ctor__((d_573_v13_)[default__.safeIndex(435, (d_573_v13_).length(0))], d_573_v13_)
                d_575_v15_ = nw111_
                index110_ = default__.safeIndex(232, ((self).f15).length(0))
                ((self).f15)[index110_] = (self).f3
                index111_ = default__.safeIndex(232, ((self).f15).length(0))
                rhs116_ = (self).f15
                rhs117_ = ((0) - ((self).f14)) - ((0) - ((self).f14))
                lhs107_ = (self).f15
                lhs108_ = default__.safeIndex(232, ((self).f15).length(0))
                d_561_cf4_ = rhs116_
                lhs107_[lhs108_] = rhs117_
            elif source6_.is_DC2:
                d_576___mcc_h4_ = source6_.cf5
                d_577___mcc_h5_ = source6_.cf6
                d_578_cf6_ = d_577___mcc_h5_
                d_579_cf5_ = d_576___mcc_h4_
                d_580_v16_: _dafny.Array
                nw112_ = _dafny.Array(_dafny.Array(None, 0), 18)
                d_580_v16_ = nw112_
                index112_ = default__.safeIndex(570, (d_580_v16_).length(0))
                (d_580_v16_)[index112_] = ((self).f15 if d_578_cf6_ else (self).f15)
                d_581_v17_: _dafny.Array
                nw113_ = _dafny.Array(D8.default()(), 26)
                d_581_v17_ = nw113_
                d_582_v18_: D8
                d_582_v18_ = D8_DC23(False, d_578_cf6_, self.f4, d_554_v2_)
                d_583_v19_: D8
                d_583_v19_ = D8_DC24(d_582_v18_)
                index113_ = default__.safeIndex(626, (d_581_v17_).length(0))
                (d_581_v17_)[index113_] = d_583_v19_
                d_584_v20_: _dafny.MultiSet
                d_584_v20_ = _dafny.MultiSet([d_553_i0_, (self).f14])
                d_585_v21_: _dafny.Seq
                d_585_v21_ = _dafny.SeqWithoutIsStrInference([d_578_cf6_, d_578_cf6_])
                index114_ = default__.safeIndex(570, (d_580_v16_).length(0))
                index115_ = default__.safeIndex(626, (d_581_v17_).length(0))
                rhs118_ = (self).f15
                rhs119_ = default__.fm26(globalState)
                rhs120_ = (d_584_v20_).set(len(d_585_v21_), default__.abs((self).f3))
                lhs109_ = d_580_v16_
                lhs110_ = default__.safeIndex(570, (d_580_v16_).length(0))
                lhs111_ = d_581_v17_
                lhs112_ = default__.safeIndex(626, (d_581_v17_).length(0))
                lhs109_[lhs110_] = rhs118_
                lhs111_[lhs112_] = rhs119_
                d_584_v20_ = rhs120_
                d_586_v22_: _dafny.Map
                d_586_v22_ = _dafny.Map({d_553_i0_: self.f4})
                d_586_v22_ = (d_586_v22_).set((0) - ((self).f3), self.f4)
                r0 = d_553_i0_
                r0 = (self).f3
            elif source6_.is_DC3:
                d_587_v23_: _dafny.Seq
                d_587_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "af"))
                d_588_v24_: _dafny.Map
                d_588_v24_ = _dafny.Map({(d_553_i0_) >= (-271): _dafny.SeqWithoutIsStrInference([d_587_v23_])})
                d_589_v25_: _dafny.Seq
                d_589_v25_ = _dafny.SeqWithoutIsStrInference([self.f4])
                d_590_v26_: _dafny.Seq
                d_590_v26_ = _dafny.SeqWithoutIsStrInference([d_587_v23_])
                d_588_v24_ = (d_588_v24_).set(not((self.f4) == ((d_589_v25_)[default__.safeIndex(d_553_i0_, len(d_589_v25_))])), d_590_v26_)
                d_591_v27_: _dafny.Array
                nw114_ = _dafny.Array(None, 17)
                nw114_[int(0)] = self.f4
                nw114_[int(1)] = True
                nw114_[int(2)] = self.f4
                nw114_[int(3)] = self.f4
                nw114_[int(4)] = not(self.f4)
                nw114_[int(5)] = self.f4
                nw114_[int(6)] = default__.fm2(d_553_i0_, globalState)
                nw114_[int(7)] = self.f4
                nw114_[int(8)] = True
                nw114_[int(9)] = self.f4
                nw114_[int(10)] = self.f4
                nw114_[int(11)] = False
                nw114_[int(12)] = True
                nw114_[int(13)] = False
                nw114_[int(14)] = self.f4
                nw114_[int(15)] = self.f4
                nw114_[int(16)] = self.f4
                d_591_v27_ = nw114_
                d_592_v28_: _dafny.Map
                d_592_v28_ = _dafny.Map({self.f4: d_591_v27_})
                d_592_v28_ = (d_592_v28_).set(True, d_591_v27_)
                d_593_v29_: C0
                nw115_ = C0()
                nw115_.ctor__(self.f4, d_591_v27_)
                d_593_v29_ = nw115_
                (self).f4 = d_593_v29_.f10
            elif source6_.is_DC0:
                d_594___mcc_h6_ = source6_.cf0
                d_595_cf0_ = d_594___mcc_h6_
                d_596_v30_: _dafny.Seq
                d_596_v30_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gjdgtgnld"))
                d_597_v31_: _dafny.Map
                d_597_v31_ = _dafny.Map({d_596_v30_: d_553_i0_})
                d_598_v32_: _dafny.Map
                d_598_v32_ = _dafny.Map({d_595_cf0_: False})
                d_597_v31_ = (d_597_v31_).set(d_596_v30_, (_dafny.MultiSet([d_595_cf0_, ((d_598_v32_)[self.f4] if (self.f4) in (d_598_v32_) else True)])).cardinality)
                index116_ = default__.safeIndex(586, ((self).f15).length(0))
                ((self).f15)[index116_] = (self).f14
                index117_ = default__.safeIndex(586, ((self).f15).length(0))
                ((self).f15)[index117_] = (self).f14
                d_599_v33_: _dafny.Array
                nw116_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 14)
                d_599_v33_ = nw116_
                index118_ = default__.safeIndex(836, (d_599_v33_).length(0))
                (d_599_v33_)[index118_] = d_596_v30_
                d_600_v34_: _dafny.Map
                d_600_v34_ = _dafny.Map({(self).f14: ((self).f15)[default__.safeIndex(586, ((self).f15).length(0))]})
                d_601_v36_: _dafny.Map
                def iife24_():
                    coll8_ = _dafny.Set()
                    compr_8_: int
                    for compr_8_ in (d_600_v34_).keys.Elements:
                        d_602_v35_: int = compr_8_
                        if (d_602_v35_) in (d_600_v34_):
                            coll8_ = coll8_.union(_dafny.Set([(d_602_v35_) * (-756)]))
                    return _dafny.Set(coll8_)
                d_601_v36_ = _dafny.Map({self.f4: len(iife24_()
                )})
                d_603_v37_: _dafny.MultiSet
                d_603_v37_ = _dafny.MultiSet([d_601_v36_, d_601_v36_])
                d_604_v38_: _dafny.Array
                nw117_ = _dafny.Array(False, 20)
                d_604_v38_ = nw117_
                d_605_v39_: C0
                nw118_ = C0()
                nw118_.ctor__(d_595_cf0_, d_604_v38_)
                d_605_v39_ = nw118_
                d_606_v40_: _dafny.Map
                d_606_v40_ = _dafny.Map({d_605_v39_: d_596_v30_})
                d_607_v41_: _dafny.Map
                d_607_v41_ = _dafny.Map({(d_603_v37_).set(d_601_v36_, default__.abs((self).f3)): ((d_606_v40_)[d_605_v39_] if (d_605_v39_) in (d_606_v40_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tpoiv")))})
                index119_ = default__.safeIndex(836, (d_599_v33_).length(0))
                (d_599_v33_)[index119_] = ((d_607_v41_)[d_603_v37_] if (d_603_v37_) in (d_607_v41_) else d_596_v30_)
                d_551_v0_ = _dafny.SeqWithoutIsStrInference([(self).f3 for d_608_i1_ in range(default__.abs(4))])
            elif True:
                d_609___mcc_h7_ = source6_.cf7
                d_610_cf7_ = d_609___mcc_h7_
                r1 = default__.fm2((0) - ((self).f3), globalState)
                d_611_v42_: _dafny.Map
                d_611_v42_ = _dafny.Map({(self).f14: True})
                d_611_v42_ = (d_611_v42_).set((self).f3, not((p0) > ((self).f14)))
                d_612_v43_: _dafny.Seq
                d_612_v43_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "viwcggtyo"))
                d_613_v44_: _dafny.Set
                d_613_v44_ = _dafny.Set({p0, len(d_612_v43_)})
                d_614_v45_: _dafny.MultiSet
                d_614_v45_ = _dafny.MultiSet([(self).f14])
                r0 = (len(d_613_v44_)) + (((self).f3 if self.f4 else (d_614_v45_).cardinality))
                d_615_v46_: _dafny.Array
                def lambda32_(d_616_i2_):
                    return self.f4

                init22_ = lambda32_
                nw119_ = _dafny.Array(None, 29)
                for i0_22_ in range(nw119_.length(0)):
                    nw119_[i0_22_] = init22_(i0_22_)
                d_615_v46_ = nw119_
                d_617_v47_: D7
                d_617_v47_ = D7_DC21(d_615_v46_, self.f4, self.f4)
                pat_let_tv5_ = d_615_v46_
                d_618_v48_: _dafny.Array
                nw120_ = _dafny.Array(None, 15)
                nw120_[int(0)] = D7_DC21(d_615_v46_, self.f4, self.f4)
                nw120_[int(1)] = d_617_v47_
                nw120_[int(2)] = d_617_v47_
                nw120_[int(3)] = d_617_v47_
                nw120_[int(4)] = d_617_v47_
                nw120_[int(5)] = d_617_v47_
                nw120_[int(6)] = d_617_v47_
                nw120_[int(7)] = d_617_v47_
                nw120_[int(8)] = d_617_v47_
                nw120_[int(9)] = d_617_v47_
                def iife25_(_pat_let8_0):
                    def iife26_(d_619_dt__update__tmp_h0_):
                        def iife27_(_pat_let9_0):
                            def iife28_(d_620_dt__update_hcf29_h0_):
                                return D7_DC21((d_619_dt__update__tmp_h0_).cf28, d_620_dt__update_hcf29_h0_, (d_619_dt__update__tmp_h0_).cf30)
                            return iife28_(_pat_let9_0)
                        return iife27_(self.f4)
                    return iife26_(_pat_let8_0)
                nw120_[int(10)] = iife25_(d_617_v47_)
                nw120_[int(11)] = D7_DC21(d_615_v46_, self.f4, self.f4)
                def iife29_(_pat_let10_0):
                    def iife30_(d_621_dt__update__tmp_h1_):
                        def iife31_(_pat_let11_0):
                            def iife32_(d_622_dt__update_hcf28_h0_):
                                return D7_DC21(d_622_dt__update_hcf28_h0_, (d_621_dt__update__tmp_h1_).cf29, (d_621_dt__update__tmp_h1_).cf30)
                            return iife32_(_pat_let11_0)
                        return iife31_(pat_let_tv5_)
                    return iife30_(_pat_let10_0)
                nw120_[int(12)] = iife29_(D7_DC21(d_615_v46_, False, self.f4))
                nw120_[int(13)] = d_617_v47_
                nw120_[int(14)] = D7_DC21(d_615_v46_, self.f4, self.f4)
                d_618_v48_ = nw120_
                d_623_v49_: _dafny.Set
                d_623_v49_ = _dafny.Set({d_618_v48_, d_618_v48_})
                d_624_v50_: _dafny.Seq
                d_624_v50_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({d_618_v48_, d_618_v48_})])
                d_623_v49_ = ((d_624_v50_)[default__.safeIndex(p0, len(d_624_v50_))]) | (d_623_v49_)
            d_625_v52_: D7
            def iife33_():
                coll9_ = _dafny.Map()
                compr_9_: int
                for compr_9_ in _dafny.IntegerRange(353, 82):
                    d_626_v51_: int = compr_9_
                    if ((353) <= (d_626_v51_)) and ((d_626_v51_) < (82)):
                        coll9_[default__.safeModuloInt(d_626_v51_, p0)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xoulfc"))
                return _dafny.Map(coll9_)
            d_625_v52_ = D7_DC20(iife33_()
)
            source7_ = d_625_v52_
            if source7_.is_DC21:
                d_627___mcc_h8_ = source7_.cf28
                d_628___mcc_h9_ = source7_.cf29
                d_629___mcc_h10_ = source7_.cf30
                d_630_cf30_ = d_629___mcc_h10_
                d_631_cf29_ = d_628___mcc_h9_
                d_632_cf28_ = d_627___mcc_h8_
                (self).f4 = d_631_cf29_
                d_633_v53_: _dafny.Map
                d_633_v53_ = _dafny.Map({p0: True})
                index120_ = default__.safeIndex(354, ((self).f15).length(0))
                ((self).f15)[index120_] = (p0) + (len(d_633_v53_))
                d_634_v54_: D7
                d_634_v54_ = D7_DC21(d_632_cf28_, d_630_cf30_, not(d_631_cf29_))
                d_635_v55_: _dafny.Map
                d_635_v55_ = _dafny.Map({self.f4: d_553_i0_})
                d_636_v56_: D3
                d_636_v56_ = D3_DC10((d_553_i0_) + ((0) - (d_553_i0_)), (d_634_v54_).cf28, ((d_635_v55_)[d_630_cf30_] if (d_630_cf30_) in (d_635_v55_) else (self).f14))
                d_637_v57_: _dafny.Seq
                d_637_v57_ = _dafny.SeqWithoutIsStrInference([d_632_cf28_])
                index121_ = default__.safeIndex(686, (p1).length(0))
                (p1)[index121_] = (d_637_v57_)[default__.safeIndex((self).f3, len(d_637_v57_))]
                index122_ = default__.safeIndex(354, ((self).f15).length(0))
                index123_ = default__.safeIndex(686, (p1).length(0))
                rhs121_ = -163
                rhs122_ = (default__.fm2(d_553_i0_, globalState) if default__.fm2((self).f3, globalState) else True)
                rhs123_ = p0
                rhs124_ = d_636_v56_
                rhs125_ = d_632_cf28_
                lhs113_ = globalState
                lhs114_ = globalState
                lhs115_ = (self).f15
                lhs116_ = default__.safeIndex(354, ((self).f15).length(0))
                lhs117_ = p1
                lhs118_ = default__.safeIndex(686, (p1).length(0))
                lhs113_.f1 = rhs121_
                lhs114_.f2 = rhs122_
                lhs115_[lhs116_] = rhs123_
                d_636_v56_ = rhs124_
                lhs117_[lhs118_] = rhs125_
                index124_ = default__.safeIndex(354, ((self).f15).length(0))
                ((self).f15)[index124_] = default__.safeDivisionInt(default__.fm0(globalState), (((self).f15)[default__.safeIndex(354, ((self).f15).length(0))] if d_631_cf29_ else (self).f14))
                d_638_v58_: C0
                nw121_ = C0()
                nw121_.ctor__(d_630_cf30_, (p1)[default__.safeIndex(686, (p1).length(0))])
                d_638_v58_ = nw121_
            elif True:
                d_639___mcc_h11_ = source7_.cf27
                d_640_cf27_ = d_639___mcc_h11_
                d_641_v59_: _dafny.Seq
                d_641_v59_ = _dafny.SeqWithoutIsStrInference([self.f4, False])
                r0 = len((d_641_v59_) + ((d_641_v59_) + (d_641_v59_)))
                d_642_v60_: _dafny.Seq
                d_642_v60_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yqnnbe"))
                r0 = len(d_642_v60_)
                d_643_v61_: _dafny.Map
                d_643_v61_ = _dafny.Map({self.f4: self.f4})
                d_644_v62_: D6
                d_644_v62_ = D6_DC18(self.f4, p0)
                d_645_v63_: _dafny.Set
                d_645_v63_ = _dafny.Set({d_553_i0_, d_553_i0_, len(d_643_v61_), (d_644_v62_).cf25, d_553_i0_})
                d_646_v64_: _dafny.Array
                nw122_ = _dafny.Array(None, 15)
                nw122_[int(0)] = self.f4
                nw122_[int(1)] = (len(d_645_v63_)) < (656)
                nw122_[int(2)] = self.f4
                nw122_[int(3)] = (default__.fm2((0) - ((0) - (default__.fm0(globalState))), globalState)) in (d_643_v61_)
                nw122_[int(4)] = (self.f4 if self.f4 else not(True))
                nw122_[int(5)] = ((self).f3) <= (p0)
                nw122_[int(6)] = self.f4
                nw122_[int(7)] = self.f4
                nw122_[int(8)] = not(False)
                nw122_[int(9)] = (not(self.f4)) in (d_643_v61_)
                nw122_[int(10)] = self.f4
                nw122_[int(11)] = self.f4
                nw122_[int(12)] = self.f4
                nw122_[int(13)] = self.f4
                nw122_[int(14)] = self.f4
                d_646_v64_ = nw122_
                index125_ = default__.safeIndex(330, (d_646_v64_).length(0))
                (d_646_v64_)[index125_] = self.f4
                index126_ = default__.safeIndex(330, (d_646_v64_).length(0))
                (d_646_v64_)[index126_] = self.f4
                d_646_v64_ = d_646_v64_
            d_647_v65_: _dafny.Seq
            d_647_v65_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "png"))
            d_648_v66_: _dafny.Map
            d_648_v66_ = _dafny.Map({len(d_647_v65_): self.f4})
            d_649_v67_: C1
            nw123_ = C1()
            nw123_.ctor__((0) - (d_553_i0_), len(d_648_v66_), self.f4)
            d_649_v67_ = nw123_
        d_650_v68_: str
        d_650_v68_ = _dafny.CodePoint('l')
        pat_let_tv6_ = d_650_v68_
        pat_let_tv7_ = d_650_v68_
        pat_let_tv8_ = d_650_v68_
        def lambda33_(source8_):
            if source8_.is_DC6:
                d_651___mcc_h12_ = source8_.cf9
                d_652___mcc_h13_ = source8_.cf10
                d_653___mcc_h14_ = source8_.cf11
                d_654___mcc_h15_ = source8_.cf12
                d_655_cf12_ = d_654___mcc_h15_
                d_656_cf11_ = d_653___mcc_h14_
                d_657_cf10_ = d_652___mcc_h13_
                d_658_cf9_ = d_651___mcc_h12_
                if d_655_cf12_:
                    return pat_let_tv6_
                elif True:
                    return _dafny.CodePoint('c')
            elif source8_.is_DC7:
                return pat_let_tv7_
            elif True:
                d_659___mcc_h16_ = source8_.cf8
                d_660_cf8_ = d_659___mcc_h16_
                return pat_let_tv8_

        d_650_v68_ = lambda33_(D1_DC5((self).f14))
        r0 = (p0) * ((self).f14)
        (globalState).f2 = self.f4
        r0 = ((self).f14 if (906) < (p0) else default__.safeDivisionInt((self).f14, (0) - ((self).f3)))
        r1 = default__.fm2(default__.safeModuloInt((self).f14, p0), globalState)
        return r0, r1

    @property
    def f14(self):
        return self._f14
    @property
    def f15(self):
        return self._f15

class C3(T0, T1):
    def  __init__(self):
        self._f4: bool = False
        self._f7: int = int(0)
        self._f3: int = int(0)
        self.f13: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f12: str = _dafny.CodePoint('D')
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f4(self):
        return self._f4
    @f4.setter
    def f4(self, value):
        self._f4 = value
    @property
    def f7(self):
        return self._f7
    @property
    def f3(self):
        return self._f3
    def ctor__(self, f12, f13, f3, f4, f7):
        (self)._f12 = f12
        (self).f13 = f13
        (self)._f3 = f3
        (self).f4 = f4
        (self)._f7 = f7

    def fm12(self, p0, p1, globalState):
        return default__.safeModuloInt((self).f7, 14)

    def fm13(self, globalState):
        return self.f4

    def m1(self, p0, globalState):
        r0: int = int(0)
        d_661_v0_: _dafny.Set
        d_661_v0_ = _dafny.Set({((self).f7) - ((self).f7), (self).f7, len(self.f13)})
        d_661_v0_ = _dafny.Set({default__.safeDivisionInt((self).f3, (self).f7)})
        d_662_v1_: _dafny.Map
        d_662_v1_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([default__.fm14(not(self.f4), (self).fm13(globalState), globalState)]): self.f4})
        d_663_v2_: _dafny.MultiSet
        d_663_v2_ = _dafny.MultiSet([True])
        (globalState).f2 = ((d_662_v1_)[_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([self.f4]), d_663_v2_])] if (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([self.f4]), d_663_v2_])) in (d_662_v1_) else ((self).f3) < (p0))
        d_664_i0_: int
        d_664_i0_ = 0
        with _dafny.label("6"):
            while self.f4:
                with _dafny.c_label("6"):
                    if (d_664_i0_) >= (100):
                        raise _dafny.Break("6")
                    d_664_i0_ = (d_664_i0_) + (1)
                    (self).m5(globalState)
                    (globalState).f2 = not((self).fm13(globalState))
                    if self.f4:
                        d_665_v3_: D0
                        d_665_v3_ = D0_DC0(self.f4)
                        d_666_v4_: _dafny.Array
                        nw124_ = _dafny.Array(None, 23)
                        nw124_[int(0)] = not(False)
                        nw124_[int(1)] = self.f4
                        nw124_[int(2)] = (d_665_v3_).cf0
                        nw124_[int(3)] = self.f4
                        nw124_[int(4)] = self.f4
                        nw124_[int(5)] = self.f4
                        nw124_[int(6)] = self.f4
                        nw124_[int(7)] = self.f4
                        nw124_[int(8)] = self.f4
                        nw124_[int(9)] = self.f4
                        nw124_[int(10)] = self.f4
                        nw124_[int(11)] = self.f4
                        nw124_[int(12)] = False
                        nw124_[int(13)] = self.f4
                        nw124_[int(14)] = self.f4
                        nw124_[int(15)] = self.f4
                        nw124_[int(16)] = self.f4
                        nw124_[int(17)] = False
                        nw124_[int(18)] = False
                        nw124_[int(19)] = self.f4
                        nw124_[int(20)] = self.f4
                        nw124_[int(21)] = False
                        nw124_[int(22)] = self.f4
                        d_666_v4_ = nw124_
                        d_667_v5_: C0
                        nw125_ = C0()
                        nw125_.ctor__((self.f4 if not(self.f4) else self.f4), d_666_v4_)
                        d_667_v5_ = nw125_
                        d_668_v6_: _dafny.Seq
                        d_668_v6_ = _dafny.SeqWithoutIsStrInference([(self).f3])
                        d_669_v7_: C0
                        nw126_ = C0()
                        nw126_.ctor__((len((default__.fm15(not(d_667_v5_.f10), (self).f3, len(d_668_v6_), globalState)).set(default__.safeIndex(p0, len(default__.fm15(not(d_667_v5_.f10), (self).f3, len(d_668_v6_), globalState))), (self).f12))) == (p0), (d_667_v5_).f11)
                        d_669_v7_ = nw126_
                        d_670_v8_: D4
                        d_670_v8_ = D4_DC11(d_661_v0_)
                        d_671_v9_: D4
                        d_671_v9_ = D4_DC11((d_670_v8_).cf18)
                        d_661_v0_ = (d_670_v8_).cf18
                        d_672_v10_: _dafny.Array
                        def lambda34_(d_673_p0_):
                            def lambda35_(d_674_i1_):
                                return (d_674_i1_) * (d_673_p0_)

                            return lambda35_

                        init23_ = lambda34_(p0)
                        nw127_ = _dafny.Array(None, 5)
                        for i0_23_ in range(nw127_.length(0)):
                            nw127_[i0_23_] = init23_(i0_23_)
                        d_672_v10_ = nw127_
                        d_672_v10_ = d_672_v10_
                        d_675_v11_: _dafny.MultiSet
                        d_675_v11_ = _dafny.MultiSet([(d_669_v7_).f11, (d_669_v7_).f11, (d_669_v7_).f11, (d_667_v5_).f11])
                        d_676_v12_: _dafny.Map
                        d_676_v12_ = _dafny.Map({d_669_v7_.f10: (d_669_v7_).fm5(292, p0, (self).f12, d_669_v7_.f10, globalState)})
                        d_677_v13_: _dafny.Map
                        d_677_v13_ = _dafny.Map({len(d_676_v12_): d_667_v5_.f10})
                        rhs126_ = default__.safeDivisionInt((self).f3, (len(d_677_v13_)) - (p0))
                        rhs127_ = (d_675_v11_) - (d_675_v11_)
                        rhs128_ = (self).f7
                        rhs129_ = d_672_v10_
                        rhs130_ = d_668_v6_
                        lhs119_ = globalState
                        r0 = rhs126_
                        d_675_v11_ = rhs127_
                        lhs119_.f1 = rhs128_
                        d_672_v10_ = rhs129_
                        d_668_v6_ = rhs130_
                    elif True:
                        d_678_v14_: _dafny.Array
                        nw128_ = _dafny.Array(None, 7)
                        nw128_[int(0)] = (self).f12
                        nw128_[int(1)] = (self).f12
                        nw128_[int(2)] = (self).f12
                        nw128_[int(3)] = (self).f12
                        nw128_[int(4)] = (self).f12
                        nw128_[int(5)] = (self).f12
                        nw128_[int(6)] = (self).f12
                        d_678_v14_ = nw128_
                        d_679_v15_: _dafny.Seq
                        d_679_v15_ = _dafny.SeqWithoutIsStrInference([d_678_v14_])
                        d_680_v16_: _dafny.MultiSet
                        d_680_v16_ = _dafny.MultiSet([(self).f3])
                        d_681_v17_: _dafny.Array
                        nw129_ = _dafny.Array(None, 13)
                        nw129_[int(0)] = d_678_v14_
                        nw129_[int(1)] = d_678_v14_
                        nw129_[int(2)] = d_678_v14_
                        nw129_[int(3)] = (d_679_v15_)[default__.safeIndex((d_680_v16_).cardinality, len(d_679_v15_))]
                        nw129_[int(4)] = d_678_v14_
                        nw129_[int(5)] = (d_679_v15_)[default__.safeIndex((self).f7, len(d_679_v15_))]
                        nw129_[int(6)] = d_678_v14_
                        nw129_[int(7)] = d_678_v14_
                        nw129_[int(8)] = d_678_v14_
                        nw129_[int(9)] = d_678_v14_
                        nw129_[int(10)] = d_678_v14_
                        nw129_[int(11)] = d_678_v14_
                        nw129_[int(12)] = d_678_v14_
                        d_681_v17_ = nw129_
                        index127_ = default__.safeIndex(448, (d_681_v17_).length(0))
                        (d_681_v17_)[index127_] = d_678_v14_
                        d_682_v18_: D3
                        d_682_v18_ = D3_DC9(d_678_v14_)
                        index128_ = default__.safeIndex(448, (d_681_v17_).length(0))
                        rhs131_ = (d_682_v18_).cf14
                        rhs132_ = (self.f4) or (self.f4)
                        lhs120_ = d_681_v17_
                        lhs121_ = default__.safeIndex(448, (d_681_v17_).length(0))
                        lhs122_ = globalState
                        lhs120_[lhs121_] = rhs131_
                        lhs122_.f2 = rhs132_
                        d_683_v19_: _dafny.Array
                        nw130_ = _dafny.Array(False, 11)
                        d_683_v19_ = nw130_
                        index129_ = default__.safeIndex(379, (d_683_v19_).length(0))
                        (d_683_v19_)[index129_] = self.f4
                        d_684_v20_: _dafny.Seq
                        d_684_v20_ = _dafny.SeqWithoutIsStrInference([p0])
                        d_685_v21_: D1
                        d_685_v21_ = D1_DC5((self).f7)
                        d_686_v22_: _dafny.Map
                        d_686_v22_ = _dafny.Map({d_685_v21_: (self).f3})
                        d_687_v23_: D1
                        d_687_v23_ = D1_DC5(((d_686_v22_)[d_685_v21_] if (d_685_v21_) in (d_686_v22_) else (self).f7))
                        d_688_v24_: _dafny.Seq
                        d_688_v24_ = _dafny.SeqWithoutIsStrInference([self.f4, self.f4, self.f4])
                        d_689_v25_: _dafny.Map
                        d_689_v25_ = _dafny.Map({(d_688_v24_)[default__.safeIndex(p0, len(d_688_v24_))]: p0})
                        d_690_v26_: _dafny.Map
                        d_690_v26_ = _dafny.Map({(self).f3: d_689_v25_})
                        d_691_v27_: _dafny.Array
                        nw131_ = _dafny.Array(None, 21)
                        nw131_[int(0)] = (self).f7
                        nw131_[int(1)] = (self).f3
                        nw131_[int(2)] = (self).f7
                        nw131_[int(3)] = (self).f3
                        nw131_[int(4)] = p0
                        nw131_[int(5)] = ((self).f7) - (len(d_684_v20_))
                        nw131_[int(6)] = p0
                        nw131_[int(7)] = ((self).f7) - (166)
                        nw131_[int(8)] = (self).f7
                        nw131_[int(9)] = (d_685_v21_).cf8
                        nw131_[int(10)] = default__.fm0(globalState)
                        nw131_[int(11)] = -367
                        nw131_[int(12)] = (407) * ((self).f3)
                        nw131_[int(13)] = len(d_690_v26_)
                        nw131_[int(14)] = default__.safeModuloInt((D1_DC5((self).f7)).cf8, p0)
                        nw131_[int(15)] = -272
                        nw131_[int(16)] = p0
                        nw131_[int(17)] = default__.safeModuloInt(len(d_688_v24_), (self).f3)
                        nw131_[int(18)] = (self).f7
                        nw131_[int(19)] = default__.safeModuloInt((self).f7, (self).f7)
                        nw131_[int(20)] = (self).f3
                        d_691_v27_ = nw131_
                        index130_ = default__.safeIndex(379, (d_683_v19_).length(0))
                        rhs133_ = self.f4
                        rhs134_ = d_691_v27_
                        lhs123_ = d_683_v19_
                        lhs124_ = default__.safeIndex(379, (d_683_v19_).length(0))
                        lhs123_[lhs124_] = rhs133_
                        d_691_v27_ = rhs134_
                        d_692_v28_: C0
                        nw132_ = C0()
                        nw132_.ctor__((d_683_v19_)[default__.safeIndex(379, (d_683_v19_).length(0))], d_683_v19_)
                        d_692_v28_ = nw132_
                        (globalState).f1 = p0
                        d_693_v29_: _dafny.Array
                        def lambda36_(d_694_v0_):
                            def lambda37_(d_695_i2_):
                                return d_694_v0_

                            return lambda37_

                        init24_ = lambda36_(d_661_v0_)
                        nw133_ = _dafny.Array(None, 2)
                        for i0_24_ in range(nw133_.length(0)):
                            nw133_[i0_24_] = init24_(i0_24_)
                        d_693_v29_ = nw133_
                        d_696_v30_: _dafny.Array
                        nw134_ = _dafny.Array(D3.default()(), 3)
                        d_696_v30_ = nw134_
                        d_697_v31_: D3
                        d_697_v31_ = D3_DC10((0) - (p0), (d_692_v28_).f11, (self).f7)
                        index131_ = default__.safeIndex(84, (d_696_v30_).length(0))
                        (d_696_v30_)[index131_] = d_697_v31_
                        d_698_v32_: _dafny.Seq
                        d_698_v32_ = _dafny.SeqWithoutIsStrInference([self.f13])
                        d_699_v33_: _dafny.MultiSet
                        d_699_v33_ = _dafny.MultiSet([self.f13, self.f13])
                        pat_let_tv9_ = d_683_v19_
                        index132_ = default__.safeIndex(379, (d_683_v19_).length(0))
                        index133_ = default__.safeIndex(84, (d_696_v30_).length(0))
                        def iife34_(_pat_let12_0):
                            def iife35_(d_700_dt__update__tmp_h0_):
                                def iife36_(_pat_let13_0):
                                    def iife37_(d_701_dt__update_hcf15_h0_):
                                        def iife38_(_pat_let14_0):
                                            def iife39_(d_702_dt__update_hcf16_h0_):
                                                return D3_DC10(d_701_dt__update_hcf15_h0_, d_702_dt__update_hcf16_h0_, (d_700_dt__update__tmp_h0_).cf17)
                                            return iife39_(_pat_let14_0)
                                        return iife38_(pat_let_tv9_)
                                    return iife37_(_pat_let13_0)
                                return iife36_((997) + ((self).f3))
                            return iife35_(_pat_let12_0)
                        rhs135_ = d_693_v29_
                        rhs136_ = ((self).f7) >= (((_dafny.MultiSet(d_698_v32_) if True else d_699_v33_)).cardinality)
                        rhs137_ = (p0) == (default__.safeModuloInt(p0, (self).f7))
                        rhs138_ = iife34_(d_697_v31_)
                        lhs125_ = d_683_v19_
                        lhs126_ = default__.safeIndex(379, (d_683_v19_).length(0))
                        lhs127_ = globalState
                        lhs128_ = d_696_v30_
                        lhs129_ = default__.safeIndex(84, (d_696_v30_).length(0))
                        d_693_v29_ = rhs135_
                        lhs125_[lhs126_] = rhs136_
                        lhs127_.f2 = rhs137_
                        lhs128_[lhs129_] = rhs138_
                    d_703_v34_: _dafny.Array
                    nw135_ = _dafny.Array(int(0), 25)
                    d_703_v34_ = nw135_
                    index134_ = default__.safeIndex(57, (d_703_v34_).length(0))
                    (d_703_v34_)[index134_] = p0
                    index135_ = default__.safeIndex(57, (d_703_v34_).length(0))
                    (d_703_v34_)[index135_] = (self).f7
                    pass
            pass
        d_704_i3_: int
        d_704_i3_ = 0
        with _dafny.label("7"):
            while self.f4:
                with _dafny.c_label("7"):
                    if (d_704_i3_) >= (100):
                        raise _dafny.Break("7")
                    d_704_i3_ = (d_704_i3_) + (1)
                    d_705_v35_: _dafny.Seq
                    d_705_v35_ = _dafny.SeqWithoutIsStrInference([p0])
                    d_706_v36_: _dafny.Map
                    d_706_v36_ = _dafny.Map({not(self.f4): d_705_v35_})
                    d_706_v36_ = (d_706_v36_).set(not(self.f4), d_705_v35_)
                    d_661_v0_ = d_661_v0_
                    d_707_v37_: _dafny.Map
                    d_707_v37_ = _dafny.Map({self.f4: 637})
                    d_708_v38_: _dafny.Array
                    nw136_ = _dafny.Array(None, 5)
                    nw136_[int(0)] = (not(self.f4)) in (d_707_v37_)
                    nw136_[int(1)] = True
                    nw136_[int(2)] = self.f4
                    nw136_[int(3)] = True
                    nw136_[int(4)] = self.f4
                    d_708_v38_ = nw136_
                    index136_ = default__.safeIndex(436, (d_708_v38_).length(0))
                    (d_708_v38_)[index136_] = not(self.f4)
                    d_709_v39_: _dafny.Map
                    d_709_v39_ = _dafny.Map({546: (self).f3})
                    index137_ = default__.safeIndex(436, (d_708_v38_).length(0))
                    (d_708_v38_)[index137_] = ((self).f7) == (len(d_709_v39_))
                    d_710_v40_: C0
                    nw137_ = C0()
                    nw137_.ctor__(((self).f7) >= ((self).f3), d_708_v38_)
                    d_710_v40_ = nw137_
                    pass
            pass
        (globalState).f2 = False
        d_711_v41_: D1
        d_711_v41_ = D1_DC7()
        def lambda38_(source9_):
            if source9_.is_DC6:
                d_712___mcc_h0_ = source9_.cf9
                d_713___mcc_h1_ = source9_.cf10
                d_714___mcc_h2_ = source9_.cf11
                d_715___mcc_h3_ = source9_.cf12
                d_716_cf12_ = d_715___mcc_h3_
                d_717_cf11_ = d_714___mcc_h2_
                d_718_cf10_ = d_713___mcc_h1_
                d_719_cf9_ = d_712___mcc_h0_
                if not(False):
                    return False
                elif True:
                    return True
            elif source9_.is_DC7:
                return self.f4
            elif True:
                d_720___mcc_h4_ = source9_.cf8
                d_721_cf8_ = d_720___mcc_h4_
                return (len(_dafny.SeqWithoutIsStrInference([d_721_cf8_ for d_722_i4_ in range(default__.abs(300))]))) in (_dafny.MultiSet([(self).f3]))

        rhs139_ = (self).fm12(default__.fm2(len(d_661_v0_), globalState), self.f4, globalState)
        rhs140_ = lambda38_(d_711_v41_)
        rhs141_ = ((p0) - (len(self.f13))) * ((self).f3)
        lhs130_ = globalState
        lhs131_ = globalState
        lhs132_ = globalState
        lhs130_.f1 = rhs139_
        lhs131_.f2 = rhs140_
        lhs132_.f1 = rhs141_
        d_723_v42_: _dafny.Map
        d_723_v42_ = _dafny.Map({d_663_v2_: (self).f7})
        r0 = default__.safeModuloInt(default__.safeDivisionInt(((d_723_v42_)[_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([self.f4]))] if (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([self.f4]))) in (d_723_v42_) else (self).f3), len(d_661_v0_)), (self).f3)
        return r0

    def m5(self, globalState):
        hi6_ = default__.safeDivisionInt((self).f7, (self).f3)
        for d_724_i0_ in range((self).f7, hi6_):
            d_725_v0_: _dafny.Array
            nw138_ = _dafny.Array(False, 25)
            d_725_v0_ = nw138_
            d_726_v1_: C0
            nw139_ = C0()
            nw139_.ctor__(not(((0) - ((self).f7)) < (len(_dafny.SeqWithoutIsStrInference([(self).f12 for d_727_i1_ in range(default__.abs(591))])))), d_725_v0_)
            d_726_v1_ = nw139_
            (globalState).f1 = (self).f3
            d_728_v2_: _dafny.Seq
            d_728_v2_ = _dafny.SeqWithoutIsStrInference([d_726_v1_.f10])
            d_729_v3_: _dafny.MultiSet
            d_729_v3_ = _dafny.MultiSet([d_728_v2_, d_728_v2_])
            (globalState).f1 = (((d_729_v3_).intersection(d_729_v3_)) - ((default__.fm16(self.f4, self.f4, globalState)) - (default__.fm16(d_726_v1_.f10, True, globalState)))).cardinality
            d_730_v4_: _dafny.Map
            d_730_v4_ = _dafny.Map({d_726_v1_.f10: self.f4})
            d_731_v5_: C0
            nw140_ = C0()
            nw140_.ctor__(((d_730_v4_)[self.f4] if (self.f4) in (d_730_v4_) else d_726_v1_.f10), (d_726_v1_).f11)
            d_731_v5_ = nw140_
        d_732_v6_: _dafny.Array
        nw141_ = _dafny.Array(int(0), 28)
        d_732_v6_ = nw141_
        d_733_v7_: D0
        d_733_v7_ = D0_DC0((_dafny.SeqWithoutIsStrInference([d_732_v6_, d_732_v6_, d_732_v6_])) == (_dafny.SeqWithoutIsStrInference([d_732_v6_])))
        source10_ = d_733_v7_
        if source10_.is_DC1:
            d_734___mcc_h0_ = source10_.cf1
            d_735___mcc_h1_ = source10_.cf2
            d_736___mcc_h2_ = source10_.cf3
            d_737___mcc_h3_ = source10_.cf4
            d_738_cf4_ = d_737___mcc_h3_
            d_739_cf3_ = d_736___mcc_h2_
            d_740_cf2_ = d_735___mcc_h1_
            d_741_cf1_ = d_734___mcc_h0_
            (globalState).f1 = -77
            d_742_v8_: _dafny.Array
            def lambda39_(d_743_cf1_):
                def lambda40_(d_744_i2_):
                    return d_743_cf1_

                return lambda40_

            init25_ = lambda39_(d_741_cf1_)
            nw142_ = _dafny.Array(None, 11)
            for i0_25_ in range(nw142_.length(0)):
                nw142_[i0_25_] = init25_(i0_25_)
            d_742_v8_ = nw142_
            index138_ = default__.safeIndex(967, (d_742_v8_).length(0))
            (d_742_v8_)[index138_] = d_740_cf2_
            d_745_v9_: _dafny.Seq
            d_745_v9_ = _dafny.SeqWithoutIsStrInference([d_732_v6_])
            d_746_v10_: _dafny.MultiSet
            d_746_v10_ = _dafny.MultiSet([(d_745_v9_)[default__.safeIndex((self).f3, len(d_745_v9_))]])
            d_747_v11_: _dafny.Map
            d_747_v11_ = _dafny.Map({not(d_740_cf2_): len((self.f13).set(default__.safeIndex((0) - ((d_746_v10_).cardinality), len(self.f13)), (self).f12))})
            index139_ = default__.safeIndex(967, (d_742_v8_).length(0))
            rhs142_ = d_741_cf1_
            rhs143_ = d_747_v11_
            lhs133_ = d_742_v8_
            lhs134_ = default__.safeIndex(967, (d_742_v8_).length(0))
            lhs133_[lhs134_] = rhs142_
            d_747_v11_ = rhs143_
            (globalState).f2 = ((self).f7) <= (((self).f7) - ((self).f7))
            index140_ = default__.safeIndex(967, (d_742_v8_).length(0))
            (d_742_v8_)[index140_] = d_740_cf2_
        elif source10_.is_DC2:
            d_748___mcc_h4_ = source10_.cf5
            d_749___mcc_h5_ = source10_.cf6
            d_750_cf6_ = d_749___mcc_h5_
            d_751_cf5_ = d_748___mcc_h4_
            d_752_v12_: _dafny.Seq
            d_752_v12_ = _dafny.SeqWithoutIsStrInference([d_750_cf6_, d_750_cf6_])
            d_753_v13_: _dafny.Seq
            d_753_v13_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([self.f4]), (d_752_v12_).set(default__.safeIndex((self).f7, len(d_752_v12_)), self.f4), d_752_v12_, d_752_v12_])
            d_754_v14_: D0
            d_754_v14_ = D0_DC1(False, self.f4, (self).f12, d_732_v6_)
            d_755_v15_: _dafny.Map
            d_755_v15_ = _dafny.Map({(d_753_v13_)[default__.safeIndex(len(default__.fm15(d_750_cf6_, (0) - ((self).f3), len(self.f13), globalState)), len(d_753_v13_))]: (self.f13) + ((self.f13).set(default__.safeIndex((0) - ((self).f3), len(self.f13)), (d_754_v14_).cf3))})
            d_755_v15_ = d_755_v15_
            (self).f4 = d_750_cf6_
            d_756_v16_: _dafny.Array
            def lambda41_(d_757_cf6_):
                def lambda42_(d_758_i3_):
                    return d_757_cf6_

                return lambda42_

            init26_ = lambda41_(d_750_cf6_)
            nw143_ = _dafny.Array(None, 27)
            for i0_26_ in range(nw143_.length(0)):
                nw143_[i0_26_] = init26_(i0_26_)
            d_756_v16_ = nw143_
            d_759_v17_: C0
            nw144_ = C0()
            nw144_.ctor__(self.f4, d_756_v16_)
            d_759_v17_ = nw144_
            d_760_v18_: _dafny.Map
            d_760_v18_ = _dafny.Map({(self).f12: d_750_cf6_})
            d_761_v20_: _dafny.Seq
            def iife40_():
                coll10_ = _dafny.Map()
                compr_10_: str
                for compr_10_ in (self.f13).Elements:
                    d_762_v19_: str = compr_10_
                    if (d_762_v19_) in (self.f13):
                        coll10_[d_762_v19_] = d_759_v17_.f10
                return _dafny.Map(coll10_)
            d_761_v20_ = _dafny.SeqWithoutIsStrInference([d_760_v18_, (iife40_()
            ).set((self).f12, self.f4), default__.fm17(_dafny.Map({d_759_v17_.f10: d_759_v17_.f10}), globalState), d_760_v18_, d_760_v18_])
            (globalState).f1 = (len((d_761_v20_)[default__.safeIndex((self).f7, len(d_761_v20_))])) + ((self).f3)
        elif source10_.is_DC3:
            d_763_v21_: _dafny.Array
            nw145_ = _dafny.Array(_dafny.CodePoint('D'), 25)
            d_763_v21_ = nw145_
            d_764_v22_: D3
            d_764_v22_ = D3_DC9(d_763_v21_)
            d_765_v23_: _dafny.Seq
            d_765_v23_ = _dafny.SeqWithoutIsStrInference([d_764_v22_, d_764_v22_, d_764_v22_, d_764_v22_])
            d_766_v24_: _dafny.Array
            nw146_ = _dafny.Array(None, 29)
            nw146_[int(0)] = d_765_v23_
            nw146_[int(1)] = _dafny.SeqWithoutIsStrInference([d_764_v22_])
            nw146_[int(2)] = d_765_v23_
            nw146_[int(3)] = (d_765_v23_) + (d_765_v23_)
            nw146_[int(4)] = _dafny.SeqWithoutIsStrInference([d_764_v22_, d_764_v22_, d_764_v22_])
            nw146_[int(5)] = d_765_v23_
            nw146_[int(6)] = d_765_v23_
            nw146_[int(7)] = d_765_v23_
            nw146_[int(8)] = d_765_v23_
            nw146_[int(9)] = d_765_v23_
            nw146_[int(10)] = ((d_765_v23_) + (d_765_v23_)).set(default__.safeIndex((self).f7, len((d_765_v23_) + (d_765_v23_))), d_764_v22_)
            nw146_[int(11)] = _dafny.SeqWithoutIsStrInference([d_764_v22_])
            nw146_[int(12)] = d_765_v23_
            nw146_[int(13)] = d_765_v23_
            nw146_[int(14)] = d_765_v23_
            nw146_[int(15)] = d_765_v23_
            nw146_[int(16)] = _dafny.SeqWithoutIsStrInference([d_764_v22_, d_764_v22_])
            nw146_[int(17)] = d_765_v23_
            nw146_[int(18)] = _dafny.SeqWithoutIsStrInference([d_764_v22_])
            nw146_[int(19)] = d_765_v23_
            nw146_[int(20)] = (d_765_v23_) + (d_765_v23_)
            nw146_[int(21)] = _dafny.SeqWithoutIsStrInference([d_764_v22_])
            nw146_[int(22)] = _dafny.SeqWithoutIsStrInference([d_764_v22_, d_764_v22_])
            nw146_[int(23)] = (d_765_v23_) + (d_765_v23_)
            nw146_[int(24)] = d_765_v23_
            nw146_[int(25)] = _dafny.SeqWithoutIsStrInference([D3_DC9(d_763_v21_)])
            nw146_[int(26)] = _dafny.SeqWithoutIsStrInference([d_764_v22_, D3_DC9(d_763_v21_), d_764_v22_])
            nw146_[int(27)] = d_765_v23_
            nw146_[int(28)] = d_765_v23_
            d_766_v24_ = nw146_
            index141_ = default__.safeIndex(83, (d_766_v24_).length(0))
            (d_766_v24_)[index141_] = d_765_v23_
            index142_ = default__.safeIndex(83, (d_766_v24_).length(0))
            (d_766_v24_)[index142_] = (d_765_v23_ if self.f4 else (d_765_v23_) + (d_765_v23_))
            d_767_v25_: _dafny.Map
            d_767_v25_ = _dafny.Map({True: (self).f7})
            d_768_v26_: _dafny.Seq
            d_768_v26_ = _dafny.SeqWithoutIsStrInference([(d_767_v25_) | (d_767_v25_), d_767_v25_])
            d_769_v27_: _dafny.Seq
            d_769_v27_ = _dafny.SeqWithoutIsStrInference([False])
            d_770_v28_: _dafny.Map
            d_770_v28_ = _dafny.Map({default__.fm0(globalState): self.f4})
            d_771_v30_: _dafny.MultiSet
            d_771_v30_ = _dafny.MultiSet([(self).f12, (self).f12])
            d_772_v31_: D0
            d_772_v31_ = D0_DC2(d_771_v30_, (D0_DC2(d_771_v30_, self.f4)).cf6)
            d_773_v32_: T0
            nw147_ = C2()
            nw147_.ctor__(704, d_732_v6_, (self).f7, self.f4)
            d_773_v32_ = nw147_
            d_774_v33_: _dafny.Set
            d_774_v33_ = _dafny.Set({False})
            d_775_v34_: D8
            d_775_v34_ = D8_DC23(self.f4, d_773_v32_.f4, d_773_v32_.f4, (self).f12)
            d_776_v35_: _dafny.Map
            d_776_v35_ = _dafny.Map({d_773_v32_.f4: self.f4})
            d_777_v36_: _dafny.Array
            nw148_ = _dafny.Array(None, 22)
            nw148_[int(0)] = self.f4
            nw148_[int(1)] = self.f4
            nw148_[int(2)] = self.f4
            nw148_[int(3)] = self.f4
            nw148_[int(4)] = (self.f4) and (True)
            nw148_[int(5)] = self.f4
            nw148_[int(6)] = (d_769_v27_)[default__.safeIndex((self).f7, len(d_769_v27_))]
            nw148_[int(7)] = ((self).f7) not in (d_770_v28_)
            nw148_[int(8)] = self.f4
            nw148_[int(9)] = self.f4
            nw148_[int(10)] = self.f4
            nw148_[int(11)] = ((default__.fm14(self.f4, self.f4, globalState)).cardinality) == (-524)
            nw148_[int(12)] = self.f4
            nw148_[int(13)] = (-900) >= (-434)
            def iife41_():
                coll11_ = _dafny.Set()
                compr_11_: str
                for compr_11_ in (self.f13).Elements:
                    d_778_v29_: str = compr_11_
                    if (d_778_v29_) in (self.f13):
                        coll11_ = coll11_.union(_dafny.Set([d_778_v29_]))
                return _dafny.Set(coll11_)
            nw148_[int(14)] = (default__.fm18(d_772_v31_, (D1_DC6(d_773_v32_, default__.fm20(len(d_774_v33_), globalState), self.f4, d_773_v32_.f4)).cf11, d_773_v32_.f4, globalState)).issubset(iife41_()
            )
            nw148_[int(15)] = d_773_v32_.f4
            nw148_[int(16)] = d_773_v32_.f4
            nw148_[int(17)] = (d_775_v34_).cf32
            nw148_[int(18)] = d_773_v32_.f4
            nw148_[int(19)] = self.f4
            nw148_[int(20)] = ((d_776_v35_)[self.f4] if (self.f4) in (d_776_v35_) else self.f4)
            nw148_[int(21)] = self.f4
            d_777_v36_ = nw148_
            d_779_v37_: _dafny.Map
            d_779_v37_ = _dafny.Map({self.f4: self.f13})
            index143_ = default__.safeIndex(521, (d_777_v36_).length(0))
            (d_777_v36_)[index143_] = (len(d_779_v37_)) >= (len(d_774_v33_))
            d_780_v38_: _dafny.Seq
            d_780_v38_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kmrehf")))])
            index144_ = default__.safeIndex(521, (d_777_v36_).length(0))
            rhs144_ = (((self).f7 if (d_775_v34_).cf32 else 900)) * (len(default__.fm27(globalState)))
            rhs145_ = (d_768_v26_).set(default__.safeIndex((d_773_v32_).f3, len(d_768_v26_)), d_767_v25_)
            rhs146_ = ((self.f13)[default__.safeIndex((d_773_v32_).f3, len(self.f13))]) in (default__.fm15(self.f4, 318, (0) - (len(d_780_v38_)), globalState))
            lhs135_ = globalState
            lhs136_ = d_777_v36_
            lhs137_ = default__.safeIndex(521, (d_777_v36_).length(0))
            lhs135_.f1 = rhs144_
            d_768_v26_ = rhs145_
            lhs136_[lhs137_] = rhs146_
            (self).f4 = (self).fm13(globalState)
            d_781_v39_: str
            d_781_v39_ = _dafny.CodePoint('l')
            d_781_v39_ = (self).f12
        elif source10_.is_DC0:
            d_782___mcc_h6_ = source10_.cf0
            d_783_cf0_ = d_782___mcc_h6_
            d_784_v40_: _dafny.Array
            nw149_ = _dafny.Array(False, 6)
            d_784_v40_ = nw149_
            d_785_v41_: _dafny.Map
            d_785_v41_ = _dafny.Map({d_784_v40_: len(_dafny.SeqWithoutIsStrInference([(self).f7]))})
            d_785_v41_ = (d_785_v41_).set(d_784_v40_, (self).f3)
            if (self).fm13(globalState):
                d_786_v42_: D3
                d_786_v42_ = D3_DC10((self).f3, d_784_v40_, ((self).f7) * ((self).f3))
                d_786_v42_ = d_786_v42_
                d_787_v43_: _dafny.Seq
                d_787_v43_ = _dafny.SeqWithoutIsStrInference([(D0_DC0(d_783_cf0_)).cf0])
                d_788_v44_: _dafny.Map
                d_788_v44_ = _dafny.Map({(self).f12: default__.fm2(len(d_787_v43_), globalState)})
                d_789_v45_: C2
                nw150_ = C2()
                nw150_.ctor__((0) - ((self).fm12(d_783_cf0_, True, globalState)), d_732_v6_, len((d_788_v44_) | (d_788_v44_)), self.f4)
                d_789_v45_ = nw150_
                (globalState).f2 = (not(d_783_cf0_)) or (d_783_cf0_)
                d_790_v46_: _dafny.Map
                d_790_v46_ = _dafny.Map({self.f4: d_783_cf0_})
                d_791_v47_: D6
                d_791_v47_ = D6_DC18(((d_790_v46_)[self.f4] if (self.f4) in (d_790_v46_) else d_783_cf0_), (self).f7)
                d_792_v48_: _dafny.Map
                d_792_v48_ = _dafny.Map({(d_789_v45_).f14: default__.fm2((d_791_v47_).cf25, globalState)})
                d_793_v49_: _dafny.Seq
                d_793_v49_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h")))])
                rhs147_ = (d_792_v48_ if (d_793_v49_) != (d_793_v49_) else d_792_v48_)
                rhs148_ = (d_789_v45_).f14
                rhs149_ = default__.safeDivisionInt((d_789_v45_).f14, (self).f7)
                rhs150_ = 128
                lhs138_ = globalState
                lhs139_ = globalState
                lhs140_ = globalState
                d_792_v48_ = rhs147_
                lhs138_.f1 = rhs148_
                lhs139_.f1 = rhs149_
                lhs140_.f1 = rhs150_
                (self).f13 = (self.f13) + (_dafny.SeqWithoutIsStrInference([(self).f12 for d_794_i4_ in range(default__.abs(659))]))
            elif True:
                d_795_v50_: _dafny.Seq
                d_795_v50_ = _dafny.SeqWithoutIsStrInference([not(d_783_cf0_), d_783_cf0_])
                (globalState).f2 = (d_795_v50_)[default__.safeIndex(len(self.f13), len(d_795_v50_))]
                d_796_v51_: _dafny.Set
                d_796_v51_ = _dafny.Set({True})
                d_797_v52_: D9
                d_797_v52_ = D9_DC25(d_796_v51_)
                rhs151_ = ((d_797_v52_).cf37) - (d_796_v51_)
                rhs152_ = ((self).fm13(globalState)) or (self.f4)
                lhs141_ = globalState
                d_796_v51_ = rhs151_
                lhs141_.f2 = rhs152_
                index145_ = default__.safeIndex(555, (d_784_v40_).length(0))
                (d_784_v40_)[index145_] = not(d_783_cf0_)
                d_798_v53_: _dafny.MultiSet
                d_798_v53_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gfobxgt")))])
                d_799_v54_: _dafny.MultiSet
                d_799_v54_ = _dafny.MultiSet([default__.fm15(False, (0) - ((self).f3), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vbacq"))), globalState), self.f13])
                d_800_v55_: _dafny.Set
                d_800_v55_ = _dafny.Set({(self).f12, (self).f12, (self).f12})
                index146_ = default__.safeIndex(555, (d_784_v40_).length(0))
                (d_784_v40_)[index146_] = not (not((d_798_v53_).isdisjoint(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(d_799_v54_).cardinality, len(d_800_v55_)]))))) or (self.f4)
                (globalState).f1 = (self).f3
                (globalState).f2 = self.f4
            index147_ = default__.safeIndex(283, (d_784_v40_).length(0))
            (d_784_v40_)[index147_] = (default__.fm0(globalState)) > (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mibpnhp"))))
            d_801_v56_: _dafny.MultiSet
            d_801_v56_ = _dafny.MultiSet([(self).f3, (self).f3, (self).f7, (self).f7])
            index148_ = default__.safeIndex(283, (d_784_v40_).length(0))
            (d_784_v40_)[index148_] = (d_801_v56_).isdisjoint(d_801_v56_)
            index149_ = default__.safeIndex(216, (d_732_v6_).length(0))
            (d_732_v6_)[index149_] = default__.safeModuloInt((self).f3, (self).f7)
            d_802_v57_: _dafny.Map
            d_802_v57_ = _dafny.Map({(self).f7: d_801_v56_})
            index150_ = default__.safeIndex(216, (d_732_v6_).length(0))
            (d_732_v6_)[index150_] = len((d_802_v57_) | (d_802_v57_))
        elif True:
            d_803___mcc_h7_ = source10_.cf7
            d_804_cf7_ = d_803___mcc_h7_
            d_805_v58_: _dafny.Map
            d_805_v58_ = _dafny.Map({(self).f3: not(((self).f12) in (self.f13))})
            d_806_v59_: D9
            d_806_v59_ = D9_DC26((self).f3, True)
            d_805_v58_ = (d_805_v58_).set((self).f7, (d_806_v59_).cf39)
            (globalState).f2 = self.f4
            d_807_v60_: D8
            d_807_v60_ = D8_DC23(self.f4, False, True, (self).f12)
            d_808_v61_: C2
            nw151_ = C2()
            nw151_.ctor__((self).f3, d_732_v6_, ((self).f3) - ((self).f7), (d_807_v60_).cf33)
            d_808_v61_ = nw151_
            d_809_v62_: _dafny.Set
            d_809_v62_ = _dafny.Set({False, self.f4, self.f4})
            d_810_v63_: _dafny.Array
            nw152_ = _dafny.Array(None, 22)
            nw152_[int(0)] = self.f4
            nw152_[int(1)] = self.f4
            nw152_[int(2)] = ((self).f3) not in (_dafny.SeqWithoutIsStrInference([(d_808_v61_).f14]))
            nw152_[int(3)] = self.f4
            nw152_[int(4)] = not(self.f4)
            nw152_[int(5)] = (self.f4) and (self.f4)
            nw152_[int(6)] = ((d_808_v61_).f14) > (759)
            nw152_[int(7)] = (self.f4 if False else self.f4)
            nw152_[int(8)] = self.f4
            nw152_[int(9)] = ((self).f7) != ((self).f3)
            nw152_[int(10)] = self.f4
            nw152_[int(11)] = self.f4
            nw152_[int(12)] = self.f4
            nw152_[int(13)] = not(False)
            nw152_[int(14)] = self.f4
            nw152_[int(15)] = self.f4
            nw152_[int(16)] = self.f4
            nw152_[int(17)] = (_dafny.Set({self.f4})).isdisjoint(d_809_v62_)
            nw152_[int(18)] = self.f4
            nw152_[int(19)] = self.f4
            nw152_[int(20)] = self.f4
            nw152_[int(21)] = (self).fm13(globalState)
            d_810_v63_ = nw152_
            index151_ = default__.safeIndex(513, (d_810_v63_).length(0))
            (d_810_v63_)[index151_] = self.f4
            index152_ = default__.safeIndex(513, (d_810_v63_).length(0))
            (d_810_v63_)[index152_] = (not((not(True)) and (self.f4)) if (False if self.f4 else not(self.f4)) else self.f4)
        d_811_v64_: str
        d_811_v64_ = _dafny.CodePoint('g')
        d_812_v65_: D8
        d_812_v65_ = D8_DC23(self.f4, self.f4, self.f4, d_811_v64_)
        d_811_v64_ = (d_812_v65_).cf35
        hi7_ = (0) - ((self).f3)
        for d_813_i5_ in range((self).f7, hi7_):
            d_814_v66_: _dafny.Array
            nw153_ = _dafny.Array(D6.default()(), 21)
            d_814_v66_ = nw153_
            d_815_v67_: _dafny.Map
            d_815_v67_ = _dafny.Map({self.f4: self.f4})
            d_816_v68_: D6
            d_816_v68_ = D6_DC19(len(d_815_v67_))
            index153_ = default__.safeIndex(13, (d_814_v66_).length(0))
            (d_814_v66_)[index153_] = d_816_v68_
            index154_ = default__.safeIndex(13, (d_814_v66_).length(0))
            def iife42_(_pat_let15_0):
                def iife43_(d_817_dt__update__tmp_h0_):
                    def iife44_(_pat_let16_0):
                        def iife45_(d_818_dt__update_hcf26_h0_):
                            return D6_DC19(d_818_dt__update_hcf26_h0_)
                        return iife45_(_pat_let16_0)
                    return iife44_(d_813_i5_)
                return iife43_(_pat_let15_0)
            (d_814_v66_)[index154_] = iife42_(d_816_v68_)
            d_819_v69_: _dafny.Map
            d_819_v69_ = _dafny.Map({default__.fm28(not(self.f4), (self).f3, self.f4, globalState): (self).f3})
            d_820_v70_: _dafny.Seq
            d_820_v70_ = _dafny.SeqWithoutIsStrInference([(0) - ((self).f7)])
            d_821_v71_: _dafny.Map
            d_821_v71_ = _dafny.Map({d_820_v70_: _dafny.SeqWithoutIsStrInference([d_811_v64_ for d_822_i6_ in range(default__.abs(-752))])})
            d_819_v69_ = (d_819_v69_).set((D10_DC28(d_821_v71_)).cf41, d_813_i5_)
            (globalState).f2 = not ((d_813_i5_) >= ((self).f7)) or (self.f4)
            d_823_v72_: _dafny.Array
            def lambda43_(d_824_i7_):
                return self.f4

            init27_ = lambda43_
            nw154_ = _dafny.Array(None, 3)
            for i0_27_ in range(nw154_.length(0)):
                nw154_[i0_27_] = init27_(i0_27_)
            d_823_v72_ = nw154_
            index155_ = default__.safeIndex(363, (d_823_v72_).length(0))
            (d_823_v72_)[index155_] = self.f4
            index156_ = default__.safeIndex(363, (d_823_v72_).length(0))
            (d_823_v72_)[index156_] = self.f4
        d_825_v73_: D0
        d_825_v73_ = D0_DC3()
        source11_ = d_825_v73_
        if source11_.is_DC1:
            d_826___mcc_h8_ = source11_.cf1
            d_827___mcc_h9_ = source11_.cf2
            d_828___mcc_h10_ = source11_.cf3
            d_829___mcc_h11_ = source11_.cf4
            d_830_cf4_ = d_829___mcc_h11_
            d_831_cf3_ = d_828___mcc_h10_
            d_832_cf2_ = d_827___mcc_h9_
            d_833_cf1_ = d_826___mcc_h8_
            (globalState).f1 = ((self).f7 if False else 338)
            (self).f4 = d_833_cf1_
            d_834_v74_: _dafny.Map
            d_834_v74_ = _dafny.Map({d_831_cf3_: D9_DC26((self).f7, False)})
            d_835_v75_: _dafny.Set
            d_835_v75_ = _dafny.Set({(self).f7, len(d_834_v74_)})
            (self).f13 = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ymicw"))) + (self.f13)).set(default__.safeIndex(len(d_835_v75_), len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ymicw"))) + (self.f13))), d_811_v64_)
            index157_ = default__.safeIndex(829, (d_732_v6_).length(0))
            (d_732_v6_)[index157_] = ((self).f3 if d_832_cf2_ else 476)
            d_836_v76_: _dafny.Array
            def lambda44_(d_837_cf1_):
                def lambda45_(d_838_i8_):
                    return d_837_cf1_

                return lambda45_

            init28_ = lambda44_(d_833_cf1_)
            nw155_ = _dafny.Array(None, 1)
            for i0_28_ in range(nw155_.length(0)):
                nw155_[i0_28_] = init28_(i0_28_)
            d_836_v76_ = nw155_
            index158_ = default__.safeIndex(613, (d_836_v76_).length(0))
            (d_836_v76_)[index158_] = self.f4
            d_839_v77_: _dafny.Seq
            d_839_v77_ = _dafny.SeqWithoutIsStrInference([not(True)])
            index159_ = default__.safeIndex(829, (d_732_v6_).length(0))
            index160_ = default__.safeIndex(613, (d_836_v76_).length(0))
            rhs153_ = not ((d_839_v77_)[default__.safeIndex((0) - ((self).f3), len(d_839_v77_))]) or (d_832_cf2_)
            rhs154_ = 52
            rhs155_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dcb"))) + (self.f13)) == (self.f13)
            rhs156_ = not((d_839_v77_)[default__.safeIndex(default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "athho"))), (self).f3), len(d_839_v77_))])
            rhs157_ = d_832_cf2_
            lhs142_ = self
            lhs143_ = d_732_v6_
            lhs144_ = default__.safeIndex(829, (d_732_v6_).length(0))
            lhs145_ = self
            lhs146_ = d_836_v76_
            lhs147_ = default__.safeIndex(613, (d_836_v76_).length(0))
            lhs148_ = globalState
            lhs142_.f4 = rhs153_
            lhs143_[lhs144_] = rhs154_
            lhs145_.f4 = rhs155_
            lhs146_[lhs147_] = rhs156_
            lhs148_.f2 = rhs157_
        elif source11_.is_DC2:
            d_840___mcc_h12_ = source11_.cf5
            d_841___mcc_h13_ = source11_.cf6
            d_842_cf6_ = d_841___mcc_h13_
            d_843_cf5_ = d_840___mcc_h12_
            (globalState).f1 = default__.safeDivisionInt((self).f3, (self).f3)
            (self).f4 = (len((self.f13) + (self.f13))) != ((self).f3)
            d_844_v78_: _dafny.Array
            nw156_ = _dafny.Array(False, 23)
            d_844_v78_ = nw156_
            d_845_v79_: C0
            nw157_ = C0()
            nw157_.ctor__(not(d_842_cf6_), d_844_v78_)
            d_845_v79_ = nw157_
            index161_ = default__.safeIndex(624, (d_844_v78_).length(0))
            (d_844_v78_)[index161_] = d_842_cf6_
            index162_ = default__.safeIndex(624, (d_844_v78_).length(0))
            (d_844_v78_)[index162_] = (_dafny.SeqWithoutIsStrInference([d_811_v64_ for d_846_i9_ in range(default__.abs(728))])) == (self.f13)
        elif source11_.is_DC3:
            (globalState).f1 = (0) - ((self).f3)
            d_847_v80_: _dafny.MultiSet
            d_847_v80_ = _dafny.MultiSet([(self).f7, (self).f7, (self).f3, 551])
            d_848_v81_: _dafny.MultiSet
            d_848_v81_ = _dafny.MultiSet([(d_847_v80_).cardinality])
            (globalState).f2 = (d_848_v81_).issubset(d_848_v81_)
            d_849_v82_: _dafny.Seq
            d_849_v82_ = _dafny.SeqWithoutIsStrInference([self.f4, False, self.f4, self.f4])
            d_850_v83_: _dafny.MultiSet
            d_850_v83_ = _dafny.MultiSet([(d_849_v82_)[default__.safeIndex((self).f3, len(d_849_v82_))], self.f4])
            (globalState).f1 = default__.safeModuloInt((0) - (((d_850_v83_) - (d_850_v83_)).cardinality), (self).f3)
            (globalState).f1 = (279 if (561) == ((self).f7) else (self).f7)
        elif source11_.is_DC0:
            d_851___mcc_h14_ = source11_.cf0
            d_852_cf0_ = d_851___mcc_h14_
            rhs158_ = d_852_cf0_
            rhs159_ = default__.fm2((self).f3, globalState)
            lhs149_ = globalState
            lhs150_ = self
            lhs149_.f2 = rhs158_
            lhs150_.f4 = rhs159_
            pat_let_tv10_ = d_811_v64_
            pat_let_tv11_ = d_852_cf0_
            d_853_v84_: _dafny.Array
            nw158_ = _dafny.Array(None, 13)
            nw158_[int(0)] = d_812_v65_
            nw158_[int(1)] = D8_DC23(self.f4, self.f4, True, (self).f12)
            nw158_[int(2)] = d_812_v65_
            nw158_[int(3)] = D8_DC23(self.f4, True, self.f4, d_811_v64_)
            nw158_[int(4)] = d_812_v65_
            nw158_[int(5)] = d_812_v65_
            nw158_[int(6)] = D8_DC23(self.f4, d_852_cf0_, self.f4, d_811_v64_)
            nw158_[int(7)] = d_812_v65_
            nw158_[int(8)] = D8_DC23(d_852_cf0_, d_852_cf0_, self.f4, (self).f12)
            nw158_[int(9)] = d_812_v65_
            nw158_[int(10)] = d_812_v65_
            nw158_[int(11)] = d_812_v65_
            def iife46_(_pat_let17_0):
                def iife47_(d_854_dt__update__tmp_h1_):
                    def iife48_(_pat_let18_0):
                        def iife49_(d_855_dt__update_hcf35_h0_):
                            def iife50_(_pat_let19_0):
                                def iife51_(d_856_dt__update_hcf32_h0_):
                                    return D8_DC23(d_856_dt__update_hcf32_h0_, (d_854_dt__update__tmp_h1_).cf33, (d_854_dt__update__tmp_h1_).cf34, d_855_dt__update_hcf35_h0_)
                                return iife51_(_pat_let19_0)
                            return iife50_(pat_let_tv11_)
                        return iife49_(_pat_let18_0)
                    return iife48_(pat_let_tv10_)
                return iife47_(_pat_let17_0)
            nw158_[int(12)] = iife46_(d_812_v65_)
            d_853_v84_ = nw158_
            d_857_v85_: _dafny.Map
            d_857_v85_ = _dafny.Map({(self).f7: d_853_v84_})
            d_857_v85_ = _dafny.Map({638: d_853_v84_})
            (self).f4 = True
            d_858_v86_: _dafny.MultiSet
            d_858_v86_ = _dafny.MultiSet([self.f4, self.f4, d_852_cf0_, False, self.f4])
            d_859_v87_: _dafny.MultiSet
            d_859_v87_ = _dafny.MultiSet([len(self.f13), (self).f7, default__.safeModuloInt((self).f3, (d_858_v86_).cardinality), (0) - ((self).f7)])
            d_860_v88_: _dafny.Seq
            d_860_v88_ = _dafny.SeqWithoutIsStrInference([(self).f3, (self).f3, (self).f3])
            d_861_v89_: D11
            d_861_v89_ = D11_DC31(d_860_v88_)
            d_859_v87_ = _dafny.MultiSet((d_861_v89_).cf49)
        elif True:
            d_862___mcc_h15_ = source11_.cf7
            d_863_cf7_ = d_862___mcc_h15_
            d_864_v91_: _dafny.Set
            d_864_v91_ = _dafny.Set({(self).f3})
            def iife52_():
                coll12_ = _dafny.Set()
                compr_12_: int
                for compr_12_ in _dafny.IntegerRange(819, 503):
                    d_865_v90_: int = compr_12_
                    if ((819) <= (d_865_v90_)) and ((d_865_v90_) < (503)):
                        coll12_ = coll12_.union(_dafny.Set([default__.safeModuloInt(d_865_v90_, (self).f3)]))
                return _dafny.Set(coll12_)
            (globalState).f2 = (d_864_v91_).issubset(iife52_()
            )
            d_866_v92_: _dafny.Map
            d_866_v92_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "upfnlrddl"))): self.f4})
            d_867_v94_: _dafny.Map
            d_867_v94_ = _dafny.Map({_dafny.CodePoint('p'): len(d_864_v91_)})
            d_868_v97_: _dafny.Seq
            d_868_v97_ = _dafny.SeqWithoutIsStrInference([44, (self).f3])
            d_869_v98_: _dafny.Seq
            d_869_v98_ = _dafny.SeqWithoutIsStrInference([self.f4])
            d_870_v100_: _dafny.Array
            nw159_ = _dafny.Array(None, 28)
            nw159_[int(0)] = (d_866_v92_) | (d_866_v92_)
            nw159_[int(1)] = (d_866_v92_).set(default__.fm0(globalState), self.f4)
            nw159_[int(2)] = d_866_v92_
            nw159_[int(3)] = d_866_v92_
            nw159_[int(4)] = ((d_866_v92_).set((self).f3, self.f4)) | (d_866_v92_)
            nw159_[int(5)] = _dafny.Map({(self).f7: self.f4})
            nw159_[int(6)] = d_866_v92_
            nw159_[int(7)] = (d_866_v92_).set((self).f7, self.f4)
            nw159_[int(8)] = (d_866_v92_) | (_dafny.Map({(self).f3: self.f4}))
            nw159_[int(9)] = d_866_v92_
            nw159_[int(10)] = d_866_v92_
            nw159_[int(11)] = d_866_v92_
            def iife53_():
                coll13_ = _dafny.Map()
                compr_13_: int
                for compr_13_ in _dafny.IntegerRange(-833, -986):
                    d_871_v93_: int = compr_13_
                    if ((-833) <= (d_871_v93_)) and ((d_871_v93_) < (-986)):
                        coll13_[(d_871_v93_) + ((self).f7)] = self.f4
                return _dafny.Map(coll13_)
            nw159_[int(12)] = (d_866_v92_) | (iife53_()
            )
            nw159_[int(13)] = d_866_v92_
            nw159_[int(14)] = _dafny.Map({((d_867_v94_)[d_811_v64_] if (d_811_v64_) in (d_867_v94_) else (self).f7): self.f4})
            nw159_[int(15)] = d_866_v92_
            def iife54_():
                coll14_ = _dafny.Map()
                compr_14_: int
                for compr_14_ in _dafny.IntegerRange(626, 259):
                    d_872_v95_: int = compr_14_
                    if ((626) <= (d_872_v95_)) and ((d_872_v95_) < (259)):
                        coll14_[(d_872_v95_) - ((self).f7)] = self.f4
                return _dafny.Map(coll14_)
            nw159_[int(16)] = (iife54_()
            ) | (_dafny.Map({607: self.f4}))
            nw159_[int(17)] = d_866_v92_
            nw159_[int(18)] = d_866_v92_
            nw159_[int(19)] = d_866_v92_
            nw159_[int(20)] = d_866_v92_
            def iife55_():
                coll15_ = _dafny.Map()
                compr_15_: int
                for compr_15_ in (d_868_v97_).Elements:
                    d_873_v96_: int = compr_15_
                    if (d_873_v96_) in (d_868_v97_):
                        coll15_[default__.safeModuloInt(d_873_v96_, (self).f3)] = self.f4
                return _dafny.Map(coll15_)
            nw159_[int(21)] = iife55_()
            
            nw159_[int(22)] = (d_866_v92_) | (d_866_v92_)
            nw159_[int(23)] = (d_866_v92_) | (d_866_v92_)
            nw159_[int(24)] = _dafny.Map({(self).f3: not(self.f4)})
            nw159_[int(25)] = ((_dafny.Map({(self).f7: self.f4})).set((self).f3, False)) | (d_866_v92_)
            nw159_[int(26)] = _dafny.Map({211: (d_869_v98_)[default__.safeIndex((self).f7, len(d_869_v98_))]})
            def iife56_():
                coll16_ = _dafny.Map()
                compr_16_: int
                for compr_16_ in _dafny.IntegerRange(937, 258):
                    d_874_v99_: int = compr_16_
                    if ((937) <= (d_874_v99_)) and ((d_874_v99_) < (258)):
                        coll16_[(d_874_v99_) - ((self).f7)] = self.f4
                return _dafny.Map(coll16_)
            nw159_[int(27)] = (d_866_v92_) | (iife56_()
            )
            d_870_v100_ = nw159_
            index163_ = default__.safeIndex(496, (d_870_v100_).length(0))
            (d_870_v100_)[index163_] = d_866_v92_
            d_875_v101_: _dafny.MultiSet
            d_875_v101_ = _dafny.MultiSet([self.f4, self.f4])
            index164_ = default__.safeIndex(496, (d_870_v100_).length(0))
            rhs160_ = _dafny.Map({(self).f7: (len(d_868_v97_)) > ((self).f7)})
            rhs161_ = (d_875_v101_).ispropersubset(d_875_v101_)
            lhs151_ = d_870_v100_
            lhs152_ = default__.safeIndex(496, (d_870_v100_).length(0))
            lhs153_ = globalState
            lhs151_[lhs152_] = rhs160_
            lhs153_.f2 = rhs161_
            (globalState).f2 = self.f4
            d_876_v102_: C1
            nw160_ = C1()
            nw160_.ctor__((self).f7, 885, default__.fm2((self).f3, globalState))
            d_876_v102_ = nw160_
        (self).f4 = True

    @property
    def f12(self):
        return self._f12

class C4(T1, T0):
    def  __init__(self):
        self._f4: bool = False
        self._f7: int = int(0)
        self._f3: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    @property
    def f4(self):
        return self._f4
    @f4.setter
    def f4(self, value):
        self._f4 = value
    @property
    def f7(self):
        return self._f7
    @property
    def f3(self):
        return self._f3
    def ctor__(self, f7, f3, f4):
        (self)._f7 = f7
        (self)._f3 = f3
        (self).f4 = f4

    def m1(self, p0, globalState):
        r0: int = int(0)
        (globalState).f2 = self.f4
        d_877_v0_: _dafny.Array
        def lambda46_(d_878_i0_):
            return (_dafny.SeqWithoutIsStrInference([not(self.f4)])) == (_dafny.SeqWithoutIsStrInference([not(self.f4), self.f4]))

        init29_ = lambda46_
        nw161_ = _dafny.Array(None, 7)
        for i0_29_ in range(nw161_.length(0)):
            nw161_[i0_29_] = init29_(i0_29_)
        d_877_v0_ = nw161_
        index165_ = default__.safeIndex(80, (d_877_v0_).length(0))
        (d_877_v0_)[index165_] = self.f4
        index166_ = default__.safeIndex(80, (d_877_v0_).length(0))
        (d_877_v0_)[index166_] = True
        hi8_ = 972
        for d_879_i1_ in range((self).f7, hi8_):
            d_880_v1_: _dafny.Array
            nw162_ = _dafny.Array(int(0), 19)
            d_880_v1_ = nw162_
            d_881_v2_: _dafny.Set
            d_881_v2_ = _dafny.Set({d_880_v1_})
            d_882_v3_: _dafny.Array
            nw163_ = _dafny.Array(_dafny.Seq({}), 1)
            d_882_v3_ = nw163_
            d_883_v4_: _dafny.Seq
            d_883_v4_ = _dafny.SeqWithoutIsStrInference([(self).f3, p0, (self).f3, d_879_i1_])
            index167_ = default__.safeIndex(549, (d_882_v3_).length(0))
            (d_882_v3_)[index167_] = (d_883_v4_) + (d_883_v4_)
            d_884_v5_: _dafny.Map
            d_884_v5_ = _dafny.Map({(0) - (default__.fm0(globalState)): d_881_v2_})
            index168_ = default__.safeIndex(549, (d_882_v3_).length(0))
            rhs162_ = self.f4
            rhs163_ = ((d_884_v5_)[p0] if (p0) in (d_884_v5_) else d_881_v2_)
            rhs164_ = d_883_v4_
            lhs154_ = globalState
            lhs155_ = d_882_v3_
            lhs156_ = default__.safeIndex(549, (d_882_v3_).length(0))
            lhs154_.f2 = rhs162_
            d_881_v2_ = rhs163_
            lhs155_[lhs156_] = rhs164_
            d_885_v6_: _dafny.Seq
            d_885_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g"))
            (globalState).f1 = (0) - (len((d_885_v6_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_886_i2_ in range(default__.abs(-923))]))))
            d_887_v7_: C0
            nw164_ = C0()
            nw164_.ctor__((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dmaqyb"))) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_888_i3_ in range(default__.abs(486))])), d_877_v0_)
            d_887_v7_ = nw164_
            d_889_v8_: _dafny.MultiSet
            d_889_v8_ = _dafny.MultiSet([d_879_i1_])
            d_890_v9_: str
            d_890_v9_ = _dafny.CodePoint('a')
            d_891_v10_: _dafny.MultiSet
            d_891_v10_ = _dafny.MultiSet([d_890_v9_])
            d_892_v11_: D0
            d_892_v11_ = D0_DC2(d_891_v10_, d_887_v7_.f10)
            d_893_v12_: _dafny.Map
            d_893_v12_ = _dafny.Map({d_892_v11_: (d_877_v0_)[default__.safeIndex(80, (d_877_v0_).length(0))]})
            pat_let_tv12_ = d_891_v10_
            pat_let_tv13_ = d_891_v10_
            def iife57_(_pat_let20_0):
                def iife58_(d_894_dt__update__tmp_h1_):
                    def iife59_(_pat_let21_0):
                        def iife60_(d_895_dt__update_hcf5_h1_):
                            return D0_DC2(d_895_dt__update_hcf5_h1_, (d_894_dt__update__tmp_h1_).cf6)
                        return iife60_(_pat_let21_0)
                    return iife59_(pat_let_tv12_)
                return iife58_(_pat_let20_0)
            def iife61_(_pat_let22_0):
                def iife62_(d_896_dt__update__tmp_h0_):
                    def iife63_(_pat_let23_0):
                        def iife64_(d_897_dt__update_hcf5_h0_):
                            return D0_DC2(d_897_dt__update_hcf5_h0_, (d_896_dt__update__tmp_h0_).cf6)
                        return iife64_(_pat_let23_0)
                    return iife63_(pat_let_tv13_)
                return iife62_(_pat_let22_0)
            if (d_889_v8_).isdisjoint(default__.fm6(((d_893_v12_)[iife57_(D0_DC2(d_891_v10_, True))] if (iife61_(D0_DC2(d_891_v10_, True))) in (d_893_v12_) else self.f4), (d_877_v0_)[default__.safeIndex(80, (d_877_v0_).length(0))], globalState)):
                d_898_v13_: D0
                d_898_v13_ = D0_DC1(not(self.f4), d_887_v7_.f10, d_890_v9_, d_880_v1_)
                d_899_v14_: _dafny.Seq
                d_899_v14_ = _dafny.SeqWithoutIsStrInference([(d_898_v13_).cf1])
                d_900_v15_: _dafny.Map
                d_900_v15_ = _dafny.Map({d_899_v14_: d_885_v6_})
                d_901_v16_: _dafny.Set
                d_901_v16_ = _dafny.Set({d_879_i1_})
                d_902_v18_: _dafny.Map
                d_902_v18_ = _dafny.Map({d_885_v6_: (self).f7})
                def iife65_():
                    coll17_ = _dafny.Map()
                    compr_17_: int
                    for compr_17_ in _dafny.IntegerRange(792, -411):
                        d_903_v17_: int = compr_17_
                        if ((792) <= (d_903_v17_)) and ((d_903_v17_) < (-411)):
                            coll17_[(d_903_v17_) * (((d_889_v8_)[(self).f3] if ((self).f3) in (d_889_v8_) else (self).f7))] = len(_dafny.Map({d_879_i1_: d_887_v7_.f10}))
                    return _dafny.Map(coll17_)
                rhs165_ = (d_890_v9_ if not (self.f4) or (d_887_v7_.f10) else d_890_v9_)
                rhs166_ = (d_900_v15_) | ((d_900_v15_) | (d_900_v15_))
                rhs167_ = d_901_v16_
                rhs168_ = len(default__.fm7(iife65_()
                , (self).f3, d_902_v18_, globalState))
                lhs157_ = globalState
                d_890_v9_ = rhs165_
                d_900_v15_ = rhs166_
                d_901_v16_ = rhs167_
                lhs157_.f1 = rhs168_
                (globalState).f1 = default__.safeModuloInt((self).f7, p0)
                d_904_v19_: _dafny.Map
                d_904_v19_ = _dafny.Map({d_887_v7_.f10: (self).f7})
                index169_ = default__.safeIndex(696, (d_880_v1_).length(0))
                (d_880_v1_)[index169_] = len(d_904_v19_)
                d_905_v20_: D1
                d_905_v20_ = D1_DC7()
                d_906_v21_: _dafny.Map
                d_906_v21_ = _dafny.Map({D1_DC7(): self.f4})
                d_907_v22_: _dafny.MultiSet
                d_907_v22_ = _dafny.MultiSet([self.f4, d_887_v7_.f10, d_887_v7_.f10])
                index170_ = default__.safeIndex(696, (d_880_v1_).length(0))
                rhs169_ = not(d_887_v7_.f10)
                rhs170_ = (self).f7
                rhs171_ = (d_905_v20_) in (d_906_v21_)
                rhs172_ = ((len(d_885_v6_)) * ((d_907_v22_).cardinality)) - ((self).f7)
                lhs158_ = globalState
                lhs159_ = globalState
                lhs160_ = globalState
                lhs161_ = d_880_v1_
                lhs162_ = default__.safeIndex(696, (d_880_v1_).length(0))
                lhs158_.f2 = rhs169_
                lhs159_.f1 = rhs170_
                lhs160_.f2 = rhs171_
                lhs161_[lhs162_] = rhs172_
                d_908_v23_: C0
                nw165_ = C0()
                nw165_.ctor__((self.f4) or (self.f4), (d_887_v7_).f11)
                d_908_v23_ = nw165_
                d_909_v24_: D0
                d_909_v24_ = D0_DC3()
                d_909_v24_ = d_909_v24_
            elif True:
                (globalState).f2 = (d_877_v0_)[default__.safeIndex(80, (d_877_v0_).length(0))]
                (globalState).f1 = (self).f7
                d_910_v25_: _dafny.Map
                d_910_v25_ = _dafny.Map({d_887_v7_.f10: not(self.f4)})
                index171_ = default__.safeIndex(950, (d_880_v1_).length(0))
                (d_880_v1_)[index171_] = (d_879_i1_) * ((self).f7)
                index172_ = default__.safeIndex(950, (d_880_v1_).length(0))
                def iife66_():
                    coll18_ = _dafny.Set()
                    compr_18_: int
                    for compr_18_ in _dafny.IntegerRange(149, 861):
                        d_911_v26_: int = compr_18_
                        if ((149) <= (d_911_v26_)) and ((d_911_v26_) < (861)):
                            coll18_ = coll18_.union(_dafny.Set([default__.safeModuloInt(d_911_v26_, d_879_i1_)]))
                    return _dafny.Set(coll18_)
                rhs173_ = d_887_v7_.f10
                rhs174_ = d_910_v25_
                rhs175_ = len(iife66_()
                )
                rhs176_ = (0) - ((default__.safeDivisionInt(p0, (self).f3)) * (default__.fm0(globalState)))
                lhs163_ = self
                lhs164_ = globalState
                lhs165_ = d_880_v1_
                lhs166_ = default__.safeIndex(950, (d_880_v1_).length(0))
                lhs163_.f4 = rhs173_
                d_910_v25_ = rhs174_
                lhs164_.f1 = rhs175_
                lhs165_[lhs166_] = rhs176_
                d_912_v27_: D0
                d_912_v27_ = D0_DC0(d_887_v7_.f10)
                d_885_v6_ = default__.fm8(False, D1_DC7(), d_912_v27_, p0, globalState)
                (globalState).f2 = not(True)
        d_913_v28_: _dafny.Seq
        d_913_v28_ = _dafny.SeqWithoutIsStrInference([231, (0) - (len(_dafny.Set({(d_877_v0_)[default__.safeIndex(80, (d_877_v0_).length(0))], self.f4})))])
        r0 = (d_913_v28_)[default__.safeIndex((p0) - ((self).f3), len(d_913_v28_))]
        d_914_v29_: _dafny.Array
        d_914_v29_ = d_877_v0_
        d_915_v30_: _dafny.Seq
        d_915_v30_ = _dafny.SeqWithoutIsStrInference([(d_914_v29_)])
        d_916_v31_: _dafny.Array
        nw166_ = _dafny.Array(None, 19)
        nw166_[int(0)] = d_877_v0_
        nw166_[int(1)] = (d_877_v0_ if (d_877_v0_)[default__.safeIndex(80, (d_877_v0_).length(0))] else d_877_v0_)
        nw166_[int(2)] = d_877_v0_
        nw166_[int(3)] = d_877_v0_
        nw166_[int(4)] = d_877_v0_
        nw166_[int(5)] = d_877_v0_
        nw166_[int(6)] = d_877_v0_
        nw166_[int(7)] = d_877_v0_
        nw166_[int(8)] = d_877_v0_
        nw166_[int(9)] = d_877_v0_
        nw166_[int(10)] = d_877_v0_
        nw166_[int(11)] = d_877_v0_
        nw166_[int(12)] = d_877_v0_
        nw166_[int(13)] = d_877_v0_
        nw166_[int(14)] = d_877_v0_
        nw166_[int(15)] = d_877_v0_
        nw166_[int(16)] = (d_915_v30_)[default__.safeIndex((0) - (p0), len(d_915_v30_))]
        nw166_[int(17)] = (d_877_v0_ if (d_877_v0_)[default__.safeIndex(80, (d_877_v0_).length(0))] else d_877_v0_)
        nw166_[int(18)] = d_877_v0_
        d_916_v31_ = nw166_
        index173_ = default__.safeIndex(203, (d_916_v31_).length(0))
        (d_916_v31_)[index173_] = d_877_v0_
        index174_ = default__.safeIndex(203, (d_916_v31_).length(0))
        (d_916_v31_)[index174_] = d_877_v0_
        d_917_v32_: C0
        nw167_ = C0()
        nw167_.ctor__((d_877_v0_)[default__.safeIndex(80, (d_877_v0_).length(0))], (d_915_v30_)[default__.safeIndex(919, len(d_915_v30_))])
        d_917_v32_ = nw167_
        d_918_v33_: D0
        d_918_v33_ = D0_DC0(d_917_v32_.f10)
        pat_let_tv14_ = p0
        pat_let_tv15_ = d_877_v0_
        pat_let_tv16_ = d_877_v0_
        def lambda47_(source12_):
            if source12_.is_DC1:
                d_919___mcc_h0_ = source12_.cf1
                d_920___mcc_h1_ = source12_.cf2
                d_921___mcc_h2_ = source12_.cf3
                d_922___mcc_h3_ = source12_.cf4
                d_923_cf4_ = d_922___mcc_h3_
                d_924_cf3_ = d_921___mcc_h2_
                d_925_cf2_ = d_920___mcc_h1_
                d_926_cf1_ = d_919___mcc_h0_
                return (pat_let_tv14_) + (len(_dafny.Map({(pat_let_tv16_)[default__.safeIndex(80, (pat_let_tv15_).length(0))]: True})))
            elif source12_.is_DC2:
                d_927___mcc_h4_ = source12_.cf5
                d_928___mcc_h5_ = source12_.cf6
                d_929_cf6_ = d_928___mcc_h5_
                d_930_cf5_ = d_927___mcc_h4_
                return (self).f7
            elif source12_.is_DC3:
                return 480
            elif source12_.is_DC0:
                d_931___mcc_h6_ = source12_.cf0
                d_932_cf0_ = d_931___mcc_h6_
                return (self).f7
            elif True:
                d_933___mcc_h7_ = source12_.cf7
                d_934_cf7_ = d_933___mcc_h7_
                return (self).f7

        r0 = (0) - (lambda47_(d_918_v33_))
        return r0

    def m4(self, p0, globalState):
        d_935_i0_: int
        d_935_i0_ = 0
        with _dafny.label("8"):
            while self.f4:
                with _dafny.c_label("8"):
                    if (d_935_i0_) >= (100):
                        raise _dafny.Break("8")
                    d_935_i0_ = (d_935_i0_) + (1)
                    d_936_v0_: str
                    d_936_v0_ = _dafny.CodePoint('w')
                    d_937_v1_: _dafny.Map
                    d_937_v1_ = _dafny.Map({not(self.f4): d_936_v0_})
                    d_938_v2_: _dafny.Array
                    nw168_ = _dafny.Array(None, 6)
                    nw168_[int(0)] = d_936_v0_
                    nw168_[int(1)] = d_936_v0_
                    nw168_[int(2)] = d_936_v0_
                    nw168_[int(3)] = ((d_937_v1_)[default__.fm2(866, globalState)] if (default__.fm2(866, globalState)) in (d_937_v1_) else d_936_v0_)
                    nw168_[int(4)] = d_936_v0_
                    nw168_[int(5)] = d_936_v0_
                    d_938_v2_ = nw168_
                    d_939_v3_: _dafny.Seq
                    d_939_v3_ = _dafny.SeqWithoutIsStrInference([d_938_v2_, d_938_v2_])
                    d_940_v4_: _dafny.Map
                    d_940_v4_ = _dafny.Map({self.f4: (self).f7})
                    d_941_v5_: _dafny.MultiSet
                    d_941_v5_ = _dafny.MultiSet([default__.fm0(globalState), (self).f3, (self).f3, (self).f3])
                    d_942_v6_: D3
                    d_942_v6_ = D3_DC9(d_938_v2_)
                    d_943_v7_: _dafny.Map
                    d_943_v7_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hynbt"))): d_938_v2_})
                    d_944_v8_: _dafny.Array
                    nw169_ = _dafny.Array(None, 22)
                    nw169_[int(0)] = d_938_v2_
                    nw169_[int(1)] = d_938_v2_
                    nw169_[int(2)] = d_938_v2_
                    nw169_[int(3)] = d_938_v2_
                    nw169_[int(4)] = (d_938_v2_ if self.f4 else d_938_v2_)
                    nw169_[int(5)] = d_938_v2_
                    nw169_[int(6)] = d_938_v2_
                    nw169_[int(7)] = (d_939_v3_)[default__.safeIndex(len(default__.fm9((p0)[default__.safeIndex(len(d_940_v4_), len(p0))], self.f4, globalState)), len(d_939_v3_))]
                    nw169_[int(8)] = (d_939_v3_)[default__.safeIndex((d_941_v5_).cardinality, len(d_939_v3_))]
                    nw169_[int(9)] = (d_942_v6_).cf14
                    nw169_[int(10)] = d_938_v2_
                    nw169_[int(11)] = d_938_v2_
                    nw169_[int(12)] = d_938_v2_
                    nw169_[int(13)] = d_938_v2_
                    nw169_[int(14)] = d_938_v2_
                    nw169_[int(15)] = d_938_v2_
                    nw169_[int(16)] = d_938_v2_
                    nw169_[int(17)] = d_938_v2_
                    nw169_[int(18)] = (((d_943_v7_)[(self).f3] if ((self).f3) in (d_943_v7_) else d_938_v2_) if self.f4 else d_938_v2_)
                    nw169_[int(19)] = d_938_v2_
                    nw169_[int(20)] = d_938_v2_
                    nw169_[int(21)] = d_938_v2_
                    d_944_v8_ = nw169_
                    d_944_v8_ = d_944_v8_
                    (globalState).f1 = ((self).f7) - (len((_dafny.SeqWithoutIsStrInference([(self).f3])).set(default__.safeIndex(len(default__.fm10(self.f4, self.f4, (self).f3, _dafny.Map({self.f4: p0}), globalState)), len(_dafny.SeqWithoutIsStrInference([(self).f3]))), 730)))
                    (self).f4 = not(self.f4)
                    if not(True):
                        (globalState).f2 = False
                        (globalState).f1 = (0) - (default__.safeModuloInt(default__.safeDivisionInt(-964, (self).f3), (self).f3))
                        (globalState).f2 = self.f4
                        (self).f4 = False
                        (globalState).f2 = True
                    elif True:
                        d_945_v9_: _dafny.Seq
                        d_945_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "evx"))
                        d_945_v9_ = p0
                        d_946_v10_: _dafny.Array
                        nw170_ = _dafny.Array(False, 26)
                        d_946_v10_ = nw170_
                        index175_ = default__.safeIndex(52, (d_946_v10_).length(0))
                        (d_946_v10_)[index175_] = self.f4
                        index176_ = default__.safeIndex(52, (d_946_v10_).length(0))
                        (d_946_v10_)[index176_] = self.f4
                        d_947_v11_: _dafny.Set
                        d_947_v11_ = _dafny.Set({(self).f7})
                        d_948_v13_: _dafny.Seq
                        def iife67_():
                            coll19_ = _dafny.Map()
                            compr_19_: int
                            for compr_19_ in _dafny.IntegerRange(309, 185):
                                d_949_v12_: int = compr_19_
                                if ((309) <= (d_949_v12_)) and ((d_949_v12_) < (185)):
                                    coll19_[(d_949_v12_) + (514)] = True
                            return _dafny.Map(coll19_)
                        d_948_v13_ = _dafny.SeqWithoutIsStrInference([len(iife67_()
                        ), (self).f3])
                        d_950_v14_: _dafny.Array
                        nw171_ = _dafny.Array(None, 16)
                        nw171_[int(0)] = len(d_947_v11_)
                        nw171_[int(1)] = 613
                        nw171_[int(2)] = (self).f3
                        nw171_[int(3)] = (self).f3
                        nw171_[int(4)] = default__.fm0(globalState)
                        nw171_[int(5)] = (self).f3
                        nw171_[int(6)] = (self).f3
                        nw171_[int(7)] = 373
                        nw171_[int(8)] = len(d_948_v13_)
                        nw171_[int(9)] = (self).f7
                        nw171_[int(10)] = len(d_947_v11_)
                        nw171_[int(11)] = 473
                        nw171_[int(12)] = (self).f7
                        nw171_[int(13)] = (self).f7
                        nw171_[int(14)] = (self).f3
                        nw171_[int(15)] = (self).f7
                        d_950_v14_ = nw171_
                        d_951_v15_: _dafny.Array
                        nw172_ = _dafny.Array(None, 18)
                        nw172_[int(0)] = d_950_v14_
                        nw172_[int(1)] = d_950_v14_
                        nw172_[int(2)] = d_950_v14_
                        nw172_[int(3)] = d_950_v14_
                        nw172_[int(4)] = d_950_v14_
                        nw172_[int(5)] = d_950_v14_
                        nw172_[int(6)] = d_950_v14_
                        nw172_[int(7)] = d_950_v14_
                        nw172_[int(8)] = d_950_v14_
                        nw172_[int(9)] = d_950_v14_
                        nw172_[int(10)] = d_950_v14_
                        nw172_[int(11)] = d_950_v14_
                        nw172_[int(12)] = d_950_v14_
                        nw172_[int(13)] = d_950_v14_
                        nw172_[int(14)] = d_950_v14_
                        nw172_[int(15)] = d_950_v14_
                        nw172_[int(16)] = d_950_v14_
                        nw172_[int(17)] = d_950_v14_
                        d_951_v15_ = nw172_
                        d_952_v16_: _dafny.Map
                        d_952_v16_ = _dafny.Map({(self).f3: d_951_v15_})
                        d_952_v16_ = (d_952_v16_).set(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h"))), d_951_v15_)
                        d_953_v17_: _dafny.Map
                        d_953_v17_ = _dafny.Map({self.f4: d_945_v9_})
                        (globalState).f1 = ((0) - ((self).f3) if (d_948_v13_) != (_dafny.SeqWithoutIsStrInference([len(((d_953_v17_)[self.f4] if (self.f4) in (d_953_v17_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eu"))))])) else default__.safeDivisionInt(122, (0) - ((self).f3)))
                        d_954_v18_: C0
                        nw173_ = C0()
                        nw173_.ctor__(self.f4, d_946_v10_)
                        d_954_v18_ = nw173_
                    pass
            pass
        d_955_v19_: str
        d_955_v19_ = _dafny.CodePoint('k')
        d_955_v19_ = d_955_v19_
        d_956_v20_: _dafny.Map
        d_956_v20_ = _dafny.Map({((self).f3) * ((self).f3): self.f4})
        d_957_v21_: _dafny.Set
        d_957_v21_ = _dafny.Set({self.f4, self.f4})
        d_958_v22_: _dafny.MultiSet
        d_958_v22_ = _dafny.MultiSet([(self).f3, (0) - (len(d_957_v21_)), (self).f3, (self).f7, (0) - ((self).f3)])
        if ((d_956_v20_)[(d_958_v22_).cardinality] if ((d_958_v22_).cardinality) in (d_956_v20_) else self.f4):
            d_959_v23_: _dafny.Array
            nw174_ = _dafny.Array(False, 23)
            d_959_v23_ = nw174_
            d_960_v24_: C0
            nw175_ = C0()
            nw175_.ctor__(self.f4, d_959_v23_)
            d_960_v24_ = nw175_
            d_961_v25_: _dafny.Map
            d_961_v25_ = _dafny.Map({(self).f3: (self).f7})
            d_961_v25_ = (d_961_v25_).set((0) - (default__.fm0(globalState)), len((d_957_v21_) | (d_957_v21_)))
            d_962_v26_: _dafny.Map
            d_962_v26_ = _dafny.Map({d_960_v24_.f10: p0})
            (globalState).f1 = (len(default__.fm10(d_960_v24_.f10, self.f4, (self).f7, d_962_v26_, globalState))) - ((self).f3)
            (self).f4 = True
            (globalState).f1 = (self).f7
        elif True:
            d_963_v27_: _dafny.Map
            d_963_v27_ = _dafny.Map({self.f4: self.f4})
            d_963_v27_ = (d_963_v27_).set(self.f4, self.f4)
            d_964_v28_: _dafny.Array
            nw176_ = _dafny.Array(None, 2)
            nw176_[int(0)] = (self).f3
            nw176_[int(1)] = (self).f7
            d_964_v28_ = nw176_
            index177_ = default__.safeIndex(657, (d_964_v28_).length(0))
            (d_964_v28_)[index177_] = default__.fm0(globalState)
            d_965_v29_: _dafny.Map
            d_965_v29_ = _dafny.Map({(self).f7: (self).f7})
            d_966_v30_: _dafny.Seq
            d_966_v30_ = _dafny.SeqWithoutIsStrInference([d_965_v29_])
            d_967_v31_: _dafny.Map
            d_967_v31_ = _dafny.Map({p0: self.f4})
            d_968_v32_: _dafny.Array
            nw177_ = _dafny.Array(None, 10)
            nw177_[int(0)] = self.f4
            nw177_[int(1)] = True
            nw177_[int(2)] = self.f4
            nw177_[int(3)] = self.f4
            nw177_[int(4)] = True
            nw177_[int(5)] = self.f4
            nw177_[int(6)] = self.f4
            nw177_[int(7)] = ((d_967_v31_)[p0] if (p0) in (d_967_v31_) else self.f4)
            nw177_[int(8)] = self.f4
            nw177_[int(9)] = self.f4
            d_968_v32_ = nw177_
            d_969_v33_: C0
            nw178_ = C0()
            nw178_.ctor__(self.f4, d_968_v32_)
            d_969_v33_ = nw178_
            d_970_v34_: _dafny.Map
            d_970_v34_ = _dafny.Map({(self).f3: d_969_v33_})
            d_971_v35_: _dafny.Seq
            d_971_v35_ = _dafny.SeqWithoutIsStrInference([d_970_v34_, d_970_v34_])
            index178_ = default__.safeIndex(657, (d_964_v28_).length(0))
            (d_964_v28_)[index178_] = len(((d_971_v35_ if ((d_966_v30_)[default__.safeIndex((self).f3, len(d_966_v30_))]) != ((default__.fm11((self).f7, globalState)).set((self).f3, (self).f7)) else d_971_v35_)).set(default__.safeIndex((0) - (((self).f3 if d_969_v33_.f10 else (self).f7)), len((d_971_v35_ if ((d_966_v30_)[default__.safeIndex((self).f3, len(d_966_v30_))]) != ((default__.fm11((self).f7, globalState)).set((self).f3, (self).f7)) else d_971_v35_))), d_970_v34_))
            d_972_v36_: D11
            d_972_v36_ = D11_DC32(d_969_v33_.f10, (self).f3)
            d_973_v37_: T1
            nw179_ = C3()
            nw179_.ctor__(d_955_v19_, p0, (d_972_v36_).cf51, self.f4, (self).f7)
            d_973_v37_ = nw179_
            d_974_v38_: _dafny.Array
            nw180_ = _dafny.Array(_dafny.Seq({}), 26)
            d_974_v38_ = nw180_
            d_975_v39_: _dafny.Seq
            d_975_v39_ = _dafny.SeqWithoutIsStrInference([(self).f7, (d_973_v37_).f7])
            index179_ = default__.safeIndex(856, (d_974_v38_).length(0))
            (d_974_v38_)[index179_] = (d_975_v39_) + (d_975_v39_)
            d_976_v40_: D1
            d_976_v40_ = D1_DC7()
            index180_ = default__.safeIndex(657, (d_964_v28_).length(0))
            index181_ = default__.safeIndex(856, (d_974_v38_).length(0))
            rhs177_ = ((self).f7) - ((self).f7)
            rhs178_ = d_973_v37_
            rhs179_ = d_975_v39_
            rhs180_ = d_976_v40_
            lhs167_ = d_964_v28_
            lhs168_ = default__.safeIndex(657, (d_964_v28_).length(0))
            lhs169_ = d_974_v38_
            lhs170_ = default__.safeIndex(856, (d_974_v38_).length(0))
            lhs167_[lhs168_] = rhs177_
            d_973_v37_ = rhs178_
            lhs169_[lhs170_] = rhs179_
            d_976_v40_ = rhs180_
            d_977_v41_: D7
            d_977_v41_ = D7_DC20(_dafny.Map({(d_964_v28_)[default__.safeIndex(657, (d_964_v28_).length(0))]: p0}))
            d_978_v42_: _dafny.Set
            d_978_v42_ = _dafny.Set({(d_973_v37_).f3})
            d_979_v43_: _dafny.Map
            d_979_v43_ = _dafny.Map({len(d_978_v42_): p0})
            pat_let_tv17_ = d_979_v43_
            def iife68_(_pat_let24_0):
                def iife69_(d_980_dt__update__tmp_h0_):
                    def iife70_(_pat_let25_0):
                        def iife71_(d_981_dt__update_hcf27_h0_):
                            return D7_DC20(d_981_dt__update_hcf27_h0_)
                        return iife71_(_pat_let25_0)
                    return iife70_(pat_let_tv17_)
                return iife69_(_pat_let24_0)
            d_977_v41_ = iife68_(d_977_v41_)
            d_982_v44_: _dafny.Seq
            d_982_v44_ = _dafny.SeqWithoutIsStrInference([p0, p0, p0])
            if (_dafny.SeqWithoutIsStrInference([p0, p0, p0, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gjq"))])) < (d_982_v44_):
                d_983_v45_: _dafny.MultiSet
                d_983_v45_ = _dafny.MultiSet([d_964_v28_])
                d_984_v46_: _dafny.Map
                d_984_v46_ = _dafny.Map({d_983_v45_: (d_964_v28_)[default__.safeIndex(657, (d_964_v28_).length(0))]})
                d_984_v46_ = (d_984_v46_).set((d_983_v45_) | (d_983_v45_), (d_973_v37_).f3)
                (d_969_v33_).f10 = d_969_v33_.f10
                (d_969_v33_).f10 = not(d_969_v33_.f10)
                d_985_v47_: _dafny.MultiSet
                d_985_v47_ = _dafny.MultiSet([d_975_v39_])
                (globalState).f1 = (0) - ((d_985_v47_).cardinality)
                (globalState).f1 = (d_973_v37_).f3
            elif True:
                (d_969_v33_).f10 = default__.fm2((d_973_v37_).f7, globalState)
                d_986_v48_: _dafny.Seq
                d_986_v48_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gxc"))
                rhs181_ = not(d_969_v33_.f10)
                rhs182_ = ((p0) + (d_986_v48_)) + ((d_986_v48_) + (d_986_v48_))
                lhs171_ = d_973_v37_
                lhs171_.f4 = rhs181_
                d_986_v48_ = rhs182_
                d_987_v49_: _dafny.MultiSet
                d_987_v49_ = _dafny.MultiSet([d_963_v27_, (_dafny.Map({not(self.f4): True})).set(d_973_v37_.f4, d_973_v37_.f4), d_963_v27_])
                d_988_v50_: _dafny.Map
                d_988_v50_ = _dafny.Map({default__.safeModuloInt((self).f3, (d_964_v28_)[default__.safeIndex(657, (d_964_v28_).length(0))]): d_987_v49_})
                d_988_v50_ = (d_988_v50_).set((0) - ((d_973_v37_).f3), d_987_v49_)
                index182_ = default__.safeIndex(986, ((d_969_v33_).f11).length(0))
                ((d_969_v33_).f11)[index182_] = self.f4
                index183_ = default__.safeIndex(986, ((d_969_v33_).f11).length(0))
                rhs183_ = d_973_v37_.f4
                rhs184_ = ((d_973_v37_).f3) >= ((d_973_v37_).f3)
                lhs172_ = globalState
                lhs173_ = (d_969_v33_).f11
                lhs174_ = default__.safeIndex(986, ((d_969_v33_).f11).length(0))
                lhs172_.f2 = rhs183_
                lhs173_[lhs174_] = rhs184_
                d_989_v51_: _dafny.Seq
                d_989_v51_ = _dafny.SeqWithoutIsStrInference([(False) and (self.f4)])
                d_989_v51_ = d_989_v51_
        (self).f4 = self.f4
        (globalState).f1 = (self).f7
        d_990_v52_: _dafny.Array
        nw181_ = _dafny.Array(_dafny.Seq({}), 6)
        d_990_v52_ = nw181_
        d_990_v52_ = d_990_v52_


class C5(T1, T0):
    def  __init__(self):
        self._f4: bool = False
        self._f7: int = int(0)
        self._f3: int = int(0)
        self._f9: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    @property
    def f4(self):
        return self._f4
    @f4.setter
    def f4(self, value):
        self._f4 = value
    @property
    def f7(self):
        return self._f7
    @property
    def f3(self):
        return self._f3
    def ctor__(self, f9, f7, f3, f4):
        (self)._f9 = f9
        (self)._f7 = f7
        (self)._f3 = f3
        (self).f4 = f4

    def m1(self, p0, globalState):
        r0: int = int(0)
        d_991_v0_: _dafny.Array
        nw182_ = _dafny.Array(None, 9)
        d_991_v0_ = nw182_
        d_992_v1_: _dafny.Array
        nw183_ = _dafny.Array(_dafny.Array(None, 0), 6)
        d_992_v1_ = nw183_
        d_993_v2_: _dafny.Seq
        d_993_v2_ = _dafny.SeqWithoutIsStrInference([d_991_v0_, d_991_v0_])
        rhs185_ = (d_993_v2_)[default__.safeIndex((_dafny.MultiSet([p0, (self).f3, (self).f3])).cardinality, len(d_993_v2_))]
        rhs186_ = d_992_v1_
        rhs187_ = ((p0 if self.f4 else (self).f7)) * (p0)
        lhs175_ = globalState
        d_991_v0_ = rhs185_
        d_992_v1_ = rhs186_
        lhs175_.f1 = rhs187_
        hi9_ = ((self).f3) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ptajbtlxg"))))
        for d_994_i0_ in range((self).f3, hi9_):
            d_995_v3_: _dafny.Array
            nw184_ = _dafny.Array(False, 8)
            d_995_v3_ = nw184_
            d_995_v3_ = d_995_v3_
            d_996_v4_: _dafny.Seq
            d_996_v4_ = _dafny.SeqWithoutIsStrInference([p0, d_994_i0_, (self).f3])
            d_997_v5_: _dafny.Set
            d_997_v5_ = _dafny.Set({d_996_v4_, d_996_v4_})
            d_998_v6_: _dafny.Map
            d_998_v6_ = _dafny.Map({d_994_i0_: _dafny.Set({self.f4, not(False)})})
            d_999_v7_: _dafny.Seq
            d_999_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "efq"))
            d_1000_v8_: _dafny.Array
            nw185_ = _dafny.Array(None, 17)
            nw185_[int(0)] = len(_dafny.SeqWithoutIsStrInference([(0) - (d_994_i0_) for d_1001_i1_ in range(default__.abs(606))]))
            nw185_[int(1)] = p0
            nw185_[int(2)] = p0
            nw185_[int(3)] = len(d_997_v5_)
            nw185_[int(4)] = -369
            nw185_[int(5)] = (self).f3
            nw185_[int(6)] = (self).f3
            nw185_[int(7)] = p0
            nw185_[int(8)] = p0
            nw185_[int(9)] = p0
            nw185_[int(10)] = len(d_998_v6_)
            nw185_[int(11)] = 687
            nw185_[int(12)] = (self).f7
            nw185_[int(13)] = d_994_i0_
            nw185_[int(14)] = (self).f3
            nw185_[int(15)] = (self).f3
            nw185_[int(16)] = len(d_999_v7_)
            d_1000_v8_ = nw185_
            d_1002_v9_: C2
            nw186_ = C2()
            nw186_.ctor__((d_994_i0_) + (d_994_i0_), d_1000_v8_, (self).f7, self.f4)
            d_1002_v9_ = nw186_
            d_1003_v10_: _dafny.Seq
            d_1003_v10_ = _dafny.SeqWithoutIsStrInference([self.f4])
            d_1004_v11_: _dafny.Map
            d_1004_v11_ = _dafny.Map({len(d_1003_v10_): d_999_v7_})
            d_1005_v12_: _dafny.Seq
            d_1005_v12_ = _dafny.SeqWithoutIsStrInference([self.f4, self.f4, self.f4, (d_1004_v11_) != (d_1004_v11_)])
            d_1006_v13_: _dafny.Seq
            d_1006_v13_ = _dafny.SeqWithoutIsStrInference([d_1005_v12_])
            d_1005_v12_ = (d_1006_v13_)[default__.safeIndex((self).f7, len(d_1006_v13_))]
            d_1007_v14_: _dafny.Array
            def lambda48_(d_1008_i2_):
                return D6_DC18(False, (self).f7)

            init30_ = lambda48_
            nw187_ = _dafny.Array(None, 3)
            for i0_30_ in range(nw187_.length(0)):
                nw187_[i0_30_] = init30_(i0_30_)
            d_1007_v14_ = nw187_
            index184_ = default__.safeIndex(98, (d_1007_v14_).length(0))
            (d_1007_v14_)[index184_] = default__.fm29(globalState)
            d_1009_v15_: D6
            d_1009_v15_ = D6_DC18(not(not((self.f4 if self.f4 else self.f4))), ((d_1002_v9_).f14) - (p0))
            index185_ = default__.safeIndex(98, (d_1007_v14_).length(0))
            (d_1007_v14_)[index185_] = d_1009_v15_
        d_1010_v16_: str
        d_1010_v16_ = _dafny.CodePoint('s')
        d_1011_v17_: _dafny.MultiSet
        d_1011_v17_ = _dafny.MultiSet([d_1010_v16_])
        d_1012_v18_: _dafny.Seq
        d_1012_v18_ = _dafny.SeqWithoutIsStrInference([(self).f7])
        d_1013_v19_: _dafny.Map
        d_1013_v19_ = _dafny.Map({d_1012_v18_: d_1010_v16_})
        d_1014_v20_: _dafny.Map
        d_1014_v20_ = _dafny.Map({len(d_1013_v19_): self.f4})
        d_1015_v21_: _dafny.MultiSet
        d_1015_v21_ = _dafny.MultiSet([((d_1011_v17_)[d_1010_v16_] if (d_1010_v16_) in (d_1011_v17_) else (self).f7), len(_dafny.Map({self.f4: ((d_1014_v20_)[-496] if (-496) in (d_1014_v20_) else self.f4)}))])
        d_1016_i3_: int
        d_1016_i3_ = 0
        with _dafny.label("9"):
            while (d_1015_v21_).isdisjoint(d_1015_v21_):
                with _dafny.c_label("9"):
                    if (d_1016_i3_) >= (100):
                        raise _dafny.Break("9")
                    d_1016_i3_ = (d_1016_i3_) + (1)
                    d_1017_v22_: _dafny.Set
                    d_1017_v22_ = _dafny.Set({self.f4, self.f4})
                    d_1018_v23_: D6
                    d_1018_v23_ = D6_DC19(len(d_1017_v22_))
                    pat_let_tv18_ = p0
                    d_1019_v24_: _dafny.Set
                    def iife72_(_pat_let26_0):
                        def iife73_(d_1020_dt__update__tmp_h0_):
                            def iife74_(_pat_let27_0):
                                def iife75_(d_1021_dt__update_hcf26_h0_):
                                    return D6_DC19(d_1021_dt__update_hcf26_h0_)
                                return iife75_(_pat_let27_0)
                            return iife74_((0) - (pat_let_tv18_))
                        return iife73_(_pat_let26_0)
                    d_1019_v24_ = _dafny.Set({d_1018_v23_, d_1018_v23_, d_1018_v23_, iife72_(d_1018_v23_), D6_DC19((self).f3)})
                    d_1019_v24_ = d_1019_v24_
                    d_1022_v25_: _dafny.Map
                    d_1022_v25_ = _dafny.Map({self.f4: ((0) - (p0)) + (829)})
                    d_1022_v25_ = (d_1022_v25_).set(((0) - ((self).f7)) <= (p0), (self).f3)
                    r0 = p0
                    d_1023_v26_: _dafny.Seq
                    d_1023_v26_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rkni"))
                    d_1024_v27_: _dafny.Map
                    d_1024_v27_ = _dafny.Map({p0: d_1023_v26_})
                    d_1024_v27_ = (d_1024_v27_).set((self).f3, (((d_1024_v27_)[p0] if (p0) in (d_1024_v27_) else d_1023_v26_)) + (d_1023_v26_))
                    pass
            pass
        (self).f4 = self.f4
        d_1025_v28_: _dafny.Array
        def lambda49_(d_1026_i4_):
            return default__.safeModuloInt(d_1026_i4_, 453)

        init31_ = lambda49_
        nw188_ = _dafny.Array(None, 9)
        for i0_31_ in range(nw188_.length(0)):
            nw188_[i0_31_] = init31_(i0_31_)
        d_1025_v28_ = nw188_
        d_1027_v29_: _dafny.Set
        d_1027_v29_ = _dafny.Set({(self).f3})
        d_1028_v30_: _dafny.Seq
        d_1028_v30_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hnjykaw"))
        nw189_ = _dafny.Array(None, 8)
        nw189_[int(0)] = p0
        nw189_[int(1)] = p0
        nw189_[int(2)] = p0
        nw189_[int(3)] = (self).f7
        nw189_[int(4)] = default__.safeDivisionInt(len(d_1027_v29_), len(d_1028_v30_))
        nw189_[int(5)] = ((self).f3) * ((self).f3)
        nw189_[int(6)] = 327
        nw189_[int(7)] = (self).f7
        d_1025_v28_ = nw189_
        rhs188_ = (self).f7
        rhs189_ = self.f4
        rhs190_ = d_1028_v30_
        lhs176_ = globalState
        lhs177_ = globalState
        lhs176_.f1 = rhs188_
        lhs177_.f2 = rhs189_
        d_1028_v30_ = rhs190_
        r0 = 195
        return r0

    def m2(self, p0, p1, p2, p3, globalState):
        r0: _dafny.MultiSet = _dafny.MultiSet({})
        r1: str = _dafny.CodePoint('D')
        r2: _dafny.Map = _dafny.Map({})
        r3: bool = False
        d_1029_v0_: _dafny.Seq
        d_1029_v0_ = _dafny.SeqWithoutIsStrInference([self.f4])
        d_1030_v1_: D1
        d_1030_v1_ = D1_DC7()
        d_1031_v2_: D0
        d_1031_v2_ = D0_DC0(self.f4)
        (globalState).f1 = (((self).f7) - (len(d_1029_v0_))) * (len(default__.fm8(self.f4, d_1030_v1_, d_1031_v2_, (self).f3, globalState)))
        d_1032_v3_: _dafny.Array
        def lambda50_(d_1033_p0_):
            def lambda51_(d_1034_i1_):
                return d_1033_p0_

            return lambda51_

        init32_ = lambda50_(p0)
        nw190_ = _dafny.Array(None, 9)
        for i0_32_ in range(nw190_.length(0)):
            nw190_[i0_32_] = init32_(i0_32_)
        d_1032_v3_ = nw190_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_1032_v3_).length(0)):
            d_1035_i0_: int = guard_loop_2_
            if (True) and (((0) <= (d_1035_i0_)) and ((d_1035_i0_) < ((d_1032_v3_).length(0)))):
                (d_1032_v3_)[(d_1035_i0_)] = (d_1029_v0_)[default__.safeIndex(len(_dafny.Map({(self).f7: p3})), len(d_1029_v0_))]
        d_1036_v4_: _dafny.Seq
        d_1036_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lgks"))
        d_1037_v5_: str
        d_1037_v5_ = _dafny.CodePoint('f')
        d_1038_v6_: T0
        nw191_ = C3()
        nw191_.ctor__(d_1037_v5_, d_1036_v4_, (self).f7, self.f4, (self).f3)
        d_1038_v6_ = nw191_
        d_1039_v7_: D1
        d_1039_v7_ = D1_DC6(d_1038_v6_, d_1036_v4_, d_1038_v6_.f4, p3)
        rhs191_ = d_1036_v4_
        rhs192_ = ((d_1039_v7_).cf10) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uuopatwnv"))).set(default__.safeIndex(840, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uuopatwnv")))), d_1037_v5_))
        d_1036_v4_ = rhs191_
        d_1036_v4_ = rhs192_
        rhs193_ = (d_1038_v6_).f3
        rhs194_ = d_1029_v0_
        rhs195_ = (d_1038_v6_).f3
        lhs178_ = globalState
        lhs179_ = globalState
        lhs178_.f1 = rhs193_
        d_1029_v0_ = rhs194_
        lhs179_.f1 = rhs195_
        d_1040_i2_: int
        d_1040_i2_ = 0
        with _dafny.label("10"):
            while self.f4:
                with _dafny.c_label("10"):
                    if (d_1040_i2_) >= (100):
                        raise _dafny.Break("10")
                    d_1040_i2_ = (d_1040_i2_) + (1)
                    d_1041_v8_: _dafny.Map
                    d_1041_v8_ = _dafny.Map({(d_1038_v6_).f3: d_1038_v6_.f4})
                    d_1042_v9_: D8
                    d_1042_v9_ = D8_DC22(d_1041_v8_)
                    d_1043_v10_: _dafny.Array
                    nw192_ = _dafny.Array(_dafny.Map({}), 6)
                    d_1043_v10_ = nw192_
                    d_1044_v11_: _dafny.Seq
                    d_1044_v11_ = _dafny.SeqWithoutIsStrInference([d_1043_v10_])
                    d_1045_v12_: D12
                    d_1045_v12_ = D12_DC34(d_1043_v10_)
                    d_1046_v13_: _dafny.Array
                    nw193_ = _dafny.Array(None, 24)
                    nw193_[int(0)] = d_1043_v10_
                    nw193_[int(1)] = d_1043_v10_
                    nw193_[int(2)] = d_1043_v10_
                    nw193_[int(3)] = d_1043_v10_
                    nw193_[int(4)] = (d_1044_v11_)[default__.safeIndex((d_1038_v6_).f3, len(d_1044_v11_))]
                    nw193_[int(5)] = d_1043_v10_
                    nw193_[int(6)] = (d_1043_v10_ if d_1038_v6_.f4 else d_1043_v10_)
                    nw193_[int(7)] = d_1043_v10_
                    nw193_[int(8)] = d_1043_v10_
                    nw193_[int(9)] = d_1043_v10_
                    nw193_[int(10)] = d_1043_v10_
                    nw193_[int(11)] = d_1043_v10_
                    nw193_[int(12)] = d_1043_v10_
                    nw193_[int(13)] = d_1043_v10_
                    nw193_[int(14)] = d_1043_v10_
                    nw193_[int(15)] = d_1043_v10_
                    nw193_[int(16)] = (D12_DC34(d_1043_v10_)).cf55
                    nw193_[int(17)] = d_1043_v10_
                    nw193_[int(18)] = (d_1045_v12_).cf55
                    nw193_[int(19)] = d_1043_v10_
                    nw193_[int(20)] = d_1043_v10_
                    nw193_[int(21)] = d_1043_v10_
                    nw193_[int(22)] = d_1043_v10_
                    nw193_[int(23)] = d_1043_v10_
                    d_1046_v13_ = nw193_
                    index186_ = default__.safeIndex(69, (d_1046_v13_).length(0))
                    (d_1046_v13_)[index186_] = d_1043_v10_
                    index187_ = default__.safeIndex(745, (d_1032_v3_).length(0))
                    (d_1032_v3_)[index187_] = p3
                    d_1047_v14_: _dafny.MultiSet
                    d_1047_v14_ = _dafny.MultiSet([(self).f3, (0) - ((self).f7)])
                    pat_let_tv19_ = d_1041_v8_
                    pat_let_tv20_ = d_1038_v6_
                    index188_ = default__.safeIndex(69, (d_1046_v13_).length(0))
                    index189_ = default__.safeIndex(745, (d_1032_v3_).length(0))
                    def iife76_(_pat_let28_0):
                        def iife77_(d_1048_dt__update__tmp_h0_):
                            def iife78_(_pat_let29_0):
                                def iife79_(d_1049_dt__update_hcf31_h0_):
                                    return D8_DC22(d_1049_dt__update_hcf31_h0_)
                                return iife79_(_pat_let29_0)
                            return iife78_((pat_let_tv19_).set((pat_let_tv20_).f3, self.f4))
                        return iife77_(_pat_let28_0)
                    rhs196_ = iife76_(d_1042_v9_)
                    rhs197_ = d_1043_v10_
                    rhs198_ = not(((d_1047_v14_).cardinality) == (((self).f7 if d_1038_v6_.f4 else default__.fm0(globalState))))
                    rhs199_ = (p3 if p3 else p3)
                    lhs180_ = d_1046_v13_
                    lhs181_ = default__.safeIndex(69, (d_1046_v13_).length(0))
                    lhs182_ = d_1032_v3_
                    lhs183_ = default__.safeIndex(745, (d_1032_v3_).length(0))
                    lhs184_ = d_1038_v6_
                    d_1042_v9_ = rhs196_
                    lhs180_[lhs181_] = rhs197_
                    lhs182_[lhs183_] = rhs198_
                    lhs184_.f4 = rhs199_
                    d_1036_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ncpxl"))
                    d_1050_v15_: _dafny.Seq
                    d_1050_v15_ = _dafny.SeqWithoutIsStrInference([p2])
                    d_1051_v16_: D13
                    d_1051_v16_ = D13_DC38(d_1050_v15_)
                    index190_ = default__.safeIndex(745, (d_1032_v3_).length(0))
                    (d_1032_v3_)[index190_] = ((d_1051_v16_).cf68) <= (_dafny.SeqWithoutIsStrInference([p2]))
                    d_1052_v17_: bool
                    d_1053_v18_: int
                    d_1054_v19_: _dafny.Seq
                    d_1055_v20_: int
                    out10_: bool
                    out11_: int
                    out12_: _dafny.Seq
                    out13_: int
                    out10_, out11_, out12_, out13_ = (self).m3((D13_DC39(_dafny.Map({default__.fm2((d_1038_v6_).f3, globalState): (self).f7}))).cf69, len((d_1036_v4_) + (d_1036_v4_)), globalState)
                    d_1052_v17_ = out10_
                    d_1053_v18_ = out11_
                    d_1054_v19_ = out12_
                    d_1055_v20_ = out13_
                    pass
            pass
        d_1056_v21_: _dafny.Seq
        d_1056_v21_ = _dafny.SeqWithoutIsStrInference([534])
        d_1057_v22_: _dafny.Map
        d_1057_v22_ = _dafny.Map({d_1056_v21_: d_1036_v4_})
        d_1058_v23_: D10
        d_1058_v23_ = D10_DC28(d_1057_v22_)
        pat_let_tv21_ = d_1036_v4_
        pat_let_tv22_ = d_1037_v5_
        pat_let_tv23_ = d_1057_v22_
        def lambda52_(source13_):
            if source13_.is_DC29:
                d_1059___mcc_h0_ = source13_.cf42
                d_1060___mcc_h1_ = source13_.cf43
                d_1061_cf43_ = d_1060___mcc_h1_
                d_1062_cf42_ = d_1059___mcc_h0_
                return pat_let_tv21_
            elif source13_.is_DC30:
                d_1063___mcc_h2_ = source13_.cf44
                d_1064___mcc_h3_ = source13_.cf45
                d_1065___mcc_h4_ = source13_.cf46
                d_1066___mcc_h5_ = source13_.cf47
                d_1067___mcc_h6_ = source13_.cf48
                d_1068_cf48_ = d_1067___mcc_h6_
                d_1069_cf47_ = d_1066___mcc_h5_
                d_1070_cf46_ = d_1065___mcc_h4_
                d_1071_cf45_ = d_1064___mcc_h3_
                d_1072_cf44_ = d_1063___mcc_h2_
                return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "whfgm"))
            elif True:
                d_1073___mcc_h7_ = source13_.cf41
                d_1074_cf41_ = d_1073___mcc_h7_
                return _dafny.SeqWithoutIsStrInference([pat_let_tv22_ for d_1075_i3_ in range(default__.abs(894))])

        def iife80_(_pat_let30_0):
            def iife81_(d_1076_dt__update__tmp_h1_):
                def iife82_(_pat_let31_0):
                    def iife83_(d_1077_dt__update_hcf41_h0_):
                        return D10_DC28(d_1077_dt__update_hcf41_h0_)
                    return iife83_(_pat_let31_0)
                return iife82_(pat_let_tv23_)
            return iife81_(_pat_let30_0)
        (globalState).f1 = len(lambda52_(iife80_(d_1058_v23_)))
        d_1078_v24_: D12
        d_1078_v24_ = D12_DC36(p0, d_1038_v6_.f4, (d_1038_v6_).f3)
        d_1079_v25_: _dafny.MultiSet
        d_1079_v25_ = _dafny.MultiSet([(self.f4 if (d_1078_v24_).cf61 else p3)])
        r0 = d_1079_v25_
        d_1080_v26_: _dafny.MultiSet
        d_1080_v26_ = _dafny.MultiSet([(d_1079_v25_).cardinality])
        r1 = (d_1037_v5_ if (d_1080_v26_) != (d_1080_v26_) else _dafny.CodePoint('o'))
        r2 = _dafny.Map({(self).f9: d_1037_v5_})
        r3 = self.f4
        return r0, r1, r2, r3

    def m3(self, p0, p1, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r3: int = int(0)
        hi10_ = default__.fm0(globalState)
        for d_1081_i0_ in range((0) - ((self).f3), hi10_):
            d_1082_v0_: _dafny.Array
            nw194_ = _dafny.Array(_dafny.Array(None, 0), 4)
            d_1082_v0_ = nw194_
            rhs200_ = d_1082_v0_
            rhs201_ = ((self).f3) + (((default__.fm30(self.f4, (self).f7, (self).f3, globalState)).cf66) - ((self).f3))
            d_1082_v0_ = rhs200_
            r1 = rhs201_
            d_1083_v1_: _dafny.Seq
            d_1083_v1_ = _dafny.SeqWithoutIsStrInference([self.f4])
            d_1084_v2_: _dafny.Seq
            d_1084_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "epm"))
            d_1085_v3_: _dafny.Seq
            d_1085_v3_ = _dafny.SeqWithoutIsStrInference([True, (d_1083_v1_)[default__.safeIndex(len(d_1084_v2_), len(d_1083_v1_))]])
            d_1083_v1_ = d_1083_v1_
            d_1086_v4_: _dafny.Set
            d_1086_v4_ = _dafny.Set({self.f4, self.f4, self.f4, self.f4, self.f4})
            d_1087_v5_: _dafny.Seq
            d_1087_v5_ = _dafny.SeqWithoutIsStrInference([(len(_dafny.SeqWithoutIsStrInference([(self).f9]))) - ((self).f3)])
            rhs202_ = ((d_1086_v4_) - (d_1086_v4_)) - (d_1086_v4_)
            rhs203_ = (0) - ((0) - ((d_1081_i0_) * (d_1081_i0_)))
            rhs204_ = (self.f4) or (True)
            rhs205_ = ((d_1087_v5_) + (d_1087_v5_)) + (d_1087_v5_)
            lhs185_ = globalState
            lhs186_ = globalState
            d_1086_v4_ = rhs202_
            lhs185_.f1 = rhs203_
            lhs186_.f2 = rhs204_
            d_1087_v5_ = rhs205_
            d_1088_v6_: _dafny.Array
            def lambda53_(d_1089_v2_):
                def lambda54_(d_1090_i1_):
                    return (d_1090_i1_) + (len(d_1089_v2_))

                return lambda54_

            init33_ = lambda53_(d_1084_v2_)
            nw195_ = _dafny.Array(None, 19)
            for i0_33_ in range(nw195_.length(0)):
                nw195_[i0_33_] = init33_(i0_33_)
            d_1088_v6_ = nw195_
            index191_ = default__.safeIndex(468, (d_1088_v6_).length(0))
            (d_1088_v6_)[index191_] = ((self).f3) * ((self).f3)
            index192_ = default__.safeIndex(468, (d_1088_v6_).length(0))
            rhs206_ = default__.fm0(globalState)
            rhs207_ = (d_1086_v4_).issubset(_dafny.Set({self.f4}))
            rhs208_ = (0) - ((d_1081_i0_) - (p1))
            lhs187_ = globalState
            lhs188_ = d_1088_v6_
            lhs189_ = default__.safeIndex(468, (d_1088_v6_).length(0))
            lhs187_.f1 = rhs206_
            r0 = rhs207_
            lhs188_[lhs189_] = rhs208_
        hi11_ = (self).f3
        for d_1091_i2_ in range(p1, hi11_):
            index193_ = default__.safeIndex(104, ((self).f9).length(0))
            ((self).f9)[index193_] = (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hpexysybx")))) <= ((self).f7)
            d_1092_v7_: _dafny.MultiSet
            d_1092_v7_ = _dafny.MultiSet([self.f4, self.f4, self.f4])
            d_1093_v8_: _dafny.Map
            d_1093_v8_ = _dafny.Map({default__.safeModuloInt((self).f7, (self).f7): default__.fm2(d_1091_i2_, globalState)})
            d_1094_v9_: _dafny.Seq
            d_1094_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "he"))
            index194_ = default__.safeIndex(104, ((self).f9).length(0))
            rhs209_ = (d_1092_v7_).issubset(default__.fm14(self.f4, False, globalState))
            rhs210_ = (self).f7
            rhs211_ = ((d_1093_v8_)[(self).f7] if ((self).f7) in (d_1093_v8_) else not(self.f4))
            rhs212_ = ((self).f7) * (len(d_1094_v9_))
            rhs213_ = (((self).f7 if self.f4 else p1)) - ((0) - ((self).f3))
            lhs190_ = (self).f9
            lhs191_ = default__.safeIndex(104, ((self).f9).length(0))
            lhs192_ = globalState
            lhs193_ = globalState
            lhs190_[lhs191_] = rhs209_
            r1 = rhs210_
            lhs192_.f2 = rhs211_
            lhs193_.f1 = rhs212_
            r3 = rhs213_
            d_1095_v10_: _dafny.Array
            nw196_ = _dafny.Array(_dafny.Array(None, 0), 2)
            d_1095_v10_ = nw196_
            d_1096_v11_: D6
            d_1096_v11_ = D6_DC18(((self).f9)[default__.safeIndex(104, ((self).f9).length(0))], len(d_1094_v9_))
            d_1097_v12_: _dafny.Array
            nw197_ = _dafny.Array(None, 5)
            nw197_[int(0)] = ((self).f9)[default__.safeIndex(104, ((self).f9).length(0))]
            nw197_[int(1)] = self.f4
            nw197_[int(2)] = self.f4
            nw197_[int(3)] = (d_1096_v11_).cf24
            nw197_[int(4)] = not(default__.fm2(p1, globalState))
            d_1097_v12_ = nw197_
            index195_ = default__.safeIndex(923, (d_1095_v10_).length(0))
            (d_1095_v10_)[index195_] = d_1097_v12_
            d_1098_v13_: _dafny.Seq
            d_1098_v13_ = _dafny.SeqWithoutIsStrInference([True, ((self).f9)[default__.safeIndex(104, ((self).f9).length(0))], ((self).f9)[default__.safeIndex(104, ((self).f9).length(0))], not(((self).f9)[default__.safeIndex(104, ((self).f9).length(0))])])
            d_1099_v14_: D12
            d_1099_v14_ = D12_DC35(d_1097_v12_, self.f4, d_1098_v13_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tgykad")), p1)
            index196_ = default__.safeIndex(923, (d_1095_v10_).length(0))
            (d_1095_v10_)[index196_] = (d_1099_v14_).cf56
            r1 = default__.safeDivisionInt(default__.safeModuloInt((self).f7, (self).f3), (0) - (p1))
            d_1100_v15_: _dafny.Map
            d_1100_v15_ = _dafny.Map({self.f4: False})
            d_1101_v16_: _dafny.Array
            nw198_ = _dafny.Array(_dafny.Map({}), 17)
            d_1101_v16_ = nw198_
            d_1102_v17_: D12
            d_1102_v17_ = D12_DC34(d_1101_v16_)
            d_1103_v18_: _dafny.Map
            d_1103_v18_ = _dafny.Map({len(d_1100_v15_): d_1102_v17_})
            d_1104_v19_: _dafny.MultiSet
            d_1104_v19_ = _dafny.MultiSet([d_1103_v18_])
            d_1104_v19_ = ((d_1104_v19_ if True else d_1104_v19_) if False else d_1104_v19_)
        d_1105_v20_: _dafny.Map
        d_1105_v20_ = _dafny.Map({(self).f7: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q"))})
        d_1106_v21_: D7
        d_1106_v21_ = D7_DC20(d_1105_v20_)
        d_1107_v22_: _dafny.MultiSet
        d_1107_v22_ = _dafny.MultiSet([d_1106_v21_, d_1106_v21_])
        if ((D7_DC20(d_1105_v20_) if self.f4 else d_1106_v21_)) in ((d_1107_v22_) | (d_1107_v22_)):
            d_1108_v23_: str
            d_1108_v23_ = _dafny.CodePoint('v')
            d_1109_v24_: _dafny.Set
            d_1109_v24_ = _dafny.Set({d_1108_v23_})
            d_1110_v25_: _dafny.Map
            d_1110_v25_ = _dafny.Map({d_1109_v24_: d_1108_v23_})
            d_1111_v26_: C1
            nw199_ = C1()
            nw199_.ctor__(len(d_1110_v25_), (self).f3, self.f4)
            d_1111_v26_ = nw199_
            d_1112_v27_: _dafny.Set
            d_1112_v27_ = _dafny.Set({d_1111_v26_, d_1111_v26_, d_1111_v26_, d_1111_v26_})
            d_1112_v27_ = (_dafny.Set({d_1111_v26_, d_1111_v26_, d_1111_v26_})) | (_dafny.Set({d_1111_v26_}))
            d_1113_v28_: C4
            nw200_ = C4()
            nw200_.ctor__((self).f3, (self).f3, (p1) >= ((self).f7))
            d_1113_v28_ = nw200_
            r1 = (self).f7
            d_1114_v29_: _dafny.Array
            nw201_ = _dafny.Array(_dafny.Seq({}), 12)
            d_1114_v29_ = nw201_
            d_1115_v30_: _dafny.Seq
            d_1115_v30_ = _dafny.SeqWithoutIsStrInference([self.f4])
            d_1116_v31_: _dafny.Seq
            d_1116_v31_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s"))
            d_1117_v32_: D12
            d_1117_v32_ = D12_DC35((self).f9, True, d_1115_v30_, d_1116_v31_, (self).f3)
            d_1118_v33_: _dafny.Seq
            d_1118_v33_ = _dafny.SeqWithoutIsStrInference([(d_1117_v32_).cf57])
            index197_ = default__.safeIndex(652, (d_1114_v29_).length(0))
            (d_1114_v29_)[index197_] = d_1115_v30_
            index198_ = default__.safeIndex(652, (d_1114_v29_).length(0))
            (d_1114_v29_)[index198_] = d_1115_v30_
            d_1119_v34_: _dafny.Array
            nw202_ = _dafny.Array(False, 2)
            d_1119_v34_ = nw202_
            d_1119_v34_ = (self).f9
        elif True:
            d_1120_v35_: _dafny.Set
            d_1120_v35_ = _dafny.Set({self.f4})
            d_1121_v36_: _dafny.Seq
            d_1121_v36_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tjqkmu"))
            rhs214_ = ((d_1120_v35_).intersection(_dafny.Set({False, self.f4}))) | (d_1120_v35_)
            rhs215_ = ((p1 if self.f4 else (self).f7)) <= (len((d_1121_v36_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "spyelrt")))))
            lhs194_ = self
            d_1120_v35_ = rhs214_
            lhs194_.f4 = rhs215_
            d_1122_v37_: _dafny.Map
            d_1122_v37_ = _dafny.Map({self.f4: self.f4})
            d_1122_v37_ = (d_1122_v37_).set(self.f4, ((self).f3) == (p1))
            d_1123_v38_: _dafny.MultiSet
            d_1123_v38_ = _dafny.MultiSet([(self).f7, (362) + ((self).f3), (self).f7])
            d_1123_v38_ = ((d_1123_v38_).set((0) - ((self).f3), default__.abs(p1))).set(((p0)[self.f4] if (self.f4) in (p0) else p1), default__.abs((0) - ((p1) - (p1))))
            d_1124_v39_: str
            d_1124_v39_ = _dafny.CodePoint('o')
            d_1125_v40_: _dafny.Map
            d_1125_v40_ = _dafny.Map({(d_1121_v36_).set(default__.safeIndex(p1, len(d_1121_v36_)), d_1124_v39_): (self).f7})
            d_1126_v41_: D12
            d_1126_v41_ = D12_DC37(d_1122_v37_, self.f4, ((d_1125_v40_)[_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cmrcf"))] if (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cmrcf"))) in (d_1125_v40_) else (self).f3), (self).f3)
            r3 = (d_1126_v41_).cf67
            d_1127_v42_: _dafny.Map
            d_1127_v42_ = _dafny.Map({((self).f3) + (-416): (self).f7})
            d_1127_v42_ = (d_1127_v42_).set((self).f3, (self).f3)
        d_1128_v43_: _dafny.Seq
        d_1128_v43_ = _dafny.SeqWithoutIsStrInference([self.f4])
        (globalState).f1 = (default__.safeModuloInt(p1, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_1129_i3_ in range(default__.abs(-567))]))) if (d_1128_v43_)[default__.safeIndex((self).f7, len(d_1128_v43_))] else 762)
        d_1130_v44_: C4
        nw203_ = C4()
        nw203_.ctor__((0) - (((self).f7) - ((self).f3)), (0) - (default__.fm0(globalState)), default__.fm2(len(d_1128_v43_), globalState))
        d_1130_v44_ = nw203_
        d_1131_v45_: _dafny.Array
        nw204_ = _dafny.Array(int(0), 13)
        d_1131_v45_ = nw204_
        d_1131_v45_ = d_1131_v45_
        r0 = self.f4
        d_1132_v46_: _dafny.Seq
        d_1132_v46_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gtmwd"))
        r1 = len(d_1132_v46_)
        r2 = default__.fm15(self.f4, default__.safeModuloInt(p1, (self).f3), 639, globalState)
        r3 = (0) - ((self).f3)
        return r0, r1, r2, r3

    @property
    def f9(self):
        return self._f9

class C6(T0, T1):
    def  __init__(self):
        self._f4: bool = False
        self._f7: int = int(0)
        self._f3: int = int(0)
        self._f8: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C6"
    @property
    def f4(self):
        return self._f4
    @f4.setter
    def f4(self, value):
        self._f4 = value
    @property
    def f7(self):
        return self._f7
    @property
    def f3(self):
        return self._f3
    def ctor__(self, f8, f3, f4, f7):
        (self)._f8 = f8
        (self)._f3 = f3
        (self).f4 = f4
        (self)._f7 = f7

    def m1(self, p0, globalState):
        r0: int = int(0)
        if self.f4:
            d_1133_v0_: _dafny.Seq
            d_1133_v0_ = _dafny.SeqWithoutIsStrInference([self.f4])
            d_1134_v1_: _dafny.Map
            d_1134_v1_ = _dafny.Map({self.f4: len(d_1133_v0_)})
            d_1134_v1_ = (d_1134_v1_).set((self.f4) and (self.f4), ((self).f3 if self.f4 else (self).f3))
            d_1135_v2_: C4
            nw205_ = C4()
            nw205_.ctor__(default__.safeModuloInt(p0, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qvoflbfav")))), ((self).f3) * ((self).f3), self.f4)
            d_1135_v2_ = nw205_
            d_1136_v3_: _dafny.Array
            nw206_ = _dafny.Array(False, 21)
            d_1136_v3_ = nw206_
            d_1137_v4_: _dafny.Array
            nw207_ = _dafny.Array(None, 4)
            nw207_[int(0)] = d_1136_v3_
            nw207_[int(1)] = d_1136_v3_
            nw207_[int(2)] = d_1136_v3_
            nw207_[int(3)] = d_1136_v3_
            d_1137_v4_ = nw207_
            d_1137_v4_ = d_1137_v4_
            (globalState).f1 = ((self).f7) + ((0) - ((0) - ((self).f8)))
            d_1138_v5_: C1
            nw208_ = C1()
            nw208_.ctor__(default__.fm0(globalState), (self).f8, self.f4)
            d_1138_v5_ = nw208_
        elif True:
            d_1139_v6_: _dafny.Seq
            d_1139_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kbtb"))
            d_1140_v7_: _dafny.MultiSet
            d_1140_v7_ = _dafny.MultiSet([(d_1139_v6_) + (d_1139_v6_), d_1139_v6_])
            d_1141_v8_: _dafny.Seq
            d_1141_v8_ = _dafny.SeqWithoutIsStrInference([d_1139_v6_, d_1139_v6_])
            d_1140_v7_ = ((d_1140_v7_ if self.f4 else (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oeofpknod")), d_1139_v6_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ou")), d_1139_v6_])).set(d_1139_v6_, default__.abs(p0)))) - ((d_1140_v7_).intersection((_dafny.MultiSet(d_1141_v8_)).set(d_1139_v6_, default__.abs(p0))))
            d_1142_v9_: _dafny.Array
            nw209_ = _dafny.Array(False, 13)
            d_1142_v9_ = nw209_
            index199_ = default__.safeIndex(608, (d_1142_v9_).length(0))
            (d_1142_v9_)[index199_] = self.f4
            index200_ = default__.safeIndex(608, (d_1142_v9_).length(0))
            (d_1142_v9_)[index200_] = self.f4
            d_1143_v10_: str
            d_1143_v10_ = _dafny.CodePoint('o')
            d_1143_v10_ = d_1143_v10_
            d_1144_v11_: _dafny.Array
            nw210_ = _dafny.Array(_dafny.Map({}), 2)
            d_1144_v11_ = nw210_
            index201_ = default__.safeIndex(690, (d_1144_v11_).length(0))
            (d_1144_v11_)[index201_] = _dafny.Map({p0: (d_1142_v9_)[default__.safeIndex(608, (d_1142_v9_).length(0))]})
            d_1145_v12_: _dafny.Map
            d_1145_v12_ = _dafny.Map({(self).f3: default__.fm2(p0, globalState)})
            index202_ = default__.safeIndex(690, (d_1144_v11_).length(0))
            (d_1144_v11_)[index202_] = (d_1145_v12_) | (d_1145_v12_)
            rhs216_ = ((self).f3 if (d_1142_v9_)[default__.safeIndex(608, (d_1142_v9_).length(0))] else (self).f7)
            rhs217_ = ((d_1139_v6_) + (_dafny.SeqWithoutIsStrInference([d_1143_v10_ for d_1146_i0_ in range(default__.abs(588))]))).set(default__.safeIndex(default__.safeModuloInt((self).f8, p0), len((d_1139_v6_) + (_dafny.SeqWithoutIsStrInference([d_1143_v10_ for d_1147_i0_ in range(default__.abs(588))])))), d_1143_v10_)
            rhs218_ = (d_1142_v9_)[default__.safeIndex(608, (d_1142_v9_).length(0))]
            rhs219_ = (default__.fm20((self).f8, globalState)) + (d_1139_v6_)
            rhs220_ = True
            lhs195_ = globalState
            lhs196_ = globalState
            lhs197_ = globalState
            lhs195_.f1 = rhs216_
            d_1139_v6_ = rhs217_
            lhs196_.f2 = rhs218_
            d_1139_v6_ = rhs219_
            lhs197_.f2 = rhs220_
        d_1148_v13_: _dafny.Array
        def lambda55_(d_1149_i1_):
            return self.f4

        init34_ = lambda55_
        nw211_ = _dafny.Array(None, 4)
        for i0_34_ in range(nw211_.length(0)):
            nw211_[i0_34_] = init34_(i0_34_)
        d_1148_v13_ = nw211_
        d_1150_v14_: C5
        nw212_ = C5()
        nw212_.ctor__(d_1148_v13_, ((self).f8 if not(self.f4) else (self).f8), (self).f7, not(self.f4))
        d_1150_v14_ = nw212_
        d_1151_v15_: _dafny.Seq
        d_1151_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "euk"))
        d_1152_v16_: _dafny.Seq
        d_1152_v16_ = _dafny.SeqWithoutIsStrInference([d_1151_v15_, d_1151_v15_])
        if not((d_1151_v15_) not in ((d_1152_v16_) + (d_1152_v16_))):
            d_1153_v17_: _dafny.Map
            d_1153_v17_ = _dafny.Map({False: self.f4})
            d_1154_v18_: _dafny.Set
            d_1154_v18_ = _dafny.Set({len(d_1153_v17_), (self).f8, len(_dafny.Set({(self).f7, (self).f7}))})
            d_1155_v19_: D4
            d_1155_v19_ = D4_DC11(d_1154_v18_)
            source14_ = d_1155_v19_
            if source14_.is_DC12:
                d_1156_v20_: _dafny.Array
                def lambda56_(d_1157_i2_):
                    return default__.safeModuloInt(d_1157_i2_, (self).f3)

                init35_ = lambda56_
                nw213_ = _dafny.Array(None, 15)
                for i0_35_ in range(nw213_.length(0)):
                    nw213_[i0_35_] = init35_(i0_35_)
                d_1156_v20_ = nw213_
                index203_ = default__.safeIndex(61, (d_1156_v20_).length(0))
                (d_1156_v20_)[index203_] = (self).f7
                index204_ = default__.safeIndex(61, (d_1156_v20_).length(0))
                (d_1156_v20_)[index204_] = default__.fm0(globalState)
                d_1158_v21_: C2
                nw214_ = C2()
                nw214_.ctor__(default__.safeModuloInt((self).f3, -240), d_1156_v20_, (self).f3, self.f4)
                d_1158_v21_ = nw214_
                (globalState).f2 = self.f4
                (globalState).f1 = -741
            elif source14_.is_DC11:
                d_1159___mcc_h0_ = source14_.cf18
                d_1160_cf18_ = d_1159___mcc_h0_
                (globalState).f1 = (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jacm")))) - (default__.safeModuloInt((self).f8, (self).f7))
                d_1151_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cmeqw"))
                d_1161_v22_: _dafny.Set
                d_1161_v22_ = _dafny.Set({self.f4})
                d_1162_v23_: _dafny.MultiSet
                d_1162_v23_ = _dafny.MultiSet([(self).f3, p0, len(d_1161_v22_)])
                d_1163_v24_: _dafny.Map
                d_1163_v24_ = _dafny.Map({(self).f3: self.f4})
                d_1164_v25_: _dafny.Seq
                d_1164_v25_ = _dafny.SeqWithoutIsStrInference([self.f4, self.f4, ((d_1163_v24_)[p0] if (p0) in (d_1163_v24_) else self.f4)])
                d_1165_v26_: D3
                d_1165_v26_ = D3_DC10(len(d_1151_v15_), (d_1150_v14_).f9, ((d_1162_v23_)[(0) - ((self).f8)] if ((0) - ((self).f8)) in (d_1162_v23_) else len(d_1164_v25_)))
                d_1166_v28_: C5
                nw215_ = C5()
                def iife84_():
                    coll20_ = _dafny.Set()
                    compr_20_: int
                    for compr_20_ in (d_1163_v24_).keys.Elements:
                        d_1167_v27_: int = compr_20_
                        if (d_1167_v27_) in (d_1163_v24_):
                            coll20_ = coll20_.union(_dafny.Set([(d_1167_v27_) + (-533)]))
                    return _dafny.Set(coll20_)
                nw215_.ctor__(((d_1165_v26_).cf16 if self.f4 else (d_1150_v14_).f9), len(iife84_()
                ), (_dafny.MultiSet(d_1164_v25_)).cardinality, default__.fm2((self).f8, globalState))
                d_1166_v28_ = nw215_
                d_1168_v29_: _dafny.Map
                d_1168_v29_ = _dafny.Map({self.f4: d_1148_v13_})
                d_1169_v30_: C0
                nw216_ = C0()
                nw216_.ctor__(self.f4, ((d_1168_v29_)[self.f4] if (self.f4) in (d_1168_v29_) else (d_1150_v14_).f9))
                d_1169_v30_ = nw216_
            elif True:
                d_1170___mcc_h1_ = source14_.cf19
                d_1171_cf19_ = d_1170___mcc_h1_
                d_1172_v31_: _dafny.Array
                def lambda57_(d_1173_i3_):
                    return (d_1173_i3_) * (862)

                init36_ = lambda57_
                nw217_ = _dafny.Array(None, 28)
                for i0_36_ in range(nw217_.length(0)):
                    nw217_[i0_36_] = init36_(i0_36_)
                d_1172_v31_ = nw217_
                index205_ = default__.safeIndex(5, (d_1172_v31_).length(0))
                (d_1172_v31_)[index205_] = p0
                d_1174_v32_: str
                d_1174_v32_ = _dafny.CodePoint('j')
                index206_ = default__.safeIndex(5, (d_1172_v31_).length(0))
                rhs221_ = default__.safeModuloInt(((0) - ((self).f8)) + ((self).f7), (0) - (len(default__.fm9(d_1174_v32_, self.f4, globalState))))
                rhs222_ = not(self.f4)
                lhs198_ = d_1172_v31_
                lhs199_ = default__.safeIndex(5, (d_1172_v31_).length(0))
                lhs200_ = globalState
                lhs198_[lhs199_] = rhs221_
                lhs200_.f2 = rhs222_
                index207_ = default__.safeIndex(5, (d_1172_v31_).length(0))
                (d_1172_v31_)[index207_] = (p0 if self.f4 else 819)
                d_1175_v33_: _dafny.Seq
                d_1175_v33_ = _dafny.SeqWithoutIsStrInference([((self).f3) - (p0), (self).f3, (self).f8])
                d_1176_v34_: _dafny.Seq
                d_1176_v34_ = _dafny.SeqWithoutIsStrInference([self.f4])
                d_1177_v35_: _dafny.Map
                d_1177_v35_ = _dafny.Map({(self).f7: d_1174_v32_})
                d_1178_v36_: _dafny.Set
                d_1178_v36_ = _dafny.Set({d_1174_v32_, ((d_1177_v35_)[(d_1172_v31_)[default__.safeIndex(5, (d_1172_v31_).length(0))]] if ((d_1172_v31_)[default__.safeIndex(5, (d_1172_v31_).length(0))]) in (d_1177_v35_) else _dafny.CodePoint('k')), d_1174_v32_})
                d_1179_v37_: _dafny.Map
                d_1179_v37_ = _dafny.Map({len(d_1178_v36_): self.f4})
                d_1180_v38_: C1
                nw218_ = C1()
                nw218_.ctor__((self).f3, len(_dafny.SeqWithoutIsStrInference([(self).f8 for d_1181_i4_ in range(default__.abs(542))])), self.f4)
                d_1180_v38_ = nw218_
                d_1182_v39_: _dafny.Map
                d_1182_v39_ = _dafny.Map({233: d_1180_v38_})
                d_1183_v40_: _dafny.MultiSet
                d_1183_v40_ = _dafny.MultiSet([(d_1180_v38_ if self.f4 else d_1180_v38_), d_1180_v38_, d_1180_v38_, ((d_1182_v39_)[(self).f8] if ((self).f8) in (d_1182_v39_) else d_1180_v38_), d_1180_v38_])
                index208_ = default__.safeIndex(5, (d_1172_v31_).length(0))
                rhs223_ = (d_1175_v33_).set(default__.safeIndex(len((_dafny.Map({len(d_1176_v34_): self.f4}) if self.f4 else d_1179_v37_)), len(d_1175_v33_)), p0)
                rhs224_ = ((d_1183_v40_)[(d_1180_v38_ if self.f4 else d_1180_v38_)] if ((d_1180_v38_ if self.f4 else d_1180_v38_)) in (d_1183_v40_) else default__.safeModuloInt(p0, (self).f8))
                rhs225_ = (d_1176_v34_)[default__.safeIndex(((self).f3) + (default__.fm0(globalState)), len(d_1176_v34_))]
                rhs226_ = self.f4
                rhs227_ = (d_1172_v31_)[default__.safeIndex(5, (d_1172_v31_).length(0))]
                lhs201_ = d_1172_v31_
                lhs202_ = default__.safeIndex(5, (d_1172_v31_).length(0))
                lhs203_ = globalState
                lhs204_ = self
                lhs205_ = globalState
                d_1175_v33_ = rhs223_
                lhs201_[lhs202_] = rhs224_
                lhs203_.f2 = rhs225_
                lhs204_.f4 = rhs226_
                lhs205_.f1 = rhs227_
                d_1172_v31_ = d_1172_v31_
            d_1184_v41_: _dafny.Map
            d_1184_v41_ = _dafny.Map({(self.f4) or (not(self.f4)): (self).f3})
            d_1185_v42_: C4
            nw219_ = C4()
            nw219_.ctor__(-621, (self).f7, self.f4)
            d_1185_v42_ = nw219_
            d_1186_v43_: _dafny.Seq
            d_1186_v43_ = _dafny.SeqWithoutIsStrInference([d_1185_v42_, d_1185_v42_])
            d_1187_v44_: _dafny.Map
            d_1187_v44_ = _dafny.Map({not(self.f4): d_1186_v43_})
            d_1188_v45_: _dafny.Seq
            d_1188_v45_ = ((d_1187_v44_)[True] if (True) in (d_1187_v44_) else d_1186_v43_)
            d_1184_v41_ = (d_1184_v41_).set(self.f4, len(((d_1188_v45_)) + (d_1186_v43_)))
            d_1189_v46_: str
            d_1189_v46_ = _dafny.CodePoint('f')
            d_1190_v47_: _dafny.Seq
            d_1190_v47_ = _dafny.SeqWithoutIsStrInference([self.f4])
            d_1191_v48_: C3
            nw220_ = C3()
            nw220_.ctor__(d_1189_v46_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iqrkkklva")), (self).f8, self.f4, len(d_1190_v47_))
            d_1191_v48_ = nw220_
            d_1192_v49_: _dafny.Map
            d_1192_v49_ = _dafny.Map({d_1191_v48_: self.f4})
            d_1193_v50_: _dafny.Map
            d_1193_v50_ = _dafny.Map({self.f4: d_1192_v49_})
            d_1192_v49_ = ((d_1192_v49_) | (d_1192_v49_)) | ((((d_1193_v50_)[self.f4] if (self.f4) in (d_1193_v50_) else (_dafny.Map({d_1191_v48_: self.f4})).set(d_1191_v48_, self.f4))) | (d_1192_v49_))
            d_1194_v51_: _dafny.Array
            def lambda58_(d_1195_i5_):
                return (d_1195_i5_) - ((self).f8)

            init37_ = lambda58_
            nw221_ = _dafny.Array(None, 15)
            for i0_37_ in range(nw221_.length(0)):
                nw221_[i0_37_] = init37_(i0_37_)
            d_1194_v51_ = nw221_
            d_1196_v52_: _dafny.MultiSet
            d_1196_v52_ = _dafny.MultiSet([d_1194_v51_, d_1194_v51_])
            index209_ = default__.safeIndex(958, (d_1194_v51_).length(0))
            (d_1194_v51_)[index209_] = (d_1196_v52_).cardinality
            index210_ = default__.safeIndex(958, (d_1194_v51_).length(0))
            (d_1194_v51_)[index210_] = default__.safeDivisionInt(46, (self).f7)
            d_1197_v53_: _dafny.Set
            d_1197_v53_ = _dafny.Set({self.f4})
            d_1198_v54_: _dafny.Map
            d_1198_v54_ = _dafny.Map({(self).f8: d_1197_v53_})
            index211_ = default__.safeIndex(958, (d_1194_v51_).length(0))
            (d_1194_v51_)[index211_] = len(((((d_1198_v54_)[p0] if (p0) in (d_1198_v54_) else d_1197_v53_)) - (d_1197_v53_)) - (d_1197_v53_))
        elif True:
            d_1199_v55_: _dafny.Seq
            d_1199_v55_ = _dafny.SeqWithoutIsStrInference([(d_1150_v14_).f9, (d_1150_v14_).f9, (d_1150_v14_).f9])
            d_1199_v55_ = ((d_1199_v55_).set(default__.safeIndex((self).f7, len(d_1199_v55_)), (d_1150_v14_).f9)) + (((d_1199_v55_).set(default__.safeIndex(p0, len(d_1199_v55_)), d_1148_v13_)) + (d_1199_v55_))
            d_1200_v56_: _dafny.Map
            d_1200_v56_ = _dafny.Map({not(self.f4): self.f4})
            d_1200_v56_ = d_1200_v56_
            d_1201_v57_: _dafny.Seq
            d_1201_v57_ = _dafny.SeqWithoutIsStrInference([self.f4, self.f4])
            d_1202_v58_: _dafny.Set
            d_1202_v58_ = _dafny.Set({((d_1201_v57_).set(default__.safeIndex((self).f8, len(d_1201_v57_)), self.f4)) + (d_1201_v57_), d_1201_v57_, d_1201_v57_, d_1201_v57_})
            d_1203_v59_: _dafny.Seq
            d_1203_v59_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({d_1201_v57_}), (d_1202_v58_) | (_dafny.Set({d_1201_v57_}))])
            d_1202_v58_ = (d_1203_v59_)[default__.safeIndex(p0, len(d_1203_v59_))]
            d_1204_v60_: _dafny.MultiSet
            d_1204_v60_ = _dafny.MultiSet([self.f4])
            d_1205_v61_: _dafny.Set
            d_1205_v61_ = _dafny.Set({(d_1204_v60_).cardinality, 818})
            d_1151_v15_ = default__.fm15((len(d_1205_v61_)) < ((self).f8), (self).f8, (self).f8, globalState)
            r0 = (0) - ((0) - (((self).f7 if (self.f4) or (self.f4) else (0) - (-893))))
        (globalState).f2 = self.f4
        d_1206_v62_: D4
        d_1206_v62_ = D4_DC12()
        d_1207_v63_: _dafny.Set
        d_1207_v63_ = _dafny.Set({self.f4})
        d_1208_v64_: _dafny.Seq
        d_1208_v64_ = _dafny.SeqWithoutIsStrInference([self.f4])
        d_1209_v65_: _dafny.Seq
        d_1209_v65_ = _dafny.SeqWithoutIsStrInference([p0])
        d_1210_v66_: _dafny.Map
        d_1210_v66_ = _dafny.Map({len(d_1209_v65_): (self).f8})
        d_1211_v67_: D3
        d_1211_v67_ = D3_DC10((self).f7, (d_1150_v14_).f9, (0) - ((self).f3))
        d_1212_v68_: _dafny.Array
        nw222_ = _dafny.Array(None, 24)
        nw222_[int(0)] = (self).f3
        nw222_[int(1)] = p0
        nw222_[int(2)] = (self).f8
        nw222_[int(3)] = 688
        nw222_[int(4)] = len(d_1151_v15_)
        nw222_[int(5)] = len(d_1207_v63_)
        nw222_[int(6)] = p0
        nw222_[int(7)] = (self).f8
        nw222_[int(8)] = (self).f7
        nw222_[int(9)] = (self).f8
        nw222_[int(10)] = len(d_1208_v64_)
        nw222_[int(11)] = default__.fm0(globalState)
        nw222_[int(12)] = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_1213_i6_ in range(default__.abs(831))]))
        nw222_[int(13)] = (_dafny.MultiSet(d_1209_v65_)).cardinality
        nw222_[int(14)] = len((d_1210_v66_).set(p0, p0))
        nw222_[int(15)] = p0
        nw222_[int(16)] = (d_1211_v67_).cf17
        nw222_[int(17)] = len(default__.fm31(globalState))
        nw222_[int(18)] = (self).f3
        nw222_[int(19)] = p0
        nw222_[int(20)] = (self).f8
        nw222_[int(21)] = p0
        nw222_[int(22)] = (self).f8
        nw222_[int(23)] = (self).f3
        d_1212_v68_ = nw222_
        d_1214_v69_: _dafny.Seq
        d_1214_v69_ = _dafny.SeqWithoutIsStrInference([d_1212_v68_])
        d_1215_v70_: C2
        nw223_ = C2()
        nw223_.ctor__(p0, d_1212_v68_, (self).f7, self.f4)
        d_1215_v70_ = nw223_
        d_1216_v71_: _dafny.Seq
        d_1216_v71_ = _dafny.SeqWithoutIsStrInference([d_1215_v70_, d_1215_v70_, d_1215_v70_])
        d_1217_v72_: _dafny.Seq
        d_1217_v72_ = _dafny.SeqWithoutIsStrInference([d_1214_v69_])
        d_1218_v73_: str
        d_1218_v73_ = _dafny.CodePoint('b')
        d_1219_v74_: D8
        d_1219_v74_ = D8_DC23(self.f4, self.f4, self.f4, d_1218_v73_)
        d_1220_v75_: D0
        d_1220_v75_ = D0_DC1(self.f4, self.f4, d_1218_v73_, (d_1215_v70_).f15)
        pat_let_tv24_ = d_1209_v65_
        pat_let_tv25_ = d_1209_v65_
        pat_let_tv26_ = d_1209_v65_
        pat_let_tv27_ = d_1215_v70_
        pat_let_tv28_ = d_1151_v15_
        pat_let_tv29_ = d_1215_v70_
        pat_let_tv30_ = d_1151_v15_
        pat_let_tv31_ = d_1218_v73_
        pat_let_tv32_ = d_1218_v73_
        def lambda59_(source15_):
            if source15_.is_DC23:
                d_1221___mcc_h2_ = source15_.cf32
                d_1222___mcc_h3_ = source15_.cf33
                d_1223___mcc_h4_ = source15_.cf34
                d_1224___mcc_h5_ = source15_.cf35
                d_1225_cf35_ = d_1224___mcc_h5_
                d_1226_cf34_ = d_1223___mcc_h4_
                d_1227_cf33_ = d_1222___mcc_h3_
                d_1228_cf32_ = d_1221___mcc_h2_
                return (pat_let_tv24_).set(default__.safeIndex((_dafny.MultiSet([(D0_DC0(d_1226_cf34_)).cf0])).cardinality, len(pat_let_tv25_)), (self).f8)
            elif source15_.is_DC22:
                d_1229___mcc_h6_ = source15_.cf31
                d_1230_cf31_ = d_1229___mcc_h6_
                return (pat_let_tv26_) + (_dafny.SeqWithoutIsStrInference([(0) - ((pat_let_tv27_).f14), (_dafny.MultiSet([self.f4])).cardinality]))
            elif True:
                d_1231___mcc_h7_ = source15_.cf36
                d_1232_cf36_ = d_1231___mcc_h7_
                return _dafny.SeqWithoutIsStrInference([len((pat_let_tv28_).set(default__.safeIndex((pat_let_tv29_).f14, len(pat_let_tv30_)), pat_let_tv31_)) for d_1233_i7_ in range(default__.abs(724))])

        def iife85_(_pat_let32_0):
            def iife86_(d_1234_dt__update__tmp_h0_):
                def iife87_(_pat_let33_0):
                    def iife88_(d_1235_dt__update_hcf3_h0_):
                        return D0_DC1((d_1234_dt__update__tmp_h0_).cf1, (d_1234_dt__update__tmp_h0_).cf2, d_1235_dt__update_hcf3_h0_, (d_1234_dt__update__tmp_h0_).cf4)
                    return iife88_(_pat_let33_0)
                return iife87_(pat_let_tv32_)
            return iife86_(_pat_let32_0)
        rhs228_ = (d_1206_v62_ if self.f4 else D4_DC12())
        rhs229_ = (d_1217_v72_)[default__.safeIndex((d_1215_v70_).f14, len(d_1217_v72_))]
        rhs230_ = lambda59_(d_1219_v74_)
        rhs231_ = (iife85_(d_1220_v75_)).cf2
        rhs232_ = ((d_1216_v71_) + (d_1216_v71_)) + (d_1216_v71_)
        lhs206_ = globalState
        d_1206_v62_ = rhs228_
        d_1214_v69_ = rhs229_
        d_1209_v65_ = rhs230_
        lhs206_.f2 = rhs231_
        d_1216_v71_ = rhs232_
        d_1236_v76_: _dafny.Map
        d_1236_v76_ = _dafny.Map({False: len(d_1208_v64_)})
        d_1237_v77_: _dafny.MultiSet
        d_1237_v77_ = _dafny.MultiSet([len(d_1209_v65_), ((d_1236_v76_)[self.f4] if (self.f4) in (d_1236_v76_) else (self).f3), (self).f7, (self).f8])
        d_1238_v78_: C3
        nw224_ = C3()
        nw224_.ctor__(d_1218_v73_, _dafny.SeqWithoutIsStrInference([d_1218_v73_ for d_1239_i8_ in range(default__.abs(-483))]), len(d_1151_v15_), self.f4, (((d_1237_v77_).set(p0, default__.abs(p0))).set((self).f8, default__.abs(-529))).cardinality)
        d_1238_v78_ = nw224_
        r0 = (d_1215_v70_).f14
        return r0

    @property
    def f8(self):
        return self._f8

class C7(T0):
    def  __init__(self):
        self._f4: bool = False
        self._f3: int = int(0)
        self.f5: bool = False
        self._f6: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C7"
    @property
    def f4(self):
        return self._f4
    @f4.setter
    def f4(self, value):
        self._f4 = value
    @property
    def f3(self):
        return self._f3
    def ctor__(self, f5, f6, f3, f4):
        (self).f5 = f5
        (self)._f6 = f6
        (self)._f3 = f3
        (self).f4 = f4

    def fm3(self, p0, globalState):
        return _dafny.CodePoint('f')

    def m0(self, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: int = int(0)
        (self).f5 = self.f5
        d_1240_v0_: D1
        d_1240_v0_ = D1_DC5((self).f3)
        hi12_ = ((self).f3) + ((d_1240_v0_).cf8)
        for d_1241_i0_ in range((self).f6, hi12_):
            d_1242_v1_: _dafny.Seq
            d_1242_v1_ = _dafny.SeqWithoutIsStrInference([self.f4])
            d_1243_v2_: _dafny.Seq
            d_1243_v2_ = _dafny.SeqWithoutIsStrInference([not((d_1242_v1_)[default__.safeIndex(242, len(d_1242_v1_))]), self.f5, self.f4, self.f4, self.f4])
            d_1244_v3_: _dafny.Array
            def lambda60_(d_1245_i1_):
                return self.f4

            init38_ = lambda60_
            nw225_ = _dafny.Array(None, 4)
            for i0_38_ in range(nw225_.length(0)):
                nw225_[i0_38_] = init38_(i0_38_)
            d_1244_v3_ = nw225_
            d_1246_v4_: _dafny.Map
            d_1246_v4_ = _dafny.Map({d_1244_v3_: self.f4})
            d_1247_v5_: _dafny.Seq
            d_1247_v5_ = _dafny.SeqWithoutIsStrInference([(self).f6])
            d_1248_v6_: D0
            d_1248_v6_ = D0_DC0(self.f4)
            d_1249_v7_: _dafny.Seq
            d_1249_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wyblf"))
            d_1250_v8_: _dafny.MultiSet
            d_1250_v8_ = _dafny.MultiSet([self.f5])
            d_1251_v9_: _dafny.Array
            nw226_ = _dafny.Array(None, 24)
            nw226_[int(0)] = self.f5
            nw226_[int(1)] = (d_1242_v1_)[default__.safeIndex((self).f6, len(d_1242_v1_))]
            nw226_[int(2)] = self.f4
            nw226_[int(3)] = self.f5
            nw226_[int(4)] = self.f4
            nw226_[int(5)] = (default__.fm4(self.f5, ((d_1246_v4_)[d_1244_v3_] if (d_1244_v3_) in (d_1246_v4_) else self.f5), self.f5, self.f4, globalState)) < (d_1247_v5_)
            nw226_[int(6)] = self.f5
            nw226_[int(7)] = ((self).f3) != (d_1241_i0_)
            nw226_[int(8)] = self.f4
            nw226_[int(9)] = (d_1248_v6_).cf0
            nw226_[int(10)] = self.f4
            nw226_[int(11)] = not(self.f4)
            nw226_[int(12)] = self.f5
            nw226_[int(13)] = self.f5
            nw226_[int(14)] = self.f4
            nw226_[int(15)] = ((self).f6) < (len(d_1249_v7_))
            nw226_[int(16)] = self.f5
            nw226_[int(17)] = self.f5
            nw226_[int(18)] = (_dafny.MultiSet([self.f5, self.f4])).isdisjoint(d_1250_v8_)
            nw226_[int(19)] = self.f4
            nw226_[int(20)] = not (self.f5) or (self.f5)
            nw226_[int(21)] = self.f4
            nw226_[int(22)] = self.f4
            nw226_[int(23)] = not(not(not(self.f4)))
            d_1251_v9_ = nw226_
            d_1252_v10_: _dafny.Seq
            d_1252_v10_ = _dafny.SeqWithoutIsStrInference([d_1251_v9_, d_1244_v3_, d_1251_v9_])
            d_1251_v9_ = (d_1252_v10_)[default__.safeIndex((self).f3, len(d_1252_v10_))]
            d_1253_v11_: _dafny.Map
            d_1253_v11_ = _dafny.Map({default__.safeDivisionInt((0) - (d_1241_i0_), (self).f6): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u"))})
            d_1253_v11_ = (d_1253_v11_).set((self).f3, d_1249_v7_)
            (globalState).f1 = (self).f6
            d_1254_v12_: C6
            nw227_ = C6()
            nw227_.ctor__(default__.safeDivisionInt((self).f3, (0) - ((self).f3)), (self).f3, True, (self).f6)
            d_1254_v12_ = nw227_
        d_1255_v13_: _dafny.Array
        def lambda61_(d_1256_i3_):
            return (d_1256_i3_) + ((self).f3)

        init39_ = lambda61_
        nw228_ = _dafny.Array(None, 27)
        for i0_39_ in range(nw228_.length(0)):
            nw228_[i0_39_] = init39_(i0_39_)
        d_1255_v13_ = nw228_
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_1255_v13_).length(0)):
            d_1257_i2_: int = guard_loop_3_
            if (True) and (((0) <= (d_1257_i2_)) and ((d_1257_i2_) < ((d_1255_v13_).length(0)))):
                (d_1255_v13_)[(d_1257_i2_)] = default__.safeModuloInt(d_1257_i2_, 947)
        d_1258_v14_: _dafny.Seq
        d_1258_v14_ = _dafny.SeqWithoutIsStrInference([self.f4])
        index212_ = default__.safeIndex(527, (d_1255_v13_).length(0))
        (d_1255_v13_)[index212_] = len(d_1258_v14_)
        index213_ = default__.safeIndex(527, (d_1255_v13_).length(0))
        (d_1255_v13_)[index213_] = default__.fm0(globalState)
        (self).f4 = self.f4
        (self).f4 = (self.f4) and (self.f4)
        r0 = d_1258_v14_
        r1 = (self).f3
        return r0, r1

    @property
    def f6(self):
        return self._f6
