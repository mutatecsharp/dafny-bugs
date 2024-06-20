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
    def fm2(globalState):
        return _dafny.MultiSet([(910) + (260)])

    @staticmethod
    def fm3(p0, p1, p2, globalState):
        return (len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uirae"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))))) * (21)

    @staticmethod
    def fm4(p0, p1, p2, p3, globalState):
        return (_dafny.Map({False: len(_dafny.Set({True, True}))})) | (_dafny.Map({False: -93}))

    @staticmethod
    def fm5(globalState):
        return _dafny.Map({False: (D9_DC26(_dafny.MultiSet([-657, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_0_i0_ in range(default__.abs(193))]))]))).cf38})

    @staticmethod
    def fm6(p0, p1, p2, globalState):
        return _dafny.Map({(56) - (len(_dafny.SeqWithoutIsStrInference([True, True]))): (len((D10_DC28(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_1_i0_ in range(default__.abs(384))]))}))).cf39) if True else (0) - ((0) - (len(_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y"))})))))})

    @staticmethod
    def fm7(p0, p1, globalState):
        def iife0_():
            coll0_ = _dafny.Set()
            compr_0_: _dafny.Seq
            for compr_0_ in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_2_i0_ in range(default__.abs(76))])])).Elements:
                d_3_v0_: _dafny.Seq = compr_0_
                if (d_3_v0_) in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_2_i0_ in range(default__.abs(76))])])):
                    coll0_ = coll0_.union(_dafny.Set([d_3_v0_]))
            return _dafny.Set(coll0_)
        return (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qaoo"))})).intersection((iife0_()
        ).intersection(_dafny.Set({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_4_i1_ in range(default__.abs(-523))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "flq"))})))

    @staticmethod
    def fm8(p0, p1, p2, p3, globalState):
        return not(((0) - ((478) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d")))))) != ((_dafny.MultiSet([128, 805])).cardinality))

    @staticmethod
    def fm9(p0, p1, p2, p3, globalState):
        def iife1_():
            coll1_ = _dafny.Set()
            compr_1_: int
            for compr_1_ in (_dafny.SeqWithoutIsStrInference([989, (_dafny.MultiSet([_dafny.CodePoint('s')])).cardinality])).Elements:
                d_5_v0_: int = compr_1_
                if (d_5_v0_) in (_dafny.SeqWithoutIsStrInference([989, (_dafny.MultiSet([_dafny.CodePoint('s')])).cardinality])):
                    coll1_ = coll1_.union(_dafny.Set([default__.safeDivisionInt(d_5_v0_, -206)]))
            return _dafny.Set(coll1_)
        def iife2_():
            coll2_ = _dafny.Map()
            compr_2_: int
            for compr_2_ in _dafny.IntegerRange(832, 101):
                d_6_v1_: int = compr_2_
                if ((832) <= (d_6_v1_)) and ((d_6_v1_) < (101)):
                    coll2_[default__.safeDivisionInt(d_6_v1_, 228)] = not(True)
            return _dafny.Map(coll2_)
        return D1_DC2(_dafny.CodePoint('i'), 216, default__.safeDivisionInt((0) - (len(iife1_()
)), len(_dafny.Map({_dafny.CodePoint('n'): D10_DC29((0) - ((0) - (len(_dafny.SeqWithoutIsStrInference([False])))), len(iife2_()
))}))))

    @staticmethod
    def fm10(globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sh"))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pcvumw"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yl"))))

    @staticmethod
    def m1(p0, p1, globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r1: C0 = None
        r2: _dafny.Seq = _dafny.Seq({})
        r3: bool = False
        if p1:
            d_7_v0_: int
            d_7_v0_ = 845
            d_7_v0_ = d_7_v0_
            d_8_v1_: C1
            nw0_ = C1()
            nw0_.ctor__()
            d_8_v1_ = nw0_
            d_9_v3_: _dafny.Seq
            d_9_v3_ = _dafny.SeqWithoutIsStrInference([p1])
            d_10_v4_: _dafny.Map
            d_10_v4_ = _dafny.Map({d_7_v0_: d_9_v3_})
            d_11_v5_: _dafny.Map
            def iife3_():
                coll3_ = _dafny.Map()
                compr_3_: int
                for compr_3_ in (d_10_v4_).keys.Elements:
                    d_12_v2_: int = compr_3_
                    if (d_12_v2_) in (d_10_v4_):
                        coll3_[(d_12_v2_) - (d_7_v0_)] = d_7_v0_
                return _dafny.Map(coll3_)
            d_11_v5_ = _dafny.Map({d_7_v0_: default__.fm8(d_7_v0_, iife3_()
            , d_7_v0_, True, globalState)})
            d_13_v6_: _dafny.Map
            d_13_v6_ = _dafny.Map({d_8_v1_: d_11_v5_})
            d_14_v7_: D8
            d_14_v7_ = D8_DC22(_dafny.Map({d_8_v1_: _dafny.Map({d_7_v0_: p1})}))
            d_15_v8_: D8
            d_15_v8_ = D8_DC22((d_14_v7_).cf29)
            d_13_v6_ = ((d_13_v6_) | ((d_14_v7_).cf29)) | ((_dafny.Map({d_8_v1_: d_11_v5_})).set(d_8_v1_, d_11_v5_))
            d_16_v9_: _dafny.Array
            nw1_ = _dafny.Array(_dafny.Seq({}), 23)
            d_16_v9_ = nw1_
            d_17_v10_: _dafny.Seq
            d_17_v10_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_7_v0_])])
            index0_ = default__.safeIndex(445, (d_16_v9_).length(0))
            (d_16_v9_)[index0_] = d_17_v10_
            d_18_v11_: _dafny.Map
            d_18_v11_ = _dafny.Map({d_7_v0_: d_7_v0_})
            d_19_v12_: _dafny.Set
            d_19_v12_ = _dafny.Set({p1, p1})
            d_20_v13_: _dafny.Map
            d_20_v13_ = _dafny.Map({p1: default__.fm2(globalState)})
            d_21_v14_: bool
            d_21_v14_ = False
            d_22_v15_: C0
            nw2_ = C0()
            nw2_.ctor__(d_20_v13_, d_7_v0_, d_21_v14_)
            d_22_v15_ = nw2_
            d_23_v16_: _dafny.Map
            d_23_v16_ = _dafny.Map({p1: 303})
            d_24_v17_: _dafny.Map
            d_24_v17_ = _dafny.Map({True: d_23_v16_})
            d_25_v18_: _dafny.MultiSet
            d_25_v18_ = _dafny.MultiSet([d_7_v0_])
            index1_ = default__.safeIndex(445, (d_16_v9_).length(0))
            rhs0_ = default__.fm8(d_7_v0_, (d_18_v11_).set(len(d_19_v12_), d_7_v0_), d_7_v0_, p1, globalState)
            rhs1_ = (0) - ((d_7_v0_ if (p1) and (True) else d_7_v0_))
            rhs2_ = d_22_v15_
            rhs3_ = default__.safeDivisionInt(default__.fm3(p0, p1, p1, globalState), (0) - ((d_7_v0_) * (547)))
            rhs4_ = (d_17_v10_) + ((d_17_v10_) + ((d_17_v10_).set(default__.safeIndex(len(d_24_v17_), len(d_17_v10_)), (d_25_v18_).set(d_7_v0_, default__.abs(d_7_v0_)))))
            lhs0_ = d_16_v9_
            lhs1_ = default__.safeIndex(445, (d_16_v9_).length(0))
            r3 = rhs0_
            d_7_v0_ = rhs1_
            r1 = rhs2_
            d_7_v0_ = rhs3_
            lhs0_[lhs1_] = rhs4_
            d_26_v19_: _dafny.Seq
            d_26_v19_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pbwt"))
            rhs5_ = (0) - ((d_7_v0_ if p1 else d_7_v0_))
            rhs6_ = d_7_v0_
            rhs7_ = (d_7_v0_ if p1 else (d_22_v15_).fm0(d_7_v0_, p0, p1, d_7_v0_, globalState))
            rhs8_ = d_26_v19_
            d_7_v0_ = rhs5_
            d_7_v0_ = rhs6_
            d_7_v0_ = rhs7_
            r0 = rhs8_
            d_23_v16_ = (d_23_v16_).set(True, len(d_26_v19_))
        elif True:
            d_27_v20_: _dafny.Array
            nw3_ = _dafny.Array(int(0), 12)
            d_27_v20_ = nw3_
            d_28_v21_: _dafny.Seq
            d_28_v21_ = _dafny.SeqWithoutIsStrInference([d_27_v20_, d_27_v20_, d_27_v20_])
            d_28_v21_ = d_28_v21_
            d_29_v22_: int
            d_29_v22_ = 736
            d_29_v22_ = (d_29_v22_) - (default__.safeModuloInt(-887, d_29_v22_))
            d_30_v23_: _dafny.MultiSet
            d_30_v23_ = _dafny.MultiSet([d_29_v22_])
            d_31_v24_: _dafny.Map
            d_31_v24_ = _dafny.Map({False: d_30_v23_})
            d_32_v25_: bool
            d_32_v25_ = p1
            d_33_v26_: T0
            nw4_ = C0()
            nw4_.ctor__(d_31_v24_, len(_dafny.SeqWithoutIsStrInference([not(p1)])), d_32_v25_)
            d_33_v26_ = nw4_
            d_34_v27_: _dafny.MultiSet
            d_34_v27_ = _dafny.MultiSet([d_33_v26_])
            d_35_v28_: _dafny.MultiSet
            d_35_v28_ = _dafny.MultiSet([d_34_v27_, d_34_v27_])
            d_35_v28_ = (_dafny.MultiSet([d_34_v27_])) - (d_35_v28_)
            d_36_v29_: _dafny.Map
            d_36_v29_ = _dafny.Map({(d_33_v26_).f6: (d_33_v26_).f6})
            def iife4_():
                coll4_ = _dafny.Map()
                compr_4_: int
                for compr_4_ in _dafny.IntegerRange(454, -189):
                    d_37_v30_: int = compr_4_
                    if ((454) <= (d_37_v30_)) and ((d_37_v30_) < (-189)):
                        coll4_[(d_37_v30_) + ((_dafny.MultiSet([_dafny.Set({p1})])).cardinality)] = (0) - (d_29_v22_)
                return _dafny.Map(coll4_)
            if not (default__.fm8((0) - (d_29_v22_), d_36_v29_, len(iife4_()
            ), p1, globalState)) or (p1):
                d_38_v31_: C0
                nw5_ = C0()
                nw5_.ctor__(_dafny.Map({not(True): d_30_v23_}), d_29_v22_, p1)
                d_38_v31_ = nw5_
                d_29_v22_ = d_29_v22_
                d_39_v32_: _dafny.Set
                d_39_v32_ = _dafny.Set({p1})
                d_40_v33_: _dafny.Map
                d_40_v33_ = _dafny.Map({d_39_v32_: False})
                d_29_v22_ = (default__.safeModuloInt((d_33_v26_).f6, (d_33_v26_).f6)) - (len(d_40_v33_))
                d_41_v34_: D6
                d_41_v34_ = D6_DC15(d_38_v31_)
                d_42_v35_: _dafny.Seq
                d_42_v35_ = _dafny.SeqWithoutIsStrInference([d_38_v31_, d_38_v31_, d_38_v31_, d_38_v31_])
                d_43_v36_: _dafny.Seq
                d_43_v36_ = _dafny.SeqWithoutIsStrInference([p1])
                d_44_v37_: _dafny.Map
                d_44_v37_ = _dafny.Map({(d_43_v36_)[default__.safeIndex(-721, len(d_43_v36_))]: d_29_v22_})
                d_45_v38_: _dafny.Seq
                d_45_v38_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "njpmb"))
                d_46_v39_: _dafny.Array
                nw6_ = _dafny.Array(None, 10)
                nw6_[int(0)] = d_38_v31_
                nw6_[int(1)] = d_38_v31_
                nw6_[int(2)] = d_38_v31_
                nw6_[int(3)] = d_38_v31_
                nw6_[int(4)] = (d_41_v34_).cf21
                nw6_[int(5)] = d_38_v31_
                nw6_[int(6)] = d_38_v31_
                nw6_[int(7)] = (d_42_v35_)[default__.safeIndex(len((d_44_v37_).set(p1, len(d_45_v38_))), len(d_42_v35_))]
                nw6_[int(8)] = d_38_v31_
                nw6_[int(9)] = d_38_v31_
                d_46_v39_ = nw6_
                d_46_v39_ = d_46_v39_
                d_47_v40_: _dafny.Map
                d_47_v40_ = _dafny.Map({p1: p1})
                d_47_v40_ = (d_47_v40_).set(p1, not((d_45_v38_) < (d_45_v38_)))
            elif True:
                d_48_v41_: _dafny.Array
                nw7_ = _dafny.Array(False, 14)
                d_48_v41_ = nw7_
                index2_ = default__.safeIndex(831, (d_48_v41_).length(0))
                (d_48_v41_)[index2_] = p1
                d_49_v42_: _dafny.Map
                d_49_v42_ = _dafny.Map({d_29_v22_: p1})
                d_50_v43_: _dafny.Seq
                d_50_v43_ = _dafny.SeqWithoutIsStrInference([((d_49_v42_)[(0) - ((0) - (d_29_v22_))] if ((0) - ((0) - (d_29_v22_))) in (d_49_v42_) else p1)])
                index3_ = default__.safeIndex(831, (d_48_v41_).length(0))
                (d_48_v41_)[index3_] = ((d_33_v26_).f6) < (len(d_50_v43_))
                d_29_v22_ = (0) - (d_29_v22_)
                d_51_v44_: C1
                nw8_ = C1()
                nw8_.ctor__()
                d_51_v44_ = nw8_
                d_52_v45_: _dafny.Map
                d_52_v45_ = _dafny.Map({d_29_v22_: d_51_v44_})
                d_53_v46_: D1
                d_53_v46_ = D1_DC2(_dafny.CodePoint('u'), d_29_v22_, -318)
                d_54_v47_: _dafny.Set
                d_54_v47_ = _dafny.Set({len(d_50_v43_)})
                d_55_v48_: _dafny.Set
                d_55_v48_ = _dafny.Set({(D1_DC2((d_53_v46_).cf2, (d_33_v26_).f6, (0) - (len(d_54_v47_)))).cf4})
                d_52_v45_ = (d_52_v45_).set((d_33_v26_).f6, (d_51_v44_ if not(default__.fm8((d_33_v26_).f6, _dafny.Map({(d_33_v26_).f6: (d_33_v26_).f6}), len(d_54_v47_), p1, globalState)) else d_51_v44_))
                index4_ = default__.safeIndex(831, (d_48_v41_).length(0))
                (d_48_v41_)[index4_] = p1
                d_56_v49_: C1
                nw9_ = C1()
                nw9_.ctor__()
                d_56_v49_ = nw9_
            d_57_v50_: C1
            nw10_ = C1()
            nw10_.ctor__()
            d_57_v50_ = nw10_
        d_58_v51_: _dafny.Array
        def lambda0_(d_59_p1_):
            def lambda1_(d_60_i0_):
                return default__.safeModuloInt(d_60_i0_, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ovbwu"))).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_59_p1_])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ovbwu")))), _dafny.CodePoint('n'))))

            return lambda1_

        init0_ = lambda0_(p1)
        nw11_ = _dafny.Array(None, 19)
        for i0_0_ in range(nw11_.length(0)):
            nw11_[i0_0_] = init0_(i0_0_)
        d_58_v51_ = nw11_
        d_61_v52_: int
        d_61_v52_ = 926
        index5_ = default__.safeIndex(847, (d_58_v51_).length(0))
        (d_58_v51_)[index5_] = (0) - (d_61_v52_)
        index6_ = default__.safeIndex(847, (d_58_v51_).length(0))
        (d_58_v51_)[index6_] = d_61_v52_
        d_62_v53_: _dafny.Map
        d_62_v53_ = _dafny.Map({p1: _dafny.MultiSet([d_61_v52_])})
        d_63_v54_: bool
        d_63_v54_ = p1
        d_64_v55_: T0
        nw12_ = C0()
        nw12_.ctor__(d_62_v53_, (d_58_v51_)[default__.safeIndex(847, (d_58_v51_).length(0))], d_63_v54_)
        d_64_v55_ = nw12_
        d_65_v56_: D7
        d_65_v56_ = D7_DC21(p1, d_64_v55_)
        if (d_65_v56_).cf27:
            d_61_v52_ = (0) - (((d_64_v55_).f6) - ((d_58_v51_)[default__.safeIndex(847, (d_58_v51_).length(0))]))
            d_66_v57_: _dafny.Seq
            d_66_v57_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fqqxnh"))
            d_67_v58_: _dafny.MultiSet
            d_67_v58_ = _dafny.MultiSet([(d_64_v55_).f6])
            d_68_v59_: _dafny.MultiSet
            d_68_v59_ = _dafny.MultiSet([d_58_v51_, d_58_v51_])
            d_69_v60_: D3
            d_69_v60_ = D3_DC6(d_68_v59_)
            d_70_v61_: C1
            nw13_ = C1()
            nw13_.ctor__()
            d_70_v61_ = nw13_
            d_71_v62_: _dafny.Map
            d_71_v62_ = _dafny.Map({d_69_v60_: d_70_v61_})
            d_72_v63_: _dafny.Seq
            d_72_v63_ = _dafny.SeqWithoutIsStrInference([(d_64_v55_).f6, ((d_64_v55_).f6) * ((d_64_v55_).f6), default__.safeDivisionInt((d_64_v55_).f6, len(d_66_v57_)), ((d_67_v58_) - (d_67_v58_)).cardinality, ((d_58_v51_)[default__.safeIndex(847, (d_58_v51_).length(0))]) * (len(d_71_v62_))])
            index7_ = default__.safeIndex(847, (d_58_v51_).length(0))
            (d_58_v51_)[index7_] = (0) - ((0) - ((d_72_v63_)[default__.safeIndex((0) - ((0) - (d_61_v52_)), len(d_72_v63_))]))
            r3 = p1
            d_73_v64_: C0
            nw14_ = C0()
            nw14_.ctor__(default__.fm5(globalState), d_61_v52_, d_63_v54_)
            d_73_v64_ = nw14_
            d_74_v65_: _dafny.Map
            d_74_v65_ = _dafny.Map({d_73_v64_: not(False)})
            d_74_v65_ = (d_74_v65_).set(d_73_v64_, p1)
            d_75_v66_: _dafny.Set
            d_75_v66_ = _dafny.Set({d_72_v63_})
            index8_ = default__.safeIndex(847, (d_58_v51_).length(0))
            rhs9_ = -252
            rhs10_ = (d_66_v57_) + ((d_66_v57_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jnqvr"))))
            rhs11_ = d_61_v52_
            rhs12_ = _dafny.Set({d_72_v63_, d_72_v63_})
            lhs2_ = d_58_v51_
            lhs3_ = default__.safeIndex(847, (d_58_v51_).length(0))
            lhs2_[lhs3_] = rhs9_
            d_66_v57_ = rhs10_
            d_61_v52_ = rhs11_
            d_75_v66_ = rhs12_
        elif True:
            d_76_v67_: D5
            d_76_v67_ = D5_DC12((d_61_v52_) >= ((d_58_v51_)[default__.safeIndex(847, (d_58_v51_).length(0))]))
            d_76_v67_ = d_76_v67_
            d_77_v68_: _dafny.Seq
            d_77_v68_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ibjngh"))
            r3 = ((d_77_v68_) + (default__.fm10(globalState))) <= (d_77_v68_)
            d_78_v69_: _dafny.Seq
            d_78_v69_ = _dafny.SeqWithoutIsStrInference([len(d_77_v68_), len(d_77_v68_), 751, 427])
            d_79_v70_: _dafny.Map
            d_79_v70_ = _dafny.Map({not(False): d_64_v55_})
            d_80_v71_: _dafny.Map
            d_80_v71_ = _dafny.Map({d_78_v69_: d_79_v70_})
            d_81_v72_: _dafny.Map
            d_81_v72_ = _dafny.Map({d_61_v52_: (d_58_v51_)[default__.safeIndex(847, (d_58_v51_).length(0))]})
            r3 = ((((d_80_v71_)[d_78_v69_] if (d_78_v69_) in (d_80_v71_) else d_79_v70_)).set(default__.fm8((d_64_v55_).f6, d_81_v72_, 902, p1, globalState), d_64_v55_)) not in (_dafny.SeqWithoutIsStrInference([d_79_v70_]))
            r3 = (((d_58_v51_)[default__.safeIndex(847, (d_58_v51_).length(0))]) - ((d_64_v55_).f6)) == ((d_64_v55_).f6)
            d_82_v73_: D1
            d_82_v73_ = D1_DC2(p0, d_61_v52_, (d_58_v51_)[default__.safeIndex(847, (d_58_v51_).length(0))])
            r3 = (p1 if (d_82_v73_) not in (_dafny.SeqWithoutIsStrInference([D1_DC2(p0, d_61_v52_, (d_58_v51_)[default__.safeIndex(847, (d_58_v51_).length(0))]), d_82_v73_])) else (557) == ((d_64_v55_).f6))
        d_83_v74_: C0
        nw15_ = C0()
        nw15_.ctor__((d_62_v53_) | (d_62_v53_), d_61_v52_, d_63_v54_)
        d_83_v74_ = nw15_
        if p1:
            d_84_v75_: _dafny.Seq
            d_84_v75_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b"))
            r0 = d_84_v75_
            d_85_v76_: C1
            nw16_ = C1()
            nw16_.ctor__()
            d_85_v76_ = nw16_
            d_86_v77_: int
            out0_: int
            out0_ = (d_85_v76_).m0(p1, p0, p1, globalState)
            d_86_v77_ = out0_
            d_87_v78_: _dafny.Seq
            d_87_v78_ = _dafny.SeqWithoutIsStrInference([d_84_v75_])
            rhs13_ = d_87_v78_
            rhs14_ = d_83_v74_
            rhs15_ = ((d_58_v51_)[default__.safeIndex(847, (d_58_v51_).length(0))]) + (((d_58_v51_)[default__.safeIndex(847, (d_58_v51_).length(0))]) * (d_86_v77_))
            rhs16_ = ((405) + (375)) <= ((d_64_v55_).f6)
            d_87_v78_ = rhs13_
            d_83_v74_ = rhs14_
            d_86_v77_ = rhs15_
            r3 = rhs16_
            d_88_v79_: _dafny.Map
            d_88_v79_ = _dafny.Map({d_64_v55_: d_64_v55_})
            d_64_v55_ = (((d_88_v79_)[d_64_v55_] if (d_64_v55_) in (d_88_v79_) else d_64_v55_) if p1 else d_64_v55_)
        elif True:
            d_89_v80_: _dafny.Array
            def lambda2_(d_90_p1_):
                def lambda3_(d_91_i1_):
                    return d_90_p1_

                return lambda3_

            init1_ = lambda2_(p1)
            nw17_ = _dafny.Array(None, 4)
            for i0_1_ in range(nw17_.length(0)):
                nw17_[i0_1_] = init1_(i0_1_)
            d_89_v80_ = nw17_
            index9_ = default__.safeIndex(804, (d_89_v80_).length(0))
            (d_89_v80_)[index9_] = p1
            index10_ = default__.safeIndex(804, (d_89_v80_).length(0))
            (d_89_v80_)[index10_] = not(p1)
            d_92_v81_: _dafny.Set
            d_92_v81_ = _dafny.Set({d_89_v80_})
            d_92_v81_ = (d_92_v81_).intersection(_dafny.Set({d_89_v80_, d_89_v80_, d_89_v80_}))
            d_93_v82_: str
            d_93_v82_ = _dafny.CodePoint('a')
            d_93_v82_ = p0
            d_61_v52_ = (d_58_v51_)[default__.safeIndex(847, (d_58_v51_).length(0))]
            if (d_89_v80_)[default__.safeIndex(804, (d_89_v80_).length(0))]:
                d_94_v83_: _dafny.Set
                d_94_v83_ = _dafny.Set({d_93_v82_})
                index11_ = default__.safeIndex(804, (d_89_v80_).length(0))
                index12_ = default__.safeIndex(804, (d_89_v80_).length(0))
                rhs17_ = (d_94_v83_).issubset(d_94_v83_)
                rhs18_ = (False if (d_89_v80_)[default__.safeIndex(804, (d_89_v80_).length(0))] else (d_89_v80_)[default__.safeIndex(804, (d_89_v80_).length(0))])
                lhs4_ = d_89_v80_
                lhs5_ = default__.safeIndex(804, (d_89_v80_).length(0))
                lhs6_ = d_89_v80_
                lhs7_ = default__.safeIndex(804, (d_89_v80_).length(0))
                lhs4_[lhs5_] = rhs17_
                lhs6_[lhs7_] = rhs18_
                d_95_v84_: _dafny.Seq
                d_95_v84_ = _dafny.SeqWithoutIsStrInference([p1, True, p1, p1])
                d_96_v85_: _dafny.MultiSet
                d_96_v85_ = _dafny.MultiSet([(d_64_v55_).f6])
                d_97_v86_: _dafny.Seq
                d_97_v86_ = _dafny.SeqWithoutIsStrInference([(d_96_v85_).cardinality])
                d_98_v87_: _dafny.Set
                d_98_v87_ = _dafny.Set({not(p1), (len(_dafny.Map({(d_89_v80_)[default__.safeIndex(804, (d_89_v80_).length(0))]: (d_95_v84_)[default__.safeIndex((d_58_v51_)[default__.safeIndex(847, (d_58_v51_).length(0))], len(d_95_v84_))]}))) in (d_97_v86_), p1})
                (globalState).f5 = d_98_v87_
                d_99_v88_: C1
                nw18_ = C1()
                nw18_.ctor__()
                d_99_v88_ = nw18_
                d_100_v89_: _dafny.Array
                def lambda4_(d_101_v86_):
                    def lambda5_(d_102_i2_):
                        return d_101_v86_

                    return lambda5_

                init2_ = lambda4_(d_97_v86_)
                nw19_ = _dafny.Array(None, 27)
                for i0_2_ in range(nw19_.length(0)):
                    nw19_[i0_2_] = init2_(i0_2_)
                d_100_v89_ = nw19_
                index13_ = default__.safeIndex(862, (d_100_v89_).length(0))
                (d_100_v89_)[index13_] = d_97_v86_
                index14_ = default__.safeIndex(862, (d_100_v89_).length(0))
                (d_100_v89_)[index14_] = _dafny.SeqWithoutIsStrInference([len(_dafny.Set({d_61_v52_})) for d_103_i3_ in range(default__.abs(937))])
                d_104_v90_: D4
                d_104_v90_ = D4_DC9(d_89_v80_)
                d_105_v91_: _dafny.Map
                d_105_v91_ = _dafny.Map({d_104_v90_: d_93_v82_})
                d_93_v82_ = ((d_105_v91_)[D4_DC9(d_89_v80_)] if (D4_DC9(d_89_v80_)) in (d_105_v91_) else p0)
            elif True:
                d_106_v92_: C0
                nw20_ = C0()
                nw20_.ctor__((d_83_v74_.f8) | (d_83_v74_.f8), d_61_v52_, p1)
                d_106_v92_ = nw20_
                d_107_v93_: C1
                nw21_ = C1()
                nw21_.ctor__()
                d_107_v93_ = nw21_
                d_108_v94_: _dafny.Map
                d_108_v94_ = _dafny.Map({(d_89_v80_)[default__.safeIndex(804, (d_89_v80_).length(0))]: d_107_v93_})
                d_109_v95_: _dafny.MultiSet
                d_109_v95_ = _dafny.MultiSet([p1, (d_89_v80_)[default__.safeIndex(804, (d_89_v80_).length(0))], p1])
                index15_ = default__.safeIndex(804, (d_89_v80_).length(0))
                (d_89_v80_)[index15_] = not ((d_83_v74_).fm1((d_64_v55_).f6, False, len(d_108_v94_), d_109_v95_, globalState)) or (p1)
                d_110_v96_: _dafny.Seq
                d_110_v96_ = _dafny.SeqWithoutIsStrInference([(0) - ((d_64_v55_).f6)])
                d_111_v97_: _dafny.Seq
                d_111_v97_ = _dafny.SeqWithoutIsStrInference([(d_89_v80_)[default__.safeIndex(804, (d_89_v80_).length(0))], (d_110_v96_) == (d_110_v96_), p1])
                index16_ = default__.safeIndex(804, (d_89_v80_).length(0))
                rhs19_ = d_107_v93_
                rhs20_ = not((d_111_v97_)[default__.safeIndex((d_64_v55_).f6, len(d_111_v97_))])
                rhs21_ = (d_89_v80_)[default__.safeIndex(804, (d_89_v80_).length(0))]
                lhs8_ = d_89_v80_
                lhs9_ = default__.safeIndex(804, (d_89_v80_).length(0))
                d_107_v93_ = rhs19_
                lhs8_[lhs9_] = rhs20_
                r3 = rhs21_
                d_112_v98_: int
                out1_: int
                out1_ = (d_107_v93_).m0(not(True), d_93_v82_, (d_64_v55_).f7, globalState)
                d_112_v98_ = out1_
                d_113_v99_: _dafny.Seq
                d_113_v99_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aqwo"))
                r0 = d_113_v99_
        d_114_v100_: _dafny.Array
        def lambda6_(d_115_p1_):
            def lambda7_(d_116_i4_):
                return (_dafny.Set({_dafny.Map({d_115_p1_: d_115_p1_}), _dafny.Map({d_115_p1_: not(d_115_p1_)}), _dafny.Map({d_115_p1_: d_115_p1_})})).intersection(_dafny.Set({_dafny.Map({not(d_115_p1_): d_115_p1_})}))

            return lambda7_

        init3_ = lambda6_(p1)
        nw22_ = _dafny.Array(None, 5)
        for i0_3_ in range(nw22_.length(0)):
            nw22_[i0_3_] = init3_(i0_3_)
        d_114_v100_ = nw22_
        d_117_v101_: _dafny.Set
        d_117_v101_ = _dafny.Set({_dafny.Map({p1: False})})
        index17_ = default__.safeIndex(712, (d_114_v100_).length(0))
        (d_114_v100_)[index17_] = (d_117_v101_ if p1 else d_117_v101_)
        index18_ = default__.safeIndex(712, (d_114_v100_).length(0))
        (d_114_v100_)[index18_] = d_117_v101_
        d_118_v102_: _dafny.Seq
        d_118_v102_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pkeaema"))
        r0 = d_118_v102_
        r1 = d_83_v74_
        def iife5_():
            coll5_ = _dafny.Map()
            compr_5_: int
            for compr_5_ in _dafny.IntegerRange(146, 220):
                d_119_v103_: int = compr_5_
                if ((146) <= (d_119_v103_)) and ((d_119_v103_) < (220)):
                    coll5_[(d_119_v103_) - (len(_dafny.Set({p1})))] = d_61_v52_
            return _dafny.Map(coll5_)
        r2 = _dafny.SeqWithoutIsStrInference([iife5_()
 for d_120_i5_ in range(default__.abs(-574))])
        r3 = p1
        return r0, r1, r2, r3

    @staticmethod
    def Main(noArgsParameter__):
        d_121_v0_: int
        d_121_v0_ = 197
        d_122_v1_: _dafny.MultiSet
        d_122_v1_ = _dafny.MultiSet([d_121_v0_])
        d_123_v2_: _dafny.Array
        nw23_ = _dafny.Array(None, 13)
        nw23_[int(0)] = _dafny.MultiSet([d_121_v0_])
        nw23_[int(1)] = d_122_v1_
        nw23_[int(2)] = d_122_v1_
        nw23_[int(3)] = d_122_v1_
        nw23_[int(4)] = d_122_v1_
        nw23_[int(5)] = _dafny.MultiSet([d_121_v0_, d_121_v0_, 265])
        nw23_[int(6)] = d_122_v1_
        nw23_[int(7)] = d_122_v1_
        nw23_[int(8)] = d_122_v1_
        nw23_[int(9)] = d_122_v1_
        nw23_[int(10)] = d_122_v1_
        nw23_[int(11)] = d_122_v1_
        nw23_[int(12)] = d_122_v1_
        d_123_v2_ = nw23_
        d_124_v3_: _dafny.Array
        nw24_ = _dafny.Array(int(0), 2)
        d_124_v3_ = nw24_
        d_125_v4_: _dafny.Map
        d_125_v4_ = _dafny.Map({d_121_v0_: d_124_v3_})
        d_126_v5_: bool
        d_126_v5_ = False
        d_127_v6_: _dafny.Set
        d_127_v6_ = _dafny.Set({d_126_v5_, d_126_v5_, d_126_v5_, d_126_v5_})
        d_128_globalState_: GlobalState
        nw25_ = GlobalState()
        nw25_.ctor__(d_123_v2_, d_123_v2_, True, d_125_v4_, _dafny.CodePoint('x'), (d_127_v6_) | (d_127_v6_))
        d_128_globalState_ = nw25_
        hi0_ = d_121_v0_
        for d_129_i0_ in range((d_121_v0_) + (302), hi0_):
            d_130_v7_: _dafny.Seq
            d_130_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a"))
            d_131_v8_: _dafny.Seq
            d_131_v8_ = _dafny.SeqWithoutIsStrInference([d_126_v5_, d_126_v5_])
            d_132_v9_: str
            d_132_v9_ = _dafny.CodePoint('b')
            d_133_v10_: _dafny.Set
            d_133_v10_ = _dafny.Set({d_132_v9_})
            d_134_v11_: bool
            d_134_v11_ = d_126_v5_
            d_135_v12_: _dafny.Map
            d_135_v12_ = _dafny.Map({d_132_v9_: False})
            d_136_v14_: _dafny.Array
            nw26_ = _dafny.Array(None, 21)
            nw26_[int(0)] = d_126_v5_
            nw26_[int(1)] = d_126_v5_
            nw26_[int(2)] = d_126_v5_
            nw26_[int(3)] = (d_130_v7_) != ((d_130_v7_).set(default__.safeIndex(len(d_131_v8_), len(d_130_v7_)), d_132_v9_))
            nw26_[int(4)] = d_126_v5_
            nw26_[int(5)] = d_126_v5_
            nw26_[int(6)] = d_126_v5_
            nw26_[int(7)] = (d_121_v0_) == (d_129_i0_)
            nw26_[int(8)] = d_126_v5_
            nw26_[int(9)] = (len(d_133_v10_)) >= (d_129_i0_)
            nw26_[int(10)] = d_126_v5_
            nw26_[int(11)] = d_126_v5_
            nw26_[int(12)] = d_126_v5_
            nw26_[int(13)] = (d_134_v11_)
            nw26_[int(14)] = d_126_v5_
            nw26_[int(15)] = d_126_v5_
            nw26_[int(16)] = (d_126_v5_) not in (_dafny.MultiSet([d_126_v5_]))
            nw26_[int(17)] = d_126_v5_
            def iife6_():
                coll6_ = _dafny.Set()
                compr_6_: str
                for compr_6_ in (d_135_v12_).keys.Elements:
                    d_137_v13_: str = compr_6_
                    if (d_137_v13_) in (d_135_v12_):
                        coll6_ = coll6_.union(_dafny.Set([d_137_v13_]))
                return _dafny.Set(coll6_)
            nw26_[int(18)] = (d_133_v10_).ispropersubset(iife6_()
            )
            nw26_[int(19)] = d_126_v5_
            nw26_[int(20)] = d_126_v5_
            d_136_v14_ = nw26_
            d_138_v15_: _dafny.Set
            d_138_v15_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b"))})
            d_139_v16_: _dafny.Seq
            d_139_v16_ = _dafny.SeqWithoutIsStrInference([-894, (0) - (d_121_v0_)])
            d_140_v17_: _dafny.Set
            d_140_v17_ = _dafny.Set({(d_129_i0_ if (d_134_v11_) else len(d_131_v8_)), default__.safeDivisionInt(d_121_v0_, ((d_122_v1_).set(320, default__.abs((d_139_v16_)[default__.safeIndex(d_121_v0_, len(d_139_v16_))]))).cardinality)})
            d_141_v18_: _dafny.Map
            d_141_v18_ = _dafny.Map({d_139_v16_: d_136_v14_})
            d_142_v19_: _dafny.Seq
            d_142_v19_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({len(d_127_v6_): d_121_v0_})])
            d_143_v20_: _dafny.Seq
            d_143_v20_ = _dafny.SeqWithoutIsStrInference([d_140_v17_])
            rhs22_ = d_126_v5_
            rhs23_ = ((d_141_v18_)[((d_139_v16_).set(default__.safeIndex(-122, len(d_139_v16_)), len((d_142_v19_)[default__.safeIndex(d_121_v0_, len(d_142_v19_))]))) + (d_139_v16_)] if (((d_139_v16_).set(default__.safeIndex(-122, len(d_139_v16_)), len((d_142_v19_)[default__.safeIndex(d_121_v0_, len(d_142_v19_))]))) + (d_139_v16_)) in (d_141_v18_) else d_136_v14_)
            rhs24_ = d_138_v15_
            rhs25_ = (d_140_v17_).intersection((d_143_v20_)[default__.safeIndex(d_129_i0_, len(d_143_v20_))])
            d_126_v5_ = rhs22_
            d_136_v14_ = rhs23_
            d_138_v15_ = rhs24_
            d_140_v17_ = rhs25_
            d_144_v21_: C1
            nw27_ = C1()
            nw27_.ctor__()
            d_144_v21_ = nw27_
            d_126_v5_ = d_126_v5_
            d_121_v0_ = d_121_v0_
        hi1_ = d_121_v0_
        for d_145_i1_ in range(71, hi1_):
            d_146_v22_: str
            d_146_v22_ = _dafny.CodePoint('r')
            d_147_v23_: _dafny.Seq
            d_147_v23_ = _dafny.SeqWithoutIsStrInference([d_145_i1_])
            d_148_v24_: D1
            d_148_v24_ = D1_DC2(d_146_v22_, d_121_v0_, len(d_147_v23_))
            d_149_v25_: D3
            d_149_v25_ = D3_DC7(d_148_v24_, d_126_v5_)
            d_126_v5_ = (d_149_v25_).cf15
            d_150_v26_: _dafny.Seq
            d_150_v26_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dbh"))
            d_151_v27_: _dafny.MultiSet
            d_151_v27_ = _dafny.MultiSet([d_150_v26_, (d_150_v26_).set(default__.safeIndex(len(d_147_v23_), len(d_150_v26_)), d_146_v22_), d_150_v26_])
            d_151_v27_ = d_151_v27_
            d_152_v28_: C1
            nw28_ = C1()
            nw28_.ctor__()
            d_152_v28_ = nw28_
            d_152_v28_ = d_152_v28_
            d_153_v29_: bool
            d_153_v29_ = not(d_126_v5_)
            d_154_v30_: C0
            nw29_ = C0()
            nw29_.ctor__(default__.fm5(d_128_globalState_), (d_145_i1_) + (d_121_v0_), d_153_v29_)
            d_154_v30_ = nw29_
        d_155_v31_: _dafny.Seq
        d_155_v31_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nt"))
        hi2_ = d_121_v0_
        for d_156_i2_ in range(len(d_155_v31_), hi2_):
            d_157_v32_: D2
            d_157_v32_ = D2_DC4(d_126_v5_, d_121_v0_)
            d_158_v33_: _dafny.Array
            nw30_ = _dafny.Array(None, 8)
            nw30_[int(0)] = d_126_v5_
            nw30_[int(1)] = d_126_v5_
            nw30_[int(2)] = d_126_v5_
            nw30_[int(3)] = d_126_v5_
            nw30_[int(4)] = d_126_v5_
            nw30_[int(5)] = d_126_v5_
            nw30_[int(6)] = (d_157_v32_).cf6
            nw30_[int(7)] = d_126_v5_
            d_158_v33_ = nw30_
            index19_ = default__.safeIndex(148, (d_158_v33_).length(0))
            (d_158_v33_)[index19_] = (d_156_i2_) > (d_156_i2_)
            d_159_v34_: str
            d_159_v34_ = _dafny.CodePoint('u')
            d_160_v35_: _dafny.Map
            d_160_v35_ = _dafny.Map({d_156_i2_: d_156_i2_})
            d_161_v36_: _dafny.Set
            d_161_v36_ = _dafny.Set({d_155_v31_})
            d_162_v37_: D2
            d_162_v37_ = D2_DC5((default__.fm6(default__.fm3(d_159_v34_, d_126_v5_, d_126_v5_, d_128_globalState_), d_126_v5_, d_156_i2_, d_128_globalState_) if d_126_v5_ else d_160_v35_), (default__.fm7(d_126_v5_, d_126_v5_, d_128_globalState_)).isdisjoint(d_161_v36_), d_126_v5_, d_126_v5_, d_126_v5_)
            d_163_v38_: _dafny.Seq
            d_163_v38_ = _dafny.SeqWithoutIsStrInference([d_126_v5_])
            d_164_v39_: _dafny.Map
            d_164_v39_ = _dafny.Map({(d_163_v38_)[default__.safeIndex(len(d_163_v38_), len(d_163_v38_))]: d_126_v5_})
            index20_ = default__.safeIndex(148, (d_158_v33_).length(0))
            rhs26_ = not(d_126_v5_)
            rhs27_ = (((d_164_v39_)[False] if (False) in (d_164_v39_) else not(d_126_v5_))) == ((d_156_i2_) >= (default__.fm3(d_159_v34_, d_126_v5_, d_126_v5_, d_128_globalState_)))
            rhs28_ = D2_DC5(d_160_v35_, d_126_v5_, d_126_v5_, d_126_v5_, d_126_v5_)
            lhs10_ = d_158_v33_
            lhs11_ = default__.safeIndex(148, (d_158_v33_).length(0))
            d_126_v5_ = rhs26_
            lhs10_[lhs11_] = rhs27_
            d_162_v37_ = rhs28_
            d_165_v40_: _dafny.Map
            d_165_v40_ = _dafny.Map({True: d_124_v3_})
            pat_let_tv0_ = d_159_v34_
            def iife7_(_pat_let0_0):
                def iife8_(d_166_dt__update__tmp_h0_):
                    def iife9_(_pat_let1_0):
                        def iife10_(d_167_dt__update_hcf2_h0_):
                            return D1_DC2(d_167_dt__update_hcf2_h0_, (d_166_dt__update__tmp_h0_).cf3, (d_166_dt__update__tmp_h0_).cf4)
                        return iife10_(_pat_let1_0)
                    return iife9_(pat_let_tv0_)
                return iife8_(_pat_let0_0)
            source0_ = iife7_((D1_DC2(d_159_v34_, d_121_v0_, len(d_165_v40_)) if (d_158_v33_)[default__.safeIndex(148, (d_158_v33_).length(0))] else D1_DC2(d_159_v34_, len(d_155_v31_), (0) - (d_121_v0_))))
            if source0_.is_DC2:
                d_168___mcc_h0_ = source0_.cf2
                d_169___mcc_h1_ = source0_.cf3
                d_170___mcc_h2_ = source0_.cf4
                d_171_cf4_ = d_170___mcc_h2_
                d_172_cf3_ = d_169___mcc_h1_
                d_173_cf2_ = d_168___mcc_h0_
                d_174_v41_: _dafny.Seq
                d_174_v41_ = _dafny.SeqWithoutIsStrInference([d_121_v0_, d_156_i2_])
                d_175_v42_: _dafny.Map
                d_175_v42_ = _dafny.Map({(d_163_v38_)[default__.safeIndex(d_121_v0_, len(d_163_v38_))]: _dafny.MultiSet(d_174_v41_)})
                d_176_v43_: bool
                d_176_v43_ = d_126_v5_
                d_177_v44_: C0
                nw31_ = C0()
                nw31_.ctor__(d_175_v42_, len(d_155_v31_), d_176_v43_)
                d_177_v44_ = nw31_
                d_177_v44_ = d_177_v44_
                d_124_v3_ = d_124_v3_
                index21_ = default__.safeIndex(148, (d_158_v33_).length(0))
                (d_158_v33_)[index21_] = (d_155_v31_) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nwf")))
                d_178_v45_: _dafny.Map
                d_178_v45_ = _dafny.Map({d_159_v34_: d_126_v5_})
                d_172_cf3_ = len((d_178_v45_) | (d_178_v45_))
            elif True:
                d_179___mcc_h3_ = source0_.cf1
                d_180_cf1_ = d_179___mcc_h3_
                d_126_v5_ = (d_158_v33_)[default__.safeIndex(148, (d_158_v33_).length(0))]
                d_181_v46_: C1
                nw32_ = C1()
                nw32_.ctor__()
                d_181_v46_ = nw32_
                d_182_v47_: _dafny.Array
                nw33_ = _dafny.Array(None, 15)
                nw33_[int(0)] = d_181_v46_
                nw33_[int(1)] = d_181_v46_
                nw33_[int(2)] = d_181_v46_
                nw33_[int(3)] = d_181_v46_
                nw33_[int(4)] = d_181_v46_
                nw33_[int(5)] = (d_181_v46_ if (d_158_v33_)[default__.safeIndex(148, (d_158_v33_).length(0))] else d_181_v46_)
                nw33_[int(6)] = d_181_v46_
                nw33_[int(7)] = d_181_v46_
                nw33_[int(8)] = d_181_v46_
                nw33_[int(9)] = d_181_v46_
                nw33_[int(10)] = d_181_v46_
                nw33_[int(11)] = d_181_v46_
                nw33_[int(12)] = d_181_v46_
                nw33_[int(13)] = (d_181_v46_ if d_126_v5_ else d_181_v46_)
                nw33_[int(14)] = d_181_v46_
                d_182_v47_ = nw33_
                index22_ = default__.safeIndex(9, (d_182_v47_).length(0))
                (d_182_v47_)[index22_] = d_181_v46_
                index23_ = default__.safeIndex(9, (d_182_v47_).length(0))
                (d_182_v47_)[index23_] = d_181_v46_
                d_183_v48_: D4
                d_183_v48_ = D4_DC9(d_158_v33_)
                d_158_v33_ = (d_183_v48_).cf17
                d_124_v3_ = d_124_v3_
            if not(d_126_v5_):
                d_184_v49_: _dafny.Array
                nw34_ = _dafny.Array(_dafny.Seq({}), 16)
                d_184_v49_ = nw34_
                d_185_v50_: _dafny.Seq
                d_185_v50_ = _dafny.SeqWithoutIsStrInference([d_156_i2_, d_121_v0_, d_156_i2_])
                index24_ = default__.safeIndex(55, (d_184_v49_).length(0))
                (d_184_v49_)[index24_] = _dafny.SeqWithoutIsStrInference([(d_185_v50_)[default__.safeIndex(d_121_v0_, len(d_185_v50_))], len(d_160_v35_), d_156_i2_, (d_122_v1_).cardinality])
                index25_ = default__.safeIndex(55, (d_184_v49_).length(0))
                (d_184_v49_)[index25_] = d_185_v50_
                d_126_v5_ = not(not (default__.fm8(d_121_v0_, d_160_v35_, ((d_122_v1_)[254] if (254) in (d_122_v1_) else len(d_163_v38_)), d_126_v5_, d_128_globalState_)) or (False))
                index26_ = default__.safeIndex(148, (d_158_v33_).length(0))
                (d_158_v33_)[index26_] = ((_dafny.SeqWithoutIsStrInference([(d_158_v33_)[default__.safeIndex(148, (d_158_v33_).length(0))]])) <= (_dafny.SeqWithoutIsStrInference([d_126_v5_])) if (d_158_v33_)[default__.safeIndex(148, (d_158_v33_).length(0))] else (d_158_v33_)[default__.safeIndex(148, (d_158_v33_).length(0))])
                d_186_v51_: _dafny.Map
                d_186_v51_ = _dafny.Map({(d_158_v33_)[default__.safeIndex(148, (d_158_v33_).length(0))]: (d_122_v1_) - (_dafny.MultiSet([d_121_v0_]))})
                d_186_v51_ = (d_186_v51_).set((d_158_v33_)[default__.safeIndex(148, (d_158_v33_).length(0))], default__.fm2(d_128_globalState_))
                d_187_v52_: _dafny.Map
                d_187_v52_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lpuwc")): default__.fm3(d_159_v34_, (d_158_v33_)[default__.safeIndex(148, (d_158_v33_).length(0))], d_126_v5_, d_128_globalState_)})
                d_188_v53_: bool
                d_188_v53_ = (d_158_v33_)[default__.safeIndex(148, (d_158_v33_).length(0))]
                d_189_v54_: C0
                nw35_ = C0()
                nw35_.ctor__(d_186_v51_, (d_122_v1_).cardinality, d_188_v53_)
                d_189_v54_ = nw35_
                d_190_v55_: _dafny.Map
                d_190_v55_ = _dafny.Map({d_189_v54_: d_126_v5_})
                d_187_v52_ = (d_187_v52_).set(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bo")), len((d_190_v55_).set(d_189_v54_, (d_158_v33_)[default__.safeIndex(148, (d_158_v33_).length(0))])))
            elif True:
                d_121_v0_ = d_121_v0_
                d_191_v56_: _dafny.Seq
                d_191_v56_ = _dafny.SeqWithoutIsStrInference([d_124_v3_])
                d_192_v57_: D5
                d_192_v57_ = D5_DC11(d_191_v56_)
                index27_ = default__.safeIndex(148, (d_158_v33_).length(0))
                (d_158_v33_)[index27_] = default__.fm8(d_156_i2_, d_160_v35_, d_121_v0_, default__.fm8(len((d_192_v57_).cf18), (d_160_v35_).set(d_156_i2_, d_121_v0_), d_156_i2_, d_126_v5_, d_128_globalState_), d_128_globalState_)
                d_191_v56_ = _dafny.SeqWithoutIsStrInference([d_124_v3_, d_124_v3_, d_124_v3_, d_124_v3_, d_124_v3_])
                d_121_v0_ = d_156_i2_
                d_193_v58_: C1
                nw36_ = C1()
                nw36_.ctor__()
                d_193_v58_ = nw36_
            d_126_v5_ = default__.fm8(len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([d_159_v34_ for d_194_i4_ in range(default__.abs(682))])) for d_195_i3_ in range(default__.abs(700))])), d_160_v35_, d_156_i2_, d_126_v5_, d_128_globalState_)
        if d_126_v5_:
            d_196_v59_: _dafny.Array
            def lambda8_(d_197_v5_):
                def lambda9_(d_198_i5_):
                    return d_197_v5_

                return lambda9_

            init4_ = lambda8_(d_126_v5_)
            nw37_ = _dafny.Array(None, 19)
            for i0_4_ in range(nw37_.length(0)):
                nw37_[i0_4_] = init4_(i0_4_)
            d_196_v59_ = nw37_
            d_199_v60_: str
            d_199_v60_ = _dafny.CodePoint('m')
            d_200_v61_: _dafny.Map
            d_200_v61_ = _dafny.Map({d_199_v60_: d_121_v0_})
            index28_ = default__.safeIndex(594, (d_196_v59_).length(0))
            (d_196_v59_)[index28_] = not((d_122_v1_).isdisjoint(_dafny.MultiSet([d_121_v0_, d_121_v0_, d_121_v0_, len(d_200_v61_)])))
            index29_ = default__.safeIndex(594, (d_196_v59_).length(0))
            (d_196_v59_)[index29_] = d_126_v5_
            d_201_v62_: _dafny.Set
            d_201_v62_ = _dafny.Set({d_196_v59_})
            index30_ = default__.safeIndex(594, (d_196_v59_).length(0))
            (d_196_v59_)[index30_] = (_dafny.Set({d_196_v59_, d_196_v59_})).issubset((d_201_v62_).intersection(d_201_v62_))
            d_121_v0_ = 705
            d_202_v63_: _dafny.Map
            d_202_v63_ = _dafny.Map({True: d_122_v1_})
            d_203_v64_: bool
            d_203_v64_ = (d_196_v59_)[default__.safeIndex(594, (d_196_v59_).length(0))]
            d_204_v65_: C0
            nw38_ = C0()
            nw38_.ctor__(d_202_v63_, d_121_v0_, d_203_v64_)
            d_204_v65_ = nw38_
            d_205_v66_: _dafny.Map
            d_205_v66_ = _dafny.Map({(d_196_v59_)[default__.safeIndex(594, (d_196_v59_).length(0))]: d_121_v0_})
            d_206_v67_: _dafny.Set
            d_206_v67_ = _dafny.Set({len(d_127_v6_), 238, len(d_205_v66_)})
            d_207_v68_: _dafny.Seq
            d_207_v68_ = _dafny.SeqWithoutIsStrInference([d_206_v67_])
            d_208_v69_: D6
            d_208_v69_ = D6_DC15(d_204_v65_)
            d_209_v70_: D7
            d_209_v70_ = D7_DC19(d_207_v68_)
            rhs29_ = (d_208_v69_).cf21
            rhs30_ = d_121_v0_
            rhs31_ = (d_209_v70_).cf25
            d_204_v65_ = rhs29_
            d_121_v0_ = rhs30_
            d_207_v68_ = rhs31_
            d_210_v71_: C1
            nw39_ = C1()
            nw39_.ctor__()
            d_210_v71_ = nw39_
            d_210_v71_ = d_210_v71_
        elif True:
            d_211_v72_: _dafny.Map
            d_211_v72_ = _dafny.Map({(0) - (d_121_v0_): d_121_v0_})
            d_212_v73_: _dafny.MultiSet
            d_212_v73_ = _dafny.MultiSet([d_126_v5_])
            d_211_v72_ = (d_211_v72_).set(((d_212_v73_).cardinality) + (d_121_v0_), (0) - (d_121_v0_))
            d_121_v0_ = (d_121_v0_) - (default__.safeModuloInt(d_121_v0_, d_121_v0_))
            d_213_v74_: str
            d_213_v74_ = _dafny.CodePoint('n')
            d_155_v31_ = (d_155_v31_).set(default__.safeIndex(791, len(d_155_v31_)), d_213_v74_)
            d_214_v75_: _dafny.Array
            def lambda10_(d_215_v0_):
                def lambda11_(d_216_i6_):
                    return _dafny.SeqWithoutIsStrInference([d_215_v0_, 535])

                return lambda11_

            init5_ = lambda10_(d_121_v0_)
            nw40_ = _dafny.Array(None, 16)
            for i0_5_ in range(nw40_.length(0)):
                nw40_[i0_5_] = init5_(i0_5_)
            d_214_v75_ = nw40_
            d_217_v76_: _dafny.Seq
            d_217_v76_ = _dafny.SeqWithoutIsStrInference([117])
            index31_ = default__.safeIndex(606, (d_214_v75_).length(0))
            (d_214_v75_)[index31_] = (d_217_v76_).set(default__.safeIndex(973, len(d_217_v76_)), 462)
            index32_ = default__.safeIndex(606, (d_214_v75_).length(0))
            (d_214_v75_)[index32_] = d_217_v76_
            d_218_v77_: _dafny.Seq
            d_219_v78_: C0
            d_220_v79_: _dafny.Seq
            d_221_v80_: bool
            out2_: _dafny.Seq
            out3_: C0
            out4_: _dafny.Seq
            out5_: bool
            out2_, out3_, out4_, out5_ = default__.m1(d_213_v74_, d_126_v5_, d_128_globalState_)
            d_218_v77_ = out2_
            d_219_v78_ = out3_
            d_220_v79_ = out4_
            d_221_v80_ = out5_
        d_222_v81_: _dafny.Map
        d_222_v81_ = _dafny.Map({d_121_v0_: d_126_v5_})
        d_222_v81_ = (d_222_v81_).set(default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([d_121_v0_])), 623), d_126_v5_)
        if True:
            d_223_v82_: str
            d_223_v82_ = _dafny.CodePoint('c')
            d_224_v83_: _dafny.Seq
            d_225_v84_: C0
            d_226_v85_: _dafny.Seq
            d_227_v86_: bool
            out6_: _dafny.Seq
            out7_: C0
            out8_: _dafny.Seq
            out9_: bool
            out6_, out7_, out8_, out9_ = default__.m1((d_223_v82_ if True else d_223_v82_), d_126_v5_, d_128_globalState_)
            d_224_v83_ = out6_
            d_225_v84_ = out7_
            d_226_v85_ = out8_
            d_227_v86_ = out9_
            d_126_v5_ = d_126_v5_
            d_121_v0_ = d_121_v0_
            d_228_v87_: _dafny.Seq
            d_229_v88_: C0
            d_230_v89_: _dafny.Seq
            d_231_v90_: bool
            out10_: _dafny.Seq
            out11_: C0
            out12_: _dafny.Seq
            out13_: bool
            out10_, out11_, out12_, out13_ = default__.m1(d_223_v82_, d_126_v5_, d_128_globalState_)
            d_228_v87_ = out10_
            d_229_v88_ = out11_
            d_230_v89_ = out12_
            d_231_v90_ = out13_
            d_231_v90_ = True
        elif True:
            d_232_v91_: _dafny.Seq
            d_232_v91_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_233_i7_ in range(default__.abs(162))]))])
            d_234_v92_: _dafny.Map
            d_234_v92_ = _dafny.Map({d_126_v5_: _dafny.MultiSet(d_232_v91_)})
            d_235_v93_: str
            d_235_v93_ = _dafny.CodePoint('x')
            d_236_v94_: C0
            nw41_ = C0()
            nw41_.ctor__(d_234_v92_, default__.fm3(d_235_v93_, d_126_v5_, d_126_v5_, d_128_globalState_), False)
            d_236_v94_ = nw41_
            d_237_v95_: _dafny.Seq
            d_237_v95_ = _dafny.SeqWithoutIsStrInference([d_236_v94_])
            d_238_v96_: D6
            d_238_v96_ = D6_DC15(d_236_v94_)
            d_239_v97_: _dafny.Seq
            d_239_v97_ = _dafny.SeqWithoutIsStrInference([((d_237_v95_) + (d_237_v95_)).set(default__.safeIndex(len(d_155_v31_), len((d_237_v95_) + (d_237_v95_))), (d_238_v96_).cf21), d_237_v95_, d_237_v95_])
            d_239_v97_ = _dafny.SeqWithoutIsStrInference([d_237_v95_])
            index33_ = default__.safeIndex(190, (d_124_v3_).length(0))
            (d_124_v3_)[index33_] = (d_121_v0_) - (d_121_v0_)
            index34_ = default__.safeIndex(190, (d_124_v3_).length(0))
            (d_124_v3_)[index34_] = (d_121_v0_) * (d_121_v0_)
            d_240_v98_: _dafny.Seq
            d_240_v98_ = _dafny.SeqWithoutIsStrInference([False, d_126_v5_, d_126_v5_])
            d_241_v99_: _dafny.Seq
            d_241_v99_ = _dafny.SeqWithoutIsStrInference([d_240_v98_, d_240_v98_, (d_240_v98_).set(default__.safeIndex((d_124_v3_)[default__.safeIndex(190, (d_124_v3_).length(0))], len(d_240_v98_)), d_126_v5_), d_240_v98_, d_240_v98_])
            d_242_v100_: bool
            d_242_v100_ = False
            d_243_v101_: C0
            nw42_ = C0()
            nw42_.ctor__(_dafny.Map({d_126_v5_: _dafny.MultiSet([(d_124_v3_)[default__.safeIndex(190, (d_124_v3_).length(0))], len(_dafny.SeqWithoutIsStrInference([(d_124_v3_)[default__.safeIndex(190, (d_124_v3_).length(0))], d_121_v0_]))])}), len(d_241_v99_), not(d_126_v5_))
            d_243_v101_ = nw42_
            d_244_v103_: _dafny.Set
            def iife11_():
                coll7_ = _dafny.Map()
                compr_7_: int
                for compr_7_ in _dafny.IntegerRange(930, 440):
                    d_245_v102_: int = compr_7_
                    if ((930) <= (d_245_v102_)) and ((d_245_v102_) < (440)):
                        coll7_[(d_245_v102_) * ((d_124_v3_)[default__.safeIndex(190, (d_124_v3_).length(0))])] = True
                return _dafny.Map(coll7_)
            d_244_v103_ = _dafny.Set({d_222_v81_, iife11_()
            })
            d_246_v104_: _dafny.Map
            d_246_v104_ = _dafny.Map({len(d_240_v98_): d_244_v103_})
            d_246_v104_ = (d_246_v104_).set(-48, d_244_v103_)
            d_247_v105_: _dafny.Array
            def lambda12_(d_248_v5_):
                def lambda13_(d_249_i8_):
                    return (_dafny.Map({not(d_248_v5_): d_248_v5_})) | (_dafny.Map({True: not(d_248_v5_)}))

                return lambda13_

            init6_ = lambda12_(d_126_v5_)
            nw43_ = _dafny.Array(None, 27)
            for i0_6_ in range(nw43_.length(0)):
                nw43_[i0_6_] = init6_(i0_6_)
            d_247_v105_ = nw43_
            d_250_v106_: _dafny.Map
            d_250_v106_ = _dafny.Map({not(d_126_v5_): d_126_v5_})
            index35_ = default__.safeIndex(503, (d_247_v105_).length(0))
            (d_247_v105_)[index35_] = (d_250_v106_) | (_dafny.Map({d_126_v5_: d_126_v5_}))
            index36_ = default__.safeIndex(503, (d_247_v105_).length(0))
            (d_247_v105_)[index36_] = _dafny.Map({False: d_126_v5_})
        hi3_ = d_121_v0_
        for d_251_i9_ in range(d_121_v0_, hi3_):
            d_252_v107_: _dafny.Map
            d_252_v107_ = _dafny.Map({not(True): d_122_v1_})
            d_253_v108_: bool
            d_253_v108_ = d_126_v5_
            d_254_v109_: C0
            nw44_ = C0()
            nw44_.ctor__(d_252_v107_, (d_121_v0_) * ((0) - (len(d_155_v31_))), d_253_v108_)
            d_254_v109_ = nw44_
            d_254_v109_ = d_254_v109_
            d_255_v110_: _dafny.Array
            nw45_ = _dafny.Array(False, 13)
            d_255_v110_ = nw45_
            index37_ = default__.safeIndex(24, (d_255_v110_).length(0))
            (d_255_v110_)[index37_] = d_126_v5_
            d_256_v111_: _dafny.MultiSet
            d_256_v111_ = _dafny.MultiSet([not (False) or (d_126_v5_), d_126_v5_])
            index38_ = default__.safeIndex(24, (d_255_v110_).length(0))
            rhs32_ = d_121_v0_
            rhs33_ = (d_122_v1_).intersection(d_122_v1_)
            rhs34_ = d_126_v5_
            rhs35_ = ((d_256_v111_)[False] if (False) in (d_256_v111_) else d_121_v0_)
            lhs12_ = d_255_v110_
            lhs13_ = default__.safeIndex(24, (d_255_v110_).length(0))
            d_121_v0_ = rhs32_
            d_122_v1_ = rhs33_
            lhs12_[lhs13_] = rhs34_
            d_121_v0_ = rhs35_
            index39_ = default__.safeIndex(24, (d_255_v110_).length(0))
            (d_255_v110_)[index39_] = (((d_256_v111_).intersection(d_256_v111_)).cardinality) != (d_251_i9_)
            d_257_v112_: C1
            nw46_ = C1()
            nw46_.ctor__()
            d_257_v112_ = nw46_
        d_258_v113_: _dafny.Array
        nw47_ = _dafny.Array(False, 28)
        d_258_v113_ = nw47_
        index40_ = default__.safeIndex(201, (d_258_v113_).length(0))
        (d_258_v113_)[index40_] = True
        index41_ = default__.safeIndex(201, (d_258_v113_).length(0))
        (d_258_v113_)[index41_] = d_126_v5_
        d_259_v114_: _dafny.Map
        d_259_v114_ = _dafny.Map({(d_258_v113_)[default__.safeIndex(201, (d_258_v113_).length(0))]: (d_258_v113_)[default__.safeIndex(201, (d_258_v113_).length(0))]})
        d_260_v115_: _dafny.Seq
        d_260_v115_ = _dafny.SeqWithoutIsStrInference([d_121_v0_])
        d_259_v114_ = (d_259_v114_).set((d_258_v113_)[default__.safeIndex(201, (d_258_v113_).length(0))], (d_121_v0_) in (d_260_v115_))
        d_261_v116_: str
        d_261_v116_ = _dafny.CodePoint('i')
        d_261_v116_ = d_261_v116_
        d_261_v116_ = _dafny.CodePoint('h')
        d_262_v117_: _dafny.Map
        d_262_v117_ = _dafny.Map({d_121_v0_: d_121_v0_})
        d_263_v118_: _dafny.Map
        d_263_v118_ = _dafny.Map({len(d_262_v117_): d_121_v0_})
        d_264_v119_: _dafny.Seq
        d_264_v119_ = _dafny.SeqWithoutIsStrInference([d_126_v5_])
        hi4_ = (((d_263_v118_)[len(_dafny.SeqWithoutIsStrInference([d_261_v116_ for d_265_i11_ in range(default__.abs(329))]))] if (len(_dafny.SeqWithoutIsStrInference([d_261_v116_ for d_266_i11_ in range(default__.abs(329))]))) in (d_263_v118_) else d_121_v0_)) - (len(d_264_v119_))
        for d_267_i10_ in range((d_121_v0_) * (d_121_v0_), hi4_):
            d_126_v5_ = not(d_126_v5_)
            if d_126_v5_:
                d_268_v120_: _dafny.Set
                d_268_v120_ = _dafny.Set({d_261_v116_})
                index42_ = default__.safeIndex(201, (d_258_v113_).length(0))
                (d_258_v113_)[index42_] = ((-528) - (d_267_i10_)) >= (len(d_268_v120_))
                d_269_v121_: _dafny.Seq
                d_270_v122_: C0
                d_271_v123_: _dafny.Seq
                d_272_v124_: bool
                out14_: _dafny.Seq
                out15_: C0
                out16_: _dafny.Seq
                out17_: bool
                out14_, out15_, out16_, out17_ = default__.m1((d_155_v31_)[default__.safeIndex((0) - (d_121_v0_), len(d_155_v31_))], (d_258_v113_)[default__.safeIndex(201, (d_258_v113_).length(0))], d_128_globalState_)
                d_269_v121_ = out14_
                d_270_v122_ = out15_
                d_271_v123_ = out16_
                d_272_v124_ = out17_
                d_273_v125_: bool
                d_273_v125_ = d_126_v5_
                d_274_v126_: T0
                nw48_ = C0()
                nw48_.ctor__(default__.fm5(d_128_globalState_), default__.safeDivisionInt(d_267_i10_, -74), d_273_v125_)
                d_274_v126_ = nw48_
                d_274_v126_ = d_274_v126_
                d_121_v0_ = 201
                d_121_v0_ = (default__.fm9(d_267_i10_, d_272_v124_, (d_258_v113_)[default__.safeIndex(201, (d_258_v113_).length(0))], len(d_269_v121_), d_128_globalState_)).cf4
            elif True:
                d_275_v127_: _dafny.Map
                d_275_v127_ = _dafny.Map({(d_258_v113_)[default__.safeIndex(201, (d_258_v113_).length(0))]: default__.fm3(d_261_v116_, d_126_v5_, (d_258_v113_)[default__.safeIndex(201, (d_258_v113_).length(0))], d_128_globalState_)})
                d_275_v127_ = (d_275_v127_).set(not((d_121_v0_) != (d_267_i10_)), len(d_264_v119_))
                d_276_v128_: _dafny.Map
                d_276_v128_ = _dafny.Map({d_127_v6_: d_126_v5_})
                d_276_v128_ = (d_276_v128_).set(d_127_v6_, d_126_v5_)
                d_260_v115_ = (d_260_v115_) + (d_260_v115_)
                d_126_v5_ = ((d_258_v113_)[default__.safeIndex(201, (d_258_v113_).length(0))]) or ((d_258_v113_)[default__.safeIndex(201, (d_258_v113_).length(0))])
                d_277_v129_: _dafny.Seq
                d_278_v130_: C0
                d_279_v131_: _dafny.Seq
                d_280_v132_: bool
                out18_: _dafny.Seq
                out19_: C0
                out20_: _dafny.Seq
                out21_: bool
                out18_, out19_, out20_, out21_ = default__.m1((d_261_v116_ if d_126_v5_ else d_261_v116_), (d_258_v113_)[default__.safeIndex(201, (d_258_v113_).length(0))], d_128_globalState_)
                d_277_v129_ = out18_
                d_278_v130_ = out19_
                d_279_v131_ = out20_
                d_280_v132_ = out21_
            d_281_v133_: _dafny.Map
            d_281_v133_ = _dafny.Map({(d_258_v113_)[default__.safeIndex(201, (d_258_v113_).length(0))]: d_121_v0_})
            d_282_v134_: _dafny.Map
            d_282_v134_ = _dafny.Map({d_126_v5_: d_122_v1_})
            d_283_v135_: bool
            d_283_v135_ = (d_258_v113_)[default__.safeIndex(201, (d_258_v113_).length(0))]
            d_284_v136_: T0
            nw49_ = C0()
            nw49_.ctor__(d_282_v134_, len(d_155_v31_), d_283_v135_)
            d_284_v136_ = nw49_
            d_285_v137_: _dafny.Map
            d_285_v137_ = _dafny.Map({d_284_v136_: d_155_v31_})
            d_263_v118_ = (d_263_v118_).set(d_121_v0_, ((d_281_v133_)[(d_258_v113_)[default__.safeIndex(201, (d_258_v113_).length(0))]] if ((d_258_v113_)[default__.safeIndex(201, (d_258_v113_).length(0))]) in (d_281_v133_) else len(d_285_v137_)))
            d_121_v0_ = d_267_i10_
        d_126_v5_ = d_126_v5_
        d_286_v138_: _dafny.Seq
        d_286_v138_ = _dafny.SeqWithoutIsStrInference([d_124_v3_, d_124_v3_, d_124_v3_])
        d_287_v139_: D5
        d_287_v139_ = D5_DC11(d_286_v138_)
        rhs36_ = (d_258_v113_ if d_126_v5_ else d_258_v113_)
        rhs37_ = d_287_v139_
        d_258_v113_ = rhs36_
        d_287_v139_ = rhs37_
        d_126_v5_ = d_126_v5_
        d_288_v140_: _dafny.Array
        nw50_ = _dafny.Array(_dafny.Seq({}), 11)
        d_288_v140_ = nw50_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_288_v140_).length(0)):
            d_289_i12_: int = guard_loop_0_
            if (True) and (((0) <= (d_289_i12_)) and ((d_289_i12_) < ((d_288_v140_).length(0)))):
                (d_288_v140_)[(d_289_i12_)] = d_260_v115_
        _dafny.print(_dafny.string_of(d_121_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_122_v1_) == (_dafny.MultiSet([197]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_123_v2_)[0]) == (_dafny.MultiSet([197]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_123_v2_)[1]) == (_dafny.MultiSet([197]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_123_v2_)[2]) == (_dafny.MultiSet([197]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_123_v2_)[3]) == (_dafny.MultiSet([197]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_123_v2_)[4]) == (_dafny.MultiSet([197]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_123_v2_)[5]) == (_dafny.MultiSet([197, 197, 265]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_123_v2_)[6]) == (_dafny.MultiSet([197]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_123_v2_)[7]) == (_dafny.MultiSet([197]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_123_v2_)[8]) == (_dafny.MultiSet([197]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_123_v2_)[9]) == (_dafny.MultiSet([197]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_123_v2_)[10]) == (_dafny.MultiSet([197]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_123_v2_)[11]) == (_dafny.MultiSet([197]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_123_v2_)[12]) == (_dafny.MultiSet([197]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_125_v4_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_126_v5_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_127_v6_) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_128_globalState_).f0)[0]) == (_dafny.MultiSet([197]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_128_globalState_).f0)[1]) == (_dafny.MultiSet([197]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_128_globalState_).f0)[2]) == (_dafny.MultiSet([197]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_128_globalState_).f0)[3]) == (_dafny.MultiSet([197]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_128_globalState_).f0)[4]) == (_dafny.MultiSet([197]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_128_globalState_).f0)[5]) == (_dafny.MultiSet([197, 197, 265]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_128_globalState_).f0)[6]) == (_dafny.MultiSet([197]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_128_globalState_).f0)[7]) == (_dafny.MultiSet([197]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_128_globalState_).f0)[8]) == (_dafny.MultiSet([197]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_128_globalState_).f0)[9]) == (_dafny.MultiSet([197]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_128_globalState_).f0)[10]) == (_dafny.MultiSet([197]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_128_globalState_).f0)[11]) == (_dafny.MultiSet([197]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_128_globalState_).f0)[12]) == (_dafny.MultiSet([197]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_128_globalState_).f1)[0]) == (_dafny.MultiSet([197]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_128_globalState_).f1)[1]) == (_dafny.MultiSet([197]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_128_globalState_).f1)[2]) == (_dafny.MultiSet([197]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_128_globalState_).f1)[3]) == (_dafny.MultiSet([197]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_128_globalState_).f1)[4]) == (_dafny.MultiSet([197]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_128_globalState_).f1)[5]) == (_dafny.MultiSet([197, 197, 265]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_128_globalState_).f1)[6]) == (_dafny.MultiSet([197]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_128_globalState_).f1)[7]) == (_dafny.MultiSet([197]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_128_globalState_).f1)[8]) == (_dafny.MultiSet([197]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_128_globalState_).f1)[9]) == (_dafny.MultiSet([197]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_128_globalState_).f1)[10]) == (_dafny.MultiSet([197]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_128_globalState_).f1)[11]) == (_dafny.MultiSet([197]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_128_globalState_).f1)[12]) == (_dafny.MultiSet([197]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_128_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_128_globalState_).f3)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_128_globalState_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_128_globalState_.f5) == (_dafny.Set({True, False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_155_v31_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_222_v81_) == (_dafny.Map({196: False, 1: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_258_v113_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_259_v114_) == (_dafny.Map({False: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_260_v115_) == (_dafny.SeqWithoutIsStrInference([196]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_261_v116_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_262_v117_) == (_dafny.Map({196: 196}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_263_v118_) == (_dafny.Map({1: 196}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_264_v119_) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_286_v138_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_287_v139_).cf18)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_288_v140_)[0]) == (_dafny.SeqWithoutIsStrInference([196]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_288_v140_)[1]) == (_dafny.SeqWithoutIsStrInference([196]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_288_v140_)[2]) == (_dafny.SeqWithoutIsStrInference([196]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_288_v140_)[3]) == (_dafny.SeqWithoutIsStrInference([196]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_288_v140_)[4]) == (_dafny.SeqWithoutIsStrInference([196]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_288_v140_)[5]) == (_dafny.SeqWithoutIsStrInference([196]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_288_v140_)[6]) == (_dafny.SeqWithoutIsStrInference([196]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_288_v140_)[7]) == (_dafny.SeqWithoutIsStrInference([196]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_288_v140_)[8]) == (_dafny.SeqWithoutIsStrInference([196]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_288_v140_)[9]) == (_dafny.SeqWithoutIsStrInference([196]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_288_v140_)[10]) == (_dafny.SeqWithoutIsStrInference([196]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: False
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)

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
        return lambda: D1_DC2(_dafny.CodePoint('D'), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D1_DC2)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D1_DC1)

class D1_DC2(D1, NamedTuple('DC2', [('cf2', Any), ('cf3', Any), ('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC2({_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)}, {_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC2) and self.cf2 == __o.cf2 and self.cf3 == __o.cf3 and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC1(D1, NamedTuple('DC1', [('cf1', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC1({_dafny.string_of(self.cf1)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC1) and self.cf1 == __o.cf1
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC4(False, int(0))
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

class D2_DC4(D2, NamedTuple('DC4', [('cf6', Any), ('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC4({_dafny.string_of(self.cf6)}, {_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC4) and self.cf6 == __o.cf6 and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC5(D2, NamedTuple('DC5', [('cf8', Any), ('cf9', Any), ('cf10', Any), ('cf11', Any), ('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC5({_dafny.string_of(self.cf8)}, {_dafny.string_of(self.cf9)}, {_dafny.string_of(self.cf10)}, {_dafny.string_of(self.cf11)}, {_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC5) and self.cf8 == __o.cf8 and self.cf9 == __o.cf9 and self.cf10 == __o.cf10 and self.cf11 == __o.cf11 and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC3(D2, NamedTuple('DC3', [('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC3({_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC3) and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC7(D1.default()(), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D3_DC7)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D3_DC6)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D3_DC8)

class D3_DC7(D3, NamedTuple('DC7', [('cf14', Any), ('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC7({_dafny.string_of(self.cf14)}, {_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC7) and self.cf14 == __o.cf14 and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC6(D3, NamedTuple('DC6', [('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC6({_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC6) and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC8(D3, NamedTuple('DC8', [('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC8({_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC8) and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC10()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D4_DC10)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D4_DC9)

class D4_DC10(D4, NamedTuple('DC10', [])):
    def __dafnystr__(self) -> str:
        return f'D4.DC10'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC10)
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC9(D4, NamedTuple('DC9', [('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC9({_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC9) and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC12(False)
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

class D5_DC12(D5, NamedTuple('DC12', [('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC12({_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC12) and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC13(D5, NamedTuple('DC13', [])):
    def __dafnystr__(self) -> str:
        return f'D5.DC13'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC13)
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC11(D5, NamedTuple('DC11', [('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC11({_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC11) and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC14(D5, NamedTuple('DC14', [('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC14({_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC14) and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC16(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D6_DC16)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D6_DC17)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D6_DC18)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D6_DC15)

class D6_DC16(D6, NamedTuple('DC16', [('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC16({_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC16) and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC17(D6, NamedTuple('DC17', [('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC17({_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC17) and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC18(D6, NamedTuple('DC18', [('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC18({_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC18) and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC15(D6, NamedTuple('DC15', [('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC15({_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC15) and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC20(_dafny.Array(None, 0))
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

class D7_DC20(D7, NamedTuple('DC20', [('cf26', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC20({_dafny.string_of(self.cf26)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC20) and self.cf26 == __o.cf26
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC21(D7, NamedTuple('DC21', [('cf27', Any), ('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC21({_dafny.string_of(self.cf27)}, {_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC21) and self.cf27 == __o.cf27 and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC19(D7, NamedTuple('DC19', [('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC19({_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC19) and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC23(False, _dafny.CodePoint('D'), False, _dafny.Seq({}))
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
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D8_DC25)

class D8_DC23(D8, NamedTuple('DC23', [('cf30', Any), ('cf31', Any), ('cf32', Any), ('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC23({_dafny.string_of(self.cf30)}, {_dafny.string_of(self.cf31)}, {_dafny.string_of(self.cf32)}, {_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC23) and self.cf30 == __o.cf30 and self.cf31 == __o.cf31 and self.cf32 == __o.cf32 and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC24(D8, NamedTuple('DC24', [('cf34', Any), ('cf35', Any), ('cf36', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC24({_dafny.string_of(self.cf34)}, {_dafny.string_of(self.cf35)}, {_dafny.string_of(self.cf36)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC24) and self.cf34 == __o.cf34 and self.cf35 == __o.cf35 and self.cf36 == __o.cf36
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC22(D8, NamedTuple('DC22', [('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC22({_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC22) and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC25(D8, NamedTuple('DC25', [('cf37', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC25({_dafny.string_of(self.cf37)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC25) and self.cf37 == __o.cf37
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC27()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D9_DC27)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D9_DC26)

class D9_DC27(D9, NamedTuple('DC27', [])):
    def __dafnystr__(self) -> str:
        return f'D9.DC27'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC27)
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC26(D9, NamedTuple('DC26', [('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC26({_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC26) and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: D10_DC29(int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D10_DC29)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D10_DC28)

class D10_DC29(D10, NamedTuple('DC29', [('cf40', Any), ('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC29({_dafny.string_of(self.cf40)}, {_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC29) and self.cf40 == __o.cf40 and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC28(D10, NamedTuple('DC28', [('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC28({_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC28) and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass

class GlobalState:
    def  __init__(self):
        self.f5: _dafny.Set = _dafny.Set({})
        self._f0: _dafny.Array = _dafny.Array(None, 0)
        self._f1: _dafny.Array = _dafny.Array(None, 0)
        self._f2: bool = False
        self._f3: _dafny.Map = _dafny.Map({})
        self._f4: str = _dafny.CodePoint('D')
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5):
        (self)._f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2
        (self)._f3 = f3
        (self)._f4 = f4
        (self).f5 = f5

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

class C0(T0):
    def  __init__(self):
        self._f6: int = int(0)
        self._f7: bool = False
        self.f8: _dafny.Map = _dafny.Map({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    @property
    def f6(self):
        return self._f6
    @property
    def f7(self):
        return self._f7
    def ctor__(self, f8, f6, f7):
        (self).f8 = f8
        (self)._f6 = f6
        (self)._f7 = f7

    def fm0(self, p0, p1, p2, p3, globalState):
        return default__.safeModuloInt(default__.safeModuloInt((0) - ((self).f6), (self).f6), (self).f6)

    def fm1(self, p0, p1, p2, p3, globalState):
        return not(((self).f6) > (((self).f6 if True else (self).f6)))


class C1:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    def ctor__(self):
        pass
        pass

    def m0(self, p0, p1, p2, globalState):
        r0: int = int(0)
        d_290_v0_: _dafny.Seq
        d_290_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mdcxtl"))
        d_291_v1_: int
        d_291_v1_ = 362
        d_292_v2_: _dafny.Set
        d_292_v2_ = _dafny.Set({_dafny.CodePoint('t'), p1, (d_290_v0_)[default__.safeIndex(d_291_v1_, len(d_290_v0_))]})
        d_292_v2_ = ((d_292_v2_) - (_dafny.Set({p1}))) | (d_292_v2_)
        d_293_i0_: int
        d_293_i0_ = 0
        with _dafny.label("0"):
            while p0:
                with _dafny.c_label("0"):
                    if (d_293_i0_) >= (100):
                        raise _dafny.Break("0")
                    d_293_i0_ = (d_293_i0_) + (1)
                    d_294_v3_: bool
                    d_294_v3_ = False
                    d_294_v3_ = (p2)
                    d_295_v4_: _dafny.Map
                    d_295_v4_ = _dafny.Map({d_294_v3_: default__.fm2(globalState)})
                    d_296_v5_: D1
                    d_296_v5_ = D1_DC1(d_295_v4_)
                    d_297_v6_: C0
                    nw51_ = C0()
                    nw51_.ctor__((d_296_v5_).cf1, d_291_v1_, p2)
                    d_297_v6_ = nw51_
                    d_298_v7_: _dafny.MultiSet
                    d_298_v7_ = _dafny.MultiSet([p0])
                    d_299_v8_: _dafny.MultiSet
                    d_299_v8_ = _dafny.MultiSet([d_297_v6_, d_297_v6_])
                    if (d_297_v6_).fm1(d_291_v1_, not((d_297_v6_).fm1(d_291_v1_, d_294_v3_, d_291_v1_, d_298_v7_, globalState)), default__.safeModuloInt((d_299_v8_).cardinality, (d_298_v7_).cardinality), (d_298_v7_).intersection(d_298_v7_), globalState):
                        d_300_v9_: _dafny.Seq
                        d_300_v9_ = _dafny.SeqWithoutIsStrInference([(d_298_v7_).ispropersubset(d_298_v7_)])
                        d_300_v9_ = d_300_v9_
                        r0 = d_291_v1_
                        d_301_v10_: _dafny.Set
                        d_301_v10_ = _dafny.Set({p0})
                        d_302_v11_: C0
                        nw52_ = C0()
                        nw52_.ctor__(d_297_v6_.f8, len(d_301_v10_), p2)
                        d_302_v11_ = nw52_
                        d_303_v12_: _dafny.Array
                        nw53_ = _dafny.Array(int(0), 7)
                        d_303_v12_ = nw53_
                        index43_ = default__.safeIndex(831, (d_303_v12_).length(0))
                        (d_303_v12_)[index43_] = default__.safeDivisionInt(d_291_v1_, d_291_v1_)
                        index44_ = default__.safeIndex(831, (d_303_v12_).length(0))
                        (d_303_v12_)[index44_] = d_291_v1_
                        d_304_v13_: _dafny.Array
                        nw54_ = _dafny.Array(None, 5)
                        nw54_[int(0)] = False
                        nw54_[int(1)] = d_294_v3_
                        nw54_[int(2)] = False
                        nw54_[int(3)] = d_294_v3_
                        nw54_[int(4)] = (_dafny.Set({not(d_294_v3_)})).isdisjoint(d_301_v10_)
                        d_304_v13_ = nw54_
                        index45_ = default__.safeIndex(784, (d_304_v13_).length(0))
                        (d_304_v13_)[index45_] = (341) <= (d_291_v1_)
                        index46_ = default__.safeIndex(784, (d_304_v13_).length(0))
                        (d_304_v13_)[index46_] = p0
                    elif True:
                        r0 = d_291_v1_
                        d_305_v14_: _dafny.Array
                        nw55_ = _dafny.Array(int(0), 14)
                        d_305_v14_ = nw55_
                        index47_ = default__.safeIndex(116, (d_305_v14_).length(0))
                        (d_305_v14_)[index47_] = default__.safeModuloInt(d_291_v1_, d_291_v1_)
                        index48_ = default__.safeIndex(116, (d_305_v14_).length(0))
                        (d_305_v14_)[index48_] = 47
                        d_291_v1_ = (d_305_v14_)[default__.safeIndex(116, (d_305_v14_).length(0))]
                        d_306_v15_: C0
                        nw56_ = C0()
                        nw56_.ctor__(d_297_v6_.f8, d_291_v1_, True)
                        d_306_v15_ = nw56_
                        d_307_v16_: _dafny.Seq
                        d_307_v16_ = _dafny.SeqWithoutIsStrInference([(d_305_v14_)[default__.safeIndex(116, (d_305_v14_).length(0))]])
                        d_308_v17_: _dafny.MultiSet
                        d_308_v17_ = _dafny.MultiSet([len(d_307_v16_)])
                        d_309_v18_: _dafny.Map
                        d_309_v18_ = _dafny.Map({d_308_v17_: (d_305_v14_)[default__.safeIndex(116, (d_305_v14_).length(0))]})
                        d_310_v19_: _dafny.Seq
                        d_310_v19_ = _dafny.SeqWithoutIsStrInference([d_294_v3_])
                        d_309_v18_ = (d_309_v18_).set(_dafny.MultiSet([d_291_v1_, len(d_310_v19_)]), d_291_v1_)
                    d_311_v20_: _dafny.MultiSet
                    d_311_v20_ = _dafny.MultiSet([d_291_v1_])
                    d_312_v21_: C0
                    nw57_ = C0()
                    nw57_.ctor__((_dafny.Map({p0: d_311_v20_})) | (d_295_v4_), default__.safeModuloInt(d_291_v1_, d_291_v1_), p2)
                    d_312_v21_ = nw57_
                    pass
            pass
        d_313_v22_: bool
        d_313_v22_ = False
        d_313_v22_ = not((default__.fm3(p1, d_313_v22_, d_313_v22_, globalState)) != (d_291_v1_))
        d_314_i1_: int
        d_314_i1_ = 0
        with _dafny.label("1"):
            while p0:
                with _dafny.c_label("1"):
                    if (d_314_i1_) >= (100):
                        raise _dafny.Break("1")
                    d_314_i1_ = (d_314_i1_) + (1)
                    d_315_v23_: _dafny.Array
                    def lambda14_(d_316_p0_):
                        def lambda15_(d_317_i2_):
                            return d_316_p0_

                        return lambda15_

                    init7_ = lambda14_(p0)
                    nw58_ = _dafny.Array(None, 18)
                    for i0_7_ in range(nw58_.length(0)):
                        nw58_[i0_7_] = init7_(i0_7_)
                    d_315_v23_ = nw58_
                    d_318_v24_: _dafny.Set
                    d_318_v24_ = _dafny.Set({d_315_v23_})
                    d_318_v24_ = d_318_v24_
                    d_319_v25_: _dafny.Array
                    nw59_ = _dafny.Array(int(0), 16)
                    d_319_v25_ = nw59_
                    d_320_v26_: _dafny.Map
                    d_320_v26_ = _dafny.Map({not(p0): d_319_v25_})
                    d_321_v27_: _dafny.Array
                    nw60_ = _dafny.Array(None, 21)
                    nw60_[int(0)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rymhitw")))
                    nw60_[int(1)] = d_291_v1_
                    nw60_[int(2)] = 687
                    nw60_[int(3)] = d_291_v1_
                    nw60_[int(4)] = 77
                    nw60_[int(5)] = len(d_290_v0_)
                    nw60_[int(6)] = (d_291_v1_) + (d_291_v1_)
                    nw60_[int(7)] = (d_291_v1_) + (d_291_v1_)
                    nw60_[int(8)] = d_291_v1_
                    nw60_[int(9)] = 899
                    nw60_[int(10)] = 397
                    nw60_[int(11)] = (d_291_v1_) + (d_291_v1_)
                    nw60_[int(12)] = d_291_v1_
                    nw60_[int(13)] = len((_dafny.SeqWithoutIsStrInference([p1 for d_322_i3_ in range(default__.abs(-807))])) + (d_290_v0_))
                    nw60_[int(14)] = len((d_320_v26_).set(d_313_v22_, d_319_v25_))
                    nw60_[int(15)] = d_291_v1_
                    nw60_[int(16)] = (0) - (default__.safeDivisionInt(len(d_290_v0_), d_291_v1_))
                    nw60_[int(17)] = 140
                    nw60_[int(18)] = (d_291_v1_) + (d_291_v1_)
                    nw60_[int(19)] = (d_291_v1_) - (d_291_v1_)
                    nw60_[int(20)] = (0) - (d_291_v1_)
                    d_321_v27_ = nw60_
                    index49_ = default__.safeIndex(17, (d_319_v25_).length(0))
                    (d_319_v25_)[index49_] = (d_291_v1_) * (56)
                    index50_ = default__.safeIndex(17, (d_319_v25_).length(0))
                    (d_319_v25_)[index50_] = (d_291_v1_) * ((-528) + (d_291_v1_))
                    if (d_291_v1_) < ((len(_dafny.Set({(d_319_v25_)[default__.safeIndex(17, (d_319_v25_).length(0))], d_291_v1_}))) * (len((D2_DC3(_dafny.Map({d_313_v22_: d_291_v1_}))).cf5))):
                        d_313_v22_ = p0
                        index51_ = default__.safeIndex(17, (d_319_v25_).length(0))
                        (d_319_v25_)[index51_] = (d_319_v25_)[default__.safeIndex(17, (d_319_v25_).length(0))]
                        d_323_v28_: D1
                        d_323_v28_ = D1_DC2(p1, -281, d_291_v1_)
                        d_324_v29_: _dafny.MultiSet
                        d_324_v29_ = _dafny.MultiSet([d_323_v28_, d_323_v28_, d_323_v28_, d_323_v28_])
                        d_313_v22_ = (d_324_v29_).issubset(_dafny.MultiSet([d_323_v28_]))
                        d_325_v30_: _dafny.MultiSet
                        d_325_v30_ = _dafny.MultiSet([(d_319_v25_)[default__.safeIndex(17, (d_319_v25_).length(0))]])
                        d_326_v31_: _dafny.Seq
                        d_326_v31_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_313_v22_: d_325_v30_})])
                        d_327_v32_: _dafny.Seq
                        d_327_v32_ = _dafny.SeqWithoutIsStrInference([d_313_v22_])
                        d_328_v33_: _dafny.Seq
                        d_328_v33_ = _dafny.SeqWithoutIsStrInference([d_327_v32_])
                        d_329_v34_: C0
                        nw61_ = C0()
                        nw61_.ctor__((d_326_v31_)[default__.safeIndex(len(d_328_v33_), len(d_326_v31_))], (d_319_v25_)[default__.safeIndex(17, (d_319_v25_).length(0))], p2)
                        d_329_v34_ = nw61_
                        d_330_v35_: C0
                        nw62_ = C0()
                        nw62_.ctor__(_dafny.Map({p0: d_325_v30_}), (d_319_v25_)[default__.safeIndex(17, (d_319_v25_).length(0))], p2)
                        d_330_v35_ = nw62_
                    elif True:
                        d_331_v36_: _dafny.Map
                        d_331_v36_ = _dafny.Map({(0) - ((0) - (d_291_v1_)): (0) - (-582)})
                        d_332_v37_: _dafny.Map
                        d_332_v37_ = _dafny.Map({p0: len(default__.fm4(d_331_v36_, p1, p0, d_313_v22_, globalState))})
                        d_333_v38_: D2
                        d_333_v38_ = D2_DC3(d_332_v37_)
                        d_333_v38_ = d_333_v38_
                        d_334_v39_: _dafny.Seq
                        d_334_v39_ = _dafny.SeqWithoutIsStrInference([d_313_v22_, p0, d_313_v22_, d_313_v22_, p0])
                        d_332_v37_ = (d_332_v37_).set(p0, (0) - ((0) - (default__.safeDivisionInt(len(d_334_v39_), (0) - (len(d_290_v0_))))))
                        d_313_v22_ = d_313_v22_
                        d_335_v40_: _dafny.Set
                        d_335_v40_ = _dafny.Set({p0, p0})
                        index52_ = default__.safeIndex(572, (d_315_v23_).length(0))
                        (d_315_v23_)[index52_] = (d_335_v40_).ispropersubset(d_335_v40_)
                        index53_ = default__.safeIndex(572, (d_315_v23_).length(0))
                        (d_315_v23_)[index53_] = (d_313_v22_) or (p0)
                        index54_ = default__.safeIndex(17, (d_319_v25_).length(0))
                        (d_319_v25_)[index54_] = (d_319_v25_)[default__.safeIndex(17, (d_319_v25_).length(0))]
                    r0 = (0) - ((0) - (d_291_v1_))
                    pass
            pass
        d_291_v1_ = (D2_DC4(p0, d_291_v1_)).cf7
        d_336_v41_: _dafny.Set
        d_336_v41_ = _dafny.Set({980})
        d_337_v42_: _dafny.Seq
        d_337_v42_ = _dafny.SeqWithoutIsStrInference([d_291_v1_, d_291_v1_])
        d_338_v43_: _dafny.Array
        nw63_ = _dafny.Array(None, 23)
        nw63_[int(0)] = d_291_v1_
        nw63_[int(1)] = d_291_v1_
        nw63_[int(2)] = d_291_v1_
        nw63_[int(3)] = d_291_v1_
        nw63_[int(4)] = d_291_v1_
        nw63_[int(5)] = 107
        nw63_[int(6)] = d_291_v1_
        nw63_[int(7)] = d_291_v1_
        nw63_[int(8)] = d_291_v1_
        nw63_[int(9)] = d_291_v1_
        nw63_[int(10)] = d_291_v1_
        nw63_[int(11)] = d_291_v1_
        nw63_[int(12)] = len(d_336_v41_)
        nw63_[int(13)] = (d_337_v42_)[default__.safeIndex(d_291_v1_, len(d_337_v42_))]
        nw63_[int(14)] = d_291_v1_
        nw63_[int(15)] = 696
        nw63_[int(16)] = d_291_v1_
        nw63_[int(17)] = d_291_v1_
        nw63_[int(18)] = d_291_v1_
        nw63_[int(19)] = d_291_v1_
        nw63_[int(20)] = d_291_v1_
        nw63_[int(21)] = d_291_v1_
        nw63_[int(22)] = d_291_v1_
        d_338_v43_ = nw63_
        d_339_v44_: _dafny.MultiSet
        d_339_v44_ = _dafny.MultiSet([d_338_v43_])
        rhs38_ = d_290_v0_
        rhs39_ = ((D3_DC6(d_339_v44_)).cf13).intersection(d_339_v44_)
        rhs40_ = not(d_313_v22_)
        d_290_v0_ = rhs38_
        d_339_v44_ = rhs39_
        d_313_v22_ = rhs40_
        r0 = default__.safeDivisionInt(d_291_v1_, (d_291_v1_) - (d_291_v1_))
        return r0

