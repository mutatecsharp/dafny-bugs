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
        return default__.safeDivisionInt((_dafny.MultiSet([(D4_DC13(len(_dafny.SeqWithoutIsStrInference([47])), -267, True)).cf20, False, False, not(True), False])).cardinality, (len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hrdfpvvp")) for d_0_i0_ in range(default__.abs(39))]))) - (len(_dafny.SeqWithoutIsStrInference([True, (D0_DC0(True)).cf0]))))

    @staticmethod
    def fm1(globalState):
        source0_ = D0_DC2(D0_DC1(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([True, True])]), 484))
        if source0_.is_DC1:
            d_1___mcc_h0_ = source0_.cf1
            d_2___mcc_h1_ = source0_.cf2
            d_3_cf2_ = d_2___mcc_h1_
            d_4_cf1_ = d_1___mcc_h0_
            return D0_DC1(d_4_cf1_, len(_dafny.Map({_dafny.CodePoint('m'): d_3_cf2_})))
        elif source0_.is_DC0:
            d_5___mcc_h2_ = source0_.cf0
            d_6_cf0_ = d_5___mcc_h2_
            return D0_DC1(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_6_cf0_])])), 170)
        elif True:
            d_7___mcc_h3_ = source0_.cf3
            d_8_cf3_ = d_7___mcc_h3_
            return D0_DC1(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([not(False)])]), 261)

    @staticmethod
    def fm2(p0, p1, globalState):
        def iife0_():
            coll0_ = _dafny.Set()
            compr_0_: int
            for compr_0_ in _dafny.IntegerRange(481, 21):
                d_9_v0_: int = compr_0_
                if ((481) <= (d_9_v0_)) and ((d_9_v0_) < (21)):
                    coll0_ = coll0_.union(_dafny.Set([(d_9_v0_) - (-863)]))
            return _dafny.Set(coll0_)
        return (_dafny.Set({(0) - (len(_dafny.SeqWithoutIsStrInference([-262])))})) | (iife0_()
        )

    @staticmethod
    def fm3(p0, globalState):
        return (_dafny.Map({_dafny.Map({False: _dafny.CodePoint('n')}): False})) == (_dafny.Map({_dafny.Map({not(True): _dafny.CodePoint('i')}): True}))

    @staticmethod
    def fm6(p0, p1, p2, p3, globalState):
        return (D3_DC9(_dafny.SeqWithoutIsStrInference([False]))).cf10

    @staticmethod
    def fm7(globalState):
        return (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True])), ((D11_DC32(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "re")) for d_10_i0_ in range(default__.abs(702))])))).cf55).cardinality, len(_dafny.SeqWithoutIsStrInference([517, 44]))])) + ((_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([False]))])) + (_dafny.SeqWithoutIsStrInference([833, (0) - (len(_dafny.Map({not(True): not(False)})))])))

    @staticmethod
    def fm8(p0, p1, p2, p3, globalState):
        return ((_dafny.Map({False: False})) | ((D12_DC35(_dafny.Map({False: True}))).cf58)) | (_dafny.Map({False: not(False)}))

    @staticmethod
    def fm9(globalState):
        def iife1_():
            coll1_ = _dafny.Map()
            compr_1_: int
            for compr_1_ in _dafny.IntegerRange(-10, 691):
                d_11_v0_: int = compr_1_
                if ((-10) <= (d_11_v0_)) and ((d_11_v0_) < (691)):
                    coll1_[(d_11_v0_) + (155)] = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f')]))
            return _dafny.Map(coll1_)
        return D1_DC3(_dafny.SeqWithoutIsStrInference([len(iife1_()
)]))

    @staticmethod
    def fm10(p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hrux"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_12_i0_ in range(default__.abs(346))]))

    @staticmethod
    def fm11(p0, p1, p2, p3, globalState):
        return _dafny.CodePoint('l')

    @staticmethod
    def fm12(p0, p1, p2, globalState):
        return (_dafny.MultiSet([False])) | (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(D3_DC11(len(_dafny.SeqWithoutIsStrInference([False])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ihqlrtng"))), True, not(True))).cf15])))

    @staticmethod
    def fm13(p0, globalState):
        return ((_dafny.Map({False: -488})) | (_dafny.Map({True: (0) - (-168)}))) | ((_dafny.Map({False: -929})) | (_dafny.Map({False: -715})))

    @staticmethod
    def fm14(globalState):
        def iife2_():
            coll2_ = _dafny.Set()
            compr_2_: str
            for compr_2_ in (_dafny.MultiSet([_dafny.CodePoint('b'), _dafny.CodePoint('d')])).Elements:
                d_13_v0_: str = compr_2_
                if (d_13_v0_) in (_dafny.MultiSet([_dafny.CodePoint('b'), _dafny.CodePoint('d')])):
                    coll2_ = coll2_.union(_dafny.Set([d_13_v0_]))
            return _dafny.Set(coll2_)
        return (_dafny.Set({_dafny.CodePoint('u')})) - ((iife2_()
        ) | (_dafny.Set({_dafny.CodePoint('q'), _dafny.CodePoint('n')})))

    @staticmethod
    def fm15(p0, p1, p2, globalState):
        return _dafny.Map({738: (True) == (not(not(False)))})

    @staticmethod
    def fm16(p0, p1, p2, p3, globalState):
        def iife3_():
            coll3_ = _dafny.Map()
            compr_3_: int
            for compr_3_ in (_dafny.Set({260})).Elements:
                d_14_v0_: int = compr_3_
                if (d_14_v0_) in (_dafny.Set({260})):
                    coll3_[default__.safeModuloInt(d_14_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uvbomwi"))))] = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_15_i0_ in range(default__.abs(786))]))
            return _dafny.Map(coll3_)
        def iife4_():
            coll4_ = _dafny.Map()
            compr_4_: int
            for compr_4_ in _dafny.IntegerRange(970, -978):
                d_16_v1_: int = compr_4_
                if ((970) <= (d_16_v1_)) and ((d_16_v1_) < (-978)):
                    coll4_[(d_16_v1_) * (137)] = 484
            return _dafny.Map(coll4_)
        return (_dafny.Map({-350: len(_dafny.Set({_dafny.CodePoint('q'), _dafny.CodePoint('l')}))})) | ((iife3_()
        ) | (iife4_()
        ))

    @staticmethod
    def fm17(p0, p1, p2, globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True])])

    @staticmethod
    def m0(p0, p1, p2, globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        hi0_ = p0
        for d_17_i0_ in range(p0, hi0_):
            d_18_v0_: _dafny.Array
            def lambda0_(d_19_p0_, d_20_i0_):
                def lambda1_(d_21_i1_):
                    return (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_19_p0_, len(_dafny.Set({d_19_p0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wq")))})), (_dafny.MultiSet([d_19_p0_])).cardinality]) for d_22_i2_ in range(default__.abs(899))]) if False else _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_20_i0_, d_20_i0_]), _dafny.SeqWithoutIsStrInference([143 for d_23_i3_ in range(default__.abs(466))]), _dafny.SeqWithoutIsStrInference([d_19_p0_])]))

                return lambda1_

            init0_ = lambda0_(p0, d_17_i0_)
            nw0_ = _dafny.Array(None, 25)
            for i0_0_ in range(nw0_.length(0)):
                nw0_[i0_0_] = init0_(i0_0_)
            d_18_v0_ = nw0_
            d_24_v1_: _dafny.Set
            d_24_v1_ = _dafny.Set({(0) - (p0)})
            d_25_v2_: _dafny.Seq
            d_25_v2_ = _dafny.SeqWithoutIsStrInference([p0, len(d_24_v1_), p0])
            d_26_v3_: _dafny.Seq
            d_26_v3_ = _dafny.SeqWithoutIsStrInference([d_25_v2_])
            index0_ = default__.safeIndex(570, (d_18_v0_).length(0))
            (d_18_v0_)[index0_] = d_26_v3_
            d_27_v4_: bool
            d_27_v4_ = False
            d_28_v5_: _dafny.Array
            nw1_ = _dafny.Array(_dafny.Set({}), 26)
            d_28_v5_ = nw1_
            index1_ = default__.safeIndex(698, (d_28_v5_).length(0))
            (d_28_v5_)[index1_] = d_24_v1_
            d_29_v6_: str
            d_29_v6_ = _dafny.CodePoint('n')
            d_30_v7_: _dafny.Map
            d_30_v7_ = _dafny.Map({p0: d_29_v6_})
            d_31_v8_: C0
            nw2_ = C0()
            nw2_.ctor__(d_17_i0_, (0) - (d_17_i0_), True)
            d_31_v8_ = nw2_
            d_32_v9_: D2
            d_32_v9_ = D2_DC7(len(d_30_v7_), d_31_v8_)
            index2_ = default__.safeIndex(570, (d_18_v0_).length(0))
            index3_ = default__.safeIndex(698, (d_28_v5_).length(0))
            rhs0_ = (_dafny.SeqWithoutIsStrInference([d_25_v2_ for d_33_i4_ in range(default__.abs(625))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p0 for d_34_i5_ in range(default__.abs(-914))])]))
            rhs1_ = d_27_v4_
            rhs2_ = (d_32_v9_).cf7
            rhs3_ = (d_31_v8_).f9
            rhs4_ = d_24_v1_
            lhs0_ = d_18_v0_
            lhs1_ = default__.safeIndex(570, (d_18_v0_).length(0))
            lhs2_ = globalState
            lhs3_ = globalState
            lhs4_ = d_28_v5_
            lhs5_ = default__.safeIndex(698, (d_28_v5_).length(0))
            lhs0_[lhs1_] = rhs0_
            d_27_v4_ = rhs1_
            lhs2_.f3 = rhs2_
            lhs3_.f6 = rhs3_
            lhs4_[lhs5_] = rhs4_
            d_35_v10_: _dafny.Array
            nw3_ = _dafny.Array(None, 21)
            d_35_v10_ = nw3_
            index4_ = default__.safeIndex(694, (d_35_v10_).length(0))
            (d_35_v10_)[index4_] = d_31_v8_
            index5_ = default__.safeIndex(694, (d_35_v10_).length(0))
            (d_35_v10_)[index5_] = d_31_v8_
            d_27_v4_ = d_27_v4_
            if d_27_v4_:
                d_36_v11_: _dafny.Array
                def lambda2_(d_37_p1_):
                    def lambda3_(d_38_i6_):
                        return default__.safeModuloInt(d_38_i6_, (d_37_p1_).cardinality)

                    return lambda3_

                init1_ = lambda2_(p1)
                nw4_ = _dafny.Array(None, 3)
                for i0_1_ in range(nw4_.length(0)):
                    nw4_[i0_1_] = init1_(i0_1_)
                d_36_v11_ = nw4_
                d_39_v12_: _dafny.Seq
                d_39_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kn"))
                d_40_v13_: _dafny.Set
                d_40_v13_ = _dafny.Set({d_31_v8_})
                d_41_v14_: _dafny.MultiSet
                d_41_v14_ = _dafny.MultiSet([d_40_v13_])
                index6_ = default__.safeIndex(243, (d_36_v11_).length(0))
                (d_36_v11_)[index6_] = default__.safeDivisionInt(len(d_39_v12_), ((d_41_v14_)[_dafny.Set({(d_35_v10_)[default__.safeIndex(694, (d_35_v10_).length(0))], d_31_v8_})] if (_dafny.Set({(d_35_v10_)[default__.safeIndex(694, (d_35_v10_).length(0))], d_31_v8_})) in (d_41_v14_) else (d_25_v2_)[default__.safeIndex(len(d_39_v12_), len(d_25_v2_))]))
                index7_ = default__.safeIndex(243, (d_36_v11_).length(0))
                (d_36_v11_)[index7_] = 551
                d_42_v15_: _dafny.Map
                d_42_v15_ = _dafny.Map({(d_31_v8_).f9: (d_31_v8_).f9})
                (globalState).f6 = (42) * (default__.safeDivisionInt((d_31_v8_).f9, len(d_42_v15_)))
                d_43_v16_: _dafny.Seq
                d_43_v16_ = _dafny.SeqWithoutIsStrInference([d_27_v4_, d_27_v4_, (not(not(d_27_v4_))) not in (p1)])
                d_44_v17_: D4
                d_44_v17_ = D4_DC14((d_35_v10_)[default__.safeIndex(694, (d_35_v10_).length(0))], d_27_v4_)
                d_45_v18_: _dafny.Map
                d_45_v18_ = _dafny.Map({(d_35_v10_)[default__.safeIndex(694, (d_35_v10_).length(0))]: d_29_v6_})
                d_46_v19_: _dafny.Map
                d_46_v19_ = _dafny.Map({d_44_v17_: ((d_45_v18_)[(d_35_v10_)[default__.safeIndex(694, (d_35_v10_).length(0))]] if ((d_35_v10_)[default__.safeIndex(694, (d_35_v10_).length(0))]) in (d_45_v18_) else d_29_v6_)})
                d_47_v20_: _dafny.MultiSet
                d_47_v20_ = _dafny.MultiSet([len(d_46_v19_)])
                d_48_v21_: T0
                nw5_ = C0()
                nw5_.ctor__(d_17_i0_, len(d_43_v16_), d_27_v4_)
                d_48_v21_ = nw5_
                d_49_v22_: _dafny.Seq
                d_49_v22_ = _dafny.SeqWithoutIsStrInference([d_48_v21_, d_48_v21_, d_48_v21_])
                index8_ = default__.safeIndex(243, (d_36_v11_).length(0))
                rhs5_ = d_43_v16_
                rhs6_ = d_29_v6_
                rhs7_ = (0) - ((d_47_v20_).cardinality)
                rhs8_ = (((0) - (len(d_49_v22_))) * ((d_36_v11_)[default__.safeIndex(243, (d_36_v11_).length(0))])) + ((d_31_v8_).f8)
                lhs6_ = globalState
                lhs7_ = d_36_v11_
                lhs8_ = default__.safeIndex(243, (d_36_v11_).length(0))
                d_43_v16_ = rhs5_
                d_29_v6_ = rhs6_
                lhs6_.f6 = rhs7_
                lhs7_[lhs8_] = rhs8_
                d_50_v23_: D6
                d_50_v23_ = D6_DC17(d_48_v21_.f7, False, d_27_v4_)
                d_51_v24_: _dafny.Seq
                d_51_v24_ = _dafny.SeqWithoutIsStrInference([d_50_v23_, d_50_v23_, d_50_v23_])
                d_27_v4_ = ((d_51_v24_) + (_dafny.SeqWithoutIsStrInference([d_50_v23_]))) < ((_dafny.SeqWithoutIsStrInference([d_50_v23_ for d_52_i7_ in range(default__.abs(576))])) + (((d_51_v24_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ljh"))), len(d_51_v24_)), d_50_v23_)).set(default__.safeIndex((d_31_v8_).f8, len((d_51_v24_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ljh"))), len(d_51_v24_)), d_50_v23_))), D6_DC17(d_48_v21_.f7, d_48_v21_.f7, d_48_v21_.f7))))
                d_53_v25_: _dafny.Seq
                d_53_v25_ = _dafny.SeqWithoutIsStrInference([d_43_v16_])
                d_54_v26_: _dafny.MultiSet
                d_54_v26_ = _dafny.MultiSet([d_43_v16_, d_43_v16_, p2, p2])
                d_55_v27_: _dafny.Seq
                d_55_v27_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet(d_53_v25_), d_54_v26_])
                d_56_v28_: D0
                d_56_v28_ = D0_DC1((d_55_v27_)[default__.safeIndex(len(d_39_v12_), len(d_55_v27_))], default__.safeModuloInt(233, p0))
                rhs9_ = d_43_v16_
                rhs10_ = (d_31_v8_ if default__.fm3((d_36_v11_)[default__.safeIndex(243, (d_36_v11_).length(0))], globalState) else (d_35_v10_)[default__.safeIndex(694, (d_35_v10_).length(0))])
                rhs11_ = d_56_v28_
                d_43_v16_ = rhs9_
                d_31_v8_ = rhs10_
                d_56_v28_ = rhs11_
            elif True:
                d_24_v1_ = (d_28_v5_)[default__.safeIndex(698, (d_28_v5_).length(0))]
                d_57_v29_: _dafny.Seq
                d_57_v29_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dled"))
                d_58_v30_: _dafny.Seq
                d_58_v30_ = _dafny.SeqWithoutIsStrInference([d_57_v29_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vsxctdar")), d_57_v29_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gmyfssb"))])
                d_58_v30_ = d_58_v30_
                d_59_v31_: C0
                nw6_ = C0()
                nw6_.ctor__((d_31_v8_).f8, (d_31_v8_).f9, (p0) > (len(_dafny.SeqWithoutIsStrInference([d_29_v6_ for d_60_i8_ in range(default__.abs(641))]))))
                d_59_v31_ = nw6_
                d_61_v32_: C0
                nw7_ = C0()
                nw7_.ctor__(len((_dafny.SeqWithoutIsStrInference([d_29_v6_ for d_62_i9_ in range(default__.abs(38))])).set(default__.safeIndex((d_31_v8_).f8, len(_dafny.SeqWithoutIsStrInference([d_29_v6_ for d_63_i9_ in range(default__.abs(38))]))), d_29_v6_)), 403, d_27_v4_)
                d_61_v32_ = nw7_
                d_64_v33_: _dafny.Map
                d_64_v33_ = _dafny.Map({p0: False})
                d_65_v34_: _dafny.Map
                d_65_v34_ = _dafny.Map({((d_64_v33_)[(0) - (d_17_i0_)] if ((0) - (d_17_i0_)) in (d_64_v33_) else d_27_v4_): (d_25_v2_)[default__.safeIndex((0) - (p0), len(d_25_v2_))]})
                d_66_v35_: D3
                d_66_v35_ = D3_DC10(d_29_v6_, d_27_v4_)
                d_65_v34_ = (d_65_v34_).set((d_66_v35_).cf12, default__.safeModuloInt((d_31_v8_).f9, -391))
        d_67_v36_: bool
        d_67_v36_ = True
        d_68_v37_: _dafny.Seq
        d_68_v37_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "deqbibs"))
        hi1_ = p0
        for d_69_i10_ in range((len(d_68_v37_) if not(d_67_v36_) else len(d_68_v37_)), hi1_):
            d_70_v38_: _dafny.Array
            nw8_ = _dafny.Array(_dafny.MultiSet({}), 8)
            d_70_v38_ = nw8_
            d_71_v39_: _dafny.Seq
            d_71_v39_ = _dafny.SeqWithoutIsStrInference([p0])
            d_72_v40_: str
            d_72_v40_ = _dafny.CodePoint('x')
            index9_ = default__.safeIndex(102, (d_70_v38_).length(0))
            (d_70_v38_)[index9_] = _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([default__.fm0(340, d_67_v36_, d_71_v39_, d_72_v40_, globalState), p0, d_69_i10_])]))
            d_73_v41_: _dafny.Array
            def lambda4_(d_74_v40_):
                def lambda5_(d_75_i11_):
                    return (d_74_v40_) in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uyntxwc")))

                return lambda5_

            init2_ = lambda4_(d_72_v40_)
            nw9_ = _dafny.Array(None, 5)
            for i0_2_ in range(nw9_.length(0)):
                nw9_[i0_2_] = init2_(i0_2_)
            d_73_v41_ = nw9_
            index10_ = default__.safeIndex(756, (d_73_v41_).length(0))
            (d_73_v41_)[index10_] = (p0) == (p0)
            d_76_v42_: _dafny.Array
            nw10_ = _dafny.Array(_dafny.MultiSet({}), 10)
            d_76_v42_ = nw10_
            index11_ = default__.safeIndex(780, (d_76_v42_).length(0))
            (d_76_v42_)[index11_] = p1
            d_77_v43_: D6
            d_77_v43_ = D6_DC17(d_67_v36_, d_67_v36_, d_67_v36_)
            d_78_v44_: _dafny.Map
            d_78_v44_ = _dafny.Map({d_77_v43_: p0})
            d_79_v45_: _dafny.MultiSet
            d_79_v45_ = _dafny.MultiSet([d_78_v44_, d_78_v44_, d_78_v44_, d_78_v44_])
            d_80_v46_: C0
            nw11_ = C0()
            nw11_.ctor__(666, d_69_i10_, False)
            d_80_v46_ = nw11_
            d_81_v47_: D2
            d_81_v47_ = D2_DC7(p0, d_80_v46_)
            d_82_v48_: _dafny.MultiSet
            d_82_v48_ = _dafny.MultiSet([(0) - (p0), (d_81_v47_).cf7, p0, p0])
            d_83_v49_: D1
            d_83_v49_ = D1_DC3(d_71_v39_)
            d_84_v50_: D7
            d_84_v50_ = D7_DC19(_dafny.MultiSet(d_71_v39_))
            d_85_v51_: _dafny.MultiSet
            d_85_v51_ = _dafny.MultiSet([d_82_v48_, _dafny.MultiSet((d_83_v49_).cf4), (d_84_v50_).cf29, d_82_v48_])
            index12_ = default__.safeIndex(102, (d_70_v38_).length(0))
            index13_ = default__.safeIndex(756, (d_73_v41_).length(0))
            index14_ = default__.safeIndex(780, (d_76_v42_).length(0))
            rhs12_ = not (not(False)) or ((d_79_v45_).isdisjoint(d_79_v45_))
            rhs13_ = d_85_v51_
            rhs14_ = d_73_v41_
            rhs15_ = (default__.fm7(globalState)) != (d_71_v39_)
            rhs16_ = _dafny.MultiSet(p2)
            lhs9_ = d_70_v38_
            lhs10_ = default__.safeIndex(102, (d_70_v38_).length(0))
            lhs11_ = d_73_v41_
            lhs12_ = default__.safeIndex(756, (d_73_v41_).length(0))
            lhs13_ = d_76_v42_
            lhs14_ = default__.safeIndex(780, (d_76_v42_).length(0))
            d_67_v36_ = rhs12_
            lhs9_[lhs10_] = rhs13_
            d_73_v41_ = rhs14_
            lhs11_[lhs12_] = rhs15_
            lhs13_[lhs14_] = rhs16_
            d_86_v52_: T0
            nw12_ = C0()
            nw12_.ctor__((d_80_v46_).f9, p0, not((d_73_v41_)[default__.safeIndex(756, (d_73_v41_).length(0))]))
            d_86_v52_ = nw12_
            d_87_v53_: _dafny.Seq
            d_87_v53_ = _dafny.SeqWithoutIsStrInference([d_86_v52_])
            d_88_v54_: _dafny.Array
            nw13_ = _dafny.Array(None, 4)
            nw13_[int(0)] = p0
            nw13_[int(1)] = default__.safeDivisionInt((d_80_v46_).f8, len(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xpqepkn")): (d_87_v53_)[default__.safeIndex(921, len(d_87_v53_))]})))
            nw13_[int(2)] = (d_80_v46_).f9
            nw13_[int(3)] = default__.safeModuloInt(d_69_i10_, (d_80_v46_).f9)
            d_88_v54_ = nw13_
            index15_ = default__.safeIndex(916, (d_88_v54_).length(0))
            (d_88_v54_)[index15_] = p0
            index16_ = default__.safeIndex(916, (d_88_v54_).length(0))
            (d_88_v54_)[index16_] = (d_80_v46_).f8
            d_89_v55_: _dafny.Seq
            d_89_v55_ = _dafny.SeqWithoutIsStrInference([d_80_v46_])
            d_90_v56_: _dafny.Set
            d_90_v56_ = _dafny.Set({d_69_i10_, len(d_89_v55_)})
            d_91_v57_: _dafny.Set
            d_91_v57_ = _dafny.Set({(d_90_v56_) | (_dafny.Set({len(p2)}))})
            d_91_v57_ = _dafny.Set({d_90_v56_})
            (globalState).f3 = -293
        d_92_v58_: _dafny.Map
        d_92_v58_ = _dafny.Map({p0: d_68_v37_})
        d_68_v37_ = ((d_92_v58_)[(0) - (((p1)[False] if (False) in (p1) else p0))] if ((0) - (((p1)[False] if (False) in (p1) else p0))) in (d_92_v58_) else d_68_v37_)
        if (p0) != (784):
            d_93_v59_: _dafny.Set
            d_93_v59_ = _dafny.Set({(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lmjjoibd"))) + (d_68_v37_), d_68_v37_, d_68_v37_, d_68_v37_})
            d_93_v59_ = ((d_93_v59_).intersection(d_93_v59_)) - (d_93_v59_)
            if (len(d_68_v37_)) >= (106):
                d_94_v60_: _dafny.Seq
                d_94_v60_ = _dafny.SeqWithoutIsStrInference([p0])
                d_95_v61_: str
                d_95_v61_ = _dafny.CodePoint('d')
                d_96_v62_: T0
                nw14_ = C0()
                nw14_.ctor__(default__.fm0(p0, d_67_v36_, d_94_v60_, d_95_v61_, globalState), len(d_68_v37_), d_67_v36_)
                d_96_v62_ = nw14_
                d_97_v63_: _dafny.MultiSet
                d_97_v63_ = _dafny.MultiSet([d_96_v62_, d_96_v62_])
                d_67_v36_ = ((d_97_v63_ if d_67_v36_ else d_97_v63_)).ispropersubset(d_97_v63_)
                d_98_v64_: C0
                nw15_ = C0()
                nw15_.ctor__(520, p0, d_67_v36_)
                d_98_v64_ = nw15_
                d_99_v65_: _dafny.Seq
                d_99_v65_ = _dafny.SeqWithoutIsStrInference([d_98_v64_])
                d_100_v66_: _dafny.Array
                nw16_ = _dafny.Array(None, 16)
                nw16_[int(0)] = default__.fm0(p0, d_67_v36_, _dafny.SeqWithoutIsStrInference([len(p2)]), d_95_v61_, globalState)
                nw16_[int(1)] = 483
                nw16_[int(2)] = len(d_99_v65_)
                nw16_[int(3)] = p0
                nw16_[int(4)] = (d_98_v64_).f9
                nw16_[int(5)] = (d_98_v64_).f8
                nw16_[int(6)] = (d_98_v64_).f9
                nw16_[int(7)] = (d_98_v64_).f8
                nw16_[int(8)] = -166
                nw16_[int(9)] = (d_98_v64_).f9
                nw16_[int(10)] = (d_98_v64_).f9
                nw16_[int(11)] = (d_98_v64_).f8
                nw16_[int(12)] = (d_98_v64_).f9
                nw16_[int(13)] = (d_98_v64_).f8
                nw16_[int(14)] = (0) - ((d_94_v60_)[default__.safeIndex(p0, len(d_94_v60_))])
                nw16_[int(15)] = (0) - ((d_98_v64_).f8)
                d_100_v66_ = nw16_
                d_101_v67_: D7
                d_101_v67_ = D7_DC20(d_67_v36_, d_96_v62_.f7, d_98_v64_)
                d_102_v68_: _dafny.Map
                d_102_v68_ = _dafny.Map({d_100_v66_: (d_101_v67_ if d_67_v36_ else d_101_v67_)})
                d_103_v69_: _dafny.Map
                d_103_v69_ = _dafny.Map({d_67_v36_: d_100_v66_})
                d_104_v70_: _dafny.Seq
                d_104_v70_ = _dafny.SeqWithoutIsStrInference([d_102_v68_])
                d_102_v68_ = (_dafny.Map({((d_103_v69_)[d_67_v36_] if (d_67_v36_) in (d_103_v69_) else d_100_v66_): d_101_v67_})) | ((d_104_v70_)[default__.safeIndex((d_98_v64_).f8, len(d_104_v70_))])
                d_105_v71_: _dafny.Map
                d_105_v71_ = _dafny.Map({d_95_v61_: d_96_v62_})
                d_106_v72_: _dafny.Map
                d_106_v72_ = _dafny.Map({d_96_v62_.f7: (d_98_v64_).f9})
                d_107_v73_: _dafny.MultiSet
                d_107_v73_ = _dafny.MultiSet([d_106_v72_, d_106_v72_])
                d_108_v74_: _dafny.Array
                nw17_ = _dafny.Array(None, 15)
                nw17_[int(0)] = (d_105_v71_).set(default__.fm11(d_107_v73_, p0, d_68_v37_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aicyyk")), globalState), d_96_v62_)
                nw17_[int(1)] = d_105_v71_
                nw17_[int(2)] = d_105_v71_
                nw17_[int(3)] = d_105_v71_
                nw17_[int(4)] = d_105_v71_
                nw17_[int(5)] = _dafny.Map({d_95_v61_: d_96_v62_})
                nw17_[int(6)] = d_105_v71_
                nw17_[int(7)] = d_105_v71_
                nw17_[int(8)] = d_105_v71_
                nw17_[int(9)] = d_105_v71_
                nw17_[int(10)] = d_105_v71_
                nw17_[int(11)] = d_105_v71_
                nw17_[int(12)] = d_105_v71_
                nw17_[int(13)] = d_105_v71_
                nw17_[int(14)] = (d_105_v71_).set(d_95_v61_, d_96_v62_)
                d_108_v74_ = nw17_
                d_109_v75_: _dafny.Set
                d_109_v75_ = _dafny.Set({d_94_v60_, d_94_v60_})
                d_110_v76_: _dafny.Map
                d_110_v76_ = _dafny.Map({d_108_v74_: d_109_v75_})
                d_111_v77_: D8
                d_111_v77_ = D8_DC22(d_109_v75_)
                d_110_v76_ = (d_110_v76_).set(d_108_v74_, ((d_111_v77_).cf34).intersection(d_109_v75_))
                d_112_v78_: _dafny.Array
                def lambda6_(d_113_v62_):
                    def lambda7_(d_114_i12_):
                        return d_113_v62_.f7

                    return lambda7_

                init3_ = lambda6_(d_96_v62_)
                nw18_ = _dafny.Array(None, 28)
                for i0_3_ in range(nw18_.length(0)):
                    nw18_[i0_3_] = init3_(i0_3_)
                d_112_v78_ = nw18_
                d_115_v79_: _dafny.Map
                d_115_v79_ = _dafny.Map({d_68_v37_: d_112_v78_})
                d_115_v79_ = ((d_115_v79_) | (d_115_v79_)) | (d_115_v79_)
                d_116_v80_: _dafny.Seq
                d_116_v80_ = _dafny.SeqWithoutIsStrInference([d_112_v78_, d_112_v78_])
                d_116_v80_ = d_116_v80_
            elif True:
                d_117_v81_: _dafny.Map
                d_117_v81_ = _dafny.Map({p0: d_67_v36_})
                d_118_v82_: _dafny.Seq
                d_118_v82_ = _dafny.SeqWithoutIsStrInference([d_117_v81_, default__.fm15(True, d_67_v36_, p0, globalState)])
                d_67_v36_ = (_dafny.SeqWithoutIsStrInference([d_117_v81_])) != ((_dafny.SeqWithoutIsStrInference([d_117_v81_, d_117_v81_, default__.fm15(d_67_v36_, default__.fm3(p0, globalState), 137, globalState), _dafny.Map({(0) - (p0): d_67_v36_})])) + (d_118_v82_))
                d_119_v83_: _dafny.Array
                nw19_ = _dafny.Array(_dafny.CodePoint('D'), 7)
                d_119_v83_ = nw19_
                d_120_v84_: str
                d_120_v84_ = _dafny.CodePoint('f')
                index17_ = default__.safeIndex(974, (d_119_v83_).length(0))
                (d_119_v83_)[index17_] = d_120_v84_
                index18_ = default__.safeIndex(974, (d_119_v83_).length(0))
                (d_119_v83_)[index18_] = d_120_v84_
                d_121_v85_: _dafny.Map
                d_121_v85_ = _dafny.Map({d_67_v36_: False})
                d_121_v85_ = (d_121_v85_).set(True, (d_67_v36_) or (d_67_v36_))
                d_122_v86_: _dafny.Set
                d_122_v86_ = _dafny.Set({p0, -138})
                d_123_v88_: _dafny.MultiSet
                def iife5_():
                    coll5_ = _dafny.Set()
                    compr_5_: int
                    for compr_5_ in (d_122_v86_).Elements:
                        d_124_v87_: int = compr_5_
                        if (d_124_v87_) in (d_122_v86_):
                            coll5_ = coll5_.union(_dafny.Set([default__.safeDivisionInt(d_124_v87_, 948)]))
                    return _dafny.Set(coll5_)
                d_123_v88_ = _dafny.MultiSet([(d_122_v86_).intersection(iife5_()
                )])
                d_125_v90_: _dafny.Set
                d_125_v90_ = _dafny.Set({d_67_v36_})
                d_126_v91_: _dafny.MultiSet
                d_126_v91_ = _dafny.MultiSet([d_125_v90_])
                d_127_v92_: D9
                d_127_v92_ = D9_DC27(d_126_v91_)
                d_128_v93_: _dafny.Seq
                d_128_v93_ = _dafny.SeqWithoutIsStrInference([d_126_v91_, d_126_v91_, (d_127_v92_).cf43, d_126_v91_, d_126_v91_])
                d_129_v94_: _dafny.MultiSet
                d_129_v94_ = _dafny.MultiSet([p0])
                d_130_v95_: _dafny.Map
                d_130_v95_ = _dafny.Map({((d_129_v94_)[123] if (123) in (d_129_v94_) else p0): p0})
                def iife6_():
                    coll6_ = _dafny.Set()
                    compr_6_: int
                    for compr_6_ in _dafny.IntegerRange(113, 926):
                        d_131_v89_: int = compr_6_
                        if ((113) <= (d_131_v89_)) and ((d_131_v89_) < (926)):
                            coll6_ = coll6_.union(_dafny.Set([(d_131_v89_) - (p0)]))
                    return _dafny.Set(coll6_)
                def iife7_():
                    coll7_ = _dafny.Set()
                    compr_7_: int
                    for compr_7_ in _dafny.IntegerRange(113, 926):
                        d_132_v89_: int = compr_7_
                        if ((113) <= (d_132_v89_)) and ((d_132_v89_) < (926)):
                            coll7_ = coll7_.union(_dafny.Set([(d_132_v89_) - (p0)]))
                    return _dafny.Set(coll7_)
                rhs17_ = d_67_v36_
                rhs18_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uownf"))
                rhs19_ = ((d_123_v88_)[(d_122_v86_) | (iife6_()
                )] if ((d_122_v86_) | (iife7_()
                )) in (d_123_v88_) else ((d_128_v93_)[default__.safeIndex(len(d_130_v95_), len(d_128_v93_))]).cardinality)
                lhs15_ = globalState
                d_67_v36_ = rhs17_
                r0 = rhs18_
                lhs15_.f6 = rhs19_
                d_133_v96_: C0
                nw20_ = C0()
                nw20_.ctor__(len(default__.fm16(d_67_v36_, len(d_68_v37_), d_67_v36_, d_67_v36_, globalState)), p0, d_67_v36_)
                d_133_v96_ = nw20_
                d_134_v97_: D6
                d_134_v97_ = D6_DC18(d_133_v96_)
                d_135_v98_: _dafny.Array
                nw21_ = _dafny.Array(None, 6)
                nw21_[int(0)] = d_134_v97_
                nw21_[int(1)] = d_134_v97_
                nw21_[int(2)] = D6_DC18(d_133_v96_)
                nw21_[int(3)] = d_134_v97_
                nw21_[int(4)] = d_134_v97_
                nw21_[int(5)] = d_134_v97_
                d_135_v98_ = nw21_
                index19_ = default__.safeIndex(70, (d_135_v98_).length(0))
                (d_135_v98_)[index19_] = D6_DC18(d_133_v96_)
                pat_let_tv0_ = d_133_v96_
                index20_ = default__.safeIndex(70, (d_135_v98_).length(0))
                def iife8_(_pat_let0_0):
                    def iife9_(d_136_dt__update__tmp_h0_):
                        def iife10_(_pat_let1_0):
                            def iife11_(d_137_dt__update_hcf28_h0_):
                                return D6_DC18(d_137_dt__update_hcf28_h0_)
                            return iife11_(_pat_let1_0)
                        return iife10_(pat_let_tv0_)
                    return iife9_(_pat_let0_0)
                (d_135_v98_)[index20_] = iife8_(d_134_v97_)
            d_138_v99_: T0
            nw22_ = C0()
            nw22_.ctor__((0) - (p0), p0, d_67_v36_)
            d_138_v99_ = nw22_
            d_139_v100_: C0
            nw23_ = C0()
            nw23_.ctor__(default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([d_138_v99_, d_138_v99_])), ((p1)[d_138_v99_.f7] if (d_138_v99_.f7) in (p1) else p0)), p0, d_138_v99_.f7)
            d_139_v100_ = nw23_
            (globalState).f6 = (d_139_v100_).f8
            (d_138_v99_).f7 = d_67_v36_
        elif True:
            (globalState).f3 = p0
            d_67_v36_ = d_67_v36_
            d_140_v101_: T0
            nw24_ = C0()
            nw24_.ctor__((p1).cardinality, p0, default__.fm3(p0, globalState))
            d_140_v101_ = nw24_
            d_141_v102_: _dafny.Set
            d_141_v102_ = _dafny.Set({d_140_v101_})
            if (d_140_v101_) not in (d_141_v102_):
                (globalState).f3 = (p0) - (p0)
                d_142_v103_: _dafny.Array
                def lambda8_(d_143_v101_):
                    def lambda9_(d_144_i13_):
                        return d_143_v101_.f7

                    return lambda9_

                init4_ = lambda8_(d_140_v101_)
                nw25_ = _dafny.Array(None, 28)
                for i0_4_ in range(nw25_.length(0)):
                    nw25_[i0_4_] = init4_(i0_4_)
                d_142_v103_ = nw25_
                d_145_v104_: _dafny.MultiSet
                d_145_v104_ = _dafny.MultiSet([d_142_v103_])
                d_145_v104_ = (_dafny.MultiSet([d_142_v103_])) | (d_145_v104_)
                d_146_v105_: C0
                nw26_ = C0()
                nw26_.ctor__(p0, (0) - (p0), d_67_v36_)
                d_146_v105_ = nw26_
                d_147_v106_: D7
                d_147_v106_ = D7_DC20(True, False, d_146_v105_)
                d_146_v105_ = (d_147_v106_).cf32
                d_148_v107_: C0
                nw27_ = C0()
                nw27_.ctor__(default__.safeDivisionInt(p0, len(_dafny.Map({p0: d_140_v101_}))), default__.safeDivisionInt((d_146_v105_).f9, (d_146_v105_).f8), default__.fm3((d_146_v105_).f8, globalState))
                d_148_v107_ = nw27_
                d_146_v105_ = d_146_v105_
            elif True:
                (d_140_v101_).f7 = (p2)[default__.safeIndex(p0, len(p2))]
                d_149_v108_: _dafny.Array
                nw28_ = _dafny.Array(False, 27)
                d_149_v108_ = nw28_
                index21_ = default__.safeIndex(243, (d_149_v108_).length(0))
                (d_149_v108_)[index21_] = False
                d_150_v109_: _dafny.Map
                d_150_v109_ = _dafny.Map({p0: p1})
                index22_ = default__.safeIndex(243, (d_149_v108_).length(0))
                (d_149_v108_)[index22_] = (p1) != ((((d_150_v109_)[p0] if (p0) in (d_150_v109_) else p1) if not(d_140_v101_.f7) else p1))
                d_151_v110_: _dafny.Map
                d_151_v110_ = _dafny.Map({True: (d_149_v108_)[default__.safeIndex(243, (d_149_v108_).length(0))]})
                d_152_v111_: _dafny.Seq
                d_152_v111_ = _dafny.SeqWithoutIsStrInference([len(d_151_v110_)])
                d_153_v112_: _dafny.Seq
                d_153_v112_ = _dafny.SeqWithoutIsStrInference([len(d_152_v111_)])
                d_154_v113_: D8
                d_154_v113_ = D8_DC22(_dafny.Set({d_152_v111_}))
                d_155_v114_: _dafny.Set
                d_155_v114_ = _dafny.Set({d_153_v112_, d_152_v111_})
                d_154_v113_ = D8_DC22(d_155_v114_)
                d_156_v115_: C0
                nw29_ = C0()
                nw29_.ctor__((d_140_v101_).fm4(globalState), len(_dafny.Map({False: p2})), (d_149_v108_)[default__.safeIndex(243, (d_149_v108_).length(0))])
                d_156_v115_ = nw29_
                d_157_v116_: _dafny.Map
                d_157_v116_ = _dafny.Map({d_156_v115_: p0})
                d_157_v116_ = (d_157_v116_) | (d_157_v116_)
                d_140_v101_ = d_140_v101_
            d_158_v117_: C0
            nw30_ = C0()
            nw30_.ctor__(len(d_68_v37_), p0, d_140_v101_.f7)
            d_158_v117_ = nw30_
            d_159_v118_: D4
            d_159_v118_ = D4_DC14(d_158_v117_, d_140_v101_.f7)
            d_160_v119_: C0
            nw31_ = C0()
            nw31_.ctor__(p0, (p0) - (p0), (d_159_v118_).cf22)
            d_160_v119_ = nw31_
            d_161_v120_: D0
            d_161_v120_ = D0_DC0(d_67_v36_)
            pat_let_tv1_ = d_140_v101_
            def iife12_(_pat_let2_0):
                def iife13_(d_162_dt__update__tmp_h1_):
                    def iife14_(_pat_let3_0):
                        def iife15_(d_163_dt__update_hcf0_h0_):
                            return D0_DC0(d_163_dt__update_hcf0_h0_)
                        return iife15_(_pat_let3_0)
                    return iife14_(pat_let_tv1_.f7)
                return iife13_(_pat_let2_0)
            (d_140_v101_).f7 = (iife12_(d_161_v120_)).cf0
        d_164_v121_: _dafny.Array
        def lambda10_(d_165_i14_):
            return (d_165_i14_) * (-771)

        init5_ = lambda10_
        nw32_ = _dafny.Array(None, 7)
        for i0_5_ in range(nw32_.length(0)):
            nw32_[i0_5_] = init5_(i0_5_)
        d_164_v121_ = nw32_
        d_166_v122_: _dafny.Map
        d_166_v122_ = _dafny.Map({d_67_v36_: d_164_v121_})
        d_166_v122_ = (d_166_v122_).set(default__.fm3(p0, globalState), d_164_v121_)
        d_167_v123_: _dafny.MultiSet
        d_167_v123_ = _dafny.MultiSet([p2])
        d_168_v124_: D0
        d_168_v124_ = D0_DC1(_dafny.MultiSet([p2, p2]), (0) - (p0))
        d_169_v125_: _dafny.MultiSet
        d_169_v125_ = _dafny.MultiSet([d_168_v124_])
        hi2_ = ((_dafny.MultiSet([D0_DC1(d_167_v123_, p0)])) | (d_169_v125_)).cardinality
        for d_170_i15_ in range(p0, hi2_):
            d_171_v126_: _dafny.Seq
            d_171_v126_ = _dafny.SeqWithoutIsStrInference([p0])
            source1_ = D1_DC3((d_171_v126_) + (d_171_v126_))
            if source1_.is_DC4:
                r0 = d_68_v37_
                d_172_v127_: _dafny.Array
                nw33_ = _dafny.Array(None, 27)
                d_172_v127_ = nw33_
                d_173_v128_: C0
                nw34_ = C0()
                nw34_.ctor__(d_170_i15_, (d_171_v126_)[default__.safeIndex(d_170_i15_, len(d_171_v126_))], d_67_v36_)
                d_173_v128_ = nw34_
                d_174_v129_: _dafny.Map
                d_174_v129_ = _dafny.Map({d_173_v128_: d_173_v128_})
                d_175_v130_: _dafny.Map
                d_175_v130_ = _dafny.Map({d_67_v36_: d_174_v129_})
                d_176_v131_: str
                d_176_v131_ = _dafny.CodePoint('q')
                d_177_v132_: _dafny.MultiSet
                d_177_v132_ = _dafny.MultiSet([480, default__.fm0(len(((d_175_v130_)[d_67_v36_] if (d_67_v36_) in (d_175_v130_) else d_174_v129_)), d_67_v36_, _dafny.SeqWithoutIsStrInference([(d_173_v128_).f9 for d_178_i16_ in range(default__.abs(927))]), d_176_v131_, globalState), (0) - (len(d_171_v126_)), len(_dafny.SeqWithoutIsStrInference([d_176_v131_ for d_179_i17_ in range(default__.abs(474))]))])
                d_180_v133_: C0
                nw35_ = C0()
                nw35_.ctor__(p0, (d_177_v132_).cardinality, d_67_v36_)
                d_180_v133_ = nw35_
                index23_ = default__.safeIndex(808, (d_172_v127_).length(0))
                (d_172_v127_)[index23_] = d_173_v128_
                index24_ = default__.safeIndex(808, (d_172_v127_).length(0))
                (d_172_v127_)[index24_] = d_180_v133_
                (globalState).f6 = ((d_173_v128_).f9) * (default__.safeDivisionInt(-448, d_170_i15_))
                (globalState).f3 = (d_180_v133_).f9
            elif source1_.is_DC3:
                d_181___mcc_h0_ = source1_.cf4
                d_182_cf4_ = d_181___mcc_h0_
                d_183_v134_: _dafny.Seq
                d_183_v134_ = d_68_v37_
                d_184_v135_: _dafny.MultiSet
                d_184_v135_ = _dafny.MultiSet([d_182_cf4_, d_182_cf4_])
                d_185_v136_: str
                d_185_v136_ = _dafny.CodePoint('t')
                d_186_v137_: _dafny.Array
                nw36_ = _dafny.Array(None, 29)
                nw36_[int(0)] = d_68_v37_
                nw36_[int(1)] = d_68_v37_
                nw36_[int(2)] = d_68_v37_
                nw36_[int(3)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_187_i18_ in range(default__.abs(-246))])
                nw36_[int(4)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iqglrvh"))
                nw36_[int(5)] = d_68_v37_
                nw36_[int(6)] = (d_68_v37_ if d_67_v36_ else d_68_v37_)
                nw36_[int(7)] = (d_183_v134_)
                nw36_[int(8)] = d_68_v37_
                nw36_[int(9)] = d_68_v37_
                nw36_[int(10)] = (d_68_v37_) + (d_68_v37_)
                nw36_[int(11)] = (d_68_v37_) + (d_68_v37_)
                nw36_[int(12)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ugemjeiny"))
                nw36_[int(13)] = d_68_v37_
                nw36_[int(14)] = d_68_v37_
                nw36_[int(15)] = d_68_v37_
                nw36_[int(16)] = d_68_v37_
                nw36_[int(17)] = d_68_v37_
                nw36_[int(18)] = d_68_v37_
                nw36_[int(19)] = (default__.fm10(p0, d_67_v36_, d_170_i15_, d_184_v135_, globalState)) + (d_68_v37_)
                nw36_[int(20)] = d_68_v37_
                nw36_[int(21)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "foi"))) + (d_68_v37_)
                nw36_[int(22)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ptuxtb"))).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ptuxtb")))), d_185_v136_)
                nw36_[int(23)] = d_68_v37_
                nw36_[int(24)] = d_68_v37_
                nw36_[int(25)] = d_68_v37_
                nw36_[int(26)] = d_68_v37_
                nw36_[int(27)] = d_68_v37_
                nw36_[int(28)] = d_68_v37_
                d_186_v137_ = nw36_
                index25_ = default__.safeIndex(936, (d_186_v137_).length(0))
                (d_186_v137_)[index25_] = d_68_v37_
                index26_ = default__.safeIndex(936, (d_186_v137_).length(0))
                (d_186_v137_)[index26_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jquu"))
                index27_ = default__.safeIndex(933, (d_164_v121_).length(0))
                (d_164_v121_)[index27_] = p0
                index28_ = default__.safeIndex(933, (d_164_v121_).length(0))
                (d_164_v121_)[index28_] = d_170_i15_
                d_188_v138_: _dafny.Map
                d_188_v138_ = _dafny.Map({p2: (d_164_v121_)[default__.safeIndex(933, (d_164_v121_).length(0))]})
                d_189_v139_: _dafny.Map
                d_189_v139_ = _dafny.Map({d_188_v138_: d_185_v136_})
                def iife16_():
                    coll8_ = _dafny.Map()
                    compr_8_: _dafny.Seq
                    for compr_8_ in (default__.fm17(d_67_v36_, d_67_v36_, d_67_v36_, globalState)).Elements:
                        d_190_v140_: _dafny.Seq = compr_8_
                        if (d_190_v140_) in (default__.fm17(d_67_v36_, d_67_v36_, d_67_v36_, globalState)):
                            coll8_[d_190_v140_] = d_170_i15_
                    return _dafny.Map(coll8_)
                def iife17_():
                    coll9_ = _dafny.Map()
                    compr_9_: _dafny.Seq
                    for compr_9_ in (default__.fm17(d_67_v36_, d_67_v36_, d_67_v36_, globalState)).Elements:
                        d_191_v140_: _dafny.Seq = compr_9_
                        if (d_191_v140_) in (default__.fm17(d_67_v36_, d_67_v36_, d_67_v36_, globalState)):
                            coll9_[d_191_v140_] = d_170_i15_
                    return _dafny.Map(coll9_)
                d_185_v136_ = ((d_189_v139_)[iife16_()
                ] if (iife17_()
                ) in (d_189_v139_) else _dafny.CodePoint('g'))
                d_192_v141_: C0
                nw37_ = C0()
                nw37_.ctor__(len(p2), d_170_i15_, False)
                d_192_v141_ = nw37_
            elif True:
                d_193___mcc_h1_ = source1_.cf5
                d_194_cf5_ = d_193___mcc_h1_
                d_67_v36_ = d_67_v36_
                d_67_v36_ = (d_68_v37_) not in (_dafny.MultiSet([d_68_v37_, d_68_v37_, d_68_v37_, d_68_v37_, d_68_v37_]))
                d_195_v142_: _dafny.MultiSet
                d_195_v142_ = _dafny.MultiSet([d_67_v36_])
                d_195_v142_ = p1
                d_196_v143_: _dafny.Map
                d_196_v143_ = _dafny.Map({d_67_v36_: d_170_i15_})
                d_195_v142_ = ((default__.fm12(d_170_i15_, d_170_i15_, d_196_v143_, globalState)).set(d_67_v36_, default__.abs(d_170_i15_))) | (p1)
            d_197_v144_: _dafny.Set
            d_197_v144_ = _dafny.Set({d_170_i15_})
            d_198_v145_: _dafny.MultiSet
            d_198_v145_ = _dafny.MultiSet([d_197_v144_])
            d_199_v146_: T0
            nw38_ = C0()
            nw38_.ctor__((((d_198_v145_)[d_197_v144_] if (d_197_v144_) in (d_198_v145_) else p0)) + (d_170_i15_), (0) - (d_170_i15_), d_67_v36_)
            d_199_v146_ = nw38_
            d_200_v147_: _dafny.Map
            d_200_v147_ = _dafny.Map({d_170_i15_: d_199_v146_})
            d_199_v146_ = ((d_200_v147_)[d_170_i15_] if (d_170_i15_) in (d_200_v147_) else d_199_v146_)
            d_201_v148_: _dafny.Array
            nw39_ = _dafny.Array(_dafny.Array(None, 0), 26)
            d_201_v148_ = nw39_
            index29_ = default__.safeIndex(300, (d_201_v148_).length(0))
            (d_201_v148_)[index29_] = d_164_v121_
            index30_ = default__.safeIndex(300, (d_201_v148_).length(0))
            nw40_ = _dafny.Array(int(0), 13)
            (d_201_v148_)[index30_] = nw40_
            (d_199_v146_).f7 = d_199_v146_.f7
        r0 = ((d_68_v37_) + (d_68_v37_)) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oc"))) + (d_68_v37_))
        return r0

    @staticmethod
    def Main(noArgsParameter__):
        d_202_v0_: _dafny.Seq
        d_202_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qfxbbb"))
        d_203_v1_: int
        d_203_v1_ = 779
        d_204_v2_: str
        d_204_v2_ = _dafny.CodePoint('p')
        d_205_v3_: _dafny.Map
        d_205_v3_ = _dafny.Map({d_203_v1_: d_204_v2_})
        d_206_v4_: bool
        d_206_v4_ = False
        d_207_v5_: _dafny.Map
        d_207_v5_ = _dafny.Map({((d_205_v3_)[len(d_202_v0_)] if (len(d_202_v0_)) in (d_205_v3_) else d_204_v2_): d_206_v4_})
        d_208_v6_: _dafny.MultiSet
        d_208_v6_ = _dafny.MultiSet([-166])
        d_209_v7_: _dafny.MultiSet
        d_209_v7_ = _dafny.MultiSet([_dafny.MultiSet([(d_208_v6_).cardinality])])
        d_210_globalState_: GlobalState
        nw41_ = GlobalState()
        nw41_.ctor__(True, d_202_v0_, True, 225, (d_207_v5_) | (d_207_v5_), d_209_v7_, 852)
        d_210_globalState_ = nw41_
        (d_210_globalState_).f6 = 155
        def iife18_():
            coll10_ = _dafny.Map()
            compr_10_: str
            for compr_10_ in (d_202_v0_).Elements:
                d_244_v8_: str = compr_10_
                if (d_244_v8_) in (d_202_v0_):
                    coll10_[d_244_v8_] = d_203_v1_
            return _dafny.Map(coll10_)
        hi3_ = d_203_v1_
        for d_211_i0_ in range((len((_dafny.SeqWithoutIsStrInference([d_203_v1_])).set(default__.safeIndex(len(iife18_()
        ), len(_dafny.SeqWithoutIsStrInference([d_203_v1_]))), d_203_v1_))) + (83), hi3_):
            d_212_v9_: _dafny.Seq
            d_212_v9_ = _dafny.SeqWithoutIsStrInference([d_211_i0_, 832, d_203_v1_, d_203_v1_])
            if (d_212_v9_) < (d_212_v9_):
                d_213_v10_: _dafny.Array
                def lambda11_(d_214_v6_):
                    def lambda12_(d_215_i1_):
                        return d_214_v6_

                    return lambda12_

                init6_ = lambda11_(d_208_v6_)
                nw42_ = _dafny.Array(None, 8)
                for i0_6_ in range(nw42_.length(0)):
                    nw42_[i0_6_] = init6_(i0_6_)
                d_213_v10_ = nw42_
                d_213_v10_ = d_213_v10_
                d_216_v11_: _dafny.MultiSet
                d_216_v11_ = _dafny.MultiSet([d_206_v4_, not(d_206_v4_)])
                d_217_v12_: _dafny.Seq
                out0_: _dafny.Seq
                out0_ = default__.m0(d_211_i0_, (d_216_v11_).intersection(d_216_v11_), (_dafny.SeqWithoutIsStrInference([d_206_v4_])).set(default__.safeIndex(len(d_212_v9_), len(_dafny.SeqWithoutIsStrInference([d_206_v4_]))), d_206_v4_), d_210_globalState_)
                d_217_v12_ = out0_
                rhs20_ = (d_211_i0_ if d_206_v4_ else len(d_202_v0_))
                rhs21_ = d_211_i0_
                lhs16_ = d_210_globalState_
                lhs17_ = d_210_globalState_
                lhs16_.f6 = rhs20_
                lhs17_.f3 = rhs21_
                d_206_v4_ = d_206_v4_
                d_218_v13_: D0
                d_218_v13_ = D0_DC0(d_206_v4_)
                d_206_v4_ = (d_218_v13_).cf0
            elif True:
                d_219_v14_: _dafny.Set
                d_219_v14_ = _dafny.Set({d_211_i0_, d_203_v1_})
                rhs22_ = default__.fm0((default__.fm1(d_210_globalState_)).cf2, True, d_212_v9_, d_204_v2_, d_210_globalState_)
                rhs23_ = default__.safeDivisionInt(d_203_v1_, len(d_219_v14_))
                lhs18_ = d_210_globalState_
                lhs19_ = d_210_globalState_
                lhs18_.f3 = rhs22_
                lhs19_.f3 = rhs23_
                d_220_v15_: _dafny.Map
                d_220_v15_ = _dafny.Map({d_206_v4_: _dafny.MultiSet([928, d_203_v1_])})
                d_221_v16_: _dafny.Map
                d_221_v16_ = _dafny.Map({((d_220_v15_)[d_206_v4_] if (d_206_v4_) in (d_220_v15_) else d_208_v6_): False})
                d_219_v14_ = default__.fm2(d_221_v16_, _dafny.CodePoint('d'), d_210_globalState_)
                d_222_v17_: _dafny.Map
                d_222_v17_ = _dafny.Map({d_211_i0_: d_203_v1_})
                d_223_v18_: _dafny.Seq
                d_223_v18_ = _dafny.SeqWithoutIsStrInference([d_206_v4_, d_206_v4_])
                d_224_v19_: _dafny.MultiSet
                d_224_v19_ = _dafny.MultiSet([d_223_v18_])
                d_225_v20_: _dafny.Map
                d_225_v20_ = _dafny.Map({d_206_v4_: False})
                d_226_v21_: D0
                d_226_v21_ = D0_DC1(d_224_v19_, len(d_225_v20_))
                d_227_v22_: _dafny.Map
                d_227_v22_ = _dafny.Map({d_202_v0_: d_206_v4_})
                d_228_v23_: _dafny.Map
                d_228_v23_ = _dafny.Map({d_227_v22_: d_206_v4_})
                d_229_v24_: _dafny.Set
                d_229_v24_ = _dafny.Set({d_206_v4_})
                rhs24_ = ((d_203_v1_ if d_206_v4_ else d_211_i0_) if default__.fm3(d_203_v1_, d_210_globalState_) else (d_226_v21_).cf2)
                rhs25_ = d_222_v17_
                rhs26_ = ((d_228_v23_)[d_227_v22_] if (d_227_v22_) in (d_228_v23_) else (d_229_v24_).ispropersubset(d_229_v24_))
                rhs27_ = d_206_v4_
                lhs20_ = d_210_globalState_
                lhs20_.f3 = rhs24_
                d_222_v17_ = rhs25_
                d_206_v4_ = rhs26_
                d_206_v4_ = rhs27_
                d_206_v4_ = d_206_v4_
                d_230_v25_: _dafny.Array
                nw43_ = _dafny.Array(int(0), 14)
                d_230_v25_ = nw43_
                index31_ = default__.safeIndex(794, (d_230_v25_).length(0))
                (d_230_v25_)[index31_] = d_211_i0_
                index32_ = default__.safeIndex(794, (d_230_v25_).length(0))
                (d_230_v25_)[index32_] = d_203_v1_
            d_231_v26_: _dafny.Seq
            d_231_v26_ = _dafny.SeqWithoutIsStrInference([d_206_v4_])
            d_232_v27_: D0
            d_232_v27_ = D0_DC0(False)
            d_233_v28_: _dafny.Array
            nw44_ = _dafny.Array(None, 10)
            nw44_[int(0)] = d_206_v4_
            nw44_[int(1)] = (d_231_v26_)[default__.safeIndex(464, len(d_231_v26_))]
            nw44_[int(2)] = d_206_v4_
            nw44_[int(3)] = d_206_v4_
            nw44_[int(4)] = d_206_v4_
            nw44_[int(5)] = d_206_v4_
            nw44_[int(6)] = True
            nw44_[int(7)] = d_206_v4_
            nw44_[int(8)] = d_206_v4_
            nw44_[int(9)] = (d_232_v27_).cf0
            d_233_v28_ = nw44_
            index33_ = default__.safeIndex(685, (d_233_v28_).length(0))
            (d_233_v28_)[index33_] = d_206_v4_
            index34_ = default__.safeIndex(685, (d_233_v28_).length(0))
            (d_233_v28_)[index34_] = d_206_v4_
            d_234_v29_: _dafny.Set
            d_234_v29_ = _dafny.Set({d_204_v2_})
            d_235_v30_: _dafny.MultiSet
            d_235_v30_ = _dafny.MultiSet([(d_234_v29_).intersection(d_234_v29_), d_234_v29_])
            d_236_v31_: _dafny.Seq
            d_236_v31_ = _dafny.SeqWithoutIsStrInference([d_234_v29_])
            d_235_v30_ = ((_dafny.MultiSet(d_236_v31_)) | (d_235_v30_)).set(d_234_v29_, default__.abs(d_211_i0_))
            d_237_v32_: _dafny.Map
            d_237_v32_ = _dafny.Map({False: d_211_i0_})
            d_238_v33_: _dafny.Seq
            d_238_v33_ = _dafny.SeqWithoutIsStrInference([d_233_v28_, d_233_v28_])
            d_239_v34_: _dafny.Seq
            d_239_v34_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([(0) - (d_203_v1_)]), d_208_v6_])
            d_240_v35_: D1
            d_240_v35_ = D1_DC3(d_212_v9_)
            d_241_v36_: _dafny.Array
            nw45_ = _dafny.Array(None, 21)
            nw45_[int(0)] = len((_dafny.SeqWithoutIsStrInference([(0) - ((_dafny.MultiSet([False])).cardinality) for d_242_i2_ in range(default__.abs(-332))])).set(default__.safeIndex(len(d_202_v0_), len(_dafny.SeqWithoutIsStrInference([(0) - ((_dafny.MultiSet([False])).cardinality) for d_243_i2_ in range(default__.abs(-332))]))), len(d_237_v32_)))
            nw45_[int(1)] = 742
            nw45_[int(2)] = len(d_237_v32_)
            nw45_[int(3)] = d_203_v1_
            nw45_[int(4)] = d_211_i0_
            nw45_[int(5)] = d_203_v1_
            nw45_[int(6)] = d_211_i0_
            nw45_[int(7)] = 967
            nw45_[int(8)] = d_211_i0_
            nw45_[int(9)] = 568
            nw45_[int(10)] = len(d_238_v33_)
            nw45_[int(11)] = d_203_v1_
            nw45_[int(12)] = 869
            nw45_[int(13)] = default__.safeModuloInt(d_203_v1_, d_211_i0_)
            nw45_[int(14)] = d_203_v1_
            nw45_[int(15)] = (d_211_i0_) - (d_203_v1_)
            nw45_[int(16)] = d_203_v1_
            nw45_[int(17)] = default__.safeDivisionInt(default__.fm0(len(d_239_v34_), (d_233_v28_)[default__.safeIndex(685, (d_233_v28_).length(0))], d_212_v9_, d_204_v2_, d_210_globalState_), default__.fm0(len(d_202_v0_), (d_233_v28_)[default__.safeIndex(685, (d_233_v28_).length(0))], (d_240_v35_).cf4, d_204_v2_, d_210_globalState_))
            nw45_[int(18)] = (0) - (d_211_i0_)
            nw45_[int(19)] = 975
            nw45_[int(20)] = ((0) - (d_211_i0_)) + (d_211_i0_)
            d_241_v36_ = nw45_
            index35_ = default__.safeIndex(938, (d_241_v36_).length(0))
            (d_241_v36_)[index35_] = d_203_v1_
            index36_ = default__.safeIndex(938, (d_241_v36_).length(0))
            rhs28_ = d_211_i0_
            rhs29_ = d_211_i0_
            rhs30_ = d_211_i0_
            lhs21_ = d_241_v36_
            lhs22_ = default__.safeIndex(938, (d_241_v36_).length(0))
            lhs23_ = d_210_globalState_
            lhs21_[lhs22_] = rhs28_
            d_203_v1_ = rhs29_
            lhs23_.f6 = rhs30_
        d_245_v37_: _dafny.Set
        d_245_v37_ = _dafny.Set({d_204_v2_})
        d_246_v38_: _dafny.Set
        d_246_v38_ = _dafny.Set({d_245_v37_, d_245_v37_, d_245_v37_, _dafny.Set({d_204_v2_, d_204_v2_}), d_245_v37_})
        d_247_v39_: _dafny.Seq
        d_247_v39_ = _dafny.SeqWithoutIsStrInference([d_246_v38_])
        d_248_v40_: _dafny.Map
        d_248_v40_ = _dafny.Map({d_203_v1_: d_203_v1_})
        hi4_ = (d_203_v1_) - (len(d_248_v40_))
        for d_249_i3_ in range(len((d_247_v39_)[default__.safeIndex(d_203_v1_, len(d_247_v39_))]), hi4_):
            if d_206_v4_:
                d_250_v41_: _dafny.Array
                nw46_ = _dafny.Array(_dafny.Seq({}), 1)
                d_250_v41_ = nw46_
                d_251_v42_: _dafny.Array
                def lambda13_(d_252_v1_):
                    def lambda14_(d_253_i4_):
                        return default__.safeModuloInt(d_253_i4_, d_252_v1_)

                    return lambda14_

                init7_ = lambda13_(d_203_v1_)
                nw47_ = _dafny.Array(None, 18)
                for i0_7_ in range(nw47_.length(0)):
                    nw47_[i0_7_] = init7_(i0_7_)
                d_251_v42_ = nw47_
                d_254_v43_: _dafny.Map
                d_254_v43_ = _dafny.Map({d_250_v41_: d_251_v42_})
                d_254_v43_ = (d_254_v43_).set(d_250_v41_, d_251_v42_)
                d_255_v44_: C0
                nw48_ = C0()
                nw48_.ctor__(d_203_v1_, default__.safeModuloInt(d_249_i3_, 163), d_206_v4_)
                d_255_v44_ = nw48_
                index37_ = default__.safeIndex(326, (d_251_v42_).length(0))
                (d_251_v42_)[index37_] = default__.safeDivisionInt((d_255_v44_).f8, d_249_i3_)
                index38_ = default__.safeIndex(326, (d_251_v42_).length(0))
                (d_251_v42_)[index38_] = (0) - (d_203_v1_)
                d_256_v45_: _dafny.Set
                d_256_v45_ = _dafny.Set({d_206_v4_})
                d_257_v46_: _dafny.MultiSet
                d_257_v46_ = _dafny.MultiSet([default__.fm3(d_203_v1_, d_210_globalState_)])
                d_258_v47_: D1
                d_258_v47_ = D1_DC4()
                d_259_v48_: _dafny.Map
                d_259_v48_ = _dafny.Map({d_258_v47_: d_206_v4_})
                d_260_v49_: _dafny.Seq
                out1_: _dafny.Seq
                out1_ = default__.m0((len(d_256_v45_)) + (d_249_i3_), d_257_v46_, (default__.fm6((d_255_v44_).f9, _dafny.CodePoint('p'), ((d_259_v48_)[d_258_v47_] if (d_258_v47_) in (d_259_v48_) else d_206_v4_), d_206_v4_, d_210_globalState_)) + (default__.fm6(d_203_v1_, d_204_v2_, d_206_v4_, d_206_v4_, d_210_globalState_)), d_210_globalState_)
                d_260_v49_ = out1_
                d_261_v50_: _dafny.Seq
                d_261_v50_ = _dafny.SeqWithoutIsStrInference([d_206_v4_, d_206_v4_])
                d_262_v51_: _dafny.Seq
                d_262_v51_ = _dafny.SeqWithoutIsStrInference([(d_261_v50_)[default__.safeIndex((d_255_v44_).f8, len(d_261_v50_))]])
                d_263_v52_: _dafny.Seq
                out2_: _dafny.Seq
                out2_ = default__.m0(496, d_257_v46_, d_262_v51_, d_210_globalState_)
                d_263_v52_ = out2_
            elif True:
                d_206_v4_ = (d_208_v6_).isdisjoint(d_208_v6_)
                d_264_v53_: _dafny.Array
                def lambda15_(d_265_v4_):
                    def lambda16_(d_266_i5_):
                        return d_265_v4_

                    return lambda16_

                init8_ = lambda15_(d_206_v4_)
                nw49_ = _dafny.Array(None, 26)
                for i0_8_ in range(nw49_.length(0)):
                    nw49_[i0_8_] = init8_(i0_8_)
                d_264_v53_ = nw49_
                d_267_v54_: _dafny.Map
                d_267_v54_ = _dafny.Map({d_249_i3_: d_264_v53_})
                d_203_v1_ = len(((d_267_v54_) | (_dafny.Map({d_203_v1_: d_264_v53_}))) | (d_267_v54_))
                d_268_v55_: _dafny.Map
                d_268_v55_ = _dafny.Map({d_206_v4_: d_249_i3_})
                d_269_v56_: _dafny.Seq
                d_269_v56_ = _dafny.SeqWithoutIsStrInference([not(True)])
                d_270_v57_: _dafny.Seq
                d_270_v57_ = _dafny.SeqWithoutIsStrInference([d_203_v1_, d_203_v1_])
                d_271_v58_: _dafny.Map
                d_271_v58_ = _dafny.Map({default__.fm0(len(d_268_v55_), (d_269_v56_)[default__.safeIndex((0) - (len(d_202_v0_)), len(d_269_v56_))], d_270_v57_, d_204_v2_, d_210_globalState_): d_206_v4_})
                d_272_v59_: C0
                nw50_ = C0()
                nw50_.ctor__((969) - (d_249_i3_), default__.safeDivisionInt(d_203_v1_, len(d_268_v55_)), ((d_271_v58_)[d_203_v1_] if (d_203_v1_) in (d_271_v58_) else d_206_v4_))
                d_272_v59_ = nw50_
                d_273_v60_: _dafny.Array
                nw51_ = _dafny.Array(D1.default()(), 12)
                d_273_v60_ = nw51_
                d_274_v61_: D1
                d_274_v61_ = D1_DC3(_dafny.SeqWithoutIsStrInference([146]))
                index39_ = default__.safeIndex(998, (d_273_v60_).length(0))
                (d_273_v60_)[index39_] = d_274_v61_
                index40_ = default__.safeIndex(998, (d_273_v60_).length(0))
                (d_273_v60_)[index40_] = d_274_v61_
                d_275_v62_: _dafny.MultiSet
                d_275_v62_ = _dafny.MultiSet([d_206_v4_])
                d_276_v63_: _dafny.Seq
                out3_: _dafny.Seq
                out3_ = default__.m0(((d_272_v59_).f9 if d_206_v4_ else len(d_269_v56_)), (_dafny.MultiSet(d_269_v56_)).intersection(d_275_v62_), (d_269_v56_).set(default__.safeIndex(d_249_i3_, len(d_269_v56_)), d_206_v4_), d_210_globalState_)
                d_276_v63_ = out3_
            (d_210_globalState_).f6 = default__.safeDivisionInt(default__.safeDivisionInt((0) - (d_249_i3_), len(d_202_v0_)), d_203_v1_)
            d_277_v64_: _dafny.Seq
            d_277_v64_ = _dafny.SeqWithoutIsStrInference([d_203_v1_])
            d_278_v65_: _dafny.Array
            nw52_ = _dafny.Array(None, 11)
            nw52_[int(0)] = True
            nw52_[int(1)] = d_206_v4_
            nw52_[int(2)] = d_206_v4_
            nw52_[int(3)] = False
            nw52_[int(4)] = not(d_206_v4_)
            nw52_[int(5)] = d_206_v4_
            nw52_[int(6)] = not(d_206_v4_)
            nw52_[int(7)] = d_206_v4_
            nw52_[int(8)] = default__.fm3((0) - (d_249_i3_), d_210_globalState_)
            nw52_[int(9)] = not(False)
            nw52_[int(10)] = d_206_v4_
            d_278_v65_ = nw52_
            d_279_v66_: _dafny.Map
            d_279_v66_ = _dafny.Map({(d_277_v64_)[default__.safeIndex(d_249_i3_, len(d_277_v64_))]: d_278_v65_})
            d_280_v67_: _dafny.Set
            d_280_v67_ = _dafny.Set({d_206_v4_})
            d_279_v66_ = (d_279_v66_).set(len(d_280_v67_), d_278_v65_)
            if d_206_v4_:
                d_281_v68_: _dafny.MultiSet
                d_281_v68_ = _dafny.MultiSet([d_206_v4_])
                rhs31_ = (d_203_v1_) * (len(_dafny.Set({(0) - (d_249_i3_), d_203_v1_, d_203_v1_})))
                rhs32_ = (d_281_v68_) != (_dafny.MultiSet([d_206_v4_]))
                lhs24_ = d_210_globalState_
                lhs24_.f3 = rhs31_
                d_206_v4_ = rhs32_
                d_282_v69_: C0
                nw53_ = C0()
                nw53_.ctor__(d_249_i3_, d_249_i3_, (d_206_v4_) == (d_206_v4_))
                d_282_v69_ = nw53_
                d_206_v4_ = not((D0_DC0(d_206_v4_)).cf0)
                d_283_v70_: _dafny.Seq
                d_283_v70_ = _dafny.SeqWithoutIsStrInference([d_281_v68_])
                d_284_v71_: _dafny.Seq
                d_284_v71_ = _dafny.SeqWithoutIsStrInference([d_206_v4_])
                d_285_v72_: _dafny.Seq
                out4_: _dafny.Seq
                out4_ = default__.m0(default__.safeDivisionInt(d_249_i3_, default__.fm0(d_203_v1_, d_206_v4_, d_277_v64_, d_204_v2_, d_210_globalState_)), ((d_281_v68_).set(d_206_v4_, default__.abs((0) - (d_249_i3_)))) - ((d_283_v70_)[default__.safeIndex((d_282_v69_).f9, len(d_283_v70_))]), d_284_v71_, d_210_globalState_)
                d_285_v72_ = out4_
                d_206_v4_ = default__.fm3((d_282_v69_).f9, d_210_globalState_)
            elif True:
                d_206_v4_ = d_206_v4_
                d_286_v73_: _dafny.Seq
                d_286_v73_ = _dafny.SeqWithoutIsStrInference([d_206_v4_])
                d_287_v74_: _dafny.Seq
                d_287_v74_ = _dafny.SeqWithoutIsStrInference([d_206_v4_, (d_286_v73_)[default__.safeIndex((d_208_v6_).cardinality, len(d_286_v73_))]])
                index41_ = default__.safeIndex(891, (d_278_v65_).length(0))
                (d_278_v65_)[index41_] = (d_286_v73_)[default__.safeIndex(d_203_v1_, len(d_286_v73_))]
                index42_ = default__.safeIndex(891, (d_278_v65_).length(0))
                (d_278_v65_)[index42_] = d_206_v4_
                d_288_v75_: _dafny.MultiSet
                d_288_v75_ = _dafny.MultiSet([d_286_v73_])
                d_289_v76_: D0
                d_289_v76_ = D0_DC1(d_288_v75_, d_203_v1_)
                d_290_v77_: D0
                d_290_v77_ = D0_DC2(d_289_v76_)
                pat_let_tv2_ = d_210_globalState_
                def iife19_(_pat_let4_0):
                    def iife20_(d_291_dt__update__tmp_h0_):
                        def iife21_(_pat_let5_0):
                            def iife22_(d_292_dt__update_hcf3_h0_):
                                return D0_DC2(d_292_dt__update_hcf3_h0_)
                            return iife22_(_pat_let5_0)
                        return iife21_(default__.fm1(pat_let_tv2_))
                    return iife20_(_pat_let4_0)
                d_290_v77_ = iife19_((d_290_v77_ if d_206_v4_ else d_290_v77_))
                index43_ = default__.safeIndex(891, (d_278_v65_).length(0))
                (d_278_v65_)[index43_] = default__.fm3((d_249_i3_ if d_206_v4_ else d_249_i3_), d_210_globalState_)
                (d_210_globalState_).f3 = d_203_v1_
        hi5_ = d_203_v1_
        for d_293_i6_ in range(d_203_v1_, hi5_):
            d_294_v78_: C0
            nw54_ = C0()
            nw54_.ctor__((0) - ((0) - (d_203_v1_)), d_293_i6_, d_206_v4_)
            d_294_v78_ = nw54_
            d_205_v3_ = (d_205_v3_).set((d_294_v78_).f9, d_204_v2_)
            d_295_v79_: _dafny.Seq
            d_295_v79_ = _dafny.SeqWithoutIsStrInference([d_206_v4_, d_206_v4_, d_206_v4_])
            d_296_v80_: _dafny.Seq
            out5_: _dafny.Seq
            out5_ = default__.m0((0) - ((d_294_v78_).f8), _dafny.MultiSet([d_206_v4_, False]), d_295_v79_, d_210_globalState_)
            d_296_v80_ = out5_
            (d_210_globalState_).f3 = (0) - (d_293_i6_)
        d_297_v81_: _dafny.Seq
        d_297_v81_ = _dafny.SeqWithoutIsStrInference([d_203_v1_])
        d_298_v82_: _dafny.Seq
        d_298_v82_ = _dafny.SeqWithoutIsStrInference([d_203_v1_, (d_297_v81_)[default__.safeIndex(d_203_v1_, len(d_297_v81_))]])
        d_299_v83_: D1
        d_299_v83_ = D1_DC3(d_297_v81_)
        d_300_i7_: int
        d_300_i7_ = 0
        with _dafny.label("0"):
            pat_let_tv3_ = d_206_v4_
            pat_let_tv4_ = d_206_v4_
            pat_let_tv5_ = d_203_v1_
            def lambda17_(source2_):
                if source2_.is_DC4:
                    return pat_let_tv3_
                elif source2_.is_DC3:
                    d_313___mcc_h0_ = source2_.cf4
                    d_314_cf4_ = d_313___mcc_h0_
                    return pat_let_tv4_
                elif True:
                    d_315___mcc_h1_ = source2_.cf5
                    d_316_cf5_ = d_315___mcc_h1_
                    return (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nol")))) > (pat_let_tv5_)

            while lambda17_(d_299_v83_):
                with _dafny.c_label("0"):
                    if (d_300_i7_) >= (100):
                        raise _dafny.Break("0")
                    d_300_i7_ = (d_300_i7_) + (1)
                    d_301_v84_: _dafny.Array
                    nw55_ = _dafny.Array(False, 14)
                    d_301_v84_ = nw55_
                    d_302_v85_: _dafny.Seq
                    d_302_v85_ = _dafny.SeqWithoutIsStrInference([d_206_v4_])
                    index44_ = default__.safeIndex(798, (d_301_v84_).length(0))
                    (d_301_v84_)[index44_] = (d_302_v85_) != (d_302_v85_)
                    d_303_v86_: D0
                    d_303_v86_ = D0_DC0(d_206_v4_)
                    d_304_v87_: C0
                    nw56_ = C0()
                    nw56_.ctor__(d_203_v1_, 715, (d_303_v86_).cf0)
                    d_304_v87_ = nw56_
                    d_305_v88_: _dafny.Map
                    d_305_v88_ = _dafny.Map({d_304_v87_: d_206_v4_})
                    d_306_v89_: _dafny.MultiSet
                    d_306_v89_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([not(d_206_v4_), d_206_v4_]), d_302_v85_])
                    index45_ = default__.safeIndex(798, (d_301_v84_).length(0))
                    (d_301_v84_)[index45_] = ((d_305_v88_)[d_304_v87_] if (d_304_v87_) in (d_305_v88_) else (d_306_v89_).issubset(d_306_v89_))
                    d_307_v90_: _dafny.Array
                    nw57_ = _dafny.Array(int(0), 16)
                    d_307_v90_ = nw57_
                    index46_ = default__.safeIndex(384, (d_307_v90_).length(0))
                    (d_307_v90_)[index46_] = (d_304_v87_).f9
                    d_308_v91_: _dafny.Map
                    d_308_v91_ = _dafny.Map({d_206_v4_: (d_301_v84_)[default__.safeIndex(798, (d_301_v84_).length(0))]})
                    d_309_v92_: _dafny.Set
                    d_309_v92_ = _dafny.Set({d_206_v4_, (d_303_v86_).cf0, (d_308_v91_) != (d_308_v91_), default__.fm3((d_304_v87_).f9, d_210_globalState_), ((d_301_v84_)[default__.safeIndex(798, (d_301_v84_).length(0))]) in (d_302_v85_)})
                    index47_ = default__.safeIndex(798, (d_301_v84_).length(0))
                    index48_ = default__.safeIndex(384, (d_307_v90_).length(0))
                    rhs33_ = d_206_v4_
                    rhs34_ = (d_304_v87_).f9
                    rhs35_ = len(d_309_v92_)
                    lhs25_ = d_301_v84_
                    lhs26_ = default__.safeIndex(798, (d_301_v84_).length(0))
                    lhs27_ = d_210_globalState_
                    lhs28_ = d_307_v90_
                    lhs29_ = default__.safeIndex(384, (d_307_v90_).length(0))
                    lhs25_[lhs26_] = rhs33_
                    lhs27_.f6 = rhs34_
                    lhs28_[lhs29_] = rhs35_
                    d_310_v93_: _dafny.Array
                    nw58_ = _dafny.Array(None, 11)
                    d_310_v93_ = nw58_
                    nw59_ = _dafny.Array(None, 27)
                    d_310_v93_ = nw59_
                    d_311_v94_: _dafny.MultiSet
                    d_311_v94_ = _dafny.MultiSet([not(False), not(not(not(d_206_v4_))), d_206_v4_, not(True), (d_301_v84_)[default__.safeIndex(798, (d_301_v84_).length(0))]])
                    d_312_v95_: _dafny.Seq
                    out6_: _dafny.Seq
                    out6_ = default__.m0((0) - (d_203_v1_), d_311_v94_, d_302_v85_, d_210_globalState_)
                    d_312_v95_ = out6_
                    pass
            pass
        d_206_v4_ = d_206_v4_
        d_317_v96_: D0
        d_317_v96_ = D0_DC0(default__.fm3(d_203_v1_, d_210_globalState_))
        if (d_317_v96_).cf0:
            d_318_v97_: _dafny.Array
            nw60_ = _dafny.Array(False, 18)
            d_318_v97_ = nw60_
            index49_ = default__.safeIndex(29, (d_318_v97_).length(0))
            (d_318_v97_)[index49_] = d_206_v4_
            d_319_v98_: _dafny.Map
            d_319_v98_ = _dafny.Map({d_206_v4_: d_203_v1_})
            d_320_v99_: _dafny.Array
            nw61_ = _dafny.Array(None, 4)
            nw61_[int(0)] = (0) - (d_203_v1_)
            nw61_[int(1)] = d_203_v1_
            nw61_[int(2)] = ((d_319_v98_)[d_206_v4_] if (d_206_v4_) in (d_319_v98_) else len(_dafny.Map({d_318_v97_: d_203_v1_})))
            nw61_[int(3)] = d_203_v1_
            d_320_v99_ = nw61_
            index50_ = default__.safeIndex(938, (d_320_v99_).length(0))
            (d_320_v99_)[index50_] = 736
            d_321_v100_: _dafny.Seq
            d_321_v100_ = _dafny.SeqWithoutIsStrInference([d_206_v4_, d_206_v4_])
            d_322_v101_: _dafny.MultiSet
            d_322_v101_ = _dafny.MultiSet([d_321_v100_])
            d_323_v102_: D0
            d_323_v102_ = D0_DC1(d_322_v101_, d_203_v1_)
            d_324_v103_: _dafny.Map
            d_324_v103_ = _dafny.Map({d_298_v82_: d_203_v1_})
            index51_ = default__.safeIndex(29, (d_318_v97_).length(0))
            index52_ = default__.safeIndex(938, (d_320_v99_).length(0))
            rhs36_ = (d_323_v102_).cf2
            rhs37_ = (default__.fm7(d_210_globalState_)) not in (d_324_v103_)
            rhs38_ = d_203_v1_
            lhs30_ = d_210_globalState_
            lhs31_ = d_318_v97_
            lhs32_ = default__.safeIndex(29, (d_318_v97_).length(0))
            lhs33_ = d_320_v99_
            lhs34_ = default__.safeIndex(938, (d_320_v99_).length(0))
            lhs30_.f6 = rhs36_
            lhs31_[lhs32_] = rhs37_
            lhs33_[lhs34_] = rhs38_
            d_325_v104_: _dafny.Array
            nw62_ = _dafny.Array(_dafny.Map({}), 24)
            d_325_v104_ = nw62_
            d_326_v105_: _dafny.Array
            nw63_ = _dafny.Array(None, 4)
            nw63_[int(0)] = d_325_v104_
            nw63_[int(1)] = d_325_v104_
            nw63_[int(2)] = d_325_v104_
            nw63_[int(3)] = d_325_v104_
            d_326_v105_ = nw63_
            d_327_v106_: _dafny.Map
            d_327_v106_ = _dafny.Map({(d_320_v99_)[default__.safeIndex(938, (d_320_v99_).length(0))]: d_325_v104_})
            d_328_v107_: _dafny.MultiSet
            d_328_v107_ = _dafny.MultiSet([False, (d_318_v97_)[default__.safeIndex(29, (d_318_v97_).length(0))]])
            index53_ = default__.safeIndex(718, (d_326_v105_).length(0))
            (d_326_v105_)[index53_] = ((d_327_v106_)[(d_328_v107_).cardinality] if ((d_328_v107_).cardinality) in (d_327_v106_) else d_325_v104_)
            index54_ = default__.safeIndex(718, (d_326_v105_).length(0))
            rhs39_ = (d_321_v100_)[default__.safeIndex(268, len(d_321_v100_))]
            rhs40_ = d_325_v104_
            lhs35_ = d_326_v105_
            lhs36_ = default__.safeIndex(718, (d_326_v105_).length(0))
            d_206_v4_ = rhs39_
            lhs35_[lhs36_] = rhs40_
            d_329_v108_: C0
            nw64_ = C0()
            nw64_.ctor__((d_320_v99_)[default__.safeIndex(938, (d_320_v99_).length(0))], d_203_v1_, d_206_v4_)
            d_329_v108_ = nw64_
            d_330_v109_: D1
            d_330_v109_ = D1_DC4()
            d_331_v110_: _dafny.Set
            d_331_v110_ = _dafny.Set({d_330_v109_})
            index55_ = default__.safeIndex(938, (d_320_v99_).length(0))
            index56_ = default__.safeIndex(29, (d_318_v97_).length(0))
            rhs41_ = default__.safeModuloInt(default__.safeModuloInt((d_329_v108_).f9, len(d_331_v110_)), (d_329_v108_).f9)
            rhs42_ = (_dafny.MultiSet((d_297_v81_) + (d_298_v82_))).isdisjoint(d_208_v6_)
            lhs37_ = d_320_v99_
            lhs38_ = default__.safeIndex(938, (d_320_v99_).length(0))
            lhs39_ = d_318_v97_
            lhs40_ = default__.safeIndex(29, (d_318_v97_).length(0))
            lhs37_[lhs38_] = rhs41_
            lhs39_[lhs40_] = rhs42_
            if (d_318_v97_)[default__.safeIndex(29, (d_318_v97_).length(0))]:
                d_332_v111_: _dafny.Map
                d_332_v111_ = _dafny.Map({False: (d_318_v97_)[default__.safeIndex(29, (d_318_v97_).length(0))]})
                d_333_v112_: _dafny.Map
                d_333_v112_ = _dafny.Map({d_332_v111_: d_204_v2_})
                d_333_v112_ = (d_333_v112_).set(_dafny.Map({(d_318_v97_)[default__.safeIndex(29, (d_318_v97_).length(0))]: d_206_v4_}), d_204_v2_)
                d_334_v113_: T0
                nw65_ = C0()
                nw65_.ctor__((d_329_v108_).f9, (d_320_v99_)[default__.safeIndex(938, (d_320_v99_).length(0))], d_206_v4_)
                d_334_v113_ = nw65_
                d_335_v114_: _dafny.Map
                d_335_v114_ = _dafny.Map({d_334_v113_: (d_318_v97_)[default__.safeIndex(29, (d_318_v97_).length(0))]})
                index57_ = default__.safeIndex(29, (d_318_v97_).length(0))
                (d_318_v97_)[index57_] = ((d_335_v114_)[d_334_v113_] if (d_334_v113_) in (d_335_v114_) else True)
                d_336_v115_: _dafny.Seq
                out7_: _dafny.Seq
                out7_ = default__.m0(907, _dafny.MultiSet([d_206_v4_, (d_318_v97_)[default__.safeIndex(29, (d_318_v97_).length(0))]]), d_321_v100_, d_210_globalState_)
                d_336_v115_ = out7_
                rhs43_ = (d_329_v108_).f8
                rhs44_ = d_320_v99_
                rhs45_ = d_332_v111_
                rhs46_ = (d_329_v108_).f9
                rhs47_ = d_329_v108_
                lhs41_ = d_210_globalState_
                lhs42_ = d_210_globalState_
                lhs41_.f6 = rhs43_
                d_320_v99_ = rhs44_
                d_332_v111_ = rhs45_
                lhs42_.f3 = rhs46_
                d_329_v108_ = rhs47_
                index58_ = default__.safeIndex(938, (d_320_v99_).length(0))
                (d_320_v99_)[index58_] = (d_329_v108_).f8
            elif True:
                index59_ = default__.safeIndex(29, (d_318_v97_).length(0))
                index60_ = default__.safeIndex(938, (d_320_v99_).length(0))
                rhs48_ = d_206_v4_
                rhs49_ = (d_329_v108_).f9
                lhs43_ = d_318_v97_
                lhs44_ = default__.safeIndex(29, (d_318_v97_).length(0))
                lhs45_ = d_320_v99_
                lhs46_ = default__.safeIndex(938, (d_320_v99_).length(0))
                lhs43_[lhs44_] = rhs48_
                lhs45_[lhs46_] = rhs49_
                d_337_v116_: _dafny.Map
                d_337_v116_ = _dafny.Map({(d_318_v97_)[default__.safeIndex(29, (d_318_v97_).length(0))]: d_206_v4_})
                (d_210_globalState_).f3 = (default__.safeModuloInt(len(d_337_v116_), (d_298_v82_)[default__.safeIndex((_dafny.MultiSet([d_206_v4_, d_206_v4_])).cardinality, len(d_298_v82_))])) * (default__.fm0((d_320_v99_)[default__.safeIndex(938, (d_320_v99_).length(0))], (d_318_v97_)[default__.safeIndex(29, (d_318_v97_).length(0))], _dafny.SeqWithoutIsStrInference([(d_329_v108_).f9, (d_329_v108_).f8, (d_320_v99_)[default__.safeIndex(938, (d_320_v99_).length(0))]]), d_204_v2_, d_210_globalState_))
                d_338_v117_: _dafny.Seq
                d_338_v117_ = _dafny.SeqWithoutIsStrInference([d_202_v0_])
                index61_ = default__.safeIndex(29, (d_318_v97_).length(0))
                index62_ = default__.safeIndex(29, (d_318_v97_).length(0))
                rhs50_ = ((d_329_v108_).f8) < ((d_329_v108_).fm5(default__.fm8((d_329_v108_).f9, (d_320_v99_)[default__.safeIndex(938, (d_320_v99_).length(0))], 356, (d_320_v99_)[default__.safeIndex(938, (d_320_v99_).length(0))], d_210_globalState_), d_204_v2_, d_203_v1_, d_206_v4_, d_210_globalState_))
                rhs51_ = ((d_299_v83_ if (d_318_v97_)[default__.safeIndex(29, (d_318_v97_).length(0))] else d_299_v83_) if ((d_338_v117_)[default__.safeIndex((d_329_v108_).f9, len(d_338_v117_))]) != (d_202_v0_) else default__.fm9(d_210_globalState_))
                rhs52_ = not ((d_318_v97_)[default__.safeIndex(29, (d_318_v97_).length(0))]) or (d_206_v4_)
                lhs47_ = d_318_v97_
                lhs48_ = default__.safeIndex(29, (d_318_v97_).length(0))
                lhs49_ = d_318_v97_
                lhs50_ = default__.safeIndex(29, (d_318_v97_).length(0))
                lhs47_[lhs48_] = rhs50_
                d_299_v83_ = rhs51_
                lhs49_[lhs50_] = rhs52_
                d_337_v116_ = (d_337_v116_).set((d_318_v97_)[default__.safeIndex(29, (d_318_v97_).length(0))], (d_318_v97_)[default__.safeIndex(29, (d_318_v97_).length(0))])
                d_339_v118_: _dafny.Seq
                out8_: _dafny.Seq
                out8_ = default__.m0((d_329_v108_).f8, _dafny.MultiSet([(d_318_v97_)[default__.safeIndex(29, (d_318_v97_).length(0))]]), d_321_v100_, d_210_globalState_)
                d_339_v118_ = out8_
        elif True:
            d_340_v119_: _dafny.Seq
            d_340_v119_ = _dafny.SeqWithoutIsStrInference([d_206_v4_])
            d_341_v120_: _dafny.MultiSet
            d_341_v120_ = _dafny.MultiSet([d_340_v119_, d_340_v119_])
            d_342_v121_: D0
            d_342_v121_ = D0_DC1(d_341_v120_, d_203_v1_)
            d_343_v122_: _dafny.MultiSet
            d_343_v122_ = _dafny.MultiSet([D0_DC1(d_341_v120_, len(d_340_v119_)), d_342_v121_])
            d_206_v4_ = not((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_342_v121_]))).ispropersubset(d_343_v122_))
            d_344_v123_: _dafny.MultiSet
            d_344_v123_ = _dafny.MultiSet([d_206_v4_, d_206_v4_])
            d_345_v124_: _dafny.Seq
            out9_: _dafny.Seq
            out9_ = default__.m0(len(d_202_v0_), d_344_v123_, (d_340_v119_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_206_v4_])), len(d_340_v119_)), d_206_v4_), d_210_globalState_)
            d_345_v124_ = out9_
            d_346_v125_: _dafny.Map
            d_346_v125_ = _dafny.Map({d_203_v1_: False})
            d_206_v4_ = not((default__.fm3(default__.fm0(d_203_v1_, d_206_v4_, d_297_v81_, d_204_v2_, d_210_globalState_), d_210_globalState_)) and (((d_346_v125_)[d_203_v1_] if (d_203_v1_) in (d_346_v125_) else d_206_v4_)))
            d_347_v126_: _dafny.Seq
            out10_: _dafny.Seq
            out10_ = default__.m0(d_203_v1_, (d_344_v123_) - (d_344_v123_), default__.fm6(d_203_v1_, d_204_v2_, d_206_v4_, d_206_v4_, d_210_globalState_), d_210_globalState_)
            d_347_v126_ = out10_
            d_348_v127_: _dafny.Map
            d_348_v127_ = _dafny.Map({d_206_v4_: (0) - (d_203_v1_)})
            d_349_v128_: _dafny.Set
            d_349_v128_ = _dafny.Set({((d_348_v127_)[not(d_206_v4_)] if (not(d_206_v4_)) in (d_348_v127_) else d_203_v1_)})
            d_349_v128_ = d_349_v128_
        d_350_v129_: _dafny.MultiSet
        d_350_v129_ = _dafny.MultiSet([d_206_v4_, d_206_v4_, d_206_v4_, d_206_v4_, default__.fm3(d_203_v1_, d_210_globalState_)])
        d_351_v130_: _dafny.Map
        d_351_v130_ = _dafny.Map({d_204_v2_: d_350_v129_})
        d_351_v130_ = (d_351_v130_).set(d_204_v2_, d_350_v129_)
        d_206_v4_ = d_206_v4_
        d_206_v4_ = (d_317_v96_).cf0
        d_352_v132_: C0
        nw66_ = C0()
        def iife23_():
            coll11_ = _dafny.Set()
            compr_11_: int
            for compr_11_ in _dafny.IntegerRange(416, 754):
                d_353_v131_: int = compr_11_
                if ((416) <= (d_353_v131_)) and ((d_353_v131_) < (754)):
                    coll11_ = coll11_.union(_dafny.Set([(d_353_v131_) * (d_203_v1_)]))
            return _dafny.Set(coll11_)
        nw66_.ctor__(d_203_v1_, len(iife23_()
        ), d_206_v4_)
        d_352_v132_ = nw66_
        d_354_v133_: D2
        d_354_v133_ = D2_DC6(d_352_v132_)
        d_355_v134_: _dafny.Array
        nw67_ = _dafny.Array(None, 29)
        nw67_[int(0)] = d_352_v132_
        nw67_[int(1)] = d_352_v132_
        nw67_[int(2)] = (d_354_v133_).cf6
        nw67_[int(3)] = d_352_v132_
        nw67_[int(4)] = d_352_v132_
        nw67_[int(5)] = d_352_v132_
        nw67_[int(6)] = (d_352_v132_ if False else d_352_v132_)
        nw67_[int(7)] = d_352_v132_
        nw67_[int(8)] = d_352_v132_
        nw67_[int(9)] = d_352_v132_
        nw67_[int(10)] = d_352_v132_
        nw67_[int(11)] = d_352_v132_
        nw67_[int(12)] = d_352_v132_
        nw67_[int(13)] = d_352_v132_
        nw67_[int(14)] = d_352_v132_
        nw67_[int(15)] = d_352_v132_
        nw67_[int(16)] = d_352_v132_
        nw67_[int(17)] = d_352_v132_
        nw67_[int(18)] = d_352_v132_
        nw67_[int(19)] = d_352_v132_
        nw67_[int(20)] = d_352_v132_
        nw67_[int(21)] = d_352_v132_
        nw67_[int(22)] = d_352_v132_
        nw67_[int(23)] = d_352_v132_
        nw67_[int(24)] = d_352_v132_
        nw67_[int(25)] = (d_352_v132_ if d_206_v4_ else d_352_v132_)
        nw67_[int(26)] = (d_352_v132_ if d_206_v4_ else d_352_v132_)
        nw67_[int(27)] = d_352_v132_
        nw67_[int(28)] = d_352_v132_
        d_355_v134_ = nw67_
        index63_ = default__.safeIndex(5, (d_355_v134_).length(0))
        (d_355_v134_)[index63_] = d_352_v132_
        d_356_v135_: _dafny.Array
        def lambda18_(d_357_v4_, d_358_v2_, d_359_v132_):
            def lambda19_(d_360_i8_):
                return D0_DC1(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_357_v4_, d_357_v4_]), (D3_DC9(_dafny.SeqWithoutIsStrInference([d_357_v4_, d_357_v4_]))).cf10, _dafny.SeqWithoutIsStrInference([(D3_DC10(d_358_v2_, d_357_v4_)).cf12])]), (d_359_v132_).f9)

            return lambda19_

        init9_ = lambda18_(d_206_v4_, d_204_v2_, d_352_v132_)
        nw68_ = _dafny.Array(None, 4)
        for i0_9_ in range(nw68_.length(0)):
            nw68_[i0_9_] = init9_(i0_9_)
        d_356_v135_ = nw68_
        d_361_v136_: _dafny.Seq
        d_361_v136_ = _dafny.SeqWithoutIsStrInference([d_206_v4_])
        d_362_v137_: _dafny.MultiSet
        d_362_v137_ = _dafny.MultiSet([d_361_v136_])
        d_363_v138_: D0
        d_363_v138_ = D0_DC1(d_362_v137_, (d_352_v132_).f9)
        index64_ = default__.safeIndex(568, (d_356_v135_).length(0))
        (d_356_v135_)[index64_] = d_363_v138_
        d_364_v139_: _dafny.Array
        nw69_ = _dafny.Array(int(0), 16)
        d_364_v139_ = nw69_
        d_365_v140_: _dafny.Seq
        d_365_v140_ = _dafny.SeqWithoutIsStrInference([d_364_v139_])
        d_366_v141_: _dafny.MultiSet
        d_366_v141_ = _dafny.MultiSet([d_297_v81_])
        index65_ = default__.safeIndex(5, (d_355_v134_).length(0))
        index66_ = default__.safeIndex(568, (d_356_v135_).length(0))
        rhs53_ = len((d_365_v140_) + (_dafny.SeqWithoutIsStrInference([d_364_v139_])))
        rhs54_ = d_352_v132_
        rhs55_ = d_363_v138_
        rhs56_ = default__.fm10(len(d_361_v136_), not(((d_352_v132_).f8) <= ((d_352_v132_).f8)), d_203_v1_, (d_366_v141_) - (d_366_v141_), d_210_globalState_)
        lhs51_ = d_210_globalState_
        lhs52_ = d_355_v134_
        lhs53_ = default__.safeIndex(5, (d_355_v134_).length(0))
        lhs54_ = d_356_v135_
        lhs55_ = default__.safeIndex(568, (d_356_v135_).length(0))
        lhs51_.f6 = rhs53_
        lhs52_[lhs53_] = rhs54_
        lhs54_[lhs55_] = rhs55_
        d_202_v0_ = rhs56_
        d_367_v142_: D3
        d_367_v142_ = D3_DC10(d_204_v2_, d_206_v4_)
        source3_ = d_367_v142_
        if source3_.is_DC10:
            d_368___mcc_h2_ = source3_.cf11
            d_369___mcc_h3_ = source3_.cf12
            d_370_cf12_ = d_369___mcc_h3_
            d_371_cf11_ = d_368___mcc_h2_
            d_372_v143_: _dafny.MultiSet
            d_372_v143_ = _dafny.MultiSet([_dafny.Map({d_206_v4_: (d_352_v132_).f8})])
            d_371_cf11_ = default__.fm11(d_372_v143_, ((d_208_v6_).cardinality) + (d_203_v1_), d_202_v0_, d_202_v0_, d_210_globalState_)
            d_206_v4_ = d_206_v4_
            d_373_v144_: _dafny.Map
            d_373_v144_ = _dafny.Map({d_206_v4_: d_203_v1_})
            d_374_v145_: _dafny.Map
            d_374_v145_ = _dafny.Map({d_364_v139_: d_373_v144_})
            d_375_v146_: _dafny.Array
            nw70_ = _dafny.Array(None, 8)
            nw70_[int(0)] = d_204_v2_
            nw70_[int(1)] = default__.fm11(_dafny.MultiSet([d_373_v144_, ((d_374_v145_)[d_364_v139_] if (d_364_v139_) in (d_374_v145_) else d_373_v144_)]), d_203_v1_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "unvm")), d_202_v0_, d_210_globalState_)
            nw70_[int(2)] = _dafny.CodePoint('l')
            nw70_[int(3)] = d_371_cf11_
            nw70_[int(4)] = _dafny.CodePoint('a')
            nw70_[int(5)] = d_371_cf11_
            nw70_[int(6)] = d_371_cf11_
            nw70_[int(7)] = d_371_cf11_
            d_375_v146_ = nw70_
            index67_ = default__.safeIndex(842, (d_375_v146_).length(0))
            (d_375_v146_)[index67_] = d_204_v2_
            index68_ = default__.safeIndex(842, (d_375_v146_).length(0))
            (d_375_v146_)[index68_] = d_371_cf11_
            source4_ = d_317_v96_
            if source4_.is_DC1:
                d_376___mcc_h9_ = source4_.cf1
                d_377___mcc_h10_ = source4_.cf2
                d_378_cf2_ = d_377___mcc_h10_
                d_379_cf1_ = d_376___mcc_h9_
                (d_210_globalState_).f3 = (d_352_v132_).f9
                d_378_cf2_ = ((d_352_v132_).f9) + ((0) - (d_203_v1_))
                d_380_v147_: _dafny.Map
                d_380_v147_ = _dafny.Map({d_248_v40_: d_297_v81_})
                d_381_v148_: _dafny.Set
                d_381_v148_ = _dafny.Set({len(d_298_v82_)})
                d_380_v147_ = (d_380_v147_).set((d_248_v40_).set((d_352_v132_).f9, len(d_381_v148_)), d_298_v82_)
                (d_210_globalState_).f6 = (d_352_v132_).f8
            elif source4_.is_DC0:
                d_382___mcc_h11_ = source4_.cf0
                d_383_cf0_ = d_382___mcc_h11_
                d_384_v149_: _dafny.Seq
                out11_: _dafny.Seq
                out11_ = default__.m0(default__.safeModuloInt((d_352_v132_).f9, d_203_v1_), d_350_v129_, d_361_v136_, d_210_globalState_)
                d_384_v149_ = out11_
                (d_210_globalState_).f6 = default__.safeModuloInt(d_203_v1_, (d_352_v132_).f9)
                d_385_v150_: _dafny.Seq
                out12_: _dafny.Seq
                out12_ = default__.m0((d_352_v132_).f9, d_350_v129_, d_361_v136_, d_210_globalState_)
                d_385_v150_ = out12_
                (d_210_globalState_).f6 = (d_352_v132_).f8
            elif True:
                d_386___mcc_h12_ = source4_.cf3
                d_387_cf3_ = d_386___mcc_h12_
                d_371_cf11_ = d_371_cf11_
                d_388_v151_: _dafny.Map
                d_388_v151_ = _dafny.Map({(d_352_v132_).f8: _dafny.Map({d_364_v139_: d_203_v1_})})
                d_389_v152_: _dafny.Map
                d_389_v152_ = _dafny.Map({d_364_v139_: d_203_v1_})
                d_388_v151_ = (d_388_v151_).set((d_352_v132_).f9, d_389_v152_)
                d_390_v153_: _dafny.Seq
                out13_: _dafny.Seq
                out13_ = default__.m0((d_352_v132_).f9, default__.fm12((d_352_v132_).f8, len(_dafny.SeqWithoutIsStrInference([(d_375_v146_)[default__.safeIndex(842, (d_375_v146_).length(0))] for d_391_i9_ in range(default__.abs(546))])), default__.fm13((d_352_v132_).f8, d_210_globalState_), d_210_globalState_), d_361_v136_, d_210_globalState_)
                d_390_v153_ = out13_
                (d_210_globalState_).f3 = len(d_390_v153_)
        elif source3_.is_DC11:
            d_392___mcc_h4_ = source3_.cf13
            d_393___mcc_h5_ = source3_.cf14
            d_394___mcc_h6_ = source3_.cf15
            d_395___mcc_h7_ = source3_.cf16
            d_396_cf16_ = d_395___mcc_h7_
            d_397_cf15_ = d_394___mcc_h6_
            d_398_cf14_ = d_393___mcc_h5_
            d_399_cf13_ = d_392___mcc_h4_
            d_400_v154_: _dafny.Array
            def lambda20_(d_401_cf16_):
                def lambda21_(d_402_i10_):
                    return (d_401_cf16_) and (d_401_cf16_)

                return lambda21_

            init10_ = lambda20_(d_396_cf16_)
            nw71_ = _dafny.Array(None, 1)
            for i0_10_ in range(nw71_.length(0)):
                nw71_[i0_10_] = init10_(i0_10_)
            d_400_v154_ = nw71_
            index69_ = default__.safeIndex(526, (d_400_v154_).length(0))
            (d_400_v154_)[index69_] = d_206_v4_
            index70_ = default__.safeIndex(526, (d_400_v154_).length(0))
            (d_400_v154_)[index70_] = not((390) >= (((d_352_v132_).f8) * (d_203_v1_)))
            d_403_v155_: D4
            d_403_v155_ = D4_DC12(d_364_v139_)
            d_404_v156_: _dafny.Array
            nw72_ = _dafny.Array(None, 12)
            nw72_[int(0)] = d_364_v139_
            nw72_[int(1)] = d_364_v139_
            nw72_[int(2)] = d_364_v139_
            nw72_[int(3)] = d_364_v139_
            nw72_[int(4)] = d_364_v139_
            nw72_[int(5)] = (d_403_v155_).cf17
            nw72_[int(6)] = d_364_v139_
            nw72_[int(7)] = (d_364_v139_ if d_206_v4_ else d_364_v139_)
            nw72_[int(8)] = d_364_v139_
            nw72_[int(9)] = d_364_v139_
            nw72_[int(10)] = d_364_v139_
            nw72_[int(11)] = d_364_v139_
            d_404_v156_ = nw72_
            index71_ = default__.safeIndex(603, (d_404_v156_).length(0))
            (d_404_v156_)[index71_] = d_364_v139_
            index72_ = default__.safeIndex(603, (d_404_v156_).length(0))
            rhs57_ = d_364_v139_
            rhs58_ = (d_352_v132_).f8
            lhs56_ = d_404_v156_
            lhs57_ = default__.safeIndex(603, (d_404_v156_).length(0))
            lhs56_[lhs57_] = rhs57_
            d_399_cf13_ = rhs58_
            d_405_v157_: _dafny.Map
            d_405_v157_ = _dafny.Map({not(True): default__.fm3(d_203_v1_, d_210_globalState_)})
            d_406_v158_: _dafny.Set
            d_406_v158_ = _dafny.Set({d_248_v40_, _dafny.Map({len(d_207_v5_): len(d_405_v157_)}), d_248_v40_})
            if not((d_406_v158_).isdisjoint(d_406_v158_)):
                (d_210_globalState_).f3 = ((d_352_v132_).f9 if d_396_cf16_ else (d_352_v132_).f8)
                d_407_v159_: _dafny.Map
                d_407_v159_ = _dafny.Map({d_400_v154_: (d_400_v154_)[default__.safeIndex(526, (d_400_v154_).length(0))]})
                index73_ = default__.safeIndex(5, (d_355_v134_).length(0))
                nw73_ = C0()
                nw73_.ctor__(d_203_v1_, len(d_407_v159_), d_397_cf15_)
                (d_355_v134_)[index73_] = nw73_
                d_202_v0_ = ((d_202_v0_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_408_i11_ in range(default__.abs(-203))]))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uw"))) + (d_202_v0_))
                d_409_v160_: _dafny.Array
                nw74_ = _dafny.Array(None, 14)
                d_409_v160_ = nw74_
                d_410_v161_: T0
                nw75_ = C0()
                nw75_.ctor__((d_352_v132_).f8, -634, (d_400_v154_)[default__.safeIndex(526, (d_400_v154_).length(0))])
                d_410_v161_ = nw75_
                index74_ = default__.safeIndex(585, (d_409_v160_).length(0))
                (d_409_v160_)[index74_] = d_410_v161_
                index75_ = default__.safeIndex(585, (d_409_v160_).length(0))
                (d_409_v160_)[index75_] = d_410_v161_
                d_202_v0_ = d_202_v0_
            elif True:
                d_411_v162_: _dafny.Array
                nw76_ = _dafny.Array(D1.default()(), 1)
                d_411_v162_ = nw76_
                index76_ = default__.safeIndex(240, (d_411_v162_).length(0))
                (d_411_v162_)[index76_] = d_299_v83_
                index77_ = default__.safeIndex(240, (d_411_v162_).length(0))
                (d_411_v162_)[index77_] = D1_DC3(d_298_v82_)
                (d_210_globalState_).f6 = (((d_352_v132_).f8) + ((d_352_v132_).f9)) - (default__.safeModuloInt((d_352_v132_).f9, d_399_cf13_))
                d_412_v163_: _dafny.Map
                d_412_v163_ = _dafny.Map({(d_352_v132_).f9: (d_400_v154_)[default__.safeIndex(526, (d_400_v154_).length(0))]})
                d_412_v163_ = (d_412_v163_).set((d_352_v132_).f9, d_397_cf15_)
                d_397_cf15_ = d_397_cf15_
                index78_ = default__.safeIndex(526, (d_400_v154_).length(0))
                (d_400_v154_)[index78_] = (d_400_v154_)[default__.safeIndex(526, (d_400_v154_).length(0))]
            d_413_v164_: _dafny.Set
            d_413_v164_ = _dafny.Set({(d_352_v132_).f9, (0) - ((d_352_v132_).f8)})
            def iife24_():
                coll12_ = _dafny.Set()
                compr_12_: int
                for compr_12_ in (d_413_v164_).Elements:
                    d_414_v165_: int = compr_12_
                    if (d_414_v165_) in (d_413_v164_):
                        coll12_ = coll12_.union(_dafny.Set([(d_414_v165_) + (len(_dafny.SeqWithoutIsStrInference([True, True])))]))
                return _dafny.Set(coll12_)
            if (iife24_()
            ) == ((d_413_v164_) - (d_413_v164_)):
                d_415_v166_: C0
                nw77_ = C0()
                nw77_.ctor__(d_203_v1_, (d_352_v132_).f8, (d_400_v154_)[default__.safeIndex(526, (d_400_v154_).length(0))])
                d_415_v166_ = nw77_
                d_203_v1_ = ((620) * ((d_352_v132_).f8)) + ((d_363_v138_).cf2)
                d_416_v167_: D3
                d_416_v167_ = D3_DC11(len(d_405_v157_), 642, default__.fm3((d_415_v166_).fm4(d_210_globalState_), d_210_globalState_), (d_400_v154_)[default__.safeIndex(526, (d_400_v154_).length(0))])
                d_417_v168_: _dafny.Seq
                d_417_v168_ = _dafny.SeqWithoutIsStrInference([d_416_v167_])
                d_417_v168_ = d_417_v168_
                d_397_cf15_ = (not(((d_400_v154_)[default__.safeIndex(526, (d_400_v154_).length(0))] if d_396_cf16_ else d_397_cf15_))) and (True)
                d_418_v169_: _dafny.Map
                d_418_v169_ = _dafny.Map({False: (d_415_v166_).f9})
                index79_ = default__.safeIndex(526, (d_400_v154_).length(0))
                (d_400_v154_)[index79_] = ((len(d_418_v169_)) == ((d_415_v166_).f9) if (d_361_v136_)[default__.safeIndex(d_398_cf14_, len(d_361_v136_))] else d_397_cf15_)
            elif True:
                d_419_v170_: _dafny.Map
                d_419_v170_ = _dafny.Map({d_399_cf13_: (d_202_v0_).set(default__.safeIndex(d_398_cf14_, len(d_202_v0_)), d_204_v2_)})
                d_420_v171_: _dafny.Seq
                out14_: _dafny.Seq
                out14_ = default__.m0((d_352_v132_).f9, d_350_v129_, default__.fm6(len(((d_419_v170_)[(d_352_v132_).f8] if ((d_352_v132_).f8) in (d_419_v170_) else d_202_v0_)), d_204_v2_, not(False), default__.fm3(d_399_cf13_, d_210_globalState_), d_210_globalState_), d_210_globalState_)
                d_420_v171_ = out14_
                rhs59_ = d_350_v129_
                rhs60_ = (d_203_v1_) + (len(_dafny.SeqWithoutIsStrInference([True, d_206_v4_, (d_400_v154_)[default__.safeIndex(526, (d_400_v154_).length(0))]])))
                lhs58_ = d_210_globalState_
                d_350_v129_ = rhs59_
                lhs58_.f3 = rhs60_
                index80_ = default__.safeIndex(526, (d_400_v154_).length(0))
                (d_400_v154_)[index80_] = d_397_cf15_
                d_421_v172_: _dafny.Array
                nw78_ = _dafny.Array(_dafny.Array(None, 0), 12)
                d_421_v172_ = nw78_
                index81_ = default__.safeIndex(306, (d_421_v172_).length(0))
                (d_421_v172_)[index81_] = d_400_v154_
                index82_ = default__.safeIndex(306, (d_421_v172_).length(0))
                rhs61_ = d_398_cf14_
                rhs62_ = d_400_v154_
                rhs63_ = d_350_v129_
                lhs59_ = d_210_globalState_
                lhs60_ = d_421_v172_
                lhs61_ = default__.safeIndex(306, (d_421_v172_).length(0))
                lhs59_.f3 = rhs61_
                lhs60_[lhs61_] = rhs62_
                d_350_v129_ = rhs63_
                d_422_v173_: _dafny.Seq
                out15_: _dafny.Seq
                out15_ = default__.m0((d_352_v132_).fm5(d_405_v157_, d_204_v2_, 213, (d_400_v154_)[default__.safeIndex(526, (d_400_v154_).length(0))], d_210_globalState_), d_350_v129_, _dafny.SeqWithoutIsStrInference([d_206_v4_]), d_210_globalState_)
                d_422_v173_ = out15_
        elif True:
            d_423___mcc_h8_ = source3_.cf10
            d_424_cf10_ = d_423___mcc_h8_
            if not(not(not((d_206_v4_) == (d_206_v4_)))):
                d_425_v174_: _dafny.Array
                nw79_ = _dafny.Array(False, 20)
                d_425_v174_ = nw79_
                index83_ = default__.safeIndex(369, (d_425_v174_).length(0))
                (d_425_v174_)[index83_] = (not(d_206_v4_)) == (d_206_v4_)
                index84_ = default__.safeIndex(369, (d_425_v174_).length(0))
                (d_425_v174_)[index84_] = (len(d_248_v40_)) != ((d_352_v132_).f8)
                (d_210_globalState_).f6 = len((d_202_v0_) + (d_202_v0_))
                d_426_v175_: _dafny.Array
                nw80_ = _dafny.Array(None, 1)
                d_426_v175_ = nw80_
                d_427_v176_: T0
                nw81_ = C0()
                nw81_.ctor__((0) - ((d_352_v132_).f8), (d_352_v132_).f8, d_206_v4_)
                d_427_v176_ = nw81_
                index85_ = default__.safeIndex(6, (d_426_v175_).length(0))
                (d_426_v175_)[index85_] = d_427_v176_
                index86_ = default__.safeIndex(6, (d_426_v175_).length(0))
                (d_426_v175_)[index86_] = d_427_v176_
                d_428_v177_: _dafny.Seq
                out16_: _dafny.Seq
                out16_ = default__.m0((d_352_v132_).f8, d_350_v129_, d_424_cf10_, d_210_globalState_)
                d_428_v177_ = out16_
                index87_ = default__.safeIndex(369, (d_425_v174_).length(0))
                (d_425_v174_)[index87_] = default__.fm3(((d_352_v132_).f8) * ((d_352_v132_).f9), d_210_globalState_)
            elif True:
                d_429_v178_: _dafny.Seq
                d_429_v178_ = _dafny.SeqWithoutIsStrInference([d_352_v132_, d_352_v132_])
                d_430_v179_: T0
                nw82_ = C0()
                nw82_.ctor__(len(d_202_v0_), len(d_429_v178_), d_206_v4_)
                d_430_v179_ = nw82_
                d_431_v180_: _dafny.Seq
                d_431_v180_ = _dafny.SeqWithoutIsStrInference([d_430_v179_])
                d_432_v181_: _dafny.MultiSet
                d_432_v181_ = _dafny.MultiSet([d_430_v179_])
                d_433_v182_: _dafny.Array
                nw83_ = _dafny.Array(None, 23)
                nw83_[int(0)] = d_430_v179_
                nw83_[int(1)] = d_430_v179_
                nw83_[int(2)] = d_430_v179_
                nw83_[int(3)] = d_430_v179_
                nw83_[int(4)] = (d_431_v180_)[default__.safeIndex(((d_432_v181_).set(d_430_v179_, default__.abs((d_352_v132_).f8))).cardinality, len(d_431_v180_))]
                nw83_[int(5)] = d_430_v179_
                nw83_[int(6)] = d_430_v179_
                nw83_[int(7)] = d_430_v179_
                nw83_[int(8)] = d_430_v179_
                nw83_[int(9)] = d_430_v179_
                nw83_[int(10)] = d_430_v179_
                nw83_[int(11)] = d_430_v179_
                nw83_[int(12)] = d_430_v179_
                nw83_[int(13)] = d_430_v179_
                nw83_[int(14)] = d_430_v179_
                nw83_[int(15)] = d_430_v179_
                nw83_[int(16)] = d_430_v179_
                nw83_[int(17)] = d_430_v179_
                nw83_[int(18)] = d_430_v179_
                nw83_[int(19)] = (d_430_v179_ if True else d_430_v179_)
                nw83_[int(20)] = d_430_v179_
                nw83_[int(21)] = d_430_v179_
                nw83_[int(22)] = d_430_v179_
                d_433_v182_ = nw83_
                d_434_v183_: T0
                d_434_v183_ = d_430_v179_
                d_435_v184_: _dafny.Map
                d_435_v184_ = _dafny.Map({793: d_430_v179_})
                nw84_ = _dafny.Array(None, 20)
                nw84_[int(0)] = d_430_v179_
                nw84_[int(1)] = d_430_v179_
                nw84_[int(2)] = d_430_v179_
                nw84_[int(3)] = d_430_v179_
                nw84_[int(4)] = d_430_v179_
                nw84_[int(5)] = d_430_v179_
                nw84_[int(6)] = d_430_v179_
                nw84_[int(7)] = d_430_v179_
                nw84_[int(8)] = d_430_v179_
                nw84_[int(9)] = d_430_v179_
                nw84_[int(10)] = d_430_v179_
                nw84_[int(11)] = d_430_v179_
                nw84_[int(12)] = (d_434_v183_)
                nw84_[int(13)] = d_430_v179_
                nw84_[int(14)] = d_430_v179_
                nw84_[int(15)] = d_430_v179_
                nw84_[int(16)] = d_430_v179_
                nw84_[int(17)] = ((d_435_v184_)[(d_352_v132_).f8] if ((d_352_v132_).f8) in (d_435_v184_) else d_430_v179_)
                nw84_[int(18)] = d_430_v179_
                nw84_[int(19)] = d_430_v179_
                d_433_v182_ = nw84_
                d_436_v185_: _dafny.Seq
                out17_: _dafny.Seq
                out17_ = default__.m0(d_203_v1_, d_350_v129_, d_361_v136_, d_210_globalState_)
                d_436_v185_ = out17_
                d_437_v186_: _dafny.Array
                def lambda22_(d_438_i12_):
                    return True

                init11_ = lambda22_
                nw85_ = _dafny.Array(None, 26)
                for i0_11_ in range(nw85_.length(0)):
                    nw85_[i0_11_] = init11_(i0_11_)
                d_437_v186_ = nw85_
                d_439_v187_: _dafny.Map
                d_439_v187_ = _dafny.Map({d_352_v132_: len(_dafny.Map({144: d_430_v179_}))})
                rhs64_ = d_437_v186_
                rhs65_ = d_430_v179_.f7
                rhs66_ = (d_203_v1_) * (((d_352_v132_).f9) - ((0) - (len(_dafny.SeqWithoutIsStrInference([d_439_v187_])))))
                rhs67_ = d_364_v139_
                rhs68_ = ((d_352_v132_).f8) * ((0) - ((d_352_v132_).f8))
                lhs62_ = d_210_globalState_
                d_437_v186_ = rhs64_
                d_206_v4_ = rhs65_
                lhs62_.f6 = rhs66_
                d_364_v139_ = rhs67_
                d_203_v1_ = rhs68_
                d_440_v188_: _dafny.MultiSet
                d_440_v188_ = _dafny.MultiSet([d_437_v186_])
                d_441_v189_: D3
                d_441_v189_ = D3_DC11((d_352_v132_).f9, (d_430_v179_).fm4(d_210_globalState_), (d_440_v188_).issubset(d_440_v188_), True)
                d_441_v189_ = d_441_v189_
                d_424_cf10_ = (d_424_cf10_) + (d_361_v136_)
            d_206_v4_ = not (not(d_206_v4_)) or (((d_352_v132_).f9) == ((d_352_v132_).f9))
            d_442_v190_: _dafny.Seq
            out18_: _dafny.Seq
            out18_ = default__.m0((d_352_v132_).f9, d_350_v129_, d_424_cf10_, d_210_globalState_)
            d_442_v190_ = out18_
            rhs69_ = d_206_v4_
            rhs70_ = (d_352_v132_).f8
            lhs63_ = d_210_globalState_
            d_206_v4_ = rhs69_
            lhs63_.f6 = rhs70_
        d_206_v4_ = d_206_v4_
        rhs71_ = (d_361_v136_)[default__.safeIndex(((d_352_v132_).f8 if d_206_v4_ else (d_352_v132_).f9), len(d_361_v136_))]
        rhs72_ = True
        rhs73_ = (0) - ((d_203_v1_) + (d_203_v1_))
        rhs74_ = (d_352_v132_).f9
        lhs64_ = d_210_globalState_
        lhs65_ = d_210_globalState_
        d_206_v4_ = rhs71_
        d_206_v4_ = rhs72_
        lhs64_.f3 = rhs73_
        lhs65_.f3 = rhs74_
        if not (d_206_v4_) or (True):
            d_443_v191_: _dafny.Seq
            out19_: _dafny.Seq
            out19_ = default__.m0(d_203_v1_, _dafny.MultiSet(d_361_v136_), (d_361_v136_) + (_dafny.SeqWithoutIsStrInference([d_206_v4_, d_206_v4_, d_206_v4_])), d_210_globalState_)
            d_443_v191_ = out19_
            d_206_v4_ = True
            d_206_v4_ = (d_202_v0_) != (d_443_v191_)
            d_444_v192_: _dafny.Array
            nw86_ = _dafny.Array(False, 18)
            d_444_v192_ = nw86_
            d_445_v193_: D1
            d_445_v193_ = D1_DC3((d_298_v82_).set(default__.safeIndex(len(_dafny.Map({d_444_v192_: d_203_v1_})), len(d_298_v82_)), (d_352_v132_).f8))
            d_446_v194_: D1
            d_446_v194_ = D1_DC5(d_445_v193_)
            source5_ = d_446_v194_
            if source5_.is_DC4:
                d_447_v195_: _dafny.Array
                nw87_ = _dafny.Array(_dafny.Seq({}), 11)
                d_447_v195_ = nw87_
                index88_ = default__.safeIndex(719, (d_447_v195_).length(0))
                (d_447_v195_)[index88_] = d_297_v81_
                index89_ = default__.safeIndex(719, (d_447_v195_).length(0))
                (d_447_v195_)[index89_] = d_297_v81_
                d_448_v196_: _dafny.Set
                d_448_v196_ = _dafny.Set({-939, (d_352_v132_).f8})
                d_449_v197_: D6
                d_449_v197_ = D6_DC16(d_448_v196_)
                (d_210_globalState_).f3 = (len((d_449_v197_).cf24)) + (d_203_v1_)
                d_450_v198_: _dafny.Map
                d_450_v198_ = _dafny.Map({(d_352_v132_).f8: False})
                d_450_v198_ = (d_450_v198_).set(d_203_v1_, False)
                d_451_v199_: _dafny.Set
                d_451_v199_ = _dafny.Set({d_206_v4_})
                d_452_v200_: _dafny.Map
                d_452_v200_ = _dafny.Map({default__.safeDivisionInt(len((d_361_v136_).set(default__.safeIndex((d_352_v132_).f8, len(d_361_v136_)), d_206_v4_)), len(d_202_v0_)): d_451_v199_})
                d_452_v200_ = (d_452_v200_).set(default__.safeModuloInt((0) - ((d_298_v82_)[default__.safeIndex(d_203_v1_, len(d_298_v82_))]), -792), (d_451_v199_) - (d_451_v199_))
            elif source5_.is_DC3:
                d_453___mcc_h13_ = source5_.cf4
                d_454_cf4_ = d_453___mcc_h13_
                d_206_v4_ = not(d_206_v4_)
                d_455_v201_: T0
                nw88_ = C0()
                nw88_.ctor__(965, (d_352_v132_).f8, default__.fm3(d_203_v1_, d_210_globalState_))
                d_455_v201_ = nw88_
                d_456_v202_: _dafny.Set
                d_456_v202_ = _dafny.Set({d_455_v201_.f7, d_455_v201_.f7, d_455_v201_.f7, False, d_455_v201_.f7})
                d_455_v201_ = (d_455_v201_ if ((d_352_v132_).f9) != (default__.fm0(len(d_456_v202_), True, _dafny.SeqWithoutIsStrInference([(d_352_v132_).f9, d_203_v1_, (d_352_v132_).f9]), d_204_v2_, d_210_globalState_)) else d_455_v201_)
                d_457_v203_: _dafny.Map
                d_457_v203_ = _dafny.Map({d_455_v201_.f7: (0) - (d_203_v1_)})
                d_458_v204_: C0
                nw89_ = C0()
                nw89_.ctor__(len((d_457_v203_) | (d_457_v203_)), (d_352_v132_).f8, False)
                d_458_v204_ = nw89_
                (d_210_globalState_).f3 = (((d_458_v204_).fm4(d_210_globalState_)) - ((d_352_v132_).f8)) + ((d_352_v132_).f9)
            elif True:
                d_459___mcc_h14_ = source5_.cf5
                d_460_cf5_ = d_459___mcc_h14_
                index90_ = default__.safeIndex(5, (d_355_v134_).length(0))
                (d_355_v134_)[index90_] = (d_352_v132_ if ((d_352_v132_).f9) == ((d_352_v132_).f8) else d_352_v132_)
                d_461_v205_: _dafny.Map
                d_461_v205_ = _dafny.Map({(d_352_v132_).f9: d_444_v192_})
                d_461_v205_ = (d_461_v205_).set(((d_352_v132_).f9 if d_206_v4_ else (d_352_v132_).f9), d_444_v192_)
                d_206_v4_ = (default__.fm14(d_210_globalState_)).issubset(d_245_v37_)
                index91_ = default__.safeIndex(748, (d_364_v139_).length(0))
                (d_364_v139_)[index91_] = (d_352_v132_).f8
                index92_ = default__.safeIndex(748, (d_364_v139_).length(0))
                (d_364_v139_)[index92_] = (d_352_v132_).f9
            d_206_v4_ = d_206_v4_
        elif True:
            (d_210_globalState_).f3 = d_203_v1_
            d_462_v206_: _dafny.Map
            d_462_v206_ = _dafny.Map({_dafny.Set({not(d_206_v4_)}): d_364_v139_})
            d_463_v207_: _dafny.Set
            d_463_v207_ = _dafny.Set({d_206_v4_})
            d_462_v206_ = (d_462_v206_).set(d_463_v207_, d_364_v139_)
            (d_210_globalState_).f3 = (d_352_v132_).f8
            if d_206_v4_:
                d_464_v208_: C0
                nw90_ = C0()
                nw90_.ctor__(((d_352_v132_).f8) * ((d_352_v132_).f9), d_203_v1_, d_206_v4_)
                d_464_v208_ = nw90_
                d_465_v209_: _dafny.MultiSet
                d_465_v209_ = _dafny.MultiSet([d_352_v132_, d_352_v132_])
                d_466_v210_: _dafny.Map
                d_466_v210_ = _dafny.Map({d_465_v209_: len(_dafny.Set({True, d_206_v4_, d_206_v4_}))})
                d_206_v4_ = not(((d_465_v209_).intersection(d_465_v209_)) not in (d_466_v210_))
                (d_210_globalState_).f6 = 65
                d_467_v212_: _dafny.Set
                def iife25_():
                    coll13_ = _dafny.Set()
                    compr_13_: int
                    for compr_13_ in (d_297_v81_).Elements:
                        d_468_v211_: int = compr_13_
                        if (d_468_v211_) in (d_297_v81_):
                            coll13_ = coll13_.union(_dafny.Set([(d_468_v211_) * (len(_dafny.Map({_dafny.SeqWithoutIsStrInference([False]): 931})))]))
                    return _dafny.Set(coll13_)
                d_467_v212_ = _dafny.Set({len(d_202_v0_), default__.safeModuloInt(len(iife25_()
                ), (d_352_v132_).f8)})
                d_467_v212_ = ((d_467_v212_).intersection(d_467_v212_)).intersection(d_467_v212_)
                d_204_v2_ = d_204_v2_
            elif True:
                d_202_v0_ = (default__.fm10((d_352_v132_).f8, d_206_v4_, ((_dafny.MultiSet([d_206_v4_, d_206_v4_])).set(default__.fm3((d_352_v132_).f8, d_210_globalState_), default__.abs((d_352_v132_).f8))).cardinality, d_366_v141_, d_210_globalState_)) + (d_202_v0_)
                d_352_v132_ = (d_355_v134_)[default__.safeIndex(5, (d_355_v134_).length(0))]
                index93_ = default__.safeIndex(54, (d_364_v139_).length(0))
                (d_364_v139_)[index93_] = (0) - ((d_352_v132_).f8)
                index94_ = default__.safeIndex(54, (d_364_v139_).length(0))
                (d_364_v139_)[index94_] = (d_203_v1_) - ((d_352_v132_).f9)
                d_469_v213_: _dafny.Map
                d_469_v213_ = _dafny.Map({not(d_206_v4_): _dafny.Set({d_206_v4_})})
                d_470_v214_: _dafny.Seq
                out20_: _dafny.Seq
                out20_ = default__.m0(len(((d_469_v213_)[d_206_v4_] if (d_206_v4_) in (d_469_v213_) else d_463_v207_)), d_350_v129_, d_361_v136_, d_210_globalState_)
                d_470_v214_ = out20_
                d_471_v215_: C0
                nw91_ = C0()
                nw91_.ctor__((d_352_v132_).f8, default__.safeModuloInt((d_352_v132_).f9, (d_352_v132_).f8), ((d_364_v139_)[default__.safeIndex(54, (d_364_v139_).length(0))]) >= ((d_352_v132_).f8))
                d_471_v215_ = nw91_
            d_472_v216_: _dafny.Map
            d_472_v216_ = _dafny.Map({d_206_v4_: d_206_v4_})
            d_473_v217_: _dafny.Map
            d_473_v217_ = _dafny.Map({d_206_v4_: not(not((d_361_v136_)[default__.safeIndex(len(d_472_v216_), len(d_361_v136_))]))})
            d_474_v218_: D3
            d_474_v218_ = D3_DC9(((d_361_v136_).set(default__.safeIndex((d_352_v132_).f9, len(d_361_v136_)), False)).set(default__.safeIndex((d_352_v132_).fm5(d_472_v216_, d_204_v2_, 472, d_206_v4_, d_210_globalState_), len((d_361_v136_).set(default__.safeIndex((d_352_v132_).f9, len(d_361_v136_)), False))), d_206_v4_))
            d_475_v219_: _dafny.Seq
            out21_: _dafny.Seq
            out21_ = default__.m0(default__.safeDivisionInt((d_352_v132_).f8, d_203_v1_), _dafny.MultiSet([d_206_v4_, d_206_v4_]), (d_474_v218_).cf10, d_210_globalState_)
            d_475_v219_ = out21_
        d_203_v1_ = (d_203_v1_) - ((d_366_v141_).cardinality)
        _dafny.print((d_202_v0_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_203_v1_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_204_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_205_v3_) == (_dafny.Map({779: _dafny.CodePoint('p')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_206_v4_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_207_v5_) == (_dafny.Map({_dafny.CodePoint('p'): False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_208_v6_) == (_dafny.MultiSet([-166]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_209_v7_) == (_dafny.MultiSet([_dafny.MultiSet([1])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_210_globalState_).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_210_globalState_).f1).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_210_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_210_globalState_.f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_210_globalState_.f4) == (_dafny.Map({_dafny.CodePoint('p'): False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_210_globalState_).f5) == (_dafny.MultiSet([_dafny.MultiSet([1])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_210_globalState_.f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_245_v37_) == (_dafny.Set({_dafny.CodePoint('p')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_246_v38_) == (_dafny.Set({_dafny.Set({_dafny.CodePoint('p')})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_247_v39_) == (_dafny.SeqWithoutIsStrInference([_dafny.Set({_dafny.Set({_dafny.CodePoint('p')})})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_248_v40_) == (_dafny.Map({778: 778}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_297_v81_) == (_dafny.SeqWithoutIsStrInference([2]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_298_v82_) == (_dafny.SeqWithoutIsStrInference([2, 2]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_299_v83_).cf4) == (_dafny.SeqWithoutIsStrInference([2]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_300_i7_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_317_v96_).cf0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_350_v129_) == (_dafny.MultiSet([True, True, True, True, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_351_v130_) == (_dafny.Map({_dafny.CodePoint('p'): _dafny.MultiSet([True, True, True, True, False])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_352_v132_).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_352_v132_).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_354_v133_).cf6).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_354_v133_).cf6).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_354_v133_).cf6.f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[0]).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[0]).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_355_v134_)[0].f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[1]).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[1]).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_355_v134_)[1].f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[2]).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[2]).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_355_v134_)[2].f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[3]).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[3]).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_355_v134_)[3].f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[4]).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[4]).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_355_v134_)[4].f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[5]).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[5]).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_355_v134_)[5].f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[6]).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[6]).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_355_v134_)[6].f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[7]).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[7]).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_355_v134_)[7].f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[8]).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[8]).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_355_v134_)[8].f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[9]).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[9]).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_355_v134_)[9].f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[10]).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[10]).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_355_v134_)[10].f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[11]).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[11]).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_355_v134_)[11].f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[12]).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[12]).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_355_v134_)[12].f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[13]).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[13]).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_355_v134_)[13].f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[14]).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[14]).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_355_v134_)[14].f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[15]).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[15]).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_355_v134_)[15].f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[16]).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[16]).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_355_v134_)[16].f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[17]).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[17]).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_355_v134_)[17].f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[18]).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[18]).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_355_v134_)[18].f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[19]).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[19]).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_355_v134_)[19].f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[20]).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[20]).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_355_v134_)[20].f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[21]).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[21]).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_355_v134_)[21].f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[22]).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[22]).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_355_v134_)[22].f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[23]).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[23]).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_355_v134_)[23].f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[24]).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[24]).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_355_v134_)[24].f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[25]).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[25]).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_355_v134_)[25].f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[26]).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[26]).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_355_v134_)[26].f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[27]).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[27]).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_355_v134_)[27].f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[28]).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_355_v134_)[28]).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_355_v134_)[28].f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_356_v135_)[0]).cf1) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([False])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_356_v135_)[0]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_356_v135_)[1]).cf1) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([False, False]), _dafny.SeqWithoutIsStrInference([False, False]), _dafny.SeqWithoutIsStrInference([False])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_356_v135_)[1]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_356_v135_)[2]).cf1) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([False, False]), _dafny.SeqWithoutIsStrInference([False, False]), _dafny.SeqWithoutIsStrInference([False])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_356_v135_)[2]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_356_v135_)[3]).cf1) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([False, False]), _dafny.SeqWithoutIsStrInference([False, False]), _dafny.SeqWithoutIsStrInference([False])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_356_v135_)[3]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_361_v136_) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_362_v137_) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([False])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_363_v138_).cf1) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([False])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_363_v138_).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_364_v139_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_365_v140_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_366_v141_) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([2])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_367_v142_).cf11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_367_v142_).cf12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(_dafny.MultiSet({}), int(0))
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
        return lambda: D1_DC4()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D1_DC5)

class D1_DC4(D1, NamedTuple('DC4', [])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4)
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC3(D1, NamedTuple('DC3', [('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC5(D1, NamedTuple('DC5', [('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC7(int(0), None)
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

class D2_DC7(D2, NamedTuple('DC7', [('cf7', Any), ('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf7 == __o.cf7 and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC6(D2, NamedTuple('DC6', [('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf6 == __o.cf6
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
        return lambda: D3_DC10(_dafny.CodePoint('D'), False)
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

class D3_DC11(D3, NamedTuple('DC11', [('cf13', Any), ('cf14', Any), ('cf15', Any), ('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC11({_dafny.string_of(self.cf13)}, {_dafny.string_of(self.cf14)}, {_dafny.string_of(self.cf15)}, {_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC11) and self.cf13 == __o.cf13 and self.cf14 == __o.cf14 and self.cf15 == __o.cf15 and self.cf16 == __o.cf16
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
        return lambda: D4_DC13(int(0), int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D4_DC13)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D4_DC14)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D4_DC12)

class D4_DC13(D4, NamedTuple('DC13', [('cf18', Any), ('cf19', Any), ('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC13({_dafny.string_of(self.cf18)}, {_dafny.string_of(self.cf19)}, {_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC13) and self.cf18 == __o.cf18 and self.cf19 == __o.cf19 and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC14(D4, NamedTuple('DC14', [('cf21', Any), ('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC14({_dafny.string_of(self.cf21)}, {_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC14) and self.cf21 == __o.cf21 and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC12(D4, NamedTuple('DC12', [('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC12({_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC12) and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: None
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D5_DC15)

class D5_DC15(D5, NamedTuple('DC15', [('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC15({_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC15) and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC17(False, False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D6_DC17)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D6_DC18)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D6_DC16)

class D6_DC17(D6, NamedTuple('DC17', [('cf25', Any), ('cf26', Any), ('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC17({_dafny.string_of(self.cf25)}, {_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC17) and self.cf25 == __o.cf25 and self.cf26 == __o.cf26 and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC18(D6, NamedTuple('DC18', [('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC18({_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC18) and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC16(D6, NamedTuple('DC16', [('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC16({_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC16) and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC20(False, False, None)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D7_DC20)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D7_DC19)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D7_DC21)

class D7_DC20(D7, NamedTuple('DC20', [('cf30', Any), ('cf31', Any), ('cf32', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC20({_dafny.string_of(self.cf30)}, {_dafny.string_of(self.cf31)}, {_dafny.string_of(self.cf32)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC20) and self.cf30 == __o.cf30 and self.cf31 == __o.cf31 and self.cf32 == __o.cf32
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC19(D7, NamedTuple('DC19', [('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC19({_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC19) and self.cf29 == __o.cf29
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
        return lambda: D8_DC23(False, False, None)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D8_DC23)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D8_DC24)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D8_DC25)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D8_DC22)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D8_DC26)

class D8_DC23(D8, NamedTuple('DC23', [('cf35', Any), ('cf36', Any), ('cf37', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC23({_dafny.string_of(self.cf35)}, {_dafny.string_of(self.cf36)}, {_dafny.string_of(self.cf37)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC23) and self.cf35 == __o.cf35 and self.cf36 == __o.cf36 and self.cf37 == __o.cf37
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC24(D8, NamedTuple('DC24', [('cf38', Any), ('cf39', Any), ('cf40', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC24({_dafny.string_of(self.cf38)}, {_dafny.string_of(self.cf39)}, {_dafny.string_of(self.cf40)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC24) and self.cf38 == __o.cf38 and self.cf39 == __o.cf39 and self.cf40 == __o.cf40
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC25(D8, NamedTuple('DC25', [('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC25({_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC25) and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC22(D8, NamedTuple('DC22', [('cf34', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC22({_dafny.string_of(self.cf34)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC22) and self.cf34 == __o.cf34
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC26(D8, NamedTuple('DC26', [('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC26({_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC26) and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC28(False, int(0), _dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D9_DC28)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D9_DC29)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D9_DC30)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D9_DC27)

class D9_DC28(D9, NamedTuple('DC28', [('cf44', Any), ('cf45', Any), ('cf46', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC28({_dafny.string_of(self.cf44)}, {_dafny.string_of(self.cf45)}, {_dafny.string_of(self.cf46)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC28) and self.cf44 == __o.cf44 and self.cf45 == __o.cf45 and self.cf46 == __o.cf46
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC29(D9, NamedTuple('DC29', [('cf47', Any), ('cf48', Any), ('cf49', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC29({_dafny.string_of(self.cf47)}, {_dafny.string_of(self.cf48)}, {_dafny.string_of(self.cf49)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC29) and self.cf47 == __o.cf47 and self.cf48 == __o.cf48 and self.cf49 == __o.cf49
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC30(D9, NamedTuple('DC30', [('cf50', Any), ('cf51', Any), ('cf52', Any), ('cf53', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC30({_dafny.string_of(self.cf50)}, {_dafny.string_of(self.cf51)}, {_dafny.string_of(self.cf52)}, {_dafny.string_of(self.cf53)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC30) and self.cf50 == __o.cf50 and self.cf51 == __o.cf51 and self.cf52 == __o.cf52 and self.cf53 == __o.cf53
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC27(D9, NamedTuple('DC27', [('cf43', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC27({_dafny.string_of(self.cf43)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC27) and self.cf43 == __o.cf43
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D10_DC31)

class D10_DC31(D10, NamedTuple('DC31', [('cf54', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC31({self.cf54.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC31) and self.cf54 == __o.cf54
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC33(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D11_DC33)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D11_DC32)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D11_DC34)

class D11_DC33(D11, NamedTuple('DC33', [('cf56', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC33({_dafny.string_of(self.cf56)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC33) and self.cf56 == __o.cf56
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC32(D11, NamedTuple('DC32', [('cf55', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC32({_dafny.string_of(self.cf55)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC32) and self.cf55 == __o.cf55
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC34(D11, NamedTuple('DC34', [('cf57', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC34({_dafny.string_of(self.cf57)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC34) and self.cf57 == __o.cf57
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC36(int(0), int(0), False, int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D12_DC36)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D12_DC37)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D12_DC35)

class D12_DC36(D12, NamedTuple('DC36', [('cf59', Any), ('cf60', Any), ('cf61', Any), ('cf62', Any), ('cf63', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC36({_dafny.string_of(self.cf59)}, {_dafny.string_of(self.cf60)}, {_dafny.string_of(self.cf61)}, {_dafny.string_of(self.cf62)}, {_dafny.string_of(self.cf63)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC36) and self.cf59 == __o.cf59 and self.cf60 == __o.cf60 and self.cf61 == __o.cf61 and self.cf62 == __o.cf62 and self.cf63 == __o.cf63
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC37(D12, NamedTuple('DC37', [('cf64', Any), ('cf65', Any), ('cf66', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC37({_dafny.string_of(self.cf64)}, {_dafny.string_of(self.cf65)}, {_dafny.string_of(self.cf66)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC37) and self.cf64 == __o.cf64 and self.cf65 == __o.cf65 and self.cf66 == __o.cf66
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC35(D12, NamedTuple('DC35', [('cf58', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC35({_dafny.string_of(self.cf58)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC35) and self.cf58 == __o.cf58
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    @property
    def f7(self):
        return self._f7
    @f7.setter
    def f7(self, value):
        self._f7 = value

class GlobalState:
    def  __init__(self):
        self.f3: int = int(0)
        self.f4: _dafny.Map = _dafny.Map({})
        self.f6: int = int(0)
        self._f0: bool = False
        self._f1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f2: bool = False
        self._f5: _dafny.MultiSet = _dafny.MultiSet({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6):
        (self)._f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2
        (self).f3 = f3
        (self).f4 = f4
        (self)._f5 = f5
        (self).f6 = f6

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
    def f5(self):
        return self._f5

class C0(T0):
    def  __init__(self):
        self._f7: bool = False
        self._f8: int = int(0)
        self._f9: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    @property
    def f7(self):
        return self._f7
    @f7.setter
    def f7(self, value):
        self._f7 = value
    def ctor__(self, f8, f9, f7):
        (self)._f8 = f8
        (self)._f9 = f9
        (self).f7 = f7

    def fm4(self, globalState):
        return (-635) - ((self).f9)

    def fm5(self, p0, p1, p2, p3, globalState):
        return len(_dafny.Map({(self.f7 if self.f7 else self.f7): self.f7}))

    @property
    def f8(self):
        return self._f8
    @property
    def f9(self):
        return self._f9
